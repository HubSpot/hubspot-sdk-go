// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package marketing

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
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

// SubscriptionV4StatusService contains methods and other services that help with
// interacting with the Hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSubscriptionV4StatusService] method instead.
type SubscriptionV4StatusService struct {
	Options []option.RequestOption
}

// NewSubscriptionV4StatusService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewSubscriptionV4StatusService(opts ...option.RequestOption) (r SubscriptionV4StatusService) {
	r = SubscriptionV4StatusService{}
	r.Options = opts
	return
}

// Set the subscription status of a specific contact.
func (r *SubscriptionV4StatusService) Update(ctx context.Context, subscriberIDString string, body SubscriptionV4StatusUpdateParams, opts ...option.RequestOption) (res *ActionResponseWithResultsPublicStatus, err error) {
	opts = slices.Concat(r.Options, opts)
	if subscriberIDString == "" {
		err = errors.New("missing required subscriberIdString parameter")
		return
	}
	path := fmt.Sprintf("communication-preferences/v4/statuses/%s", subscriberIDString)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Batch retrieve subscription statuses for a set of contacts.
func (r *SubscriptionV4StatusService) BatchGet(ctx context.Context, params SubscriptionV4StatusBatchGetParams, opts ...option.RequestOption) (res *BatchResponsePublicStatusBulkResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "communication-preferences/v4/statuses/batch/read"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Checks whether a set of contacts have opted out of all communications.
func (r *SubscriptionV4StatusService) BatchGetUnsubscribeAllStatus(ctx context.Context, params SubscriptionV4StatusBatchGetUnsubscribeAllStatusParams, opts ...option.RequestOption) (res *BatchResponsePublicWideStatusBulkResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "communication-preferences/v4/statuses/batch/unsubscribe-all/read"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Unsubscribe a set of contacts from all email subscriptions.
func (r *SubscriptionV4StatusService) BatchUnsubscribeAll(ctx context.Context, params SubscriptionV4StatusBatchUnsubscribeAllParams, opts ...option.RequestOption) (res *BatchResponsePublicBulkOptOutFromAllResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "communication-preferences/v4/statuses/batch/unsubscribe-all"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Update the subscription status for a set of contacts.
func (r *SubscriptionV4StatusService) BatchUpdate(ctx context.Context, body SubscriptionV4StatusBatchUpdateParams, opts ...option.RequestOption) (res *BatchResponsePublicStatus, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "communication-preferences/v4/statuses/batch/write"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieve a contact's current email subscription preferences.
func (r *SubscriptionV4StatusService) Get(ctx context.Context, subscriberIDString string, query SubscriptionV4StatusGetParams, opts ...option.RequestOption) (res *ActionResponseWithResultsPublicStatus, err error) {
	opts = slices.Concat(r.Options, opts)
	if subscriberIDString == "" {
		err = errors.New("missing required subscriberIdString parameter")
		return
	}
	path := fmt.Sprintf("communication-preferences/v4/statuses/%s", subscriberIDString)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Check whether a contact has unsubscribed from all email subscriptions. If a
// contact has not opted out of all communications, the response `results` array
// will be empty.
func (r *SubscriptionV4StatusService) GetUnsubscribeAllStatus(ctx context.Context, subscriberIDString string, query SubscriptionV4StatusGetUnsubscribeAllStatusParams, opts ...option.RequestOption) (res *ActionResponseWithResultsPublicWideStatus, err error) {
	opts = slices.Concat(r.Options, opts)
	if subscriberIDString == "" {
		err = errors.New("missing required subscriberIdString parameter")
		return
	}
	path := fmt.Sprintf("communication-preferences/v4/statuses/%s/unsubscribe-all", subscriberIDString)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Unsubscribe a contact from all email subscriptions.
func (r *SubscriptionV4StatusService) UnsubscribeAll(ctx context.Context, subscriberIDString string, body SubscriptionV4StatusUnsubscribeAllParams, opts ...option.RequestOption) (res *ActionResponseWithResultsPublicStatus, err error) {
	opts = slices.Concat(r.Options, opts)
	if subscriberIDString == "" {
		err = errors.New("missing required subscriberIdString parameter")
		return
	}
	path := fmt.Sprintf("communication-preferences/v4/statuses/%s/unsubscribe-all", subscriberIDString)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type SubscriptionV4StatusUpdateParams struct {
	PartialPublicStatusRequest PartialPublicStatusRequestParam
	paramObj
}

func (r SubscriptionV4StatusUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.PartialPublicStatusRequest)
}
func (r *SubscriptionV4StatusUpdateParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.PartialPublicStatusRequest)
}

type SubscriptionV4StatusBatchGetParams struct {
	// The channel type for the subscription type. Currently, the only supported
	// channel type is `EMAIL`.
	//
	// Any of "EMAIL".
	Channel SubscriptionV4StatusBatchGetParamsChannel `query:"channel,omitzero,required" json:"-"`
	// Wrapper for providing an array of strings as inputs.
	BatchInputString shared.BatchInputStringParam
	// If you have the
	// [business unit add-on](https://developers.hubspot.com/beta-docs/guides/api/settings/business-units-api),
	// include this parameter to filter results by business unit ID. The default
	// Account business unit will always use `0`.
	BusinessUnitID param.Opt[int64] `query:"businessUnitId,omitzero" json:"-"`
	paramObj
}

func (r SubscriptionV4StatusBatchGetParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.BatchInputString)
}
func (r *SubscriptionV4StatusBatchGetParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.BatchInputString)
}

// URLQuery serializes [SubscriptionV4StatusBatchGetParams]'s query parameters as
// `url.Values`.
func (r SubscriptionV4StatusBatchGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// The channel type for the subscription type. Currently, the only supported
// channel type is `EMAIL`.
type SubscriptionV4StatusBatchGetParamsChannel string

const (
	SubscriptionV4StatusBatchGetParamsChannelEmail SubscriptionV4StatusBatchGetParamsChannel = "EMAIL"
)

type SubscriptionV4StatusBatchGetUnsubscribeAllStatusParams struct {
	// The channel type for the subscription type. Currently, the only supported
	// channel type is `EMAIL`.
	//
	// Any of "EMAIL".
	Channel SubscriptionV4StatusBatchGetUnsubscribeAllStatusParamsChannel `query:"channel,omitzero,required" json:"-"`
	// Wrapper for providing an array of strings as inputs.
	BatchInputString shared.BatchInputStringParam
	// If you have the
	// [business unit add-on](https://developers.hubspot.com/beta-docs/guides/api/settings/business-units-api),
	// include this parameter to filter results by business unit ID. The default
	// Account business unit will always use `0`.
	BusinessUnitID param.Opt[int64] `query:"businessUnitId,omitzero" json:"-"`
	paramObj
}

func (r SubscriptionV4StatusBatchGetUnsubscribeAllStatusParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.BatchInputString)
}
func (r *SubscriptionV4StatusBatchGetUnsubscribeAllStatusParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.BatchInputString)
}

// URLQuery serializes [SubscriptionV4StatusBatchGetUnsubscribeAllStatusParams]'s
// query parameters as `url.Values`.
func (r SubscriptionV4StatusBatchGetUnsubscribeAllStatusParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// The channel type for the subscription type. Currently, the only supported
// channel type is `EMAIL`.
type SubscriptionV4StatusBatchGetUnsubscribeAllStatusParamsChannel string

const (
	SubscriptionV4StatusBatchGetUnsubscribeAllStatusParamsChannelEmail SubscriptionV4StatusBatchGetUnsubscribeAllStatusParamsChannel = "EMAIL"
)

type SubscriptionV4StatusBatchUnsubscribeAllParams struct {
	// The channel type for the subscription type. Currently, the only supported
	// channel type is `EMAIL`.
	//
	// Any of "EMAIL".
	Channel SubscriptionV4StatusBatchUnsubscribeAllParamsChannel `query:"channel,omitzero,required" json:"-"`
	// Wrapper for providing an array of strings as inputs.
	BatchInputString shared.BatchInputStringParam
	// If you have the
	// [business unit add-on](https://developers.hubspot.com/beta-docs/guides/api/settings/business-units-api),
	// include this parameter to filter results by business unit ID. The default
	// Account business unit will always use `0`.
	BusinessUnitID param.Opt[int64] `query:"businessUnitId,omitzero" json:"-"`
	// Set to `true` to include the details of the updated subscription statuses in the
	// response. Not including this parameter will result in an empty response.
	Verbose param.Opt[bool] `query:"verbose,omitzero" json:"-"`
	paramObj
}

func (r SubscriptionV4StatusBatchUnsubscribeAllParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.BatchInputString)
}
func (r *SubscriptionV4StatusBatchUnsubscribeAllParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.BatchInputString)
}

// URLQuery serializes [SubscriptionV4StatusBatchUnsubscribeAllParams]'s query
// parameters as `url.Values`.
func (r SubscriptionV4StatusBatchUnsubscribeAllParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// The channel type for the subscription type. Currently, the only supported
// channel type is `EMAIL`.
type SubscriptionV4StatusBatchUnsubscribeAllParamsChannel string

const (
	SubscriptionV4StatusBatchUnsubscribeAllParamsChannelEmail SubscriptionV4StatusBatchUnsubscribeAllParamsChannel = "EMAIL"
)

type SubscriptionV4StatusBatchUpdateParams struct {
	BatchInputPublicStatusRequest BatchInputPublicStatusRequestParam
	paramObj
}

func (r SubscriptionV4StatusBatchUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.BatchInputPublicStatusRequest)
}
func (r *SubscriptionV4StatusBatchUpdateParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.BatchInputPublicStatusRequest)
}

type SubscriptionV4StatusGetParams struct {
	// The channel type for the subscription type. Currently, the only supported
	// channel type is `EMAIL`.
	//
	// Any of "EMAIL".
	Channel SubscriptionV4StatusGetParamsChannel `query:"channel,omitzero,required" json:"-"`
	// If you have the
	// [business unit add-on](https://developers.hubspot.com/beta-docs/guides/api/settings/business-units-api),
	// include this parameter to filter results by business unit ID. The default
	// Account business unit will always use `0`.
	BusinessUnitID param.Opt[int64] `query:"businessUnitId,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SubscriptionV4StatusGetParams]'s query parameters as
// `url.Values`.
func (r SubscriptionV4StatusGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// The channel type for the subscription type. Currently, the only supported
// channel type is `EMAIL`.
type SubscriptionV4StatusGetParamsChannel string

const (
	SubscriptionV4StatusGetParamsChannelEmail SubscriptionV4StatusGetParamsChannel = "EMAIL"
)

type SubscriptionV4StatusGetUnsubscribeAllStatusParams struct {
	// The channel type for the subscription type. Currently, the only supported
	// channel type is `EMAIL`.
	//
	// Any of "EMAIL".
	Channel SubscriptionV4StatusGetUnsubscribeAllStatusParamsChannel `query:"channel,omitzero,required" json:"-"`
	// If you have the
	// [business unit add-on](https://developers.hubspot.com/beta-docs/guides/api/settings/business-units-api),
	// include this parameter to filter results by business unit ID. The default
	// Account business unit will always use `0`.
	BusinessUnitID param.Opt[int64] `query:"businessUnitId,omitzero" json:"-"`
	// Set to `true` to include the details of the updated subscription statuses in the
	// response. Not including this parameter will result in an empty response.
	Verbose param.Opt[bool] `query:"verbose,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SubscriptionV4StatusGetUnsubscribeAllStatusParams]'s query
// parameters as `url.Values`.
func (r SubscriptionV4StatusGetUnsubscribeAllStatusParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// The channel type for the subscription type. Currently, the only supported
// channel type is `EMAIL`.
type SubscriptionV4StatusGetUnsubscribeAllStatusParamsChannel string

const (
	SubscriptionV4StatusGetUnsubscribeAllStatusParamsChannelEmail SubscriptionV4StatusGetUnsubscribeAllStatusParamsChannel = "EMAIL"
)

type SubscriptionV4StatusUnsubscribeAllParams struct {
	// The channel type for the subscription type. Currently, the only supported
	// channel type is `EMAIL`.
	//
	// Any of "EMAIL".
	Channel SubscriptionV4StatusUnsubscribeAllParamsChannel `query:"channel,omitzero,required" json:"-"`
	// If you have the
	// [business unit add-on](https://developers.hubspot.com/beta-docs/guides/api/settings/business-units-api),
	// include this parameter to filter results by business unit ID. The default
	// Account business unit will always use `0`.
	BusinessUnitID param.Opt[int64] `query:"businessUnitId,omitzero" json:"-"`
	// Set to `true` to include the details of the updated subscription statuses in the
	// response. Not including this parameter will result in an empty response.
	Verbose param.Opt[bool] `query:"verbose,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SubscriptionV4StatusUnsubscribeAllParams]'s query
// parameters as `url.Values`.
func (r SubscriptionV4StatusUnsubscribeAllParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// The channel type for the subscription type. Currently, the only supported
// channel type is `EMAIL`.
type SubscriptionV4StatusUnsubscribeAllParamsChannel string

const (
	SubscriptionV4StatusUnsubscribeAllParamsChannelEmail SubscriptionV4StatusUnsubscribeAllParamsChannel = "EMAIL"
)
