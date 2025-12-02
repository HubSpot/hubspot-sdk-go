// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package scheduler

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/stainless-sdks/hubspot-sdk-go/internal/apiquery"
	shimjson "github.com/stainless-sdks/hubspot-sdk-go/internal/encoding/json"
	"github.com/stainless-sdks/hubspot-sdk-go/internal/requestconfig"
	"github.com/stainless-sdks/hubspot-sdk-go/option"
	"github.com/stainless-sdks/hubspot-sdk-go/packages/pagination"
	"github.com/stainless-sdks/hubspot-sdk-go/packages/param"
)

// MeetingMeetingsLinkService contains methods and other services that help with
// interacting with the hubspot API.
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
func (r *MeetingMeetingsLinkService) List(ctx context.Context, query MeetingMeetingsLinkListParams, opts ...option.RequestOption) (res *pagination.Page[ExternalLinkMetadata], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "scheduler/v3/meetings/meeting-links"
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// Get a paged list meeting scheduling pages
func (r *MeetingMeetingsLinkService) ListAutoPaging(ctx context.Context, query MeetingMeetingsLinkListParams, opts ...option.RequestOption) *pagination.PageAutoPager[ExternalLinkMetadata] {
	return pagination.NewPageAutoPager(r.List(ctx, query, opts...))
}

// Book a meeting for a specified meeting page.
func (r *MeetingMeetingsLinkService) Book(ctx context.Context, body MeetingMeetingsLinkBookParams, opts ...option.RequestOption) (res *ExternalMeetingBookingResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "scheduler/v3/meetings/meeting-links/book"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get the next availability times for a meeting page.
func (r *MeetingMeetingsLinkService) GetAvailabilityBySlug(ctx context.Context, slug string, query MeetingMeetingsLinkGetAvailabilityBySlugParams, opts ...option.RequestOption) (res *ExternalLinkAvailabilityAndBusyTimes, err error) {
	opts = slices.Concat(r.Options, opts)
	if slug == "" {
		err = errors.New("missing required slug parameter")
		return
	}
	path := fmt.Sprintf("scheduler/v3/meetings/meeting-links/book/availability-page/%s", slug)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Get details about the initial information necessary for a meeting scheduler.
func (r *MeetingMeetingsLinkService) GetBookingInfoBySlug(ctx context.Context, slug string, query MeetingMeetingsLinkGetBookingInfoBySlugParams, opts ...option.RequestOption) (res *ExternalBookingInfo, err error) {
	opts = slices.Concat(r.Options, opts)
	if slug == "" {
		err = errors.New("missing required slug parameter")
		return
	}
	path := fmt.Sprintf("scheduler/v3/meetings/meeting-links/book/%s", slug)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type MeetingMeetingsLinkListParams struct {
	// The paging cursor token of the last successfully read resource will be returned
	// as the `paging.next.after` JSON property of a paged response containing more
	// results.
	After param.Opt[string] `query:"after,omitzero" json:"-"`
	// The maximum number of results to display per page.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Retrieve scheduling pages with a specified name.
	Name param.Opt[string] `query:"name,omitzero" json:"-"`
	// Filter the response to scheduling pages created by the specified user.
	OrganizerUserID param.Opt[string] `query:"organizerUserId,omitzero" json:"-"`
	// Filter the response to the specific type of meeting.
	Type param.Opt[string] `query:"type,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [MeetingMeetingsLinkListParams]'s query parameters as
// `url.Values`.
func (r MeetingMeetingsLinkListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
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

type MeetingMeetingsLinkGetAvailabilityBySlugParams struct {
	// Return times in response based on specified time zone.
	Timezone string `query:"timezone,required" json:"-"`
	// Get times for a different month.
	MonthOffset param.Opt[int64] `query:"monthOffset,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [MeetingMeetingsLinkGetAvailabilityBySlugParams]'s query
// parameters as `url.Values`.
func (r MeetingMeetingsLinkGetAvailabilityBySlugParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type MeetingMeetingsLinkGetBookingInfoBySlugParams struct {
	// Return times in response based on specified time zone.
	Timezone string `query:"timezone,required" json:"-"`
	paramObj
}

// URLQuery serializes [MeetingMeetingsLinkGetBookingInfoBySlugParams]'s query
// parameters as `url.Values`.
func (r MeetingMeetingsLinkGetBookingInfoBySlugParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
