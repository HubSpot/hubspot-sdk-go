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

// Create a batch of CRM object snapshots for the specified portal. This endpoint
// allows you to capture the state of CRM objects at a specific point in time,
// which can be useful for auditing or historical analysis. The request requires a
// list of CRM object snapshot requests, each specifying the portal ID, object ID,
// object type ID, and properties to include in the snapshot.
func (r *WebhookService) NewCrmSnapshots(ctx context.Context, body WebhookNewCrmSnapshotsParams, opts ...option.RequestOption) (res *CrmObjectSnapshotBatchResponse, err error) {
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

// Create a new webhook subscription for the specified portal in the HubSpot
// account. This endpoint allows you to define the subscription details, including
// the types of events you want to subscribe to. The request body must include the
// necessary subscription information as defined by the SubscriptionUpsertRequest
// schema.
func (r *WebhookService) NewJournalSubscription(ctx context.Context, body WebhookNewJournalSubscriptionParams, opts ...option.RequestOption) (res *SubscriptionResponse1, err error) {
	opts = slices.Concat(r.options, opts)
	path := "webhooks-journal/subscriptions/2026-03"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Create a new filter for a webhook subscription in your HubSpot account. This
// endpoint allows you to define specific conditions that a webhook event must meet
// to trigger the subscription. It is useful for managing and customizing the
// behavior of webhook subscriptions based on specific criteria.
func (r *WebhookService) NewSubscriptionFilter(ctx context.Context, body WebhookNewSubscriptionFilterParams, opts ...option.RequestOption) (res *FilterCreateResponse, err error) {
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
// longer needed or relevant.
func (r *WebhookService) DeleteJournalSubscription(ctx context.Context, subscriptionID int64, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := fmt.Sprintf("webhooks-journal/subscriptions/2026-03/%v", subscriptionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Delete a webhook journal subscription for a specific portal. This operation
// removes the subscription associated with the given portalId, and no content is
// returned upon successful deletion.
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

// Delete a specific filter associated with a webhook journal subscription. This
// operation is useful for managing and cleaning up filters that are no longer
// needed in your subscription setup. The endpoint requires the unique identifier
// of the filter to be deleted.
func (r *WebhookService) DeleteSubscriptionFilter(ctx context.Context, filterID int64, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := fmt.Sprintf("webhooks-journal/subscriptions/2026-03/filters/%v", filterID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Retrieve the earliest batch of webhook journal entries up to the specified
// count. This endpoint is useful for fetching historical webhook data in batches,
// allowing you to process or analyze the earliest entries first.
func (r *WebhookService) GetEarliestJournalBatch(ctx context.Context, count int64, query WebhookGetEarliestJournalBatchParams, opts ...option.RequestOption) (res *BatchResponseJournalFetchResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := fmt.Sprintf("webhooks-journal/journal/2026-03/batch/earliest/%v", count)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Retrieve the earliest entry from the webhooks journal for the specified version.
// This endpoint is useful for accessing the oldest records available in the
// journal, which can be helpful for auditing or historical data analysis.
func (r *WebhookService) GetEarliestJournalEntry(ctx context.Context, query WebhookGetEarliestJournalEntryParams, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "webhooks-journal/journal/2026-03/earliest"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Retrieve the earliest batch of webhook journal entries based on the specified
// count. This endpoint is useful for fetching a specific number of the earliest
// entries in the webhook journal for analysis or processing.
func (r *WebhookService) GetEarliestLocalJournalBatch(ctx context.Context, count int64, query WebhookGetEarliestLocalJournalBatchParams, opts ...option.RequestOption) (res *BatchResponseJournalFetchResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := fmt.Sprintf("webhooks-journal/journal-local/2026-03/batch/earliest/%v", count)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Retrieve the earliest entry from the webhooks journal for the specified portal.
// This endpoint is useful for accessing the oldest records in the journal, which
// can be helpful for auditing or tracking purposes.
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

// Perform a batch read operation on the webhooks journal for the specified date.
// This endpoint allows you to retrieve multiple entries from the webhooks journal
// in a single request, which can be useful for processing large amounts of data
// efficiently.
func (r *WebhookService) GetJournalBatchByRequest(ctx context.Context, params WebhookGetJournalBatchByRequestParams, opts ...option.RequestOption) (res *BatchResponseJournalFetchResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "webhooks-journal/journal/2026-03/batch/read"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Retrieve a batch of webhook journal entries starting from a specified offset.
// This endpoint allows you to fetch a specified number of entries, making it
// useful for paginating through large sets of webhook journal data.
func (r *WebhookService) GetJournalBatchFromOffset(ctx context.Context, count int64, params WebhookGetJournalBatchFromOffsetParams, opts ...option.RequestOption) (res *BatchResponseJournalFetchResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if params.Offset == "" {
		err = errors.New("missing required offset parameter")
		return nil, err
	}
	path := fmt.Sprintf("webhooks-journal/journal/2026-03/batch/%s/next/%v", url.PathEscape(params.Offset), count)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// Retrieve the status of a specific webhook journal entry using its status ID.
// This endpoint is useful for checking the current state of a webhook process,
// such as whether it is pending, in progress, completed, failed, or expired.
func (r *WebhookService) GetJournalStatus(ctx context.Context, statusID string, opts ...option.RequestOption) (res *SnapshotStatusResponse, err error) {
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
// subscription's configuration and status within the HubSpot account.
func (r *WebhookService) GetJournalSubscription(ctx context.Context, subscriptionID int64, opts ...option.RequestOption) (res *SubscriptionResponse1, err error) {
	opts = slices.Concat(r.options, opts)
	path := fmt.Sprintf("webhooks-journal/subscriptions/2026-03/%v", subscriptionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Retrieve the latest batch of webhook journal entries. This endpoint allows you
// to specify the number of entries to fetch, providing a way to access recent
// webhook activity within your HubSpot account.
func (r *WebhookService) GetLatestJournalBatch(ctx context.Context, count int64, query WebhookGetLatestJournalBatchParams, opts ...option.RequestOption) (res *BatchResponseJournalFetchResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := fmt.Sprintf("webhooks-journal/journal/2026-03/batch/latest/%v", count)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Retrieve the latest entries from the webhooks journal for the specified portal.
// This endpoint is useful for accessing the most recent webhook events processed
// by your HubSpot account. It allows you to filter the results by the portal ID to
// ensure you are retrieving data relevant to a specific installation.
func (r *WebhookService) GetLatestJournalEntry(ctx context.Context, query WebhookGetLatestJournalEntryParams, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "webhooks-journal/journal/2026-03/latest"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Retrieve the latest batch of webhook journal entries. This endpoint is useful
// for accessing the most recent data entries processed by the webhook journal. It
// requires specifying the number of entries to retrieve.
func (r *WebhookService) GetLatestLocalJournalBatch(ctx context.Context, count int64, query WebhookGetLatestLocalJournalBatchParams, opts ...option.RequestOption) (res *BatchResponseJournalFetchResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := fmt.Sprintf("webhooks-journal/journal-local/2026-03/batch/latest/%v", count)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Retrieve the latest entries from the webhooks journal for the specified portal.
// This endpoint is useful for accessing the most recent webhook events that have
// been logged, allowing you to process or analyze them as needed.
func (r *WebhookService) GetLatestLocalJournalEntry(ctx context.Context, query WebhookGetLatestLocalJournalEntryParams, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "webhooks-journal/journal-local/2026-03/latest"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Perform a batch read operation on the webhooks journal. This endpoint allows you
// to read multiple entries from the journal in a single request. It requires a
// JSON request body specifying the inputs to be read. The response includes the
// results of the batch read operation, and may return multiple statuses if there
// are errors.
func (r *WebhookService) GetLocalJournalBatchByRequest(ctx context.Context, params WebhookGetLocalJournalBatchByRequestParams, opts ...option.RequestOption) (res *BatchResponseJournalFetchResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "webhooks-journal/journal-local/2026-03/batch/read"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Retrieve a batch of webhook journal entries starting from a specified offset.
// This endpoint allows you to fetch a defined number of entries, facilitating the
// processing of webhook data in manageable chunks.
func (r *WebhookService) GetLocalJournalBatchFromOffset(ctx context.Context, count int64, params WebhookGetLocalJournalBatchFromOffsetParams, opts ...option.RequestOption) (res *BatchResponseJournalFetchResponse, err error) {
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
// ID. This endpoint is useful for monitoring the progress or completion of webhook
// processing tasks.
func (r *WebhookService) GetLocalJournalStatus(ctx context.Context, statusID string, opts ...option.RequestOption) (res *SnapshotStatusResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if statusID == "" {
		err = errors.New("missing required statusId parameter")
		return nil, err
	}
	path := fmt.Sprintf("webhooks-journal/journal-local/2026-03/status/%s", statusID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Retrieve the next batch of webhook journal entries starting from a specified
// offset. This endpoint is useful for paginating through large sets of webhook
// data, allowing you to continue fetching entries from where you last left off.
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
// offset. This endpoint is useful for paginating through webhook journal data in a
// sequential manner, allowing you to fetch entries beyond a given point.
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

// Retrieve details of a specific filter associated with a webhook subscription in
// the HubSpot account. This endpoint is useful for accessing the configuration and
// conditions of a filter by its unique identifier.
func (r *WebhookService) GetSubscriptionFilter(ctx context.Context, filterID int64, opts ...option.RequestOption) (res *FilterResponse, err error) {
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

// Retrieve a list of webhook journal subscriptions for the specified API version.
// This endpoint provides details about each subscription, including actions,
// object types, and associated properties. It is useful for managing and reviewing
// current webhook subscriptions.
func (r *WebhookService) ListJournalSubscriptions(ctx context.Context, opts ...option.RequestOption) (res *CollectionResponseSubscriptionResponseNoPaging, err error) {
	opts = slices.Concat(r.options, opts)
	path := "webhooks-journal/subscriptions/2026-03"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Retrieve the filters associated with a specific webhook subscription in the
// HubSpot account. This endpoint is useful for obtaining detailed information
// about the filters applied to a given subscription, identified by its
// subscription ID.
func (r *WebhookService) ListSubscriptionFilters(ctx context.Context, subscriptionID int64, opts ...option.RequestOption) (res *[]FilterResponse, err error) {
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

type ActionOverrideRequest struct {
	// An array of strings, each representing an associated object type ID relevant to
	// the action override.
	AssociatedObjectTypeIDs []string `json:"associatedObjectTypeIds"`
	// An array of integers representing list IDs that are associated with the action
	// override. The integers are in int64 format.
	ListIDs []int64 `json:"listIds"`
	// An array of integers, each representing an object ID for which the action
	// override is applicable. The integers are in int64 format.
	ObjectIDs []int64 `json:"objectIds"`
	// An array of strings representing the properties to be overridden in the action.
	// Each string corresponds to a property name.
	Properties []string `json:"properties"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AssociatedObjectTypeIDs respjson.Field
		ListIDs                 respjson.Field
		ObjectIDs               respjson.Field
		Properties              respjson.Field
		ExtraFields             map[string]respjson.Field
		raw                     string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ActionOverrideRequest) RawJSON() string { return r.JSON.raw }
func (r *ActionOverrideRequest) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties EventTypeID, Properties, SubscriptionType are required.
type AppLifecycleEventSubscriptionUpsertRequestParam struct {
	EventTypeID string   `json:"eventTypeId" api:"required"`
	Properties  []string `json:"properties,omitzero" api:"required"`
	// Any of "OBJECT", "ASSOCIATION", "EVENT", "APP_LIFECYCLE_EVENT",
	// "LIST_MEMBERSHIP", "GDPR_PRIVACY_DELETION".
	SubscriptionType AppLifecycleEventSubscriptionUpsertRequestSubscriptionType `json:"subscriptionType,omitzero" api:"required"`
	paramObj
}

func (r AppLifecycleEventSubscriptionUpsertRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow AppLifecycleEventSubscriptionUpsertRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AppLifecycleEventSubscriptionUpsertRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AppLifecycleEventSubscriptionUpsertRequestSubscriptionType string

const (
	AppLifecycleEventSubscriptionUpsertRequestSubscriptionTypeObject              AppLifecycleEventSubscriptionUpsertRequestSubscriptionType = "OBJECT"
	AppLifecycleEventSubscriptionUpsertRequestSubscriptionTypeAssociation         AppLifecycleEventSubscriptionUpsertRequestSubscriptionType = "ASSOCIATION"
	AppLifecycleEventSubscriptionUpsertRequestSubscriptionTypeEvent               AppLifecycleEventSubscriptionUpsertRequestSubscriptionType = "EVENT"
	AppLifecycleEventSubscriptionUpsertRequestSubscriptionTypeAppLifecycleEvent   AppLifecycleEventSubscriptionUpsertRequestSubscriptionType = "APP_LIFECYCLE_EVENT"
	AppLifecycleEventSubscriptionUpsertRequestSubscriptionTypeListMembership      AppLifecycleEventSubscriptionUpsertRequestSubscriptionType = "LIST_MEMBERSHIP"
	AppLifecycleEventSubscriptionUpsertRequestSubscriptionTypeGdprPrivacyDeletion AppLifecycleEventSubscriptionUpsertRequestSubscriptionType = "GDPR_PRIVACY_DELETION"
)

// The properties Actions, AssociatedObjectTypeIDs, ObjectIDs, ObjectTypeID,
// PortalID, SubscriptionType are required.
type AssociationSubscriptionUpsertRequestParam struct {
	// Any of "CREATE", "UPDATE", "DELETE", "MERGE", "RESTORE", "ASSOCIATION_ADDED",
	// "ASSOCIATION_REMOVED", "SNAPSHOT", "APP_INSTALL", "APP_UNINSTALL",
	// "ADDED_TO_LIST", "REMOVED_FROM_LIST", "GDPR_DELETE".
	Actions                 []string `json:"actions,omitzero" api:"required"`
	AssociatedObjectTypeIDs []string `json:"associatedObjectTypeIds,omitzero" api:"required"`
	ObjectIDs               []int64  `json:"objectIds,omitzero" api:"required"`
	ObjectTypeID            string   `json:"objectTypeId" api:"required"`
	PortalID                int64    `json:"portalId" api:"required"`
	// Any of "OBJECT", "ASSOCIATION", "EVENT", "APP_LIFECYCLE_EVENT",
	// "LIST_MEMBERSHIP", "GDPR_PRIVACY_DELETION".
	SubscriptionType AssociationSubscriptionUpsertRequestSubscriptionType `json:"subscriptionType,omitzero" api:"required"`
	paramObj
}

func (r AssociationSubscriptionUpsertRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow AssociationSubscriptionUpsertRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AssociationSubscriptionUpsertRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AssociationSubscriptionUpsertRequestSubscriptionType string

const (
	AssociationSubscriptionUpsertRequestSubscriptionTypeObject              AssociationSubscriptionUpsertRequestSubscriptionType = "OBJECT"
	AssociationSubscriptionUpsertRequestSubscriptionTypeAssociation         AssociationSubscriptionUpsertRequestSubscriptionType = "ASSOCIATION"
	AssociationSubscriptionUpsertRequestSubscriptionTypeEvent               AssociationSubscriptionUpsertRequestSubscriptionType = "EVENT"
	AssociationSubscriptionUpsertRequestSubscriptionTypeAppLifecycleEvent   AssociationSubscriptionUpsertRequestSubscriptionType = "APP_LIFECYCLE_EVENT"
	AssociationSubscriptionUpsertRequestSubscriptionTypeListMembership      AssociationSubscriptionUpsertRequestSubscriptionType = "LIST_MEMBERSHIP"
	AssociationSubscriptionUpsertRequestSubscriptionTypeGdprPrivacyDeletion AssociationSubscriptionUpsertRequestSubscriptionType = "GDPR_PRIVACY_DELETION"
)

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

type BatchResponseJournalFetchResponse struct {
	// The date and time when the batch operation was completed, in ISO 8601 format.
	CompletedAt time.Time `json:"completedAt" api:"required" format:"date-time"`
	// An array of results from the batch operation, each represented as a
	// JournalFetchResponse object.
	Results []JournalFetchResponse `json:"results" api:"required"`
	// The date and time when the batch operation started, in ISO 8601 format.
	StartedAt time.Time `json:"startedAt" api:"required" format:"date-time"`
	// The current status of the batch operation. Valid values include 'PENDING',
	// 'PROCESSING', 'CANCELED', and 'COMPLETE'.
	//
	// Any of "CANCELED", "COMPLETE", "PENDING", "PROCESSING".
	Status BatchResponseJournalFetchResponseStatus `json:"status" api:"required"`
	// A map of link names to associated URIs related to the batch operation.
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
func (r BatchResponseJournalFetchResponse) RawJSON() string { return r.JSON.raw }
func (r *BatchResponseJournalFetchResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The current status of the batch operation. Valid values include 'PENDING',
// 'PROCESSING', 'CANCELED', and 'COMPLETE'.
type BatchResponseJournalFetchResponseStatus string

const (
	BatchResponseJournalFetchResponseStatusCanceled   BatchResponseJournalFetchResponseStatus = "CANCELED"
	BatchResponseJournalFetchResponseStatusComplete   BatchResponseJournalFetchResponseStatus = "COMPLETE"
	BatchResponseJournalFetchResponseStatusPending    BatchResponseJournalFetchResponseStatus = "PENDING"
	BatchResponseJournalFetchResponseStatusProcessing BatchResponseJournalFetchResponseStatus = "PROCESSING"
)

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

type CollectionResponseSubscriptionResponseNoPaging struct {
	// An array of SubscriptionResponse objects, each representing a subscription's
	// details such as actions, appId, createdAt, and other relevant properties.
	Results []SubscriptionResponse1 `json:"results" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Results     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CollectionResponseSubscriptionResponseNoPaging) RawJSON() string { return r.JSON.raw }
func (r *CollectionResponseSubscriptionResponseNoPaging) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Condition struct {
	// A string indicating the type of filter being applied. Valid value is
	// 'CRM_OBJECT_PROPERTY'.
	//
	// Any of "CRM_OBJECT_PROPERTY".
	FilterType ConditionFilterType `json:"filterType" api:"required"`
	// A string specifying the operation to be performed in the condition. Valid values
	// include 'EQ', 'N_EQ', 'LT', 'GT', 'LTE', 'GTE', 'CONTAINS', 'STARTS_WITH',
	// 'ENDS_WITH', 'IN', 'NOT_IN', 'IS_EMPTY', and 'IS_NOT_EMPTY'.
	//
	// Any of "CONTAINS", "ENDS_WITH", "EQ", "GT", "GTE", "IN", "IS_EMPTY",
	// "IS_NOT_EMPTY", "LT", "LTE", "N_EQ", "NOT_IN", "STARTS_WITH".
	Operator ConditionOperator `json:"operator" api:"required"`
	// A string representing the specific property of the CRM object that the condition
	// applies to.
	Property string `json:"property" api:"required"`
	// A string representing the value to be compared against the specified property
	// when using single-value operators.
	Value string `json:"value"`
	// An array of strings used to specify multiple values for comparison when using
	// operators that support multiple values, such as 'IN' or 'NOT_IN'.
	Values []string `json:"values"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FilterType  respjson.Field
		Operator    respjson.Field
		Property    respjson.Field
		Value       respjson.Field
		Values      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Condition) RawJSON() string { return r.JSON.raw }
func (r *Condition) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this Condition to a ConditionParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// ConditionParam.Overrides()
func (r Condition) ToParam() ConditionParam {
	return param.Override[ConditionParam](json.RawMessage(r.RawJSON()))
}

// A string indicating the type of filter being applied. Valid value is
// 'CRM_OBJECT_PROPERTY'.
type ConditionFilterType string

const (
	ConditionFilterTypeCrmObjectProperty ConditionFilterType = "CRM_OBJECT_PROPERTY"
)

// A string specifying the operation to be performed in the condition. Valid values
// include 'EQ', 'N_EQ', 'LT', 'GT', 'LTE', 'GTE', 'CONTAINS', 'STARTS_WITH',
// 'ENDS_WITH', 'IN', 'NOT_IN', 'IS_EMPTY', and 'IS_NOT_EMPTY'.
type ConditionOperator string

const (
	ConditionOperatorContains   ConditionOperator = "CONTAINS"
	ConditionOperatorEndsWith   ConditionOperator = "ENDS_WITH"
	ConditionOperatorEq         ConditionOperator = "EQ"
	ConditionOperatorGt         ConditionOperator = "GT"
	ConditionOperatorGte        ConditionOperator = "GTE"
	ConditionOperatorIn         ConditionOperator = "IN"
	ConditionOperatorIsEmpty    ConditionOperator = "IS_EMPTY"
	ConditionOperatorIsNotEmpty ConditionOperator = "IS_NOT_EMPTY"
	ConditionOperatorLt         ConditionOperator = "LT"
	ConditionOperatorLte        ConditionOperator = "LTE"
	ConditionOperatorNEq        ConditionOperator = "N_EQ"
	ConditionOperatorNotIn      ConditionOperator = "NOT_IN"
	ConditionOperatorStartsWith ConditionOperator = "STARTS_WITH"
)

// The properties FilterType, Operator, Property are required.
type ConditionParam struct {
	// A string indicating the type of filter being applied. Valid value is
	// 'CRM_OBJECT_PROPERTY'.
	//
	// Any of "CRM_OBJECT_PROPERTY".
	FilterType ConditionFilterType `json:"filterType,omitzero" api:"required"`
	// A string specifying the operation to be performed in the condition. Valid values
	// include 'EQ', 'N_EQ', 'LT', 'GT', 'LTE', 'GTE', 'CONTAINS', 'STARTS_WITH',
	// 'ENDS_WITH', 'IN', 'NOT_IN', 'IS_EMPTY', and 'IS_NOT_EMPTY'.
	//
	// Any of "CONTAINS", "ENDS_WITH", "EQ", "GT", "GTE", "IN", "IS_EMPTY",
	// "IS_NOT_EMPTY", "LT", "LTE", "N_EQ", "NOT_IN", "STARTS_WITH".
	Operator ConditionOperator `json:"operator,omitzero" api:"required"`
	// A string representing the specific property of the CRM object that the condition
	// applies to.
	Property string `json:"property" api:"required"`
	// A string representing the value to be compared against the specified property
	// when using single-value operators.
	Value param.Opt[string] `json:"value,omitzero"`
	// An array of strings used to specify multiple values for comparison when using
	// operators that support multiple values, such as 'IN' or 'NOT_IN'.
	Values []string `json:"values,omitzero"`
	paramObj
}

func (r ConditionParam) MarshalJSON() (data []byte, err error) {
	type shadow ConditionParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConditionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property SnapshotRequests is required.
type CrmObjectSnapshotBatchRequestParam struct {
	// An array of CrmObjectSnapshotRequest objects, each representing a request to
	// create a snapshot for a specific CRM object. This property is required.
	SnapshotRequests []CrmObjectSnapshotRequestParam `json:"snapshotRequests,omitzero" api:"required"`
	paramObj
}

func (r CrmObjectSnapshotBatchRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow CrmObjectSnapshotBatchRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CrmObjectSnapshotBatchRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CrmObjectSnapshotBatchResponse struct {
	// An array of CrmObjectSnapshotResponse objects, each representing the result of a
	// snapshot operation for a specific CRM object. This property is required.
	SnapshotResponses []CrmObjectSnapshotResponse `json:"snapshotResponses" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		SnapshotResponses respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CrmObjectSnapshotBatchResponse) RawJSON() string { return r.JSON.raw }
func (r *CrmObjectSnapshotBatchResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties ObjectID, ObjectTypeID, PortalID, Properties are required.
type CrmObjectSnapshotRequestParam struct {
	// An integer representing the unique identifier of the CRM object for which the
	// snapshot is requested.
	ObjectID int64 `json:"objectId" api:"required"`
	// A string representing the type identifier of the CRM object, specifying what
	// kind of object it is within HubSpot.
	ObjectTypeID string `json:"objectTypeId" api:"required"`
	// An integer representing the unique identifier of the HubSpot account (portal)
	// where the CRM object resides.
	PortalID int64 `json:"portalId" api:"required"`
	// An array of strings, each representing a property of the CRM object that should
	// be included in the snapshot.
	Properties []string `json:"properties,omitzero" api:"required"`
	paramObj
}

func (r CrmObjectSnapshotRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow CrmObjectSnapshotRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CrmObjectSnapshotRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CrmObjectSnapshotResponse struct {
	// An integer representing the unique identifier of the CRM object for which the
	// snapshot is taken.
	ObjectID int64 `json:"objectId" api:"required"`
	// A string indicating the type of the CRM object, such as contact, company, or
	// deal.
	ObjectTypeID string `json:"objectTypeId" api:"required"`
	// An integer representing the unique identifier of the HubSpot portal associated
	// with the CRM object.
	PortalID int64 `json:"portalId" api:"required"`
	// A UUID string representing the status identifier of the snapshot request,
	// indicating the current state of the snapshot process.
	SnapshotStatusID string `json:"snapshotStatusId" api:"required" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ObjectID         respjson.Field
		ObjectTypeID     respjson.Field
		PortalID         respjson.Field
		SnapshotStatusID respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CrmObjectSnapshotResponse) RawJSON() string { return r.JSON.raw }
func (r *CrmObjectSnapshotResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Defines a single condition for searching CRM objects, specifying the property to
// filter on, the operator to use (such as equals, greater than, or contains), and
// the value(s) to compare against.
type Filter struct {
	// An array of conditions that define the criteria for the filter. Each condition
	// specifies a property, an operator, and optionally a value or values.
	Conditions []Condition `json:"conditions" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Conditions  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Filter) RawJSON() string { return r.JSON.raw }
func (r *Filter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this Filter to a FilterParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// FilterParam.Overrides()
func (r Filter) ToParam() FilterParam {
	return param.Override[FilterParam](json.RawMessage(r.RawJSON()))
}

// Defines a single condition for searching CRM objects, specifying the property to
// filter on, the operator to use (such as equals, greater than, or contains), and
// the value(s) to compare against.
//
// The property Conditions is required.
type FilterParam struct {
	// An array of conditions that define the criteria for the filter. Each condition
	// specifies a property, an operator, and optionally a value or values.
	Conditions []ConditionParam `json:"conditions,omitzero" api:"required"`
	paramObj
}

func (r FilterParam) MarshalJSON() (data []byte, err error) {
	type shadow FilterParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FilterParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Filter, SubscriptionID are required.
type FilterCreateRequestParam struct {
	// Defines a single condition for searching CRM objects, specifying the property to
	// filter on, the operator to use (such as equals, greater than, or contains), and
	// the value(s) to compare against.
	Filter FilterParam `json:"filter,omitzero" api:"required"`
	// The unique identifier of the subscription to which the filter will be applied.
	// It is an integer formatted as int64.
	SubscriptionID int64 `json:"subscriptionId" api:"required"`
	paramObj
}

func (r FilterCreateRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow FilterCreateRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FilterCreateRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FilterCreateResponse struct {
	// The unique identifier for the created filter. It is an integer formatted as
	// int64.
	FilterID int64 `json:"filterId" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FilterID    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FilterCreateResponse) RawJSON() string { return r.JSON.raw }
func (r *FilterCreateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FilterResponse struct {
	// The unique identifier for the filter. It is an integer in int64 format.
	ID int64 `json:"id" api:"required"`
	// A Unix timestamp in milliseconds indicating when the filter was created.
	CreatedAt int64 `json:"createdAt" api:"required"`
	// Defines a single condition for searching CRM objects, specifying the property to
	// filter on, the operator to use (such as equals, greater than, or contains), and
	// the value(s) to compare against.
	Filter Filter `json:"filter" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		Filter      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FilterResponse) RawJSON() string { return r.JSON.raw }
func (r *FilterResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Actions, ObjectTypeID, PortalID, SubscriptionType are required.
type GdprPrivacyDeletionSubscriptionUpsertRequestParam struct {
	// Any of "CREATE", "UPDATE", "DELETE", "MERGE", "RESTORE", "ASSOCIATION_ADDED",
	// "ASSOCIATION_REMOVED", "SNAPSHOT", "APP_INSTALL", "APP_UNINSTALL",
	// "ADDED_TO_LIST", "REMOVED_FROM_LIST", "GDPR_DELETE".
	Actions      []string `json:"actions,omitzero" api:"required"`
	ObjectTypeID string   `json:"objectTypeId" api:"required"`
	PortalID     int64    `json:"portalId" api:"required"`
	// Any of "OBJECT", "ASSOCIATION", "EVENT", "APP_LIFECYCLE_EVENT",
	// "LIST_MEMBERSHIP", "GDPR_PRIVACY_DELETION".
	SubscriptionType GdprPrivacyDeletionSubscriptionUpsertRequestSubscriptionType `json:"subscriptionType,omitzero" api:"required"`
	paramObj
}

func (r GdprPrivacyDeletionSubscriptionUpsertRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow GdprPrivacyDeletionSubscriptionUpsertRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *GdprPrivacyDeletionSubscriptionUpsertRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GdprPrivacyDeletionSubscriptionUpsertRequestSubscriptionType string

const (
	GdprPrivacyDeletionSubscriptionUpsertRequestSubscriptionTypeObject              GdprPrivacyDeletionSubscriptionUpsertRequestSubscriptionType = "OBJECT"
	GdprPrivacyDeletionSubscriptionUpsertRequestSubscriptionTypeAssociation         GdprPrivacyDeletionSubscriptionUpsertRequestSubscriptionType = "ASSOCIATION"
	GdprPrivacyDeletionSubscriptionUpsertRequestSubscriptionTypeEvent               GdprPrivacyDeletionSubscriptionUpsertRequestSubscriptionType = "EVENT"
	GdprPrivacyDeletionSubscriptionUpsertRequestSubscriptionTypeAppLifecycleEvent   GdprPrivacyDeletionSubscriptionUpsertRequestSubscriptionType = "APP_LIFECYCLE_EVENT"
	GdprPrivacyDeletionSubscriptionUpsertRequestSubscriptionTypeListMembership      GdprPrivacyDeletionSubscriptionUpsertRequestSubscriptionType = "LIST_MEMBERSHIP"
	GdprPrivacyDeletionSubscriptionUpsertRequestSubscriptionTypeGdprPrivacyDeletion GdprPrivacyDeletionSubscriptionUpsertRequestSubscriptionType = "GDPR_PRIVACY_DELETION"
)

type JournalFetchResponse struct {
	// The unique identifier for the current offset of the journal entry, formatted as
	// a UUID.
	CurrentOffset string `json:"currentOffset" api:"required" format:"uuid"`
	// The date and time when the URL will expire, in ISO 8601 format.
	ExpiresAt time.Time `json:"expiresAt" api:"required" format:"date-time"`
	// The URL where the journal entry can be accessed. It is a string.
	URL string `json:"url" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CurrentOffset respjson.Field
		ExpiresAt     respjson.Field
		URL           respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r JournalFetchResponse) RawJSON() string { return r.JSON.raw }
func (r *JournalFetchResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Actions, ListIDs, ObjectIDs, PortalID, SubscriptionType are
// required.
type ListMembershipSubscriptionUpsertRequestParam struct {
	// Any of "CREATE", "UPDATE", "DELETE", "MERGE", "RESTORE", "ASSOCIATION_ADDED",
	// "ASSOCIATION_REMOVED", "SNAPSHOT", "APP_INSTALL", "APP_UNINSTALL",
	// "ADDED_TO_LIST", "REMOVED_FROM_LIST", "GDPR_DELETE".
	Actions   []string `json:"actions,omitzero" api:"required"`
	ListIDs   []int64  `json:"listIds,omitzero" api:"required"`
	ObjectIDs []int64  `json:"objectIds,omitzero" api:"required"`
	PortalID  int64    `json:"portalId" api:"required"`
	// Any of "OBJECT", "ASSOCIATION", "EVENT", "APP_LIFECYCLE_EVENT",
	// "LIST_MEMBERSHIP", "GDPR_PRIVACY_DELETION".
	SubscriptionType ListMembershipSubscriptionUpsertRequestSubscriptionType `json:"subscriptionType,omitzero" api:"required"`
	paramObj
}

func (r ListMembershipSubscriptionUpsertRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow ListMembershipSubscriptionUpsertRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ListMembershipSubscriptionUpsertRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ListMembershipSubscriptionUpsertRequestSubscriptionType string

const (
	ListMembershipSubscriptionUpsertRequestSubscriptionTypeObject              ListMembershipSubscriptionUpsertRequestSubscriptionType = "OBJECT"
	ListMembershipSubscriptionUpsertRequestSubscriptionTypeAssociation         ListMembershipSubscriptionUpsertRequestSubscriptionType = "ASSOCIATION"
	ListMembershipSubscriptionUpsertRequestSubscriptionTypeEvent               ListMembershipSubscriptionUpsertRequestSubscriptionType = "EVENT"
	ListMembershipSubscriptionUpsertRequestSubscriptionTypeAppLifecycleEvent   ListMembershipSubscriptionUpsertRequestSubscriptionType = "APP_LIFECYCLE_EVENT"
	ListMembershipSubscriptionUpsertRequestSubscriptionTypeListMembership      ListMembershipSubscriptionUpsertRequestSubscriptionType = "LIST_MEMBERSHIP"
	ListMembershipSubscriptionUpsertRequestSubscriptionTypeGdprPrivacyDeletion ListMembershipSubscriptionUpsertRequestSubscriptionType = "GDPR_PRIVACY_DELETION"
)

// The properties Actions, ObjectIDs, ObjectTypeID, PortalID, Properties,
// SubscriptionType are required.
type ObjectSubscriptionUpsertRequestParam struct {
	// Any of "CREATE", "UPDATE", "DELETE", "MERGE", "RESTORE", "ASSOCIATION_ADDED",
	// "ASSOCIATION_REMOVED", "SNAPSHOT", "APP_INSTALL", "APP_UNINSTALL",
	// "ADDED_TO_LIST", "REMOVED_FROM_LIST", "GDPR_DELETE".
	Actions      []string `json:"actions,omitzero" api:"required"`
	ObjectIDs    []int64  `json:"objectIds,omitzero" api:"required"`
	ObjectTypeID string   `json:"objectTypeId" api:"required"`
	PortalID     int64    `json:"portalId" api:"required"`
	Properties   []string `json:"properties,omitzero" api:"required"`
	// Any of "OBJECT", "ASSOCIATION", "EVENT", "APP_LIFECYCLE_EVENT",
	// "LIST_MEMBERSHIP", "GDPR_PRIVACY_DELETION".
	SubscriptionType ObjectSubscriptionUpsertRequestSubscriptionType `json:"subscriptionType,omitzero" api:"required"`
	paramObj
}

func (r ObjectSubscriptionUpsertRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow ObjectSubscriptionUpsertRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ObjectSubscriptionUpsertRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ObjectSubscriptionUpsertRequestSubscriptionType string

const (
	ObjectSubscriptionUpsertRequestSubscriptionTypeObject              ObjectSubscriptionUpsertRequestSubscriptionType = "OBJECT"
	ObjectSubscriptionUpsertRequestSubscriptionTypeAssociation         ObjectSubscriptionUpsertRequestSubscriptionType = "ASSOCIATION"
	ObjectSubscriptionUpsertRequestSubscriptionTypeEvent               ObjectSubscriptionUpsertRequestSubscriptionType = "EVENT"
	ObjectSubscriptionUpsertRequestSubscriptionTypeAppLifecycleEvent   ObjectSubscriptionUpsertRequestSubscriptionType = "APP_LIFECYCLE_EVENT"
	ObjectSubscriptionUpsertRequestSubscriptionTypeListMembership      ObjectSubscriptionUpsertRequestSubscriptionType = "LIST_MEMBERSHIP"
	ObjectSubscriptionUpsertRequestSubscriptionTypeGdprPrivacyDeletion ObjectSubscriptionUpsertRequestSubscriptionType = "GDPR_PRIVACY_DELETION"
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

type SnapshotStatusResponse struct {
	// The unique identifier for the snapshot operation, represented as a UUID.
	ID string `json:"id" api:"required" format:"uuid"`
	// The timestamp indicating when the snapshot operation was initiated, represented
	// as a Unix timestamp in milliseconds.
	InitiatedAt int64 `json:"initiatedAt" api:"required"`
	// The current status of the snapshot. Valid values include 'PENDING',
	// 'IN_PROGRESS', 'COMPLETED', 'FAILED', and 'EXPIRED'.
	//
	// Any of "COMPLETED", "EXPIRED", "FAILED", "IN_PROGRESS", "PENDING".
	Status SnapshotStatusResponseStatus `json:"status" api:"required"`
	// The timestamp indicating when the snapshot operation was completed, represented
	// as a Unix timestamp in milliseconds.
	CompletedAt int64 `json:"completedAt"`
	// A code representing the error that occurred, if any. Possible values are
	// 'TIMEOUT', 'VALIDATION_ERROR', 'INTERNAL_ERROR', and 'PERMISSION_DENIED'.
	//
	// Any of "INTERNAL_ERROR", "PERMISSION_DENIED", "TIMEOUT", "VALIDATION_ERROR".
	ErrorCode SnapshotStatusResponseErrorCode `json:"errorCode"`
	// A descriptive message providing additional information about the snapshot
	// operation or error.
	Message string `json:"message"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		InitiatedAt respjson.Field
		Status      respjson.Field
		CompletedAt respjson.Field
		ErrorCode   respjson.Field
		Message     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SnapshotStatusResponse) RawJSON() string { return r.JSON.raw }
func (r *SnapshotStatusResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The current status of the snapshot. Valid values include 'PENDING',
// 'IN_PROGRESS', 'COMPLETED', 'FAILED', and 'EXPIRED'.
type SnapshotStatusResponseStatus string

const (
	SnapshotStatusResponseStatusCompleted  SnapshotStatusResponseStatus = "COMPLETED"
	SnapshotStatusResponseStatusExpired    SnapshotStatusResponseStatus = "EXPIRED"
	SnapshotStatusResponseStatusFailed     SnapshotStatusResponseStatus = "FAILED"
	SnapshotStatusResponseStatusInProgress SnapshotStatusResponseStatus = "IN_PROGRESS"
	SnapshotStatusResponseStatusPending    SnapshotStatusResponseStatus = "PENDING"
)

// A code representing the error that occurred, if any. Possible values are
// 'TIMEOUT', 'VALIDATION_ERROR', 'INTERNAL_ERROR', and 'PERMISSION_DENIED'.
type SnapshotStatusResponseErrorCode string

const (
	SnapshotStatusResponseErrorCodeInternalError    SnapshotStatusResponseErrorCode = "INTERNAL_ERROR"
	SnapshotStatusResponseErrorCodePermissionDenied SnapshotStatusResponseErrorCode = "PERMISSION_DENIED"
	SnapshotStatusResponseErrorCodeTimeout          SnapshotStatusResponseErrorCode = "TIMEOUT"
	SnapshotStatusResponseErrorCodeValidationError  SnapshotStatusResponseErrorCode = "VALIDATION_ERROR"
)

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

type SubscriptionResponse1 struct {
	// The unique identifier for the subscription. It is an integer formatted as int64.
	ID int64 `json:"id" api:"required"`
	// A list of actions that trigger the subscription. Possible values include
	// 'CREATE', 'UPDATE', 'DELETE', 'MERGE', 'RESTORE', 'ASSOCIATION_ADDED',
	// 'ASSOCIATION_REMOVED', 'SNAPSHOT', 'APP_INSTALL', 'APP_UNINSTALL',
	// 'ADDED_TO_LIST', 'REMOVED_FROM_LIST', and 'GDPR_DELETE'.
	//
	// Any of "CREATE", "UPDATE", "DELETE", "MERGE", "RESTORE", "ASSOCIATION_ADDED",
	// "ASSOCIATION_REMOVED", "SNAPSHOT", "APP_INSTALL", "APP_UNINSTALL",
	// "ADDED_TO_LIST", "REMOVED_FROM_LIST", "GDPR_DELETE".
	Actions []string `json:"actions" api:"required"`
	// The unique identifier for the app associated with the subscription. It is an
	// integer formatted as int64.
	AppID int64 `json:"appId" api:"required"`
	// The date and time when the subscription was created, in ISO 8601 format.
	CreatedAt time.Time `json:"createdAt" api:"required" format:"date-time"`
	// The identifier for the object type associated with the subscription. It is a
	// string.
	ObjectTypeID string `json:"objectTypeId" api:"required"`
	// The type of subscription, which can be one of the following: 'OBJECT',
	// 'ASSOCIATION', 'EVENT', 'APP_LIFECYCLE_EVENT', 'LIST_MEMBERSHIP', or
	// 'GDPR_PRIVACY_DELETION'.
	//
	// Any of "APP_LIFECYCLE_EVENT", "ASSOCIATION", "EVENT", "GDPR_PRIVACY_DELETION",
	// "LIST_MEMBERSHIP", "OBJECT".
	SubscriptionType SubscriptionResponse1SubscriptionType `json:"subscriptionType" api:"required"`
	// The date and time when the subscription was last updated, in ISO 8601 format.
	UpdatedAt time.Time `json:"updatedAt" api:"required" format:"date-time"`
	// An object containing action overrides, where each key is an action and the value
	// is an ActionOverrideRequest object.
	ActionOverrides map[string]ActionOverrideRequest `json:"actionOverrides"`
	// A list of associated object type IDs. Each ID is a string.
	AssociatedObjectTypeIDs []string `json:"associatedObjectTypeIds"`
	// The ID of the user who created the subscription. It is an integer formatted as
	// int64.
	CreatedBy int64 `json:"createdBy"`
	// The date and time when the subscription was deleted, in ISO 8601 format, if
	// applicable.
	DeletedAt time.Time `json:"deletedAt" format:"date-time"`
	// A list of list IDs associated with the subscription. Each ID is an integer
	// formatted as int64.
	ListIDs []int64 `json:"listIds"`
	// A list of object IDs associated with the subscription. Each ID is an integer
	// formatted as int64.
	ObjectIDs []int64 `json:"objectIds"`
	// The unique identifier for the portal associated with the subscription. It is an
	// integer formatted as int64.
	PortalID int64 `json:"portalId"`
	// A list of property names associated with the subscription. Each property is a
	// string.
	Properties []string `json:"properties"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                      respjson.Field
		Actions                 respjson.Field
		AppID                   respjson.Field
		CreatedAt               respjson.Field
		ObjectTypeID            respjson.Field
		SubscriptionType        respjson.Field
		UpdatedAt               respjson.Field
		ActionOverrides         respjson.Field
		AssociatedObjectTypeIDs respjson.Field
		CreatedBy               respjson.Field
		DeletedAt               respjson.Field
		ListIDs                 respjson.Field
		ObjectIDs               respjson.Field
		PortalID                respjson.Field
		Properties              respjson.Field
		ExtraFields             map[string]respjson.Field
		raw                     string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SubscriptionResponse1) RawJSON() string { return r.JSON.raw }
func (r *SubscriptionResponse1) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The type of subscription, which can be one of the following: 'OBJECT',
// 'ASSOCIATION', 'EVENT', 'APP_LIFECYCLE_EVENT', 'LIST_MEMBERSHIP', or
// 'GDPR_PRIVACY_DELETION'.
type SubscriptionResponse1SubscriptionType string

const (
	SubscriptionResponse1SubscriptionTypeAppLifecycleEvent   SubscriptionResponse1SubscriptionType = "APP_LIFECYCLE_EVENT"
	SubscriptionResponse1SubscriptionTypeAssociation         SubscriptionResponse1SubscriptionType = "ASSOCIATION"
	SubscriptionResponse1SubscriptionTypeEvent               SubscriptionResponse1SubscriptionType = "EVENT"
	SubscriptionResponse1SubscriptionTypeGdprPrivacyDeletion SubscriptionResponse1SubscriptionType = "GDPR_PRIVACY_DELETION"
	SubscriptionResponse1SubscriptionTypeListMembership      SubscriptionResponse1SubscriptionType = "LIST_MEMBERSHIP"
	SubscriptionResponse1SubscriptionTypeObject              SubscriptionResponse1SubscriptionType = "OBJECT"
)

func SubscriptionUpsertRequestParamOfAppLifecycleEventSubscriptionUpsertRequest(eventTypeID string, properties []string, subscriptionType AppLifecycleEventSubscriptionUpsertRequestSubscriptionType) SubscriptionUpsertRequestUnionParam {
	var variant AppLifecycleEventSubscriptionUpsertRequestParam
	variant.EventTypeID = eventTypeID
	variant.Properties = properties
	variant.SubscriptionType = subscriptionType
	return SubscriptionUpsertRequestUnionParam{OfAppLifecycleEventSubscriptionUpsertRequest: &variant}
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type SubscriptionUpsertRequestUnionParam struct {
	OfObjectSubscriptionUpsertRequest              *ObjectSubscriptionUpsertRequestParam              `json:",omitzero,inline"`
	OfAssociationSubscriptionUpsertRequest         *AssociationSubscriptionUpsertRequestParam         `json:",omitzero,inline"`
	OfAppLifecycleEventSubscriptionUpsertRequest   *AppLifecycleEventSubscriptionUpsertRequestParam   `json:",omitzero,inline"`
	OfListMembershipSubscriptionUpsertRequest      *ListMembershipSubscriptionUpsertRequestParam      `json:",omitzero,inline"`
	OfGdprPrivacyDeletionSubscriptionUpsertRequest *GdprPrivacyDeletionSubscriptionUpsertRequestParam `json:",omitzero,inline"`
	paramUnion
}

func (u SubscriptionUpsertRequestUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfObjectSubscriptionUpsertRequest,
		u.OfAssociationSubscriptionUpsertRequest,
		u.OfAppLifecycleEventSubscriptionUpsertRequest,
		u.OfListMembershipSubscriptionUpsertRequest,
		u.OfGdprPrivacyDeletionSubscriptionUpsertRequest)
}
func (u *SubscriptionUpsertRequestUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

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
	CrmObjectSnapshotBatchRequest CrmObjectSnapshotBatchRequestParam
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
	SubscriptionUpsertRequest SubscriptionUpsertRequestUnionParam
	paramObj
}

func (r WebhookNewJournalSubscriptionParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.SubscriptionUpsertRequest)
}
func (r *WebhookNewJournalSubscriptionParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebhookNewSubscriptionFilterParams struct {
	FilterCreateRequest FilterCreateRequestParam
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
	// The ID of the portal installation to filter the webhook journal entries by. This
	// is an integer value.
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
	// The ID of the portal installation to filter the journal entries. It is an
	// integer.
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
	// The ID of the portal where the webhooks are installed. This is an integer value.
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
	// The ID of the portal installation to filter the journal entries by. This
	// parameter is optional and should be an integer.
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
	// The ID of the portal where the webhooks are installed. This is an integer value.
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
	// The ID of the portal installation. This is an integer value used to specify the
	// portal context for the request.
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
	// The ID of the portal installation. This is an integer value used to identify the
	// specific portal.
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
	// The ID of the portal installation to filter the journal entries. It is an
	// integer value.
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
	// The ID of the portal installation. This parameter is optional and used to filter
	// the journal entries by a specific portal.
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
	// The ID of the portal for which to retrieve the latest journal entries. This
	// parameter is optional and should be an integer.
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
	// optional and is used to specify the target portal.
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
	// The ID of the portal installation. This is an integer value used to specify the
	// portal context for the request.
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
	// The ID of the portal installation to filter the webhook journal entries. This is
	// an optional parameter.
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
	// The ID of the portal where the webhook is installed. This is an integer value.
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
