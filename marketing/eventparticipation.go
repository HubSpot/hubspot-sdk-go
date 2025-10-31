// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package marketing

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
	"github.com/stainless-sdks/hubspot-sdk-go/packages/param"
)

// EventParticipationService contains methods and other services that help with
// interacting with the Hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEventParticipationService] method instead.
type EventParticipationService struct {
	Options []option.RequestOption
}

// NewEventParticipationService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewEventParticipationService(opts ...option.RequestOption) (r EventParticipationService) {
	r = EventParticipationService{}
	r.Options = opts
	return
}

// Read Marketing event's participations counters by externalAccountId and
// externalEventId pair.
func (r *EventParticipationService) GetByExternalAccountAndEventID(ctx context.Context, externalEventID string, query EventParticipationGetByExternalAccountAndEventIDParams, opts ...option.RequestOption) (res *AttendanceCounters, err error) {
	opts = slices.Concat(r.Options, opts)
	if query.ExternalAccountID == "" {
		err = errors.New("missing required externalAccountId parameter")
		return
	}
	if externalEventID == "" {
		err = errors.New("missing required externalEventId parameter")
		return
	}
	path := fmt.Sprintf("marketing/v3/marketing-events/participations/%s/%s", query.ExternalAccountID, externalEventID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Read Marketing event's participations counters by internal identifier
// marketingEventId.
func (r *EventParticipationService) GetByID(ctx context.Context, marketingEventID int64, opts ...option.RequestOption) (res *AttendanceCounters, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("marketing/v3/marketing-events/participations/%v", marketingEventID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Read Contact's participations by identifier - email or internal id.
func (r *EventParticipationService) ListBreakdownByContact(ctx context.Context, contactIdentifier string, query EventParticipationListBreakdownByContactParams, opts ...option.RequestOption) (res *CollectionResponseWithTotalParticipationBreakdownForwardPaging, err error) {
	opts = slices.Concat(r.Options, opts)
	if contactIdentifier == "" {
		err = errors.New("missing required contactIdentifier parameter")
		return
	}
	path := fmt.Sprintf("marketing/v3/marketing-events/participations/contacts/%s/breakdown", contactIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Read Marketing event's participations breakdown with optional filters by
// externalAccountId and externalEventId pair.
func (r *EventParticipationService) ListBreakdownByExternalAccountAndEventID(ctx context.Context, externalEventID string, params EventParticipationListBreakdownByExternalAccountAndEventIDParams, opts ...option.RequestOption) (res *CollectionResponseWithTotalParticipationBreakdownForwardPaging, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.ExternalAccountID == "" {
		err = errors.New("missing required externalAccountId parameter")
		return
	}
	if externalEventID == "" {
		err = errors.New("missing required externalEventId parameter")
		return
	}
	path := fmt.Sprintf("marketing/v3/marketing-events/participations/%s/%s/breakdown", params.ExternalAccountID, externalEventID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return
}

// Read Marketing event's participations breakdown with optional filters by
// internal identifier marketingEventId.
func (r *EventParticipationService) ListBreakdownByID(ctx context.Context, marketingEventID int64, query EventParticipationListBreakdownByIDParams, opts ...option.RequestOption) (res *CollectionResponseWithTotalParticipationBreakdownForwardPaging, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("marketing/v3/marketing-events/participations/%v/breakdown", marketingEventID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type EventParticipationGetByExternalAccountAndEventIDParams struct {
	ExternalAccountID string `path:"externalAccountId,required" json:"-"`
	paramObj
}

type EventParticipationListBreakdownByContactParams struct {
	// The cursor indicating the position of the last retrieved item.
	After param.Opt[string] `query:"after,omitzero" json:"-"`
	// The limit for response size. The default value is 10, the max number is 100
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// The participation state value. It may be REGISTERED, CANCELLED, ATTENDED,
	// NO_SHOW
	State param.Opt[string] `query:"state,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [EventParticipationListBreakdownByContactParams]'s query
// parameters as `url.Values`.
func (r EventParticipationListBreakdownByContactParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type EventParticipationListBreakdownByExternalAccountAndEventIDParams struct {
	ExternalAccountID string `path:"externalAccountId,required" json:"-"`
	// The cursor indicating the position of the last retrieved item.
	After param.Opt[string] `query:"after,omitzero" json:"-"`
	// The identifier of the Contact. It may be email or internal id.
	ContactIdentifier param.Opt[string] `query:"contactIdentifier,omitzero" json:"-"`
	// The limit for response size. The default value is 10, the max number is 100
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// The participation state value. It may be REGISTERED, CANCELLED, ATTENDED,
	// NO_SHOW
	State param.Opt[string] `query:"state,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes
// [EventParticipationListBreakdownByExternalAccountAndEventIDParams]'s query
// parameters as `url.Values`.
func (r EventParticipationListBreakdownByExternalAccountAndEventIDParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type EventParticipationListBreakdownByIDParams struct {
	// The cursor indicating the position of the last retrieved item.
	After param.Opt[string] `query:"after,omitzero" json:"-"`
	// The identifier of the Contact. It may be email or internal id.
	ContactIdentifier param.Opt[string] `query:"contactIdentifier,omitzero" json:"-"`
	// The limit for response size. The default value is 10, the max number is 100
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// The participation state value. It may be REGISTERED, CANCELLED, ATTENDED,
	// NO_SHOW
	State param.Opt[string] `query:"state,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [EventParticipationListBreakdownByIDParams]'s query
// parameters as `url.Values`.
func (r EventParticipationListBreakdownByIDParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
