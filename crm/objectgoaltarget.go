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

// ObjectGoalTargetService contains methods and other services that help with
// interacting with the hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewObjectGoalTargetService] method instead.
type ObjectGoalTargetService struct {
	options []option.RequestOption
	Batch   ObjectGoalTargetBatchService
}

// NewObjectGoalTargetService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewObjectGoalTargetService(opts ...option.RequestOption) (r ObjectGoalTargetService) {
	r = ObjectGoalTargetService{}
	r.options = opts
	r.Batch = NewObjectGoalTargetBatchService(opts...)
	return
}

// Create a goal target with the given properties and return a copy of the object,
// including the ID. Documentation and examples for creating standard goal targets
// is provided.
func (r *ObjectGoalTargetService) New(ctx context.Context, body ObjectGoalTargetNewParams, opts ...option.RequestOption) (res *SimplePublicObject, err error) {
	opts = slices.Concat(r.options, opts)
	path := "crm/objects/2026-03/goal_targets"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Perform a partial update of an Object identified by `{goalTargetId}`or
// optionally a unique property value as specified by the `idProperty` query param.
// `{goalTargetId}` refers to the internal object ID by default, and the
// `idProperty` query param refers to a property whose values are unique for the
// object. Provided property values will be overwritten. Read-only and non-existent
// properties will result in an error. Properties values can be cleared by passing
// an empty string.
func (r *ObjectGoalTargetService) Update(ctx context.Context, goalTargetID string, params ObjectGoalTargetUpdateParams, opts ...option.RequestOption) (res *SimplePublicObject, err error) {
	opts = slices.Concat(r.options, opts)
	if goalTargetID == "" {
		err = errors.New("missing required goalTargetId parameter")
		return nil, err
	}
	path := fmt.Sprintf("crm/objects/2026-03/goal_targets/%s", url.PathEscape(goalTargetID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return res, err
}

// Read a page of goal targets. Control what is returned via the `properties` query
// param.
func (r *ObjectGoalTargetService) List(ctx context.Context, query ObjectGoalTargetListParams, opts ...option.RequestOption) (res *pagination.Page[SimplePublicObjectWithAssociations], err error) {
	var raw *http.Response
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "crm/objects/2026-03/goal_targets"
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

// Read a page of goal targets. Control what is returned via the `properties` query
// param.
func (r *ObjectGoalTargetService) ListAutoPaging(ctx context.Context, query ObjectGoalTargetListParams, opts ...option.RequestOption) *pagination.PageAutoPager[SimplePublicObjectWithAssociations] {
	return pagination.NewPageAutoPager(r.List(ctx, query, opts...))
}

// Delete a goal target by `{goalTargetId}` to the recycling bin.
func (r *ObjectGoalTargetService) Delete(ctx context.Context, goalTargetID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if goalTargetID == "" {
		err = errors.New("missing required goalTargetId parameter")
		return err
	}
	path := fmt.Sprintf("crm/objects/2026-03/goal_targets/%s", url.PathEscape(goalTargetID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Retrieve a goal target by its ID. You can specify what is returned using the
// `properties` query parameter.
func (r *ObjectGoalTargetService) Get(ctx context.Context, goalTargetID string, query ObjectGoalTargetGetParams, opts ...option.RequestOption) (res *SimplePublicObjectWithAssociations, err error) {
	opts = slices.Concat(r.options, opts)
	if goalTargetID == "" {
		err = errors.New("missing required goalTargetId parameter")
		return nil, err
	}
	path := fmt.Sprintf("crm/objects/2026-03/goal_targets/%s", url.PathEscape(goalTargetID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Search for goal targets using specified criteria.
func (r *ObjectGoalTargetService) Search(ctx context.Context, body ObjectGoalTargetSearchParams, opts ...option.RequestOption) (res *CollectionResponseWithTotalSimplePublicObject, err error) {
	opts = slices.Concat(r.options, opts)
	path := "crm/objects/2026-03/goal_targets/search"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type ObjectGoalTargetNewParams struct {
	// Is the input object used to create a new CRM object, containing the properties
	// to be set and optional associations to link the new record with other CRM
	// objects.
	SimplePublicObjectInputForCreate SimplePublicObjectInputForCreateParam
	paramObj
}

func (r ObjectGoalTargetNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.SimplePublicObjectInputForCreate)
}
func (r *ObjectGoalTargetNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ObjectGoalTargetUpdateParams struct {
	// Represents the input required to create or update a CRM object, containing an
	// object with property names and their corresponding values.
	SimplePublicObjectInput SimplePublicObjectInputParam
	// The name of a property whose values are unique for this object type
	IDProperty param.Opt[string] `query:"idProperty,omitzero" json:"-"`
	paramObj
}

func (r ObjectGoalTargetUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.SimplePublicObjectInput)
}
func (r *ObjectGoalTargetUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [ObjectGoalTargetUpdateParams]'s query parameters as
// `url.Values`.
func (r ObjectGoalTargetUpdateParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ObjectGoalTargetListParams struct {
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

// URLQuery serializes [ObjectGoalTargetListParams]'s query parameters as
// `url.Values`.
func (r ObjectGoalTargetListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ObjectGoalTargetGetParams struct {
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

// URLQuery serializes [ObjectGoalTargetGetParams]'s query parameters as
// `url.Values`.
func (r ObjectGoalTargetGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ObjectGoalTargetSearchParams struct {
	// Describes a search request
	PublicObjectSearchRequest PublicObjectSearchRequestParam
	paramObj
}

func (r ObjectGoalTargetSearchParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.PublicObjectSearchRequest)
}
func (r *ObjectGoalTargetSearchParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
