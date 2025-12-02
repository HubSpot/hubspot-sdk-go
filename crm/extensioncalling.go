// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package crm

import (
	"time"

	"github.com/stainless-sdks/hubspot-sdk-go/internal/apijson"
	"github.com/stainless-sdks/hubspot-sdk-go/option"
	"github.com/stainless-sdks/hubspot-sdk-go/packages/param"
	"github.com/stainless-sdks/hubspot-sdk-go/packages/respjson"
)

// ExtensionCallingService contains methods and other services that help with
// interacting with the hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewExtensionCallingService] method instead.
type ExtensionCallingService struct {
	Options                   []option.RequestOption
	ChannelConnectionSettings ExtensionCallingChannelConnectionSettingService
	RecordingSettings         ExtensionCallingRecordingSettingService
	Settings                  ExtensionCallingSettingService
	Transcripts               ExtensionCallingTranscriptService
}

// NewExtensionCallingService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewExtensionCallingService(opts ...option.RequestOption) (r ExtensionCallingService) {
	r = ExtensionCallingService{}
	r.Options = opts
	r.ChannelConnectionSettings = NewExtensionCallingChannelConnectionSettingService(opts...)
	r.RecordingSettings = NewExtensionCallingRecordingSettingService(opts...)
	r.Settings = NewExtensionCallingSettingService(opts...)
	r.Transcripts = NewExtensionCallingTranscriptService(opts...)
	return
}

type ChannelConnectionSettingsPatchRequestParam struct {
	IsReady param.Opt[bool]   `json:"isReady,omitzero"`
	URL     param.Opt[string] `json:"url,omitzero"`
	paramObj
}

func (r ChannelConnectionSettingsPatchRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow ChannelConnectionSettingsPatchRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ChannelConnectionSettingsPatchRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties IsReady, URL are required.
type ChannelConnectionSettingsRequestParam struct {
	IsReady bool   `json:"isReady,required"`
	URL     string `json:"url,required"`
	paramObj
}

func (r ChannelConnectionSettingsRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow ChannelConnectionSettingsRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ChannelConnectionSettingsRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChannelConnectionSettingsResponse struct {
	CreatedAt time.Time `json:"createdAt,required" format:"date-time"`
	IsReady   bool      `json:"isReady,required"`
	UpdatedAt time.Time `json:"updatedAt,required" format:"date-time"`
	URL       string    `json:"url,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreatedAt   respjson.Field
		IsReady     respjson.Field
		UpdatedAt   respjson.Field
		URL         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChannelConnectionSettingsResponse) RawJSON() string { return r.JSON.raw }
func (r *ChannelConnectionSettingsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property EngagementID is required.
type MarkRecordingAsReadyRequestParam struct {
	EngagementID int64 `json:"engagementId,required"`
	paramObj
}

func (r MarkRecordingAsReadyRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow MarkRecordingAsReadyRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MarkRecordingAsReadyRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RecordingSettingsPatchRequestParam struct {
	URLToRetrieveAuthedRecording param.Opt[string] `json:"urlToRetrieveAuthedRecording,omitzero"`
	paramObj
}

func (r RecordingSettingsPatchRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow RecordingSettingsPatchRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RecordingSettingsPatchRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property URLToRetrieveAuthedRecording is required.
type RecordingSettingsRequestParam struct {
	URLToRetrieveAuthedRecording string `json:"urlToRetrieveAuthedRecording,required"`
	paramObj
}

func (r RecordingSettingsRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow RecordingSettingsRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RecordingSettingsRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RecordingSettingsResponse struct {
	// The URL used to retrieve authenticated call recordings.
	URLToRetrieveAuthedRecording string `json:"urlToRetrieveAuthedRecording,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		URLToRetrieveAuthedRecording respjson.Field
		ExtraFields                  map[string]respjson.Field
		raw                          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RecordingSettingsResponse) RawJSON() string { return r.JSON.raw }
func (r *RecordingSettingsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SettingsPatchRequestParam struct {
	Height                 param.Opt[int64]  `json:"height,omitzero"`
	IsReady                param.Opt[bool]   `json:"isReady,omitzero"`
	Name                   param.Opt[string] `json:"name,omitzero"`
	SupportsCustomObjects  param.Opt[bool]   `json:"supportsCustomObjects,omitzero"`
	SupportsInboundCalling param.Opt[bool]   `json:"supportsInboundCalling,omitzero"`
	URL                    param.Opt[string] `json:"url,omitzero"`
	UsesCallingWindow      param.Opt[bool]   `json:"usesCallingWindow,omitzero"`
	UsesRemote             param.Opt[bool]   `json:"usesRemote,omitzero"`
	Width                  param.Opt[int64]  `json:"width,omitzero"`
	paramObj
}

func (r SettingsPatchRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow SettingsPatchRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SettingsPatchRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Height, IsReady, Name, SupportsCustomObjects,
// SupportsInboundCalling, URL, UsesCallingWindow, UsesRemote, Width are required.
type SettingsRequestParam struct {
	Height                 int64  `json:"height,required"`
	IsReady                bool   `json:"isReady,required"`
	Name                   string `json:"name,required"`
	SupportsCustomObjects  bool   `json:"supportsCustomObjects,required"`
	SupportsInboundCalling bool   `json:"supportsInboundCalling,required"`
	URL                    string `json:"url,required"`
	UsesCallingWindow      bool   `json:"usesCallingWindow,required"`
	UsesRemote             bool   `json:"usesRemote,required"`
	Width                  int64  `json:"width,required"`
	paramObj
}

func (r SettingsRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow SettingsRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SettingsRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
