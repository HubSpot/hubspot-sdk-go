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

	"github.com/HubSpot/hubspot-sdk-go/internal/apijson"
	"github.com/HubSpot/hubspot-sdk-go/internal/apiquery"
	"github.com/HubSpot/hubspot-sdk-go/internal/requestconfig"
	"github.com/HubSpot/hubspot-sdk-go/option"
	"github.com/HubSpot/hubspot-sdk-go/packages/pagination"
	"github.com/HubSpot/hubspot-sdk-go/packages/param"
	"github.com/HubSpot/hubspot-sdk-go/packages/respjson"
	"github.com/HubSpot/hubspot-sdk-go/shared"
)

// PageService contains methods and other services that help with interacting with
// the hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPageService] method instead.
type PageService struct {
	options       []option.RequestOption
	ABTests       PageABTestService
	Batch         PageBatchService
	Folders       PageFolderService
	LandingPages  PageLandingPageService
	MultiLanguage PageMultiLanguageService
	WebsitePages  PageWebsitePageService
}

// NewPageService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewPageService(opts ...option.RequestOption) (r PageService) {
	r = PageService{}
	r.options = opts
	r.ABTests = NewPageABTestService(opts...)
	r.Batch = NewPageBatchService(opts...)
	r.Folders = NewPageFolderService(opts...)
	r.LandingPages = NewPageLandingPageService(opts...)
	r.MultiLanguage = NewPageMultiLanguageService(opts...)
	r.WebsitePages = NewPageWebsitePageService(opts...)
	return
}

