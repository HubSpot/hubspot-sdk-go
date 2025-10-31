// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package scheduler

import (
	"context"
	"encoding/json"
	"net/http"
	"slices"

	shimjson "github.com/stainless-sdks/hubspot-sdk-go/internal/encoding/json"
	"github.com/stainless-sdks/hubspot-sdk-go/internal/requestconfig"
	"github.com/stainless-sdks/hubspot-sdk-go/option"
)

// MeetingCalendarService contains methods and other services that help with
// interacting with the Hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMeetingCalendarService] method instead.
type MeetingCalendarService struct {
	Options []option.RequestOption
}

// NewMeetingCalendarService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewMeetingCalendarService(opts ...option.RequestOption) (r MeetingCalendarService) {
	r = MeetingCalendarService{}
	r.Options = opts
	return
}

func (r *MeetingCalendarService) New(ctx context.Context, body MeetingCalendarNewParams, opts ...option.RequestOption) (res *ExternalCalenderMeetingEventResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "scheduler/v3/meetings/calendar"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type MeetingCalendarNewParams struct {
	ExternalCalendarMeetingEventCreateRequest ExternalCalendarMeetingEventCreateRequestParam
	paramObj
}

func (r MeetingCalendarNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.ExternalCalendarMeetingEventCreateRequest)
}
func (r *MeetingCalendarNewParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.ExternalCalendarMeetingEventCreateRequest)
}
