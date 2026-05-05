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
	"time"

	"github.com/HubSpot/hubspot-sdk-go/internal/apijson"
	"github.com/HubSpot/hubspot-sdk-go/internal/apiquery"
	shimjson "github.com/HubSpot/hubspot-sdk-go/internal/encoding/json"
	"github.com/HubSpot/hubspot-sdk-go/internal/requestconfig"
	"github.com/HubSpot/hubspot-sdk-go/option"
	"github.com/HubSpot/hubspot-sdk-go/packages/param"
	"github.com/HubSpot/hubspot-sdk-go/packages/respjson"
	"github.com/HubSpot/hubspot-sdk-go/shared"
)

// MediaBridgeService contains methods and other services that help with
// interacting with the hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMediaBridgeService] method instead.
type MediaBridgeService struct {
	options []option.RequestOption
	Batch   MediaBridgeBatchService
}

// NewMediaBridgeService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewMediaBridgeService(opts ...option.RequestOption) (r MediaBridgeService) {
	r = MediaBridgeService{}
	r.options = opts
	r.Batch = NewMediaBridgeBatchService(opts...)
	return
}

// Create a new association definition for the specified object type.
func (r *MediaBridgeService) NewAssociation(ctx context.Context, objectType string, params MediaBridgeNewAssociationParams, opts ...option.RequestOption) (res *shared.BaseAssociationDefinition, err error) {
	opts = slices.Concat(r.options, opts)
	if objectType == "" {
		err = errors.New("missing required objectType parameter")
		return nil, err
	}
	path := fmt.Sprintf("media-bridge/2026-03/%v/schemas/%s/associations", params.AppID, url.PathEscape(objectType))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Create an event containing the viewers attention span details for the media.
func (r *MediaBridgeService) NewAttentionSpanEvent(ctx context.Context, body MediaBridgeNewAttentionSpanEventParams, opts ...option.RequestOption) (res *AttentionSpanEvent, err error) {
	opts = slices.Concat(r.options, opts)
	path := "media-bridge/2026-03/events/attention-span"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Create an event for when a user begins playing a piece of media.
func (r *MediaBridgeService) NewMediaPlayedEvent(ctx context.Context, body MediaBridgeNewMediaPlayedEventParams, opts ...option.RequestOption) (res *MediaPlayedEvent, err error) {
	opts = slices.Concat(r.options, opts)
	path := "media-bridge/2026-03/events/media-played"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Create an event representing a user reaching quarterly milestones in a piece of
// media they're viewing.
func (r *MediaBridgeService) NewMediaPlayedPercentEvent(ctx context.Context, body MediaBridgeNewMediaPlayedPercentEventParams, opts ...option.RequestOption) (res *MediaPlayedPercentageEvent, err error) {
	opts = slices.Concat(r.options, opts)
	path := "media-bridge/2026-03/events/media-played-percent"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Create a new media object type
func (r *MediaBridgeService) NewObjectType(ctx context.Context, appID int64, body MediaBridgeNewObjectTypeParams, opts ...option.RequestOption) (res *BulkIntegratorObjectCreationResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := fmt.Sprintf("media-bridge/2026-03/%v/settings/object-definitions", appID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Set up a new oEmbed domain for your media bridge app.
func (r *MediaBridgeService) NewOembedDomain(ctx context.Context, appID int64, body MediaBridgeNewOembedDomainParams, opts ...option.RequestOption) (res *IntegratorOEmbedDomainModel, err error) {
	opts = slices.Concat(r.options, opts)
	path := fmt.Sprintf("media-bridge/2026-03/%v/settings/oembed-domains", appID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Create a new property for the specified media type
func (r *MediaBridgeService) NewProperty(ctx context.Context, objectType string, params MediaBridgeNewPropertyParams, opts ...option.RequestOption) (res *MediaBridgeProperty, err error) {
	opts = slices.Concat(r.options, opts)
	if objectType == "" {
		err = errors.New("missing required objectType parameter")
		return nil, err
	}
	path := fmt.Sprintf("media-bridge/2026-03/%v/properties/%s", params.AppID, url.PathEscape(objectType))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Create a new property group for the specified object type.
func (r *MediaBridgeService) NewPropertyGroup(ctx context.Context, objectType string, params MediaBridgeNewPropertyGroupParams, opts ...option.RequestOption) (res *shared.PropertyGroup, err error) {
	opts = slices.Concat(r.options, opts)
	if objectType == "" {
		err = errors.New("missing required objectType parameter")
		return nil, err
	}
	path := fmt.Sprintf("media-bridge/2026-03/%v/properties/%s/groups", params.AppID, url.PathEscape(objectType))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

func (r *MediaBridgeService) NewVideoAssociationDefinition(ctx context.Context, appID int64, opts ...option.RequestOption) (res *shared.BaseAssociationDefinition, err error) {
	opts = slices.Concat(r.options, opts)
	path := fmt.Sprintf("media-bridge/2026-03/%v/settings/video-association-definition", appID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

// Delete an existing association definition for an object type.
func (r *MediaBridgeService) DeleteAssociation(ctx context.Context, associationID string, body MediaBridgeDeleteAssociationParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if body.ObjectType == "" {
		err = errors.New("missing required objectType parameter")
		return err
	}
	if associationID == "" {
		err = errors.New("missing required associationId parameter")
		return err
	}
	path := fmt.Sprintf("media-bridge/2026-03/%v/schemas/%s/associations/%s", body.AppID, url.PathEscape(body.ObjectType), url.PathEscape(associationID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Delete an existing oEmbed domain.
func (r *MediaBridgeService) DeleteOembedDomain(ctx context.Context, appID int64, body MediaBridgeDeleteOembedDomainParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := fmt.Sprintf("media-bridge/2026-03/%v/settings/oembed-domains", appID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, nil, opts...)
	return err
}

// Delete an existing property for an object type.
func (r *MediaBridgeService) DeleteProperty(ctx context.Context, propertyName string, body MediaBridgeDeletePropertyParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if body.ObjectType == "" {
		err = errors.New("missing required objectType parameter")
		return err
	}
	if propertyName == "" {
		err = errors.New("missing required propertyName parameter")
		return err
	}
	path := fmt.Sprintf("media-bridge/2026-03/%v/properties/%s/%s", body.AppID, url.PathEscape(body.ObjectType), url.PathEscape(propertyName))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Delete an existing property group by name
func (r *MediaBridgeService) DeletePropertyGroup(ctx context.Context, groupName string, body MediaBridgeDeletePropertyGroupParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if body.ObjectType == "" {
		err = errors.New("missing required objectType parameter")
		return err
	}
	if groupName == "" {
		err = errors.New("missing required groupName parameter")
		return err
	}
	path := fmt.Sprintf("media-bridge/2026-03/%v/properties/%s/groups/%s", body.AppID, url.PathEscape(body.ObjectType), url.PathEscape(groupName))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Get the visibility settings for media bridge events for your apps.
func (r *MediaBridgeService) GetEventVisibilitySettings(ctx context.Context, appID int64, opts ...option.RequestOption) (res *EventVisibilityResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := fmt.Sprintf("media-bridge/2026-03/%v/settings/event-visibility", appID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Get the details for an existing oEmbed domain.
func (r *MediaBridgeService) GetOembedDomain(ctx context.Context, oEmbedDomainID string, query MediaBridgeGetOembedDomainParams, opts ...option.RequestOption) (res *IntegratorOEmbedDomainModel, err error) {
	opts = slices.Concat(r.options, opts)
	if oEmbedDomainID == "" {
		err = errors.New("missing required oEmbedDomainId parameter")
		return nil, err
	}
	path := fmt.Sprintf("media-bridge/2026-03/%v/settings/oembed-domains/%s", query.AppID, url.PathEscape(oEmbedDomainID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Get the details for an existing property by name.
func (r *MediaBridgeService) GetProperty(ctx context.Context, propertyName string, params MediaBridgeGetPropertyParams, opts ...option.RequestOption) (res *MediaBridgeProperty, err error) {
	opts = slices.Concat(r.options, opts)
	if params.ObjectType == "" {
		err = errors.New("missing required objectType parameter")
		return nil, err
	}
	if propertyName == "" {
		err = errors.New("missing required propertyName parameter")
		return nil, err
	}
	path := fmt.Sprintf("media-bridge/2026-03/%v/properties/%s/%s", params.AppID, url.PathEscape(params.ObjectType), url.PathEscape(propertyName))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// Get the details of an existing property group by name.
func (r *MediaBridgeService) GetPropertyGroup(ctx context.Context, groupName string, query MediaBridgeGetPropertyGroupParams, opts ...option.RequestOption) (res *shared.PropertyGroup, err error) {
	opts = slices.Concat(r.options, opts)
	if query.ObjectType == "" {
		err = errors.New("missing required objectType parameter")
		return nil, err
	}
	if groupName == "" {
		err = errors.New("missing required groupName parameter")
		return nil, err
	}
	path := fmt.Sprintf("media-bridge/2026-03/%v/properties/%s/groups/%s", query.AppID, url.PathEscape(query.ObjectType), url.PathEscape(groupName))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Get the schema for a specified object type.
func (r *MediaBridgeService) GetSchema(ctx context.Context, objectType string, query MediaBridgeGetSchemaParams, opts ...option.RequestOption) (res *ObjectSchema, err error) {
	opts = slices.Concat(r.options, opts)
	if objectType == "" {
		err = errors.New("missing required objectType parameter")
		return nil, err
	}
	path := fmt.Sprintf("media-bridge/2026-03/%v/schemas/%s", query.AppID, url.PathEscape(objectType))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Get the existing objects types that belong to the specified media type.
func (r *MediaBridgeService) ListObjectTypesByMediaType(ctx context.Context, mediaType MediaBridgeListObjectTypesByMediaTypeParamsMediaType, params MediaBridgeListObjectTypesByMediaTypeParams, opts ...option.RequestOption) (res *ObjectDefinitionResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := fmt.Sprintf("media-bridge/2026-03/%v/settings/object-definitions/%v", params.AppID, mediaType)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// Get the details for existing oEmbed domains for your app
func (r *MediaBridgeService) ListOembedDomains(ctx context.Context, appID int64, query MediaBridgeListOembedDomainsParams, opts ...option.RequestOption) (res *OEmbedDomainsCollectionResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := fmt.Sprintf("media-bridge/2026-03/%v/settings/oembed-domains", appID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Get the existing properties defined for a media object type.
func (r *MediaBridgeService) ListProperties(ctx context.Context, objectType string, params MediaBridgeListPropertiesParams, opts ...option.RequestOption) (res *CollectionResponsePropertyNoPaging, err error) {
	opts = slices.Concat(r.options, opts)
	if objectType == "" {
		err = errors.New("missing required objectType parameter")
		return nil, err
	}
	path := fmt.Sprintf("media-bridge/2026-03/%v/properties/%s", params.AppID, url.PathEscape(objectType))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// Get the property groups for a specified object type.
func (r *MediaBridgeService) ListPropertyGroups(ctx context.Context, objectType string, query MediaBridgeListPropertyGroupsParams, opts ...option.RequestOption) (res *shared.CollectionResponsePropertyGroupNoPaging, err error) {
	opts = slices.Concat(r.options, opts)
	if objectType == "" {
		err = errors.New("missing required objectType parameter")
		return nil, err
	}
	path := fmt.Sprintf("media-bridge/2026-03/%v/properties/%s/groups", query.AppID, url.PathEscape(objectType))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Get the schemas for all object types.
func (r *MediaBridgeService) ListSchemas(ctx context.Context, appID int64, query MediaBridgeListSchemasParams, opts ...option.RequestOption) (res *CollectionResponseObjectSchemaNoPaging, err error) {
	opts = slices.Concat(r.options, opts)
	path := fmt.Sprintf("media-bridge/2026-03/%v/schemas", appID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Register the name that your app will display when a user is selecting media
// bridge items.
//
// Deprecated: deprecated
func (r *MediaBridgeService) RegisterAppName(ctx context.Context, appID int64, body MediaBridgeRegisterAppNameParams, opts ...option.RequestOption) (res *MediaBridgeProviderRegistrationResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := fmt.Sprintf("media-bridge/2026-03/%v/settings/register", appID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Set the visibility settings for media bridge events created by your app.
func (r *MediaBridgeService) UpdateEventVisibilitySettings(ctx context.Context, appID int64, body MediaBridgeUpdateEventVisibilitySettingsParams, opts ...option.RequestOption) (res *EventVisibilityChange, err error) {
	opts = slices.Concat(r.options, opts)
	path := fmt.Sprintf("media-bridge/2026-03/%v/settings/event-visibility", appID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

// Update an existing oEmbed domain.
func (r *MediaBridgeService) UpdateOembedDomain(ctx context.Context, oEmbedDomainID string, params MediaBridgeUpdateOembedDomainParams, opts ...option.RequestOption) (res *IntegratorOEmbedDomainModel, err error) {
	opts = slices.Concat(r.options, opts)
	if oEmbedDomainID == "" {
		err = errors.New("missing required oEmbedDomainId parameter")
		return nil, err
	}
	path := fmt.Sprintf("media-bridge/2026-03/%v/settings/oembed-domains/%s", params.AppID, url.PathEscape(oEmbedDomainID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return res, err
}

// Update an existing property for an object type.
func (r *MediaBridgeService) UpdateProperty(ctx context.Context, propertyName string, params MediaBridgeUpdatePropertyParams, opts ...option.RequestOption) (res *MediaBridgeProperty, err error) {
	opts = slices.Concat(r.options, opts)
	if params.ObjectType == "" {
		err = errors.New("missing required objectType parameter")
		return nil, err
	}
	if propertyName == "" {
		err = errors.New("missing required propertyName parameter")
		return nil, err
	}
	path := fmt.Sprintf("media-bridge/2026-03/%v/properties/%s/%s", params.AppID, url.PathEscape(params.ObjectType), url.PathEscape(propertyName))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return res, err
}

// Update an existing property group by name.
func (r *MediaBridgeService) UpdatePropertyGroup(ctx context.Context, groupName string, params MediaBridgeUpdatePropertyGroupParams, opts ...option.RequestOption) (res *shared.PropertyGroup, err error) {
	opts = slices.Concat(r.options, opts)
	if params.ObjectType == "" {
		err = errors.New("missing required objectType parameter")
		return nil, err
	}
	if groupName == "" {
		err = errors.New("missing required groupName parameter")
		return nil, err
	}
	path := fmt.Sprintf("media-bridge/2026-03/%v/properties/%s/groups/%s", params.AppID, url.PathEscape(params.ObjectType), url.PathEscape(groupName))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return res, err
}

// Update the schema for an existing object type
func (r *MediaBridgeService) UpdateSchema(ctx context.Context, objectType string, params MediaBridgeUpdateSchemaParams, opts ...option.RequestOption) (res *shared.BaseObjectTypeDefinition, err error) {
	opts = slices.Concat(r.options, opts)
	if objectType == "" {
		err = errors.New("missing required objectType parameter")
		return nil, err
	}
	path := fmt.Sprintf("media-bridge/2026-03/%v/schemas/%s", params.AppID, url.PathEscape(objectType))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return res, err
}

// Update the name that your app will display when a user is selecting media bridge
// items.
func (r *MediaBridgeService) UpdateSettings(ctx context.Context, appID int64, body MediaBridgeUpdateSettingsParams, opts ...option.RequestOption) (res *MediaBridgeProviderRegistrationResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := fmt.Sprintf("media-bridge/2026-03/%v/settings", appID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return res, err
}

// The properties TotalPercentPlayed, TotalSecondsPlayed are required.
type AttentionSpanCalculatedValuesParam struct {
	TotalPercentPlayed float64 `json:"totalPercentPlayed" api:"required"`
	TotalSecondsPlayed int64   `json:"totalSecondsPlayed" api:"required"`
	paramObj
}

func (r AttentionSpanCalculatedValuesParam) MarshalJSON() (data []byte, err error) {
	type shadow AttentionSpanCalculatedValuesParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AttentionSpanCalculatedValuesParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AttentionSpanEvent struct {
	// The ID of the contact in HubSpot’s system that consumed the media. This can be
	// fetched using HubSpot's Get contact by usertoken (utk) API. The API also
	// supports supplying a usertoken, and will handle converting this into a contact
	// ID automatically.
	ContactID                    int64  `json:"contactId" api:"required"`
	MediaBridgeID                int64  `json:"mediaBridgeId" api:"required"`
	MediaBridgeObjectCoordinates string `json:"mediaBridgeObjectCoordinates" api:"required"`
	MediaBridgeObjectTypeID      string `json:"mediaBridgeObjectTypeId" api:"required"`
	MediaName                    string `json:"mediaName" api:"required"`
	// Any of "AUDIO", "DOCUMENT", "IMAGE", "OTHER", "VIDEO".
	MediaType AttentionSpanEventMediaType `json:"mediaType" api:"required"`
	// The timestamp at which this event occurred, in milliseconds since the epoch.
	OccurredTimestamp int64  `json:"occurredTimestamp" api:"required"`
	PercentRange      string `json:"percentRange" api:"required"`
	// The ID of the HubSpot account.
	PortalID   int64  `json:"portalId" api:"required"`
	ProviderID int64  `json:"providerId" api:"required"`
	SessionID  string `json:"sessionId" api:"required"`
	// The percent of the media that the user consumed. Providers may calculate this
	// differently depending on how they consider repeated views of the same portion of
	// media. For this reason, the API will not attempt to validate totalPercentWatched
	// against the attention span information for the event. If it is missing, HubSpot
	// will calculate this from the attention span map as follows: (number of spans
	// with a value of 1 or more)/(Total number of spans).
	TotalPercentPlayed float64 `json:"totalPercentPlayed" api:"required"`
	// Any of "EMAIL", "EXTERNAL_PAGE".
	ExternalPlayContext AttentionSpanEventExternalPlayContext `json:"externalPlayContext"`
	MediaURL            string                                `json:"mediaUrl"`
	// The ID of the page, if hosted on HubSpot. Required for HubSpot pages.
	PageID int64 `json:"pageId"`
	// The name of the page. Required if the page is not hosted on HubSpot.
	PageName              string `json:"pageName"`
	PageObjectCoordinates string `json:"pageObjectCoordinates"`
	// The URL of the page that an event happened on. Required if the page is not
	// hosted on HubSpot.
	PageURL string `json:"pageUrl"`
	// This is the raw data which provides the most granular data about spans of the
	// media, and how many times each span was consumed by the user. For example, for a
	// 10 second video where each second is a span, if a visitor watches the first 5
	// seconds of the video, then restarts the video and watches the first 2 seconds
	// again, the resulting `rawDataString` would be
	// `“0=2;1=2;2=1;3=1;4=1;5=0;6=0;7=0;8=0;9=0;”`.
	RawData string `json:"rawData"`
	// The seconds that a user spent consuming the media. The media bridge calculates
	// this as `totalPercentPlayed`\*`mediaDuration`. If a provider would like this to
	// be calculated differently, they can provide the pre-calculated value when they
	// create the event.
	TotalSecondsPlayed int64 `json:"totalSecondsPlayed"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ContactID                    respjson.Field
		MediaBridgeID                respjson.Field
		MediaBridgeObjectCoordinates respjson.Field
		MediaBridgeObjectTypeID      respjson.Field
		MediaName                    respjson.Field
		MediaType                    respjson.Field
		OccurredTimestamp            respjson.Field
		PercentRange                 respjson.Field
		PortalID                     respjson.Field
		ProviderID                   respjson.Field
		SessionID                    respjson.Field
		TotalPercentPlayed           respjson.Field
		ExternalPlayContext          respjson.Field
		MediaURL                     respjson.Field
		PageID                       respjson.Field
		PageName                     respjson.Field
		PageObjectCoordinates        respjson.Field
		PageURL                      respjson.Field
		RawData                      respjson.Field
		TotalSecondsPlayed           respjson.Field
		ExtraFields                  map[string]respjson.Field
		raw                          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AttentionSpanEvent) RawJSON() string { return r.JSON.raw }
func (r *AttentionSpanEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AttentionSpanEventMediaType string

const (
	AttentionSpanEventMediaTypeAudio    AttentionSpanEventMediaType = "AUDIO"
	AttentionSpanEventMediaTypeDocument AttentionSpanEventMediaType = "DOCUMENT"
	AttentionSpanEventMediaTypeImage    AttentionSpanEventMediaType = "IMAGE"
	AttentionSpanEventMediaTypeOther    AttentionSpanEventMediaType = "OTHER"
	AttentionSpanEventMediaTypeVideo    AttentionSpanEventMediaType = "VIDEO"
)

type AttentionSpanEventExternalPlayContext string

const (
	AttentionSpanEventExternalPlayContextEmail        AttentionSpanEventExternalPlayContext = "EMAIL"
	AttentionSpanEventExternalPlayContextExternalPage AttentionSpanEventExternalPlayContext = "EXTERNAL_PAGE"
)

// The properties MediaType, OccurredTimestamp, RawDataMap, SessionID are required.
type AttentionSpanEventRequestParam struct {
	// Any of "AUDIO", "DOCUMENT", "IMAGE", "OTHER", "VIDEO".
	MediaType         AttentionSpanEventRequestMediaType `json:"mediaType,omitzero" api:"required"`
	OccurredTimestamp int64                              `json:"occurredTimestamp" api:"required"`
	RawDataMap        map[string]int64                   `json:"rawDataMap,omitzero" api:"required"`
	SessionID         string                             `json:"sessionId" api:"required"`
	Hsenc             param.Opt[string]                  `json:"_hsenc,omitzero"`
	ContactID         param.Opt[int64]                   `json:"contactId,omitzero"`
	ContactUtk        param.Opt[string]                  `json:"contactUtk,omitzero"`
	ExternalID        param.Opt[string]                  `json:"externalId,omitzero"`
	MediaBridgeID     param.Opt[int64]                   `json:"mediaBridgeId,omitzero"`
	MediaName         param.Opt[string]                  `json:"mediaName,omitzero"`
	MediaURL          param.Opt[string]                  `json:"mediaUrl,omitzero"`
	PageID            param.Opt[int64]                   `json:"pageId,omitzero"`
	PageName          param.Opt[string]                  `json:"pageName,omitzero"`
	PageURL           param.Opt[string]                  `json:"pageUrl,omitzero"`
	RawDataString     param.Opt[string]                  `json:"rawDataString,omitzero"`
	DerivedValues     AttentionSpanCalculatedValuesParam `json:"derivedValues,omitzero"`
	// Any of "EMAIL", "EXTERNAL_PAGE".
	ExternalPlayContext AttentionSpanEventRequestExternalPlayContext `json:"externalPlayContext,omitzero"`
	paramObj
}

func (r AttentionSpanEventRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow AttentionSpanEventRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AttentionSpanEventRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AttentionSpanEventRequestMediaType string

const (
	AttentionSpanEventRequestMediaTypeAudio    AttentionSpanEventRequestMediaType = "AUDIO"
	AttentionSpanEventRequestMediaTypeDocument AttentionSpanEventRequestMediaType = "DOCUMENT"
	AttentionSpanEventRequestMediaTypeImage    AttentionSpanEventRequestMediaType = "IMAGE"
	AttentionSpanEventRequestMediaTypeOther    AttentionSpanEventRequestMediaType = "OTHER"
	AttentionSpanEventRequestMediaTypeVideo    AttentionSpanEventRequestMediaType = "VIDEO"
)

type AttentionSpanEventRequestExternalPlayContext string

const (
	AttentionSpanEventRequestExternalPlayContextEmail        AttentionSpanEventRequestExternalPlayContext = "EMAIL"
	AttentionSpanEventRequestExternalPlayContextExternalPage AttentionSpanEventRequestExternalPlayContext = "EXTERNAL_PAGE"
)

type BatchResponseProperty struct {
	CompletedAt time.Time             `json:"completedAt" api:"required" format:"date-time"`
	Results     []MediaBridgeProperty `json:"results" api:"required"`
	StartedAt   time.Time             `json:"startedAt" api:"required" format:"date-time"`
	// Any of "CANCELED", "COMPLETE", "PENDING", "PROCESSING".
	Status      BatchResponsePropertyStatus `json:"status" api:"required"`
	Links       map[string]string           `json:"links"`
	RequestedAt time.Time                   `json:"requestedAt" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CompletedAt respjson.Field
		Results     respjson.Field
		StartedAt   respjson.Field
		Status      respjson.Field
		Links       respjson.Field
		RequestedAt respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BatchResponseProperty) RawJSON() string { return r.JSON.raw }
func (r *BatchResponseProperty) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BatchResponsePropertyStatus string

const (
	BatchResponsePropertyStatusCanceled   BatchResponsePropertyStatus = "CANCELED"
	BatchResponsePropertyStatusComplete   BatchResponsePropertyStatus = "COMPLETE"
	BatchResponsePropertyStatusPending    BatchResponsePropertyStatus = "PENDING"
	BatchResponsePropertyStatusProcessing BatchResponsePropertyStatus = "PROCESSING"
)

type BulkIntegratorObjectCreationResponse struct {
	CreatedObjects map[string]IntegratorObjectCreationResponse `json:"createdObjects" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreatedObjects respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BulkIntegratorObjectCreationResponse) RawJSON() string { return r.JSON.raw }
func (r *BulkIntegratorObjectCreationResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CaseChangeTestExtensionData struct {
	// Any of "ANGRY", "HAPPY", "SAD", "SARCASTIC".
	Mood CaseChangeTestExtensionDataMood `json:"mood" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Mood        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CaseChangeTestExtensionData) RawJSON() string { return r.JSON.raw }
func (r *CaseChangeTestExtensionData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CaseChangeTestExtensionDataMood string

const (
	CaseChangeTestExtensionDataMoodAngry     CaseChangeTestExtensionDataMood = "ANGRY"
	CaseChangeTestExtensionDataMoodHappy     CaseChangeTestExtensionDataMood = "HAPPY"
	CaseChangeTestExtensionDataMoodSad       CaseChangeTestExtensionDataMood = "SAD"
	CaseChangeTestExtensionDataMoodSarcastic CaseChangeTestExtensionDataMood = "SARCASTIC"
)

type CollectionResponseObjectSchemaNoPaging struct {
	Results []ObjectSchema `json:"results" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Results     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CollectionResponseObjectSchemaNoPaging) RawJSON() string { return r.JSON.raw }
func (r *CollectionResponseObjectSchemaNoPaging) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CollectionResponsePropertyNoPaging struct {
	Results []Property1 `json:"results" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Results     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CollectionResponsePropertyNoPaging) RawJSON() string { return r.JSON.raw }
func (r *CollectionResponsePropertyNoPaging) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DefaultRequirements struct {
	Gates []string `json:"gates" api:"required"`
	// Any of "AND", "OR".
	Operator   DefaultRequirementsOperator `json:"operator" api:"required"`
	ScopeNames []string                    `json:"scopeNames" api:"required"`
	Settings   []string                    `json:"settings" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Gates       respjson.Field
		Operator    respjson.Field
		ScopeNames  respjson.Field
		Settings    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DefaultRequirements) RawJSON() string { return r.JSON.raw }
func (r *DefaultRequirements) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DefaultRequirementsOperator string

const (
	DefaultRequirementsOperatorAnd DefaultRequirementsOperator = "AND"
	DefaultRequirementsOperatorOr  DefaultRequirementsOperator = "OR"
)

type DefinitionSource struct {
	// Any of "GLOBAL", "HAVEN_BRANCH", "OBJECT_TYPE", "PORTAL".
	Type DefinitionSourceType `json:"type" api:"required"`
	Name string               `json:"name"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DefinitionSource) RawJSON() string { return r.JSON.raw }
func (r *DefinitionSource) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DefinitionSourceType string

const (
	DefinitionSourceTypeGlobal      DefinitionSourceType = "GLOBAL"
	DefinitionSourceTypeHavenBranch DefinitionSourceType = "HAVEN_BRANCH"
	DefinitionSourceTypeObjectType  DefinitionSourceType = "OBJECT_TYPE"
	DefinitionSourceTypePortal      DefinitionSourceType = "PORTAL"
)

type Endpoints struct {
	Discovery bool     `json:"discovery" api:"required"`
	Schemes   []string `json:"schemes" api:"required"`
	URL       string   `json:"url" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Discovery   respjson.Field
		Schemes     respjson.Field
		URL         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Endpoints) RawJSON() string { return r.JSON.raw }
func (r *Endpoints) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this Endpoints to a EndpointsParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// EndpointsParam.Overrides()
func (r Endpoints) ToParam() EndpointsParam {
	return param.Override[EndpointsParam](json.RawMessage(r.RawJSON()))
}

// The properties Discovery, Schemes, URL are required.
type EndpointsParam struct {
	Discovery bool     `json:"discovery" api:"required"`
	Schemes   []string `json:"schemes,omitzero" api:"required"`
	URL       string   `json:"url" api:"required"`
	paramObj
}

func (r EndpointsParam) MarshalJSON() (data []byte, err error) {
	type shadow EndpointsParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EndpointsParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EventVisibilityChange struct {
	// Any of "ALL", "ATTENTION_SPAN", "MEDIA_PLAYS", "MEDIA_PLAYS_PERCENT".
	EventType       EventVisibilityChangeEventType `json:"eventType" api:"required"`
	UpdatedAt       int64                          `json:"updatedAt" api:"required"`
	ShowInReporting bool                           `json:"showInReporting"`
	ShowInTimeline  bool                           `json:"showInTimeline"`
	ShowInWorkflows bool                           `json:"showInWorkflows"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		EventType       respjson.Field
		UpdatedAt       respjson.Field
		ShowInReporting respjson.Field
		ShowInTimeline  respjson.Field
		ShowInWorkflows respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EventVisibilityChange) RawJSON() string { return r.JSON.raw }
func (r *EventVisibilityChange) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this EventVisibilityChange to a EventVisibilityChangeParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// EventVisibilityChangeParam.Overrides()
func (r EventVisibilityChange) ToParam() EventVisibilityChangeParam {
	return param.Override[EventVisibilityChangeParam](json.RawMessage(r.RawJSON()))
}

type EventVisibilityChangeEventType string

const (
	EventVisibilityChangeEventTypeAll               EventVisibilityChangeEventType = "ALL"
	EventVisibilityChangeEventTypeAttentionSpan     EventVisibilityChangeEventType = "ATTENTION_SPAN"
	EventVisibilityChangeEventTypeMediaPlays        EventVisibilityChangeEventType = "MEDIA_PLAYS"
	EventVisibilityChangeEventTypeMediaPlaysPercent EventVisibilityChangeEventType = "MEDIA_PLAYS_PERCENT"
)

// The properties EventType, UpdatedAt are required.
type EventVisibilityChangeParam struct {
	// Any of "ALL", "ATTENTION_SPAN", "MEDIA_PLAYS", "MEDIA_PLAYS_PERCENT".
	EventType       EventVisibilityChangeEventType `json:"eventType,omitzero" api:"required"`
	UpdatedAt       int64                          `json:"updatedAt" api:"required"`
	ShowInReporting param.Opt[bool]                `json:"showInReporting,omitzero"`
	ShowInTimeline  param.Opt[bool]                `json:"showInTimeline,omitzero"`
	ShowInWorkflows param.Opt[bool]                `json:"showInWorkflows,omitzero"`
	paramObj
}

func (r EventVisibilityChangeParam) MarshalJSON() (data []byte, err error) {
	type shadow EventVisibilityChangeParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EventVisibilityChangeParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EventVisibilityResponse struct {
	CreatedAt          time.Time               `json:"createdAt" api:"required" format:"date-time"`
	VisibilitySettings []EventVisibilityChange `json:"visibilitySettings" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreatedAt          respjson.Field
		VisibilitySettings respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EventVisibilityResponse) RawJSON() string { return r.JSON.raw }
func (r *EventVisibilityResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExtensionData struct {
	// Any of "OK", "ERROR", "TIMEOUT".
	ExtensionStatusMap                  map[string]string                   `json:"extensionStatusMap" api:"required"`
	Tags                                []string                            `json:"tags" api:"required"`
	CaseChangeTestExtensionData         CaseChangeTestExtensionData         `json:"caseChangeTestExtensionData"`
	OptionDecoratorsExtensionData       OptionDecoratorsExtensionData       `json:"optionDecoratorsExtensionData"`
	RequiredPropertiesExtensionData     RequiredPropertiesExtensionData     `json:"requiredPropertiesExtensionData"`
	SoftRequiredPropertiesExtensionData SoftRequiredPropertiesExtensionData `json:"softRequiredPropertiesExtensionData"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ExtensionStatusMap                  respjson.Field
		Tags                                respjson.Field
		CaseChangeTestExtensionData         respjson.Field
		OptionDecoratorsExtensionData       respjson.Field
		RequiredPropertiesExtensionData     respjson.Field
		SoftRequiredPropertiesExtensionData respjson.Field
		ExtraFields                         map[string]respjson.Field
		raw                                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExtensionData) RawJSON() string { return r.JSON.raw }
func (r *ExtensionData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExternalOptionsMetaData struct {
	Filter              FilteringMetaData `json:"filter"`
	RelatedObjectTypeID string            `json:"relatedObjectTypeId"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Filter              respjson.Field
		RelatedObjectTypeID respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExternalOptionsMetaData) RawJSON() string { return r.JSON.raw }
func (r *ExternalOptionsMetaData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FieldLevelPermission struct {
	AccessLevel string `json:"accessLevel" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AccessLevel respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FieldLevelPermission) RawJSON() string { return r.JSON.raw }
func (r *FieldLevelPermission) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FilteringMetaData struct {
	IncludeHelpdeskRoutableTeamsOnly bool     `json:"includeHelpdeskRoutableTeamsOnly" api:"required"`
	IncludeUnconfirmedUsers          bool     `json:"includeUnconfirmedUsers" api:"required"`
	ListProcessingTypes              []string `json:"listProcessingTypes" api:"required"`
	PipelineIDs                      []string `json:"pipelineIds" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		IncludeHelpdeskRoutableTeamsOnly respjson.Field
		IncludeUnconfirmedUsers          respjson.Field
		ListProcessingTypes              respjson.Field
		PipelineIDs                      respjson.Field
		ExtraFields                      map[string]respjson.Field
		raw                              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FilteringMetaData) RawJSON() string { return r.JSON.raw }
func (r *FilteringMetaData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Group struct {
	Deleted          bool   `json:"deleted" api:"required"`
	DisplayName      string `json:"displayName" api:"required"`
	DisplayOrder     int64  `json:"displayOrder" api:"required"`
	FulcrumPortalID  int64  `json:"fulcrumPortalId" api:"required"`
	FulcrumTimestamp int64  `json:"fulcrumTimestamp" api:"required"`
	HubSpotDefined   bool   `json:"hubspotDefined" api:"required"`
	Name             string `json:"name" api:"required"`
	PortalID         int64  `json:"portalId" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Deleted          respjson.Field
		DisplayName      respjson.Field
		DisplayOrder     respjson.Field
		FulcrumPortalID  respjson.Field
		FulcrumTimestamp respjson.Field
		HubSpotDefined   respjson.Field
		Name             respjson.Field
		PortalID         respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Group) RawJSON() string { return r.JSON.raw }
func (r *Group) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GroupView struct {
	DisplayName      string `json:"displayName" api:"required"`
	DisplayOrder     int64  `json:"displayOrder" api:"required"`
	FulcrumPortalID  int64  `json:"fulcrumPortalId" api:"required"`
	FulcrumTimestamp int64  `json:"fulcrumTimestamp" api:"required"`
	HubSpotDefined   bool   `json:"hubspotDefined" api:"required"`
	Name             string `json:"name" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DisplayName      respjson.Field
		DisplayOrder     respjson.Field
		FulcrumPortalID  respjson.Field
		FulcrumTimestamp respjson.Field
		HubSpotDefined   respjson.Field
		Name             respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GroupView) RawJSON() string { return r.JSON.raw }
func (r *GroupView) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InboundDBObjectType struct {
	ID                          int64    `json:"id" api:"required"`
	AllowsSensitiveProperties   bool     `json:"allowsSensitiveProperties" api:"required"`
	CreateDatePropertyName      string   `json:"createDatePropertyName" api:"required"`
	DefaultSearchPropertyNames  []string `json:"defaultSearchPropertyNames" api:"required"`
	Deleted                     bool     `json:"deleted" api:"required"`
	FullyQualifiedName          string   `json:"fullyQualifiedName" api:"required"`
	HasCustomProperties         bool     `json:"hasCustomProperties" api:"required"`
	HasDefaultProperties        bool     `json:"hasDefaultProperties" api:"required"`
	HasExternalObjectIDs        bool     `json:"hasExternalObjectIds" api:"required"`
	HasOwners                   bool     `json:"hasOwners" api:"required"`
	HasPipelines                bool     `json:"hasPipelines" api:"required"`
	IndexedForFiltersAndReports bool     `json:"indexedForFiltersAndReports" api:"required"`
	LastModifiedPropertyName    string   `json:"lastModifiedPropertyName" api:"required"`
	// Any of "CMS_HUBDB", "HUBSPOT", "HUBSPOT_EVENT", "INTEGRATION",
	// "INTEGRATION_EVENT", "PORTAL_SPECIFIC", "PORTAL_SPECIFIC_EVENT", "WORK",
	// "WORK_SUB".
	MetaType           InboundDBObjectTypeMetaType `json:"metaType" api:"required"`
	MetaTypeID         int64                       `json:"metaTypeId" api:"required"`
	Name               string                      `json:"name" api:"required"`
	ObjectTypeID       string                      `json:"objectTypeId" api:"required"`
	ObjectTypeIDString string                      `json:"objectTypeIdString" api:"required"`
	// Any of "ALL_OR_NONE", "DO_NOT_CHECK_PERMISSIONS", "EXPLICIT", "OWNER_BASED",
	// "TEAM_BASED".
	PermissioningType                  InboundDBObjectTypePermissioningType `json:"permissioningType" api:"required"`
	PipelinePropertyName               string                               `json:"pipelinePropertyName" api:"required"`
	PipelineStagePropertyName          string                               `json:"pipelineStagePropertyName" api:"required"`
	RequiredProperties                 []string                             `json:"requiredProperties" api:"required"`
	Restorable                         bool                                 `json:"restorable" api:"required"`
	ScopeMappings                      []ScopeMapping                       `json:"scopeMappings" api:"required"`
	SecondaryDisplayLabelPropertyNames []string                             `json:"secondaryDisplayLabelPropertyNames" api:"required"`
	AccessScopeName                    string                               `json:"accessScopeName"`
	CreatedAt                          int64                                `json:"createdAt"`
	Description                        string                               `json:"description"`
	IntegrationAppID                   int64                                `json:"integrationAppId"`
	JanusGroup                         string                               `json:"janusGroup"`
	OwnerPortalID                      int64                                `json:"ownerPortalId"`
	PipelineCloseDatePropertyName      string                               `json:"pipelineCloseDatePropertyName"`
	PipelineTimeToClosePropertyName    string                               `json:"pipelineTimeToClosePropertyName"`
	PluralForm                         string                               `json:"pluralForm"`
	PrimaryDisplayLabelPropertyName    string                               `json:"primaryDisplayLabelPropertyName"`
	ReadScopeName                      string                               `json:"readScopeName"`
	SingularForm                       string                               `json:"singularForm"`
	// Any of "Deprecated", "In development", "Live".
	Status InboundDBObjectTypeStatus `json:"status"`
	// Any of "Customer-facing", "Customer-facing public API", "Customer-facing UI",
	// "Internal only".
	Visibility     InboundDBObjectTypeVisibility `json:"visibility"`
	WriteScopeName string                        `json:"writeScopeName"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                                 respjson.Field
		AllowsSensitiveProperties          respjson.Field
		CreateDatePropertyName             respjson.Field
		DefaultSearchPropertyNames         respjson.Field
		Deleted                            respjson.Field
		FullyQualifiedName                 respjson.Field
		HasCustomProperties                respjson.Field
		HasDefaultProperties               respjson.Field
		HasExternalObjectIDs               respjson.Field
		HasOwners                          respjson.Field
		HasPipelines                       respjson.Field
		IndexedForFiltersAndReports        respjson.Field
		LastModifiedPropertyName           respjson.Field
		MetaType                           respjson.Field
		MetaTypeID                         respjson.Field
		Name                               respjson.Field
		ObjectTypeID                       respjson.Field
		ObjectTypeIDString                 respjson.Field
		PermissioningType                  respjson.Field
		PipelinePropertyName               respjson.Field
		PipelineStagePropertyName          respjson.Field
		RequiredProperties                 respjson.Field
		Restorable                         respjson.Field
		ScopeMappings                      respjson.Field
		SecondaryDisplayLabelPropertyNames respjson.Field
		AccessScopeName                    respjson.Field
		CreatedAt                          respjson.Field
		Description                        respjson.Field
		IntegrationAppID                   respjson.Field
		JanusGroup                         respjson.Field
		OwnerPortalID                      respjson.Field
		PipelineCloseDatePropertyName      respjson.Field
		PipelineTimeToClosePropertyName    respjson.Field
		PluralForm                         respjson.Field
		PrimaryDisplayLabelPropertyName    respjson.Field
		ReadScopeName                      respjson.Field
		SingularForm                       respjson.Field
		Status                             respjson.Field
		Visibility                         respjson.Field
		WriteScopeName                     respjson.Field
		ExtraFields                        map[string]respjson.Field
		raw                                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InboundDBObjectType) RawJSON() string { return r.JSON.raw }
func (r *InboundDBObjectType) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InboundDBObjectTypeMetaType string

const (
	InboundDBObjectTypeMetaTypeCmsHubdb            InboundDBObjectTypeMetaType = "CMS_HUBDB"
	InboundDBObjectTypeMetaTypeHubSpot             InboundDBObjectTypeMetaType = "HUBSPOT"
	InboundDBObjectTypeMetaTypeHubSpotEvent        InboundDBObjectTypeMetaType = "HUBSPOT_EVENT"
	InboundDBObjectTypeMetaTypeIntegration         InboundDBObjectTypeMetaType = "INTEGRATION"
	InboundDBObjectTypeMetaTypeIntegrationEvent    InboundDBObjectTypeMetaType = "INTEGRATION_EVENT"
	InboundDBObjectTypeMetaTypePortalSpecific      InboundDBObjectTypeMetaType = "PORTAL_SPECIFIC"
	InboundDBObjectTypeMetaTypePortalSpecificEvent InboundDBObjectTypeMetaType = "PORTAL_SPECIFIC_EVENT"
	InboundDBObjectTypeMetaTypeWork                InboundDBObjectTypeMetaType = "WORK"
	InboundDBObjectTypeMetaTypeWorkSub             InboundDBObjectTypeMetaType = "WORK_SUB"
)

type InboundDBObjectTypePermissioningType string

const (
	InboundDBObjectTypePermissioningTypeAllOrNone             InboundDBObjectTypePermissioningType = "ALL_OR_NONE"
	InboundDBObjectTypePermissioningTypeDoNotCheckPermissions InboundDBObjectTypePermissioningType = "DO_NOT_CHECK_PERMISSIONS"
	InboundDBObjectTypePermissioningTypeExplicit              InboundDBObjectTypePermissioningType = "EXPLICIT"
	InboundDBObjectTypePermissioningTypeOwnerBased            InboundDBObjectTypePermissioningType = "OWNER_BASED"
	InboundDBObjectTypePermissioningTypeTeamBased             InboundDBObjectTypePermissioningType = "TEAM_BASED"
)

type InboundDBObjectTypeStatus string

const (
	InboundDBObjectTypeStatusDeprecated    InboundDBObjectTypeStatus = "Deprecated"
	InboundDBObjectTypeStatusInDevelopment InboundDBObjectTypeStatus = "In development"
	InboundDBObjectTypeStatusLive          InboundDBObjectTypeStatus = "Live"
)

type InboundDBObjectTypeVisibility string

const (
	InboundDBObjectTypeVisibilityCustomerFacing          InboundDBObjectTypeVisibility = "Customer-facing"
	InboundDBObjectTypeVisibilityCustomerFacingPublicAPI InboundDBObjectTypeVisibility = "Customer-facing public API"
	InboundDBObjectTypeVisibilityCustomerFacingUi        InboundDBObjectTypeVisibility = "Customer-facing UI"
	InboundDBObjectTypeVisibilityInternalOnly            InboundDBObjectTypeVisibility = "Internal only"
)

type IntegratorOEmbedDomainModel struct {
	ID        int64     `json:"id" api:"required"`
	AppID     int64     `json:"appId" api:"required"`
	CreatedAt int64     `json:"createdAt" api:"required"`
	DeletedAt int64     `json:"deletedAt" api:"required"`
	Endpoints Endpoints `json:"endpoints" api:"required"`
	PortalID  int64     `json:"portalId" api:"required"`
	UpdatedAt int64     `json:"updatedAt" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		AppID       respjson.Field
		CreatedAt   respjson.Field
		DeletedAt   respjson.Field
		Endpoints   respjson.Field
		PortalID    respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IntegratorOEmbedDomainModel) RawJSON() string { return r.JSON.raw }
func (r *IntegratorOEmbedDomainModel) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Endpoints is required.
type IntegratorOEmbedDomainRequestParam struct {
	Endpoints EndpointsParam   `json:"endpoints,omitzero" api:"required"`
	PortalID  param.Opt[int64] `json:"portalId,omitzero"`
	paramObj
}

func (r IntegratorOEmbedDomainRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow IntegratorOEmbedDomainRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *IntegratorOEmbedDomainRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property MediaTypes is required.
type IntegratorObjectCreationRequestParam struct {
	// Any of "VIDEO", "AUDIO", "DOCUMENT", "OTHER", "IMAGE".
	MediaTypes []string `json:"mediaTypes,omitzero" api:"required"`
	paramObj
}

func (r IntegratorObjectCreationRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow IntegratorObjectCreationRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *IntegratorObjectCreationRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IntegratorObjectCreationResponse struct {
	ObjectType     InboundDBObjectType  `json:"objectType" api:"required"`
	Properties     []PropertyDefinition `json:"properties" api:"required"`
	PropertyGroups []Group              `json:"propertyGroups" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ObjectType     respjson.Field
		Properties     respjson.Field
		PropertyGroups respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IntegratorObjectCreationResponse) RawJSON() string { return r.JSON.raw }
func (r *IntegratorObjectCreationResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LookupAssociationSpec struct {
	// Defines the type, direction, and details of the relationship between two CRM
	// objects.
	AssociationSpec shared.AssociationSpec `json:"associationSpec" api:"required"`
	// Any of "ONE_TO_MANY", "ONE_TO_ONE".
	Cardinality    LookupAssociationSpecCardinality `json:"cardinality"`
	MaxToObjectIDs int64                            `json:"maxToObjectIds"`
	ToObjectTypeID string                           `json:"toObjectTypeId"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AssociationSpec respjson.Field
		Cardinality     respjson.Field
		MaxToObjectIDs  respjson.Field
		ToObjectTypeID  respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LookupAssociationSpec) RawJSON() string { return r.JSON.raw }
func (r *LookupAssociationSpec) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LookupAssociationSpecCardinality string

const (
	LookupAssociationSpecCardinalityOneToMany LookupAssociationSpecCardinality = "ONE_TO_MANY"
	LookupAssociationSpecCardinalityOneToOne  LookupAssociationSpecCardinality = "ONE_TO_ONE"
)

// A HubSpot property
type MediaBridgeProperty struct {
	// Object types permitted to use this property.
	AllowedObjectTypes []ObjectTypeIDProto `json:"allowedObjectTypes" api:"required"`
	// Whether the property is a calculated field.
	Calculated bool `json:"calculated" api:"required"`
	CanArchive bool `json:"canArchive" api:"required"`
	CanRestore bool `json:"canRestore" api:"required"`
	// The timestamp when the property was created, in ISO 8601 format.
	CreatedAt int64 `json:"createdAt" api:"required"`
	// The ID of the user who created the property.
	CreatedUserID int64 `json:"createdUserId" api:"required"`
	// The name of the related currency property.
	CurrencyPropertyName string `json:"currencyPropertyName" api:"required"`
	// Indicates the sensitivity level of the property, such as "non_sensitive",
	// "sensitive", or "highly_sensitive".
	//
	// Any of "high", "none", "standard".
	DataSensitivity MediaBridgePropertyDataSensitivity `json:"dataSensitivity" api:"required"`
	// Any of "absolute", "absolute_with_relative", "time_since", "time_until".
	DateDisplayHint MediaBridgePropertyDateDisplayHint `json:"dateDisplayHint" api:"required"`
	// Whether the property has been deleted.
	Deleted bool `json:"deleted" api:"required"`
	// A summary of the property's purpose.
	Description string `json:"description" api:"required"`
	// The mode in which the property is displayed. Can be: "current_value" or
	// "all_unique_versions".
	//
	// Any of "all_unique_versions", "current_value".
	DisplayMode MediaBridgePropertyDisplayMode `json:"displayMode" api:"required"`
	// The position of the item relative to others in the list.
	DisplayOrder                int64 `json:"displayOrder" api:"required"`
	EnforceMultivalueUniqueness bool  `json:"enforceMultivalueUniqueness" api:"required"`
	// Applicable only for enumeration type properties. Should be set to true with a
	// 'referencedObjectType' of 'OWNER'. Otherwise false.
	ExternalOptions bool `json:"externalOptions" api:"required"`
	// When externalOptions is true, indicates the property's option values will be
	// populated from other systems (e.g., "OWNER" for the hubspot_owner_id property).
	ExternalOptionsReferenceType string `json:"externalOptionsReferenceType" api:"required"`
	// Deprecated. Whether the property is marked as a favorite.
	Favorited bool `json:"favorited" api:"required"`
	// Deprecated. The order position when marked as favorited.
	FavoritedOrder int64 `json:"favoritedOrder" api:"required"`
	// Determines how the property will appear in HubSpot's UI or on a form. Learn more
	// in the properties API guide.
	FieldType string `json:"fieldType" api:"required"`
	// Whether the property can appear on forms.
	FormField bool `json:"formField" api:"required"`
	// The ID of the user who last updated the property.
	FromUserID int64 `json:"fromUserId" api:"required"`
	// The name of the group to which the property is assigned.
	GroupName string `json:"groupName" api:"required"`
	// Whether the property is a unique identifier property.
	HasUniqueValue bool `json:"hasUniqueValue" api:"required"`
	// Whether or not the property will be hidden from the HubSpot UI. It's recommended
	// that this be set to false for custom properties.
	Hidden bool `json:"hidden" api:"required"`
	// A boolean value set to true for HubSpot default properties.
	HubSpotDefined bool `json:"hubspotDefined" api:"required"`
	// For default properties, whether the property has been customized. Equivalent to
	// the 'isCustomizedDefault' field.
	IsCustomizedDefault bool `json:"isCustomizedDefault" api:"required"`
	// Whether the property can contain multiple values.
	IsMultiValued bool `json:"isMultiValued" api:"required"`
	// For default properties, whether the property has been customized. Equivalent to
	// the 'isCustomizedDefault' field.
	IsPartial bool `json:"isPartial" api:"required"`
	// The display label for the property.
	Label string `json:"label" api:"required"`
	// Whether the property definition can be customized but not deleted.
	MutableDefinitionNotDeletable bool `json:"mutableDefinitionNotDeletable" api:"required"`
	// The internal name for the property.
	Name string `json:"name" api:"required"`
	// Hint for how a number property is displayed and validated in HubSpot's UI. Can
	// be: "unformatted", "formatted", "currency", "percentage", "duration", or
	// "probability".
	//
	// Any of "currency", "duration", "formatted", "percentage", "probability",
	// "unformatted".
	NumberDisplayHint MediaBridgePropertyNumberDisplayHint `json:"numberDisplayHint" api:"required"`
	// A list of valid options for the property. This field is required for enumerated
	// properties.
	Options []shared.AutomationActionsOption `json:"options" api:"required"`
	// Whether options can be modified after creation.
	OptionsAreMutable bool `json:"optionsAreMutable" api:"required"`
	// Specifies how to sort property options. Can be either "DISPLAY_ORDER" to defer
	// to the displayOrder field, or "ALPHABETICAL".
	//
	// Any of "ALPHABETICAL", "DISPLAY_ORDER".
	OptionSortStrategy MediaBridgePropertyOptionSortStrategy `json:"optionSortStrategy" api:"required"`
	OwningAppID        int64                                 `json:"owningAppId" api:"required"`
	// The ID of the HubSpot account where the property is defined.
	PortalID int64 `json:"portalId" api:"required"`
	// Whether the property's description is read-only.
	ReadOnlyDefinition bool `json:"readOnlyDefinition" api:"required"`
	// Indicates if the property's value is read-only.
	ReadOnlyValue bool `json:"readOnlyValue" api:"required"`
	// Deprecated. Use externalOptionsReferenceType instead.
	//
	// Any of "ABANDONED_CART", "ACCEPTANCE_TEST", "AD", "AD_ACCOUNT", "AD_CAMPAIGN",
	// "AD_GROUP", "AI_FORECAST", "ALL_PAGES", "APPROVAL", "APPROVAL_STEP",
	// "ATTRIBUTION", "AUDIENCE", "AUTOMATION_JOURNEY", "AUTOMATION_PLATFORM_FLOW",
	// "AUTOMATION_PLATFORM_FLOW_ACTION", "BET_ALERT", "BET_DELIVERABLE_SERVICE",
	// "BLOG_LISTING_PAGE", "BLOG_POST", "CALL", "CAMPAIGN", "CAMPAIGN_BUDGET_ITEM",
	// "CAMPAIGN_SPEND_ITEM", "CAMPAIGN_STEP", "CAMPAIGN_TEMPLATE",
	// "CAMPAIGN_TEMPLATE_STEP", "CART", "CASE_STUDY", "CHATFLOW", "CLIP", "CMS_URL",
	// "COMBO_EVENT_CONFIGURATION", "COMMERCE_PAYMENT", "COMMUNICATION", "COMPANY",
	// "CONTACT", "CONTACT_CREATE_ATTRIBUTION", "CONTENT", "CONTENT_AUDIT",
	// "CONTENT_AUDIT_PAGE", "CONVERSATION", "CONVERSATION_INBOX",
	// "CONVERSATION_SESSION", "CRM_OBJECTS_DUMMY_TYPE", "CRM_PIPELINES_DUMMY_TYPE",
	// "CTA", "CTA_VARIANT", "DATA_PRIVACY_CONSENT", "DATA_SYNC_STATE", "DEAL",
	// "DEAL_CREATE_ATTRIBUTION", "DEAL_REGISTRATION", "DEAL_SPLIT", "DISCOUNT",
	// "DISCOUNT_CODE", "DISCOUNT_TEMPLATE", "EMAIL", "ENGAGEMENT", "EXPORT",
	// "EXTERNAL_WEB_URL", "FEE", "FEEDBACK_SUBMISSION", "FEEDBACK_SURVEY",
	// "FILE_MANAGER_FILE", "FILE_MANAGER_FOLDER", "FOLDER", "FORECAST", "FORM",
	// "FORM_SUBMISSION_INBOUNDDB", "GOAL_TARGET", "GOAL_TARGET_GROUP",
	// "GOAL_TEMPLATE", "GSC_PROPERTY", "HUB", "IMPORT", "INVOICE", "KEYWORD",
	// "KNOWLEDGE_ARTICLE", "LANDING_PAGE", "LEAD", "LINE_ITEM", "MARKETING_CALENDAR",
	// "MARKETING_CAMPAIGN_UTM", "MARKETING_EMAIL", "MARKETING_EVENT",
	// "MARKETING_EVENT_ATTENDANCE", "MARKETING_SMS", "MEDIA_BRIDGE", "MEETING_EVENT",
	// "MIC", "NOTE", "OBJECT_LIST", "ORDER", "OWNER", "PARTNER_ACCOUNT",
	// "PARTNER_CLIENT", "PARTNER_CLIENT_REVENUE", "PARTNER_SERVICE", "PAYMENT_LINK",
	// "PAYMENT_SCHEDULE", "PAYMENT_SCHEDULE_INSTALLMENT", "PERMISSIONS_TESTING",
	// "PLAYBOOK", "PLAYBOOK_QUESTION", "PLAYBOOK_SUBMISSION",
	// "PLAYBOOK_SUBMISSION_ANSWER", "PLAYLIST", "PLAYLIST_FOLDER", "PODCAST_EPISODE",
	// "PORTAL", "PORTAL_OBJECT_SYNC_MESSAGE", "POSTAL_MAIL", "PRIVACY_SCANNER_COOKIE",
	// "PRODUCT", "PRODUCT_OR_FOLDER", "PROPERTY_INFO",
	// "PROSPECTING_AGENT_CONTACT_ASSIGNMENT", "PUBLISHING_TASK",
	// "QUARANTINED_SUBMISSION", "QUOTA", "QUOTE", "QUOTE_FIELD", "QUOTE_MODULE",
	// "QUOTE_MODULE_FIELD", "QUOTE_TEMPLATE", "RESTORABLE_CRM_OBJECT", "ROSTER",
	// "ROSTER_MEMBER", "SALES_DOCUMENT", "SALES_TASK", "SALES_WORKLOAD",
	// "SALESFORCE_SYNC_ERROR", "SCHEDULING_PAGE", "SCHEMAS_BACKEND_TEST",
	// "SCORE_CONFIGURATION", "SEQUENCE", "SEQUENCE_ENROLLMENT", "SEQUENCE_STEP",
	// "SEQUENCE_STEP_ENROLLMENT", "SERVICE", "SITE_PAGE", "SNIPPET",
	// "SOCIAL_BROADCAST", "SOCIAL_CHANNEL", "SOCIAL_POST", "SOCIAL_PROFILE",
	// "SOX_PROTECTED_DUMMY_TYPE", "SOX_PROTECTED_TEST_TYPE", "SUBMISSION_TAG",
	// "SUBSCRIPTION", "TASK", "TASK_TEMPLATE", "TAX", "TEMPLATE", "TICKET", "UNKNOWN",
	// "UNSUBSCRIBE", "USER", "VIEW", "VIEW_BLOCK", "WEB_INTERACTIVE".
	ReferencedObjectType MediaBridgePropertyReferencedObjectType `json:"referencedObjectType" api:"required"`
	// Whether the property is searchable globaly.
	SearchableInGlobalSearch bool `json:"searchableInGlobalSearch" api:"required"`
	// Any of "NONE", "NOT_ANALYZED_TEXT".
	SearchTextAnalysisMode MediaBridgePropertySearchTextAnalysisMode `json:"searchTextAnalysisMode" api:"required"`
	// When sensitiveData is true, lists the type of sensitive data contained in the
	// property (e.g., "HIPAA").
	SensitiveDataCategories []string `json:"sensitiveDataCategories" api:"required"`
	// Whether to show the currency symbol in HubSpot's UI.
	ShowCurrencySymbol bool `json:"showCurrencySymbol" api:"required"`
	// Hint for how the text is displayed and validated in HubSpot's UI. Can be:
	// "unformatted_single_line", "multi_line", "email", "phone_number", "domain_name",
	// "ip_address", "physical_address", or "postal_code".
	//
	// Any of "domain_name", "email", "ip_address", "multi_line", "phone_number",
	// "physical_address", "postal_code", "unformatted_single_line".
	TextDisplayHint MediaBridgePropertyTextDisplayHint `json:"textDisplayHint" api:"required"`
	// The data type of the property, such as string or number.
	//
	// Any of "bool", "currency_number", "date", "datetime", "enumeration", "json",
	// "number", "object_coordinates", "phone_number", "string".
	Type MediaBridgePropertyType `json:"type" api:"required"`
	// The timestamp when the property was last updated, in ISO 8601 format.
	UpdatedAt int64 `json:"updatedAt" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AllowedObjectTypes            respjson.Field
		Calculated                    respjson.Field
		CanArchive                    respjson.Field
		CanRestore                    respjson.Field
		CreatedAt                     respjson.Field
		CreatedUserID                 respjson.Field
		CurrencyPropertyName          respjson.Field
		DataSensitivity               respjson.Field
		DateDisplayHint               respjson.Field
		Deleted                       respjson.Field
		Description                   respjson.Field
		DisplayMode                   respjson.Field
		DisplayOrder                  respjson.Field
		EnforceMultivalueUniqueness   respjson.Field
		ExternalOptions               respjson.Field
		ExternalOptionsReferenceType  respjson.Field
		Favorited                     respjson.Field
		FavoritedOrder                respjson.Field
		FieldType                     respjson.Field
		FormField                     respjson.Field
		FromUserID                    respjson.Field
		GroupName                     respjson.Field
		HasUniqueValue                respjson.Field
		Hidden                        respjson.Field
		HubSpotDefined                respjson.Field
		IsCustomizedDefault           respjson.Field
		IsMultiValued                 respjson.Field
		IsPartial                     respjson.Field
		Label                         respjson.Field
		MutableDefinitionNotDeletable respjson.Field
		Name                          respjson.Field
		NumberDisplayHint             respjson.Field
		Options                       respjson.Field
		OptionsAreMutable             respjson.Field
		OptionSortStrategy            respjson.Field
		OwningAppID                   respjson.Field
		PortalID                      respjson.Field
		ReadOnlyDefinition            respjson.Field
		ReadOnlyValue                 respjson.Field
		ReferencedObjectType          respjson.Field
		SearchableInGlobalSearch      respjson.Field
		SearchTextAnalysisMode        respjson.Field
		SensitiveDataCategories       respjson.Field
		ShowCurrencySymbol            respjson.Field
		TextDisplayHint               respjson.Field
		Type                          respjson.Field
		UpdatedAt                     respjson.Field
		ExtraFields                   map[string]respjson.Field
		raw                           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MediaBridgeProperty) RawJSON() string { return r.JSON.raw }
func (r *MediaBridgeProperty) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Indicates the sensitivity level of the property, such as "non_sensitive",
// "sensitive", or "highly_sensitive".
type MediaBridgePropertyDataSensitivity string

const (
	MediaBridgePropertyDataSensitivityHigh     MediaBridgePropertyDataSensitivity = "high"
	MediaBridgePropertyDataSensitivityNone     MediaBridgePropertyDataSensitivity = "none"
	MediaBridgePropertyDataSensitivityStandard MediaBridgePropertyDataSensitivity = "standard"
)

type MediaBridgePropertyDateDisplayHint string

const (
	MediaBridgePropertyDateDisplayHintAbsolute             MediaBridgePropertyDateDisplayHint = "absolute"
	MediaBridgePropertyDateDisplayHintAbsoluteWithRelative MediaBridgePropertyDateDisplayHint = "absolute_with_relative"
	MediaBridgePropertyDateDisplayHintTimeSince            MediaBridgePropertyDateDisplayHint = "time_since"
	MediaBridgePropertyDateDisplayHintTimeUntil            MediaBridgePropertyDateDisplayHint = "time_until"
)

// The mode in which the property is displayed. Can be: "current_value" or
// "all_unique_versions".
type MediaBridgePropertyDisplayMode string

const (
	MediaBridgePropertyDisplayModeAllUniqueVersions MediaBridgePropertyDisplayMode = "all_unique_versions"
	MediaBridgePropertyDisplayModeCurrentValue      MediaBridgePropertyDisplayMode = "current_value"
)

// Hint for how a number property is displayed and validated in HubSpot's UI. Can
// be: "unformatted", "formatted", "currency", "percentage", "duration", or
// "probability".
type MediaBridgePropertyNumberDisplayHint string

const (
	MediaBridgePropertyNumberDisplayHintCurrency    MediaBridgePropertyNumberDisplayHint = "currency"
	MediaBridgePropertyNumberDisplayHintDuration    MediaBridgePropertyNumberDisplayHint = "duration"
	MediaBridgePropertyNumberDisplayHintFormatted   MediaBridgePropertyNumberDisplayHint = "formatted"
	MediaBridgePropertyNumberDisplayHintPercentage  MediaBridgePropertyNumberDisplayHint = "percentage"
	MediaBridgePropertyNumberDisplayHintProbability MediaBridgePropertyNumberDisplayHint = "probability"
	MediaBridgePropertyNumberDisplayHintUnformatted MediaBridgePropertyNumberDisplayHint = "unformatted"
)

// Specifies how to sort property options. Can be either "DISPLAY_ORDER" to defer
// to the displayOrder field, or "ALPHABETICAL".
type MediaBridgePropertyOptionSortStrategy string

const (
	MediaBridgePropertyOptionSortStrategyAlphabetical MediaBridgePropertyOptionSortStrategy = "ALPHABETICAL"
	MediaBridgePropertyOptionSortStrategyDisplayOrder MediaBridgePropertyOptionSortStrategy = "DISPLAY_ORDER"
)

// Deprecated. Use externalOptionsReferenceType instead.
type MediaBridgePropertyReferencedObjectType string

const (
	MediaBridgePropertyReferencedObjectTypeAbandonedCart                     MediaBridgePropertyReferencedObjectType = "ABANDONED_CART"
	MediaBridgePropertyReferencedObjectTypeAcceptanceTest                    MediaBridgePropertyReferencedObjectType = "ACCEPTANCE_TEST"
	MediaBridgePropertyReferencedObjectTypeAd                                MediaBridgePropertyReferencedObjectType = "AD"
	MediaBridgePropertyReferencedObjectTypeAdAccount                         MediaBridgePropertyReferencedObjectType = "AD_ACCOUNT"
	MediaBridgePropertyReferencedObjectTypeAdCampaign                        MediaBridgePropertyReferencedObjectType = "AD_CAMPAIGN"
	MediaBridgePropertyReferencedObjectTypeAdGroup                           MediaBridgePropertyReferencedObjectType = "AD_GROUP"
	MediaBridgePropertyReferencedObjectTypeAIForecast                        MediaBridgePropertyReferencedObjectType = "AI_FORECAST"
	MediaBridgePropertyReferencedObjectTypeAllPages                          MediaBridgePropertyReferencedObjectType = "ALL_PAGES"
	MediaBridgePropertyReferencedObjectTypeApproval                          MediaBridgePropertyReferencedObjectType = "APPROVAL"
	MediaBridgePropertyReferencedObjectTypeApprovalStep                      MediaBridgePropertyReferencedObjectType = "APPROVAL_STEP"
	MediaBridgePropertyReferencedObjectTypeAttribution                       MediaBridgePropertyReferencedObjectType = "ATTRIBUTION"
	MediaBridgePropertyReferencedObjectTypeAudience                          MediaBridgePropertyReferencedObjectType = "AUDIENCE"
	MediaBridgePropertyReferencedObjectTypeAutomationJourney                 MediaBridgePropertyReferencedObjectType = "AUTOMATION_JOURNEY"
	MediaBridgePropertyReferencedObjectTypeAutomationPlatformFlow            MediaBridgePropertyReferencedObjectType = "AUTOMATION_PLATFORM_FLOW"
	MediaBridgePropertyReferencedObjectTypeAutomationPlatformFlowAction      MediaBridgePropertyReferencedObjectType = "AUTOMATION_PLATFORM_FLOW_ACTION"
	MediaBridgePropertyReferencedObjectTypeBetAlert                          MediaBridgePropertyReferencedObjectType = "BET_ALERT"
	MediaBridgePropertyReferencedObjectTypeBetDeliverableService             MediaBridgePropertyReferencedObjectType = "BET_DELIVERABLE_SERVICE"
	MediaBridgePropertyReferencedObjectTypeBlogListingPage                   MediaBridgePropertyReferencedObjectType = "BLOG_LISTING_PAGE"
	MediaBridgePropertyReferencedObjectTypeBlogPost                          MediaBridgePropertyReferencedObjectType = "BLOG_POST"
	MediaBridgePropertyReferencedObjectTypeCall                              MediaBridgePropertyReferencedObjectType = "CALL"
	MediaBridgePropertyReferencedObjectTypeCampaign                          MediaBridgePropertyReferencedObjectType = "CAMPAIGN"
	MediaBridgePropertyReferencedObjectTypeCampaignBudgetItem                MediaBridgePropertyReferencedObjectType = "CAMPAIGN_BUDGET_ITEM"
	MediaBridgePropertyReferencedObjectTypeCampaignSpendItem                 MediaBridgePropertyReferencedObjectType = "CAMPAIGN_SPEND_ITEM"
	MediaBridgePropertyReferencedObjectTypeCampaignStep                      MediaBridgePropertyReferencedObjectType = "CAMPAIGN_STEP"
	MediaBridgePropertyReferencedObjectTypeCampaignTemplate                  MediaBridgePropertyReferencedObjectType = "CAMPAIGN_TEMPLATE"
	MediaBridgePropertyReferencedObjectTypeCampaignTemplateStep              MediaBridgePropertyReferencedObjectType = "CAMPAIGN_TEMPLATE_STEP"
	MediaBridgePropertyReferencedObjectTypeCart                              MediaBridgePropertyReferencedObjectType = "CART"
	MediaBridgePropertyReferencedObjectTypeCaseStudy                         MediaBridgePropertyReferencedObjectType = "CASE_STUDY"
	MediaBridgePropertyReferencedObjectTypeChatflow                          MediaBridgePropertyReferencedObjectType = "CHATFLOW"
	MediaBridgePropertyReferencedObjectTypeClip                              MediaBridgePropertyReferencedObjectType = "CLIP"
	MediaBridgePropertyReferencedObjectTypeCmsURL                            MediaBridgePropertyReferencedObjectType = "CMS_URL"
	MediaBridgePropertyReferencedObjectTypeComboEventConfiguration           MediaBridgePropertyReferencedObjectType = "COMBO_EVENT_CONFIGURATION"
	MediaBridgePropertyReferencedObjectTypeCommercePayment                   MediaBridgePropertyReferencedObjectType = "COMMERCE_PAYMENT"
	MediaBridgePropertyReferencedObjectTypeCommunication                     MediaBridgePropertyReferencedObjectType = "COMMUNICATION"
	MediaBridgePropertyReferencedObjectTypeCompany                           MediaBridgePropertyReferencedObjectType = "COMPANY"
	MediaBridgePropertyReferencedObjectTypeContact                           MediaBridgePropertyReferencedObjectType = "CONTACT"
	MediaBridgePropertyReferencedObjectTypeContactCreateAttribution          MediaBridgePropertyReferencedObjectType = "CONTACT_CREATE_ATTRIBUTION"
	MediaBridgePropertyReferencedObjectTypeContent                           MediaBridgePropertyReferencedObjectType = "CONTENT"
	MediaBridgePropertyReferencedObjectTypeContentAudit                      MediaBridgePropertyReferencedObjectType = "CONTENT_AUDIT"
	MediaBridgePropertyReferencedObjectTypeContentAuditPage                  MediaBridgePropertyReferencedObjectType = "CONTENT_AUDIT_PAGE"
	MediaBridgePropertyReferencedObjectTypeConversation                      MediaBridgePropertyReferencedObjectType = "CONVERSATION"
	MediaBridgePropertyReferencedObjectTypeConversationInbox                 MediaBridgePropertyReferencedObjectType = "CONVERSATION_INBOX"
	MediaBridgePropertyReferencedObjectTypeConversationSession               MediaBridgePropertyReferencedObjectType = "CONVERSATION_SESSION"
	MediaBridgePropertyReferencedObjectTypeCrmObjectsDummyType               MediaBridgePropertyReferencedObjectType = "CRM_OBJECTS_DUMMY_TYPE"
	MediaBridgePropertyReferencedObjectTypeCrmPipelinesDummyType             MediaBridgePropertyReferencedObjectType = "CRM_PIPELINES_DUMMY_TYPE"
	MediaBridgePropertyReferencedObjectTypeCta                               MediaBridgePropertyReferencedObjectType = "CTA"
	MediaBridgePropertyReferencedObjectTypeCtaVariant                        MediaBridgePropertyReferencedObjectType = "CTA_VARIANT"
	MediaBridgePropertyReferencedObjectTypeDataPrivacyConsent                MediaBridgePropertyReferencedObjectType = "DATA_PRIVACY_CONSENT"
	MediaBridgePropertyReferencedObjectTypeDataSyncState                     MediaBridgePropertyReferencedObjectType = "DATA_SYNC_STATE"
	MediaBridgePropertyReferencedObjectTypeDeal                              MediaBridgePropertyReferencedObjectType = "DEAL"
	MediaBridgePropertyReferencedObjectTypeDealCreateAttribution             MediaBridgePropertyReferencedObjectType = "DEAL_CREATE_ATTRIBUTION"
	MediaBridgePropertyReferencedObjectTypeDealRegistration                  MediaBridgePropertyReferencedObjectType = "DEAL_REGISTRATION"
	MediaBridgePropertyReferencedObjectTypeDealSplit                         MediaBridgePropertyReferencedObjectType = "DEAL_SPLIT"
	MediaBridgePropertyReferencedObjectTypeDiscount                          MediaBridgePropertyReferencedObjectType = "DISCOUNT"
	MediaBridgePropertyReferencedObjectTypeDiscountCode                      MediaBridgePropertyReferencedObjectType = "DISCOUNT_CODE"
	MediaBridgePropertyReferencedObjectTypeDiscountTemplate                  MediaBridgePropertyReferencedObjectType = "DISCOUNT_TEMPLATE"
	MediaBridgePropertyReferencedObjectTypeEmail                             MediaBridgePropertyReferencedObjectType = "EMAIL"
	MediaBridgePropertyReferencedObjectTypeEngagement                        MediaBridgePropertyReferencedObjectType = "ENGAGEMENT"
	MediaBridgePropertyReferencedObjectTypeExport                            MediaBridgePropertyReferencedObjectType = "EXPORT"
	MediaBridgePropertyReferencedObjectTypeExternalWebURL                    MediaBridgePropertyReferencedObjectType = "EXTERNAL_WEB_URL"
	MediaBridgePropertyReferencedObjectTypeFee                               MediaBridgePropertyReferencedObjectType = "FEE"
	MediaBridgePropertyReferencedObjectTypeFeedbackSubmission                MediaBridgePropertyReferencedObjectType = "FEEDBACK_SUBMISSION"
	MediaBridgePropertyReferencedObjectTypeFeedbackSurvey                    MediaBridgePropertyReferencedObjectType = "FEEDBACK_SURVEY"
	MediaBridgePropertyReferencedObjectTypeFileManagerFile                   MediaBridgePropertyReferencedObjectType = "FILE_MANAGER_FILE"
	MediaBridgePropertyReferencedObjectTypeFileManagerFolder                 MediaBridgePropertyReferencedObjectType = "FILE_MANAGER_FOLDER"
	MediaBridgePropertyReferencedObjectTypeFolder                            MediaBridgePropertyReferencedObjectType = "FOLDER"
	MediaBridgePropertyReferencedObjectTypeForecast                          MediaBridgePropertyReferencedObjectType = "FORECAST"
	MediaBridgePropertyReferencedObjectTypeForm                              MediaBridgePropertyReferencedObjectType = "FORM"
	MediaBridgePropertyReferencedObjectTypeFormSubmissionInbounddb           MediaBridgePropertyReferencedObjectType = "FORM_SUBMISSION_INBOUNDDB"
	MediaBridgePropertyReferencedObjectTypeGoalTarget                        MediaBridgePropertyReferencedObjectType = "GOAL_TARGET"
	MediaBridgePropertyReferencedObjectTypeGoalTargetGroup                   MediaBridgePropertyReferencedObjectType = "GOAL_TARGET_GROUP"
	MediaBridgePropertyReferencedObjectTypeGoalTemplate                      MediaBridgePropertyReferencedObjectType = "GOAL_TEMPLATE"
	MediaBridgePropertyReferencedObjectTypeGscProperty                       MediaBridgePropertyReferencedObjectType = "GSC_PROPERTY"
	MediaBridgePropertyReferencedObjectTypeHub                               MediaBridgePropertyReferencedObjectType = "HUB"
	MediaBridgePropertyReferencedObjectTypeImport                            MediaBridgePropertyReferencedObjectType = "IMPORT"
	MediaBridgePropertyReferencedObjectTypeInvoice                           MediaBridgePropertyReferencedObjectType = "INVOICE"
	MediaBridgePropertyReferencedObjectTypeKeyword                           MediaBridgePropertyReferencedObjectType = "KEYWORD"
	MediaBridgePropertyReferencedObjectTypeKnowledgeArticle                  MediaBridgePropertyReferencedObjectType = "KNOWLEDGE_ARTICLE"
	MediaBridgePropertyReferencedObjectTypeLandingPage                       MediaBridgePropertyReferencedObjectType = "LANDING_PAGE"
	MediaBridgePropertyReferencedObjectTypeLead                              MediaBridgePropertyReferencedObjectType = "LEAD"
	MediaBridgePropertyReferencedObjectTypeLineItem                          MediaBridgePropertyReferencedObjectType = "LINE_ITEM"
	MediaBridgePropertyReferencedObjectTypeMarketingCalendar                 MediaBridgePropertyReferencedObjectType = "MARKETING_CALENDAR"
	MediaBridgePropertyReferencedObjectTypeMarketingCampaignUtm              MediaBridgePropertyReferencedObjectType = "MARKETING_CAMPAIGN_UTM"
	MediaBridgePropertyReferencedObjectTypeMarketingEmail                    MediaBridgePropertyReferencedObjectType = "MARKETING_EMAIL"
	MediaBridgePropertyReferencedObjectTypeMarketingEvent                    MediaBridgePropertyReferencedObjectType = "MARKETING_EVENT"
	MediaBridgePropertyReferencedObjectTypeMarketingEventAttendance          MediaBridgePropertyReferencedObjectType = "MARKETING_EVENT_ATTENDANCE"
	MediaBridgePropertyReferencedObjectTypeMarketingSMS                      MediaBridgePropertyReferencedObjectType = "MARKETING_SMS"
	MediaBridgePropertyReferencedObjectTypeMediaBridge                       MediaBridgePropertyReferencedObjectType = "MEDIA_BRIDGE"
	MediaBridgePropertyReferencedObjectTypeMeetingEvent                      MediaBridgePropertyReferencedObjectType = "MEETING_EVENT"
	MediaBridgePropertyReferencedObjectTypeMic                               MediaBridgePropertyReferencedObjectType = "MIC"
	MediaBridgePropertyReferencedObjectTypeNote                              MediaBridgePropertyReferencedObjectType = "NOTE"
	MediaBridgePropertyReferencedObjectTypeObjectList                        MediaBridgePropertyReferencedObjectType = "OBJECT_LIST"
	MediaBridgePropertyReferencedObjectTypeOrder                             MediaBridgePropertyReferencedObjectType = "ORDER"
	MediaBridgePropertyReferencedObjectTypeOwner                             MediaBridgePropertyReferencedObjectType = "OWNER"
	MediaBridgePropertyReferencedObjectTypePartnerAccount                    MediaBridgePropertyReferencedObjectType = "PARTNER_ACCOUNT"
	MediaBridgePropertyReferencedObjectTypePartnerClient                     MediaBridgePropertyReferencedObjectType = "PARTNER_CLIENT"
	MediaBridgePropertyReferencedObjectTypePartnerClientRevenue              MediaBridgePropertyReferencedObjectType = "PARTNER_CLIENT_REVENUE"
	MediaBridgePropertyReferencedObjectTypePartnerService                    MediaBridgePropertyReferencedObjectType = "PARTNER_SERVICE"
	MediaBridgePropertyReferencedObjectTypePaymentLink                       MediaBridgePropertyReferencedObjectType = "PAYMENT_LINK"
	MediaBridgePropertyReferencedObjectTypePaymentSchedule                   MediaBridgePropertyReferencedObjectType = "PAYMENT_SCHEDULE"
	MediaBridgePropertyReferencedObjectTypePaymentScheduleInstallment        MediaBridgePropertyReferencedObjectType = "PAYMENT_SCHEDULE_INSTALLMENT"
	MediaBridgePropertyReferencedObjectTypePermissionsTesting                MediaBridgePropertyReferencedObjectType = "PERMISSIONS_TESTING"
	MediaBridgePropertyReferencedObjectTypePlaybook                          MediaBridgePropertyReferencedObjectType = "PLAYBOOK"
	MediaBridgePropertyReferencedObjectTypePlaybookQuestion                  MediaBridgePropertyReferencedObjectType = "PLAYBOOK_QUESTION"
	MediaBridgePropertyReferencedObjectTypePlaybookSubmission                MediaBridgePropertyReferencedObjectType = "PLAYBOOK_SUBMISSION"
	MediaBridgePropertyReferencedObjectTypePlaybookSubmissionAnswer          MediaBridgePropertyReferencedObjectType = "PLAYBOOK_SUBMISSION_ANSWER"
	MediaBridgePropertyReferencedObjectTypePlaylist                          MediaBridgePropertyReferencedObjectType = "PLAYLIST"
	MediaBridgePropertyReferencedObjectTypePlaylistFolder                    MediaBridgePropertyReferencedObjectType = "PLAYLIST_FOLDER"
	MediaBridgePropertyReferencedObjectTypePodcastEpisode                    MediaBridgePropertyReferencedObjectType = "PODCAST_EPISODE"
	MediaBridgePropertyReferencedObjectTypePortal                            MediaBridgePropertyReferencedObjectType = "PORTAL"
	MediaBridgePropertyReferencedObjectTypePortalObjectSyncMessage           MediaBridgePropertyReferencedObjectType = "PORTAL_OBJECT_SYNC_MESSAGE"
	MediaBridgePropertyReferencedObjectTypePostalMail                        MediaBridgePropertyReferencedObjectType = "POSTAL_MAIL"
	MediaBridgePropertyReferencedObjectTypePrivacyScannerCookie              MediaBridgePropertyReferencedObjectType = "PRIVACY_SCANNER_COOKIE"
	MediaBridgePropertyReferencedObjectTypeProduct                           MediaBridgePropertyReferencedObjectType = "PRODUCT"
	MediaBridgePropertyReferencedObjectTypeProductOrFolder                   MediaBridgePropertyReferencedObjectType = "PRODUCT_OR_FOLDER"
	MediaBridgePropertyReferencedObjectTypePropertyInfo                      MediaBridgePropertyReferencedObjectType = "PROPERTY_INFO"
	MediaBridgePropertyReferencedObjectTypeProspectingAgentContactAssignment MediaBridgePropertyReferencedObjectType = "PROSPECTING_AGENT_CONTACT_ASSIGNMENT"
	MediaBridgePropertyReferencedObjectTypePublishingTask                    MediaBridgePropertyReferencedObjectType = "PUBLISHING_TASK"
	MediaBridgePropertyReferencedObjectTypeQuarantinedSubmission             MediaBridgePropertyReferencedObjectType = "QUARANTINED_SUBMISSION"
	MediaBridgePropertyReferencedObjectTypeQuota                             MediaBridgePropertyReferencedObjectType = "QUOTA"
	MediaBridgePropertyReferencedObjectTypeQuote                             MediaBridgePropertyReferencedObjectType = "QUOTE"
	MediaBridgePropertyReferencedObjectTypeQuoteField                        MediaBridgePropertyReferencedObjectType = "QUOTE_FIELD"
	MediaBridgePropertyReferencedObjectTypeQuoteModule                       MediaBridgePropertyReferencedObjectType = "QUOTE_MODULE"
	MediaBridgePropertyReferencedObjectTypeQuoteModuleField                  MediaBridgePropertyReferencedObjectType = "QUOTE_MODULE_FIELD"
	MediaBridgePropertyReferencedObjectTypeQuoteTemplate                     MediaBridgePropertyReferencedObjectType = "QUOTE_TEMPLATE"
	MediaBridgePropertyReferencedObjectTypeRestorableCrmObject               MediaBridgePropertyReferencedObjectType = "RESTORABLE_CRM_OBJECT"
	MediaBridgePropertyReferencedObjectTypeRoster                            MediaBridgePropertyReferencedObjectType = "ROSTER"
	MediaBridgePropertyReferencedObjectTypeRosterMember                      MediaBridgePropertyReferencedObjectType = "ROSTER_MEMBER"
	MediaBridgePropertyReferencedObjectTypeSalesDocument                     MediaBridgePropertyReferencedObjectType = "SALES_DOCUMENT"
	MediaBridgePropertyReferencedObjectTypeSalesTask                         MediaBridgePropertyReferencedObjectType = "SALES_TASK"
	MediaBridgePropertyReferencedObjectTypeSalesWorkload                     MediaBridgePropertyReferencedObjectType = "SALES_WORKLOAD"
	MediaBridgePropertyReferencedObjectTypeSalesforceSyncError               MediaBridgePropertyReferencedObjectType = "SALESFORCE_SYNC_ERROR"
	MediaBridgePropertyReferencedObjectTypeSchedulingPage                    MediaBridgePropertyReferencedObjectType = "SCHEDULING_PAGE"
	MediaBridgePropertyReferencedObjectTypeSchemasBackendTest                MediaBridgePropertyReferencedObjectType = "SCHEMAS_BACKEND_TEST"
	MediaBridgePropertyReferencedObjectTypeScoreConfiguration                MediaBridgePropertyReferencedObjectType = "SCORE_CONFIGURATION"
	MediaBridgePropertyReferencedObjectTypeSequence                          MediaBridgePropertyReferencedObjectType = "SEQUENCE"
	MediaBridgePropertyReferencedObjectTypeSequenceEnrollment                MediaBridgePropertyReferencedObjectType = "SEQUENCE_ENROLLMENT"
	MediaBridgePropertyReferencedObjectTypeSequenceStep                      MediaBridgePropertyReferencedObjectType = "SEQUENCE_STEP"
	MediaBridgePropertyReferencedObjectTypeSequenceStepEnrollment            MediaBridgePropertyReferencedObjectType = "SEQUENCE_STEP_ENROLLMENT"
	MediaBridgePropertyReferencedObjectTypeService                           MediaBridgePropertyReferencedObjectType = "SERVICE"
	MediaBridgePropertyReferencedObjectTypeSitePage                          MediaBridgePropertyReferencedObjectType = "SITE_PAGE"
	MediaBridgePropertyReferencedObjectTypeSnippet                           MediaBridgePropertyReferencedObjectType = "SNIPPET"
	MediaBridgePropertyReferencedObjectTypeSocialBroadcast                   MediaBridgePropertyReferencedObjectType = "SOCIAL_BROADCAST"
	MediaBridgePropertyReferencedObjectTypeSocialChannel                     MediaBridgePropertyReferencedObjectType = "SOCIAL_CHANNEL"
	MediaBridgePropertyReferencedObjectTypeSocialPost                        MediaBridgePropertyReferencedObjectType = "SOCIAL_POST"
	MediaBridgePropertyReferencedObjectTypeSocialProfile                     MediaBridgePropertyReferencedObjectType = "SOCIAL_PROFILE"
	MediaBridgePropertyReferencedObjectTypeSoxProtectedDummyType             MediaBridgePropertyReferencedObjectType = "SOX_PROTECTED_DUMMY_TYPE"
	MediaBridgePropertyReferencedObjectTypeSoxProtectedTestType              MediaBridgePropertyReferencedObjectType = "SOX_PROTECTED_TEST_TYPE"
	MediaBridgePropertyReferencedObjectTypeSubmissionTag                     MediaBridgePropertyReferencedObjectType = "SUBMISSION_TAG"
	MediaBridgePropertyReferencedObjectTypeSubscription                      MediaBridgePropertyReferencedObjectType = "SUBSCRIPTION"
	MediaBridgePropertyReferencedObjectTypeTask                              MediaBridgePropertyReferencedObjectType = "TASK"
	MediaBridgePropertyReferencedObjectTypeTaskTemplate                      MediaBridgePropertyReferencedObjectType = "TASK_TEMPLATE"
	MediaBridgePropertyReferencedObjectTypeTax                               MediaBridgePropertyReferencedObjectType = "TAX"
	MediaBridgePropertyReferencedObjectTypeTemplate                          MediaBridgePropertyReferencedObjectType = "TEMPLATE"
	MediaBridgePropertyReferencedObjectTypeTicket                            MediaBridgePropertyReferencedObjectType = "TICKET"
	MediaBridgePropertyReferencedObjectTypeUnknown                           MediaBridgePropertyReferencedObjectType = "UNKNOWN"
	MediaBridgePropertyReferencedObjectTypeUnsubscribe                       MediaBridgePropertyReferencedObjectType = "UNSUBSCRIBE"
	MediaBridgePropertyReferencedObjectTypeUser                              MediaBridgePropertyReferencedObjectType = "USER"
	MediaBridgePropertyReferencedObjectTypeView                              MediaBridgePropertyReferencedObjectType = "VIEW"
	MediaBridgePropertyReferencedObjectTypeViewBlock                         MediaBridgePropertyReferencedObjectType = "VIEW_BLOCK"
	MediaBridgePropertyReferencedObjectTypeWebInteractive                    MediaBridgePropertyReferencedObjectType = "WEB_INTERACTIVE"
)

type MediaBridgePropertySearchTextAnalysisMode string

const (
	MediaBridgePropertySearchTextAnalysisModeNone            MediaBridgePropertySearchTextAnalysisMode = "NONE"
	MediaBridgePropertySearchTextAnalysisModeNotAnalyzedText MediaBridgePropertySearchTextAnalysisMode = "NOT_ANALYZED_TEXT"
)

// Hint for how the text is displayed and validated in HubSpot's UI. Can be:
// "unformatted_single_line", "multi_line", "email", "phone_number", "domain_name",
// "ip_address", "physical_address", or "postal_code".
type MediaBridgePropertyTextDisplayHint string

const (
	MediaBridgePropertyTextDisplayHintDomainName            MediaBridgePropertyTextDisplayHint = "domain_name"
	MediaBridgePropertyTextDisplayHintEmail                 MediaBridgePropertyTextDisplayHint = "email"
	MediaBridgePropertyTextDisplayHintIPAddress             MediaBridgePropertyTextDisplayHint = "ip_address"
	MediaBridgePropertyTextDisplayHintMultiLine             MediaBridgePropertyTextDisplayHint = "multi_line"
	MediaBridgePropertyTextDisplayHintPhoneNumber           MediaBridgePropertyTextDisplayHint = "phone_number"
	MediaBridgePropertyTextDisplayHintPhysicalAddress       MediaBridgePropertyTextDisplayHint = "physical_address"
	MediaBridgePropertyTextDisplayHintPostalCode            MediaBridgePropertyTextDisplayHint = "postal_code"
	MediaBridgePropertyTextDisplayHintUnformattedSingleLine MediaBridgePropertyTextDisplayHint = "unformatted_single_line"
)

// The data type of the property, such as string or number.
type MediaBridgePropertyType string

const (
	MediaBridgePropertyTypeBool              MediaBridgePropertyType = "bool"
	MediaBridgePropertyTypeCurrencyNumber    MediaBridgePropertyType = "currency_number"
	MediaBridgePropertyTypeDate              MediaBridgePropertyType = "date"
	MediaBridgePropertyTypeDatetime          MediaBridgePropertyType = "datetime"
	MediaBridgePropertyTypeEnumeration       MediaBridgePropertyType = "enumeration"
	MediaBridgePropertyTypeJson              MediaBridgePropertyType = "json"
	MediaBridgePropertyTypeNumber            MediaBridgePropertyType = "number"
	MediaBridgePropertyTypeObjectCoordinates MediaBridgePropertyType = "object_coordinates"
	MediaBridgePropertyTypePhoneNumber       MediaBridgePropertyType = "phone_number"
	MediaBridgePropertyTypeString            MediaBridgePropertyType = "string"
)

type MediaBridgePropertyUpdateParam struct {
	CalculationFormula   param.Opt[string] `json:"calculationFormula,omitzero"`
	CurrencyPropertyName param.Opt[string] `json:"currencyPropertyName,omitzero"`
	Description          param.Opt[string] `json:"description,omitzero"`
	DisplayOrder         param.Opt[int64]  `json:"displayOrder,omitzero"`
	FormField            param.Opt[bool]   `json:"formField,omitzero"`
	GroupName            param.Opt[string] `json:"groupName,omitzero"`
	HasUniqueValue       param.Opt[bool]   `json:"hasUniqueValue,omitzero"`
	Hidden               param.Opt[bool]   `json:"hidden,omitzero"`
	Label                param.Opt[string] `json:"label,omitzero"`
	ShowCurrencySymbol   param.Opt[bool]   `json:"showCurrencySymbol,omitzero"`
	// Any of "booleancheckbox", "calculation_equation", "checkbox", "date", "file",
	// "html", "number", "phonenumber", "radio", "select", "text", "textarea".
	FieldType MediaBridgePropertyUpdateFieldType `json:"fieldType,omitzero"`
	// Any of "currency", "duration", "formatted", "percentage", "probability",
	// "unformatted".
	NumberDisplayHint MediaBridgePropertyUpdateNumberDisplayHint `json:"numberDisplayHint,omitzero"`
	Options           []shared.OptionInputParam                  `json:"options,omitzero"`
	// Any of "bool", "date", "datetime", "enumeration", "number", "phone_number",
	// "string".
	Type MediaBridgePropertyUpdateType `json:"type,omitzero"`
	paramObj
}

func (r MediaBridgePropertyUpdateParam) MarshalJSON() (data []byte, err error) {
	type shadow MediaBridgePropertyUpdateParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MediaBridgePropertyUpdateParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgePropertyUpdateFieldType string

const (
	MediaBridgePropertyUpdateFieldTypeBooleancheckbox     MediaBridgePropertyUpdateFieldType = "booleancheckbox"
	MediaBridgePropertyUpdateFieldTypeCalculationEquation MediaBridgePropertyUpdateFieldType = "calculation_equation"
	MediaBridgePropertyUpdateFieldTypeCheckbox            MediaBridgePropertyUpdateFieldType = "checkbox"
	MediaBridgePropertyUpdateFieldTypeDate                MediaBridgePropertyUpdateFieldType = "date"
	MediaBridgePropertyUpdateFieldTypeFile                MediaBridgePropertyUpdateFieldType = "file"
	MediaBridgePropertyUpdateFieldTypeHTML                MediaBridgePropertyUpdateFieldType = "html"
	MediaBridgePropertyUpdateFieldTypeNumber              MediaBridgePropertyUpdateFieldType = "number"
	MediaBridgePropertyUpdateFieldTypePhonenumber         MediaBridgePropertyUpdateFieldType = "phonenumber"
	MediaBridgePropertyUpdateFieldTypeRadio               MediaBridgePropertyUpdateFieldType = "radio"
	MediaBridgePropertyUpdateFieldTypeSelect              MediaBridgePropertyUpdateFieldType = "select"
	MediaBridgePropertyUpdateFieldTypeText                MediaBridgePropertyUpdateFieldType = "text"
	MediaBridgePropertyUpdateFieldTypeTextarea            MediaBridgePropertyUpdateFieldType = "textarea"
)

type MediaBridgePropertyUpdateNumberDisplayHint string

const (
	MediaBridgePropertyUpdateNumberDisplayHintCurrency    MediaBridgePropertyUpdateNumberDisplayHint = "currency"
	MediaBridgePropertyUpdateNumberDisplayHintDuration    MediaBridgePropertyUpdateNumberDisplayHint = "duration"
	MediaBridgePropertyUpdateNumberDisplayHintFormatted   MediaBridgePropertyUpdateNumberDisplayHint = "formatted"
	MediaBridgePropertyUpdateNumberDisplayHintPercentage  MediaBridgePropertyUpdateNumberDisplayHint = "percentage"
	MediaBridgePropertyUpdateNumberDisplayHintProbability MediaBridgePropertyUpdateNumberDisplayHint = "probability"
	MediaBridgePropertyUpdateNumberDisplayHintUnformatted MediaBridgePropertyUpdateNumberDisplayHint = "unformatted"
)

type MediaBridgePropertyUpdateType string

const (
	MediaBridgePropertyUpdateTypeBool        MediaBridgePropertyUpdateType = "bool"
	MediaBridgePropertyUpdateTypeDate        MediaBridgePropertyUpdateType = "date"
	MediaBridgePropertyUpdateTypeDatetime    MediaBridgePropertyUpdateType = "datetime"
	MediaBridgePropertyUpdateTypeEnumeration MediaBridgePropertyUpdateType = "enumeration"
	MediaBridgePropertyUpdateTypeNumber      MediaBridgePropertyUpdateType = "number"
	MediaBridgePropertyUpdateTypePhoneNumber MediaBridgePropertyUpdateType = "phone_number"
	MediaBridgePropertyUpdateTypeString      MediaBridgePropertyUpdateType = "string"
)

// The property UpdatedAt is required.
type MediaBridgeProviderPartialParam struct {
	UpdatedAt               int64             `json:"updatedAt" api:"required"`
	AllowImportOnDisconnect param.Opt[bool]   `json:"allowImportOnDisconnect,omitzero"`
	ModuleName              param.Opt[string] `json:"moduleName,omitzero"`
	Name                    param.Opt[string] `json:"name,omitzero"`
	paramObj
}

func (r MediaBridgeProviderPartialParam) MarshalJSON() (data []byte, err error) {
	type shadow MediaBridgeProviderPartialParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MediaBridgeProviderPartialParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeProviderRegistrationResponse struct {
	AppID int64  `json:"appId" api:"required"`
	Name  string `json:"name" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AppID       respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MediaBridgeProviderRegistrationResponse) RawJSON() string { return r.JSON.raw }
func (r *MediaBridgeProviderRegistrationResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaPlayedEvent struct {
	ContactID                    int64  `json:"contactId" api:"required"`
	MediaBridgeID                int64  `json:"mediaBridgeId" api:"required"`
	MediaBridgeObjectCoordinates string `json:"mediaBridgeObjectCoordinates" api:"required"`
	MediaBridgeObjectTypeID      string `json:"mediaBridgeObjectTypeId" api:"required"`
	MediaName                    string `json:"mediaName" api:"required"`
	// Any of "AUDIO", "DOCUMENT", "IMAGE", "OTHER", "VIDEO".
	MediaType         MediaPlayedEventMediaType `json:"mediaType" api:"required"`
	OccurredTimestamp int64                     `json:"occurredTimestamp" api:"required"`
	PortalID          int64                     `json:"portalId" api:"required"`
	ProviderID        int64                     `json:"providerId" api:"required"`
	SessionID         string                    `json:"sessionId" api:"required"`
	// Any of "STARTED", "VIEWED".
	State MediaPlayedEventState `json:"state" api:"required"`
	// Any of "EMAIL", "EXTERNAL_PAGE".
	ExternalPlayContext   MediaPlayedEventExternalPlayContext `json:"externalPlayContext"`
	IframeURL             string                              `json:"iframeUrl"`
	MediaURL              string                              `json:"mediaUrl"`
	PageID                int64                               `json:"pageId"`
	PageName              string                              `json:"pageName"`
	PageObjectCoordinates string                              `json:"pageObjectCoordinates"`
	PageURL               string                              `json:"pageUrl"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ContactID                    respjson.Field
		MediaBridgeID                respjson.Field
		MediaBridgeObjectCoordinates respjson.Field
		MediaBridgeObjectTypeID      respjson.Field
		MediaName                    respjson.Field
		MediaType                    respjson.Field
		OccurredTimestamp            respjson.Field
		PortalID                     respjson.Field
		ProviderID                   respjson.Field
		SessionID                    respjson.Field
		State                        respjson.Field
		ExternalPlayContext          respjson.Field
		IframeURL                    respjson.Field
		MediaURL                     respjson.Field
		PageID                       respjson.Field
		PageName                     respjson.Field
		PageObjectCoordinates        respjson.Field
		PageURL                      respjson.Field
		ExtraFields                  map[string]respjson.Field
		raw                          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MediaPlayedEvent) RawJSON() string { return r.JSON.raw }
func (r *MediaPlayedEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaPlayedEventMediaType string

const (
	MediaPlayedEventMediaTypeAudio    MediaPlayedEventMediaType = "AUDIO"
	MediaPlayedEventMediaTypeDocument MediaPlayedEventMediaType = "DOCUMENT"
	MediaPlayedEventMediaTypeImage    MediaPlayedEventMediaType = "IMAGE"
	MediaPlayedEventMediaTypeOther    MediaPlayedEventMediaType = "OTHER"
	MediaPlayedEventMediaTypeVideo    MediaPlayedEventMediaType = "VIDEO"
)

type MediaPlayedEventState string

const (
	MediaPlayedEventStateStarted MediaPlayedEventState = "STARTED"
	MediaPlayedEventStateViewed  MediaPlayedEventState = "VIEWED"
)

type MediaPlayedEventExternalPlayContext string

const (
	MediaPlayedEventExternalPlayContextEmail        MediaPlayedEventExternalPlayContext = "EMAIL"
	MediaPlayedEventExternalPlayContextExternalPage MediaPlayedEventExternalPlayContext = "EXTERNAL_PAGE"
)

// The properties MediaType, OccurredTimestamp, SessionID, State are required.
type MediaPlayedEventRequestParam struct {
	// Any of "AUDIO", "DOCUMENT", "IMAGE", "OTHER", "VIDEO".
	MediaType         MediaPlayedEventRequestMediaType `json:"mediaType,omitzero" api:"required"`
	OccurredTimestamp int64                            `json:"occurredTimestamp" api:"required"`
	SessionID         string                           `json:"sessionId" api:"required"`
	// Any of "STARTED", "VIEWED".
	State         MediaPlayedEventRequestState `json:"state,omitzero" api:"required"`
	Hsenc         param.Opt[string]            `json:"_hsenc,omitzero"`
	ContactID     param.Opt[int64]             `json:"contactId,omitzero"`
	ContactUtk    param.Opt[string]            `json:"contactUtk,omitzero"`
	ExternalID    param.Opt[string]            `json:"externalId,omitzero"`
	IframeURL     param.Opt[string]            `json:"iframeUrl,omitzero"`
	MediaBridgeID param.Opt[int64]             `json:"mediaBridgeId,omitzero"`
	MediaName     param.Opt[string]            `json:"mediaName,omitzero"`
	MediaURL      param.Opt[string]            `json:"mediaUrl,omitzero"`
	PageID        param.Opt[int64]             `json:"pageId,omitzero"`
	PageName      param.Opt[string]            `json:"pageName,omitzero"`
	PageURL       param.Opt[string]            `json:"pageUrl,omitzero"`
	// Any of "EMAIL", "EXTERNAL_PAGE".
	ExternalPlayContext MediaPlayedEventRequestExternalPlayContext `json:"externalPlayContext,omitzero"`
	paramObj
}

func (r MediaPlayedEventRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow MediaPlayedEventRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MediaPlayedEventRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaPlayedEventRequestMediaType string

const (
	MediaPlayedEventRequestMediaTypeAudio    MediaPlayedEventRequestMediaType = "AUDIO"
	MediaPlayedEventRequestMediaTypeDocument MediaPlayedEventRequestMediaType = "DOCUMENT"
	MediaPlayedEventRequestMediaTypeImage    MediaPlayedEventRequestMediaType = "IMAGE"
	MediaPlayedEventRequestMediaTypeOther    MediaPlayedEventRequestMediaType = "OTHER"
	MediaPlayedEventRequestMediaTypeVideo    MediaPlayedEventRequestMediaType = "VIDEO"
)

type MediaPlayedEventRequestState string

const (
	MediaPlayedEventRequestStateStarted MediaPlayedEventRequestState = "STARTED"
	MediaPlayedEventRequestStateViewed  MediaPlayedEventRequestState = "VIEWED"
)

type MediaPlayedEventRequestExternalPlayContext string

const (
	MediaPlayedEventRequestExternalPlayContextEmail        MediaPlayedEventRequestExternalPlayContext = "EMAIL"
	MediaPlayedEventRequestExternalPlayContextExternalPage MediaPlayedEventRequestExternalPlayContext = "EXTERNAL_PAGE"
)

type MediaPlayedPercentageEvent struct {
	// The ID of the contact in HubSpot’s system that consumed the media. This can be
	// fetched using HubSpot's Get contact by usertoken (utk) API. The API also
	// supports supplying a usertoken, and will handle converting this into a contact
	// ID automatically.
	ContactID                    int64  `json:"contactId" api:"required"`
	MediaBridgeID                int64  `json:"mediaBridgeId" api:"required"`
	MediaBridgeObjectCoordinates string `json:"mediaBridgeObjectCoordinates" api:"required"`
	MediaBridgeObjectTypeID      string `json:"mediaBridgeObjectTypeId" api:"required"`
	MediaName                    string `json:"mediaName" api:"required"`
	// Any of "AUDIO", "DOCUMENT", "IMAGE", "OTHER", "VIDEO".
	MediaType         MediaPlayedPercentageEventMediaType `json:"mediaType" api:"required"`
	OccurredTimestamp int64                               `json:"occurredTimestamp" api:"required"`
	PlayedPercent     int64                               `json:"playedPercent" api:"required"`
	// The ID of the HubSpot account.
	PortalID   int64  `json:"portalId" api:"required"`
	ProviderID int64  `json:"providerId" api:"required"`
	SessionID  string `json:"sessionId" api:"required"`
	// Any of "EMAIL", "EXTERNAL_PAGE".
	ExternalPlayContext MediaPlayedPercentageEventExternalPlayContext `json:"externalPlayContext"`
	MediaURL            string                                        `json:"mediaUrl"`
	// The content ID of the page that an event happened on, for HubSpot pages.
	// Required if the page is a HubSpot page.
	PageID int64 `json:"pageId"`
	// The name or title of the page that an event happened on. Required for
	// non-HubSpot pages.
	PageName              string `json:"pageName"`
	PageObjectCoordinates string `json:"pageObjectCoordinates"`
	// The URL of the page that an event happened on. Required for non-HubSpot pages.
	PageURL string `json:"pageUrl"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ContactID                    respjson.Field
		MediaBridgeID                respjson.Field
		MediaBridgeObjectCoordinates respjson.Field
		MediaBridgeObjectTypeID      respjson.Field
		MediaName                    respjson.Field
		MediaType                    respjson.Field
		OccurredTimestamp            respjson.Field
		PlayedPercent                respjson.Field
		PortalID                     respjson.Field
		ProviderID                   respjson.Field
		SessionID                    respjson.Field
		ExternalPlayContext          respjson.Field
		MediaURL                     respjson.Field
		PageID                       respjson.Field
		PageName                     respjson.Field
		PageObjectCoordinates        respjson.Field
		PageURL                      respjson.Field
		ExtraFields                  map[string]respjson.Field
		raw                          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MediaPlayedPercentageEvent) RawJSON() string { return r.JSON.raw }
func (r *MediaPlayedPercentageEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaPlayedPercentageEventMediaType string

const (
	MediaPlayedPercentageEventMediaTypeAudio    MediaPlayedPercentageEventMediaType = "AUDIO"
	MediaPlayedPercentageEventMediaTypeDocument MediaPlayedPercentageEventMediaType = "DOCUMENT"
	MediaPlayedPercentageEventMediaTypeImage    MediaPlayedPercentageEventMediaType = "IMAGE"
	MediaPlayedPercentageEventMediaTypeOther    MediaPlayedPercentageEventMediaType = "OTHER"
	MediaPlayedPercentageEventMediaTypeVideo    MediaPlayedPercentageEventMediaType = "VIDEO"
)

type MediaPlayedPercentageEventExternalPlayContext string

const (
	MediaPlayedPercentageEventExternalPlayContextEmail        MediaPlayedPercentageEventExternalPlayContext = "EMAIL"
	MediaPlayedPercentageEventExternalPlayContextExternalPage MediaPlayedPercentageEventExternalPlayContext = "EXTERNAL_PAGE"
)

// The properties MediaType, OccurredTimestamp, PlayedPercent, SessionID are
// required.
type MediaPlayedPercentageEventRequestParam struct {
	// Any of "AUDIO", "DOCUMENT", "IMAGE", "OTHER", "VIDEO".
	MediaType         MediaPlayedPercentageEventRequestMediaType `json:"mediaType,omitzero" api:"required"`
	OccurredTimestamp int64                                      `json:"occurredTimestamp" api:"required"`
	PlayedPercent     int64                                      `json:"playedPercent" api:"required"`
	SessionID         string                                     `json:"sessionId" api:"required"`
	Hsenc             param.Opt[string]                          `json:"_hsenc,omitzero"`
	ContactID         param.Opt[int64]                           `json:"contactId,omitzero"`
	ContactUtk        param.Opt[string]                          `json:"contactUtk,omitzero"`
	ExternalID        param.Opt[string]                          `json:"externalId,omitzero"`
	MediaBridgeID     param.Opt[int64]                           `json:"mediaBridgeId,omitzero"`
	MediaName         param.Opt[string]                          `json:"mediaName,omitzero"`
	MediaURL          param.Opt[string]                          `json:"mediaUrl,omitzero"`
	PageID            param.Opt[int64]                           `json:"pageId,omitzero"`
	PageName          param.Opt[string]                          `json:"pageName,omitzero"`
	PageURL           param.Opt[string]                          `json:"pageUrl,omitzero"`
	// Any of "EMAIL", "EXTERNAL_PAGE".
	ExternalPlayContext MediaPlayedPercentageEventRequestExternalPlayContext `json:"externalPlayContext,omitzero"`
	paramObj
}

func (r MediaPlayedPercentageEventRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow MediaPlayedPercentageEventRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MediaPlayedPercentageEventRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaPlayedPercentageEventRequestMediaType string

const (
	MediaPlayedPercentageEventRequestMediaTypeAudio    MediaPlayedPercentageEventRequestMediaType = "AUDIO"
	MediaPlayedPercentageEventRequestMediaTypeDocument MediaPlayedPercentageEventRequestMediaType = "DOCUMENT"
	MediaPlayedPercentageEventRequestMediaTypeImage    MediaPlayedPercentageEventRequestMediaType = "IMAGE"
	MediaPlayedPercentageEventRequestMediaTypeOther    MediaPlayedPercentageEventRequestMediaType = "OTHER"
	MediaPlayedPercentageEventRequestMediaTypeVideo    MediaPlayedPercentageEventRequestMediaType = "VIDEO"
)

type MediaPlayedPercentageEventRequestExternalPlayContext string

const (
	MediaPlayedPercentageEventRequestExternalPlayContextEmail        MediaPlayedPercentageEventRequestExternalPlayContext = "EMAIL"
	MediaPlayedPercentageEventRequestExternalPlayContextExternalPage MediaPlayedPercentageEventRequestExternalPlayContext = "EXTERNAL_PAGE"
)

type OEmbedDomainsCollectionResponse struct {
	Results    []IntegratorOEmbedDomainModel `json:"results" api:"required"`
	TotalCount int64                         `json:"totalCount"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Results     respjson.Field
		TotalCount  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OEmbedDomainsCollectionResponse) RawJSON() string { return r.JSON.raw }
func (r *OEmbedDomainsCollectionResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ObjectDefinitionResponse struct {
	ObjectTypeID   string               `json:"objectTypeId" api:"required"`
	ObjectTypeName string               `json:"objectTypeName" api:"required"`
	Properties     []PropertyDefinition `json:"properties" api:"required"`
	PropertyGroups []GroupView          `json:"propertyGroups" api:"required"`
	Schema         InboundDBObjectType  `json:"schema"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ObjectTypeID   respjson.Field
		ObjectTypeName respjson.Field
		Properties     respjson.Field
		PropertyGroups respjson.Field
		Schema         respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ObjectDefinitionResponse) RawJSON() string { return r.JSON.raw }
func (r *ObjectDefinitionResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ObjectSchema struct {
	ID                         string                             `json:"id" api:"required"`
	AllowsSensitiveProperties  bool                               `json:"allowsSensitiveProperties" api:"required"`
	Archived                   bool                               `json:"archived" api:"required"`
	Associations               []shared.BaseAssociationDefinition `json:"associations" api:"required"`
	FullyQualifiedName         string                             `json:"fullyQualifiedName" api:"required"`
	Labels                     shared.ObjectTypeDefinitionLabels  `json:"labels" api:"required"`
	Name                       string                             `json:"name" api:"required"`
	ObjectTypeID               string                             `json:"objectTypeId" api:"required"`
	Properties                 []Property1                        `json:"properties" api:"required"`
	RequiredProperties         []string                           `json:"requiredProperties" api:"required"`
	SearchableProperties       []string                           `json:"searchableProperties" api:"required"`
	SecondaryDisplayProperties []string                           `json:"secondaryDisplayProperties" api:"required"`
	CreatedAt                  time.Time                          `json:"createdAt" format:"date-time"`
	CreatedByUserID            int64                              `json:"createdByUserId"`
	Description                string                             `json:"description"`
	PrimaryDisplayProperty     string                             `json:"primaryDisplayProperty"`
	UpdatedAt                  time.Time                          `json:"updatedAt" format:"date-time"`
	UpdatedByUserID            int64                              `json:"updatedByUserId"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                         respjson.Field
		AllowsSensitiveProperties  respjson.Field
		Archived                   respjson.Field
		Associations               respjson.Field
		FullyQualifiedName         respjson.Field
		Labels                     respjson.Field
		Name                       respjson.Field
		ObjectTypeID               respjson.Field
		Properties                 respjson.Field
		RequiredProperties         respjson.Field
		SearchableProperties       respjson.Field
		SecondaryDisplayProperties respjson.Field
		CreatedAt                  respjson.Field
		CreatedByUserID            respjson.Field
		Description                respjson.Field
		PrimaryDisplayProperty     respjson.Field
		UpdatedAt                  respjson.Field
		UpdatedByUserID            respjson.Field
		ExtraFields                map[string]respjson.Field
		raw                        string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ObjectSchema) RawJSON() string { return r.JSON.raw }
func (r *ObjectSchema) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ObjectTypeIDProto struct {
	InnerID    int64 `json:"innerId" api:"required"`
	MetaTypeID int64 `json:"metaTypeId" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		InnerID     respjson.Field
		MetaTypeID  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ObjectTypeIDProto) RawJSON() string { return r.JSON.raw }
func (r *ObjectTypeIDProto) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Option1 struct {
	Hidden       bool   `json:"hidden" api:"required"`
	Label        string `json:"label" api:"required"`
	Value        string `json:"value" api:"required"`
	Description  string `json:"description"`
	DisplayOrder int64  `json:"displayOrder"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Hidden       respjson.Field
		Label        respjson.Field
		Value        respjson.Field
		Description  respjson.Field
		DisplayOrder respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Option1) RawJSON() string { return r.JSON.raw }
func (r *Option1) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OptionDecorations struct {
	// Any of "BLACK", "BLUE", "BLUE_LIGHT", "GRAY", "GREEN", "GREEN_LIGHT", "ORANGE",
	// "ORANGE_LIGHT", "PINK", "PINK_LIGHT", "PURPLE", "PURPLE_LIGHT", "RED",
	// "RED_LIGHT", "TEAL", "TEAL_LIGHT", "YELLOW", "YELLOW_LIGHT".
	Color OptionDecorationsColor `json:"color" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Color       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OptionDecorations) RawJSON() string { return r.JSON.raw }
func (r *OptionDecorations) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OptionDecorationsColor string

const (
	OptionDecorationsColorBlack       OptionDecorationsColor = "BLACK"
	OptionDecorationsColorBlue        OptionDecorationsColor = "BLUE"
	OptionDecorationsColorBlueLight   OptionDecorationsColor = "BLUE_LIGHT"
	OptionDecorationsColorGray        OptionDecorationsColor = "GRAY"
	OptionDecorationsColorGreen       OptionDecorationsColor = "GREEN"
	OptionDecorationsColorGreenLight  OptionDecorationsColor = "GREEN_LIGHT"
	OptionDecorationsColorOrange      OptionDecorationsColor = "ORANGE"
	OptionDecorationsColorOrangeLight OptionDecorationsColor = "ORANGE_LIGHT"
	OptionDecorationsColorPink        OptionDecorationsColor = "PINK"
	OptionDecorationsColorPinkLight   OptionDecorationsColor = "PINK_LIGHT"
	OptionDecorationsColorPurple      OptionDecorationsColor = "PURPLE"
	OptionDecorationsColorPurpleLight OptionDecorationsColor = "PURPLE_LIGHT"
	OptionDecorationsColorRed         OptionDecorationsColor = "RED"
	OptionDecorationsColorRedLight    OptionDecorationsColor = "RED_LIGHT"
	OptionDecorationsColorTeal        OptionDecorationsColor = "TEAL"
	OptionDecorationsColorTealLight   OptionDecorationsColor = "TEAL_LIGHT"
	OptionDecorationsColorYellow      OptionDecorationsColor = "YELLOW"
	OptionDecorationsColorYellowLight OptionDecorationsColor = "YELLOW_LIGHT"
)

type OptionDecoratorsExtensionData struct {
	OptionDecorators map[string]OptionDecorations `json:"optionDecorators" api:"required"`
	// Any of "LABEL_ONLY", "LABEL_WITH_BADGE", "LABEL_WITH_COLOR".
	OptionDecoratorStyle OptionDecoratorsExtensionDataOptionDecoratorStyle `json:"optionDecoratorStyle" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		OptionDecorators     respjson.Field
		OptionDecoratorStyle respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OptionDecoratorsExtensionData) RawJSON() string { return r.JSON.raw }
func (r *OptionDecoratorsExtensionData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OptionDecoratorsExtensionDataOptionDecoratorStyle string

const (
	OptionDecoratorsExtensionDataOptionDecoratorStyleLabelOnly      OptionDecoratorsExtensionDataOptionDecoratorStyle = "LABEL_ONLY"
	OptionDecoratorsExtensionDataOptionDecoratorStyleLabelWithBadge OptionDecoratorsExtensionDataOptionDecoratorStyle = "LABEL_WITH_BADGE"
	OptionDecoratorsExtensionDataOptionDecoratorStyleLabelWithColor OptionDecoratorsExtensionDataOptionDecoratorStyle = "LABEL_WITH_COLOR"
)

type Property1 struct {
	Description          string    `json:"description" api:"required"`
	FieldType            string    `json:"fieldType" api:"required"`
	GroupName            string    `json:"groupName" api:"required"`
	Label                string    `json:"label" api:"required"`
	Name                 string    `json:"name" api:"required"`
	Options              []Option1 `json:"options" api:"required"`
	Type                 string    `json:"type" api:"required"`
	Archived             bool      `json:"archived"`
	ArchivedAt           time.Time `json:"archivedAt" format:"date-time"`
	Calculated           bool      `json:"calculated"`
	CalculationFormula   string    `json:"calculationFormula"`
	CreatedAt            time.Time `json:"createdAt" format:"date-time"`
	CreatedUserID        string    `json:"createdUserId"`
	CurrencyPropertyName string    `json:"currencyPropertyName"`
	// Any of "highly_sensitive", "non_sensitive", "sensitive".
	DataSensitivity Property1DataSensitivity `json:"dataSensitivity"`
	// Any of "absolute", "absolute_with_relative", "time_since", "time_until".
	DateDisplayHint      Property1DateDisplayHint            `json:"dateDisplayHint"`
	DisplayOrder         int64                               `json:"displayOrder"`
	ExternalOptions      bool                                `json:"externalOptions"`
	FormField            bool                                `json:"formField"`
	HasUniqueValue       bool                                `json:"hasUniqueValue"`
	Hidden               bool                                `json:"hidden"`
	HubSpotDefined       bool                                `json:"hubspotDefined"`
	ModificationMetadata shared.PropertyModificationMetadata `json:"modificationMetadata"`
	// Any of "currency", "duration", "formatted", "percentage", "probability",
	// "unformatted".
	NumberDisplayHint       Property1NumberDisplayHint `json:"numberDisplayHint"`
	ReferencedObjectType    string                     `json:"referencedObjectType"`
	SensitiveDataCategories []string                   `json:"sensitiveDataCategories"`
	ShowCurrencySymbol      bool                       `json:"showCurrencySymbol"`
	UpdatedAt               time.Time                  `json:"updatedAt" format:"date-time"`
	UpdatedUserID           string                     `json:"updatedUserId"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Description             respjson.Field
		FieldType               respjson.Field
		GroupName               respjson.Field
		Label                   respjson.Field
		Name                    respjson.Field
		Options                 respjson.Field
		Type                    respjson.Field
		Archived                respjson.Field
		ArchivedAt              respjson.Field
		Calculated              respjson.Field
		CalculationFormula      respjson.Field
		CreatedAt               respjson.Field
		CreatedUserID           respjson.Field
		CurrencyPropertyName    respjson.Field
		DataSensitivity         respjson.Field
		DateDisplayHint         respjson.Field
		DisplayOrder            respjson.Field
		ExternalOptions         respjson.Field
		FormField               respjson.Field
		HasUniqueValue          respjson.Field
		Hidden                  respjson.Field
		HubSpotDefined          respjson.Field
		ModificationMetadata    respjson.Field
		NumberDisplayHint       respjson.Field
		ReferencedObjectType    respjson.Field
		SensitiveDataCategories respjson.Field
		ShowCurrencySymbol      respjson.Field
		UpdatedAt               respjson.Field
		UpdatedUserID           respjson.Field
		ExtraFields             map[string]respjson.Field
		raw                     string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Property1) RawJSON() string { return r.JSON.raw }
func (r *Property1) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Property1DataSensitivity string

const (
	Property1DataSensitivityHighlySensitive Property1DataSensitivity = "highly_sensitive"
	Property1DataSensitivityNonSensitive    Property1DataSensitivity = "non_sensitive"
	Property1DataSensitivitySensitive       Property1DataSensitivity = "sensitive"
)

type Property1DateDisplayHint string

const (
	Property1DateDisplayHintAbsolute             Property1DateDisplayHint = "absolute"
	Property1DateDisplayHintAbsoluteWithRelative Property1DateDisplayHint = "absolute_with_relative"
	Property1DateDisplayHintTimeSince            Property1DateDisplayHint = "time_since"
	Property1DateDisplayHintTimeUntil            Property1DateDisplayHint = "time_until"
)

type Property1NumberDisplayHint string

const (
	Property1NumberDisplayHintCurrency    Property1NumberDisplayHint = "currency"
	Property1NumberDisplayHintDuration    Property1NumberDisplayHint = "duration"
	Property1NumberDisplayHintFormatted   Property1NumberDisplayHint = "formatted"
	Property1NumberDisplayHintPercentage  Property1NumberDisplayHint = "percentage"
	Property1NumberDisplayHintProbability Property1NumberDisplayHint = "probability"
	Property1NumberDisplayHintUnformatted Property1NumberDisplayHint = "unformatted"
)

type PropertyDefinition struct {
	ObjectTypeID string `json:"objectTypeId" api:"required"`
	// A HubSpot property
	Property                 MediaBridgeProperty      `json:"property" api:"required"`
	CalculationExpression    any                      `json:"calculationExpression"`
	CalculationFormula       string                   `json:"calculationFormula"`
	DefinitionSource         PropertyDefinitionSource `json:"definitionSource"`
	ExtensionData            ExtensionData            `json:"extensionData"`
	ExternalOptionsMetaData  ExternalOptionsMetaData  `json:"externalOptionsMetaData"`
	FulcrumPortalID          int64                    `json:"fulcrumPortalId"`
	FulcrumTimestamp         int64                    `json:"fulcrumTimestamp"`
	JanusGroup               string                   `json:"janusGroup"`
	LookupAssociationSpec    LookupAssociationSpec    `json:"lookupAssociationSpec"`
	Permission               FieldLevelPermission     `json:"permission"`
	PropertyDefinitionSource DefinitionSource         `json:"propertyDefinitionSource"`
	PropertyRequirements     DefaultRequirements      `json:"propertyRequirements"`
	RollupExpression         RollupExpression         `json:"rollupExpression"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ObjectTypeID             respjson.Field
		Property                 respjson.Field
		CalculationExpression    respjson.Field
		CalculationFormula       respjson.Field
		DefinitionSource         respjson.Field
		ExtensionData            respjson.Field
		ExternalOptionsMetaData  respjson.Field
		FulcrumPortalID          respjson.Field
		FulcrumTimestamp         respjson.Field
		JanusGroup               respjson.Field
		LookupAssociationSpec    respjson.Field
		Permission               respjson.Field
		PropertyDefinitionSource respjson.Field
		PropertyRequirements     respjson.Field
		RollupExpression         respjson.Field
		ExtraFields              map[string]respjson.Field
		raw                      string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PropertyDefinition) RawJSON() string { return r.JSON.raw }
func (r *PropertyDefinition) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PropertyDefinitionSource struct {
	// Any of "GLOBAL", "HAVEN_BRANCH", "OBJECT_TYPE", "PORTAL".
	Type PropertyDefinitionSourceType `json:"type" api:"required"`
	Name string                       `json:"name"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PropertyDefinitionSource) RawJSON() string { return r.JSON.raw }
func (r *PropertyDefinitionSource) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PropertyDefinitionSourceType string

const (
	PropertyDefinitionSourceTypeGlobal      PropertyDefinitionSourceType = "GLOBAL"
	PropertyDefinitionSourceTypeHavenBranch PropertyDefinitionSourceType = "HAVEN_BRANCH"
	PropertyDefinitionSourceTypeObjectType  PropertyDefinitionSourceType = "OBJECT_TYPE"
	PropertyDefinitionSourceTypePortal      PropertyDefinitionSourceType = "PORTAL"
)

type RequiredPropertiesExtensionData struct {
	IsRequiredProperty bool `json:"isRequiredProperty" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		IsRequiredProperty respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RequiredPropertiesExtensionData) RawJSON() string { return r.JSON.raw }
func (r *RequiredPropertiesExtensionData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RollupExpression struct {
	AssociationTypes []shared.AssociationSpec `json:"associationTypes" api:"required"`
	// Any of "AVERAGE", "COUNT", "EARLIEST_VALUE", "LATEST_VALUE", "MAX", "MAX_BY",
	// "MIN", "MIN_BY", "REFERENCED_ID_SET", "REFERENCED_STRING_SET",
	// "REFERENCED_STRING_SET_INTERSECTION", "SUM", "SYNC_MAX_BY", "SYNC_MIN_BY",
	// "SYNC_VALUE", "UNKNOWN_ROLLUP_OPERATOR".
	RollupOperator              RollupExpressionRollupOperator `json:"rollupOperator" api:"required"`
	SourceObjectTypeID          string                         `json:"sourceObjectTypeId" api:"required"`
	SourcePropertyName          string                         `json:"sourcePropertyName" api:"required"`
	ConditionalExpression       any                            `json:"conditionalExpression"`
	ConditionalFormula          string                         `json:"conditionalFormula"`
	EmptyRollupValue            string                         `json:"emptyRollupValue"`
	SourceCompareByPropertyName string                         `json:"sourceCompareByPropertyName"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AssociationTypes            respjson.Field
		RollupOperator              respjson.Field
		SourceObjectTypeID          respjson.Field
		SourcePropertyName          respjson.Field
		ConditionalExpression       respjson.Field
		ConditionalFormula          respjson.Field
		EmptyRollupValue            respjson.Field
		SourceCompareByPropertyName respjson.Field
		ExtraFields                 map[string]respjson.Field
		raw                         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RollupExpression) RawJSON() string { return r.JSON.raw }
func (r *RollupExpression) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RollupExpressionRollupOperator string

const (
	RollupExpressionRollupOperatorAverage                         RollupExpressionRollupOperator = "AVERAGE"
	RollupExpressionRollupOperatorCount                           RollupExpressionRollupOperator = "COUNT"
	RollupExpressionRollupOperatorEarliestValue                   RollupExpressionRollupOperator = "EARLIEST_VALUE"
	RollupExpressionRollupOperatorLatestValue                     RollupExpressionRollupOperator = "LATEST_VALUE"
	RollupExpressionRollupOperatorMax                             RollupExpressionRollupOperator = "MAX"
	RollupExpressionRollupOperatorMaxBy                           RollupExpressionRollupOperator = "MAX_BY"
	RollupExpressionRollupOperatorMin                             RollupExpressionRollupOperator = "MIN"
	RollupExpressionRollupOperatorMinBy                           RollupExpressionRollupOperator = "MIN_BY"
	RollupExpressionRollupOperatorReferencedIDSet                 RollupExpressionRollupOperator = "REFERENCED_ID_SET"
	RollupExpressionRollupOperatorReferencedStringSet             RollupExpressionRollupOperator = "REFERENCED_STRING_SET"
	RollupExpressionRollupOperatorReferencedStringSetIntersection RollupExpressionRollupOperator = "REFERENCED_STRING_SET_INTERSECTION"
	RollupExpressionRollupOperatorSum                             RollupExpressionRollupOperator = "SUM"
	RollupExpressionRollupOperatorSyncMaxBy                       RollupExpressionRollupOperator = "SYNC_MAX_BY"
	RollupExpressionRollupOperatorSyncMinBy                       RollupExpressionRollupOperator = "SYNC_MIN_BY"
	RollupExpressionRollupOperatorSyncValue                       RollupExpressionRollupOperator = "SYNC_VALUE"
	RollupExpressionRollupOperatorUnknownRollupOperator           RollupExpressionRollupOperator = "UNKNOWN_ROLLUP_OPERATOR"
)

type ScopeMapping struct {
	// Any of "ALL", "OWNED", "TEAM_OWNED", "UNASSIGNED".
	AccessLevel ScopeMappingAccessLevel `json:"accessLevel" api:"required"`
	// Any of "COMMUNICATE", "DELETE", "EDIT", "EDIT_ASSOCIATION", "MERGE", "VIEW".
	RequestAction ScopeMappingRequestAction `json:"requestAction" api:"required"`
	ScopeName     string                    `json:"scopeName" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AccessLevel   respjson.Field
		RequestAction respjson.Field
		ScopeName     respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ScopeMapping) RawJSON() string { return r.JSON.raw }
func (r *ScopeMapping) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ScopeMappingAccessLevel string

const (
	ScopeMappingAccessLevelAll        ScopeMappingAccessLevel = "ALL"
	ScopeMappingAccessLevelOwned      ScopeMappingAccessLevel = "OWNED"
	ScopeMappingAccessLevelTeamOwned  ScopeMappingAccessLevel = "TEAM_OWNED"
	ScopeMappingAccessLevelUnassigned ScopeMappingAccessLevel = "UNASSIGNED"
)

type ScopeMappingRequestAction string

const (
	ScopeMappingRequestActionCommunicate     ScopeMappingRequestAction = "COMMUNICATE"
	ScopeMappingRequestActionDelete          ScopeMappingRequestAction = "DELETE"
	ScopeMappingRequestActionEdit            ScopeMappingRequestAction = "EDIT"
	ScopeMappingRequestActionEditAssociation ScopeMappingRequestAction = "EDIT_ASSOCIATION"
	ScopeMappingRequestActionMerge           ScopeMappingRequestAction = "MERGE"
	ScopeMappingRequestActionView            ScopeMappingRequestAction = "VIEW"
)

type SoftRequiredPropertiesExtensionData struct {
	IsSoftRequiredProperty bool `json:"isSoftRequiredProperty" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		IsSoftRequiredProperty respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SoftRequiredPropertiesExtensionData) RawJSON() string { return r.JSON.raw }
func (r *SoftRequiredPropertiesExtensionData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeNewAssociationParams struct {
	AppID                    int64 `path:"appId" api:"required" json:"-"`
	AssociationDefinitionEgg shared.AssociationDefinitionEggParam
	paramObj
}

func (r MediaBridgeNewAssociationParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.AssociationDefinitionEgg)
}
func (r *MediaBridgeNewAssociationParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeNewAttentionSpanEventParams struct {
	AttentionSpanEventRequest AttentionSpanEventRequestParam
	paramObj
}

func (r MediaBridgeNewAttentionSpanEventParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.AttentionSpanEventRequest)
}
func (r *MediaBridgeNewAttentionSpanEventParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeNewMediaPlayedEventParams struct {
	MediaPlayedEventRequest MediaPlayedEventRequestParam
	paramObj
}

func (r MediaBridgeNewMediaPlayedEventParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.MediaPlayedEventRequest)
}
func (r *MediaBridgeNewMediaPlayedEventParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeNewMediaPlayedPercentEventParams struct {
	MediaPlayedPercentageEventRequest MediaPlayedPercentageEventRequestParam
	paramObj
}

func (r MediaBridgeNewMediaPlayedPercentEventParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.MediaPlayedPercentageEventRequest)
}
func (r *MediaBridgeNewMediaPlayedPercentEventParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeNewObjectTypeParams struct {
	IntegratorObjectCreationRequest IntegratorObjectCreationRequestParam
	paramObj
}

func (r MediaBridgeNewObjectTypeParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.IntegratorObjectCreationRequest)
}
func (r *MediaBridgeNewObjectTypeParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeNewOembedDomainParams struct {
	IntegratorOEmbedDomainRequest IntegratorOEmbedDomainRequestParam
	paramObj
}

func (r MediaBridgeNewOembedDomainParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.IntegratorOEmbedDomainRequest)
}
func (r *MediaBridgeNewOembedDomainParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeNewPropertyParams struct {
	AppID          int64 `path:"appId" api:"required" json:"-"`
	PropertyCreate shared.PropertyCreateParam
	paramObj
}

func (r MediaBridgeNewPropertyParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.PropertyCreate)
}
func (r *MediaBridgeNewPropertyParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeNewPropertyGroupParams struct {
	AppID               int64 `path:"appId" api:"required" json:"-"`
	PropertyGroupCreate shared.PropertyGroupCreateParam
	paramObj
}

func (r MediaBridgeNewPropertyGroupParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.PropertyGroupCreate)
}
func (r *MediaBridgeNewPropertyGroupParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeDeleteAssociationParams struct {
	AppID      int64  `path:"appId" api:"required" json:"-"`
	ObjectType string `path:"objectType" api:"required" json:"-"`
	paramObj
}

type MediaBridgeDeleteOembedDomainParams struct {
	ID             param.Opt[int64] `query:"id,omitzero" json:"-"`
	DomainPortalID param.Opt[int64] `query:"domainPortalId,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [MediaBridgeDeleteOembedDomainParams]'s query parameters as
// `url.Values`.
func (r MediaBridgeDeleteOembedDomainParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type MediaBridgeDeletePropertyParams struct {
	AppID      int64  `path:"appId" api:"required" json:"-"`
	ObjectType string `path:"objectType" api:"required" json:"-"`
	paramObj
}

type MediaBridgeDeletePropertyGroupParams struct {
	AppID      int64  `path:"appId" api:"required" json:"-"`
	ObjectType string `path:"objectType" api:"required" json:"-"`
	paramObj
}

type MediaBridgeGetOembedDomainParams struct {
	AppID int64 `path:"appId" api:"required" json:"-"`
	paramObj
}

type MediaBridgeGetPropertyParams struct {
	AppID      int64  `path:"appId" api:"required" json:"-"`
	ObjectType string `path:"objectType" api:"required" json:"-"`
	// Whether to return only results that have been archived.
	Archived   param.Opt[bool]   `query:"archived,omitzero" json:"-"`
	Properties param.Opt[string] `query:"properties,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [MediaBridgeGetPropertyParams]'s query parameters as
// `url.Values`.
func (r MediaBridgeGetPropertyParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type MediaBridgeGetPropertyGroupParams struct {
	AppID      int64  `path:"appId" api:"required" json:"-"`
	ObjectType string `path:"objectType" api:"required" json:"-"`
	paramObj
}

type MediaBridgeGetSchemaParams struct {
	AppID int64 `path:"appId" api:"required" json:"-"`
	paramObj
}

type MediaBridgeListObjectTypesByMediaTypeParams struct {
	AppID                 int64           `path:"appId" api:"required" json:"-"`
	IncludeFullDefinition param.Opt[bool] `query:"includeFullDefinition,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [MediaBridgeListObjectTypesByMediaTypeParams]'s query
// parameters as `url.Values`.
func (r MediaBridgeListObjectTypesByMediaTypeParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type MediaBridgeListObjectTypesByMediaTypeParamsMediaType string

const (
	MediaBridgeListObjectTypesByMediaTypeParamsMediaTypeAudio    MediaBridgeListObjectTypesByMediaTypeParamsMediaType = "AUDIO"
	MediaBridgeListObjectTypesByMediaTypeParamsMediaTypeDocument MediaBridgeListObjectTypesByMediaTypeParamsMediaType = "DOCUMENT"
	MediaBridgeListObjectTypesByMediaTypeParamsMediaTypeImage    MediaBridgeListObjectTypesByMediaTypeParamsMediaType = "IMAGE"
	MediaBridgeListObjectTypesByMediaTypeParamsMediaTypeOther    MediaBridgeListObjectTypesByMediaTypeParamsMediaType = "OTHER"
	MediaBridgeListObjectTypesByMediaTypeParamsMediaTypeVideo    MediaBridgeListObjectTypesByMediaTypeParamsMediaType = "VIDEO"
)

type MediaBridgeListOembedDomainsParams struct {
	DomainPortalID param.Opt[int64] `query:"domainPortalId,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [MediaBridgeListOembedDomainsParams]'s query parameters as
// `url.Values`.
func (r MediaBridgeListOembedDomainsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type MediaBridgeListPropertiesParams struct {
	AppID int64 `path:"appId" api:"required" json:"-"`
	// Whether to return only results that have been archived.
	Archived   param.Opt[bool]   `query:"archived,omitzero" json:"-"`
	Properties param.Opt[string] `query:"properties,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [MediaBridgeListPropertiesParams]'s query parameters as
// `url.Values`.
func (r MediaBridgeListPropertiesParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type MediaBridgeListPropertyGroupsParams struct {
	AppID int64 `path:"appId" api:"required" json:"-"`
	paramObj
}

type MediaBridgeListSchemasParams struct {
	// Whether to return only results that have been archived.
	Archived param.Opt[bool] `query:"archived,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [MediaBridgeListSchemasParams]'s query parameters as
// `url.Values`.
func (r MediaBridgeListSchemasParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type MediaBridgeRegisterAppNameParams struct {
	MediaBridgeProviderPartial MediaBridgeProviderPartialParam
	paramObj
}

func (r MediaBridgeRegisterAppNameParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.MediaBridgeProviderPartial)
}
func (r *MediaBridgeRegisterAppNameParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeUpdateEventVisibilitySettingsParams struct {
	EventVisibilityChange EventVisibilityChangeParam
	paramObj
}

func (r MediaBridgeUpdateEventVisibilitySettingsParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.EventVisibilityChange)
}
func (r *MediaBridgeUpdateEventVisibilitySettingsParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeUpdateOembedDomainParams struct {
	AppID                         int64 `path:"appId" api:"required" json:"-"`
	IntegratorOEmbedDomainRequest IntegratorOEmbedDomainRequestParam
	paramObj
}

func (r MediaBridgeUpdateOembedDomainParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.IntegratorOEmbedDomainRequest)
}
func (r *MediaBridgeUpdateOembedDomainParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeUpdatePropertyParams struct {
	AppID                     int64  `path:"appId" api:"required" json:"-"`
	ObjectType                string `path:"objectType" api:"required" json:"-"`
	MediaBridgePropertyUpdate MediaBridgePropertyUpdateParam
	paramObj
}

func (r MediaBridgeUpdatePropertyParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.MediaBridgePropertyUpdate)
}
func (r *MediaBridgeUpdatePropertyParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeUpdatePropertyGroupParams struct {
	AppID               int64  `path:"appId" api:"required" json:"-"`
	ObjectType          string `path:"objectType" api:"required" json:"-"`
	PropertyGroupUpdate shared.PropertyGroupUpdateParam
	paramObj
}

func (r MediaBridgeUpdatePropertyGroupParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.PropertyGroupUpdate)
}
func (r *MediaBridgeUpdatePropertyGroupParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeUpdateSchemaParams struct {
	AppID                     int64 `path:"appId" api:"required" json:"-"`
	ObjectTypeDefinitionPatch shared.ObjectTypeDefinitionPatchParam
	paramObj
}

func (r MediaBridgeUpdateSchemaParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.ObjectTypeDefinitionPatch)
}
func (r *MediaBridgeUpdateSchemaParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeUpdateSettingsParams struct {
	MediaBridgeProviderPartial MediaBridgeProviderPartialParam
	paramObj
}

func (r MediaBridgeUpdateSettingsParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.MediaBridgeProviderPartial)
}
func (r *MediaBridgeUpdateSettingsParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
