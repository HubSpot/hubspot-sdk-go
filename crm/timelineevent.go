// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package crm

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"slices"

	shimjson "github.com/stainless-sdks/hubspot-sdk-go/internal/encoding/json"
	"github.com/stainless-sdks/hubspot-sdk-go/internal/requestconfig"
	"github.com/stainless-sdks/hubspot-sdk-go/option"
)

// TimelineEventService contains methods and other services that help with
// interacting with the Hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTimelineEventService] method instead.
type TimelineEventService struct {
	Options []option.RequestOption
}

// NewTimelineEventService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewTimelineEventService(opts ...option.RequestOption) (r TimelineEventService) {
	r = TimelineEventService{}
	r.Options = opts
	return
}

// Send a single instance of event data to a specified event type.
func (r *TimelineEventService) New(ctx context.Context, body TimelineEventNewParams, opts ...option.RequestOption) (res *TimelineEventResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "integrators/timeline/v3/events"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Batch create multiple instances of timeline events based on an event template.
// Once created, these event are immutable on the object timeline and cannot be
// modified. If the event template was configured to update object properties via
// `objectPropertyName`, this call will also attempt to updates those properties,
// or add them if they don't exist.
func (r *TimelineEventService) BatchNew(ctx context.Context, body TimelineEventBatchNewParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "integrators/timeline/v3/events/batch/create"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Retrieve an event instance, specified by template ID and event ID.
func (r *TimelineEventService) Get(ctx context.Context, eventID string, query TimelineEventGetParams, opts ...option.RequestOption) (res *TimelineEventResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if query.EventTemplateID == "" {
		err = errors.New("missing required eventTemplateId parameter")
		return
	}
	if eventID == "" {
		err = errors.New("missing required eventId parameter")
		return
	}
	path := fmt.Sprintf("integrators/timeline/v3/events/%s/%s", query.EventTemplateID, eventID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Retrieve details for a specific event, specified by template ID and event ID.
func (r *TimelineEventService) GetDetail(ctx context.Context, eventID string, query TimelineEventGetDetailParams, opts ...option.RequestOption) (res *EventDetail, err error) {
	opts = slices.Concat(r.Options, opts)
	if query.EventTemplateID == "" {
		err = errors.New("missing required eventTemplateId parameter")
		return
	}
	if eventID == "" {
		err = errors.New("missing required eventId parameter")
		return
	}
	path := fmt.Sprintf("integrators/timeline/v3/events/%s/%s/detail", query.EventTemplateID, eventID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type TimelineEventNewParams struct {
	// The state of the timeline event.
	TimelineEvent TimelineEventParam
	paramObj
}

func (r TimelineEventNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.TimelineEvent)
}
func (r *TimelineEventNewParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.TimelineEvent)
}

type TimelineEventBatchNewParams struct {
	// Used to create timeline events in batches.
	BatchInputTimelineEvent BatchInputTimelineEventParam
	paramObj
}

func (r TimelineEventBatchNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.BatchInputTimelineEvent)
}
func (r *TimelineEventBatchNewParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.BatchInputTimelineEvent)
}

type TimelineEventGetParams struct {
	EventTemplateID string `path:"eventTemplateId,required" json:"-"`
	paramObj
}

type TimelineEventGetDetailParams struct {
	EventTemplateID string `path:"eventTemplateId,required" json:"-"`
	paramObj
}
