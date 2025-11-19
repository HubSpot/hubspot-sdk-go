// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package crm

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"slices"

	shimjson "github.com/stainless-sdks/hubspot-sdk-go/internal/encoding/json"
	"github.com/stainless-sdks/hubspot-sdk-go/internal/requestconfig"
	"github.com/stainless-sdks/hubspot-sdk-go/option"
	"github.com/stainless-sdks/hubspot-sdk-go/webhooks"
)

// ExtensionCallingSettingService contains methods and other services that help
// with interacting with the Hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewExtensionCallingSettingService] method instead.
type ExtensionCallingSettingService struct {
	Options []option.RequestOption
}

// NewExtensionCallingSettingService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewExtensionCallingSettingService(opts ...option.RequestOption) (r ExtensionCallingSettingService) {
	r = ExtensionCallingSettingService{}
	r.Options = opts
	return
}

// Set the menu label, target iframe URL, and dimensions for your calling
// extension.
func (r *ExtensionCallingSettingService) New(ctx context.Context, appID int64, body ExtensionCallingSettingNewParams, opts ...option.RequestOption) (res *webhooks.SettingsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("crm/v3/extensions/calling/%v/settings", appID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Update existing calling extension settings.
func (r *ExtensionCallingSettingService) Update(ctx context.Context, appID int64, body ExtensionCallingSettingUpdateParams, opts ...option.RequestOption) (res *webhooks.SettingsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("crm/v3/extensions/calling/%v/settings", appID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Delete a calling extension. This will remove your service as an option for all
// connected accounts.
func (r *ExtensionCallingSettingService) Delete(ctx context.Context, appID int64, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := fmt.Sprintf("crm/v3/extensions/calling/%v/settings", appID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Retrieve the settings configured for the app.
func (r *ExtensionCallingSettingService) Get(ctx context.Context, appID int64, opts ...option.RequestOption) (res *webhooks.SettingsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("crm/v3/extensions/calling/%v/settings", appID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type ExtensionCallingSettingNewParams struct {
	SettingsRequest SettingsRequestParam
	paramObj
}

func (r ExtensionCallingSettingNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.SettingsRequest)
}
func (r *ExtensionCallingSettingNewParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.SettingsRequest)
}

type ExtensionCallingSettingUpdateParams struct {
	SettingsPatchRequest SettingsPatchRequestParam
	paramObj
}

func (r ExtensionCallingSettingUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.SettingsPatchRequest)
}
func (r *ExtensionCallingSettingUpdateParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.SettingsPatchRequest)
}
