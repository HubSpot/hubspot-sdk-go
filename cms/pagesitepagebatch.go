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

// PageSitePageBatchService contains methods and other services that help with
// interacting with the hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPageSitePageBatchService] method instead.
type PageSitePageBatchService struct {
	options []option.RequestOption
}

// NewPageSitePageBatchService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewPageSitePageBatchService(opts ...option.RequestOption) (r PageSitePageBatchService) {
	r = PageSitePageBatchService{}
	r.options = opts
	return
}

// Create a batch of website pages as specified in the request body.
func (r *PageSitePageBatchService) NewSitePages(ctx context.Context, body PageSitePageBatchNewSitePagesParams, opts ...option.RequestOption) (res *BatchResponsePage, err error) {
	opts = slices.Concat(r.options, opts)
	path := "cms/pages/2026-03/site-pages/batch/create"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Delete a batch of website pages as specified in the request body. Note that this
// is not the same as the dashboard `archive` function. To perform a dashboard
// `archive` send an normal update with the `archivedInDashboard` field set to
// `true`.
func (r *PageSitePageBatchService) DeleteSitePages(ctx context.Context, body PageSitePageBatchDeleteSitePagesParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "cms/pages/2026-03/site-pages/batch/archive"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return err
}

// Retrieve a batch of website pages as specified in the request body.
func (r *PageSitePageBatchService) GetSitePages(ctx context.Context, params PageSitePageBatchGetSitePagesParams, opts ...option.RequestOption) (res *BatchResponsePage, err error) {
	opts = slices.Concat(r.options, opts)
	path := "cms/pages/2026-03/site-pages/batch/read"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Update a batch of website pages as specified in the request body.
func (r *PageSitePageBatchService) UpdateSitePages(ctx context.Context, params PageSitePageBatchUpdateSitePagesParams, opts ...option.RequestOption) (res *BatchResponsePage, err error) {
	opts = slices.Concat(r.options, opts)
	path := "cms/pages/2026-03/site-pages/batch/update"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

type PageSitePageBatchNewSitePagesParams struct {
	BatchInputPage BatchInputPageParam
	paramObj
}

func (r PageSitePageBatchNewSitePagesParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.BatchInputPage)
}
func (r *PageSitePageBatchNewSitePagesParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PageSitePageBatchDeleteSitePagesParams struct {
	BatchInputString shared.BatchInputStringParam
	paramObj
}

func (r PageSitePageBatchDeleteSitePagesParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.BatchInputString)
}
func (r *PageSitePageBatchDeleteSitePagesParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PageSitePageBatchGetSitePagesParams struct {
	BatchInputString shared.BatchInputStringParam
	// Whether to return only results that have been archived.
	Archived param.Opt[bool] `query:"archived,omitzero" json:"-"`
	paramObj
}

func (r PageSitePageBatchGetSitePagesParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.BatchInputString)
}
func (r *PageSitePageBatchGetSitePagesParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [PageSitePageBatchGetSitePagesParams]'s query parameters as
// `url.Values`.
func (r PageSitePageBatchGetSitePagesParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type PageSitePageBatchUpdateSitePagesParams struct {
	BatchInputJsonNode BatchInputJsonNodeParam
	// Whether to return only results that have been archived.
	Archived param.Opt[bool] `query:"archived,omitzero" json:"-"`
	paramObj
}

func (r PageSitePageBatchUpdateSitePagesParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.BatchInputJsonNode)
}
func (r *PageSitePageBatchUpdateSitePagesParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [PageSitePageBatchUpdateSitePagesParams]'s query parameters
// as `url.Values`.
func (r PageSitePageBatchUpdateSitePagesParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
