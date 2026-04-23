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
	shimjson "github.com/stainless-sdks/hubspot-sdk-go/internal/encoding/json"
	"github.com/stainless-sdks/hubspot-sdk-go/internal/requestconfig"
	"github.com/stainless-sdks/hubspot-sdk-go/option"
)

// FeatureFlagBatchService contains methods and other services that help with
// interacting with the hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewFeatureFlagBatchService] method instead.
type FeatureFlagBatchService struct {
	options []option.RequestOption
}

// NewFeatureFlagBatchService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewFeatureFlagBatchService(opts ...option.RequestOption) (r FeatureFlagBatchService) {
	r = FeatureFlagBatchService{}
	r.options = opts
	return
}

// Delete an account-level flag state for multiple HubSpot accounts at once. Use
// this endpoint to manage flag exposure for groups of HubSpot accounts.
func (r *FeatureFlagBatchService) Delete(ctx context.Context, flagName string, params FeatureFlagBatchDeleteParams, opts ...option.RequestOption) (res *PortalFlagStateBatchResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if flagName == "" {
		err = errors.New("missing required flagName parameter")
		return nil, err
	}
	path := fmt.Sprintf("feature-flags/2026-03/%v/flags/%s/portals/batch/delete", params.AppID, url.PathEscape(flagName))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Set the portal flag state for multiple HubSpot accounts at once. Use this
// endpoint to manage flag exposure for groups of HubSpot accounts.
func (r *FeatureFlagBatchService) Upsert(ctx context.Context, flagName string, params FeatureFlagBatchUpsertParams, opts ...option.RequestOption) (res *PortalFlagStateBatchResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if flagName == "" {
		err = errors.New("missing required flagName parameter")
		return nil, err
	}
	path := fmt.Sprintf("feature-flags/2026-03/%v/flags/%s/portals/batch/upsert", params.AppID, url.PathEscape(flagName))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

type FeatureFlagBatchDeleteParams struct {
	AppID                             int64 `path:"appId" api:"required" json:"-"`
	PortalFlagStateBatchDeleteRequest PortalFlagStateBatchDeleteRequestParam
	paramObj
}

func (r FeatureFlagBatchDeleteParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.PortalFlagStateBatchDeleteRequest)
}
func (r *FeatureFlagBatchDeleteParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FeatureFlagBatchUpsertParams struct {
	AppID                          int64 `path:"appId" api:"required" json:"-"`
	PortalFlagStateBatchPutRequest PortalFlagStateBatchPutRequestParam
	paramObj
}

func (r FeatureFlagBatchUpsertParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.PortalFlagStateBatchPutRequest)
}
func (r *FeatureFlagBatchUpsertParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
