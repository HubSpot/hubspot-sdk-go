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
// with the Hubspot API.
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
	// Any of "string", "number", "bool", "datetime", "enumeration", "date",
	// "phone_number", "currency_number", "json", "object_coordinates".
	Type                         FieldTypeDefinitionType `json:"type,required"`
	Description                  string                  `json:"description"`
	ExternalOptionsReferenceType string                  `json:"externalOptionsReferenceType"`
	// Any of "booleancheckbox", "checkbox", "date", "file", "number", "phonenumber",
	// "radio", "select", "text", "textarea", "calculation_equation",
	// "calculation_rollup", "calculation_score", "calculation_read_time", "unknown",
	// "html".
	FieldType  FieldTypeDefinitionFieldType `json:"fieldType"`
	HelpText   string                       `json:"helpText"`
	Label      string                       `json:"label"`
	OptionsURL string                       `json:"optionsUrl"`
	// Any of "CONTACT", "COMPANY", "DEAL", "ENGAGEMENT", "TICKET", "OWNER", "PRODUCT",
	// "LINE_ITEM", "BET_DELIVERABLE_SERVICE", "CONTENT", "CONVERSATION", "BET_ALERT",
	// "PORTAL", "QUOTE", "FORM_SUBMISSION_INBOUNDDB", "QUOTA", "UNSUBSCRIBE",
	// "COMMUNICATION", "FEEDBACK_SUBMISSION", "ATTRIBUTION", "SALESFORCE_SYNC_ERROR",
	// "RESTORABLE_CRM_OBJECT", "HUB", "LANDING_PAGE", "PRODUCT_OR_FOLDER", "TASK",
	// "FORM", "MARKETING_EMAIL", "AD_ACCOUNT", "AD_CAMPAIGN", "AD_GROUP", "AD",
	// "KEYWORD", "CAMPAIGN", "SOCIAL_CHANNEL", "SOCIAL_POST", "SITE_PAGE",
	// "BLOG_POST", "IMPORT", "EXPORT", "CTA", "TASK_TEMPLATE",
	// "AUTOMATION_PLATFORM_FLOW", "OBJECT_LIST", "NOTE", "MEETING_EVENT", "CALL",
	// "EMAIL", "PUBLISHING_TASK", "CONVERSATION_SESSION",
	// "CONTACT_CREATE_ATTRIBUTION", "INVOICE", "MARKETING_EVENT",
	// "CONVERSATION_INBOX", "CHATFLOW", "MEDIA_BRIDGE", "SEQUENCE", "SEQUENCE_STEP",
	// "FORECAST", "SNIPPET", "TEMPLATE", "DEAL_CREATE_ATTRIBUTION", "QUOTE_TEMPLATE",
	// "QUOTE_MODULE", "QUOTE_MODULE_FIELD", "QUOTE_FIELD", "SEQUENCE_ENROLLMENT",
	// "SUBSCRIPTION", "ACCEPTANCE_TEST", "SOCIAL_BROADCAST", "DEAL_SPLIT",
	// "DEAL_REGISTRATION", "GOAL_TARGET", "GOAL_TARGET_GROUP",
	// "PORTAL_OBJECT_SYNC_MESSAGE", "FILE_MANAGER_FILE", "FILE_MANAGER_FOLDER",
	// "SEQUENCE_STEP_ENROLLMENT", "APPROVAL", "APPROVAL_STEP", "CTA_VARIANT",
	// "SALES_DOCUMENT", "DISCOUNT", "FEE", "TAX", "MARKETING_CALENDAR",
	// "PERMISSIONS_TESTING", "PRIVACY_SCANNER_COOKIE", "DATA_SYNC_STATE",
	// "WEB_INTERACTIVE", "PLAYBOOK", "FOLDER", "PLAYBOOK_QUESTION",
	// "PLAYBOOK_SUBMISSION", "PLAYBOOK_SUBMISSION_ANSWER", "COMMERCE_PAYMENT",
	// "GSC_PROPERTY", "SOX_PROTECTED_DUMMY_TYPE", "BLOG_LISTING_PAGE",
	// "QUARANTINED_SUBMISSION", "PAYMENT_SCHEDULE", "PAYMENT_SCHEDULE_INSTALLMENT",
	// "MARKETING_CAMPAIGN_UTM", "DISCOUNT_TEMPLATE", "DISCOUNT_CODE",
	// "FEEDBACK_SURVEY", "CMS_URL", "SALES_TASK", "SALES_WORKLOAD", "USER",
	// "POSTAL_MAIL", "SCHEMAS_BACKEND_TEST", "PAYMENT_LINK", "SUBMISSION_TAG",
	// "CAMPAIGN_STEP", "SCHEDULING_PAGE", "SOX_PROTECTED_TEST_TYPE", "ORDER",
	// "MARKETING_SMS", "PARTNER_ACCOUNT", "CAMPAIGN_TEMPLATE",
	// "CAMPAIGN_TEMPLATE_STEP", "PLAYLIST", "CLIP", "CAMPAIGN_BUDGET_ITEM",
	// "CAMPAIGN_SPEND_ITEM", "MIC", "CONTENT_AUDIT", "CONTENT_AUDIT_PAGE",
	// "PLAYLIST_FOLDER", "LEAD", "ABANDONED_CART", "EXTERNAL_WEB_URL", "VIEW",
	// "VIEW_BLOCK", "ROSTER", "CART", "AUTOMATION_PLATFORM_FLOW_ACTION",
	// "SOCIAL_PROFILE", "PARTNER_CLIENT", "ROSTER_MEMBER",
	// "MARKETING_EVENT_ATTENDANCE", "ALL_PAGES", "AI_FORECAST",
	// "CRM_PIPELINES_DUMMY_TYPE", "KNOWLEDGE_ARTICLE", "PROPERTY_INFO",
	// "DATA_PRIVACY_CONSENT", "GOAL_TEMPLATE", "SCORE_CONFIGURATION", "AUDIENCE",
	// "PARTNER_CLIENT_REVENUE", "AUTOMATION_JOURNEY", "UNKNOWN".
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
	FieldTypeDefinitionTypeString            FieldTypeDefinitionType = "string"
	FieldTypeDefinitionTypeNumber            FieldTypeDefinitionType = "number"
	FieldTypeDefinitionTypeBool              FieldTypeDefinitionType = "bool"
	FieldTypeDefinitionTypeDatetime          FieldTypeDefinitionType = "datetime"
	FieldTypeDefinitionTypeEnumeration       FieldTypeDefinitionType = "enumeration"
	FieldTypeDefinitionTypeDate              FieldTypeDefinitionType = "date"
	FieldTypeDefinitionTypePhoneNumber       FieldTypeDefinitionType = "phone_number"
	FieldTypeDefinitionTypeCurrencyNumber    FieldTypeDefinitionType = "currency_number"
	FieldTypeDefinitionTypeJson              FieldTypeDefinitionType = "json"
	FieldTypeDefinitionTypeObjectCoordinates FieldTypeDefinitionType = "object_coordinates"
)

