// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package crm

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/HubSpot/hubspot-sdk-go/internal/apiform"
	"github.com/HubSpot/hubspot-sdk-go/internal/apijson"
	"github.com/HubSpot/hubspot-sdk-go/internal/apiquery"
	"github.com/HubSpot/hubspot-sdk-go/internal/requestconfig"
	"github.com/HubSpot/hubspot-sdk-go/option"
	"github.com/HubSpot/hubspot-sdk-go/packages/pagination"
	"github.com/HubSpot/hubspot-sdk-go/packages/param"
	"github.com/HubSpot/hubspot-sdk-go/packages/respjson"
	"github.com/HubSpot/hubspot-sdk-go/shared"
)

// ImportService contains methods and other services that help with interacting
// with the hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewImportService] method instead.
type ImportService struct {
	options []option.RequestOption
}

// NewImportService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewImportService(opts ...option.RequestOption) (r ImportService) {
	r = ImportService{}
	r.options = opts
	return
}

func (r *ImportService) New(ctx context.Context, body ImportNewParams, opts ...option.RequestOption) (res *PublicImportResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "crm/imports/2026-03"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

func (r *ImportService) List(ctx context.Context, query ImportListParams, opts ...option.RequestOption) (res *pagination.Page[PublicImportResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "crm/imports/2026-03"
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

func (r *ImportService) ListAutoPaging(ctx context.Context, query ImportListParams, opts ...option.RequestOption) *pagination.PageAutoPager[PublicImportResponse] {
	return pagination.NewPageAutoPager(r.List(ctx, query, opts...))
}

func (r *ImportService) Cancel(ctx context.Context, importID int64, opts ...option.RequestOption) (res *shared.ActionResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := fmt.Sprintf("crm/imports/2026-03/%v/cancel", importID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

func (r *ImportService) Get(ctx context.Context, importID int64, opts ...option.RequestOption) (res *PublicImportResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := fmt.Sprintf("crm/imports/2026-03/%v", importID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

func (r *ImportService) ListErrors(ctx context.Context, importID int64, query ImportListErrorsParams, opts ...option.RequestOption) (res *pagination.Page[PublicImportError], err error) {
	var raw *http.Response
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := fmt.Sprintf("crm/imports/2026-03/%v/errors", importID)
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

func (r *ImportService) ListErrorsAutoPaging(ctx context.Context, importID int64, query ImportListErrorsParams, opts ...option.RequestOption) *pagination.PageAutoPager[PublicImportError] {
	return pagination.NewPageAutoPager(r.ListErrors(ctx, importID, query, opts...))
}

type CollectionResponsePublicImportErrorForwardPaging struct {
	Results []PublicImportError  `json:"results" api:"required"`
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
func (r CollectionResponsePublicImportErrorForwardPaging) RawJSON() string { return r.JSON.raw }
func (r *CollectionResponsePublicImportErrorForwardPaging) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CollectionResponsePublicImportResponseForwardPaging struct {
	Results []PublicImportResponse `json:"results" api:"required"`
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
func (r CollectionResponsePublicImportResponseForwardPaging) RawJSON() string { return r.JSON.raw }
func (r *CollectionResponsePublicImportResponseForwardPaging) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ImportRowCore struct {
	AdditionalRowData []string `json:"additionalRowData" api:"required"`
	// Indicates whether this row contains values that were encrypted.
	ContainsEncryptedProperties bool `json:"containsEncryptedProperties" api:"required"`
	// The unique identifier of the uploaded file containing this row.
	FileID int64 `json:"fileId" api:"required"`
	// The 1-indexed line number of this row in the source file. Line number 0 is
	// reserved for file-wide errors that don't correspond to a specific row.
	LineNumber int64    `json:"lineNumber" api:"required"`
	RowData    []string `json:"rowData" api:"required"`
	// The name of the spreadsheet sheet/page containing this row.
	PageName string `json:"pageName"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AdditionalRowData           respjson.Field
		ContainsEncryptedProperties respjson.Field
		FileID                      respjson.Field
		LineNumber                  respjson.Field
		RowData                     respjson.Field
		PageName                    respjson.Field
		ExtraFields                 map[string]respjson.Field
		raw                         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ImportRowCore) RawJSON() string { return r.JSON.raw }
func (r *ImportRowCore) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ImportTemplate struct {
	// The unique identifier for the specific saved template or previous import being
	// referenced.
	TemplateID int64 `json:"templateId" api:"required"`
	// The classification of what type of template this represents, and what is its
	// origin or purpose.
	//
	// Any of "admin_defined", "previous_import", "user_file".
	TemplateType ImportTemplateTemplateType `json:"templateType" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		TemplateID   respjson.Field
		TemplateType respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ImportTemplate) RawJSON() string { return r.JSON.raw }
func (r *ImportTemplate) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The classification of what type of template this represents, and what is its
// origin or purpose.
type ImportTemplateTemplateType string

const (
	ImportTemplateTemplateTypeAdminDefined   ImportTemplateTemplateType = "admin_defined"
	ImportTemplateTemplateTypePreviousImport ImportTemplateTemplateType = "previous_import"
	ImportTemplateTemplateTypeUserFile       ImportTemplateTemplateType = "user_file"
)

type PublicImportError struct {
	// A unique, stable identifier for this specific error.
	ID string `json:"id" api:"required"`
	// The epoch millisecond timestamp when this error was recorded.
	CreatedAt int64 `json:"createdAt" api:"required"`
	// The classification of what went wrong during import processing.
	//
	// Any of "AMBIGUOUS_ENUMERATION_OPTION", "ASSOCIATION_LABEL_NOT_FOUND",
	// "ASSOCIATION_LIMIT_EXCEEDED", "ASSOCIATION_RECORD_NOT_FOUND",
	// "COLUMN_TOO_LARGE", "COULD_NOT_FIND_BUSINESS_UNIT", "COULD_NOT_FIND_OWNER",
	// "COULD_NOT_PARSE_DATE", "COULD_NOT_PARSE_NUMBER", "COULD_NOT_PARSE_ROW",
	// "COULD_NOT_PARSE_TERM", "CREATE_ONLY_IMPORT", "DUPLICATE_ALTERNATE_ID",
	// "DUPLICATE_ASSOCIATION_ID", "DUPLICATE_EVENT", "DUPLICATE_OBJECT_ID",
	// "DUPLICATE_RECORD_ID", "DUPLICATE_ROW_CONTENT", "DUPLICATE_UNIQUE_CREATION_KEY",
	// "DUPLICATE_UNIQUE_PROPERTY_VALUE", "FAILED_TO_CREATE_ASSOCIATION",
	// "FAILED_TO_FIND_RECORD_FOR_ASSOCIATIONS", "FAILED_TO_OPT_OUT_CONTACT",
	// "FAILED_TO_PROCESS_OBJECT_WITH_EMPTY_PROPERTY_VALUES", "FAILED_VALIDATION",
	// "FILE_NOT_FOUND", "GDPR_BLACKLISTED_EMAIL", "INCORRECT_NUMBER_OF_COLUMNS",
	// "INVALID_ALTERNATE_ID", "INVALID_ASSOCIATION_IDENTIFIER",
	// "INVALID_ASSOCIATION_KEY", "INVALID_COLUMN_CONFIGURATION",
	// "INVALID_CUSTOM_PROPERTY_VALIDATION", "INVALID_DOMAIN", "INVALID_EMAIL",
	// "INVALID_ENUM_FILE_ID_OR_URL", "INVALID_ENUMERATION_OPTION", "INVALID_EVENT",
	// "INVALID_EVENT_TIMESTAMP", "INVALID_FILE_TYPE", "INVALID_NUMBER_SIZE",
	// "INVALID_OBJECT_ID", "INVALID_PROPERTY_VALUE_FORMAT", "INVALID_RECORD_ID",
	// "INVALID_REQUIRED_PROPERTY", "INVALID_SHEET_COUNT", "INVALID_SPREADSHEET",
	// "LIMIT_EXCEEDED", "MANY_ERRORS_IN_ROW", "MISSING_EVENT_DEFINITION",
	// "MISSING_EVENT_TIMESTAMP", "MISSING_OBJECT_DEFINITION",
	// "MISSING_REQUIRED_PROPERTY", "MULTIPLE_COMPANIES_WITH_THIS_DOMAIN",
	// "MULTIPLE_OWNERS_FOUND", "NO_OBJECT_ID_FROM_ASSOCIATION_IDENTIFIER",
	// "OUTSIDE_VALID_TERM_RANGE", "OUTSIDE_VALID_TIME_RANGE",
	// "PORTAL_WIDE_CUSTOM_OBJECT_LIMIT_EXCEEDED", "PROPERTY_DEFINITION_NOT_FOUND",
	// "PROPERTY_VALUE_NOT_FOUND", "ROW_DATA_TOO_LARGE",
	// "SECONDARY_EMAIL_WRITE_FAILURE", "UNKNOWN_ASSOCIATION_RECORD_ID",
	// "UNKNOWN_BAD_REQUEST", "UNKNOWN_ERROR", "UPDATE_ONLY_IMPORT".
	ErrorType  PublicImportErrorErrorType `json:"errorType" api:"required"`
	SourceData ImportRowCore              `json:"sourceData" api:"required"`
	// A human-readable error message.
	ErrorMessage string `json:"errorMessage"`
	// Additional human-readable context about the error.
	ExtraContext string `json:"extraContext"`
	// Represents a single custom property of a marketing event, storing its name,
	// value, metadata (like source, timestamp, and sensitivity), and related audit
	// information for tracking changes.
	InvalidPropertyValue shared.PropertyValue `json:"invalidPropertyValue"`
	// The raw string value from the import file that caused the validation failure.
	InvalidValue string `json:"invalidValue"`
	// A convenience accessor that returns either the value from `invalidPropertyValue`
	// or `invalidValue`, whichever is present (preferring the property value).
	InvalidValueToDisplay string `json:"invalidValueToDisplay"`
	// The zero-based column index in the import file where the error occurred
	KnownColumnNumber int64 `json:"knownColumnNumber"`
	// The CRM object type affected by this error.
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
	ObjectType PublicImportErrorObjectType `json:"objectType"`
	// The modern object type identifier for the CRM object affected by this error.
	ObjectTypeID string `json:"objectTypeId"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                    respjson.Field
		CreatedAt             respjson.Field
		ErrorType             respjson.Field
		SourceData            respjson.Field
		ErrorMessage          respjson.Field
		ExtraContext          respjson.Field
		InvalidPropertyValue  respjson.Field
		InvalidValue          respjson.Field
		InvalidValueToDisplay respjson.Field
		KnownColumnNumber     respjson.Field
		ObjectType            respjson.Field
		ObjectTypeID          respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PublicImportError) RawJSON() string { return r.JSON.raw }
func (r *PublicImportError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The classification of what went wrong during import processing.
type PublicImportErrorErrorType string

const (
	PublicImportErrorErrorTypeAmbiguousEnumerationOption                   PublicImportErrorErrorType = "AMBIGUOUS_ENUMERATION_OPTION"
	PublicImportErrorErrorTypeAssociationLabelNotFound                     PublicImportErrorErrorType = "ASSOCIATION_LABEL_NOT_FOUND"
	PublicImportErrorErrorTypeAssociationLimitExceeded                     PublicImportErrorErrorType = "ASSOCIATION_LIMIT_EXCEEDED"
	PublicImportErrorErrorTypeAssociationRecordNotFound                    PublicImportErrorErrorType = "ASSOCIATION_RECORD_NOT_FOUND"
	PublicImportErrorErrorTypeColumnTooLarge                               PublicImportErrorErrorType = "COLUMN_TOO_LARGE"
	PublicImportErrorErrorTypeCouldNotFindBusinessUnit                     PublicImportErrorErrorType = "COULD_NOT_FIND_BUSINESS_UNIT"
	PublicImportErrorErrorTypeCouldNotFindOwner                            PublicImportErrorErrorType = "COULD_NOT_FIND_OWNER"
	PublicImportErrorErrorTypeCouldNotParseDate                            PublicImportErrorErrorType = "COULD_NOT_PARSE_DATE"
	PublicImportErrorErrorTypeCouldNotParseNumber                          PublicImportErrorErrorType = "COULD_NOT_PARSE_NUMBER"
	PublicImportErrorErrorTypeCouldNotParseRow                             PublicImportErrorErrorType = "COULD_NOT_PARSE_ROW"
	PublicImportErrorErrorTypeCouldNotParseTerm                            PublicImportErrorErrorType = "COULD_NOT_PARSE_TERM"
	PublicImportErrorErrorTypeCreateOnlyImport                             PublicImportErrorErrorType = "CREATE_ONLY_IMPORT"
	PublicImportErrorErrorTypeDuplicateAlternateID                         PublicImportErrorErrorType = "DUPLICATE_ALTERNATE_ID"
	PublicImportErrorErrorTypeDuplicateAssociationID                       PublicImportErrorErrorType = "DUPLICATE_ASSOCIATION_ID"
	PublicImportErrorErrorTypeDuplicateEvent                               PublicImportErrorErrorType = "DUPLICATE_EVENT"
	PublicImportErrorErrorTypeDuplicateObjectID                            PublicImportErrorErrorType = "DUPLICATE_OBJECT_ID"
	PublicImportErrorErrorTypeDuplicateRecordID                            PublicImportErrorErrorType = "DUPLICATE_RECORD_ID"
	PublicImportErrorErrorTypeDuplicateRowContent                          PublicImportErrorErrorType = "DUPLICATE_ROW_CONTENT"
	PublicImportErrorErrorTypeDuplicateUniqueCreationKey                   PublicImportErrorErrorType = "DUPLICATE_UNIQUE_CREATION_KEY"
	PublicImportErrorErrorTypeDuplicateUniquePropertyValue                 PublicImportErrorErrorType = "DUPLICATE_UNIQUE_PROPERTY_VALUE"
	PublicImportErrorErrorTypeFailedToCreateAssociation                    PublicImportErrorErrorType = "FAILED_TO_CREATE_ASSOCIATION"
	PublicImportErrorErrorTypeFailedToFindRecordForAssociations            PublicImportErrorErrorType = "FAILED_TO_FIND_RECORD_FOR_ASSOCIATIONS"
	PublicImportErrorErrorTypeFailedToOptOutContact                        PublicImportErrorErrorType = "FAILED_TO_OPT_OUT_CONTACT"
	PublicImportErrorErrorTypeFailedToProcessObjectWithEmptyPropertyValues PublicImportErrorErrorType = "FAILED_TO_PROCESS_OBJECT_WITH_EMPTY_PROPERTY_VALUES"
	PublicImportErrorErrorTypeFailedValidation                             PublicImportErrorErrorType = "FAILED_VALIDATION"
	PublicImportErrorErrorTypeFileNotFound                                 PublicImportErrorErrorType = "FILE_NOT_FOUND"
	PublicImportErrorErrorTypeGdprBlacklistedEmail                         PublicImportErrorErrorType = "GDPR_BLACKLISTED_EMAIL"
	PublicImportErrorErrorTypeIncorrectNumberOfColumns                     PublicImportErrorErrorType = "INCORRECT_NUMBER_OF_COLUMNS"
	PublicImportErrorErrorTypeInvalidAlternateID                           PublicImportErrorErrorType = "INVALID_ALTERNATE_ID"
	PublicImportErrorErrorTypeInvalidAssociationIdentifier                 PublicImportErrorErrorType = "INVALID_ASSOCIATION_IDENTIFIER"
	PublicImportErrorErrorTypeInvalidAssociationKey                        PublicImportErrorErrorType = "INVALID_ASSOCIATION_KEY"
	PublicImportErrorErrorTypeInvalidColumnConfiguration                   PublicImportErrorErrorType = "INVALID_COLUMN_CONFIGURATION"
	PublicImportErrorErrorTypeInvalidCustomPropertyValidation              PublicImportErrorErrorType = "INVALID_CUSTOM_PROPERTY_VALIDATION"
	PublicImportErrorErrorTypeInvalidDomain                                PublicImportErrorErrorType = "INVALID_DOMAIN"
	PublicImportErrorErrorTypeInvalidEmail                                 PublicImportErrorErrorType = "INVALID_EMAIL"
	PublicImportErrorErrorTypeInvalidEnumFileIDOrURL                       PublicImportErrorErrorType = "INVALID_ENUM_FILE_ID_OR_URL"
	PublicImportErrorErrorTypeInvalidEnumerationOption                     PublicImportErrorErrorType = "INVALID_ENUMERATION_OPTION"
	PublicImportErrorErrorTypeInvalidEvent                                 PublicImportErrorErrorType = "INVALID_EVENT"
	PublicImportErrorErrorTypeInvalidEventTimestamp                        PublicImportErrorErrorType = "INVALID_EVENT_TIMESTAMP"
	PublicImportErrorErrorTypeInvalidFileType                              PublicImportErrorErrorType = "INVALID_FILE_TYPE"
	PublicImportErrorErrorTypeInvalidNumberSize                            PublicImportErrorErrorType = "INVALID_NUMBER_SIZE"
	PublicImportErrorErrorTypeInvalidObjectID                              PublicImportErrorErrorType = "INVALID_OBJECT_ID"
	PublicImportErrorErrorTypeInvalidPropertyValueFormat                   PublicImportErrorErrorType = "INVALID_PROPERTY_VALUE_FORMAT"
	PublicImportErrorErrorTypeInvalidRecordID                              PublicImportErrorErrorType = "INVALID_RECORD_ID"
	PublicImportErrorErrorTypeInvalidRequiredProperty                      PublicImportErrorErrorType = "INVALID_REQUIRED_PROPERTY"
	PublicImportErrorErrorTypeInvalidSheetCount                            PublicImportErrorErrorType = "INVALID_SHEET_COUNT"
	PublicImportErrorErrorTypeInvalidSpreadsheet                           PublicImportErrorErrorType = "INVALID_SPREADSHEET"
	PublicImportErrorErrorTypeLimitExceeded                                PublicImportErrorErrorType = "LIMIT_EXCEEDED"
	PublicImportErrorErrorTypeManyErrorsInRow                              PublicImportErrorErrorType = "MANY_ERRORS_IN_ROW"
	PublicImportErrorErrorTypeMissingEventDefinition                       PublicImportErrorErrorType = "MISSING_EVENT_DEFINITION"
	PublicImportErrorErrorTypeMissingEventTimestamp                        PublicImportErrorErrorType = "MISSING_EVENT_TIMESTAMP"
	PublicImportErrorErrorTypeMissingObjectDefinition                      PublicImportErrorErrorType = "MISSING_OBJECT_DEFINITION"
	PublicImportErrorErrorTypeMissingRequiredProperty                      PublicImportErrorErrorType = "MISSING_REQUIRED_PROPERTY"
	PublicImportErrorErrorTypeMultipleCompaniesWithThisDomain              PublicImportErrorErrorType = "MULTIPLE_COMPANIES_WITH_THIS_DOMAIN"
	PublicImportErrorErrorTypeMultipleOwnersFound                          PublicImportErrorErrorType = "MULTIPLE_OWNERS_FOUND"
	PublicImportErrorErrorTypeNoObjectIDFromAssociationIdentifier          PublicImportErrorErrorType = "NO_OBJECT_ID_FROM_ASSOCIATION_IDENTIFIER"
	PublicImportErrorErrorTypeOutsideValidTermRange                        PublicImportErrorErrorType = "OUTSIDE_VALID_TERM_RANGE"
	PublicImportErrorErrorTypeOutsideValidTimeRange                        PublicImportErrorErrorType = "OUTSIDE_VALID_TIME_RANGE"
	PublicImportErrorErrorTypePortalWideCustomObjectLimitExceeded          PublicImportErrorErrorType = "PORTAL_WIDE_CUSTOM_OBJECT_LIMIT_EXCEEDED"
	PublicImportErrorErrorTypePropertyDefinitionNotFound                   PublicImportErrorErrorType = "PROPERTY_DEFINITION_NOT_FOUND"
	PublicImportErrorErrorTypePropertyValueNotFound                        PublicImportErrorErrorType = "PROPERTY_VALUE_NOT_FOUND"
	PublicImportErrorErrorTypeRowDataTooLarge                              PublicImportErrorErrorType = "ROW_DATA_TOO_LARGE"
	PublicImportErrorErrorTypeSecondaryEmailWriteFailure                   PublicImportErrorErrorType = "SECONDARY_EMAIL_WRITE_FAILURE"
	PublicImportErrorErrorTypeUnknownAssociationRecordID                   PublicImportErrorErrorType = "UNKNOWN_ASSOCIATION_RECORD_ID"
	PublicImportErrorErrorTypeUnknownBadRequest                            PublicImportErrorErrorType = "UNKNOWN_BAD_REQUEST"
	PublicImportErrorErrorTypeUnknownError                                 PublicImportErrorErrorType = "UNKNOWN_ERROR"
	PublicImportErrorErrorTypeUpdateOnlyImport                             PublicImportErrorErrorType = "UPDATE_ONLY_IMPORT"
)

// The CRM object type affected by this error.
type PublicImportErrorObjectType string

const (
	PublicImportErrorObjectTypeAbandonedCart                     PublicImportErrorObjectType = "ABANDONED_CART"
	PublicImportErrorObjectTypeAcceptanceTest                    PublicImportErrorObjectType = "ACCEPTANCE_TEST"
	PublicImportErrorObjectTypeAd                                PublicImportErrorObjectType = "AD"
	PublicImportErrorObjectTypeAdAccount                         PublicImportErrorObjectType = "AD_ACCOUNT"
	PublicImportErrorObjectTypeAdCampaign                        PublicImportErrorObjectType = "AD_CAMPAIGN"
	PublicImportErrorObjectTypeAdGroup                           PublicImportErrorObjectType = "AD_GROUP"
	PublicImportErrorObjectTypeAIForecast                        PublicImportErrorObjectType = "AI_FORECAST"
	PublicImportErrorObjectTypeAllPages                          PublicImportErrorObjectType = "ALL_PAGES"
	PublicImportErrorObjectTypeApproval                          PublicImportErrorObjectType = "APPROVAL"
	PublicImportErrorObjectTypeApprovalStep                      PublicImportErrorObjectType = "APPROVAL_STEP"
	PublicImportErrorObjectTypeAttribution                       PublicImportErrorObjectType = "ATTRIBUTION"
	PublicImportErrorObjectTypeAudience                          PublicImportErrorObjectType = "AUDIENCE"
	PublicImportErrorObjectTypeAutomationJourney                 PublicImportErrorObjectType = "AUTOMATION_JOURNEY"
	PublicImportErrorObjectTypeAutomationPlatformFlow            PublicImportErrorObjectType = "AUTOMATION_PLATFORM_FLOW"
	PublicImportErrorObjectTypeAutomationPlatformFlowAction      PublicImportErrorObjectType = "AUTOMATION_PLATFORM_FLOW_ACTION"
	PublicImportErrorObjectTypeBetAlert                          PublicImportErrorObjectType = "BET_ALERT"
	PublicImportErrorObjectTypeBetDeliverableService             PublicImportErrorObjectType = "BET_DELIVERABLE_SERVICE"
	PublicImportErrorObjectTypeBlogListingPage                   PublicImportErrorObjectType = "BLOG_LISTING_PAGE"
	PublicImportErrorObjectTypeBlogPost                          PublicImportErrorObjectType = "BLOG_POST"
	PublicImportErrorObjectTypeCall                              PublicImportErrorObjectType = "CALL"
	PublicImportErrorObjectTypeCampaign                          PublicImportErrorObjectType = "CAMPAIGN"
	PublicImportErrorObjectTypeCampaignBudgetItem                PublicImportErrorObjectType = "CAMPAIGN_BUDGET_ITEM"
	PublicImportErrorObjectTypeCampaignSpendItem                 PublicImportErrorObjectType = "CAMPAIGN_SPEND_ITEM"
	PublicImportErrorObjectTypeCampaignStep                      PublicImportErrorObjectType = "CAMPAIGN_STEP"
	PublicImportErrorObjectTypeCampaignTemplate                  PublicImportErrorObjectType = "CAMPAIGN_TEMPLATE"
	PublicImportErrorObjectTypeCampaignTemplateStep              PublicImportErrorObjectType = "CAMPAIGN_TEMPLATE_STEP"
	PublicImportErrorObjectTypeCart                              PublicImportErrorObjectType = "CART"
	PublicImportErrorObjectTypeCaseStudy                         PublicImportErrorObjectType = "CASE_STUDY"
	PublicImportErrorObjectTypeChatflow                          PublicImportErrorObjectType = "CHATFLOW"
	PublicImportErrorObjectTypeClip                              PublicImportErrorObjectType = "CLIP"
	PublicImportErrorObjectTypeCmsURL                            PublicImportErrorObjectType = "CMS_URL"
	PublicImportErrorObjectTypeComboEventConfiguration           PublicImportErrorObjectType = "COMBO_EVENT_CONFIGURATION"
	PublicImportErrorObjectTypeCommercePayment                   PublicImportErrorObjectType = "COMMERCE_PAYMENT"
	PublicImportErrorObjectTypeCommunication                     PublicImportErrorObjectType = "COMMUNICATION"
	PublicImportErrorObjectTypeCompany                           PublicImportErrorObjectType = "COMPANY"
	PublicImportErrorObjectTypeContact                           PublicImportErrorObjectType = "CONTACT"
	PublicImportErrorObjectTypeContactCreateAttribution          PublicImportErrorObjectType = "CONTACT_CREATE_ATTRIBUTION"
	PublicImportErrorObjectTypeContent                           PublicImportErrorObjectType = "CONTENT"
	PublicImportErrorObjectTypeContentAudit                      PublicImportErrorObjectType = "CONTENT_AUDIT"
	PublicImportErrorObjectTypeContentAuditPage                  PublicImportErrorObjectType = "CONTENT_AUDIT_PAGE"
	PublicImportErrorObjectTypeConversation                      PublicImportErrorObjectType = "CONVERSATION"
	PublicImportErrorObjectTypeConversationInbox                 PublicImportErrorObjectType = "CONVERSATION_INBOX"
	PublicImportErrorObjectTypeConversationSession               PublicImportErrorObjectType = "CONVERSATION_SESSION"
	PublicImportErrorObjectTypeCrmObjectsDummyType               PublicImportErrorObjectType = "CRM_OBJECTS_DUMMY_TYPE"
	PublicImportErrorObjectTypeCrmPipelinesDummyType             PublicImportErrorObjectType = "CRM_PIPELINES_DUMMY_TYPE"
	PublicImportErrorObjectTypeCta                               PublicImportErrorObjectType = "CTA"
	PublicImportErrorObjectTypeCtaVariant                        PublicImportErrorObjectType = "CTA_VARIANT"
	PublicImportErrorObjectTypeDataPrivacyConsent                PublicImportErrorObjectType = "DATA_PRIVACY_CONSENT"
	PublicImportErrorObjectTypeDataSyncState                     PublicImportErrorObjectType = "DATA_SYNC_STATE"
	PublicImportErrorObjectTypeDeal                              PublicImportErrorObjectType = "DEAL"
	PublicImportErrorObjectTypeDealCreateAttribution             PublicImportErrorObjectType = "DEAL_CREATE_ATTRIBUTION"
	PublicImportErrorObjectTypeDealRegistration                  PublicImportErrorObjectType = "DEAL_REGISTRATION"
	PublicImportErrorObjectTypeDealSplit                         PublicImportErrorObjectType = "DEAL_SPLIT"
	PublicImportErrorObjectTypeDiscount                          PublicImportErrorObjectType = "DISCOUNT"
	PublicImportErrorObjectTypeDiscountCode                      PublicImportErrorObjectType = "DISCOUNT_CODE"
	PublicImportErrorObjectTypeDiscountTemplate                  PublicImportErrorObjectType = "DISCOUNT_TEMPLATE"
	PublicImportErrorObjectTypeEmail                             PublicImportErrorObjectType = "EMAIL"
	PublicImportErrorObjectTypeEngagement                        PublicImportErrorObjectType = "ENGAGEMENT"
	PublicImportErrorObjectTypeExport                            PublicImportErrorObjectType = "EXPORT"
	PublicImportErrorObjectTypeExternalWebURL                    PublicImportErrorObjectType = "EXTERNAL_WEB_URL"
	PublicImportErrorObjectTypeFee                               PublicImportErrorObjectType = "FEE"
	PublicImportErrorObjectTypeFeedbackSubmission                PublicImportErrorObjectType = "FEEDBACK_SUBMISSION"
	PublicImportErrorObjectTypeFeedbackSurvey                    PublicImportErrorObjectType = "FEEDBACK_SURVEY"
	PublicImportErrorObjectTypeFileManagerFile                   PublicImportErrorObjectType = "FILE_MANAGER_FILE"
	PublicImportErrorObjectTypeFileManagerFolder                 PublicImportErrorObjectType = "FILE_MANAGER_FOLDER"
	PublicImportErrorObjectTypeFolder                            PublicImportErrorObjectType = "FOLDER"
	PublicImportErrorObjectTypeForecast                          PublicImportErrorObjectType = "FORECAST"
	PublicImportErrorObjectTypeForm                              PublicImportErrorObjectType = "FORM"
	PublicImportErrorObjectTypeFormSubmissionInbounddb           PublicImportErrorObjectType = "FORM_SUBMISSION_INBOUNDDB"
	PublicImportErrorObjectTypeGoalTarget                        PublicImportErrorObjectType = "GOAL_TARGET"
	PublicImportErrorObjectTypeGoalTargetGroup                   PublicImportErrorObjectType = "GOAL_TARGET_GROUP"
	PublicImportErrorObjectTypeGoalTemplate                      PublicImportErrorObjectType = "GOAL_TEMPLATE"
	PublicImportErrorObjectTypeGscProperty                       PublicImportErrorObjectType = "GSC_PROPERTY"
	PublicImportErrorObjectTypeHub                               PublicImportErrorObjectType = "HUB"
	PublicImportErrorObjectTypeImport                            PublicImportErrorObjectType = "IMPORT"
	PublicImportErrorObjectTypeInvoice                           PublicImportErrorObjectType = "INVOICE"
	PublicImportErrorObjectTypeKeyword                           PublicImportErrorObjectType = "KEYWORD"
	PublicImportErrorObjectTypeKnowledgeArticle                  PublicImportErrorObjectType = "KNOWLEDGE_ARTICLE"
	PublicImportErrorObjectTypeLandingPage                       PublicImportErrorObjectType = "LANDING_PAGE"
	PublicImportErrorObjectTypeLead                              PublicImportErrorObjectType = "LEAD"
	PublicImportErrorObjectTypeLineItem                          PublicImportErrorObjectType = "LINE_ITEM"
	PublicImportErrorObjectTypeMarketingCalendar                 PublicImportErrorObjectType = "MARKETING_CALENDAR"
	PublicImportErrorObjectTypeMarketingCampaignUtm              PublicImportErrorObjectType = "MARKETING_CAMPAIGN_UTM"
	PublicImportErrorObjectTypeMarketingEmail                    PublicImportErrorObjectType = "MARKETING_EMAIL"
	PublicImportErrorObjectTypeMarketingEvent                    PublicImportErrorObjectType = "MARKETING_EVENT"
	PublicImportErrorObjectTypeMarketingEventAttendance          PublicImportErrorObjectType = "MARKETING_EVENT_ATTENDANCE"
	PublicImportErrorObjectTypeMarketingSMS                      PublicImportErrorObjectType = "MARKETING_SMS"
	PublicImportErrorObjectTypeMediaBridge                       PublicImportErrorObjectType = "MEDIA_BRIDGE"
	PublicImportErrorObjectTypeMeetingEvent                      PublicImportErrorObjectType = "MEETING_EVENT"
	PublicImportErrorObjectTypeMic                               PublicImportErrorObjectType = "MIC"
	PublicImportErrorObjectTypeNote                              PublicImportErrorObjectType = "NOTE"
	PublicImportErrorObjectTypeObjectList                        PublicImportErrorObjectType = "OBJECT_LIST"
	PublicImportErrorObjectTypeOrder                             PublicImportErrorObjectType = "ORDER"
	PublicImportErrorObjectTypeOwner                             PublicImportErrorObjectType = "OWNER"
	PublicImportErrorObjectTypePartnerAccount                    PublicImportErrorObjectType = "PARTNER_ACCOUNT"
	PublicImportErrorObjectTypePartnerClient                     PublicImportErrorObjectType = "PARTNER_CLIENT"
	PublicImportErrorObjectTypePartnerClientRevenue              PublicImportErrorObjectType = "PARTNER_CLIENT_REVENUE"
	PublicImportErrorObjectTypePartnerService                    PublicImportErrorObjectType = "PARTNER_SERVICE"
	PublicImportErrorObjectTypePaymentLink                       PublicImportErrorObjectType = "PAYMENT_LINK"
	PublicImportErrorObjectTypePaymentSchedule                   PublicImportErrorObjectType = "PAYMENT_SCHEDULE"
	PublicImportErrorObjectTypePaymentScheduleInstallment        PublicImportErrorObjectType = "PAYMENT_SCHEDULE_INSTALLMENT"
	PublicImportErrorObjectTypePermissionsTesting                PublicImportErrorObjectType = "PERMISSIONS_TESTING"
	PublicImportErrorObjectTypePlaybook                          PublicImportErrorObjectType = "PLAYBOOK"
	PublicImportErrorObjectTypePlaybookQuestion                  PublicImportErrorObjectType = "PLAYBOOK_QUESTION"
	PublicImportErrorObjectTypePlaybookSubmission                PublicImportErrorObjectType = "PLAYBOOK_SUBMISSION"
	PublicImportErrorObjectTypePlaybookSubmissionAnswer          PublicImportErrorObjectType = "PLAYBOOK_SUBMISSION_ANSWER"
	PublicImportErrorObjectTypePlaylist                          PublicImportErrorObjectType = "PLAYLIST"
	PublicImportErrorObjectTypePlaylistFolder                    PublicImportErrorObjectType = "PLAYLIST_FOLDER"
	PublicImportErrorObjectTypePodcastEpisode                    PublicImportErrorObjectType = "PODCAST_EPISODE"
	PublicImportErrorObjectTypePortal                            PublicImportErrorObjectType = "PORTAL"
	PublicImportErrorObjectTypePortalObjectSyncMessage           PublicImportErrorObjectType = "PORTAL_OBJECT_SYNC_MESSAGE"
	PublicImportErrorObjectTypePostalMail                        PublicImportErrorObjectType = "POSTAL_MAIL"
	PublicImportErrorObjectTypePrivacyScannerCookie              PublicImportErrorObjectType = "PRIVACY_SCANNER_COOKIE"
	PublicImportErrorObjectTypeProduct                           PublicImportErrorObjectType = "PRODUCT"
	PublicImportErrorObjectTypeProductOrFolder                   PublicImportErrorObjectType = "PRODUCT_OR_FOLDER"
	PublicImportErrorObjectTypePropertyInfo                      PublicImportErrorObjectType = "PROPERTY_INFO"
	PublicImportErrorObjectTypeProspectingAgentContactAssignment PublicImportErrorObjectType = "PROSPECTING_AGENT_CONTACT_ASSIGNMENT"
	PublicImportErrorObjectTypePublishingTask                    PublicImportErrorObjectType = "PUBLISHING_TASK"
	PublicImportErrorObjectTypeQuarantinedSubmission             PublicImportErrorObjectType = "QUARANTINED_SUBMISSION"
	PublicImportErrorObjectTypeQuota                             PublicImportErrorObjectType = "QUOTA"
	PublicImportErrorObjectTypeQuote                             PublicImportErrorObjectType = "QUOTE"
	PublicImportErrorObjectTypeQuoteField                        PublicImportErrorObjectType = "QUOTE_FIELD"
	PublicImportErrorObjectTypeQuoteModule                       PublicImportErrorObjectType = "QUOTE_MODULE"
	PublicImportErrorObjectTypeQuoteModuleField                  PublicImportErrorObjectType = "QUOTE_MODULE_FIELD"
	PublicImportErrorObjectTypeQuoteTemplate                     PublicImportErrorObjectType = "QUOTE_TEMPLATE"
	PublicImportErrorObjectTypeRestorableCrmObject               PublicImportErrorObjectType = "RESTORABLE_CRM_OBJECT"
	PublicImportErrorObjectTypeRoster                            PublicImportErrorObjectType = "ROSTER"
	PublicImportErrorObjectTypeRosterMember                      PublicImportErrorObjectType = "ROSTER_MEMBER"
	PublicImportErrorObjectTypeSalesDocument                     PublicImportErrorObjectType = "SALES_DOCUMENT"
	PublicImportErrorObjectTypeSalesTask                         PublicImportErrorObjectType = "SALES_TASK"
	PublicImportErrorObjectTypeSalesWorkload                     PublicImportErrorObjectType = "SALES_WORKLOAD"
	PublicImportErrorObjectTypeSalesforceSyncError               PublicImportErrorObjectType = "SALESFORCE_SYNC_ERROR"
	PublicImportErrorObjectTypeSchedulingPage                    PublicImportErrorObjectType = "SCHEDULING_PAGE"
	PublicImportErrorObjectTypeSchemasBackendTest                PublicImportErrorObjectType = "SCHEMAS_BACKEND_TEST"
	PublicImportErrorObjectTypeScoreConfiguration                PublicImportErrorObjectType = "SCORE_CONFIGURATION"
	PublicImportErrorObjectTypeSequence                          PublicImportErrorObjectType = "SEQUENCE"
	PublicImportErrorObjectTypeSequenceEnrollment                PublicImportErrorObjectType = "SEQUENCE_ENROLLMENT"
	PublicImportErrorObjectTypeSequenceStep                      PublicImportErrorObjectType = "SEQUENCE_STEP"
	PublicImportErrorObjectTypeSequenceStepEnrollment            PublicImportErrorObjectType = "SEQUENCE_STEP_ENROLLMENT"
	PublicImportErrorObjectTypeService                           PublicImportErrorObjectType = "SERVICE"
	PublicImportErrorObjectTypeSitePage                          PublicImportErrorObjectType = "SITE_PAGE"
	PublicImportErrorObjectTypeSnippet                           PublicImportErrorObjectType = "SNIPPET"
	PublicImportErrorObjectTypeSocialBroadcast                   PublicImportErrorObjectType = "SOCIAL_BROADCAST"
	PublicImportErrorObjectTypeSocialChannel                     PublicImportErrorObjectType = "SOCIAL_CHANNEL"
	PublicImportErrorObjectTypeSocialPost                        PublicImportErrorObjectType = "SOCIAL_POST"
	PublicImportErrorObjectTypeSocialProfile                     PublicImportErrorObjectType = "SOCIAL_PROFILE"
	PublicImportErrorObjectTypeSoxProtectedDummyType             PublicImportErrorObjectType = "SOX_PROTECTED_DUMMY_TYPE"
	PublicImportErrorObjectTypeSoxProtectedTestType              PublicImportErrorObjectType = "SOX_PROTECTED_TEST_TYPE"
	PublicImportErrorObjectTypeSubmissionTag                     PublicImportErrorObjectType = "SUBMISSION_TAG"
	PublicImportErrorObjectTypeSubscription                      PublicImportErrorObjectType = "SUBSCRIPTION"
	PublicImportErrorObjectTypeTask                              PublicImportErrorObjectType = "TASK"
	PublicImportErrorObjectTypeTaskTemplate                      PublicImportErrorObjectType = "TASK_TEMPLATE"
	PublicImportErrorObjectTypeTax                               PublicImportErrorObjectType = "TAX"
	PublicImportErrorObjectTypeTemplate                          PublicImportErrorObjectType = "TEMPLATE"
	PublicImportErrorObjectTypeTicket                            PublicImportErrorObjectType = "TICKET"
	PublicImportErrorObjectTypeUnknown                           PublicImportErrorObjectType = "UNKNOWN"
	PublicImportErrorObjectTypeUnsubscribe                       PublicImportErrorObjectType = "UNSUBSCRIBE"
	PublicImportErrorObjectTypeUser                              PublicImportErrorObjectType = "USER"
	PublicImportErrorObjectTypeView                              PublicImportErrorObjectType = "VIEW"
	PublicImportErrorObjectTypeViewBlock                         PublicImportErrorObjectType = "VIEW_BLOCK"
	PublicImportErrorObjectTypeWebInteractive                    PublicImportErrorObjectType = "WEB_INTERACTIVE"
)

type PublicImportMetadata struct {
	// Summarized outcomes of each row a developer attempted to import into HubSpot.
	Counters map[string]int64 `json:"counters" api:"required"`
	// The IDs of files uploaded in the File Manager API.
	FileIDs []string `json:"fileIds" api:"required"`
	// The lists containing the imported objects.
	ObjectLists []PublicObjectListRecord `json:"objectLists" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Counters    respjson.Field
		FileIDs     respjson.Field
		ObjectLists respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PublicImportMetadata) RawJSON() string { return r.JSON.raw }
func (r *PublicImportMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PublicImportResponse struct {
	// The unique identifier for this import.
	ID string `json:"id" api:"required"`
	// The timestamp when the object was created, in ISO 8601 format.
	CreatedAt           time.Time            `json:"createdAt" api:"required" format:"date-time"`
	MappedObjectTypeIDs []string             `json:"mappedObjectTypeIds" api:"required"`
	Metadata            PublicImportMetadata `json:"metadata" api:"required"`
	// Whether or not the import is a list of people disqualified from receiving
	// emails.
	OptOutImport bool `json:"optOutImport" api:"required"`
	// The status of the import.
	//
	// Any of "CANCELED", "DEFERRED", "DONE", "FAILED", "PROCESSING", "REVERTED",
	// "STARTED".
	State PublicImportResponseState `json:"state" api:"required"`
	// The timestamp when the import record was last updated, formatted as an ISO 8601
	// instant.
	UpdatedAt time.Time `json:"updatedAt" api:"required" format:"date-time"`
	// The user-provided name for this import.
	ImportName string `json:"importName"`
	// The complete import request configuration as a JSON object.
	ImportRequestJson any `json:"importRequestJson"`
	// Indicates where/how the import was initiated.
	//
	// Any of "API", "CRM_UI", "IMPORT", "MOBILE_ANDROID", "MOBILE_IOS", "SALESFORCE".
	ImportSource   PublicImportResponseImportSource `json:"importSource"`
	ImportTemplate ImportTemplate                   `json:"importTemplate"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                  respjson.Field
		CreatedAt           respjson.Field
		MappedObjectTypeIDs respjson.Field
		Metadata            respjson.Field
		OptOutImport        respjson.Field
		State               respjson.Field
		UpdatedAt           respjson.Field
		ImportName          respjson.Field
		ImportRequestJson   respjson.Field
		ImportSource        respjson.Field
		ImportTemplate      respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PublicImportResponse) RawJSON() string { return r.JSON.raw }
func (r *PublicImportResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the import.
type PublicImportResponseState string

const (
	PublicImportResponseStateCanceled   PublicImportResponseState = "CANCELED"
	PublicImportResponseStateDeferred   PublicImportResponseState = "DEFERRED"
	PublicImportResponseStateDone       PublicImportResponseState = "DONE"
	PublicImportResponseStateFailed     PublicImportResponseState = "FAILED"
	PublicImportResponseStateProcessing PublicImportResponseState = "PROCESSING"
	PublicImportResponseStateReverted   PublicImportResponseState = "REVERTED"
	PublicImportResponseStateStarted    PublicImportResponseState = "STARTED"
)

// Indicates where/how the import was initiated.
type PublicImportResponseImportSource string

const (
	PublicImportResponseImportSourceAPI           PublicImportResponseImportSource = "API"
	PublicImportResponseImportSourceCrmUi         PublicImportResponseImportSource = "CRM_UI"
	PublicImportResponseImportSourceImport        PublicImportResponseImportSource = "IMPORT"
	PublicImportResponseImportSourceMobileAndroid PublicImportResponseImportSource = "MOBILE_ANDROID"
	PublicImportResponseImportSourceMobileIos     PublicImportResponseImportSource = "MOBILE_IOS"
	PublicImportResponseImportSourceSalesforce    PublicImportResponseImportSource = "SALESFORCE"
)

type PublicObjectListRecord struct {
	// The ID of the list containing the imported objects.
	ListID string `json:"listId" api:"required"`
	// The type of object contained in the list.
	ObjectType string `json:"objectType" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ListID      respjson.Field
		ObjectType  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PublicObjectListRecord) RawJSON() string { return r.JSON.raw }
func (r *PublicObjectListRecord) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ImportNewParams struct {
	ImportRequest param.Opt[string] `json:"importRequest,omitzero"`
	Files         io.Reader         `json:"files,omitzero" format:"binary"`
	paramObj
}

func (r ImportNewParams) MarshalMultipart() (data []byte, contentType string, err error) {
	buf := bytes.NewBuffer(nil)
	writer := multipart.NewWriter(buf)
	err = apiform.MarshalRoot(r, writer)
	if err == nil {
		err = apiform.WriteExtras(writer, r.ExtraFields())
	}
	if err != nil {
		writer.Close()
		return nil, "", err
	}
	err = writer.Close()
	if err != nil {
		return nil, "", err
	}
	return buf.Bytes(), writer.FormDataContentType(), nil
}

type ImportListParams struct {
	// The paging cursor token of the last successfully read resource will be returned
	// as the `paging.next.after` JSON property of a paged response containing more
	// results.
	After param.Opt[string] `query:"after,omitzero" json:"-"`
	// The maximum number of results to display per page.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ImportListParams]'s query parameters as `url.Values`.
func (r ImportListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ImportListErrorsParams struct {
	// The paging cursor token of the last successfully read resource will be returned
	// as the `paging.next.after` JSON property of a paged response containing more
	// results.
	After               param.Opt[string] `query:"after,omitzero" json:"-"`
	IncludeErrorMessage param.Opt[bool]   `query:"includeErrorMessage,omitzero" json:"-"`
	IncludeRowData      param.Opt[bool]   `query:"includeRowData,omitzero" json:"-"`
	// The maximum number of results to display per page.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ImportListErrorsParams]'s query parameters as `url.Values`.
func (r ImportListErrorsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
