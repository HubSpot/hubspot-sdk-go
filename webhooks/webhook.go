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
// with the hubspot API.
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
	// Any of "CANCELED", "COMPLETE", "PENDING", "PROCESSING".
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
	BatchResponseSubscriptionResponseStatusCanceled   BatchResponseSubscriptionResponseStatus = "CANCELED"
	BatchResponseSubscriptionResponseStatusComplete   BatchResponseSubscriptionResponseStatus = "COMPLETE"
	BatchResponseSubscriptionResponseStatusPending    BatchResponseSubscriptionResponseStatus = "PENDING"
	BatchResponseSubscriptionResponseStatusProcessing BatchResponseSubscriptionResponseStatus = "PROCESSING"
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
	// Any of "company.associationChange", "company.creation", "company.deletion",
	// "company.merge", "company.propertyChange", "company.restore",
	// "contact.associationChange", "contact.creation", "contact.deletion",
	// "contact.merge", "contact.privacyDeletion", "contact.propertyChange",
	// "contact.restore", "conversation.creation", "conversation.deletion",
	// "conversation.newMessage", "conversation.privacyDeletion",
	// "conversation.propertyChange", "deal.associationChange", "deal.creation",
	// "deal.deletion", "deal.merge", "deal.propertyChange", "deal.restore",
	// "line_item.associationChange", "line_item.creation", "line_item.deletion",
	// "line_item.merge", "line_item.propertyChange", "line_item.restore",
	// "object.associationChange", "object.creation", "object.deletion",
	// "object.merge", "object.propertyChange", "object.restore", "product.creation",
	// "product.deletion", "product.merge", "product.propertyChange",
	// "product.restore", "ticket.associationChange", "ticket.creation",
	// "ticket.deletion", "ticket.merge", "ticket.propertyChange", "ticket.restore".
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
	// Any of "company.associationChange", "company.creation", "company.deletion",
	// "company.merge", "company.propertyChange", "company.restore",
	// "contact.associationChange", "contact.creation", "contact.deletion",
	// "contact.merge", "contact.privacyDeletion", "contact.propertyChange",
	// "contact.restore", "conversation.creation", "conversation.deletion",
	// "conversation.newMessage", "conversation.privacyDeletion",
	// "conversation.propertyChange", "deal.associationChange", "deal.creation",
	// "deal.deletion", "deal.merge", "deal.propertyChange", "deal.restore",
	// "line_item.associationChange", "line_item.creation", "line_item.deletion",
	// "line_item.merge", "line_item.propertyChange", "line_item.restore",
	// "object.associationChange", "object.creation", "object.deletion",
	// "object.merge", "object.propertyChange", "object.restore", "product.creation",
	// "product.deletion", "product.merge", "product.propertyChange",
	// "product.restore", "ticket.associationChange", "ticket.creation",
	// "ticket.deletion", "ticket.merge", "ticket.propertyChange", "ticket.restore".
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
