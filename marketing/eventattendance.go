// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package marketing

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
	"github.com/stainless-sdks/hubspot-sdk-go/packages/param"
)

// EventAttendanceService contains methods and other services that help with
// interacting with the Hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEventAttendanceService] method instead.
type EventAttendanceService struct {
	Options []option.RequestOption
}

// NewEventAttendanceService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewEventAttendanceService(opts ...option.RequestOption) (r EventAttendanceService) {
	r = EventAttendanceService{}
	r.Options = opts
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
func (r *EventAttendanceService) NewByEventIDAndContactID(ctx context.Context, subscriberState string, params EventAttendanceNewByEventIDAndContactIDParams, opts ...option.RequestOption) (res *BatchResponseSubscriberVidResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.ObjectID == "" {
		err = errors.New("missing required objectId parameter")
		return
	}
	if subscriberState == "" {
		err = errors.New("missing required subscriberState parameter")
		return
	}
	path := fmt.Sprintf("marketing/v3/marketing-events/%s/attendance/%s/create", params.ObjectID, subscriberState)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
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
func (r *EventAttendanceService) NewByEventIDAndEmail(ctx context.Context, subscriberState string, params EventAttendanceNewByEventIDAndEmailParams, opts ...option.RequestOption) (res *BatchResponseSubscriberEmailResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.ObjectID == "" {
		err = errors.New("missing required objectId parameter")
		return
	}
	if subscriberState == "" {
		err = errors.New("missing required subscriberState parameter")
		return
	}
	path := fmt.Sprintf("marketing/v3/marketing-events/%s/attendance/%s/email-create", params.ObjectID, subscriberState)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
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
func (r *EventAttendanceService) NewByExternalEventIDAndContactID(ctx context.Context, subscriberState string, params EventAttendanceNewByExternalEventIDAndContactIDParams, opts ...option.RequestOption) (res *BatchResponseSubscriberVidResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.ExternalEventID == "" {
		err = errors.New("missing required externalEventId parameter")
		return
	}
	if subscriberState == "" {
		err = errors.New("missing required subscriberState parameter")
		return
	}
	path := fmt.Sprintf("marketing/v3/marketing-events/attendance/%s/%s/create", params.ExternalEventID, subscriberState)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
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
func (r *EventAttendanceService) NewByExternalEventIDAndEmail(ctx context.Context, subscriberState string, params EventAttendanceNewByExternalEventIDAndEmailParams, opts ...option.RequestOption) (res *BatchResponseSubscriberEmailResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.ExternalEventID == "" {
		err = errors.New("missing required externalEventId parameter")
		return
	}
	if subscriberState == "" {
		err = errors.New("missing required subscriberState parameter")
		return
	}
	path := fmt.Sprintf("marketing/v3/marketing-events/attendance/%s/%s/email-create", params.ExternalEventID, subscriberState)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

type EventAttendanceNewByEventIDAndContactIDParams struct {
	ObjectID                           string `path:"objectId,required" json:"-"`
	BatchInputMarketingEventSubscriber BatchInputMarketingEventSubscriberParam
	paramObj
}

func (r EventAttendanceNewByEventIDAndContactIDParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.BatchInputMarketingEventSubscriber)
}
func (r *EventAttendanceNewByEventIDAndContactIDParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.BatchInputMarketingEventSubscriber)
}

type EventAttendanceNewByEventIDAndEmailParams struct {
	ObjectID                                string `path:"objectId,required" json:"-"`
	BatchInputMarketingEventEmailSubscriber BatchInputMarketingEventEmailSubscriberParam
	paramObj
}

func (r EventAttendanceNewByEventIDAndEmailParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.BatchInputMarketingEventEmailSubscriber)
}
func (r *EventAttendanceNewByEventIDAndEmailParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.BatchInputMarketingEventEmailSubscriber)
}

type EventAttendanceNewByExternalEventIDAndContactIDParams struct {
	ExternalEventID                    string `path:"externalEventId,required" json:"-"`
	BatchInputMarketingEventSubscriber BatchInputMarketingEventSubscriberParam
	// The accountId that is associated with this marketing event in the external event
	// application
	ExternalAccountID param.Opt[string] `query:"externalAccountId,omitzero" json:"-"`
	paramObj
}

func (r EventAttendanceNewByExternalEventIDAndContactIDParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.BatchInputMarketingEventSubscriber)
}
func (r *EventAttendanceNewByExternalEventIDAndContactIDParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.BatchInputMarketingEventSubscriber)
}

// URLQuery serializes [EventAttendanceNewByExternalEventIDAndContactIDParams]'s
// query parameters as `url.Values`.
func (r EventAttendanceNewByExternalEventIDAndContactIDParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type EventAttendanceNewByExternalEventIDAndEmailParams struct {
	ExternalEventID                         string `path:"externalEventId,required" json:"-"`
	BatchInputMarketingEventEmailSubscriber BatchInputMarketingEventEmailSubscriberParam
	// The accountId that is associated with this marketing event in the external event
	// application
	ExternalAccountID param.Opt[string] `query:"externalAccountId,omitzero" json:"-"`
	paramObj
}

func (r EventAttendanceNewByExternalEventIDAndEmailParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.BatchInputMarketingEventEmailSubscriber)
}
func (r *EventAttendanceNewByExternalEventIDAndEmailParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.BatchInputMarketingEventEmailSubscriber)
}

// URLQuery serializes [EventAttendanceNewByExternalEventIDAndEmailParams]'s query
// parameters as `url.Values`.
func (r EventAttendanceNewByExternalEventIDAndEmailParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
