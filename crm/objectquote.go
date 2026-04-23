// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package crm

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/stainless-sdks/hubspot-sdk-go/internal/apijson"
	"github.com/stainless-sdks/hubspot-sdk-go/internal/apiquery"
	shimjson "github.com/stainless-sdks/hubspot-sdk-go/internal/encoding/json"
	"github.com/stainless-sdks/hubspot-sdk-go/internal/requestconfig"
	"github.com/stainless-sdks/hubspot-sdk-go/option"
	"github.com/stainless-sdks/hubspot-sdk-go/packages/pagination"
	"github.com/stainless-sdks/hubspot-sdk-go/packages/param"
)

// ObjectQuoteService contains methods and other services that help with
// interacting with the hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewObjectQuoteService] method instead.
type ObjectQuoteService struct {
	options []option.RequestOption
	Batch   ObjectQuoteBatchService
}

// NewObjectQuoteService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewObjectQuoteService(opts ...option.RequestOption) (r ObjectQuoteService) {
	r = ObjectQuoteService{}
	r.options = opts
	r.Batch = NewObjectQuoteBatchService(opts...)
	return
}

// Create a quote with the given properties and return a copy of the object,
// including the ID. Documentation and examples for creating standard quotes is
// provided.
func (r *ObjectQuoteService) New(ctx context.Context, body ObjectQuoteNewParams, opts ...option.RequestOption) (res *SimplePublicObject, err error) {
	opts = slices.Concat(r.options, opts)
	path := "crm/objects/2026-03/quotes"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Perform a partial update of an Object identified by `{quoteId}`or optionally a
// unique property value as specified by the `idProperty` query param. `{quoteId}`
// refers to the internal object ID by default, and the `idProperty` query param
// refers to a property whose values are unique for the object. Provided property
// values will be overwritten. Read-only and non-existent properties will result in
// an error. Properties values can be cleared by passing an empty string.
func (r *ObjectQuoteService) Update(ctx context.Context, quoteID string, params ObjectQuoteUpdateParams, opts ...option.RequestOption) (res *SimplePublicObject, err error) {
	opts = slices.Concat(r.options, opts)
	if quoteID == "" {
		err = errors.New("missing required quoteId parameter")
		return nil, err
	}
	path := fmt.Sprintf("crm/objects/2026-03/quotes/%s", url.PathEscape(quoteID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return res, err
}

// Read a page of quotes. Control what is returned via the `properties` query
// param.
func (r *ObjectQuoteService) List(ctx context.Context, query ObjectQuoteListParams, opts ...option.RequestOption) (res *pagination.Page[SimplePublicObjectWithAssociations], err error) {
	var raw *http.Response
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "crm/objects/2026-03/quotes"
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

// Read a page of quotes. Control what is returned via the `properties` query
// param.
func (r *ObjectQuoteService) ListAutoPaging(ctx context.Context, query ObjectQuoteListParams, opts ...option.RequestOption) *pagination.PageAutoPager[SimplePublicObjectWithAssociations] {
	return pagination.NewPageAutoPager(r.List(ctx, query, opts...))
}

// Move an Object identified by `{quoteId}` to the recycling bin.
func (r *ObjectQuoteService) Delete(ctx context.Context, quoteID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if quoteID == "" {
		err = errors.New("missing required quoteId parameter")
		return err
	}
	path := fmt.Sprintf("crm/objects/2026-03/quotes/%s", url.PathEscape(quoteID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Read an Object identified by `{quoteId}`. `{quoteId}` refers to the internal
// object ID by default, or optionally any unique property value as specified by
// the `idProperty` query param. Control what is returned via the `properties`
// query param.
func (r *ObjectQuoteService) Get(ctx context.Context, quoteID string, query ObjectQuoteGetParams, opts ...option.RequestOption) (res *SimplePublicObjectWithAssociations, err error) {
	opts = slices.Concat(r.options, opts)
	if quoteID == "" {
		err = errors.New("missing required quoteId parameter")
		return nil, err
	}
	path := fmt.Sprintf("crm/objects/2026-03/quotes/%s", url.PathEscape(quoteID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Execute a search for quotes based on the criteria defined in the request body,
// such as filters, properties, and sorting options. This endpoint allows for
// detailed querying of quote records to retrieve specific data sets.
func (r *ObjectQuoteService) Search(ctx context.Context, body ObjectQuoteSearchParams, opts ...option.RequestOption) (res *CollectionResponseWithTotalSimplePublicObject, err error) {
	opts = slices.Concat(r.options, opts)
	path := "crm/objects/2026-03/quotes/search"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type ObjectQuoteNewParams struct {
	// Is the input object used to create a new CRM object, containing the properties
	// to be set and optional associations to link the new record with other CRM
	// objects.
	SimplePublicObjectInputForCreate SimplePublicObjectInputForCreateParam
	paramObj
}

func (r ObjectQuoteNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.SimplePublicObjectInputForCreate)
}
func (r *ObjectQuoteNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ObjectQuoteUpdateParams struct {
	// Represents the input required to create or update a CRM object, containing an
	// object with property names and their corresponding values.
	SimplePublicObjectInput SimplePublicObjectInputParam
	// The name of a property whose values are unique for this object type
	IDProperty param.Opt[string] `query:"idProperty,omitzero" json:"-"`
	paramObj
}

func (r ObjectQuoteUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.SimplePublicObjectInput)
}
func (r *ObjectQuoteUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [ObjectQuoteUpdateParams]'s query parameters as
// `url.Values`.
func (r ObjectQuoteUpdateParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ObjectQuoteListParams struct {
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
	// the maximum number of objects that can be read by a single request.
	PropertiesWithHistory []string `query:"propertiesWithHistory,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ObjectQuoteListParams]'s query parameters as `url.Values`.
func (r ObjectQuoteListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ObjectQuoteGetParams struct {
	// Whether to return only results that have been archived.
	Archived param.Opt[bool] `query:"archived,omitzero" json:"-"`
	// The name of a property whose values are unique for this object type
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

// URLQuery serializes [ObjectQuoteGetParams]'s query parameters as `url.Values`.
func (r ObjectQuoteGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ObjectQuoteSearchParams struct {
	// Describes a search request
	PublicObjectSearchRequest PublicObjectSearchRequestParam
	paramObj
}

func (r ObjectQuoteSearchParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.PublicObjectSearchRequest)
}
func (r *ObjectQuoteSearchParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
