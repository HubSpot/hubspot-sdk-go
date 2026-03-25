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
	"time"

	"github.com/stainless-sdks/hubspot-sdk-go/events"
	"github.com/stainless-sdks/hubspot-sdk-go/internal/apijson"
	"github.com/stainless-sdks/hubspot-sdk-go/internal/apiquery"
	shimjson "github.com/stainless-sdks/hubspot-sdk-go/internal/encoding/json"
	"github.com/stainless-sdks/hubspot-sdk-go/internal/requestconfig"
	"github.com/stainless-sdks/hubspot-sdk-go/option"
	"github.com/stainless-sdks/hubspot-sdk-go/packages/pagination"
	"github.com/stainless-sdks/hubspot-sdk-go/packages/param"
	"github.com/stainless-sdks/hubspot-sdk-go/packages/respjson"
	"github.com/stainless-sdks/hubspot-sdk-go/shared"
)

// ActionService contains methods and other services that help with interacting
// with the hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewActionService] method instead.
type ActionService struct {
	Options []option.RequestOption
}

// NewActionService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewActionService(opts ...option.RequestOption) (r ActionService) {
	r = ActionService{}
	r.Options = opts
	return
}

func (r *ActionService) New(ctx context.Context, appID int64, body ActionNewParams, opts ...option.RequestOption) (res *PublicActionDefinition, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("automation/actions/2026-03/%v", appID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

func (r *ActionService) Update(ctx context.Context, definitionID string, params ActionUpdateParams, opts ...option.RequestOption) (res *PublicActionDefinition, err error) {
	opts = slices.Concat(r.Options, opts)
	if definitionID == "" {
		err = errors.New("missing required definitionId parameter")
		return nil, err
	}
	path := fmt.Sprintf("automation/actions/2026-03/%v/%s", params.AppID, definitionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return res, err
}

func (r *ActionService) List(ctx context.Context, definitionID string, params ActionListParams, opts ...option.RequestOption) (res *pagination.Page[PublicActionRevision], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if definitionID == "" {
		err = errors.New("missing required definitionId parameter")
		return nil, err
	}
	path := fmt.Sprintf("automation/actions/2026-03/%v/%s/revisions", params.AppID, definitionID)
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

func (r *ActionService) ListAutoPaging(ctx context.Context, definitionID string, params ActionListParams, opts ...option.RequestOption) *pagination.PageAutoPager[PublicActionRevision] {
	return pagination.NewPageAutoPager(r.List(ctx, definitionID, params, opts...))
}

func (r *ActionService) Delete(ctx context.Context, functionID string, body ActionDeleteParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if body.DefinitionID == "" {
		err = errors.New("missing required definitionId parameter")
		return err
	}
	if functionID == "" {
		err = errors.New("missing required functionId parameter")
		return err
	}
	path := fmt.Sprintf("automation/actions/2026-03/%v/%s/functions/%v/%s", body.AppID, body.DefinitionID, body.FunctionType, functionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

func (r *ActionService) Complete(ctx context.Context, callbackID string, body ActionCompleteParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if callbackID == "" {
		err = errors.New("missing required callbackId parameter")
		return err
	}
	path := fmt.Sprintf("automation/actions/callbacks/2026-03/%s/complete", callbackID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return err
}

func (r *ActionService) CompleteBatch(ctx context.Context, body ActionCompleteBatchParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "automation/actions/callbacks/2026-03/complete"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return err
}

func (r *ActionService) NewOrReplace(ctx context.Context, functionID string, params ActionNewOrReplaceParams, opts ...option.RequestOption) (res *PublicActionFunctionIdentifier, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.DefinitionID == "" {
		err = errors.New("missing required definitionId parameter")
		return nil, err
	}
	if functionID == "" {
		err = errors.New("missing required functionId parameter")
		return nil, err
	}
	path := fmt.Sprintf("automation/actions/2026-03/%v/%s/functions/%v/%s", params.AppID, params.DefinitionID, params.FunctionType, functionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return res, err
}

func (r *ActionService) NewOrReplaceByFunctionType(ctx context.Context, functionType ActionNewOrReplaceByFunctionTypeParamsFunctionType, params ActionNewOrReplaceByFunctionTypeParams, opts ...option.RequestOption) (res *PublicActionFunctionIdentifier, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.DefinitionID == "" {
		err = errors.New("missing required definitionId parameter")
		return nil, err
	}
	path := fmt.Sprintf("automation/actions/2026-03/%v/%s/functions/%v", params.AppID, params.DefinitionID, functionType)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return res, err
}

func (r *ActionService) NewRequiresObject(ctx context.Context, definitionID string, params ActionNewRequiresObjectParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if definitionID == "" {
		err = errors.New("missing required definitionId parameter")
		return err
	}
	path := fmt.Sprintf("automation/actions/2026-03/%v/%s/requires-object", params.AppID, definitionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, nil, opts...)
	return err
}

func (r *ActionService) DeleteByFunctionType(ctx context.Context, functionType ActionDeleteByFunctionTypeParamsFunctionType, body ActionDeleteByFunctionTypeParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if body.DefinitionID == "" {
		err = errors.New("missing required definitionId parameter")
		return err
	}
	path := fmt.Sprintf("automation/actions/2026-03/%v/%s/functions/%v", body.AppID, body.DefinitionID, functionType)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

func (r *ActionService) Get(ctx context.Context, revisionID string, query ActionGetParams, opts ...option.RequestOption) (res *PublicActionRevision, err error) {
	opts = slices.Concat(r.Options, opts)
	if query.DefinitionID == "" {
		err = errors.New("missing required definitionId parameter")
		return nil, err
	}
	if revisionID == "" {
		err = errors.New("missing required revisionId parameter")
		return nil, err
	}
	path := fmt.Sprintf("automation/actions/2026-03/%v/%s/revisions/%s", query.AppID, query.DefinitionID, revisionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

func (r *ActionService) GetByFunctionType(ctx context.Context, functionType ActionGetByFunctionTypeParamsFunctionType, query ActionGetByFunctionTypeParams, opts ...option.RequestOption) (res *PublicActionFunction, err error) {
	opts = slices.Concat(r.Options, opts)
	if query.DefinitionID == "" {
		err = errors.New("missing required definitionId parameter")
		return nil, err
	}
	path := fmt.Sprintf("automation/actions/2026-03/%v/%s/functions/%v", query.AppID, query.DefinitionID, functionType)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

func (r *ActionService) GetRequiresObject(ctx context.Context, definitionID string, query ActionGetRequiresObjectParams, opts ...option.RequestOption) (res *PublicActionDefinitionRequiresObjectResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if definitionID == "" {
		err = errors.New("missing required definitionId parameter")
		return nil, err
	}
	path := fmt.Sprintf("automation/actions/2026-03/%v/%s/requires-object", query.AppID, definitionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// The properties ActionExecutionIndex, EnrollmentID are required.
type ActionExecutionIndexIdentifierParam struct {
	ActionExecutionIndex int64 `json:"actionExecutionIndex" api:"required"`
	EnrollmentID         int64 `json:"enrollmentId" api:"required"`
	paramObj
}

func (r ActionExecutionIndexIdentifierParam) MarshalJSON() (data []byte, err error) {
	type shadow ActionExecutionIndexIdentifierParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ActionExecutionIndexIdentifierParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties AgentID, ChirpAIContextObject, Source are required.
type AgentRequestContextParam struct {
	AgentID              int64                     `json:"agentId" api:"required"`
	ChirpAIContextObject ChirpAIContextObjectParam `json:"chirpAiContextObject,omitzero" api:"required"`
	// Any of "AGENTS".
	Source       AgentRequestContextSource `json:"source,omitzero" api:"required"`
	TrajectoryID param.Opt[string]         `json:"trajectoryId,omitzero"`
	paramObj
}

func (r AgentRequestContextParam) MarshalJSON() (data []byte, err error) {
	type shadow AgentRequestContextParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AgentRequestContextParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AgentRequestContextSource string

const (
	AgentRequestContextSourceAgents AgentRequestContextSource = "AGENTS"
)

type ArrayFieldSchema struct {
	Items ArrayFieldSchemaItemsUnion `json:"items" api:"required"`
	// Any of "ARRAY".
	Type ArrayFieldSchemaType `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Items       respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ArrayFieldSchema) RawJSON() string { return r.JSON.raw }
func (r *ArrayFieldSchema) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this ArrayFieldSchema to a ArrayFieldSchemaParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// ArrayFieldSchemaParam.Overrides()
func (r ArrayFieldSchema) ToParam() ArrayFieldSchemaParam {
	return param.Override[ArrayFieldSchemaParam](json.RawMessage(r.RawJSON()))
}

// ArrayFieldSchemaItemsUnion contains all possible properties and values from
// [IntegerFieldSchema], [LongFieldSchema], [DoubleFieldSchema],
// [StringFieldSchema], [BooleanFieldSchema], [ArrayFieldSchema],
// [ObjectFieldSchema].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ArrayFieldSchemaItemsUnion struct {
	Type string `json:"type"`
	// This field is a union of [int64], [int64], [float64]
	Maximum ArrayFieldSchemaItemsUnionMaximum `json:"maximum"`
	// This field is a union of [int64], [int64], [float64]
	Minimum ArrayFieldSchemaItemsUnionMinimum `json:"minimum"`
	// This field is from variant [StringFieldSchema].
	Format StringFieldSchemaFormat `json:"format"`
	// This field is from variant [ArrayFieldSchema].
	Items ArrayFieldSchemaItemsUnion `json:"items"`
	// This field is from variant [ObjectFieldSchema].
	Properties any `json:"properties"`
	JSON       struct {
		Type       respjson.Field
		Maximum    respjson.Field
		Minimum    respjson.Field
		Format     respjson.Field
		Items      respjson.Field
		Properties respjson.Field
		raw        string
	} `json:"-"`
}

func (u ArrayFieldSchemaItemsUnion) AsInteger() (v IntegerFieldSchema) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ArrayFieldSchemaItemsUnion) AsLong() (v LongFieldSchema) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ArrayFieldSchemaItemsUnion) AsDouble() (v DoubleFieldSchema) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ArrayFieldSchemaItemsUnion) AsString() (v StringFieldSchema) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ArrayFieldSchemaItemsUnion) AsBoolean() (v BooleanFieldSchema) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ArrayFieldSchemaItemsUnion) AsArray() (v ArrayFieldSchema) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ArrayFieldSchemaItemsUnion) AsObject() (v ObjectFieldSchema) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ArrayFieldSchemaItemsUnion) RawJSON() string { return u.JSON.raw }

func (r *ArrayFieldSchemaItemsUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ArrayFieldSchemaItemsUnionMaximum is an implicit subunion of
// [ArrayFieldSchemaItemsUnion]. ArrayFieldSchemaItemsUnionMaximum provides
// convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [ArrayFieldSchemaItemsUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfInt OfFloat]
type ArrayFieldSchemaItemsUnionMaximum struct {
	// This field will be present if the value is a [int64] instead of an object.
	OfInt int64 `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	JSON    struct {
		OfInt   respjson.Field
		OfFloat respjson.Field
		raw     string
	} `json:"-"`
}

func (r *ArrayFieldSchemaItemsUnionMaximum) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ArrayFieldSchemaItemsUnionMinimum is an implicit subunion of
// [ArrayFieldSchemaItemsUnion]. ArrayFieldSchemaItemsUnionMinimum provides
// convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [ArrayFieldSchemaItemsUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfInt OfFloat]
type ArrayFieldSchemaItemsUnionMinimum struct {
	// This field will be present if the value is a [int64] instead of an object.
	OfInt int64 `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	JSON    struct {
		OfInt   respjson.Field
		OfFloat respjson.Field
		raw     string
	} `json:"-"`
}

func (r *ArrayFieldSchemaItemsUnionMinimum) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ArrayFieldSchemaType string

const (
	ArrayFieldSchemaTypeArray ArrayFieldSchemaType = "ARRAY"
)

// The properties Items, Type are required.
type ArrayFieldSchemaParam struct {
	Items ArrayFieldSchemaItemsUnionParam `json:"items,omitzero" api:"required"`
	// Any of "ARRAY".
	Type ArrayFieldSchemaType `json:"type,omitzero" api:"required"`
	paramObj
}

func (r ArrayFieldSchemaParam) MarshalJSON() (data []byte, err error) {
	type shadow ArrayFieldSchemaParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ArrayFieldSchemaParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ArrayFieldSchemaItemsUnionParam struct {
	OfInteger *IntegerFieldSchemaParam `json:",omitzero,inline"`
	OfLong    *LongFieldSchemaParam    `json:",omitzero,inline"`
	OfDouble  *DoubleFieldSchemaParam  `json:",omitzero,inline"`
	OfString  *StringFieldSchemaParam  `json:",omitzero,inline"`
	OfBoolean *BooleanFieldSchemaParam `json:",omitzero,inline"`
	OfArray   *ArrayFieldSchemaParam   `json:",omitzero,inline"`
	OfObject  *ObjectFieldSchemaParam  `json:",omitzero,inline"`
	paramUnion
}

func (u ArrayFieldSchemaItemsUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfInteger,
		u.OfLong,
		u.OfDouble,
		u.OfString,
		u.OfBoolean,
		u.OfArray,
		u.OfObject)
}
func (u *ArrayFieldSchemaItemsUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ArrayFieldSchemaItemsUnionParam) asAny() any {
	if !param.IsOmitted(u.OfInteger) {
		return u.OfInteger
	} else if !param.IsOmitted(u.OfLong) {
		return u.OfLong
	} else if !param.IsOmitted(u.OfDouble) {
		return u.OfDouble
	} else if !param.IsOmitted(u.OfString) {
		return u.OfString
	} else if !param.IsOmitted(u.OfBoolean) {
		return u.OfBoolean
	} else if !param.IsOmitted(u.OfArray) {
		return u.OfArray
	} else if !param.IsOmitted(u.OfObject) {
		return u.OfObject
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ArrayFieldSchemaItemsUnionParam) GetFormat() *string {
	if vt := u.OfString; vt != nil {
		return (*string)(&vt.Format)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ArrayFieldSchemaItemsUnionParam) GetItems() *ArrayFieldSchemaItemsUnionParam {
	if vt := u.OfArray; vt != nil {
		return &vt.Items
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ArrayFieldSchemaItemsUnionParam) GetProperties() *any {
	if vt := u.OfObject; vt != nil {
		return &vt.Properties
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ArrayFieldSchemaItemsUnionParam) GetType() *string {
	if vt := u.OfInteger; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfLong; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfDouble; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfString; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfBoolean; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfArray; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfObject; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

// Returns a subunion which exports methods to access subproperties
//
// Or use AsAny() to get the underlying value
func (u ArrayFieldSchemaItemsUnionParam) GetMaximum() (res arrayFieldSchemaItemsUnionParamMaximum) {
	if vt := u.OfInteger; vt != nil && vt.Maximum.Valid() {
		res.any = &vt.Maximum.Value
	} else if vt := u.OfLong; vt != nil && vt.Maximum.Valid() {
		res.any = &vt.Maximum.Value
	} else if vt := u.OfDouble; vt != nil && vt.Maximum.Valid() {
		res.any = &vt.Maximum.Value
	}
	return
}

// Can have the runtime types [*int64], [*float64]
type arrayFieldSchemaItemsUnionParamMaximum struct{ any }

// Use the following switch statement to get the type of the union:
//
//	switch u.AsAny().(type) {
//	case *int64:
//	case *float64:
//	default:
//	    fmt.Errorf("not present")
//	}
func (u arrayFieldSchemaItemsUnionParamMaximum) AsAny() any { return u.any }

// Returns a subunion which exports methods to access subproperties
//
// Or use AsAny() to get the underlying value
func (u ArrayFieldSchemaItemsUnionParam) GetMinimum() (res arrayFieldSchemaItemsUnionParamMinimum) {
	if vt := u.OfInteger; vt != nil && vt.Minimum.Valid() {
		res.any = &vt.Minimum.Value
	} else if vt := u.OfLong; vt != nil && vt.Minimum.Valid() {
		res.any = &vt.Minimum.Value
	} else if vt := u.OfDouble; vt != nil && vt.Minimum.Valid() {
		res.any = &vt.Minimum.Value
	}
	return
}

// Can have the runtime types [*int64], [*float64]
type arrayFieldSchemaItemsUnionParamMinimum struct{ any }

// Use the following switch statement to get the type of the union:
//
//	switch u.AsAny().(type) {
//	case *int64:
//	case *float64:
//	default:
//	    fmt.Errorf("not present")
//	}
func (u arrayFieldSchemaItemsUnionParamMinimum) AsAny() any { return u.any }

// The property Inputs is required.
type BatchInputCallbackCompletionBatchRequestParam struct {
	Inputs []CallbackCompletionBatchRequestParam `json:"inputs,omitzero" api:"required"`
	paramObj
}

func (r BatchInputCallbackCompletionBatchRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow BatchInputCallbackCompletionBatchRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BatchInputCallbackCompletionBatchRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BooleanFieldSchema struct {
	// Any of "BOOLEAN".
	Type BooleanFieldSchemaType `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BooleanFieldSchema) RawJSON() string { return r.JSON.raw }
func (r *BooleanFieldSchema) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this BooleanFieldSchema to a BooleanFieldSchemaParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// BooleanFieldSchemaParam.Overrides()
func (r BooleanFieldSchema) ToParam() BooleanFieldSchemaParam {
	return param.Override[BooleanFieldSchemaParam](json.RawMessage(r.RawJSON()))
}

type BooleanFieldSchemaType string

const (
	BooleanFieldSchemaTypeBoolean BooleanFieldSchemaType = "BOOLEAN"
)

// The property Type is required.
type BooleanFieldSchemaParam struct {
	// Any of "BOOLEAN".
	Type BooleanFieldSchemaType `json:"type,omitzero" api:"required"`
	paramObj
}

func (r BooleanFieldSchemaParam) MarshalJSON() (data []byte, err error) {
	type shadow BooleanFieldSchemaParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BooleanFieldSchemaParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties CallbackID, OutputFields, TypedOutputs are required.
type CallbackCompletionBatchRequestParam struct {
	CallbackID        string                                                 `json:"callbackId" api:"required"`
	OutputFields      map[string]string                                      `json:"outputFields,omitzero" api:"required"`
	TypedOutputs      any                                                    `json:"typedOutputs,omitzero" api:"required"`
	FailureReasonType param.Opt[string]                                      `json:"failureReasonType,omitzero"`
	RequestContext    CallbackCompletionBatchRequestRequestContextUnionParam `json:"requestContext,omitzero"`
	paramObj
}

func (r CallbackCompletionBatchRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow CallbackCompletionBatchRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CallbackCompletionBatchRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type CallbackCompletionBatchRequestRequestContextUnionParam struct {
	OfWorkflows  *WorkflowsRequestContextParam  `json:",omitzero,inline"`
	OfAgents     *AgentRequestContextParam      `json:",omitzero,inline"`
	OfCopilot    *CopilotRequestContextParam    `json:",omitzero,inline"`
	OfStandalone *StandaloneRequestContextParam `json:",omitzero,inline"`
	OfTest       *TestRequestContextParam       `json:",omitzero,inline"`
	paramUnion
}

func (u CallbackCompletionBatchRequestRequestContextUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfWorkflows,
		u.OfAgents,
		u.OfCopilot,
		u.OfStandalone,
		u.OfTest)
}
func (u *CallbackCompletionBatchRequestRequestContextUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *CallbackCompletionBatchRequestRequestContextUnionParam) asAny() any {
	if !param.IsOmitted(u.OfWorkflows) {
		return u.OfWorkflows
	} else if !param.IsOmitted(u.OfAgents) {
		return u.OfAgents
	} else if !param.IsOmitted(u.OfCopilot) {
		return u.OfCopilot
	} else if !param.IsOmitted(u.OfStandalone) {
		return u.OfStandalone
	} else if !param.IsOmitted(u.OfTest) {
		return u.OfTest
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u CallbackCompletionBatchRequestRequestContextUnionParam) GetWorkflowID() *int64 {
	if vt := u.OfWorkflows; vt != nil {
		return &vt.WorkflowID
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u CallbackCompletionBatchRequestRequestContextUnionParam) GetActionExecutionIndexIdentifier() *ActionExecutionIndexIdentifierParam {
	if vt := u.OfWorkflows; vt != nil {
		return &vt.ActionExecutionIndexIdentifier
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u CallbackCompletionBatchRequestRequestContextUnionParam) GetActionID() *int64 {
	if vt := u.OfWorkflows; vt != nil && vt.ActionID.Valid() {
		return &vt.ActionID.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u CallbackCompletionBatchRequestRequestContextUnionParam) GetAgentID() *int64 {
	if vt := u.OfAgents; vt != nil {
		return &vt.AgentID
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u CallbackCompletionBatchRequestRequestContextUnionParam) GetSource() *string {
	if vt := u.OfWorkflows; vt != nil {
		return (*string)(&vt.Source)
	} else if vt := u.OfAgents; vt != nil {
		return (*string)(&vt.Source)
	} else if vt := u.OfCopilot; vt != nil {
		return (*string)(&vt.Source)
	} else if vt := u.OfStandalone; vt != nil {
		return (*string)(&vt.Source)
	} else if vt := u.OfTest; vt != nil {
		return (*string)(&vt.Source)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u CallbackCompletionBatchRequestRequestContextUnionParam) GetTrajectoryID() *string {
	if vt := u.OfAgents; vt != nil && vt.TrajectoryID.Valid() {
		return &vt.TrajectoryID.Value
	} else if vt := u.OfCopilot; vt != nil && vt.TrajectoryID.Valid() {
		return &vt.TrajectoryID.Value
	} else if vt := u.OfStandalone; vt != nil && vt.TrajectoryID.Valid() {
		return &vt.TrajectoryID.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's ChirpAIContextObject property, if
// present.
func (u CallbackCompletionBatchRequestRequestContextUnionParam) GetChirpAIContextObject() *ChirpAIContextObjectParam {
	if vt := u.OfAgents; vt != nil {
		return &vt.ChirpAIContextObject
	} else if vt := u.OfStandalone; vt != nil {
		return &vt.ChirpAIContextObject
	}
	return nil
}

// The properties OutputFields, TypedOutputs are required.
type CallbackCompletionRequestParam struct {
	OutputFields      map[string]string                                 `json:"outputFields,omitzero" api:"required"`
	TypedOutputs      any                                               `json:"typedOutputs,omitzero" api:"required"`
	FailureReasonType param.Opt[string]                                 `json:"failureReasonType,omitzero"`
	RequestContext    CallbackCompletionRequestRequestContextUnionParam `json:"requestContext,omitzero"`
	paramObj
}

func (r CallbackCompletionRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow CallbackCompletionRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CallbackCompletionRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type CallbackCompletionRequestRequestContextUnionParam struct {
	OfWorkflows  *WorkflowsRequestContextParam  `json:",omitzero,inline"`
	OfAgents     *AgentRequestContextParam      `json:",omitzero,inline"`
	OfCopilot    *CopilotRequestContextParam    `json:",omitzero,inline"`
	OfStandalone *StandaloneRequestContextParam `json:",omitzero,inline"`
	OfTest       *TestRequestContextParam       `json:",omitzero,inline"`
	paramUnion
}

func (u CallbackCompletionRequestRequestContextUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfWorkflows,
		u.OfAgents,
		u.OfCopilot,
		u.OfStandalone,
		u.OfTest)
}
func (u *CallbackCompletionRequestRequestContextUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *CallbackCompletionRequestRequestContextUnionParam) asAny() any {
	if !param.IsOmitted(u.OfWorkflows) {
		return u.OfWorkflows
	} else if !param.IsOmitted(u.OfAgents) {
		return u.OfAgents
	} else if !param.IsOmitted(u.OfCopilot) {
		return u.OfCopilot
	} else if !param.IsOmitted(u.OfStandalone) {
		return u.OfStandalone
	} else if !param.IsOmitted(u.OfTest) {
		return u.OfTest
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u CallbackCompletionRequestRequestContextUnionParam) GetWorkflowID() *int64 {
	if vt := u.OfWorkflows; vt != nil {
		return &vt.WorkflowID
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u CallbackCompletionRequestRequestContextUnionParam) GetActionExecutionIndexIdentifier() *ActionExecutionIndexIdentifierParam {
	if vt := u.OfWorkflows; vt != nil {
		return &vt.ActionExecutionIndexIdentifier
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u CallbackCompletionRequestRequestContextUnionParam) GetActionID() *int64 {
	if vt := u.OfWorkflows; vt != nil && vt.ActionID.Valid() {
		return &vt.ActionID.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u CallbackCompletionRequestRequestContextUnionParam) GetAgentID() *int64 {
	if vt := u.OfAgents; vt != nil {
		return &vt.AgentID
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u CallbackCompletionRequestRequestContextUnionParam) GetSource() *string {
	if vt := u.OfWorkflows; vt != nil {
		return (*string)(&vt.Source)
	} else if vt := u.OfAgents; vt != nil {
		return (*string)(&vt.Source)
	} else if vt := u.OfCopilot; vt != nil {
		return (*string)(&vt.Source)
	} else if vt := u.OfStandalone; vt != nil {
		return (*string)(&vt.Source)
	} else if vt := u.OfTest; vt != nil {
		return (*string)(&vt.Source)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u CallbackCompletionRequestRequestContextUnionParam) GetTrajectoryID() *string {
	if vt := u.OfAgents; vt != nil && vt.TrajectoryID.Valid() {
		return &vt.TrajectoryID.Value
	} else if vt := u.OfCopilot; vt != nil && vt.TrajectoryID.Valid() {
		return &vt.TrajectoryID.Value
	} else if vt := u.OfStandalone; vt != nil && vt.TrajectoryID.Valid() {
		return &vt.TrajectoryID.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's ChirpAIContextObject property, if
// present.
func (u CallbackCompletionRequestRequestContextUnionParam) GetChirpAIContextObject() *ChirpAIContextObjectParam {
	if vt := u.OfAgents; vt != nil {
		return &vt.ChirpAIContextObject
	} else if vt := u.OfStandalone; vt != nil {
		return &vt.ChirpAIContextObject
	}
	return nil
}

// The properties ApplicationGroup, ApplicationID, Metadata, OtelContextHolder,
// UnstructuredSources are required.
type ChirpAIContextObjectParam struct {
	ApplicationGroup  string            `json:"applicationGroup" api:"required"`
	ApplicationID     string            `json:"applicationId" api:"required"`
	Metadata          map[string]string `json:"metadata,omitzero" api:"required"`
	OtelContextHolder map[string]string `json:"otelContextHolder,omitzero" api:"required"`
	// Any of "NONE", "USER_INPUT", "LOGGED_EMAIL", "VIDEO_CALL", "AUDIO_CALL",
	// "CALL_TRANSCRIPT", "MEETING_TRANSCRIPT", "FORMS", "FEEDBACK_SURVEY", "PDF",
	// "QUOTE", "INVOICE", "OTHER_ATTACHMENT_DOC", "WHATSAPP", "SMS", "CHAT",
	// "FACEBOOK_MESSENGER", "CUSTOM_CHANNEL_OR_API", "MANY", "NOTE", "DERIVED".
	UnstructuredSources []string           `json:"unstructuredSources,omitzero" api:"required"`
	FeatureID           param.Opt[string]  `json:"featureId,omitzero"`
	InferenceID         param.Opt[string]  `json:"inferenceId,omitzero"`
	TrajectoryID        param.Opt[string]  `json:"trajectoryId,omitzero" format:"uuid"`
	ComplianceIDs       ComplianceIDsParam `json:"complianceIds,omitzero"`
	paramObj
}

func (r ChirpAIContextObjectParam) MarshalJSON() (data []byte, err error) {
	type shadow ChirpAIContextObjectParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ChirpAIContextObjectParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CollectionResponsePublicActionRevisionForwardPaging struct {
	Results []PublicActionRevision `json:"results" api:"required"`
	Paging  shared.ForwardPaging   `json:"paging"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Results     respjson.Field
		Paging      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CollectionResponsePublicActionRevisionForwardPaging) RawJSON() string { return r.JSON.raw }
func (r *CollectionResponsePublicActionRevisionForwardPaging) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties ContactIDs, PortalIDs, UserIDs are required.
type ComplianceIDsParam struct {
	ContactIDs        []ContactIDParam  `json:"contactIds,omitzero" api:"required"`
	PortalIDs         []int64           `json:"portalIds,omitzero" api:"required"`
	UserIDs           []int64           `json:"userIds,omitzero" api:"required"`
	NoContactIDReason param.Opt[string] `json:"noContactIdReason,omitzero"`
	NoPortalIDReason  param.Opt[string] `json:"noPortalIdReason,omitzero"`
	NoUserIDReason    param.Opt[string] `json:"noUserIdReason,omitzero"`
	paramObj
}

func (r ComplianceIDsParam) MarshalJSON() (data []byte, err error) {
	type shadow ComplianceIDsParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ComplianceIDsParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property PortalID is required.
type ContactIDParam struct {
	PortalID int64             `json:"portalId" api:"required"`
	Email    param.Opt[string] `json:"email,omitzero"`
	Vid      param.Opt[int64]  `json:"vid,omitzero"`
	paramObj
}

func (r ContactIDParam) MarshalJSON() (data []byte, err error) {
	type shadow ContactIDParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ContactIDParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Source is required.
type CopilotRequestContextParam struct {
	// Any of "COPILOT".
	Source       CopilotRequestContextSource `json:"source,omitzero" api:"required"`
	TrajectoryID param.Opt[string]           `json:"trajectoryId,omitzero"`
	paramObj
}

func (r CopilotRequestContextParam) MarshalJSON() (data []byte, err error) {
	type shadow CopilotRequestContextParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CopilotRequestContextParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CopilotRequestContextSource string

const (
	CopilotRequestContextSourceCopilot CopilotRequestContextSource = "COPILOT"
)

type DoubleFieldSchema struct {
	// Any of "DOUBLE".
	Type    DoubleFieldSchemaType `json:"type" api:"required"`
	Maximum float64               `json:"maximum"`
	Minimum float64               `json:"minimum"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		Maximum     respjson.Field
		Minimum     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DoubleFieldSchema) RawJSON() string { return r.JSON.raw }
func (r *DoubleFieldSchema) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this DoubleFieldSchema to a DoubleFieldSchemaParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// DoubleFieldSchemaParam.Overrides()
func (r DoubleFieldSchema) ToParam() DoubleFieldSchemaParam {
	return param.Override[DoubleFieldSchemaParam](json.RawMessage(r.RawJSON()))
}

type DoubleFieldSchemaType string

const (
	DoubleFieldSchemaTypeDouble DoubleFieldSchemaType = "DOUBLE"
)

// The property Type is required.
type DoubleFieldSchemaParam struct {
	// Any of "DOUBLE".
	Type    DoubleFieldSchemaType `json:"type,omitzero" api:"required"`
	Maximum param.Opt[float64]    `json:"maximum,omitzero"`
	Minimum param.Opt[float64]    `json:"minimum,omitzero"`
	paramObj
}

func (r DoubleFieldSchemaParam) MarshalJSON() (data []byte, err error) {
	type shadow DoubleFieldSchemaParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DoubleFieldSchemaParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FieldTypeDefinition struct {
	ExternalOptions bool                           `json:"externalOptions" api:"required"`
	Name            string                         `json:"name" api:"required"`
	Options         []events.Option                `json:"options" api:"required"`
	Schema          FieldTypeDefinitionSchemaUnion `json:"schema" api:"required"`
	// Any of "bool", "currency_number", "date", "datetime", "enumeration", "json",
	// "number", "object_coordinates", "phone_number", "string".
	Type                         FieldTypeDefinitionType `json:"type" api:"required"`
	UseChirp                     bool                    `json:"useChirp" api:"required"`
	Description                  string                  `json:"description"`
	ExternalOptionsReferenceType string                  `json:"externalOptionsReferenceType"`
	// Any of "booleancheckbox", "calculation_equation", "calculation_read_time",
	// "calculation_rollup", "calculation_score", "checkbox", "date", "file", "html",
	// "number", "phonenumber", "radio", "select", "text", "textarea", "unknown".
	FieldType  FieldTypeDefinitionFieldType `json:"fieldType"`
	HelpText   string                       `json:"helpText"`
	Label      string                       `json:"label"`
	OptionsURL string                       `json:"optionsUrl"`
	// Any of "ABANDONED_CART", "ACCEPTANCE_TEST", "AD", "AD_ACCOUNT", "AD_CAMPAIGN",
	// "AD_GROUP", "AI_FORECAST", "ALL_PAGES", "APPROVAL", "APPROVAL_STEP",
	// "ATTRIBUTION", "AUDIENCE", "AUTOMATION_JOURNEY", "AUTOMATION_PLATFORM_FLOW",
	// "AUTOMATION_PLATFORM_FLOW_ACTION", "BET_ALERT", "BET_DELIVERABLE_SERVICE",
	// "BLOG_LISTING_PAGE", "BLOG_POST", "CALL", "CAMPAIGN", "CAMPAIGN_BUDGET_ITEM",
	// "CAMPAIGN_SPEND_ITEM", "CAMPAIGN_STEP", "CAMPAIGN_TEMPLATE",
	// "CAMPAIGN_TEMPLATE_STEP", "CART", "CASE_STUDY", "CHATFLOW", "CLIP", "CMS_URL",
	// "COMBO_EVENT_CONFIGURATION", "COMMERCE_PAYMENT", "COMMUNICATION", "COMPANY",
	// "CONTACT", "CONTACT_CREATE_ATTRIBUTION", "CONTENT", "CONTENT_AUDIT",
	// "CONTENT_AUDIT_PAGE", "CONVERSATION", "CONVERSATION_INBOX",
	// "CONVERSATION_SESSION", "CRM_OBJECTS_DUMMY_TYPE", "CRM_PIPELINES_DUMMY_TYPE",
	// "CTA", "CTA_VARIANT", "DATA_PRIVACY_CONSENT", "DATA_SYNC_STATE", "DEAL",
	// "DEAL_CREATE_ATTRIBUTION", "DEAL_REGISTRATION", "DEAL_SPLIT", "DISCOUNT",
	// "DISCOUNT_CODE", "DISCOUNT_TEMPLATE", "EMAIL", "ENGAGEMENT", "EXPORT",
	// "EXTERNAL_WEB_URL", "FEE", "FEEDBACK_SUBMISSION", "FEEDBACK_SURVEY",
	// "FILE_MANAGER_FILE", "FILE_MANAGER_FOLDER", "FOLDER", "FORECAST", "FORM",
	// "FORM_SUBMISSION_INBOUNDDB", "GOAL_TARGET", "GOAL_TARGET_GROUP",
	// "GOAL_TEMPLATE", "GSC_PROPERTY", "HUB", "IMPORT", "INVOICE", "KEYWORD",
	// "KNOWLEDGE_ARTICLE", "LANDING_PAGE", "LEAD", "LINE_ITEM", "MARKETING_CALENDAR",
	// "MARKETING_CAMPAIGN_UTM", "MARKETING_EMAIL", "MARKETING_EVENT",
	// "MARKETING_EVENT_ATTENDANCE", "MARKETING_SMS", "MEDIA_BRIDGE", "MEETING_EVENT",
	// "MIC", "NOTE", "OBJECT_LIST", "ORDER", "OWNER", "PARTNER_ACCOUNT",
	// "PARTNER_CLIENT", "PARTNER_CLIENT_REVENUE", "PARTNER_SERVICE", "PAYMENT_LINK",
	// "PAYMENT_SCHEDULE", "PAYMENT_SCHEDULE_INSTALLMENT", "PERMISSIONS_TESTING",
	// "PLAYBOOK", "PLAYBOOK_QUESTION", "PLAYBOOK_SUBMISSION",
	// "PLAYBOOK_SUBMISSION_ANSWER", "PLAYLIST", "PLAYLIST_FOLDER", "PODCAST_EPISODE",
	// "PORTAL", "PORTAL_OBJECT_SYNC_MESSAGE", "POSTAL_MAIL", "PRIVACY_SCANNER_COOKIE",
	// "PRODUCT", "PRODUCT_OR_FOLDER", "PROPERTY_INFO",
	// "PROSPECTING_AGENT_CONTACT_ASSIGNMENT", "PUBLISHING_TASK",
	// "QUARANTINED_SUBMISSION", "QUOTA", "QUOTE", "QUOTE_FIELD", "QUOTE_MODULE",
	// "QUOTE_MODULE_FIELD", "QUOTE_TEMPLATE", "RESTORABLE_CRM_OBJECT", "ROSTER",
	// "ROSTER_MEMBER", "SALES_DOCUMENT", "SALES_TASK", "SALES_WORKLOAD",
	// "SALESFORCE_SYNC_ERROR", "SCHEDULING_PAGE", "SCHEMAS_BACKEND_TEST",
	// "SCORE_CONFIGURATION", "SEQUENCE", "SEQUENCE_ENROLLMENT", "SEQUENCE_STEP",
	// "SEQUENCE_STEP_ENROLLMENT", "SERVICE", "SITE_PAGE", "SNIPPET",
	// "SOCIAL_BROADCAST", "SOCIAL_CHANNEL", "SOCIAL_POST", "SOCIAL_PROFILE",
	// "SOX_PROTECTED_DUMMY_TYPE", "SOX_PROTECTED_TEST_TYPE", "SUBMISSION_TAG",
	// "SUBSCRIPTION", "TASK", "TASK_TEMPLATE", "TAX", "TEMPLATE", "TICKET", "UNKNOWN",
	// "UNSUBSCRIBE", "USER", "VIEW", "VIEW_BLOCK", "WEB_INTERACTIVE".
	ReferencedObjectType FieldTypeDefinitionReferencedObjectType `json:"referencedObjectType"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ExternalOptions              respjson.Field
		Name                         respjson.Field
		Options                      respjson.Field
		Schema                       respjson.Field
		Type                         respjson.Field
		UseChirp                     respjson.Field
		Description                  respjson.Field
		ExternalOptionsReferenceType respjson.Field
		FieldType                    respjson.Field
		HelpText                     respjson.Field
		Label                        respjson.Field
		OptionsURL                   respjson.Field
		ReferencedObjectType         respjson.Field
		ExtraFields                  map[string]respjson.Field
		raw                          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FieldTypeDefinition) RawJSON() string { return r.JSON.raw }
func (r *FieldTypeDefinition) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this FieldTypeDefinition to a FieldTypeDefinitionParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// FieldTypeDefinitionParam.Overrides()
func (r FieldTypeDefinition) ToParam() FieldTypeDefinitionParam {
	return param.Override[FieldTypeDefinitionParam](json.RawMessage(r.RawJSON()))
}

// FieldTypeDefinitionSchemaUnion contains all possible properties and values from
// [IntegerFieldSchema], [LongFieldSchema], [DoubleFieldSchema],
// [StringFieldSchema], [BooleanFieldSchema], [ArrayFieldSchema],
// [ObjectFieldSchema].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type FieldTypeDefinitionSchemaUnion struct {
	Type string `json:"type"`
	// This field is a union of [int64], [int64], [float64]
	Maximum FieldTypeDefinitionSchemaUnionMaximum `json:"maximum"`
	// This field is a union of [int64], [int64], [float64]
	Minimum FieldTypeDefinitionSchemaUnionMinimum `json:"minimum"`
	// This field is from variant [StringFieldSchema].
	Format StringFieldSchemaFormat `json:"format"`
	// This field is from variant [ArrayFieldSchema].
	Items ArrayFieldSchemaItemsUnion `json:"items"`
	// This field is from variant [ObjectFieldSchema].
	Properties any `json:"properties"`
	JSON       struct {
		Type       respjson.Field
		Maximum    respjson.Field
		Minimum    respjson.Field
		Format     respjson.Field
		Items      respjson.Field
		Properties respjson.Field
		raw        string
	} `json:"-"`
}

func (u FieldTypeDefinitionSchemaUnion) AsInteger() (v IntegerFieldSchema) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FieldTypeDefinitionSchemaUnion) AsLong() (v LongFieldSchema) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FieldTypeDefinitionSchemaUnion) AsDouble() (v DoubleFieldSchema) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FieldTypeDefinitionSchemaUnion) AsString() (v StringFieldSchema) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FieldTypeDefinitionSchemaUnion) AsBoolean() (v BooleanFieldSchema) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FieldTypeDefinitionSchemaUnion) AsArray() (v ArrayFieldSchema) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FieldTypeDefinitionSchemaUnion) AsObject() (v ObjectFieldSchema) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u FieldTypeDefinitionSchemaUnion) RawJSON() string { return u.JSON.raw }

func (r *FieldTypeDefinitionSchemaUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// FieldTypeDefinitionSchemaUnionMaximum is an implicit subunion of
// [FieldTypeDefinitionSchemaUnion]. FieldTypeDefinitionSchemaUnionMaximum provides
// convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [FieldTypeDefinitionSchemaUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfInt OfFloat]
type FieldTypeDefinitionSchemaUnionMaximum struct {
	// This field will be present if the value is a [int64] instead of an object.
	OfInt int64 `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	JSON    struct {
		OfInt   respjson.Field
		OfFloat respjson.Field
		raw     string
	} `json:"-"`
}

func (r *FieldTypeDefinitionSchemaUnionMaximum) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// FieldTypeDefinitionSchemaUnionMinimum is an implicit subunion of
// [FieldTypeDefinitionSchemaUnion]. FieldTypeDefinitionSchemaUnionMinimum provides
// convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [FieldTypeDefinitionSchemaUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfInt OfFloat]
type FieldTypeDefinitionSchemaUnionMinimum struct {
	// This field will be present if the value is a [int64] instead of an object.
	OfInt int64 `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	JSON    struct {
		OfInt   respjson.Field
		OfFloat respjson.Field
		raw     string
	} `json:"-"`
}

func (r *FieldTypeDefinitionSchemaUnionMinimum) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FieldTypeDefinitionType string

const (
	FieldTypeDefinitionTypeBool              FieldTypeDefinitionType = "bool"
	FieldTypeDefinitionTypeCurrencyNumber    FieldTypeDefinitionType = "currency_number"
	FieldTypeDefinitionTypeDate              FieldTypeDefinitionType = "date"
	FieldTypeDefinitionTypeDatetime          FieldTypeDefinitionType = "datetime"
	FieldTypeDefinitionTypeEnumeration       FieldTypeDefinitionType = "enumeration"
	FieldTypeDefinitionTypeJson              FieldTypeDefinitionType = "json"
	FieldTypeDefinitionTypeNumber            FieldTypeDefinitionType = "number"
	FieldTypeDefinitionTypeObjectCoordinates FieldTypeDefinitionType = "object_coordinates"
	FieldTypeDefinitionTypePhoneNumber       FieldTypeDefinitionType = "phone_number"
	FieldTypeDefinitionTypeString            FieldTypeDefinitionType = "string"
)

type FieldTypeDefinitionFieldType string

const (
	FieldTypeDefinitionFieldTypeBooleancheckbox     FieldTypeDefinitionFieldType = "booleancheckbox"
	FieldTypeDefinitionFieldTypeCalculationEquation FieldTypeDefinitionFieldType = "calculation_equation"
	FieldTypeDefinitionFieldTypeCalculationReadTime FieldTypeDefinitionFieldType = "calculation_read_time"
	FieldTypeDefinitionFieldTypeCalculationRollup   FieldTypeDefinitionFieldType = "calculation_rollup"
	FieldTypeDefinitionFieldTypeCalculationScore    FieldTypeDefinitionFieldType = "calculation_score"
	FieldTypeDefinitionFieldTypeCheckbox            FieldTypeDefinitionFieldType = "checkbox"
	FieldTypeDefinitionFieldTypeDate                FieldTypeDefinitionFieldType = "date"
	FieldTypeDefinitionFieldTypeFile                FieldTypeDefinitionFieldType = "file"
	FieldTypeDefinitionFieldTypeHTML                FieldTypeDefinitionFieldType = "html"
	FieldTypeDefinitionFieldTypeNumber              FieldTypeDefinitionFieldType = "number"
	FieldTypeDefinitionFieldTypePhonenumber         FieldTypeDefinitionFieldType = "phonenumber"
	FieldTypeDefinitionFieldTypeRadio               FieldTypeDefinitionFieldType = "radio"
	FieldTypeDefinitionFieldTypeSelect              FieldTypeDefinitionFieldType = "select"
	FieldTypeDefinitionFieldTypeText                FieldTypeDefinitionFieldType = "text"
	FieldTypeDefinitionFieldTypeTextarea            FieldTypeDefinitionFieldType = "textarea"
	FieldTypeDefinitionFieldTypeUnknown             FieldTypeDefinitionFieldType = "unknown"
)

type FieldTypeDefinitionReferencedObjectType string

const (
	FieldTypeDefinitionReferencedObjectTypeAbandonedCart                     FieldTypeDefinitionReferencedObjectType = "ABANDONED_CART"
	FieldTypeDefinitionReferencedObjectTypeAcceptanceTest                    FieldTypeDefinitionReferencedObjectType = "ACCEPTANCE_TEST"
	FieldTypeDefinitionReferencedObjectTypeAd                                FieldTypeDefinitionReferencedObjectType = "AD"
	FieldTypeDefinitionReferencedObjectTypeAdAccount                         FieldTypeDefinitionReferencedObjectType = "AD_ACCOUNT"
	FieldTypeDefinitionReferencedObjectTypeAdCampaign                        FieldTypeDefinitionReferencedObjectType = "AD_CAMPAIGN"
	FieldTypeDefinitionReferencedObjectTypeAdGroup                           FieldTypeDefinitionReferencedObjectType = "AD_GROUP"
	FieldTypeDefinitionReferencedObjectTypeAIForecast                        FieldTypeDefinitionReferencedObjectType = "AI_FORECAST"
	FieldTypeDefinitionReferencedObjectTypeAllPages                          FieldTypeDefinitionReferencedObjectType = "ALL_PAGES"
	FieldTypeDefinitionReferencedObjectTypeApproval                          FieldTypeDefinitionReferencedObjectType = "APPROVAL"
	FieldTypeDefinitionReferencedObjectTypeApprovalStep                      FieldTypeDefinitionReferencedObjectType = "APPROVAL_STEP"
	FieldTypeDefinitionReferencedObjectTypeAttribution                       FieldTypeDefinitionReferencedObjectType = "ATTRIBUTION"
	FieldTypeDefinitionReferencedObjectTypeAudience                          FieldTypeDefinitionReferencedObjectType = "AUDIENCE"
	FieldTypeDefinitionReferencedObjectTypeAutomationJourney                 FieldTypeDefinitionReferencedObjectType = "AUTOMATION_JOURNEY"
	FieldTypeDefinitionReferencedObjectTypeAutomationPlatformFlow            FieldTypeDefinitionReferencedObjectType = "AUTOMATION_PLATFORM_FLOW"
	FieldTypeDefinitionReferencedObjectTypeAutomationPlatformFlowAction      FieldTypeDefinitionReferencedObjectType = "AUTOMATION_PLATFORM_FLOW_ACTION"
	FieldTypeDefinitionReferencedObjectTypeBetAlert                          FieldTypeDefinitionReferencedObjectType = "BET_ALERT"
	FieldTypeDefinitionReferencedObjectTypeBetDeliverableService             FieldTypeDefinitionReferencedObjectType = "BET_DELIVERABLE_SERVICE"
	FieldTypeDefinitionReferencedObjectTypeBlogListingPage                   FieldTypeDefinitionReferencedObjectType = "BLOG_LISTING_PAGE"
	FieldTypeDefinitionReferencedObjectTypeBlogPost                          FieldTypeDefinitionReferencedObjectType = "BLOG_POST"
	FieldTypeDefinitionReferencedObjectTypeCall                              FieldTypeDefinitionReferencedObjectType = "CALL"
	FieldTypeDefinitionReferencedObjectTypeCampaign                          FieldTypeDefinitionReferencedObjectType = "CAMPAIGN"
	FieldTypeDefinitionReferencedObjectTypeCampaignBudgetItem                FieldTypeDefinitionReferencedObjectType = "CAMPAIGN_BUDGET_ITEM"
	FieldTypeDefinitionReferencedObjectTypeCampaignSpendItem                 FieldTypeDefinitionReferencedObjectType = "CAMPAIGN_SPEND_ITEM"
	FieldTypeDefinitionReferencedObjectTypeCampaignStep                      FieldTypeDefinitionReferencedObjectType = "CAMPAIGN_STEP"
	FieldTypeDefinitionReferencedObjectTypeCampaignTemplate                  FieldTypeDefinitionReferencedObjectType = "CAMPAIGN_TEMPLATE"
	FieldTypeDefinitionReferencedObjectTypeCampaignTemplateStep              FieldTypeDefinitionReferencedObjectType = "CAMPAIGN_TEMPLATE_STEP"
	FieldTypeDefinitionReferencedObjectTypeCart                              FieldTypeDefinitionReferencedObjectType = "CART"
	FieldTypeDefinitionReferencedObjectTypeCaseStudy                         FieldTypeDefinitionReferencedObjectType = "CASE_STUDY"
	FieldTypeDefinitionReferencedObjectTypeChatflow                          FieldTypeDefinitionReferencedObjectType = "CHATFLOW"
	FieldTypeDefinitionReferencedObjectTypeClip                              FieldTypeDefinitionReferencedObjectType = "CLIP"
	FieldTypeDefinitionReferencedObjectTypeCmsURL                            FieldTypeDefinitionReferencedObjectType = "CMS_URL"
	FieldTypeDefinitionReferencedObjectTypeComboEventConfiguration           FieldTypeDefinitionReferencedObjectType = "COMBO_EVENT_CONFIGURATION"
	FieldTypeDefinitionReferencedObjectTypeCommercePayment                   FieldTypeDefinitionReferencedObjectType = "COMMERCE_PAYMENT"
	FieldTypeDefinitionReferencedObjectTypeCommunication                     FieldTypeDefinitionReferencedObjectType = "COMMUNICATION"
	FieldTypeDefinitionReferencedObjectTypeCompany                           FieldTypeDefinitionReferencedObjectType = "COMPANY"
	FieldTypeDefinitionReferencedObjectTypeContact                           FieldTypeDefinitionReferencedObjectType = "CONTACT"
	FieldTypeDefinitionReferencedObjectTypeContactCreateAttribution          FieldTypeDefinitionReferencedObjectType = "CONTACT_CREATE_ATTRIBUTION"
	FieldTypeDefinitionReferencedObjectTypeContent                           FieldTypeDefinitionReferencedObjectType = "CONTENT"
	FieldTypeDefinitionReferencedObjectTypeContentAudit                      FieldTypeDefinitionReferencedObjectType = "CONTENT_AUDIT"
	FieldTypeDefinitionReferencedObjectTypeContentAuditPage                  FieldTypeDefinitionReferencedObjectType = "CONTENT_AUDIT_PAGE"
	FieldTypeDefinitionReferencedObjectTypeConversation                      FieldTypeDefinitionReferencedObjectType = "CONVERSATION"
	FieldTypeDefinitionReferencedObjectTypeConversationInbox                 FieldTypeDefinitionReferencedObjectType = "CONVERSATION_INBOX"
	FieldTypeDefinitionReferencedObjectTypeConversationSession               FieldTypeDefinitionReferencedObjectType = "CONVERSATION_SESSION"
	FieldTypeDefinitionReferencedObjectTypeCrmObjectsDummyType               FieldTypeDefinitionReferencedObjectType = "CRM_OBJECTS_DUMMY_TYPE"
	FieldTypeDefinitionReferencedObjectTypeCrmPipelinesDummyType             FieldTypeDefinitionReferencedObjectType = "CRM_PIPELINES_DUMMY_TYPE"
	FieldTypeDefinitionReferencedObjectTypeCta                               FieldTypeDefinitionReferencedObjectType = "CTA"
	FieldTypeDefinitionReferencedObjectTypeCtaVariant                        FieldTypeDefinitionReferencedObjectType = "CTA_VARIANT"
	FieldTypeDefinitionReferencedObjectTypeDataPrivacyConsent                FieldTypeDefinitionReferencedObjectType = "DATA_PRIVACY_CONSENT"
	FieldTypeDefinitionReferencedObjectTypeDataSyncState                     FieldTypeDefinitionReferencedObjectType = "DATA_SYNC_STATE"
	FieldTypeDefinitionReferencedObjectTypeDeal                              FieldTypeDefinitionReferencedObjectType = "DEAL"
	FieldTypeDefinitionReferencedObjectTypeDealCreateAttribution             FieldTypeDefinitionReferencedObjectType = "DEAL_CREATE_ATTRIBUTION"
	FieldTypeDefinitionReferencedObjectTypeDealRegistration                  FieldTypeDefinitionReferencedObjectType = "DEAL_REGISTRATION"
	FieldTypeDefinitionReferencedObjectTypeDealSplit                         FieldTypeDefinitionReferencedObjectType = "DEAL_SPLIT"
	FieldTypeDefinitionReferencedObjectTypeDiscount                          FieldTypeDefinitionReferencedObjectType = "DISCOUNT"
	FieldTypeDefinitionReferencedObjectTypeDiscountCode                      FieldTypeDefinitionReferencedObjectType = "DISCOUNT_CODE"
	FieldTypeDefinitionReferencedObjectTypeDiscountTemplate                  FieldTypeDefinitionReferencedObjectType = "DISCOUNT_TEMPLATE"
	FieldTypeDefinitionReferencedObjectTypeEmail                             FieldTypeDefinitionReferencedObjectType = "EMAIL"
	FieldTypeDefinitionReferencedObjectTypeEngagement                        FieldTypeDefinitionReferencedObjectType = "ENGAGEMENT"
	FieldTypeDefinitionReferencedObjectTypeExport                            FieldTypeDefinitionReferencedObjectType = "EXPORT"
	FieldTypeDefinitionReferencedObjectTypeExternalWebURL                    FieldTypeDefinitionReferencedObjectType = "EXTERNAL_WEB_URL"
	FieldTypeDefinitionReferencedObjectTypeFee                               FieldTypeDefinitionReferencedObjectType = "FEE"
	FieldTypeDefinitionReferencedObjectTypeFeedbackSubmission                FieldTypeDefinitionReferencedObjectType = "FEEDBACK_SUBMISSION"
	FieldTypeDefinitionReferencedObjectTypeFeedbackSurvey                    FieldTypeDefinitionReferencedObjectType = "FEEDBACK_SURVEY"
	FieldTypeDefinitionReferencedObjectTypeFileManagerFile                   FieldTypeDefinitionReferencedObjectType = "FILE_MANAGER_FILE"
	FieldTypeDefinitionReferencedObjectTypeFileManagerFolder                 FieldTypeDefinitionReferencedObjectType = "FILE_MANAGER_FOLDER"
	FieldTypeDefinitionReferencedObjectTypeFolder                            FieldTypeDefinitionReferencedObjectType = "FOLDER"
	FieldTypeDefinitionReferencedObjectTypeForecast                          FieldTypeDefinitionReferencedObjectType = "FORECAST"
	FieldTypeDefinitionReferencedObjectTypeForm                              FieldTypeDefinitionReferencedObjectType = "FORM"
	FieldTypeDefinitionReferencedObjectTypeFormSubmissionInbounddb           FieldTypeDefinitionReferencedObjectType = "FORM_SUBMISSION_INBOUNDDB"
	FieldTypeDefinitionReferencedObjectTypeGoalTarget                        FieldTypeDefinitionReferencedObjectType = "GOAL_TARGET"
	FieldTypeDefinitionReferencedObjectTypeGoalTargetGroup                   FieldTypeDefinitionReferencedObjectType = "GOAL_TARGET_GROUP"
	FieldTypeDefinitionReferencedObjectTypeGoalTemplate                      FieldTypeDefinitionReferencedObjectType = "GOAL_TEMPLATE"
	FieldTypeDefinitionReferencedObjectTypeGscProperty                       FieldTypeDefinitionReferencedObjectType = "GSC_PROPERTY"
	FieldTypeDefinitionReferencedObjectTypeHub                               FieldTypeDefinitionReferencedObjectType = "HUB"
	FieldTypeDefinitionReferencedObjectTypeImport                            FieldTypeDefinitionReferencedObjectType = "IMPORT"
	FieldTypeDefinitionReferencedObjectTypeInvoice                           FieldTypeDefinitionReferencedObjectType = "INVOICE"
	FieldTypeDefinitionReferencedObjectTypeKeyword                           FieldTypeDefinitionReferencedObjectType = "KEYWORD"
	FieldTypeDefinitionReferencedObjectTypeKnowledgeArticle                  FieldTypeDefinitionReferencedObjectType = "KNOWLEDGE_ARTICLE"
	FieldTypeDefinitionReferencedObjectTypeLandingPage                       FieldTypeDefinitionReferencedObjectType = "LANDING_PAGE"
	FieldTypeDefinitionReferencedObjectTypeLead                              FieldTypeDefinitionReferencedObjectType = "LEAD"
	FieldTypeDefinitionReferencedObjectTypeLineItem                          FieldTypeDefinitionReferencedObjectType = "LINE_ITEM"
	FieldTypeDefinitionReferencedObjectTypeMarketingCalendar                 FieldTypeDefinitionReferencedObjectType = "MARKETING_CALENDAR"
	FieldTypeDefinitionReferencedObjectTypeMarketingCampaignUtm              FieldTypeDefinitionReferencedObjectType = "MARKETING_CAMPAIGN_UTM"
	FieldTypeDefinitionReferencedObjectTypeMarketingEmail                    FieldTypeDefinitionReferencedObjectType = "MARKETING_EMAIL"
	FieldTypeDefinitionReferencedObjectTypeMarketingEvent                    FieldTypeDefinitionReferencedObjectType = "MARKETING_EVENT"
	FieldTypeDefinitionReferencedObjectTypeMarketingEventAttendance          FieldTypeDefinitionReferencedObjectType = "MARKETING_EVENT_ATTENDANCE"
	FieldTypeDefinitionReferencedObjectTypeMarketingSMS                      FieldTypeDefinitionReferencedObjectType = "MARKETING_SMS"
	FieldTypeDefinitionReferencedObjectTypeMediaBridge                       FieldTypeDefinitionReferencedObjectType = "MEDIA_BRIDGE"
	FieldTypeDefinitionReferencedObjectTypeMeetingEvent                      FieldTypeDefinitionReferencedObjectType = "MEETING_EVENT"
	FieldTypeDefinitionReferencedObjectTypeMic                               FieldTypeDefinitionReferencedObjectType = "MIC"
	FieldTypeDefinitionReferencedObjectTypeNote                              FieldTypeDefinitionReferencedObjectType = "NOTE"
	FieldTypeDefinitionReferencedObjectTypeObjectList                        FieldTypeDefinitionReferencedObjectType = "OBJECT_LIST"
	FieldTypeDefinitionReferencedObjectTypeOrder                             FieldTypeDefinitionReferencedObjectType = "ORDER"
	FieldTypeDefinitionReferencedObjectTypeOwner                             FieldTypeDefinitionReferencedObjectType = "OWNER"
	FieldTypeDefinitionReferencedObjectTypePartnerAccount                    FieldTypeDefinitionReferencedObjectType = "PARTNER_ACCOUNT"
	FieldTypeDefinitionReferencedObjectTypePartnerClient                     FieldTypeDefinitionReferencedObjectType = "PARTNER_CLIENT"
	FieldTypeDefinitionReferencedObjectTypePartnerClientRevenue              FieldTypeDefinitionReferencedObjectType = "PARTNER_CLIENT_REVENUE"
	FieldTypeDefinitionReferencedObjectTypePartnerService                    FieldTypeDefinitionReferencedObjectType = "PARTNER_SERVICE"
	FieldTypeDefinitionReferencedObjectTypePaymentLink                       FieldTypeDefinitionReferencedObjectType = "PAYMENT_LINK"
	FieldTypeDefinitionReferencedObjectTypePaymentSchedule                   FieldTypeDefinitionReferencedObjectType = "PAYMENT_SCHEDULE"
	FieldTypeDefinitionReferencedObjectTypePaymentScheduleInstallment        FieldTypeDefinitionReferencedObjectType = "PAYMENT_SCHEDULE_INSTALLMENT"
	FieldTypeDefinitionReferencedObjectTypePermissionsTesting                FieldTypeDefinitionReferencedObjectType = "PERMISSIONS_TESTING"
	FieldTypeDefinitionReferencedObjectTypePlaybook                          FieldTypeDefinitionReferencedObjectType = "PLAYBOOK"
	FieldTypeDefinitionReferencedObjectTypePlaybookQuestion                  FieldTypeDefinitionReferencedObjectType = "PLAYBOOK_QUESTION"
	FieldTypeDefinitionReferencedObjectTypePlaybookSubmission                FieldTypeDefinitionReferencedObjectType = "PLAYBOOK_SUBMISSION"
	FieldTypeDefinitionReferencedObjectTypePlaybookSubmissionAnswer          FieldTypeDefinitionReferencedObjectType = "PLAYBOOK_SUBMISSION_ANSWER"
	FieldTypeDefinitionReferencedObjectTypePlaylist                          FieldTypeDefinitionReferencedObjectType = "PLAYLIST"
	FieldTypeDefinitionReferencedObjectTypePlaylistFolder                    FieldTypeDefinitionReferencedObjectType = "PLAYLIST_FOLDER"
	FieldTypeDefinitionReferencedObjectTypePodcastEpisode                    FieldTypeDefinitionReferencedObjectType = "PODCAST_EPISODE"
	FieldTypeDefinitionReferencedObjectTypePortal                            FieldTypeDefinitionReferencedObjectType = "PORTAL"
	FieldTypeDefinitionReferencedObjectTypePortalObjectSyncMessage           FieldTypeDefinitionReferencedObjectType = "PORTAL_OBJECT_SYNC_MESSAGE"
	FieldTypeDefinitionReferencedObjectTypePostalMail                        FieldTypeDefinitionReferencedObjectType = "POSTAL_MAIL"
	FieldTypeDefinitionReferencedObjectTypePrivacyScannerCookie              FieldTypeDefinitionReferencedObjectType = "PRIVACY_SCANNER_COOKIE"
	FieldTypeDefinitionReferencedObjectTypeProduct                           FieldTypeDefinitionReferencedObjectType = "PRODUCT"
	FieldTypeDefinitionReferencedObjectTypeProductOrFolder                   FieldTypeDefinitionReferencedObjectType = "PRODUCT_OR_FOLDER"
	FieldTypeDefinitionReferencedObjectTypePropertyInfo                      FieldTypeDefinitionReferencedObjectType = "PROPERTY_INFO"
	FieldTypeDefinitionReferencedObjectTypeProspectingAgentContactAssignment FieldTypeDefinitionReferencedObjectType = "PROSPECTING_AGENT_CONTACT_ASSIGNMENT"
	FieldTypeDefinitionReferencedObjectTypePublishingTask                    FieldTypeDefinitionReferencedObjectType = "PUBLISHING_TASK"
	FieldTypeDefinitionReferencedObjectTypeQuarantinedSubmission             FieldTypeDefinitionReferencedObjectType = "QUARANTINED_SUBMISSION"
	FieldTypeDefinitionReferencedObjectTypeQuota                             FieldTypeDefinitionReferencedObjectType = "QUOTA"
	FieldTypeDefinitionReferencedObjectTypeQuote                             FieldTypeDefinitionReferencedObjectType = "QUOTE"
	FieldTypeDefinitionReferencedObjectTypeQuoteField                        FieldTypeDefinitionReferencedObjectType = "QUOTE_FIELD"
	FieldTypeDefinitionReferencedObjectTypeQuoteModule                       FieldTypeDefinitionReferencedObjectType = "QUOTE_MODULE"
	FieldTypeDefinitionReferencedObjectTypeQuoteModuleField                  FieldTypeDefinitionReferencedObjectType = "QUOTE_MODULE_FIELD"
	FieldTypeDefinitionReferencedObjectTypeQuoteTemplate                     FieldTypeDefinitionReferencedObjectType = "QUOTE_TEMPLATE"
	FieldTypeDefinitionReferencedObjectTypeRestorableCrmObject               FieldTypeDefinitionReferencedObjectType = "RESTORABLE_CRM_OBJECT"
	FieldTypeDefinitionReferencedObjectTypeRoster                            FieldTypeDefinitionReferencedObjectType = "ROSTER"
	FieldTypeDefinitionReferencedObjectTypeRosterMember                      FieldTypeDefinitionReferencedObjectType = "ROSTER_MEMBER"
	FieldTypeDefinitionReferencedObjectTypeSalesDocument                     FieldTypeDefinitionReferencedObjectType = "SALES_DOCUMENT"
	FieldTypeDefinitionReferencedObjectTypeSalesTask                         FieldTypeDefinitionReferencedObjectType = "SALES_TASK"
	FieldTypeDefinitionReferencedObjectTypeSalesWorkload                     FieldTypeDefinitionReferencedObjectType = "SALES_WORKLOAD"
	FieldTypeDefinitionReferencedObjectTypeSalesforceSyncError               FieldTypeDefinitionReferencedObjectType = "SALESFORCE_SYNC_ERROR"
	FieldTypeDefinitionReferencedObjectTypeSchedulingPage                    FieldTypeDefinitionReferencedObjectType = "SCHEDULING_PAGE"
	FieldTypeDefinitionReferencedObjectTypeSchemasBackendTest                FieldTypeDefinitionReferencedObjectType = "SCHEMAS_BACKEND_TEST"
	FieldTypeDefinitionReferencedObjectTypeScoreConfiguration                FieldTypeDefinitionReferencedObjectType = "SCORE_CONFIGURATION"
	FieldTypeDefinitionReferencedObjectTypeSequence                          FieldTypeDefinitionReferencedObjectType = "SEQUENCE"
	FieldTypeDefinitionReferencedObjectTypeSequenceEnrollment                FieldTypeDefinitionReferencedObjectType = "SEQUENCE_ENROLLMENT"
	FieldTypeDefinitionReferencedObjectTypeSequenceStep                      FieldTypeDefinitionReferencedObjectType = "SEQUENCE_STEP"
	FieldTypeDefinitionReferencedObjectTypeSequenceStepEnrollment            FieldTypeDefinitionReferencedObjectType = "SEQUENCE_STEP_ENROLLMENT"
	FieldTypeDefinitionReferencedObjectTypeService                           FieldTypeDefinitionReferencedObjectType = "SERVICE"
	FieldTypeDefinitionReferencedObjectTypeSitePage                          FieldTypeDefinitionReferencedObjectType = "SITE_PAGE"
	FieldTypeDefinitionReferencedObjectTypeSnippet                           FieldTypeDefinitionReferencedObjectType = "SNIPPET"
	FieldTypeDefinitionReferencedObjectTypeSocialBroadcast                   FieldTypeDefinitionReferencedObjectType = "SOCIAL_BROADCAST"
	FieldTypeDefinitionReferencedObjectTypeSocialChannel                     FieldTypeDefinitionReferencedObjectType = "SOCIAL_CHANNEL"
	FieldTypeDefinitionReferencedObjectTypeSocialPost                        FieldTypeDefinitionReferencedObjectType = "SOCIAL_POST"
	FieldTypeDefinitionReferencedObjectTypeSocialProfile                     FieldTypeDefinitionReferencedObjectType = "SOCIAL_PROFILE"
	FieldTypeDefinitionReferencedObjectTypeSoxProtectedDummyType             FieldTypeDefinitionReferencedObjectType = "SOX_PROTECTED_DUMMY_TYPE"
	FieldTypeDefinitionReferencedObjectTypeSoxProtectedTestType              FieldTypeDefinitionReferencedObjectType = "SOX_PROTECTED_TEST_TYPE"
	FieldTypeDefinitionReferencedObjectTypeSubmissionTag                     FieldTypeDefinitionReferencedObjectType = "SUBMISSION_TAG"
	FieldTypeDefinitionReferencedObjectTypeSubscription                      FieldTypeDefinitionReferencedObjectType = "SUBSCRIPTION"
	FieldTypeDefinitionReferencedObjectTypeTask                              FieldTypeDefinitionReferencedObjectType = "TASK"
	FieldTypeDefinitionReferencedObjectTypeTaskTemplate                      FieldTypeDefinitionReferencedObjectType = "TASK_TEMPLATE"
	FieldTypeDefinitionReferencedObjectTypeTax                               FieldTypeDefinitionReferencedObjectType = "TAX"
	FieldTypeDefinitionReferencedObjectTypeTemplate                          FieldTypeDefinitionReferencedObjectType = "TEMPLATE"
	FieldTypeDefinitionReferencedObjectTypeTicket                            FieldTypeDefinitionReferencedObjectType = "TICKET"
	FieldTypeDefinitionReferencedObjectTypeUnknown                           FieldTypeDefinitionReferencedObjectType = "UNKNOWN"
	FieldTypeDefinitionReferencedObjectTypeUnsubscribe                       FieldTypeDefinitionReferencedObjectType = "UNSUBSCRIBE"
	FieldTypeDefinitionReferencedObjectTypeUser                              FieldTypeDefinitionReferencedObjectType = "USER"
	FieldTypeDefinitionReferencedObjectTypeView                              FieldTypeDefinitionReferencedObjectType = "VIEW"
	FieldTypeDefinitionReferencedObjectTypeViewBlock                         FieldTypeDefinitionReferencedObjectType = "VIEW_BLOCK"
	FieldTypeDefinitionReferencedObjectTypeWebInteractive                    FieldTypeDefinitionReferencedObjectType = "WEB_INTERACTIVE"
)

// The properties ExternalOptions, Name, Options, Schema, Type, UseChirp are
// required.
type FieldTypeDefinitionParam struct {
	ExternalOptions bool                                `json:"externalOptions" api:"required"`
	Name            string                              `json:"name" api:"required"`
	Options         []events.OptionParam                `json:"options,omitzero" api:"required"`
	Schema          FieldTypeDefinitionSchemaUnionParam `json:"schema,omitzero" api:"required"`
	// Any of "bool", "currency_number", "date", "datetime", "enumeration", "json",
	// "number", "object_coordinates", "phone_number", "string".
	Type                         FieldTypeDefinitionType `json:"type,omitzero" api:"required"`
	UseChirp                     bool                    `json:"useChirp" api:"required"`
	Description                  param.Opt[string]       `json:"description,omitzero"`
	ExternalOptionsReferenceType param.Opt[string]       `json:"externalOptionsReferenceType,omitzero"`
	HelpText                     param.Opt[string]       `json:"helpText,omitzero"`
	Label                        param.Opt[string]       `json:"label,omitzero"`
	OptionsURL                   param.Opt[string]       `json:"optionsUrl,omitzero"`
	// Any of "booleancheckbox", "calculation_equation", "calculation_read_time",
	// "calculation_rollup", "calculation_score", "checkbox", "date", "file", "html",
	// "number", "phonenumber", "radio", "select", "text", "textarea", "unknown".
	FieldType FieldTypeDefinitionFieldType `json:"fieldType,omitzero"`
	// Any of "ABANDONED_CART", "ACCEPTANCE_TEST", "AD", "AD_ACCOUNT", "AD_CAMPAIGN",
	// "AD_GROUP", "AI_FORECAST", "ALL_PAGES", "APPROVAL", "APPROVAL_STEP",
	// "ATTRIBUTION", "AUDIENCE", "AUTOMATION_JOURNEY", "AUTOMATION_PLATFORM_FLOW",
	// "AUTOMATION_PLATFORM_FLOW_ACTION", "BET_ALERT", "BET_DELIVERABLE_SERVICE",
	// "BLOG_LISTING_PAGE", "BLOG_POST", "CALL", "CAMPAIGN", "CAMPAIGN_BUDGET_ITEM",
	// "CAMPAIGN_SPEND_ITEM", "CAMPAIGN_STEP", "CAMPAIGN_TEMPLATE",
	// "CAMPAIGN_TEMPLATE_STEP", "CART", "CASE_STUDY", "CHATFLOW", "CLIP", "CMS_URL",
	// "COMBO_EVENT_CONFIGURATION", "COMMERCE_PAYMENT", "COMMUNICATION", "COMPANY",
	// "CONTACT", "CONTACT_CREATE_ATTRIBUTION", "CONTENT", "CONTENT_AUDIT",
	// "CONTENT_AUDIT_PAGE", "CONVERSATION", "CONVERSATION_INBOX",
	// "CONVERSATION_SESSION", "CRM_OBJECTS_DUMMY_TYPE", "CRM_PIPELINES_DUMMY_TYPE",
	// "CTA", "CTA_VARIANT", "DATA_PRIVACY_CONSENT", "DATA_SYNC_STATE", "DEAL",
	// "DEAL_CREATE_ATTRIBUTION", "DEAL_REGISTRATION", "DEAL_SPLIT", "DISCOUNT",
	// "DISCOUNT_CODE", "DISCOUNT_TEMPLATE", "EMAIL", "ENGAGEMENT", "EXPORT",
	// "EXTERNAL_WEB_URL", "FEE", "FEEDBACK_SUBMISSION", "FEEDBACK_SURVEY",
	// "FILE_MANAGER_FILE", "FILE_MANAGER_FOLDER", "FOLDER", "FORECAST", "FORM",
	// "FORM_SUBMISSION_INBOUNDDB", "GOAL_TARGET", "GOAL_TARGET_GROUP",
	// "GOAL_TEMPLATE", "GSC_PROPERTY", "HUB", "IMPORT", "INVOICE", "KEYWORD",
	// "KNOWLEDGE_ARTICLE", "LANDING_PAGE", "LEAD", "LINE_ITEM", "MARKETING_CALENDAR",
	// "MARKETING_CAMPAIGN_UTM", "MARKETING_EMAIL", "MARKETING_EVENT",
	// "MARKETING_EVENT_ATTENDANCE", "MARKETING_SMS", "MEDIA_BRIDGE", "MEETING_EVENT",
	// "MIC", "NOTE", "OBJECT_LIST", "ORDER", "OWNER", "PARTNER_ACCOUNT",
	// "PARTNER_CLIENT", "PARTNER_CLIENT_REVENUE", "PARTNER_SERVICE", "PAYMENT_LINK",
	// "PAYMENT_SCHEDULE", "PAYMENT_SCHEDULE_INSTALLMENT", "PERMISSIONS_TESTING",
	// "PLAYBOOK", "PLAYBOOK_QUESTION", "PLAYBOOK_SUBMISSION",
	// "PLAYBOOK_SUBMISSION_ANSWER", "PLAYLIST", "PLAYLIST_FOLDER", "PODCAST_EPISODE",
	// "PORTAL", "PORTAL_OBJECT_SYNC_MESSAGE", "POSTAL_MAIL", "PRIVACY_SCANNER_COOKIE",
	// "PRODUCT", "PRODUCT_OR_FOLDER", "PROPERTY_INFO",
	// "PROSPECTING_AGENT_CONTACT_ASSIGNMENT", "PUBLISHING_TASK",
	// "QUARANTINED_SUBMISSION", "QUOTA", "QUOTE", "QUOTE_FIELD", "QUOTE_MODULE",
	// "QUOTE_MODULE_FIELD", "QUOTE_TEMPLATE", "RESTORABLE_CRM_OBJECT", "ROSTER",
	// "ROSTER_MEMBER", "SALES_DOCUMENT", "SALES_TASK", "SALES_WORKLOAD",
	// "SALESFORCE_SYNC_ERROR", "SCHEDULING_PAGE", "SCHEMAS_BACKEND_TEST",
	// "SCORE_CONFIGURATION", "SEQUENCE", "SEQUENCE_ENROLLMENT", "SEQUENCE_STEP",
	// "SEQUENCE_STEP_ENROLLMENT", "SERVICE", "SITE_PAGE", "SNIPPET",
	// "SOCIAL_BROADCAST", "SOCIAL_CHANNEL", "SOCIAL_POST", "SOCIAL_PROFILE",
	// "SOX_PROTECTED_DUMMY_TYPE", "SOX_PROTECTED_TEST_TYPE", "SUBMISSION_TAG",
	// "SUBSCRIPTION", "TASK", "TASK_TEMPLATE", "TAX", "TEMPLATE", "TICKET", "UNKNOWN",
	// "UNSUBSCRIBE", "USER", "VIEW", "VIEW_BLOCK", "WEB_INTERACTIVE".
	ReferencedObjectType FieldTypeDefinitionReferencedObjectType `json:"referencedObjectType,omitzero"`
	paramObj
}

func (r FieldTypeDefinitionParam) MarshalJSON() (data []byte, err error) {
	type shadow FieldTypeDefinitionParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FieldTypeDefinitionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type FieldTypeDefinitionSchemaUnionParam struct {
	OfInteger *IntegerFieldSchemaParam `json:",omitzero,inline"`
	OfLong    *LongFieldSchemaParam    `json:",omitzero,inline"`
	OfDouble  *DoubleFieldSchemaParam  `json:",omitzero,inline"`
	OfString  *StringFieldSchemaParam  `json:",omitzero,inline"`
	OfBoolean *BooleanFieldSchemaParam `json:",omitzero,inline"`
	OfArray   *ArrayFieldSchemaParam   `json:",omitzero,inline"`
	OfObject  *ObjectFieldSchemaParam  `json:",omitzero,inline"`
	paramUnion
}

func (u FieldTypeDefinitionSchemaUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfInteger,
		u.OfLong,
		u.OfDouble,
		u.OfString,
		u.OfBoolean,
		u.OfArray,
		u.OfObject)
}
func (u *FieldTypeDefinitionSchemaUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *FieldTypeDefinitionSchemaUnionParam) asAny() any {
	if !param.IsOmitted(u.OfInteger) {
		return u.OfInteger
	} else if !param.IsOmitted(u.OfLong) {
		return u.OfLong
	} else if !param.IsOmitted(u.OfDouble) {
		return u.OfDouble
	} else if !param.IsOmitted(u.OfString) {
		return u.OfString
	} else if !param.IsOmitted(u.OfBoolean) {
		return u.OfBoolean
	} else if !param.IsOmitted(u.OfArray) {
		return u.OfArray
	} else if !param.IsOmitted(u.OfObject) {
		return u.OfObject
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u FieldTypeDefinitionSchemaUnionParam) GetFormat() *string {
	if vt := u.OfString; vt != nil {
		return (*string)(&vt.Format)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u FieldTypeDefinitionSchemaUnionParam) GetItems() *ArrayFieldSchemaItemsUnionParam {
	if vt := u.OfArray; vt != nil {
		return &vt.Items
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u FieldTypeDefinitionSchemaUnionParam) GetProperties() *any {
	if vt := u.OfObject; vt != nil {
		return &vt.Properties
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u FieldTypeDefinitionSchemaUnionParam) GetType() *string {
	if vt := u.OfInteger; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfLong; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfDouble; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfString; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfBoolean; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfArray; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfObject; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

// Returns a subunion which exports methods to access subproperties
//
// Or use AsAny() to get the underlying value
func (u FieldTypeDefinitionSchemaUnionParam) GetMaximum() (res fieldTypeDefinitionSchemaUnionParamMaximum) {
	if vt := u.OfInteger; vt != nil && vt.Maximum.Valid() {
		res.any = &vt.Maximum.Value
	} else if vt := u.OfLong; vt != nil && vt.Maximum.Valid() {
		res.any = &vt.Maximum.Value
	} else if vt := u.OfDouble; vt != nil && vt.Maximum.Valid() {
		res.any = &vt.Maximum.Value
	}
	return
}

// Can have the runtime types [*int64], [*float64]
type fieldTypeDefinitionSchemaUnionParamMaximum struct{ any }

// Use the following switch statement to get the type of the union:
//
//	switch u.AsAny().(type) {
//	case *int64:
//	case *float64:
//	default:
//	    fmt.Errorf("not present")
//	}
func (u fieldTypeDefinitionSchemaUnionParamMaximum) AsAny() any { return u.any }

// Returns a subunion which exports methods to access subproperties
//
// Or use AsAny() to get the underlying value
func (u FieldTypeDefinitionSchemaUnionParam) GetMinimum() (res fieldTypeDefinitionSchemaUnionParamMinimum) {
	if vt := u.OfInteger; vt != nil && vt.Minimum.Valid() {
		res.any = &vt.Minimum.Value
	} else if vt := u.OfLong; vt != nil && vt.Minimum.Valid() {
		res.any = &vt.Minimum.Value
	} else if vt := u.OfDouble; vt != nil && vt.Minimum.Valid() {
		res.any = &vt.Minimum.Value
	}
	return
}

// Can have the runtime types [*int64], [*float64]
type fieldTypeDefinitionSchemaUnionParamMinimum struct{ any }

// Use the following switch statement to get the type of the union:
//
//	switch u.AsAny().(type) {
//	case *int64:
//	case *float64:
//	default:
//	    fmt.Errorf("not present")
//	}
func (u fieldTypeDefinitionSchemaUnionParamMinimum) AsAny() any { return u.any }

type IntegerFieldSchema struct {
	// Any of "INTEGER".
	Type    IntegerFieldSchemaType `json:"type" api:"required"`
	Maximum int64                  `json:"maximum"`
	Minimum int64                  `json:"minimum"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		Maximum     respjson.Field
		Minimum     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IntegerFieldSchema) RawJSON() string { return r.JSON.raw }
func (r *IntegerFieldSchema) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this IntegerFieldSchema to a IntegerFieldSchemaParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// IntegerFieldSchemaParam.Overrides()
func (r IntegerFieldSchema) ToParam() IntegerFieldSchemaParam {
	return param.Override[IntegerFieldSchemaParam](json.RawMessage(r.RawJSON()))
}

type IntegerFieldSchemaType string

const (
	IntegerFieldSchemaTypeInteger IntegerFieldSchemaType = "INTEGER"
)

// The property Type is required.
type IntegerFieldSchemaParam struct {
	// Any of "INTEGER".
	Type    IntegerFieldSchemaType `json:"type,omitzero" api:"required"`
	Maximum param.Opt[int64]       `json:"maximum,omitzero"`
	Minimum param.Opt[int64]       `json:"minimum,omitzero"`
	paramObj
}

func (r IntegerFieldSchemaParam) MarshalJSON() (data []byte, err error) {
	type shadow IntegerFieldSchemaParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *IntegerFieldSchemaParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LongFieldSchema struct {
	// Any of "LONG".
	Type    LongFieldSchemaType `json:"type" api:"required"`
	Maximum int64               `json:"maximum"`
	Minimum int64               `json:"minimum"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		Maximum     respjson.Field
		Minimum     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LongFieldSchema) RawJSON() string { return r.JSON.raw }
func (r *LongFieldSchema) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this LongFieldSchema to a LongFieldSchemaParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// LongFieldSchemaParam.Overrides()
func (r LongFieldSchema) ToParam() LongFieldSchemaParam {
	return param.Override[LongFieldSchemaParam](json.RawMessage(r.RawJSON()))
}

type LongFieldSchemaType string

const (
	LongFieldSchemaTypeLong LongFieldSchemaType = "LONG"
)

// The property Type is required.
type LongFieldSchemaParam struct {
	// Any of "LONG".
	Type    LongFieldSchemaType `json:"type,omitzero" api:"required"`
	Maximum param.Opt[int64]    `json:"maximum,omitzero"`
	Minimum param.Opt[int64]    `json:"minimum,omitzero"`
	paramObj
}

func (r LongFieldSchemaParam) MarshalJSON() (data []byte, err error) {
	type shadow LongFieldSchemaParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LongFieldSchemaParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ObjectFieldSchema struct {
	Properties any `json:"properties" api:"required"`
	// Any of "OBJECT".
	Type ObjectFieldSchemaType `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Properties  respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ObjectFieldSchema) RawJSON() string { return r.JSON.raw }
func (r *ObjectFieldSchema) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this ObjectFieldSchema to a ObjectFieldSchemaParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// ObjectFieldSchemaParam.Overrides()
func (r ObjectFieldSchema) ToParam() ObjectFieldSchemaParam {
	return param.Override[ObjectFieldSchemaParam](json.RawMessage(r.RawJSON()))
}

type ObjectFieldSchemaType string

const (
	ObjectFieldSchemaTypeObject ObjectFieldSchemaType = "OBJECT"
)

// The properties Properties, Type are required.
type ObjectFieldSchemaParam struct {
	Properties any `json:"properties,omitzero" api:"required"`
	// Any of "OBJECT".
	Type ObjectFieldSchemaType `json:"type,omitzero" api:"required"`
	paramObj
}

func (r ObjectFieldSchemaParam) MarshalJSON() (data []byte, err error) {
	type shadow ObjectFieldSchemaParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ObjectFieldSchemaParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OutputFieldDefinition struct {
	TypeDefinition FieldTypeDefinition `json:"typeDefinition" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		TypeDefinition respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OutputFieldDefinition) RawJSON() string { return r.JSON.raw }
func (r *OutputFieldDefinition) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this OutputFieldDefinition to a OutputFieldDefinitionParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// OutputFieldDefinitionParam.Overrides()
func (r OutputFieldDefinition) ToParam() OutputFieldDefinitionParam {
	return param.Override[OutputFieldDefinitionParam](json.RawMessage(r.RawJSON()))
}

// The property TypeDefinition is required.
type OutputFieldDefinitionParam struct {
	TypeDefinition FieldTypeDefinitionParam `json:"typeDefinition,omitzero" api:"required"`
	paramObj
}

func (r OutputFieldDefinitionParam) MarshalJSON() (data []byte, err error) {
	type shadow OutputFieldDefinitionParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OutputFieldDefinitionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PublicActionDefinition struct {
	ID                     string                                            `json:"id" api:"required"`
	ActionURL              string                                            `json:"actionUrl" api:"required"`
	Functions              []PublicActionFunctionIdentifier                  `json:"functions" api:"required"`
	InputFields            []PublicInputFieldDefinition                      `json:"inputFields" api:"required"`
	Labels                 map[string]PublicActionLabels                     `json:"labels" api:"required"`
	ObjectTypes            []string                                          `json:"objectTypes" api:"required"`
	Published              bool                                              `json:"published" api:"required"`
	RevisionID             string                                            `json:"revisionId" api:"required"`
	ArchivedAt             int64                                             `json:"archivedAt"`
	ExecutionRules         []PublicExecutionTranslationRule                  `json:"executionRules"`
	InputFieldDependencies []PublicActionDefinitionInputFieldDependencyUnion `json:"inputFieldDependencies"`
	ObjectRequestOptions   PublicObjectRequestOptions                        `json:"objectRequestOptions"`
	OutputFields           []OutputFieldDefinition                           `json:"outputFields"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                     respjson.Field
		ActionURL              respjson.Field
		Functions              respjson.Field
		InputFields            respjson.Field
		Labels                 respjson.Field
		ObjectTypes            respjson.Field
		Published              respjson.Field
		RevisionID             respjson.Field
		ArchivedAt             respjson.Field
		ExecutionRules         respjson.Field
		InputFieldDependencies respjson.Field
		ObjectRequestOptions   respjson.Field
		OutputFields           respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PublicActionDefinition) RawJSON() string { return r.JSON.raw }
func (r *PublicActionDefinition) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PublicActionDefinitionInputFieldDependencyUnion contains all possible properties
// and values from [PublicSingleFieldDependency],
// [PublicConditionalSingleFieldDependency].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type PublicActionDefinitionInputFieldDependencyUnion struct {
	ControllingFieldName string   `json:"controllingFieldName"`
	DependencyType       string   `json:"dependencyType"`
	DependentFieldNames  []string `json:"dependentFieldNames"`
	// This field is from variant [PublicConditionalSingleFieldDependency].
	ControllingFieldValue string `json:"controllingFieldValue"`
	JSON                  struct {
		ControllingFieldName  respjson.Field
		DependencyType        respjson.Field
		DependentFieldNames   respjson.Field
		ControllingFieldValue respjson.Field
		raw                   string
	} `json:"-"`
}

func (u PublicActionDefinitionInputFieldDependencyUnion) AsSingleField() (v PublicSingleFieldDependency) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicActionDefinitionInputFieldDependencyUnion) AsConditionalSingleField() (v PublicConditionalSingleFieldDependency) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u PublicActionDefinitionInputFieldDependencyUnion) RawJSON() string { return u.JSON.raw }

func (r *PublicActionDefinitionInputFieldDependencyUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties ActionURL, Functions, InputFields, Labels, ObjectTypes, Published
// are required.
type PublicActionDefinitionEggParam struct {
	ActionURL              string                                                    `json:"actionUrl" api:"required"`
	Functions              []PublicActionFunctionParam                               `json:"functions,omitzero" api:"required"`
	InputFields            []PublicInputFieldDefinitionParam                         `json:"inputFields,omitzero" api:"required"`
	Labels                 map[string]PublicActionLabelsParam                        `json:"labels,omitzero" api:"required"`
	ObjectTypes            []string                                                  `json:"objectTypes,omitzero" api:"required"`
	Published              bool                                                      `json:"published" api:"required"`
	ArchivedAt             param.Opt[int64]                                          `json:"archivedAt,omitzero"`
	ExecutionRules         []PublicExecutionTranslationRuleParam                     `json:"executionRules,omitzero"`
	InputFieldDependencies []PublicActionDefinitionEggInputFieldDependencyUnionParam `json:"inputFieldDependencies,omitzero"`
	ObjectRequestOptions   PublicObjectRequestOptionsParam                           `json:"objectRequestOptions,omitzero"`
	OutputFields           []OutputFieldDefinitionParam                              `json:"outputFields,omitzero"`
	paramObj
}

func (r PublicActionDefinitionEggParam) MarshalJSON() (data []byte, err error) {
	type shadow PublicActionDefinitionEggParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PublicActionDefinitionEggParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type PublicActionDefinitionEggInputFieldDependencyUnionParam struct {
	OfSingleField            *PublicSingleFieldDependencyParam            `json:",omitzero,inline"`
	OfConditionalSingleField *PublicConditionalSingleFieldDependencyParam `json:",omitzero,inline"`
	paramUnion
}

func (u PublicActionDefinitionEggInputFieldDependencyUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfSingleField, u.OfConditionalSingleField)
}
func (u *PublicActionDefinitionEggInputFieldDependencyUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *PublicActionDefinitionEggInputFieldDependencyUnionParam) asAny() any {
	if !param.IsOmitted(u.OfSingleField) {
		return u.OfSingleField
	} else if !param.IsOmitted(u.OfConditionalSingleField) {
		return u.OfConditionalSingleField
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u PublicActionDefinitionEggInputFieldDependencyUnionParam) GetControllingFieldValue() *string {
	if vt := u.OfConditionalSingleField; vt != nil {
		return &vt.ControllingFieldValue
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u PublicActionDefinitionEggInputFieldDependencyUnionParam) GetControllingFieldName() *string {
	if vt := u.OfSingleField; vt != nil {
		return (*string)(&vt.ControllingFieldName)
	} else if vt := u.OfConditionalSingleField; vt != nil {
		return (*string)(&vt.ControllingFieldName)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u PublicActionDefinitionEggInputFieldDependencyUnionParam) GetDependencyType() *string {
	if vt := u.OfSingleField; vt != nil {
		return (*string)(&vt.DependencyType)
	} else if vt := u.OfConditionalSingleField; vt != nil {
		return (*string)(&vt.DependencyType)
	}
	return nil
}

// Returns a pointer to the underlying variant's DependentFieldNames property, if
// present.
func (u PublicActionDefinitionEggInputFieldDependencyUnionParam) GetDependentFieldNames() []string {
	if vt := u.OfSingleField; vt != nil {
		return vt.DependentFieldNames
	} else if vt := u.OfConditionalSingleField; vt != nil {
		return vt.DependentFieldNames
	}
	return nil
}

type PublicActionDefinitionPatchParam struct {
	ActionURL              param.Opt[string]                                           `json:"actionUrl,omitzero"`
	Published              param.Opt[bool]                                             `json:"published,omitzero"`
	ExecutionRules         []PublicExecutionTranslationRuleParam                       `json:"executionRules,omitzero"`
	InputFieldDependencies []PublicActionDefinitionPatchInputFieldDependencyUnionParam `json:"inputFieldDependencies,omitzero"`
	InputFields            []PublicInputFieldDefinitionParam                           `json:"inputFields,omitzero"`
	Labels                 map[string]PublicActionLabelsParam                          `json:"labels,omitzero"`
	ObjectRequestOptions   PublicObjectRequestOptionsParam                             `json:"objectRequestOptions,omitzero"`
	ObjectTypes            []string                                                    `json:"objectTypes,omitzero"`
	OutputFields           []OutputFieldDefinitionParam                                `json:"outputFields,omitzero"`
	paramObj
}

func (r PublicActionDefinitionPatchParam) MarshalJSON() (data []byte, err error) {
	type shadow PublicActionDefinitionPatchParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PublicActionDefinitionPatchParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type PublicActionDefinitionPatchInputFieldDependencyUnionParam struct {
	OfSingleField            *PublicSingleFieldDependencyParam            `json:",omitzero,inline"`
	OfConditionalSingleField *PublicConditionalSingleFieldDependencyParam `json:",omitzero,inline"`
	paramUnion
}

func (u PublicActionDefinitionPatchInputFieldDependencyUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfSingleField, u.OfConditionalSingleField)
}
func (u *PublicActionDefinitionPatchInputFieldDependencyUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *PublicActionDefinitionPatchInputFieldDependencyUnionParam) asAny() any {
	if !param.IsOmitted(u.OfSingleField) {
		return u.OfSingleField
	} else if !param.IsOmitted(u.OfConditionalSingleField) {
		return u.OfConditionalSingleField
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u PublicActionDefinitionPatchInputFieldDependencyUnionParam) GetControllingFieldValue() *string {
	if vt := u.OfConditionalSingleField; vt != nil {
		return &vt.ControllingFieldValue
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u PublicActionDefinitionPatchInputFieldDependencyUnionParam) GetControllingFieldName() *string {
	if vt := u.OfSingleField; vt != nil {
		return (*string)(&vt.ControllingFieldName)
	} else if vt := u.OfConditionalSingleField; vt != nil {
		return (*string)(&vt.ControllingFieldName)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u PublicActionDefinitionPatchInputFieldDependencyUnionParam) GetDependencyType() *string {
	if vt := u.OfSingleField; vt != nil {
		return (*string)(&vt.DependencyType)
	} else if vt := u.OfConditionalSingleField; vt != nil {
		return (*string)(&vt.DependencyType)
	}
	return nil
}

// Returns a pointer to the underlying variant's DependentFieldNames property, if
// present.
func (u PublicActionDefinitionPatchInputFieldDependencyUnionParam) GetDependentFieldNames() []string {
	if vt := u.OfSingleField; vt != nil {
		return vt.DependentFieldNames
	} else if vt := u.OfConditionalSingleField; vt != nil {
		return vt.DependentFieldNames
	}
	return nil
}

// The property RequiresObject is required.
type PublicActionDefinitionRequiresObjectRequestParam struct {
	RequiresObject bool `json:"requiresObject" api:"required"`
	paramObj
}

func (r PublicActionDefinitionRequiresObjectRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow PublicActionDefinitionRequiresObjectRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PublicActionDefinitionRequiresObjectRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PublicActionDefinitionRequiresObjectResponse struct {
	RequiresObject bool `json:"requiresObject" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		RequiresObject respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PublicActionDefinitionRequiresObjectResponse) RawJSON() string { return r.JSON.raw }
func (r *PublicActionDefinitionRequiresObjectResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PublicActionFunction struct {
	FunctionSource string `json:"functionSource" api:"required"`
	// Any of "POST_ACTION_EXECUTION", "POST_FETCH_OPTIONS", "PRE_ACTION_EXECUTION",
	// "PRE_FETCH_OPTIONS".
	FunctionType PublicActionFunctionFunctionType `json:"functionType" api:"required"`
	ID           string                           `json:"id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FunctionSource respjson.Field
		FunctionType   respjson.Field
		ID             respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PublicActionFunction) RawJSON() string { return r.JSON.raw }
func (r *PublicActionFunction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this PublicActionFunction to a PublicActionFunctionParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// PublicActionFunctionParam.Overrides()
func (r PublicActionFunction) ToParam() PublicActionFunctionParam {
	return param.Override[PublicActionFunctionParam](json.RawMessage(r.RawJSON()))
}

type PublicActionFunctionFunctionType string

const (
	PublicActionFunctionFunctionTypePostActionExecution PublicActionFunctionFunctionType = "POST_ACTION_EXECUTION"
	PublicActionFunctionFunctionTypePostFetchOptions    PublicActionFunctionFunctionType = "POST_FETCH_OPTIONS"
	PublicActionFunctionFunctionTypePreActionExecution  PublicActionFunctionFunctionType = "PRE_ACTION_EXECUTION"
	PublicActionFunctionFunctionTypePreFetchOptions     PublicActionFunctionFunctionType = "PRE_FETCH_OPTIONS"
)

// The properties FunctionSource, FunctionType are required.
type PublicActionFunctionParam struct {
	FunctionSource string `json:"functionSource" api:"required"`
	// Any of "POST_ACTION_EXECUTION", "POST_FETCH_OPTIONS", "PRE_ACTION_EXECUTION",
	// "PRE_FETCH_OPTIONS".
	FunctionType PublicActionFunctionFunctionType `json:"functionType,omitzero" api:"required"`
	ID           param.Opt[string]                `json:"id,omitzero"`
	paramObj
}

func (r PublicActionFunctionParam) MarshalJSON() (data []byte, err error) {
	type shadow PublicActionFunctionParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PublicActionFunctionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PublicActionFunctionIdentifier struct {
	// Any of "POST_ACTION_EXECUTION", "POST_FETCH_OPTIONS", "PRE_ACTION_EXECUTION",
	// "PRE_FETCH_OPTIONS".
	FunctionType PublicActionFunctionIdentifierFunctionType `json:"functionType" api:"required"`
	ID           string                                     `json:"id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FunctionType respjson.Field
		ID           respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PublicActionFunctionIdentifier) RawJSON() string { return r.JSON.raw }
func (r *PublicActionFunctionIdentifier) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PublicActionFunctionIdentifierFunctionType string

const (
	PublicActionFunctionIdentifierFunctionTypePostActionExecution PublicActionFunctionIdentifierFunctionType = "POST_ACTION_EXECUTION"
	PublicActionFunctionIdentifierFunctionTypePostFetchOptions    PublicActionFunctionIdentifierFunctionType = "POST_FETCH_OPTIONS"
	PublicActionFunctionIdentifierFunctionTypePreActionExecution  PublicActionFunctionIdentifierFunctionType = "PRE_ACTION_EXECUTION"
	PublicActionFunctionIdentifierFunctionTypePreFetchOptions     PublicActionFunctionIdentifierFunctionType = "PRE_FETCH_OPTIONS"
)

type PublicActionLabels struct {
	ActionName             string                       `json:"actionName" api:"required"`
	ActionCardContent      string                       `json:"actionCardContent"`
	ActionDescription      string                       `json:"actionDescription"`
	AppDisplayName         string                       `json:"appDisplayName"`
	ExecutionRules         map[string]string            `json:"executionRules"`
	InputFieldDescriptions map[string]string            `json:"inputFieldDescriptions"`
	InputFieldLabels       map[string]string            `json:"inputFieldLabels"`
	InputFieldOptionLabels map[string]map[string]string `json:"inputFieldOptionLabels"`
	OutputFieldLabels      map[string]string            `json:"outputFieldLabels"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ActionName             respjson.Field
		ActionCardContent      respjson.Field
		ActionDescription      respjson.Field
		AppDisplayName         respjson.Field
		ExecutionRules         respjson.Field
		InputFieldDescriptions respjson.Field
		InputFieldLabels       respjson.Field
		InputFieldOptionLabels respjson.Field
		OutputFieldLabels      respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PublicActionLabels) RawJSON() string { return r.JSON.raw }
func (r *PublicActionLabels) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this PublicActionLabels to a PublicActionLabelsParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// PublicActionLabelsParam.Overrides()
func (r PublicActionLabels) ToParam() PublicActionLabelsParam {
	return param.Override[PublicActionLabelsParam](json.RawMessage(r.RawJSON()))
}

// The property ActionName is required.
type PublicActionLabelsParam struct {
	ActionName             string                       `json:"actionName" api:"required"`
	ActionCardContent      param.Opt[string]            `json:"actionCardContent,omitzero"`
	ActionDescription      param.Opt[string]            `json:"actionDescription,omitzero"`
	AppDisplayName         param.Opt[string]            `json:"appDisplayName,omitzero"`
	ExecutionRules         map[string]string            `json:"executionRules,omitzero"`
	InputFieldDescriptions map[string]string            `json:"inputFieldDescriptions,omitzero"`
	InputFieldLabels       map[string]string            `json:"inputFieldLabels,omitzero"`
	InputFieldOptionLabels map[string]map[string]string `json:"inputFieldOptionLabels,omitzero"`
	OutputFieldLabels      map[string]string            `json:"outputFieldLabels,omitzero"`
	paramObj
}

func (r PublicActionLabelsParam) MarshalJSON() (data []byte, err error) {
	type shadow PublicActionLabelsParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PublicActionLabelsParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PublicActionRevision struct {
	ID         string                 `json:"id" api:"required"`
	CreatedAt  time.Time              `json:"createdAt" api:"required" format:"date-time"`
	Definition PublicActionDefinition `json:"definition" api:"required"`
	RevisionID string                 `json:"revisionId" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		Definition  respjson.Field
		RevisionID  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PublicActionRevision) RawJSON() string { return r.JSON.raw }
func (r *PublicActionRevision) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PublicConditionalSingleFieldDependency struct {
	ControllingFieldName  string `json:"controllingFieldName" api:"required"`
	ControllingFieldValue string `json:"controllingFieldValue" api:"required"`
	// Any of "CONDITIONAL_SINGLE_FIELD".
	DependencyType      PublicConditionalSingleFieldDependencyDependencyType `json:"dependencyType" api:"required"`
	DependentFieldNames []string                                             `json:"dependentFieldNames" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ControllingFieldName  respjson.Field
		ControllingFieldValue respjson.Field
		DependencyType        respjson.Field
		DependentFieldNames   respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PublicConditionalSingleFieldDependency) RawJSON() string { return r.JSON.raw }
func (r *PublicConditionalSingleFieldDependency) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this PublicConditionalSingleFieldDependency to a
// PublicConditionalSingleFieldDependencyParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// PublicConditionalSingleFieldDependencyParam.Overrides()
func (r PublicConditionalSingleFieldDependency) ToParam() PublicConditionalSingleFieldDependencyParam {
	return param.Override[PublicConditionalSingleFieldDependencyParam](json.RawMessage(r.RawJSON()))
}

type PublicConditionalSingleFieldDependencyDependencyType string

const (
	PublicConditionalSingleFieldDependencyDependencyTypeConditionalSingleField PublicConditionalSingleFieldDependencyDependencyType = "CONDITIONAL_SINGLE_FIELD"
)

// The properties ControllingFieldName, ControllingFieldValue, DependencyType,
// DependentFieldNames are required.
type PublicConditionalSingleFieldDependencyParam struct {
	ControllingFieldName  string `json:"controllingFieldName" api:"required"`
	ControllingFieldValue string `json:"controllingFieldValue" api:"required"`
	// Any of "CONDITIONAL_SINGLE_FIELD".
	DependencyType      PublicConditionalSingleFieldDependencyDependencyType `json:"dependencyType,omitzero" api:"required"`
	DependentFieldNames []string                                             `json:"dependentFieldNames,omitzero" api:"required"`
	paramObj
}

func (r PublicConditionalSingleFieldDependencyParam) MarshalJSON() (data []byte, err error) {
	type shadow PublicConditionalSingleFieldDependencyParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PublicConditionalSingleFieldDependencyParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PublicExecutionTranslationRule struct {
	Conditions map[string]any `json:"conditions" api:"required"`
	LabelName  string         `json:"labelName" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Conditions  respjson.Field
		LabelName   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PublicExecutionTranslationRule) RawJSON() string { return r.JSON.raw }
func (r *PublicExecutionTranslationRule) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this PublicExecutionTranslationRule to a
// PublicExecutionTranslationRuleParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// PublicExecutionTranslationRuleParam.Overrides()
func (r PublicExecutionTranslationRule) ToParam() PublicExecutionTranslationRuleParam {
	return param.Override[PublicExecutionTranslationRuleParam](json.RawMessage(r.RawJSON()))
}

// The properties Conditions, LabelName are required.
type PublicExecutionTranslationRuleParam struct {
	Conditions map[string]any `json:"conditions,omitzero" api:"required"`
	LabelName  string         `json:"labelName" api:"required"`
	paramObj
}

func (r PublicExecutionTranslationRuleParam) MarshalJSON() (data []byte, err error) {
	type shadow PublicExecutionTranslationRuleParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PublicExecutionTranslationRuleParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PublicFieldTypeDefinition struct {
	Name    string         `json:"name" api:"required"`
	Options []PublicOption `json:"options" api:"required"`
	// Any of "bool", "date", "datetime", "enumeration", "json", "number",
	// "object_coordinates", "phone_number", "string".
	Type        PublicFieldTypeDefinitionType `json:"type" api:"required"`
	Description string                        `json:"description"`
	// Any of "booleancheckbox", "calculation_equation", "checkbox", "date", "file",
	// "html", "number", "phonenumber", "radio", "select", "text", "textarea".
	FieldType  PublicFieldTypeDefinitionFieldType `json:"fieldType"`
	HelpText   string                             `json:"helpText"`
	Label      string                             `json:"label"`
	OptionsURL string                             `json:"optionsUrl"`
	// Any of "OWNER".
	ReferencedObjectType PublicFieldTypeDefinitionReferencedObjectType `json:"referencedObjectType"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name                 respjson.Field
		Options              respjson.Field
		Type                 respjson.Field
		Description          respjson.Field
		FieldType            respjson.Field
		HelpText             respjson.Field
		Label                respjson.Field
		OptionsURL           respjson.Field
		ReferencedObjectType respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PublicFieldTypeDefinition) RawJSON() string { return r.JSON.raw }
func (r *PublicFieldTypeDefinition) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this PublicFieldTypeDefinition to a
// PublicFieldTypeDefinitionParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// PublicFieldTypeDefinitionParam.Overrides()
func (r PublicFieldTypeDefinition) ToParam() PublicFieldTypeDefinitionParam {
	return param.Override[PublicFieldTypeDefinitionParam](json.RawMessage(r.RawJSON()))
}

type PublicFieldTypeDefinitionType string

const (
	PublicFieldTypeDefinitionTypeBool              PublicFieldTypeDefinitionType = "bool"
	PublicFieldTypeDefinitionTypeDate              PublicFieldTypeDefinitionType = "date"
	PublicFieldTypeDefinitionTypeDatetime          PublicFieldTypeDefinitionType = "datetime"
	PublicFieldTypeDefinitionTypeEnumeration       PublicFieldTypeDefinitionType = "enumeration"
	PublicFieldTypeDefinitionTypeJson              PublicFieldTypeDefinitionType = "json"
	PublicFieldTypeDefinitionTypeNumber            PublicFieldTypeDefinitionType = "number"
	PublicFieldTypeDefinitionTypeObjectCoordinates PublicFieldTypeDefinitionType = "object_coordinates"
	PublicFieldTypeDefinitionTypePhoneNumber       PublicFieldTypeDefinitionType = "phone_number"
	PublicFieldTypeDefinitionTypeString            PublicFieldTypeDefinitionType = "string"
)

type PublicFieldTypeDefinitionFieldType string

const (
	PublicFieldTypeDefinitionFieldTypeBooleancheckbox     PublicFieldTypeDefinitionFieldType = "booleancheckbox"
	PublicFieldTypeDefinitionFieldTypeCalculationEquation PublicFieldTypeDefinitionFieldType = "calculation_equation"
	PublicFieldTypeDefinitionFieldTypeCheckbox            PublicFieldTypeDefinitionFieldType = "checkbox"
	PublicFieldTypeDefinitionFieldTypeDate                PublicFieldTypeDefinitionFieldType = "date"
	PublicFieldTypeDefinitionFieldTypeFile                PublicFieldTypeDefinitionFieldType = "file"
	PublicFieldTypeDefinitionFieldTypeHTML                PublicFieldTypeDefinitionFieldType = "html"
	PublicFieldTypeDefinitionFieldTypeNumber              PublicFieldTypeDefinitionFieldType = "number"
	PublicFieldTypeDefinitionFieldTypePhonenumber         PublicFieldTypeDefinitionFieldType = "phonenumber"
	PublicFieldTypeDefinitionFieldTypeRadio               PublicFieldTypeDefinitionFieldType = "radio"
	PublicFieldTypeDefinitionFieldTypeSelect              PublicFieldTypeDefinitionFieldType = "select"
	PublicFieldTypeDefinitionFieldTypeText                PublicFieldTypeDefinitionFieldType = "text"
	PublicFieldTypeDefinitionFieldTypeTextarea            PublicFieldTypeDefinitionFieldType = "textarea"
)

type PublicFieldTypeDefinitionReferencedObjectType string

const (
	PublicFieldTypeDefinitionReferencedObjectTypeOwner PublicFieldTypeDefinitionReferencedObjectType = "OWNER"
)

// The properties Name, Options, Type are required.
type PublicFieldTypeDefinitionParam struct {
	Name    string              `json:"name" api:"required"`
	Options []PublicOptionParam `json:"options,omitzero" api:"required"`
	// Any of "bool", "date", "datetime", "enumeration", "json", "number",
	// "object_coordinates", "phone_number", "string".
	Type        PublicFieldTypeDefinitionType `json:"type,omitzero" api:"required"`
	Description param.Opt[string]             `json:"description,omitzero"`
	HelpText    param.Opt[string]             `json:"helpText,omitzero"`
	Label       param.Opt[string]             `json:"label,omitzero"`
	OptionsURL  param.Opt[string]             `json:"optionsUrl,omitzero"`
	// Any of "booleancheckbox", "calculation_equation", "checkbox", "date", "file",
	// "html", "number", "phonenumber", "radio", "select", "text", "textarea".
	FieldType PublicFieldTypeDefinitionFieldType `json:"fieldType,omitzero"`
	// Any of "OWNER".
	ReferencedObjectType PublicFieldTypeDefinitionReferencedObjectType `json:"referencedObjectType,omitzero"`
	paramObj
}

func (r PublicFieldTypeDefinitionParam) MarshalJSON() (data []byte, err error) {
	type shadow PublicFieldTypeDefinitionParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PublicFieldTypeDefinitionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PublicInputFieldDefinition struct {
	IsRequired     bool                      `json:"isRequired" api:"required"`
	TypeDefinition PublicFieldTypeDefinition `json:"typeDefinition" api:"required"`
	// Any of "STATIC_VALUE", "OBJECT_PROPERTY".
	SupportedValueTypes []string `json:"supportedValueTypes"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		IsRequired          respjson.Field
		TypeDefinition      respjson.Field
		SupportedValueTypes respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PublicInputFieldDefinition) RawJSON() string { return r.JSON.raw }
func (r *PublicInputFieldDefinition) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this PublicInputFieldDefinition to a
// PublicInputFieldDefinitionParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// PublicInputFieldDefinitionParam.Overrides()
func (r PublicInputFieldDefinition) ToParam() PublicInputFieldDefinitionParam {
	return param.Override[PublicInputFieldDefinitionParam](json.RawMessage(r.RawJSON()))
}

// The properties IsRequired, TypeDefinition are required.
type PublicInputFieldDefinitionParam struct {
	IsRequired     bool                           `json:"isRequired" api:"required"`
	TypeDefinition PublicFieldTypeDefinitionParam `json:"typeDefinition,omitzero" api:"required"`
	// Any of "STATIC_VALUE", "OBJECT_PROPERTY".
	SupportedValueTypes []string `json:"supportedValueTypes,omitzero"`
	paramObj
}

func (r PublicInputFieldDefinitionParam) MarshalJSON() (data []byte, err error) {
	type shadow PublicInputFieldDefinitionParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PublicInputFieldDefinitionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PublicObjectRequestOptions struct {
	Properties []string `json:"properties" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Properties  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PublicObjectRequestOptions) RawJSON() string { return r.JSON.raw }
func (r *PublicObjectRequestOptions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this PublicObjectRequestOptions to a
// PublicObjectRequestOptionsParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// PublicObjectRequestOptionsParam.Overrides()
func (r PublicObjectRequestOptions) ToParam() PublicObjectRequestOptionsParam {
	return param.Override[PublicObjectRequestOptionsParam](json.RawMessage(r.RawJSON()))
}

// The property Properties is required.
type PublicObjectRequestOptionsParam struct {
	Properties []string `json:"properties,omitzero" api:"required"`
	paramObj
}

func (r PublicObjectRequestOptionsParam) MarshalJSON() (data []byte, err error) {
	type shadow PublicObjectRequestOptionsParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PublicObjectRequestOptionsParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PublicOption struct {
	Label        string `json:"label" api:"required"`
	Value        string `json:"value" api:"required"`
	Description  string `json:"description"`
	DisplayOrder int64  `json:"displayOrder"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Label        respjson.Field
		Value        respjson.Field
		Description  respjson.Field
		DisplayOrder respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PublicOption) RawJSON() string { return r.JSON.raw }
func (r *PublicOption) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this PublicOption to a PublicOptionParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// PublicOptionParam.Overrides()
func (r PublicOption) ToParam() PublicOptionParam {
	return param.Override[PublicOptionParam](json.RawMessage(r.RawJSON()))
}

// The properties Label, Value are required.
type PublicOptionParam struct {
	Label        string            `json:"label" api:"required"`
	Value        string            `json:"value" api:"required"`
	Description  param.Opt[string] `json:"description,omitzero"`
	DisplayOrder param.Opt[int64]  `json:"displayOrder,omitzero"`
	paramObj
}

func (r PublicOptionParam) MarshalJSON() (data []byte, err error) {
	type shadow PublicOptionParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PublicOptionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PublicSingleFieldDependency struct {
	ControllingFieldName string `json:"controllingFieldName" api:"required"`
	// Any of "SINGLE_FIELD".
	DependencyType      PublicSingleFieldDependencyDependencyType `json:"dependencyType" api:"required"`
	DependentFieldNames []string                                  `json:"dependentFieldNames" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ControllingFieldName respjson.Field
		DependencyType       respjson.Field
		DependentFieldNames  respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PublicSingleFieldDependency) RawJSON() string { return r.JSON.raw }
func (r *PublicSingleFieldDependency) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this PublicSingleFieldDependency to a
// PublicSingleFieldDependencyParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// PublicSingleFieldDependencyParam.Overrides()
func (r PublicSingleFieldDependency) ToParam() PublicSingleFieldDependencyParam {
	return param.Override[PublicSingleFieldDependencyParam](json.RawMessage(r.RawJSON()))
}

type PublicSingleFieldDependencyDependencyType string

const (
	PublicSingleFieldDependencyDependencyTypeSingleField PublicSingleFieldDependencyDependencyType = "SINGLE_FIELD"
)

// The properties ControllingFieldName, DependencyType, DependentFieldNames are
// required.
type PublicSingleFieldDependencyParam struct {
	ControllingFieldName string `json:"controllingFieldName" api:"required"`
	// Any of "SINGLE_FIELD".
	DependencyType      PublicSingleFieldDependencyDependencyType `json:"dependencyType,omitzero" api:"required"`
	DependentFieldNames []string                                  `json:"dependentFieldNames,omitzero" api:"required"`
	paramObj
}

func (r PublicSingleFieldDependencyParam) MarshalJSON() (data []byte, err error) {
	type shadow PublicSingleFieldDependencyParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PublicSingleFieldDependencyParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties ChirpAIContextObject, Source are required.
type StandaloneRequestContextParam struct {
	ChirpAIContextObject ChirpAIContextObjectParam `json:"chirpAiContextObject,omitzero" api:"required"`
	// Any of "STANDALONE".
	Source       StandaloneRequestContextSource `json:"source,omitzero" api:"required"`
	TrajectoryID param.Opt[string]              `json:"trajectoryId,omitzero"`
	paramObj
}

func (r StandaloneRequestContextParam) MarshalJSON() (data []byte, err error) {
	type shadow StandaloneRequestContextParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *StandaloneRequestContextParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type StandaloneRequestContextSource string

const (
	StandaloneRequestContextSourceStandalone StandaloneRequestContextSource = "STANDALONE"
)

type StringFieldSchema struct {
	// Any of "STRING".
	Type StringFieldSchemaType `json:"type" api:"required"`
	// Any of "DATE", "DATE_TIME", "OBJECT_COORDINATE", "TIME", "URI".
	Format StringFieldSchemaFormat `json:"format"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		Format      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r StringFieldSchema) RawJSON() string { return r.JSON.raw }
func (r *StringFieldSchema) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this StringFieldSchema to a StringFieldSchemaParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// StringFieldSchemaParam.Overrides()
func (r StringFieldSchema) ToParam() StringFieldSchemaParam {
	return param.Override[StringFieldSchemaParam](json.RawMessage(r.RawJSON()))
}

type StringFieldSchemaType string

const (
	StringFieldSchemaTypeString StringFieldSchemaType = "STRING"
)

type StringFieldSchemaFormat string

const (
	StringFieldSchemaFormatDate             StringFieldSchemaFormat = "DATE"
	StringFieldSchemaFormatDateTime         StringFieldSchemaFormat = "DATE_TIME"
	StringFieldSchemaFormatObjectCoordinate StringFieldSchemaFormat = "OBJECT_COORDINATE"
	StringFieldSchemaFormatTime             StringFieldSchemaFormat = "TIME"
	StringFieldSchemaFormatUri              StringFieldSchemaFormat = "URI"
)

// The property Type is required.
type StringFieldSchemaParam struct {
	// Any of "STRING".
	Type StringFieldSchemaType `json:"type,omitzero" api:"required"`
	// Any of "DATE", "DATE_TIME", "OBJECT_COORDINATE", "TIME", "URI".
	Format StringFieldSchemaFormat `json:"format,omitzero"`
	paramObj
}

func (r StringFieldSchemaParam) MarshalJSON() (data []byte, err error) {
	type shadow StringFieldSchemaParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *StringFieldSchemaParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Source is required.
type TestRequestContextParam struct {
	// Any of "TEST".
	Source TestRequestContextSource `json:"source,omitzero" api:"required"`
	paramObj
}

func (r TestRequestContextParam) MarshalJSON() (data []byte, err error) {
	type shadow TestRequestContextParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TestRequestContextParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TestRequestContextSource string

const (
	TestRequestContextSourceTest TestRequestContextSource = "TEST"
)

// The properties Source, WorkflowID are required.
type WorkflowsRequestContextParam struct {
	// Any of "WORKFLOWS".
	Source                         WorkflowsRequestContextSource       `json:"source,omitzero" api:"required"`
	WorkflowID                     int64                               `json:"workflowId" api:"required"`
	ActionID                       param.Opt[int64]                    `json:"actionId,omitzero"`
	ActionExecutionIndexIdentifier ActionExecutionIndexIdentifierParam `json:"actionExecutionIndexIdentifier,omitzero"`
	paramObj
}

func (r WorkflowsRequestContextParam) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowsRequestContextParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowsRequestContextParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkflowsRequestContextSource string

const (
	WorkflowsRequestContextSourceWorkflows WorkflowsRequestContextSource = "WORKFLOWS"
)

type ActionNewParams struct {
	PublicActionDefinitionEgg PublicActionDefinitionEggParam
	paramObj
}

func (r ActionNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.PublicActionDefinitionEgg)
}
func (r *ActionNewParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.PublicActionDefinitionEgg)
}

type ActionUpdateParams struct {
	AppID                       int64 `path:"appId" api:"required" json:"-"`
	PublicActionDefinitionPatch PublicActionDefinitionPatchParam
	paramObj
}

func (r ActionUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.PublicActionDefinitionPatch)
}
func (r *ActionUpdateParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.PublicActionDefinitionPatch)
}

type ActionListParams struct {
	AppID int64 `path:"appId" api:"required" json:"-"`
	// The paging cursor token of the last successfully read resource will be returned
	// as the `paging.next.after` JSON property of a paged response containing more
	// results.
	After param.Opt[string] `query:"after,omitzero" json:"-"`
	// The maximum number of results to display per page.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ActionListParams]'s query parameters as `url.Values`.
func (r ActionListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ActionDeleteParams struct {
	AppID        int64  `path:"appId" api:"required" json:"-"`
	DefinitionID string `path:"definitionId" api:"required" json:"-"`
	// Any of "POST_ACTION_EXECUTION", "POST_FETCH_OPTIONS", "PRE_ACTION_EXECUTION",
	// "PRE_FETCH_OPTIONS".
	FunctionType ActionDeleteParamsFunctionType `path:"functionType,omitzero" api:"required" json:"-"`
	paramObj
}

type ActionDeleteParamsFunctionType string

const (
	ActionDeleteParamsFunctionTypePostActionExecution ActionDeleteParamsFunctionType = "POST_ACTION_EXECUTION"
	ActionDeleteParamsFunctionTypePostFetchOptions    ActionDeleteParamsFunctionType = "POST_FETCH_OPTIONS"
	ActionDeleteParamsFunctionTypePreActionExecution  ActionDeleteParamsFunctionType = "PRE_ACTION_EXECUTION"
	ActionDeleteParamsFunctionTypePreFetchOptions     ActionDeleteParamsFunctionType = "PRE_FETCH_OPTIONS"
)

type ActionCompleteParams struct {
	CallbackCompletionRequest CallbackCompletionRequestParam
	paramObj
}

func (r ActionCompleteParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.CallbackCompletionRequest)
}
func (r *ActionCompleteParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.CallbackCompletionRequest)
}

type ActionCompleteBatchParams struct {
	BatchInputCallbackCompletionBatchRequest BatchInputCallbackCompletionBatchRequestParam
	paramObj
}

func (r ActionCompleteBatchParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.BatchInputCallbackCompletionBatchRequest)
}
func (r *ActionCompleteBatchParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.BatchInputCallbackCompletionBatchRequest)
}

type ActionNewOrReplaceParams struct {
	AppID        int64  `path:"appId" api:"required" json:"-"`
	DefinitionID string `path:"definitionId" api:"required" json:"-"`
	// Any of "POST_ACTION_EXECUTION", "POST_FETCH_OPTIONS", "PRE_ACTION_EXECUTION",
	// "PRE_FETCH_OPTIONS".
	FunctionType ActionNewOrReplaceParamsFunctionType `path:"functionType,omitzero" api:"required" json:"-"`
	Body         string
	paramObj
}

func (r ActionNewOrReplaceParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.Body)
}
func (r *ActionNewOrReplaceParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.Body)
}

type ActionNewOrReplaceParamsFunctionType string

const (
	ActionNewOrReplaceParamsFunctionTypePostActionExecution ActionNewOrReplaceParamsFunctionType = "POST_ACTION_EXECUTION"
	ActionNewOrReplaceParamsFunctionTypePostFetchOptions    ActionNewOrReplaceParamsFunctionType = "POST_FETCH_OPTIONS"
	ActionNewOrReplaceParamsFunctionTypePreActionExecution  ActionNewOrReplaceParamsFunctionType = "PRE_ACTION_EXECUTION"
	ActionNewOrReplaceParamsFunctionTypePreFetchOptions     ActionNewOrReplaceParamsFunctionType = "PRE_FETCH_OPTIONS"
)

type ActionNewOrReplaceByFunctionTypeParams struct {
	AppID        int64  `path:"appId" api:"required" json:"-"`
	DefinitionID string `path:"definitionId" api:"required" json:"-"`
	Body         string
	paramObj
}

func (r ActionNewOrReplaceByFunctionTypeParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.Body)
}
func (r *ActionNewOrReplaceByFunctionTypeParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.Body)
}

type ActionNewOrReplaceByFunctionTypeParamsFunctionType string

const (
	ActionNewOrReplaceByFunctionTypeParamsFunctionTypePostActionExecution ActionNewOrReplaceByFunctionTypeParamsFunctionType = "POST_ACTION_EXECUTION"
	ActionNewOrReplaceByFunctionTypeParamsFunctionTypePostFetchOptions    ActionNewOrReplaceByFunctionTypeParamsFunctionType = "POST_FETCH_OPTIONS"
	ActionNewOrReplaceByFunctionTypeParamsFunctionTypePreActionExecution  ActionNewOrReplaceByFunctionTypeParamsFunctionType = "PRE_ACTION_EXECUTION"
	ActionNewOrReplaceByFunctionTypeParamsFunctionTypePreFetchOptions     ActionNewOrReplaceByFunctionTypeParamsFunctionType = "PRE_FETCH_OPTIONS"
)

type ActionNewRequiresObjectParams struct {
	AppID                                       int64 `path:"appId" api:"required" json:"-"`
	PublicActionDefinitionRequiresObjectRequest PublicActionDefinitionRequiresObjectRequestParam
	paramObj
}

func (r ActionNewRequiresObjectParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.PublicActionDefinitionRequiresObjectRequest)
}
func (r *ActionNewRequiresObjectParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.PublicActionDefinitionRequiresObjectRequest)
}

type ActionDeleteByFunctionTypeParams struct {
	AppID        int64  `path:"appId" api:"required" json:"-"`
	DefinitionID string `path:"definitionId" api:"required" json:"-"`
	paramObj
}

type ActionDeleteByFunctionTypeParamsFunctionType string

const (
	ActionDeleteByFunctionTypeParamsFunctionTypePostActionExecution ActionDeleteByFunctionTypeParamsFunctionType = "POST_ACTION_EXECUTION"
	ActionDeleteByFunctionTypeParamsFunctionTypePostFetchOptions    ActionDeleteByFunctionTypeParamsFunctionType = "POST_FETCH_OPTIONS"
	ActionDeleteByFunctionTypeParamsFunctionTypePreActionExecution  ActionDeleteByFunctionTypeParamsFunctionType = "PRE_ACTION_EXECUTION"
	ActionDeleteByFunctionTypeParamsFunctionTypePreFetchOptions     ActionDeleteByFunctionTypeParamsFunctionType = "PRE_FETCH_OPTIONS"
)

type ActionGetParams struct {
	AppID        int64  `path:"appId" api:"required" json:"-"`
	DefinitionID string `path:"definitionId" api:"required" json:"-"`
	paramObj
}

type ActionGetByFunctionTypeParams struct {
	AppID        int64  `path:"appId" api:"required" json:"-"`
	DefinitionID string `path:"definitionId" api:"required" json:"-"`
	paramObj
}

type ActionGetByFunctionTypeParamsFunctionType string

const (
	ActionGetByFunctionTypeParamsFunctionTypePostActionExecution ActionGetByFunctionTypeParamsFunctionType = "POST_ACTION_EXECUTION"
	ActionGetByFunctionTypeParamsFunctionTypePostFetchOptions    ActionGetByFunctionTypeParamsFunctionType = "POST_FETCH_OPTIONS"
	ActionGetByFunctionTypeParamsFunctionTypePreActionExecution  ActionGetByFunctionTypeParamsFunctionType = "PRE_ACTION_EXECUTION"
	ActionGetByFunctionTypeParamsFunctionTypePreFetchOptions     ActionGetByFunctionTypeParamsFunctionType = "PRE_FETCH_OPTIONS"
)

type ActionGetRequiresObjectParams struct {
	AppID int64 `path:"appId" api:"required" json:"-"`
	paramObj
}
