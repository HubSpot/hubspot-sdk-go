// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cms

import (
	"encoding/json"
	"time"

	"github.com/stainless-sdks/hubspot-sdk-go/internal/apijson"
	"github.com/stainless-sdks/hubspot-sdk-go/option"
	"github.com/stainless-sdks/hubspot-sdk-go/packages/param"
	"github.com/stainless-sdks/hubspot-sdk-go/packages/respjson"
	"github.com/stainless-sdks/hubspot-sdk-go/shared"
)

// PageService contains methods and other services that help with interacting with
// the hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPageService] method instead.
type PageService struct {
	Options      []option.RequestOption
	LandingPages PageLandingPageService
	SitePages    PageSitePageService
}

// NewPageService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewPageService(opts ...option.RequestOption) (r PageService) {
	r = PageService{}
	r.Options = opts
	r.LandingPages = NewPageLandingPageService(opts...)
	r.SitePages = NewPageSitePageService(opts...)
	return
}

// Request body object for ending A/B tests.
//
// The properties AbTestID, WinnerID are required.
type AbTestEndRequestVNextParam struct {
	// ID of the test to end.
	AbTestID string `json:"abTestId,required"`
	// ID of the object to designate as the test winner.
	WinnerID string `json:"winnerId,required"`
	paramObj
}

