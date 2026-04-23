// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cms

import (
	"context"
	"net/http"
	"net/url"
	"slices"

	"github.com/stainless-sdks/hubspot-sdk-go/internal/apijson"
	"github.com/stainless-sdks/hubspot-sdk-go/internal/apiquery"
	shimjson "github.com/stainless-sdks/hubspot-sdk-go/internal/encoding/json"
	"github.com/stainless-sdks/hubspot-sdk-go/internal/requestconfig"
	"github.com/stainless-sdks/hubspot-sdk-go/option"
	"github.com/stainless-sdks/hubspot-sdk-go/packages/param"
	"github.com/stainless-sdks/hubspot-sdk-go/shared"
)

// BlogPostBatchService contains methods and other services that help with
// interacting with the hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBlogPostBatchService] method instead.
type BlogPostBatchService struct {
	options []option.RequestOption
}

// NewBlogPostBatchService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewBlogPostBatchService(opts ...option.RequestOption) (r BlogPostBatchService) {
	r = BlogPostBatchService{}
	r.options = opts
	return
}

// Create a batch of blog posts, specifying their content in the request body.
func (r *BlogPostBatchService) New(ctx context.Context, body BlogPostBatchNewParams, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "cms/blogs/2026-03/posts/batch/create"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Update a batch of blog posts.
func (r *BlogPostBatchService) Update(ctx context.Context, params BlogPostBatchUpdateParams, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "cms/blogs/2026-03/posts/batch/update"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Delete a blog post by ID. Note: This is not the same as the in-app `archive`
// function. To perform a dashboard `archive` send an normal update with the
// `archivedInDashboard` field set to `true`.
func (r *BlogPostBatchService) Delete(ctx context.Context, body BlogPostBatchDeleteParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "cms/blogs/2026-03/posts/batch/archive"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return err
}

// Retrieve a batch of blog posts by ID. identified in the request body.
func (r *BlogPostBatchService) Get(ctx context.Context, params BlogPostBatchGetParams, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "cms/blogs/2026-03/posts/batch/read"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

type BlogPostBatchNewParams struct {
	BatchInputBlogPost BatchInputBlogPostParam
	paramObj
}

func (r BlogPostBatchNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.BatchInputBlogPost)
}
func (r *BlogPostBatchNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BlogPostBatchUpdateParams struct {
	BatchInputJsonNode BatchInputJsonNodeParam
	// Whether to return only results that have been archived.
	Archived param.Opt[bool] `query:"archived,omitzero" json:"-"`
	paramObj
}

func (r BlogPostBatchUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.BatchInputJsonNode)
}
func (r *BlogPostBatchUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [BlogPostBatchUpdateParams]'s query parameters as
// `url.Values`.
func (r BlogPostBatchUpdateParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type BlogPostBatchDeleteParams struct {
	BatchInputString shared.BatchInputStringParam
	paramObj
}

func (r BlogPostBatchDeleteParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.BatchInputString)
}
func (r *BlogPostBatchDeleteParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BlogPostBatchGetParams struct {
	BatchInputString shared.BatchInputStringParam
	// Whether to return only results that have been archived.
	Archived param.Opt[bool] `query:"archived,omitzero" json:"-"`
	paramObj
}

func (r BlogPostBatchGetParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.BatchInputString)
}
func (r *BlogPostBatchGetParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [BlogPostBatchGetParams]'s query parameters as `url.Values`.
func (r BlogPostBatchGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
