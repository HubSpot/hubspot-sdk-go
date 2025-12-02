// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cms

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/stainless-sdks/hubspot-sdk-go/internal/apijson"
	"github.com/stainless-sdks/hubspot-sdk-go/internal/apiquery"
	"github.com/stainless-sdks/hubspot-sdk-go/internal/requestconfig"
	"github.com/stainless-sdks/hubspot-sdk-go/option"
	"github.com/stainless-sdks/hubspot-sdk-go/packages/param"
	"github.com/stainless-sdks/hubspot-sdk-go/packages/respjson"
)

// SiteSearchService contains methods and other services that help with interacting
// with the hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSiteSearchService] method instead.
type SiteSearchService struct {
	Options []option.RequestOption
}

// NewSiteSearchService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewSiteSearchService(opts ...option.RequestOption) (r SiteSearchService) {
	r = SiteSearchService{}
	r.Options = opts
	return
}

// For a given account and document ID (page ID, blog post ID, HubDB row ID, etc.),
// return all indexed data for that document. This is useful when debugging why a
// particular document is not returned from a custom search.
func (r *SiteSearchService) GetIndexedData(ctx context.Context, contentID string, query SiteSearchGetIndexedDataParams, opts ...option.RequestOption) (res *IndexedData, err error) {
	opts = slices.Concat(r.Options, opts)
	if contentID == "" {
		err = errors.New("missing required contentId parameter")
		return
	}
	path := fmt.Sprintf("cms/v3/site-search/indexed-data/%s", contentID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Returns any website content matching the given search criteria for a given
// HubSpot account. Searches can be filtered by content type, domain, or URL path.
func (r *SiteSearchService) Search(ctx context.Context, query SiteSearchSearchParams, opts ...option.RequestOption) (res *PublicSearchResults, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "cms/v3/site-search/search"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// An individual search result.
type ContentSearchResult struct {
	// The ID of the content.
	ID int64 `json:"id,required"`
	// The domain the document is hosted on.
	Domain string `json:"domain,required"`
	// The matching score of the document.
	Score float64 `json:"score,required"`
	// The type of document. Can be `SITE_PAGE`, `LANDING_PAGE`, `BLOG_POST`,
	// `LISTING_PAGE`, or `KNOWLEDGE_ARTICLE`.
	//
	// Any of "BLOG_POST", "KNOWLEDGE_ARTICLE", "LANDING_PAGE", "LISTING_PAGE",
	// "SITE_PAGE".
	Type ContentSearchResultType `json:"type,required"`
	// The url of the document.
	URL string `json:"url,required"`
	// Name of the author.
	AuthorFullName string `json:"authorFullName"`
	// For knowledge articles, the category of the article.
	Category string `json:"category"`
	// The ID of the document in HubSpot.
	CombinedID string `json:"combinedId"`
	// The result's description. The content will be determined by the value of
	// `length` in the request.
	Description string `json:"description"`
	// URL of the featured image.
	FeaturedImageURL string `json:"featuredImageUrl"`
	// The document's language.
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
	Language ContentSearchResultLanguage `json:"language"`
	// The date the content was published.
	PublishedDate int64 `json:"publishedDate"`
	// If a dynamic page, the row ID in the HubDB table.
	RowID int64 `json:"rowId"`
	// For knowledge articles, the subcategory of the article.
	Subcategory string `json:"subcategory"`
	// If a dynamic page, the ID of the HubDB table.
	TableID int64 `json:"tableId"`
	// If a blog post, the tags associated with it.
	Tags []string `json:"tags"`
	// The title of the returned document.
	Title string `json:"title"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		Domain           respjson.Field
		Score            respjson.Field
		Type             respjson.Field
		URL              respjson.Field
		AuthorFullName   respjson.Field
		Category         respjson.Field
		CombinedID       respjson.Field
		Description      respjson.Field
		FeaturedImageURL respjson.Field
		Language         respjson.Field
		PublishedDate    respjson.Field
		RowID            respjson.Field
		Subcategory      respjson.Field
		TableID          respjson.Field
		Tags             respjson.Field
		Title            respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ContentSearchResult) RawJSON() string { return r.JSON.raw }
func (r *ContentSearchResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The type of document. Can be `SITE_PAGE`, `LANDING_PAGE`, `BLOG_POST`,
// `LISTING_PAGE`, or `KNOWLEDGE_ARTICLE`.
type ContentSearchResultType string

const (
	ContentSearchResultTypeBlogPost         ContentSearchResultType = "BLOG_POST"
	ContentSearchResultTypeKnowledgeArticle ContentSearchResultType = "KNOWLEDGE_ARTICLE"
	ContentSearchResultTypeLandingPage      ContentSearchResultType = "LANDING_PAGE"
	ContentSearchResultTypeListingPage      ContentSearchResultType = "LISTING_PAGE"
	ContentSearchResultTypeSitePage         ContentSearchResultType = "SITE_PAGE"
)

// The document's language.
type ContentSearchResultLanguage string

const (
	ContentSearchResultLanguageAf     ContentSearchResultLanguage = "af"
	ContentSearchResultLanguageAfNa   ContentSearchResultLanguage = "af-na"
	ContentSearchResultLanguageAfZa   ContentSearchResultLanguage = "af-za"
	ContentSearchResultLanguageAgq    ContentSearchResultLanguage = "agq"
	ContentSearchResultLanguageAgqCm  ContentSearchResultLanguage = "agq-cm"
	ContentSearchResultLanguageAk     ContentSearchResultLanguage = "ak"
	ContentSearchResultLanguageAkGh   ContentSearchResultLanguage = "ak-gh"
	ContentSearchResultLanguageAm     ContentSearchResultLanguage = "am"
	ContentSearchResultLanguageAmEt   ContentSearchResultLanguage = "am-et"
	ContentSearchResultLanguageAr     ContentSearchResultLanguage = "ar"
	ContentSearchResultLanguageAr001  ContentSearchResultLanguage = "ar-001"
	ContentSearchResultLanguageArAe   ContentSearchResultLanguage = "ar-ae"
	ContentSearchResultLanguageArBh   ContentSearchResultLanguage = "ar-bh"
	ContentSearchResultLanguageArDj   ContentSearchResultLanguage = "ar-dj"
	ContentSearchResultLanguageArDz   ContentSearchResultLanguage = "ar-dz"
	ContentSearchResultLanguageArEg   ContentSearchResultLanguage = "ar-eg"
	ContentSearchResultLanguageArEh   ContentSearchResultLanguage = "ar-eh"
	ContentSearchResultLanguageArEr   ContentSearchResultLanguage = "ar-er"
	ContentSearchResultLanguageArIl   ContentSearchResultLanguage = "ar-il"
	ContentSearchResultLanguageArIq   ContentSearchResultLanguage = "ar-iq"
	ContentSearchResultLanguageArJo   ContentSearchResultLanguage = "ar-jo"
	ContentSearchResultLanguageArKm   ContentSearchResultLanguage = "ar-km"
	ContentSearchResultLanguageArKw   ContentSearchResultLanguage = "ar-kw"
	ContentSearchResultLanguageArLb   ContentSearchResultLanguage = "ar-lb"
	ContentSearchResultLanguageArLy   ContentSearchResultLanguage = "ar-ly"
	ContentSearchResultLanguageArMa   ContentSearchResultLanguage = "ar-ma"
	ContentSearchResultLanguageArMr   ContentSearchResultLanguage = "ar-mr"
	ContentSearchResultLanguageArOm   ContentSearchResultLanguage = "ar-om"
	ContentSearchResultLanguageArPs   ContentSearchResultLanguage = "ar-ps"
	ContentSearchResultLanguageArQa   ContentSearchResultLanguage = "ar-qa"
	ContentSearchResultLanguageArSa   ContentSearchResultLanguage = "ar-sa"
	ContentSearchResultLanguageArSd   ContentSearchResultLanguage = "ar-sd"
	ContentSearchResultLanguageArSo   ContentSearchResultLanguage = "ar-so"
	ContentSearchResultLanguageArSS   ContentSearchResultLanguage = "ar-ss"
	ContentSearchResultLanguageArSy   ContentSearchResultLanguage = "ar-sy"
	ContentSearchResultLanguageArTd   ContentSearchResultLanguage = "ar-td"
	ContentSearchResultLanguageArTn   ContentSearchResultLanguage = "ar-tn"
	ContentSearchResultLanguageArYe   ContentSearchResultLanguage = "ar-ye"
	ContentSearchResultLanguageAs     ContentSearchResultLanguage = "as"
	ContentSearchResultLanguageAsIn   ContentSearchResultLanguage = "as-in"
	ContentSearchResultLanguageAsa    ContentSearchResultLanguage = "asa"
	ContentSearchResultLanguageAsaTz  ContentSearchResultLanguage = "asa-tz"
	ContentSearchResultLanguageAst    ContentSearchResultLanguage = "ast"
	ContentSearchResultLanguageAstEs  ContentSearchResultLanguage = "ast-es"
	ContentSearchResultLanguageAz     ContentSearchResultLanguage = "az"
	ContentSearchResultLanguageAzAz   ContentSearchResultLanguage = "az-az"
	ContentSearchResultLanguageBas    ContentSearchResultLanguage = "bas"
	ContentSearchResultLanguageBasCm  ContentSearchResultLanguage = "bas-cm"
	ContentSearchResultLanguageBe     ContentSearchResultLanguage = "be"
	ContentSearchResultLanguageBeBy   ContentSearchResultLanguage = "be-by"
	ContentSearchResultLanguageBem    ContentSearchResultLanguage = "bem"
	ContentSearchResultLanguageBemZm  ContentSearchResultLanguage = "bem-zm"
	ContentSearchResultLanguageBez    ContentSearchResultLanguage = "bez"
	ContentSearchResultLanguageBezTz  ContentSearchResultLanguage = "bez-tz"
	ContentSearchResultLanguageBg     ContentSearchResultLanguage = "bg"
	ContentSearchResultLanguageBgBg   ContentSearchResultLanguage = "bg-bg"
	ContentSearchResultLanguageBm     ContentSearchResultLanguage = "bm"
	ContentSearchResultLanguageBmMl   ContentSearchResultLanguage = "bm-ml"
	ContentSearchResultLanguageBn     ContentSearchResultLanguage = "bn"
	ContentSearchResultLanguageBnBd   ContentSearchResultLanguage = "bn-bd"
	ContentSearchResultLanguageBnIn   ContentSearchResultLanguage = "bn-in"
	ContentSearchResultLanguageBo     ContentSearchResultLanguage = "bo"
	ContentSearchResultLanguageBoCn   ContentSearchResultLanguage = "bo-cn"
	ContentSearchResultLanguageBoIn   ContentSearchResultLanguage = "bo-in"
	ContentSearchResultLanguageBr     ContentSearchResultLanguage = "br"
	ContentSearchResultLanguageBrFr   ContentSearchResultLanguage = "br-fr"
	ContentSearchResultLanguageBrx    ContentSearchResultLanguage = "brx"
	ContentSearchResultLanguageBrxIn  ContentSearchResultLanguage = "brx-in"
	ContentSearchResultLanguageBs     ContentSearchResultLanguage = "bs"
	ContentSearchResultLanguageBsBa   ContentSearchResultLanguage = "bs-ba"
	ContentSearchResultLanguageCa     ContentSearchResultLanguage = "ca"
	ContentSearchResultLanguageCaAd   ContentSearchResultLanguage = "ca-ad"
	ContentSearchResultLanguageCaEs   ContentSearchResultLanguage = "ca-es"
	ContentSearchResultLanguageCaFr   ContentSearchResultLanguage = "ca-fr"
	ContentSearchResultLanguageCaIt   ContentSearchResultLanguage = "ca-it"
	ContentSearchResultLanguageCcp    ContentSearchResultLanguage = "ccp"
	ContentSearchResultLanguageCcpBd  ContentSearchResultLanguage = "ccp-bd"
	ContentSearchResultLanguageCcpIn  ContentSearchResultLanguage = "ccp-in"
	ContentSearchResultLanguageCe     ContentSearchResultLanguage = "ce"
	ContentSearchResultLanguageCeRu   ContentSearchResultLanguage = "ce-ru"
	ContentSearchResultLanguageCeb    ContentSearchResultLanguage = "ceb"
	ContentSearchResultLanguageCebPh  ContentSearchResultLanguage = "ceb-ph"
	ContentSearchResultLanguageCgg    ContentSearchResultLanguage = "cgg"
	ContentSearchResultLanguageCggUg  ContentSearchResultLanguage = "cgg-ug"
	ContentSearchResultLanguageChr    ContentSearchResultLanguage = "chr"
	ContentSearchResultLanguageChrUs  ContentSearchResultLanguage = "chr-us"
	ContentSearchResultLanguageCkb    ContentSearchResultLanguage = "ckb"
	ContentSearchResultLanguageCkbIq  ContentSearchResultLanguage = "ckb-iq"
	ContentSearchResultLanguageCkbIr  ContentSearchResultLanguage = "ckb-ir"
	ContentSearchResultLanguageCs     ContentSearchResultLanguage = "cs"
	ContentSearchResultLanguageCsCz   ContentSearchResultLanguage = "cs-cz"
	ContentSearchResultLanguageCu     ContentSearchResultLanguage = "cu"
	ContentSearchResultLanguageCuRu   ContentSearchResultLanguage = "cu-ru"
	ContentSearchResultLanguageCy     ContentSearchResultLanguage = "cy"
	ContentSearchResultLanguageCyGB   ContentSearchResultLanguage = "cy-gb"
	ContentSearchResultLanguageDa     ContentSearchResultLanguage = "da"
	ContentSearchResultLanguageDaDk   ContentSearchResultLanguage = "da-dk"
	ContentSearchResultLanguageDaGl   ContentSearchResultLanguage = "da-gl"
	ContentSearchResultLanguageDav    ContentSearchResultLanguage = "dav"
	ContentSearchResultLanguageDavKe  ContentSearchResultLanguage = "dav-ke"
	ContentSearchResultLanguageDe     ContentSearchResultLanguage = "de"
	ContentSearchResultLanguageDeAt   ContentSearchResultLanguage = "de-at"
	ContentSearchResultLanguageDeBe   ContentSearchResultLanguage = "de-be"
	ContentSearchResultLanguageDeCh   ContentSearchResultLanguage = "de-ch"
	ContentSearchResultLanguageDeDe   ContentSearchResultLanguage = "de-de"
	ContentSearchResultLanguageDeGr   ContentSearchResultLanguage = "de-gr"
	ContentSearchResultLanguageDeIt   ContentSearchResultLanguage = "de-it"
	ContentSearchResultLanguageDeLi   ContentSearchResultLanguage = "de-li"
	ContentSearchResultLanguageDeLu   ContentSearchResultLanguage = "de-lu"
	ContentSearchResultLanguageDje    ContentSearchResultLanguage = "dje"
	ContentSearchResultLanguageDjeNe  ContentSearchResultLanguage = "dje-ne"
	ContentSearchResultLanguageDoi    ContentSearchResultLanguage = "doi"
	ContentSearchResultLanguageDoiIn  ContentSearchResultLanguage = "doi-in"
	ContentSearchResultLanguageDsb    ContentSearchResultLanguage = "dsb"
	ContentSearchResultLanguageDsbDe  ContentSearchResultLanguage = "dsb-de"
	ContentSearchResultLanguageDua    ContentSearchResultLanguage = "dua"
	ContentSearchResultLanguageDuaCm  ContentSearchResultLanguage = "dua-cm"
	ContentSearchResultLanguageDyo    ContentSearchResultLanguage = "dyo"
	ContentSearchResultLanguageDyoSn  ContentSearchResultLanguage = "dyo-sn"
	ContentSearchResultLanguageDz     ContentSearchResultLanguage = "dz"
	ContentSearchResultLanguageDzBt   ContentSearchResultLanguage = "dz-bt"
	ContentSearchResultLanguageEbu    ContentSearchResultLanguage = "ebu"
	ContentSearchResultLanguageEbuKe  ContentSearchResultLanguage = "ebu-ke"
	ContentSearchResultLanguageEe     ContentSearchResultLanguage = "ee"
	ContentSearchResultLanguageEeGh   ContentSearchResultLanguage = "ee-gh"
	ContentSearchResultLanguageEeTg   ContentSearchResultLanguage = "ee-tg"
	ContentSearchResultLanguageEl     ContentSearchResultLanguage = "el"
	ContentSearchResultLanguageElCy   ContentSearchResultLanguage = "el-cy"
	ContentSearchResultLanguageElGr   ContentSearchResultLanguage = "el-gr"
	ContentSearchResultLanguageEn     ContentSearchResultLanguage = "en"
	ContentSearchResultLanguageEn001  ContentSearchResultLanguage = "en-001"
	ContentSearchResultLanguageEn150  ContentSearchResultLanguage = "en-150"
	ContentSearchResultLanguageEnAe   ContentSearchResultLanguage = "en-ae"
	ContentSearchResultLanguageEnAg   ContentSearchResultLanguage = "en-ag"
	ContentSearchResultLanguageEnAI   ContentSearchResultLanguage = "en-ai"
	ContentSearchResultLanguageEnAs   ContentSearchResultLanguage = "en-as"
	ContentSearchResultLanguageEnAt   ContentSearchResultLanguage = "en-at"
	ContentSearchResultLanguageEnAu   ContentSearchResultLanguage = "en-au"
	ContentSearchResultLanguageEnBb   ContentSearchResultLanguage = "en-bb"
	ContentSearchResultLanguageEnBe   ContentSearchResultLanguage = "en-be"
	ContentSearchResultLanguageEnBi   ContentSearchResultLanguage = "en-bi"
	ContentSearchResultLanguageEnBm   ContentSearchResultLanguage = "en-bm"
	ContentSearchResultLanguageEnBs   ContentSearchResultLanguage = "en-bs"
	ContentSearchResultLanguageEnBw   ContentSearchResultLanguage = "en-bw"
	ContentSearchResultLanguageEnBz   ContentSearchResultLanguage = "en-bz"
	ContentSearchResultLanguageEnCa   ContentSearchResultLanguage = "en-ca"
	ContentSearchResultLanguageEnCc   ContentSearchResultLanguage = "en-cc"
	ContentSearchResultLanguageEnCh   ContentSearchResultLanguage = "en-ch"
	ContentSearchResultLanguageEnCk   ContentSearchResultLanguage = "en-ck"
	ContentSearchResultLanguageEnCm   ContentSearchResultLanguage = "en-cm"
	ContentSearchResultLanguageEnCn   ContentSearchResultLanguage = "en-cn"
	ContentSearchResultLanguageEnCx   ContentSearchResultLanguage = "en-cx"
	ContentSearchResultLanguageEnCy   ContentSearchResultLanguage = "en-cy"
	ContentSearchResultLanguageEnDe   ContentSearchResultLanguage = "en-de"
	ContentSearchResultLanguageEnDg   ContentSearchResultLanguage = "en-dg"
	ContentSearchResultLanguageEnDk   ContentSearchResultLanguage = "en-dk"
	ContentSearchResultLanguageEnDm   ContentSearchResultLanguage = "en-dm"
	ContentSearchResultLanguageEnEr   ContentSearchResultLanguage = "en-er"
	ContentSearchResultLanguageEnFi   ContentSearchResultLanguage = "en-fi"
	ContentSearchResultLanguageEnFj   ContentSearchResultLanguage = "en-fj"
	ContentSearchResultLanguageEnFk   ContentSearchResultLanguage = "en-fk"
	ContentSearchResultLanguageEnFm   ContentSearchResultLanguage = "en-fm"
	ContentSearchResultLanguageEnGB   ContentSearchResultLanguage = "en-gb"
	ContentSearchResultLanguageEnGd   ContentSearchResultLanguage = "en-gd"
	ContentSearchResultLanguageEnGg   ContentSearchResultLanguage = "en-gg"
	ContentSearchResultLanguageEnGh   ContentSearchResultLanguage = "en-gh"
	ContentSearchResultLanguageEnGi   ContentSearchResultLanguage = "en-gi"
	ContentSearchResultLanguageEnGm   ContentSearchResultLanguage = "en-gm"
	ContentSearchResultLanguageEnGu   ContentSearchResultLanguage = "en-gu"
	ContentSearchResultLanguageEnGy   ContentSearchResultLanguage = "en-gy"
	ContentSearchResultLanguageEnHk   ContentSearchResultLanguage = "en-hk"
	ContentSearchResultLanguageEnIe   ContentSearchResultLanguage = "en-ie"
	ContentSearchResultLanguageEnIl   ContentSearchResultLanguage = "en-il"
	ContentSearchResultLanguageEnIm   ContentSearchResultLanguage = "en-im"
	ContentSearchResultLanguageEnIn   ContentSearchResultLanguage = "en-in"
	ContentSearchResultLanguageEnIo   ContentSearchResultLanguage = "en-io"
	ContentSearchResultLanguageEnJe   ContentSearchResultLanguage = "en-je"
	ContentSearchResultLanguageEnJm   ContentSearchResultLanguage = "en-jm"
	ContentSearchResultLanguageEnKe   ContentSearchResultLanguage = "en-ke"
	ContentSearchResultLanguageEnKi   ContentSearchResultLanguage = "en-ki"
	ContentSearchResultLanguageEnKn   ContentSearchResultLanguage = "en-kn"
	ContentSearchResultLanguageEnKy   ContentSearchResultLanguage = "en-ky"
	ContentSearchResultLanguageEnLc   ContentSearchResultLanguage = "en-lc"
	ContentSearchResultLanguageEnLr   ContentSearchResultLanguage = "en-lr"
	ContentSearchResultLanguageEnLs   ContentSearchResultLanguage = "en-ls"
	ContentSearchResultLanguageEnLu   ContentSearchResultLanguage = "en-lu"
	ContentSearchResultLanguageEnMg   ContentSearchResultLanguage = "en-mg"
	ContentSearchResultLanguageEnMh   ContentSearchResultLanguage = "en-mh"
	ContentSearchResultLanguageEnMo   ContentSearchResultLanguage = "en-mo"
	ContentSearchResultLanguageEnMp   ContentSearchResultLanguage = "en-mp"
	ContentSearchResultLanguageEnMs   ContentSearchResultLanguage = "en-ms"
	ContentSearchResultLanguageEnMt   ContentSearchResultLanguage = "en-mt"
	ContentSearchResultLanguageEnMu   ContentSearchResultLanguage = "en-mu"
	ContentSearchResultLanguageEnMw   ContentSearchResultLanguage = "en-mw"
	ContentSearchResultLanguageEnMx   ContentSearchResultLanguage = "en-mx"
	ContentSearchResultLanguageEnMy   ContentSearchResultLanguage = "en-my"
	ContentSearchResultLanguageEnNa   ContentSearchResultLanguage = "en-na"
	ContentSearchResultLanguageEnNf   ContentSearchResultLanguage = "en-nf"
	ContentSearchResultLanguageEnNg   ContentSearchResultLanguage = "en-ng"
	ContentSearchResultLanguageEnNl   ContentSearchResultLanguage = "en-nl"
	ContentSearchResultLanguageEnNr   ContentSearchResultLanguage = "en-nr"
	ContentSearchResultLanguageEnNu   ContentSearchResultLanguage = "en-nu"
	ContentSearchResultLanguageEnNz   ContentSearchResultLanguage = "en-nz"
	ContentSearchResultLanguageEnPg   ContentSearchResultLanguage = "en-pg"
	ContentSearchResultLanguageEnPh   ContentSearchResultLanguage = "en-ph"
	ContentSearchResultLanguageEnPk   ContentSearchResultLanguage = "en-pk"
	ContentSearchResultLanguageEnPn   ContentSearchResultLanguage = "en-pn"
	ContentSearchResultLanguageEnPr   ContentSearchResultLanguage = "en-pr"
	ContentSearchResultLanguageEnPw   ContentSearchResultLanguage = "en-pw"
	ContentSearchResultLanguageEnRw   ContentSearchResultLanguage = "en-rw"
	ContentSearchResultLanguageEnSb   ContentSearchResultLanguage = "en-sb"
	ContentSearchResultLanguageEnSc   ContentSearchResultLanguage = "en-sc"
	ContentSearchResultLanguageEnSd   ContentSearchResultLanguage = "en-sd"
	ContentSearchResultLanguageEnSe   ContentSearchResultLanguage = "en-se"
	ContentSearchResultLanguageEnSg   ContentSearchResultLanguage = "en-sg"
	ContentSearchResultLanguageEnSh   ContentSearchResultLanguage = "en-sh"
	ContentSearchResultLanguageEnSi   ContentSearchResultLanguage = "en-si"
	ContentSearchResultLanguageEnSl   ContentSearchResultLanguage = "en-sl"
	ContentSearchResultLanguageEnSS   ContentSearchResultLanguage = "en-ss"
	ContentSearchResultLanguageEnSx   ContentSearchResultLanguage = "en-sx"
	ContentSearchResultLanguageEnSz   ContentSearchResultLanguage = "en-sz"
	ContentSearchResultLanguageEnTc   ContentSearchResultLanguage = "en-tc"
	ContentSearchResultLanguageEnTk   ContentSearchResultLanguage = "en-tk"
	ContentSearchResultLanguageEnTo   ContentSearchResultLanguage = "en-to"
	ContentSearchResultLanguageEnTt   ContentSearchResultLanguage = "en-tt"
	ContentSearchResultLanguageEnTv   ContentSearchResultLanguage = "en-tv"
	ContentSearchResultLanguageEnTz   ContentSearchResultLanguage = "en-tz"
	ContentSearchResultLanguageEnUg   ContentSearchResultLanguage = "en-ug"
	ContentSearchResultLanguageEnUm   ContentSearchResultLanguage = "en-um"
	ContentSearchResultLanguageEnUs   ContentSearchResultLanguage = "en-us"
	ContentSearchResultLanguageEnVc   ContentSearchResultLanguage = "en-vc"
	ContentSearchResultLanguageEnVg   ContentSearchResultLanguage = "en-vg"
	ContentSearchResultLanguageEnVi   ContentSearchResultLanguage = "en-vi"
	ContentSearchResultLanguageEnVu   ContentSearchResultLanguage = "en-vu"
	ContentSearchResultLanguageEnWs   ContentSearchResultLanguage = "en-ws"
	ContentSearchResultLanguageEnZa   ContentSearchResultLanguage = "en-za"
	ContentSearchResultLanguageEnZm   ContentSearchResultLanguage = "en-zm"
	ContentSearchResultLanguageEnZw   ContentSearchResultLanguage = "en-zw"
	ContentSearchResultLanguageEo     ContentSearchResultLanguage = "eo"
	ContentSearchResultLanguageEo001  ContentSearchResultLanguage = "eo-001"
	ContentSearchResultLanguageEs     ContentSearchResultLanguage = "es"
	ContentSearchResultLanguageEs419  ContentSearchResultLanguage = "es-419"
	ContentSearchResultLanguageEsAr   ContentSearchResultLanguage = "es-ar"
	ContentSearchResultLanguageEsBo   ContentSearchResultLanguage = "es-bo"
	ContentSearchResultLanguageEsBr   ContentSearchResultLanguage = "es-br"
	ContentSearchResultLanguageEsBz   ContentSearchResultLanguage = "es-bz"
	ContentSearchResultLanguageEsCl   ContentSearchResultLanguage = "es-cl"
	ContentSearchResultLanguageEsCo   ContentSearchResultLanguage = "es-co"
	ContentSearchResultLanguageEsCr   ContentSearchResultLanguage = "es-cr"
	ContentSearchResultLanguageEsCu   ContentSearchResultLanguage = "es-cu"
	ContentSearchResultLanguageEsDo   ContentSearchResultLanguage = "es-do"
	ContentSearchResultLanguageEsEa   ContentSearchResultLanguage = "es-ea"
	ContentSearchResultLanguageEsEc   ContentSearchResultLanguage = "es-ec"
	ContentSearchResultLanguageEsEs   ContentSearchResultLanguage = "es-es"
	ContentSearchResultLanguageEsGq   ContentSearchResultLanguage = "es-gq"
	ContentSearchResultLanguageEsGt   ContentSearchResultLanguage = "es-gt"
	ContentSearchResultLanguageEsHn   ContentSearchResultLanguage = "es-hn"
	ContentSearchResultLanguageEsIc   ContentSearchResultLanguage = "es-ic"
	ContentSearchResultLanguageEsMx   ContentSearchResultLanguage = "es-mx"
	ContentSearchResultLanguageEsNi   ContentSearchResultLanguage = "es-ni"
	ContentSearchResultLanguageEsPa   ContentSearchResultLanguage = "es-pa"
	ContentSearchResultLanguageEsPe   ContentSearchResultLanguage = "es-pe"
	ContentSearchResultLanguageEsPh   ContentSearchResultLanguage = "es-ph"
	ContentSearchResultLanguageEsPr   ContentSearchResultLanguage = "es-pr"
	ContentSearchResultLanguageEsPy   ContentSearchResultLanguage = "es-py"
	ContentSearchResultLanguageEsSv   ContentSearchResultLanguage = "es-sv"
	ContentSearchResultLanguageEsUs   ContentSearchResultLanguage = "es-us"
	ContentSearchResultLanguageEsUy   ContentSearchResultLanguage = "es-uy"
	ContentSearchResultLanguageEsVe   ContentSearchResultLanguage = "es-ve"
	ContentSearchResultLanguageEt     ContentSearchResultLanguage = "et"
	ContentSearchResultLanguageEtEe   ContentSearchResultLanguage = "et-ee"
	ContentSearchResultLanguageEu     ContentSearchResultLanguage = "eu"
	ContentSearchResultLanguageEuEs   ContentSearchResultLanguage = "eu-es"
	ContentSearchResultLanguageEwo    ContentSearchResultLanguage = "ewo"
	ContentSearchResultLanguageEwoCm  ContentSearchResultLanguage = "ewo-cm"
	ContentSearchResultLanguageFa     ContentSearchResultLanguage = "fa"
	ContentSearchResultLanguageFaAf   ContentSearchResultLanguage = "fa-af"
	ContentSearchResultLanguageFaIr   ContentSearchResultLanguage = "fa-ir"
	ContentSearchResultLanguageFf     ContentSearchResultLanguage = "ff"
	ContentSearchResultLanguageFfBf   ContentSearchResultLanguage = "ff-bf"
	ContentSearchResultLanguageFfCm   ContentSearchResultLanguage = "ff-cm"
	ContentSearchResultLanguageFfGh   ContentSearchResultLanguage = "ff-gh"
	ContentSearchResultLanguageFfGm   ContentSearchResultLanguage = "ff-gm"
	ContentSearchResultLanguageFfGn   ContentSearchResultLanguage = "ff-gn"
	ContentSearchResultLanguageFfGw   ContentSearchResultLanguage = "ff-gw"
	ContentSearchResultLanguageFfLr   ContentSearchResultLanguage = "ff-lr"
	ContentSearchResultLanguageFfMr   ContentSearchResultLanguage = "ff-mr"
	ContentSearchResultLanguageFfNe   ContentSearchResultLanguage = "ff-ne"
	ContentSearchResultLanguageFfNg   ContentSearchResultLanguage = "ff-ng"
	ContentSearchResultLanguageFfSl   ContentSearchResultLanguage = "ff-sl"
	ContentSearchResultLanguageFfSn   ContentSearchResultLanguage = "ff-sn"
	ContentSearchResultLanguageFi     ContentSearchResultLanguage = "fi"
	ContentSearchResultLanguageFiFi   ContentSearchResultLanguage = "fi-fi"
	ContentSearchResultLanguageFil    ContentSearchResultLanguage = "fil"
	ContentSearchResultLanguageFilPh  ContentSearchResultLanguage = "fil-ph"
	ContentSearchResultLanguageFo     ContentSearchResultLanguage = "fo"
	ContentSearchResultLanguageFoDk   ContentSearchResultLanguage = "fo-dk"
	ContentSearchResultLanguageFoFo   ContentSearchResultLanguage = "fo-fo"
	ContentSearchResultLanguageFr     ContentSearchResultLanguage = "fr"
	ContentSearchResultLanguageFrBe   ContentSearchResultLanguage = "fr-be"
	ContentSearchResultLanguageFrBf   ContentSearchResultLanguage = "fr-bf"
	ContentSearchResultLanguageFrBi   ContentSearchResultLanguage = "fr-bi"
	ContentSearchResultLanguageFrBj   ContentSearchResultLanguage = "fr-bj"
	ContentSearchResultLanguageFrBl   ContentSearchResultLanguage = "fr-bl"
	ContentSearchResultLanguageFrCa   ContentSearchResultLanguage = "fr-ca"
	ContentSearchResultLanguageFrCd   ContentSearchResultLanguage = "fr-cd"
	ContentSearchResultLanguageFrCf   ContentSearchResultLanguage = "fr-cf"
	ContentSearchResultLanguageFrCg   ContentSearchResultLanguage = "fr-cg"
	ContentSearchResultLanguageFrCh   ContentSearchResultLanguage = "fr-ch"
	ContentSearchResultLanguageFrCi   ContentSearchResultLanguage = "fr-ci"
	ContentSearchResultLanguageFrCm   ContentSearchResultLanguage = "fr-cm"
	ContentSearchResultLanguageFrDj   ContentSearchResultLanguage = "fr-dj"
	ContentSearchResultLanguageFrDz   ContentSearchResultLanguage = "fr-dz"
	ContentSearchResultLanguageFrFr   ContentSearchResultLanguage = "fr-fr"
	ContentSearchResultLanguageFrGa   ContentSearchResultLanguage = "fr-ga"
	ContentSearchResultLanguageFrGf   ContentSearchResultLanguage = "fr-gf"
	ContentSearchResultLanguageFrGn   ContentSearchResultLanguage = "fr-gn"
	ContentSearchResultLanguageFrGp   ContentSearchResultLanguage = "fr-gp"
	ContentSearchResultLanguageFrGq   ContentSearchResultLanguage = "fr-gq"
	ContentSearchResultLanguageFrHt   ContentSearchResultLanguage = "fr-ht"
	ContentSearchResultLanguageFrKm   ContentSearchResultLanguage = "fr-km"
	ContentSearchResultLanguageFrLu   ContentSearchResultLanguage = "fr-lu"
	ContentSearchResultLanguageFrMa   ContentSearchResultLanguage = "fr-ma"
	ContentSearchResultLanguageFrMc   ContentSearchResultLanguage = "fr-mc"
	ContentSearchResultLanguageFrMf   ContentSearchResultLanguage = "fr-mf"
	ContentSearchResultLanguageFrMg   ContentSearchResultLanguage = "fr-mg"
	ContentSearchResultLanguageFrMl   ContentSearchResultLanguage = "fr-ml"
	ContentSearchResultLanguageFrMq   ContentSearchResultLanguage = "fr-mq"
	ContentSearchResultLanguageFrMr   ContentSearchResultLanguage = "fr-mr"
	ContentSearchResultLanguageFrMu   ContentSearchResultLanguage = "fr-mu"
	ContentSearchResultLanguageFrNc   ContentSearchResultLanguage = "fr-nc"
	ContentSearchResultLanguageFrNe   ContentSearchResultLanguage = "fr-ne"
	ContentSearchResultLanguageFrPf   ContentSearchResultLanguage = "fr-pf"
	ContentSearchResultLanguageFrPm   ContentSearchResultLanguage = "fr-pm"
	ContentSearchResultLanguageFrRe   ContentSearchResultLanguage = "fr-re"
	ContentSearchResultLanguageFrRw   ContentSearchResultLanguage = "fr-rw"
	ContentSearchResultLanguageFrSc   ContentSearchResultLanguage = "fr-sc"
	ContentSearchResultLanguageFrSn   ContentSearchResultLanguage = "fr-sn"
	ContentSearchResultLanguageFrSy   ContentSearchResultLanguage = "fr-sy"
	ContentSearchResultLanguageFrTd   ContentSearchResultLanguage = "fr-td"
	ContentSearchResultLanguageFrTg   ContentSearchResultLanguage = "fr-tg"
	ContentSearchResultLanguageFrTn   ContentSearchResultLanguage = "fr-tn"
	ContentSearchResultLanguageFrVu   ContentSearchResultLanguage = "fr-vu"
	ContentSearchResultLanguageFrWf   ContentSearchResultLanguage = "fr-wf"
	ContentSearchResultLanguageFrYt   ContentSearchResultLanguage = "fr-yt"
	ContentSearchResultLanguageFur    ContentSearchResultLanguage = "fur"
	ContentSearchResultLanguageFurIt  ContentSearchResultLanguage = "fur-it"
	ContentSearchResultLanguageFy     ContentSearchResultLanguage = "fy"
	ContentSearchResultLanguageFyNl   ContentSearchResultLanguage = "fy-nl"
	ContentSearchResultLanguageGa     ContentSearchResultLanguage = "ga"
	ContentSearchResultLanguageGaGB   ContentSearchResultLanguage = "ga-gb"
	ContentSearchResultLanguageGaIe   ContentSearchResultLanguage = "ga-ie"
	ContentSearchResultLanguageGd     ContentSearchResultLanguage = "gd"
	ContentSearchResultLanguageGdGB   ContentSearchResultLanguage = "gd-gb"
	ContentSearchResultLanguageGl     ContentSearchResultLanguage = "gl"
	ContentSearchResultLanguageGlEs   ContentSearchResultLanguage = "gl-es"
	ContentSearchResultLanguageGsw    ContentSearchResultLanguage = "gsw"
	ContentSearchResultLanguageGswCh  ContentSearchResultLanguage = "gsw-ch"
	ContentSearchResultLanguageGswFr  ContentSearchResultLanguage = "gsw-fr"
	ContentSearchResultLanguageGswLi  ContentSearchResultLanguage = "gsw-li"
	ContentSearchResultLanguageGu     ContentSearchResultLanguage = "gu"
	ContentSearchResultLanguageGuIn   ContentSearchResultLanguage = "gu-in"
	ContentSearchResultLanguageGuz    ContentSearchResultLanguage = "guz"
	ContentSearchResultLanguageGuzKe  ContentSearchResultLanguage = "guz-ke"
	ContentSearchResultLanguageGv     ContentSearchResultLanguage = "gv"
	ContentSearchResultLanguageGvIm   ContentSearchResultLanguage = "gv-im"
	ContentSearchResultLanguageHa     ContentSearchResultLanguage = "ha"
	ContentSearchResultLanguageHaGh   ContentSearchResultLanguage = "ha-gh"
	ContentSearchResultLanguageHaNe   ContentSearchResultLanguage = "ha-ne"
	ContentSearchResultLanguageHaNg   ContentSearchResultLanguage = "ha-ng"
	ContentSearchResultLanguageHaw    ContentSearchResultLanguage = "haw"
	ContentSearchResultLanguageHawUs  ContentSearchResultLanguage = "haw-us"
	ContentSearchResultLanguageHe     ContentSearchResultLanguage = "he"
	ContentSearchResultLanguageHeIl   ContentSearchResultLanguage = "he-il"
	ContentSearchResultLanguageHi     ContentSearchResultLanguage = "hi"
	ContentSearchResultLanguageHiIn   ContentSearchResultLanguage = "hi-in"
	ContentSearchResultLanguageHr     ContentSearchResultLanguage = "hr"
	ContentSearchResultLanguageHrBa   ContentSearchResultLanguage = "hr-ba"
	ContentSearchResultLanguageHrHr   ContentSearchResultLanguage = "hr-hr"
	ContentSearchResultLanguageHsb    ContentSearchResultLanguage = "hsb"
	ContentSearchResultLanguageHsbDe  ContentSearchResultLanguage = "hsb-de"
	ContentSearchResultLanguageHu     ContentSearchResultLanguage = "hu"
	ContentSearchResultLanguageHuHu   ContentSearchResultLanguage = "hu-hu"
	ContentSearchResultLanguageHy     ContentSearchResultLanguage = "hy"
	ContentSearchResultLanguageHyAm   ContentSearchResultLanguage = "hy-am"
	ContentSearchResultLanguageIa     ContentSearchResultLanguage = "ia"
	ContentSearchResultLanguageIa001  ContentSearchResultLanguage = "ia-001"
	ContentSearchResultLanguageID     ContentSearchResultLanguage = "id"
	ContentSearchResultLanguageIDID   ContentSearchResultLanguage = "id-id"
	ContentSearchResultLanguageIg     ContentSearchResultLanguage = "ig"
	ContentSearchResultLanguageIgNg   ContentSearchResultLanguage = "ig-ng"
	ContentSearchResultLanguageIi     ContentSearchResultLanguage = "ii"
	ContentSearchResultLanguageIiCn   ContentSearchResultLanguage = "ii-cn"
	ContentSearchResultLanguageIs     ContentSearchResultLanguage = "is"
	ContentSearchResultLanguageIsIs   ContentSearchResultLanguage = "is-is"
	ContentSearchResultLanguageIt     ContentSearchResultLanguage = "it"
	ContentSearchResultLanguageItCh   ContentSearchResultLanguage = "it-ch"
	ContentSearchResultLanguageItIt   ContentSearchResultLanguage = "it-it"
	ContentSearchResultLanguageItSm   ContentSearchResultLanguage = "it-sm"
	ContentSearchResultLanguageItVa   ContentSearchResultLanguage = "it-va"
	ContentSearchResultLanguageJa     ContentSearchResultLanguage = "ja"
	ContentSearchResultLanguageJaJp   ContentSearchResultLanguage = "ja-jp"
	ContentSearchResultLanguageJgo    ContentSearchResultLanguage = "jgo"
	ContentSearchResultLanguageJgoCm  ContentSearchResultLanguage = "jgo-cm"
	ContentSearchResultLanguageJmc    ContentSearchResultLanguage = "jmc"
	ContentSearchResultLanguageJmcTz  ContentSearchResultLanguage = "jmc-tz"
	ContentSearchResultLanguageJv     ContentSearchResultLanguage = "jv"
	ContentSearchResultLanguageJvID   ContentSearchResultLanguage = "jv-id"
	ContentSearchResultLanguageKa     ContentSearchResultLanguage = "ka"
	ContentSearchResultLanguageKaGe   ContentSearchResultLanguage = "ka-ge"
	ContentSearchResultLanguageKab    ContentSearchResultLanguage = "kab"
	ContentSearchResultLanguageKabDz  ContentSearchResultLanguage = "kab-dz"
	ContentSearchResultLanguageKam    ContentSearchResultLanguage = "kam"
	ContentSearchResultLanguageKamKe  ContentSearchResultLanguage = "kam-ke"
	ContentSearchResultLanguageKde    ContentSearchResultLanguage = "kde"
	ContentSearchResultLanguageKdeTz  ContentSearchResultLanguage = "kde-tz"
	ContentSearchResultLanguageKea    ContentSearchResultLanguage = "kea"
	ContentSearchResultLanguageKeaCv  ContentSearchResultLanguage = "kea-cv"
	ContentSearchResultLanguageKhq    ContentSearchResultLanguage = "khq"
	ContentSearchResultLanguageKhqMl  ContentSearchResultLanguage = "khq-ml"
	ContentSearchResultLanguageKi     ContentSearchResultLanguage = "ki"
	ContentSearchResultLanguageKiKe   ContentSearchResultLanguage = "ki-ke"
	ContentSearchResultLanguageKk     ContentSearchResultLanguage = "kk"
	ContentSearchResultLanguageKkKz   ContentSearchResultLanguage = "kk-kz"
	ContentSearchResultLanguageKkj    ContentSearchResultLanguage = "kkj"
	ContentSearchResultLanguageKkjCm  ContentSearchResultLanguage = "kkj-cm"
	ContentSearchResultLanguageKl     ContentSearchResultLanguage = "kl"
	ContentSearchResultLanguageKlGl   ContentSearchResultLanguage = "kl-gl"
	ContentSearchResultLanguageKln    ContentSearchResultLanguage = "kln"
	ContentSearchResultLanguageKlnKe  ContentSearchResultLanguage = "kln-ke"
	ContentSearchResultLanguageKm     ContentSearchResultLanguage = "km"
	ContentSearchResultLanguageKmKh   ContentSearchResultLanguage = "km-kh"
	ContentSearchResultLanguageKn     ContentSearchResultLanguage = "kn"
	ContentSearchResultLanguageKnIn   ContentSearchResultLanguage = "kn-in"
	ContentSearchResultLanguageKo     ContentSearchResultLanguage = "ko"
	ContentSearchResultLanguageKoKp   ContentSearchResultLanguage = "ko-kp"
	ContentSearchResultLanguageKoKr   ContentSearchResultLanguage = "ko-kr"
	ContentSearchResultLanguageKok    ContentSearchResultLanguage = "kok"
	ContentSearchResultLanguageKokIn  ContentSearchResultLanguage = "kok-in"
	ContentSearchResultLanguageKs     ContentSearchResultLanguage = "ks"
	ContentSearchResultLanguageKsIn   ContentSearchResultLanguage = "ks-in"
	ContentSearchResultLanguageKsb    ContentSearchResultLanguage = "ksb"
	ContentSearchResultLanguageKsbTz  ContentSearchResultLanguage = "ksb-tz"
	ContentSearchResultLanguageKsf    ContentSearchResultLanguage = "ksf"
	ContentSearchResultLanguageKsfCm  ContentSearchResultLanguage = "ksf-cm"
	ContentSearchResultLanguageKsh    ContentSearchResultLanguage = "ksh"
	ContentSearchResultLanguageKshDe  ContentSearchResultLanguage = "ksh-de"
	ContentSearchResultLanguageKu     ContentSearchResultLanguage = "ku"
	ContentSearchResultLanguageKuTr   ContentSearchResultLanguage = "ku-tr"
	ContentSearchResultLanguageKw     ContentSearchResultLanguage = "kw"
	ContentSearchResultLanguageKwGB   ContentSearchResultLanguage = "kw-gb"
	ContentSearchResultLanguageKy     ContentSearchResultLanguage = "ky"
	ContentSearchResultLanguageKyKg   ContentSearchResultLanguage = "ky-kg"
	ContentSearchResultLanguageLag    ContentSearchResultLanguage = "lag"
	ContentSearchResultLanguageLagTz  ContentSearchResultLanguage = "lag-tz"
	ContentSearchResultLanguageLb     ContentSearchResultLanguage = "lb"
	ContentSearchResultLanguageLbLu   ContentSearchResultLanguage = "lb-lu"
	ContentSearchResultLanguageLg     ContentSearchResultLanguage = "lg"
	ContentSearchResultLanguageLgUg   ContentSearchResultLanguage = "lg-ug"
	ContentSearchResultLanguageLkt    ContentSearchResultLanguage = "lkt"
	ContentSearchResultLanguageLktUs  ContentSearchResultLanguage = "lkt-us"
	ContentSearchResultLanguageLn     ContentSearchResultLanguage = "ln"
	ContentSearchResultLanguageLnAo   ContentSearchResultLanguage = "ln-ao"
	ContentSearchResultLanguageLnCd   ContentSearchResultLanguage = "ln-cd"
	ContentSearchResultLanguageLnCf   ContentSearchResultLanguage = "ln-cf"
	ContentSearchResultLanguageLnCg   ContentSearchResultLanguage = "ln-cg"
	ContentSearchResultLanguageLo     ContentSearchResultLanguage = "lo"
	ContentSearchResultLanguageLoLa   ContentSearchResultLanguage = "lo-la"
	ContentSearchResultLanguageLrc    ContentSearchResultLanguage = "lrc"
	ContentSearchResultLanguageLrcIq  ContentSearchResultLanguage = "lrc-iq"
	ContentSearchResultLanguageLrcIr  ContentSearchResultLanguage = "lrc-ir"
	ContentSearchResultLanguageLt     ContentSearchResultLanguage = "lt"
	ContentSearchResultLanguageLtLt   ContentSearchResultLanguage = "lt-lt"
	ContentSearchResultLanguageLu     ContentSearchResultLanguage = "lu"
	ContentSearchResultLanguageLuCd   ContentSearchResultLanguage = "lu-cd"
	ContentSearchResultLanguageLuo    ContentSearchResultLanguage = "luo"
	ContentSearchResultLanguageLuoKe  ContentSearchResultLanguage = "luo-ke"
	ContentSearchResultLanguageLuy    ContentSearchResultLanguage = "luy"
	ContentSearchResultLanguageLuyKe  ContentSearchResultLanguage = "luy-ke"
	ContentSearchResultLanguageLv     ContentSearchResultLanguage = "lv"
	ContentSearchResultLanguageLvLv   ContentSearchResultLanguage = "lv-lv"
	ContentSearchResultLanguageMai    ContentSearchResultLanguage = "mai"
	ContentSearchResultLanguageMaiIn  ContentSearchResultLanguage = "mai-in"
	ContentSearchResultLanguageMas    ContentSearchResultLanguage = "mas"
	ContentSearchResultLanguageMasKe  ContentSearchResultLanguage = "mas-ke"
	ContentSearchResultLanguageMasTz  ContentSearchResultLanguage = "mas-tz"
	ContentSearchResultLanguageMer    ContentSearchResultLanguage = "mer"
	ContentSearchResultLanguageMerKe  ContentSearchResultLanguage = "mer-ke"
	ContentSearchResultLanguageMfe    ContentSearchResultLanguage = "mfe"
	ContentSearchResultLanguageMfeMu  ContentSearchResultLanguage = "mfe-mu"
	ContentSearchResultLanguageMg     ContentSearchResultLanguage = "mg"
	ContentSearchResultLanguageMgMg   ContentSearchResultLanguage = "mg-mg"
	ContentSearchResultLanguageMgh    ContentSearchResultLanguage = "mgh"
	ContentSearchResultLanguageMghMz  ContentSearchResultLanguage = "mgh-mz"
	ContentSearchResultLanguageMgo    ContentSearchResultLanguage = "mgo"
	ContentSearchResultLanguageMgoCm  ContentSearchResultLanguage = "mgo-cm"
	ContentSearchResultLanguageMi     ContentSearchResultLanguage = "mi"
	ContentSearchResultLanguageMiNz   ContentSearchResultLanguage = "mi-nz"
	ContentSearchResultLanguageMk     ContentSearchResultLanguage = "mk"
	ContentSearchResultLanguageMkMk   ContentSearchResultLanguage = "mk-mk"
	ContentSearchResultLanguageMl     ContentSearchResultLanguage = "ml"
	ContentSearchResultLanguageMlIn   ContentSearchResultLanguage = "ml-in"
	ContentSearchResultLanguageMn     ContentSearchResultLanguage = "mn"
	ContentSearchResultLanguageMnMn   ContentSearchResultLanguage = "mn-mn"
	ContentSearchResultLanguageMni    ContentSearchResultLanguage = "mni"
	ContentSearchResultLanguageMniIn  ContentSearchResultLanguage = "mni-in"
	ContentSearchResultLanguageMr     ContentSearchResultLanguage = "mr"
	ContentSearchResultLanguageMrIn   ContentSearchResultLanguage = "mr-in"
	ContentSearchResultLanguageMs     ContentSearchResultLanguage = "ms"
	ContentSearchResultLanguageMsBn   ContentSearchResultLanguage = "ms-bn"
	ContentSearchResultLanguageMsID   ContentSearchResultLanguage = "ms-id"
	ContentSearchResultLanguageMsMy   ContentSearchResultLanguage = "ms-my"
	ContentSearchResultLanguageMsSg   ContentSearchResultLanguage = "ms-sg"
	ContentSearchResultLanguageMt     ContentSearchResultLanguage = "mt"
	ContentSearchResultLanguageMtMt   ContentSearchResultLanguage = "mt-mt"
	ContentSearchResultLanguageMua    ContentSearchResultLanguage = "mua"
	ContentSearchResultLanguageMuaCm  ContentSearchResultLanguage = "mua-cm"
	ContentSearchResultLanguageMy     ContentSearchResultLanguage = "my"
	ContentSearchResultLanguageMyMm   ContentSearchResultLanguage = "my-mm"
	ContentSearchResultLanguageMzn    ContentSearchResultLanguage = "mzn"
	ContentSearchResultLanguageMznIr  ContentSearchResultLanguage = "mzn-ir"
	ContentSearchResultLanguageNaq    ContentSearchResultLanguage = "naq"
	ContentSearchResultLanguageNaqNa  ContentSearchResultLanguage = "naq-na"
	ContentSearchResultLanguageNb     ContentSearchResultLanguage = "nb"
	ContentSearchResultLanguageNbNo   ContentSearchResultLanguage = "nb-no"
	ContentSearchResultLanguageNbSj   ContentSearchResultLanguage = "nb-sj"
	ContentSearchResultLanguageNd     ContentSearchResultLanguage = "nd"
	ContentSearchResultLanguageNdZw   ContentSearchResultLanguage = "nd-zw"
	ContentSearchResultLanguageNds    ContentSearchResultLanguage = "nds"
	ContentSearchResultLanguageNdsDe  ContentSearchResultLanguage = "nds-de"
	ContentSearchResultLanguageNdsNl  ContentSearchResultLanguage = "nds-nl"
	ContentSearchResultLanguageNe     ContentSearchResultLanguage = "ne"
	ContentSearchResultLanguageNeIn   ContentSearchResultLanguage = "ne-in"
	ContentSearchResultLanguageNeNp   ContentSearchResultLanguage = "ne-np"
	ContentSearchResultLanguageNl     ContentSearchResultLanguage = "nl"
	ContentSearchResultLanguageNlAw   ContentSearchResultLanguage = "nl-aw"
	ContentSearchResultLanguageNlBe   ContentSearchResultLanguage = "nl-be"
	ContentSearchResultLanguageNlBq   ContentSearchResultLanguage = "nl-bq"
	ContentSearchResultLanguageNlCh   ContentSearchResultLanguage = "nl-ch"
	ContentSearchResultLanguageNlCw   ContentSearchResultLanguage = "nl-cw"
	ContentSearchResultLanguageNlLu   ContentSearchResultLanguage = "nl-lu"
	ContentSearchResultLanguageNlNl   ContentSearchResultLanguage = "nl-nl"
	ContentSearchResultLanguageNlSr   ContentSearchResultLanguage = "nl-sr"
	ContentSearchResultLanguageNlSx   ContentSearchResultLanguage = "nl-sx"
	ContentSearchResultLanguageNmg    ContentSearchResultLanguage = "nmg"
	ContentSearchResultLanguageNmgCm  ContentSearchResultLanguage = "nmg-cm"
	ContentSearchResultLanguageNn     ContentSearchResultLanguage = "nn"
	ContentSearchResultLanguageNnNo   ContentSearchResultLanguage = "nn-no"
	ContentSearchResultLanguageNnh    ContentSearchResultLanguage = "nnh"
	ContentSearchResultLanguageNnhCm  ContentSearchResultLanguage = "nnh-cm"
	ContentSearchResultLanguageNo     ContentSearchResultLanguage = "no"
	ContentSearchResultLanguageNoNo   ContentSearchResultLanguage = "no-no"
	ContentSearchResultLanguageNus    ContentSearchResultLanguage = "nus"
	ContentSearchResultLanguageNusSS  ContentSearchResultLanguage = "nus-ss"
	ContentSearchResultLanguageNyn    ContentSearchResultLanguage = "nyn"
	ContentSearchResultLanguageNynUg  ContentSearchResultLanguage = "nyn-ug"
	ContentSearchResultLanguageOm     ContentSearchResultLanguage = "om"
	ContentSearchResultLanguageOmEt   ContentSearchResultLanguage = "om-et"
	ContentSearchResultLanguageOmKe   ContentSearchResultLanguage = "om-ke"
	ContentSearchResultLanguageOr     ContentSearchResultLanguage = "or"
	ContentSearchResultLanguageOrIn   ContentSearchResultLanguage = "or-in"
	ContentSearchResultLanguageOs     ContentSearchResultLanguage = "os"
	ContentSearchResultLanguageOsGe   ContentSearchResultLanguage = "os-ge"
	ContentSearchResultLanguageOsRu   ContentSearchResultLanguage = "os-ru"
	ContentSearchResultLanguagePa     ContentSearchResultLanguage = "pa"
	ContentSearchResultLanguagePaIn   ContentSearchResultLanguage = "pa-in"
	ContentSearchResultLanguagePaPk   ContentSearchResultLanguage = "pa-pk"
	ContentSearchResultLanguagePcm    ContentSearchResultLanguage = "pcm"
	ContentSearchResultLanguagePcmNg  ContentSearchResultLanguage = "pcm-ng"
	ContentSearchResultLanguagePl     ContentSearchResultLanguage = "pl"
	ContentSearchResultLanguagePlPl   ContentSearchResultLanguage = "pl-pl"
	ContentSearchResultLanguagePrg    ContentSearchResultLanguage = "prg"
	ContentSearchResultLanguagePrg001 ContentSearchResultLanguage = "prg-001"
	ContentSearchResultLanguagePs     ContentSearchResultLanguage = "ps"
	ContentSearchResultLanguagePsAf   ContentSearchResultLanguage = "ps-af"
	ContentSearchResultLanguagePsPk   ContentSearchResultLanguage = "ps-pk"
	ContentSearchResultLanguagePt     ContentSearchResultLanguage = "pt"
	ContentSearchResultLanguagePtAo   ContentSearchResultLanguage = "pt-ao"
	ContentSearchResultLanguagePtBr   ContentSearchResultLanguage = "pt-br"
	ContentSearchResultLanguagePtCh   ContentSearchResultLanguage = "pt-ch"
	ContentSearchResultLanguagePtCv   ContentSearchResultLanguage = "pt-cv"
	ContentSearchResultLanguagePtGq   ContentSearchResultLanguage = "pt-gq"
	ContentSearchResultLanguagePtGw   ContentSearchResultLanguage = "pt-gw"
	ContentSearchResultLanguagePtLu   ContentSearchResultLanguage = "pt-lu"
	ContentSearchResultLanguagePtMo   ContentSearchResultLanguage = "pt-mo"
	ContentSearchResultLanguagePtMz   ContentSearchResultLanguage = "pt-mz"
	ContentSearchResultLanguagePtPt   ContentSearchResultLanguage = "pt-pt"
	ContentSearchResultLanguagePtSt   ContentSearchResultLanguage = "pt-st"
	ContentSearchResultLanguagePtTl   ContentSearchResultLanguage = "pt-tl"
	ContentSearchResultLanguageQu     ContentSearchResultLanguage = "qu"
	ContentSearchResultLanguageQuBo   ContentSearchResultLanguage = "qu-bo"
	ContentSearchResultLanguageQuEc   ContentSearchResultLanguage = "qu-ec"
	ContentSearchResultLanguageQuPe   ContentSearchResultLanguage = "qu-pe"
	ContentSearchResultLanguageRm     ContentSearchResultLanguage = "rm"
	ContentSearchResultLanguageRmCh   ContentSearchResultLanguage = "rm-ch"
	ContentSearchResultLanguageRn     ContentSearchResultLanguage = "rn"
	ContentSearchResultLanguageRnBi   ContentSearchResultLanguage = "rn-bi"
	ContentSearchResultLanguageRo     ContentSearchResultLanguage = "ro"
	ContentSearchResultLanguageRoMd   ContentSearchResultLanguage = "ro-md"
	ContentSearchResultLanguageRoRo   ContentSearchResultLanguage = "ro-ro"
	ContentSearchResultLanguageRof    ContentSearchResultLanguage = "rof"
	ContentSearchResultLanguageRofTz  ContentSearchResultLanguage = "rof-tz"
	ContentSearchResultLanguageRu     ContentSearchResultLanguage = "ru"
	ContentSearchResultLanguageRuBy   ContentSearchResultLanguage = "ru-by"
	ContentSearchResultLanguageRuKg   ContentSearchResultLanguage = "ru-kg"
	ContentSearchResultLanguageRuKz   ContentSearchResultLanguage = "ru-kz"
	ContentSearchResultLanguageRuMd   ContentSearchResultLanguage = "ru-md"
	ContentSearchResultLanguageRuRu   ContentSearchResultLanguage = "ru-ru"
	ContentSearchResultLanguageRuUa   ContentSearchResultLanguage = "ru-ua"
	ContentSearchResultLanguageRw     ContentSearchResultLanguage = "rw"
	ContentSearchResultLanguageRwRw   ContentSearchResultLanguage = "rw-rw"
	ContentSearchResultLanguageRwk    ContentSearchResultLanguage = "rwk"
	ContentSearchResultLanguageRwkTz  ContentSearchResultLanguage = "rwk-tz"
	ContentSearchResultLanguageSa     ContentSearchResultLanguage = "sa"
	ContentSearchResultLanguageSaIn   ContentSearchResultLanguage = "sa-in"
	ContentSearchResultLanguageSah    ContentSearchResultLanguage = "sah"
	ContentSearchResultLanguageSahRu  ContentSearchResultLanguage = "sah-ru"
	ContentSearchResultLanguageSaq    ContentSearchResultLanguage = "saq"
	ContentSearchResultLanguageSaqKe  ContentSearchResultLanguage = "saq-ke"
	ContentSearchResultLanguageSat    ContentSearchResultLanguage = "sat"
	ContentSearchResultLanguageSatIn  ContentSearchResultLanguage = "sat-in"
	ContentSearchResultLanguageSbp    ContentSearchResultLanguage = "sbp"
	ContentSearchResultLanguageSbpTz  ContentSearchResultLanguage = "sbp-tz"
	ContentSearchResultLanguageSd     ContentSearchResultLanguage = "sd"
	ContentSearchResultLanguageSdIn   ContentSearchResultLanguage = "sd-in"
	ContentSearchResultLanguageSdPk   ContentSearchResultLanguage = "sd-pk"
	ContentSearchResultLanguageSe     ContentSearchResultLanguage = "se"
	ContentSearchResultLanguageSeFi   ContentSearchResultLanguage = "se-fi"
	ContentSearchResultLanguageSeNo   ContentSearchResultLanguage = "se-no"
	ContentSearchResultLanguageSeSe   ContentSearchResultLanguage = "se-se"
	ContentSearchResultLanguageSeh    ContentSearchResultLanguage = "seh"
	ContentSearchResultLanguageSehMz  ContentSearchResultLanguage = "seh-mz"
	ContentSearchResultLanguageSes    ContentSearchResultLanguage = "ses"
	ContentSearchResultLanguageSesMl  ContentSearchResultLanguage = "ses-ml"
	ContentSearchResultLanguageSg     ContentSearchResultLanguage = "sg"
	ContentSearchResultLanguageSgCf   ContentSearchResultLanguage = "sg-cf"
	ContentSearchResultLanguageShi    ContentSearchResultLanguage = "shi"
	ContentSearchResultLanguageShiMa  ContentSearchResultLanguage = "shi-ma"
	ContentSearchResultLanguageSi     ContentSearchResultLanguage = "si"
	ContentSearchResultLanguageSiLk   ContentSearchResultLanguage = "si-lk"
	ContentSearchResultLanguageSk     ContentSearchResultLanguage = "sk"
	ContentSearchResultLanguageSkSk   ContentSearchResultLanguage = "sk-sk"
	ContentSearchResultLanguageSl     ContentSearchResultLanguage = "sl"
	ContentSearchResultLanguageSlSi   ContentSearchResultLanguage = "sl-si"
	ContentSearchResultLanguageSmn    ContentSearchResultLanguage = "smn"
	ContentSearchResultLanguageSmnFi  ContentSearchResultLanguage = "smn-fi"
	ContentSearchResultLanguageSn     ContentSearchResultLanguage = "sn"
	ContentSearchResultLanguageSnZw   ContentSearchResultLanguage = "sn-zw"
	ContentSearchResultLanguageSo     ContentSearchResultLanguage = "so"
	ContentSearchResultLanguageSoDj   ContentSearchResultLanguage = "so-dj"
	ContentSearchResultLanguageSoEt   ContentSearchResultLanguage = "so-et"
	ContentSearchResultLanguageSoKe   ContentSearchResultLanguage = "so-ke"
	ContentSearchResultLanguageSoSo   ContentSearchResultLanguage = "so-so"
	ContentSearchResultLanguageSq     ContentSearchResultLanguage = "sq"
	ContentSearchResultLanguageSqAl   ContentSearchResultLanguage = "sq-al"
	ContentSearchResultLanguageSqMk   ContentSearchResultLanguage = "sq-mk"
	ContentSearchResultLanguageSqXk   ContentSearchResultLanguage = "sq-xk"
	ContentSearchResultLanguageSr     ContentSearchResultLanguage = "sr"
	ContentSearchResultLanguageSrBa   ContentSearchResultLanguage = "sr-ba"
	ContentSearchResultLanguageSrCs   ContentSearchResultLanguage = "sr-cs"
	ContentSearchResultLanguageSrMe   ContentSearchResultLanguage = "sr-me"
	ContentSearchResultLanguageSrRs   ContentSearchResultLanguage = "sr-rs"
	ContentSearchResultLanguageSrXk   ContentSearchResultLanguage = "sr-xk"
	ContentSearchResultLanguageSu     ContentSearchResultLanguage = "su"
	ContentSearchResultLanguageSuID   ContentSearchResultLanguage = "su-id"
	ContentSearchResultLanguageSv     ContentSearchResultLanguage = "sv"
	ContentSearchResultLanguageSvAx   ContentSearchResultLanguage = "sv-ax"
	ContentSearchResultLanguageSvFi   ContentSearchResultLanguage = "sv-fi"
	ContentSearchResultLanguageSvSe   ContentSearchResultLanguage = "sv-se"
	ContentSearchResultLanguageSw     ContentSearchResultLanguage = "sw"
	ContentSearchResultLanguageSwCd   ContentSearchResultLanguage = "sw-cd"
	ContentSearchResultLanguageSwKe   ContentSearchResultLanguage = "sw-ke"
	ContentSearchResultLanguageSwTz   ContentSearchResultLanguage = "sw-tz"
	ContentSearchResultLanguageSwUg   ContentSearchResultLanguage = "sw-ug"
	ContentSearchResultLanguageSy     ContentSearchResultLanguage = "sy"
	ContentSearchResultLanguageTa     ContentSearchResultLanguage = "ta"
	ContentSearchResultLanguageTaIn   ContentSearchResultLanguage = "ta-in"
	ContentSearchResultLanguageTaLk   ContentSearchResultLanguage = "ta-lk"
	ContentSearchResultLanguageTaMy   ContentSearchResultLanguage = "ta-my"
	ContentSearchResultLanguageTaSg   ContentSearchResultLanguage = "ta-sg"
	ContentSearchResultLanguageTe     ContentSearchResultLanguage = "te"
	ContentSearchResultLanguageTeIn   ContentSearchResultLanguage = "te-in"
	ContentSearchResultLanguageTeo    ContentSearchResultLanguage = "teo"
	ContentSearchResultLanguageTeoKe  ContentSearchResultLanguage = "teo-ke"
	ContentSearchResultLanguageTeoUg  ContentSearchResultLanguage = "teo-ug"
	ContentSearchResultLanguageTg     ContentSearchResultLanguage = "tg"
	ContentSearchResultLanguageTgTj   ContentSearchResultLanguage = "tg-tj"
	ContentSearchResultLanguageTh     ContentSearchResultLanguage = "th"
	ContentSearchResultLanguageThTh   ContentSearchResultLanguage = "th-th"
	ContentSearchResultLanguageTi     ContentSearchResultLanguage = "ti"
	ContentSearchResultLanguageTiEr   ContentSearchResultLanguage = "ti-er"
	ContentSearchResultLanguageTiEt   ContentSearchResultLanguage = "ti-et"
	ContentSearchResultLanguageTk     ContentSearchResultLanguage = "tk"
	ContentSearchResultLanguageTkTm   ContentSearchResultLanguage = "tk-tm"
	ContentSearchResultLanguageTl     ContentSearchResultLanguage = "tl"
	ContentSearchResultLanguageTo     ContentSearchResultLanguage = "to"
	ContentSearchResultLanguageToTo   ContentSearchResultLanguage = "to-to"
	ContentSearchResultLanguageTr     ContentSearchResultLanguage = "tr"
	ContentSearchResultLanguageTrCy   ContentSearchResultLanguage = "tr-cy"
	ContentSearchResultLanguageTrTr   ContentSearchResultLanguage = "tr-tr"
	ContentSearchResultLanguageTt     ContentSearchResultLanguage = "tt"
	ContentSearchResultLanguageTtRu   ContentSearchResultLanguage = "tt-ru"
	ContentSearchResultLanguageTwq    ContentSearchResultLanguage = "twq"
	ContentSearchResultLanguageTwqNe  ContentSearchResultLanguage = "twq-ne"
	ContentSearchResultLanguageTzm    ContentSearchResultLanguage = "tzm"
	ContentSearchResultLanguageTzmMa  ContentSearchResultLanguage = "tzm-ma"
	ContentSearchResultLanguageUg     ContentSearchResultLanguage = "ug"
	ContentSearchResultLanguageUgCn   ContentSearchResultLanguage = "ug-cn"
	ContentSearchResultLanguageUk     ContentSearchResultLanguage = "uk"
	ContentSearchResultLanguageUkUa   ContentSearchResultLanguage = "uk-ua"
	ContentSearchResultLanguageUr     ContentSearchResultLanguage = "ur"
	ContentSearchResultLanguageUrIn   ContentSearchResultLanguage = "ur-in"
	ContentSearchResultLanguageUrPk   ContentSearchResultLanguage = "ur-pk"
	ContentSearchResultLanguageUz     ContentSearchResultLanguage = "uz"
	ContentSearchResultLanguageUzAf   ContentSearchResultLanguage = "uz-af"
	ContentSearchResultLanguageUzUz   ContentSearchResultLanguage = "uz-uz"
	ContentSearchResultLanguageVai    ContentSearchResultLanguage = "vai"
	ContentSearchResultLanguageVaiLr  ContentSearchResultLanguage = "vai-lr"
	ContentSearchResultLanguageVi     ContentSearchResultLanguage = "vi"
	ContentSearchResultLanguageViVn   ContentSearchResultLanguage = "vi-vn"
	ContentSearchResultLanguageVo     ContentSearchResultLanguage = "vo"
	ContentSearchResultLanguageVo001  ContentSearchResultLanguage = "vo-001"
	ContentSearchResultLanguageVun    ContentSearchResultLanguage = "vun"
	ContentSearchResultLanguageVunTz  ContentSearchResultLanguage = "vun-tz"
	ContentSearchResultLanguageWae    ContentSearchResultLanguage = "wae"
	ContentSearchResultLanguageWaeCh  ContentSearchResultLanguage = "wae-ch"
	ContentSearchResultLanguageWo     ContentSearchResultLanguage = "wo"
	ContentSearchResultLanguageWoSn   ContentSearchResultLanguage = "wo-sn"
	ContentSearchResultLanguageXh     ContentSearchResultLanguage = "xh"
	ContentSearchResultLanguageXhZa   ContentSearchResultLanguage = "xh-za"
	ContentSearchResultLanguageXog    ContentSearchResultLanguage = "xog"
	ContentSearchResultLanguageXogUg  ContentSearchResultLanguage = "xog-ug"
	ContentSearchResultLanguageYav    ContentSearchResultLanguage = "yav"
	ContentSearchResultLanguageYavCm  ContentSearchResultLanguage = "yav-cm"
	ContentSearchResultLanguageYi     ContentSearchResultLanguage = "yi"
	ContentSearchResultLanguageYi001  ContentSearchResultLanguage = "yi-001"
	ContentSearchResultLanguageYo     ContentSearchResultLanguage = "yo"
	ContentSearchResultLanguageYoBj   ContentSearchResultLanguage = "yo-bj"
	ContentSearchResultLanguageYoNg   ContentSearchResultLanguage = "yo-ng"
	ContentSearchResultLanguageYue    ContentSearchResultLanguage = "yue"
	ContentSearchResultLanguageYueCn  ContentSearchResultLanguage = "yue-cn"
	ContentSearchResultLanguageYueHk  ContentSearchResultLanguage = "yue-hk"
	ContentSearchResultLanguageZgh    ContentSearchResultLanguage = "zgh"
	ContentSearchResultLanguageZghMa  ContentSearchResultLanguage = "zgh-ma"
	ContentSearchResultLanguageZh     ContentSearchResultLanguage = "zh"
	ContentSearchResultLanguageZhCn   ContentSearchResultLanguage = "zh-cn"
	ContentSearchResultLanguageZhHans ContentSearchResultLanguage = "zh-hans"
	ContentSearchResultLanguageZhHant ContentSearchResultLanguage = "zh-hant"
	ContentSearchResultLanguageZhHk   ContentSearchResultLanguage = "zh-hk"
	ContentSearchResultLanguageZhMo   ContentSearchResultLanguage = "zh-mo"
	ContentSearchResultLanguageZhSg   ContentSearchResultLanguage = "zh-sg"
	ContentSearchResultLanguageZhTw   ContentSearchResultLanguage = "zh-tw"
	ContentSearchResultLanguageZu     ContentSearchResultLanguage = "zu"
	ContentSearchResultLanguageZuZa   ContentSearchResultLanguage = "zu-za"
)

// The indexed data in HubSpot
type IndexedData struct {
	// The ID of the document in HubSpot.
	ID string `json:"id,required"`
	// The indexed fields in HubSpot.
	Fields map[string]IndexedField `json:"fields,required"`
	// The type of document. Can be `SITE_PAGE`, `LANDING_PAGE`, `BLOG_POST`,
	// `LISTING_PAGE`, or `KNOWLEDGE_ARTICLE`.
	//
	// Any of "BLOG_POST", "KNOWLEDGE_ARTICLE", "LANDING_PAGE", "LISTING_PAGE",
	// "SITE_PAGE".
	Type IndexedDataType `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Fields      respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IndexedData) RawJSON() string { return r.JSON.raw }
func (r *IndexedData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The type of document. Can be `SITE_PAGE`, `LANDING_PAGE`, `BLOG_POST`,
// `LISTING_PAGE`, or `KNOWLEDGE_ARTICLE`.
type IndexedDataType string

const (
	IndexedDataTypeBlogPost         IndexedDataType = "BLOG_POST"
	IndexedDataTypeKnowledgeArticle IndexedDataType = "KNOWLEDGE_ARTICLE"
	IndexedDataTypeLandingPage      IndexedDataType = "LANDING_PAGE"
	IndexedDataTypeListingPage      IndexedDataType = "LISTING_PAGE"
	IndexedDataTypeSitePage         IndexedDataType = "SITE_PAGE"
)

type IndexedField struct {
	MetadataField bool   `json:"metadataField,required"`
	Name          string `json:"name,required"`
	Value         any    `json:"value,required"`
	Values        []any  `json:"values,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		MetadataField respjson.Field
		Name          respjson.Field
		Value         respjson.Field
		Values        respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IndexedField) RawJSON() string { return r.JSON.raw }
func (r *IndexedField) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PublicSearchResults struct {
	Limit      int64                 `json:"limit,required"`
	Offset     int64                 `json:"offset,required"`
	Page       int64                 `json:"page,required"`
	Results    []ContentSearchResult `json:"results,required"`
	Total      int64                 `json:"total,required"`
	SearchTerm string                `json:"searchTerm"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Limit       respjson.Field
		Offset      respjson.Field
		Page        respjson.Field
		Results     respjson.Field
		Total       respjson.Field
		SearchTerm  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PublicSearchResults) RawJSON() string { return r.JSON.raw }
func (r *PublicSearchResults) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SiteSearchGetIndexedDataParams struct {
	// The type of document. Can be one of `SITE_PAGE`, `BLOG_POST`, or
	// `KNOWLEDGE_ARTICLE`.
	//
	// Any of "BLOG_POST", "KNOWLEDGE_ARTICLE", "LANDING_PAGE", "LISTING_PAGE",
	// "SITE_PAGE".
	Type SiteSearchGetIndexedDataParamsType `query:"type,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SiteSearchGetIndexedDataParams]'s query parameters as
// `url.Values`.
func (r SiteSearchGetIndexedDataParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// The type of document. Can be one of `SITE_PAGE`, `BLOG_POST`, or
// `KNOWLEDGE_ARTICLE`.
type SiteSearchGetIndexedDataParamsType string

const (
	SiteSearchGetIndexedDataParamsTypeBlogPost         SiteSearchGetIndexedDataParamsType = "BLOG_POST"
	SiteSearchGetIndexedDataParamsTypeKnowledgeArticle SiteSearchGetIndexedDataParamsType = "KNOWLEDGE_ARTICLE"
	SiteSearchGetIndexedDataParamsTypeLandingPage      SiteSearchGetIndexedDataParamsType = "LANDING_PAGE"
	SiteSearchGetIndexedDataParamsTypeListingPage      SiteSearchGetIndexedDataParamsType = "LISTING_PAGE"
	SiteSearchGetIndexedDataParamsTypeSitePage         SiteSearchGetIndexedDataParamsType = "SITE_PAGE"
)

type SiteSearchSearchParams struct {
	// Specifies whether or not you are showing autocomplete results. Defaults to
	// false.
	Autocomplete param.Opt[bool] `query:"autocomplete,omitzero" json:"-"`
	// Specifies the maximum amount a result will be boosted based on its view count.
	// Defaults to 5.0. Read more about elasticsearch boosting
	// [here](https://www.elastic.co/guide/en/elasticsearch/reference/current/mapping-boost.html).
	BoostLimit param.Opt[float64] `query:"boostLimit,omitzero" json:"-"`
	// Specifies a relative time window where scores of documents published outside
	// this time window decay. This can only be used for blog posts. For example,
	// boostRecent=10d will boost documents published within the last 10 days.
	// Supported timeunits are ms (milliseconds), s (seconds), m (minutes), h (hours),
	// d (days).
	BoostRecent param.Opt[string] `query:"boostRecent,omitzero" json:"-"`
	// Specify a HubDB query to further filter the search results.
	HubdbQuery param.Opt[string] `query:"hubdbQuery,omitzero" json:"-"`
	// Specifies the number of results to be returned in a single response. Defaults to
	// `10`. Maximum value is `100`.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Inverts the behavior of the pathPrefix filter when set to `false`. Defaults to
	// `true`.
	MatchPrefix param.Opt[bool] `query:"matchPrefix,omitzero" json:"-"`
	// Used to page through the results. If there are more results than specified by
	// the `limit` parameter, you will need to use the value of offset returned in the
	// previous request to get the next set of results.
	Offset param.Opt[int64] `query:"offset,omitzero" json:"-"`
	// Specifies how strongly a result is boosted based on its view count. Defaults to
	// 1.0.
	PopularityBoost param.Opt[float64] `query:"popularityBoost,omitzero" json:"-"`
	// The term to search for.
	Q param.Opt[string] `query:"q,omitzero" json:"-"`
	// Specifies a specific HubDB table to search. Only returns results from the
	// specified table. Can be used in tandem with the `hubdbQuery` parameter to
	// further filter results.
	TableID param.Opt[int64] `query:"tableId,omitzero" json:"-"`
	// A domain to match search results for. Multiple domains can be provided with &.
	Domain []string `query:"domain,omitzero" json:"-"`
	// Specifies which blog(s) to be searched by blog ID. Can be used multiple times to
	// search more than one blog.
	GroupID []int64 `query:"groupId,omitzero" json:"-"`
	// Specifies the language of content to be searched. This value must be a valid
	// [ISO 639-1 language code](https://en.wikipedia.org/wiki/List_of_ISO_639-1_codes)
	// (e.g. `es` for Spanish)
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
	Language SiteSearchSearchParamsLanguage `query:"language,omitzero" json:"-"`
	// Specifies the length of the search results. Can be set to `LONG` or `SHORT`.
	// `SHORT` will return the first 128 characters of the content's meta description.
	// `LONG` will build a more detailed content snippet based on the html/content of
	// the page.
	//
	// Any of "LONG", "SHORT".
	Length SiteSearchSearchParamsLength `query:"length,omitzero" json:"-"`
	// Specifies a path prefix to filter search results. Will only return results with
	// URL paths that start with the specified parameter. Can be used multiple times.
	PathPrefix []string `query:"pathPrefix,omitzero" json:"-"`
	// Specifies which properties to include in the search. Options include `title`,
	// `description`, and `html`. All properties will be searched by default.
	Property []string `query:"property,omitzero" json:"-"`
	// Specifies the type of content to search. Can be one or more of SITE_PAGE,
	// LANDING_PAGE, BLOG_POST, LISTING_PAGE, and KNOWLEDGE_ARTICLE. Defaults to all
	// content types except LANDING_PAGE and KNOWLEDGE_ARTICLE
	//
	// Any of "LANDING_PAGE", "BLOG_POST", "SITE_PAGE", "KNOWLEDGE_ARTICLE",
	// "LISTING_PAGE".
	Type []string `query:"type,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SiteSearchSearchParams]'s query parameters as `url.Values`.
func (r SiteSearchSearchParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Specifies the language of content to be searched. This value must be a valid
// [ISO 639-1 language code](https://en.wikipedia.org/wiki/List_of_ISO_639-1_codes)
// (e.g. `es` for Spanish)
type SiteSearchSearchParamsLanguage string

const (
	SiteSearchSearchParamsLanguageAf     SiteSearchSearchParamsLanguage = "af"
	SiteSearchSearchParamsLanguageAfNa   SiteSearchSearchParamsLanguage = "af-na"
	SiteSearchSearchParamsLanguageAfZa   SiteSearchSearchParamsLanguage = "af-za"
	SiteSearchSearchParamsLanguageAgq    SiteSearchSearchParamsLanguage = "agq"
	SiteSearchSearchParamsLanguageAgqCm  SiteSearchSearchParamsLanguage = "agq-cm"
	SiteSearchSearchParamsLanguageAk     SiteSearchSearchParamsLanguage = "ak"
	SiteSearchSearchParamsLanguageAkGh   SiteSearchSearchParamsLanguage = "ak-gh"
	SiteSearchSearchParamsLanguageAm     SiteSearchSearchParamsLanguage = "am"
	SiteSearchSearchParamsLanguageAmEt   SiteSearchSearchParamsLanguage = "am-et"
	SiteSearchSearchParamsLanguageAr     SiteSearchSearchParamsLanguage = "ar"
	SiteSearchSearchParamsLanguageAr001  SiteSearchSearchParamsLanguage = "ar-001"
	SiteSearchSearchParamsLanguageArAe   SiteSearchSearchParamsLanguage = "ar-ae"
	SiteSearchSearchParamsLanguageArBh   SiteSearchSearchParamsLanguage = "ar-bh"
	SiteSearchSearchParamsLanguageArDj   SiteSearchSearchParamsLanguage = "ar-dj"
	SiteSearchSearchParamsLanguageArDz   SiteSearchSearchParamsLanguage = "ar-dz"
	SiteSearchSearchParamsLanguageArEg   SiteSearchSearchParamsLanguage = "ar-eg"
	SiteSearchSearchParamsLanguageArEh   SiteSearchSearchParamsLanguage = "ar-eh"
	SiteSearchSearchParamsLanguageArEr   SiteSearchSearchParamsLanguage = "ar-er"
	SiteSearchSearchParamsLanguageArIl   SiteSearchSearchParamsLanguage = "ar-il"
	SiteSearchSearchParamsLanguageArIq   SiteSearchSearchParamsLanguage = "ar-iq"
	SiteSearchSearchParamsLanguageArJo   SiteSearchSearchParamsLanguage = "ar-jo"
	SiteSearchSearchParamsLanguageArKm   SiteSearchSearchParamsLanguage = "ar-km"
	SiteSearchSearchParamsLanguageArKw   SiteSearchSearchParamsLanguage = "ar-kw"
	SiteSearchSearchParamsLanguageArLb   SiteSearchSearchParamsLanguage = "ar-lb"
	SiteSearchSearchParamsLanguageArLy   SiteSearchSearchParamsLanguage = "ar-ly"
	SiteSearchSearchParamsLanguageArMa   SiteSearchSearchParamsLanguage = "ar-ma"
	SiteSearchSearchParamsLanguageArMr   SiteSearchSearchParamsLanguage = "ar-mr"
	SiteSearchSearchParamsLanguageArOm   SiteSearchSearchParamsLanguage = "ar-om"
	SiteSearchSearchParamsLanguageArPs   SiteSearchSearchParamsLanguage = "ar-ps"
	SiteSearchSearchParamsLanguageArQa   SiteSearchSearchParamsLanguage = "ar-qa"
	SiteSearchSearchParamsLanguageArSa   SiteSearchSearchParamsLanguage = "ar-sa"
	SiteSearchSearchParamsLanguageArSd   SiteSearchSearchParamsLanguage = "ar-sd"
	SiteSearchSearchParamsLanguageArSo   SiteSearchSearchParamsLanguage = "ar-so"
	SiteSearchSearchParamsLanguageArSS   SiteSearchSearchParamsLanguage = "ar-ss"
	SiteSearchSearchParamsLanguageArSy   SiteSearchSearchParamsLanguage = "ar-sy"
	SiteSearchSearchParamsLanguageArTd   SiteSearchSearchParamsLanguage = "ar-td"
	SiteSearchSearchParamsLanguageArTn   SiteSearchSearchParamsLanguage = "ar-tn"
	SiteSearchSearchParamsLanguageArYe   SiteSearchSearchParamsLanguage = "ar-ye"
	SiteSearchSearchParamsLanguageAs     SiteSearchSearchParamsLanguage = "as"
	SiteSearchSearchParamsLanguageAsIn   SiteSearchSearchParamsLanguage = "as-in"
	SiteSearchSearchParamsLanguageAsa    SiteSearchSearchParamsLanguage = "asa"
	SiteSearchSearchParamsLanguageAsaTz  SiteSearchSearchParamsLanguage = "asa-tz"
	SiteSearchSearchParamsLanguageAst    SiteSearchSearchParamsLanguage = "ast"
	SiteSearchSearchParamsLanguageAstEs  SiteSearchSearchParamsLanguage = "ast-es"
	SiteSearchSearchParamsLanguageAz     SiteSearchSearchParamsLanguage = "az"
	SiteSearchSearchParamsLanguageAzAz   SiteSearchSearchParamsLanguage = "az-az"
	SiteSearchSearchParamsLanguageBas    SiteSearchSearchParamsLanguage = "bas"
	SiteSearchSearchParamsLanguageBasCm  SiteSearchSearchParamsLanguage = "bas-cm"
	SiteSearchSearchParamsLanguageBe     SiteSearchSearchParamsLanguage = "be"
	SiteSearchSearchParamsLanguageBeBy   SiteSearchSearchParamsLanguage = "be-by"
	SiteSearchSearchParamsLanguageBem    SiteSearchSearchParamsLanguage = "bem"
	SiteSearchSearchParamsLanguageBemZm  SiteSearchSearchParamsLanguage = "bem-zm"
	SiteSearchSearchParamsLanguageBez    SiteSearchSearchParamsLanguage = "bez"
	SiteSearchSearchParamsLanguageBezTz  SiteSearchSearchParamsLanguage = "bez-tz"
	SiteSearchSearchParamsLanguageBg     SiteSearchSearchParamsLanguage = "bg"
	SiteSearchSearchParamsLanguageBgBg   SiteSearchSearchParamsLanguage = "bg-bg"
	SiteSearchSearchParamsLanguageBm     SiteSearchSearchParamsLanguage = "bm"
	SiteSearchSearchParamsLanguageBmMl   SiteSearchSearchParamsLanguage = "bm-ml"
	SiteSearchSearchParamsLanguageBn     SiteSearchSearchParamsLanguage = "bn"
	SiteSearchSearchParamsLanguageBnBd   SiteSearchSearchParamsLanguage = "bn-bd"
	SiteSearchSearchParamsLanguageBnIn   SiteSearchSearchParamsLanguage = "bn-in"
	SiteSearchSearchParamsLanguageBo     SiteSearchSearchParamsLanguage = "bo"
	SiteSearchSearchParamsLanguageBoCn   SiteSearchSearchParamsLanguage = "bo-cn"
	SiteSearchSearchParamsLanguageBoIn   SiteSearchSearchParamsLanguage = "bo-in"
	SiteSearchSearchParamsLanguageBr     SiteSearchSearchParamsLanguage = "br"
	SiteSearchSearchParamsLanguageBrFr   SiteSearchSearchParamsLanguage = "br-fr"
	SiteSearchSearchParamsLanguageBrx    SiteSearchSearchParamsLanguage = "brx"
	SiteSearchSearchParamsLanguageBrxIn  SiteSearchSearchParamsLanguage = "brx-in"
	SiteSearchSearchParamsLanguageBs     SiteSearchSearchParamsLanguage = "bs"
	SiteSearchSearchParamsLanguageBsBa   SiteSearchSearchParamsLanguage = "bs-ba"
	SiteSearchSearchParamsLanguageCa     SiteSearchSearchParamsLanguage = "ca"
	SiteSearchSearchParamsLanguageCaAd   SiteSearchSearchParamsLanguage = "ca-ad"
	SiteSearchSearchParamsLanguageCaEs   SiteSearchSearchParamsLanguage = "ca-es"
	SiteSearchSearchParamsLanguageCaFr   SiteSearchSearchParamsLanguage = "ca-fr"
	SiteSearchSearchParamsLanguageCaIt   SiteSearchSearchParamsLanguage = "ca-it"
	SiteSearchSearchParamsLanguageCcp    SiteSearchSearchParamsLanguage = "ccp"
	SiteSearchSearchParamsLanguageCcpBd  SiteSearchSearchParamsLanguage = "ccp-bd"
	SiteSearchSearchParamsLanguageCcpIn  SiteSearchSearchParamsLanguage = "ccp-in"
	SiteSearchSearchParamsLanguageCe     SiteSearchSearchParamsLanguage = "ce"
	SiteSearchSearchParamsLanguageCeRu   SiteSearchSearchParamsLanguage = "ce-ru"
	SiteSearchSearchParamsLanguageCeb    SiteSearchSearchParamsLanguage = "ceb"
	SiteSearchSearchParamsLanguageCebPh  SiteSearchSearchParamsLanguage = "ceb-ph"
	SiteSearchSearchParamsLanguageCgg    SiteSearchSearchParamsLanguage = "cgg"
	SiteSearchSearchParamsLanguageCggUg  SiteSearchSearchParamsLanguage = "cgg-ug"
	SiteSearchSearchParamsLanguageChr    SiteSearchSearchParamsLanguage = "chr"
	SiteSearchSearchParamsLanguageChrUs  SiteSearchSearchParamsLanguage = "chr-us"
	SiteSearchSearchParamsLanguageCkb    SiteSearchSearchParamsLanguage = "ckb"
	SiteSearchSearchParamsLanguageCkbIq  SiteSearchSearchParamsLanguage = "ckb-iq"
	SiteSearchSearchParamsLanguageCkbIr  SiteSearchSearchParamsLanguage = "ckb-ir"
	SiteSearchSearchParamsLanguageCs     SiteSearchSearchParamsLanguage = "cs"
	SiteSearchSearchParamsLanguageCsCz   SiteSearchSearchParamsLanguage = "cs-cz"
	SiteSearchSearchParamsLanguageCu     SiteSearchSearchParamsLanguage = "cu"
	SiteSearchSearchParamsLanguageCuRu   SiteSearchSearchParamsLanguage = "cu-ru"
	SiteSearchSearchParamsLanguageCy     SiteSearchSearchParamsLanguage = "cy"
	SiteSearchSearchParamsLanguageCyGB   SiteSearchSearchParamsLanguage = "cy-gb"
	SiteSearchSearchParamsLanguageDa     SiteSearchSearchParamsLanguage = "da"
	SiteSearchSearchParamsLanguageDaDk   SiteSearchSearchParamsLanguage = "da-dk"
	SiteSearchSearchParamsLanguageDaGl   SiteSearchSearchParamsLanguage = "da-gl"
	SiteSearchSearchParamsLanguageDav    SiteSearchSearchParamsLanguage = "dav"
	SiteSearchSearchParamsLanguageDavKe  SiteSearchSearchParamsLanguage = "dav-ke"
	SiteSearchSearchParamsLanguageDe     SiteSearchSearchParamsLanguage = "de"
	SiteSearchSearchParamsLanguageDeAt   SiteSearchSearchParamsLanguage = "de-at"
	SiteSearchSearchParamsLanguageDeBe   SiteSearchSearchParamsLanguage = "de-be"
	SiteSearchSearchParamsLanguageDeCh   SiteSearchSearchParamsLanguage = "de-ch"
	SiteSearchSearchParamsLanguageDeDe   SiteSearchSearchParamsLanguage = "de-de"
	SiteSearchSearchParamsLanguageDeGr   SiteSearchSearchParamsLanguage = "de-gr"
	SiteSearchSearchParamsLanguageDeIt   SiteSearchSearchParamsLanguage = "de-it"
	SiteSearchSearchParamsLanguageDeLi   SiteSearchSearchParamsLanguage = "de-li"
	SiteSearchSearchParamsLanguageDeLu   SiteSearchSearchParamsLanguage = "de-lu"
	SiteSearchSearchParamsLanguageDje    SiteSearchSearchParamsLanguage = "dje"
	SiteSearchSearchParamsLanguageDjeNe  SiteSearchSearchParamsLanguage = "dje-ne"
	SiteSearchSearchParamsLanguageDoi    SiteSearchSearchParamsLanguage = "doi"
	SiteSearchSearchParamsLanguageDoiIn  SiteSearchSearchParamsLanguage = "doi-in"
	SiteSearchSearchParamsLanguageDsb    SiteSearchSearchParamsLanguage = "dsb"
	SiteSearchSearchParamsLanguageDsbDe  SiteSearchSearchParamsLanguage = "dsb-de"
	SiteSearchSearchParamsLanguageDua    SiteSearchSearchParamsLanguage = "dua"
	SiteSearchSearchParamsLanguageDuaCm  SiteSearchSearchParamsLanguage = "dua-cm"
	SiteSearchSearchParamsLanguageDyo    SiteSearchSearchParamsLanguage = "dyo"
	SiteSearchSearchParamsLanguageDyoSn  SiteSearchSearchParamsLanguage = "dyo-sn"
	SiteSearchSearchParamsLanguageDz     SiteSearchSearchParamsLanguage = "dz"
	SiteSearchSearchParamsLanguageDzBt   SiteSearchSearchParamsLanguage = "dz-bt"
	SiteSearchSearchParamsLanguageEbu    SiteSearchSearchParamsLanguage = "ebu"
	SiteSearchSearchParamsLanguageEbuKe  SiteSearchSearchParamsLanguage = "ebu-ke"
	SiteSearchSearchParamsLanguageEe     SiteSearchSearchParamsLanguage = "ee"
	SiteSearchSearchParamsLanguageEeGh   SiteSearchSearchParamsLanguage = "ee-gh"
	SiteSearchSearchParamsLanguageEeTg   SiteSearchSearchParamsLanguage = "ee-tg"
	SiteSearchSearchParamsLanguageEl     SiteSearchSearchParamsLanguage = "el"
	SiteSearchSearchParamsLanguageElCy   SiteSearchSearchParamsLanguage = "el-cy"
	SiteSearchSearchParamsLanguageElGr   SiteSearchSearchParamsLanguage = "el-gr"
	SiteSearchSearchParamsLanguageEn     SiteSearchSearchParamsLanguage = "en"
	SiteSearchSearchParamsLanguageEn001  SiteSearchSearchParamsLanguage = "en-001"
	SiteSearchSearchParamsLanguageEn150  SiteSearchSearchParamsLanguage = "en-150"
	SiteSearchSearchParamsLanguageEnAe   SiteSearchSearchParamsLanguage = "en-ae"
	SiteSearchSearchParamsLanguageEnAg   SiteSearchSearchParamsLanguage = "en-ag"
	SiteSearchSearchParamsLanguageEnAI   SiteSearchSearchParamsLanguage = "en-ai"
	SiteSearchSearchParamsLanguageEnAs   SiteSearchSearchParamsLanguage = "en-as"
	SiteSearchSearchParamsLanguageEnAt   SiteSearchSearchParamsLanguage = "en-at"
	SiteSearchSearchParamsLanguageEnAu   SiteSearchSearchParamsLanguage = "en-au"
	SiteSearchSearchParamsLanguageEnBb   SiteSearchSearchParamsLanguage = "en-bb"
	SiteSearchSearchParamsLanguageEnBe   SiteSearchSearchParamsLanguage = "en-be"
	SiteSearchSearchParamsLanguageEnBi   SiteSearchSearchParamsLanguage = "en-bi"
	SiteSearchSearchParamsLanguageEnBm   SiteSearchSearchParamsLanguage = "en-bm"
	SiteSearchSearchParamsLanguageEnBs   SiteSearchSearchParamsLanguage = "en-bs"
	SiteSearchSearchParamsLanguageEnBw   SiteSearchSearchParamsLanguage = "en-bw"
	SiteSearchSearchParamsLanguageEnBz   SiteSearchSearchParamsLanguage = "en-bz"
	SiteSearchSearchParamsLanguageEnCa   SiteSearchSearchParamsLanguage = "en-ca"
	SiteSearchSearchParamsLanguageEnCc   SiteSearchSearchParamsLanguage = "en-cc"
	SiteSearchSearchParamsLanguageEnCh   SiteSearchSearchParamsLanguage = "en-ch"
	SiteSearchSearchParamsLanguageEnCk   SiteSearchSearchParamsLanguage = "en-ck"
	SiteSearchSearchParamsLanguageEnCm   SiteSearchSearchParamsLanguage = "en-cm"
	SiteSearchSearchParamsLanguageEnCn   SiteSearchSearchParamsLanguage = "en-cn"
	SiteSearchSearchParamsLanguageEnCx   SiteSearchSearchParamsLanguage = "en-cx"
	SiteSearchSearchParamsLanguageEnCy   SiteSearchSearchParamsLanguage = "en-cy"
	SiteSearchSearchParamsLanguageEnDe   SiteSearchSearchParamsLanguage = "en-de"
	SiteSearchSearchParamsLanguageEnDg   SiteSearchSearchParamsLanguage = "en-dg"
	SiteSearchSearchParamsLanguageEnDk   SiteSearchSearchParamsLanguage = "en-dk"
	SiteSearchSearchParamsLanguageEnDm   SiteSearchSearchParamsLanguage = "en-dm"
	SiteSearchSearchParamsLanguageEnEr   SiteSearchSearchParamsLanguage = "en-er"
	SiteSearchSearchParamsLanguageEnFi   SiteSearchSearchParamsLanguage = "en-fi"
	SiteSearchSearchParamsLanguageEnFj   SiteSearchSearchParamsLanguage = "en-fj"
	SiteSearchSearchParamsLanguageEnFk   SiteSearchSearchParamsLanguage = "en-fk"
	SiteSearchSearchParamsLanguageEnFm   SiteSearchSearchParamsLanguage = "en-fm"
	SiteSearchSearchParamsLanguageEnGB   SiteSearchSearchParamsLanguage = "en-gb"
	SiteSearchSearchParamsLanguageEnGd   SiteSearchSearchParamsLanguage = "en-gd"
	SiteSearchSearchParamsLanguageEnGg   SiteSearchSearchParamsLanguage = "en-gg"
	SiteSearchSearchParamsLanguageEnGh   SiteSearchSearchParamsLanguage = "en-gh"
	SiteSearchSearchParamsLanguageEnGi   SiteSearchSearchParamsLanguage = "en-gi"
	SiteSearchSearchParamsLanguageEnGm   SiteSearchSearchParamsLanguage = "en-gm"
	SiteSearchSearchParamsLanguageEnGu   SiteSearchSearchParamsLanguage = "en-gu"
	SiteSearchSearchParamsLanguageEnGy   SiteSearchSearchParamsLanguage = "en-gy"
	SiteSearchSearchParamsLanguageEnHk   SiteSearchSearchParamsLanguage = "en-hk"
	SiteSearchSearchParamsLanguageEnIe   SiteSearchSearchParamsLanguage = "en-ie"
	SiteSearchSearchParamsLanguageEnIl   SiteSearchSearchParamsLanguage = "en-il"
	SiteSearchSearchParamsLanguageEnIm   SiteSearchSearchParamsLanguage = "en-im"
	SiteSearchSearchParamsLanguageEnIn   SiteSearchSearchParamsLanguage = "en-in"
	SiteSearchSearchParamsLanguageEnIo   SiteSearchSearchParamsLanguage = "en-io"
	SiteSearchSearchParamsLanguageEnJe   SiteSearchSearchParamsLanguage = "en-je"
	SiteSearchSearchParamsLanguageEnJm   SiteSearchSearchParamsLanguage = "en-jm"
	SiteSearchSearchParamsLanguageEnKe   SiteSearchSearchParamsLanguage = "en-ke"
	SiteSearchSearchParamsLanguageEnKi   SiteSearchSearchParamsLanguage = "en-ki"
	SiteSearchSearchParamsLanguageEnKn   SiteSearchSearchParamsLanguage = "en-kn"
	SiteSearchSearchParamsLanguageEnKy   SiteSearchSearchParamsLanguage = "en-ky"
	SiteSearchSearchParamsLanguageEnLc   SiteSearchSearchParamsLanguage = "en-lc"
	SiteSearchSearchParamsLanguageEnLr   SiteSearchSearchParamsLanguage = "en-lr"
	SiteSearchSearchParamsLanguageEnLs   SiteSearchSearchParamsLanguage = "en-ls"
	SiteSearchSearchParamsLanguageEnLu   SiteSearchSearchParamsLanguage = "en-lu"
	SiteSearchSearchParamsLanguageEnMg   SiteSearchSearchParamsLanguage = "en-mg"
	SiteSearchSearchParamsLanguageEnMh   SiteSearchSearchParamsLanguage = "en-mh"
	SiteSearchSearchParamsLanguageEnMo   SiteSearchSearchParamsLanguage = "en-mo"
	SiteSearchSearchParamsLanguageEnMp   SiteSearchSearchParamsLanguage = "en-mp"
	SiteSearchSearchParamsLanguageEnMs   SiteSearchSearchParamsLanguage = "en-ms"
	SiteSearchSearchParamsLanguageEnMt   SiteSearchSearchParamsLanguage = "en-mt"
	SiteSearchSearchParamsLanguageEnMu   SiteSearchSearchParamsLanguage = "en-mu"
	SiteSearchSearchParamsLanguageEnMw   SiteSearchSearchParamsLanguage = "en-mw"
	SiteSearchSearchParamsLanguageEnMx   SiteSearchSearchParamsLanguage = "en-mx"
	SiteSearchSearchParamsLanguageEnMy   SiteSearchSearchParamsLanguage = "en-my"
	SiteSearchSearchParamsLanguageEnNa   SiteSearchSearchParamsLanguage = "en-na"
	SiteSearchSearchParamsLanguageEnNf   SiteSearchSearchParamsLanguage = "en-nf"
	SiteSearchSearchParamsLanguageEnNg   SiteSearchSearchParamsLanguage = "en-ng"
	SiteSearchSearchParamsLanguageEnNl   SiteSearchSearchParamsLanguage = "en-nl"
	SiteSearchSearchParamsLanguageEnNr   SiteSearchSearchParamsLanguage = "en-nr"
	SiteSearchSearchParamsLanguageEnNu   SiteSearchSearchParamsLanguage = "en-nu"
	SiteSearchSearchParamsLanguageEnNz   SiteSearchSearchParamsLanguage = "en-nz"
	SiteSearchSearchParamsLanguageEnPg   SiteSearchSearchParamsLanguage = "en-pg"
	SiteSearchSearchParamsLanguageEnPh   SiteSearchSearchParamsLanguage = "en-ph"
	SiteSearchSearchParamsLanguageEnPk   SiteSearchSearchParamsLanguage = "en-pk"
	SiteSearchSearchParamsLanguageEnPn   SiteSearchSearchParamsLanguage = "en-pn"
	SiteSearchSearchParamsLanguageEnPr   SiteSearchSearchParamsLanguage = "en-pr"
	SiteSearchSearchParamsLanguageEnPw   SiteSearchSearchParamsLanguage = "en-pw"
	SiteSearchSearchParamsLanguageEnRw   SiteSearchSearchParamsLanguage = "en-rw"
	SiteSearchSearchParamsLanguageEnSb   SiteSearchSearchParamsLanguage = "en-sb"
	SiteSearchSearchParamsLanguageEnSc   SiteSearchSearchParamsLanguage = "en-sc"
	SiteSearchSearchParamsLanguageEnSd   SiteSearchSearchParamsLanguage = "en-sd"
	SiteSearchSearchParamsLanguageEnSe   SiteSearchSearchParamsLanguage = "en-se"
	SiteSearchSearchParamsLanguageEnSg   SiteSearchSearchParamsLanguage = "en-sg"
	SiteSearchSearchParamsLanguageEnSh   SiteSearchSearchParamsLanguage = "en-sh"
	SiteSearchSearchParamsLanguageEnSi   SiteSearchSearchParamsLanguage = "en-si"
	SiteSearchSearchParamsLanguageEnSl   SiteSearchSearchParamsLanguage = "en-sl"
	SiteSearchSearchParamsLanguageEnSS   SiteSearchSearchParamsLanguage = "en-ss"
	SiteSearchSearchParamsLanguageEnSx   SiteSearchSearchParamsLanguage = "en-sx"
	SiteSearchSearchParamsLanguageEnSz   SiteSearchSearchParamsLanguage = "en-sz"
	SiteSearchSearchParamsLanguageEnTc   SiteSearchSearchParamsLanguage = "en-tc"
	SiteSearchSearchParamsLanguageEnTk   SiteSearchSearchParamsLanguage = "en-tk"
	SiteSearchSearchParamsLanguageEnTo   SiteSearchSearchParamsLanguage = "en-to"
	SiteSearchSearchParamsLanguageEnTt   SiteSearchSearchParamsLanguage = "en-tt"
	SiteSearchSearchParamsLanguageEnTv   SiteSearchSearchParamsLanguage = "en-tv"
	SiteSearchSearchParamsLanguageEnTz   SiteSearchSearchParamsLanguage = "en-tz"
	SiteSearchSearchParamsLanguageEnUg   SiteSearchSearchParamsLanguage = "en-ug"
	SiteSearchSearchParamsLanguageEnUm   SiteSearchSearchParamsLanguage = "en-um"
	SiteSearchSearchParamsLanguageEnUs   SiteSearchSearchParamsLanguage = "en-us"
	SiteSearchSearchParamsLanguageEnVc   SiteSearchSearchParamsLanguage = "en-vc"
	SiteSearchSearchParamsLanguageEnVg   SiteSearchSearchParamsLanguage = "en-vg"
	SiteSearchSearchParamsLanguageEnVi   SiteSearchSearchParamsLanguage = "en-vi"
	SiteSearchSearchParamsLanguageEnVu   SiteSearchSearchParamsLanguage = "en-vu"
	SiteSearchSearchParamsLanguageEnWs   SiteSearchSearchParamsLanguage = "en-ws"
	SiteSearchSearchParamsLanguageEnZa   SiteSearchSearchParamsLanguage = "en-za"
	SiteSearchSearchParamsLanguageEnZm   SiteSearchSearchParamsLanguage = "en-zm"
	SiteSearchSearchParamsLanguageEnZw   SiteSearchSearchParamsLanguage = "en-zw"
	SiteSearchSearchParamsLanguageEo     SiteSearchSearchParamsLanguage = "eo"
	SiteSearchSearchParamsLanguageEo001  SiteSearchSearchParamsLanguage = "eo-001"
	SiteSearchSearchParamsLanguageEs     SiteSearchSearchParamsLanguage = "es"
	SiteSearchSearchParamsLanguageEs419  SiteSearchSearchParamsLanguage = "es-419"
	SiteSearchSearchParamsLanguageEsAr   SiteSearchSearchParamsLanguage = "es-ar"
	SiteSearchSearchParamsLanguageEsBo   SiteSearchSearchParamsLanguage = "es-bo"
	SiteSearchSearchParamsLanguageEsBr   SiteSearchSearchParamsLanguage = "es-br"
	SiteSearchSearchParamsLanguageEsBz   SiteSearchSearchParamsLanguage = "es-bz"
	SiteSearchSearchParamsLanguageEsCl   SiteSearchSearchParamsLanguage = "es-cl"
	SiteSearchSearchParamsLanguageEsCo   SiteSearchSearchParamsLanguage = "es-co"
	SiteSearchSearchParamsLanguageEsCr   SiteSearchSearchParamsLanguage = "es-cr"
	SiteSearchSearchParamsLanguageEsCu   SiteSearchSearchParamsLanguage = "es-cu"
	SiteSearchSearchParamsLanguageEsDo   SiteSearchSearchParamsLanguage = "es-do"
	SiteSearchSearchParamsLanguageEsEa   SiteSearchSearchParamsLanguage = "es-ea"
	SiteSearchSearchParamsLanguageEsEc   SiteSearchSearchParamsLanguage = "es-ec"
	SiteSearchSearchParamsLanguageEsEs   SiteSearchSearchParamsLanguage = "es-es"
	SiteSearchSearchParamsLanguageEsGq   SiteSearchSearchParamsLanguage = "es-gq"
	SiteSearchSearchParamsLanguageEsGt   SiteSearchSearchParamsLanguage = "es-gt"
	SiteSearchSearchParamsLanguageEsHn   SiteSearchSearchParamsLanguage = "es-hn"
	SiteSearchSearchParamsLanguageEsIc   SiteSearchSearchParamsLanguage = "es-ic"
	SiteSearchSearchParamsLanguageEsMx   SiteSearchSearchParamsLanguage = "es-mx"
	SiteSearchSearchParamsLanguageEsNi   SiteSearchSearchParamsLanguage = "es-ni"
	SiteSearchSearchParamsLanguageEsPa   SiteSearchSearchParamsLanguage = "es-pa"
	SiteSearchSearchParamsLanguageEsPe   SiteSearchSearchParamsLanguage = "es-pe"
	SiteSearchSearchParamsLanguageEsPh   SiteSearchSearchParamsLanguage = "es-ph"
	SiteSearchSearchParamsLanguageEsPr   SiteSearchSearchParamsLanguage = "es-pr"
	SiteSearchSearchParamsLanguageEsPy   SiteSearchSearchParamsLanguage = "es-py"
	SiteSearchSearchParamsLanguageEsSv   SiteSearchSearchParamsLanguage = "es-sv"
	SiteSearchSearchParamsLanguageEsUs   SiteSearchSearchParamsLanguage = "es-us"
	SiteSearchSearchParamsLanguageEsUy   SiteSearchSearchParamsLanguage = "es-uy"
	SiteSearchSearchParamsLanguageEsVe   SiteSearchSearchParamsLanguage = "es-ve"
	SiteSearchSearchParamsLanguageEt     SiteSearchSearchParamsLanguage = "et"
	SiteSearchSearchParamsLanguageEtEe   SiteSearchSearchParamsLanguage = "et-ee"
	SiteSearchSearchParamsLanguageEu     SiteSearchSearchParamsLanguage = "eu"
	SiteSearchSearchParamsLanguageEuEs   SiteSearchSearchParamsLanguage = "eu-es"
	SiteSearchSearchParamsLanguageEwo    SiteSearchSearchParamsLanguage = "ewo"
	SiteSearchSearchParamsLanguageEwoCm  SiteSearchSearchParamsLanguage = "ewo-cm"
	SiteSearchSearchParamsLanguageFa     SiteSearchSearchParamsLanguage = "fa"
	SiteSearchSearchParamsLanguageFaAf   SiteSearchSearchParamsLanguage = "fa-af"
	SiteSearchSearchParamsLanguageFaIr   SiteSearchSearchParamsLanguage = "fa-ir"
	SiteSearchSearchParamsLanguageFf     SiteSearchSearchParamsLanguage = "ff"
	SiteSearchSearchParamsLanguageFfBf   SiteSearchSearchParamsLanguage = "ff-bf"
	SiteSearchSearchParamsLanguageFfCm   SiteSearchSearchParamsLanguage = "ff-cm"
	SiteSearchSearchParamsLanguageFfGh   SiteSearchSearchParamsLanguage = "ff-gh"
	SiteSearchSearchParamsLanguageFfGm   SiteSearchSearchParamsLanguage = "ff-gm"
	SiteSearchSearchParamsLanguageFfGn   SiteSearchSearchParamsLanguage = "ff-gn"
	SiteSearchSearchParamsLanguageFfGw   SiteSearchSearchParamsLanguage = "ff-gw"
	SiteSearchSearchParamsLanguageFfLr   SiteSearchSearchParamsLanguage = "ff-lr"
	SiteSearchSearchParamsLanguageFfMr   SiteSearchSearchParamsLanguage = "ff-mr"
	SiteSearchSearchParamsLanguageFfNe   SiteSearchSearchParamsLanguage = "ff-ne"
	SiteSearchSearchParamsLanguageFfNg   SiteSearchSearchParamsLanguage = "ff-ng"
	SiteSearchSearchParamsLanguageFfSl   SiteSearchSearchParamsLanguage = "ff-sl"
	SiteSearchSearchParamsLanguageFfSn   SiteSearchSearchParamsLanguage = "ff-sn"
	SiteSearchSearchParamsLanguageFi     SiteSearchSearchParamsLanguage = "fi"
	SiteSearchSearchParamsLanguageFiFi   SiteSearchSearchParamsLanguage = "fi-fi"
	SiteSearchSearchParamsLanguageFil    SiteSearchSearchParamsLanguage = "fil"
	SiteSearchSearchParamsLanguageFilPh  SiteSearchSearchParamsLanguage = "fil-ph"
	SiteSearchSearchParamsLanguageFo     SiteSearchSearchParamsLanguage = "fo"
	SiteSearchSearchParamsLanguageFoDk   SiteSearchSearchParamsLanguage = "fo-dk"
	SiteSearchSearchParamsLanguageFoFo   SiteSearchSearchParamsLanguage = "fo-fo"
	SiteSearchSearchParamsLanguageFr     SiteSearchSearchParamsLanguage = "fr"
	SiteSearchSearchParamsLanguageFrBe   SiteSearchSearchParamsLanguage = "fr-be"
	SiteSearchSearchParamsLanguageFrBf   SiteSearchSearchParamsLanguage = "fr-bf"
	SiteSearchSearchParamsLanguageFrBi   SiteSearchSearchParamsLanguage = "fr-bi"
	SiteSearchSearchParamsLanguageFrBj   SiteSearchSearchParamsLanguage = "fr-bj"
	SiteSearchSearchParamsLanguageFrBl   SiteSearchSearchParamsLanguage = "fr-bl"
	SiteSearchSearchParamsLanguageFrCa   SiteSearchSearchParamsLanguage = "fr-ca"
	SiteSearchSearchParamsLanguageFrCd   SiteSearchSearchParamsLanguage = "fr-cd"
	SiteSearchSearchParamsLanguageFrCf   SiteSearchSearchParamsLanguage = "fr-cf"
	SiteSearchSearchParamsLanguageFrCg   SiteSearchSearchParamsLanguage = "fr-cg"
	SiteSearchSearchParamsLanguageFrCh   SiteSearchSearchParamsLanguage = "fr-ch"
	SiteSearchSearchParamsLanguageFrCi   SiteSearchSearchParamsLanguage = "fr-ci"
	SiteSearchSearchParamsLanguageFrCm   SiteSearchSearchParamsLanguage = "fr-cm"
	SiteSearchSearchParamsLanguageFrDj   SiteSearchSearchParamsLanguage = "fr-dj"
	SiteSearchSearchParamsLanguageFrDz   SiteSearchSearchParamsLanguage = "fr-dz"
	SiteSearchSearchParamsLanguageFrFr   SiteSearchSearchParamsLanguage = "fr-fr"
	SiteSearchSearchParamsLanguageFrGa   SiteSearchSearchParamsLanguage = "fr-ga"
	SiteSearchSearchParamsLanguageFrGf   SiteSearchSearchParamsLanguage = "fr-gf"
	SiteSearchSearchParamsLanguageFrGn   SiteSearchSearchParamsLanguage = "fr-gn"
	SiteSearchSearchParamsLanguageFrGp   SiteSearchSearchParamsLanguage = "fr-gp"
	SiteSearchSearchParamsLanguageFrGq   SiteSearchSearchParamsLanguage = "fr-gq"
	SiteSearchSearchParamsLanguageFrHt   SiteSearchSearchParamsLanguage = "fr-ht"
	SiteSearchSearchParamsLanguageFrKm   SiteSearchSearchParamsLanguage = "fr-km"
	SiteSearchSearchParamsLanguageFrLu   SiteSearchSearchParamsLanguage = "fr-lu"
	SiteSearchSearchParamsLanguageFrMa   SiteSearchSearchParamsLanguage = "fr-ma"
	SiteSearchSearchParamsLanguageFrMc   SiteSearchSearchParamsLanguage = "fr-mc"
	SiteSearchSearchParamsLanguageFrMf   SiteSearchSearchParamsLanguage = "fr-mf"
	SiteSearchSearchParamsLanguageFrMg   SiteSearchSearchParamsLanguage = "fr-mg"
	SiteSearchSearchParamsLanguageFrMl   SiteSearchSearchParamsLanguage = "fr-ml"
	SiteSearchSearchParamsLanguageFrMq   SiteSearchSearchParamsLanguage = "fr-mq"
	SiteSearchSearchParamsLanguageFrMr   SiteSearchSearchParamsLanguage = "fr-mr"
	SiteSearchSearchParamsLanguageFrMu   SiteSearchSearchParamsLanguage = "fr-mu"
	SiteSearchSearchParamsLanguageFrNc   SiteSearchSearchParamsLanguage = "fr-nc"
	SiteSearchSearchParamsLanguageFrNe   SiteSearchSearchParamsLanguage = "fr-ne"
	SiteSearchSearchParamsLanguageFrPf   SiteSearchSearchParamsLanguage = "fr-pf"
	SiteSearchSearchParamsLanguageFrPm   SiteSearchSearchParamsLanguage = "fr-pm"
	SiteSearchSearchParamsLanguageFrRe   SiteSearchSearchParamsLanguage = "fr-re"
	SiteSearchSearchParamsLanguageFrRw   SiteSearchSearchParamsLanguage = "fr-rw"
	SiteSearchSearchParamsLanguageFrSc   SiteSearchSearchParamsLanguage = "fr-sc"
	SiteSearchSearchParamsLanguageFrSn   SiteSearchSearchParamsLanguage = "fr-sn"
	SiteSearchSearchParamsLanguageFrSy   SiteSearchSearchParamsLanguage = "fr-sy"
	SiteSearchSearchParamsLanguageFrTd   SiteSearchSearchParamsLanguage = "fr-td"
	SiteSearchSearchParamsLanguageFrTg   SiteSearchSearchParamsLanguage = "fr-tg"
	SiteSearchSearchParamsLanguageFrTn   SiteSearchSearchParamsLanguage = "fr-tn"
	SiteSearchSearchParamsLanguageFrVu   SiteSearchSearchParamsLanguage = "fr-vu"
	SiteSearchSearchParamsLanguageFrWf   SiteSearchSearchParamsLanguage = "fr-wf"
	SiteSearchSearchParamsLanguageFrYt   SiteSearchSearchParamsLanguage = "fr-yt"
	SiteSearchSearchParamsLanguageFur    SiteSearchSearchParamsLanguage = "fur"
	SiteSearchSearchParamsLanguageFurIt  SiteSearchSearchParamsLanguage = "fur-it"
	SiteSearchSearchParamsLanguageFy     SiteSearchSearchParamsLanguage = "fy"
	SiteSearchSearchParamsLanguageFyNl   SiteSearchSearchParamsLanguage = "fy-nl"
	SiteSearchSearchParamsLanguageGa     SiteSearchSearchParamsLanguage = "ga"
	SiteSearchSearchParamsLanguageGaGB   SiteSearchSearchParamsLanguage = "ga-gb"
	SiteSearchSearchParamsLanguageGaIe   SiteSearchSearchParamsLanguage = "ga-ie"
	SiteSearchSearchParamsLanguageGd     SiteSearchSearchParamsLanguage = "gd"
	SiteSearchSearchParamsLanguageGdGB   SiteSearchSearchParamsLanguage = "gd-gb"
	SiteSearchSearchParamsLanguageGl     SiteSearchSearchParamsLanguage = "gl"
	SiteSearchSearchParamsLanguageGlEs   SiteSearchSearchParamsLanguage = "gl-es"
	SiteSearchSearchParamsLanguageGsw    SiteSearchSearchParamsLanguage = "gsw"
	SiteSearchSearchParamsLanguageGswCh  SiteSearchSearchParamsLanguage = "gsw-ch"
	SiteSearchSearchParamsLanguageGswFr  SiteSearchSearchParamsLanguage = "gsw-fr"
	SiteSearchSearchParamsLanguageGswLi  SiteSearchSearchParamsLanguage = "gsw-li"
	SiteSearchSearchParamsLanguageGu     SiteSearchSearchParamsLanguage = "gu"
	SiteSearchSearchParamsLanguageGuIn   SiteSearchSearchParamsLanguage = "gu-in"
	SiteSearchSearchParamsLanguageGuz    SiteSearchSearchParamsLanguage = "guz"
	SiteSearchSearchParamsLanguageGuzKe  SiteSearchSearchParamsLanguage = "guz-ke"
	SiteSearchSearchParamsLanguageGv     SiteSearchSearchParamsLanguage = "gv"
	SiteSearchSearchParamsLanguageGvIm   SiteSearchSearchParamsLanguage = "gv-im"
	SiteSearchSearchParamsLanguageHa     SiteSearchSearchParamsLanguage = "ha"
	SiteSearchSearchParamsLanguageHaGh   SiteSearchSearchParamsLanguage = "ha-gh"
	SiteSearchSearchParamsLanguageHaNe   SiteSearchSearchParamsLanguage = "ha-ne"
	SiteSearchSearchParamsLanguageHaNg   SiteSearchSearchParamsLanguage = "ha-ng"
	SiteSearchSearchParamsLanguageHaw    SiteSearchSearchParamsLanguage = "haw"
	SiteSearchSearchParamsLanguageHawUs  SiteSearchSearchParamsLanguage = "haw-us"
	SiteSearchSearchParamsLanguageHe     SiteSearchSearchParamsLanguage = "he"
	SiteSearchSearchParamsLanguageHeIl   SiteSearchSearchParamsLanguage = "he-il"
	SiteSearchSearchParamsLanguageHi     SiteSearchSearchParamsLanguage = "hi"
	SiteSearchSearchParamsLanguageHiIn   SiteSearchSearchParamsLanguage = "hi-in"
	SiteSearchSearchParamsLanguageHr     SiteSearchSearchParamsLanguage = "hr"
	SiteSearchSearchParamsLanguageHrBa   SiteSearchSearchParamsLanguage = "hr-ba"
	SiteSearchSearchParamsLanguageHrHr   SiteSearchSearchParamsLanguage = "hr-hr"
	SiteSearchSearchParamsLanguageHsb    SiteSearchSearchParamsLanguage = "hsb"
	SiteSearchSearchParamsLanguageHsbDe  SiteSearchSearchParamsLanguage = "hsb-de"
	SiteSearchSearchParamsLanguageHu     SiteSearchSearchParamsLanguage = "hu"
	SiteSearchSearchParamsLanguageHuHu   SiteSearchSearchParamsLanguage = "hu-hu"
	SiteSearchSearchParamsLanguageHy     SiteSearchSearchParamsLanguage = "hy"
	SiteSearchSearchParamsLanguageHyAm   SiteSearchSearchParamsLanguage = "hy-am"
	SiteSearchSearchParamsLanguageIa     SiteSearchSearchParamsLanguage = "ia"
	SiteSearchSearchParamsLanguageIa001  SiteSearchSearchParamsLanguage = "ia-001"
	SiteSearchSearchParamsLanguageID     SiteSearchSearchParamsLanguage = "id"
	SiteSearchSearchParamsLanguageIDID   SiteSearchSearchParamsLanguage = "id-id"
	SiteSearchSearchParamsLanguageIg     SiteSearchSearchParamsLanguage = "ig"
	SiteSearchSearchParamsLanguageIgNg   SiteSearchSearchParamsLanguage = "ig-ng"
	SiteSearchSearchParamsLanguageIi     SiteSearchSearchParamsLanguage = "ii"
	SiteSearchSearchParamsLanguageIiCn   SiteSearchSearchParamsLanguage = "ii-cn"
	SiteSearchSearchParamsLanguageIs     SiteSearchSearchParamsLanguage = "is"
	SiteSearchSearchParamsLanguageIsIs   SiteSearchSearchParamsLanguage = "is-is"
	SiteSearchSearchParamsLanguageIt     SiteSearchSearchParamsLanguage = "it"
	SiteSearchSearchParamsLanguageItCh   SiteSearchSearchParamsLanguage = "it-ch"
	SiteSearchSearchParamsLanguageItIt   SiteSearchSearchParamsLanguage = "it-it"
	SiteSearchSearchParamsLanguageItSm   SiteSearchSearchParamsLanguage = "it-sm"
	SiteSearchSearchParamsLanguageItVa   SiteSearchSearchParamsLanguage = "it-va"
	SiteSearchSearchParamsLanguageJa     SiteSearchSearchParamsLanguage = "ja"
	SiteSearchSearchParamsLanguageJaJp   SiteSearchSearchParamsLanguage = "ja-jp"
	SiteSearchSearchParamsLanguageJgo    SiteSearchSearchParamsLanguage = "jgo"
	SiteSearchSearchParamsLanguageJgoCm  SiteSearchSearchParamsLanguage = "jgo-cm"
	SiteSearchSearchParamsLanguageJmc    SiteSearchSearchParamsLanguage = "jmc"
	SiteSearchSearchParamsLanguageJmcTz  SiteSearchSearchParamsLanguage = "jmc-tz"
	SiteSearchSearchParamsLanguageJv     SiteSearchSearchParamsLanguage = "jv"
	SiteSearchSearchParamsLanguageJvID   SiteSearchSearchParamsLanguage = "jv-id"
	SiteSearchSearchParamsLanguageKa     SiteSearchSearchParamsLanguage = "ka"
	SiteSearchSearchParamsLanguageKaGe   SiteSearchSearchParamsLanguage = "ka-ge"
	SiteSearchSearchParamsLanguageKab    SiteSearchSearchParamsLanguage = "kab"
	SiteSearchSearchParamsLanguageKabDz  SiteSearchSearchParamsLanguage = "kab-dz"
	SiteSearchSearchParamsLanguageKam    SiteSearchSearchParamsLanguage = "kam"
	SiteSearchSearchParamsLanguageKamKe  SiteSearchSearchParamsLanguage = "kam-ke"
	SiteSearchSearchParamsLanguageKde    SiteSearchSearchParamsLanguage = "kde"
	SiteSearchSearchParamsLanguageKdeTz  SiteSearchSearchParamsLanguage = "kde-tz"
	SiteSearchSearchParamsLanguageKea    SiteSearchSearchParamsLanguage = "kea"
	SiteSearchSearchParamsLanguageKeaCv  SiteSearchSearchParamsLanguage = "kea-cv"
	SiteSearchSearchParamsLanguageKhq    SiteSearchSearchParamsLanguage = "khq"
	SiteSearchSearchParamsLanguageKhqMl  SiteSearchSearchParamsLanguage = "khq-ml"
	SiteSearchSearchParamsLanguageKi     SiteSearchSearchParamsLanguage = "ki"
	SiteSearchSearchParamsLanguageKiKe   SiteSearchSearchParamsLanguage = "ki-ke"
	SiteSearchSearchParamsLanguageKk     SiteSearchSearchParamsLanguage = "kk"
	SiteSearchSearchParamsLanguageKkKz   SiteSearchSearchParamsLanguage = "kk-kz"
	SiteSearchSearchParamsLanguageKkj    SiteSearchSearchParamsLanguage = "kkj"
	SiteSearchSearchParamsLanguageKkjCm  SiteSearchSearchParamsLanguage = "kkj-cm"
	SiteSearchSearchParamsLanguageKl     SiteSearchSearchParamsLanguage = "kl"
	SiteSearchSearchParamsLanguageKlGl   SiteSearchSearchParamsLanguage = "kl-gl"
	SiteSearchSearchParamsLanguageKln    SiteSearchSearchParamsLanguage = "kln"
	SiteSearchSearchParamsLanguageKlnKe  SiteSearchSearchParamsLanguage = "kln-ke"
	SiteSearchSearchParamsLanguageKm     SiteSearchSearchParamsLanguage = "km"
	SiteSearchSearchParamsLanguageKmKh   SiteSearchSearchParamsLanguage = "km-kh"
	SiteSearchSearchParamsLanguageKn     SiteSearchSearchParamsLanguage = "kn"
	SiteSearchSearchParamsLanguageKnIn   SiteSearchSearchParamsLanguage = "kn-in"
	SiteSearchSearchParamsLanguageKo     SiteSearchSearchParamsLanguage = "ko"
	SiteSearchSearchParamsLanguageKoKp   SiteSearchSearchParamsLanguage = "ko-kp"
	SiteSearchSearchParamsLanguageKoKr   SiteSearchSearchParamsLanguage = "ko-kr"
	SiteSearchSearchParamsLanguageKok    SiteSearchSearchParamsLanguage = "kok"
	SiteSearchSearchParamsLanguageKokIn  SiteSearchSearchParamsLanguage = "kok-in"
	SiteSearchSearchParamsLanguageKs     SiteSearchSearchParamsLanguage = "ks"
	SiteSearchSearchParamsLanguageKsIn   SiteSearchSearchParamsLanguage = "ks-in"
	SiteSearchSearchParamsLanguageKsb    SiteSearchSearchParamsLanguage = "ksb"
	SiteSearchSearchParamsLanguageKsbTz  SiteSearchSearchParamsLanguage = "ksb-tz"
	SiteSearchSearchParamsLanguageKsf    SiteSearchSearchParamsLanguage = "ksf"
	SiteSearchSearchParamsLanguageKsfCm  SiteSearchSearchParamsLanguage = "ksf-cm"
	SiteSearchSearchParamsLanguageKsh    SiteSearchSearchParamsLanguage = "ksh"
	SiteSearchSearchParamsLanguageKshDe  SiteSearchSearchParamsLanguage = "ksh-de"
	SiteSearchSearchParamsLanguageKu     SiteSearchSearchParamsLanguage = "ku"
	SiteSearchSearchParamsLanguageKuTr   SiteSearchSearchParamsLanguage = "ku-tr"
	SiteSearchSearchParamsLanguageKw     SiteSearchSearchParamsLanguage = "kw"
	SiteSearchSearchParamsLanguageKwGB   SiteSearchSearchParamsLanguage = "kw-gb"
	SiteSearchSearchParamsLanguageKy     SiteSearchSearchParamsLanguage = "ky"
	SiteSearchSearchParamsLanguageKyKg   SiteSearchSearchParamsLanguage = "ky-kg"
	SiteSearchSearchParamsLanguageLag    SiteSearchSearchParamsLanguage = "lag"
	SiteSearchSearchParamsLanguageLagTz  SiteSearchSearchParamsLanguage = "lag-tz"
	SiteSearchSearchParamsLanguageLb     SiteSearchSearchParamsLanguage = "lb"
	SiteSearchSearchParamsLanguageLbLu   SiteSearchSearchParamsLanguage = "lb-lu"
	SiteSearchSearchParamsLanguageLg     SiteSearchSearchParamsLanguage = "lg"
	SiteSearchSearchParamsLanguageLgUg   SiteSearchSearchParamsLanguage = "lg-ug"
	SiteSearchSearchParamsLanguageLkt    SiteSearchSearchParamsLanguage = "lkt"
	SiteSearchSearchParamsLanguageLktUs  SiteSearchSearchParamsLanguage = "lkt-us"
	SiteSearchSearchParamsLanguageLn     SiteSearchSearchParamsLanguage = "ln"
	SiteSearchSearchParamsLanguageLnAo   SiteSearchSearchParamsLanguage = "ln-ao"
	SiteSearchSearchParamsLanguageLnCd   SiteSearchSearchParamsLanguage = "ln-cd"
	SiteSearchSearchParamsLanguageLnCf   SiteSearchSearchParamsLanguage = "ln-cf"
	SiteSearchSearchParamsLanguageLnCg   SiteSearchSearchParamsLanguage = "ln-cg"
	SiteSearchSearchParamsLanguageLo     SiteSearchSearchParamsLanguage = "lo"
	SiteSearchSearchParamsLanguageLoLa   SiteSearchSearchParamsLanguage = "lo-la"
	SiteSearchSearchParamsLanguageLrc    SiteSearchSearchParamsLanguage = "lrc"
	SiteSearchSearchParamsLanguageLrcIq  SiteSearchSearchParamsLanguage = "lrc-iq"
	SiteSearchSearchParamsLanguageLrcIr  SiteSearchSearchParamsLanguage = "lrc-ir"
	SiteSearchSearchParamsLanguageLt     SiteSearchSearchParamsLanguage = "lt"
	SiteSearchSearchParamsLanguageLtLt   SiteSearchSearchParamsLanguage = "lt-lt"
	SiteSearchSearchParamsLanguageLu     SiteSearchSearchParamsLanguage = "lu"
	SiteSearchSearchParamsLanguageLuCd   SiteSearchSearchParamsLanguage = "lu-cd"
	SiteSearchSearchParamsLanguageLuo    SiteSearchSearchParamsLanguage = "luo"
	SiteSearchSearchParamsLanguageLuoKe  SiteSearchSearchParamsLanguage = "luo-ke"
	SiteSearchSearchParamsLanguageLuy    SiteSearchSearchParamsLanguage = "luy"
	SiteSearchSearchParamsLanguageLuyKe  SiteSearchSearchParamsLanguage = "luy-ke"
	SiteSearchSearchParamsLanguageLv     SiteSearchSearchParamsLanguage = "lv"
	SiteSearchSearchParamsLanguageLvLv   SiteSearchSearchParamsLanguage = "lv-lv"
	SiteSearchSearchParamsLanguageMai    SiteSearchSearchParamsLanguage = "mai"
	SiteSearchSearchParamsLanguageMaiIn  SiteSearchSearchParamsLanguage = "mai-in"
	SiteSearchSearchParamsLanguageMas    SiteSearchSearchParamsLanguage = "mas"
	SiteSearchSearchParamsLanguageMasKe  SiteSearchSearchParamsLanguage = "mas-ke"
	SiteSearchSearchParamsLanguageMasTz  SiteSearchSearchParamsLanguage = "mas-tz"
	SiteSearchSearchParamsLanguageMer    SiteSearchSearchParamsLanguage = "mer"
	SiteSearchSearchParamsLanguageMerKe  SiteSearchSearchParamsLanguage = "mer-ke"
	SiteSearchSearchParamsLanguageMfe    SiteSearchSearchParamsLanguage = "mfe"
	SiteSearchSearchParamsLanguageMfeMu  SiteSearchSearchParamsLanguage = "mfe-mu"
	SiteSearchSearchParamsLanguageMg     SiteSearchSearchParamsLanguage = "mg"
	SiteSearchSearchParamsLanguageMgMg   SiteSearchSearchParamsLanguage = "mg-mg"
	SiteSearchSearchParamsLanguageMgh    SiteSearchSearchParamsLanguage = "mgh"
	SiteSearchSearchParamsLanguageMghMz  SiteSearchSearchParamsLanguage = "mgh-mz"
	SiteSearchSearchParamsLanguageMgo    SiteSearchSearchParamsLanguage = "mgo"
	SiteSearchSearchParamsLanguageMgoCm  SiteSearchSearchParamsLanguage = "mgo-cm"
	SiteSearchSearchParamsLanguageMi     SiteSearchSearchParamsLanguage = "mi"
	SiteSearchSearchParamsLanguageMiNz   SiteSearchSearchParamsLanguage = "mi-nz"
	SiteSearchSearchParamsLanguageMk     SiteSearchSearchParamsLanguage = "mk"
	SiteSearchSearchParamsLanguageMkMk   SiteSearchSearchParamsLanguage = "mk-mk"
	SiteSearchSearchParamsLanguageMl     SiteSearchSearchParamsLanguage = "ml"
	SiteSearchSearchParamsLanguageMlIn   SiteSearchSearchParamsLanguage = "ml-in"
	SiteSearchSearchParamsLanguageMn     SiteSearchSearchParamsLanguage = "mn"
	SiteSearchSearchParamsLanguageMnMn   SiteSearchSearchParamsLanguage = "mn-mn"
	SiteSearchSearchParamsLanguageMni    SiteSearchSearchParamsLanguage = "mni"
	SiteSearchSearchParamsLanguageMniIn  SiteSearchSearchParamsLanguage = "mni-in"
	SiteSearchSearchParamsLanguageMr     SiteSearchSearchParamsLanguage = "mr"
	SiteSearchSearchParamsLanguageMrIn   SiteSearchSearchParamsLanguage = "mr-in"
	SiteSearchSearchParamsLanguageMs     SiteSearchSearchParamsLanguage = "ms"
	SiteSearchSearchParamsLanguageMsBn   SiteSearchSearchParamsLanguage = "ms-bn"
	SiteSearchSearchParamsLanguageMsID   SiteSearchSearchParamsLanguage = "ms-id"
	SiteSearchSearchParamsLanguageMsMy   SiteSearchSearchParamsLanguage = "ms-my"
	SiteSearchSearchParamsLanguageMsSg   SiteSearchSearchParamsLanguage = "ms-sg"
	SiteSearchSearchParamsLanguageMt     SiteSearchSearchParamsLanguage = "mt"
	SiteSearchSearchParamsLanguageMtMt   SiteSearchSearchParamsLanguage = "mt-mt"
	SiteSearchSearchParamsLanguageMua    SiteSearchSearchParamsLanguage = "mua"
	SiteSearchSearchParamsLanguageMuaCm  SiteSearchSearchParamsLanguage = "mua-cm"
	SiteSearchSearchParamsLanguageMy     SiteSearchSearchParamsLanguage = "my"
	SiteSearchSearchParamsLanguageMyMm   SiteSearchSearchParamsLanguage = "my-mm"
	SiteSearchSearchParamsLanguageMzn    SiteSearchSearchParamsLanguage = "mzn"
	SiteSearchSearchParamsLanguageMznIr  SiteSearchSearchParamsLanguage = "mzn-ir"
	SiteSearchSearchParamsLanguageNaq    SiteSearchSearchParamsLanguage = "naq"
	SiteSearchSearchParamsLanguageNaqNa  SiteSearchSearchParamsLanguage = "naq-na"
	SiteSearchSearchParamsLanguageNb     SiteSearchSearchParamsLanguage = "nb"
	SiteSearchSearchParamsLanguageNbNo   SiteSearchSearchParamsLanguage = "nb-no"
	SiteSearchSearchParamsLanguageNbSj   SiteSearchSearchParamsLanguage = "nb-sj"
	SiteSearchSearchParamsLanguageNd     SiteSearchSearchParamsLanguage = "nd"
	SiteSearchSearchParamsLanguageNdZw   SiteSearchSearchParamsLanguage = "nd-zw"
	SiteSearchSearchParamsLanguageNds    SiteSearchSearchParamsLanguage = "nds"
	SiteSearchSearchParamsLanguageNdsDe  SiteSearchSearchParamsLanguage = "nds-de"
	SiteSearchSearchParamsLanguageNdsNl  SiteSearchSearchParamsLanguage = "nds-nl"
	SiteSearchSearchParamsLanguageNe     SiteSearchSearchParamsLanguage = "ne"
	SiteSearchSearchParamsLanguageNeIn   SiteSearchSearchParamsLanguage = "ne-in"
	SiteSearchSearchParamsLanguageNeNp   SiteSearchSearchParamsLanguage = "ne-np"
	SiteSearchSearchParamsLanguageNl     SiteSearchSearchParamsLanguage = "nl"
	SiteSearchSearchParamsLanguageNlAw   SiteSearchSearchParamsLanguage = "nl-aw"
	SiteSearchSearchParamsLanguageNlBe   SiteSearchSearchParamsLanguage = "nl-be"
	SiteSearchSearchParamsLanguageNlBq   SiteSearchSearchParamsLanguage = "nl-bq"
	SiteSearchSearchParamsLanguageNlCh   SiteSearchSearchParamsLanguage = "nl-ch"
	SiteSearchSearchParamsLanguageNlCw   SiteSearchSearchParamsLanguage = "nl-cw"
	SiteSearchSearchParamsLanguageNlLu   SiteSearchSearchParamsLanguage = "nl-lu"
	SiteSearchSearchParamsLanguageNlNl   SiteSearchSearchParamsLanguage = "nl-nl"
	SiteSearchSearchParamsLanguageNlSr   SiteSearchSearchParamsLanguage = "nl-sr"
	SiteSearchSearchParamsLanguageNlSx   SiteSearchSearchParamsLanguage = "nl-sx"
	SiteSearchSearchParamsLanguageNmg    SiteSearchSearchParamsLanguage = "nmg"
	SiteSearchSearchParamsLanguageNmgCm  SiteSearchSearchParamsLanguage = "nmg-cm"
	SiteSearchSearchParamsLanguageNn     SiteSearchSearchParamsLanguage = "nn"
	SiteSearchSearchParamsLanguageNnNo   SiteSearchSearchParamsLanguage = "nn-no"
	SiteSearchSearchParamsLanguageNnh    SiteSearchSearchParamsLanguage = "nnh"
	SiteSearchSearchParamsLanguageNnhCm  SiteSearchSearchParamsLanguage = "nnh-cm"
	SiteSearchSearchParamsLanguageNo     SiteSearchSearchParamsLanguage = "no"
	SiteSearchSearchParamsLanguageNoNo   SiteSearchSearchParamsLanguage = "no-no"
	SiteSearchSearchParamsLanguageNus    SiteSearchSearchParamsLanguage = "nus"
	SiteSearchSearchParamsLanguageNusSS  SiteSearchSearchParamsLanguage = "nus-ss"
	SiteSearchSearchParamsLanguageNyn    SiteSearchSearchParamsLanguage = "nyn"
	SiteSearchSearchParamsLanguageNynUg  SiteSearchSearchParamsLanguage = "nyn-ug"
	SiteSearchSearchParamsLanguageOm     SiteSearchSearchParamsLanguage = "om"
	SiteSearchSearchParamsLanguageOmEt   SiteSearchSearchParamsLanguage = "om-et"
	SiteSearchSearchParamsLanguageOmKe   SiteSearchSearchParamsLanguage = "om-ke"
	SiteSearchSearchParamsLanguageOr     SiteSearchSearchParamsLanguage = "or"
	SiteSearchSearchParamsLanguageOrIn   SiteSearchSearchParamsLanguage = "or-in"
	SiteSearchSearchParamsLanguageOs     SiteSearchSearchParamsLanguage = "os"
	SiteSearchSearchParamsLanguageOsGe   SiteSearchSearchParamsLanguage = "os-ge"
	SiteSearchSearchParamsLanguageOsRu   SiteSearchSearchParamsLanguage = "os-ru"
	SiteSearchSearchParamsLanguagePa     SiteSearchSearchParamsLanguage = "pa"
	SiteSearchSearchParamsLanguagePaIn   SiteSearchSearchParamsLanguage = "pa-in"
	SiteSearchSearchParamsLanguagePaPk   SiteSearchSearchParamsLanguage = "pa-pk"
	SiteSearchSearchParamsLanguagePcm    SiteSearchSearchParamsLanguage = "pcm"
	SiteSearchSearchParamsLanguagePcmNg  SiteSearchSearchParamsLanguage = "pcm-ng"
	SiteSearchSearchParamsLanguagePl     SiteSearchSearchParamsLanguage = "pl"
	SiteSearchSearchParamsLanguagePlPl   SiteSearchSearchParamsLanguage = "pl-pl"
	SiteSearchSearchParamsLanguagePrg    SiteSearchSearchParamsLanguage = "prg"
	SiteSearchSearchParamsLanguagePrg001 SiteSearchSearchParamsLanguage = "prg-001"
	SiteSearchSearchParamsLanguagePs     SiteSearchSearchParamsLanguage = "ps"
	SiteSearchSearchParamsLanguagePsAf   SiteSearchSearchParamsLanguage = "ps-af"
	SiteSearchSearchParamsLanguagePsPk   SiteSearchSearchParamsLanguage = "ps-pk"
	SiteSearchSearchParamsLanguagePt     SiteSearchSearchParamsLanguage = "pt"
	SiteSearchSearchParamsLanguagePtAo   SiteSearchSearchParamsLanguage = "pt-ao"
	SiteSearchSearchParamsLanguagePtBr   SiteSearchSearchParamsLanguage = "pt-br"
	SiteSearchSearchParamsLanguagePtCh   SiteSearchSearchParamsLanguage = "pt-ch"
	SiteSearchSearchParamsLanguagePtCv   SiteSearchSearchParamsLanguage = "pt-cv"
	SiteSearchSearchParamsLanguagePtGq   SiteSearchSearchParamsLanguage = "pt-gq"
	SiteSearchSearchParamsLanguagePtGw   SiteSearchSearchParamsLanguage = "pt-gw"
	SiteSearchSearchParamsLanguagePtLu   SiteSearchSearchParamsLanguage = "pt-lu"
	SiteSearchSearchParamsLanguagePtMo   SiteSearchSearchParamsLanguage = "pt-mo"
	SiteSearchSearchParamsLanguagePtMz   SiteSearchSearchParamsLanguage = "pt-mz"
	SiteSearchSearchParamsLanguagePtPt   SiteSearchSearchParamsLanguage = "pt-pt"
	SiteSearchSearchParamsLanguagePtSt   SiteSearchSearchParamsLanguage = "pt-st"
	SiteSearchSearchParamsLanguagePtTl   SiteSearchSearchParamsLanguage = "pt-tl"
	SiteSearchSearchParamsLanguageQu     SiteSearchSearchParamsLanguage = "qu"
	SiteSearchSearchParamsLanguageQuBo   SiteSearchSearchParamsLanguage = "qu-bo"
	SiteSearchSearchParamsLanguageQuEc   SiteSearchSearchParamsLanguage = "qu-ec"
	SiteSearchSearchParamsLanguageQuPe   SiteSearchSearchParamsLanguage = "qu-pe"
	SiteSearchSearchParamsLanguageRm     SiteSearchSearchParamsLanguage = "rm"
	SiteSearchSearchParamsLanguageRmCh   SiteSearchSearchParamsLanguage = "rm-ch"
	SiteSearchSearchParamsLanguageRn     SiteSearchSearchParamsLanguage = "rn"
	SiteSearchSearchParamsLanguageRnBi   SiteSearchSearchParamsLanguage = "rn-bi"
	SiteSearchSearchParamsLanguageRo     SiteSearchSearchParamsLanguage = "ro"
	SiteSearchSearchParamsLanguageRoMd   SiteSearchSearchParamsLanguage = "ro-md"
	SiteSearchSearchParamsLanguageRoRo   SiteSearchSearchParamsLanguage = "ro-ro"
	SiteSearchSearchParamsLanguageRof    SiteSearchSearchParamsLanguage = "rof"
	SiteSearchSearchParamsLanguageRofTz  SiteSearchSearchParamsLanguage = "rof-tz"
	SiteSearchSearchParamsLanguageRu     SiteSearchSearchParamsLanguage = "ru"
	SiteSearchSearchParamsLanguageRuBy   SiteSearchSearchParamsLanguage = "ru-by"
	SiteSearchSearchParamsLanguageRuKg   SiteSearchSearchParamsLanguage = "ru-kg"
	SiteSearchSearchParamsLanguageRuKz   SiteSearchSearchParamsLanguage = "ru-kz"
	SiteSearchSearchParamsLanguageRuMd   SiteSearchSearchParamsLanguage = "ru-md"
	SiteSearchSearchParamsLanguageRuRu   SiteSearchSearchParamsLanguage = "ru-ru"
	SiteSearchSearchParamsLanguageRuUa   SiteSearchSearchParamsLanguage = "ru-ua"
	SiteSearchSearchParamsLanguageRw     SiteSearchSearchParamsLanguage = "rw"
	SiteSearchSearchParamsLanguageRwRw   SiteSearchSearchParamsLanguage = "rw-rw"
	SiteSearchSearchParamsLanguageRwk    SiteSearchSearchParamsLanguage = "rwk"
	SiteSearchSearchParamsLanguageRwkTz  SiteSearchSearchParamsLanguage = "rwk-tz"
	SiteSearchSearchParamsLanguageSa     SiteSearchSearchParamsLanguage = "sa"
	SiteSearchSearchParamsLanguageSaIn   SiteSearchSearchParamsLanguage = "sa-in"
	SiteSearchSearchParamsLanguageSah    SiteSearchSearchParamsLanguage = "sah"
	SiteSearchSearchParamsLanguageSahRu  SiteSearchSearchParamsLanguage = "sah-ru"
	SiteSearchSearchParamsLanguageSaq    SiteSearchSearchParamsLanguage = "saq"
	SiteSearchSearchParamsLanguageSaqKe  SiteSearchSearchParamsLanguage = "saq-ke"
	SiteSearchSearchParamsLanguageSat    SiteSearchSearchParamsLanguage = "sat"
	SiteSearchSearchParamsLanguageSatIn  SiteSearchSearchParamsLanguage = "sat-in"
	SiteSearchSearchParamsLanguageSbp    SiteSearchSearchParamsLanguage = "sbp"
	SiteSearchSearchParamsLanguageSbpTz  SiteSearchSearchParamsLanguage = "sbp-tz"
	SiteSearchSearchParamsLanguageSd     SiteSearchSearchParamsLanguage = "sd"
	SiteSearchSearchParamsLanguageSdIn   SiteSearchSearchParamsLanguage = "sd-in"
	SiteSearchSearchParamsLanguageSdPk   SiteSearchSearchParamsLanguage = "sd-pk"
	SiteSearchSearchParamsLanguageSe     SiteSearchSearchParamsLanguage = "se"
	SiteSearchSearchParamsLanguageSeFi   SiteSearchSearchParamsLanguage = "se-fi"
	SiteSearchSearchParamsLanguageSeNo   SiteSearchSearchParamsLanguage = "se-no"
	SiteSearchSearchParamsLanguageSeSe   SiteSearchSearchParamsLanguage = "se-se"
	SiteSearchSearchParamsLanguageSeh    SiteSearchSearchParamsLanguage = "seh"
	SiteSearchSearchParamsLanguageSehMz  SiteSearchSearchParamsLanguage = "seh-mz"
	SiteSearchSearchParamsLanguageSes    SiteSearchSearchParamsLanguage = "ses"
	SiteSearchSearchParamsLanguageSesMl  SiteSearchSearchParamsLanguage = "ses-ml"
	SiteSearchSearchParamsLanguageSg     SiteSearchSearchParamsLanguage = "sg"
	SiteSearchSearchParamsLanguageSgCf   SiteSearchSearchParamsLanguage = "sg-cf"
	SiteSearchSearchParamsLanguageShi    SiteSearchSearchParamsLanguage = "shi"
	SiteSearchSearchParamsLanguageShiMa  SiteSearchSearchParamsLanguage = "shi-ma"
	SiteSearchSearchParamsLanguageSi     SiteSearchSearchParamsLanguage = "si"
	SiteSearchSearchParamsLanguageSiLk   SiteSearchSearchParamsLanguage = "si-lk"
	SiteSearchSearchParamsLanguageSk     SiteSearchSearchParamsLanguage = "sk"
	SiteSearchSearchParamsLanguageSkSk   SiteSearchSearchParamsLanguage = "sk-sk"
	SiteSearchSearchParamsLanguageSl     SiteSearchSearchParamsLanguage = "sl"
	SiteSearchSearchParamsLanguageSlSi   SiteSearchSearchParamsLanguage = "sl-si"
	SiteSearchSearchParamsLanguageSmn    SiteSearchSearchParamsLanguage = "smn"
	SiteSearchSearchParamsLanguageSmnFi  SiteSearchSearchParamsLanguage = "smn-fi"
	SiteSearchSearchParamsLanguageSn     SiteSearchSearchParamsLanguage = "sn"
	SiteSearchSearchParamsLanguageSnZw   SiteSearchSearchParamsLanguage = "sn-zw"
	SiteSearchSearchParamsLanguageSo     SiteSearchSearchParamsLanguage = "so"
	SiteSearchSearchParamsLanguageSoDj   SiteSearchSearchParamsLanguage = "so-dj"
	SiteSearchSearchParamsLanguageSoEt   SiteSearchSearchParamsLanguage = "so-et"
	SiteSearchSearchParamsLanguageSoKe   SiteSearchSearchParamsLanguage = "so-ke"
	SiteSearchSearchParamsLanguageSoSo   SiteSearchSearchParamsLanguage = "so-so"
	SiteSearchSearchParamsLanguageSq     SiteSearchSearchParamsLanguage = "sq"
	SiteSearchSearchParamsLanguageSqAl   SiteSearchSearchParamsLanguage = "sq-al"
	SiteSearchSearchParamsLanguageSqMk   SiteSearchSearchParamsLanguage = "sq-mk"
	SiteSearchSearchParamsLanguageSqXk   SiteSearchSearchParamsLanguage = "sq-xk"
	SiteSearchSearchParamsLanguageSr     SiteSearchSearchParamsLanguage = "sr"
	SiteSearchSearchParamsLanguageSrBa   SiteSearchSearchParamsLanguage = "sr-ba"
	SiteSearchSearchParamsLanguageSrCs   SiteSearchSearchParamsLanguage = "sr-cs"
	SiteSearchSearchParamsLanguageSrMe   SiteSearchSearchParamsLanguage = "sr-me"
	SiteSearchSearchParamsLanguageSrRs   SiteSearchSearchParamsLanguage = "sr-rs"
	SiteSearchSearchParamsLanguageSrXk   SiteSearchSearchParamsLanguage = "sr-xk"
	SiteSearchSearchParamsLanguageSu     SiteSearchSearchParamsLanguage = "su"
	SiteSearchSearchParamsLanguageSuID   SiteSearchSearchParamsLanguage = "su-id"
	SiteSearchSearchParamsLanguageSv     SiteSearchSearchParamsLanguage = "sv"
	SiteSearchSearchParamsLanguageSvAx   SiteSearchSearchParamsLanguage = "sv-ax"
	SiteSearchSearchParamsLanguageSvFi   SiteSearchSearchParamsLanguage = "sv-fi"
	SiteSearchSearchParamsLanguageSvSe   SiteSearchSearchParamsLanguage = "sv-se"
	SiteSearchSearchParamsLanguageSw     SiteSearchSearchParamsLanguage = "sw"
	SiteSearchSearchParamsLanguageSwCd   SiteSearchSearchParamsLanguage = "sw-cd"
	SiteSearchSearchParamsLanguageSwKe   SiteSearchSearchParamsLanguage = "sw-ke"
	SiteSearchSearchParamsLanguageSwTz   SiteSearchSearchParamsLanguage = "sw-tz"
	SiteSearchSearchParamsLanguageSwUg   SiteSearchSearchParamsLanguage = "sw-ug"
	SiteSearchSearchParamsLanguageSy     SiteSearchSearchParamsLanguage = "sy"
	SiteSearchSearchParamsLanguageTa     SiteSearchSearchParamsLanguage = "ta"
	SiteSearchSearchParamsLanguageTaIn   SiteSearchSearchParamsLanguage = "ta-in"
	SiteSearchSearchParamsLanguageTaLk   SiteSearchSearchParamsLanguage = "ta-lk"
	SiteSearchSearchParamsLanguageTaMy   SiteSearchSearchParamsLanguage = "ta-my"
	SiteSearchSearchParamsLanguageTaSg   SiteSearchSearchParamsLanguage = "ta-sg"
	SiteSearchSearchParamsLanguageTe     SiteSearchSearchParamsLanguage = "te"
	SiteSearchSearchParamsLanguageTeIn   SiteSearchSearchParamsLanguage = "te-in"
	SiteSearchSearchParamsLanguageTeo    SiteSearchSearchParamsLanguage = "teo"
	SiteSearchSearchParamsLanguageTeoKe  SiteSearchSearchParamsLanguage = "teo-ke"
	SiteSearchSearchParamsLanguageTeoUg  SiteSearchSearchParamsLanguage = "teo-ug"
	SiteSearchSearchParamsLanguageTg     SiteSearchSearchParamsLanguage = "tg"
	SiteSearchSearchParamsLanguageTgTj   SiteSearchSearchParamsLanguage = "tg-tj"
	SiteSearchSearchParamsLanguageTh     SiteSearchSearchParamsLanguage = "th"
	SiteSearchSearchParamsLanguageThTh   SiteSearchSearchParamsLanguage = "th-th"
	SiteSearchSearchParamsLanguageTi     SiteSearchSearchParamsLanguage = "ti"
	SiteSearchSearchParamsLanguageTiEr   SiteSearchSearchParamsLanguage = "ti-er"
	SiteSearchSearchParamsLanguageTiEt   SiteSearchSearchParamsLanguage = "ti-et"
	SiteSearchSearchParamsLanguageTk     SiteSearchSearchParamsLanguage = "tk"
	SiteSearchSearchParamsLanguageTkTm   SiteSearchSearchParamsLanguage = "tk-tm"
	SiteSearchSearchParamsLanguageTl     SiteSearchSearchParamsLanguage = "tl"
	SiteSearchSearchParamsLanguageTo     SiteSearchSearchParamsLanguage = "to"
	SiteSearchSearchParamsLanguageToTo   SiteSearchSearchParamsLanguage = "to-to"
	SiteSearchSearchParamsLanguageTr     SiteSearchSearchParamsLanguage = "tr"
	SiteSearchSearchParamsLanguageTrCy   SiteSearchSearchParamsLanguage = "tr-cy"
	SiteSearchSearchParamsLanguageTrTr   SiteSearchSearchParamsLanguage = "tr-tr"
	SiteSearchSearchParamsLanguageTt     SiteSearchSearchParamsLanguage = "tt"
	SiteSearchSearchParamsLanguageTtRu   SiteSearchSearchParamsLanguage = "tt-ru"
	SiteSearchSearchParamsLanguageTwq    SiteSearchSearchParamsLanguage = "twq"
	SiteSearchSearchParamsLanguageTwqNe  SiteSearchSearchParamsLanguage = "twq-ne"
	SiteSearchSearchParamsLanguageTzm    SiteSearchSearchParamsLanguage = "tzm"
	SiteSearchSearchParamsLanguageTzmMa  SiteSearchSearchParamsLanguage = "tzm-ma"
	SiteSearchSearchParamsLanguageUg     SiteSearchSearchParamsLanguage = "ug"
	SiteSearchSearchParamsLanguageUgCn   SiteSearchSearchParamsLanguage = "ug-cn"
	SiteSearchSearchParamsLanguageUk     SiteSearchSearchParamsLanguage = "uk"
	SiteSearchSearchParamsLanguageUkUa   SiteSearchSearchParamsLanguage = "uk-ua"
	SiteSearchSearchParamsLanguageUr     SiteSearchSearchParamsLanguage = "ur"
	SiteSearchSearchParamsLanguageUrIn   SiteSearchSearchParamsLanguage = "ur-in"
	SiteSearchSearchParamsLanguageUrPk   SiteSearchSearchParamsLanguage = "ur-pk"
	SiteSearchSearchParamsLanguageUz     SiteSearchSearchParamsLanguage = "uz"
	SiteSearchSearchParamsLanguageUzAf   SiteSearchSearchParamsLanguage = "uz-af"
	SiteSearchSearchParamsLanguageUzUz   SiteSearchSearchParamsLanguage = "uz-uz"
	SiteSearchSearchParamsLanguageVai    SiteSearchSearchParamsLanguage = "vai"
	SiteSearchSearchParamsLanguageVaiLr  SiteSearchSearchParamsLanguage = "vai-lr"
	SiteSearchSearchParamsLanguageVi     SiteSearchSearchParamsLanguage = "vi"
	SiteSearchSearchParamsLanguageViVn   SiteSearchSearchParamsLanguage = "vi-vn"
	SiteSearchSearchParamsLanguageVo     SiteSearchSearchParamsLanguage = "vo"
	SiteSearchSearchParamsLanguageVo001  SiteSearchSearchParamsLanguage = "vo-001"
	SiteSearchSearchParamsLanguageVun    SiteSearchSearchParamsLanguage = "vun"
	SiteSearchSearchParamsLanguageVunTz  SiteSearchSearchParamsLanguage = "vun-tz"
	SiteSearchSearchParamsLanguageWae    SiteSearchSearchParamsLanguage = "wae"
	SiteSearchSearchParamsLanguageWaeCh  SiteSearchSearchParamsLanguage = "wae-ch"
	SiteSearchSearchParamsLanguageWo     SiteSearchSearchParamsLanguage = "wo"
	SiteSearchSearchParamsLanguageWoSn   SiteSearchSearchParamsLanguage = "wo-sn"
	SiteSearchSearchParamsLanguageXh     SiteSearchSearchParamsLanguage = "xh"
	SiteSearchSearchParamsLanguageXhZa   SiteSearchSearchParamsLanguage = "xh-za"
	SiteSearchSearchParamsLanguageXog    SiteSearchSearchParamsLanguage = "xog"
	SiteSearchSearchParamsLanguageXogUg  SiteSearchSearchParamsLanguage = "xog-ug"
	SiteSearchSearchParamsLanguageYav    SiteSearchSearchParamsLanguage = "yav"
	SiteSearchSearchParamsLanguageYavCm  SiteSearchSearchParamsLanguage = "yav-cm"
	SiteSearchSearchParamsLanguageYi     SiteSearchSearchParamsLanguage = "yi"
	SiteSearchSearchParamsLanguageYi001  SiteSearchSearchParamsLanguage = "yi-001"
	SiteSearchSearchParamsLanguageYo     SiteSearchSearchParamsLanguage = "yo"
	SiteSearchSearchParamsLanguageYoBj   SiteSearchSearchParamsLanguage = "yo-bj"
	SiteSearchSearchParamsLanguageYoNg   SiteSearchSearchParamsLanguage = "yo-ng"
	SiteSearchSearchParamsLanguageYue    SiteSearchSearchParamsLanguage = "yue"
	SiteSearchSearchParamsLanguageYueCn  SiteSearchSearchParamsLanguage = "yue-cn"
	SiteSearchSearchParamsLanguageYueHk  SiteSearchSearchParamsLanguage = "yue-hk"
	SiteSearchSearchParamsLanguageZgh    SiteSearchSearchParamsLanguage = "zgh"
	SiteSearchSearchParamsLanguageZghMa  SiteSearchSearchParamsLanguage = "zgh-ma"
	SiteSearchSearchParamsLanguageZh     SiteSearchSearchParamsLanguage = "zh"
	SiteSearchSearchParamsLanguageZhCn   SiteSearchSearchParamsLanguage = "zh-cn"
	SiteSearchSearchParamsLanguageZhHans SiteSearchSearchParamsLanguage = "zh-hans"
	SiteSearchSearchParamsLanguageZhHant SiteSearchSearchParamsLanguage = "zh-hant"
	SiteSearchSearchParamsLanguageZhHk   SiteSearchSearchParamsLanguage = "zh-hk"
	SiteSearchSearchParamsLanguageZhMo   SiteSearchSearchParamsLanguage = "zh-mo"
	SiteSearchSearchParamsLanguageZhSg   SiteSearchSearchParamsLanguage = "zh-sg"
	SiteSearchSearchParamsLanguageZhTw   SiteSearchSearchParamsLanguage = "zh-tw"
	SiteSearchSearchParamsLanguageZu     SiteSearchSearchParamsLanguage = "zu"
	SiteSearchSearchParamsLanguageZuZa   SiteSearchSearchParamsLanguage = "zu-za"
)

// Specifies the length of the search results. Can be set to `LONG` or `SHORT`.
// `SHORT` will return the first 128 characters of the content's meta description.
// `LONG` will build a more detailed content snippet based on the html/content of
// the page.
type SiteSearchSearchParamsLength string

const (
	SiteSearchSearchParamsLengthLong  SiteSearchSearchParamsLength = "LONG"
	SiteSearchSearchParamsLengthShort SiteSearchSearchParamsLength = "SHORT"
)
