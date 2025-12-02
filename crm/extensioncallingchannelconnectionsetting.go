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

// ExtensionCallingChannelConnectionSettingService contains methods and other
// services that help with interacting with the hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewExtensionCallingChannelConnectionSettingService] method instead.
type ExtensionCallingChannelConnectionSettingService struct {
	Options []option.RequestOption
}

// NewExtensionCallingChannelConnectionSettingService generates a new service that
// applies the given options to each request. These options are applied after the
// parent client's options (if there is one), and before any request-specific
// options.
func NewExtensionCallingChannelConnectionSettingService(opts ...option.RequestOption) (r ExtensionCallingChannelConnectionSettingService) {
	r = ExtensionCallingChannelConnectionSettingService{}
	r.Options = opts
	return
}

func (r *ExtensionCallingChannelConnectionSettingService) New(ctx context.Context, appID int64, body ExtensionCallingChannelConnectionSettingNewParams, opts ...option.RequestOption) (res *ChannelConnectionSettingsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("crm/v3/extensions/calling/%v/settings/channel-connection", appID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

func (r *ExtensionCallingChannelConnectionSettingService) Update(ctx context.Context, appID int64, body ExtensionCallingChannelConnectionSettingUpdateParams, opts ...option.RequestOption) (res *ChannelConnectionSettingsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("crm/v3/extensions/calling/%v/settings/channel-connection", appID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

func (r *ExtensionCallingChannelConnectionSettingService) Delete(ctx context.Context, appID int64, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := fmt.Sprintf("crm/v3/extensions/calling/%v/settings/channel-connection", appID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

func (r *ExtensionCallingChannelConnectionSettingService) Get(ctx context.Context, appID int64, opts ...option.RequestOption) (res *ChannelConnectionSettingsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("crm/v3/extensions/calling/%v/settings/channel-connection", appID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type ExtensionCallingChannelConnectionSettingNewParams struct {
	ChannelConnectionSettingsRequest ChannelConnectionSettingsRequestParam
	paramObj
}

func (r ExtensionCallingChannelConnectionSettingNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.ChannelConnectionSettingsRequest)
}
func (r *ExtensionCallingChannelConnectionSettingNewParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.ChannelConnectionSettingsRequest)
}

type ExtensionCallingChannelConnectionSettingUpdateParams struct {
	ChannelConnectionSettingsPatchRequest ChannelConnectionSettingsPatchRequestParam
	paramObj
}

func (r ExtensionCallingChannelConnectionSettingUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.ChannelConnectionSettingsPatchRequest)
}
func (r *ExtensionCallingChannelConnectionSettingUpdateParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.ChannelConnectionSettingsPatchRequest)
}
