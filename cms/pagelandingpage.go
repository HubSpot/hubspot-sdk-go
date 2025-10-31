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

	"github.com/stainless-sdks/hubspot-sdk-go/internal/apiquery"
	shimjson "github.com/stainless-sdks/hubspot-sdk-go/internal/encoding/json"
	"github.com/stainless-sdks/hubspot-sdk-go/internal/requestconfig"
	"github.com/stainless-sdks/hubspot-sdk-go/option"
	"github.com/stainless-sdks/hubspot-sdk-go/packages/pagination"
	"github.com/stainless-sdks/hubspot-sdk-go/packages/param"
	"github.com/stainless-sdks/hubspot-sdk-go/shared"
)

// PageLandingPageService contains methods and other services that help with
// interacting with the Hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPageLandingPageService] method instead.
type PageLandingPageService struct {
	Options []option.RequestOption
}

// NewPageLandingPageService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewPageLandingPageService(opts ...option.RequestOption) (r PageLandingPageService) {
	r = PageLandingPageService{}
	r.Options = opts
	return
}

// Create a new Landing Page
func (r *PageLandingPageService) New(ctx context.Context, body PageLandingPageNewParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "cms/v3/pages/landing-pages"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Sparse updates a single Landing Page object identified by the id in the path.
// You only need to specify the column values that you are modifying.
func (r *PageLandingPageService) Update(ctx context.Context, objectID string, params PageLandingPageUpdateParams, opts ...option.RequestOption) (res *Page, err error) {
	opts = slices.Concat(r.Options, opts)
	if objectID == "" {
		err = errors.New("missing required objectId parameter")
		return
	}
	path := fmt.Sprintf("cms/v3/pages/landing-pages/%s", objectID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return
}

// Get the list of landing pages. Supports paging and filtering. This method would
// be useful for an integration that examined these models and used an external
// service to suggest edits.
func (r *PageLandingPageService) List(ctx context.Context, query PageLandingPageListParams, opts ...option.RequestOption) (res *pagination.Page[Page], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "cms/v3/pages/landing-pages"
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

// Get the list of landing pages. Supports paging and filtering. This method would
// be useful for an integration that examined these models and used an external
// service to suggest edits.
func (r *PageLandingPageService) ListAutoPaging(ctx context.Context, query PageLandingPageListParams, opts ...option.RequestOption) *pagination.PageAutoPager[Page] {
	return pagination.NewPageAutoPager(r.List(ctx, query, opts...))
}

// Delete the Landing Page object identified by the id in the path.
func (r *PageLandingPageService) Delete(ctx context.Context, objectID string, body PageLandingPageDeleteParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if objectID == "" {
		err = errors.New("missing required objectId parameter")
		return
	}
	path := fmt.Sprintf("cms/v3/pages/landing-pages/%s", objectID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, nil, opts...)
	return
}

// Attach a landing page to a multi-language group.
func (r *PageLandingPageService) AttachToLangGroup(ctx context.Context, body PageLandingPageAttachToLangGroupParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "cms/v3/pages/landing-pages/multi-language/attach-to-lang-group"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Clone a Landing Page
func (r *PageLandingPageService) Clone(ctx context.Context, body PageLandingPageCloneParams, opts ...option.RequestOption) (res *Page, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "cms/v3/pages/landing-pages/clone"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Create a new A/B test variation based on the information provided in the request
// body.
func (r *PageLandingPageService) NewAbTestVariation(ctx context.Context, body PageLandingPageNewAbTestVariationParams, opts ...option.RequestOption) (res *Page, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "cms/v3/pages/landing-pages/ab-test/create-variation"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Create the Landing Page objects detailed in the request body.
func (r *PageLandingPageService) NewBatch(ctx context.Context, body PageLandingPageNewBatchParams, opts ...option.RequestOption) (res *BatchResponsePage, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "cms/v3/pages/landing-pages/batch/create"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Create a new Folder
func (r *PageLandingPageService) NewFolder(ctx context.Context, body PageLandingPageNewFolderParams, opts ...option.RequestOption) (res *ContentFolder, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "cms/v3/pages/landing-pages/folders"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Create the Folder objects detailed in the request body.
func (r *PageLandingPageService) NewFoldersBatch(ctx context.Context, body PageLandingPageNewFoldersBatchParams, opts ...option.RequestOption) (res *BatchResponseContentFolder, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "cms/v3/pages/landing-pages/folders/batch/create"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Create a new language variation from an existing landing page
func (r *PageLandingPageService) NewLanguageVariation(ctx context.Context, body PageLandingPageNewLanguageVariationParams, opts ...option.RequestOption) (res *Page, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "cms/v3/pages/landing-pages/multi-language/create-language-variation"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Delete the Landing Page objects identified in the request body. Note: This is
// not the same as the dashboard `archive` function. To perform a dashboard
// `archive` send an normal update with the `archivedInDashboard` field set to
// true.
func (r *PageLandingPageService) DeleteBatch(ctx context.Context, body PageLandingPageDeleteBatchParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "cms/v3/pages/landing-pages/batch/archive"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Delete the Folder object identified by the id in the path.
func (r *PageLandingPageService) DeleteFolder(ctx context.Context, objectID string, body PageLandingPageDeleteFolderParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if objectID == "" {
		err = errors.New("missing required objectId parameter")
		return
	}
	path := fmt.Sprintf("cms/v3/pages/landing-pages/folders/%s", objectID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, nil, opts...)
	return
}

// Delete the Folder objects identified in the request body.
func (r *PageLandingPageService) DeleteFoldersBatch(ctx context.Context, body PageLandingPageDeleteFoldersBatchParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "cms/v3/pages/landing-pages/folders/batch/archive"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Detach a landing page from a multi-language group.
func (r *PageLandingPageService) DetachFromLangGroup(ctx context.Context, body PageLandingPageDetachFromLangGroupParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "cms/v3/pages/landing-pages/multi-language/detach-from-lang-group"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// End an active A/B test and designate a winner.
func (r *PageLandingPageService) EndAbTest(ctx context.Context, body PageLandingPageEndAbTestParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "cms/v3/pages/landing-pages/ab-test/end"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Retrieve the Landing Page object identified by the id in the path.
func (r *PageLandingPageService) Get(ctx context.Context, objectID string, query PageLandingPageGetParams, opts ...option.RequestOption) (res *Page, err error) {
	opts = slices.Concat(r.Options, opts)
	if objectID == "" {
		err = errors.New("missing required objectId parameter")
		return
	}
	path := fmt.Sprintf("cms/v3/pages/landing-pages/%s", objectID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Retrieve the Landing Page objects identified in the request body.
func (r *PageLandingPageService) GetBatch(ctx context.Context, params PageLandingPageGetBatchParams, opts ...option.RequestOption) (res *BatchResponsePage, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "cms/v3/pages/landing-pages/batch/read"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Retrieve the full draft version of the Landing Page.
func (r *PageLandingPageService) GetDraft(ctx context.Context, objectID string, opts ...option.RequestOption) (res *Page, err error) {
	opts = slices.Concat(r.Options, opts)
	if objectID == "" {
		err = errors.New("missing required objectId parameter")
		return
	}
	path := fmt.Sprintf("cms/v3/pages/landing-pages/%s/draft", objectID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Retrieve the Folder object identified by the id in the path.
func (r *PageLandingPageService) GetFolder(ctx context.Context, objectID string, query PageLandingPageGetFolderParams, opts ...option.RequestOption) (res *ContentFolder, err error) {
	opts = slices.Concat(r.Options, opts)
	if objectID == "" {
		err = errors.New("missing required objectId parameter")
		return
	}
	path := fmt.Sprintf("cms/v3/pages/landing-pages/folders/%s", objectID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Retrieves a previous version of a Folder
func (r *PageLandingPageService) GetFolderRevision(ctx context.Context, revisionID string, query PageLandingPageGetFolderRevisionParams, opts ...option.RequestOption) (res *VersionContentFolder, err error) {
	opts = slices.Concat(r.Options, opts)
	if query.ObjectID == "" {
		err = errors.New("missing required objectId parameter")
		return
	}
	if revisionID == "" {
		err = errors.New("missing required revisionId parameter")
		return
	}
	path := fmt.Sprintf("cms/v3/pages/landing-pages/folders/%s/revisions/%s", query.ObjectID, revisionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update the Folder objects identified in the request body.
func (r *PageLandingPageService) GetFoldersBatch(ctx context.Context, params PageLandingPageGetFoldersBatchParams, opts ...option.RequestOption) (res *BatchResponseContentFolder, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "cms/v3/pages/landing-pages/folders/batch/read"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Retrieves a previous version of a Landing Page
func (r *PageLandingPageService) GetRevision(ctx context.Context, revisionID string, query PageLandingPageGetRevisionParams, opts ...option.RequestOption) (res *VersionPage, err error) {
	opts = slices.Concat(r.Options, opts)
	if query.ObjectID == "" {
		err = errors.New("missing required objectId parameter")
		return
	}
	if revisionID == "" {
		err = errors.New("missing required revisionId parameter")
		return
	}
	path := fmt.Sprintf("cms/v3/pages/landing-pages/%s/revisions/%s", query.ObjectID, revisionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Retrieves all the previous versions of a Folder.
func (r *PageLandingPageService) ListFolderRevisions(ctx context.Context, objectID string, query PageLandingPageListFolderRevisionsParams, opts ...option.RequestOption) (res *CollectionResponseWithTotalVersionContentFolder, err error) {
	opts = slices.Concat(r.Options, opts)
	if objectID == "" {
		err = errors.New("missing required objectId parameter")
		return
	}
	path := fmt.Sprintf("cms/v3/pages/landing-pages/folders/%s/revisions", objectID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Get the list of Landing Page Folders. Supports paging and filtering. This method
// would be useful for an integration that examined these models and used an
// external service to suggest edits.
func (r *PageLandingPageService) ListFolders(ctx context.Context, query PageLandingPageListFoldersParams, opts ...option.RequestOption) (res *CollectionResponseWithTotalContentFolderForwardPaging, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "cms/v3/pages/landing-pages/folders"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Retrieves all the previous versions of a Landing Page.
func (r *PageLandingPageService) ListRevisions(ctx context.Context, objectID string, query PageLandingPageListRevisionsParams, opts ...option.RequestOption) (res *CollectionResponseWithTotalVersionPage, err error) {
	opts = slices.Concat(r.Options, opts)
	if objectID == "" {
		err = errors.New("missing required objectId parameter")
		return
	}
	path := fmt.Sprintf("cms/v3/pages/landing-pages/%s/revisions", objectID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Take any changes from the draft version of the Landing Page and apply them to
// the live version.
func (r *PageLandingPageService) PublishDraft(ctx context.Context, objectID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if objectID == "" {
		err = errors.New("missing required objectId parameter")
		return
	}
	path := fmt.Sprintf("cms/v3/pages/landing-pages/%s/draft/push-live", objectID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, nil, opts...)
	return
}

// Rerun a previous A/B test.
func (r *PageLandingPageService) RerunAbTest(ctx context.Context, body PageLandingPageRerunAbTestParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "cms/v3/pages/landing-pages/ab-test/rerun"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Discards any edits and resets the draft to the live version.
func (r *PageLandingPageService) ResetDraft(ctx context.Context, objectID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if objectID == "" {
		err = errors.New("missing required objectId parameter")
		return
	}
	path := fmt.Sprintf("cms/v3/pages/landing-pages/%s/draft/reset", objectID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, nil, opts...)
	return
}

// Takes a specified version of a Folder and restores it.
func (r *PageLandingPageService) RestoreFolderRevision(ctx context.Context, revisionID string, body PageLandingPageRestoreFolderRevisionParams, opts ...option.RequestOption) (res *ContentFolder, err error) {
	opts = slices.Concat(r.Options, opts)
	if body.ObjectID == "" {
		err = errors.New("missing required objectId parameter")
		return
	}
	if revisionID == "" {
		err = errors.New("missing required revisionId parameter")
		return
	}
	path := fmt.Sprintf("cms/v3/pages/landing-pages/folders/%s/revisions/%s/restore", body.ObjectID, revisionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Takes a specified version of a Landing Page and restores it.
func (r *PageLandingPageService) RestoreRevision(ctx context.Context, revisionID string, body PageLandingPageRestoreRevisionParams, opts ...option.RequestOption) (res *Page, err error) {
	opts = slices.Concat(r.Options, opts)
	if body.ObjectID == "" {
		err = errors.New("missing required objectId parameter")
		return
	}
	if revisionID == "" {
		err = errors.New("missing required revisionId parameter")
		return
	}
	path := fmt.Sprintf("cms/v3/pages/landing-pages/%s/revisions/%s/restore", body.ObjectID, revisionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Takes a specified version of a Landing Page, sets it as the new draft version of
// the Landing Page.
func (r *PageLandingPageService) RestoreRevisionToDraft(ctx context.Context, revisionID int64, body PageLandingPageRestoreRevisionToDraftParams, opts ...option.RequestOption) (res *Page, err error) {
	opts = slices.Concat(r.Options, opts)
	if body.ObjectID == "" {
		err = errors.New("missing required objectId parameter")
		return
	}
	path := fmt.Sprintf("cms/v3/pages/landing-pages/%s/revisions/%v/restore-to-draft", body.ObjectID, revisionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Schedule a Landing Page to be Published
func (r *PageLandingPageService) Schedule(ctx context.Context, body PageLandingPageScheduleParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "cms/v3/pages/landing-pages/schedule"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Set a landing page as the primary language of a multi-language group.
func (r *PageLandingPageService) SetNewLangPrimary(ctx context.Context, body PageLandingPageSetNewLangPrimaryParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "cms/v3/pages/landing-pages/multi-language/set-new-lang-primary"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, nil, opts...)
	return
}

// Update the Landing Page objects identified in the request body.
func (r *PageLandingPageService) UpdateBatch(ctx context.Context, params PageLandingPageUpdateBatchParams, opts ...option.RequestOption) (res *BatchResponsePage, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "cms/v3/pages/landing-pages/batch/update"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Sparse updates the draft version of a single Landing Page object identified by
// the id in the path. You only need to specify the column values that you are
// modifying.
func (r *PageLandingPageService) UpdateDraft(ctx context.Context, objectID string, body PageLandingPageUpdateDraftParams, opts ...option.RequestOption) (res *Page, err error) {
	opts = slices.Concat(r.Options, opts)
	if objectID == "" {
		err = errors.New("missing required objectId parameter")
		return
	}
	path := fmt.Sprintf("cms/v3/pages/landing-pages/%s/draft", objectID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Sparse updates a single Folder object identified by the id in the path. You only
// need to specify the column values that you are modifying.
func (r *PageLandingPageService) UpdateFolder(ctx context.Context, objectID string, params PageLandingPageUpdateFolderParams, opts ...option.RequestOption) (res *ContentFolder, err error) {
	opts = slices.Concat(r.Options, opts)
	if objectID == "" {
		err = errors.New("missing required objectId parameter")
		return
	}
	path := fmt.Sprintf("cms/v3/pages/landing-pages/folders/%s", objectID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return
}

// Update the Folder objects identified in the request body.
func (r *PageLandingPageService) UpdateFoldersBatch(ctx context.Context, params PageLandingPageUpdateFoldersBatchParams, opts ...option.RequestOption) (res *BatchResponseContentFolder, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "cms/v3/pages/landing-pages/folders/batch/update"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Explicitly set new languages for each landing page in a multi-language group.
func (r *PageLandingPageService) UpdateLanguages(ctx context.Context, body PageLandingPageUpdateLanguagesParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "cms/v3/pages/landing-pages/multi-language/update-languages"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

type PageLandingPageNewParams struct {
	// Model definition for a landing page or site page.
	Page PageParam
	paramObj
}

func (r PageLandingPageNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.Page)
}
func (r *PageLandingPageNewParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.Page)
}

type PageLandingPageUpdateParams struct {
	// Model definition for a landing page or site page.
	Page PageParam
	// Specifies whether to update deleted Landing Pages. Defaults to `false`.
	Archived param.Opt[bool] `query:"archived,omitzero" json:"-"`
	paramObj
}

func (r PageLandingPageUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.Page)
}
func (r *PageLandingPageUpdateParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.Page)
}

// URLQuery serializes [PageLandingPageUpdateParams]'s query parameters as
// `url.Values`.
func (r PageLandingPageUpdateParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type PageLandingPageListParams struct {
	// The cursor token value to get the next set of results. You can get this from the
	// `paging.next.after` JSON property of a paged response containing more results.
	After param.Opt[string] `query:"after,omitzero" json:"-"`
	// Specifies whether to return deleted Landing Pages. Defaults to `false`.
	Archived param.Opt[bool] `query:"archived,omitzero" json:"-"`
	// Only return Landing Pages created after the specified time.
	CreatedAfter param.Opt[time.Time] `query:"createdAfter,omitzero" format:"date-time" json:"-"`
	// Only return Landing Pages created at exactly the specified time.
	CreatedAt param.Opt[time.Time] `query:"createdAt,omitzero" format:"date-time" json:"-"`
	// Only return Landing Pages created before the specified time.
	CreatedBefore param.Opt[time.Time] `query:"createdBefore,omitzero" format:"date-time" json:"-"`
	// The maximum number of results to return. Default is 100.
	Limit    param.Opt[int64]  `query:"limit,omitzero" json:"-"`
	Property param.Opt[string] `query:"property,omitzero" json:"-"`
	// Only return Landing Pages last updated after the specified time.
	UpdatedAfter param.Opt[time.Time] `query:"updatedAfter,omitzero" format:"date-time" json:"-"`
	// Only return Landing Pages last updated at exactly the specified time.
	UpdatedAt param.Opt[time.Time] `query:"updatedAt,omitzero" format:"date-time" json:"-"`
	// Only return Landing Pages last updated before the specified time.
	UpdatedBefore param.Opt[time.Time] `query:"updatedBefore,omitzero" format:"date-time" json:"-"`
	// Specifies which fields to use for sorting results. Valid fields are `name`,
	// `createdAt`, `updatedAt`, `createdBy`, `updatedBy`. `createdAt` will be used by
	// default.
	Sort []string `query:"sort,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [PageLandingPageListParams]'s query parameters as
// `url.Values`.
func (r PageLandingPageListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type PageLandingPageDeleteParams struct {
	// Whether to return only results that have been archived.
	Archived param.Opt[bool] `query:"archived,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [PageLandingPageDeleteParams]'s query parameters as
// `url.Values`.
func (r PageLandingPageDeleteParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type PageLandingPageAttachToLangGroupParams struct {
	// Request body object for attaching objects to multi-language groups.
	AttachToLangPrimaryRequestVNext AttachToLangPrimaryRequestVNextParam
	paramObj
}

func (r PageLandingPageAttachToLangGroupParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.AttachToLangPrimaryRequestVNext)
}
func (r *PageLandingPageAttachToLangGroupParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.AttachToLangPrimaryRequestVNext)
}

type PageLandingPageCloneParams struct {
	// Request body object for cloning content.
	ContentCloneRequestVNext ContentCloneRequestVNextParam
	paramObj
}

func (r PageLandingPageCloneParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.ContentCloneRequestVNext)
}
func (r *PageLandingPageCloneParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.ContentCloneRequestVNext)
}

type PageLandingPageNewAbTestVariationParams struct {
	// Request body object for creating A/B tests.
	AbTestCreateRequestVNext shared.AbTestCreateRequestVNextParam
	paramObj
}

func (r PageLandingPageNewAbTestVariationParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.AbTestCreateRequestVNext)
}
func (r *PageLandingPageNewAbTestVariationParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.AbTestCreateRequestVNext)
}

type PageLandingPageNewBatchParams struct {
	// Wrapper for providing an array of pages as inputs.
	BatchInputPage BatchInputPageParam
	paramObj
}

func (r PageLandingPageNewBatchParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.BatchInputPage)
}
func (r *PageLandingPageNewBatchParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.BatchInputPage)
}

type PageLandingPageNewFolderParams struct {
	// Model definition for a content folder.
	ContentFolder ContentFolderParam
	paramObj
}

func (r PageLandingPageNewFolderParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.ContentFolder)
}
func (r *PageLandingPageNewFolderParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.ContentFolder)
}

type PageLandingPageNewFoldersBatchParams struct {
	// Wrapper for providing an array of content folders as inputs.
	BatchInputContentFolder BatchInputContentFolderParam
	paramObj
}

func (r PageLandingPageNewFoldersBatchParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.BatchInputContentFolder)
}
func (r *PageLandingPageNewFoldersBatchParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.BatchInputContentFolder)
}

type PageLandingPageNewLanguageVariationParams struct {
	// Request body object for creating new language variant content.
	ContentLanguageCloneRequestVNext ContentLanguageCloneRequestVNextParam
	paramObj
}

func (r PageLandingPageNewLanguageVariationParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.ContentLanguageCloneRequestVNext)
}
func (r *PageLandingPageNewLanguageVariationParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.ContentLanguageCloneRequestVNext)
}

type PageLandingPageDeleteBatchParams struct {
	// Wrapper for providing an array of strings as inputs.
	BatchInputString shared.BatchInputStringParam
	paramObj
}

func (r PageLandingPageDeleteBatchParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.BatchInputString)
}
func (r *PageLandingPageDeleteBatchParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.BatchInputString)
}

type PageLandingPageDeleteFolderParams struct {
	// Whether to return only results that have been archived.
	Archived param.Opt[bool] `query:"archived,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [PageLandingPageDeleteFolderParams]'s query parameters as
// `url.Values`.
func (r PageLandingPageDeleteFolderParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type PageLandingPageDeleteFoldersBatchParams struct {
	// Wrapper for providing an array of strings as inputs.
	BatchInputString shared.BatchInputStringParam
	paramObj
}

func (r PageLandingPageDeleteFoldersBatchParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.BatchInputString)
}
func (r *PageLandingPageDeleteFoldersBatchParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.BatchInputString)
}

type PageLandingPageDetachFromLangGroupParams struct {
	// Request body object for detaching objects from multi-language groups.
	DetachFromLangGroupRequestVNext DetachFromLangGroupRequestVNextParam
	paramObj
}

func (r PageLandingPageDetachFromLangGroupParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.DetachFromLangGroupRequestVNext)
}
func (r *PageLandingPageDetachFromLangGroupParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.DetachFromLangGroupRequestVNext)
}

type PageLandingPageEndAbTestParams struct {
	// Request body object for ending A/B tests.
	AbTestEndRequestVNext AbTestEndRequestVNextParam
	paramObj
}

func (r PageLandingPageEndAbTestParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.AbTestEndRequestVNext)
}
func (r *PageLandingPageEndAbTestParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.AbTestEndRequestVNext)
}

type PageLandingPageGetParams struct {
	// Specifies whether to return deleted Landing Pages. Defaults to `false`.
	Archived param.Opt[bool]   `query:"archived,omitzero" json:"-"`
	Property param.Opt[string] `query:"property,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [PageLandingPageGetParams]'s query parameters as
// `url.Values`.
func (r PageLandingPageGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type PageLandingPageGetBatchParams struct {
	// Wrapper for providing an array of strings as inputs.
	BatchInputString shared.BatchInputStringParam
	// Specifies whether to return deleted Landing Pages. Defaults to `false`.
	Archived param.Opt[bool] `query:"archived,omitzero" json:"-"`
	paramObj
}

func (r PageLandingPageGetBatchParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.BatchInputString)
}
func (r *PageLandingPageGetBatchParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.BatchInputString)
}

// URLQuery serializes [PageLandingPageGetBatchParams]'s query parameters as
// `url.Values`.
func (r PageLandingPageGetBatchParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type PageLandingPageGetFolderParams struct {
	// Specifies whether to return deleted Folders. Defaults to `false`.
	Archived param.Opt[bool]   `query:"archived,omitzero" json:"-"`
	Property param.Opt[string] `query:"property,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [PageLandingPageGetFolderParams]'s query parameters as
// `url.Values`.
func (r PageLandingPageGetFolderParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type PageLandingPageGetFolderRevisionParams struct {
	ObjectID string `path:"objectId,required" json:"-"`
	paramObj
}

type PageLandingPageGetFoldersBatchParams struct {
	// Wrapper for providing an array of strings as inputs.
	BatchInputString shared.BatchInputStringParam
	// Specifies whether to return deleted Folders. Defaults to `false`.
	Archived param.Opt[bool] `query:"archived,omitzero" json:"-"`
	paramObj
}

func (r PageLandingPageGetFoldersBatchParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.BatchInputString)
}
func (r *PageLandingPageGetFoldersBatchParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.BatchInputString)
}

// URLQuery serializes [PageLandingPageGetFoldersBatchParams]'s query parameters as
// `url.Values`.
func (r PageLandingPageGetFoldersBatchParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type PageLandingPageGetRevisionParams struct {
	ObjectID string `path:"objectId,required" json:"-"`
	paramObj
}

type PageLandingPageListFolderRevisionsParams struct {
	// The cursor token value to get the next set of results. You can get this from the
	// `paging.next.after` JSON property of a paged response containing more results.
	After  param.Opt[string] `query:"after,omitzero" json:"-"`
	Before param.Opt[string] `query:"before,omitzero" json:"-"`
	// The maximum number of results to return. Default is 100.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [PageLandingPageListFolderRevisionsParams]'s query
// parameters as `url.Values`.
func (r PageLandingPageListFolderRevisionsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type PageLandingPageListFoldersParams struct {
	// The cursor token value to get the next set of results. You can get this from the
	// `paging.next.after` JSON property of a paged response containing more results.
	After param.Opt[string] `query:"after,omitzero" json:"-"`
	// Specifies whether to return deleted Folders. Defaults to `false`.
	Archived param.Opt[bool] `query:"archived,omitzero" json:"-"`
	// Only return Folders created after the specified time.
	CreatedAfter param.Opt[time.Time] `query:"createdAfter,omitzero" format:"date-time" json:"-"`
	// Only return Folders created at exactly the specified time.
	CreatedAt param.Opt[time.Time] `query:"createdAt,omitzero" format:"date-time" json:"-"`
	// Only return Folders created before the specified time.
	CreatedBefore param.Opt[time.Time] `query:"createdBefore,omitzero" format:"date-time" json:"-"`
	// The maximum number of results to return. Default is 100.
	Limit    param.Opt[int64]  `query:"limit,omitzero" json:"-"`
	Property param.Opt[string] `query:"property,omitzero" json:"-"`
	// Only return Folders last updated after the specified time.
	UpdatedAfter param.Opt[time.Time] `query:"updatedAfter,omitzero" format:"date-time" json:"-"`
	// Only return Folders last updated at exactly the specified time.
	UpdatedAt param.Opt[time.Time] `query:"updatedAt,omitzero" format:"date-time" json:"-"`
	// Only return Folders last updated before the specified time.
	UpdatedBefore param.Opt[time.Time] `query:"updatedBefore,omitzero" format:"date-time" json:"-"`
	// Specifies which fields to use for sorting results. Valid fields are `name`,
	// `createdAt`, `updatedAt`, `createdBy`, `updatedBy`. `createdAt` will be used by
	// default.
	Sort []string `query:"sort,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [PageLandingPageListFoldersParams]'s query parameters as
// `url.Values`.
func (r PageLandingPageListFoldersParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type PageLandingPageListRevisionsParams struct {
	// The cursor token value to get the next set of results. You can get this from the
	// `paging.next.after` JSON property of a paged response containing more results.
	After  param.Opt[string] `query:"after,omitzero" json:"-"`
	Before param.Opt[string] `query:"before,omitzero" json:"-"`
	// The maximum number of results to return. Default is 100.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [PageLandingPageListRevisionsParams]'s query parameters as
// `url.Values`.
func (r PageLandingPageListRevisionsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type PageLandingPageRerunAbTestParams struct {
	// Request body object for rerunning A/B tests.
	AbTestRerunRequestVNext AbTestRerunRequestVNextParam
	paramObj
}

func (r PageLandingPageRerunAbTestParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.AbTestRerunRequestVNext)
}
func (r *PageLandingPageRerunAbTestParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.AbTestRerunRequestVNext)
}

type PageLandingPageRestoreFolderRevisionParams struct {
	ObjectID string `path:"objectId,required" json:"-"`
	paramObj
}

type PageLandingPageRestoreRevisionParams struct {
	ObjectID string `path:"objectId,required" json:"-"`
	paramObj
}

type PageLandingPageRestoreRevisionToDraftParams struct {
	ObjectID string `path:"objectId,required" json:"-"`
	paramObj
}

type PageLandingPageScheduleParams struct {
	// Request body object for scheduling the publish of content
	ContentScheduleRequestVNext ContentScheduleRequestVNextParam
	paramObj
}

func (r PageLandingPageScheduleParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.ContentScheduleRequestVNext)
}
func (r *PageLandingPageScheduleParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.ContentScheduleRequestVNext)
}

type PageLandingPageSetNewLangPrimaryParams struct {
	// Request body object for setting a new primary language.
	SetNewLanguagePrimaryRequestVNext SetNewLanguagePrimaryRequestVNextParam
	paramObj
}

func (r PageLandingPageSetNewLangPrimaryParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.SetNewLanguagePrimaryRequestVNext)
}
func (r *PageLandingPageSetNewLangPrimaryParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.SetNewLanguagePrimaryRequestVNext)
}

type PageLandingPageUpdateBatchParams struct {
	// Wrapper for providing an array of JSON nodes as inputs.
	BatchInputJsonNode BatchInputJsonNodeParam
	// Specifies whether to update deleted Landing Pages. Defaults to `false`.
	Archived param.Opt[bool] `query:"archived,omitzero" json:"-"`
	paramObj
}

func (r PageLandingPageUpdateBatchParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.BatchInputJsonNode)
}
func (r *PageLandingPageUpdateBatchParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.BatchInputJsonNode)
}

// URLQuery serializes [PageLandingPageUpdateBatchParams]'s query parameters as
// `url.Values`.
func (r PageLandingPageUpdateBatchParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type PageLandingPageUpdateDraftParams struct {
	// Model definition for a landing page or site page.
	Page PageParam
	paramObj
}

func (r PageLandingPageUpdateDraftParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.Page)
}
func (r *PageLandingPageUpdateDraftParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.Page)
}

type PageLandingPageUpdateFolderParams struct {
	// Model definition for a content folder.
	ContentFolder ContentFolderParam
	// Specifies whether to update deleted Folders. Defaults to `false`.
	Archived param.Opt[bool] `query:"archived,omitzero" json:"-"`
	paramObj
}

func (r PageLandingPageUpdateFolderParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.ContentFolder)
}
func (r *PageLandingPageUpdateFolderParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.ContentFolder)
}

// URLQuery serializes [PageLandingPageUpdateFolderParams]'s query parameters as
// `url.Values`.
func (r PageLandingPageUpdateFolderParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type PageLandingPageUpdateFoldersBatchParams struct {
	// Wrapper for providing an array of JSON nodes as inputs.
	BatchInputJsonNode BatchInputJsonNodeParam
	// Whether to return only results that have been archived.
	Archived param.Opt[bool] `query:"archived,omitzero" json:"-"`
	paramObj
}

func (r PageLandingPageUpdateFoldersBatchParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.BatchInputJsonNode)
}
func (r *PageLandingPageUpdateFoldersBatchParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.BatchInputJsonNode)
}

// URLQuery serializes [PageLandingPageUpdateFoldersBatchParams]'s query parameters
// as `url.Values`.
func (r PageLandingPageUpdateFoldersBatchParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type PageLandingPageUpdateLanguagesParams struct {
	// Request object for updating languages within a multi-language group.
	UpdateLanguagesRequestVNext UpdateLanguagesRequestVNextParam
	paramObj
}

func (r PageLandingPageUpdateLanguagesParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.UpdateLanguagesRequestVNext)
}
func (r *PageLandingPageUpdateLanguagesParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.UpdateLanguagesRequestVNext)
}
