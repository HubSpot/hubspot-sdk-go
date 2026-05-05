// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cms

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/HubSpot/hubspot-sdk-go/internal/apijson"
	shimjson "github.com/HubSpot/hubspot-sdk-go/internal/encoding/json"
	"github.com/HubSpot/hubspot-sdk-go/internal/requestconfig"
	"github.com/HubSpot/hubspot-sdk-go/option"
	"github.com/HubSpot/hubspot-sdk-go/shared"
)

// MediaBridgeBatchService contains methods and other services that help with
// interacting with the hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMediaBridgeBatchService] method instead.
type MediaBridgeBatchService struct {
	options []option.RequestOption
}

// NewMediaBridgeBatchService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewMediaBridgeBatchService(opts ...option.RequestOption) (r MediaBridgeBatchService) {
	r = MediaBridgeBatchService{}
	r.options = opts
	return
}

// Create a batch of properties of the specified object type.
func (r *MediaBridgeBatchService) New(ctx context.Context, objectType string, params MediaBridgeBatchNewParams, opts ...option.RequestOption) (res *BatchResponseProperty, err error) {
	opts = slices.Concat(r.options, opts)
	if objectType == "" {
		err = errors.New("missing required objectType parameter")
		return nil, err
	}
	path := fmt.Sprintf("media-bridge/2026-03/%v/properties/%s/batch/create", params.AppID, url.PathEscape(objectType))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Archive a batch of existing properties for the specified types.
func (r *MediaBridgeBatchService) Delete(ctx context.Context, objectType string, params MediaBridgeBatchDeleteParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if objectType == "" {
		err = errors.New("missing required objectType parameter")
		return err
	}
	path := fmt.Sprintf("media-bridge/2026-03/%v/properties/%s/batch/archive", params.AppID, url.PathEscape(objectType))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, nil, opts...)
	return err
}

// Get the details for a batch of properties for a specified object type.
func (r *MediaBridgeBatchService) Get(ctx context.Context, objectType string, params MediaBridgeBatchGetParams, opts ...option.RequestOption) (res *BatchResponseProperty, err error) {
	opts = slices.Concat(r.options, opts)
	if objectType == "" {
		err = errors.New("missing required objectType parameter")
		return nil, err
	}
	path := fmt.Sprintf("media-bridge/2026-03/%v/properties/%s/batch/read", params.AppID, url.PathEscape(objectType))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

type MediaBridgeBatchNewParams struct {
	AppID                    int64 `path:"appId" api:"required" json:"-"`
	BatchInputPropertyCreate shared.BatchInputPropertyCreateParam
	paramObj
}

func (r MediaBridgeBatchNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.BatchInputPropertyCreate)
}
func (r *MediaBridgeBatchNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeBatchDeleteParams struct {
	AppID                  int64 `path:"appId" api:"required" json:"-"`
	BatchInputPropertyName shared.BatchInputPropertyNameParam
	paramObj
}

func (r MediaBridgeBatchDeleteParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.BatchInputPropertyName)
}
func (r *MediaBridgeBatchDeleteParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeBatchGetParams struct {
	AppID                      int64 `path:"appId" api:"required" json:"-"`
	BatchReadInputPropertyName shared.BatchReadInputPropertyNameParam
	paramObj
}

func (r MediaBridgeBatchGetParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.BatchReadInputPropertyName)
}
func (r *MediaBridgeBatchGetParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
