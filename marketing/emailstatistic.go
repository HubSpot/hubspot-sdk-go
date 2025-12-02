// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package marketing

import (
	"context"
	"net/http"
	"net/url"
	"slices"

	"github.com/stainless-sdks/hubspot-sdk-go/internal/apiquery"
	"github.com/stainless-sdks/hubspot-sdk-go/internal/requestconfig"
	"github.com/stainless-sdks/hubspot-sdk-go/option"
	"github.com/stainless-sdks/hubspot-sdk-go/packages/param"
)

// EmailStatisticService contains methods and other services that help with
// interacting with the hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEmailStatisticService] method instead.
type EmailStatisticService struct {
	Options []option.RequestOption
}

// NewEmailStatisticService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewEmailStatisticService(opts ...option.RequestOption) (r EmailStatisticService) {
	r = EmailStatisticService{}
	r.Options = opts
	return
}

// Use this endpoint to get aggregated statistics of emails sent in a specified
// time span. It also returns the list of emails that were sent during the time
// span.
func (r *EmailStatisticService) Get(ctx context.Context, query EmailStatisticGetParams, opts ...option.RequestOption) (res *AggregateEmailStatistics, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "marketing/v3/emails/statistics/list"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Get aggregated statistics in intervals for a specified time span. Each interval
// contains aggregated statistics of the emails that were sent in that time.
func (r *EmailStatisticService) GetHistogram(ctx context.Context, query EmailStatisticGetHistogramParams, opts ...option.RequestOption) (res *CollectionResponseWithTotalEmailStatisticIntervalNoPaging, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "marketing/v3/emails/statistics/histogram"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type EmailStatisticGetParams struct {
	// The end timestamp of the time span, in ISO8601 representation.
	EndTimestamp param.Opt[string] `query:"endTimestamp,omitzero" json:"-"`
	// Specifies which email properties should be returned. All properties will be
	// returned by default.
	Property param.Opt[string] `query:"property,omitzero" json:"-"`
	// The start timestamp of the time span, in ISO8601 representation.
	StartTimestamp param.Opt[string] `query:"startTimestamp,omitzero" json:"-"`
	// Filter by email IDs. Only include statistics of emails with these IDs.
	EmailIDs []int64 `query:"emailIds,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [EmailStatisticGetParams]'s query parameters as
// `url.Values`.
func (r EmailStatisticGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type EmailStatisticGetHistogramParams struct {
	// The end timestamp of the time span, in ISO8601 representation.
	EndTimestamp param.Opt[string] `query:"endTimestamp,omitzero" json:"-"`
	// The start timestamp of the time span, in ISO8601 representation.
	StartTimestamp param.Opt[string] `query:"startTimestamp,omitzero" json:"-"`
	// Filter by email IDs. Only include statistics of emails with these IDs.
	EmailIDs []int64 `query:"emailIds,omitzero" json:"-"`
	// The interval to aggregate statistics for.
	//
	// Any of "DAY", "HOUR", "MINUTE", "MONTH", "QUARTER", "QUARTER_HOUR", "SECOND",
	// "WEEK", "YEAR".
	Interval EmailStatisticGetHistogramParamsInterval `query:"interval,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [EmailStatisticGetHistogramParams]'s query parameters as
// `url.Values`.
func (r EmailStatisticGetHistogramParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// The interval to aggregate statistics for.
type EmailStatisticGetHistogramParamsInterval string

const (
	EmailStatisticGetHistogramParamsIntervalDay         EmailStatisticGetHistogramParamsInterval = "DAY"
	EmailStatisticGetHistogramParamsIntervalHour        EmailStatisticGetHistogramParamsInterval = "HOUR"
	EmailStatisticGetHistogramParamsIntervalMinute      EmailStatisticGetHistogramParamsInterval = "MINUTE"
	EmailStatisticGetHistogramParamsIntervalMonth       EmailStatisticGetHistogramParamsInterval = "MONTH"
	EmailStatisticGetHistogramParamsIntervalQuarter     EmailStatisticGetHistogramParamsInterval = "QUARTER"
	EmailStatisticGetHistogramParamsIntervalQuarterHour EmailStatisticGetHistogramParamsInterval = "QUARTER_HOUR"
	EmailStatisticGetHistogramParamsIntervalSecond      EmailStatisticGetHistogramParamsInterval = "SECOND"
	EmailStatisticGetHistogramParamsIntervalWeek        EmailStatisticGetHistogramParamsInterval = "WEEK"
	EmailStatisticGetHistogramParamsIntervalYear        EmailStatisticGetHistogramParamsInterval = "YEAR"
)