type FieldTypeDefinitionFieldType string

const (
	FieldTypeDefinitionFieldTypeBooleancheckbox     FieldTypeDefinitionFieldType = "booleancheckbox"
	FieldTypeDefinitionFieldTypeCheckbox            FieldTypeDefinitionFieldType = "checkbox"
	FieldTypeDefinitionFieldTypeDate                FieldTypeDefinitionFieldType = "date"
	FieldTypeDefinitionFieldTypeFile                FieldTypeDefinitionFieldType = "file"
	FieldTypeDefinitionFieldTypeNumber              FieldTypeDefinitionFieldType = "number"
	FieldTypeDefinitionFieldTypePhonenumber         FieldTypeDefinitionFieldType = "phonenumber"
	FieldTypeDefinitionFieldTypeRadio               FieldTypeDefinitionFieldType = "radio"
	FieldTypeDefinitionFieldTypeSelect              FieldTypeDefinitionFieldType = "select"
	FieldTypeDefinitionFieldTypeText                FieldTypeDefinitionFieldType = "text"
	FieldTypeDefinitionFieldTypeTextarea            FieldTypeDefinitionFieldType = "textarea"
	FieldTypeDefinitionFieldTypeCalculationEquation FieldTypeDefinitionFieldType = "calculation_equation"
	FieldTypeDefinitionFieldTypeCalculationRollup   FieldTypeDefinitionFieldType = "calculation_rollup"
	FieldTypeDefinitionFieldTypeCalculationScore    FieldTypeDefinitionFieldType = "calculation_score"
	FieldTypeDefinitionFieldTypeCalculationReadTime FieldTypeDefinitionFieldType = "calculation_read_time"
	FieldTypeDefinitionFieldTypeUnknown             FieldTypeDefinitionFieldType = "unknown"
	FieldTypeDefinitionFieldTypeHTML                FieldTypeDefinitionFieldType = "html"
)

type FieldTypeDefinitionReferencedObjectType string

