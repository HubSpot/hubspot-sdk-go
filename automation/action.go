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
	Options     []option.RequestOption
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
	r.Options = opts
	r.Callbacks = NewActionCallbackService(opts...)
	r.Definitions = NewActionDefinitionService(opts...)
	r.Functions = NewActionFunctionService(opts...)
	r.Revisions = NewActionRevisionService(opts...)
	return
}

// The property Inputs is required.
type BatchInputCallbackCompletionBatchRequestParam struct {
	Inputs []CallbackCompletionBatchRequestParam `json:"inputs,omitzero,required"`
	paramObj
}

func (r BatchInputCallbackCompletionBatchRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow BatchInputCallbackCompletionBatchRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BatchInputCallbackCompletionBatchRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties CallbackID, OutputFields are required.
type CallbackCompletionBatchRequestParam struct {
	CallbackID   string            `json:"callbackId,required"`
	OutputFields map[string]string `json:"outputFields,omitzero,required"`
	paramObj
}

func (r CallbackCompletionBatchRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow CallbackCompletionBatchRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CallbackCompletionBatchRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property OutputFields is required.
type CallbackCompletionRequestParam struct {
	OutputFields map[string]string `json:"outputFields,omitzero,required"`
	paramObj
}

func (r CallbackCompletionRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow CallbackCompletionRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CallbackCompletionRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CollectionResponsePublicActionDefinitionForwardPaging struct {
	Results []PublicActionDefinition `json:"results,required"`
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
	Results []PublicActionFunctionIdentifier `json:"results,required"`
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
	Results []PublicActionRevision `json:"results,required"`
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

type FieldTypeDefinition struct {
	ExternalOptions bool            `json:"externalOptions,required"`
	Name            string          `json:"name,required"`
	Options         []shared.Option `json:"options,required"`
	// Any of "bool", "currency_number", "date", "datetime", "enumeration", "json",
	// "number", "object_coordinates", "phone_number", "string".
	Type                         FieldTypeDefinitionType `json:"type,required"`
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
	// "CAMPAIGN_TEMPLATE_STEP", "CART", "CHATFLOW", "CLIP", "CMS_URL",
	// "COMMERCE_PAYMENT", "COMMUNICATION", "COMPANY", "CONTACT",
	// "CONTACT_CREATE_ATTRIBUTION", "CONTENT", "CONTENT_AUDIT", "CONTENT_AUDIT_PAGE",
	// "CONVERSATION", "CONVERSATION_INBOX", "CONVERSATION_SESSION",
	// "CRM_PIPELINES_DUMMY_TYPE", "CTA", "CTA_VARIANT", "DATA_PRIVACY_CONSENT",
	// "DATA_SYNC_STATE", "DEAL", "DEAL_CREATE_ATTRIBUTION", "DEAL_REGISTRATION",
	// "DEAL_SPLIT", "DISCOUNT", "DISCOUNT_CODE", "DISCOUNT_TEMPLATE", "EMAIL",
	// "ENGAGEMENT", "EXPORT", "EXTERNAL_WEB_URL", "FEE", "FEEDBACK_SUBMISSION",
	// "FEEDBACK_SURVEY", "FILE_MANAGER_FILE", "FILE_MANAGER_FOLDER", "FOLDER",
	// "FORECAST", "FORM", "FORM_SUBMISSION_INBOUNDDB", "GOAL_TARGET",
	// "GOAL_TARGET_GROUP", "GOAL_TEMPLATE", "GSC_PROPERTY", "HUB", "IMPORT",
	// "INVOICE", "KEYWORD", "KNOWLEDGE_ARTICLE", "LANDING_PAGE", "LEAD", "LINE_ITEM",
	// "MARKETING_CALENDAR", "MARKETING_CAMPAIGN_UTM", "MARKETING_EMAIL",
	// "MARKETING_EVENT", "MARKETING_EVENT_ATTENDANCE", "MARKETING_SMS",
	// "MEDIA_BRIDGE", "MEETING_EVENT", "MIC", "NOTE", "OBJECT_LIST", "ORDER", "OWNER",
	// "PARTNER_ACCOUNT", "PARTNER_CLIENT", "PARTNER_CLIENT_REVENUE", "PAYMENT_LINK",
	// "PAYMENT_SCHEDULE", "PAYMENT_SCHEDULE_INSTALLMENT", "PERMISSIONS_TESTING",
	// "PLAYBOOK", "PLAYBOOK_QUESTION", "PLAYBOOK_SUBMISSION",
	// "PLAYBOOK_SUBMISSION_ANSWER", "PLAYLIST", "PLAYLIST_FOLDER", "PORTAL",
	// "PORTAL_OBJECT_SYNC_MESSAGE", "POSTAL_MAIL", "PRIVACY_SCANNER_COOKIE",
	// "PRODUCT", "PRODUCT_OR_FOLDER", "PROPERTY_INFO", "PUBLISHING_TASK",
	// "QUARANTINED_SUBMISSION", "QUOTA", "QUOTE", "QUOTE_FIELD", "QUOTE_MODULE",
	// "QUOTE_MODULE_FIELD", "QUOTE_TEMPLATE", "RESTORABLE_CRM_OBJECT", "ROSTER",
	// "ROSTER_MEMBER", "SALES_DOCUMENT", "SALES_TASK", "SALES_WORKLOAD",
	// "SALESFORCE_SYNC_ERROR", "SCHEDULING_PAGE", "SCHEMAS_BACKEND_TEST",
	// "SCORE_CONFIGURATION", "SEQUENCE", "SEQUENCE_ENROLLMENT", "SEQUENCE_STEP",
	// "SEQUENCE_STEP_ENROLLMENT", "SITE_PAGE", "SNIPPET", "SOCIAL_BROADCAST",
	// "SOCIAL_CHANNEL", "SOCIAL_POST", "SOCIAL_PROFILE", "SOX_PROTECTED_DUMMY_TYPE",
	// "SOX_PROTECTED_TEST_TYPE", "SUBMISSION_TAG", "SUBSCRIPTION", "TASK",
	// "TASK_TEMPLATE", "TAX", "TEMPLATE", "TICKET", "UNKNOWN", "UNSUBSCRIBE", "USER",
	// "VIEW", "VIEW_BLOCK", "WEB_INTERACTIVE".
	ReferencedObjectType FieldTypeDefinitionReferencedObjectType `json:"referencedObjectType"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ExternalOptions              respjson.Field
		Name                         respjson.Field
		Options                      respjson.Field
		Type                         respjson.Field
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
	FieldTypeDefinitionReferencedObjectTypeAbandonedCart                FieldTypeDefinitionReferencedObjectType = "ABANDONED_CART"
	FieldTypeDefinitionReferencedObjectTypeAcceptanceTest               FieldTypeDefinitionReferencedObjectType = "ACCEPTANCE_TEST"
	FieldTypeDefinitionReferencedObjectTypeAd                           FieldTypeDefinitionReferencedObjectType = "AD"
	FieldTypeDefinitionReferencedObjectTypeAdAccount                    FieldTypeDefinitionReferencedObjectType = "AD_ACCOUNT"
	FieldTypeDefinitionReferencedObjectTypeAdCampaign                   FieldTypeDefinitionReferencedObjectType = "AD_CAMPAIGN"
	FieldTypeDefinitionReferencedObjectTypeAdGroup                      FieldTypeDefinitionReferencedObjectType = "AD_GROUP"
	FieldTypeDefinitionReferencedObjectTypeAIForecast                   FieldTypeDefinitionReferencedObjectType = "AI_FORECAST"
	FieldTypeDefinitionReferencedObjectTypeAllPages                     FieldTypeDefinitionReferencedObjectType = "ALL_PAGES"
	FieldTypeDefinitionReferencedObjectTypeApproval                     FieldTypeDefinitionReferencedObjectType = "APPROVAL"
	FieldTypeDefinitionReferencedObjectTypeApprovalStep                 FieldTypeDefinitionReferencedObjectType = "APPROVAL_STEP"
	FieldTypeDefinitionReferencedObjectTypeAttribution                  FieldTypeDefinitionReferencedObjectType = "ATTRIBUTION"
	FieldTypeDefinitionReferencedObjectTypeAudience                     FieldTypeDefinitionReferencedObjectType = "AUDIENCE"
	FieldTypeDefinitionReferencedObjectTypeAutomationJourney            FieldTypeDefinitionReferencedObjectType = "AUTOMATION_JOURNEY"
	FieldTypeDefinitionReferencedObjectTypeAutomationPlatformFlow       FieldTypeDefinitionReferencedObjectType = "AUTOMATION_PLATFORM_FLOW"
	FieldTypeDefinitionReferencedObjectTypeAutomationPlatformFlowAction FieldTypeDefinitionReferencedObjectType = "AUTOMATION_PLATFORM_FLOW_ACTION"
	FieldTypeDefinitionReferencedObjectTypeBetAlert                     FieldTypeDefinitionReferencedObjectType = "BET_ALERT"
	FieldTypeDefinitionReferencedObjectTypeBetDeliverableService        FieldTypeDefinitionReferencedObjectType = "BET_DELIVERABLE_SERVICE"
	FieldTypeDefinitionReferencedObjectTypeBlogListingPage              FieldTypeDefinitionReferencedObjectType = "BLOG_LISTING_PAGE"
	FieldTypeDefinitionReferencedObjectTypeBlogPost                     FieldTypeDefinitionReferencedObjectType = "BLOG_POST"
	FieldTypeDefinitionReferencedObjectTypeCall                         FieldTypeDefinitionReferencedObjectType = "CALL"
	FieldTypeDefinitionReferencedObjectTypeCampaign                     FieldTypeDefinitionReferencedObjectType = "CAMPAIGN"
	FieldTypeDefinitionReferencedObjectTypeCampaignBudgetItem           FieldTypeDefinitionReferencedObjectType = "CAMPAIGN_BUDGET_ITEM"
	FieldTypeDefinitionReferencedObjectTypeCampaignSpendItem            FieldTypeDefinitionReferencedObjectType = "CAMPAIGN_SPEND_ITEM"
	FieldTypeDefinitionReferencedObjectTypeCampaignStep                 FieldTypeDefinitionReferencedObjectType = "CAMPAIGN_STEP"
	FieldTypeDefinitionReferencedObjectTypeCampaignTemplate             FieldTypeDefinitionReferencedObjectType = "CAMPAIGN_TEMPLATE"
	FieldTypeDefinitionReferencedObjectTypeCampaignTemplateStep         FieldTypeDefinitionReferencedObjectType = "CAMPAIGN_TEMPLATE_STEP"
	FieldTypeDefinitionReferencedObjectTypeCart                         FieldTypeDefinitionReferencedObjectType = "CART"
	FieldTypeDefinitionReferencedObjectTypeChatflow                     FieldTypeDefinitionReferencedObjectType = "CHATFLOW"
	FieldTypeDefinitionReferencedObjectTypeClip                         FieldTypeDefinitionReferencedObjectType = "CLIP"
	FieldTypeDefinitionReferencedObjectTypeCmsURL                       FieldTypeDefinitionReferencedObjectType = "CMS_URL"
	FieldTypeDefinitionReferencedObjectTypeCommercePayment              FieldTypeDefinitionReferencedObjectType = "COMMERCE_PAYMENT"
	FieldTypeDefinitionReferencedObjectTypeCommunication                FieldTypeDefinitionReferencedObjectType = "COMMUNICATION"
	FieldTypeDefinitionReferencedObjectTypeCompany                      FieldTypeDefinitionReferencedObjectType = "COMPANY"
	FieldTypeDefinitionReferencedObjectTypeContact                      FieldTypeDefinitionReferencedObjectType = "CONTACT"
	FieldTypeDefinitionReferencedObjectTypeContactCreateAttribution     FieldTypeDefinitionReferencedObjectType = "CONTACT_CREATE_ATTRIBUTION"
	FieldTypeDefinitionReferencedObjectTypeContent                      FieldTypeDefinitionReferencedObjectType = "CONTENT"
	FieldTypeDefinitionReferencedObjectTypeContentAudit                 FieldTypeDefinitionReferencedObjectType = "CONTENT_AUDIT"
	FieldTypeDefinitionReferencedObjectTypeContentAuditPage             FieldTypeDefinitionReferencedObjectType = "CONTENT_AUDIT_PAGE"
	FieldTypeDefinitionReferencedObjectTypeConversation                 FieldTypeDefinitionReferencedObjectType = "CONVERSATION"
	FieldTypeDefinitionReferencedObjectTypeConversationInbox            FieldTypeDefinitionReferencedObjectType = "CONVERSATION_INBOX"
	FieldTypeDefinitionReferencedObjectTypeConversationSession          FieldTypeDefinitionReferencedObjectType = "CONVERSATION_SESSION"
	FieldTypeDefinitionReferencedObjectTypeCrmPipelinesDummyType        FieldTypeDefinitionReferencedObjectType = "CRM_PIPELINES_DUMMY_TYPE"
	FieldTypeDefinitionReferencedObjectTypeCta                          FieldTypeDefinitionReferencedObjectType = "CTA"
	FieldTypeDefinitionReferencedObjectTypeCtaVariant                   FieldTypeDefinitionReferencedObjectType = "CTA_VARIANT"
	FieldTypeDefinitionReferencedObjectTypeDataPrivacyConsent           FieldTypeDefinitionReferencedObjectType = "DATA_PRIVACY_CONSENT"
	FieldTypeDefinitionReferencedObjectTypeDataSyncState                FieldTypeDefinitionReferencedObjectType = "DATA_SYNC_STATE"
	FieldTypeDefinitionReferencedObjectTypeDeal                         FieldTypeDefinitionReferencedObjectType = "DEAL"
	FieldTypeDefinitionReferencedObjectTypeDealCreateAttribution        FieldTypeDefinitionReferencedObjectType = "DEAL_CREATE_ATTRIBUTION"
	FieldTypeDefinitionReferencedObjectTypeDealRegistration             FieldTypeDefinitionReferencedObjectType = "DEAL_REGISTRATION"
	FieldTypeDefinitionReferencedObjectTypeDealSplit                    FieldTypeDefinitionReferencedObjectType = "DEAL_SPLIT"
	FieldTypeDefinitionReferencedObjectTypeDiscount                     FieldTypeDefinitionReferencedObjectType = "DISCOUNT"
	FieldTypeDefinitionReferencedObjectTypeDiscountCode                 FieldTypeDefinitionReferencedObjectType = "DISCOUNT_CODE"
	FieldTypeDefinitionReferencedObjectTypeDiscountTemplate             FieldTypeDefinitionReferencedObjectType = "DISCOUNT_TEMPLATE"
	FieldTypeDefinitionReferencedObjectTypeEmail                        FieldTypeDefinitionReferencedObjectType = "EMAIL"
	FieldTypeDefinitionReferencedObjectTypeEngagement                   FieldTypeDefinitionReferencedObjectType = "ENGAGEMENT"
	FieldTypeDefinitionReferencedObjectTypeExport                       FieldTypeDefinitionReferencedObjectType = "EXPORT"
	FieldTypeDefinitionReferencedObjectTypeExternalWebURL               FieldTypeDefinitionReferencedObjectType = "EXTERNAL_WEB_URL"
	FieldTypeDefinitionReferencedObjectTypeFee                          FieldTypeDefinitionReferencedObjectType = "FEE"
	FieldTypeDefinitionReferencedObjectTypeFeedbackSubmission           FieldTypeDefinitionReferencedObjectType = "FEEDBACK_SUBMISSION"
	FieldTypeDefinitionReferencedObjectTypeFeedbackSurvey               FieldTypeDefinitionReferencedObjectType = "FEEDBACK_SURVEY"
	FieldTypeDefinitionReferencedObjectTypeFileManagerFile              FieldTypeDefinitionReferencedObjectType = "FILE_MANAGER_FILE"
	FieldTypeDefinitionReferencedObjectTypeFileManagerFolder            FieldTypeDefinitionReferencedObjectType = "FILE_MANAGER_FOLDER"
	FieldTypeDefinitionReferencedObjectTypeFolder                       FieldTypeDefinitionReferencedObjectType = "FOLDER"
	FieldTypeDefinitionReferencedObjectTypeForecast                     FieldTypeDefinitionReferencedObjectType = "FORECAST"
	FieldTypeDefinitionReferencedObjectTypeForm                         FieldTypeDefinitionReferencedObjectType = "FORM"
	FieldTypeDefinitionReferencedObjectTypeFormSubmissionInbounddb      FieldTypeDefinitionReferencedObjectType = "FORM_SUBMISSION_INBOUNDDB"
	FieldTypeDefinitionReferencedObjectTypeGoalTarget                   FieldTypeDefinitionReferencedObjectType = "GOAL_TARGET"
	FieldTypeDefinitionReferencedObjectTypeGoalTargetGroup              FieldTypeDefinitionReferencedObjectType = "GOAL_TARGET_GROUP"
	FieldTypeDefinitionReferencedObjectTypeGoalTemplate                 FieldTypeDefinitionReferencedObjectType = "GOAL_TEMPLATE"
	FieldTypeDefinitionReferencedObjectTypeGscProperty                  FieldTypeDefinitionReferencedObjectType = "GSC_PROPERTY"
	FieldTypeDefinitionReferencedObjectTypeHub                          FieldTypeDefinitionReferencedObjectType = "HUB"
	FieldTypeDefinitionReferencedObjectTypeImport                       FieldTypeDefinitionReferencedObjectType = "IMPORT"
	FieldTypeDefinitionReferencedObjectTypeInvoice                      FieldTypeDefinitionReferencedObjectType = "INVOICE"
	FieldTypeDefinitionReferencedObjectTypeKeyword                      FieldTypeDefinitionReferencedObjectType = "KEYWORD"
	FieldTypeDefinitionReferencedObjectTypeKnowledgeArticle             FieldTypeDefinitionReferencedObjectType = "KNOWLEDGE_ARTICLE"
	FieldTypeDefinitionReferencedObjectTypeLandingPage                  FieldTypeDefinitionReferencedObjectType = "LANDING_PAGE"
	FieldTypeDefinitionReferencedObjectTypeLead                         FieldTypeDefinitionReferencedObjectType = "LEAD"
	FieldTypeDefinitionReferencedObjectTypeLineItem                     FieldTypeDefinitionReferencedObjectType = "LINE_ITEM"
	FieldTypeDefinitionReferencedObjectTypeMarketingCalendar            FieldTypeDefinitionReferencedObjectType = "MARKETING_CALENDAR"
	FieldTypeDefinitionReferencedObjectTypeMarketingCampaignUtm         FieldTypeDefinitionReferencedObjectType = "MARKETING_CAMPAIGN_UTM"
	FieldTypeDefinitionReferencedObjectTypeMarketingEmail               FieldTypeDefinitionReferencedObjectType = "MARKETING_EMAIL"
	FieldTypeDefinitionReferencedObjectTypeMarketingEvent               FieldTypeDefinitionReferencedObjectType = "MARKETING_EVENT"
	FieldTypeDefinitionReferencedObjectTypeMarketingEventAttendance     FieldTypeDefinitionReferencedObjectType = "MARKETING_EVENT_ATTENDANCE"
	FieldTypeDefinitionReferencedObjectTypeMarketingSMS                 FieldTypeDefinitionReferencedObjectType = "MARKETING_SMS"
	FieldTypeDefinitionReferencedObjectTypeMediaBridge                  FieldTypeDefinitionReferencedObjectType = "MEDIA_BRIDGE"
	FieldTypeDefinitionReferencedObjectTypeMeetingEvent                 FieldTypeDefinitionReferencedObjectType = "MEETING_EVENT"
	FieldTypeDefinitionReferencedObjectTypeMic                          FieldTypeDefinitionReferencedObjectType = "MIC"
	FieldTypeDefinitionReferencedObjectTypeNote                         FieldTypeDefinitionReferencedObjectType = "NOTE"
	FieldTypeDefinitionReferencedObjectTypeObjectList                   FieldTypeDefinitionReferencedObjectType = "OBJECT_LIST"
	FieldTypeDefinitionReferencedObjectTypeOrder                        FieldTypeDefinitionReferencedObjectType = "ORDER"
	FieldTypeDefinitionReferencedObjectTypeOwner                        FieldTypeDefinitionReferencedObjectType = "OWNER"
	FieldTypeDefinitionReferencedObjectTypePartnerAccount               FieldTypeDefinitionReferencedObjectType = "PARTNER_ACCOUNT"
	FieldTypeDefinitionReferencedObjectTypePartnerClient                FieldTypeDefinitionReferencedObjectType = "PARTNER_CLIENT"
	FieldTypeDefinitionReferencedObjectTypePartnerClientRevenue         FieldTypeDefinitionReferencedObjectType = "PARTNER_CLIENT_REVENUE"
	FieldTypeDefinitionReferencedObjectTypePaymentLink                  FieldTypeDefinitionReferencedObjectType = "PAYMENT_LINK"
	FieldTypeDefinitionReferencedObjectTypePaymentSchedule              FieldTypeDefinitionReferencedObjectType = "PAYMENT_SCHEDULE"
	FieldTypeDefinitionReferencedObjectTypePaymentScheduleInstallment   FieldTypeDefinitionReferencedObjectType = "PAYMENT_SCHEDULE_INSTALLMENT"
	FieldTypeDefinitionReferencedObjectTypePermissionsTesting           FieldTypeDefinitionReferencedObjectType = "PERMISSIONS_TESTING"
	FieldTypeDefinitionReferencedObjectTypePlaybook                     FieldTypeDefinitionReferencedObjectType = "PLAYBOOK"
	FieldTypeDefinitionReferencedObjectTypePlaybookQuestion             FieldTypeDefinitionReferencedObjectType = "PLAYBOOK_QUESTION"
	FieldTypeDefinitionReferencedObjectTypePlaybookSubmission           FieldTypeDefinitionReferencedObjectType = "PLAYBOOK_SUBMISSION"
	FieldTypeDefinitionReferencedObjectTypePlaybookSubmissionAnswer     FieldTypeDefinitionReferencedObjectType = "PLAYBOOK_SUBMISSION_ANSWER"
	FieldTypeDefinitionReferencedObjectTypePlaylist                     FieldTypeDefinitionReferencedObjectType = "PLAYLIST"
	FieldTypeDefinitionReferencedObjectTypePlaylistFolder               FieldTypeDefinitionReferencedObjectType = "PLAYLIST_FOLDER"
	FieldTypeDefinitionReferencedObjectTypePortal                       FieldTypeDefinitionReferencedObjectType = "PORTAL"
	FieldTypeDefinitionReferencedObjectTypePortalObjectSyncMessage      FieldTypeDefinitionReferencedObjectType = "PORTAL_OBJECT_SYNC_MESSAGE"
	FieldTypeDefinitionReferencedObjectTypePostalMail                   FieldTypeDefinitionReferencedObjectType = "POSTAL_MAIL"
	FieldTypeDefinitionReferencedObjectTypePrivacyScannerCookie         FieldTypeDefinitionReferencedObjectType = "PRIVACY_SCANNER_COOKIE"
	FieldTypeDefinitionReferencedObjectTypeProduct                      FieldTypeDefinitionReferencedObjectType = "PRODUCT"
	FieldTypeDefinitionReferencedObjectTypeProductOrFolder              FieldTypeDefinitionReferencedObjectType = "PRODUCT_OR_FOLDER"
	FieldTypeDefinitionReferencedObjectTypePropertyInfo                 FieldTypeDefinitionReferencedObjectType = "PROPERTY_INFO"
	FieldTypeDefinitionReferencedObjectTypePublishingTask               FieldTypeDefinitionReferencedObjectType = "PUBLISHING_TASK"
	FieldTypeDefinitionReferencedObjectTypeQuarantinedSubmission        FieldTypeDefinitionReferencedObjectType = "QUARANTINED_SUBMISSION"
	FieldTypeDefinitionReferencedObjectTypeQuota                        FieldTypeDefinitionReferencedObjectType = "QUOTA"
	FieldTypeDefinitionReferencedObjectTypeQuote                        FieldTypeDefinitionReferencedObjectType = "QUOTE"
	FieldTypeDefinitionReferencedObjectTypeQuoteField                   FieldTypeDefinitionReferencedObjectType = "QUOTE_FIELD"
	FieldTypeDefinitionReferencedObjectTypeQuoteModule                  FieldTypeDefinitionReferencedObjectType = "QUOTE_MODULE"
	FieldTypeDefinitionReferencedObjectTypeQuoteModuleField             FieldTypeDefinitionReferencedObjectType = "QUOTE_MODULE_FIELD"
	FieldTypeDefinitionReferencedObjectTypeQuoteTemplate                FieldTypeDefinitionReferencedObjectType = "QUOTE_TEMPLATE"
	FieldTypeDefinitionReferencedObjectTypeRestorableCrmObject          FieldTypeDefinitionReferencedObjectType = "RESTORABLE_CRM_OBJECT"
	FieldTypeDefinitionReferencedObjectTypeRoster                       FieldTypeDefinitionReferencedObjectType = "ROSTER"
	FieldTypeDefinitionReferencedObjectTypeRosterMember                 FieldTypeDefinitionReferencedObjectType = "ROSTER_MEMBER"
	FieldTypeDefinitionReferencedObjectTypeSalesDocument                FieldTypeDefinitionReferencedObjectType = "SALES_DOCUMENT"
	FieldTypeDefinitionReferencedObjectTypeSalesTask                    FieldTypeDefinitionReferencedObjectType = "SALES_TASK"
	FieldTypeDefinitionReferencedObjectTypeSalesWorkload                FieldTypeDefinitionReferencedObjectType = "SALES_WORKLOAD"
	FieldTypeDefinitionReferencedObjectTypeSalesforceSyncError          FieldTypeDefinitionReferencedObjectType = "SALESFORCE_SYNC_ERROR"
	FieldTypeDefinitionReferencedObjectTypeSchedulingPage               FieldTypeDefinitionReferencedObjectType = "SCHEDULING_PAGE"
	FieldTypeDefinitionReferencedObjectTypeSchemasBackendTest           FieldTypeDefinitionReferencedObjectType = "SCHEMAS_BACKEND_TEST"
	FieldTypeDefinitionReferencedObjectTypeScoreConfiguration           FieldTypeDefinitionReferencedObjectType = "SCORE_CONFIGURATION"
	FieldTypeDefinitionReferencedObjectTypeSequence                     FieldTypeDefinitionReferencedObjectType = "SEQUENCE"
	FieldTypeDefinitionReferencedObjectTypeSequenceEnrollment           FieldTypeDefinitionReferencedObjectType = "SEQUENCE_ENROLLMENT"
	FieldTypeDefinitionReferencedObjectTypeSequenceStep                 FieldTypeDefinitionReferencedObjectType = "SEQUENCE_STEP"
	FieldTypeDefinitionReferencedObjectTypeSequenceStepEnrollment       FieldTypeDefinitionReferencedObjectType = "SEQUENCE_STEP_ENROLLMENT"
	FieldTypeDefinitionReferencedObjectTypeSitePage                     FieldTypeDefinitionReferencedObjectType = "SITE_PAGE"
	FieldTypeDefinitionReferencedObjectTypeSnippet                      FieldTypeDefinitionReferencedObjectType = "SNIPPET"
	FieldTypeDefinitionReferencedObjectTypeSocialBroadcast              FieldTypeDefinitionReferencedObjectType = "SOCIAL_BROADCAST"
	FieldTypeDefinitionReferencedObjectTypeSocialChannel                FieldTypeDefinitionReferencedObjectType = "SOCIAL_CHANNEL"
	FieldTypeDefinitionReferencedObjectTypeSocialPost                   FieldTypeDefinitionReferencedObjectType = "SOCIAL_POST"
	FieldTypeDefinitionReferencedObjectTypeSocialProfile                FieldTypeDefinitionReferencedObjectType = "SOCIAL_PROFILE"
	FieldTypeDefinitionReferencedObjectTypeSoxProtectedDummyType        FieldTypeDefinitionReferencedObjectType = "SOX_PROTECTED_DUMMY_TYPE"
	FieldTypeDefinitionReferencedObjectTypeSoxProtectedTestType         FieldTypeDefinitionReferencedObjectType = "SOX_PROTECTED_TEST_TYPE"
	FieldTypeDefinitionReferencedObjectTypeSubmissionTag                FieldTypeDefinitionReferencedObjectType = "SUBMISSION_TAG"
	FieldTypeDefinitionReferencedObjectTypeSubscription                 FieldTypeDefinitionReferencedObjectType = "SUBSCRIPTION"
	FieldTypeDefinitionReferencedObjectTypeTask                         FieldTypeDefinitionReferencedObjectType = "TASK"
	FieldTypeDefinitionReferencedObjectTypeTaskTemplate                 FieldTypeDefinitionReferencedObjectType = "TASK_TEMPLATE"
	FieldTypeDefinitionReferencedObjectTypeTax                          FieldTypeDefinitionReferencedObjectType = "TAX"
	FieldTypeDefinitionReferencedObjectTypeTemplate                     FieldTypeDefinitionReferencedObjectType = "TEMPLATE"
	FieldTypeDefinitionReferencedObjectTypeTicket                       FieldTypeDefinitionReferencedObjectType = "TICKET"
	FieldTypeDefinitionReferencedObjectTypeUnknown                      FieldTypeDefinitionReferencedObjectType = "UNKNOWN"
	FieldTypeDefinitionReferencedObjectTypeUnsubscribe                  FieldTypeDefinitionReferencedObjectType = "UNSUBSCRIBE"
	FieldTypeDefinitionReferencedObjectTypeUser                         FieldTypeDefinitionReferencedObjectType = "USER"
	FieldTypeDefinitionReferencedObjectTypeView                         FieldTypeDefinitionReferencedObjectType = "VIEW"
	FieldTypeDefinitionReferencedObjectTypeViewBlock                    FieldTypeDefinitionReferencedObjectType = "VIEW_BLOCK"
	FieldTypeDefinitionReferencedObjectTypeWebInteractive               FieldTypeDefinitionReferencedObjectType = "WEB_INTERACTIVE"
)

// The properties ExternalOptions, Name, Options, Type are required.
type FieldTypeDefinitionParam struct {
	ExternalOptions bool                 `json:"externalOptions,required"`
	Name            string               `json:"name,required"`
	Options         []shared.OptionParam `json:"options,omitzero,required"`
	// Any of "bool", "currency_number", "date", "datetime", "enumeration", "json",
	// "number", "object_coordinates", "phone_number", "string".
	Type                         FieldTypeDefinitionType `json:"type,omitzero,required"`
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
	// "CAMPAIGN_TEMPLATE_STEP", "CART", "CHATFLOW", "CLIP", "CMS_URL",
	// "COMMERCE_PAYMENT", "COMMUNICATION", "COMPANY", "CONTACT",
	// "CONTACT_CREATE_ATTRIBUTION", "CONTENT", "CONTENT_AUDIT", "CONTENT_AUDIT_PAGE",
	// "CONVERSATION", "CONVERSATION_INBOX", "CONVERSATION_SESSION",
	// "CRM_PIPELINES_DUMMY_TYPE", "CTA", "CTA_VARIANT", "DATA_PRIVACY_CONSENT",
	// "DATA_SYNC_STATE", "DEAL", "DEAL_CREATE_ATTRIBUTION", "DEAL_REGISTRATION",
	// "DEAL_SPLIT", "DISCOUNT", "DISCOUNT_CODE", "DISCOUNT_TEMPLATE", "EMAIL",
	// "ENGAGEMENT", "EXPORT", "EXTERNAL_WEB_URL", "FEE", "FEEDBACK_SUBMISSION",
	// "FEEDBACK_SURVEY", "FILE_MANAGER_FILE", "FILE_MANAGER_FOLDER", "FOLDER",
	// "FORECAST", "FORM", "FORM_SUBMISSION_INBOUNDDB", "GOAL_TARGET",
	// "GOAL_TARGET_GROUP", "GOAL_TEMPLATE", "GSC_PROPERTY", "HUB", "IMPORT",
	// "INVOICE", "KEYWORD", "KNOWLEDGE_ARTICLE", "LANDING_PAGE", "LEAD", "LINE_ITEM",
	// "MARKETING_CALENDAR", "MARKETING_CAMPAIGN_UTM", "MARKETING_EMAIL",
	// "MARKETING_EVENT", "MARKETING_EVENT_ATTENDANCE", "MARKETING_SMS",
	// "MEDIA_BRIDGE", "MEETING_EVENT", "MIC", "NOTE", "OBJECT_LIST", "ORDER", "OWNER",
	// "PARTNER_ACCOUNT", "PARTNER_CLIENT", "PARTNER_CLIENT_REVENUE", "PAYMENT_LINK",
	// "PAYMENT_SCHEDULE", "PAYMENT_SCHEDULE_INSTALLMENT", "PERMISSIONS_TESTING",
	// "PLAYBOOK", "PLAYBOOK_QUESTION", "PLAYBOOK_SUBMISSION",
	// "PLAYBOOK_SUBMISSION_ANSWER", "PLAYLIST", "PLAYLIST_FOLDER", "PORTAL",
	// "PORTAL_OBJECT_SYNC_MESSAGE", "POSTAL_MAIL", "PRIVACY_SCANNER_COOKIE",
	// "PRODUCT", "PRODUCT_OR_FOLDER", "PROPERTY_INFO", "PUBLISHING_TASK",
	// "QUARANTINED_SUBMISSION", "QUOTA", "QUOTE", "QUOTE_FIELD", "QUOTE_MODULE",
	// "QUOTE_MODULE_FIELD", "QUOTE_TEMPLATE", "RESTORABLE_CRM_OBJECT", "ROSTER",
	// "ROSTER_MEMBER", "SALES_DOCUMENT", "SALES_TASK", "SALES_WORKLOAD",
	// "SALESFORCE_SYNC_ERROR", "SCHEDULING_PAGE", "SCHEMAS_BACKEND_TEST",
	// "SCORE_CONFIGURATION", "SEQUENCE", "SEQUENCE_ENROLLMENT", "SEQUENCE_STEP",
	// "SEQUENCE_STEP_ENROLLMENT", "SITE_PAGE", "SNIPPET", "SOCIAL_BROADCAST",
	// "SOCIAL_CHANNEL", "SOCIAL_POST", "SOCIAL_PROFILE", "SOX_PROTECTED_DUMMY_TYPE",
	// "SOX_PROTECTED_TEST_TYPE", "SUBMISSION_TAG", "SUBSCRIPTION", "TASK",
	// "TASK_TEMPLATE", "TAX", "TEMPLATE", "TICKET", "UNKNOWN", "UNSUBSCRIBE", "USER",
	// "VIEW", "VIEW_BLOCK", "WEB_INTERACTIVE".
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

type InputFieldDefinition struct {
	IsRequired          bool                `json:"isRequired,required"`
	TypeDefinition      FieldTypeDefinition `json:"typeDefinition,required"`
	AutomationFieldType string              `json:"automationFieldType"`
	// Any of "STATIC_VALUE", "OBJECT_PROPERTY", "FIELD_DATA",
	// "FETCHED_OBJECT_PROPERTY", "ENROLLMENT_EVENT_PROPERTY".
	SupportedValueTypes []string `json:"supportedValueTypes"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		IsRequired          respjson.Field
		TypeDefinition      respjson.Field
		AutomationFieldType respjson.Field
		SupportedValueTypes respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InputFieldDefinition) RawJSON() string { return r.JSON.raw }
func (r *InputFieldDefinition) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this InputFieldDefinition to a InputFieldDefinitionParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// InputFieldDefinitionParam.Overrides()
func (r InputFieldDefinition) ToParam() InputFieldDefinitionParam {
	return param.Override[InputFieldDefinitionParam](json.RawMessage(r.RawJSON()))
}

// The properties IsRequired, TypeDefinition are required.
type InputFieldDefinitionParam struct {
	IsRequired          bool                     `json:"isRequired,required"`
	TypeDefinition      FieldTypeDefinitionParam `json:"typeDefinition,omitzero,required"`
	AutomationFieldType param.Opt[string]        `json:"automationFieldType,omitzero"`
	// Any of "STATIC_VALUE", "OBJECT_PROPERTY", "FIELD_DATA",
	// "FETCHED_OBJECT_PROPERTY", "ENROLLMENT_EVENT_PROPERTY".
	SupportedValueTypes []string `json:"supportedValueTypes,omitzero"`
	paramObj
}

func (r InputFieldDefinitionParam) MarshalJSON() (data []byte, err error) {
	type shadow InputFieldDefinitionParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InputFieldDefinitionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OutputFieldDefinition struct {
	TypeDefinition FieldTypeDefinition `json:"typeDefinition,required"`
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
	TypeDefinition FieldTypeDefinitionParam `json:"typeDefinition,omitzero,required"`
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
	ID                     string                                            `json:"id,required"`
	ActionURL              string                                            `json:"actionUrl,required"`
	Functions              []PublicActionFunctionIdentifier                  `json:"functions,required"`
	InputFields            []InputFieldDefinition                            `json:"inputFields,required"`
	Labels                 map[string]PublicActionLabels                     `json:"labels,required"`
	ObjectTypes            []string                                          `json:"objectTypes,required"`
	Published              bool                                              `json:"published,required"`
	RevisionID             string                                            `json:"revisionId,required"`
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
	ActionURL              string                                                    `json:"actionUrl,required"`
	Functions              []PublicActionFunctionParam                               `json:"functions,omitzero,required"`
	InputFields            []InputFieldDefinitionParam                               `json:"inputFields,omitzero,required"`
	Labels                 map[string]PublicActionLabelsParam                        `json:"labels,omitzero,required"`
	ObjectTypes            []string                                                  `json:"objectTypes,omitzero,required"`
	Published              bool                                                      `json:"published,required"`
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
	InputFields            []InputFieldDefinitionParam                                 `json:"inputFields,omitzero"`
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

type PublicActionFunction struct {
	FunctionSource string `json:"functionSource,required"`
	// Any of "POST_ACTION_EXECUTION", "POST_FETCH_OPTIONS", "PRE_ACTION_EXECUTION",
	// "PRE_FETCH_OPTIONS".
	FunctionType PublicActionFunctionFunctionType `json:"functionType,required"`
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
	FunctionSource string `json:"functionSource,required"`
	// Any of "POST_ACTION_EXECUTION", "POST_FETCH_OPTIONS", "PRE_ACTION_EXECUTION",
	// "PRE_FETCH_OPTIONS".
	FunctionType PublicActionFunctionFunctionType `json:"functionType,omitzero,required"`
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
	FunctionType PublicActionFunctionIdentifierFunctionType `json:"functionType,required"`
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
	ActionName             string                       `json:"actionName,required"`
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
	ActionName             string                       `json:"actionName,required"`
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
	ID         string                 `json:"id,required"`
	CreatedAt  time.Time              `json:"createdAt,required" format:"date-time"`
	Definition PublicActionDefinition `json:"definition,required"`
	RevisionID string                 `json:"revisionId,required"`
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
	ControllingFieldName  string `json:"controllingFieldName,required"`
	ControllingFieldValue string `json:"controllingFieldValue,required"`
	// Any of "CONDITIONAL_SINGLE_FIELD".
	DependencyType      PublicConditionalSingleFieldDependencyDependencyType `json:"dependencyType,required"`
	DependentFieldNames []string                                             `json:"dependentFieldNames,required"`
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
	ControllingFieldName  string `json:"controllingFieldName,required"`
	ControllingFieldValue string `json:"controllingFieldValue,required"`
	// Any of "CONDITIONAL_SINGLE_FIELD".
	DependencyType      PublicConditionalSingleFieldDependencyDependencyType `json:"dependencyType,omitzero,required"`
	DependentFieldNames []string                                             `json:"dependentFieldNames,omitzero,required"`
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
	Conditions map[string]any `json:"conditions,required"`
	LabelName  string         `json:"labelName,required"`
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
	Conditions map[string]any `json:"conditions,omitzero,required"`
	LabelName  string         `json:"labelName,required"`
	paramObj
}

func (r PublicExecutionTranslationRuleParam) MarshalJSON() (data []byte, err error) {
	type shadow PublicExecutionTranslationRuleParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PublicExecutionTranslationRuleParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PublicObjectRequestOptions struct {
	Properties []string `json:"properties,required"`
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
	Properties []string `json:"properties,omitzero,required"`
	paramObj
}

func (r PublicObjectRequestOptionsParam) MarshalJSON() (data []byte, err error) {
	type shadow PublicObjectRequestOptionsParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PublicObjectRequestOptionsParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PublicSingleFieldDependency struct {
	ControllingFieldName string `json:"controllingFieldName,required"`
	// Any of "SINGLE_FIELD".
	DependencyType      PublicSingleFieldDependencyDependencyType `json:"dependencyType,required"`
	DependentFieldNames []string                                  `json:"dependentFieldNames,required"`
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
	ControllingFieldName string `json:"controllingFieldName,required"`
	// Any of "SINGLE_FIELD".
	DependencyType      PublicSingleFieldDependencyDependencyType `json:"dependencyType,omitzero,required"`
	DependentFieldNames []string                                  `json:"dependentFieldNames,omitzero,required"`
	paramObj
}

func (r PublicSingleFieldDependencyParam) MarshalJSON() (data []byte, err error) {
	type shadow PublicSingleFieldDependencyParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PublicSingleFieldDependencyParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
