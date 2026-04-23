// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package automation

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/stainless-sdks/hubspot-sdk-go/internal/apijson"
	shimjson "github.com/stainless-sdks/hubspot-sdk-go/internal/encoding/json"
	"github.com/stainless-sdks/hubspot-sdk-go/internal/requestconfig"
	"github.com/stainless-sdks/hubspot-sdk-go/option"
)

// ActionFunctionService contains methods and other services that help with
// interacting with the hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewActionFunctionService] method instead.
type ActionFunctionService struct {
	options []option.RequestOption
}

// NewActionFunctionService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewActionFunctionService(opts ...option.RequestOption) (r ActionFunctionService) {
	r = ActionFunctionService{}
	r.options = opts
	return
}

// Retrieve all functions included in a definition.
func (r *ActionFunctionService) List(ctx context.Context, definitionID string, query ActionFunctionListParams, opts ...option.RequestOption) (res *CollectionResponsePublicActionFunctionIdentifierNoPaging, err error) {
	opts = slices.Concat(r.options, opts)
	if definitionID == "" {
		err = errors.New("missing required definitionId parameter")
		return nil, err
	}
	path := fmt.Sprintf("automation/actions/2026-03/%v/%s/functions", query.AppID, url.PathEscape(definitionID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Archive a function for a specific definition.
func (r *ActionFunctionService) Delete(ctx context.Context, functionID string, body ActionFunctionDeleteParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if body.DefinitionID == "" {
		err = errors.New("missing required definitionId parameter")
		return err
	}
	if functionID == "" {
		err = errors.New("missing required functionId parameter")
		return err
	}
	path := fmt.Sprintf("automation/actions/2026-03/%v/%s/functions/%v/%s", body.AppID, url.PathEscape(body.DefinitionID), body.FunctionType, url.PathEscape(functionID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Update a function for a given definition by ID.
func (r *ActionFunctionService) NewOrReplace(ctx context.Context, functionID string, params ActionFunctionNewOrReplaceParams, opts ...option.RequestOption) (res *PublicActionFunctionIdentifier, err error) {
	opts = slices.Concat(r.options, opts)
	if params.DefinitionID == "" {
		err = errors.New("missing required definitionId parameter")
		return nil, err
	}
	if functionID == "" {
		err = errors.New("missing required functionId parameter")
		return nil, err
	}
	path := fmt.Sprintf("automation/actions/2026-03/%v/%s/functions/%v/%s", params.AppID, url.PathEscape(params.DefinitionID), params.FunctionType, url.PathEscape(functionID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return res, err
}

// Add a function for a given definition.
func (r *ActionFunctionService) NewOrReplaceByFunctionType(ctx context.Context, functionType ActionFunctionNewOrReplaceByFunctionTypeParamsFunctionType, params ActionFunctionNewOrReplaceByFunctionTypeParams, opts ...option.RequestOption) (res *PublicActionFunctionIdentifier, err error) {
	opts = slices.Concat(r.options, opts)
	if params.DefinitionID == "" {
		err = errors.New("missing required definitionId parameter")
		return nil, err
	}
	path := fmt.Sprintf("automation/actions/2026-03/%v/%s/functions/%v", params.AppID, url.PathEscape(params.DefinitionID), functionType)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return res, err
}

// Delete a function within a given definition.
func (r *ActionFunctionService) DeleteByFunctionType(ctx context.Context, functionType ActionFunctionDeleteByFunctionTypeParamsFunctionType, body ActionFunctionDeleteByFunctionTypeParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if body.DefinitionID == "" {
		err = errors.New("missing required definitionId parameter")
		return err
	}
	path := fmt.Sprintf("automation/actions/2026-03/%v/%s/functions/%v", body.AppID, url.PathEscape(body.DefinitionID), functionType)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Retrieve a specific function from a given definition.
func (r *ActionFunctionService) Get(ctx context.Context, functionID string, query ActionFunctionGetParams, opts ...option.RequestOption) (res *PublicActionFunction, err error) {
	opts = slices.Concat(r.options, opts)
	if query.DefinitionID == "" {
		err = errors.New("missing required definitionId parameter")
		return nil, err
	}
	if functionID == "" {
		err = errors.New("missing required functionId parameter")
		return nil, err
	}
	path := fmt.Sprintf("automation/actions/2026-03/%v/%s/functions/%v/%s", query.AppID, url.PathEscape(query.DefinitionID), query.FunctionType, url.PathEscape(functionID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Retrieve functions of a specific type for a given definition.
func (r *ActionFunctionService) GetByFunctionType(ctx context.Context, functionType ActionFunctionGetByFunctionTypeParamsFunctionType, query ActionFunctionGetByFunctionTypeParams, opts ...option.RequestOption) (res *PublicActionFunction, err error) {
	opts = slices.Concat(r.options, opts)
	if query.DefinitionID == "" {
		err = errors.New("missing required definitionId parameter")
		return nil, err
	}
	path := fmt.Sprintf("automation/actions/2026-03/%v/%s/functions/%v", query.AppID, url.PathEscape(query.DefinitionID), functionType)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type ActionFunctionListParams struct {
	AppID int64 `path:"appId" api:"required" json:"-"`
	paramObj
}

type ActionFunctionDeleteParams struct {
	AppID        int64  `path:"appId" api:"required" json:"-"`
	DefinitionID string `path:"definitionId" api:"required" json:"-"`
	// Any of "POST_ACTION_EXECUTION", "POST_FETCH_OPTIONS", "PRE_ACTION_EXECUTION",
	// "PRE_FETCH_OPTIONS".
	FunctionType ActionFunctionDeleteParamsFunctionType `path:"functionType,omitzero" api:"required" json:"-"`
	paramObj
}

type ActionFunctionDeleteParamsFunctionType string

const (
	ActionFunctionDeleteParamsFunctionTypePostActionExecution ActionFunctionDeleteParamsFunctionType = "POST_ACTION_EXECUTION"
	ActionFunctionDeleteParamsFunctionTypePostFetchOptions    ActionFunctionDeleteParamsFunctionType = "POST_FETCH_OPTIONS"
	ActionFunctionDeleteParamsFunctionTypePreActionExecution  ActionFunctionDeleteParamsFunctionType = "PRE_ACTION_EXECUTION"
	ActionFunctionDeleteParamsFunctionTypePreFetchOptions     ActionFunctionDeleteParamsFunctionType = "PRE_FETCH_OPTIONS"
)

type ActionFunctionNewOrReplaceParams struct {
	AppID        int64  `path:"appId" api:"required" json:"-"`
	DefinitionID string `path:"definitionId" api:"required" json:"-"`
	// Any of "POST_ACTION_EXECUTION", "POST_FETCH_OPTIONS", "PRE_ACTION_EXECUTION",
	// "PRE_FETCH_OPTIONS".
	FunctionType ActionFunctionNewOrReplaceParamsFunctionType `path:"functionType,omitzero" api:"required" json:"-"`
	Body         string
	paramObj
}

func (r ActionFunctionNewOrReplaceParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.Body)
}
func (r *ActionFunctionNewOrReplaceParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ActionFunctionNewOrReplaceParamsFunctionType string

const (
	ActionFunctionNewOrReplaceParamsFunctionTypePostActionExecution ActionFunctionNewOrReplaceParamsFunctionType = "POST_ACTION_EXECUTION"
	ActionFunctionNewOrReplaceParamsFunctionTypePostFetchOptions    ActionFunctionNewOrReplaceParamsFunctionType = "POST_FETCH_OPTIONS"
	ActionFunctionNewOrReplaceParamsFunctionTypePreActionExecution  ActionFunctionNewOrReplaceParamsFunctionType = "PRE_ACTION_EXECUTION"
	ActionFunctionNewOrReplaceParamsFunctionTypePreFetchOptions     ActionFunctionNewOrReplaceParamsFunctionType = "PRE_FETCH_OPTIONS"
)

type ActionFunctionNewOrReplaceByFunctionTypeParams struct {
	AppID        int64  `path:"appId" api:"required" json:"-"`
	DefinitionID string `path:"definitionId" api:"required" json:"-"`
	Body         string
	paramObj
}

func (r ActionFunctionNewOrReplaceByFunctionTypeParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.Body)
}
func (r *ActionFunctionNewOrReplaceByFunctionTypeParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ActionFunctionNewOrReplaceByFunctionTypeParamsFunctionType string

const (
	ActionFunctionNewOrReplaceByFunctionTypeParamsFunctionTypePostActionExecution ActionFunctionNewOrReplaceByFunctionTypeParamsFunctionType = "POST_ACTION_EXECUTION"
	ActionFunctionNewOrReplaceByFunctionTypeParamsFunctionTypePostFetchOptions    ActionFunctionNewOrReplaceByFunctionTypeParamsFunctionType = "POST_FETCH_OPTIONS"
	ActionFunctionNewOrReplaceByFunctionTypeParamsFunctionTypePreActionExecution  ActionFunctionNewOrReplaceByFunctionTypeParamsFunctionType = "PRE_ACTION_EXECUTION"
	ActionFunctionNewOrReplaceByFunctionTypeParamsFunctionTypePreFetchOptions     ActionFunctionNewOrReplaceByFunctionTypeParamsFunctionType = "PRE_FETCH_OPTIONS"
)

type ActionFunctionDeleteByFunctionTypeParams struct {
	AppID        int64  `path:"appId" api:"required" json:"-"`
	DefinitionID string `path:"definitionId" api:"required" json:"-"`
	paramObj
}

type ActionFunctionDeleteByFunctionTypeParamsFunctionType string

const (
	ActionFunctionDeleteByFunctionTypeParamsFunctionTypePostActionExecution ActionFunctionDeleteByFunctionTypeParamsFunctionType = "POST_ACTION_EXECUTION"
	ActionFunctionDeleteByFunctionTypeParamsFunctionTypePostFetchOptions    ActionFunctionDeleteByFunctionTypeParamsFunctionType = "POST_FETCH_OPTIONS"
	ActionFunctionDeleteByFunctionTypeParamsFunctionTypePreActionExecution  ActionFunctionDeleteByFunctionTypeParamsFunctionType = "PRE_ACTION_EXECUTION"
	ActionFunctionDeleteByFunctionTypeParamsFunctionTypePreFetchOptions     ActionFunctionDeleteByFunctionTypeParamsFunctionType = "PRE_FETCH_OPTIONS"
)

type ActionFunctionGetParams struct {
	AppID        int64  `path:"appId" api:"required" json:"-"`
	DefinitionID string `path:"definitionId" api:"required" json:"-"`
	// Any of "POST_ACTION_EXECUTION", "POST_FETCH_OPTIONS", "PRE_ACTION_EXECUTION",
	// "PRE_FETCH_OPTIONS".
	FunctionType ActionFunctionGetParamsFunctionType `path:"functionType,omitzero" api:"required" json:"-"`
	paramObj
}

type ActionFunctionGetParamsFunctionType string

const (
	ActionFunctionGetParamsFunctionTypePostActionExecution ActionFunctionGetParamsFunctionType = "POST_ACTION_EXECUTION"
	ActionFunctionGetParamsFunctionTypePostFetchOptions    ActionFunctionGetParamsFunctionType = "POST_FETCH_OPTIONS"
	ActionFunctionGetParamsFunctionTypePreActionExecution  ActionFunctionGetParamsFunctionType = "PRE_ACTION_EXECUTION"
	ActionFunctionGetParamsFunctionTypePreFetchOptions     ActionFunctionGetParamsFunctionType = "PRE_FETCH_OPTIONS"
)

type ActionFunctionGetByFunctionTypeParams struct {
	AppID        int64  `path:"appId" api:"required" json:"-"`
	DefinitionID string `path:"definitionId" api:"required" json:"-"`
	paramObj
}

type ActionFunctionGetByFunctionTypeParamsFunctionType string

const (
	ActionFunctionGetByFunctionTypeParamsFunctionTypePostActionExecution ActionFunctionGetByFunctionTypeParamsFunctionType = "POST_ACTION_EXECUTION"
	ActionFunctionGetByFunctionTypeParamsFunctionTypePostFetchOptions    ActionFunctionGetByFunctionTypeParamsFunctionType = "POST_FETCH_OPTIONS"
	ActionFunctionGetByFunctionTypeParamsFunctionTypePreActionExecution  ActionFunctionGetByFunctionTypeParamsFunctionType = "PRE_ACTION_EXECUTION"
	ActionFunctionGetByFunctionTypeParamsFunctionTypePreFetchOptions     ActionFunctionGetByFunctionTypeParamsFunctionType = "PRE_FETCH_OPTIONS"
)
