// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package automation

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

// ActionDefinitionService contains methods and other services that help with
// interacting with the Hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewActionDefinitionService] method instead.
type ActionDefinitionService struct {
	Options []option.RequestOption
}

// NewActionDefinitionService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewActionDefinitionService(opts ...option.RequestOption) (r ActionDefinitionService) {
	r = ActionDefinitionService{}
	r.Options = opts
	return
}

// Create a new custom workflow action.
func (r *ActionDefinitionService) New(ctx context.Context, appID int64, body ActionDefinitionNewParams, opts ...option.RequestOption) (res *PublicActionDefinition, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("automation/v4/actions/%v", appID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Update an existing action definition by ID.
func (r *ActionDefinitionService) Update(ctx context.Context, definitionID string, params ActionDefinitionUpdateParams, opts ...option.RequestOption) (res *PublicActionDefinition, err error) {
	opts = slices.Concat(r.Options, opts)
	if definitionID == "" {
		err = errors.New("missing required definitionId parameter")
		return
	}
	path := fmt.Sprintf("automation/v4/actions/%v/%s", params.AppID, definitionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return
}

// Retrieve custom workflow action definitions by app ID.
func (r *ActionDefinitionService) List(ctx context.Context, appID int64, query ActionDefinitionListParams, opts ...option.RequestOption) (res *pagination.Page[PublicActionDefinition], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := fmt.Sprintf("automation/v4/actions/%v", appID)
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
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if definitionID == "" {
		err = errors.New("missing required definitionId parameter")
		return
	}
	path := fmt.Sprintf("automation/v4/actions/%v/%s", body.AppID, definitionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Retrieve a custom workflow action definition by ID.
func (r *ActionDefinitionService) Get(ctx context.Context, definitionID string, params ActionDefinitionGetParams, opts ...option.RequestOption) (res *PublicActionDefinition, err error) {
	opts = slices.Concat(r.Options, opts)
	if definitionID == "" {
		err = errors.New("missing required definitionId parameter")
		return
	}
	path := fmt.Sprintf("automation/v4/actions/%v/%s", params.AppID, definitionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return
}

type ActionDefinitionNewParams struct {
	PublicActionDefinitionEgg PublicActionDefinitionEggParam
	paramObj
}

func (r ActionDefinitionNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.PublicActionDefinitionEgg)
}
func (r *ActionDefinitionNewParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.PublicActionDefinitionEgg)
}

type ActionDefinitionUpdateParams struct {
	AppID                       int64 `path:"appId,required" json:"-"`
	PublicActionDefinitionPatch PublicActionDefinitionPatchParam
	paramObj
}

func (r ActionDefinitionUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.PublicActionDefinitionPatch)
}
func (r *ActionDefinitionUpdateParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.PublicActionDefinitionPatch)
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
	AppID int64 `path:"appId,required" json:"-"`
	paramObj
}

type ActionDefinitionGetParams struct {
	AppID int64 `path:"appId,required" json:"-"`
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
