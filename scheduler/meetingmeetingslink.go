// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package scheduler

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

// MeetingMeetingsLinkService contains methods and other services that help with
// interacting with the Hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMeetingMeetingsLinkService] method instead.
type MeetingMeetingsLinkService struct {
	Options []option.RequestOption
}

// NewMeetingMeetingsLinkService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewMeetingMeetingsLinkService(opts ...option.RequestOption) (r MeetingMeetingsLinkService) {
	r = MeetingMeetingsLinkService{}
	r.Options = opts
	return
}

// Get a paged list meeting scheduling pages
func (r *MeetingMeetingsLinkService) List(ctx context.Context, opts ...option.RequestOption) (res *CollectionResponseWithTotalExternalLinkMetadataForwardPaging, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "scheduler/v3/meetings/meeting-links"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Book a meeting for a specified meeting page.
func (r *MeetingMeetingsLinkService) Book(ctx context.Context, body MeetingMeetingsLinkBookParams, opts ...option.RequestOption) (res *ExternalMeetingBookingResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "scheduler/v3/meetings/meeting-links/book"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get the next availability times for a meeting page.
func (r *MeetingMeetingsLinkService) GetAvailabilityBySlug(ctx context.Context, slug string, opts ...option.RequestOption) (res *ExternalLinkAvailabilityAndBusyTimes, err error) {
	opts = slices.Concat(r.Options, opts)
	if slug == "" {
		err = errors.New("missing required slug parameter")
		return
	}
	path := fmt.Sprintf("scheduler/v3/meetings/meeting-links/book/availability-page/%s", slug)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Get details about the initial information necessary for a meeting scheduler.
func (r *MeetingMeetingsLinkService) GetBookingInfoBySlug(ctx context.Context, slug string, opts ...option.RequestOption) (res *ExternalBookingInfo, err error) {
	opts = slices.Concat(r.Options, opts)
	if slug == "" {
		err = errors.New("missing required slug parameter")
		return
	}
	path := fmt.Sprintf("scheduler/v3/meetings/meeting-links/book/%s", slug)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type MeetingMeetingsLinkBookParams struct {
	ExternalMeetingBooking ExternalMeetingBookingParam
	paramObj
}

func (r MeetingMeetingsLinkBookParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.ExternalMeetingBooking)
}
func (r *MeetingMeetingsLinkBookParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.ExternalMeetingBooking)
}
