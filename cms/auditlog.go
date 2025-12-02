// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cms

import (
	"context"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/stainless-sdks/hubspot-sdk-go/internal/apijson"
	"github.com/stainless-sdks/hubspot-sdk-go/internal/apiquery"
	"github.com/stainless-sdks/hubspot-sdk-go/internal/requestconfig"
	"github.com/stainless-sdks/hubspot-sdk-go/option"
	"github.com/stainless-sdks/hubspot-sdk-go/packages/pagination"
	"github.com/stainless-sdks/hubspot-sdk-go/packages/param"
	"github.com/stainless-sdks/hubspot-sdk-go/packages/respjson"
	"github.com/stainless-sdks/hubspot-sdk-go/shared"
)

// AuditLogService contains methods and other services that help with interacting
// with the hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAuditLogService] method instead.
type AuditLogService struct {
	Options []option.RequestOption
}

// NewAuditLogService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAuditLogService(opts ...option.RequestOption) (r AuditLogService) {
	r = AuditLogService{}
	r.Options = opts
	return
}

// Returns audit logs based on filters.
func (r *AuditLogService) List(ctx context.Context, query AuditLogListParams, opts ...option.RequestOption) (res *pagination.Page[PublicAuditLog], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "cms/v3/audit-logs/"
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

// Returns audit logs based on filters.
func (r *AuditLogService) ListAutoPaging(ctx context.Context, query AuditLogListParams, opts ...option.RequestOption) *pagination.PageAutoPager[PublicAuditLog] {
	return pagination.NewPageAutoPager(r.List(ctx, query, opts...))
}

// The collection of audit logs.
type CollectionResponsePublicAuditLog struct {
	Results []PublicAuditLog `json:"results,required"`
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
	Event PublicAuditLogEvent `json:"event,required"`
	// The name of the user who caused the event.
	FullName string `json:"fullName,required"`
	// The ID of the object.
	ObjectID string `json:"objectId,required"`
	// The internal name of the object in HubSpot.
	ObjectName string `json:"objectName,required"`
	// The type of the object (BLOG, LANDING_PAGE, DOMAIN, HUBDB_TABLE etc.)
	//
	// Any of "BLOG", "BLOG_POST", "CONTENT_SETTINGS", "CSS", "CTA", "DOMAIN", "EMAIL",
	// "FILE", "GLOBAL_MODULE", "HUBDB_TABLE", "JS", "KNOWLEDGE_BASE",
	// "KNOWLEDGE_BASE_ARTICLE", "LANDING_PAGE", "MODULE", "SERVERLESS_FUNCTION",
	// "TEMPLATE", "THEME", "URL_MAPPING", "WEBSITE_PAGE".
	ObjectType PublicAuditLogObjectType `json:"objectType,required"`
	// The timestamp at which the event occurred.
	Timestamp time.Time `json:"timestamp,required" format:"date-time"`
	// The ID of the user who caused the event.
	UserID string `json:"userId,required"`
	// Supplementary metadata associated with the audit log entry. It provides
	// additional context about the audited event (ex: rows deleted/updated for a HubDB
	// event, the specific fields that were changed for a Content Settings event).
	Meta any `json:"meta"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Event       respjson.Field
		FullName    respjson.Field
		ObjectID    respjson.Field
		ObjectName  respjson.Field
		ObjectType  respjson.Field
		Timestamp   respjson.Field
		UserID      respjson.Field
		Meta        respjson.Field
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
	PublicAuditLogObjectTypeServerlessFunction   PublicAuditLogObjectType = "SERVERLESS_FUNCTION"
	PublicAuditLogObjectTypeTemplate             PublicAuditLogObjectType = "TEMPLATE"
	PublicAuditLogObjectTypeTheme                PublicAuditLogObjectType = "THEME"
	PublicAuditLogObjectTypeURLMapping           PublicAuditLogObjectType = "URL_MAPPING"
	PublicAuditLogObjectTypeWebsitePage          PublicAuditLogObjectType = "WEBSITE_PAGE"
)

type AuditLogListParams struct {
	// Timestamp after which audit logs will be returned
	After param.Opt[string] `query:"after,omitzero" json:"-"`
	// Timestamp before which audit logs will be returned
	Before param.Opt[string] `query:"before,omitzero" json:"-"`
	// The number of logs to return.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Comma separated list of event types to filter by (CREATED, UPDATED, PUBLISHED,
	// DELETED, UNPUBLISHED).
	EventType []string `query:"eventType,omitzero" json:"-"`
	// Comma separated list of object ids to filter by.
	ObjectID []string `query:"objectId,omitzero" json:"-"`
	// Comma separated list of object types to filter by (BLOG, LANDING_PAGE, DOMAIN,
	// HUBDB_TABLE etc.)
	ObjectType []string `query:"objectType,omitzero" json:"-"`
	// The sort direction for the audit logs. (Can only sort by timestamp).
	Sort []string `query:"sort,omitzero" json:"-"`
	// Comma separated list of user ids to filter by.
	UserID []string `query:"userId,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [AuditLogListParams]'s query parameters as `url.Values`.
func (r AuditLogListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
