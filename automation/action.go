// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package automation

import (
	"encoding/json"
	"time"

	"github.com/stainless-sdks/hubspot-sdk-go/internal/apijson"
	"github.com/stainless-sdks/hubspot-sdk-go/option"
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
	options     []option.RequestOption
	Callbacks   ActionCallbackService
	Definitions ActionDefinitionService
	Functions   ActionFunctionService
	Revisions   ActionRevisionService
}

// NewActionService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewActionService(opts ...option.RequestOption) (r ActionService) {
	r = ActionService{}
	r.options = opts
	r.Callbacks = NewActionCallbackService(opts...)
	r.Definitions = NewActionDefinitionService(opts...)
	r.Functions = NewActionFunctionService(opts...)
	r.Revisions = NewActionRevisionService(opts...)
	return
}

// The properties ActionExecutionIndex, EnrollmentID are required.
type ActionExecutionIndexIdentifierParam struct {
	// The index number representing the execution order of the action.
	ActionExecutionIndex int64 `json:"actionExecutionIndex" api:"required"`
	// The ID associated with the enrollment process.
	EnrollmentID int64 `json:"enrollmentId" api:"required"`
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
	// The unique identifier for the agent making the request.
	AgentID              int64                     `json:"agentId" api:"required"`
	ChirpAIContextObject ChirpAIContextObjectParam `json:"chirpAiContextObject,omitzero" api:"required"`
	// Indicates the source of the request, with the default value being 'AGENTS'.
	//
	// Any of "AGENTS".
	Source AgentRequestContextSource `json:"source,omitzero" api:"required"`
	// The unique identifier for the trajectory associated with the agent request.
	TrajectoryID param.Opt[string] `json:"trajectoryId,omitzero"`
	paramObj
}

