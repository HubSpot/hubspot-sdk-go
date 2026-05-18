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

// PageLandingPageBatchService contains methods and other services that help with
// interacting with the hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPageLandingPageBatchService] method instead.
type PageLandingPageBatchService struct {
	options []option.RequestOption
}

// NewPageLandingPageBatchService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewPageLandingPageBatchService(opts ...option.RequestOption) (r PageLandingPageBatchService) {
	r = PageLandingPageBatchService{}
	r.options = opts
	return
}

// Create a batch of landing pages as detailed in the request body.
func (r *PageLandingPageBatchService) NewLandingPages(ctx context.Context, body PageLandingPageBatchNewLandingPagesParams, opts ...option.RequestOption) (res *BatchResponsePage, err error) {
	opts = slices.Concat(r.options, opts)
	path := "cms/pages/2026-03/landing-pages/batch/create"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Delete landing pages specified by ID in the request body. Note: this is not the
// same as the dashboard `archive` function. To perform a dashboard `archive` send
// an normal update with the `archivedInDashboard` field set to `true`.
func (r *PageLandingPageBatchService) DeleteLandingPages(ctx context.Context, body PageLandingPageBatchDeleteLandingPagesParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "cms/pages/2026-03/landing-pages/batch/archive"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return err
}

// Retrieve a batch of landing pages as specified in the request body.
func (r *PageLandingPageBatchService) GetLandingPages(ctx context.Context, params PageLandingPageBatchGetLandingPagesParams, opts ...option.RequestOption) (res *BatchResponsePage, err error) {
	opts = slices.Concat(r.options, opts)
	path := "cms/pages/2026-03/landing-pages/batch/read"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Update a batch of landing pages as specified in the request body.
func (r *PageLandingPageBatchService) UpdateLandingPages(ctx context.Context, params PageLandingPageBatchUpdateLandingPagesParams, opts ...option.RequestOption) (res *BatchResponsePage, err error) {
	opts = slices.Concat(r.options, opts)
	path := "cms/pages/2026-03/landing-pages/batch/update"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

type PageLandingPageBatchNewLandingPagesParams struct {
	BatchInputPage BatchInputPageParam
	paramObj
}

func (r PageLandingPageBatchNewLandingPagesParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.BatchInputPage)
}
func (r *PageLandingPageBatchNewLandingPagesParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PageLandingPageBatchDeleteLandingPagesParams struct {
	BatchInputString shared.BatchInputStringParam
	paramObj
}

func (r PageLandingPageBatchDeleteLandingPagesParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.BatchInputString)
}
func (r *PageLandingPageBatchDeleteLandingPagesParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PageLandingPageBatchGetLandingPagesParams struct {
	BatchInputString shared.BatchInputStringParam
	// Whether to return only results that have been archived.
	Archived param.Opt[bool] `query:"archived,omitzero" json:"-"`
	paramObj
}

func (r PageLandingPageBatchGetLandingPagesParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.BatchInputString)
}
func (r *PageLandingPageBatchGetLandingPagesParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [PageLandingPageBatchGetLandingPagesParams]'s query
// parameters as `url.Values`.
func (r PageLandingPageBatchGetLandingPagesParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type PageLandingPageBatchUpdateLandingPagesParams struct {
	BatchInputJsonNode BatchInputJsonNodeParam
	// Whether to return only results that have been archived.
	Archived param.Opt[bool] `query:"archived,omitzero" json:"-"`
	paramObj
}

func (r PageLandingPageBatchUpdateLandingPagesParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.BatchInputJsonNode)
}
func (r *PageLandingPageBatchUpdateLandingPagesParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [PageLandingPageBatchUpdateLandingPagesParams]'s query
// parameters as `url.Values`.
func (r PageLandingPageBatchUpdateLandingPagesParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
