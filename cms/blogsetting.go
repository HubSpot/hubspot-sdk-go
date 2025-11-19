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
	"github.com/stainless-sdks/hubspot-sdk-go/marketing"
	"github.com/stainless-sdks/hubspot-sdk-go/option"
	"github.com/stainless-sdks/hubspot-sdk-go/packages/pagination"
	"github.com/stainless-sdks/hubspot-sdk-go/packages/param"
	"github.com/stainless-sdks/hubspot-sdk-go/packages/respjson"
	"github.com/stainless-sdks/hubspot-sdk-go/shared"
)

// BlogSettingService contains methods and other services that help with
// interacting with the Hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBlogSettingService] method instead.
type BlogSettingService struct {
	Options []option.RequestOption
}

// NewBlogSettingService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewBlogSettingService(opts ...option.RequestOption) (r BlogSettingService) {
	r = BlogSettingService{}
	r.Options = opts
	return
}

// Get the list of Blogs. Supports paging and filtering. This method would be
// useful for an integration that examined these models and used an external
// service to suggest edits.
func (r *BlogSettingService) List(ctx context.Context, query BlogSettingListParams, opts ...option.RequestOption) (res *pagination.Page[Blog], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "cms/v3/blog-settings/settings"
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

// Get the list of Blogs. Supports paging and filtering. This method would be
// useful for an integration that examined these models and used an external
// service to suggest edits.
func (r *BlogSettingService) ListAutoPaging(ctx context.Context, query BlogSettingListParams, opts ...option.RequestOption) *pagination.PageAutoPager[Blog] {
	return pagination.NewPageAutoPager(r.List(ctx, query, opts...))
}

// Attach a blog to a multi-language group.
func (r *BlogSettingService) AttachToLangGroup(ctx context.Context, body BlogSettingAttachToLangGroupParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "cms/v3/blog-settings/settings/multi-language/attach-to-lang-group"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Create a new language variation from an existing blog
func (r *BlogSettingService) NewLanguageVariation(ctx context.Context, body BlogSettingNewLanguageVariationParams, opts ...option.RequestOption) (res *Blog, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "cms/v3/blog-settings/settings/multi-language/create-language-variation"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Detach a blog from a multi-language group.
func (r *BlogSettingService) DetachFromLangGroup(ctx context.Context, body BlogSettingDetachFromLangGroupParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "cms/v3/blog-settings/settings/multi-language/detach-from-lang-group"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Retrieve the Blog object identified by the id in the path.
func (r *BlogSettingService) Get(ctx context.Context, blogID string, opts ...option.RequestOption) (res *Blog, err error) {
	opts = slices.Concat(r.Options, opts)
	if blogID == "" {
		err = errors.New("missing required blogId parameter")
		return
	}
	path := fmt.Sprintf("cms/v3/blog-settings/settings/%s", blogID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Retrieves a previous version of a Blog
func (r *BlogSettingService) GetRevision(ctx context.Context, revisionID string, query BlogSettingGetRevisionParams, opts ...option.RequestOption) (res *VersionBlog, err error) {
	opts = slices.Concat(r.Options, opts)
	if query.BlogID == "" {
		err = errors.New("missing required blogId parameter")
		return
	}
	if revisionID == "" {
		err = errors.New("missing required revisionId parameter")
		return
	}
	path := fmt.Sprintf("cms/v3/blog-settings/settings/%s/revisions/%s", query.BlogID, revisionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Retrieves all the previous versions of a Blog
func (r *BlogSettingService) ListRevisions(ctx context.Context, blogID string, query BlogSettingListRevisionsParams, opts ...option.RequestOption) (res *CollectionResponseWithTotalVersionBlog, err error) {
	opts = slices.Concat(r.Options, opts)
	if blogID == "" {
		err = errors.New("missing required blogId parameter")
		return
	}
	path := fmt.Sprintf("cms/v3/blog-settings/settings/%s/revisions", blogID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Set a blog as the primary language of a multi-language group.
func (r *BlogSettingService) SetNewLangPrimary(ctx context.Context, body BlogSettingSetNewLangPrimaryParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "cms/v3/blog-settings/settings/multi-language/set-new-lang-primary"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, nil, opts...)
	return
}

// Explicitly set new languages for each blog in a multi-language group.
func (r *BlogSettingService) UpdateLanguages(ctx context.Context, body BlogSettingUpdateLanguagesParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "cms/v3/blog-settings/settings/multi-language/update-languages"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

type Blog struct {
	// The unique ID of the Blog.
	ID          string `json:"id,required"`
	AbsoluteURL string `json:"absoluteUrl,required"`
	// Boolean determining whether or not this blog allows public comments.
	AllowComments bool      `json:"allowComments,required"`
	Created       time.Time `json:"created,required" format:"date-time"`
	// The timestamp (ISO8601 format) when this Blog was deleted.
	DeletedAt time.Time `json:"deletedAt,required" format:"date-time"`
	// The Description of this Blog.
	Description string `json:"description,required"`
	// The html title of this Blog.
	HTMLTitle string `json:"htmlTitle,required"`
	// The explicitly defined language of the Blog. If null, the Blog will default to
	// the language of the Domain.
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
	// "gv-im", "ha", "ha-gh", "ha-ne", "ha-ng", "haw", "haw-us", "he", "hi", "hi-in",
	// "hr", "hr-ba", "hr-hr", "hsb", "hsb-de", "hu", "hu-hu", "hy", "hy-am", "ia",
	// "ia-001", "id", "ig", "ig-ng", "ii", "ii-cn", "id-id", "is", "is-is", "it",
	// "it-ch", "it-it", "it-sm", "it-va", "he-il", "ja", "ja-jp", "jgo", "jgo-cm",
	// "yi", "yi-001", "jmc", "jmc-tz", "jv", "jv-id", "ka", "ka-ge", "kab", "kab-dz",
	// "kam", "kam-ke", "kde", "kde-tz", "kea", "kea-cv", "khq", "khq-ml", "ki",
	// "ki-ke", "kk", "kk-kz", "kkj", "kkj-cm", "kl", "kl-gl", "kln", "kln-ke", "km",
	// "km-kh", "kn", "kn-in", "ko", "ko-kp", "ko-kr", "kok", "kok-in", "ks", "ks-in",
	// "ksb", "ksb-tz", "ksf", "ksf-cm", "ksh", "ksh-de", "kw", "kw-gb", "ku", "ku-tr",
	// "ky", "ky-kg", "lag", "lag-tz", "lb", "lb-lu", "lg", "lg-ug", "lkt", "lkt-us",
	// "ln", "ln-ao", "ln-cd", "ln-cf", "ln-cg", "lo", "lo-la", "lrc", "lrc-iq",
	// "lrc-ir", "lt", "lt-lt", "lu", "lu-cd", "luo", "luo-ke", "luy", "luy-ke", "lv",
	// "lv-lv", "mai", "mai-in", "mas", "mas-ke", "mas-tz", "mer", "mer-ke", "mfe",
	// "mfe-mu", "mg", "mg-mg", "mgh", "mgh-mz", "mgo", "mgo-cm", "mi", "mi-nz", "mk",
	// "mk-mk", "ml", "ml-in", "mn", "mn-mn", "mni", "mni-in", "mr", "mr-in", "ms",
	// "ms-bn", "ms-id", "ms-my", "ms-sg", "mt", "mt-mt", "mua", "mua-cm", "my",
	// "my-mm", "mzn", "mzn-ir", "naq", "naq-na", "nb", "nb-no", "nb-sj", "nd",
	// "nd-zw", "nds", "nds-de", "nds-nl", "ne", "ne-in", "ne-np", "nl", "nl-aw",
	// "nl-be", "nl-ch", "nl-bq", "nl-cw", "nl-lu", "nl-nl", "nl-sr", "nl-sx", "nmg",
	// "nmg-cm", "nn", "nn-no", "nnh", "nnh-cm", "no", "no-no", "nus", "nus-ss", "nyn",
	// "nyn-ug", "om", "om-et", "om-ke", "or", "or-in", "os", "os-ge", "os-ru", "pa",
	// "pa-in", "pa-pk", "pcm", "pcm-ng", "pl", "pl-pl", "prg", "prg-001", "ps",
	// "ps-af", "ps-pk", "pt", "pt-ao", "pt-br", "pt-ch", "pt-cv", "pt-gq", "pt-gw",
	// "pt-lu", "pt-mo", "pt-mz", "pt-pt", "pt-st", "pt-tl", "qu", "qu-bo", "qu-ec",
	// "qu-pe", "rm", "rm-ch", "rn", "rn-bi", "ro", "ro-md", "ro-ro", "rof", "rof-tz",
	// "ru", "ru-by", "ru-kg", "ru-kz", "ru-md", "ru-ru", "ru-ua", "rw", "rw-rw",
	// "rwk", "rwk-tz", "sa", "sa-in", "sah", "sah-ru", "saq", "saq-ke", "sat",
	// "sat-in", "sbp", "sbp-tz", "sd", "sd-in", "sd-pk", "se", "se-fi", "se-no",
	// "se-se", "seh", "seh-mz", "ses", "ses-ml", "sg", "sg-cf", "shi", "shi-ma", "si",
	// "si-lk", "sk", "sk-sk", "sl", "sl-si", "smn", "smn-fi", "sn", "sn-zw", "so",
	// "so-dj", "so-et", "so-ke", "so-so", "sq", "sq-al", "sq-mk", "sq-xk", "sr",
	// "sr-ba", "sr-cs", "sr-me", "sr-rs", "sr-xk", "su", "su-id", "sv", "sv-ax",
	// "sv-fi", "sv-se", "sw", "sw-cd", "sw-ke", "sw-tz", "sw-ug", "sy", "ta", "ta-in",
	// "ta-lk", "ta-my", "ta-sg", "te", "te-in", "teo", "teo-ke", "teo-ug", "tg",
	// "tg-tj", "th", "th-th", "ti", "ti-er", "ti-et", "tk", "tk-tm", "tl", "to",
	// "to-to", "tr", "tr-cy", "tr-tr", "tt", "tt-ru", "twq", "twq-ne", "tzm",
	// "tzm-ma", "ug", "ug-cn", "uk", "uk-ua", "ur", "ur-in", "ur-pk", "uz", "uz-af",
	// "uz-uz", "vai", "vai-lr", "vi", "vi-vn", "vo", "vo-001", "vun", "vun-tz", "wae",
	// "wae-ch", "wo", "wo-sn", "xh", "xh-za", "xog", "xog-ug", "yav", "yav-cm", "yo",
	// "yo-bj", "yo-ng", "yue", "yue-cn", "yue-hk", "zgh", "zgh-ma", "zh", "zh-cn",
	// "zh-hk", "zh-mo", "zh-sg", "zh-tw", "zh-hans", "zh-hant", "zu", "zu-za".
	Language BlogLanguage `json:"language,required"`
	// The internal name of the blog.
	Name string `json:"name,required"`
	// Rules for require member registration to access private content.
	PublicAccessRules []PublicAccessRule `json:"publicAccessRules,required"`
	// Boolean to determine whether or not to respect publicAccessRules.
	PublicAccessRulesEnabled bool `json:"publicAccessRulesEnabled,required"`
	// The public title of this Blog.
	PublicTitle string `json:"publicTitle,required"`
	// The path of the this blog. This field is appended to the domain to construct the
	// url of this blog.
	Slug string `json:"slug,required"`
	// ID of the primary Blog this object was translated from.
	TranslatedFromID string    `json:"translatedFromId,required"`
	Updated          time.Time `json:"updated,required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                       respjson.Field
		AbsoluteURL              respjson.Field
		AllowComments            respjson.Field
		Created                  respjson.Field
		DeletedAt                respjson.Field
		Description              respjson.Field
		HTMLTitle                respjson.Field
		Language                 respjson.Field
		Name                     respjson.Field
		PublicAccessRules        respjson.Field
		PublicAccessRulesEnabled respjson.Field
		PublicTitle              respjson.Field
		Slug                     respjson.Field
		TranslatedFromID         respjson.Field
		Updated                  respjson.Field
		ExtraFields              map[string]respjson.Field
		raw                      string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Blog) RawJSON() string { return r.JSON.raw }
func (r *Blog) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The explicitly defined language of the Blog. If null, the Blog will default to
// the language of the Domain.
type BlogLanguage string

const (
	BlogLanguageAf     BlogLanguage = "af"
	BlogLanguageAfNa   BlogLanguage = "af-na"
	BlogLanguageAfZa   BlogLanguage = "af-za"
	BlogLanguageAgq    BlogLanguage = "agq"
	BlogLanguageAgqCm  BlogLanguage = "agq-cm"
	BlogLanguageAk     BlogLanguage = "ak"
	BlogLanguageAkGh   BlogLanguage = "ak-gh"
	BlogLanguageAm     BlogLanguage = "am"
	BlogLanguageAmEt   BlogLanguage = "am-et"
	BlogLanguageAr     BlogLanguage = "ar"
	BlogLanguageAr001  BlogLanguage = "ar-001"
	BlogLanguageArAe   BlogLanguage = "ar-ae"
	BlogLanguageArBh   BlogLanguage = "ar-bh"
	BlogLanguageArDj   BlogLanguage = "ar-dj"
	BlogLanguageArDz   BlogLanguage = "ar-dz"
	BlogLanguageArEg   BlogLanguage = "ar-eg"
	BlogLanguageArEh   BlogLanguage = "ar-eh"
	BlogLanguageArEr   BlogLanguage = "ar-er"
	BlogLanguageArIl   BlogLanguage = "ar-il"
	BlogLanguageArIq   BlogLanguage = "ar-iq"
	BlogLanguageArJo   BlogLanguage = "ar-jo"
	BlogLanguageArKm   BlogLanguage = "ar-km"
	BlogLanguageArKw   BlogLanguage = "ar-kw"
	BlogLanguageArLb   BlogLanguage = "ar-lb"
	BlogLanguageArLy   BlogLanguage = "ar-ly"
	BlogLanguageArMa   BlogLanguage = "ar-ma"
	BlogLanguageArMr   BlogLanguage = "ar-mr"
	BlogLanguageArOm   BlogLanguage = "ar-om"
	BlogLanguageArPs   BlogLanguage = "ar-ps"
	BlogLanguageArQa   BlogLanguage = "ar-qa"
	BlogLanguageArSa   BlogLanguage = "ar-sa"
	BlogLanguageArSd   BlogLanguage = "ar-sd"
	BlogLanguageArSo   BlogLanguage = "ar-so"
	BlogLanguageArSS   BlogLanguage = "ar-ss"
	BlogLanguageArSy   BlogLanguage = "ar-sy"
	BlogLanguageArTd   BlogLanguage = "ar-td"
	BlogLanguageArTn   BlogLanguage = "ar-tn"
	BlogLanguageArYe   BlogLanguage = "ar-ye"
	BlogLanguageAs     BlogLanguage = "as"
	BlogLanguageAsIn   BlogLanguage = "as-in"
	BlogLanguageAsa    BlogLanguage = "asa"
	BlogLanguageAsaTz  BlogLanguage = "asa-tz"
	BlogLanguageAst    BlogLanguage = "ast"
	BlogLanguageAstEs  BlogLanguage = "ast-es"
	BlogLanguageAz     BlogLanguage = "az"
	BlogLanguageAzAz   BlogLanguage = "az-az"
	BlogLanguageBas    BlogLanguage = "bas"
	BlogLanguageBasCm  BlogLanguage = "bas-cm"
	BlogLanguageBe     BlogLanguage = "be"
	BlogLanguageBeBy   BlogLanguage = "be-by"
	BlogLanguageBem    BlogLanguage = "bem"
	BlogLanguageBemZm  BlogLanguage = "bem-zm"
	BlogLanguageBez    BlogLanguage = "bez"
	BlogLanguageBezTz  BlogLanguage = "bez-tz"
	BlogLanguageBg     BlogLanguage = "bg"
	BlogLanguageBgBg   BlogLanguage = "bg-bg"
	BlogLanguageBm     BlogLanguage = "bm"
	BlogLanguageBmMl   BlogLanguage = "bm-ml"
	BlogLanguageBn     BlogLanguage = "bn"
	BlogLanguageBnBd   BlogLanguage = "bn-bd"
	BlogLanguageBnIn   BlogLanguage = "bn-in"
	BlogLanguageBo     BlogLanguage = "bo"
	BlogLanguageBoCn   BlogLanguage = "bo-cn"
	BlogLanguageBoIn   BlogLanguage = "bo-in"
	BlogLanguageBr     BlogLanguage = "br"
	BlogLanguageBrFr   BlogLanguage = "br-fr"
	BlogLanguageBrx    BlogLanguage = "brx"
	BlogLanguageBrxIn  BlogLanguage = "brx-in"
	BlogLanguageBs     BlogLanguage = "bs"
	BlogLanguageBsBa   BlogLanguage = "bs-ba"
	BlogLanguageCa     BlogLanguage = "ca"
	BlogLanguageCaAd   BlogLanguage = "ca-ad"
	BlogLanguageCaEs   BlogLanguage = "ca-es"
	BlogLanguageCaFr   BlogLanguage = "ca-fr"
	BlogLanguageCaIt   BlogLanguage = "ca-it"
	BlogLanguageCcp    BlogLanguage = "ccp"
	BlogLanguageCcpBd  BlogLanguage = "ccp-bd"
	BlogLanguageCcpIn  BlogLanguage = "ccp-in"
	BlogLanguageCe     BlogLanguage = "ce"
	BlogLanguageCeRu   BlogLanguage = "ce-ru"
	BlogLanguageCeb    BlogLanguage = "ceb"
	BlogLanguageCebPh  BlogLanguage = "ceb-ph"
	BlogLanguageCgg    BlogLanguage = "cgg"
	BlogLanguageCggUg  BlogLanguage = "cgg-ug"
	BlogLanguageChr    BlogLanguage = "chr"
	BlogLanguageChrUs  BlogLanguage = "chr-us"
	BlogLanguageCkb    BlogLanguage = "ckb"
	BlogLanguageCkbIq  BlogLanguage = "ckb-iq"
	BlogLanguageCkbIr  BlogLanguage = "ckb-ir"
	BlogLanguageCs     BlogLanguage = "cs"
	BlogLanguageCsCz   BlogLanguage = "cs-cz"
	BlogLanguageCu     BlogLanguage = "cu"
	BlogLanguageCuRu   BlogLanguage = "cu-ru"
	BlogLanguageCy     BlogLanguage = "cy"
	BlogLanguageCyGB   BlogLanguage = "cy-gb"
	BlogLanguageDa     BlogLanguage = "da"
	BlogLanguageDaDk   BlogLanguage = "da-dk"
	BlogLanguageDaGl   BlogLanguage = "da-gl"
	BlogLanguageDav    BlogLanguage = "dav"
	BlogLanguageDavKe  BlogLanguage = "dav-ke"
	BlogLanguageDe     BlogLanguage = "de"
	BlogLanguageDeAt   BlogLanguage = "de-at"
	BlogLanguageDeBe   BlogLanguage = "de-be"
	BlogLanguageDeCh   BlogLanguage = "de-ch"
	BlogLanguageDeDe   BlogLanguage = "de-de"
	BlogLanguageDeGr   BlogLanguage = "de-gr"
	BlogLanguageDeIt   BlogLanguage = "de-it"
	BlogLanguageDeLi   BlogLanguage = "de-li"
	BlogLanguageDeLu   BlogLanguage = "de-lu"
	BlogLanguageDje    BlogLanguage = "dje"
	BlogLanguageDjeNe  BlogLanguage = "dje-ne"
	BlogLanguageDoi    BlogLanguage = "doi"
	BlogLanguageDoiIn  BlogLanguage = "doi-in"
	BlogLanguageDsb    BlogLanguage = "dsb"
	BlogLanguageDsbDe  BlogLanguage = "dsb-de"
	BlogLanguageDua    BlogLanguage = "dua"
	BlogLanguageDuaCm  BlogLanguage = "dua-cm"
	BlogLanguageDyo    BlogLanguage = "dyo"
	BlogLanguageDyoSn  BlogLanguage = "dyo-sn"
	BlogLanguageDz     BlogLanguage = "dz"
	BlogLanguageDzBt   BlogLanguage = "dz-bt"
	BlogLanguageEbu    BlogLanguage = "ebu"
	BlogLanguageEbuKe  BlogLanguage = "ebu-ke"
	BlogLanguageEe     BlogLanguage = "ee"
	BlogLanguageEeGh   BlogLanguage = "ee-gh"
	BlogLanguageEeTg   BlogLanguage = "ee-tg"
	BlogLanguageEl     BlogLanguage = "el"
	BlogLanguageElCy   BlogLanguage = "el-cy"
	BlogLanguageElGr   BlogLanguage = "el-gr"
	BlogLanguageEn     BlogLanguage = "en"
	BlogLanguageEn001  BlogLanguage = "en-001"
	BlogLanguageEn150  BlogLanguage = "en-150"
	BlogLanguageEnAe   BlogLanguage = "en-ae"
	BlogLanguageEnAg   BlogLanguage = "en-ag"
	BlogLanguageEnAI   BlogLanguage = "en-ai"
	BlogLanguageEnAs   BlogLanguage = "en-as"
	BlogLanguageEnAt   BlogLanguage = "en-at"
	BlogLanguageEnAu   BlogLanguage = "en-au"
	BlogLanguageEnBb   BlogLanguage = "en-bb"
	BlogLanguageEnBe   BlogLanguage = "en-be"
	BlogLanguageEnBi   BlogLanguage = "en-bi"
	BlogLanguageEnBm   BlogLanguage = "en-bm"
	BlogLanguageEnBs   BlogLanguage = "en-bs"
	BlogLanguageEnBw   BlogLanguage = "en-bw"
	BlogLanguageEnBz   BlogLanguage = "en-bz"
	BlogLanguageEnCa   BlogLanguage = "en-ca"
	BlogLanguageEnCc   BlogLanguage = "en-cc"
	BlogLanguageEnCh   BlogLanguage = "en-ch"
	BlogLanguageEnCk   BlogLanguage = "en-ck"
	BlogLanguageEnCm   BlogLanguage = "en-cm"
	BlogLanguageEnCn   BlogLanguage = "en-cn"
	BlogLanguageEnCx   BlogLanguage = "en-cx"
	BlogLanguageEnCy   BlogLanguage = "en-cy"
	BlogLanguageEnDe   BlogLanguage = "en-de"
	BlogLanguageEnDg   BlogLanguage = "en-dg"
	BlogLanguageEnDk   BlogLanguage = "en-dk"
	BlogLanguageEnDm   BlogLanguage = "en-dm"
	BlogLanguageEnEr   BlogLanguage = "en-er"
	BlogLanguageEnFi   BlogLanguage = "en-fi"
	BlogLanguageEnFj   BlogLanguage = "en-fj"
	BlogLanguageEnFk   BlogLanguage = "en-fk"
	BlogLanguageEnFm   BlogLanguage = "en-fm"
	BlogLanguageEnGB   BlogLanguage = "en-gb"
	BlogLanguageEnGd   BlogLanguage = "en-gd"
	BlogLanguageEnGg   BlogLanguage = "en-gg"
	BlogLanguageEnGh   BlogLanguage = "en-gh"
	BlogLanguageEnGi   BlogLanguage = "en-gi"
	BlogLanguageEnGm   BlogLanguage = "en-gm"
	BlogLanguageEnGu   BlogLanguage = "en-gu"
	BlogLanguageEnGy   BlogLanguage = "en-gy"
	BlogLanguageEnHk   BlogLanguage = "en-hk"
	BlogLanguageEnIe   BlogLanguage = "en-ie"
	BlogLanguageEnIl   BlogLanguage = "en-il"
	BlogLanguageEnIm   BlogLanguage = "en-im"
	BlogLanguageEnIn   BlogLanguage = "en-in"
	BlogLanguageEnIo   BlogLanguage = "en-io"
	BlogLanguageEnJe   BlogLanguage = "en-je"
	BlogLanguageEnJm   BlogLanguage = "en-jm"
	BlogLanguageEnKe   BlogLanguage = "en-ke"
	BlogLanguageEnKi   BlogLanguage = "en-ki"
	BlogLanguageEnKn   BlogLanguage = "en-kn"
	BlogLanguageEnKy   BlogLanguage = "en-ky"
	BlogLanguageEnLc   BlogLanguage = "en-lc"
	BlogLanguageEnLr   BlogLanguage = "en-lr"
	BlogLanguageEnLs   BlogLanguage = "en-ls"
	BlogLanguageEnLu   BlogLanguage = "en-lu"
	BlogLanguageEnMg   BlogLanguage = "en-mg"
	BlogLanguageEnMh   BlogLanguage = "en-mh"
	BlogLanguageEnMo   BlogLanguage = "en-mo"
	BlogLanguageEnMp   BlogLanguage = "en-mp"
	BlogLanguageEnMs   BlogLanguage = "en-ms"
	BlogLanguageEnMt   BlogLanguage = "en-mt"
	BlogLanguageEnMu   BlogLanguage = "en-mu"
	BlogLanguageEnMw   BlogLanguage = "en-mw"
	BlogLanguageEnMx   BlogLanguage = "en-mx"
	BlogLanguageEnMy   BlogLanguage = "en-my"
	BlogLanguageEnNa   BlogLanguage = "en-na"
	BlogLanguageEnNf   BlogLanguage = "en-nf"
	BlogLanguageEnNg   BlogLanguage = "en-ng"
	BlogLanguageEnNl   BlogLanguage = "en-nl"
	BlogLanguageEnNr   BlogLanguage = "en-nr"
	BlogLanguageEnNu   BlogLanguage = "en-nu"
	BlogLanguageEnNz   BlogLanguage = "en-nz"
	BlogLanguageEnPg   BlogLanguage = "en-pg"
	BlogLanguageEnPh   BlogLanguage = "en-ph"
	BlogLanguageEnPk   BlogLanguage = "en-pk"
	BlogLanguageEnPn   BlogLanguage = "en-pn"
	BlogLanguageEnPr   BlogLanguage = "en-pr"
	BlogLanguageEnPw   BlogLanguage = "en-pw"
	BlogLanguageEnRw   BlogLanguage = "en-rw"
	BlogLanguageEnSb   BlogLanguage = "en-sb"
	BlogLanguageEnSc   BlogLanguage = "en-sc"
	BlogLanguageEnSd   BlogLanguage = "en-sd"
	BlogLanguageEnSe   BlogLanguage = "en-se"
	BlogLanguageEnSg   BlogLanguage = "en-sg"
	BlogLanguageEnSh   BlogLanguage = "en-sh"
	BlogLanguageEnSi   BlogLanguage = "en-si"
	BlogLanguageEnSl   BlogLanguage = "en-sl"
	BlogLanguageEnSS   BlogLanguage = "en-ss"
	BlogLanguageEnSx   BlogLanguage = "en-sx"
	BlogLanguageEnSz   BlogLanguage = "en-sz"
	BlogLanguageEnTc   BlogLanguage = "en-tc"
	BlogLanguageEnTk   BlogLanguage = "en-tk"
	BlogLanguageEnTo   BlogLanguage = "en-to"
	BlogLanguageEnTt   BlogLanguage = "en-tt"
	BlogLanguageEnTv   BlogLanguage = "en-tv"
	BlogLanguageEnTz   BlogLanguage = "en-tz"
	BlogLanguageEnUg   BlogLanguage = "en-ug"
	BlogLanguageEnUm   BlogLanguage = "en-um"
	BlogLanguageEnUs   BlogLanguage = "en-us"
	BlogLanguageEnVc   BlogLanguage = "en-vc"
	BlogLanguageEnVg   BlogLanguage = "en-vg"
	BlogLanguageEnVi   BlogLanguage = "en-vi"
	BlogLanguageEnVu   BlogLanguage = "en-vu"
	BlogLanguageEnWs   BlogLanguage = "en-ws"
	BlogLanguageEnZa   BlogLanguage = "en-za"
	BlogLanguageEnZm   BlogLanguage = "en-zm"
	BlogLanguageEnZw   BlogLanguage = "en-zw"
	BlogLanguageEo     BlogLanguage = "eo"
	BlogLanguageEo001  BlogLanguage = "eo-001"
	BlogLanguageEs     BlogLanguage = "es"
	BlogLanguageEs419  BlogLanguage = "es-419"
	BlogLanguageEsAr   BlogLanguage = "es-ar"
	BlogLanguageEsBo   BlogLanguage = "es-bo"
	BlogLanguageEsBr   BlogLanguage = "es-br"
	BlogLanguageEsBz   BlogLanguage = "es-bz"
	BlogLanguageEsCl   BlogLanguage = "es-cl"
	BlogLanguageEsCo   BlogLanguage = "es-co"
	BlogLanguageEsCr   BlogLanguage = "es-cr"
	BlogLanguageEsCu   BlogLanguage = "es-cu"
	BlogLanguageEsDo   BlogLanguage = "es-do"
	BlogLanguageEsEa   BlogLanguage = "es-ea"
	BlogLanguageEsEc   BlogLanguage = "es-ec"
	BlogLanguageEsEs   BlogLanguage = "es-es"
	BlogLanguageEsGq   BlogLanguage = "es-gq"
	BlogLanguageEsGt   BlogLanguage = "es-gt"
	BlogLanguageEsHn   BlogLanguage = "es-hn"
	BlogLanguageEsIc   BlogLanguage = "es-ic"
	BlogLanguageEsMx   BlogLanguage = "es-mx"
	BlogLanguageEsNi   BlogLanguage = "es-ni"
	BlogLanguageEsPa   BlogLanguage = "es-pa"
	BlogLanguageEsPe   BlogLanguage = "es-pe"
	BlogLanguageEsPh   BlogLanguage = "es-ph"
	BlogLanguageEsPr   BlogLanguage = "es-pr"
	BlogLanguageEsPy   BlogLanguage = "es-py"
	BlogLanguageEsSv   BlogLanguage = "es-sv"
	BlogLanguageEsUs   BlogLanguage = "es-us"
	BlogLanguageEsUy   BlogLanguage = "es-uy"
	BlogLanguageEsVe   BlogLanguage = "es-ve"
	BlogLanguageEt     BlogLanguage = "et"
	BlogLanguageEtEe   BlogLanguage = "et-ee"
	BlogLanguageEu     BlogLanguage = "eu"
	BlogLanguageEuEs   BlogLanguage = "eu-es"
	BlogLanguageEwo    BlogLanguage = "ewo"
	BlogLanguageEwoCm  BlogLanguage = "ewo-cm"
	BlogLanguageFa     BlogLanguage = "fa"
	BlogLanguageFaAf   BlogLanguage = "fa-af"
	BlogLanguageFaIr   BlogLanguage = "fa-ir"
	BlogLanguageFf     BlogLanguage = "ff"
	BlogLanguageFfBf   BlogLanguage = "ff-bf"
	BlogLanguageFfCm   BlogLanguage = "ff-cm"
	BlogLanguageFfGh   BlogLanguage = "ff-gh"
	BlogLanguageFfGm   BlogLanguage = "ff-gm"
	BlogLanguageFfGn   BlogLanguage = "ff-gn"
	BlogLanguageFfGw   BlogLanguage = "ff-gw"
	BlogLanguageFfLr   BlogLanguage = "ff-lr"
	BlogLanguageFfMr   BlogLanguage = "ff-mr"
	BlogLanguageFfNe   BlogLanguage = "ff-ne"
	BlogLanguageFfNg   BlogLanguage = "ff-ng"
	BlogLanguageFfSl   BlogLanguage = "ff-sl"
	BlogLanguageFfSn   BlogLanguage = "ff-sn"
	BlogLanguageFi     BlogLanguage = "fi"
	BlogLanguageFiFi   BlogLanguage = "fi-fi"
	BlogLanguageFil    BlogLanguage = "fil"
	BlogLanguageFilPh  BlogLanguage = "fil-ph"
	BlogLanguageFo     BlogLanguage = "fo"
	BlogLanguageFoDk   BlogLanguage = "fo-dk"
	BlogLanguageFoFo   BlogLanguage = "fo-fo"
	BlogLanguageFr     BlogLanguage = "fr"
	BlogLanguageFrBe   BlogLanguage = "fr-be"
	BlogLanguageFrBf   BlogLanguage = "fr-bf"
	BlogLanguageFrBi   BlogLanguage = "fr-bi"
	BlogLanguageFrBj   BlogLanguage = "fr-bj"
	BlogLanguageFrBl   BlogLanguage = "fr-bl"
	BlogLanguageFrCa   BlogLanguage = "fr-ca"
	BlogLanguageFrCd   BlogLanguage = "fr-cd"
	BlogLanguageFrCf   BlogLanguage = "fr-cf"
	BlogLanguageFrCg   BlogLanguage = "fr-cg"
	BlogLanguageFrCh   BlogLanguage = "fr-ch"
	BlogLanguageFrCi   BlogLanguage = "fr-ci"
	BlogLanguageFrCm   BlogLanguage = "fr-cm"
	BlogLanguageFrDj   BlogLanguage = "fr-dj"
	BlogLanguageFrDz   BlogLanguage = "fr-dz"
	BlogLanguageFrFr   BlogLanguage = "fr-fr"
	BlogLanguageFrGa   BlogLanguage = "fr-ga"
	BlogLanguageFrGf   BlogLanguage = "fr-gf"
	BlogLanguageFrGn   BlogLanguage = "fr-gn"
	BlogLanguageFrGp   BlogLanguage = "fr-gp"
	BlogLanguageFrGq   BlogLanguage = "fr-gq"
	BlogLanguageFrHt   BlogLanguage = "fr-ht"
	BlogLanguageFrKm   BlogLanguage = "fr-km"
	BlogLanguageFrLu   BlogLanguage = "fr-lu"
	BlogLanguageFrMa   BlogLanguage = "fr-ma"
	BlogLanguageFrMc   BlogLanguage = "fr-mc"
	BlogLanguageFrMf   BlogLanguage = "fr-mf"
	BlogLanguageFrMg   BlogLanguage = "fr-mg"
	BlogLanguageFrMl   BlogLanguage = "fr-ml"
	BlogLanguageFrMq   BlogLanguage = "fr-mq"
	BlogLanguageFrMr   BlogLanguage = "fr-mr"
	BlogLanguageFrMu   BlogLanguage = "fr-mu"
	BlogLanguageFrNc   BlogLanguage = "fr-nc"
	BlogLanguageFrNe   BlogLanguage = "fr-ne"
	BlogLanguageFrPf   BlogLanguage = "fr-pf"
	BlogLanguageFrPm   BlogLanguage = "fr-pm"
	BlogLanguageFrRe   BlogLanguage = "fr-re"
	BlogLanguageFrRw   BlogLanguage = "fr-rw"
	BlogLanguageFrSc   BlogLanguage = "fr-sc"
	BlogLanguageFrSn   BlogLanguage = "fr-sn"
	BlogLanguageFrSy   BlogLanguage = "fr-sy"
	BlogLanguageFrTd   BlogLanguage = "fr-td"
	BlogLanguageFrTg   BlogLanguage = "fr-tg"
	BlogLanguageFrTn   BlogLanguage = "fr-tn"
	BlogLanguageFrVu   BlogLanguage = "fr-vu"
	BlogLanguageFrWf   BlogLanguage = "fr-wf"
	BlogLanguageFrYt   BlogLanguage = "fr-yt"
	BlogLanguageFur    BlogLanguage = "fur"
	BlogLanguageFurIt  BlogLanguage = "fur-it"
	BlogLanguageFy     BlogLanguage = "fy"
	BlogLanguageFyNl   BlogLanguage = "fy-nl"
	BlogLanguageGa     BlogLanguage = "ga"
	BlogLanguageGaGB   BlogLanguage = "ga-gb"
	BlogLanguageGaIe   BlogLanguage = "ga-ie"
	BlogLanguageGd     BlogLanguage = "gd"
	BlogLanguageGdGB   BlogLanguage = "gd-gb"
	BlogLanguageGl     BlogLanguage = "gl"
	BlogLanguageGlEs   BlogLanguage = "gl-es"
	BlogLanguageGsw    BlogLanguage = "gsw"
	BlogLanguageGswCh  BlogLanguage = "gsw-ch"
	BlogLanguageGswFr  BlogLanguage = "gsw-fr"
	BlogLanguageGswLi  BlogLanguage = "gsw-li"
	BlogLanguageGu     BlogLanguage = "gu"
	BlogLanguageGuIn   BlogLanguage = "gu-in"
	BlogLanguageGuz    BlogLanguage = "guz"
	BlogLanguageGuzKe  BlogLanguage = "guz-ke"
	BlogLanguageGv     BlogLanguage = "gv"
	BlogLanguageGvIm   BlogLanguage = "gv-im"
	BlogLanguageHa     BlogLanguage = "ha"
	BlogLanguageHaGh   BlogLanguage = "ha-gh"
	BlogLanguageHaNe   BlogLanguage = "ha-ne"
	BlogLanguageHaNg   BlogLanguage = "ha-ng"
	BlogLanguageHaw    BlogLanguage = "haw"
	BlogLanguageHawUs  BlogLanguage = "haw-us"
	BlogLanguageHe     BlogLanguage = "he"
	BlogLanguageHi     BlogLanguage = "hi"
	BlogLanguageHiIn   BlogLanguage = "hi-in"
	BlogLanguageHr     BlogLanguage = "hr"
	BlogLanguageHrBa   BlogLanguage = "hr-ba"
	BlogLanguageHrHr   BlogLanguage = "hr-hr"
	BlogLanguageHsb    BlogLanguage = "hsb"
	BlogLanguageHsbDe  BlogLanguage = "hsb-de"
	BlogLanguageHu     BlogLanguage = "hu"
	BlogLanguageHuHu   BlogLanguage = "hu-hu"
	BlogLanguageHy     BlogLanguage = "hy"
	BlogLanguageHyAm   BlogLanguage = "hy-am"
	BlogLanguageIa     BlogLanguage = "ia"
	BlogLanguageIa001  BlogLanguage = "ia-001"
	BlogLanguageID     BlogLanguage = "id"
	BlogLanguageIg     BlogLanguage = "ig"
	BlogLanguageIgNg   BlogLanguage = "ig-ng"
	BlogLanguageIi     BlogLanguage = "ii"
	BlogLanguageIiCn   BlogLanguage = "ii-cn"
	BlogLanguageIDID   BlogLanguage = "id-id"
	BlogLanguageIs     BlogLanguage = "is"
	BlogLanguageIsIs   BlogLanguage = "is-is"
	BlogLanguageIt     BlogLanguage = "it"
	BlogLanguageItCh   BlogLanguage = "it-ch"
	BlogLanguageItIt   BlogLanguage = "it-it"
	BlogLanguageItSm   BlogLanguage = "it-sm"
	BlogLanguageItVa   BlogLanguage = "it-va"
	BlogLanguageHeIl   BlogLanguage = "he-il"
	BlogLanguageJa     BlogLanguage = "ja"
	BlogLanguageJaJp   BlogLanguage = "ja-jp"
	BlogLanguageJgo    BlogLanguage = "jgo"
	BlogLanguageJgoCm  BlogLanguage = "jgo-cm"
	BlogLanguageYi     BlogLanguage = "yi"
	BlogLanguageYi001  BlogLanguage = "yi-001"
	BlogLanguageJmc    BlogLanguage = "jmc"
	BlogLanguageJmcTz  BlogLanguage = "jmc-tz"
	BlogLanguageJv     BlogLanguage = "jv"
	BlogLanguageJvID   BlogLanguage = "jv-id"
	BlogLanguageKa     BlogLanguage = "ka"
	BlogLanguageKaGe   BlogLanguage = "ka-ge"
	BlogLanguageKab    BlogLanguage = "kab"
	BlogLanguageKabDz  BlogLanguage = "kab-dz"
	BlogLanguageKam    BlogLanguage = "kam"
	BlogLanguageKamKe  BlogLanguage = "kam-ke"
	BlogLanguageKde    BlogLanguage = "kde"
	BlogLanguageKdeTz  BlogLanguage = "kde-tz"
	BlogLanguageKea    BlogLanguage = "kea"
	BlogLanguageKeaCv  BlogLanguage = "kea-cv"
	BlogLanguageKhq    BlogLanguage = "khq"
	BlogLanguageKhqMl  BlogLanguage = "khq-ml"
	BlogLanguageKi     BlogLanguage = "ki"
	BlogLanguageKiKe   BlogLanguage = "ki-ke"
	BlogLanguageKk     BlogLanguage = "kk"
	BlogLanguageKkKz   BlogLanguage = "kk-kz"
	BlogLanguageKkj    BlogLanguage = "kkj"
	BlogLanguageKkjCm  BlogLanguage = "kkj-cm"
	BlogLanguageKl     BlogLanguage = "kl"
	BlogLanguageKlGl   BlogLanguage = "kl-gl"
	BlogLanguageKln    BlogLanguage = "kln"
	BlogLanguageKlnKe  BlogLanguage = "kln-ke"
	BlogLanguageKm     BlogLanguage = "km"
	BlogLanguageKmKh   BlogLanguage = "km-kh"
	BlogLanguageKn     BlogLanguage = "kn"
	BlogLanguageKnIn   BlogLanguage = "kn-in"
	BlogLanguageKo     BlogLanguage = "ko"
	BlogLanguageKoKp   BlogLanguage = "ko-kp"
	BlogLanguageKoKr   BlogLanguage = "ko-kr"
	BlogLanguageKok    BlogLanguage = "kok"
	BlogLanguageKokIn  BlogLanguage = "kok-in"
	BlogLanguageKs     BlogLanguage = "ks"
	BlogLanguageKsIn   BlogLanguage = "ks-in"
	BlogLanguageKsb    BlogLanguage = "ksb"
	BlogLanguageKsbTz  BlogLanguage = "ksb-tz"
	BlogLanguageKsf    BlogLanguage = "ksf"
	BlogLanguageKsfCm  BlogLanguage = "ksf-cm"
	BlogLanguageKsh    BlogLanguage = "ksh"
	BlogLanguageKshDe  BlogLanguage = "ksh-de"
	BlogLanguageKw     BlogLanguage = "kw"
	BlogLanguageKwGB   BlogLanguage = "kw-gb"
	BlogLanguageKu     BlogLanguage = "ku"
	BlogLanguageKuTr   BlogLanguage = "ku-tr"
	BlogLanguageKy     BlogLanguage = "ky"
	BlogLanguageKyKg   BlogLanguage = "ky-kg"
	BlogLanguageLag    BlogLanguage = "lag"
	BlogLanguageLagTz  BlogLanguage = "lag-tz"
	BlogLanguageLb     BlogLanguage = "lb"
	BlogLanguageLbLu   BlogLanguage = "lb-lu"
	BlogLanguageLg     BlogLanguage = "lg"
	BlogLanguageLgUg   BlogLanguage = "lg-ug"
	BlogLanguageLkt    BlogLanguage = "lkt"
	BlogLanguageLktUs  BlogLanguage = "lkt-us"
	BlogLanguageLn     BlogLanguage = "ln"
	BlogLanguageLnAo   BlogLanguage = "ln-ao"
	BlogLanguageLnCd   BlogLanguage = "ln-cd"
	BlogLanguageLnCf   BlogLanguage = "ln-cf"
	BlogLanguageLnCg   BlogLanguage = "ln-cg"
	BlogLanguageLo     BlogLanguage = "lo"
	BlogLanguageLoLa   BlogLanguage = "lo-la"
	BlogLanguageLrc    BlogLanguage = "lrc"
	BlogLanguageLrcIq  BlogLanguage = "lrc-iq"
	BlogLanguageLrcIr  BlogLanguage = "lrc-ir"
	BlogLanguageLt     BlogLanguage = "lt"
	BlogLanguageLtLt   BlogLanguage = "lt-lt"
	BlogLanguageLu     BlogLanguage = "lu"
	BlogLanguageLuCd   BlogLanguage = "lu-cd"
	BlogLanguageLuo    BlogLanguage = "luo"
	BlogLanguageLuoKe  BlogLanguage = "luo-ke"
	BlogLanguageLuy    BlogLanguage = "luy"
	BlogLanguageLuyKe  BlogLanguage = "luy-ke"
	BlogLanguageLv     BlogLanguage = "lv"
	BlogLanguageLvLv   BlogLanguage = "lv-lv"
	BlogLanguageMai    BlogLanguage = "mai"
	BlogLanguageMaiIn  BlogLanguage = "mai-in"
	BlogLanguageMas    BlogLanguage = "mas"
	BlogLanguageMasKe  BlogLanguage = "mas-ke"
	BlogLanguageMasTz  BlogLanguage = "mas-tz"
	BlogLanguageMer    BlogLanguage = "mer"
	BlogLanguageMerKe  BlogLanguage = "mer-ke"
	BlogLanguageMfe    BlogLanguage = "mfe"
	BlogLanguageMfeMu  BlogLanguage = "mfe-mu"
	BlogLanguageMg     BlogLanguage = "mg"
	BlogLanguageMgMg   BlogLanguage = "mg-mg"
	BlogLanguageMgh    BlogLanguage = "mgh"
	BlogLanguageMghMz  BlogLanguage = "mgh-mz"
	BlogLanguageMgo    BlogLanguage = "mgo"
	BlogLanguageMgoCm  BlogLanguage = "mgo-cm"
	BlogLanguageMi     BlogLanguage = "mi"
	BlogLanguageMiNz   BlogLanguage = "mi-nz"
	BlogLanguageMk     BlogLanguage = "mk"
	BlogLanguageMkMk   BlogLanguage = "mk-mk"
	BlogLanguageMl     BlogLanguage = "ml"
	BlogLanguageMlIn   BlogLanguage = "ml-in"
	BlogLanguageMn     BlogLanguage = "mn"
	BlogLanguageMnMn   BlogLanguage = "mn-mn"
	BlogLanguageMni    BlogLanguage = "mni"
	BlogLanguageMniIn  BlogLanguage = "mni-in"
	BlogLanguageMr     BlogLanguage = "mr"
	BlogLanguageMrIn   BlogLanguage = "mr-in"
	BlogLanguageMs     BlogLanguage = "ms"
	BlogLanguageMsBn   BlogLanguage = "ms-bn"
	BlogLanguageMsID   BlogLanguage = "ms-id"
	BlogLanguageMsMy   BlogLanguage = "ms-my"
	BlogLanguageMsSg   BlogLanguage = "ms-sg"
	BlogLanguageMt     BlogLanguage = "mt"
	BlogLanguageMtMt   BlogLanguage = "mt-mt"
	BlogLanguageMua    BlogLanguage = "mua"
	BlogLanguageMuaCm  BlogLanguage = "mua-cm"
	BlogLanguageMy     BlogLanguage = "my"
	BlogLanguageMyMm   BlogLanguage = "my-mm"
	BlogLanguageMzn    BlogLanguage = "mzn"
	BlogLanguageMznIr  BlogLanguage = "mzn-ir"
	BlogLanguageNaq    BlogLanguage = "naq"
	BlogLanguageNaqNa  BlogLanguage = "naq-na"
	BlogLanguageNb     BlogLanguage = "nb"
	BlogLanguageNbNo   BlogLanguage = "nb-no"
	BlogLanguageNbSj   BlogLanguage = "nb-sj"
	BlogLanguageNd     BlogLanguage = "nd"
	BlogLanguageNdZw   BlogLanguage = "nd-zw"
	BlogLanguageNds    BlogLanguage = "nds"
	BlogLanguageNdsDe  BlogLanguage = "nds-de"
	BlogLanguageNdsNl  BlogLanguage = "nds-nl"
	BlogLanguageNe     BlogLanguage = "ne"
	BlogLanguageNeIn   BlogLanguage = "ne-in"
	BlogLanguageNeNp   BlogLanguage = "ne-np"
	BlogLanguageNl     BlogLanguage = "nl"
	BlogLanguageNlAw   BlogLanguage = "nl-aw"
	BlogLanguageNlBe   BlogLanguage = "nl-be"
	BlogLanguageNlCh   BlogLanguage = "nl-ch"
	BlogLanguageNlBq   BlogLanguage = "nl-bq"
	BlogLanguageNlCw   BlogLanguage = "nl-cw"
	BlogLanguageNlLu   BlogLanguage = "nl-lu"
	BlogLanguageNlNl   BlogLanguage = "nl-nl"
	BlogLanguageNlSr   BlogLanguage = "nl-sr"
	BlogLanguageNlSx   BlogLanguage = "nl-sx"
	BlogLanguageNmg    BlogLanguage = "nmg"
	BlogLanguageNmgCm  BlogLanguage = "nmg-cm"
	BlogLanguageNn     BlogLanguage = "nn"
	BlogLanguageNnNo   BlogLanguage = "nn-no"
	BlogLanguageNnh    BlogLanguage = "nnh"
	BlogLanguageNnhCm  BlogLanguage = "nnh-cm"
	BlogLanguageNo     BlogLanguage = "no"
	BlogLanguageNoNo   BlogLanguage = "no-no"
	BlogLanguageNus    BlogLanguage = "nus"
	BlogLanguageNusSS  BlogLanguage = "nus-ss"
	BlogLanguageNyn    BlogLanguage = "nyn"
	BlogLanguageNynUg  BlogLanguage = "nyn-ug"
	BlogLanguageOm     BlogLanguage = "om"
	BlogLanguageOmEt   BlogLanguage = "om-et"
	BlogLanguageOmKe   BlogLanguage = "om-ke"
	BlogLanguageOr     BlogLanguage = "or"
	BlogLanguageOrIn   BlogLanguage = "or-in"
	BlogLanguageOs     BlogLanguage = "os"
	BlogLanguageOsGe   BlogLanguage = "os-ge"
	BlogLanguageOsRu   BlogLanguage = "os-ru"
	BlogLanguagePa     BlogLanguage = "pa"
	BlogLanguagePaIn   BlogLanguage = "pa-in"
	BlogLanguagePaPk   BlogLanguage = "pa-pk"
	BlogLanguagePcm    BlogLanguage = "pcm"
	BlogLanguagePcmNg  BlogLanguage = "pcm-ng"
	BlogLanguagePl     BlogLanguage = "pl"
	BlogLanguagePlPl   BlogLanguage = "pl-pl"
	BlogLanguagePrg    BlogLanguage = "prg"
	BlogLanguagePrg001 BlogLanguage = "prg-001"
	BlogLanguagePs     BlogLanguage = "ps"
	BlogLanguagePsAf   BlogLanguage = "ps-af"
	BlogLanguagePsPk   BlogLanguage = "ps-pk"
	BlogLanguagePt     BlogLanguage = "pt"
	BlogLanguagePtAo   BlogLanguage = "pt-ao"
	BlogLanguagePtBr   BlogLanguage = "pt-br"
	BlogLanguagePtCh   BlogLanguage = "pt-ch"
	BlogLanguagePtCv   BlogLanguage = "pt-cv"
	BlogLanguagePtGq   BlogLanguage = "pt-gq"
	BlogLanguagePtGw   BlogLanguage = "pt-gw"
	BlogLanguagePtLu   BlogLanguage = "pt-lu"
	BlogLanguagePtMo   BlogLanguage = "pt-mo"
	BlogLanguagePtMz   BlogLanguage = "pt-mz"
	BlogLanguagePtPt   BlogLanguage = "pt-pt"
	BlogLanguagePtSt   BlogLanguage = "pt-st"
	BlogLanguagePtTl   BlogLanguage = "pt-tl"
	BlogLanguageQu     BlogLanguage = "qu"
	BlogLanguageQuBo   BlogLanguage = "qu-bo"
	BlogLanguageQuEc   BlogLanguage = "qu-ec"
	BlogLanguageQuPe   BlogLanguage = "qu-pe"
	BlogLanguageRm     BlogLanguage = "rm"
	BlogLanguageRmCh   BlogLanguage = "rm-ch"
	BlogLanguageRn     BlogLanguage = "rn"
	BlogLanguageRnBi   BlogLanguage = "rn-bi"
	BlogLanguageRo     BlogLanguage = "ro"
	BlogLanguageRoMd   BlogLanguage = "ro-md"
	BlogLanguageRoRo   BlogLanguage = "ro-ro"
	BlogLanguageRof    BlogLanguage = "rof"
	BlogLanguageRofTz  BlogLanguage = "rof-tz"
	BlogLanguageRu     BlogLanguage = "ru"
	BlogLanguageRuBy   BlogLanguage = "ru-by"
	BlogLanguageRuKg   BlogLanguage = "ru-kg"
	BlogLanguageRuKz   BlogLanguage = "ru-kz"
	BlogLanguageRuMd   BlogLanguage = "ru-md"
	BlogLanguageRuRu   BlogLanguage = "ru-ru"
	BlogLanguageRuUa   BlogLanguage = "ru-ua"
	BlogLanguageRw     BlogLanguage = "rw"
	BlogLanguageRwRw   BlogLanguage = "rw-rw"
	BlogLanguageRwk    BlogLanguage = "rwk"
	BlogLanguageRwkTz  BlogLanguage = "rwk-tz"
	BlogLanguageSa     BlogLanguage = "sa"
	BlogLanguageSaIn   BlogLanguage = "sa-in"
	BlogLanguageSah    BlogLanguage = "sah"
	BlogLanguageSahRu  BlogLanguage = "sah-ru"
	BlogLanguageSaq    BlogLanguage = "saq"
	BlogLanguageSaqKe  BlogLanguage = "saq-ke"
	BlogLanguageSat    BlogLanguage = "sat"
	BlogLanguageSatIn  BlogLanguage = "sat-in"
	BlogLanguageSbp    BlogLanguage = "sbp"
	BlogLanguageSbpTz  BlogLanguage = "sbp-tz"
	BlogLanguageSd     BlogLanguage = "sd"
	BlogLanguageSdIn   BlogLanguage = "sd-in"
	BlogLanguageSdPk   BlogLanguage = "sd-pk"
	BlogLanguageSe     BlogLanguage = "se"
	BlogLanguageSeFi   BlogLanguage = "se-fi"
	BlogLanguageSeNo   BlogLanguage = "se-no"
	BlogLanguageSeSe   BlogLanguage = "se-se"
	BlogLanguageSeh    BlogLanguage = "seh"
	BlogLanguageSehMz  BlogLanguage = "seh-mz"
	BlogLanguageSes    BlogLanguage = "ses"
	BlogLanguageSesMl  BlogLanguage = "ses-ml"
	BlogLanguageSg     BlogLanguage = "sg"
	BlogLanguageSgCf   BlogLanguage = "sg-cf"
	BlogLanguageShi    BlogLanguage = "shi"
	BlogLanguageShiMa  BlogLanguage = "shi-ma"
	BlogLanguageSi     BlogLanguage = "si"
	BlogLanguageSiLk   BlogLanguage = "si-lk"
	BlogLanguageSk     BlogLanguage = "sk"
	BlogLanguageSkSk   BlogLanguage = "sk-sk"
	BlogLanguageSl     BlogLanguage = "sl"
	BlogLanguageSlSi   BlogLanguage = "sl-si"
	BlogLanguageSmn    BlogLanguage = "smn"
	BlogLanguageSmnFi  BlogLanguage = "smn-fi"
	BlogLanguageSn     BlogLanguage = "sn"
	BlogLanguageSnZw   BlogLanguage = "sn-zw"
	BlogLanguageSo     BlogLanguage = "so"
	BlogLanguageSoDj   BlogLanguage = "so-dj"
	BlogLanguageSoEt   BlogLanguage = "so-et"
	BlogLanguageSoKe   BlogLanguage = "so-ke"
	BlogLanguageSoSo   BlogLanguage = "so-so"
	BlogLanguageSq     BlogLanguage = "sq"
	BlogLanguageSqAl   BlogLanguage = "sq-al"
	BlogLanguageSqMk   BlogLanguage = "sq-mk"
	BlogLanguageSqXk   BlogLanguage = "sq-xk"
	BlogLanguageSr     BlogLanguage = "sr"
	BlogLanguageSrBa   BlogLanguage = "sr-ba"
	BlogLanguageSrCs   BlogLanguage = "sr-cs"
	BlogLanguageSrMe   BlogLanguage = "sr-me"
	BlogLanguageSrRs   BlogLanguage = "sr-rs"
	BlogLanguageSrXk   BlogLanguage = "sr-xk"
	BlogLanguageSu     BlogLanguage = "su"
	BlogLanguageSuID   BlogLanguage = "su-id"
	BlogLanguageSv     BlogLanguage = "sv"
	BlogLanguageSvAx   BlogLanguage = "sv-ax"
	BlogLanguageSvFi   BlogLanguage = "sv-fi"
	BlogLanguageSvSe   BlogLanguage = "sv-se"
	BlogLanguageSw     BlogLanguage = "sw"
	BlogLanguageSwCd   BlogLanguage = "sw-cd"
	BlogLanguageSwKe   BlogLanguage = "sw-ke"
	BlogLanguageSwTz   BlogLanguage = "sw-tz"
	BlogLanguageSwUg   BlogLanguage = "sw-ug"
	BlogLanguageSy     BlogLanguage = "sy"
	BlogLanguageTa     BlogLanguage = "ta"
	BlogLanguageTaIn   BlogLanguage = "ta-in"
	BlogLanguageTaLk   BlogLanguage = "ta-lk"
	BlogLanguageTaMy   BlogLanguage = "ta-my"
	BlogLanguageTaSg   BlogLanguage = "ta-sg"
	BlogLanguageTe     BlogLanguage = "te"
	BlogLanguageTeIn   BlogLanguage = "te-in"
	BlogLanguageTeo    BlogLanguage = "teo"
	BlogLanguageTeoKe  BlogLanguage = "teo-ke"
	BlogLanguageTeoUg  BlogLanguage = "teo-ug"
	BlogLanguageTg     BlogLanguage = "tg"
	BlogLanguageTgTj   BlogLanguage = "tg-tj"
	BlogLanguageTh     BlogLanguage = "th"
	BlogLanguageThTh   BlogLanguage = "th-th"
	BlogLanguageTi     BlogLanguage = "ti"
	BlogLanguageTiEr   BlogLanguage = "ti-er"
	BlogLanguageTiEt   BlogLanguage = "ti-et"
	BlogLanguageTk     BlogLanguage = "tk"
	BlogLanguageTkTm   BlogLanguage = "tk-tm"
	BlogLanguageTl     BlogLanguage = "tl"
	BlogLanguageTo     BlogLanguage = "to"
	BlogLanguageToTo   BlogLanguage = "to-to"
	BlogLanguageTr     BlogLanguage = "tr"
	BlogLanguageTrCy   BlogLanguage = "tr-cy"
	BlogLanguageTrTr   BlogLanguage = "tr-tr"
	BlogLanguageTt     BlogLanguage = "tt"
	BlogLanguageTtRu   BlogLanguage = "tt-ru"
	BlogLanguageTwq    BlogLanguage = "twq"
	BlogLanguageTwqNe  BlogLanguage = "twq-ne"
	BlogLanguageTzm    BlogLanguage = "tzm"
	BlogLanguageTzmMa  BlogLanguage = "tzm-ma"
	BlogLanguageUg     BlogLanguage = "ug"
	BlogLanguageUgCn   BlogLanguage = "ug-cn"
	BlogLanguageUk     BlogLanguage = "uk"
	BlogLanguageUkUa   BlogLanguage = "uk-ua"
	BlogLanguageUr     BlogLanguage = "ur"
	BlogLanguageUrIn   BlogLanguage = "ur-in"
	BlogLanguageUrPk   BlogLanguage = "ur-pk"
	BlogLanguageUz     BlogLanguage = "uz"
	BlogLanguageUzAf   BlogLanguage = "uz-af"
	BlogLanguageUzUz   BlogLanguage = "uz-uz"
	BlogLanguageVai    BlogLanguage = "vai"
	BlogLanguageVaiLr  BlogLanguage = "vai-lr"
	BlogLanguageVi     BlogLanguage = "vi"
	BlogLanguageViVn   BlogLanguage = "vi-vn"
	BlogLanguageVo     BlogLanguage = "vo"
	BlogLanguageVo001  BlogLanguage = "vo-001"
	BlogLanguageVun    BlogLanguage = "vun"
	BlogLanguageVunTz  BlogLanguage = "vun-tz"
	BlogLanguageWae    BlogLanguage = "wae"
	BlogLanguageWaeCh  BlogLanguage = "wae-ch"
	BlogLanguageWo     BlogLanguage = "wo"
	BlogLanguageWoSn   BlogLanguage = "wo-sn"
	BlogLanguageXh     BlogLanguage = "xh"
	BlogLanguageXhZa   BlogLanguage = "xh-za"
	BlogLanguageXog    BlogLanguage = "xog"
	BlogLanguageXogUg  BlogLanguage = "xog-ug"
	BlogLanguageYav    BlogLanguage = "yav"
	BlogLanguageYavCm  BlogLanguage = "yav-cm"
	BlogLanguageYo     BlogLanguage = "yo"
	BlogLanguageYoBj   BlogLanguage = "yo-bj"
	BlogLanguageYoNg   BlogLanguage = "yo-ng"
	BlogLanguageYue    BlogLanguage = "yue"
	BlogLanguageYueCn  BlogLanguage = "yue-cn"
	BlogLanguageYueHk  BlogLanguage = "yue-hk"
	BlogLanguageZgh    BlogLanguage = "zgh"
	BlogLanguageZghMa  BlogLanguage = "zgh-ma"
	BlogLanguageZh     BlogLanguage = "zh"
	BlogLanguageZhCn   BlogLanguage = "zh-cn"
	BlogLanguageZhHk   BlogLanguage = "zh-hk"
	BlogLanguageZhMo   BlogLanguage = "zh-mo"
	BlogLanguageZhSg   BlogLanguage = "zh-sg"
	BlogLanguageZhTw   BlogLanguage = "zh-tw"
	BlogLanguageZhHans BlogLanguage = "zh-hans"
	BlogLanguageZhHant BlogLanguage = "zh-hant"
	BlogLanguageZu     BlogLanguage = "zu"
	BlogLanguageZuZa   BlogLanguage = "zu-za"
)

// Request body object for creating new language variant blog.
//
// The property ID is required.
type BlogLanguageCloneRequestVNextParam struct {
	// ID of blog to clone.
	ID string `json:"id,required"`
	// Target language of new variant.
	Language param.Opt[string] `json:"language,omitzero"`
	// Language of primary blog to clone.
	PrimaryLanguage param.Opt[string] `json:"primaryLanguage,omitzero"`
	// Path to this blog.
	Slug param.Opt[string] `json:"slug,omitzero"`
	paramObj
}

func (r BlogLanguageCloneRequestVNextParam) MarshalJSON() (data []byte, err error) {
	type shadow BlogLanguageCloneRequestVNextParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BlogLanguageCloneRequestVNextParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Response object for collections of blogs with pagination information.
type CollectionResponseWithTotalBlogForwardPaging struct {
	// Collection of blogs.
	Results []Blog `json:"results,required"`
	// Total number of blogs.
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
func (r CollectionResponseWithTotalBlogForwardPaging) RawJSON() string { return r.JSON.raw }
func (r *CollectionResponseWithTotalBlogForwardPaging) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Response object for collections of blog versions with pagination information.
type CollectionResponseWithTotalVersionBlog struct {
	// Collection of blog versions.
	Results []VersionBlog `json:"results,required"`
	// Total number of blog versions.
	Total int64 `json:"total,required"`
	// Contains information pagination of results.
	Paging marketing.Paging `json:"paging"`
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
func (r CollectionResponseWithTotalVersionBlog) RawJSON() string { return r.JSON.raw }
func (r *CollectionResponseWithTotalVersionBlog) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Model definition for a Version Blog. Contains metadata describing the version of
// the Blog. It can be used to view edit history of the settings.
type VersionBlog struct {
	// The id of the version.
	ID        string    `json:"id,required"`
	Object    Blog      `json:"object,required"`
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
func (r VersionBlog) RawJSON() string { return r.JSON.raw }
func (r *VersionBlog) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BlogSettingListParams struct {
	// The cursor token value to get the next set of results. You can get this from the
	// `paging.next.after` JSON property of a paged response containing more results.
	After param.Opt[string] `query:"after,omitzero" json:"-"`
	// Specifies whether to return archived Blogs. Defaults to `false`.
	Archived param.Opt[bool] `query:"archived,omitzero" json:"-"`
	// Only return Blogs created after the specified time.
	CreatedAfter param.Opt[time.Time] `query:"createdAfter,omitzero" format:"date-time" json:"-"`
	// Only return Blogs created at exactly the specified time.
	CreatedAt param.Opt[time.Time] `query:"createdAt,omitzero" format:"date-time" json:"-"`
	// Only return Blogs created before the specified time.
	CreatedBefore param.Opt[time.Time] `query:"createdBefore,omitzero" format:"date-time" json:"-"`
	// The maximum number of results to return. Default is 100.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Only return Blogs last updated after the specified time.
	UpdatedAfter param.Opt[time.Time] `query:"updatedAfter,omitzero" format:"date-time" json:"-"`
	// Only return Blogs last updated at exactly the specified time.
	UpdatedAt param.Opt[time.Time] `query:"updatedAt,omitzero" format:"date-time" json:"-"`
	// Only return Blogs last updated before the specified time.
	UpdatedBefore param.Opt[time.Time] `query:"updatedBefore,omitzero" format:"date-time" json:"-"`
	// Specifies which fields to use for sorting results. Valid fields are `name` and
	// `id`
	Sort []string `query:"sort,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [BlogSettingListParams]'s query parameters as `url.Values`.
func (r BlogSettingListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type BlogSettingAttachToLangGroupParams struct {
	// Request body object for attaching objects to multi-language groups.
	AttachToLangPrimaryRequestVNext AttachToLangPrimaryRequestVNextParam
	paramObj
}

func (r BlogSettingAttachToLangGroupParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.AttachToLangPrimaryRequestVNext)
}
func (r *BlogSettingAttachToLangGroupParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.AttachToLangPrimaryRequestVNext)
}

type BlogSettingNewLanguageVariationParams struct {
	// Request body object for creating new language variant blog.
	BlogLanguageCloneRequestVNext BlogLanguageCloneRequestVNextParam
	paramObj
}

func (r BlogSettingNewLanguageVariationParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.BlogLanguageCloneRequestVNext)
}
func (r *BlogSettingNewLanguageVariationParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.BlogLanguageCloneRequestVNext)
}

type BlogSettingDetachFromLangGroupParams struct {
	// Request body object for detaching objects from multi-language groups.
	DetachFromLangGroupRequestVNext DetachFromLangGroupRequestVNextParam
	paramObj
}

func (r BlogSettingDetachFromLangGroupParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.DetachFromLangGroupRequestVNext)
}
func (r *BlogSettingDetachFromLangGroupParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.DetachFromLangGroupRequestVNext)
}

type BlogSettingGetRevisionParams struct {
	BlogID string `path:"blogId,required" json:"-"`
	paramObj
}

type BlogSettingListRevisionsParams struct {
	// The cursor token value to get the next set of results. You can get this from the
	// `paging.next.after` JSON property of a paged response containing more results.
	After  param.Opt[string] `query:"after,omitzero" json:"-"`
	Before param.Opt[string] `query:"before,omitzero" json:"-"`
	// The maximum number of results to return. Default is 100.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [BlogSettingListRevisionsParams]'s query parameters as
// `url.Values`.
func (r BlogSettingListRevisionsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type BlogSettingSetNewLangPrimaryParams struct {
	// Request body object for setting a new primary language.
	SetNewLanguagePrimaryRequestVNext SetNewLanguagePrimaryRequestVNextParam
	paramObj
}

func (r BlogSettingSetNewLangPrimaryParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.SetNewLanguagePrimaryRequestVNext)
}
func (r *BlogSettingSetNewLangPrimaryParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.SetNewLanguagePrimaryRequestVNext)
}

type BlogSettingUpdateLanguagesParams struct {
	// Request object for updating languages within a multi-language group.
	UpdateLanguagesRequestVNext UpdateLanguagesRequestVNextParam
	paramObj
}

func (r BlogSettingUpdateLanguagesParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.UpdateLanguagesRequestVNext)
}
func (r *BlogSettingUpdateLanguagesParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.UpdateLanguagesRequestVNext)
}
