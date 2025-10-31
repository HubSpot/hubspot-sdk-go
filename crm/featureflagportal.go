// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package crm

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/stainless-sdks/hubspot-sdk-go/internal/apijson"
	"github.com/stainless-sdks/hubspot-sdk-go/internal/requestconfig"
	"github.com/stainless-sdks/hubspot-sdk-go/option"
	"github.com/stainless-sdks/hubspot-sdk-go/packages/param"
	"github.com/stainless-sdks/hubspot-sdk-go/packages/respjson"
)

// FeatureFlagPortalService contains methods and other services that help with
// interacting with the Hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewFeatureFlagPortalService] method instead.
type FeatureFlagPortalService struct {
	Options []option.RequestOption
}

// NewFeatureFlagPortalService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewFeatureFlagPortalService(opts ...option.RequestOption) (r FeatureFlagPortalService) {
	r = FeatureFlagPortalService{}
	r.Options = opts
	return
}

// Specify an account-level flag state for a specific HubSpot account.
func (r *FeatureFlagPortalService) Update(ctx context.Context, portalID int64, params FeatureFlagPortalUpdateParams, opts ...option.RequestOption) (res *FeatureFlagPortalUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.FlagName == "" {
		err = errors.New("missing required flagName parameter")
		return
	}
	path := fmt.Sprintf("feature-flags/v3/%v/flags/%s/portals/%v", params.AppID, params.FlagName, portalID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return
}

// Delete an account-level flag state for a specific HubSpot account. No request
// body is included.
func (r *FeatureFlagPortalService) Delete(ctx context.Context, portalID int64, body FeatureFlagPortalDeleteParams, opts ...option.RequestOption) (res *FeatureFlagPortalDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if body.FlagName == "" {
		err = errors.New("missing required flagName parameter")
		return
	}
	path := fmt.Sprintf("feature-flags/v3/%v/flags/%s/portals/%v", body.AppID, body.FlagName, portalID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Delete an account-level flag state for multiple HubSpot accounts at once. Use
// this endpoint to manage flag exposure for groups of HubSpot accounts.
func (r *FeatureFlagPortalService) BatchDelete(ctx context.Context, flagName string, params FeatureFlagPortalBatchDeleteParams, opts ...option.RequestOption) (res *FeatureFlagPortalBatchDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if flagName == "" {
		err = errors.New("missing required flagName parameter")
		return
	}
	path := fmt.Sprintf("feature-flags/v3/%v/flags/%s/portals/batch/delete", params.AppID, flagName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Set the portal flag state for multiple HubSpot accounts at once. Use this
// endpoint to manage flag exposure for groups of HubSpot accounts.
func (r *FeatureFlagPortalService) BatchUpsert(ctx context.Context, flagName string, params FeatureFlagPortalBatchUpsertParams, opts ...option.RequestOption) (res *FeatureFlagPortalBatchUpsertResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if flagName == "" {
		err = errors.New("missing required flagName parameter")
		return
	}
	path := fmt.Sprintf("feature-flags/v3/%v/flags/%s/portals/batch/upsert", params.AppID, flagName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Retrieve the account-level flag state of a specific HubSpot account.
func (r *FeatureFlagPortalService) Get(ctx context.Context, portalID int64, query FeatureFlagPortalGetParams, opts ...option.RequestOption) (res *FeatureFlagPortalGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if query.FlagName == "" {
		err = errors.New("missing required flagName parameter")
		return
	}
	path := fmt.Sprintf("feature-flags/v3/%v/flags/%s/portals/%v", query.AppID, query.FlagName, portalID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type FeatureFlagPortalUpdateResponse struct {
	AppID    int64  `json:"appId,required"`
	FlagName string `json:"flagName,required"`
	// Any of "OFF", "ON", "ABSENT".
	FlagState FeatureFlagPortalUpdateResponseFlagState `json:"flagState,required"`
	PortalID  int64                                    `json:"portalId,required"`
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
func (r FeatureFlagPortalUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *FeatureFlagPortalUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FeatureFlagPortalUpdateResponseFlagState string

const (
	FeatureFlagPortalUpdateResponseFlagStateOff    FeatureFlagPortalUpdateResponseFlagState = "OFF"
	FeatureFlagPortalUpdateResponseFlagStateOn     FeatureFlagPortalUpdateResponseFlagState = "ON"
	FeatureFlagPortalUpdateResponseFlagStateAbsent FeatureFlagPortalUpdateResponseFlagState = "ABSENT"
)

type FeatureFlagPortalDeleteResponse struct {
	AppID    int64  `json:"appId,required"`
	FlagName string `json:"flagName,required"`
	// Any of "OFF", "ON", "ABSENT".
	FlagState FeatureFlagPortalDeleteResponseFlagState `json:"flagState,required"`
	PortalID  int64                                    `json:"portalId,required"`
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
func (r FeatureFlagPortalDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *FeatureFlagPortalDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FeatureFlagPortalDeleteResponseFlagState string

const (
	FeatureFlagPortalDeleteResponseFlagStateOff    FeatureFlagPortalDeleteResponseFlagState = "OFF"
	FeatureFlagPortalDeleteResponseFlagStateOn     FeatureFlagPortalDeleteResponseFlagState = "ON"
	FeatureFlagPortalDeleteResponseFlagStateAbsent FeatureFlagPortalDeleteResponseFlagState = "ABSENT"
)

type FeatureFlagPortalBatchDeleteResponse struct {
	PortalFlagStates []FeatureFlagPortalBatchDeleteResponsePortalFlagState `json:"portalFlagStates,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PortalFlagStates respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FeatureFlagPortalBatchDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *FeatureFlagPortalBatchDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FeatureFlagPortalBatchDeleteResponsePortalFlagState struct {
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
func (r FeatureFlagPortalBatchDeleteResponsePortalFlagState) RawJSON() string { return r.JSON.raw }
func (r *FeatureFlagPortalBatchDeleteResponsePortalFlagState) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FeatureFlagPortalBatchUpsertResponse struct {
	PortalFlagStates []FeatureFlagPortalBatchUpsertResponsePortalFlagState `json:"portalFlagStates,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PortalFlagStates respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FeatureFlagPortalBatchUpsertResponse) RawJSON() string { return r.JSON.raw }
func (r *FeatureFlagPortalBatchUpsertResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FeatureFlagPortalBatchUpsertResponsePortalFlagState struct {
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
func (r FeatureFlagPortalBatchUpsertResponsePortalFlagState) RawJSON() string { return r.JSON.raw }
func (r *FeatureFlagPortalBatchUpsertResponsePortalFlagState) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FeatureFlagPortalGetResponse struct {
	AppID    int64  `json:"appId,required"`
	FlagName string `json:"flagName,required"`
	// Any of "OFF", "ON", "ABSENT".
	FlagState FeatureFlagPortalGetResponseFlagState `json:"flagState,required"`
	PortalID  int64                                 `json:"portalId,required"`
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
func (r FeatureFlagPortalGetResponse) RawJSON() string { return r.JSON.raw }
func (r *FeatureFlagPortalGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FeatureFlagPortalGetResponseFlagState string

const (
	FeatureFlagPortalGetResponseFlagStateOff    FeatureFlagPortalGetResponseFlagState = "OFF"
	FeatureFlagPortalGetResponseFlagStateOn     FeatureFlagPortalGetResponseFlagState = "ON"
	FeatureFlagPortalGetResponseFlagStateAbsent FeatureFlagPortalGetResponseFlagState = "ABSENT"
)

type FeatureFlagPortalUpdateParams struct {
	AppID    int64  `path:"appId,required" json:"-"`
	FlagName string `path:"flagName,required" json:"-"`
	// Any of "OFF", "ON", "ABSENT".
	FlagState FeatureFlagPortalUpdateParamsFlagState `json:"flagState,omitzero,required"`
	paramObj
}

func (r FeatureFlagPortalUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow FeatureFlagPortalUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FeatureFlagPortalUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FeatureFlagPortalUpdateParamsFlagState string

const (
	FeatureFlagPortalUpdateParamsFlagStateOff    FeatureFlagPortalUpdateParamsFlagState = "OFF"
	FeatureFlagPortalUpdateParamsFlagStateOn     FeatureFlagPortalUpdateParamsFlagState = "ON"
	FeatureFlagPortalUpdateParamsFlagStateAbsent FeatureFlagPortalUpdateParamsFlagState = "ABSENT"
)

type FeatureFlagPortalDeleteParams struct {
	AppID    int64  `path:"appId,required" json:"-"`
	FlagName string `path:"flagName,required" json:"-"`
	paramObj
}

type FeatureFlagPortalBatchDeleteParams struct {
	AppID     int64   `path:"appId,required" json:"-"`
	PortalIDs []int64 `json:"portalIds,omitzero,required"`
	paramObj
}

func (r FeatureFlagPortalBatchDeleteParams) MarshalJSON() (data []byte, err error) {
	type shadow FeatureFlagPortalBatchDeleteParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FeatureFlagPortalBatchDeleteParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FeatureFlagPortalBatchUpsertParams struct {
	AppID        int64                                           `path:"appId,required" json:"-"`
	PortalStates []FeatureFlagPortalBatchUpsertParamsPortalState `json:"portalStates,omitzero,required"`
	paramObj
}

func (r FeatureFlagPortalBatchUpsertParams) MarshalJSON() (data []byte, err error) {
	type shadow FeatureFlagPortalBatchUpsertParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FeatureFlagPortalBatchUpsertParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties FlagState, PortalID are required.
type FeatureFlagPortalBatchUpsertParamsPortalState struct {
	// Any of "OFF", "ON", "ABSENT".
	FlagState string `json:"flagState,omitzero,required"`
	PortalID  int64  `json:"portalId,required"`
	paramObj
}

func (r FeatureFlagPortalBatchUpsertParamsPortalState) MarshalJSON() (data []byte, err error) {
	type shadow FeatureFlagPortalBatchUpsertParamsPortalState
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FeatureFlagPortalBatchUpsertParamsPortalState) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[FeatureFlagPortalBatchUpsertParamsPortalState](
		"flagState", "OFF", "ON", "ABSENT",
	)
}

type FeatureFlagPortalGetParams struct {
	AppID    int64  `path:"appId,required" json:"-"`
	FlagName string `path:"flagName,required" json:"-"`
	paramObj
}
