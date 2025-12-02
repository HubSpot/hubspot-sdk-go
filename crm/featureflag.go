// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package crm

import (
	"github.com/stainless-sdks/hubspot-sdk-go/internal/apijson"
	"github.com/stainless-sdks/hubspot-sdk-go/option"
	"github.com/stainless-sdks/hubspot-sdk-go/packages/param"
	"github.com/stainless-sdks/hubspot-sdk-go/packages/respjson"
)

// FeatureFlagService contains methods and other services that help with
// interacting with the hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewFeatureFlagService] method instead.
type FeatureFlagService struct {
	Options []option.RequestOption
	Apps    FeatureFlagAppService
	Portals FeatureFlagPortalService
}

// NewFeatureFlagService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewFeatureFlagService(opts ...option.RequestOption) (r FeatureFlagService) {
	r = FeatureFlagService{}
	r.Options = opts
	r.Apps = NewFeatureFlagAppService(opts...)
	r.Portals = NewFeatureFlagPortalService(opts...)
	return
}

// The properties FlagState, PortalID are required.
type BatchPortalEntryParam struct {
	// Any of "ABSENT", "OFF", "ON".
	FlagState BatchPortalEntryFlagState `json:"flagState,omitzero,required"`
	PortalID  int64                     `json:"portalId,required"`
	paramObj
}

func (r BatchPortalEntryParam) MarshalJSON() (data []byte, err error) {
	type shadow BatchPortalEntryParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BatchPortalEntryParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BatchPortalEntryFlagState string

const (
	BatchPortalEntryFlagStateAbsent BatchPortalEntryFlagState = "ABSENT"
	BatchPortalEntryFlagStateOff    BatchPortalEntryFlagState = "OFF"
	BatchPortalEntryFlagStateOn     BatchPortalEntryFlagState = "ON"
)

// The property DefaultState is required.
type FlagPutRequestParam struct {
	// Any of "ABSENT", "OFF", "ON".
	DefaultState FlagPutRequestDefaultState `json:"defaultState,omitzero,required"`
	// Any of "ABSENT", "OFF", "ON".
	OverrideState FlagPutRequestOverrideState `json:"overrideState,omitzero"`
	paramObj
}

func (r FlagPutRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow FlagPutRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FlagPutRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FlagPutRequestDefaultState string

const (
	FlagPutRequestDefaultStateAbsent FlagPutRequestDefaultState = "ABSENT"
	FlagPutRequestDefaultStateOff    FlagPutRequestDefaultState = "OFF"
	FlagPutRequestDefaultStateOn     FlagPutRequestDefaultState = "ON"
)

type FlagPutRequestOverrideState string

const (
	FlagPutRequestOverrideStateAbsent FlagPutRequestOverrideState = "ABSENT"
	FlagPutRequestOverrideStateOff    FlagPutRequestOverrideState = "OFF"
	FlagPutRequestOverrideStateOn     FlagPutRequestOverrideState = "ON"
)

type FlagResponse struct {
	AppID int64 `json:"appId,required"`
	// Any of "ABSENT", "OFF", "ON".
	DefaultState FlagResponseDefaultState `json:"defaultState,required"`
	FlagName     string                   `json:"flagName,required"`
	// Any of "ABSENT", "OFF", "ON".
	OverrideState FlagResponseOverrideState `json:"overrideState"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AppID         respjson.Field
		DefaultState  respjson.Field
		FlagName      respjson.Field
		OverrideState respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FlagResponse) RawJSON() string { return r.JSON.raw }
func (r *FlagResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FlagResponseDefaultState string

const (
	FlagResponseDefaultStateAbsent FlagResponseDefaultState = "ABSENT"
	FlagResponseDefaultStateOff    FlagResponseDefaultState = "OFF"
	FlagResponseDefaultStateOn     FlagResponseDefaultState = "ON"
)

type FlagResponseOverrideState string

const (
	FlagResponseOverrideStateAbsent FlagResponseOverrideState = "ABSENT"
	FlagResponseOverrideStateOff    FlagResponseOverrideState = "OFF"
	FlagResponseOverrideStateOn     FlagResponseOverrideState = "ON"
)

// The property PortalIDs is required.
type PortalFlagStateBatchDeleteRequestParam struct {
	PortalIDs []int64 `json:"portalIds,omitzero,required"`
	paramObj
}

func (r PortalFlagStateBatchDeleteRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow PortalFlagStateBatchDeleteRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PortalFlagStateBatchDeleteRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property PortalStates is required.
type PortalFlagStateBatchPutRequestParam struct {
	PortalStates []BatchPortalEntryParam `json:"portalStates,omitzero,required"`
	paramObj
}

func (r PortalFlagStateBatchPutRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow PortalFlagStateBatchPutRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PortalFlagStateBatchPutRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PortalFlagStateBatchResponse struct {
	PortalFlagStates []PortalFlagStateResponse `json:"portalFlagStates,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PortalFlagStates respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PortalFlagStateBatchResponse) RawJSON() string { return r.JSON.raw }
func (r *PortalFlagStateBatchResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property FlagState is required.
type PortalFlagStatePutRequestParam struct {
	// Any of "ABSENT", "OFF", "ON".
	FlagState PortalFlagStatePutRequestFlagState `json:"flagState,omitzero,required"`
	paramObj
}

func (r PortalFlagStatePutRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow PortalFlagStatePutRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PortalFlagStatePutRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PortalFlagStatePutRequestFlagState string

const (
	PortalFlagStatePutRequestFlagStateAbsent PortalFlagStatePutRequestFlagState = "ABSENT"
	PortalFlagStatePutRequestFlagStateOff    PortalFlagStatePutRequestFlagState = "OFF"
	PortalFlagStatePutRequestFlagStateOn     PortalFlagStatePutRequestFlagState = "ON"
)

type PortalFlagStateResponse struct {
	AppID    int64  `json:"appId,required"`
	FlagName string `json:"flagName,required"`
	// Any of "ABSENT", "OFF", "ON".
	FlagState PortalFlagStateResponseFlagState `json:"flagState,required"`
	PortalID  int64                            `json:"portalId,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AppID       respjson.Field
		FlagName    respjson.Field
		FlagState   respjson.Field
		PortalID    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PortalFlagStateResponse) RawJSON() string { return r.JSON.raw }
func (r *PortalFlagStateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PortalFlagStateResponseFlagState string

const (
	PortalFlagStateResponseFlagStateAbsent PortalFlagStateResponseFlagState = "ABSENT"
	PortalFlagStateResponseFlagStateOff    PortalFlagStateResponseFlagState = "OFF"
	PortalFlagStateResponseFlagStateOn     PortalFlagStateResponseFlagState = "ON"
)
