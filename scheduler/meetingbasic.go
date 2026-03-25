// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package scheduler

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/stainless-sdks/hubspot-sdk-go/internal/apiquery"
	"github.com/stainless-sdks/hubspot-sdk-go/internal/requestconfig"
	"github.com/stainless-sdks/hubspot-sdk-go/option"
	"github.com/stainless-sdks/hubspot-sdk-go/packages/pagination"
	"github.com/stainless-sdks/hubspot-sdk-go/packages/param"
)

// MeetingBasicService contains methods and other services that help with
// interacting with the hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMeetingBasicService] method instead.
type MeetingBasicService struct {
	Options []option.RequestOption
}

// NewMeetingBasicService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewMeetingBasicService(opts ...option.RequestOption) (r MeetingBasicService) {
	r = MeetingBasicService{}
	r.Options = opts
	return
}

// Get a paged list meeting scheduling pages
func (r *MeetingBasicService) List(ctx context.Context, query MeetingBasicListParams, opts ...option.RequestOption) (res *pagination.Page[ExternalLinkMetadata], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "scheduler/2026-03/meetings/meeting-links"
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
func (r *MeetingBasicService) ListAutoPaging(ctx context.Context, query MeetingBasicListParams, opts ...option.RequestOption) *pagination.PageAutoPager[ExternalLinkMetadata] {
	return pagination.NewPageAutoPager(r.List(ctx, query, opts...))
}

// Get the next availability times for a meeting page.
func (r *MeetingBasicService) GetAvailabilityBySlug(ctx context.Context, slug string, query MeetingBasicGetAvailabilityBySlugParams, opts ...option.RequestOption) (res *ExternalLinkAvailabilityAndBusyTimes, err error) {
	opts = slices.Concat(r.Options, opts)
	if slug == "" {
		err = errors.New("missing required slug parameter")
		return nil, err
	}
	path := fmt.Sprintf("scheduler/2026-03/meetings/meeting-links/book/availability-page/%s", slug)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Get details about the initial information necessary for a meeting scheduler.
func (r *MeetingBasicService) GetBookingInfoBySlug(ctx context.Context, slug string, query MeetingBasicGetBookingInfoBySlugParams, opts ...option.RequestOption) (res *ExternalBookingInfo, err error) {
	opts = slices.Concat(r.Options, opts)
	if slug == "" {
		err = errors.New("missing required slug parameter")
		return nil, err
	}
	path := fmt.Sprintf("scheduler/2026-03/meetings/meeting-links/book/%s", slug)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

type MeetingBasicListParams struct {
	// The paging cursor token of the last successfully read resource will be returned
	// as the `paging.next.after` JSON property of a paged response containing more
	// results.
	After param.Opt[string] `query:"after,omitzero" json:"-"`
	// The maximum number of results to display per page.
	Limit           param.Opt[int64]  `query:"limit,omitzero" json:"-"`
	Name            param.Opt[string] `query:"name,omitzero" json:"-"`
	OrganizerUserID param.Opt[string] `query:"organizerUserId,omitzero" json:"-"`
	// Any of "GROUP_CALENDAR", "PERSONAL_LINK", "ROUND_ROBIN_CALENDAR".
	Type MeetingBasicListParamsType `query:"type,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [MeetingBasicListParams]'s query parameters as `url.Values`.
func (r MeetingBasicListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type MeetingBasicListParamsType string

const (
	MeetingBasicListParamsTypeGroupCalendar      MeetingBasicListParamsType = "GROUP_CALENDAR"
	MeetingBasicListParamsTypePersonalLink       MeetingBasicListParamsType = "PERSONAL_LINK"
	MeetingBasicListParamsTypeRoundRobinCalendar MeetingBasicListParamsType = "ROUND_ROBIN_CALENDAR"
)

type MeetingBasicGetAvailabilityBySlugParams struct {
	Timezone    string           `query:"timezone" api:"required" json:"-"`
	MonthOffset param.Opt[int64] `query:"monthOffset,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [MeetingBasicGetAvailabilityBySlugParams]'s query parameters
// as `url.Values`.
func (r MeetingBasicGetAvailabilityBySlugParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type MeetingBasicGetBookingInfoBySlugParams struct {
	Timezone string `query:"timezone" api:"required" json:"-"`
	paramObj
}

// URLQuery serializes [MeetingBasicGetBookingInfoBySlugParams]'s query parameters
// as `url.Values`.
func (r MeetingBasicGetBookingInfoBySlugParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
