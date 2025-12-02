// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package crm

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"slices"

	shimjson "github.com/stainless-sdks/hubspot-sdk-go/internal/encoding/json"
	"github.com/stainless-sdks/hubspot-sdk-go/internal/requestconfig"
	"github.com/stainless-sdks/hubspot-sdk-go/option"
)

// FeatureFlagPortalService contains methods and other services that help with
// interacting with the hubspot API.
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
func (r *FeatureFlagPortalService) Update(ctx context.Context, portalID int64, params FeatureFlagPortalUpdateParams, opts ...option.RequestOption) (res *PortalFlagStateResponse, err error) {
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
func (r *FeatureFlagPortalService) Delete(ctx context.Context, portalID int64, body FeatureFlagPortalDeleteParams, opts ...option.RequestOption) (res *PortalFlagStateResponse, err error) {
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
func (r *FeatureFlagPortalService) BatchDelete(ctx context.Context, flagName string, params FeatureFlagPortalBatchDeleteParams, opts ...option.RequestOption) (res *PortalFlagStateBatchResponse, err error) {
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
func (r *FeatureFlagPortalService) BatchUpsert(ctx context.Context, flagName string, params FeatureFlagPortalBatchUpsertParams, opts ...option.RequestOption) (res *PortalFlagStateBatchResponse, err error) {
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
func (r *FeatureFlagPortalService) Get(ctx context.Context, portalID int64, query FeatureFlagPortalGetParams, opts ...option.RequestOption) (res *PortalFlagStateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if query.FlagName == "" {
		err = errors.New("missing required flagName parameter")
		return
	}
	path := fmt.Sprintf("feature-flags/v3/%v/flags/%s/portals/%v", query.AppID, query.FlagName, portalID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type FeatureFlagPortalUpdateParams struct {
	AppID                     int64  `path:"appId,required" json:"-"`
	FlagName                  string `path:"flagName,required" json:"-"`
	PortalFlagStatePutRequest PortalFlagStatePutRequestParam
	paramObj
}

func (r FeatureFlagPortalUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.PortalFlagStatePutRequest)
}
func (r *FeatureFlagPortalUpdateParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.PortalFlagStatePutRequest)
}

type FeatureFlagPortalDeleteParams struct {
	AppID    int64  `path:"appId,required" json:"-"`
	FlagName string `path:"flagName,required" json:"-"`
	paramObj
}

type FeatureFlagPortalBatchDeleteParams struct {
	AppID                             int64 `path:"appId,required" json:"-"`
	PortalFlagStateBatchDeleteRequest PortalFlagStateBatchDeleteRequestParam
	paramObj
}

func (r FeatureFlagPortalBatchDeleteParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.PortalFlagStateBatchDeleteRequest)
}
func (r *FeatureFlagPortalBatchDeleteParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.PortalFlagStateBatchDeleteRequest)
}

type FeatureFlagPortalBatchUpsertParams struct {
	AppID                          int64 `path:"appId,required" json:"-"`
	PortalFlagStateBatchPutRequest PortalFlagStateBatchPutRequestParam
	paramObj
}

func (r FeatureFlagPortalBatchUpsertParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.PortalFlagStateBatchPutRequest)
}
func (r *FeatureFlagPortalBatchUpsertParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.PortalFlagStateBatchPutRequest)
}

type FeatureFlagPortalGetParams struct {
	AppID    int64  `path:"appId,required" json:"-"`
	FlagName string `path:"flagName,required" json:"-"`
	paramObj
}
