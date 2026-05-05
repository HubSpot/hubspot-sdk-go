// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cms

import (
	"context"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/HubSpot/hubspot-sdk-go/internal/apijson"
	"github.com/HubSpot/hubspot-sdk-go/internal/apiquery"
	shimjson "github.com/HubSpot/hubspot-sdk-go/internal/encoding/json"
	"github.com/HubSpot/hubspot-sdk-go/internal/requestconfig"
	"github.com/HubSpot/hubspot-sdk-go/option"
	"github.com/HubSpot/hubspot-sdk-go/packages/pagination"
	"github.com/HubSpot/hubspot-sdk-go/packages/param"
	"github.com/HubSpot/hubspot-sdk-go/packages/respjson"
	"github.com/HubSpot/hubspot-sdk-go/shared"
)

// AuditLogService contains methods and other services that help with interacting
// with the hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAuditLogService] method instead.
type AuditLogService struct {
	options []option.RequestOption
}

// NewAuditLogService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAuditLogService(opts ...option.RequestOption) (r AuditLogService) {
	r = AuditLogService{}
	r.options = opts
	return
}

func (r *AuditLogService) List(ctx context.Context, query AuditLogListParams, opts ...option.RequestOption) (res *pagination.Page[PublicAuditLog], err error) {
	var raw *http.Response
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "cms/audit-logs/2026-03"
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

func (r *AuditLogService) ListAutoPaging(ctx context.Context, query AuditLogListParams, opts ...option.RequestOption) *pagination.PageAutoPager[PublicAuditLog] {
	return pagination.NewPageAutoPager(r.List(ctx, query, opts...))
}

func (r *AuditLogService) Export(ctx context.Context, body AuditLogExportParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "cms/audit-logs/2026-03/export"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return err
}

// The property ObjectType is required.
type CmsAuditLoggingExportFiltersParam struct {
	ObjectType []string `json:"objectType,omitzero" api:"required"`
	paramObj
}

func (r CmsAuditLoggingExportFiltersParam) MarshalJSON() (data []byte, err error) {
	type shadow CmsAuditLoggingExportFiltersParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CmsAuditLoggingExportFiltersParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Email, Format, PortalID, RecipientUserIDs,
// ShouldMarkExportFileAsSensitive, Type are required.
type CmsAuditLoggingExportSettingsParam struct {
	Email string `json:"email" api:"required"`
	// Any of "CSV", "XLS", "XLSX".
	Format                          CmsAuditLoggingExportSettingsFormat `json:"format,omitzero" api:"required"`
	PortalID                        int64                               `json:"portalId" api:"required"`
	RecipientUserIDs                []int64                             `json:"recipientUserIds,omitzero" api:"required"`
	ShouldMarkExportFileAsSensitive bool                                `json:"shouldMarkExportFileAsSensitive" api:"required"`
	Type                            string                              `json:"type" api:"required"`
	Partition                       param.Opt[int64]                    `json:"partition,omitzero"`
	UserID                          param.Opt[int64]                    `json:"userId,omitzero"`
	UserTimeZone                    param.Opt[string]                   `json:"userTimeZone,omitzero"`
	Filters                         CmsAuditLoggingExportFiltersParam   `json:"filters,omitzero"`
	paramObj
}

func (r CmsAuditLoggingExportSettingsParam) MarshalJSON() (data []byte, err error) {
	type shadow CmsAuditLoggingExportSettingsParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CmsAuditLoggingExportSettingsParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CmsAuditLoggingExportSettingsFormat string

const (
	CmsAuditLoggingExportSettingsFormatCsv  CmsAuditLoggingExportSettingsFormat = "CSV"
	CmsAuditLoggingExportSettingsFormatXls  CmsAuditLoggingExportSettingsFormat = "XLS"
	CmsAuditLoggingExportSettingsFormatXlsx CmsAuditLoggingExportSettingsFormat = "XLSX"
)

type CollectionResponsePublicAuditLog struct {
	Results []PublicAuditLog `json:"results" api:"required"`
	Paging  shared.Paging    `json:"paging"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Results     respjson.Field
		Paging      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CollectionResponsePublicAuditLog) RawJSON() string { return r.JSON.raw }
func (r *CollectionResponsePublicAuditLog) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PublicAuditLog struct {
	// The type of event that took place (CREATED, UPDATED, PUBLISHED, DELETED,
	// UNPUBLISHED).
	//
	// Any of "CREATED", "DELETED", "PUBLISHED", "RESTORE", "UNPUBLISHED", "UPDATED".
	Event PublicAuditLogEvent `json:"event" api:"required"`
	// The name of the user who caused the event.
	FullName string `json:"fullName" api:"required"`
	// Supplementary metadata associated with the audit log entry. It provides
	// additional context about the audited event (ex: rows deleted/updated for a HubDB
	// event, the specific fields that were changed for a Content Settings event).
	Meta any `json:"meta" api:"required"`
	// The ID of the object.
	ObjectID string `json:"objectId" api:"required"`
	// The internal name of the object in HubSpot.
	ObjectName string `json:"objectName" api:"required"`
	// The type of the object (BLOG, LANDING_PAGE, DOMAIN, HUBDB_TABLE etc.)
	//
	// Any of "BLOG", "BLOG_POST", "CASE_STUDY", "CONTENT_SETTINGS", "CSS", "CTA",
	// "DOMAIN", "EMAIL", "FILE", "GLOBAL_MODULE", "HUBDB_TABLE", "JS",
	// "KNOWLEDGE_BASE", "KNOWLEDGE_BASE_ARTICLE", "LANDING_PAGE", "MODULE", "PODCAST",
	// "QUOTE", "SERVERLESS_FUNCTION", "TEMPLATE", "THEME", "URL_MAPPING",
	// "WEB_INTERACTIVE", "WEBSITE_PAGE".
	ObjectType PublicAuditLogObjectType `json:"objectType" api:"required"`
	// The timestamp at which the event occurred.
	Timestamp time.Time `json:"timestamp" api:"required" format:"date-time"`
	// The ID of the user who caused the event.
	UserID string `json:"userId" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Event       respjson.Field
		FullName    respjson.Field
		Meta        respjson.Field
		ObjectID    respjson.Field
		ObjectName  respjson.Field
		ObjectType  respjson.Field
		Timestamp   respjson.Field
		UserID      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PublicAuditLog) RawJSON() string { return r.JSON.raw }
func (r *PublicAuditLog) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The type of event that took place (CREATED, UPDATED, PUBLISHED, DELETED,
// UNPUBLISHED).
type PublicAuditLogEvent string

const (
	PublicAuditLogEventCreated     PublicAuditLogEvent = "CREATED"
	PublicAuditLogEventDeleted     PublicAuditLogEvent = "DELETED"
	PublicAuditLogEventPublished   PublicAuditLogEvent = "PUBLISHED"
	PublicAuditLogEventRestore     PublicAuditLogEvent = "RESTORE"
	PublicAuditLogEventUnpublished PublicAuditLogEvent = "UNPUBLISHED"
	PublicAuditLogEventUpdated     PublicAuditLogEvent = "UPDATED"
)

// The type of the object (BLOG, LANDING_PAGE, DOMAIN, HUBDB_TABLE etc.)
type PublicAuditLogObjectType string

const (
	PublicAuditLogObjectTypeBlog                 PublicAuditLogObjectType = "BLOG"
	PublicAuditLogObjectTypeBlogPost             PublicAuditLogObjectType = "BLOG_POST"
	PublicAuditLogObjectTypeCaseStudy            PublicAuditLogObjectType = "CASE_STUDY"
	PublicAuditLogObjectTypeContentSettings      PublicAuditLogObjectType = "CONTENT_SETTINGS"
	PublicAuditLogObjectTypeCss                  PublicAuditLogObjectType = "CSS"
	PublicAuditLogObjectTypeCta                  PublicAuditLogObjectType = "CTA"
	PublicAuditLogObjectTypeDomain               PublicAuditLogObjectType = "DOMAIN"
	PublicAuditLogObjectTypeEmail                PublicAuditLogObjectType = "EMAIL"
	PublicAuditLogObjectTypeFile                 PublicAuditLogObjectType = "FILE"
	PublicAuditLogObjectTypeGlobalModule         PublicAuditLogObjectType = "GLOBAL_MODULE"
	PublicAuditLogObjectTypeHubdbTable           PublicAuditLogObjectType = "HUBDB_TABLE"
	PublicAuditLogObjectTypeJs                   PublicAuditLogObjectType = "JS"
	PublicAuditLogObjectTypeKnowledgeBase        PublicAuditLogObjectType = "KNOWLEDGE_BASE"
	PublicAuditLogObjectTypeKnowledgeBaseArticle PublicAuditLogObjectType = "KNOWLEDGE_BASE_ARTICLE"
	PublicAuditLogObjectTypeLandingPage          PublicAuditLogObjectType = "LANDING_PAGE"
	PublicAuditLogObjectTypeModule               PublicAuditLogObjectType = "MODULE"
	PublicAuditLogObjectTypePodcast              PublicAuditLogObjectType = "PODCAST"
	PublicAuditLogObjectTypeQuote                PublicAuditLogObjectType = "QUOTE"
	PublicAuditLogObjectTypeServerlessFunction   PublicAuditLogObjectType = "SERVERLESS_FUNCTION"
	PublicAuditLogObjectTypeTemplate             PublicAuditLogObjectType = "TEMPLATE"
	PublicAuditLogObjectTypeTheme                PublicAuditLogObjectType = "THEME"
	PublicAuditLogObjectTypeURLMapping           PublicAuditLogObjectType = "URL_MAPPING"
	PublicAuditLogObjectTypeWebInteractive       PublicAuditLogObjectType = "WEB_INTERACTIVE"
	PublicAuditLogObjectTypeWebsitePage          PublicAuditLogObjectType = "WEBSITE_PAGE"
)

type AuditLogListParams struct {
	// The paging cursor token of the last successfully read resource will be returned
	// as the `paging.next.after` JSON property of a paged response containing more
	// results.
	After  param.Opt[string] `query:"after,omitzero" json:"-"`
	Before param.Opt[string] `query:"before,omitzero" json:"-"`
	// The maximum number of results to display per page.
	Limit      param.Opt[int64] `query:"limit,omitzero" json:"-"`
	EventType  []string         `query:"eventType,omitzero" json:"-"`
	ObjectID   []string         `query:"objectId,omitzero" json:"-"`
	ObjectType []string         `query:"objectType,omitzero" json:"-"`
	Sort       []string         `query:"sort,omitzero" json:"-"`
	UserID     []string         `query:"userId,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [AuditLogListParams]'s query parameters as `url.Values`.
func (r AuditLogListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AuditLogExportParams struct {
	CmsAuditLoggingExportSettings CmsAuditLoggingExportSettingsParam
	paramObj
}

func (r AuditLogExportParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.CmsAuditLoggingExportSettings)
}
func (r *AuditLogExportParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
