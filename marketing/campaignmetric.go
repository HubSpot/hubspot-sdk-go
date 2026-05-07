// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package marketing

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/HubSpot/hubspot-sdk-go/internal/apiquery"
	"github.com/HubSpot/hubspot-sdk-go/internal/requestconfig"
	"github.com/HubSpot/hubspot-sdk-go/option"
	"github.com/HubSpot/hubspot-sdk-go/packages/pagination"
	"github.com/HubSpot/hubspot-sdk-go/packages/param"
)

// CampaignMetricService contains methods and other services that help with
// interacting with the hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCampaignMetricService] method instead.
type CampaignMetricService struct {
	options []option.RequestOption
}

// NewCampaignMetricService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCampaignMetricService(opts ...option.RequestOption) (r CampaignMetricService) {
	r = CampaignMetricService{}
	r.options = opts
	return
}

// This endpoint retrieves key attribution metrics for a specified campaign, such
// as sessions, new contacts, and influenced contacts.
func (r *CampaignMetricService) GetAttributionMetrics(ctx context.Context, campaignGuid string, query CampaignMetricGetAttributionMetricsParams, opts ...option.RequestOption) (res *MetricsCounters, err error) {
	opts = slices.Concat(r.options, opts)
	if campaignGuid == "" {
		err = errors.New("missing required campaignGuid parameter")
		return nil, err
	}
	path := fmt.Sprintf("marketing/campaigns/2026-03/%s/reports/metrics", url.PathEscape(campaignGuid))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Fetch revenue attribution report data for a specified campaign
func (r *CampaignMetricService) GetRevenueAttribution(ctx context.Context, campaignGuid string, query CampaignMetricGetRevenueAttributionParams, opts ...option.RequestOption) (res *RevenueAttributionAggregate, err error) {
	opts = slices.Concat(r.options, opts)
	if campaignGuid == "" {
		err = errors.New("missing required campaignGuid parameter")
		return nil, err
	}
	path := fmt.Sprintf("marketing/campaigns/2026-03/%s/reports/revenue", url.PathEscape(campaignGuid))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Fetch the list of contact IDs for the specified campaign and contact type
func (r *CampaignMetricService) ListContactIDsByType(ctx context.Context, contactType string, params CampaignMetricListContactIDsByTypeParams, opts ...option.RequestOption) (res *pagination.Page[ContactReference], err error) {
	var raw *http.Response
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.CampaignGuid == "" {
		err = errors.New("missing required campaignGuid parameter")
		return nil, err
	}
	if contactType == "" {
		err = errors.New("missing required contactType parameter")
		return nil, err
	}
	path := fmt.Sprintf("marketing/campaigns/2026-03/%s/reports/contacts/%s", url.PathEscape(params.CampaignGuid), url.PathEscape(contactType))
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, params, &res, opts...)
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

// Fetch the list of contact IDs for the specified campaign and contact type
func (r *CampaignMetricService) ListContactIDsByTypeAutoPaging(ctx context.Context, contactType string, params CampaignMetricListContactIDsByTypeParams, opts ...option.RequestOption) *pagination.PageAutoPager[ContactReference] {
	return pagination.NewPageAutoPager(r.ListContactIDsByType(ctx, contactType, params, opts...))
}

type CampaignMetricGetAttributionMetricsParams struct {
	// The end date for fetching attribution data, in YYYY-MM-DD format. Optional.
	// Example: 2000-01-27
	EndDate param.Opt[string] `query:"endDate,omitzero" json:"-"`
	// The start date for fetching attribution data, in YYYY-MM-DD format. Optional.
	// Example: 2000-01-20
	StartDate param.Opt[string] `query:"startDate,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [CampaignMetricGetAttributionMetricsParams]'s query
// parameters as `url.Values`.
func (r CampaignMetricGetAttributionMetricsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type CampaignMetricGetRevenueAttributionParams struct {
	// The revenue attribution model used to calculate deal revenue credit. Defaults to
	// LINEAR if not specified. Enum values: LINEAR, FIRST_INTERACTION,
	// LAST_INTERACTION, FULL_PATH, U_SHAPED, W_SHAPED, TIME_DECAY, J_SHAPED,
	// INVERSE_J_SHAPED
	AttributionModel param.Opt[string] `query:"attributionModel,omitzero" json:"-"`
	// End date to fetch attribution data, YYYY-MM-DD
	EndDate param.Opt[string] `query:"endDate,omitzero" json:"-"`
	// Start date to fetch attribution data, YYYY-MM-DD
	StartDate param.Opt[string] `query:"startDate,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [CampaignMetricGetRevenueAttributionParams]'s query
// parameters as `url.Values`.
func (r CampaignMetricGetRevenueAttributionParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type CampaignMetricListContactIDsByTypeParams struct {
	CampaignGuid string `path:"campaignGuid" api:"required" json:"-"`
	// The paging cursor token of the last successfully read resource will be returned
	// as the `paging.next.after` JSON property of a paged response containing more
	// results.
	After param.Opt[string] `query:"after,omitzero" json:"-"`
	// The end date for fetching contact data, in YYYY-MM-DD format.
	// Optional. Example: 2000-01-27
	EndDate param.Opt[string] `query:"endDate,omitzero" json:"-"`
	// The maximum number of results to display per page.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// The start date for fetching contact data, in YYYY-MM-DD format.
	// Optional. Example: 2000-01-20
	StartDate param.Opt[string] `query:"startDate,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [CampaignMetricListContactIDsByTypeParams]'s query
// parameters as `url.Values`.
func (r CampaignMetricListContactIDsByTypeParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
