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
	"github.com/HubSpot/hubspot-sdk-go/packages/pagination"
	"github.com/HubSpot/hubspot-sdk-go/packages/param"
)

// PageSitePageService contains methods and other services that help with
// interacting with the hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPageSitePageService] method instead.
type PageSitePageService struct {
	options       []option.RequestOption
	AbTest        PageSitePageAbTestService
	Batch         PageSitePageBatchService
	Draft         PageSitePageDraftService
	MultiLanguage PageSitePageMultiLanguageService
	Revisions     PageSitePageRevisionService
}

// NewPageSitePageService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewPageSitePageService(opts ...option.RequestOption) (r PageSitePageService) {
	r = PageSitePageService{}
	r.options = opts
	r.AbTest = NewPageSitePageAbTestService(opts...)
	r.Batch = NewPageSitePageBatchService(opts...)
	r.Draft = NewPageSitePageDraftService(opts...)
	r.MultiLanguage = NewPageSitePageMultiLanguageService(opts...)
	r.Revisions = NewPageSitePageRevisionService(opts...)
	return
}

// Create a new website page.
func (r *PageSitePageService) New(ctx context.Context, body PageSitePageNewParams, opts ...option.RequestOption) (res *PagesPage, err error) {
	opts = slices.Concat(r.options, opts)
	path := "cms/pages/2026-03/site-pages"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Partially updates a single website page, specified by its ID. You only need to
// specify the column values that you are modifying.
func (r *PageSitePageService) Update(ctx context.Context, objectID string, params PageSitePageUpdateParams, opts ...option.RequestOption) (res *PagesPage, err error) {
	opts = slices.Concat(r.options, opts)
	if objectID == "" {
		err = errors.New("missing required objectId parameter")
		return nil, err
	}
	path := fmt.Sprintf("cms/pages/2026-03/site-pages/%s", url.PathEscape(objectID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return res, err
}

// Retrieve all website pages. Supports paging and filtering. This method would be
// useful for an integration that examined these models and used an external
// service to suggest edits.
func (r *PageSitePageService) List(ctx context.Context, query PageSitePageListParams, opts ...option.RequestOption) (res *pagination.Page[PagesPage], err error) {
	var raw *http.Response
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "cms/pages/2026-03/site-pages"
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

// Retrieve all website pages. Supports paging and filtering. This method would be
// useful for an integration that examined these models and used an external
// service to suggest edits.
func (r *PageSitePageService) ListAutoPaging(ctx context.Context, query PageSitePageListParams, opts ...option.RequestOption) *pagination.PageAutoPager[PagesPage] {
	return pagination.NewPageAutoPager(r.List(ctx, query, opts...))
}

// Delete a website page, specified by its ID.
func (r *PageSitePageService) Delete(ctx context.Context, objectID string, body PageSitePageDeleteParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if objectID == "" {
		err = errors.New("missing required objectId parameter")
		return err
	}
	path := fmt.Sprintf("cms/pages/2026-03/site-pages/%s", url.PathEscape(objectID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, nil, opts...)
	return err
}

// Create a copy of an existing website page.
func (r *PageSitePageService) Clone(ctx context.Context, body PageSitePageCloneParams, opts ...option.RequestOption) (res *PagesPage, err error) {
	opts = slices.Concat(r.options, opts)
	path := "cms/pages/2026-03/site-pages/clone"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Retrieve a website page by its ID.
func (r *PageSitePageService) Get(ctx context.Context, objectID string, query PageSitePageGetParams, opts ...option.RequestOption) (res *PagesPage, err error) {
	opts = slices.Concat(r.options, opts)
	if objectID == "" {
		err = errors.New("missing required objectId parameter")
		return nil, err
	}
	path := fmt.Sprintf("cms/pages/2026-03/site-pages/%s", url.PathEscape(objectID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Schedule a website page to published at a future time.
func (r *PageSitePageService) Schedule(ctx context.Context, body PageSitePageScheduleParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "cms/pages/2026-03/site-pages/schedule"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return err
}

type PageSitePageNewParams struct {
	PagesPage PagesPageParam
	paramObj
}

func (r PageSitePageNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.PagesPage)
}
func (r *PageSitePageNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PageSitePageUpdateParams struct {
	PagesPage PagesPageParam
	// Whether to return only results that have been archived.
	Archived param.Opt[bool] `query:"archived,omitzero" json:"-"`
	paramObj
}

func (r PageSitePageUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.PagesPage)
}
func (r *PageSitePageUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
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
	// The paging cursor token of the last successfully read resource will be returned
	// as the `paging.next.after` JSON property of a paged response containing more
	// results.
	After param.Opt[string] `query:"after,omitzero" json:"-"`
	// Whether to return only results that have been archived.
	Archived param.Opt[bool] `query:"archived,omitzero" json:"-"`
	// Filter pages created after a specific date and time.
	CreatedAfter param.Opt[time.Time] `query:"createdAfter,omitzero" format:"date-time" json:"-"`
	// Filter pages by the exact creation timestamp. Format is date-time.
	CreatedAt param.Opt[time.Time] `query:"createdAt,omitzero" format:"date-time" json:"-"`
	// Filter pages created before a specific date-time.
	CreatedBefore param.Opt[time.Time] `query:"createdBefore,omitzero" format:"date-time" json:"-"`
	// The maximum number of results to display per page.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Specify properties to include in the response.
	Property param.Opt[string] `query:"property,omitzero" json:"-"`
	// Filter pages updated after the specified date-time.
	UpdatedAfter param.Opt[time.Time] `query:"updatedAfter,omitzero" format:"date-time" json:"-"`
	// Filter pages by their exact update timestamp in ISO 8601 format.
	UpdatedAt param.Opt[time.Time] `query:"updatedAt,omitzero" format:"date-time" json:"-"`
	// Filter pages updated before a specific date and time. Format should be
	// date-time.
	UpdatedBefore param.Opt[time.Time] `query:"updatedBefore,omitzero" format:"date-time" json:"-"`
	// Specify the order of results. Accepts an array of field names to sort by.
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

type PageSitePageCloneParams struct {
	ContentCloneRequestVNext ContentCloneRequestVNextParam
	paramObj
}

func (r PageSitePageCloneParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.ContentCloneRequestVNext)
}
func (r *PageSitePageCloneParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PageSitePageGetParams struct {
	// Whether to return only results that have been archived.
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

type PageSitePageScheduleParams struct {
	ContentScheduleRequestVNext ContentScheduleRequestVNextParam
	paramObj
}

func (r PageSitePageScheduleParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.ContentScheduleRequestVNext)
}
func (r *PageSitePageScheduleParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
