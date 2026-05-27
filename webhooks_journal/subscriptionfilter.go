// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package webhooks_journal

import (
	"context"
	"fmt"
	"net/http"
	"slices"

	"github.com/HubSpot/hubspot-sdk-go/internal/apijson"
	shimjson "github.com/HubSpot/hubspot-sdk-go/internal/encoding/json"
	"github.com/HubSpot/hubspot-sdk-go/internal/requestconfig"
	"github.com/HubSpot/hubspot-sdk-go/option"
	"github.com/HubSpot/hubspot-sdk-go/shared"
)

// SubscriptionFilterService contains methods and other services that help with
// interacting with the hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSubscriptionFilterService] method instead.
type SubscriptionFilterService struct {
	options []option.RequestOption
}

// NewSubscriptionFilterService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewSubscriptionFilterService(opts ...option.RequestOption) (r SubscriptionFilterService) {
	r = SubscriptionFilterService{}
	r.options = opts
	return
}

// Create a new filter for a specific webhook subscription in the HubSpot account.
// This endpoint allows you to define conditions that determine when a webhook
// should be triggered. The filter is associated with a subscription identified by
// its ID, and the request must include the filter details.
func (r *SubscriptionFilterService) New(ctx context.Context, body SubscriptionFilterNewParams, opts ...option.RequestOption) (res *shared.FilterCreateResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "webhooks-journal/subscriptions/2026-03/filters"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Retrieve the filters associated with a specific webhook subscription. This
// endpoint allows you to view the filters applied to a subscription, which can
// help in managing and understanding the conditions set for webhook events.
func (r *SubscriptionFilterService) List(ctx context.Context, subscriptionID int64, opts ...option.RequestOption) (res *[]shared.FilterResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := fmt.Sprintf("webhooks-journal/subscriptions/2026-03/filters/subscription/%v", subscriptionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Remove a specific filter from the webhooks journal subscriptions. This operation
// is useful for managing and cleaning up filters that are no longer needed. Once
// deleted, the filter cannot be recovered.
func (r *SubscriptionFilterService) Delete(ctx context.Context, filterID int64, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := fmt.Sprintf("webhooks-journal/subscriptions/2026-03/filters/%v", filterID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Retrieve a specific filter associated with a webhook journal subscription. This
// endpoint allows you to access the details of the filter identified by the
// filterId, which is useful for managing and understanding the conditions applied
// to webhook events.
func (r *SubscriptionFilterService) Get(ctx context.Context, filterID int64, opts ...option.RequestOption) (res *shared.FilterResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := fmt.Sprintf("webhooks-journal/subscriptions/2026-03/filters/%v", filterID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type SubscriptionFilterNewParams struct {
	FilterCreateRequest shared.FilterCreateRequestParam
	paramObj
}

func (r SubscriptionFilterNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.FilterCreateRequest)
}
func (r *SubscriptionFilterNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
