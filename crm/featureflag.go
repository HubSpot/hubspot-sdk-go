// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package crm

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/stainless-sdks/hubspot-sdk-go/internal/apijson"
	"github.com/stainless-sdks/hubspot-sdk-go/internal/apiquery"
	shimjson "github.com/stainless-sdks/hubspot-sdk-go/internal/encoding/json"
	"github.com/stainless-sdks/hubspot-sdk-go/internal/requestconfig"
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
	options []option.RequestOption
	Batch   FeatureFlagBatchService
}

// NewFeatureFlagService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewFeatureFlagService(opts ...option.RequestOption) (r FeatureFlagService) {
	r = FeatureFlagService{}
	r.options = opts
	r.Batch = NewFeatureFlagBatchService(opts...)
	return
}

// Set a feature flag for an app. For example, update the `hs-hide-crm-cards`
// flag's `defaultState` to `ON` to hide classic CRM cards from new installs.
func (r *FeatureFlagService) Update(ctx context.Context, flagName string, params FeatureFlagUpdateParams, opts ...option.RequestOption) (res *FlagResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if flagName == "" {
		err = errors.New("missing required flagName parameter")
		return nil, err
	}
	path := fmt.Sprintf("feature-flags/2026-03/%v/flags/%s", params.AppID, url.PathEscape(flagName))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return res, err
}

