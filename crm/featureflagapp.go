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
	"github.com/stainless-sdks/hubspot-sdk-go/internal/requestconfig"
	"github.com/stainless-sdks/hubspot-sdk-go/option"
	"github.com/stainless-sdks/hubspot-sdk-go/packages/param"
	"github.com/stainless-sdks/hubspot-sdk-go/packages/respjson"
)

// FeatureFlagAppService contains methods and other services that help with
// interacting with the Hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewFeatureFlagAppService] method instead.
type FeatureFlagAppService struct {
	Options []option.RequestOption
}

// NewFeatureFlagAppService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewFeatureFlagAppService(opts ...option.RequestOption) (r FeatureFlagAppService) {
	r = FeatureFlagAppService{}
	r.Options = opts
	return
}

// Set a feature flag for an app. For example, update the `hs-hide-crm-cards`
// flag's `defaultState` to `ON` to hide classic CRM cards from new installs.
func (r *FeatureFlagAppService) Update(ctx context.Context, flagName string, params FeatureFlagAppUpdateParams, opts ...option.RequestOption) (res *FeatureFlagAppUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if flagName == "" {
		err = errors.New("missing required flagName parameter")
		return
	}
	path := fmt.Sprintf("feature-flags/v3/%v/flags/%s", params.AppID, flagName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return
}

// Delete a feature flag in an app. For example, delete the `hs-release-app-cards`
// flag after all accounts have been migrated.
func (r *FeatureFlagAppService) Delete(ctx context.Context, flagName string, body FeatureFlagAppDeleteParams, opts ...option.RequestOption) (res *FeatureFlagAppDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if flagName == "" {
		err = errors.New("missing required flagName parameter")
		return
	}
	path := fmt.Sprintf("feature-flags/v3/%v/flags/%s", body.AppID, flagName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Retrieve the current status of the app's feature flags. No request body is
// included.
func (r *FeatureFlagAppService) Get(ctx context.Context, flagName string, query FeatureFlagAppGetParams, opts ...option.RequestOption) (res *FeatureFlagAppGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if flagName == "" {
		err = errors.New("missing required flagName parameter")
		return
	}
	path := fmt.Sprintf("feature-flags/v3/%v/flags/%s", query.AppID, flagName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Retrieve a list of HubSpot accounts with an account-level flag setting for the
// specified app. No request body is included.
func (r *FeatureFlagAppService) ListPortals(ctx context.Context, flagName string, params FeatureFlagAppListPortalsParams, opts ...option.RequestOption) (res *FeatureFlagAppListPortalsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if flagName == "" {
		err = errors.New("missing required flagName parameter")
		return
	}
	path := fmt.Sprintf("feature-flags/v3/%v/flags/%s/portals", params.AppID, flagName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return
}

type FeatureFlagAppUpdateResponse struct {
	AppID int64 `json:"appId,required"`
	// Any of "OFF", "ON", "ABSENT".
	DefaultState FeatureFlagAppUpdateResponseDefaultState `json:"defaultState,required"`
	FlagName     string                                   `json:"flagName,required"`
	// Any of "OFF", "ON", "ABSENT".
	OverrideState FeatureFlagAppUpdateResponseOverrideState `json:"overrideState"`
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
func (r FeatureFlagAppUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *FeatureFlagAppUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FeatureFlagAppUpdateResponseDefaultState string

const (
	FeatureFlagAppUpdateResponseDefaultStateOff    FeatureFlagAppUpdateResponseDefaultState = "OFF"
	FeatureFlagAppUpdateResponseDefaultStateOn     FeatureFlagAppUpdateResponseDefaultState = "ON"
	FeatureFlagAppUpdateResponseDefaultStateAbsent FeatureFlagAppUpdateResponseDefaultState = "ABSENT"
)

type FeatureFlagAppUpdateResponseOverrideState string

const (
	FeatureFlagAppUpdateResponseOverrideStateOff    FeatureFlagAppUpdateResponseOverrideState = "OFF"
	FeatureFlagAppUpdateResponseOverrideStateOn     FeatureFlagAppUpdateResponseOverrideState = "ON"
	FeatureFlagAppUpdateResponseOverrideStateAbsent FeatureFlagAppUpdateResponseOverrideState = "ABSENT"
)

type FeatureFlagAppDeleteResponse struct {
	AppID int64 `json:"appId,required"`
	// Any of "OFF", "ON", "ABSENT".
	DefaultState FeatureFlagAppDeleteResponseDefaultState `json:"defaultState,required"`
	FlagName     string                                   `json:"flagName,required"`
	// Any of "OFF", "ON", "ABSENT".
	OverrideState FeatureFlagAppDeleteResponseOverrideState `json:"overrideState"`
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
func (r FeatureFlagAppDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *FeatureFlagAppDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FeatureFlagAppDeleteResponseDefaultState string

const (
	FeatureFlagAppDeleteResponseDefaultStateOff    FeatureFlagAppDeleteResponseDefaultState = "OFF"
	FeatureFlagAppDeleteResponseDefaultStateOn     FeatureFlagAppDeleteResponseDefaultState = "ON"
	FeatureFlagAppDeleteResponseDefaultStateAbsent FeatureFlagAppDeleteResponseDefaultState = "ABSENT"
)

type FeatureFlagAppDeleteResponseOverrideState string

const (
	FeatureFlagAppDeleteResponseOverrideStateOff    FeatureFlagAppDeleteResponseOverrideState = "OFF"
	FeatureFlagAppDeleteResponseOverrideStateOn     FeatureFlagAppDeleteResponseOverrideState = "ON"
	FeatureFlagAppDeleteResponseOverrideStateAbsent FeatureFlagAppDeleteResponseOverrideState = "ABSENT"
)

type FeatureFlagAppGetResponse struct {
	AppID int64 `json:"appId,required"`
	// Any of "OFF", "ON", "ABSENT".
	DefaultState FeatureFlagAppGetResponseDefaultState `json:"defaultState,required"`
	FlagName     string                                `json:"flagName,required"`
	// Any of "OFF", "ON", "ABSENT".
	OverrideState FeatureFlagAppGetResponseOverrideState `json:"overrideState"`
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
func (r FeatureFlagAppGetResponse) RawJSON() string { return r.JSON.raw }
func (r *FeatureFlagAppGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FeatureFlagAppGetResponseDefaultState string

const (
	FeatureFlagAppGetResponseDefaultStateOff    FeatureFlagAppGetResponseDefaultState = "OFF"
	FeatureFlagAppGetResponseDefaultStateOn     FeatureFlagAppGetResponseDefaultState = "ON"
	FeatureFlagAppGetResponseDefaultStateAbsent FeatureFlagAppGetResponseDefaultState = "ABSENT"
)

type FeatureFlagAppGetResponseOverrideState string

const (
	FeatureFlagAppGetResponseOverrideStateOff    FeatureFlagAppGetResponseOverrideState = "OFF"
	FeatureFlagAppGetResponseOverrideStateOn     FeatureFlagAppGetResponseOverrideState = "ON"
	FeatureFlagAppGetResponseOverrideStateAbsent FeatureFlagAppGetResponseOverrideState = "ABSENT"
)

type FeatureFlagAppListPortalsResponse struct {
	PortalFlagStates []FeatureFlagAppListPortalsResponsePortalFlagState `json:"portalFlagStates,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PortalFlagStates respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FeatureFlagAppListPortalsResponse) RawJSON() string { return r.JSON.raw }
func (r *FeatureFlagAppListPortalsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FeatureFlagAppListPortalsResponsePortalFlagState struct {
	AppID    int64  `json:"appId,required"`
	FlagName string `json:"flagName,required"`
	// Any of "OFF", "ON", "ABSENT".
	FlagState string `json:"flagState,required"`
	PortalID  int64  `json:"portalId,required"`
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
func (r FeatureFlagAppListPortalsResponsePortalFlagState) RawJSON() string { return r.JSON.raw }
func (r *FeatureFlagAppListPortalsResponsePortalFlagState) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FeatureFlagAppUpdateParams struct {
	AppID int64 `path:"appId,required" json:"-"`
	// Any of "OFF", "ON", "ABSENT".
	DefaultState FeatureFlagAppUpdateParamsDefaultState `json:"defaultState,omitzero,required"`
	// Any of "OFF", "ON", "ABSENT".
	OverrideState FeatureFlagAppUpdateParamsOverrideState `json:"overrideState,omitzero"`
	paramObj
}

func (r FeatureFlagAppUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow FeatureFlagAppUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FeatureFlagAppUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FeatureFlagAppUpdateParamsDefaultState string

const (
	FeatureFlagAppUpdateParamsDefaultStateOff    FeatureFlagAppUpdateParamsDefaultState = "OFF"
	FeatureFlagAppUpdateParamsDefaultStateOn     FeatureFlagAppUpdateParamsDefaultState = "ON"
	FeatureFlagAppUpdateParamsDefaultStateAbsent FeatureFlagAppUpdateParamsDefaultState = "ABSENT"
)

type FeatureFlagAppUpdateParamsOverrideState string

const (
	FeatureFlagAppUpdateParamsOverrideStateOff    FeatureFlagAppUpdateParamsOverrideState = "OFF"
	FeatureFlagAppUpdateParamsOverrideStateOn     FeatureFlagAppUpdateParamsOverrideState = "ON"
	FeatureFlagAppUpdateParamsOverrideStateAbsent FeatureFlagAppUpdateParamsOverrideState = "ABSENT"
)

type FeatureFlagAppDeleteParams struct {
	AppID int64 `path:"appId,required" json:"-"`
	paramObj
}

type FeatureFlagAppGetParams struct {
	AppID int64 `path:"appId,required" json:"-"`
	paramObj
}

type FeatureFlagAppListPortalsParams struct {
	AppID int64 `path:"appId,required" json:"-"`
	// The maximum number of results to return in a single request.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// The initial account ID for listing, enabling pagination.
	StartPortalID param.Opt[int64] `query:"startPortalId,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [FeatureFlagAppListPortalsParams]'s query parameters as
// `url.Values`.
func (r FeatureFlagAppListPortalsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
