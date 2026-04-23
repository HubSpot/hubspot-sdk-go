// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package scheduler

import (
	"context"
	"net/http"
	"net/url"
	"slices"

	"github.com/stainless-sdks/hubspot-sdk-go/internal/apijson"
	"github.com/stainless-sdks/hubspot-sdk-go/internal/apiquery"
	shimjson "github.com/stainless-sdks/hubspot-sdk-go/internal/encoding/json"
	"github.com/stainless-sdks/hubspot-sdk-go/internal/requestconfig"
	"github.com/stainless-sdks/hubspot-sdk-go/option"
)

// MeetingAdvancedService contains methods and other services that help with
// interacting with the hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMeetingAdvancedService] method instead.
type MeetingAdvancedService struct {
	options []option.RequestOption
}

// NewMeetingAdvancedService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewMeetingAdvancedService(opts ...option.RequestOption) (r MeetingAdvancedService) {
	r = MeetingAdvancedService{}
	r.options = opts
	return
}

// Create a new calendar event and meeting object by providing the necessary
// details such as associations, email reminders, meeting object properties, and
// timezone.
func (r *MeetingAdvancedService) New(ctx context.Context, params MeetingAdvancedNewParams, opts ...option.RequestOption) (res *ExternalCalenderMeetingEventResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "scheduler/2026-03/meetings/calendar"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Book a meeting for a specified meeting page.
func (r *MeetingAdvancedService) Book(ctx context.Context, body MeetingAdvancedBookParams, opts ...option.RequestOption) (res *ExternalMeetingBookingResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "scheduler/2026-03/meetings/meeting-links/book"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type MeetingAdvancedNewParams struct {
	OrganizerUserID                           string `query:"organizerUserId" api:"required" json:"-"`
	ExternalCalendarMeetingEventCreateRequest ExternalCalendarMeetingEventCreateRequestParam
	paramObj
}

func (r MeetingAdvancedNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.ExternalCalendarMeetingEventCreateRequest)
}
func (r *MeetingAdvancedNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [MeetingAdvancedNewParams]'s query parameters as
// `url.Values`.
func (r MeetingAdvancedNewParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type MeetingAdvancedBookParams struct {
	ExternalMeetingBooking ExternalMeetingBookingParam
	paramObj
}

func (r MeetingAdvancedBookParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.ExternalMeetingBooking)
}
func (r *MeetingAdvancedBookParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
