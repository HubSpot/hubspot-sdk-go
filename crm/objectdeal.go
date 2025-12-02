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

// ObjectDealService contains methods and other services that help with interacting
// with the hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewObjectDealService] method instead.
type ObjectDealService struct {
	Options []option.RequestOption
	Batch   ObjectDealBatchService
}

// NewObjectDealService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewObjectDealService(opts ...option.RequestOption) (r ObjectDealService) {
	r = ObjectDealService{}
	r.Options = opts
	r.Batch = NewObjectDealBatchService(opts...)
	return
}

// Create a deal with the given properties and return a copy of the object,
// including the ID. Documentation and examples for creating standard deals is
// provided.
func (r *ObjectDealService) New(ctx context.Context, body ObjectDealNewParams, opts ...option.RequestOption) (res *CreatedResponseSimplePublicObject, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "crm/v3/objects/0-3"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Perform a partial update of an Object identified by `{dealId}`or optionally a
// unique property value as specified by the `idProperty` query param. `{dealId}`
// refers to the internal object ID by default, and the `idProperty` query param
// refers to a property whose values are unique for the object. Provided property
// values will be overwritten. Read-only and non-existent properties will result in
// an error. Properties values can be cleared by passing an empty string.
func (r *ObjectDealService) Update(ctx context.Context, dealID string, params ObjectDealUpdateParams, opts ...option.RequestOption) (res *SimplePublicObject, err error) {
	opts = slices.Concat(r.Options, opts)
	if dealID == "" {
		err = errors.New("missing required dealId parameter")
		return
	}
	path := fmt.Sprintf("crm/v3/objects/0-3/%s", dealID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return
}

// Read a page of deals. Control what is returned via the `properties` query param.
func (r *ObjectDealService) List(ctx context.Context, query ObjectDealListParams, opts ...option.RequestOption) (res *pagination.Page[SimplePublicObjectWithAssociations], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "crm/v3/objects/0-3"
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

// Read a page of deals. Control what is returned via the `properties` query param.
func (r *ObjectDealService) ListAutoPaging(ctx context.Context, query ObjectDealListParams, opts ...option.RequestOption) *pagination.PageAutoPager[SimplePublicObjectWithAssociations] {
	return pagination.NewPageAutoPager(r.List(ctx, query, opts...))
}

// Move an Object identified by `{dealId}` to the recycling bin.
func (r *ObjectDealService) Delete(ctx context.Context, dealID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if dealID == "" {
		err = errors.New("missing required dealId parameter")
		return
	}
	path := fmt.Sprintf("crm/v3/objects/0-3/%s", dealID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Read an Object identified by `{dealId}`. `{dealId}` refers to the internal
// object ID by default, or optionally any unique property value as specified by
// the `idProperty` query param. Control what is returned via the `properties`
// query param.
func (r *ObjectDealService) Get(ctx context.Context, dealID string, query ObjectDealGetParams, opts ...option.RequestOption) (res *SimplePublicObjectWithAssociations, err error) {
	opts = slices.Concat(r.Options, opts)
	if dealID == "" {
		err = errors.New("missing required dealId parameter")
		return
	}
	path := fmt.Sprintf("crm/v3/objects/0-3/%s", dealID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Combine two deals of the same type into a single deal.
func (r *ObjectDealService) Merge(ctx context.Context, body ObjectDealMergeParams, opts ...option.RequestOption) (res *SimplePublicObject, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "crm/v3/objects/0-3/merge"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Search for deals using specified criteria and filters.
func (r *ObjectDealService) Search(ctx context.Context, body ObjectDealSearchParams, opts ...option.RequestOption) (res *CollectionResponseWithTotalSimplePublicObject, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "crm/v3/objects/0-3/search"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type ObjectDealNewParams struct {
	// Is the input object used to create a new CRM object, containing the properties
	// to be set and optional associations to link the new record with other CRM
	// objects.
	SimplePublicObjectInputForCreate SimplePublicObjectInputForCreateParam
	paramObj
}

func (r ObjectDealNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.SimplePublicObjectInputForCreate)
}
func (r *ObjectDealNewParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.SimplePublicObjectInputForCreate)
}

type ObjectDealUpdateParams struct {
	// Represents the input required to create or update a CRM object, containing an
	// object with property names and their corresponding values.
	SimplePublicObjectInput SimplePublicObjectInputParam
	// The name of a property whose values are unique for this object
	IDProperty param.Opt[string] `query:"idProperty,omitzero" json:"-"`
	paramObj
}

func (r ObjectDealUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.SimplePublicObjectInput)
}
func (r *ObjectDealUpdateParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.SimplePublicObjectInput)
}

// URLQuery serializes [ObjectDealUpdateParams]'s query parameters as `url.Values`.
func (r ObjectDealUpdateParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ObjectDealListParams struct {
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
	// the maximum number of deals that can be read by a single request.
	PropertiesWithHistory []string `query:"propertiesWithHistory,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ObjectDealListParams]'s query parameters as `url.Values`.
func (r ObjectDealListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ObjectDealGetParams struct {
	// Whether to return only results that have been archived.
	Archived param.Opt[bool] `query:"archived,omitzero" json:"-"`
	// The name of a property whose values are unique for this object
	IDProperty param.Opt[string] `query:"idProperty,omitzero" json:"-"`
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

// URLQuery serializes [ObjectDealGetParams]'s query parameters as `url.Values`.
func (r ObjectDealGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ObjectDealMergeParams struct {
	PublicMergeInput PublicMergeInputParam
	paramObj
}

func (r ObjectDealMergeParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.PublicMergeInput)
}
func (r *ObjectDealMergeParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.PublicMergeInput)
}

type ObjectDealSearchParams struct {
	// Describes a search request
	PublicObjectSearchRequest PublicObjectSearchRequestParam
	paramObj
}

func (r ObjectDealSearchParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.PublicObjectSearchRequest)
}
func (r *ObjectDealSearchParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.PublicObjectSearchRequest)
}
