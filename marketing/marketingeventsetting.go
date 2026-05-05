// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package marketing

import (
	"context"
	"fmt"
	"net/http"
	"slices"

	"github.com/HubSpot/hubspot-sdk-go/internal/apijson"
	shimjson "github.com/HubSpot/hubspot-sdk-go/internal/encoding/json"
	"github.com/HubSpot/hubspot-sdk-go/internal/requestconfig"
	"github.com/HubSpot/hubspot-sdk-go/option"
)

// MarketingEventSettingService contains methods and other services that help with
// interacting with the hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMarketingEventSettingService] method instead.
type MarketingEventSettingService struct {
	options []option.RequestOption
}

// NewMarketingEventSettingService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewMarketingEventSettingService(opts ...option.RequestOption) (r MarketingEventSettingService) {
	r = MarketingEventSettingService{}
	r.options = opts
	return
}

// Create or update the current settings for the application.
func (r *MarketingEventSettingService) NewOrUpdate(ctx context.Context, appID int64, body MarketingEventSettingNewOrUpdateParams, opts ...option.RequestOption) (res *EventDetailSettings, err error) {
	opts = slices.Concat(r.options, opts)
	path := fmt.Sprintf("marketing/marketing-events/2026-03/%v/settings", appID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Retrieve the current settings for the application.
func (r *MarketingEventSettingService) Get(ctx context.Context, appID int64, opts ...option.RequestOption) (res *EventDetailSettings, err error) {
	opts = slices.Concat(r.options, opts)
	path := fmt.Sprintf("marketing/marketing-events/2026-03/%v/settings", appID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type MarketingEventSettingNewOrUpdateParams struct {
	EventDetailSettingsURL EventDetailSettingsURLParam
	paramObj
}

func (r MarketingEventSettingNewOrUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.EventDetailSettingsURL)
}
func (r *MarketingEventSettingNewOrUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
