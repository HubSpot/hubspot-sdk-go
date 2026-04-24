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
	"github.com/stainless-sdks/hubspot-sdk-go/internal/requestconfig"
	"github.com/stainless-sdks/hubspot-sdk-go/option"
	"github.com/stainless-sdks/hubspot-sdk-go/packages/pagination"
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
func (r *PageService) RestoreLandingPageRevision(ctx context.Context, revisionID string, body PageRestoreLandingPageRevisionParams, opts ...option.RequestOption) (res *PagesPage, err error) {
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
func (r *PageService) RestoreLandingPageRevisionToDraft(ctx context.Context, revisionID int64, body PageRestoreLandingPageRevisionToDraftParams, opts ...option.RequestOption) (res *PagesPage, err error) {
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
func (r *PageService) RestoreSitePageRevision(ctx context.Context, revisionID string, body PageRestoreSitePageRevisionParams, opts ...option.RequestOption) (res *PagesPage, err error) {
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
func (r *PageService) RestoreSitePageRevisionToDraft(ctx context.Context, revisionID int64, body PageRestoreSitePageRevisionToDraftParams, opts ...option.RequestOption) (res *PagesPage, err error) {
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
	Inputs []PagesPageParam `json:"inputs,omitzero" api:"required"`
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
	Results []PagesPage `json:"results" api:"required"`
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
	Results []PagesPage `json:"results" api:"required"`
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

type PageVersion struct {
	ID        string             `json:"id" api:"required"`
	Object    PagesPage          `json:"object" api:"required"`
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

type PagesPage struct {
	// The unique ID of the page.
	ID string `json:"id" api:"required"`
	// The status of the AB test associated with this page, if applicable
	//
	// Any of "automated_loser_variant", "automated_master", "automated_variant",
	// "loser_variant", "mab_master", "mab_variant", "master", "variant".
	AbStatus PagesPageAbStatus `json:"abStatus" api:"required"`
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
	ContentTypeCategory PagesPageContentTypeCategory `json:"contentTypeCategory" api:"required"`
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
	CurrentState PagesPageCurrentState `json:"currentState" api:"required"`
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
	Language PagesPageLanguage `json:"language" api:"required"`
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
func (r PagesPage) RawJSON() string { return r.JSON.raw }
func (r *PagesPage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this PagesPage to a PagesPageParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// PagesPageParam.Overrides()
func (r PagesPage) ToParam() PagesPageParam {
	return param.Override[PagesPageParam](json.RawMessage(r.RawJSON()))
}

// The status of the AB test associated with this page, if applicable
type PagesPageAbStatus string

const (
	PagesPageAbStatusAutomatedLoserVariant PagesPageAbStatus = "automated_loser_variant"
	PagesPageAbStatusAutomatedMaster       PagesPageAbStatus = "automated_master"
	PagesPageAbStatusAutomatedVariant      PagesPageAbStatus = "automated_variant"
	PagesPageAbStatusLoserVariant          PagesPageAbStatus = "loser_variant"
	PagesPageAbStatusMabMaster             PagesPageAbStatus = "mab_master"
	PagesPageAbStatusMabVariant            PagesPageAbStatus = "mab_variant"
	PagesPageAbStatusMaster                PagesPageAbStatus = "master"
	PagesPageAbStatusVariant               PagesPageAbStatus = "variant"
)

// An ENUM descibing the type of this object. Should be either LANDING_PAGE or
// SITE_PAGE.
type PagesPageContentTypeCategory string

const (
	PagesPageContentTypeCategory0  PagesPageContentTypeCategory = "0"
	PagesPageContentTypeCategory1  PagesPageContentTypeCategory = "1"
	PagesPageContentTypeCategory10 PagesPageContentTypeCategory = "10"
	PagesPageContentTypeCategory11 PagesPageContentTypeCategory = "11"
	PagesPageContentTypeCategory12 PagesPageContentTypeCategory = "12"
	PagesPageContentTypeCategory13 PagesPageContentTypeCategory = "13"
	PagesPageContentTypeCategory14 PagesPageContentTypeCategory = "14"
	PagesPageContentTypeCategory15 PagesPageContentTypeCategory = "15"
	PagesPageContentTypeCategory16 PagesPageContentTypeCategory = "16"
	PagesPageContentTypeCategory17 PagesPageContentTypeCategory = "17"
	PagesPageContentTypeCategory18 PagesPageContentTypeCategory = "18"
	PagesPageContentTypeCategory19 PagesPageContentTypeCategory = "19"
	PagesPageContentTypeCategory2  PagesPageContentTypeCategory = "2"
	PagesPageContentTypeCategory20 PagesPageContentTypeCategory = "20"
	PagesPageContentTypeCategory21 PagesPageContentTypeCategory = "21"
	PagesPageContentTypeCategory22 PagesPageContentTypeCategory = "22"
	PagesPageContentTypeCategory3  PagesPageContentTypeCategory = "3"
	PagesPageContentTypeCategory4  PagesPageContentTypeCategory = "4"
	PagesPageContentTypeCategory5  PagesPageContentTypeCategory = "5"
	PagesPageContentTypeCategory6  PagesPageContentTypeCategory = "6"
	PagesPageContentTypeCategory7  PagesPageContentTypeCategory = "7"
	PagesPageContentTypeCategory8  PagesPageContentTypeCategory = "8"
	PagesPageContentTypeCategory9  PagesPageContentTypeCategory = "9"
)

// A generated ENUM descibing the current state of this page.
type PagesPageCurrentState string

const (
	PagesPageCurrentStateAgentGenerated          PagesPageCurrentState = "AGENT_GENERATED"
	PagesPageCurrentStateAutomated               PagesPageCurrentState = "AUTOMATED"
	PagesPageCurrentStateAutomatedAb             PagesPageCurrentState = "AUTOMATED_AB"
	PagesPageCurrentStateAutomatedAbVariant      PagesPageCurrentState = "AUTOMATED_AB_VARIANT"
	PagesPageCurrentStateAutomatedDraft          PagesPageCurrentState = "AUTOMATED_DRAFT"
	PagesPageCurrentStateAutomatedDraftAb        PagesPageCurrentState = "AUTOMATED_DRAFT_AB"
	PagesPageCurrentStateAutomatedDraftAbvariant PagesPageCurrentState = "AUTOMATED_DRAFT_ABVARIANT"
	PagesPageCurrentStateAutomatedForForm        PagesPageCurrentState = "AUTOMATED_FOR_FORM"
	PagesPageCurrentStateAutomatedForFormBuffer  PagesPageCurrentState = "AUTOMATED_FOR_FORM_BUFFER"
	PagesPageCurrentStateAutomatedForFormDraft   PagesPageCurrentState = "AUTOMATED_FOR_FORM_DRAFT"
	PagesPageCurrentStateAutomatedForFormLegacy  PagesPageCurrentState = "AUTOMATED_FOR_FORM_LEGACY"
	PagesPageCurrentStateAutomatedLoserAbvariant PagesPageCurrentState = "AUTOMATED_LOSER_ABVARIANT"
	PagesPageCurrentStateAutomatedSending        PagesPageCurrentState = "AUTOMATED_SENDING"
	PagesPageCurrentStateBlogEmailDraft          PagesPageCurrentState = "BLOG_EMAIL_DRAFT"
	PagesPageCurrentStateBlogEmailPublished      PagesPageCurrentState = "BLOG_EMAIL_PUBLISHED"
	PagesPageCurrentStateDraft                   PagesPageCurrentState = "DRAFT"
	PagesPageCurrentStateDraftAb                 PagesPageCurrentState = "DRAFT_AB"
	PagesPageCurrentStateDraftAbVariant          PagesPageCurrentState = "DRAFT_AB_VARIANT"
	PagesPageCurrentStateError                   PagesPageCurrentState = "ERROR"
	PagesPageCurrentStateLoserAbVariant          PagesPageCurrentState = "LOSER_AB_VARIANT"
	PagesPageCurrentStatePageStub                PagesPageCurrentState = "PAGE_STUB"
	PagesPageCurrentStatePreProcessing           PagesPageCurrentState = "PRE_PROCESSING"
	PagesPageCurrentStateProcessing              PagesPageCurrentState = "PROCESSING"
	PagesPageCurrentStatePublished               PagesPageCurrentState = "PUBLISHED"
	PagesPageCurrentStatePublishedAb             PagesPageCurrentState = "PUBLISHED_AB"
	PagesPageCurrentStatePublishedAbVariant      PagesPageCurrentState = "PUBLISHED_AB_VARIANT"
	PagesPageCurrentStatePublishedOrScheduled    PagesPageCurrentState = "PUBLISHED_OR_SCHEDULED"
	PagesPageCurrentStateRssToEmailDraft         PagesPageCurrentState = "RSS_TO_EMAIL_DRAFT"
	PagesPageCurrentStateRssToEmailPublished     PagesPageCurrentState = "RSS_TO_EMAIL_PUBLISHED"
	PagesPageCurrentStateScheduled               PagesPageCurrentState = "SCHEDULED"
	PagesPageCurrentStateScheduledAb             PagesPageCurrentState = "SCHEDULED_AB"
	PagesPageCurrentStateScheduledOrPublished    PagesPageCurrentState = "SCHEDULED_OR_PUBLISHED"
)

// The explicitly defined ISO 639 language code of the page. If null, the page will
// default to the language of the Domain.
type PagesPageLanguage string

const (
	PagesPageLanguageAa     PagesPageLanguage = "aa"
	PagesPageLanguageAb     PagesPageLanguage = "ab"
	PagesPageLanguageAe     PagesPageLanguage = "ae"
	PagesPageLanguageAf     PagesPageLanguage = "af"
	PagesPageLanguageAfNa   PagesPageLanguage = "af-na"
	PagesPageLanguageAfZa   PagesPageLanguage = "af-za"
	PagesPageLanguageAgq    PagesPageLanguage = "agq"
	PagesPageLanguageAgqCm  PagesPageLanguage = "agq-cm"
	PagesPageLanguageAk     PagesPageLanguage = "ak"
	PagesPageLanguageAkGh   PagesPageLanguage = "ak-gh"
	PagesPageLanguageAm     PagesPageLanguage = "am"
	PagesPageLanguageAmEt   PagesPageLanguage = "am-et"
	PagesPageLanguageAn     PagesPageLanguage = "an"
	PagesPageLanguageAnn    PagesPageLanguage = "ann"
	PagesPageLanguageAnnNg  PagesPageLanguage = "ann-ng"
	PagesPageLanguageAr     PagesPageLanguage = "ar"
	PagesPageLanguageAr001  PagesPageLanguage = "ar-001"
	PagesPageLanguageArAe   PagesPageLanguage = "ar-ae"
	PagesPageLanguageArBh   PagesPageLanguage = "ar-bh"
	PagesPageLanguageArDj   PagesPageLanguage = "ar-dj"
	PagesPageLanguageArDz   PagesPageLanguage = "ar-dz"
	PagesPageLanguageArEg   PagesPageLanguage = "ar-eg"
	PagesPageLanguageArEh   PagesPageLanguage = "ar-eh"
	PagesPageLanguageArEr   PagesPageLanguage = "ar-er"
	PagesPageLanguageArIl   PagesPageLanguage = "ar-il"
	PagesPageLanguageArIq   PagesPageLanguage = "ar-iq"
	PagesPageLanguageArJo   PagesPageLanguage = "ar-jo"
	PagesPageLanguageArKm   PagesPageLanguage = "ar-km"
	PagesPageLanguageArKw   PagesPageLanguage = "ar-kw"
	PagesPageLanguageArLb   PagesPageLanguage = "ar-lb"
	PagesPageLanguageArLy   PagesPageLanguage = "ar-ly"
	PagesPageLanguageArMa   PagesPageLanguage = "ar-ma"
	PagesPageLanguageArMr   PagesPageLanguage = "ar-mr"
	PagesPageLanguageArOm   PagesPageLanguage = "ar-om"
	PagesPageLanguageArPs   PagesPageLanguage = "ar-ps"
	PagesPageLanguageArQa   PagesPageLanguage = "ar-qa"
	PagesPageLanguageArSa   PagesPageLanguage = "ar-sa"
	PagesPageLanguageArSd   PagesPageLanguage = "ar-sd"
	PagesPageLanguageArSo   PagesPageLanguage = "ar-so"
	PagesPageLanguageArSS   PagesPageLanguage = "ar-ss"
	PagesPageLanguageArSy   PagesPageLanguage = "ar-sy"
	PagesPageLanguageArTd   PagesPageLanguage = "ar-td"
	PagesPageLanguageArTn   PagesPageLanguage = "ar-tn"
	PagesPageLanguageArYe   PagesPageLanguage = "ar-ye"
	PagesPageLanguageAs     PagesPageLanguage = "as"
	PagesPageLanguageAsIn   PagesPageLanguage = "as-in"
	PagesPageLanguageAsa    PagesPageLanguage = "asa"
	PagesPageLanguageAsaTz  PagesPageLanguage = "asa-tz"
	PagesPageLanguageAst    PagesPageLanguage = "ast"
	PagesPageLanguageAstEs  PagesPageLanguage = "ast-es"
	PagesPageLanguageAv     PagesPageLanguage = "av"
	PagesPageLanguageAy     PagesPageLanguage = "ay"
	PagesPageLanguageAz     PagesPageLanguage = "az"
	PagesPageLanguageAzAz   PagesPageLanguage = "az-az"
	PagesPageLanguageBa     PagesPageLanguage = "ba"
	PagesPageLanguageBas    PagesPageLanguage = "bas"
	PagesPageLanguageBasCm  PagesPageLanguage = "bas-cm"
	PagesPageLanguageBe     PagesPageLanguage = "be"
	PagesPageLanguageBeBy   PagesPageLanguage = "be-by"
	PagesPageLanguageBem    PagesPageLanguage = "bem"
	PagesPageLanguageBemZm  PagesPageLanguage = "bem-zm"
	PagesPageLanguageBez    PagesPageLanguage = "bez"
	PagesPageLanguageBezTz  PagesPageLanguage = "bez-tz"
	PagesPageLanguageBg     PagesPageLanguage = "bg"
	PagesPageLanguageBgBg   PagesPageLanguage = "bg-bg"
	PagesPageLanguageBgc    PagesPageLanguage = "bgc"
	PagesPageLanguageBgcIn  PagesPageLanguage = "bgc-in"
	PagesPageLanguageBho    PagesPageLanguage = "bho"
	PagesPageLanguageBhoIn  PagesPageLanguage = "bho-in"
	PagesPageLanguageBi     PagesPageLanguage = "bi"
	PagesPageLanguageBm     PagesPageLanguage = "bm"
	PagesPageLanguageBmMl   PagesPageLanguage = "bm-ml"
	PagesPageLanguageBn     PagesPageLanguage = "bn"
	PagesPageLanguageBnBd   PagesPageLanguage = "bn-bd"
	PagesPageLanguageBnIn   PagesPageLanguage = "bn-in"
	PagesPageLanguageBo     PagesPageLanguage = "bo"
	PagesPageLanguageBoCn   PagesPageLanguage = "bo-cn"
	PagesPageLanguageBoIn   PagesPageLanguage = "bo-in"
	PagesPageLanguageBr     PagesPageLanguage = "br"
	PagesPageLanguageBrFr   PagesPageLanguage = "br-fr"
	PagesPageLanguageBrx    PagesPageLanguage = "brx"
	PagesPageLanguageBrxIn  PagesPageLanguage = "brx-in"
	PagesPageLanguageBs     PagesPageLanguage = "bs"
	PagesPageLanguageBsBa   PagesPageLanguage = "bs-ba"
	PagesPageLanguageCa     PagesPageLanguage = "ca"
	PagesPageLanguageCaAd   PagesPageLanguage = "ca-ad"
	PagesPageLanguageCaEs   PagesPageLanguage = "ca-es"
	PagesPageLanguageCaFr   PagesPageLanguage = "ca-fr"
	PagesPageLanguageCaIt   PagesPageLanguage = "ca-it"
	PagesPageLanguageCcp    PagesPageLanguage = "ccp"
	PagesPageLanguageCcpBd  PagesPageLanguage = "ccp-bd"
	PagesPageLanguageCcpIn  PagesPageLanguage = "ccp-in"
	PagesPageLanguageCe     PagesPageLanguage = "ce"
	PagesPageLanguageCeRu   PagesPageLanguage = "ce-ru"
	PagesPageLanguageCeb    PagesPageLanguage = "ceb"
	PagesPageLanguageCebPh  PagesPageLanguage = "ceb-ph"
	PagesPageLanguageCgg    PagesPageLanguage = "cgg"
	PagesPageLanguageCggUg  PagesPageLanguage = "cgg-ug"
	PagesPageLanguageCh     PagesPageLanguage = "ch"
	PagesPageLanguageChr    PagesPageLanguage = "chr"
	PagesPageLanguageChrUs  PagesPageLanguage = "chr-us"
	PagesPageLanguageCkb    PagesPageLanguage = "ckb"
	PagesPageLanguageCkbIq  PagesPageLanguage = "ckb-iq"
	PagesPageLanguageCkbIr  PagesPageLanguage = "ckb-ir"
	PagesPageLanguageCo     PagesPageLanguage = "co"
	PagesPageLanguageCr     PagesPageLanguage = "cr"
	PagesPageLanguageCs     PagesPageLanguage = "cs"
	PagesPageLanguageCsCz   PagesPageLanguage = "cs-cz"
	PagesPageLanguageCu     PagesPageLanguage = "cu"
	PagesPageLanguageCuRu   PagesPageLanguage = "cu-ru"
	PagesPageLanguageCv     PagesPageLanguage = "cv"
	PagesPageLanguageCvRu   PagesPageLanguage = "cv-ru"
	PagesPageLanguageCy     PagesPageLanguage = "cy"
	PagesPageLanguageCyGB   PagesPageLanguage = "cy-gb"
	PagesPageLanguageDa     PagesPageLanguage = "da"
	PagesPageLanguageDaDk   PagesPageLanguage = "da-dk"
	PagesPageLanguageDaGl   PagesPageLanguage = "da-gl"
	PagesPageLanguageDav    PagesPageLanguage = "dav"
	PagesPageLanguageDavKe  PagesPageLanguage = "dav-ke"
	PagesPageLanguageDe     PagesPageLanguage = "de"
	PagesPageLanguageDeAt   PagesPageLanguage = "de-at"
	PagesPageLanguageDeBe   PagesPageLanguage = "de-be"
	PagesPageLanguageDeCh   PagesPageLanguage = "de-ch"
	PagesPageLanguageDeDe   PagesPageLanguage = "de-de"
	PagesPageLanguageDeGr   PagesPageLanguage = "de-gr"
	PagesPageLanguageDeIt   PagesPageLanguage = "de-it"
	PagesPageLanguageDeLi   PagesPageLanguage = "de-li"
	PagesPageLanguageDeLu   PagesPageLanguage = "de-lu"
	PagesPageLanguageDje    PagesPageLanguage = "dje"
	PagesPageLanguageDjeNe  PagesPageLanguage = "dje-ne"
	PagesPageLanguageDoi    PagesPageLanguage = "doi"
	PagesPageLanguageDoiIn  PagesPageLanguage = "doi-in"
	PagesPageLanguageDsb    PagesPageLanguage = "dsb"
	PagesPageLanguageDsbDe  PagesPageLanguage = "dsb-de"
	PagesPageLanguageDua    PagesPageLanguage = "dua"
	PagesPageLanguageDuaCm  PagesPageLanguage = "dua-cm"
	PagesPageLanguageDv     PagesPageLanguage = "dv"
	PagesPageLanguageDyo    PagesPageLanguage = "dyo"
	PagesPageLanguageDyoSn  PagesPageLanguage = "dyo-sn"
	PagesPageLanguageDz     PagesPageLanguage = "dz"
	PagesPageLanguageDzBt   PagesPageLanguage = "dz-bt"
	PagesPageLanguageEbu    PagesPageLanguage = "ebu"
	PagesPageLanguageEbuKe  PagesPageLanguage = "ebu-ke"
	PagesPageLanguageEe     PagesPageLanguage = "ee"
	PagesPageLanguageEeGh   PagesPageLanguage = "ee-gh"
	PagesPageLanguageEeTg   PagesPageLanguage = "ee-tg"
	PagesPageLanguageEl     PagesPageLanguage = "el"
	PagesPageLanguageElCy   PagesPageLanguage = "el-cy"
	PagesPageLanguageElGr   PagesPageLanguage = "el-gr"
	PagesPageLanguageEn     PagesPageLanguage = "en"
	PagesPageLanguageEn001  PagesPageLanguage = "en-001"
	PagesPageLanguageEn150  PagesPageLanguage = "en-150"
	PagesPageLanguageEnAe   PagesPageLanguage = "en-ae"
	PagesPageLanguageEnAg   PagesPageLanguage = "en-ag"
	PagesPageLanguageEnAI   PagesPageLanguage = "en-ai"
	PagesPageLanguageEnAs   PagesPageLanguage = "en-as"
	PagesPageLanguageEnAt   PagesPageLanguage = "en-at"
	PagesPageLanguageEnAu   PagesPageLanguage = "en-au"
	PagesPageLanguageEnBb   PagesPageLanguage = "en-bb"
	PagesPageLanguageEnBe   PagesPageLanguage = "en-be"
	PagesPageLanguageEnBi   PagesPageLanguage = "en-bi"
	PagesPageLanguageEnBm   PagesPageLanguage = "en-bm"
	PagesPageLanguageEnBs   PagesPageLanguage = "en-bs"
	PagesPageLanguageEnBw   PagesPageLanguage = "en-bw"
	PagesPageLanguageEnBz   PagesPageLanguage = "en-bz"
	PagesPageLanguageEnCa   PagesPageLanguage = "en-ca"
	PagesPageLanguageEnCc   PagesPageLanguage = "en-cc"
	PagesPageLanguageEnCh   PagesPageLanguage = "en-ch"
	PagesPageLanguageEnCk   PagesPageLanguage = "en-ck"
	PagesPageLanguageEnCm   PagesPageLanguage = "en-cm"
	PagesPageLanguageEnCn   PagesPageLanguage = "en-cn"
	PagesPageLanguageEnCx   PagesPageLanguage = "en-cx"
	PagesPageLanguageEnCy   PagesPageLanguage = "en-cy"
	PagesPageLanguageEnDe   PagesPageLanguage = "en-de"
	PagesPageLanguageEnDg   PagesPageLanguage = "en-dg"
	PagesPageLanguageEnDk   PagesPageLanguage = "en-dk"
	PagesPageLanguageEnDm   PagesPageLanguage = "en-dm"
	PagesPageLanguageEnEe   PagesPageLanguage = "en-ee"
	PagesPageLanguageEnEg   PagesPageLanguage = "en-eg"
	PagesPageLanguageEnEr   PagesPageLanguage = "en-er"
	PagesPageLanguageEnEs   PagesPageLanguage = "en-es"
	PagesPageLanguageEnFi   PagesPageLanguage = "en-fi"
	PagesPageLanguageEnFj   PagesPageLanguage = "en-fj"
	PagesPageLanguageEnFk   PagesPageLanguage = "en-fk"
	PagesPageLanguageEnFm   PagesPageLanguage = "en-fm"
	PagesPageLanguageEnFr   PagesPageLanguage = "en-fr"
	PagesPageLanguageEnGB   PagesPageLanguage = "en-gb"
	PagesPageLanguageEnGd   PagesPageLanguage = "en-gd"
	PagesPageLanguageEnGg   PagesPageLanguage = "en-gg"
	PagesPageLanguageEnGh   PagesPageLanguage = "en-gh"
	PagesPageLanguageEnGi   PagesPageLanguage = "en-gi"
	PagesPageLanguageEnGm   PagesPageLanguage = "en-gm"
	PagesPageLanguageEnGu   PagesPageLanguage = "en-gu"
	PagesPageLanguageEnGy   PagesPageLanguage = "en-gy"
	PagesPageLanguageEnHk   PagesPageLanguage = "en-hk"
	PagesPageLanguageEnID   PagesPageLanguage = "en-id"
	PagesPageLanguageEnIe   PagesPageLanguage = "en-ie"
	PagesPageLanguageEnIl   PagesPageLanguage = "en-il"
	PagesPageLanguageEnIm   PagesPageLanguage = "en-im"
	PagesPageLanguageEnIn   PagesPageLanguage = "en-in"
	PagesPageLanguageEnIo   PagesPageLanguage = "en-io"
	PagesPageLanguageEnJe   PagesPageLanguage = "en-je"
	PagesPageLanguageEnJm   PagesPageLanguage = "en-jm"
	PagesPageLanguageEnKe   PagesPageLanguage = "en-ke"
	PagesPageLanguageEnKi   PagesPageLanguage = "en-ki"
	PagesPageLanguageEnKn   PagesPageLanguage = "en-kn"
	PagesPageLanguageEnKy   PagesPageLanguage = "en-ky"
	PagesPageLanguageEnLc   PagesPageLanguage = "en-lc"
	PagesPageLanguageEnLr   PagesPageLanguage = "en-lr"
	PagesPageLanguageEnLs   PagesPageLanguage = "en-ls"
	PagesPageLanguageEnLu   PagesPageLanguage = "en-lu"
	PagesPageLanguageEnMg   PagesPageLanguage = "en-mg"
	PagesPageLanguageEnMh   PagesPageLanguage = "en-mh"
	PagesPageLanguageEnMo   PagesPageLanguage = "en-mo"
	PagesPageLanguageEnMp   PagesPageLanguage = "en-mp"
	PagesPageLanguageEnMs   PagesPageLanguage = "en-ms"
	PagesPageLanguageEnMt   PagesPageLanguage = "en-mt"
	PagesPageLanguageEnMu   PagesPageLanguage = "en-mu"
	PagesPageLanguageEnMv   PagesPageLanguage = "en-mv"
	PagesPageLanguageEnMw   PagesPageLanguage = "en-mw"
	PagesPageLanguageEnMx   PagesPageLanguage = "en-mx"
	PagesPageLanguageEnMy   PagesPageLanguage = "en-my"
	PagesPageLanguageEnNa   PagesPageLanguage = "en-na"
	PagesPageLanguageEnNf   PagesPageLanguage = "en-nf"
	PagesPageLanguageEnNg   PagesPageLanguage = "en-ng"
	PagesPageLanguageEnNl   PagesPageLanguage = "en-nl"
	PagesPageLanguageEnNr   PagesPageLanguage = "en-nr"
	PagesPageLanguageEnNu   PagesPageLanguage = "en-nu"
	PagesPageLanguageEnNz   PagesPageLanguage = "en-nz"
	PagesPageLanguageEnPg   PagesPageLanguage = "en-pg"
	PagesPageLanguageEnPh   PagesPageLanguage = "en-ph"
	PagesPageLanguageEnPk   PagesPageLanguage = "en-pk"
	PagesPageLanguageEnPn   PagesPageLanguage = "en-pn"
	PagesPageLanguageEnPr   PagesPageLanguage = "en-pr"
	PagesPageLanguageEnPt   PagesPageLanguage = "en-pt"
	PagesPageLanguageEnPw   PagesPageLanguage = "en-pw"
	PagesPageLanguageEnRw   PagesPageLanguage = "en-rw"
	PagesPageLanguageEnSb   PagesPageLanguage = "en-sb"
	PagesPageLanguageEnSc   PagesPageLanguage = "en-sc"
	PagesPageLanguageEnSd   PagesPageLanguage = "en-sd"
	PagesPageLanguageEnSe   PagesPageLanguage = "en-se"
	PagesPageLanguageEnSg   PagesPageLanguage = "en-sg"
	PagesPageLanguageEnSh   PagesPageLanguage = "en-sh"
	PagesPageLanguageEnSi   PagesPageLanguage = "en-si"
	PagesPageLanguageEnSl   PagesPageLanguage = "en-sl"
	PagesPageLanguageEnSS   PagesPageLanguage = "en-ss"
	PagesPageLanguageEnSx   PagesPageLanguage = "en-sx"
	PagesPageLanguageEnSz   PagesPageLanguage = "en-sz"
	PagesPageLanguageEnTc   PagesPageLanguage = "en-tc"
	PagesPageLanguageEnTh   PagesPageLanguage = "en-th"
	PagesPageLanguageEnTk   PagesPageLanguage = "en-tk"
	PagesPageLanguageEnTn   PagesPageLanguage = "en-tn"
	PagesPageLanguageEnTo   PagesPageLanguage = "en-to"
	PagesPageLanguageEnTt   PagesPageLanguage = "en-tt"
	PagesPageLanguageEnTv   PagesPageLanguage = "en-tv"
	PagesPageLanguageEnTz   PagesPageLanguage = "en-tz"
	PagesPageLanguageEnUg   PagesPageLanguage = "en-ug"
	PagesPageLanguageEnUm   PagesPageLanguage = "en-um"
	PagesPageLanguageEnUs   PagesPageLanguage = "en-us"
	PagesPageLanguageEnVc   PagesPageLanguage = "en-vc"
	PagesPageLanguageEnVg   PagesPageLanguage = "en-vg"
	PagesPageLanguageEnVi   PagesPageLanguage = "en-vi"
	PagesPageLanguageEnVn   PagesPageLanguage = "en-vn"
	PagesPageLanguageEnVu   PagesPageLanguage = "en-vu"
	PagesPageLanguageEnWs   PagesPageLanguage = "en-ws"
	PagesPageLanguageEnZa   PagesPageLanguage = "en-za"
	PagesPageLanguageEnZm   PagesPageLanguage = "en-zm"
	PagesPageLanguageEnZw   PagesPageLanguage = "en-zw"
	PagesPageLanguageEo     PagesPageLanguage = "eo"
	PagesPageLanguageEo001  PagesPageLanguage = "eo-001"
	PagesPageLanguageEs     PagesPageLanguage = "es"
	PagesPageLanguageEs419  PagesPageLanguage = "es-419"
	PagesPageLanguageEsAr   PagesPageLanguage = "es-ar"
	PagesPageLanguageEsBo   PagesPageLanguage = "es-bo"
	PagesPageLanguageEsBr   PagesPageLanguage = "es-br"
	PagesPageLanguageEsBz   PagesPageLanguage = "es-bz"
	PagesPageLanguageEsCl   PagesPageLanguage = "es-cl"
	PagesPageLanguageEsCo   PagesPageLanguage = "es-co"
	PagesPageLanguageEsCr   PagesPageLanguage = "es-cr"
	PagesPageLanguageEsCu   PagesPageLanguage = "es-cu"
	PagesPageLanguageEsDo   PagesPageLanguage = "es-do"
	PagesPageLanguageEsEa   PagesPageLanguage = "es-ea"
	PagesPageLanguageEsEc   PagesPageLanguage = "es-ec"
	PagesPageLanguageEsEs   PagesPageLanguage = "es-es"
	PagesPageLanguageEsGq   PagesPageLanguage = "es-gq"
	PagesPageLanguageEsGt   PagesPageLanguage = "es-gt"
	PagesPageLanguageEsHn   PagesPageLanguage = "es-hn"
	PagesPageLanguageEsIc   PagesPageLanguage = "es-ic"
	PagesPageLanguageEsMx   PagesPageLanguage = "es-mx"
	PagesPageLanguageEsNi   PagesPageLanguage = "es-ni"
	PagesPageLanguageEsPa   PagesPageLanguage = "es-pa"
	PagesPageLanguageEsPe   PagesPageLanguage = "es-pe"
	PagesPageLanguageEsPh   PagesPageLanguage = "es-ph"
	PagesPageLanguageEsPr   PagesPageLanguage = "es-pr"
	PagesPageLanguageEsPy   PagesPageLanguage = "es-py"
	PagesPageLanguageEsSv   PagesPageLanguage = "es-sv"
	PagesPageLanguageEsUs   PagesPageLanguage = "es-us"
	PagesPageLanguageEsUy   PagesPageLanguage = "es-uy"
	PagesPageLanguageEsVe   PagesPageLanguage = "es-ve"
	PagesPageLanguageEt     PagesPageLanguage = "et"
	PagesPageLanguageEtEe   PagesPageLanguage = "et-ee"
	PagesPageLanguageEu     PagesPageLanguage = "eu"
	PagesPageLanguageEuEs   PagesPageLanguage = "eu-es"
	PagesPageLanguageEwo    PagesPageLanguage = "ewo"
	PagesPageLanguageEwoCm  PagesPageLanguage = "ewo-cm"
	PagesPageLanguageFa     PagesPageLanguage = "fa"
	PagesPageLanguageFaAf   PagesPageLanguage = "fa-af"
	PagesPageLanguageFaIr   PagesPageLanguage = "fa-ir"
	PagesPageLanguageFf     PagesPageLanguage = "ff"
	PagesPageLanguageFfBf   PagesPageLanguage = "ff-bf"
	PagesPageLanguageFfCm   PagesPageLanguage = "ff-cm"
	PagesPageLanguageFfGh   PagesPageLanguage = "ff-gh"
	PagesPageLanguageFfGm   PagesPageLanguage = "ff-gm"
	PagesPageLanguageFfGn   PagesPageLanguage = "ff-gn"
	PagesPageLanguageFfGw   PagesPageLanguage = "ff-gw"
	PagesPageLanguageFfLr   PagesPageLanguage = "ff-lr"
	PagesPageLanguageFfMr   PagesPageLanguage = "ff-mr"
	PagesPageLanguageFfNe   PagesPageLanguage = "ff-ne"
	PagesPageLanguageFfNg   PagesPageLanguage = "ff-ng"
	PagesPageLanguageFfSl   PagesPageLanguage = "ff-sl"
	PagesPageLanguageFfSn   PagesPageLanguage = "ff-sn"
	PagesPageLanguageFi     PagesPageLanguage = "fi"
	PagesPageLanguageFiFi   PagesPageLanguage = "fi-fi"
	PagesPageLanguageFil    PagesPageLanguage = "fil"
	PagesPageLanguageFilPh  PagesPageLanguage = "fil-ph"
	PagesPageLanguageFj     PagesPageLanguage = "fj"
	PagesPageLanguageFo     PagesPageLanguage = "fo"
	PagesPageLanguageFoDk   PagesPageLanguage = "fo-dk"
	PagesPageLanguageFoFo   PagesPageLanguage = "fo-fo"
	PagesPageLanguageFr     PagesPageLanguage = "fr"
	PagesPageLanguageFrBe   PagesPageLanguage = "fr-be"
	PagesPageLanguageFrBf   PagesPageLanguage = "fr-bf"
	PagesPageLanguageFrBi   PagesPageLanguage = "fr-bi"
	PagesPageLanguageFrBj   PagesPageLanguage = "fr-bj"
	PagesPageLanguageFrBl   PagesPageLanguage = "fr-bl"
	PagesPageLanguageFrCa   PagesPageLanguage = "fr-ca"
	PagesPageLanguageFrCd   PagesPageLanguage = "fr-cd"
	PagesPageLanguageFrCf   PagesPageLanguage = "fr-cf"
	PagesPageLanguageFrCg   PagesPageLanguage = "fr-cg"
	PagesPageLanguageFrCh   PagesPageLanguage = "fr-ch"
	PagesPageLanguageFrCi   PagesPageLanguage = "fr-ci"
	PagesPageLanguageFrCm   PagesPageLanguage = "fr-cm"
	PagesPageLanguageFrDj   PagesPageLanguage = "fr-dj"
	PagesPageLanguageFrDz   PagesPageLanguage = "fr-dz"
	PagesPageLanguageFrFr   PagesPageLanguage = "fr-fr"
	PagesPageLanguageFrGa   PagesPageLanguage = "fr-ga"
	PagesPageLanguageFrGf   PagesPageLanguage = "fr-gf"
	PagesPageLanguageFrGn   PagesPageLanguage = "fr-gn"
	PagesPageLanguageFrGp   PagesPageLanguage = "fr-gp"
	PagesPageLanguageFrGq   PagesPageLanguage = "fr-gq"
	PagesPageLanguageFrHt   PagesPageLanguage = "fr-ht"
	PagesPageLanguageFrKm   PagesPageLanguage = "fr-km"
	PagesPageLanguageFrLu   PagesPageLanguage = "fr-lu"
	PagesPageLanguageFrMa   PagesPageLanguage = "fr-ma"
	PagesPageLanguageFrMc   PagesPageLanguage = "fr-mc"
	PagesPageLanguageFrMf   PagesPageLanguage = "fr-mf"
	PagesPageLanguageFrMg   PagesPageLanguage = "fr-mg"
	PagesPageLanguageFrMl   PagesPageLanguage = "fr-ml"
	PagesPageLanguageFrMq   PagesPageLanguage = "fr-mq"
	PagesPageLanguageFrMr   PagesPageLanguage = "fr-mr"
	PagesPageLanguageFrMu   PagesPageLanguage = "fr-mu"
	PagesPageLanguageFrNc   PagesPageLanguage = "fr-nc"
	PagesPageLanguageFrNe   PagesPageLanguage = "fr-ne"
	PagesPageLanguageFrPf   PagesPageLanguage = "fr-pf"
	PagesPageLanguageFrPm   PagesPageLanguage = "fr-pm"
	PagesPageLanguageFrRe   PagesPageLanguage = "fr-re"
	PagesPageLanguageFrRw   PagesPageLanguage = "fr-rw"
	PagesPageLanguageFrSc   PagesPageLanguage = "fr-sc"
	PagesPageLanguageFrSn   PagesPageLanguage = "fr-sn"
	PagesPageLanguageFrSy   PagesPageLanguage = "fr-sy"
	PagesPageLanguageFrTd   PagesPageLanguage = "fr-td"
	PagesPageLanguageFrTg   PagesPageLanguage = "fr-tg"
	PagesPageLanguageFrTn   PagesPageLanguage = "fr-tn"
	PagesPageLanguageFrVu   PagesPageLanguage = "fr-vu"
	PagesPageLanguageFrWf   PagesPageLanguage = "fr-wf"
	PagesPageLanguageFrYt   PagesPageLanguage = "fr-yt"
	PagesPageLanguageFrr    PagesPageLanguage = "frr"
	PagesPageLanguageFrrDe  PagesPageLanguage = "frr-de"
	PagesPageLanguageFur    PagesPageLanguage = "fur"
	PagesPageLanguageFurIt  PagesPageLanguage = "fur-it"
	PagesPageLanguageFy     PagesPageLanguage = "fy"
	PagesPageLanguageFyNl   PagesPageLanguage = "fy-nl"
	PagesPageLanguageGa     PagesPageLanguage = "ga"
	PagesPageLanguageGaGB   PagesPageLanguage = "ga-gb"
	PagesPageLanguageGaIe   PagesPageLanguage = "ga-ie"
	PagesPageLanguageGd     PagesPageLanguage = "gd"
	PagesPageLanguageGdGB   PagesPageLanguage = "gd-gb"
	PagesPageLanguageGl     PagesPageLanguage = "gl"
	PagesPageLanguageGlEs   PagesPageLanguage = "gl-es"
	PagesPageLanguageGn     PagesPageLanguage = "gn"
	PagesPageLanguageGsw    PagesPageLanguage = "gsw"
	PagesPageLanguageGswCh  PagesPageLanguage = "gsw-ch"
	PagesPageLanguageGswFr  PagesPageLanguage = "gsw-fr"
	PagesPageLanguageGswLi  PagesPageLanguage = "gsw-li"
	PagesPageLanguageGu     PagesPageLanguage = "gu"
	PagesPageLanguageGuIn   PagesPageLanguage = "gu-in"
	PagesPageLanguageGuz    PagesPageLanguage = "guz"
	PagesPageLanguageGuzKe  PagesPageLanguage = "guz-ke"
	PagesPageLanguageGv     PagesPageLanguage = "gv"
	PagesPageLanguageGvIm   PagesPageLanguage = "gv-im"
	PagesPageLanguageHa     PagesPageLanguage = "ha"
	PagesPageLanguageHaGh   PagesPageLanguage = "ha-gh"
	PagesPageLanguageHaNe   PagesPageLanguage = "ha-ne"
	PagesPageLanguageHaNg   PagesPageLanguage = "ha-ng"
	PagesPageLanguageHaw    PagesPageLanguage = "haw"
	PagesPageLanguageHawUs  PagesPageLanguage = "haw-us"
	PagesPageLanguageHe     PagesPageLanguage = "he"
	PagesPageLanguageHeIl   PagesPageLanguage = "he-il"
	PagesPageLanguageHi     PagesPageLanguage = "hi"
	PagesPageLanguageHiIn   PagesPageLanguage = "hi-in"
	PagesPageLanguageHmn    PagesPageLanguage = "hmn"
	PagesPageLanguageHo     PagesPageLanguage = "ho"
	PagesPageLanguageHr     PagesPageLanguage = "hr"
	PagesPageLanguageHrBa   PagesPageLanguage = "hr-ba"
	PagesPageLanguageHrHr   PagesPageLanguage = "hr-hr"
	PagesPageLanguageHsb    PagesPageLanguage = "hsb"
	PagesPageLanguageHsbDe  PagesPageLanguage = "hsb-de"
	PagesPageLanguageHt     PagesPageLanguage = "ht"
	PagesPageLanguageHu     PagesPageLanguage = "hu"
	PagesPageLanguageHuHu   PagesPageLanguage = "hu-hu"
	PagesPageLanguageHy     PagesPageLanguage = "hy"
	PagesPageLanguageHyAm   PagesPageLanguage = "hy-am"
	PagesPageLanguageHz     PagesPageLanguage = "hz"
	PagesPageLanguageIa     PagesPageLanguage = "ia"
	PagesPageLanguageIa001  PagesPageLanguage = "ia-001"
	PagesPageLanguageID     PagesPageLanguage = "id"
	PagesPageLanguageIDID   PagesPageLanguage = "id-id"
	PagesPageLanguageIe     PagesPageLanguage = "ie"
	PagesPageLanguageIg     PagesPageLanguage = "ig"
	PagesPageLanguageIgNg   PagesPageLanguage = "ig-ng"
	PagesPageLanguageIi     PagesPageLanguage = "ii"
	PagesPageLanguageIiCn   PagesPageLanguage = "ii-cn"
	PagesPageLanguageIk     PagesPageLanguage = "ik"
	PagesPageLanguageIo     PagesPageLanguage = "io"
	PagesPageLanguageIs     PagesPageLanguage = "is"
	PagesPageLanguageIsIs   PagesPageLanguage = "is-is"
	PagesPageLanguageIt     PagesPageLanguage = "it"
	PagesPageLanguageItCh   PagesPageLanguage = "it-ch"
	PagesPageLanguageItIt   PagesPageLanguage = "it-it"
	PagesPageLanguageItSm   PagesPageLanguage = "it-sm"
	PagesPageLanguageItVa   PagesPageLanguage = "it-va"
	PagesPageLanguageIu     PagesPageLanguage = "iu"
	PagesPageLanguageJa     PagesPageLanguage = "ja"
	PagesPageLanguageJaJp   PagesPageLanguage = "ja-jp"
	PagesPageLanguageJgo    PagesPageLanguage = "jgo"
	PagesPageLanguageJgoCm  PagesPageLanguage = "jgo-cm"
	PagesPageLanguageJmc    PagesPageLanguage = "jmc"
	PagesPageLanguageJmcTz  PagesPageLanguage = "jmc-tz"
	PagesPageLanguageJv     PagesPageLanguage = "jv"
	PagesPageLanguageJvID   PagesPageLanguage = "jv-id"
	PagesPageLanguageKa     PagesPageLanguage = "ka"
	PagesPageLanguageKaGe   PagesPageLanguage = "ka-ge"
	PagesPageLanguageKab    PagesPageLanguage = "kab"
	PagesPageLanguageKabDz  PagesPageLanguage = "kab-dz"
	PagesPageLanguageKam    PagesPageLanguage = "kam"
	PagesPageLanguageKamKe  PagesPageLanguage = "kam-ke"
	PagesPageLanguageKar    PagesPageLanguage = "kar"
	PagesPageLanguageKde    PagesPageLanguage = "kde"
	PagesPageLanguageKdeTz  PagesPageLanguage = "kde-tz"
	PagesPageLanguageKea    PagesPageLanguage = "kea"
	PagesPageLanguageKeaCv  PagesPageLanguage = "kea-cv"
	PagesPageLanguageKg     PagesPageLanguage = "kg"
	PagesPageLanguageKgp    PagesPageLanguage = "kgp"
	PagesPageLanguageKgpBr  PagesPageLanguage = "kgp-br"
	PagesPageLanguageKh     PagesPageLanguage = "kh"
	PagesPageLanguageKhq    PagesPageLanguage = "khq"
	PagesPageLanguageKhqMl  PagesPageLanguage = "khq-ml"
	PagesPageLanguageKi     PagesPageLanguage = "ki"
	PagesPageLanguageKiKe   PagesPageLanguage = "ki-ke"
	PagesPageLanguageKj     PagesPageLanguage = "kj"
	PagesPageLanguageKk     PagesPageLanguage = "kk"
	PagesPageLanguageKkKz   PagesPageLanguage = "kk-kz"
	PagesPageLanguageKkj    PagesPageLanguage = "kkj"
	PagesPageLanguageKkjCm  PagesPageLanguage = "kkj-cm"
	PagesPageLanguageKl     PagesPageLanguage = "kl"
	PagesPageLanguageKlGl   PagesPageLanguage = "kl-gl"
	PagesPageLanguageKln    PagesPageLanguage = "kln"
	PagesPageLanguageKlnKe  PagesPageLanguage = "kln-ke"
	PagesPageLanguageKm     PagesPageLanguage = "km"
	PagesPageLanguageKmKh   PagesPageLanguage = "km-kh"
	PagesPageLanguageKn     PagesPageLanguage = "kn"
	PagesPageLanguageKnIn   PagesPageLanguage = "kn-in"
	PagesPageLanguageKo     PagesPageLanguage = "ko"
	PagesPageLanguageKoKp   PagesPageLanguage = "ko-kp"
	PagesPageLanguageKoKr   PagesPageLanguage = "ko-kr"
	PagesPageLanguageKok    PagesPageLanguage = "kok"
	PagesPageLanguageKokIn  PagesPageLanguage = "kok-in"
	PagesPageLanguageKr     PagesPageLanguage = "kr"
	PagesPageLanguageKs     PagesPageLanguage = "ks"
	PagesPageLanguageKsIn   PagesPageLanguage = "ks-in"
	PagesPageLanguageKsb    PagesPageLanguage = "ksb"
	PagesPageLanguageKsbTz  PagesPageLanguage = "ksb-tz"
	PagesPageLanguageKsf    PagesPageLanguage = "ksf"
	PagesPageLanguageKsfCm  PagesPageLanguage = "ksf-cm"
	PagesPageLanguageKsh    PagesPageLanguage = "ksh"
	PagesPageLanguageKshDe  PagesPageLanguage = "ksh-de"
	PagesPageLanguageKu     PagesPageLanguage = "ku"
	PagesPageLanguageKuTr   PagesPageLanguage = "ku-tr"
	PagesPageLanguageKv     PagesPageLanguage = "kv"
	PagesPageLanguageKw     PagesPageLanguage = "kw"
	PagesPageLanguageKwGB   PagesPageLanguage = "kw-gb"
	PagesPageLanguageKy     PagesPageLanguage = "ky"
	PagesPageLanguageKyKg   PagesPageLanguage = "ky-kg"
	PagesPageLanguageLa     PagesPageLanguage = "la"
	PagesPageLanguageLag    PagesPageLanguage = "lag"
	PagesPageLanguageLagTz  PagesPageLanguage = "lag-tz"
	PagesPageLanguageLb     PagesPageLanguage = "lb"
	PagesPageLanguageLbLu   PagesPageLanguage = "lb-lu"
	PagesPageLanguageLg     PagesPageLanguage = "lg"
	PagesPageLanguageLgUg   PagesPageLanguage = "lg-ug"
	PagesPageLanguageLi     PagesPageLanguage = "li"
	PagesPageLanguageLkt    PagesPageLanguage = "lkt"
	PagesPageLanguageLktUs  PagesPageLanguage = "lkt-us"
	PagesPageLanguageLn     PagesPageLanguage = "ln"
	PagesPageLanguageLnAo   PagesPageLanguage = "ln-ao"
	PagesPageLanguageLnCd   PagesPageLanguage = "ln-cd"
	PagesPageLanguageLnCf   PagesPageLanguage = "ln-cf"
	PagesPageLanguageLnCg   PagesPageLanguage = "ln-cg"
	PagesPageLanguageLo     PagesPageLanguage = "lo"
	PagesPageLanguageLoLa   PagesPageLanguage = "lo-la"
	PagesPageLanguageLrc    PagesPageLanguage = "lrc"
	PagesPageLanguageLrcIq  PagesPageLanguage = "lrc-iq"
	PagesPageLanguageLrcIr  PagesPageLanguage = "lrc-ir"
	PagesPageLanguageLt     PagesPageLanguage = "lt"
	PagesPageLanguageLtLt   PagesPageLanguage = "lt-lt"
	PagesPageLanguageLu     PagesPageLanguage = "lu"
	PagesPageLanguageLuCd   PagesPageLanguage = "lu-cd"
	PagesPageLanguageLuo    PagesPageLanguage = "luo"
	PagesPageLanguageLuoKe  PagesPageLanguage = "luo-ke"
	PagesPageLanguageLuy    PagesPageLanguage = "luy"
	PagesPageLanguageLuyKe  PagesPageLanguage = "luy-ke"
	PagesPageLanguageLv     PagesPageLanguage = "lv"
	PagesPageLanguageLvLv   PagesPageLanguage = "lv-lv"
	PagesPageLanguageMai    PagesPageLanguage = "mai"
	PagesPageLanguageMaiIn  PagesPageLanguage = "mai-in"
	PagesPageLanguageMas    PagesPageLanguage = "mas"
	PagesPageLanguageMasKe  PagesPageLanguage = "mas-ke"
	PagesPageLanguageMasTz  PagesPageLanguage = "mas-tz"
	PagesPageLanguageMdf    PagesPageLanguage = "mdf"
	PagesPageLanguageMdfRu  PagesPageLanguage = "mdf-ru"
	PagesPageLanguageMer    PagesPageLanguage = "mer"
	PagesPageLanguageMerKe  PagesPageLanguage = "mer-ke"
	PagesPageLanguageMfe    PagesPageLanguage = "mfe"
	PagesPageLanguageMfeMu  PagesPageLanguage = "mfe-mu"
	PagesPageLanguageMg     PagesPageLanguage = "mg"
	PagesPageLanguageMgMg   PagesPageLanguage = "mg-mg"
	PagesPageLanguageMgh    PagesPageLanguage = "mgh"
	PagesPageLanguageMghMz  PagesPageLanguage = "mgh-mz"
	PagesPageLanguageMgo    PagesPageLanguage = "mgo"
	PagesPageLanguageMgoCm  PagesPageLanguage = "mgo-cm"
	PagesPageLanguageMh     PagesPageLanguage = "mh"
	PagesPageLanguageMi     PagesPageLanguage = "mi"
	PagesPageLanguageMiNz   PagesPageLanguage = "mi-nz"
	PagesPageLanguageMk     PagesPageLanguage = "mk"
	PagesPageLanguageMkMk   PagesPageLanguage = "mk-mk"
	PagesPageLanguageMl     PagesPageLanguage = "ml"
	PagesPageLanguageMlIn   PagesPageLanguage = "ml-in"
	PagesPageLanguageMn     PagesPageLanguage = "mn"
	PagesPageLanguageMnMn   PagesPageLanguage = "mn-mn"
	PagesPageLanguageMni    PagesPageLanguage = "mni"
	PagesPageLanguageMniIn  PagesPageLanguage = "mni-in"
	PagesPageLanguageMr     PagesPageLanguage = "mr"
	PagesPageLanguageMrIn   PagesPageLanguage = "mr-in"
	PagesPageLanguageMs     PagesPageLanguage = "ms"
	PagesPageLanguageMsBn   PagesPageLanguage = "ms-bn"
	PagesPageLanguageMsID   PagesPageLanguage = "ms-id"
	PagesPageLanguageMsMy   PagesPageLanguage = "ms-my"
	PagesPageLanguageMsSg   PagesPageLanguage = "ms-sg"
	PagesPageLanguageMt     PagesPageLanguage = "mt"
	PagesPageLanguageMtMt   PagesPageLanguage = "mt-mt"
	PagesPageLanguageMua    PagesPageLanguage = "mua"
	PagesPageLanguageMuaCm  PagesPageLanguage = "mua-cm"
	PagesPageLanguageMy     PagesPageLanguage = "my"
	PagesPageLanguageMyMm   PagesPageLanguage = "my-mm"
	PagesPageLanguageMzn    PagesPageLanguage = "mzn"
	PagesPageLanguageMznIr  PagesPageLanguage = "mzn-ir"
	PagesPageLanguageNa     PagesPageLanguage = "na"
	PagesPageLanguageNaq    PagesPageLanguage = "naq"
	PagesPageLanguageNaqNa  PagesPageLanguage = "naq-na"
	PagesPageLanguageNb     PagesPageLanguage = "nb"
	PagesPageLanguageNbNo   PagesPageLanguage = "nb-no"
	PagesPageLanguageNbSj   PagesPageLanguage = "nb-sj"
	PagesPageLanguageNd     PagesPageLanguage = "nd"
	PagesPageLanguageNdZw   PagesPageLanguage = "nd-zw"
	PagesPageLanguageNds    PagesPageLanguage = "nds"
	PagesPageLanguageNdsDe  PagesPageLanguage = "nds-de"
	PagesPageLanguageNdsNl  PagesPageLanguage = "nds-nl"
	PagesPageLanguageNe     PagesPageLanguage = "ne"
	PagesPageLanguageNeIn   PagesPageLanguage = "ne-in"
	PagesPageLanguageNeNp   PagesPageLanguage = "ne-np"
	PagesPageLanguageNg     PagesPageLanguage = "ng"
	PagesPageLanguageNl     PagesPageLanguage = "nl"
	PagesPageLanguageNlAw   PagesPageLanguage = "nl-aw"
	PagesPageLanguageNlBe   PagesPageLanguage = "nl-be"
	PagesPageLanguageNlBq   PagesPageLanguage = "nl-bq"
	PagesPageLanguageNlCh   PagesPageLanguage = "nl-ch"
	PagesPageLanguageNlCw   PagesPageLanguage = "nl-cw"
	PagesPageLanguageNlLu   PagesPageLanguage = "nl-lu"
	PagesPageLanguageNlNl   PagesPageLanguage = "nl-nl"
	PagesPageLanguageNlSr   PagesPageLanguage = "nl-sr"
	PagesPageLanguageNlSx   PagesPageLanguage = "nl-sx"
	PagesPageLanguageNmg    PagesPageLanguage = "nmg"
	PagesPageLanguageNmgCm  PagesPageLanguage = "nmg-cm"
	PagesPageLanguageNn     PagesPageLanguage = "nn"
	PagesPageLanguageNnNo   PagesPageLanguage = "nn-no"
	PagesPageLanguageNnh    PagesPageLanguage = "nnh"
	PagesPageLanguageNnhCm  PagesPageLanguage = "nnh-cm"
	PagesPageLanguageNo     PagesPageLanguage = "no"
	PagesPageLanguageNoNo   PagesPageLanguage = "no-no"
	PagesPageLanguageNr     PagesPageLanguage = "nr"
	PagesPageLanguageNus    PagesPageLanguage = "nus"
	PagesPageLanguageNusSS  PagesPageLanguage = "nus-ss"
	PagesPageLanguageNv     PagesPageLanguage = "nv"
	PagesPageLanguageNy     PagesPageLanguage = "ny"
	PagesPageLanguageNyn    PagesPageLanguage = "nyn"
	PagesPageLanguageNynUg  PagesPageLanguage = "nyn-ug"
	PagesPageLanguageOc     PagesPageLanguage = "oc"
	PagesPageLanguageOcEs   PagesPageLanguage = "oc-es"
	PagesPageLanguageOcFr   PagesPageLanguage = "oc-fr"
	PagesPageLanguageOj     PagesPageLanguage = "oj"
	PagesPageLanguageOm     PagesPageLanguage = "om"
	PagesPageLanguageOmEt   PagesPageLanguage = "om-et"
	PagesPageLanguageOmKe   PagesPageLanguage = "om-ke"
	PagesPageLanguageOr     PagesPageLanguage = "or"
	PagesPageLanguageOrIn   PagesPageLanguage = "or-in"
	PagesPageLanguageOs     PagesPageLanguage = "os"
	PagesPageLanguageOsGe   PagesPageLanguage = "os-ge"
	PagesPageLanguageOsRu   PagesPageLanguage = "os-ru"
	PagesPageLanguagePa     PagesPageLanguage = "pa"
	PagesPageLanguagePaIn   PagesPageLanguage = "pa-in"
	PagesPageLanguagePaPk   PagesPageLanguage = "pa-pk"
	PagesPageLanguagePcm    PagesPageLanguage = "pcm"
	PagesPageLanguagePcmNg  PagesPageLanguage = "pcm-ng"
	PagesPageLanguagePi     PagesPageLanguage = "pi"
	PagesPageLanguagePis    PagesPageLanguage = "pis"
	PagesPageLanguagePisSb  PagesPageLanguage = "pis-sb"
	PagesPageLanguagePl     PagesPageLanguage = "pl"
	PagesPageLanguagePlPl   PagesPageLanguage = "pl-pl"
	PagesPageLanguagePrg    PagesPageLanguage = "prg"
	PagesPageLanguagePrg001 PagesPageLanguage = "prg-001"
	PagesPageLanguagePs     PagesPageLanguage = "ps"
	PagesPageLanguagePsAf   PagesPageLanguage = "ps-af"
	PagesPageLanguagePsPk   PagesPageLanguage = "ps-pk"
	PagesPageLanguagePt     PagesPageLanguage = "pt"
	PagesPageLanguagePtAo   PagesPageLanguage = "pt-ao"
	PagesPageLanguagePtBr   PagesPageLanguage = "pt-br"
	PagesPageLanguagePtCh   PagesPageLanguage = "pt-ch"
	PagesPageLanguagePtCv   PagesPageLanguage = "pt-cv"
	PagesPageLanguagePtGq   PagesPageLanguage = "pt-gq"
	PagesPageLanguagePtGw   PagesPageLanguage = "pt-gw"
	PagesPageLanguagePtLu   PagesPageLanguage = "pt-lu"
	PagesPageLanguagePtMo   PagesPageLanguage = "pt-mo"
	PagesPageLanguagePtMz   PagesPageLanguage = "pt-mz"
	PagesPageLanguagePtPt   PagesPageLanguage = "pt-pt"
	PagesPageLanguagePtSt   PagesPageLanguage = "pt-st"
	PagesPageLanguagePtTl   PagesPageLanguage = "pt-tl"
	PagesPageLanguageQu     PagesPageLanguage = "qu"
	PagesPageLanguageQuBo   PagesPageLanguage = "qu-bo"
	PagesPageLanguageQuEc   PagesPageLanguage = "qu-ec"
	PagesPageLanguageQuPe   PagesPageLanguage = "qu-pe"
	PagesPageLanguageRaj    PagesPageLanguage = "raj"
	PagesPageLanguageRajIn  PagesPageLanguage = "raj-in"
	PagesPageLanguageRm     PagesPageLanguage = "rm"
	PagesPageLanguageRmCh   PagesPageLanguage = "rm-ch"
	PagesPageLanguageRn     PagesPageLanguage = "rn"
	PagesPageLanguageRnBi   PagesPageLanguage = "rn-bi"
	PagesPageLanguageRo     PagesPageLanguage = "ro"
	PagesPageLanguageRoMd   PagesPageLanguage = "ro-md"
	PagesPageLanguageRoRo   PagesPageLanguage = "ro-ro"
	PagesPageLanguageRof    PagesPageLanguage = "rof"
	PagesPageLanguageRofTz  PagesPageLanguage = "rof-tz"
	PagesPageLanguageRu     PagesPageLanguage = "ru"
	PagesPageLanguageRuBy   PagesPageLanguage = "ru-by"
	PagesPageLanguageRuKg   PagesPageLanguage = "ru-kg"
	PagesPageLanguageRuKz   PagesPageLanguage = "ru-kz"
	PagesPageLanguageRuMd   PagesPageLanguage = "ru-md"
	PagesPageLanguageRuRu   PagesPageLanguage = "ru-ru"
	PagesPageLanguageRuUa   PagesPageLanguage = "ru-ua"
	PagesPageLanguageRw     PagesPageLanguage = "rw"
	PagesPageLanguageRwRw   PagesPageLanguage = "rw-rw"
	PagesPageLanguageRwk    PagesPageLanguage = "rwk"
	PagesPageLanguageRwkTz  PagesPageLanguage = "rwk-tz"
	PagesPageLanguageSa     PagesPageLanguage = "sa"
	PagesPageLanguageSaIn   PagesPageLanguage = "sa-in"
	PagesPageLanguageSah    PagesPageLanguage = "sah"
	PagesPageLanguageSahRu  PagesPageLanguage = "sah-ru"
	PagesPageLanguageSaq    PagesPageLanguage = "saq"
	PagesPageLanguageSaqKe  PagesPageLanguage = "saq-ke"
	PagesPageLanguageSat    PagesPageLanguage = "sat"
	PagesPageLanguageSatIn  PagesPageLanguage = "sat-in"
	PagesPageLanguageSbp    PagesPageLanguage = "sbp"
	PagesPageLanguageSbpTz  PagesPageLanguage = "sbp-tz"
	PagesPageLanguageSc     PagesPageLanguage = "sc"
	PagesPageLanguageScIt   PagesPageLanguage = "sc-it"
	PagesPageLanguageSd     PagesPageLanguage = "sd"
	PagesPageLanguageSdIn   PagesPageLanguage = "sd-in"
	PagesPageLanguageSdPk   PagesPageLanguage = "sd-pk"
	PagesPageLanguageSe     PagesPageLanguage = "se"
	PagesPageLanguageSeFi   PagesPageLanguage = "se-fi"
	PagesPageLanguageSeNo   PagesPageLanguage = "se-no"
	PagesPageLanguageSeSe   PagesPageLanguage = "se-se"
	PagesPageLanguageSeh    PagesPageLanguage = "seh"
	PagesPageLanguageSehMz  PagesPageLanguage = "seh-mz"
	PagesPageLanguageSes    PagesPageLanguage = "ses"
	PagesPageLanguageSesMl  PagesPageLanguage = "ses-ml"
	PagesPageLanguageSg     PagesPageLanguage = "sg"
	PagesPageLanguageSgCf   PagesPageLanguage = "sg-cf"
	PagesPageLanguageShi    PagesPageLanguage = "shi"
	PagesPageLanguageShiMa  PagesPageLanguage = "shi-ma"
	PagesPageLanguageSi     PagesPageLanguage = "si"
	PagesPageLanguageSiLk   PagesPageLanguage = "si-lk"
	PagesPageLanguageSk     PagesPageLanguage = "sk"
	PagesPageLanguageSkSk   PagesPageLanguage = "sk-sk"
	PagesPageLanguageSl     PagesPageLanguage = "sl"
	PagesPageLanguageSlSi   PagesPageLanguage = "sl-si"
	PagesPageLanguageSm     PagesPageLanguage = "sm"
	PagesPageLanguageSmn    PagesPageLanguage = "smn"
	PagesPageLanguageSmnFi  PagesPageLanguage = "smn-fi"
	PagesPageLanguageSMS    PagesPageLanguage = "sms"
	PagesPageLanguageSMSFi  PagesPageLanguage = "sms-fi"
	PagesPageLanguageSn     PagesPageLanguage = "sn"
	PagesPageLanguageSnZw   PagesPageLanguage = "sn-zw"
	PagesPageLanguageSo     PagesPageLanguage = "so"
	PagesPageLanguageSoDj   PagesPageLanguage = "so-dj"
	PagesPageLanguageSoEt   PagesPageLanguage = "so-et"
	PagesPageLanguageSoKe   PagesPageLanguage = "so-ke"
	PagesPageLanguageSoSo   PagesPageLanguage = "so-so"
	PagesPageLanguageSq     PagesPageLanguage = "sq"
	PagesPageLanguageSqAl   PagesPageLanguage = "sq-al"
	PagesPageLanguageSqMk   PagesPageLanguage = "sq-mk"
	PagesPageLanguageSqXk   PagesPageLanguage = "sq-xk"
	PagesPageLanguageSr     PagesPageLanguage = "sr"
	PagesPageLanguageSrBa   PagesPageLanguage = "sr-ba"
	PagesPageLanguageSrCs   PagesPageLanguage = "sr-cs"
	PagesPageLanguageSrMe   PagesPageLanguage = "sr-me"
	PagesPageLanguageSrRs   PagesPageLanguage = "sr-rs"
	PagesPageLanguageSrXk   PagesPageLanguage = "sr-xk"
	PagesPageLanguageSS     PagesPageLanguage = "ss"
	PagesPageLanguageSt     PagesPageLanguage = "st"
	PagesPageLanguageSu     PagesPageLanguage = "su"
	PagesPageLanguageSuID   PagesPageLanguage = "su-id"
	PagesPageLanguageSv     PagesPageLanguage = "sv"
	PagesPageLanguageSvAx   PagesPageLanguage = "sv-ax"
	PagesPageLanguageSvFi   PagesPageLanguage = "sv-fi"
	PagesPageLanguageSvSe   PagesPageLanguage = "sv-se"
	PagesPageLanguageSw     PagesPageLanguage = "sw"
	PagesPageLanguageSwCd   PagesPageLanguage = "sw-cd"
	PagesPageLanguageSwKe   PagesPageLanguage = "sw-ke"
	PagesPageLanguageSwTz   PagesPageLanguage = "sw-tz"
	PagesPageLanguageSwUg   PagesPageLanguage = "sw-ug"
	PagesPageLanguageSy     PagesPageLanguage = "sy"
	PagesPageLanguageTa     PagesPageLanguage = "ta"
	PagesPageLanguageTaIn   PagesPageLanguage = "ta-in"
	PagesPageLanguageTaLk   PagesPageLanguage = "ta-lk"
	PagesPageLanguageTaMy   PagesPageLanguage = "ta-my"
	PagesPageLanguageTaSg   PagesPageLanguage = "ta-sg"
	PagesPageLanguageTe     PagesPageLanguage = "te"
	PagesPageLanguageTeIn   PagesPageLanguage = "te-in"
	PagesPageLanguageTeo    PagesPageLanguage = "teo"
	PagesPageLanguageTeoKe  PagesPageLanguage = "teo-ke"
	PagesPageLanguageTeoUg  PagesPageLanguage = "teo-ug"
	PagesPageLanguageTg     PagesPageLanguage = "tg"
	PagesPageLanguageTgTj   PagesPageLanguage = "tg-tj"
	PagesPageLanguageTh     PagesPageLanguage = "th"
	PagesPageLanguageThTh   PagesPageLanguage = "th-th"
	PagesPageLanguageTi     PagesPageLanguage = "ti"
	PagesPageLanguageTiEr   PagesPageLanguage = "ti-er"
	PagesPageLanguageTiEt   PagesPageLanguage = "ti-et"
	PagesPageLanguageTk     PagesPageLanguage = "tk"
	PagesPageLanguageTkTm   PagesPageLanguage = "tk-tm"
	PagesPageLanguageTl     PagesPageLanguage = "tl"
	PagesPageLanguageTn     PagesPageLanguage = "tn"
	PagesPageLanguageTo     PagesPageLanguage = "to"
	PagesPageLanguageToTo   PagesPageLanguage = "to-to"
	PagesPageLanguageTok    PagesPageLanguage = "tok"
	PagesPageLanguageTok001 PagesPageLanguage = "tok-001"
	PagesPageLanguageTr     PagesPageLanguage = "tr"
	PagesPageLanguageTrCy   PagesPageLanguage = "tr-cy"
	PagesPageLanguageTrTr   PagesPageLanguage = "tr-tr"
	PagesPageLanguageTs     PagesPageLanguage = "ts"
	PagesPageLanguageTt     PagesPageLanguage = "tt"
	PagesPageLanguageTtRu   PagesPageLanguage = "tt-ru"
	PagesPageLanguageTw     PagesPageLanguage = "tw"
	PagesPageLanguageTwq    PagesPageLanguage = "twq"
	PagesPageLanguageTwqNe  PagesPageLanguage = "twq-ne"
	PagesPageLanguageTy     PagesPageLanguage = "ty"
	PagesPageLanguageTzm    PagesPageLanguage = "tzm"
	PagesPageLanguageTzmMa  PagesPageLanguage = "tzm-ma"
	PagesPageLanguageUg     PagesPageLanguage = "ug"
	PagesPageLanguageUgCn   PagesPageLanguage = "ug-cn"
	PagesPageLanguageUk     PagesPageLanguage = "uk"
	PagesPageLanguageUkUa   PagesPageLanguage = "uk-ua"
	PagesPageLanguageUr     PagesPageLanguage = "ur"
	PagesPageLanguageUrIn   PagesPageLanguage = "ur-in"
	PagesPageLanguageUrPk   PagesPageLanguage = "ur-pk"
	PagesPageLanguageUz     PagesPageLanguage = "uz"
	PagesPageLanguageUzAf   PagesPageLanguage = "uz-af"
	PagesPageLanguageUzUz   PagesPageLanguage = "uz-uz"
	PagesPageLanguageVai    PagesPageLanguage = "vai"
	PagesPageLanguageVaiLr  PagesPageLanguage = "vai-lr"
	PagesPageLanguageVe     PagesPageLanguage = "ve"
	PagesPageLanguageVi     PagesPageLanguage = "vi"
	PagesPageLanguageViVn   PagesPageLanguage = "vi-vn"
	PagesPageLanguageVo     PagesPageLanguage = "vo"
	PagesPageLanguageVo001  PagesPageLanguage = "vo-001"
	PagesPageLanguageVun    PagesPageLanguage = "vun"
	PagesPageLanguageVunTz  PagesPageLanguage = "vun-tz"
	PagesPageLanguageWa     PagesPageLanguage = "wa"
	PagesPageLanguageWae    PagesPageLanguage = "wae"
	PagesPageLanguageWaeCh  PagesPageLanguage = "wae-ch"
	PagesPageLanguageWo     PagesPageLanguage = "wo"
	PagesPageLanguageWoSn   PagesPageLanguage = "wo-sn"
	PagesPageLanguageXh     PagesPageLanguage = "xh"
	PagesPageLanguageXhZa   PagesPageLanguage = "xh-za"
	PagesPageLanguageXog    PagesPageLanguage = "xog"
	PagesPageLanguageXogUg  PagesPageLanguage = "xog-ug"
	PagesPageLanguageYav    PagesPageLanguage = "yav"
	PagesPageLanguageYavCm  PagesPageLanguage = "yav-cm"
	PagesPageLanguageYi     PagesPageLanguage = "yi"
	PagesPageLanguageYi001  PagesPageLanguage = "yi-001"
	PagesPageLanguageYo     PagesPageLanguage = "yo"
	PagesPageLanguageYoBj   PagesPageLanguage = "yo-bj"
	PagesPageLanguageYoNg   PagesPageLanguage = "yo-ng"
	PagesPageLanguageYrl    PagesPageLanguage = "yrl"
	PagesPageLanguageYrlBr  PagesPageLanguage = "yrl-br"
	PagesPageLanguageYrlCo  PagesPageLanguage = "yrl-co"
	PagesPageLanguageYrlVe  PagesPageLanguage = "yrl-ve"
	PagesPageLanguageYue    PagesPageLanguage = "yue"
	PagesPageLanguageYueCn  PagesPageLanguage = "yue-cn"
	PagesPageLanguageYueHk  PagesPageLanguage = "yue-hk"
	PagesPageLanguageZa     PagesPageLanguage = "za"
	PagesPageLanguageZgh    PagesPageLanguage = "zgh"
	PagesPageLanguageZghMa  PagesPageLanguage = "zgh-ma"
	PagesPageLanguageZh     PagesPageLanguage = "zh"
	PagesPageLanguageZhCn   PagesPageLanguage = "zh-cn"
	PagesPageLanguageZhHans PagesPageLanguage = "zh-hans"
	PagesPageLanguageZhHant PagesPageLanguage = "zh-hant"
	PagesPageLanguageZhHk   PagesPageLanguage = "zh-hk"
	PagesPageLanguageZhMo   PagesPageLanguage = "zh-mo"
	PagesPageLanguageZhSg   PagesPageLanguage = "zh-sg"
	PagesPageLanguageZhTw   PagesPageLanguage = "zh-tw"
	PagesPageLanguageZu     PagesPageLanguage = "zu"
	PagesPageLanguageZuZa   PagesPageLanguage = "zu-za"
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
type PagesPageParam struct {
	// The unique ID of the page.
	ID string `json:"id" api:"required"`
	// The status of the AB test associated with this page, if applicable
	//
	// Any of "automated_loser_variant", "automated_master", "automated_variant",
	// "loser_variant", "mab_master", "mab_variant", "master", "variant".
	AbStatus PagesPageAbStatus `json:"abStatus,omitzero" api:"required"`
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
	ContentTypeCategory PagesPageContentTypeCategory `json:"contentTypeCategory,omitzero" api:"required"`
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
	CurrentState PagesPageCurrentState `json:"currentState,omitzero" api:"required"`
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
	Language PagesPageLanguage `json:"language,omitzero" api:"required"`
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

func (r PagesPageParam) MarshalJSON() (data []byte, err error) {
	type shadow PagesPageParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PagesPageParam) UnmarshalJSON(data []byte) error {
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
