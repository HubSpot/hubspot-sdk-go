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

	"github.com/HubSpot/hubspot-sdk-go/internal/apijson"
	"github.com/HubSpot/hubspot-sdk-go/internal/apiquery"
	shimjson "github.com/HubSpot/hubspot-sdk-go/internal/encoding/json"
	"github.com/HubSpot/hubspot-sdk-go/internal/requestconfig"
	"github.com/HubSpot/hubspot-sdk-go/option"
	"github.com/HubSpot/hubspot-sdk-go/packages/param"
)

// BlogAuthorService contains methods and other services that help with interacting
// with the hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBlogAuthorService] method instead.
type BlogAuthorService struct {
	options []option.RequestOption
	Batch   BlogAuthorBatchService
}

// NewBlogAuthorService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewBlogAuthorService(opts ...option.RequestOption) (r BlogAuthorService) {
	r = BlogAuthorService{}
	r.options = opts
	r.Batch = NewBlogAuthorBatchService(opts...)
	return
}

// Create a new Blog Author.
func (r *BlogAuthorService) New(ctx context.Context, body BlogAuthorNewParams, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "cms/blogs/2026-03/authors"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Sparse updates a single Blog Author object identified by the id in the path. All
// the column values need not be specified. Only the that need to be modified can
// be specified.
func (r *BlogAuthorService) Update(ctx context.Context, objectID string, params BlogAuthorUpdateParams, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if objectID == "" {
		err = errors.New("missing required objectId parameter")
		return nil, err
	}
	path := fmt.Sprintf("cms/blogs/2026-03/authors/%s", url.PathEscape(objectID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return res, err
}

// Get the list of blog authors. Supports paging and filtering. This method would
// be useful for an integration that examined these models and used an external
// service to suggest edits.
func (r *BlogAuthorService) List(ctx context.Context, query BlogAuthorListParams, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "cms/blogs/2026-03/authors"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Delete the Blog Author object identified by the id in the path.
func (r *BlogAuthorService) Delete(ctx context.Context, objectID string, body BlogAuthorDeleteParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if objectID == "" {
		err = errors.New("missing required objectId parameter")
		return err
	}
	path := fmt.Sprintf("cms/blogs/2026-03/authors/%s", url.PathEscape(objectID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, nil, opts...)
	return err
}

// Attach a Blog Author to a multi-language group.
func (r *BlogAuthorService) AttachToLangGroup(ctx context.Context, body BlogAuthorAttachToLangGroupParams, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "cms/blogs/2026-03/authors/multi-language/attach-to-lang-group"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Create a new language variation from an existing Blog Author.
func (r *BlogAuthorService) NewLanguageVariation(ctx context.Context, body BlogAuthorNewLanguageVariationParams, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "cms/blogs/2026-03/authors/multi-language/create-language-variation"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Detach a Blog Author from a multi-language group.
func (r *BlogAuthorService) DetachFromLangGroup(ctx context.Context, body BlogAuthorDetachFromLangGroupParams, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "cms/blogs/2026-03/authors/multi-language/detach-from-lang-group"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Retrieve the Blog Author object identified by the id in the path.
func (r *BlogAuthorService) Get(ctx context.Context, objectID string, query BlogAuthorGetParams, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if objectID == "" {
		err = errors.New("missing required objectId parameter")
		return nil, err
	}
	path := fmt.Sprintf("cms/blogs/2026-03/authors/%s", url.PathEscape(objectID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

func (r *BlogAuthorService) GetCursor(ctx context.Context, query BlogAuthorGetCursorParams, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "cms/blogs/2026-03/authors/cursor"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

func (r *BlogAuthorService) GetCursorByQuery(ctx context.Context, query BlogAuthorGetCursorByQueryParams, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "cms/blogs/2026-03/authors/cursor/query"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

func (r *BlogAuthorService) GetPostsCursor(ctx context.Context, query BlogAuthorGetPostsCursorParams, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "cms/blogs/2026-03/posts/cursor"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

func (r *BlogAuthorService) GetPostsCursorByQuery(ctx context.Context, query BlogAuthorGetPostsCursorByQueryParams, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "cms/blogs/2026-03/posts/cursor/query"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

func (r *BlogAuthorService) GetTagsCursor(ctx context.Context, query BlogAuthorGetTagsCursorParams, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "cms/blogs/2026-03/tags/cursor"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

func (r *BlogAuthorService) GetTagsCursorByQuery(ctx context.Context, query BlogAuthorGetTagsCursorByQueryParams, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "cms/blogs/2026-03/tags/cursor/query"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Set a Blog Author as the primary language of a multi-language group.
func (r *BlogAuthorService) SetNewLangPrimary(ctx context.Context, body BlogAuthorSetNewLangPrimaryParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "cms/blogs/2026-03/authors/multi-language/set-new-lang-primary"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, nil, opts...)
	return err
}

// Explicitly set new languages for each Blog Author in a multi-language group.
func (r *BlogAuthorService) UpdateLanguages(ctx context.Context, body BlogAuthorUpdateLanguagesParams, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "cms/blogs/2026-03/authors/multi-language/update-languages"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// The property Inputs is required.
type BatchInputBlogAuthorParam struct {
	// Blog authors to input.
	Inputs []BlogAuthorParam `json:"inputs,omitzero" api:"required"`
	paramObj
}

func (r BatchInputBlogAuthorParam) MarshalJSON() (data []byte, err error) {
	type shadow BatchInputBlogAuthorParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BatchInputBlogAuthorParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties ID, Avatar, Bio, Created, DeletedAt, DisplayName, Email,
// Facebook, FullName, Language, Linkedin, Name, Slug, TranslatedFromID, Twitter,
// Updated, Website are required.
type BlogAuthorParam struct {
	// The unique ID of the Blog Author.
	ID string `json:"id" api:"required"`
	// URL to the blog author's avatar, if supplying a custom one.
	Avatar string `json:"avatar" api:"required"`
	// A short biography of the blog author.
	Bio string `json:"bio" api:"required"`
	// The timestamp (ISO8601 format) when this Blog Author was created.
	Created time.Time `json:"created" api:"required" format:"date-time"`
	// The timestamp (ISO8601 format) when this Blog Author was deleted.
	DeletedAt time.Time `json:"deletedAt" api:"required" format:"date-time"`
	// The full name of the Blog Author to be displayed.
	DisplayName string `json:"displayName" api:"required"`
	// Email address of the Blog Author.
	Email string `json:"email" api:"required"`
	// URL to the Blog Author's Facebook page.
	Facebook string `json:"facebook" api:"required"`
	// The full, unabbreviated name of the blog author, typically their first and last
	// name combined.
	FullName string `json:"fullName" api:"required"`
	// The explicitly defined ISO 639 language code of the blog author.
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
	Language BlogAuthorLanguage `json:"language,omitzero" api:"required"`
	// URL to the blog author's LinkedIn page.
	Linkedin string `json:"linkedin" api:"required"`
	// The name field for the blog author. (This appears to be a shorter or alternative
	// name field compared to fullName.)
	Name string `json:"name" api:"required"`
	// A URL-friendly identifier for the blog author that can be used to reference the
	// author in URLs. Typically generated from the author's name and contains
	// lowercase letters, hyphens, and underscores.
	Slug string `json:"slug" api:"required"`
	// ID of the primary blog author this object was translated from.
	TranslatedFromID int64 `json:"translatedFromId" api:"required"`
	// URL or username of the Twitter account associated with the Blog Author. This
	// will be normalized into the Twitter url for said user.
	Twitter string `json:"twitter" api:"required"`
	// The timestamp (ISO8601 format) when this Blog Author was updated.
	Updated time.Time `json:"updated" api:"required" format:"date-time"`
	// URL to the website of the Blog Author.
	Website string `json:"website" api:"required"`
	paramObj
}

func (r BlogAuthorParam) MarshalJSON() (data []byte, err error) {
	type shadow BlogAuthorParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BlogAuthorParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The explicitly defined ISO 639 language code of the blog author.
type BlogAuthorLanguage string

const (
	BlogAuthorLanguageAa     BlogAuthorLanguage = "aa"
	BlogAuthorLanguageAb     BlogAuthorLanguage = "ab"
	BlogAuthorLanguageAe     BlogAuthorLanguage = "ae"
	BlogAuthorLanguageAf     BlogAuthorLanguage = "af"
	BlogAuthorLanguageAfNa   BlogAuthorLanguage = "af-na"
	BlogAuthorLanguageAfZa   BlogAuthorLanguage = "af-za"
	BlogAuthorLanguageAgq    BlogAuthorLanguage = "agq"
	BlogAuthorLanguageAgqCm  BlogAuthorLanguage = "agq-cm"
	BlogAuthorLanguageAk     BlogAuthorLanguage = "ak"
	BlogAuthorLanguageAkGh   BlogAuthorLanguage = "ak-gh"
	BlogAuthorLanguageAm     BlogAuthorLanguage = "am"
	BlogAuthorLanguageAmEt   BlogAuthorLanguage = "am-et"
	BlogAuthorLanguageAn     BlogAuthorLanguage = "an"
	BlogAuthorLanguageAnn    BlogAuthorLanguage = "ann"
	BlogAuthorLanguageAnnNg  BlogAuthorLanguage = "ann-ng"
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
	BlogAuthorLanguageAv     BlogAuthorLanguage = "av"
	BlogAuthorLanguageAy     BlogAuthorLanguage = "ay"
	BlogAuthorLanguageAz     BlogAuthorLanguage = "az"
	BlogAuthorLanguageAzAz   BlogAuthorLanguage = "az-az"
	BlogAuthorLanguageBa     BlogAuthorLanguage = "ba"
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
	BlogAuthorLanguageBgc    BlogAuthorLanguage = "bgc"
	BlogAuthorLanguageBgcIn  BlogAuthorLanguage = "bgc-in"
	BlogAuthorLanguageBho    BlogAuthorLanguage = "bho"
	BlogAuthorLanguageBhoIn  BlogAuthorLanguage = "bho-in"
	BlogAuthorLanguageBi     BlogAuthorLanguage = "bi"
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
	BlogAuthorLanguageCh     BlogAuthorLanguage = "ch"
	BlogAuthorLanguageChr    BlogAuthorLanguage = "chr"
	BlogAuthorLanguageChrUs  BlogAuthorLanguage = "chr-us"
	BlogAuthorLanguageCkb    BlogAuthorLanguage = "ckb"
	BlogAuthorLanguageCkbIq  BlogAuthorLanguage = "ckb-iq"
	BlogAuthorLanguageCkbIr  BlogAuthorLanguage = "ckb-ir"
	BlogAuthorLanguageCo     BlogAuthorLanguage = "co"
	BlogAuthorLanguageCr     BlogAuthorLanguage = "cr"
	BlogAuthorLanguageCs     BlogAuthorLanguage = "cs"
	BlogAuthorLanguageCsCz   BlogAuthorLanguage = "cs-cz"
	BlogAuthorLanguageCu     BlogAuthorLanguage = "cu"
	BlogAuthorLanguageCuRu   BlogAuthorLanguage = "cu-ru"
	BlogAuthorLanguageCv     BlogAuthorLanguage = "cv"
	BlogAuthorLanguageCvRu   BlogAuthorLanguage = "cv-ru"
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
	BlogAuthorLanguageDv     BlogAuthorLanguage = "dv"
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
	BlogAuthorLanguageEnEe   BlogAuthorLanguage = "en-ee"
	BlogAuthorLanguageEnEg   BlogAuthorLanguage = "en-eg"
	BlogAuthorLanguageEnEr   BlogAuthorLanguage = "en-er"
	BlogAuthorLanguageEnEs   BlogAuthorLanguage = "en-es"
	BlogAuthorLanguageEnFi   BlogAuthorLanguage = "en-fi"
	BlogAuthorLanguageEnFj   BlogAuthorLanguage = "en-fj"
	BlogAuthorLanguageEnFk   BlogAuthorLanguage = "en-fk"
	BlogAuthorLanguageEnFm   BlogAuthorLanguage = "en-fm"
	BlogAuthorLanguageEnFr   BlogAuthorLanguage = "en-fr"
	BlogAuthorLanguageEnGB   BlogAuthorLanguage = "en-gb"
	BlogAuthorLanguageEnGd   BlogAuthorLanguage = "en-gd"
	BlogAuthorLanguageEnGg   BlogAuthorLanguage = "en-gg"
	BlogAuthorLanguageEnGh   BlogAuthorLanguage = "en-gh"
	BlogAuthorLanguageEnGi   BlogAuthorLanguage = "en-gi"
	BlogAuthorLanguageEnGm   BlogAuthorLanguage = "en-gm"
	BlogAuthorLanguageEnGu   BlogAuthorLanguage = "en-gu"
	BlogAuthorLanguageEnGy   BlogAuthorLanguage = "en-gy"
	BlogAuthorLanguageEnHk   BlogAuthorLanguage = "en-hk"
	BlogAuthorLanguageEnID   BlogAuthorLanguage = "en-id"
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
	BlogAuthorLanguageEnMv   BlogAuthorLanguage = "en-mv"
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
	BlogAuthorLanguageEnPt   BlogAuthorLanguage = "en-pt"
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
	BlogAuthorLanguageEnTh   BlogAuthorLanguage = "en-th"
	BlogAuthorLanguageEnTk   BlogAuthorLanguage = "en-tk"
	BlogAuthorLanguageEnTn   BlogAuthorLanguage = "en-tn"
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
	BlogAuthorLanguageEnVn   BlogAuthorLanguage = "en-vn"
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
	BlogAuthorLanguageFj     BlogAuthorLanguage = "fj"
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
	BlogAuthorLanguageFrr    BlogAuthorLanguage = "frr"
	BlogAuthorLanguageFrrDe  BlogAuthorLanguage = "frr-de"
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
	BlogAuthorLanguageGn     BlogAuthorLanguage = "gn"
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
	BlogAuthorLanguageHmn    BlogAuthorLanguage = "hmn"
	BlogAuthorLanguageHo     BlogAuthorLanguage = "ho"
	BlogAuthorLanguageHr     BlogAuthorLanguage = "hr"
	BlogAuthorLanguageHrBa   BlogAuthorLanguage = "hr-ba"
	BlogAuthorLanguageHrHr   BlogAuthorLanguage = "hr-hr"
	BlogAuthorLanguageHsb    BlogAuthorLanguage = "hsb"
	BlogAuthorLanguageHsbDe  BlogAuthorLanguage = "hsb-de"
	BlogAuthorLanguageHt     BlogAuthorLanguage = "ht"
	BlogAuthorLanguageHu     BlogAuthorLanguage = "hu"
	BlogAuthorLanguageHuHu   BlogAuthorLanguage = "hu-hu"
	BlogAuthorLanguageHy     BlogAuthorLanguage = "hy"
	BlogAuthorLanguageHyAm   BlogAuthorLanguage = "hy-am"
	BlogAuthorLanguageHz     BlogAuthorLanguage = "hz"
	BlogAuthorLanguageIa     BlogAuthorLanguage = "ia"
	BlogAuthorLanguageIa001  BlogAuthorLanguage = "ia-001"
	BlogAuthorLanguageID     BlogAuthorLanguage = "id"
	BlogAuthorLanguageIDID   BlogAuthorLanguage = "id-id"
	BlogAuthorLanguageIe     BlogAuthorLanguage = "ie"
	BlogAuthorLanguageIg     BlogAuthorLanguage = "ig"
	BlogAuthorLanguageIgNg   BlogAuthorLanguage = "ig-ng"
	BlogAuthorLanguageIi     BlogAuthorLanguage = "ii"
	BlogAuthorLanguageIiCn   BlogAuthorLanguage = "ii-cn"
	BlogAuthorLanguageIk     BlogAuthorLanguage = "ik"
	BlogAuthorLanguageIo     BlogAuthorLanguage = "io"
	BlogAuthorLanguageIs     BlogAuthorLanguage = "is"
	BlogAuthorLanguageIsIs   BlogAuthorLanguage = "is-is"
	BlogAuthorLanguageIt     BlogAuthorLanguage = "it"
	BlogAuthorLanguageItCh   BlogAuthorLanguage = "it-ch"
	BlogAuthorLanguageItIt   BlogAuthorLanguage = "it-it"
	BlogAuthorLanguageItSm   BlogAuthorLanguage = "it-sm"
	BlogAuthorLanguageItVa   BlogAuthorLanguage = "it-va"
	BlogAuthorLanguageIu     BlogAuthorLanguage = "iu"
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
	BlogAuthorLanguageKar    BlogAuthorLanguage = "kar"
	BlogAuthorLanguageKde    BlogAuthorLanguage = "kde"
	BlogAuthorLanguageKdeTz  BlogAuthorLanguage = "kde-tz"
	BlogAuthorLanguageKea    BlogAuthorLanguage = "kea"
	BlogAuthorLanguageKeaCv  BlogAuthorLanguage = "kea-cv"
	BlogAuthorLanguageKg     BlogAuthorLanguage = "kg"
	BlogAuthorLanguageKgp    BlogAuthorLanguage = "kgp"
	BlogAuthorLanguageKgpBr  BlogAuthorLanguage = "kgp-br"
	BlogAuthorLanguageKh     BlogAuthorLanguage = "kh"
	BlogAuthorLanguageKhq    BlogAuthorLanguage = "khq"
	BlogAuthorLanguageKhqMl  BlogAuthorLanguage = "khq-ml"
	BlogAuthorLanguageKi     BlogAuthorLanguage = "ki"
	BlogAuthorLanguageKiKe   BlogAuthorLanguage = "ki-ke"
	BlogAuthorLanguageKj     BlogAuthorLanguage = "kj"
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
	BlogAuthorLanguageKr     BlogAuthorLanguage = "kr"
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
	BlogAuthorLanguageKv     BlogAuthorLanguage = "kv"
	BlogAuthorLanguageKw     BlogAuthorLanguage = "kw"
	BlogAuthorLanguageKwGB   BlogAuthorLanguage = "kw-gb"
	BlogAuthorLanguageKy     BlogAuthorLanguage = "ky"
	BlogAuthorLanguageKyKg   BlogAuthorLanguage = "ky-kg"
	BlogAuthorLanguageLa     BlogAuthorLanguage = "la"
	BlogAuthorLanguageLag    BlogAuthorLanguage = "lag"
	BlogAuthorLanguageLagTz  BlogAuthorLanguage = "lag-tz"
	BlogAuthorLanguageLb     BlogAuthorLanguage = "lb"
	BlogAuthorLanguageLbLu   BlogAuthorLanguage = "lb-lu"
	BlogAuthorLanguageLg     BlogAuthorLanguage = "lg"
	BlogAuthorLanguageLgUg   BlogAuthorLanguage = "lg-ug"
	BlogAuthorLanguageLi     BlogAuthorLanguage = "li"
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
	BlogAuthorLanguageMdf    BlogAuthorLanguage = "mdf"
	BlogAuthorLanguageMdfRu  BlogAuthorLanguage = "mdf-ru"
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
	BlogAuthorLanguageMh     BlogAuthorLanguage = "mh"
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
	BlogAuthorLanguageNa     BlogAuthorLanguage = "na"
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
	BlogAuthorLanguageNg     BlogAuthorLanguage = "ng"
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
	BlogAuthorLanguageNr     BlogAuthorLanguage = "nr"
	BlogAuthorLanguageNus    BlogAuthorLanguage = "nus"
	BlogAuthorLanguageNusSS  BlogAuthorLanguage = "nus-ss"
	BlogAuthorLanguageNv     BlogAuthorLanguage = "nv"
	BlogAuthorLanguageNy     BlogAuthorLanguage = "ny"
	BlogAuthorLanguageNyn    BlogAuthorLanguage = "nyn"
	BlogAuthorLanguageNynUg  BlogAuthorLanguage = "nyn-ug"
	BlogAuthorLanguageOc     BlogAuthorLanguage = "oc"
	BlogAuthorLanguageOcEs   BlogAuthorLanguage = "oc-es"
	BlogAuthorLanguageOcFr   BlogAuthorLanguage = "oc-fr"
	BlogAuthorLanguageOj     BlogAuthorLanguage = "oj"
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
	BlogAuthorLanguagePi     BlogAuthorLanguage = "pi"
	BlogAuthorLanguagePis    BlogAuthorLanguage = "pis"
	BlogAuthorLanguagePisSb  BlogAuthorLanguage = "pis-sb"
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
	BlogAuthorLanguageRaj    BlogAuthorLanguage = "raj"
	BlogAuthorLanguageRajIn  BlogAuthorLanguage = "raj-in"
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
	BlogAuthorLanguageSc     BlogAuthorLanguage = "sc"
	BlogAuthorLanguageScIt   BlogAuthorLanguage = "sc-it"
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
	BlogAuthorLanguageSm     BlogAuthorLanguage = "sm"
	BlogAuthorLanguageSmn    BlogAuthorLanguage = "smn"
	BlogAuthorLanguageSmnFi  BlogAuthorLanguage = "smn-fi"
	BlogAuthorLanguageSMS    BlogAuthorLanguage = "sms"
	BlogAuthorLanguageSMSFi  BlogAuthorLanguage = "sms-fi"
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
	BlogAuthorLanguageSS     BlogAuthorLanguage = "ss"
	BlogAuthorLanguageSt     BlogAuthorLanguage = "st"
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
	BlogAuthorLanguageTn     BlogAuthorLanguage = "tn"
	BlogAuthorLanguageTo     BlogAuthorLanguage = "to"
	BlogAuthorLanguageToTo   BlogAuthorLanguage = "to-to"
	BlogAuthorLanguageTok    BlogAuthorLanguage = "tok"
	BlogAuthorLanguageTok001 BlogAuthorLanguage = "tok-001"
	BlogAuthorLanguageTr     BlogAuthorLanguage = "tr"
	BlogAuthorLanguageTrCy   BlogAuthorLanguage = "tr-cy"
	BlogAuthorLanguageTrTr   BlogAuthorLanguage = "tr-tr"
	BlogAuthorLanguageTs     BlogAuthorLanguage = "ts"
	BlogAuthorLanguageTt     BlogAuthorLanguage = "tt"
	BlogAuthorLanguageTtRu   BlogAuthorLanguage = "tt-ru"
	BlogAuthorLanguageTw     BlogAuthorLanguage = "tw"
	BlogAuthorLanguageTwq    BlogAuthorLanguage = "twq"
	BlogAuthorLanguageTwqNe  BlogAuthorLanguage = "twq-ne"
	BlogAuthorLanguageTy     BlogAuthorLanguage = "ty"
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
	BlogAuthorLanguageVe     BlogAuthorLanguage = "ve"
	BlogAuthorLanguageVi     BlogAuthorLanguage = "vi"
	BlogAuthorLanguageViVn   BlogAuthorLanguage = "vi-vn"
	BlogAuthorLanguageVo     BlogAuthorLanguage = "vo"
	BlogAuthorLanguageVo001  BlogAuthorLanguage = "vo-001"
	BlogAuthorLanguageVun    BlogAuthorLanguage = "vun"
	BlogAuthorLanguageVunTz  BlogAuthorLanguage = "vun-tz"
	BlogAuthorLanguageWa     BlogAuthorLanguage = "wa"
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
	BlogAuthorLanguageYrl    BlogAuthorLanguage = "yrl"
	BlogAuthorLanguageYrlBr  BlogAuthorLanguage = "yrl-br"
	BlogAuthorLanguageYrlCo  BlogAuthorLanguage = "yrl-co"
	BlogAuthorLanguageYrlVe  BlogAuthorLanguage = "yrl-ve"
	BlogAuthorLanguageYue    BlogAuthorLanguage = "yue"
	BlogAuthorLanguageYueCn  BlogAuthorLanguage = "yue-cn"
	BlogAuthorLanguageYueHk  BlogAuthorLanguage = "yue-hk"
	BlogAuthorLanguageZa     BlogAuthorLanguage = "za"
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

// The properties ID, BlogAuthor are required.
type BlogAuthorCloneRequestVNextParam struct {
	// ID of the object to be cloned.
	ID         string          `json:"id" api:"required"`
	BlogAuthor BlogAuthorParam `json:"blogAuthor,omitzero" api:"required"`
	// Language of newly cloned object.
	Language param.Opt[string] `json:"language,omitzero"`
	// Primary language in multi-language group.
	PrimaryLanguage param.Opt[string] `json:"primaryLanguage,omitzero"`
	UsePublished    param.Opt[bool]   `json:"usePublished,omitzero"`
	paramObj
}

func (r BlogAuthorCloneRequestVNextParam) MarshalJSON() (data []byte, err error) {
	type shadow BlogAuthorCloneRequestVNextParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BlogAuthorCloneRequestVNextParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BlogAuthorNewParams struct {
	BlogAuthor BlogAuthorParam
	paramObj
}

func (r BlogAuthorNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.BlogAuthor)
}
func (r *BlogAuthorNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BlogAuthorUpdateParams struct {
	BlogAuthor BlogAuthorParam
	// Whether to return only results that have been archived.
	Archived param.Opt[bool] `query:"archived,omitzero" json:"-"`
	paramObj
}

func (r BlogAuthorUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.BlogAuthor)
}
func (r *BlogAuthorUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [BlogAuthorUpdateParams]'s query parameters as `url.Values`.
func (r BlogAuthorUpdateParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type BlogAuthorListParams struct {
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
	AttachToLangPrimaryRequestVNext AttachToLangPrimaryRequestVNextParam
	paramObj
}

func (r BlogAuthorAttachToLangGroupParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.AttachToLangPrimaryRequestVNext)
}
func (r *BlogAuthorAttachToLangGroupParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BlogAuthorNewLanguageVariationParams struct {
	BlogAuthorCloneRequestVNext BlogAuthorCloneRequestVNextParam
	paramObj
}

func (r BlogAuthorNewLanguageVariationParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.BlogAuthorCloneRequestVNext)
}
func (r *BlogAuthorNewLanguageVariationParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BlogAuthorDetachFromLangGroupParams struct {
	DetachFromLangGroupRequestVNext DetachFromLangGroupRequestVNextParam
	paramObj
}

func (r BlogAuthorDetachFromLangGroupParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.DetachFromLangGroupRequestVNext)
}
func (r *BlogAuthorDetachFromLangGroupParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BlogAuthorGetParams struct {
	// Whether to return only results that have been archived.
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

type BlogAuthorGetCursorParams struct {
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

// URLQuery serializes [BlogAuthorGetCursorParams]'s query parameters as
// `url.Values`.
func (r BlogAuthorGetCursorParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type BlogAuthorGetCursorByQueryParams struct {
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

// URLQuery serializes [BlogAuthorGetCursorByQueryParams]'s query parameters as
// `url.Values`.
func (r BlogAuthorGetCursorByQueryParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type BlogAuthorGetPostsCursorParams struct {
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

// URLQuery serializes [BlogAuthorGetPostsCursorParams]'s query parameters as
// `url.Values`.
func (r BlogAuthorGetPostsCursorParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type BlogAuthorGetPostsCursorByQueryParams struct {
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

// URLQuery serializes [BlogAuthorGetPostsCursorByQueryParams]'s query parameters
// as `url.Values`.
func (r BlogAuthorGetPostsCursorByQueryParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type BlogAuthorGetTagsCursorParams struct {
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

// URLQuery serializes [BlogAuthorGetTagsCursorParams]'s query parameters as
// `url.Values`.
func (r BlogAuthorGetTagsCursorParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type BlogAuthorGetTagsCursorByQueryParams struct {
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

// URLQuery serializes [BlogAuthorGetTagsCursorByQueryParams]'s query parameters as
// `url.Values`.
func (r BlogAuthorGetTagsCursorByQueryParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type BlogAuthorSetNewLangPrimaryParams struct {
	SetNewLanguagePrimaryRequestVNext SetNewLanguagePrimaryRequestVNextParam
	paramObj
}

func (r BlogAuthorSetNewLangPrimaryParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.SetNewLanguagePrimaryRequestVNext)
}
func (r *BlogAuthorSetNewLangPrimaryParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BlogAuthorUpdateLanguagesParams struct {
	UpdateLanguagesRequestVNext UpdateLanguagesRequestVNextParam
	paramObj
}

func (r BlogAuthorUpdateLanguagesParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.UpdateLanguagesRequestVNext)
}
func (r *BlogAuthorUpdateLanguagesParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