func (r AbTestEndRequestVNextParam) MarshalJSON() (data []byte, err error) {
	type shadow AbTestEndRequestVNextParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AbTestEndRequestVNextParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Request body object for rerunning A/B tests.
//
// The properties AbTestID, VariationID are required.
type AbTestRerunRequestVNextParam struct {
	// ID of the test to rerun.
	AbTestID string `json:"abTestId,required"`
	// ID of the object to reactivate as a test variation.
	VariationID string `json:"variationId,required"`
	paramObj
}

func (r AbTestRerunRequestVNextParam) MarshalJSON() (data []byte, err error) {
	type shadow AbTestRerunRequestVNextParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AbTestRerunRequestVNextParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Wrapper for providing an array of content folders as inputs.
//
// The property Inputs is required.
type BatchInputContentFolderParam struct {
	// Content folders to input.
	Inputs []ContentFolderParam `json:"inputs,omitzero,required"`
	paramObj
}

func (r BatchInputContentFolderParam) MarshalJSON() (data []byte, err error) {
	type shadow BatchInputContentFolderParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BatchInputContentFolderParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Wrapper for providing an array of pages as inputs.
//
// The property Inputs is required.
type BatchInputPageParam struct {
	// Pages to input.
	Inputs []PageParam `json:"inputs,omitzero,required"`
	paramObj
}

func (r BatchInputPageParam) MarshalJSON() (data []byte, err error) {
	type shadow BatchInputPageParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BatchInputPageParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Response object for successful batch operations on content folders.
type BatchResponseContentFolder struct {
	CompletedAt time.Time       `json:"completedAt,required" format:"date-time"`
	Results     []ContentFolder `json:"results,required"`
	StartedAt   time.Time       `json:"startedAt,required" format:"date-time"`
	// Any of "CANCELED", "COMPLETE", "PENDING", "PROCESSING".
	Status      BatchResponseContentFolderStatus `json:"status,required"`
	Links       map[string]string                `json:"links"`
	RequestedAt time.Time                        `json:"requestedAt" format:"date-time"`
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
func (r BatchResponseContentFolder) RawJSON() string { return r.JSON.raw }
func (r *BatchResponseContentFolder) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BatchResponseContentFolderStatus string

const (
	BatchResponseContentFolderStatusCanceled   BatchResponseContentFolderStatus = "CANCELED"
	BatchResponseContentFolderStatusComplete   BatchResponseContentFolderStatus = "COMPLETE"
	BatchResponseContentFolderStatusPending    BatchResponseContentFolderStatus = "PENDING"
	BatchResponseContentFolderStatusProcessing BatchResponseContentFolderStatus = "PROCESSING"
)

// Response object for successful batch operations on pages.
type BatchResponsePage struct {
	CompletedAt time.Time `json:"completedAt,required" format:"date-time"`
	Results     []Page    `json:"results,required"`
	StartedAt   time.Time `json:"startedAt,required" format:"date-time"`
	// Any of "CANCELED", "COMPLETE", "PENDING", "PROCESSING".
	Status      BatchResponsePageStatus `json:"status,required"`
	Links       map[string]string       `json:"links"`
	RequestedAt time.Time               `json:"requestedAt" format:"date-time"`
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
func (r BatchResponsePage) RawJSON() string { return r.JSON.raw }
func (r *BatchResponsePage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BatchResponsePageStatus string

const (
	BatchResponsePageStatusCanceled   BatchResponsePageStatus = "CANCELED"
	BatchResponsePageStatusComplete   BatchResponsePageStatus = "COMPLETE"
	BatchResponsePageStatusPending    BatchResponsePageStatus = "PENDING"
	BatchResponsePageStatusProcessing BatchResponsePageStatus = "PROCESSING"
)

// Response object for collections of content folders with pagination information.
type CollectionResponseWithTotalContentFolderForwardPaging struct {
	// Collection of content folders.
	Results []ContentFolder `json:"results,required"`
	// Total number of content folders.
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
func (r CollectionResponseWithTotalContentFolderForwardPaging) RawJSON() string { return r.JSON.raw }
func (r *CollectionResponseWithTotalContentFolderForwardPaging) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Response object for collections of pages with pagination information.
type CollectionResponseWithTotalPageForwardPaging struct {
	// Collection of pages.
	Results []Page `json:"results,required"`
	// Total number of pages.
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
func (r CollectionResponseWithTotalPageForwardPaging) RawJSON() string { return r.JSON.raw }
func (r *CollectionResponseWithTotalPageForwardPaging) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Response object for collections of content folder versions with pagination
// information.
type CollectionResponseWithTotalVersionContentFolder struct {
	// Collection of content folder versions.
	Results []VersionContentFolder `json:"results,required"`
	// Total number of content folder versions.
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
func (r CollectionResponseWithTotalVersionContentFolder) RawJSON() string { return r.JSON.raw }
func (r *CollectionResponseWithTotalVersionContentFolder) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Response object for collections of page versions with pagination information.
type CollectionResponseWithTotalVersionPage struct {
	// Collection of page versions.
	Results []VersionPage `json:"results,required"`
	// Total number of page versions.
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
func (r CollectionResponseWithTotalVersionPage) RawJSON() string { return r.JSON.raw }
func (r *CollectionResponseWithTotalVersionPage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Model definition for a content folder.
type ContentFolder struct {
	// The unique ID of the content folder.
	ID string `json:"id,required"`
	// The type of object this folder applies to. Should always be LANDING_PAGE.
	Category int64     `json:"category,required"`
	Created  time.Time `json:"created,required" format:"date-time"`
	// The timestamp (ISO8601 format) when this content folder was deleted.
	DeletedAt time.Time `json:"deletedAt,required" format:"date-time"`
	// The name of the folder which will show up in the app dashboard
	Name string `json:"name,required"`
	// The ID of the content folder this folder is nested under
	ParentFolderID int64     `json:"parentFolderId,required"`
	Updated        time.Time `json:"updated,required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID             respjson.Field
		Category       respjson.Field
		Created        respjson.Field
		DeletedAt      respjson.Field
		Name           respjson.Field
		ParentFolderID respjson.Field
		Updated        respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ContentFolder) RawJSON() string { return r.JSON.raw }
func (r *ContentFolder) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this ContentFolder to a ContentFolderParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// ContentFolderParam.Overrides()
func (r ContentFolder) ToParam() ContentFolderParam {
	return param.Override[ContentFolderParam](json.RawMessage(r.RawJSON()))
}

// Model definition for a content folder.
//
// The properties ID, Category, Created, DeletedAt, Name, ParentFolderID, Updated
// are required.
type ContentFolderParam struct {
	// The unique ID of the content folder.
	ID string `json:"id,required"`
	// The type of object this folder applies to. Should always be LANDING_PAGE.
	Category int64     `json:"category,required"`
	Created  time.Time `json:"created,required" format:"date-time"`
	// The timestamp (ISO8601 format) when this content folder was deleted.
	DeletedAt time.Time `json:"deletedAt,required" format:"date-time"`
	// The name of the folder which will show up in the app dashboard
	Name string `json:"name,required"`
	// The ID of the content folder this folder is nested under
	ParentFolderID int64     `json:"parentFolderId,required"`
	Updated        time.Time `json:"updated,required" format:"date-time"`
	paramObj
}

func (r ContentFolderParam) MarshalJSON() (data []byte, err error) {
	type shadow ContentFolderParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ContentFolderParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Request body object for creating new language variant content.
//
// The property ID is required.
type ContentLanguageCloneRequestVNextParam struct {
	// ID of content to clone.
	ID string `json:"id,required"`
	// Target language of new variant.
	Language param.Opt[string] `json:"language,omitzero"`
	// Language of primary content to clone.
	PrimaryLanguage param.Opt[string] `json:"primaryLanguage,omitzero"`
	paramObj
}

func (r ContentLanguageCloneRequestVNextParam) MarshalJSON() (data []byte, err error) {
	type shadow ContentLanguageCloneRequestVNextParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ContentLanguageCloneRequestVNextParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Model definition for a landing page or site page.
type Page struct {
	// The unique ID of the page.
	ID string `json:"id,required"`
	// The status of the AB test associated with this page, if applicable
	//
	// Any of "automated_loser_variant", "automated_master", "automated_variant",
	// "loser_variant", "mab_master", "mab_variant", "master", "variant".
	AbStatus PageAbStatus `json:"abStatus,required"`
	// The ID of the AB test associated with this page, if applicable
	AbTestID string `json:"abTestId,required"`
	// The timestamp (ISO8601 format) when this page was deleted.
	ArchivedAt time.Time `json:"archivedAt,required" format:"date-time"`
	// If True, the page will not show up in your dashboard, although the page could
	// still be live.
	ArchivedInDashboard bool `json:"archivedInDashboard,required"`
	// List of stylesheets to attach to this page. These stylesheets are attached to
	// just this page. Order of precedence is bottom to top, just like in the HTML.
	AttachedStylesheets []map[string]any `json:"attachedStylesheets,required"`
	// The name of the user that updated this page.
	AuthorName string `json:"authorName,required"`
	// The GUID of the marketing campaign this page is a part of.
	Campaign string `json:"campaign,required"`
	// ID of the type of object this is. Should always .
	CategoryID     int64  `json:"categoryId,required"`
	ContentGroupID string `json:"contentGroupId,required"`
	// An ENUM descibing the type of this object. Should be either LANDING_PAGE or
	// SITE_PAGE.
	//
	// Any of "0", "1", "10", "11", "12", "2", "3", "4", "5", "6", "7", "8", "9".
	ContentTypeCategory PageContentTypeCategory `json:"contentTypeCategory,required"`
	Created             time.Time               `json:"created,required" format:"date-time"`
	// The ID of the user that created this page.
	CreatedByID        string `json:"createdById,required"`
	CurrentlyPublished bool   `json:"currentlyPublished,required"`
	// A generated ENUM descibing the current state of this page.
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
	CurrentState PageCurrentState `json:"currentState,required"`
	// The domain this page will resolve to. If null, the page will default to the
	// primary domain for this content type.
	Domain                    string `json:"domain,required"`
	DynamicPageDataSourceID   string `json:"dynamicPageDataSourceId,required"`
	DynamicPageDataSourceType int64  `json:"dynamicPageDataSourceType,required"`
	// The ID of the HubDB table this page references, if applicable
	DynamicPageHubDBTableID string `json:"dynamicPageHubDbTableId,required"`
	// Boolean to determine whether or not the styles from the template should be
	// applied.
	EnableDomainStylesheets bool `json:"enableDomainStylesheets,required"`
	// Boolean to determine whether or not the styles from the template should be
	// applied.
	EnableLayoutStylesheets bool `json:"enableLayoutStylesheets,required"`
	// The featuredImage of this page.
	FeaturedImage string `json:"featuredImage,required"`
	// Alt Text of the featuredImage.
	FeaturedImageAltText string `json:"featuredImageAltText,required"`
	// The ID of the associated folder this landing page is organized under in the app
	// dashboard.
	FolderID string `json:"folderId,required"`
	// Custom HTML for embed codes, javascript that should be placed before the </body>
	// tag of the page.
	FooterHTML string `json:"footerHtml,required"`
	// Custom HTML for embed codes, javascript, etc. that goes in the <head> tag of the
	// page.
	HeadHTML string `json:"headHtml,required"`
	// The html title of this page.
	HTMLTitle string `json:"htmlTitle,required"`
	// Boolean to determine whether or not the Primary CSS Files should be applied.
	IncludeDefaultCustomCss bool `json:"includeDefaultCustomCss,required"`
	// The explicitly defined ISO 639 language code of the page. If null, the page will
	// default to the language of the Domain.
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
	Language       PageLanguage             `json:"language,required"`
	LayoutSections map[string]LayoutSection `json:"layoutSections,required"`
	// Optional override to set the URL to be used in the rel=canonical link tag on the
	// page.
	LinkRelCanonicalURL string `json:"linkRelCanonicalUrl,required"`
	// The ID of the MAB test (or dynamic test) associated with this page, if
	// applicable
	MabExperimentID string `json:"mabExperimentId,required"`
	// A description that goes in <meta> tag on the page.
	MetaDescription string `json:"metaDescription,required"`
	// The internal name of the page.
	Name string `json:"name,required"`
	// The date at which this page should expire and begin redirecting to another url
	// or page.
	PageExpiryDate int64 `json:"pageExpiryDate,required"`
	// Boolean describing if the page expiration feature is enabled for this page
	PageExpiryEnabled bool `json:"pageExpiryEnabled,required"`
	// The ID of another page this page's url should redirect to once this page
	// expires. Should only set this or pageExpiryRedirectUrl.
	PageExpiryRedirectID int64 `json:"pageExpiryRedirectId,required"`
	// The URL this page's url should redirect to once this page expires. Should only
	// set this or pageExpiryRedirectId.
	PageExpiryRedirectURL string `json:"pageExpiryRedirectUrl,required"`
	// A generated Boolean describing whether or not this page is currently expired and
	// being redirected.
	PageRedirected bool `json:"pageRedirected,required"`
	// Set this to create a password protected page. Entering the password will be
	// required to view the page.
	Password string `json:"password,required"`
	// Rules for require member registration to access private content.
	PublicAccessRules []PublicAccessRule `json:"publicAccessRules,required"`
	// Boolean to determine whether or not to respect publicAccessRules.
	PublicAccessRulesEnabled bool `json:"publicAccessRulesEnabled,required"`
	// The date (ISO8601 format) the page is to be published at.
	PublishDate time.Time `json:"publishDate,required" format:"date-time"`
	// Set this to true if you want to be published immediately when the schedule
	// publish endpoint is called, and to ignore the publish_date setting.
	PublishImmediately bool `json:"publishImmediately,required"`
	// The path of the this page. This field is appended to the domain to construct the
	// url of this page.
	Slug string `json:"slug,required"`
	// An ENUM descibing the current state of this page.
	State string `json:"state,required"`
	// Details the type of page this is. Should always be landing_page or site_page
	Subcategory string `json:"subcategory,required"`
	// String detailing the path of the template used for this page.
	TemplatePath        string         `json:"templatePath,required"`
	ThemeSettingsValues map[string]any `json:"themeSettingsValues,required"`
	// ID of the primary page this object was translated from.
	TranslatedFromID string                                   `json:"translatedFromId,required"`
	Translations     map[string]PagesContentLanguageVariation `json:"translations,required"`
	Updated          time.Time                                `json:"updated,required" format:"date-time"`
	// The ID of the user that updated this page.
	UpdatedByID string `json:"updatedById,required"`
	// A generated field representing the URL of this page.
	URL string `json:"url,required"`
	// Boolean to determine if this page should use a featuredImage.
	UseFeaturedImage bool `json:"useFeaturedImage,required"`
	// A data structure containing the data for all the modules inside the containers
	// for this page. This will only be populated if the page has widget containers.
	WidgetContainers map[string]any `json:"widgetContainers,required"`
	// A data structure containing the data for all the modules for this page.
	Widgets map[string]any `json:"widgets,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                        respjson.Field
		AbStatus                  respjson.Field
		AbTestID                  respjson.Field
		ArchivedAt                respjson.Field
		ArchivedInDashboard       respjson.Field
		AttachedStylesheets       respjson.Field
		AuthorName                respjson.Field
		Campaign                  respjson.Field
		CategoryID                respjson.Field
		ContentGroupID            respjson.Field
		ContentTypeCategory       respjson.Field
		Created                   respjson.Field
		CreatedByID               respjson.Field
		CurrentlyPublished        respjson.Field
		CurrentState              respjson.Field
		Domain                    respjson.Field
		DynamicPageDataSourceID   respjson.Field
		DynamicPageDataSourceType respjson.Field
		DynamicPageHubDBTableID   respjson.Field
		EnableDomainStylesheets   respjson.Field
		EnableLayoutStylesheets   respjson.Field
		FeaturedImage             respjson.Field
		FeaturedImageAltText      respjson.Field
		FolderID                  respjson.Field
		FooterHTML                respjson.Field
		HeadHTML                  respjson.Field
		HTMLTitle                 respjson.Field
		IncludeDefaultCustomCss   respjson.Field
		Language                  respjson.Field
		LayoutSections            respjson.Field
		LinkRelCanonicalURL       respjson.Field
		MabExperimentID           respjson.Field
		MetaDescription           respjson.Field
		Name                      respjson.Field
		PageExpiryDate            respjson.Field
		PageExpiryEnabled         respjson.Field
		PageExpiryRedirectID      respjson.Field
		PageExpiryRedirectURL     respjson.Field
		PageRedirected            respjson.Field
		Password                  respjson.Field
		PublicAccessRules         respjson.Field
		PublicAccessRulesEnabled  respjson.Field
		PublishDate               respjson.Field
		PublishImmediately        respjson.Field
		Slug                      respjson.Field
		State                     respjson.Field
		Subcategory               respjson.Field
		TemplatePath              respjson.Field
		ThemeSettingsValues       respjson.Field
		TranslatedFromID          respjson.Field
		Translations              respjson.Field
		Updated                   respjson.Field
		UpdatedByID               respjson.Field
		URL                       respjson.Field
		UseFeaturedImage          respjson.Field
		WidgetContainers          respjson.Field
		Widgets                   respjson.Field
		ExtraFields               map[string]respjson.Field
		raw                       string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Page) RawJSON() string { return r.JSON.raw }
func (r *Page) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this Page to a PageParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// PageParam.Overrides()
func (r Page) ToParam() PageParam {
	return param.Override[PageParam](json.RawMessage(r.RawJSON()))
}

// The status of the AB test associated with this page, if applicable
type PageAbStatus string

const (
	PageAbStatusAutomatedLoserVariant PageAbStatus = "automated_loser_variant"
	PageAbStatusAutomatedMaster       PageAbStatus = "automated_master"
	PageAbStatusAutomatedVariant      PageAbStatus = "automated_variant"
	PageAbStatusLoserVariant          PageAbStatus = "loser_variant"
	PageAbStatusMabMaster             PageAbStatus = "mab_master"
	PageAbStatusMabVariant            PageAbStatus = "mab_variant"
	PageAbStatusMaster                PageAbStatus = "master"
	PageAbStatusVariant               PageAbStatus = "variant"
)

// An ENUM descibing the type of this object. Should be either LANDING_PAGE or
// SITE_PAGE.
type PageContentTypeCategory string

const (
	PageContentTypeCategory0  PageContentTypeCategory = "0"
	PageContentTypeCategory1  PageContentTypeCategory = "1"
	PageContentTypeCategory10 PageContentTypeCategory = "10"
	PageContentTypeCategory11 PageContentTypeCategory = "11"
	PageContentTypeCategory12 PageContentTypeCategory = "12"
	PageContentTypeCategory2  PageContentTypeCategory = "2"
	PageContentTypeCategory3  PageContentTypeCategory = "3"
	PageContentTypeCategory4  PageContentTypeCategory = "4"
	PageContentTypeCategory5  PageContentTypeCategory = "5"
	PageContentTypeCategory6  PageContentTypeCategory = "6"
	PageContentTypeCategory7  PageContentTypeCategory = "7"
	PageContentTypeCategory8  PageContentTypeCategory = "8"
	PageContentTypeCategory9  PageContentTypeCategory = "9"
)

// A generated ENUM descibing the current state of this page.
type PageCurrentState string

const (
	PageCurrentStateAutomated               PageCurrentState = "AUTOMATED"
	PageCurrentStateAutomatedAb             PageCurrentState = "AUTOMATED_AB"
	PageCurrentStateAutomatedAbVariant      PageCurrentState = "AUTOMATED_AB_VARIANT"
	PageCurrentStateAutomatedDraft          PageCurrentState = "AUTOMATED_DRAFT"
	PageCurrentStateAutomatedDraftAb        PageCurrentState = "AUTOMATED_DRAFT_AB"
	PageCurrentStateAutomatedDraftAbvariant PageCurrentState = "AUTOMATED_DRAFT_ABVARIANT"
	PageCurrentStateAutomatedForForm        PageCurrentState = "AUTOMATED_FOR_FORM"
	PageCurrentStateAutomatedForFormBuffer  PageCurrentState = "AUTOMATED_FOR_FORM_BUFFER"
	PageCurrentStateAutomatedForFormDraft   PageCurrentState = "AUTOMATED_FOR_FORM_DRAFT"
	PageCurrentStateAutomatedForFormLegacy  PageCurrentState = "AUTOMATED_FOR_FORM_LEGACY"
	PageCurrentStateAutomatedLoserAbvariant PageCurrentState = "AUTOMATED_LOSER_ABVARIANT"
	PageCurrentStateAutomatedSending        PageCurrentState = "AUTOMATED_SENDING"
	PageCurrentStateBlogEmailDraft          PageCurrentState = "BLOG_EMAIL_DRAFT"
	PageCurrentStateBlogEmailPublished      PageCurrentState = "BLOG_EMAIL_PUBLISHED"
	PageCurrentStateDraft                   PageCurrentState = "DRAFT"
	PageCurrentStateDraftAb                 PageCurrentState = "DRAFT_AB"
	PageCurrentStateDraftAbVariant          PageCurrentState = "DRAFT_AB_VARIANT"
	PageCurrentStateError                   PageCurrentState = "ERROR"
	PageCurrentStateLoserAbVariant          PageCurrentState = "LOSER_AB_VARIANT"
	PageCurrentStatePageStub                PageCurrentState = "PAGE_STUB"
	PageCurrentStatePreProcessing           PageCurrentState = "PRE_PROCESSING"
	PageCurrentStateProcessing              PageCurrentState = "PROCESSING"
	PageCurrentStatePublished               PageCurrentState = "PUBLISHED"
	PageCurrentStatePublishedAb             PageCurrentState = "PUBLISHED_AB"
	PageCurrentStatePublishedAbVariant      PageCurrentState = "PUBLISHED_AB_VARIANT"
	PageCurrentStatePublishedOrScheduled    PageCurrentState = "PUBLISHED_OR_SCHEDULED"
	PageCurrentStateRssToEmailDraft         PageCurrentState = "RSS_TO_EMAIL_DRAFT"
	PageCurrentStateRssToEmailPublished     PageCurrentState = "RSS_TO_EMAIL_PUBLISHED"
	PageCurrentStateScheduled               PageCurrentState = "SCHEDULED"
	PageCurrentStateScheduledAb             PageCurrentState = "SCHEDULED_AB"
	PageCurrentStateScheduledOrPublished    PageCurrentState = "SCHEDULED_OR_PUBLISHED"
)

// The explicitly defined ISO 639 language code of the page. If null, the page will
// default to the language of the Domain.
type PageLanguage string

const (
	PageLanguageAf     PageLanguage = "af"
	PageLanguageAfNa   PageLanguage = "af-na"
	PageLanguageAfZa   PageLanguage = "af-za"
	PageLanguageAgq    PageLanguage = "agq"
	PageLanguageAgqCm  PageLanguage = "agq-cm"
	PageLanguageAk     PageLanguage = "ak"
	PageLanguageAkGh   PageLanguage = "ak-gh"
	PageLanguageAm     PageLanguage = "am"
	PageLanguageAmEt   PageLanguage = "am-et"
	PageLanguageAr     PageLanguage = "ar"
	PageLanguageAr001  PageLanguage = "ar-001"
	PageLanguageArAe   PageLanguage = "ar-ae"
	PageLanguageArBh   PageLanguage = "ar-bh"
	PageLanguageArDj   PageLanguage = "ar-dj"
	PageLanguageArDz   PageLanguage = "ar-dz"
	PageLanguageArEg   PageLanguage = "ar-eg"
	PageLanguageArEh   PageLanguage = "ar-eh"
	PageLanguageArEr   PageLanguage = "ar-er"
	PageLanguageArIl   PageLanguage = "ar-il"
	PageLanguageArIq   PageLanguage = "ar-iq"
	PageLanguageArJo   PageLanguage = "ar-jo"
	PageLanguageArKm   PageLanguage = "ar-km"
	PageLanguageArKw   PageLanguage = "ar-kw"
	PageLanguageArLb   PageLanguage = "ar-lb"
	PageLanguageArLy   PageLanguage = "ar-ly"
	PageLanguageArMa   PageLanguage = "ar-ma"
	PageLanguageArMr   PageLanguage = "ar-mr"
	PageLanguageArOm   PageLanguage = "ar-om"
	PageLanguageArPs   PageLanguage = "ar-ps"
	PageLanguageArQa   PageLanguage = "ar-qa"
	PageLanguageArSa   PageLanguage = "ar-sa"
	PageLanguageArSd   PageLanguage = "ar-sd"
	PageLanguageArSo   PageLanguage = "ar-so"
	PageLanguageArSS   PageLanguage = "ar-ss"
	PageLanguageArSy   PageLanguage = "ar-sy"
	PageLanguageArTd   PageLanguage = "ar-td"
	PageLanguageArTn   PageLanguage = "ar-tn"
	PageLanguageArYe   PageLanguage = "ar-ye"
	PageLanguageAs     PageLanguage = "as"
	PageLanguageAsIn   PageLanguage = "as-in"
	PageLanguageAsa    PageLanguage = "asa"
	PageLanguageAsaTz  PageLanguage = "asa-tz"
	PageLanguageAst    PageLanguage = "ast"
	PageLanguageAstEs  PageLanguage = "ast-es"
	PageLanguageAz     PageLanguage = "az"
	PageLanguageAzAz   PageLanguage = "az-az"
	PageLanguageBas    PageLanguage = "bas"
	PageLanguageBasCm  PageLanguage = "bas-cm"
	PageLanguageBe     PageLanguage = "be"
	PageLanguageBeBy   PageLanguage = "be-by"
	PageLanguageBem    PageLanguage = "bem"
	PageLanguageBemZm  PageLanguage = "bem-zm"
	PageLanguageBez    PageLanguage = "bez"
	PageLanguageBezTz  PageLanguage = "bez-tz"
	PageLanguageBg     PageLanguage = "bg"
	PageLanguageBgBg   PageLanguage = "bg-bg"
	PageLanguageBm     PageLanguage = "bm"
	PageLanguageBmMl   PageLanguage = "bm-ml"
	PageLanguageBn     PageLanguage = "bn"
	PageLanguageBnBd   PageLanguage = "bn-bd"
	PageLanguageBnIn   PageLanguage = "bn-in"
	PageLanguageBo     PageLanguage = "bo"
	PageLanguageBoCn   PageLanguage = "bo-cn"
	PageLanguageBoIn   PageLanguage = "bo-in"
	PageLanguageBr     PageLanguage = "br"
	PageLanguageBrFr   PageLanguage = "br-fr"
	PageLanguageBrx    PageLanguage = "brx"
	PageLanguageBrxIn  PageLanguage = "brx-in"
	PageLanguageBs     PageLanguage = "bs"
	PageLanguageBsBa   PageLanguage = "bs-ba"
	PageLanguageCa     PageLanguage = "ca"
	PageLanguageCaAd   PageLanguage = "ca-ad"
	PageLanguageCaEs   PageLanguage = "ca-es"
	PageLanguageCaFr   PageLanguage = "ca-fr"
	PageLanguageCaIt   PageLanguage = "ca-it"
	PageLanguageCcp    PageLanguage = "ccp"
	PageLanguageCcpBd  PageLanguage = "ccp-bd"
	PageLanguageCcpIn  PageLanguage = "ccp-in"
	PageLanguageCe     PageLanguage = "ce"
	PageLanguageCeRu   PageLanguage = "ce-ru"
	PageLanguageCeb    PageLanguage = "ceb"
	PageLanguageCebPh  PageLanguage = "ceb-ph"
	PageLanguageCgg    PageLanguage = "cgg"
	PageLanguageCggUg  PageLanguage = "cgg-ug"
	PageLanguageChr    PageLanguage = "chr"
	PageLanguageChrUs  PageLanguage = "chr-us"
	PageLanguageCkb    PageLanguage = "ckb"
	PageLanguageCkbIq  PageLanguage = "ckb-iq"
	PageLanguageCkbIr  PageLanguage = "ckb-ir"
	PageLanguageCs     PageLanguage = "cs"
	PageLanguageCsCz   PageLanguage = "cs-cz"
	PageLanguageCu     PageLanguage = "cu"
	PageLanguageCuRu   PageLanguage = "cu-ru"
	PageLanguageCy     PageLanguage = "cy"
	PageLanguageCyGB   PageLanguage = "cy-gb"
	PageLanguageDa     PageLanguage = "da"
	PageLanguageDaDk   PageLanguage = "da-dk"
	PageLanguageDaGl   PageLanguage = "da-gl"
	PageLanguageDav    PageLanguage = "dav"
	PageLanguageDavKe  PageLanguage = "dav-ke"
	PageLanguageDe     PageLanguage = "de"
	PageLanguageDeAt   PageLanguage = "de-at"
	PageLanguageDeBe   PageLanguage = "de-be"
	PageLanguageDeCh   PageLanguage = "de-ch"
	PageLanguageDeDe   PageLanguage = "de-de"
	PageLanguageDeGr   PageLanguage = "de-gr"
	PageLanguageDeIt   PageLanguage = "de-it"
	PageLanguageDeLi   PageLanguage = "de-li"
	PageLanguageDeLu   PageLanguage = "de-lu"
	PageLanguageDje    PageLanguage = "dje"
	PageLanguageDjeNe  PageLanguage = "dje-ne"
	PageLanguageDoi    PageLanguage = "doi"
	PageLanguageDoiIn  PageLanguage = "doi-in"
	PageLanguageDsb    PageLanguage = "dsb"
	PageLanguageDsbDe  PageLanguage = "dsb-de"
	PageLanguageDua    PageLanguage = "dua"
	PageLanguageDuaCm  PageLanguage = "dua-cm"
	PageLanguageDyo    PageLanguage = "dyo"
	PageLanguageDyoSn  PageLanguage = "dyo-sn"
	PageLanguageDz     PageLanguage = "dz"
	PageLanguageDzBt   PageLanguage = "dz-bt"
	PageLanguageEbu    PageLanguage = "ebu"
	PageLanguageEbuKe  PageLanguage = "ebu-ke"
	PageLanguageEe     PageLanguage = "ee"
	PageLanguageEeGh   PageLanguage = "ee-gh"
	PageLanguageEeTg   PageLanguage = "ee-tg"
	PageLanguageEl     PageLanguage = "el"
	PageLanguageElCy   PageLanguage = "el-cy"
	PageLanguageElGr   PageLanguage = "el-gr"
	PageLanguageEn     PageLanguage = "en"
	PageLanguageEn001  PageLanguage = "en-001"
	PageLanguageEn150  PageLanguage = "en-150"
	PageLanguageEnAe   PageLanguage = "en-ae"
	PageLanguageEnAg   PageLanguage = "en-ag"
	PageLanguageEnAI   PageLanguage = "en-ai"
	PageLanguageEnAs   PageLanguage = "en-as"
	PageLanguageEnAt   PageLanguage = "en-at"
	PageLanguageEnAu   PageLanguage = "en-au"
	PageLanguageEnBb   PageLanguage = "en-bb"
	PageLanguageEnBe   PageLanguage = "en-be"
	PageLanguageEnBi   PageLanguage = "en-bi"
	PageLanguageEnBm   PageLanguage = "en-bm"
	PageLanguageEnBs   PageLanguage = "en-bs"
	PageLanguageEnBw   PageLanguage = "en-bw"
	PageLanguageEnBz   PageLanguage = "en-bz"
	PageLanguageEnCa   PageLanguage = "en-ca"
	PageLanguageEnCc   PageLanguage = "en-cc"
	PageLanguageEnCh   PageLanguage = "en-ch"
	PageLanguageEnCk   PageLanguage = "en-ck"
	PageLanguageEnCm   PageLanguage = "en-cm"
	PageLanguageEnCn   PageLanguage = "en-cn"
	PageLanguageEnCx   PageLanguage = "en-cx"
	PageLanguageEnCy   PageLanguage = "en-cy"
	PageLanguageEnDe   PageLanguage = "en-de"
	PageLanguageEnDg   PageLanguage = "en-dg"
	PageLanguageEnDk   PageLanguage = "en-dk"
	PageLanguageEnDm   PageLanguage = "en-dm"
	PageLanguageEnEr   PageLanguage = "en-er"
	PageLanguageEnFi   PageLanguage = "en-fi"
	PageLanguageEnFj   PageLanguage = "en-fj"
	PageLanguageEnFk   PageLanguage = "en-fk"
	PageLanguageEnFm   PageLanguage = "en-fm"
	PageLanguageEnGB   PageLanguage = "en-gb"
	PageLanguageEnGd   PageLanguage = "en-gd"
	PageLanguageEnGg   PageLanguage = "en-gg"
	PageLanguageEnGh   PageLanguage = "en-gh"
	PageLanguageEnGi   PageLanguage = "en-gi"
	PageLanguageEnGm   PageLanguage = "en-gm"
	PageLanguageEnGu   PageLanguage = "en-gu"
	PageLanguageEnGy   PageLanguage = "en-gy"
	PageLanguageEnHk   PageLanguage = "en-hk"
	PageLanguageEnIe   PageLanguage = "en-ie"
	PageLanguageEnIl   PageLanguage = "en-il"
	PageLanguageEnIm   PageLanguage = "en-im"
	PageLanguageEnIn   PageLanguage = "en-in"
	PageLanguageEnIo   PageLanguage = "en-io"
	PageLanguageEnJe   PageLanguage = "en-je"
	PageLanguageEnJm   PageLanguage = "en-jm"
	PageLanguageEnKe   PageLanguage = "en-ke"
	PageLanguageEnKi   PageLanguage = "en-ki"
	PageLanguageEnKn   PageLanguage = "en-kn"
	PageLanguageEnKy   PageLanguage = "en-ky"
	PageLanguageEnLc   PageLanguage = "en-lc"
	PageLanguageEnLr   PageLanguage = "en-lr"
	PageLanguageEnLs   PageLanguage = "en-ls"
	PageLanguageEnLu   PageLanguage = "en-lu"
	PageLanguageEnMg   PageLanguage = "en-mg"
	PageLanguageEnMh   PageLanguage = "en-mh"
	PageLanguageEnMo   PageLanguage = "en-mo"
	PageLanguageEnMp   PageLanguage = "en-mp"
	PageLanguageEnMs   PageLanguage = "en-ms"
	PageLanguageEnMt   PageLanguage = "en-mt"
	PageLanguageEnMu   PageLanguage = "en-mu"
	PageLanguageEnMw   PageLanguage = "en-mw"
	PageLanguageEnMx   PageLanguage = "en-mx"
	PageLanguageEnMy   PageLanguage = "en-my"
	PageLanguageEnNa   PageLanguage = "en-na"
	PageLanguageEnNf   PageLanguage = "en-nf"
	PageLanguageEnNg   PageLanguage = "en-ng"
	PageLanguageEnNl   PageLanguage = "en-nl"
	PageLanguageEnNr   PageLanguage = "en-nr"
	PageLanguageEnNu   PageLanguage = "en-nu"
	PageLanguageEnNz   PageLanguage = "en-nz"
	PageLanguageEnPg   PageLanguage = "en-pg"
	PageLanguageEnPh   PageLanguage = "en-ph"
	PageLanguageEnPk   PageLanguage = "en-pk"
	PageLanguageEnPn   PageLanguage = "en-pn"
	PageLanguageEnPr   PageLanguage = "en-pr"
	PageLanguageEnPw   PageLanguage = "en-pw"
	PageLanguageEnRw   PageLanguage = "en-rw"
	PageLanguageEnSb   PageLanguage = "en-sb"
	PageLanguageEnSc   PageLanguage = "en-sc"
	PageLanguageEnSd   PageLanguage = "en-sd"
	PageLanguageEnSe   PageLanguage = "en-se"
	PageLanguageEnSg   PageLanguage = "en-sg"
	PageLanguageEnSh   PageLanguage = "en-sh"
	PageLanguageEnSi   PageLanguage = "en-si"
	PageLanguageEnSl   PageLanguage = "en-sl"
	PageLanguageEnSS   PageLanguage = "en-ss"
	PageLanguageEnSx   PageLanguage = "en-sx"
	PageLanguageEnSz   PageLanguage = "en-sz"
	PageLanguageEnTc   PageLanguage = "en-tc"
	PageLanguageEnTk   PageLanguage = "en-tk"
	PageLanguageEnTo   PageLanguage = "en-to"
	PageLanguageEnTt   PageLanguage = "en-tt"
	PageLanguageEnTv   PageLanguage = "en-tv"
	PageLanguageEnTz   PageLanguage = "en-tz"
	PageLanguageEnUg   PageLanguage = "en-ug"
	PageLanguageEnUm   PageLanguage = "en-um"
	PageLanguageEnUs   PageLanguage = "en-us"
	PageLanguageEnVc   PageLanguage = "en-vc"
	PageLanguageEnVg   PageLanguage = "en-vg"
	PageLanguageEnVi   PageLanguage = "en-vi"
	PageLanguageEnVu   PageLanguage = "en-vu"
	PageLanguageEnWs   PageLanguage = "en-ws"
	PageLanguageEnZa   PageLanguage = "en-za"
	PageLanguageEnZm   PageLanguage = "en-zm"
	PageLanguageEnZw   PageLanguage = "en-zw"
	PageLanguageEo     PageLanguage = "eo"
	PageLanguageEo001  PageLanguage = "eo-001"
	PageLanguageEs     PageLanguage = "es"
	PageLanguageEs419  PageLanguage = "es-419"
	PageLanguageEsAr   PageLanguage = "es-ar"
	PageLanguageEsBo   PageLanguage = "es-bo"
	PageLanguageEsBr   PageLanguage = "es-br"
	PageLanguageEsBz   PageLanguage = "es-bz"
	PageLanguageEsCl   PageLanguage = "es-cl"
	PageLanguageEsCo   PageLanguage = "es-co"
	PageLanguageEsCr   PageLanguage = "es-cr"
	PageLanguageEsCu   PageLanguage = "es-cu"
	PageLanguageEsDo   PageLanguage = "es-do"
	PageLanguageEsEa   PageLanguage = "es-ea"
	PageLanguageEsEc   PageLanguage = "es-ec"
	PageLanguageEsEs   PageLanguage = "es-es"
	PageLanguageEsGq   PageLanguage = "es-gq"
	PageLanguageEsGt   PageLanguage = "es-gt"
	PageLanguageEsHn   PageLanguage = "es-hn"
	PageLanguageEsIc   PageLanguage = "es-ic"
	PageLanguageEsMx   PageLanguage = "es-mx"
	PageLanguageEsNi   PageLanguage = "es-ni"
	PageLanguageEsPa   PageLanguage = "es-pa"
	PageLanguageEsPe   PageLanguage = "es-pe"
	PageLanguageEsPh   PageLanguage = "es-ph"
	PageLanguageEsPr   PageLanguage = "es-pr"
	PageLanguageEsPy   PageLanguage = "es-py"
	PageLanguageEsSv   PageLanguage = "es-sv"
	PageLanguageEsUs   PageLanguage = "es-us"
	PageLanguageEsUy   PageLanguage = "es-uy"
	PageLanguageEsVe   PageLanguage = "es-ve"
	PageLanguageEt     PageLanguage = "et"
	PageLanguageEtEe   PageLanguage = "et-ee"
	PageLanguageEu     PageLanguage = "eu"
	PageLanguageEuEs   PageLanguage = "eu-es"
	PageLanguageEwo    PageLanguage = "ewo"
	PageLanguageEwoCm  PageLanguage = "ewo-cm"
	PageLanguageFa     PageLanguage = "fa"
	PageLanguageFaAf   PageLanguage = "fa-af"
	PageLanguageFaIr   PageLanguage = "fa-ir"
	PageLanguageFf     PageLanguage = "ff"
	PageLanguageFfBf   PageLanguage = "ff-bf"
	PageLanguageFfCm   PageLanguage = "ff-cm"
	PageLanguageFfGh   PageLanguage = "ff-gh"
	PageLanguageFfGm   PageLanguage = "ff-gm"
	PageLanguageFfGn   PageLanguage = "ff-gn"
	PageLanguageFfGw   PageLanguage = "ff-gw"
	PageLanguageFfLr   PageLanguage = "ff-lr"
	PageLanguageFfMr   PageLanguage = "ff-mr"
	PageLanguageFfNe   PageLanguage = "ff-ne"
	PageLanguageFfNg   PageLanguage = "ff-ng"
	PageLanguageFfSl   PageLanguage = "ff-sl"
	PageLanguageFfSn   PageLanguage = "ff-sn"
	PageLanguageFi     PageLanguage = "fi"
	PageLanguageFiFi   PageLanguage = "fi-fi"
	PageLanguageFil    PageLanguage = "fil"
	PageLanguageFilPh  PageLanguage = "fil-ph"
	PageLanguageFo     PageLanguage = "fo"
	PageLanguageFoDk   PageLanguage = "fo-dk"
	PageLanguageFoFo   PageLanguage = "fo-fo"
	PageLanguageFr     PageLanguage = "fr"
	PageLanguageFrBe   PageLanguage = "fr-be"
	PageLanguageFrBf   PageLanguage = "fr-bf"
	PageLanguageFrBi   PageLanguage = "fr-bi"
	PageLanguageFrBj   PageLanguage = "fr-bj"
	PageLanguageFrBl   PageLanguage = "fr-bl"
	PageLanguageFrCa   PageLanguage = "fr-ca"
	PageLanguageFrCd   PageLanguage = "fr-cd"
	PageLanguageFrCf   PageLanguage = "fr-cf"
	PageLanguageFrCg   PageLanguage = "fr-cg"
	PageLanguageFrCh   PageLanguage = "fr-ch"
	PageLanguageFrCi   PageLanguage = "fr-ci"
	PageLanguageFrCm   PageLanguage = "fr-cm"
	PageLanguageFrDj   PageLanguage = "fr-dj"
	PageLanguageFrDz   PageLanguage = "fr-dz"
	PageLanguageFrFr   PageLanguage = "fr-fr"
	PageLanguageFrGa   PageLanguage = "fr-ga"
	PageLanguageFrGf   PageLanguage = "fr-gf"
	PageLanguageFrGn   PageLanguage = "fr-gn"
	PageLanguageFrGp   PageLanguage = "fr-gp"
	PageLanguageFrGq   PageLanguage = "fr-gq"
	PageLanguageFrHt   PageLanguage = "fr-ht"
	PageLanguageFrKm   PageLanguage = "fr-km"
	PageLanguageFrLu   PageLanguage = "fr-lu"
	PageLanguageFrMa   PageLanguage = "fr-ma"
	PageLanguageFrMc   PageLanguage = "fr-mc"
	PageLanguageFrMf   PageLanguage = "fr-mf"
	PageLanguageFrMg   PageLanguage = "fr-mg"
	PageLanguageFrMl   PageLanguage = "fr-ml"
	PageLanguageFrMq   PageLanguage = "fr-mq"
	PageLanguageFrMr   PageLanguage = "fr-mr"
	PageLanguageFrMu   PageLanguage = "fr-mu"
	PageLanguageFrNc   PageLanguage = "fr-nc"
	PageLanguageFrNe   PageLanguage = "fr-ne"
	PageLanguageFrPf   PageLanguage = "fr-pf"
	PageLanguageFrPm   PageLanguage = "fr-pm"
	PageLanguageFrRe   PageLanguage = "fr-re"
	PageLanguageFrRw   PageLanguage = "fr-rw"
	PageLanguageFrSc   PageLanguage = "fr-sc"
	PageLanguageFrSn   PageLanguage = "fr-sn"
	PageLanguageFrSy   PageLanguage = "fr-sy"
	PageLanguageFrTd   PageLanguage = "fr-td"
	PageLanguageFrTg   PageLanguage = "fr-tg"
	PageLanguageFrTn   PageLanguage = "fr-tn"
	PageLanguageFrVu   PageLanguage = "fr-vu"
	PageLanguageFrWf   PageLanguage = "fr-wf"
	PageLanguageFrYt   PageLanguage = "fr-yt"
	PageLanguageFur    PageLanguage = "fur"
	PageLanguageFurIt  PageLanguage = "fur-it"
	PageLanguageFy     PageLanguage = "fy"
	PageLanguageFyNl   PageLanguage = "fy-nl"
	PageLanguageGa     PageLanguage = "ga"
	PageLanguageGaGB   PageLanguage = "ga-gb"
	PageLanguageGaIe   PageLanguage = "ga-ie"
	PageLanguageGd     PageLanguage = "gd"
	PageLanguageGdGB   PageLanguage = "gd-gb"
	PageLanguageGl     PageLanguage = "gl"
	PageLanguageGlEs   PageLanguage = "gl-es"
	PageLanguageGsw    PageLanguage = "gsw"
	PageLanguageGswCh  PageLanguage = "gsw-ch"
	PageLanguageGswFr  PageLanguage = "gsw-fr"
	PageLanguageGswLi  PageLanguage = "gsw-li"
	PageLanguageGu     PageLanguage = "gu"
	PageLanguageGuIn   PageLanguage = "gu-in"
	PageLanguageGuz    PageLanguage = "guz"
	PageLanguageGuzKe  PageLanguage = "guz-ke"
	PageLanguageGv     PageLanguage = "gv"
	PageLanguageGvIm   PageLanguage = "gv-im"
	PageLanguageHa     PageLanguage = "ha"
	PageLanguageHaGh   PageLanguage = "ha-gh"
	PageLanguageHaNe   PageLanguage = "ha-ne"
	PageLanguageHaNg   PageLanguage = "ha-ng"
	PageLanguageHaw    PageLanguage = "haw"
	PageLanguageHawUs  PageLanguage = "haw-us"
	PageLanguageHe     PageLanguage = "he"
	PageLanguageHeIl   PageLanguage = "he-il"
	PageLanguageHi     PageLanguage = "hi"
	PageLanguageHiIn   PageLanguage = "hi-in"
	PageLanguageHr     PageLanguage = "hr"
	PageLanguageHrBa   PageLanguage = "hr-ba"
	PageLanguageHrHr   PageLanguage = "hr-hr"
	PageLanguageHsb    PageLanguage = "hsb"
	PageLanguageHsbDe  PageLanguage = "hsb-de"
	PageLanguageHu     PageLanguage = "hu"
	PageLanguageHuHu   PageLanguage = "hu-hu"
	PageLanguageHy     PageLanguage = "hy"
	PageLanguageHyAm   PageLanguage = "hy-am"
	PageLanguageIa     PageLanguage = "ia"
	PageLanguageIa001  PageLanguage = "ia-001"
	PageLanguageID     PageLanguage = "id"
	PageLanguageIDID   PageLanguage = "id-id"
	PageLanguageIg     PageLanguage = "ig"
	PageLanguageIgNg   PageLanguage = "ig-ng"
	PageLanguageIi     PageLanguage = "ii"
	PageLanguageIiCn   PageLanguage = "ii-cn"
	PageLanguageIs     PageLanguage = "is"
	PageLanguageIsIs   PageLanguage = "is-is"
	PageLanguageIt     PageLanguage = "it"
	PageLanguageItCh   PageLanguage = "it-ch"
	PageLanguageItIt   PageLanguage = "it-it"
	PageLanguageItSm   PageLanguage = "it-sm"
	PageLanguageItVa   PageLanguage = "it-va"
	PageLanguageJa     PageLanguage = "ja"
	PageLanguageJaJp   PageLanguage = "ja-jp"
	PageLanguageJgo    PageLanguage = "jgo"
	PageLanguageJgoCm  PageLanguage = "jgo-cm"
	PageLanguageJmc    PageLanguage = "jmc"
	PageLanguageJmcTz  PageLanguage = "jmc-tz"
	PageLanguageJv     PageLanguage = "jv"
	PageLanguageJvID   PageLanguage = "jv-id"
	PageLanguageKa     PageLanguage = "ka"
	PageLanguageKaGe   PageLanguage = "ka-ge"
	PageLanguageKab    PageLanguage = "kab"
	PageLanguageKabDz  PageLanguage = "kab-dz"
	PageLanguageKam    PageLanguage = "kam"
	PageLanguageKamKe  PageLanguage = "kam-ke"
	PageLanguageKde    PageLanguage = "kde"
	PageLanguageKdeTz  PageLanguage = "kde-tz"
	PageLanguageKea    PageLanguage = "kea"
	PageLanguageKeaCv  PageLanguage = "kea-cv"
	PageLanguageKhq    PageLanguage = "khq"
	PageLanguageKhqMl  PageLanguage = "khq-ml"
	PageLanguageKi     PageLanguage = "ki"
	PageLanguageKiKe   PageLanguage = "ki-ke"
	PageLanguageKk     PageLanguage = "kk"
	PageLanguageKkKz   PageLanguage = "kk-kz"
	PageLanguageKkj    PageLanguage = "kkj"
	PageLanguageKkjCm  PageLanguage = "kkj-cm"
	PageLanguageKl     PageLanguage = "kl"
	PageLanguageKlGl   PageLanguage = "kl-gl"
	PageLanguageKln    PageLanguage = "kln"
	PageLanguageKlnKe  PageLanguage = "kln-ke"
	PageLanguageKm     PageLanguage = "km"
	PageLanguageKmKh   PageLanguage = "km-kh"
	PageLanguageKn     PageLanguage = "kn"
	PageLanguageKnIn   PageLanguage = "kn-in"
	PageLanguageKo     PageLanguage = "ko"
	PageLanguageKoKp   PageLanguage = "ko-kp"
	PageLanguageKoKr   PageLanguage = "ko-kr"
	PageLanguageKok    PageLanguage = "kok"
	PageLanguageKokIn  PageLanguage = "kok-in"
	PageLanguageKs     PageLanguage = "ks"
	PageLanguageKsIn   PageLanguage = "ks-in"
	PageLanguageKsb    PageLanguage = "ksb"
	PageLanguageKsbTz  PageLanguage = "ksb-tz"
	PageLanguageKsf    PageLanguage = "ksf"
	PageLanguageKsfCm  PageLanguage = "ksf-cm"
	PageLanguageKsh    PageLanguage = "ksh"
	PageLanguageKshDe  PageLanguage = "ksh-de"
	PageLanguageKu     PageLanguage = "ku"
	PageLanguageKuTr   PageLanguage = "ku-tr"
	PageLanguageKw     PageLanguage = "kw"
	PageLanguageKwGB   PageLanguage = "kw-gb"
	PageLanguageKy     PageLanguage = "ky"
	PageLanguageKyKg   PageLanguage = "ky-kg"
	PageLanguageLag    PageLanguage = "lag"
	PageLanguageLagTz  PageLanguage = "lag-tz"
	PageLanguageLb     PageLanguage = "lb"
	PageLanguageLbLu   PageLanguage = "lb-lu"
	PageLanguageLg     PageLanguage = "lg"
	PageLanguageLgUg   PageLanguage = "lg-ug"
	PageLanguageLkt    PageLanguage = "lkt"
	PageLanguageLktUs  PageLanguage = "lkt-us"
	PageLanguageLn     PageLanguage = "ln"
	PageLanguageLnAo   PageLanguage = "ln-ao"
	PageLanguageLnCd   PageLanguage = "ln-cd"
	PageLanguageLnCf   PageLanguage = "ln-cf"
	PageLanguageLnCg   PageLanguage = "ln-cg"
	PageLanguageLo     PageLanguage = "lo"
	PageLanguageLoLa   PageLanguage = "lo-la"
	PageLanguageLrc    PageLanguage = "lrc"
	PageLanguageLrcIq  PageLanguage = "lrc-iq"
	PageLanguageLrcIr  PageLanguage = "lrc-ir"
	PageLanguageLt     PageLanguage = "lt"
	PageLanguageLtLt   PageLanguage = "lt-lt"
	PageLanguageLu     PageLanguage = "lu"
	PageLanguageLuCd   PageLanguage = "lu-cd"
	PageLanguageLuo    PageLanguage = "luo"
	PageLanguageLuoKe  PageLanguage = "luo-ke"
	PageLanguageLuy    PageLanguage = "luy"
	PageLanguageLuyKe  PageLanguage = "luy-ke"
	PageLanguageLv     PageLanguage = "lv"
	PageLanguageLvLv   PageLanguage = "lv-lv"
	PageLanguageMai    PageLanguage = "mai"
	PageLanguageMaiIn  PageLanguage = "mai-in"
	PageLanguageMas    PageLanguage = "mas"
	PageLanguageMasKe  PageLanguage = "mas-ke"
	PageLanguageMasTz  PageLanguage = "mas-tz"
	PageLanguageMer    PageLanguage = "mer"
	PageLanguageMerKe  PageLanguage = "mer-ke"
	PageLanguageMfe    PageLanguage = "mfe"
	PageLanguageMfeMu  PageLanguage = "mfe-mu"
	PageLanguageMg     PageLanguage = "mg"
	PageLanguageMgMg   PageLanguage = "mg-mg"
	PageLanguageMgh    PageLanguage = "mgh"
	PageLanguageMghMz  PageLanguage = "mgh-mz"
	PageLanguageMgo    PageLanguage = "mgo"
	PageLanguageMgoCm  PageLanguage = "mgo-cm"
	PageLanguageMi     PageLanguage = "mi"
	PageLanguageMiNz   PageLanguage = "mi-nz"
	PageLanguageMk     PageLanguage = "mk"
	PageLanguageMkMk   PageLanguage = "mk-mk"
	PageLanguageMl     PageLanguage = "ml"
	PageLanguageMlIn   PageLanguage = "ml-in"
	PageLanguageMn     PageLanguage = "mn"
	PageLanguageMnMn   PageLanguage = "mn-mn"
	PageLanguageMni    PageLanguage = "mni"
	PageLanguageMniIn  PageLanguage = "mni-in"
	PageLanguageMr     PageLanguage = "mr"
	PageLanguageMrIn   PageLanguage = "mr-in"
	PageLanguageMs     PageLanguage = "ms"
	PageLanguageMsBn   PageLanguage = "ms-bn"
	PageLanguageMsID   PageLanguage = "ms-id"
	PageLanguageMsMy   PageLanguage = "ms-my"
	PageLanguageMsSg   PageLanguage = "ms-sg"
	PageLanguageMt     PageLanguage = "mt"
	PageLanguageMtMt   PageLanguage = "mt-mt"
	PageLanguageMua    PageLanguage = "mua"
	PageLanguageMuaCm  PageLanguage = "mua-cm"
	PageLanguageMy     PageLanguage = "my"
	PageLanguageMyMm   PageLanguage = "my-mm"
	PageLanguageMzn    PageLanguage = "mzn"
	PageLanguageMznIr  PageLanguage = "mzn-ir"
	PageLanguageNaq    PageLanguage = "naq"
	PageLanguageNaqNa  PageLanguage = "naq-na"
	PageLanguageNb     PageLanguage = "nb"
	PageLanguageNbNo   PageLanguage = "nb-no"
	PageLanguageNbSj   PageLanguage = "nb-sj"
	PageLanguageNd     PageLanguage = "nd"
	PageLanguageNdZw   PageLanguage = "nd-zw"
	PageLanguageNds    PageLanguage = "nds"
	PageLanguageNdsDe  PageLanguage = "nds-de"
	PageLanguageNdsNl  PageLanguage = "nds-nl"
	PageLanguageNe     PageLanguage = "ne"
	PageLanguageNeIn   PageLanguage = "ne-in"
	PageLanguageNeNp   PageLanguage = "ne-np"
	PageLanguageNl     PageLanguage = "nl"
	PageLanguageNlAw   PageLanguage = "nl-aw"
	PageLanguageNlBe   PageLanguage = "nl-be"
	PageLanguageNlBq   PageLanguage = "nl-bq"
	PageLanguageNlCh   PageLanguage = "nl-ch"
	PageLanguageNlCw   PageLanguage = "nl-cw"
	PageLanguageNlLu   PageLanguage = "nl-lu"
	PageLanguageNlNl   PageLanguage = "nl-nl"
	PageLanguageNlSr   PageLanguage = "nl-sr"
	PageLanguageNlSx   PageLanguage = "nl-sx"
	PageLanguageNmg    PageLanguage = "nmg"
	PageLanguageNmgCm  PageLanguage = "nmg-cm"
	PageLanguageNn     PageLanguage = "nn"
	PageLanguageNnNo   PageLanguage = "nn-no"
	PageLanguageNnh    PageLanguage = "nnh"
	PageLanguageNnhCm  PageLanguage = "nnh-cm"
	PageLanguageNo     PageLanguage = "no"
	PageLanguageNoNo   PageLanguage = "no-no"
	PageLanguageNus    PageLanguage = "nus"
	PageLanguageNusSS  PageLanguage = "nus-ss"
	PageLanguageNyn    PageLanguage = "nyn"
	PageLanguageNynUg  PageLanguage = "nyn-ug"
	PageLanguageOm     PageLanguage = "om"
	PageLanguageOmEt   PageLanguage = "om-et"
	PageLanguageOmKe   PageLanguage = "om-ke"
	PageLanguageOr     PageLanguage = "or"
	PageLanguageOrIn   PageLanguage = "or-in"
	PageLanguageOs     PageLanguage = "os"
	PageLanguageOsGe   PageLanguage = "os-ge"
	PageLanguageOsRu   PageLanguage = "os-ru"
	PageLanguagePa     PageLanguage = "pa"
	PageLanguagePaIn   PageLanguage = "pa-in"
	PageLanguagePaPk   PageLanguage = "pa-pk"
	PageLanguagePcm    PageLanguage = "pcm"
	PageLanguagePcmNg  PageLanguage = "pcm-ng"
	PageLanguagePl     PageLanguage = "pl"
	PageLanguagePlPl   PageLanguage = "pl-pl"
	PageLanguagePrg    PageLanguage = "prg"
	PageLanguagePrg001 PageLanguage = "prg-001"
	PageLanguagePs     PageLanguage = "ps"
	PageLanguagePsAf   PageLanguage = "ps-af"
	PageLanguagePsPk   PageLanguage = "ps-pk"
	PageLanguagePt     PageLanguage = "pt"
	PageLanguagePtAo   PageLanguage = "pt-ao"
	PageLanguagePtBr   PageLanguage = "pt-br"
	PageLanguagePtCh   PageLanguage = "pt-ch"
	PageLanguagePtCv   PageLanguage = "pt-cv"
	PageLanguagePtGq   PageLanguage = "pt-gq"
	PageLanguagePtGw   PageLanguage = "pt-gw"
	PageLanguagePtLu   PageLanguage = "pt-lu"
	PageLanguagePtMo   PageLanguage = "pt-mo"
	PageLanguagePtMz   PageLanguage = "pt-mz"
	PageLanguagePtPt   PageLanguage = "pt-pt"
	PageLanguagePtSt   PageLanguage = "pt-st"
	PageLanguagePtTl   PageLanguage = "pt-tl"
	PageLanguageQu     PageLanguage = "qu"
	PageLanguageQuBo   PageLanguage = "qu-bo"
	PageLanguageQuEc   PageLanguage = "qu-ec"
	PageLanguageQuPe   PageLanguage = "qu-pe"
	PageLanguageRm     PageLanguage = "rm"
	PageLanguageRmCh   PageLanguage = "rm-ch"
	PageLanguageRn     PageLanguage = "rn"
	PageLanguageRnBi   PageLanguage = "rn-bi"
	PageLanguageRo     PageLanguage = "ro"
	PageLanguageRoMd   PageLanguage = "ro-md"
	PageLanguageRoRo   PageLanguage = "ro-ro"
	PageLanguageRof    PageLanguage = "rof"
	PageLanguageRofTz  PageLanguage = "rof-tz"
	PageLanguageRu     PageLanguage = "ru"
	PageLanguageRuBy   PageLanguage = "ru-by"
	PageLanguageRuKg   PageLanguage = "ru-kg"
	PageLanguageRuKz   PageLanguage = "ru-kz"
	PageLanguageRuMd   PageLanguage = "ru-md"
	PageLanguageRuRu   PageLanguage = "ru-ru"
	PageLanguageRuUa   PageLanguage = "ru-ua"
	PageLanguageRw     PageLanguage = "rw"
	PageLanguageRwRw   PageLanguage = "rw-rw"
	PageLanguageRwk    PageLanguage = "rwk"
	PageLanguageRwkTz  PageLanguage = "rwk-tz"
	PageLanguageSa     PageLanguage = "sa"
	PageLanguageSaIn   PageLanguage = "sa-in"
	PageLanguageSah    PageLanguage = "sah"
	PageLanguageSahRu  PageLanguage = "sah-ru"
	PageLanguageSaq    PageLanguage = "saq"
	PageLanguageSaqKe  PageLanguage = "saq-ke"
	PageLanguageSat    PageLanguage = "sat"
	PageLanguageSatIn  PageLanguage = "sat-in"
	PageLanguageSbp    PageLanguage = "sbp"
	PageLanguageSbpTz  PageLanguage = "sbp-tz"
	PageLanguageSd     PageLanguage = "sd"
	PageLanguageSdIn   PageLanguage = "sd-in"
	PageLanguageSdPk   PageLanguage = "sd-pk"
	PageLanguageSe     PageLanguage = "se"
	PageLanguageSeFi   PageLanguage = "se-fi"
	PageLanguageSeNo   PageLanguage = "se-no"
	PageLanguageSeSe   PageLanguage = "se-se"
	PageLanguageSeh    PageLanguage = "seh"
	PageLanguageSehMz  PageLanguage = "seh-mz"
	PageLanguageSes    PageLanguage = "ses"
	PageLanguageSesMl  PageLanguage = "ses-ml"
	PageLanguageSg     PageLanguage = "sg"
	PageLanguageSgCf   PageLanguage = "sg-cf"
	PageLanguageShi    PageLanguage = "shi"
	PageLanguageShiMa  PageLanguage = "shi-ma"
	PageLanguageSi     PageLanguage = "si"
	PageLanguageSiLk   PageLanguage = "si-lk"
	PageLanguageSk     PageLanguage = "sk"
	PageLanguageSkSk   PageLanguage = "sk-sk"
	PageLanguageSl     PageLanguage = "sl"
	PageLanguageSlSi   PageLanguage = "sl-si"
	PageLanguageSmn    PageLanguage = "smn"
	PageLanguageSmnFi  PageLanguage = "smn-fi"
	PageLanguageSn     PageLanguage = "sn"
	PageLanguageSnZw   PageLanguage = "sn-zw"
	PageLanguageSo     PageLanguage = "so"
	PageLanguageSoDj   PageLanguage = "so-dj"
	PageLanguageSoEt   PageLanguage = "so-et"
	PageLanguageSoKe   PageLanguage = "so-ke"
	PageLanguageSoSo   PageLanguage = "so-so"
	PageLanguageSq     PageLanguage = "sq"
	PageLanguageSqAl   PageLanguage = "sq-al"
	PageLanguageSqMk   PageLanguage = "sq-mk"
	PageLanguageSqXk   PageLanguage = "sq-xk"
	PageLanguageSr     PageLanguage = "sr"
	PageLanguageSrBa   PageLanguage = "sr-ba"
	PageLanguageSrCs   PageLanguage = "sr-cs"
	PageLanguageSrMe   PageLanguage = "sr-me"
	PageLanguageSrRs   PageLanguage = "sr-rs"
	PageLanguageSrXk   PageLanguage = "sr-xk"
	PageLanguageSu     PageLanguage = "su"
	PageLanguageSuID   PageLanguage = "su-id"
	PageLanguageSv     PageLanguage = "sv"
	PageLanguageSvAx   PageLanguage = "sv-ax"
	PageLanguageSvFi   PageLanguage = "sv-fi"
	PageLanguageSvSe   PageLanguage = "sv-se"
	PageLanguageSw     PageLanguage = "sw"
	PageLanguageSwCd   PageLanguage = "sw-cd"
	PageLanguageSwKe   PageLanguage = "sw-ke"
	PageLanguageSwTz   PageLanguage = "sw-tz"
	PageLanguageSwUg   PageLanguage = "sw-ug"
	PageLanguageSy     PageLanguage = "sy"
	PageLanguageTa     PageLanguage = "ta"
	PageLanguageTaIn   PageLanguage = "ta-in"
	PageLanguageTaLk   PageLanguage = "ta-lk"
	PageLanguageTaMy   PageLanguage = "ta-my"
	PageLanguageTaSg   PageLanguage = "ta-sg"
	PageLanguageTe     PageLanguage = "te"
	PageLanguageTeIn   PageLanguage = "te-in"
	PageLanguageTeo    PageLanguage = "teo"
	PageLanguageTeoKe  PageLanguage = "teo-ke"
	PageLanguageTeoUg  PageLanguage = "teo-ug"
	PageLanguageTg     PageLanguage = "tg"
	PageLanguageTgTj   PageLanguage = "tg-tj"
	PageLanguageTh     PageLanguage = "th"
	PageLanguageThTh   PageLanguage = "th-th"
	PageLanguageTi     PageLanguage = "ti"
	PageLanguageTiEr   PageLanguage = "ti-er"
	PageLanguageTiEt   PageLanguage = "ti-et"
	PageLanguageTk     PageLanguage = "tk"
	PageLanguageTkTm   PageLanguage = "tk-tm"
	PageLanguageTl     PageLanguage = "tl"
	PageLanguageTo     PageLanguage = "to"
	PageLanguageToTo   PageLanguage = "to-to"
	PageLanguageTr     PageLanguage = "tr"
	PageLanguageTrCy   PageLanguage = "tr-cy"
	PageLanguageTrTr   PageLanguage = "tr-tr"
	PageLanguageTt     PageLanguage = "tt"
	PageLanguageTtRu   PageLanguage = "tt-ru"
	PageLanguageTwq    PageLanguage = "twq"
	PageLanguageTwqNe  PageLanguage = "twq-ne"
	PageLanguageTzm    PageLanguage = "tzm"
	PageLanguageTzmMa  PageLanguage = "tzm-ma"
	PageLanguageUg     PageLanguage = "ug"
	PageLanguageUgCn   PageLanguage = "ug-cn"
	PageLanguageUk     PageLanguage = "uk"
	PageLanguageUkUa   PageLanguage = "uk-ua"
	PageLanguageUr     PageLanguage = "ur"
	PageLanguageUrIn   PageLanguage = "ur-in"
	PageLanguageUrPk   PageLanguage = "ur-pk"
	PageLanguageUz     PageLanguage = "uz"
	PageLanguageUzAf   PageLanguage = "uz-af"
	PageLanguageUzUz   PageLanguage = "uz-uz"
	PageLanguageVai    PageLanguage = "vai"
	PageLanguageVaiLr  PageLanguage = "vai-lr"
	PageLanguageVi     PageLanguage = "vi"
	PageLanguageViVn   PageLanguage = "vi-vn"
	PageLanguageVo     PageLanguage = "vo"
	PageLanguageVo001  PageLanguage = "vo-001"
	PageLanguageVun    PageLanguage = "vun"
	PageLanguageVunTz  PageLanguage = "vun-tz"
	PageLanguageWae    PageLanguage = "wae"
	PageLanguageWaeCh  PageLanguage = "wae-ch"
	PageLanguageWo     PageLanguage = "wo"
	PageLanguageWoSn   PageLanguage = "wo-sn"
	PageLanguageXh     PageLanguage = "xh"
	PageLanguageXhZa   PageLanguage = "xh-za"
	PageLanguageXog    PageLanguage = "xog"
	PageLanguageXogUg  PageLanguage = "xog-ug"
	PageLanguageYav    PageLanguage = "yav"
	PageLanguageYavCm  PageLanguage = "yav-cm"
	PageLanguageYi     PageLanguage = "yi"
	PageLanguageYi001  PageLanguage = "yi-001"
	PageLanguageYo     PageLanguage = "yo"
	PageLanguageYoBj   PageLanguage = "yo-bj"
	PageLanguageYoNg   PageLanguage = "yo-ng"
	PageLanguageYue    PageLanguage = "yue"
	PageLanguageYueCn  PageLanguage = "yue-cn"
	PageLanguageYueHk  PageLanguage = "yue-hk"
	PageLanguageZgh    PageLanguage = "zgh"
	PageLanguageZghMa  PageLanguage = "zgh-ma"
	PageLanguageZh     PageLanguage = "zh"
	PageLanguageZhCn   PageLanguage = "zh-cn"
	PageLanguageZhHans PageLanguage = "zh-hans"
	PageLanguageZhHant PageLanguage = "zh-hant"
	PageLanguageZhHk   PageLanguage = "zh-hk"
	PageLanguageZhMo   PageLanguage = "zh-mo"
	PageLanguageZhSg   PageLanguage = "zh-sg"
	PageLanguageZhTw   PageLanguage = "zh-tw"
	PageLanguageZu     PageLanguage = "zu"
	PageLanguageZuZa   PageLanguage = "zu-za"
)

// Model definition for a landing page or site page.
//
// The properties ID, AbStatus, AbTestID, ArchivedAt, ArchivedInDashboard,
// AttachedStylesheets, AuthorName, Campaign, CategoryID, ContentGroupID,
// ContentTypeCategory, Created, CreatedByID, CurrentlyPublished, CurrentState,
// Domain, DynamicPageDataSourceID, DynamicPageDataSourceType,
// DynamicPageHubDBTableID, EnableDomainStylesheets, EnableLayoutStylesheets,
// FeaturedImage, FeaturedImageAltText, FolderID, FooterHTML, HeadHTML, HTMLTitle,
// IncludeDefaultCustomCss, Language, LayoutSections, LinkRelCanonicalURL,
// MabExperimentID, MetaDescription, Name, PageExpiryDate, PageExpiryEnabled,
// PageExpiryRedirectID, PageExpiryRedirectURL, PageRedirected, Password,
// PublicAccessRules, PublicAccessRulesEnabled, PublishDate, PublishImmediately,
// Slug, State, Subcategory, TemplatePath, ThemeSettingsValues, TranslatedFromID,
// Translations, Updated, UpdatedByID, URL, UseFeaturedImage, WidgetContainers,
// Widgets are required.
type PageParam struct {
	// The unique ID of the page.
	ID string `json:"id,required"`
	// The status of the AB test associated with this page, if applicable
	//
	// Any of "automated_loser_variant", "automated_master", "automated_variant",
	// "loser_variant", "mab_master", "mab_variant", "master", "variant".
	AbStatus PageAbStatus `json:"abStatus,omitzero,required"`
	// The ID of the AB test associated with this page, if applicable
	AbTestID string `json:"abTestId,required"`
	// The timestamp (ISO8601 format) when this page was deleted.
	ArchivedAt time.Time `json:"archivedAt,required" format:"date-time"`
	// If True, the page will not show up in your dashboard, although the page could
	// still be live.
	ArchivedInDashboard bool `json:"archivedInDashboard,required"`
	// List of stylesheets to attach to this page. These stylesheets are attached to
	// just this page. Order of precedence is bottom to top, just like in the HTML.
	AttachedStylesheets []map[string]any `json:"attachedStylesheets,omitzero,required"`
	// The name of the user that updated this page.
	AuthorName string `json:"authorName,required"`
	// The GUID of the marketing campaign this page is a part of.
	Campaign string `json:"campaign,required"`
	// ID of the type of object this is. Should always .
	CategoryID     int64  `json:"categoryId,required"`
	ContentGroupID string `json:"contentGroupId,required"`
	// An ENUM descibing the type of this object. Should be either LANDING_PAGE or
	// SITE_PAGE.
	//
	// Any of "0", "1", "10", "11", "12", "2", "3", "4", "5", "6", "7", "8", "9".
	ContentTypeCategory PageContentTypeCategory `json:"contentTypeCategory,omitzero,required"`
	Created             time.Time               `json:"created,required" format:"date-time"`
	// The ID of the user that created this page.
	CreatedByID        string `json:"createdById,required"`
	CurrentlyPublished bool   `json:"currentlyPublished,required"`
	// A generated ENUM descibing the current state of this page.
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
	CurrentState PageCurrentState `json:"currentState,omitzero,required"`
	// The domain this page will resolve to. If null, the page will default to the
	// primary domain for this content type.
	Domain                    string `json:"domain,required"`
	DynamicPageDataSourceID   string `json:"dynamicPageDataSourceId,required"`
	DynamicPageDataSourceType int64  `json:"dynamicPageDataSourceType,required"`
	// The ID of the HubDB table this page references, if applicable
	DynamicPageHubDBTableID string `json:"dynamicPageHubDbTableId,required"`
	// Boolean to determine whether or not the styles from the template should be
	// applied.
	EnableDomainStylesheets bool `json:"enableDomainStylesheets,required"`
	// Boolean to determine whether or not the styles from the template should be
	// applied.
	EnableLayoutStylesheets bool `json:"enableLayoutStylesheets,required"`
	// The featuredImage of this page.
	FeaturedImage string `json:"featuredImage,required"`
	// Alt Text of the featuredImage.
	FeaturedImageAltText string `json:"featuredImageAltText,required"`
	// The ID of the associated folder this landing page is organized under in the app
	// dashboard.
	FolderID string `json:"folderId,required"`
	// Custom HTML for embed codes, javascript that should be placed before the </body>
	// tag of the page.
	FooterHTML string `json:"footerHtml,required"`
	// Custom HTML for embed codes, javascript, etc. that goes in the <head> tag of the
	// page.
	HeadHTML string `json:"headHtml,required"`
	// The html title of this page.
	HTMLTitle string `json:"htmlTitle,required"`
	// Boolean to determine whether or not the Primary CSS Files should be applied.
	IncludeDefaultCustomCss bool `json:"includeDefaultCustomCss,required"`
	// The explicitly defined ISO 639 language code of the page. If null, the page will
	// default to the language of the Domain.
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
	Language       PageLanguage                  `json:"language,omitzero,required"`
	LayoutSections map[string]LayoutSectionParam `json:"layoutSections,omitzero,required"`
	// Optional override to set the URL to be used in the rel=canonical link tag on the
	// page.
	LinkRelCanonicalURL string `json:"linkRelCanonicalUrl,required"`
	// The ID of the MAB test (or dynamic test) associated with this page, if
	// applicable
	MabExperimentID string `json:"mabExperimentId,required"`
	// A description that goes in <meta> tag on the page.
	MetaDescription string `json:"metaDescription,required"`
	// The internal name of the page.
	Name string `json:"name,required"`
	// The date at which this page should expire and begin redirecting to another url
	// or page.
	PageExpiryDate int64 `json:"pageExpiryDate,required"`
	// Boolean describing if the page expiration feature is enabled for this page
	PageExpiryEnabled bool `json:"pageExpiryEnabled,required"`
	// The ID of another page this page's url should redirect to once this page
	// expires. Should only set this or pageExpiryRedirectUrl.
	PageExpiryRedirectID int64 `json:"pageExpiryRedirectId,required"`
	// The URL this page's url should redirect to once this page expires. Should only
	// set this or pageExpiryRedirectId.
	PageExpiryRedirectURL string `json:"pageExpiryRedirectUrl,required"`
	// A generated Boolean describing whether or not this page is currently expired and
	// being redirected.
	PageRedirected bool `json:"pageRedirected,required"`
	// Set this to create a password protected page. Entering the password will be
	// required to view the page.
	Password string `json:"password,required"`
	// Rules for require member registration to access private content.
	PublicAccessRules []PublicAccessRule `json:"publicAccessRules,omitzero,required"`
	// Boolean to determine whether or not to respect publicAccessRules.
	PublicAccessRulesEnabled bool `json:"publicAccessRulesEnabled,required"`
	// The date (ISO8601 format) the page is to be published at.
	PublishDate time.Time `json:"publishDate,required" format:"date-time"`
	// Set this to true if you want to be published immediately when the schedule
	// publish endpoint is called, and to ignore the publish_date setting.
	PublishImmediately bool `json:"publishImmediately,required"`
	// The path of the this page. This field is appended to the domain to construct the
	// url of this page.
	Slug string `json:"slug,required"`
	// An ENUM descibing the current state of this page.
	State string `json:"state,required"`
	// Details the type of page this is. Should always be landing_page or site_page
	Subcategory string `json:"subcategory,required"`
	// String detailing the path of the template used for this page.
	TemplatePath        string         `json:"templatePath,required"`
	ThemeSettingsValues map[string]any `json:"themeSettingsValues,omitzero,required"`
	// ID of the primary page this object was translated from.
	TranslatedFromID string                                        `json:"translatedFromId,required"`
	Translations     map[string]PagesContentLanguageVariationParam `json:"translations,omitzero,required"`
	Updated          time.Time                                     `json:"updated,required" format:"date-time"`
	// The ID of the user that updated this page.
	UpdatedByID string `json:"updatedById,required"`
	// A generated field representing the URL of this page.
	URL string `json:"url,required"`
	// Boolean to determine if this page should use a featuredImage.
	UseFeaturedImage bool `json:"useFeaturedImage,required"`
	// A data structure containing the data for all the modules inside the containers
	// for this page. This will only be populated if the page has widget containers.
	WidgetContainers map[string]any `json:"widgetContainers,omitzero,required"`
	// A data structure containing the data for all the modules for this page.
	Widgets map[string]any `json:"widgets,omitzero,required"`
	paramObj
}

func (r PageParam) MarshalJSON() (data []byte, err error) {
	type shadow PageParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PageParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PagesContentLanguageVariation struct {
	ID                       int64              `json:"id,required"`
	ArchivedInDashboard      bool               `json:"archivedInDashboard,required"`
	AuthorName               string             `json:"authorName,required"`
	Campaign                 string             `json:"campaign,required"`
	Created                  time.Time          `json:"created,required" format:"date-time"`
	Name                     string             `json:"name,required"`
	Password                 string             `json:"password,required"`
	PublicAccessRules        []PublicAccessRule `json:"publicAccessRules,required"`
	PublicAccessRulesEnabled bool               `json:"publicAccessRulesEnabled,required"`
	PublishDate              time.Time          `json:"publishDate,required" format:"date-time"`
	Slug                     string             `json:"slug,required"`
	State                    string             `json:"state,required"`
	Updated                  time.Time          `json:"updated,required" format:"date-time"`
	TagIDs                   []int64            `json:"tagIds"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                       respjson.Field
		ArchivedInDashboard      respjson.Field
		AuthorName               respjson.Field
		Campaign                 respjson.Field
		Created                  respjson.Field
		Name                     respjson.Field
		Password                 respjson.Field
		PublicAccessRules        respjson.Field
		PublicAccessRulesEnabled respjson.Field
		PublishDate              respjson.Field
		Slug                     respjson.Field
		State                    respjson.Field
		Updated                  respjson.Field
		TagIDs                   respjson.Field
		ExtraFields              map[string]respjson.Field
		raw                      string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PagesContentLanguageVariation) RawJSON() string { return r.JSON.raw }
func (r *PagesContentLanguageVariation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this PagesContentLanguageVariation to a
// PagesContentLanguageVariationParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// PagesContentLanguageVariationParam.Overrides()
func (r PagesContentLanguageVariation) ToParam() PagesContentLanguageVariationParam {
	return param.Override[PagesContentLanguageVariationParam](json.RawMessage(r.RawJSON()))
}

// The properties ID, ArchivedInDashboard, AuthorName, Campaign, Created, Name,
// Password, PublicAccessRules, PublicAccessRulesEnabled, PublishDate, Slug, State,
// Updated are required.
type PagesContentLanguageVariationParam struct {
	ID                       int64              `json:"id,required"`
	ArchivedInDashboard      bool               `json:"archivedInDashboard,required"`
	AuthorName               string             `json:"authorName,required"`
	Campaign                 string             `json:"campaign,required"`
	Created                  time.Time          `json:"created,required" format:"date-time"`
	Name                     string             `json:"name,required"`
	Password                 string             `json:"password,required"`
	PublicAccessRules        []PublicAccessRule `json:"publicAccessRules,omitzero,required"`
	PublicAccessRulesEnabled bool               `json:"publicAccessRulesEnabled,required"`
	PublishDate              time.Time          `json:"publishDate,required" format:"date-time"`
	Slug                     string             `json:"slug,required"`
	State                    string             `json:"state,required"`
	Updated                  time.Time          `json:"updated,required" format:"date-time"`
	TagIDs                   []int64            `json:"tagIds,omitzero"`
	paramObj
}

func (r PagesContentLanguageVariationParam) MarshalJSON() (data []byte, err error) {
	type shadow PagesContentLanguageVariationParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PagesContentLanguageVariationParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Model definition for a content folder version. Contains metadata describing the
// version of the folder. It can be used to view edit history of a folder.
type VersionContentFolder struct {
	// ID of this folder version.
	ID string `json:"id,required"`
	// Model definition for a content folder.
	Object    ContentFolder `json:"object,required"`
	UpdatedAt time.Time     `json:"updatedAt,required" format:"date-time"`
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
func (r VersionContentFolder) RawJSON() string { return r.JSON.raw }
func (r *VersionContentFolder) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Model definition for a landing page or site page version. Contains metadata
// describing the version of the page. It can be used to view edit history of a
// page.
type VersionPage struct {
	// ID of this page version.
	ID string `json:"id,required"`
	// Model definition for a landing page or site page.
	Object    Page      `json:"object,required"`
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
func (r VersionPage) RawJSON() string { return r.JSON.raw }
func (r *VersionPage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
