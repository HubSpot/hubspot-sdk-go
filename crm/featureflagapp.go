// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package crm

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/stainless-sdks/hubspot-sdk-go/internal/apiquery"
	shimjson "github.com/stainless-sdks/hubspot-sdk-go/internal/encoding/json"
	"github.com/stainless-sdks/hubspot-sdk-go/internal/requestconfig"
	"github.com/stainless-sdks/hubspot-sdk-go/option"
	"github.com/stainless-sdks/hubspot-sdk-go/packages/param"
)

// FeatureFlagAppService contains methods and other services that help with
// interacting with the hubspot API.
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
func (r *FeatureFlagAppService) Update(ctx context.Context, flagName string, params FeatureFlagAppUpdateParams, opts ...option.RequestOption) (res *FlagResponse, err error) {
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
func (r *FeatureFlagAppService) Delete(ctx context.Context, flagName string, body FeatureFlagAppDeleteParams, opts ...option.RequestOption) (res *FlagResponse, err error) {
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
func (r *FeatureFlagAppService) Get(ctx context.Context, flagName string, query FeatureFlagAppGetParams, opts ...option.RequestOption) (res *FlagResponse, err error) {
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
func (r *FeatureFlagAppService) ListPortals(ctx context.Context, flagName string, params FeatureFlagAppListPortalsParams, opts ...option.RequestOption) (res *PortalFlagStateBatchResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if flagName == "" {
		err = errors.New("missing required flagName parameter")
		return
	}
	path := fmt.Sprintf("feature-flags/v3/%v/flags/%s/portals", params.AppID, flagName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return
}

type FeatureFlagAppUpdateParams struct {
	AppID          int64 `path:"appId,required" json:"-"`
	FlagPutRequest FlagPutRequestParam
	paramObj
}

func (r FeatureFlagAppUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.FlagPutRequest)
}
func (r *FeatureFlagAppUpdateParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.FlagPutRequest)
}

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
