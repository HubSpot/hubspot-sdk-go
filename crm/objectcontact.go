// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package crm

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/stainless-sdks/hubspot-sdk-go/internal/apiquery"
	shimjson "github.com/stainless-sdks/hubspot-sdk-go/internal/encoding/json"
	"github.com/stainless-sdks/hubspot-sdk-go/internal/requestconfig"
	"github.com/stainless-sdks/hubspot-sdk-go/option"
	"github.com/stainless-sdks/hubspot-sdk-go/packages/pagination"
	"github.com/stainless-sdks/hubspot-sdk-go/packages/param"
)

// ObjectContactService contains methods and other services that help with
// interacting with the Hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewObjectContactService] method instead.
type ObjectContactService struct {
	Options []option.RequestOption
	Batch   ObjectContactBatchService
}

// NewObjectContactService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewObjectContactService(opts ...option.RequestOption) (r ObjectContactService) {
	r = ObjectContactService{}
	r.Options = opts
	r.Batch = NewObjectContactBatchService(opts...)
	return
}

// Create a single contact. Include a `properties` object to define
// [property values](https://developers.hubspot.com/docs/guides/api/crm/properties)
// for the contact, along with an `associations` array to define
// [associations](https://developers.hubspot.com/docs/guides/api/crm/associations/associations-v4)
// with other CRM records.
func (r *ObjectContactService) New(ctx context.Context, body ObjectContactNewParams, opts ...option.RequestOption) (res *CreatedResponseSimplePublicObject, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "crm/v3/objects/contacts"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Update an existing contact, identified by ID or email/unique property value. To
// identify a contact by ID, include the ID in the request URL path. To identify a
// contact by their email or other unique property, include the email/property
// value in the request URL path, and add the `idProperty` query parameter
// (`/crm/v3/objects/contacts/jon@website.com?idProperty=email`). Provided property
// values will be overwritten. Read-only and non-existent properties will result in
// an error. Properties values can be cleared by passing an empty string.
func (r *ObjectContactService) Update(ctx context.Context, contactID string, body ObjectContactUpdateParams, opts ...option.RequestOption) (res *SimplePublicObject, err error) {
	opts = slices.Concat(r.Options, opts)
	if contactID == "" {
		err = errors.New("missing required contactId parameter")
		return
	}
	path := fmt.Sprintf("crm/v3/objects/contacts/%s", contactID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Retrieve all contacts, using query parameters to specify the information that
// gets returned.
func (r *ObjectContactService) List(ctx context.Context, query ObjectContactListParams, opts ...option.RequestOption) (res *pagination.Page[SimplePublicObjectWithAssociations], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "crm/v3/objects/contacts"
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// Retrieve all contacts, using query parameters to specify the information that
// gets returned.
func (r *ObjectContactService) ListAutoPaging(ctx context.Context, query ObjectContactListParams, opts ...option.RequestOption) *pagination.PageAutoPager[SimplePublicObjectWithAssociations] {
	return pagination.NewPageAutoPager(r.List(ctx, query, opts...))
}

// Delete a contact by ID. Deleted contacts can be restored within 90 days of
// deletion. Learn more about the
// [data impacted by contact deletions](https://knowledge.hubspot.com/privacy-and-consent/understand-restorable-and-permanent-contact-deletions)
// and how to
// [restore archived records](https://knowledge.hubspot.com/records/restore-deleted-records).
func (r *ObjectContactService) Delete(ctx context.Context, contactID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if contactID == "" {
		err = errors.New("missing required contactId parameter")
		return
	}
	path := fmt.Sprintf("crm/v3/objects/contacts/%s", contactID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Permanently delete a contact and all associated content to follow GDPR. Use
// optional property `idProperty` set to `email` to identify contact by email
// address. If email address is not found, the email address will be added to a
// blocklist and prevent it from being used in the future. Learn more about
// [permanently deleting contacts](https://knowledge.hubspot.com/privacy-and-consent/how-do-i-perform-a-gdpr-delete-in-hubspot).
func (r *ObjectContactService) GdprDelete(ctx context.Context, body ObjectContactGdprDeleteParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "crm/v3/objects/contacts/gdpr-delete"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Retrieve a contact by its ID (`contactId`) or by a unique property
// (`idProperty`). You can specify what is returned using the `properties` query
// parameter.
func (r *ObjectContactService) Get(ctx context.Context, contactID string, query ObjectContactGetParams, opts ...option.RequestOption) (res *SimplePublicObjectWithAssociations, err error) {
	opts = slices.Concat(r.Options, opts)
	if contactID == "" {
		err = errors.New("missing required contactId parameter")
		return
	}
	path := fmt.Sprintf("crm/v3/objects/contacts/%s", contactID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Merge two contact records. Learn more about
// [merging records](https://knowledge.hubspot.com/records/merge-records).
func (r *ObjectContactService) Merge(ctx context.Context, body ObjectContactMergeParams, opts ...option.RequestOption) (res *SimplePublicObject, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "crm/v3/objects/contacts/merge"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Search for contacts by filtering on properties, searching through associations,
// and sorting results. Learn more about
// [CRM search](https://developers.hubspot.com/docs/guides/api/crm/search#make-a-search-request).
func (r *ObjectContactService) Search(ctx context.Context, body ObjectContactSearchParams, opts ...option.RequestOption) (res *CollectionResponseWithTotalSimplePublicObject, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "crm/v3/objects/contacts/search"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type ObjectContactNewParams struct {
	// Is the input object used to create a new CRM object, containing the properties
	// to be set and optional associations to link the new record with other CRM
	// objects.
	SimplePublicObjectInputForCreate SimplePublicObjectInputForCreateParam
	paramObj
}

func (r ObjectContactNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.SimplePublicObjectInputForCreate)
}
func (r *ObjectContactNewParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.SimplePublicObjectInputForCreate)
}

type ObjectContactUpdateParams struct {
	// Represents the input required to create or update a CRM object, containing an
	// object with property names and their corresponding values.
	SimplePublicObjectInput SimplePublicObjectInputParam
	paramObj
}

func (r ObjectContactUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.SimplePublicObjectInput)
}
func (r *ObjectContactUpdateParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.SimplePublicObjectInput)
}

type ObjectContactListParams struct {
	// The paging cursor token of the last successfully read resource will be returned
	// as the `paging.next.after` JSON property of a paged response containing more
	// results.
	After param.Opt[string] `query:"after,omitzero" json:"-"`
	// Whether to return only results that have been archived.
	Archived param.Opt[bool] `query:"archived,omitzero" json:"-"`
	// The maximum number of results to display per page.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// A comma separated list of object types to retrieve associated IDs for. If any of
	// the specified associations do not exist, they will be ignored.
	Associations []string `query:"associations,omitzero" json:"-"`
	// A comma separated list of the properties to be returned in the response. If any
	// of the specified properties are not present on the requested object(s), they
	// will be ignored.
	Properties []string `query:"properties,omitzero" json:"-"`
	// A comma separated list of the properties to be returned along with their history
	// of previous values. If any of the specified properties are not present on the
	// requested object(s), they will be ignored. Usage of this parameter will reduce
	// the maximum number of contacts that can be read by a single request.
	PropertiesWithHistory []string `query:"propertiesWithHistory,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ObjectContactListParams]'s query parameters as
// `url.Values`.
func (r ObjectContactListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ObjectContactGdprDeleteParams struct {
	PublicGdprDeleteInput PublicGdprDeleteInputParam
	paramObj
}

func (r ObjectContactGdprDeleteParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.PublicGdprDeleteInput)
}
func (r *ObjectContactGdprDeleteParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.PublicGdprDeleteInput)
}

type ObjectContactGetParams struct {
	// Whether to return only results that have been archived.
	Archived param.Opt[bool] `query:"archived,omitzero" json:"-"`
	// A comma separated list of object types to retrieve associated IDs for. If any of
	// the specified associations do not exist, they will be ignored.
	Associations []string `query:"associations,omitzero" json:"-"`
	// A comma separated list of the properties to be returned in the response. If any
	// of the specified properties are not present on the requested object(s), they
	// will be ignored.
	Properties []string `query:"properties,omitzero" json:"-"`
	// A comma separated list of the properties to be returned along with their history
	// of previous values. If any of the specified properties are not present on the
	// requested object(s), they will be ignored.
	PropertiesWithHistory []string `query:"propertiesWithHistory,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ObjectContactGetParams]'s query parameters as `url.Values`.
func (r ObjectContactGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ObjectContactMergeParams struct {
	PublicMergeInput PublicMergeInputParam
	paramObj
}

func (r ObjectContactMergeParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.PublicMergeInput)
}
func (r *ObjectContactMergeParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.PublicMergeInput)
}

type ObjectContactSearchParams struct {
	// Describes a search request
	PublicObjectSearchRequest PublicObjectSearchRequestParam
	paramObj
}

func (r ObjectContactSearchParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.PublicObjectSearchRequest)
}
func (r *ObjectContactSearchParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.PublicObjectSearchRequest)
}
