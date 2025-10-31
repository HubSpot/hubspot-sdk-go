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
)

// ExtensionVideoConferencingSettingService contains methods and other services
// that help with interacting with the Hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewExtensionVideoConferencingSettingService] method instead.
type ExtensionVideoConferencingSettingService struct {
	Options []option.RequestOption
}

// NewExtensionVideoConferencingSettingService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewExtensionVideoConferencingSettingService(opts ...option.RequestOption) (r ExtensionVideoConferencingSettingService) {
	r = ExtensionVideoConferencingSettingService{}
	r.Options = opts
	return
}

// Updates the settings for a video conference application with the specified ID.
func (r *ExtensionVideoConferencingSettingService) Update(ctx context.Context, appID int64, body ExtensionVideoConferencingSettingUpdateParams, opts ...option.RequestOption) (res *ExternalSettings, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("crm/v3/extensions/videoconferencing/settings/%v", appID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Deletes the settings for a video conference application with the specified ID.
func (r *ExtensionVideoConferencingSettingService) Delete(ctx context.Context, appID int64, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := fmt.Sprintf("crm/v3/extensions/videoconferencing/settings/%v", appID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Return the settings for a video conference application with the specified ID.
func (r *ExtensionVideoConferencingSettingService) Get(ctx context.Context, appID int64, opts ...option.RequestOption) (res *ExternalSettings, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("crm/v3/extensions/videoconferencing/settings/%v", appID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type ExtensionVideoConferencingSettingUpdateParams struct {
	// The URLs of the various actions provided by the video conferencing application.
	// All URLs must use the `https` protocol.
	ExternalSettings ExternalSettingsParam
	paramObj
}

func (r ExtensionVideoConferencingSettingUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.ExternalSettings)
}
func (r *ExtensionVideoConferencingSettingUpdateParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.ExternalSettings)
}
