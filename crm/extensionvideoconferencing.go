// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package crm

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"slices"

	"github.com/stainless-sdks/hubspot-sdk-go/internal/apijson"
	shimjson "github.com/stainless-sdks/hubspot-sdk-go/internal/encoding/json"
	"github.com/stainless-sdks/hubspot-sdk-go/internal/requestconfig"
	"github.com/stainless-sdks/hubspot-sdk-go/option"
	"github.com/stainless-sdks/hubspot-sdk-go/packages/param"
	"github.com/stainless-sdks/hubspot-sdk-go/packages/respjson"
)

// ExtensionVideoConferencingService contains methods and other services that help
// with interacting with the hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewExtensionVideoConferencingService] method instead.
type ExtensionVideoConferencingService struct {
	options []option.RequestOption
}

// NewExtensionVideoConferencingService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewExtensionVideoConferencingService(opts ...option.RequestOption) (r ExtensionVideoConferencingService) {
	r = ExtensionVideoConferencingService{}
	r.options = opts
	return
}

// Create or update video conference extension settings for your app
func (r *ExtensionVideoConferencingService) Update(ctx context.Context, appID int64, body ExtensionVideoConferencingUpdateParams, opts ...option.RequestOption) (res *ExternalSettings, err error) {
	opts = slices.Concat(r.options, opts)
	path := fmt.Sprintf("crm/extensions/videoconferencing/2026-03/settings/%v", appID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return res, err
}

// Delete video conference extension settings for your app
func (r *ExtensionVideoConferencingService) Delete(ctx context.Context, appID int64, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := fmt.Sprintf("crm/extensions/videoconferencing/2026-03/settings/%v", appID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Fetch video conference extension settings for your app
func (r *ExtensionVideoConferencingService) Get(ctx context.Context, appID int64, opts ...option.RequestOption) (res *ExternalSettings, err error) {
	opts = slices.Concat(r.options, opts)
	path := fmt.Sprintf("crm/extensions/videoconferencing/2026-03/settings/%v", appID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type ExternalSettings struct {
	CreateMeetingURL string `json:"createMeetingUrl" api:"required"`
	DeleteMeetingURL string `json:"deleteMeetingUrl"`
	FetchAccountsUri string `json:"fetchAccountsUri"`
	UpdateMeetingURL string `json:"updateMeetingUrl"`
	UserVerifyURL    string `json:"userVerifyUrl"`
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

// The property CreateMeetingURL is required.
type ExternalSettingsParam struct {
	CreateMeetingURL string            `json:"createMeetingUrl" api:"required"`
	DeleteMeetingURL param.Opt[string] `json:"deleteMeetingUrl,omitzero"`
	FetchAccountsUri param.Opt[string] `json:"fetchAccountsUri,omitzero"`
	UpdateMeetingURL param.Opt[string] `json:"updateMeetingUrl,omitzero"`
	UserVerifyURL    param.Opt[string] `json:"userVerifyUrl,omitzero"`
	paramObj
}

func (r ExternalSettingsParam) MarshalJSON() (data []byte, err error) {
	type shadow ExternalSettingsParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExternalSettingsParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExtensionVideoConferencingUpdateParams struct {
	ExternalSettings ExternalSettingsParam
	paramObj
}

func (r ExtensionVideoConferencingUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.ExternalSettings)
}
func (r *ExtensionVideoConferencingUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
