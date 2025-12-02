// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cms

import (
	"context"
	"encoding/json"
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
	"github.com/stainless-sdks/hubspot-sdk-go/packages/pagination"
	"github.com/stainless-sdks/hubspot-sdk-go/packages/param"
	"github.com/stainless-sdks/hubspot-sdk-go/packages/respjson"
	"github.com/stainless-sdks/hubspot-sdk-go/shared"
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

// Create a new blog post, specifying its content in the request body.
func (r *BlogPostService) New(ctx context.Context, body BlogPostNewParams, opts ...option.RequestOption) (res *BlogPost, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "cms/v3/blogs/posts"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Partially updates a single blog post by ID. You only need to specify the values
// that you want to update.
func (r *BlogPostService) Update(ctx context.Context, objectID string, params BlogPostUpdateParams, opts ...option.RequestOption) (res *BlogPost, err error) {
	opts = slices.Concat(r.Options, opts)
	if objectID == "" {
		err = errors.New("missing required objectId parameter")
		return
	}
	path := fmt.Sprintf("cms/v3/blogs/posts/%s", objectID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return
}

// Retrieve all blog posts, with paging and filtering options. This method would be
// useful for an integration that ingests posts and suggests edits.
func (r *BlogPostService) List(ctx context.Context, query BlogPostListParams, opts ...option.RequestOption) (res *pagination.Page[BlogPost], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "cms/v3/blogs/posts"
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

// Retrieve all blog posts, with paging and filtering options. This method would be
// useful for an integration that ingests posts and suggests edits.
func (r *BlogPostService) ListAutoPaging(ctx context.Context, query BlogPostListParams, opts ...option.RequestOption) *pagination.PageAutoPager[BlogPost] {
	return pagination.NewPageAutoPager(r.List(ctx, query, opts...))
}

// Delete a blog post by ID.
func (r *BlogPostService) Delete(ctx context.Context, objectID string, body BlogPostDeleteParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if objectID == "" {
		err = errors.New("missing required objectId parameter")
		return
	}
	path := fmt.Sprintf("cms/v3/blogs/posts/%s", objectID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, nil, opts...)
	return
}

// Attach a blog post to a
// [multi-language group](https://developers.hubspot.com/docs/guides/cms/content/multi-language-content).
func (r *BlogPostService) AttachToLangGroup(ctx context.Context, body BlogPostAttachToLangGroupParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "cms/v3/blogs/posts/multi-language/attach-to-lang-group"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Clone a blog post, making a copy of it in a new blog post.
func (r *BlogPostService) Clone(ctx context.Context, body BlogPostCloneParams, opts ...option.RequestOption) (res *BlogPost, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "cms/v3/blogs/posts/clone"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Create a new language variation from an existing blog post
func (r *BlogPostService) NewLangVariation(ctx context.Context, body BlogPostNewLangVariationParams, opts ...option.RequestOption) (res *BlogPost, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "cms/v3/blogs/posts/multi-language/create-language-variation"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Detach a blog post from a
// [multi-language group](https://developers.hubspot.com/docs/guides/cms/content/multi-language-content).
func (r *BlogPostService) DetachFromLangGroup(ctx context.Context, body BlogPostDetachFromLangGroupParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "cms/v3/blogs/posts/multi-language/detach-from-lang-group"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Retrieve a blog post by the post ID.
func (r *BlogPostService) Get(ctx context.Context, objectID string, query BlogPostGetParams, opts ...option.RequestOption) (res *BlogPost, err error) {
	opts = slices.Concat(r.Options, opts)
	if objectID == "" {
		err = errors.New("missing required objectId parameter")
		return
	}
	path := fmt.Sprintf("cms/v3/blogs/posts/%s", objectID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Retrieve the full draft version of a blog post.
func (r *BlogPostService) GetDraftByID(ctx context.Context, objectID string, opts ...option.RequestOption) (res *BlogPost, err error) {
	opts = slices.Concat(r.Options, opts)
	if objectID == "" {
		err = errors.New("missing required objectId parameter")
		return
	}
	path := fmt.Sprintf("cms/v3/blogs/posts/%s/draft", objectID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Retrieve a previous version of a blog post.
func (r *BlogPostService) GetPreviousVersion(ctx context.Context, revisionID string, query BlogPostGetPreviousVersionParams, opts ...option.RequestOption) (res *VersionBlogPost, err error) {
	opts = slices.Concat(r.Options, opts)
	if query.ObjectID == "" {
		err = errors.New("missing required objectId parameter")
		return
	}
	if revisionID == "" {
		err = errors.New("missing required revisionId parameter")
		return
	}
	path := fmt.Sprintf("cms/v3/blogs/posts/%s/revisions/%s", query.ObjectID, revisionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Retrieve all the previous versions of a blog post.
func (r *BlogPostService) GetPreviousVersions(ctx context.Context, objectID string, query BlogPostGetPreviousVersionsParams, opts ...option.RequestOption) (res *pagination.Page[VersionBlogPost], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if objectID == "" {
		err = errors.New("missing required objectId parameter")
		return
	}
	path := fmt.Sprintf("cms/v3/blogs/posts/%s/revisions", objectID)
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

// Retrieve all the previous versions of a blog post.
func (r *BlogPostService) GetPreviousVersionsAutoPaging(ctx context.Context, objectID string, query BlogPostGetPreviousVersionsParams, opts ...option.RequestOption) *pagination.PageAutoPager[VersionBlogPost] {
	return pagination.NewPageAutoPager(r.GetPreviousVersions(ctx, objectID, query, opts...))
}

// Publish the draft version of the blog post, sending its content to the live
// page.
func (r *BlogPostService) PushLive(ctx context.Context, objectID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if objectID == "" {
		err = errors.New("missing required objectId parameter")
		return
	}
	path := fmt.Sprintf("cms/v3/blogs/posts/%s/draft/push-live", objectID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, nil, opts...)
	return
}

// Discard all drafted content, resetting the draft to contain the content in the
// currently published version.
func (r *BlogPostService) ResetDraft(ctx context.Context, objectID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if objectID == "" {
		err = errors.New("missing required objectId parameter")
		return
	}
	path := fmt.Sprintf("cms/v3/blogs/posts/%s/draft/reset", objectID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, nil, opts...)
	return
}

// Restores a blog post to one of its previous versions.
func (r *BlogPostService) RestorePreviousVersion(ctx context.Context, revisionID string, body BlogPostRestorePreviousVersionParams, opts ...option.RequestOption) (res *BlogPost, err error) {
	opts = slices.Concat(r.Options, opts)
	if body.ObjectID == "" {
		err = errors.New("missing required objectId parameter")
		return
	}
	if revisionID == "" {
		err = errors.New("missing required revisionId parameter")
		return
	}
	path := fmt.Sprintf("cms/v3/blogs/posts/%s/revisions/%s/restore", body.ObjectID, revisionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Takes a specified version of a blog post, sets it as the new draft version of
// the blog post.
func (r *BlogPostService) RestorePreviousVersionToDraft(ctx context.Context, revisionID int64, body BlogPostRestorePreviousVersionToDraftParams, opts ...option.RequestOption) (res *BlogPost, err error) {
	opts = slices.Concat(r.Options, opts)
	if body.ObjectID == "" {
		err = errors.New("missing required objectId parameter")
		return
	}
	path := fmt.Sprintf("cms/v3/blogs/posts/%s/revisions/%v/restore-to-draft", body.ObjectID, revisionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Schedule a blog post to be published at a specified time.
func (r *BlogPostService) Schedule(ctx context.Context, body BlogPostScheduleParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "cms/v3/blogs/posts/schedule"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Set the primary language of a
// [multi-language group](https://developers.hubspot.com/docs/guides/cms/content/multi-language-content)
// to the language of the provided post (specified as an ID in the request body)
func (r *BlogPostService) SetLangPrimary(ctx context.Context, body BlogPostSetLangPrimaryParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "cms/v3/blogs/posts/multi-language/set-new-lang-primary"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, nil, opts...)
	return
}

// Partially updates the draft version of a single blog post by ID. You only need
// to specify the values that you want to update.
func (r *BlogPostService) UpdateDraft(ctx context.Context, objectID string, body BlogPostUpdateDraftParams, opts ...option.RequestOption) (res *BlogPost, err error) {
	opts = slices.Concat(r.Options, opts)
	if objectID == "" {
		err = errors.New("missing required objectId parameter")
		return
	}
	path := fmt.Sprintf("cms/v3/blogs/posts/%s/draft", objectID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Explicitly set new languages for each post in a
// [multi-language group](https://developers.hubspot.com/docs/guides/cms/content/multi-language-content).
func (r *BlogPostService) UpdateLangs(ctx context.Context, body BlogPostUpdateLangsParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "cms/v3/blogs/posts/multi-language/update-languages"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Wrapper for providing an array of blog posts as inputs.
//
// The property Inputs is required.
type BatchInputBlogPostParam struct {
	// Blog posts to input.
	Inputs []BlogPostParam `json:"inputs,omitzero,required"`
	paramObj
}

func (r BatchInputBlogPostParam) MarshalJSON() (data []byte, err error) {
	type shadow BatchInputBlogPostParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BatchInputBlogPostParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Response object for batch operations on blog posts.
type BatchResponseBlogPost struct {
	// Time of batch operation completion.
	CompletedAt time.Time `json:"completedAt,required" format:"date-time"`
	// Results of batch operation.
	Results []BlogPost `json:"results,required"`
	// Time of batch operation start.
	StartedAt time.Time `json:"startedAt,required" format:"date-time"`
	// Status of batch operation.
	//
	// Any of "CANCELED", "COMPLETE", "PENDING", "PROCESSING".
	Status BatchResponseBlogPostStatus `json:"status,required"`
	// Links associated with batch operation.
	Links map[string]string `json:"links"`
	// Time of batch operation request.
	RequestedAt time.Time `json:"requestedAt" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CompletedAt respjson.Field
		Results     respjson.Field
		StartedAt   respjson.Field
		Status      respjson.Field
		Links       respjson.Field
		RequestedAt respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BatchResponseBlogPost) RawJSON() string { return r.JSON.raw }
func (r *BatchResponseBlogPost) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Status of batch operation.
type BatchResponseBlogPostStatus string

const (
	BatchResponseBlogPostStatusCanceled   BatchResponseBlogPostStatus = "CANCELED"
	BatchResponseBlogPostStatusComplete   BatchResponseBlogPostStatus = "COMPLETE"
	BatchResponseBlogPostStatusPending    BatchResponseBlogPostStatus = "PENDING"
	BatchResponseBlogPostStatusProcessing BatchResponseBlogPostStatus = "PROCESSING"
)

// Model definition for a Blog Post.
type BlogPost struct {
	// The unique ID of the blog post.
	ID string `json:"id,required"`
	// Any of "automated_loser_variant", "automated_master", "automated_variant",
	// "loser_variant", "mab_master", "mab_variant", "master", "variant".
	AbStatus BlogPostAbStatus `json:"abStatus,required"`
	AbTestID string           `json:"abTestId,required"`
	// The timestamp (ISO8601 format) when this Blog Post was deleted.
	ArchivedAt int64 `json:"archivedAt,required"`
	// If True, the post will not show up in your dashboard, although the post could
	// still be live.
	ArchivedInDashboard bool `json:"archivedInDashboard,required"`
	// List of stylesheets to attach to this blog post. These stylesheets are attached
	// to just this page. Order of precedence is bottom to top, just like in the HTML.
	AttachedStylesheets []map[string]any `json:"attachedStylesheets,required"`
	// The name of the blog author associated with the post.
	AuthorName string `json:"authorName,required"`
	// The ID of the blog author associated with this post.
	BlogAuthorID string `json:"blogAuthorId,required"`
	// The GUID of the marketing campaign the post is associated with.
	Campaign string `json:"campaign,required"`
	// ID of the object type.
	CategoryID int64 `json:"categoryId,required"`
	// The ID of the post's parent blog.
	ContentGroupID string `json:"contentGroupId,required"`
	// An ENUM descibing the type of this object. Should always be BLOG_POST.
	//
	// Any of "0", "1", "10", "11", "12", "13", "14", "15", "2", "3", "4", "5", "6",
	// "7", "8", "9".
	ContentTypeCategory BlogPostContentTypeCategory `json:"contentTypeCategory,required"`
	Created             time.Time                   `json:"created,required" format:"date-time"`
	// The ID of the user that created the post.
	CreatedByID        string `json:"createdById,required"`
	CurrentlyPublished bool   `json:"currentlyPublished,required"`
	// A generated ENUM descibing the current state of this Blog Post. Should always
	// match state.
	//
	// Any of "AUTOMATED", "AUTOMATED_AB", "AUTOMATED_AB_VARIANT", "AUTOMATED_DRAFT",
	// "AUTOMATED_DRAFT_AB", "AUTOMATED_DRAFT_ABVARIANT", "AUTOMATED_FOR_FORM",
	// "AUTOMATED_FOR_FORM_BUFFER", "AUTOMATED_FOR_FORM_DRAFT",
	// "AUTOMATED_FOR_FORM_LEGACY", "AUTOMATED_LOSER_ABVARIANT", "AUTOMATED_SENDING",
	// "BLOG_EMAIL_DRAFT", "BLOG_EMAIL_PUBLISHED", "DRAFT", "DRAFT_AB",
	// "DRAFT_AB_VARIANT", "ERROR", "LOSER_AB_VARIANT", "PAGE_STUB", "PRE_PROCESSING",
	// "PROCESSING", "PUBLISHED", "PUBLISHED_AB", "PUBLISHED_AB_VARIANT",
	// "PUBLISHED_OR_SCHEDULED", "RSS_TO_EMAIL_DRAFT", "RSS_TO_EMAIL_PUBLISHED",
	// "SCHEDULED", "SCHEDULED_AB", "SCHEDULED_OR_PUBLISHED".
	CurrentState BlogPostCurrentState `json:"currentState,required"`
	// The domain that the post lives on. If null, the post will default to the domain
	// of the parent blog.
	Domain                    string `json:"domain,required"`
	DynamicPageDataSourceID   string `json:"dynamicPageDataSourceId,required"`
	DynamicPageDataSourceType int64  `json:"dynamicPageDataSourceType,required"`
	// For dynamic HubDB pages, the ID of the HubDB table this post references.
	DynamicPageHubDBTableID string `json:"dynamicPageHubDbTableId,required"`
	// Boolean to determine whether or not the styles from the template should be
	// applied.
	EnableDomainStylesheets bool `json:"enableDomainStylesheets,required"`
	// Boolean to allow overriding the AMP settings for the blog.
	EnableGoogleAmpOutputOverride bool `json:"enableGoogleAmpOutputOverride,required"`
	// Boolean to determine whether or not the styles from the template should be
	// applied.
	EnableLayoutStylesheets bool `json:"enableLayoutStylesheets,required"`
	// The featuredImage of this Blog Post.
	FeaturedImage string `json:"featuredImage,required"`
	// Alt Text of the featuredImage.
	FeaturedImageAltText string `json:"featuredImageAltText,required"`
	FolderID             string `json:"folderId,required"`
	// Custom HTML for embed codes, javascript that should be placed before the </body>
	// tag of the page.
	FooterHTML string `json:"footerHtml,required"`
	// Custom HTML for embed codes, javascript, etc. that goes in the <head> tag of the
	// page.
	HeadHTML string `json:"headHtml,required"`
	// The HTML title of the post.
	HTMLTitle string `json:"htmlTitle,required"`
	// Boolean to determine whether or not the Primary CSS Files should be applied.
	IncludeDefaultCustomCss bool `json:"includeDefaultCustomCss,required"`
	// The explicitly defined ISO 639 language code of the post. If null, the post will
	// default to the language of the parent blog.
	//
	// Any of "af", "af-na", "af-za", "agq", "agq-cm", "ak", "ak-gh", "am", "am-et",
	// "ar", "ar-001", "ar-ae", "ar-bh", "ar-dj", "ar-dz", "ar-eg", "ar-eh", "ar-er",
	// "ar-il", "ar-iq", "ar-jo", "ar-km", "ar-kw", "ar-lb", "ar-ly", "ar-ma", "ar-mr",
	// "ar-om", "ar-ps", "ar-qa", "ar-sa", "ar-sd", "ar-so", "ar-ss", "ar-sy", "ar-td",
	// "ar-tn", "ar-ye", "as", "as-in", "asa", "asa-tz", "ast", "ast-es", "az",
	// "az-az", "bas", "bas-cm", "be", "be-by", "bem", "bem-zm", "bez", "bez-tz", "bg",
	// "bg-bg", "bm", "bm-ml", "bn", "bn-bd", "bn-in", "bo", "bo-cn", "bo-in", "br",
	// "br-fr", "brx", "brx-in", "bs", "bs-ba", "ca", "ca-ad", "ca-es", "ca-fr",
	// "ca-it", "ccp", "ccp-bd", "ccp-in", "ce", "ce-ru", "ceb", "ceb-ph", "cgg",
	// "cgg-ug", "chr", "chr-us", "ckb", "ckb-iq", "ckb-ir", "cs", "cs-cz", "cu",
	// "cu-ru", "cy", "cy-gb", "da", "da-dk", "da-gl", "dav", "dav-ke", "de", "de-at",
	// "de-be", "de-ch", "de-de", "de-gr", "de-it", "de-li", "de-lu", "dje", "dje-ne",
	// "doi", "doi-in", "dsb", "dsb-de", "dua", "dua-cm", "dyo", "dyo-sn", "dz",
	// "dz-bt", "ebu", "ebu-ke", "ee", "ee-gh", "ee-tg", "el", "el-cy", "el-gr", "en",
	// "en-001", "en-150", "en-ae", "en-ag", "en-ai", "en-as", "en-at", "en-au",
	// "en-bb", "en-be", "en-bi", "en-bm", "en-bs", "en-bw", "en-bz", "en-ca", "en-cc",
	// "en-ch", "en-ck", "en-cm", "en-cn", "en-cx", "en-cy", "en-de", "en-dg", "en-dk",
	// "en-dm", "en-ee", "en-er", "en-fi", "en-fj", "en-fk", "en-fm", "en-fr", "en-gb",
	// "en-gd", "en-gg", "en-gh", "en-gi", "en-gm", "en-gu", "en-gy", "en-hk", "en-ie",
	// "en-il", "en-im", "en-in", "en-io", "en-je", "en-jm", "en-ke", "en-ki", "en-kn",
	// "en-ky", "en-lc", "en-lr", "en-ls", "en-lu", "en-mg", "en-mh", "en-mo", "en-mp",
	// "en-ms", "en-mt", "en-mu", "en-mw", "en-mx", "en-my", "en-na", "en-nf", "en-ng",
	// "en-nl", "en-nr", "en-nu", "en-nz", "en-pg", "en-ph", "en-pk", "en-pn", "en-pr",
	// "en-pw", "en-rw", "en-sb", "en-sc", "en-sd", "en-se", "en-sg", "en-sh", "en-si",
	// "en-sl", "en-ss", "en-sx", "en-sz", "en-tc", "en-tk", "en-to", "en-tt", "en-tv",
	// "en-tz", "en-ug", "en-um", "en-us", "en-vc", "en-vg", "en-vi", "en-vu", "en-ws",
	// "en-za", "en-zm", "en-zw", "eo", "eo-001", "es", "es-419", "es-ar", "es-bo",
	// "es-br", "es-bz", "es-cl", "es-co", "es-cr", "es-cu", "es-do", "es-ea", "es-ec",
	// "es-es", "es-gq", "es-gt", "es-hn", "es-ic", "es-mx", "es-ni", "es-pa", "es-pe",
	// "es-ph", "es-pr", "es-py", "es-sv", "es-us", "es-uy", "es-ve", "et", "et-ee",
	// "eu", "eu-es", "ewo", "ewo-cm", "fa", "fa-af", "fa-ir", "ff", "ff-bf", "ff-cm",
	// "ff-gh", "ff-gm", "ff-gn", "ff-gw", "ff-lr", "ff-mr", "ff-ne", "ff-ng", "ff-sl",
	// "ff-sn", "fi", "fi-fi", "fil", "fil-ph", "fo", "fo-dk", "fo-fo", "fr", "fr-be",
	// "fr-bf", "fr-bi", "fr-bj", "fr-bl", "fr-ca", "fr-cd", "fr-cf", "fr-cg", "fr-ch",
	// "fr-ci", "fr-cm", "fr-dj", "fr-dz", "fr-fr", "fr-ga", "fr-gf", "fr-gn", "fr-gp",
	// "fr-gq", "fr-ht", "fr-km", "fr-lu", "fr-ma", "fr-mc", "fr-mf", "fr-mg", "fr-ml",
	// "fr-mq", "fr-mr", "fr-mu", "fr-nc", "fr-ne", "fr-pf", "fr-pm", "fr-re", "fr-rw",
	// "fr-sc", "fr-sn", "fr-sy", "fr-td", "fr-tg", "fr-tn", "fr-vu", "fr-wf", "fr-yt",
	// "fur", "fur-it", "fy", "fy-nl", "ga", "ga-gb", "ga-ie", "gd", "gd-gb", "gl",
	// "gl-es", "gsw", "gsw-ch", "gsw-fr", "gsw-li", "gu", "gu-in", "guz", "guz-ke",
	// "gv", "gv-im", "ha", "ha-gh", "ha-ne", "ha-ng", "haw", "haw-us", "he", "he-il",
	// "hi", "hi-in", "hr", "hr-ba", "hr-hr", "hsb", "hsb-de", "hu", "hu-hu", "hy",
	// "hy-am", "ia", "ia-001", "id", "id-id", "ig", "ig-ng", "ii", "ii-cn", "is",
	// "is-is", "it", "it-ch", "it-it", "it-sm", "it-va", "ja", "ja-jp", "jgo",
	// "jgo-cm", "jmc", "jmc-tz", "jv", "jv-id", "ka", "ka-ge", "kab", "kab-dz", "kam",
	// "kam-ke", "kde", "kde-tz", "kea", "kea-cv", "khq", "khq-ml", "ki", "ki-ke",
	// "kk", "kk-kz", "kkj", "kkj-cm", "kl", "kl-gl", "kln", "kln-ke", "km", "km-kh",
	// "kn", "kn-in", "ko", "ko-kp", "ko-kr", "kok", "kok-in", "ks", "ks-in", "ksb",
	// "ksb-tz", "ksf", "ksf-cm", "ksh", "ksh-de", "ku", "ku-tr", "kw", "kw-gb", "ky",
	// "ky-kg", "lag", "lag-tz", "lb", "lb-lu", "lg", "lg-ug", "lkt", "lkt-us", "ln",
	// "ln-ao", "ln-cd", "ln-cf", "ln-cg", "lo", "lo-la", "lrc", "lrc-iq", "lrc-ir",
	// "lt", "lt-lt", "lu", "lu-cd", "luo", "luo-ke", "luy", "luy-ke", "lv", "lv-lv",
	// "mai", "mai-in", "mas", "mas-ke", "mas-tz", "mer", "mer-ke", "mfe", "mfe-mu",
	// "mg", "mg-mg", "mgh", "mgh-mz", "mgo", "mgo-cm", "mi", "mi-nz", "mk", "mk-mk",
	// "ml", "ml-in", "mn", "mn-mn", "mni", "mni-in", "mr", "mr-in", "ms", "ms-bn",
	// "ms-id", "ms-my", "ms-sg", "mt", "mt-mt", "mua", "mua-cm", "my", "my-mm", "mzn",
	// "mzn-ir", "naq", "naq-na", "nb", "nb-no", "nb-sj", "nd", "nd-zw", "nds",
	// "nds-de", "nds-nl", "ne", "ne-in", "ne-np", "nl", "nl-aw", "nl-be", "nl-bq",
	// "nl-ch", "nl-cw", "nl-lu", "nl-nl", "nl-sr", "nl-sx", "nmg", "nmg-cm", "nn",
	// "nn-no", "nnh", "nnh-cm", "no", "no-no", "nus", "nus-ss", "nyn", "nyn-ug", "om",
	// "om-et", "om-ke", "or", "or-in", "os", "os-ge", "os-ru", "pa", "pa-in", "pa-pk",
	// "pcm", "pcm-ng", "pl", "pl-pl", "prg", "prg-001", "ps", "ps-af", "ps-pk", "pt",
	// "pt-ao", "pt-br", "pt-ch", "pt-cv", "pt-gq", "pt-gw", "pt-lu", "pt-mo", "pt-mz",
	// "pt-pt", "pt-st", "pt-tl", "qu", "qu-bo", "qu-ec", "qu-pe", "rm", "rm-ch", "rn",
	// "rn-bi", "ro", "ro-md", "ro-ro", "rof", "rof-tz", "ru", "ru-by", "ru-kg",
	// "ru-kz", "ru-md", "ru-ru", "ru-ua", "rw", "rw-rw", "rwk", "rwk-tz", "sa",
	// "sa-in", "sah", "sah-ru", "saq", "saq-ke", "sat", "sat-in", "sbp", "sbp-tz",
	// "sd", "sd-in", "sd-pk", "se", "se-fi", "se-no", "se-se", "seh", "seh-mz", "ses",
	// "ses-ml", "sg", "sg-cf", "shi", "shi-ma", "si", "si-lk", "sk", "sk-sk", "sl",
	// "sl-si", "smn", "smn-fi", "sn", "sn-zw", "so", "so-dj", "so-et", "so-ke",
	// "so-so", "sq", "sq-al", "sq-mk", "sq-xk", "sr", "sr-ba", "sr-cs", "sr-me",
	// "sr-rs", "sr-xk", "su", "su-id", "sv", "sv-ax", "sv-fi", "sv-se", "sw", "sw-cd",
	// "sw-ke", "sw-tz", "sw-ug", "sy", "ta", "ta-in", "ta-lk", "ta-my", "ta-sg", "te",
	// "te-in", "teo", "teo-ke", "teo-ug", "tg", "tg-tj", "th", "th-th", "ti", "ti-er",
	// "ti-et", "tk", "tk-tm", "tl", "to", "to-to", "tr", "tr-cy", "tr-tr", "tt",
	// "tt-ru", "twq", "twq-ne", "tzm", "tzm-ma", "ug", "ug-cn", "uk", "uk-ua", "ur",
	// "ur-in", "ur-pk", "uz", "uz-af", "uz-uz", "vai", "vai-lr", "vi", "vi-vn", "vo",
	// "vo-001", "vun", "vun-tz", "wae", "wae-ch", "wo", "wo-sn", "xh", "xh-za", "xog",
	// "xog-ug", "yav", "yav-cm", "yi", "yi-001", "yo", "yo-bj", "yo-ng", "yue",
	// "yue-cn", "yue-hk", "zgh", "zgh-ma", "zh", "zh-cn", "zh-hans", "zh-hant",
	// "zh-hk", "zh-mo", "zh-sg", "zh-tw", "zu", "zu-za".
	Language       BlogPostLanguage         `json:"language,required"`
	LayoutSections map[string]LayoutSection `json:"layoutSections,required"`
	// Optional override to set the URL to be used in the rel=canonical link tag on the
	// page.
	LinkRelCanonicalURL string `json:"linkRelCanonicalUrl,required"`
	MabExperimentID     string `json:"mabExperimentId,required"`
	// A description that goes in <meta> tag on the page.
	MetaDescription string `json:"metaDescription,required"`
	// The internal name of the post.
	Name                  string `json:"name,required"`
	PageExpiryDate        int64  `json:"pageExpiryDate,required"`
	PageExpiryEnabled     bool   `json:"pageExpiryEnabled,required"`
	PageExpiryRedirectID  int64  `json:"pageExpiryRedirectId,required"`
	PageExpiryRedirectURL string `json:"pageExpiryRedirectUrl,required"`
	// Set this to create a password protected page. Entering the password will be
	// required to view the page.
	Password string `json:"password,required"`
	// The HTML of the main post body.
	PostBody string `json:"postBody,required"`
	// The summary of the blog post that will appear on the main listing page.
	PostSummary string `json:"postSummary,required"`
	// Rules for require member registration to access private content.
	PublicAccessRules []PublicAccessRule `json:"publicAccessRules,required"`
	// Boolean to determine whether or not to respect publicAccessRules.
	PublicAccessRulesEnabled bool `json:"publicAccessRulesEnabled,required"`
	// The date (ISO8601 format) the blog post is to be published at.
	PublishDate time.Time `json:"publishDate,required" format:"date-time"`
	// Set this to true if you want to be published immediately when the schedule
	// publish endpoint is called, and to ignore the publish_date setting.
	PublishImmediately bool `json:"publishImmediately,required"`
	// The contents of the RSS body for this Blog Post.
	RssBody string `json:"rssBody,required"`
	// The contents of the RSS summary for this Blog Post.
	RssSummary string `json:"rssSummary,required"`
	// The URL slug of the blog post. This field is appended to the domain to construct
	// the url of this post.
	Slug string `json:"slug,required"`
	// An enumeration describing the current publish state of the post.
	State string `json:"state,required"`
	// The IDs of the tags associated with this post.
	TagIDs              []int64        `json:"tagIds,required"`
	ThemeSettingsValues map[string]any `json:"themeSettingsValues,required"`
	// ID of the primary blog post that this post was translated from.
	TranslatedFromID string                                   `json:"translatedFromId,required"`
	Translations     map[string]PagesContentLanguageVariation `json:"translations,required"`
	Updated          time.Time                                `json:"updated,required" format:"date-time"`
	// The ID of the user that updated the post.
	UpdatedByID string `json:"updatedById,required"`
	// A generated field representing the URL of this blog post.
	URL string `json:"url,required"`
	// Boolean to determine if this post should use a featured image.
	UseFeaturedImage bool `json:"useFeaturedImage,required"`
	// A data structure containing the data for all the modules inside the containers
	// for this post. This will only be populated if the page has widget containers.
	WidgetContainers map[string]any `json:"widgetContainers,required"`
	// A data structure containing the data for all the modules for this page.
	Widgets map[string]any `json:"widgets,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                            respjson.Field
		AbStatus                      respjson.Field
		AbTestID                      respjson.Field
		ArchivedAt                    respjson.Field
		ArchivedInDashboard           respjson.Field
		AttachedStylesheets           respjson.Field
		AuthorName                    respjson.Field
		BlogAuthorID                  respjson.Field
		Campaign                      respjson.Field
		CategoryID                    respjson.Field
		ContentGroupID                respjson.Field
		ContentTypeCategory           respjson.Field
		Created                       respjson.Field
		CreatedByID                   respjson.Field
		CurrentlyPublished            respjson.Field
		CurrentState                  respjson.Field
		Domain                        respjson.Field
		DynamicPageDataSourceID       respjson.Field
		DynamicPageDataSourceType     respjson.Field
		DynamicPageHubDBTableID       respjson.Field
		EnableDomainStylesheets       respjson.Field
		EnableGoogleAmpOutputOverride respjson.Field
		EnableLayoutStylesheets       respjson.Field
		FeaturedImage                 respjson.Field
		FeaturedImageAltText          respjson.Field
		FolderID                      respjson.Field
		FooterHTML                    respjson.Field
		HeadHTML                      respjson.Field
		HTMLTitle                     respjson.Field
		IncludeDefaultCustomCss       respjson.Field
		Language                      respjson.Field
		LayoutSections                respjson.Field
		LinkRelCanonicalURL           respjson.Field
		MabExperimentID               respjson.Field
		MetaDescription               respjson.Field
		Name                          respjson.Field
		PageExpiryDate                respjson.Field
		PageExpiryEnabled             respjson.Field
		PageExpiryRedirectID          respjson.Field
		PageExpiryRedirectURL         respjson.Field
		Password                      respjson.Field
		PostBody                      respjson.Field
		PostSummary                   respjson.Field
		PublicAccessRules             respjson.Field
		PublicAccessRulesEnabled      respjson.Field
		PublishDate                   respjson.Field
		PublishImmediately            respjson.Field
		RssBody                       respjson.Field
		RssSummary                    respjson.Field
		Slug                          respjson.Field
		State                         respjson.Field
		TagIDs                        respjson.Field
		ThemeSettingsValues           respjson.Field
		TranslatedFromID              respjson.Field
		Translations                  respjson.Field
		Updated                       respjson.Field
		UpdatedByID                   respjson.Field
		URL                           respjson.Field
		UseFeaturedImage              respjson.Field
		WidgetContainers              respjson.Field
		Widgets                       respjson.Field
		ExtraFields                   map[string]respjson.Field
		raw                           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BlogPost) RawJSON() string { return r.JSON.raw }
func (r *BlogPost) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this BlogPost to a BlogPostParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// BlogPostParam.Overrides()
func (r BlogPost) ToParam() BlogPostParam {
	return param.Override[BlogPostParam](json.RawMessage(r.RawJSON()))
}

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
	BlogPostContentTypeCategory2  BlogPostContentTypeCategory = "2"
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
	BlogPostLanguageAf     BlogPostLanguage = "af"
	BlogPostLanguageAfNa   BlogPostLanguage = "af-na"
	BlogPostLanguageAfZa   BlogPostLanguage = "af-za"
	BlogPostLanguageAgq    BlogPostLanguage = "agq"
	BlogPostLanguageAgqCm  BlogPostLanguage = "agq-cm"
	BlogPostLanguageAk     BlogPostLanguage = "ak"
	BlogPostLanguageAkGh   BlogPostLanguage = "ak-gh"
	BlogPostLanguageAm     BlogPostLanguage = "am"
	BlogPostLanguageAmEt   BlogPostLanguage = "am-et"
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
	BlogPostLanguageAz     BlogPostLanguage = "az"
	BlogPostLanguageAzAz   BlogPostLanguage = "az-az"
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
	BlogPostLanguageChr    BlogPostLanguage = "chr"
	BlogPostLanguageChrUs  BlogPostLanguage = "chr-us"
	BlogPostLanguageCkb    BlogPostLanguage = "ckb"
	BlogPostLanguageCkbIq  BlogPostLanguage = "ckb-iq"
	BlogPostLanguageCkbIr  BlogPostLanguage = "ckb-ir"
	BlogPostLanguageCs     BlogPostLanguage = "cs"
	BlogPostLanguageCsCz   BlogPostLanguage = "cs-cz"
	BlogPostLanguageCu     BlogPostLanguage = "cu"
	BlogPostLanguageCuRu   BlogPostLanguage = "cu-ru"
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
	BlogPostLanguageEnEr   BlogPostLanguage = "en-er"
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
	BlogPostLanguageEnTk   BlogPostLanguage = "en-tk"
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
	BlogPostLanguageHr     BlogPostLanguage = "hr"
	BlogPostLanguageHrBa   BlogPostLanguage = "hr-ba"
	BlogPostLanguageHrHr   BlogPostLanguage = "hr-hr"
	BlogPostLanguageHsb    BlogPostLanguage = "hsb"
	BlogPostLanguageHsbDe  BlogPostLanguage = "hsb-de"
	BlogPostLanguageHu     BlogPostLanguage = "hu"
	BlogPostLanguageHuHu   BlogPostLanguage = "hu-hu"
	BlogPostLanguageHy     BlogPostLanguage = "hy"
	BlogPostLanguageHyAm   BlogPostLanguage = "hy-am"
	BlogPostLanguageIa     BlogPostLanguage = "ia"
	BlogPostLanguageIa001  BlogPostLanguage = "ia-001"
	BlogPostLanguageID     BlogPostLanguage = "id"
	BlogPostLanguageIDID   BlogPostLanguage = "id-id"
	BlogPostLanguageIg     BlogPostLanguage = "ig"
	BlogPostLanguageIgNg   BlogPostLanguage = "ig-ng"
	BlogPostLanguageIi     BlogPostLanguage = "ii"
	BlogPostLanguageIiCn   BlogPostLanguage = "ii-cn"
	BlogPostLanguageIs     BlogPostLanguage = "is"
	BlogPostLanguageIsIs   BlogPostLanguage = "is-is"
	BlogPostLanguageIt     BlogPostLanguage = "it"
	BlogPostLanguageItCh   BlogPostLanguage = "it-ch"
	BlogPostLanguageItIt   BlogPostLanguage = "it-it"
	BlogPostLanguageItSm   BlogPostLanguage = "it-sm"
	BlogPostLanguageItVa   BlogPostLanguage = "it-va"
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
	BlogPostLanguageKde    BlogPostLanguage = "kde"
	BlogPostLanguageKdeTz  BlogPostLanguage = "kde-tz"
	BlogPostLanguageKea    BlogPostLanguage = "kea"
	BlogPostLanguageKeaCv  BlogPostLanguage = "kea-cv"
	BlogPostLanguageKhq    BlogPostLanguage = "khq"
	BlogPostLanguageKhqMl  BlogPostLanguage = "khq-ml"
	BlogPostLanguageKi     BlogPostLanguage = "ki"
	BlogPostLanguageKiKe   BlogPostLanguage = "ki-ke"
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
	BlogPostLanguageKw     BlogPostLanguage = "kw"
	BlogPostLanguageKwGB   BlogPostLanguage = "kw-gb"
	BlogPostLanguageKy     BlogPostLanguage = "ky"
	BlogPostLanguageKyKg   BlogPostLanguage = "ky-kg"
	BlogPostLanguageLag    BlogPostLanguage = "lag"
	BlogPostLanguageLagTz  BlogPostLanguage = "lag-tz"
	BlogPostLanguageLb     BlogPostLanguage = "lb"
	BlogPostLanguageLbLu   BlogPostLanguage = "lb-lu"
	BlogPostLanguageLg     BlogPostLanguage = "lg"
	BlogPostLanguageLgUg   BlogPostLanguage = "lg-ug"
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
	BlogPostLanguageNus    BlogPostLanguage = "nus"
	BlogPostLanguageNusSS  BlogPostLanguage = "nus-ss"
	BlogPostLanguageNyn    BlogPostLanguage = "nyn"
	BlogPostLanguageNynUg  BlogPostLanguage = "nyn-ug"
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
	BlogPostLanguageSmn    BlogPostLanguage = "smn"
	BlogPostLanguageSmnFi  BlogPostLanguage = "smn-fi"
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
	BlogPostLanguageTo     BlogPostLanguage = "to"
	BlogPostLanguageToTo   BlogPostLanguage = "to-to"
	BlogPostLanguageTr     BlogPostLanguage = "tr"
	BlogPostLanguageTrCy   BlogPostLanguage = "tr-cy"
	BlogPostLanguageTrTr   BlogPostLanguage = "tr-tr"
	BlogPostLanguageTt     BlogPostLanguage = "tt"
	BlogPostLanguageTtRu   BlogPostLanguage = "tt-ru"
	BlogPostLanguageTwq    BlogPostLanguage = "twq"
	BlogPostLanguageTwqNe  BlogPostLanguage = "twq-ne"
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
	BlogPostLanguageVi     BlogPostLanguage = "vi"
	BlogPostLanguageViVn   BlogPostLanguage = "vi-vn"
	BlogPostLanguageVo     BlogPostLanguage = "vo"
	BlogPostLanguageVo001  BlogPostLanguage = "vo-001"
	BlogPostLanguageVun    BlogPostLanguage = "vun"
	BlogPostLanguageVunTz  BlogPostLanguage = "vun-tz"
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
	BlogPostLanguageYue    BlogPostLanguage = "yue"
	BlogPostLanguageYueCn  BlogPostLanguage = "yue-cn"
	BlogPostLanguageYueHk  BlogPostLanguage = "yue-hk"
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

// Model definition for a Blog Post.
//
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
	ID string `json:"id,required"`
	// Any of "automated_loser_variant", "automated_master", "automated_variant",
	// "loser_variant", "mab_master", "mab_variant", "master", "variant".
	AbStatus BlogPostAbStatus `json:"abStatus,omitzero,required"`
	AbTestID string           `json:"abTestId,required"`
	// The timestamp (ISO8601 format) when this Blog Post was deleted.
	ArchivedAt int64 `json:"archivedAt,required"`
	// If True, the post will not show up in your dashboard, although the post could
	// still be live.
	ArchivedInDashboard bool `json:"archivedInDashboard,required"`
	// List of stylesheets to attach to this blog post. These stylesheets are attached
	// to just this page. Order of precedence is bottom to top, just like in the HTML.
	AttachedStylesheets []map[string]any `json:"attachedStylesheets,omitzero,required"`
	// The name of the blog author associated with the post.
	AuthorName string `json:"authorName,required"`
	// The ID of the blog author associated with this post.
	BlogAuthorID string `json:"blogAuthorId,required"`
	// The GUID of the marketing campaign the post is associated with.
	Campaign string `json:"campaign,required"`
	// ID of the object type.
	CategoryID int64 `json:"categoryId,required"`
	// The ID of the post's parent blog.
	ContentGroupID string `json:"contentGroupId,required"`
	// An ENUM descibing the type of this object. Should always be BLOG_POST.
	//
	// Any of "0", "1", "10", "11", "12", "13", "14", "15", "2", "3", "4", "5", "6",
	// "7", "8", "9".
	ContentTypeCategory BlogPostContentTypeCategory `json:"contentTypeCategory,omitzero,required"`
	Created             time.Time                   `json:"created,required" format:"date-time"`
	// The ID of the user that created the post.
	CreatedByID        string `json:"createdById,required"`
	CurrentlyPublished bool   `json:"currentlyPublished,required"`
	// A generated ENUM descibing the current state of this Blog Post. Should always
	// match state.
	//
	// Any of "AUTOMATED", "AUTOMATED_AB", "AUTOMATED_AB_VARIANT", "AUTOMATED_DRAFT",
	// "AUTOMATED_DRAFT_AB", "AUTOMATED_DRAFT_ABVARIANT", "AUTOMATED_FOR_FORM",
	// "AUTOMATED_FOR_FORM_BUFFER", "AUTOMATED_FOR_FORM_DRAFT",
	// "AUTOMATED_FOR_FORM_LEGACY", "AUTOMATED_LOSER_ABVARIANT", "AUTOMATED_SENDING",
	// "BLOG_EMAIL_DRAFT", "BLOG_EMAIL_PUBLISHED", "DRAFT", "DRAFT_AB",
	// "DRAFT_AB_VARIANT", "ERROR", "LOSER_AB_VARIANT", "PAGE_STUB", "PRE_PROCESSING",
	// "PROCESSING", "PUBLISHED", "PUBLISHED_AB", "PUBLISHED_AB_VARIANT",
	// "PUBLISHED_OR_SCHEDULED", "RSS_TO_EMAIL_DRAFT", "RSS_TO_EMAIL_PUBLISHED",
	// "SCHEDULED", "SCHEDULED_AB", "SCHEDULED_OR_PUBLISHED".
	CurrentState BlogPostCurrentState `json:"currentState,omitzero,required"`
	// The domain that the post lives on. If null, the post will default to the domain
	// of the parent blog.
	Domain                    string `json:"domain,required"`
	DynamicPageDataSourceID   string `json:"dynamicPageDataSourceId,required"`
	DynamicPageDataSourceType int64  `json:"dynamicPageDataSourceType,required"`
	// For dynamic HubDB pages, the ID of the HubDB table this post references.
	DynamicPageHubDBTableID string `json:"dynamicPageHubDbTableId,required"`
	// Boolean to determine whether or not the styles from the template should be
	// applied.
	EnableDomainStylesheets bool `json:"enableDomainStylesheets,required"`
	// Boolean to allow overriding the AMP settings for the blog.
	EnableGoogleAmpOutputOverride bool `json:"enableGoogleAmpOutputOverride,required"`
	// Boolean to determine whether or not the styles from the template should be
	// applied.
	EnableLayoutStylesheets bool `json:"enableLayoutStylesheets,required"`
	// The featuredImage of this Blog Post.
	FeaturedImage string `json:"featuredImage,required"`
	// Alt Text of the featuredImage.
	FeaturedImageAltText string `json:"featuredImageAltText,required"`
	FolderID             string `json:"folderId,required"`
	// Custom HTML for embed codes, javascript that should be placed before the </body>
	// tag of the page.
	FooterHTML string `json:"footerHtml,required"`
	// Custom HTML for embed codes, javascript, etc. that goes in the <head> tag of the
	// page.
	HeadHTML string `json:"headHtml,required"`
	// The HTML title of the post.
	HTMLTitle string `json:"htmlTitle,required"`
	// Boolean to determine whether or not the Primary CSS Files should be applied.
	IncludeDefaultCustomCss bool `json:"includeDefaultCustomCss,required"`
	// The explicitly defined ISO 639 language code of the post. If null, the post will
	// default to the language of the parent blog.
	//
	// Any of "af", "af-na", "af-za", "agq", "agq-cm", "ak", "ak-gh", "am", "am-et",
	// "ar", "ar-001", "ar-ae", "ar-bh", "ar-dj", "ar-dz", "ar-eg", "ar-eh", "ar-er",
	// "ar-il", "ar-iq", "ar-jo", "ar-km", "ar-kw", "ar-lb", "ar-ly", "ar-ma", "ar-mr",
	// "ar-om", "ar-ps", "ar-qa", "ar-sa", "ar-sd", "ar-so", "ar-ss", "ar-sy", "ar-td",
	// "ar-tn", "ar-ye", "as", "as-in", "asa", "asa-tz", "ast", "ast-es", "az",
	// "az-az", "bas", "bas-cm", "be", "be-by", "bem", "bem-zm", "bez", "bez-tz", "bg",
	// "bg-bg", "bm", "bm-ml", "bn", "bn-bd", "bn-in", "bo", "bo-cn", "bo-in", "br",
	// "br-fr", "brx", "brx-in", "bs", "bs-ba", "ca", "ca-ad", "ca-es", "ca-fr",
	// "ca-it", "ccp", "ccp-bd", "ccp-in", "ce", "ce-ru", "ceb", "ceb-ph", "cgg",
	// "cgg-ug", "chr", "chr-us", "ckb", "ckb-iq", "ckb-ir", "cs", "cs-cz", "cu",
	// "cu-ru", "cy", "cy-gb", "da", "da-dk", "da-gl", "dav", "dav-ke", "de", "de-at",
	// "de-be", "de-ch", "de-de", "de-gr", "de-it", "de-li", "de-lu", "dje", "dje-ne",
	// "doi", "doi-in", "dsb", "dsb-de", "dua", "dua-cm", "dyo", "dyo-sn", "dz",
	// "dz-bt", "ebu", "ebu-ke", "ee", "ee-gh", "ee-tg", "el", "el-cy", "el-gr", "en",
	// "en-001", "en-150", "en-ae", "en-ag", "en-ai", "en-as", "en-at", "en-au",
	// "en-bb", "en-be", "en-bi", "en-bm", "en-bs", "en-bw", "en-bz", "en-ca", "en-cc",
	// "en-ch", "en-ck", "en-cm", "en-cn", "en-cx", "en-cy", "en-de", "en-dg", "en-dk",
	// "en-dm", "en-ee", "en-er", "en-fi", "en-fj", "en-fk", "en-fm", "en-fr", "en-gb",
	// "en-gd", "en-gg", "en-gh", "en-gi", "en-gm", "en-gu", "en-gy", "en-hk", "en-ie",
	// "en-il", "en-im", "en-in", "en-io", "en-je", "en-jm", "en-ke", "en-ki", "en-kn",
	// "en-ky", "en-lc", "en-lr", "en-ls", "en-lu", "en-mg", "en-mh", "en-mo", "en-mp",
	// "en-ms", "en-mt", "en-mu", "en-mw", "en-mx", "en-my", "en-na", "en-nf", "en-ng",
	// "en-nl", "en-nr", "en-nu", "en-nz", "en-pg", "en-ph", "en-pk", "en-pn", "en-pr",
	// "en-pw", "en-rw", "en-sb", "en-sc", "en-sd", "en-se", "en-sg", "en-sh", "en-si",
	// "en-sl", "en-ss", "en-sx", "en-sz", "en-tc", "en-tk", "en-to", "en-tt", "en-tv",
	// "en-tz", "en-ug", "en-um", "en-us", "en-vc", "en-vg", "en-vi", "en-vu", "en-ws",
	// "en-za", "en-zm", "en-zw", "eo", "eo-001", "es", "es-419", "es-ar", "es-bo",
	// "es-br", "es-bz", "es-cl", "es-co", "es-cr", "es-cu", "es-do", "es-ea", "es-ec",
	// "es-es", "es-gq", "es-gt", "es-hn", "es-ic", "es-mx", "es-ni", "es-pa", "es-pe",
	// "es-ph", "es-pr", "es-py", "es-sv", "es-us", "es-uy", "es-ve", "et", "et-ee",
	// "eu", "eu-es", "ewo", "ewo-cm", "fa", "fa-af", "fa-ir", "ff", "ff-bf", "ff-cm",
	// "ff-gh", "ff-gm", "ff-gn", "ff-gw", "ff-lr", "ff-mr", "ff-ne", "ff-ng", "ff-sl",
	// "ff-sn", "fi", "fi-fi", "fil", "fil-ph", "fo", "fo-dk", "fo-fo", "fr", "fr-be",
	// "fr-bf", "fr-bi", "fr-bj", "fr-bl", "fr-ca", "fr-cd", "fr-cf", "fr-cg", "fr-ch",
	// "fr-ci", "fr-cm", "fr-dj", "fr-dz", "fr-fr", "fr-ga", "fr-gf", "fr-gn", "fr-gp",
	// "fr-gq", "fr-ht", "fr-km", "fr-lu", "fr-ma", "fr-mc", "fr-mf", "fr-mg", "fr-ml",
	// "fr-mq", "fr-mr", "fr-mu", "fr-nc", "fr-ne", "fr-pf", "fr-pm", "fr-re", "fr-rw",
	// "fr-sc", "fr-sn", "fr-sy", "fr-td", "fr-tg", "fr-tn", "fr-vu", "fr-wf", "fr-yt",
	// "fur", "fur-it", "fy", "fy-nl", "ga", "ga-gb", "ga-ie", "gd", "gd-gb", "gl",
	// "gl-es", "gsw", "gsw-ch", "gsw-fr", "gsw-li", "gu", "gu-in", "guz", "guz-ke",
	// "gv", "gv-im", "ha", "ha-gh", "ha-ne", "ha-ng", "haw", "haw-us", "he", "he-il",
	// "hi", "hi-in", "hr", "hr-ba", "hr-hr", "hsb", "hsb-de", "hu", "hu-hu", "hy",
	// "hy-am", "ia", "ia-001", "id", "id-id", "ig", "ig-ng", "ii", "ii-cn", "is",
	// "is-is", "it", "it-ch", "it-it", "it-sm", "it-va", "ja", "ja-jp", "jgo",
	// "jgo-cm", "jmc", "jmc-tz", "jv", "jv-id", "ka", "ka-ge", "kab", "kab-dz", "kam",
	// "kam-ke", "kde", "kde-tz", "kea", "kea-cv", "khq", "khq-ml", "ki", "ki-ke",
	// "kk", "kk-kz", "kkj", "kkj-cm", "kl", "kl-gl", "kln", "kln-ke", "km", "km-kh",
	// "kn", "kn-in", "ko", "ko-kp", "ko-kr", "kok", "kok-in", "ks", "ks-in", "ksb",
	// "ksb-tz", "ksf", "ksf-cm", "ksh", "ksh-de", "ku", "ku-tr", "kw", "kw-gb", "ky",
	// "ky-kg", "lag", "lag-tz", "lb", "lb-lu", "lg", "lg-ug", "lkt", "lkt-us", "ln",
	// "ln-ao", "ln-cd", "ln-cf", "ln-cg", "lo", "lo-la", "lrc", "lrc-iq", "lrc-ir",
	// "lt", "lt-lt", "lu", "lu-cd", "luo", "luo-ke", "luy", "luy-ke", "lv", "lv-lv",
	// "mai", "mai-in", "mas", "mas-ke", "mas-tz", "mer", "mer-ke", "mfe", "mfe-mu",
	// "mg", "mg-mg", "mgh", "mgh-mz", "mgo", "mgo-cm", "mi", "mi-nz", "mk", "mk-mk",
	// "ml", "ml-in", "mn", "mn-mn", "mni", "mni-in", "mr", "mr-in", "ms", "ms-bn",
	// "ms-id", "ms-my", "ms-sg", "mt", "mt-mt", "mua", "mua-cm", "my", "my-mm", "mzn",
	// "mzn-ir", "naq", "naq-na", "nb", "nb-no", "nb-sj", "nd", "nd-zw", "nds",
	// "nds-de", "nds-nl", "ne", "ne-in", "ne-np", "nl", "nl-aw", "nl-be", "nl-bq",
	// "nl-ch", "nl-cw", "nl-lu", "nl-nl", "nl-sr", "nl-sx", "nmg", "nmg-cm", "nn",
	// "nn-no", "nnh", "nnh-cm", "no", "no-no", "nus", "nus-ss", "nyn", "nyn-ug", "om",
	// "om-et", "om-ke", "or", "or-in", "os", "os-ge", "os-ru", "pa", "pa-in", "pa-pk",
	// "pcm", "pcm-ng", "pl", "pl-pl", "prg", "prg-001", "ps", "ps-af", "ps-pk", "pt",
	// "pt-ao", "pt-br", "pt-ch", "pt-cv", "pt-gq", "pt-gw", "pt-lu", "pt-mo", "pt-mz",
	// "pt-pt", "pt-st", "pt-tl", "qu", "qu-bo", "qu-ec", "qu-pe", "rm", "rm-ch", "rn",
	// "rn-bi", "ro", "ro-md", "ro-ro", "rof", "rof-tz", "ru", "ru-by", "ru-kg",
	// "ru-kz", "ru-md", "ru-ru", "ru-ua", "rw", "rw-rw", "rwk", "rwk-tz", "sa",
	// "sa-in", "sah", "sah-ru", "saq", "saq-ke", "sat", "sat-in", "sbp", "sbp-tz",
	// "sd", "sd-in", "sd-pk", "se", "se-fi", "se-no", "se-se", "seh", "seh-mz", "ses",
	// "ses-ml", "sg", "sg-cf", "shi", "shi-ma", "si", "si-lk", "sk", "sk-sk", "sl",
	// "sl-si", "smn", "smn-fi", "sn", "sn-zw", "so", "so-dj", "so-et", "so-ke",
	// "so-so", "sq", "sq-al", "sq-mk", "sq-xk", "sr", "sr-ba", "sr-cs", "sr-me",
	// "sr-rs", "sr-xk", "su", "su-id", "sv", "sv-ax", "sv-fi", "sv-se", "sw", "sw-cd",
	// "sw-ke", "sw-tz", "sw-ug", "sy", "ta", "ta-in", "ta-lk", "ta-my", "ta-sg", "te",
	// "te-in", "teo", "teo-ke", "teo-ug", "tg", "tg-tj", "th", "th-th", "ti", "ti-er",
	// "ti-et", "tk", "tk-tm", "tl", "to", "to-to", "tr", "tr-cy", "tr-tr", "tt",
	// "tt-ru", "twq", "twq-ne", "tzm", "tzm-ma", "ug", "ug-cn", "uk", "uk-ua", "ur",
	// "ur-in", "ur-pk", "uz", "uz-af", "uz-uz", "vai", "vai-lr", "vi", "vi-vn", "vo",
	// "vo-001", "vun", "vun-tz", "wae", "wae-ch", "wo", "wo-sn", "xh", "xh-za", "xog",
	// "xog-ug", "yav", "yav-cm", "yi", "yi-001", "yo", "yo-bj", "yo-ng", "yue",
	// "yue-cn", "yue-hk", "zgh", "zgh-ma", "zh", "zh-cn", "zh-hans", "zh-hant",
	// "zh-hk", "zh-mo", "zh-sg", "zh-tw", "zu", "zu-za".
	Language       BlogPostLanguage              `json:"language,omitzero,required"`
	LayoutSections map[string]LayoutSectionParam `json:"layoutSections,omitzero,required"`
	// Optional override to set the URL to be used in the rel=canonical link tag on the
	// page.
	LinkRelCanonicalURL string `json:"linkRelCanonicalUrl,required"`
	MabExperimentID     string `json:"mabExperimentId,required"`
	// A description that goes in <meta> tag on the page.
	MetaDescription string `json:"metaDescription,required"`
	// The internal name of the post.
	Name                  string `json:"name,required"`
	PageExpiryDate        int64  `json:"pageExpiryDate,required"`
	PageExpiryEnabled     bool   `json:"pageExpiryEnabled,required"`
	PageExpiryRedirectID  int64  `json:"pageExpiryRedirectId,required"`
	PageExpiryRedirectURL string `json:"pageExpiryRedirectUrl,required"`
	// Set this to create a password protected page. Entering the password will be
	// required to view the page.
	Password string `json:"password,required"`
	// The HTML of the main post body.
	PostBody string `json:"postBody,required"`
	// The summary of the blog post that will appear on the main listing page.
	PostSummary string `json:"postSummary,required"`
	// Rules for require member registration to access private content.
	PublicAccessRules []PublicAccessRule `json:"publicAccessRules,omitzero,required"`
	// Boolean to determine whether or not to respect publicAccessRules.
	PublicAccessRulesEnabled bool `json:"publicAccessRulesEnabled,required"`
	// The date (ISO8601 format) the blog post is to be published at.
	PublishDate time.Time `json:"publishDate,required" format:"date-time"`
	// Set this to true if you want to be published immediately when the schedule
	// publish endpoint is called, and to ignore the publish_date setting.
	PublishImmediately bool `json:"publishImmediately,required"`
	// The contents of the RSS body for this Blog Post.
	RssBody string `json:"rssBody,required"`
	// The contents of the RSS summary for this Blog Post.
	RssSummary string `json:"rssSummary,required"`
	// The URL slug of the blog post. This field is appended to the domain to construct
	// the url of this post.
	Slug string `json:"slug,required"`
	// An enumeration describing the current publish state of the post.
	State string `json:"state,required"`
	// The IDs of the tags associated with this post.
	TagIDs              []int64        `json:"tagIds,omitzero,required"`
	ThemeSettingsValues map[string]any `json:"themeSettingsValues,omitzero,required"`
	// ID of the primary blog post that this post was translated from.
	TranslatedFromID string                                        `json:"translatedFromId,required"`
	Translations     map[string]PagesContentLanguageVariationParam `json:"translations,omitzero,required"`
	Updated          time.Time                                     `json:"updated,required" format:"date-time"`
	// The ID of the user that updated the post.
	UpdatedByID string `json:"updatedById,required"`
	// A generated field representing the URL of this blog post.
	URL string `json:"url,required"`
	// Boolean to determine if this post should use a featured image.
	UseFeaturedImage bool `json:"useFeaturedImage,required"`
	// A data structure containing the data for all the modules inside the containers
	// for this post. This will only be populated if the page has widget containers.
	WidgetContainers map[string]any `json:"widgetContainers,omitzero,required"`
	// A data structure containing the data for all the modules for this page.
	Widgets map[string]any `json:"widgets,omitzero,required"`
	paramObj
}

func (r BlogPostParam) MarshalJSON() (data []byte, err error) {
	type shadow BlogPostParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BlogPostParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Request body object for creating new blog post language variant.
//
// The property ID is required.
type BlogPostLanguageCloneRequestVNextParam struct {
	// ID of blog post to clone.
	ID string `json:"id,required"`
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

// Response object for collections of blog posts with pagination information.
type CollectionResponseWithTotalBlogPostForwardPaging struct {
	// Collection of blog posts.
	Results []BlogPost `json:"results,required"`
	// Total number of blog posts.
	Total  int64                `json:"total,required"`
	Paging shared.ForwardPaging `json:"paging"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Results     respjson.Field
		Total       respjson.Field
		Paging      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CollectionResponseWithTotalBlogPostForwardPaging) RawJSON() string { return r.JSON.raw }
func (r *CollectionResponseWithTotalBlogPostForwardPaging) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Response object for collections of blog post versions with pagination
// information.
type CollectionResponseWithTotalVersionBlogPost struct {
	// Collection of blog post versions.
	Results []VersionBlogPost `json:"results,required"`
	// Total number of blog post versions.
	Total  int64         `json:"total,required"`
	Paging shared.Paging `json:"paging"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Results     respjson.Field
		Total       respjson.Field
		Paging      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CollectionResponseWithTotalVersionBlogPost) RawJSON() string { return r.JSON.raw }
func (r *CollectionResponseWithTotalVersionBlogPost) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Margin = any

type Padding = any

// Model definition of a version of a blog post.
type VersionBlogPost struct {
	// The id of the version.
	ID string `json:"id,required"`
	// Model definition for a Blog Post.
	Object    BlogPost  `json:"object,required"`
	UpdatedAt time.Time `json:"updatedAt,required" format:"date-time"`
	// Model definition for a version user. Contains addition information about the
	// user who created a version.
	User shared.VersionUser `json:"user,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Object      respjson.Field
		UpdatedAt   respjson.Field
		User        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VersionBlogPost) RawJSON() string { return r.JSON.raw }
func (r *VersionBlogPost) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BlogPostNewParams struct {
	// Model definition for a Blog Post.
	BlogPost BlogPostParam
	paramObj
}

func (r BlogPostNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.BlogPost)
}
func (r *BlogPostNewParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.BlogPost)
}

type BlogPostUpdateParams struct {
	// Model definition for a Blog Post.
	BlogPost BlogPostParam
	// Specifies whether to update deleted blog posts. Defaults to `false`.
	Archived param.Opt[bool] `query:"archived,omitzero" json:"-"`
	paramObj
}

func (r BlogPostUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.BlogPost)
}
func (r *BlogPostUpdateParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.BlogPost)
}

// URLQuery serializes [BlogPostUpdateParams]'s query parameters as `url.Values`.
func (r BlogPostUpdateParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type BlogPostListParams struct {
	// The cursor token value to get the next set of results. You can get this from the
	// `paging.next.after` JSON property of a paged response containing more results.
	After param.Opt[string] `query:"after,omitzero" json:"-"`
	// Specifies whether to return deleted blog posts. Defaults to `false`.
	Archived param.Opt[bool] `query:"archived,omitzero" json:"-"`
	// Only return blog posts created after the specified time.
	CreatedAfter param.Opt[time.Time] `query:"createdAfter,omitzero" format:"date-time" json:"-"`
	// Only return blog posts created at exactly the specified time.
	CreatedAt param.Opt[time.Time] `query:"createdAt,omitzero" format:"date-time" json:"-"`
	// Only return blog posts created before the specified time.
	CreatedBefore param.Opt[time.Time] `query:"createdBefore,omitzero" format:"date-time" json:"-"`
	// The maximum number of results to return. Default is 20.
	Limit    param.Opt[int64]  `query:"limit,omitzero" json:"-"`
	Property param.Opt[string] `query:"property,omitzero" json:"-"`
	// Only return blog posts last updated after the specified time.
	UpdatedAfter param.Opt[time.Time] `query:"updatedAfter,omitzero" format:"date-time" json:"-"`
	// Only return blog posts last updated at exactly the specified time.
	UpdatedAt param.Opt[time.Time] `query:"updatedAt,omitzero" format:"date-time" json:"-"`
	// Only return blog posts last updated before the specified time.
	UpdatedBefore param.Opt[time.Time] `query:"updatedBefore,omitzero" format:"date-time" json:"-"`
	// Specifies which fields to use for sorting results. Valid fields are `createdAt`
	// (default), `name`, `updatedAt`, `createdBy`, `updatedBy`.
	Sort []string `query:"sort,omitzero" json:"-"`
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
	// Whether to return only results that have been deleted.
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
	// Request body object for attaching objects to multi-language groups.
	AttachToLangPrimaryRequestVNext AttachToLangPrimaryRequestVNextParam
	paramObj
}

func (r BlogPostAttachToLangGroupParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.AttachToLangPrimaryRequestVNext)
}
func (r *BlogPostAttachToLangGroupParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.AttachToLangPrimaryRequestVNext)
}

type BlogPostCloneParams struct {
	// Request body object for cloning content.
	ContentCloneRequestVNext ContentCloneRequestVNextParam
	paramObj
}

func (r BlogPostCloneParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.ContentCloneRequestVNext)
}
func (r *BlogPostCloneParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.ContentCloneRequestVNext)
}

type BlogPostNewLangVariationParams struct {
	// Request body object for creating new blog post language variant.
	BlogPostLanguageCloneRequestVNext BlogPostLanguageCloneRequestVNextParam
	paramObj
}

func (r BlogPostNewLangVariationParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.BlogPostLanguageCloneRequestVNext)
}
func (r *BlogPostNewLangVariationParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.BlogPostLanguageCloneRequestVNext)
}

type BlogPostDetachFromLangGroupParams struct {
	// Request body object for detaching objects from multi-language groups.
	DetachFromLangGroupRequestVNext DetachFromLangGroupRequestVNextParam
	paramObj
}

func (r BlogPostDetachFromLangGroupParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.DetachFromLangGroupRequestVNext)
}
func (r *BlogPostDetachFromLangGroupParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.DetachFromLangGroupRequestVNext)
}

type BlogPostGetParams struct {
	// Specifies whether to return deleted blog posts. Defaults to `false`.
	Archived param.Opt[bool] `query:"archived,omitzero" json:"-"`
	// Specific properties to return.
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
	ObjectID string `path:"objectId,required" json:"-"`
	paramObj
}

type BlogPostGetPreviousVersionsParams struct {
	// The cursor token value to get the next set of results. You can get this from the
	// `paging.next.after` JSON property of a paged response containing more results.
	After  param.Opt[string] `query:"after,omitzero" json:"-"`
	Before param.Opt[string] `query:"before,omitzero" json:"-"`
	// The maximum number of results to return. Default is 100.
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
	ObjectID string `path:"objectId,required" json:"-"`
	paramObj
}

type BlogPostRestorePreviousVersionToDraftParams struct {
	ObjectID string `path:"objectId,required" json:"-"`
	paramObj
}

type BlogPostScheduleParams struct {
	// Request body object for scheduling the publish of content
	ContentScheduleRequestVNext ContentScheduleRequestVNextParam
	paramObj
}

func (r BlogPostScheduleParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.ContentScheduleRequestVNext)
}
func (r *BlogPostScheduleParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.ContentScheduleRequestVNext)
}

type BlogPostSetLangPrimaryParams struct {
	// Request body object for setting a new primary language.
	SetNewLanguagePrimaryRequestVNext SetNewLanguagePrimaryRequestVNextParam
	paramObj
}

func (r BlogPostSetLangPrimaryParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.SetNewLanguagePrimaryRequestVNext)
}
func (r *BlogPostSetLangPrimaryParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.SetNewLanguagePrimaryRequestVNext)
}

type BlogPostUpdateDraftParams struct {
	// Model definition for a Blog Post.
	BlogPost BlogPostParam
	paramObj
}

func (r BlogPostUpdateDraftParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.BlogPost)
}
func (r *BlogPostUpdateDraftParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.BlogPost)
}

type BlogPostUpdateLangsParams struct {
	// Request object for updating languages within a multi-language group.
	UpdateLanguagesRequestVNext UpdateLanguagesRequestVNextParam
	paramObj
}

func (r BlogPostUpdateLangsParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.UpdateLanguagesRequestVNext)
}
func (r *BlogPostUpdateLangsParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.UpdateLanguagesRequestVNext)
}
