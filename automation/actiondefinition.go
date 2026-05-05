// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package automation

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/HubSpot/hubspot-sdk-go/internal/apijson"
	"github.com/HubSpot/hubspot-sdk-go/internal/apiquery"
	shimjson "github.com/HubSpot/hubspot-sdk-go/internal/encoding/json"
	"github.com/HubSpot/hubspot-sdk-go/internal/requestconfig"
	"github.com/HubSpot/hubspot-sdk-go/option"
	"github.com/HubSpot/hubspot-sdk-go/packages/pagination"
	"github.com/HubSpot/hubspot-sdk-go/packages/param"
)

// ActionDefinitionService contains methods and other services that help with
// interacting with the hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewActionDefinitionService] method instead.
type ActionDefinitionService struct {
	options []option.RequestOption
}

// NewActionDefinitionService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewActionDefinitionService(opts ...option.RequestOption) (r ActionDefinitionService) {
	r = ActionDefinitionService{}
	r.options = opts
	return
}

// Create a new custom workflow action.
func (r *ActionDefinitionService) New(ctx context.Context, appID int64, body ActionDefinitionNewParams, opts ...option.RequestOption) (res *PublicActionDefinition, err error) {
	opts = slices.Concat(r.options, opts)
	path := fmt.Sprintf("automation/actions/2026-03/%v", appID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Update an existing action definition by ID.
func (r *ActionDefinitionService) Update(ctx context.Context, definitionID string, params ActionDefinitionUpdateParams, opts ...option.RequestOption) (res *PublicActionDefinition, err error) {
	opts = slices.Concat(r.options, opts)
	if definitionID == "" {
		err = errors.New("missing required definitionId parameter")
		return nil, err
	}
	path := fmt.Sprintf("automation/actions/2026-03/%v/%s", params.AppID, url.PathEscape(definitionID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return res, err
}

// Retrieve custom workflow action definitions by app ID.
func (r *ActionDefinitionService) List(ctx context.Context, appID int64, query ActionDefinitionListParams, opts ...option.RequestOption) (res *pagination.Page[PublicActionDefinition], err error) {
	var raw *http.Response
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := fmt.Sprintf("automation/actions/2026-03/%v", appID)
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

// Retrieve custom workflow action definitions by app ID.
func (r *ActionDefinitionService) ListAutoPaging(ctx context.Context, appID int64, query ActionDefinitionListParams, opts ...option.RequestOption) *pagination.PageAutoPager[PublicActionDefinition] {
	return pagination.NewPageAutoPager(r.List(ctx, appID, query, opts...))
}

// Delete an action definition by ID.
func (r *ActionDefinitionService) Delete(ctx context.Context, definitionID string, body ActionDefinitionDeleteParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if definitionID == "" {
		err = errors.New("missing required definitionId parameter")
		return err
	}
	path := fmt.Sprintf("automation/actions/2026-03/%v/%s", body.AppID, url.PathEscape(definitionID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Set whether a custom action definition requires an object.
func (r *ActionDefinitionService) NewRequiresObject(ctx context.Context, definitionID string, params ActionDefinitionNewRequiresObjectParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if definitionID == "" {
		err = errors.New("missing required definitionId parameter")
		return err
	}
	path := fmt.Sprintf("automation/actions/2026-03/%v/%s/requires-object", params.AppID, url.PathEscape(definitionID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, nil, opts...)
	return err
}

// Retrieve a custom workflow action definition by ID.
func (r *ActionDefinitionService) Get(ctx context.Context, definitionID string, params ActionDefinitionGetParams, opts ...option.RequestOption) (res *PublicActionDefinition, err error) {
	opts = slices.Concat(r.options, opts)
	if definitionID == "" {
		err = errors.New("missing required definitionId parameter")
		return nil, err
	}
	path := fmt.Sprintf("automation/actions/2026-03/%v/%s", params.AppID, url.PathEscape(definitionID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// Retrieve whether a custom action definition requires an object.
func (r *ActionDefinitionService) GetRequiresObject(ctx context.Context, definitionID string, query ActionDefinitionGetRequiresObjectParams, opts ...option.RequestOption) (res *PublicActionDefinitionRequiresObjectResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if definitionID == "" {
		err = errors.New("missing required definitionId parameter")
		return nil, err
	}
	path := fmt.Sprintf("automation/actions/2026-03/%v/%s/requires-object", query.AppID, url.PathEscape(definitionID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type ActionDefinitionNewParams struct {
	PublicActionDefinitionEgg PublicActionDefinitionEggParam
	paramObj
}

func (r ActionDefinitionNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.PublicActionDefinitionEgg)
}
func (r *ActionDefinitionNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ActionDefinitionUpdateParams struct {
	AppID                       int64 `path:"appId" api:"required" json:"-"`
	PublicActionDefinitionPatch PublicActionDefinitionPatchParam
	paramObj
}

func (r ActionDefinitionUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.PublicActionDefinitionPatch)
}
func (r *ActionDefinitionUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ActionDefinitionListParams struct {
	// The paging cursor token of the last successfully read resource will be returned
	// as the `paging.next.after` JSON property of a paged response containing more
	// results.
	After param.Opt[string] `query:"after,omitzero" json:"-"`
	// Whether to return only results that have been archived.
	Archived param.Opt[bool] `query:"archived,omitzero" json:"-"`
	// The maximum number of results to display per page.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ActionDefinitionListParams]'s query parameters as
// `url.Values`.
func (r ActionDefinitionListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ActionDefinitionDeleteParams struct {
	AppID int64 `path:"appId" api:"required" json:"-"`
	paramObj
}

type ActionDefinitionNewRequiresObjectParams struct {
	AppID                                       int64 `path:"appId" api:"required" json:"-"`
	PublicActionDefinitionRequiresObjectRequest PublicActionDefinitionRequiresObjectRequestParam
	paramObj
}

func (r ActionDefinitionNewRequiresObjectParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.PublicActionDefinitionRequiresObjectRequest)
}
func (r *ActionDefinitionNewRequiresObjectParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ActionDefinitionGetParams struct {
	AppID int64 `path:"appId" api:"required" json:"-"`
	// Whether to return only results that have been archived.
	Archived param.Opt[bool] `query:"archived,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ActionDefinitionGetParams]'s query parameters as
// `url.Values`.
func (r ActionDefinitionGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ActionDefinitionGetRequiresObjectParams struct {
	AppID int64 `path:"appId" api:"required" json:"-"`
	paramObj
}
