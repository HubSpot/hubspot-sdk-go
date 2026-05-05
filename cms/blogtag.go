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

// BlogTagService contains methods and other services that help with interacting
// with the hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBlogTagService] method instead.
type BlogTagService struct {
	options []option.RequestOption
	Batch   BlogTagBatchService
}

// NewBlogTagService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewBlogTagService(opts ...option.RequestOption) (r BlogTagService) {
	r = BlogTagService{}
	r.options = opts
	r.Batch = NewBlogTagBatchService(opts...)
	return
}

// Create a new Blog Tag.
func (r *BlogTagService) New(ctx context.Context, body BlogTagNewParams, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "cms/blogs/2026-03/tags"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Sparse updates a single Blog Tag object identified by the id in the path. All
// the column values need not be specified. Only the that need to be modified can
// be specified.
func (r *BlogTagService) Update(ctx context.Context, objectID string, params BlogTagUpdateParams, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if objectID == "" {
		err = errors.New("missing required objectId parameter")
		return nil, err
	}
	path := fmt.Sprintf("cms/blogs/2026-03/tags/%s", url.PathEscape(objectID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return res, err
}

// Get the list of blog tags. Supports paging and filtering. This method would be
// useful for an integration that examined these models and used an external
// service to suggest edits.
func (r *BlogTagService) List(ctx context.Context, query BlogTagListParams, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "cms/blogs/2026-03/tags"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Delete the Blog Tag object identified by the id in the path.
func (r *BlogTagService) Delete(ctx context.Context, objectID string, body BlogTagDeleteParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if objectID == "" {
		err = errors.New("missing required objectId parameter")
		return err
	}
	path := fmt.Sprintf("cms/blogs/2026-03/tags/%s", url.PathEscape(objectID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, nil, opts...)
	return err
}

// Attach a Blog Tag to a multi-language group.
func (r *BlogTagService) AttachToLangGroup(ctx context.Context, body BlogTagAttachToLangGroupParams, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "cms/blogs/2026-03/tags/multi-language/attach-to-lang-group"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Create a new language variation from an existing Blog Tag
func (r *BlogTagService) NewLangVariation(ctx context.Context, body BlogTagNewLangVariationParams, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "cms/blogs/2026-03/tags/multi-language/create-language-variation"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Detach a Blog Tag from a multi-language group.
func (r *BlogTagService) DetachFromLangGroup(ctx context.Context, body BlogTagDetachFromLangGroupParams, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "cms/blogs/2026-03/tags/multi-language/detach-from-lang-group"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Retrieve the Blog Tag object identified by the id in the path.
func (r *BlogTagService) Get(ctx context.Context, objectID string, query BlogTagGetParams, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if objectID == "" {
		err = errors.New("missing required objectId parameter")
		return nil, err
	}
	path := fmt.Sprintf("cms/blogs/2026-03/tags/%s", url.PathEscape(objectID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

func (r *BlogTagService) ListAuthorsCursor(ctx context.Context, query BlogTagListAuthorsCursorParams, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "cms/blogs/2026-03/authors/cursor"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

func (r *BlogTagService) ListAuthorsCursorByQuery(ctx context.Context, query BlogTagListAuthorsCursorByQueryParams, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "cms/blogs/2026-03/authors/cursor/query"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

func (r *BlogTagService) ListCursor(ctx context.Context, query BlogTagListCursorParams, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "cms/blogs/2026-03/tags/cursor"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

func (r *BlogTagService) ListCursorByQuery(ctx context.Context, query BlogTagListCursorByQueryParams, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "cms/blogs/2026-03/tags/cursor/query"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

func (r *BlogTagService) ListPostsCursor(ctx context.Context, query BlogTagListPostsCursorParams, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "cms/blogs/2026-03/posts/cursor"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

func (r *BlogTagService) ListPostsCursorByQuery(ctx context.Context, query BlogTagListPostsCursorByQueryParams, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "cms/blogs/2026-03/posts/cursor/query"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Set a Blog Tag as the primary language of a multi-language group.
func (r *BlogTagService) SetLangPrimary(ctx context.Context, body BlogTagSetLangPrimaryParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "cms/blogs/2026-03/tags/multi-language/set-new-lang-primary"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, nil, opts...)
	return err
}

// Explicitly set new languages for each Blog Tag in a multi-language group.
func (r *BlogTagService) UpdateLangs(ctx context.Context, body BlogTagUpdateLangsParams, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "cms/blogs/2026-03/tags/multi-language/update-languages"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// The property Inputs is required.
type BatchInputTagParam struct {
	// Blog tags to input.
	Inputs []TagParam `json:"inputs,omitzero" api:"required"`
	paramObj
}

func (r BatchInputTagParam) MarshalJSON() (data []byte, err error) {
	type shadow BatchInputTagParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BatchInputTagParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties ID, Created, DeletedAt, Language, Name, Slug, TranslatedFromID,
// Updated are required.
type TagParam struct {
	// The unique ID of the Blog Tag.
	ID string `json:"id" api:"required"`
	// The timestamp (ISO8601 format) when this Blog Tag was created.
	Created time.Time `json:"created" api:"required" format:"date-time"`
	// The timestamp (ISO8601 format) when this Blog Tag was deleted.
	DeletedAt time.Time `json:"deletedAt" api:"required" format:"date-time"`
	// The explicitly defined ISO 639 language code of the tag.
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
	Language TagLanguage `json:"language,omitzero" api:"required"`
	// The name of the tag.
	Name string `json:"name" api:"required"`
	Slug string `json:"slug" api:"required"`
	// ID of the primary tag this object was translated from.
	TranslatedFromID int64 `json:"translatedFromId" api:"required"`
	// The timestamp (ISO8601 format) when this Blog Tag was updated.
	Updated time.Time `json:"updated" api:"required" format:"date-time"`
	paramObj
}

func (r TagParam) MarshalJSON() (data []byte, err error) {
	type shadow TagParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TagParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The explicitly defined ISO 639 language code of the tag.
type TagLanguage string

const (
	TagLanguageAa     TagLanguage = "aa"
	TagLanguageAb     TagLanguage = "ab"
	TagLanguageAe     TagLanguage = "ae"
	TagLanguageAf     TagLanguage = "af"
	TagLanguageAfNa   TagLanguage = "af-na"
	TagLanguageAfZa   TagLanguage = "af-za"
	TagLanguageAgq    TagLanguage = "agq"
	TagLanguageAgqCm  TagLanguage = "agq-cm"
	TagLanguageAk     TagLanguage = "ak"
	TagLanguageAkGh   TagLanguage = "ak-gh"
	TagLanguageAm     TagLanguage = "am"
	TagLanguageAmEt   TagLanguage = "am-et"
	TagLanguageAn     TagLanguage = "an"
	TagLanguageAnn    TagLanguage = "ann"
	TagLanguageAnnNg  TagLanguage = "ann-ng"
	TagLanguageAr     TagLanguage = "ar"
	TagLanguageAr001  TagLanguage = "ar-001"
	TagLanguageArAe   TagLanguage = "ar-ae"
	TagLanguageArBh   TagLanguage = "ar-bh"
	TagLanguageArDj   TagLanguage = "ar-dj"
	TagLanguageArDz   TagLanguage = "ar-dz"
	TagLanguageArEg   TagLanguage = "ar-eg"
	TagLanguageArEh   TagLanguage = "ar-eh"
	TagLanguageArEr   TagLanguage = "ar-er"
	TagLanguageArIl   TagLanguage = "ar-il"
	TagLanguageArIq   TagLanguage = "ar-iq"
	TagLanguageArJo   TagLanguage = "ar-jo"
	TagLanguageArKm   TagLanguage = "ar-km"
	TagLanguageArKw   TagLanguage = "ar-kw"
	TagLanguageArLb   TagLanguage = "ar-lb"
	TagLanguageArLy   TagLanguage = "ar-ly"
	TagLanguageArMa   TagLanguage = "ar-ma"
	TagLanguageArMr   TagLanguage = "ar-mr"
	TagLanguageArOm   TagLanguage = "ar-om"
	TagLanguageArPs   TagLanguage = "ar-ps"
	TagLanguageArQa   TagLanguage = "ar-qa"
	TagLanguageArSa   TagLanguage = "ar-sa"
	TagLanguageArSd   TagLanguage = "ar-sd"
	TagLanguageArSo   TagLanguage = "ar-so"
	TagLanguageArSS   TagLanguage = "ar-ss"
	TagLanguageArSy   TagLanguage = "ar-sy"
	TagLanguageArTd   TagLanguage = "ar-td"
	TagLanguageArTn   TagLanguage = "ar-tn"
	TagLanguageArYe   TagLanguage = "ar-ye"
	TagLanguageAs     TagLanguage = "as"
	TagLanguageAsIn   TagLanguage = "as-in"
	TagLanguageAsa    TagLanguage = "asa"
	TagLanguageAsaTz  TagLanguage = "asa-tz"
	TagLanguageAst    TagLanguage = "ast"
	TagLanguageAstEs  TagLanguage = "ast-es"
	TagLanguageAv     TagLanguage = "av"
	TagLanguageAy     TagLanguage = "ay"
	TagLanguageAz     TagLanguage = "az"
	TagLanguageAzAz   TagLanguage = "az-az"
	TagLanguageBa     TagLanguage = "ba"
	TagLanguageBas    TagLanguage = "bas"
	TagLanguageBasCm  TagLanguage = "bas-cm"
	TagLanguageBe     TagLanguage = "be"
	TagLanguageBeBy   TagLanguage = "be-by"
	TagLanguageBem    TagLanguage = "bem"
	TagLanguageBemZm  TagLanguage = "bem-zm"
	TagLanguageBez    TagLanguage = "bez"
	TagLanguageBezTz  TagLanguage = "bez-tz"
	TagLanguageBg     TagLanguage = "bg"
	TagLanguageBgBg   TagLanguage = "bg-bg"
	TagLanguageBgc    TagLanguage = "bgc"
	TagLanguageBgcIn  TagLanguage = "bgc-in"
	TagLanguageBho    TagLanguage = "bho"
	TagLanguageBhoIn  TagLanguage = "bho-in"
	TagLanguageBi     TagLanguage = "bi"
	TagLanguageBm     TagLanguage = "bm"
	TagLanguageBmMl   TagLanguage = "bm-ml"
	TagLanguageBn     TagLanguage = "bn"
	TagLanguageBnBd   TagLanguage = "bn-bd"
	TagLanguageBnIn   TagLanguage = "bn-in"
	TagLanguageBo     TagLanguage = "bo"
	TagLanguageBoCn   TagLanguage = "bo-cn"
	TagLanguageBoIn   TagLanguage = "bo-in"
	TagLanguageBr     TagLanguage = "br"
	TagLanguageBrFr   TagLanguage = "br-fr"
	TagLanguageBrx    TagLanguage = "brx"
	TagLanguageBrxIn  TagLanguage = "brx-in"
	TagLanguageBs     TagLanguage = "bs"
	TagLanguageBsBa   TagLanguage = "bs-ba"
	TagLanguageCa     TagLanguage = "ca"
	TagLanguageCaAd   TagLanguage = "ca-ad"
	TagLanguageCaEs   TagLanguage = "ca-es"
	TagLanguageCaFr   TagLanguage = "ca-fr"
	TagLanguageCaIt   TagLanguage = "ca-it"
	TagLanguageCcp    TagLanguage = "ccp"
	TagLanguageCcpBd  TagLanguage = "ccp-bd"
	TagLanguageCcpIn  TagLanguage = "ccp-in"
	TagLanguageCe     TagLanguage = "ce"
	TagLanguageCeRu   TagLanguage = "ce-ru"
	TagLanguageCeb    TagLanguage = "ceb"
	TagLanguageCebPh  TagLanguage = "ceb-ph"
	TagLanguageCgg    TagLanguage = "cgg"
	TagLanguageCggUg  TagLanguage = "cgg-ug"
	TagLanguageCh     TagLanguage = "ch"
	TagLanguageChr    TagLanguage = "chr"
	TagLanguageChrUs  TagLanguage = "chr-us"
	TagLanguageCkb    TagLanguage = "ckb"
	TagLanguageCkbIq  TagLanguage = "ckb-iq"
	TagLanguageCkbIr  TagLanguage = "ckb-ir"
	TagLanguageCo     TagLanguage = "co"
	TagLanguageCr     TagLanguage = "cr"
	TagLanguageCs     TagLanguage = "cs"
	TagLanguageCsCz   TagLanguage = "cs-cz"
	TagLanguageCu     TagLanguage = "cu"
	TagLanguageCuRu   TagLanguage = "cu-ru"
	TagLanguageCv     TagLanguage = "cv"
	TagLanguageCvRu   TagLanguage = "cv-ru"
	TagLanguageCy     TagLanguage = "cy"
	TagLanguageCyGB   TagLanguage = "cy-gb"
	TagLanguageDa     TagLanguage = "da"
	TagLanguageDaDk   TagLanguage = "da-dk"
	TagLanguageDaGl   TagLanguage = "da-gl"
	TagLanguageDav    TagLanguage = "dav"
	TagLanguageDavKe  TagLanguage = "dav-ke"
	TagLanguageDe     TagLanguage = "de"
	TagLanguageDeAt   TagLanguage = "de-at"
	TagLanguageDeBe   TagLanguage = "de-be"
	TagLanguageDeCh   TagLanguage = "de-ch"
	TagLanguageDeDe   TagLanguage = "de-de"
	TagLanguageDeGr   TagLanguage = "de-gr"
	TagLanguageDeIt   TagLanguage = "de-it"
	TagLanguageDeLi   TagLanguage = "de-li"
	TagLanguageDeLu   TagLanguage = "de-lu"
	TagLanguageDje    TagLanguage = "dje"
	TagLanguageDjeNe  TagLanguage = "dje-ne"
	TagLanguageDoi    TagLanguage = "doi"
	TagLanguageDoiIn  TagLanguage = "doi-in"
	TagLanguageDsb    TagLanguage = "dsb"
	TagLanguageDsbDe  TagLanguage = "dsb-de"
	TagLanguageDua    TagLanguage = "dua"
	TagLanguageDuaCm  TagLanguage = "dua-cm"
	TagLanguageDv     TagLanguage = "dv"
	TagLanguageDyo    TagLanguage = "dyo"
	TagLanguageDyoSn  TagLanguage = "dyo-sn"
	TagLanguageDz     TagLanguage = "dz"
	TagLanguageDzBt   TagLanguage = "dz-bt"
	TagLanguageEbu    TagLanguage = "ebu"
	TagLanguageEbuKe  TagLanguage = "ebu-ke"
	TagLanguageEe     TagLanguage = "ee"
	TagLanguageEeGh   TagLanguage = "ee-gh"
	TagLanguageEeTg   TagLanguage = "ee-tg"
	TagLanguageEl     TagLanguage = "el"
	TagLanguageElCy   TagLanguage = "el-cy"
	TagLanguageElGr   TagLanguage = "el-gr"
	TagLanguageEn     TagLanguage = "en"
	TagLanguageEn001  TagLanguage = "en-001"
	TagLanguageEn150  TagLanguage = "en-150"
	TagLanguageEnAe   TagLanguage = "en-ae"
	TagLanguageEnAg   TagLanguage = "en-ag"
	TagLanguageEnAI   TagLanguage = "en-ai"
	TagLanguageEnAs   TagLanguage = "en-as"
	TagLanguageEnAt   TagLanguage = "en-at"
	TagLanguageEnAu   TagLanguage = "en-au"
	TagLanguageEnBb   TagLanguage = "en-bb"
	TagLanguageEnBe   TagLanguage = "en-be"
	TagLanguageEnBi   TagLanguage = "en-bi"
	TagLanguageEnBm   TagLanguage = "en-bm"
	TagLanguageEnBs   TagLanguage = "en-bs"
	TagLanguageEnBw   TagLanguage = "en-bw"
	TagLanguageEnBz   TagLanguage = "en-bz"
	TagLanguageEnCa   TagLanguage = "en-ca"
	TagLanguageEnCc   TagLanguage = "en-cc"
	TagLanguageEnCh   TagLanguage = "en-ch"
	TagLanguageEnCk   TagLanguage = "en-ck"
	TagLanguageEnCm   TagLanguage = "en-cm"
	TagLanguageEnCn   TagLanguage = "en-cn"
	TagLanguageEnCx   TagLanguage = "en-cx"
	TagLanguageEnCy   TagLanguage = "en-cy"
	TagLanguageEnDe   TagLanguage = "en-de"
	TagLanguageEnDg   TagLanguage = "en-dg"
	TagLanguageEnDk   TagLanguage = "en-dk"
	TagLanguageEnDm   TagLanguage = "en-dm"
	TagLanguageEnEe   TagLanguage = "en-ee"
	TagLanguageEnEg   TagLanguage = "en-eg"
	TagLanguageEnEr   TagLanguage = "en-er"
	TagLanguageEnEs   TagLanguage = "en-es"
	TagLanguageEnFi   TagLanguage = "en-fi"
	TagLanguageEnFj   TagLanguage = "en-fj"
	TagLanguageEnFk   TagLanguage = "en-fk"
	TagLanguageEnFm   TagLanguage = "en-fm"
	TagLanguageEnFr   TagLanguage = "en-fr"
	TagLanguageEnGB   TagLanguage = "en-gb"
	TagLanguageEnGd   TagLanguage = "en-gd"
	TagLanguageEnGg   TagLanguage = "en-gg"
	TagLanguageEnGh   TagLanguage = "en-gh"
	TagLanguageEnGi   TagLanguage = "en-gi"
	TagLanguageEnGm   TagLanguage = "en-gm"
	TagLanguageEnGu   TagLanguage = "en-gu"
	TagLanguageEnGy   TagLanguage = "en-gy"
	TagLanguageEnHk   TagLanguage = "en-hk"
	TagLanguageEnID   TagLanguage = "en-id"
	TagLanguageEnIe   TagLanguage = "en-ie"
	TagLanguageEnIl   TagLanguage = "en-il"
	TagLanguageEnIm   TagLanguage = "en-im"
	TagLanguageEnIn   TagLanguage = "en-in"
	TagLanguageEnIo   TagLanguage = "en-io"
	TagLanguageEnJe   TagLanguage = "en-je"
	TagLanguageEnJm   TagLanguage = "en-jm"
	TagLanguageEnKe   TagLanguage = "en-ke"
	TagLanguageEnKi   TagLanguage = "en-ki"
	TagLanguageEnKn   TagLanguage = "en-kn"
	TagLanguageEnKy   TagLanguage = "en-ky"
	TagLanguageEnLc   TagLanguage = "en-lc"
	TagLanguageEnLr   TagLanguage = "en-lr"
	TagLanguageEnLs   TagLanguage = "en-ls"
	TagLanguageEnLu   TagLanguage = "en-lu"
	TagLanguageEnMg   TagLanguage = "en-mg"
	TagLanguageEnMh   TagLanguage = "en-mh"
	TagLanguageEnMo   TagLanguage = "en-mo"
	TagLanguageEnMp   TagLanguage = "en-mp"
	TagLanguageEnMs   TagLanguage = "en-ms"
	TagLanguageEnMt   TagLanguage = "en-mt"
	TagLanguageEnMu   TagLanguage = "en-mu"
	TagLanguageEnMv   TagLanguage = "en-mv"
	TagLanguageEnMw   TagLanguage = "en-mw"
	TagLanguageEnMx   TagLanguage = "en-mx"
	TagLanguageEnMy   TagLanguage = "en-my"
	TagLanguageEnNa   TagLanguage = "en-na"
	TagLanguageEnNf   TagLanguage = "en-nf"
	TagLanguageEnNg   TagLanguage = "en-ng"
	TagLanguageEnNl   TagLanguage = "en-nl"
	TagLanguageEnNr   TagLanguage = "en-nr"
	TagLanguageEnNu   TagLanguage = "en-nu"
	TagLanguageEnNz   TagLanguage = "en-nz"
	TagLanguageEnPg   TagLanguage = "en-pg"
	TagLanguageEnPh   TagLanguage = "en-ph"
	TagLanguageEnPk   TagLanguage = "en-pk"
	TagLanguageEnPn   TagLanguage = "en-pn"
	TagLanguageEnPr   TagLanguage = "en-pr"
	TagLanguageEnPt   TagLanguage = "en-pt"
	TagLanguageEnPw   TagLanguage = "en-pw"
	TagLanguageEnRw   TagLanguage = "en-rw"
	TagLanguageEnSb   TagLanguage = "en-sb"
	TagLanguageEnSc   TagLanguage = "en-sc"
	TagLanguageEnSd   TagLanguage = "en-sd"
	TagLanguageEnSe   TagLanguage = "en-se"
	TagLanguageEnSg   TagLanguage = "en-sg"
	TagLanguageEnSh   TagLanguage = "en-sh"
	TagLanguageEnSi   TagLanguage = "en-si"
	TagLanguageEnSl   TagLanguage = "en-sl"
	TagLanguageEnSS   TagLanguage = "en-ss"
	TagLanguageEnSx   TagLanguage = "en-sx"
	TagLanguageEnSz   TagLanguage = "en-sz"
	TagLanguageEnTc   TagLanguage = "en-tc"
	TagLanguageEnTh   TagLanguage = "en-th"
	TagLanguageEnTk   TagLanguage = "en-tk"
	TagLanguageEnTn   TagLanguage = "en-tn"
	TagLanguageEnTo   TagLanguage = "en-to"
	TagLanguageEnTt   TagLanguage = "en-tt"
	TagLanguageEnTv   TagLanguage = "en-tv"
	TagLanguageEnTz   TagLanguage = "en-tz"
	TagLanguageEnUg   TagLanguage = "en-ug"
	TagLanguageEnUm   TagLanguage = "en-um"
	TagLanguageEnUs   TagLanguage = "en-us"
	TagLanguageEnVc   TagLanguage = "en-vc"
	TagLanguageEnVg   TagLanguage = "en-vg"
	TagLanguageEnVi   TagLanguage = "en-vi"
	TagLanguageEnVn   TagLanguage = "en-vn"
	TagLanguageEnVu   TagLanguage = "en-vu"
	TagLanguageEnWs   TagLanguage = "en-ws"
	TagLanguageEnZa   TagLanguage = "en-za"
	TagLanguageEnZm   TagLanguage = "en-zm"
	TagLanguageEnZw   TagLanguage = "en-zw"
	TagLanguageEo     TagLanguage = "eo"
	TagLanguageEo001  TagLanguage = "eo-001"
	TagLanguageEs     TagLanguage = "es"
	TagLanguageEs419  TagLanguage = "es-419"
	TagLanguageEsAr   TagLanguage = "es-ar"
	TagLanguageEsBo   TagLanguage = "es-bo"
	TagLanguageEsBr   TagLanguage = "es-br"
	TagLanguageEsBz   TagLanguage = "es-bz"
	TagLanguageEsCl   TagLanguage = "es-cl"
	TagLanguageEsCo   TagLanguage = "es-co"
	TagLanguageEsCr   TagLanguage = "es-cr"
	TagLanguageEsCu   TagLanguage = "es-cu"
	TagLanguageEsDo   TagLanguage = "es-do"
	TagLanguageEsEa   TagLanguage = "es-ea"
	TagLanguageEsEc   TagLanguage = "es-ec"
	TagLanguageEsEs   TagLanguage = "es-es"
	TagLanguageEsGq   TagLanguage = "es-gq"
	TagLanguageEsGt   TagLanguage = "es-gt"
	TagLanguageEsHn   TagLanguage = "es-hn"
	TagLanguageEsIc   TagLanguage = "es-ic"
	TagLanguageEsMx   TagLanguage = "es-mx"
	TagLanguageEsNi   TagLanguage = "es-ni"
	TagLanguageEsPa   TagLanguage = "es-pa"
	TagLanguageEsPe   TagLanguage = "es-pe"
	TagLanguageEsPh   TagLanguage = "es-ph"
	TagLanguageEsPr   TagLanguage = "es-pr"
	TagLanguageEsPy   TagLanguage = "es-py"
	TagLanguageEsSv   TagLanguage = "es-sv"
	TagLanguageEsUs   TagLanguage = "es-us"
	TagLanguageEsUy   TagLanguage = "es-uy"
	TagLanguageEsVe   TagLanguage = "es-ve"
	TagLanguageEt     TagLanguage = "et"
	TagLanguageEtEe   TagLanguage = "et-ee"
	TagLanguageEu     TagLanguage = "eu"
	TagLanguageEuEs   TagLanguage = "eu-es"
	TagLanguageEwo    TagLanguage = "ewo"
	TagLanguageEwoCm  TagLanguage = "ewo-cm"
	TagLanguageFa     TagLanguage = "fa"
	TagLanguageFaAf   TagLanguage = "fa-af"
	TagLanguageFaIr   TagLanguage = "fa-ir"
	TagLanguageFf     TagLanguage = "ff"
	TagLanguageFfBf   TagLanguage = "ff-bf"
	TagLanguageFfCm   TagLanguage = "ff-cm"
	TagLanguageFfGh   TagLanguage = "ff-gh"
	TagLanguageFfGm   TagLanguage = "ff-gm"
	TagLanguageFfGn   TagLanguage = "ff-gn"
	TagLanguageFfGw   TagLanguage = "ff-gw"
	TagLanguageFfLr   TagLanguage = "ff-lr"
	TagLanguageFfMr   TagLanguage = "ff-mr"
	TagLanguageFfNe   TagLanguage = "ff-ne"
	TagLanguageFfNg   TagLanguage = "ff-ng"
	TagLanguageFfSl   TagLanguage = "ff-sl"
	TagLanguageFfSn   TagLanguage = "ff-sn"
	TagLanguageFi     TagLanguage = "fi"
	TagLanguageFiFi   TagLanguage = "fi-fi"
	TagLanguageFil    TagLanguage = "fil"
	TagLanguageFilPh  TagLanguage = "fil-ph"
	TagLanguageFj     TagLanguage = "fj"
	TagLanguageFo     TagLanguage = "fo"
	TagLanguageFoDk   TagLanguage = "fo-dk"
	TagLanguageFoFo   TagLanguage = "fo-fo"
	TagLanguageFr     TagLanguage = "fr"
	TagLanguageFrBe   TagLanguage = "fr-be"
	TagLanguageFrBf   TagLanguage = "fr-bf"
	TagLanguageFrBi   TagLanguage = "fr-bi"
	TagLanguageFrBj   TagLanguage = "fr-bj"
	TagLanguageFrBl   TagLanguage = "fr-bl"
	TagLanguageFrCa   TagLanguage = "fr-ca"
	TagLanguageFrCd   TagLanguage = "fr-cd"
	TagLanguageFrCf   TagLanguage = "fr-cf"
	TagLanguageFrCg   TagLanguage = "fr-cg"
	TagLanguageFrCh   TagLanguage = "fr-ch"
	TagLanguageFrCi   TagLanguage = "fr-ci"
	TagLanguageFrCm   TagLanguage = "fr-cm"
	TagLanguageFrDj   TagLanguage = "fr-dj"
	TagLanguageFrDz   TagLanguage = "fr-dz"
	TagLanguageFrFr   TagLanguage = "fr-fr"
	TagLanguageFrGa   TagLanguage = "fr-ga"
	TagLanguageFrGf   TagLanguage = "fr-gf"
	TagLanguageFrGn   TagLanguage = "fr-gn"
	TagLanguageFrGp   TagLanguage = "fr-gp"
	TagLanguageFrGq   TagLanguage = "fr-gq"
	TagLanguageFrHt   TagLanguage = "fr-ht"
	TagLanguageFrKm   TagLanguage = "fr-km"
	TagLanguageFrLu   TagLanguage = "fr-lu"
	TagLanguageFrMa   TagLanguage = "fr-ma"
	TagLanguageFrMc   TagLanguage = "fr-mc"
	TagLanguageFrMf   TagLanguage = "fr-mf"
	TagLanguageFrMg   TagLanguage = "fr-mg"
	TagLanguageFrMl   TagLanguage = "fr-ml"
	TagLanguageFrMq   TagLanguage = "fr-mq"
	TagLanguageFrMr   TagLanguage = "fr-mr"
	TagLanguageFrMu   TagLanguage = "fr-mu"
	TagLanguageFrNc   TagLanguage = "fr-nc"
	TagLanguageFrNe   TagLanguage = "fr-ne"
	TagLanguageFrPf   TagLanguage = "fr-pf"
	TagLanguageFrPm   TagLanguage = "fr-pm"
	TagLanguageFrRe   TagLanguage = "fr-re"
	TagLanguageFrRw   TagLanguage = "fr-rw"
	TagLanguageFrSc   TagLanguage = "fr-sc"
	TagLanguageFrSn   TagLanguage = "fr-sn"
	TagLanguageFrSy   TagLanguage = "fr-sy"
	TagLanguageFrTd   TagLanguage = "fr-td"
	TagLanguageFrTg   TagLanguage = "fr-tg"
	TagLanguageFrTn   TagLanguage = "fr-tn"
	TagLanguageFrVu   TagLanguage = "fr-vu"
	TagLanguageFrWf   TagLanguage = "fr-wf"
	TagLanguageFrYt   TagLanguage = "fr-yt"
	TagLanguageFrr    TagLanguage = "frr"
	TagLanguageFrrDe  TagLanguage = "frr-de"
	TagLanguageFur    TagLanguage = "fur"
	TagLanguageFurIt  TagLanguage = "fur-it"
	TagLanguageFy     TagLanguage = "fy"
	TagLanguageFyNl   TagLanguage = "fy-nl"
	TagLanguageGa     TagLanguage = "ga"
	TagLanguageGaGB   TagLanguage = "ga-gb"
	TagLanguageGaIe   TagLanguage = "ga-ie"
	TagLanguageGd     TagLanguage = "gd"
	TagLanguageGdGB   TagLanguage = "gd-gb"
	TagLanguageGl     TagLanguage = "gl"
	TagLanguageGlEs   TagLanguage = "gl-es"
	TagLanguageGn     TagLanguage = "gn"
	TagLanguageGsw    TagLanguage = "gsw"
	TagLanguageGswCh  TagLanguage = "gsw-ch"
	TagLanguageGswFr  TagLanguage = "gsw-fr"
	TagLanguageGswLi  TagLanguage = "gsw-li"
	TagLanguageGu     TagLanguage = "gu"
	TagLanguageGuIn   TagLanguage = "gu-in"
	TagLanguageGuz    TagLanguage = "guz"
	TagLanguageGuzKe  TagLanguage = "guz-ke"
	TagLanguageGv     TagLanguage = "gv"
	TagLanguageGvIm   TagLanguage = "gv-im"
	TagLanguageHa     TagLanguage = "ha"
	TagLanguageHaGh   TagLanguage = "ha-gh"
	TagLanguageHaNe   TagLanguage = "ha-ne"
	TagLanguageHaNg   TagLanguage = "ha-ng"
	TagLanguageHaw    TagLanguage = "haw"
	TagLanguageHawUs  TagLanguage = "haw-us"
	TagLanguageHe     TagLanguage = "he"
	TagLanguageHeIl   TagLanguage = "he-il"
	TagLanguageHi     TagLanguage = "hi"
	TagLanguageHiIn   TagLanguage = "hi-in"
	TagLanguageHmn    TagLanguage = "hmn"
	TagLanguageHo     TagLanguage = "ho"
	TagLanguageHr     TagLanguage = "hr"
	TagLanguageHrBa   TagLanguage = "hr-ba"
	TagLanguageHrHr   TagLanguage = "hr-hr"
	TagLanguageHsb    TagLanguage = "hsb"
	TagLanguageHsbDe  TagLanguage = "hsb-de"
	TagLanguageHt     TagLanguage = "ht"
	TagLanguageHu     TagLanguage = "hu"
	TagLanguageHuHu   TagLanguage = "hu-hu"
	TagLanguageHy     TagLanguage = "hy"
	TagLanguageHyAm   TagLanguage = "hy-am"
	TagLanguageHz     TagLanguage = "hz"
	TagLanguageIa     TagLanguage = "ia"
	TagLanguageIa001  TagLanguage = "ia-001"
	TagLanguageID     TagLanguage = "id"
	TagLanguageIDID   TagLanguage = "id-id"
	TagLanguageIe     TagLanguage = "ie"
	TagLanguageIg     TagLanguage = "ig"
	TagLanguageIgNg   TagLanguage = "ig-ng"
	TagLanguageIi     TagLanguage = "ii"
	TagLanguageIiCn   TagLanguage = "ii-cn"
	TagLanguageIk     TagLanguage = "ik"
	TagLanguageIo     TagLanguage = "io"
	TagLanguageIs     TagLanguage = "is"
	TagLanguageIsIs   TagLanguage = "is-is"
	TagLanguageIt     TagLanguage = "it"
	TagLanguageItCh   TagLanguage = "it-ch"
	TagLanguageItIt   TagLanguage = "it-it"
	TagLanguageItSm   TagLanguage = "it-sm"
	TagLanguageItVa   TagLanguage = "it-va"
	TagLanguageIu     TagLanguage = "iu"
	TagLanguageJa     TagLanguage = "ja"
	TagLanguageJaJp   TagLanguage = "ja-jp"
	TagLanguageJgo    TagLanguage = "jgo"
	TagLanguageJgoCm  TagLanguage = "jgo-cm"
	TagLanguageJmc    TagLanguage = "jmc"
	TagLanguageJmcTz  TagLanguage = "jmc-tz"
	TagLanguageJv     TagLanguage = "jv"
	TagLanguageJvID   TagLanguage = "jv-id"
	TagLanguageKa     TagLanguage = "ka"
	TagLanguageKaGe   TagLanguage = "ka-ge"
	TagLanguageKab    TagLanguage = "kab"
	TagLanguageKabDz  TagLanguage = "kab-dz"
	TagLanguageKam    TagLanguage = "kam"
	TagLanguageKamKe  TagLanguage = "kam-ke"
	TagLanguageKar    TagLanguage = "kar"
	TagLanguageKde    TagLanguage = "kde"
	TagLanguageKdeTz  TagLanguage = "kde-tz"
	TagLanguageKea    TagLanguage = "kea"
	TagLanguageKeaCv  TagLanguage = "kea-cv"
	TagLanguageKg     TagLanguage = "kg"
	TagLanguageKgp    TagLanguage = "kgp"
	TagLanguageKgpBr  TagLanguage = "kgp-br"
	TagLanguageKh     TagLanguage = "kh"
	TagLanguageKhq    TagLanguage = "khq"
	TagLanguageKhqMl  TagLanguage = "khq-ml"
	TagLanguageKi     TagLanguage = "ki"
	TagLanguageKiKe   TagLanguage = "ki-ke"
	TagLanguageKj     TagLanguage = "kj"
	TagLanguageKk     TagLanguage = "kk"
	TagLanguageKkKz   TagLanguage = "kk-kz"
	TagLanguageKkj    TagLanguage = "kkj"
	TagLanguageKkjCm  TagLanguage = "kkj-cm"
	TagLanguageKl     TagLanguage = "kl"
	TagLanguageKlGl   TagLanguage = "kl-gl"
	TagLanguageKln    TagLanguage = "kln"
	TagLanguageKlnKe  TagLanguage = "kln-ke"
	TagLanguageKm     TagLanguage = "km"
	TagLanguageKmKh   TagLanguage = "km-kh"
	TagLanguageKn     TagLanguage = "kn"
	TagLanguageKnIn   TagLanguage = "kn-in"
	TagLanguageKo     TagLanguage = "ko"
	TagLanguageKoKp   TagLanguage = "ko-kp"
	TagLanguageKoKr   TagLanguage = "ko-kr"
	TagLanguageKok    TagLanguage = "kok"
	TagLanguageKokIn  TagLanguage = "kok-in"
	TagLanguageKr     TagLanguage = "kr"
	TagLanguageKs     TagLanguage = "ks"
	TagLanguageKsIn   TagLanguage = "ks-in"
	TagLanguageKsb    TagLanguage = "ksb"
	TagLanguageKsbTz  TagLanguage = "ksb-tz"
	TagLanguageKsf    TagLanguage = "ksf"
	TagLanguageKsfCm  TagLanguage = "ksf-cm"
	TagLanguageKsh    TagLanguage = "ksh"
	TagLanguageKshDe  TagLanguage = "ksh-de"
	TagLanguageKu     TagLanguage = "ku"
	TagLanguageKuTr   TagLanguage = "ku-tr"
	TagLanguageKv     TagLanguage = "kv"
	TagLanguageKw     TagLanguage = "kw"
	TagLanguageKwGB   TagLanguage = "kw-gb"
	TagLanguageKy     TagLanguage = "ky"
	TagLanguageKyKg   TagLanguage = "ky-kg"
	TagLanguageLa     TagLanguage = "la"
	TagLanguageLag    TagLanguage = "lag"
	TagLanguageLagTz  TagLanguage = "lag-tz"
	TagLanguageLb     TagLanguage = "lb"
	TagLanguageLbLu   TagLanguage = "lb-lu"
	TagLanguageLg     TagLanguage = "lg"
	TagLanguageLgUg   TagLanguage = "lg-ug"
	TagLanguageLi     TagLanguage = "li"
	TagLanguageLkt    TagLanguage = "lkt"
	TagLanguageLktUs  TagLanguage = "lkt-us"
	TagLanguageLn     TagLanguage = "ln"
	TagLanguageLnAo   TagLanguage = "ln-ao"
	TagLanguageLnCd   TagLanguage = "ln-cd"
	TagLanguageLnCf   TagLanguage = "ln-cf"
	TagLanguageLnCg   TagLanguage = "ln-cg"
	TagLanguageLo     TagLanguage = "lo"
	TagLanguageLoLa   TagLanguage = "lo-la"
	TagLanguageLrc    TagLanguage = "lrc"
	TagLanguageLrcIq  TagLanguage = "lrc-iq"
	TagLanguageLrcIr  TagLanguage = "lrc-ir"
	TagLanguageLt     TagLanguage = "lt"
	TagLanguageLtLt   TagLanguage = "lt-lt"
	TagLanguageLu     TagLanguage = "lu"
	TagLanguageLuCd   TagLanguage = "lu-cd"
	TagLanguageLuo    TagLanguage = "luo"
	TagLanguageLuoKe  TagLanguage = "luo-ke"
	TagLanguageLuy    TagLanguage = "luy"
	TagLanguageLuyKe  TagLanguage = "luy-ke"
	TagLanguageLv     TagLanguage = "lv"
	TagLanguageLvLv   TagLanguage = "lv-lv"
	TagLanguageMai    TagLanguage = "mai"
	TagLanguageMaiIn  TagLanguage = "mai-in"
	TagLanguageMas    TagLanguage = "mas"
	TagLanguageMasKe  TagLanguage = "mas-ke"
	TagLanguageMasTz  TagLanguage = "mas-tz"
	TagLanguageMdf    TagLanguage = "mdf"
	TagLanguageMdfRu  TagLanguage = "mdf-ru"
	TagLanguageMer    TagLanguage = "mer"
	TagLanguageMerKe  TagLanguage = "mer-ke"
	TagLanguageMfe    TagLanguage = "mfe"
	TagLanguageMfeMu  TagLanguage = "mfe-mu"
	TagLanguageMg     TagLanguage = "mg"
	TagLanguageMgMg   TagLanguage = "mg-mg"
	TagLanguageMgh    TagLanguage = "mgh"
	TagLanguageMghMz  TagLanguage = "mgh-mz"
	TagLanguageMgo    TagLanguage = "mgo"
	TagLanguageMgoCm  TagLanguage = "mgo-cm"
	TagLanguageMh     TagLanguage = "mh"
	TagLanguageMi     TagLanguage = "mi"
	TagLanguageMiNz   TagLanguage = "mi-nz"
	TagLanguageMk     TagLanguage = "mk"
	TagLanguageMkMk   TagLanguage = "mk-mk"
	TagLanguageMl     TagLanguage = "ml"
	TagLanguageMlIn   TagLanguage = "ml-in"
	TagLanguageMn     TagLanguage = "mn"
	TagLanguageMnMn   TagLanguage = "mn-mn"
	TagLanguageMni    TagLanguage = "mni"
	TagLanguageMniIn  TagLanguage = "mni-in"
	TagLanguageMr     TagLanguage = "mr"
	TagLanguageMrIn   TagLanguage = "mr-in"
	TagLanguageMs     TagLanguage = "ms"
	TagLanguageMsBn   TagLanguage = "ms-bn"
	TagLanguageMsID   TagLanguage = "ms-id"
	TagLanguageMsMy   TagLanguage = "ms-my"
	TagLanguageMsSg   TagLanguage = "ms-sg"
	TagLanguageMt     TagLanguage = "mt"
	TagLanguageMtMt   TagLanguage = "mt-mt"
	TagLanguageMua    TagLanguage = "mua"
	TagLanguageMuaCm  TagLanguage = "mua-cm"
	TagLanguageMy     TagLanguage = "my"
	TagLanguageMyMm   TagLanguage = "my-mm"
	TagLanguageMzn    TagLanguage = "mzn"
	TagLanguageMznIr  TagLanguage = "mzn-ir"
	TagLanguageNa     TagLanguage = "na"
	TagLanguageNaq    TagLanguage = "naq"
	TagLanguageNaqNa  TagLanguage = "naq-na"
	TagLanguageNb     TagLanguage = "nb"
	TagLanguageNbNo   TagLanguage = "nb-no"
	TagLanguageNbSj   TagLanguage = "nb-sj"
	TagLanguageNd     TagLanguage = "nd"
	TagLanguageNdZw   TagLanguage = "nd-zw"
	TagLanguageNds    TagLanguage = "nds"
	TagLanguageNdsDe  TagLanguage = "nds-de"
	TagLanguageNdsNl  TagLanguage = "nds-nl"
	TagLanguageNe     TagLanguage = "ne"
	TagLanguageNeIn   TagLanguage = "ne-in"
	TagLanguageNeNp   TagLanguage = "ne-np"
	TagLanguageNg     TagLanguage = "ng"
	TagLanguageNl     TagLanguage = "nl"
	TagLanguageNlAw   TagLanguage = "nl-aw"
	TagLanguageNlBe   TagLanguage = "nl-be"
	TagLanguageNlBq   TagLanguage = "nl-bq"
	TagLanguageNlCh   TagLanguage = "nl-ch"
	TagLanguageNlCw   TagLanguage = "nl-cw"
	TagLanguageNlLu   TagLanguage = "nl-lu"
	TagLanguageNlNl   TagLanguage = "nl-nl"
	TagLanguageNlSr   TagLanguage = "nl-sr"
	TagLanguageNlSx   TagLanguage = "nl-sx"
	TagLanguageNmg    TagLanguage = "nmg"
	TagLanguageNmgCm  TagLanguage = "nmg-cm"
	TagLanguageNn     TagLanguage = "nn"
	TagLanguageNnNo   TagLanguage = "nn-no"
	TagLanguageNnh    TagLanguage = "nnh"
	TagLanguageNnhCm  TagLanguage = "nnh-cm"
	TagLanguageNo     TagLanguage = "no"
	TagLanguageNoNo   TagLanguage = "no-no"
	TagLanguageNr     TagLanguage = "nr"
	TagLanguageNus    TagLanguage = "nus"
	TagLanguageNusSS  TagLanguage = "nus-ss"
	TagLanguageNv     TagLanguage = "nv"
	TagLanguageNy     TagLanguage = "ny"
	TagLanguageNyn    TagLanguage = "nyn"
	TagLanguageNynUg  TagLanguage = "nyn-ug"
	TagLanguageOc     TagLanguage = "oc"
	TagLanguageOcEs   TagLanguage = "oc-es"
	TagLanguageOcFr   TagLanguage = "oc-fr"
	TagLanguageOj     TagLanguage = "oj"
	TagLanguageOm     TagLanguage = "om"
	TagLanguageOmEt   TagLanguage = "om-et"
	TagLanguageOmKe   TagLanguage = "om-ke"
	TagLanguageOr     TagLanguage = "or"
	TagLanguageOrIn   TagLanguage = "or-in"
	TagLanguageOs     TagLanguage = "os"
	TagLanguageOsGe   TagLanguage = "os-ge"
	TagLanguageOsRu   TagLanguage = "os-ru"
	TagLanguagePa     TagLanguage = "pa"
	TagLanguagePaIn   TagLanguage = "pa-in"
	TagLanguagePaPk   TagLanguage = "pa-pk"
	TagLanguagePcm    TagLanguage = "pcm"
	TagLanguagePcmNg  TagLanguage = "pcm-ng"
	TagLanguagePi     TagLanguage = "pi"
	TagLanguagePis    TagLanguage = "pis"
	TagLanguagePisSb  TagLanguage = "pis-sb"
	TagLanguagePl     TagLanguage = "pl"
	TagLanguagePlPl   TagLanguage = "pl-pl"
	TagLanguagePrg    TagLanguage = "prg"
	TagLanguagePrg001 TagLanguage = "prg-001"
	TagLanguagePs     TagLanguage = "ps"
	TagLanguagePsAf   TagLanguage = "ps-af"
	TagLanguagePsPk   TagLanguage = "ps-pk"
	TagLanguagePt     TagLanguage = "pt"
	TagLanguagePtAo   TagLanguage = "pt-ao"
	TagLanguagePtBr   TagLanguage = "pt-br"
	TagLanguagePtCh   TagLanguage = "pt-ch"
	TagLanguagePtCv   TagLanguage = "pt-cv"
	TagLanguagePtGq   TagLanguage = "pt-gq"
	TagLanguagePtGw   TagLanguage = "pt-gw"
	TagLanguagePtLu   TagLanguage = "pt-lu"
	TagLanguagePtMo   TagLanguage = "pt-mo"
	TagLanguagePtMz   TagLanguage = "pt-mz"
	TagLanguagePtPt   TagLanguage = "pt-pt"
	TagLanguagePtSt   TagLanguage = "pt-st"
	TagLanguagePtTl   TagLanguage = "pt-tl"
	TagLanguageQu     TagLanguage = "qu"
	TagLanguageQuBo   TagLanguage = "qu-bo"
	TagLanguageQuEc   TagLanguage = "qu-ec"
	TagLanguageQuPe   TagLanguage = "qu-pe"
	TagLanguageRaj    TagLanguage = "raj"
	TagLanguageRajIn  TagLanguage = "raj-in"
	TagLanguageRm     TagLanguage = "rm"
	TagLanguageRmCh   TagLanguage = "rm-ch"
	TagLanguageRn     TagLanguage = "rn"
	TagLanguageRnBi   TagLanguage = "rn-bi"
	TagLanguageRo     TagLanguage = "ro"
	TagLanguageRoMd   TagLanguage = "ro-md"
	TagLanguageRoRo   TagLanguage = "ro-ro"
	TagLanguageRof    TagLanguage = "rof"
	TagLanguageRofTz  TagLanguage = "rof-tz"
	TagLanguageRu     TagLanguage = "ru"
	TagLanguageRuBy   TagLanguage = "ru-by"
	TagLanguageRuKg   TagLanguage = "ru-kg"
	TagLanguageRuKz   TagLanguage = "ru-kz"
	TagLanguageRuMd   TagLanguage = "ru-md"
	TagLanguageRuRu   TagLanguage = "ru-ru"
	TagLanguageRuUa   TagLanguage = "ru-ua"
	TagLanguageRw     TagLanguage = "rw"
	TagLanguageRwRw   TagLanguage = "rw-rw"
	TagLanguageRwk    TagLanguage = "rwk"
	TagLanguageRwkTz  TagLanguage = "rwk-tz"
	TagLanguageSa     TagLanguage = "sa"
	TagLanguageSaIn   TagLanguage = "sa-in"
	TagLanguageSah    TagLanguage = "sah"
	TagLanguageSahRu  TagLanguage = "sah-ru"
	TagLanguageSaq    TagLanguage = "saq"
	TagLanguageSaqKe  TagLanguage = "saq-ke"
	TagLanguageSat    TagLanguage = "sat"
	TagLanguageSatIn  TagLanguage = "sat-in"
	TagLanguageSbp    TagLanguage = "sbp"
	TagLanguageSbpTz  TagLanguage = "sbp-tz"
	TagLanguageSc     TagLanguage = "sc"
	TagLanguageScIt   TagLanguage = "sc-it"
	TagLanguageSd     TagLanguage = "sd"
	TagLanguageSdIn   TagLanguage = "sd-in"
	TagLanguageSdPk   TagLanguage = "sd-pk"
	TagLanguageSe     TagLanguage = "se"
	TagLanguageSeFi   TagLanguage = "se-fi"
	TagLanguageSeNo   TagLanguage = "se-no"
	TagLanguageSeSe   TagLanguage = "se-se"
	TagLanguageSeh    TagLanguage = "seh"
	TagLanguageSehMz  TagLanguage = "seh-mz"
	TagLanguageSes    TagLanguage = "ses"
	TagLanguageSesMl  TagLanguage = "ses-ml"
	TagLanguageSg     TagLanguage = "sg"
	TagLanguageSgCf   TagLanguage = "sg-cf"
	TagLanguageShi    TagLanguage = "shi"
	TagLanguageShiMa  TagLanguage = "shi-ma"
	TagLanguageSi     TagLanguage = "si"
	TagLanguageSiLk   TagLanguage = "si-lk"
	TagLanguageSk     TagLanguage = "sk"
	TagLanguageSkSk   TagLanguage = "sk-sk"
	TagLanguageSl     TagLanguage = "sl"
	TagLanguageSlSi   TagLanguage = "sl-si"
	TagLanguageSm     TagLanguage = "sm"
	TagLanguageSmn    TagLanguage = "smn"
	TagLanguageSmnFi  TagLanguage = "smn-fi"
	TagLanguageSMS    TagLanguage = "sms"
	TagLanguageSMSFi  TagLanguage = "sms-fi"
	TagLanguageSn     TagLanguage = "sn"
	TagLanguageSnZw   TagLanguage = "sn-zw"
	TagLanguageSo     TagLanguage = "so"
	TagLanguageSoDj   TagLanguage = "so-dj"
	TagLanguageSoEt   TagLanguage = "so-et"
	TagLanguageSoKe   TagLanguage = "so-ke"
	TagLanguageSoSo   TagLanguage = "so-so"
	TagLanguageSq     TagLanguage = "sq"
	TagLanguageSqAl   TagLanguage = "sq-al"
	TagLanguageSqMk   TagLanguage = "sq-mk"
	TagLanguageSqXk   TagLanguage = "sq-xk"
	TagLanguageSr     TagLanguage = "sr"
	TagLanguageSrBa   TagLanguage = "sr-ba"
	TagLanguageSrCs   TagLanguage = "sr-cs"
	TagLanguageSrMe   TagLanguage = "sr-me"
	TagLanguageSrRs   TagLanguage = "sr-rs"
	TagLanguageSrXk   TagLanguage = "sr-xk"
	TagLanguageSS     TagLanguage = "ss"
	TagLanguageSt     TagLanguage = "st"
	TagLanguageSu     TagLanguage = "su"
	TagLanguageSuID   TagLanguage = "su-id"
	TagLanguageSv     TagLanguage = "sv"
	TagLanguageSvAx   TagLanguage = "sv-ax"
	TagLanguageSvFi   TagLanguage = "sv-fi"
	TagLanguageSvSe   TagLanguage = "sv-se"
	TagLanguageSw     TagLanguage = "sw"
	TagLanguageSwCd   TagLanguage = "sw-cd"
	TagLanguageSwKe   TagLanguage = "sw-ke"
	TagLanguageSwTz   TagLanguage = "sw-tz"
	TagLanguageSwUg   TagLanguage = "sw-ug"
	TagLanguageSy     TagLanguage = "sy"
	TagLanguageTa     TagLanguage = "ta"
	TagLanguageTaIn   TagLanguage = "ta-in"
	TagLanguageTaLk   TagLanguage = "ta-lk"
	TagLanguageTaMy   TagLanguage = "ta-my"
	TagLanguageTaSg   TagLanguage = "ta-sg"
	TagLanguageTe     TagLanguage = "te"
	TagLanguageTeIn   TagLanguage = "te-in"
	TagLanguageTeo    TagLanguage = "teo"
	TagLanguageTeoKe  TagLanguage = "teo-ke"
	TagLanguageTeoUg  TagLanguage = "teo-ug"
	TagLanguageTg     TagLanguage = "tg"
	TagLanguageTgTj   TagLanguage = "tg-tj"
	TagLanguageTh     TagLanguage = "th"
	TagLanguageThTh   TagLanguage = "th-th"
	TagLanguageTi     TagLanguage = "ti"
	TagLanguageTiEr   TagLanguage = "ti-er"
	TagLanguageTiEt   TagLanguage = "ti-et"
	TagLanguageTk     TagLanguage = "tk"
	TagLanguageTkTm   TagLanguage = "tk-tm"
	TagLanguageTl     TagLanguage = "tl"
	TagLanguageTn     TagLanguage = "tn"
	TagLanguageTo     TagLanguage = "to"
	TagLanguageToTo   TagLanguage = "to-to"
	TagLanguageTok    TagLanguage = "tok"
	TagLanguageTok001 TagLanguage = "tok-001"
	TagLanguageTr     TagLanguage = "tr"
	TagLanguageTrCy   TagLanguage = "tr-cy"
	TagLanguageTrTr   TagLanguage = "tr-tr"
	TagLanguageTs     TagLanguage = "ts"
	TagLanguageTt     TagLanguage = "tt"
	TagLanguageTtRu   TagLanguage = "tt-ru"
	TagLanguageTw     TagLanguage = "tw"
	TagLanguageTwq    TagLanguage = "twq"
	TagLanguageTwqNe  TagLanguage = "twq-ne"
	TagLanguageTy     TagLanguage = "ty"
	TagLanguageTzm    TagLanguage = "tzm"
	TagLanguageTzmMa  TagLanguage = "tzm-ma"
	TagLanguageUg     TagLanguage = "ug"
	TagLanguageUgCn   TagLanguage = "ug-cn"
	TagLanguageUk     TagLanguage = "uk"
	TagLanguageUkUa   TagLanguage = "uk-ua"
	TagLanguageUr     TagLanguage = "ur"
	TagLanguageUrIn   TagLanguage = "ur-in"
	TagLanguageUrPk   TagLanguage = "ur-pk"
	TagLanguageUz     TagLanguage = "uz"
	TagLanguageUzAf   TagLanguage = "uz-af"
	TagLanguageUzUz   TagLanguage = "uz-uz"
	TagLanguageVai    TagLanguage = "vai"
	TagLanguageVaiLr  TagLanguage = "vai-lr"
	TagLanguageVe     TagLanguage = "ve"
	TagLanguageVi     TagLanguage = "vi"
	TagLanguageViVn   TagLanguage = "vi-vn"
	TagLanguageVo     TagLanguage = "vo"
	TagLanguageVo001  TagLanguage = "vo-001"
	TagLanguageVun    TagLanguage = "vun"
	TagLanguageVunTz  TagLanguage = "vun-tz"
	TagLanguageWa     TagLanguage = "wa"
	TagLanguageWae    TagLanguage = "wae"
	TagLanguageWaeCh  TagLanguage = "wae-ch"
	TagLanguageWo     TagLanguage = "wo"
	TagLanguageWoSn   TagLanguage = "wo-sn"
	TagLanguageXh     TagLanguage = "xh"
	TagLanguageXhZa   TagLanguage = "xh-za"
	TagLanguageXog    TagLanguage = "xog"
	TagLanguageXogUg  TagLanguage = "xog-ug"
	TagLanguageYav    TagLanguage = "yav"
	TagLanguageYavCm  TagLanguage = "yav-cm"
	TagLanguageYi     TagLanguage = "yi"
	TagLanguageYi001  TagLanguage = "yi-001"
	TagLanguageYo     TagLanguage = "yo"
	TagLanguageYoBj   TagLanguage = "yo-bj"
	TagLanguageYoNg   TagLanguage = "yo-ng"
	TagLanguageYrl    TagLanguage = "yrl"
	TagLanguageYrlBr  TagLanguage = "yrl-br"
	TagLanguageYrlCo  TagLanguage = "yrl-co"
	TagLanguageYrlVe  TagLanguage = "yrl-ve"
	TagLanguageYue    TagLanguage = "yue"
	TagLanguageYueCn  TagLanguage = "yue-cn"
	TagLanguageYueHk  TagLanguage = "yue-hk"
	TagLanguageZa     TagLanguage = "za"
	TagLanguageZgh    TagLanguage = "zgh"
	TagLanguageZghMa  TagLanguage = "zgh-ma"
	TagLanguageZh     TagLanguage = "zh"
	TagLanguageZhCn   TagLanguage = "zh-cn"
	TagLanguageZhHans TagLanguage = "zh-hans"
	TagLanguageZhHant TagLanguage = "zh-hant"
	TagLanguageZhHk   TagLanguage = "zh-hk"
	TagLanguageZhMo   TagLanguage = "zh-mo"
	TagLanguageZhSg   TagLanguage = "zh-sg"
	TagLanguageZhTw   TagLanguage = "zh-tw"
	TagLanguageZu     TagLanguage = "zu"
	TagLanguageZuZa   TagLanguage = "zu-za"
)

// The properties ID, Name are required.
type TagCloneRequestVNextParam struct {
	// ID of the object to be cloned.
	ID string `json:"id" api:"required"`
	// Name of newly cloned blog tag.
	Name string `json:"name" api:"required"`
	// Target language of new variant.
	Language param.Opt[string] `json:"language,omitzero"`
	// Language of primary blog tag to clone.
	PrimaryLanguage param.Opt[string] `json:"primaryLanguage,omitzero"`
	paramObj
}

func (r TagCloneRequestVNextParam) MarshalJSON() (data []byte, err error) {
	type shadow TagCloneRequestVNextParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TagCloneRequestVNextParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BlogTagNewParams struct {
	Tag TagParam
	paramObj
}

func (r BlogTagNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.Tag)
}
func (r *BlogTagNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BlogTagUpdateParams struct {
	Tag TagParam
	// Whether to return only results that have been archived.
	Archived param.Opt[bool] `query:"archived,omitzero" json:"-"`
	paramObj
}

func (r BlogTagUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.Tag)
}
func (r *BlogTagUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [BlogTagUpdateParams]'s query parameters as `url.Values`.
func (r BlogTagUpdateParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type BlogTagListParams struct {
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

// URLQuery serializes [BlogTagListParams]'s query parameters as `url.Values`.
func (r BlogTagListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type BlogTagDeleteParams struct {
	// Whether to return only results that have been archived.
	Archived param.Opt[bool] `query:"archived,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [BlogTagDeleteParams]'s query parameters as `url.Values`.
func (r BlogTagDeleteParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type BlogTagAttachToLangGroupParams struct {
	AttachToLangPrimaryRequestVNext AttachToLangPrimaryRequestVNextParam
	paramObj
}

func (r BlogTagAttachToLangGroupParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.AttachToLangPrimaryRequestVNext)
}
func (r *BlogTagAttachToLangGroupParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BlogTagNewLangVariationParams struct {
	TagCloneRequestVNext TagCloneRequestVNextParam
	paramObj
}

func (r BlogTagNewLangVariationParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.TagCloneRequestVNext)
}
func (r *BlogTagNewLangVariationParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BlogTagDetachFromLangGroupParams struct {
	DetachFromLangGroupRequestVNext DetachFromLangGroupRequestVNextParam
	paramObj
}

func (r BlogTagDetachFromLangGroupParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.DetachFromLangGroupRequestVNext)
}
func (r *BlogTagDetachFromLangGroupParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BlogTagGetParams struct {
	// Whether to return only results that have been archived.
	Archived param.Opt[bool]   `query:"archived,omitzero" json:"-"`
	Property param.Opt[string] `query:"property,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [BlogTagGetParams]'s query parameters as `url.Values`.
func (r BlogTagGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type BlogTagListAuthorsCursorParams struct {
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

// URLQuery serializes [BlogTagListAuthorsCursorParams]'s query parameters as
// `url.Values`.
func (r BlogTagListAuthorsCursorParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type BlogTagListAuthorsCursorByQueryParams struct {
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

// URLQuery serializes [BlogTagListAuthorsCursorByQueryParams]'s query parameters
// as `url.Values`.
func (r BlogTagListAuthorsCursorByQueryParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type BlogTagListCursorParams struct {
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

// URLQuery serializes [BlogTagListCursorParams]'s query parameters as
// `url.Values`.
func (r BlogTagListCursorParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type BlogTagListCursorByQueryParams struct {
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

// URLQuery serializes [BlogTagListCursorByQueryParams]'s query parameters as
// `url.Values`.
func (r BlogTagListCursorByQueryParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type BlogTagListPostsCursorParams struct {
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

// URLQuery serializes [BlogTagListPostsCursorParams]'s query parameters as
// `url.Values`.
func (r BlogTagListPostsCursorParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type BlogTagListPostsCursorByQueryParams struct {
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

// URLQuery serializes [BlogTagListPostsCursorByQueryParams]'s query parameters as
// `url.Values`.
func (r BlogTagListPostsCursorByQueryParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type BlogTagSetLangPrimaryParams struct {
	SetNewLanguagePrimaryRequestVNext SetNewLanguagePrimaryRequestVNextParam
	paramObj
}

func (r BlogTagSetLangPrimaryParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.SetNewLanguagePrimaryRequestVNext)
}
func (r *BlogTagSetLangPrimaryParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BlogTagUpdateLangsParams struct {
	UpdateLanguagesRequestVNext UpdateLanguagesRequestVNextParam
	paramObj
}

func (r BlogTagUpdateLangsParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.UpdateLanguagesRequestVNext)
}
func (r *BlogTagUpdateLangsParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
