// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package communication_preferences

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

// StatusBatchService contains methods and other services that help with
// interacting with the hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewStatusBatchService] method instead.
type StatusBatchService struct {
	options []option.RequestOption
}

// NewStatusBatchService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewStatusBatchService(opts ...option.RequestOption) (r StatusBatchService) {
	r = StatusBatchService{}
	r.options = opts
	return
}

// Checks whether a set of contacts have opted out of all communications.
func (r *StatusBatchService) GetUnsubscribeAllStatuses(ctx context.Context, params StatusBatchGetUnsubscribeAllStatusesParams, opts ...option.RequestOption) (res *BatchResponsePublicWideStatusBulkResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "communication-preferences/2026-03/statuses/batch/unsubscribe-all/read"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Batch retrieve subscription statuses for a set of contacts.
func (r *StatusBatchService) Read(ctx context.Context, params StatusBatchReadParams, opts ...option.RequestOption) (res *BatchResponsePublicStatusBulkResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "communication-preferences/2026-03/statuses/batch/read"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Unsubscribe a set of contacts from all email subscriptions.
func (r *StatusBatchService) UnsubscribeAll(ctx context.Context, params StatusBatchUnsubscribeAllParams, opts ...option.RequestOption) (res *BatchResponsePublicBulkOptOutFromAllResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "communication-preferences/2026-03/statuses/batch/unsubscribe-all"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Update the subscription status for a set of contacts.
func (r *StatusBatchService) UpdateStatuses(ctx context.Context, body StatusBatchUpdateStatusesParams, opts ...option.RequestOption) (res *BatchResponsePublicStatus, err error) {
	opts = slices.Concat(r.options, opts)
	path := "communication-preferences/2026-03/statuses/batch/write"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type StatusBatchGetUnsubscribeAllStatusesParams struct {
	// The communication channel to filter the unsubscribe statuses. This parameter is
	// required and currently supports 'EMAIL' as a valid value.
	//
	// Any of "EMAIL".
	Channel          StatusBatchGetUnsubscribeAllStatusesParamsChannel `query:"channel,omitzero" api:"required" json:"-"`
	BatchInputString shared.BatchInputStringParam
	// The ID of the business unit to filter the results. This is an optional
	// parameter.
	BusinessUnitID param.Opt[int64] `query:"businessUnitId,omitzero" json:"-"`
	paramObj
}

func (r StatusBatchGetUnsubscribeAllStatusesParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.BatchInputString)
}
func (r *StatusBatchGetUnsubscribeAllStatusesParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [StatusBatchGetUnsubscribeAllStatusesParams]'s query
// parameters as `url.Values`.
func (r StatusBatchGetUnsubscribeAllStatusesParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// The communication channel to filter the unsubscribe statuses. This parameter is
// required and currently supports 'EMAIL' as a valid value.
type StatusBatchGetUnsubscribeAllStatusesParamsChannel string

const (
	StatusBatchGetUnsubscribeAllStatusesParamsChannelEmail StatusBatchGetUnsubscribeAllStatusesParamsChannel = "EMAIL"
)

type StatusBatchReadParams struct {
	// The communication channel to filter the subscription statuses. Must be 'EMAIL'.
	//
	// Any of "EMAIL".
	Channel          StatusBatchReadParamsChannel `query:"channel,omitzero" api:"required" json:"-"`
	BatchInputString shared.BatchInputStringParam
	// An optional integer representing the business unit ID. This parameter helps to
	// filter the results based on the specific business unit.
	BusinessUnitID param.Opt[int64] `query:"businessUnitId,omitzero" json:"-"`
	paramObj
}

func (r StatusBatchReadParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.BatchInputString)
}
func (r *StatusBatchReadParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [StatusBatchReadParams]'s query parameters as `url.Values`.
func (r StatusBatchReadParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// The communication channel to filter the subscription statuses. Must be 'EMAIL'.
type StatusBatchReadParamsChannel string

const (
	StatusBatchReadParamsChannelEmail StatusBatchReadParamsChannel = "EMAIL"
)

type StatusBatchUnsubscribeAllParams struct {
	// The communication channel from which subscribers will be unsubscribed. This
	// parameter is required and currently supports only 'EMAIL'.
	//
	// Any of "EMAIL".
	Channel          StatusBatchUnsubscribeAllParamsChannel `query:"channel,omitzero" api:"required" json:"-"`
	BatchInputString shared.BatchInputStringParam
	// An optional integer representing the business unit ID for which the operation is
	// being performed.
	BusinessUnitID param.Opt[int64] `query:"businessUnitId,omitzero" json:"-"`
	// A boolean indicating whether to include detailed information in the response.
	// Defaults to false.
	Verbose param.Opt[bool] `query:"verbose,omitzero" json:"-"`
	paramObj
}

func (r StatusBatchUnsubscribeAllParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.BatchInputString)
}
func (r *StatusBatchUnsubscribeAllParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [StatusBatchUnsubscribeAllParams]'s query parameters as
// `url.Values`.
func (r StatusBatchUnsubscribeAllParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// The communication channel from which subscribers will be unsubscribed. This
// parameter is required and currently supports only 'EMAIL'.
type StatusBatchUnsubscribeAllParamsChannel string

const (
	StatusBatchUnsubscribeAllParamsChannelEmail StatusBatchUnsubscribeAllParamsChannel = "EMAIL"
)

type StatusBatchUpdateStatusesParams struct {
	BatchInputPublicStatusRequest BatchInputPublicStatusRequestParam
	paramObj
}

func (r StatusBatchUpdateStatusesParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.BatchInputPublicStatusRequest)
}
func (r *StatusBatchUpdateStatusesParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