const (
	FieldTypeDefinitionReferencedObjectTypeContact                      FieldTypeDefinitionReferencedObjectType = "CONTACT"
	FieldTypeDefinitionReferencedObjectTypeCompany                      FieldTypeDefinitionReferencedObjectType = "COMPANY"
	FieldTypeDefinitionReferencedObjectTypeDeal                         FieldTypeDefinitionReferencedObjectType = "DEAL"
	FieldTypeDefinitionReferencedObjectTypeEngagement                   FieldTypeDefinitionReferencedObjectType = "ENGAGEMENT"
	FieldTypeDefinitionReferencedObjectTypeTicket                       FieldTypeDefinitionReferencedObjectType = "TICKET"
	FieldTypeDefinitionReferencedObjectTypeOwner                        FieldTypeDefinitionReferencedObjectType = "OWNER"
	FieldTypeDefinitionReferencedObjectTypeProduct                      FieldTypeDefinitionReferencedObjectType = "PRODUCT"
	FieldTypeDefinitionReferencedObjectTypeLineItem                     FieldTypeDefinitionReferencedObjectType = "LINE_ITEM"
	FieldTypeDefinitionReferencedObjectTypeBetDeliverableService        FieldTypeDefinitionReferencedObjectType = "BET_DELIVERABLE_SERVICE"
	FieldTypeDefinitionReferencedObjectTypeContent                      FieldTypeDefinitionReferencedObjectType = "CONTENT"
	FieldTypeDefinitionReferencedObjectTypeConversation                 FieldTypeDefinitionReferencedObjectType = "CONVERSATION"
	FieldTypeDefinitionReferencedObjectTypeBetAlert                     FieldTypeDefinitionReferencedObjectType = "BET_ALERT"
	FieldTypeDefinitionReferencedObjectTypePortal                       FieldTypeDefinitionReferencedObjectType = "PORTAL"
	FieldTypeDefinitionReferencedObjectTypeQuote                        FieldTypeDefinitionReferencedObjectType = "QUOTE"
	FieldTypeDefinitionReferencedObjectTypeFormSubmissionInbounddb      FieldTypeDefinitionReferencedObjectType = "FORM_SUBMISSION_INBOUNDDB"
	FieldTypeDefinitionReferencedObjectTypeQuota                        FieldTypeDefinitionReferencedObjectType = "QUOTA"
	FieldTypeDefinitionReferencedObjectTypeUnsubscribe                  FieldTypeDefinitionReferencedObjectType = "UNSUBSCRIBE"
	FieldTypeDefinitionReferencedObjectTypeCommunication                FieldTypeDefinitionReferencedObjectType = "COMMUNICATION"
	FieldTypeDefinitionReferencedObjectTypeFeedbackSubmission           FieldTypeDefinitionReferencedObjectType = "FEEDBACK_SUBMISSION"
	FieldTypeDefinitionReferencedObjectTypeAttribution                  FieldTypeDefinitionReferencedObjectType = "ATTRIBUTION"
	FieldTypeDefinitionReferencedObjectTypeSalesforceSyncError          FieldTypeDefinitionReferencedObjectType = "SALESFORCE_SYNC_ERROR"
	FieldTypeDefinitionReferencedObjectTypeRestorableCRMObject          FieldTypeDefinitionReferencedObjectType = "RESTORABLE_CRM_OBJECT"
	FieldTypeDefinitionReferencedObjectTypeHub                          FieldTypeDefinitionReferencedObjectType = "HUB"
	FieldTypeDefinitionReferencedObjectTypeLandingPage                  FieldTypeDefinitionReferencedObjectType = "LANDING_PAGE"
	FieldTypeDefinitionReferencedObjectTypeProductOrFolder              FieldTypeDefinitionReferencedObjectType = "PRODUCT_OR_FOLDER"
	FieldTypeDefinitionReferencedObjectTypeTask                         FieldTypeDefinitionReferencedObjectType = "TASK"
	FieldTypeDefinitionReferencedObjectTypeForm                         FieldTypeDefinitionReferencedObjectType = "FORM"
	FieldTypeDefinitionReferencedObjectTypeMarketingEmail               FieldTypeDefinitionReferencedObjectType = "MARKETING_EMAIL"
	FieldTypeDefinitionReferencedObjectTypeAdAccount                    FieldTypeDefinitionReferencedObjectType = "AD_ACCOUNT"
	FieldTypeDefinitionReferencedObjectTypeAdCampaign                   FieldTypeDefinitionReferencedObjectType = "AD_CAMPAIGN"
	FieldTypeDefinitionReferencedObjectTypeAdGroup                      FieldTypeDefinitionReferencedObjectType = "AD_GROUP"
	FieldTypeDefinitionReferencedObjectTypeAd                           FieldTypeDefinitionReferencedObjectType = "AD"
	FieldTypeDefinitionReferencedObjectTypeKeyword                      FieldTypeDefinitionReferencedObjectType = "KEYWORD"
	FieldTypeDefinitionReferencedObjectTypeCampaign                     FieldTypeDefinitionReferencedObjectType = "CAMPAIGN"
	FieldTypeDefinitionReferencedObjectTypeSocialChannel                FieldTypeDefinitionReferencedObjectType = "SOCIAL_CHANNEL"
	FieldTypeDefinitionReferencedObjectTypeSocialPost                   FieldTypeDefinitionReferencedObjectType = "SOCIAL_POST"
	FieldTypeDefinitionReferencedObjectTypeSitePage                     FieldTypeDefinitionReferencedObjectType = "SITE_PAGE"
	FieldTypeDefinitionReferencedObjectTypeBlogPost                     FieldTypeDefinitionReferencedObjectType = "BLOG_POST"
	FieldTypeDefinitionReferencedObjectTypeImport                       FieldTypeDefinitionReferencedObjectType = "IMPORT"
	FieldTypeDefinitionReferencedObjectTypeExport                       FieldTypeDefinitionReferencedObjectType = "EXPORT"
	FieldTypeDefinitionReferencedObjectTypeCta                          FieldTypeDefinitionReferencedObjectType = "CTA"
	FieldTypeDefinitionReferencedObjectTypeTaskTemplate                 FieldTypeDefinitionReferencedObjectType = "TASK_TEMPLATE"
	FieldTypeDefinitionReferencedObjectTypeAutomationPlatformFlow       FieldTypeDefinitionReferencedObjectType = "AUTOMATION_PLATFORM_FLOW"
	FieldTypeDefinitionReferencedObjectTypeObjectList                   FieldTypeDefinitionReferencedObjectType = "OBJECT_LIST"
	FieldTypeDefinitionReferencedObjectTypeNote                         FieldTypeDefinitionReferencedObjectType = "NOTE"
	FieldTypeDefinitionReferencedObjectTypeMeetingEvent                 FieldTypeDefinitionReferencedObjectType = "MEETING_EVENT"
	FieldTypeDefinitionReferencedObjectTypeCall                         FieldTypeDefinitionReferencedObjectType = "CALL"
	FieldTypeDefinitionReferencedObjectTypeEmail                        FieldTypeDefinitionReferencedObjectType = "EMAIL"
	FieldTypeDefinitionReferencedObjectTypePublishingTask               FieldTypeDefinitionReferencedObjectType = "PUBLISHING_TASK"
	FieldTypeDefinitionReferencedObjectTypeConversationSession          FieldTypeDefinitionReferencedObjectType = "CONVERSATION_SESSION"
	FieldTypeDefinitionReferencedObjectTypeContactCreateAttribution     FieldTypeDefinitionReferencedObjectType = "CONTACT_CREATE_ATTRIBUTION"
	FieldTypeDefinitionReferencedObjectTypeInvoice                      FieldTypeDefinitionReferencedObjectType = "INVOICE"
	FieldTypeDefinitionReferencedObjectTypeMarketingEvent               FieldTypeDefinitionReferencedObjectType = "MARKETING_EVENT"
	FieldTypeDefinitionReferencedObjectTypeConversationInbox            FieldTypeDefinitionReferencedObjectType = "CONVERSATION_INBOX"
	FieldTypeDefinitionReferencedObjectTypeChatflow                     FieldTypeDefinitionReferencedObjectType = "CHATFLOW"
	FieldTypeDefinitionReferencedObjectTypeMediaBridge                  FieldTypeDefinitionReferencedObjectType = "MEDIA_BRIDGE"
	FieldTypeDefinitionReferencedObjectTypeSequence                     FieldTypeDefinitionReferencedObjectType = "SEQUENCE"
	FieldTypeDefinitionReferencedObjectTypeSequenceStep                 FieldTypeDefinitionReferencedObjectType = "SEQUENCE_STEP"
	FieldTypeDefinitionReferencedObjectTypeForecast                     FieldTypeDefinitionReferencedObjectType = "FORECAST"
	FieldTypeDefinitionReferencedObjectTypeSnippet                      FieldTypeDefinitionReferencedObjectType = "SNIPPET"
	FieldTypeDefinitionReferencedObjectTypeTemplate                     FieldTypeDefinitionReferencedObjectType = "TEMPLATE"
	FieldTypeDefinitionReferencedObjectTypeDealCreateAttribution        FieldTypeDefinitionReferencedObjectType = "DEAL_CREATE_ATTRIBUTION"
	FieldTypeDefinitionReferencedObjectTypeQuoteTemplate                FieldTypeDefinitionReferencedObjectType = "QUOTE_TEMPLATE"
	FieldTypeDefinitionReferencedObjectTypeQuoteModule                  FieldTypeDefinitionReferencedObjectType = "QUOTE_MODULE"
	FieldTypeDefinitionReferencedObjectTypeQuoteModuleField             FieldTypeDefinitionReferencedObjectType = "QUOTE_MODULE_FIELD"
	FieldTypeDefinitionReferencedObjectTypeQuoteField                   FieldTypeDefinitionReferencedObjectType = "QUOTE_FIELD"
	FieldTypeDefinitionReferencedObjectTypeSequenceEnrollment           FieldTypeDefinitionReferencedObjectType = "SEQUENCE_ENROLLMENT"
	FieldTypeDefinitionReferencedObjectTypeSubscription                 FieldTypeDefinitionReferencedObjectType = "SUBSCRIPTION"
	FieldTypeDefinitionReferencedObjectTypeAcceptanceTest               FieldTypeDefinitionReferencedObjectType = "ACCEPTANCE_TEST"
	FieldTypeDefinitionReferencedObjectTypeSocialBroadcast              FieldTypeDefinitionReferencedObjectType = "SOCIAL_BROADCAST"
	FieldTypeDefinitionReferencedObjectTypeDealSplit                    FieldTypeDefinitionReferencedObjectType = "DEAL_SPLIT"
	FieldTypeDefinitionReferencedObjectTypeDealRegistration             FieldTypeDefinitionReferencedObjectType = "DEAL_REGISTRATION"
	FieldTypeDefinitionReferencedObjectTypeGoalTarget                   FieldTypeDefinitionReferencedObjectType = "GOAL_TARGET"
	FieldTypeDefinitionReferencedObjectTypeGoalTargetGroup              FieldTypeDefinitionReferencedObjectType = "GOAL_TARGET_GROUP"
	FieldTypeDefinitionReferencedObjectTypePortalObjectSyncMessage      FieldTypeDefinitionReferencedObjectType = "PORTAL_OBJECT_SYNC_MESSAGE"
	FieldTypeDefinitionReferencedObjectTypeFileManagerFile              FieldTypeDefinitionReferencedObjectType = "FILE_MANAGER_FILE"
	FieldTypeDefinitionReferencedObjectTypeFileManagerFolder            FieldTypeDefinitionReferencedObjectType = "FILE_MANAGER_FOLDER"
	FieldTypeDefinitionReferencedObjectTypeSequenceStepEnrollment       FieldTypeDefinitionReferencedObjectType = "SEQUENCE_STEP_ENROLLMENT"
	FieldTypeDefinitionReferencedObjectTypeApproval                     FieldTypeDefinitionReferencedObjectType = "APPROVAL"
	FieldTypeDefinitionReferencedObjectTypeApprovalStep                 FieldTypeDefinitionReferencedObjectType = "APPROVAL_STEP"
	FieldTypeDefinitionReferencedObjectTypeCtaVariant                   FieldTypeDefinitionReferencedObjectType = "CTA_VARIANT"
	FieldTypeDefinitionReferencedObjectTypeSalesDocument                FieldTypeDefinitionReferencedObjectType = "SALES_DOCUMENT"
	FieldTypeDefinitionReferencedObjectTypeDiscount                     FieldTypeDefinitionReferencedObjectType = "DISCOUNT"
	FieldTypeDefinitionReferencedObjectTypeFee                          FieldTypeDefinitionReferencedObjectType = "FEE"
	FieldTypeDefinitionReferencedObjectTypeTax                          FieldTypeDefinitionReferencedObjectType = "TAX"
	FieldTypeDefinitionReferencedObjectTypeMarketingCalendar            FieldTypeDefinitionReferencedObjectType = "MARKETING_CALENDAR"
	FieldTypeDefinitionReferencedObjectTypePermissionsTesting           FieldTypeDefinitionReferencedObjectType = "PERMISSIONS_TESTING"
	FieldTypeDefinitionReferencedObjectTypePrivacyScannerCookie         FieldTypeDefinitionReferencedObjectType = "PRIVACY_SCANNER_COOKIE"
	FieldTypeDefinitionReferencedObjectTypeDataSyncState                FieldTypeDefinitionReferencedObjectType = "DATA_SYNC_STATE"
	FieldTypeDefinitionReferencedObjectTypeWebInteractive               FieldTypeDefinitionReferencedObjectType = "WEB_INTERACTIVE"
	FieldTypeDefinitionReferencedObjectTypePlaybook                     FieldTypeDefinitionReferencedObjectType = "PLAYBOOK"
	FieldTypeDefinitionReferencedObjectTypeFolder                       FieldTypeDefinitionReferencedObjectType = "FOLDER"
	FieldTypeDefinitionReferencedObjectTypePlaybookQuestion             FieldTypeDefinitionReferencedObjectType = "PLAYBOOK_QUESTION"
	FieldTypeDefinitionReferencedObjectTypePlaybookSubmission           FieldTypeDefinitionReferencedObjectType = "PLAYBOOK_SUBMISSION"
	FieldTypeDefinitionReferencedObjectTypePlaybookSubmissionAnswer     FieldTypeDefinitionReferencedObjectType = "PLAYBOOK_SUBMISSION_ANSWER"
	FieldTypeDefinitionReferencedObjectTypeCommercePayment              FieldTypeDefinitionReferencedObjectType = "COMMERCE_PAYMENT"
	FieldTypeDefinitionReferencedObjectTypeGscProperty                  FieldTypeDefinitionReferencedObjectType = "GSC_PROPERTY"
	FieldTypeDefinitionReferencedObjectTypeSoxProtectedDummyType        FieldTypeDefinitionReferencedObjectType = "SOX_PROTECTED_DUMMY_TYPE"
	FieldTypeDefinitionReferencedObjectTypeBlogListingPage              FieldTypeDefinitionReferencedObjectType = "BLOG_LISTING_PAGE"
	FieldTypeDefinitionReferencedObjectTypeQuarantinedSubmission        FieldTypeDefinitionReferencedObjectType = "QUARANTINED_SUBMISSION"
	FieldTypeDefinitionReferencedObjectTypePaymentSchedule              FieldTypeDefinitionReferencedObjectType = "PAYMENT_SCHEDULE"
	FieldTypeDefinitionReferencedObjectTypePaymentScheduleInstallment   FieldTypeDefinitionReferencedObjectType = "PAYMENT_SCHEDULE_INSTALLMENT"
	FieldTypeDefinitionReferencedObjectTypeMarketingCampaignUtm         FieldTypeDefinitionReferencedObjectType = "MARKETING_CAMPAIGN_UTM"
	FieldTypeDefinitionReferencedObjectTypeDiscountTemplate             FieldTypeDefinitionReferencedObjectType = "DISCOUNT_TEMPLATE"
	FieldTypeDefinitionReferencedObjectTypeDiscountCode                 FieldTypeDefinitionReferencedObjectType = "DISCOUNT_CODE"
	FieldTypeDefinitionReferencedObjectTypeFeedbackSurvey               FieldTypeDefinitionReferencedObjectType = "FEEDBACK_SURVEY"
	FieldTypeDefinitionReferencedObjectTypeCmsURL                       FieldTypeDefinitionReferencedObjectType = "CMS_URL"
	FieldTypeDefinitionReferencedObjectTypeSalesTask                    FieldTypeDefinitionReferencedObjectType = "SALES_TASK"
	FieldTypeDefinitionReferencedObjectTypeSalesWorkload                FieldTypeDefinitionReferencedObjectType = "SALES_WORKLOAD"
	FieldTypeDefinitionReferencedObjectTypeUser                         FieldTypeDefinitionReferencedObjectType = "USER"
	FieldTypeDefinitionReferencedObjectTypePostalMail                   FieldTypeDefinitionReferencedObjectType = "POSTAL_MAIL"
	FieldTypeDefinitionReferencedObjectTypeSchemasBackendTest           FieldTypeDefinitionReferencedObjectType = "SCHEMAS_BACKEND_TEST"
	FieldTypeDefinitionReferencedObjectTypePaymentLink                  FieldTypeDefinitionReferencedObjectType = "PAYMENT_LINK"
	FieldTypeDefinitionReferencedObjectTypeSubmissionTag                FieldTypeDefinitionReferencedObjectType = "SUBMISSION_TAG"
	FieldTypeDefinitionReferencedObjectTypeCampaignStep                 FieldTypeDefinitionReferencedObjectType = "CAMPAIGN_STEP"
	FieldTypeDefinitionReferencedObjectTypeSchedulingPage               FieldTypeDefinitionReferencedObjectType = "SCHEDULING_PAGE"
	FieldTypeDefinitionReferencedObjectTypeSoxProtectedTestType         FieldTypeDefinitionReferencedObjectType = "SOX_PROTECTED_TEST_TYPE"
	FieldTypeDefinitionReferencedObjectTypeOrder                        FieldTypeDefinitionReferencedObjectType = "ORDER"
	FieldTypeDefinitionReferencedObjectTypeMarketingSMS                 FieldTypeDefinitionReferencedObjectType = "MARKETING_SMS"
	FieldTypeDefinitionReferencedObjectTypePartnerAccount               FieldTypeDefinitionReferencedObjectType = "PARTNER_ACCOUNT"
	FieldTypeDefinitionReferencedObjectTypeCampaignTemplate             FieldTypeDefinitionReferencedObjectType = "CAMPAIGN_TEMPLATE"
	FieldTypeDefinitionReferencedObjectTypeCampaignTemplateStep         FieldTypeDefinitionReferencedObjectType = "CAMPAIGN_TEMPLATE_STEP"
	FieldTypeDefinitionReferencedObjectTypePlaylist                     FieldTypeDefinitionReferencedObjectType = "PLAYLIST"
	FieldTypeDefinitionReferencedObjectTypeClip                         FieldTypeDefinitionReferencedObjectType = "CLIP"
	FieldTypeDefinitionReferencedObjectTypeCampaignBudgetItem           FieldTypeDefinitionReferencedObjectType = "CAMPAIGN_BUDGET_ITEM"
	FieldTypeDefinitionReferencedObjectTypeCampaignSpendItem            FieldTypeDefinitionReferencedObjectType = "CAMPAIGN_SPEND_ITEM"
	FieldTypeDefinitionReferencedObjectTypeMic                          FieldTypeDefinitionReferencedObjectType = "MIC"
	FieldTypeDefinitionReferencedObjectTypeContentAudit                 FieldTypeDefinitionReferencedObjectType = "CONTENT_AUDIT"
	FieldTypeDefinitionReferencedObjectTypeContentAuditPage             FieldTypeDefinitionReferencedObjectType = "CONTENT_AUDIT_PAGE"
	FieldTypeDefinitionReferencedObjectTypePlaylistFolder               FieldTypeDefinitionReferencedObjectType = "PLAYLIST_FOLDER"
	FieldTypeDefinitionReferencedObjectTypeLead                         FieldTypeDefinitionReferencedObjectType = "LEAD"
	FieldTypeDefinitionReferencedObjectTypeAbandonedCart                FieldTypeDefinitionReferencedObjectType = "ABANDONED_CART"
	FieldTypeDefinitionReferencedObjectTypeExternalWebURL               FieldTypeDefinitionReferencedObjectType = "EXTERNAL_WEB_URL"
	FieldTypeDefinitionReferencedObjectTypeView                         FieldTypeDefinitionReferencedObjectType = "VIEW"
	FieldTypeDefinitionReferencedObjectTypeViewBlock                    FieldTypeDefinitionReferencedObjectType = "VIEW_BLOCK"
	FieldTypeDefinitionReferencedObjectTypeRoster                       FieldTypeDefinitionReferencedObjectType = "ROSTER"
	FieldTypeDefinitionReferencedObjectTypeCart                         FieldTypeDefinitionReferencedObjectType = "CART"
	FieldTypeDefinitionReferencedObjectTypeAutomationPlatformFlowAction FieldTypeDefinitionReferencedObjectType = "AUTOMATION_PLATFORM_FLOW_ACTION"
	FieldTypeDefinitionReferencedObjectTypeSocialProfile                FieldTypeDefinitionReferencedObjectType = "SOCIAL_PROFILE"
	FieldTypeDefinitionReferencedObjectTypePartnerClient                FieldTypeDefinitionReferencedObjectType = "PARTNER_CLIENT"
	FieldTypeDefinitionReferencedObjectTypeRosterMember                 FieldTypeDefinitionReferencedObjectType = "ROSTER_MEMBER"
	FieldTypeDefinitionReferencedObjectTypeMarketingEventAttendance     FieldTypeDefinitionReferencedObjectType = "MARKETING_EVENT_ATTENDANCE"
	FieldTypeDefinitionReferencedObjectTypeAllPages                     FieldTypeDefinitionReferencedObjectType = "ALL_PAGES"
	FieldTypeDefinitionReferencedObjectTypeAIForecast                   FieldTypeDefinitionReferencedObjectType = "AI_FORECAST"
	FieldTypeDefinitionReferencedObjectTypeCRMPipelinesDummyType        FieldTypeDefinitionReferencedObjectType = "CRM_PIPELINES_DUMMY_TYPE"
	FieldTypeDefinitionReferencedObjectTypeKnowledgeArticle             FieldTypeDefinitionReferencedObjectType = "KNOWLEDGE_ARTICLE"
	FieldTypeDefinitionReferencedObjectTypePropertyInfo                 FieldTypeDefinitionReferencedObjectType = "PROPERTY_INFO"
	FieldTypeDefinitionReferencedObjectTypeDataPrivacyConsent           FieldTypeDefinitionReferencedObjectType = "DATA_PRIVACY_CONSENT"
	FieldTypeDefinitionReferencedObjectTypeGoalTemplate                 FieldTypeDefinitionReferencedObjectType = "GOAL_TEMPLATE"
	FieldTypeDefinitionReferencedObjectTypeScoreConfiguration           FieldTypeDefinitionReferencedObjectType = "SCORE_CONFIGURATION"
	FieldTypeDefinitionReferencedObjectTypeAudience                     FieldTypeDefinitionReferencedObjectType = "AUDIENCE"
	FieldTypeDefinitionReferencedObjectTypePartnerClientRevenue         FieldTypeDefinitionReferencedObjectType = "PARTNER_CLIENT_REVENUE"
	FieldTypeDefinitionReferencedObjectTypeAutomationJourney            FieldTypeDefinitionReferencedObjectType = "AUTOMATION_JOURNEY"
	FieldTypeDefinitionReferencedObjectTypeUnknown                      FieldTypeDefinitionReferencedObjectType = "UNKNOWN"
)

