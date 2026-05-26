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

// SubscriptionService contains methods and other services that help with
// interacting with the hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSubscriptionService] method instead.
type SubscriptionService struct {
	options []option.RequestOption
	Filters SubscriptionFilterService
}

// NewSubscriptionService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewSubscriptionService(opts ...option.RequestOption) (r SubscriptionService) {
	r = SubscriptionService{}
	r.options = opts
	r.Filters = NewSubscriptionFilterService(opts...)
	return
}

// Create a new subscription in the Webhooks Journal for the specified version.
// This endpoint allows you to define the subscription details by providing the
// necessary information in the request body. It supports various types of
// subscriptions, including object, association, event, app lifecycle event, list
// membership, and GDPR privacy deletion. Ensure that all required fields are
// included in the request to successfully create a subscription.
func (r *SubscriptionService) New(ctx context.Context, body SubscriptionNewParams, opts ...option.RequestOption) (res *JournalSubscriptionResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "webhooks-journal/subscriptions/2026-03"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Retrieve a list of webhook journal subscriptions for the specified version. This
// endpoint allows you to view all active subscriptions without pagination. It is
// useful for monitoring and managing webhook subscriptions in your HubSpot
// account.
func (r *SubscriptionService) List(ctx context.Context, opts ...option.RequestOption) (res *JournalCollectionResponseSubscriptionResponseNoPaging, err error) {
	opts = slices.Concat(r.options, opts)
	path := "webhooks-journal/subscriptions/2026-03"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Delete a specific webhook journal subscription using its unique identifier. This
// operation is useful for managing and cleaning up subscriptions that are no
// longer needed in your HubSpot account.
func (r *SubscriptionService) Delete(ctx context.Context, subscriptionID int64, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := fmt.Sprintf("webhooks-journal/subscriptions/2026-03/%v", subscriptionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Delete a webhook journal subscription for a specific portal. This operation
// removes the subscription associated with the given portalId, ensuring that no
// further webhook events are sent for this portal. Use this endpoint to manage and
// clean up subscriptions that are no longer needed.
func (r *SubscriptionService) DeleteForPortal(ctx context.Context, portalID int64, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := fmt.Sprintf("webhooks-journal/subscriptions/2026-03/portals/%v", portalID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Retrieve details of a specific webhook subscription using its unique identifier.
// This endpoint is useful for obtaining information about a particular
// subscription, such as its actions, object type, and associated properties.
func (r *SubscriptionService) Get(ctx context.Context, subscriptionID int64, opts ...option.RequestOption) (res *JournalSubscriptionResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := fmt.Sprintf("webhooks-journal/subscriptions/2026-03/%v", subscriptionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type SubscriptionNewParams struct {
	SubscriptionUpsertRequest shared.SubscriptionUpsertRequestUnionParam
	paramObj
}

func (r SubscriptionNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.SubscriptionUpsertRequest)
}
func (r *SubscriptionNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
