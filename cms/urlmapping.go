// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cms

import (
	"context"
	"fmt"
	"net/http"
	"slices"

	"github.com/stainless-sdks/hubspot-sdk-go/internal/apijson"
	shimjson "github.com/stainless-sdks/hubspot-sdk-go/internal/encoding/json"
	"github.com/stainless-sdks/hubspot-sdk-go/internal/requestconfig"
	"github.com/stainless-sdks/hubspot-sdk-go/option"
	"github.com/stainless-sdks/hubspot-sdk-go/packages/param"
)

// URLMappingService contains methods and other services that help with interacting
// with the hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewURLMappingService] method instead.
type URLMappingService struct {
	options []option.RequestOption
}

// NewURLMappingService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewURLMappingService(opts ...option.RequestOption) (r URLMappingService) {
	r = URLMappingService{}
	r.options = opts
	return
}

// Create a new URL mapping in your HubSpot account. This endpoint allows you to
// define URL redirections and mappings, which can be useful for managing site
// navigation and SEO. The request body must include all required properties of the
// UrlMapping schema.
func (r *URLMappingService) New(ctx context.Context, body URLMappingNewParams, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "url-mappings/2026-03/url-mappings"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Retrieve a list of URL mappings from the HubSpot account. This endpoint provides
// access to URL mapping configurations, which can be used to manage and redirect
// URLs within the HubSpot CMS. It is useful for understanding how URLs are
// structured and redirected in your content management setup.
func (r *URLMappingService) List(ctx context.Context, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "url-mappings/2026-03/url-mappings"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Delete a specific URL mapping in your HubSpot account using its unique
// identifier. This operation will remove the URL mapping permanently, and it
// requires appropriate write and delete permissions.
func (r *URLMappingService) Delete(ctx context.Context, id int64, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := fmt.Sprintf("url-mappings/2026-03/url-mappings/%v", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Retrieve a specific URL mapping by its unique identifier. This endpoint is
// useful for obtaining details about a particular URL mapping configuration within
// your HubSpot account. It requires the ID of the URL mapping as a path parameter.
func (r *URLMappingService) Get(ctx context.Context, id int64, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := fmt.Sprintf("url-mappings/2026-03/url-mappings/%v", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// The properties ID, CdnPurgeEmbargoTime, ContentGroupID, CosObjectType, Created,
// CreatedByID, DeletedAt, Destination, InternallyCreated, IsActive,
// IsMatchFullURL, IsMatchQueryString, IsOnlyAfterNotFound, IsPattern,
// IsProtocolAgnostic, IsRegex, IsTrailingSlashOptional, Label, LastUsedAt, Name,
// Note, PortalID, Precedence, RedirectStyle, RoutePrefix, Updated, UpdatedByID are
// required.
type URLMappingsURLMappingParam struct {
	// The unique identifier for the URL mapping, represented as a 64-bit integer.
	ID int64 `json:"id" api:"required"`
	// A Unix timestamp in milliseconds indicating the embargo time for CDN purge
	// related to the URL mapping.
	CdnPurgeEmbargoTime int64 `json:"cdnPurgeEmbargoTime" api:"required"`
	// A 64-bit integer representing the content group associated with the URL mapping.
	ContentGroupID int64 `json:"contentGroupId" api:"required"`
	// A string representing the type of content object associated with the URL
	// mapping. Valid values include various content types such as 'CONTENT', 'LAYOUT',
	// 'FILE', etc.
	//
	// Any of "ACCESS_GROUP_MEMBERSHIP", "APP_PAGE", "BLOCK", "BLOG", "BLOG_AUTHOR",
	// "BRAND_BUSINESS_UNIT", "BRAND_SETTINGS", "CONTACT_MEMBERSHIP", "CONTENT",
	// "CONTENT_EMBED", "CONTENT_FOLDER", "CONTENT_GROUP", "CRM_OBJECT",
	// "CRM_OBJECT_TYPE", "CUSTOM_WIDGET", "CUSTOMER_PORTAL", "DATA_QUERY",
	// "DESIGN_FOLDER", "DOMAIN", "DOMAIN_SETTINGS", "EMAIL_ADDRESS",
	// "EXTENSION_RESOURCE", "FILE", "FOLDER", "FOLLOW_ME", "FORM", "GLOBAL_CONTENT",
	// "HUBDB_TABLE", "HUBDB_TABLE_ROW", "IMAGE", "JS_PROJECT_COMPONENT",
	// "KNOWLEDGE_BASE", "KNOWLEDGE_CATEGORY", "KNOWLEDGE_CATEGORY_TRANSLATION",
	// "KNOWLEDGE_HOMEPAGE_CATEGORY", "LAYOUT", "LAYOUT_SECTION", "LIST_MEMBERSHIP",
	// "MARKETPLACE_LISTING", "PASSWORD_PROTECTED", "PAYMENT", "PERSONALIZATION_TOKEN",
	// "PLACEMENT", "PROJECT", "QUOTE_TEMPLATE", "RAW_ASSET", "REDIRECT_URL",
	// "SECTION", "SERVERLESS_FUNCTION", "SITE_MAP", "SITE_MENU", "SITE_SETTINGS",
	// "SUBSCRIPTIONS_SETTINGS", "TAG", "THEME", "THEME_SETTINGS",
	// "UNRESTRICTED_ACCESS", "URL_MAPPING", "VIDEO_PLAYER", "WIDGET", "WORKFLOW".
	CosObjectType URLMappingsURLMappingCosObjectType `json:"cosObjectType,omitzero" api:"required"`
	// A Unix timestamp in milliseconds indicating when the URL mapping was created.
	Created int64 `json:"created" api:"required"`
	// The identifier of the user who created the URL mapping.
	CreatedByID int64 `json:"createdById" api:"required"`
	// A Unix timestamp in milliseconds indicating when the URL mapping was deleted.
	DeletedAt int64 `json:"deletedAt" api:"required"`
	// The destination URL to which the routePrefix is redirected.
	Destination string `json:"destination" api:"required"`
	// A boolean indicating if the URL mapping was created internally by the system.
	InternallyCreated bool `json:"internallyCreated" api:"required"`
	// A boolean indicating if the URL mapping is currently active.
	IsActive bool `json:"isActive" api:"required"`
	// A boolean indicating if the full URL should be matched.
	IsMatchFullURL bool `json:"isMatchFullUrl" api:"required"`
	// A boolean indicating if the query string should be matched.
	IsMatchQueryString bool `json:"isMatchQueryString" api:"required"`
	// A boolean indicating if the mapping should only be applied after a 404 Not Found
	// response.
	IsOnlyAfterNotFound bool `json:"isOnlyAfterNotFound" api:"required"`
	// A boolean indicating if the routePrefix is a pattern.
	IsPattern bool `json:"isPattern" api:"required"`
	// A boolean indicating if the mapping should ignore the URL protocol (http/https).
	IsProtocolAgnostic bool `json:"isProtocolAgnostic" api:"required"`
	// A boolean indicating if the routePrefix should be treated as a regular
	// expression.
	IsRegex bool `json:"isRegex" api:"required"`
	// A boolean indicating if the trailing slash in the URL is optional.
	IsTrailingSlashOptional bool `json:"isTrailingSlashOptional" api:"required"`
	// A label for the URL mapping.
	Label      string `json:"label" api:"required"`
	LastUsedAt int64  `json:"lastUsedAt" api:"required"`
	// The name of the URL mapping.
	Name string `json:"name" api:"required"`
	// A string containing notes about the URL mapping.
	Note string `json:"note" api:"required"`
	// The identifier for the HubSpot portal associated with this URL mapping.
	PortalID int64 `json:"portalId" api:"required"`
	// An integer representing the precedence of the URL mapping, used to determine
	// order of evaluation.
	Precedence int64 `json:"precedence" api:"required"`
	// An integer representing the style of redirection used.
	RedirectStyle int64 `json:"redirectStyle" api:"required"`
	// The prefix of the URL path that is being mapped.
	RoutePrefix string `json:"routePrefix" api:"required"`
	// A Unix timestamp in milliseconds indicating when the URL mapping was last
	// updated.
	Updated int64 `json:"updated" api:"required"`
	// The identifier of the user who last updated the URL mapping.
	UpdatedByID int64 `json:"updatedById" api:"required"`
	paramObj
}

func (r URLMappingsURLMappingParam) MarshalJSON() (data []byte, err error) {
	type shadow URLMappingsURLMappingParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *URLMappingsURLMappingParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A string representing the type of content object associated with the URL
// mapping. Valid values include various content types such as 'CONTENT', 'LAYOUT',
// 'FILE', etc.
type URLMappingsURLMappingCosObjectType string

const (
	URLMappingsURLMappingCosObjectTypeAccessGroupMembership        URLMappingsURLMappingCosObjectType = "ACCESS_GROUP_MEMBERSHIP"
	URLMappingsURLMappingCosObjectTypeAppPage                      URLMappingsURLMappingCosObjectType = "APP_PAGE"
	URLMappingsURLMappingCosObjectTypeBlock                        URLMappingsURLMappingCosObjectType = "BLOCK"
	URLMappingsURLMappingCosObjectTypeBlog                         URLMappingsURLMappingCosObjectType = "BLOG"
	URLMappingsURLMappingCosObjectTypeBlogAuthor                   URLMappingsURLMappingCosObjectType = "BLOG_AUTHOR"
	URLMappingsURLMappingCosObjectTypeBrandBusinessUnit            URLMappingsURLMappingCosObjectType = "BRAND_BUSINESS_UNIT"
	URLMappingsURLMappingCosObjectTypeBrandSettings                URLMappingsURLMappingCosObjectType = "BRAND_SETTINGS"
	URLMappingsURLMappingCosObjectTypeContactMembership            URLMappingsURLMappingCosObjectType = "CONTACT_MEMBERSHIP"
	URLMappingsURLMappingCosObjectTypeContent                      URLMappingsURLMappingCosObjectType = "CONTENT"
	URLMappingsURLMappingCosObjectTypeContentEmbed                 URLMappingsURLMappingCosObjectType = "CONTENT_EMBED"
	URLMappingsURLMappingCosObjectTypeContentFolder                URLMappingsURLMappingCosObjectType = "CONTENT_FOLDER"
	URLMappingsURLMappingCosObjectTypeContentGroup                 URLMappingsURLMappingCosObjectType = "CONTENT_GROUP"
	URLMappingsURLMappingCosObjectTypeCrmObject                    URLMappingsURLMappingCosObjectType = "CRM_OBJECT"
	URLMappingsURLMappingCosObjectTypeCrmObjectType                URLMappingsURLMappingCosObjectType = "CRM_OBJECT_TYPE"
	URLMappingsURLMappingCosObjectTypeCustomWidget                 URLMappingsURLMappingCosObjectType = "CUSTOM_WIDGET"
	URLMappingsURLMappingCosObjectTypeCustomerPortal               URLMappingsURLMappingCosObjectType = "CUSTOMER_PORTAL"
	URLMappingsURLMappingCosObjectTypeDataQuery                    URLMappingsURLMappingCosObjectType = "DATA_QUERY"
	URLMappingsURLMappingCosObjectTypeDesignFolder                 URLMappingsURLMappingCosObjectType = "DESIGN_FOLDER"
	URLMappingsURLMappingCosObjectTypeDomain                       URLMappingsURLMappingCosObjectType = "DOMAIN"
	URLMappingsURLMappingCosObjectTypeDomainSettings               URLMappingsURLMappingCosObjectType = "DOMAIN_SETTINGS"
	URLMappingsURLMappingCosObjectTypeEmailAddress                 URLMappingsURLMappingCosObjectType = "EMAIL_ADDRESS"
	URLMappingsURLMappingCosObjectTypeExtensionResource            URLMappingsURLMappingCosObjectType = "EXTENSION_RESOURCE"
	URLMappingsURLMappingCosObjectTypeFile                         URLMappingsURLMappingCosObjectType = "FILE"
	URLMappingsURLMappingCosObjectTypeFolder                       URLMappingsURLMappingCosObjectType = "FOLDER"
	URLMappingsURLMappingCosObjectTypeFollowMe                     URLMappingsURLMappingCosObjectType = "FOLLOW_ME"
	URLMappingsURLMappingCosObjectTypeForm                         URLMappingsURLMappingCosObjectType = "FORM"
	URLMappingsURLMappingCosObjectTypeGlobalContent                URLMappingsURLMappingCosObjectType = "GLOBAL_CONTENT"
	URLMappingsURLMappingCosObjectTypeHubdbTable                   URLMappingsURLMappingCosObjectType = "HUBDB_TABLE"
	URLMappingsURLMappingCosObjectTypeHubdbTableRow                URLMappingsURLMappingCosObjectType = "HUBDB_TABLE_ROW"
	URLMappingsURLMappingCosObjectTypeImage                        URLMappingsURLMappingCosObjectType = "IMAGE"
	URLMappingsURLMappingCosObjectTypeJsProjectComponent           URLMappingsURLMappingCosObjectType = "JS_PROJECT_COMPONENT"
	URLMappingsURLMappingCosObjectTypeKnowledgeBase                URLMappingsURLMappingCosObjectType = "KNOWLEDGE_BASE"
	URLMappingsURLMappingCosObjectTypeKnowledgeCategory            URLMappingsURLMappingCosObjectType = "KNOWLEDGE_CATEGORY"
	URLMappingsURLMappingCosObjectTypeKnowledgeCategoryTranslation URLMappingsURLMappingCosObjectType = "KNOWLEDGE_CATEGORY_TRANSLATION"
	URLMappingsURLMappingCosObjectTypeKnowledgeHomepageCategory    URLMappingsURLMappingCosObjectType = "KNOWLEDGE_HOMEPAGE_CATEGORY"
	URLMappingsURLMappingCosObjectTypeLayout                       URLMappingsURLMappingCosObjectType = "LAYOUT"
	URLMappingsURLMappingCosObjectTypeLayoutSection                URLMappingsURLMappingCosObjectType = "LAYOUT_SECTION"
	URLMappingsURLMappingCosObjectTypeListMembership               URLMappingsURLMappingCosObjectType = "LIST_MEMBERSHIP"
	URLMappingsURLMappingCosObjectTypeMarketplaceListing           URLMappingsURLMappingCosObjectType = "MARKETPLACE_LISTING"
	URLMappingsURLMappingCosObjectTypePasswordProtected            URLMappingsURLMappingCosObjectType = "PASSWORD_PROTECTED"
	URLMappingsURLMappingCosObjectTypePayment                      URLMappingsURLMappingCosObjectType = "PAYMENT"
	URLMappingsURLMappingCosObjectTypePersonalizationToken         URLMappingsURLMappingCosObjectType = "PERSONALIZATION_TOKEN"
	URLMappingsURLMappingCosObjectTypePlacement                    URLMappingsURLMappingCosObjectType = "PLACEMENT"
	URLMappingsURLMappingCosObjectTypeProject                      URLMappingsURLMappingCosObjectType = "PROJECT"
	URLMappingsURLMappingCosObjectTypeQuoteTemplate                URLMappingsURLMappingCosObjectType = "QUOTE_TEMPLATE"
	URLMappingsURLMappingCosObjectTypeRawAsset                     URLMappingsURLMappingCosObjectType = "RAW_ASSET"
	URLMappingsURLMappingCosObjectTypeRedirectURL                  URLMappingsURLMappingCosObjectType = "REDIRECT_URL"
	URLMappingsURLMappingCosObjectTypeSection                      URLMappingsURLMappingCosObjectType = "SECTION"
	URLMappingsURLMappingCosObjectTypeServerlessFunction           URLMappingsURLMappingCosObjectType = "SERVERLESS_FUNCTION"
	URLMappingsURLMappingCosObjectTypeSiteMap                      URLMappingsURLMappingCosObjectType = "SITE_MAP"
	URLMappingsURLMappingCosObjectTypeSiteMenu                     URLMappingsURLMappingCosObjectType = "SITE_MENU"
	URLMappingsURLMappingCosObjectTypeSiteSettings                 URLMappingsURLMappingCosObjectType = "SITE_SETTINGS"
	URLMappingsURLMappingCosObjectTypeSubscriptionsSettings        URLMappingsURLMappingCosObjectType = "SUBSCRIPTIONS_SETTINGS"
	URLMappingsURLMappingCosObjectTypeTag                          URLMappingsURLMappingCosObjectType = "TAG"
	URLMappingsURLMappingCosObjectTypeTheme                        URLMappingsURLMappingCosObjectType = "THEME"
	URLMappingsURLMappingCosObjectTypeThemeSettings                URLMappingsURLMappingCosObjectType = "THEME_SETTINGS"
	URLMappingsURLMappingCosObjectTypeUnrestrictedAccess           URLMappingsURLMappingCosObjectType = "UNRESTRICTED_ACCESS"
	URLMappingsURLMappingCosObjectTypeURLMapping                   URLMappingsURLMappingCosObjectType = "URL_MAPPING"
	URLMappingsURLMappingCosObjectTypeVideoPlayer                  URLMappingsURLMappingCosObjectType = "VIDEO_PLAYER"
	URLMappingsURLMappingCosObjectTypeWidget                       URLMappingsURLMappingCosObjectType = "WIDGET"
	URLMappingsURLMappingCosObjectTypeWorkflow                     URLMappingsURLMappingCosObjectType = "WORKFLOW"
)

type URLMappingNewParams struct {
	URLMappingsURLMapping URLMappingsURLMappingParam
	paramObj
}

func (r URLMappingNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.URLMappingsURLMapping)
}
func (r *URLMappingNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
