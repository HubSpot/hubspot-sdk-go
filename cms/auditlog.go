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
)

// AuditLogService contains methods and other services that help with interacting
// with the Hubspot API.
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
func (r *AuditLogService) List(ctx context.Context, query AuditLogListParams, opts ...option.RequestOption) (res *pagination.Page[AuditLogListResponse], err error) {
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
func (r *AuditLogService) ListAutoPaging(ctx context.Context, query AuditLogListParams, opts ...option.RequestOption) *pagination.PageAutoPager[AuditLogListResponse] {
	return pagination.NewPageAutoPager(r.List(ctx, query, opts...))
}

type AuditLogListResponse struct {
	// The type of event that took place (CREATED, UPDATED, PUBLISHED, DELETED,
	// UNPUBLISHED).
	//
	// Any of "CREATED", "UPDATED", "PUBLISHED", "DELETED", "UNPUBLISHED", "RESTORE".
	Event AuditLogListResponseEvent `json:"event,required"`
	// The name of the user who caused the event.
	FullName string `json:"fullName,required"`
	// The ID of the object.
	ObjectID string `json:"objectId,required"`
	// The internal name of the object in HubSpot.
	ObjectName string `json:"objectName,required"`
	// The type of the object (BLOG, LANDING_PAGE, DOMAIN, HUBDB_TABLE etc.)
	//
	// Any of "BLOG", "BLOG_POST", "LANDING_PAGE", "WEBSITE_PAGE", "TEMPLATE",
	// "MODULE", "GLOBAL_MODULE", "SERVERLESS_FUNCTION", "DOMAIN", "URL_MAPPING",
	// "EMAIL", "CONTENT_SETTINGS", "HUBDB_TABLE", "KNOWLEDGE_BASE_ARTICLE",
	// "KNOWLEDGE_BASE", "THEME", "CSS", "JS", "CTA", "FILE".
	ObjectType AuditLogListResponseObjectType `json:"objectType,required"`
	// The timestamp at which the event occurred.
	Timestamp time.Time `json:"timestamp,required" format:"date-time"`
	// The ID of the user who caused the event.
	UserID string `json:"userId,required"`
	Meta   any    `json:"meta"`
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
func (r AuditLogListResponse) RawJSON() string { return r.JSON.raw }
func (r *AuditLogListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The type of event that took place (CREATED, UPDATED, PUBLISHED, DELETED,
// UNPUBLISHED).
type AuditLogListResponseEvent string

const (
	AuditLogListResponseEventCreated     AuditLogListResponseEvent = "CREATED"
	AuditLogListResponseEventUpdated     AuditLogListResponseEvent = "UPDATED"
	AuditLogListResponseEventPublished   AuditLogListResponseEvent = "PUBLISHED"
	AuditLogListResponseEventDeleted     AuditLogListResponseEvent = "DELETED"
	AuditLogListResponseEventUnpublished AuditLogListResponseEvent = "UNPUBLISHED"
	AuditLogListResponseEventRestore     AuditLogListResponseEvent = "RESTORE"
)

// The type of the object (BLOG, LANDING_PAGE, DOMAIN, HUBDB_TABLE etc.)
type AuditLogListResponseObjectType string

const (
	AuditLogListResponseObjectTypeBlog                 AuditLogListResponseObjectType = "BLOG"
	AuditLogListResponseObjectTypeBlogPost             AuditLogListResponseObjectType = "BLOG_POST"
	AuditLogListResponseObjectTypeLandingPage          AuditLogListResponseObjectType = "LANDING_PAGE"
	AuditLogListResponseObjectTypeWebsitePage          AuditLogListResponseObjectType = "WEBSITE_PAGE"
	AuditLogListResponseObjectTypeTemplate             AuditLogListResponseObjectType = "TEMPLATE"
	AuditLogListResponseObjectTypeModule               AuditLogListResponseObjectType = "MODULE"
	AuditLogListResponseObjectTypeGlobalModule         AuditLogListResponseObjectType = "GLOBAL_MODULE"
	AuditLogListResponseObjectTypeServerlessFunction   AuditLogListResponseObjectType = "SERVERLESS_FUNCTION"
	AuditLogListResponseObjectTypeDomain               AuditLogListResponseObjectType = "DOMAIN"
	AuditLogListResponseObjectTypeURLMapping           AuditLogListResponseObjectType = "URL_MAPPING"
	AuditLogListResponseObjectTypeEmail                AuditLogListResponseObjectType = "EMAIL"
	AuditLogListResponseObjectTypeContentSettings      AuditLogListResponseObjectType = "CONTENT_SETTINGS"
	AuditLogListResponseObjectTypeHubdbTable           AuditLogListResponseObjectType = "HUBDB_TABLE"
	AuditLogListResponseObjectTypeKnowledgeBaseArticle AuditLogListResponseObjectType = "KNOWLEDGE_BASE_ARTICLE"
	AuditLogListResponseObjectTypeKnowledgeBase        AuditLogListResponseObjectType = "KNOWLEDGE_BASE"
	AuditLogListResponseObjectTypeTheme                AuditLogListResponseObjectType = "THEME"
	AuditLogListResponseObjectTypeCss                  AuditLogListResponseObjectType = "CSS"
	AuditLogListResponseObjectTypeJs                   AuditLogListResponseObjectType = "JS"
	AuditLogListResponseObjectTypeCta                  AuditLogListResponseObjectType = "CTA"
	AuditLogListResponseObjectTypeFile                 AuditLogListResponseObjectType = "FILE"
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
