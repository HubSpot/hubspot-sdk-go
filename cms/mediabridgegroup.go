// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cms

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/stainless-sdks/hubspot-sdk-go/crm"
	"github.com/stainless-sdks/hubspot-sdk-go/internal/apijson"
	shimjson "github.com/stainless-sdks/hubspot-sdk-go/internal/encoding/json"
	"github.com/stainless-sdks/hubspot-sdk-go/internal/requestconfig"
	"github.com/stainless-sdks/hubspot-sdk-go/option"
	"github.com/stainless-sdks/hubspot-sdk-go/packages/respjson"
	"github.com/stainless-sdks/hubspot-sdk-go/shared"
)

// MediaBridgeGroupService contains methods and other services that help with
// interacting with the Hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMediaBridgeGroupService] method instead.
type MediaBridgeGroupService struct {
	Options []option.RequestOption
}

// NewMediaBridgeGroupService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewMediaBridgeGroupService(opts ...option.RequestOption) (r MediaBridgeGroupService) {
	r = MediaBridgeGroupService{}
	r.Options = opts
	return
}

// Create a new property group for the specified object type.
func (r *MediaBridgeGroupService) New(ctx context.Context, objectType string, params MediaBridgeGroupNewParams, opts ...option.RequestOption) (res *crm.PropertyGroup, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.AppID == "" {
		err = errors.New("missing required appId parameter")
		return
	}
	if objectType == "" {
		err = errors.New("missing required objectType parameter")
		return
	}
	path := fmt.Sprintf("media-bridge/v1/%s/properties/%s/groups", params.AppID, objectType)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Get the property groups for a specified object type.
func (r *MediaBridgeGroupService) List(ctx context.Context, objectType string, query MediaBridgeGroupListParams, opts ...option.RequestOption) (res *MediaBridgeGroupListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if query.AppID == "" {
		err = errors.New("missing required appId parameter")
		return
	}
	if objectType == "" {
		err = errors.New("missing required objectType parameter")
		return
	}
	path := fmt.Sprintf("media-bridge/v1/%s/properties/%s/groups", query.AppID, objectType)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Delete an existing property group by name
func (r *MediaBridgeGroupService) DeleteByName(ctx context.Context, groupName string, body MediaBridgeGroupDeleteByNameParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if body.AppID == "" {
		err = errors.New("missing required appId parameter")
		return
	}
	if body.ObjectType == "" {
		err = errors.New("missing required objectType parameter")
		return
	}
	if groupName == "" {
		err = errors.New("missing required groupName parameter")
		return
	}
	path := fmt.Sprintf("media-bridge/v1/%s/properties/%s/groups/%s", body.AppID, body.ObjectType, groupName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Get the details of an existing property group by name.
func (r *MediaBridgeGroupService) GetByName(ctx context.Context, groupName string, query MediaBridgeGroupGetByNameParams, opts ...option.RequestOption) (res *crm.PropertyGroup, err error) {
	opts = slices.Concat(r.Options, opts)
	if query.AppID == "" {
		err = errors.New("missing required appId parameter")
		return
	}
	if query.ObjectType == "" {
		err = errors.New("missing required objectType parameter")
		return
	}
	if groupName == "" {
		err = errors.New("missing required groupName parameter")
		return
	}
	path := fmt.Sprintf("media-bridge/v1/%s/properties/%s/groups/%s", query.AppID, query.ObjectType, groupName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update an existing property group by name.
func (r *MediaBridgeGroupService) UpdateByName(ctx context.Context, groupName string, params MediaBridgeGroupUpdateByNameParams, opts ...option.RequestOption) (res *crm.PropertyGroup, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.AppID == "" {
		err = errors.New("missing required appId parameter")
		return
	}
	if params.ObjectType == "" {
		err = errors.New("missing required objectType parameter")
		return
	}
	if groupName == "" {
		err = errors.New("missing required groupName parameter")
		return
	}
	path := fmt.Sprintf("media-bridge/v1/%s/properties/%s/groups/%s", params.AppID, params.ObjectType, groupName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return
}

type MediaBridgeGroupListResponse struct {
	Results []crm.PropertyGroup `json:"results,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Results     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MediaBridgeGroupListResponse) RawJSON() string { return r.JSON.raw }
func (r *MediaBridgeGroupListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeGroupNewParams struct {
	AppID               string `path:"appId,required" json:"-"`
	PropertyGroupCreate shared.PropertyGroupCreateParam
	paramObj
}

func (r MediaBridgeGroupNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.PropertyGroupCreate)
}
func (r *MediaBridgeGroupNewParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.PropertyGroupCreate)
}

type MediaBridgeGroupListParams struct {
	AppID string `path:"appId,required" json:"-"`
	paramObj
}

type MediaBridgeGroupDeleteByNameParams struct {
	AppID      string `path:"appId,required" json:"-"`
	ObjectType string `path:"objectType,required" json:"-"`
	paramObj
}

type MediaBridgeGroupGetByNameParams struct {
	AppID      string `path:"appId,required" json:"-"`
	ObjectType string `path:"objectType,required" json:"-"`
	paramObj
}

type MediaBridgeGroupUpdateByNameParams struct {
	AppID               string `path:"appId,required" json:"-"`
	ObjectType          string `path:"objectType,required" json:"-"`
	PropertyGroupUpdate shared.PropertyGroupUpdateParam
	paramObj
}

func (r MediaBridgeGroupUpdateByNameParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.PropertyGroupUpdate)
}
func (r *MediaBridgeGroupUpdateByNameParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.PropertyGroupUpdate)
}
