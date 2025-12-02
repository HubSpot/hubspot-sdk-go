// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cms

import (
	"context"
	"encoding/json"
	"net/http"
	"slices"

	shimjson "github.com/stainless-sdks/hubspot-sdk-go/internal/encoding/json"
	"github.com/stainless-sdks/hubspot-sdk-go/internal/requestconfig"
	"github.com/stainless-sdks/hubspot-sdk-go/option"
)

// MediaBridgeEventService contains methods and other services that help with
// interacting with the hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMediaBridgeEventService] method instead.
type MediaBridgeEventService struct {
	Options []option.RequestOption
}

// NewMediaBridgeEventService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewMediaBridgeEventService(opts ...option.RequestOption) (r MediaBridgeEventService) {
	r = MediaBridgeEventService{}
	r.Options = opts
	return
}

// Create an event containing the viewers attention span details for the media.
func (r *MediaBridgeEventService) NewAttentionSpanEvent(ctx context.Context, body MediaBridgeEventNewAttentionSpanEventParams, opts ...option.RequestOption) (res *AttentionSpanEvent, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "media-bridge/v1/events/attention-span"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Create an event for when a user begins playing a piece of media.
func (r *MediaBridgeEventService) NewMediaPlayedEvent(ctx context.Context, body MediaBridgeEventNewMediaPlayedEventParams, opts ...option.RequestOption) (res *MediaPlayedEvent, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "media-bridge/v1/events/media-played"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Create an event representing a user reaching quarterly milestones in a piece of
// media they're viewing.
func (r *MediaBridgeEventService) NewMediaPlayedPercentEvent(ctx context.Context, body MediaBridgeEventNewMediaPlayedPercentEventParams, opts ...option.RequestOption) (res *MediaPlayedPercentageEvent, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "media-bridge/v1/events/media-played-percent"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type MediaBridgeEventNewAttentionSpanEventParams struct {
	AttentionSpanEventRequest AttentionSpanEventRequestParam
	paramObj
}

func (r MediaBridgeEventNewAttentionSpanEventParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.AttentionSpanEventRequest)
}
func (r *MediaBridgeEventNewAttentionSpanEventParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.AttentionSpanEventRequest)
}

type MediaBridgeEventNewMediaPlayedEventParams struct {
	MediaPlayedEventRequest MediaPlayedEventRequestParam
	paramObj
}

func (r MediaBridgeEventNewMediaPlayedEventParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.MediaPlayedEventRequest)
}
func (r *MediaBridgeEventNewMediaPlayedEventParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.MediaPlayedEventRequest)
}

type MediaBridgeEventNewMediaPlayedPercentEventParams struct {
	MediaPlayedPercentageEventRequest MediaPlayedPercentageEventRequestParam
	paramObj
}

func (r MediaBridgeEventNewMediaPlayedPercentEventParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.MediaPlayedPercentageEventRequest)
}
func (r *MediaBridgeEventNewMediaPlayedPercentEventParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.MediaPlayedPercentageEventRequest)
}
