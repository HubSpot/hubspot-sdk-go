// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package marketing

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"slices"

	shimjson "github.com/stainless-sdks/hubspot-sdk-go/internal/encoding/json"
	"github.com/stainless-sdks/hubspot-sdk-go/internal/requestconfig"
	"github.com/stainless-sdks/hubspot-sdk-go/option"
)

// EventSettingService contains methods and other services that help with
// interacting with the Hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEventSettingService] method instead.
type EventSettingService struct {
	Options []option.RequestOption
}

// NewEventSettingService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewEventSettingService(opts ...option.RequestOption) (r EventSettingService) {
	r = EventSettingService{}
	r.Options = opts
	return
}

// Create or update the current settings for the application.
func (r *EventSettingService) NewOrUpdate(ctx context.Context, appID int64, body EventSettingNewOrUpdateParams, opts ...option.RequestOption) (res *EventDetailSettings, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("marketing/v3/marketing-events/%v/settings", appID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieve the current settings for the application.
func (r *EventSettingService) Get(ctx context.Context, appID int64, opts ...option.RequestOption) (res *EventDetailSettings, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("marketing/v3/marketing-events/%v/settings", appID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type EventSettingNewOrUpdateParams struct {
	EventDetailSettingsURL EventDetailSettingsURLParam
	paramObj
}

func (r EventSettingNewOrUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.EventDetailSettingsURL)
}
func (r *EventSettingNewOrUpdateParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.EventDetailSettingsURL)
}
