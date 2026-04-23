// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package events

import (
	"time"

	"github.com/stainless-sdks/hubspot-sdk-go/internal/apijson"
	"github.com/stainless-sdks/hubspot-sdk-go/option"
	"github.com/stainless-sdks/hubspot-sdk-go/packages/param"
)

// EventService contains methods and other services that help with interacting with
// the hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEventService] method instead.
type EventService struct {
	options     []option.RequestOption
	Definitions DefinitionService
	Occurrences OccurrenceService
	Send        SendService
}

// NewEventService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewEventService(opts ...option.RequestOption) (r EventService) {
	r = EventService{}
	r.options = opts
	r.Definitions = NewDefinitionService(opts...)
	r.Occurrences = NewOccurrenceService(opts...)
	r.Send = NewSendService(opts...)
	return
}

// The property Inputs is required.
type BatchedBehavioralEventHTTPCompletionRequestParam struct {
	Inputs []BehavioralEventHTTPCompletionRequestParam `json:"inputs,omitzero" api:"required"`
	paramObj
}

func (r BatchedBehavioralEventHTTPCompletionRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow BatchedBehavioralEventHTTPCompletionRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BatchedBehavioralEventHTTPCompletionRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties EventName, Properties are required.
type BehavioralEventHTTPCompletionRequestParam struct {
	// Internal name of the event-type to trigger
	EventName string `json:"eventName" api:"required"`
	// Map of properties for the event in the format property internal name - property
	// value
	Properties map[string]string `json:"properties,omitzero" api:"required"`
	// Email of visitor
	Email param.Opt[string] `json:"email,omitzero"`
	// The object id that this event occurred on. Could be a contact id or a visitor
	// id.
	ObjectID param.Opt[string] `json:"objectId,omitzero"`
	// The time when this event occurred (if any). If this isn't set, the current time
	// will be used
	OccurredAt param.Opt[time.Time] `json:"occurredAt,omitzero" format:"date-time"`
	// User token
	Utk  param.Opt[string] `json:"utk,omitzero"`
	Uuid param.Opt[string] `json:"uuid,omitzero"`
	paramObj
}

func (r BehavioralEventHTTPCompletionRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow BehavioralEventHTTPCompletionRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BehavioralEventHTTPCompletionRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
