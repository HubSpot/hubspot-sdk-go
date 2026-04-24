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

// PageLandingPageService contains methods and other services that help with
// interacting with the hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPageLandingPageService] method instead.
type PageLandingPageService struct {
	options []option.RequestOption
}

// NewPageLandingPageService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewPageLandingPageService(opts ...option.RequestOption) (r PageLandingPageService) {
	r = PageLandingPageService{}
	r.options = opts
	return
}

// Create a new landing page.
func (r *PageLandingPageService) New(ctx context.Context, body PageLandingPageNewParams, opts ...option.RequestOption) (res *PageData, err error) {
	opts = slices.Concat(r.options, opts)
	path := "cms/pages/2026-03/landing-pages"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Sparse updates a single Landing Page object identified by the id in the path.
// You only need to specify the column values that you are modifying.
func (r *PageLandingPageService) Update(ctx context.Context, objectID string, params PageLandingPageUpdateParams, opts ...option.RequestOption) (res *PageData, err error) {
	opts = slices.Concat(r.options, opts)
	if objectID == "" {
		err = errors.New("missing required objectId parameter")
		return nil, err
	}
	path := fmt.Sprintf("cms/pages/2026-03/landing-pages/%s", url.PathEscape(objectID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return res, err
}

// Retrieve a list of landing pages in your HubSpot account. This endpoint allows
// you to filter landing pages based on creation and update timestamps, sort them,
// and paginate through results. You can also choose to include archived pages or
// specify certain properties to be included in the response.
func (r *PageLandingPageService) List(ctx context.Context, query PageLandingPageListParams, opts ...option.RequestOption) (res *pagination.Page[PageData], err error) {
	var raw *http.Response
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "cms/pages/2026-03/landing-pages"
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

// Retrieve a list of landing pages in your HubSpot account. This endpoint allows
// you to filter landing pages based on creation and update timestamps, sort them,
// and paginate through results. You can also choose to include archived pages or
// specify certain properties to be included in the response.
func (r *PageLandingPageService) ListAutoPaging(ctx context.Context, query PageLandingPageListParams, opts ...option.RequestOption) *pagination.PageAutoPager[PageData] {
	return pagination.NewPageAutoPager(r.List(ctx, query, opts...))
}

// Delete a landing page, specified by its ID.
func (r *PageLandingPageService) Delete(ctx context.Context, objectID string, body PageLandingPageDeleteParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if objectID == "" {
		err = errors.New("missing required objectId parameter")
		return err
	}
	path := fmt.Sprintf("cms/pages/2026-03/landing-pages/%s", url.PathEscape(objectID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, nil, opts...)
	return err
}

// Create a copy of an existing landing page.
func (r *PageLandingPageService) Clone(ctx context.Context, body PageLandingPageCloneParams, opts ...option.RequestOption) (res *PageData, err error) {
	opts = slices.Concat(r.options, opts)
	path := "cms/pages/2026-03/landing-pages/clone"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Retrieve a landing page, specified by its ID.
func (r *PageLandingPageService) Get(ctx context.Context, objectID string, query PageLandingPageGetParams, opts ...option.RequestOption) (res *PageData, err error) {
	opts = slices.Concat(r.options, opts)
	if objectID == "" {
		err = errors.New("missing required objectId parameter")
		return nil, err
	}
	path := fmt.Sprintf("cms/pages/2026-03/landing-pages/%s", url.PathEscape(objectID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Retrieve the full draft version of a landing page, specified by page ID.
func (r *PageLandingPageService) GetDraft(ctx context.Context, objectID string, opts ...option.RequestOption) (res *PageData, err error) {
	opts = slices.Concat(r.options, opts)
	if objectID == "" {
		err = errors.New("missing required objectId parameter")
		return nil, err
	}
	path := fmt.Sprintf("cms/pages/2026-03/landing-pages/%s/draft", url.PathEscape(objectID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Take any changes from the draft version of the Landing Page and apply them to
// the live version.
func (r *PageLandingPageService) PushDraftLive(ctx context.Context, objectID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if objectID == "" {
		err = errors.New("missing required objectId parameter")
		return err
	}
	path := fmt.Sprintf("cms/pages/2026-03/landing-pages/%s/draft/push-live", url.PathEscape(objectID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, nil, opts...)
	return err
}

// Discards any edits and resets the draft to match the live version.
func (r *PageLandingPageService) ResetDraft(ctx context.Context, objectID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if objectID == "" {
		err = errors.New("missing required objectId parameter")
		return err
	}
	path := fmt.Sprintf("cms/pages/2026-03/landing-pages/%s/draft/reset", url.PathEscape(objectID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, nil, opts...)
	return err
}

// Schedule a landing page to be published.
func (r *PageLandingPageService) Schedule(ctx context.Context, body PageLandingPageScheduleParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "cms/pages/2026-03/landing-pages/schedule"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return err
}

// Partially updates the draft version of a single landing page, specified by its
// ID. You only need to specify the column values that you are modifying.
func (r *PageLandingPageService) UpdateDraft(ctx context.Context, objectID string, body PageLandingPageUpdateDraftParams, opts ...option.RequestOption) (res *PageData, err error) {
	opts = slices.Concat(r.options, opts)
	if objectID == "" {
		err = errors.New("missing required objectId parameter")
		return nil, err
	}
	path := fmt.Sprintf("cms/pages/2026-03/landing-pages/%s/draft", url.PathEscape(objectID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

type PageLandingPageNewParams struct {
	PageData PageDataParam
	paramObj
}

func (r PageLandingPageNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.PageData)
}
func (r *PageLandingPageNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PageLandingPageUpdateParams struct {
	PageData PageDataParam
	// Whether to return only results that have been archived.
	Archived param.Opt[bool] `query:"archived,omitzero" json:"-"`
	paramObj
}

func (r PageLandingPageUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.PageData)
}
func (r *PageLandingPageUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
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
	// A cursor token for pagination. Use the value from the previous response's
	// paging.next.after field.
	After param.Opt[string] `query:"after,omitzero" json:"-"`
	// Whether to return only results that have been archived.
	Archived param.Opt[bool] `query:"archived,omitzero" json:"-"`
	// Filter landing pages created after a specific date and time.
	CreatedAfter param.Opt[time.Time] `query:"createdAfter,omitzero" format:"date-time" json:"-"`
	// Filter landing pages by their creation timestamp.
	CreatedAt param.Opt[time.Time] `query:"createdAt,omitzero" format:"date-time" json:"-"`
	// Filter landing pages created before a specific date and time.
	CreatedBefore param.Opt[time.Time] `query:"createdBefore,omitzero" format:"date-time" json:"-"`
	// The maximum number of results to display per page.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Specify which properties of the landing pages to include in the response.
	Property param.Opt[string] `query:"property,omitzero" json:"-"`
	// Filter landing pages updated after a specific date and time.
	UpdatedAfter param.Opt[time.Time] `query:"updatedAfter,omitzero" format:"date-time" json:"-"`
	// Filter landing pages by their last updated timestamp.
	UpdatedAt param.Opt[time.Time] `query:"updatedAt,omitzero" format:"date-time" json:"-"`
	// Filter landing pages updated before a specific date and time.
	UpdatedBefore param.Opt[time.Time] `query:"updatedBefore,omitzero" format:"date-time" json:"-"`
	// Specify the order in which results are returned. Accepts an array of strings.
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

type PageLandingPageCloneParams struct {
	ContentCloneRequestVNext ContentCloneRequestVNextParam
	paramObj
}

func (r PageLandingPageCloneParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.ContentCloneRequestVNext)
}
func (r *PageLandingPageCloneParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PageLandingPageGetParams struct {
	// Whether to return only results that have been archived.
	Archived param.Opt[bool] `query:"archived,omitzero" json:"-"`
	// A specific property of the landing page to include in the response.
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

type PageLandingPageScheduleParams struct {
	ContentScheduleRequestVNext ContentScheduleRequestVNextParam
	paramObj
}

func (r PageLandingPageScheduleParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.ContentScheduleRequestVNext)
}
func (r *PageLandingPageScheduleParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PageLandingPageUpdateDraftParams struct {
	PageData PageDataParam
	paramObj
}

func (r PageLandingPageUpdateDraftParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.PageData)
}
func (r *PageLandingPageUpdateDraftParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
