// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package webhooks

import (
	"encoding/json"
	"time"

	"github.com/stainless-sdks/hubspot-sdk-go/internal/apijson"
	"github.com/stainless-sdks/hubspot-sdk-go/option"
	"github.com/stainless-sdks/hubspot-sdk-go/packages/param"
	"github.com/stainless-sdks/hubspot-sdk-go/packages/respjson"
)

// WebhookService contains methods and other services that help with interacting
// with the Hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWebhookService] method instead.
type WebhookService struct {
	Options       []option.RequestOption
	Settings      SettingService
	Subscriptions SubscriptionService
}

// NewWebhookService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewWebhookService(opts ...option.RequestOption) (r WebhookService) {
	r = WebhookService{}
	r.Options = opts
	r.Settings = NewSettingService(opts...)
	r.Subscriptions = NewSubscriptionService(opts...)
	return
}

// The property Inputs is required.
type BatchInputSubscriptionBatchUpdateRequestParam struct {
	Inputs []SubscriptionBatchUpdateRequestParam `json:"inputs,omitzero,required"`
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
	// The date and time when the batch operation was completed.
	CompletedAt time.Time `json:"completedAt,required" format:"date-time"`
	// The list of results from the batch operation.
	Results []SubscriptionResponse `json:"results,required"`
	// The date and time when the batch operation started.
	StartedAt time.Time `json:"startedAt,required" format:"date-time"`
	// The current status of the batch operation, which can be PENDING, PROCESSING,
	// CANCELED, or COMPLETE.
	//
	// Any of "PENDING", "PROCESSING", "CANCELED", "COMPLETE".
	Status BatchResponseSubscriptionResponseStatus `json:"status,required"`
	// A collection of related links associated with the batch operation.
	Links map[string]string `json:"links"`
	// The date and time when the batch operation was requested.
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

// The current status of the batch operation, which can be PENDING, PROCESSING,
// CANCELED, or COMPLETE.
type BatchResponseSubscriptionResponseStatus string

const (
	BatchResponseSubscriptionResponseStatusPending    BatchResponseSubscriptionResponseStatus = "PENDING"
	BatchResponseSubscriptionResponseStatusProcessing BatchResponseSubscriptionResponseStatus = "PROCESSING"
	BatchResponseSubscriptionResponseStatusCanceled   BatchResponseSubscriptionResponseStatus = "CANCELED"
	BatchResponseSubscriptionResponseStatusComplete   BatchResponseSubscriptionResponseStatus = "COMPLETE"
)

// New or updated webhook settings for an app.
//
// The properties TargetURL, Throttling are required.
type SettingsChangeRequestParam struct {
	// A publicly available URL for HubSpot to call where event payloads will be
	// delivered.
	TargetURL string `json:"targetUrl,required"`
	// Configuration details for webhook throttling.
	Throttling ThrottlingSettingsParam `json:"throttling,omitzero,required"`
	paramObj
}

func (r SettingsChangeRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow SettingsChangeRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SettingsChangeRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Webhook settings for an app.
type SettingsResponse struct {
	// When this subscription was created. Formatted as milliseconds from the
	// [Unix epoch](#).
	CreatedAt time.Time `json:"createdAt,required" format:"date-time"`
	// A publicly available URL for HubSpot to call where event payloads will be
	// delivered. See [link-so-some-doc](#) for details about the format of these event
	// payloads.
	TargetURL string `json:"targetUrl,required"`
	// Configuration details for webhook throttling.
	Throttling ThrottlingSettings `json:"throttling,required"`
	// When this subscription was last updated. Formatted as milliseconds from the
	// [Unix epoch](#).
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
	ID     int64 `json:"id,required"`
	Active bool  `json:"active,required"`
	paramObj
}

func (r SubscriptionBatchUpdateRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow SubscriptionBatchUpdateRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SubscriptionBatchUpdateRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// New webhook settings for an app.
//
// The property EventType is required.
type SubscriptionCreateRequestParam struct {
	// Type of event to listen for. Can be one of `create`, `delete`,
	// `deletedForPrivacy`, or `propertyChange`.
	//
	// Any of "contact.propertyChange", "company.propertyChange",
	// "deal.propertyChange", "ticket.propertyChange", "product.propertyChange",
	// "line_item.propertyChange", "contact.creation", "contact.deletion",
	// "contact.privacyDeletion", "company.creation", "company.deletion",
	// "deal.creation", "deal.deletion", "ticket.creation", "ticket.deletion",
	// "product.creation", "product.deletion", "line_item.creation",
	// "line_item.deletion", "conversation.creation", "conversation.deletion",
	// "conversation.newMessage", "conversation.privacyDeletion",
	// "conversation.propertyChange", "contact.merge", "company.merge", "deal.merge",
	// "ticket.merge", "product.merge", "line_item.merge", "contact.restore",
	// "company.restore", "deal.restore", "ticket.restore", "product.restore",
	// "line_item.restore", "contact.associationChange", "company.associationChange",
	// "deal.associationChange", "ticket.associationChange",
	// "line_item.associationChange", "object.propertyChange", "object.creation",
	// "object.deletion", "object.merge", "object.restore", "object.associationChange".
	EventType SubscriptionCreateRequestEventType `json:"eventType,omitzero,required"`
	// Determines if the subscription is active or paused. Defaults to false.
	Active       param.Opt[bool]   `json:"active,omitzero"`
	ObjectTypeID param.Opt[string] `json:"objectTypeId,omitzero"`
	// The internal name of the property to monitor for changes. Only applies when
	// `eventType` is `propertyChange`.
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

// Type of event to listen for. Can be one of `create`, `delete`,
// `deletedForPrivacy`, or `propertyChange`.
type SubscriptionCreateRequestEventType string

const (
	SubscriptionCreateRequestEventTypeContactPropertyChange       SubscriptionCreateRequestEventType = "contact.propertyChange"
	SubscriptionCreateRequestEventTypeCompanyPropertyChange       SubscriptionCreateRequestEventType = "company.propertyChange"
	SubscriptionCreateRequestEventTypeDealPropertyChange          SubscriptionCreateRequestEventType = "deal.propertyChange"
	SubscriptionCreateRequestEventTypeTicketPropertyChange        SubscriptionCreateRequestEventType = "ticket.propertyChange"
	SubscriptionCreateRequestEventTypeProductPropertyChange       SubscriptionCreateRequestEventType = "product.propertyChange"
	SubscriptionCreateRequestEventTypeLineItemPropertyChange      SubscriptionCreateRequestEventType = "line_item.propertyChange"
	SubscriptionCreateRequestEventTypeContactCreation             SubscriptionCreateRequestEventType = "contact.creation"
	SubscriptionCreateRequestEventTypeContactDeletion             SubscriptionCreateRequestEventType = "contact.deletion"
	SubscriptionCreateRequestEventTypeContactPrivacyDeletion      SubscriptionCreateRequestEventType = "contact.privacyDeletion"
	SubscriptionCreateRequestEventTypeCompanyCreation             SubscriptionCreateRequestEventType = "company.creation"
	SubscriptionCreateRequestEventTypeCompanyDeletion             SubscriptionCreateRequestEventType = "company.deletion"
	SubscriptionCreateRequestEventTypeDealCreation                SubscriptionCreateRequestEventType = "deal.creation"
	SubscriptionCreateRequestEventTypeDealDeletion                SubscriptionCreateRequestEventType = "deal.deletion"
	SubscriptionCreateRequestEventTypeTicketCreation              SubscriptionCreateRequestEventType = "ticket.creation"
	SubscriptionCreateRequestEventTypeTicketDeletion              SubscriptionCreateRequestEventType = "ticket.deletion"
	SubscriptionCreateRequestEventTypeProductCreation             SubscriptionCreateRequestEventType = "product.creation"
	SubscriptionCreateRequestEventTypeProductDeletion             SubscriptionCreateRequestEventType = "product.deletion"
	SubscriptionCreateRequestEventTypeLineItemCreation            SubscriptionCreateRequestEventType = "line_item.creation"
	SubscriptionCreateRequestEventTypeLineItemDeletion            SubscriptionCreateRequestEventType = "line_item.deletion"
	SubscriptionCreateRequestEventTypeConversationCreation        SubscriptionCreateRequestEventType = "conversation.creation"
	SubscriptionCreateRequestEventTypeConversationDeletion        SubscriptionCreateRequestEventType = "conversation.deletion"
	SubscriptionCreateRequestEventTypeConversationNewMessage      SubscriptionCreateRequestEventType = "conversation.newMessage"
	SubscriptionCreateRequestEventTypeConversationPrivacyDeletion SubscriptionCreateRequestEventType = "conversation.privacyDeletion"
	SubscriptionCreateRequestEventTypeConversationPropertyChange  SubscriptionCreateRequestEventType = "conversation.propertyChange"
	SubscriptionCreateRequestEventTypeContactMerge                SubscriptionCreateRequestEventType = "contact.merge"
	SubscriptionCreateRequestEventTypeCompanyMerge                SubscriptionCreateRequestEventType = "company.merge"
	SubscriptionCreateRequestEventTypeDealMerge                   SubscriptionCreateRequestEventType = "deal.merge"
	SubscriptionCreateRequestEventTypeTicketMerge                 SubscriptionCreateRequestEventType = "ticket.merge"
	SubscriptionCreateRequestEventTypeProductMerge                SubscriptionCreateRequestEventType = "product.merge"
	SubscriptionCreateRequestEventTypeLineItemMerge               SubscriptionCreateRequestEventType = "line_item.merge"
	SubscriptionCreateRequestEventTypeContactRestore              SubscriptionCreateRequestEventType = "contact.restore"
	SubscriptionCreateRequestEventTypeCompanyRestore              SubscriptionCreateRequestEventType = "company.restore"
	SubscriptionCreateRequestEventTypeDealRestore                 SubscriptionCreateRequestEventType = "deal.restore"
	SubscriptionCreateRequestEventTypeTicketRestore               SubscriptionCreateRequestEventType = "ticket.restore"
	SubscriptionCreateRequestEventTypeProductRestore              SubscriptionCreateRequestEventType = "product.restore"
	SubscriptionCreateRequestEventTypeLineItemRestore             SubscriptionCreateRequestEventType = "line_item.restore"
	SubscriptionCreateRequestEventTypeContactAssociationChange    SubscriptionCreateRequestEventType = "contact.associationChange"
	SubscriptionCreateRequestEventTypeCompanyAssociationChange    SubscriptionCreateRequestEventType = "company.associationChange"
	SubscriptionCreateRequestEventTypeDealAssociationChange       SubscriptionCreateRequestEventType = "deal.associationChange"
	SubscriptionCreateRequestEventTypeTicketAssociationChange     SubscriptionCreateRequestEventType = "ticket.associationChange"
	SubscriptionCreateRequestEventTypeLineItemAssociationChange   SubscriptionCreateRequestEventType = "line_item.associationChange"
	SubscriptionCreateRequestEventTypeObjectPropertyChange        SubscriptionCreateRequestEventType = "object.propertyChange"
	SubscriptionCreateRequestEventTypeObjectCreation              SubscriptionCreateRequestEventType = "object.creation"
	SubscriptionCreateRequestEventTypeObjectDeletion              SubscriptionCreateRequestEventType = "object.deletion"
	SubscriptionCreateRequestEventTypeObjectMerge                 SubscriptionCreateRequestEventType = "object.merge"
	SubscriptionCreateRequestEventTypeObjectRestore               SubscriptionCreateRequestEventType = "object.restore"
	SubscriptionCreateRequestEventTypeObjectAssociationChange     SubscriptionCreateRequestEventType = "object.associationChange"
)

// List of event subscriptions for your app
type SubscriptionListResponse struct {
	// List of event subscriptions for your app
	Results []SubscriptionResponse `json:"results,required"`
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

// Updated details for the subscription.
type SubscriptionPatchRequestParam struct {
	// Determines if the subscription is active or paused.
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

// Complete details for an event subscription.
type SubscriptionResponse struct {
	// The unique ID of the subscription.
	ID string `json:"id,required"`
	// Determines if the subscription is active or paused.
	Active bool `json:"active,required"`
	// When this subscription was created. Formatted as milliseconds from the
	// [Unix epoch](#).
	CreatedAt time.Time `json:"createdAt,required" format:"date-time"`
	// Type of event to listen for. Can be one of `create`, `delete`,
	// `deletedForPrivacy`, or `propertyChange`.
	//
	// Any of "contact.propertyChange", "company.propertyChange",
	// "deal.propertyChange", "ticket.propertyChange", "product.propertyChange",
	// "line_item.propertyChange", "contact.creation", "contact.deletion",
	// "contact.privacyDeletion", "company.creation", "company.deletion",
	// "deal.creation", "deal.deletion", "ticket.creation", "ticket.deletion",
	// "product.creation", "product.deletion", "line_item.creation",
	// "line_item.deletion", "conversation.creation", "conversation.deletion",
	// "conversation.newMessage", "conversation.privacyDeletion",
	// "conversation.propertyChange", "contact.merge", "company.merge", "deal.merge",
	// "ticket.merge", "product.merge", "line_item.merge", "contact.restore",
	// "company.restore", "deal.restore", "ticket.restore", "product.restore",
	// "line_item.restore", "contact.associationChange", "company.associationChange",
	// "deal.associationChange", "ticket.associationChange",
	// "line_item.associationChange", "object.propertyChange", "object.creation",
	// "object.deletion", "object.merge", "object.restore", "object.associationChange".
	EventType SubscriptionResponseEventType `json:"eventType,required"`
	// The identifier of the object type associated with the subscription.
	ObjectTypeID string `json:"objectTypeId"`
	// The internal name of the property being monitored for changes. Only applies when
	// `eventType` is `propertyChange`.
	PropertyName string `json:"propertyName"`
	// When this subscription was last updated. Formatted as milliseconds from the
	// [Unix epoch](#).
	UpdatedAt time.Time `json:"updatedAt" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		Active       respjson.Field
		CreatedAt    respjson.Field
		EventType    respjson.Field
		ObjectTypeID respjson.Field
		PropertyName respjson.Field
		UpdatedAt    respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SubscriptionResponse) RawJSON() string { return r.JSON.raw }
func (r *SubscriptionResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Type of event to listen for. Can be one of `create`, `delete`,
// `deletedForPrivacy`, or `propertyChange`.
type SubscriptionResponseEventType string

const (
	SubscriptionResponseEventTypeContactPropertyChange       SubscriptionResponseEventType = "contact.propertyChange"
	SubscriptionResponseEventTypeCompanyPropertyChange       SubscriptionResponseEventType = "company.propertyChange"
	SubscriptionResponseEventTypeDealPropertyChange          SubscriptionResponseEventType = "deal.propertyChange"
	SubscriptionResponseEventTypeTicketPropertyChange        SubscriptionResponseEventType = "ticket.propertyChange"
	SubscriptionResponseEventTypeProductPropertyChange       SubscriptionResponseEventType = "product.propertyChange"
	SubscriptionResponseEventTypeLineItemPropertyChange      SubscriptionResponseEventType = "line_item.propertyChange"
	SubscriptionResponseEventTypeContactCreation             SubscriptionResponseEventType = "contact.creation"
	SubscriptionResponseEventTypeContactDeletion             SubscriptionResponseEventType = "contact.deletion"
	SubscriptionResponseEventTypeContactPrivacyDeletion      SubscriptionResponseEventType = "contact.privacyDeletion"
	SubscriptionResponseEventTypeCompanyCreation             SubscriptionResponseEventType = "company.creation"
	SubscriptionResponseEventTypeCompanyDeletion             SubscriptionResponseEventType = "company.deletion"
	SubscriptionResponseEventTypeDealCreation                SubscriptionResponseEventType = "deal.creation"
	SubscriptionResponseEventTypeDealDeletion                SubscriptionResponseEventType = "deal.deletion"
	SubscriptionResponseEventTypeTicketCreation              SubscriptionResponseEventType = "ticket.creation"
	SubscriptionResponseEventTypeTicketDeletion              SubscriptionResponseEventType = "ticket.deletion"
	SubscriptionResponseEventTypeProductCreation             SubscriptionResponseEventType = "product.creation"
	SubscriptionResponseEventTypeProductDeletion             SubscriptionResponseEventType = "product.deletion"
	SubscriptionResponseEventTypeLineItemCreation            SubscriptionResponseEventType = "line_item.creation"
	SubscriptionResponseEventTypeLineItemDeletion            SubscriptionResponseEventType = "line_item.deletion"
	SubscriptionResponseEventTypeConversationCreation        SubscriptionResponseEventType = "conversation.creation"
	SubscriptionResponseEventTypeConversationDeletion        SubscriptionResponseEventType = "conversation.deletion"
	SubscriptionResponseEventTypeConversationNewMessage      SubscriptionResponseEventType = "conversation.newMessage"
	SubscriptionResponseEventTypeConversationPrivacyDeletion SubscriptionResponseEventType = "conversation.privacyDeletion"
	SubscriptionResponseEventTypeConversationPropertyChange  SubscriptionResponseEventType = "conversation.propertyChange"
	SubscriptionResponseEventTypeContactMerge                SubscriptionResponseEventType = "contact.merge"
	SubscriptionResponseEventTypeCompanyMerge                SubscriptionResponseEventType = "company.merge"
	SubscriptionResponseEventTypeDealMerge                   SubscriptionResponseEventType = "deal.merge"
	SubscriptionResponseEventTypeTicketMerge                 SubscriptionResponseEventType = "ticket.merge"
	SubscriptionResponseEventTypeProductMerge                SubscriptionResponseEventType = "product.merge"
	SubscriptionResponseEventTypeLineItemMerge               SubscriptionResponseEventType = "line_item.merge"
	SubscriptionResponseEventTypeContactRestore              SubscriptionResponseEventType = "contact.restore"
	SubscriptionResponseEventTypeCompanyRestore              SubscriptionResponseEventType = "company.restore"
	SubscriptionResponseEventTypeDealRestore                 SubscriptionResponseEventType = "deal.restore"
	SubscriptionResponseEventTypeTicketRestore               SubscriptionResponseEventType = "ticket.restore"
	SubscriptionResponseEventTypeProductRestore              SubscriptionResponseEventType = "product.restore"
	SubscriptionResponseEventTypeLineItemRestore             SubscriptionResponseEventType = "line_item.restore"
	SubscriptionResponseEventTypeContactAssociationChange    SubscriptionResponseEventType = "contact.associationChange"
	SubscriptionResponseEventTypeCompanyAssociationChange    SubscriptionResponseEventType = "company.associationChange"
	SubscriptionResponseEventTypeDealAssociationChange       SubscriptionResponseEventType = "deal.associationChange"
	SubscriptionResponseEventTypeTicketAssociationChange     SubscriptionResponseEventType = "ticket.associationChange"
	SubscriptionResponseEventTypeLineItemAssociationChange   SubscriptionResponseEventType = "line_item.associationChange"
	SubscriptionResponseEventTypeObjectPropertyChange        SubscriptionResponseEventType = "object.propertyChange"
	SubscriptionResponseEventTypeObjectCreation              SubscriptionResponseEventType = "object.creation"
	SubscriptionResponseEventTypeObjectDeletion              SubscriptionResponseEventType = "object.deletion"
	SubscriptionResponseEventTypeObjectMerge                 SubscriptionResponseEventType = "object.merge"
	SubscriptionResponseEventTypeObjectRestore               SubscriptionResponseEventType = "object.restore"
	SubscriptionResponseEventTypeObjectAssociationChange     SubscriptionResponseEventType = "object.associationChange"
)

// Configuration details for webhook throttling.
type ThrottlingSettings struct {
	// The maximum number of concurrent HTTP requests HubSpot will attempt to make to
	// your app.
	MaxConcurrentRequests int64 `json:"maxConcurrentRequests,required"`
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

// Configuration details for webhook throttling.
//
// The property MaxConcurrentRequests is required.
type ThrottlingSettingsParam struct {
	// The maximum number of concurrent HTTP requests HubSpot will attempt to make to
	// your app.
	MaxConcurrentRequests int64 `json:"maxConcurrentRequests,required"`
	paramObj
}

func (r ThrottlingSettingsParam) MarshalJSON() (data []byte, err error) {
	type shadow ThrottlingSettingsParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ThrottlingSettingsParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
