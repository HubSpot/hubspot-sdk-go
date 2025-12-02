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
	"github.com/stainless-sdks/hubspot-sdk-go/shared"
)

// PropertyGroupService contains methods and other services that help with
// interacting with the hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPropertyGroupService] method instead.
type PropertyGroupService struct {
	Options []option.RequestOption
}

// NewPropertyGroupService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewPropertyGroupService(opts ...option.RequestOption) (r PropertyGroupService) {
	r = PropertyGroupService{}
	r.Options = opts
	return
}

// Create and return a copy of a new property group.
func (r *PropertyGroupService) New(ctx context.Context, objectType string, body PropertyGroupNewParams, opts ...option.RequestOption) (res *CreatedResponsePropertyGroup, err error) {
	opts = slices.Concat(r.Options, opts)
	if objectType == "" {
		err = errors.New("missing required objectType parameter")
		return
	}
	path := fmt.Sprintf("crm/v3/properties/%s/groups", objectType)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Perform a partial update of a property group identified by {groupName}. Provided
// fields will be overwritten.
func (r *PropertyGroupService) Update(ctx context.Context, groupName string, params PropertyGroupUpdateParams, opts ...option.RequestOption) (res *PropertyGroup, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.ObjectType == "" {
		err = errors.New("missing required objectType parameter")
		return
	}
	if groupName == "" {
		err = errors.New("missing required groupName parameter")
		return
	}
	path := fmt.Sprintf("crm/v3/properties/%s/groups/%s", params.ObjectType, groupName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return
}

// Read all existing property groups for the specified object type and HubSpot
// account.
func (r *PropertyGroupService) List(ctx context.Context, objectType string, query PropertyGroupListParams, opts ...option.RequestOption) (res *CollectionResponsePropertyGroup, err error) {
	opts = slices.Concat(r.Options, opts)
	if objectType == "" {
		err = errors.New("missing required objectType parameter")
		return
	}
	path := fmt.Sprintf("crm/v3/properties/%s/groups", objectType)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Move a property group identified by {groupName} to the recycling bin.
func (r *PropertyGroupService) Delete(ctx context.Context, groupName string, body PropertyGroupDeleteParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if body.ObjectType == "" {
		err = errors.New("missing required objectType parameter")
		return
	}
	if groupName == "" {
		err = errors.New("missing required groupName parameter")
		return
	}
	path := fmt.Sprintf("crm/v3/properties/%s/groups/%s", body.ObjectType, groupName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Read a property group identified by {groupName}.
func (r *PropertyGroupService) Get(ctx context.Context, groupName string, params PropertyGroupGetParams, opts ...option.RequestOption) (res *PropertyGroup, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.ObjectType == "" {
		err = errors.New("missing required objectType parameter")
		return
	}
	if groupName == "" {
		err = errors.New("missing required groupName parameter")
		return
	}
	path := fmt.Sprintf("crm/v3/properties/%s/groups/%s", params.ObjectType, groupName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return
}

type PropertyGroupNewParams struct {
	PropertyGroupCreate shared.PropertyGroupCreateParam
	paramObj
}

func (r PropertyGroupNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.PropertyGroupCreate)
}
func (r *PropertyGroupNewParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.PropertyGroupCreate)
}

type PropertyGroupUpdateParams struct {
	ObjectType          string `path:"objectType,required" json:"-"`
	PropertyGroupUpdate shared.PropertyGroupUpdateParam
	paramObj
}

func (r PropertyGroupUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.PropertyGroupUpdate)
}
func (r *PropertyGroupUpdateParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.PropertyGroupUpdate)
}

type PropertyGroupListParams struct {
	Locale param.Opt[string] `query:"locale,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [PropertyGroupListParams]'s query parameters as
// `url.Values`.
func (r PropertyGroupListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type PropertyGroupDeleteParams struct {
	ObjectType string `path:"objectType,required" json:"-"`
	paramObj
}

type PropertyGroupGetParams struct {
	ObjectType string            `path:"objectType,required" json:"-"`
	Locale     param.Opt[string] `query:"locale,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [PropertyGroupGetParams]'s query parameters as `url.Values`.
func (r PropertyGroupGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
