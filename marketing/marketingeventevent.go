// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package marketing

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/HubSpot/hubspot-sdk-go/internal/apijson"
	"github.com/HubSpot/hubspot-sdk-go/internal/apiquery"
	shimjson "github.com/HubSpot/hubspot-sdk-go/internal/encoding/json"
	"github.com/HubSpot/hubspot-sdk-go/internal/requestconfig"
	"github.com/HubSpot/hubspot-sdk-go/option"
)

// MarketingEventEventService contains methods and other services that help with
// interacting with the hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMarketingEventEventService] method instead.
type MarketingEventEventService struct {
	options []option.RequestOption
}

// NewMarketingEventEventService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewMarketingEventEventService(opts ...option.RequestOption) (r MarketingEventEventService) {
	r = MarketingEventEventService{}
	r.options = opts
	return
}

// Mark a marketing event as cancelled.
func (r *MarketingEventEventService) CancelByExternalEventID(ctx context.Context, externalEventID string, body MarketingEventEventCancelByExternalEventIDParams, opts ...option.RequestOption) (res *MarketingEventDefaultResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if externalEventID == "" {
		err = errors.New("missing required externalEventId parameter")
		return nil, err
	}
	path := fmt.Sprintf("marketing/marketing-events/2026-03/events/%s/cancel", url.PathEscape(externalEventID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Mark a marketing event as completed
func (r *MarketingEventEventService) CompleteByExternalEventID(ctx context.Context, externalEventID string, params MarketingEventEventCompleteByExternalEventIDParams, opts ...option.RequestOption) (res *MarketingEventDefaultResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if externalEventID == "" {
		err = errors.New("missing required externalEventId parameter")
		return nil, err
	}
	path := fmt.Sprintf("marketing/marketing-events/2026-03/events/%s/complete", url.PathEscape(externalEventID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

type MarketingEventEventCancelByExternalEventIDParams struct {
	ExternalAccountID string `query:"externalAccountId" api:"required" json:"-"`
	paramObj
}

// URLQuery serializes [MarketingEventEventCancelByExternalEventIDParams]'s query
// parameters as `url.Values`.
func (r MarketingEventEventCancelByExternalEventIDParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type MarketingEventEventCompleteByExternalEventIDParams struct {
	ExternalAccountID                   string `query:"externalAccountId" api:"required" json:"-"`
	MarketingEventCompleteRequestParams MarketingEventCompleteRequestParams
	paramObj
}

func (r MarketingEventEventCompleteByExternalEventIDParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.MarketingEventCompleteRequestParams)
}
func (r *MarketingEventEventCompleteByExternalEventIDParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [MarketingEventEventCompleteByExternalEventIDParams]'s query
// parameters as `url.Values`.
func (r MarketingEventEventCompleteByExternalEventIDParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
