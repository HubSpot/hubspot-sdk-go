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

	"github.com/stainless-sdks/hubspot-sdk-go/internal/apijson"
	"github.com/stainless-sdks/hubspot-sdk-go/internal/apiquery"
	shimjson "github.com/stainless-sdks/hubspot-sdk-go/internal/encoding/json"
	"github.com/stainless-sdks/hubspot-sdk-go/internal/requestconfig"
	"github.com/stainless-sdks/hubspot-sdk-go/option"
	"github.com/stainless-sdks/hubspot-sdk-go/packages/pagination"
	"github.com/stainless-sdks/hubspot-sdk-go/packages/param"
)

// PageWebsitePageService contains methods and other services that help with
// interacting with the hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPageWebsitePageService] method instead.
type PageWebsitePageService struct {
	options []option.RequestOption
}

// NewPageWebsitePageService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewPageWebsitePageService(opts ...option.RequestOption) (r PageWebsitePageService) {
	r = PageWebsitePageService{}
	r.options = opts
	return
}

// Create a new website page.
func (r *PageWebsitePageService) New(ctx context.Context, body PageWebsitePageNewParams, opts ...option.RequestOption) (res *PageData, err error) {
	opts = slices.Concat(r.options, opts)
	path := "cms/pages/2026-03/site-pages"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Partially updates a single website page, specified by its ID. You only need to
// specify the column values that you are modifying.
func (r *PageWebsitePageService) Update(ctx context.Context, objectID string, params PageWebsitePageUpdateParams, opts ...option.RequestOption) (res *PageData, err error) {
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
func (r *PageWebsitePageService) List(ctx context.Context, query PageWebsitePageListParams, opts ...option.RequestOption) (res *pagination.Page[PageData], err error) {
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
func (r *PageWebsitePageService) ListAutoPaging(ctx context.Context, query PageWebsitePageListParams, opts ...option.RequestOption) *pagination.PageAutoPager[PageData] {
	return pagination.NewPageAutoPager(r.List(ctx, query, opts...))
}

// Delete a website page, specified by its ID.
func (r *PageWebsitePageService) Delete(ctx context.Context, objectID string, body PageWebsitePageDeleteParams, opts ...option.RequestOption) (err error) {
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
func (r *PageWebsitePageService) Clone(ctx context.Context, body PageWebsitePageCloneParams, opts ...option.RequestOption) (res *PageData, err error) {
	opts = slices.Concat(r.options, opts)
	path := "cms/pages/2026-03/site-pages/clone"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Retrieve a website page by its ID.
func (r *PageWebsitePageService) Get(ctx context.Context, objectID string, query PageWebsitePageGetParams, opts ...option.RequestOption) (res *PageData, err error) {
	opts = slices.Concat(r.options, opts)
	if objectID == "" {
		err = errors.New("missing required objectId parameter")
		return nil, err
	}
	path := fmt.Sprintf("cms/pages/2026-03/site-pages/%s", url.PathEscape(objectID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Retrieve the full draft version of a website page, specified by its ID.
func (r *PageWebsitePageService) GetDraft(ctx context.Context, objectID string, opts ...option.RequestOption) (res *PageData, err error) {
	opts = slices.Concat(r.options, opts)
	if objectID == "" {
		err = errors.New("missing required objectId parameter")
		return nil, err
	}
	path := fmt.Sprintf("cms/pages/2026-03/site-pages/%s/draft", url.PathEscape(objectID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Take any changes from the draft version of the website page and apply them to
// the live version.
func (r *PageWebsitePageService) PublishDraft(ctx context.Context, objectID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if objectID == "" {
		err = errors.New("missing required objectId parameter")
		return err
	}
	path := fmt.Sprintf("cms/pages/2026-03/site-pages/%s/draft/push-live", url.PathEscape(objectID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, nil, opts...)
	return err
}

// Schedule a website page to published at a future time.
func (r *PageWebsitePageService) Schedule(ctx context.Context, body PageWebsitePageScheduleParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "cms/pages/2026-03/site-pages/schedule"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return err
}

// Set a landing page as the primary language of a multi-language group.
func (r *PageWebsitePageService) SetNewLangPrimary(ctx context.Context, body PageWebsitePageSetNewLangPrimaryParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "cms/pages/2026-03/landing-pages/multi-language/set-new-lang-primary"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, nil, opts...)
	return err
}

// Partially update the draft version of a website page, specified by page ID. You
// only need to specify the values for the details that you're modifying.
func (r *PageWebsitePageService) UpdateDraft(ctx context.Context, objectID string, body PageWebsitePageUpdateDraftParams, opts ...option.RequestOption) (res *PageData, err error) {
	opts = slices.Concat(r.options, opts)
	if objectID == "" {
		err = errors.New("missing required objectId parameter")
		return nil, err
	}
	path := fmt.Sprintf("cms/pages/2026-03/site-pages/%s/draft", url.PathEscape(objectID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

type PageWebsitePageNewParams struct {
	PageData PageDataParam
	paramObj
}

func (r PageWebsitePageNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.PageData)
}
func (r *PageWebsitePageNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PageWebsitePageUpdateParams struct {
	PageData PageDataParam
	// Whether to return only results that have been archived.
	Archived param.Opt[bool] `query:"archived,omitzero" json:"-"`
	paramObj
}

func (r PageWebsitePageUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.PageData)
}
func (r *PageWebsitePageUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [PageWebsitePageUpdateParams]'s query parameters as
// `url.Values`.
func (r PageWebsitePageUpdateParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type PageWebsitePageListParams struct {
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

// URLQuery serializes [PageWebsitePageListParams]'s query parameters as
// `url.Values`.
func (r PageWebsitePageListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type PageWebsitePageDeleteParams struct {
	// Whether to return only results that have been archived.
	Archived param.Opt[bool] `query:"archived,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [PageWebsitePageDeleteParams]'s query parameters as
// `url.Values`.
func (r PageWebsitePageDeleteParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type PageWebsitePageCloneParams struct {
	ContentCloneRequestVNext ContentCloneRequestVNextParam
	paramObj
}

func (r PageWebsitePageCloneParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.ContentCloneRequestVNext)
}
func (r *PageWebsitePageCloneParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PageWebsitePageGetParams struct {
	// Whether to return only results that have been archived.
	Archived param.Opt[bool]   `query:"archived,omitzero" json:"-"`
	Property param.Opt[string] `query:"property,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [PageWebsitePageGetParams]'s query parameters as
// `url.Values`.
func (r PageWebsitePageGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type PageWebsitePageScheduleParams struct {
	ContentScheduleRequestVNext ContentScheduleRequestVNextParam
	paramObj
}

func (r PageWebsitePageScheduleParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.ContentScheduleRequestVNext)
}
func (r *PageWebsitePageScheduleParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PageWebsitePageSetNewLangPrimaryParams struct {
	SetNewLanguagePrimaryRequestVNext SetNewLanguagePrimaryRequestVNextParam
	paramObj
}

func (r PageWebsitePageSetNewLangPrimaryParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.SetNewLanguagePrimaryRequestVNext)
}
func (r *PageWebsitePageSetNewLangPrimaryParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PageWebsitePageUpdateDraftParams struct {
	PageData PageDataParam
	paramObj
}

func (r PageWebsitePageUpdateDraftParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.PageData)
}
func (r *PageWebsitePageUpdateDraftParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
