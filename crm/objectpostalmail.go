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

// ObjectPostalMailService contains methods and other services that help with
// interacting with the hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewObjectPostalMailService] method instead.
type ObjectPostalMailService struct {
	Options []option.RequestOption
	Batch   ObjectPostalMailBatchService
}

// NewObjectPostalMailService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewObjectPostalMailService(opts ...option.RequestOption) (r ObjectPostalMailService) {
	r = ObjectPostalMailService{}
	r.Options = opts
	r.Batch = NewObjectPostalMailBatchService(opts...)
	return
}

// Create a postal mail object with the given properties and return a copy of the
// object, including the ID.
func (r *ObjectPostalMailService) New(ctx context.Context, body ObjectPostalMailNewParams, opts ...option.RequestOption) (res *CreatedResponseSimplePublicObject, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "crm/v3/objects/postal_mail"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

func (r *ObjectPostalMailService) Update(ctx context.Context, postalMailID string, params ObjectPostalMailUpdateParams, opts ...option.RequestOption) (res *SimplePublicObject, err error) {
	opts = slices.Concat(r.Options, opts)
	if postalMailID == "" {
		err = errors.New("missing required postalMailId parameter")
		return
	}
	path := fmt.Sprintf("crm/v3/objects/postal_mail/%s", postalMailID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return
}

func (r *ObjectPostalMailService) List(ctx context.Context, query ObjectPostalMailListParams, opts ...option.RequestOption) (res *pagination.Page[SimplePublicObjectWithAssociations], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "crm/v3/objects/postal_mail"
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

func (r *ObjectPostalMailService) ListAutoPaging(ctx context.Context, query ObjectPostalMailListParams, opts ...option.RequestOption) *pagination.PageAutoPager[SimplePublicObjectWithAssociations] {
	return pagination.NewPageAutoPager(r.List(ctx, query, opts...))
}

// Move the postal mail object with the ID `{postalMailId}` to the recycling bin.
func (r *ObjectPostalMailService) Delete(ctx context.Context, postalMailID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if postalMailID == "" {
		err = errors.New("missing required postalMailId parameter")
		return
	}
	path := fmt.Sprintf("crm/v3/objects/postal_mail/%s", postalMailID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

func (r *ObjectPostalMailService) Get(ctx context.Context, postalMailID string, query ObjectPostalMailGetParams, opts ...option.RequestOption) (res *SimplePublicObjectWithAssociations, err error) {
	opts = slices.Concat(r.Options, opts)
	if postalMailID == "" {
		err = errors.New("missing required postalMailId parameter")
		return
	}
	path := fmt.Sprintf("crm/v3/objects/postal_mail/%s", postalMailID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Search for postal mail objects using specific criteria in the request.
func (r *ObjectPostalMailService) Search(ctx context.Context, body ObjectPostalMailSearchParams, opts ...option.RequestOption) (res *CollectionResponseWithTotalSimplePublicObject, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "crm/v3/objects/postal_mail/search"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type ObjectPostalMailNewParams struct {
	// Is the input object used to create a new CRM object, containing the properties
	// to be set and optional associations to link the new record with other CRM
	// objects.
	SimplePublicObjectInputForCreate SimplePublicObjectInputForCreateParam
	paramObj
}

func (r ObjectPostalMailNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.SimplePublicObjectInputForCreate)
}
func (r *ObjectPostalMailNewParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.SimplePublicObjectInputForCreate)
}

type ObjectPostalMailUpdateParams struct {
	// Represents the input required to create or update a CRM object, containing an
	// object with property names and their corresponding values.
	SimplePublicObjectInput SimplePublicObjectInputParam
	IDProperty              param.Opt[string] `query:"idProperty,omitzero" json:"-"`
	paramObj
}

func (r ObjectPostalMailUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.SimplePublicObjectInput)
}
func (r *ObjectPostalMailUpdateParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.SimplePublicObjectInput)
}

// URLQuery serializes [ObjectPostalMailUpdateParams]'s query parameters as
// `url.Values`.
func (r ObjectPostalMailUpdateParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ObjectPostalMailListParams struct {
	After                 param.Opt[string] `query:"after,omitzero" json:"-"`
	Archived              param.Opt[bool]   `query:"archived,omitzero" json:"-"`
	Limit                 param.Opt[int64]  `query:"limit,omitzero" json:"-"`
	Associations          []string          `query:"associations,omitzero" json:"-"`
	Properties            []string          `query:"properties,omitzero" json:"-"`
	PropertiesWithHistory []string          `query:"propertiesWithHistory,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ObjectPostalMailListParams]'s query parameters as
// `url.Values`.
func (r ObjectPostalMailListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ObjectPostalMailGetParams struct {
	Archived              param.Opt[bool]   `query:"archived,omitzero" json:"-"`
	IDProperty            param.Opt[string] `query:"idProperty,omitzero" json:"-"`
	Associations          []string          `query:"associations,omitzero" json:"-"`
	Properties            []string          `query:"properties,omitzero" json:"-"`
	PropertiesWithHistory []string          `query:"propertiesWithHistory,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ObjectPostalMailGetParams]'s query parameters as
// `url.Values`.
func (r ObjectPostalMailGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ObjectPostalMailSearchParams struct {
	// Describes a search request
	PublicObjectSearchRequest PublicObjectSearchRequestParam
	paramObj
}

func (r ObjectPostalMailSearchParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.PublicObjectSearchRequest)
}
func (r *ObjectPostalMailSearchParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.PublicObjectSearchRequest)
}
