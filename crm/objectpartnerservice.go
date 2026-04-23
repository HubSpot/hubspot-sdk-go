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

// ObjectPartnerServiceService contains methods and other services that help with
// interacting with the hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewObjectPartnerServiceService] method instead.
type ObjectPartnerServiceService struct {
	options []option.RequestOption
	Batch   ObjectPartnerServiceBatchService
}

// NewObjectPartnerServiceService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewObjectPartnerServiceService(opts ...option.RequestOption) (r ObjectPartnerServiceService) {
	r = ObjectPartnerServiceService{}
	r.options = opts
	r.Batch = NewObjectPartnerServiceBatchService(opts...)
	return
}

// Perform a partial update of an Object identified by `{partnerServiceId}`or
// optionally a unique property value as specified by the `idProperty` query param.
// `{partnerServiceId}` refers to the internal object ID by default, and the
// `idProperty` query param refers to a property whose values are unique for the
// object. Provided property values will be overwritten. Read-only and non-existent
// properties will result in an error. Properties values can be cleared by passing
// an empty string.
func (r *ObjectPartnerServiceService) Update(ctx context.Context, partnerServiceID string, params ObjectPartnerServiceUpdateParams, opts ...option.RequestOption) (res *SimplePublicObject, err error) {
	opts = slices.Concat(r.options, opts)
	if partnerServiceID == "" {
		err = errors.New("missing required partnerServiceId parameter")
		return nil, err
	}
	path := fmt.Sprintf("crm/objects/2026-03/partner_services/%s", url.PathEscape(partnerServiceID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return res, err
}

// Retrieve a list of associations for a specific partner service, filtered by the
// type of associated object.
func (r *ObjectPartnerServiceService) List(ctx context.Context, toObjectType string, params ObjectPartnerServiceListParams, opts ...option.RequestOption) (res *pagination.Page[MultiAssociatedObjectWithLabel], err error) {
	var raw *http.Response
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.PartnerServiceID == "" {
		err = errors.New("missing required partnerServiceId parameter")
		return nil, err
	}
	if toObjectType == "" {
		err = errors.New("missing required toObjectType parameter")
		return nil, err
	}
	path := fmt.Sprintf("crm/objects/2026-03/partner_services/%s/associations/%s", url.PathEscape(params.PartnerServiceID), url.PathEscape(toObjectType))
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, params, &res, opts...)
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

// Retrieve a list of associations for a specific partner service, filtered by the
// type of associated object.
func (r *ObjectPartnerServiceService) ListAutoPaging(ctx context.Context, toObjectType string, params ObjectPartnerServiceListParams, opts ...option.RequestOption) *pagination.PageAutoPager[MultiAssociatedObjectWithLabel] {
	return pagination.NewPageAutoPager(r.List(ctx, toObjectType, params, opts...))
}

// Read an Object identified by `{partnerServiceId}`. `{partnerServiceId}` refers
// to the internal object ID by default, or optionally any unique property value as
// specified by the `idProperty` query param. Control what is returned via the
// `properties` query param.
func (r *ObjectPartnerServiceService) Get(ctx context.Context, partnerServiceID string, query ObjectPartnerServiceGetParams, opts ...option.RequestOption) (res *SimplePublicObjectWithAssociations, err error) {
	opts = slices.Concat(r.options, opts)
	if partnerServiceID == "" {
		err = errors.New("missing required partnerServiceId parameter")
		return nil, err
	}
	path := fmt.Sprintf("crm/objects/2026-03/partner_services/%s", url.PathEscape(partnerServiceID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Execute a search query to find partner services based on defined filters,
// properties, and sorting options. This endpoint allows you to retrieve a
// collection of partner services that match the specified search criteria.
func (r *ObjectPartnerServiceService) Search(ctx context.Context, body ObjectPartnerServiceSearchParams, opts ...option.RequestOption) (res *CollectionResponseWithTotalSimplePublicObject, err error) {
	opts = slices.Concat(r.options, opts)
	path := "crm/objects/2026-03/partner_services/search"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type ObjectPartnerServiceUpdateParams struct {
	// Represents the input required to create or update a CRM object, containing an
	// object with property names and their corresponding values.
	SimplePublicObjectInput SimplePublicObjectInputParam
	// The name of a property whose values are unique for this object type
	IDProperty param.Opt[string] `query:"idProperty,omitzero" json:"-"`
	paramObj
}

func (r ObjectPartnerServiceUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.SimplePublicObjectInput)
}
func (r *ObjectPartnerServiceUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [ObjectPartnerServiceUpdateParams]'s query parameters as
// `url.Values`.
func (r ObjectPartnerServiceUpdateParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ObjectPartnerServiceListParams struct {
	PartnerServiceID string `path:"partnerServiceId" api:"required" json:"-"`
	// The paging cursor token of the last successfully read resource will be returned
	// as the `paging.next.after` JSON property of a paged response containing more
	// results.
	After param.Opt[string] `query:"after,omitzero" json:"-"`
	// The maximum number of results to display per page.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ObjectPartnerServiceListParams]'s query parameters as
// `url.Values`.
func (r ObjectPartnerServiceListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ObjectPartnerServiceGetParams struct {
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

// URLQuery serializes [ObjectPartnerServiceGetParams]'s query parameters as
// `url.Values`.
func (r ObjectPartnerServiceGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ObjectPartnerServiceSearchParams struct {
	// Describes a search request
	PublicObjectSearchRequest PublicObjectSearchRequestParam
	paramObj
}

func (r ObjectPartnerServiceSearchParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.PublicObjectSearchRequest)
}
func (r *ObjectPartnerServiceSearchParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
