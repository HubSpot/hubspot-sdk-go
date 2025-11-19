// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package events

import (
	"context"
	"encoding/json"
	"net/http"
	"slices"
	"time"

	"github.com/stainless-sdks/hubspot-sdk-go/internal/apijson"
	shimjson "github.com/stainless-sdks/hubspot-sdk-go/internal/encoding/json"
	"github.com/stainless-sdks/hubspot-sdk-go/internal/requestconfig"
	"github.com/stainless-sdks/hubspot-sdk-go/option"
	"github.com/stainless-sdks/hubspot-sdk-go/packages/param"
)

// SendService contains methods and other services that help with interacting with
// the Hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSendService] method instead.
type SendService struct {
	Options []option.RequestOption
}

// NewSendService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewSendService(opts ...option.RequestOption) (r SendService) {
	r = SendService{}
	r.Options = opts
	return
}

// Send data for a single event completion.
func (r *SendService) Send(ctx context.Context, body SendSendParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "events/v3/send"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Send multiple event completions at once.
func (r *SendService) SendBatch(ctx context.Context, body SendSendBatchParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "events/v3/send/batch"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// The property Inputs is required.
type BatchedBehavioralEventHTTPCompletionRequestParam struct {
	Inputs []BehavioralEventHTTPCompletionRequestParam `json:"inputs,omitzero,required"`
	paramObj
}

func (r BatchedBehavioralEventHTTPCompletionRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow BatchedBehavioralEventHTTPCompletionRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BatchedBehavioralEventHTTPCompletionRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property EventName is required.
type BehavioralEventHTTPCompletionRequestParam struct {
	// The internal name of the event (`pe<portalID>_eventName`). Can be retrieved
	// through the
	// [event definitions API](https://developers.hubspot.com/docs/reference/api/analytics-and-events/custom-events/custom-event-definitions#get-%2Fevents%2Fv3%2Fevent-definitions)
	// or in
	// [HubSpot's UI](https://knowledge.hubspot.com/reports/create-custom-behavioral-events-with-the-code-wizard#find-internal-name).
	EventName string `json:"eventName,required"`
	// The visitor's email address. Used for associating the event data with a CRM
	// record.
	Email param.Opt[string] `json:"email,omitzero"`
	// The ID of the object that completed the event (e.g., contact ID or visitor ID).
	ObjectID param.Opt[string] `json:"objectId,omitzero"`
	// The time when this event occurred. If this isn't set, the current time will be
	// used.
	OccurredAt param.Opt[time.Time] `json:"occurredAt,omitzero" format:"date-time"`
	// The visitor's usertoken. Used for associating the event data with a CRM record.
	Utk param.Opt[string] `json:"utk,omitzero"`
	// Include a universally unique identifier to assign a unique ID to the event
	// completion. Can be useful for matching data between HubSpot and other external
	// systems.
	Uuid param.Opt[string] `json:"uuid,omitzero"`
	// The event properties to update. Takes the format of key-value pairs (property
	// internal name and property value). Learn more about
	// [HubSpot's default event properties](https://developers.hubspot.com/docs/guides/api/analytics-and-events/custom-events/custom-event-definitions#hubspot-s-default-event-properties).
	Properties map[string]string `json:"properties,omitzero"`
	paramObj
}

func (r BehavioralEventHTTPCompletionRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow BehavioralEventHTTPCompletionRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BehavioralEventHTTPCompletionRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SendSendParams struct {
	BehavioralEventHTTPCompletionRequest BehavioralEventHTTPCompletionRequestParam
	paramObj
}

func (r SendSendParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.BehavioralEventHTTPCompletionRequest)
}
func (r *SendSendParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.BehavioralEventHTTPCompletionRequest)
}

type SendSendBatchParams struct {
	BatchedBehavioralEventHTTPCompletionRequest BatchedBehavioralEventHTTPCompletionRequestParam
	paramObj
}

func (r SendSendBatchParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.BatchedBehavioralEventHTTPCompletionRequest)
}
func (r *SendSendBatchParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.BatchedBehavioralEventHTTPCompletionRequest)
}