// The properties ExternalOptions, Name, Options, Type are required.
type FieldTypeDefinitionParam struct {
	ExternalOptions bool                 `json:"externalOptions,required"`
	Name            string               `json:"name,required"`
	Options         []shared.OptionParam `json:"options,omitzero,required"`
	// Any of "string", "number", "bool", "datetime", "enumeration", "date",
	// "phone_number", "currency_number", "json", "object_coordinates".
	Type                         FieldTypeDefinitionType `json:"type,omitzero,required"`
	Description                  param.Opt[string]       `json:"description,omitzero"`
	ExternalOptionsReferenceType param.Opt[string]       `json:"externalOptionsReferenceType,omitzero"`
	HelpText                     param.Opt[string]       `json:"helpText,omitzero"`
	Label                        param.Opt[string]       `json:"label,omitzero"`
	OptionsURL                   param.Opt[string]       `json:"optionsUrl,omitzero"`
	// Any of "booleancheckbox", "checkbox", "date", "file", "number", "phonenumber",
	// "radio", "select", "text", "textarea", "calculation_equation",
	// "calculation_rollup", "calculation_score", "calculation_read_time", "unknown",
	// "html".
	FieldType FieldTypeDefinitionFieldType `json:"fieldType,omitzero"`
	// Any of "CONTACT", "COMPANY", "DEAL", "ENGAGEMENT", "TICKET", "OWNER", "PRODUCT",
	// "LINE_ITEM", "BET_DELIVERABLE_SERVICE", "CONTENT", "CONVERSATION", "BET_ALERT",
	// "PORTAL", "QUOTE", "FORM_SUBMISSION_INBOUNDDB", "QUOTA", "UNSUBSCRIBE",
	// "COMMUNICATION", "FEEDBACK_SUBMISSION", "ATTRIBUTION", "SALESFORCE_SYNC_ERROR",
	// "RESTORABLE_CRM_OBJECT", "HUB", "LANDING_PAGE", "PRODUCT_OR_FOLDER", "TASK",
	// "FORM", "MARKETING_EMAIL", "AD_ACCOUNT", "AD_CAMPAIGN", "AD_GROUP", "AD",
	// "KEYWORD", "CAMPAIGN", "SOCIAL_CHANNEL", "SOCIAL_POST", "SITE_PAGE",
	// "BLOG_POST", "IMPORT", "EXPORT", "CTA", "TASK_TEMPLATE",
	// "AUTOMATION_PLATFORM_FLOW", "OBJECT_LIST", "NOTE", "MEETING_EVENT", "CALL",
	// "EMAIL", "PUBLISHING_TASK", "CONVERSATION_SESSION",
	// "CONTACT_CREATE_ATTRIBUTION", "INVOICE", "MARKETING_EVENT",
	// "CONVERSATION_INBOX", "CHATFLOW", "MEDIA_BRIDGE", "SEQUENCE", "SEQUENCE_STEP",
	// "FORECAST", "SNIPPET", "TEMPLATE", "DEAL_CREATE_ATTRIBUTION", "QUOTE_TEMPLATE",
	// "QUOTE_MODULE", "QUOTE_MODULE_FIELD", "QUOTE_FIELD", "SEQUENCE_ENROLLMENT",
	// "SUBSCRIPTION", "ACCEPTANCE_TEST", "SOCIAL_BROADCAST", "DEAL_SPLIT",
	// "DEAL_REGISTRATION", "GOAL_TARGET", "GOAL_TARGET_GROUP",
	// "PORTAL_OBJECT_SYNC_MESSAGE", "FILE_MANAGER_FILE", "FILE_MANAGER_FOLDER",
	// "SEQUENCE_STEP_ENROLLMENT", "APPROVAL", "APPROVAL_STEP", "CTA_VARIANT",
	// "SALES_DOCUMENT", "DISCOUNT", "FEE", "TAX", "MARKETING_CALENDAR",
	// "PERMISSIONS_TESTING", "PRIVACY_SCANNER_COOKIE", "DATA_SYNC_STATE",
	// "WEB_INTERACTIVE", "PLAYBOOK", "FOLDER", "PLAYBOOK_QUESTION",
	// "PLAYBOOK_SUBMISSION", "PLAYBOOK_SUBMISSION_ANSWER", "COMMERCE_PAYMENT",
	// "GSC_PROPERTY", "SOX_PROTECTED_DUMMY_TYPE", "BLOG_LISTING_PAGE",
	// "QUARANTINED_SUBMISSION", "PAYMENT_SCHEDULE", "PAYMENT_SCHEDULE_INSTALLMENT",
	// "MARKETING_CAMPAIGN_UTM", "DISCOUNT_TEMPLATE", "DISCOUNT_CODE",
	// "FEEDBACK_SURVEY", "CMS_URL", "SALES_TASK", "SALES_WORKLOAD", "USER",
	// "POSTAL_MAIL", "SCHEMAS_BACKEND_TEST", "PAYMENT_LINK", "SUBMISSION_TAG",
	// "CAMPAIGN_STEP", "SCHEDULING_PAGE", "SOX_PROTECTED_TEST_TYPE", "ORDER",
	// "MARKETING_SMS", "PARTNER_ACCOUNT", "CAMPAIGN_TEMPLATE",
	// "CAMPAIGN_TEMPLATE_STEP", "PLAYLIST", "CLIP", "CAMPAIGN_BUDGET_ITEM",
	// "CAMPAIGN_SPEND_ITEM", "MIC", "CONTENT_AUDIT", "CONTENT_AUDIT_PAGE",
	// "PLAYLIST_FOLDER", "LEAD", "ABANDONED_CART", "EXTERNAL_WEB_URL", "VIEW",
	// "VIEW_BLOCK", "ROSTER", "CART", "AUTOMATION_PLATFORM_FLOW_ACTION",
	// "SOCIAL_PROFILE", "PARTNER_CLIENT", "ROSTER_MEMBER",
	// "MARKETING_EVENT_ATTENDANCE", "ALL_PAGES", "AI_FORECAST",
	// "CRM_PIPELINES_DUMMY_TYPE", "KNOWLEDGE_ARTICLE", "PROPERTY_INFO",
	// "DATA_PRIVACY_CONSENT", "GOAL_TEMPLATE", "SCORE_CONFIGURATION", "AUDIENCE",
	// "PARTNER_CLIENT_REVENUE", "AUTOMATION_JOURNEY", "UNKNOWN".
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
	// Any of "PRE_ACTION_EXECUTION", "PRE_FETCH_OPTIONS", "POST_FETCH_OPTIONS",
	// "POST_ACTION_EXECUTION".
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
	PublicActionFunctionFunctionTypePreActionExecution  PublicActionFunctionFunctionType = "PRE_ACTION_EXECUTION"
	PublicActionFunctionFunctionTypePreFetchOptions     PublicActionFunctionFunctionType = "PRE_FETCH_OPTIONS"
	PublicActionFunctionFunctionTypePostFetchOptions    PublicActionFunctionFunctionType = "POST_FETCH_OPTIONS"
	PublicActionFunctionFunctionTypePostActionExecution PublicActionFunctionFunctionType = "POST_ACTION_EXECUTION"
)

// The properties FunctionSource, FunctionType are required.
type PublicActionFunctionParam struct {
	FunctionSource string `json:"functionSource,required"`
	// Any of "PRE_ACTION_EXECUTION", "PRE_FETCH_OPTIONS", "POST_FETCH_OPTIONS",
	// "POST_ACTION_EXECUTION".
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
	// Any of "PRE_ACTION_EXECUTION", "PRE_FETCH_OPTIONS", "POST_FETCH_OPTIONS",
	// "POST_ACTION_EXECUTION".
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
	PublicActionFunctionIdentifierFunctionTypePreActionExecution  PublicActionFunctionIdentifierFunctionType = "PRE_ACTION_EXECUTION"
	PublicActionFunctionIdentifierFunctionTypePreFetchOptions     PublicActionFunctionIdentifierFunctionType = "PRE_FETCH_OPTIONS"
	PublicActionFunctionIdentifierFunctionTypePostFetchOptions    PublicActionFunctionIdentifierFunctionType = "POST_FETCH_OPTIONS"
	PublicActionFunctionIdentifierFunctionTypePostActionExecution PublicActionFunctionIdentifierFunctionType = "POST_ACTION_EXECUTION"
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
