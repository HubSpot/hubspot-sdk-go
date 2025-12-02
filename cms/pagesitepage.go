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

// PageSitePageService contains methods and other services that help with
// interacting with the hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPageSitePageService] method instead.
type PageSitePageService struct {
	Options []option.RequestOption
}

// NewPageSitePageService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewPageSitePageService(opts ...option.RequestOption) (r PageSitePageService) {
	r = PageSitePageService{}
	r.Options = opts
	return
}

// Create a new Site Page
func (r *PageSitePageService) New(ctx context.Context, body PageSitePageNewParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "cms/v3/pages/site-pages"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Sparse updates a single Site Page object identified by the id in the path. You
// only need to specify the column values that you are modifying.
func (r *PageSitePageService) Update(ctx context.Context, objectID string, params PageSitePageUpdateParams, opts ...option.RequestOption) (res *Page, err error) {
	opts = slices.Concat(r.Options, opts)
	if objectID == "" {
		err = errors.New("missing required objectId parameter")
		return
	}
	path := fmt.Sprintf("cms/v3/pages/site-pages/%s", objectID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return
}

// Get the list of site pages. Supports paging and filtering. This method would be
// useful for an integration that examined these models and used an external
// service to suggest edits.
func (r *PageSitePageService) List(ctx context.Context, query PageSitePageListParams, opts ...option.RequestOption) (res *pagination.Page[Page], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "cms/v3/pages/site-pages"
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

// Get the list of site pages. Supports paging and filtering. This method would be
// useful for an integration that examined these models and used an external
// service to suggest edits.
func (r *PageSitePageService) ListAutoPaging(ctx context.Context, query PageSitePageListParams, opts ...option.RequestOption) *pagination.PageAutoPager[Page] {
	return pagination.NewPageAutoPager(r.List(ctx, query, opts...))
}

// Delete the Site Page object identified by the id in the path.
func (r *PageSitePageService) Delete(ctx context.Context, objectID string, body PageSitePageDeleteParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if objectID == "" {
		err = errors.New("missing required objectId parameter")
		return
	}
	path := fmt.Sprintf("cms/v3/pages/site-pages/%s", objectID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, nil, opts...)
	return
}

// Attach a site page to a multi-language group.
func (r *PageSitePageService) AttachToLangGroup(ctx context.Context, body PageSitePageAttachToLangGroupParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "cms/v3/pages/site-pages/multi-language/attach-to-lang-group"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Clone a Site Page
func (r *PageSitePageService) Clone(ctx context.Context, body PageSitePageCloneParams, opts ...option.RequestOption) (res *Page, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "cms/v3/pages/site-pages/clone"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Create a new A/B test variation based on the information provided in the request
// body.
func (r *PageSitePageService) NewAbTestVariation(ctx context.Context, body PageSitePageNewAbTestVariationParams, opts ...option.RequestOption) (res *Page, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "cms/v3/pages/site-pages/ab-test/create-variation"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Create the Site Page objects detailed in the request body.
func (r *PageSitePageService) NewBatch(ctx context.Context, body PageSitePageNewBatchParams, opts ...option.RequestOption) (res *BatchResponsePage, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "cms/v3/pages/site-pages/batch/create"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Create a new language variation from an existing site page
func (r *PageSitePageService) NewLanguageVariation(ctx context.Context, body PageSitePageNewLanguageVariationParams, opts ...option.RequestOption) (res *Page, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "cms/v3/pages/site-pages/multi-language/create-language-variation"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Delete the Site Page objects identified in the request body. Note: This is not
// the same as the dashboard `archive` function. To perform a dashboard `archive`
// send an normal update with the `archivedInDashboard` field set to true.
func (r *PageSitePageService) DeleteBatch(ctx context.Context, body PageSitePageDeleteBatchParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "cms/v3/pages/site-pages/batch/archive"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Detach a site page from a multi-language group.
func (r *PageSitePageService) DetachFromLangGroup(ctx context.Context, body PageSitePageDetachFromLangGroupParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "cms/v3/pages/site-pages/multi-language/detach-from-lang-group"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// End an active A/B test and designate a winner.
func (r *PageSitePageService) EndAbTest(ctx context.Context, body PageSitePageEndAbTestParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "cms/v3/pages/site-pages/ab-test/end"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Retrieve the Site Page object identified by the id in the path.
func (r *PageSitePageService) Get(ctx context.Context, objectID string, query PageSitePageGetParams, opts ...option.RequestOption) (res *Page, err error) {
	opts = slices.Concat(r.Options, opts)
	if objectID == "" {
		err = errors.New("missing required objectId parameter")
		return
	}
	path := fmt.Sprintf("cms/v3/pages/site-pages/%s", objectID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Retrieve the Site Page objects identified in the request body.
func (r *PageSitePageService) GetBatch(ctx context.Context, params PageSitePageGetBatchParams, opts ...option.RequestOption) (res *BatchResponsePage, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "cms/v3/pages/site-pages/batch/read"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Retrieve the full draft version of the Site Page.
func (r *PageSitePageService) GetDraft(ctx context.Context, objectID string, opts ...option.RequestOption) (res *Page, err error) {
	opts = slices.Concat(r.Options, opts)
	if objectID == "" {
		err = errors.New("missing required objectId parameter")
		return
	}
	path := fmt.Sprintf("cms/v3/pages/site-pages/%s/draft", objectID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Retrieves a previous version of a Site Page
func (r *PageSitePageService) GetRevision(ctx context.Context, revisionID string, query PageSitePageGetRevisionParams, opts ...option.RequestOption) (res *VersionPage, err error) {
	opts = slices.Concat(r.Options, opts)
	if query.ObjectID == "" {
		err = errors.New("missing required objectId parameter")
		return
	}
	if revisionID == "" {
		err = errors.New("missing required revisionId parameter")
		return
	}
	path := fmt.Sprintf("cms/v3/pages/site-pages/%s/revisions/%s", query.ObjectID, revisionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Retrieves all the previous versions of a Site Page.
func (r *PageSitePageService) ListRevisions(ctx context.Context, objectID string, query PageSitePageListRevisionsParams, opts ...option.RequestOption) (res *pagination.Page[VersionPage], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if objectID == "" {
		err = errors.New("missing required objectId parameter")
		return
	}
	path := fmt.Sprintf("cms/v3/pages/site-pages/%s/revisions", objectID)
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

// Retrieves all the previous versions of a Site Page.
func (r *PageSitePageService) ListRevisionsAutoPaging(ctx context.Context, objectID string, query PageSitePageListRevisionsParams, opts ...option.RequestOption) *pagination.PageAutoPager[VersionPage] {
	return pagination.NewPageAutoPager(r.ListRevisions(ctx, objectID, query, opts...))
}

// Take any changes from the draft version of the Site Page and apply them to the
// live version.
func (r *PageSitePageService) PublishDraft(ctx context.Context, objectID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if objectID == "" {
		err = errors.New("missing required objectId parameter")
		return
	}
	path := fmt.Sprintf("cms/v3/pages/site-pages/%s/draft/push-live", objectID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, nil, opts...)
	return
}

// Rerun a previous A/B test.
func (r *PageSitePageService) RerunAbTest(ctx context.Context, body PageSitePageRerunAbTestParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "cms/v3/pages/site-pages/ab-test/rerun"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Discards any edits and resets the draft to the live version.
func (r *PageSitePageService) ResetDraft(ctx context.Context, objectID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if objectID == "" {
		err = errors.New("missing required objectId parameter")
		return
	}
	path := fmt.Sprintf("cms/v3/pages/site-pages/%s/draft/reset", objectID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, nil, opts...)
	return
}

// Takes a specified version of a Site Page and restores it.
func (r *PageSitePageService) RestoreRevision(ctx context.Context, revisionID string, body PageSitePageRestoreRevisionParams, opts ...option.RequestOption) (res *Page, err error) {
	opts = slices.Concat(r.Options, opts)
	if body.ObjectID == "" {
		err = errors.New("missing required objectId parameter")
		return
	}
	if revisionID == "" {
		err = errors.New("missing required revisionId parameter")
		return
	}
	path := fmt.Sprintf("cms/v3/pages/site-pages/%s/revisions/%s/restore", body.ObjectID, revisionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Takes a specified version of a Site Page, sets it as the new draft version of
// the Site Page.
func (r *PageSitePageService) RestoreRevisionToDraft(ctx context.Context, revisionID int64, body PageSitePageRestoreRevisionToDraftParams, opts ...option.RequestOption) (res *Page, err error) {
	opts = slices.Concat(r.Options, opts)
	if body.ObjectID == "" {
		err = errors.New("missing required objectId parameter")
		return
	}
	path := fmt.Sprintf("cms/v3/pages/site-pages/%s/revisions/%v/restore-to-draft", body.ObjectID, revisionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Schedule a Site Page to be Published
func (r *PageSitePageService) Schedule(ctx context.Context, body PageSitePageScheduleParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "cms/v3/pages/site-pages/schedule"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Set a site page as the primary language of a multi-language group.
func (r *PageSitePageService) SetNewLangPrimary(ctx context.Context, body PageSitePageSetNewLangPrimaryParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "cms/v3/pages/site-pages/multi-language/set-new-lang-primary"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, nil, opts...)
	return
}

// Update the Site Page objects identified in the request body.
func (r *PageSitePageService) UpdateBatch(ctx context.Context, params PageSitePageUpdateBatchParams, opts ...option.RequestOption) (res *BatchResponsePage, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "cms/v3/pages/site-pages/batch/update"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Sparse updates the draft version of a single Site Page object identified by the
// id in the path. You only need to specify the column values that you are
// modifying.
func (r *PageSitePageService) UpdateDraft(ctx context.Context, objectID string, body PageSitePageUpdateDraftParams, opts ...option.RequestOption) (res *Page, err error) {
	opts = slices.Concat(r.Options, opts)
	if objectID == "" {
		err = errors.New("missing required objectId parameter")
		return
	}
	path := fmt.Sprintf("cms/v3/pages/site-pages/%s/draft", objectID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Explicitly set new languages for each site page in a multi-language group.
func (r *PageSitePageService) UpdateLanguages(ctx context.Context, body PageSitePageUpdateLanguagesParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "cms/v3/pages/site-pages/multi-language/update-languages"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

type PageSitePageNewParams struct {
	// Model definition for a landing page or site page.
	Page PageParam
	paramObj
}

func (r PageSitePageNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.Page)
}
func (r *PageSitePageNewParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.Page)
}

type PageSitePageUpdateParams struct {
	// Model definition for a landing page or site page.
	Page PageParam
	// Specifies whether to update deleted Site Pages. Defaults to `false`.
	Archived param.Opt[bool] `query:"archived,omitzero" json:"-"`
	paramObj
}

func (r PageSitePageUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.Page)
}
func (r *PageSitePageUpdateParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.Page)
}

// URLQuery serializes [PageSitePageUpdateParams]'s query parameters as
// `url.Values`.
func (r PageSitePageUpdateParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type PageSitePageListParams struct {
	// The cursor token value to get the next set of results. You can get this from the
	// `paging.next.after` JSON property of a paged response containing more results.
	After param.Opt[string] `query:"after,omitzero" json:"-"`
	// Specifies whether to return deleted Site Pages. Defaults to `false`.
	Archived param.Opt[bool] `query:"archived,omitzero" json:"-"`
	// Only return Site Pages created after the specified time.
	CreatedAfter param.Opt[time.Time] `query:"createdAfter,omitzero" format:"date-time" json:"-"`
	// Only return Site Pages created at exactly the specified time.
	CreatedAt param.Opt[time.Time] `query:"createdAt,omitzero" format:"date-time" json:"-"`
	// Only return Site Pages created before the specified time.
	CreatedBefore param.Opt[time.Time] `query:"createdBefore,omitzero" format:"date-time" json:"-"`
	// The maximum number of results to return. Default is 100.
	Limit    param.Opt[int64]  `query:"limit,omitzero" json:"-"`
	Property param.Opt[string] `query:"property,omitzero" json:"-"`
	// Only return Site Pages last updated after the specified time.
	UpdatedAfter param.Opt[time.Time] `query:"updatedAfter,omitzero" format:"date-time" json:"-"`
	// Only return Site Pages last updated at exactly the specified time.
	UpdatedAt param.Opt[time.Time] `query:"updatedAt,omitzero" format:"date-time" json:"-"`
	// Only return Site Pages last updated before the specified time.
	UpdatedBefore param.Opt[time.Time] `query:"updatedBefore,omitzero" format:"date-time" json:"-"`
	// Specifies which fields to use for sorting results. Valid fields are `name`,
	// `createdAt`, `updatedAt`, `createdBy`, `updatedBy`. `createdAt` will be used by
	// default.
	Sort []string `query:"sort,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [PageSitePageListParams]'s query parameters as `url.Values`.
func (r PageSitePageListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type PageSitePageDeleteParams struct {
	// Whether to return only results that have been archived.
	Archived param.Opt[bool] `query:"archived,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [PageSitePageDeleteParams]'s query parameters as
// `url.Values`.
func (r PageSitePageDeleteParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type PageSitePageAttachToLangGroupParams struct {
	// Request body object for attaching objects to multi-language groups.
	AttachToLangPrimaryRequestVNext AttachToLangPrimaryRequestVNextParam
	paramObj
}

func (r PageSitePageAttachToLangGroupParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.AttachToLangPrimaryRequestVNext)
}
func (r *PageSitePageAttachToLangGroupParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.AttachToLangPrimaryRequestVNext)
}

type PageSitePageCloneParams struct {
	// Request body object for cloning content.
	ContentCloneRequestVNext ContentCloneRequestVNextParam
	paramObj
}

func (r PageSitePageCloneParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.ContentCloneRequestVNext)
}
func (r *PageSitePageCloneParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.ContentCloneRequestVNext)
}

type PageSitePageNewAbTestVariationParams struct {
	// Request body object for creating A/B tests.
	AbTestCreateRequestVNext shared.AbTestCreateRequestVNextParam
	paramObj
}

func (r PageSitePageNewAbTestVariationParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.AbTestCreateRequestVNext)
}
func (r *PageSitePageNewAbTestVariationParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.AbTestCreateRequestVNext)
}

type PageSitePageNewBatchParams struct {
	// Wrapper for providing an array of pages as inputs.
	BatchInputPage BatchInputPageParam
	paramObj
}

func (r PageSitePageNewBatchParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.BatchInputPage)
}
func (r *PageSitePageNewBatchParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.BatchInputPage)
}

type PageSitePageNewLanguageVariationParams struct {
	// Request body object for creating new language variant content.
	ContentLanguageCloneRequestVNext ContentLanguageCloneRequestVNextParam
	paramObj
}

func (r PageSitePageNewLanguageVariationParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.ContentLanguageCloneRequestVNext)
}
func (r *PageSitePageNewLanguageVariationParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.ContentLanguageCloneRequestVNext)
}

type PageSitePageDeleteBatchParams struct {
	// Wrapper for providing an array of strings as inputs.
	BatchInputString shared.BatchInputStringParam
	paramObj
}

func (r PageSitePageDeleteBatchParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.BatchInputString)
}
func (r *PageSitePageDeleteBatchParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.BatchInputString)
}

type PageSitePageDetachFromLangGroupParams struct {
	// Request body object for detaching objects from multi-language groups.
	DetachFromLangGroupRequestVNext DetachFromLangGroupRequestVNextParam
	paramObj
}

func (r PageSitePageDetachFromLangGroupParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.DetachFromLangGroupRequestVNext)
}
func (r *PageSitePageDetachFromLangGroupParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.DetachFromLangGroupRequestVNext)
}

type PageSitePageEndAbTestParams struct {
	// Request body object for ending A/B tests.
	AbTestEndRequestVNext AbTestEndRequestVNextParam
	paramObj
}

func (r PageSitePageEndAbTestParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.AbTestEndRequestVNext)
}
func (r *PageSitePageEndAbTestParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.AbTestEndRequestVNext)
}

type PageSitePageGetParams struct {
	// Specifies whether to return deleted Site Pages. Defaults to `false`.
	Archived param.Opt[bool]   `query:"archived,omitzero" json:"-"`
	Property param.Opt[string] `query:"property,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [PageSitePageGetParams]'s query parameters as `url.Values`.
func (r PageSitePageGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type PageSitePageGetBatchParams struct {
	// Wrapper for providing an array of strings as inputs.
	BatchInputString shared.BatchInputStringParam
	// Specifies whether to return deleted Site Pages. Defaults to `false`.
	Archived param.Opt[bool] `query:"archived,omitzero" json:"-"`
	paramObj
}

func (r PageSitePageGetBatchParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.BatchInputString)
}
func (r *PageSitePageGetBatchParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.BatchInputString)
}

// URLQuery serializes [PageSitePageGetBatchParams]'s query parameters as
// `url.Values`.
func (r PageSitePageGetBatchParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type PageSitePageGetRevisionParams struct {
	ObjectID string `path:"objectId,required" json:"-"`
	paramObj
}

type PageSitePageListRevisionsParams struct {
	// The cursor token value to get the next set of results. You can get this from the
	// `paging.next.after` JSON property of a paged response containing more results.
	After  param.Opt[string] `query:"after,omitzero" json:"-"`
	Before param.Opt[string] `query:"before,omitzero" json:"-"`
	// The maximum number of results to return. Default is 100.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [PageSitePageListRevisionsParams]'s query parameters as
// `url.Values`.
func (r PageSitePageListRevisionsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type PageSitePageRerunAbTestParams struct {
	// Request body object for rerunning A/B tests.
	AbTestRerunRequestVNext AbTestRerunRequestVNextParam
	paramObj
}

func (r PageSitePageRerunAbTestParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.AbTestRerunRequestVNext)
}
func (r *PageSitePageRerunAbTestParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.AbTestRerunRequestVNext)
}

type PageSitePageRestoreRevisionParams struct {
	ObjectID string `path:"objectId,required" json:"-"`
	paramObj
}

type PageSitePageRestoreRevisionToDraftParams struct {
	ObjectID string `path:"objectId,required" json:"-"`
	paramObj
}

type PageSitePageScheduleParams struct {
	// Request body object for scheduling the publish of content
	ContentScheduleRequestVNext ContentScheduleRequestVNextParam
	paramObj
}

func (r PageSitePageScheduleParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.ContentScheduleRequestVNext)
}
func (r *PageSitePageScheduleParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.ContentScheduleRequestVNext)
}

type PageSitePageSetNewLangPrimaryParams struct {
	// Request body object for setting a new primary language.
	SetNewLanguagePrimaryRequestVNext SetNewLanguagePrimaryRequestVNextParam
	paramObj
}

func (r PageSitePageSetNewLangPrimaryParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.SetNewLanguagePrimaryRequestVNext)
}
func (r *PageSitePageSetNewLangPrimaryParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.SetNewLanguagePrimaryRequestVNext)
}

type PageSitePageUpdateBatchParams struct {
	// Wrapper for providing an array of JSON nodes as inputs.
	BatchInputJsonNode BatchInputJsonNodeParam
	// Specifies whether to update deleted Site Pages. Defaults to `false`.
	Archived param.Opt[bool] `query:"archived,omitzero" json:"-"`
	paramObj
}

func (r PageSitePageUpdateBatchParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.BatchInputJsonNode)
}
func (r *PageSitePageUpdateBatchParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.BatchInputJsonNode)
}

// URLQuery serializes [PageSitePageUpdateBatchParams]'s query parameters as
// `url.Values`.
func (r PageSitePageUpdateBatchParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type PageSitePageUpdateDraftParams struct {
	// Model definition for a landing page or site page.
	Page PageParam
	paramObj
}

func (r PageSitePageUpdateDraftParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.Page)
}
func (r *PageSitePageUpdateDraftParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.Page)
}

type PageSitePageUpdateLanguagesParams struct {
	// Request object for updating languages within a multi-language group.
	UpdateLanguagesRequestVNext UpdateLanguagesRequestVNextParam
	paramObj
}

func (r PageSitePageUpdateLanguagesParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.UpdateLanguagesRequestVNext)
}
func (r *PageSitePageUpdateLanguagesParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.UpdateLanguagesRequestVNext)
}
