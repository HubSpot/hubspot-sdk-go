// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cms

import (
	"context"
	"net/http"
	"net/url"
	"slices"

	"github.com/HubSpot/hubspot-sdk-go/internal/apijson"
	"github.com/HubSpot/hubspot-sdk-go/internal/apiquery"
	shimjson "github.com/HubSpot/hubspot-sdk-go/internal/encoding/json"
	"github.com/HubSpot/hubspot-sdk-go/internal/requestconfig"
	"github.com/HubSpot/hubspot-sdk-go/option"
	"github.com/HubSpot/hubspot-sdk-go/packages/param"
	"github.com/HubSpot/hubspot-sdk-go/shared"
)

// PageBatchService contains methods and other services that help with interacting
// with the hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPageBatchService] method instead.
type PageBatchService struct {
	options []option.RequestOption
}

// NewPageBatchService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewPageBatchService(opts ...option.RequestOption) (r PageBatchService) {
	r = PageBatchService{}
	r.options = opts
	return
}

// Create a batch of folders as detailed in the request body.
func (r *PageBatchService) NewFolders(ctx context.Context, body PageBatchNewFoldersParams, opts ...option.RequestOption) (res *BatchResponseContentFolder, err error) {
	opts = slices.Concat(r.options, opts)
	path := "cms/pages/2026-03/landing-pages/folders/batch/create"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Create a batch of landing pages as detailed in the request body.
func (r *PageBatchService) NewLandingPages(ctx context.Context, body PageBatchNewLandingPagesParams, opts ...option.RequestOption) (res *BatchResponsePage, err error) {
	opts = slices.Concat(r.options, opts)
	path := "cms/pages/2026-03/landing-pages/batch/create"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Create a batch of website pages as specified in the request body.
func (r *PageBatchService) NewSitePages(ctx context.Context, body PageBatchNewSitePagesParams, opts ...option.RequestOption) (res *BatchResponsePage, err error) {
	opts = slices.Concat(r.options, opts)
	path := "cms/pages/2026-03/site-pages/batch/create"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Delete a batch of folders as specified in the request body.
func (r *PageBatchService) DeleteFolders(ctx context.Context, body PageBatchDeleteFoldersParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "cms/pages/2026-03/landing-pages/folders/batch/archive"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return err
}

// Delete landing pages specified by ID in the request body. Note: this is not the
// same as the dashboard `archive` function. To perform a dashboard `archive` send
// an normal update with the `archivedInDashboard` field set to `true`.
func (r *PageBatchService) DeleteLandingPages(ctx context.Context, body PageBatchDeleteLandingPagesParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "cms/pages/2026-03/landing-pages/batch/archive"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return err
}

// Delete a batch of website pages as specified in the request body. Note that this
// is not the same as the dashboard `archive` function. To perform a dashboard
// `archive` send an normal update with the `archivedInDashboard` field set to
// `true`.
func (r *PageBatchService) DeleteSitePages(ctx context.Context, body PageBatchDeleteSitePagesParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "cms/pages/2026-03/site-pages/batch/archive"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return err
}

// Retrieve a batch of landing pages as specified in the request body.
func (r *PageBatchService) GetLandingPages(ctx context.Context, params PageBatchGetLandingPagesParams, opts ...option.RequestOption) (res *BatchResponsePage, err error) {
	opts = slices.Concat(r.options, opts)
	path := "cms/pages/2026-03/landing-pages/batch/read"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Retrieve a batch of website pages as specified in the request body.
func (r *PageBatchService) GetSitePages(ctx context.Context, params PageBatchGetSitePagesParams, opts ...option.RequestOption) (res *BatchResponsePage, err error) {
	opts = slices.Concat(r.options, opts)
	path := "cms/pages/2026-03/site-pages/batch/read"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Update a batch of landing page folders as specified in the request body.
func (r *PageBatchService) UpdateFolders(ctx context.Context, params PageBatchUpdateFoldersParams, opts ...option.RequestOption) (res *BatchResponseContentFolder, err error) {
	opts = slices.Concat(r.options, opts)
	path := "cms/pages/2026-03/landing-pages/folders/batch/update"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Update a batch of landing pages as specified in the request body.
func (r *PageBatchService) UpdateLandingPages(ctx context.Context, params PageBatchUpdateLandingPagesParams, opts ...option.RequestOption) (res *BatchResponsePage, err error) {
	opts = slices.Concat(r.options, opts)
	path := "cms/pages/2026-03/landing-pages/batch/update"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Update a batch of website pages as specified in the request body.
func (r *PageBatchService) UpdateSitePages(ctx context.Context, params PageBatchUpdateSitePagesParams, opts ...option.RequestOption) (res *BatchResponsePage, err error) {
	opts = slices.Concat(r.options, opts)
	path := "cms/pages/2026-03/site-pages/batch/update"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

type PageBatchNewFoldersParams struct {
	BatchInputContentFolder BatchInputContentFolderParam
	paramObj
}

func (r PageBatchNewFoldersParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.BatchInputContentFolder)
}
func (r *PageBatchNewFoldersParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PageBatchNewLandingPagesParams struct {
	BatchInputPage BatchInputPageParam
	paramObj
}

func (r PageBatchNewLandingPagesParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.BatchInputPage)
}
func (r *PageBatchNewLandingPagesParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PageBatchNewSitePagesParams struct {
	BatchInputPage BatchInputPageParam
	paramObj
}

func (r PageBatchNewSitePagesParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.BatchInputPage)
}
func (r *PageBatchNewSitePagesParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PageBatchDeleteFoldersParams struct {
	BatchInputString shared.BatchInputStringParam
	paramObj
}

func (r PageBatchDeleteFoldersParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.BatchInputString)
}
func (r *PageBatchDeleteFoldersParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PageBatchDeleteLandingPagesParams struct {
	BatchInputString shared.BatchInputStringParam
	paramObj
}

func (r PageBatchDeleteLandingPagesParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.BatchInputString)
}
func (r *PageBatchDeleteLandingPagesParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PageBatchDeleteSitePagesParams struct {
	BatchInputString shared.BatchInputStringParam
	paramObj
}

func (r PageBatchDeleteSitePagesParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.BatchInputString)
}
func (r *PageBatchDeleteSitePagesParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PageBatchGetLandingPagesParams struct {
	BatchInputString shared.BatchInputStringParam
	// Whether to return only results that have been archived.
	Archived param.Opt[bool] `query:"archived,omitzero" json:"-"`
	paramObj
}

func (r PageBatchGetLandingPagesParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.BatchInputString)
}
func (r *PageBatchGetLandingPagesParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [PageBatchGetLandingPagesParams]'s query parameters as
// `url.Values`.
func (r PageBatchGetLandingPagesParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type PageBatchGetSitePagesParams struct {
	BatchInputString shared.BatchInputStringParam
	// Whether to return only results that have been archived.
	Archived param.Opt[bool] `query:"archived,omitzero" json:"-"`
	paramObj
}

func (r PageBatchGetSitePagesParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.BatchInputString)
}
func (r *PageBatchGetSitePagesParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [PageBatchGetSitePagesParams]'s query parameters as
// `url.Values`.
func (r PageBatchGetSitePagesParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type PageBatchUpdateFoldersParams struct {
	BatchInputJsonNode BatchInputJsonNodeParam
	// Whether to return only results that have been archived.
	Archived param.Opt[bool] `query:"archived,omitzero" json:"-"`
	paramObj
}

func (r PageBatchUpdateFoldersParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.BatchInputJsonNode)
}
func (r *PageBatchUpdateFoldersParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [PageBatchUpdateFoldersParams]'s query parameters as
// `url.Values`.
func (r PageBatchUpdateFoldersParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type PageBatchUpdateLandingPagesParams struct {
	BatchInputJsonNode BatchInputJsonNodeParam
	// Whether to return only results that have been archived.
	Archived param.Opt[bool] `query:"archived,omitzero" json:"-"`
	paramObj
}

func (r PageBatchUpdateLandingPagesParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.BatchInputJsonNode)
}
func (r *PageBatchUpdateLandingPagesParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [PageBatchUpdateLandingPagesParams]'s query parameters as
// `url.Values`.
func (r PageBatchUpdateLandingPagesParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type PageBatchUpdateSitePagesParams struct {
	BatchInputJsonNode BatchInputJsonNodeParam
	// Whether to return only results that have been archived.
	Archived param.Opt[bool] `query:"archived,omitzero" json:"-"`
	paramObj
}

func (r PageBatchUpdateSitePagesParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.BatchInputJsonNode)
}
func (r *PageBatchUpdateSitePagesParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [PageBatchUpdateSitePagesParams]'s query parameters as
// `url.Values`.
func (r PageBatchUpdateSitePagesParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
