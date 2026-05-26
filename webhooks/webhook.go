// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package webhooks

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/HubSpot/hubspot-sdk-go/internal/apijson"
	"github.com/HubSpot/hubspot-sdk-go/internal/apiquery"
	shimjson "github.com/HubSpot/hubspot-sdk-go/internal/encoding/json"
	"github.com/HubSpot/hubspot-sdk-go/internal/requestconfig"
	"github.com/HubSpot/hubspot-sdk-go/option"
	"github.com/HubSpot/hubspot-sdk-go/packages/param"
	"github.com/HubSpot/hubspot-sdk-go/packages/respjson"
	"github.com/HubSpot/hubspot-sdk-go/shared"
	"github.com/HubSpot/hubspot-sdk-go/webhooks_journal"
)

// WebhookService contains methods and other services that help with interacting
// with the hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWebhookService] method instead.
type WebhookService struct {
	options []option.RequestOption
}

// NewWebhookService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewWebhookService(opts ...option.RequestOption) (r WebhookService) {
	r = WebhookService{}
	r.options = opts
	return
}

// Batch create event subscriptions for the specified app.
func (r *WebhookService) NewBatchEventSubscriptions(ctx context.Context, appID int64, body WebhookNewBatchEventSubscriptionsParams, opts ...option.RequestOption) (res *BatchResponseSubscriptionResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := fmt.Sprintf("webhooks/2026-03/%v/subscriptions/batch/update", appID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Create a batch of CRM object snapshots in HubSpot. This endpoint is used to
// capture the current state of specified CRM objects for later reference or
// analysis. It requires a JSON payload containing the details of the CRM objects
// to snapshot. This operation is exempt from daily and ten-secondly rate limits.
func (r *WebhookService) NewCrmSnapshots(ctx context.Context, body WebhookNewCrmSnapshotsParams, opts ...option.RequestOption) (res *shared.CrmObjectSnapshotBatchResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "webhooks-journal/snapshots/2026-03/crm"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Create new event subscription for the specified app.
func (r *WebhookService) NewEventSubscription(ctx context.Context, appID int64, body WebhookNewEventSubscriptionParams, opts ...option.RequestOption) (res *SubscriptionResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := fmt.Sprintf("webhooks/2026-03/%v/subscriptions", appID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Create a new subscription in the Webhooks Journal for the specified version.
// This endpoint allows you to define the subscription details by providing the
// necessary information in the request body. It supports various types of
// subscriptions, including object, association, event, app lifecycle event, list
// membership, and GDPR privacy deletion. Ensure that all required fields are
// included in the request to successfully create a subscription.
func (r *WebhookService) NewJournalSubscription(ctx context.Context, body WebhookNewJournalSubscriptionParams, opts ...option.RequestOption) (res *webhooks_journal.JournalSubscriptionResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "webhooks-journal/subscriptions/2026-03"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Create a new filter for a specific webhook subscription in the HubSpot account.
// This endpoint allows you to define conditions that determine when a webhook
// should be triggered. The filter is associated with a subscription identified by
// its ID, and the request must include the filter details.
func (r *WebhookService) NewSubscriptionFilter(ctx context.Context, body WebhookNewSubscriptionFilterParams, opts ...option.RequestOption) (res *shared.FilterCreateResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "webhooks-journal/subscriptions/2026-03/filters"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Delete an existing event subscription by ID.
func (r *WebhookService) DeleteEventSubscription(ctx context.Context, subscriptionID int64, body WebhookDeleteEventSubscriptionParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := fmt.Sprintf("webhooks/2026-03/%v/subscriptions/%v", body.AppID, subscriptionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Delete a specific webhook journal subscription using its unique identifier. This
// operation is useful for managing and cleaning up subscriptions that are no
// longer needed in your HubSpot account.
func (r *WebhookService) DeleteJournalSubscription(ctx context.Context, subscriptionID int64, opts ...option.RequestOption) (err error) {
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
func (r *WebhookService) DeleteJournalSubscriptionForPortal(ctx context.Context, portalID int64, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := fmt.Sprintf("webhooks-journal/subscriptions/2026-03/portals/%v", portalID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Delete the webhook settings for the specified app. Event subscriptions will not
// be deleted, but will be paused until another webhook is created.
func (r *WebhookService) DeleteSettings(ctx context.Context, appID int64, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := fmt.Sprintf("webhooks/2026-03/%v/settings", appID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Remove a specific filter from the webhooks journal subscriptions. This operation
// is useful for managing and cleaning up filters that are no longer needed. Once
// deleted, the filter cannot be recovered.
func (r *WebhookService) DeleteSubscriptionFilter(ctx context.Context, filterID int64, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := fmt.Sprintf("webhooks-journal/subscriptions/2026-03/filters/%v", filterID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Retrieve the earliest batch of webhook journal entries for a specified count.
// This endpoint is useful for accessing historical webhook data in batches,
// allowing you to process or analyze older entries. The number of entries
// retrieved is determined by the count parameter.
func (r *WebhookService) GetEarliestJournalBatch(ctx context.Context, count int64, query WebhookGetEarliestJournalBatchParams, opts ...option.RequestOption) (res *shared.BatchResponseJournalFetchResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := fmt.Sprintf("webhooks-journal/journal/2026-03/batch/earliest/%v", count)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Retrieve the earliest entry from the webhooks journal for the specified portal.
// This endpoint is useful for accessing the first recorded webhook event in the
// journal, which can be helpful for auditing or debugging purposes.
func (r *WebhookService) GetEarliestJournalEntry(ctx context.Context, query WebhookGetEarliestJournalEntryParams, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "webhooks-journal/journal/2026-03/earliest"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Retrieve the earliest batch of webhook journal entries. This endpoint is useful
// for accessing the oldest available data in the webhook journal, allowing users
// to process or analyze historical webhook events. The number of entries to fetch
// is specified by the 'count' path parameter.
func (r *WebhookService) GetEarliestLocalJournalBatch(ctx context.Context, count int64, query WebhookGetEarliestLocalJournalBatchParams, opts ...option.RequestOption) (res *shared.BatchResponseJournalFetchResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := fmt.Sprintf("webhooks-journal/journal-local/2026-03/batch/earliest/%v", count)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Retrieve the earliest webhook journal entries for the specified portal. This
// endpoint can be used to access the oldest records available in the webhook
// journal, which may be useful for auditing or historical analysis.
func (r *WebhookService) GetEarliestLocalJournalEntry(ctx context.Context, query WebhookGetEarliestLocalJournalEntryParams, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "webhooks-journal/journal-local/2026-03/earliest"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Retrieve a specific event subscription by ID.
func (r *WebhookService) GetEventSubscription(ctx context.Context, subscriptionID int64, query WebhookGetEventSubscriptionParams, opts ...option.RequestOption) (res *SubscriptionResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := fmt.Sprintf("webhooks/2026-03/%v/subscriptions/%v", query.AppID, subscriptionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Execute a batch read operation on the webhooks journal for the specified date,
// 2026-03. This endpoint allows you to retrieve multiple entries from the webhooks
// journal in a single request, which can be useful for processing large amounts of
// data efficiently. Ensure that the request body is provided in the required
// format.
func (r *WebhookService) GetJournalBatchByRequest(ctx context.Context, params WebhookGetJournalBatchByRequestParams, opts ...option.RequestOption) (res *shared.BatchResponseJournalFetchResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "webhooks-journal/journal/2026-03/batch/read"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Retrieve a batch of webhook journal entries starting from a specified offset.
// This endpoint allows you to fetch a defined number of entries, which can be
// useful for processing large datasets in manageable chunks.
func (r *WebhookService) GetJournalBatchFromOffset(ctx context.Context, count int64, params WebhookGetJournalBatchFromOffsetParams, opts ...option.RequestOption) (res *shared.BatchResponseJournalFetchResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if params.Offset == "" {
		err = errors.New("missing required offset parameter")
		return nil, err
	}
	path := fmt.Sprintf("webhooks-journal/journal/2026-03/batch/%s/next/%v", url.PathEscape(params.Offset), count)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// Retrieve the status of a specific webhook journal entry using its unique status
// ID. This endpoint provides detailed information about the status, including
// whether it is pending, in progress, completed, failed, or expired. It is useful
// for monitoring and managing the state of webhook journal entries.
func (r *WebhookService) GetJournalStatus(ctx context.Context, statusID string, opts ...option.RequestOption) (res *shared.SnapshotStatusResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if statusID == "" {
		err = errors.New("missing required statusId parameter")
		return nil, err
	}
	path := fmt.Sprintf("webhooks-journal/journal/2026-03/status/%s", statusID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Retrieve details of a specific webhook subscription using its unique identifier.
// This endpoint is useful for obtaining information about a particular
// subscription, such as its actions, object type, and associated properties.
func (r *WebhookService) GetJournalSubscription(ctx context.Context, subscriptionID int64, opts ...option.RequestOption) (res *webhooks_journal.JournalSubscriptionResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := fmt.Sprintf("webhooks-journal/subscriptions/2026-03/%v", subscriptionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Retrieve the latest batch of webhook journal entries up to the specified count.
// This endpoint is useful for fetching recent webhook data for analysis or
// processing. The count parameter determines the maximum number of entries to
// return.
func (r *WebhookService) GetLatestJournalBatch(ctx context.Context, count int64, query WebhookGetLatestJournalBatchParams, opts ...option.RequestOption) (res *shared.BatchResponseJournalFetchResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := fmt.Sprintf("webhooks-journal/journal/2026-03/batch/latest/%v", count)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Retrieve the latest entries from the webhooks journal for the specified portal.
// This endpoint is useful for accessing the most recent webhook events and their
// statuses, allowing you to monitor and debug webhook activity effectively.
func (r *WebhookService) GetLatestJournalEntry(ctx context.Context, query WebhookGetLatestJournalEntryParams, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "webhooks-journal/journal/2026-03/latest"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Retrieve the latest batch of webhook journal entries. This endpoint allows you
// to specify the number of entries to fetch, providing a way to access the most
// recent webhook events processed by your HubSpot account.
func (r *WebhookService) GetLatestLocalJournalBatch(ctx context.Context, count int64, query WebhookGetLatestLocalJournalBatchParams, opts ...option.RequestOption) (res *shared.BatchResponseJournalFetchResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := fmt.Sprintf("webhooks-journal/journal-local/2026-03/batch/latest/%v", count)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Retrieve the latest entries from the webhooks journal for the specified portal.
// This endpoint is useful for accessing the most recent webhook events that have
// been logged, allowing for real-time monitoring or debugging of webhook
// activities.
func (r *WebhookService) GetLatestLocalJournalEntry(ctx context.Context, query WebhookGetLatestLocalJournalEntryParams, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "webhooks-journal/journal-local/2026-03/latest"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Execute a batch read operation on the webhooks journal. This endpoint allows you
// to retrieve a batch of webhook journal entries by providing the necessary input
// data. It is useful for processing multiple records in a single request,
// streamlining data retrieval tasks.
func (r *WebhookService) GetLocalJournalBatchByRequest(ctx context.Context, params WebhookGetLocalJournalBatchByRequestParams, opts ...option.RequestOption) (res *shared.BatchResponseJournalFetchResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "webhooks-journal/journal-local/2026-03/batch/read"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Retrieve a batch of webhook journal entries starting from a specified offset.
// This endpoint is useful for paginating through large sets of webhook data. The
// number of entries returned is determined by the 'count' parameter.
func (r *WebhookService) GetLocalJournalBatchFromOffset(ctx context.Context, count int64, params WebhookGetLocalJournalBatchFromOffsetParams, opts ...option.RequestOption) (res *shared.BatchResponseJournalFetchResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if params.Offset == "" {
		err = errors.New("missing required offset parameter")
		return nil, err
	}
	path := fmt.Sprintf("webhooks-journal/journal-local/2026-03/batch/%s/next/%v", url.PathEscape(params.Offset), count)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// Retrieve the status of a specific webhook journal entry using its unique status
// ID. This endpoint is useful for monitoring the progress or outcome of webhook
// journal entries, allowing you to check if an entry is pending, in progress,
// completed, failed, or expired.
func (r *WebhookService) GetLocalJournalStatus(ctx context.Context, statusID string, opts ...option.RequestOption) (res *shared.SnapshotStatusResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if statusID == "" {
		err = errors.New("missing required statusId parameter")
		return nil, err
	}
	path := fmt.Sprintf("webhooks-journal/journal-local/2026-03/status/%s", statusID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Retrieve the next set of entries from the webhooks journal starting from a
// specified offset. This endpoint is useful for paginating through journal entries
// to process or analyze webhook events sequentially.
func (r *WebhookService) GetNextJournalEntries(ctx context.Context, offset string, query WebhookGetNextJournalEntriesParams, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if offset == "" {
		err = errors.New("missing required offset parameter")
		return nil, err
	}
	path := fmt.Sprintf("webhooks-journal/journal/2026-03/offset/%s/next", url.PathEscape(offset))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Retrieve the next set of webhook journal entries starting from a specified
// offset. This endpoint is useful for paginating through large sets of webhook
// data, allowing you to continue from where a previous request left off.
func (r *WebhookService) GetNextLocalJournalEntries(ctx context.Context, offset string, query WebhookGetNextLocalJournalEntriesParams, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if offset == "" {
		err = errors.New("missing required offset parameter")
		return nil, err
	}
	path := fmt.Sprintf("webhooks-journal/journal-local/2026-03/offset/%s/next", url.PathEscape(offset))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Retrieve the webhook settings for the specified app, including the webhook’s
// target URL, throttle configuration, and create/update date.
func (r *WebhookService) GetSettings(ctx context.Context, appID int64, opts ...option.RequestOption) (res *SettingsResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := fmt.Sprintf("webhooks/2026-03/%v/settings", appID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Retrieve a specific filter associated with a webhook journal subscription. This
// endpoint allows you to access the details of the filter identified by the
// filterId, which is useful for managing and understanding the conditions applied
// to webhook events.
func (r *WebhookService) GetSubscriptionFilter(ctx context.Context, filterID int64, opts ...option.RequestOption) (res *shared.FilterResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := fmt.Sprintf("webhooks-journal/subscriptions/2026-03/filters/%v", filterID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Retrieve event subscriptions for the specified app.
func (r *WebhookService) ListEventSubscriptions(ctx context.Context, appID int64, opts ...option.RequestOption) (res *SubscriptionListResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := fmt.Sprintf("webhooks/2026-03/%v/subscriptions", appID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Retrieve a list of webhook journal subscriptions for the specified version. This
// endpoint allows you to view all active subscriptions without pagination. It is
// useful for monitoring and managing webhook subscriptions in your HubSpot
// account.
func (r *WebhookService) ListJournalSubscriptions(ctx context.Context, opts ...option.RequestOption) (res *webhooks_journal.JournalCollectionResponseSubscriptionResponseNoPaging, err error) {
	opts = slices.Concat(r.options, opts)
	path := "webhooks-journal/subscriptions/2026-03"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Retrieve the filters associated with a specific webhook subscription. This
// endpoint allows you to view the filters applied to a subscription, which can
// help in managing and understanding the conditions set for webhook events.
func (r *WebhookService) ListSubscriptionFilters(ctx context.Context, subscriptionID int64, opts ...option.RequestOption) (res *[]shared.FilterResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := fmt.Sprintf("webhooks-journal/subscriptions/2026-03/filters/subscription/%v", subscriptionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Update an existing event subscription by ID.
func (r *WebhookService) UpdateEventSubscription(ctx context.Context, subscriptionID int64, params WebhookUpdateEventSubscriptionParams, opts ...option.RequestOption) (res *SubscriptionResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := fmt.Sprintf("webhooks/2026-03/%v/subscriptions/%v", params.AppID, subscriptionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return res, err
}

// Update webhook settings for the specified app.
func (r *WebhookService) UpdateSettings(ctx context.Context, appID int64, body WebhookUpdateSettingsParams, opts ...option.RequestOption) (res *SettingsResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := fmt.Sprintf("webhooks/2026-03/%v/settings", appID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return res, err
}

// The property Inputs is required.
type BatchInputSubscriptionBatchUpdateRequestParam struct {
	// An array of SubscriptionBatchUpdateRequest objects, each representing a
	// subscription to be updated. This property is required.
	Inputs []SubscriptionBatchUpdateRequestParam `json:"inputs,omitzero" api:"required"`
	paramObj
}

func (r BatchInputSubscriptionBatchUpdateRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow BatchInputSubscriptionBatchUpdateRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BatchInputSubscriptionBatchUpdateRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BatchResponseSubscriptionResponse struct {
	// The date and time when the batch operation was completed, in ISO 8601 format.
	CompletedAt time.Time `json:"completedAt" api:"required" format:"date-time"`
	// An array containing the results of the batch operation, with each item
	// representing an individual subscription response.
	Results []SubscriptionResponse `json:"results" api:"required"`
	// The date and time when the batch operation started, in ISO 8601 format.
	StartedAt time.Time `json:"startedAt" api:"required" format:"date-time"`
	// The current status of the batch operation. Valid values include 'PENDING',
	// 'PROCESSING', 'CANCELED', and 'COMPLETE'.
	//
	// Any of "CANCELED", "COMPLETE", "PENDING", "PROCESSING".
	Status BatchResponseSubscriptionResponseStatus `json:"status" api:"required"`
	// A map of link names to associated URIs providing additional information about
	// the batch operation.
	Links map[string]string `json:"links"`
	// The date and time when the batch operation was requested, in ISO 8601 format.
	RequestedAt time.Time `json:"requestedAt" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CompletedAt respjson.Field
		Results     respjson.Field
		StartedAt   respjson.Field
		Status      respjson.Field
		Links       respjson.Field
		RequestedAt respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BatchResponseSubscriptionResponse) RawJSON() string { return r.JSON.raw }
func (r *BatchResponseSubscriptionResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The current status of the batch operation. Valid values include 'PENDING',
// 'PROCESSING', 'CANCELED', and 'COMPLETE'.
type BatchResponseSubscriptionResponseStatus string

const (
	BatchResponseSubscriptionResponseStatusCanceled   BatchResponseSubscriptionResponseStatus = "CANCELED"
	BatchResponseSubscriptionResponseStatusComplete   BatchResponseSubscriptionResponseStatus = "COMPLETE"
	BatchResponseSubscriptionResponseStatusPending    BatchResponseSubscriptionResponseStatus = "PENDING"
	BatchResponseSubscriptionResponseStatusProcessing BatchResponseSubscriptionResponseStatus = "PROCESSING"
)

// The properties TargetURL, Throttling are required.
type SettingsChangeRequestParam struct {
	// The URL to which webhook events will be sent. It is a string.
	TargetURL  string                  `json:"targetUrl" api:"required"`
	Throttling ThrottlingSettingsParam `json:"throttling,omitzero" api:"required"`
	paramObj
}

func (r SettingsChangeRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow SettingsChangeRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SettingsChangeRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SettingsResponse struct {
	// The date and time when the webhook settings were created, in ISO 8601 format.
	CreatedAt time.Time `json:"createdAt" api:"required" format:"date-time"`
	// The URL to which the webhook events will be sent. It is a string.
	TargetURL  string             `json:"targetUrl" api:"required"`
	Throttling ThrottlingSettings `json:"throttling" api:"required"`
	// The date and time when the webhook settings were last updated, in ISO 8601
	// format.
	UpdatedAt time.Time `json:"updatedAt" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreatedAt   respjson.Field
		TargetURL   respjson.Field
		Throttling  respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SettingsResponse) RawJSON() string { return r.JSON.raw }
func (r *SettingsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties ID, Active are required.
type SubscriptionBatchUpdateRequestParam struct {
	// The unique identifier for the subscription. It is an integer.
	ID int64 `json:"id" api:"required"`
	// A boolean indicating whether the subscription is active.
	Active bool `json:"active" api:"required"`
	paramObj
}

func (r SubscriptionBatchUpdateRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow SubscriptionBatchUpdateRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SubscriptionBatchUpdateRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Active, EventType are required.
type SubscriptionCreateRequestParam struct {
	// A boolean indicating whether the subscription is active.
	Active bool `json:"active" api:"required"`
	// A string representing the type of event to subscribe to. Valid values include
	// various property changes, creations, deletions, merges, restorations,
	// association changes, and event completions.
	//
	// Any of "company.associationChange", "company.creation", "company.deletion",
	// "company.merge", "company.propertyChange", "company.restore",
	// "contact.associationChange", "contact.creation", "contact.deletion",
	// "contact.merge", "contact.privacyDeletion", "contact.propertyChange",
	// "contact.restore", "conversation.creation", "conversation.deletion",
	// "conversation.newMessage", "conversation.privacyDeletion",
	// "conversation.propertyChange", "deal.associationChange", "deal.creation",
	// "deal.deletion", "deal.merge", "deal.propertyChange", "deal.restore",
	// "event.completed", "line_item.associationChange", "line_item.creation",
	// "line_item.deletion", "line_item.merge", "line_item.propertyChange",
	// "line_item.restore", "object.associationChange", "object.creation",
	// "object.deletion", "object.merge", "object.propertyChange", "object.restore",
	// "product.creation", "product.deletion", "product.merge",
	// "product.propertyChange", "product.restore", "ticket.associationChange",
	// "ticket.creation", "ticket.deletion", "ticket.merge", "ticket.propertyChange",
	// "ticket.restore".
	EventType SubscriptionCreateRequestEventType `json:"eventType,omitzero" api:"required"`
	// A string providing a human-readable name for the event type.
	EventTypeName param.Opt[string] `json:"eventTypeName,omitzero"`
	// A string representing the ID of the object type associated with the
	// subscription.
	ObjectTypeID param.Opt[string] `json:"objectTypeId,omitzero"`
	// A string indicating the specific property name related to the event type, if
	// applicable.
	PropertyName param.Opt[string] `json:"propertyName,omitzero"`
	paramObj
}

func (r SubscriptionCreateRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow SubscriptionCreateRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SubscriptionCreateRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A string representing the type of event to subscribe to. Valid values include
// various property changes, creations, deletions, merges, restorations,
// association changes, and event completions.
type SubscriptionCreateRequestEventType string

const (
	SubscriptionCreateRequestEventTypeCompanyAssociationChange    SubscriptionCreateRequestEventType = "company.associationChange"
	SubscriptionCreateRequestEventTypeCompanyCreation             SubscriptionCreateRequestEventType = "company.creation"
	SubscriptionCreateRequestEventTypeCompanyDeletion             SubscriptionCreateRequestEventType = "company.deletion"
	SubscriptionCreateRequestEventTypeCompanyMerge                SubscriptionCreateRequestEventType = "company.merge"
	SubscriptionCreateRequestEventTypeCompanyPropertyChange       SubscriptionCreateRequestEventType = "company.propertyChange"
	SubscriptionCreateRequestEventTypeCompanyRestore              SubscriptionCreateRequestEventType = "company.restore"
	SubscriptionCreateRequestEventTypeContactAssociationChange    SubscriptionCreateRequestEventType = "contact.associationChange"
	SubscriptionCreateRequestEventTypeContactCreation             SubscriptionCreateRequestEventType = "contact.creation"
	SubscriptionCreateRequestEventTypeContactDeletion             SubscriptionCreateRequestEventType = "contact.deletion"
	SubscriptionCreateRequestEventTypeContactMerge                SubscriptionCreateRequestEventType = "contact.merge"
	SubscriptionCreateRequestEventTypeContactPrivacyDeletion      SubscriptionCreateRequestEventType = "contact.privacyDeletion"
	SubscriptionCreateRequestEventTypeContactPropertyChange       SubscriptionCreateRequestEventType = "contact.propertyChange"
	SubscriptionCreateRequestEventTypeContactRestore              SubscriptionCreateRequestEventType = "contact.restore"
	SubscriptionCreateRequestEventTypeConversationCreation        SubscriptionCreateRequestEventType = "conversation.creation"
	SubscriptionCreateRequestEventTypeConversationDeletion        SubscriptionCreateRequestEventType = "conversation.deletion"
	SubscriptionCreateRequestEventTypeConversationNewMessage      SubscriptionCreateRequestEventType = "conversation.newMessage"
	SubscriptionCreateRequestEventTypeConversationPrivacyDeletion SubscriptionCreateRequestEventType = "conversation.privacyDeletion"
	SubscriptionCreateRequestEventTypeConversationPropertyChange  SubscriptionCreateRequestEventType = "conversation.propertyChange"
	SubscriptionCreateRequestEventTypeDealAssociationChange       SubscriptionCreateRequestEventType = "deal.associationChange"
	SubscriptionCreateRequestEventTypeDealCreation                SubscriptionCreateRequestEventType = "deal.creation"
	SubscriptionCreateRequestEventTypeDealDeletion                SubscriptionCreateRequestEventType = "deal.deletion"
	SubscriptionCreateRequestEventTypeDealMerge                   SubscriptionCreateRequestEventType = "deal.merge"
	SubscriptionCreateRequestEventTypeDealPropertyChange          SubscriptionCreateRequestEventType = "deal.propertyChange"
	SubscriptionCreateRequestEventTypeDealRestore                 SubscriptionCreateRequestEventType = "deal.restore"
	SubscriptionCreateRequestEventTypeEventCompleted              SubscriptionCreateRequestEventType = "event.completed"
	SubscriptionCreateRequestEventTypeLineItemAssociationChange   SubscriptionCreateRequestEventType = "line_item.associationChange"
	SubscriptionCreateRequestEventTypeLineItemCreation            SubscriptionCreateRequestEventType = "line_item.creation"
	SubscriptionCreateRequestEventTypeLineItemDeletion            SubscriptionCreateRequestEventType = "line_item.deletion"
	SubscriptionCreateRequestEventTypeLineItemMerge               SubscriptionCreateRequestEventType = "line_item.merge"
	SubscriptionCreateRequestEventTypeLineItemPropertyChange      SubscriptionCreateRequestEventType = "line_item.propertyChange"
	SubscriptionCreateRequestEventTypeLineItemRestore             SubscriptionCreateRequestEventType = "line_item.restore"
	SubscriptionCreateRequestEventTypeObjectAssociationChange     SubscriptionCreateRequestEventType = "object.associationChange"
	SubscriptionCreateRequestEventTypeObjectCreation              SubscriptionCreateRequestEventType = "object.creation"
	SubscriptionCreateRequestEventTypeObjectDeletion              SubscriptionCreateRequestEventType = "object.deletion"
	SubscriptionCreateRequestEventTypeObjectMerge                 SubscriptionCreateRequestEventType = "object.merge"
	SubscriptionCreateRequestEventTypeObjectPropertyChange        SubscriptionCreateRequestEventType = "object.propertyChange"
	SubscriptionCreateRequestEventTypeObjectRestore               SubscriptionCreateRequestEventType = "object.restore"
	SubscriptionCreateRequestEventTypeProductCreation             SubscriptionCreateRequestEventType = "product.creation"
	SubscriptionCreateRequestEventTypeProductDeletion             SubscriptionCreateRequestEventType = "product.deletion"
	SubscriptionCreateRequestEventTypeProductMerge                SubscriptionCreateRequestEventType = "product.merge"
	SubscriptionCreateRequestEventTypeProductPropertyChange       SubscriptionCreateRequestEventType = "product.propertyChange"
	SubscriptionCreateRequestEventTypeProductRestore              SubscriptionCreateRequestEventType = "product.restore"
	SubscriptionCreateRequestEventTypeTicketAssociationChange     SubscriptionCreateRequestEventType = "ticket.associationChange"
	SubscriptionCreateRequestEventTypeTicketCreation              SubscriptionCreateRequestEventType = "ticket.creation"
	SubscriptionCreateRequestEventTypeTicketDeletion              SubscriptionCreateRequestEventType = "ticket.deletion"
	SubscriptionCreateRequestEventTypeTicketMerge                 SubscriptionCreateRequestEventType = "ticket.merge"
	SubscriptionCreateRequestEventTypeTicketPropertyChange        SubscriptionCreateRequestEventType = "ticket.propertyChange"
	SubscriptionCreateRequestEventTypeTicketRestore               SubscriptionCreateRequestEventType = "ticket.restore"
)

type SubscriptionListResponse struct {
	// An array of subscription responses, each detailing a specific subscription's
	// properties and status.
	Results []SubscriptionResponse `json:"results" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Results     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SubscriptionListResponse) RawJSON() string { return r.JSON.raw }
func (r *SubscriptionListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SubscriptionPatchRequestParam struct {
	// A boolean indicating whether the subscription is active. If true, the
	// subscription is active; if false, it is inactive.
	Active param.Opt[bool] `json:"active,omitzero"`
	paramObj
}

func (r SubscriptionPatchRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow SubscriptionPatchRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SubscriptionPatchRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SubscriptionResponse struct {
	// The unique identifier for the subscription. It is an integer formatted as int64.
	ID string `json:"id" api:"required"`
	// A boolean indicating whether the subscription is currently active.
	Active bool `json:"active" api:"required"`
	// The date and time when the subscription was created, in ISO 8601 format.
	CreatedAt time.Time `json:"createdAt" api:"required" format:"date-time"`
	// The type of event that triggers the subscription. Valid values include various
	// property changes, creations, deletions, merges, restores, and association
	// changes for different HubSpot objects.
	//
	// Any of "company.associationChange", "company.creation", "company.deletion",
	// "company.merge", "company.propertyChange", "company.restore",
	// "contact.associationChange", "contact.creation", "contact.deletion",
	// "contact.merge", "contact.privacyDeletion", "contact.propertyChange",
	// "contact.restore", "conversation.creation", "conversation.deletion",
	// "conversation.newMessage", "conversation.privacyDeletion",
	// "conversation.propertyChange", "deal.associationChange", "deal.creation",
	// "deal.deletion", "deal.merge", "deal.propertyChange", "deal.restore",
	// "event.completed", "line_item.associationChange", "line_item.creation",
	// "line_item.deletion", "line_item.merge", "line_item.propertyChange",
	// "line_item.restore", "object.associationChange", "object.creation",
	// "object.deletion", "object.merge", "object.propertyChange", "object.restore",
	// "product.creation", "product.deletion", "product.merge",
	// "product.propertyChange", "product.restore", "ticket.associationChange",
	// "ticket.creation", "ticket.deletion", "ticket.merge", "ticket.propertyChange",
	// "ticket.restore".
	EventType SubscriptionResponseEventType `json:"eventType" api:"required"`
	// The name of the event type for the subscription.
	EventTypeName string `json:"eventTypeName"`
	// The identifier for the object type associated with the subscription. It is a
	// string.
	ObjectTypeID string `json:"objectTypeId"`
	// The name of the property associated with the subscription event, if applicable.
	PropertyName string `json:"propertyName"`
	// The date and time when the subscription was last updated, in ISO 8601 format.
	UpdatedAt time.Time `json:"updatedAt" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		Active        respjson.Field
		CreatedAt     respjson.Field
		EventType     respjson.Field
		EventTypeName respjson.Field
		ObjectTypeID  respjson.Field
		PropertyName  respjson.Field
		UpdatedAt     respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SubscriptionResponse) RawJSON() string { return r.JSON.raw }
func (r *SubscriptionResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The type of event that triggers the subscription. Valid values include various
// property changes, creations, deletions, merges, restores, and association
// changes for different HubSpot objects.
type SubscriptionResponseEventType string

const (
	SubscriptionResponseEventTypeCompanyAssociationChange    SubscriptionResponseEventType = "company.associationChange"
	SubscriptionResponseEventTypeCompanyCreation             SubscriptionResponseEventType = "company.creation"
	SubscriptionResponseEventTypeCompanyDeletion             SubscriptionResponseEventType = "company.deletion"
	SubscriptionResponseEventTypeCompanyMerge                SubscriptionResponseEventType = "company.merge"
	SubscriptionResponseEventTypeCompanyPropertyChange       SubscriptionResponseEventType = "company.propertyChange"
	SubscriptionResponseEventTypeCompanyRestore              SubscriptionResponseEventType = "company.restore"
	SubscriptionResponseEventTypeContactAssociationChange    SubscriptionResponseEventType = "contact.associationChange"
	SubscriptionResponseEventTypeContactCreation             SubscriptionResponseEventType = "contact.creation"
	SubscriptionResponseEventTypeContactDeletion             SubscriptionResponseEventType = "contact.deletion"
	SubscriptionResponseEventTypeContactMerge                SubscriptionResponseEventType = "contact.merge"
	SubscriptionResponseEventTypeContactPrivacyDeletion      SubscriptionResponseEventType = "contact.privacyDeletion"
	SubscriptionResponseEventTypeContactPropertyChange       SubscriptionResponseEventType = "contact.propertyChange"
	SubscriptionResponseEventTypeContactRestore              SubscriptionResponseEventType = "contact.restore"
	SubscriptionResponseEventTypeConversationCreation        SubscriptionResponseEventType = "conversation.creation"
	SubscriptionResponseEventTypeConversationDeletion        SubscriptionResponseEventType = "conversation.deletion"
	SubscriptionResponseEventTypeConversationNewMessage      SubscriptionResponseEventType = "conversation.newMessage"
	SubscriptionResponseEventTypeConversationPrivacyDeletion SubscriptionResponseEventType = "conversation.privacyDeletion"
	SubscriptionResponseEventTypeConversationPropertyChange  SubscriptionResponseEventType = "conversation.propertyChange"
	SubscriptionResponseEventTypeDealAssociationChange       SubscriptionResponseEventType = "deal.associationChange"
	SubscriptionResponseEventTypeDealCreation                SubscriptionResponseEventType = "deal.creation"
	SubscriptionResponseEventTypeDealDeletion                SubscriptionResponseEventType = "deal.deletion"
	SubscriptionResponseEventTypeDealMerge                   SubscriptionResponseEventType = "deal.merge"
	SubscriptionResponseEventTypeDealPropertyChange          SubscriptionResponseEventType = "deal.propertyChange"
	SubscriptionResponseEventTypeDealRestore                 SubscriptionResponseEventType = "deal.restore"
	SubscriptionResponseEventTypeEventCompleted              SubscriptionResponseEventType = "event.completed"
	SubscriptionResponseEventTypeLineItemAssociationChange   SubscriptionResponseEventType = "line_item.associationChange"
	SubscriptionResponseEventTypeLineItemCreation            SubscriptionResponseEventType = "line_item.creation"
	SubscriptionResponseEventTypeLineItemDeletion            SubscriptionResponseEventType = "line_item.deletion"
	SubscriptionResponseEventTypeLineItemMerge               SubscriptionResponseEventType = "line_item.merge"
	SubscriptionResponseEventTypeLineItemPropertyChange      SubscriptionResponseEventType = "line_item.propertyChange"
	SubscriptionResponseEventTypeLineItemRestore             SubscriptionResponseEventType = "line_item.restore"
	SubscriptionResponseEventTypeObjectAssociationChange     SubscriptionResponseEventType = "object.associationChange"
	SubscriptionResponseEventTypeObjectCreation              SubscriptionResponseEventType = "object.creation"
	SubscriptionResponseEventTypeObjectDeletion              SubscriptionResponseEventType = "object.deletion"
	SubscriptionResponseEventTypeObjectMerge                 SubscriptionResponseEventType = "object.merge"
	SubscriptionResponseEventTypeObjectPropertyChange        SubscriptionResponseEventType = "object.propertyChange"
	SubscriptionResponseEventTypeObjectRestore               SubscriptionResponseEventType = "object.restore"
	SubscriptionResponseEventTypeProductCreation             SubscriptionResponseEventType = "product.creation"
	SubscriptionResponseEventTypeProductDeletion             SubscriptionResponseEventType = "product.deletion"
	SubscriptionResponseEventTypeProductMerge                SubscriptionResponseEventType = "product.merge"
	SubscriptionResponseEventTypeProductPropertyChange       SubscriptionResponseEventType = "product.propertyChange"
	SubscriptionResponseEventTypeProductRestore              SubscriptionResponseEventType = "product.restore"
	SubscriptionResponseEventTypeTicketAssociationChange     SubscriptionResponseEventType = "ticket.associationChange"
	SubscriptionResponseEventTypeTicketCreation              SubscriptionResponseEventType = "ticket.creation"
	SubscriptionResponseEventTypeTicketDeletion              SubscriptionResponseEventType = "ticket.deletion"
	SubscriptionResponseEventTypeTicketMerge                 SubscriptionResponseEventType = "ticket.merge"
	SubscriptionResponseEventTypeTicketPropertyChange        SubscriptionResponseEventType = "ticket.propertyChange"
	SubscriptionResponseEventTypeTicketRestore               SubscriptionResponseEventType = "ticket.restore"
)

type ThrottlingSettings struct {
	// The maximum number of concurrent requests allowed. This is an integer value.
	MaxConcurrentRequests int64 `json:"maxConcurrentRequests" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		MaxConcurrentRequests respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ThrottlingSettings) RawJSON() string { return r.JSON.raw }
func (r *ThrottlingSettings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this ThrottlingSettings to a ThrottlingSettingsParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// ThrottlingSettingsParam.Overrides()
func (r ThrottlingSettings) ToParam() ThrottlingSettingsParam {
	return param.Override[ThrottlingSettingsParam](json.RawMessage(r.RawJSON()))
}

// The property MaxConcurrentRequests is required.
type ThrottlingSettingsParam struct {
	// The maximum number of concurrent requests allowed. This is an integer value.
	MaxConcurrentRequests int64 `json:"maxConcurrentRequests" api:"required"`
	paramObj
}

func (r ThrottlingSettingsParam) MarshalJSON() (data []byte, err error) {
	type shadow ThrottlingSettingsParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ThrottlingSettingsParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebhookNewBatchEventSubscriptionsParams struct {
	BatchInputSubscriptionBatchUpdateRequest BatchInputSubscriptionBatchUpdateRequestParam
	paramObj
}

func (r WebhookNewBatchEventSubscriptionsParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.BatchInputSubscriptionBatchUpdateRequest)
}
func (r *WebhookNewBatchEventSubscriptionsParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebhookNewCrmSnapshotsParams struct {
	CrmObjectSnapshotBatchRequest shared.CrmObjectSnapshotBatchRequestParam
	paramObj
}

func (r WebhookNewCrmSnapshotsParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.CrmObjectSnapshotBatchRequest)
}
func (r *WebhookNewCrmSnapshotsParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebhookNewEventSubscriptionParams struct {
	SubscriptionCreateRequest SubscriptionCreateRequestParam
	paramObj
}

func (r WebhookNewEventSubscriptionParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.SubscriptionCreateRequest)
}
func (r *WebhookNewEventSubscriptionParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebhookNewJournalSubscriptionParams struct {
	SubscriptionUpsertRequest shared.SubscriptionUpsertRequestUnionParam
	paramObj
}

func (r WebhookNewJournalSubscriptionParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.SubscriptionUpsertRequest)
}
func (r *WebhookNewJournalSubscriptionParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebhookNewSubscriptionFilterParams struct {
	FilterCreateRequest shared.FilterCreateRequestParam
	paramObj
}

func (r WebhookNewSubscriptionFilterParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.FilterCreateRequest)
}
func (r *WebhookNewSubscriptionFilterParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebhookDeleteEventSubscriptionParams struct {
	AppID int64 `path:"appId" api:"required" json:"-"`
	paramObj
}

type WebhookGetEarliestJournalBatchParams struct {
	// The ID of the portal installation. This is an integer value that specifies which
	// portal's data to access.
	InstallPortalID param.Opt[int64] `query:"installPortalId,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [WebhookGetEarliestJournalBatchParams]'s query parameters as
// `url.Values`.
func (r WebhookGetEarliestJournalBatchParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type WebhookGetEarliestJournalEntryParams struct {
	// The ID of the portal installation to filter the journal entries by. This is an
	// integer value.
	InstallPortalID param.Opt[int64] `query:"installPortalId,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [WebhookGetEarliestJournalEntryParams]'s query parameters as
// `url.Values`.
func (r WebhookGetEarliestJournalEntryParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type WebhookGetEarliestLocalJournalBatchParams struct {
	// The ID of the portal installation to filter the webhook journal entries. This is
	// an optional integer parameter.
	InstallPortalID param.Opt[int64] `query:"installPortalId,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [WebhookGetEarliestLocalJournalBatchParams]'s query
// parameters as `url.Values`.
func (r WebhookGetEarliestLocalJournalBatchParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type WebhookGetEarliestLocalJournalEntryParams struct {
	// The ID of the portal for which to retrieve the earliest webhook journal entries.
	// This parameter is optional and should be an integer.
	InstallPortalID param.Opt[int64] `query:"installPortalId,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [WebhookGetEarliestLocalJournalEntryParams]'s query
// parameters as `url.Values`.
func (r WebhookGetEarliestLocalJournalEntryParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type WebhookGetEventSubscriptionParams struct {
	AppID int64 `path:"appId" api:"required" json:"-"`
	paramObj
}

type WebhookGetJournalBatchByRequestParams struct {
	BatchInputString shared.BatchInputStringParam
	// An integer representing the ID of the portal installation for which the webhooks
	// journal data should be retrieved.
	InstallPortalID param.Opt[int64] `query:"installPortalId,omitzero" json:"-"`
	paramObj
}

func (r WebhookGetJournalBatchByRequestParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.BatchInputString)
}
func (r *WebhookGetJournalBatchByRequestParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [WebhookGetJournalBatchByRequestParams]'s query parameters
// as `url.Values`.
func (r WebhookGetJournalBatchByRequestParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type WebhookGetJournalBatchFromOffsetParams struct {
	Offset string `path:"offset" api:"required" json:"-"`
	// The ID of the portal installation. This is an integer value.
	InstallPortalID param.Opt[int64] `query:"installPortalId,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [WebhookGetJournalBatchFromOffsetParams]'s query parameters
// as `url.Values`.
func (r WebhookGetJournalBatchFromOffsetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type WebhookGetLatestJournalBatchParams struct {
	// The ID of the portal installation. This is an integer value used to specify the
	// portal context for the request.
	InstallPortalID param.Opt[int64] `query:"installPortalId,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [WebhookGetLatestJournalBatchParams]'s query parameters as
// `url.Values`.
func (r WebhookGetLatestJournalBatchParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type WebhookGetLatestJournalEntryParams struct {
	// The unique identifier of the portal installation for which to retrieve the
	// latest journal entries. This parameter is optional and should be an integer.
	InstallPortalID param.Opt[int64] `query:"installPortalId,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [WebhookGetLatestJournalEntryParams]'s query parameters as
// `url.Values`.
func (r WebhookGetLatestJournalEntryParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type WebhookGetLatestLocalJournalBatchParams struct {
	// The ID of the portal where the webhook journal is installed. This parameter is
	// optional and used to specify the target portal.
	InstallPortalID param.Opt[int64] `query:"installPortalId,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [WebhookGetLatestLocalJournalBatchParams]'s query parameters
// as `url.Values`.
func (r WebhookGetLatestLocalJournalBatchParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type WebhookGetLatestLocalJournalEntryParams struct {
	// The ID of the portal for which to retrieve the latest journal entries. This is
	// an integer value.
	InstallPortalID param.Opt[int64] `query:"installPortalId,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [WebhookGetLatestLocalJournalEntryParams]'s query parameters
// as `url.Values`.
func (r WebhookGetLatestLocalJournalEntryParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type WebhookGetLocalJournalBatchByRequestParams struct {
	BatchInputString shared.BatchInputStringParam
	// The ID of the portal where the webhooks are installed. This parameter is
	// optional and is used to specify the target portal for the operation.
	InstallPortalID param.Opt[int64] `query:"installPortalId,omitzero" json:"-"`
	paramObj
}

func (r WebhookGetLocalJournalBatchByRequestParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.BatchInputString)
}
func (r *WebhookGetLocalJournalBatchByRequestParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [WebhookGetLocalJournalBatchByRequestParams]'s query
// parameters as `url.Values`.
func (r WebhookGetLocalJournalBatchByRequestParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type WebhookGetLocalJournalBatchFromOffsetParams struct {
	Offset string `path:"offset" api:"required" json:"-"`
	// The ID of the portal where the webhooks are installed. This is an optional
	// parameter.
	InstallPortalID param.Opt[int64] `query:"installPortalId,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [WebhookGetLocalJournalBatchFromOffsetParams]'s query
// parameters as `url.Values`.
func (r WebhookGetLocalJournalBatchFromOffsetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type WebhookGetNextJournalEntriesParams struct {
	// The ID of the portal where the webhooks are installed. This is an integer value.
	InstallPortalID param.Opt[int64] `query:"installPortalId,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [WebhookGetNextJournalEntriesParams]'s query parameters as
// `url.Values`.
func (r WebhookGetNextJournalEntriesParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type WebhookGetNextLocalJournalEntriesParams struct {
	// The ID of the portal installation to filter the webhook journal entries. This is
	// an integer value.
	InstallPortalID param.Opt[int64] `query:"installPortalId,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [WebhookGetNextLocalJournalEntriesParams]'s query parameters
// as `url.Values`.
func (r WebhookGetNextLocalJournalEntriesParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type WebhookUpdateEventSubscriptionParams struct {
	AppID                    int64 `path:"appId" api:"required" json:"-"`
	SubscriptionPatchRequest SubscriptionPatchRequestParam
	paramObj
}

func (r WebhookUpdateEventSubscriptionParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.SubscriptionPatchRequest)
}
func (r *WebhookUpdateEventSubscriptionParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebhookUpdateSettingsParams struct {
	SettingsChangeRequest SettingsChangeRequestParam
	paramObj
}

func (r WebhookUpdateSettingsParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.SettingsChangeRequest)
}
func (r *WebhookUpdateSettingsParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
