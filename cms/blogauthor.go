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

// BlogAuthorService contains methods and other services that help with interacting
// with the hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBlogAuthorService] method instead.
type BlogAuthorService struct {
	Options []option.RequestOption
}

// NewBlogAuthorService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewBlogAuthorService(opts ...option.RequestOption) (r BlogAuthorService) {
	r = BlogAuthorService{}
	r.Options = opts
	return
}

// Create a new Blog Author.
func (r *BlogAuthorService) New(ctx context.Context, body BlogAuthorNewParams, opts ...option.RequestOption) (res *BlogAuthor, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "cms/v3/blogs/authors"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Sparse updates a single Blog Author object identified by the id in the path. All
// the column values need not be specified. Only the that need to be modified can
// be specified.
func (r *BlogAuthorService) Update(ctx context.Context, objectID string, params BlogAuthorUpdateParams, opts ...option.RequestOption) (res *BlogAuthor, err error) {
	opts = slices.Concat(r.Options, opts)
	if objectID == "" {
		err = errors.New("missing required objectId parameter")
		return
	}
	path := fmt.Sprintf("cms/v3/blogs/authors/%s", objectID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return
}

// Get the list of blog authors. Supports paging and filtering. This method would
// be useful for an integration that examined these models and used an external
// service to suggest edits.
func (r *BlogAuthorService) List(ctx context.Context, query BlogAuthorListParams, opts ...option.RequestOption) (res *pagination.Page[BlogAuthor], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "cms/v3/blogs/authors"
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

// Get the list of blog authors. Supports paging and filtering. This method would
// be useful for an integration that examined these models and used an external
// service to suggest edits.
func (r *BlogAuthorService) ListAutoPaging(ctx context.Context, query BlogAuthorListParams, opts ...option.RequestOption) *pagination.PageAutoPager[BlogAuthor] {
	return pagination.NewPageAutoPager(r.List(ctx, query, opts...))
}

// Delete the Blog Author object identified by the id in the path.
func (r *BlogAuthorService) Delete(ctx context.Context, objectID string, body BlogAuthorDeleteParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if objectID == "" {
		err = errors.New("missing required objectId parameter")
		return
	}
	path := fmt.Sprintf("cms/v3/blogs/authors/%s", objectID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, nil, opts...)
	return
}

// Attach a Blog Author to a multi-language group.
func (r *BlogAuthorService) AttachToLangGroup(ctx context.Context, body BlogAuthorAttachToLangGroupParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "cms/v3/blogs/authors/multi-language/attach-to-lang-group"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Create the Blog Author objects detailed in the request body.
func (r *BlogAuthorService) NewBatch(ctx context.Context, body BlogAuthorNewBatchParams, opts ...option.RequestOption) (res *BatchResponseBlogAuthor, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "cms/v3/blogs/authors/batch/create"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Create a new language variation from an existing Blog Author.
func (r *BlogAuthorService) NewLanguageVariation(ctx context.Context, body BlogAuthorNewLanguageVariationParams, opts ...option.RequestOption) (res *BlogAuthor, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "cms/v3/blogs/authors/multi-language/create-language-variation"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Delete the Blog Author objects identified in the request body.
func (r *BlogAuthorService) DeleteBatch(ctx context.Context, body BlogAuthorDeleteBatchParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "cms/v3/blogs/authors/batch/archive"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Detach a Blog Author from a multi-language group.
func (r *BlogAuthorService) DetachFromLangGroup(ctx context.Context, body BlogAuthorDetachFromLangGroupParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "cms/v3/blogs/authors/multi-language/detach-from-lang-group"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Retrieve the Blog Author object identified by the id in the path.
func (r *BlogAuthorService) Get(ctx context.Context, objectID string, query BlogAuthorGetParams, opts ...option.RequestOption) (res *BlogAuthor, err error) {
	opts = slices.Concat(r.Options, opts)
	if objectID == "" {
		err = errors.New("missing required objectId parameter")
		return
	}
	path := fmt.Sprintf("cms/v3/blogs/authors/%s", objectID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Retrieve the Blog Author objects identified in the request body.
func (r *BlogAuthorService) GetBatch(ctx context.Context, params BlogAuthorGetBatchParams, opts ...option.RequestOption) (res *BatchResponseBlogAuthor, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "cms/v3/blogs/authors/batch/read"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Set a Blog Author as the primary language of a multi-language group.
func (r *BlogAuthorService) SetNewLangPrimary(ctx context.Context, body BlogAuthorSetNewLangPrimaryParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "cms/v3/blogs/authors/multi-language/set-new-lang-primary"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, nil, opts...)
	return
}

// Update the Blog Author objects identified in the request body.
func (r *BlogAuthorService) UpdateBatch(ctx context.Context, params BlogAuthorUpdateBatchParams, opts ...option.RequestOption) (res *BatchResponseBlogAuthor, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "cms/v3/blogs/authors/batch/update"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Explicitly set new languages for each Blog Author in a multi-language group.
func (r *BlogAuthorService) UpdateLanguages(ctx context.Context, body BlogAuthorUpdateLanguagesParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "cms/v3/blogs/authors/multi-language/update-languages"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Wrapper for providing an array of blog authors as inputs.
//
// The property Inputs is required.
type BatchInputBlogAuthorParam struct {
	// Blog authors to input.
	Inputs []BlogAuthorParam `json:"inputs,omitzero,required"`
	paramObj
}

func (r BatchInputBlogAuthorParam) MarshalJSON() (data []byte, err error) {
	type shadow BatchInputBlogAuthorParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BatchInputBlogAuthorParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Response object for batch operations on blog authors.
type BatchResponseBlogAuthor struct {
	// Time of batch operation completion.
	CompletedAt time.Time `json:"completedAt,required" format:"date-time"`
	// Results of batch operation.
	Results []BlogAuthor `json:"results,required"`
	// Time of batch operation start.
	StartedAt time.Time `json:"startedAt,required" format:"date-time"`
	// Status of batch operation.
	//
	// Any of "CANCELED", "COMPLETE", "PENDING", "PROCESSING".
	Status BatchResponseBlogAuthorStatus `json:"status,required"`
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
func (r BatchResponseBlogAuthor) RawJSON() string { return r.JSON.raw }
func (r *BatchResponseBlogAuthor) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Status of batch operation.
type BatchResponseBlogAuthorStatus string

const (
	BatchResponseBlogAuthorStatusCanceled   BatchResponseBlogAuthorStatus = "CANCELED"
	BatchResponseBlogAuthorStatusComplete   BatchResponseBlogAuthorStatus = "COMPLETE"
	BatchResponseBlogAuthorStatusPending    BatchResponseBlogAuthorStatus = "PENDING"
	BatchResponseBlogAuthorStatusProcessing BatchResponseBlogAuthorStatus = "PROCESSING"
)

// Model definition for a Blog Author.
type BlogAuthor struct {
	// The unique ID of the Blog Author.
	ID string `json:"id,required"`
	// URL to the blog author's avatar, if supplying a custom one.
	Avatar string `json:"avatar,required"`
	// A short biography of the blog author.
	Bio     string    `json:"bio,required"`
	Created time.Time `json:"created,required" format:"date-time"`
	// The timestamp (ISO8601 format) when this Blog Author was deleted.
	DeletedAt time.Time `json:"deletedAt,required" format:"date-time"`
	// The full name of the Blog Author to be displayed.
	DisplayName string `json:"displayName,required"`
	// Email address of the Blog Author.
	Email string `json:"email,required"`
	// URL to the Blog Author's Facebook page.
	Facebook string `json:"facebook,required"`
	FullName string `json:"fullName,required"`
	// The explicitly defined ISO 639 language code of the blog author.
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
	// "en-dm", "en-er", "en-fi", "en-fj", "en-fk", "en-fm", "en-gb", "en-gd", "en-gg",
	// "en-gh", "en-gi", "en-gm", "en-gu", "en-gy", "en-hk", "en-ie", "en-il", "en-im",
	// "en-in", "en-io", "en-je", "en-jm", "en-ke", "en-ki", "en-kn", "en-ky", "en-lc",
	// "en-lr", "en-ls", "en-lu", "en-mg", "en-mh", "en-mo", "en-mp", "en-ms", "en-mt",
	// "en-mu", "en-mw", "en-mx", "en-my", "en-na", "en-nf", "en-ng", "en-nl", "en-nr",
	// "en-nu", "en-nz", "en-pg", "en-ph", "en-pk", "en-pn", "en-pr", "en-pw", "en-rw",
	// "en-sb", "en-sc", "en-sd", "en-se", "en-sg", "en-sh", "en-si", "en-sl", "en-ss",
	// "en-sx", "en-sz", "en-tc", "en-tk", "en-to", "en-tt", "en-tv", "en-tz", "en-ug",
	// "en-um", "en-us", "en-vc", "en-vg", "en-vi", "en-vu", "en-ws", "en-za", "en-zm",
	// "en-zw", "eo", "eo-001", "es", "es-419", "es-ar", "es-bo", "es-br", "es-bz",
	// "es-cl", "es-co", "es-cr", "es-cu", "es-do", "es-ea", "es-ec", "es-es", "es-gq",
	// "es-gt", "es-hn", "es-ic", "es-mx", "es-ni", "es-pa", "es-pe", "es-ph", "es-pr",
	// "es-py", "es-sv", "es-us", "es-uy", "es-ve", "et", "et-ee", "eu", "eu-es",
	// "ewo", "ewo-cm", "fa", "fa-af", "fa-ir", "ff", "ff-bf", "ff-cm", "ff-gh",
	// "ff-gm", "ff-gn", "ff-gw", "ff-lr", "ff-mr", "ff-ne", "ff-ng", "ff-sl", "ff-sn",
	// "fi", "fi-fi", "fil", "fil-ph", "fo", "fo-dk", "fo-fo", "fr", "fr-be", "fr-bf",
	// "fr-bi", "fr-bj", "fr-bl", "fr-ca", "fr-cd", "fr-cf", "fr-cg", "fr-ch", "fr-ci",
	// "fr-cm", "fr-dj", "fr-dz", "fr-fr", "fr-ga", "fr-gf", "fr-gn", "fr-gp", "fr-gq",
	// "fr-ht", "fr-km", "fr-lu", "fr-ma", "fr-mc", "fr-mf", "fr-mg", "fr-ml", "fr-mq",
	// "fr-mr", "fr-mu", "fr-nc", "fr-ne", "fr-pf", "fr-pm", "fr-re", "fr-rw", "fr-sc",
	// "fr-sn", "fr-sy", "fr-td", "fr-tg", "fr-tn", "fr-vu", "fr-wf", "fr-yt", "fur",
	// "fur-it", "fy", "fy-nl", "ga", "ga-gb", "ga-ie", "gd", "gd-gb", "gl", "gl-es",
	// "gsw", "gsw-ch", "gsw-fr", "gsw-li", "gu", "gu-in", "guz", "guz-ke", "gv",
	// "gv-im", "ha", "ha-gh", "ha-ne", "ha-ng", "haw", "haw-us", "he", "he-il", "hi",
	// "hi-in", "hr", "hr-ba", "hr-hr", "hsb", "hsb-de", "hu", "hu-hu", "hy", "hy-am",
	// "ia", "ia-001", "id", "id-id", "ig", "ig-ng", "ii", "ii-cn", "is", "is-is",
	// "it", "it-ch", "it-it", "it-sm", "it-va", "ja", "ja-jp", "jgo", "jgo-cm", "jmc",
	// "jmc-tz", "jv", "jv-id", "ka", "ka-ge", "kab", "kab-dz", "kam", "kam-ke", "kde",
	// "kde-tz", "kea", "kea-cv", "khq", "khq-ml", "ki", "ki-ke", "kk", "kk-kz", "kkj",
	// "kkj-cm", "kl", "kl-gl", "kln", "kln-ke", "km", "km-kh", "kn", "kn-in", "ko",
	// "ko-kp", "ko-kr", "kok", "kok-in", "ks", "ks-in", "ksb", "ksb-tz", "ksf",
	// "ksf-cm", "ksh", "ksh-de", "ku", "ku-tr", "kw", "kw-gb", "ky", "ky-kg", "lag",
	// "lag-tz", "lb", "lb-lu", "lg", "lg-ug", "lkt", "lkt-us", "ln", "ln-ao", "ln-cd",
	// "ln-cf", "ln-cg", "lo", "lo-la", "lrc", "lrc-iq", "lrc-ir", "lt", "lt-lt", "lu",
	// "lu-cd", "luo", "luo-ke", "luy", "luy-ke", "lv", "lv-lv", "mai", "mai-in",
	// "mas", "mas-ke", "mas-tz", "mer", "mer-ke", "mfe", "mfe-mu", "mg", "mg-mg",
	// "mgh", "mgh-mz", "mgo", "mgo-cm", "mi", "mi-nz", "mk", "mk-mk", "ml", "ml-in",
	// "mn", "mn-mn", "mni", "mni-in", "mr", "mr-in", "ms", "ms-bn", "ms-id", "ms-my",
	// "ms-sg", "mt", "mt-mt", "mua", "mua-cm", "my", "my-mm", "mzn", "mzn-ir", "naq",
	// "naq-na", "nb", "nb-no", "nb-sj", "nd", "nd-zw", "nds", "nds-de", "nds-nl",
	// "ne", "ne-in", "ne-np", "nl", "nl-aw", "nl-be", "nl-bq", "nl-ch", "nl-cw",
	// "nl-lu", "nl-nl", "nl-sr", "nl-sx", "nmg", "nmg-cm", "nn", "nn-no", "nnh",
	// "nnh-cm", "no", "no-no", "nus", "nus-ss", "nyn", "nyn-ug", "om", "om-et",
	// "om-ke", "or", "or-in", "os", "os-ge", "os-ru", "pa", "pa-in", "pa-pk", "pcm",
	// "pcm-ng", "pl", "pl-pl", "prg", "prg-001", "ps", "ps-af", "ps-pk", "pt",
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
	Language BlogAuthorLanguage `json:"language,required"`
	// URL to the blog author's LinkedIn page.
	Linkedin string `json:"linkedin,required"`
	Name     string `json:"name,required"`
	Slug     string `json:"slug,required"`
	// ID of the primary blog author this object was translated from.
	TranslatedFromID int64 `json:"translatedFromId,required"`
	// URL or username of the Twitter account associated with the Blog Author. This
	// will be normalized into the Twitter url for said user.
	Twitter string    `json:"twitter,required"`
	Updated time.Time `json:"updated,required" format:"date-time"`
	// URL to the website of the Blog Author.
	Website string `json:"website,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		Avatar           respjson.Field
		Bio              respjson.Field
		Created          respjson.Field
		DeletedAt        respjson.Field
		DisplayName      respjson.Field
		Email            respjson.Field
		Facebook         respjson.Field
		FullName         respjson.Field
		Language         respjson.Field
		Linkedin         respjson.Field
		Name             respjson.Field
		Slug             respjson.Field
		TranslatedFromID respjson.Field
		Twitter          respjson.Field
		Updated          respjson.Field
		Website          respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BlogAuthor) RawJSON() string { return r.JSON.raw }
func (r *BlogAuthor) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this BlogAuthor to a BlogAuthorParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// BlogAuthorParam.Overrides()
func (r BlogAuthor) ToParam() BlogAuthorParam {
	return param.Override[BlogAuthorParam](json.RawMessage(r.RawJSON()))
}

// The explicitly defined ISO 639 language code of the blog author.
type BlogAuthorLanguage string

const (
	BlogAuthorLanguageAf     BlogAuthorLanguage = "af"
	BlogAuthorLanguageAfNa   BlogAuthorLanguage = "af-na"
	BlogAuthorLanguageAfZa   BlogAuthorLanguage = "af-za"
	BlogAuthorLanguageAgq    BlogAuthorLanguage = "agq"
	BlogAuthorLanguageAgqCm  BlogAuthorLanguage = "agq-cm"
	BlogAuthorLanguageAk     BlogAuthorLanguage = "ak"
	BlogAuthorLanguageAkGh   BlogAuthorLanguage = "ak-gh"
	BlogAuthorLanguageAm     BlogAuthorLanguage = "am"
	BlogAuthorLanguageAmEt   BlogAuthorLanguage = "am-et"
	BlogAuthorLanguageAr     BlogAuthorLanguage = "ar"
	BlogAuthorLanguageAr001  BlogAuthorLanguage = "ar-001"
	BlogAuthorLanguageArAe   BlogAuthorLanguage = "ar-ae"
	BlogAuthorLanguageArBh   BlogAuthorLanguage = "ar-bh"
	BlogAuthorLanguageArDj   BlogAuthorLanguage = "ar-dj"
	BlogAuthorLanguageArDz   BlogAuthorLanguage = "ar-dz"
	BlogAuthorLanguageArEg   BlogAuthorLanguage = "ar-eg"
	BlogAuthorLanguageArEh   BlogAuthorLanguage = "ar-eh"
	BlogAuthorLanguageArEr   BlogAuthorLanguage = "ar-er"
	BlogAuthorLanguageArIl   BlogAuthorLanguage = "ar-il"
	BlogAuthorLanguageArIq   BlogAuthorLanguage = "ar-iq"
	BlogAuthorLanguageArJo   BlogAuthorLanguage = "ar-jo"
	BlogAuthorLanguageArKm   BlogAuthorLanguage = "ar-km"
	BlogAuthorLanguageArKw   BlogAuthorLanguage = "ar-kw"
	BlogAuthorLanguageArLb   BlogAuthorLanguage = "ar-lb"
	BlogAuthorLanguageArLy   BlogAuthorLanguage = "ar-ly"
	BlogAuthorLanguageArMa   BlogAuthorLanguage = "ar-ma"
	BlogAuthorLanguageArMr   BlogAuthorLanguage = "ar-mr"
	BlogAuthorLanguageArOm   BlogAuthorLanguage = "ar-om"
	BlogAuthorLanguageArPs   BlogAuthorLanguage = "ar-ps"
	BlogAuthorLanguageArQa   BlogAuthorLanguage = "ar-qa"
	BlogAuthorLanguageArSa   BlogAuthorLanguage = "ar-sa"
	BlogAuthorLanguageArSd   BlogAuthorLanguage = "ar-sd"
	BlogAuthorLanguageArSo   BlogAuthorLanguage = "ar-so"
	BlogAuthorLanguageArSS   BlogAuthorLanguage = "ar-ss"
	BlogAuthorLanguageArSy   BlogAuthorLanguage = "ar-sy"
	BlogAuthorLanguageArTd   BlogAuthorLanguage = "ar-td"
	BlogAuthorLanguageArTn   BlogAuthorLanguage = "ar-tn"
	BlogAuthorLanguageArYe   BlogAuthorLanguage = "ar-ye"
	BlogAuthorLanguageAs     BlogAuthorLanguage = "as"
	BlogAuthorLanguageAsIn   BlogAuthorLanguage = "as-in"
	BlogAuthorLanguageAsa    BlogAuthorLanguage = "asa"
	BlogAuthorLanguageAsaTz  BlogAuthorLanguage = "asa-tz"
	BlogAuthorLanguageAst    BlogAuthorLanguage = "ast"
	BlogAuthorLanguageAstEs  BlogAuthorLanguage = "ast-es"
	BlogAuthorLanguageAz     BlogAuthorLanguage = "az"
	BlogAuthorLanguageAzAz   BlogAuthorLanguage = "az-az"
	BlogAuthorLanguageBas    BlogAuthorLanguage = "bas"
	BlogAuthorLanguageBasCm  BlogAuthorLanguage = "bas-cm"
	BlogAuthorLanguageBe     BlogAuthorLanguage = "be"
	BlogAuthorLanguageBeBy   BlogAuthorLanguage = "be-by"
	BlogAuthorLanguageBem    BlogAuthorLanguage = "bem"
	BlogAuthorLanguageBemZm  BlogAuthorLanguage = "bem-zm"
	BlogAuthorLanguageBez    BlogAuthorLanguage = "bez"
	BlogAuthorLanguageBezTz  BlogAuthorLanguage = "bez-tz"
	BlogAuthorLanguageBg     BlogAuthorLanguage = "bg"
	BlogAuthorLanguageBgBg   BlogAuthorLanguage = "bg-bg"
	BlogAuthorLanguageBm     BlogAuthorLanguage = "bm"
	BlogAuthorLanguageBmMl   BlogAuthorLanguage = "bm-ml"
	BlogAuthorLanguageBn     BlogAuthorLanguage = "bn"
	BlogAuthorLanguageBnBd   BlogAuthorLanguage = "bn-bd"
	BlogAuthorLanguageBnIn   BlogAuthorLanguage = "bn-in"
	BlogAuthorLanguageBo     BlogAuthorLanguage = "bo"
	BlogAuthorLanguageBoCn   BlogAuthorLanguage = "bo-cn"
	BlogAuthorLanguageBoIn   BlogAuthorLanguage = "bo-in"
	BlogAuthorLanguageBr     BlogAuthorLanguage = "br"
	BlogAuthorLanguageBrFr   BlogAuthorLanguage = "br-fr"
	BlogAuthorLanguageBrx    BlogAuthorLanguage = "brx"
	BlogAuthorLanguageBrxIn  BlogAuthorLanguage = "brx-in"
	BlogAuthorLanguageBs     BlogAuthorLanguage = "bs"
	BlogAuthorLanguageBsBa   BlogAuthorLanguage = "bs-ba"
	BlogAuthorLanguageCa     BlogAuthorLanguage = "ca"
	BlogAuthorLanguageCaAd   BlogAuthorLanguage = "ca-ad"
	BlogAuthorLanguageCaEs   BlogAuthorLanguage = "ca-es"
	BlogAuthorLanguageCaFr   BlogAuthorLanguage = "ca-fr"
	BlogAuthorLanguageCaIt   BlogAuthorLanguage = "ca-it"
	BlogAuthorLanguageCcp    BlogAuthorLanguage = "ccp"
	BlogAuthorLanguageCcpBd  BlogAuthorLanguage = "ccp-bd"
	BlogAuthorLanguageCcpIn  BlogAuthorLanguage = "ccp-in"
	BlogAuthorLanguageCe     BlogAuthorLanguage = "ce"
	BlogAuthorLanguageCeRu   BlogAuthorLanguage = "ce-ru"
	BlogAuthorLanguageCeb    BlogAuthorLanguage = "ceb"
	BlogAuthorLanguageCebPh  BlogAuthorLanguage = "ceb-ph"
	BlogAuthorLanguageCgg    BlogAuthorLanguage = "cgg"
	BlogAuthorLanguageCggUg  BlogAuthorLanguage = "cgg-ug"
	BlogAuthorLanguageChr    BlogAuthorLanguage = "chr"
	BlogAuthorLanguageChrUs  BlogAuthorLanguage = "chr-us"
	BlogAuthorLanguageCkb    BlogAuthorLanguage = "ckb"
	BlogAuthorLanguageCkbIq  BlogAuthorLanguage = "ckb-iq"
	BlogAuthorLanguageCkbIr  BlogAuthorLanguage = "ckb-ir"
	BlogAuthorLanguageCs     BlogAuthorLanguage = "cs"
	BlogAuthorLanguageCsCz   BlogAuthorLanguage = "cs-cz"
	BlogAuthorLanguageCu     BlogAuthorLanguage = "cu"
	BlogAuthorLanguageCuRu   BlogAuthorLanguage = "cu-ru"
	BlogAuthorLanguageCy     BlogAuthorLanguage = "cy"
	BlogAuthorLanguageCyGB   BlogAuthorLanguage = "cy-gb"
	BlogAuthorLanguageDa     BlogAuthorLanguage = "da"
	BlogAuthorLanguageDaDk   BlogAuthorLanguage = "da-dk"
	BlogAuthorLanguageDaGl   BlogAuthorLanguage = "da-gl"
	BlogAuthorLanguageDav    BlogAuthorLanguage = "dav"
	BlogAuthorLanguageDavKe  BlogAuthorLanguage = "dav-ke"
	BlogAuthorLanguageDe     BlogAuthorLanguage = "de"
	BlogAuthorLanguageDeAt   BlogAuthorLanguage = "de-at"
	BlogAuthorLanguageDeBe   BlogAuthorLanguage = "de-be"
	BlogAuthorLanguageDeCh   BlogAuthorLanguage = "de-ch"
	BlogAuthorLanguageDeDe   BlogAuthorLanguage = "de-de"
	BlogAuthorLanguageDeGr   BlogAuthorLanguage = "de-gr"
	BlogAuthorLanguageDeIt   BlogAuthorLanguage = "de-it"
	BlogAuthorLanguageDeLi   BlogAuthorLanguage = "de-li"
	BlogAuthorLanguageDeLu   BlogAuthorLanguage = "de-lu"
	BlogAuthorLanguageDje    BlogAuthorLanguage = "dje"
	BlogAuthorLanguageDjeNe  BlogAuthorLanguage = "dje-ne"
	BlogAuthorLanguageDoi    BlogAuthorLanguage = "doi"
	BlogAuthorLanguageDoiIn  BlogAuthorLanguage = "doi-in"
	BlogAuthorLanguageDsb    BlogAuthorLanguage = "dsb"
	BlogAuthorLanguageDsbDe  BlogAuthorLanguage = "dsb-de"
	BlogAuthorLanguageDua    BlogAuthorLanguage = "dua"
	BlogAuthorLanguageDuaCm  BlogAuthorLanguage = "dua-cm"
	BlogAuthorLanguageDyo    BlogAuthorLanguage = "dyo"
	BlogAuthorLanguageDyoSn  BlogAuthorLanguage = "dyo-sn"
	BlogAuthorLanguageDz     BlogAuthorLanguage = "dz"
	BlogAuthorLanguageDzBt   BlogAuthorLanguage = "dz-bt"
	BlogAuthorLanguageEbu    BlogAuthorLanguage = "ebu"
	BlogAuthorLanguageEbuKe  BlogAuthorLanguage = "ebu-ke"
	BlogAuthorLanguageEe     BlogAuthorLanguage = "ee"
	BlogAuthorLanguageEeGh   BlogAuthorLanguage = "ee-gh"
	BlogAuthorLanguageEeTg   BlogAuthorLanguage = "ee-tg"
	BlogAuthorLanguageEl     BlogAuthorLanguage = "el"
	BlogAuthorLanguageElCy   BlogAuthorLanguage = "el-cy"
	BlogAuthorLanguageElGr   BlogAuthorLanguage = "el-gr"
	BlogAuthorLanguageEn     BlogAuthorLanguage = "en"
	BlogAuthorLanguageEn001  BlogAuthorLanguage = "en-001"
	BlogAuthorLanguageEn150  BlogAuthorLanguage = "en-150"
	BlogAuthorLanguageEnAe   BlogAuthorLanguage = "en-ae"
	BlogAuthorLanguageEnAg   BlogAuthorLanguage = "en-ag"
	BlogAuthorLanguageEnAI   BlogAuthorLanguage = "en-ai"
	BlogAuthorLanguageEnAs   BlogAuthorLanguage = "en-as"
	BlogAuthorLanguageEnAt   BlogAuthorLanguage = "en-at"
	BlogAuthorLanguageEnAu   BlogAuthorLanguage = "en-au"
	BlogAuthorLanguageEnBb   BlogAuthorLanguage = "en-bb"
	BlogAuthorLanguageEnBe   BlogAuthorLanguage = "en-be"
	BlogAuthorLanguageEnBi   BlogAuthorLanguage = "en-bi"
	BlogAuthorLanguageEnBm   BlogAuthorLanguage = "en-bm"
	BlogAuthorLanguageEnBs   BlogAuthorLanguage = "en-bs"
	BlogAuthorLanguageEnBw   BlogAuthorLanguage = "en-bw"
	BlogAuthorLanguageEnBz   BlogAuthorLanguage = "en-bz"
	BlogAuthorLanguageEnCa   BlogAuthorLanguage = "en-ca"
	BlogAuthorLanguageEnCc   BlogAuthorLanguage = "en-cc"
	BlogAuthorLanguageEnCh   BlogAuthorLanguage = "en-ch"
	BlogAuthorLanguageEnCk   BlogAuthorLanguage = "en-ck"
	BlogAuthorLanguageEnCm   BlogAuthorLanguage = "en-cm"
	BlogAuthorLanguageEnCn   BlogAuthorLanguage = "en-cn"
	BlogAuthorLanguageEnCx   BlogAuthorLanguage = "en-cx"
	BlogAuthorLanguageEnCy   BlogAuthorLanguage = "en-cy"
	BlogAuthorLanguageEnDe   BlogAuthorLanguage = "en-de"
	BlogAuthorLanguageEnDg   BlogAuthorLanguage = "en-dg"
	BlogAuthorLanguageEnDk   BlogAuthorLanguage = "en-dk"
	BlogAuthorLanguageEnDm   BlogAuthorLanguage = "en-dm"
	BlogAuthorLanguageEnEr   BlogAuthorLanguage = "en-er"
	BlogAuthorLanguageEnFi   BlogAuthorLanguage = "en-fi"
	BlogAuthorLanguageEnFj   BlogAuthorLanguage = "en-fj"
	BlogAuthorLanguageEnFk   BlogAuthorLanguage = "en-fk"
	BlogAuthorLanguageEnFm   BlogAuthorLanguage = "en-fm"
	BlogAuthorLanguageEnGB   BlogAuthorLanguage = "en-gb"
	BlogAuthorLanguageEnGd   BlogAuthorLanguage = "en-gd"
	BlogAuthorLanguageEnGg   BlogAuthorLanguage = "en-gg"
	BlogAuthorLanguageEnGh   BlogAuthorLanguage = "en-gh"
	BlogAuthorLanguageEnGi   BlogAuthorLanguage = "en-gi"
	BlogAuthorLanguageEnGm   BlogAuthorLanguage = "en-gm"
	BlogAuthorLanguageEnGu   BlogAuthorLanguage = "en-gu"
	BlogAuthorLanguageEnGy   BlogAuthorLanguage = "en-gy"
	BlogAuthorLanguageEnHk   BlogAuthorLanguage = "en-hk"
	BlogAuthorLanguageEnIe   BlogAuthorLanguage = "en-ie"
	BlogAuthorLanguageEnIl   BlogAuthorLanguage = "en-il"
	BlogAuthorLanguageEnIm   BlogAuthorLanguage = "en-im"
	BlogAuthorLanguageEnIn   BlogAuthorLanguage = "en-in"
	BlogAuthorLanguageEnIo   BlogAuthorLanguage = "en-io"
	BlogAuthorLanguageEnJe   BlogAuthorLanguage = "en-je"
	BlogAuthorLanguageEnJm   BlogAuthorLanguage = "en-jm"
	BlogAuthorLanguageEnKe   BlogAuthorLanguage = "en-ke"
	BlogAuthorLanguageEnKi   BlogAuthorLanguage = "en-ki"
	BlogAuthorLanguageEnKn   BlogAuthorLanguage = "en-kn"
	BlogAuthorLanguageEnKy   BlogAuthorLanguage = "en-ky"
	BlogAuthorLanguageEnLc   BlogAuthorLanguage = "en-lc"
	BlogAuthorLanguageEnLr   BlogAuthorLanguage = "en-lr"
	BlogAuthorLanguageEnLs   BlogAuthorLanguage = "en-ls"
	BlogAuthorLanguageEnLu   BlogAuthorLanguage = "en-lu"
	BlogAuthorLanguageEnMg   BlogAuthorLanguage = "en-mg"
	BlogAuthorLanguageEnMh   BlogAuthorLanguage = "en-mh"
	BlogAuthorLanguageEnMo   BlogAuthorLanguage = "en-mo"
	BlogAuthorLanguageEnMp   BlogAuthorLanguage = "en-mp"
	BlogAuthorLanguageEnMs   BlogAuthorLanguage = "en-ms"
	BlogAuthorLanguageEnMt   BlogAuthorLanguage = "en-mt"
	BlogAuthorLanguageEnMu   BlogAuthorLanguage = "en-mu"
	BlogAuthorLanguageEnMw   BlogAuthorLanguage = "en-mw"
	BlogAuthorLanguageEnMx   BlogAuthorLanguage = "en-mx"
	BlogAuthorLanguageEnMy   BlogAuthorLanguage = "en-my"
	BlogAuthorLanguageEnNa   BlogAuthorLanguage = "en-na"
	BlogAuthorLanguageEnNf   BlogAuthorLanguage = "en-nf"
	BlogAuthorLanguageEnNg   BlogAuthorLanguage = "en-ng"
	BlogAuthorLanguageEnNl   BlogAuthorLanguage = "en-nl"
	BlogAuthorLanguageEnNr   BlogAuthorLanguage = "en-nr"
	BlogAuthorLanguageEnNu   BlogAuthorLanguage = "en-nu"
	BlogAuthorLanguageEnNz   BlogAuthorLanguage = "en-nz"
	BlogAuthorLanguageEnPg   BlogAuthorLanguage = "en-pg"
	BlogAuthorLanguageEnPh   BlogAuthorLanguage = "en-ph"
	BlogAuthorLanguageEnPk   BlogAuthorLanguage = "en-pk"
	BlogAuthorLanguageEnPn   BlogAuthorLanguage = "en-pn"
	BlogAuthorLanguageEnPr   BlogAuthorLanguage = "en-pr"
	BlogAuthorLanguageEnPw   BlogAuthorLanguage = "en-pw"
	BlogAuthorLanguageEnRw   BlogAuthorLanguage = "en-rw"
	BlogAuthorLanguageEnSb   BlogAuthorLanguage = "en-sb"
	BlogAuthorLanguageEnSc   BlogAuthorLanguage = "en-sc"
	BlogAuthorLanguageEnSd   BlogAuthorLanguage = "en-sd"
	BlogAuthorLanguageEnSe   BlogAuthorLanguage = "en-se"
	BlogAuthorLanguageEnSg   BlogAuthorLanguage = "en-sg"
	BlogAuthorLanguageEnSh   BlogAuthorLanguage = "en-sh"
	BlogAuthorLanguageEnSi   BlogAuthorLanguage = "en-si"
	BlogAuthorLanguageEnSl   BlogAuthorLanguage = "en-sl"
	BlogAuthorLanguageEnSS   BlogAuthorLanguage = "en-ss"
	BlogAuthorLanguageEnSx   BlogAuthorLanguage = "en-sx"
	BlogAuthorLanguageEnSz   BlogAuthorLanguage = "en-sz"
	BlogAuthorLanguageEnTc   BlogAuthorLanguage = "en-tc"
	BlogAuthorLanguageEnTk   BlogAuthorLanguage = "en-tk"
	BlogAuthorLanguageEnTo   BlogAuthorLanguage = "en-to"
	BlogAuthorLanguageEnTt   BlogAuthorLanguage = "en-tt"
	BlogAuthorLanguageEnTv   BlogAuthorLanguage = "en-tv"
	BlogAuthorLanguageEnTz   BlogAuthorLanguage = "en-tz"
	BlogAuthorLanguageEnUg   BlogAuthorLanguage = "en-ug"
	BlogAuthorLanguageEnUm   BlogAuthorLanguage = "en-um"
	BlogAuthorLanguageEnUs   BlogAuthorLanguage = "en-us"
	BlogAuthorLanguageEnVc   BlogAuthorLanguage = "en-vc"
	BlogAuthorLanguageEnVg   BlogAuthorLanguage = "en-vg"
	BlogAuthorLanguageEnVi   BlogAuthorLanguage = "en-vi"
	BlogAuthorLanguageEnVu   BlogAuthorLanguage = "en-vu"
	BlogAuthorLanguageEnWs   BlogAuthorLanguage = "en-ws"
	BlogAuthorLanguageEnZa   BlogAuthorLanguage = "en-za"
	BlogAuthorLanguageEnZm   BlogAuthorLanguage = "en-zm"
	BlogAuthorLanguageEnZw   BlogAuthorLanguage = "en-zw"
	BlogAuthorLanguageEo     BlogAuthorLanguage = "eo"
	BlogAuthorLanguageEo001  BlogAuthorLanguage = "eo-001"
	BlogAuthorLanguageEs     BlogAuthorLanguage = "es"
	BlogAuthorLanguageEs419  BlogAuthorLanguage = "es-419"
	BlogAuthorLanguageEsAr   BlogAuthorLanguage = "es-ar"
	BlogAuthorLanguageEsBo   BlogAuthorLanguage = "es-bo"
	BlogAuthorLanguageEsBr   BlogAuthorLanguage = "es-br"
	BlogAuthorLanguageEsBz   BlogAuthorLanguage = "es-bz"
	BlogAuthorLanguageEsCl   BlogAuthorLanguage = "es-cl"
	BlogAuthorLanguageEsCo   BlogAuthorLanguage = "es-co"
	BlogAuthorLanguageEsCr   BlogAuthorLanguage = "es-cr"
	BlogAuthorLanguageEsCu   BlogAuthorLanguage = "es-cu"
	BlogAuthorLanguageEsDo   BlogAuthorLanguage = "es-do"
	BlogAuthorLanguageEsEa   BlogAuthorLanguage = "es-ea"
	BlogAuthorLanguageEsEc   BlogAuthorLanguage = "es-ec"
	BlogAuthorLanguageEsEs   BlogAuthorLanguage = "es-es"
	BlogAuthorLanguageEsGq   BlogAuthorLanguage = "es-gq"
	BlogAuthorLanguageEsGt   BlogAuthorLanguage = "es-gt"
	BlogAuthorLanguageEsHn   BlogAuthorLanguage = "es-hn"
	BlogAuthorLanguageEsIc   BlogAuthorLanguage = "es-ic"
	BlogAuthorLanguageEsMx   BlogAuthorLanguage = "es-mx"
	BlogAuthorLanguageEsNi   BlogAuthorLanguage = "es-ni"
	BlogAuthorLanguageEsPa   BlogAuthorLanguage = "es-pa"
	BlogAuthorLanguageEsPe   BlogAuthorLanguage = "es-pe"
	BlogAuthorLanguageEsPh   BlogAuthorLanguage = "es-ph"
	BlogAuthorLanguageEsPr   BlogAuthorLanguage = "es-pr"
	BlogAuthorLanguageEsPy   BlogAuthorLanguage = "es-py"
	BlogAuthorLanguageEsSv   BlogAuthorLanguage = "es-sv"
	BlogAuthorLanguageEsUs   BlogAuthorLanguage = "es-us"
	BlogAuthorLanguageEsUy   BlogAuthorLanguage = "es-uy"
	BlogAuthorLanguageEsVe   BlogAuthorLanguage = "es-ve"
	BlogAuthorLanguageEt     BlogAuthorLanguage = "et"
	BlogAuthorLanguageEtEe   BlogAuthorLanguage = "et-ee"
	BlogAuthorLanguageEu     BlogAuthorLanguage = "eu"
	BlogAuthorLanguageEuEs   BlogAuthorLanguage = "eu-es"
	BlogAuthorLanguageEwo    BlogAuthorLanguage = "ewo"
	BlogAuthorLanguageEwoCm  BlogAuthorLanguage = "ewo-cm"
	BlogAuthorLanguageFa     BlogAuthorLanguage = "fa"
	BlogAuthorLanguageFaAf   BlogAuthorLanguage = "fa-af"
	BlogAuthorLanguageFaIr   BlogAuthorLanguage = "fa-ir"
	BlogAuthorLanguageFf     BlogAuthorLanguage = "ff"
	BlogAuthorLanguageFfBf   BlogAuthorLanguage = "ff-bf"
	BlogAuthorLanguageFfCm   BlogAuthorLanguage = "ff-cm"
	BlogAuthorLanguageFfGh   BlogAuthorLanguage = "ff-gh"
	BlogAuthorLanguageFfGm   BlogAuthorLanguage = "ff-gm"
	BlogAuthorLanguageFfGn   BlogAuthorLanguage = "ff-gn"
	BlogAuthorLanguageFfGw   BlogAuthorLanguage = "ff-gw"
	BlogAuthorLanguageFfLr   BlogAuthorLanguage = "ff-lr"
	BlogAuthorLanguageFfMr   BlogAuthorLanguage = "ff-mr"
	BlogAuthorLanguageFfNe   BlogAuthorLanguage = "ff-ne"
	BlogAuthorLanguageFfNg   BlogAuthorLanguage = "ff-ng"
	BlogAuthorLanguageFfSl   BlogAuthorLanguage = "ff-sl"
	BlogAuthorLanguageFfSn   BlogAuthorLanguage = "ff-sn"
	BlogAuthorLanguageFi     BlogAuthorLanguage = "fi"
	BlogAuthorLanguageFiFi   BlogAuthorLanguage = "fi-fi"
	BlogAuthorLanguageFil    BlogAuthorLanguage = "fil"
	BlogAuthorLanguageFilPh  BlogAuthorLanguage = "fil-ph"
	BlogAuthorLanguageFo     BlogAuthorLanguage = "fo"
	BlogAuthorLanguageFoDk   BlogAuthorLanguage = "fo-dk"
	BlogAuthorLanguageFoFo   BlogAuthorLanguage = "fo-fo"
	BlogAuthorLanguageFr     BlogAuthorLanguage = "fr"
	BlogAuthorLanguageFrBe   BlogAuthorLanguage = "fr-be"
	BlogAuthorLanguageFrBf   BlogAuthorLanguage = "fr-bf"
	BlogAuthorLanguageFrBi   BlogAuthorLanguage = "fr-bi"
	BlogAuthorLanguageFrBj   BlogAuthorLanguage = "fr-bj"
	BlogAuthorLanguageFrBl   BlogAuthorLanguage = "fr-bl"
	BlogAuthorLanguageFrCa   BlogAuthorLanguage = "fr-ca"
	BlogAuthorLanguageFrCd   BlogAuthorLanguage = "fr-cd"
	BlogAuthorLanguageFrCf   BlogAuthorLanguage = "fr-cf"
	BlogAuthorLanguageFrCg   BlogAuthorLanguage = "fr-cg"
	BlogAuthorLanguageFrCh   BlogAuthorLanguage = "fr-ch"
	BlogAuthorLanguageFrCi   BlogAuthorLanguage = "fr-ci"
	BlogAuthorLanguageFrCm   BlogAuthorLanguage = "fr-cm"
	BlogAuthorLanguageFrDj   BlogAuthorLanguage = "fr-dj"
	BlogAuthorLanguageFrDz   BlogAuthorLanguage = "fr-dz"
	BlogAuthorLanguageFrFr   BlogAuthorLanguage = "fr-fr"
	BlogAuthorLanguageFrGa   BlogAuthorLanguage = "fr-ga"
	BlogAuthorLanguageFrGf   BlogAuthorLanguage = "fr-gf"
	BlogAuthorLanguageFrGn   BlogAuthorLanguage = "fr-gn"
	BlogAuthorLanguageFrGp   BlogAuthorLanguage = "fr-gp"
	BlogAuthorLanguageFrGq   BlogAuthorLanguage = "fr-gq"
	BlogAuthorLanguageFrHt   BlogAuthorLanguage = "fr-ht"
	BlogAuthorLanguageFrKm   BlogAuthorLanguage = "fr-km"
	BlogAuthorLanguageFrLu   BlogAuthorLanguage = "fr-lu"
	BlogAuthorLanguageFrMa   BlogAuthorLanguage = "fr-ma"
	BlogAuthorLanguageFrMc   BlogAuthorLanguage = "fr-mc"
	BlogAuthorLanguageFrMf   BlogAuthorLanguage = "fr-mf"
	BlogAuthorLanguageFrMg   BlogAuthorLanguage = "fr-mg"
	BlogAuthorLanguageFrMl   BlogAuthorLanguage = "fr-ml"
	BlogAuthorLanguageFrMq   BlogAuthorLanguage = "fr-mq"
	BlogAuthorLanguageFrMr   BlogAuthorLanguage = "fr-mr"
	BlogAuthorLanguageFrMu   BlogAuthorLanguage = "fr-mu"
	BlogAuthorLanguageFrNc   BlogAuthorLanguage = "fr-nc"
	BlogAuthorLanguageFrNe   BlogAuthorLanguage = "fr-ne"
	BlogAuthorLanguageFrPf   BlogAuthorLanguage = "fr-pf"
	BlogAuthorLanguageFrPm   BlogAuthorLanguage = "fr-pm"
	BlogAuthorLanguageFrRe   BlogAuthorLanguage = "fr-re"
	BlogAuthorLanguageFrRw   BlogAuthorLanguage = "fr-rw"
	BlogAuthorLanguageFrSc   BlogAuthorLanguage = "fr-sc"
	BlogAuthorLanguageFrSn   BlogAuthorLanguage = "fr-sn"
	BlogAuthorLanguageFrSy   BlogAuthorLanguage = "fr-sy"
	BlogAuthorLanguageFrTd   BlogAuthorLanguage = "fr-td"
	BlogAuthorLanguageFrTg   BlogAuthorLanguage = "fr-tg"
	BlogAuthorLanguageFrTn   BlogAuthorLanguage = "fr-tn"
	BlogAuthorLanguageFrVu   BlogAuthorLanguage = "fr-vu"
	BlogAuthorLanguageFrWf   BlogAuthorLanguage = "fr-wf"
	BlogAuthorLanguageFrYt   BlogAuthorLanguage = "fr-yt"
	BlogAuthorLanguageFur    BlogAuthorLanguage = "fur"
	BlogAuthorLanguageFurIt  BlogAuthorLanguage = "fur-it"
	BlogAuthorLanguageFy     BlogAuthorLanguage = "fy"
	BlogAuthorLanguageFyNl   BlogAuthorLanguage = "fy-nl"
	BlogAuthorLanguageGa     BlogAuthorLanguage = "ga"
	BlogAuthorLanguageGaGB   BlogAuthorLanguage = "ga-gb"
	BlogAuthorLanguageGaIe   BlogAuthorLanguage = "ga-ie"
	BlogAuthorLanguageGd     BlogAuthorLanguage = "gd"
	BlogAuthorLanguageGdGB   BlogAuthorLanguage = "gd-gb"
	BlogAuthorLanguageGl     BlogAuthorLanguage = "gl"
	BlogAuthorLanguageGlEs   BlogAuthorLanguage = "gl-es"
	BlogAuthorLanguageGsw    BlogAuthorLanguage = "gsw"
	BlogAuthorLanguageGswCh  BlogAuthorLanguage = "gsw-ch"
	BlogAuthorLanguageGswFr  BlogAuthorLanguage = "gsw-fr"
	BlogAuthorLanguageGswLi  BlogAuthorLanguage = "gsw-li"
	BlogAuthorLanguageGu     BlogAuthorLanguage = "gu"
	BlogAuthorLanguageGuIn   BlogAuthorLanguage = "gu-in"
	BlogAuthorLanguageGuz    BlogAuthorLanguage = "guz"
	BlogAuthorLanguageGuzKe  BlogAuthorLanguage = "guz-ke"
	BlogAuthorLanguageGv     BlogAuthorLanguage = "gv"
	BlogAuthorLanguageGvIm   BlogAuthorLanguage = "gv-im"
	BlogAuthorLanguageHa     BlogAuthorLanguage = "ha"
	BlogAuthorLanguageHaGh   BlogAuthorLanguage = "ha-gh"
	BlogAuthorLanguageHaNe   BlogAuthorLanguage = "ha-ne"
	BlogAuthorLanguageHaNg   BlogAuthorLanguage = "ha-ng"
	BlogAuthorLanguageHaw    BlogAuthorLanguage = "haw"
	BlogAuthorLanguageHawUs  BlogAuthorLanguage = "haw-us"
	BlogAuthorLanguageHe     BlogAuthorLanguage = "he"
	BlogAuthorLanguageHeIl   BlogAuthorLanguage = "he-il"
	BlogAuthorLanguageHi     BlogAuthorLanguage = "hi"
	BlogAuthorLanguageHiIn   BlogAuthorLanguage = "hi-in"
	BlogAuthorLanguageHr     BlogAuthorLanguage = "hr"
	BlogAuthorLanguageHrBa   BlogAuthorLanguage = "hr-ba"
	BlogAuthorLanguageHrHr   BlogAuthorLanguage = "hr-hr"
	BlogAuthorLanguageHsb    BlogAuthorLanguage = "hsb"
	BlogAuthorLanguageHsbDe  BlogAuthorLanguage = "hsb-de"
	BlogAuthorLanguageHu     BlogAuthorLanguage = "hu"
	BlogAuthorLanguageHuHu   BlogAuthorLanguage = "hu-hu"
	BlogAuthorLanguageHy     BlogAuthorLanguage = "hy"
	BlogAuthorLanguageHyAm   BlogAuthorLanguage = "hy-am"
	BlogAuthorLanguageIa     BlogAuthorLanguage = "ia"
	BlogAuthorLanguageIa001  BlogAuthorLanguage = "ia-001"
	BlogAuthorLanguageID     BlogAuthorLanguage = "id"
	BlogAuthorLanguageIDID   BlogAuthorLanguage = "id-id"
	BlogAuthorLanguageIg     BlogAuthorLanguage = "ig"
	BlogAuthorLanguageIgNg   BlogAuthorLanguage = "ig-ng"
	BlogAuthorLanguageIi     BlogAuthorLanguage = "ii"
	BlogAuthorLanguageIiCn   BlogAuthorLanguage = "ii-cn"
	BlogAuthorLanguageIs     BlogAuthorLanguage = "is"
	BlogAuthorLanguageIsIs   BlogAuthorLanguage = "is-is"
	BlogAuthorLanguageIt     BlogAuthorLanguage = "it"
	BlogAuthorLanguageItCh   BlogAuthorLanguage = "it-ch"
	BlogAuthorLanguageItIt   BlogAuthorLanguage = "it-it"
	BlogAuthorLanguageItSm   BlogAuthorLanguage = "it-sm"
	BlogAuthorLanguageItVa   BlogAuthorLanguage = "it-va"
	BlogAuthorLanguageJa     BlogAuthorLanguage = "ja"
	BlogAuthorLanguageJaJp   BlogAuthorLanguage = "ja-jp"
	BlogAuthorLanguageJgo    BlogAuthorLanguage = "jgo"
	BlogAuthorLanguageJgoCm  BlogAuthorLanguage = "jgo-cm"
	BlogAuthorLanguageJmc    BlogAuthorLanguage = "jmc"
	BlogAuthorLanguageJmcTz  BlogAuthorLanguage = "jmc-tz"
	BlogAuthorLanguageJv     BlogAuthorLanguage = "jv"
	BlogAuthorLanguageJvID   BlogAuthorLanguage = "jv-id"
	BlogAuthorLanguageKa     BlogAuthorLanguage = "ka"
	BlogAuthorLanguageKaGe   BlogAuthorLanguage = "ka-ge"
	BlogAuthorLanguageKab    BlogAuthorLanguage = "kab"
	BlogAuthorLanguageKabDz  BlogAuthorLanguage = "kab-dz"
	BlogAuthorLanguageKam    BlogAuthorLanguage = "kam"
	BlogAuthorLanguageKamKe  BlogAuthorLanguage = "kam-ke"
	BlogAuthorLanguageKde    BlogAuthorLanguage = "kde"
	BlogAuthorLanguageKdeTz  BlogAuthorLanguage = "kde-tz"
	BlogAuthorLanguageKea    BlogAuthorLanguage = "kea"
	BlogAuthorLanguageKeaCv  BlogAuthorLanguage = "kea-cv"
	BlogAuthorLanguageKhq    BlogAuthorLanguage = "khq"
	BlogAuthorLanguageKhqMl  BlogAuthorLanguage = "khq-ml"
	BlogAuthorLanguageKi     BlogAuthorLanguage = "ki"
	BlogAuthorLanguageKiKe   BlogAuthorLanguage = "ki-ke"
	BlogAuthorLanguageKk     BlogAuthorLanguage = "kk"
	BlogAuthorLanguageKkKz   BlogAuthorLanguage = "kk-kz"
	BlogAuthorLanguageKkj    BlogAuthorLanguage = "kkj"
	BlogAuthorLanguageKkjCm  BlogAuthorLanguage = "kkj-cm"
	BlogAuthorLanguageKl     BlogAuthorLanguage = "kl"
	BlogAuthorLanguageKlGl   BlogAuthorLanguage = "kl-gl"
	BlogAuthorLanguageKln    BlogAuthorLanguage = "kln"
	BlogAuthorLanguageKlnKe  BlogAuthorLanguage = "kln-ke"
	BlogAuthorLanguageKm     BlogAuthorLanguage = "km"
	BlogAuthorLanguageKmKh   BlogAuthorLanguage = "km-kh"
	BlogAuthorLanguageKn     BlogAuthorLanguage = "kn"
	BlogAuthorLanguageKnIn   BlogAuthorLanguage = "kn-in"
	BlogAuthorLanguageKo     BlogAuthorLanguage = "ko"
	BlogAuthorLanguageKoKp   BlogAuthorLanguage = "ko-kp"
	BlogAuthorLanguageKoKr   BlogAuthorLanguage = "ko-kr"
	BlogAuthorLanguageKok    BlogAuthorLanguage = "kok"
	BlogAuthorLanguageKokIn  BlogAuthorLanguage = "kok-in"
	BlogAuthorLanguageKs     BlogAuthorLanguage = "ks"
	BlogAuthorLanguageKsIn   BlogAuthorLanguage = "ks-in"
	BlogAuthorLanguageKsb    BlogAuthorLanguage = "ksb"
	BlogAuthorLanguageKsbTz  BlogAuthorLanguage = "ksb-tz"
	BlogAuthorLanguageKsf    BlogAuthorLanguage = "ksf"
	BlogAuthorLanguageKsfCm  BlogAuthorLanguage = "ksf-cm"
	BlogAuthorLanguageKsh    BlogAuthorLanguage = "ksh"
	BlogAuthorLanguageKshDe  BlogAuthorLanguage = "ksh-de"
	BlogAuthorLanguageKu     BlogAuthorLanguage = "ku"
	BlogAuthorLanguageKuTr   BlogAuthorLanguage = "ku-tr"
	BlogAuthorLanguageKw     BlogAuthorLanguage = "kw"
	BlogAuthorLanguageKwGB   BlogAuthorLanguage = "kw-gb"
	BlogAuthorLanguageKy     BlogAuthorLanguage = "ky"
	BlogAuthorLanguageKyKg   BlogAuthorLanguage = "ky-kg"
	BlogAuthorLanguageLag    BlogAuthorLanguage = "lag"
	BlogAuthorLanguageLagTz  BlogAuthorLanguage = "lag-tz"
	BlogAuthorLanguageLb     BlogAuthorLanguage = "lb"
	BlogAuthorLanguageLbLu   BlogAuthorLanguage = "lb-lu"
	BlogAuthorLanguageLg     BlogAuthorLanguage = "lg"
	BlogAuthorLanguageLgUg   BlogAuthorLanguage = "lg-ug"
	BlogAuthorLanguageLkt    BlogAuthorLanguage = "lkt"
	BlogAuthorLanguageLktUs  BlogAuthorLanguage = "lkt-us"
	BlogAuthorLanguageLn     BlogAuthorLanguage = "ln"
	BlogAuthorLanguageLnAo   BlogAuthorLanguage = "ln-ao"
	BlogAuthorLanguageLnCd   BlogAuthorLanguage = "ln-cd"
	BlogAuthorLanguageLnCf   BlogAuthorLanguage = "ln-cf"
	BlogAuthorLanguageLnCg   BlogAuthorLanguage = "ln-cg"
	BlogAuthorLanguageLo     BlogAuthorLanguage = "lo"
	BlogAuthorLanguageLoLa   BlogAuthorLanguage = "lo-la"
	BlogAuthorLanguageLrc    BlogAuthorLanguage = "lrc"
	BlogAuthorLanguageLrcIq  BlogAuthorLanguage = "lrc-iq"
	BlogAuthorLanguageLrcIr  BlogAuthorLanguage = "lrc-ir"
	BlogAuthorLanguageLt     BlogAuthorLanguage = "lt"
	BlogAuthorLanguageLtLt   BlogAuthorLanguage = "lt-lt"
	BlogAuthorLanguageLu     BlogAuthorLanguage = "lu"
	BlogAuthorLanguageLuCd   BlogAuthorLanguage = "lu-cd"
	BlogAuthorLanguageLuo    BlogAuthorLanguage = "luo"
	BlogAuthorLanguageLuoKe  BlogAuthorLanguage = "luo-ke"
	BlogAuthorLanguageLuy    BlogAuthorLanguage = "luy"
	BlogAuthorLanguageLuyKe  BlogAuthorLanguage = "luy-ke"
	BlogAuthorLanguageLv     BlogAuthorLanguage = "lv"
	BlogAuthorLanguageLvLv   BlogAuthorLanguage = "lv-lv"
	BlogAuthorLanguageMai    BlogAuthorLanguage = "mai"
	BlogAuthorLanguageMaiIn  BlogAuthorLanguage = "mai-in"
	BlogAuthorLanguageMas    BlogAuthorLanguage = "mas"
	BlogAuthorLanguageMasKe  BlogAuthorLanguage = "mas-ke"
	BlogAuthorLanguageMasTz  BlogAuthorLanguage = "mas-tz"
	BlogAuthorLanguageMer    BlogAuthorLanguage = "mer"
	BlogAuthorLanguageMerKe  BlogAuthorLanguage = "mer-ke"
	BlogAuthorLanguageMfe    BlogAuthorLanguage = "mfe"
	BlogAuthorLanguageMfeMu  BlogAuthorLanguage = "mfe-mu"
	BlogAuthorLanguageMg     BlogAuthorLanguage = "mg"
	BlogAuthorLanguageMgMg   BlogAuthorLanguage = "mg-mg"
	BlogAuthorLanguageMgh    BlogAuthorLanguage = "mgh"
	BlogAuthorLanguageMghMz  BlogAuthorLanguage = "mgh-mz"
	BlogAuthorLanguageMgo    BlogAuthorLanguage = "mgo"
	BlogAuthorLanguageMgoCm  BlogAuthorLanguage = "mgo-cm"
	BlogAuthorLanguageMi     BlogAuthorLanguage = "mi"
	BlogAuthorLanguageMiNz   BlogAuthorLanguage = "mi-nz"
	BlogAuthorLanguageMk     BlogAuthorLanguage = "mk"
	BlogAuthorLanguageMkMk   BlogAuthorLanguage = "mk-mk"
	BlogAuthorLanguageMl     BlogAuthorLanguage = "ml"
	BlogAuthorLanguageMlIn   BlogAuthorLanguage = "ml-in"
	BlogAuthorLanguageMn     BlogAuthorLanguage = "mn"
	BlogAuthorLanguageMnMn   BlogAuthorLanguage = "mn-mn"
	BlogAuthorLanguageMni    BlogAuthorLanguage = "mni"
	BlogAuthorLanguageMniIn  BlogAuthorLanguage = "mni-in"
	BlogAuthorLanguageMr     BlogAuthorLanguage = "mr"
	BlogAuthorLanguageMrIn   BlogAuthorLanguage = "mr-in"
	BlogAuthorLanguageMs     BlogAuthorLanguage = "ms"
	BlogAuthorLanguageMsBn   BlogAuthorLanguage = "ms-bn"
	BlogAuthorLanguageMsID   BlogAuthorLanguage = "ms-id"
	BlogAuthorLanguageMsMy   BlogAuthorLanguage = "ms-my"
	BlogAuthorLanguageMsSg   BlogAuthorLanguage = "ms-sg"
	BlogAuthorLanguageMt     BlogAuthorLanguage = "mt"
	BlogAuthorLanguageMtMt   BlogAuthorLanguage = "mt-mt"
	BlogAuthorLanguageMua    BlogAuthorLanguage = "mua"
	BlogAuthorLanguageMuaCm  BlogAuthorLanguage = "mua-cm"
	BlogAuthorLanguageMy     BlogAuthorLanguage = "my"
	BlogAuthorLanguageMyMm   BlogAuthorLanguage = "my-mm"
	BlogAuthorLanguageMzn    BlogAuthorLanguage = "mzn"
	BlogAuthorLanguageMznIr  BlogAuthorLanguage = "mzn-ir"
	BlogAuthorLanguageNaq    BlogAuthorLanguage = "naq"
	BlogAuthorLanguageNaqNa  BlogAuthorLanguage = "naq-na"
	BlogAuthorLanguageNb     BlogAuthorLanguage = "nb"
	BlogAuthorLanguageNbNo   BlogAuthorLanguage = "nb-no"
	BlogAuthorLanguageNbSj   BlogAuthorLanguage = "nb-sj"
	BlogAuthorLanguageNd     BlogAuthorLanguage = "nd"
	BlogAuthorLanguageNdZw   BlogAuthorLanguage = "nd-zw"
	BlogAuthorLanguageNds    BlogAuthorLanguage = "nds"
	BlogAuthorLanguageNdsDe  BlogAuthorLanguage = "nds-de"
	BlogAuthorLanguageNdsNl  BlogAuthorLanguage = "nds-nl"
	BlogAuthorLanguageNe     BlogAuthorLanguage = "ne"
	BlogAuthorLanguageNeIn   BlogAuthorLanguage = "ne-in"
	BlogAuthorLanguageNeNp   BlogAuthorLanguage = "ne-np"
	BlogAuthorLanguageNl     BlogAuthorLanguage = "nl"
	BlogAuthorLanguageNlAw   BlogAuthorLanguage = "nl-aw"
	BlogAuthorLanguageNlBe   BlogAuthorLanguage = "nl-be"
	BlogAuthorLanguageNlBq   BlogAuthorLanguage = "nl-bq"
	BlogAuthorLanguageNlCh   BlogAuthorLanguage = "nl-ch"
	BlogAuthorLanguageNlCw   BlogAuthorLanguage = "nl-cw"
	BlogAuthorLanguageNlLu   BlogAuthorLanguage = "nl-lu"
	BlogAuthorLanguageNlNl   BlogAuthorLanguage = "nl-nl"
	BlogAuthorLanguageNlSr   BlogAuthorLanguage = "nl-sr"
	BlogAuthorLanguageNlSx   BlogAuthorLanguage = "nl-sx"
	BlogAuthorLanguageNmg    BlogAuthorLanguage = "nmg"
	BlogAuthorLanguageNmgCm  BlogAuthorLanguage = "nmg-cm"
	BlogAuthorLanguageNn     BlogAuthorLanguage = "nn"
	BlogAuthorLanguageNnNo   BlogAuthorLanguage = "nn-no"
	BlogAuthorLanguageNnh    BlogAuthorLanguage = "nnh"
	BlogAuthorLanguageNnhCm  BlogAuthorLanguage = "nnh-cm"
	BlogAuthorLanguageNo     BlogAuthorLanguage = "no"
	BlogAuthorLanguageNoNo   BlogAuthorLanguage = "no-no"
	BlogAuthorLanguageNus    BlogAuthorLanguage = "nus"
	BlogAuthorLanguageNusSS  BlogAuthorLanguage = "nus-ss"
	BlogAuthorLanguageNyn    BlogAuthorLanguage = "nyn"
	BlogAuthorLanguageNynUg  BlogAuthorLanguage = "nyn-ug"
	BlogAuthorLanguageOm     BlogAuthorLanguage = "om"
	BlogAuthorLanguageOmEt   BlogAuthorLanguage = "om-et"
	BlogAuthorLanguageOmKe   BlogAuthorLanguage = "om-ke"
	BlogAuthorLanguageOr     BlogAuthorLanguage = "or"
	BlogAuthorLanguageOrIn   BlogAuthorLanguage = "or-in"
	BlogAuthorLanguageOs     BlogAuthorLanguage = "os"
	BlogAuthorLanguageOsGe   BlogAuthorLanguage = "os-ge"
	BlogAuthorLanguageOsRu   BlogAuthorLanguage = "os-ru"
	BlogAuthorLanguagePa     BlogAuthorLanguage = "pa"
	BlogAuthorLanguagePaIn   BlogAuthorLanguage = "pa-in"
	BlogAuthorLanguagePaPk   BlogAuthorLanguage = "pa-pk"
	BlogAuthorLanguagePcm    BlogAuthorLanguage = "pcm"
	BlogAuthorLanguagePcmNg  BlogAuthorLanguage = "pcm-ng"
	BlogAuthorLanguagePl     BlogAuthorLanguage = "pl"
	BlogAuthorLanguagePlPl   BlogAuthorLanguage = "pl-pl"
	BlogAuthorLanguagePrg    BlogAuthorLanguage = "prg"
	BlogAuthorLanguagePrg001 BlogAuthorLanguage = "prg-001"
	BlogAuthorLanguagePs     BlogAuthorLanguage = "ps"
	BlogAuthorLanguagePsAf   BlogAuthorLanguage = "ps-af"
	BlogAuthorLanguagePsPk   BlogAuthorLanguage = "ps-pk"
	BlogAuthorLanguagePt     BlogAuthorLanguage = "pt"
	BlogAuthorLanguagePtAo   BlogAuthorLanguage = "pt-ao"
	BlogAuthorLanguagePtBr   BlogAuthorLanguage = "pt-br"
	BlogAuthorLanguagePtCh   BlogAuthorLanguage = "pt-ch"
	BlogAuthorLanguagePtCv   BlogAuthorLanguage = "pt-cv"
	BlogAuthorLanguagePtGq   BlogAuthorLanguage = "pt-gq"
	BlogAuthorLanguagePtGw   BlogAuthorLanguage = "pt-gw"
	BlogAuthorLanguagePtLu   BlogAuthorLanguage = "pt-lu"
	BlogAuthorLanguagePtMo   BlogAuthorLanguage = "pt-mo"
	BlogAuthorLanguagePtMz   BlogAuthorLanguage = "pt-mz"
	BlogAuthorLanguagePtPt   BlogAuthorLanguage = "pt-pt"
	BlogAuthorLanguagePtSt   BlogAuthorLanguage = "pt-st"
	BlogAuthorLanguagePtTl   BlogAuthorLanguage = "pt-tl"
	BlogAuthorLanguageQu     BlogAuthorLanguage = "qu"
	BlogAuthorLanguageQuBo   BlogAuthorLanguage = "qu-bo"
	BlogAuthorLanguageQuEc   BlogAuthorLanguage = "qu-ec"
	BlogAuthorLanguageQuPe   BlogAuthorLanguage = "qu-pe"
	BlogAuthorLanguageRm     BlogAuthorLanguage = "rm"
	BlogAuthorLanguageRmCh   BlogAuthorLanguage = "rm-ch"
	BlogAuthorLanguageRn     BlogAuthorLanguage = "rn"
	BlogAuthorLanguageRnBi   BlogAuthorLanguage = "rn-bi"
	BlogAuthorLanguageRo     BlogAuthorLanguage = "ro"
	BlogAuthorLanguageRoMd   BlogAuthorLanguage = "ro-md"
	BlogAuthorLanguageRoRo   BlogAuthorLanguage = "ro-ro"
	BlogAuthorLanguageRof    BlogAuthorLanguage = "rof"
	BlogAuthorLanguageRofTz  BlogAuthorLanguage = "rof-tz"
	BlogAuthorLanguageRu     BlogAuthorLanguage = "ru"
	BlogAuthorLanguageRuBy   BlogAuthorLanguage = "ru-by"
	BlogAuthorLanguageRuKg   BlogAuthorLanguage = "ru-kg"
	BlogAuthorLanguageRuKz   BlogAuthorLanguage = "ru-kz"
	BlogAuthorLanguageRuMd   BlogAuthorLanguage = "ru-md"
	BlogAuthorLanguageRuRu   BlogAuthorLanguage = "ru-ru"
	BlogAuthorLanguageRuUa   BlogAuthorLanguage = "ru-ua"
	BlogAuthorLanguageRw     BlogAuthorLanguage = "rw"
	BlogAuthorLanguageRwRw   BlogAuthorLanguage = "rw-rw"
	BlogAuthorLanguageRwk    BlogAuthorLanguage = "rwk"
	BlogAuthorLanguageRwkTz  BlogAuthorLanguage = "rwk-tz"
	BlogAuthorLanguageSa     BlogAuthorLanguage = "sa"
	BlogAuthorLanguageSaIn   BlogAuthorLanguage = "sa-in"
	BlogAuthorLanguageSah    BlogAuthorLanguage = "sah"
	BlogAuthorLanguageSahRu  BlogAuthorLanguage = "sah-ru"
	BlogAuthorLanguageSaq    BlogAuthorLanguage = "saq"
	BlogAuthorLanguageSaqKe  BlogAuthorLanguage = "saq-ke"
	BlogAuthorLanguageSat    BlogAuthorLanguage = "sat"
	BlogAuthorLanguageSatIn  BlogAuthorLanguage = "sat-in"
	BlogAuthorLanguageSbp    BlogAuthorLanguage = "sbp"
	BlogAuthorLanguageSbpTz  BlogAuthorLanguage = "sbp-tz"
	BlogAuthorLanguageSd     BlogAuthorLanguage = "sd"
	BlogAuthorLanguageSdIn   BlogAuthorLanguage = "sd-in"
	BlogAuthorLanguageSdPk   BlogAuthorLanguage = "sd-pk"
	BlogAuthorLanguageSe     BlogAuthorLanguage = "se"
	BlogAuthorLanguageSeFi   BlogAuthorLanguage = "se-fi"
	BlogAuthorLanguageSeNo   BlogAuthorLanguage = "se-no"
	BlogAuthorLanguageSeSe   BlogAuthorLanguage = "se-se"
	BlogAuthorLanguageSeh    BlogAuthorLanguage = "seh"
	BlogAuthorLanguageSehMz  BlogAuthorLanguage = "seh-mz"
	BlogAuthorLanguageSes    BlogAuthorLanguage = "ses"
	BlogAuthorLanguageSesMl  BlogAuthorLanguage = "ses-ml"
	BlogAuthorLanguageSg     BlogAuthorLanguage = "sg"
	BlogAuthorLanguageSgCf   BlogAuthorLanguage = "sg-cf"
	BlogAuthorLanguageShi    BlogAuthorLanguage = "shi"
	BlogAuthorLanguageShiMa  BlogAuthorLanguage = "shi-ma"
	BlogAuthorLanguageSi     BlogAuthorLanguage = "si"
	BlogAuthorLanguageSiLk   BlogAuthorLanguage = "si-lk"
	BlogAuthorLanguageSk     BlogAuthorLanguage = "sk"
	BlogAuthorLanguageSkSk   BlogAuthorLanguage = "sk-sk"
	BlogAuthorLanguageSl     BlogAuthorLanguage = "sl"
	BlogAuthorLanguageSlSi   BlogAuthorLanguage = "sl-si"
	BlogAuthorLanguageSmn    BlogAuthorLanguage = "smn"
	BlogAuthorLanguageSmnFi  BlogAuthorLanguage = "smn-fi"
	BlogAuthorLanguageSn     BlogAuthorLanguage = "sn"
	BlogAuthorLanguageSnZw   BlogAuthorLanguage = "sn-zw"
	BlogAuthorLanguageSo     BlogAuthorLanguage = "so"
	BlogAuthorLanguageSoDj   BlogAuthorLanguage = "so-dj"
	BlogAuthorLanguageSoEt   BlogAuthorLanguage = "so-et"
	BlogAuthorLanguageSoKe   BlogAuthorLanguage = "so-ke"
	BlogAuthorLanguageSoSo   BlogAuthorLanguage = "so-so"
	BlogAuthorLanguageSq     BlogAuthorLanguage = "sq"
	BlogAuthorLanguageSqAl   BlogAuthorLanguage = "sq-al"
	BlogAuthorLanguageSqMk   BlogAuthorLanguage = "sq-mk"
	BlogAuthorLanguageSqXk   BlogAuthorLanguage = "sq-xk"
	BlogAuthorLanguageSr     BlogAuthorLanguage = "sr"
	BlogAuthorLanguageSrBa   BlogAuthorLanguage = "sr-ba"
	BlogAuthorLanguageSrCs   BlogAuthorLanguage = "sr-cs"
	BlogAuthorLanguageSrMe   BlogAuthorLanguage = "sr-me"
	BlogAuthorLanguageSrRs   BlogAuthorLanguage = "sr-rs"
	BlogAuthorLanguageSrXk   BlogAuthorLanguage = "sr-xk"
	BlogAuthorLanguageSu     BlogAuthorLanguage = "su"
	BlogAuthorLanguageSuID   BlogAuthorLanguage = "su-id"
	BlogAuthorLanguageSv     BlogAuthorLanguage = "sv"
	BlogAuthorLanguageSvAx   BlogAuthorLanguage = "sv-ax"
	BlogAuthorLanguageSvFi   BlogAuthorLanguage = "sv-fi"
	BlogAuthorLanguageSvSe   BlogAuthorLanguage = "sv-se"
	BlogAuthorLanguageSw     BlogAuthorLanguage = "sw"
	BlogAuthorLanguageSwCd   BlogAuthorLanguage = "sw-cd"
	BlogAuthorLanguageSwKe   BlogAuthorLanguage = "sw-ke"
	BlogAuthorLanguageSwTz   BlogAuthorLanguage = "sw-tz"
	BlogAuthorLanguageSwUg   BlogAuthorLanguage = "sw-ug"
	BlogAuthorLanguageSy     BlogAuthorLanguage = "sy"
	BlogAuthorLanguageTa     BlogAuthorLanguage = "ta"
	BlogAuthorLanguageTaIn   BlogAuthorLanguage = "ta-in"
	BlogAuthorLanguageTaLk   BlogAuthorLanguage = "ta-lk"
	BlogAuthorLanguageTaMy   BlogAuthorLanguage = "ta-my"
	BlogAuthorLanguageTaSg   BlogAuthorLanguage = "ta-sg"
	BlogAuthorLanguageTe     BlogAuthorLanguage = "te"
	BlogAuthorLanguageTeIn   BlogAuthorLanguage = "te-in"
	BlogAuthorLanguageTeo    BlogAuthorLanguage = "teo"
	BlogAuthorLanguageTeoKe  BlogAuthorLanguage = "teo-ke"
	BlogAuthorLanguageTeoUg  BlogAuthorLanguage = "teo-ug"
	BlogAuthorLanguageTg     BlogAuthorLanguage = "tg"
	BlogAuthorLanguageTgTj   BlogAuthorLanguage = "tg-tj"
	BlogAuthorLanguageTh     BlogAuthorLanguage = "th"
	BlogAuthorLanguageThTh   BlogAuthorLanguage = "th-th"
	BlogAuthorLanguageTi     BlogAuthorLanguage = "ti"
	BlogAuthorLanguageTiEr   BlogAuthorLanguage = "ti-er"
	BlogAuthorLanguageTiEt   BlogAuthorLanguage = "ti-et"
	BlogAuthorLanguageTk     BlogAuthorLanguage = "tk"
	BlogAuthorLanguageTkTm   BlogAuthorLanguage = "tk-tm"
	BlogAuthorLanguageTl     BlogAuthorLanguage = "tl"
	BlogAuthorLanguageTo     BlogAuthorLanguage = "to"
	BlogAuthorLanguageToTo   BlogAuthorLanguage = "to-to"
	BlogAuthorLanguageTr     BlogAuthorLanguage = "tr"
	BlogAuthorLanguageTrCy   BlogAuthorLanguage = "tr-cy"
	BlogAuthorLanguageTrTr   BlogAuthorLanguage = "tr-tr"
	BlogAuthorLanguageTt     BlogAuthorLanguage = "tt"
	BlogAuthorLanguageTtRu   BlogAuthorLanguage = "tt-ru"
	BlogAuthorLanguageTwq    BlogAuthorLanguage = "twq"
	BlogAuthorLanguageTwqNe  BlogAuthorLanguage = "twq-ne"
	BlogAuthorLanguageTzm    BlogAuthorLanguage = "tzm"
	BlogAuthorLanguageTzmMa  BlogAuthorLanguage = "tzm-ma"
	BlogAuthorLanguageUg     BlogAuthorLanguage = "ug"
	BlogAuthorLanguageUgCn   BlogAuthorLanguage = "ug-cn"
	BlogAuthorLanguageUk     BlogAuthorLanguage = "uk"
	BlogAuthorLanguageUkUa   BlogAuthorLanguage = "uk-ua"
	BlogAuthorLanguageUr     BlogAuthorLanguage = "ur"
	BlogAuthorLanguageUrIn   BlogAuthorLanguage = "ur-in"
	BlogAuthorLanguageUrPk   BlogAuthorLanguage = "ur-pk"
	BlogAuthorLanguageUz     BlogAuthorLanguage = "uz"
	BlogAuthorLanguageUzAf   BlogAuthorLanguage = "uz-af"
	BlogAuthorLanguageUzUz   BlogAuthorLanguage = "uz-uz"
	BlogAuthorLanguageVai    BlogAuthorLanguage = "vai"
	BlogAuthorLanguageVaiLr  BlogAuthorLanguage = "vai-lr"
	BlogAuthorLanguageVi     BlogAuthorLanguage = "vi"
	BlogAuthorLanguageViVn   BlogAuthorLanguage = "vi-vn"
	BlogAuthorLanguageVo     BlogAuthorLanguage = "vo"
	BlogAuthorLanguageVo001  BlogAuthorLanguage = "vo-001"
	BlogAuthorLanguageVun    BlogAuthorLanguage = "vun"
	BlogAuthorLanguageVunTz  BlogAuthorLanguage = "vun-tz"
	BlogAuthorLanguageWae    BlogAuthorLanguage = "wae"
	BlogAuthorLanguageWaeCh  BlogAuthorLanguage = "wae-ch"
	BlogAuthorLanguageWo     BlogAuthorLanguage = "wo"
	BlogAuthorLanguageWoSn   BlogAuthorLanguage = "wo-sn"
	BlogAuthorLanguageXh     BlogAuthorLanguage = "xh"
	BlogAuthorLanguageXhZa   BlogAuthorLanguage = "xh-za"
	BlogAuthorLanguageXog    BlogAuthorLanguage = "xog"
	BlogAuthorLanguageXogUg  BlogAuthorLanguage = "xog-ug"
	BlogAuthorLanguageYav    BlogAuthorLanguage = "yav"
	BlogAuthorLanguageYavCm  BlogAuthorLanguage = "yav-cm"
	BlogAuthorLanguageYi     BlogAuthorLanguage = "yi"
	BlogAuthorLanguageYi001  BlogAuthorLanguage = "yi-001"
	BlogAuthorLanguageYo     BlogAuthorLanguage = "yo"
	BlogAuthorLanguageYoBj   BlogAuthorLanguage = "yo-bj"
	BlogAuthorLanguageYoNg   BlogAuthorLanguage = "yo-ng"
	BlogAuthorLanguageYue    BlogAuthorLanguage = "yue"
	BlogAuthorLanguageYueCn  BlogAuthorLanguage = "yue-cn"
	BlogAuthorLanguageYueHk  BlogAuthorLanguage = "yue-hk"
	BlogAuthorLanguageZgh    BlogAuthorLanguage = "zgh"
	BlogAuthorLanguageZghMa  BlogAuthorLanguage = "zgh-ma"
	BlogAuthorLanguageZh     BlogAuthorLanguage = "zh"
	BlogAuthorLanguageZhCn   BlogAuthorLanguage = "zh-cn"
	BlogAuthorLanguageZhHans BlogAuthorLanguage = "zh-hans"
	BlogAuthorLanguageZhHant BlogAuthorLanguage = "zh-hant"
	BlogAuthorLanguageZhHk   BlogAuthorLanguage = "zh-hk"
	BlogAuthorLanguageZhMo   BlogAuthorLanguage = "zh-mo"
	BlogAuthorLanguageZhSg   BlogAuthorLanguage = "zh-sg"
	BlogAuthorLanguageZhTw   BlogAuthorLanguage = "zh-tw"
	BlogAuthorLanguageZu     BlogAuthorLanguage = "zu"
	BlogAuthorLanguageZuZa   BlogAuthorLanguage = "zu-za"
)

// Model definition for a Blog Author.
//
// The properties ID, Avatar, Bio, Created, DeletedAt, DisplayName, Email,
// Facebook, FullName, Language, Linkedin, Name, Slug, TranslatedFromID, Twitter,
// Updated, Website are required.
type BlogAuthorParam struct {
	// The unique ID of the Blog Author.
	ID string `json:"id,required"`
	// URL to the blog author's avatar, if supplying a custom one.
	Avatar string `json:"avatar,required"`
	// A short biography of the blog author.
	Bio     string    `json:"bio,required"`
	Created time.Time `json:"created,required" format:"date-time"`
	// The timestamp (ISO8601 format) when this Blog Author was deleted.
	DeletedAt time.Time `json:"deletedAt,required" format:"date-time"`
	// The full name of the Blog Author to be displayed.
	DisplayName string `json:"displayName,required"`
	// Email address of the Blog Author.
	Email string `json:"email,required"`
	// URL to the Blog Author's Facebook page.
	Facebook string `json:"facebook,required"`
	FullName string `json:"fullName,required"`
	// The explicitly defined ISO 639 language code of the blog author.
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
	// "en-dm", "en-er", "en-fi", "en-fj", "en-fk", "en-fm", "en-gb", "en-gd", "en-gg",
	// "en-gh", "en-gi", "en-gm", "en-gu", "en-gy", "en-hk", "en-ie", "en-il", "en-im",
	// "en-in", "en-io", "en-je", "en-jm", "en-ke", "en-ki", "en-kn", "en-ky", "en-lc",
	// "en-lr", "en-ls", "en-lu", "en-mg", "en-mh", "en-mo", "en-mp", "en-ms", "en-mt",
	// "en-mu", "en-mw", "en-mx", "en-my", "en-na", "en-nf", "en-ng", "en-nl", "en-nr",
	// "en-nu", "en-nz", "en-pg", "en-ph", "en-pk", "en-pn", "en-pr", "en-pw", "en-rw",
	// "en-sb", "en-sc", "en-sd", "en-se", "en-sg", "en-sh", "en-si", "en-sl", "en-ss",
	// "en-sx", "en-sz", "en-tc", "en-tk", "en-to", "en-tt", "en-tv", "en-tz", "en-ug",
	// "en-um", "en-us", "en-vc", "en-vg", "en-vi", "en-vu", "en-ws", "en-za", "en-zm",
	// "en-zw", "eo", "eo-001", "es", "es-419", "es-ar", "es-bo", "es-br", "es-bz",
	// "es-cl", "es-co", "es-cr", "es-cu", "es-do", "es-ea", "es-ec", "es-es", "es-gq",
	// "es-gt", "es-hn", "es-ic", "es-mx", "es-ni", "es-pa", "es-pe", "es-ph", "es-pr",
	// "es-py", "es-sv", "es-us", "es-uy", "es-ve", "et", "et-ee", "eu", "eu-es",
	// "ewo", "ewo-cm", "fa", "fa-af", "fa-ir", "ff", "ff-bf", "ff-cm", "ff-gh",
	// "ff-gm", "ff-gn", "ff-gw", "ff-lr", "ff-mr", "ff-ne", "ff-ng", "ff-sl", "ff-sn",
	// "fi", "fi-fi", "fil", "fil-ph", "fo", "fo-dk", "fo-fo", "fr", "fr-be", "fr-bf",
	// "fr-bi", "fr-bj", "fr-bl", "fr-ca", "fr-cd", "fr-cf", "fr-cg", "fr-ch", "fr-ci",
	// "fr-cm", "fr-dj", "fr-dz", "fr-fr", "fr-ga", "fr-gf", "fr-gn", "fr-gp", "fr-gq",
	// "fr-ht", "fr-km", "fr-lu", "fr-ma", "fr-mc", "fr-mf", "fr-mg", "fr-ml", "fr-mq",
	// "fr-mr", "fr-mu", "fr-nc", "fr-ne", "fr-pf", "fr-pm", "fr-re", "fr-rw", "fr-sc",
	// "fr-sn", "fr-sy", "fr-td", "fr-tg", "fr-tn", "fr-vu", "fr-wf", "fr-yt", "fur",
	// "fur-it", "fy", "fy-nl", "ga", "ga-gb", "ga-ie", "gd", "gd-gb", "gl", "gl-es",
	// "gsw", "gsw-ch", "gsw-fr", "gsw-li", "gu", "gu-in", "guz", "guz-ke", "gv",
	// "gv-im", "ha", "ha-gh", "ha-ne", "ha-ng", "haw", "haw-us", "he", "he-il", "hi",
	// "hi-in", "hr", "hr-ba", "hr-hr", "hsb", "hsb-de", "hu", "hu-hu", "hy", "hy-am",
	// "ia", "ia-001", "id", "id-id", "ig", "ig-ng", "ii", "ii-cn", "is", "is-is",
	// "it", "it-ch", "it-it", "it-sm", "it-va", "ja", "ja-jp", "jgo", "jgo-cm", "jmc",
	// "jmc-tz", "jv", "jv-id", "ka", "ka-ge", "kab", "kab-dz", "kam", "kam-ke", "kde",
	// "kde-tz", "kea", "kea-cv", "khq", "khq-ml", "ki", "ki-ke", "kk", "kk-kz", "kkj",
	// "kkj-cm", "kl", "kl-gl", "kln", "kln-ke", "km", "km-kh", "kn", "kn-in", "ko",
	// "ko-kp", "ko-kr", "kok", "kok-in", "ks", "ks-in", "ksb", "ksb-tz", "ksf",
	// "ksf-cm", "ksh", "ksh-de", "ku", "ku-tr", "kw", "kw-gb", "ky", "ky-kg", "lag",
	// "lag-tz", "lb", "lb-lu", "lg", "lg-ug", "lkt", "lkt-us", "ln", "ln-ao", "ln-cd",
	// "ln-cf", "ln-cg", "lo", "lo-la", "lrc", "lrc-iq", "lrc-ir", "lt", "lt-lt", "lu",
	// "lu-cd", "luo", "luo-ke", "luy", "luy-ke", "lv", "lv-lv", "mai", "mai-in",
	// "mas", "mas-ke", "mas-tz", "mer", "mer-ke", "mfe", "mfe-mu", "mg", "mg-mg",
	// "mgh", "mgh-mz", "mgo", "mgo-cm", "mi", "mi-nz", "mk", "mk-mk", "ml", "ml-in",
	// "mn", "mn-mn", "mni", "mni-in", "mr", "mr-in", "ms", "ms-bn", "ms-id", "ms-my",
	// "ms-sg", "mt", "mt-mt", "mua", "mua-cm", "my", "my-mm", "mzn", "mzn-ir", "naq",
	// "naq-na", "nb", "nb-no", "nb-sj", "nd", "nd-zw", "nds", "nds-de", "nds-nl",
	// "ne", "ne-in", "ne-np", "nl", "nl-aw", "nl-be", "nl-bq", "nl-ch", "nl-cw",
	// "nl-lu", "nl-nl", "nl-sr", "nl-sx", "nmg", "nmg-cm", "nn", "nn-no", "nnh",
	// "nnh-cm", "no", "no-no", "nus", "nus-ss", "nyn", "nyn-ug", "om", "om-et",
	// "om-ke", "or", "or-in", "os", "os-ge", "os-ru", "pa", "pa-in", "pa-pk", "pcm",
	// "pcm-ng", "pl", "pl-pl", "prg", "prg-001", "ps", "ps-af", "ps-pk", "pt",
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
	Language BlogAuthorLanguage `json:"language,omitzero,required"`
	// URL to the blog author's LinkedIn page.
	Linkedin string `json:"linkedin,required"`
	Name     string `json:"name,required"`
	Slug     string `json:"slug,required"`
	// ID of the primary blog author this object was translated from.
	TranslatedFromID int64 `json:"translatedFromId,required"`
	// URL or username of the Twitter account associated with the Blog Author. This
	// will be normalized into the Twitter url for said user.
	Twitter string    `json:"twitter,required"`
	Updated time.Time `json:"updated,required" format:"date-time"`
	// URL to the website of the Blog Author.
	Website string `json:"website,required"`
	paramObj
}

func (r BlogAuthorParam) MarshalJSON() (data []byte, err error) {
	type shadow BlogAuthorParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BlogAuthorParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Request body object for cloning blog authors.
//
// The properties ID, BlogAuthor are required.
type BlogAuthorCloneRequestVNextParam struct {
	// ID of the object to be cloned.
	ID string `json:"id,required"`
	// Model definition for a Blog Author.
	BlogAuthor BlogAuthorParam `json:"blogAuthor,omitzero,required"`
	// Language of newly cloned object.
	Language param.Opt[string] `json:"language,omitzero"`
	// Primary language in multi-language group.
	PrimaryLanguage param.Opt[string] `json:"primaryLanguage,omitzero"`
	paramObj
}

func (r BlogAuthorCloneRequestVNextParam) MarshalJSON() (data []byte, err error) {
	type shadow BlogAuthorCloneRequestVNextParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BlogAuthorCloneRequestVNextParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Response object for collections of blog authors with pagination information.
type CollectionResponseWithTotalBlogAuthorForwardPaging struct {
	// Collection of blog authors.
	Results []BlogAuthor `json:"results,required"`
	// Total number of blog authors.
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
func (r CollectionResponseWithTotalBlogAuthorForwardPaging) RawJSON() string { return r.JSON.raw }
func (r *CollectionResponseWithTotalBlogAuthorForwardPaging) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BlogAuthorNewParams struct {
	// Model definition for a Blog Author.
	BlogAuthor BlogAuthorParam
	paramObj
}

func (r BlogAuthorNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.BlogAuthor)
}
func (r *BlogAuthorNewParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.BlogAuthor)
}

type BlogAuthorUpdateParams struct {
	// Model definition for a Blog Author.
	BlogAuthor BlogAuthorParam
	// Specifies whether to update deleted Blog Authors. Defaults to `false`.
	Archived param.Opt[bool] `query:"archived,omitzero" json:"-"`
	paramObj
}

func (r BlogAuthorUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.BlogAuthor)
}
func (r *BlogAuthorUpdateParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.BlogAuthor)
}

// URLQuery serializes [BlogAuthorUpdateParams]'s query parameters as `url.Values`.
func (r BlogAuthorUpdateParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type BlogAuthorListParams struct {
	// The cursor token value to get the next set of results. You can get this from the
	// `paging.next.after` JSON property of a paged response containing more results.
	After param.Opt[string] `query:"after,omitzero" json:"-"`
	// Specifies whether to return deleted Blog Authors. Defaults to `false`.
	Archived param.Opt[bool] `query:"archived,omitzero" json:"-"`
	// Only return Blog Authors created after the specified time.
	CreatedAfter param.Opt[time.Time] `query:"createdAfter,omitzero" format:"date-time" json:"-"`
	// Only return Blog Authors created at exactly the specified time.
	CreatedAt param.Opt[time.Time] `query:"createdAt,omitzero" format:"date-time" json:"-"`
	// Only return Blog Authors created before the specified time.
	CreatedBefore param.Opt[time.Time] `query:"createdBefore,omitzero" format:"date-time" json:"-"`
	// The maximum number of results to return. Default is 100.
	Limit    param.Opt[int64]  `query:"limit,omitzero" json:"-"`
	Property param.Opt[string] `query:"property,omitzero" json:"-"`
	// Only return Blog Authors last updated after the specified time.
	UpdatedAfter param.Opt[time.Time] `query:"updatedAfter,omitzero" format:"date-time" json:"-"`
	// Only return Blog Authors last updated at exactly the specified time.
	UpdatedAt param.Opt[time.Time] `query:"updatedAt,omitzero" format:"date-time" json:"-"`
	// Only return Blog Authors last updated before the specified time.
	UpdatedBefore param.Opt[time.Time] `query:"updatedBefore,omitzero" format:"date-time" json:"-"`
	// Specifies which fields to use for sorting results. Valid fields are `name`,
	// `createdAt`, `updatedAt`, `createdBy`, `updatedBy`. `createdAt` will be used by
	// default.
	Sort []string `query:"sort,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [BlogAuthorListParams]'s query parameters as `url.Values`.
func (r BlogAuthorListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type BlogAuthorDeleteParams struct {
	// Whether to return only results that have been archived.
	Archived param.Opt[bool] `query:"archived,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [BlogAuthorDeleteParams]'s query parameters as `url.Values`.
func (r BlogAuthorDeleteParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type BlogAuthorAttachToLangGroupParams struct {
	// Request body object for attaching objects to multi-language groups.
	AttachToLangPrimaryRequestVNext AttachToLangPrimaryRequestVNextParam
	paramObj
}

func (r BlogAuthorAttachToLangGroupParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.AttachToLangPrimaryRequestVNext)
}
func (r *BlogAuthorAttachToLangGroupParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.AttachToLangPrimaryRequestVNext)
}

type BlogAuthorNewBatchParams struct {
	// Wrapper for providing an array of blog authors as inputs.
	BatchInputBlogAuthor BatchInputBlogAuthorParam
	paramObj
}

func (r BlogAuthorNewBatchParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.BatchInputBlogAuthor)
}
func (r *BlogAuthorNewBatchParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.BatchInputBlogAuthor)
}

type BlogAuthorNewLanguageVariationParams struct {
	// Request body object for cloning blog authors.
	BlogAuthorCloneRequestVNext BlogAuthorCloneRequestVNextParam
	paramObj
}

func (r BlogAuthorNewLanguageVariationParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.BlogAuthorCloneRequestVNext)
}
func (r *BlogAuthorNewLanguageVariationParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.BlogAuthorCloneRequestVNext)
}

type BlogAuthorDeleteBatchParams struct {
	// Wrapper for providing an array of strings as inputs.
	BatchInputString shared.BatchInputStringParam
	paramObj
}

func (r BlogAuthorDeleteBatchParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.BatchInputString)
}
func (r *BlogAuthorDeleteBatchParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.BatchInputString)
}

type BlogAuthorDetachFromLangGroupParams struct {
	// Request body object for detaching objects from multi-language groups.
	DetachFromLangGroupRequestVNext DetachFromLangGroupRequestVNextParam
	paramObj
}

func (r BlogAuthorDetachFromLangGroupParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.DetachFromLangGroupRequestVNext)
}
func (r *BlogAuthorDetachFromLangGroupParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.DetachFromLangGroupRequestVNext)
}

type BlogAuthorGetParams struct {
	// Specifies whether to return deleted Blog Authors. Defaults to `false`.
	Archived param.Opt[bool]   `query:"archived,omitzero" json:"-"`
	Property param.Opt[string] `query:"property,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [BlogAuthorGetParams]'s query parameters as `url.Values`.
func (r BlogAuthorGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type BlogAuthorGetBatchParams struct {
	// Wrapper for providing an array of strings as inputs.
	BatchInputString shared.BatchInputStringParam
	// Specifies whether to return deleted Blog Authors. Defaults to `false`.
	Archived param.Opt[bool] `query:"archived,omitzero" json:"-"`
	paramObj
}

func (r BlogAuthorGetBatchParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.BatchInputString)
}
func (r *BlogAuthorGetBatchParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.BatchInputString)
}

// URLQuery serializes [BlogAuthorGetBatchParams]'s query parameters as
// `url.Values`.
func (r BlogAuthorGetBatchParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type BlogAuthorSetNewLangPrimaryParams struct {
	// Request body object for setting a new primary language.
	SetNewLanguagePrimaryRequestVNext SetNewLanguagePrimaryRequestVNextParam
	paramObj
}

func (r BlogAuthorSetNewLangPrimaryParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.SetNewLanguagePrimaryRequestVNext)
}
func (r *BlogAuthorSetNewLangPrimaryParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.SetNewLanguagePrimaryRequestVNext)
}

type BlogAuthorUpdateBatchParams struct {
	// Wrapper for providing an array of JSON nodes as inputs.
	BatchInputJsonNode BatchInputJsonNodeParam
	// Specifies whether to update deleted Blog Authors. Defaults to `false`.
	Archived param.Opt[bool] `query:"archived,omitzero" json:"-"`
	paramObj
}

func (r BlogAuthorUpdateBatchParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.BatchInputJsonNode)
}
func (r *BlogAuthorUpdateBatchParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.BatchInputJsonNode)
}

// URLQuery serializes [BlogAuthorUpdateBatchParams]'s query parameters as
// `url.Values`.
func (r BlogAuthorUpdateBatchParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type BlogAuthorUpdateLanguagesParams struct {
	// Request object for updating languages within a multi-language group.
	UpdateLanguagesRequestVNext UpdateLanguagesRequestVNextParam
	paramObj
}

func (r BlogAuthorUpdateLanguagesParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.UpdateLanguagesRequestVNext)
}
func (r *BlogAuthorUpdateLanguagesParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.UpdateLanguagesRequestVNext)
}
