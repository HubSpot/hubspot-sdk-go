// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package webhooks_journal

import (
	"time"

	"github.com/HubSpot/hubspot-sdk-go/internal/apijson"
	"github.com/HubSpot/hubspot-sdk-go/option"
	"github.com/HubSpot/hubspot-sdk-go/packages/respjson"
	"github.com/HubSpot/hubspot-sdk-go/shared"
)

// WebhooksJournalService contains methods and other services that help with
// interacting with the hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWebhooksJournalService] method instead.
type WebhooksJournalService struct {
	options       []option.RequestOption
	Journal       JournalService
	JournalLocal  JournalLocalService
	Snapshots     SnapshotService
	Subscriptions SubscriptionService
}

// NewWebhooksJournalService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewWebhooksJournalService(opts ...option.RequestOption) (r WebhooksJournalService) {
	r = WebhooksJournalService{}
	r.options = opts
	r.Journal = NewJournalService(opts...)
	r.JournalLocal = NewJournalLocalService(opts...)
	r.Snapshots = NewSnapshotService(opts...)
	r.Subscriptions = NewSubscriptionService(opts...)
	return
}

type JournalCollectionResponseSubscriptionResponseNoPaging struct {
	// An array of subscription responses, where each item contains details about a
	// specific subscription. Each item follows the SubscriptionResponse schema.
	Results []JournalSubscriptionResponse `json:"results" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Results     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r JournalCollectionResponseSubscriptionResponseNoPaging) RawJSON() string { return r.JSON.raw }
func (r *JournalCollectionResponseSubscriptionResponseNoPaging) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type JournalSubscriptionResponse struct {
	// The unique identifier for the subscription, represented as an integer.
	ID int64 `json:"id" api:"required"`
	// An array of actions associated with the subscription. Valid actions include
	// 'CREATE', 'UPDATE', 'DELETE', 'MERGE', 'RESTORE', 'ASSOCIATION_ADDED',
	// 'ASSOCIATION_REMOVED', 'SNAPSHOT', 'APP_INSTALL', 'APP_UNINSTALL',
	// 'ADDED_TO_LIST', 'REMOVED_FROM_LIST', and 'GDPR_DELETE'.
	//
	// Any of "CREATE", "UPDATE", "DELETE", "MERGE", "RESTORE", "ASSOCIATION_ADDED",
	// "ASSOCIATION_REMOVED", "SNAPSHOT", "APP_INSTALL", "APP_UNINSTALL",
	// "ADDED_TO_LIST", "REMOVED_FROM_LIST", "GDPR_DELETE".
	Actions []string `json:"actions" api:"required"`
	// The unique identifier for the app associated with the subscription, represented
	// as an integer.
	AppID int64 `json:"appId" api:"required"`
	// The date and time when the subscription was created, in ISO 8601 format.
	CreatedAt time.Time `json:"createdAt" api:"required" format:"date-time"`
	// The identifier for the type of object associated with the subscription,
	// represented as a string.
	ObjectTypeID string `json:"objectTypeId" api:"required"`
	// The type of subscription, indicating the nature of events it pertains to. Valid
	// values include 'OBJECT', 'ASSOCIATION', 'EVENT', 'APP_LIFECYCLE_EVENT',
	// 'LIST_MEMBERSHIP', and 'GDPR_PRIVACY_DELETION'.
	//
	// Any of "APP_LIFECYCLE_EVENT", "ASSOCIATION", "EVENT", "GDPR_PRIVACY_DELETION",
	// "LIST_MEMBERSHIP", "OBJECT".
	SubscriptionType JournalSubscriptionResponseSubscriptionType `json:"subscriptionType" api:"required"`
	// The date and time when the subscription was last updated, in ISO 8601 format.
	UpdatedAt time.Time `json:"updatedAt" api:"required" format:"date-time"`
	// An object containing overrides for actions, where each key is an action and the
	// value is an ActionOverrideRequest object.
	ActionOverrides map[string]shared.ActionOverrideRequest `json:"actionOverrides"`
	// An array of strings representing the IDs of associated object types.
	AssociatedObjectTypeIDs []string `json:"associatedObjectTypeIds"`
	// The unique identifier of the user who created the subscription, represented as
	// an integer.
	CreatedBy int64 `json:"createdBy"`
	// The date and time when the subscription was deleted, in ISO 8601 format, if
	// applicable.
	DeletedAt time.Time `json:"deletedAt" format:"date-time"`
	// An array of integers representing the IDs of lists associated with the
	// subscription.
	ListIDs []int64 `json:"listIds"`
	// An array of integers representing the IDs of objects associated with the
	// subscription.
	ObjectIDs []int64 `json:"objectIds"`
	// The unique identifier for the portal associated with the subscription,
	// represented as an integer.
	PortalID int64 `json:"portalId"`
	// An array of strings representing the properties associated with the
	// subscription.
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
func (r JournalSubscriptionResponse) RawJSON() string { return r.JSON.raw }
func (r *JournalSubscriptionResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The type of subscription, indicating the nature of events it pertains to. Valid
// values include 'OBJECT', 'ASSOCIATION', 'EVENT', 'APP_LIFECYCLE_EVENT',
// 'LIST_MEMBERSHIP', and 'GDPR_PRIVACY_DELETION'.
type JournalSubscriptionResponseSubscriptionType string

const (
	JournalSubscriptionResponseSubscriptionTypeAppLifecycleEvent   JournalSubscriptionResponseSubscriptionType = "APP_LIFECYCLE_EVENT"
	JournalSubscriptionResponseSubscriptionTypeAssociation         JournalSubscriptionResponseSubscriptionType = "ASSOCIATION"
	JournalSubscriptionResponseSubscriptionTypeEvent               JournalSubscriptionResponseSubscriptionType = "EVENT"
	JournalSubscriptionResponseSubscriptionTypeGdprPrivacyDeletion JournalSubscriptionResponseSubscriptionType = "GDPR_PRIVACY_DELETION"
	JournalSubscriptionResponseSubscriptionTypeListMembership      JournalSubscriptionResponseSubscriptionType = "LIST_MEMBERSHIP"
	JournalSubscriptionResponseSubscriptionTypeObject              JournalSubscriptionResponseSubscriptionType = "OBJECT"
)
