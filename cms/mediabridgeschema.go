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

	"github.com/stainless-sdks/hubspot-sdk-go/crm"
	"github.com/stainless-sdks/hubspot-sdk-go/events"
	"github.com/stainless-sdks/hubspot-sdk-go/internal/apijson"
	"github.com/stainless-sdks/hubspot-sdk-go/internal/apiquery"
	shimjson "github.com/stainless-sdks/hubspot-sdk-go/internal/encoding/json"
	"github.com/stainless-sdks/hubspot-sdk-go/internal/requestconfig"
	"github.com/stainless-sdks/hubspot-sdk-go/option"
	"github.com/stainless-sdks/hubspot-sdk-go/packages/param"
	"github.com/stainless-sdks/hubspot-sdk-go/packages/respjson"
	"github.com/stainless-sdks/hubspot-sdk-go/shared"
)

// MediaBridgeSchemaService contains methods and other services that help with
// interacting with the hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMediaBridgeSchemaService] method instead.
type MediaBridgeSchemaService struct {
	Options []option.RequestOption
}

// NewMediaBridgeSchemaService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewMediaBridgeSchemaService(opts ...option.RequestOption) (r MediaBridgeSchemaService) {
	r = MediaBridgeSchemaService{}
	r.Options = opts
	return
}

// Update the schema for an existing object type
func (r *MediaBridgeSchemaService) Update(ctx context.Context, objectType string, params MediaBridgeSchemaUpdateParams, opts ...option.RequestOption) (res *crm.ObjectsSchemasObjectTypeDefinition, err error) {
	opts = slices.Concat(r.Options, opts)
	if objectType == "" {
		err = errors.New("missing required objectType parameter")
		return
	}
	path := fmt.Sprintf("media-bridge/v1/%v/schemas/%s", params.AppID, objectType)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return
}

// Get the schemas for all object types.
func (r *MediaBridgeSchemaService) List(ctx context.Context, appID int64, query MediaBridgeSchemaListParams, opts ...option.RequestOption) (res *MediaBridgeSchemaListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("media-bridge/v1/%v/schemas", appID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Create a new association definition for the specified object type.
func (r *MediaBridgeSchemaService) NewAssociation(ctx context.Context, objectType string, params MediaBridgeSchemaNewAssociationParams, opts ...option.RequestOption) (res *events.AssociationDefinition, err error) {
	opts = slices.Concat(r.Options, opts)
	if objectType == "" {
		err = errors.New("missing required objectType parameter")
		return
	}
	path := fmt.Sprintf("media-bridge/v1/%v/schemas/%s/associations", params.AppID, objectType)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Delete an existing association definition for an object type.
func (r *MediaBridgeSchemaService) DeleteAssociation(ctx context.Context, associationID string, body MediaBridgeSchemaDeleteAssociationParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if body.ObjectType == "" {
		err = errors.New("missing required objectType parameter")
		return
	}
	if associationID == "" {
		err = errors.New("missing required associationId parameter")
		return
	}
	path := fmt.Sprintf("media-bridge/v1/%v/schemas/%s/associations/%s", body.AppID, body.ObjectType, associationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Get the schema for a specified object type.
func (r *MediaBridgeSchemaService) Get(ctx context.Context, objectType string, query MediaBridgeSchemaGetParams, opts ...option.RequestOption) (res *crm.ObjectSchema, err error) {
	opts = slices.Concat(r.Options, opts)
	if objectType == "" {
		err = errors.New("missing required objectType parameter")
		return
	}
	path := fmt.Sprintf("media-bridge/v1/%v/schemas/%s", query.AppID, objectType)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type MediaBridgeSchemaListResponse struct {
	Results []crm.ObjectSchema `json:"results,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Results     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MediaBridgeSchemaListResponse) RawJSON() string { return r.JSON.raw }
func (r *MediaBridgeSchemaListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeSchemaUpdateParams struct {
	AppID int64 `path:"appId,required" json:"-"`
	// Defines attributes to update on an object type.
	ObjectTypeDefinitionPatch crm.ObjectTypeDefinitionPatchParam
	paramObj
}

func (r MediaBridgeSchemaUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.ObjectTypeDefinitionPatch)
}
func (r *MediaBridgeSchemaUpdateParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.ObjectTypeDefinitionPatch)
}

type MediaBridgeSchemaListParams struct {
	// Whether to return only results that have been archived.
	Archived param.Opt[bool] `query:"archived,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [MediaBridgeSchemaListParams]'s query parameters as
// `url.Values`.
func (r MediaBridgeSchemaListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type MediaBridgeSchemaNewAssociationParams struct {
	AppID                    int64 `path:"appId,required" json:"-"`
	AssociationDefinitionEgg shared.AssociationDefinitionEggParam
	paramObj
}

func (r MediaBridgeSchemaNewAssociationParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.AssociationDefinitionEgg)
}
func (r *MediaBridgeSchemaNewAssociationParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.AssociationDefinitionEgg)
}

type MediaBridgeSchemaDeleteAssociationParams struct {
	AppID      int64  `path:"appId,required" json:"-"`
	ObjectType string `path:"objectType,required" json:"-"`
	paramObj
}

type MediaBridgeSchemaGetParams struct {
	AppID int64 `path:"appId,required" json:"-"`
	paramObj
}
