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

// ExtensionCallingRecordingSettingService contains methods and other services that
// help with interacting with the hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewExtensionCallingRecordingSettingService] method instead.
type ExtensionCallingRecordingSettingService struct {
	Options []option.RequestOption
}

// NewExtensionCallingRecordingSettingService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewExtensionCallingRecordingSettingService(opts ...option.RequestOption) (r ExtensionCallingRecordingSettingService) {
	r = ExtensionCallingRecordingSettingService{}
	r.Options = opts
	return
}

func (r *ExtensionCallingRecordingSettingService) New(ctx context.Context, appID int64, body ExtensionCallingRecordingSettingNewParams, opts ...option.RequestOption) (res *RecordingSettingsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("crm/v3/extensions/calling/%v/settings/recording", appID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

func (r *ExtensionCallingRecordingSettingService) Update(ctx context.Context, appID int64, body ExtensionCallingRecordingSettingUpdateParams, opts ...option.RequestOption) (res *RecordingSettingsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("crm/v3/extensions/calling/%v/settings/recording", appID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

func (r *ExtensionCallingRecordingSettingService) Get(ctx context.Context, appID int64, opts ...option.RequestOption) (res *RecordingSettingsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("crm/v3/extensions/calling/%v/settings/recording", appID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

func (r *ExtensionCallingRecordingSettingService) MarkReady(ctx context.Context, body ExtensionCallingRecordingSettingMarkReadyParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "crm/v3/extensions/calling/recordings/ready"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

type ExtensionCallingRecordingSettingNewParams struct {
	RecordingSettingsRequest RecordingSettingsRequestParam
	paramObj
}

func (r ExtensionCallingRecordingSettingNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.RecordingSettingsRequest)
}
func (r *ExtensionCallingRecordingSettingNewParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.RecordingSettingsRequest)
}

type ExtensionCallingRecordingSettingUpdateParams struct {
	RecordingSettingsPatchRequest RecordingSettingsPatchRequestParam
	paramObj
}

func (r ExtensionCallingRecordingSettingUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.RecordingSettingsPatchRequest)
}
func (r *ExtensionCallingRecordingSettingUpdateParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.RecordingSettingsPatchRequest)
}

type ExtensionCallingRecordingSettingMarkReadyParams struct {
	MarkRecordingAsReadyRequest MarkRecordingAsReadyRequestParam
	paramObj
}

func (r ExtensionCallingRecordingSettingMarkReadyParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.MarkRecordingAsReadyRequest)
}
func (r *ExtensionCallingRecordingSettingMarkReadyParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.MarkRecordingAsReadyRequest)
}