func (r AgentRequestContextParam) MarshalJSON() (data []byte, err error) {
	type shadow AgentRequestContextParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AgentRequestContextParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Indicates the source of the request, with the default value being 'AGENTS'.
type AgentRequestContextSource string

const (
	AgentRequestContextSourceAgents AgentRequestContextSource = "AGENTS"
)

type ArrayFieldSchema struct {
	Items any `json:"items" api:"required"`
	// Specifies that the field is of type 'ARRAY'.
	//
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

// Specifies that the field is of type 'ARRAY'.
type ArrayFieldSchemaType string

const (
	ArrayFieldSchemaTypeArray ArrayFieldSchemaType = "ARRAY"
)

// The properties Items, Type are required.
type ArrayFieldSchemaParam struct {
	Items any `json:"items,omitzero" api:"required"`
	// Specifies that the field is of type 'ARRAY'.
	//
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
	// Specifies the field type as BOOLEAN, indicating that the field can hold a true
	// or false value.
	//
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

// Specifies the field type as BOOLEAN, indicating that the field can hold a true
// or false value.
type BooleanFieldSchemaType string

const (
	BooleanFieldSchemaTypeBoolean BooleanFieldSchemaType = "BOOLEAN"
)

// The property Type is required.
type BooleanFieldSchemaParam struct {
	// Specifies the field type as BOOLEAN, indicating that the field can hold a true
	// or false value.
	//
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
	// The unique identifier for the callback.
	CallbackID string `json:"callbackId" api:"required"`
	// Holds the output fields for the callback completion.
	OutputFields map[string]string `json:"outputFields,omitzero" api:"required"`
	// Contains the typed outputs for the callback completion.
	TypedOutputs any `json:"typedOutputs,omitzero" api:"required"`
	// Specifies the type of failure reason for the callback completion.
	FailureReasonType param.Opt[string] `json:"failureReasonType,omitzero"`
	// Defines the context of the request, which can be one of several predefined
	// types.
	RequestContext CallbackCompletionBatchRequestRequestContextUnionParam `json:"requestContext,omitzero"`
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

// The properties OutputFields, TypedOutputs are required.
type CallbackCompletionRequestParam struct {
	// Contains the output fields associated with the callback, with each field
	// represented as a key-value pair.
	OutputFields map[string]string `json:"outputFields,omitzero" api:"required"`
	// Holds the typed outputs related to the callback, structured as an object.
	TypedOutputs any `json:"typedOutputs,omitzero" api:"required"`
	// Indicates the reason for the failure of a callback completion.
	FailureReasonType param.Opt[string] `json:"failureReasonType,omitzero"`
	// Specifies the context in which the request is made, which can be one of several
	// predefined contexts.
	RequestContext CallbackCompletionRequestRequestContextUnionParam `json:"requestContext,omitzero"`
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

// The properties ApplicationGroup, ApplicationID, Metadata, OtelContextHolder,
// UnstructuredSources are required.
type ChirpAIContextObjectParam struct {
	// The group to which the application belongs.
	ApplicationGroup string `json:"applicationGroup" api:"required"`
	// The identifier for the application associated with the context.
	ApplicationID string `json:"applicationId" api:"required"`
	// Additional metadata related to the context, represented as key-value pairs.
	Metadata map[string]string `json:"metadata,omitzero" api:"required"`
	// Holds OpenTelemetry context information as key-value pairs.
	OtelContextHolder map[string]string `json:"otelContextHolder,omitzero" api:"required"`
	// Any of "NONE", "USER_INPUT", "LOGGED_EMAIL", "VIDEO_CALL", "AUDIO_CALL",
	// "CALL_TRANSCRIPT", "MEETING_TRANSCRIPT", "FORMS", "FEEDBACK_SURVEY", "PDF",
	// "QUOTE", "INVOICE", "OTHER_ATTACHMENT_DOC", "WHATSAPP", "SMS", "CHAT",
	// "FACEBOOK_MESSENGER", "CUSTOM_CHANNEL_OR_API", "MANY", "NOTE", "DERIVED".
	UnstructuredSources []string          `json:"unstructuredSources,omitzero" api:"required"`
	ConversationID      param.Opt[string] `json:"conversationId,omitzero"`
	// The identifier for the feature associated with the context.
	FeatureID param.Opt[string] `json:"featureId,omitzero"`
	// The identifier for the inference associated with the context.
	InferenceID param.Opt[string] `json:"inferenceId,omitzero"`
	// The identifier for the trajectory, formatted as a UUID.
	TrajectoryID  param.Opt[string]  `json:"trajectoryId,omitzero" format:"uuid"`
	ComplianceIDs ComplianceIDsParam `json:"complianceIds,omitzero"`
	paramObj
}

func (r ChirpAIContextObjectParam) MarshalJSON() (data []byte, err error) {
	type shadow ChirpAIContextObjectParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ChirpAIContextObjectParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CollectionResponsePublicActionDefinitionForwardPaging struct {
	Results []PublicActionDefinition `json:"results" api:"required"`
	Paging  shared.ForwardPaging     `json:"paging"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Results     respjson.Field
		Paging      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CollectionResponsePublicActionDefinitionForwardPaging) RawJSON() string { return r.JSON.raw }
func (r *CollectionResponsePublicActionDefinitionForwardPaging) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CollectionResponsePublicActionFunctionIdentifierNoPaging struct {
	Results []PublicActionFunctionIdentifier `json:"results" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Results     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CollectionResponsePublicActionFunctionIdentifierNoPaging) RawJSON() string { return r.JSON.raw }
func (r *CollectionResponsePublicActionFunctionIdentifierNoPaging) UnmarshalJSON(data []byte) error {
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
	ContactIDs []ContactIDParam `json:"contactIds,omitzero" api:"required"`
	PortalIDs  []int64          `json:"portalIds,omitzero" api:"required"`
	UserIDs    []int64          `json:"userIds,omitzero" api:"required"`
	// The reason why no contact ID is available.
	NoContactIDReason param.Opt[string] `json:"noContactIdReason,omitzero"`
	// The reason why no portal ID is available.
	NoPortalIDReason param.Opt[string] `json:"noPortalIdReason,omitzero"`
	// The reason why no user ID is available.
	NoUserIDReason param.Opt[string] `json:"noUserIdReason,omitzero"`
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
	// The ID of the portal associated with the contact.
	PortalID int64 `json:"portalId" api:"required"`
	// The email address of the contact.
	Email param.Opt[string] `json:"email,omitzero"`
	// The unique identifier for the contact.
	Vid param.Opt[int64] `json:"vid,omitzero"`
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
	// Indicates the source of the request, with the default value being 'COPILOT'.
	//
	// Any of "COPILOT".
	Source CopilotRequestContextSource `json:"source,omitzero" api:"required"`
	// The unique identifier for the trajectory.
	TrajectoryID param.Opt[string] `json:"trajectoryId,omitzero"`
	paramObj
}

func (r CopilotRequestContextParam) MarshalJSON() (data []byte, err error) {
	type shadow CopilotRequestContextParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CopilotRequestContextParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Indicates the source of the request, with the default value being 'COPILOT'.
type CopilotRequestContextSource string

const (
	CopilotRequestContextSourceCopilot CopilotRequestContextSource = "COPILOT"
)

type DoubleFieldSchema struct {
	// Indicates the field type as DOUBLE.
	//
	// Any of "DOUBLE".
	Type DoubleFieldSchemaType `json:"type" api:"required"`
	// The maximum allowable value for the double field.
	Maximum float64 `json:"maximum"`
	// The minimum allowable value for the double field.
	Minimum float64 `json:"minimum"`
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

// Indicates the field type as DOUBLE.
type DoubleFieldSchemaType string

const (
	DoubleFieldSchemaTypeDouble DoubleFieldSchemaType = "DOUBLE"
)

// The property Type is required.
type DoubleFieldSchemaParam struct {
	// Indicates the field type as DOUBLE.
	//
	// Any of "DOUBLE".
	Type DoubleFieldSchemaType `json:"type,omitzero" api:"required"`
	// The maximum allowable value for the double field.
	Maximum param.Opt[float64] `json:"maximum,omitzero"`
	// The minimum allowable value for the double field.
	Minimum param.Opt[float64] `json:"minimum,omitzero"`
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
	// Indicates whether the field's options are sourced externally.
	ExternalOptions bool `json:"externalOptions" api:"required"`
	// The unique identifier for the field.
	Name    string                           `json:"name" api:"required"`
	Options []shared.AutomationActionsOption `json:"options" api:"required"`
	// Defines the structure and constraints of the field.
	Schema FieldTypeDefinitionSchemaUnion `json:"schema" api:"required"`
	// Specifies the data type of the field, with accepted values like bool, date,
	// datetime, enumeration, json, number, object_coordinates, phone_number, string.
	//
	// Any of "bool", "currency_number", "date", "datetime", "enumeration", "json",
	// "number", "object_coordinates", "phone_number", "string".
	Type FieldTypeDefinitionType `json:"type" api:"required"`
	// Specifies whether the field uses the Chirp feature.
	UseChirp bool `json:"useChirp" api:"required"`
	// A detailed explanation of the field's purpose and usage.
	Description string `json:"description"`
	// Specifies the type of external reference for options.
	ExternalOptionsReferenceType string `json:"externalOptionsReferenceType"`
	// Describes the field's type in the UI, with accepted values like booleancheckbox,
	// calculation_equation, checkbox, date, file, html, number, phonenumber, radio,
	// select, text, textarea, unknown.
	//
	// Any of "booleancheckbox", "calculation_equation", "calculation_read_time",
	// "calculation_rollup", "calculation_score", "checkbox", "date", "file", "html",
	// "number", "phonenumber", "radio", "select", "text", "textarea", "unknown".
	FieldType FieldTypeDefinitionFieldType `json:"fieldType"`
	// Additional information or guidance about the field.
	HelpText string `json:"helpText"`
	// The user-friendly label for the field.
	Label string `json:"label"`
	// A URL that provides options for the field.
	OptionsURL string `json:"optionsUrl"`
	// Indicates the type of object that the field references, with accepted values
	// like OWNER.
	//
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
	Items any `json:"items"`
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

// Specifies the data type of the field, with accepted values like bool, date,
// datetime, enumeration, json, number, object_coordinates, phone_number, string.
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

// Describes the field's type in the UI, with accepted values like booleancheckbox,
// calculation_equation, checkbox, date, file, html, number, phonenumber, radio,
// select, text, textarea, unknown.
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

// Indicates the type of object that the field references, with accepted values
// like OWNER.
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
	// Indicates whether the field's options are sourced externally.
	ExternalOptions bool `json:"externalOptions" api:"required"`
	// The unique identifier for the field.
	Name    string                                `json:"name" api:"required"`
	Options []shared.AutomationActionsOptionParam `json:"options,omitzero" api:"required"`
	// Defines the structure and constraints of the field.
	Schema FieldTypeDefinitionSchemaUnionParam `json:"schema,omitzero" api:"required"`
	// Specifies the data type of the field, with accepted values like bool, date,
	// datetime, enumeration, json, number, object_coordinates, phone_number, string.
	//
	// Any of "bool", "currency_number", "date", "datetime", "enumeration", "json",
	// "number", "object_coordinates", "phone_number", "string".
	Type FieldTypeDefinitionType `json:"type,omitzero" api:"required"`
	// Specifies whether the field uses the Chirp feature.
	UseChirp bool `json:"useChirp" api:"required"`
	// A detailed explanation of the field's purpose and usage.
	Description param.Opt[string] `json:"description,omitzero"`
	// Specifies the type of external reference for options.
	ExternalOptionsReferenceType param.Opt[string] `json:"externalOptionsReferenceType,omitzero"`
	// Additional information or guidance about the field.
	HelpText param.Opt[string] `json:"helpText,omitzero"`
	// The user-friendly label for the field.
	Label param.Opt[string] `json:"label,omitzero"`
	// A URL that provides options for the field.
	OptionsURL param.Opt[string] `json:"optionsUrl,omitzero"`
	// Describes the field's type in the UI, with accepted values like booleancheckbox,
	// calculation_equation, checkbox, date, file, html, number, phonenumber, radio,
	// select, text, textarea, unknown.
	//
	// Any of "booleancheckbox", "calculation_equation", "calculation_read_time",
	// "calculation_rollup", "calculation_score", "checkbox", "date", "file", "html",
	// "number", "phonenumber", "radio", "select", "text", "textarea", "unknown".
	FieldType FieldTypeDefinitionFieldType `json:"fieldType,omitzero"`
	// Indicates the type of object that the field references, with accepted values
	// like OWNER.
	//
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

type IntegerFieldSchema struct {
	// The type of the field, which is set to INTEGER.
	//
	// Any of "INTEGER".
	Type IntegerFieldSchemaType `json:"type" api:"required"`
	// The maximum value allowed for the integer field.
	Maximum int64 `json:"maximum"`
	// The minimum value allowed for the integer field.
	Minimum int64 `json:"minimum"`
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

// The type of the field, which is set to INTEGER.
type IntegerFieldSchemaType string

const (
	IntegerFieldSchemaTypeInteger IntegerFieldSchemaType = "INTEGER"
)

// The property Type is required.
type IntegerFieldSchemaParam struct {
	// The type of the field, which is set to INTEGER.
	//
	// Any of "INTEGER".
	Type IntegerFieldSchemaType `json:"type,omitzero" api:"required"`
	// The maximum value allowed for the integer field.
	Maximum param.Opt[int64] `json:"maximum,omitzero"`
	// The minimum value allowed for the integer field.
	Minimum param.Opt[int64] `json:"minimum,omitzero"`
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
	// The type of the field, which is LONG by default.
	//
	// Any of "LONG".
	Type LongFieldSchemaType `json:"type" api:"required"`
	// The maximum value allowed for the long field.
	Maximum int64 `json:"maximum"`
	// The minimum value allowed for the long field.
	Minimum int64 `json:"minimum"`
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

// The type of the field, which is LONG by default.
type LongFieldSchemaType string

const (
	LongFieldSchemaTypeLong LongFieldSchemaType = "LONG"
)

// The property Type is required.
type LongFieldSchemaParam struct {
	// The type of the field, which is LONG by default.
	//
	// Any of "LONG".
	Type LongFieldSchemaType `json:"type,omitzero" api:"required"`
	// The maximum value allowed for the long field.
	Maximum param.Opt[int64] `json:"maximum,omitzero"`
	// The minimum value allowed for the long field.
	Minimum param.Opt[int64] `json:"minimum,omitzero"`
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
	// Contains the properties of the object.
	Properties any `json:"properties" api:"required"`
	// Specifies the type of the field, which is 'OBJECT' by default.
	//
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

// Specifies the type of the field, which is 'OBJECT' by default.
type ObjectFieldSchemaType string

const (
	ObjectFieldSchemaTypeObject ObjectFieldSchemaType = "OBJECT"
)

// The properties Properties, Type are required.
type ObjectFieldSchemaParam struct {
	// Contains the properties of the object.
	Properties any `json:"properties,omitzero" api:"required"`
	// Specifies the type of the field, which is 'OBJECT' by default.
	//
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
	// The URL endpoint where the action is executed.
	ActionURL   string                            `json:"actionUrl" api:"required"`
	Functions   []PublicActionFunctionParam       `json:"functions,omitzero" api:"required"`
	InputFields []PublicInputFieldDefinitionParam `json:"inputFields,omitzero" api:"required"`
	// Holds various labels associated with the action, including names and
	// descriptions.
	Labels      map[string]PublicActionLabelsParam `json:"labels,omitzero" api:"required"`
	ObjectTypes []string                           `json:"objectTypes,omitzero" api:"required"`
	// Indicates whether the action is published and available for use.
	Published bool `json:"published" api:"required"`
	// The timestamp indicating when the action was archived.
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

type PublicActionDefinitionPatchParam struct {
	// The URL endpoint where the action is executed.
	ActionURL param.Opt[string] `json:"actionUrl,omitzero"`
	// Indicates whether the action is published and available for use.
	Published              param.Opt[bool]                                             `json:"published,omitzero"`
	ExecutionRules         []PublicExecutionTranslationRuleParam                       `json:"executionRules,omitzero"`
	InputFieldDependencies []PublicActionDefinitionPatchInputFieldDependencyUnionParam `json:"inputFieldDependencies,omitzero"`
	InputFields            []PublicInputFieldDefinitionParam                           `json:"inputFields,omitzero"`
	// Contains labels for the action, including names and descriptions.
	Labels               map[string]PublicActionLabelsParam `json:"labels,omitzero"`
	ObjectRequestOptions PublicObjectRequestOptionsParam    `json:"objectRequestOptions,omitzero"`
	ObjectTypes          []string                           `json:"objectTypes,omitzero"`
	OutputFields         []OutputFieldDefinitionParam       `json:"outputFields,omitzero"`
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

// The property RequiresObject is required.
type PublicActionDefinitionRequiresObjectRequestParam struct {
	// Indicates whether a custom action definition requires an associated object.
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
	// Indicates whether a custom action definition requires an object.
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
	// The source code or script that defines the function's behavior.
	FunctionSource string `json:"functionSource" api:"required"`
	// The type of function, with accepted values: POST_ACTION_EXECUTION,
	// POST_FETCH_OPTIONS, PRE_ACTION_EXECUTION, PRE_FETCH_OPTIONS.
	//
	// Any of "POST_ACTION_EXECUTION", "POST_FETCH_OPTIONS", "PRE_ACTION_EXECUTION",
	// "PRE_FETCH_OPTIONS".
	FunctionType PublicActionFunctionFunctionType `json:"functionType" api:"required"`
	// The unique identifier for the action function.
	ID string `json:"id"`
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

// The type of function, with accepted values: POST_ACTION_EXECUTION,
// POST_FETCH_OPTIONS, PRE_ACTION_EXECUTION, PRE_FETCH_OPTIONS.
type PublicActionFunctionFunctionType string

const (
	PublicActionFunctionFunctionTypePostActionExecution PublicActionFunctionFunctionType = "POST_ACTION_EXECUTION"
	PublicActionFunctionFunctionTypePostFetchOptions    PublicActionFunctionFunctionType = "POST_FETCH_OPTIONS"
	PublicActionFunctionFunctionTypePreActionExecution  PublicActionFunctionFunctionType = "PRE_ACTION_EXECUTION"
	PublicActionFunctionFunctionTypePreFetchOptions     PublicActionFunctionFunctionType = "PRE_FETCH_OPTIONS"
)

// The properties FunctionSource, FunctionType are required.
type PublicActionFunctionParam struct {
	// The source code or script that defines the function's behavior.
	FunctionSource string `json:"functionSource" api:"required"`
	// The type of function, with accepted values: POST_ACTION_EXECUTION,
	// POST_FETCH_OPTIONS, PRE_ACTION_EXECUTION, PRE_FETCH_OPTIONS.
	//
	// Any of "POST_ACTION_EXECUTION", "POST_FETCH_OPTIONS", "PRE_ACTION_EXECUTION",
	// "PRE_FETCH_OPTIONS".
	FunctionType PublicActionFunctionFunctionType `json:"functionType,omitzero" api:"required"`
	// The unique identifier for the action function.
	ID param.Opt[string] `json:"id,omitzero"`
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
	// The type of function, with accepted values: POST_ACTION_EXECUTION,
	// POST_FETCH_OPTIONS, PRE_ACTION_EXECUTION, PRE_FETCH_OPTIONS.
	//
	// Any of "POST_ACTION_EXECUTION", "POST_FETCH_OPTIONS", "PRE_ACTION_EXECUTION",
	// "PRE_FETCH_OPTIONS".
	FunctionType PublicActionFunctionIdentifierFunctionType `json:"functionType" api:"required"`
	// The unique identifier for the function.
	ID string `json:"id"`
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

// The type of function, with accepted values: POST_ACTION_EXECUTION,
// POST_FETCH_OPTIONS, PRE_ACTION_EXECUTION, PRE_FETCH_OPTIONS.
type PublicActionFunctionIdentifierFunctionType string

const (
	PublicActionFunctionIdentifierFunctionTypePostActionExecution PublicActionFunctionIdentifierFunctionType = "POST_ACTION_EXECUTION"
	PublicActionFunctionIdentifierFunctionTypePostFetchOptions    PublicActionFunctionIdentifierFunctionType = "POST_FETCH_OPTIONS"
	PublicActionFunctionIdentifierFunctionTypePreActionExecution  PublicActionFunctionIdentifierFunctionType = "PRE_ACTION_EXECUTION"
	PublicActionFunctionIdentifierFunctionTypePreFetchOptions     PublicActionFunctionIdentifierFunctionType = "PRE_FETCH_OPTIONS"
)

type PublicActionLabels struct {
	// The name of the action.
	ActionName string `json:"actionName" api:"required"`
	// Content displayed on the action card.
	ActionCardContent string `json:"actionCardContent"`
	// A description of what the action does.
	ActionDescription string `json:"actionDescription"`
	// The display name of the application associated with the action.
	AppDisplayName string `json:"appDisplayName"`
	// Rules that govern the execution of the action.
	ExecutionRules map[string]string `json:"executionRules"`
	// Descriptions for each input field.
	InputFieldDescriptions map[string]string `json:"inputFieldDescriptions"`
	// Labels for the input fields.
	InputFieldLabels map[string]string `json:"inputFieldLabels"`
	// Labels for the options available in input fields.
	InputFieldOptionLabels map[string]map[string]string `json:"inputFieldOptionLabels"`
	// Labels for the output fields.
	OutputFieldLabels map[string]string `json:"outputFieldLabels"`
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
	// The name of the action.
	ActionName string `json:"actionName" api:"required"`
	// Content displayed on the action card.
	ActionCardContent param.Opt[string] `json:"actionCardContent,omitzero"`
	// A description of what the action does.
	ActionDescription param.Opt[string] `json:"actionDescription,omitzero"`
	// The display name of the application associated with the action.
	AppDisplayName param.Opt[string] `json:"appDisplayName,omitzero"`
	// Rules that govern the execution of the action.
	ExecutionRules map[string]string `json:"executionRules,omitzero"`
	// Descriptions for each input field.
	InputFieldDescriptions map[string]string `json:"inputFieldDescriptions,omitzero"`
	// Labels for the input fields.
	InputFieldLabels map[string]string `json:"inputFieldLabels,omitzero"`
	// Labels for the options available in input fields.
	InputFieldOptionLabels map[string]map[string]string `json:"inputFieldOptionLabels,omitzero"`
	// Labels for the output fields.
	OutputFieldLabels map[string]string `json:"outputFieldLabels,omitzero"`
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
	// The unique identifier for the action revision.
	ID string `json:"id" api:"required"`
	// The date and time when the action revision was created.
	CreatedAt  time.Time              `json:"createdAt" api:"required" format:"date-time"`
	Definition PublicActionDefinition `json:"definition" api:"required"`
	// The unique identifier for the specific revision of the action.
	RevisionID string `json:"revisionId" api:"required"`
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
	// The name of the field that determines the dependency.
	ControllingFieldName string `json:"controllingFieldName" api:"required"`
	// The value of the controlling field that triggers the dependency.
	ControllingFieldValue string `json:"controllingFieldValue" api:"required"`
	// The type of dependency, with the default value being CONDITIONAL_SINGLE_FIELD.
	//
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

// The type of dependency, with the default value being CONDITIONAL_SINGLE_FIELD.
type PublicConditionalSingleFieldDependencyDependencyType string

const (
	PublicConditionalSingleFieldDependencyDependencyTypeConditionalSingleField PublicConditionalSingleFieldDependencyDependencyType = "CONDITIONAL_SINGLE_FIELD"
)

// The properties ControllingFieldName, ControllingFieldValue, DependencyType,
// DependentFieldNames are required.
type PublicConditionalSingleFieldDependencyParam struct {
	// The name of the field that determines the dependency.
	ControllingFieldName string `json:"controllingFieldName" api:"required"`
	// The value of the controlling field that triggers the dependency.
	ControllingFieldValue string `json:"controllingFieldValue" api:"required"`
	// The type of dependency, with the default value being CONDITIONAL_SINGLE_FIELD.
	//
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
	// Defines the conditions that must be met for the execution rule to apply.
	Conditions map[string]any `json:"conditions" api:"required"`
	// Specifies the name of the label associated with the execution rule.
	LabelName string `json:"labelName" api:"required"`
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
	// Defines the conditions that must be met for the execution rule to apply.
	Conditions map[string]any `json:"conditions,omitzero" api:"required"`
	// Specifies the name of the label associated with the execution rule.
	LabelName string `json:"labelName" api:"required"`
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
	// The internal name used to identify the field.
	Name    string         `json:"name" api:"required"`
	Options []PublicOption `json:"options" api:"required"`
	// The data type of the field, with accepted values including bool, date, datetime,
	// enumeration, json, number, object_coordinates, phone_number, and string.
	//
	// Any of "bool", "date", "datetime", "enumeration", "json", "number",
	// "object_coordinates", "phone_number", "string".
	Type PublicFieldTypeDefinitionType `json:"type" api:"required"`
	// A detailed explanation of the field's purpose.
	Description string `json:"description"`
	// The type of field, with accepted values including booleancheckbox,
	// calculation_equation, checkbox, date, file, html, number, phonenumber, radio,
	// select, text, and textarea.
	//
	// Any of "booleancheckbox", "calculation_equation", "checkbox", "date", "file",
	// "html", "number", "phonenumber", "radio", "select", "text", "textarea".
	FieldType PublicFieldTypeDefinitionFieldType `json:"fieldType"`
	// Additional information or guidance about the field.
	HelpText string `json:"helpText"`
	// A user-friendly name for the field.
	Label string `json:"label"`
	// A URL that provides options for the field.
	OptionsURL string `json:"optionsUrl"`
	// The type of object that the field references, with accepted values including
	// OWNER.
	//
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

// The data type of the field, with accepted values including bool, date, datetime,
// enumeration, json, number, object_coordinates, phone_number, and string.
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

// The type of field, with accepted values including booleancheckbox,
// calculation_equation, checkbox, date, file, html, number, phonenumber, radio,
// select, text, and textarea.
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

// The type of object that the field references, with accepted values including
// OWNER.
type PublicFieldTypeDefinitionReferencedObjectType string

const (
	PublicFieldTypeDefinitionReferencedObjectTypeOwner PublicFieldTypeDefinitionReferencedObjectType = "OWNER"
)

// The properties Name, Options, Type are required.
type PublicFieldTypeDefinitionParam struct {
	// The internal name used to identify the field.
	Name    string              `json:"name" api:"required"`
	Options []PublicOptionParam `json:"options,omitzero" api:"required"`
	// The data type of the field, with accepted values including bool, date, datetime,
	// enumeration, json, number, object_coordinates, phone_number, and string.
	//
	// Any of "bool", "date", "datetime", "enumeration", "json", "number",
	// "object_coordinates", "phone_number", "string".
	Type PublicFieldTypeDefinitionType `json:"type,omitzero" api:"required"`
	// A detailed explanation of the field's purpose.
	Description param.Opt[string] `json:"description,omitzero"`
	// Additional information or guidance about the field.
	HelpText param.Opt[string] `json:"helpText,omitzero"`
	// A user-friendly name for the field.
	Label param.Opt[string] `json:"label,omitzero"`
	// A URL that provides options for the field.
	OptionsURL param.Opt[string] `json:"optionsUrl,omitzero"`
	// The type of field, with accepted values including booleancheckbox,
	// calculation_equation, checkbox, date, file, html, number, phonenumber, radio,
	// select, text, and textarea.
	//
	// Any of "booleancheckbox", "calculation_equation", "checkbox", "date", "file",
	// "html", "number", "phonenumber", "radio", "select", "text", "textarea".
	FieldType PublicFieldTypeDefinitionFieldType `json:"fieldType,omitzero"`
	// The type of object that the field references, with accepted values including
	// OWNER.
	//
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
	// Indicates whether the input field is mandatory.
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
	// Indicates whether the input field is mandatory.
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
	// A user-friendly label that identifies the option.
	Label string `json:"label" api:"required"`
	// The actual value of the option.
	Value string `json:"value" api:"required"`
	// A description of the option.
	Description string `json:"description"`
	// The position of the option relative to others in the list.
	DisplayOrder int64 `json:"displayOrder"`
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
	// A user-friendly label that identifies the option.
	Label string `json:"label" api:"required"`
	// The actual value of the option.
	Value string `json:"value" api:"required"`
	// A description of the option.
	Description param.Opt[string] `json:"description,omitzero"`
	// The position of the option relative to others in the list.
	DisplayOrder param.Opt[int64] `json:"displayOrder,omitzero"`
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
	// The name of the field that controls the dependency.
	ControllingFieldName string `json:"controllingFieldName" api:"required"`
	// The type of dependency, with the default value being 'SINGLE_FIELD'.
	//
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

// The type of dependency, with the default value being 'SINGLE_FIELD'.
type PublicSingleFieldDependencyDependencyType string

const (
	PublicSingleFieldDependencyDependencyTypeSingleField PublicSingleFieldDependencyDependencyType = "SINGLE_FIELD"
)

// The properties ControllingFieldName, DependencyType, DependentFieldNames are
// required.
type PublicSingleFieldDependencyParam struct {
	// The name of the field that controls the dependency.
	ControllingFieldName string `json:"controllingFieldName" api:"required"`
	// The type of dependency, with the default value being 'SINGLE_FIELD'.
	//
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
	// Indicates the source of the request, with the default value being 'STANDALONE'.
	//
	// Any of "STANDALONE".
	Source StandaloneRequestContextSource `json:"source,omitzero" api:"required"`
	// A unique identifier for tracking the trajectory of the request.
	TrajectoryID param.Opt[string] `json:"trajectoryId,omitzero"`
	paramObj
}

func (r StandaloneRequestContextParam) MarshalJSON() (data []byte, err error) {
	type shadow StandaloneRequestContextParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *StandaloneRequestContextParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Indicates the source of the request, with the default value being 'STANDALONE'.
type StandaloneRequestContextSource string

const (
	StandaloneRequestContextSourceStandalone StandaloneRequestContextSource = "STANDALONE"
)

type StringFieldSchema struct {
	// Indicates that the type is a string, with the default value being STRING.
	//
	// Any of "STRING".
	Type StringFieldSchemaType `json:"type" api:"required"`
	// Specifies the format of the string, with accepted values: DATE, DATE_TIME,
	// OBJECT_COORDINATE, TIME, URI.
	//
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

// Indicates that the type is a string, with the default value being STRING.
type StringFieldSchemaType string

const (
	StringFieldSchemaTypeString StringFieldSchemaType = "STRING"
)

// Specifies the format of the string, with accepted values: DATE, DATE_TIME,
// OBJECT_COORDINATE, TIME, URI.
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
	// Indicates that the type is a string, with the default value being STRING.
	//
	// Any of "STRING".
	Type StringFieldSchemaType `json:"type,omitzero" api:"required"`
	// Specifies the format of the string, with accepted values: DATE, DATE_TIME,
	// OBJECT_COORDINATE, TIME, URI.
	//
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
	// Indicates the source of the test request, with the only accepted value being
	// 'TEST'.
	//
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

// Indicates the source of the test request, with the only accepted value being
// 'TEST'.
type TestRequestContextSource string

const (
	TestRequestContextSourceTest TestRequestContextSource = "TEST"
)

// The properties Source, WorkflowID are required.
type WorkflowsRequestContextParam struct {
	// Indicates the source of the request, with the default value being WORKFLOWS.
	//
	// Any of "WORKFLOWS".
	Source WorkflowsRequestContextSource `json:"source,omitzero" api:"required"`
	// The ID of the workflow associated with the request context.
	WorkflowID int64 `json:"workflowId" api:"required"`
	// The ID of the action within the workflow context.
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

// Indicates the source of the request, with the default value being WORKFLOWS.
type WorkflowsRequestContextSource string

const (
	WorkflowsRequestContextSourceWorkflows WorkflowsRequestContextSource = "WORKFLOWS"
)
