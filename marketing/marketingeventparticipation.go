// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package marketing

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/HubSpot/hubspot-sdk-go/internal/apiquery"
	"github.com/HubSpot/hubspot-sdk-go/internal/requestconfig"
	"github.com/HubSpot/hubspot-sdk-go/option"
	"github.com/HubSpot/hubspot-sdk-go/packages/pagination"
	"github.com/HubSpot/hubspot-sdk-go/packages/param"
)

// MarketingEventParticipationService contains methods and other services that help
// with interacting with the hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMarketingEventParticipationService] method instead.
type MarketingEventParticipationService struct {
	options []option.RequestOption
}

// NewMarketingEventParticipationService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewMarketingEventParticipationService(opts ...option.RequestOption) (r MarketingEventParticipationService) {
	r = MarketingEventParticipationService{}
	r.options = opts
	return
}

// Read Marketing event's participations counters by externalAccountId and
// externalEventId pair.
func (r *MarketingEventParticipationService) GetByExternalAccountAndEventID(ctx context.Context, externalEventID string, query MarketingEventParticipationGetByExternalAccountAndEventIDParams, opts ...option.RequestOption) (res *AttendanceCounters, err error) {
	opts = slices.Concat(r.options, opts)
	if query.ExternalAccountID == "" {
		err = errors.New("missing required externalAccountId parameter")
		return nil, err
	}
	if externalEventID == "" {
		err = errors.New("missing required externalEventId parameter")
		return nil, err
	}
	path := fmt.Sprintf("marketing/marketing-events/2026-03/participations/%s/%s", url.PathEscape(query.ExternalAccountID), url.PathEscape(externalEventID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Read Marketing event's participations counters by internal identifier
// marketingEventId.
func (r *MarketingEventParticipationService) GetByID(ctx context.Context, marketingEventID int64, opts ...option.RequestOption) (res *AttendanceCounters, err error) {
	opts = slices.Concat(r.options, opts)
	path := fmt.Sprintf("marketing/marketing-events/2026-03/participations/%v", marketingEventID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Read Contact's participations by identifier - email or internal id.
func (r *MarketingEventParticipationService) ListBreakdownByContact(ctx context.Context, contactIdentifier string, query MarketingEventParticipationListBreakdownByContactParams, opts ...option.RequestOption) (res *pagination.Page[ParticipationBreakdown], err error) {
	var raw *http.Response
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if contactIdentifier == "" {
		err = errors.New("missing required contactIdentifier parameter")
		return nil, err
	}
	path := fmt.Sprintf("marketing/marketing-events/2026-03/participations/contacts/%s/breakdown", url.PathEscape(contactIdentifier))
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

// Read Contact's participations by identifier - email or internal id.
func (r *MarketingEventParticipationService) ListBreakdownByContactAutoPaging(ctx context.Context, contactIdentifier string, query MarketingEventParticipationListBreakdownByContactParams, opts ...option.RequestOption) *pagination.PageAutoPager[ParticipationBreakdown] {
	return pagination.NewPageAutoPager(r.ListBreakdownByContact(ctx, contactIdentifier, query, opts...))
}

// Read Marketing event's participations breakdown with optional filters by
// externalAccountId and externalEventId pair.
func (r *MarketingEventParticipationService) ListBreakdownByExternalAccountAndEventID(ctx context.Context, externalEventID string, params MarketingEventParticipationListBreakdownByExternalAccountAndEventIDParams, opts ...option.RequestOption) (res *pagination.Page[ParticipationBreakdown], err error) {
	var raw *http.Response
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.ExternalAccountID == "" {
		err = errors.New("missing required externalAccountId parameter")
		return nil, err
	}
	if externalEventID == "" {
		err = errors.New("missing required externalEventId parameter")
		return nil, err
	}
	path := fmt.Sprintf("marketing/marketing-events/2026-03/participations/%s/%s/breakdown", url.PathEscape(params.ExternalAccountID), url.PathEscape(externalEventID))
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, params, &res, opts...)
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

// Read Marketing event's participations breakdown with optional filters by
// externalAccountId and externalEventId pair.
func (r *MarketingEventParticipationService) ListBreakdownByExternalAccountAndEventIDAutoPaging(ctx context.Context, externalEventID string, params MarketingEventParticipationListBreakdownByExternalAccountAndEventIDParams, opts ...option.RequestOption) *pagination.PageAutoPager[ParticipationBreakdown] {
	return pagination.NewPageAutoPager(r.ListBreakdownByExternalAccountAndEventID(ctx, externalEventID, params, opts...))
}

// Read Marketing event's participations breakdown with optional filters by
// internal identifier marketingEventId.
func (r *MarketingEventParticipationService) ListBreakdownByID(ctx context.Context, marketingEventID int64, query MarketingEventParticipationListBreakdownByIDParams, opts ...option.RequestOption) (res *pagination.Page[ParticipationBreakdown], err error) {
	var raw *http.Response
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := fmt.Sprintf("marketing/marketing-events/2026-03/participations/%v/breakdown", marketingEventID)
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

// Read Marketing event's participations breakdown with optional filters by
// internal identifier marketingEventId.
func (r *MarketingEventParticipationService) ListBreakdownByIDAutoPaging(ctx context.Context, marketingEventID int64, query MarketingEventParticipationListBreakdownByIDParams, opts ...option.RequestOption) *pagination.PageAutoPager[ParticipationBreakdown] {
	return pagination.NewPageAutoPager(r.ListBreakdownByID(ctx, marketingEventID, query, opts...))
}

type MarketingEventParticipationGetByExternalAccountAndEventIDParams struct {
	ExternalAccountID string `path:"externalAccountId" api:"required" json:"-"`
	paramObj
}

type MarketingEventParticipationListBreakdownByContactParams struct {
	// The cursor indicating the position of the last retrieved item.
	After param.Opt[string] `query:"after,omitzero" json:"-"`
	// The limit for response size. The default value is 10, the max number is 100
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// The participation state value. It may be REGISTERED, CANCELLED, ATTENDED,
	// NO_SHOW
	State param.Opt[string] `query:"state,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [MarketingEventParticipationListBreakdownByContactParams]'s
// query parameters as `url.Values`.
func (r MarketingEventParticipationListBreakdownByContactParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type MarketingEventParticipationListBreakdownByExternalAccountAndEventIDParams struct {
	ExternalAccountID string `path:"externalAccountId" api:"required" json:"-"`
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
// [MarketingEventParticipationListBreakdownByExternalAccountAndEventIDParams]'s
// query parameters as `url.Values`.
func (r MarketingEventParticipationListBreakdownByExternalAccountAndEventIDParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type MarketingEventParticipationListBreakdownByIDParams struct {
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

// URLQuery serializes [MarketingEventParticipationListBreakdownByIDParams]'s query
// parameters as `url.Values`.
func (r MarketingEventParticipationListBreakdownByIDParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
