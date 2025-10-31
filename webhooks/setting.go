// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package webhooks

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

// SettingService contains methods and other services that help with interacting
// with the Hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSettingService] method instead.
type SettingService struct {
	Options []option.RequestOption
}

// NewSettingService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewSettingService(opts ...option.RequestOption) (r SettingService) {
	r = SettingService{}
	r.Options = opts
	return
}

// Update webhook settings for the specified app.
func (r *SettingService) Update(ctx context.Context, appID int64, body SettingUpdateParams, opts ...option.RequestOption) (res *SettingsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("webhooks/v3/%v/settings", appID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Retrieve the webhook settings for the specified app, including the webhook’s
// target URL, throttle configuration, and create/update date.
func (r *SettingService) List(ctx context.Context, appID int64, opts ...option.RequestOption) (res *SettingsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("webhooks/v3/%v/settings", appID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Delete the webhook settings for the specified app. Event subscriptions will not
// be deleted, but will be paused until another webhook is created.
func (r *SettingService) Delete(ctx context.Context, appID int64, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := fmt.Sprintf("webhooks/v3/%v/settings", appID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

type SettingUpdateParams struct {
	// New or updated webhook settings for an app.
	SettingsChangeRequest SettingsChangeRequestParam
	paramObj
}

func (r SettingUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.SettingsChangeRequest)
}
func (r *SettingUpdateParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.SettingsChangeRequest)
}
