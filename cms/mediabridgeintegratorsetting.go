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
)

// MediaBridgeIntegratorSettingService contains methods and other services that
// help with interacting with the hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMediaBridgeIntegratorSettingService] method instead.
type MediaBridgeIntegratorSettingService struct {
	Options []option.RequestOption
}

// NewMediaBridgeIntegratorSettingService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewMediaBridgeIntegratorSettingService(opts ...option.RequestOption) (r MediaBridgeIntegratorSettingService) {
	r = MediaBridgeIntegratorSettingService{}
	r.Options = opts
	return
}

// Create a new media object type
func (r *MediaBridgeIntegratorSettingService) NewObjectDefinition(ctx context.Context, appID int64, body MediaBridgeIntegratorSettingNewObjectDefinitionParams, opts ...option.RequestOption) (res *BulkIntegratorObjectCreationResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("media-bridge/v1/%v/settings/object-definitions", appID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Set up a new oEmbed domain for your media bridge app.
func (r *MediaBridgeIntegratorSettingService) NewOembedDomain(ctx context.Context, appID int64, body MediaBridgeIntegratorSettingNewOembedDomainParams, opts ...option.RequestOption) (res *IntegratorOEmbedDomainModel, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("media-bridge/v1/%v/settings/oembed-domains", appID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Delete an existing oEmbed domain.
func (r *MediaBridgeIntegratorSettingService) DeleteOembedDomain(ctx context.Context, appID int64, body MediaBridgeIntegratorSettingDeleteOembedDomainParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := fmt.Sprintf("media-bridge/v1/%v/settings/oembed-domains", appID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, nil, opts...)
	return
}

// Get the visibility settings for media bridge events for your apps.
func (r *MediaBridgeIntegratorSettingService) GetEventVisibilitySettings(ctx context.Context, appID int64, opts ...option.RequestOption) (res *EventVisibilityResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("media-bridge/v1/%v/settings/event-visibility", appID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Get the existing objects types that belong to the specified media type.
func (r *MediaBridgeIntegratorSettingService) GetObjectDefinitionsByMediaType(ctx context.Context, mediaType MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeParamsMediaType, params MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeParams, opts ...option.RequestOption) (res *ObjectDefinitionResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("media-bridge/v1/%v/settings/object-definitions/%v", params.AppID, mediaType)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return
}

// Get the details for an existing oEmbed domain.
func (r *MediaBridgeIntegratorSettingService) GetOembedDomain(ctx context.Context, oEmbedDomainID string, query MediaBridgeIntegratorSettingGetOembedDomainParams, opts ...option.RequestOption) (res *IntegratorOEmbedDomainModel, err error) {
	opts = slices.Concat(r.Options, opts)
	if oEmbedDomainID == "" {
		err = errors.New("missing required oEmbedDomainId parameter")
		return
	}
	path := fmt.Sprintf("media-bridge/v1/%v/settings/oembed-domains/%s", query.AppID, oEmbedDomainID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Get the details for existing oEmbed domains for your app
func (r *MediaBridgeIntegratorSettingService) ListOembedDomains(ctx context.Context, appID int64, query MediaBridgeIntegratorSettingListOembedDomainsParams, opts ...option.RequestOption) (res *OEmbedDomainsCollectionResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("media-bridge/v1/%v/settings/oembed-domains", appID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Register the name that your app will display when a user is selecting media
// bridge items.
//
// Deprecated: deprecated
func (r *MediaBridgeIntegratorSettingService) RegisterAppName(ctx context.Context, appID int64, body MediaBridgeIntegratorSettingRegisterAppNameParams, opts ...option.RequestOption) (res *MediaBridgeProviderRegistrationResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("media-bridge/v1/%v/settings/register", appID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Update the name that your app will display when a user is selecting media bridge
// items.
func (r *MediaBridgeIntegratorSettingService) UpdateAppName(ctx context.Context, appID int64, body MediaBridgeIntegratorSettingUpdateAppNameParams, opts ...option.RequestOption) (res *MediaBridgeProviderRegistrationResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("media-bridge/v1/%v/settings", appID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Set the visibility settings for media bridge events created by your app.
func (r *MediaBridgeIntegratorSettingService) UpdateEventVisibilitySettings(ctx context.Context, appID int64, body MediaBridgeIntegratorSettingUpdateEventVisibilitySettingsParams, opts ...option.RequestOption) (res *EventVisibilityChange, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("media-bridge/v1/%v/settings/event-visibility", appID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Update an existing oEmbed domain.
func (r *MediaBridgeIntegratorSettingService) UpdateOembedDomain(ctx context.Context, oEmbedDomainID string, params MediaBridgeIntegratorSettingUpdateOembedDomainParams, opts ...option.RequestOption) (res *IntegratorOEmbedDomainModel, err error) {
	opts = slices.Concat(r.Options, opts)
	if oEmbedDomainID == "" {
		err = errors.New("missing required oEmbedDomainId parameter")
		return
	}
	path := fmt.Sprintf("media-bridge/v1/%v/settings/oembed-domains/%s", params.AppID, oEmbedDomainID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return
}

type MediaBridgeIntegratorSettingNewObjectDefinitionParams struct {
	IntegratorObjectCreationRequest IntegratorObjectCreationRequestParam
	paramObj
}

func (r MediaBridgeIntegratorSettingNewObjectDefinitionParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.IntegratorObjectCreationRequest)
}
func (r *MediaBridgeIntegratorSettingNewObjectDefinitionParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.IntegratorObjectCreationRequest)
}

type MediaBridgeIntegratorSettingNewOembedDomainParams struct {
	IntegratorOEmbedDomainRequest IntegratorOEmbedDomainRequestParam
	paramObj
}

func (r MediaBridgeIntegratorSettingNewOembedDomainParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.IntegratorOEmbedDomainRequest)
}
func (r *MediaBridgeIntegratorSettingNewOembedDomainParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.IntegratorOEmbedDomainRequest)
}

type MediaBridgeIntegratorSettingDeleteOembedDomainParams struct {
	// The ID of the oEmbed to delete.
	ID param.Opt[int64] `query:"id,omitzero" json:"-"`
	// Filter response by Hub ID.
	DomainPortalID param.Opt[int64] `query:"domainPortalId,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [MediaBridgeIntegratorSettingDeleteOembedDomainParams]'s
// query parameters as `url.Values`.
func (r MediaBridgeIntegratorSettingDeleteOembedDomainParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeParams struct {
	AppID int64 `path:"appId,required" json:"-"`
	// Include the full definition in the response.
	IncludeFullDefinition param.Opt[bool] `query:"includeFullDefinition,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes
// [MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeParams]'s query
// parameters as `url.Values`.
func (r MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeParamsMediaType string

const (
	MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeParamsMediaTypeAudio    MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeParamsMediaType = "AUDIO"
	MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeParamsMediaTypeDocument MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeParamsMediaType = "DOCUMENT"
	MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeParamsMediaTypeImage    MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeParamsMediaType = "IMAGE"
	MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeParamsMediaTypeOther    MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeParamsMediaType = "OTHER"
	MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeParamsMediaTypeVideo    MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeParamsMediaType = "VIDEO"
)

type MediaBridgeIntegratorSettingGetOembedDomainParams struct {
	AppID int64 `path:"appId,required" json:"-"`
	paramObj
}

type MediaBridgeIntegratorSettingListOembedDomainsParams struct {
	// Filter response by Hub ID.
	DomainPortalID param.Opt[int64] `query:"domainPortalId,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [MediaBridgeIntegratorSettingListOembedDomainsParams]'s
// query parameters as `url.Values`.
func (r MediaBridgeIntegratorSettingListOembedDomainsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type MediaBridgeIntegratorSettingRegisterAppNameParams struct {
	MediaBridgeProviderPartial MediaBridgeProviderPartialParam
	paramObj
}

func (r MediaBridgeIntegratorSettingRegisterAppNameParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.MediaBridgeProviderPartial)
}
func (r *MediaBridgeIntegratorSettingRegisterAppNameParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.MediaBridgeProviderPartial)
}

type MediaBridgeIntegratorSettingUpdateAppNameParams struct {
	MediaBridgeProviderPartial MediaBridgeProviderPartialParam
	paramObj
}

func (r MediaBridgeIntegratorSettingUpdateAppNameParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.MediaBridgeProviderPartial)
}
func (r *MediaBridgeIntegratorSettingUpdateAppNameParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.MediaBridgeProviderPartial)
}

type MediaBridgeIntegratorSettingUpdateEventVisibilitySettingsParams struct {
	EventVisibilityChange EventVisibilityChangeParam
	paramObj
}

func (r MediaBridgeIntegratorSettingUpdateEventVisibilitySettingsParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.EventVisibilityChange)
}
func (r *MediaBridgeIntegratorSettingUpdateEventVisibilitySettingsParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.EventVisibilityChange)
}

type MediaBridgeIntegratorSettingUpdateOembedDomainParams struct {
	AppID                         int64 `path:"appId,required" json:"-"`
	IntegratorOEmbedDomainRequest IntegratorOEmbedDomainRequestParam
	paramObj
}

func (r MediaBridgeIntegratorSettingUpdateOembedDomainParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.IntegratorOEmbedDomainRequest)
}
func (r *MediaBridgeIntegratorSettingUpdateOembedDomainParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.IntegratorOEmbedDomainRequest)
}
