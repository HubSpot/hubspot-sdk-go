// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cms

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
	"github.com/stainless-sdks/hubspot-sdk-go/shared"
)

// MediaBridgePropertyService contains methods and other services that help with
// interacting with the hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMediaBridgePropertyService] method instead.
type MediaBridgePropertyService struct {
	Options []option.RequestOption
}

// NewMediaBridgePropertyService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewMediaBridgePropertyService(opts ...option.RequestOption) (r MediaBridgePropertyService) {
	r = MediaBridgePropertyService{}
	r.Options = opts
	return
}

// Create a new property for the specified media type
func (r *MediaBridgePropertyService) New(ctx context.Context, objectType string, params MediaBridgePropertyNewParams, opts ...option.RequestOption) (res *shared.Property, err error) {
	opts = slices.Concat(r.Options, opts)
	if objectType == "" {
		err = errors.New("missing required objectType parameter")
		return
	}
	path := fmt.Sprintf("media-bridge/v1/%v/properties/%s", params.AppID, objectType)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Update an existing property for an object type.
func (r *MediaBridgePropertyService) Update(ctx context.Context, propertyName string, params MediaBridgePropertyUpdateParams, opts ...option.RequestOption) (res *shared.Property, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.ObjectType == "" {
		err = errors.New("missing required objectType parameter")
		return
	}
	if propertyName == "" {
		err = errors.New("missing required propertyName parameter")
		return
	}
	path := fmt.Sprintf("media-bridge/v1/%v/properties/%s/%s", params.AppID, params.ObjectType, propertyName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return
}

// Get the existing properties defined for a media object type.
func (r *MediaBridgePropertyService) List(ctx context.Context, objectType string, params MediaBridgePropertyListParams, opts ...option.RequestOption) (res *CollectionResponsePropertyNoPaging, err error) {
	opts = slices.Concat(r.Options, opts)
	if objectType == "" {
		err = errors.New("missing required objectType parameter")
		return
	}
	path := fmt.Sprintf("media-bridge/v1/%v/properties/%s", params.AppID, objectType)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return
}

// Delete an existing property for an object type.
func (r *MediaBridgePropertyService) Delete(ctx context.Context, propertyName string, body MediaBridgePropertyDeleteParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if body.ObjectType == "" {
		err = errors.New("missing required objectType parameter")
		return
	}
	if propertyName == "" {
		err = errors.New("missing required propertyName parameter")
		return
	}
	path := fmt.Sprintf("media-bridge/v1/%v/properties/%s/%s", body.AppID, body.ObjectType, propertyName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Create a batch of properties of the specified object type.
func (r *MediaBridgePropertyService) NewBatch(ctx context.Context, objectType string, params MediaBridgePropertyNewBatchParams, opts ...option.RequestOption) (res *shared.BatchResponseProperty, err error) {
	opts = slices.Concat(r.Options, opts)
	if objectType == "" {
		err = errors.New("missing required objectType parameter")
		return
	}
	path := fmt.Sprintf("media-bridge/v1/%v/properties/%s/batch/create", params.AppID, objectType)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Archive a batch of existing properties for the specified types.
func (r *MediaBridgePropertyService) DeleteBatch(ctx context.Context, objectType string, params MediaBridgePropertyDeleteBatchParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if objectType == "" {
		err = errors.New("missing required objectType parameter")
		return
	}
	path := fmt.Sprintf("media-bridge/v1/%v/properties/%s/batch/archive", params.AppID, objectType)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, nil, opts...)
	return
}

// Get the details for an existing property by name.
func (r *MediaBridgePropertyService) Get(ctx context.Context, propertyName string, params MediaBridgePropertyGetParams, opts ...option.RequestOption) (res *shared.Property, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.ObjectType == "" {
		err = errors.New("missing required objectType parameter")
		return
	}
	if propertyName == "" {
		err = errors.New("missing required propertyName parameter")
		return
	}
	path := fmt.Sprintf("media-bridge/v1/%v/properties/%s/%s", params.AppID, params.ObjectType, propertyName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return
}

// Get the details for a batch of properties for a specified object type.
func (r *MediaBridgePropertyService) GetBatch(ctx context.Context, objectType string, params MediaBridgePropertyGetBatchParams, opts ...option.RequestOption) (res *shared.BatchResponseProperty, err error) {
	opts = slices.Concat(r.Options, opts)
	if objectType == "" {
		err = errors.New("missing required objectType parameter")
		return
	}
	path := fmt.Sprintf("media-bridge/v1/%v/properties/%s/batch/read", params.AppID, objectType)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

type MediaBridgePropertyNewParams struct {
	AppID          int64 `path:"appId,required" json:"-"`
	PropertyCreate shared.PropertyCreateParam
	paramObj
}

func (r MediaBridgePropertyNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.PropertyCreate)
}
func (r *MediaBridgePropertyNewParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.PropertyCreate)
}

type MediaBridgePropertyUpdateParams struct {
	AppID                     int64  `path:"appId,required" json:"-"`
	ObjectType                string `path:"objectType,required" json:"-"`
	MediaBridgePropertyUpdate MediaBridgePropertyUpdateParam
	paramObj
}

func (r MediaBridgePropertyUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.MediaBridgePropertyUpdate)
}
func (r *MediaBridgePropertyUpdateParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.MediaBridgePropertyUpdate)
}

type MediaBridgePropertyListParams struct {
	AppID int64 `path:"appId,required" json:"-"`
	// Whether to return only results that have been archived.
	Archived param.Opt[bool] `query:"archived,omitzero" json:"-"`
	// Filter the response to the specified properties.
	Properties param.Opt[string] `query:"properties,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [MediaBridgePropertyListParams]'s query parameters as
// `url.Values`.
func (r MediaBridgePropertyListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type MediaBridgePropertyDeleteParams struct {
	AppID      int64  `path:"appId,required" json:"-"`
	ObjectType string `path:"objectType,required" json:"-"`
	paramObj
}

type MediaBridgePropertyNewBatchParams struct {
	AppID                    int64 `path:"appId,required" json:"-"`
	BatchInputPropertyCreate shared.BatchInputPropertyCreateParam
	paramObj
}

func (r MediaBridgePropertyNewBatchParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.BatchInputPropertyCreate)
}
func (r *MediaBridgePropertyNewBatchParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.BatchInputPropertyCreate)
}

type MediaBridgePropertyDeleteBatchParams struct {
	AppID                  int64 `path:"appId,required" json:"-"`
	BatchInputPropertyName shared.BatchInputPropertyNameParam
	paramObj
}

func (r MediaBridgePropertyDeleteBatchParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.BatchInputPropertyName)
}
func (r *MediaBridgePropertyDeleteBatchParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.BatchInputPropertyName)
}

type MediaBridgePropertyGetParams struct {
	AppID      int64  `path:"appId,required" json:"-"`
	ObjectType string `path:"objectType,required" json:"-"`
	// Whether to return only results that have been archived.
	Archived param.Opt[bool] `query:"archived,omitzero" json:"-"`
	// Limit the response to only include the specified properties.
	Properties param.Opt[string] `query:"properties,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [MediaBridgePropertyGetParams]'s query parameters as
// `url.Values`.
func (r MediaBridgePropertyGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type MediaBridgePropertyGetBatchParams struct {
	AppID                      int64 `path:"appId,required" json:"-"`
	BatchReadInputPropertyName shared.BatchReadInputPropertyNameParam
	paramObj
}

func (r MediaBridgePropertyGetBatchParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.BatchReadInputPropertyName)
}
func (r *MediaBridgePropertyGetBatchParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.BatchReadInputPropertyName)
}
