// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package marketing

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/stainless-sdks/hubspot-sdk-go/internal/apiquery"
	"github.com/stainless-sdks/hubspot-sdk-go/internal/requestconfig"
	"github.com/stainless-sdks/hubspot-sdk-go/option"
	"github.com/stainless-sdks/hubspot-sdk-go/packages/param"
)

// CampaignReportService contains methods and other services that help with
// interacting with the Hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCampaignReportService] method instead.
type CampaignReportService struct {
	Options []option.RequestOption
}

// NewCampaignReportService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCampaignReportService(opts ...option.RequestOption) (r CampaignReportService) {
	r = CampaignReportService{}
	r.Options = opts
	return
}

// This endpoint retrieves key attribution metrics for a specified campaign, such
// as sessions, new contacts, and influenced contacts.
func (r *CampaignReportService) GetAttributionMetrics(ctx context.Context, campaignGuid string, query CampaignReportGetAttributionMetricsParams, opts ...option.RequestOption) (res *MetricsCounters, err error) {
	opts = slices.Concat(r.Options, opts)
	if campaignGuid == "" {
		err = errors.New("missing required campaignGuid parameter")
		return
	}
	path := fmt.Sprintf("marketing/v3/campaigns/%s/reports/metrics", campaignGuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Fetch revenue attribution report data for a specified campaign
func (r *CampaignReportService) GetRevenueAttribution(ctx context.Context, campaignGuid string, query CampaignReportGetRevenueAttributionParams, opts ...option.RequestOption) (res *RevenueAttributionAggregate, err error) {
	opts = slices.Concat(r.Options, opts)
	if campaignGuid == "" {
		err = errors.New("missing required campaignGuid parameter")
		return
	}
	path := fmt.Sprintf("marketing/v3/campaigns/%s/reports/revenue", campaignGuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Fetch the list of contact IDs for the specified campaign and contact type
func (r *CampaignReportService) ListContactIDsByType(ctx context.Context, contactType string, params CampaignReportListContactIDsByTypeParams, opts ...option.RequestOption) (res *CollectionResponseContactReferenceForwardPaging, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.CampaignGuid == "" {
		err = errors.New("missing required campaignGuid parameter")
		return
	}
	if contactType == "" {
		err = errors.New("missing required contactType parameter")
		return
	}
	path := fmt.Sprintf("marketing/v3/campaigns/%s/reports/contacts/%s", params.CampaignGuid, contactType)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return
}

type CampaignReportGetAttributionMetricsParams struct {
	// End date for the report data, formatted as YYYY-MM-DD. Default value: Current
	// date
	EndDate param.Opt[string] `query:"endDate,omitzero" json:"-"`
	// The start date for the report data, formatted as YYYY-MM-DD. Default value:
	// 2006-01-01
	StartDate param.Opt[string] `query:"startDate,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [CampaignReportGetAttributionMetricsParams]'s query
// parameters as `url.Values`.
func (r CampaignReportGetAttributionMetricsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type CampaignReportGetRevenueAttributionParams struct {
	// Allowed values: LINEAR, FIRST_INTERACTION, LAST_INTERACTION, FULL_PATH,
	// U_SHAPED, W_SHAPED, TIME_DECAY, J_SHAPED, INVERSE_J_SHAPED Default value: LINEAR
	AttributionModel param.Opt[string] `query:"attributionModel,omitzero" json:"-"`
	// End date for the report data, formatted as YYYY-MM-DD. Default value: Current
	// date
	EndDate param.Opt[string] `query:"endDate,omitzero" json:"-"`
	// The start date for the report data, formatted as YYYY-MM-DD. Default value:
	// 2006-01-01
	StartDate param.Opt[string] `query:"startDate,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [CampaignReportGetRevenueAttributionParams]'s query
// parameters as `url.Values`.
func (r CampaignReportGetRevenueAttributionParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type CampaignReportListContactIDsByTypeParams struct {
	CampaignGuid string `path:"campaignGuid,required" json:"-"`
	// A cursor for pagination. If provided, the results will start after the given
	// cursor. Example: NTI1Cg%3D%3D
	After param.Opt[string] `query:"after,omitzero" json:"-"`
	// End date for the report data, formatted as YYYY-MM-DD. Default value: Current
	// date
	EndDate param.Opt[string] `query:"endDate,omitzero" json:"-"`
	// Limit for the number of contacts to fetch Default: 100
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// The start date for the report data, formatted as YYYY-MM-DD. Default value:
	// 2006-01-01
	StartDate param.Opt[string] `query:"startDate,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [CampaignReportListContactIDsByTypeParams]'s query
// parameters as `url.Values`.
func (r CampaignReportListContactIDsByTypeParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
