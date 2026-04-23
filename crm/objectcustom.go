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

// ObjectCustomService contains methods and other services that help with
// interacting with the hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewObjectCustomService] method instead.
type ObjectCustomService struct {
	options []option.RequestOption
	Batch   ObjectCustomBatchService
}

// NewObjectCustomService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewObjectCustomService(opts ...option.RequestOption) (r ObjectCustomService) {
	r = ObjectCustomService{}
	r.options = opts
	r.Batch = NewObjectCustomBatchService(opts...)
	return
}

// Create a CRM object with the given properties and return a copy of the object,
// including the ID. Documentation and examples for creating standard objects is
// provided.
func (r *ObjectCustomService) New(ctx context.Context, objectType string, body ObjectCustomNewParams, opts ...option.RequestOption) (res *SimplePublicObject, err error) {
	opts = slices.Concat(r.options, opts)
	if objectType == "" {
		err = errors.New("missing required objectType parameter")
		return nil, err
	}
	path := fmt.Sprintf("crm/objects/2026-03/%s", url.PathEscape(objectType))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Perform a partial update of an Object identified by `{objectId}`or optionally a
// unique property value as specified by the `idProperty` query param. `{objectId}`
// refers to the internal object ID by default, and the `idProperty` query param
// refers to a property whose values are unique for the object. Provided property
// values will be overwritten. Read-only and non-existent properties will result in
// an error. Properties values can be cleared by passing an empty string.
func (r *ObjectCustomService) Update(ctx context.Context, objectID string, params ObjectCustomUpdateParams, opts ...option.RequestOption) (res *SimplePublicObject, err error) {
	opts = slices.Concat(r.options, opts)
	if params.ObjectType == "" {
		err = errors.New("missing required objectType parameter")
		return nil, err
	}
	if objectID == "" {
		err = errors.New("missing required objectId parameter")
		return nil, err
	}
	path := fmt.Sprintf("crm/objects/2026-03/%s/%s", url.PathEscape(params.ObjectType), url.PathEscape(objectID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return res, err
}

// Read a page of objects. Control what is returned via the `properties` query
// param.
func (r *ObjectCustomService) List(ctx context.Context, objectType string, query ObjectCustomListParams, opts ...option.RequestOption) (res *pagination.Page[SimplePublicObjectWithAssociations], err error) {
	var raw *http.Response
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if objectType == "" {
		err = errors.New("missing required objectType parameter")
		return nil, err
	}
	path := fmt.Sprintf("crm/objects/2026-03/%s", url.PathEscape(objectType))
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

// Read a page of objects. Control what is returned via the `properties` query
// param.
func (r *ObjectCustomService) ListAutoPaging(ctx context.Context, objectType string, query ObjectCustomListParams, opts ...option.RequestOption) *pagination.PageAutoPager[SimplePublicObjectWithAssociations] {
	return pagination.NewPageAutoPager(r.List(ctx, objectType, query, opts...))
}

// Move an Object identified by `{objectId}` to the recycling bin.
func (r *ObjectCustomService) Delete(ctx context.Context, objectID string, body ObjectCustomDeleteParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if body.ObjectType == "" {
		err = errors.New("missing required objectType parameter")
		return err
	}
	if objectID == "" {
		err = errors.New("missing required objectId parameter")
		return err
	}
	path := fmt.Sprintf("crm/objects/2026-03/%s/%s", url.PathEscape(body.ObjectType), url.PathEscape(objectID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Read an Object identified by `{objectId}`. `{objectId}` refers to the internal
// object ID by default, or optionally any unique property value as specified by
// the `idProperty` query param. Control what is returned via the `properties`
// query param.
func (r *ObjectCustomService) Get(ctx context.Context, objectID string, params ObjectCustomGetParams, opts ...option.RequestOption) (res *SimplePublicObjectWithAssociations, err error) {
	opts = slices.Concat(r.options, opts)
	if params.ObjectType == "" {
		err = errors.New("missing required objectType parameter")
		return nil, err
	}
	if objectID == "" {
		err = errors.New("missing required objectId parameter")
		return nil, err
	}
	path := fmt.Sprintf("crm/objects/2026-03/%s/%s", url.PathEscape(params.ObjectType), url.PathEscape(objectID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// Merge two CRM objects of the same type by specifying one as the primary object
// and the other as the object to be merged into it.
func (r *ObjectCustomService) Merge(ctx context.Context, objectType string, body ObjectCustomMergeParams, opts ...option.RequestOption) (res *SimplePublicObject, err error) {
	opts = slices.Concat(r.options, opts)
	if objectType == "" {
		err = errors.New("missing required objectType parameter")
		return nil, err
	}
	path := fmt.Sprintf("crm/objects/2026-03/%s/merge", url.PathEscape(objectType))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

func (r *ObjectCustomService) Search(ctx context.Context, objectType string, body ObjectCustomSearchParams, opts ...option.RequestOption) (res *CollectionResponseWithTotalSimplePublicObject, err error) {
	opts = slices.Concat(r.options, opts)
	if objectType == "" {
		err = errors.New("missing required objectType parameter")
		return nil, err
	}
	path := fmt.Sprintf("crm/objects/2026-03/%s/search", url.PathEscape(objectType))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type ObjectCustomNewParams struct {
	// Is the input object used to create a new CRM object, containing the properties
	// to be set and optional associations to link the new record with other CRM
	// objects.
	SimplePublicObjectInputForCreate SimplePublicObjectInputForCreateParam
	paramObj
}

func (r ObjectCustomNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.SimplePublicObjectInputForCreate)
}
func (r *ObjectCustomNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ObjectCustomUpdateParams struct {
	ObjectType string `path:"objectType" api:"required" json:"-"`
	// Represents the input required to create or update a CRM object, containing an
	// object with property names and their corresponding values.
	SimplePublicObjectInput SimplePublicObjectInputParam
	// The name of a property whose values are unique for this object type
	IDProperty param.Opt[string] `query:"idProperty,omitzero" json:"-"`
	paramObj
}

func (r ObjectCustomUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.SimplePublicObjectInput)
}
func (r *ObjectCustomUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [ObjectCustomUpdateParams]'s query parameters as
// `url.Values`.
func (r ObjectCustomUpdateParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ObjectCustomListParams struct {
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

// URLQuery serializes [ObjectCustomListParams]'s query parameters as `url.Values`.
func (r ObjectCustomListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ObjectCustomDeleteParams struct {
	ObjectType string `path:"objectType" api:"required" json:"-"`
	paramObj
}

type ObjectCustomGetParams struct {
	ObjectType string `path:"objectType" api:"required" json:"-"`
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

// URLQuery serializes [ObjectCustomGetParams]'s query parameters as `url.Values`.
func (r ObjectCustomGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ObjectCustomMergeParams struct {
	// Input data for merging two records.
	PublicMergeInput PublicMergeInputParam
	paramObj
}

func (r ObjectCustomMergeParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.PublicMergeInput)
}
func (r *ObjectCustomMergeParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ObjectCustomSearchParams struct {
	// Describes a search request
	PublicObjectSearchRequest PublicObjectSearchRequestParam
	paramObj
}

func (r ObjectCustomSearchParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.PublicObjectSearchRequest)
}
func (r *ObjectCustomSearchParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
