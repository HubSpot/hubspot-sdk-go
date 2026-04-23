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
)

// MarketingEventSubscriberStateService contains methods and other services that
// help with interacting with the hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMarketingEventSubscriberStateService] method instead.
type MarketingEventSubscriberStateService struct {
	options []option.RequestOption
}

// NewMarketingEventSubscriberStateService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewMarketingEventSubscriberStateService(opts ...option.RequestOption) (r MarketingEventSubscriberStateService) {
	r = MarketingEventSubscriberStateService{}
	r.options = opts
	return
}

// Record a subscriber state between multiple HubSpot contacts and a marketing
// event, using contact email addresses. Note that the contact must already exist
// in HubSpot; a contact will not be created. The contactProperties field is used
// only when creating a new contact. These properties will not update existing
// contacts.
func (r *MarketingEventSubscriberStateService) RecordByEmail(ctx context.Context, subscriberState string, params MarketingEventSubscriberStateRecordByEmailParams, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if params.ExternalEventID == "" {
		err = errors.New("missing required externalEventId parameter")
		return nil, err
	}
	if subscriberState == "" {
		err = errors.New("missing required subscriberState parameter")
		return nil, err
	}
	path := fmt.Sprintf("marketing/marketing-events/2026-03/events/%s/%s/email-upsert", url.PathEscape(params.ExternalEventID), url.PathEscape(subscriberState))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Record a subscriber state between multiple HubSpot contacts and a marketing
// event, using HubSpot contact IDs. Note that the contact must already exist in
// HubSpot; a contact will not be created.
func (r *MarketingEventSubscriberStateService) RecordByID(ctx context.Context, subscriberState string, params MarketingEventSubscriberStateRecordByIDParams, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if params.ExternalEventID == "" {
		err = errors.New("missing required externalEventId parameter")
		return nil, err
	}
	if subscriberState == "" {
		err = errors.New("missing required subscriberState parameter")
		return nil, err
	}
	path := fmt.Sprintf("marketing/marketing-events/2026-03/events/%s/%s/upsert", url.PathEscape(params.ExternalEventID), url.PathEscape(subscriberState))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

type MarketingEventSubscriberStateRecordByEmailParams struct {
	ExternalEventID                         string `path:"externalEventId" api:"required" json:"-"`
	ExternalAccountID                       string `query:"externalAccountId" api:"required" json:"-"`
	BatchInputMarketingEventEmailSubscriber BatchInputMarketingEventEmailSubscriberParam
	paramObj
}

func (r MarketingEventSubscriberStateRecordByEmailParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.BatchInputMarketingEventEmailSubscriber)
}
func (r *MarketingEventSubscriberStateRecordByEmailParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [MarketingEventSubscriberStateRecordByEmailParams]'s query
// parameters as `url.Values`.
func (r MarketingEventSubscriberStateRecordByEmailParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type MarketingEventSubscriberStateRecordByIDParams struct {
	ExternalEventID                    string `path:"externalEventId" api:"required" json:"-"`
	ExternalAccountID                  string `query:"externalAccountId" api:"required" json:"-"`
	BatchInputMarketingEventSubscriber BatchInputMarketingEventSubscriberParam
	paramObj
}

func (r MarketingEventSubscriberStateRecordByIDParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.BatchInputMarketingEventSubscriber)
}
func (r *MarketingEventSubscriberStateRecordByIDParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [MarketingEventSubscriberStateRecordByIDParams]'s query
// parameters as `url.Values`.
func (r MarketingEventSubscriberStateRecordByIDParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
