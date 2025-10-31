// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package crm

import (
	"encoding/json"

	"github.com/stainless-sdks/hubspot-sdk-go/internal/apijson"
	"github.com/stainless-sdks/hubspot-sdk-go/option"
	"github.com/stainless-sdks/hubspot-sdk-go/packages/param"
	"github.com/stainless-sdks/hubspot-sdk-go/packages/respjson"
)

// ExtensionVideoconferencingService contains methods and other services that help
// with interacting with the Hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewExtensionVideoconferencingService] method instead.
type ExtensionVideoconferencingService struct {
	Options []option.RequestOption
}

// NewExtensionVideoconferencingService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewExtensionVideoconferencingService(opts ...option.RequestOption) (r ExtensionVideoconferencingService) {
	r = ExtensionVideoconferencingService{}
	r.Options = opts
	return
}

// The URLs of the various actions provided by the video conferencing application.
// All URLs must use the `https` protocol.
type ExternalSettings struct {
	// The URL that HubSpot will send requests to create a new video conference.
	CreateMeetingURL string `json:"createMeetingUrl,required"`
	// The URL that HubSpot will send notifications of meetings that have been deleted
	// in HubSpot.
	DeleteMeetingURL string `json:"deleteMeetingUrl"`
	FetchAccountsUri string `json:"fetchAccountsUri"`
	// The URL that HubSpot will send updates to existing meetings. Typically called
	// when the user changes the topic or times of a meeting.
	UpdateMeetingURL string `json:"updateMeetingUrl"`
	// The URL that HubSpot will use to verify that a user exists in the video
	// conference application.
	UserVerifyURL string `json:"userVerifyUrl"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreateMeetingURL respjson.Field
		DeleteMeetingURL respjson.Field
		FetchAccountsUri respjson.Field
		UpdateMeetingURL respjson.Field
		UserVerifyURL    respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExternalSettings) RawJSON() string { return r.JSON.raw }
func (r *ExternalSettings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this ExternalSettings to a ExternalSettingsParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// ExternalSettingsParam.Overrides()
func (r ExternalSettings) ToParam() ExternalSettingsParam {
	return param.Override[ExternalSettingsParam](json.RawMessage(r.RawJSON()))
}

// The URLs of the various actions provided by the video conferencing application.
// All URLs must use the `https` protocol.
//
// The property CreateMeetingURL is required.
type ExternalSettingsParam struct {
	// The URL that HubSpot will send requests to create a new video conference.
	CreateMeetingURL string `json:"createMeetingUrl,required"`
	// The URL that HubSpot will send notifications of meetings that have been deleted
	// in HubSpot.
	DeleteMeetingURL param.Opt[string] `json:"deleteMeetingUrl,omitzero"`
	FetchAccountsUri param.Opt[string] `json:"fetchAccountsUri,omitzero"`
	// The URL that HubSpot will send updates to existing meetings. Typically called
	// when the user changes the topic or times of a meeting.
	UpdateMeetingURL param.Opt[string] `json:"updateMeetingUrl,omitzero"`
	// The URL that HubSpot will use to verify that a user exists in the video
	// conference application.
	UserVerifyURL param.Opt[string] `json:"userVerifyUrl,omitzero"`
	paramObj
}

func (r ExternalSettingsParam) MarshalJSON() (data []byte, err error) {
	type shadow ExternalSettingsParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExternalSettingsParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
