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
// with the Hubspot API.
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
func (r *SiteSearchService) GetIndexedData(ctx context.Context, contentID string, query SiteSearchGetIndexedDataParams, opts ...option.RequestOption) (res *SiteSearchGetIndexedDataResponse, err error) {
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
func (r *SiteSearchService) Search(ctx context.Context, query SiteSearchSearchParams, opts ...option.RequestOption) (res *SiteSearchSearchResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "cms/v3/site-search/search"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// The indexed data in HubSpot
type SiteSearchGetIndexedDataResponse struct {
	// The ID of the document in HubSpot.
	ID string `json:"id,required"`
	// The indexed fields in HubSpot.
	Fields map[string]SiteSearchGetIndexedDataResponseField `json:"fields,required"`
	// The type of document. Can be `SITE_PAGE`, `LANDING_PAGE`, `BLOG_POST`,
	// `LISTING_PAGE`, or `KNOWLEDGE_ARTICLE`.
	//
	// Any of "LANDING_PAGE", "BLOG_POST", "SITE_PAGE", "KNOWLEDGE_ARTICLE",
	// "LISTING_PAGE".
	Type SiteSearchGetIndexedDataResponseType `json:"type,required"`
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
func (r SiteSearchGetIndexedDataResponse) RawJSON() string { return r.JSON.raw }
func (r *SiteSearchGetIndexedDataResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SiteSearchGetIndexedDataResponseField struct {
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
func (r SiteSearchGetIndexedDataResponseField) RawJSON() string { return r.JSON.raw }
func (r *SiteSearchGetIndexedDataResponseField) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The type of document. Can be `SITE_PAGE`, `LANDING_PAGE`, `BLOG_POST`,
// `LISTING_PAGE`, or `KNOWLEDGE_ARTICLE`.
type SiteSearchGetIndexedDataResponseType string

const (
	SiteSearchGetIndexedDataResponseTypeLandingPage      SiteSearchGetIndexedDataResponseType = "LANDING_PAGE"
	SiteSearchGetIndexedDataResponseTypeBlogPost         SiteSearchGetIndexedDataResponseType = "BLOG_POST"
	SiteSearchGetIndexedDataResponseTypeSitePage         SiteSearchGetIndexedDataResponseType = "SITE_PAGE"
	SiteSearchGetIndexedDataResponseTypeKnowledgeArticle SiteSearchGetIndexedDataResponseType = "KNOWLEDGE_ARTICLE"
	SiteSearchGetIndexedDataResponseTypeListingPage      SiteSearchGetIndexedDataResponseType = "LISTING_PAGE"
)

type SiteSearchSearchResponse struct {
	Limit      int64                            `json:"limit,required"`
	Offset     int64                            `json:"offset,required"`
	Page       int64                            `json:"page,required"`
	Results    []SiteSearchSearchResponseResult `json:"results,required"`
	Total      int64                            `json:"total,required"`
	SearchTerm string                           `json:"searchTerm"`
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
func (r SiteSearchSearchResponse) RawJSON() string { return r.JSON.raw }
func (r *SiteSearchSearchResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An individual search result.
type SiteSearchSearchResponseResult struct {
	// The ID of the content.
	ID int64 `json:"id,required"`
	// The domain the document is hosted on.
	Domain string `json:"domain,required"`
	// The matching score of the document.
	Score float64 `json:"score,required"`
	// The type of document. Can be `SITE_PAGE`, `LANDING_PAGE`, `BLOG_POST`,
	// `LISTING_PAGE`, or `KNOWLEDGE_ARTICLE`.
	//
	// Any of "LANDING_PAGE", "BLOG_POST", "SITE_PAGE", "KNOWLEDGE_ARTICLE",
	// "LISTING_PAGE".
	Type string `json:"type,required"`
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
	Language string `json:"language"`
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
func (r SiteSearchSearchResponseResult) RawJSON() string { return r.JSON.raw }
func (r *SiteSearchSearchResponseResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SiteSearchGetIndexedDataParams struct {
	// The type of document. Can be one of `SITE_PAGE`, `BLOG_POST`, or
	// `KNOWLEDGE_ARTICLE`.
	//
	// Any of "LANDING_PAGE", "BLOG_POST", "SITE_PAGE", "KNOWLEDGE_ARTICLE",
	// "LISTING_PAGE".
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
	SiteSearchGetIndexedDataParamsTypeLandingPage      SiteSearchGetIndexedDataParamsType = "LANDING_PAGE"
	SiteSearchGetIndexedDataParamsTypeBlogPost         SiteSearchGetIndexedDataParamsType = "BLOG_POST"
	SiteSearchGetIndexedDataParamsTypeSitePage         SiteSearchGetIndexedDataParamsType = "SITE_PAGE"
	SiteSearchGetIndexedDataParamsTypeKnowledgeArticle SiteSearchGetIndexedDataParamsType = "KNOWLEDGE_ARTICLE"
	SiteSearchGetIndexedDataParamsTypeListingPage      SiteSearchGetIndexedDataParamsType = "LISTING_PAGE"
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
	Language SiteSearchSearchParamsLanguage `query:"language,omitzero" json:"-"`
	// Specifies the length of the search results. Can be set to `LONG` or `SHORT`.
	// `SHORT` will return the first 128 characters of the content's meta description.
	// `LONG` will build a more detailed content snippet based on the html/content of
	// the page.
	//
	// Any of "SHORT", "LONG".
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
	SiteSearchSearchParamsLanguageIg     SiteSearchSearchParamsLanguage = "ig"
	SiteSearchSearchParamsLanguageIgNg   SiteSearchSearchParamsLanguage = "ig-ng"
	SiteSearchSearchParamsLanguageIi     SiteSearchSearchParamsLanguage = "ii"
	SiteSearchSearchParamsLanguageIiCn   SiteSearchSearchParamsLanguage = "ii-cn"
	SiteSearchSearchParamsLanguageIDID   SiteSearchSearchParamsLanguage = "id-id"
	SiteSearchSearchParamsLanguageIs     SiteSearchSearchParamsLanguage = "is"
	SiteSearchSearchParamsLanguageIsIs   SiteSearchSearchParamsLanguage = "is-is"
	SiteSearchSearchParamsLanguageIt     SiteSearchSearchParamsLanguage = "it"
	SiteSearchSearchParamsLanguageItCh   SiteSearchSearchParamsLanguage = "it-ch"
	SiteSearchSearchParamsLanguageItIt   SiteSearchSearchParamsLanguage = "it-it"
	SiteSearchSearchParamsLanguageItSm   SiteSearchSearchParamsLanguage = "it-sm"
	SiteSearchSearchParamsLanguageItVa   SiteSearchSearchParamsLanguage = "it-va"
	SiteSearchSearchParamsLanguageHeIl   SiteSearchSearchParamsLanguage = "he-il"
	SiteSearchSearchParamsLanguageJa     SiteSearchSearchParamsLanguage = "ja"
	SiteSearchSearchParamsLanguageJaJp   SiteSearchSearchParamsLanguage = "ja-jp"
	SiteSearchSearchParamsLanguageJgo    SiteSearchSearchParamsLanguage = "jgo"
	SiteSearchSearchParamsLanguageJgoCm  SiteSearchSearchParamsLanguage = "jgo-cm"
	SiteSearchSearchParamsLanguageYi     SiteSearchSearchParamsLanguage = "yi"
	SiteSearchSearchParamsLanguageYi001  SiteSearchSearchParamsLanguage = "yi-001"
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
	SiteSearchSearchParamsLanguageKw     SiteSearchSearchParamsLanguage = "kw"
	SiteSearchSearchParamsLanguageKwGB   SiteSearchSearchParamsLanguage = "kw-gb"
	SiteSearchSearchParamsLanguageKu     SiteSearchSearchParamsLanguage = "ku"
	SiteSearchSearchParamsLanguageKuTr   SiteSearchSearchParamsLanguage = "ku-tr"
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
	SiteSearchSearchParamsLanguageNlCh   SiteSearchSearchParamsLanguage = "nl-ch"
	SiteSearchSearchParamsLanguageNlBq   SiteSearchSearchParamsLanguage = "nl-bq"
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
	SiteSearchSearchParamsLanguageZhHk   SiteSearchSearchParamsLanguage = "zh-hk"
	SiteSearchSearchParamsLanguageZhMo   SiteSearchSearchParamsLanguage = "zh-mo"
	SiteSearchSearchParamsLanguageZhSg   SiteSearchSearchParamsLanguage = "zh-sg"
	SiteSearchSearchParamsLanguageZhTw   SiteSearchSearchParamsLanguage = "zh-tw"
	SiteSearchSearchParamsLanguageZhHans SiteSearchSearchParamsLanguage = "zh-hans"
	SiteSearchSearchParamsLanguageZhHant SiteSearchSearchParamsLanguage = "zh-hant"
	SiteSearchSearchParamsLanguageZu     SiteSearchSearchParamsLanguage = "zu"
	SiteSearchSearchParamsLanguageZuZa   SiteSearchSearchParamsLanguage = "zu-za"
)

// Specifies the length of the search results. Can be set to `LONG` or `SHORT`.
// `SHORT` will return the first 128 characters of the content's meta description.
// `LONG` will build a more detailed content snippet based on the html/content of
// the page.
type SiteSearchSearchParamsLength string

const (
	SiteSearchSearchParamsLengthShort SiteSearchSearchParamsLength = "SHORT"
	SiteSearchSearchParamsLengthLong  SiteSearchSearchParamsLength = "LONG"
)
