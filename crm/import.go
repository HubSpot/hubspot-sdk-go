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

	"github.com/stainless-sdks/hubspot-sdk-go/internal/apiform"
	"github.com/stainless-sdks/hubspot-sdk-go/internal/apijson"
	"github.com/stainless-sdks/hubspot-sdk-go/internal/apiquery"
	"github.com/stainless-sdks/hubspot-sdk-go/internal/requestconfig"
	"github.com/stainless-sdks/hubspot-sdk-go/marketing"
	"github.com/stainless-sdks/hubspot-sdk-go/option"
	"github.com/stainless-sdks/hubspot-sdk-go/packages/pagination"
	"github.com/stainless-sdks/hubspot-sdk-go/packages/param"
	"github.com/stainless-sdks/hubspot-sdk-go/packages/respjson"
	"github.com/stainless-sdks/hubspot-sdk-go/shared"
)

// ImportService contains methods and other services that help with interacting
// with the Hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewImportService] method instead.
type ImportService struct {
	Options []option.RequestOption
}

// NewImportService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewImportService(opts ...option.RequestOption) (r ImportService) {
	r = ImportService{}
	r.Options = opts
	return
}

// Begins importing data from the specified file resources. This uploads the
// corresponding file and uses the import request object to convert rows in the
// files to objects.
func (r *ImportService) New(ctx context.Context, body ImportNewParams, opts ...option.RequestOption) (res *PublicImportResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "crm/v3/imports/"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Returns a paged list of active imports for this account.
func (r *ImportService) List(ctx context.Context, query ImportListParams, opts ...option.RequestOption) (res *pagination.Page[PublicImportResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "crm/v3/imports/"
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

// Returns a paged list of active imports for this account.
func (r *ImportService) ListAutoPaging(ctx context.Context, query ImportListParams, opts ...option.RequestOption) *pagination.PageAutoPager[PublicImportResponse] {
	return pagination.NewPageAutoPager(r.List(ctx, query, opts...))
}

// This allows a developer to cancel an active import.
func (r *ImportService) Cancel(ctx context.Context, importID int64, opts ...option.RequestOption) (res *shared.ActionResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("crm/v3/imports/%v/cancel", importID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// A complete summary of an import record, including any updates.
func (r *ImportService) Get(ctx context.Context, importID int64, opts ...option.RequestOption) (res *PublicImportResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("crm/v3/imports/%v", importID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

func (r *ImportService) ListErrors(ctx context.Context, importID int64, query ImportListErrorsParams, opts ...option.RequestOption) (res *CollectionResponsePublicImportErrorForwardPaging, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("crm/v3/imports/%v/errors", importID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type CollectionResponsePublicImportErrorForwardPaging struct {
	Results []PublicImportError  `json:"results,required"`
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

type CollectionResponsePublicImportResponse struct {
	Results []PublicImportResponse `json:"results,required"`
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
func (r CollectionResponsePublicImportResponse) RawJSON() string { return r.JSON.raw }
func (r *CollectionResponsePublicImportResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ImportRowCore struct {
	ContainsEncryptedProperties bool     `json:"containsEncryptedProperties,required"`
	FileID                      int64    `json:"fileId,required"`
	LineNumber                  int64    `json:"lineNumber,required"`
	RowData                     []string `json:"rowData,required"`
	PageName                    string   `json:"pageName"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
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
	TemplateID int64 `json:"templateId,required"`
	// Any of "admin_defined", "previous_import", "user_file".
	TemplateType ImportTemplateTemplateType `json:"templateType,required"`
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

type ImportTemplateTemplateType string

const (
	ImportTemplateTemplateTypeAdminDefined   ImportTemplateTemplateType = "admin_defined"
	ImportTemplateTemplateTypePreviousImport ImportTemplateTemplateType = "previous_import"
	ImportTemplateTemplateTypeUserFile       ImportTemplateTemplateType = "user_file"
)

type PublicImportError struct {
	ID        string `json:"id,required"`
	CreatedAt int64  `json:"createdAt,required"`
	// Any of "INCORRECT_NUMBER_OF_COLUMNS", "INVALID_OBJECT_ID",
	// "INVALID_ASSOCIATION_IDENTIFIER", "NO_OBJECT_ID_FROM_ASSOCIATION_IDENTIFIER",
	// "MULTIPLE_COMPANIES_WITH_THIS_DOMAIN", "PROPERTY_DEFINITION_NOT_FOUND",
	// "PROPERTY_VALUE_NOT_FOUND", "COULD_NOT_FIND_OWNER", "MULTIPLE_OWNERS_FOUND",
	// "COULD_NOT_FIND_BUSINESS_UNIT", "COULD_NOT_PARSE_NUMBER",
	// "COULD_NOT_PARSE_DATE", "COULD_NOT_PARSE_TERM", "OUTSIDE_VALID_TIME_RANGE",
	// "OUTSIDE_VALID_TERM_RANGE", "COULD_NOT_PARSE_ROW", "INVALID_ENUMERATION_OPTION",
	// "AMBIGUOUS_ENUMERATION_OPTION", "FAILED_VALIDATION",
	// "FAILED_TO_CREATE_ASSOCIATION", "ASSOCIATION_LIMIT_EXCEEDED", "FILE_NOT_FOUND",
	// "INVALID_COLUMN_CONFIGURATION", "INVALID_FILE_TYPE", "INVALID_SPREADSHEET",
	// "INVALID_SHEET_COUNT", "FAILED_TO_PROCESS_OBJECT_WITH_EMPTY_PROPERTY_VALUES",
	// "UNKNOWN_BAD_REQUEST", "GDPR_BLACKLISTED_EMAIL", "DUPLICATE_ASSOCIATION_ID",
	// "LIMIT_EXCEEDED", "PORTAL_WIDE_CUSTOM_OBJECT_LIMIT_EXCEEDED",
	// "INVALID_ALTERNATE_ID", "INVALID_EMAIL", "SECONDARY_EMAIL_WRITE_FAILURE",
	// "INVALID_DOMAIN", "DUPLICATE_ROW_CONTENT", "INVALID_NUMBER_SIZE",
	// "UNKNOWN_ERROR", "FAILED_TO_OPT_OUT_CONTACT", "INVALID_REQUIRED_PROPERTY",
	// "MISSING_REQUIRED_PROPERTY", "DUPLICATE_ALTERNATE_ID", "DUPLICATE_OBJECT_ID",
	// "DUPLICATE_UNIQUE_PROPERTY_VALUE", "UNKNOWN_ASSOCIATION_RECORD_ID",
	// "INVALID_RECORD_ID", "DUPLICATE_RECORD_ID",
	// "INVALID_CUSTOM_PROPERTY_VALIDATION", "CREATE_ONLY_IMPORT",
	// "UPDATE_ONLY_IMPORT", "COLUMN_TOO_LARGE", "ROW_DATA_TOO_LARGE",
	// "MISSING_EVENT_TIMESTAMP", "INVALID_EVENT_TIMESTAMP", "INVALID_EVENT",
	// "DUPLICATE_EVENT", "MISSING_EVENT_DEFINITION", "INVALID_ASSOCIATION_KEY",
	// "ASSOCIATION_RECORD_NOT_FOUND", "MISSING_OBJECT_DEFINITION",
	// "ASSOCIATION_LABEL_NOT_FOUND", "MANY_ERRORS_IN_ROW".
	ErrorType    PublicImportErrorErrorType `json:"errorType,required"`
	SourceData   ImportRowCore              `json:"sourceData,required"`
	ErrorMessage string                     `json:"errorMessage"`
	ExtraContext string                     `json:"extraContext"`
	// Represents a single custom property of a marketing event, storing its name,
	// value, metadata (like source, timestamp, and sensitivity), and related audit
	// information for tracking changes.
	InvalidPropertyValue  marketing.PropertyValue `json:"invalidPropertyValue"`
	InvalidValue          string                  `json:"invalidValue"`
	InvalidValueToDisplay string                  `json:"invalidValueToDisplay"`
	KnownColumnNumber     int64                   `json:"knownColumnNumber"`
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
	// "PARTNER_CLIENT_REVENUE", "AUTOMATION_JOURNEY", "COMBO_EVENT_CONFIGURATION",
	// "CRM_OBJECTS_DUMMY_TYPE", "CASE_STUDY", "SERVICE", "PODCAST_EPISODE",
	// "PARTNER_SERVICE", "PROSPECTING_AGENT_CONTACT_ASSIGNMENT", "UNKNOWN".
	ObjectType   PublicImportErrorObjectType `json:"objectType"`
	ObjectTypeID string                      `json:"objectTypeId"`
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

type PublicImportErrorErrorType string

const (
	PublicImportErrorErrorTypeIncorrectNumberOfColumns                     PublicImportErrorErrorType = "INCORRECT_NUMBER_OF_COLUMNS"
	PublicImportErrorErrorTypeInvalidObjectID                              PublicImportErrorErrorType = "INVALID_OBJECT_ID"
	PublicImportErrorErrorTypeInvalidAssociationIdentifier                 PublicImportErrorErrorType = "INVALID_ASSOCIATION_IDENTIFIER"
	PublicImportErrorErrorTypeNoObjectIDFromAssociationIdentifier          PublicImportErrorErrorType = "NO_OBJECT_ID_FROM_ASSOCIATION_IDENTIFIER"
	PublicImportErrorErrorTypeMultipleCompaniesWithThisDomain              PublicImportErrorErrorType = "MULTIPLE_COMPANIES_WITH_THIS_DOMAIN"
	PublicImportErrorErrorTypePropertyDefinitionNotFound                   PublicImportErrorErrorType = "PROPERTY_DEFINITION_NOT_FOUND"
	PublicImportErrorErrorTypePropertyValueNotFound                        PublicImportErrorErrorType = "PROPERTY_VALUE_NOT_FOUND"
	PublicImportErrorErrorTypeCouldNotFindOwner                            PublicImportErrorErrorType = "COULD_NOT_FIND_OWNER"
	PublicImportErrorErrorTypeMultipleOwnersFound                          PublicImportErrorErrorType = "MULTIPLE_OWNERS_FOUND"
	PublicImportErrorErrorTypeCouldNotFindBusinessUnit                     PublicImportErrorErrorType = "COULD_NOT_FIND_BUSINESS_UNIT"
	PublicImportErrorErrorTypeCouldNotParseNumber                          PublicImportErrorErrorType = "COULD_NOT_PARSE_NUMBER"
	PublicImportErrorErrorTypeCouldNotParseDate                            PublicImportErrorErrorType = "COULD_NOT_PARSE_DATE"
	PublicImportErrorErrorTypeCouldNotParseTerm                            PublicImportErrorErrorType = "COULD_NOT_PARSE_TERM"
	PublicImportErrorErrorTypeOutsideValidTimeRange                        PublicImportErrorErrorType = "OUTSIDE_VALID_TIME_RANGE"
	PublicImportErrorErrorTypeOutsideValidTermRange                        PublicImportErrorErrorType = "OUTSIDE_VALID_TERM_RANGE"
	PublicImportErrorErrorTypeCouldNotParseRow                             PublicImportErrorErrorType = "COULD_NOT_PARSE_ROW"
	PublicImportErrorErrorTypeInvalidEnumerationOption                     PublicImportErrorErrorType = "INVALID_ENUMERATION_OPTION"
	PublicImportErrorErrorTypeAmbiguousEnumerationOption                   PublicImportErrorErrorType = "AMBIGUOUS_ENUMERATION_OPTION"
	PublicImportErrorErrorTypeFailedValidation                             PublicImportErrorErrorType = "FAILED_VALIDATION"
	PublicImportErrorErrorTypeFailedToCreateAssociation                    PublicImportErrorErrorType = "FAILED_TO_CREATE_ASSOCIATION"
	PublicImportErrorErrorTypeAssociationLimitExceeded                     PublicImportErrorErrorType = "ASSOCIATION_LIMIT_EXCEEDED"
	PublicImportErrorErrorTypeFileNotFound                                 PublicImportErrorErrorType = "FILE_NOT_FOUND"
	PublicImportErrorErrorTypeInvalidColumnConfiguration                   PublicImportErrorErrorType = "INVALID_COLUMN_CONFIGURATION"
	PublicImportErrorErrorTypeInvalidFileType                              PublicImportErrorErrorType = "INVALID_FILE_TYPE"
	PublicImportErrorErrorTypeInvalidSpreadsheet                           PublicImportErrorErrorType = "INVALID_SPREADSHEET"
	PublicImportErrorErrorTypeInvalidSheetCount                            PublicImportErrorErrorType = "INVALID_SHEET_COUNT"
	PublicImportErrorErrorTypeFailedToProcessObjectWithEmptyPropertyValues PublicImportErrorErrorType = "FAILED_TO_PROCESS_OBJECT_WITH_EMPTY_PROPERTY_VALUES"
	PublicImportErrorErrorTypeUnknownBadRequest                            PublicImportErrorErrorType = "UNKNOWN_BAD_REQUEST"
	PublicImportErrorErrorTypeGdprBlacklistedEmail                         PublicImportErrorErrorType = "GDPR_BLACKLISTED_EMAIL"
	PublicImportErrorErrorTypeDuplicateAssociationID                       PublicImportErrorErrorType = "DUPLICATE_ASSOCIATION_ID"
	PublicImportErrorErrorTypeLimitExceeded                                PublicImportErrorErrorType = "LIMIT_EXCEEDED"
	PublicImportErrorErrorTypePortalWideCustomObjectLimitExceeded          PublicImportErrorErrorType = "PORTAL_WIDE_CUSTOM_OBJECT_LIMIT_EXCEEDED"
	PublicImportErrorErrorTypeInvalidAlternateID                           PublicImportErrorErrorType = "INVALID_ALTERNATE_ID"
	PublicImportErrorErrorTypeInvalidEmail                                 PublicImportErrorErrorType = "INVALID_EMAIL"
	PublicImportErrorErrorTypeSecondaryEmailWriteFailure                   PublicImportErrorErrorType = "SECONDARY_EMAIL_WRITE_FAILURE"
	PublicImportErrorErrorTypeInvalidDomain                                PublicImportErrorErrorType = "INVALID_DOMAIN"
	PublicImportErrorErrorTypeDuplicateRowContent                          PublicImportErrorErrorType = "DUPLICATE_ROW_CONTENT"
	PublicImportErrorErrorTypeInvalidNumberSize                            PublicImportErrorErrorType = "INVALID_NUMBER_SIZE"
	PublicImportErrorErrorTypeUnknownError                                 PublicImportErrorErrorType = "UNKNOWN_ERROR"
	PublicImportErrorErrorTypeFailedToOptOutContact                        PublicImportErrorErrorType = "FAILED_TO_OPT_OUT_CONTACT"
	PublicImportErrorErrorTypeInvalidRequiredProperty                      PublicImportErrorErrorType = "INVALID_REQUIRED_PROPERTY"
	PublicImportErrorErrorTypeMissingRequiredProperty                      PublicImportErrorErrorType = "MISSING_REQUIRED_PROPERTY"
	PublicImportErrorErrorTypeDuplicateAlternateID                         PublicImportErrorErrorType = "DUPLICATE_ALTERNATE_ID"
	PublicImportErrorErrorTypeDuplicateObjectID                            PublicImportErrorErrorType = "DUPLICATE_OBJECT_ID"
	PublicImportErrorErrorTypeDuplicateUniquePropertyValue                 PublicImportErrorErrorType = "DUPLICATE_UNIQUE_PROPERTY_VALUE"
	PublicImportErrorErrorTypeUnknownAssociationRecordID                   PublicImportErrorErrorType = "UNKNOWN_ASSOCIATION_RECORD_ID"
	PublicImportErrorErrorTypeInvalidRecordID                              PublicImportErrorErrorType = "INVALID_RECORD_ID"
	PublicImportErrorErrorTypeDuplicateRecordID                            PublicImportErrorErrorType = "DUPLICATE_RECORD_ID"
	PublicImportErrorErrorTypeInvalidCustomPropertyValidation              PublicImportErrorErrorType = "INVALID_CUSTOM_PROPERTY_VALIDATION"
	PublicImportErrorErrorTypeCreateOnlyImport                             PublicImportErrorErrorType = "CREATE_ONLY_IMPORT"
	PublicImportErrorErrorTypeUpdateOnlyImport                             PublicImportErrorErrorType = "UPDATE_ONLY_IMPORT"
	PublicImportErrorErrorTypeColumnTooLarge                               PublicImportErrorErrorType = "COLUMN_TOO_LARGE"
	PublicImportErrorErrorTypeRowDataTooLarge                              PublicImportErrorErrorType = "ROW_DATA_TOO_LARGE"
	PublicImportErrorErrorTypeMissingEventTimestamp                        PublicImportErrorErrorType = "MISSING_EVENT_TIMESTAMP"
	PublicImportErrorErrorTypeInvalidEventTimestamp                        PublicImportErrorErrorType = "INVALID_EVENT_TIMESTAMP"
	PublicImportErrorErrorTypeInvalidEvent                                 PublicImportErrorErrorType = "INVALID_EVENT"
	PublicImportErrorErrorTypeDuplicateEvent                               PublicImportErrorErrorType = "DUPLICATE_EVENT"
	PublicImportErrorErrorTypeMissingEventDefinition                       PublicImportErrorErrorType = "MISSING_EVENT_DEFINITION"
	PublicImportErrorErrorTypeInvalidAssociationKey                        PublicImportErrorErrorType = "INVALID_ASSOCIATION_KEY"
	PublicImportErrorErrorTypeAssociationRecordNotFound                    PublicImportErrorErrorType = "ASSOCIATION_RECORD_NOT_FOUND"
	PublicImportErrorErrorTypeMissingObjectDefinition                      PublicImportErrorErrorType = "MISSING_OBJECT_DEFINITION"
	PublicImportErrorErrorTypeAssociationLabelNotFound                     PublicImportErrorErrorType = "ASSOCIATION_LABEL_NOT_FOUND"
	PublicImportErrorErrorTypeManyErrorsInRow                              PublicImportErrorErrorType = "MANY_ERRORS_IN_ROW"
)

type PublicImportErrorObjectType string

const (
	PublicImportErrorObjectTypeContact                           PublicImportErrorObjectType = "CONTACT"
	PublicImportErrorObjectTypeCompany                           PublicImportErrorObjectType = "COMPANY"
	PublicImportErrorObjectTypeDeal                              PublicImportErrorObjectType = "DEAL"
	PublicImportErrorObjectTypeEngagement                        PublicImportErrorObjectType = "ENGAGEMENT"
	PublicImportErrorObjectTypeTicket                            PublicImportErrorObjectType = "TICKET"
	PublicImportErrorObjectTypeOwner                             PublicImportErrorObjectType = "OWNER"
	PublicImportErrorObjectTypeProduct                           PublicImportErrorObjectType = "PRODUCT"
	PublicImportErrorObjectTypeLineItem                          PublicImportErrorObjectType = "LINE_ITEM"
	PublicImportErrorObjectTypeBetDeliverableService             PublicImportErrorObjectType = "BET_DELIVERABLE_SERVICE"
	PublicImportErrorObjectTypeContent                           PublicImportErrorObjectType = "CONTENT"
	PublicImportErrorObjectTypeConversation                      PublicImportErrorObjectType = "CONVERSATION"
	PublicImportErrorObjectTypeBetAlert                          PublicImportErrorObjectType = "BET_ALERT"
	PublicImportErrorObjectTypePortal                            PublicImportErrorObjectType = "PORTAL"
	PublicImportErrorObjectTypeQuote                             PublicImportErrorObjectType = "QUOTE"
	PublicImportErrorObjectTypeFormSubmissionInbounddb           PublicImportErrorObjectType = "FORM_SUBMISSION_INBOUNDDB"
	PublicImportErrorObjectTypeQuota                             PublicImportErrorObjectType = "QUOTA"
	PublicImportErrorObjectTypeUnsubscribe                       PublicImportErrorObjectType = "UNSUBSCRIBE"
	PublicImportErrorObjectTypeCommunication                     PublicImportErrorObjectType = "COMMUNICATION"
	PublicImportErrorObjectTypeFeedbackSubmission                PublicImportErrorObjectType = "FEEDBACK_SUBMISSION"
	PublicImportErrorObjectTypeAttribution                       PublicImportErrorObjectType = "ATTRIBUTION"
	PublicImportErrorObjectTypeSalesforceSyncError               PublicImportErrorObjectType = "SALESFORCE_SYNC_ERROR"
	PublicImportErrorObjectTypeRestorableCRMObject               PublicImportErrorObjectType = "RESTORABLE_CRM_OBJECT"
	PublicImportErrorObjectTypeHub                               PublicImportErrorObjectType = "HUB"
	PublicImportErrorObjectTypeLandingPage                       PublicImportErrorObjectType = "LANDING_PAGE"
	PublicImportErrorObjectTypeProductOrFolder                   PublicImportErrorObjectType = "PRODUCT_OR_FOLDER"
	PublicImportErrorObjectTypeTask                              PublicImportErrorObjectType = "TASK"
	PublicImportErrorObjectTypeForm                              PublicImportErrorObjectType = "FORM"
	PublicImportErrorObjectTypeMarketingEmail                    PublicImportErrorObjectType = "MARKETING_EMAIL"
	PublicImportErrorObjectTypeAdAccount                         PublicImportErrorObjectType = "AD_ACCOUNT"
	PublicImportErrorObjectTypeAdCampaign                        PublicImportErrorObjectType = "AD_CAMPAIGN"
	PublicImportErrorObjectTypeAdGroup                           PublicImportErrorObjectType = "AD_GROUP"
	PublicImportErrorObjectTypeAd                                PublicImportErrorObjectType = "AD"
	PublicImportErrorObjectTypeKeyword                           PublicImportErrorObjectType = "KEYWORD"
	PublicImportErrorObjectTypeCampaign                          PublicImportErrorObjectType = "CAMPAIGN"
	PublicImportErrorObjectTypeSocialChannel                     PublicImportErrorObjectType = "SOCIAL_CHANNEL"
	PublicImportErrorObjectTypeSocialPost                        PublicImportErrorObjectType = "SOCIAL_POST"
	PublicImportErrorObjectTypeSitePage                          PublicImportErrorObjectType = "SITE_PAGE"
	PublicImportErrorObjectTypeBlogPost                          PublicImportErrorObjectType = "BLOG_POST"
	PublicImportErrorObjectTypeImport                            PublicImportErrorObjectType = "IMPORT"
	PublicImportErrorObjectTypeExport                            PublicImportErrorObjectType = "EXPORT"
	PublicImportErrorObjectTypeCta                               PublicImportErrorObjectType = "CTA"
	PublicImportErrorObjectTypeTaskTemplate                      PublicImportErrorObjectType = "TASK_TEMPLATE"
	PublicImportErrorObjectTypeAutomationPlatformFlow            PublicImportErrorObjectType = "AUTOMATION_PLATFORM_FLOW"
	PublicImportErrorObjectTypeObjectList                        PublicImportErrorObjectType = "OBJECT_LIST"
	PublicImportErrorObjectTypeNote                              PublicImportErrorObjectType = "NOTE"
	PublicImportErrorObjectTypeMeetingEvent                      PublicImportErrorObjectType = "MEETING_EVENT"
	PublicImportErrorObjectTypeCall                              PublicImportErrorObjectType = "CALL"
	PublicImportErrorObjectTypeEmail                             PublicImportErrorObjectType = "EMAIL"
	PublicImportErrorObjectTypePublishingTask                    PublicImportErrorObjectType = "PUBLISHING_TASK"
	PublicImportErrorObjectTypeConversationSession               PublicImportErrorObjectType = "CONVERSATION_SESSION"
	PublicImportErrorObjectTypeContactCreateAttribution          PublicImportErrorObjectType = "CONTACT_CREATE_ATTRIBUTION"
	PublicImportErrorObjectTypeInvoice                           PublicImportErrorObjectType = "INVOICE"
	PublicImportErrorObjectTypeMarketingEvent                    PublicImportErrorObjectType = "MARKETING_EVENT"
	PublicImportErrorObjectTypeConversationInbox                 PublicImportErrorObjectType = "CONVERSATION_INBOX"
	PublicImportErrorObjectTypeChatflow                          PublicImportErrorObjectType = "CHATFLOW"
	PublicImportErrorObjectTypeMediaBridge                       PublicImportErrorObjectType = "MEDIA_BRIDGE"
	PublicImportErrorObjectTypeSequence                          PublicImportErrorObjectType = "SEQUENCE"
	PublicImportErrorObjectTypeSequenceStep                      PublicImportErrorObjectType = "SEQUENCE_STEP"
	PublicImportErrorObjectTypeForecast                          PublicImportErrorObjectType = "FORECAST"
	PublicImportErrorObjectTypeSnippet                           PublicImportErrorObjectType = "SNIPPET"
	PublicImportErrorObjectTypeTemplate                          PublicImportErrorObjectType = "TEMPLATE"
	PublicImportErrorObjectTypeDealCreateAttribution             PublicImportErrorObjectType = "DEAL_CREATE_ATTRIBUTION"
	PublicImportErrorObjectTypeQuoteTemplate                     PublicImportErrorObjectType = "QUOTE_TEMPLATE"
	PublicImportErrorObjectTypeQuoteModule                       PublicImportErrorObjectType = "QUOTE_MODULE"
	PublicImportErrorObjectTypeQuoteModuleField                  PublicImportErrorObjectType = "QUOTE_MODULE_FIELD"
	PublicImportErrorObjectTypeQuoteField                        PublicImportErrorObjectType = "QUOTE_FIELD"
	PublicImportErrorObjectTypeSequenceEnrollment                PublicImportErrorObjectType = "SEQUENCE_ENROLLMENT"
	PublicImportErrorObjectTypeSubscription                      PublicImportErrorObjectType = "SUBSCRIPTION"
	PublicImportErrorObjectTypeAcceptanceTest                    PublicImportErrorObjectType = "ACCEPTANCE_TEST"
	PublicImportErrorObjectTypeSocialBroadcast                   PublicImportErrorObjectType = "SOCIAL_BROADCAST"
	PublicImportErrorObjectTypeDealSplit                         PublicImportErrorObjectType = "DEAL_SPLIT"
	PublicImportErrorObjectTypeDealRegistration                  PublicImportErrorObjectType = "DEAL_REGISTRATION"
	PublicImportErrorObjectTypeGoalTarget                        PublicImportErrorObjectType = "GOAL_TARGET"
	PublicImportErrorObjectTypeGoalTargetGroup                   PublicImportErrorObjectType = "GOAL_TARGET_GROUP"
	PublicImportErrorObjectTypePortalObjectSyncMessage           PublicImportErrorObjectType = "PORTAL_OBJECT_SYNC_MESSAGE"
	PublicImportErrorObjectTypeFileManagerFile                   PublicImportErrorObjectType = "FILE_MANAGER_FILE"
	PublicImportErrorObjectTypeFileManagerFolder                 PublicImportErrorObjectType = "FILE_MANAGER_FOLDER"
	PublicImportErrorObjectTypeSequenceStepEnrollment            PublicImportErrorObjectType = "SEQUENCE_STEP_ENROLLMENT"
	PublicImportErrorObjectTypeApproval                          PublicImportErrorObjectType = "APPROVAL"
	PublicImportErrorObjectTypeApprovalStep                      PublicImportErrorObjectType = "APPROVAL_STEP"
	PublicImportErrorObjectTypeCtaVariant                        PublicImportErrorObjectType = "CTA_VARIANT"
	PublicImportErrorObjectTypeSalesDocument                     PublicImportErrorObjectType = "SALES_DOCUMENT"
	PublicImportErrorObjectTypeDiscount                          PublicImportErrorObjectType = "DISCOUNT"
	PublicImportErrorObjectTypeFee                               PublicImportErrorObjectType = "FEE"
	PublicImportErrorObjectTypeTax                               PublicImportErrorObjectType = "TAX"
	PublicImportErrorObjectTypeMarketingCalendar                 PublicImportErrorObjectType = "MARKETING_CALENDAR"
	PublicImportErrorObjectTypePermissionsTesting                PublicImportErrorObjectType = "PERMISSIONS_TESTING"
	PublicImportErrorObjectTypePrivacyScannerCookie              PublicImportErrorObjectType = "PRIVACY_SCANNER_COOKIE"
	PublicImportErrorObjectTypeDataSyncState                     PublicImportErrorObjectType = "DATA_SYNC_STATE"
	PublicImportErrorObjectTypeWebInteractive                    PublicImportErrorObjectType = "WEB_INTERACTIVE"
	PublicImportErrorObjectTypePlaybook                          PublicImportErrorObjectType = "PLAYBOOK"
	PublicImportErrorObjectTypeFolder                            PublicImportErrorObjectType = "FOLDER"
	PublicImportErrorObjectTypePlaybookQuestion                  PublicImportErrorObjectType = "PLAYBOOK_QUESTION"
	PublicImportErrorObjectTypePlaybookSubmission                PublicImportErrorObjectType = "PLAYBOOK_SUBMISSION"
	PublicImportErrorObjectTypePlaybookSubmissionAnswer          PublicImportErrorObjectType = "PLAYBOOK_SUBMISSION_ANSWER"
	PublicImportErrorObjectTypeCommercePayment                   PublicImportErrorObjectType = "COMMERCE_PAYMENT"
	PublicImportErrorObjectTypeGscProperty                       PublicImportErrorObjectType = "GSC_PROPERTY"
	PublicImportErrorObjectTypeSoxProtectedDummyType             PublicImportErrorObjectType = "SOX_PROTECTED_DUMMY_TYPE"
	PublicImportErrorObjectTypeBlogListingPage                   PublicImportErrorObjectType = "BLOG_LISTING_PAGE"
	PublicImportErrorObjectTypeQuarantinedSubmission             PublicImportErrorObjectType = "QUARANTINED_SUBMISSION"
	PublicImportErrorObjectTypePaymentSchedule                   PublicImportErrorObjectType = "PAYMENT_SCHEDULE"
	PublicImportErrorObjectTypePaymentScheduleInstallment        PublicImportErrorObjectType = "PAYMENT_SCHEDULE_INSTALLMENT"
	PublicImportErrorObjectTypeMarketingCampaignUtm              PublicImportErrorObjectType = "MARKETING_CAMPAIGN_UTM"
	PublicImportErrorObjectTypeDiscountTemplate                  PublicImportErrorObjectType = "DISCOUNT_TEMPLATE"
	PublicImportErrorObjectTypeDiscountCode                      PublicImportErrorObjectType = "DISCOUNT_CODE"
	PublicImportErrorObjectTypeFeedbackSurvey                    PublicImportErrorObjectType = "FEEDBACK_SURVEY"
	PublicImportErrorObjectTypeCmsURL                            PublicImportErrorObjectType = "CMS_URL"
	PublicImportErrorObjectTypeSalesTask                         PublicImportErrorObjectType = "SALES_TASK"
	PublicImportErrorObjectTypeSalesWorkload                     PublicImportErrorObjectType = "SALES_WORKLOAD"
	PublicImportErrorObjectTypeUser                              PublicImportErrorObjectType = "USER"
	PublicImportErrorObjectTypePostalMail                        PublicImportErrorObjectType = "POSTAL_MAIL"
	PublicImportErrorObjectTypeSchemasBackendTest                PublicImportErrorObjectType = "SCHEMAS_BACKEND_TEST"
	PublicImportErrorObjectTypePaymentLink                       PublicImportErrorObjectType = "PAYMENT_LINK"
	PublicImportErrorObjectTypeSubmissionTag                     PublicImportErrorObjectType = "SUBMISSION_TAG"
	PublicImportErrorObjectTypeCampaignStep                      PublicImportErrorObjectType = "CAMPAIGN_STEP"
	PublicImportErrorObjectTypeSchedulingPage                    PublicImportErrorObjectType = "SCHEDULING_PAGE"
	PublicImportErrorObjectTypeSoxProtectedTestType              PublicImportErrorObjectType = "SOX_PROTECTED_TEST_TYPE"
	PublicImportErrorObjectTypeOrder                             PublicImportErrorObjectType = "ORDER"
	PublicImportErrorObjectTypeMarketingSMS                      PublicImportErrorObjectType = "MARKETING_SMS"
	PublicImportErrorObjectTypePartnerAccount                    PublicImportErrorObjectType = "PARTNER_ACCOUNT"
	PublicImportErrorObjectTypeCampaignTemplate                  PublicImportErrorObjectType = "CAMPAIGN_TEMPLATE"
	PublicImportErrorObjectTypeCampaignTemplateStep              PublicImportErrorObjectType = "CAMPAIGN_TEMPLATE_STEP"
	PublicImportErrorObjectTypePlaylist                          PublicImportErrorObjectType = "PLAYLIST"
	PublicImportErrorObjectTypeClip                              PublicImportErrorObjectType = "CLIP"
	PublicImportErrorObjectTypeCampaignBudgetItem                PublicImportErrorObjectType = "CAMPAIGN_BUDGET_ITEM"
	PublicImportErrorObjectTypeCampaignSpendItem                 PublicImportErrorObjectType = "CAMPAIGN_SPEND_ITEM"
	PublicImportErrorObjectTypeMic                               PublicImportErrorObjectType = "MIC"
	PublicImportErrorObjectTypeContentAudit                      PublicImportErrorObjectType = "CONTENT_AUDIT"
	PublicImportErrorObjectTypeContentAuditPage                  PublicImportErrorObjectType = "CONTENT_AUDIT_PAGE"
	PublicImportErrorObjectTypePlaylistFolder                    PublicImportErrorObjectType = "PLAYLIST_FOLDER"
	PublicImportErrorObjectTypeLead                              PublicImportErrorObjectType = "LEAD"
	PublicImportErrorObjectTypeAbandonedCart                     PublicImportErrorObjectType = "ABANDONED_CART"
	PublicImportErrorObjectTypeExternalWebURL                    PublicImportErrorObjectType = "EXTERNAL_WEB_URL"
	PublicImportErrorObjectTypeView                              PublicImportErrorObjectType = "VIEW"
	PublicImportErrorObjectTypeViewBlock                         PublicImportErrorObjectType = "VIEW_BLOCK"
	PublicImportErrorObjectTypeRoster                            PublicImportErrorObjectType = "ROSTER"
	PublicImportErrorObjectTypeCart                              PublicImportErrorObjectType = "CART"
	PublicImportErrorObjectTypeAutomationPlatformFlowAction      PublicImportErrorObjectType = "AUTOMATION_PLATFORM_FLOW_ACTION"
	PublicImportErrorObjectTypeSocialProfile                     PublicImportErrorObjectType = "SOCIAL_PROFILE"
	PublicImportErrorObjectTypePartnerClient                     PublicImportErrorObjectType = "PARTNER_CLIENT"
	PublicImportErrorObjectTypeRosterMember                      PublicImportErrorObjectType = "ROSTER_MEMBER"
	PublicImportErrorObjectTypeMarketingEventAttendance          PublicImportErrorObjectType = "MARKETING_EVENT_ATTENDANCE"
	PublicImportErrorObjectTypeAllPages                          PublicImportErrorObjectType = "ALL_PAGES"
	PublicImportErrorObjectTypeAIForecast                        PublicImportErrorObjectType = "AI_FORECAST"
	PublicImportErrorObjectTypeCRMPipelinesDummyType             PublicImportErrorObjectType = "CRM_PIPELINES_DUMMY_TYPE"
	PublicImportErrorObjectTypeKnowledgeArticle                  PublicImportErrorObjectType = "KNOWLEDGE_ARTICLE"
	PublicImportErrorObjectTypePropertyInfo                      PublicImportErrorObjectType = "PROPERTY_INFO"
	PublicImportErrorObjectTypeDataPrivacyConsent                PublicImportErrorObjectType = "DATA_PRIVACY_CONSENT"
	PublicImportErrorObjectTypeGoalTemplate                      PublicImportErrorObjectType = "GOAL_TEMPLATE"
	PublicImportErrorObjectTypeScoreConfiguration                PublicImportErrorObjectType = "SCORE_CONFIGURATION"
	PublicImportErrorObjectTypeAudience                          PublicImportErrorObjectType = "AUDIENCE"
	PublicImportErrorObjectTypePartnerClientRevenue              PublicImportErrorObjectType = "PARTNER_CLIENT_REVENUE"
	PublicImportErrorObjectTypeAutomationJourney                 PublicImportErrorObjectType = "AUTOMATION_JOURNEY"
	PublicImportErrorObjectTypeComboEventConfiguration           PublicImportErrorObjectType = "COMBO_EVENT_CONFIGURATION"
	PublicImportErrorObjectTypeCRMObjectsDummyType               PublicImportErrorObjectType = "CRM_OBJECTS_DUMMY_TYPE"
	PublicImportErrorObjectTypeCaseStudy                         PublicImportErrorObjectType = "CASE_STUDY"
	PublicImportErrorObjectTypeService                           PublicImportErrorObjectType = "SERVICE"
	PublicImportErrorObjectTypePodcastEpisode                    PublicImportErrorObjectType = "PODCAST_EPISODE"
	PublicImportErrorObjectTypePartnerService                    PublicImportErrorObjectType = "PARTNER_SERVICE"
	PublicImportErrorObjectTypeProspectingAgentContactAssignment PublicImportErrorObjectType = "PROSPECTING_AGENT_CONTACT_ASSIGNMENT"
	PublicImportErrorObjectTypeUnknown                           PublicImportErrorObjectType = "UNKNOWN"
)

type PublicImportMetadata struct {
	// Summarized outcomes of each row a developer attempted to import into HubSpot.
	Counters map[string]int64 `json:"counters,required"`
	// The IDs of files uploaded in the File Manager API.
	FileIDs []string `json:"fileIds,required"`
	// The lists containing the imported objects.
	ObjectLists []PublicObjectListRecord `json:"objectLists,required"`
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
	ID                  string               `json:"id,required"`
	CreatedAt           time.Time            `json:"createdAt,required" format:"date-time"`
	MappedObjectTypeIDs []string             `json:"mappedObjectTypeIds,required"`
	Metadata            PublicImportMetadata `json:"metadata,required"`
	// Whether or not the import is a list of people disqualified from receiving
	// emails.
	OptOutImport bool `json:"optOutImport,required"`
	// The status of the import.
	//
	// Any of "STARTED", "PROCESSING", "DONE", "FAILED", "CANCELED", "DEFERRED",
	// "REVERTED".
	State             PublicImportResponseState `json:"state,required"`
	UpdatedAt         time.Time                 `json:"updatedAt,required" format:"date-time"`
	ImportName        string                    `json:"importName"`
	ImportRequestJson any                       `json:"importRequestJson"`
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
	PublicImportResponseStateStarted    PublicImportResponseState = "STARTED"
	PublicImportResponseStateProcessing PublicImportResponseState = "PROCESSING"
	PublicImportResponseStateDone       PublicImportResponseState = "DONE"
	PublicImportResponseStateFailed     PublicImportResponseState = "FAILED"
	PublicImportResponseStateCanceled   PublicImportResponseState = "CANCELED"
	PublicImportResponseStateDeferred   PublicImportResponseState = "DEFERRED"
	PublicImportResponseStateReverted   PublicImportResponseState = "REVERTED"
)

type PublicImportResponseImportSource string

const (
	PublicImportResponseImportSourceAPI           PublicImportResponseImportSource = "API"
	PublicImportResponseImportSourceCRMUi         PublicImportResponseImportSource = "CRM_UI"
	PublicImportResponseImportSourceImport        PublicImportResponseImportSource = "IMPORT"
	PublicImportResponseImportSourceMobileAndroid PublicImportResponseImportSource = "MOBILE_ANDROID"
	PublicImportResponseImportSourceMobileIos     PublicImportResponseImportSource = "MOBILE_IOS"
	PublicImportResponseImportSourceSalesforce    PublicImportResponseImportSource = "SALESFORCE"
)

type PublicObjectListRecord struct {
	// The ID of the list containing the imported objects.
	ListID string `json:"listId,required"`
	// The type of object contained in the list.
	ObjectType string `json:"objectType,required"`
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
	After  param.Opt[string] `query:"after,omitzero" json:"-"`
	Before param.Opt[string] `query:"before,omitzero" json:"-"`
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
	After param.Opt[string] `query:"after,omitzero" json:"-"`
	// Set to True to receive a message explaining the error.
	IncludeErrorMessage param.Opt[bool] `query:"includeErrorMessage,omitzero" json:"-"`
	// Set to True to receive the data values for the errored row.
	IncludeRowData param.Opt[bool] `query:"includeRowData,omitzero" json:"-"`
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