// Retrieve a previous version of a landing page, specified by page ID and revision
// ID.
func (r *PageService) GetLandingPageRevision(ctx context.Context, revisionID string, query PageGetLandingPageRevisionParams, opts ...option.RequestOption) (res *PageVersion, err error) {
	opts = slices.Concat(r.options, opts)
	if query.ObjectID == "" {
		err = errors.New("missing required objectId parameter")
		return nil, err
	}
	if revisionID == "" {
		err = errors.New("missing required revisionId parameter")
		return nil, err
	}
	path := fmt.Sprintf("cms/pages/2026-03/landing-pages/%s/revisions/%s", url.PathEscape(query.ObjectID), url.PathEscape(revisionID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Retrieve a previous version of a website page by the revision ID.
func (r *PageService) GetSitePageRevision(ctx context.Context, revisionID string, query PageGetSitePageRevisionParams, opts ...option.RequestOption) (res *PageVersion, err error) {
	opts = slices.Concat(r.options, opts)
	if query.ObjectID == "" {
		err = errors.New("missing required objectId parameter")
		return nil, err
	}
	if revisionID == "" {
		err = errors.New("missing required revisionId parameter")
		return nil, err
	}
	path := fmt.Sprintf("cms/pages/2026-03/site-pages/%s/revisions/%s", url.PathEscape(query.ObjectID), url.PathEscape(revisionID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Retrieve all the previous versions of a landing page, specified by page ID.
func (r *PageService) ListLandingPageRevisions(ctx context.Context, objectID string, query PageListLandingPageRevisionsParams, opts ...option.RequestOption) (res *pagination.Page[PageVersion], err error) {
	var raw *http.Response
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if objectID == "" {
		err = errors.New("missing required objectId parameter")
		return nil, err
	}
	path := fmt.Sprintf("cms/pages/2026-03/landing-pages/%s/revisions", url.PathEscape(objectID))
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

// Retrieve all the previous versions of a landing page, specified by page ID.
func (r *PageService) ListLandingPageRevisionsAutoPaging(ctx context.Context, objectID string, query PageListLandingPageRevisionsParams, opts ...option.RequestOption) *pagination.PageAutoPager[PageVersion] {
	return pagination.NewPageAutoPager(r.ListLandingPageRevisions(ctx, objectID, query, opts...))
}

// Retrieves all the previous versions of a website page, specified by page ID.
func (r *PageService) ListSitePageRevisions(ctx context.Context, objectID string, query PageListSitePageRevisionsParams, opts ...option.RequestOption) (res *pagination.Page[PageVersion], err error) {
	var raw *http.Response
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if objectID == "" {
		err = errors.New("missing required objectId parameter")
		return nil, err
	}
	path := fmt.Sprintf("cms/pages/2026-03/site-pages/%s/revisions", url.PathEscape(objectID))
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

// Retrieves all the previous versions of a website page, specified by page ID.
func (r *PageService) ListSitePageRevisionsAutoPaging(ctx context.Context, objectID string, query PageListSitePageRevisionsParams, opts ...option.RequestOption) *pagination.PageAutoPager[PageVersion] {
	return pagination.NewPageAutoPager(r.ListSitePageRevisions(ctx, objectID, query, opts...))
}

// Discards any edits and resets the draft to match the live version.
func (r *PageService) ResetSitePageDraft(ctx context.Context, objectID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if objectID == "" {
		err = errors.New("missing required objectId parameter")
		return err
	}
	path := fmt.Sprintf("cms/pages/2026-03/site-pages/%s/draft/reset", url.PathEscape(objectID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, nil, opts...)
	return err
}

// Restores a previous version of a landing page, specified by page ID and revision
// ID.
func (r *PageService) RestoreLandingPageRevision(ctx context.Context, revisionID string, body PageRestoreLandingPageRevisionParams, opts ...option.RequestOption) (res *PageData, err error) {
	opts = slices.Concat(r.options, opts)
	if body.ObjectID == "" {
		err = errors.New("missing required objectId parameter")
		return nil, err
	}
	if revisionID == "" {
		err = errors.New("missing required revisionId parameter")
		return nil, err
	}
	path := fmt.Sprintf("cms/pages/2026-03/landing-pages/%s/revisions/%s/restore", url.PathEscape(body.ObjectID), url.PathEscape(revisionID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

// Specify a previous version of a landing page to set as the page draft.
func (r *PageService) RestoreLandingPageRevisionToDraft(ctx context.Context, revisionID int64, body PageRestoreLandingPageRevisionToDraftParams, opts ...option.RequestOption) (res *PageData, err error) {
	opts = slices.Concat(r.options, opts)
	if body.ObjectID == "" {
		err = errors.New("missing required objectId parameter")
		return nil, err
	}
	path := fmt.Sprintf("cms/pages/2026-03/landing-pages/%s/revisions/%v/restore-to-draft", url.PathEscape(body.ObjectID), revisionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

// Restores a website page to a previous version, specified by page ID and version
// ID.
func (r *PageService) RestoreSitePageRevision(ctx context.Context, revisionID string, body PageRestoreSitePageRevisionParams, opts ...option.RequestOption) (res *PageData, err error) {
	opts = slices.Concat(r.options, opts)
	if body.ObjectID == "" {
		err = errors.New("missing required objectId parameter")
		return nil, err
	}
	if revisionID == "" {
		err = errors.New("missing required revisionId parameter")
		return nil, err
	}
	path := fmt.Sprintf("cms/pages/2026-03/site-pages/%s/revisions/%s/restore", url.PathEscape(body.ObjectID), url.PathEscape(revisionID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

// Takes a specified version of a website page and sets it as the new draft version
// of the page.
func (r *PageService) RestoreSitePageRevisionToDraft(ctx context.Context, revisionID int64, body PageRestoreSitePageRevisionToDraftParams, opts ...option.RequestOption) (res *PageData, err error) {
	opts = slices.Concat(r.options, opts)
	if body.ObjectID == "" {
		err = errors.New("missing required objectId parameter")
		return nil, err
	}
	path := fmt.Sprintf("cms/pages/2026-03/site-pages/%s/revisions/%v/restore-to-draft", url.PathEscape(body.ObjectID), revisionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

// The properties AbTestID, WinnerID are required.
type AbTestEndRequestVNextParam struct {
	// ID of the test to end.
	AbTestID string `json:"abTestId" api:"required"`
	// ID of the object to designate as the test winner.
	WinnerID string `json:"winnerId" api:"required"`
	paramObj
}

func (r AbTestEndRequestVNextParam) MarshalJSON() (data []byte, err error) {
	type shadow AbTestEndRequestVNextParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AbTestEndRequestVNextParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties AbTestID, VariationID are required.
type AbTestRerunRequestVNextParam struct {
	// ID of the test to rerun.
	AbTestID string `json:"abTestId" api:"required"`
	// ID of the object to reactivate as a test variation.
	VariationID string `json:"variationId" api:"required"`
	paramObj
}

func (r AbTestRerunRequestVNextParam) MarshalJSON() (data []byte, err error) {
	type shadow AbTestRerunRequestVNextParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AbTestRerunRequestVNextParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Inputs is required.
type BatchInputContentFolderParam struct {
	// Content folders to input.
	Inputs []ContentFolderParam `json:"inputs,omitzero" api:"required"`
	paramObj
}

func (r BatchInputContentFolderParam) MarshalJSON() (data []byte, err error) {
	type shadow BatchInputContentFolderParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BatchInputContentFolderParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Inputs is required.
type BatchInputPageParam struct {
	// Pages to input.
	Inputs []PageDataParam `json:"inputs,omitzero" api:"required"`
	paramObj
}

func (r BatchInputPageParam) MarshalJSON() (data []byte, err error) {
	type shadow BatchInputPageParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BatchInputPageParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BatchResponseContentFolder struct {
	// Time of batch operation completion.
	CompletedAt time.Time `json:"completedAt" api:"required" format:"date-time"`
	// Results of batch operation.
	Results []ContentFolder `json:"results" api:"required"`
	// Time of batch operation start.
	StartedAt time.Time `json:"startedAt" api:"required" format:"date-time"`
	// Status of batch operation.
	//
	// Any of "CANCELED", "COMPLETE", "PENDING", "PROCESSING".
	Status BatchResponseContentFolderStatus `json:"status" api:"required"`
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
func (r BatchResponseContentFolder) RawJSON() string { return r.JSON.raw }
func (r *BatchResponseContentFolder) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Status of batch operation.
type BatchResponseContentFolderStatus string

const (
	BatchResponseContentFolderStatusCanceled   BatchResponseContentFolderStatus = "CANCELED"
	BatchResponseContentFolderStatusComplete   BatchResponseContentFolderStatus = "COMPLETE"
	BatchResponseContentFolderStatusPending    BatchResponseContentFolderStatus = "PENDING"
	BatchResponseContentFolderStatusProcessing BatchResponseContentFolderStatus = "PROCESSING"
)

type BatchResponsePage struct {
	// Time of batch operation completion.
	CompletedAt time.Time `json:"completedAt" api:"required" format:"date-time"`
	// Results of batch operation.
	Results []PageData `json:"results" api:"required"`
	// Time of batch operation start.
	StartedAt time.Time `json:"startedAt" api:"required" format:"date-time"`
	// Status of batch operation.
	//
	// Any of "CANCELED", "COMPLETE", "PENDING", "PROCESSING".
	Status BatchResponsePageStatus `json:"status" api:"required"`
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
func (r BatchResponsePage) RawJSON() string { return r.JSON.raw }
func (r *BatchResponsePage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Status of batch operation.
type BatchResponsePageStatus string

const (
	BatchResponsePageStatusCanceled   BatchResponsePageStatus = "CANCELED"
	BatchResponsePageStatusComplete   BatchResponsePageStatus = "COMPLETE"
	BatchResponsePageStatusPending    BatchResponsePageStatus = "PENDING"
	BatchResponsePageStatusProcessing BatchResponsePageStatus = "PROCESSING"
)

type CollectionResponseWithTotalContentFolderForwardPaging struct {
	// Collection of content folders.
	Results []ContentFolder `json:"results" api:"required"`
	// Total number of content folders.
	Total  int64                `json:"total" api:"required"`
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

type CollectionResponseWithTotalContentFolderVersion struct {
	Results []ContentFolderVersion `json:"results" api:"required"`
	Total   int64                  `json:"total" api:"required"`
	Paging  shared.Paging          `json:"paging"`
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
func (r CollectionResponseWithTotalContentFolderVersion) RawJSON() string { return r.JSON.raw }
func (r *CollectionResponseWithTotalContentFolderVersion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CollectionResponseWithTotalPageForwardPaging struct {
	// Collection of pages.
	Results []PageData `json:"results" api:"required"`
	// Total number of pages.
	Total  int64                `json:"total" api:"required"`
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

type CollectionResponseWithTotalPageVersion struct {
	Results []PageVersion `json:"results" api:"required"`
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
func (r CollectionResponseWithTotalPageVersion) RawJSON() string { return r.JSON.raw }
func (r *CollectionResponseWithTotalPageVersion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ContentFolder struct {
	// The unique ID of the content folder.
	ID string `json:"id" api:"required"`
	// The type of object this folder applies to. Should always be LANDING_PAGE.
	Category int64 `json:"category" api:"required"`
	// The timestamp indicating when the content folder was created.
	Created time.Time `json:"created" api:"required" format:"date-time"`
	// The timestamp (ISO8601 format) when this content folder was deleted.
	DeletedAt time.Time `json:"deletedAt" api:"required" format:"date-time"`
	// The name of the folder which will show up in the app dashboard
	Name string `json:"name" api:"required"`
	// The ID of the content folder this folder is nested under
	ParentFolderID int64 `json:"parentFolderId" api:"required"`
	// The timestamp indicating when the content folder was last updated.
	Updated time.Time `json:"updated" api:"required" format:"date-time"`
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

// The properties ID, Category, Created, DeletedAt, Name, ParentFolderID, Updated
// are required.
type ContentFolderParam struct {
	// The unique ID of the content folder.
	ID string `json:"id" api:"required"`
	// The type of object this folder applies to. Should always be LANDING_PAGE.
	Category int64 `json:"category" api:"required"`
	// The timestamp indicating when the content folder was created.
	Created time.Time `json:"created" api:"required" format:"date-time"`
	// The timestamp (ISO8601 format) when this content folder was deleted.
	DeletedAt time.Time `json:"deletedAt" api:"required" format:"date-time"`
	// The name of the folder which will show up in the app dashboard
	Name string `json:"name" api:"required"`
	// The ID of the content folder this folder is nested under
	ParentFolderID int64 `json:"parentFolderId" api:"required"`
	// The timestamp indicating when the content folder was last updated.
	Updated time.Time `json:"updated" api:"required" format:"date-time"`
	paramObj
}

func (r ContentFolderParam) MarshalJSON() (data []byte, err error) {
	type shadow ContentFolderParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ContentFolderParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ContentFolderVersion struct {
	ID        string             `json:"id" api:"required"`
	Object    ContentFolder      `json:"object" api:"required"`
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
func (r ContentFolderVersion) RawJSON() string { return r.JSON.raw }
func (r *ContentFolderVersion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property ID is required.
type ContentLanguageCloneRequestVNextParam struct {
	// ID of content to clone.
	ID string `json:"id" api:"required"`
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

type PageData struct {
	// The unique ID of the page.
	ID string `json:"id" api:"required"`
	// The status of the AB test associated with this page, if applicable
	//
	// Any of "automated_loser_variant", "automated_master", "automated_variant",
	// "loser_variant", "mab_master", "mab_variant", "master", "variant".
	AbStatus PageDataAbStatus `json:"abStatus" api:"required"`
	// The ID of the AB test associated with this page, if applicable
	AbTestID string `json:"abTestId" api:"required"`
	// The timestamp (ISO8601 format) when this page was deleted.
	ArchivedAt time.Time `json:"archivedAt" api:"required" format:"date-time"`
	// If True, the page will not show up in your dashboard, although the page could
	// still be live.
	ArchivedInDashboard bool `json:"archivedInDashboard" api:"required"`
	// List of stylesheets to attach to this page. These stylesheets are attached to
	// just this page. Order of precedence is bottom to top, just like in the HTML.
	AttachedStylesheets []map[string]any `json:"attachedStylesheets" api:"required"`
	// The name of the user that updated this page.
	AuthorName string `json:"authorName" api:"required"`
	// The GUID of the marketing campaign this page is a part of.
	Campaign string `json:"campaign" api:"required"`
	// ID of the type of object this is. Should always .
	CategoryID int64 `json:"categoryId" api:"required"`
	// The unique identifier for the content group associated with the page.
	ContentGroupID string `json:"contentGroupId" api:"required"`
	// An ENUM descibing the type of this object. Should be either LANDING_PAGE or
	// SITE_PAGE.
	//
	// Any of "0", "1", "10", "11", "12", "13", "14", "15", "16", "17", "18", "19",
	// "2", "20", "21", "22", "3", "4", "5", "6", "7", "8", "9".
	ContentTypeCategory PageDataContentTypeCategory `json:"contentTypeCategory" api:"required"`
	// The timestamp indicating when the page was created.
	Created time.Time `json:"created" api:"required" format:"date-time"`
	// The ID of the user that created this page.
	CreatedByID string `json:"createdById" api:"required"`
	// Indicates whether the page is currently published.
	CurrentlyPublished bool `json:"currentlyPublished" api:"required"`
	// A generated ENUM descibing the current state of this page.
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
	CurrentState PageDataCurrentState `json:"currentState" api:"required"`
	// The domain this page will resolve to. If null, the page will default to the
	// primary domain for this content type.
	Domain string `json:"domain" api:"required"`
	// The identifier for the data source used by the dynamic page.
	DynamicPageDataSourceID string `json:"dynamicPageDataSourceId" api:"required"`
	// The type of data source used by the dynamic page.
	DynamicPageDataSourceType int64 `json:"dynamicPageDataSourceType" api:"required"`
	// The ID of the HubDB table this page references, if applicable
	DynamicPageHubDBTableID string `json:"dynamicPageHubDbTableId" api:"required"`
	// Boolean to determine whether or not the styles from the template should be
	// applied.
	EnableDomainStylesheets bool `json:"enableDomainStylesheets" api:"required"`
	// Boolean to determine whether or not the styles from the template should be
	// applied.
	EnableLayoutStylesheets bool `json:"enableLayoutStylesheets" api:"required"`
	// The featuredImage of this page.
	FeaturedImage string `json:"featuredImage" api:"required"`
	// Alt Text of the featuredImage.
	FeaturedImageAltText string `json:"featuredImageAltText" api:"required"`
	// The ID of the associated folder this landing page is organized under in the app
	// dashboard.
	FolderID string `json:"folderId" api:"required"`
	// Custom HTML for embed codes, javascript that should be placed before the </body>
	// tag of the page.
	FooterHTML string `json:"footerHtml" api:"required"`
	// Custom HTML for embed codes, javascript, etc. that goes in the <head> tag of the
	// page.
	HeadHTML string `json:"headHtml" api:"required"`
	// The html title of this page.
	HTMLTitle string `json:"htmlTitle" api:"required"`
	// Boolean to determine whether or not the Primary CSS Files should be applied.
	IncludeDefaultCustomCss bool `json:"includeDefaultCustomCss" api:"required"`
	// The explicitly defined ISO 639 language code of the page. If null, the page will
	// default to the language of the Domain.
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
	Language PageDataLanguage `json:"language" api:"required"`
	// A structure detailing the layout sections of the page.
	LayoutSections map[string]LayoutSection `json:"layoutSections" api:"required"`
	// Optional override to set the URL to be used in the rel=canonical link tag on the
	// page.
	LinkRelCanonicalURL string `json:"linkRelCanonicalUrl" api:"required"`
	// The ID of the MAB test (or dynamic test) associated with this page, if
	// applicable
	MabExperimentID string `json:"mabExperimentId" api:"required"`
	// A description that goes in <meta> tag on the page.
	MetaDescription string `json:"metaDescription" api:"required"`
	// The internal name of the page.
	Name string `json:"name" api:"required"`
	// The date at which this page should expire and begin redirecting to another url
	// or page.
	PageExpiryDate int64 `json:"pageExpiryDate" api:"required"`
	// Boolean describing if the page expiration feature is enabled for this page
	PageExpiryEnabled bool `json:"pageExpiryEnabled" api:"required"`
	// The ID of another page this page's url should redirect to once this page
	// expires. Should only set this or pageExpiryRedirectUrl.
	PageExpiryRedirectID int64 `json:"pageExpiryRedirectId" api:"required"`
	// The URL this page's url should redirect to once this page expires. Should only
	// set this or pageExpiryRedirectId.
	PageExpiryRedirectURL string `json:"pageExpiryRedirectUrl" api:"required"`
	// A generated Boolean describing whether or not this page is currently expired and
	// being redirected.
	PageRedirected bool `json:"pageRedirected" api:"required"`
	// Set this to create a password protected page. Entering the password will be
	// required to view the page.
	Password string `json:"password" api:"required"`
	// Rules for require member registration to access private content.
	PublicAccessRules []PublicAccessRule `json:"publicAccessRules" api:"required"`
	// Boolean to determine whether or not to respect publicAccessRules.
	PublicAccessRulesEnabled bool `json:"publicAccessRulesEnabled" api:"required"`
	// The date (ISO8601 format) the page is to be published at.
	PublishDate time.Time `json:"publishDate" api:"required" format:"date-time"`
	// Set this to true if you want to be published immediately when the schedule
	// publish endpoint is called, and to ignore the publish_date setting.
	PublishImmediately bool `json:"publishImmediately" api:"required"`
	// The path of the this page. This field is appended to the domain to construct the
	// url of this page.
	Slug string `json:"slug" api:"required"`
	// An ENUM descibing the current state of this page.
	State string `json:"state" api:"required"`
	// Details the type of page this is. Should always be landing_page or site_page
	Subcategory string `json:"subcategory" api:"required"`
	// String detailing the path of the template used for this page.
	TemplatePath string `json:"templatePath" api:"required"`
	// A collection of settings specific to the theme applied to the page.
	ThemeSettingsValues map[string]any `json:"themeSettingsValues" api:"required"`
	// ID of the primary page this object was translated from.
	TranslatedFromID string `json:"translatedFromId" api:"required"`
	// A map of translations for the page, each associated with a specific language
	// variation.
	Translations map[string]ContentLanguageVariation `json:"translations" api:"required"`
	// The timestamp indicating when the page was last updated.
	Updated time.Time `json:"updated" api:"required" format:"date-time"`
	// The ID of the user that updated this page.
	UpdatedByID string `json:"updatedById" api:"required"`
	// A generated field representing the URL of this page.
	URL string `json:"url" api:"required"`
	// Boolean to determine if this page should use a featuredImage.
	UseFeaturedImage bool `json:"useFeaturedImage" api:"required"`
	// A data structure containing the data for all the modules inside the containers
	// for this page. This will only be populated if the page has widget containers.
	WidgetContainers map[string]any `json:"widgetContainers" api:"required"`
	// A data structure containing the data for all the modules for this page.
	Widgets map[string]any `json:"widgets" api:"required"`
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
func (r PageData) RawJSON() string { return r.JSON.raw }
func (r *PageData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this PageData to a PageDataParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// PageDataParam.Overrides()
func (r PageData) ToParam() PageDataParam {
	return param.Override[PageDataParam](json.RawMessage(r.RawJSON()))
}

// The status of the AB test associated with this page, if applicable
type PageDataAbStatus string

const (
	PageDataAbStatusAutomatedLoserVariant PageDataAbStatus = "automated_loser_variant"
	PageDataAbStatusAutomatedMaster       PageDataAbStatus = "automated_master"
	PageDataAbStatusAutomatedVariant      PageDataAbStatus = "automated_variant"
	PageDataAbStatusLoserVariant          PageDataAbStatus = "loser_variant"
	PageDataAbStatusMabMaster             PageDataAbStatus = "mab_master"
	PageDataAbStatusMabVariant            PageDataAbStatus = "mab_variant"
	PageDataAbStatusMaster                PageDataAbStatus = "master"
	PageDataAbStatusVariant               PageDataAbStatus = "variant"
)

// An ENUM descibing the type of this object. Should be either LANDING_PAGE or
// SITE_PAGE.
type PageDataContentTypeCategory string

const (
	PageDataContentTypeCategory0  PageDataContentTypeCategory = "0"
	PageDataContentTypeCategory1  PageDataContentTypeCategory = "1"
	PageDataContentTypeCategory10 PageDataContentTypeCategory = "10"
	PageDataContentTypeCategory11 PageDataContentTypeCategory = "11"
	PageDataContentTypeCategory12 PageDataContentTypeCategory = "12"
	PageDataContentTypeCategory13 PageDataContentTypeCategory = "13"
	PageDataContentTypeCategory14 PageDataContentTypeCategory = "14"
	PageDataContentTypeCategory15 PageDataContentTypeCategory = "15"
	PageDataContentTypeCategory16 PageDataContentTypeCategory = "16"
	PageDataContentTypeCategory17 PageDataContentTypeCategory = "17"
	PageDataContentTypeCategory18 PageDataContentTypeCategory = "18"
	PageDataContentTypeCategory19 PageDataContentTypeCategory = "19"
	PageDataContentTypeCategory2  PageDataContentTypeCategory = "2"
	PageDataContentTypeCategory20 PageDataContentTypeCategory = "20"
	PageDataContentTypeCategory21 PageDataContentTypeCategory = "21"
	PageDataContentTypeCategory22 PageDataContentTypeCategory = "22"
	PageDataContentTypeCategory3  PageDataContentTypeCategory = "3"
	PageDataContentTypeCategory4  PageDataContentTypeCategory = "4"
	PageDataContentTypeCategory5  PageDataContentTypeCategory = "5"
	PageDataContentTypeCategory6  PageDataContentTypeCategory = "6"
	PageDataContentTypeCategory7  PageDataContentTypeCategory = "7"
	PageDataContentTypeCategory8  PageDataContentTypeCategory = "8"
	PageDataContentTypeCategory9  PageDataContentTypeCategory = "9"
)

// A generated ENUM descibing the current state of this page.
type PageDataCurrentState string

const (
	PageDataCurrentStateAgentGenerated          PageDataCurrentState = "AGENT_GENERATED"
	PageDataCurrentStateAutomated               PageDataCurrentState = "AUTOMATED"
	PageDataCurrentStateAutomatedAb             PageDataCurrentState = "AUTOMATED_AB"
	PageDataCurrentStateAutomatedAbVariant      PageDataCurrentState = "AUTOMATED_AB_VARIANT"
	PageDataCurrentStateAutomatedDraft          PageDataCurrentState = "AUTOMATED_DRAFT"
	PageDataCurrentStateAutomatedDraftAb        PageDataCurrentState = "AUTOMATED_DRAFT_AB"
	PageDataCurrentStateAutomatedDraftAbvariant PageDataCurrentState = "AUTOMATED_DRAFT_ABVARIANT"
	PageDataCurrentStateAutomatedForForm        PageDataCurrentState = "AUTOMATED_FOR_FORM"
	PageDataCurrentStateAutomatedForFormBuffer  PageDataCurrentState = "AUTOMATED_FOR_FORM_BUFFER"
	PageDataCurrentStateAutomatedForFormDraft   PageDataCurrentState = "AUTOMATED_FOR_FORM_DRAFT"
	PageDataCurrentStateAutomatedForFormLegacy  PageDataCurrentState = "AUTOMATED_FOR_FORM_LEGACY"
	PageDataCurrentStateAutomatedLoserAbvariant PageDataCurrentState = "AUTOMATED_LOSER_ABVARIANT"
	PageDataCurrentStateAutomatedSending        PageDataCurrentState = "AUTOMATED_SENDING"
	PageDataCurrentStateBlogEmailDraft          PageDataCurrentState = "BLOG_EMAIL_DRAFT"
	PageDataCurrentStateBlogEmailPublished      PageDataCurrentState = "BLOG_EMAIL_PUBLISHED"
	PageDataCurrentStateDraft                   PageDataCurrentState = "DRAFT"
	PageDataCurrentStateDraftAb                 PageDataCurrentState = "DRAFT_AB"
	PageDataCurrentStateDraftAbVariant          PageDataCurrentState = "DRAFT_AB_VARIANT"
	PageDataCurrentStateError                   PageDataCurrentState = "ERROR"
	PageDataCurrentStateLoserAbVariant          PageDataCurrentState = "LOSER_AB_VARIANT"
	PageDataCurrentStatePageStub                PageDataCurrentState = "PAGE_STUB"
	PageDataCurrentStatePreProcessing           PageDataCurrentState = "PRE_PROCESSING"
	PageDataCurrentStateProcessing              PageDataCurrentState = "PROCESSING"
	PageDataCurrentStatePublished               PageDataCurrentState = "PUBLISHED"
	PageDataCurrentStatePublishedAb             PageDataCurrentState = "PUBLISHED_AB"
	PageDataCurrentStatePublishedAbVariant      PageDataCurrentState = "PUBLISHED_AB_VARIANT"
	PageDataCurrentStatePublishedOrScheduled    PageDataCurrentState = "PUBLISHED_OR_SCHEDULED"
	PageDataCurrentStateRssToEmailDraft         PageDataCurrentState = "RSS_TO_EMAIL_DRAFT"
	PageDataCurrentStateRssToEmailPublished     PageDataCurrentState = "RSS_TO_EMAIL_PUBLISHED"
	PageDataCurrentStateScheduled               PageDataCurrentState = "SCHEDULED"
	PageDataCurrentStateScheduledAb             PageDataCurrentState = "SCHEDULED_AB"
	PageDataCurrentStateScheduledOrPublished    PageDataCurrentState = "SCHEDULED_OR_PUBLISHED"
)

// The explicitly defined ISO 639 language code of the page. If null, the page will
// default to the language of the Domain.
type PageDataLanguage string

const (
	PageDataLanguageAa     PageDataLanguage = "aa"
	PageDataLanguageAb     PageDataLanguage = "ab"
	PageDataLanguageAe     PageDataLanguage = "ae"
	PageDataLanguageAf     PageDataLanguage = "af"
	PageDataLanguageAfNa   PageDataLanguage = "af-na"
	PageDataLanguageAfZa   PageDataLanguage = "af-za"
	PageDataLanguageAgq    PageDataLanguage = "agq"
	PageDataLanguageAgqCm  PageDataLanguage = "agq-cm"
	PageDataLanguageAk     PageDataLanguage = "ak"
	PageDataLanguageAkGh   PageDataLanguage = "ak-gh"
	PageDataLanguageAm     PageDataLanguage = "am"
	PageDataLanguageAmEt   PageDataLanguage = "am-et"
	PageDataLanguageAn     PageDataLanguage = "an"
	PageDataLanguageAnn    PageDataLanguage = "ann"
	PageDataLanguageAnnNg  PageDataLanguage = "ann-ng"
	PageDataLanguageAr     PageDataLanguage = "ar"
	PageDataLanguageAr001  PageDataLanguage = "ar-001"
	PageDataLanguageArAe   PageDataLanguage = "ar-ae"
	PageDataLanguageArBh   PageDataLanguage = "ar-bh"
	PageDataLanguageArDj   PageDataLanguage = "ar-dj"
	PageDataLanguageArDz   PageDataLanguage = "ar-dz"
	PageDataLanguageArEg   PageDataLanguage = "ar-eg"
	PageDataLanguageArEh   PageDataLanguage = "ar-eh"
	PageDataLanguageArEr   PageDataLanguage = "ar-er"
	PageDataLanguageArIl   PageDataLanguage = "ar-il"
	PageDataLanguageArIq   PageDataLanguage = "ar-iq"
	PageDataLanguageArJo   PageDataLanguage = "ar-jo"
	PageDataLanguageArKm   PageDataLanguage = "ar-km"
	PageDataLanguageArKw   PageDataLanguage = "ar-kw"
	PageDataLanguageArLb   PageDataLanguage = "ar-lb"
	PageDataLanguageArLy   PageDataLanguage = "ar-ly"
	PageDataLanguageArMa   PageDataLanguage = "ar-ma"
	PageDataLanguageArMr   PageDataLanguage = "ar-mr"
	PageDataLanguageArOm   PageDataLanguage = "ar-om"
	PageDataLanguageArPs   PageDataLanguage = "ar-ps"
	PageDataLanguageArQa   PageDataLanguage = "ar-qa"
	PageDataLanguageArSa   PageDataLanguage = "ar-sa"
	PageDataLanguageArSd   PageDataLanguage = "ar-sd"
	PageDataLanguageArSo   PageDataLanguage = "ar-so"
	PageDataLanguageArSS   PageDataLanguage = "ar-ss"
	PageDataLanguageArSy   PageDataLanguage = "ar-sy"
	PageDataLanguageArTd   PageDataLanguage = "ar-td"
	PageDataLanguageArTn   PageDataLanguage = "ar-tn"
	PageDataLanguageArYe   PageDataLanguage = "ar-ye"
	PageDataLanguageAs     PageDataLanguage = "as"
	PageDataLanguageAsIn   PageDataLanguage = "as-in"
	PageDataLanguageAsa    PageDataLanguage = "asa"
	PageDataLanguageAsaTz  PageDataLanguage = "asa-tz"
	PageDataLanguageAst    PageDataLanguage = "ast"
	PageDataLanguageAstEs  PageDataLanguage = "ast-es"
	PageDataLanguageAv     PageDataLanguage = "av"
	PageDataLanguageAy     PageDataLanguage = "ay"
	PageDataLanguageAz     PageDataLanguage = "az"
	PageDataLanguageAzAz   PageDataLanguage = "az-az"
	PageDataLanguageBa     PageDataLanguage = "ba"
	PageDataLanguageBas    PageDataLanguage = "bas"
	PageDataLanguageBasCm  PageDataLanguage = "bas-cm"
	PageDataLanguageBe     PageDataLanguage = "be"
	PageDataLanguageBeBy   PageDataLanguage = "be-by"
	PageDataLanguageBem    PageDataLanguage = "bem"
	PageDataLanguageBemZm  PageDataLanguage = "bem-zm"
	PageDataLanguageBez    PageDataLanguage = "bez"
	PageDataLanguageBezTz  PageDataLanguage = "bez-tz"
	PageDataLanguageBg     PageDataLanguage = "bg"
	PageDataLanguageBgBg   PageDataLanguage = "bg-bg"
	PageDataLanguageBgc    PageDataLanguage = "bgc"
	PageDataLanguageBgcIn  PageDataLanguage = "bgc-in"
	PageDataLanguageBho    PageDataLanguage = "bho"
	PageDataLanguageBhoIn  PageDataLanguage = "bho-in"
	PageDataLanguageBi     PageDataLanguage = "bi"
	PageDataLanguageBm     PageDataLanguage = "bm"
	PageDataLanguageBmMl   PageDataLanguage = "bm-ml"
	PageDataLanguageBn     PageDataLanguage = "bn"
	PageDataLanguageBnBd   PageDataLanguage = "bn-bd"
	PageDataLanguageBnIn   PageDataLanguage = "bn-in"
	PageDataLanguageBo     PageDataLanguage = "bo"
	PageDataLanguageBoCn   PageDataLanguage = "bo-cn"
	PageDataLanguageBoIn   PageDataLanguage = "bo-in"
	PageDataLanguageBr     PageDataLanguage = "br"
	PageDataLanguageBrFr   PageDataLanguage = "br-fr"
	PageDataLanguageBrx    PageDataLanguage = "brx"
	PageDataLanguageBrxIn  PageDataLanguage = "brx-in"
	PageDataLanguageBs     PageDataLanguage = "bs"
	PageDataLanguageBsBa   PageDataLanguage = "bs-ba"
	PageDataLanguageCa     PageDataLanguage = "ca"
	PageDataLanguageCaAd   PageDataLanguage = "ca-ad"
	PageDataLanguageCaEs   PageDataLanguage = "ca-es"
	PageDataLanguageCaFr   PageDataLanguage = "ca-fr"
	PageDataLanguageCaIt   PageDataLanguage = "ca-it"
	PageDataLanguageCcp    PageDataLanguage = "ccp"
	PageDataLanguageCcpBd  PageDataLanguage = "ccp-bd"
	PageDataLanguageCcpIn  PageDataLanguage = "ccp-in"
	PageDataLanguageCe     PageDataLanguage = "ce"
	PageDataLanguageCeRu   PageDataLanguage = "ce-ru"
	PageDataLanguageCeb    PageDataLanguage = "ceb"
	PageDataLanguageCebPh  PageDataLanguage = "ceb-ph"
	PageDataLanguageCgg    PageDataLanguage = "cgg"
	PageDataLanguageCggUg  PageDataLanguage = "cgg-ug"
	PageDataLanguageCh     PageDataLanguage = "ch"
	PageDataLanguageChr    PageDataLanguage = "chr"
	PageDataLanguageChrUs  PageDataLanguage = "chr-us"
	PageDataLanguageCkb    PageDataLanguage = "ckb"
	PageDataLanguageCkbIq  PageDataLanguage = "ckb-iq"
	PageDataLanguageCkbIr  PageDataLanguage = "ckb-ir"
	PageDataLanguageCo     PageDataLanguage = "co"
	PageDataLanguageCr     PageDataLanguage = "cr"
	PageDataLanguageCs     PageDataLanguage = "cs"
	PageDataLanguageCsCz   PageDataLanguage = "cs-cz"
	PageDataLanguageCu     PageDataLanguage = "cu"
	PageDataLanguageCuRu   PageDataLanguage = "cu-ru"
	PageDataLanguageCv     PageDataLanguage = "cv"
	PageDataLanguageCvRu   PageDataLanguage = "cv-ru"
	PageDataLanguageCy     PageDataLanguage = "cy"
	PageDataLanguageCyGB   PageDataLanguage = "cy-gb"
	PageDataLanguageDa     PageDataLanguage = "da"
	PageDataLanguageDaDk   PageDataLanguage = "da-dk"
	PageDataLanguageDaGl   PageDataLanguage = "da-gl"
	PageDataLanguageDav    PageDataLanguage = "dav"
	PageDataLanguageDavKe  PageDataLanguage = "dav-ke"
	PageDataLanguageDe     PageDataLanguage = "de"
	PageDataLanguageDeAt   PageDataLanguage = "de-at"
	PageDataLanguageDeBe   PageDataLanguage = "de-be"
	PageDataLanguageDeCh   PageDataLanguage = "de-ch"
	PageDataLanguageDeDe   PageDataLanguage = "de-de"
	PageDataLanguageDeGr   PageDataLanguage = "de-gr"
	PageDataLanguageDeIt   PageDataLanguage = "de-it"
	PageDataLanguageDeLi   PageDataLanguage = "de-li"
	PageDataLanguageDeLu   PageDataLanguage = "de-lu"
	PageDataLanguageDje    PageDataLanguage = "dje"
	PageDataLanguageDjeNe  PageDataLanguage = "dje-ne"
	PageDataLanguageDoi    PageDataLanguage = "doi"
	PageDataLanguageDoiIn  PageDataLanguage = "doi-in"
	PageDataLanguageDsb    PageDataLanguage = "dsb"
	PageDataLanguageDsbDe  PageDataLanguage = "dsb-de"
	PageDataLanguageDua    PageDataLanguage = "dua"
	PageDataLanguageDuaCm  PageDataLanguage = "dua-cm"
	PageDataLanguageDv     PageDataLanguage = "dv"
	PageDataLanguageDyo    PageDataLanguage = "dyo"
	PageDataLanguageDyoSn  PageDataLanguage = "dyo-sn"
	PageDataLanguageDz     PageDataLanguage = "dz"
	PageDataLanguageDzBt   PageDataLanguage = "dz-bt"
	PageDataLanguageEbu    PageDataLanguage = "ebu"
	PageDataLanguageEbuKe  PageDataLanguage = "ebu-ke"
	PageDataLanguageEe     PageDataLanguage = "ee"
	PageDataLanguageEeGh   PageDataLanguage = "ee-gh"
	PageDataLanguageEeTg   PageDataLanguage = "ee-tg"
	PageDataLanguageEl     PageDataLanguage = "el"
	PageDataLanguageElCy   PageDataLanguage = "el-cy"
	PageDataLanguageElGr   PageDataLanguage = "el-gr"
	PageDataLanguageEn     PageDataLanguage = "en"
	PageDataLanguageEn001  PageDataLanguage = "en-001"
	PageDataLanguageEn150  PageDataLanguage = "en-150"
	PageDataLanguageEnAe   PageDataLanguage = "en-ae"
	PageDataLanguageEnAg   PageDataLanguage = "en-ag"
	PageDataLanguageEnAI   PageDataLanguage = "en-ai"
	PageDataLanguageEnAs   PageDataLanguage = "en-as"
	PageDataLanguageEnAt   PageDataLanguage = "en-at"
	PageDataLanguageEnAu   PageDataLanguage = "en-au"
	PageDataLanguageEnBb   PageDataLanguage = "en-bb"
	PageDataLanguageEnBe   PageDataLanguage = "en-be"
	PageDataLanguageEnBi   PageDataLanguage = "en-bi"
	PageDataLanguageEnBm   PageDataLanguage = "en-bm"
	PageDataLanguageEnBs   PageDataLanguage = "en-bs"
	PageDataLanguageEnBw   PageDataLanguage = "en-bw"
	PageDataLanguageEnBz   PageDataLanguage = "en-bz"
	PageDataLanguageEnCa   PageDataLanguage = "en-ca"
	PageDataLanguageEnCc   PageDataLanguage = "en-cc"
	PageDataLanguageEnCh   PageDataLanguage = "en-ch"
	PageDataLanguageEnCk   PageDataLanguage = "en-ck"
	PageDataLanguageEnCm   PageDataLanguage = "en-cm"
	PageDataLanguageEnCn   PageDataLanguage = "en-cn"
	PageDataLanguageEnCx   PageDataLanguage = "en-cx"
	PageDataLanguageEnCy   PageDataLanguage = "en-cy"
	PageDataLanguageEnDe   PageDataLanguage = "en-de"
	PageDataLanguageEnDg   PageDataLanguage = "en-dg"
	PageDataLanguageEnDk   PageDataLanguage = "en-dk"
	PageDataLanguageEnDm   PageDataLanguage = "en-dm"
	PageDataLanguageEnEe   PageDataLanguage = "en-ee"
	PageDataLanguageEnEg   PageDataLanguage = "en-eg"
	PageDataLanguageEnEr   PageDataLanguage = "en-er"
	PageDataLanguageEnEs   PageDataLanguage = "en-es"
	PageDataLanguageEnFi   PageDataLanguage = "en-fi"
	PageDataLanguageEnFj   PageDataLanguage = "en-fj"
	PageDataLanguageEnFk   PageDataLanguage = "en-fk"
	PageDataLanguageEnFm   PageDataLanguage = "en-fm"
	PageDataLanguageEnFr   PageDataLanguage = "en-fr"
	PageDataLanguageEnGB   PageDataLanguage = "en-gb"
	PageDataLanguageEnGd   PageDataLanguage = "en-gd"
	PageDataLanguageEnGg   PageDataLanguage = "en-gg"
	PageDataLanguageEnGh   PageDataLanguage = "en-gh"
	PageDataLanguageEnGi   PageDataLanguage = "en-gi"
	PageDataLanguageEnGm   PageDataLanguage = "en-gm"
	PageDataLanguageEnGu   PageDataLanguage = "en-gu"
	PageDataLanguageEnGy   PageDataLanguage = "en-gy"
	PageDataLanguageEnHk   PageDataLanguage = "en-hk"
	PageDataLanguageEnID   PageDataLanguage = "en-id"
	PageDataLanguageEnIe   PageDataLanguage = "en-ie"
	PageDataLanguageEnIl   PageDataLanguage = "en-il"
	PageDataLanguageEnIm   PageDataLanguage = "en-im"
	PageDataLanguageEnIn   PageDataLanguage = "en-in"
	PageDataLanguageEnIo   PageDataLanguage = "en-io"
	PageDataLanguageEnJe   PageDataLanguage = "en-je"
	PageDataLanguageEnJm   PageDataLanguage = "en-jm"
	PageDataLanguageEnKe   PageDataLanguage = "en-ke"
	PageDataLanguageEnKi   PageDataLanguage = "en-ki"
	PageDataLanguageEnKn   PageDataLanguage = "en-kn"
	PageDataLanguageEnKy   PageDataLanguage = "en-ky"
	PageDataLanguageEnLc   PageDataLanguage = "en-lc"
	PageDataLanguageEnLr   PageDataLanguage = "en-lr"
	PageDataLanguageEnLs   PageDataLanguage = "en-ls"
	PageDataLanguageEnLu   PageDataLanguage = "en-lu"
	PageDataLanguageEnMg   PageDataLanguage = "en-mg"
	PageDataLanguageEnMh   PageDataLanguage = "en-mh"
	PageDataLanguageEnMo   PageDataLanguage = "en-mo"
	PageDataLanguageEnMp   PageDataLanguage = "en-mp"
	PageDataLanguageEnMs   PageDataLanguage = "en-ms"
	PageDataLanguageEnMt   PageDataLanguage = "en-mt"
	PageDataLanguageEnMu   PageDataLanguage = "en-mu"
	PageDataLanguageEnMv   PageDataLanguage = "en-mv"
	PageDataLanguageEnMw   PageDataLanguage = "en-mw"
	PageDataLanguageEnMx   PageDataLanguage = "en-mx"
	PageDataLanguageEnMy   PageDataLanguage = "en-my"
	PageDataLanguageEnNa   PageDataLanguage = "en-na"
	PageDataLanguageEnNf   PageDataLanguage = "en-nf"
	PageDataLanguageEnNg   PageDataLanguage = "en-ng"
	PageDataLanguageEnNl   PageDataLanguage = "en-nl"
	PageDataLanguageEnNr   PageDataLanguage = "en-nr"
	PageDataLanguageEnNu   PageDataLanguage = "en-nu"
	PageDataLanguageEnNz   PageDataLanguage = "en-nz"
	PageDataLanguageEnPg   PageDataLanguage = "en-pg"
	PageDataLanguageEnPh   PageDataLanguage = "en-ph"
	PageDataLanguageEnPk   PageDataLanguage = "en-pk"
	PageDataLanguageEnPn   PageDataLanguage = "en-pn"
	PageDataLanguageEnPr   PageDataLanguage = "en-pr"
	PageDataLanguageEnPt   PageDataLanguage = "en-pt"
	PageDataLanguageEnPw   PageDataLanguage = "en-pw"
	PageDataLanguageEnRw   PageDataLanguage = "en-rw"
	PageDataLanguageEnSb   PageDataLanguage = "en-sb"
	PageDataLanguageEnSc   PageDataLanguage = "en-sc"
	PageDataLanguageEnSd   PageDataLanguage = "en-sd"
	PageDataLanguageEnSe   PageDataLanguage = "en-se"
	PageDataLanguageEnSg   PageDataLanguage = "en-sg"
	PageDataLanguageEnSh   PageDataLanguage = "en-sh"
	PageDataLanguageEnSi   PageDataLanguage = "en-si"
	PageDataLanguageEnSl   PageDataLanguage = "en-sl"
	PageDataLanguageEnSS   PageDataLanguage = "en-ss"
	PageDataLanguageEnSx   PageDataLanguage = "en-sx"
	PageDataLanguageEnSz   PageDataLanguage = "en-sz"
	PageDataLanguageEnTc   PageDataLanguage = "en-tc"
	PageDataLanguageEnTh   PageDataLanguage = "en-th"
	PageDataLanguageEnTk   PageDataLanguage = "en-tk"
	PageDataLanguageEnTn   PageDataLanguage = "en-tn"
	PageDataLanguageEnTo   PageDataLanguage = "en-to"
	PageDataLanguageEnTt   PageDataLanguage = "en-tt"
	PageDataLanguageEnTv   PageDataLanguage = "en-tv"
	PageDataLanguageEnTz   PageDataLanguage = "en-tz"
	PageDataLanguageEnUg   PageDataLanguage = "en-ug"
	PageDataLanguageEnUm   PageDataLanguage = "en-um"
	PageDataLanguageEnUs   PageDataLanguage = "en-us"
	PageDataLanguageEnVc   PageDataLanguage = "en-vc"
	PageDataLanguageEnVg   PageDataLanguage = "en-vg"
	PageDataLanguageEnVi   PageDataLanguage = "en-vi"
	PageDataLanguageEnVn   PageDataLanguage = "en-vn"
	PageDataLanguageEnVu   PageDataLanguage = "en-vu"
	PageDataLanguageEnWs   PageDataLanguage = "en-ws"
	PageDataLanguageEnZa   PageDataLanguage = "en-za"
	PageDataLanguageEnZm   PageDataLanguage = "en-zm"
	PageDataLanguageEnZw   PageDataLanguage = "en-zw"
	PageDataLanguageEo     PageDataLanguage = "eo"
	PageDataLanguageEo001  PageDataLanguage = "eo-001"
	PageDataLanguageEs     PageDataLanguage = "es"
	PageDataLanguageEs419  PageDataLanguage = "es-419"
	PageDataLanguageEsAr   PageDataLanguage = "es-ar"
	PageDataLanguageEsBo   PageDataLanguage = "es-bo"
	PageDataLanguageEsBr   PageDataLanguage = "es-br"
	PageDataLanguageEsBz   PageDataLanguage = "es-bz"
	PageDataLanguageEsCl   PageDataLanguage = "es-cl"
	PageDataLanguageEsCo   PageDataLanguage = "es-co"
	PageDataLanguageEsCr   PageDataLanguage = "es-cr"
	PageDataLanguageEsCu   PageDataLanguage = "es-cu"
	PageDataLanguageEsDo   PageDataLanguage = "es-do"
	PageDataLanguageEsEa   PageDataLanguage = "es-ea"
	PageDataLanguageEsEc   PageDataLanguage = "es-ec"
	PageDataLanguageEsEs   PageDataLanguage = "es-es"
	PageDataLanguageEsGq   PageDataLanguage = "es-gq"
	PageDataLanguageEsGt   PageDataLanguage = "es-gt"
	PageDataLanguageEsHn   PageDataLanguage = "es-hn"
	PageDataLanguageEsIc   PageDataLanguage = "es-ic"
	PageDataLanguageEsMx   PageDataLanguage = "es-mx"
	PageDataLanguageEsNi   PageDataLanguage = "es-ni"
	PageDataLanguageEsPa   PageDataLanguage = "es-pa"
	PageDataLanguageEsPe   PageDataLanguage = "es-pe"
	PageDataLanguageEsPh   PageDataLanguage = "es-ph"
	PageDataLanguageEsPr   PageDataLanguage = "es-pr"
	PageDataLanguageEsPy   PageDataLanguage = "es-py"
	PageDataLanguageEsSv   PageDataLanguage = "es-sv"
	PageDataLanguageEsUs   PageDataLanguage = "es-us"
	PageDataLanguageEsUy   PageDataLanguage = "es-uy"
	PageDataLanguageEsVe   PageDataLanguage = "es-ve"
	PageDataLanguageEt     PageDataLanguage = "et"
	PageDataLanguageEtEe   PageDataLanguage = "et-ee"
	PageDataLanguageEu     PageDataLanguage = "eu"
	PageDataLanguageEuEs   PageDataLanguage = "eu-es"
	PageDataLanguageEwo    PageDataLanguage = "ewo"
	PageDataLanguageEwoCm  PageDataLanguage = "ewo-cm"
	PageDataLanguageFa     PageDataLanguage = "fa"
	PageDataLanguageFaAf   PageDataLanguage = "fa-af"
	PageDataLanguageFaIr   PageDataLanguage = "fa-ir"
	PageDataLanguageFf     PageDataLanguage = "ff"
	PageDataLanguageFfBf   PageDataLanguage = "ff-bf"
	PageDataLanguageFfCm   PageDataLanguage = "ff-cm"
	PageDataLanguageFfGh   PageDataLanguage = "ff-gh"
	PageDataLanguageFfGm   PageDataLanguage = "ff-gm"
	PageDataLanguageFfGn   PageDataLanguage = "ff-gn"
	PageDataLanguageFfGw   PageDataLanguage = "ff-gw"
	PageDataLanguageFfLr   PageDataLanguage = "ff-lr"
	PageDataLanguageFfMr   PageDataLanguage = "ff-mr"
	PageDataLanguageFfNe   PageDataLanguage = "ff-ne"
	PageDataLanguageFfNg   PageDataLanguage = "ff-ng"
	PageDataLanguageFfSl   PageDataLanguage = "ff-sl"
	PageDataLanguageFfSn   PageDataLanguage = "ff-sn"
	PageDataLanguageFi     PageDataLanguage = "fi"
	PageDataLanguageFiFi   PageDataLanguage = "fi-fi"
	PageDataLanguageFil    PageDataLanguage = "fil"
	PageDataLanguageFilPh  PageDataLanguage = "fil-ph"
	PageDataLanguageFj     PageDataLanguage = "fj"
	PageDataLanguageFo     PageDataLanguage = "fo"
	PageDataLanguageFoDk   PageDataLanguage = "fo-dk"
	PageDataLanguageFoFo   PageDataLanguage = "fo-fo"
	PageDataLanguageFr     PageDataLanguage = "fr"
	PageDataLanguageFrBe   PageDataLanguage = "fr-be"
	PageDataLanguageFrBf   PageDataLanguage = "fr-bf"
	PageDataLanguageFrBi   PageDataLanguage = "fr-bi"
	PageDataLanguageFrBj   PageDataLanguage = "fr-bj"
	PageDataLanguageFrBl   PageDataLanguage = "fr-bl"
	PageDataLanguageFrCa   PageDataLanguage = "fr-ca"
	PageDataLanguageFrCd   PageDataLanguage = "fr-cd"
	PageDataLanguageFrCf   PageDataLanguage = "fr-cf"
	PageDataLanguageFrCg   PageDataLanguage = "fr-cg"
	PageDataLanguageFrCh   PageDataLanguage = "fr-ch"
	PageDataLanguageFrCi   PageDataLanguage = "fr-ci"
	PageDataLanguageFrCm   PageDataLanguage = "fr-cm"
	PageDataLanguageFrDj   PageDataLanguage = "fr-dj"
	PageDataLanguageFrDz   PageDataLanguage = "fr-dz"
	PageDataLanguageFrFr   PageDataLanguage = "fr-fr"
	PageDataLanguageFrGa   PageDataLanguage = "fr-ga"
	PageDataLanguageFrGf   PageDataLanguage = "fr-gf"
	PageDataLanguageFrGn   PageDataLanguage = "fr-gn"
	PageDataLanguageFrGp   PageDataLanguage = "fr-gp"
	PageDataLanguageFrGq   PageDataLanguage = "fr-gq"
	PageDataLanguageFrHt   PageDataLanguage = "fr-ht"
	PageDataLanguageFrKm   PageDataLanguage = "fr-km"
	PageDataLanguageFrLu   PageDataLanguage = "fr-lu"
	PageDataLanguageFrMa   PageDataLanguage = "fr-ma"
	PageDataLanguageFrMc   PageDataLanguage = "fr-mc"
	PageDataLanguageFrMf   PageDataLanguage = "fr-mf"
	PageDataLanguageFrMg   PageDataLanguage = "fr-mg"
	PageDataLanguageFrMl   PageDataLanguage = "fr-ml"
	PageDataLanguageFrMq   PageDataLanguage = "fr-mq"
	PageDataLanguageFrMr   PageDataLanguage = "fr-mr"
	PageDataLanguageFrMu   PageDataLanguage = "fr-mu"
	PageDataLanguageFrNc   PageDataLanguage = "fr-nc"
	PageDataLanguageFrNe   PageDataLanguage = "fr-ne"
	PageDataLanguageFrPf   PageDataLanguage = "fr-pf"
	PageDataLanguageFrPm   PageDataLanguage = "fr-pm"
	PageDataLanguageFrRe   PageDataLanguage = "fr-re"
	PageDataLanguageFrRw   PageDataLanguage = "fr-rw"
	PageDataLanguageFrSc   PageDataLanguage = "fr-sc"
	PageDataLanguageFrSn   PageDataLanguage = "fr-sn"
	PageDataLanguageFrSy   PageDataLanguage = "fr-sy"
	PageDataLanguageFrTd   PageDataLanguage = "fr-td"
	PageDataLanguageFrTg   PageDataLanguage = "fr-tg"
	PageDataLanguageFrTn   PageDataLanguage = "fr-tn"
	PageDataLanguageFrVu   PageDataLanguage = "fr-vu"
	PageDataLanguageFrWf   PageDataLanguage = "fr-wf"
	PageDataLanguageFrYt   PageDataLanguage = "fr-yt"
	PageDataLanguageFrr    PageDataLanguage = "frr"
	PageDataLanguageFrrDe  PageDataLanguage = "frr-de"
	PageDataLanguageFur    PageDataLanguage = "fur"
	PageDataLanguageFurIt  PageDataLanguage = "fur-it"
	PageDataLanguageFy     PageDataLanguage = "fy"
	PageDataLanguageFyNl   PageDataLanguage = "fy-nl"
	PageDataLanguageGa     PageDataLanguage = "ga"
	PageDataLanguageGaGB   PageDataLanguage = "ga-gb"
	PageDataLanguageGaIe   PageDataLanguage = "ga-ie"
	PageDataLanguageGd     PageDataLanguage = "gd"
	PageDataLanguageGdGB   PageDataLanguage = "gd-gb"
	PageDataLanguageGl     PageDataLanguage = "gl"
	PageDataLanguageGlEs   PageDataLanguage = "gl-es"
	PageDataLanguageGn     PageDataLanguage = "gn"
	PageDataLanguageGsw    PageDataLanguage = "gsw"
	PageDataLanguageGswCh  PageDataLanguage = "gsw-ch"
	PageDataLanguageGswFr  PageDataLanguage = "gsw-fr"
	PageDataLanguageGswLi  PageDataLanguage = "gsw-li"
	PageDataLanguageGu     PageDataLanguage = "gu"
	PageDataLanguageGuIn   PageDataLanguage = "gu-in"
	PageDataLanguageGuz    PageDataLanguage = "guz"
	PageDataLanguageGuzKe  PageDataLanguage = "guz-ke"
	PageDataLanguageGv     PageDataLanguage = "gv"
	PageDataLanguageGvIm   PageDataLanguage = "gv-im"
	PageDataLanguageHa     PageDataLanguage = "ha"
	PageDataLanguageHaGh   PageDataLanguage = "ha-gh"
	PageDataLanguageHaNe   PageDataLanguage = "ha-ne"
	PageDataLanguageHaNg   PageDataLanguage = "ha-ng"
	PageDataLanguageHaw    PageDataLanguage = "haw"
	PageDataLanguageHawUs  PageDataLanguage = "haw-us"
	PageDataLanguageHe     PageDataLanguage = "he"
	PageDataLanguageHeIl   PageDataLanguage = "he-il"
	PageDataLanguageHi     PageDataLanguage = "hi"
	PageDataLanguageHiIn   PageDataLanguage = "hi-in"
	PageDataLanguageHmn    PageDataLanguage = "hmn"
	PageDataLanguageHo     PageDataLanguage = "ho"
	PageDataLanguageHr     PageDataLanguage = "hr"
	PageDataLanguageHrBa   PageDataLanguage = "hr-ba"
	PageDataLanguageHrHr   PageDataLanguage = "hr-hr"
	PageDataLanguageHsb    PageDataLanguage = "hsb"
	PageDataLanguageHsbDe  PageDataLanguage = "hsb-de"
	PageDataLanguageHt     PageDataLanguage = "ht"
	PageDataLanguageHu     PageDataLanguage = "hu"
	PageDataLanguageHuHu   PageDataLanguage = "hu-hu"
	PageDataLanguageHy     PageDataLanguage = "hy"
	PageDataLanguageHyAm   PageDataLanguage = "hy-am"
	PageDataLanguageHz     PageDataLanguage = "hz"
	PageDataLanguageIa     PageDataLanguage = "ia"
	PageDataLanguageIa001  PageDataLanguage = "ia-001"
	PageDataLanguageID     PageDataLanguage = "id"
	PageDataLanguageIDID   PageDataLanguage = "id-id"
	PageDataLanguageIe     PageDataLanguage = "ie"
	PageDataLanguageIg     PageDataLanguage = "ig"
	PageDataLanguageIgNg   PageDataLanguage = "ig-ng"
	PageDataLanguageIi     PageDataLanguage = "ii"
	PageDataLanguageIiCn   PageDataLanguage = "ii-cn"
	PageDataLanguageIk     PageDataLanguage = "ik"
	PageDataLanguageIo     PageDataLanguage = "io"
	PageDataLanguageIs     PageDataLanguage = "is"
	PageDataLanguageIsIs   PageDataLanguage = "is-is"
	PageDataLanguageIt     PageDataLanguage = "it"
	PageDataLanguageItCh   PageDataLanguage = "it-ch"
	PageDataLanguageItIt   PageDataLanguage = "it-it"
	PageDataLanguageItSm   PageDataLanguage = "it-sm"
	PageDataLanguageItVa   PageDataLanguage = "it-va"
	PageDataLanguageIu     PageDataLanguage = "iu"
	PageDataLanguageJa     PageDataLanguage = "ja"
	PageDataLanguageJaJp   PageDataLanguage = "ja-jp"
	PageDataLanguageJgo    PageDataLanguage = "jgo"
	PageDataLanguageJgoCm  PageDataLanguage = "jgo-cm"
	PageDataLanguageJmc    PageDataLanguage = "jmc"
	PageDataLanguageJmcTz  PageDataLanguage = "jmc-tz"
	PageDataLanguageJv     PageDataLanguage = "jv"
	PageDataLanguageJvID   PageDataLanguage = "jv-id"
	PageDataLanguageKa     PageDataLanguage = "ka"
	PageDataLanguageKaGe   PageDataLanguage = "ka-ge"
	PageDataLanguageKab    PageDataLanguage = "kab"
	PageDataLanguageKabDz  PageDataLanguage = "kab-dz"
	PageDataLanguageKam    PageDataLanguage = "kam"
	PageDataLanguageKamKe  PageDataLanguage = "kam-ke"
	PageDataLanguageKar    PageDataLanguage = "kar"
	PageDataLanguageKde    PageDataLanguage = "kde"
	PageDataLanguageKdeTz  PageDataLanguage = "kde-tz"
	PageDataLanguageKea    PageDataLanguage = "kea"
	PageDataLanguageKeaCv  PageDataLanguage = "kea-cv"
	PageDataLanguageKg     PageDataLanguage = "kg"
	PageDataLanguageKgp    PageDataLanguage = "kgp"
	PageDataLanguageKgpBr  PageDataLanguage = "kgp-br"
	PageDataLanguageKh     PageDataLanguage = "kh"
	PageDataLanguageKhq    PageDataLanguage = "khq"
	PageDataLanguageKhqMl  PageDataLanguage = "khq-ml"
	PageDataLanguageKi     PageDataLanguage = "ki"
	PageDataLanguageKiKe   PageDataLanguage = "ki-ke"
	PageDataLanguageKj     PageDataLanguage = "kj"
	PageDataLanguageKk     PageDataLanguage = "kk"
	PageDataLanguageKkKz   PageDataLanguage = "kk-kz"
	PageDataLanguageKkj    PageDataLanguage = "kkj"
	PageDataLanguageKkjCm  PageDataLanguage = "kkj-cm"
	PageDataLanguageKl     PageDataLanguage = "kl"
	PageDataLanguageKlGl   PageDataLanguage = "kl-gl"
	PageDataLanguageKln    PageDataLanguage = "kln"
	PageDataLanguageKlnKe  PageDataLanguage = "kln-ke"
	PageDataLanguageKm     PageDataLanguage = "km"
	PageDataLanguageKmKh   PageDataLanguage = "km-kh"
	PageDataLanguageKn     PageDataLanguage = "kn"
	PageDataLanguageKnIn   PageDataLanguage = "kn-in"
	PageDataLanguageKo     PageDataLanguage = "ko"
	PageDataLanguageKoKp   PageDataLanguage = "ko-kp"
	PageDataLanguageKoKr   PageDataLanguage = "ko-kr"
	PageDataLanguageKok    PageDataLanguage = "kok"
	PageDataLanguageKokIn  PageDataLanguage = "kok-in"
	PageDataLanguageKr     PageDataLanguage = "kr"
	PageDataLanguageKs     PageDataLanguage = "ks"
	PageDataLanguageKsIn   PageDataLanguage = "ks-in"
	PageDataLanguageKsb    PageDataLanguage = "ksb"
	PageDataLanguageKsbTz  PageDataLanguage = "ksb-tz"
	PageDataLanguageKsf    PageDataLanguage = "ksf"
	PageDataLanguageKsfCm  PageDataLanguage = "ksf-cm"
	PageDataLanguageKsh    PageDataLanguage = "ksh"
	PageDataLanguageKshDe  PageDataLanguage = "ksh-de"
	PageDataLanguageKu     PageDataLanguage = "ku"
	PageDataLanguageKuTr   PageDataLanguage = "ku-tr"
	PageDataLanguageKv     PageDataLanguage = "kv"
	PageDataLanguageKw     PageDataLanguage = "kw"
	PageDataLanguageKwGB   PageDataLanguage = "kw-gb"
	PageDataLanguageKy     PageDataLanguage = "ky"
	PageDataLanguageKyKg   PageDataLanguage = "ky-kg"
	PageDataLanguageLa     PageDataLanguage = "la"
	PageDataLanguageLag    PageDataLanguage = "lag"
	PageDataLanguageLagTz  PageDataLanguage = "lag-tz"
	PageDataLanguageLb     PageDataLanguage = "lb"
	PageDataLanguageLbLu   PageDataLanguage = "lb-lu"
	PageDataLanguageLg     PageDataLanguage = "lg"
	PageDataLanguageLgUg   PageDataLanguage = "lg-ug"
	PageDataLanguageLi     PageDataLanguage = "li"
	PageDataLanguageLkt    PageDataLanguage = "lkt"
	PageDataLanguageLktUs  PageDataLanguage = "lkt-us"
	PageDataLanguageLn     PageDataLanguage = "ln"
	PageDataLanguageLnAo   PageDataLanguage = "ln-ao"
	PageDataLanguageLnCd   PageDataLanguage = "ln-cd"
	PageDataLanguageLnCf   PageDataLanguage = "ln-cf"
	PageDataLanguageLnCg   PageDataLanguage = "ln-cg"
	PageDataLanguageLo     PageDataLanguage = "lo"
	PageDataLanguageLoLa   PageDataLanguage = "lo-la"
	PageDataLanguageLrc    PageDataLanguage = "lrc"
	PageDataLanguageLrcIq  PageDataLanguage = "lrc-iq"
	PageDataLanguageLrcIr  PageDataLanguage = "lrc-ir"
	PageDataLanguageLt     PageDataLanguage = "lt"
	PageDataLanguageLtLt   PageDataLanguage = "lt-lt"
	PageDataLanguageLu     PageDataLanguage = "lu"
	PageDataLanguageLuCd   PageDataLanguage = "lu-cd"
	PageDataLanguageLuo    PageDataLanguage = "luo"
	PageDataLanguageLuoKe  PageDataLanguage = "luo-ke"
	PageDataLanguageLuy    PageDataLanguage = "luy"
	PageDataLanguageLuyKe  PageDataLanguage = "luy-ke"
	PageDataLanguageLv     PageDataLanguage = "lv"
	PageDataLanguageLvLv   PageDataLanguage = "lv-lv"
	PageDataLanguageMai    PageDataLanguage = "mai"
	PageDataLanguageMaiIn  PageDataLanguage = "mai-in"
	PageDataLanguageMas    PageDataLanguage = "mas"
	PageDataLanguageMasKe  PageDataLanguage = "mas-ke"
	PageDataLanguageMasTz  PageDataLanguage = "mas-tz"
	PageDataLanguageMdf    PageDataLanguage = "mdf"
	PageDataLanguageMdfRu  PageDataLanguage = "mdf-ru"
	PageDataLanguageMer    PageDataLanguage = "mer"
	PageDataLanguageMerKe  PageDataLanguage = "mer-ke"
	PageDataLanguageMfe    PageDataLanguage = "mfe"
	PageDataLanguageMfeMu  PageDataLanguage = "mfe-mu"
	PageDataLanguageMg     PageDataLanguage = "mg"
	PageDataLanguageMgMg   PageDataLanguage = "mg-mg"
	PageDataLanguageMgh    PageDataLanguage = "mgh"
	PageDataLanguageMghMz  PageDataLanguage = "mgh-mz"
	PageDataLanguageMgo    PageDataLanguage = "mgo"
	PageDataLanguageMgoCm  PageDataLanguage = "mgo-cm"
	PageDataLanguageMh     PageDataLanguage = "mh"
	PageDataLanguageMi     PageDataLanguage = "mi"
	PageDataLanguageMiNz   PageDataLanguage = "mi-nz"
	PageDataLanguageMk     PageDataLanguage = "mk"
	PageDataLanguageMkMk   PageDataLanguage = "mk-mk"
	PageDataLanguageMl     PageDataLanguage = "ml"
	PageDataLanguageMlIn   PageDataLanguage = "ml-in"
	PageDataLanguageMn     PageDataLanguage = "mn"
	PageDataLanguageMnMn   PageDataLanguage = "mn-mn"
	PageDataLanguageMni    PageDataLanguage = "mni"
	PageDataLanguageMniIn  PageDataLanguage = "mni-in"
	PageDataLanguageMr     PageDataLanguage = "mr"
	PageDataLanguageMrIn   PageDataLanguage = "mr-in"
	PageDataLanguageMs     PageDataLanguage = "ms"
	PageDataLanguageMsBn   PageDataLanguage = "ms-bn"
	PageDataLanguageMsID   PageDataLanguage = "ms-id"
	PageDataLanguageMsMy   PageDataLanguage = "ms-my"
	PageDataLanguageMsSg   PageDataLanguage = "ms-sg"
	PageDataLanguageMt     PageDataLanguage = "mt"
	PageDataLanguageMtMt   PageDataLanguage = "mt-mt"
	PageDataLanguageMua    PageDataLanguage = "mua"
	PageDataLanguageMuaCm  PageDataLanguage = "mua-cm"
	PageDataLanguageMy     PageDataLanguage = "my"
	PageDataLanguageMyMm   PageDataLanguage = "my-mm"
	PageDataLanguageMzn    PageDataLanguage = "mzn"
	PageDataLanguageMznIr  PageDataLanguage = "mzn-ir"
	PageDataLanguageNa     PageDataLanguage = "na"
	PageDataLanguageNaq    PageDataLanguage = "naq"
	PageDataLanguageNaqNa  PageDataLanguage = "naq-na"
	PageDataLanguageNb     PageDataLanguage = "nb"
	PageDataLanguageNbNo   PageDataLanguage = "nb-no"
	PageDataLanguageNbSj   PageDataLanguage = "nb-sj"
	PageDataLanguageNd     PageDataLanguage = "nd"
	PageDataLanguageNdZw   PageDataLanguage = "nd-zw"
	PageDataLanguageNds    PageDataLanguage = "nds"
	PageDataLanguageNdsDe  PageDataLanguage = "nds-de"
	PageDataLanguageNdsNl  PageDataLanguage = "nds-nl"
	PageDataLanguageNe     PageDataLanguage = "ne"
	PageDataLanguageNeIn   PageDataLanguage = "ne-in"
	PageDataLanguageNeNp   PageDataLanguage = "ne-np"
	PageDataLanguageNg     PageDataLanguage = "ng"
	PageDataLanguageNl     PageDataLanguage = "nl"
	PageDataLanguageNlAw   PageDataLanguage = "nl-aw"
	PageDataLanguageNlBe   PageDataLanguage = "nl-be"
	PageDataLanguageNlBq   PageDataLanguage = "nl-bq"
	PageDataLanguageNlCh   PageDataLanguage = "nl-ch"
	PageDataLanguageNlCw   PageDataLanguage = "nl-cw"
	PageDataLanguageNlLu   PageDataLanguage = "nl-lu"
	PageDataLanguageNlNl   PageDataLanguage = "nl-nl"
	PageDataLanguageNlSr   PageDataLanguage = "nl-sr"
	PageDataLanguageNlSx   PageDataLanguage = "nl-sx"
	PageDataLanguageNmg    PageDataLanguage = "nmg"
	PageDataLanguageNmgCm  PageDataLanguage = "nmg-cm"
	PageDataLanguageNn     PageDataLanguage = "nn"
	PageDataLanguageNnNo   PageDataLanguage = "nn-no"
	PageDataLanguageNnh    PageDataLanguage = "nnh"
	PageDataLanguageNnhCm  PageDataLanguage = "nnh-cm"
	PageDataLanguageNo     PageDataLanguage = "no"
	PageDataLanguageNoNo   PageDataLanguage = "no-no"
	PageDataLanguageNr     PageDataLanguage = "nr"
	PageDataLanguageNus    PageDataLanguage = "nus"
	PageDataLanguageNusSS  PageDataLanguage = "nus-ss"
	PageDataLanguageNv     PageDataLanguage = "nv"
	PageDataLanguageNy     PageDataLanguage = "ny"
	PageDataLanguageNyn    PageDataLanguage = "nyn"
	PageDataLanguageNynUg  PageDataLanguage = "nyn-ug"
	PageDataLanguageOc     PageDataLanguage = "oc"
	PageDataLanguageOcEs   PageDataLanguage = "oc-es"
	PageDataLanguageOcFr   PageDataLanguage = "oc-fr"
	PageDataLanguageOj     PageDataLanguage = "oj"
	PageDataLanguageOm     PageDataLanguage = "om"
	PageDataLanguageOmEt   PageDataLanguage = "om-et"
	PageDataLanguageOmKe   PageDataLanguage = "om-ke"
	PageDataLanguageOr     PageDataLanguage = "or"
	PageDataLanguageOrIn   PageDataLanguage = "or-in"
	PageDataLanguageOs     PageDataLanguage = "os"
	PageDataLanguageOsGe   PageDataLanguage = "os-ge"
	PageDataLanguageOsRu   PageDataLanguage = "os-ru"
	PageDataLanguagePa     PageDataLanguage = "pa"
	PageDataLanguagePaIn   PageDataLanguage = "pa-in"
	PageDataLanguagePaPk   PageDataLanguage = "pa-pk"
	PageDataLanguagePcm    PageDataLanguage = "pcm"
	PageDataLanguagePcmNg  PageDataLanguage = "pcm-ng"
	PageDataLanguagePi     PageDataLanguage = "pi"
	PageDataLanguagePis    PageDataLanguage = "pis"
	PageDataLanguagePisSb  PageDataLanguage = "pis-sb"
	PageDataLanguagePl     PageDataLanguage = "pl"
	PageDataLanguagePlPl   PageDataLanguage = "pl-pl"
	PageDataLanguagePrg    PageDataLanguage = "prg"
	PageDataLanguagePrg001 PageDataLanguage = "prg-001"
	PageDataLanguagePs     PageDataLanguage = "ps"
	PageDataLanguagePsAf   PageDataLanguage = "ps-af"
	PageDataLanguagePsPk   PageDataLanguage = "ps-pk"
	PageDataLanguagePt     PageDataLanguage = "pt"
	PageDataLanguagePtAo   PageDataLanguage = "pt-ao"
	PageDataLanguagePtBr   PageDataLanguage = "pt-br"
	PageDataLanguagePtCh   PageDataLanguage = "pt-ch"
	PageDataLanguagePtCv   PageDataLanguage = "pt-cv"
	PageDataLanguagePtGq   PageDataLanguage = "pt-gq"
	PageDataLanguagePtGw   PageDataLanguage = "pt-gw"
	PageDataLanguagePtLu   PageDataLanguage = "pt-lu"
	PageDataLanguagePtMo   PageDataLanguage = "pt-mo"
	PageDataLanguagePtMz   PageDataLanguage = "pt-mz"
	PageDataLanguagePtPt   PageDataLanguage = "pt-pt"
	PageDataLanguagePtSt   PageDataLanguage = "pt-st"
	PageDataLanguagePtTl   PageDataLanguage = "pt-tl"
	PageDataLanguageQu     PageDataLanguage = "qu"
	PageDataLanguageQuBo   PageDataLanguage = "qu-bo"
	PageDataLanguageQuEc   PageDataLanguage = "qu-ec"
	PageDataLanguageQuPe   PageDataLanguage = "qu-pe"
	PageDataLanguageRaj    PageDataLanguage = "raj"
	PageDataLanguageRajIn  PageDataLanguage = "raj-in"
	PageDataLanguageRm     PageDataLanguage = "rm"
	PageDataLanguageRmCh   PageDataLanguage = "rm-ch"
	PageDataLanguageRn     PageDataLanguage = "rn"
	PageDataLanguageRnBi   PageDataLanguage = "rn-bi"
	PageDataLanguageRo     PageDataLanguage = "ro"
	PageDataLanguageRoMd   PageDataLanguage = "ro-md"
	PageDataLanguageRoRo   PageDataLanguage = "ro-ro"
	PageDataLanguageRof    PageDataLanguage = "rof"
	PageDataLanguageRofTz  PageDataLanguage = "rof-tz"
	PageDataLanguageRu     PageDataLanguage = "ru"
	PageDataLanguageRuBy   PageDataLanguage = "ru-by"
	PageDataLanguageRuKg   PageDataLanguage = "ru-kg"
	PageDataLanguageRuKz   PageDataLanguage = "ru-kz"
	PageDataLanguageRuMd   PageDataLanguage = "ru-md"
	PageDataLanguageRuRu   PageDataLanguage = "ru-ru"
	PageDataLanguageRuUa   PageDataLanguage = "ru-ua"
	PageDataLanguageRw     PageDataLanguage = "rw"
	PageDataLanguageRwRw   PageDataLanguage = "rw-rw"
	PageDataLanguageRwk    PageDataLanguage = "rwk"
	PageDataLanguageRwkTz  PageDataLanguage = "rwk-tz"
	PageDataLanguageSa     PageDataLanguage = "sa"
	PageDataLanguageSaIn   PageDataLanguage = "sa-in"
	PageDataLanguageSah    PageDataLanguage = "sah"
	PageDataLanguageSahRu  PageDataLanguage = "sah-ru"
	PageDataLanguageSaq    PageDataLanguage = "saq"
	PageDataLanguageSaqKe  PageDataLanguage = "saq-ke"
	PageDataLanguageSat    PageDataLanguage = "sat"
	PageDataLanguageSatIn  PageDataLanguage = "sat-in"
	PageDataLanguageSbp    PageDataLanguage = "sbp"
	PageDataLanguageSbpTz  PageDataLanguage = "sbp-tz"
	PageDataLanguageSc     PageDataLanguage = "sc"
	PageDataLanguageScIt   PageDataLanguage = "sc-it"
	PageDataLanguageSd     PageDataLanguage = "sd"
	PageDataLanguageSdIn   PageDataLanguage = "sd-in"
	PageDataLanguageSdPk   PageDataLanguage = "sd-pk"
	PageDataLanguageSe     PageDataLanguage = "se"
	PageDataLanguageSeFi   PageDataLanguage = "se-fi"
	PageDataLanguageSeNo   PageDataLanguage = "se-no"
	PageDataLanguageSeSe   PageDataLanguage = "se-se"
	PageDataLanguageSeh    PageDataLanguage = "seh"
	PageDataLanguageSehMz  PageDataLanguage = "seh-mz"
	PageDataLanguageSes    PageDataLanguage = "ses"
	PageDataLanguageSesMl  PageDataLanguage = "ses-ml"
	PageDataLanguageSg     PageDataLanguage = "sg"
	PageDataLanguageSgCf   PageDataLanguage = "sg-cf"
	PageDataLanguageShi    PageDataLanguage = "shi"
	PageDataLanguageShiMa  PageDataLanguage = "shi-ma"
	PageDataLanguageSi     PageDataLanguage = "si"
	PageDataLanguageSiLk   PageDataLanguage = "si-lk"
	PageDataLanguageSk     PageDataLanguage = "sk"
	PageDataLanguageSkSk   PageDataLanguage = "sk-sk"
	PageDataLanguageSl     PageDataLanguage = "sl"
	PageDataLanguageSlSi   PageDataLanguage = "sl-si"
	PageDataLanguageSm     PageDataLanguage = "sm"
	PageDataLanguageSmn    PageDataLanguage = "smn"
	PageDataLanguageSmnFi  PageDataLanguage = "smn-fi"
	PageDataLanguageSMS    PageDataLanguage = "sms"
	PageDataLanguageSMSFi  PageDataLanguage = "sms-fi"
	PageDataLanguageSn     PageDataLanguage = "sn"
	PageDataLanguageSnZw   PageDataLanguage = "sn-zw"
	PageDataLanguageSo     PageDataLanguage = "so"
	PageDataLanguageSoDj   PageDataLanguage = "so-dj"
	PageDataLanguageSoEt   PageDataLanguage = "so-et"
	PageDataLanguageSoKe   PageDataLanguage = "so-ke"
	PageDataLanguageSoSo   PageDataLanguage = "so-so"
	PageDataLanguageSq     PageDataLanguage = "sq"
	PageDataLanguageSqAl   PageDataLanguage = "sq-al"
	PageDataLanguageSqMk   PageDataLanguage = "sq-mk"
	PageDataLanguageSqXk   PageDataLanguage = "sq-xk"
	PageDataLanguageSr     PageDataLanguage = "sr"
	PageDataLanguageSrBa   PageDataLanguage = "sr-ba"
	PageDataLanguageSrCs   PageDataLanguage = "sr-cs"
	PageDataLanguageSrMe   PageDataLanguage = "sr-me"
	PageDataLanguageSrRs   PageDataLanguage = "sr-rs"
	PageDataLanguageSrXk   PageDataLanguage = "sr-xk"
	PageDataLanguageSS     PageDataLanguage = "ss"
	PageDataLanguageSt     PageDataLanguage = "st"
	PageDataLanguageSu     PageDataLanguage = "su"
	PageDataLanguageSuID   PageDataLanguage = "su-id"
	PageDataLanguageSv     PageDataLanguage = "sv"
	PageDataLanguageSvAx   PageDataLanguage = "sv-ax"
	PageDataLanguageSvFi   PageDataLanguage = "sv-fi"
	PageDataLanguageSvSe   PageDataLanguage = "sv-se"
	PageDataLanguageSw     PageDataLanguage = "sw"
	PageDataLanguageSwCd   PageDataLanguage = "sw-cd"
	PageDataLanguageSwKe   PageDataLanguage = "sw-ke"
	PageDataLanguageSwTz   PageDataLanguage = "sw-tz"
	PageDataLanguageSwUg   PageDataLanguage = "sw-ug"
	PageDataLanguageSy     PageDataLanguage = "sy"
	PageDataLanguageTa     PageDataLanguage = "ta"
	PageDataLanguageTaIn   PageDataLanguage = "ta-in"
	PageDataLanguageTaLk   PageDataLanguage = "ta-lk"
	PageDataLanguageTaMy   PageDataLanguage = "ta-my"
	PageDataLanguageTaSg   PageDataLanguage = "ta-sg"
	PageDataLanguageTe     PageDataLanguage = "te"
	PageDataLanguageTeIn   PageDataLanguage = "te-in"
	PageDataLanguageTeo    PageDataLanguage = "teo"
	PageDataLanguageTeoKe  PageDataLanguage = "teo-ke"
	PageDataLanguageTeoUg  PageDataLanguage = "teo-ug"
	PageDataLanguageTg     PageDataLanguage = "tg"
	PageDataLanguageTgTj   PageDataLanguage = "tg-tj"
	PageDataLanguageTh     PageDataLanguage = "th"
	PageDataLanguageThTh   PageDataLanguage = "th-th"
	PageDataLanguageTi     PageDataLanguage = "ti"
	PageDataLanguageTiEr   PageDataLanguage = "ti-er"
	PageDataLanguageTiEt   PageDataLanguage = "ti-et"
	PageDataLanguageTk     PageDataLanguage = "tk"
	PageDataLanguageTkTm   PageDataLanguage = "tk-tm"
	PageDataLanguageTl     PageDataLanguage = "tl"
	PageDataLanguageTn     PageDataLanguage = "tn"
	PageDataLanguageTo     PageDataLanguage = "to"
	PageDataLanguageToTo   PageDataLanguage = "to-to"
	PageDataLanguageTok    PageDataLanguage = "tok"
	PageDataLanguageTok001 PageDataLanguage = "tok-001"
	PageDataLanguageTr     PageDataLanguage = "tr"
	PageDataLanguageTrCy   PageDataLanguage = "tr-cy"
	PageDataLanguageTrTr   PageDataLanguage = "tr-tr"
	PageDataLanguageTs     PageDataLanguage = "ts"
	PageDataLanguageTt     PageDataLanguage = "tt"
	PageDataLanguageTtRu   PageDataLanguage = "tt-ru"
	PageDataLanguageTw     PageDataLanguage = "tw"
	PageDataLanguageTwq    PageDataLanguage = "twq"
	PageDataLanguageTwqNe  PageDataLanguage = "twq-ne"
	PageDataLanguageTy     PageDataLanguage = "ty"
	PageDataLanguageTzm    PageDataLanguage = "tzm"
	PageDataLanguageTzmMa  PageDataLanguage = "tzm-ma"
	PageDataLanguageUg     PageDataLanguage = "ug"
	PageDataLanguageUgCn   PageDataLanguage = "ug-cn"
	PageDataLanguageUk     PageDataLanguage = "uk"
	PageDataLanguageUkUa   PageDataLanguage = "uk-ua"
	PageDataLanguageUr     PageDataLanguage = "ur"
	PageDataLanguageUrIn   PageDataLanguage = "ur-in"
	PageDataLanguageUrPk   PageDataLanguage = "ur-pk"
	PageDataLanguageUz     PageDataLanguage = "uz"
	PageDataLanguageUzAf   PageDataLanguage = "uz-af"
	PageDataLanguageUzUz   PageDataLanguage = "uz-uz"
	PageDataLanguageVai    PageDataLanguage = "vai"
	PageDataLanguageVaiLr  PageDataLanguage = "vai-lr"
	PageDataLanguageVe     PageDataLanguage = "ve"
	PageDataLanguageVi     PageDataLanguage = "vi"
	PageDataLanguageViVn   PageDataLanguage = "vi-vn"
	PageDataLanguageVo     PageDataLanguage = "vo"
	PageDataLanguageVo001  PageDataLanguage = "vo-001"
	PageDataLanguageVun    PageDataLanguage = "vun"
	PageDataLanguageVunTz  PageDataLanguage = "vun-tz"
	PageDataLanguageWa     PageDataLanguage = "wa"
	PageDataLanguageWae    PageDataLanguage = "wae"
	PageDataLanguageWaeCh  PageDataLanguage = "wae-ch"
	PageDataLanguageWo     PageDataLanguage = "wo"
	PageDataLanguageWoSn   PageDataLanguage = "wo-sn"
	PageDataLanguageXh     PageDataLanguage = "xh"
	PageDataLanguageXhZa   PageDataLanguage = "xh-za"
	PageDataLanguageXog    PageDataLanguage = "xog"
	PageDataLanguageXogUg  PageDataLanguage = "xog-ug"
	PageDataLanguageYav    PageDataLanguage = "yav"
	PageDataLanguageYavCm  PageDataLanguage = "yav-cm"
	PageDataLanguageYi     PageDataLanguage = "yi"
	PageDataLanguageYi001  PageDataLanguage = "yi-001"
	PageDataLanguageYo     PageDataLanguage = "yo"
	PageDataLanguageYoBj   PageDataLanguage = "yo-bj"
	PageDataLanguageYoNg   PageDataLanguage = "yo-ng"
	PageDataLanguageYrl    PageDataLanguage = "yrl"
	PageDataLanguageYrlBr  PageDataLanguage = "yrl-br"
	PageDataLanguageYrlCo  PageDataLanguage = "yrl-co"
	PageDataLanguageYrlVe  PageDataLanguage = "yrl-ve"
	PageDataLanguageYue    PageDataLanguage = "yue"
	PageDataLanguageYueCn  PageDataLanguage = "yue-cn"
	PageDataLanguageYueHk  PageDataLanguage = "yue-hk"
	PageDataLanguageZa     PageDataLanguage = "za"
	PageDataLanguageZgh    PageDataLanguage = "zgh"
	PageDataLanguageZghMa  PageDataLanguage = "zgh-ma"
	PageDataLanguageZh     PageDataLanguage = "zh"
	PageDataLanguageZhCn   PageDataLanguage = "zh-cn"
	PageDataLanguageZhHans PageDataLanguage = "zh-hans"
	PageDataLanguageZhHant PageDataLanguage = "zh-hant"
	PageDataLanguageZhHk   PageDataLanguage = "zh-hk"
	PageDataLanguageZhMo   PageDataLanguage = "zh-mo"
	PageDataLanguageZhSg   PageDataLanguage = "zh-sg"
	PageDataLanguageZhTw   PageDataLanguage = "zh-tw"
	PageDataLanguageZu     PageDataLanguage = "zu"
	PageDataLanguageZuZa   PageDataLanguage = "zu-za"
)

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
type PageDataParam struct {
	// The unique ID of the page.
	ID string `json:"id" api:"required"`
	// The status of the AB test associated with this page, if applicable
	//
	// Any of "automated_loser_variant", "automated_master", "automated_variant",
	// "loser_variant", "mab_master", "mab_variant", "master", "variant".
	AbStatus PageDataAbStatus `json:"abStatus,omitzero" api:"required"`
	// The ID of the AB test associated with this page, if applicable
	AbTestID string `json:"abTestId" api:"required"`
	// The timestamp (ISO8601 format) when this page was deleted.
	ArchivedAt time.Time `json:"archivedAt" api:"required" format:"date-time"`
	// If True, the page will not show up in your dashboard, although the page could
	// still be live.
	ArchivedInDashboard bool `json:"archivedInDashboard" api:"required"`
	// List of stylesheets to attach to this page. These stylesheets are attached to
	// just this page. Order of precedence is bottom to top, just like in the HTML.
	AttachedStylesheets []map[string]any `json:"attachedStylesheets,omitzero" api:"required"`
	// The name of the user that updated this page.
	AuthorName string `json:"authorName" api:"required"`
	// The GUID of the marketing campaign this page is a part of.
	Campaign string `json:"campaign" api:"required"`
	// ID of the type of object this is. Should always .
	CategoryID int64 `json:"categoryId" api:"required"`
	// The unique identifier for the content group associated with the page.
	ContentGroupID string `json:"contentGroupId" api:"required"`
	// An ENUM descibing the type of this object. Should be either LANDING_PAGE or
	// SITE_PAGE.
	//
	// Any of "0", "1", "10", "11", "12", "13", "14", "15", "16", "17", "18", "19",
	// "2", "20", "21", "22", "3", "4", "5", "6", "7", "8", "9".
	ContentTypeCategory PageDataContentTypeCategory `json:"contentTypeCategory,omitzero" api:"required"`
	// The timestamp indicating when the page was created.
	Created time.Time `json:"created" api:"required" format:"date-time"`
	// The ID of the user that created this page.
	CreatedByID string `json:"createdById" api:"required"`
	// Indicates whether the page is currently published.
	CurrentlyPublished bool `json:"currentlyPublished" api:"required"`
	// A generated ENUM descibing the current state of this page.
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
	CurrentState PageDataCurrentState `json:"currentState,omitzero" api:"required"`
	// The domain this page will resolve to. If null, the page will default to the
	// primary domain for this content type.
	Domain string `json:"domain" api:"required"`
	// The identifier for the data source used by the dynamic page.
	DynamicPageDataSourceID string `json:"dynamicPageDataSourceId" api:"required"`
	// The type of data source used by the dynamic page.
	DynamicPageDataSourceType int64 `json:"dynamicPageDataSourceType" api:"required"`
	// The ID of the HubDB table this page references, if applicable
	DynamicPageHubDBTableID string `json:"dynamicPageHubDbTableId" api:"required"`
	// Boolean to determine whether or not the styles from the template should be
	// applied.
	EnableDomainStylesheets bool `json:"enableDomainStylesheets" api:"required"`
	// Boolean to determine whether or not the styles from the template should be
	// applied.
	EnableLayoutStylesheets bool `json:"enableLayoutStylesheets" api:"required"`
	// The featuredImage of this page.
	FeaturedImage string `json:"featuredImage" api:"required"`
	// Alt Text of the featuredImage.
	FeaturedImageAltText string `json:"featuredImageAltText" api:"required"`
	// The ID of the associated folder this landing page is organized under in the app
	// dashboard.
	FolderID string `json:"folderId" api:"required"`
	// Custom HTML for embed codes, javascript that should be placed before the </body>
	// tag of the page.
	FooterHTML string `json:"footerHtml" api:"required"`
	// Custom HTML for embed codes, javascript, etc. that goes in the <head> tag of the
	// page.
	HeadHTML string `json:"headHtml" api:"required"`
	// The html title of this page.
	HTMLTitle string `json:"htmlTitle" api:"required"`
	// Boolean to determine whether or not the Primary CSS Files should be applied.
	IncludeDefaultCustomCss bool `json:"includeDefaultCustomCss" api:"required"`
	// The explicitly defined ISO 639 language code of the page. If null, the page will
	// default to the language of the Domain.
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
	Language PageDataLanguage `json:"language,omitzero" api:"required"`
	// A structure detailing the layout sections of the page.
	LayoutSections map[string]LayoutSectionParam `json:"layoutSections,omitzero" api:"required"`
	// Optional override to set the URL to be used in the rel=canonical link tag on the
	// page.
	LinkRelCanonicalURL string `json:"linkRelCanonicalUrl" api:"required"`
	// The ID of the MAB test (or dynamic test) associated with this page, if
	// applicable
	MabExperimentID string `json:"mabExperimentId" api:"required"`
	// A description that goes in <meta> tag on the page.
	MetaDescription string `json:"metaDescription" api:"required"`
	// The internal name of the page.
	Name string `json:"name" api:"required"`
	// The date at which this page should expire and begin redirecting to another url
	// or page.
	PageExpiryDate int64 `json:"pageExpiryDate" api:"required"`
	// Boolean describing if the page expiration feature is enabled for this page
	PageExpiryEnabled bool `json:"pageExpiryEnabled" api:"required"`
	// The ID of another page this page's url should redirect to once this page
	// expires. Should only set this or pageExpiryRedirectUrl.
	PageExpiryRedirectID int64 `json:"pageExpiryRedirectId" api:"required"`
	// The URL this page's url should redirect to once this page expires. Should only
	// set this or pageExpiryRedirectId.
	PageExpiryRedirectURL string `json:"pageExpiryRedirectUrl" api:"required"`
	// A generated Boolean describing whether or not this page is currently expired and
	// being redirected.
	PageRedirected bool `json:"pageRedirected" api:"required"`
	// Set this to create a password protected page. Entering the password will be
	// required to view the page.
	Password string `json:"password" api:"required"`
	// Rules for require member registration to access private content.
	PublicAccessRules []PublicAccessRule `json:"publicAccessRules,omitzero" api:"required"`
	// Boolean to determine whether or not to respect publicAccessRules.
	PublicAccessRulesEnabled bool `json:"publicAccessRulesEnabled" api:"required"`
	// The date (ISO8601 format) the page is to be published at.
	PublishDate time.Time `json:"publishDate" api:"required" format:"date-time"`
	// Set this to true if you want to be published immediately when the schedule
	// publish endpoint is called, and to ignore the publish_date setting.
	PublishImmediately bool `json:"publishImmediately" api:"required"`
	// The path of the this page. This field is appended to the domain to construct the
	// url of this page.
	Slug string `json:"slug" api:"required"`
	// An ENUM descibing the current state of this page.
	State string `json:"state" api:"required"`
	// Details the type of page this is. Should always be landing_page or site_page
	Subcategory string `json:"subcategory" api:"required"`
	// String detailing the path of the template used for this page.
	TemplatePath string `json:"templatePath" api:"required"`
	// A collection of settings specific to the theme applied to the page.
	ThemeSettingsValues map[string]any `json:"themeSettingsValues,omitzero" api:"required"`
	// ID of the primary page this object was translated from.
	TranslatedFromID string `json:"translatedFromId" api:"required"`
	// A map of translations for the page, each associated with a specific language
	// variation.
	Translations map[string]ContentLanguageVariationParam `json:"translations,omitzero" api:"required"`
	// The timestamp indicating when the page was last updated.
	Updated time.Time `json:"updated" api:"required" format:"date-time"`
	// The ID of the user that updated this page.
	UpdatedByID string `json:"updatedById" api:"required"`
	// A generated field representing the URL of this page.
	URL string `json:"url" api:"required"`
	// Boolean to determine if this page should use a featuredImage.
	UseFeaturedImage bool `json:"useFeaturedImage" api:"required"`
	// A data structure containing the data for all the modules inside the containers
	// for this page. This will only be populated if the page has widget containers.
	WidgetContainers map[string]any `json:"widgetContainers,omitzero" api:"required"`
	// A data structure containing the data for all the modules for this page.
	Widgets map[string]any `json:"widgets,omitzero" api:"required"`
	paramObj
}

func (r PageDataParam) MarshalJSON() (data []byte, err error) {
	type shadow PageDataParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PageDataParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PageVersion struct {
	ID        string             `json:"id" api:"required"`
	Object    PageData           `json:"object" api:"required"`
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
func (r PageVersion) RawJSON() string { return r.JSON.raw }
func (r *PageVersion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PageGetLandingPageRevisionParams struct {
	ObjectID string `path:"objectId" api:"required" json:"-"`
	paramObj
}

type PageGetSitePageRevisionParams struct {
	ObjectID string `path:"objectId" api:"required" json:"-"`
	paramObj
}

type PageListLandingPageRevisionsParams struct {
	// The paging cursor token of the last successfully read resource will be returned
	// as the `paging.next.after` JSON property of a paged response containing more
	// results.
	After  param.Opt[string] `query:"after,omitzero" json:"-"`
	Before param.Opt[string] `query:"before,omitzero" json:"-"`
	// The maximum number of results to display per page.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [PageListLandingPageRevisionsParams]'s query parameters as
// `url.Values`.
func (r PageListLandingPageRevisionsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type PageListSitePageRevisionsParams struct {
	// The paging cursor token of the last successfully read resource will be returned
	// as the `paging.next.after` JSON property of a paged response containing more
	// results.
	After  param.Opt[string] `query:"after,omitzero" json:"-"`
	Before param.Opt[string] `query:"before,omitzero" json:"-"`
	// The maximum number of results to display per page.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [PageListSitePageRevisionsParams]'s query parameters as
// `url.Values`.
func (r PageListSitePageRevisionsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type PageRestoreLandingPageRevisionParams struct {
	ObjectID string `path:"objectId" api:"required" json:"-"`
	paramObj
}

type PageRestoreLandingPageRevisionToDraftParams struct {
	ObjectID string `path:"objectId" api:"required" json:"-"`
	paramObj
}

type PageRestoreSitePageRevisionParams struct {
	ObjectID string `path:"objectId" api:"required" json:"-"`
	paramObj
}

type PageRestoreSitePageRevisionToDraftParams struct {
	ObjectID string `path:"objectId" api:"required" json:"-"`
	paramObj
}
