// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cms

import (
	"context"
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
	"github.com/stainless-sdks/hubspot-sdk-go/option"
	"github.com/stainless-sdks/hubspot-sdk-go/packages/param"
)

// BlogPostService contains methods and other services that help with interacting
// with the hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBlogPostService] method instead.
type BlogPostService struct {
	Options []option.RequestOption
	Batch   BlogPostBatchService
}

// NewBlogPostService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewBlogPostService(opts ...option.RequestOption) (r BlogPostService) {
	r = BlogPostService{}
	r.Options = opts
	r.Batch = NewBlogPostBatchService(opts...)
	return
}

func (r *BlogPostService) New(ctx context.Context, body BlogPostNewParams, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "cms/blogs/2026-03/posts"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

func (r *BlogPostService) Update(ctx context.Context, objectID string, params BlogPostUpdateParams, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if objectID == "" {
		err = errors.New("missing required objectId parameter")
		return nil, err
	}
	path := fmt.Sprintf("cms/blogs/2026-03/posts/%s", objectID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return res, err
}

func (r *BlogPostService) List(ctx context.Context, query BlogPostListParams, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "cms/blogs/2026-03/posts"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

func (r *BlogPostService) Delete(ctx context.Context, objectID string, body BlogPostDeleteParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if objectID == "" {
		err = errors.New("missing required objectId parameter")
		return err
	}
	path := fmt.Sprintf("cms/blogs/2026-03/posts/%s", objectID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, nil, opts...)
	return err
}

func (r *BlogPostService) AttachToLangGroup(ctx context.Context, body BlogPostAttachToLangGroupParams, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "cms/blogs/2026-03/posts/multi-language/attach-to-lang-group"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

func (r *BlogPostService) Clone(ctx context.Context, body BlogPostCloneParams, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "cms/blogs/2026-03/posts/clone"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

func (r *BlogPostService) NewLangVariation(ctx context.Context, body BlogPostNewLangVariationParams, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "cms/blogs/2026-03/posts/multi-language/create-language-variation"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

func (r *BlogPostService) DetachFromLangGroup(ctx context.Context, body BlogPostDetachFromLangGroupParams, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "cms/blogs/2026-03/posts/multi-language/detach-from-lang-group"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

func (r *BlogPostService) Get(ctx context.Context, objectID string, query BlogPostGetParams, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if objectID == "" {
		err = errors.New("missing required objectId parameter")
		return nil, err
	}
	path := fmt.Sprintf("cms/blogs/2026-03/posts/%s", objectID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

func (r *BlogPostService) GetDraftByID(ctx context.Context, objectID string, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if objectID == "" {
		err = errors.New("missing required objectId parameter")
		return nil, err
	}
	path := fmt.Sprintf("cms/blogs/2026-03/posts/%s/draft", objectID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

func (r *BlogPostService) GetPreviousVersion(ctx context.Context, revisionID string, query BlogPostGetPreviousVersionParams, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if query.ObjectID == "" {
		err = errors.New("missing required objectId parameter")
		return nil, err
	}
	if revisionID == "" {
		err = errors.New("missing required revisionId parameter")
		return nil, err
	}
	path := fmt.Sprintf("cms/blogs/2026-03/posts/%s/revisions/%s", query.ObjectID, revisionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

func (r *BlogPostService) GetPreviousVersions(ctx context.Context, objectID string, query BlogPostGetPreviousVersionsParams, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if objectID == "" {
		err = errors.New("missing required objectId parameter")
		return nil, err
	}
	path := fmt.Sprintf("cms/blogs/2026-03/posts/%s/revisions", objectID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

func (r *BlogPostService) PushLive(ctx context.Context, objectID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if objectID == "" {
		err = errors.New("missing required objectId parameter")
		return err
	}
	path := fmt.Sprintf("cms/blogs/2026-03/posts/%s/draft/push-live", objectID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, nil, opts...)
	return err
}

func (r *BlogPostService) ResetDraft(ctx context.Context, objectID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if objectID == "" {
		err = errors.New("missing required objectId parameter")
		return err
	}
	path := fmt.Sprintf("cms/blogs/2026-03/posts/%s/draft/reset", objectID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, nil, opts...)
	return err
}

func (r *BlogPostService) RestorePreviousVersion(ctx context.Context, revisionID string, body BlogPostRestorePreviousVersionParams, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if body.ObjectID == "" {
		err = errors.New("missing required objectId parameter")
		return nil, err
	}
	if revisionID == "" {
		err = errors.New("missing required revisionId parameter")
		return nil, err
	}
	path := fmt.Sprintf("cms/blogs/2026-03/posts/%s/revisions/%s/restore", body.ObjectID, revisionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

func (r *BlogPostService) RestorePreviousVersionToDraft(ctx context.Context, revisionID int64, body BlogPostRestorePreviousVersionToDraftParams, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if body.ObjectID == "" {
		err = errors.New("missing required objectId parameter")
		return nil, err
	}
	path := fmt.Sprintf("cms/blogs/2026-03/posts/%s/revisions/%v/restore-to-draft", body.ObjectID, revisionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

func (r *BlogPostService) Schedule(ctx context.Context, body BlogPostScheduleParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "cms/blogs/2026-03/posts/schedule"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return err
}

func (r *BlogPostService) SetLangPrimary(ctx context.Context, body BlogPostSetLangPrimaryParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "cms/blogs/2026-03/posts/multi-language/set-new-lang-primary"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, nil, opts...)
	return err
}

func (r *BlogPostService) UpdateDraft(ctx context.Context, objectID string, body BlogPostUpdateDraftParams, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if objectID == "" {
		err = errors.New("missing required objectId parameter")
		return nil, err
	}
	path := fmt.Sprintf("cms/blogs/2026-03/posts/%s/draft", objectID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

func (r *BlogPostService) UpdateLangs(ctx context.Context, body BlogPostUpdateLangsParams, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "cms/blogs/2026-03/posts/multi-language/update-languages"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// The properties Units, Value are required.
type AngleParam struct {
	// The unit of measurement for the angle.
	//
	// Any of "deg", "grad", "rad", "turn".
	Units AngleUnits `json:"units,omitzero" api:"required"`
	// The numerical representation of the angle.
	Value float64 `json:"value" api:"required"`
	paramObj
}

func (r AngleParam) MarshalJSON() (data []byte, err error) {
	type shadow AngleParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AngleParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The unit of measurement for the angle.
type AngleUnits string

const (
	AngleUnitsDeg  AngleUnits = "deg"
	AngleUnitsGrad AngleUnits = "grad"
	AngleUnitsRad  AngleUnits = "rad"
	AngleUnitsTurn AngleUnits = "turn"
)

// The properties BackgroundPosition, BackgroundSize, ImageURL are required.
type BackgroundImageParam struct {
	// Defines the position of the background image.
	BackgroundPosition string `json:"backgroundPosition" api:"required"`
	// Specifies the size of the background image.
	BackgroundSize string `json:"backgroundSize" api:"required"`
	// The URL of the background image.
	ImageURL string `json:"imageUrl" api:"required"`
	paramObj
}

func (r BackgroundImageParam) MarshalJSON() (data []byte, err error) {
	type shadow BackgroundImageParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BackgroundImageParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Inputs is required.
type BatchInputBlogPostParam struct {
	// Blog posts to input.
	Inputs []BlogPostParam `json:"inputs,omitzero" api:"required"`
	paramObj
}

func (r BatchInputBlogPostParam) MarshalJSON() (data []byte, err error) {
	type shadow BatchInputBlogPostParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BatchInputBlogPostParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Inputs is required.
type BatchInputJsonNodeParam struct {
	// JSON nodes to input.
	Inputs []any `json:"inputs,omitzero" api:"required"`
	paramObj
}

func (r BatchInputJsonNodeParam) MarshalJSON() (data []byte, err error) {
	type shadow BatchInputJsonNodeParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BatchInputJsonNodeParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Inputs is required.
type BatchInputStringParam struct {
	// Strings to input.
	Inputs []string `json:"inputs,omitzero" api:"required"`
	paramObj
}

func (r BatchInputStringParam) MarshalJSON() (data []byte, err error) {
	type shadow BatchInputStringParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BatchInputStringParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties ID, AbStatus, AbTestID, ArchivedAt, ArchivedInDashboard,
// AttachedStylesheets, AuthorName, BlogAuthorID, Campaign, CategoryID,
// ContentGroupID, ContentTypeCategory, Created, CreatedByID, CurrentlyPublished,
// CurrentState, Domain, DynamicPageDataSourceID, DynamicPageDataSourceType,
// DynamicPageHubDBTableID, EnableDomainStylesheets, EnableGoogleAmpOutputOverride,
// EnableLayoutStylesheets, FeaturedImage, FeaturedImageAltText, FolderID,
// FooterHTML, HeadHTML, HTMLTitle, IncludeDefaultCustomCss, Language,
// LayoutSections, LinkRelCanonicalURL, MabExperimentID, MetaDescription, Name,
// PageExpiryDate, PageExpiryEnabled, PageExpiryRedirectID, PageExpiryRedirectURL,
// Password, PostBody, PostSummary, PublicAccessRules, PublicAccessRulesEnabled,
// PublishDate, PublishImmediately, RssBody, RssSummary, Slug, State, TagIDs,
// ThemeSettingsValues, TranslatedFromID, Translations, Updated, UpdatedByID, URL,
// UseFeaturedImage, WidgetContainers, Widgets are required.
type BlogPostParam struct {
	// The unique ID of the blog post.
	ID string `json:"id" api:"required"`
	// The status of the AB test associated with this blog post, if applicable
	//
	// Available options: automated_loser_variant, automated_master, automated_variant,
	// loser_variant, mab_master, mab_variant, master, variant
	//
	// Any of "automated_loser_variant", "automated_master", "automated_variant",
	// "loser_variant", "mab_master", "mab_variant", "master", "variant".
	AbStatus BlogPostAbStatus `json:"abStatus,omitzero" api:"required"`
	// The ID of the AB test associated with this page, if applicable
	AbTestID string `json:"abTestId" api:"required"`
	// The timestamp (ISO8601 format) when this Blog Post was deleted.
	ArchivedAt int64 `json:"archivedAt" api:"required"`
	// If True, the post will not show up in your dashboard, although the post could
	// still be live.
	ArchivedInDashboard bool `json:"archivedInDashboard" api:"required"`
	// List of stylesheets to attach to this blog post. These stylesheets are attached
	// to just this page. Order of precedence is bottom to top, just like in the HTML.
	AttachedStylesheets []map[string]any `json:"attachedStylesheets,omitzero" api:"required"`
	// The name of the user who last published the blog post. For posts that haven't
	// been published yet, this property will reflect the user who initially created
	// the draft.
	AuthorName string `json:"authorName" api:"required"`
	// The ID of the blog author associated with this post.
	BlogAuthorID string `json:"blogAuthorId" api:"required"`
	// The GUID of the marketing campaign the post is associated with.
	Campaign string `json:"campaign" api:"required"`
	// ID of the object type.
	CategoryID int64 `json:"categoryId" api:"required"`
	// The ID of the post's parent blog.
	ContentGroupID string `json:"contentGroupId" api:"required"`
	// An ENUM descibing the type of this object. Should always be BLOG_POST.
	//
	// Any of "0", "1", "10", "11", "12", "13", "14", "15", "16", "17", "18", "19",
	// "2", "20", "21", "22", "3", "4", "5", "6", "7", "8", "9".
	ContentTypeCategory BlogPostContentTypeCategory `json:"contentTypeCategory,omitzero" api:"required"`
	// The timestamp (ISO8601 format) when this Blog Post was created.
	Created time.Time `json:"created" api:"required" format:"date-time"`
	// The ID of the user that created the post.
	CreatedByID string `json:"createdById" api:"required"`
	// Whether the post is published (true or false)
	CurrentlyPublished bool `json:"currentlyPublished" api:"required"`
	// A generated ENUM descibing the current state of this Blog Post. Should always
	// match state.
	//
	// Any of "AGENT_GENERATED", "AUTOMATED", "AUTOMATED_AB", "AUTOMATED_AB_VARIANT",
	// "AUTOMATED_DRAFT", "AUTOMATED_DRAFT_AB", "AUTOMATED_DRAFT_ABVARIANT",
	// "AUTOMATED_FOR_FORM", "AUTOMATED_FOR_FORM_BUFFER", "AUTOMATED_FOR_FORM_DRAFT",
	// "AUTOMATED_FOR_FORM_LEGACY", "AUTOMATED_LOSER_ABVARIANT", "AUTOMATED_SENDING",
	// "BLOG_EMAIL_DRAFT", "BLOG_EMAIL_PUBLISHED", "DRAFT", "DRAFT_AB",
	// "DRAFT_AB_VARIANT", "ERROR", "LOSER_AB_VARIANT", "PAGE_STUB", "PRE_PROCESSING",
	// "PROCESSING", "PUBLISHED", "PUBLISHED_AB", "PUBLISHED_AB_VARIANT",
	// "PUBLISHED_OR_SCHEDULED", "RSS_TO_EMAIL_DRAFT", "RSS_TO_EMAIL_PUBLISHED",
	// "SCHEDULED", "SCHEDULED_AB", "SCHEDULED_OR_PUBLISHED".
	CurrentState BlogPostCurrentState `json:"currentState,omitzero" api:"required"`
	// The domain that the post lives on. If null, the post will default to the domain
	// of the parent blog.
	Domain string `json:"domain" api:"required"`
	// The identifier for the data source used by the dynamic page.
	DynamicPageDataSourceID string `json:"dynamicPageDataSourceId" api:"required"`
	// The type of data source used by the dynamic page.
	DynamicPageDataSourceType int64 `json:"dynamicPageDataSourceType" api:"required"`
	// For dynamic HubDB pages, the ID of the HubDB table this post references.
	DynamicPageHubDBTableID string `json:"dynamicPageHubDbTableId" api:"required"`
	// Boolean to determine whether or not the styles from the template should be
	// applied.
	EnableDomainStylesheets bool `json:"enableDomainStylesheets" api:"required"`
	// Boolean to allow overriding the AMP settings for the blog.
	EnableGoogleAmpOutputOverride bool `json:"enableGoogleAmpOutputOverride" api:"required"`
	// Boolean to determine whether or not the styles from the template should be
	// applied.
	EnableLayoutStylesheets bool `json:"enableLayoutStylesheets" api:"required"`
	// The featuredImage of this Blog Post.
	FeaturedImage string `json:"featuredImage" api:"required"`
	// Alt Text of the featuredImage.
	FeaturedImageAltText string `json:"featuredImageAltText" api:"required"`
	// Unique identifier of associated folder
	FolderID string `json:"folderId" api:"required"`
	// Custom HTML for embed codes, javascript that should be placed before the </body>
	// tag of the page.
	FooterHTML string `json:"footerHtml" api:"required"`
	// Custom HTML for embed codes, javascript, etc. that goes in the <head> tag of the
	// page.
	HeadHTML string `json:"headHtml" api:"required"`
	// The HTML title of the post.
	HTMLTitle string `json:"htmlTitle" api:"required"`
	// Boolean to determine whether or not the Primary CSS Files should be applied.
	IncludeDefaultCustomCss bool `json:"includeDefaultCustomCss" api:"required"`
	// The explicitly defined ISO 639 language code of the post. If null, the post will
	// default to the language of the parent blog.
	//
	// Any of "aa", "ab", "ae", "af", "af-na", "af-za", "agq", "agq-cm", "ak", "ak-gh",
	// "am", "am-et", "an", "ann", "ann-ng", "ar", "ar-001", "ar-ae", "ar-bh", "ar-dj",
	// "ar-dz", "ar-eg", "ar-eh", "ar-er", "ar-il", "ar-iq", "ar-jo", "ar-km", "ar-kw",
	// "ar-lb", "ar-ly", "ar-ma", "ar-mr", "ar-om", "ar-ps", "ar-qa", "ar-sa", "ar-sd",
	// "ar-so", "ar-ss", "ar-sy", "ar-td", "ar-tn", "ar-ye", "as", "as-in", "asa",
	// "asa-tz", "ast", "ast-es", "av", "ay", "az", "az-az", "ba", "bas", "bas-cm",
	// "be", "be-by", "bem", "bem-zm", "bez", "bez-tz", "bg", "bg-bg", "bgc", "bgc-in",
	// "bho", "bho-in", "bi", "bm", "bm-ml", "bn", "bn-bd", "bn-in", "bo", "bo-cn",
	// "bo-in", "br", "br-fr", "brx", "brx-in", "bs", "bs-ba", "ca", "ca-ad", "ca-es",
	// "ca-fr", "ca-it", "ccp", "ccp-bd", "ccp-in", "ce", "ce-ru", "ceb", "ceb-ph",
	// "cgg", "cgg-ug", "ch", "chr", "chr-us", "ckb", "ckb-iq", "ckb-ir", "co", "cr",
	// "cs", "cs-cz", "cu", "cu-ru", "cv", "cv-ru", "cy", "cy-gb", "da", "da-dk",
	// "da-gl", "dav", "dav-ke", "de", "de-at", "de-be", "de-ch", "de-de", "de-gr",
	// "de-it", "de-li", "de-lu", "dje", "dje-ne", "doi", "doi-in", "dsb", "dsb-de",
	// "dua", "dua-cm", "dv", "dyo", "dyo-sn", "dz", "dz-bt", "ebu", "ebu-ke", "ee",
	// "ee-gh", "ee-tg", "el", "el-cy", "el-gr", "en", "en-001", "en-150", "en-ae",
	// "en-ag", "en-ai", "en-as", "en-at", "en-au", "en-bb", "en-be", "en-bi", "en-bm",
	// "en-bs", "en-bw", "en-bz", "en-ca", "en-cc", "en-ch", "en-ck", "en-cm", "en-cn",
	// "en-cx", "en-cy", "en-de", "en-dg", "en-dk", "en-dm", "en-ee", "en-eg", "en-er",
	// "en-es", "en-fi", "en-fj", "en-fk", "en-fm", "en-fr", "en-gb", "en-gd", "en-gg",
	// "en-gh", "en-gi", "en-gm", "en-gu", "en-gy", "en-hk", "en-id", "en-ie", "en-il",
	// "en-im", "en-in", "en-io", "en-je", "en-jm", "en-ke", "en-ki", "en-kn", "en-ky",
	// "en-lc", "en-lr", "en-ls", "en-lu", "en-mg", "en-mh", "en-mo", "en-mp", "en-ms",
	// "en-mt", "en-mu", "en-mv", "en-mw", "en-mx", "en-my", "en-na", "en-nf", "en-ng",
	// "en-nl", "en-nr", "en-nu", "en-nz", "en-pg", "en-ph", "en-pk", "en-pn", "en-pr",
	// "en-pt", "en-pw", "en-rw", "en-sb", "en-sc", "en-sd", "en-se", "en-sg", "en-sh",
	// "en-si", "en-sl", "en-ss", "en-sx", "en-sz", "en-tc", "en-th", "en-tk", "en-tn",
	// "en-to", "en-tt", "en-tv", "en-tz", "en-ug", "en-um", "en-us", "en-vc", "en-vg",
	// "en-vi", "en-vn", "en-vu", "en-ws", "en-za", "en-zm", "en-zw", "eo", "eo-001",
	// "es", "es-419", "es-ar", "es-bo", "es-br", "es-bz", "es-cl", "es-co", "es-cr",
	// "es-cu", "es-do", "es-ea", "es-ec", "es-es", "es-gq", "es-gt", "es-hn", "es-ic",
	// "es-mx", "es-ni", "es-pa", "es-pe", "es-ph", "es-pr", "es-py", "es-sv", "es-us",
	// "es-uy", "es-ve", "et", "et-ee", "eu", "eu-es", "ewo", "ewo-cm", "fa", "fa-af",
	// "fa-ir", "ff", "ff-bf", "ff-cm", "ff-gh", "ff-gm", "ff-gn", "ff-gw", "ff-lr",
	// "ff-mr", "ff-ne", "ff-ng", "ff-sl", "ff-sn", "fi", "fi-fi", "fil", "fil-ph",
	// "fj", "fo", "fo-dk", "fo-fo", "fr", "fr-be", "fr-bf", "fr-bi", "fr-bj", "fr-bl",
	// "fr-ca", "fr-cd", "fr-cf", "fr-cg", "fr-ch", "fr-ci", "fr-cm", "fr-dj", "fr-dz",
	// "fr-fr", "fr-ga", "fr-gf", "fr-gn", "fr-gp", "fr-gq", "fr-ht", "fr-km", "fr-lu",
	// "fr-ma", "fr-mc", "fr-mf", "fr-mg", "fr-ml", "fr-mq", "fr-mr", "fr-mu", "fr-nc",
	// "fr-ne", "fr-pf", "fr-pm", "fr-re", "fr-rw", "fr-sc", "fr-sn", "fr-sy", "fr-td",
	// "fr-tg", "fr-tn", "fr-vu", "fr-wf", "fr-yt", "frr", "frr-de", "fur", "fur-it",
	// "fy", "fy-nl", "ga", "ga-gb", "ga-ie", "gd", "gd-gb", "gl", "gl-es", "gn",
	// "gsw", "gsw-ch", "gsw-fr", "gsw-li", "gu", "gu-in", "guz", "guz-ke", "gv",
	// "gv-im", "ha", "ha-gh", "ha-ne", "ha-ng", "haw", "haw-us", "he", "he-il", "hi",
	// "hi-in", "hmn", "ho", "hr", "hr-ba", "hr-hr", "hsb", "hsb-de", "ht", "hu",
	// "hu-hu", "hy", "hy-am", "hz", "ia", "ia-001", "id", "id-id", "ie", "ig",
	// "ig-ng", "ii", "ii-cn", "ik", "io", "is", "is-is", "it", "it-ch", "it-it",
	// "it-sm", "it-va", "iu", "ja", "ja-jp", "jgo", "jgo-cm", "jmc", "jmc-tz", "jv",
	// "jv-id", "ka", "ka-ge", "kab", "kab-dz", "kam", "kam-ke", "kar", "kde",
	// "kde-tz", "kea", "kea-cv", "kg", "kgp", "kgp-br", "kh", "khq", "khq-ml", "ki",
	// "ki-ke", "kj", "kk", "kk-kz", "kkj", "kkj-cm", "kl", "kl-gl", "kln", "kln-ke",
	// "km", "km-kh", "kn", "kn-in", "ko", "ko-kp", "ko-kr", "kok", "kok-in", "kr",
	// "ks", "ks-in", "ksb", "ksb-tz", "ksf", "ksf-cm", "ksh", "ksh-de", "ku", "ku-tr",
	// "kv", "kw", "kw-gb", "ky", "ky-kg", "la", "lag", "lag-tz", "lb", "lb-lu", "lg",
	// "lg-ug", "li", "lkt", "lkt-us", "ln", "ln-ao", "ln-cd", "ln-cf", "ln-cg", "lo",
	// "lo-la", "lrc", "lrc-iq", "lrc-ir", "lt", "lt-lt", "lu", "lu-cd", "luo",
	// "luo-ke", "luy", "luy-ke", "lv", "lv-lv", "mai", "mai-in", "mas", "mas-ke",
	// "mas-tz", "mdf", "mdf-ru", "mer", "mer-ke", "mfe", "mfe-mu", "mg", "mg-mg",
	// "mgh", "mgh-mz", "mgo", "mgo-cm", "mh", "mi", "mi-nz", "mk", "mk-mk", "ml",
	// "ml-in", "mn", "mn-mn", "mni", "mni-in", "mr", "mr-in", "ms", "ms-bn", "ms-id",
	// "ms-my", "ms-sg", "mt", "mt-mt", "mua", "mua-cm", "my", "my-mm", "mzn",
	// "mzn-ir", "na", "naq", "naq-na", "nb", "nb-no", "nb-sj", "nd", "nd-zw", "nds",
	// "nds-de", "nds-nl", "ne", "ne-in", "ne-np", "ng", "nl", "nl-aw", "nl-be",
	// "nl-bq", "nl-ch", "nl-cw", "nl-lu", "nl-nl", "nl-sr", "nl-sx", "nmg", "nmg-cm",
	// "nn", "nn-no", "nnh", "nnh-cm", "no", "no-no", "nr", "nus", "nus-ss", "nv",
	// "ny", "nyn", "nyn-ug", "oc", "oc-es", "oc-fr", "oj", "om", "om-et", "om-ke",
	// "or", "or-in", "os", "os-ge", "os-ru", "pa", "pa-in", "pa-pk", "pcm", "pcm-ng",
	// "pi", "pis", "pis-sb", "pl", "pl-pl", "prg", "prg-001", "ps", "ps-af", "ps-pk",
	// "pt", "pt-ao", "pt-br", "pt-ch", "pt-cv", "pt-gq", "pt-gw", "pt-lu", "pt-mo",
	// "pt-mz", "pt-pt", "pt-st", "pt-tl", "qu", "qu-bo", "qu-ec", "qu-pe", "raj",
	// "raj-in", "rm", "rm-ch", "rn", "rn-bi", "ro", "ro-md", "ro-ro", "rof", "rof-tz",
	// "ru", "ru-by", "ru-kg", "ru-kz", "ru-md", "ru-ru", "ru-ua", "rw", "rw-rw",
	// "rwk", "rwk-tz", "sa", "sa-in", "sah", "sah-ru", "saq", "saq-ke", "sat",
	// "sat-in", "sbp", "sbp-tz", "sc", "sc-it", "sd", "sd-in", "sd-pk", "se", "se-fi",
	// "se-no", "se-se", "seh", "seh-mz", "ses", "ses-ml", "sg", "sg-cf", "shi",
	// "shi-ma", "si", "si-lk", "sk", "sk-sk", "sl", "sl-si", "sm", "smn", "smn-fi",
	// "sms", "sms-fi", "sn", "sn-zw", "so", "so-dj", "so-et", "so-ke", "so-so", "sq",
	// "sq-al", "sq-mk", "sq-xk", "sr", "sr-ba", "sr-cs", "sr-me", "sr-rs", "sr-xk",
	// "ss", "st", "su", "su-id", "sv", "sv-ax", "sv-fi", "sv-se", "sw", "sw-cd",
	// "sw-ke", "sw-tz", "sw-ug", "sy", "ta", "ta-in", "ta-lk", "ta-my", "ta-sg", "te",
	// "te-in", "teo", "teo-ke", "teo-ug", "tg", "tg-tj", "th", "th-th", "ti", "ti-er",
	// "ti-et", "tk", "tk-tm", "tl", "tn", "to", "to-to", "tok", "tok-001", "tr",
	// "tr-cy", "tr-tr", "ts", "tt", "tt-ru", "tw", "twq", "twq-ne", "ty", "tzm",
	// "tzm-ma", "ug", "ug-cn", "uk", "uk-ua", "ur", "ur-in", "ur-pk", "uz", "uz-af",
	// "uz-uz", "vai", "vai-lr", "ve", "vi", "vi-vn", "vo", "vo-001", "vun", "vun-tz",
	// "wa", "wae", "wae-ch", "wo", "wo-sn", "xh", "xh-za", "xog", "xog-ug", "yav",
	// "yav-cm", "yi", "yi-001", "yo", "yo-bj", "yo-ng", "yrl", "yrl-br", "yrl-co",
	// "yrl-ve", "yue", "yue-cn", "yue-hk", "za", "zgh", "zgh-ma", "zh", "zh-cn",
	// "zh-hans", "zh-hant", "zh-hk", "zh-mo", "zh-sg", "zh-tw", "zu", "zu-za".
	Language BlogPostLanguage `json:"language,omitzero" api:"required"`
	// A structure detailing the layout sections of the blog post.
	LayoutSections map[string]LayoutSectionParam `json:"layoutSections,omitzero" api:"required"`
	// Optional override to set the URL to be used in the rel=canonical link tag on the
	// page.
	LinkRelCanonicalURL string `json:"linkRelCanonicalUrl" api:"required"`
	// Unique identifier of the MAB Experiment
	MabExperimentID string `json:"mabExperimentId" api:"required"`
	// A description that goes in <meta> tag on the page.
	MetaDescription string `json:"metaDescription" api:"required"`
	// The internal name of the post.
	Name string `json:"name" api:"required"`
	// The date at which this blog post should expire and begin redirecting to another
	// url or page.
	PageExpiryDate int64 `json:"pageExpiryDate" api:"required"`
	// Boolean describing if the page expiration feature is enabled for this blog post.
	PageExpiryEnabled bool `json:"pageExpiryEnabled" api:"required"`
	// The ID of another page this blog post's url should redirect to once this blog
	// post expires. Should only set this or pageExpiryRedirectUrl.
	PageExpiryRedirectID int64 `json:"pageExpiryRedirectId" api:"required"`
	// The URL this blog post's url should redirect to once it expires. Should only set
	// this or pageExpiryRedirectId.
	PageExpiryRedirectURL string `json:"pageExpiryRedirectUrl" api:"required"`
	// Set this to create a password protected page. Entering the password will be
	// required to view the blog post.
	Password string `json:"password" api:"required"`
	// The HTML of the main post body.
	PostBody string `json:"postBody" api:"required"`
	// The summary of the blog post that will appear on the main listing page.
	PostSummary string `json:"postSummary" api:"required"`
	// Rules for require member registration to access private content.
	PublicAccessRules []PublicAccessRule `json:"publicAccessRules,omitzero" api:"required"`
	// Boolean to determine whether or not to respect publicAccessRules.
	PublicAccessRulesEnabled bool `json:"publicAccessRulesEnabled" api:"required"`
	// The date (ISO8601 format) the blog post is to be published at.
	PublishDate time.Time `json:"publishDate" api:"required" format:"date-time"`
	// Set this to true if you want to be published immediately when the schedule
	// publish endpoint is called, and to ignore the publish_date setting.
	PublishImmediately bool `json:"publishImmediately" api:"required"`
	// The contents of the RSS body for this Blog Post.
	RssBody string `json:"rssBody" api:"required"`
	// The contents of the RSS summary for this Blog Post.
	RssSummary string `json:"rssSummary" api:"required"`
	// The URL slug of the blog post. This field is appended to the domain to construct
	// the url of this post.
	Slug string `json:"slug" api:"required"`
	// An enumeration describing the current publish state of the post.
	State string `json:"state" api:"required"`
	// The IDs of the tags associated with this post.
	TagIDs []int64 `json:"tagIds,omitzero" api:"required"`
	// A collection of settings specific to the theme applied to the blog post.
	ThemeSettingsValues map[string]any `json:"themeSettingsValues,omitzero" api:"required"`
	// ID of the primary blog post that this post was translated from.
	TranslatedFromID string `json:"translatedFromId" api:"required"`
	// A map of translations for the blog post, each associated with a specific
	// language variation.
	Translations map[string]ContentLanguageVariationParam `json:"translations,omitzero" api:"required"`
	// The timestamp (ISO8601 format) when this Blog Post was updated.
	Updated time.Time `json:"updated" api:"required" format:"date-time"`
	// The ID of the user that updated the post.
	UpdatedByID string `json:"updatedById" api:"required"`
	// A generated field representing the URL of this blog post.
	URL string `json:"url" api:"required"`
	// Boolean to determine if this post should use a featured image.
	UseFeaturedImage bool `json:"useFeaturedImage" api:"required"`
	// A data structure containing the data for all the modules inside the containers
	// for this post. This will only be populated if the page has widget containers.
	WidgetContainers map[string]any `json:"widgetContainers,omitzero" api:"required"`
	// A data structure containing the data for all the modules for this page.
	Widgets map[string]any `json:"widgets,omitzero" api:"required"`
	paramObj
}

func (r BlogPostParam) MarshalJSON() (data []byte, err error) {
	type shadow BlogPostParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BlogPostParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the AB test associated with this blog post, if applicable
//
// Available options: automated_loser_variant, automated_master, automated_variant,
// loser_variant, mab_master, mab_variant, master, variant
type BlogPostAbStatus string

const (
	BlogPostAbStatusAutomatedLoserVariant BlogPostAbStatus = "automated_loser_variant"
	BlogPostAbStatusAutomatedMaster       BlogPostAbStatus = "automated_master"
	BlogPostAbStatusAutomatedVariant      BlogPostAbStatus = "automated_variant"
	BlogPostAbStatusLoserVariant          BlogPostAbStatus = "loser_variant"
	BlogPostAbStatusMabMaster             BlogPostAbStatus = "mab_master"
	BlogPostAbStatusMabVariant            BlogPostAbStatus = "mab_variant"
	BlogPostAbStatusMaster                BlogPostAbStatus = "master"
	BlogPostAbStatusVariant               BlogPostAbStatus = "variant"
)

// An ENUM descibing the type of this object. Should always be BLOG_POST.
type BlogPostContentTypeCategory string

const (
	BlogPostContentTypeCategory0  BlogPostContentTypeCategory = "0"
	BlogPostContentTypeCategory1  BlogPostContentTypeCategory = "1"
	BlogPostContentTypeCategory10 BlogPostContentTypeCategory = "10"
	BlogPostContentTypeCategory11 BlogPostContentTypeCategory = "11"
	BlogPostContentTypeCategory12 BlogPostContentTypeCategory = "12"
	BlogPostContentTypeCategory13 BlogPostContentTypeCategory = "13"
	BlogPostContentTypeCategory14 BlogPostContentTypeCategory = "14"
	BlogPostContentTypeCategory15 BlogPostContentTypeCategory = "15"
	BlogPostContentTypeCategory16 BlogPostContentTypeCategory = "16"
	BlogPostContentTypeCategory17 BlogPostContentTypeCategory = "17"
	BlogPostContentTypeCategory18 BlogPostContentTypeCategory = "18"
	BlogPostContentTypeCategory19 BlogPostContentTypeCategory = "19"
	BlogPostContentTypeCategory2  BlogPostContentTypeCategory = "2"
	BlogPostContentTypeCategory20 BlogPostContentTypeCategory = "20"
	BlogPostContentTypeCategory21 BlogPostContentTypeCategory = "21"
	BlogPostContentTypeCategory22 BlogPostContentTypeCategory = "22"
	BlogPostContentTypeCategory3  BlogPostContentTypeCategory = "3"
	BlogPostContentTypeCategory4  BlogPostContentTypeCategory = "4"
	BlogPostContentTypeCategory5  BlogPostContentTypeCategory = "5"
	BlogPostContentTypeCategory6  BlogPostContentTypeCategory = "6"
	BlogPostContentTypeCategory7  BlogPostContentTypeCategory = "7"
	BlogPostContentTypeCategory8  BlogPostContentTypeCategory = "8"
	BlogPostContentTypeCategory9  BlogPostContentTypeCategory = "9"
)

// A generated ENUM descibing the current state of this Blog Post. Should always
// match state.
type BlogPostCurrentState string

const (
	BlogPostCurrentStateAgentGenerated          BlogPostCurrentState = "AGENT_GENERATED"
	BlogPostCurrentStateAutomated               BlogPostCurrentState = "AUTOMATED"
	BlogPostCurrentStateAutomatedAb             BlogPostCurrentState = "AUTOMATED_AB"
	BlogPostCurrentStateAutomatedAbVariant      BlogPostCurrentState = "AUTOMATED_AB_VARIANT"
	BlogPostCurrentStateAutomatedDraft          BlogPostCurrentState = "AUTOMATED_DRAFT"
	BlogPostCurrentStateAutomatedDraftAb        BlogPostCurrentState = "AUTOMATED_DRAFT_AB"
	BlogPostCurrentStateAutomatedDraftAbvariant BlogPostCurrentState = "AUTOMATED_DRAFT_ABVARIANT"
	BlogPostCurrentStateAutomatedForForm        BlogPostCurrentState = "AUTOMATED_FOR_FORM"
	BlogPostCurrentStateAutomatedForFormBuffer  BlogPostCurrentState = "AUTOMATED_FOR_FORM_BUFFER"
	BlogPostCurrentStateAutomatedForFormDraft   BlogPostCurrentState = "AUTOMATED_FOR_FORM_DRAFT"
	BlogPostCurrentStateAutomatedForFormLegacy  BlogPostCurrentState = "AUTOMATED_FOR_FORM_LEGACY"
	BlogPostCurrentStateAutomatedLoserAbvariant BlogPostCurrentState = "AUTOMATED_LOSER_ABVARIANT"
	BlogPostCurrentStateAutomatedSending        BlogPostCurrentState = "AUTOMATED_SENDING"
	BlogPostCurrentStateBlogEmailDraft          BlogPostCurrentState = "BLOG_EMAIL_DRAFT"
	BlogPostCurrentStateBlogEmailPublished      BlogPostCurrentState = "BLOG_EMAIL_PUBLISHED"
	BlogPostCurrentStateDraft                   BlogPostCurrentState = "DRAFT"
	BlogPostCurrentStateDraftAb                 BlogPostCurrentState = "DRAFT_AB"
	BlogPostCurrentStateDraftAbVariant          BlogPostCurrentState = "DRAFT_AB_VARIANT"
	BlogPostCurrentStateError                   BlogPostCurrentState = "ERROR"
	BlogPostCurrentStateLoserAbVariant          BlogPostCurrentState = "LOSER_AB_VARIANT"
	BlogPostCurrentStatePageStub                BlogPostCurrentState = "PAGE_STUB"
	BlogPostCurrentStatePreProcessing           BlogPostCurrentState = "PRE_PROCESSING"
	BlogPostCurrentStateProcessing              BlogPostCurrentState = "PROCESSING"
	BlogPostCurrentStatePublished               BlogPostCurrentState = "PUBLISHED"
	BlogPostCurrentStatePublishedAb             BlogPostCurrentState = "PUBLISHED_AB"
	BlogPostCurrentStatePublishedAbVariant      BlogPostCurrentState = "PUBLISHED_AB_VARIANT"
	BlogPostCurrentStatePublishedOrScheduled    BlogPostCurrentState = "PUBLISHED_OR_SCHEDULED"
	BlogPostCurrentStateRssToEmailDraft         BlogPostCurrentState = "RSS_TO_EMAIL_DRAFT"
	BlogPostCurrentStateRssToEmailPublished     BlogPostCurrentState = "RSS_TO_EMAIL_PUBLISHED"
	BlogPostCurrentStateScheduled               BlogPostCurrentState = "SCHEDULED"
	BlogPostCurrentStateScheduledAb             BlogPostCurrentState = "SCHEDULED_AB"
	BlogPostCurrentStateScheduledOrPublished    BlogPostCurrentState = "SCHEDULED_OR_PUBLISHED"
)

// The explicitly defined ISO 639 language code of the post. If null, the post will
// default to the language of the parent blog.
type BlogPostLanguage string

const (
	BlogPostLanguageAa     BlogPostLanguage = "aa"
	BlogPostLanguageAb     BlogPostLanguage = "ab"
	BlogPostLanguageAe     BlogPostLanguage = "ae"
	BlogPostLanguageAf     BlogPostLanguage = "af"
	BlogPostLanguageAfNa   BlogPostLanguage = "af-na"
	BlogPostLanguageAfZa   BlogPostLanguage = "af-za"
	BlogPostLanguageAgq    BlogPostLanguage = "agq"
	BlogPostLanguageAgqCm  BlogPostLanguage = "agq-cm"
	BlogPostLanguageAk     BlogPostLanguage = "ak"
	BlogPostLanguageAkGh   BlogPostLanguage = "ak-gh"
	BlogPostLanguageAm     BlogPostLanguage = "am"
	BlogPostLanguageAmEt   BlogPostLanguage = "am-et"
	BlogPostLanguageAn     BlogPostLanguage = "an"
	BlogPostLanguageAnn    BlogPostLanguage = "ann"
	BlogPostLanguageAnnNg  BlogPostLanguage = "ann-ng"
	BlogPostLanguageAr     BlogPostLanguage = "ar"
	BlogPostLanguageAr001  BlogPostLanguage = "ar-001"
	BlogPostLanguageArAe   BlogPostLanguage = "ar-ae"
	BlogPostLanguageArBh   BlogPostLanguage = "ar-bh"
	BlogPostLanguageArDj   BlogPostLanguage = "ar-dj"
	BlogPostLanguageArDz   BlogPostLanguage = "ar-dz"
	BlogPostLanguageArEg   BlogPostLanguage = "ar-eg"
	BlogPostLanguageArEh   BlogPostLanguage = "ar-eh"
	BlogPostLanguageArEr   BlogPostLanguage = "ar-er"
	BlogPostLanguageArIl   BlogPostLanguage = "ar-il"
	BlogPostLanguageArIq   BlogPostLanguage = "ar-iq"
	BlogPostLanguageArJo   BlogPostLanguage = "ar-jo"
	BlogPostLanguageArKm   BlogPostLanguage = "ar-km"
	BlogPostLanguageArKw   BlogPostLanguage = "ar-kw"
	BlogPostLanguageArLb   BlogPostLanguage = "ar-lb"
	BlogPostLanguageArLy   BlogPostLanguage = "ar-ly"
	BlogPostLanguageArMa   BlogPostLanguage = "ar-ma"
	BlogPostLanguageArMr   BlogPostLanguage = "ar-mr"
	BlogPostLanguageArOm   BlogPostLanguage = "ar-om"
	BlogPostLanguageArPs   BlogPostLanguage = "ar-ps"
	BlogPostLanguageArQa   BlogPostLanguage = "ar-qa"
	BlogPostLanguageArSa   BlogPostLanguage = "ar-sa"
	BlogPostLanguageArSd   BlogPostLanguage = "ar-sd"
	BlogPostLanguageArSo   BlogPostLanguage = "ar-so"
	BlogPostLanguageArSS   BlogPostLanguage = "ar-ss"
	BlogPostLanguageArSy   BlogPostLanguage = "ar-sy"
	BlogPostLanguageArTd   BlogPostLanguage = "ar-td"
	BlogPostLanguageArTn   BlogPostLanguage = "ar-tn"
	BlogPostLanguageArYe   BlogPostLanguage = "ar-ye"
	BlogPostLanguageAs     BlogPostLanguage = "as"
	BlogPostLanguageAsIn   BlogPostLanguage = "as-in"
	BlogPostLanguageAsa    BlogPostLanguage = "asa"
	BlogPostLanguageAsaTz  BlogPostLanguage = "asa-tz"
	BlogPostLanguageAst    BlogPostLanguage = "ast"
	BlogPostLanguageAstEs  BlogPostLanguage = "ast-es"
	BlogPostLanguageAv     BlogPostLanguage = "av"
	BlogPostLanguageAy     BlogPostLanguage = "ay"
	BlogPostLanguageAz     BlogPostLanguage = "az"
	BlogPostLanguageAzAz   BlogPostLanguage = "az-az"
	BlogPostLanguageBa     BlogPostLanguage = "ba"
	BlogPostLanguageBas    BlogPostLanguage = "bas"
	BlogPostLanguageBasCm  BlogPostLanguage = "bas-cm"
	BlogPostLanguageBe     BlogPostLanguage = "be"
	BlogPostLanguageBeBy   BlogPostLanguage = "be-by"
	BlogPostLanguageBem    BlogPostLanguage = "bem"
	BlogPostLanguageBemZm  BlogPostLanguage = "bem-zm"
	BlogPostLanguageBez    BlogPostLanguage = "bez"
	BlogPostLanguageBezTz  BlogPostLanguage = "bez-tz"
	BlogPostLanguageBg     BlogPostLanguage = "bg"
	BlogPostLanguageBgBg   BlogPostLanguage = "bg-bg"
	BlogPostLanguageBgc    BlogPostLanguage = "bgc"
	BlogPostLanguageBgcIn  BlogPostLanguage = "bgc-in"
	BlogPostLanguageBho    BlogPostLanguage = "bho"
	BlogPostLanguageBhoIn  BlogPostLanguage = "bho-in"
	BlogPostLanguageBi     BlogPostLanguage = "bi"
	BlogPostLanguageBm     BlogPostLanguage = "bm"
	BlogPostLanguageBmMl   BlogPostLanguage = "bm-ml"
	BlogPostLanguageBn     BlogPostLanguage = "bn"
	BlogPostLanguageBnBd   BlogPostLanguage = "bn-bd"
	BlogPostLanguageBnIn   BlogPostLanguage = "bn-in"
	BlogPostLanguageBo     BlogPostLanguage = "bo"
	BlogPostLanguageBoCn   BlogPostLanguage = "bo-cn"
	BlogPostLanguageBoIn   BlogPostLanguage = "bo-in"
	BlogPostLanguageBr     BlogPostLanguage = "br"
	BlogPostLanguageBrFr   BlogPostLanguage = "br-fr"
	BlogPostLanguageBrx    BlogPostLanguage = "brx"
	BlogPostLanguageBrxIn  BlogPostLanguage = "brx-in"
	BlogPostLanguageBs     BlogPostLanguage = "bs"
	BlogPostLanguageBsBa   BlogPostLanguage = "bs-ba"
	BlogPostLanguageCa     BlogPostLanguage = "ca"
	BlogPostLanguageCaAd   BlogPostLanguage = "ca-ad"
	BlogPostLanguageCaEs   BlogPostLanguage = "ca-es"
	BlogPostLanguageCaFr   BlogPostLanguage = "ca-fr"
	BlogPostLanguageCaIt   BlogPostLanguage = "ca-it"
	BlogPostLanguageCcp    BlogPostLanguage = "ccp"
	BlogPostLanguageCcpBd  BlogPostLanguage = "ccp-bd"
	BlogPostLanguageCcpIn  BlogPostLanguage = "ccp-in"
	BlogPostLanguageCe     BlogPostLanguage = "ce"
	BlogPostLanguageCeRu   BlogPostLanguage = "ce-ru"
	BlogPostLanguageCeb    BlogPostLanguage = "ceb"
	BlogPostLanguageCebPh  BlogPostLanguage = "ceb-ph"
	BlogPostLanguageCgg    BlogPostLanguage = "cgg"
	BlogPostLanguageCggUg  BlogPostLanguage = "cgg-ug"
	BlogPostLanguageCh     BlogPostLanguage = "ch"
	BlogPostLanguageChr    BlogPostLanguage = "chr"
	BlogPostLanguageChrUs  BlogPostLanguage = "chr-us"
	BlogPostLanguageCkb    BlogPostLanguage = "ckb"
	BlogPostLanguageCkbIq  BlogPostLanguage = "ckb-iq"
	BlogPostLanguageCkbIr  BlogPostLanguage = "ckb-ir"
	BlogPostLanguageCo     BlogPostLanguage = "co"
	BlogPostLanguageCr     BlogPostLanguage = "cr"
	BlogPostLanguageCs     BlogPostLanguage = "cs"
	BlogPostLanguageCsCz   BlogPostLanguage = "cs-cz"
	BlogPostLanguageCu     BlogPostLanguage = "cu"
	BlogPostLanguageCuRu   BlogPostLanguage = "cu-ru"
	BlogPostLanguageCv     BlogPostLanguage = "cv"
	BlogPostLanguageCvRu   BlogPostLanguage = "cv-ru"
	BlogPostLanguageCy     BlogPostLanguage = "cy"
	BlogPostLanguageCyGB   BlogPostLanguage = "cy-gb"
	BlogPostLanguageDa     BlogPostLanguage = "da"
	BlogPostLanguageDaDk   BlogPostLanguage = "da-dk"
	BlogPostLanguageDaGl   BlogPostLanguage = "da-gl"
	BlogPostLanguageDav    BlogPostLanguage = "dav"
	BlogPostLanguageDavKe  BlogPostLanguage = "dav-ke"
	BlogPostLanguageDe     BlogPostLanguage = "de"
	BlogPostLanguageDeAt   BlogPostLanguage = "de-at"
	BlogPostLanguageDeBe   BlogPostLanguage = "de-be"
	BlogPostLanguageDeCh   BlogPostLanguage = "de-ch"
	BlogPostLanguageDeDe   BlogPostLanguage = "de-de"
	BlogPostLanguageDeGr   BlogPostLanguage = "de-gr"
	BlogPostLanguageDeIt   BlogPostLanguage = "de-it"
	BlogPostLanguageDeLi   BlogPostLanguage = "de-li"
	BlogPostLanguageDeLu   BlogPostLanguage = "de-lu"
	BlogPostLanguageDje    BlogPostLanguage = "dje"
	BlogPostLanguageDjeNe  BlogPostLanguage = "dje-ne"
	BlogPostLanguageDoi    BlogPostLanguage = "doi"
	BlogPostLanguageDoiIn  BlogPostLanguage = "doi-in"
	BlogPostLanguageDsb    BlogPostLanguage = "dsb"
	BlogPostLanguageDsbDe  BlogPostLanguage = "dsb-de"
	BlogPostLanguageDua    BlogPostLanguage = "dua"
	BlogPostLanguageDuaCm  BlogPostLanguage = "dua-cm"
	BlogPostLanguageDv     BlogPostLanguage = "dv"
	BlogPostLanguageDyo    BlogPostLanguage = "dyo"
	BlogPostLanguageDyoSn  BlogPostLanguage = "dyo-sn"
	BlogPostLanguageDz     BlogPostLanguage = "dz"
	BlogPostLanguageDzBt   BlogPostLanguage = "dz-bt"
	BlogPostLanguageEbu    BlogPostLanguage = "ebu"
	BlogPostLanguageEbuKe  BlogPostLanguage = "ebu-ke"
	BlogPostLanguageEe     BlogPostLanguage = "ee"
	BlogPostLanguageEeGh   BlogPostLanguage = "ee-gh"
	BlogPostLanguageEeTg   BlogPostLanguage = "ee-tg"
	BlogPostLanguageEl     BlogPostLanguage = "el"
	BlogPostLanguageElCy   BlogPostLanguage = "el-cy"
	BlogPostLanguageElGr   BlogPostLanguage = "el-gr"
	BlogPostLanguageEn     BlogPostLanguage = "en"
	BlogPostLanguageEn001  BlogPostLanguage = "en-001"
	BlogPostLanguageEn150  BlogPostLanguage = "en-150"
	BlogPostLanguageEnAe   BlogPostLanguage = "en-ae"
	BlogPostLanguageEnAg   BlogPostLanguage = "en-ag"
	BlogPostLanguageEnAI   BlogPostLanguage = "en-ai"
	BlogPostLanguageEnAs   BlogPostLanguage = "en-as"
	BlogPostLanguageEnAt   BlogPostLanguage = "en-at"
	BlogPostLanguageEnAu   BlogPostLanguage = "en-au"
	BlogPostLanguageEnBb   BlogPostLanguage = "en-bb"
	BlogPostLanguageEnBe   BlogPostLanguage = "en-be"
	BlogPostLanguageEnBi   BlogPostLanguage = "en-bi"
	BlogPostLanguageEnBm   BlogPostLanguage = "en-bm"
	BlogPostLanguageEnBs   BlogPostLanguage = "en-bs"
	BlogPostLanguageEnBw   BlogPostLanguage = "en-bw"
	BlogPostLanguageEnBz   BlogPostLanguage = "en-bz"
	BlogPostLanguageEnCa   BlogPostLanguage = "en-ca"
	BlogPostLanguageEnCc   BlogPostLanguage = "en-cc"
	BlogPostLanguageEnCh   BlogPostLanguage = "en-ch"
	BlogPostLanguageEnCk   BlogPostLanguage = "en-ck"
	BlogPostLanguageEnCm   BlogPostLanguage = "en-cm"
	BlogPostLanguageEnCn   BlogPostLanguage = "en-cn"
	BlogPostLanguageEnCx   BlogPostLanguage = "en-cx"
	BlogPostLanguageEnCy   BlogPostLanguage = "en-cy"
	BlogPostLanguageEnDe   BlogPostLanguage = "en-de"
	BlogPostLanguageEnDg   BlogPostLanguage = "en-dg"
	BlogPostLanguageEnDk   BlogPostLanguage = "en-dk"
	BlogPostLanguageEnDm   BlogPostLanguage = "en-dm"
	BlogPostLanguageEnEe   BlogPostLanguage = "en-ee"
	BlogPostLanguageEnEg   BlogPostLanguage = "en-eg"
	BlogPostLanguageEnEr   BlogPostLanguage = "en-er"
	BlogPostLanguageEnEs   BlogPostLanguage = "en-es"
	BlogPostLanguageEnFi   BlogPostLanguage = "en-fi"
	BlogPostLanguageEnFj   BlogPostLanguage = "en-fj"
	BlogPostLanguageEnFk   BlogPostLanguage = "en-fk"
	BlogPostLanguageEnFm   BlogPostLanguage = "en-fm"
	BlogPostLanguageEnFr   BlogPostLanguage = "en-fr"
	BlogPostLanguageEnGB   BlogPostLanguage = "en-gb"
	BlogPostLanguageEnGd   BlogPostLanguage = "en-gd"
	BlogPostLanguageEnGg   BlogPostLanguage = "en-gg"
	BlogPostLanguageEnGh   BlogPostLanguage = "en-gh"
	BlogPostLanguageEnGi   BlogPostLanguage = "en-gi"
	BlogPostLanguageEnGm   BlogPostLanguage = "en-gm"
	BlogPostLanguageEnGu   BlogPostLanguage = "en-gu"
	BlogPostLanguageEnGy   BlogPostLanguage = "en-gy"
	BlogPostLanguageEnHk   BlogPostLanguage = "en-hk"
	BlogPostLanguageEnID   BlogPostLanguage = "en-id"
	BlogPostLanguageEnIe   BlogPostLanguage = "en-ie"
	BlogPostLanguageEnIl   BlogPostLanguage = "en-il"
	BlogPostLanguageEnIm   BlogPostLanguage = "en-im"
	BlogPostLanguageEnIn   BlogPostLanguage = "en-in"
	BlogPostLanguageEnIo   BlogPostLanguage = "en-io"
	BlogPostLanguageEnJe   BlogPostLanguage = "en-je"
	BlogPostLanguageEnJm   BlogPostLanguage = "en-jm"
	BlogPostLanguageEnKe   BlogPostLanguage = "en-ke"
	BlogPostLanguageEnKi   BlogPostLanguage = "en-ki"
	BlogPostLanguageEnKn   BlogPostLanguage = "en-kn"
	BlogPostLanguageEnKy   BlogPostLanguage = "en-ky"
	BlogPostLanguageEnLc   BlogPostLanguage = "en-lc"
	BlogPostLanguageEnLr   BlogPostLanguage = "en-lr"
	BlogPostLanguageEnLs   BlogPostLanguage = "en-ls"
	BlogPostLanguageEnLu   BlogPostLanguage = "en-lu"
	BlogPostLanguageEnMg   BlogPostLanguage = "en-mg"
	BlogPostLanguageEnMh   BlogPostLanguage = "en-mh"
	BlogPostLanguageEnMo   BlogPostLanguage = "en-mo"
	BlogPostLanguageEnMp   BlogPostLanguage = "en-mp"
	BlogPostLanguageEnMs   BlogPostLanguage = "en-ms"
	BlogPostLanguageEnMt   BlogPostLanguage = "en-mt"
	BlogPostLanguageEnMu   BlogPostLanguage = "en-mu"
	BlogPostLanguageEnMv   BlogPostLanguage = "en-mv"
	BlogPostLanguageEnMw   BlogPostLanguage = "en-mw"
	BlogPostLanguageEnMx   BlogPostLanguage = "en-mx"
	BlogPostLanguageEnMy   BlogPostLanguage = "en-my"
	BlogPostLanguageEnNa   BlogPostLanguage = "en-na"
	BlogPostLanguageEnNf   BlogPostLanguage = "en-nf"
	BlogPostLanguageEnNg   BlogPostLanguage = "en-ng"
	BlogPostLanguageEnNl   BlogPostLanguage = "en-nl"
	BlogPostLanguageEnNr   BlogPostLanguage = "en-nr"
	BlogPostLanguageEnNu   BlogPostLanguage = "en-nu"
	BlogPostLanguageEnNz   BlogPostLanguage = "en-nz"
	BlogPostLanguageEnPg   BlogPostLanguage = "en-pg"
	BlogPostLanguageEnPh   BlogPostLanguage = "en-ph"
	BlogPostLanguageEnPk   BlogPostLanguage = "en-pk"
	BlogPostLanguageEnPn   BlogPostLanguage = "en-pn"
	BlogPostLanguageEnPr   BlogPostLanguage = "en-pr"
	BlogPostLanguageEnPt   BlogPostLanguage = "en-pt"
	BlogPostLanguageEnPw   BlogPostLanguage = "en-pw"
	BlogPostLanguageEnRw   BlogPostLanguage = "en-rw"
	BlogPostLanguageEnSb   BlogPostLanguage = "en-sb"
	BlogPostLanguageEnSc   BlogPostLanguage = "en-sc"
	BlogPostLanguageEnSd   BlogPostLanguage = "en-sd"
	BlogPostLanguageEnSe   BlogPostLanguage = "en-se"
	BlogPostLanguageEnSg   BlogPostLanguage = "en-sg"
	BlogPostLanguageEnSh   BlogPostLanguage = "en-sh"
	BlogPostLanguageEnSi   BlogPostLanguage = "en-si"
	BlogPostLanguageEnSl   BlogPostLanguage = "en-sl"
	BlogPostLanguageEnSS   BlogPostLanguage = "en-ss"
	BlogPostLanguageEnSx   BlogPostLanguage = "en-sx"
	BlogPostLanguageEnSz   BlogPostLanguage = "en-sz"
	BlogPostLanguageEnTc   BlogPostLanguage = "en-tc"
	BlogPostLanguageEnTh   BlogPostLanguage = "en-th"
	BlogPostLanguageEnTk   BlogPostLanguage = "en-tk"
	BlogPostLanguageEnTn   BlogPostLanguage = "en-tn"
	BlogPostLanguageEnTo   BlogPostLanguage = "en-to"
	BlogPostLanguageEnTt   BlogPostLanguage = "en-tt"
	BlogPostLanguageEnTv   BlogPostLanguage = "en-tv"
	BlogPostLanguageEnTz   BlogPostLanguage = "en-tz"
	BlogPostLanguageEnUg   BlogPostLanguage = "en-ug"
	BlogPostLanguageEnUm   BlogPostLanguage = "en-um"
	BlogPostLanguageEnUs   BlogPostLanguage = "en-us"
	BlogPostLanguageEnVc   BlogPostLanguage = "en-vc"
	BlogPostLanguageEnVg   BlogPostLanguage = "en-vg"
	BlogPostLanguageEnVi   BlogPostLanguage = "en-vi"
	BlogPostLanguageEnVn   BlogPostLanguage = "en-vn"
	BlogPostLanguageEnVu   BlogPostLanguage = "en-vu"
	BlogPostLanguageEnWs   BlogPostLanguage = "en-ws"
	BlogPostLanguageEnZa   BlogPostLanguage = "en-za"
	BlogPostLanguageEnZm   BlogPostLanguage = "en-zm"
	BlogPostLanguageEnZw   BlogPostLanguage = "en-zw"
	BlogPostLanguageEo     BlogPostLanguage = "eo"
	BlogPostLanguageEo001  BlogPostLanguage = "eo-001"
	BlogPostLanguageEs     BlogPostLanguage = "es"
	BlogPostLanguageEs419  BlogPostLanguage = "es-419"
	BlogPostLanguageEsAr   BlogPostLanguage = "es-ar"
	BlogPostLanguageEsBo   BlogPostLanguage = "es-bo"
	BlogPostLanguageEsBr   BlogPostLanguage = "es-br"
	BlogPostLanguageEsBz   BlogPostLanguage = "es-bz"
	BlogPostLanguageEsCl   BlogPostLanguage = "es-cl"
	BlogPostLanguageEsCo   BlogPostLanguage = "es-co"
	BlogPostLanguageEsCr   BlogPostLanguage = "es-cr"
	BlogPostLanguageEsCu   BlogPostLanguage = "es-cu"
	BlogPostLanguageEsDo   BlogPostLanguage = "es-do"
	BlogPostLanguageEsEa   BlogPostLanguage = "es-ea"
	BlogPostLanguageEsEc   BlogPostLanguage = "es-ec"
	BlogPostLanguageEsEs   BlogPostLanguage = "es-es"
	BlogPostLanguageEsGq   BlogPostLanguage = "es-gq"
	BlogPostLanguageEsGt   BlogPostLanguage = "es-gt"
	BlogPostLanguageEsHn   BlogPostLanguage = "es-hn"
	BlogPostLanguageEsIc   BlogPostLanguage = "es-ic"
	BlogPostLanguageEsMx   BlogPostLanguage = "es-mx"
	BlogPostLanguageEsNi   BlogPostLanguage = "es-ni"
	BlogPostLanguageEsPa   BlogPostLanguage = "es-pa"
	BlogPostLanguageEsPe   BlogPostLanguage = "es-pe"
	BlogPostLanguageEsPh   BlogPostLanguage = "es-ph"
	BlogPostLanguageEsPr   BlogPostLanguage = "es-pr"
	BlogPostLanguageEsPy   BlogPostLanguage = "es-py"
	BlogPostLanguageEsSv   BlogPostLanguage = "es-sv"
	BlogPostLanguageEsUs   BlogPostLanguage = "es-us"
	BlogPostLanguageEsUy   BlogPostLanguage = "es-uy"
	BlogPostLanguageEsVe   BlogPostLanguage = "es-ve"
	BlogPostLanguageEt     BlogPostLanguage = "et"
	BlogPostLanguageEtEe   BlogPostLanguage = "et-ee"
	BlogPostLanguageEu     BlogPostLanguage = "eu"
	BlogPostLanguageEuEs   BlogPostLanguage = "eu-es"
	BlogPostLanguageEwo    BlogPostLanguage = "ewo"
	BlogPostLanguageEwoCm  BlogPostLanguage = "ewo-cm"
	BlogPostLanguageFa     BlogPostLanguage = "fa"
	BlogPostLanguageFaAf   BlogPostLanguage = "fa-af"
	BlogPostLanguageFaIr   BlogPostLanguage = "fa-ir"
	BlogPostLanguageFf     BlogPostLanguage = "ff"
	BlogPostLanguageFfBf   BlogPostLanguage = "ff-bf"
	BlogPostLanguageFfCm   BlogPostLanguage = "ff-cm"
	BlogPostLanguageFfGh   BlogPostLanguage = "ff-gh"
	BlogPostLanguageFfGm   BlogPostLanguage = "ff-gm"
	BlogPostLanguageFfGn   BlogPostLanguage = "ff-gn"
	BlogPostLanguageFfGw   BlogPostLanguage = "ff-gw"
	BlogPostLanguageFfLr   BlogPostLanguage = "ff-lr"
	BlogPostLanguageFfMr   BlogPostLanguage = "ff-mr"
	BlogPostLanguageFfNe   BlogPostLanguage = "ff-ne"
	BlogPostLanguageFfNg   BlogPostLanguage = "ff-ng"
	BlogPostLanguageFfSl   BlogPostLanguage = "ff-sl"
	BlogPostLanguageFfSn   BlogPostLanguage = "ff-sn"
	BlogPostLanguageFi     BlogPostLanguage = "fi"
	BlogPostLanguageFiFi   BlogPostLanguage = "fi-fi"
	BlogPostLanguageFil    BlogPostLanguage = "fil"
	BlogPostLanguageFilPh  BlogPostLanguage = "fil-ph"
	BlogPostLanguageFj     BlogPostLanguage = "fj"
	BlogPostLanguageFo     BlogPostLanguage = "fo"
	BlogPostLanguageFoDk   BlogPostLanguage = "fo-dk"
	BlogPostLanguageFoFo   BlogPostLanguage = "fo-fo"
	BlogPostLanguageFr     BlogPostLanguage = "fr"
	BlogPostLanguageFrBe   BlogPostLanguage = "fr-be"
	BlogPostLanguageFrBf   BlogPostLanguage = "fr-bf"
	BlogPostLanguageFrBi   BlogPostLanguage = "fr-bi"
	BlogPostLanguageFrBj   BlogPostLanguage = "fr-bj"
	BlogPostLanguageFrBl   BlogPostLanguage = "fr-bl"
	BlogPostLanguageFrCa   BlogPostLanguage = "fr-ca"
	BlogPostLanguageFrCd   BlogPostLanguage = "fr-cd"
	BlogPostLanguageFrCf   BlogPostLanguage = "fr-cf"
	BlogPostLanguageFrCg   BlogPostLanguage = "fr-cg"
	BlogPostLanguageFrCh   BlogPostLanguage = "fr-ch"
	BlogPostLanguageFrCi   BlogPostLanguage = "fr-ci"
	BlogPostLanguageFrCm   BlogPostLanguage = "fr-cm"
	BlogPostLanguageFrDj   BlogPostLanguage = "fr-dj"
	BlogPostLanguageFrDz   BlogPostLanguage = "fr-dz"
	BlogPostLanguageFrFr   BlogPostLanguage = "fr-fr"
	BlogPostLanguageFrGa   BlogPostLanguage = "fr-ga"
	BlogPostLanguageFrGf   BlogPostLanguage = "fr-gf"
	BlogPostLanguageFrGn   BlogPostLanguage = "fr-gn"
	BlogPostLanguageFrGp   BlogPostLanguage = "fr-gp"
	BlogPostLanguageFrGq   BlogPostLanguage = "fr-gq"
	BlogPostLanguageFrHt   BlogPostLanguage = "fr-ht"
	BlogPostLanguageFrKm   BlogPostLanguage = "fr-km"
	BlogPostLanguageFrLu   BlogPostLanguage = "fr-lu"
	BlogPostLanguageFrMa   BlogPostLanguage = "fr-ma"
	BlogPostLanguageFrMc   BlogPostLanguage = "fr-mc"
	BlogPostLanguageFrMf   BlogPostLanguage = "fr-mf"
	BlogPostLanguageFrMg   BlogPostLanguage = "fr-mg"
	BlogPostLanguageFrMl   BlogPostLanguage = "fr-ml"
	BlogPostLanguageFrMq   BlogPostLanguage = "fr-mq"
	BlogPostLanguageFrMr   BlogPostLanguage = "fr-mr"
	BlogPostLanguageFrMu   BlogPostLanguage = "fr-mu"
	BlogPostLanguageFrNc   BlogPostLanguage = "fr-nc"
	BlogPostLanguageFrNe   BlogPostLanguage = "fr-ne"
	BlogPostLanguageFrPf   BlogPostLanguage = "fr-pf"
	BlogPostLanguageFrPm   BlogPostLanguage = "fr-pm"
	BlogPostLanguageFrRe   BlogPostLanguage = "fr-re"
	BlogPostLanguageFrRw   BlogPostLanguage = "fr-rw"
	BlogPostLanguageFrSc   BlogPostLanguage = "fr-sc"
	BlogPostLanguageFrSn   BlogPostLanguage = "fr-sn"
	BlogPostLanguageFrSy   BlogPostLanguage = "fr-sy"
	BlogPostLanguageFrTd   BlogPostLanguage = "fr-td"
	BlogPostLanguageFrTg   BlogPostLanguage = "fr-tg"
	BlogPostLanguageFrTn   BlogPostLanguage = "fr-tn"
	BlogPostLanguageFrVu   BlogPostLanguage = "fr-vu"
	BlogPostLanguageFrWf   BlogPostLanguage = "fr-wf"
	BlogPostLanguageFrYt   BlogPostLanguage = "fr-yt"
	BlogPostLanguageFrr    BlogPostLanguage = "frr"
	BlogPostLanguageFrrDe  BlogPostLanguage = "frr-de"
	BlogPostLanguageFur    BlogPostLanguage = "fur"
	BlogPostLanguageFurIt  BlogPostLanguage = "fur-it"
	BlogPostLanguageFy     BlogPostLanguage = "fy"
	BlogPostLanguageFyNl   BlogPostLanguage = "fy-nl"
	BlogPostLanguageGa     BlogPostLanguage = "ga"
	BlogPostLanguageGaGB   BlogPostLanguage = "ga-gb"
	BlogPostLanguageGaIe   BlogPostLanguage = "ga-ie"
	BlogPostLanguageGd     BlogPostLanguage = "gd"
	BlogPostLanguageGdGB   BlogPostLanguage = "gd-gb"
	BlogPostLanguageGl     BlogPostLanguage = "gl"
	BlogPostLanguageGlEs   BlogPostLanguage = "gl-es"
	BlogPostLanguageGn     BlogPostLanguage = "gn"
	BlogPostLanguageGsw    BlogPostLanguage = "gsw"
	BlogPostLanguageGswCh  BlogPostLanguage = "gsw-ch"
	BlogPostLanguageGswFr  BlogPostLanguage = "gsw-fr"
	BlogPostLanguageGswLi  BlogPostLanguage = "gsw-li"
	BlogPostLanguageGu     BlogPostLanguage = "gu"
	BlogPostLanguageGuIn   BlogPostLanguage = "gu-in"
	BlogPostLanguageGuz    BlogPostLanguage = "guz"
	BlogPostLanguageGuzKe  BlogPostLanguage = "guz-ke"
	BlogPostLanguageGv     BlogPostLanguage = "gv"
	BlogPostLanguageGvIm   BlogPostLanguage = "gv-im"
	BlogPostLanguageHa     BlogPostLanguage = "ha"
	BlogPostLanguageHaGh   BlogPostLanguage = "ha-gh"
	BlogPostLanguageHaNe   BlogPostLanguage = "ha-ne"
	BlogPostLanguageHaNg   BlogPostLanguage = "ha-ng"
	BlogPostLanguageHaw    BlogPostLanguage = "haw"
	BlogPostLanguageHawUs  BlogPostLanguage = "haw-us"
	BlogPostLanguageHe     BlogPostLanguage = "he"
	BlogPostLanguageHeIl   BlogPostLanguage = "he-il"
	BlogPostLanguageHi     BlogPostLanguage = "hi"
	BlogPostLanguageHiIn   BlogPostLanguage = "hi-in"
	BlogPostLanguageHmn    BlogPostLanguage = "hmn"
	BlogPostLanguageHo     BlogPostLanguage = "ho"
	BlogPostLanguageHr     BlogPostLanguage = "hr"
	BlogPostLanguageHrBa   BlogPostLanguage = "hr-ba"
	BlogPostLanguageHrHr   BlogPostLanguage = "hr-hr"
	BlogPostLanguageHsb    BlogPostLanguage = "hsb"
	BlogPostLanguageHsbDe  BlogPostLanguage = "hsb-de"
	BlogPostLanguageHt     BlogPostLanguage = "ht"
	BlogPostLanguageHu     BlogPostLanguage = "hu"
	BlogPostLanguageHuHu   BlogPostLanguage = "hu-hu"
	BlogPostLanguageHy     BlogPostLanguage = "hy"
	BlogPostLanguageHyAm   BlogPostLanguage = "hy-am"
	BlogPostLanguageHz     BlogPostLanguage = "hz"
	BlogPostLanguageIa     BlogPostLanguage = "ia"
	BlogPostLanguageIa001  BlogPostLanguage = "ia-001"
	BlogPostLanguageID     BlogPostLanguage = "id"
	BlogPostLanguageIDID   BlogPostLanguage = "id-id"
	BlogPostLanguageIe     BlogPostLanguage = "ie"
	BlogPostLanguageIg     BlogPostLanguage = "ig"
	BlogPostLanguageIgNg   BlogPostLanguage = "ig-ng"
	BlogPostLanguageIi     BlogPostLanguage = "ii"
	BlogPostLanguageIiCn   BlogPostLanguage = "ii-cn"
	BlogPostLanguageIk     BlogPostLanguage = "ik"
	BlogPostLanguageIo     BlogPostLanguage = "io"
	BlogPostLanguageIs     BlogPostLanguage = "is"
	BlogPostLanguageIsIs   BlogPostLanguage = "is-is"
	BlogPostLanguageIt     BlogPostLanguage = "it"
	BlogPostLanguageItCh   BlogPostLanguage = "it-ch"
	BlogPostLanguageItIt   BlogPostLanguage = "it-it"
	BlogPostLanguageItSm   BlogPostLanguage = "it-sm"
	BlogPostLanguageItVa   BlogPostLanguage = "it-va"
	BlogPostLanguageIu     BlogPostLanguage = "iu"
	BlogPostLanguageJa     BlogPostLanguage = "ja"
	BlogPostLanguageJaJp   BlogPostLanguage = "ja-jp"
	BlogPostLanguageJgo    BlogPostLanguage = "jgo"
	BlogPostLanguageJgoCm  BlogPostLanguage = "jgo-cm"
	BlogPostLanguageJmc    BlogPostLanguage = "jmc"
	BlogPostLanguageJmcTz  BlogPostLanguage = "jmc-tz"
	BlogPostLanguageJv     BlogPostLanguage = "jv"
	BlogPostLanguageJvID   BlogPostLanguage = "jv-id"
	BlogPostLanguageKa     BlogPostLanguage = "ka"
	BlogPostLanguageKaGe   BlogPostLanguage = "ka-ge"
	BlogPostLanguageKab    BlogPostLanguage = "kab"
	BlogPostLanguageKabDz  BlogPostLanguage = "kab-dz"
	BlogPostLanguageKam    BlogPostLanguage = "kam"
	BlogPostLanguageKamKe  BlogPostLanguage = "kam-ke"
	BlogPostLanguageKar    BlogPostLanguage = "kar"
	BlogPostLanguageKde    BlogPostLanguage = "kde"
	BlogPostLanguageKdeTz  BlogPostLanguage = "kde-tz"
	BlogPostLanguageKea    BlogPostLanguage = "kea"
	BlogPostLanguageKeaCv  BlogPostLanguage = "kea-cv"
	BlogPostLanguageKg     BlogPostLanguage = "kg"
	BlogPostLanguageKgp    BlogPostLanguage = "kgp"
	BlogPostLanguageKgpBr  BlogPostLanguage = "kgp-br"
	BlogPostLanguageKh     BlogPostLanguage = "kh"
	BlogPostLanguageKhq    BlogPostLanguage = "khq"
	BlogPostLanguageKhqMl  BlogPostLanguage = "khq-ml"
	BlogPostLanguageKi     BlogPostLanguage = "ki"
	BlogPostLanguageKiKe   BlogPostLanguage = "ki-ke"
	BlogPostLanguageKj     BlogPostLanguage = "kj"
	BlogPostLanguageKk     BlogPostLanguage = "kk"
	BlogPostLanguageKkKz   BlogPostLanguage = "kk-kz"
	BlogPostLanguageKkj    BlogPostLanguage = "kkj"
	BlogPostLanguageKkjCm  BlogPostLanguage = "kkj-cm"
	BlogPostLanguageKl     BlogPostLanguage = "kl"
	BlogPostLanguageKlGl   BlogPostLanguage = "kl-gl"
	BlogPostLanguageKln    BlogPostLanguage = "kln"
	BlogPostLanguageKlnKe  BlogPostLanguage = "kln-ke"
	BlogPostLanguageKm     BlogPostLanguage = "km"
	BlogPostLanguageKmKh   BlogPostLanguage = "km-kh"
	BlogPostLanguageKn     BlogPostLanguage = "kn"
	BlogPostLanguageKnIn   BlogPostLanguage = "kn-in"
	BlogPostLanguageKo     BlogPostLanguage = "ko"
	BlogPostLanguageKoKp   BlogPostLanguage = "ko-kp"
	BlogPostLanguageKoKr   BlogPostLanguage = "ko-kr"
	BlogPostLanguageKok    BlogPostLanguage = "kok"
	BlogPostLanguageKokIn  BlogPostLanguage = "kok-in"
	BlogPostLanguageKr     BlogPostLanguage = "kr"
	BlogPostLanguageKs     BlogPostLanguage = "ks"
	BlogPostLanguageKsIn   BlogPostLanguage = "ks-in"
	BlogPostLanguageKsb    BlogPostLanguage = "ksb"
	BlogPostLanguageKsbTz  BlogPostLanguage = "ksb-tz"
	BlogPostLanguageKsf    BlogPostLanguage = "ksf"
	BlogPostLanguageKsfCm  BlogPostLanguage = "ksf-cm"
	BlogPostLanguageKsh    BlogPostLanguage = "ksh"
	BlogPostLanguageKshDe  BlogPostLanguage = "ksh-de"
	BlogPostLanguageKu     BlogPostLanguage = "ku"
	BlogPostLanguageKuTr   BlogPostLanguage = "ku-tr"
	BlogPostLanguageKv     BlogPostLanguage = "kv"
	BlogPostLanguageKw     BlogPostLanguage = "kw"
	BlogPostLanguageKwGB   BlogPostLanguage = "kw-gb"
	BlogPostLanguageKy     BlogPostLanguage = "ky"
	BlogPostLanguageKyKg   BlogPostLanguage = "ky-kg"
	BlogPostLanguageLa     BlogPostLanguage = "la"
	BlogPostLanguageLag    BlogPostLanguage = "lag"
	BlogPostLanguageLagTz  BlogPostLanguage = "lag-tz"
	BlogPostLanguageLb     BlogPostLanguage = "lb"
	BlogPostLanguageLbLu   BlogPostLanguage = "lb-lu"
	BlogPostLanguageLg     BlogPostLanguage = "lg"
	BlogPostLanguageLgUg   BlogPostLanguage = "lg-ug"
	BlogPostLanguageLi     BlogPostLanguage = "li"
	BlogPostLanguageLkt    BlogPostLanguage = "lkt"
	BlogPostLanguageLktUs  BlogPostLanguage = "lkt-us"
	BlogPostLanguageLn     BlogPostLanguage = "ln"
	BlogPostLanguageLnAo   BlogPostLanguage = "ln-ao"
	BlogPostLanguageLnCd   BlogPostLanguage = "ln-cd"
	BlogPostLanguageLnCf   BlogPostLanguage = "ln-cf"
	BlogPostLanguageLnCg   BlogPostLanguage = "ln-cg"
	BlogPostLanguageLo     BlogPostLanguage = "lo"
	BlogPostLanguageLoLa   BlogPostLanguage = "lo-la"
	BlogPostLanguageLrc    BlogPostLanguage = "lrc"
	BlogPostLanguageLrcIq  BlogPostLanguage = "lrc-iq"
	BlogPostLanguageLrcIr  BlogPostLanguage = "lrc-ir"
	BlogPostLanguageLt     BlogPostLanguage = "lt"
	BlogPostLanguageLtLt   BlogPostLanguage = "lt-lt"
	BlogPostLanguageLu     BlogPostLanguage = "lu"
	BlogPostLanguageLuCd   BlogPostLanguage = "lu-cd"
	BlogPostLanguageLuo    BlogPostLanguage = "luo"
	BlogPostLanguageLuoKe  BlogPostLanguage = "luo-ke"
	BlogPostLanguageLuy    BlogPostLanguage = "luy"
	BlogPostLanguageLuyKe  BlogPostLanguage = "luy-ke"
	BlogPostLanguageLv     BlogPostLanguage = "lv"
	BlogPostLanguageLvLv   BlogPostLanguage = "lv-lv"
	BlogPostLanguageMai    BlogPostLanguage = "mai"
	BlogPostLanguageMaiIn  BlogPostLanguage = "mai-in"
	BlogPostLanguageMas    BlogPostLanguage = "mas"
	BlogPostLanguageMasKe  BlogPostLanguage = "mas-ke"
	BlogPostLanguageMasTz  BlogPostLanguage = "mas-tz"
	BlogPostLanguageMdf    BlogPostLanguage = "mdf"
	BlogPostLanguageMdfRu  BlogPostLanguage = "mdf-ru"
	BlogPostLanguageMer    BlogPostLanguage = "mer"
	BlogPostLanguageMerKe  BlogPostLanguage = "mer-ke"
	BlogPostLanguageMfe    BlogPostLanguage = "mfe"
	BlogPostLanguageMfeMu  BlogPostLanguage = "mfe-mu"
	BlogPostLanguageMg     BlogPostLanguage = "mg"
	BlogPostLanguageMgMg   BlogPostLanguage = "mg-mg"
	BlogPostLanguageMgh    BlogPostLanguage = "mgh"
	BlogPostLanguageMghMz  BlogPostLanguage = "mgh-mz"
	BlogPostLanguageMgo    BlogPostLanguage = "mgo"
	BlogPostLanguageMgoCm  BlogPostLanguage = "mgo-cm"
	BlogPostLanguageMh     BlogPostLanguage = "mh"
	BlogPostLanguageMi     BlogPostLanguage = "mi"
	BlogPostLanguageMiNz   BlogPostLanguage = "mi-nz"
	BlogPostLanguageMk     BlogPostLanguage = "mk"
	BlogPostLanguageMkMk   BlogPostLanguage = "mk-mk"
	BlogPostLanguageMl     BlogPostLanguage = "ml"
	BlogPostLanguageMlIn   BlogPostLanguage = "ml-in"
	BlogPostLanguageMn     BlogPostLanguage = "mn"
	BlogPostLanguageMnMn   BlogPostLanguage = "mn-mn"
	BlogPostLanguageMni    BlogPostLanguage = "mni"
	BlogPostLanguageMniIn  BlogPostLanguage = "mni-in"
	BlogPostLanguageMr     BlogPostLanguage = "mr"
	BlogPostLanguageMrIn   BlogPostLanguage = "mr-in"
	BlogPostLanguageMs     BlogPostLanguage = "ms"
	BlogPostLanguageMsBn   BlogPostLanguage = "ms-bn"
	BlogPostLanguageMsID   BlogPostLanguage = "ms-id"
	BlogPostLanguageMsMy   BlogPostLanguage = "ms-my"
	BlogPostLanguageMsSg   BlogPostLanguage = "ms-sg"
	BlogPostLanguageMt     BlogPostLanguage = "mt"
	BlogPostLanguageMtMt   BlogPostLanguage = "mt-mt"
	BlogPostLanguageMua    BlogPostLanguage = "mua"
	BlogPostLanguageMuaCm  BlogPostLanguage = "mua-cm"
	BlogPostLanguageMy     BlogPostLanguage = "my"
	BlogPostLanguageMyMm   BlogPostLanguage = "my-mm"
	BlogPostLanguageMzn    BlogPostLanguage = "mzn"
	BlogPostLanguageMznIr  BlogPostLanguage = "mzn-ir"
	BlogPostLanguageNa     BlogPostLanguage = "na"
	BlogPostLanguageNaq    BlogPostLanguage = "naq"
	BlogPostLanguageNaqNa  BlogPostLanguage = "naq-na"
	BlogPostLanguageNb     BlogPostLanguage = "nb"
	BlogPostLanguageNbNo   BlogPostLanguage = "nb-no"
	BlogPostLanguageNbSj   BlogPostLanguage = "nb-sj"
	BlogPostLanguageNd     BlogPostLanguage = "nd"
	BlogPostLanguageNdZw   BlogPostLanguage = "nd-zw"
	BlogPostLanguageNds    BlogPostLanguage = "nds"
	BlogPostLanguageNdsDe  BlogPostLanguage = "nds-de"
	BlogPostLanguageNdsNl  BlogPostLanguage = "nds-nl"
	BlogPostLanguageNe     BlogPostLanguage = "ne"
	BlogPostLanguageNeIn   BlogPostLanguage = "ne-in"
	BlogPostLanguageNeNp   BlogPostLanguage = "ne-np"
	BlogPostLanguageNg     BlogPostLanguage = "ng"
	BlogPostLanguageNl     BlogPostLanguage = "nl"
	BlogPostLanguageNlAw   BlogPostLanguage = "nl-aw"
	BlogPostLanguageNlBe   BlogPostLanguage = "nl-be"
	BlogPostLanguageNlBq   BlogPostLanguage = "nl-bq"
	BlogPostLanguageNlCh   BlogPostLanguage = "nl-ch"
	BlogPostLanguageNlCw   BlogPostLanguage = "nl-cw"
	BlogPostLanguageNlLu   BlogPostLanguage = "nl-lu"
	BlogPostLanguageNlNl   BlogPostLanguage = "nl-nl"
	BlogPostLanguageNlSr   BlogPostLanguage = "nl-sr"
	BlogPostLanguageNlSx   BlogPostLanguage = "nl-sx"
	BlogPostLanguageNmg    BlogPostLanguage = "nmg"
	BlogPostLanguageNmgCm  BlogPostLanguage = "nmg-cm"
	BlogPostLanguageNn     BlogPostLanguage = "nn"
	BlogPostLanguageNnNo   BlogPostLanguage = "nn-no"
	BlogPostLanguageNnh    BlogPostLanguage = "nnh"
	BlogPostLanguageNnhCm  BlogPostLanguage = "nnh-cm"
	BlogPostLanguageNo     BlogPostLanguage = "no"
	BlogPostLanguageNoNo   BlogPostLanguage = "no-no"
	BlogPostLanguageNr     BlogPostLanguage = "nr"
	BlogPostLanguageNus    BlogPostLanguage = "nus"
	BlogPostLanguageNusSS  BlogPostLanguage = "nus-ss"
	BlogPostLanguageNv     BlogPostLanguage = "nv"
	BlogPostLanguageNy     BlogPostLanguage = "ny"
	BlogPostLanguageNyn    BlogPostLanguage = "nyn"
	BlogPostLanguageNynUg  BlogPostLanguage = "nyn-ug"
	BlogPostLanguageOc     BlogPostLanguage = "oc"
	BlogPostLanguageOcEs   BlogPostLanguage = "oc-es"
	BlogPostLanguageOcFr   BlogPostLanguage = "oc-fr"
	BlogPostLanguageOj     BlogPostLanguage = "oj"
	BlogPostLanguageOm     BlogPostLanguage = "om"
	BlogPostLanguageOmEt   BlogPostLanguage = "om-et"
	BlogPostLanguageOmKe   BlogPostLanguage = "om-ke"
	BlogPostLanguageOr     BlogPostLanguage = "or"
	BlogPostLanguageOrIn   BlogPostLanguage = "or-in"
	BlogPostLanguageOs     BlogPostLanguage = "os"
	BlogPostLanguageOsGe   BlogPostLanguage = "os-ge"
	BlogPostLanguageOsRu   BlogPostLanguage = "os-ru"
	BlogPostLanguagePa     BlogPostLanguage = "pa"
	BlogPostLanguagePaIn   BlogPostLanguage = "pa-in"
	BlogPostLanguagePaPk   BlogPostLanguage = "pa-pk"
	BlogPostLanguagePcm    BlogPostLanguage = "pcm"
	BlogPostLanguagePcmNg  BlogPostLanguage = "pcm-ng"
	BlogPostLanguagePi     BlogPostLanguage = "pi"
	BlogPostLanguagePis    BlogPostLanguage = "pis"
	BlogPostLanguagePisSb  BlogPostLanguage = "pis-sb"
	BlogPostLanguagePl     BlogPostLanguage = "pl"
	BlogPostLanguagePlPl   BlogPostLanguage = "pl-pl"
	BlogPostLanguagePrg    BlogPostLanguage = "prg"
	BlogPostLanguagePrg001 BlogPostLanguage = "prg-001"
	BlogPostLanguagePs     BlogPostLanguage = "ps"
	BlogPostLanguagePsAf   BlogPostLanguage = "ps-af"
	BlogPostLanguagePsPk   BlogPostLanguage = "ps-pk"
	BlogPostLanguagePt     BlogPostLanguage = "pt"
	BlogPostLanguagePtAo   BlogPostLanguage = "pt-ao"
	BlogPostLanguagePtBr   BlogPostLanguage = "pt-br"
	BlogPostLanguagePtCh   BlogPostLanguage = "pt-ch"
	BlogPostLanguagePtCv   BlogPostLanguage = "pt-cv"
	BlogPostLanguagePtGq   BlogPostLanguage = "pt-gq"
	BlogPostLanguagePtGw   BlogPostLanguage = "pt-gw"
	BlogPostLanguagePtLu   BlogPostLanguage = "pt-lu"
	BlogPostLanguagePtMo   BlogPostLanguage = "pt-mo"
	BlogPostLanguagePtMz   BlogPostLanguage = "pt-mz"
	BlogPostLanguagePtPt   BlogPostLanguage = "pt-pt"
	BlogPostLanguagePtSt   BlogPostLanguage = "pt-st"
	BlogPostLanguagePtTl   BlogPostLanguage = "pt-tl"
	BlogPostLanguageQu     BlogPostLanguage = "qu"
	BlogPostLanguageQuBo   BlogPostLanguage = "qu-bo"
	BlogPostLanguageQuEc   BlogPostLanguage = "qu-ec"
	BlogPostLanguageQuPe   BlogPostLanguage = "qu-pe"
	BlogPostLanguageRaj    BlogPostLanguage = "raj"
	BlogPostLanguageRajIn  BlogPostLanguage = "raj-in"
	BlogPostLanguageRm     BlogPostLanguage = "rm"
	BlogPostLanguageRmCh   BlogPostLanguage = "rm-ch"
	BlogPostLanguageRn     BlogPostLanguage = "rn"
	BlogPostLanguageRnBi   BlogPostLanguage = "rn-bi"
	BlogPostLanguageRo     BlogPostLanguage = "ro"
	BlogPostLanguageRoMd   BlogPostLanguage = "ro-md"
	BlogPostLanguageRoRo   BlogPostLanguage = "ro-ro"
	BlogPostLanguageRof    BlogPostLanguage = "rof"
	BlogPostLanguageRofTz  BlogPostLanguage = "rof-tz"
	BlogPostLanguageRu     BlogPostLanguage = "ru"
	BlogPostLanguageRuBy   BlogPostLanguage = "ru-by"
	BlogPostLanguageRuKg   BlogPostLanguage = "ru-kg"
	BlogPostLanguageRuKz   BlogPostLanguage = "ru-kz"
	BlogPostLanguageRuMd   BlogPostLanguage = "ru-md"
	BlogPostLanguageRuRu   BlogPostLanguage = "ru-ru"
	BlogPostLanguageRuUa   BlogPostLanguage = "ru-ua"
	BlogPostLanguageRw     BlogPostLanguage = "rw"
	BlogPostLanguageRwRw   BlogPostLanguage = "rw-rw"
	BlogPostLanguageRwk    BlogPostLanguage = "rwk"
	BlogPostLanguageRwkTz  BlogPostLanguage = "rwk-tz"
	BlogPostLanguageSa     BlogPostLanguage = "sa"
	BlogPostLanguageSaIn   BlogPostLanguage = "sa-in"
	BlogPostLanguageSah    BlogPostLanguage = "sah"
	BlogPostLanguageSahRu  BlogPostLanguage = "sah-ru"
	BlogPostLanguageSaq    BlogPostLanguage = "saq"
	BlogPostLanguageSaqKe  BlogPostLanguage = "saq-ke"
	BlogPostLanguageSat    BlogPostLanguage = "sat"
	BlogPostLanguageSatIn  BlogPostLanguage = "sat-in"
	BlogPostLanguageSbp    BlogPostLanguage = "sbp"
	BlogPostLanguageSbpTz  BlogPostLanguage = "sbp-tz"
	BlogPostLanguageSc     BlogPostLanguage = "sc"
	BlogPostLanguageScIt   BlogPostLanguage = "sc-it"
	BlogPostLanguageSd     BlogPostLanguage = "sd"
	BlogPostLanguageSdIn   BlogPostLanguage = "sd-in"
	BlogPostLanguageSdPk   BlogPostLanguage = "sd-pk"
	BlogPostLanguageSe     BlogPostLanguage = "se"
	BlogPostLanguageSeFi   BlogPostLanguage = "se-fi"
	BlogPostLanguageSeNo   BlogPostLanguage = "se-no"
	BlogPostLanguageSeSe   BlogPostLanguage = "se-se"
	BlogPostLanguageSeh    BlogPostLanguage = "seh"
	BlogPostLanguageSehMz  BlogPostLanguage = "seh-mz"
	BlogPostLanguageSes    BlogPostLanguage = "ses"
	BlogPostLanguageSesMl  BlogPostLanguage = "ses-ml"
	BlogPostLanguageSg     BlogPostLanguage = "sg"
	BlogPostLanguageSgCf   BlogPostLanguage = "sg-cf"
	BlogPostLanguageShi    BlogPostLanguage = "shi"
	BlogPostLanguageShiMa  BlogPostLanguage = "shi-ma"
	BlogPostLanguageSi     BlogPostLanguage = "si"
	BlogPostLanguageSiLk   BlogPostLanguage = "si-lk"
	BlogPostLanguageSk     BlogPostLanguage = "sk"
	BlogPostLanguageSkSk   BlogPostLanguage = "sk-sk"
	BlogPostLanguageSl     BlogPostLanguage = "sl"
	BlogPostLanguageSlSi   BlogPostLanguage = "sl-si"
	BlogPostLanguageSm     BlogPostLanguage = "sm"
	BlogPostLanguageSmn    BlogPostLanguage = "smn"
	BlogPostLanguageSmnFi  BlogPostLanguage = "smn-fi"
	BlogPostLanguageSMS    BlogPostLanguage = "sms"
	BlogPostLanguageSMSFi  BlogPostLanguage = "sms-fi"
	BlogPostLanguageSn     BlogPostLanguage = "sn"
	BlogPostLanguageSnZw   BlogPostLanguage = "sn-zw"
	BlogPostLanguageSo     BlogPostLanguage = "so"
	BlogPostLanguageSoDj   BlogPostLanguage = "so-dj"
	BlogPostLanguageSoEt   BlogPostLanguage = "so-et"
	BlogPostLanguageSoKe   BlogPostLanguage = "so-ke"
	BlogPostLanguageSoSo   BlogPostLanguage = "so-so"
	BlogPostLanguageSq     BlogPostLanguage = "sq"
	BlogPostLanguageSqAl   BlogPostLanguage = "sq-al"
	BlogPostLanguageSqMk   BlogPostLanguage = "sq-mk"
	BlogPostLanguageSqXk   BlogPostLanguage = "sq-xk"
	BlogPostLanguageSr     BlogPostLanguage = "sr"
	BlogPostLanguageSrBa   BlogPostLanguage = "sr-ba"
	BlogPostLanguageSrCs   BlogPostLanguage = "sr-cs"
	BlogPostLanguageSrMe   BlogPostLanguage = "sr-me"
	BlogPostLanguageSrRs   BlogPostLanguage = "sr-rs"
	BlogPostLanguageSrXk   BlogPostLanguage = "sr-xk"
	BlogPostLanguageSS     BlogPostLanguage = "ss"
	BlogPostLanguageSt     BlogPostLanguage = "st"
	BlogPostLanguageSu     BlogPostLanguage = "su"
	BlogPostLanguageSuID   BlogPostLanguage = "su-id"
	BlogPostLanguageSv     BlogPostLanguage = "sv"
	BlogPostLanguageSvAx   BlogPostLanguage = "sv-ax"
	BlogPostLanguageSvFi   BlogPostLanguage = "sv-fi"
	BlogPostLanguageSvSe   BlogPostLanguage = "sv-se"
	BlogPostLanguageSw     BlogPostLanguage = "sw"
	BlogPostLanguageSwCd   BlogPostLanguage = "sw-cd"
	BlogPostLanguageSwKe   BlogPostLanguage = "sw-ke"
	BlogPostLanguageSwTz   BlogPostLanguage = "sw-tz"
	BlogPostLanguageSwUg   BlogPostLanguage = "sw-ug"
	BlogPostLanguageSy     BlogPostLanguage = "sy"
	BlogPostLanguageTa     BlogPostLanguage = "ta"
	BlogPostLanguageTaIn   BlogPostLanguage = "ta-in"
	BlogPostLanguageTaLk   BlogPostLanguage = "ta-lk"
	BlogPostLanguageTaMy   BlogPostLanguage = "ta-my"
	BlogPostLanguageTaSg   BlogPostLanguage = "ta-sg"
	BlogPostLanguageTe     BlogPostLanguage = "te"
	BlogPostLanguageTeIn   BlogPostLanguage = "te-in"
	BlogPostLanguageTeo    BlogPostLanguage = "teo"
	BlogPostLanguageTeoKe  BlogPostLanguage = "teo-ke"
	BlogPostLanguageTeoUg  BlogPostLanguage = "teo-ug"
	BlogPostLanguageTg     BlogPostLanguage = "tg"
	BlogPostLanguageTgTj   BlogPostLanguage = "tg-tj"
	BlogPostLanguageTh     BlogPostLanguage = "th"
	BlogPostLanguageThTh   BlogPostLanguage = "th-th"
	BlogPostLanguageTi     BlogPostLanguage = "ti"
	BlogPostLanguageTiEr   BlogPostLanguage = "ti-er"
	BlogPostLanguageTiEt   BlogPostLanguage = "ti-et"
	BlogPostLanguageTk     BlogPostLanguage = "tk"
	BlogPostLanguageTkTm   BlogPostLanguage = "tk-tm"
	BlogPostLanguageTl     BlogPostLanguage = "tl"
	BlogPostLanguageTn     BlogPostLanguage = "tn"
	BlogPostLanguageTo     BlogPostLanguage = "to"
	BlogPostLanguageToTo   BlogPostLanguage = "to-to"
	BlogPostLanguageTok    BlogPostLanguage = "tok"
	BlogPostLanguageTok001 BlogPostLanguage = "tok-001"
	BlogPostLanguageTr     BlogPostLanguage = "tr"
	BlogPostLanguageTrCy   BlogPostLanguage = "tr-cy"
	BlogPostLanguageTrTr   BlogPostLanguage = "tr-tr"
	BlogPostLanguageTs     BlogPostLanguage = "ts"
	BlogPostLanguageTt     BlogPostLanguage = "tt"
	BlogPostLanguageTtRu   BlogPostLanguage = "tt-ru"
	BlogPostLanguageTw     BlogPostLanguage = "tw"
	BlogPostLanguageTwq    BlogPostLanguage = "twq"
	BlogPostLanguageTwqNe  BlogPostLanguage = "twq-ne"
	BlogPostLanguageTy     BlogPostLanguage = "ty"
	BlogPostLanguageTzm    BlogPostLanguage = "tzm"
	BlogPostLanguageTzmMa  BlogPostLanguage = "tzm-ma"
	BlogPostLanguageUg     BlogPostLanguage = "ug"
	BlogPostLanguageUgCn   BlogPostLanguage = "ug-cn"
	BlogPostLanguageUk     BlogPostLanguage = "uk"
	BlogPostLanguageUkUa   BlogPostLanguage = "uk-ua"
	BlogPostLanguageUr     BlogPostLanguage = "ur"
	BlogPostLanguageUrIn   BlogPostLanguage = "ur-in"
	BlogPostLanguageUrPk   BlogPostLanguage = "ur-pk"
	BlogPostLanguageUz     BlogPostLanguage = "uz"
	BlogPostLanguageUzAf   BlogPostLanguage = "uz-af"
	BlogPostLanguageUzUz   BlogPostLanguage = "uz-uz"
	BlogPostLanguageVai    BlogPostLanguage = "vai"
	BlogPostLanguageVaiLr  BlogPostLanguage = "vai-lr"
	BlogPostLanguageVe     BlogPostLanguage = "ve"
	BlogPostLanguageVi     BlogPostLanguage = "vi"
	BlogPostLanguageViVn   BlogPostLanguage = "vi-vn"
	BlogPostLanguageVo     BlogPostLanguage = "vo"
	BlogPostLanguageVo001  BlogPostLanguage = "vo-001"
	BlogPostLanguageVun    BlogPostLanguage = "vun"
	BlogPostLanguageVunTz  BlogPostLanguage = "vun-tz"
	BlogPostLanguageWa     BlogPostLanguage = "wa"
	BlogPostLanguageWae    BlogPostLanguage = "wae"
	BlogPostLanguageWaeCh  BlogPostLanguage = "wae-ch"
	BlogPostLanguageWo     BlogPostLanguage = "wo"
	BlogPostLanguageWoSn   BlogPostLanguage = "wo-sn"
	BlogPostLanguageXh     BlogPostLanguage = "xh"
	BlogPostLanguageXhZa   BlogPostLanguage = "xh-za"
	BlogPostLanguageXog    BlogPostLanguage = "xog"
	BlogPostLanguageXogUg  BlogPostLanguage = "xog-ug"
	BlogPostLanguageYav    BlogPostLanguage = "yav"
	BlogPostLanguageYavCm  BlogPostLanguage = "yav-cm"
	BlogPostLanguageYi     BlogPostLanguage = "yi"
	BlogPostLanguageYi001  BlogPostLanguage = "yi-001"
	BlogPostLanguageYo     BlogPostLanguage = "yo"
	BlogPostLanguageYoBj   BlogPostLanguage = "yo-bj"
	BlogPostLanguageYoNg   BlogPostLanguage = "yo-ng"
	BlogPostLanguageYrl    BlogPostLanguage = "yrl"
	BlogPostLanguageYrlBr  BlogPostLanguage = "yrl-br"
	BlogPostLanguageYrlCo  BlogPostLanguage = "yrl-co"
	BlogPostLanguageYrlVe  BlogPostLanguage = "yrl-ve"
	BlogPostLanguageYue    BlogPostLanguage = "yue"
	BlogPostLanguageYueCn  BlogPostLanguage = "yue-cn"
	BlogPostLanguageYueHk  BlogPostLanguage = "yue-hk"
	BlogPostLanguageZa     BlogPostLanguage = "za"
	BlogPostLanguageZgh    BlogPostLanguage = "zgh"
	BlogPostLanguageZghMa  BlogPostLanguage = "zgh-ma"
	BlogPostLanguageZh     BlogPostLanguage = "zh"
	BlogPostLanguageZhCn   BlogPostLanguage = "zh-cn"
	BlogPostLanguageZhHans BlogPostLanguage = "zh-hans"
	BlogPostLanguageZhHant BlogPostLanguage = "zh-hant"
	BlogPostLanguageZhHk   BlogPostLanguage = "zh-hk"
	BlogPostLanguageZhMo   BlogPostLanguage = "zh-mo"
	BlogPostLanguageZhSg   BlogPostLanguage = "zh-sg"
	BlogPostLanguageZhTw   BlogPostLanguage = "zh-tw"
	BlogPostLanguageZu     BlogPostLanguage = "zu"
	BlogPostLanguageZuZa   BlogPostLanguage = "zu-za"
)

// The property ID is required.
type BlogPostLanguageCloneRequestVNextParam struct {
	// ID of blog post to clone.
	ID string `json:"id" api:"required"`
	// Target language of new variant.
	Language param.Opt[string] `json:"language,omitzero"`
	paramObj
}

func (r BlogPostLanguageCloneRequestVNextParam) MarshalJSON() (data []byte, err error) {
	type shadow BlogPostLanguageCloneRequestVNextParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BlogPostLanguageCloneRequestVNextParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Hidden, Margin, Padding are required.
type BreakpointStylesParam struct {
	// Boolean indicating if the breakpoint is visible.
	Hidden  bool         `json:"hidden" api:"required"`
	Margin  MarginParam  `json:"margin,omitzero" api:"required"`
	Padding PaddingParam `json:"padding,omitzero" api:"required"`
	paramObj
}

func (r BreakpointStylesParam) MarshalJSON() (data []byte, err error) {
	type shadow BreakpointStylesParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BreakpointStylesParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Color is required.
type ColorStopParam struct {
	Color RgbaColorParam `json:"color,omitzero" api:"required"`
	paramObj
}

func (r ColorStopParam) MarshalJSON() (data []byte, err error) {
	type shadow ColorStopParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ColorStopParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property ID is required.
type ContentCloneRequestVNextParam struct {
	// ID of the object to be cloned.
	ID string `json:"id" api:"required"`
	// Name of the cloned object.
	CloneName param.Opt[string] `json:"cloneName,omitzero"`
	paramObj
}

func (r ContentCloneRequestVNextParam) MarshalJSON() (data []byte, err error) {
	type shadow ContentCloneRequestVNextParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ContentCloneRequestVNextParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties ID, ArchivedInDashboard, AuthorName, Campaign, CampaignName,
// Created, Name, Password, PublicAccessRules, PublicAccessRulesEnabled,
// PublishDate, Slug, State, Updated are required.
type ContentLanguageVariationParam struct {
	// The unique ID of the content language variation.
	ID int64 `json:"id" api:"required"`
	// If True, the variant will not show up in your dashboard, although the post could
	// still be live.
	ArchivedInDashboard bool `json:"archivedInDashboard" api:"required"`
	// The name of the user who last published the blog post. For posts that haven't
	// been published yet, this property will reflect the user who initially created
	// the draft.
	AuthorName string `json:"authorName" api:"required"`
	// The GUID of the marketing campaign this page is a part of.
	Campaign string `json:"campaign" api:"required"`
	// Name of the associated marketing campaign.
	CampaignName string `json:"campaignName" api:"required"`
	// The timestamp (ISO8601 format) when this Blog Post was created.
	Created time.Time `json:"created" api:"required" format:"date-time"`
	// The internal name of the content language variation.
	Name string `json:"name" api:"required"`
	// Set this to create a password protected page. Entering the password will be
	// required to view the page.
	Password          string             `json:"password" api:"required"`
	PublicAccessRules []PublicAccessRule `json:"publicAccessRules,omitzero" api:"required"`
	// Boolean to determine whether or not to respect publicAccessRules.
	PublicAccessRulesEnabled bool `json:"publicAccessRulesEnabled" api:"required"`
	// The date (ISO8601 format) the page is to be published at.
	PublishDate time.Time `json:"publishDate" api:"required" format:"date-time"`
	// The path of the this page. This field is appended to the domain to construct the
	// url of this page.
	Slug string `json:"slug" api:"required"`
	// An ENUM describing the current state of this page.
	//
	// Maximum string length: 25
	State string `json:"state" api:"required"`
	// The timestamp (ISO8601 format) when this Blog Post was updated.
	Updated time.Time `json:"updated" api:"required" format:"date-time"`
	TagIDs  []int64   `json:"tagIds,omitzero"`
	paramObj
}

func (r ContentLanguageVariationParam) MarshalJSON() (data []byte, err error) {
	type shadow ContentLanguageVariationParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ContentLanguageVariationParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties ID, PublishDate are required.
type ContentScheduleRequestVNextParam struct {
	// The ID of the object to be scheduled.
	ID string `json:"id" api:"required"`
	// The date the object should transition from scheduled to published.
	PublishDate time.Time `json:"publishDate" api:"required" format:"date-time"`
	paramObj
}

func (r ContentScheduleRequestVNextParam) MarshalJSON() (data []byte, err error) {
	type shadow ContentScheduleRequestVNextParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ContentScheduleRequestVNextParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Angle, Colors, SideOrCorner are required.
type GradientParam struct {
	Angle        AngleParam        `json:"angle,omitzero" api:"required"`
	Colors       []ColorStopParam  `json:"colors,omitzero" api:"required"`
	SideOrCorner SideOrCornerParam `json:"sideOrCorner,omitzero" api:"required"`
	paramObj
}

func (r GradientParam) MarshalJSON() (data []byte, err error) {
	type shadow GradientParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *GradientParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Cells, CssClass, CssID, CssStyle, Label, Name, Params,
// RowMetaData, Rows, Styles, Type, W, X are required.
type LayoutSectionParam struct {
	Cells []LayoutSectionParam `json:"cells,omitzero" api:"required"`
	// The CSS class applied to the layout section.
	CssClass string `json:"cssClass" api:"required"`
	// The CSS ID applied to the layout section.
	CssID string `json:"cssId" api:"required"`
	// Custom CSS styles applied to the layout section.
	CssStyle string `json:"cssStyle" api:"required"`
	// The label for the layout section.
	Label string `json:"label" api:"required"`
	// The name assigned to the layout section.
	Name string `json:"name" api:"required"`
	// Parameters associated with the layout section.
	Params      map[string]any                  `json:"params,omitzero" api:"required"`
	RowMetaData []RowMetaDataParam              `json:"rowMetaData,omitzero" api:"required"`
	Rows        []map[string]LayoutSectionParam `json:"rows,omitzero" api:"required"`
	Styles      StylesParam                     `json:"styles,omitzero" api:"required"`
	// The type of the layout section.
	Type string `json:"type" api:"required"`
	// The width of the layout section.
	W int64 `json:"w" api:"required"`
	// The x-coordinate position of the layout section.
	X int64 `json:"x" api:"required"`
	paramObj
}

func (r LayoutSectionParam) MarshalJSON() (data []byte, err error) {
	type shadow LayoutSectionParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LayoutSectionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Bottom, Top are required.
type MarginParam struct {
	Bottom SizeParam `json:"bottom,omitzero" api:"required"`
	Top    SizeParam `json:"top,omitzero" api:"required"`
	paramObj
}

func (r MarginParam) MarshalJSON() (data []byte, err error) {
	type shadow MarginParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MarginParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Bottom, Left, Right, Top are required.
type PaddingParam struct {
	Bottom SizeParam `json:"bottom,omitzero" api:"required"`
	Left   SizeParam `json:"left,omitzero" api:"required"`
	Right  SizeParam `json:"right,omitzero" api:"required"`
	Top    SizeParam `json:"top,omitzero" api:"required"`
	paramObj
}

func (r PaddingParam) MarshalJSON() (data []byte, err error) {
	type shadow PaddingParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PaddingParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties A, B, G, R are required.
type RgbaColorParam struct {
	// Alpha.
	A float64 `json:"a" api:"required"`
	// Blue.
	B int64 `json:"b" api:"required"`
	// Green.
	G int64 `json:"g" api:"required"`
	// Red.
	R int64 `json:"r" api:"required"`
	paramObj
}

func (r RgbaColorParam) MarshalJSON() (data []byte, err error) {
	type shadow RgbaColorParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RgbaColorParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties CssClass, Styles are required.
type RowMetaDataParam struct {
	// The CSS class applied to the row.
	CssClass string      `json:"cssClass" api:"required"`
	Styles   StylesParam `json:"styles,omitzero" api:"required"`
	paramObj
}

func (r RowMetaDataParam) MarshalJSON() (data []byte, err error) {
	type shadow RowMetaDataParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RowMetaDataParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties HorizontalSide, VerticalSide are required.
type SideOrCornerParam struct {
	// Specifies the horizontal side of an element.
	//
	// Any of "CENTER", "LEFT", "RIGHT".
	HorizontalSide SideOrCornerHorizontalSide `json:"horizontalSide,omitzero" api:"required"`
	// Specifies the vertical side of an element.
	//
	// Any of "BOTTOM", "MIDDLE", "TOP".
	VerticalSide SideOrCornerVerticalSide `json:"verticalSide,omitzero" api:"required"`
	paramObj
}

func (r SideOrCornerParam) MarshalJSON() (data []byte, err error) {
	type shadow SideOrCornerParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SideOrCornerParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Specifies the horizontal side of an element.
type SideOrCornerHorizontalSide string

const (
	SideOrCornerHorizontalSideCenter SideOrCornerHorizontalSide = "CENTER"
	SideOrCornerHorizontalSideLeft   SideOrCornerHorizontalSide = "LEFT"
	SideOrCornerHorizontalSideRight  SideOrCornerHorizontalSide = "RIGHT"
)

// Specifies the vertical side of an element.
type SideOrCornerVerticalSide string

const (
	SideOrCornerVerticalSideBottom SideOrCornerVerticalSide = "BOTTOM"
	SideOrCornerVerticalSideMiddle SideOrCornerVerticalSide = "MIDDLE"
	SideOrCornerVerticalSideTop    SideOrCornerVerticalSide = "TOP"
)

// The properties Units, Value are required.
type SizeParam struct {
	// Any of "%", "ch", "cm", "em", "ex", "in", "lh", "mm", "pc", "pt", "px", "Q",
	// "rem", "vh", "vmax", "vmin", "vw".
	Units SizeUnits `json:"units,omitzero" api:"required"`
	Value float64   `json:"value" api:"required"`
	paramObj
}

func (r SizeParam) MarshalJSON() (data []byte, err error) {
	type shadow SizeParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SizeParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SizeUnits string

const (
	SizeUnitsUnknown0 SizeUnits = "%"
	SizeUnitsCh       SizeUnits = "ch"
	SizeUnitsCm       SizeUnits = "cm"
	SizeUnitsEm       SizeUnits = "em"
	SizeUnitsEx       SizeUnits = "ex"
	SizeUnitsIn       SizeUnits = "in"
	SizeUnitsLh       SizeUnits = "lh"
	SizeUnitsMm       SizeUnits = "mm"
	SizeUnitsPc       SizeUnits = "pc"
	SizeUnitsPt       SizeUnits = "pt"
	SizeUnitsPx       SizeUnits = "px"
	SizeUnitsQ        SizeUnits = "Q"
	SizeUnitsRem      SizeUnits = "rem"
	SizeUnitsVh       SizeUnits = "vh"
	SizeUnitsVmax     SizeUnits = "vmax"
	SizeUnitsVmin     SizeUnits = "vmin"
	SizeUnitsVw       SizeUnits = "vw"
)

// The properties BackgroundColor, BackgroundGradient, BackgroundImage,
// FlexboxPositioning, ForceFullWidthSection, MaxWidthSectionCentering,
// VerticalAlignment are required.
type StylesParam struct {
	BackgroundColor    RgbaColorParam       `json:"backgroundColor,omitzero" api:"required"`
	BackgroundGradient GradientParam        `json:"backgroundGradient,omitzero" api:"required"`
	BackgroundImage    BackgroundImageParam `json:"backgroundImage,omitzero" api:"required"`
	// Indicates whether flexbox positioning is enabled for the section.
	//
	// Any of "BOTTOM_CENTER", "BOTTOM_LEFT", "BOTTOM_RIGHT", "MIDDLE_CENTER",
	// "MIDDLE_LEFT", "MIDDLE_RIGHT", "TOP_CENTER", "TOP_LEFT", "TOP_RIGHT".
	FlexboxPositioning StylesFlexboxPositioning `json:"flexboxPositioning,omitzero" api:"required"`
	// Determines if the section should be forced to full width.
	ForceFullWidthSection bool `json:"forceFullWidthSection" api:"required"`
	// Defines the maximum width for centering the section.
	MaxWidthSectionCentering int64 `json:"maxWidthSectionCentering" api:"required"`
	// Specifies the vertical alignment of elements within the section.
	//
	// Any of "BOTTOM", "MIDDLE", "TOP".
	VerticalAlignment StylesVerticalAlignment `json:"verticalAlignment,omitzero" api:"required"`
	// Breakpoint CSS styles for margin, padding, etc...
	BreakpointStyles map[string]BreakpointStylesParam `json:"breakpointStyles,omitzero"`
	paramObj
}

func (r StylesParam) MarshalJSON() (data []byte, err error) {
	type shadow StylesParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *StylesParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Indicates whether flexbox positioning is enabled for the section.
type StylesFlexboxPositioning string

const (
	StylesFlexboxPositioningBottomCenter StylesFlexboxPositioning = "BOTTOM_CENTER"
	StylesFlexboxPositioningBottomLeft   StylesFlexboxPositioning = "BOTTOM_LEFT"
	StylesFlexboxPositioningBottomRight  StylesFlexboxPositioning = "BOTTOM_RIGHT"
	StylesFlexboxPositioningMiddleCenter StylesFlexboxPositioning = "MIDDLE_CENTER"
	StylesFlexboxPositioningMiddleLeft   StylesFlexboxPositioning = "MIDDLE_LEFT"
	StylesFlexboxPositioningMiddleRight  StylesFlexboxPositioning = "MIDDLE_RIGHT"
	StylesFlexboxPositioningTopCenter    StylesFlexboxPositioning = "TOP_CENTER"
	StylesFlexboxPositioningTopLeft      StylesFlexboxPositioning = "TOP_LEFT"
	StylesFlexboxPositioningTopRight     StylesFlexboxPositioning = "TOP_RIGHT"
)

// Specifies the vertical alignment of elements within the section.
type StylesVerticalAlignment string

const (
	StylesVerticalAlignmentBottom StylesVerticalAlignment = "BOTTOM"
	StylesVerticalAlignmentMiddle StylesVerticalAlignment = "MIDDLE"
	StylesVerticalAlignmentTop    StylesVerticalAlignment = "TOP"
)

type BlogPostNewParams struct {
	BlogPost BlogPostParam
	paramObj
}

func (r BlogPostNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.BlogPost)
}
func (r *BlogPostNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BlogPostUpdateParams struct {
	BlogPost BlogPostParam
	// Whether to return only results that have been archived.
	Archived param.Opt[bool] `query:"archived,omitzero" json:"-"`
	paramObj
}

func (r BlogPostUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.BlogPost)
}
func (r *BlogPostUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [BlogPostUpdateParams]'s query parameters as `url.Values`.
func (r BlogPostUpdateParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type BlogPostListParams struct {
	// The paging cursor token of the last successfully read resource will be returned
	// as the `paging.next.after` JSON property of a paged response containing more
	// results.
	After param.Opt[string] `query:"after,omitzero" json:"-"`
	// Whether to return only results that have been archived.
	Archived      param.Opt[bool]      `query:"archived,omitzero" json:"-"`
	CreatedAfter  param.Opt[time.Time] `query:"createdAfter,omitzero" format:"date-time" json:"-"`
	CreatedAt     param.Opt[time.Time] `query:"createdAt,omitzero" format:"date-time" json:"-"`
	CreatedBefore param.Opt[time.Time] `query:"createdBefore,omitzero" format:"date-time" json:"-"`
	// The maximum number of results to display per page.
	Limit         param.Opt[int64]     `query:"limit,omitzero" json:"-"`
	Property      param.Opt[string]    `query:"property,omitzero" json:"-"`
	UpdatedAfter  param.Opt[time.Time] `query:"updatedAfter,omitzero" format:"date-time" json:"-"`
	UpdatedAt     param.Opt[time.Time] `query:"updatedAt,omitzero" format:"date-time" json:"-"`
	UpdatedBefore param.Opt[time.Time] `query:"updatedBefore,omitzero" format:"date-time" json:"-"`
	Sort          []string             `query:"sort,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [BlogPostListParams]'s query parameters as `url.Values`.
func (r BlogPostListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type BlogPostDeleteParams struct {
	// Whether to return only results that have been archived.
	Archived param.Opt[bool] `query:"archived,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [BlogPostDeleteParams]'s query parameters as `url.Values`.
func (r BlogPostDeleteParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type BlogPostAttachToLangGroupParams struct {
	AttachToLangPrimaryRequestVNext AttachToLangPrimaryRequestVNextParam
	paramObj
}

func (r BlogPostAttachToLangGroupParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.AttachToLangPrimaryRequestVNext)
}
func (r *BlogPostAttachToLangGroupParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BlogPostCloneParams struct {
	ContentCloneRequestVNext ContentCloneRequestVNextParam
	paramObj
}

func (r BlogPostCloneParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.ContentCloneRequestVNext)
}
func (r *BlogPostCloneParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BlogPostNewLangVariationParams struct {
	BlogPostLanguageCloneRequestVNext BlogPostLanguageCloneRequestVNextParam
	paramObj
}

func (r BlogPostNewLangVariationParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.BlogPostLanguageCloneRequestVNext)
}
func (r *BlogPostNewLangVariationParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BlogPostDetachFromLangGroupParams struct {
	DetachFromLangGroupRequestVNext DetachFromLangGroupRequestVNextParam
	paramObj
}

func (r BlogPostDetachFromLangGroupParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.DetachFromLangGroupRequestVNext)
}
func (r *BlogPostDetachFromLangGroupParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BlogPostGetParams struct {
	// Whether to return only results that have been archived.
	Archived param.Opt[bool]   `query:"archived,omitzero" json:"-"`
	Property param.Opt[string] `query:"property,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [BlogPostGetParams]'s query parameters as `url.Values`.
func (r BlogPostGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type BlogPostGetPreviousVersionParams struct {
	ObjectID string `path:"objectId" api:"required" json:"-"`
	paramObj
}

type BlogPostGetPreviousVersionsParams struct {
	// The paging cursor token of the last successfully read resource will be returned
	// as the `paging.next.after` JSON property of a paged response containing more
	// results.
	After  param.Opt[string] `query:"after,omitzero" json:"-"`
	Before param.Opt[string] `query:"before,omitzero" json:"-"`
	// The maximum number of results to display per page.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [BlogPostGetPreviousVersionsParams]'s query parameters as
// `url.Values`.
func (r BlogPostGetPreviousVersionsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type BlogPostRestorePreviousVersionParams struct {
	ObjectID string `path:"objectId" api:"required" json:"-"`
	paramObj
}

type BlogPostRestorePreviousVersionToDraftParams struct {
	ObjectID string `path:"objectId" api:"required" json:"-"`
	paramObj
}

type BlogPostScheduleParams struct {
	ContentScheduleRequestVNext ContentScheduleRequestVNextParam
	paramObj
}

func (r BlogPostScheduleParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.ContentScheduleRequestVNext)
}
func (r *BlogPostScheduleParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BlogPostSetLangPrimaryParams struct {
	SetNewLanguagePrimaryRequestVNext SetNewLanguagePrimaryRequestVNextParam
	paramObj
}

func (r BlogPostSetLangPrimaryParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.SetNewLanguagePrimaryRequestVNext)
}
func (r *BlogPostSetLangPrimaryParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BlogPostUpdateDraftParams struct {
	BlogPost BlogPostParam
	paramObj
}

func (r BlogPostUpdateDraftParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.BlogPost)
}
func (r *BlogPostUpdateDraftParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BlogPostUpdateLangsParams struct {
	UpdateLanguagesRequestVNext UpdateLanguagesRequestVNextParam
	paramObj
}

func (r BlogPostUpdateLangsParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.UpdateLanguagesRequestVNext)
}
func (r *BlogPostUpdateLangsParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
