// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package crm

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"slices"

	"github.com/stainless-sdks/hubspot-sdk-go/internal/apijson"
	shimjson "github.com/stainless-sdks/hubspot-sdk-go/internal/encoding/json"
	"github.com/stainless-sdks/hubspot-sdk-go/internal/requestconfig"
	"github.com/stainless-sdks/hubspot-sdk-go/option"
	"github.com/stainless-sdks/hubspot-sdk-go/packages/param"
)

// SubscriptionService contains methods and other services that help with
// interacting with the hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSubscriptionService] method instead.
type SubscriptionService struct {
	Options []option.RequestOption
}

// NewSubscriptionService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewSubscriptionService(opts ...option.RequestOption) (r SubscriptionService) {
	r = SubscriptionService{}
	r.Options = opts
	return
}

// Cancel an active commerce subscription using the subscription ID.
func (r *SubscriptionService) Cancel(ctx context.Context, objectID int64, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := fmt.Sprintf("payments-subscriptions/v1/subscriptions/crm/%v/cancel", objectID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Pause an active subscription using the subscription ID.
func (r *SubscriptionService) Pause(ctx context.Context, objectID int64, body SubscriptionPauseParams, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := fmt.Sprintf("payments-subscriptions/v1/subscriptions/crm/%v/pause", objectID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Resume a previously paused subscription using the subscription ID.
func (r *SubscriptionService) Unpause(ctx context.Context, objectID int64, body SubscriptionUnpauseParams, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := fmt.Sprintf("payments-subscriptions/v1/subscriptions/crm/%v/unpause", objectID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type PauseSubscriptionRequestParam struct {
	PauseReason param.Opt[string] `json:"pauseReason,omitzero"`
	paramObj
}

func (r PauseSubscriptionRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow PauseSubscriptionRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PauseSubscriptionRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property ProposedNextBillingDate is required.
type UnpauseRequestParam struct {
	ProposedNextBillingDate int64 `json:"proposedNextBillingDate,required"`
	paramObj
}

func (r UnpauseRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow UnpauseRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UnpauseRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SubscriptionPauseParams struct {
	PauseSubscriptionRequest PauseSubscriptionRequestParam
	paramObj
}

func (r SubscriptionPauseParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.PauseSubscriptionRequest)
}
func (r *SubscriptionPauseParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.PauseSubscriptionRequest)
}

type SubscriptionUnpauseParams struct {
	UnpauseRequest UnpauseRequestParam
	paramObj
}

func (r SubscriptionUnpauseParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.UnpauseRequest)
}
func (r *SubscriptionUnpauseParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.UnpauseRequest)
}
