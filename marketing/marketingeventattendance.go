// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package marketing

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/stainless-sdks/hubspot-sdk-go/internal/apijson"
	"github.com/stainless-sdks/hubspot-sdk-go/internal/apiquery"
	shimjson "github.com/stainless-sdks/hubspot-sdk-go/internal/encoding/json"
	"github.com/stainless-sdks/hubspot-sdk-go/internal/requestconfig"
	"github.com/stainless-sdks/hubspot-sdk-go/option"
	"github.com/stainless-sdks/hubspot-sdk-go/packages/param"
)

// MarketingEventAttendanceService contains methods and other services that help
// with interacting with the hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMarketingEventAttendanceService] method instead.
type MarketingEventAttendanceService struct {
	options []option.RequestOption
}

// NewMarketingEventAttendanceService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewMarketingEventAttendanceService(opts ...option.RequestOption) (r MarketingEventAttendanceService) {
	r = MarketingEventAttendanceService{}
	r.options = opts
	return
}

// Records the participation of multiple HubSpot contacts in a Marketing Event
// using their HubSpot contact IDs.
//
// Additional Functionality:
//
// - Adds a timeline event to the contacts.
//
// Allowed Properties: For the state "attend":
//
// - joinedAt
// - leftAt
func (r *MarketingEventAttendanceService) NewByEventIDAndContactID(ctx context.Context, subscriberState string, params MarketingEventAttendanceNewByEventIDAndContactIDParams, opts ...option.RequestOption) (res *BatchResponseSubscriberVidResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if params.ObjectID == "" {
		err = errors.New("missing required objectId parameter")
		return nil, err
	}
	if subscriberState == "" {
		err = errors.New("missing required subscriberState parameter")
		return nil, err
	}
	path := fmt.Sprintf("marketing/marketing-events/2026-03/%s/attendance/%s/create", url.PathEscape(params.ObjectID), url.PathEscape(subscriberState))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Records the participation of multiple HubSpot contacts in a Marketing Event
// using their email addresses.
//
// If a contact does not exist, it will be automatically created. The
// contactProperties field is used exclusively for creating new contacts and will
// not update properties of existing contacts.
//
// Additional Functionality:
//
// - Adds a timeline event to the contacts.
//
// Allowed Properties: For the state "attend":
//
// - joinedAt
// - leftAt
func (r *MarketingEventAttendanceService) NewByEventIDAndEmail(ctx context.Context, subscriberState string, params MarketingEventAttendanceNewByEventIDAndEmailParams, opts ...option.RequestOption) (res *BatchResponseSubscriberEmailResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if params.ObjectID == "" {
		err = errors.New("missing required objectId parameter")
		return nil, err
	}
	if subscriberState == "" {
		err = errors.New("missing required subscriberState parameter")
		return nil, err
	}
	path := fmt.Sprintf("marketing/marketing-events/2026-03/%s/attendance/%s/email-create", url.PathEscape(params.ObjectID), url.PathEscape(subscriberState))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Records the participation of multiple HubSpot contacts in a Marketing Event
// using their HubSpot contact IDs.
//
// Additional Functionality:
//
// - Adds a timeline event to the contacts.
//
// Allowed Properties: For the state "attend":
//
// - joinedAt
// - leftAt
func (r *MarketingEventAttendanceService) NewByExternalEventIDAndContactID(ctx context.Context, subscriberState string, params MarketingEventAttendanceNewByExternalEventIDAndContactIDParams, opts ...option.RequestOption) (res *BatchResponseSubscriberVidResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if params.ExternalEventID == "" {
		err = errors.New("missing required externalEventId parameter")
		return nil, err
	}
	if subscriberState == "" {
		err = errors.New("missing required subscriberState parameter")
		return nil, err
	}
	path := fmt.Sprintf("marketing/marketing-events/2026-03/attendance/%s/%s/create", url.PathEscape(params.ExternalEventID), url.PathEscape(subscriberState))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Records the participation of multiple HubSpot contacts in a Marketing Event
// using their email addresses.
//
// If a contact does not exist, it will be automatically created. The
// contactProperties field is used exclusively for creating new contacts and will
// not update properties of existing contacts.
//
// Additional Functionality:
//
// - Adds a timeline event to the contacts.
//
// Allowed Properties: For the state "attend":
//
// - joinedAt
// - leftAt
func (r *MarketingEventAttendanceService) NewByExternalEventIDAndEmail(ctx context.Context, subscriberState string, params MarketingEventAttendanceNewByExternalEventIDAndEmailParams, opts ...option.RequestOption) (res *BatchResponseSubscriberEmailResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if params.ExternalEventID == "" {
		err = errors.New("missing required externalEventId parameter")
		return nil, err
	}
	if subscriberState == "" {
		err = errors.New("missing required subscriberState parameter")
		return nil, err
	}
	path := fmt.Sprintf("marketing/marketing-events/2026-03/attendance/%s/%s/email-create", url.PathEscape(params.ExternalEventID), url.PathEscape(subscriberState))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

type MarketingEventAttendanceNewByEventIDAndContactIDParams struct {
	ObjectID                           string `path:"objectId" api:"required" json:"-"`
	BatchInputMarketingEventSubscriber BatchInputMarketingEventSubscriberParam
	paramObj
}

func (r MarketingEventAttendanceNewByEventIDAndContactIDParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.BatchInputMarketingEventSubscriber)
}
func (r *MarketingEventAttendanceNewByEventIDAndContactIDParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MarketingEventAttendanceNewByEventIDAndEmailParams struct {
	ObjectID                                string `path:"objectId" api:"required" json:"-"`
	BatchInputMarketingEventEmailSubscriber BatchInputMarketingEventEmailSubscriberParam
	paramObj
}

func (r MarketingEventAttendanceNewByEventIDAndEmailParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.BatchInputMarketingEventEmailSubscriber)
}
func (r *MarketingEventAttendanceNewByEventIDAndEmailParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MarketingEventAttendanceNewByExternalEventIDAndContactIDParams struct {
	ExternalEventID                    string `path:"externalEventId" api:"required" json:"-"`
	BatchInputMarketingEventSubscriber BatchInputMarketingEventSubscriberParam
	ExternalAccountID                  param.Opt[string] `query:"externalAccountId,omitzero" json:"-"`
	paramObj
}

func (r MarketingEventAttendanceNewByExternalEventIDAndContactIDParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.BatchInputMarketingEventSubscriber)
}
func (r *MarketingEventAttendanceNewByExternalEventIDAndContactIDParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes
// [MarketingEventAttendanceNewByExternalEventIDAndContactIDParams]'s query
// parameters as `url.Values`.
func (r MarketingEventAttendanceNewByExternalEventIDAndContactIDParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type MarketingEventAttendanceNewByExternalEventIDAndEmailParams struct {
	ExternalEventID                         string `path:"externalEventId" api:"required" json:"-"`
	BatchInputMarketingEventEmailSubscriber BatchInputMarketingEventEmailSubscriberParam
	ExternalAccountID                       param.Opt[string] `query:"externalAccountId,omitzero" json:"-"`
	paramObj
}

func (r MarketingEventAttendanceNewByExternalEventIDAndEmailParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.BatchInputMarketingEventEmailSubscriber)
}
func (r *MarketingEventAttendanceNewByExternalEventIDAndEmailParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes
// [MarketingEventAttendanceNewByExternalEventIDAndEmailParams]'s query parameters
// as `url.Values`.
func (r MarketingEventAttendanceNewByExternalEventIDAndEmailParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
