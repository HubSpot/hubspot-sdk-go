// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cms

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"
	"slices"

	"github.com/stainless-sdks/hubspot-sdk-go/internal/apiquery"
	shimjson "github.com/stainless-sdks/hubspot-sdk-go/internal/encoding/json"
	"github.com/stainless-sdks/hubspot-sdk-go/internal/requestconfig"
	"github.com/stainless-sdks/hubspot-sdk-go/option"
	"github.com/stainless-sdks/hubspot-sdk-go/packages/param"
	"github.com/stainless-sdks/hubspot-sdk-go/shared"
)

// BlogPostBatchService contains methods and other services that help with
// interacting with the Hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBlogPostBatchService] method instead.
type BlogPostBatchService struct {
	Options []option.RequestOption
}

// NewBlogPostBatchService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewBlogPostBatchService(opts ...option.RequestOption) (r BlogPostBatchService) {
	r = BlogPostBatchService{}
	r.Options = opts
	return
}

// Create a batch of blog posts, specifying their content in the request body.
func (r *BlogPostBatchService) New(ctx context.Context, body BlogPostBatchNewParams, opts ...option.RequestOption) (res *BatchResponseBlogPost, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "cms/v3/blogs/posts/batch/create"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Update a batch of blog posts.
func (r *BlogPostBatchService) Update(ctx context.Context, params BlogPostBatchUpdateParams, opts ...option.RequestOption) (res *BatchResponseBlogPost, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "cms/v3/blogs/posts/batch/update"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Delete a blog post by ID. Note: This is not the same as the in-app `archive`
// function. To perform a dashboard `archive` send an normal update with the
// `archivedInDashboard` field set to `true`.
func (r *BlogPostBatchService) Delete(ctx context.Context, body BlogPostBatchDeleteParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "cms/v3/blogs/posts/batch/archive"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Retrieve a batch of blog posts by ID. identified in the request body.
func (r *BlogPostBatchService) Get(ctx context.Context, params BlogPostBatchGetParams, opts ...option.RequestOption) (res *BatchResponseBlogPost, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "cms/v3/blogs/posts/batch/read"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

type BlogPostBatchNewParams struct {
	// Wrapper for providing an array of blog posts as inputs.
	BatchInputBlogPost BatchInputBlogPostParam
	paramObj
}

func (r BlogPostBatchNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.BatchInputBlogPost)
}
func (r *BlogPostBatchNewParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.BatchInputBlogPost)
}

type BlogPostBatchUpdateParams struct {
	// Wrapper for providing an array of JSON nodes as inputs.
	BatchInputJsonNode BatchInputJsonNodeParam
	// Specifies whether to update deleted Blog Posts. Defaults to `false`.
	Archived param.Opt[bool] `query:"archived,omitzero" json:"-"`
	paramObj
}

func (r BlogPostBatchUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.BatchInputJsonNode)
}
func (r *BlogPostBatchUpdateParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.BatchInputJsonNode)
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
	// Wrapper for providing an array of strings as inputs.
	BatchInputString shared.BatchInputStringParam
	paramObj
}

func (r BlogPostBatchDeleteParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.BatchInputString)
}
func (r *BlogPostBatchDeleteParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.BatchInputString)
}

type BlogPostBatchGetParams struct {
	// Wrapper for providing an array of strings as inputs.
	BatchInputString shared.BatchInputStringParam
	// Specifies whether to return deleted blog posts Defaults to `false`.
	Archived param.Opt[bool] `query:"archived,omitzero" json:"-"`
	paramObj
}

func (r BlogPostBatchGetParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.BatchInputString)
}
func (r *BlogPostBatchGetParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.BatchInputString)
}

// URLQuery serializes [BlogPostBatchGetParams]'s query parameters as `url.Values`.
func (r BlogPostBatchGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
