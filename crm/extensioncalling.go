// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package crm

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"slices"
	"time"

	"github.com/stainless-sdks/hubspot-sdk-go/internal/apijson"
	shimjson "github.com/stainless-sdks/hubspot-sdk-go/internal/encoding/json"
	"github.com/stainless-sdks/hubspot-sdk-go/internal/requestconfig"
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
	options     []option.RequestOption
	Transcripts ExtensionCallingTranscriptService
}

// NewExtensionCallingService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewExtensionCallingService(opts ...option.RequestOption) (r ExtensionCallingService) {
	r = ExtensionCallingService{}
	r.options = opts
	r.Transcripts = NewExtensionCallingTranscriptService(opts...)
	return
}

// Establish new channel connection settings for the specified app.
func (r *ExtensionCallingService) NewChannelConnectionSettings(ctx context.Context, appID int64, body ExtensionCallingNewChannelConnectionSettingsParams, opts ...option.RequestOption) (res *ChannelConnectionSettingsResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := fmt.Sprintf("crm/extensions/calling/2026-03/%v/settings/channel-connection", appID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

func (r *ExtensionCallingService) NewInboundCall(ctx context.Context, body ExtensionCallingNewInboundCallParams, opts ...option.RequestOption) (res *CompletedThirdPartyCallResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "crm/extensions/calling/2026-03/inbound-call"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// This endpoint is used to mark a call recording as ready. It requires the
// engagementId to identify the specific recording.
func (r *ExtensionCallingService) NewRecordingReady(ctx context.Context, body ExtensionCallingNewRecordingReadyParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "crm/extensions/calling/2026-03/recordings/ready"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return err
}

// Create new recording settings for a specific app using the provided app ID.
func (r *ExtensionCallingService) NewRecordingSettings(ctx context.Context, appID int64, body ExtensionCallingNewRecordingSettingsParams, opts ...option.RequestOption) (res *RecordingSettingsResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := fmt.Sprintf("crm/extensions/calling/2026-03/%v/settings/recording", appID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Create new settings for the calling extension associated with the specified
// appId.
func (r *ExtensionCallingService) NewSettings(ctx context.Context, appID int64, body ExtensionCallingNewSettingsParams, opts ...option.RequestOption) (res *SettingsResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := fmt.Sprintf("crm/extensions/calling/2026-03/%v/settings", appID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Delete the channel connection settings associated with the specified app.
func (r *ExtensionCallingService) DeleteChannelConnectionSettings(ctx context.Context, appID int64, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := fmt.Sprintf("crm/extensions/calling/2026-03/%v/settings/channel-connection", appID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Remove the calling extension settings associated with the specified appId. This
// action cannot be undone.
func (r *ExtensionCallingService) DeleteSettings(ctx context.Context, appID int64, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := fmt.Sprintf("crm/extensions/calling/2026-03/%v/settings", appID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Access the current channel connection settings for the specified app.
func (r *ExtensionCallingService) GetChannelConnectionSettings(ctx context.Context, appID int64, opts ...option.RequestOption) (res *ChannelConnectionSettingsResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := fmt.Sprintf("crm/extensions/calling/2026-03/%v/settings/channel-connection", appID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Retrieve the current recording settings for a specific app using the provided
// app ID.
func (r *ExtensionCallingService) GetRecordingSettings(ctx context.Context, appID int64, opts ...option.RequestOption) (res *RecordingSettingsResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := fmt.Sprintf("crm/extensions/calling/2026-03/%v/settings/recording", appID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Retrieve the current settings of the calling extension for the specified appId.
func (r *ExtensionCallingService) GetSettings(ctx context.Context, appID int64, opts ...option.RequestOption) (res *SettingsResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := fmt.Sprintf("crm/extensions/calling/2026-03/%v/settings", appID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Modify the existing channel connection settings for the specified app.
func (r *ExtensionCallingService) UpdateChannelConnectionSettings(ctx context.Context, appID int64, body ExtensionCallingUpdateChannelConnectionSettingsParams, opts ...option.RequestOption) (res *ChannelConnectionSettingsResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := fmt.Sprintf("crm/extensions/calling/2026-03/%v/settings/channel-connection", appID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

// Update the recording settings for a specific app using the provided app ID.
func (r *ExtensionCallingService) UpdateRecordingSettings(ctx context.Context, appID int64, body ExtensionCallingUpdateRecordingSettingsParams, opts ...option.RequestOption) (res *RecordingSettingsResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := fmt.Sprintf("crm/extensions/calling/2026-03/%v/settings/recording", appID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

// Modify existing calling extension settings for the specified appId. Only the
// fields provided in the request will be updated.
func (r *ExtensionCallingService) UpdateSettings(ctx context.Context, appID int64, body ExtensionCallingUpdateSettingsParams, opts ...option.RequestOption) (res *SettingsResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := fmt.Sprintf("crm/extensions/calling/2026-03/%v/settings", appID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

type ChannelConnectionSettingsPatchRequestParam struct {
	// Indicates whether the channel connection settings are ready.
	IsReady param.Opt[bool] `json:"isReady,omitzero"`
	// The URL for the channel connection settings.
	URL param.Opt[string] `json:"url,omitzero"`
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
	// Indicates whether the channel connection settings are ready.
	IsReady bool `json:"isReady" api:"required"`
	// The URL associated with the channel connection settings.
	URL string `json:"url" api:"required"`
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
	// The date and time when the channel connection settings were created.
	CreatedAt time.Time `json:"createdAt" api:"required" format:"date-time"`
	// Indicates whether the channel connection settings are ready for use.
	IsReady bool `json:"isReady" api:"required"`
	// The date and time when the channel connection settings were last updated.
	UpdatedAt time.Time `json:"updatedAt" api:"required" format:"date-time"`
	// The URL associated with the channel connection settings.
	URL string `json:"url" api:"required"`
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

type CompanyCallerID struct {
	// Specifies the type of caller ID, which is set to 'COMPANY' by default.
	//
	// Any of "COMPANY".
	CallerIDType      CompanyCallerIDCallerIDType `json:"callerIdType" api:"required"`
	ObjectCoordinates ObjectCoordinates           `json:"objectCoordinates" api:"required"`
	// The name associated with the company caller ID.
	Name string `json:"name"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CallerIDType      respjson.Field
		ObjectCoordinates respjson.Field
		Name              respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CompanyCallerID) RawJSON() string { return r.JSON.raw }
func (r *CompanyCallerID) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Specifies the type of caller ID, which is set to 'COMPANY' by default.
type CompanyCallerIDCallerIDType string

const (
	CompanyCallerIDCallerIDTypeCompany CompanyCallerIDCallerIDType = "COMPANY"
)

// The properties CreateEngagement, EngagementProperties, ExternalCallID,
// FinalCallStatus, FromNumber, PotentialRecipientUserIDs, ToNumber are required.
type CompletedThirdPartyCallRequestParam struct {
	// Indicates whether an engagement should be created for the call.
	CreateEngagement bool `json:"createEngagement" api:"required"`
	// Contains additional properties related to the engagement.
	EngagementProperties map[string]string `json:"engagementProperties,omitzero" api:"required"`
	// The unique identifier for the call from an external system.
	ExternalCallID string `json:"externalCallId" api:"required"`
	// The final status of the call, with accepted values including: BUSY,
	// CALLING_CRM_USER, CANCELED, COMPLETED, CONNECTING, FAILED, HOLD, IN_PROGRESS,
	// MISSED, NO_ANSWER, QUEUED, RINGING, UNKNOWN.
	//
	// Any of "BUSY", "CALLING_CRM_USER", "CANCELED", "COMPLETED", "CONNECTING",
	// "FAILED", "HOLD", "IN_PROGRESS", "MISSED", "NO_ANSWER", "QUEUED", "RINGING",
	// "UNKNOWN".
	FinalCallStatus           CompletedThirdPartyCallRequestFinalCallStatus `json:"finalCallStatus,omitzero" api:"required"`
	FromNumber                FormattedPhoneNumberParam                     `json:"fromNumber,omitzero" api:"required"`
	PotentialRecipientUserIDs []int64                                       `json:"potentialRecipientUserIds,omitzero" api:"required"`
	ToNumber                  FormattedPhoneNumberParam                     `json:"toNumber,omitzero" api:"required"`
	// The timestamp indicating when the call started, formatted as a date-time string.
	CallStartedTimestamp param.Opt[time.Time] `json:"callStartedTimestamp,omitzero" format:"date-time"`
	// The duration of the call in seconds.
	DurationSeconds param.Opt[int64] `json:"durationSeconds,omitzero"`
	// The ID of the user associated with the call.
	UserID param.Opt[int64] `json:"userId,omitzero"`
	paramObj
}

func (r CompletedThirdPartyCallRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow CompletedThirdPartyCallRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CompletedThirdPartyCallRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The final status of the call, with accepted values including: BUSY,
// CALLING_CRM_USER, CANCELED, COMPLETED, CONNECTING, FAILED, HOLD, IN_PROGRESS,
// MISSED, NO_ANSWER, QUEUED, RINGING, UNKNOWN.
type CompletedThirdPartyCallRequestFinalCallStatus string

const (
	CompletedThirdPartyCallRequestFinalCallStatusBusy           CompletedThirdPartyCallRequestFinalCallStatus = "BUSY"
	CompletedThirdPartyCallRequestFinalCallStatusCallingCrmUser CompletedThirdPartyCallRequestFinalCallStatus = "CALLING_CRM_USER"
	CompletedThirdPartyCallRequestFinalCallStatusCanceled       CompletedThirdPartyCallRequestFinalCallStatus = "CANCELED"
	CompletedThirdPartyCallRequestFinalCallStatusCompleted      CompletedThirdPartyCallRequestFinalCallStatus = "COMPLETED"
	CompletedThirdPartyCallRequestFinalCallStatusConnecting     CompletedThirdPartyCallRequestFinalCallStatus = "CONNECTING"
	CompletedThirdPartyCallRequestFinalCallStatusFailed         CompletedThirdPartyCallRequestFinalCallStatus = "FAILED"
	CompletedThirdPartyCallRequestFinalCallStatusHold           CompletedThirdPartyCallRequestFinalCallStatus = "HOLD"
	CompletedThirdPartyCallRequestFinalCallStatusInProgress     CompletedThirdPartyCallRequestFinalCallStatus = "IN_PROGRESS"
	CompletedThirdPartyCallRequestFinalCallStatusMissed         CompletedThirdPartyCallRequestFinalCallStatus = "MISSED"
	CompletedThirdPartyCallRequestFinalCallStatusNoAnswer       CompletedThirdPartyCallRequestFinalCallStatus = "NO_ANSWER"
	CompletedThirdPartyCallRequestFinalCallStatusQueued         CompletedThirdPartyCallRequestFinalCallStatus = "QUEUED"
	CompletedThirdPartyCallRequestFinalCallStatusRinging        CompletedThirdPartyCallRequestFinalCallStatus = "RINGING"
	CompletedThirdPartyCallRequestFinalCallStatusUnknown        CompletedThirdPartyCallRequestFinalCallStatus = "UNKNOWN"
)

type CompletedThirdPartyCallResponse struct {
	CallerIDMatches []CompletedThirdPartyCallResponseCallerIDMatchUnion `json:"callerIdMatches" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CallerIDMatches respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CompletedThirdPartyCallResponse) RawJSON() string { return r.JSON.raw }
func (r *CompletedThirdPartyCallResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// CompletedThirdPartyCallResponseCallerIDMatchUnion contains all possible
// properties and values from [ContactCallerID], [CompanyCallerID].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type CompletedThirdPartyCallResponseCallerIDMatchUnion struct {
	CallerIDType string `json:"callerIdType"`
	// This field is from variant [ContactCallerID].
	ObjectCoordinates ObjectCoordinates `json:"objectCoordinates"`
	// This field is from variant [ContactCallerID].
	Email string `json:"email"`
	// This field is from variant [ContactCallerID].
	FirstName string `json:"firstName"`
	// This field is from variant [ContactCallerID].
	LastName string `json:"lastName"`
	// This field is from variant [CompanyCallerID].
	Name string `json:"name"`
	JSON struct {
		CallerIDType      respjson.Field
		ObjectCoordinates respjson.Field
		Email             respjson.Field
		FirstName         respjson.Field
		LastName          respjson.Field
		Name              respjson.Field
		raw               string
	} `json:"-"`
}

func (u CompletedThirdPartyCallResponseCallerIDMatchUnion) AsContact() (v ContactCallerID) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CompletedThirdPartyCallResponseCallerIDMatchUnion) AsCompany() (v CompanyCallerID) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u CompletedThirdPartyCallResponseCallerIDMatchUnion) RawJSON() string { return u.JSON.raw }

func (r *CompletedThirdPartyCallResponseCallerIDMatchUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ContactCallerID struct {
	// Specifies the type of caller ID, with the default value being CONTACT.
	//
	// Any of "CONTACT".
	CallerIDType      ContactCallerIDCallerIDType `json:"callerIdType" api:"required"`
	ObjectCoordinates ObjectCoordinates           `json:"objectCoordinates" api:"required"`
	// The email address of the contact.
	Email string `json:"email"`
	// The first name of the contact.
	FirstName string `json:"firstName"`
	// The last name of the contact.
	LastName string `json:"lastName"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CallerIDType      respjson.Field
		ObjectCoordinates respjson.Field
		Email             respjson.Field
		FirstName         respjson.Field
		LastName          respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ContactCallerID) RawJSON() string { return r.JSON.raw }
func (r *ContactCallerID) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Specifies the type of caller ID, with the default value being CONTACT.
type ContactCallerIDCallerIDType string

const (
	ContactCallerIDCallerIDTypeContact ContactCallerIDCallerIDType = "CONTACT"
)

// The properties E164Number, PhoneNumberType are required.
type FormattedPhoneNumberParam struct {
	// The phone number formatted in E.164 standard.
	E164Number string `json:"e164Number" api:"required"`
	// The type of phone number, with accepted values including FIXED_LINE, MOBILE,
	// VOIP, and others.
	//
	// Any of "FIXED_LINE", "FIXED_LINE_OR_MOBILE", "MOBILE", "PAGER",
	// "PERSONAL_NUMBER", "PREMIUM_RATE", "SHARED_COST", "TOLL_FREE", "UAN", "UNKNOWN",
	// "VOICEMAIL", "VOIP".
	PhoneNumberType FormattedPhoneNumberPhoneNumberType `json:"phoneNumberType,omitzero" api:"required"`
	// The extension number associated with the phone number.
	Extension param.Opt[string] `json:"extension,omitzero"`
	paramObj
}

func (r FormattedPhoneNumberParam) MarshalJSON() (data []byte, err error) {
	type shadow FormattedPhoneNumberParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FormattedPhoneNumberParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The type of phone number, with accepted values including FIXED_LINE, MOBILE,
// VOIP, and others.
type FormattedPhoneNumberPhoneNumberType string

const (
	FormattedPhoneNumberPhoneNumberTypeFixedLine         FormattedPhoneNumberPhoneNumberType = "FIXED_LINE"
	FormattedPhoneNumberPhoneNumberTypeFixedLineOrMobile FormattedPhoneNumberPhoneNumberType = "FIXED_LINE_OR_MOBILE"
	FormattedPhoneNumberPhoneNumberTypeMobile            FormattedPhoneNumberPhoneNumberType = "MOBILE"
	FormattedPhoneNumberPhoneNumberTypePager             FormattedPhoneNumberPhoneNumberType = "PAGER"
	FormattedPhoneNumberPhoneNumberTypePersonalNumber    FormattedPhoneNumberPhoneNumberType = "PERSONAL_NUMBER"
	FormattedPhoneNumberPhoneNumberTypePremiumRate       FormattedPhoneNumberPhoneNumberType = "PREMIUM_RATE"
	FormattedPhoneNumberPhoneNumberTypeSharedCost        FormattedPhoneNumberPhoneNumberType = "SHARED_COST"
	FormattedPhoneNumberPhoneNumberTypeTollFree          FormattedPhoneNumberPhoneNumberType = "TOLL_FREE"
	FormattedPhoneNumberPhoneNumberTypeUan               FormattedPhoneNumberPhoneNumberType = "UAN"
	FormattedPhoneNumberPhoneNumberTypeUnknown           FormattedPhoneNumberPhoneNumberType = "UNKNOWN"
	FormattedPhoneNumberPhoneNumberTypeVoicemail         FormattedPhoneNumberPhoneNumberType = "VOICEMAIL"
	FormattedPhoneNumberPhoneNumberTypeVoip              FormattedPhoneNumberPhoneNumberType = "VOIP"
)

// The property EngagementID is required.
type MarkRecordingAsReadyRequestParam struct {
	// The unique identifier for the engagement associated with the call recording.
	EngagementID int64 `json:"engagementId" api:"required"`
	paramObj
}

func (r MarkRecordingAsReadyRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow MarkRecordingAsReadyRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MarkRecordingAsReadyRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ObjectCoordinates struct {
	// The unique identifier for the object.
	ObjectID int64 `json:"objectId" api:"required"`
	// The type identifier for the object.
	ObjectTypeID string `json:"objectTypeId" api:"required"`
	// The unique identifier for the portal.
	PortalID int64 `json:"portalId" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ObjectID     respjson.Field
		ObjectTypeID respjson.Field
		PortalID     respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ObjectCoordinates) RawJSON() string { return r.JSON.raw }
func (r *ObjectCoordinates) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RecordingSettingsPatchRequestParam struct {
	// The URL used to access authenticated call recordings.
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
	// The URL used to access authenticated call recordings.
	URLToRetrieveAuthedRecording string `json:"urlToRetrieveAuthedRecording" api:"required"`
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
	URLToRetrieveAuthedRecording string `json:"urlToRetrieveAuthedRecording" api:"required"`
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
	// The height setting for the calling extension interface.
	Height param.Opt[int64] `json:"height,omitzero"`
	// Specifies whether the calling extension is ready for use.
	IsReady param.Opt[bool] `json:"isReady,omitzero"`
	// The name of the calling extension.
	Name param.Opt[string] `json:"name,omitzero"`
	// Indicates if the calling extension supports custom objects.
	SupportsCustomObjects param.Opt[bool] `json:"supportsCustomObjects,omitzero"`
	// Indicates if the calling extension supports inbound calling.
	SupportsInboundCalling param.Opt[bool] `json:"supportsInboundCalling,omitzero"`
	// The URL associated with the calling extension settings.
	URL param.Opt[string] `json:"url,omitzero"`
	// Indicates if the calling extension uses a calling window.
	UsesCallingWindow param.Opt[bool] `json:"usesCallingWindow,omitzero"`
	// Indicates if the calling extension uses a remote connection.
	UsesRemote param.Opt[bool] `json:"usesRemote,omitzero"`
	// The width setting for the calling extension interface.
	Width param.Opt[int64] `json:"width,omitzero"`
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
	// Specifies the height of the calling extension interface.
	Height int64 `json:"height" api:"required"`
	// Indicates if the calling extension is ready for use.
	IsReady bool `json:"isReady" api:"required"`
	// The name of the calling extension.
	Name string `json:"name" api:"required"`
	// Indicates if the calling extension supports custom objects.
	SupportsCustomObjects bool `json:"supportsCustomObjects" api:"required"`
	// Indicates if the calling extension supports inbound calling.
	SupportsInboundCalling bool `json:"supportsInboundCalling" api:"required"`
	// The URL associated with the calling extension.
	URL string `json:"url" api:"required"`
	// Indicates if the calling extension uses a separate calling window.
	UsesCallingWindow bool `json:"usesCallingWindow" api:"required"`
	// Indicates if the calling extension uses remote services.
	UsesRemote bool `json:"usesRemote" api:"required"`
	// Specifies the width of the calling extension interface.
	Width int64 `json:"width" api:"required"`
	paramObj
}

func (r SettingsRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow SettingsRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SettingsRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SettingsResponse struct {
	// The date and time when the calling extension settings were created.
	CreatedAt time.Time `json:"createdAt" api:"required" format:"date-time"`
	// The height of the calling extension interface.
	Height int64 `json:"height" api:"required"`
	// Specifies whether the calling extension settings are ready for use.
	IsReady bool `json:"isReady" api:"required"`
	// The name of the calling extension.
	Name string `json:"name" api:"required"`
	// Indicates if the calling extension supports custom objects.
	SupportsCustomObjects bool `json:"supportsCustomObjects" api:"required"`
	// Indicates if the calling extension supports inbound calling.
	SupportsInboundCalling bool `json:"supportsInboundCalling" api:"required"`
	// The date and time when the calling extension settings were last updated.
	UpdatedAt time.Time `json:"updatedAt" api:"required" format:"date-time"`
	// The URL associated with the calling extension.
	URL string `json:"url" api:"required"`
	// Specifies if the calling extension uses a dedicated calling window.
	UsesCallingWindow bool `json:"usesCallingWindow" api:"required"`
	// Indicates if the calling extension uses a remote service.
	UsesRemote bool `json:"usesRemote" api:"required"`
	// The width of the calling extension interface.
	Width int64 `json:"width" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreatedAt              respjson.Field
		Height                 respjson.Field
		IsReady                respjson.Field
		Name                   respjson.Field
		SupportsCustomObjects  respjson.Field
		SupportsInboundCalling respjson.Field
		UpdatedAt              respjson.Field
		URL                    respjson.Field
		UsesCallingWindow      respjson.Field
		UsesRemote             respjson.Field
		Width                  respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SettingsResponse) RawJSON() string { return r.JSON.raw }
func (r *SettingsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExtensionCallingNewChannelConnectionSettingsParams struct {
	ChannelConnectionSettingsRequest ChannelConnectionSettingsRequestParam
	paramObj
}

func (r ExtensionCallingNewChannelConnectionSettingsParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.ChannelConnectionSettingsRequest)
}
func (r *ExtensionCallingNewChannelConnectionSettingsParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExtensionCallingNewInboundCallParams struct {
	CompletedThirdPartyCallRequest CompletedThirdPartyCallRequestParam
	paramObj
}

func (r ExtensionCallingNewInboundCallParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.CompletedThirdPartyCallRequest)
}
func (r *ExtensionCallingNewInboundCallParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExtensionCallingNewRecordingReadyParams struct {
	MarkRecordingAsReadyRequest MarkRecordingAsReadyRequestParam
	paramObj
}

func (r ExtensionCallingNewRecordingReadyParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.MarkRecordingAsReadyRequest)
}
func (r *ExtensionCallingNewRecordingReadyParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExtensionCallingNewRecordingSettingsParams struct {
	RecordingSettingsRequest RecordingSettingsRequestParam
	paramObj
}

func (r ExtensionCallingNewRecordingSettingsParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.RecordingSettingsRequest)
}
func (r *ExtensionCallingNewRecordingSettingsParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExtensionCallingNewSettingsParams struct {
	SettingsRequest SettingsRequestParam
	paramObj
}

func (r ExtensionCallingNewSettingsParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.SettingsRequest)
}
func (r *ExtensionCallingNewSettingsParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExtensionCallingUpdateChannelConnectionSettingsParams struct {
	ChannelConnectionSettingsPatchRequest ChannelConnectionSettingsPatchRequestParam
	paramObj
}

func (r ExtensionCallingUpdateChannelConnectionSettingsParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.ChannelConnectionSettingsPatchRequest)
}
func (r *ExtensionCallingUpdateChannelConnectionSettingsParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExtensionCallingUpdateRecordingSettingsParams struct {
	RecordingSettingsPatchRequest RecordingSettingsPatchRequestParam
	paramObj
}

func (r ExtensionCallingUpdateRecordingSettingsParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.RecordingSettingsPatchRequest)
}
func (r *ExtensionCallingUpdateRecordingSettingsParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExtensionCallingUpdateSettingsParams struct {
	SettingsPatchRequest SettingsPatchRequestParam
	paramObj
}

func (r ExtensionCallingUpdateSettingsParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.SettingsPatchRequest)
}
func (r *ExtensionCallingUpdateSettingsParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
