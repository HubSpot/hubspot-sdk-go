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
	"github.com/HubSpot/hubspot-sdk-go/internal/requestconfig"
	"github.com/HubSpot/hubspot-sdk-go/option"
	"github.com/HubSpot/hubspot-sdk-go/packages/pagination"
	"github.com/HubSpot/hubspot-sdk-go/packages/param"
	"github.com/HubSpot/hubspot-sdk-go/packages/respjson"
	"github.com/HubSpot/hubspot-sdk-go/shared"
)

// BlogSettingService contains methods and other services that help with
// interacting with the hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBlogSettingService] method instead.
type BlogSettingService struct {
	options       []option.RequestOption
	MultiLanguage BlogSettingMultiLanguageService
}

// NewBlogSettingService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewBlogSettingService(opts ...option.RequestOption) (r BlogSettingService) {
	r = BlogSettingService{}
	r.options = opts
	r.MultiLanguage = NewBlogSettingMultiLanguageService(opts...)
	return
}

// Get the list of blogs. Results can be limited and filtered by creation or
// updated date.
func (r *BlogSettingService) List(ctx context.Context, query BlogSettingListParams, opts ...option.RequestOption) (res *pagination.Page[Blog], err error) {
	var raw *http.Response
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "cms/blog-settings/2026-03/settings"
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

// Get the list of blogs. Results can be limited and filtered by creation or
// updated date.
func (r *BlogSettingService) ListAutoPaging(ctx context.Context, query BlogSettingListParams, opts ...option.RequestOption) *pagination.PageAutoPager[Blog] {
	return pagination.NewPageAutoPager(r.List(ctx, query, opts...))
}

// Retrieve a specific blog by its ID.
func (r *BlogSettingService) Get(ctx context.Context, blogID string, opts ...option.RequestOption) (res *Blog, err error) {
	opts = slices.Concat(r.options, opts)
	if blogID == "" {
		err = errors.New("missing required blogId parameter")
		return nil, err
	}
	path := fmt.Sprintf("cms/blog-settings/2026-03/settings/%s", url.PathEscape(blogID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Get a specific blog revision.
func (r *BlogSettingService) GetRevision(ctx context.Context, revisionID string, query BlogSettingGetRevisionParams, opts ...option.RequestOption) (res *BlogVersion, err error) {
	opts = slices.Concat(r.options, opts)
	if query.BlogID == "" {
		err = errors.New("missing required blogId parameter")
		return nil, err
	}
	if revisionID == "" {
		err = errors.New("missing required revisionId parameter")
		return nil, err
	}
	path := fmt.Sprintf("cms/blog-settings/2026-03/settings/%s/revisions/%s", url.PathEscape(query.BlogID), url.PathEscape(revisionID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Get the list of blog revisions. Results can be limited and filtered by creation
// or updated date.
func (r *BlogSettingService) ListRevisions(ctx context.Context, blogID string, query BlogSettingListRevisionsParams, opts ...option.RequestOption) (res *pagination.Page[VersionBlog], err error) {
	var raw *http.Response
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if blogID == "" {
		err = errors.New("missing required blogId parameter")
		return nil, err
	}
	path := fmt.Sprintf("cms/blog-settings/2026-03/settings/%s/revisions", url.PathEscape(blogID))
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

// Get the list of blog revisions. Results can be limited and filtered by creation
// or updated date.
func (r *BlogSettingService) ListRevisionsAutoPaging(ctx context.Context, blogID string, query BlogSettingListRevisionsParams, opts ...option.RequestOption) *pagination.PageAutoPager[VersionBlog] {
	return pagination.NewPageAutoPager(r.ListRevisions(ctx, blogID, query, opts...))
}

type Blog struct {
	// The unique ID of the Blog.
	ID string `json:"id" api:"required"`
	// Blog's root URL
	AbsoluteURL string `json:"absoluteUrl" api:"required"`
	// Boolean determining whether or not this blog allows public comments.
	AllowComments bool `json:"allowComments" api:"required"`
	// The timestamp (ISO8601 format) when this blog was created.
	Created time.Time `json:"created" api:"required" format:"date-time"`
	// The timestamp (ISO8601 format) when this Blog was deleted.
	DeletedAt time.Time `json:"deletedAt" api:"required" format:"date-time"`
	// The Description of this Blog.
	Description string `json:"description" api:"required"`
	// The html title of this Blog.
	HTMLTitle string `json:"htmlTitle" api:"required"`
	// The explicitly defined language of the Blog. If null, the Blog will default to
	// the language of the Domain.
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
	Language      BlogLanguage `json:"language" api:"required"`
	ListingPageID string       `json:"listingPageId" api:"required"`
	// The internal name of the blog.
	Name string `json:"name" api:"required"`
	// Rules for require member registration to access private content.
	PublicAccessRules []PublicAccessRule `json:"publicAccessRules" api:"required"`
	// Boolean to determine whether or not to respect publicAccessRules.
	PublicAccessRulesEnabled bool `json:"publicAccessRulesEnabled" api:"required"`
	// The public title of this Blog.
	PublicTitle string `json:"publicTitle" api:"required"`
	// The path of the this blog. This field is appended to the domain to construct the
	// url of this blog.
	Slug string `json:"slug" api:"required"`
	// ID of the primary Blog this object was translated from.
	TranslatedFromID string `json:"translatedFromId" api:"required"`
	// The timestamp (ISO8601 format) when this blog was updated.
	Updated time.Time `json:"updated" api:"required" format:"date-time"`
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
		ListingPageID            respjson.Field
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
	BlogLanguageAa     BlogLanguage = "aa"
	BlogLanguageAb     BlogLanguage = "ab"
	BlogLanguageAe     BlogLanguage = "ae"
	BlogLanguageAf     BlogLanguage = "af"
	BlogLanguageAfNa   BlogLanguage = "af-na"
	BlogLanguageAfZa   BlogLanguage = "af-za"
	BlogLanguageAgq    BlogLanguage = "agq"
	BlogLanguageAgqCm  BlogLanguage = "agq-cm"
	BlogLanguageAk     BlogLanguage = "ak"
	BlogLanguageAkGh   BlogLanguage = "ak-gh"
	BlogLanguageAm     BlogLanguage = "am"
	BlogLanguageAmEt   BlogLanguage = "am-et"
	BlogLanguageAn     BlogLanguage = "an"
	BlogLanguageAnn    BlogLanguage = "ann"
	BlogLanguageAnnNg  BlogLanguage = "ann-ng"
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
	BlogLanguageAv     BlogLanguage = "av"
	BlogLanguageAy     BlogLanguage = "ay"
	BlogLanguageAz     BlogLanguage = "az"
	BlogLanguageAzAz   BlogLanguage = "az-az"
	BlogLanguageBa     BlogLanguage = "ba"
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
	BlogLanguageBgc    BlogLanguage = "bgc"
	BlogLanguageBgcIn  BlogLanguage = "bgc-in"
	BlogLanguageBho    BlogLanguage = "bho"
	BlogLanguageBhoIn  BlogLanguage = "bho-in"
	BlogLanguageBi     BlogLanguage = "bi"
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
	BlogLanguageCh     BlogLanguage = "ch"
	BlogLanguageChr    BlogLanguage = "chr"
	BlogLanguageChrUs  BlogLanguage = "chr-us"
	BlogLanguageCkb    BlogLanguage = "ckb"
	BlogLanguageCkbIq  BlogLanguage = "ckb-iq"
	BlogLanguageCkbIr  BlogLanguage = "ckb-ir"
	BlogLanguageCo     BlogLanguage = "co"
	BlogLanguageCr     BlogLanguage = "cr"
	BlogLanguageCs     BlogLanguage = "cs"
	BlogLanguageCsCz   BlogLanguage = "cs-cz"
	BlogLanguageCu     BlogLanguage = "cu"
	BlogLanguageCuRu   BlogLanguage = "cu-ru"
	BlogLanguageCv     BlogLanguage = "cv"
	BlogLanguageCvRu   BlogLanguage = "cv-ru"
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
	BlogLanguageDv     BlogLanguage = "dv"
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
	BlogLanguageEnEe   BlogLanguage = "en-ee"
	BlogLanguageEnEg   BlogLanguage = "en-eg"
	BlogLanguageEnEr   BlogLanguage = "en-er"
	BlogLanguageEnEs   BlogLanguage = "en-es"
	BlogLanguageEnFi   BlogLanguage = "en-fi"
	BlogLanguageEnFj   BlogLanguage = "en-fj"
	BlogLanguageEnFk   BlogLanguage = "en-fk"
	BlogLanguageEnFm   BlogLanguage = "en-fm"
	BlogLanguageEnFr   BlogLanguage = "en-fr"
	BlogLanguageEnGB   BlogLanguage = "en-gb"
	BlogLanguageEnGd   BlogLanguage = "en-gd"
	BlogLanguageEnGg   BlogLanguage = "en-gg"
	BlogLanguageEnGh   BlogLanguage = "en-gh"
	BlogLanguageEnGi   BlogLanguage = "en-gi"
	BlogLanguageEnGm   BlogLanguage = "en-gm"
	BlogLanguageEnGu   BlogLanguage = "en-gu"
	BlogLanguageEnGy   BlogLanguage = "en-gy"
	BlogLanguageEnHk   BlogLanguage = "en-hk"
	BlogLanguageEnID   BlogLanguage = "en-id"
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
	BlogLanguageEnMv   BlogLanguage = "en-mv"
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
	BlogLanguageEnPt   BlogLanguage = "en-pt"
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
	BlogLanguageEnTh   BlogLanguage = "en-th"
	BlogLanguageEnTk   BlogLanguage = "en-tk"
	BlogLanguageEnTn   BlogLanguage = "en-tn"
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
	BlogLanguageEnVn   BlogLanguage = "en-vn"
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
	BlogLanguageFj     BlogLanguage = "fj"
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
	BlogLanguageFrr    BlogLanguage = "frr"
	BlogLanguageFrrDe  BlogLanguage = "frr-de"
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
	BlogLanguageGn     BlogLanguage = "gn"
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
	BlogLanguageHeIl   BlogLanguage = "he-il"
	BlogLanguageHi     BlogLanguage = "hi"
	BlogLanguageHiIn   BlogLanguage = "hi-in"
	BlogLanguageHmn    BlogLanguage = "hmn"
	BlogLanguageHo     BlogLanguage = "ho"
	BlogLanguageHr     BlogLanguage = "hr"
	BlogLanguageHrBa   BlogLanguage = "hr-ba"
	BlogLanguageHrHr   BlogLanguage = "hr-hr"
	BlogLanguageHsb    BlogLanguage = "hsb"
	BlogLanguageHsbDe  BlogLanguage = "hsb-de"
	BlogLanguageHt     BlogLanguage = "ht"
	BlogLanguageHu     BlogLanguage = "hu"
	BlogLanguageHuHu   BlogLanguage = "hu-hu"
	BlogLanguageHy     BlogLanguage = "hy"
	BlogLanguageHyAm   BlogLanguage = "hy-am"
	BlogLanguageHz     BlogLanguage = "hz"
	BlogLanguageIa     BlogLanguage = "ia"
	BlogLanguageIa001  BlogLanguage = "ia-001"
	BlogLanguageID     BlogLanguage = "id"
	BlogLanguageIDID   BlogLanguage = "id-id"
	BlogLanguageIe     BlogLanguage = "ie"
	BlogLanguageIg     BlogLanguage = "ig"
	BlogLanguageIgNg   BlogLanguage = "ig-ng"
	BlogLanguageIi     BlogLanguage = "ii"
	BlogLanguageIiCn   BlogLanguage = "ii-cn"
	BlogLanguageIk     BlogLanguage = "ik"
	BlogLanguageIo     BlogLanguage = "io"
	BlogLanguageIs     BlogLanguage = "is"
	BlogLanguageIsIs   BlogLanguage = "is-is"
	BlogLanguageIt     BlogLanguage = "it"
	BlogLanguageItCh   BlogLanguage = "it-ch"
	BlogLanguageItIt   BlogLanguage = "it-it"
	BlogLanguageItSm   BlogLanguage = "it-sm"
	BlogLanguageItVa   BlogLanguage = "it-va"
	BlogLanguageIu     BlogLanguage = "iu"
	BlogLanguageJa     BlogLanguage = "ja"
	BlogLanguageJaJp   BlogLanguage = "ja-jp"
	BlogLanguageJgo    BlogLanguage = "jgo"
	BlogLanguageJgoCm  BlogLanguage = "jgo-cm"
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
	BlogLanguageKar    BlogLanguage = "kar"
	BlogLanguageKde    BlogLanguage = "kde"
	BlogLanguageKdeTz  BlogLanguage = "kde-tz"
	BlogLanguageKea    BlogLanguage = "kea"
	BlogLanguageKeaCv  BlogLanguage = "kea-cv"
	BlogLanguageKg     BlogLanguage = "kg"
	BlogLanguageKgp    BlogLanguage = "kgp"
	BlogLanguageKgpBr  BlogLanguage = "kgp-br"
	BlogLanguageKh     BlogLanguage = "kh"
	BlogLanguageKhq    BlogLanguage = "khq"
	BlogLanguageKhqMl  BlogLanguage = "khq-ml"
	BlogLanguageKi     BlogLanguage = "ki"
	BlogLanguageKiKe   BlogLanguage = "ki-ke"
	BlogLanguageKj     BlogLanguage = "kj"
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
	BlogLanguageKr     BlogLanguage = "kr"
	BlogLanguageKs     BlogLanguage = "ks"
	BlogLanguageKsIn   BlogLanguage = "ks-in"
	BlogLanguageKsb    BlogLanguage = "ksb"
	BlogLanguageKsbTz  BlogLanguage = "ksb-tz"
	BlogLanguageKsf    BlogLanguage = "ksf"
	BlogLanguageKsfCm  BlogLanguage = "ksf-cm"
	BlogLanguageKsh    BlogLanguage = "ksh"
	BlogLanguageKshDe  BlogLanguage = "ksh-de"
	BlogLanguageKu     BlogLanguage = "ku"
	BlogLanguageKuTr   BlogLanguage = "ku-tr"
	BlogLanguageKv     BlogLanguage = "kv"
	BlogLanguageKw     BlogLanguage = "kw"
	BlogLanguageKwGB   BlogLanguage = "kw-gb"
	BlogLanguageKy     BlogLanguage = "ky"
	BlogLanguageKyKg   BlogLanguage = "ky-kg"
	BlogLanguageLa     BlogLanguage = "la"
	BlogLanguageLag    BlogLanguage = "lag"
	BlogLanguageLagTz  BlogLanguage = "lag-tz"
	BlogLanguageLb     BlogLanguage = "lb"
	BlogLanguageLbLu   BlogLanguage = "lb-lu"
	BlogLanguageLg     BlogLanguage = "lg"
	BlogLanguageLgUg   BlogLanguage = "lg-ug"
	BlogLanguageLi     BlogLanguage = "li"
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
	BlogLanguageMdf    BlogLanguage = "mdf"
	BlogLanguageMdfRu  BlogLanguage = "mdf-ru"
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
	BlogLanguageMh     BlogLanguage = "mh"
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
	BlogLanguageNa     BlogLanguage = "na"
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
	BlogLanguageNg     BlogLanguage = "ng"
	BlogLanguageNl     BlogLanguage = "nl"
	BlogLanguageNlAw   BlogLanguage = "nl-aw"
	BlogLanguageNlBe   BlogLanguage = "nl-be"
	BlogLanguageNlBq   BlogLanguage = "nl-bq"
	BlogLanguageNlCh   BlogLanguage = "nl-ch"
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
	BlogLanguageNr     BlogLanguage = "nr"
	BlogLanguageNus    BlogLanguage = "nus"
	BlogLanguageNusSS  BlogLanguage = "nus-ss"
	BlogLanguageNv     BlogLanguage = "nv"
	BlogLanguageNy     BlogLanguage = "ny"
	BlogLanguageNyn    BlogLanguage = "nyn"
	BlogLanguageNynUg  BlogLanguage = "nyn-ug"
	BlogLanguageOc     BlogLanguage = "oc"
	BlogLanguageOcEs   BlogLanguage = "oc-es"
	BlogLanguageOcFr   BlogLanguage = "oc-fr"
	BlogLanguageOj     BlogLanguage = "oj"
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
	BlogLanguagePi     BlogLanguage = "pi"
	BlogLanguagePis    BlogLanguage = "pis"
	BlogLanguagePisSb  BlogLanguage = "pis-sb"
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
	BlogLanguageRaj    BlogLanguage = "raj"
	BlogLanguageRajIn  BlogLanguage = "raj-in"
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
	BlogLanguageSc     BlogLanguage = "sc"
	BlogLanguageScIt   BlogLanguage = "sc-it"
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
	BlogLanguageSm     BlogLanguage = "sm"
	BlogLanguageSmn    BlogLanguage = "smn"
	BlogLanguageSmnFi  BlogLanguage = "smn-fi"
	BlogLanguageSMS    BlogLanguage = "sms"
	BlogLanguageSMSFi  BlogLanguage = "sms-fi"
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
	BlogLanguageSS     BlogLanguage = "ss"
	BlogLanguageSt     BlogLanguage = "st"
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
	BlogLanguageTn     BlogLanguage = "tn"
	BlogLanguageTo     BlogLanguage = "to"
	BlogLanguageToTo   BlogLanguage = "to-to"
	BlogLanguageTok    BlogLanguage = "tok"
	BlogLanguageTok001 BlogLanguage = "tok-001"
	BlogLanguageTr     BlogLanguage = "tr"
	BlogLanguageTrCy   BlogLanguage = "tr-cy"
	BlogLanguageTrTr   BlogLanguage = "tr-tr"
	BlogLanguageTs     BlogLanguage = "ts"
	BlogLanguageTt     BlogLanguage = "tt"
	BlogLanguageTtRu   BlogLanguage = "tt-ru"
	BlogLanguageTw     BlogLanguage = "tw"
	BlogLanguageTwq    BlogLanguage = "twq"
	BlogLanguageTwqNe  BlogLanguage = "twq-ne"
	BlogLanguageTy     BlogLanguage = "ty"
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
	BlogLanguageVe     BlogLanguage = "ve"
	BlogLanguageVi     BlogLanguage = "vi"
	BlogLanguageViVn   BlogLanguage = "vi-vn"
	BlogLanguageVo     BlogLanguage = "vo"
	BlogLanguageVo001  BlogLanguage = "vo-001"
	BlogLanguageVun    BlogLanguage = "vun"
	BlogLanguageVunTz  BlogLanguage = "vun-tz"
	BlogLanguageWa     BlogLanguage = "wa"
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
	BlogLanguageYi     BlogLanguage = "yi"
	BlogLanguageYi001  BlogLanguage = "yi-001"
	BlogLanguageYo     BlogLanguage = "yo"
	BlogLanguageYoBj   BlogLanguage = "yo-bj"
	BlogLanguageYoNg   BlogLanguage = "yo-ng"
	BlogLanguageYrl    BlogLanguage = "yrl"
	BlogLanguageYrlBr  BlogLanguage = "yrl-br"
	BlogLanguageYrlCo  BlogLanguage = "yrl-co"
	BlogLanguageYrlVe  BlogLanguage = "yrl-ve"
	BlogLanguageYue    BlogLanguage = "yue"
	BlogLanguageYueCn  BlogLanguage = "yue-cn"
	BlogLanguageYueHk  BlogLanguage = "yue-hk"
	BlogLanguageZa     BlogLanguage = "za"
	BlogLanguageZgh    BlogLanguage = "zgh"
	BlogLanguageZghMa  BlogLanguage = "zgh-ma"
	BlogLanguageZh     BlogLanguage = "zh"
	BlogLanguageZhCn   BlogLanguage = "zh-cn"
	BlogLanguageZhHans BlogLanguage = "zh-hans"
	BlogLanguageZhHant BlogLanguage = "zh-hant"
	BlogLanguageZhHk   BlogLanguage = "zh-hk"
	BlogLanguageZhMo   BlogLanguage = "zh-mo"
	BlogLanguageZhSg   BlogLanguage = "zh-sg"
	BlogLanguageZhTw   BlogLanguage = "zh-tw"
	BlogLanguageZu     BlogLanguage = "zu"
	BlogLanguageZuZa   BlogLanguage = "zu-za"
)

// The property ID is required.
type BlogLanguageCloneRequestVNextParam struct {
	// ID of blog to clone.
	ID string `json:"id" api:"required"`
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

type BlogVersion struct {
	ID        string             `json:"id" api:"required"`
	Object    Blog               `json:"object" api:"required"`
	UpdatedAt time.Time          `json:"updatedAt" api:"required" format:"date-time"`
	User      shared.VersionUser `json:"user" api:"required"`
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
func (r BlogVersion) RawJSON() string { return r.JSON.raw }
func (r *BlogVersion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CollectionResponseWithTotalBlog struct {
	Results []Blog        `json:"results" api:"required"`
	Total   int64         `json:"total" api:"required"`
	Paging  shared.Paging `json:"paging"`
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
func (r CollectionResponseWithTotalBlog) RawJSON() string { return r.JSON.raw }
func (r *CollectionResponseWithTotalBlog) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CollectionResponseWithTotalBlogVersion struct {
	Results []VersionBlog `json:"results" api:"required"`
	Total   int64         `json:"total" api:"required"`
	Paging  shared.Paging `json:"paging"`
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
func (r CollectionResponseWithTotalBlogVersion) RawJSON() string { return r.JSON.raw }
func (r *CollectionResponseWithTotalBlogVersion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VersionBlog struct {
	// The id of the version.
	ID     string `json:"id" api:"required"`
	Object Blog   `json:"object" api:"required"`
	// The timestamp (ISO8601 format) when this blog version was updated.
	UpdatedAt time.Time          `json:"updatedAt" api:"required" format:"date-time"`
	User      shared.VersionUser `json:"user" api:"required"`
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
	UpdatedAfter  param.Opt[time.Time] `query:"updatedAfter,omitzero" format:"date-time" json:"-"`
	UpdatedAt     param.Opt[time.Time] `query:"updatedAt,omitzero" format:"date-time" json:"-"`
	UpdatedBefore param.Opt[time.Time] `query:"updatedBefore,omitzero" format:"date-time" json:"-"`
	Sort          []string             `query:"sort,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [BlogSettingListParams]'s query parameters as `url.Values`.
func (r BlogSettingListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type BlogSettingGetRevisionParams struct {
	BlogID string `path:"blogId" api:"required" json:"-"`
	paramObj
}

type BlogSettingListRevisionsParams struct {
	// The paging cursor token of the last successfully read resource will be returned
	// as the `paging.next.after` JSON property of a paged response containing more
	// results.
	After  param.Opt[string] `query:"after,omitzero" json:"-"`
	Before param.Opt[string] `query:"before,omitzero" json:"-"`
	// The maximum number of results to display per page.
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
