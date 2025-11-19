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

	"github.com/stainless-sdks/hubspot-sdk-go/internal/apijson"
	"github.com/stainless-sdks/hubspot-sdk-go/internal/apiquery"
	shimjson "github.com/stainless-sdks/hubspot-sdk-go/internal/encoding/json"
	"github.com/stainless-sdks/hubspot-sdk-go/internal/requestconfig"
	"github.com/stainless-sdks/hubspot-sdk-go/marketing"
	"github.com/stainless-sdks/hubspot-sdk-go/option"
	"github.com/stainless-sdks/hubspot-sdk-go/packages/pagination"
	"github.com/stainless-sdks/hubspot-sdk-go/packages/param"
	"github.com/stainless-sdks/hubspot-sdk-go/packages/respjson"
	"github.com/stainless-sdks/hubspot-sdk-go/shared"
)

// WorkflowService contains methods and other services that help with interacting
// with the Hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWorkflowService] method instead.
type WorkflowService struct {
	Options []option.RequestOption
}

// NewWorkflowService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewWorkflowService(opts ...option.RequestOption) (r WorkflowService) {
	r = WorkflowService{}
	r.Options = opts
	return
}

// Create a new workflow.
func (r *WorkflowService) New(ctx context.Context, body WorkflowNewParams, opts ...option.RequestOption) (res *APIFlowUnion, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "automation/v4/flows"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Update a workflow by ID.
func (r *WorkflowService) Update(ctx context.Context, flowID string, body WorkflowUpdateParams, opts ...option.RequestOption) (res *APIFlowUnion, err error) {
	opts = slices.Concat(r.Options, opts)
	if flowID == "" {
		err = errors.New("missing required flowId parameter")
		return
	}
	path := fmt.Sprintf("automation/v4/flows/%s", flowID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Retrieve all workflows from an account.
func (r *WorkflowService) List(ctx context.Context, query WorkflowListParams, opts ...option.RequestOption) (res *pagination.Page[APIFlowListing], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "automation/v4/flows"
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

// Retrieve all workflows from an account.
func (r *WorkflowService) ListAutoPaging(ctx context.Context, query WorkflowListParams, opts ...option.RequestOption) *pagination.PageAutoPager[APIFlowListing] {
	return pagination.NewPageAutoPager(r.List(ctx, query, opts...))
}

// Fully delete a workflow by ID. Deleted workflows cannot be restored via the API.
// If you need to restore an accidentally deleted flow, you'll need to contact
// support.
func (r *WorkflowService) Delete(ctx context.Context, flowID int64, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := fmt.Sprintf("automation/v4/flows/%v", flowID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Retrieve a batch of workflows by ID.
func (r *WorkflowService) BatchGet(ctx context.Context, body WorkflowBatchGetParams, opts ...option.RequestOption) (res *BatchResponseAPIFlow, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "automation/v4/flows/batch/read"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieve the IDs of v3 workflows that have been migrated to the v4 API.
func (r *WorkflowService) BatchGetIDMappings(ctx context.Context, body WorkflowBatchGetIDMappingsParams, opts ...option.RequestOption) (res *BatchResponseFlowIDWorkflowIDMappingResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "automation/v4/workflow-id-mappings/batch/read"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieve all details for a specific workflow by ID.
func (r *WorkflowService) Get(ctx context.Context, flowID string, opts ...option.RequestOption) (res *APIFlowUnion, err error) {
	opts = slices.Concat(r.Options, opts)
	if flowID == "" {
		err = errors.New("missing required flowId parameter")
		return
	}
	path := fmt.Sprintf("automation/v4/flows/%s", flowID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Retrieve emails sent by a workflow by ID.
func (r *WorkflowService) ListEmailCampaigns(ctx context.Context, query WorkflowListEmailCampaignsParams, opts ...option.RequestOption) (res *CollectionResponseAPIFlowEmailCampaign, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "automation/v4/flows/email-campaigns"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type APIAbTestBranchAction struct {
	// The ID for this action.
	ActionID     string          `json:"actionId,required"`
	TestBranches []APIConnection `json:"testBranches,required"`
	// The type of action this is, can be: "STATIC_BRANCH", "LIST_BRANCH",
	// "AB_TEST_BRANCH", "CUSTOM_CODE", "WEBHOOK", or "SINGLE_CONNECTION"
	//
	// Any of "AB_TEST_BRANCH".
	Type APIAbTestBranchActionType `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ActionID     respjson.Field
		TestBranches respjson.Field
		Type         respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r APIAbTestBranchAction) RawJSON() string { return r.JSON.raw }
func (r *APIAbTestBranchAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this APIAbTestBranchAction to a APIAbTestBranchActionParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// APIAbTestBranchActionParam.Overrides()
func (r APIAbTestBranchAction) ToParam() APIAbTestBranchActionParam {
	return param.Override[APIAbTestBranchActionParam](json.RawMessage(r.RawJSON()))
}

// The type of action this is, can be: "STATIC_BRANCH", "LIST_BRANCH",
// "AB_TEST_BRANCH", "CUSTOM_CODE", "WEBHOOK", or "SINGLE_CONNECTION"
type APIAbTestBranchActionType string

const (
	APIAbTestBranchActionTypeAbTestBranch APIAbTestBranchActionType = "AB_TEST_BRANCH"
)

// The properties ActionID, TestBranches, Type are required.
type APIAbTestBranchActionParam struct {
	// The ID for this action.
	ActionID     string               `json:"actionId,required"`
	TestBranches []APIConnectionParam `json:"testBranches,omitzero,required"`
	// The type of action this is, can be: "STATIC_BRANCH", "LIST_BRANCH",
	// "AB_TEST_BRANCH", "CUSTOM_CODE", "WEBHOOK", or "SINGLE_CONNECTION"
	//
	// Any of "AB_TEST_BRANCH".
	Type APIAbTestBranchActionType `json:"type,omitzero,required"`
	paramObj
}

func (r APIAbTestBranchActionParam) MarshalJSON() (data []byte, err error) {
	type shadow APIAbTestBranchActionParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *APIAbTestBranchActionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type APIActionDataValue struct {
	// Which action to pull data from.
	ActionID string `json:"actionId,required"`
	// The output field name for that action
	DataKey string `json:"dataKey,required"`
	// This is the type of input value. This can be one of: "FIELD_DATA",
	// "OBJECT_PROPERTY", "STATIC_VALUE", "RELATIVE_DATETIME", "TIMESTAMP",
	// "INCREMENT", "FETCHED_OBJECT_PROPERTY", "APPEND_OBJECT_PROPERTY",
	// "STATIC_APPEND_VALUE", "ENROLLMENT_EVENT_PROPERTY"
	//
	// Any of "FIELD_DATA".
	Type APIActionDataValueType `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ActionID    respjson.Field
		DataKey     respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r APIActionDataValue) RawJSON() string { return r.JSON.raw }
func (r *APIActionDataValue) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this APIActionDataValue to a APIActionDataValueParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// APIActionDataValueParam.Overrides()
func (r APIActionDataValue) ToParam() APIActionDataValueParam {
	return param.Override[APIActionDataValueParam](json.RawMessage(r.RawJSON()))
}

// This is the type of input value. This can be one of: "FIELD_DATA",
// "OBJECT_PROPERTY", "STATIC_VALUE", "RELATIVE_DATETIME", "TIMESTAMP",
// "INCREMENT", "FETCHED_OBJECT_PROPERTY", "APPEND_OBJECT_PROPERTY",
// "STATIC_APPEND_VALUE", "ENROLLMENT_EVENT_PROPERTY"
type APIActionDataValueType string

const (
	APIActionDataValueTypeFieldData APIActionDataValueType = "FIELD_DATA"
)

// The properties ActionID, DataKey, Type are required.
type APIActionDataValueParam struct {
	// Which action to pull data from.
	ActionID string `json:"actionId,required"`
	// The output field name for that action
	DataKey string `json:"dataKey,required"`
	// This is the type of input value. This can be one of: "FIELD_DATA",
	// "OBJECT_PROPERTY", "STATIC_VALUE", "RELATIVE_DATETIME", "TIMESTAMP",
	// "INCREMENT", "FETCHED_OBJECT_PROPERTY", "APPEND_OBJECT_PROPERTY",
	// "STATIC_APPEND_VALUE", "ENROLLMENT_EVENT_PROPERTY"
	//
	// Any of "FIELD_DATA".
	Type APIActionDataValueType `json:"type,omitzero,required"`
	paramObj
}

func (r APIActionDataValueParam) MarshalJSON() (data []byte, err error) {
	type shadow APIActionDataValueParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *APIActionDataValueParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type APIAppendObjectPropertyValue struct {
	// The name of the property to append data from
	AppendPropertyName string `json:"appendPropertyName,required"`
	// This is the type of input value. This can be one of: "FIELD_DATA",
	// "OBJECT_PROPERTY", "STATIC_VALUE", "RELATIVE_DATETIME", "TIMESTAMP",
	// "INCREMENT", "FETCHED_OBJECT_PROPERTY", "APPEND_OBJECT_PROPERTY",
	// "STATIC_APPEND_VALUE", "ENROLLMENT_EVENT_PROPERTY"
	//
	// Any of "APPEND_OBJECT_PROPERTY".
	Type APIAppendObjectPropertyValueType `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AppendPropertyName respjson.Field
		Type               respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r APIAppendObjectPropertyValue) RawJSON() string { return r.JSON.raw }
func (r *APIAppendObjectPropertyValue) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this APIAppendObjectPropertyValue to a
// APIAppendObjectPropertyValueParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// APIAppendObjectPropertyValueParam.Overrides()
func (r APIAppendObjectPropertyValue) ToParam() APIAppendObjectPropertyValueParam {
	return param.Override[APIAppendObjectPropertyValueParam](json.RawMessage(r.RawJSON()))
}

// This is the type of input value. This can be one of: "FIELD_DATA",
// "OBJECT_PROPERTY", "STATIC_VALUE", "RELATIVE_DATETIME", "TIMESTAMP",
// "INCREMENT", "FETCHED_OBJECT_PROPERTY", "APPEND_OBJECT_PROPERTY",
// "STATIC_APPEND_VALUE", "ENROLLMENT_EVENT_PROPERTY"
type APIAppendObjectPropertyValueType string

const (
	APIAppendObjectPropertyValueTypeAppendObjectProperty APIAppendObjectPropertyValueType = "APPEND_OBJECT_PROPERTY"
)

// The properties AppendPropertyName, Type are required.
type APIAppendObjectPropertyValueParam struct {
	// The name of the property to append data from
	AppendPropertyName string `json:"appendPropertyName,required"`
	// This is the type of input value. This can be one of: "FIELD_DATA",
	// "OBJECT_PROPERTY", "STATIC_VALUE", "RELATIVE_DATETIME", "TIMESTAMP",
	// "INCREMENT", "FETCHED_OBJECT_PROPERTY", "APPEND_OBJECT_PROPERTY",
	// "STATIC_APPEND_VALUE", "ENROLLMENT_EVENT_PROPERTY"
	//
	// Any of "APPEND_OBJECT_PROPERTY".
	Type APIAppendObjectPropertyValueType `json:"type,omitzero,required"`
	paramObj
}

func (r APIAppendObjectPropertyValueParam) MarshalJSON() (data []byte, err error) {
	type shadow APIAppendObjectPropertyValueParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *APIAppendObjectPropertyValueParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type APIAssociationDataSource struct {
	// Any of "HUBSPOT_DEFINED", "USER_DEFINED", "INTEGRATOR_DEFINED".
	AssociationCategory APIAssociationDataSourceAssociationCategory `json:"associationCategory,required"`
	AssociationTypeID   int64                                       `json:"associationTypeId,required"`
	Name                string                                      `json:"name,required"`
	ObjectTypeID        string                                      `json:"objectTypeId,required"`
	// Any of "ASSOCIATION".
	Type   APIAssociationDataSourceType `json:"type,required"`
	SortBy APISort                      `json:"sortBy"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AssociationCategory respjson.Field
		AssociationTypeID   respjson.Field
		Name                respjson.Field
		ObjectTypeID        respjson.Field
		Type                respjson.Field
		SortBy              respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r APIAssociationDataSource) RawJSON() string { return r.JSON.raw }
func (r *APIAssociationDataSource) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this APIAssociationDataSource to a
// APIAssociationDataSourceParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// APIAssociationDataSourceParam.Overrides()
func (r APIAssociationDataSource) ToParam() APIAssociationDataSourceParam {
	return param.Override[APIAssociationDataSourceParam](json.RawMessage(r.RawJSON()))
}

type APIAssociationDataSourceAssociationCategory string

const (
	APIAssociationDataSourceAssociationCategoryHubspotDefined    APIAssociationDataSourceAssociationCategory = "HUBSPOT_DEFINED"
	APIAssociationDataSourceAssociationCategoryUserDefined       APIAssociationDataSourceAssociationCategory = "USER_DEFINED"
	APIAssociationDataSourceAssociationCategoryIntegratorDefined APIAssociationDataSourceAssociationCategory = "INTEGRATOR_DEFINED"
)

type APIAssociationDataSourceType string

const (
	APIAssociationDataSourceTypeAssociation APIAssociationDataSourceType = "ASSOCIATION"
)

// The properties AssociationCategory, AssociationTypeID, Name, ObjectTypeID, Type
// are required.
type APIAssociationDataSourceParam struct {
	// Any of "HUBSPOT_DEFINED", "USER_DEFINED", "INTEGRATOR_DEFINED".
	AssociationCategory APIAssociationDataSourceAssociationCategory `json:"associationCategory,omitzero,required"`
	AssociationTypeID   int64                                       `json:"associationTypeId,required"`
	Name                string                                      `json:"name,required"`
	ObjectTypeID        string                                      `json:"objectTypeId,required"`
	// Any of "ASSOCIATION".
	Type   APIAssociationDataSourceType `json:"type,omitzero,required"`
	SortBy APISortParam                 `json:"sortBy,omitzero"`
	paramObj
}

func (r APIAssociationDataSourceParam) MarshalJSON() (data []byte, err error) {
	type shadow APIAssociationDataSourceParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *APIAssociationDataSourceParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type APIAssociationTimestampDataSource struct {
	// Any of "HUBSPOT_DEFINED", "USER_DEFINED", "INTEGRATOR_DEFINED".
	AssociationCategory APIAssociationTimestampDataSourceAssociationCategory `json:"associationCategory,required"`
	AssociationTypeID   int64                                                `json:"associationTypeId,required"`
	Name                string                                               `json:"name,required"`
	ObjectTypeID        string                                               `json:"objectTypeId,required"`
	// Any of "ASSOCIATION_TIMESTAMP".
	Type APIAssociationTimestampDataSourceType `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AssociationCategory respjson.Field
		AssociationTypeID   respjson.Field
		Name                respjson.Field
		ObjectTypeID        respjson.Field
		Type                respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r APIAssociationTimestampDataSource) RawJSON() string { return r.JSON.raw }
func (r *APIAssociationTimestampDataSource) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this APIAssociationTimestampDataSource to a
// APIAssociationTimestampDataSourceParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// APIAssociationTimestampDataSourceParam.Overrides()
func (r APIAssociationTimestampDataSource) ToParam() APIAssociationTimestampDataSourceParam {
	return param.Override[APIAssociationTimestampDataSourceParam](json.RawMessage(r.RawJSON()))
}

type APIAssociationTimestampDataSourceAssociationCategory string

const (
	APIAssociationTimestampDataSourceAssociationCategoryHubspotDefined    APIAssociationTimestampDataSourceAssociationCategory = "HUBSPOT_DEFINED"
	APIAssociationTimestampDataSourceAssociationCategoryUserDefined       APIAssociationTimestampDataSourceAssociationCategory = "USER_DEFINED"
	APIAssociationTimestampDataSourceAssociationCategoryIntegratorDefined APIAssociationTimestampDataSourceAssociationCategory = "INTEGRATOR_DEFINED"
)

type APIAssociationTimestampDataSourceType string

const (
	APIAssociationTimestampDataSourceTypeAssociationTimestamp APIAssociationTimestampDataSourceType = "ASSOCIATION_TIMESTAMP"
)

// The properties AssociationCategory, AssociationTypeID, Name, ObjectTypeID, Type
// are required.
type APIAssociationTimestampDataSourceParam struct {
	// Any of "HUBSPOT_DEFINED", "USER_DEFINED", "INTEGRATOR_DEFINED".
	AssociationCategory APIAssociationTimestampDataSourceAssociationCategory `json:"associationCategory,omitzero,required"`
	AssociationTypeID   int64                                                `json:"associationTypeId,required"`
	Name                string                                               `json:"name,required"`
	ObjectTypeID        string                                               `json:"objectTypeId,required"`
	// Any of "ASSOCIATION_TIMESTAMP".
	Type APIAssociationTimestampDataSourceType `json:"type,omitzero,required"`
	paramObj
}

func (r APIAssociationTimestampDataSourceParam) MarshalJSON() (data []byte, err error) {
	type shadow APIAssociationTimestampDataSourceParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *APIAssociationTimestampDataSourceParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type APIAuthKeyWebhookAuthSettings struct {
	// Where in the request this auth key should be located: "HEADER" or "QUERY_PARAM"
	//
	// Any of "HEADER", "QUERY_PARAM".
	Location APIAuthKeyWebhookAuthSettingsLocation `json:"location,required"`
	// The name to use for this auth key.
	Name string `json:"name,required"`
	// The secret to pass through in this auth key.
	SecretName string `json:"secretName,required"`
	// The type of webhook auth settings this is, can be: "AUTH_KEY" or "SIGNATURE"
	//
	// Any of "AUTH_KEY".
	Type APIAuthKeyWebhookAuthSettingsType `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Location    respjson.Field
		Name        respjson.Field
		SecretName  respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r APIAuthKeyWebhookAuthSettings) RawJSON() string { return r.JSON.raw }
func (r *APIAuthKeyWebhookAuthSettings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this APIAuthKeyWebhookAuthSettings to a
// APIAuthKeyWebhookAuthSettingsParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// APIAuthKeyWebhookAuthSettingsParam.Overrides()
func (r APIAuthKeyWebhookAuthSettings) ToParam() APIAuthKeyWebhookAuthSettingsParam {
	return param.Override[APIAuthKeyWebhookAuthSettingsParam](json.RawMessage(r.RawJSON()))
}

// Where in the request this auth key should be located: "HEADER" or "QUERY_PARAM"
type APIAuthKeyWebhookAuthSettingsLocation string

const (
	APIAuthKeyWebhookAuthSettingsLocationHeader     APIAuthKeyWebhookAuthSettingsLocation = "HEADER"
	APIAuthKeyWebhookAuthSettingsLocationQueryParam APIAuthKeyWebhookAuthSettingsLocation = "QUERY_PARAM"
)

// The type of webhook auth settings this is, can be: "AUTH_KEY" or "SIGNATURE"
type APIAuthKeyWebhookAuthSettingsType string

const (
	APIAuthKeyWebhookAuthSettingsTypeAuthKey APIAuthKeyWebhookAuthSettingsType = "AUTH_KEY"
)

// The properties Location, Name, SecretName, Type are required.
type APIAuthKeyWebhookAuthSettingsParam struct {
	// Where in the request this auth key should be located: "HEADER" or "QUERY_PARAM"
	//
	// Any of "HEADER", "QUERY_PARAM".
	Location APIAuthKeyWebhookAuthSettingsLocation `json:"location,omitzero,required"`
	// The name to use for this auth key.
	Name string `json:"name,required"`
	// The secret to pass through in this auth key.
	SecretName string `json:"secretName,required"`
	// The type of webhook auth settings this is, can be: "AUTH_KEY" or "SIGNATURE"
	//
	// Any of "AUTH_KEY".
	Type APIAuthKeyWebhookAuthSettingsType `json:"type,omitzero,required"`
	paramObj
}

func (r APIAuthKeyWebhookAuthSettingsParam) MarshalJSON() (data []byte, err error) {
	type shadow APIAuthKeyWebhookAuthSettingsParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *APIAuthKeyWebhookAuthSettingsParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type APIBlockedDate struct {
	DayOfMonth int64 `json:"dayOfMonth,required"`
	// Any of "JANUARY", "FEBRUARY", "MARCH", "APRIL", "MAY", "JUNE", "JULY", "AUGUST",
	// "SEPTEMBER", "OCTOBER", "NOVEMBER", "DECEMBER".
	Month APIBlockedDateMonth `json:"month,required"`
	Year  int64               `json:"year"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DayOfMonth  respjson.Field
		Month       respjson.Field
		Year        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r APIBlockedDate) RawJSON() string { return r.JSON.raw }
func (r *APIBlockedDate) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this APIBlockedDate to a APIBlockedDateParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// APIBlockedDateParam.Overrides()
func (r APIBlockedDate) ToParam() APIBlockedDateParam {
	return param.Override[APIBlockedDateParam](json.RawMessage(r.RawJSON()))
}

type APIBlockedDateMonth string

const (
	APIBlockedDateMonthJanuary   APIBlockedDateMonth = "JANUARY"
	APIBlockedDateMonthFebruary  APIBlockedDateMonth = "FEBRUARY"
	APIBlockedDateMonthMarch     APIBlockedDateMonth = "MARCH"
	APIBlockedDateMonthApril     APIBlockedDateMonth = "APRIL"
	APIBlockedDateMonthMay       APIBlockedDateMonth = "MAY"
	APIBlockedDateMonthJune      APIBlockedDateMonth = "JUNE"
	APIBlockedDateMonthJuly      APIBlockedDateMonth = "JULY"
	APIBlockedDateMonthAugust    APIBlockedDateMonth = "AUGUST"
	APIBlockedDateMonthSeptember APIBlockedDateMonth = "SEPTEMBER"
	APIBlockedDateMonthOctober   APIBlockedDateMonth = "OCTOBER"
	APIBlockedDateMonthNovember  APIBlockedDateMonth = "NOVEMBER"
	APIBlockedDateMonthDecember  APIBlockedDateMonth = "DECEMBER"
)

// The properties DayOfMonth, Month are required.
type APIBlockedDateParam struct {
	DayOfMonth int64 `json:"dayOfMonth,required"`
	// Any of "JANUARY", "FEBRUARY", "MARCH", "APRIL", "MAY", "JUNE", "JULY", "AUGUST",
	// "SEPTEMBER", "OCTOBER", "NOVEMBER", "DECEMBER".
	Month APIBlockedDateMonth `json:"month,omitzero,required"`
	Year  param.Opt[int64]    `json:"year,omitzero"`
	paramObj
}

func (r APIBlockedDateParam) MarshalJSON() (data []byte, err error) {
	type shadow APIBlockedDateParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *APIBlockedDateParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type APIConnection struct {
	EdgeType     string `json:"edgeType,required"`
	NextActionID string `json:"nextActionId,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		EdgeType     respjson.Field
		NextActionID respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r APIConnection) RawJSON() string { return r.JSON.raw }
func (r *APIConnection) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this APIConnection to a APIConnectionParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// APIConnectionParam.Overrides()
func (r APIConnection) ToParam() APIConnectionParam {
	return param.Override[APIConnectionParam](json.RawMessage(r.RawJSON()))
}

// The properties EdgeType, NextActionID are required.
type APIConnectionParam struct {
	EdgeType     string `json:"edgeType,required"`
	NextActionID string `json:"nextActionId,required"`
	paramObj
}

func (r APIConnectionParam) MarshalJSON() (data []byte, err error) {
	type shadow APIConnectionParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *APIConnectionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type APIContactFlow struct {
	ID                      string                      `json:"id,required"`
	Actions                 []APIContactFlowActionUnion `json:"actions,required"`
	BlockedDates            []APIBlockedDate            `json:"blockedDates,required"`
	CanEnrollFromSalesforce bool                        `json:"canEnrollFromSalesforce,required"`
	CreatedAt               time.Time                   `json:"createdAt,required" format:"date-time"`
	// Any of "PENDING", "COMPLETE".
	CRMObjectCreationStatus APIContactFlowCRMObjectCreationStatus `json:"crmObjectCreationStatus,required"`
	CustomProperties        map[string]string                     `json:"customProperties,required"`
	DataSources             []APIContactFlowDataSourceUnion       `json:"dataSources,required"`
	// Any of "WORKFLOW", "ACTION_SET", "UNKNOWN".
	FlowType              APIContactFlowFlowType `json:"flowType,required"`
	IsEnabled             bool                   `json:"isEnabled,required"`
	NextAvailableActionID string                 `json:"nextAvailableActionId,required"`
	ObjectTypeID          string                 `json:"objectTypeId,required"`
	RevisionID            string                 `json:"revisionId,required"`
	SuppressionListIDs    []int64                `json:"suppressionListIds,required"`
	TimeWindows           []APITimeWindow        `json:"timeWindows,required"`
	// Any of "CONTACT_FLOW", "PLATFORM_FLOW".
	Type                APIContactFlowType                    `json:"type,required"`
	UpdatedAt           time.Time                             `json:"updatedAt,required" format:"date-time"`
	Description         string                                `json:"description"`
	EnrollmentCriteria  APIContactFlowEnrollmentCriteriaUnion `json:"enrollmentCriteria"`
	EnrollmentSchedule  APIContactFlowEnrollmentScheduleUnion `json:"enrollmentSchedule"`
	EventAnchor         APIContactFlowEventAnchorUnion        `json:"eventAnchor"`
	GoalFilterBranch    APIContactFlowGoalFilterBranchUnion   `json:"goalFilterBranch"`
	Name                string                                `json:"name"`
	StartActionID       string                                `json:"startActionId"`
	UnEnrollmentSetting APIUnEnrollmentSetting                `json:"unEnrollmentSetting"`
	Uuid                string                                `json:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                      respjson.Field
		Actions                 respjson.Field
		BlockedDates            respjson.Field
		CanEnrollFromSalesforce respjson.Field
		CreatedAt               respjson.Field
		CRMObjectCreationStatus respjson.Field
		CustomProperties        respjson.Field
		DataSources             respjson.Field
		FlowType                respjson.Field
		IsEnabled               respjson.Field
		NextAvailableActionID   respjson.Field
		ObjectTypeID            respjson.Field
		RevisionID              respjson.Field
		SuppressionListIDs      respjson.Field
		TimeWindows             respjson.Field
		Type                    respjson.Field
		UpdatedAt               respjson.Field
		Description             respjson.Field
		EnrollmentCriteria      respjson.Field
		EnrollmentSchedule      respjson.Field
		EventAnchor             respjson.Field
		GoalFilterBranch        respjson.Field
		Name                    respjson.Field
		StartActionID           respjson.Field
		UnEnrollmentSetting     respjson.Field
		Uuid                    respjson.Field
		ExtraFields             map[string]respjson.Field
		raw                     string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r APIContactFlow) RawJSON() string { return r.JSON.raw }
func (r *APIContactFlow) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// APIContactFlowActionUnion contains all possible properties and values from
// [APIStaticBranchAction], [APIListBranchAction], [APIAbTestBranchAction],
// [APICustomCodeAction], [APIWebhookAction], [APISingleConnectionAction].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type APIContactFlowActionUnion struct {
	ActionID string `json:"actionId"`
	// This field is from variant [APIStaticBranchAction].
	InputValue APIStaticBranchActionInputValueUnion `json:"inputValue"`
	// This field is from variant [APIStaticBranchAction].
	StaticBranches []APIStaticBranch `json:"staticBranches"`
	Type           string            `json:"type"`
	// This field is from variant [APIStaticBranchAction].
	DefaultBranch     APIConnection `json:"defaultBranch"`
	DefaultBranchName string        `json:"defaultBranchName"`
	// This field is from variant [APIListBranchAction].
	ListBranches []APIListBranch `json:"listBranches"`
	// This field is from variant [APIAbTestBranchAction].
	TestBranches []APIConnection `json:"testBranches"`
	// This field is from variant [APICustomCodeAction].
	InputFields []APIInputVariable `json:"inputFields"`
	// This field is from variant [APICustomCodeAction].
	OutputFields []APIEnumerationOutputField `json:"outputFields"`
	// This field is from variant [APICustomCodeAction].
	Runtime string `json:"runtime"`
	// This field is from variant [APICustomCodeAction].
	SecretNames []string `json:"secretNames"`
	// This field is from variant [APICustomCodeAction].
	SourceCode string `json:"sourceCode"`
	// This field is from variant [APICustomCodeAction].
	Connection APIConnection `json:"connection"`
	// This field is from variant [APIWebhookAction].
	Method APIWebhookActionMethod `json:"method"`
	// This field is from variant [APIWebhookAction].
	QueryParams []APIInputVariable `json:"queryParams"`
	// This field is from variant [APIWebhookAction].
	WebhookURL string `json:"webhookUrl"`
	// This field is from variant [APIWebhookAction].
	AuthSettings APIWebhookActionAuthSettingsUnion `json:"authSettings"`
	// This field is from variant [APISingleConnectionAction].
	ActionTypeID string `json:"actionTypeId"`
	// This field is from variant [APISingleConnectionAction].
	ActionTypeVersion int64 `json:"actionTypeVersion"`
	// This field is from variant [APISingleConnectionAction].
	Fields map[string]any `json:"fields"`
	JSON   struct {
		ActionID          respjson.Field
		InputValue        respjson.Field
		StaticBranches    respjson.Field
		Type              respjson.Field
		DefaultBranch     respjson.Field
		DefaultBranchName respjson.Field
		ListBranches      respjson.Field
		TestBranches      respjson.Field
		InputFields       respjson.Field
		OutputFields      respjson.Field
		Runtime           respjson.Field
		SecretNames       respjson.Field
		SourceCode        respjson.Field
		Connection        respjson.Field
		Method            respjson.Field
		QueryParams       respjson.Field
		WebhookURL        respjson.Field
		AuthSettings      respjson.Field
		ActionTypeID      respjson.Field
		ActionTypeVersion respjson.Field
		Fields            respjson.Field
		raw               string
	} `json:"-"`
}

func (u APIContactFlowActionUnion) AsStaticBranch() (v APIStaticBranchAction) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u APIContactFlowActionUnion) AsListBranch() (v APIListBranchAction) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u APIContactFlowActionUnion) AsAbTestBranch() (v APIAbTestBranchAction) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u APIContactFlowActionUnion) AsCustomCode() (v APICustomCodeAction) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u APIContactFlowActionUnion) AsWebhook() (v APIWebhookAction) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u APIContactFlowActionUnion) AsSingleConnection() (v APISingleConnectionAction) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u APIContactFlowActionUnion) RawJSON() string { return u.JSON.raw }

func (r *APIContactFlowActionUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type APIContactFlowCRMObjectCreationStatus string

const (
	APIContactFlowCRMObjectCreationStatusPending  APIContactFlowCRMObjectCreationStatus = "PENDING"
	APIContactFlowCRMObjectCreationStatusComplete APIContactFlowCRMObjectCreationStatus = "COMPLETE"
)

// APIContactFlowDataSourceUnion contains all possible properties and values from
// [APIAssociationDataSource], [APIAssociationTimestampDataSource],
// [APIStaticPropertyFilterDataSource],
// [APIEnrolledRecordPropertyFilterDataSource],
// [APIDatasetFieldPropertyFilterDataSource],
// [APIEnrolledArgumentPropertyFilterDataSource].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type APIContactFlowDataSourceUnion struct {
	AssociationCategory string `json:"associationCategory"`
	AssociationTypeID   int64  `json:"associationTypeId"`
	Name                string `json:"name"`
	ObjectTypeID        string `json:"objectTypeId"`
	Type                string `json:"type"`
	// This field is from variant [APIAssociationDataSource].
	SortBy       APISort `json:"sortBy"`
	PropertyName string  `json:"propertyName"`
	// This field is from variant [APIStaticPropertyFilterDataSource].
	StaticValue string `json:"staticValue"`
	// This field is from variant [APIEnrolledRecordPropertyFilterDataSource].
	RecordFieldName string `json:"recordFieldName"`
	// This field is from variant [APIDatasetFieldPropertyFilterDataSource].
	DatasetFieldName string `json:"datasetFieldName"`
	// This field is from variant [APIEnrolledArgumentPropertyFilterDataSource].
	ArgumentName string `json:"argumentName"`
	JSON         struct {
		AssociationCategory respjson.Field
		AssociationTypeID   respjson.Field
		Name                respjson.Field
		ObjectTypeID        respjson.Field
		Type                respjson.Field
		SortBy              respjson.Field
		PropertyName        respjson.Field
		StaticValue         respjson.Field
		RecordFieldName     respjson.Field
		DatasetFieldName    respjson.Field
		ArgumentName        respjson.Field
		raw                 string
	} `json:"-"`
}

func (u APIContactFlowDataSourceUnion) AsAssociation() (v APIAssociationDataSource) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u APIContactFlowDataSourceUnion) AsAssociationTimestamp() (v APIAssociationTimestampDataSource) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u APIContactFlowDataSourceUnion) AsStaticPropertyFilter() (v APIStaticPropertyFilterDataSource) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u APIContactFlowDataSourceUnion) AsEnrolledRecordPropertyFilter() (v APIEnrolledRecordPropertyFilterDataSource) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u APIContactFlowDataSourceUnion) AsDatasetFieldPropertyFilter() (v APIDatasetFieldPropertyFilterDataSource) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u APIContactFlowDataSourceUnion) AsEnrolledArgumentPropertyFilter() (v APIEnrolledArgumentPropertyFilterDataSource) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u APIContactFlowDataSourceUnion) RawJSON() string { return u.JSON.raw }

func (r *APIContactFlowDataSourceUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type APIContactFlowFlowType string

const (
	APIContactFlowFlowTypeWorkflow  APIContactFlowFlowType = "WORKFLOW"
	APIContactFlowFlowTypeActionSet APIContactFlowFlowType = "ACTION_SET"
	APIContactFlowFlowTypeUnknown   APIContactFlowFlowType = "UNKNOWN"
)

type APIContactFlowType string

const (
	APIContactFlowTypeContactFlow  APIContactFlowType = "CONTACT_FLOW"
	APIContactFlowTypePlatformFlow APIContactFlowType = "PLATFORM_FLOW"
)

// APIContactFlowEnrollmentCriteriaUnion contains all possible properties and
// values from [APIListBasedEnrollmentCriteria], [APIEventBasedEnrollmentCriteria],
// [APIManualEnrollmentCriteria].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type APIContactFlowEnrollmentCriteriaUnion struct {
	// This field is from variant [APIListBasedEnrollmentCriteria].
	ListFilterBranch APIListBasedEnrollmentCriteriaListFilterBranchUnion `json:"listFilterBranch"`
	// This field is from variant [APIListBasedEnrollmentCriteria].
	ReEnrollmentTriggersFilterBranches []APIListBasedEnrollmentCriteriaReEnrollmentTriggersFilterBranchUnion `json:"reEnrollmentTriggersFilterBranches"`
	ShouldReEnroll                     bool                                                                  `json:"shouldReEnroll"`
	Type                               string                                                                `json:"type"`
	// This field is from variant [APIListBasedEnrollmentCriteria].
	UnEnrollObjectsNotMeetingCriteria bool `json:"unEnrollObjectsNotMeetingCriteria"`
	// This field is from variant [APIEventBasedEnrollmentCriteria].
	EventFilterBranches []shared.PublicUnifiedEventsFilterBranch `json:"eventFilterBranches"`
	// This field is from variant [APIEventBasedEnrollmentCriteria].
	ListMembershipFilterBranches []APIEventBasedEnrollmentCriteriaListMembershipFilterBranchUnion `json:"listMembershipFilterBranches"`
	// This field is from variant [APIEventBasedEnrollmentCriteria].
	RefinementCriteria APIEventBasedEnrollmentCriteriaRefinementCriteriaUnion `json:"refinementCriteria"`
	JSON               struct {
		ListFilterBranch                   respjson.Field
		ReEnrollmentTriggersFilterBranches respjson.Field
		ShouldReEnroll                     respjson.Field
		Type                               respjson.Field
		UnEnrollObjectsNotMeetingCriteria  respjson.Field
		EventFilterBranches                respjson.Field
		ListMembershipFilterBranches       respjson.Field
		RefinementCriteria                 respjson.Field
		raw                                string
	} `json:"-"`
}

func (u APIContactFlowEnrollmentCriteriaUnion) AsListBased() (v APIListBasedEnrollmentCriteria) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u APIContactFlowEnrollmentCriteriaUnion) AsEventBased() (v APIEventBasedEnrollmentCriteria) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u APIContactFlowEnrollmentCriteriaUnion) AsManual() (v APIManualEnrollmentCriteria) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u APIContactFlowEnrollmentCriteriaUnion) RawJSON() string { return u.JSON.raw }

func (r *APIContactFlowEnrollmentCriteriaUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// APIContactFlowEnrollmentScheduleUnion contains all possible properties and
// values from [APIDailyEnrollmentSchedule], [APIWeeklyEnrollmentSchedule],
// [APIMonthlySpecificDaysEnrollmentSchedule],
// [APIMonthlyRelativeDaysEnrollmentSchedule], [APIYearlyEnrollmentSchedule],
// [APIPropertyBasedEnrollmentSchedule].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type APIContactFlowEnrollmentScheduleUnion struct {
	// This field is from variant [APIDailyEnrollmentSchedule].
	TimeOfDay APITimeOfDay `json:"timeOfDay"`
	Type      string       `json:"type"`
	// This field is from variant [APIWeeklyEnrollmentSchedule].
	DaysOfWeek []string `json:"daysOfWeek"`
	// This field is from variant [APIMonthlySpecificDaysEnrollmentSchedule].
	DaysOfMonth []int64 `json:"daysOfMonth"`
	// This field is from variant [APIMonthlyRelativeDaysEnrollmentSchedule].
	MonthlyRelativeDays APIMonthlyRelativeDaysEnrollmentScheduleMonthlyRelativeDays `json:"monthlyRelativeDays"`
	// This field is from variant [APIYearlyEnrollmentSchedule].
	DayOfMonth int64 `json:"dayOfMonth"`
	// This field is from variant [APIYearlyEnrollmentSchedule].
	Month APIYearlyEnrollmentScheduleMonth `json:"month"`
	// This field is from variant [APIPropertyBasedEnrollmentSchedule].
	DateProperty string `json:"dateProperty"`
	// This field is from variant [APIPropertyBasedEnrollmentSchedule].
	DaysDelta int64 `json:"daysDelta"`
	// This field is from variant [APIPropertyBasedEnrollmentSchedule].
	Yearly bool `json:"yearly"`
	JSON   struct {
		TimeOfDay           respjson.Field
		Type                respjson.Field
		DaysOfWeek          respjson.Field
		DaysOfMonth         respjson.Field
		MonthlyRelativeDays respjson.Field
		DayOfMonth          respjson.Field
		Month               respjson.Field
		DateProperty        respjson.Field
		DaysDelta           respjson.Field
		Yearly              respjson.Field
		raw                 string
	} `json:"-"`
}

func (u APIContactFlowEnrollmentScheduleUnion) AsDaily() (v APIDailyEnrollmentSchedule) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u APIContactFlowEnrollmentScheduleUnion) AsWeekly() (v APIWeeklyEnrollmentSchedule) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u APIContactFlowEnrollmentScheduleUnion) AsMonthlySpecificDays() (v APIMonthlySpecificDaysEnrollmentSchedule) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u APIContactFlowEnrollmentScheduleUnion) AsMonthlyRelativeDays() (v APIMonthlyRelativeDaysEnrollmentSchedule) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u APIContactFlowEnrollmentScheduleUnion) AsYearly() (v APIYearlyEnrollmentSchedule) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u APIContactFlowEnrollmentScheduleUnion) AsPropertyBased() (v APIPropertyBasedEnrollmentSchedule) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u APIContactFlowEnrollmentScheduleUnion) RawJSON() string { return u.JSON.raw }

func (r *APIContactFlowEnrollmentScheduleUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// APIContactFlowEventAnchorUnion contains all possible properties and values from
// [APIContactPropertyAnchor], [APIStaticDateAnchor].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type APIContactFlowEventAnchorUnion struct {
	// This field is from variant [APIContactPropertyAnchor].
	ContactProperty string `json:"contactProperty"`
	Type            string `json:"type"`
	// This field is from variant [APIStaticDateAnchor].
	DayOfMonth int64 `json:"dayOfMonth"`
	// This field is from variant [APIStaticDateAnchor].
	Month APIStaticDateAnchorMonth `json:"month"`
	// This field is from variant [APIStaticDateAnchor].
	Year int64 `json:"year"`
	JSON struct {
		ContactProperty respjson.Field
		Type            respjson.Field
		DayOfMonth      respjson.Field
		Month           respjson.Field
		Year            respjson.Field
		raw             string
	} `json:"-"`
}

func (u APIContactFlowEventAnchorUnion) AsContactPropertyAnchor() (v APIContactPropertyAnchor) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u APIContactFlowEventAnchorUnion) AsStaticDateAnchor() (v APIStaticDateAnchor) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u APIContactFlowEventAnchorUnion) RawJSON() string { return u.JSON.raw }

func (r *APIContactFlowEventAnchorUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// APIContactFlowGoalFilterBranchUnion contains all possible properties and values
// from [shared.PublicOrFilterBranch], [shared.PublicAndFilterBranch],
// [shared.PublicNotAllFilterBranch], [shared.PublicNotAnyFilterBranch],
// [shared.PublicRestrictedFilterBranch], [shared.PublicUnifiedEventsFilterBranch],
// [shared.PublicPropertyAssociationFilterBranch],
// [shared.PublicAssociationFilterBranch].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type APIContactFlowGoalFilterBranchUnion struct {
	// This field is a union of [[]shared.PublicOrFilterBranchFilterBranchUnion],
	// [[]shared.PublicAndFilterBranchFilterBranchUnion],
	// [[]shared.PublicNotAllFilterBranchFilterBranchUnion],
	// [[]shared.PublicNotAnyFilterBranchFilterBranchUnion],
	// [[]shared.PublicRestrictedFilterBranchFilterBranchUnion],
	// [[]shared.PublicUnifiedEventsFilterBranchFilterBranchUnion],
	// [[]shared.PublicPropertyAssociationFilterBranchFilterBranchUnion],
	// [[]shared.PublicAssociationFilterBranchFilterBranchUnion]
	FilterBranches       APIContactFlowGoalFilterBranchUnionFilterBranches `json:"filterBranches"`
	FilterBranchOperator string                                            `json:"filterBranchOperator"`
	FilterBranchType     string                                            `json:"filterBranchType"`
	// This field is a union of [[]shared.PublicOrFilterBranchFilterUnion],
	// [[]shared.PublicAndFilterBranchFilterUnion],
	// [[]shared.PublicNotAllFilterBranchFilterUnion],
	// [[]shared.PublicNotAnyFilterBranchFilterUnion],
	// [[]shared.PublicRestrictedFilterBranchFilterUnion],
	// [[]shared.PublicUnifiedEventsFilterBranchFilterUnion],
	// [[]shared.PublicPropertyAssociationFilterBranchFilterUnion],
	// [[]shared.PublicAssociationFilterBranchFilterUnion]
	Filters APIContactFlowGoalFilterBranchUnionFilters `json:"filters"`
	// This field is from variant [shared.PublicUnifiedEventsFilterBranch].
	EventTypeID string `json:"eventTypeId"`
	Operator    string `json:"operator"`
	// This field is from variant [shared.PublicUnifiedEventsFilterBranch].
	CoalescingRefineBy shared.PublicUnifiedEventsFilterBranchCoalescingRefineByUnion `json:"coalescingRefineBy"`
	ObjectTypeID       string                                                        `json:"objectTypeId"`
	// This field is from variant [shared.PublicPropertyAssociationFilterBranch].
	PropertyWithObjectID string `json:"propertyWithObjectId"`
	// This field is from variant [shared.PublicAssociationFilterBranch].
	AssociationCategory string `json:"associationCategory"`
	// This field is from variant [shared.PublicAssociationFilterBranch].
	AssociationTypeID int64 `json:"associationTypeId"`
	JSON              struct {
		FilterBranches       respjson.Field
		FilterBranchOperator respjson.Field
		FilterBranchType     respjson.Field
		Filters              respjson.Field
		EventTypeID          respjson.Field
		Operator             respjson.Field
		CoalescingRefineBy   respjson.Field
		ObjectTypeID         respjson.Field
		PropertyWithObjectID respjson.Field
		AssociationCategory  respjson.Field
		AssociationTypeID    respjson.Field
		raw                  string
	} `json:"-"`
}

func (u APIContactFlowGoalFilterBranchUnion) AsOr() (v shared.PublicOrFilterBranch) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u APIContactFlowGoalFilterBranchUnion) AsAnd() (v shared.PublicAndFilterBranch) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u APIContactFlowGoalFilterBranchUnion) AsNotAll() (v shared.PublicNotAllFilterBranch) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u APIContactFlowGoalFilterBranchUnion) AsNotAny() (v shared.PublicNotAnyFilterBranch) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u APIContactFlowGoalFilterBranchUnion) AsRestricted() (v shared.PublicRestrictedFilterBranch) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u APIContactFlowGoalFilterBranchUnion) AsUnifiedEvents() (v shared.PublicUnifiedEventsFilterBranch) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u APIContactFlowGoalFilterBranchUnion) AsPropertyAssociation() (v shared.PublicPropertyAssociationFilterBranch) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u APIContactFlowGoalFilterBranchUnion) AsAssociation() (v shared.PublicAssociationFilterBranch) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u APIContactFlowGoalFilterBranchUnion) RawJSON() string { return u.JSON.raw }

func (r *APIContactFlowGoalFilterBranchUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// APIContactFlowGoalFilterBranchUnionFilterBranches is an implicit subunion of
// [APIContactFlowGoalFilterBranchUnion].
// APIContactFlowGoalFilterBranchUnionFilterBranches provides convenient access to
// the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [APIContactFlowGoalFilterBranchUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfPublicOrFilterBranchFilterBranches
// OfPublicAndFilterBranchFilterBranches OfPublicNotAllFilterBranchFilterBranches
// OfPublicNotAnyFilterBranchFilterBranches
// OfPublicRestrictedFilterBranchFilterBranches
// OfPublicUnifiedEventsFilterBranchFilterBranches
// OfPublicPropertyAssociationFilterBranchFilterBranches
// OfPublicAssociationFilterBranchFilterBranches]
type APIContactFlowGoalFilterBranchUnionFilterBranches struct {
	// This field will be present if the value is a
	// [[]shared.PublicOrFilterBranchFilterBranchUnion] instead of an object.
	OfPublicOrFilterBranchFilterBranches []shared.PublicOrFilterBranchFilterBranchUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]shared.PublicAndFilterBranchFilterBranchUnion] instead of an object.
	OfPublicAndFilterBranchFilterBranches []shared.PublicAndFilterBranchFilterBranchUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]shared.PublicNotAllFilterBranchFilterBranchUnion] instead of an object.
	OfPublicNotAllFilterBranchFilterBranches []shared.PublicNotAllFilterBranchFilterBranchUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]shared.PublicNotAnyFilterBranchFilterBranchUnion] instead of an object.
	OfPublicNotAnyFilterBranchFilterBranches []shared.PublicNotAnyFilterBranchFilterBranchUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]shared.PublicRestrictedFilterBranchFilterBranchUnion] instead of an object.
	OfPublicRestrictedFilterBranchFilterBranches []shared.PublicRestrictedFilterBranchFilterBranchUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]shared.PublicUnifiedEventsFilterBranchFilterBranchUnion] instead of an
	// object.
	OfPublicUnifiedEventsFilterBranchFilterBranches []shared.PublicUnifiedEventsFilterBranchFilterBranchUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]shared.PublicPropertyAssociationFilterBranchFilterBranchUnion] instead of an
	// object.
	OfPublicPropertyAssociationFilterBranchFilterBranches []shared.PublicPropertyAssociationFilterBranchFilterBranchUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]shared.PublicAssociationFilterBranchFilterBranchUnion] instead of an object.
	OfPublicAssociationFilterBranchFilterBranches []shared.PublicAssociationFilterBranchFilterBranchUnion `json:",inline"`
	JSON                                          struct {
		OfPublicOrFilterBranchFilterBranches                  respjson.Field
		OfPublicAndFilterBranchFilterBranches                 respjson.Field
		OfPublicNotAllFilterBranchFilterBranches              respjson.Field
		OfPublicNotAnyFilterBranchFilterBranches              respjson.Field
		OfPublicRestrictedFilterBranchFilterBranches          respjson.Field
		OfPublicUnifiedEventsFilterBranchFilterBranches       respjson.Field
		OfPublicPropertyAssociationFilterBranchFilterBranches respjson.Field
		OfPublicAssociationFilterBranchFilterBranches         respjson.Field
		raw                                                   string
	} `json:"-"`
}

func (r *APIContactFlowGoalFilterBranchUnionFilterBranches) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// APIContactFlowGoalFilterBranchUnionFilters is an implicit subunion of
// [APIContactFlowGoalFilterBranchUnion].
// APIContactFlowGoalFilterBranchUnionFilters provides convenient access to the
// sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [APIContactFlowGoalFilterBranchUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfPublicOrFilterBranchFilters OfPublicAndFilterBranchFilters
// OfPublicNotAllFilterBranchFilters OfPublicNotAnyFilterBranchFilters
// OfPublicRestrictedFilterBranchFilters OfPublicUnifiedEventsFilterBranchFilters
// OfPublicPropertyAssociationFilterBranchFilters
// OfPublicAssociationFilterBranchFilters]
type APIContactFlowGoalFilterBranchUnionFilters struct {
	// This field will be present if the value is a
	// [[]shared.PublicOrFilterBranchFilterUnion] instead of an object.
	OfPublicOrFilterBranchFilters []shared.PublicOrFilterBranchFilterUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]shared.PublicAndFilterBranchFilterUnion] instead of an object.
	OfPublicAndFilterBranchFilters []shared.PublicAndFilterBranchFilterUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]shared.PublicNotAllFilterBranchFilterUnion] instead of an object.
	OfPublicNotAllFilterBranchFilters []shared.PublicNotAllFilterBranchFilterUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]shared.PublicNotAnyFilterBranchFilterUnion] instead of an object.
	OfPublicNotAnyFilterBranchFilters []shared.PublicNotAnyFilterBranchFilterUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]shared.PublicRestrictedFilterBranchFilterUnion] instead of an object.
	OfPublicRestrictedFilterBranchFilters []shared.PublicRestrictedFilterBranchFilterUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]shared.PublicUnifiedEventsFilterBranchFilterUnion] instead of an object.
	OfPublicUnifiedEventsFilterBranchFilters []shared.PublicUnifiedEventsFilterBranchFilterUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]shared.PublicPropertyAssociationFilterBranchFilterUnion] instead of an
	// object.
	OfPublicPropertyAssociationFilterBranchFilters []shared.PublicPropertyAssociationFilterBranchFilterUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]shared.PublicAssociationFilterBranchFilterUnion] instead of an object.
	OfPublicAssociationFilterBranchFilters []shared.PublicAssociationFilterBranchFilterUnion `json:",inline"`
	JSON                                   struct {
		OfPublicOrFilterBranchFilters                  respjson.Field
		OfPublicAndFilterBranchFilters                 respjson.Field
		OfPublicNotAllFilterBranchFilters              respjson.Field
		OfPublicNotAnyFilterBranchFilters              respjson.Field
		OfPublicRestrictedFilterBranchFilters          respjson.Field
		OfPublicUnifiedEventsFilterBranchFilters       respjson.Field
		OfPublicPropertyAssociationFilterBranchFilters respjson.Field
		OfPublicAssociationFilterBranchFilters         respjson.Field
		raw                                            string
	} `json:"-"`
}

func (r *APIContactFlowGoalFilterBranchUnionFilters) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Actions, BlockedDates, CanEnrollFromSalesforce, CustomProperties,
// DataSources, FlowType, IsEnabled, ObjectTypeID, SuppressionListIDs, TimeWindows,
// Type are required.
type APIContactFlowCreateRequestParam struct {
	Actions                 []APIContactFlowCreateRequestActionUnionParam     `json:"actions,omitzero,required"`
	BlockedDates            []APIBlockedDateParam                             `json:"blockedDates,omitzero,required"`
	CanEnrollFromSalesforce bool                                              `json:"canEnrollFromSalesforce,required"`
	CustomProperties        map[string]string                                 `json:"customProperties,omitzero,required"`
	DataSources             []APIContactFlowCreateRequestDataSourceUnionParam `json:"dataSources,omitzero,required"`
	// Any of "WORKFLOW", "ACTION_SET", "UNKNOWN".
	FlowType           APIContactFlowCreateRequestFlowType `json:"flowType,omitzero,required"`
	IsEnabled          bool                                `json:"isEnabled,required"`
	ObjectTypeID       string                              `json:"objectTypeId,required"`
	SuppressionListIDs []int64                             `json:"suppressionListIds,omitzero,required"`
	TimeWindows        []APITimeWindowParam                `json:"timeWindows,omitzero,required"`
	// Any of "CONTACT_FLOW", "PLATFORM_FLOW".
	Type                APIContactFlowCreateRequestType                         `json:"type,omitzero,required"`
	Description         param.Opt[string]                                       `json:"description,omitzero"`
	Name                param.Opt[string]                                       `json:"name,omitzero"`
	StartActionID       param.Opt[string]                                       `json:"startActionId,omitzero"`
	Uuid                param.Opt[string]                                       `json:"uuid,omitzero"`
	EnrollmentCriteria  APIContactFlowCreateRequestEnrollmentCriteriaUnionParam `json:"enrollmentCriteria,omitzero"`
	EnrollmentSchedule  APIContactFlowCreateRequestEnrollmentScheduleUnionParam `json:"enrollmentSchedule,omitzero"`
	EventAnchor         APIContactFlowCreateRequestEventAnchorUnionParam        `json:"eventAnchor,omitzero"`
	GoalFilterBranch    APIContactFlowCreateRequestGoalFilterBranchUnionParam   `json:"goalFilterBranch,omitzero"`
	UnEnrollmentSetting APIUnEnrollmentSettingParam                             `json:"unEnrollmentSetting,omitzero"`
	paramObj
}

func (r APIContactFlowCreateRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow APIContactFlowCreateRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *APIContactFlowCreateRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type APIContactFlowCreateRequestActionUnionParam struct {
	OfStaticBranch     *APIStaticBranchActionParam     `json:",omitzero,inline"`
	OfListBranch       *APIListBranchActionParam       `json:",omitzero,inline"`
	OfAbTestBranch     *APIAbTestBranchActionParam     `json:",omitzero,inline"`
	OfCustomCode       *APICustomCodeActionParam       `json:",omitzero,inline"`
	OfWebhook          *APIWebhookActionParam          `json:",omitzero,inline"`
	OfSingleConnection *APISingleConnectionActionParam `json:",omitzero,inline"`
	paramUnion
}

func (u APIContactFlowCreateRequestActionUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfStaticBranch,
		u.OfListBranch,
		u.OfAbTestBranch,
		u.OfCustomCode,
		u.OfWebhook,
		u.OfSingleConnection)
}
func (u *APIContactFlowCreateRequestActionUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *APIContactFlowCreateRequestActionUnionParam) asAny() any {
	if !param.IsOmitted(u.OfStaticBranch) {
		return u.OfStaticBranch
	} else if !param.IsOmitted(u.OfListBranch) {
		return u.OfListBranch
	} else if !param.IsOmitted(u.OfAbTestBranch) {
		return u.OfAbTestBranch
	} else if !param.IsOmitted(u.OfCustomCode) {
		return u.OfCustomCode
	} else if !param.IsOmitted(u.OfWebhook) {
		return u.OfWebhook
	} else if !param.IsOmitted(u.OfSingleConnection) {
		return u.OfSingleConnection
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIContactFlowCreateRequestActionUnionParam) GetInputValue() *APIStaticBranchActionInputValueUnionParam {
	if vt := u.OfStaticBranch; vt != nil {
		return &vt.InputValue
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIContactFlowCreateRequestActionUnionParam) GetStaticBranches() []APIStaticBranchParam {
	if vt := u.OfStaticBranch; vt != nil {
		return vt.StaticBranches
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIContactFlowCreateRequestActionUnionParam) GetListBranches() []APIListBranchParam {
	if vt := u.OfListBranch; vt != nil {
		return vt.ListBranches
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIContactFlowCreateRequestActionUnionParam) GetTestBranches() []APIConnectionParam {
	if vt := u.OfAbTestBranch; vt != nil {
		return vt.TestBranches
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIContactFlowCreateRequestActionUnionParam) GetInputFields() []APIInputVariableParam {
	if vt := u.OfCustomCode; vt != nil {
		return vt.InputFields
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIContactFlowCreateRequestActionUnionParam) GetOutputFields() []APIEnumerationOutputFieldParam {
	if vt := u.OfCustomCode; vt != nil {
		return vt.OutputFields
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIContactFlowCreateRequestActionUnionParam) GetRuntime() *string {
	if vt := u.OfCustomCode; vt != nil {
		return &vt.Runtime
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIContactFlowCreateRequestActionUnionParam) GetSecretNames() []string {
	if vt := u.OfCustomCode; vt != nil {
		return vt.SecretNames
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIContactFlowCreateRequestActionUnionParam) GetSourceCode() *string {
	if vt := u.OfCustomCode; vt != nil {
		return &vt.SourceCode
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIContactFlowCreateRequestActionUnionParam) GetMethod() *string {
	if vt := u.OfWebhook; vt != nil {
		return (*string)(&vt.Method)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIContactFlowCreateRequestActionUnionParam) GetQueryParams() []APIInputVariableParam {
	if vt := u.OfWebhook; vt != nil {
		return vt.QueryParams
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIContactFlowCreateRequestActionUnionParam) GetWebhookURL() *string {
	if vt := u.OfWebhook; vt != nil {
		return &vt.WebhookURL
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIContactFlowCreateRequestActionUnionParam) GetAuthSettings() *APIWebhookActionAuthSettingsUnionParam {
	if vt := u.OfWebhook; vt != nil {
		return &vt.AuthSettings
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIContactFlowCreateRequestActionUnionParam) GetActionTypeID() *string {
	if vt := u.OfSingleConnection; vt != nil {
		return &vt.ActionTypeID
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIContactFlowCreateRequestActionUnionParam) GetActionTypeVersion() *int64 {
	if vt := u.OfSingleConnection; vt != nil {
		return &vt.ActionTypeVersion
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIContactFlowCreateRequestActionUnionParam) GetFields() map[string]any {
	if vt := u.OfSingleConnection; vt != nil {
		return vt.Fields
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIContactFlowCreateRequestActionUnionParam) GetActionID() *string {
	if vt := u.OfStaticBranch; vt != nil {
		return (*string)(&vt.ActionID)
	} else if vt := u.OfListBranch; vt != nil {
		return (*string)(&vt.ActionID)
	} else if vt := u.OfAbTestBranch; vt != nil {
		return (*string)(&vt.ActionID)
	} else if vt := u.OfCustomCode; vt != nil {
		return (*string)(&vt.ActionID)
	} else if vt := u.OfWebhook; vt != nil {
		return (*string)(&vt.ActionID)
	} else if vt := u.OfSingleConnection; vt != nil {
		return (*string)(&vt.ActionID)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIContactFlowCreateRequestActionUnionParam) GetType() *string {
	if vt := u.OfStaticBranch; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfListBranch; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfAbTestBranch; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfCustomCode; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfWebhook; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfSingleConnection; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIContactFlowCreateRequestActionUnionParam) GetDefaultBranchName() *string {
	if vt := u.OfStaticBranch; vt != nil && vt.DefaultBranchName.Valid() {
		return &vt.DefaultBranchName.Value
	} else if vt := u.OfListBranch; vt != nil && vt.DefaultBranchName.Valid() {
		return &vt.DefaultBranchName.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's DefaultBranch property, if
// present.
func (u APIContactFlowCreateRequestActionUnionParam) GetDefaultBranch() *APIConnectionParam {
	if vt := u.OfStaticBranch; vt != nil {
		return &vt.DefaultBranch
	} else if vt := u.OfListBranch; vt != nil {
		return &vt.DefaultBranch
	}
	return nil
}

// Returns a pointer to the underlying variant's Connection property, if present.
func (u APIContactFlowCreateRequestActionUnionParam) GetConnection() *APIConnectionParam {
	if vt := u.OfCustomCode; vt != nil {
		return &vt.Connection
	} else if vt := u.OfWebhook; vt != nil {
		return &vt.Connection
	} else if vt := u.OfSingleConnection; vt != nil {
		return &vt.Connection
	}
	return nil
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type APIContactFlowCreateRequestDataSourceUnionParam struct {
	OfAssociation                    *APIAssociationDataSourceParam                    `json:",omitzero,inline"`
	OfAssociationTimestamp           *APIAssociationTimestampDataSourceParam           `json:",omitzero,inline"`
	OfStaticPropertyFilter           *APIStaticPropertyFilterDataSourceParam           `json:",omitzero,inline"`
	OfEnrolledRecordPropertyFilter   *APIEnrolledRecordPropertyFilterDataSourceParam   `json:",omitzero,inline"`
	OfDatasetFieldPropertyFilter     *APIDatasetFieldPropertyFilterDataSourceParam     `json:",omitzero,inline"`
	OfEnrolledArgumentPropertyFilter *APIEnrolledArgumentPropertyFilterDataSourceParam `json:",omitzero,inline"`
	paramUnion
}

func (u APIContactFlowCreateRequestDataSourceUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfAssociation,
		u.OfAssociationTimestamp,
		u.OfStaticPropertyFilter,
		u.OfEnrolledRecordPropertyFilter,
		u.OfDatasetFieldPropertyFilter,
		u.OfEnrolledArgumentPropertyFilter)
}
func (u *APIContactFlowCreateRequestDataSourceUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *APIContactFlowCreateRequestDataSourceUnionParam) asAny() any {
	if !param.IsOmitted(u.OfAssociation) {
		return u.OfAssociation
	} else if !param.IsOmitted(u.OfAssociationTimestamp) {
		return u.OfAssociationTimestamp
	} else if !param.IsOmitted(u.OfStaticPropertyFilter) {
		return u.OfStaticPropertyFilter
	} else if !param.IsOmitted(u.OfEnrolledRecordPropertyFilter) {
		return u.OfEnrolledRecordPropertyFilter
	} else if !param.IsOmitted(u.OfDatasetFieldPropertyFilter) {
		return u.OfDatasetFieldPropertyFilter
	} else if !param.IsOmitted(u.OfEnrolledArgumentPropertyFilter) {
		return u.OfEnrolledArgumentPropertyFilter
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIContactFlowCreateRequestDataSourceUnionParam) GetStaticValue() *string {
	if vt := u.OfStaticPropertyFilter; vt != nil {
		return &vt.StaticValue
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIContactFlowCreateRequestDataSourceUnionParam) GetRecordFieldName() *string {
	if vt := u.OfEnrolledRecordPropertyFilter; vt != nil {
		return &vt.RecordFieldName
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIContactFlowCreateRequestDataSourceUnionParam) GetDatasetFieldName() *string {
	if vt := u.OfDatasetFieldPropertyFilter; vt != nil {
		return &vt.DatasetFieldName
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIContactFlowCreateRequestDataSourceUnionParam) GetArgumentName() *string {
	if vt := u.OfEnrolledArgumentPropertyFilter; vt != nil {
		return &vt.ArgumentName
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIContactFlowCreateRequestDataSourceUnionParam) GetAssociationCategory() *string {
	if vt := u.OfAssociation; vt != nil {
		return (*string)(&vt.AssociationCategory)
	} else if vt := u.OfAssociationTimestamp; vt != nil {
		return (*string)(&vt.AssociationCategory)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIContactFlowCreateRequestDataSourceUnionParam) GetAssociationTypeID() *int64 {
	if vt := u.OfAssociation; vt != nil {
		return (*int64)(&vt.AssociationTypeID)
	} else if vt := u.OfAssociationTimestamp; vt != nil {
		return (*int64)(&vt.AssociationTypeID)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIContactFlowCreateRequestDataSourceUnionParam) GetName() *string {
	if vt := u.OfAssociation; vt != nil {
		return (*string)(&vt.Name)
	} else if vt := u.OfAssociationTimestamp; vt != nil {
		return (*string)(&vt.Name)
	} else if vt := u.OfStaticPropertyFilter; vt != nil {
		return (*string)(&vt.Name)
	} else if vt := u.OfEnrolledRecordPropertyFilter; vt != nil {
		return (*string)(&vt.Name)
	} else if vt := u.OfDatasetFieldPropertyFilter; vt != nil {
		return (*string)(&vt.Name)
	} else if vt := u.OfEnrolledArgumentPropertyFilter; vt != nil {
		return (*string)(&vt.Name)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIContactFlowCreateRequestDataSourceUnionParam) GetObjectTypeID() *string {
	if vt := u.OfAssociation; vt != nil {
		return (*string)(&vt.ObjectTypeID)
	} else if vt := u.OfAssociationTimestamp; vt != nil {
		return (*string)(&vt.ObjectTypeID)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIContactFlowCreateRequestDataSourceUnionParam) GetType() *string {
	if vt := u.OfAssociation; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfAssociationTimestamp; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfStaticPropertyFilter; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfEnrolledRecordPropertyFilter; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfDatasetFieldPropertyFilter; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfEnrolledArgumentPropertyFilter; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIContactFlowCreateRequestDataSourceUnionParam) GetPropertyName() *string {
	if vt := u.OfStaticPropertyFilter; vt != nil {
		return (*string)(&vt.PropertyName)
	} else if vt := u.OfEnrolledRecordPropertyFilter; vt != nil {
		return (*string)(&vt.PropertyName)
	} else if vt := u.OfDatasetFieldPropertyFilter; vt != nil {
		return (*string)(&vt.PropertyName)
	} else if vt := u.OfEnrolledArgumentPropertyFilter; vt != nil {
		return (*string)(&vt.PropertyName)
	}
	return nil
}

// Returns a pointer to the underlying variant's SortBy property, if present.
func (u APIContactFlowCreateRequestDataSourceUnionParam) GetSortBy() *APISortParam {
	if vt := u.OfAssociation; vt != nil {
		return &vt.SortBy
	} else if vt := u.OfStaticPropertyFilter; vt != nil {
		return &vt.SortBy
	} else if vt := u.OfEnrolledRecordPropertyFilter; vt != nil {
		return &vt.SortBy
	} else if vt := u.OfDatasetFieldPropertyFilter; vt != nil {
		return &vt.SortBy
	} else if vt := u.OfEnrolledArgumentPropertyFilter; vt != nil {
		return &vt.SortBy
	}
	return nil
}

type APIContactFlowCreateRequestFlowType string

const (
	APIContactFlowCreateRequestFlowTypeWorkflow  APIContactFlowCreateRequestFlowType = "WORKFLOW"
	APIContactFlowCreateRequestFlowTypeActionSet APIContactFlowCreateRequestFlowType = "ACTION_SET"
	APIContactFlowCreateRequestFlowTypeUnknown   APIContactFlowCreateRequestFlowType = "UNKNOWN"
)

type APIContactFlowCreateRequestType string

const (
	APIContactFlowCreateRequestTypeContactFlow  APIContactFlowCreateRequestType = "CONTACT_FLOW"
	APIContactFlowCreateRequestTypePlatformFlow APIContactFlowCreateRequestType = "PLATFORM_FLOW"
)

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type APIContactFlowCreateRequestEnrollmentCriteriaUnionParam struct {
	OfListBased  *APIListBasedEnrollmentCriteriaParam  `json:",omitzero,inline"`
	OfEventBased *APIEventBasedEnrollmentCriteriaParam `json:",omitzero,inline"`
	OfManual     *APIManualEnrollmentCriteriaParam     `json:",omitzero,inline"`
	paramUnion
}

func (u APIContactFlowCreateRequestEnrollmentCriteriaUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfListBased, u.OfEventBased, u.OfManual)
}
func (u *APIContactFlowCreateRequestEnrollmentCriteriaUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *APIContactFlowCreateRequestEnrollmentCriteriaUnionParam) asAny() any {
	if !param.IsOmitted(u.OfListBased) {
		return u.OfListBased
	} else if !param.IsOmitted(u.OfEventBased) {
		return u.OfEventBased
	} else if !param.IsOmitted(u.OfManual) {
		return u.OfManual
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIContactFlowCreateRequestEnrollmentCriteriaUnionParam) GetListFilterBranch() *APIListBasedEnrollmentCriteriaListFilterBranchUnionParam {
	if vt := u.OfListBased; vt != nil {
		return &vt.ListFilterBranch
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIContactFlowCreateRequestEnrollmentCriteriaUnionParam) GetReEnrollmentTriggersFilterBranches() []APIListBasedEnrollmentCriteriaReEnrollmentTriggersFilterBranchUnionParam {
	if vt := u.OfListBased; vt != nil {
		return vt.ReEnrollmentTriggersFilterBranches
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIContactFlowCreateRequestEnrollmentCriteriaUnionParam) GetUnEnrollObjectsNotMeetingCriteria() *bool {
	if vt := u.OfListBased; vt != nil {
		return &vt.UnEnrollObjectsNotMeetingCriteria
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIContactFlowCreateRequestEnrollmentCriteriaUnionParam) GetEventFilterBranches() []shared.PublicUnifiedEventsFilterBranchParam {
	if vt := u.OfEventBased; vt != nil {
		return vt.EventFilterBranches
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIContactFlowCreateRequestEnrollmentCriteriaUnionParam) GetListMembershipFilterBranches() []APIEventBasedEnrollmentCriteriaListMembershipFilterBranchUnionParam {
	if vt := u.OfEventBased; vt != nil {
		return vt.ListMembershipFilterBranches
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIContactFlowCreateRequestEnrollmentCriteriaUnionParam) GetRefinementCriteria() *APIEventBasedEnrollmentCriteriaRefinementCriteriaUnionParam {
	if vt := u.OfEventBased; vt != nil {
		return &vt.RefinementCriteria
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIContactFlowCreateRequestEnrollmentCriteriaUnionParam) GetShouldReEnroll() *bool {
	if vt := u.OfListBased; vt != nil {
		return (*bool)(&vt.ShouldReEnroll)
	} else if vt := u.OfEventBased; vt != nil {
		return (*bool)(&vt.ShouldReEnroll)
	} else if vt := u.OfManual; vt != nil {
		return (*bool)(&vt.ShouldReEnroll)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIContactFlowCreateRequestEnrollmentCriteriaUnionParam) GetType() *string {
	if vt := u.OfListBased; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfEventBased; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfManual; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type APIContactFlowCreateRequestEnrollmentScheduleUnionParam struct {
	OfDaily               *APIDailyEnrollmentScheduleParam               `json:",omitzero,inline"`
	OfWeekly              *APIWeeklyEnrollmentScheduleParam              `json:",omitzero,inline"`
	OfMonthlySpecificDays *APIMonthlySpecificDaysEnrollmentScheduleParam `json:",omitzero,inline"`
	OfMonthlyRelativeDays *APIMonthlyRelativeDaysEnrollmentScheduleParam `json:",omitzero,inline"`
	OfYearly              *APIYearlyEnrollmentScheduleParam              `json:",omitzero,inline"`
	OfPropertyBased       *APIPropertyBasedEnrollmentScheduleParam       `json:",omitzero,inline"`
	paramUnion
}

func (u APIContactFlowCreateRequestEnrollmentScheduleUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfDaily,
		u.OfWeekly,
		u.OfMonthlySpecificDays,
		u.OfMonthlyRelativeDays,
		u.OfYearly,
		u.OfPropertyBased)
}
func (u *APIContactFlowCreateRequestEnrollmentScheduleUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *APIContactFlowCreateRequestEnrollmentScheduleUnionParam) asAny() any {
	if !param.IsOmitted(u.OfDaily) {
		return u.OfDaily
	} else if !param.IsOmitted(u.OfWeekly) {
		return u.OfWeekly
	} else if !param.IsOmitted(u.OfMonthlySpecificDays) {
		return u.OfMonthlySpecificDays
	} else if !param.IsOmitted(u.OfMonthlyRelativeDays) {
		return u.OfMonthlyRelativeDays
	} else if !param.IsOmitted(u.OfYearly) {
		return u.OfYearly
	} else if !param.IsOmitted(u.OfPropertyBased) {
		return u.OfPropertyBased
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIContactFlowCreateRequestEnrollmentScheduleUnionParam) GetDaysOfWeek() []string {
	if vt := u.OfWeekly; vt != nil {
		return vt.DaysOfWeek
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIContactFlowCreateRequestEnrollmentScheduleUnionParam) GetDaysOfMonth() []int64 {
	if vt := u.OfMonthlySpecificDays; vt != nil {
		return vt.DaysOfMonth
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIContactFlowCreateRequestEnrollmentScheduleUnionParam) GetMonthlyRelativeDays() *string {
	if vt := u.OfMonthlyRelativeDays; vt != nil {
		return (*string)(&vt.MonthlyRelativeDays)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIContactFlowCreateRequestEnrollmentScheduleUnionParam) GetDayOfMonth() *int64 {
	if vt := u.OfYearly; vt != nil {
		return &vt.DayOfMonth
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIContactFlowCreateRequestEnrollmentScheduleUnionParam) GetMonth() *string {
	if vt := u.OfYearly; vt != nil {
		return (*string)(&vt.Month)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIContactFlowCreateRequestEnrollmentScheduleUnionParam) GetDateProperty() *string {
	if vt := u.OfPropertyBased; vt != nil {
		return &vt.DateProperty
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIContactFlowCreateRequestEnrollmentScheduleUnionParam) GetDaysDelta() *int64 {
	if vt := u.OfPropertyBased; vt != nil {
		return &vt.DaysDelta
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIContactFlowCreateRequestEnrollmentScheduleUnionParam) GetYearly() *bool {
	if vt := u.OfPropertyBased; vt != nil {
		return &vt.Yearly
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIContactFlowCreateRequestEnrollmentScheduleUnionParam) GetType() *string {
	if vt := u.OfDaily; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfWeekly; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfMonthlySpecificDays; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfMonthlyRelativeDays; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfYearly; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfPropertyBased; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

// Returns a pointer to the underlying variant's TimeOfDay property, if present.
func (u APIContactFlowCreateRequestEnrollmentScheduleUnionParam) GetTimeOfDay() *APITimeOfDayParam {
	if vt := u.OfDaily; vt != nil {
		return &vt.TimeOfDay
	} else if vt := u.OfWeekly; vt != nil {
		return &vt.TimeOfDay
	} else if vt := u.OfMonthlySpecificDays; vt != nil {
		return &vt.TimeOfDay
	} else if vt := u.OfMonthlyRelativeDays; vt != nil {
		return &vt.TimeOfDay
	} else if vt := u.OfYearly; vt != nil {
		return &vt.TimeOfDay
	} else if vt := u.OfPropertyBased; vt != nil {
		return &vt.TimeOfDay
	}
	return nil
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type APIContactFlowCreateRequestEventAnchorUnionParam struct {
	OfContactPropertyAnchor *APIContactPropertyAnchorParam `json:",omitzero,inline"`
	OfStaticDateAnchor      *APIStaticDateAnchorParam      `json:",omitzero,inline"`
	paramUnion
}

func (u APIContactFlowCreateRequestEventAnchorUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfContactPropertyAnchor, u.OfStaticDateAnchor)
}
func (u *APIContactFlowCreateRequestEventAnchorUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *APIContactFlowCreateRequestEventAnchorUnionParam) asAny() any {
	if !param.IsOmitted(u.OfContactPropertyAnchor) {
		return u.OfContactPropertyAnchor
	} else if !param.IsOmitted(u.OfStaticDateAnchor) {
		return u.OfStaticDateAnchor
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIContactFlowCreateRequestEventAnchorUnionParam) GetContactProperty() *string {
	if vt := u.OfContactPropertyAnchor; vt != nil {
		return &vt.ContactProperty
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIContactFlowCreateRequestEventAnchorUnionParam) GetDayOfMonth() *int64 {
	if vt := u.OfStaticDateAnchor; vt != nil {
		return &vt.DayOfMonth
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIContactFlowCreateRequestEventAnchorUnionParam) GetMonth() *string {
	if vt := u.OfStaticDateAnchor; vt != nil {
		return (*string)(&vt.Month)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIContactFlowCreateRequestEventAnchorUnionParam) GetYear() *int64 {
	if vt := u.OfStaticDateAnchor; vt != nil && vt.Year.Valid() {
		return &vt.Year.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIContactFlowCreateRequestEventAnchorUnionParam) GetType() *string {
	if vt := u.OfContactPropertyAnchor; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfStaticDateAnchor; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type APIContactFlowCreateRequestGoalFilterBranchUnionParam struct {
	OfOr                  *shared.PublicOrFilterBranchParam                  `json:",omitzero,inline"`
	OfAnd                 *shared.PublicAndFilterBranchParam                 `json:",omitzero,inline"`
	OfNotAll              *shared.PublicNotAllFilterBranchParam              `json:",omitzero,inline"`
	OfNotAny              *shared.PublicNotAnyFilterBranchParam              `json:",omitzero,inline"`
	OfRestricted          *shared.PublicRestrictedFilterBranchParam          `json:",omitzero,inline"`
	OfUnifiedEvents       *shared.PublicUnifiedEventsFilterBranchParam       `json:",omitzero,inline"`
	OfPropertyAssociation *shared.PublicPropertyAssociationFilterBranchParam `json:",omitzero,inline"`
	OfAssociation         *shared.PublicAssociationFilterBranchParam         `json:",omitzero,inline"`
	paramUnion
}

func (u APIContactFlowCreateRequestGoalFilterBranchUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfOr,
		u.OfAnd,
		u.OfNotAll,
		u.OfNotAny,
		u.OfRestricted,
		u.OfUnifiedEvents,
		u.OfPropertyAssociation,
		u.OfAssociation)
}
func (u *APIContactFlowCreateRequestGoalFilterBranchUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *APIContactFlowCreateRequestGoalFilterBranchUnionParam) asAny() any {
	if !param.IsOmitted(u.OfOr) {
		return u.OfOr
	} else if !param.IsOmitted(u.OfAnd) {
		return u.OfAnd
	} else if !param.IsOmitted(u.OfNotAll) {
		return u.OfNotAll
	} else if !param.IsOmitted(u.OfNotAny) {
		return u.OfNotAny
	} else if !param.IsOmitted(u.OfRestricted) {
		return u.OfRestricted
	} else if !param.IsOmitted(u.OfUnifiedEvents) {
		return u.OfUnifiedEvents
	} else if !param.IsOmitted(u.OfPropertyAssociation) {
		return u.OfPropertyAssociation
	} else if !param.IsOmitted(u.OfAssociation) {
		return u.OfAssociation
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIContactFlowCreateRequestGoalFilterBranchUnionParam) GetEventTypeID() *string {
	if vt := u.OfUnifiedEvents; vt != nil {
		return &vt.EventTypeID
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIContactFlowCreateRequestGoalFilterBranchUnionParam) GetCoalescingRefineBy() *shared.PublicUnifiedEventsFilterBranchCoalescingRefineByUnionParam {
	if vt := u.OfUnifiedEvents; vt != nil {
		return &vt.CoalescingRefineBy
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIContactFlowCreateRequestGoalFilterBranchUnionParam) GetPropertyWithObjectID() *string {
	if vt := u.OfPropertyAssociation; vt != nil {
		return &vt.PropertyWithObjectID
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIContactFlowCreateRequestGoalFilterBranchUnionParam) GetAssociationCategory() *string {
	if vt := u.OfAssociation; vt != nil {
		return &vt.AssociationCategory
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIContactFlowCreateRequestGoalFilterBranchUnionParam) GetAssociationTypeID() *int64 {
	if vt := u.OfAssociation; vt != nil {
		return &vt.AssociationTypeID
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIContactFlowCreateRequestGoalFilterBranchUnionParam) GetFilterBranchOperator() *string {
	if vt := u.OfOr; vt != nil {
		return (*string)(&vt.FilterBranchOperator)
	} else if vt := u.OfAnd; vt != nil {
		return (*string)(&vt.FilterBranchOperator)
	} else if vt := u.OfNotAll; vt != nil {
		return (*string)(&vt.FilterBranchOperator)
	} else if vt := u.OfNotAny; vt != nil {
		return (*string)(&vt.FilterBranchOperator)
	} else if vt := u.OfRestricted; vt != nil {
		return (*string)(&vt.FilterBranchOperator)
	} else if vt := u.OfUnifiedEvents; vt != nil {
		return (*string)(&vt.FilterBranchOperator)
	} else if vt := u.OfPropertyAssociation; vt != nil {
		return (*string)(&vt.FilterBranchOperator)
	} else if vt := u.OfAssociation; vt != nil {
		return (*string)(&vt.FilterBranchOperator)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIContactFlowCreateRequestGoalFilterBranchUnionParam) GetFilterBranchType() *string {
	if vt := u.OfOr; vt != nil {
		return (*string)(&vt.FilterBranchType)
	} else if vt := u.OfAnd; vt != nil {
		return (*string)(&vt.FilterBranchType)
	} else if vt := u.OfNotAll; vt != nil {
		return (*string)(&vt.FilterBranchType)
	} else if vt := u.OfNotAny; vt != nil {
		return (*string)(&vt.FilterBranchType)
	} else if vt := u.OfRestricted; vt != nil {
		return (*string)(&vt.FilterBranchType)
	} else if vt := u.OfUnifiedEvents; vt != nil {
		return (*string)(&vt.FilterBranchType)
	} else if vt := u.OfPropertyAssociation; vt != nil {
		return (*string)(&vt.FilterBranchType)
	} else if vt := u.OfAssociation; vt != nil {
		return (*string)(&vt.FilterBranchType)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIContactFlowCreateRequestGoalFilterBranchUnionParam) GetOperator() *string {
	if vt := u.OfUnifiedEvents; vt != nil {
		return (*string)(&vt.Operator)
	} else if vt := u.OfPropertyAssociation; vt != nil {
		return (*string)(&vt.Operator)
	} else if vt := u.OfAssociation; vt != nil {
		return (*string)(&vt.Operator)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIContactFlowCreateRequestGoalFilterBranchUnionParam) GetObjectTypeID() *string {
	if vt := u.OfPropertyAssociation; vt != nil {
		return (*string)(&vt.ObjectTypeID)
	} else if vt := u.OfAssociation; vt != nil {
		return (*string)(&vt.ObjectTypeID)
	}
	return nil
}

// Returns a subunion which exports methods to access subproperties
//
// Or use AsAny() to get the underlying value
func (u APIContactFlowCreateRequestGoalFilterBranchUnionParam) GetFilterBranches() (res apiContactFlowCreateRequestGoalFilterBranchUnionParamFilterBranches) {
	if vt := u.OfOr; vt != nil {
		res.any = &vt.FilterBranches
	} else if vt := u.OfAnd; vt != nil {
		res.any = &vt.FilterBranches
	} else if vt := u.OfNotAll; vt != nil {
		res.any = &vt.FilterBranches
	} else if vt := u.OfNotAny; vt != nil {
		res.any = &vt.FilterBranches
	} else if vt := u.OfRestricted; vt != nil {
		res.any = &vt.FilterBranches
	} else if vt := u.OfUnifiedEvents; vt != nil {
		res.any = &vt.FilterBranches
	} else if vt := u.OfPropertyAssociation; vt != nil {
		res.any = &vt.FilterBranches
	} else if vt := u.OfAssociation; vt != nil {
		res.any = &vt.FilterBranches
	}
	return
}

// Can have the runtime types
// [_[]shared.PublicOrFilterBranchFilterBranchUnionParam],
// [_[]shared.PublicAndFilterBranchFilterBranchUnionParam],
// [_[]shared.PublicNotAllFilterBranchFilterBranchUnionParam],
// [_[]shared.PublicNotAnyFilterBranchFilterBranchUnionParam],
// [_[]shared.PublicRestrictedFilterBranchFilterBranchUnionParam],
// [_[]shared.PublicUnifiedEventsFilterBranchFilterBranchUnionParam],
// [_[]shared.PublicPropertyAssociationFilterBranchFilterBranchUnionParam],
// [_[]shared.PublicAssociationFilterBranchFilterBranchUnionParam]
type apiContactFlowCreateRequestGoalFilterBranchUnionParamFilterBranches struct{ any }

// Use the following switch statement to get the type of the union:
//
//	switch u.AsAny().(type) {
//	case *[]shared.PublicOrFilterBranchFilterBranchUnionParam:
//	case *[]shared.PublicAndFilterBranchFilterBranchUnionParam:
//	case *[]shared.PublicNotAllFilterBranchFilterBranchUnionParam:
//	case *[]shared.PublicNotAnyFilterBranchFilterBranchUnionParam:
//	case *[]shared.PublicRestrictedFilterBranchFilterBranchUnionParam:
//	case *[]shared.PublicUnifiedEventsFilterBranchFilterBranchUnionParam:
//	case *[]shared.PublicPropertyAssociationFilterBranchFilterBranchUnionParam:
//	case *[]shared.PublicAssociationFilterBranchFilterBranchUnionParam:
//	default:
//	    fmt.Errorf("not present")
//	}
func (u apiContactFlowCreateRequestGoalFilterBranchUnionParamFilterBranches) AsAny() any {
	return u.any
}

// Returns a subunion which exports methods to access subproperties
//
// Or use AsAny() to get the underlying value
func (u APIContactFlowCreateRequestGoalFilterBranchUnionParam) GetFilters() (res apiContactFlowCreateRequestGoalFilterBranchUnionParamFilters) {
	if vt := u.OfOr; vt != nil {
		res.any = &vt.Filters
	} else if vt := u.OfAnd; vt != nil {
		res.any = &vt.Filters
	} else if vt := u.OfNotAll; vt != nil {
		res.any = &vt.Filters
	} else if vt := u.OfNotAny; vt != nil {
		res.any = &vt.Filters
	} else if vt := u.OfRestricted; vt != nil {
		res.any = &vt.Filters
	} else if vt := u.OfUnifiedEvents; vt != nil {
		res.any = &vt.Filters
	} else if vt := u.OfPropertyAssociation; vt != nil {
		res.any = &vt.Filters
	} else if vt := u.OfAssociation; vt != nil {
		res.any = &vt.Filters
	}
	return
}

// Can have the runtime types [_[]shared.PublicOrFilterBranchFilterUnionParam],
// [_[]shared.PublicAndFilterBranchFilterUnionParam],
// [_[]shared.PublicNotAllFilterBranchFilterUnionParam],
// [_[]shared.PublicNotAnyFilterBranchFilterUnionParam],
// [_[]shared.PublicRestrictedFilterBranchFilterUnionParam],
// [_[]shared.PublicUnifiedEventsFilterBranchFilterUnionParam],
// [_[]shared.PublicPropertyAssociationFilterBranchFilterUnionParam],
// [_[]shared.PublicAssociationFilterBranchFilterUnionParam]
type apiContactFlowCreateRequestGoalFilterBranchUnionParamFilters struct{ any }

// Use the following switch statement to get the type of the union:
//
//	switch u.AsAny().(type) {
//	case *[]shared.PublicOrFilterBranchFilterUnionParam:
//	case *[]shared.PublicAndFilterBranchFilterUnionParam:
//	case *[]shared.PublicNotAllFilterBranchFilterUnionParam:
//	case *[]shared.PublicNotAnyFilterBranchFilterUnionParam:
//	case *[]shared.PublicRestrictedFilterBranchFilterUnionParam:
//	case *[]shared.PublicUnifiedEventsFilterBranchFilterUnionParam:
//	case *[]shared.PublicPropertyAssociationFilterBranchFilterUnionParam:
//	case *[]shared.PublicAssociationFilterBranchFilterUnionParam:
//	default:
//	    fmt.Errorf("not present")
//	}
func (u apiContactFlowCreateRequestGoalFilterBranchUnionParamFilters) AsAny() any { return u.any }

// The properties Actions, BlockedDates, CanEnrollFromSalesforce, CustomProperties,
// IsEnabled, RevisionID, SuppressionListIDs, TimeWindows, Type are required.
type APIContactFlowPutRequestParam struct {
	Actions                 []APIContactFlowPutRequestActionUnionParam `json:"actions,omitzero,required"`
	BlockedDates            []APIBlockedDateParam                      `json:"blockedDates,omitzero,required"`
	CanEnrollFromSalesforce bool                                       `json:"canEnrollFromSalesforce,required"`
	CustomProperties        map[string]string                          `json:"customProperties,omitzero,required"`
	IsEnabled               bool                                       `json:"isEnabled,required"`
	RevisionID              string                                     `json:"revisionId,required"`
	SuppressionListIDs      []int64                                    `json:"suppressionListIds,omitzero,required"`
	TimeWindows             []APITimeWindowParam                       `json:"timeWindows,omitzero,required"`
	// Any of "CONTACT_FLOW", "PLATFORM_FLOW".
	Type                APIContactFlowPutRequestType                         `json:"type,omitzero,required"`
	Description         param.Opt[string]                                    `json:"description,omitzero"`
	Name                param.Opt[string]                                    `json:"name,omitzero"`
	StartActionID       param.Opt[string]                                    `json:"startActionId,omitzero"`
	Uuid                param.Opt[string]                                    `json:"uuid,omitzero"`
	EnrollmentCriteria  APIContactFlowPutRequestEnrollmentCriteriaUnionParam `json:"enrollmentCriteria,omitzero"`
	EnrollmentSchedule  APIContactFlowPutRequestEnrollmentScheduleUnionParam `json:"enrollmentSchedule,omitzero"`
	EventAnchor         APIContactFlowPutRequestEventAnchorUnionParam        `json:"eventAnchor,omitzero"`
	GoalFilterBranch    APIContactFlowPutRequestGoalFilterBranchUnionParam   `json:"goalFilterBranch,omitzero"`
	UnEnrollmentSetting APIUnEnrollmentSettingParam                          `json:"unEnrollmentSetting,omitzero"`
	paramObj
}

func (r APIContactFlowPutRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow APIContactFlowPutRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *APIContactFlowPutRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type APIContactFlowPutRequestActionUnionParam struct {
	OfStaticBranch     *APIStaticBranchActionParam     `json:",omitzero,inline"`
	OfListBranch       *APIListBranchActionParam       `json:",omitzero,inline"`
	OfAbTestBranch     *APIAbTestBranchActionParam     `json:",omitzero,inline"`
	OfCustomCode       *APICustomCodeActionParam       `json:",omitzero,inline"`
	OfWebhook          *APIWebhookActionParam          `json:",omitzero,inline"`
	OfSingleConnection *APISingleConnectionActionParam `json:",omitzero,inline"`
	paramUnion
}

func (u APIContactFlowPutRequestActionUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfStaticBranch,
		u.OfListBranch,
		u.OfAbTestBranch,
		u.OfCustomCode,
		u.OfWebhook,
		u.OfSingleConnection)
}
func (u *APIContactFlowPutRequestActionUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *APIContactFlowPutRequestActionUnionParam) asAny() any {
	if !param.IsOmitted(u.OfStaticBranch) {
		return u.OfStaticBranch
	} else if !param.IsOmitted(u.OfListBranch) {
		return u.OfListBranch
	} else if !param.IsOmitted(u.OfAbTestBranch) {
		return u.OfAbTestBranch
	} else if !param.IsOmitted(u.OfCustomCode) {
		return u.OfCustomCode
	} else if !param.IsOmitted(u.OfWebhook) {
		return u.OfWebhook
	} else if !param.IsOmitted(u.OfSingleConnection) {
		return u.OfSingleConnection
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIContactFlowPutRequestActionUnionParam) GetInputValue() *APIStaticBranchActionInputValueUnionParam {
	if vt := u.OfStaticBranch; vt != nil {
		return &vt.InputValue
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIContactFlowPutRequestActionUnionParam) GetStaticBranches() []APIStaticBranchParam {
	if vt := u.OfStaticBranch; vt != nil {
		return vt.StaticBranches
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIContactFlowPutRequestActionUnionParam) GetListBranches() []APIListBranchParam {
	if vt := u.OfListBranch; vt != nil {
		return vt.ListBranches
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIContactFlowPutRequestActionUnionParam) GetTestBranches() []APIConnectionParam {
	if vt := u.OfAbTestBranch; vt != nil {
		return vt.TestBranches
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIContactFlowPutRequestActionUnionParam) GetInputFields() []APIInputVariableParam {
	if vt := u.OfCustomCode; vt != nil {
		return vt.InputFields
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIContactFlowPutRequestActionUnionParam) GetOutputFields() []APIEnumerationOutputFieldParam {
	if vt := u.OfCustomCode; vt != nil {
		return vt.OutputFields
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIContactFlowPutRequestActionUnionParam) GetRuntime() *string {
	if vt := u.OfCustomCode; vt != nil {
		return &vt.Runtime
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIContactFlowPutRequestActionUnionParam) GetSecretNames() []string {
	if vt := u.OfCustomCode; vt != nil {
		return vt.SecretNames
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIContactFlowPutRequestActionUnionParam) GetSourceCode() *string {
	if vt := u.OfCustomCode; vt != nil {
		return &vt.SourceCode
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIContactFlowPutRequestActionUnionParam) GetMethod() *string {
	if vt := u.OfWebhook; vt != nil {
		return (*string)(&vt.Method)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIContactFlowPutRequestActionUnionParam) GetQueryParams() []APIInputVariableParam {
	if vt := u.OfWebhook; vt != nil {
		return vt.QueryParams
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIContactFlowPutRequestActionUnionParam) GetWebhookURL() *string {
	if vt := u.OfWebhook; vt != nil {
		return &vt.WebhookURL
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIContactFlowPutRequestActionUnionParam) GetAuthSettings() *APIWebhookActionAuthSettingsUnionParam {
	if vt := u.OfWebhook; vt != nil {
		return &vt.AuthSettings
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIContactFlowPutRequestActionUnionParam) GetActionTypeID() *string {
	if vt := u.OfSingleConnection; vt != nil {
		return &vt.ActionTypeID
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIContactFlowPutRequestActionUnionParam) GetActionTypeVersion() *int64 {
	if vt := u.OfSingleConnection; vt != nil {
		return &vt.ActionTypeVersion
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIContactFlowPutRequestActionUnionParam) GetFields() map[string]any {
	if vt := u.OfSingleConnection; vt != nil {
		return vt.Fields
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIContactFlowPutRequestActionUnionParam) GetActionID() *string {
	if vt := u.OfStaticBranch; vt != nil {
		return (*string)(&vt.ActionID)
	} else if vt := u.OfListBranch; vt != nil {
		return (*string)(&vt.ActionID)
	} else if vt := u.OfAbTestBranch; vt != nil {
		return (*string)(&vt.ActionID)
	} else if vt := u.OfCustomCode; vt != nil {
		return (*string)(&vt.ActionID)
	} else if vt := u.OfWebhook; vt != nil {
		return (*string)(&vt.ActionID)
	} else if vt := u.OfSingleConnection; vt != nil {
		return (*string)(&vt.ActionID)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIContactFlowPutRequestActionUnionParam) GetType() *string {
	if vt := u.OfStaticBranch; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfListBranch; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfAbTestBranch; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfCustomCode; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfWebhook; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfSingleConnection; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIContactFlowPutRequestActionUnionParam) GetDefaultBranchName() *string {
	if vt := u.OfStaticBranch; vt != nil && vt.DefaultBranchName.Valid() {
		return &vt.DefaultBranchName.Value
	} else if vt := u.OfListBranch; vt != nil && vt.DefaultBranchName.Valid() {
		return &vt.DefaultBranchName.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's DefaultBranch property, if
// present.
func (u APIContactFlowPutRequestActionUnionParam) GetDefaultBranch() *APIConnectionParam {
	if vt := u.OfStaticBranch; vt != nil {
		return &vt.DefaultBranch
	} else if vt := u.OfListBranch; vt != nil {
		return &vt.DefaultBranch
	}
	return nil
}

// Returns a pointer to the underlying variant's Connection property, if present.
func (u APIContactFlowPutRequestActionUnionParam) GetConnection() *APIConnectionParam {
	if vt := u.OfCustomCode; vt != nil {
		return &vt.Connection
	} else if vt := u.OfWebhook; vt != nil {
		return &vt.Connection
	} else if vt := u.OfSingleConnection; vt != nil {
		return &vt.Connection
	}
	return nil
}

type APIContactFlowPutRequestType string

const (
	APIContactFlowPutRequestTypeContactFlow  APIContactFlowPutRequestType = "CONTACT_FLOW"
	APIContactFlowPutRequestTypePlatformFlow APIContactFlowPutRequestType = "PLATFORM_FLOW"
)

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type APIContactFlowPutRequestEnrollmentCriteriaUnionParam struct {
	OfListBased  *APIListBasedEnrollmentCriteriaParam  `json:",omitzero,inline"`
	OfEventBased *APIEventBasedEnrollmentCriteriaParam `json:",omitzero,inline"`
	OfManual     *APIManualEnrollmentCriteriaParam     `json:",omitzero,inline"`
	paramUnion
}

func (u APIContactFlowPutRequestEnrollmentCriteriaUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfListBased, u.OfEventBased, u.OfManual)
}
func (u *APIContactFlowPutRequestEnrollmentCriteriaUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *APIContactFlowPutRequestEnrollmentCriteriaUnionParam) asAny() any {
	if !param.IsOmitted(u.OfListBased) {
		return u.OfListBased
	} else if !param.IsOmitted(u.OfEventBased) {
		return u.OfEventBased
	} else if !param.IsOmitted(u.OfManual) {
		return u.OfManual
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIContactFlowPutRequestEnrollmentCriteriaUnionParam) GetListFilterBranch() *APIListBasedEnrollmentCriteriaListFilterBranchUnionParam {
	if vt := u.OfListBased; vt != nil {
		return &vt.ListFilterBranch
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIContactFlowPutRequestEnrollmentCriteriaUnionParam) GetReEnrollmentTriggersFilterBranches() []APIListBasedEnrollmentCriteriaReEnrollmentTriggersFilterBranchUnionParam {
	if vt := u.OfListBased; vt != nil {
		return vt.ReEnrollmentTriggersFilterBranches
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIContactFlowPutRequestEnrollmentCriteriaUnionParam) GetUnEnrollObjectsNotMeetingCriteria() *bool {
	if vt := u.OfListBased; vt != nil {
		return &vt.UnEnrollObjectsNotMeetingCriteria
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIContactFlowPutRequestEnrollmentCriteriaUnionParam) GetEventFilterBranches() []shared.PublicUnifiedEventsFilterBranchParam {
	if vt := u.OfEventBased; vt != nil {
		return vt.EventFilterBranches
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIContactFlowPutRequestEnrollmentCriteriaUnionParam) GetListMembershipFilterBranches() []APIEventBasedEnrollmentCriteriaListMembershipFilterBranchUnionParam {
	if vt := u.OfEventBased; vt != nil {
		return vt.ListMembershipFilterBranches
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIContactFlowPutRequestEnrollmentCriteriaUnionParam) GetRefinementCriteria() *APIEventBasedEnrollmentCriteriaRefinementCriteriaUnionParam {
	if vt := u.OfEventBased; vt != nil {
		return &vt.RefinementCriteria
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIContactFlowPutRequestEnrollmentCriteriaUnionParam) GetShouldReEnroll() *bool {
	if vt := u.OfListBased; vt != nil {
		return (*bool)(&vt.ShouldReEnroll)
	} else if vt := u.OfEventBased; vt != nil {
		return (*bool)(&vt.ShouldReEnroll)
	} else if vt := u.OfManual; vt != nil {
		return (*bool)(&vt.ShouldReEnroll)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIContactFlowPutRequestEnrollmentCriteriaUnionParam) GetType() *string {
	if vt := u.OfListBased; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfEventBased; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfManual; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type APIContactFlowPutRequestEnrollmentScheduleUnionParam struct {
	OfDaily               *APIDailyEnrollmentScheduleParam               `json:",omitzero,inline"`
	OfWeekly              *APIWeeklyEnrollmentScheduleParam              `json:",omitzero,inline"`
	OfMonthlySpecificDays *APIMonthlySpecificDaysEnrollmentScheduleParam `json:",omitzero,inline"`
	OfMonthlyRelativeDays *APIMonthlyRelativeDaysEnrollmentScheduleParam `json:",omitzero,inline"`
	OfYearly              *APIYearlyEnrollmentScheduleParam              `json:",omitzero,inline"`
	OfPropertyBased       *APIPropertyBasedEnrollmentScheduleParam       `json:",omitzero,inline"`
	paramUnion
}

func (u APIContactFlowPutRequestEnrollmentScheduleUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfDaily,
		u.OfWeekly,
		u.OfMonthlySpecificDays,
		u.OfMonthlyRelativeDays,
		u.OfYearly,
		u.OfPropertyBased)
}
func (u *APIContactFlowPutRequestEnrollmentScheduleUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *APIContactFlowPutRequestEnrollmentScheduleUnionParam) asAny() any {
	if !param.IsOmitted(u.OfDaily) {
		return u.OfDaily
	} else if !param.IsOmitted(u.OfWeekly) {
		return u.OfWeekly
	} else if !param.IsOmitted(u.OfMonthlySpecificDays) {
		return u.OfMonthlySpecificDays
	} else if !param.IsOmitted(u.OfMonthlyRelativeDays) {
		return u.OfMonthlyRelativeDays
	} else if !param.IsOmitted(u.OfYearly) {
		return u.OfYearly
	} else if !param.IsOmitted(u.OfPropertyBased) {
		return u.OfPropertyBased
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIContactFlowPutRequestEnrollmentScheduleUnionParam) GetDaysOfWeek() []string {
	if vt := u.OfWeekly; vt != nil {
		return vt.DaysOfWeek
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIContactFlowPutRequestEnrollmentScheduleUnionParam) GetDaysOfMonth() []int64 {
	if vt := u.OfMonthlySpecificDays; vt != nil {
		return vt.DaysOfMonth
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIContactFlowPutRequestEnrollmentScheduleUnionParam) GetMonthlyRelativeDays() *string {
	if vt := u.OfMonthlyRelativeDays; vt != nil {
		return (*string)(&vt.MonthlyRelativeDays)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIContactFlowPutRequestEnrollmentScheduleUnionParam) GetDayOfMonth() *int64 {
	if vt := u.OfYearly; vt != nil {
		return &vt.DayOfMonth
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIContactFlowPutRequestEnrollmentScheduleUnionParam) GetMonth() *string {
	if vt := u.OfYearly; vt != nil {
		return (*string)(&vt.Month)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIContactFlowPutRequestEnrollmentScheduleUnionParam) GetDateProperty() *string {
	if vt := u.OfPropertyBased; vt != nil {
		return &vt.DateProperty
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIContactFlowPutRequestEnrollmentScheduleUnionParam) GetDaysDelta() *int64 {
	if vt := u.OfPropertyBased; vt != nil {
		return &vt.DaysDelta
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIContactFlowPutRequestEnrollmentScheduleUnionParam) GetYearly() *bool {
	if vt := u.OfPropertyBased; vt != nil {
		return &vt.Yearly
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIContactFlowPutRequestEnrollmentScheduleUnionParam) GetType() *string {
	if vt := u.OfDaily; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfWeekly; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfMonthlySpecificDays; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfMonthlyRelativeDays; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfYearly; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfPropertyBased; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

// Returns a pointer to the underlying variant's TimeOfDay property, if present.
func (u APIContactFlowPutRequestEnrollmentScheduleUnionParam) GetTimeOfDay() *APITimeOfDayParam {
	if vt := u.OfDaily; vt != nil {
		return &vt.TimeOfDay
	} else if vt := u.OfWeekly; vt != nil {
		return &vt.TimeOfDay
	} else if vt := u.OfMonthlySpecificDays; vt != nil {
		return &vt.TimeOfDay
	} else if vt := u.OfMonthlyRelativeDays; vt != nil {
		return &vt.TimeOfDay
	} else if vt := u.OfYearly; vt != nil {
		return &vt.TimeOfDay
	} else if vt := u.OfPropertyBased; vt != nil {
		return &vt.TimeOfDay
	}
	return nil
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type APIContactFlowPutRequestEventAnchorUnionParam struct {
	OfContactPropertyAnchor *APIContactPropertyAnchorParam `json:",omitzero,inline"`
	OfStaticDateAnchor      *APIStaticDateAnchorParam      `json:",omitzero,inline"`
	paramUnion
}

func (u APIContactFlowPutRequestEventAnchorUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfContactPropertyAnchor, u.OfStaticDateAnchor)
}
func (u *APIContactFlowPutRequestEventAnchorUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *APIContactFlowPutRequestEventAnchorUnionParam) asAny() any {
	if !param.IsOmitted(u.OfContactPropertyAnchor) {
		return u.OfContactPropertyAnchor
	} else if !param.IsOmitted(u.OfStaticDateAnchor) {
		return u.OfStaticDateAnchor
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIContactFlowPutRequestEventAnchorUnionParam) GetContactProperty() *string {
	if vt := u.OfContactPropertyAnchor; vt != nil {
		return &vt.ContactProperty
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIContactFlowPutRequestEventAnchorUnionParam) GetDayOfMonth() *int64 {
	if vt := u.OfStaticDateAnchor; vt != nil {
		return &vt.DayOfMonth
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIContactFlowPutRequestEventAnchorUnionParam) GetMonth() *string {
	if vt := u.OfStaticDateAnchor; vt != nil {
		return (*string)(&vt.Month)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIContactFlowPutRequestEventAnchorUnionParam) GetYear() *int64 {
	if vt := u.OfStaticDateAnchor; vt != nil && vt.Year.Valid() {
		return &vt.Year.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIContactFlowPutRequestEventAnchorUnionParam) GetType() *string {
	if vt := u.OfContactPropertyAnchor; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfStaticDateAnchor; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type APIContactFlowPutRequestGoalFilterBranchUnionParam struct {
	OfOr                  *shared.PublicOrFilterBranchParam                  `json:",omitzero,inline"`
	OfAnd                 *shared.PublicAndFilterBranchParam                 `json:",omitzero,inline"`
	OfNotAll              *shared.PublicNotAllFilterBranchParam              `json:",omitzero,inline"`
	OfNotAny              *shared.PublicNotAnyFilterBranchParam              `json:",omitzero,inline"`
	OfRestricted          *shared.PublicRestrictedFilterBranchParam          `json:",omitzero,inline"`
	OfUnifiedEvents       *shared.PublicUnifiedEventsFilterBranchParam       `json:",omitzero,inline"`
	OfPropertyAssociation *shared.PublicPropertyAssociationFilterBranchParam `json:",omitzero,inline"`
	OfAssociation         *shared.PublicAssociationFilterBranchParam         `json:",omitzero,inline"`
	paramUnion
}

func (u APIContactFlowPutRequestGoalFilterBranchUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfOr,
		u.OfAnd,
		u.OfNotAll,
		u.OfNotAny,
		u.OfRestricted,
		u.OfUnifiedEvents,
		u.OfPropertyAssociation,
		u.OfAssociation)
}
func (u *APIContactFlowPutRequestGoalFilterBranchUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *APIContactFlowPutRequestGoalFilterBranchUnionParam) asAny() any {
	if !param.IsOmitted(u.OfOr) {
		return u.OfOr
	} else if !param.IsOmitted(u.OfAnd) {
		return u.OfAnd
	} else if !param.IsOmitted(u.OfNotAll) {
		return u.OfNotAll
	} else if !param.IsOmitted(u.OfNotAny) {
		return u.OfNotAny
	} else if !param.IsOmitted(u.OfRestricted) {
		return u.OfRestricted
	} else if !param.IsOmitted(u.OfUnifiedEvents) {
		return u.OfUnifiedEvents
	} else if !param.IsOmitted(u.OfPropertyAssociation) {
		return u.OfPropertyAssociation
	} else if !param.IsOmitted(u.OfAssociation) {
		return u.OfAssociation
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIContactFlowPutRequestGoalFilterBranchUnionParam) GetEventTypeID() *string {
	if vt := u.OfUnifiedEvents; vt != nil {
		return &vt.EventTypeID
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIContactFlowPutRequestGoalFilterBranchUnionParam) GetCoalescingRefineBy() *shared.PublicUnifiedEventsFilterBranchCoalescingRefineByUnionParam {
	if vt := u.OfUnifiedEvents; vt != nil {
		return &vt.CoalescingRefineBy
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIContactFlowPutRequestGoalFilterBranchUnionParam) GetPropertyWithObjectID() *string {
	if vt := u.OfPropertyAssociation; vt != nil {
		return &vt.PropertyWithObjectID
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIContactFlowPutRequestGoalFilterBranchUnionParam) GetAssociationCategory() *string {
	if vt := u.OfAssociation; vt != nil {
		return &vt.AssociationCategory
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIContactFlowPutRequestGoalFilterBranchUnionParam) GetAssociationTypeID() *int64 {
	if vt := u.OfAssociation; vt != nil {
		return &vt.AssociationTypeID
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIContactFlowPutRequestGoalFilterBranchUnionParam) GetFilterBranchOperator() *string {
	if vt := u.OfOr; vt != nil {
		return (*string)(&vt.FilterBranchOperator)
	} else if vt := u.OfAnd; vt != nil {
		return (*string)(&vt.FilterBranchOperator)
	} else if vt := u.OfNotAll; vt != nil {
		return (*string)(&vt.FilterBranchOperator)
	} else if vt := u.OfNotAny; vt != nil {
		return (*string)(&vt.FilterBranchOperator)
	} else if vt := u.OfRestricted; vt != nil {
		return (*string)(&vt.FilterBranchOperator)
	} else if vt := u.OfUnifiedEvents; vt != nil {
		return (*string)(&vt.FilterBranchOperator)
	} else if vt := u.OfPropertyAssociation; vt != nil {
		return (*string)(&vt.FilterBranchOperator)
	} else if vt := u.OfAssociation; vt != nil {
		return (*string)(&vt.FilterBranchOperator)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIContactFlowPutRequestGoalFilterBranchUnionParam) GetFilterBranchType() *string {
	if vt := u.OfOr; vt != nil {
		return (*string)(&vt.FilterBranchType)
	} else if vt := u.OfAnd; vt != nil {
		return (*string)(&vt.FilterBranchType)
	} else if vt := u.OfNotAll; vt != nil {
		return (*string)(&vt.FilterBranchType)
	} else if vt := u.OfNotAny; vt != nil {
		return (*string)(&vt.FilterBranchType)
	} else if vt := u.OfRestricted; vt != nil {
		return (*string)(&vt.FilterBranchType)
	} else if vt := u.OfUnifiedEvents; vt != nil {
		return (*string)(&vt.FilterBranchType)
	} else if vt := u.OfPropertyAssociation; vt != nil {
		return (*string)(&vt.FilterBranchType)
	} else if vt := u.OfAssociation; vt != nil {
		return (*string)(&vt.FilterBranchType)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIContactFlowPutRequestGoalFilterBranchUnionParam) GetOperator() *string {
	if vt := u.OfUnifiedEvents; vt != nil {
		return (*string)(&vt.Operator)
	} else if vt := u.OfPropertyAssociation; vt != nil {
		return (*string)(&vt.Operator)
	} else if vt := u.OfAssociation; vt != nil {
		return (*string)(&vt.Operator)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIContactFlowPutRequestGoalFilterBranchUnionParam) GetObjectTypeID() *string {
	if vt := u.OfPropertyAssociation; vt != nil {
		return (*string)(&vt.ObjectTypeID)
	} else if vt := u.OfAssociation; vt != nil {
		return (*string)(&vt.ObjectTypeID)
	}
	return nil
}

// Returns a subunion which exports methods to access subproperties
//
// Or use AsAny() to get the underlying value
func (u APIContactFlowPutRequestGoalFilterBranchUnionParam) GetFilterBranches() (res apiContactFlowPutRequestGoalFilterBranchUnionParamFilterBranches) {
	if vt := u.OfOr; vt != nil {
		res.any = &vt.FilterBranches
	} else if vt := u.OfAnd; vt != nil {
		res.any = &vt.FilterBranches
	} else if vt := u.OfNotAll; vt != nil {
		res.any = &vt.FilterBranches
	} else if vt := u.OfNotAny; vt != nil {
		res.any = &vt.FilterBranches
	} else if vt := u.OfRestricted; vt != nil {
		res.any = &vt.FilterBranches
	} else if vt := u.OfUnifiedEvents; vt != nil {
		res.any = &vt.FilterBranches
	} else if vt := u.OfPropertyAssociation; vt != nil {
		res.any = &vt.FilterBranches
	} else if vt := u.OfAssociation; vt != nil {
		res.any = &vt.FilterBranches
	}
	return
}

// Can have the runtime types
// [_[]shared.PublicOrFilterBranchFilterBranchUnionParam],
// [_[]shared.PublicAndFilterBranchFilterBranchUnionParam],
// [_[]shared.PublicNotAllFilterBranchFilterBranchUnionParam],
// [_[]shared.PublicNotAnyFilterBranchFilterBranchUnionParam],
// [_[]shared.PublicRestrictedFilterBranchFilterBranchUnionParam],
// [_[]shared.PublicUnifiedEventsFilterBranchFilterBranchUnionParam],
// [_[]shared.PublicPropertyAssociationFilterBranchFilterBranchUnionParam],
// [_[]shared.PublicAssociationFilterBranchFilterBranchUnionParam]
type apiContactFlowPutRequestGoalFilterBranchUnionParamFilterBranches struct{ any }

// Use the following switch statement to get the type of the union:
//
//	switch u.AsAny().(type) {
//	case *[]shared.PublicOrFilterBranchFilterBranchUnionParam:
//	case *[]shared.PublicAndFilterBranchFilterBranchUnionParam:
//	case *[]shared.PublicNotAllFilterBranchFilterBranchUnionParam:
//	case *[]shared.PublicNotAnyFilterBranchFilterBranchUnionParam:
//	case *[]shared.PublicRestrictedFilterBranchFilterBranchUnionParam:
//	case *[]shared.PublicUnifiedEventsFilterBranchFilterBranchUnionParam:
//	case *[]shared.PublicPropertyAssociationFilterBranchFilterBranchUnionParam:
//	case *[]shared.PublicAssociationFilterBranchFilterBranchUnionParam:
//	default:
//	    fmt.Errorf("not present")
//	}
func (u apiContactFlowPutRequestGoalFilterBranchUnionParamFilterBranches) AsAny() any { return u.any }

// Returns a subunion which exports methods to access subproperties
//
// Or use AsAny() to get the underlying value
func (u APIContactFlowPutRequestGoalFilterBranchUnionParam) GetFilters() (res apiContactFlowPutRequestGoalFilterBranchUnionParamFilters) {
	if vt := u.OfOr; vt != nil {
		res.any = &vt.Filters
	} else if vt := u.OfAnd; vt != nil {
		res.any = &vt.Filters
	} else if vt := u.OfNotAll; vt != nil {
		res.any = &vt.Filters
	} else if vt := u.OfNotAny; vt != nil {
		res.any = &vt.Filters
	} else if vt := u.OfRestricted; vt != nil {
		res.any = &vt.Filters
	} else if vt := u.OfUnifiedEvents; vt != nil {
		res.any = &vt.Filters
	} else if vt := u.OfPropertyAssociation; vt != nil {
		res.any = &vt.Filters
	} else if vt := u.OfAssociation; vt != nil {
		res.any = &vt.Filters
	}
	return
}

// Can have the runtime types [_[]shared.PublicOrFilterBranchFilterUnionParam],
// [_[]shared.PublicAndFilterBranchFilterUnionParam],
// [_[]shared.PublicNotAllFilterBranchFilterUnionParam],
// [_[]shared.PublicNotAnyFilterBranchFilterUnionParam],
// [_[]shared.PublicRestrictedFilterBranchFilterUnionParam],
// [_[]shared.PublicUnifiedEventsFilterBranchFilterUnionParam],
// [_[]shared.PublicPropertyAssociationFilterBranchFilterUnionParam],
// [_[]shared.PublicAssociationFilterBranchFilterUnionParam]
type apiContactFlowPutRequestGoalFilterBranchUnionParamFilters struct{ any }

// Use the following switch statement to get the type of the union:
//
//	switch u.AsAny().(type) {
//	case *[]shared.PublicOrFilterBranchFilterUnionParam:
//	case *[]shared.PublicAndFilterBranchFilterUnionParam:
//	case *[]shared.PublicNotAllFilterBranchFilterUnionParam:
//	case *[]shared.PublicNotAnyFilterBranchFilterUnionParam:
//	case *[]shared.PublicRestrictedFilterBranchFilterUnionParam:
//	case *[]shared.PublicUnifiedEventsFilterBranchFilterUnionParam:
//	case *[]shared.PublicPropertyAssociationFilterBranchFilterUnionParam:
//	case *[]shared.PublicAssociationFilterBranchFilterUnionParam:
//	default:
//	    fmt.Errorf("not present")
//	}
func (u apiContactFlowPutRequestGoalFilterBranchUnionParamFilters) AsAny() any { return u.any }

type APIContactPropertyAnchor struct {
	// A date property on the contact to use as the anchor point of this workflow.
	ContactProperty string `json:"contactProperty,required"`
	// The type of event anchor this is, can be: "CONTACT_PROPERTY_ANCHOR" or
	// "STATIC_DATE_ANCHOR"
	//
	// Any of "CONTACT_PROPERTY_ANCHOR".
	Type APIContactPropertyAnchorType `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ContactProperty respjson.Field
		Type            respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r APIContactPropertyAnchor) RawJSON() string { return r.JSON.raw }
func (r *APIContactPropertyAnchor) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this APIContactPropertyAnchor to a
// APIContactPropertyAnchorParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// APIContactPropertyAnchorParam.Overrides()
func (r APIContactPropertyAnchor) ToParam() APIContactPropertyAnchorParam {
	return param.Override[APIContactPropertyAnchorParam](json.RawMessage(r.RawJSON()))
}

// The type of event anchor this is, can be: "CONTACT_PROPERTY_ANCHOR" or
// "STATIC_DATE_ANCHOR"
type APIContactPropertyAnchorType string

const (
	APIContactPropertyAnchorTypeContactPropertyAnchor APIContactPropertyAnchorType = "CONTACT_PROPERTY_ANCHOR"
)

// The properties ContactProperty, Type are required.
type APIContactPropertyAnchorParam struct {
	// A date property on the contact to use as the anchor point of this workflow.
	ContactProperty string `json:"contactProperty,required"`
	// The type of event anchor this is, can be: "CONTACT_PROPERTY_ANCHOR" or
	// "STATIC_DATE_ANCHOR"
	//
	// Any of "CONTACT_PROPERTY_ANCHOR".
	Type APIContactPropertyAnchorType `json:"type,omitzero,required"`
	paramObj
}

func (r APIContactPropertyAnchorParam) MarshalJSON() (data []byte, err error) {
	type shadow APIContactPropertyAnchorParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *APIContactPropertyAnchorParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type APICustomCodeAction struct {
	// The ID for this action.
	ActionID    string             `json:"actionId,required"`
	InputFields []APIInputVariable `json:"inputFields,required"`
	// The list of output fields that this custom action makes available to the rest of
	// the flow.
	OutputFields []APIEnumerationOutputField `json:"outputFields,required"`
	// The runtime to use to execute the source code. Supported runtimes are:
	// "NODE16X", "NODE20X", "PYTHON39"
	Runtime string `json:"runtime,required"`
	// The names of any "secrets" setup in this portal that will be used in this
	// action.
	SecretNames []string `json:"secretNames,required"`
	// The source code to execute when this action executes.
	SourceCode string `json:"sourceCode,required"`
	// The type of action this is, can be: "STATIC_BRANCH", "LIST_BRANCH",
	// "AB_TEST_BRANCH", "CUSTOM_CODE", "WEBHOOK", or "SINGLE_CONNECTION"
	//
	// Any of "CUSTOM_CODE".
	Type       APICustomCodeActionType `json:"type,required"`
	Connection APIConnection           `json:"connection"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ActionID     respjson.Field
		InputFields  respjson.Field
		OutputFields respjson.Field
		Runtime      respjson.Field
		SecretNames  respjson.Field
		SourceCode   respjson.Field
		Type         respjson.Field
		Connection   respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r APICustomCodeAction) RawJSON() string { return r.JSON.raw }
func (r *APICustomCodeAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this APICustomCodeAction to a APICustomCodeActionParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// APICustomCodeActionParam.Overrides()
func (r APICustomCodeAction) ToParam() APICustomCodeActionParam {
	return param.Override[APICustomCodeActionParam](json.RawMessage(r.RawJSON()))
}

// The type of action this is, can be: "STATIC_BRANCH", "LIST_BRANCH",
// "AB_TEST_BRANCH", "CUSTOM_CODE", "WEBHOOK", or "SINGLE_CONNECTION"
type APICustomCodeActionType string

const (
	APICustomCodeActionTypeCustomCode APICustomCodeActionType = "CUSTOM_CODE"
)

// The properties ActionID, InputFields, OutputFields, Runtime, SecretNames,
// SourceCode, Type are required.
type APICustomCodeActionParam struct {
	// The ID for this action.
	ActionID    string                  `json:"actionId,required"`
	InputFields []APIInputVariableParam `json:"inputFields,omitzero,required"`
	// The list of output fields that this custom action makes available to the rest of
	// the flow.
	OutputFields []APIEnumerationOutputFieldParam `json:"outputFields,omitzero,required"`
	// The runtime to use to execute the source code. Supported runtimes are:
	// "NODE16X", "NODE20X", "PYTHON39"
	Runtime string `json:"runtime,required"`
	// The names of any "secrets" setup in this portal that will be used in this
	// action.
	SecretNames []string `json:"secretNames,omitzero,required"`
	// The source code to execute when this action executes.
	SourceCode string `json:"sourceCode,required"`
	// The type of action this is, can be: "STATIC_BRANCH", "LIST_BRANCH",
	// "AB_TEST_BRANCH", "CUSTOM_CODE", "WEBHOOK", or "SINGLE_CONNECTION"
	//
	// Any of "CUSTOM_CODE".
	Type       APICustomCodeActionType `json:"type,omitzero,required"`
	Connection APIConnectionParam      `json:"connection,omitzero"`
	paramObj
}

func (r APICustomCodeActionParam) MarshalJSON() (data []byte, err error) {
	type shadow APICustomCodeActionParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *APICustomCodeActionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type APIDailyEnrollmentSchedule struct {
	TimeOfDay APITimeOfDay `json:"timeOfDay,required"`
	// The type of enrollment schedule this is, can be: "DAILY", "WEEKLY",
	// "MONTHLY_SPECIFIC_DAYS", "MONTHLY_RELATIVE_DAYS", "YEARLY"
	//
	// Any of "DAILY".
	Type APIDailyEnrollmentScheduleType `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		TimeOfDay   respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r APIDailyEnrollmentSchedule) RawJSON() string { return r.JSON.raw }
func (r *APIDailyEnrollmentSchedule) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this APIDailyEnrollmentSchedule to a
// APIDailyEnrollmentScheduleParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// APIDailyEnrollmentScheduleParam.Overrides()
func (r APIDailyEnrollmentSchedule) ToParam() APIDailyEnrollmentScheduleParam {
	return param.Override[APIDailyEnrollmentScheduleParam](json.RawMessage(r.RawJSON()))
}

// The type of enrollment schedule this is, can be: "DAILY", "WEEKLY",
// "MONTHLY_SPECIFIC_DAYS", "MONTHLY_RELATIVE_DAYS", "YEARLY"
type APIDailyEnrollmentScheduleType string

const (
	APIDailyEnrollmentScheduleTypeDaily APIDailyEnrollmentScheduleType = "DAILY"
)

// The properties TimeOfDay, Type are required.
type APIDailyEnrollmentScheduleParam struct {
	TimeOfDay APITimeOfDayParam `json:"timeOfDay,omitzero,required"`
	// The type of enrollment schedule this is, can be: "DAILY", "WEEKLY",
	// "MONTHLY_SPECIFIC_DAYS", "MONTHLY_RELATIVE_DAYS", "YEARLY"
	//
	// Any of "DAILY".
	Type APIDailyEnrollmentScheduleType `json:"type,omitzero,required"`
	paramObj
}

func (r APIDailyEnrollmentScheduleParam) MarshalJSON() (data []byte, err error) {
	type shadow APIDailyEnrollmentScheduleParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *APIDailyEnrollmentScheduleParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type APIDatasetFieldPropertyFilterDataSource struct {
	DatasetFieldName string `json:"datasetFieldName,required"`
	Name             string `json:"name,required"`
	PropertyName     string `json:"propertyName,required"`
	// Any of "DATASET_FIELD_PROPERTY_FILTER".
	Type   APIDatasetFieldPropertyFilterDataSourceType `json:"type,required"`
	SortBy APISort                                     `json:"sortBy"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DatasetFieldName respjson.Field
		Name             respjson.Field
		PropertyName     respjson.Field
		Type             respjson.Field
		SortBy           respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r APIDatasetFieldPropertyFilterDataSource) RawJSON() string { return r.JSON.raw }
func (r *APIDatasetFieldPropertyFilterDataSource) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this APIDatasetFieldPropertyFilterDataSource to a
// APIDatasetFieldPropertyFilterDataSourceParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// APIDatasetFieldPropertyFilterDataSourceParam.Overrides()
func (r APIDatasetFieldPropertyFilterDataSource) ToParam() APIDatasetFieldPropertyFilterDataSourceParam {
	return param.Override[APIDatasetFieldPropertyFilterDataSourceParam](json.RawMessage(r.RawJSON()))
}

type APIDatasetFieldPropertyFilterDataSourceType string

const (
	APIDatasetFieldPropertyFilterDataSourceTypeDatasetFieldPropertyFilter APIDatasetFieldPropertyFilterDataSourceType = "DATASET_FIELD_PROPERTY_FILTER"
)

// The properties DatasetFieldName, Name, PropertyName, Type are required.
type APIDatasetFieldPropertyFilterDataSourceParam struct {
	DatasetFieldName string `json:"datasetFieldName,required"`
	Name             string `json:"name,required"`
	PropertyName     string `json:"propertyName,required"`
	// Any of "DATASET_FIELD_PROPERTY_FILTER".
	Type   APIDatasetFieldPropertyFilterDataSourceType `json:"type,omitzero,required"`
	SortBy APISortParam                                `json:"sortBy,omitzero"`
	paramObj
}

func (r APIDatasetFieldPropertyFilterDataSourceParam) MarshalJSON() (data []byte, err error) {
	type shadow APIDatasetFieldPropertyFilterDataSourceParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *APIDatasetFieldPropertyFilterDataSourceParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type APIEnrolledArgumentPropertyFilterDataSource struct {
	ArgumentName string `json:"argumentName,required"`
	Name         string `json:"name,required"`
	PropertyName string `json:"propertyName,required"`
	// Any of "ENROLLED_ARGUMENT_PROPERTY_FILTER".
	Type   APIEnrolledArgumentPropertyFilterDataSourceType `json:"type,required"`
	SortBy APISort                                         `json:"sortBy"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ArgumentName respjson.Field
		Name         respjson.Field
		PropertyName respjson.Field
		Type         respjson.Field
		SortBy       respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r APIEnrolledArgumentPropertyFilterDataSource) RawJSON() string { return r.JSON.raw }
func (r *APIEnrolledArgumentPropertyFilterDataSource) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this APIEnrolledArgumentPropertyFilterDataSource to a
// APIEnrolledArgumentPropertyFilterDataSourceParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// APIEnrolledArgumentPropertyFilterDataSourceParam.Overrides()
func (r APIEnrolledArgumentPropertyFilterDataSource) ToParam() APIEnrolledArgumentPropertyFilterDataSourceParam {
	return param.Override[APIEnrolledArgumentPropertyFilterDataSourceParam](json.RawMessage(r.RawJSON()))
}

type APIEnrolledArgumentPropertyFilterDataSourceType string

const (
	APIEnrolledArgumentPropertyFilterDataSourceTypeEnrolledArgumentPropertyFilter APIEnrolledArgumentPropertyFilterDataSourceType = "ENROLLED_ARGUMENT_PROPERTY_FILTER"
)

// The properties ArgumentName, Name, PropertyName, Type are required.
type APIEnrolledArgumentPropertyFilterDataSourceParam struct {
	ArgumentName string `json:"argumentName,required"`
	Name         string `json:"name,required"`
	PropertyName string `json:"propertyName,required"`
	// Any of "ENROLLED_ARGUMENT_PROPERTY_FILTER".
	Type   APIEnrolledArgumentPropertyFilterDataSourceType `json:"type,omitzero,required"`
	SortBy APISortParam                                    `json:"sortBy,omitzero"`
	paramObj
}

func (r APIEnrolledArgumentPropertyFilterDataSourceParam) MarshalJSON() (data []byte, err error) {
	type shadow APIEnrolledArgumentPropertyFilterDataSourceParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *APIEnrolledArgumentPropertyFilterDataSourceParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type APIEnrolledRecordPropertyFilterDataSource struct {
	Name            string `json:"name,required"`
	PropertyName    string `json:"propertyName,required"`
	RecordFieldName string `json:"recordFieldName,required"`
	// Any of "ENROLLED_RECORD_PROPERTY_FILTER".
	Type   APIEnrolledRecordPropertyFilterDataSourceType `json:"type,required"`
	SortBy APISort                                       `json:"sortBy"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name            respjson.Field
		PropertyName    respjson.Field
		RecordFieldName respjson.Field
		Type            respjson.Field
		SortBy          respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r APIEnrolledRecordPropertyFilterDataSource) RawJSON() string { return r.JSON.raw }
func (r *APIEnrolledRecordPropertyFilterDataSource) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this APIEnrolledRecordPropertyFilterDataSource to a
// APIEnrolledRecordPropertyFilterDataSourceParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// APIEnrolledRecordPropertyFilterDataSourceParam.Overrides()
func (r APIEnrolledRecordPropertyFilterDataSource) ToParam() APIEnrolledRecordPropertyFilterDataSourceParam {
	return param.Override[APIEnrolledRecordPropertyFilterDataSourceParam](json.RawMessage(r.RawJSON()))
}

type APIEnrolledRecordPropertyFilterDataSourceType string

const (
	APIEnrolledRecordPropertyFilterDataSourceTypeEnrolledRecordPropertyFilter APIEnrolledRecordPropertyFilterDataSourceType = "ENROLLED_RECORD_PROPERTY_FILTER"
)

// The properties Name, PropertyName, RecordFieldName, Type are required.
type APIEnrolledRecordPropertyFilterDataSourceParam struct {
	Name            string `json:"name,required"`
	PropertyName    string `json:"propertyName,required"`
	RecordFieldName string `json:"recordFieldName,required"`
	// Any of "ENROLLED_RECORD_PROPERTY_FILTER".
	Type   APIEnrolledRecordPropertyFilterDataSourceType `json:"type,omitzero,required"`
	SortBy APISortParam                                  `json:"sortBy,omitzero"`
	paramObj
}

func (r APIEnrolledRecordPropertyFilterDataSourceParam) MarshalJSON() (data []byte, err error) {
	type shadow APIEnrolledRecordPropertyFilterDataSourceParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *APIEnrolledRecordPropertyFilterDataSourceParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type APIEnrollmentEventPropertyValue struct {
	EnrollmentEventPropertyToken string `json:"enrollmentEventPropertyToken,required"`
	// Any of "ENROLLMENT_EVENT_PROPERTY".
	Type APIEnrollmentEventPropertyValueType `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		EnrollmentEventPropertyToken respjson.Field
		Type                         respjson.Field
		ExtraFields                  map[string]respjson.Field
		raw                          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r APIEnrollmentEventPropertyValue) RawJSON() string { return r.JSON.raw }
func (r *APIEnrollmentEventPropertyValue) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this APIEnrollmentEventPropertyValue to a
// APIEnrollmentEventPropertyValueParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// APIEnrollmentEventPropertyValueParam.Overrides()
func (r APIEnrollmentEventPropertyValue) ToParam() APIEnrollmentEventPropertyValueParam {
	return param.Override[APIEnrollmentEventPropertyValueParam](json.RawMessage(r.RawJSON()))
}

type APIEnrollmentEventPropertyValueType string

const (
	APIEnrollmentEventPropertyValueTypeEnrollmentEventProperty APIEnrollmentEventPropertyValueType = "ENROLLMENT_EVENT_PROPERTY"
)

// The properties EnrollmentEventPropertyToken, Type are required.
type APIEnrollmentEventPropertyValueParam struct {
	EnrollmentEventPropertyToken string `json:"enrollmentEventPropertyToken,required"`
	// Any of "ENROLLMENT_EVENT_PROPERTY".
	Type APIEnrollmentEventPropertyValueType `json:"type,omitzero,required"`
	paramObj
}

func (r APIEnrollmentEventPropertyValueParam) MarshalJSON() (data []byte, err error) {
	type shadow APIEnrollmentEventPropertyValueParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *APIEnrollmentEventPropertyValueParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type APIEnumerationOutputField struct {
	Name    string   `json:"name,required"`
	Options []string `json:"options,required"`
	// Any of "ENUMERATION".
	Type APIEnumerationOutputFieldType `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		Options     respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r APIEnumerationOutputField) RawJSON() string { return r.JSON.raw }
func (r *APIEnumerationOutputField) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this APIEnumerationOutputField to a
// APIEnumerationOutputFieldParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// APIEnumerationOutputFieldParam.Overrides()
func (r APIEnumerationOutputField) ToParam() APIEnumerationOutputFieldParam {
	return param.Override[APIEnumerationOutputFieldParam](json.RawMessage(r.RawJSON()))
}

type APIEnumerationOutputFieldType string

const (
	APIEnumerationOutputFieldTypeEnumeration APIEnumerationOutputFieldType = "ENUMERATION"
)

// The properties Name, Options, Type are required.
type APIEnumerationOutputFieldParam struct {
	Name    string   `json:"name,required"`
	Options []string `json:"options,omitzero,required"`
	// Any of "ENUMERATION".
	Type APIEnumerationOutputFieldType `json:"type,omitzero,required"`
	paramObj
}

func (r APIEnumerationOutputFieldParam) MarshalJSON() (data []byte, err error) {
	type shadow APIEnumerationOutputFieldParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *APIEnumerationOutputFieldParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type APIEventBasedEnrollmentCriteria struct {
	EventFilterBranches []shared.PublicUnifiedEventsFilterBranch `json:"eventFilterBranches,required"`
	// If you want to listen to list-membership events (an object was added to a list,
	// an object was removed from a list) you need to use this
	// `listMembershipFilterBranches` property instead of `eventFilterBranches`,
	// because list membership events work differently.
	ListMembershipFilterBranches []APIEventBasedEnrollmentCriteriaListMembershipFilterBranchUnion `json:"listMembershipFilterBranches,required"`
	// Whether or not the same object can enroll in this workflow twice.
	ShouldReEnroll bool `json:"shouldReEnroll,required"`
	// The type of enrollment criteria this is, this can be "LIST_BASED",
	// "EVENT_BASED", or "MANUAL".
	//
	// Any of "EVENT_BASED".
	Type APIEventBasedEnrollmentCriteriaType `json:"type,required"`
	// List-based criteria to further refine which contacts will enroll in this flow.
	RefinementCriteria APIEventBasedEnrollmentCriteriaRefinementCriteriaUnion `json:"refinementCriteria"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		EventFilterBranches          respjson.Field
		ListMembershipFilterBranches respjson.Field
		ShouldReEnroll               respjson.Field
		Type                         respjson.Field
		RefinementCriteria           respjson.Field
		ExtraFields                  map[string]respjson.Field
		raw                          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r APIEventBasedEnrollmentCriteria) RawJSON() string { return r.JSON.raw }
func (r *APIEventBasedEnrollmentCriteria) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this APIEventBasedEnrollmentCriteria to a
// APIEventBasedEnrollmentCriteriaParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// APIEventBasedEnrollmentCriteriaParam.Overrides()
func (r APIEventBasedEnrollmentCriteria) ToParam() APIEventBasedEnrollmentCriteriaParam {
	return param.Override[APIEventBasedEnrollmentCriteriaParam](json.RawMessage(r.RawJSON()))
}

// APIEventBasedEnrollmentCriteriaListMembershipFilterBranchUnion contains all
// possible properties and values from [shared.PublicOrFilterBranch],
// [shared.PublicAndFilterBranch], [shared.PublicNotAllFilterBranch],
// [shared.PublicNotAnyFilterBranch], [shared.PublicRestrictedFilterBranch],
// [shared.PublicUnifiedEventsFilterBranch],
// [shared.PublicPropertyAssociationFilterBranch],
// [shared.PublicAssociationFilterBranch].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type APIEventBasedEnrollmentCriteriaListMembershipFilterBranchUnion struct {
	// This field is a union of [[]shared.PublicOrFilterBranchFilterBranchUnion],
	// [[]shared.PublicAndFilterBranchFilterBranchUnion],
	// [[]shared.PublicNotAllFilterBranchFilterBranchUnion],
	// [[]shared.PublicNotAnyFilterBranchFilterBranchUnion],
	// [[]shared.PublicRestrictedFilterBranchFilterBranchUnion],
	// [[]shared.PublicUnifiedEventsFilterBranchFilterBranchUnion],
	// [[]shared.PublicPropertyAssociationFilterBranchFilterBranchUnion],
	// [[]shared.PublicAssociationFilterBranchFilterBranchUnion]
	FilterBranches       APIEventBasedEnrollmentCriteriaListMembershipFilterBranchUnionFilterBranches `json:"filterBranches"`
	FilterBranchOperator string                                                                       `json:"filterBranchOperator"`
	FilterBranchType     string                                                                       `json:"filterBranchType"`
	// This field is a union of [[]shared.PublicOrFilterBranchFilterUnion],
	// [[]shared.PublicAndFilterBranchFilterUnion],
	// [[]shared.PublicNotAllFilterBranchFilterUnion],
	// [[]shared.PublicNotAnyFilterBranchFilterUnion],
	// [[]shared.PublicRestrictedFilterBranchFilterUnion],
	// [[]shared.PublicUnifiedEventsFilterBranchFilterUnion],
	// [[]shared.PublicPropertyAssociationFilterBranchFilterUnion],
	// [[]shared.PublicAssociationFilterBranchFilterUnion]
	Filters APIEventBasedEnrollmentCriteriaListMembershipFilterBranchUnionFilters `json:"filters"`
	// This field is from variant [shared.PublicUnifiedEventsFilterBranch].
	EventTypeID string `json:"eventTypeId"`
	Operator    string `json:"operator"`
	// This field is from variant [shared.PublicUnifiedEventsFilterBranch].
	CoalescingRefineBy shared.PublicUnifiedEventsFilterBranchCoalescingRefineByUnion `json:"coalescingRefineBy"`
	ObjectTypeID       string                                                        `json:"objectTypeId"`
	// This field is from variant [shared.PublicPropertyAssociationFilterBranch].
	PropertyWithObjectID string `json:"propertyWithObjectId"`
	// This field is from variant [shared.PublicAssociationFilterBranch].
	AssociationCategory string `json:"associationCategory"`
	// This field is from variant [shared.PublicAssociationFilterBranch].
	AssociationTypeID int64 `json:"associationTypeId"`
	JSON              struct {
		FilterBranches       respjson.Field
		FilterBranchOperator respjson.Field
		FilterBranchType     respjson.Field
		Filters              respjson.Field
		EventTypeID          respjson.Field
		Operator             respjson.Field
		CoalescingRefineBy   respjson.Field
		ObjectTypeID         respjson.Field
		PropertyWithObjectID respjson.Field
		AssociationCategory  respjson.Field
		AssociationTypeID    respjson.Field
		raw                  string
	} `json:"-"`
}

func (u APIEventBasedEnrollmentCriteriaListMembershipFilterBranchUnion) AsOr() (v shared.PublicOrFilterBranch) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u APIEventBasedEnrollmentCriteriaListMembershipFilterBranchUnion) AsAnd() (v shared.PublicAndFilterBranch) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u APIEventBasedEnrollmentCriteriaListMembershipFilterBranchUnion) AsNotAll() (v shared.PublicNotAllFilterBranch) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u APIEventBasedEnrollmentCriteriaListMembershipFilterBranchUnion) AsNotAny() (v shared.PublicNotAnyFilterBranch) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u APIEventBasedEnrollmentCriteriaListMembershipFilterBranchUnion) AsRestricted() (v shared.PublicRestrictedFilterBranch) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u APIEventBasedEnrollmentCriteriaListMembershipFilterBranchUnion) AsUnifiedEvents() (v shared.PublicUnifiedEventsFilterBranch) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u APIEventBasedEnrollmentCriteriaListMembershipFilterBranchUnion) AsPropertyAssociation() (v shared.PublicPropertyAssociationFilterBranch) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u APIEventBasedEnrollmentCriteriaListMembershipFilterBranchUnion) AsAssociation() (v shared.PublicAssociationFilterBranch) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u APIEventBasedEnrollmentCriteriaListMembershipFilterBranchUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *APIEventBasedEnrollmentCriteriaListMembershipFilterBranchUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// APIEventBasedEnrollmentCriteriaListMembershipFilterBranchUnionFilterBranches is
// an implicit subunion of
// [APIEventBasedEnrollmentCriteriaListMembershipFilterBranchUnion].
// APIEventBasedEnrollmentCriteriaListMembershipFilterBranchUnionFilterBranches
// provides convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [APIEventBasedEnrollmentCriteriaListMembershipFilterBranchUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfPublicOrFilterBranchFilterBranches
// OfPublicAndFilterBranchFilterBranches OfPublicNotAllFilterBranchFilterBranches
// OfPublicNotAnyFilterBranchFilterBranches
// OfPublicRestrictedFilterBranchFilterBranches
// OfPublicUnifiedEventsFilterBranchFilterBranches
// OfPublicPropertyAssociationFilterBranchFilterBranches
// OfPublicAssociationFilterBranchFilterBranches]
type APIEventBasedEnrollmentCriteriaListMembershipFilterBranchUnionFilterBranches struct {
	// This field will be present if the value is a
	// [[]shared.PublicOrFilterBranchFilterBranchUnion] instead of an object.
	OfPublicOrFilterBranchFilterBranches []shared.PublicOrFilterBranchFilterBranchUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]shared.PublicAndFilterBranchFilterBranchUnion] instead of an object.
	OfPublicAndFilterBranchFilterBranches []shared.PublicAndFilterBranchFilterBranchUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]shared.PublicNotAllFilterBranchFilterBranchUnion] instead of an object.
	OfPublicNotAllFilterBranchFilterBranches []shared.PublicNotAllFilterBranchFilterBranchUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]shared.PublicNotAnyFilterBranchFilterBranchUnion] instead of an object.
	OfPublicNotAnyFilterBranchFilterBranches []shared.PublicNotAnyFilterBranchFilterBranchUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]shared.PublicRestrictedFilterBranchFilterBranchUnion] instead of an object.
	OfPublicRestrictedFilterBranchFilterBranches []shared.PublicRestrictedFilterBranchFilterBranchUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]shared.PublicUnifiedEventsFilterBranchFilterBranchUnion] instead of an
	// object.
	OfPublicUnifiedEventsFilterBranchFilterBranches []shared.PublicUnifiedEventsFilterBranchFilterBranchUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]shared.PublicPropertyAssociationFilterBranchFilterBranchUnion] instead of an
	// object.
	OfPublicPropertyAssociationFilterBranchFilterBranches []shared.PublicPropertyAssociationFilterBranchFilterBranchUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]shared.PublicAssociationFilterBranchFilterBranchUnion] instead of an object.
	OfPublicAssociationFilterBranchFilterBranches []shared.PublicAssociationFilterBranchFilterBranchUnion `json:",inline"`
	JSON                                          struct {
		OfPublicOrFilterBranchFilterBranches                  respjson.Field
		OfPublicAndFilterBranchFilterBranches                 respjson.Field
		OfPublicNotAllFilterBranchFilterBranches              respjson.Field
		OfPublicNotAnyFilterBranchFilterBranches              respjson.Field
		OfPublicRestrictedFilterBranchFilterBranches          respjson.Field
		OfPublicUnifiedEventsFilterBranchFilterBranches       respjson.Field
		OfPublicPropertyAssociationFilterBranchFilterBranches respjson.Field
		OfPublicAssociationFilterBranchFilterBranches         respjson.Field
		raw                                                   string
	} `json:"-"`
}

func (r *APIEventBasedEnrollmentCriteriaListMembershipFilterBranchUnionFilterBranches) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// APIEventBasedEnrollmentCriteriaListMembershipFilterBranchUnionFilters is an
// implicit subunion of
// [APIEventBasedEnrollmentCriteriaListMembershipFilterBranchUnion].
// APIEventBasedEnrollmentCriteriaListMembershipFilterBranchUnionFilters provides
// convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [APIEventBasedEnrollmentCriteriaListMembershipFilterBranchUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfPublicOrFilterBranchFilters OfPublicAndFilterBranchFilters
// OfPublicNotAllFilterBranchFilters OfPublicNotAnyFilterBranchFilters
// OfPublicRestrictedFilterBranchFilters OfPublicUnifiedEventsFilterBranchFilters
// OfPublicPropertyAssociationFilterBranchFilters
// OfPublicAssociationFilterBranchFilters]
type APIEventBasedEnrollmentCriteriaListMembershipFilterBranchUnionFilters struct {
	// This field will be present if the value is a
	// [[]shared.PublicOrFilterBranchFilterUnion] instead of an object.
	OfPublicOrFilterBranchFilters []shared.PublicOrFilterBranchFilterUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]shared.PublicAndFilterBranchFilterUnion] instead of an object.
	OfPublicAndFilterBranchFilters []shared.PublicAndFilterBranchFilterUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]shared.PublicNotAllFilterBranchFilterUnion] instead of an object.
	OfPublicNotAllFilterBranchFilters []shared.PublicNotAllFilterBranchFilterUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]shared.PublicNotAnyFilterBranchFilterUnion] instead of an object.
	OfPublicNotAnyFilterBranchFilters []shared.PublicNotAnyFilterBranchFilterUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]shared.PublicRestrictedFilterBranchFilterUnion] instead of an object.
	OfPublicRestrictedFilterBranchFilters []shared.PublicRestrictedFilterBranchFilterUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]shared.PublicUnifiedEventsFilterBranchFilterUnion] instead of an object.
	OfPublicUnifiedEventsFilterBranchFilters []shared.PublicUnifiedEventsFilterBranchFilterUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]shared.PublicPropertyAssociationFilterBranchFilterUnion] instead of an
	// object.
	OfPublicPropertyAssociationFilterBranchFilters []shared.PublicPropertyAssociationFilterBranchFilterUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]shared.PublicAssociationFilterBranchFilterUnion] instead of an object.
	OfPublicAssociationFilterBranchFilters []shared.PublicAssociationFilterBranchFilterUnion `json:",inline"`
	JSON                                   struct {
		OfPublicOrFilterBranchFilters                  respjson.Field
		OfPublicAndFilterBranchFilters                 respjson.Field
		OfPublicNotAllFilterBranchFilters              respjson.Field
		OfPublicNotAnyFilterBranchFilters              respjson.Field
		OfPublicRestrictedFilterBranchFilters          respjson.Field
		OfPublicUnifiedEventsFilterBranchFilters       respjson.Field
		OfPublicPropertyAssociationFilterBranchFilters respjson.Field
		OfPublicAssociationFilterBranchFilters         respjson.Field
		raw                                            string
	} `json:"-"`
}

func (r *APIEventBasedEnrollmentCriteriaListMembershipFilterBranchUnionFilters) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The type of enrollment criteria this is, this can be "LIST_BASED",
// "EVENT_BASED", or "MANUAL".
type APIEventBasedEnrollmentCriteriaType string

const (
	APIEventBasedEnrollmentCriteriaTypeEventBased APIEventBasedEnrollmentCriteriaType = "EVENT_BASED"
)

// APIEventBasedEnrollmentCriteriaRefinementCriteriaUnion contains all possible
// properties and values from [shared.PublicOrFilterBranch],
// [shared.PublicAndFilterBranch], [shared.PublicNotAllFilterBranch],
// [shared.PublicNotAnyFilterBranch], [shared.PublicRestrictedFilterBranch],
// [shared.PublicUnifiedEventsFilterBranch],
// [shared.PublicPropertyAssociationFilterBranch],
// [shared.PublicAssociationFilterBranch].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type APIEventBasedEnrollmentCriteriaRefinementCriteriaUnion struct {
	// This field is a union of [[]shared.PublicOrFilterBranchFilterBranchUnion],
	// [[]shared.PublicAndFilterBranchFilterBranchUnion],
	// [[]shared.PublicNotAllFilterBranchFilterBranchUnion],
	// [[]shared.PublicNotAnyFilterBranchFilterBranchUnion],
	// [[]shared.PublicRestrictedFilterBranchFilterBranchUnion],
	// [[]shared.PublicUnifiedEventsFilterBranchFilterBranchUnion],
	// [[]shared.PublicPropertyAssociationFilterBranchFilterBranchUnion],
	// [[]shared.PublicAssociationFilterBranchFilterBranchUnion]
	FilterBranches       APIEventBasedEnrollmentCriteriaRefinementCriteriaUnionFilterBranches `json:"filterBranches"`
	FilterBranchOperator string                                                               `json:"filterBranchOperator"`
	FilterBranchType     string                                                               `json:"filterBranchType"`
	// This field is a union of [[]shared.PublicOrFilterBranchFilterUnion],
	// [[]shared.PublicAndFilterBranchFilterUnion],
	// [[]shared.PublicNotAllFilterBranchFilterUnion],
	// [[]shared.PublicNotAnyFilterBranchFilterUnion],
	// [[]shared.PublicRestrictedFilterBranchFilterUnion],
	// [[]shared.PublicUnifiedEventsFilterBranchFilterUnion],
	// [[]shared.PublicPropertyAssociationFilterBranchFilterUnion],
	// [[]shared.PublicAssociationFilterBranchFilterUnion]
	Filters APIEventBasedEnrollmentCriteriaRefinementCriteriaUnionFilters `json:"filters"`
	// This field is from variant [shared.PublicUnifiedEventsFilterBranch].
	EventTypeID string `json:"eventTypeId"`
	Operator    string `json:"operator"`
	// This field is from variant [shared.PublicUnifiedEventsFilterBranch].
	CoalescingRefineBy shared.PublicUnifiedEventsFilterBranchCoalescingRefineByUnion `json:"coalescingRefineBy"`
	ObjectTypeID       string                                                        `json:"objectTypeId"`
	// This field is from variant [shared.PublicPropertyAssociationFilterBranch].
	PropertyWithObjectID string `json:"propertyWithObjectId"`
	// This field is from variant [shared.PublicAssociationFilterBranch].
	AssociationCategory string `json:"associationCategory"`
	// This field is from variant [shared.PublicAssociationFilterBranch].
	AssociationTypeID int64 `json:"associationTypeId"`
	JSON              struct {
		FilterBranches       respjson.Field
		FilterBranchOperator respjson.Field
		FilterBranchType     respjson.Field
		Filters              respjson.Field
		EventTypeID          respjson.Field
		Operator             respjson.Field
		CoalescingRefineBy   respjson.Field
		ObjectTypeID         respjson.Field
		PropertyWithObjectID respjson.Field
		AssociationCategory  respjson.Field
		AssociationTypeID    respjson.Field
		raw                  string
	} `json:"-"`
}

func (u APIEventBasedEnrollmentCriteriaRefinementCriteriaUnion) AsOr() (v shared.PublicOrFilterBranch) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u APIEventBasedEnrollmentCriteriaRefinementCriteriaUnion) AsAnd() (v shared.PublicAndFilterBranch) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u APIEventBasedEnrollmentCriteriaRefinementCriteriaUnion) AsNotAll() (v shared.PublicNotAllFilterBranch) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u APIEventBasedEnrollmentCriteriaRefinementCriteriaUnion) AsNotAny() (v shared.PublicNotAnyFilterBranch) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u APIEventBasedEnrollmentCriteriaRefinementCriteriaUnion) AsRestricted() (v shared.PublicRestrictedFilterBranch) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u APIEventBasedEnrollmentCriteriaRefinementCriteriaUnion) AsUnifiedEvents() (v shared.PublicUnifiedEventsFilterBranch) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u APIEventBasedEnrollmentCriteriaRefinementCriteriaUnion) AsPropertyAssociation() (v shared.PublicPropertyAssociationFilterBranch) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u APIEventBasedEnrollmentCriteriaRefinementCriteriaUnion) AsAssociation() (v shared.PublicAssociationFilterBranch) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u APIEventBasedEnrollmentCriteriaRefinementCriteriaUnion) RawJSON() string { return u.JSON.raw }

func (r *APIEventBasedEnrollmentCriteriaRefinementCriteriaUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// APIEventBasedEnrollmentCriteriaRefinementCriteriaUnionFilterBranches is an
// implicit subunion of [APIEventBasedEnrollmentCriteriaRefinementCriteriaUnion].
// APIEventBasedEnrollmentCriteriaRefinementCriteriaUnionFilterBranches provides
// convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [APIEventBasedEnrollmentCriteriaRefinementCriteriaUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfPublicOrFilterBranchFilterBranches
// OfPublicAndFilterBranchFilterBranches OfPublicNotAllFilterBranchFilterBranches
// OfPublicNotAnyFilterBranchFilterBranches
// OfPublicRestrictedFilterBranchFilterBranches
// OfPublicUnifiedEventsFilterBranchFilterBranches
// OfPublicPropertyAssociationFilterBranchFilterBranches
// OfPublicAssociationFilterBranchFilterBranches]
type APIEventBasedEnrollmentCriteriaRefinementCriteriaUnionFilterBranches struct {
	// This field will be present if the value is a
	// [[]shared.PublicOrFilterBranchFilterBranchUnion] instead of an object.
	OfPublicOrFilterBranchFilterBranches []shared.PublicOrFilterBranchFilterBranchUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]shared.PublicAndFilterBranchFilterBranchUnion] instead of an object.
	OfPublicAndFilterBranchFilterBranches []shared.PublicAndFilterBranchFilterBranchUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]shared.PublicNotAllFilterBranchFilterBranchUnion] instead of an object.
	OfPublicNotAllFilterBranchFilterBranches []shared.PublicNotAllFilterBranchFilterBranchUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]shared.PublicNotAnyFilterBranchFilterBranchUnion] instead of an object.
	OfPublicNotAnyFilterBranchFilterBranches []shared.PublicNotAnyFilterBranchFilterBranchUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]shared.PublicRestrictedFilterBranchFilterBranchUnion] instead of an object.
	OfPublicRestrictedFilterBranchFilterBranches []shared.PublicRestrictedFilterBranchFilterBranchUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]shared.PublicUnifiedEventsFilterBranchFilterBranchUnion] instead of an
	// object.
	OfPublicUnifiedEventsFilterBranchFilterBranches []shared.PublicUnifiedEventsFilterBranchFilterBranchUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]shared.PublicPropertyAssociationFilterBranchFilterBranchUnion] instead of an
	// object.
	OfPublicPropertyAssociationFilterBranchFilterBranches []shared.PublicPropertyAssociationFilterBranchFilterBranchUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]shared.PublicAssociationFilterBranchFilterBranchUnion] instead of an object.
	OfPublicAssociationFilterBranchFilterBranches []shared.PublicAssociationFilterBranchFilterBranchUnion `json:",inline"`
	JSON                                          struct {
		OfPublicOrFilterBranchFilterBranches                  respjson.Field
		OfPublicAndFilterBranchFilterBranches                 respjson.Field
		OfPublicNotAllFilterBranchFilterBranches              respjson.Field
		OfPublicNotAnyFilterBranchFilterBranches              respjson.Field
		OfPublicRestrictedFilterBranchFilterBranches          respjson.Field
		OfPublicUnifiedEventsFilterBranchFilterBranches       respjson.Field
		OfPublicPropertyAssociationFilterBranchFilterBranches respjson.Field
		OfPublicAssociationFilterBranchFilterBranches         respjson.Field
		raw                                                   string
	} `json:"-"`
}

func (r *APIEventBasedEnrollmentCriteriaRefinementCriteriaUnionFilterBranches) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// APIEventBasedEnrollmentCriteriaRefinementCriteriaUnionFilters is an implicit
// subunion of [APIEventBasedEnrollmentCriteriaRefinementCriteriaUnion].
// APIEventBasedEnrollmentCriteriaRefinementCriteriaUnionFilters provides
// convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [APIEventBasedEnrollmentCriteriaRefinementCriteriaUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfPublicOrFilterBranchFilters OfPublicAndFilterBranchFilters
// OfPublicNotAllFilterBranchFilters OfPublicNotAnyFilterBranchFilters
// OfPublicRestrictedFilterBranchFilters OfPublicUnifiedEventsFilterBranchFilters
// OfPublicPropertyAssociationFilterBranchFilters
// OfPublicAssociationFilterBranchFilters]
type APIEventBasedEnrollmentCriteriaRefinementCriteriaUnionFilters struct {
	// This field will be present if the value is a
	// [[]shared.PublicOrFilterBranchFilterUnion] instead of an object.
	OfPublicOrFilterBranchFilters []shared.PublicOrFilterBranchFilterUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]shared.PublicAndFilterBranchFilterUnion] instead of an object.
	OfPublicAndFilterBranchFilters []shared.PublicAndFilterBranchFilterUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]shared.PublicNotAllFilterBranchFilterUnion] instead of an object.
	OfPublicNotAllFilterBranchFilters []shared.PublicNotAllFilterBranchFilterUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]shared.PublicNotAnyFilterBranchFilterUnion] instead of an object.
	OfPublicNotAnyFilterBranchFilters []shared.PublicNotAnyFilterBranchFilterUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]shared.PublicRestrictedFilterBranchFilterUnion] instead of an object.
	OfPublicRestrictedFilterBranchFilters []shared.PublicRestrictedFilterBranchFilterUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]shared.PublicUnifiedEventsFilterBranchFilterUnion] instead of an object.
	OfPublicUnifiedEventsFilterBranchFilters []shared.PublicUnifiedEventsFilterBranchFilterUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]shared.PublicPropertyAssociationFilterBranchFilterUnion] instead of an
	// object.
	OfPublicPropertyAssociationFilterBranchFilters []shared.PublicPropertyAssociationFilterBranchFilterUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]shared.PublicAssociationFilterBranchFilterUnion] instead of an object.
	OfPublicAssociationFilterBranchFilters []shared.PublicAssociationFilterBranchFilterUnion `json:",inline"`
	JSON                                   struct {
		OfPublicOrFilterBranchFilters                  respjson.Field
		OfPublicAndFilterBranchFilters                 respjson.Field
		OfPublicNotAllFilterBranchFilters              respjson.Field
		OfPublicNotAnyFilterBranchFilters              respjson.Field
		OfPublicRestrictedFilterBranchFilters          respjson.Field
		OfPublicUnifiedEventsFilterBranchFilters       respjson.Field
		OfPublicPropertyAssociationFilterBranchFilters respjson.Field
		OfPublicAssociationFilterBranchFilters         respjson.Field
		raw                                            string
	} `json:"-"`
}

func (r *APIEventBasedEnrollmentCriteriaRefinementCriteriaUnionFilters) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties EventFilterBranches, ListMembershipFilterBranches,
// ShouldReEnroll, Type are required.
type APIEventBasedEnrollmentCriteriaParam struct {
	EventFilterBranches []shared.PublicUnifiedEventsFilterBranchParam `json:"eventFilterBranches,omitzero,required"`
	// If you want to listen to list-membership events (an object was added to a list,
	// an object was removed from a list) you need to use this
	// `listMembershipFilterBranches` property instead of `eventFilterBranches`,
	// because list membership events work differently.
	ListMembershipFilterBranches []APIEventBasedEnrollmentCriteriaListMembershipFilterBranchUnionParam `json:"listMembershipFilterBranches,omitzero,required"`
	// Whether or not the same object can enroll in this workflow twice.
	ShouldReEnroll bool `json:"shouldReEnroll,required"`
	// The type of enrollment criteria this is, this can be "LIST_BASED",
	// "EVENT_BASED", or "MANUAL".
	//
	// Any of "EVENT_BASED".
	Type APIEventBasedEnrollmentCriteriaType `json:"type,omitzero,required"`
	// List-based criteria to further refine which contacts will enroll in this flow.
	RefinementCriteria APIEventBasedEnrollmentCriteriaRefinementCriteriaUnionParam `json:"refinementCriteria,omitzero"`
	paramObj
}

func (r APIEventBasedEnrollmentCriteriaParam) MarshalJSON() (data []byte, err error) {
	type shadow APIEventBasedEnrollmentCriteriaParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *APIEventBasedEnrollmentCriteriaParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type APIEventBasedEnrollmentCriteriaListMembershipFilterBranchUnionParam struct {
	OfOr                  *shared.PublicOrFilterBranchParam                  `json:",omitzero,inline"`
	OfAnd                 *shared.PublicAndFilterBranchParam                 `json:",omitzero,inline"`
	OfNotAll              *shared.PublicNotAllFilterBranchParam              `json:",omitzero,inline"`
	OfNotAny              *shared.PublicNotAnyFilterBranchParam              `json:",omitzero,inline"`
	OfRestricted          *shared.PublicRestrictedFilterBranchParam          `json:",omitzero,inline"`
	OfUnifiedEvents       *shared.PublicUnifiedEventsFilterBranchParam       `json:",omitzero,inline"`
	OfPropertyAssociation *shared.PublicPropertyAssociationFilterBranchParam `json:",omitzero,inline"`
	OfAssociation         *shared.PublicAssociationFilterBranchParam         `json:",omitzero,inline"`
	paramUnion
}

func (u APIEventBasedEnrollmentCriteriaListMembershipFilterBranchUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfOr,
		u.OfAnd,
		u.OfNotAll,
		u.OfNotAny,
		u.OfRestricted,
		u.OfUnifiedEvents,
		u.OfPropertyAssociation,
		u.OfAssociation)
}
func (u *APIEventBasedEnrollmentCriteriaListMembershipFilterBranchUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *APIEventBasedEnrollmentCriteriaListMembershipFilterBranchUnionParam) asAny() any {
	if !param.IsOmitted(u.OfOr) {
		return u.OfOr
	} else if !param.IsOmitted(u.OfAnd) {
		return u.OfAnd
	} else if !param.IsOmitted(u.OfNotAll) {
		return u.OfNotAll
	} else if !param.IsOmitted(u.OfNotAny) {
		return u.OfNotAny
	} else if !param.IsOmitted(u.OfRestricted) {
		return u.OfRestricted
	} else if !param.IsOmitted(u.OfUnifiedEvents) {
		return u.OfUnifiedEvents
	} else if !param.IsOmitted(u.OfPropertyAssociation) {
		return u.OfPropertyAssociation
	} else if !param.IsOmitted(u.OfAssociation) {
		return u.OfAssociation
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIEventBasedEnrollmentCriteriaListMembershipFilterBranchUnionParam) GetEventTypeID() *string {
	if vt := u.OfUnifiedEvents; vt != nil {
		return &vt.EventTypeID
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIEventBasedEnrollmentCriteriaListMembershipFilterBranchUnionParam) GetCoalescingRefineBy() *shared.PublicUnifiedEventsFilterBranchCoalescingRefineByUnionParam {
	if vt := u.OfUnifiedEvents; vt != nil {
		return &vt.CoalescingRefineBy
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIEventBasedEnrollmentCriteriaListMembershipFilterBranchUnionParam) GetPropertyWithObjectID() *string {
	if vt := u.OfPropertyAssociation; vt != nil {
		return &vt.PropertyWithObjectID
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIEventBasedEnrollmentCriteriaListMembershipFilterBranchUnionParam) GetAssociationCategory() *string {
	if vt := u.OfAssociation; vt != nil {
		return &vt.AssociationCategory
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIEventBasedEnrollmentCriteriaListMembershipFilterBranchUnionParam) GetAssociationTypeID() *int64 {
	if vt := u.OfAssociation; vt != nil {
		return &vt.AssociationTypeID
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIEventBasedEnrollmentCriteriaListMembershipFilterBranchUnionParam) GetFilterBranchOperator() *string {
	if vt := u.OfOr; vt != nil {
		return (*string)(&vt.FilterBranchOperator)
	} else if vt := u.OfAnd; vt != nil {
		return (*string)(&vt.FilterBranchOperator)
	} else if vt := u.OfNotAll; vt != nil {
		return (*string)(&vt.FilterBranchOperator)
	} else if vt := u.OfNotAny; vt != nil {
		return (*string)(&vt.FilterBranchOperator)
	} else if vt := u.OfRestricted; vt != nil {
		return (*string)(&vt.FilterBranchOperator)
	} else if vt := u.OfUnifiedEvents; vt != nil {
		return (*string)(&vt.FilterBranchOperator)
	} else if vt := u.OfPropertyAssociation; vt != nil {
		return (*string)(&vt.FilterBranchOperator)
	} else if vt := u.OfAssociation; vt != nil {
		return (*string)(&vt.FilterBranchOperator)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIEventBasedEnrollmentCriteriaListMembershipFilterBranchUnionParam) GetFilterBranchType() *string {
	if vt := u.OfOr; vt != nil {
		return (*string)(&vt.FilterBranchType)
	} else if vt := u.OfAnd; vt != nil {
		return (*string)(&vt.FilterBranchType)
	} else if vt := u.OfNotAll; vt != nil {
		return (*string)(&vt.FilterBranchType)
	} else if vt := u.OfNotAny; vt != nil {
		return (*string)(&vt.FilterBranchType)
	} else if vt := u.OfRestricted; vt != nil {
		return (*string)(&vt.FilterBranchType)
	} else if vt := u.OfUnifiedEvents; vt != nil {
		return (*string)(&vt.FilterBranchType)
	} else if vt := u.OfPropertyAssociation; vt != nil {
		return (*string)(&vt.FilterBranchType)
	} else if vt := u.OfAssociation; vt != nil {
		return (*string)(&vt.FilterBranchType)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIEventBasedEnrollmentCriteriaListMembershipFilterBranchUnionParam) GetOperator() *string {
	if vt := u.OfUnifiedEvents; vt != nil {
		return (*string)(&vt.Operator)
	} else if vt := u.OfPropertyAssociation; vt != nil {
		return (*string)(&vt.Operator)
	} else if vt := u.OfAssociation; vt != nil {
		return (*string)(&vt.Operator)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIEventBasedEnrollmentCriteriaListMembershipFilterBranchUnionParam) GetObjectTypeID() *string {
	if vt := u.OfPropertyAssociation; vt != nil {
		return (*string)(&vt.ObjectTypeID)
	} else if vt := u.OfAssociation; vt != nil {
		return (*string)(&vt.ObjectTypeID)
	}
	return nil
}

// Returns a subunion which exports methods to access subproperties
//
// Or use AsAny() to get the underlying value
func (u APIEventBasedEnrollmentCriteriaListMembershipFilterBranchUnionParam) GetFilterBranches() (res apiEventBasedEnrollmentCriteriaListMembershipFilterBranchUnionParamFilterBranches) {
	if vt := u.OfOr; vt != nil {
		res.any = &vt.FilterBranches
	} else if vt := u.OfAnd; vt != nil {
		res.any = &vt.FilterBranches
	} else if vt := u.OfNotAll; vt != nil {
		res.any = &vt.FilterBranches
	} else if vt := u.OfNotAny; vt != nil {
		res.any = &vt.FilterBranches
	} else if vt := u.OfRestricted; vt != nil {
		res.any = &vt.FilterBranches
	} else if vt := u.OfUnifiedEvents; vt != nil {
		res.any = &vt.FilterBranches
	} else if vt := u.OfPropertyAssociation; vt != nil {
		res.any = &vt.FilterBranches
	} else if vt := u.OfAssociation; vt != nil {
		res.any = &vt.FilterBranches
	}
	return
}

// Can have the runtime types
// [_[]shared.PublicOrFilterBranchFilterBranchUnionParam],
// [_[]shared.PublicAndFilterBranchFilterBranchUnionParam],
// [_[]shared.PublicNotAllFilterBranchFilterBranchUnionParam],
// [_[]shared.PublicNotAnyFilterBranchFilterBranchUnionParam],
// [_[]shared.PublicRestrictedFilterBranchFilterBranchUnionParam],
// [_[]shared.PublicUnifiedEventsFilterBranchFilterBranchUnionParam],
// [_[]shared.PublicPropertyAssociationFilterBranchFilterBranchUnionParam],
// [_[]shared.PublicAssociationFilterBranchFilterBranchUnionParam]
type apiEventBasedEnrollmentCriteriaListMembershipFilterBranchUnionParamFilterBranches struct{ any }

// Use the following switch statement to get the type of the union:
//
//	switch u.AsAny().(type) {
//	case *[]shared.PublicOrFilterBranchFilterBranchUnionParam:
//	case *[]shared.PublicAndFilterBranchFilterBranchUnionParam:
//	case *[]shared.PublicNotAllFilterBranchFilterBranchUnionParam:
//	case *[]shared.PublicNotAnyFilterBranchFilterBranchUnionParam:
//	case *[]shared.PublicRestrictedFilterBranchFilterBranchUnionParam:
//	case *[]shared.PublicUnifiedEventsFilterBranchFilterBranchUnionParam:
//	case *[]shared.PublicPropertyAssociationFilterBranchFilterBranchUnionParam:
//	case *[]shared.PublicAssociationFilterBranchFilterBranchUnionParam:
//	default:
//	    fmt.Errorf("not present")
//	}
func (u apiEventBasedEnrollmentCriteriaListMembershipFilterBranchUnionParamFilterBranches) AsAny() any {
	return u.any
}

// Returns a subunion which exports methods to access subproperties
//
// Or use AsAny() to get the underlying value
func (u APIEventBasedEnrollmentCriteriaListMembershipFilterBranchUnionParam) GetFilters() (res apiEventBasedEnrollmentCriteriaListMembershipFilterBranchUnionParamFilters) {
	if vt := u.OfOr; vt != nil {
		res.any = &vt.Filters
	} else if vt := u.OfAnd; vt != nil {
		res.any = &vt.Filters
	} else if vt := u.OfNotAll; vt != nil {
		res.any = &vt.Filters
	} else if vt := u.OfNotAny; vt != nil {
		res.any = &vt.Filters
	} else if vt := u.OfRestricted; vt != nil {
		res.any = &vt.Filters
	} else if vt := u.OfUnifiedEvents; vt != nil {
		res.any = &vt.Filters
	} else if vt := u.OfPropertyAssociation; vt != nil {
		res.any = &vt.Filters
	} else if vt := u.OfAssociation; vt != nil {
		res.any = &vt.Filters
	}
	return
}

// Can have the runtime types [_[]shared.PublicOrFilterBranchFilterUnionParam],
// [_[]shared.PublicAndFilterBranchFilterUnionParam],
// [_[]shared.PublicNotAllFilterBranchFilterUnionParam],
// [_[]shared.PublicNotAnyFilterBranchFilterUnionParam],
// [_[]shared.PublicRestrictedFilterBranchFilterUnionParam],
// [_[]shared.PublicUnifiedEventsFilterBranchFilterUnionParam],
// [_[]shared.PublicPropertyAssociationFilterBranchFilterUnionParam],
// [_[]shared.PublicAssociationFilterBranchFilterUnionParam]
type apiEventBasedEnrollmentCriteriaListMembershipFilterBranchUnionParamFilters struct{ any }

// Use the following switch statement to get the type of the union:
//
//	switch u.AsAny().(type) {
//	case *[]shared.PublicOrFilterBranchFilterUnionParam:
//	case *[]shared.PublicAndFilterBranchFilterUnionParam:
//	case *[]shared.PublicNotAllFilterBranchFilterUnionParam:
//	case *[]shared.PublicNotAnyFilterBranchFilterUnionParam:
//	case *[]shared.PublicRestrictedFilterBranchFilterUnionParam:
//	case *[]shared.PublicUnifiedEventsFilterBranchFilterUnionParam:
//	case *[]shared.PublicPropertyAssociationFilterBranchFilterUnionParam:
//	case *[]shared.PublicAssociationFilterBranchFilterUnionParam:
//	default:
//	    fmt.Errorf("not present")
//	}
func (u apiEventBasedEnrollmentCriteriaListMembershipFilterBranchUnionParamFilters) AsAny() any {
	return u.any
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type APIEventBasedEnrollmentCriteriaRefinementCriteriaUnionParam struct {
	OfOr                  *shared.PublicOrFilterBranchParam                  `json:",omitzero,inline"`
	OfAnd                 *shared.PublicAndFilterBranchParam                 `json:",omitzero,inline"`
	OfNotAll              *shared.PublicNotAllFilterBranchParam              `json:",omitzero,inline"`
	OfNotAny              *shared.PublicNotAnyFilterBranchParam              `json:",omitzero,inline"`
	OfRestricted          *shared.PublicRestrictedFilterBranchParam          `json:",omitzero,inline"`
	OfUnifiedEvents       *shared.PublicUnifiedEventsFilterBranchParam       `json:",omitzero,inline"`
	OfPropertyAssociation *shared.PublicPropertyAssociationFilterBranchParam `json:",omitzero,inline"`
	OfAssociation         *shared.PublicAssociationFilterBranchParam         `json:",omitzero,inline"`
	paramUnion
}

func (u APIEventBasedEnrollmentCriteriaRefinementCriteriaUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfOr,
		u.OfAnd,
		u.OfNotAll,
		u.OfNotAny,
		u.OfRestricted,
		u.OfUnifiedEvents,
		u.OfPropertyAssociation,
		u.OfAssociation)
}
func (u *APIEventBasedEnrollmentCriteriaRefinementCriteriaUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *APIEventBasedEnrollmentCriteriaRefinementCriteriaUnionParam) asAny() any {
	if !param.IsOmitted(u.OfOr) {
		return u.OfOr
	} else if !param.IsOmitted(u.OfAnd) {
		return u.OfAnd
	} else if !param.IsOmitted(u.OfNotAll) {
		return u.OfNotAll
	} else if !param.IsOmitted(u.OfNotAny) {
		return u.OfNotAny
	} else if !param.IsOmitted(u.OfRestricted) {
		return u.OfRestricted
	} else if !param.IsOmitted(u.OfUnifiedEvents) {
		return u.OfUnifiedEvents
	} else if !param.IsOmitted(u.OfPropertyAssociation) {
		return u.OfPropertyAssociation
	} else if !param.IsOmitted(u.OfAssociation) {
		return u.OfAssociation
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIEventBasedEnrollmentCriteriaRefinementCriteriaUnionParam) GetEventTypeID() *string {
	if vt := u.OfUnifiedEvents; vt != nil {
		return &vt.EventTypeID
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIEventBasedEnrollmentCriteriaRefinementCriteriaUnionParam) GetCoalescingRefineBy() *shared.PublicUnifiedEventsFilterBranchCoalescingRefineByUnionParam {
	if vt := u.OfUnifiedEvents; vt != nil {
		return &vt.CoalescingRefineBy
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIEventBasedEnrollmentCriteriaRefinementCriteriaUnionParam) GetPropertyWithObjectID() *string {
	if vt := u.OfPropertyAssociation; vt != nil {
		return &vt.PropertyWithObjectID
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIEventBasedEnrollmentCriteriaRefinementCriteriaUnionParam) GetAssociationCategory() *string {
	if vt := u.OfAssociation; vt != nil {
		return &vt.AssociationCategory
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIEventBasedEnrollmentCriteriaRefinementCriteriaUnionParam) GetAssociationTypeID() *int64 {
	if vt := u.OfAssociation; vt != nil {
		return &vt.AssociationTypeID
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIEventBasedEnrollmentCriteriaRefinementCriteriaUnionParam) GetFilterBranchOperator() *string {
	if vt := u.OfOr; vt != nil {
		return (*string)(&vt.FilterBranchOperator)
	} else if vt := u.OfAnd; vt != nil {
		return (*string)(&vt.FilterBranchOperator)
	} else if vt := u.OfNotAll; vt != nil {
		return (*string)(&vt.FilterBranchOperator)
	} else if vt := u.OfNotAny; vt != nil {
		return (*string)(&vt.FilterBranchOperator)
	} else if vt := u.OfRestricted; vt != nil {
		return (*string)(&vt.FilterBranchOperator)
	} else if vt := u.OfUnifiedEvents; vt != nil {
		return (*string)(&vt.FilterBranchOperator)
	} else if vt := u.OfPropertyAssociation; vt != nil {
		return (*string)(&vt.FilterBranchOperator)
	} else if vt := u.OfAssociation; vt != nil {
		return (*string)(&vt.FilterBranchOperator)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIEventBasedEnrollmentCriteriaRefinementCriteriaUnionParam) GetFilterBranchType() *string {
	if vt := u.OfOr; vt != nil {
		return (*string)(&vt.FilterBranchType)
	} else if vt := u.OfAnd; vt != nil {
		return (*string)(&vt.FilterBranchType)
	} else if vt := u.OfNotAll; vt != nil {
		return (*string)(&vt.FilterBranchType)
	} else if vt := u.OfNotAny; vt != nil {
		return (*string)(&vt.FilterBranchType)
	} else if vt := u.OfRestricted; vt != nil {
		return (*string)(&vt.FilterBranchType)
	} else if vt := u.OfUnifiedEvents; vt != nil {
		return (*string)(&vt.FilterBranchType)
	} else if vt := u.OfPropertyAssociation; vt != nil {
		return (*string)(&vt.FilterBranchType)
	} else if vt := u.OfAssociation; vt != nil {
		return (*string)(&vt.FilterBranchType)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIEventBasedEnrollmentCriteriaRefinementCriteriaUnionParam) GetOperator() *string {
	if vt := u.OfUnifiedEvents; vt != nil {
		return (*string)(&vt.Operator)
	} else if vt := u.OfPropertyAssociation; vt != nil {
		return (*string)(&vt.Operator)
	} else if vt := u.OfAssociation; vt != nil {
		return (*string)(&vt.Operator)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIEventBasedEnrollmentCriteriaRefinementCriteriaUnionParam) GetObjectTypeID() *string {
	if vt := u.OfPropertyAssociation; vt != nil {
		return (*string)(&vt.ObjectTypeID)
	} else if vt := u.OfAssociation; vt != nil {
		return (*string)(&vt.ObjectTypeID)
	}
	return nil
}

// Returns a subunion which exports methods to access subproperties
//
// Or use AsAny() to get the underlying value
func (u APIEventBasedEnrollmentCriteriaRefinementCriteriaUnionParam) GetFilterBranches() (res apiEventBasedEnrollmentCriteriaRefinementCriteriaUnionParamFilterBranches) {
	if vt := u.OfOr; vt != nil {
		res.any = &vt.FilterBranches
	} else if vt := u.OfAnd; vt != nil {
		res.any = &vt.FilterBranches
	} else if vt := u.OfNotAll; vt != nil {
		res.any = &vt.FilterBranches
	} else if vt := u.OfNotAny; vt != nil {
		res.any = &vt.FilterBranches
	} else if vt := u.OfRestricted; vt != nil {
		res.any = &vt.FilterBranches
	} else if vt := u.OfUnifiedEvents; vt != nil {
		res.any = &vt.FilterBranches
	} else if vt := u.OfPropertyAssociation; vt != nil {
		res.any = &vt.FilterBranches
	} else if vt := u.OfAssociation; vt != nil {
		res.any = &vt.FilterBranches
	}
	return
}

// Can have the runtime types
// [_[]shared.PublicOrFilterBranchFilterBranchUnionParam],
// [_[]shared.PublicAndFilterBranchFilterBranchUnionParam],
// [_[]shared.PublicNotAllFilterBranchFilterBranchUnionParam],
// [_[]shared.PublicNotAnyFilterBranchFilterBranchUnionParam],
// [_[]shared.PublicRestrictedFilterBranchFilterBranchUnionParam],
// [_[]shared.PublicUnifiedEventsFilterBranchFilterBranchUnionParam],
// [_[]shared.PublicPropertyAssociationFilterBranchFilterBranchUnionParam],
// [_[]shared.PublicAssociationFilterBranchFilterBranchUnionParam]
type apiEventBasedEnrollmentCriteriaRefinementCriteriaUnionParamFilterBranches struct{ any }

// Use the following switch statement to get the type of the union:
//
//	switch u.AsAny().(type) {
//	case *[]shared.PublicOrFilterBranchFilterBranchUnionParam:
//	case *[]shared.PublicAndFilterBranchFilterBranchUnionParam:
//	case *[]shared.PublicNotAllFilterBranchFilterBranchUnionParam:
//	case *[]shared.PublicNotAnyFilterBranchFilterBranchUnionParam:
//	case *[]shared.PublicRestrictedFilterBranchFilterBranchUnionParam:
//	case *[]shared.PublicUnifiedEventsFilterBranchFilterBranchUnionParam:
//	case *[]shared.PublicPropertyAssociationFilterBranchFilterBranchUnionParam:
//	case *[]shared.PublicAssociationFilterBranchFilterBranchUnionParam:
//	default:
//	    fmt.Errorf("not present")
//	}
func (u apiEventBasedEnrollmentCriteriaRefinementCriteriaUnionParamFilterBranches) AsAny() any {
	return u.any
}

// Returns a subunion which exports methods to access subproperties
//
// Or use AsAny() to get the underlying value
func (u APIEventBasedEnrollmentCriteriaRefinementCriteriaUnionParam) GetFilters() (res apiEventBasedEnrollmentCriteriaRefinementCriteriaUnionParamFilters) {
	if vt := u.OfOr; vt != nil {
		res.any = &vt.Filters
	} else if vt := u.OfAnd; vt != nil {
		res.any = &vt.Filters
	} else if vt := u.OfNotAll; vt != nil {
		res.any = &vt.Filters
	} else if vt := u.OfNotAny; vt != nil {
		res.any = &vt.Filters
	} else if vt := u.OfRestricted; vt != nil {
		res.any = &vt.Filters
	} else if vt := u.OfUnifiedEvents; vt != nil {
		res.any = &vt.Filters
	} else if vt := u.OfPropertyAssociation; vt != nil {
		res.any = &vt.Filters
	} else if vt := u.OfAssociation; vt != nil {
		res.any = &vt.Filters
	}
	return
}

// Can have the runtime types [_[]shared.PublicOrFilterBranchFilterUnionParam],
// [_[]shared.PublicAndFilterBranchFilterUnionParam],
// [_[]shared.PublicNotAllFilterBranchFilterUnionParam],
// [_[]shared.PublicNotAnyFilterBranchFilterUnionParam],
// [_[]shared.PublicRestrictedFilterBranchFilterUnionParam],
// [_[]shared.PublicUnifiedEventsFilterBranchFilterUnionParam],
// [_[]shared.PublicPropertyAssociationFilterBranchFilterUnionParam],
// [_[]shared.PublicAssociationFilterBranchFilterUnionParam]
type apiEventBasedEnrollmentCriteriaRefinementCriteriaUnionParamFilters struct{ any }

// Use the following switch statement to get the type of the union:
//
//	switch u.AsAny().(type) {
//	case *[]shared.PublicOrFilterBranchFilterUnionParam:
//	case *[]shared.PublicAndFilterBranchFilterUnionParam:
//	case *[]shared.PublicNotAllFilterBranchFilterUnionParam:
//	case *[]shared.PublicNotAnyFilterBranchFilterUnionParam:
//	case *[]shared.PublicRestrictedFilterBranchFilterUnionParam:
//	case *[]shared.PublicUnifiedEventsFilterBranchFilterUnionParam:
//	case *[]shared.PublicPropertyAssociationFilterBranchFilterUnionParam:
//	case *[]shared.PublicAssociationFilterBranchFilterUnionParam:
//	default:
//	    fmt.Errorf("not present")
//	}
func (u apiEventBasedEnrollmentCriteriaRefinementCriteriaUnionParamFilters) AsAny() any { return u.any }

type APIFetchedObjectPropertyValue struct {
	// The token to use to identify the object property to use
	PropertyToken string `json:"propertyToken,required"`
	// This is the type of input value. This can be one of: "FIELD_DATA",
	// "OBJECT_PROPERTY", "STATIC_VALUE", "RELATIVE_DATETIME", "TIMESTAMP",
	// "INCREMENT", "FETCHED_OBJECT_PROPERTY", "APPEND_OBJECT_PROPERTY",
	// "STATIC_APPEND_VALUE", "ENROLLMENT_EVENT_PROPERTY"
	//
	// Any of "FETCHED_OBJECT_PROPERTY".
	Type APIFetchedObjectPropertyValueType `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PropertyToken respjson.Field
		Type          respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r APIFetchedObjectPropertyValue) RawJSON() string { return r.JSON.raw }
func (r *APIFetchedObjectPropertyValue) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this APIFetchedObjectPropertyValue to a
// APIFetchedObjectPropertyValueParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// APIFetchedObjectPropertyValueParam.Overrides()
func (r APIFetchedObjectPropertyValue) ToParam() APIFetchedObjectPropertyValueParam {
	return param.Override[APIFetchedObjectPropertyValueParam](json.RawMessage(r.RawJSON()))
}

// This is the type of input value. This can be one of: "FIELD_DATA",
// "OBJECT_PROPERTY", "STATIC_VALUE", "RELATIVE_DATETIME", "TIMESTAMP",
// "INCREMENT", "FETCHED_OBJECT_PROPERTY", "APPEND_OBJECT_PROPERTY",
// "STATIC_APPEND_VALUE", "ENROLLMENT_EVENT_PROPERTY"
type APIFetchedObjectPropertyValueType string

const (
	APIFetchedObjectPropertyValueTypeFetchedObjectProperty APIFetchedObjectPropertyValueType = "FETCHED_OBJECT_PROPERTY"
)

// The properties PropertyToken, Type are required.
type APIFetchedObjectPropertyValueParam struct {
	// The token to use to identify the object property to use
	PropertyToken string `json:"propertyToken,required"`
	// This is the type of input value. This can be one of: "FIELD_DATA",
	// "OBJECT_PROPERTY", "STATIC_VALUE", "RELATIVE_DATETIME", "TIMESTAMP",
	// "INCREMENT", "FETCHED_OBJECT_PROPERTY", "APPEND_OBJECT_PROPERTY",
	// "STATIC_APPEND_VALUE", "ENROLLMENT_EVENT_PROPERTY"
	//
	// Any of "FETCHED_OBJECT_PROPERTY".
	Type APIFetchedObjectPropertyValueType `json:"type,omitzero,required"`
	paramObj
}

func (r APIFetchedObjectPropertyValueParam) MarshalJSON() (data []byte, err error) {
	type shadow APIFetchedObjectPropertyValueParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *APIFetchedObjectPropertyValueParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// APIFlowUnion contains all possible properties and values from [APIContactFlow],
// [APIPlatformFlow].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type APIFlowUnion struct {
	ID string `json:"id"`
	// This field is a union of [[]APIContactFlowActionUnion],
	// [[]APIPlatformFlowActionUnion]
	Actions      APIFlowUnionActions `json:"actions"`
	BlockedDates []APIBlockedDate    `json:"blockedDates"`
	// This field is from variant [APIContactFlow].
	CanEnrollFromSalesforce bool      `json:"canEnrollFromSalesforce"`
	CreatedAt               time.Time `json:"createdAt"`
	CRMObjectCreationStatus string    `json:"crmObjectCreationStatus"`
	CustomProperties        string    `json:"customProperties"`
	// This field is a union of [[]APIContactFlowDataSourceUnion],
	// [[]APIPlatformFlowDataSourceUnion]
	DataSources           APIFlowUnionDataSources `json:"dataSources"`
	FlowType              string                  `json:"flowType"`
	IsEnabled             bool                    `json:"isEnabled"`
	NextAvailableActionID string                  `json:"nextAvailableActionId"`
	ObjectTypeID          string                  `json:"objectTypeId"`
	RevisionID            string                  `json:"revisionId"`
	// This field is from variant [APIContactFlow].
	SuppressionListIDs []int64         `json:"suppressionListIds"`
	TimeWindows        []APITimeWindow `json:"timeWindows"`
	Type               string          `json:"type"`
	UpdatedAt          time.Time       `json:"updatedAt"`
	Description        string          `json:"description"`
	// This field is a union of [APIContactFlowEnrollmentCriteriaUnion],
	// [APIPlatformFlowEnrollmentCriteriaUnion]
	EnrollmentCriteria APIFlowUnionEnrollmentCriteria `json:"enrollmentCriteria"`
	// This field is a union of [APIContactFlowEnrollmentScheduleUnion],
	// [APIPlatformFlowEnrollmentScheduleUnion]
	EnrollmentSchedule APIFlowUnionEnrollmentSchedule `json:"enrollmentSchedule"`
	// This field is from variant [APIContactFlow].
	EventAnchor APIContactFlowEventAnchorUnion `json:"eventAnchor"`
	// This field is from variant [APIContactFlow].
	GoalFilterBranch APIContactFlowGoalFilterBranchUnion `json:"goalFilterBranch"`
	Name             string                              `json:"name"`
	StartActionID    string                              `json:"startActionId"`
	// This field is from variant [APIContactFlow].
	UnEnrollmentSetting APIUnEnrollmentSetting `json:"unEnrollmentSetting"`
	Uuid                string                 `json:"uuid"`
	// This field is from variant [APIPlatformFlow].
	SuppressionFilterBranch APIPlatformFlowSuppressionFilterBranchUnion `json:"suppressionFilterBranch"`
	JSON                    struct {
		ID                      respjson.Field
		Actions                 respjson.Field
		BlockedDates            respjson.Field
		CanEnrollFromSalesforce respjson.Field
		CreatedAt               respjson.Field
		CRMObjectCreationStatus respjson.Field
		CustomProperties        respjson.Field
		DataSources             respjson.Field
		FlowType                respjson.Field
		IsEnabled               respjson.Field
		NextAvailableActionID   respjson.Field
		ObjectTypeID            respjson.Field
		RevisionID              respjson.Field
		SuppressionListIDs      respjson.Field
		TimeWindows             respjson.Field
		Type                    respjson.Field
		UpdatedAt               respjson.Field
		Description             respjson.Field
		EnrollmentCriteria      respjson.Field
		EnrollmentSchedule      respjson.Field
		EventAnchor             respjson.Field
		GoalFilterBranch        respjson.Field
		Name                    respjson.Field
		StartActionID           respjson.Field
		UnEnrollmentSetting     respjson.Field
		Uuid                    respjson.Field
		SuppressionFilterBranch respjson.Field
		raw                     string
	} `json:"-"`
}

func (u APIFlowUnion) AsAPIContactFlow() (v APIContactFlow) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u APIFlowUnion) AsAPIPlatformFlow() (v APIPlatformFlow) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u APIFlowUnion) RawJSON() string { return u.JSON.raw }

func (r *APIFlowUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// APIFlowUnionActions is an implicit subunion of [APIFlowUnion].
// APIFlowUnionActions provides convenient access to the sub-properties of the
// union.
//
// For type safety it is recommended to directly use a variant of the
// [APIFlowUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfAPIContactFlowActions OfAPIPlatformFlowActions]
type APIFlowUnionActions struct {
	// This field will be present if the value is a [[]APIContactFlowActionUnion]
	// instead of an object.
	OfAPIContactFlowActions []APIContactFlowActionUnion `json:",inline"`
	// This field will be present if the value is a [[]APIPlatformFlowActionUnion]
	// instead of an object.
	OfAPIPlatformFlowActions []APIPlatformFlowActionUnion `json:",inline"`
	JSON                     struct {
		OfAPIContactFlowActions  respjson.Field
		OfAPIPlatformFlowActions respjson.Field
		raw                      string
	} `json:"-"`
}

func (r *APIFlowUnionActions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// APIFlowUnionDataSources is an implicit subunion of [APIFlowUnion].
// APIFlowUnionDataSources provides convenient access to the sub-properties of the
// union.
//
// For type safety it is recommended to directly use a variant of the
// [APIFlowUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfAPIContactFlowDataSources OfAPIPlatformFlowDataSources]
type APIFlowUnionDataSources struct {
	// This field will be present if the value is a [[]APIContactFlowDataSourceUnion]
	// instead of an object.
	OfAPIContactFlowDataSources []APIContactFlowDataSourceUnion `json:",inline"`
	// This field will be present if the value is a [[]APIPlatformFlowDataSourceUnion]
	// instead of an object.
	OfAPIPlatformFlowDataSources []APIPlatformFlowDataSourceUnion `json:",inline"`
	JSON                         struct {
		OfAPIContactFlowDataSources  respjson.Field
		OfAPIPlatformFlowDataSources respjson.Field
		raw                          string
	} `json:"-"`
}

func (r *APIFlowUnionDataSources) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// APIFlowUnionEnrollmentCriteria is an implicit subunion of [APIFlowUnion].
// APIFlowUnionEnrollmentCriteria provides convenient access to the sub-properties
// of the union.
//
// For type safety it is recommended to directly use a variant of the
// [APIFlowUnion].
type APIFlowUnionEnrollmentCriteria struct {
	// This field is from variant [APIContactFlowEnrollmentCriteriaUnion],
	// [APIPlatformFlowEnrollmentCriteriaUnion].
	ListFilterBranch APIListBasedEnrollmentCriteriaListFilterBranchUnion `json:"listFilterBranch"`
	// This field is from variant [APIContactFlowEnrollmentCriteriaUnion],
	// [APIPlatformFlowEnrollmentCriteriaUnion].
	ReEnrollmentTriggersFilterBranches []APIListBasedEnrollmentCriteriaReEnrollmentTriggersFilterBranchUnion `json:"reEnrollmentTriggersFilterBranches"`
	ShouldReEnroll                     bool                                                                  `json:"shouldReEnroll"`
	Type                               string                                                                `json:"type"`
	// This field is from variant [APIContactFlowEnrollmentCriteriaUnion],
	// [APIPlatformFlowEnrollmentCriteriaUnion].
	UnEnrollObjectsNotMeetingCriteria bool `json:"unEnrollObjectsNotMeetingCriteria"`
	// This field is from variant [APIContactFlowEnrollmentCriteriaUnion],
	// [APIPlatformFlowEnrollmentCriteriaUnion].
	EventFilterBranches []shared.PublicUnifiedEventsFilterBranch `json:"eventFilterBranches"`
	// This field is from variant [APIContactFlowEnrollmentCriteriaUnion],
	// [APIPlatformFlowEnrollmentCriteriaUnion].
	ListMembershipFilterBranches []APIEventBasedEnrollmentCriteriaListMembershipFilterBranchUnion `json:"listMembershipFilterBranches"`
	// This field is from variant [APIContactFlowEnrollmentCriteriaUnion],
	// [APIPlatformFlowEnrollmentCriteriaUnion].
	RefinementCriteria APIEventBasedEnrollmentCriteriaRefinementCriteriaUnion `json:"refinementCriteria"`
	JSON               struct {
		ListFilterBranch                   respjson.Field
		ReEnrollmentTriggersFilterBranches respjson.Field
		ShouldReEnroll                     respjson.Field
		Type                               respjson.Field
		UnEnrollObjectsNotMeetingCriteria  respjson.Field
		EventFilterBranches                respjson.Field
		ListMembershipFilterBranches       respjson.Field
		RefinementCriteria                 respjson.Field
		raw                                string
	} `json:"-"`
}

func (r *APIFlowUnionEnrollmentCriteria) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// APIFlowUnionEnrollmentSchedule is an implicit subunion of [APIFlowUnion].
// APIFlowUnionEnrollmentSchedule provides convenient access to the sub-properties
// of the union.
//
// For type safety it is recommended to directly use a variant of the
// [APIFlowUnion].
type APIFlowUnionEnrollmentSchedule struct {
	// This field is from variant [APIContactFlowEnrollmentScheduleUnion],
	// [APIPlatformFlowEnrollmentScheduleUnion].
	TimeOfDay APITimeOfDay `json:"timeOfDay"`
	Type      string       `json:"type"`
	// This field is from variant [APIContactFlowEnrollmentScheduleUnion],
	// [APIPlatformFlowEnrollmentScheduleUnion].
	DaysOfWeek []string `json:"daysOfWeek"`
	// This field is from variant [APIContactFlowEnrollmentScheduleUnion],
	// [APIPlatformFlowEnrollmentScheduleUnion].
	DaysOfMonth []int64 `json:"daysOfMonth"`
	// This field is from variant [APIContactFlowEnrollmentScheduleUnion],
	// [APIPlatformFlowEnrollmentScheduleUnion].
	MonthlyRelativeDays APIMonthlyRelativeDaysEnrollmentScheduleMonthlyRelativeDays `json:"monthlyRelativeDays"`
	// This field is from variant [APIContactFlowEnrollmentScheduleUnion],
	// [APIPlatformFlowEnrollmentScheduleUnion].
	DayOfMonth int64 `json:"dayOfMonth"`
	// This field is from variant [APIContactFlowEnrollmentScheduleUnion],
	// [APIPlatformFlowEnrollmentScheduleUnion].
	Month APIYearlyEnrollmentScheduleMonth `json:"month"`
	// This field is from variant [APIContactFlowEnrollmentScheduleUnion],
	// [APIPlatformFlowEnrollmentScheduleUnion].
	DateProperty string `json:"dateProperty"`
	// This field is from variant [APIContactFlowEnrollmentScheduleUnion],
	// [APIPlatformFlowEnrollmentScheduleUnion].
	DaysDelta int64 `json:"daysDelta"`
	// This field is from variant [APIContactFlowEnrollmentScheduleUnion],
	// [APIPlatformFlowEnrollmentScheduleUnion].
	Yearly bool `json:"yearly"`
	JSON   struct {
		TimeOfDay           respjson.Field
		Type                respjson.Field
		DaysOfWeek          respjson.Field
		DaysOfMonth         respjson.Field
		MonthlyRelativeDays respjson.Field
		DayOfMonth          respjson.Field
		Month               respjson.Field
		DateProperty        respjson.Field
		DaysDelta           respjson.Field
		Yearly              respjson.Field
		raw                 string
	} `json:"-"`
}

func (r *APIFlowUnionEnrollmentSchedule) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties FlowID, Type are required.
type APIFlowBatchFetchFlowIDCoordinateParam struct {
	FlowID string `json:"flowId,required"`
	// Any of "FLOW_ID".
	Type APIFlowBatchFetchFlowIDCoordinateType `json:"type,omitzero,required"`
	paramObj
}

func (r APIFlowBatchFetchFlowIDCoordinateParam) MarshalJSON() (data []byte, err error) {
	type shadow APIFlowBatchFetchFlowIDCoordinateParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *APIFlowBatchFetchFlowIDCoordinateParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type APIFlowBatchFetchFlowIDCoordinateType string

const (
	APIFlowBatchFetchFlowIDCoordinateTypeFlowID APIFlowBatchFetchFlowIDCoordinateType = "FLOW_ID"
)

// The properties FlowMigrationStatuses, Type are required.
type APIFlowBatchFetchMigrationFlowIDCoordinateParam struct {
	// The flowId from the V4 API
	FlowMigrationStatuses string `json:"flowMigrationStatuses,required"`
	// The type of input this is, can be FLOW_ID or WORKFLOW_ID
	//
	// Any of "FLOW_ID".
	Type APIFlowBatchFetchMigrationFlowIDCoordinateType `json:"type,omitzero,required"`
	paramObj
}

func (r APIFlowBatchFetchMigrationFlowIDCoordinateParam) MarshalJSON() (data []byte, err error) {
	type shadow APIFlowBatchFetchMigrationFlowIDCoordinateParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *APIFlowBatchFetchMigrationFlowIDCoordinateParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The type of input this is, can be FLOW_ID or WORKFLOW_ID
type APIFlowBatchFetchMigrationFlowIDCoordinateType string

const (
	APIFlowBatchFetchMigrationFlowIDCoordinateTypeFlowID APIFlowBatchFetchMigrationFlowIDCoordinateType = "FLOW_ID"
)

// The properties FlowMigrationStatusForClassicWorkflows, Type are required.
type APIFlowBatchFetchMigrationWorkflowIDCoordinateParam struct {
	// The workflowId from the V3 API
	FlowMigrationStatusForClassicWorkflows string `json:"flowMigrationStatusForClassicWorkflows,required"`
	// The type of input this is, can be FLOW_ID or WORKFLOW_ID
	//
	// Any of "WORKFLOW_ID".
	Type APIFlowBatchFetchMigrationWorkflowIDCoordinateType `json:"type,omitzero,required"`
	paramObj
}

func (r APIFlowBatchFetchMigrationWorkflowIDCoordinateParam) MarshalJSON() (data []byte, err error) {
	type shadow APIFlowBatchFetchMigrationWorkflowIDCoordinateParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *APIFlowBatchFetchMigrationWorkflowIDCoordinateParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The type of input this is, can be FLOW_ID or WORKFLOW_ID
type APIFlowBatchFetchMigrationWorkflowIDCoordinateType string

const (
	APIFlowBatchFetchMigrationWorkflowIDCoordinateTypeWorkflowID APIFlowBatchFetchMigrationWorkflowIDCoordinateType = "WORKFLOW_ID"
)

// The property Inputs is required.
type APIFlowBatchInputParam struct {
	Inputs []APIFlowBatchFetchFlowIDCoordinateParam `json:"inputs,omitzero,required"`
	paramObj
}

func (r APIFlowBatchInputParam) MarshalJSON() (data []byte, err error) {
	type shadow APIFlowBatchInputParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *APIFlowBatchInputParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Inputs is required.
type APIFlowBatchMigrationInputParam struct {
	Inputs []APIFlowBatchMigrationInputInputUnionParam `json:"inputs,omitzero,required"`
	paramObj
}

func (r APIFlowBatchMigrationInputParam) MarshalJSON() (data []byte, err error) {
	type shadow APIFlowBatchMigrationInputParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *APIFlowBatchMigrationInputParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type APIFlowBatchMigrationInputInputUnionParam struct {
	OfFlowID     *APIFlowBatchFetchMigrationFlowIDCoordinateParam     `json:",omitzero,inline"`
	OfWorkflowID *APIFlowBatchFetchMigrationWorkflowIDCoordinateParam `json:",omitzero,inline"`
	paramUnion
}

func (u APIFlowBatchMigrationInputInputUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfFlowID, u.OfWorkflowID)
}
func (u *APIFlowBatchMigrationInputInputUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *APIFlowBatchMigrationInputInputUnionParam) asAny() any {
	if !param.IsOmitted(u.OfFlowID) {
		return u.OfFlowID
	} else if !param.IsOmitted(u.OfWorkflowID) {
		return u.OfWorkflowID
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIFlowBatchMigrationInputInputUnionParam) GetFlowMigrationStatuses() *string {
	if vt := u.OfFlowID; vt != nil {
		return &vt.FlowMigrationStatuses
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIFlowBatchMigrationInputInputUnionParam) GetFlowMigrationStatusForClassicWorkflows() *string {
	if vt := u.OfWorkflowID; vt != nil {
		return &vt.FlowMigrationStatusForClassicWorkflows
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIFlowBatchMigrationInputInputUnionParam) GetType() *string {
	if vt := u.OfFlowID; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfWorkflowID; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type APIFlowCreateRequestUnionParam struct {
	OfAPIContactFlowCreateRequest  *APIContactFlowCreateRequestParam  `json:",omitzero,inline"`
	OfAPIPlatformFlowCreateRequest *APIPlatformFlowCreateRequestParam `json:",omitzero,inline"`
	paramUnion
}

func (u APIFlowCreateRequestUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfAPIContactFlowCreateRequest, u.OfAPIPlatformFlowCreateRequest)
}
func (u *APIFlowCreateRequestUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *APIFlowCreateRequestUnionParam) asAny() any {
	if !param.IsOmitted(u.OfAPIContactFlowCreateRequest) {
		return u.OfAPIContactFlowCreateRequest
	} else if !param.IsOmitted(u.OfAPIPlatformFlowCreateRequest) {
		return u.OfAPIPlatformFlowCreateRequest
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIFlowCreateRequestUnionParam) GetCanEnrollFromSalesforce() *bool {
	if vt := u.OfAPIContactFlowCreateRequest; vt != nil {
		return &vt.CanEnrollFromSalesforce
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIFlowCreateRequestUnionParam) GetSuppressionListIDs() []int64 {
	if vt := u.OfAPIContactFlowCreateRequest; vt != nil {
		return vt.SuppressionListIDs
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIFlowCreateRequestUnionParam) GetEventAnchor() *APIContactFlowCreateRequestEventAnchorUnionParam {
	if vt := u.OfAPIContactFlowCreateRequest; vt != nil {
		return &vt.EventAnchor
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIFlowCreateRequestUnionParam) GetGoalFilterBranch() *APIContactFlowCreateRequestGoalFilterBranchUnionParam {
	if vt := u.OfAPIContactFlowCreateRequest; vt != nil {
		return &vt.GoalFilterBranch
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIFlowCreateRequestUnionParam) GetUnEnrollmentSetting() *APIUnEnrollmentSettingParam {
	if vt := u.OfAPIContactFlowCreateRequest; vt != nil {
		return &vt.UnEnrollmentSetting
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIFlowCreateRequestUnionParam) GetSuppressionFilterBranch() *APIPlatformFlowCreateRequestSuppressionFilterBranchUnionParam {
	if vt := u.OfAPIPlatformFlowCreateRequest; vt != nil {
		return &vt.SuppressionFilterBranch
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIFlowCreateRequestUnionParam) GetFlowType() *string {
	if vt := u.OfAPIContactFlowCreateRequest; vt != nil {
		return (*string)(&vt.FlowType)
	} else if vt := u.OfAPIPlatformFlowCreateRequest; vt != nil {
		return (*string)(&vt.FlowType)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIFlowCreateRequestUnionParam) GetIsEnabled() *bool {
	if vt := u.OfAPIContactFlowCreateRequest; vt != nil {
		return (*bool)(&vt.IsEnabled)
	} else if vt := u.OfAPIPlatformFlowCreateRequest; vt != nil {
		return (*bool)(&vt.IsEnabled)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIFlowCreateRequestUnionParam) GetObjectTypeID() *string {
	if vt := u.OfAPIContactFlowCreateRequest; vt != nil {
		return (*string)(&vt.ObjectTypeID)
	} else if vt := u.OfAPIPlatformFlowCreateRequest; vt != nil {
		return (*string)(&vt.ObjectTypeID)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIFlowCreateRequestUnionParam) GetType() *string {
	if vt := u.OfAPIContactFlowCreateRequest; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfAPIPlatformFlowCreateRequest; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIFlowCreateRequestUnionParam) GetDescription() *string {
	if vt := u.OfAPIContactFlowCreateRequest; vt != nil && vt.Description.Valid() {
		return &vt.Description.Value
	} else if vt := u.OfAPIPlatformFlowCreateRequest; vt != nil && vt.Description.Valid() {
		return &vt.Description.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIFlowCreateRequestUnionParam) GetName() *string {
	if vt := u.OfAPIContactFlowCreateRequest; vt != nil && vt.Name.Valid() {
		return &vt.Name.Value
	} else if vt := u.OfAPIPlatformFlowCreateRequest; vt != nil && vt.Name.Valid() {
		return &vt.Name.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIFlowCreateRequestUnionParam) GetStartActionID() *string {
	if vt := u.OfAPIContactFlowCreateRequest; vt != nil && vt.StartActionID.Valid() {
		return &vt.StartActionID.Value
	} else if vt := u.OfAPIPlatformFlowCreateRequest; vt != nil && vt.StartActionID.Valid() {
		return &vt.StartActionID.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIFlowCreateRequestUnionParam) GetUuid() *string {
	if vt := u.OfAPIContactFlowCreateRequest; vt != nil && vt.Uuid.Valid() {
		return &vt.Uuid.Value
	} else if vt := u.OfAPIPlatformFlowCreateRequest; vt != nil && vt.Uuid.Valid() {
		return &vt.Uuid.Value
	}
	return nil
}

// Returns a subunion which exports methods to access subproperties
//
// Or use AsAny() to get the underlying value
func (u APIFlowCreateRequestUnionParam) GetActions() (res apiFlowCreateRequestUnionParamActions) {
	if vt := u.OfAPIContactFlowCreateRequest; vt != nil {
		res.any = &vt.Actions
	} else if vt := u.OfAPIPlatformFlowCreateRequest; vt != nil {
		res.any = &vt.Actions
	}
	return
}

// Can have the runtime types [_[]APIContactFlowCreateRequestActionUnionParam],
// [_[]APIPlatformFlowCreateRequestActionUnionParam]
type apiFlowCreateRequestUnionParamActions struct{ any }

// Use the following switch statement to get the type of the union:
//
//	switch u.AsAny().(type) {
//	case *[]automation.APIContactFlowCreateRequestActionUnionParam:
//	case *[]automation.APIPlatformFlowCreateRequestActionUnionParam:
//	default:
//	    fmt.Errorf("not present")
//	}
func (u apiFlowCreateRequestUnionParamActions) AsAny() any { return u.any }

// Returns a pointer to the underlying variant's BlockedDates property, if present.
func (u APIFlowCreateRequestUnionParam) GetBlockedDates() []APIBlockedDateParam {
	if vt := u.OfAPIContactFlowCreateRequest; vt != nil {
		return vt.BlockedDates
	} else if vt := u.OfAPIPlatformFlowCreateRequest; vt != nil {
		return vt.BlockedDates
	}
	return nil
}

// Returns a pointer to the underlying variant's CustomProperties property, if
// present.
func (u APIFlowCreateRequestUnionParam) GetCustomProperties() map[string]string {
	if vt := u.OfAPIContactFlowCreateRequest; vt != nil {
		return vt.CustomProperties
	} else if vt := u.OfAPIPlatformFlowCreateRequest; vt != nil {
		return vt.CustomProperties
	}
	return nil
}

// Returns a subunion which exports methods to access subproperties
//
// Or use AsAny() to get the underlying value
func (u APIFlowCreateRequestUnionParam) GetDataSources() (res apiFlowCreateRequestUnionParamDataSources) {
	if vt := u.OfAPIContactFlowCreateRequest; vt != nil {
		res.any = &vt.DataSources
	} else if vt := u.OfAPIPlatformFlowCreateRequest; vt != nil {
		res.any = &vt.DataSources
	}
	return
}

// Can have the runtime types [_[]APIContactFlowCreateRequestDataSourceUnionParam],
// [_[]APIPlatformFlowCreateRequestDataSourceUnionParam]
type apiFlowCreateRequestUnionParamDataSources struct{ any }

// Use the following switch statement to get the type of the union:
//
//	switch u.AsAny().(type) {
//	case *[]automation.APIContactFlowCreateRequestDataSourceUnionParam:
//	case *[]automation.APIPlatformFlowCreateRequestDataSourceUnionParam:
//	default:
//	    fmt.Errorf("not present")
//	}
func (u apiFlowCreateRequestUnionParamDataSources) AsAny() any { return u.any }

// Returns a pointer to the underlying variant's TimeWindows property, if present.
func (u APIFlowCreateRequestUnionParam) GetTimeWindows() []APITimeWindowParam {
	if vt := u.OfAPIContactFlowCreateRequest; vt != nil {
		return vt.TimeWindows
	} else if vt := u.OfAPIPlatformFlowCreateRequest; vt != nil {
		return vt.TimeWindows
	}
	return nil
}

// Returns a subunion which exports methods to access subproperties
//
// Or use AsAny() to get the underlying value
func (u APIFlowCreateRequestUnionParam) GetEnrollmentCriteria() (res apiFlowCreateRequestUnionParamEnrollmentCriteria) {
	if vt := u.OfAPIContactFlowCreateRequest; vt != nil {
		res.any = vt.EnrollmentCriteria.asAny()
	} else if vt := u.OfAPIPlatformFlowCreateRequest; vt != nil {
		res.any = vt.EnrollmentCriteria.asAny()
	}
	return
}

// Can have the runtime types [*APIListBasedEnrollmentCriteriaParam],
// [*APIEventBasedEnrollmentCriteriaParam], [*APIManualEnrollmentCriteriaParam]
type apiFlowCreateRequestUnionParamEnrollmentCriteria struct{ any }

// Use the following switch statement to get the type of the union:
//
//	switch u.AsAny().(type) {
//	case *automation.APIListBasedEnrollmentCriteriaParam:
//	case *automation.APIEventBasedEnrollmentCriteriaParam:
//	case *automation.APIManualEnrollmentCriteriaParam:
//	default:
//	    fmt.Errorf("not present")
//	}
func (u apiFlowCreateRequestUnionParamEnrollmentCriteria) AsAny() any { return u.any }

// Returns a pointer to the underlying variant's property, if present.
func (u apiFlowCreateRequestUnionParamEnrollmentCriteria) GetShouldReEnroll() *bool {
	switch vt := u.any.(type) {
	case *APIContactFlowCreateRequestEnrollmentCriteriaUnionParam:
		return vt.GetShouldReEnroll()
	case *APIPlatformFlowCreateRequestEnrollmentCriteriaUnionParam:
		return vt.GetShouldReEnroll()
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u apiFlowCreateRequestUnionParamEnrollmentCriteria) GetType() *string {
	switch vt := u.any.(type) {
	case *APIContactFlowCreateRequestEnrollmentCriteriaUnionParam:
		return vt.GetType()
	case *APIPlatformFlowCreateRequestEnrollmentCriteriaUnionParam:
		return vt.GetType()
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u apiFlowCreateRequestUnionParamEnrollmentCriteria) GetUnEnrollObjectsNotMeetingCriteria() *bool {
	switch vt := u.any.(type) {
	case *APIContactFlowCreateRequestEnrollmentCriteriaUnionParam:
		return vt.GetUnEnrollObjectsNotMeetingCriteria()
	case *APIPlatformFlowCreateRequestEnrollmentCriteriaUnionParam:
		return vt.GetUnEnrollObjectsNotMeetingCriteria()
	}
	return nil
}

// Returns a pointer to the underlying variant's ListFilterBranch property, if
// present.
func (u apiFlowCreateRequestUnionParamEnrollmentCriteria) GetListFilterBranch() *APIListBasedEnrollmentCriteriaListFilterBranchUnionParam {
	switch vt := u.any.(type) {
	case *APIContactFlowCreateRequestEnrollmentCriteriaUnionParam:
		return vt.GetListFilterBranch()
	case *APIPlatformFlowCreateRequestEnrollmentCriteriaUnionParam:
		return vt.GetListFilterBranch()
	}
	return nil
}

// Returns a pointer to the underlying variant's ReEnrollmentTriggersFilterBranches
// property, if present.
func (u apiFlowCreateRequestUnionParamEnrollmentCriteria) GetReEnrollmentTriggersFilterBranches() []APIListBasedEnrollmentCriteriaReEnrollmentTriggersFilterBranchUnionParam {
	switch vt := u.any.(type) {
	case *APIContactFlowCreateRequestEnrollmentCriteriaUnionParam:
		return vt.GetReEnrollmentTriggersFilterBranches()
	case *APIPlatformFlowCreateRequestEnrollmentCriteriaUnionParam:
		return vt.GetReEnrollmentTriggersFilterBranches()
	}
	return nil
}

// Returns a pointer to the underlying variant's EventFilterBranches property, if
// present.
func (u apiFlowCreateRequestUnionParamEnrollmentCriteria) GetEventFilterBranches() []shared.PublicUnifiedEventsFilterBranchParam {
	switch vt := u.any.(type) {
	case *APIContactFlowCreateRequestEnrollmentCriteriaUnionParam:
		return vt.GetEventFilterBranches()
	case *APIPlatformFlowCreateRequestEnrollmentCriteriaUnionParam:
		return vt.GetEventFilterBranches()
	}
	return nil
}

// Returns a pointer to the underlying variant's ListMembershipFilterBranches
// property, if present.
func (u apiFlowCreateRequestUnionParamEnrollmentCriteria) GetListMembershipFilterBranches() []APIEventBasedEnrollmentCriteriaListMembershipFilterBranchUnionParam {
	switch vt := u.any.(type) {
	case *APIContactFlowCreateRequestEnrollmentCriteriaUnionParam:
		return vt.GetListMembershipFilterBranches()
	case *APIPlatformFlowCreateRequestEnrollmentCriteriaUnionParam:
		return vt.GetListMembershipFilterBranches()
	}
	return nil
}

// Returns a pointer to the underlying variant's RefinementCriteria property, if
// present.
func (u apiFlowCreateRequestUnionParamEnrollmentCriteria) GetRefinementCriteria() *APIEventBasedEnrollmentCriteriaRefinementCriteriaUnionParam {
	switch vt := u.any.(type) {
	case *APIContactFlowCreateRequestEnrollmentCriteriaUnionParam:
		return vt.GetRefinementCriteria()
	case *APIPlatformFlowCreateRequestEnrollmentCriteriaUnionParam:
		return vt.GetRefinementCriteria()
	}
	return nil
}

// Returns a subunion which exports methods to access subproperties
//
// Or use AsAny() to get the underlying value
func (u APIFlowCreateRequestUnionParam) GetEnrollmentSchedule() (res apiFlowCreateRequestUnionParamEnrollmentSchedule) {
	if vt := u.OfAPIContactFlowCreateRequest; vt != nil {
		res.any = vt.EnrollmentSchedule.asAny()
	} else if vt := u.OfAPIPlatformFlowCreateRequest; vt != nil {
		res.any = vt.EnrollmentSchedule.asAny()
	}
	return
}

// Can have the runtime types [*APIDailyEnrollmentScheduleParam],
// [*APIWeeklyEnrollmentScheduleParam],
// [*APIMonthlySpecificDaysEnrollmentScheduleParam],
// [*APIMonthlyRelativeDaysEnrollmentScheduleParam],
// [*APIYearlyEnrollmentScheduleParam], [*APIPropertyBasedEnrollmentScheduleParam]
type apiFlowCreateRequestUnionParamEnrollmentSchedule struct{ any }

// Use the following switch statement to get the type of the union:
//
//	switch u.AsAny().(type) {
//	case *automation.APIDailyEnrollmentScheduleParam:
//	case *automation.APIWeeklyEnrollmentScheduleParam:
//	case *automation.APIMonthlySpecificDaysEnrollmentScheduleParam:
//	case *automation.APIMonthlyRelativeDaysEnrollmentScheduleParam:
//	case *automation.APIYearlyEnrollmentScheduleParam:
//	case *automation.APIPropertyBasedEnrollmentScheduleParam:
//	default:
//	    fmt.Errorf("not present")
//	}
func (u apiFlowCreateRequestUnionParamEnrollmentSchedule) AsAny() any { return u.any }

// Returns a pointer to the underlying variant's property, if present.
func (u apiFlowCreateRequestUnionParamEnrollmentSchedule) GetType() *string {
	switch vt := u.any.(type) {
	case *APIContactFlowCreateRequestEnrollmentScheduleUnionParam:
		return vt.GetType()
	case *APIPlatformFlowCreateRequestEnrollmentScheduleUnionParam:
		return vt.GetType()
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u apiFlowCreateRequestUnionParamEnrollmentSchedule) GetMonthlyRelativeDays() *string {
	switch vt := u.any.(type) {
	case *APIContactFlowCreateRequestEnrollmentScheduleUnionParam:
		return vt.GetMonthlyRelativeDays()
	case *APIPlatformFlowCreateRequestEnrollmentScheduleUnionParam:
		return vt.GetMonthlyRelativeDays()
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u apiFlowCreateRequestUnionParamEnrollmentSchedule) GetDayOfMonth() *int64 {
	switch vt := u.any.(type) {
	case *APIContactFlowCreateRequestEnrollmentScheduleUnionParam:
		return vt.GetDayOfMonth()
	case *APIPlatformFlowCreateRequestEnrollmentScheduleUnionParam:
		return vt.GetDayOfMonth()
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u apiFlowCreateRequestUnionParamEnrollmentSchedule) GetMonth() *string {
	switch vt := u.any.(type) {
	case *APIContactFlowCreateRequestEnrollmentScheduleUnionParam:
		return vt.GetMonth()
	case *APIPlatformFlowCreateRequestEnrollmentScheduleUnionParam:
		return vt.GetMonth()
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u apiFlowCreateRequestUnionParamEnrollmentSchedule) GetDateProperty() *string {
	switch vt := u.any.(type) {
	case *APIContactFlowCreateRequestEnrollmentScheduleUnionParam:
		return vt.GetDateProperty()
	case *APIPlatformFlowCreateRequestEnrollmentScheduleUnionParam:
		return vt.GetDateProperty()
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u apiFlowCreateRequestUnionParamEnrollmentSchedule) GetDaysDelta() *int64 {
	switch vt := u.any.(type) {
	case *APIContactFlowCreateRequestEnrollmentScheduleUnionParam:
		return vt.GetDaysDelta()
	case *APIPlatformFlowCreateRequestEnrollmentScheduleUnionParam:
		return vt.GetDaysDelta()
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u apiFlowCreateRequestUnionParamEnrollmentSchedule) GetYearly() *bool {
	switch vt := u.any.(type) {
	case *APIContactFlowCreateRequestEnrollmentScheduleUnionParam:
		return vt.GetYearly()
	case *APIPlatformFlowCreateRequestEnrollmentScheduleUnionParam:
		return vt.GetYearly()
	}
	return nil
}

// Returns a pointer to the underlying variant's TimeOfDay property, if present.
func (u apiFlowCreateRequestUnionParamEnrollmentSchedule) GetTimeOfDay() *APITimeOfDayParam {
	switch vt := u.any.(type) {
	case *APIContactFlowCreateRequestEnrollmentScheduleUnionParam:
		return vt.GetTimeOfDay()
	case *APIPlatformFlowCreateRequestEnrollmentScheduleUnionParam:
		return vt.GetTimeOfDay()
	}
	return nil
}

// Returns a pointer to the underlying variant's DaysOfWeek property, if present.
func (u apiFlowCreateRequestUnionParamEnrollmentSchedule) GetDaysOfWeek() []string {
	switch vt := u.any.(type) {
	case *APIContactFlowCreateRequestEnrollmentScheduleUnionParam:
		return vt.GetDaysOfWeek()
	case *APIPlatformFlowCreateRequestEnrollmentScheduleUnionParam:
		return vt.GetDaysOfWeek()
	}
	return nil
}

// Returns a pointer to the underlying variant's DaysOfMonth property, if present.
func (u apiFlowCreateRequestUnionParamEnrollmentSchedule) GetDaysOfMonth() []int64 {
	switch vt := u.any.(type) {
	case *APIContactFlowCreateRequestEnrollmentScheduleUnionParam:
		return vt.GetDaysOfMonth()
	case *APIPlatformFlowCreateRequestEnrollmentScheduleUnionParam:
		return vt.GetDaysOfMonth()
	}
	return nil
}

type APIFlowEmailCampaign struct {
	EmailCampaignID string `json:"emailCampaignId,required"`
	EmailContentID  string `json:"emailContentId,required"`
	FlowID          string `json:"flowId,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		EmailCampaignID respjson.Field
		EmailContentID  respjson.Field
		FlowID          respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r APIFlowEmailCampaign) RawJSON() string { return r.JSON.raw }
func (r *APIFlowEmailCampaign) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type APIFlowListing struct {
	// The unique ID for this flow. This is auto-generated when creating the flow.
	ID string `json:"id,required"`
	// The timestamp this flow was created.
	CreatedAt time.Time `json:"createdAt,required" format:"date-time"`
	// Deprecated. Will be removed.
	FlowType string `json:"flowType,required"`
	// This controls whether or not the flow is "enabled" if it's actively listening
	// for enrollment triggers and executing actions. If this is `false` the flow is
	// not accepting any enrollments or executing any actions.
	IsEnabled bool `json:"isEnabled,required"`
	// The CRM object type for objects that can be enrolled into this flow.
	ObjectTypeID string `json:"objectTypeId,required"`
	// Deprecated. Will be removed.
	RevisionID string `json:"revisionId,required"`
	// The timestamp this flow was last updated.
	UpdatedAt time.Time `json:"updatedAt,required" format:"date-time"`
	// The user-provided name for this flow. Names get auto-created for workflows that
	// are created without a name.
	Name string `json:"name"`
	// An optional unique key for this flow. This is only unique per-portal.
	Uuid string `json:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		CreatedAt    respjson.Field
		FlowType     respjson.Field
		IsEnabled    respjson.Field
		ObjectTypeID respjson.Field
		RevisionID   respjson.Field
		UpdatedAt    respjson.Field
		Name         respjson.Field
		Uuid         respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r APIFlowListing) RawJSON() string { return r.JSON.raw }
func (r *APIFlowListing) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type APIFlowPutRequestUnionParam struct {
	OfAPIContactFlowPutRequest  *APIContactFlowPutRequestParam  `json:",omitzero,inline"`
	OfAPIPlatformFlowPutRequest *APIPlatformFlowPutRequestParam `json:",omitzero,inline"`
	paramUnion
}

func (u APIFlowPutRequestUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfAPIContactFlowPutRequest, u.OfAPIPlatformFlowPutRequest)
}
func (u *APIFlowPutRequestUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *APIFlowPutRequestUnionParam) asAny() any {
	if !param.IsOmitted(u.OfAPIContactFlowPutRequest) {
		return u.OfAPIContactFlowPutRequest
	} else if !param.IsOmitted(u.OfAPIPlatformFlowPutRequest) {
		return u.OfAPIPlatformFlowPutRequest
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIFlowPutRequestUnionParam) GetCanEnrollFromSalesforce() *bool {
	if vt := u.OfAPIContactFlowPutRequest; vt != nil {
		return &vt.CanEnrollFromSalesforce
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIFlowPutRequestUnionParam) GetSuppressionListIDs() []int64 {
	if vt := u.OfAPIContactFlowPutRequest; vt != nil {
		return vt.SuppressionListIDs
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIFlowPutRequestUnionParam) GetEventAnchor() *APIContactFlowPutRequestEventAnchorUnionParam {
	if vt := u.OfAPIContactFlowPutRequest; vt != nil {
		return &vt.EventAnchor
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIFlowPutRequestUnionParam) GetGoalFilterBranch() *APIContactFlowPutRequestGoalFilterBranchUnionParam {
	if vt := u.OfAPIContactFlowPutRequest; vt != nil {
		return &vt.GoalFilterBranch
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIFlowPutRequestUnionParam) GetUnEnrollmentSetting() *APIUnEnrollmentSettingParam {
	if vt := u.OfAPIContactFlowPutRequest; vt != nil {
		return &vt.UnEnrollmentSetting
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIFlowPutRequestUnionParam) GetSuppressionFilterBranch() *APIPlatformFlowPutRequestSuppressionFilterBranchUnionParam {
	if vt := u.OfAPIPlatformFlowPutRequest; vt != nil {
		return &vt.SuppressionFilterBranch
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIFlowPutRequestUnionParam) GetIsEnabled() *bool {
	if vt := u.OfAPIContactFlowPutRequest; vt != nil {
		return (*bool)(&vt.IsEnabled)
	} else if vt := u.OfAPIPlatformFlowPutRequest; vt != nil {
		return (*bool)(&vt.IsEnabled)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIFlowPutRequestUnionParam) GetRevisionID() *string {
	if vt := u.OfAPIContactFlowPutRequest; vt != nil {
		return (*string)(&vt.RevisionID)
	} else if vt := u.OfAPIPlatformFlowPutRequest; vt != nil {
		return (*string)(&vt.RevisionID)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIFlowPutRequestUnionParam) GetType() *string {
	if vt := u.OfAPIContactFlowPutRequest; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfAPIPlatformFlowPutRequest; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIFlowPutRequestUnionParam) GetDescription() *string {
	if vt := u.OfAPIContactFlowPutRequest; vt != nil && vt.Description.Valid() {
		return &vt.Description.Value
	} else if vt := u.OfAPIPlatformFlowPutRequest; vt != nil && vt.Description.Valid() {
		return &vt.Description.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIFlowPutRequestUnionParam) GetName() *string {
	if vt := u.OfAPIContactFlowPutRequest; vt != nil && vt.Name.Valid() {
		return &vt.Name.Value
	} else if vt := u.OfAPIPlatformFlowPutRequest; vt != nil && vt.Name.Valid() {
		return &vt.Name.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIFlowPutRequestUnionParam) GetStartActionID() *string {
	if vt := u.OfAPIContactFlowPutRequest; vt != nil && vt.StartActionID.Valid() {
		return &vt.StartActionID.Value
	} else if vt := u.OfAPIPlatformFlowPutRequest; vt != nil && vt.StartActionID.Valid() {
		return &vt.StartActionID.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIFlowPutRequestUnionParam) GetUuid() *string {
	if vt := u.OfAPIContactFlowPutRequest; vt != nil && vt.Uuid.Valid() {
		return &vt.Uuid.Value
	} else if vt := u.OfAPIPlatformFlowPutRequest; vt != nil && vt.Uuid.Valid() {
		return &vt.Uuid.Value
	}
	return nil
}

// Returns a subunion which exports methods to access subproperties
//
// Or use AsAny() to get the underlying value
func (u APIFlowPutRequestUnionParam) GetActions() (res apiFlowPutRequestUnionParamActions) {
	if vt := u.OfAPIContactFlowPutRequest; vt != nil {
		res.any = &vt.Actions
	} else if vt := u.OfAPIPlatformFlowPutRequest; vt != nil {
		res.any = &vt.Actions
	}
	return
}

// Can have the runtime types [_[]APIContactFlowPutRequestActionUnionParam],
// [_[]APIPlatformFlowPutRequestActionUnionParam]
type apiFlowPutRequestUnionParamActions struct{ any }

// Use the following switch statement to get the type of the union:
//
//	switch u.AsAny().(type) {
//	case *[]automation.APIContactFlowPutRequestActionUnionParam:
//	case *[]automation.APIPlatformFlowPutRequestActionUnionParam:
//	default:
//	    fmt.Errorf("not present")
//	}
func (u apiFlowPutRequestUnionParamActions) AsAny() any { return u.any }

// Returns a pointer to the underlying variant's BlockedDates property, if present.
func (u APIFlowPutRequestUnionParam) GetBlockedDates() []APIBlockedDateParam {
	if vt := u.OfAPIContactFlowPutRequest; vt != nil {
		return vt.BlockedDates
	} else if vt := u.OfAPIPlatformFlowPutRequest; vt != nil {
		return vt.BlockedDates
	}
	return nil
}

// Returns a pointer to the underlying variant's CustomProperties property, if
// present.
func (u APIFlowPutRequestUnionParam) GetCustomProperties() map[string]string {
	if vt := u.OfAPIContactFlowPutRequest; vt != nil {
		return vt.CustomProperties
	} else if vt := u.OfAPIPlatformFlowPutRequest; vt != nil {
		return vt.CustomProperties
	}
	return nil
}

// Returns a pointer to the underlying variant's TimeWindows property, if present.
func (u APIFlowPutRequestUnionParam) GetTimeWindows() []APITimeWindowParam {
	if vt := u.OfAPIContactFlowPutRequest; vt != nil {
		return vt.TimeWindows
	} else if vt := u.OfAPIPlatformFlowPutRequest; vt != nil {
		return vt.TimeWindows
	}
	return nil
}

// Returns a subunion which exports methods to access subproperties
//
// Or use AsAny() to get the underlying value
func (u APIFlowPutRequestUnionParam) GetEnrollmentCriteria() (res apiFlowPutRequestUnionParamEnrollmentCriteria) {
	if vt := u.OfAPIContactFlowPutRequest; vt != nil {
		res.any = vt.EnrollmentCriteria.asAny()
	} else if vt := u.OfAPIPlatformFlowPutRequest; vt != nil {
		res.any = vt.EnrollmentCriteria.asAny()
	}
	return
}

// Can have the runtime types [*APIListBasedEnrollmentCriteriaParam],
// [*APIEventBasedEnrollmentCriteriaParam], [*APIManualEnrollmentCriteriaParam]
type apiFlowPutRequestUnionParamEnrollmentCriteria struct{ any }

// Use the following switch statement to get the type of the union:
//
//	switch u.AsAny().(type) {
//	case *automation.APIListBasedEnrollmentCriteriaParam:
//	case *automation.APIEventBasedEnrollmentCriteriaParam:
//	case *automation.APIManualEnrollmentCriteriaParam:
//	default:
//	    fmt.Errorf("not present")
//	}
func (u apiFlowPutRequestUnionParamEnrollmentCriteria) AsAny() any { return u.any }

// Returns a pointer to the underlying variant's property, if present.
func (u apiFlowPutRequestUnionParamEnrollmentCriteria) GetShouldReEnroll() *bool {
	switch vt := u.any.(type) {
	case *APIContactFlowPutRequestEnrollmentCriteriaUnionParam:
		return vt.GetShouldReEnroll()
	case *APIPlatformFlowPutRequestEnrollmentCriteriaUnionParam:
		return vt.GetShouldReEnroll()
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u apiFlowPutRequestUnionParamEnrollmentCriteria) GetType() *string {
	switch vt := u.any.(type) {
	case *APIContactFlowPutRequestEnrollmentCriteriaUnionParam:
		return vt.GetType()
	case *APIPlatformFlowPutRequestEnrollmentCriteriaUnionParam:
		return vt.GetType()
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u apiFlowPutRequestUnionParamEnrollmentCriteria) GetUnEnrollObjectsNotMeetingCriteria() *bool {
	switch vt := u.any.(type) {
	case *APIContactFlowPutRequestEnrollmentCriteriaUnionParam:
		return vt.GetUnEnrollObjectsNotMeetingCriteria()
	case *APIPlatformFlowPutRequestEnrollmentCriteriaUnionParam:
		return vt.GetUnEnrollObjectsNotMeetingCriteria()
	}
	return nil
}

// Returns a pointer to the underlying variant's ListFilterBranch property, if
// present.
func (u apiFlowPutRequestUnionParamEnrollmentCriteria) GetListFilterBranch() *APIListBasedEnrollmentCriteriaListFilterBranchUnionParam {
	switch vt := u.any.(type) {
	case *APIContactFlowPutRequestEnrollmentCriteriaUnionParam:
		return vt.GetListFilterBranch()
	case *APIPlatformFlowPutRequestEnrollmentCriteriaUnionParam:
		return vt.GetListFilterBranch()
	}
	return nil
}

// Returns a pointer to the underlying variant's ReEnrollmentTriggersFilterBranches
// property, if present.
func (u apiFlowPutRequestUnionParamEnrollmentCriteria) GetReEnrollmentTriggersFilterBranches() []APIListBasedEnrollmentCriteriaReEnrollmentTriggersFilterBranchUnionParam {
	switch vt := u.any.(type) {
	case *APIContactFlowPutRequestEnrollmentCriteriaUnionParam:
		return vt.GetReEnrollmentTriggersFilterBranches()
	case *APIPlatformFlowPutRequestEnrollmentCriteriaUnionParam:
		return vt.GetReEnrollmentTriggersFilterBranches()
	}
	return nil
}

// Returns a pointer to the underlying variant's EventFilterBranches property, if
// present.
func (u apiFlowPutRequestUnionParamEnrollmentCriteria) GetEventFilterBranches() []shared.PublicUnifiedEventsFilterBranchParam {
	switch vt := u.any.(type) {
	case *APIContactFlowPutRequestEnrollmentCriteriaUnionParam:
		return vt.GetEventFilterBranches()
	case *APIPlatformFlowPutRequestEnrollmentCriteriaUnionParam:
		return vt.GetEventFilterBranches()
	}
	return nil
}

// Returns a pointer to the underlying variant's ListMembershipFilterBranches
// property, if present.
func (u apiFlowPutRequestUnionParamEnrollmentCriteria) GetListMembershipFilterBranches() []APIEventBasedEnrollmentCriteriaListMembershipFilterBranchUnionParam {
	switch vt := u.any.(type) {
	case *APIContactFlowPutRequestEnrollmentCriteriaUnionParam:
		return vt.GetListMembershipFilterBranches()
	case *APIPlatformFlowPutRequestEnrollmentCriteriaUnionParam:
		return vt.GetListMembershipFilterBranches()
	}
	return nil
}

// Returns a pointer to the underlying variant's RefinementCriteria property, if
// present.
func (u apiFlowPutRequestUnionParamEnrollmentCriteria) GetRefinementCriteria() *APIEventBasedEnrollmentCriteriaRefinementCriteriaUnionParam {
	switch vt := u.any.(type) {
	case *APIContactFlowPutRequestEnrollmentCriteriaUnionParam:
		return vt.GetRefinementCriteria()
	case *APIPlatformFlowPutRequestEnrollmentCriteriaUnionParam:
		return vt.GetRefinementCriteria()
	}
	return nil
}

// Returns a subunion which exports methods to access subproperties
//
// Or use AsAny() to get the underlying value
func (u APIFlowPutRequestUnionParam) GetEnrollmentSchedule() (res apiFlowPutRequestUnionParamEnrollmentSchedule) {
	if vt := u.OfAPIContactFlowPutRequest; vt != nil {
		res.any = vt.EnrollmentSchedule.asAny()
	} else if vt := u.OfAPIPlatformFlowPutRequest; vt != nil {
		res.any = vt.EnrollmentSchedule.asAny()
	}
	return
}

// Can have the runtime types [*APIDailyEnrollmentScheduleParam],
// [*APIWeeklyEnrollmentScheduleParam],
// [*APIMonthlySpecificDaysEnrollmentScheduleParam],
// [*APIMonthlyRelativeDaysEnrollmentScheduleParam],
// [*APIYearlyEnrollmentScheduleParam], [*APIPropertyBasedEnrollmentScheduleParam]
type apiFlowPutRequestUnionParamEnrollmentSchedule struct{ any }

// Use the following switch statement to get the type of the union:
//
//	switch u.AsAny().(type) {
//	case *automation.APIDailyEnrollmentScheduleParam:
//	case *automation.APIWeeklyEnrollmentScheduleParam:
//	case *automation.APIMonthlySpecificDaysEnrollmentScheduleParam:
//	case *automation.APIMonthlyRelativeDaysEnrollmentScheduleParam:
//	case *automation.APIYearlyEnrollmentScheduleParam:
//	case *automation.APIPropertyBasedEnrollmentScheduleParam:
//	default:
//	    fmt.Errorf("not present")
//	}
func (u apiFlowPutRequestUnionParamEnrollmentSchedule) AsAny() any { return u.any }

// Returns a pointer to the underlying variant's property, if present.
func (u apiFlowPutRequestUnionParamEnrollmentSchedule) GetType() *string {
	switch vt := u.any.(type) {
	case *APIContactFlowPutRequestEnrollmentScheduleUnionParam:
		return vt.GetType()
	case *APIPlatformFlowPutRequestEnrollmentScheduleUnionParam:
		return vt.GetType()
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u apiFlowPutRequestUnionParamEnrollmentSchedule) GetMonthlyRelativeDays() *string {
	switch vt := u.any.(type) {
	case *APIContactFlowPutRequestEnrollmentScheduleUnionParam:
		return vt.GetMonthlyRelativeDays()
	case *APIPlatformFlowPutRequestEnrollmentScheduleUnionParam:
		return vt.GetMonthlyRelativeDays()
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u apiFlowPutRequestUnionParamEnrollmentSchedule) GetDayOfMonth() *int64 {
	switch vt := u.any.(type) {
	case *APIContactFlowPutRequestEnrollmentScheduleUnionParam:
		return vt.GetDayOfMonth()
	case *APIPlatformFlowPutRequestEnrollmentScheduleUnionParam:
		return vt.GetDayOfMonth()
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u apiFlowPutRequestUnionParamEnrollmentSchedule) GetMonth() *string {
	switch vt := u.any.(type) {
	case *APIContactFlowPutRequestEnrollmentScheduleUnionParam:
		return vt.GetMonth()
	case *APIPlatformFlowPutRequestEnrollmentScheduleUnionParam:
		return vt.GetMonth()
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u apiFlowPutRequestUnionParamEnrollmentSchedule) GetDateProperty() *string {
	switch vt := u.any.(type) {
	case *APIContactFlowPutRequestEnrollmentScheduleUnionParam:
		return vt.GetDateProperty()
	case *APIPlatformFlowPutRequestEnrollmentScheduleUnionParam:
		return vt.GetDateProperty()
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u apiFlowPutRequestUnionParamEnrollmentSchedule) GetDaysDelta() *int64 {
	switch vt := u.any.(type) {
	case *APIContactFlowPutRequestEnrollmentScheduleUnionParam:
		return vt.GetDaysDelta()
	case *APIPlatformFlowPutRequestEnrollmentScheduleUnionParam:
		return vt.GetDaysDelta()
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u apiFlowPutRequestUnionParamEnrollmentSchedule) GetYearly() *bool {
	switch vt := u.any.(type) {
	case *APIContactFlowPutRequestEnrollmentScheduleUnionParam:
		return vt.GetYearly()
	case *APIPlatformFlowPutRequestEnrollmentScheduleUnionParam:
		return vt.GetYearly()
	}
	return nil
}

// Returns a pointer to the underlying variant's TimeOfDay property, if present.
func (u apiFlowPutRequestUnionParamEnrollmentSchedule) GetTimeOfDay() *APITimeOfDayParam {
	switch vt := u.any.(type) {
	case *APIContactFlowPutRequestEnrollmentScheduleUnionParam:
		return vt.GetTimeOfDay()
	case *APIPlatformFlowPutRequestEnrollmentScheduleUnionParam:
		return vt.GetTimeOfDay()
	}
	return nil
}

// Returns a pointer to the underlying variant's DaysOfWeek property, if present.
func (u apiFlowPutRequestUnionParamEnrollmentSchedule) GetDaysOfWeek() []string {
	switch vt := u.any.(type) {
	case *APIContactFlowPutRequestEnrollmentScheduleUnionParam:
		return vt.GetDaysOfWeek()
	case *APIPlatformFlowPutRequestEnrollmentScheduleUnionParam:
		return vt.GetDaysOfWeek()
	}
	return nil
}

// Returns a pointer to the underlying variant's DaysOfMonth property, if present.
func (u apiFlowPutRequestUnionParamEnrollmentSchedule) GetDaysOfMonth() []int64 {
	switch vt := u.any.(type) {
	case *APIContactFlowPutRequestEnrollmentScheduleUnionParam:
		return vt.GetDaysOfMonth()
	case *APIPlatformFlowPutRequestEnrollmentScheduleUnionParam:
		return vt.GetDaysOfMonth()
	}
	return nil
}

type APIIncrementValue struct {
	// The amount be which to increment
	IncrementAmount float64 `json:"incrementAmount,required"`
	// This is the type of input value. This can be one of: "FIELD_DATA",
	// "OBJECT_PROPERTY", "STATIC_VALUE", "RELATIVE_DATETIME", "TIMESTAMP",
	// "INCREMENT", "FETCHED_OBJECT_PROPERTY", "APPEND_OBJECT_PROPERTY",
	// "STATIC_APPEND_VALUE", "ENROLLMENT_EVENT_PROPERTY"
	//
	// Any of "INCREMENT".
	Type APIIncrementValueType `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		IncrementAmount respjson.Field
		Type            respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r APIIncrementValue) RawJSON() string { return r.JSON.raw }
func (r *APIIncrementValue) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this APIIncrementValue to a APIIncrementValueParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// APIIncrementValueParam.Overrides()
func (r APIIncrementValue) ToParam() APIIncrementValueParam {
	return param.Override[APIIncrementValueParam](json.RawMessage(r.RawJSON()))
}

// This is the type of input value. This can be one of: "FIELD_DATA",
// "OBJECT_PROPERTY", "STATIC_VALUE", "RELATIVE_DATETIME", "TIMESTAMP",
// "INCREMENT", "FETCHED_OBJECT_PROPERTY", "APPEND_OBJECT_PROPERTY",
// "STATIC_APPEND_VALUE", "ENROLLMENT_EVENT_PROPERTY"
type APIIncrementValueType string

const (
	APIIncrementValueTypeIncrement APIIncrementValueType = "INCREMENT"
)

// The properties IncrementAmount, Type are required.
type APIIncrementValueParam struct {
	// The amount be which to increment
	IncrementAmount float64 `json:"incrementAmount,required"`
	// This is the type of input value. This can be one of: "FIELD_DATA",
	// "OBJECT_PROPERTY", "STATIC_VALUE", "RELATIVE_DATETIME", "TIMESTAMP",
	// "INCREMENT", "FETCHED_OBJECT_PROPERTY", "APPEND_OBJECT_PROPERTY",
	// "STATIC_APPEND_VALUE", "ENROLLMENT_EVENT_PROPERTY"
	//
	// Any of "INCREMENT".
	Type APIIncrementValueType `json:"type,omitzero,required"`
	paramObj
}

func (r APIIncrementValueParam) MarshalJSON() (data []byte, err error) {
	type shadow APIIncrementValueParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *APIIncrementValueParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type APIInputVariable struct {
	Name  string                     `json:"name,required"`
	Value APIInputVariableValueUnion `json:"value,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r APIInputVariable) RawJSON() string { return r.JSON.raw }
func (r *APIInputVariable) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this APIInputVariable to a APIInputVariableParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// APIInputVariableParam.Overrides()
func (r APIInputVariable) ToParam() APIInputVariableParam {
	return param.Override[APIInputVariableParam](json.RawMessage(r.RawJSON()))
}

// APIInputVariableValueUnion contains all possible properties and values from
// [APIActionDataValue], [APIObjectPropertyValue], [APIStaticValue],
// [APIRelativeDateTimeValue], [APITimestampValue], [APIIncrementValue],
// [APIFetchedObjectPropertyValue], [APIAppendObjectPropertyValue],
// [APIStaticAppendValue], [APIEnrollmentEventPropertyValue].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type APIInputVariableValueUnion struct {
	// This field is from variant [APIActionDataValue].
	ActionID string `json:"actionId"`
	// This field is from variant [APIActionDataValue].
	DataKey string `json:"dataKey"`
	Type    string `json:"type"`
	// This field is from variant [APIObjectPropertyValue].
	PropertyName string `json:"propertyName"`
	// This field is from variant [APIStaticValue].
	StaticValue string `json:"staticValue"`
	// This field is from variant [APIRelativeDateTimeValue].
	TimeDelay APITimeDelay `json:"timeDelay"`
	// This field is from variant [APITimestampValue].
	TimestampType APITimestampValueTimestampType `json:"timestampType"`
	// This field is from variant [APIIncrementValue].
	IncrementAmount float64 `json:"incrementAmount"`
	// This field is from variant [APIFetchedObjectPropertyValue].
	PropertyToken string `json:"propertyToken"`
	// This field is from variant [APIAppendObjectPropertyValue].
	AppendPropertyName string `json:"appendPropertyName"`
	// This field is from variant [APIStaticAppendValue].
	StaticAppendValue string `json:"staticAppendValue"`
	// This field is from variant [APIEnrollmentEventPropertyValue].
	EnrollmentEventPropertyToken string `json:"enrollmentEventPropertyToken"`
	JSON                         struct {
		ActionID                     respjson.Field
		DataKey                      respjson.Field
		Type                         respjson.Field
		PropertyName                 respjson.Field
		StaticValue                  respjson.Field
		TimeDelay                    respjson.Field
		TimestampType                respjson.Field
		IncrementAmount              respjson.Field
		PropertyToken                respjson.Field
		AppendPropertyName           respjson.Field
		StaticAppendValue            respjson.Field
		EnrollmentEventPropertyToken respjson.Field
		raw                          string
	} `json:"-"`
}

func (u APIInputVariableValueUnion) AsFieldData() (v APIActionDataValue) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u APIInputVariableValueUnion) AsObjectProperty() (v APIObjectPropertyValue) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u APIInputVariableValueUnion) AsStaticValue() (v APIStaticValue) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u APIInputVariableValueUnion) AsRelativeDatetime() (v APIRelativeDateTimeValue) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u APIInputVariableValueUnion) AsTimestamp() (v APITimestampValue) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u APIInputVariableValueUnion) AsIncrement() (v APIIncrementValue) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u APIInputVariableValueUnion) AsFetchedObjectProperty() (v APIFetchedObjectPropertyValue) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u APIInputVariableValueUnion) AsAppendObjectProperty() (v APIAppendObjectPropertyValue) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u APIInputVariableValueUnion) AsStaticAppendValue() (v APIStaticAppendValue) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u APIInputVariableValueUnion) AsEnrollmentEventProperty() (v APIEnrollmentEventPropertyValue) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u APIInputVariableValueUnion) RawJSON() string { return u.JSON.raw }

func (r *APIInputVariableValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Name, Value are required.
type APIInputVariableParam struct {
	Name  string                          `json:"name,required"`
	Value APIInputVariableValueUnionParam `json:"value,omitzero,required"`
	paramObj
}

func (r APIInputVariableParam) MarshalJSON() (data []byte, err error) {
	type shadow APIInputVariableParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *APIInputVariableParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type APIInputVariableValueUnionParam struct {
	OfFieldData               *APIActionDataValueParam              `json:",omitzero,inline"`
	OfObjectProperty          *APIObjectPropertyValueParam          `json:",omitzero,inline"`
	OfStaticValue             *APIStaticValueParam                  `json:",omitzero,inline"`
	OfRelativeDatetime        *APIRelativeDateTimeValueParam        `json:",omitzero,inline"`
	OfTimestamp               *APITimestampValueParam               `json:",omitzero,inline"`
	OfIncrement               *APIIncrementValueParam               `json:",omitzero,inline"`
	OfFetchedObjectProperty   *APIFetchedObjectPropertyValueParam   `json:",omitzero,inline"`
	OfAppendObjectProperty    *APIAppendObjectPropertyValueParam    `json:",omitzero,inline"`
	OfStaticAppendValue       *APIStaticAppendValueParam            `json:",omitzero,inline"`
	OfEnrollmentEventProperty *APIEnrollmentEventPropertyValueParam `json:",omitzero,inline"`
	paramUnion
}

func (u APIInputVariableValueUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfFieldData,
		u.OfObjectProperty,
		u.OfStaticValue,
		u.OfRelativeDatetime,
		u.OfTimestamp,
		u.OfIncrement,
		u.OfFetchedObjectProperty,
		u.OfAppendObjectProperty,
		u.OfStaticAppendValue,
		u.OfEnrollmentEventProperty)
}
func (u *APIInputVariableValueUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *APIInputVariableValueUnionParam) asAny() any {
	if !param.IsOmitted(u.OfFieldData) {
		return u.OfFieldData
	} else if !param.IsOmitted(u.OfObjectProperty) {
		return u.OfObjectProperty
	} else if !param.IsOmitted(u.OfStaticValue) {
		return u.OfStaticValue
	} else if !param.IsOmitted(u.OfRelativeDatetime) {
		return u.OfRelativeDatetime
	} else if !param.IsOmitted(u.OfTimestamp) {
		return u.OfTimestamp
	} else if !param.IsOmitted(u.OfIncrement) {
		return u.OfIncrement
	} else if !param.IsOmitted(u.OfFetchedObjectProperty) {
		return u.OfFetchedObjectProperty
	} else if !param.IsOmitted(u.OfAppendObjectProperty) {
		return u.OfAppendObjectProperty
	} else if !param.IsOmitted(u.OfStaticAppendValue) {
		return u.OfStaticAppendValue
	} else if !param.IsOmitted(u.OfEnrollmentEventProperty) {
		return u.OfEnrollmentEventProperty
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIInputVariableValueUnionParam) GetActionID() *string {
	if vt := u.OfFieldData; vt != nil {
		return &vt.ActionID
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIInputVariableValueUnionParam) GetDataKey() *string {
	if vt := u.OfFieldData; vt != nil {
		return &vt.DataKey
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIInputVariableValueUnionParam) GetPropertyName() *string {
	if vt := u.OfObjectProperty; vt != nil {
		return &vt.PropertyName
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIInputVariableValueUnionParam) GetStaticValue() *string {
	if vt := u.OfStaticValue; vt != nil {
		return &vt.StaticValue
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIInputVariableValueUnionParam) GetTimeDelay() *APITimeDelayParam {
	if vt := u.OfRelativeDatetime; vt != nil {
		return &vt.TimeDelay
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIInputVariableValueUnionParam) GetTimestampType() *string {
	if vt := u.OfTimestamp; vt != nil {
		return (*string)(&vt.TimestampType)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIInputVariableValueUnionParam) GetIncrementAmount() *float64 {
	if vt := u.OfIncrement; vt != nil {
		return &vt.IncrementAmount
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIInputVariableValueUnionParam) GetPropertyToken() *string {
	if vt := u.OfFetchedObjectProperty; vt != nil {
		return &vt.PropertyToken
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIInputVariableValueUnionParam) GetAppendPropertyName() *string {
	if vt := u.OfAppendObjectProperty; vt != nil {
		return &vt.AppendPropertyName
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIInputVariableValueUnionParam) GetStaticAppendValue() *string {
	if vt := u.OfStaticAppendValue; vt != nil {
		return &vt.StaticAppendValue
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIInputVariableValueUnionParam) GetEnrollmentEventPropertyToken() *string {
	if vt := u.OfEnrollmentEventProperty; vt != nil {
		return &vt.EnrollmentEventPropertyToken
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIInputVariableValueUnionParam) GetType() *string {
	if vt := u.OfFieldData; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfObjectProperty; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfStaticValue; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfRelativeDatetime; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfTimestamp; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfIncrement; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfFetchedObjectProperty; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfAppendObjectProperty; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfStaticAppendValue; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfEnrollmentEventProperty; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

type APIListBasedEnrollmentCriteria struct {
	// The list filter branch that represents the enrollment trigger to this flow.
	ListFilterBranch APIListBasedEnrollmentCriteriaListFilterBranchUnion `json:"listFilterBranch,required"`
	// A list of filter branches to listen for in order to re-enroll objects into this
	// workflow.
	ReEnrollmentTriggersFilterBranches []APIListBasedEnrollmentCriteriaReEnrollmentTriggersFilterBranchUnion `json:"reEnrollmentTriggersFilterBranches,required"`
	// Whether or not the same object can enroll in this workflow twice.
	ShouldReEnroll bool `json:"shouldReEnroll,required"`
	// The type of enrollment criteria this is, this can be "LIST_BASED",
	// "EVENT_BASED", or "MANUAL".
	//
	// Any of "LIST_BASED".
	Type APIListBasedEnrollmentCriteriaType `json:"type,required"`
	// Whether or not to remove objects from this workflow if they stop meeting the
	// enrollment criteria.
	UnEnrollObjectsNotMeetingCriteria bool `json:"unEnrollObjectsNotMeetingCriteria,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ListFilterBranch                   respjson.Field
		ReEnrollmentTriggersFilterBranches respjson.Field
		ShouldReEnroll                     respjson.Field
		Type                               respjson.Field
		UnEnrollObjectsNotMeetingCriteria  respjson.Field
		ExtraFields                        map[string]respjson.Field
		raw                                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r APIListBasedEnrollmentCriteria) RawJSON() string { return r.JSON.raw }
func (r *APIListBasedEnrollmentCriteria) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this APIListBasedEnrollmentCriteria to a
// APIListBasedEnrollmentCriteriaParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// APIListBasedEnrollmentCriteriaParam.Overrides()
func (r APIListBasedEnrollmentCriteria) ToParam() APIListBasedEnrollmentCriteriaParam {
	return param.Override[APIListBasedEnrollmentCriteriaParam](json.RawMessage(r.RawJSON()))
}

// APIListBasedEnrollmentCriteriaListFilterBranchUnion contains all possible
// properties and values from [shared.PublicOrFilterBranch],
// [shared.PublicAndFilterBranch], [shared.PublicNotAllFilterBranch],
// [shared.PublicNotAnyFilterBranch], [shared.PublicRestrictedFilterBranch],
// [shared.PublicUnifiedEventsFilterBranch],
// [shared.PublicPropertyAssociationFilterBranch],
// [shared.PublicAssociationFilterBranch].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type APIListBasedEnrollmentCriteriaListFilterBranchUnion struct {
	// This field is a union of [[]shared.PublicOrFilterBranchFilterBranchUnion],
	// [[]shared.PublicAndFilterBranchFilterBranchUnion],
	// [[]shared.PublicNotAllFilterBranchFilterBranchUnion],
	// [[]shared.PublicNotAnyFilterBranchFilterBranchUnion],
	// [[]shared.PublicRestrictedFilterBranchFilterBranchUnion],
	// [[]shared.PublicUnifiedEventsFilterBranchFilterBranchUnion],
	// [[]shared.PublicPropertyAssociationFilterBranchFilterBranchUnion],
	// [[]shared.PublicAssociationFilterBranchFilterBranchUnion]
	FilterBranches       APIListBasedEnrollmentCriteriaListFilterBranchUnionFilterBranches `json:"filterBranches"`
	FilterBranchOperator string                                                            `json:"filterBranchOperator"`
	FilterBranchType     string                                                            `json:"filterBranchType"`
	// This field is a union of [[]shared.PublicOrFilterBranchFilterUnion],
	// [[]shared.PublicAndFilterBranchFilterUnion],
	// [[]shared.PublicNotAllFilterBranchFilterUnion],
	// [[]shared.PublicNotAnyFilterBranchFilterUnion],
	// [[]shared.PublicRestrictedFilterBranchFilterUnion],
	// [[]shared.PublicUnifiedEventsFilterBranchFilterUnion],
	// [[]shared.PublicPropertyAssociationFilterBranchFilterUnion],
	// [[]shared.PublicAssociationFilterBranchFilterUnion]
	Filters APIListBasedEnrollmentCriteriaListFilterBranchUnionFilters `json:"filters"`
	// This field is from variant [shared.PublicUnifiedEventsFilterBranch].
	EventTypeID string `json:"eventTypeId"`
	Operator    string `json:"operator"`
	// This field is from variant [shared.PublicUnifiedEventsFilterBranch].
	CoalescingRefineBy shared.PublicUnifiedEventsFilterBranchCoalescingRefineByUnion `json:"coalescingRefineBy"`
	ObjectTypeID       string                                                        `json:"objectTypeId"`
	// This field is from variant [shared.PublicPropertyAssociationFilterBranch].
	PropertyWithObjectID string `json:"propertyWithObjectId"`
	// This field is from variant [shared.PublicAssociationFilterBranch].
	AssociationCategory string `json:"associationCategory"`
	// This field is from variant [shared.PublicAssociationFilterBranch].
	AssociationTypeID int64 `json:"associationTypeId"`
	JSON              struct {
		FilterBranches       respjson.Field
		FilterBranchOperator respjson.Field
		FilterBranchType     respjson.Field
		Filters              respjson.Field
		EventTypeID          respjson.Field
		Operator             respjson.Field
		CoalescingRefineBy   respjson.Field
		ObjectTypeID         respjson.Field
		PropertyWithObjectID respjson.Field
		AssociationCategory  respjson.Field
		AssociationTypeID    respjson.Field
		raw                  string
	} `json:"-"`
}

func (u APIListBasedEnrollmentCriteriaListFilterBranchUnion) AsOr() (v shared.PublicOrFilterBranch) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u APIListBasedEnrollmentCriteriaListFilterBranchUnion) AsAnd() (v shared.PublicAndFilterBranch) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u APIListBasedEnrollmentCriteriaListFilterBranchUnion) AsNotAll() (v shared.PublicNotAllFilterBranch) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u APIListBasedEnrollmentCriteriaListFilterBranchUnion) AsNotAny() (v shared.PublicNotAnyFilterBranch) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u APIListBasedEnrollmentCriteriaListFilterBranchUnion) AsRestricted() (v shared.PublicRestrictedFilterBranch) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u APIListBasedEnrollmentCriteriaListFilterBranchUnion) AsUnifiedEvents() (v shared.PublicUnifiedEventsFilterBranch) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u APIListBasedEnrollmentCriteriaListFilterBranchUnion) AsPropertyAssociation() (v shared.PublicPropertyAssociationFilterBranch) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u APIListBasedEnrollmentCriteriaListFilterBranchUnion) AsAssociation() (v shared.PublicAssociationFilterBranch) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u APIListBasedEnrollmentCriteriaListFilterBranchUnion) RawJSON() string { return u.JSON.raw }

func (r *APIListBasedEnrollmentCriteriaListFilterBranchUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// APIListBasedEnrollmentCriteriaListFilterBranchUnionFilterBranches is an implicit
// subunion of [APIListBasedEnrollmentCriteriaListFilterBranchUnion].
// APIListBasedEnrollmentCriteriaListFilterBranchUnionFilterBranches provides
// convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [APIListBasedEnrollmentCriteriaListFilterBranchUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfPublicOrFilterBranchFilterBranches
// OfPublicAndFilterBranchFilterBranches OfPublicNotAllFilterBranchFilterBranches
// OfPublicNotAnyFilterBranchFilterBranches
// OfPublicRestrictedFilterBranchFilterBranches
// OfPublicUnifiedEventsFilterBranchFilterBranches
// OfPublicPropertyAssociationFilterBranchFilterBranches
// OfPublicAssociationFilterBranchFilterBranches]
type APIListBasedEnrollmentCriteriaListFilterBranchUnionFilterBranches struct {
	// This field will be present if the value is a
	// [[]shared.PublicOrFilterBranchFilterBranchUnion] instead of an object.
	OfPublicOrFilterBranchFilterBranches []shared.PublicOrFilterBranchFilterBranchUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]shared.PublicAndFilterBranchFilterBranchUnion] instead of an object.
	OfPublicAndFilterBranchFilterBranches []shared.PublicAndFilterBranchFilterBranchUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]shared.PublicNotAllFilterBranchFilterBranchUnion] instead of an object.
	OfPublicNotAllFilterBranchFilterBranches []shared.PublicNotAllFilterBranchFilterBranchUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]shared.PublicNotAnyFilterBranchFilterBranchUnion] instead of an object.
	OfPublicNotAnyFilterBranchFilterBranches []shared.PublicNotAnyFilterBranchFilterBranchUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]shared.PublicRestrictedFilterBranchFilterBranchUnion] instead of an object.
	OfPublicRestrictedFilterBranchFilterBranches []shared.PublicRestrictedFilterBranchFilterBranchUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]shared.PublicUnifiedEventsFilterBranchFilterBranchUnion] instead of an
	// object.
	OfPublicUnifiedEventsFilterBranchFilterBranches []shared.PublicUnifiedEventsFilterBranchFilterBranchUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]shared.PublicPropertyAssociationFilterBranchFilterBranchUnion] instead of an
	// object.
	OfPublicPropertyAssociationFilterBranchFilterBranches []shared.PublicPropertyAssociationFilterBranchFilterBranchUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]shared.PublicAssociationFilterBranchFilterBranchUnion] instead of an object.
	OfPublicAssociationFilterBranchFilterBranches []shared.PublicAssociationFilterBranchFilterBranchUnion `json:",inline"`
	JSON                                          struct {
		OfPublicOrFilterBranchFilterBranches                  respjson.Field
		OfPublicAndFilterBranchFilterBranches                 respjson.Field
		OfPublicNotAllFilterBranchFilterBranches              respjson.Field
		OfPublicNotAnyFilterBranchFilterBranches              respjson.Field
		OfPublicRestrictedFilterBranchFilterBranches          respjson.Field
		OfPublicUnifiedEventsFilterBranchFilterBranches       respjson.Field
		OfPublicPropertyAssociationFilterBranchFilterBranches respjson.Field
		OfPublicAssociationFilterBranchFilterBranches         respjson.Field
		raw                                                   string
	} `json:"-"`
}

func (r *APIListBasedEnrollmentCriteriaListFilterBranchUnionFilterBranches) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// APIListBasedEnrollmentCriteriaListFilterBranchUnionFilters is an implicit
// subunion of [APIListBasedEnrollmentCriteriaListFilterBranchUnion].
// APIListBasedEnrollmentCriteriaListFilterBranchUnionFilters provides convenient
// access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [APIListBasedEnrollmentCriteriaListFilterBranchUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfPublicOrFilterBranchFilters OfPublicAndFilterBranchFilters
// OfPublicNotAllFilterBranchFilters OfPublicNotAnyFilterBranchFilters
// OfPublicRestrictedFilterBranchFilters OfPublicUnifiedEventsFilterBranchFilters
// OfPublicPropertyAssociationFilterBranchFilters
// OfPublicAssociationFilterBranchFilters]
type APIListBasedEnrollmentCriteriaListFilterBranchUnionFilters struct {
	// This field will be present if the value is a
	// [[]shared.PublicOrFilterBranchFilterUnion] instead of an object.
	OfPublicOrFilterBranchFilters []shared.PublicOrFilterBranchFilterUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]shared.PublicAndFilterBranchFilterUnion] instead of an object.
	OfPublicAndFilterBranchFilters []shared.PublicAndFilterBranchFilterUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]shared.PublicNotAllFilterBranchFilterUnion] instead of an object.
	OfPublicNotAllFilterBranchFilters []shared.PublicNotAllFilterBranchFilterUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]shared.PublicNotAnyFilterBranchFilterUnion] instead of an object.
	OfPublicNotAnyFilterBranchFilters []shared.PublicNotAnyFilterBranchFilterUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]shared.PublicRestrictedFilterBranchFilterUnion] instead of an object.
	OfPublicRestrictedFilterBranchFilters []shared.PublicRestrictedFilterBranchFilterUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]shared.PublicUnifiedEventsFilterBranchFilterUnion] instead of an object.
	OfPublicUnifiedEventsFilterBranchFilters []shared.PublicUnifiedEventsFilterBranchFilterUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]shared.PublicPropertyAssociationFilterBranchFilterUnion] instead of an
	// object.
	OfPublicPropertyAssociationFilterBranchFilters []shared.PublicPropertyAssociationFilterBranchFilterUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]shared.PublicAssociationFilterBranchFilterUnion] instead of an object.
	OfPublicAssociationFilterBranchFilters []shared.PublicAssociationFilterBranchFilterUnion `json:",inline"`
	JSON                                   struct {
		OfPublicOrFilterBranchFilters                  respjson.Field
		OfPublicAndFilterBranchFilters                 respjson.Field
		OfPublicNotAllFilterBranchFilters              respjson.Field
		OfPublicNotAnyFilterBranchFilters              respjson.Field
		OfPublicRestrictedFilterBranchFilters          respjson.Field
		OfPublicUnifiedEventsFilterBranchFilters       respjson.Field
		OfPublicPropertyAssociationFilterBranchFilters respjson.Field
		OfPublicAssociationFilterBranchFilters         respjson.Field
		raw                                            string
	} `json:"-"`
}

func (r *APIListBasedEnrollmentCriteriaListFilterBranchUnionFilters) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// APIListBasedEnrollmentCriteriaReEnrollmentTriggersFilterBranchUnion contains all
// possible properties and values from [shared.PublicOrFilterBranch],
// [shared.PublicAndFilterBranch], [shared.PublicNotAllFilterBranch],
// [shared.PublicNotAnyFilterBranch], [shared.PublicRestrictedFilterBranch],
// [shared.PublicUnifiedEventsFilterBranch],
// [shared.PublicPropertyAssociationFilterBranch],
// [shared.PublicAssociationFilterBranch].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type APIListBasedEnrollmentCriteriaReEnrollmentTriggersFilterBranchUnion struct {
	// This field is a union of [[]shared.PublicOrFilterBranchFilterBranchUnion],
	// [[]shared.PublicAndFilterBranchFilterBranchUnion],
	// [[]shared.PublicNotAllFilterBranchFilterBranchUnion],
	// [[]shared.PublicNotAnyFilterBranchFilterBranchUnion],
	// [[]shared.PublicRestrictedFilterBranchFilterBranchUnion],
	// [[]shared.PublicUnifiedEventsFilterBranchFilterBranchUnion],
	// [[]shared.PublicPropertyAssociationFilterBranchFilterBranchUnion],
	// [[]shared.PublicAssociationFilterBranchFilterBranchUnion]
	FilterBranches       APIListBasedEnrollmentCriteriaReEnrollmentTriggersFilterBranchUnionFilterBranches `json:"filterBranches"`
	FilterBranchOperator string                                                                            `json:"filterBranchOperator"`
	FilterBranchType     string                                                                            `json:"filterBranchType"`
	// This field is a union of [[]shared.PublicOrFilterBranchFilterUnion],
	// [[]shared.PublicAndFilterBranchFilterUnion],
	// [[]shared.PublicNotAllFilterBranchFilterUnion],
	// [[]shared.PublicNotAnyFilterBranchFilterUnion],
	// [[]shared.PublicRestrictedFilterBranchFilterUnion],
	// [[]shared.PublicUnifiedEventsFilterBranchFilterUnion],
	// [[]shared.PublicPropertyAssociationFilterBranchFilterUnion],
	// [[]shared.PublicAssociationFilterBranchFilterUnion]
	Filters APIListBasedEnrollmentCriteriaReEnrollmentTriggersFilterBranchUnionFilters `json:"filters"`
	// This field is from variant [shared.PublicUnifiedEventsFilterBranch].
	EventTypeID string `json:"eventTypeId"`
	Operator    string `json:"operator"`
	// This field is from variant [shared.PublicUnifiedEventsFilterBranch].
	CoalescingRefineBy shared.PublicUnifiedEventsFilterBranchCoalescingRefineByUnion `json:"coalescingRefineBy"`
	ObjectTypeID       string                                                        `json:"objectTypeId"`
	// This field is from variant [shared.PublicPropertyAssociationFilterBranch].
	PropertyWithObjectID string `json:"propertyWithObjectId"`
	// This field is from variant [shared.PublicAssociationFilterBranch].
	AssociationCategory string `json:"associationCategory"`
	// This field is from variant [shared.PublicAssociationFilterBranch].
	AssociationTypeID int64 `json:"associationTypeId"`
	JSON              struct {
		FilterBranches       respjson.Field
		FilterBranchOperator respjson.Field
		FilterBranchType     respjson.Field
		Filters              respjson.Field
		EventTypeID          respjson.Field
		Operator             respjson.Field
		CoalescingRefineBy   respjson.Field
		ObjectTypeID         respjson.Field
		PropertyWithObjectID respjson.Field
		AssociationCategory  respjson.Field
		AssociationTypeID    respjson.Field
		raw                  string
	} `json:"-"`
}

func (u APIListBasedEnrollmentCriteriaReEnrollmentTriggersFilterBranchUnion) AsOr() (v shared.PublicOrFilterBranch) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u APIListBasedEnrollmentCriteriaReEnrollmentTriggersFilterBranchUnion) AsAnd() (v shared.PublicAndFilterBranch) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u APIListBasedEnrollmentCriteriaReEnrollmentTriggersFilterBranchUnion) AsNotAll() (v shared.PublicNotAllFilterBranch) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u APIListBasedEnrollmentCriteriaReEnrollmentTriggersFilterBranchUnion) AsNotAny() (v shared.PublicNotAnyFilterBranch) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u APIListBasedEnrollmentCriteriaReEnrollmentTriggersFilterBranchUnion) AsRestricted() (v shared.PublicRestrictedFilterBranch) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u APIListBasedEnrollmentCriteriaReEnrollmentTriggersFilterBranchUnion) AsUnifiedEvents() (v shared.PublicUnifiedEventsFilterBranch) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u APIListBasedEnrollmentCriteriaReEnrollmentTriggersFilterBranchUnion) AsPropertyAssociation() (v shared.PublicPropertyAssociationFilterBranch) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u APIListBasedEnrollmentCriteriaReEnrollmentTriggersFilterBranchUnion) AsAssociation() (v shared.PublicAssociationFilterBranch) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u APIListBasedEnrollmentCriteriaReEnrollmentTriggersFilterBranchUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *APIListBasedEnrollmentCriteriaReEnrollmentTriggersFilterBranchUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// APIListBasedEnrollmentCriteriaReEnrollmentTriggersFilterBranchUnionFilterBranches
// is an implicit subunion of
// [APIListBasedEnrollmentCriteriaReEnrollmentTriggersFilterBranchUnion].
// APIListBasedEnrollmentCriteriaReEnrollmentTriggersFilterBranchUnionFilterBranches
// provides convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [APIListBasedEnrollmentCriteriaReEnrollmentTriggersFilterBranchUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfPublicOrFilterBranchFilterBranches
// OfPublicAndFilterBranchFilterBranches OfPublicNotAllFilterBranchFilterBranches
// OfPublicNotAnyFilterBranchFilterBranches
// OfPublicRestrictedFilterBranchFilterBranches
// OfPublicUnifiedEventsFilterBranchFilterBranches
// OfPublicPropertyAssociationFilterBranchFilterBranches
// OfPublicAssociationFilterBranchFilterBranches]
type APIListBasedEnrollmentCriteriaReEnrollmentTriggersFilterBranchUnionFilterBranches struct {
	// This field will be present if the value is a
	// [[]shared.PublicOrFilterBranchFilterBranchUnion] instead of an object.
	OfPublicOrFilterBranchFilterBranches []shared.PublicOrFilterBranchFilterBranchUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]shared.PublicAndFilterBranchFilterBranchUnion] instead of an object.
	OfPublicAndFilterBranchFilterBranches []shared.PublicAndFilterBranchFilterBranchUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]shared.PublicNotAllFilterBranchFilterBranchUnion] instead of an object.
	OfPublicNotAllFilterBranchFilterBranches []shared.PublicNotAllFilterBranchFilterBranchUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]shared.PublicNotAnyFilterBranchFilterBranchUnion] instead of an object.
	OfPublicNotAnyFilterBranchFilterBranches []shared.PublicNotAnyFilterBranchFilterBranchUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]shared.PublicRestrictedFilterBranchFilterBranchUnion] instead of an object.
	OfPublicRestrictedFilterBranchFilterBranches []shared.PublicRestrictedFilterBranchFilterBranchUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]shared.PublicUnifiedEventsFilterBranchFilterBranchUnion] instead of an
	// object.
	OfPublicUnifiedEventsFilterBranchFilterBranches []shared.PublicUnifiedEventsFilterBranchFilterBranchUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]shared.PublicPropertyAssociationFilterBranchFilterBranchUnion] instead of an
	// object.
	OfPublicPropertyAssociationFilterBranchFilterBranches []shared.PublicPropertyAssociationFilterBranchFilterBranchUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]shared.PublicAssociationFilterBranchFilterBranchUnion] instead of an object.
	OfPublicAssociationFilterBranchFilterBranches []shared.PublicAssociationFilterBranchFilterBranchUnion `json:",inline"`
	JSON                                          struct {
		OfPublicOrFilterBranchFilterBranches                  respjson.Field
		OfPublicAndFilterBranchFilterBranches                 respjson.Field
		OfPublicNotAllFilterBranchFilterBranches              respjson.Field
		OfPublicNotAnyFilterBranchFilterBranches              respjson.Field
		OfPublicRestrictedFilterBranchFilterBranches          respjson.Field
		OfPublicUnifiedEventsFilterBranchFilterBranches       respjson.Field
		OfPublicPropertyAssociationFilterBranchFilterBranches respjson.Field
		OfPublicAssociationFilterBranchFilterBranches         respjson.Field
		raw                                                   string
	} `json:"-"`
}

func (r *APIListBasedEnrollmentCriteriaReEnrollmentTriggersFilterBranchUnionFilterBranches) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// APIListBasedEnrollmentCriteriaReEnrollmentTriggersFilterBranchUnionFilters is an
// implicit subunion of
// [APIListBasedEnrollmentCriteriaReEnrollmentTriggersFilterBranchUnion].
// APIListBasedEnrollmentCriteriaReEnrollmentTriggersFilterBranchUnionFilters
// provides convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [APIListBasedEnrollmentCriteriaReEnrollmentTriggersFilterBranchUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfPublicOrFilterBranchFilters OfPublicAndFilterBranchFilters
// OfPublicNotAllFilterBranchFilters OfPublicNotAnyFilterBranchFilters
// OfPublicRestrictedFilterBranchFilters OfPublicUnifiedEventsFilterBranchFilters
// OfPublicPropertyAssociationFilterBranchFilters
// OfPublicAssociationFilterBranchFilters]
type APIListBasedEnrollmentCriteriaReEnrollmentTriggersFilterBranchUnionFilters struct {
	// This field will be present if the value is a
	// [[]shared.PublicOrFilterBranchFilterUnion] instead of an object.
	OfPublicOrFilterBranchFilters []shared.PublicOrFilterBranchFilterUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]shared.PublicAndFilterBranchFilterUnion] instead of an object.
	OfPublicAndFilterBranchFilters []shared.PublicAndFilterBranchFilterUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]shared.PublicNotAllFilterBranchFilterUnion] instead of an object.
	OfPublicNotAllFilterBranchFilters []shared.PublicNotAllFilterBranchFilterUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]shared.PublicNotAnyFilterBranchFilterUnion] instead of an object.
	OfPublicNotAnyFilterBranchFilters []shared.PublicNotAnyFilterBranchFilterUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]shared.PublicRestrictedFilterBranchFilterUnion] instead of an object.
	OfPublicRestrictedFilterBranchFilters []shared.PublicRestrictedFilterBranchFilterUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]shared.PublicUnifiedEventsFilterBranchFilterUnion] instead of an object.
	OfPublicUnifiedEventsFilterBranchFilters []shared.PublicUnifiedEventsFilterBranchFilterUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]shared.PublicPropertyAssociationFilterBranchFilterUnion] instead of an
	// object.
	OfPublicPropertyAssociationFilterBranchFilters []shared.PublicPropertyAssociationFilterBranchFilterUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]shared.PublicAssociationFilterBranchFilterUnion] instead of an object.
	OfPublicAssociationFilterBranchFilters []shared.PublicAssociationFilterBranchFilterUnion `json:",inline"`
	JSON                                   struct {
		OfPublicOrFilterBranchFilters                  respjson.Field
		OfPublicAndFilterBranchFilters                 respjson.Field
		OfPublicNotAllFilterBranchFilters              respjson.Field
		OfPublicNotAnyFilterBranchFilters              respjson.Field
		OfPublicRestrictedFilterBranchFilters          respjson.Field
		OfPublicUnifiedEventsFilterBranchFilters       respjson.Field
		OfPublicPropertyAssociationFilterBranchFilters respjson.Field
		OfPublicAssociationFilterBranchFilters         respjson.Field
		raw                                            string
	} `json:"-"`
}

func (r *APIListBasedEnrollmentCriteriaReEnrollmentTriggersFilterBranchUnionFilters) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The type of enrollment criteria this is, this can be "LIST_BASED",
// "EVENT_BASED", or "MANUAL".
type APIListBasedEnrollmentCriteriaType string

const (
	APIListBasedEnrollmentCriteriaTypeListBased APIListBasedEnrollmentCriteriaType = "LIST_BASED"
)

// The properties ListFilterBranch, ReEnrollmentTriggersFilterBranches,
// ShouldReEnroll, Type, UnEnrollObjectsNotMeetingCriteria are required.
type APIListBasedEnrollmentCriteriaParam struct {
	// The list filter branch that represents the enrollment trigger to this flow.
	ListFilterBranch APIListBasedEnrollmentCriteriaListFilterBranchUnionParam `json:"listFilterBranch,omitzero,required"`
	// A list of filter branches to listen for in order to re-enroll objects into this
	// workflow.
	ReEnrollmentTriggersFilterBranches []APIListBasedEnrollmentCriteriaReEnrollmentTriggersFilterBranchUnionParam `json:"reEnrollmentTriggersFilterBranches,omitzero,required"`
	// Whether or not the same object can enroll in this workflow twice.
	ShouldReEnroll bool `json:"shouldReEnroll,required"`
	// The type of enrollment criteria this is, this can be "LIST_BASED",
	// "EVENT_BASED", or "MANUAL".
	//
	// Any of "LIST_BASED".
	Type APIListBasedEnrollmentCriteriaType `json:"type,omitzero,required"`
	// Whether or not to remove objects from this workflow if they stop meeting the
	// enrollment criteria.
	UnEnrollObjectsNotMeetingCriteria bool `json:"unEnrollObjectsNotMeetingCriteria,required"`
	paramObj
}

func (r APIListBasedEnrollmentCriteriaParam) MarshalJSON() (data []byte, err error) {
	type shadow APIListBasedEnrollmentCriteriaParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *APIListBasedEnrollmentCriteriaParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type APIListBasedEnrollmentCriteriaListFilterBranchUnionParam struct {
	OfOr                  *shared.PublicOrFilterBranchParam                  `json:",omitzero,inline"`
	OfAnd                 *shared.PublicAndFilterBranchParam                 `json:",omitzero,inline"`
	OfNotAll              *shared.PublicNotAllFilterBranchParam              `json:",omitzero,inline"`
	OfNotAny              *shared.PublicNotAnyFilterBranchParam              `json:",omitzero,inline"`
	OfRestricted          *shared.PublicRestrictedFilterBranchParam          `json:",omitzero,inline"`
	OfUnifiedEvents       *shared.PublicUnifiedEventsFilterBranchParam       `json:",omitzero,inline"`
	OfPropertyAssociation *shared.PublicPropertyAssociationFilterBranchParam `json:",omitzero,inline"`
	OfAssociation         *shared.PublicAssociationFilterBranchParam         `json:",omitzero,inline"`
	paramUnion
}

func (u APIListBasedEnrollmentCriteriaListFilterBranchUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfOr,
		u.OfAnd,
		u.OfNotAll,
		u.OfNotAny,
		u.OfRestricted,
		u.OfUnifiedEvents,
		u.OfPropertyAssociation,
		u.OfAssociation)
}
func (u *APIListBasedEnrollmentCriteriaListFilterBranchUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *APIListBasedEnrollmentCriteriaListFilterBranchUnionParam) asAny() any {
	if !param.IsOmitted(u.OfOr) {
		return u.OfOr
	} else if !param.IsOmitted(u.OfAnd) {
		return u.OfAnd
	} else if !param.IsOmitted(u.OfNotAll) {
		return u.OfNotAll
	} else if !param.IsOmitted(u.OfNotAny) {
		return u.OfNotAny
	} else if !param.IsOmitted(u.OfRestricted) {
		return u.OfRestricted
	} else if !param.IsOmitted(u.OfUnifiedEvents) {
		return u.OfUnifiedEvents
	} else if !param.IsOmitted(u.OfPropertyAssociation) {
		return u.OfPropertyAssociation
	} else if !param.IsOmitted(u.OfAssociation) {
		return u.OfAssociation
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIListBasedEnrollmentCriteriaListFilterBranchUnionParam) GetEventTypeID() *string {
	if vt := u.OfUnifiedEvents; vt != nil {
		return &vt.EventTypeID
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIListBasedEnrollmentCriteriaListFilterBranchUnionParam) GetCoalescingRefineBy() *shared.PublicUnifiedEventsFilterBranchCoalescingRefineByUnionParam {
	if vt := u.OfUnifiedEvents; vt != nil {
		return &vt.CoalescingRefineBy
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIListBasedEnrollmentCriteriaListFilterBranchUnionParam) GetPropertyWithObjectID() *string {
	if vt := u.OfPropertyAssociation; vt != nil {
		return &vt.PropertyWithObjectID
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIListBasedEnrollmentCriteriaListFilterBranchUnionParam) GetAssociationCategory() *string {
	if vt := u.OfAssociation; vt != nil {
		return &vt.AssociationCategory
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIListBasedEnrollmentCriteriaListFilterBranchUnionParam) GetAssociationTypeID() *int64 {
	if vt := u.OfAssociation; vt != nil {
		return &vt.AssociationTypeID
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIListBasedEnrollmentCriteriaListFilterBranchUnionParam) GetFilterBranchOperator() *string {
	if vt := u.OfOr; vt != nil {
		return (*string)(&vt.FilterBranchOperator)
	} else if vt := u.OfAnd; vt != nil {
		return (*string)(&vt.FilterBranchOperator)
	} else if vt := u.OfNotAll; vt != nil {
		return (*string)(&vt.FilterBranchOperator)
	} else if vt := u.OfNotAny; vt != nil {
		return (*string)(&vt.FilterBranchOperator)
	} else if vt := u.OfRestricted; vt != nil {
		return (*string)(&vt.FilterBranchOperator)
	} else if vt := u.OfUnifiedEvents; vt != nil {
		return (*string)(&vt.FilterBranchOperator)
	} else if vt := u.OfPropertyAssociation; vt != nil {
		return (*string)(&vt.FilterBranchOperator)
	} else if vt := u.OfAssociation; vt != nil {
		return (*string)(&vt.FilterBranchOperator)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIListBasedEnrollmentCriteriaListFilterBranchUnionParam) GetFilterBranchType() *string {
	if vt := u.OfOr; vt != nil {
		return (*string)(&vt.FilterBranchType)
	} else if vt := u.OfAnd; vt != nil {
		return (*string)(&vt.FilterBranchType)
	} else if vt := u.OfNotAll; vt != nil {
		return (*string)(&vt.FilterBranchType)
	} else if vt := u.OfNotAny; vt != nil {
		return (*string)(&vt.FilterBranchType)
	} else if vt := u.OfRestricted; vt != nil {
		return (*string)(&vt.FilterBranchType)
	} else if vt := u.OfUnifiedEvents; vt != nil {
		return (*string)(&vt.FilterBranchType)
	} else if vt := u.OfPropertyAssociation; vt != nil {
		return (*string)(&vt.FilterBranchType)
	} else if vt := u.OfAssociation; vt != nil {
		return (*string)(&vt.FilterBranchType)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIListBasedEnrollmentCriteriaListFilterBranchUnionParam) GetOperator() *string {
	if vt := u.OfUnifiedEvents; vt != nil {
		return (*string)(&vt.Operator)
	} else if vt := u.OfPropertyAssociation; vt != nil {
		return (*string)(&vt.Operator)
	} else if vt := u.OfAssociation; vt != nil {
		return (*string)(&vt.Operator)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIListBasedEnrollmentCriteriaListFilterBranchUnionParam) GetObjectTypeID() *string {
	if vt := u.OfPropertyAssociation; vt != nil {
		return (*string)(&vt.ObjectTypeID)
	} else if vt := u.OfAssociation; vt != nil {
		return (*string)(&vt.ObjectTypeID)
	}
	return nil
}

// Returns a subunion which exports methods to access subproperties
//
// Or use AsAny() to get the underlying value
func (u APIListBasedEnrollmentCriteriaListFilterBranchUnionParam) GetFilterBranches() (res apiListBasedEnrollmentCriteriaListFilterBranchUnionParamFilterBranches) {
	if vt := u.OfOr; vt != nil {
		res.any = &vt.FilterBranches
	} else if vt := u.OfAnd; vt != nil {
		res.any = &vt.FilterBranches
	} else if vt := u.OfNotAll; vt != nil {
		res.any = &vt.FilterBranches
	} else if vt := u.OfNotAny; vt != nil {
		res.any = &vt.FilterBranches
	} else if vt := u.OfRestricted; vt != nil {
		res.any = &vt.FilterBranches
	} else if vt := u.OfUnifiedEvents; vt != nil {
		res.any = &vt.FilterBranches
	} else if vt := u.OfPropertyAssociation; vt != nil {
		res.any = &vt.FilterBranches
	} else if vt := u.OfAssociation; vt != nil {
		res.any = &vt.FilterBranches
	}
	return
}

// Can have the runtime types
// [_[]shared.PublicOrFilterBranchFilterBranchUnionParam],
// [_[]shared.PublicAndFilterBranchFilterBranchUnionParam],
// [_[]shared.PublicNotAllFilterBranchFilterBranchUnionParam],
// [_[]shared.PublicNotAnyFilterBranchFilterBranchUnionParam],
// [_[]shared.PublicRestrictedFilterBranchFilterBranchUnionParam],
// [_[]shared.PublicUnifiedEventsFilterBranchFilterBranchUnionParam],
// [_[]shared.PublicPropertyAssociationFilterBranchFilterBranchUnionParam],
// [_[]shared.PublicAssociationFilterBranchFilterBranchUnionParam]
type apiListBasedEnrollmentCriteriaListFilterBranchUnionParamFilterBranches struct{ any }

// Use the following switch statement to get the type of the union:
//
//	switch u.AsAny().(type) {
//	case *[]shared.PublicOrFilterBranchFilterBranchUnionParam:
//	case *[]shared.PublicAndFilterBranchFilterBranchUnionParam:
//	case *[]shared.PublicNotAllFilterBranchFilterBranchUnionParam:
//	case *[]shared.PublicNotAnyFilterBranchFilterBranchUnionParam:
//	case *[]shared.PublicRestrictedFilterBranchFilterBranchUnionParam:
//	case *[]shared.PublicUnifiedEventsFilterBranchFilterBranchUnionParam:
//	case *[]shared.PublicPropertyAssociationFilterBranchFilterBranchUnionParam:
//	case *[]shared.PublicAssociationFilterBranchFilterBranchUnionParam:
//	default:
//	    fmt.Errorf("not present")
//	}
func (u apiListBasedEnrollmentCriteriaListFilterBranchUnionParamFilterBranches) AsAny() any {
	return u.any
}

// Returns a subunion which exports methods to access subproperties
//
// Or use AsAny() to get the underlying value
func (u APIListBasedEnrollmentCriteriaListFilterBranchUnionParam) GetFilters() (res apiListBasedEnrollmentCriteriaListFilterBranchUnionParamFilters) {
	if vt := u.OfOr; vt != nil {
		res.any = &vt.Filters
	} else if vt := u.OfAnd; vt != nil {
		res.any = &vt.Filters
	} else if vt := u.OfNotAll; vt != nil {
		res.any = &vt.Filters
	} else if vt := u.OfNotAny; vt != nil {
		res.any = &vt.Filters
	} else if vt := u.OfRestricted; vt != nil {
		res.any = &vt.Filters
	} else if vt := u.OfUnifiedEvents; vt != nil {
		res.any = &vt.Filters
	} else if vt := u.OfPropertyAssociation; vt != nil {
		res.any = &vt.Filters
	} else if vt := u.OfAssociation; vt != nil {
		res.any = &vt.Filters
	}
	return
}

// Can have the runtime types [_[]shared.PublicOrFilterBranchFilterUnionParam],
// [_[]shared.PublicAndFilterBranchFilterUnionParam],
// [_[]shared.PublicNotAllFilterBranchFilterUnionParam],
// [_[]shared.PublicNotAnyFilterBranchFilterUnionParam],
// [_[]shared.PublicRestrictedFilterBranchFilterUnionParam],
// [_[]shared.PublicUnifiedEventsFilterBranchFilterUnionParam],
// [_[]shared.PublicPropertyAssociationFilterBranchFilterUnionParam],
// [_[]shared.PublicAssociationFilterBranchFilterUnionParam]
type apiListBasedEnrollmentCriteriaListFilterBranchUnionParamFilters struct{ any }

// Use the following switch statement to get the type of the union:
//
//	switch u.AsAny().(type) {
//	case *[]shared.PublicOrFilterBranchFilterUnionParam:
//	case *[]shared.PublicAndFilterBranchFilterUnionParam:
//	case *[]shared.PublicNotAllFilterBranchFilterUnionParam:
//	case *[]shared.PublicNotAnyFilterBranchFilterUnionParam:
//	case *[]shared.PublicRestrictedFilterBranchFilterUnionParam:
//	case *[]shared.PublicUnifiedEventsFilterBranchFilterUnionParam:
//	case *[]shared.PublicPropertyAssociationFilterBranchFilterUnionParam:
//	case *[]shared.PublicAssociationFilterBranchFilterUnionParam:
//	default:
//	    fmt.Errorf("not present")
//	}
func (u apiListBasedEnrollmentCriteriaListFilterBranchUnionParamFilters) AsAny() any { return u.any }

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type APIListBasedEnrollmentCriteriaReEnrollmentTriggersFilterBranchUnionParam struct {
	OfOr                  *shared.PublicOrFilterBranchParam                  `json:",omitzero,inline"`
	OfAnd                 *shared.PublicAndFilterBranchParam                 `json:",omitzero,inline"`
	OfNotAll              *shared.PublicNotAllFilterBranchParam              `json:",omitzero,inline"`
	OfNotAny              *shared.PublicNotAnyFilterBranchParam              `json:",omitzero,inline"`
	OfRestricted          *shared.PublicRestrictedFilterBranchParam          `json:",omitzero,inline"`
	OfUnifiedEvents       *shared.PublicUnifiedEventsFilterBranchParam       `json:",omitzero,inline"`
	OfPropertyAssociation *shared.PublicPropertyAssociationFilterBranchParam `json:",omitzero,inline"`
	OfAssociation         *shared.PublicAssociationFilterBranchParam         `json:",omitzero,inline"`
	paramUnion
}

func (u APIListBasedEnrollmentCriteriaReEnrollmentTriggersFilterBranchUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfOr,
		u.OfAnd,
		u.OfNotAll,
		u.OfNotAny,
		u.OfRestricted,
		u.OfUnifiedEvents,
		u.OfPropertyAssociation,
		u.OfAssociation)
}
func (u *APIListBasedEnrollmentCriteriaReEnrollmentTriggersFilterBranchUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *APIListBasedEnrollmentCriteriaReEnrollmentTriggersFilterBranchUnionParam) asAny() any {
	if !param.IsOmitted(u.OfOr) {
		return u.OfOr
	} else if !param.IsOmitted(u.OfAnd) {
		return u.OfAnd
	} else if !param.IsOmitted(u.OfNotAll) {
		return u.OfNotAll
	} else if !param.IsOmitted(u.OfNotAny) {
		return u.OfNotAny
	} else if !param.IsOmitted(u.OfRestricted) {
		return u.OfRestricted
	} else if !param.IsOmitted(u.OfUnifiedEvents) {
		return u.OfUnifiedEvents
	} else if !param.IsOmitted(u.OfPropertyAssociation) {
		return u.OfPropertyAssociation
	} else if !param.IsOmitted(u.OfAssociation) {
		return u.OfAssociation
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIListBasedEnrollmentCriteriaReEnrollmentTriggersFilterBranchUnionParam) GetEventTypeID() *string {
	if vt := u.OfUnifiedEvents; vt != nil {
		return &vt.EventTypeID
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIListBasedEnrollmentCriteriaReEnrollmentTriggersFilterBranchUnionParam) GetCoalescingRefineBy() *shared.PublicUnifiedEventsFilterBranchCoalescingRefineByUnionParam {
	if vt := u.OfUnifiedEvents; vt != nil {
		return &vt.CoalescingRefineBy
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIListBasedEnrollmentCriteriaReEnrollmentTriggersFilterBranchUnionParam) GetPropertyWithObjectID() *string {
	if vt := u.OfPropertyAssociation; vt != nil {
		return &vt.PropertyWithObjectID
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIListBasedEnrollmentCriteriaReEnrollmentTriggersFilterBranchUnionParam) GetAssociationCategory() *string {
	if vt := u.OfAssociation; vt != nil {
		return &vt.AssociationCategory
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIListBasedEnrollmentCriteriaReEnrollmentTriggersFilterBranchUnionParam) GetAssociationTypeID() *int64 {
	if vt := u.OfAssociation; vt != nil {
		return &vt.AssociationTypeID
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIListBasedEnrollmentCriteriaReEnrollmentTriggersFilterBranchUnionParam) GetFilterBranchOperator() *string {
	if vt := u.OfOr; vt != nil {
		return (*string)(&vt.FilterBranchOperator)
	} else if vt := u.OfAnd; vt != nil {
		return (*string)(&vt.FilterBranchOperator)
	} else if vt := u.OfNotAll; vt != nil {
		return (*string)(&vt.FilterBranchOperator)
	} else if vt := u.OfNotAny; vt != nil {
		return (*string)(&vt.FilterBranchOperator)
	} else if vt := u.OfRestricted; vt != nil {
		return (*string)(&vt.FilterBranchOperator)
	} else if vt := u.OfUnifiedEvents; vt != nil {
		return (*string)(&vt.FilterBranchOperator)
	} else if vt := u.OfPropertyAssociation; vt != nil {
		return (*string)(&vt.FilterBranchOperator)
	} else if vt := u.OfAssociation; vt != nil {
		return (*string)(&vt.FilterBranchOperator)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIListBasedEnrollmentCriteriaReEnrollmentTriggersFilterBranchUnionParam) GetFilterBranchType() *string {
	if vt := u.OfOr; vt != nil {
		return (*string)(&vt.FilterBranchType)
	} else if vt := u.OfAnd; vt != nil {
		return (*string)(&vt.FilterBranchType)
	} else if vt := u.OfNotAll; vt != nil {
		return (*string)(&vt.FilterBranchType)
	} else if vt := u.OfNotAny; vt != nil {
		return (*string)(&vt.FilterBranchType)
	} else if vt := u.OfRestricted; vt != nil {
		return (*string)(&vt.FilterBranchType)
	} else if vt := u.OfUnifiedEvents; vt != nil {
		return (*string)(&vt.FilterBranchType)
	} else if vt := u.OfPropertyAssociation; vt != nil {
		return (*string)(&vt.FilterBranchType)
	} else if vt := u.OfAssociation; vt != nil {
		return (*string)(&vt.FilterBranchType)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIListBasedEnrollmentCriteriaReEnrollmentTriggersFilterBranchUnionParam) GetOperator() *string {
	if vt := u.OfUnifiedEvents; vt != nil {
		return (*string)(&vt.Operator)
	} else if vt := u.OfPropertyAssociation; vt != nil {
		return (*string)(&vt.Operator)
	} else if vt := u.OfAssociation; vt != nil {
		return (*string)(&vt.Operator)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIListBasedEnrollmentCriteriaReEnrollmentTriggersFilterBranchUnionParam) GetObjectTypeID() *string {
	if vt := u.OfPropertyAssociation; vt != nil {
		return (*string)(&vt.ObjectTypeID)
	} else if vt := u.OfAssociation; vt != nil {
		return (*string)(&vt.ObjectTypeID)
	}
	return nil
}

// Returns a subunion which exports methods to access subproperties
//
// Or use AsAny() to get the underlying value
func (u APIListBasedEnrollmentCriteriaReEnrollmentTriggersFilterBranchUnionParam) GetFilterBranches() (res apiListBasedEnrollmentCriteriaReEnrollmentTriggersFilterBranchUnionParamFilterBranches) {
	if vt := u.OfOr; vt != nil {
		res.any = &vt.FilterBranches
	} else if vt := u.OfAnd; vt != nil {
		res.any = &vt.FilterBranches
	} else if vt := u.OfNotAll; vt != nil {
		res.any = &vt.FilterBranches
	} else if vt := u.OfNotAny; vt != nil {
		res.any = &vt.FilterBranches
	} else if vt := u.OfRestricted; vt != nil {
		res.any = &vt.FilterBranches
	} else if vt := u.OfUnifiedEvents; vt != nil {
		res.any = &vt.FilterBranches
	} else if vt := u.OfPropertyAssociation; vt != nil {
		res.any = &vt.FilterBranches
	} else if vt := u.OfAssociation; vt != nil {
		res.any = &vt.FilterBranches
	}
	return
}

// Can have the runtime types
// [_[]shared.PublicOrFilterBranchFilterBranchUnionParam],
// [_[]shared.PublicAndFilterBranchFilterBranchUnionParam],
// [_[]shared.PublicNotAllFilterBranchFilterBranchUnionParam],
// [_[]shared.PublicNotAnyFilterBranchFilterBranchUnionParam],
// [_[]shared.PublicRestrictedFilterBranchFilterBranchUnionParam],
// [_[]shared.PublicUnifiedEventsFilterBranchFilterBranchUnionParam],
// [_[]shared.PublicPropertyAssociationFilterBranchFilterBranchUnionParam],
// [_[]shared.PublicAssociationFilterBranchFilterBranchUnionParam]
type apiListBasedEnrollmentCriteriaReEnrollmentTriggersFilterBranchUnionParamFilterBranches struct{ any }

// Use the following switch statement to get the type of the union:
//
//	switch u.AsAny().(type) {
//	case *[]shared.PublicOrFilterBranchFilterBranchUnionParam:
//	case *[]shared.PublicAndFilterBranchFilterBranchUnionParam:
//	case *[]shared.PublicNotAllFilterBranchFilterBranchUnionParam:
//	case *[]shared.PublicNotAnyFilterBranchFilterBranchUnionParam:
//	case *[]shared.PublicRestrictedFilterBranchFilterBranchUnionParam:
//	case *[]shared.PublicUnifiedEventsFilterBranchFilterBranchUnionParam:
//	case *[]shared.PublicPropertyAssociationFilterBranchFilterBranchUnionParam:
//	case *[]shared.PublicAssociationFilterBranchFilterBranchUnionParam:
//	default:
//	    fmt.Errorf("not present")
//	}
func (u apiListBasedEnrollmentCriteriaReEnrollmentTriggersFilterBranchUnionParamFilterBranches) AsAny() any {
	return u.any
}

// Returns a subunion which exports methods to access subproperties
//
// Or use AsAny() to get the underlying value
func (u APIListBasedEnrollmentCriteriaReEnrollmentTriggersFilterBranchUnionParam) GetFilters() (res apiListBasedEnrollmentCriteriaReEnrollmentTriggersFilterBranchUnionParamFilters) {
	if vt := u.OfOr; vt != nil {
		res.any = &vt.Filters
	} else if vt := u.OfAnd; vt != nil {
		res.any = &vt.Filters
	} else if vt := u.OfNotAll; vt != nil {
		res.any = &vt.Filters
	} else if vt := u.OfNotAny; vt != nil {
		res.any = &vt.Filters
	} else if vt := u.OfRestricted; vt != nil {
		res.any = &vt.Filters
	} else if vt := u.OfUnifiedEvents; vt != nil {
		res.any = &vt.Filters
	} else if vt := u.OfPropertyAssociation; vt != nil {
		res.any = &vt.Filters
	} else if vt := u.OfAssociation; vt != nil {
		res.any = &vt.Filters
	}
	return
}

// Can have the runtime types [_[]shared.PublicOrFilterBranchFilterUnionParam],
// [_[]shared.PublicAndFilterBranchFilterUnionParam],
// [_[]shared.PublicNotAllFilterBranchFilterUnionParam],
// [_[]shared.PublicNotAnyFilterBranchFilterUnionParam],
// [_[]shared.PublicRestrictedFilterBranchFilterUnionParam],
// [_[]shared.PublicUnifiedEventsFilterBranchFilterUnionParam],
// [_[]shared.PublicPropertyAssociationFilterBranchFilterUnionParam],
// [_[]shared.PublicAssociationFilterBranchFilterUnionParam]
type apiListBasedEnrollmentCriteriaReEnrollmentTriggersFilterBranchUnionParamFilters struct{ any }

// Use the following switch statement to get the type of the union:
//
//	switch u.AsAny().(type) {
//	case *[]shared.PublicOrFilterBranchFilterUnionParam:
//	case *[]shared.PublicAndFilterBranchFilterUnionParam:
//	case *[]shared.PublicNotAllFilterBranchFilterUnionParam:
//	case *[]shared.PublicNotAnyFilterBranchFilterUnionParam:
//	case *[]shared.PublicRestrictedFilterBranchFilterUnionParam:
//	case *[]shared.PublicUnifiedEventsFilterBranchFilterUnionParam:
//	case *[]shared.PublicPropertyAssociationFilterBranchFilterUnionParam:
//	case *[]shared.PublicAssociationFilterBranchFilterUnionParam:
//	default:
//	    fmt.Errorf("not present")
//	}
func (u apiListBasedEnrollmentCriteriaReEnrollmentTriggersFilterBranchUnionParamFilters) AsAny() any {
	return u.any
}

type APIListBranch struct {
	// The name of this branch
	BranchName string        `json:"branchName"`
	Connection APIConnection `json:"connection"`
	// The list criteria that determine when to execute this branch. The first matching
	// branch will execute.
	FilterBranch APIListBranchFilterBranchUnion `json:"filterBranch"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BranchName   respjson.Field
		Connection   respjson.Field
		FilterBranch respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r APIListBranch) RawJSON() string { return r.JSON.raw }
func (r *APIListBranch) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this APIListBranch to a APIListBranchParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// APIListBranchParam.Overrides()
func (r APIListBranch) ToParam() APIListBranchParam {
	return param.Override[APIListBranchParam](json.RawMessage(r.RawJSON()))
}

// APIListBranchFilterBranchUnion contains all possible properties and values from
// [shared.PublicOrFilterBranch], [shared.PublicAndFilterBranch],
// [shared.PublicNotAllFilterBranch], [shared.PublicNotAnyFilterBranch],
// [shared.PublicRestrictedFilterBranch], [shared.PublicUnifiedEventsFilterBranch],
// [shared.PublicPropertyAssociationFilterBranch],
// [shared.PublicAssociationFilterBranch].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type APIListBranchFilterBranchUnion struct {
	// This field is a union of [[]shared.PublicOrFilterBranchFilterBranchUnion],
	// [[]shared.PublicAndFilterBranchFilterBranchUnion],
	// [[]shared.PublicNotAllFilterBranchFilterBranchUnion],
	// [[]shared.PublicNotAnyFilterBranchFilterBranchUnion],
	// [[]shared.PublicRestrictedFilterBranchFilterBranchUnion],
	// [[]shared.PublicUnifiedEventsFilterBranchFilterBranchUnion],
	// [[]shared.PublicPropertyAssociationFilterBranchFilterBranchUnion],
	// [[]shared.PublicAssociationFilterBranchFilterBranchUnion]
	FilterBranches       APIListBranchFilterBranchUnionFilterBranches `json:"filterBranches"`
	FilterBranchOperator string                                       `json:"filterBranchOperator"`
	FilterBranchType     string                                       `json:"filterBranchType"`
	// This field is a union of [[]shared.PublicOrFilterBranchFilterUnion],
	// [[]shared.PublicAndFilterBranchFilterUnion],
	// [[]shared.PublicNotAllFilterBranchFilterUnion],
	// [[]shared.PublicNotAnyFilterBranchFilterUnion],
	// [[]shared.PublicRestrictedFilterBranchFilterUnion],
	// [[]shared.PublicUnifiedEventsFilterBranchFilterUnion],
	// [[]shared.PublicPropertyAssociationFilterBranchFilterUnion],
	// [[]shared.PublicAssociationFilterBranchFilterUnion]
	Filters APIListBranchFilterBranchUnionFilters `json:"filters"`
	// This field is from variant [shared.PublicUnifiedEventsFilterBranch].
	EventTypeID string `json:"eventTypeId"`
	Operator    string `json:"operator"`
	// This field is from variant [shared.PublicUnifiedEventsFilterBranch].
	CoalescingRefineBy shared.PublicUnifiedEventsFilterBranchCoalescingRefineByUnion `json:"coalescingRefineBy"`
	ObjectTypeID       string                                                        `json:"objectTypeId"`
	// This field is from variant [shared.PublicPropertyAssociationFilterBranch].
	PropertyWithObjectID string `json:"propertyWithObjectId"`
	// This field is from variant [shared.PublicAssociationFilterBranch].
	AssociationCategory string `json:"associationCategory"`
	// This field is from variant [shared.PublicAssociationFilterBranch].
	AssociationTypeID int64 `json:"associationTypeId"`
	JSON              struct {
		FilterBranches       respjson.Field
		FilterBranchOperator respjson.Field
		FilterBranchType     respjson.Field
		Filters              respjson.Field
		EventTypeID          respjson.Field
		Operator             respjson.Field
		CoalescingRefineBy   respjson.Field
		ObjectTypeID         respjson.Field
		PropertyWithObjectID respjson.Field
		AssociationCategory  respjson.Field
		AssociationTypeID    respjson.Field
		raw                  string
	} `json:"-"`
}

func (u APIListBranchFilterBranchUnion) AsOr() (v shared.PublicOrFilterBranch) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u APIListBranchFilterBranchUnion) AsAnd() (v shared.PublicAndFilterBranch) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u APIListBranchFilterBranchUnion) AsNotAll() (v shared.PublicNotAllFilterBranch) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u APIListBranchFilterBranchUnion) AsNotAny() (v shared.PublicNotAnyFilterBranch) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u APIListBranchFilterBranchUnion) AsRestricted() (v shared.PublicRestrictedFilterBranch) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u APIListBranchFilterBranchUnion) AsUnifiedEvents() (v shared.PublicUnifiedEventsFilterBranch) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u APIListBranchFilterBranchUnion) AsPropertyAssociation() (v shared.PublicPropertyAssociationFilterBranch) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u APIListBranchFilterBranchUnion) AsAssociation() (v shared.PublicAssociationFilterBranch) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u APIListBranchFilterBranchUnion) RawJSON() string { return u.JSON.raw }

func (r *APIListBranchFilterBranchUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// APIListBranchFilterBranchUnionFilterBranches is an implicit subunion of
// [APIListBranchFilterBranchUnion]. APIListBranchFilterBranchUnionFilterBranches
// provides convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [APIListBranchFilterBranchUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfPublicOrFilterBranchFilterBranches
// OfPublicAndFilterBranchFilterBranches OfPublicNotAllFilterBranchFilterBranches
// OfPublicNotAnyFilterBranchFilterBranches
// OfPublicRestrictedFilterBranchFilterBranches
// OfPublicUnifiedEventsFilterBranchFilterBranches
// OfPublicPropertyAssociationFilterBranchFilterBranches
// OfPublicAssociationFilterBranchFilterBranches]
type APIListBranchFilterBranchUnionFilterBranches struct {
	// This field will be present if the value is a
	// [[]shared.PublicOrFilterBranchFilterBranchUnion] instead of an object.
	OfPublicOrFilterBranchFilterBranches []shared.PublicOrFilterBranchFilterBranchUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]shared.PublicAndFilterBranchFilterBranchUnion] instead of an object.
	OfPublicAndFilterBranchFilterBranches []shared.PublicAndFilterBranchFilterBranchUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]shared.PublicNotAllFilterBranchFilterBranchUnion] instead of an object.
	OfPublicNotAllFilterBranchFilterBranches []shared.PublicNotAllFilterBranchFilterBranchUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]shared.PublicNotAnyFilterBranchFilterBranchUnion] instead of an object.
	OfPublicNotAnyFilterBranchFilterBranches []shared.PublicNotAnyFilterBranchFilterBranchUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]shared.PublicRestrictedFilterBranchFilterBranchUnion] instead of an object.
	OfPublicRestrictedFilterBranchFilterBranches []shared.PublicRestrictedFilterBranchFilterBranchUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]shared.PublicUnifiedEventsFilterBranchFilterBranchUnion] instead of an
	// object.
	OfPublicUnifiedEventsFilterBranchFilterBranches []shared.PublicUnifiedEventsFilterBranchFilterBranchUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]shared.PublicPropertyAssociationFilterBranchFilterBranchUnion] instead of an
	// object.
	OfPublicPropertyAssociationFilterBranchFilterBranches []shared.PublicPropertyAssociationFilterBranchFilterBranchUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]shared.PublicAssociationFilterBranchFilterBranchUnion] instead of an object.
	OfPublicAssociationFilterBranchFilterBranches []shared.PublicAssociationFilterBranchFilterBranchUnion `json:",inline"`
	JSON                                          struct {
		OfPublicOrFilterBranchFilterBranches                  respjson.Field
		OfPublicAndFilterBranchFilterBranches                 respjson.Field
		OfPublicNotAllFilterBranchFilterBranches              respjson.Field
		OfPublicNotAnyFilterBranchFilterBranches              respjson.Field
		OfPublicRestrictedFilterBranchFilterBranches          respjson.Field
		OfPublicUnifiedEventsFilterBranchFilterBranches       respjson.Field
		OfPublicPropertyAssociationFilterBranchFilterBranches respjson.Field
		OfPublicAssociationFilterBranchFilterBranches         respjson.Field
		raw                                                   string
	} `json:"-"`
}

func (r *APIListBranchFilterBranchUnionFilterBranches) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// APIListBranchFilterBranchUnionFilters is an implicit subunion of
// [APIListBranchFilterBranchUnion]. APIListBranchFilterBranchUnionFilters provides
// convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [APIListBranchFilterBranchUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfPublicOrFilterBranchFilters OfPublicAndFilterBranchFilters
// OfPublicNotAllFilterBranchFilters OfPublicNotAnyFilterBranchFilters
// OfPublicRestrictedFilterBranchFilters OfPublicUnifiedEventsFilterBranchFilters
// OfPublicPropertyAssociationFilterBranchFilters
// OfPublicAssociationFilterBranchFilters]
type APIListBranchFilterBranchUnionFilters struct {
	// This field will be present if the value is a
	// [[]shared.PublicOrFilterBranchFilterUnion] instead of an object.
	OfPublicOrFilterBranchFilters []shared.PublicOrFilterBranchFilterUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]shared.PublicAndFilterBranchFilterUnion] instead of an object.
	OfPublicAndFilterBranchFilters []shared.PublicAndFilterBranchFilterUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]shared.PublicNotAllFilterBranchFilterUnion] instead of an object.
	OfPublicNotAllFilterBranchFilters []shared.PublicNotAllFilterBranchFilterUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]shared.PublicNotAnyFilterBranchFilterUnion] instead of an object.
	OfPublicNotAnyFilterBranchFilters []shared.PublicNotAnyFilterBranchFilterUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]shared.PublicRestrictedFilterBranchFilterUnion] instead of an object.
	OfPublicRestrictedFilterBranchFilters []shared.PublicRestrictedFilterBranchFilterUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]shared.PublicUnifiedEventsFilterBranchFilterUnion] instead of an object.
	OfPublicUnifiedEventsFilterBranchFilters []shared.PublicUnifiedEventsFilterBranchFilterUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]shared.PublicPropertyAssociationFilterBranchFilterUnion] instead of an
	// object.
	OfPublicPropertyAssociationFilterBranchFilters []shared.PublicPropertyAssociationFilterBranchFilterUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]shared.PublicAssociationFilterBranchFilterUnion] instead of an object.
	OfPublicAssociationFilterBranchFilters []shared.PublicAssociationFilterBranchFilterUnion `json:",inline"`
	JSON                                   struct {
		OfPublicOrFilterBranchFilters                  respjson.Field
		OfPublicAndFilterBranchFilters                 respjson.Field
		OfPublicNotAllFilterBranchFilters              respjson.Field
		OfPublicNotAnyFilterBranchFilters              respjson.Field
		OfPublicRestrictedFilterBranchFilters          respjson.Field
		OfPublicUnifiedEventsFilterBranchFilters       respjson.Field
		OfPublicPropertyAssociationFilterBranchFilters respjson.Field
		OfPublicAssociationFilterBranchFilters         respjson.Field
		raw                                            string
	} `json:"-"`
}

func (r *APIListBranchFilterBranchUnionFilters) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type APIListBranchParam struct {
	// The name of this branch
	BranchName param.Opt[string]  `json:"branchName,omitzero"`
	Connection APIConnectionParam `json:"connection,omitzero"`
	// The list criteria that determine when to execute this branch. The first matching
	// branch will execute.
	FilterBranch APIListBranchFilterBranchUnionParam `json:"filterBranch,omitzero"`
	paramObj
}

func (r APIListBranchParam) MarshalJSON() (data []byte, err error) {
	type shadow APIListBranchParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *APIListBranchParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type APIListBranchFilterBranchUnionParam struct {
	OfOr                  *shared.PublicOrFilterBranchParam                  `json:",omitzero,inline"`
	OfAnd                 *shared.PublicAndFilterBranchParam                 `json:",omitzero,inline"`
	OfNotAll              *shared.PublicNotAllFilterBranchParam              `json:",omitzero,inline"`
	OfNotAny              *shared.PublicNotAnyFilterBranchParam              `json:",omitzero,inline"`
	OfRestricted          *shared.PublicRestrictedFilterBranchParam          `json:",omitzero,inline"`
	OfUnifiedEvents       *shared.PublicUnifiedEventsFilterBranchParam       `json:",omitzero,inline"`
	OfPropertyAssociation *shared.PublicPropertyAssociationFilterBranchParam `json:",omitzero,inline"`
	OfAssociation         *shared.PublicAssociationFilterBranchParam         `json:",omitzero,inline"`
	paramUnion
}

func (u APIListBranchFilterBranchUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfOr,
		u.OfAnd,
		u.OfNotAll,
		u.OfNotAny,
		u.OfRestricted,
		u.OfUnifiedEvents,
		u.OfPropertyAssociation,
		u.OfAssociation)
}
func (u *APIListBranchFilterBranchUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *APIListBranchFilterBranchUnionParam) asAny() any {
	if !param.IsOmitted(u.OfOr) {
		return u.OfOr
	} else if !param.IsOmitted(u.OfAnd) {
		return u.OfAnd
	} else if !param.IsOmitted(u.OfNotAll) {
		return u.OfNotAll
	} else if !param.IsOmitted(u.OfNotAny) {
		return u.OfNotAny
	} else if !param.IsOmitted(u.OfRestricted) {
		return u.OfRestricted
	} else if !param.IsOmitted(u.OfUnifiedEvents) {
		return u.OfUnifiedEvents
	} else if !param.IsOmitted(u.OfPropertyAssociation) {
		return u.OfPropertyAssociation
	} else if !param.IsOmitted(u.OfAssociation) {
		return u.OfAssociation
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIListBranchFilterBranchUnionParam) GetEventTypeID() *string {
	if vt := u.OfUnifiedEvents; vt != nil {
		return &vt.EventTypeID
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIListBranchFilterBranchUnionParam) GetCoalescingRefineBy() *shared.PublicUnifiedEventsFilterBranchCoalescingRefineByUnionParam {
	if vt := u.OfUnifiedEvents; vt != nil {
		return &vt.CoalescingRefineBy
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIListBranchFilterBranchUnionParam) GetPropertyWithObjectID() *string {
	if vt := u.OfPropertyAssociation; vt != nil {
		return &vt.PropertyWithObjectID
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIListBranchFilterBranchUnionParam) GetAssociationCategory() *string {
	if vt := u.OfAssociation; vt != nil {
		return &vt.AssociationCategory
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIListBranchFilterBranchUnionParam) GetAssociationTypeID() *int64 {
	if vt := u.OfAssociation; vt != nil {
		return &vt.AssociationTypeID
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIListBranchFilterBranchUnionParam) GetFilterBranchOperator() *string {
	if vt := u.OfOr; vt != nil {
		return (*string)(&vt.FilterBranchOperator)
	} else if vt := u.OfAnd; vt != nil {
		return (*string)(&vt.FilterBranchOperator)
	} else if vt := u.OfNotAll; vt != nil {
		return (*string)(&vt.FilterBranchOperator)
	} else if vt := u.OfNotAny; vt != nil {
		return (*string)(&vt.FilterBranchOperator)
	} else if vt := u.OfRestricted; vt != nil {
		return (*string)(&vt.FilterBranchOperator)
	} else if vt := u.OfUnifiedEvents; vt != nil {
		return (*string)(&vt.FilterBranchOperator)
	} else if vt := u.OfPropertyAssociation; vt != nil {
		return (*string)(&vt.FilterBranchOperator)
	} else if vt := u.OfAssociation; vt != nil {
		return (*string)(&vt.FilterBranchOperator)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIListBranchFilterBranchUnionParam) GetFilterBranchType() *string {
	if vt := u.OfOr; vt != nil {
		return (*string)(&vt.FilterBranchType)
	} else if vt := u.OfAnd; vt != nil {
		return (*string)(&vt.FilterBranchType)
	} else if vt := u.OfNotAll; vt != nil {
		return (*string)(&vt.FilterBranchType)
	} else if vt := u.OfNotAny; vt != nil {
		return (*string)(&vt.FilterBranchType)
	} else if vt := u.OfRestricted; vt != nil {
		return (*string)(&vt.FilterBranchType)
	} else if vt := u.OfUnifiedEvents; vt != nil {
		return (*string)(&vt.FilterBranchType)
	} else if vt := u.OfPropertyAssociation; vt != nil {
		return (*string)(&vt.FilterBranchType)
	} else if vt := u.OfAssociation; vt != nil {
		return (*string)(&vt.FilterBranchType)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIListBranchFilterBranchUnionParam) GetOperator() *string {
	if vt := u.OfUnifiedEvents; vt != nil {
		return (*string)(&vt.Operator)
	} else if vt := u.OfPropertyAssociation; vt != nil {
		return (*string)(&vt.Operator)
	} else if vt := u.OfAssociation; vt != nil {
		return (*string)(&vt.Operator)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIListBranchFilterBranchUnionParam) GetObjectTypeID() *string {
	if vt := u.OfPropertyAssociation; vt != nil {
		return (*string)(&vt.ObjectTypeID)
	} else if vt := u.OfAssociation; vt != nil {
		return (*string)(&vt.ObjectTypeID)
	}
	return nil
}

// Returns a subunion which exports methods to access subproperties
//
// Or use AsAny() to get the underlying value
func (u APIListBranchFilterBranchUnionParam) GetFilterBranches() (res apiListBranchFilterBranchUnionParamFilterBranches) {
	if vt := u.OfOr; vt != nil {
		res.any = &vt.FilterBranches
	} else if vt := u.OfAnd; vt != nil {
		res.any = &vt.FilterBranches
	} else if vt := u.OfNotAll; vt != nil {
		res.any = &vt.FilterBranches
	} else if vt := u.OfNotAny; vt != nil {
		res.any = &vt.FilterBranches
	} else if vt := u.OfRestricted; vt != nil {
		res.any = &vt.FilterBranches
	} else if vt := u.OfUnifiedEvents; vt != nil {
		res.any = &vt.FilterBranches
	} else if vt := u.OfPropertyAssociation; vt != nil {
		res.any = &vt.FilterBranches
	} else if vt := u.OfAssociation; vt != nil {
		res.any = &vt.FilterBranches
	}
	return
}

// Can have the runtime types
// [_[]shared.PublicOrFilterBranchFilterBranchUnionParam],
// [_[]shared.PublicAndFilterBranchFilterBranchUnionParam],
// [_[]shared.PublicNotAllFilterBranchFilterBranchUnionParam],
// [_[]shared.PublicNotAnyFilterBranchFilterBranchUnionParam],
// [_[]shared.PublicRestrictedFilterBranchFilterBranchUnionParam],
// [_[]shared.PublicUnifiedEventsFilterBranchFilterBranchUnionParam],
// [_[]shared.PublicPropertyAssociationFilterBranchFilterBranchUnionParam],
// [_[]shared.PublicAssociationFilterBranchFilterBranchUnionParam]
type apiListBranchFilterBranchUnionParamFilterBranches struct{ any }

// Use the following switch statement to get the type of the union:
//
//	switch u.AsAny().(type) {
//	case *[]shared.PublicOrFilterBranchFilterBranchUnionParam:
//	case *[]shared.PublicAndFilterBranchFilterBranchUnionParam:
//	case *[]shared.PublicNotAllFilterBranchFilterBranchUnionParam:
//	case *[]shared.PublicNotAnyFilterBranchFilterBranchUnionParam:
//	case *[]shared.PublicRestrictedFilterBranchFilterBranchUnionParam:
//	case *[]shared.PublicUnifiedEventsFilterBranchFilterBranchUnionParam:
//	case *[]shared.PublicPropertyAssociationFilterBranchFilterBranchUnionParam:
//	case *[]shared.PublicAssociationFilterBranchFilterBranchUnionParam:
//	default:
//	    fmt.Errorf("not present")
//	}
func (u apiListBranchFilterBranchUnionParamFilterBranches) AsAny() any { return u.any }

// Returns a subunion which exports methods to access subproperties
//
// Or use AsAny() to get the underlying value
func (u APIListBranchFilterBranchUnionParam) GetFilters() (res apiListBranchFilterBranchUnionParamFilters) {
	if vt := u.OfOr; vt != nil {
		res.any = &vt.Filters
	} else if vt := u.OfAnd; vt != nil {
		res.any = &vt.Filters
	} else if vt := u.OfNotAll; vt != nil {
		res.any = &vt.Filters
	} else if vt := u.OfNotAny; vt != nil {
		res.any = &vt.Filters
	} else if vt := u.OfRestricted; vt != nil {
		res.any = &vt.Filters
	} else if vt := u.OfUnifiedEvents; vt != nil {
		res.any = &vt.Filters
	} else if vt := u.OfPropertyAssociation; vt != nil {
		res.any = &vt.Filters
	} else if vt := u.OfAssociation; vt != nil {
		res.any = &vt.Filters
	}
	return
}

// Can have the runtime types [_[]shared.PublicOrFilterBranchFilterUnionParam],
// [_[]shared.PublicAndFilterBranchFilterUnionParam],
// [_[]shared.PublicNotAllFilterBranchFilterUnionParam],
// [_[]shared.PublicNotAnyFilterBranchFilterUnionParam],
// [_[]shared.PublicRestrictedFilterBranchFilterUnionParam],
// [_[]shared.PublicUnifiedEventsFilterBranchFilterUnionParam],
// [_[]shared.PublicPropertyAssociationFilterBranchFilterUnionParam],
// [_[]shared.PublicAssociationFilterBranchFilterUnionParam]
type apiListBranchFilterBranchUnionParamFilters struct{ any }

// Use the following switch statement to get the type of the union:
//
//	switch u.AsAny().(type) {
//	case *[]shared.PublicOrFilterBranchFilterUnionParam:
//	case *[]shared.PublicAndFilterBranchFilterUnionParam:
//	case *[]shared.PublicNotAllFilterBranchFilterUnionParam:
//	case *[]shared.PublicNotAnyFilterBranchFilterUnionParam:
//	case *[]shared.PublicRestrictedFilterBranchFilterUnionParam:
//	case *[]shared.PublicUnifiedEventsFilterBranchFilterUnionParam:
//	case *[]shared.PublicPropertyAssociationFilterBranchFilterUnionParam:
//	case *[]shared.PublicAssociationFilterBranchFilterUnionParam:
//	default:
//	    fmt.Errorf("not present")
//	}
func (u apiListBranchFilterBranchUnionParamFilters) AsAny() any { return u.any }

type APIListBranchAction struct {
	// The ID for this action.
	ActionID     string          `json:"actionId,required"`
	ListBranches []APIListBranch `json:"listBranches,required"`
	// The type of action this is, can be: "STATIC_BRANCH", "LIST_BRANCH",
	// "AB_TEST_BRANCH", "CUSTOM_CODE", "WEBHOOK", or "SINGLE_CONNECTION"
	//
	// Any of "LIST_BRANCH".
	Type          APIListBranchActionType `json:"type,required"`
	DefaultBranch APIConnection           `json:"defaultBranch"`
	// The name of the default branch, the branch that gets executed if the object does
	// not match any of the `listBranch` criteria.
	DefaultBranchName string `json:"defaultBranchName"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ActionID          respjson.Field
		ListBranches      respjson.Field
		Type              respjson.Field
		DefaultBranch     respjson.Field
		DefaultBranchName respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r APIListBranchAction) RawJSON() string { return r.JSON.raw }
func (r *APIListBranchAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this APIListBranchAction to a APIListBranchActionParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// APIListBranchActionParam.Overrides()
func (r APIListBranchAction) ToParam() APIListBranchActionParam {
	return param.Override[APIListBranchActionParam](json.RawMessage(r.RawJSON()))
}

// The type of action this is, can be: "STATIC_BRANCH", "LIST_BRANCH",
// "AB_TEST_BRANCH", "CUSTOM_CODE", "WEBHOOK", or "SINGLE_CONNECTION"
type APIListBranchActionType string

const (
	APIListBranchActionTypeListBranch APIListBranchActionType = "LIST_BRANCH"
)

// The properties ActionID, ListBranches, Type are required.
type APIListBranchActionParam struct {
	// The ID for this action.
	ActionID     string               `json:"actionId,required"`
	ListBranches []APIListBranchParam `json:"listBranches,omitzero,required"`
	// The type of action this is, can be: "STATIC_BRANCH", "LIST_BRANCH",
	// "AB_TEST_BRANCH", "CUSTOM_CODE", "WEBHOOK", or "SINGLE_CONNECTION"
	//
	// Any of "LIST_BRANCH".
	Type APIListBranchActionType `json:"type,omitzero,required"`
	// The name of the default branch, the branch that gets executed if the object does
	// not match any of the `listBranch` criteria.
	DefaultBranchName param.Opt[string]  `json:"defaultBranchName,omitzero"`
	DefaultBranch     APIConnectionParam `json:"defaultBranch,omitzero"`
	paramObj
}

func (r APIListBranchActionParam) MarshalJSON() (data []byte, err error) {
	type shadow APIListBranchActionParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *APIListBranchActionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type APIManualEnrollmentCriteria struct {
	// Whether or not the same object can enroll in this workflow twice.
	ShouldReEnroll bool `json:"shouldReEnroll,required"`
	// The type of enrollment criteria this is, this can be "LIST_BASED",
	// "EVENT_BASED", or "MANUAL".
	//
	// Any of "MANUAL".
	Type APIManualEnrollmentCriteriaType `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ShouldReEnroll respjson.Field
		Type           respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r APIManualEnrollmentCriteria) RawJSON() string { return r.JSON.raw }
func (r *APIManualEnrollmentCriteria) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this APIManualEnrollmentCriteria to a
// APIManualEnrollmentCriteriaParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// APIManualEnrollmentCriteriaParam.Overrides()
func (r APIManualEnrollmentCriteria) ToParam() APIManualEnrollmentCriteriaParam {
	return param.Override[APIManualEnrollmentCriteriaParam](json.RawMessage(r.RawJSON()))
}

// The type of enrollment criteria this is, this can be "LIST_BASED",
// "EVENT_BASED", or "MANUAL".
type APIManualEnrollmentCriteriaType string

const (
	APIManualEnrollmentCriteriaTypeManual APIManualEnrollmentCriteriaType = "MANUAL"
)

// The properties ShouldReEnroll, Type are required.
type APIManualEnrollmentCriteriaParam struct {
	// Whether or not the same object can enroll in this workflow twice.
	ShouldReEnroll bool `json:"shouldReEnroll,required"`
	// The type of enrollment criteria this is, this can be "LIST_BASED",
	// "EVENT_BASED", or "MANUAL".
	//
	// Any of "MANUAL".
	Type APIManualEnrollmentCriteriaType `json:"type,omitzero,required"`
	paramObj
}

func (r APIManualEnrollmentCriteriaParam) MarshalJSON() (data []byte, err error) {
	type shadow APIManualEnrollmentCriteriaParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *APIManualEnrollmentCriteriaParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type APIMonthlyRelativeDaysEnrollmentSchedule struct {
	// Can be either "LAST_DAY_OF_MONTH" or "FIRST_MONDAY_OF_MONTH"
	//
	// Any of "LAST_DAY_OF_MONTH", "FIRST_MONDAY_OF_MONTH".
	MonthlyRelativeDays APIMonthlyRelativeDaysEnrollmentScheduleMonthlyRelativeDays `json:"monthlyRelativeDays,required"`
	TimeOfDay           APITimeOfDay                                                `json:"timeOfDay,required"`
	// The type of enrollment schedule this is, can be: "DAILY", "WEEKLY",
	// "MONTHLY_SPECIFIC_DAYS", "MONTHLY_RELATIVE_DAYS", "YEARLY"
	//
	// Any of "MONTHLY_RELATIVE_DAYS".
	Type APIMonthlyRelativeDaysEnrollmentScheduleType `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		MonthlyRelativeDays respjson.Field
		TimeOfDay           respjson.Field
		Type                respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r APIMonthlyRelativeDaysEnrollmentSchedule) RawJSON() string { return r.JSON.raw }
func (r *APIMonthlyRelativeDaysEnrollmentSchedule) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this APIMonthlyRelativeDaysEnrollmentSchedule to a
// APIMonthlyRelativeDaysEnrollmentScheduleParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// APIMonthlyRelativeDaysEnrollmentScheduleParam.Overrides()
func (r APIMonthlyRelativeDaysEnrollmentSchedule) ToParam() APIMonthlyRelativeDaysEnrollmentScheduleParam {
	return param.Override[APIMonthlyRelativeDaysEnrollmentScheduleParam](json.RawMessage(r.RawJSON()))
}

// Can be either "LAST_DAY_OF_MONTH" or "FIRST_MONDAY_OF_MONTH"
type APIMonthlyRelativeDaysEnrollmentScheduleMonthlyRelativeDays string

const (
	APIMonthlyRelativeDaysEnrollmentScheduleMonthlyRelativeDaysLastDayOfMonth     APIMonthlyRelativeDaysEnrollmentScheduleMonthlyRelativeDays = "LAST_DAY_OF_MONTH"
	APIMonthlyRelativeDaysEnrollmentScheduleMonthlyRelativeDaysFirstMondayOfMonth APIMonthlyRelativeDaysEnrollmentScheduleMonthlyRelativeDays = "FIRST_MONDAY_OF_MONTH"
)

// The type of enrollment schedule this is, can be: "DAILY", "WEEKLY",
// "MONTHLY_SPECIFIC_DAYS", "MONTHLY_RELATIVE_DAYS", "YEARLY"
type APIMonthlyRelativeDaysEnrollmentScheduleType string

const (
	APIMonthlyRelativeDaysEnrollmentScheduleTypeMonthlyRelativeDays APIMonthlyRelativeDaysEnrollmentScheduleType = "MONTHLY_RELATIVE_DAYS"
)

// The properties MonthlyRelativeDays, TimeOfDay, Type are required.
type APIMonthlyRelativeDaysEnrollmentScheduleParam struct {
	// Can be either "LAST_DAY_OF_MONTH" or "FIRST_MONDAY_OF_MONTH"
	//
	// Any of "LAST_DAY_OF_MONTH", "FIRST_MONDAY_OF_MONTH".
	MonthlyRelativeDays APIMonthlyRelativeDaysEnrollmentScheduleMonthlyRelativeDays `json:"monthlyRelativeDays,omitzero,required"`
	TimeOfDay           APITimeOfDayParam                                           `json:"timeOfDay,omitzero,required"`
	// The type of enrollment schedule this is, can be: "DAILY", "WEEKLY",
	// "MONTHLY_SPECIFIC_DAYS", "MONTHLY_RELATIVE_DAYS", "YEARLY"
	//
	// Any of "MONTHLY_RELATIVE_DAYS".
	Type APIMonthlyRelativeDaysEnrollmentScheduleType `json:"type,omitzero,required"`
	paramObj
}

func (r APIMonthlyRelativeDaysEnrollmentScheduleParam) MarshalJSON() (data []byte, err error) {
	type shadow APIMonthlyRelativeDaysEnrollmentScheduleParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *APIMonthlyRelativeDaysEnrollmentScheduleParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type APIMonthlySpecificDaysEnrollmentSchedule struct {
	// Which days of the month to run this workflow on.
	DaysOfMonth []int64      `json:"daysOfMonth,required"`
	TimeOfDay   APITimeOfDay `json:"timeOfDay,required"`
	// The type of enrollment schedule this is, can be: "DAILY", "WEEKLY",
	// "MONTHLY_SPECIFIC_DAYS", "MONTHLY_RELATIVE_DAYS", "YEARLY"
	//
	// Any of "MONTHLY_SPECIFIC_DAYS".
	Type APIMonthlySpecificDaysEnrollmentScheduleType `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DaysOfMonth respjson.Field
		TimeOfDay   respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r APIMonthlySpecificDaysEnrollmentSchedule) RawJSON() string { return r.JSON.raw }
func (r *APIMonthlySpecificDaysEnrollmentSchedule) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this APIMonthlySpecificDaysEnrollmentSchedule to a
// APIMonthlySpecificDaysEnrollmentScheduleParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// APIMonthlySpecificDaysEnrollmentScheduleParam.Overrides()
func (r APIMonthlySpecificDaysEnrollmentSchedule) ToParam() APIMonthlySpecificDaysEnrollmentScheduleParam {
	return param.Override[APIMonthlySpecificDaysEnrollmentScheduleParam](json.RawMessage(r.RawJSON()))
}

// The type of enrollment schedule this is, can be: "DAILY", "WEEKLY",
// "MONTHLY_SPECIFIC_DAYS", "MONTHLY_RELATIVE_DAYS", "YEARLY"
type APIMonthlySpecificDaysEnrollmentScheduleType string

const (
	APIMonthlySpecificDaysEnrollmentScheduleTypeMonthlySpecificDays APIMonthlySpecificDaysEnrollmentScheduleType = "MONTHLY_SPECIFIC_DAYS"
)

// The properties DaysOfMonth, TimeOfDay, Type are required.
type APIMonthlySpecificDaysEnrollmentScheduleParam struct {
	// Which days of the month to run this workflow on.
	DaysOfMonth []int64           `json:"daysOfMonth,omitzero,required"`
	TimeOfDay   APITimeOfDayParam `json:"timeOfDay,omitzero,required"`
	// The type of enrollment schedule this is, can be: "DAILY", "WEEKLY",
	// "MONTHLY_SPECIFIC_DAYS", "MONTHLY_RELATIVE_DAYS", "YEARLY"
	//
	// Any of "MONTHLY_SPECIFIC_DAYS".
	Type APIMonthlySpecificDaysEnrollmentScheduleType `json:"type,omitzero,required"`
	paramObj
}

func (r APIMonthlySpecificDaysEnrollmentScheduleParam) MarshalJSON() (data []byte, err error) {
	type shadow APIMonthlySpecificDaysEnrollmentScheduleParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *APIMonthlySpecificDaysEnrollmentScheduleParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type APIObjectPropertyValue struct {
	// The property name to pull data from.
	PropertyName string `json:"propertyName,required"`
	// This is the type of input value. This can be one of: "FIELD_DATA",
	// "OBJECT_PROPERTY", "STATIC_VALUE", "RELATIVE_DATETIME", "TIMESTAMP",
	// "INCREMENT", "FETCHED_OBJECT_PROPERTY", "APPEND_OBJECT_PROPERTY",
	// "STATIC_APPEND_VALUE", "ENROLLMENT_EVENT_PROPERTY"
	//
	// Any of "OBJECT_PROPERTY".
	Type APIObjectPropertyValueType `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PropertyName respjson.Field
		Type         respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r APIObjectPropertyValue) RawJSON() string { return r.JSON.raw }
func (r *APIObjectPropertyValue) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this APIObjectPropertyValue to a APIObjectPropertyValueParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// APIObjectPropertyValueParam.Overrides()
func (r APIObjectPropertyValue) ToParam() APIObjectPropertyValueParam {
	return param.Override[APIObjectPropertyValueParam](json.RawMessage(r.RawJSON()))
}

// This is the type of input value. This can be one of: "FIELD_DATA",
// "OBJECT_PROPERTY", "STATIC_VALUE", "RELATIVE_DATETIME", "TIMESTAMP",
// "INCREMENT", "FETCHED_OBJECT_PROPERTY", "APPEND_OBJECT_PROPERTY",
// "STATIC_APPEND_VALUE", "ENROLLMENT_EVENT_PROPERTY"
type APIObjectPropertyValueType string

const (
	APIObjectPropertyValueTypeObjectProperty APIObjectPropertyValueType = "OBJECT_PROPERTY"
)

// The properties PropertyName, Type are required.
type APIObjectPropertyValueParam struct {
	// The property name to pull data from.
	PropertyName string `json:"propertyName,required"`
	// This is the type of input value. This can be one of: "FIELD_DATA",
	// "OBJECT_PROPERTY", "STATIC_VALUE", "RELATIVE_DATETIME", "TIMESTAMP",
	// "INCREMENT", "FETCHED_OBJECT_PROPERTY", "APPEND_OBJECT_PROPERTY",
	// "STATIC_APPEND_VALUE", "ENROLLMENT_EVENT_PROPERTY"
	//
	// Any of "OBJECT_PROPERTY".
	Type APIObjectPropertyValueType `json:"type,omitzero,required"`
	paramObj
}

func (r APIObjectPropertyValueParam) MarshalJSON() (data []byte, err error) {
	type shadow APIObjectPropertyValueParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *APIObjectPropertyValueParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type APIPlatformFlow struct {
	ID           string                       `json:"id,required"`
	Actions      []APIPlatformFlowActionUnion `json:"actions,required"`
	BlockedDates []APIBlockedDate             `json:"blockedDates,required"`
	CreatedAt    time.Time                    `json:"createdAt,required" format:"date-time"`
	// Any of "PENDING", "COMPLETE".
	CRMObjectCreationStatus APIPlatformFlowCRMObjectCreationStatus `json:"crmObjectCreationStatus,required"`
	CustomProperties        map[string]string                      `json:"customProperties,required"`
	DataSources             []APIPlatformFlowDataSourceUnion       `json:"dataSources,required"`
	// Any of "WORKFLOW", "ACTION_SET", "UNKNOWN".
	FlowType              APIPlatformFlowFlowType `json:"flowType,required"`
	IsEnabled             bool                    `json:"isEnabled,required"`
	NextAvailableActionID string                  `json:"nextAvailableActionId,required"`
	ObjectTypeID          string                  `json:"objectTypeId,required"`
	RevisionID            string                  `json:"revisionId,required"`
	TimeWindows           []APITimeWindow         `json:"timeWindows,required"`
	// Any of "CONTACT_FLOW", "PLATFORM_FLOW".
	Type                    APIPlatformFlowType                         `json:"type,required"`
	UpdatedAt               time.Time                                   `json:"updatedAt,required" format:"date-time"`
	Description             string                                      `json:"description"`
	EnrollmentCriteria      APIPlatformFlowEnrollmentCriteriaUnion      `json:"enrollmentCriteria"`
	EnrollmentSchedule      APIPlatformFlowEnrollmentScheduleUnion      `json:"enrollmentSchedule"`
	Name                    string                                      `json:"name"`
	StartActionID           string                                      `json:"startActionId"`
	SuppressionFilterBranch APIPlatformFlowSuppressionFilterBranchUnion `json:"suppressionFilterBranch"`
	Uuid                    string                                      `json:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                      respjson.Field
		Actions                 respjson.Field
		BlockedDates            respjson.Field
		CreatedAt               respjson.Field
		CRMObjectCreationStatus respjson.Field
		CustomProperties        respjson.Field
		DataSources             respjson.Field
		FlowType                respjson.Field
		IsEnabled               respjson.Field
		NextAvailableActionID   respjson.Field
		ObjectTypeID            respjson.Field
		RevisionID              respjson.Field
		TimeWindows             respjson.Field
		Type                    respjson.Field
		UpdatedAt               respjson.Field
		Description             respjson.Field
		EnrollmentCriteria      respjson.Field
		EnrollmentSchedule      respjson.Field
		Name                    respjson.Field
		StartActionID           respjson.Field
		SuppressionFilterBranch respjson.Field
		Uuid                    respjson.Field
		ExtraFields             map[string]respjson.Field
		raw                     string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r APIPlatformFlow) RawJSON() string { return r.JSON.raw }
func (r *APIPlatformFlow) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// APIPlatformFlowActionUnion contains all possible properties and values from
// [APIStaticBranchAction], [APIListBranchAction], [APIAbTestBranchAction],
// [APICustomCodeAction], [APIWebhookAction], [APISingleConnectionAction].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type APIPlatformFlowActionUnion struct {
	ActionID string `json:"actionId"`
	// This field is from variant [APIStaticBranchAction].
	InputValue APIStaticBranchActionInputValueUnion `json:"inputValue"`
	// This field is from variant [APIStaticBranchAction].
	StaticBranches []APIStaticBranch `json:"staticBranches"`
	Type           string            `json:"type"`
	// This field is from variant [APIStaticBranchAction].
	DefaultBranch     APIConnection `json:"defaultBranch"`
	DefaultBranchName string        `json:"defaultBranchName"`
	// This field is from variant [APIListBranchAction].
	ListBranches []APIListBranch `json:"listBranches"`
	// This field is from variant [APIAbTestBranchAction].
	TestBranches []APIConnection `json:"testBranches"`
	// This field is from variant [APICustomCodeAction].
	InputFields []APIInputVariable `json:"inputFields"`
	// This field is from variant [APICustomCodeAction].
	OutputFields []APIEnumerationOutputField `json:"outputFields"`
	// This field is from variant [APICustomCodeAction].
	Runtime string `json:"runtime"`
	// This field is from variant [APICustomCodeAction].
	SecretNames []string `json:"secretNames"`
	// This field is from variant [APICustomCodeAction].
	SourceCode string `json:"sourceCode"`
	// This field is from variant [APICustomCodeAction].
	Connection APIConnection `json:"connection"`
	// This field is from variant [APIWebhookAction].
	Method APIWebhookActionMethod `json:"method"`
	// This field is from variant [APIWebhookAction].
	QueryParams []APIInputVariable `json:"queryParams"`
	// This field is from variant [APIWebhookAction].
	WebhookURL string `json:"webhookUrl"`
	// This field is from variant [APIWebhookAction].
	AuthSettings APIWebhookActionAuthSettingsUnion `json:"authSettings"`
	// This field is from variant [APISingleConnectionAction].
	ActionTypeID string `json:"actionTypeId"`
	// This field is from variant [APISingleConnectionAction].
	ActionTypeVersion int64 `json:"actionTypeVersion"`
	// This field is from variant [APISingleConnectionAction].
	Fields map[string]any `json:"fields"`
	JSON   struct {
		ActionID          respjson.Field
		InputValue        respjson.Field
		StaticBranches    respjson.Field
		Type              respjson.Field
		DefaultBranch     respjson.Field
		DefaultBranchName respjson.Field
		ListBranches      respjson.Field
		TestBranches      respjson.Field
		InputFields       respjson.Field
		OutputFields      respjson.Field
		Runtime           respjson.Field
		SecretNames       respjson.Field
		SourceCode        respjson.Field
		Connection        respjson.Field
		Method            respjson.Field
		QueryParams       respjson.Field
		WebhookURL        respjson.Field
		AuthSettings      respjson.Field
		ActionTypeID      respjson.Field
		ActionTypeVersion respjson.Field
		Fields            respjson.Field
		raw               string
	} `json:"-"`
}

func (u APIPlatformFlowActionUnion) AsStaticBranch() (v APIStaticBranchAction) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u APIPlatformFlowActionUnion) AsListBranch() (v APIListBranchAction) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u APIPlatformFlowActionUnion) AsAbTestBranch() (v APIAbTestBranchAction) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u APIPlatformFlowActionUnion) AsCustomCode() (v APICustomCodeAction) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u APIPlatformFlowActionUnion) AsWebhook() (v APIWebhookAction) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u APIPlatformFlowActionUnion) AsSingleConnection() (v APISingleConnectionAction) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u APIPlatformFlowActionUnion) RawJSON() string { return u.JSON.raw }

func (r *APIPlatformFlowActionUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type APIPlatformFlowCRMObjectCreationStatus string

const (
	APIPlatformFlowCRMObjectCreationStatusPending  APIPlatformFlowCRMObjectCreationStatus = "PENDING"
	APIPlatformFlowCRMObjectCreationStatusComplete APIPlatformFlowCRMObjectCreationStatus = "COMPLETE"
)

// APIPlatformFlowDataSourceUnion contains all possible properties and values from
// [APIAssociationDataSource], [APIAssociationTimestampDataSource],
// [APIStaticPropertyFilterDataSource],
// [APIEnrolledRecordPropertyFilterDataSource],
// [APIDatasetFieldPropertyFilterDataSource],
// [APIEnrolledArgumentPropertyFilterDataSource].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type APIPlatformFlowDataSourceUnion struct {
	AssociationCategory string `json:"associationCategory"`
	AssociationTypeID   int64  `json:"associationTypeId"`
	Name                string `json:"name"`
	ObjectTypeID        string `json:"objectTypeId"`
	Type                string `json:"type"`
	// This field is from variant [APIAssociationDataSource].
	SortBy       APISort `json:"sortBy"`
	PropertyName string  `json:"propertyName"`
	// This field is from variant [APIStaticPropertyFilterDataSource].
	StaticValue string `json:"staticValue"`
	// This field is from variant [APIEnrolledRecordPropertyFilterDataSource].
	RecordFieldName string `json:"recordFieldName"`
	// This field is from variant [APIDatasetFieldPropertyFilterDataSource].
	DatasetFieldName string `json:"datasetFieldName"`
	// This field is from variant [APIEnrolledArgumentPropertyFilterDataSource].
	ArgumentName string `json:"argumentName"`
	JSON         struct {
		AssociationCategory respjson.Field
		AssociationTypeID   respjson.Field
		Name                respjson.Field
		ObjectTypeID        respjson.Field
		Type                respjson.Field
		SortBy              respjson.Field
		PropertyName        respjson.Field
		StaticValue         respjson.Field
		RecordFieldName     respjson.Field
		DatasetFieldName    respjson.Field
		ArgumentName        respjson.Field
		raw                 string
	} `json:"-"`
}

func (u APIPlatformFlowDataSourceUnion) AsAssociation() (v APIAssociationDataSource) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u APIPlatformFlowDataSourceUnion) AsAssociationTimestamp() (v APIAssociationTimestampDataSource) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u APIPlatformFlowDataSourceUnion) AsStaticPropertyFilter() (v APIStaticPropertyFilterDataSource) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u APIPlatformFlowDataSourceUnion) AsEnrolledRecordPropertyFilter() (v APIEnrolledRecordPropertyFilterDataSource) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u APIPlatformFlowDataSourceUnion) AsDatasetFieldPropertyFilter() (v APIDatasetFieldPropertyFilterDataSource) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u APIPlatformFlowDataSourceUnion) AsEnrolledArgumentPropertyFilter() (v APIEnrolledArgumentPropertyFilterDataSource) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u APIPlatformFlowDataSourceUnion) RawJSON() string { return u.JSON.raw }

func (r *APIPlatformFlowDataSourceUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type APIPlatformFlowFlowType string

const (
	APIPlatformFlowFlowTypeWorkflow  APIPlatformFlowFlowType = "WORKFLOW"
	APIPlatformFlowFlowTypeActionSet APIPlatformFlowFlowType = "ACTION_SET"
	APIPlatformFlowFlowTypeUnknown   APIPlatformFlowFlowType = "UNKNOWN"
)

type APIPlatformFlowType string

const (
	APIPlatformFlowTypeContactFlow  APIPlatformFlowType = "CONTACT_FLOW"
	APIPlatformFlowTypePlatformFlow APIPlatformFlowType = "PLATFORM_FLOW"
)

// APIPlatformFlowEnrollmentCriteriaUnion contains all possible properties and
// values from [APIListBasedEnrollmentCriteria], [APIEventBasedEnrollmentCriteria],
// [APIManualEnrollmentCriteria].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type APIPlatformFlowEnrollmentCriteriaUnion struct {
	// This field is from variant [APIListBasedEnrollmentCriteria].
	ListFilterBranch APIListBasedEnrollmentCriteriaListFilterBranchUnion `json:"listFilterBranch"`
	// This field is from variant [APIListBasedEnrollmentCriteria].
	ReEnrollmentTriggersFilterBranches []APIListBasedEnrollmentCriteriaReEnrollmentTriggersFilterBranchUnion `json:"reEnrollmentTriggersFilterBranches"`
	ShouldReEnroll                     bool                                                                  `json:"shouldReEnroll"`
	Type                               string                                                                `json:"type"`
	// This field is from variant [APIListBasedEnrollmentCriteria].
	UnEnrollObjectsNotMeetingCriteria bool `json:"unEnrollObjectsNotMeetingCriteria"`
	// This field is from variant [APIEventBasedEnrollmentCriteria].
	EventFilterBranches []shared.PublicUnifiedEventsFilterBranch `json:"eventFilterBranches"`
	// This field is from variant [APIEventBasedEnrollmentCriteria].
	ListMembershipFilterBranches []APIEventBasedEnrollmentCriteriaListMembershipFilterBranchUnion `json:"listMembershipFilterBranches"`
	// This field is from variant [APIEventBasedEnrollmentCriteria].
	RefinementCriteria APIEventBasedEnrollmentCriteriaRefinementCriteriaUnion `json:"refinementCriteria"`
	JSON               struct {
		ListFilterBranch                   respjson.Field
		ReEnrollmentTriggersFilterBranches respjson.Field
		ShouldReEnroll                     respjson.Field
		Type                               respjson.Field
		UnEnrollObjectsNotMeetingCriteria  respjson.Field
		EventFilterBranches                respjson.Field
		ListMembershipFilterBranches       respjson.Field
		RefinementCriteria                 respjson.Field
		raw                                string
	} `json:"-"`
}

func (u APIPlatformFlowEnrollmentCriteriaUnion) AsListBased() (v APIListBasedEnrollmentCriteria) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u APIPlatformFlowEnrollmentCriteriaUnion) AsEventBased() (v APIEventBasedEnrollmentCriteria) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u APIPlatformFlowEnrollmentCriteriaUnion) AsManual() (v APIManualEnrollmentCriteria) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u APIPlatformFlowEnrollmentCriteriaUnion) RawJSON() string { return u.JSON.raw }

func (r *APIPlatformFlowEnrollmentCriteriaUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// APIPlatformFlowEnrollmentScheduleUnion contains all possible properties and
// values from [APIDailyEnrollmentSchedule], [APIWeeklyEnrollmentSchedule],
// [APIMonthlySpecificDaysEnrollmentSchedule],
// [APIMonthlyRelativeDaysEnrollmentSchedule], [APIYearlyEnrollmentSchedule],
// [APIPropertyBasedEnrollmentSchedule].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type APIPlatformFlowEnrollmentScheduleUnion struct {
	// This field is from variant [APIDailyEnrollmentSchedule].
	TimeOfDay APITimeOfDay `json:"timeOfDay"`
	Type      string       `json:"type"`
	// This field is from variant [APIWeeklyEnrollmentSchedule].
	DaysOfWeek []string `json:"daysOfWeek"`
	// This field is from variant [APIMonthlySpecificDaysEnrollmentSchedule].
	DaysOfMonth []int64 `json:"daysOfMonth"`
	// This field is from variant [APIMonthlyRelativeDaysEnrollmentSchedule].
	MonthlyRelativeDays APIMonthlyRelativeDaysEnrollmentScheduleMonthlyRelativeDays `json:"monthlyRelativeDays"`
	// This field is from variant [APIYearlyEnrollmentSchedule].
	DayOfMonth int64 `json:"dayOfMonth"`
	// This field is from variant [APIYearlyEnrollmentSchedule].
	Month APIYearlyEnrollmentScheduleMonth `json:"month"`
	// This field is from variant [APIPropertyBasedEnrollmentSchedule].
	DateProperty string `json:"dateProperty"`
	// This field is from variant [APIPropertyBasedEnrollmentSchedule].
	DaysDelta int64 `json:"daysDelta"`
	// This field is from variant [APIPropertyBasedEnrollmentSchedule].
	Yearly bool `json:"yearly"`
	JSON   struct {
		TimeOfDay           respjson.Field
		Type                respjson.Field
		DaysOfWeek          respjson.Field
		DaysOfMonth         respjson.Field
		MonthlyRelativeDays respjson.Field
		DayOfMonth          respjson.Field
		Month               respjson.Field
		DateProperty        respjson.Field
		DaysDelta           respjson.Field
		Yearly              respjson.Field
		raw                 string
	} `json:"-"`
}

func (u APIPlatformFlowEnrollmentScheduleUnion) AsDaily() (v APIDailyEnrollmentSchedule) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u APIPlatformFlowEnrollmentScheduleUnion) AsWeekly() (v APIWeeklyEnrollmentSchedule) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u APIPlatformFlowEnrollmentScheduleUnion) AsMonthlySpecificDays() (v APIMonthlySpecificDaysEnrollmentSchedule) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u APIPlatformFlowEnrollmentScheduleUnion) AsMonthlyRelativeDays() (v APIMonthlyRelativeDaysEnrollmentSchedule) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u APIPlatformFlowEnrollmentScheduleUnion) AsYearly() (v APIYearlyEnrollmentSchedule) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u APIPlatformFlowEnrollmentScheduleUnion) AsPropertyBased() (v APIPropertyBasedEnrollmentSchedule) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u APIPlatformFlowEnrollmentScheduleUnion) RawJSON() string { return u.JSON.raw }

func (r *APIPlatformFlowEnrollmentScheduleUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// APIPlatformFlowSuppressionFilterBranchUnion contains all possible properties and
// values from [shared.PublicOrFilterBranch], [shared.PublicAndFilterBranch],
// [shared.PublicNotAllFilterBranch], [shared.PublicNotAnyFilterBranch],
// [shared.PublicRestrictedFilterBranch], [shared.PublicUnifiedEventsFilterBranch],
// [shared.PublicPropertyAssociationFilterBranch],
// [shared.PublicAssociationFilterBranch].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type APIPlatformFlowSuppressionFilterBranchUnion struct {
	// This field is a union of [[]shared.PublicOrFilterBranchFilterBranchUnion],
	// [[]shared.PublicAndFilterBranchFilterBranchUnion],
	// [[]shared.PublicNotAllFilterBranchFilterBranchUnion],
	// [[]shared.PublicNotAnyFilterBranchFilterBranchUnion],
	// [[]shared.PublicRestrictedFilterBranchFilterBranchUnion],
	// [[]shared.PublicUnifiedEventsFilterBranchFilterBranchUnion],
	// [[]shared.PublicPropertyAssociationFilterBranchFilterBranchUnion],
	// [[]shared.PublicAssociationFilterBranchFilterBranchUnion]
	FilterBranches       APIPlatformFlowSuppressionFilterBranchUnionFilterBranches `json:"filterBranches"`
	FilterBranchOperator string                                                    `json:"filterBranchOperator"`
	FilterBranchType     string                                                    `json:"filterBranchType"`
	// This field is a union of [[]shared.PublicOrFilterBranchFilterUnion],
	// [[]shared.PublicAndFilterBranchFilterUnion],
	// [[]shared.PublicNotAllFilterBranchFilterUnion],
	// [[]shared.PublicNotAnyFilterBranchFilterUnion],
	// [[]shared.PublicRestrictedFilterBranchFilterUnion],
	// [[]shared.PublicUnifiedEventsFilterBranchFilterUnion],
	// [[]shared.PublicPropertyAssociationFilterBranchFilterUnion],
	// [[]shared.PublicAssociationFilterBranchFilterUnion]
	Filters APIPlatformFlowSuppressionFilterBranchUnionFilters `json:"filters"`
	// This field is from variant [shared.PublicUnifiedEventsFilterBranch].
	EventTypeID string `json:"eventTypeId"`
	Operator    string `json:"operator"`
	// This field is from variant [shared.PublicUnifiedEventsFilterBranch].
	CoalescingRefineBy shared.PublicUnifiedEventsFilterBranchCoalescingRefineByUnion `json:"coalescingRefineBy"`
	ObjectTypeID       string                                                        `json:"objectTypeId"`
	// This field is from variant [shared.PublicPropertyAssociationFilterBranch].
	PropertyWithObjectID string `json:"propertyWithObjectId"`
	// This field is from variant [shared.PublicAssociationFilterBranch].
	AssociationCategory string `json:"associationCategory"`
	// This field is from variant [shared.PublicAssociationFilterBranch].
	AssociationTypeID int64 `json:"associationTypeId"`
	JSON              struct {
		FilterBranches       respjson.Field
		FilterBranchOperator respjson.Field
		FilterBranchType     respjson.Field
		Filters              respjson.Field
		EventTypeID          respjson.Field
		Operator             respjson.Field
		CoalescingRefineBy   respjson.Field
		ObjectTypeID         respjson.Field
		PropertyWithObjectID respjson.Field
		AssociationCategory  respjson.Field
		AssociationTypeID    respjson.Field
		raw                  string
	} `json:"-"`
}

func (u APIPlatformFlowSuppressionFilterBranchUnion) AsOr() (v shared.PublicOrFilterBranch) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u APIPlatformFlowSuppressionFilterBranchUnion) AsAnd() (v shared.PublicAndFilterBranch) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u APIPlatformFlowSuppressionFilterBranchUnion) AsNotAll() (v shared.PublicNotAllFilterBranch) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u APIPlatformFlowSuppressionFilterBranchUnion) AsNotAny() (v shared.PublicNotAnyFilterBranch) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u APIPlatformFlowSuppressionFilterBranchUnion) AsRestricted() (v shared.PublicRestrictedFilterBranch) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u APIPlatformFlowSuppressionFilterBranchUnion) AsUnifiedEvents() (v shared.PublicUnifiedEventsFilterBranch) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u APIPlatformFlowSuppressionFilterBranchUnion) AsPropertyAssociation() (v shared.PublicPropertyAssociationFilterBranch) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u APIPlatformFlowSuppressionFilterBranchUnion) AsAssociation() (v shared.PublicAssociationFilterBranch) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u APIPlatformFlowSuppressionFilterBranchUnion) RawJSON() string { return u.JSON.raw }

func (r *APIPlatformFlowSuppressionFilterBranchUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// APIPlatformFlowSuppressionFilterBranchUnionFilterBranches is an implicit
// subunion of [APIPlatformFlowSuppressionFilterBranchUnion].
// APIPlatformFlowSuppressionFilterBranchUnionFilterBranches provides convenient
// access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [APIPlatformFlowSuppressionFilterBranchUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfPublicOrFilterBranchFilterBranches
// OfPublicAndFilterBranchFilterBranches OfPublicNotAllFilterBranchFilterBranches
// OfPublicNotAnyFilterBranchFilterBranches
// OfPublicRestrictedFilterBranchFilterBranches
// OfPublicUnifiedEventsFilterBranchFilterBranches
// OfPublicPropertyAssociationFilterBranchFilterBranches
// OfPublicAssociationFilterBranchFilterBranches]
type APIPlatformFlowSuppressionFilterBranchUnionFilterBranches struct {
	// This field will be present if the value is a
	// [[]shared.PublicOrFilterBranchFilterBranchUnion] instead of an object.
	OfPublicOrFilterBranchFilterBranches []shared.PublicOrFilterBranchFilterBranchUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]shared.PublicAndFilterBranchFilterBranchUnion] instead of an object.
	OfPublicAndFilterBranchFilterBranches []shared.PublicAndFilterBranchFilterBranchUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]shared.PublicNotAllFilterBranchFilterBranchUnion] instead of an object.
	OfPublicNotAllFilterBranchFilterBranches []shared.PublicNotAllFilterBranchFilterBranchUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]shared.PublicNotAnyFilterBranchFilterBranchUnion] instead of an object.
	OfPublicNotAnyFilterBranchFilterBranches []shared.PublicNotAnyFilterBranchFilterBranchUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]shared.PublicRestrictedFilterBranchFilterBranchUnion] instead of an object.
	OfPublicRestrictedFilterBranchFilterBranches []shared.PublicRestrictedFilterBranchFilterBranchUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]shared.PublicUnifiedEventsFilterBranchFilterBranchUnion] instead of an
	// object.
	OfPublicUnifiedEventsFilterBranchFilterBranches []shared.PublicUnifiedEventsFilterBranchFilterBranchUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]shared.PublicPropertyAssociationFilterBranchFilterBranchUnion] instead of an
	// object.
	OfPublicPropertyAssociationFilterBranchFilterBranches []shared.PublicPropertyAssociationFilterBranchFilterBranchUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]shared.PublicAssociationFilterBranchFilterBranchUnion] instead of an object.
	OfPublicAssociationFilterBranchFilterBranches []shared.PublicAssociationFilterBranchFilterBranchUnion `json:",inline"`
	JSON                                          struct {
		OfPublicOrFilterBranchFilterBranches                  respjson.Field
		OfPublicAndFilterBranchFilterBranches                 respjson.Field
		OfPublicNotAllFilterBranchFilterBranches              respjson.Field
		OfPublicNotAnyFilterBranchFilterBranches              respjson.Field
		OfPublicRestrictedFilterBranchFilterBranches          respjson.Field
		OfPublicUnifiedEventsFilterBranchFilterBranches       respjson.Field
		OfPublicPropertyAssociationFilterBranchFilterBranches respjson.Field
		OfPublicAssociationFilterBranchFilterBranches         respjson.Field
		raw                                                   string
	} `json:"-"`
}

func (r *APIPlatformFlowSuppressionFilterBranchUnionFilterBranches) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// APIPlatformFlowSuppressionFilterBranchUnionFilters is an implicit subunion of
// [APIPlatformFlowSuppressionFilterBranchUnion].
// APIPlatformFlowSuppressionFilterBranchUnionFilters provides convenient access to
// the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [APIPlatformFlowSuppressionFilterBranchUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfPublicOrFilterBranchFilters OfPublicAndFilterBranchFilters
// OfPublicNotAllFilterBranchFilters OfPublicNotAnyFilterBranchFilters
// OfPublicRestrictedFilterBranchFilters OfPublicUnifiedEventsFilterBranchFilters
// OfPublicPropertyAssociationFilterBranchFilters
// OfPublicAssociationFilterBranchFilters]
type APIPlatformFlowSuppressionFilterBranchUnionFilters struct {
	// This field will be present if the value is a
	// [[]shared.PublicOrFilterBranchFilterUnion] instead of an object.
	OfPublicOrFilterBranchFilters []shared.PublicOrFilterBranchFilterUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]shared.PublicAndFilterBranchFilterUnion] instead of an object.
	OfPublicAndFilterBranchFilters []shared.PublicAndFilterBranchFilterUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]shared.PublicNotAllFilterBranchFilterUnion] instead of an object.
	OfPublicNotAllFilterBranchFilters []shared.PublicNotAllFilterBranchFilterUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]shared.PublicNotAnyFilterBranchFilterUnion] instead of an object.
	OfPublicNotAnyFilterBranchFilters []shared.PublicNotAnyFilterBranchFilterUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]shared.PublicRestrictedFilterBranchFilterUnion] instead of an object.
	OfPublicRestrictedFilterBranchFilters []shared.PublicRestrictedFilterBranchFilterUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]shared.PublicUnifiedEventsFilterBranchFilterUnion] instead of an object.
	OfPublicUnifiedEventsFilterBranchFilters []shared.PublicUnifiedEventsFilterBranchFilterUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]shared.PublicPropertyAssociationFilterBranchFilterUnion] instead of an
	// object.
	OfPublicPropertyAssociationFilterBranchFilters []shared.PublicPropertyAssociationFilterBranchFilterUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]shared.PublicAssociationFilterBranchFilterUnion] instead of an object.
	OfPublicAssociationFilterBranchFilters []shared.PublicAssociationFilterBranchFilterUnion `json:",inline"`
	JSON                                   struct {
		OfPublicOrFilterBranchFilters                  respjson.Field
		OfPublicAndFilterBranchFilters                 respjson.Field
		OfPublicNotAllFilterBranchFilters              respjson.Field
		OfPublicNotAnyFilterBranchFilters              respjson.Field
		OfPublicRestrictedFilterBranchFilters          respjson.Field
		OfPublicUnifiedEventsFilterBranchFilters       respjson.Field
		OfPublicPropertyAssociationFilterBranchFilters respjson.Field
		OfPublicAssociationFilterBranchFilters         respjson.Field
		raw                                            string
	} `json:"-"`
}

func (r *APIPlatformFlowSuppressionFilterBranchUnionFilters) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Actions, BlockedDates, CustomProperties, DataSources, FlowType,
// IsEnabled, ObjectTypeID, TimeWindows, Type are required.
type APIPlatformFlowCreateRequestParam struct {
	Actions          []APIPlatformFlowCreateRequestActionUnionParam     `json:"actions,omitzero,required"`
	BlockedDates     []APIBlockedDateParam                              `json:"blockedDates,omitzero,required"`
	CustomProperties map[string]string                                  `json:"customProperties,omitzero,required"`
	DataSources      []APIPlatformFlowCreateRequestDataSourceUnionParam `json:"dataSources,omitzero,required"`
	// Any of "WORKFLOW", "ACTION_SET", "UNKNOWN".
	FlowType     APIPlatformFlowCreateRequestFlowType `json:"flowType,omitzero,required"`
	IsEnabled    bool                                 `json:"isEnabled,required"`
	ObjectTypeID string                               `json:"objectTypeId,required"`
	TimeWindows  []APITimeWindowParam                 `json:"timeWindows,omitzero,required"`
	// Any of "CONTACT_FLOW", "PLATFORM_FLOW".
	Type                    APIPlatformFlowCreateRequestType                              `json:"type,omitzero,required"`
	Description             param.Opt[string]                                             `json:"description,omitzero"`
	Name                    param.Opt[string]                                             `json:"name,omitzero"`
	StartActionID           param.Opt[string]                                             `json:"startActionId,omitzero"`
	Uuid                    param.Opt[string]                                             `json:"uuid,omitzero"`
	EnrollmentCriteria      APIPlatformFlowCreateRequestEnrollmentCriteriaUnionParam      `json:"enrollmentCriteria,omitzero"`
	EnrollmentSchedule      APIPlatformFlowCreateRequestEnrollmentScheduleUnionParam      `json:"enrollmentSchedule,omitzero"`
	SuppressionFilterBranch APIPlatformFlowCreateRequestSuppressionFilterBranchUnionParam `json:"suppressionFilterBranch,omitzero"`
	paramObj
}

func (r APIPlatformFlowCreateRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow APIPlatformFlowCreateRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *APIPlatformFlowCreateRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type APIPlatformFlowCreateRequestActionUnionParam struct {
	OfStaticBranch     *APIStaticBranchActionParam     `json:",omitzero,inline"`
	OfListBranch       *APIListBranchActionParam       `json:",omitzero,inline"`
	OfAbTestBranch     *APIAbTestBranchActionParam     `json:",omitzero,inline"`
	OfCustomCode       *APICustomCodeActionParam       `json:",omitzero,inline"`
	OfWebhook          *APIWebhookActionParam          `json:",omitzero,inline"`
	OfSingleConnection *APISingleConnectionActionParam `json:",omitzero,inline"`
	paramUnion
}

func (u APIPlatformFlowCreateRequestActionUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfStaticBranch,
		u.OfListBranch,
		u.OfAbTestBranch,
		u.OfCustomCode,
		u.OfWebhook,
		u.OfSingleConnection)
}
func (u *APIPlatformFlowCreateRequestActionUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *APIPlatformFlowCreateRequestActionUnionParam) asAny() any {
	if !param.IsOmitted(u.OfStaticBranch) {
		return u.OfStaticBranch
	} else if !param.IsOmitted(u.OfListBranch) {
		return u.OfListBranch
	} else if !param.IsOmitted(u.OfAbTestBranch) {
		return u.OfAbTestBranch
	} else if !param.IsOmitted(u.OfCustomCode) {
		return u.OfCustomCode
	} else if !param.IsOmitted(u.OfWebhook) {
		return u.OfWebhook
	} else if !param.IsOmitted(u.OfSingleConnection) {
		return u.OfSingleConnection
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIPlatformFlowCreateRequestActionUnionParam) GetInputValue() *APIStaticBranchActionInputValueUnionParam {
	if vt := u.OfStaticBranch; vt != nil {
		return &vt.InputValue
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIPlatformFlowCreateRequestActionUnionParam) GetStaticBranches() []APIStaticBranchParam {
	if vt := u.OfStaticBranch; vt != nil {
		return vt.StaticBranches
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIPlatformFlowCreateRequestActionUnionParam) GetListBranches() []APIListBranchParam {
	if vt := u.OfListBranch; vt != nil {
		return vt.ListBranches
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIPlatformFlowCreateRequestActionUnionParam) GetTestBranches() []APIConnectionParam {
	if vt := u.OfAbTestBranch; vt != nil {
		return vt.TestBranches
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIPlatformFlowCreateRequestActionUnionParam) GetInputFields() []APIInputVariableParam {
	if vt := u.OfCustomCode; vt != nil {
		return vt.InputFields
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIPlatformFlowCreateRequestActionUnionParam) GetOutputFields() []APIEnumerationOutputFieldParam {
	if vt := u.OfCustomCode; vt != nil {
		return vt.OutputFields
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIPlatformFlowCreateRequestActionUnionParam) GetRuntime() *string {
	if vt := u.OfCustomCode; vt != nil {
		return &vt.Runtime
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIPlatformFlowCreateRequestActionUnionParam) GetSecretNames() []string {
	if vt := u.OfCustomCode; vt != nil {
		return vt.SecretNames
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIPlatformFlowCreateRequestActionUnionParam) GetSourceCode() *string {
	if vt := u.OfCustomCode; vt != nil {
		return &vt.SourceCode
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIPlatformFlowCreateRequestActionUnionParam) GetMethod() *string {
	if vt := u.OfWebhook; vt != nil {
		return (*string)(&vt.Method)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIPlatformFlowCreateRequestActionUnionParam) GetQueryParams() []APIInputVariableParam {
	if vt := u.OfWebhook; vt != nil {
		return vt.QueryParams
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIPlatformFlowCreateRequestActionUnionParam) GetWebhookURL() *string {
	if vt := u.OfWebhook; vt != nil {
		return &vt.WebhookURL
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIPlatformFlowCreateRequestActionUnionParam) GetAuthSettings() *APIWebhookActionAuthSettingsUnionParam {
	if vt := u.OfWebhook; vt != nil {
		return &vt.AuthSettings
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIPlatformFlowCreateRequestActionUnionParam) GetActionTypeID() *string {
	if vt := u.OfSingleConnection; vt != nil {
		return &vt.ActionTypeID
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIPlatformFlowCreateRequestActionUnionParam) GetActionTypeVersion() *int64 {
	if vt := u.OfSingleConnection; vt != nil {
		return &vt.ActionTypeVersion
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIPlatformFlowCreateRequestActionUnionParam) GetFields() map[string]any {
	if vt := u.OfSingleConnection; vt != nil {
		return vt.Fields
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIPlatformFlowCreateRequestActionUnionParam) GetActionID() *string {
	if vt := u.OfStaticBranch; vt != nil {
		return (*string)(&vt.ActionID)
	} else if vt := u.OfListBranch; vt != nil {
		return (*string)(&vt.ActionID)
	} else if vt := u.OfAbTestBranch; vt != nil {
		return (*string)(&vt.ActionID)
	} else if vt := u.OfCustomCode; vt != nil {
		return (*string)(&vt.ActionID)
	} else if vt := u.OfWebhook; vt != nil {
		return (*string)(&vt.ActionID)
	} else if vt := u.OfSingleConnection; vt != nil {
		return (*string)(&vt.ActionID)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIPlatformFlowCreateRequestActionUnionParam) GetType() *string {
	if vt := u.OfStaticBranch; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfListBranch; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfAbTestBranch; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfCustomCode; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfWebhook; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfSingleConnection; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIPlatformFlowCreateRequestActionUnionParam) GetDefaultBranchName() *string {
	if vt := u.OfStaticBranch; vt != nil && vt.DefaultBranchName.Valid() {
		return &vt.DefaultBranchName.Value
	} else if vt := u.OfListBranch; vt != nil && vt.DefaultBranchName.Valid() {
		return &vt.DefaultBranchName.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's DefaultBranch property, if
// present.
func (u APIPlatformFlowCreateRequestActionUnionParam) GetDefaultBranch() *APIConnectionParam {
	if vt := u.OfStaticBranch; vt != nil {
		return &vt.DefaultBranch
	} else if vt := u.OfListBranch; vt != nil {
		return &vt.DefaultBranch
	}
	return nil
}

// Returns a pointer to the underlying variant's Connection property, if present.
func (u APIPlatformFlowCreateRequestActionUnionParam) GetConnection() *APIConnectionParam {
	if vt := u.OfCustomCode; vt != nil {
		return &vt.Connection
	} else if vt := u.OfWebhook; vt != nil {
		return &vt.Connection
	} else if vt := u.OfSingleConnection; vt != nil {
		return &vt.Connection
	}
	return nil
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type APIPlatformFlowCreateRequestDataSourceUnionParam struct {
	OfAssociation                    *APIAssociationDataSourceParam                    `json:",omitzero,inline"`
	OfAssociationTimestamp           *APIAssociationTimestampDataSourceParam           `json:",omitzero,inline"`
	OfStaticPropertyFilter           *APIStaticPropertyFilterDataSourceParam           `json:",omitzero,inline"`
	OfEnrolledRecordPropertyFilter   *APIEnrolledRecordPropertyFilterDataSourceParam   `json:",omitzero,inline"`
	OfDatasetFieldPropertyFilter     *APIDatasetFieldPropertyFilterDataSourceParam     `json:",omitzero,inline"`
	OfEnrolledArgumentPropertyFilter *APIEnrolledArgumentPropertyFilterDataSourceParam `json:",omitzero,inline"`
	paramUnion
}

func (u APIPlatformFlowCreateRequestDataSourceUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfAssociation,
		u.OfAssociationTimestamp,
		u.OfStaticPropertyFilter,
		u.OfEnrolledRecordPropertyFilter,
		u.OfDatasetFieldPropertyFilter,
		u.OfEnrolledArgumentPropertyFilter)
}
func (u *APIPlatformFlowCreateRequestDataSourceUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *APIPlatformFlowCreateRequestDataSourceUnionParam) asAny() any {
	if !param.IsOmitted(u.OfAssociation) {
		return u.OfAssociation
	} else if !param.IsOmitted(u.OfAssociationTimestamp) {
		return u.OfAssociationTimestamp
	} else if !param.IsOmitted(u.OfStaticPropertyFilter) {
		return u.OfStaticPropertyFilter
	} else if !param.IsOmitted(u.OfEnrolledRecordPropertyFilter) {
		return u.OfEnrolledRecordPropertyFilter
	} else if !param.IsOmitted(u.OfDatasetFieldPropertyFilter) {
		return u.OfDatasetFieldPropertyFilter
	} else if !param.IsOmitted(u.OfEnrolledArgumentPropertyFilter) {
		return u.OfEnrolledArgumentPropertyFilter
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIPlatformFlowCreateRequestDataSourceUnionParam) GetStaticValue() *string {
	if vt := u.OfStaticPropertyFilter; vt != nil {
		return &vt.StaticValue
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIPlatformFlowCreateRequestDataSourceUnionParam) GetRecordFieldName() *string {
	if vt := u.OfEnrolledRecordPropertyFilter; vt != nil {
		return &vt.RecordFieldName
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIPlatformFlowCreateRequestDataSourceUnionParam) GetDatasetFieldName() *string {
	if vt := u.OfDatasetFieldPropertyFilter; vt != nil {
		return &vt.DatasetFieldName
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIPlatformFlowCreateRequestDataSourceUnionParam) GetArgumentName() *string {
	if vt := u.OfEnrolledArgumentPropertyFilter; vt != nil {
		return &vt.ArgumentName
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIPlatformFlowCreateRequestDataSourceUnionParam) GetAssociationCategory() *string {
	if vt := u.OfAssociation; vt != nil {
		return (*string)(&vt.AssociationCategory)
	} else if vt := u.OfAssociationTimestamp; vt != nil {
		return (*string)(&vt.AssociationCategory)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIPlatformFlowCreateRequestDataSourceUnionParam) GetAssociationTypeID() *int64 {
	if vt := u.OfAssociation; vt != nil {
		return (*int64)(&vt.AssociationTypeID)
	} else if vt := u.OfAssociationTimestamp; vt != nil {
		return (*int64)(&vt.AssociationTypeID)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIPlatformFlowCreateRequestDataSourceUnionParam) GetName() *string {
	if vt := u.OfAssociation; vt != nil {
		return (*string)(&vt.Name)
	} else if vt := u.OfAssociationTimestamp; vt != nil {
		return (*string)(&vt.Name)
	} else if vt := u.OfStaticPropertyFilter; vt != nil {
		return (*string)(&vt.Name)
	} else if vt := u.OfEnrolledRecordPropertyFilter; vt != nil {
		return (*string)(&vt.Name)
	} else if vt := u.OfDatasetFieldPropertyFilter; vt != nil {
		return (*string)(&vt.Name)
	} else if vt := u.OfEnrolledArgumentPropertyFilter; vt != nil {
		return (*string)(&vt.Name)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIPlatformFlowCreateRequestDataSourceUnionParam) GetObjectTypeID() *string {
	if vt := u.OfAssociation; vt != nil {
		return (*string)(&vt.ObjectTypeID)
	} else if vt := u.OfAssociationTimestamp; vt != nil {
		return (*string)(&vt.ObjectTypeID)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIPlatformFlowCreateRequestDataSourceUnionParam) GetType() *string {
	if vt := u.OfAssociation; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfAssociationTimestamp; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfStaticPropertyFilter; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfEnrolledRecordPropertyFilter; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfDatasetFieldPropertyFilter; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfEnrolledArgumentPropertyFilter; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIPlatformFlowCreateRequestDataSourceUnionParam) GetPropertyName() *string {
	if vt := u.OfStaticPropertyFilter; vt != nil {
		return (*string)(&vt.PropertyName)
	} else if vt := u.OfEnrolledRecordPropertyFilter; vt != nil {
		return (*string)(&vt.PropertyName)
	} else if vt := u.OfDatasetFieldPropertyFilter; vt != nil {
		return (*string)(&vt.PropertyName)
	} else if vt := u.OfEnrolledArgumentPropertyFilter; vt != nil {
		return (*string)(&vt.PropertyName)
	}
	return nil
}

// Returns a pointer to the underlying variant's SortBy property, if present.
func (u APIPlatformFlowCreateRequestDataSourceUnionParam) GetSortBy() *APISortParam {
	if vt := u.OfAssociation; vt != nil {
		return &vt.SortBy
	} else if vt := u.OfStaticPropertyFilter; vt != nil {
		return &vt.SortBy
	} else if vt := u.OfEnrolledRecordPropertyFilter; vt != nil {
		return &vt.SortBy
	} else if vt := u.OfDatasetFieldPropertyFilter; vt != nil {
		return &vt.SortBy
	} else if vt := u.OfEnrolledArgumentPropertyFilter; vt != nil {
		return &vt.SortBy
	}
	return nil
}

type APIPlatformFlowCreateRequestFlowType string

const (
	APIPlatformFlowCreateRequestFlowTypeWorkflow  APIPlatformFlowCreateRequestFlowType = "WORKFLOW"
	APIPlatformFlowCreateRequestFlowTypeActionSet APIPlatformFlowCreateRequestFlowType = "ACTION_SET"
	APIPlatformFlowCreateRequestFlowTypeUnknown   APIPlatformFlowCreateRequestFlowType = "UNKNOWN"
)

type APIPlatformFlowCreateRequestType string

const (
	APIPlatformFlowCreateRequestTypeContactFlow  APIPlatformFlowCreateRequestType = "CONTACT_FLOW"
	APIPlatformFlowCreateRequestTypePlatformFlow APIPlatformFlowCreateRequestType = "PLATFORM_FLOW"
)

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type APIPlatformFlowCreateRequestEnrollmentCriteriaUnionParam struct {
	OfListBased  *APIListBasedEnrollmentCriteriaParam  `json:",omitzero,inline"`
	OfEventBased *APIEventBasedEnrollmentCriteriaParam `json:",omitzero,inline"`
	OfManual     *APIManualEnrollmentCriteriaParam     `json:",omitzero,inline"`
	paramUnion
}

func (u APIPlatformFlowCreateRequestEnrollmentCriteriaUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfListBased, u.OfEventBased, u.OfManual)
}
func (u *APIPlatformFlowCreateRequestEnrollmentCriteriaUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *APIPlatformFlowCreateRequestEnrollmentCriteriaUnionParam) asAny() any {
	if !param.IsOmitted(u.OfListBased) {
		return u.OfListBased
	} else if !param.IsOmitted(u.OfEventBased) {
		return u.OfEventBased
	} else if !param.IsOmitted(u.OfManual) {
		return u.OfManual
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIPlatformFlowCreateRequestEnrollmentCriteriaUnionParam) GetListFilterBranch() *APIListBasedEnrollmentCriteriaListFilterBranchUnionParam {
	if vt := u.OfListBased; vt != nil {
		return &vt.ListFilterBranch
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIPlatformFlowCreateRequestEnrollmentCriteriaUnionParam) GetReEnrollmentTriggersFilterBranches() []APIListBasedEnrollmentCriteriaReEnrollmentTriggersFilterBranchUnionParam {
	if vt := u.OfListBased; vt != nil {
		return vt.ReEnrollmentTriggersFilterBranches
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIPlatformFlowCreateRequestEnrollmentCriteriaUnionParam) GetUnEnrollObjectsNotMeetingCriteria() *bool {
	if vt := u.OfListBased; vt != nil {
		return &vt.UnEnrollObjectsNotMeetingCriteria
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIPlatformFlowCreateRequestEnrollmentCriteriaUnionParam) GetEventFilterBranches() []shared.PublicUnifiedEventsFilterBranchParam {
	if vt := u.OfEventBased; vt != nil {
		return vt.EventFilterBranches
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIPlatformFlowCreateRequestEnrollmentCriteriaUnionParam) GetListMembershipFilterBranches() []APIEventBasedEnrollmentCriteriaListMembershipFilterBranchUnionParam {
	if vt := u.OfEventBased; vt != nil {
		return vt.ListMembershipFilterBranches
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIPlatformFlowCreateRequestEnrollmentCriteriaUnionParam) GetRefinementCriteria() *APIEventBasedEnrollmentCriteriaRefinementCriteriaUnionParam {
	if vt := u.OfEventBased; vt != nil {
		return &vt.RefinementCriteria
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIPlatformFlowCreateRequestEnrollmentCriteriaUnionParam) GetShouldReEnroll() *bool {
	if vt := u.OfListBased; vt != nil {
		return (*bool)(&vt.ShouldReEnroll)
	} else if vt := u.OfEventBased; vt != nil {
		return (*bool)(&vt.ShouldReEnroll)
	} else if vt := u.OfManual; vt != nil {
		return (*bool)(&vt.ShouldReEnroll)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIPlatformFlowCreateRequestEnrollmentCriteriaUnionParam) GetType() *string {
	if vt := u.OfListBased; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfEventBased; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfManual; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type APIPlatformFlowCreateRequestEnrollmentScheduleUnionParam struct {
	OfDaily               *APIDailyEnrollmentScheduleParam               `json:",omitzero,inline"`
	OfWeekly              *APIWeeklyEnrollmentScheduleParam              `json:",omitzero,inline"`
	OfMonthlySpecificDays *APIMonthlySpecificDaysEnrollmentScheduleParam `json:",omitzero,inline"`
	OfMonthlyRelativeDays *APIMonthlyRelativeDaysEnrollmentScheduleParam `json:",omitzero,inline"`
	OfYearly              *APIYearlyEnrollmentScheduleParam              `json:",omitzero,inline"`
	OfPropertyBased       *APIPropertyBasedEnrollmentScheduleParam       `json:",omitzero,inline"`
	paramUnion
}

func (u APIPlatformFlowCreateRequestEnrollmentScheduleUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfDaily,
		u.OfWeekly,
		u.OfMonthlySpecificDays,
		u.OfMonthlyRelativeDays,
		u.OfYearly,
		u.OfPropertyBased)
}
func (u *APIPlatformFlowCreateRequestEnrollmentScheduleUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *APIPlatformFlowCreateRequestEnrollmentScheduleUnionParam) asAny() any {
	if !param.IsOmitted(u.OfDaily) {
		return u.OfDaily
	} else if !param.IsOmitted(u.OfWeekly) {
		return u.OfWeekly
	} else if !param.IsOmitted(u.OfMonthlySpecificDays) {
		return u.OfMonthlySpecificDays
	} else if !param.IsOmitted(u.OfMonthlyRelativeDays) {
		return u.OfMonthlyRelativeDays
	} else if !param.IsOmitted(u.OfYearly) {
		return u.OfYearly
	} else if !param.IsOmitted(u.OfPropertyBased) {
		return u.OfPropertyBased
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIPlatformFlowCreateRequestEnrollmentScheduleUnionParam) GetDaysOfWeek() []string {
	if vt := u.OfWeekly; vt != nil {
		return vt.DaysOfWeek
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIPlatformFlowCreateRequestEnrollmentScheduleUnionParam) GetDaysOfMonth() []int64 {
	if vt := u.OfMonthlySpecificDays; vt != nil {
		return vt.DaysOfMonth
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIPlatformFlowCreateRequestEnrollmentScheduleUnionParam) GetMonthlyRelativeDays() *string {
	if vt := u.OfMonthlyRelativeDays; vt != nil {
		return (*string)(&vt.MonthlyRelativeDays)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIPlatformFlowCreateRequestEnrollmentScheduleUnionParam) GetDayOfMonth() *int64 {
	if vt := u.OfYearly; vt != nil {
		return &vt.DayOfMonth
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIPlatformFlowCreateRequestEnrollmentScheduleUnionParam) GetMonth() *string {
	if vt := u.OfYearly; vt != nil {
		return (*string)(&vt.Month)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIPlatformFlowCreateRequestEnrollmentScheduleUnionParam) GetDateProperty() *string {
	if vt := u.OfPropertyBased; vt != nil {
		return &vt.DateProperty
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIPlatformFlowCreateRequestEnrollmentScheduleUnionParam) GetDaysDelta() *int64 {
	if vt := u.OfPropertyBased; vt != nil {
		return &vt.DaysDelta
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIPlatformFlowCreateRequestEnrollmentScheduleUnionParam) GetYearly() *bool {
	if vt := u.OfPropertyBased; vt != nil {
		return &vt.Yearly
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIPlatformFlowCreateRequestEnrollmentScheduleUnionParam) GetType() *string {
	if vt := u.OfDaily; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfWeekly; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfMonthlySpecificDays; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfMonthlyRelativeDays; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfYearly; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfPropertyBased; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

// Returns a pointer to the underlying variant's TimeOfDay property, if present.
func (u APIPlatformFlowCreateRequestEnrollmentScheduleUnionParam) GetTimeOfDay() *APITimeOfDayParam {
	if vt := u.OfDaily; vt != nil {
		return &vt.TimeOfDay
	} else if vt := u.OfWeekly; vt != nil {
		return &vt.TimeOfDay
	} else if vt := u.OfMonthlySpecificDays; vt != nil {
		return &vt.TimeOfDay
	} else if vt := u.OfMonthlyRelativeDays; vt != nil {
		return &vt.TimeOfDay
	} else if vt := u.OfYearly; vt != nil {
		return &vt.TimeOfDay
	} else if vt := u.OfPropertyBased; vt != nil {
		return &vt.TimeOfDay
	}
	return nil
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type APIPlatformFlowCreateRequestSuppressionFilterBranchUnionParam struct {
	OfOr                  *shared.PublicOrFilterBranchParam                  `json:",omitzero,inline"`
	OfAnd                 *shared.PublicAndFilterBranchParam                 `json:",omitzero,inline"`
	OfNotAll              *shared.PublicNotAllFilterBranchParam              `json:",omitzero,inline"`
	OfNotAny              *shared.PublicNotAnyFilterBranchParam              `json:",omitzero,inline"`
	OfRestricted          *shared.PublicRestrictedFilterBranchParam          `json:",omitzero,inline"`
	OfUnifiedEvents       *shared.PublicUnifiedEventsFilterBranchParam       `json:",omitzero,inline"`
	OfPropertyAssociation *shared.PublicPropertyAssociationFilterBranchParam `json:",omitzero,inline"`
	OfAssociation         *shared.PublicAssociationFilterBranchParam         `json:",omitzero,inline"`
	paramUnion
}

func (u APIPlatformFlowCreateRequestSuppressionFilterBranchUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfOr,
		u.OfAnd,
		u.OfNotAll,
		u.OfNotAny,
		u.OfRestricted,
		u.OfUnifiedEvents,
		u.OfPropertyAssociation,
		u.OfAssociation)
}
func (u *APIPlatformFlowCreateRequestSuppressionFilterBranchUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *APIPlatformFlowCreateRequestSuppressionFilterBranchUnionParam) asAny() any {
	if !param.IsOmitted(u.OfOr) {
		return u.OfOr
	} else if !param.IsOmitted(u.OfAnd) {
		return u.OfAnd
	} else if !param.IsOmitted(u.OfNotAll) {
		return u.OfNotAll
	} else if !param.IsOmitted(u.OfNotAny) {
		return u.OfNotAny
	} else if !param.IsOmitted(u.OfRestricted) {
		return u.OfRestricted
	} else if !param.IsOmitted(u.OfUnifiedEvents) {
		return u.OfUnifiedEvents
	} else if !param.IsOmitted(u.OfPropertyAssociation) {
		return u.OfPropertyAssociation
	} else if !param.IsOmitted(u.OfAssociation) {
		return u.OfAssociation
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIPlatformFlowCreateRequestSuppressionFilterBranchUnionParam) GetEventTypeID() *string {
	if vt := u.OfUnifiedEvents; vt != nil {
		return &vt.EventTypeID
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIPlatformFlowCreateRequestSuppressionFilterBranchUnionParam) GetCoalescingRefineBy() *shared.PublicUnifiedEventsFilterBranchCoalescingRefineByUnionParam {
	if vt := u.OfUnifiedEvents; vt != nil {
		return &vt.CoalescingRefineBy
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIPlatformFlowCreateRequestSuppressionFilterBranchUnionParam) GetPropertyWithObjectID() *string {
	if vt := u.OfPropertyAssociation; vt != nil {
		return &vt.PropertyWithObjectID
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIPlatformFlowCreateRequestSuppressionFilterBranchUnionParam) GetAssociationCategory() *string {
	if vt := u.OfAssociation; vt != nil {
		return &vt.AssociationCategory
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIPlatformFlowCreateRequestSuppressionFilterBranchUnionParam) GetAssociationTypeID() *int64 {
	if vt := u.OfAssociation; vt != nil {
		return &vt.AssociationTypeID
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIPlatformFlowCreateRequestSuppressionFilterBranchUnionParam) GetFilterBranchOperator() *string {
	if vt := u.OfOr; vt != nil {
		return (*string)(&vt.FilterBranchOperator)
	} else if vt := u.OfAnd; vt != nil {
		return (*string)(&vt.FilterBranchOperator)
	} else if vt := u.OfNotAll; vt != nil {
		return (*string)(&vt.FilterBranchOperator)
	} else if vt := u.OfNotAny; vt != nil {
		return (*string)(&vt.FilterBranchOperator)
	} else if vt := u.OfRestricted; vt != nil {
		return (*string)(&vt.FilterBranchOperator)
	} else if vt := u.OfUnifiedEvents; vt != nil {
		return (*string)(&vt.FilterBranchOperator)
	} else if vt := u.OfPropertyAssociation; vt != nil {
		return (*string)(&vt.FilterBranchOperator)
	} else if vt := u.OfAssociation; vt != nil {
		return (*string)(&vt.FilterBranchOperator)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIPlatformFlowCreateRequestSuppressionFilterBranchUnionParam) GetFilterBranchType() *string {
	if vt := u.OfOr; vt != nil {
		return (*string)(&vt.FilterBranchType)
	} else if vt := u.OfAnd; vt != nil {
		return (*string)(&vt.FilterBranchType)
	} else if vt := u.OfNotAll; vt != nil {
		return (*string)(&vt.FilterBranchType)
	} else if vt := u.OfNotAny; vt != nil {
		return (*string)(&vt.FilterBranchType)
	} else if vt := u.OfRestricted; vt != nil {
		return (*string)(&vt.FilterBranchType)
	} else if vt := u.OfUnifiedEvents; vt != nil {
		return (*string)(&vt.FilterBranchType)
	} else if vt := u.OfPropertyAssociation; vt != nil {
		return (*string)(&vt.FilterBranchType)
	} else if vt := u.OfAssociation; vt != nil {
		return (*string)(&vt.FilterBranchType)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIPlatformFlowCreateRequestSuppressionFilterBranchUnionParam) GetOperator() *string {
	if vt := u.OfUnifiedEvents; vt != nil {
		return (*string)(&vt.Operator)
	} else if vt := u.OfPropertyAssociation; vt != nil {
		return (*string)(&vt.Operator)
	} else if vt := u.OfAssociation; vt != nil {
		return (*string)(&vt.Operator)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIPlatformFlowCreateRequestSuppressionFilterBranchUnionParam) GetObjectTypeID() *string {
	if vt := u.OfPropertyAssociation; vt != nil {
		return (*string)(&vt.ObjectTypeID)
	} else if vt := u.OfAssociation; vt != nil {
		return (*string)(&vt.ObjectTypeID)
	}
	return nil
}

// Returns a subunion which exports methods to access subproperties
//
// Or use AsAny() to get the underlying value
func (u APIPlatformFlowCreateRequestSuppressionFilterBranchUnionParam) GetFilterBranches() (res apiPlatformFlowCreateRequestSuppressionFilterBranchUnionParamFilterBranches) {
	if vt := u.OfOr; vt != nil {
		res.any = &vt.FilterBranches
	} else if vt := u.OfAnd; vt != nil {
		res.any = &vt.FilterBranches
	} else if vt := u.OfNotAll; vt != nil {
		res.any = &vt.FilterBranches
	} else if vt := u.OfNotAny; vt != nil {
		res.any = &vt.FilterBranches
	} else if vt := u.OfRestricted; vt != nil {
		res.any = &vt.FilterBranches
	} else if vt := u.OfUnifiedEvents; vt != nil {
		res.any = &vt.FilterBranches
	} else if vt := u.OfPropertyAssociation; vt != nil {
		res.any = &vt.FilterBranches
	} else if vt := u.OfAssociation; vt != nil {
		res.any = &vt.FilterBranches
	}
	return
}

// Can have the runtime types
// [_[]shared.PublicOrFilterBranchFilterBranchUnionParam],
// [_[]shared.PublicAndFilterBranchFilterBranchUnionParam],
// [_[]shared.PublicNotAllFilterBranchFilterBranchUnionParam],
// [_[]shared.PublicNotAnyFilterBranchFilterBranchUnionParam],
// [_[]shared.PublicRestrictedFilterBranchFilterBranchUnionParam],
// [_[]shared.PublicUnifiedEventsFilterBranchFilterBranchUnionParam],
// [_[]shared.PublicPropertyAssociationFilterBranchFilterBranchUnionParam],
// [_[]shared.PublicAssociationFilterBranchFilterBranchUnionParam]
type apiPlatformFlowCreateRequestSuppressionFilterBranchUnionParamFilterBranches struct{ any }

// Use the following switch statement to get the type of the union:
//
//	switch u.AsAny().(type) {
//	case *[]shared.PublicOrFilterBranchFilterBranchUnionParam:
//	case *[]shared.PublicAndFilterBranchFilterBranchUnionParam:
//	case *[]shared.PublicNotAllFilterBranchFilterBranchUnionParam:
//	case *[]shared.PublicNotAnyFilterBranchFilterBranchUnionParam:
//	case *[]shared.PublicRestrictedFilterBranchFilterBranchUnionParam:
//	case *[]shared.PublicUnifiedEventsFilterBranchFilterBranchUnionParam:
//	case *[]shared.PublicPropertyAssociationFilterBranchFilterBranchUnionParam:
//	case *[]shared.PublicAssociationFilterBranchFilterBranchUnionParam:
//	default:
//	    fmt.Errorf("not present")
//	}
func (u apiPlatformFlowCreateRequestSuppressionFilterBranchUnionParamFilterBranches) AsAny() any {
	return u.any
}

// Returns a subunion which exports methods to access subproperties
//
// Or use AsAny() to get the underlying value
func (u APIPlatformFlowCreateRequestSuppressionFilterBranchUnionParam) GetFilters() (res apiPlatformFlowCreateRequestSuppressionFilterBranchUnionParamFilters) {
	if vt := u.OfOr; vt != nil {
		res.any = &vt.Filters
	} else if vt := u.OfAnd; vt != nil {
		res.any = &vt.Filters
	} else if vt := u.OfNotAll; vt != nil {
		res.any = &vt.Filters
	} else if vt := u.OfNotAny; vt != nil {
		res.any = &vt.Filters
	} else if vt := u.OfRestricted; vt != nil {
		res.any = &vt.Filters
	} else if vt := u.OfUnifiedEvents; vt != nil {
		res.any = &vt.Filters
	} else if vt := u.OfPropertyAssociation; vt != nil {
		res.any = &vt.Filters
	} else if vt := u.OfAssociation; vt != nil {
		res.any = &vt.Filters
	}
	return
}

// Can have the runtime types [_[]shared.PublicOrFilterBranchFilterUnionParam],
// [_[]shared.PublicAndFilterBranchFilterUnionParam],
// [_[]shared.PublicNotAllFilterBranchFilterUnionParam],
// [_[]shared.PublicNotAnyFilterBranchFilterUnionParam],
// [_[]shared.PublicRestrictedFilterBranchFilterUnionParam],
// [_[]shared.PublicUnifiedEventsFilterBranchFilterUnionParam],
// [_[]shared.PublicPropertyAssociationFilterBranchFilterUnionParam],
// [_[]shared.PublicAssociationFilterBranchFilterUnionParam]
type apiPlatformFlowCreateRequestSuppressionFilterBranchUnionParamFilters struct{ any }

// Use the following switch statement to get the type of the union:
//
//	switch u.AsAny().(type) {
//	case *[]shared.PublicOrFilterBranchFilterUnionParam:
//	case *[]shared.PublicAndFilterBranchFilterUnionParam:
//	case *[]shared.PublicNotAllFilterBranchFilterUnionParam:
//	case *[]shared.PublicNotAnyFilterBranchFilterUnionParam:
//	case *[]shared.PublicRestrictedFilterBranchFilterUnionParam:
//	case *[]shared.PublicUnifiedEventsFilterBranchFilterUnionParam:
//	case *[]shared.PublicPropertyAssociationFilterBranchFilterUnionParam:
//	case *[]shared.PublicAssociationFilterBranchFilterUnionParam:
//	default:
//	    fmt.Errorf("not present")
//	}
func (u apiPlatformFlowCreateRequestSuppressionFilterBranchUnionParamFilters) AsAny() any {
	return u.any
}

// The properties Actions, BlockedDates, CustomProperties, IsEnabled, RevisionID,
// TimeWindows, Type are required.
type APIPlatformFlowPutRequestParam struct {
	Actions          []APIPlatformFlowPutRequestActionUnionParam `json:"actions,omitzero,required"`
	BlockedDates     []APIBlockedDateParam                       `json:"blockedDates,omitzero,required"`
	CustomProperties map[string]string                           `json:"customProperties,omitzero,required"`
	IsEnabled        bool                                        `json:"isEnabled,required"`
	RevisionID       string                                      `json:"revisionId,required"`
	TimeWindows      []APITimeWindowParam                        `json:"timeWindows,omitzero,required"`
	// Any of "CONTACT_FLOW", "PLATFORM_FLOW".
	Type                    APIPlatformFlowPutRequestType                              `json:"type,omitzero,required"`
	Description             param.Opt[string]                                          `json:"description,omitzero"`
	Name                    param.Opt[string]                                          `json:"name,omitzero"`
	StartActionID           param.Opt[string]                                          `json:"startActionId,omitzero"`
	Uuid                    param.Opt[string]                                          `json:"uuid,omitzero"`
	EnrollmentCriteria      APIPlatformFlowPutRequestEnrollmentCriteriaUnionParam      `json:"enrollmentCriteria,omitzero"`
	EnrollmentSchedule      APIPlatformFlowPutRequestEnrollmentScheduleUnionParam      `json:"enrollmentSchedule,omitzero"`
	SuppressionFilterBranch APIPlatformFlowPutRequestSuppressionFilterBranchUnionParam `json:"suppressionFilterBranch,omitzero"`
	paramObj
}

func (r APIPlatformFlowPutRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow APIPlatformFlowPutRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *APIPlatformFlowPutRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type APIPlatformFlowPutRequestActionUnionParam struct {
	OfStaticBranch     *APIStaticBranchActionParam     `json:",omitzero,inline"`
	OfListBranch       *APIListBranchActionParam       `json:",omitzero,inline"`
	OfAbTestBranch     *APIAbTestBranchActionParam     `json:",omitzero,inline"`
	OfCustomCode       *APICustomCodeActionParam       `json:",omitzero,inline"`
	OfWebhook          *APIWebhookActionParam          `json:",omitzero,inline"`
	OfSingleConnection *APISingleConnectionActionParam `json:",omitzero,inline"`
	paramUnion
}

func (u APIPlatformFlowPutRequestActionUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfStaticBranch,
		u.OfListBranch,
		u.OfAbTestBranch,
		u.OfCustomCode,
		u.OfWebhook,
		u.OfSingleConnection)
}
func (u *APIPlatformFlowPutRequestActionUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *APIPlatformFlowPutRequestActionUnionParam) asAny() any {
	if !param.IsOmitted(u.OfStaticBranch) {
		return u.OfStaticBranch
	} else if !param.IsOmitted(u.OfListBranch) {
		return u.OfListBranch
	} else if !param.IsOmitted(u.OfAbTestBranch) {
		return u.OfAbTestBranch
	} else if !param.IsOmitted(u.OfCustomCode) {
		return u.OfCustomCode
	} else if !param.IsOmitted(u.OfWebhook) {
		return u.OfWebhook
	} else if !param.IsOmitted(u.OfSingleConnection) {
		return u.OfSingleConnection
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIPlatformFlowPutRequestActionUnionParam) GetInputValue() *APIStaticBranchActionInputValueUnionParam {
	if vt := u.OfStaticBranch; vt != nil {
		return &vt.InputValue
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIPlatformFlowPutRequestActionUnionParam) GetStaticBranches() []APIStaticBranchParam {
	if vt := u.OfStaticBranch; vt != nil {
		return vt.StaticBranches
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIPlatformFlowPutRequestActionUnionParam) GetListBranches() []APIListBranchParam {
	if vt := u.OfListBranch; vt != nil {
		return vt.ListBranches
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIPlatformFlowPutRequestActionUnionParam) GetTestBranches() []APIConnectionParam {
	if vt := u.OfAbTestBranch; vt != nil {
		return vt.TestBranches
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIPlatformFlowPutRequestActionUnionParam) GetInputFields() []APIInputVariableParam {
	if vt := u.OfCustomCode; vt != nil {
		return vt.InputFields
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIPlatformFlowPutRequestActionUnionParam) GetOutputFields() []APIEnumerationOutputFieldParam {
	if vt := u.OfCustomCode; vt != nil {
		return vt.OutputFields
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIPlatformFlowPutRequestActionUnionParam) GetRuntime() *string {
	if vt := u.OfCustomCode; vt != nil {
		return &vt.Runtime
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIPlatformFlowPutRequestActionUnionParam) GetSecretNames() []string {
	if vt := u.OfCustomCode; vt != nil {
		return vt.SecretNames
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIPlatformFlowPutRequestActionUnionParam) GetSourceCode() *string {
	if vt := u.OfCustomCode; vt != nil {
		return &vt.SourceCode
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIPlatformFlowPutRequestActionUnionParam) GetMethod() *string {
	if vt := u.OfWebhook; vt != nil {
		return (*string)(&vt.Method)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIPlatformFlowPutRequestActionUnionParam) GetQueryParams() []APIInputVariableParam {
	if vt := u.OfWebhook; vt != nil {
		return vt.QueryParams
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIPlatformFlowPutRequestActionUnionParam) GetWebhookURL() *string {
	if vt := u.OfWebhook; vt != nil {
		return &vt.WebhookURL
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIPlatformFlowPutRequestActionUnionParam) GetAuthSettings() *APIWebhookActionAuthSettingsUnionParam {
	if vt := u.OfWebhook; vt != nil {
		return &vt.AuthSettings
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIPlatformFlowPutRequestActionUnionParam) GetActionTypeID() *string {
	if vt := u.OfSingleConnection; vt != nil {
		return &vt.ActionTypeID
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIPlatformFlowPutRequestActionUnionParam) GetActionTypeVersion() *int64 {
	if vt := u.OfSingleConnection; vt != nil {
		return &vt.ActionTypeVersion
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIPlatformFlowPutRequestActionUnionParam) GetFields() map[string]any {
	if vt := u.OfSingleConnection; vt != nil {
		return vt.Fields
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIPlatformFlowPutRequestActionUnionParam) GetActionID() *string {
	if vt := u.OfStaticBranch; vt != nil {
		return (*string)(&vt.ActionID)
	} else if vt := u.OfListBranch; vt != nil {
		return (*string)(&vt.ActionID)
	} else if vt := u.OfAbTestBranch; vt != nil {
		return (*string)(&vt.ActionID)
	} else if vt := u.OfCustomCode; vt != nil {
		return (*string)(&vt.ActionID)
	} else if vt := u.OfWebhook; vt != nil {
		return (*string)(&vt.ActionID)
	} else if vt := u.OfSingleConnection; vt != nil {
		return (*string)(&vt.ActionID)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIPlatformFlowPutRequestActionUnionParam) GetType() *string {
	if vt := u.OfStaticBranch; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfListBranch; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfAbTestBranch; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfCustomCode; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfWebhook; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfSingleConnection; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIPlatformFlowPutRequestActionUnionParam) GetDefaultBranchName() *string {
	if vt := u.OfStaticBranch; vt != nil && vt.DefaultBranchName.Valid() {
		return &vt.DefaultBranchName.Value
	} else if vt := u.OfListBranch; vt != nil && vt.DefaultBranchName.Valid() {
		return &vt.DefaultBranchName.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's DefaultBranch property, if
// present.
func (u APIPlatformFlowPutRequestActionUnionParam) GetDefaultBranch() *APIConnectionParam {
	if vt := u.OfStaticBranch; vt != nil {
		return &vt.DefaultBranch
	} else if vt := u.OfListBranch; vt != nil {
		return &vt.DefaultBranch
	}
	return nil
}

// Returns a pointer to the underlying variant's Connection property, if present.
func (u APIPlatformFlowPutRequestActionUnionParam) GetConnection() *APIConnectionParam {
	if vt := u.OfCustomCode; vt != nil {
		return &vt.Connection
	} else if vt := u.OfWebhook; vt != nil {
		return &vt.Connection
	} else if vt := u.OfSingleConnection; vt != nil {
		return &vt.Connection
	}
	return nil
}

type APIPlatformFlowPutRequestType string

const (
	APIPlatformFlowPutRequestTypeContactFlow  APIPlatformFlowPutRequestType = "CONTACT_FLOW"
	APIPlatformFlowPutRequestTypePlatformFlow APIPlatformFlowPutRequestType = "PLATFORM_FLOW"
)

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type APIPlatformFlowPutRequestEnrollmentCriteriaUnionParam struct {
	OfListBased  *APIListBasedEnrollmentCriteriaParam  `json:",omitzero,inline"`
	OfEventBased *APIEventBasedEnrollmentCriteriaParam `json:",omitzero,inline"`
	OfManual     *APIManualEnrollmentCriteriaParam     `json:",omitzero,inline"`
	paramUnion
}

func (u APIPlatformFlowPutRequestEnrollmentCriteriaUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfListBased, u.OfEventBased, u.OfManual)
}
func (u *APIPlatformFlowPutRequestEnrollmentCriteriaUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *APIPlatformFlowPutRequestEnrollmentCriteriaUnionParam) asAny() any {
	if !param.IsOmitted(u.OfListBased) {
		return u.OfListBased
	} else if !param.IsOmitted(u.OfEventBased) {
		return u.OfEventBased
	} else if !param.IsOmitted(u.OfManual) {
		return u.OfManual
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIPlatformFlowPutRequestEnrollmentCriteriaUnionParam) GetListFilterBranch() *APIListBasedEnrollmentCriteriaListFilterBranchUnionParam {
	if vt := u.OfListBased; vt != nil {
		return &vt.ListFilterBranch
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIPlatformFlowPutRequestEnrollmentCriteriaUnionParam) GetReEnrollmentTriggersFilterBranches() []APIListBasedEnrollmentCriteriaReEnrollmentTriggersFilterBranchUnionParam {
	if vt := u.OfListBased; vt != nil {
		return vt.ReEnrollmentTriggersFilterBranches
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIPlatformFlowPutRequestEnrollmentCriteriaUnionParam) GetUnEnrollObjectsNotMeetingCriteria() *bool {
	if vt := u.OfListBased; vt != nil {
		return &vt.UnEnrollObjectsNotMeetingCriteria
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIPlatformFlowPutRequestEnrollmentCriteriaUnionParam) GetEventFilterBranches() []shared.PublicUnifiedEventsFilterBranchParam {
	if vt := u.OfEventBased; vt != nil {
		return vt.EventFilterBranches
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIPlatformFlowPutRequestEnrollmentCriteriaUnionParam) GetListMembershipFilterBranches() []APIEventBasedEnrollmentCriteriaListMembershipFilterBranchUnionParam {
	if vt := u.OfEventBased; vt != nil {
		return vt.ListMembershipFilterBranches
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIPlatformFlowPutRequestEnrollmentCriteriaUnionParam) GetRefinementCriteria() *APIEventBasedEnrollmentCriteriaRefinementCriteriaUnionParam {
	if vt := u.OfEventBased; vt != nil {
		return &vt.RefinementCriteria
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIPlatformFlowPutRequestEnrollmentCriteriaUnionParam) GetShouldReEnroll() *bool {
	if vt := u.OfListBased; vt != nil {
		return (*bool)(&vt.ShouldReEnroll)
	} else if vt := u.OfEventBased; vt != nil {
		return (*bool)(&vt.ShouldReEnroll)
	} else if vt := u.OfManual; vt != nil {
		return (*bool)(&vt.ShouldReEnroll)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIPlatformFlowPutRequestEnrollmentCriteriaUnionParam) GetType() *string {
	if vt := u.OfListBased; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfEventBased; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfManual; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type APIPlatformFlowPutRequestEnrollmentScheduleUnionParam struct {
	OfDaily               *APIDailyEnrollmentScheduleParam               `json:",omitzero,inline"`
	OfWeekly              *APIWeeklyEnrollmentScheduleParam              `json:",omitzero,inline"`
	OfMonthlySpecificDays *APIMonthlySpecificDaysEnrollmentScheduleParam `json:",omitzero,inline"`
	OfMonthlyRelativeDays *APIMonthlyRelativeDaysEnrollmentScheduleParam `json:",omitzero,inline"`
	OfYearly              *APIYearlyEnrollmentScheduleParam              `json:",omitzero,inline"`
	OfPropertyBased       *APIPropertyBasedEnrollmentScheduleParam       `json:",omitzero,inline"`
	paramUnion
}

func (u APIPlatformFlowPutRequestEnrollmentScheduleUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfDaily,
		u.OfWeekly,
		u.OfMonthlySpecificDays,
		u.OfMonthlyRelativeDays,
		u.OfYearly,
		u.OfPropertyBased)
}
func (u *APIPlatformFlowPutRequestEnrollmentScheduleUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *APIPlatformFlowPutRequestEnrollmentScheduleUnionParam) asAny() any {
	if !param.IsOmitted(u.OfDaily) {
		return u.OfDaily
	} else if !param.IsOmitted(u.OfWeekly) {
		return u.OfWeekly
	} else if !param.IsOmitted(u.OfMonthlySpecificDays) {
		return u.OfMonthlySpecificDays
	} else if !param.IsOmitted(u.OfMonthlyRelativeDays) {
		return u.OfMonthlyRelativeDays
	} else if !param.IsOmitted(u.OfYearly) {
		return u.OfYearly
	} else if !param.IsOmitted(u.OfPropertyBased) {
		return u.OfPropertyBased
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIPlatformFlowPutRequestEnrollmentScheduleUnionParam) GetDaysOfWeek() []string {
	if vt := u.OfWeekly; vt != nil {
		return vt.DaysOfWeek
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIPlatformFlowPutRequestEnrollmentScheduleUnionParam) GetDaysOfMonth() []int64 {
	if vt := u.OfMonthlySpecificDays; vt != nil {
		return vt.DaysOfMonth
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIPlatformFlowPutRequestEnrollmentScheduleUnionParam) GetMonthlyRelativeDays() *string {
	if vt := u.OfMonthlyRelativeDays; vt != nil {
		return (*string)(&vt.MonthlyRelativeDays)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIPlatformFlowPutRequestEnrollmentScheduleUnionParam) GetDayOfMonth() *int64 {
	if vt := u.OfYearly; vt != nil {
		return &vt.DayOfMonth
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIPlatformFlowPutRequestEnrollmentScheduleUnionParam) GetMonth() *string {
	if vt := u.OfYearly; vt != nil {
		return (*string)(&vt.Month)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIPlatformFlowPutRequestEnrollmentScheduleUnionParam) GetDateProperty() *string {
	if vt := u.OfPropertyBased; vt != nil {
		return &vt.DateProperty
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIPlatformFlowPutRequestEnrollmentScheduleUnionParam) GetDaysDelta() *int64 {
	if vt := u.OfPropertyBased; vt != nil {
		return &vt.DaysDelta
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIPlatformFlowPutRequestEnrollmentScheduleUnionParam) GetYearly() *bool {
	if vt := u.OfPropertyBased; vt != nil {
		return &vt.Yearly
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIPlatformFlowPutRequestEnrollmentScheduleUnionParam) GetType() *string {
	if vt := u.OfDaily; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfWeekly; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfMonthlySpecificDays; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfMonthlyRelativeDays; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfYearly; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfPropertyBased; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

// Returns a pointer to the underlying variant's TimeOfDay property, if present.
func (u APIPlatformFlowPutRequestEnrollmentScheduleUnionParam) GetTimeOfDay() *APITimeOfDayParam {
	if vt := u.OfDaily; vt != nil {
		return &vt.TimeOfDay
	} else if vt := u.OfWeekly; vt != nil {
		return &vt.TimeOfDay
	} else if vt := u.OfMonthlySpecificDays; vt != nil {
		return &vt.TimeOfDay
	} else if vt := u.OfMonthlyRelativeDays; vt != nil {
		return &vt.TimeOfDay
	} else if vt := u.OfYearly; vt != nil {
		return &vt.TimeOfDay
	} else if vt := u.OfPropertyBased; vt != nil {
		return &vt.TimeOfDay
	}
	return nil
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type APIPlatformFlowPutRequestSuppressionFilterBranchUnionParam struct {
	OfOr                  *shared.PublicOrFilterBranchParam                  `json:",omitzero,inline"`
	OfAnd                 *shared.PublicAndFilterBranchParam                 `json:",omitzero,inline"`
	OfNotAll              *shared.PublicNotAllFilterBranchParam              `json:",omitzero,inline"`
	OfNotAny              *shared.PublicNotAnyFilterBranchParam              `json:",omitzero,inline"`
	OfRestricted          *shared.PublicRestrictedFilterBranchParam          `json:",omitzero,inline"`
	OfUnifiedEvents       *shared.PublicUnifiedEventsFilterBranchParam       `json:",omitzero,inline"`
	OfPropertyAssociation *shared.PublicPropertyAssociationFilterBranchParam `json:",omitzero,inline"`
	OfAssociation         *shared.PublicAssociationFilterBranchParam         `json:",omitzero,inline"`
	paramUnion
}

func (u APIPlatformFlowPutRequestSuppressionFilterBranchUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfOr,
		u.OfAnd,
		u.OfNotAll,
		u.OfNotAny,
		u.OfRestricted,
		u.OfUnifiedEvents,
		u.OfPropertyAssociation,
		u.OfAssociation)
}
func (u *APIPlatformFlowPutRequestSuppressionFilterBranchUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *APIPlatformFlowPutRequestSuppressionFilterBranchUnionParam) asAny() any {
	if !param.IsOmitted(u.OfOr) {
		return u.OfOr
	} else if !param.IsOmitted(u.OfAnd) {
		return u.OfAnd
	} else if !param.IsOmitted(u.OfNotAll) {
		return u.OfNotAll
	} else if !param.IsOmitted(u.OfNotAny) {
		return u.OfNotAny
	} else if !param.IsOmitted(u.OfRestricted) {
		return u.OfRestricted
	} else if !param.IsOmitted(u.OfUnifiedEvents) {
		return u.OfUnifiedEvents
	} else if !param.IsOmitted(u.OfPropertyAssociation) {
		return u.OfPropertyAssociation
	} else if !param.IsOmitted(u.OfAssociation) {
		return u.OfAssociation
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIPlatformFlowPutRequestSuppressionFilterBranchUnionParam) GetEventTypeID() *string {
	if vt := u.OfUnifiedEvents; vt != nil {
		return &vt.EventTypeID
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIPlatformFlowPutRequestSuppressionFilterBranchUnionParam) GetCoalescingRefineBy() *shared.PublicUnifiedEventsFilterBranchCoalescingRefineByUnionParam {
	if vt := u.OfUnifiedEvents; vt != nil {
		return &vt.CoalescingRefineBy
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIPlatformFlowPutRequestSuppressionFilterBranchUnionParam) GetPropertyWithObjectID() *string {
	if vt := u.OfPropertyAssociation; vt != nil {
		return &vt.PropertyWithObjectID
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIPlatformFlowPutRequestSuppressionFilterBranchUnionParam) GetAssociationCategory() *string {
	if vt := u.OfAssociation; vt != nil {
		return &vt.AssociationCategory
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIPlatformFlowPutRequestSuppressionFilterBranchUnionParam) GetAssociationTypeID() *int64 {
	if vt := u.OfAssociation; vt != nil {
		return &vt.AssociationTypeID
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIPlatformFlowPutRequestSuppressionFilterBranchUnionParam) GetFilterBranchOperator() *string {
	if vt := u.OfOr; vt != nil {
		return (*string)(&vt.FilterBranchOperator)
	} else if vt := u.OfAnd; vt != nil {
		return (*string)(&vt.FilterBranchOperator)
	} else if vt := u.OfNotAll; vt != nil {
		return (*string)(&vt.FilterBranchOperator)
	} else if vt := u.OfNotAny; vt != nil {
		return (*string)(&vt.FilterBranchOperator)
	} else if vt := u.OfRestricted; vt != nil {
		return (*string)(&vt.FilterBranchOperator)
	} else if vt := u.OfUnifiedEvents; vt != nil {
		return (*string)(&vt.FilterBranchOperator)
	} else if vt := u.OfPropertyAssociation; vt != nil {
		return (*string)(&vt.FilterBranchOperator)
	} else if vt := u.OfAssociation; vt != nil {
		return (*string)(&vt.FilterBranchOperator)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIPlatformFlowPutRequestSuppressionFilterBranchUnionParam) GetFilterBranchType() *string {
	if vt := u.OfOr; vt != nil {
		return (*string)(&vt.FilterBranchType)
	} else if vt := u.OfAnd; vt != nil {
		return (*string)(&vt.FilterBranchType)
	} else if vt := u.OfNotAll; vt != nil {
		return (*string)(&vt.FilterBranchType)
	} else if vt := u.OfNotAny; vt != nil {
		return (*string)(&vt.FilterBranchType)
	} else if vt := u.OfRestricted; vt != nil {
		return (*string)(&vt.FilterBranchType)
	} else if vt := u.OfUnifiedEvents; vt != nil {
		return (*string)(&vt.FilterBranchType)
	} else if vt := u.OfPropertyAssociation; vt != nil {
		return (*string)(&vt.FilterBranchType)
	} else if vt := u.OfAssociation; vt != nil {
		return (*string)(&vt.FilterBranchType)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIPlatformFlowPutRequestSuppressionFilterBranchUnionParam) GetOperator() *string {
	if vt := u.OfUnifiedEvents; vt != nil {
		return (*string)(&vt.Operator)
	} else if vt := u.OfPropertyAssociation; vt != nil {
		return (*string)(&vt.Operator)
	} else if vt := u.OfAssociation; vt != nil {
		return (*string)(&vt.Operator)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIPlatformFlowPutRequestSuppressionFilterBranchUnionParam) GetObjectTypeID() *string {
	if vt := u.OfPropertyAssociation; vt != nil {
		return (*string)(&vt.ObjectTypeID)
	} else if vt := u.OfAssociation; vt != nil {
		return (*string)(&vt.ObjectTypeID)
	}
	return nil
}

// Returns a subunion which exports methods to access subproperties
//
// Or use AsAny() to get the underlying value
func (u APIPlatformFlowPutRequestSuppressionFilterBranchUnionParam) GetFilterBranches() (res apiPlatformFlowPutRequestSuppressionFilterBranchUnionParamFilterBranches) {
	if vt := u.OfOr; vt != nil {
		res.any = &vt.FilterBranches
	} else if vt := u.OfAnd; vt != nil {
		res.any = &vt.FilterBranches
	} else if vt := u.OfNotAll; vt != nil {
		res.any = &vt.FilterBranches
	} else if vt := u.OfNotAny; vt != nil {
		res.any = &vt.FilterBranches
	} else if vt := u.OfRestricted; vt != nil {
		res.any = &vt.FilterBranches
	} else if vt := u.OfUnifiedEvents; vt != nil {
		res.any = &vt.FilterBranches
	} else if vt := u.OfPropertyAssociation; vt != nil {
		res.any = &vt.FilterBranches
	} else if vt := u.OfAssociation; vt != nil {
		res.any = &vt.FilterBranches
	}
	return
}

// Can have the runtime types
// [_[]shared.PublicOrFilterBranchFilterBranchUnionParam],
// [_[]shared.PublicAndFilterBranchFilterBranchUnionParam],
// [_[]shared.PublicNotAllFilterBranchFilterBranchUnionParam],
// [_[]shared.PublicNotAnyFilterBranchFilterBranchUnionParam],
// [_[]shared.PublicRestrictedFilterBranchFilterBranchUnionParam],
// [_[]shared.PublicUnifiedEventsFilterBranchFilterBranchUnionParam],
// [_[]shared.PublicPropertyAssociationFilterBranchFilterBranchUnionParam],
// [_[]shared.PublicAssociationFilterBranchFilterBranchUnionParam]
type apiPlatformFlowPutRequestSuppressionFilterBranchUnionParamFilterBranches struct{ any }

// Use the following switch statement to get the type of the union:
//
//	switch u.AsAny().(type) {
//	case *[]shared.PublicOrFilterBranchFilterBranchUnionParam:
//	case *[]shared.PublicAndFilterBranchFilterBranchUnionParam:
//	case *[]shared.PublicNotAllFilterBranchFilterBranchUnionParam:
//	case *[]shared.PublicNotAnyFilterBranchFilterBranchUnionParam:
//	case *[]shared.PublicRestrictedFilterBranchFilterBranchUnionParam:
//	case *[]shared.PublicUnifiedEventsFilterBranchFilterBranchUnionParam:
//	case *[]shared.PublicPropertyAssociationFilterBranchFilterBranchUnionParam:
//	case *[]shared.PublicAssociationFilterBranchFilterBranchUnionParam:
//	default:
//	    fmt.Errorf("not present")
//	}
func (u apiPlatformFlowPutRequestSuppressionFilterBranchUnionParamFilterBranches) AsAny() any {
	return u.any
}

// Returns a subunion which exports methods to access subproperties
//
// Or use AsAny() to get the underlying value
func (u APIPlatformFlowPutRequestSuppressionFilterBranchUnionParam) GetFilters() (res apiPlatformFlowPutRequestSuppressionFilterBranchUnionParamFilters) {
	if vt := u.OfOr; vt != nil {
		res.any = &vt.Filters
	} else if vt := u.OfAnd; vt != nil {
		res.any = &vt.Filters
	} else if vt := u.OfNotAll; vt != nil {
		res.any = &vt.Filters
	} else if vt := u.OfNotAny; vt != nil {
		res.any = &vt.Filters
	} else if vt := u.OfRestricted; vt != nil {
		res.any = &vt.Filters
	} else if vt := u.OfUnifiedEvents; vt != nil {
		res.any = &vt.Filters
	} else if vt := u.OfPropertyAssociation; vt != nil {
		res.any = &vt.Filters
	} else if vt := u.OfAssociation; vt != nil {
		res.any = &vt.Filters
	}
	return
}

// Can have the runtime types [_[]shared.PublicOrFilterBranchFilterUnionParam],
// [_[]shared.PublicAndFilterBranchFilterUnionParam],
// [_[]shared.PublicNotAllFilterBranchFilterUnionParam],
// [_[]shared.PublicNotAnyFilterBranchFilterUnionParam],
// [_[]shared.PublicRestrictedFilterBranchFilterUnionParam],
// [_[]shared.PublicUnifiedEventsFilterBranchFilterUnionParam],
// [_[]shared.PublicPropertyAssociationFilterBranchFilterUnionParam],
// [_[]shared.PublicAssociationFilterBranchFilterUnionParam]
type apiPlatformFlowPutRequestSuppressionFilterBranchUnionParamFilters struct{ any }

// Use the following switch statement to get the type of the union:
//
//	switch u.AsAny().(type) {
//	case *[]shared.PublicOrFilterBranchFilterUnionParam:
//	case *[]shared.PublicAndFilterBranchFilterUnionParam:
//	case *[]shared.PublicNotAllFilterBranchFilterUnionParam:
//	case *[]shared.PublicNotAnyFilterBranchFilterUnionParam:
//	case *[]shared.PublicRestrictedFilterBranchFilterUnionParam:
//	case *[]shared.PublicUnifiedEventsFilterBranchFilterUnionParam:
//	case *[]shared.PublicPropertyAssociationFilterBranchFilterUnionParam:
//	case *[]shared.PublicAssociationFilterBranchFilterUnionParam:
//	default:
//	    fmt.Errorf("not present")
//	}
func (u apiPlatformFlowPutRequestSuppressionFilterBranchUnionParamFilters) AsAny() any { return u.any }

type APIPropertyBasedEnrollmentSchedule struct {
	DateProperty string       `json:"dateProperty,required"`
	DaysDelta    int64        `json:"daysDelta,required"`
	TimeOfDay    APITimeOfDay `json:"timeOfDay,required"`
	// Any of "PROPERTY_BASED".
	Type   APIPropertyBasedEnrollmentScheduleType `json:"type,required"`
	Yearly bool                                   `json:"yearly,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DateProperty respjson.Field
		DaysDelta    respjson.Field
		TimeOfDay    respjson.Field
		Type         respjson.Field
		Yearly       respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r APIPropertyBasedEnrollmentSchedule) RawJSON() string { return r.JSON.raw }
func (r *APIPropertyBasedEnrollmentSchedule) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this APIPropertyBasedEnrollmentSchedule to a
// APIPropertyBasedEnrollmentScheduleParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// APIPropertyBasedEnrollmentScheduleParam.Overrides()
func (r APIPropertyBasedEnrollmentSchedule) ToParam() APIPropertyBasedEnrollmentScheduleParam {
	return param.Override[APIPropertyBasedEnrollmentScheduleParam](json.RawMessage(r.RawJSON()))
}

type APIPropertyBasedEnrollmentScheduleType string

const (
	APIPropertyBasedEnrollmentScheduleTypePropertyBased APIPropertyBasedEnrollmentScheduleType = "PROPERTY_BASED"
)

// The properties DateProperty, DaysDelta, TimeOfDay, Type, Yearly are required.
type APIPropertyBasedEnrollmentScheduleParam struct {
	DateProperty string            `json:"dateProperty,required"`
	DaysDelta    int64             `json:"daysDelta,required"`
	TimeOfDay    APITimeOfDayParam `json:"timeOfDay,omitzero,required"`
	// Any of "PROPERTY_BASED".
	Type   APIPropertyBasedEnrollmentScheduleType `json:"type,omitzero,required"`
	Yearly bool                                   `json:"yearly,required"`
	paramObj
}

func (r APIPropertyBasedEnrollmentScheduleParam) MarshalJSON() (data []byte, err error) {
	type shadow APIPropertyBasedEnrollmentScheduleParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *APIPropertyBasedEnrollmentScheduleParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type APIRelativeDateTimeValue struct {
	TimeDelay APITimeDelay `json:"timeDelay,required"`
	// This is the type of input value. This can be one of: "FIELD_DATA",
	// "OBJECT_PROPERTY", "STATIC_VALUE", "RELATIVE_DATETIME", "TIMESTAMP",
	// "INCREMENT", "FETCHED_OBJECT_PROPERTY", "APPEND_OBJECT_PROPERTY",
	// "STATIC_APPEND_VALUE", "ENROLLMENT_EVENT_PROPERTY"
	//
	// Any of "RELATIVE_DATETIME".
	Type APIRelativeDateTimeValueType `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		TimeDelay   respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r APIRelativeDateTimeValue) RawJSON() string { return r.JSON.raw }
func (r *APIRelativeDateTimeValue) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this APIRelativeDateTimeValue to a
// APIRelativeDateTimeValueParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// APIRelativeDateTimeValueParam.Overrides()
func (r APIRelativeDateTimeValue) ToParam() APIRelativeDateTimeValueParam {
	return param.Override[APIRelativeDateTimeValueParam](json.RawMessage(r.RawJSON()))
}

// This is the type of input value. This can be one of: "FIELD_DATA",
// "OBJECT_PROPERTY", "STATIC_VALUE", "RELATIVE_DATETIME", "TIMESTAMP",
// "INCREMENT", "FETCHED_OBJECT_PROPERTY", "APPEND_OBJECT_PROPERTY",
// "STATIC_APPEND_VALUE", "ENROLLMENT_EVENT_PROPERTY"
type APIRelativeDateTimeValueType string

const (
	APIRelativeDateTimeValueTypeRelativeDatetime APIRelativeDateTimeValueType = "RELATIVE_DATETIME"
)

// The properties TimeDelay, Type are required.
type APIRelativeDateTimeValueParam struct {
	TimeDelay APITimeDelayParam `json:"timeDelay,omitzero,required"`
	// This is the type of input value. This can be one of: "FIELD_DATA",
	// "OBJECT_PROPERTY", "STATIC_VALUE", "RELATIVE_DATETIME", "TIMESTAMP",
	// "INCREMENT", "FETCHED_OBJECT_PROPERTY", "APPEND_OBJECT_PROPERTY",
	// "STATIC_APPEND_VALUE", "ENROLLMENT_EVENT_PROPERTY"
	//
	// Any of "RELATIVE_DATETIME".
	Type APIRelativeDateTimeValueType `json:"type,omitzero,required"`
	paramObj
}

func (r APIRelativeDateTimeValueParam) MarshalJSON() (data []byte, err error) {
	type shadow APIRelativeDateTimeValueParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *APIRelativeDateTimeValueParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type APISignatureWebhookAuthSettings struct {
	// The appId that this signature will be generated for.
	AppID int64 `json:"appId,required"`
	// The type of webhook auth settings this is, can be: "AUTH_KEY" or "SIGNATURE"
	//
	// Any of "SIGNATURE".
	Type APISignatureWebhookAuthSettingsType `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AppID       respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r APISignatureWebhookAuthSettings) RawJSON() string { return r.JSON.raw }
func (r *APISignatureWebhookAuthSettings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this APISignatureWebhookAuthSettings to a
// APISignatureWebhookAuthSettingsParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// APISignatureWebhookAuthSettingsParam.Overrides()
func (r APISignatureWebhookAuthSettings) ToParam() APISignatureWebhookAuthSettingsParam {
	return param.Override[APISignatureWebhookAuthSettingsParam](json.RawMessage(r.RawJSON()))
}

// The type of webhook auth settings this is, can be: "AUTH_KEY" or "SIGNATURE"
type APISignatureWebhookAuthSettingsType string

const (
	APISignatureWebhookAuthSettingsTypeSignature APISignatureWebhookAuthSettingsType = "SIGNATURE"
)

// The properties AppID, Type are required.
type APISignatureWebhookAuthSettingsParam struct {
	// The appId that this signature will be generated for.
	AppID int64 `json:"appId,required"`
	// The type of webhook auth settings this is, can be: "AUTH_KEY" or "SIGNATURE"
	//
	// Any of "SIGNATURE".
	Type APISignatureWebhookAuthSettingsType `json:"type,omitzero,required"`
	paramObj
}

func (r APISignatureWebhookAuthSettingsParam) MarshalJSON() (data []byte, err error) {
	type shadow APISignatureWebhookAuthSettingsParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *APISignatureWebhookAuthSettingsParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type APISingleConnectionAction struct {
	// The ID for this action.
	ActionID string `json:"actionId,required"`
	// The ID of the actionType to use.
	ActionTypeID string `json:"actionTypeId,required"`
	// The version of this actionType to use.
	ActionTypeVersion int64 `json:"actionTypeVersion,required"`
	// The fields to pass into this action. Different action types accept different
	// fields.
	Fields map[string]any `json:"fields,required"`
	// The type of action this is, can be: "STATIC_BRANCH", "LIST_BRANCH",
	// "AB_TEST_BRANCH", "CUSTOM_CODE", "WEBHOOK", or "SINGLE_CONNECTION"
	//
	// Any of "SINGLE_CONNECTION".
	Type       APISingleConnectionActionType `json:"type,required"`
	Connection APIConnection                 `json:"connection"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ActionID          respjson.Field
		ActionTypeID      respjson.Field
		ActionTypeVersion respjson.Field
		Fields            respjson.Field
		Type              respjson.Field
		Connection        respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r APISingleConnectionAction) RawJSON() string { return r.JSON.raw }
func (r *APISingleConnectionAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this APISingleConnectionAction to a
// APISingleConnectionActionParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// APISingleConnectionActionParam.Overrides()
func (r APISingleConnectionAction) ToParam() APISingleConnectionActionParam {
	return param.Override[APISingleConnectionActionParam](json.RawMessage(r.RawJSON()))
}

// The type of action this is, can be: "STATIC_BRANCH", "LIST_BRANCH",
// "AB_TEST_BRANCH", "CUSTOM_CODE", "WEBHOOK", or "SINGLE_CONNECTION"
type APISingleConnectionActionType string

const (
	APISingleConnectionActionTypeSingleConnection APISingleConnectionActionType = "SINGLE_CONNECTION"
)

// The properties ActionID, ActionTypeID, ActionTypeVersion, Fields, Type are
// required.
type APISingleConnectionActionParam struct {
	// The ID for this action.
	ActionID string `json:"actionId,required"`
	// The ID of the actionType to use.
	ActionTypeID string `json:"actionTypeId,required"`
	// The version of this actionType to use.
	ActionTypeVersion int64 `json:"actionTypeVersion,required"`
	// The fields to pass into this action. Different action types accept different
	// fields.
	Fields map[string]any `json:"fields,omitzero,required"`
	// The type of action this is, can be: "STATIC_BRANCH", "LIST_BRANCH",
	// "AB_TEST_BRANCH", "CUSTOM_CODE", "WEBHOOK", or "SINGLE_CONNECTION"
	//
	// Any of "SINGLE_CONNECTION".
	Type       APISingleConnectionActionType `json:"type,omitzero,required"`
	Connection APIConnectionParam            `json:"connection,omitzero"`
	paramObj
}

func (r APISingleConnectionActionParam) MarshalJSON() (data []byte, err error) {
	type shadow APISingleConnectionActionParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *APISingleConnectionActionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type APISort struct {
	// Any of "ASC", "DESC".
	Order    APISortOrder `json:"order,required"`
	Property string       `json:"property,required"`
	Missing  string       `json:"missing"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Order       respjson.Field
		Property    respjson.Field
		Missing     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r APISort) RawJSON() string { return r.JSON.raw }
func (r *APISort) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this APISort to a APISortParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// APISortParam.Overrides()
func (r APISort) ToParam() APISortParam {
	return param.Override[APISortParam](json.RawMessage(r.RawJSON()))
}

type APISortOrder string

const (
	APISortOrderAsc  APISortOrder = "ASC"
	APISortOrderDesc APISortOrder = "DESC"
)

// The properties Order, Property are required.
type APISortParam struct {
	// Any of "ASC", "DESC".
	Order    APISortOrder      `json:"order,omitzero,required"`
	Property string            `json:"property,required"`
	Missing  param.Opt[string] `json:"missing,omitzero"`
	paramObj
}

func (r APISortParam) MarshalJSON() (data []byte, err error) {
	type shadow APISortParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *APISortParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type APIStaticAppendValue struct {
	// The value to append
	StaticAppendValue string `json:"staticAppendValue,required"`
	// This is the type of input value. This can be one of: "FIELD_DATA",
	// "OBJECT_PROPERTY", "STATIC_VALUE", "RELATIVE_DATETIME", "TIMESTAMP",
	// "INCREMENT", "FETCHED_OBJECT_PROPERTY", "APPEND_OBJECT_PROPERTY",
	// "STATIC_APPEND_VALUE", "ENROLLMENT_EVENT_PROPERTY"
	//
	// Any of "STATIC_APPEND_VALUE".
	Type APIStaticAppendValueType `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		StaticAppendValue respjson.Field
		Type              respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r APIStaticAppendValue) RawJSON() string { return r.JSON.raw }
func (r *APIStaticAppendValue) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this APIStaticAppendValue to a APIStaticAppendValueParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// APIStaticAppendValueParam.Overrides()
func (r APIStaticAppendValue) ToParam() APIStaticAppendValueParam {
	return param.Override[APIStaticAppendValueParam](json.RawMessage(r.RawJSON()))
}

// This is the type of input value. This can be one of: "FIELD_DATA",
// "OBJECT_PROPERTY", "STATIC_VALUE", "RELATIVE_DATETIME", "TIMESTAMP",
// "INCREMENT", "FETCHED_OBJECT_PROPERTY", "APPEND_OBJECT_PROPERTY",
// "STATIC_APPEND_VALUE", "ENROLLMENT_EVENT_PROPERTY"
type APIStaticAppendValueType string

const (
	APIStaticAppendValueTypeStaticAppendValue APIStaticAppendValueType = "STATIC_APPEND_VALUE"
)

// The properties StaticAppendValue, Type are required.
type APIStaticAppendValueParam struct {
	// The value to append
	StaticAppendValue string `json:"staticAppendValue,required"`
	// This is the type of input value. This can be one of: "FIELD_DATA",
	// "OBJECT_PROPERTY", "STATIC_VALUE", "RELATIVE_DATETIME", "TIMESTAMP",
	// "INCREMENT", "FETCHED_OBJECT_PROPERTY", "APPEND_OBJECT_PROPERTY",
	// "STATIC_APPEND_VALUE", "ENROLLMENT_EVENT_PROPERTY"
	//
	// Any of "STATIC_APPEND_VALUE".
	Type APIStaticAppendValueType `json:"type,omitzero,required"`
	paramObj
}

func (r APIStaticAppendValueParam) MarshalJSON() (data []byte, err error) {
	type shadow APIStaticAppendValueParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *APIStaticAppendValueParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type APIStaticBranch struct {
	// If value to check for. If the value of the `inputValue` matches this
	// `branchValue` than this `connection` will get traversed.
	BranchValue string        `json:"branchValue,required"`
	Connection  APIConnection `json:"connection"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BranchValue respjson.Field
		Connection  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r APIStaticBranch) RawJSON() string { return r.JSON.raw }
func (r *APIStaticBranch) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this APIStaticBranch to a APIStaticBranchParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// APIStaticBranchParam.Overrides()
func (r APIStaticBranch) ToParam() APIStaticBranchParam {
	return param.Override[APIStaticBranchParam](json.RawMessage(r.RawJSON()))
}

// The property BranchValue is required.
type APIStaticBranchParam struct {
	// If value to check for. If the value of the `inputValue` matches this
	// `branchValue` than this `connection` will get traversed.
	BranchValue string             `json:"branchValue,required"`
	Connection  APIConnectionParam `json:"connection,omitzero"`
	paramObj
}

func (r APIStaticBranchParam) MarshalJSON() (data []byte, err error) {
	type shadow APIStaticBranchParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *APIStaticBranchParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type APIStaticBranchAction struct {
	// The ID for this action.
	ActionID string `json:"actionId,required"`
	// The input value to branch off of.
	InputValue     APIStaticBranchActionInputValueUnion `json:"inputValue,required"`
	StaticBranches []APIStaticBranch                    `json:"staticBranches,required"`
	// The type of action this is, can be: "STATIC_BRANCH", "LIST_BRANCH",
	// "AB_TEST_BRANCH", "CUSTOM_CODE", "WEBHOOK", or "SINGLE_CONNECTION"
	//
	// Any of "STATIC_BRANCH".
	Type          APIStaticBranchActionType `json:"type,required"`
	DefaultBranch APIConnection             `json:"defaultBranch"`
	// The name of the default branch, the branch that gets executed if `inputValue`
	// does not match any of the `staticBranches`.
	DefaultBranchName string `json:"defaultBranchName"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ActionID          respjson.Field
		InputValue        respjson.Field
		StaticBranches    respjson.Field
		Type              respjson.Field
		DefaultBranch     respjson.Field
		DefaultBranchName respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r APIStaticBranchAction) RawJSON() string { return r.JSON.raw }
func (r *APIStaticBranchAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this APIStaticBranchAction to a APIStaticBranchActionParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// APIStaticBranchActionParam.Overrides()
func (r APIStaticBranchAction) ToParam() APIStaticBranchActionParam {
	return param.Override[APIStaticBranchActionParam](json.RawMessage(r.RawJSON()))
}

// APIStaticBranchActionInputValueUnion contains all possible properties and values
// from [APIActionDataValue], [APIObjectPropertyValue], [APIStaticValue],
// [APIRelativeDateTimeValue], [APITimestampValue], [APIIncrementValue],
// [APIFetchedObjectPropertyValue], [APIAppendObjectPropertyValue],
// [APIStaticAppendValue], [APIEnrollmentEventPropertyValue].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type APIStaticBranchActionInputValueUnion struct {
	// This field is from variant [APIActionDataValue].
	ActionID string `json:"actionId"`
	// This field is from variant [APIActionDataValue].
	DataKey string `json:"dataKey"`
	Type    string `json:"type"`
	// This field is from variant [APIObjectPropertyValue].
	PropertyName string `json:"propertyName"`
	// This field is from variant [APIStaticValue].
	StaticValue string `json:"staticValue"`
	// This field is from variant [APIRelativeDateTimeValue].
	TimeDelay APITimeDelay `json:"timeDelay"`
	// This field is from variant [APITimestampValue].
	TimestampType APITimestampValueTimestampType `json:"timestampType"`
	// This field is from variant [APIIncrementValue].
	IncrementAmount float64 `json:"incrementAmount"`
	// This field is from variant [APIFetchedObjectPropertyValue].
	PropertyToken string `json:"propertyToken"`
	// This field is from variant [APIAppendObjectPropertyValue].
	AppendPropertyName string `json:"appendPropertyName"`
	// This field is from variant [APIStaticAppendValue].
	StaticAppendValue string `json:"staticAppendValue"`
	// This field is from variant [APIEnrollmentEventPropertyValue].
	EnrollmentEventPropertyToken string `json:"enrollmentEventPropertyToken"`
	JSON                         struct {
		ActionID                     respjson.Field
		DataKey                      respjson.Field
		Type                         respjson.Field
		PropertyName                 respjson.Field
		StaticValue                  respjson.Field
		TimeDelay                    respjson.Field
		TimestampType                respjson.Field
		IncrementAmount              respjson.Field
		PropertyToken                respjson.Field
		AppendPropertyName           respjson.Field
		StaticAppendValue            respjson.Field
		EnrollmentEventPropertyToken respjson.Field
		raw                          string
	} `json:"-"`
}

func (u APIStaticBranchActionInputValueUnion) AsFieldData() (v APIActionDataValue) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u APIStaticBranchActionInputValueUnion) AsObjectProperty() (v APIObjectPropertyValue) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u APIStaticBranchActionInputValueUnion) AsStaticValue() (v APIStaticValue) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u APIStaticBranchActionInputValueUnion) AsRelativeDatetime() (v APIRelativeDateTimeValue) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u APIStaticBranchActionInputValueUnion) AsTimestamp() (v APITimestampValue) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u APIStaticBranchActionInputValueUnion) AsIncrement() (v APIIncrementValue) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u APIStaticBranchActionInputValueUnion) AsFetchedObjectProperty() (v APIFetchedObjectPropertyValue) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u APIStaticBranchActionInputValueUnion) AsAppendObjectProperty() (v APIAppendObjectPropertyValue) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u APIStaticBranchActionInputValueUnion) AsStaticAppendValue() (v APIStaticAppendValue) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u APIStaticBranchActionInputValueUnion) AsEnrollmentEventProperty() (v APIEnrollmentEventPropertyValue) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u APIStaticBranchActionInputValueUnion) RawJSON() string { return u.JSON.raw }

func (r *APIStaticBranchActionInputValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The type of action this is, can be: "STATIC_BRANCH", "LIST_BRANCH",
// "AB_TEST_BRANCH", "CUSTOM_CODE", "WEBHOOK", or "SINGLE_CONNECTION"
type APIStaticBranchActionType string

const (
	APIStaticBranchActionTypeStaticBranch APIStaticBranchActionType = "STATIC_BRANCH"
)

// The properties ActionID, InputValue, StaticBranches, Type are required.
type APIStaticBranchActionParam struct {
	// The ID for this action.
	ActionID string `json:"actionId,required"`
	// The input value to branch off of.
	InputValue     APIStaticBranchActionInputValueUnionParam `json:"inputValue,omitzero,required"`
	StaticBranches []APIStaticBranchParam                    `json:"staticBranches,omitzero,required"`
	// The type of action this is, can be: "STATIC_BRANCH", "LIST_BRANCH",
	// "AB_TEST_BRANCH", "CUSTOM_CODE", "WEBHOOK", or "SINGLE_CONNECTION"
	//
	// Any of "STATIC_BRANCH".
	Type APIStaticBranchActionType `json:"type,omitzero,required"`
	// The name of the default branch, the branch that gets executed if `inputValue`
	// does not match any of the `staticBranches`.
	DefaultBranchName param.Opt[string]  `json:"defaultBranchName,omitzero"`
	DefaultBranch     APIConnectionParam `json:"defaultBranch,omitzero"`
	paramObj
}

func (r APIStaticBranchActionParam) MarshalJSON() (data []byte, err error) {
	type shadow APIStaticBranchActionParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *APIStaticBranchActionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type APIStaticBranchActionInputValueUnionParam struct {
	OfFieldData               *APIActionDataValueParam              `json:",omitzero,inline"`
	OfObjectProperty          *APIObjectPropertyValueParam          `json:",omitzero,inline"`
	OfStaticValue             *APIStaticValueParam                  `json:",omitzero,inline"`
	OfRelativeDatetime        *APIRelativeDateTimeValueParam        `json:",omitzero,inline"`
	OfTimestamp               *APITimestampValueParam               `json:",omitzero,inline"`
	OfIncrement               *APIIncrementValueParam               `json:",omitzero,inline"`
	OfFetchedObjectProperty   *APIFetchedObjectPropertyValueParam   `json:",omitzero,inline"`
	OfAppendObjectProperty    *APIAppendObjectPropertyValueParam    `json:",omitzero,inline"`
	OfStaticAppendValue       *APIStaticAppendValueParam            `json:",omitzero,inline"`
	OfEnrollmentEventProperty *APIEnrollmentEventPropertyValueParam `json:",omitzero,inline"`
	paramUnion
}

func (u APIStaticBranchActionInputValueUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfFieldData,
		u.OfObjectProperty,
		u.OfStaticValue,
		u.OfRelativeDatetime,
		u.OfTimestamp,
		u.OfIncrement,
		u.OfFetchedObjectProperty,
		u.OfAppendObjectProperty,
		u.OfStaticAppendValue,
		u.OfEnrollmentEventProperty)
}
func (u *APIStaticBranchActionInputValueUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *APIStaticBranchActionInputValueUnionParam) asAny() any {
	if !param.IsOmitted(u.OfFieldData) {
		return u.OfFieldData
	} else if !param.IsOmitted(u.OfObjectProperty) {
		return u.OfObjectProperty
	} else if !param.IsOmitted(u.OfStaticValue) {
		return u.OfStaticValue
	} else if !param.IsOmitted(u.OfRelativeDatetime) {
		return u.OfRelativeDatetime
	} else if !param.IsOmitted(u.OfTimestamp) {
		return u.OfTimestamp
	} else if !param.IsOmitted(u.OfIncrement) {
		return u.OfIncrement
	} else if !param.IsOmitted(u.OfFetchedObjectProperty) {
		return u.OfFetchedObjectProperty
	} else if !param.IsOmitted(u.OfAppendObjectProperty) {
		return u.OfAppendObjectProperty
	} else if !param.IsOmitted(u.OfStaticAppendValue) {
		return u.OfStaticAppendValue
	} else if !param.IsOmitted(u.OfEnrollmentEventProperty) {
		return u.OfEnrollmentEventProperty
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIStaticBranchActionInputValueUnionParam) GetActionID() *string {
	if vt := u.OfFieldData; vt != nil {
		return &vt.ActionID
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIStaticBranchActionInputValueUnionParam) GetDataKey() *string {
	if vt := u.OfFieldData; vt != nil {
		return &vt.DataKey
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIStaticBranchActionInputValueUnionParam) GetPropertyName() *string {
	if vt := u.OfObjectProperty; vt != nil {
		return &vt.PropertyName
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIStaticBranchActionInputValueUnionParam) GetStaticValue() *string {
	if vt := u.OfStaticValue; vt != nil {
		return &vt.StaticValue
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIStaticBranchActionInputValueUnionParam) GetTimeDelay() *APITimeDelayParam {
	if vt := u.OfRelativeDatetime; vt != nil {
		return &vt.TimeDelay
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIStaticBranchActionInputValueUnionParam) GetTimestampType() *string {
	if vt := u.OfTimestamp; vt != nil {
		return (*string)(&vt.TimestampType)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIStaticBranchActionInputValueUnionParam) GetIncrementAmount() *float64 {
	if vt := u.OfIncrement; vt != nil {
		return &vt.IncrementAmount
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIStaticBranchActionInputValueUnionParam) GetPropertyToken() *string {
	if vt := u.OfFetchedObjectProperty; vt != nil {
		return &vt.PropertyToken
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIStaticBranchActionInputValueUnionParam) GetAppendPropertyName() *string {
	if vt := u.OfAppendObjectProperty; vt != nil {
		return &vt.AppendPropertyName
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIStaticBranchActionInputValueUnionParam) GetStaticAppendValue() *string {
	if vt := u.OfStaticAppendValue; vt != nil {
		return &vt.StaticAppendValue
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIStaticBranchActionInputValueUnionParam) GetEnrollmentEventPropertyToken() *string {
	if vt := u.OfEnrollmentEventProperty; vt != nil {
		return &vt.EnrollmentEventPropertyToken
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIStaticBranchActionInputValueUnionParam) GetType() *string {
	if vt := u.OfFieldData; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfObjectProperty; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfStaticValue; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfRelativeDatetime; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfTimestamp; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfIncrement; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfFetchedObjectProperty; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfAppendObjectProperty; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfStaticAppendValue; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfEnrollmentEventProperty; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

type APIStaticDateAnchor struct {
	// The day of the date to anchor on
	DayOfMonth int64 `json:"dayOfMonth,required"`
	// The month of the date to anchor on
	//
	// Any of "JANUARY", "FEBRUARY", "MARCH", "APRIL", "MAY", "JUNE", "JULY", "AUGUST",
	// "SEPTEMBER", "OCTOBER", "NOVEMBER", "DECEMBER".
	Month APIStaticDateAnchorMonth `json:"month,required"`
	// The type of event anchor this is, can be: "CONTACT_PROPERTY_ANCHOR" or
	// "STATIC_DATE_ANCHOR"
	//
	// Any of "STATIC_DATE_ANCHOR".
	Type APIStaticDateAnchorType `json:"type,required"`
	// The year of the date to anchor on. If this is not provided then this flow will
	// re-run each year.
	Year int64 `json:"year"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DayOfMonth  respjson.Field
		Month       respjson.Field
		Type        respjson.Field
		Year        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r APIStaticDateAnchor) RawJSON() string { return r.JSON.raw }
func (r *APIStaticDateAnchor) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this APIStaticDateAnchor to a APIStaticDateAnchorParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// APIStaticDateAnchorParam.Overrides()
func (r APIStaticDateAnchor) ToParam() APIStaticDateAnchorParam {
	return param.Override[APIStaticDateAnchorParam](json.RawMessage(r.RawJSON()))
}

// The month of the date to anchor on
type APIStaticDateAnchorMonth string

const (
	APIStaticDateAnchorMonthJanuary   APIStaticDateAnchorMonth = "JANUARY"
	APIStaticDateAnchorMonthFebruary  APIStaticDateAnchorMonth = "FEBRUARY"
	APIStaticDateAnchorMonthMarch     APIStaticDateAnchorMonth = "MARCH"
	APIStaticDateAnchorMonthApril     APIStaticDateAnchorMonth = "APRIL"
	APIStaticDateAnchorMonthMay       APIStaticDateAnchorMonth = "MAY"
	APIStaticDateAnchorMonthJune      APIStaticDateAnchorMonth = "JUNE"
	APIStaticDateAnchorMonthJuly      APIStaticDateAnchorMonth = "JULY"
	APIStaticDateAnchorMonthAugust    APIStaticDateAnchorMonth = "AUGUST"
	APIStaticDateAnchorMonthSeptember APIStaticDateAnchorMonth = "SEPTEMBER"
	APIStaticDateAnchorMonthOctober   APIStaticDateAnchorMonth = "OCTOBER"
	APIStaticDateAnchorMonthNovember  APIStaticDateAnchorMonth = "NOVEMBER"
	APIStaticDateAnchorMonthDecember  APIStaticDateAnchorMonth = "DECEMBER"
)

// The type of event anchor this is, can be: "CONTACT_PROPERTY_ANCHOR" or
// "STATIC_DATE_ANCHOR"
type APIStaticDateAnchorType string

const (
	APIStaticDateAnchorTypeStaticDateAnchor APIStaticDateAnchorType = "STATIC_DATE_ANCHOR"
)

// The properties DayOfMonth, Month, Type are required.
type APIStaticDateAnchorParam struct {
	// The day of the date to anchor on
	DayOfMonth int64 `json:"dayOfMonth,required"`
	// The month of the date to anchor on
	//
	// Any of "JANUARY", "FEBRUARY", "MARCH", "APRIL", "MAY", "JUNE", "JULY", "AUGUST",
	// "SEPTEMBER", "OCTOBER", "NOVEMBER", "DECEMBER".
	Month APIStaticDateAnchorMonth `json:"month,omitzero,required"`
	// The type of event anchor this is, can be: "CONTACT_PROPERTY_ANCHOR" or
	// "STATIC_DATE_ANCHOR"
	//
	// Any of "STATIC_DATE_ANCHOR".
	Type APIStaticDateAnchorType `json:"type,omitzero,required"`
	// The year of the date to anchor on. If this is not provided then this flow will
	// re-run each year.
	Year param.Opt[int64] `json:"year,omitzero"`
	paramObj
}

func (r APIStaticDateAnchorParam) MarshalJSON() (data []byte, err error) {
	type shadow APIStaticDateAnchorParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *APIStaticDateAnchorParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type APIStaticPropertyFilterDataSource struct {
	Name         string `json:"name,required"`
	PropertyName string `json:"propertyName,required"`
	StaticValue  string `json:"staticValue,required"`
	// Any of "STATIC_PROPERTY_FILTER".
	Type   APIStaticPropertyFilterDataSourceType `json:"type,required"`
	SortBy APISort                               `json:"sortBy"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name         respjson.Field
		PropertyName respjson.Field
		StaticValue  respjson.Field
		Type         respjson.Field
		SortBy       respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r APIStaticPropertyFilterDataSource) RawJSON() string { return r.JSON.raw }
func (r *APIStaticPropertyFilterDataSource) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this APIStaticPropertyFilterDataSource to a
// APIStaticPropertyFilterDataSourceParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// APIStaticPropertyFilterDataSourceParam.Overrides()
func (r APIStaticPropertyFilterDataSource) ToParam() APIStaticPropertyFilterDataSourceParam {
	return param.Override[APIStaticPropertyFilterDataSourceParam](json.RawMessage(r.RawJSON()))
}

type APIStaticPropertyFilterDataSourceType string

const (
	APIStaticPropertyFilterDataSourceTypeStaticPropertyFilter APIStaticPropertyFilterDataSourceType = "STATIC_PROPERTY_FILTER"
)

// The properties Name, PropertyName, StaticValue, Type are required.
type APIStaticPropertyFilterDataSourceParam struct {
	Name         string `json:"name,required"`
	PropertyName string `json:"propertyName,required"`
	StaticValue  string `json:"staticValue,required"`
	// Any of "STATIC_PROPERTY_FILTER".
	Type   APIStaticPropertyFilterDataSourceType `json:"type,omitzero,required"`
	SortBy APISortParam                          `json:"sortBy,omitzero"`
	paramObj
}

func (r APIStaticPropertyFilterDataSourceParam) MarshalJSON() (data []byte, err error) {
	type shadow APIStaticPropertyFilterDataSourceParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *APIStaticPropertyFilterDataSourceParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type APIStaticTimeZoneStrategy struct {
	TimeZoneID string `json:"timeZoneId,required"`
	// Any of "STATIC_TIME_ZONE".
	Type APIStaticTimeZoneStrategyType `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		TimeZoneID  respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r APIStaticTimeZoneStrategy) RawJSON() string { return r.JSON.raw }
func (r *APIStaticTimeZoneStrategy) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this APIStaticTimeZoneStrategy to a
// APIStaticTimeZoneStrategyParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// APIStaticTimeZoneStrategyParam.Overrides()
func (r APIStaticTimeZoneStrategy) ToParam() APIStaticTimeZoneStrategyParam {
	return param.Override[APIStaticTimeZoneStrategyParam](json.RawMessage(r.RawJSON()))
}

type APIStaticTimeZoneStrategyType string

const (
	APIStaticTimeZoneStrategyTypeStaticTimeZone APIStaticTimeZoneStrategyType = "STATIC_TIME_ZONE"
)

// The properties TimeZoneID, Type are required.
type APIStaticTimeZoneStrategyParam struct {
	TimeZoneID string `json:"timeZoneId,required"`
	// Any of "STATIC_TIME_ZONE".
	Type APIStaticTimeZoneStrategyType `json:"type,omitzero,required"`
	paramObj
}

func (r APIStaticTimeZoneStrategyParam) MarshalJSON() (data []byte, err error) {
	type shadow APIStaticTimeZoneStrategyParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *APIStaticTimeZoneStrategyParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type APIStaticValue struct {
	// A static value to use as the input
	StaticValue string `json:"staticValue,required"`
	// This is the type of input value. This can be one of: "FIELD_DATA",
	// "OBJECT_PROPERTY", "STATIC_VALUE", "RELATIVE_DATETIME", "TIMESTAMP",
	// "INCREMENT", "FETCHED_OBJECT_PROPERTY", "APPEND_OBJECT_PROPERTY",
	// "STATIC_APPEND_VALUE", "ENROLLMENT_EVENT_PROPERTY"
	//
	// Any of "STATIC_VALUE".
	Type APIStaticValueType `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		StaticValue respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r APIStaticValue) RawJSON() string { return r.JSON.raw }
func (r *APIStaticValue) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this APIStaticValue to a APIStaticValueParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// APIStaticValueParam.Overrides()
func (r APIStaticValue) ToParam() APIStaticValueParam {
	return param.Override[APIStaticValueParam](json.RawMessage(r.RawJSON()))
}

// This is the type of input value. This can be one of: "FIELD_DATA",
// "OBJECT_PROPERTY", "STATIC_VALUE", "RELATIVE_DATETIME", "TIMESTAMP",
// "INCREMENT", "FETCHED_OBJECT_PROPERTY", "APPEND_OBJECT_PROPERTY",
// "STATIC_APPEND_VALUE", "ENROLLMENT_EVENT_PROPERTY"
type APIStaticValueType string

const (
	APIStaticValueTypeStaticValue APIStaticValueType = "STATIC_VALUE"
)

// The properties StaticValue, Type are required.
type APIStaticValueParam struct {
	// A static value to use as the input
	StaticValue string `json:"staticValue,required"`
	// This is the type of input value. This can be one of: "FIELD_DATA",
	// "OBJECT_PROPERTY", "STATIC_VALUE", "RELATIVE_DATETIME", "TIMESTAMP",
	// "INCREMENT", "FETCHED_OBJECT_PROPERTY", "APPEND_OBJECT_PROPERTY",
	// "STATIC_APPEND_VALUE", "ENROLLMENT_EVENT_PROPERTY"
	//
	// Any of "STATIC_VALUE".
	Type APIStaticValueType `json:"type,omitzero,required"`
	paramObj
}

func (r APIStaticValueParam) MarshalJSON() (data []byte, err error) {
	type shadow APIStaticValueParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *APIStaticValueParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type APITimeDelay struct {
	// Any of "MONDAY", "TUESDAY", "WEDNESDAY", "THURSDAY", "FRIDAY", "SATURDAY",
	// "SUNDAY".
	DaysOfWeek []string `json:"daysOfWeek,required"`
	Delta      int64    `json:"delta,required"`
	// Any of "NANOS", "MICROS", "MILLIS", "SECONDS", "MINUTES", "HOURS", "HALF_DAYS",
	// "DAYS", "WEEKS", "MONTHS", "YEARS", "DECADES", "CENTURIES", "MILLENNIA", "ERAS",
	// "FOREVER".
	TimeUnit         APITimeDelayTimeUnit      `json:"timeUnit,required"`
	TimeOfDay        APITimeOfDay              `json:"timeOfDay"`
	TimeZoneStrategy APIStaticTimeZoneStrategy `json:"timeZoneStrategy"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DaysOfWeek       respjson.Field
		Delta            respjson.Field
		TimeUnit         respjson.Field
		TimeOfDay        respjson.Field
		TimeZoneStrategy respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r APITimeDelay) RawJSON() string { return r.JSON.raw }
func (r *APITimeDelay) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this APITimeDelay to a APITimeDelayParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// APITimeDelayParam.Overrides()
func (r APITimeDelay) ToParam() APITimeDelayParam {
	return param.Override[APITimeDelayParam](json.RawMessage(r.RawJSON()))
}

type APITimeDelayTimeUnit string

const (
	APITimeDelayTimeUnitNanos     APITimeDelayTimeUnit = "NANOS"
	APITimeDelayTimeUnitMicros    APITimeDelayTimeUnit = "MICROS"
	APITimeDelayTimeUnitMillis    APITimeDelayTimeUnit = "MILLIS"
	APITimeDelayTimeUnitSeconds   APITimeDelayTimeUnit = "SECONDS"
	APITimeDelayTimeUnitMinutes   APITimeDelayTimeUnit = "MINUTES"
	APITimeDelayTimeUnitHours     APITimeDelayTimeUnit = "HOURS"
	APITimeDelayTimeUnitHalfDays  APITimeDelayTimeUnit = "HALF_DAYS"
	APITimeDelayTimeUnitDays      APITimeDelayTimeUnit = "DAYS"
	APITimeDelayTimeUnitWeeks     APITimeDelayTimeUnit = "WEEKS"
	APITimeDelayTimeUnitMonths    APITimeDelayTimeUnit = "MONTHS"
	APITimeDelayTimeUnitYears     APITimeDelayTimeUnit = "YEARS"
	APITimeDelayTimeUnitDecades   APITimeDelayTimeUnit = "DECADES"
	APITimeDelayTimeUnitCenturies APITimeDelayTimeUnit = "CENTURIES"
	APITimeDelayTimeUnitMillennia APITimeDelayTimeUnit = "MILLENNIA"
	APITimeDelayTimeUnitEras      APITimeDelayTimeUnit = "ERAS"
	APITimeDelayTimeUnitForever   APITimeDelayTimeUnit = "FOREVER"
)

// The properties DaysOfWeek, Delta, TimeUnit are required.
type APITimeDelayParam struct {
	// Any of "MONDAY", "TUESDAY", "WEDNESDAY", "THURSDAY", "FRIDAY", "SATURDAY",
	// "SUNDAY".
	DaysOfWeek []string `json:"daysOfWeek,omitzero,required"`
	Delta      int64    `json:"delta,required"`
	// Any of "NANOS", "MICROS", "MILLIS", "SECONDS", "MINUTES", "HOURS", "HALF_DAYS",
	// "DAYS", "WEEKS", "MONTHS", "YEARS", "DECADES", "CENTURIES", "MILLENNIA", "ERAS",
	// "FOREVER".
	TimeUnit         APITimeDelayTimeUnit           `json:"timeUnit,omitzero,required"`
	TimeOfDay        APITimeOfDayParam              `json:"timeOfDay,omitzero"`
	TimeZoneStrategy APIStaticTimeZoneStrategyParam `json:"timeZoneStrategy,omitzero"`
	paramObj
}

func (r APITimeDelayParam) MarshalJSON() (data []byte, err error) {
	type shadow APITimeDelayParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *APITimeDelayParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type APITimeOfDay struct {
	Hour   int64 `json:"hour,required"`
	Minute int64 `json:"minute,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Hour        respjson.Field
		Minute      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r APITimeOfDay) RawJSON() string { return r.JSON.raw }
func (r *APITimeOfDay) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this APITimeOfDay to a APITimeOfDayParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// APITimeOfDayParam.Overrides()
func (r APITimeOfDay) ToParam() APITimeOfDayParam {
	return param.Override[APITimeOfDayParam](json.RawMessage(r.RawJSON()))
}

// The properties Hour, Minute are required.
type APITimeOfDayParam struct {
	Hour   int64 `json:"hour,required"`
	Minute int64 `json:"minute,required"`
	paramObj
}

func (r APITimeOfDayParam) MarshalJSON() (data []byte, err error) {
	type shadow APITimeOfDayParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *APITimeOfDayParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type APITimestampValue struct {
	// Currently only EXECUTION_TIME is supported.
	//
	// Any of "EXECUTION_TIME".
	TimestampType APITimestampValueTimestampType `json:"timestampType,required"`
	// This is the type of input value. This can be one of: "FIELD_DATA",
	// "OBJECT_PROPERTY", "STATIC_VALUE", "RELATIVE_DATETIME", "TIMESTAMP",
	// "INCREMENT", "FETCHED_OBJECT_PROPERTY", "APPEND_OBJECT_PROPERTY",
	// "STATIC_APPEND_VALUE", "ENROLLMENT_EVENT_PROPERTY"
	//
	// Any of "TIMESTAMP".
	Type APITimestampValueType `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		TimestampType respjson.Field
		Type          respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r APITimestampValue) RawJSON() string { return r.JSON.raw }
func (r *APITimestampValue) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this APITimestampValue to a APITimestampValueParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// APITimestampValueParam.Overrides()
func (r APITimestampValue) ToParam() APITimestampValueParam {
	return param.Override[APITimestampValueParam](json.RawMessage(r.RawJSON()))
}

// Currently only EXECUTION_TIME is supported.
type APITimestampValueTimestampType string

const (
	APITimestampValueTimestampTypeExecutionTime APITimestampValueTimestampType = "EXECUTION_TIME"
)

// This is the type of input value. This can be one of: "FIELD_DATA",
// "OBJECT_PROPERTY", "STATIC_VALUE", "RELATIVE_DATETIME", "TIMESTAMP",
// "INCREMENT", "FETCHED_OBJECT_PROPERTY", "APPEND_OBJECT_PROPERTY",
// "STATIC_APPEND_VALUE", "ENROLLMENT_EVENT_PROPERTY"
type APITimestampValueType string

const (
	APITimestampValueTypeTimestamp APITimestampValueType = "TIMESTAMP"
)

// The properties TimestampType, Type are required.
type APITimestampValueParam struct {
	// Currently only EXECUTION_TIME is supported.
	//
	// Any of "EXECUTION_TIME".
	TimestampType APITimestampValueTimestampType `json:"timestampType,omitzero,required"`
	// This is the type of input value. This can be one of: "FIELD_DATA",
	// "OBJECT_PROPERTY", "STATIC_VALUE", "RELATIVE_DATETIME", "TIMESTAMP",
	// "INCREMENT", "FETCHED_OBJECT_PROPERTY", "APPEND_OBJECT_PROPERTY",
	// "STATIC_APPEND_VALUE", "ENROLLMENT_EVENT_PROPERTY"
	//
	// Any of "TIMESTAMP".
	Type APITimestampValueType `json:"type,omitzero,required"`
	paramObj
}

func (r APITimestampValueParam) MarshalJSON() (data []byte, err error) {
	type shadow APITimestampValueParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *APITimestampValueParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type APITimeWindow struct {
	// Any of "MONDAY", "TUESDAY", "WEDNESDAY", "THURSDAY", "FRIDAY", "SATURDAY",
	// "SUNDAY".
	Day       APITimeWindowDay `json:"day,required"`
	EndTime   APITimeOfDay     `json:"endTime,required"`
	StartTime APITimeOfDay     `json:"startTime,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Day         respjson.Field
		EndTime     respjson.Field
		StartTime   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r APITimeWindow) RawJSON() string { return r.JSON.raw }
func (r *APITimeWindow) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this APITimeWindow to a APITimeWindowParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// APITimeWindowParam.Overrides()
func (r APITimeWindow) ToParam() APITimeWindowParam {
	return param.Override[APITimeWindowParam](json.RawMessage(r.RawJSON()))
}

type APITimeWindowDay string

const (
	APITimeWindowDayMonday    APITimeWindowDay = "MONDAY"
	APITimeWindowDayTuesday   APITimeWindowDay = "TUESDAY"
	APITimeWindowDayWednesday APITimeWindowDay = "WEDNESDAY"
	APITimeWindowDayThursday  APITimeWindowDay = "THURSDAY"
	APITimeWindowDayFriday    APITimeWindowDay = "FRIDAY"
	APITimeWindowDaySaturday  APITimeWindowDay = "SATURDAY"
	APITimeWindowDaySunday    APITimeWindowDay = "SUNDAY"
)

// The properties Day, EndTime, StartTime are required.
type APITimeWindowParam struct {
	// Any of "MONDAY", "TUESDAY", "WEDNESDAY", "THURSDAY", "FRIDAY", "SATURDAY",
	// "SUNDAY".
	Day       APITimeWindowDay  `json:"day,omitzero,required"`
	EndTime   APITimeOfDayParam `json:"endTime,omitzero,required"`
	StartTime APITimeOfDayParam `json:"startTime,omitzero,required"`
	paramObj
}

func (r APITimeWindowParam) MarshalJSON() (data []byte, err error) {
	type shadow APITimeWindowParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *APITimeWindowParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type APIUnEnrollmentSetting struct {
	// The IDs of the flows to unenroll an object in if it's enrolled in this flow.
	FlowIDs []string `json:"flowIds,required"`
	// The type of unenrollment to perform:
	//
	// "ALL" - unenroll the object from all other flows
	//
	// "SELECTIVE" - only unenroll the object from the flows specified in `flowIds`
	//
	// Any of "ALL", "SELECTIVE".
	Type APIUnEnrollmentSettingType `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FlowIDs     respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r APIUnEnrollmentSetting) RawJSON() string { return r.JSON.raw }
func (r *APIUnEnrollmentSetting) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this APIUnEnrollmentSetting to a APIUnEnrollmentSettingParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// APIUnEnrollmentSettingParam.Overrides()
func (r APIUnEnrollmentSetting) ToParam() APIUnEnrollmentSettingParam {
	return param.Override[APIUnEnrollmentSettingParam](json.RawMessage(r.RawJSON()))
}

// The type of unenrollment to perform:
//
// "ALL" - unenroll the object from all other flows
//
// "SELECTIVE" - only unenroll the object from the flows specified in `flowIds`
type APIUnEnrollmentSettingType string

const (
	APIUnEnrollmentSettingTypeAll       APIUnEnrollmentSettingType = "ALL"
	APIUnEnrollmentSettingTypeSelective APIUnEnrollmentSettingType = "SELECTIVE"
)

// The properties FlowIDs, Type are required.
type APIUnEnrollmentSettingParam struct {
	// The IDs of the flows to unenroll an object in if it's enrolled in this flow.
	FlowIDs []string `json:"flowIds,omitzero,required"`
	// The type of unenrollment to perform:
	//
	// "ALL" - unenroll the object from all other flows
	//
	// "SELECTIVE" - only unenroll the object from the flows specified in `flowIds`
	//
	// Any of "ALL", "SELECTIVE".
	Type APIUnEnrollmentSettingType `json:"type,omitzero,required"`
	paramObj
}

func (r APIUnEnrollmentSettingParam) MarshalJSON() (data []byte, err error) {
	type shadow APIUnEnrollmentSettingParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *APIUnEnrollmentSettingParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type APIWebhookAction struct {
	// The ID for this action.
	ActionID string `json:"actionId,required"`
	// The HTTP method to use when calling the webhook URL
	//
	// Any of "CONNECT", "DELETE", "GET", "HEAD", "OPTIONS", "PATCH", "POST", "PUT",
	// "TRACE".
	Method      APIWebhookActionMethod `json:"method,required"`
	QueryParams []APIInputVariable     `json:"queryParams,required"`
	// The type of action this is, can be: "STATIC_BRANCH", "LIST_BRANCH",
	// "AB_TEST_BRANCH", "CUSTOM_CODE", "WEBHOOK", or "SINGLE_CONNECTION"
	//
	// Any of "WEBHOOK".
	Type APIWebhookActionType `json:"type,required"`
	// The URL to call each time this action is executed.
	WebhookURL string `json:"webhookUrl,required"`
	// The type of auth to use when calling the webhook endpoint.
	AuthSettings APIWebhookActionAuthSettingsUnion `json:"authSettings"`
	Connection   APIConnection                     `json:"connection"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ActionID     respjson.Field
		Method       respjson.Field
		QueryParams  respjson.Field
		Type         respjson.Field
		WebhookURL   respjson.Field
		AuthSettings respjson.Field
		Connection   respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r APIWebhookAction) RawJSON() string { return r.JSON.raw }
func (r *APIWebhookAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this APIWebhookAction to a APIWebhookActionParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// APIWebhookActionParam.Overrides()
func (r APIWebhookAction) ToParam() APIWebhookActionParam {
	return param.Override[APIWebhookActionParam](json.RawMessage(r.RawJSON()))
}

// The HTTP method to use when calling the webhook URL
type APIWebhookActionMethod string

const (
	APIWebhookActionMethodConnect APIWebhookActionMethod = "CONNECT"
	APIWebhookActionMethodDelete  APIWebhookActionMethod = "DELETE"
	APIWebhookActionMethodGet     APIWebhookActionMethod = "GET"
	APIWebhookActionMethodHead    APIWebhookActionMethod = "HEAD"
	APIWebhookActionMethodOptions APIWebhookActionMethod = "OPTIONS"
	APIWebhookActionMethodPatch   APIWebhookActionMethod = "PATCH"
	APIWebhookActionMethodPost    APIWebhookActionMethod = "POST"
	APIWebhookActionMethodPut     APIWebhookActionMethod = "PUT"
	APIWebhookActionMethodTrace   APIWebhookActionMethod = "TRACE"
)

// The type of action this is, can be: "STATIC_BRANCH", "LIST_BRANCH",
// "AB_TEST_BRANCH", "CUSTOM_CODE", "WEBHOOK", or "SINGLE_CONNECTION"
type APIWebhookActionType string

const (
	APIWebhookActionTypeWebhook APIWebhookActionType = "WEBHOOK"
)

// APIWebhookActionAuthSettingsUnion contains all possible properties and values
// from [APIAuthKeyWebhookAuthSettings], [APISignatureWebhookAuthSettings].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type APIWebhookActionAuthSettingsUnion struct {
	// This field is from variant [APIAuthKeyWebhookAuthSettings].
	Location APIAuthKeyWebhookAuthSettingsLocation `json:"location"`
	// This field is from variant [APIAuthKeyWebhookAuthSettings].
	Name string `json:"name"`
	// This field is from variant [APIAuthKeyWebhookAuthSettings].
	SecretName string `json:"secretName"`
	Type       string `json:"type"`
	// This field is from variant [APISignatureWebhookAuthSettings].
	AppID int64 `json:"appId"`
	JSON  struct {
		Location   respjson.Field
		Name       respjson.Field
		SecretName respjson.Field
		Type       respjson.Field
		AppID      respjson.Field
		raw        string
	} `json:"-"`
}

func (u APIWebhookActionAuthSettingsUnion) AsAuthKey() (v APIAuthKeyWebhookAuthSettings) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u APIWebhookActionAuthSettingsUnion) AsSignature() (v APISignatureWebhookAuthSettings) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u APIWebhookActionAuthSettingsUnion) RawJSON() string { return u.JSON.raw }

func (r *APIWebhookActionAuthSettingsUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties ActionID, Method, QueryParams, Type, WebhookURL are required.
type APIWebhookActionParam struct {
	// The ID for this action.
	ActionID string `json:"actionId,required"`
	// The HTTP method to use when calling the webhook URL
	//
	// Any of "CONNECT", "DELETE", "GET", "HEAD", "OPTIONS", "PATCH", "POST", "PUT",
	// "TRACE".
	Method      APIWebhookActionMethod  `json:"method,omitzero,required"`
	QueryParams []APIInputVariableParam `json:"queryParams,omitzero,required"`
	// The type of action this is, can be: "STATIC_BRANCH", "LIST_BRANCH",
	// "AB_TEST_BRANCH", "CUSTOM_CODE", "WEBHOOK", or "SINGLE_CONNECTION"
	//
	// Any of "WEBHOOK".
	Type APIWebhookActionType `json:"type,omitzero,required"`
	// The URL to call each time this action is executed.
	WebhookURL string `json:"webhookUrl,required"`
	// The type of auth to use when calling the webhook endpoint.
	AuthSettings APIWebhookActionAuthSettingsUnionParam `json:"authSettings,omitzero"`
	Connection   APIConnectionParam                     `json:"connection,omitzero"`
	paramObj
}

func (r APIWebhookActionParam) MarshalJSON() (data []byte, err error) {
	type shadow APIWebhookActionParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *APIWebhookActionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type APIWebhookActionAuthSettingsUnionParam struct {
	OfAuthKey   *APIAuthKeyWebhookAuthSettingsParam   `json:",omitzero,inline"`
	OfSignature *APISignatureWebhookAuthSettingsParam `json:",omitzero,inline"`
	paramUnion
}

func (u APIWebhookActionAuthSettingsUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfAuthKey, u.OfSignature)
}
func (u *APIWebhookActionAuthSettingsUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *APIWebhookActionAuthSettingsUnionParam) asAny() any {
	if !param.IsOmitted(u.OfAuthKey) {
		return u.OfAuthKey
	} else if !param.IsOmitted(u.OfSignature) {
		return u.OfSignature
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIWebhookActionAuthSettingsUnionParam) GetLocation() *string {
	if vt := u.OfAuthKey; vt != nil {
		return (*string)(&vt.Location)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIWebhookActionAuthSettingsUnionParam) GetName() *string {
	if vt := u.OfAuthKey; vt != nil {
		return &vt.Name
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIWebhookActionAuthSettingsUnionParam) GetSecretName() *string {
	if vt := u.OfAuthKey; vt != nil {
		return &vt.SecretName
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIWebhookActionAuthSettingsUnionParam) GetAppID() *int64 {
	if vt := u.OfSignature; vt != nil {
		return &vt.AppID
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u APIWebhookActionAuthSettingsUnionParam) GetType() *string {
	if vt := u.OfAuthKey; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfSignature; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

type APIWeeklyEnrollmentSchedule struct {
	// Which days of the week to allow enrollments.
	//
	// Any of "MONDAY", "TUESDAY", "WEDNESDAY", "THURSDAY", "FRIDAY", "SATURDAY",
	// "SUNDAY".
	DaysOfWeek []string     `json:"daysOfWeek,required"`
	TimeOfDay  APITimeOfDay `json:"timeOfDay,required"`
	// The type of enrollment schedule this is, can be: "DAILY", "WEEKLY",
	// "MONTHLY_SPECIFIC_DAYS", "MONTHLY_RELATIVE_DAYS", "YEARLY"
	//
	// Any of "WEEKLY".
	Type APIWeeklyEnrollmentScheduleType `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DaysOfWeek  respjson.Field
		TimeOfDay   respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r APIWeeklyEnrollmentSchedule) RawJSON() string { return r.JSON.raw }
func (r *APIWeeklyEnrollmentSchedule) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this APIWeeklyEnrollmentSchedule to a
// APIWeeklyEnrollmentScheduleParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// APIWeeklyEnrollmentScheduleParam.Overrides()
func (r APIWeeklyEnrollmentSchedule) ToParam() APIWeeklyEnrollmentScheduleParam {
	return param.Override[APIWeeklyEnrollmentScheduleParam](json.RawMessage(r.RawJSON()))
}

// The type of enrollment schedule this is, can be: "DAILY", "WEEKLY",
// "MONTHLY_SPECIFIC_DAYS", "MONTHLY_RELATIVE_DAYS", "YEARLY"
type APIWeeklyEnrollmentScheduleType string

const (
	APIWeeklyEnrollmentScheduleTypeWeekly APIWeeklyEnrollmentScheduleType = "WEEKLY"
)

// The properties DaysOfWeek, TimeOfDay, Type are required.
type APIWeeklyEnrollmentScheduleParam struct {
	// Which days of the week to allow enrollments.
	//
	// Any of "MONDAY", "TUESDAY", "WEDNESDAY", "THURSDAY", "FRIDAY", "SATURDAY",
	// "SUNDAY".
	DaysOfWeek []string          `json:"daysOfWeek,omitzero,required"`
	TimeOfDay  APITimeOfDayParam `json:"timeOfDay,omitzero,required"`
	// The type of enrollment schedule this is, can be: "DAILY", "WEEKLY",
	// "MONTHLY_SPECIFIC_DAYS", "MONTHLY_RELATIVE_DAYS", "YEARLY"
	//
	// Any of "WEEKLY".
	Type APIWeeklyEnrollmentScheduleType `json:"type,omitzero,required"`
	paramObj
}

func (r APIWeeklyEnrollmentScheduleParam) MarshalJSON() (data []byte, err error) {
	type shadow APIWeeklyEnrollmentScheduleParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *APIWeeklyEnrollmentScheduleParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type APIYearlyEnrollmentSchedule struct {
	// The day of the date each year to run this flow.
	DayOfMonth int64 `json:"dayOfMonth,required"`
	// The month of the date each year to run this flow.
	//
	// Any of "JANUARY", "FEBRUARY", "MARCH", "APRIL", "MAY", "JUNE", "JULY", "AUGUST",
	// "SEPTEMBER", "OCTOBER", "NOVEMBER", "DECEMBER".
	Month     APIYearlyEnrollmentScheduleMonth `json:"month,required"`
	TimeOfDay APITimeOfDay                     `json:"timeOfDay,required"`
	// The type of enrollment schedule this is, can be: "DAILY", "WEEKLY",
	// "MONTHLY_SPECIFIC_DAYS", "MONTHLY_RELATIVE_DAYS", "YEARLY"
	//
	// Any of "YEARLY".
	Type APIYearlyEnrollmentScheduleType `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DayOfMonth  respjson.Field
		Month       respjson.Field
		TimeOfDay   respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r APIYearlyEnrollmentSchedule) RawJSON() string { return r.JSON.raw }
func (r *APIYearlyEnrollmentSchedule) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this APIYearlyEnrollmentSchedule to a
// APIYearlyEnrollmentScheduleParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// APIYearlyEnrollmentScheduleParam.Overrides()
func (r APIYearlyEnrollmentSchedule) ToParam() APIYearlyEnrollmentScheduleParam {
	return param.Override[APIYearlyEnrollmentScheduleParam](json.RawMessage(r.RawJSON()))
}

// The month of the date each year to run this flow.
type APIYearlyEnrollmentScheduleMonth string

const (
	APIYearlyEnrollmentScheduleMonthJanuary   APIYearlyEnrollmentScheduleMonth = "JANUARY"
	APIYearlyEnrollmentScheduleMonthFebruary  APIYearlyEnrollmentScheduleMonth = "FEBRUARY"
	APIYearlyEnrollmentScheduleMonthMarch     APIYearlyEnrollmentScheduleMonth = "MARCH"
	APIYearlyEnrollmentScheduleMonthApril     APIYearlyEnrollmentScheduleMonth = "APRIL"
	APIYearlyEnrollmentScheduleMonthMay       APIYearlyEnrollmentScheduleMonth = "MAY"
	APIYearlyEnrollmentScheduleMonthJune      APIYearlyEnrollmentScheduleMonth = "JUNE"
	APIYearlyEnrollmentScheduleMonthJuly      APIYearlyEnrollmentScheduleMonth = "JULY"
	APIYearlyEnrollmentScheduleMonthAugust    APIYearlyEnrollmentScheduleMonth = "AUGUST"
	APIYearlyEnrollmentScheduleMonthSeptember APIYearlyEnrollmentScheduleMonth = "SEPTEMBER"
	APIYearlyEnrollmentScheduleMonthOctober   APIYearlyEnrollmentScheduleMonth = "OCTOBER"
	APIYearlyEnrollmentScheduleMonthNovember  APIYearlyEnrollmentScheduleMonth = "NOVEMBER"
	APIYearlyEnrollmentScheduleMonthDecember  APIYearlyEnrollmentScheduleMonth = "DECEMBER"
)

// The type of enrollment schedule this is, can be: "DAILY", "WEEKLY",
// "MONTHLY_SPECIFIC_DAYS", "MONTHLY_RELATIVE_DAYS", "YEARLY"
type APIYearlyEnrollmentScheduleType string

const (
	APIYearlyEnrollmentScheduleTypeYearly APIYearlyEnrollmentScheduleType = "YEARLY"
)

// The properties DayOfMonth, Month, TimeOfDay, Type are required.
type APIYearlyEnrollmentScheduleParam struct {
	// The day of the date each year to run this flow.
	DayOfMonth int64 `json:"dayOfMonth,required"`
	// The month of the date each year to run this flow.
	//
	// Any of "JANUARY", "FEBRUARY", "MARCH", "APRIL", "MAY", "JUNE", "JULY", "AUGUST",
	// "SEPTEMBER", "OCTOBER", "NOVEMBER", "DECEMBER".
	Month     APIYearlyEnrollmentScheduleMonth `json:"month,omitzero,required"`
	TimeOfDay APITimeOfDayParam                `json:"timeOfDay,omitzero,required"`
	// The type of enrollment schedule this is, can be: "DAILY", "WEEKLY",
	// "MONTHLY_SPECIFIC_DAYS", "MONTHLY_RELATIVE_DAYS", "YEARLY"
	//
	// Any of "YEARLY".
	Type APIYearlyEnrollmentScheduleType `json:"type,omitzero,required"`
	paramObj
}

func (r APIYearlyEnrollmentScheduleParam) MarshalJSON() (data []byte, err error) {
	type shadow APIYearlyEnrollmentScheduleParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *APIYearlyEnrollmentScheduleParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BatchResponseAPIFlow struct {
	CompletedAt time.Time      `json:"completedAt,required" format:"date-time"`
	Results     []APIFlowUnion `json:"results,required"`
	StartedAt   time.Time      `json:"startedAt,required" format:"date-time"`
	// Any of "PENDING", "PROCESSING", "CANCELED", "COMPLETE".
	Status      BatchResponseAPIFlowStatus `json:"status,required"`
	Links       map[string]string          `json:"links"`
	RequestedAt time.Time                  `json:"requestedAt" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CompletedAt respjson.Field
		Results     respjson.Field
		StartedAt   respjson.Field
		Status      respjson.Field
		Links       respjson.Field
		RequestedAt respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BatchResponseAPIFlow) RawJSON() string { return r.JSON.raw }
func (r *BatchResponseAPIFlow) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BatchResponseAPIFlowStatus string

const (
	BatchResponseAPIFlowStatusPending    BatchResponseAPIFlowStatus = "PENDING"
	BatchResponseAPIFlowStatusProcessing BatchResponseAPIFlowStatus = "PROCESSING"
	BatchResponseAPIFlowStatusCanceled   BatchResponseAPIFlowStatus = "CANCELED"
	BatchResponseAPIFlowStatusComplete   BatchResponseAPIFlowStatus = "COMPLETE"
)

type BatchResponseFlowIDWorkflowIDMappingResponse struct {
	CompletedAt time.Time                         `json:"completedAt,required" format:"date-time"`
	Results     []FlowIDWorkflowIDMappingResponse `json:"results,required"`
	StartedAt   time.Time                         `json:"startedAt,required" format:"date-time"`
	// Any of "PENDING", "PROCESSING", "CANCELED", "COMPLETE".
	Status      BatchResponseFlowIDWorkflowIDMappingResponseStatus `json:"status,required"`
	Links       map[string]string                                  `json:"links"`
	RequestedAt time.Time                                          `json:"requestedAt" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CompletedAt respjson.Field
		Results     respjson.Field
		StartedAt   respjson.Field
		Status      respjson.Field
		Links       respjson.Field
		RequestedAt respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BatchResponseFlowIDWorkflowIDMappingResponse) RawJSON() string { return r.JSON.raw }
func (r *BatchResponseFlowIDWorkflowIDMappingResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BatchResponseFlowIDWorkflowIDMappingResponseStatus string

const (
	BatchResponseFlowIDWorkflowIDMappingResponseStatusPending    BatchResponseFlowIDWorkflowIDMappingResponseStatus = "PENDING"
	BatchResponseFlowIDWorkflowIDMappingResponseStatusProcessing BatchResponseFlowIDWorkflowIDMappingResponseStatus = "PROCESSING"
	BatchResponseFlowIDWorkflowIDMappingResponseStatusCanceled   BatchResponseFlowIDWorkflowIDMappingResponseStatus = "CANCELED"
	BatchResponseFlowIDWorkflowIDMappingResponseStatusComplete   BatchResponseFlowIDWorkflowIDMappingResponseStatus = "COMPLETE"
)

type CollectionResponseAPIFlowEmailCampaign struct {
	Results []APIFlowEmailCampaign `json:"results,required"`
	// Contains information pagination of results.
	Paging marketing.Paging `json:"paging"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Results     respjson.Field
		Paging      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CollectionResponseAPIFlowEmailCampaign) RawJSON() string { return r.JSON.raw }
func (r *CollectionResponseAPIFlowEmailCampaign) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CollectionResponseAPIFlowListingForwardPaging struct {
	Results []APIFlowListing     `json:"results,required"`
	Paging  shared.ForwardPaging `json:"paging"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Results     respjson.Field
		Paging      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CollectionResponseAPIFlowListingForwardPaging) RawJSON() string { return r.JSON.raw }
func (r *CollectionResponseAPIFlowListingForwardPaging) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FlowIDWorkflowIDMappingResponse struct {
	FlowID     int64 `json:"flowId,required"`
	WorkflowID int64 `json:"workflowId,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FlowID      respjson.Field
		WorkflowID  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FlowIDWorkflowIDMappingResponse) RawJSON() string { return r.JSON.raw }
func (r *FlowIDWorkflowIDMappingResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkflowNewParams struct {
	APIFlowCreateRequest APIFlowCreateRequestUnionParam
	paramObj
}

func (r WorkflowNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.APIFlowCreateRequest)
}
func (r *WorkflowNewParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.APIFlowCreateRequest)
}

type WorkflowUpdateParams struct {
	APIFlowPutRequest APIFlowPutRequestUnionParam
	paramObj
}

func (r WorkflowUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.APIFlowPutRequest)
}
func (r *WorkflowUpdateParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.APIFlowPutRequest)
}

type WorkflowListParams struct {
	// The paging cursor token of the last successfully read resource will be returned
	// as the `paging.next.after` JSON property of a paged response containing more
	// results.
	After param.Opt[string] `query:"after,omitzero" json:"-"`
	// The maximum number of results to display per page.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [WorkflowListParams]'s query parameters as `url.Values`.
func (r WorkflowListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type WorkflowBatchGetParams struct {
	APIFlowBatchInput APIFlowBatchInputParam
	paramObj
}

func (r WorkflowBatchGetParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.APIFlowBatchInput)
}
func (r *WorkflowBatchGetParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.APIFlowBatchInput)
}

type WorkflowBatchGetIDMappingsParams struct {
	APIFlowBatchMigrationInput APIFlowBatchMigrationInputParam
	paramObj
}

func (r WorkflowBatchGetIDMappingsParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.APIFlowBatchMigrationInput)
}
func (r *WorkflowBatchGetIDMappingsParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.APIFlowBatchMigrationInput)
}

type WorkflowListEmailCampaignsParams struct {
	// The paging cursor token of the last successfully read resource will be returned
	// as the `paging.next.after` JSON property of a paged response containing more
	// results.
	After  param.Opt[string] `query:"after,omitzero" json:"-"`
	Before param.Opt[string] `query:"before,omitzero" json:"-"`
	// The maximum number of results to display per page.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// The ID of the workflow.
	FlowID []string `query:"flowId,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [WorkflowListEmailCampaignsParams]'s query parameters as
// `url.Values`.
func (r WorkflowListEmailCampaignsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