// Delete a feature flag in an app. For example, delete the `hs-release-app-cards`
// flag after all accounts have been migrated.
func (r *FeatureFlagService) Delete(ctx context.Context, flagName string, body FeatureFlagDeleteParams, opts ...option.RequestOption) (res *FlagResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if flagName == "" {
		err = errors.New("missing required flagName parameter")
		return nil, err
	}
	path := fmt.Sprintf("feature-flags/2026-03/%v/flags/%s", body.AppID, url.PathEscape(flagName))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

// Delete an account-level flag state for a specific HubSpot account. No request
// body is included.
func (r *FeatureFlagService) DeletePortalState(ctx context.Context, portalID int64, body FeatureFlagDeletePortalStateParams, opts ...option.RequestOption) (res *PortalFlagStateResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if body.FlagName == "" {
		err = errors.New("missing required flagName parameter")
		return nil, err
	}
	path := fmt.Sprintf("feature-flags/2026-03/%v/flags/%s/portals/%v", body.AppID, url.PathEscape(body.FlagName), portalID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

// Retrieve the current status of the app's feature flags. No request body is
// included.
func (r *FeatureFlagService) Get(ctx context.Context, flagName string, query FeatureFlagGetParams, opts ...option.RequestOption) (res *FlagResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if flagName == "" {
		err = errors.New("missing required flagName parameter")
		return nil, err
	}
	path := fmt.Sprintf("feature-flags/2026-03/%v/flags/%s", query.AppID, url.PathEscape(flagName))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Retrieve the account-level flag state of a specific HubSpot account.
func (r *FeatureFlagService) GetPortalState(ctx context.Context, portalID int64, query FeatureFlagGetPortalStateParams, opts ...option.RequestOption) (res *PortalFlagStateResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if query.FlagName == "" {
		err = errors.New("missing required flagName parameter")
		return nil, err
	}
	path := fmt.Sprintf("feature-flags/2026-03/%v/flags/%s/portals/%v", query.AppID, url.PathEscape(query.FlagName), portalID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

func (r *FeatureFlagService) ListAll(ctx context.Context, appID int64, opts ...option.RequestOption) (res *FlagsForAppResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := fmt.Sprintf("feature-flags/2026-03/%v/flags/all", appID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Retrieve a list of HubSpot accounts with an account-level flag setting for the
// specified app. No request body is included.
func (r *FeatureFlagService) ListPortals(ctx context.Context, flagName string, params FeatureFlagListPortalsParams, opts ...option.RequestOption) (res *PortalFlagStateBatchResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if flagName == "" {
		err = errors.New("missing required flagName parameter")
		return nil, err
	}
	path := fmt.Sprintf("feature-flags/2026-03/%v/flags/%s/portals", params.AppID, url.PathEscape(flagName))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// Specify an account-level flag state for a specific HubSpot account.
func (r *FeatureFlagService) UpdatePortalState(ctx context.Context, portalID int64, params FeatureFlagUpdatePortalStateParams, opts ...option.RequestOption) (res *PortalFlagStateResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if params.FlagName == "" {
		err = errors.New("missing required flagName parameter")
		return nil, err
	}
	path := fmt.Sprintf("feature-flags/2026-03/%v/flags/%s/portals/%v", params.AppID, url.PathEscape(params.FlagName), portalID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return res, err
}

// The properties FlagState, PortalID are required.
type BatchPortalEntryParam struct {
	// The flag state for this portal (e.g. ON or OFF)
	//
	// Any of "ABSENT", "OFF", "ON".
	FlagState BatchPortalEntryFlagState `json:"flagState,omitzero" api:"required"`
	// The ID of the portal
	PortalID int64 `json:"portalId" api:"required"`
	paramObj
}

func (r BatchPortalEntryParam) MarshalJSON() (data []byte, err error) {
	type shadow BatchPortalEntryParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BatchPortalEntryParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The flag state for this portal (e.g. ON or OFF)
type BatchPortalEntryFlagState string

const (
	BatchPortalEntryFlagStateAbsent BatchPortalEntryFlagState = "ABSENT"
	BatchPortalEntryFlagStateOff    BatchPortalEntryFlagState = "OFF"
	BatchPortalEntryFlagStateOn     BatchPortalEntryFlagState = "ON"
)

// The property DefaultState is required.
type FlagPutRequestParam struct {
	// The state that the flag should have if there are no overrides for a particular
	// portal
	//
	// Any of "ABSENT", "OFF", "ON".
	DefaultState FlagPutRequestDefaultState `json:"defaultState,omitzero" api:"required"`
	// A flag value that supercedes all other overrides, including portal-level values.
	// Mostly used for things like emergency overrides
	//
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

// The state that the flag should have if there are no overrides for a particular
// portal
type FlagPutRequestDefaultState string

const (
	FlagPutRequestDefaultStateAbsent FlagPutRequestDefaultState = "ABSENT"
	FlagPutRequestDefaultStateOff    FlagPutRequestDefaultState = "OFF"
	FlagPutRequestDefaultStateOn     FlagPutRequestDefaultState = "ON"
)

// A flag value that supercedes all other overrides, including portal-level values.
// Mostly used for things like emergency overrides
type FlagPutRequestOverrideState string

const (
	FlagPutRequestOverrideStateAbsent FlagPutRequestOverrideState = "ABSENT"
	FlagPutRequestOverrideStateOff    FlagPutRequestOverrideState = "OFF"
	FlagPutRequestOverrideStateOn     FlagPutRequestOverrideState = "ON"
)

type FlagResponse struct {
	// The ID of the app
	AppID int64 `json:"appId" api:"required"`
	// The flag state for any portal that doesn't have an override value
	//
	// Any of "ABSENT", "OFF", "ON".
	DefaultState FlagResponseDefaultState `json:"defaultState" api:"required"`
	// The name of the flag
	FlagName string `json:"flagName" api:"required"`
	// An optional flag value that overrides all others for this flag name and app,
	// including portal-level values
	//
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

// The flag state for any portal that doesn't have an override value
type FlagResponseDefaultState string

const (
	FlagResponseDefaultStateAbsent FlagResponseDefaultState = "ABSENT"
	FlagResponseDefaultStateOff    FlagResponseDefaultState = "OFF"
	FlagResponseDefaultStateOn     FlagResponseDefaultState = "ON"
)

// An optional flag value that overrides all others for this flag name and app,
// including portal-level values
type FlagResponseOverrideState string

const (
	FlagResponseOverrideStateAbsent FlagResponseOverrideState = "ABSENT"
	FlagResponseOverrideStateOff    FlagResponseOverrideState = "OFF"
	FlagResponseOverrideStateOn     FlagResponseOverrideState = "ON"
)

type FlagsForAppResponse struct {
	FlagsForApp []string `json:"flagsForApp" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FlagsForApp respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FlagsForAppResponse) RawJSON() string { return r.JSON.raw }
func (r *FlagsForAppResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property PortalIDs is required.
type PortalFlagStateBatchDeleteRequestParam struct {
	PortalIDs []int64 `json:"portalIds,omitzero" api:"required"`
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
	PortalStates []BatchPortalEntryParam `json:"portalStates,omitzero" api:"required"`
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
	PortalFlagStates []PortalFlagStateResponse `json:"portalFlagStates" api:"required"`
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
	// The state that the given flag should be in for this portal
	//
	// Any of "ABSENT", "OFF", "ON".
	FlagState PortalFlagStatePutRequestFlagState `json:"flagState,omitzero" api:"required"`
	paramObj
}

func (r PortalFlagStatePutRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow PortalFlagStatePutRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PortalFlagStatePutRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The state that the given flag should be in for this portal
type PortalFlagStatePutRequestFlagState string

const (
	PortalFlagStatePutRequestFlagStateAbsent PortalFlagStatePutRequestFlagState = "ABSENT"
	PortalFlagStatePutRequestFlagStateOff    PortalFlagStatePutRequestFlagState = "OFF"
	PortalFlagStatePutRequestFlagStateOn     PortalFlagStatePutRequestFlagState = "ON"
)

type PortalFlagStateResponse struct {
	// The ID of the app
	AppID int64 `json:"appId" api:"required"`
	// The name of the flag
	FlagName string `json:"flagName" api:"required"`
	// The state of the flag for this portal
	//
	// Any of "ABSENT", "OFF", "ON".
	FlagState PortalFlagStateResponseFlagState `json:"flagState" api:"required"`
	// The ID of the portal
	PortalID int64 `json:"portalId" api:"required"`
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

// The state of the flag for this portal
type PortalFlagStateResponseFlagState string

const (
	PortalFlagStateResponseFlagStateAbsent PortalFlagStateResponseFlagState = "ABSENT"
	PortalFlagStateResponseFlagStateOff    PortalFlagStateResponseFlagState = "OFF"
	PortalFlagStateResponseFlagStateOn     PortalFlagStateResponseFlagState = "ON"
)

type FeatureFlagUpdateParams struct {
	AppID          int64 `path:"appId" api:"required" json:"-"`
	FlagPutRequest FlagPutRequestParam
	paramObj
}

func (r FeatureFlagUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.FlagPutRequest)
}
func (r *FeatureFlagUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FeatureFlagDeleteParams struct {
	AppID int64 `path:"appId" api:"required" json:"-"`
	paramObj
}

type FeatureFlagDeletePortalStateParams struct {
	AppID    int64  `path:"appId" api:"required" json:"-"`
	FlagName string `path:"flagName" api:"required" json:"-"`
	paramObj
}

type FeatureFlagGetParams struct {
	AppID int64 `path:"appId" api:"required" json:"-"`
	paramObj
}

type FeatureFlagGetPortalStateParams struct {
	AppID    int64  `path:"appId" api:"required" json:"-"`
	FlagName string `path:"flagName" api:"required" json:"-"`
	paramObj
}

type FeatureFlagListPortalsParams struct {
	AppID int64 `path:"appId" api:"required" json:"-"`
	// The maximum number of results to display per page.
	Limit         param.Opt[int64] `query:"limit,omitzero" json:"-"`
	StartPortalID param.Opt[int64] `query:"startPortalId,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [FeatureFlagListPortalsParams]'s query parameters as
// `url.Values`.
func (r FeatureFlagListPortalsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type FeatureFlagUpdatePortalStateParams struct {
	AppID                     int64  `path:"appId" api:"required" json:"-"`
	FlagName                  string `path:"flagName" api:"required" json:"-"`
	PortalFlagStatePutRequest PortalFlagStatePutRequestParam
	paramObj
}

func (r FeatureFlagUpdatePortalStateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.PortalFlagStatePutRequest)
}
func (r *FeatureFlagUpdatePortalStateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
