// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package webhooks

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"slices"

	shimjson "github.com/stainless-sdks/hubspot-sdk-go/internal/encoding/json"
	"github.com/stainless-sdks/hubspot-sdk-go/internal/requestconfig"
	"github.com/stainless-sdks/hubspot-sdk-go/option"
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

// Create new event subscription for the specified app.
func (r *SubscriptionService) New(ctx context.Context, appID int64, body SubscriptionNewParams, opts ...option.RequestOption) (res *SubscriptionResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("webhooks/v3/%v/subscriptions", appID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Update an existing event subscription by ID.
func (r *SubscriptionService) Update(ctx context.Context, subscriptionID int64, params SubscriptionUpdateParams, opts ...option.RequestOption) (res *SubscriptionResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("webhooks/v3/%v/subscriptions/%v", params.AppID, subscriptionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return
}

// Retrieve event subscriptions for the specified app.
func (r *SubscriptionService) List(ctx context.Context, appID int64, opts ...option.RequestOption) (res *SubscriptionListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("webhooks/v3/%v/subscriptions", appID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Delete an existing event subscription by ID.
func (r *SubscriptionService) Delete(ctx context.Context, subscriptionID int64, body SubscriptionDeleteParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := fmt.Sprintf("webhooks/v3/%v/subscriptions/%v", body.AppID, subscriptionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Retrieve a specific event subscription by ID.
func (r *SubscriptionService) Get(ctx context.Context, subscriptionID int64, query SubscriptionGetParams, opts ...option.RequestOption) (res *SubscriptionResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("webhooks/v3/%v/subscriptions/%v", query.AppID, subscriptionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Batch create event subscriptions for the specified app.
func (r *SubscriptionService) UpdateBatch(ctx context.Context, appID int64, body SubscriptionUpdateBatchParams, opts ...option.RequestOption) (res *BatchResponseSubscriptionResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("webhooks/v3/%v/subscriptions/batch/update", appID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type SubscriptionNewParams struct {
	// New webhook settings for an app.
	SubscriptionCreateRequest SubscriptionCreateRequestParam
	paramObj
}

func (r SubscriptionNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.SubscriptionCreateRequest)
}
func (r *SubscriptionNewParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.SubscriptionCreateRequest)
}

type SubscriptionUpdateParams struct {
	AppID int64 `path:"appId,required" json:"-"`
	// Updated details for the subscription.
	SubscriptionPatchRequest SubscriptionPatchRequestParam
	paramObj
}

func (r SubscriptionUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.SubscriptionPatchRequest)
}
func (r *SubscriptionUpdateParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.SubscriptionPatchRequest)
}

type SubscriptionDeleteParams struct {
	AppID int64 `path:"appId,required" json:"-"`
	paramObj
}

type SubscriptionGetParams struct {
	AppID int64 `path:"appId,required" json:"-"`
	paramObj
}

type SubscriptionUpdateBatchParams struct {
	BatchInputSubscriptionBatchUpdateRequest BatchInputSubscriptionBatchUpdateRequestParam
	paramObj
}

func (r SubscriptionUpdateBatchParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.BatchInputSubscriptionBatchUpdateRequest)
}
func (r *SubscriptionUpdateBatchParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.BatchInputSubscriptionBatchUpdateRequest)
}
