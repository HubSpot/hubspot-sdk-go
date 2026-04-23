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

// ObjectPartnerClientService contains methods and other services that help with
// interacting with the hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewObjectPartnerClientService] method instead.
type ObjectPartnerClientService struct {
	options []option.RequestOption
	Batch   ObjectPartnerClientBatchService
}

// NewObjectPartnerClientService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewObjectPartnerClientService(opts ...option.RequestOption) (r ObjectPartnerClientService) {
	r = ObjectPartnerClientService{}
	r.options = opts
	r.Batch = NewObjectPartnerClientBatchService(opts...)
	return
}

// Update the specified properties of an existing partner client.
func (r *ObjectPartnerClientService) Update(ctx context.Context, partnerClientID string, params ObjectPartnerClientUpdateParams, opts ...option.RequestOption) (res *SimplePublicObject, err error) {
	opts = slices.Concat(r.options, opts)
	if partnerClientID == "" {
		err = errors.New("missing required partnerClientId parameter")
		return nil, err
	}
	path := fmt.Sprintf("crm/objects/2026-03/partner_clients/%s", url.PathEscape(partnerClientID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return res, err
}

// Retrieve a list of partner clients with optional filtering by deleted status,
// associations, and specific properties. The response can be paginated using the
// 'after' parameter.
func (r *ObjectPartnerClientService) List(ctx context.Context, query ObjectPartnerClientListParams, opts ...option.RequestOption) (res *pagination.Page[SimplePublicObjectWithAssociations], err error) {
	var raw *http.Response
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "crm/objects/2026-03/partner_clients"
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

// Retrieve a list of partner clients with optional filtering by deleted status,
// associations, and specific properties. The response can be paginated using the
// 'after' parameter.
func (r *ObjectPartnerClientService) ListAutoPaging(ctx context.Context, query ObjectPartnerClientListParams, opts ...option.RequestOption) *pagination.PageAutoPager[SimplePublicObjectWithAssociations] {
	return pagination.NewPageAutoPager(r.List(ctx, query, opts...))
}

// Retrieve detailed information about a specific partner client, including
// selected properties and associations. This endpoint is useful for accessing
// comprehensive client data for analysis or integration purposes.
func (r *ObjectPartnerClientService) Get(ctx context.Context, partnerClientID string, query ObjectPartnerClientGetParams, opts ...option.RequestOption) (res *SimplePublicObjectWithAssociations, err error) {
	opts = slices.Concat(r.options, opts)
	if partnerClientID == "" {
		err = errors.New("missing required partnerClientId parameter")
		return nil, err
	}
	path := fmt.Sprintf("crm/objects/2026-03/partner_clients/%s", url.PathEscape(partnerClientID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Retrieve a list of associations for a specific partner client based on the
// specified object type.
func (r *ObjectPartnerClientService) ListAssociations(ctx context.Context, toObjectType string, params ObjectPartnerClientListAssociationsParams, opts ...option.RequestOption) (res *pagination.Page[MultiAssociatedObjectWithLabel], err error) {
	var raw *http.Response
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.PartnerClientID == "" {
		err = errors.New("missing required partnerClientId parameter")
		return nil, err
	}
	if toObjectType == "" {
		err = errors.New("missing required toObjectType parameter")
		return nil, err
	}
	path := fmt.Sprintf("crm/objects/2026-03/partner_clients/%s/associations/%s", url.PathEscape(params.PartnerClientID), url.PathEscape(toObjectType))
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

// Retrieve a list of associations for a specific partner client based on the
// specified object type.
func (r *ObjectPartnerClientService) ListAssociationsAutoPaging(ctx context.Context, toObjectType string, params ObjectPartnerClientListAssociationsParams, opts ...option.RequestOption) *pagination.PageAutoPager[MultiAssociatedObjectWithLabel] {
	return pagination.NewPageAutoPager(r.ListAssociations(ctx, toObjectType, params, opts...))
}

// Execute a search for partner clients based on defined filters, properties, and
// sorting options. This endpoint allows you to retrieve partner client data that
// matches the search criteria, facilitating integration and data synchronization
// with third-party systems.
func (r *ObjectPartnerClientService) Search(ctx context.Context, body ObjectPartnerClientSearchParams, opts ...option.RequestOption) (res *CollectionResponseWithTotalSimplePublicObject, err error) {
	opts = slices.Concat(r.options, opts)
	path := "crm/objects/2026-03/partner_clients/search"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type ObjectPartnerClientUpdateParams struct {
	// Represents the input required to create or update a CRM object, containing an
	// object with property names and their corresponding values.
	SimplePublicObjectInput SimplePublicObjectInputParam
	// The name of a property whose values are unique for this object type
	IDProperty param.Opt[string] `query:"idProperty,omitzero" json:"-"`
	paramObj
}

func (r ObjectPartnerClientUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.SimplePublicObjectInput)
}
func (r *ObjectPartnerClientUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [ObjectPartnerClientUpdateParams]'s query parameters as
// `url.Values`.
func (r ObjectPartnerClientUpdateParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ObjectPartnerClientListParams struct {
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

// URLQuery serializes [ObjectPartnerClientListParams]'s query parameters as
// `url.Values`.
func (r ObjectPartnerClientListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ObjectPartnerClientGetParams struct {
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

// URLQuery serializes [ObjectPartnerClientGetParams]'s query parameters as
// `url.Values`.
func (r ObjectPartnerClientGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ObjectPartnerClientListAssociationsParams struct {
	PartnerClientID string `path:"partnerClientId" api:"required" json:"-"`
	// The paging cursor token of the last successfully read resource will be returned
	// as the `paging.next.after` JSON property of a paged response containing more
	// results.
	After param.Opt[string] `query:"after,omitzero" json:"-"`
	// The maximum number of results to display per page.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ObjectPartnerClientListAssociationsParams]'s query
// parameters as `url.Values`.
func (r ObjectPartnerClientListAssociationsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ObjectPartnerClientSearchParams struct {
	// Describes a search request
	PublicObjectSearchRequest PublicObjectSearchRequestParam
	paramObj
}

func (r ObjectPartnerClientSearchParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.PublicObjectSearchRequest)
}
func (r *ObjectPartnerClientSearchParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
