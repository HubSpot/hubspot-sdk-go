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

// BlogAuthorBatchService contains methods and other services that help with
// interacting with the hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBlogAuthorBatchService] method instead.
type BlogAuthorBatchService struct {
	options []option.RequestOption
}

// NewBlogAuthorBatchService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewBlogAuthorBatchService(opts ...option.RequestOption) (r BlogAuthorBatchService) {
	r = BlogAuthorBatchService{}
	r.options = opts
	return
}

// Create the Blog Author objects detailed in the request body.
func (r *BlogAuthorBatchService) New(ctx context.Context, body BlogAuthorBatchNewParams, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "cms/blogs/2026-03/authors/batch/create"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Update the Blog Author objects identified in the request body.
func (r *BlogAuthorBatchService) Update(ctx context.Context, params BlogAuthorBatchUpdateParams, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "cms/blogs/2026-03/authors/batch/update"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Delete the Blog Author objects identified in the request body.
func (r *BlogAuthorBatchService) Delete(ctx context.Context, body BlogAuthorBatchDeleteParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "cms/blogs/2026-03/authors/batch/archive"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return err
}

// Retrieve the Blog Author objects identified in the request body.
func (r *BlogAuthorBatchService) Get(ctx context.Context, params BlogAuthorBatchGetParams, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "cms/blogs/2026-03/authors/batch/read"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

type BlogAuthorBatchNewParams struct {
	BatchInputBlogAuthor BatchInputBlogAuthorParam
	paramObj
}

func (r BlogAuthorBatchNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.BatchInputBlogAuthor)
}
func (r *BlogAuthorBatchNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BlogAuthorBatchUpdateParams struct {
	BatchInputJsonNode BatchInputJsonNodeParam
	// Whether to return only results that have been archived.
	Archived param.Opt[bool] `query:"archived,omitzero" json:"-"`
	paramObj
}

func (r BlogAuthorBatchUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.BatchInputJsonNode)
}
func (r *BlogAuthorBatchUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [BlogAuthorBatchUpdateParams]'s query parameters as
// `url.Values`.
func (r BlogAuthorBatchUpdateParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type BlogAuthorBatchDeleteParams struct {
	BatchInputString shared.BatchInputStringParam
	paramObj
}

func (r BlogAuthorBatchDeleteParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.BatchInputString)
}
func (r *BlogAuthorBatchDeleteParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BlogAuthorBatchGetParams struct {
	BatchInputString shared.BatchInputStringParam
	// Whether to return only results that have been archived.
	Archived param.Opt[bool] `query:"archived,omitzero" json:"-"`
	paramObj
}

func (r BlogAuthorBatchGetParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.BatchInputString)
}
func (r *BlogAuthorBatchGetParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [BlogAuthorBatchGetParams]'s query parameters as
// `url.Values`.
func (r BlogAuthorBatchGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
