// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cms

import (
	"context"
	"net/http"
	"slices"

	"github.com/stainless-sdks/hubspot-sdk-go/internal/apijson"
	"github.com/stainless-sdks/hubspot-sdk-go/internal/requestconfig"
	"github.com/stainless-sdks/hubspot-sdk-go/option"
	"github.com/stainless-sdks/hubspot-sdk-go/packages/param"
	"github.com/stainless-sdks/hubspot-sdk-go/packages/respjson"
)

// MediaBridgeEventService contains methods and other services that help with
// interacting with the Hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMediaBridgeEventService] method instead.
type MediaBridgeEventService struct {
	Options []option.RequestOption
}

// NewMediaBridgeEventService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewMediaBridgeEventService(opts ...option.RequestOption) (r MediaBridgeEventService) {
	r = MediaBridgeEventService{}
	r.Options = opts
	return
}

// Create an event containing the viewers attention span details for the media.
func (r *MediaBridgeEventService) NewAttentionSpanEvent(ctx context.Context, body MediaBridgeEventNewAttentionSpanEventParams, opts ...option.RequestOption) (res *MediaBridgeEventNewAttentionSpanEventResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "media-bridge/v1/events/attention-span"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Create an event for when a user begins playing a piece of media.
func (r *MediaBridgeEventService) NewMediaPlayedEvent(ctx context.Context, body MediaBridgeEventNewMediaPlayedEventParams, opts ...option.RequestOption) (res *MediaBridgeEventNewMediaPlayedEventResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "media-bridge/v1/events/media-played"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Create an event representing a user reaching quarterly milestones in a piece of
// media they're viewing.
func (r *MediaBridgeEventService) NewMediaPlayedPercentEvent(ctx context.Context, body MediaBridgeEventNewMediaPlayedPercentEventParams, opts ...option.RequestOption) (res *MediaBridgeEventNewMediaPlayedPercentEventResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "media-bridge/v1/events/media-played-percent"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type MediaBridgeEventNewAttentionSpanEventResponse struct {
	// The ID of the contact in HubSpot’s system that consumed the media. This can be
	// fetched using HubSpot's Get contact by usertoken (utk) API. The API also
	// supports supplying a usertoken, and will handle converting this into a contact
	// ID automatically.
	ContactID                    int64  `json:"contactId,required"`
	MediaBridgeID                int64  `json:"mediaBridgeId,required"`
	MediaBridgeObjectCoordinates string `json:"mediaBridgeObjectCoordinates,required"`
	MediaBridgeObjectTypeID      string `json:"mediaBridgeObjectTypeId,required"`
	MediaName                    string `json:"mediaName,required"`
	// Any of "VIDEO", "AUDIO", "DOCUMENT", "OTHER", "IMAGE".
	MediaType MediaBridgeEventNewAttentionSpanEventResponseMediaType `json:"mediaType,required"`
	// The timestamp at which this event occurred, in milliseconds since the epoch.
	OccurredTimestamp int64  `json:"occurredTimestamp,required"`
	PercentRange      string `json:"percentRange,required"`
	// The ID of the HubSpot account.
	PortalID   int64  `json:"portalId,required"`
	ProviderID int64  `json:"providerId,required"`
	SessionID  string `json:"sessionId,required"`
	// The percent of the media that the user consumed. Providers may calculate this
	// differently depending on how they consider repeated views of the same portion of
	// media. For this reason, the API will not attempt to validate totalPercentWatched
	// against the attention span information for the event. If it is missing, HubSpot
	// will calculate this from the attention span map as follows: (number of spans
	// with a value of 1 or more)/(Total number of spans).
	TotalPercentPlayed float64 `json:"totalPercentPlayed,required"`
	MediaURL           string  `json:"mediaUrl"`
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
func (r MediaBridgeEventNewAttentionSpanEventResponse) RawJSON() string { return r.JSON.raw }
func (r *MediaBridgeEventNewAttentionSpanEventResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeEventNewAttentionSpanEventResponseMediaType string

const (
	MediaBridgeEventNewAttentionSpanEventResponseMediaTypeVideo    MediaBridgeEventNewAttentionSpanEventResponseMediaType = "VIDEO"
	MediaBridgeEventNewAttentionSpanEventResponseMediaTypeAudio    MediaBridgeEventNewAttentionSpanEventResponseMediaType = "AUDIO"
	MediaBridgeEventNewAttentionSpanEventResponseMediaTypeDocument MediaBridgeEventNewAttentionSpanEventResponseMediaType = "DOCUMENT"
	MediaBridgeEventNewAttentionSpanEventResponseMediaTypeOther    MediaBridgeEventNewAttentionSpanEventResponseMediaType = "OTHER"
	MediaBridgeEventNewAttentionSpanEventResponseMediaTypeImage    MediaBridgeEventNewAttentionSpanEventResponseMediaType = "IMAGE"
)

type MediaBridgeEventNewMediaPlayedEventResponse struct {
	ContactID                    int64  `json:"contactId,required"`
	MediaBridgeID                int64  `json:"mediaBridgeId,required"`
	MediaBridgeObjectCoordinates string `json:"mediaBridgeObjectCoordinates,required"`
	MediaBridgeObjectTypeID      string `json:"mediaBridgeObjectTypeId,required"`
	MediaName                    string `json:"mediaName,required"`
	// Any of "VIDEO", "AUDIO", "DOCUMENT", "OTHER", "IMAGE".
	MediaType         MediaBridgeEventNewMediaPlayedEventResponseMediaType `json:"mediaType,required"`
	OccurredTimestamp int64                                                `json:"occurredTimestamp,required"`
	PortalID          int64                                                `json:"portalId,required"`
	ProviderID        int64                                                `json:"providerId,required"`
	SessionID         string                                               `json:"sessionId,required"`
	// Any of "STARTED", "VIEWED".
	State                 MediaBridgeEventNewMediaPlayedEventResponseState `json:"state,required"`
	IframeURL             string                                           `json:"iframeUrl"`
	MediaURL              string                                           `json:"mediaUrl"`
	PageID                int64                                            `json:"pageId"`
	PageName              string                                           `json:"pageName"`
	PageObjectCoordinates string                                           `json:"pageObjectCoordinates"`
	PageURL               string                                           `json:"pageUrl"`
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
func (r MediaBridgeEventNewMediaPlayedEventResponse) RawJSON() string { return r.JSON.raw }
func (r *MediaBridgeEventNewMediaPlayedEventResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeEventNewMediaPlayedEventResponseMediaType string

const (
	MediaBridgeEventNewMediaPlayedEventResponseMediaTypeVideo    MediaBridgeEventNewMediaPlayedEventResponseMediaType = "VIDEO"
	MediaBridgeEventNewMediaPlayedEventResponseMediaTypeAudio    MediaBridgeEventNewMediaPlayedEventResponseMediaType = "AUDIO"
	MediaBridgeEventNewMediaPlayedEventResponseMediaTypeDocument MediaBridgeEventNewMediaPlayedEventResponseMediaType = "DOCUMENT"
	MediaBridgeEventNewMediaPlayedEventResponseMediaTypeOther    MediaBridgeEventNewMediaPlayedEventResponseMediaType = "OTHER"
	MediaBridgeEventNewMediaPlayedEventResponseMediaTypeImage    MediaBridgeEventNewMediaPlayedEventResponseMediaType = "IMAGE"
)

type MediaBridgeEventNewMediaPlayedEventResponseState string

const (
	MediaBridgeEventNewMediaPlayedEventResponseStateStarted MediaBridgeEventNewMediaPlayedEventResponseState = "STARTED"
	MediaBridgeEventNewMediaPlayedEventResponseStateViewed  MediaBridgeEventNewMediaPlayedEventResponseState = "VIEWED"
)

type MediaBridgeEventNewMediaPlayedPercentEventResponse struct {
	// The ID of the contact in HubSpot’s system that consumed the media. This can be
	// fetched using HubSpot's Get contact by usertoken (utk) API. The API also
	// supports supplying a usertoken, and will handle converting this into a contact
	// ID automatically.
	ContactID                    int64  `json:"contactId,required"`
	MediaBridgeID                int64  `json:"mediaBridgeId,required"`
	MediaBridgeObjectCoordinates string `json:"mediaBridgeObjectCoordinates,required"`
	MediaBridgeObjectTypeID      string `json:"mediaBridgeObjectTypeId,required"`
	MediaName                    string `json:"mediaName,required"`
	// Any of "VIDEO", "AUDIO", "DOCUMENT", "OTHER", "IMAGE".
	MediaType         MediaBridgeEventNewMediaPlayedPercentEventResponseMediaType `json:"mediaType,required"`
	OccurredTimestamp int64                                                       `json:"occurredTimestamp,required"`
	PlayedPercent     int64                                                       `json:"playedPercent,required"`
	// The ID of the HubSpot account.
	PortalID   int64  `json:"portalId,required"`
	ProviderID int64  `json:"providerId,required"`
	SessionID  string `json:"sessionId,required"`
	MediaURL   string `json:"mediaUrl"`
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
func (r MediaBridgeEventNewMediaPlayedPercentEventResponse) RawJSON() string { return r.JSON.raw }
func (r *MediaBridgeEventNewMediaPlayedPercentEventResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeEventNewMediaPlayedPercentEventResponseMediaType string

const (
	MediaBridgeEventNewMediaPlayedPercentEventResponseMediaTypeVideo    MediaBridgeEventNewMediaPlayedPercentEventResponseMediaType = "VIDEO"
	MediaBridgeEventNewMediaPlayedPercentEventResponseMediaTypeAudio    MediaBridgeEventNewMediaPlayedPercentEventResponseMediaType = "AUDIO"
	MediaBridgeEventNewMediaPlayedPercentEventResponseMediaTypeDocument MediaBridgeEventNewMediaPlayedPercentEventResponseMediaType = "DOCUMENT"
	MediaBridgeEventNewMediaPlayedPercentEventResponseMediaTypeOther    MediaBridgeEventNewMediaPlayedPercentEventResponseMediaType = "OTHER"
	MediaBridgeEventNewMediaPlayedPercentEventResponseMediaTypeImage    MediaBridgeEventNewMediaPlayedPercentEventResponseMediaType = "IMAGE"
)

type MediaBridgeEventNewAttentionSpanEventParams struct {
	// Any of "VIDEO", "AUDIO", "DOCUMENT", "OTHER", "IMAGE".
	MediaType         MediaBridgeEventNewAttentionSpanEventParamsMediaType     `json:"mediaType,omitzero,required"`
	OccurredTimestamp int64                                                    `json:"occurredTimestamp,required"`
	RawDataMap        map[string]int64                                         `json:"rawDataMap,omitzero,required"`
	SessionID         string                                                   `json:"sessionId,required"`
	Hsenc             param.Opt[string]                                        `json:"_hsenc,omitzero"`
	ContactID         param.Opt[int64]                                         `json:"contactId,omitzero"`
	ContactUtk        param.Opt[string]                                        `json:"contactUtk,omitzero"`
	ExternalID        param.Opt[string]                                        `json:"externalId,omitzero"`
	MediaBridgeID     param.Opt[int64]                                         `json:"mediaBridgeId,omitzero"`
	MediaName         param.Opt[string]                                        `json:"mediaName,omitzero"`
	MediaURL          param.Opt[string]                                        `json:"mediaUrl,omitzero"`
	PageID            param.Opt[int64]                                         `json:"pageId,omitzero"`
	PageName          param.Opt[string]                                        `json:"pageName,omitzero"`
	PageURL           param.Opt[string]                                        `json:"pageUrl,omitzero"`
	RawDataString     param.Opt[string]                                        `json:"rawDataString,omitzero"`
	DerivedValues     MediaBridgeEventNewAttentionSpanEventParamsDerivedValues `json:"derivedValues,omitzero"`
	paramObj
}

func (r MediaBridgeEventNewAttentionSpanEventParams) MarshalJSON() (data []byte, err error) {
	type shadow MediaBridgeEventNewAttentionSpanEventParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MediaBridgeEventNewAttentionSpanEventParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeEventNewAttentionSpanEventParamsMediaType string

const (
	MediaBridgeEventNewAttentionSpanEventParamsMediaTypeVideo    MediaBridgeEventNewAttentionSpanEventParamsMediaType = "VIDEO"
	MediaBridgeEventNewAttentionSpanEventParamsMediaTypeAudio    MediaBridgeEventNewAttentionSpanEventParamsMediaType = "AUDIO"
	MediaBridgeEventNewAttentionSpanEventParamsMediaTypeDocument MediaBridgeEventNewAttentionSpanEventParamsMediaType = "DOCUMENT"
	MediaBridgeEventNewAttentionSpanEventParamsMediaTypeOther    MediaBridgeEventNewAttentionSpanEventParamsMediaType = "OTHER"
	MediaBridgeEventNewAttentionSpanEventParamsMediaTypeImage    MediaBridgeEventNewAttentionSpanEventParamsMediaType = "IMAGE"
)

// The properties TotalPercentPlayed, TotalSecondsPlayed are required.
type MediaBridgeEventNewAttentionSpanEventParamsDerivedValues struct {
	TotalPercentPlayed float64 `json:"totalPercentPlayed,required"`
	TotalSecondsPlayed int64   `json:"totalSecondsPlayed,required"`
	paramObj
}

func (r MediaBridgeEventNewAttentionSpanEventParamsDerivedValues) MarshalJSON() (data []byte, err error) {
	type shadow MediaBridgeEventNewAttentionSpanEventParamsDerivedValues
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MediaBridgeEventNewAttentionSpanEventParamsDerivedValues) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeEventNewMediaPlayedEventParams struct {
	// Any of "VIDEO", "AUDIO", "DOCUMENT", "OTHER", "IMAGE".
	MediaType         MediaBridgeEventNewMediaPlayedEventParamsMediaType `json:"mediaType,omitzero,required"`
	OccurredTimestamp int64                                              `json:"occurredTimestamp,required"`
	SessionID         string                                             `json:"sessionId,required"`
	// Any of "STARTED", "VIEWED".
	State         MediaBridgeEventNewMediaPlayedEventParamsState `json:"state,omitzero,required"`
	Hsenc         param.Opt[string]                              `json:"_hsenc,omitzero"`
	ContactID     param.Opt[int64]                               `json:"contactId,omitzero"`
	ContactUtk    param.Opt[string]                              `json:"contactUtk,omitzero"`
	ExternalID    param.Opt[string]                              `json:"externalId,omitzero"`
	IframeURL     param.Opt[string]                              `json:"iframeUrl,omitzero"`
	MediaBridgeID param.Opt[int64]                               `json:"mediaBridgeId,omitzero"`
	MediaName     param.Opt[string]                              `json:"mediaName,omitzero"`
	MediaURL      param.Opt[string]                              `json:"mediaUrl,omitzero"`
	PageID        param.Opt[int64]                               `json:"pageId,omitzero"`
	PageName      param.Opt[string]                              `json:"pageName,omitzero"`
	PageURL       param.Opt[string]                              `json:"pageUrl,omitzero"`
	paramObj
}

func (r MediaBridgeEventNewMediaPlayedEventParams) MarshalJSON() (data []byte, err error) {
	type shadow MediaBridgeEventNewMediaPlayedEventParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MediaBridgeEventNewMediaPlayedEventParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeEventNewMediaPlayedEventParamsMediaType string

const (
	MediaBridgeEventNewMediaPlayedEventParamsMediaTypeVideo    MediaBridgeEventNewMediaPlayedEventParamsMediaType = "VIDEO"
	MediaBridgeEventNewMediaPlayedEventParamsMediaTypeAudio    MediaBridgeEventNewMediaPlayedEventParamsMediaType = "AUDIO"
	MediaBridgeEventNewMediaPlayedEventParamsMediaTypeDocument MediaBridgeEventNewMediaPlayedEventParamsMediaType = "DOCUMENT"
	MediaBridgeEventNewMediaPlayedEventParamsMediaTypeOther    MediaBridgeEventNewMediaPlayedEventParamsMediaType = "OTHER"
	MediaBridgeEventNewMediaPlayedEventParamsMediaTypeImage    MediaBridgeEventNewMediaPlayedEventParamsMediaType = "IMAGE"
)

type MediaBridgeEventNewMediaPlayedEventParamsState string

const (
	MediaBridgeEventNewMediaPlayedEventParamsStateStarted MediaBridgeEventNewMediaPlayedEventParamsState = "STARTED"
	MediaBridgeEventNewMediaPlayedEventParamsStateViewed  MediaBridgeEventNewMediaPlayedEventParamsState = "VIEWED"
)

type MediaBridgeEventNewMediaPlayedPercentEventParams struct {
	// Any of "VIDEO", "AUDIO", "DOCUMENT", "OTHER", "IMAGE".
	MediaType         MediaBridgeEventNewMediaPlayedPercentEventParamsMediaType `json:"mediaType,omitzero,required"`
	OccurredTimestamp int64                                                     `json:"occurredTimestamp,required"`
	PlayedPercent     int64                                                     `json:"playedPercent,required"`
	SessionID         string                                                    `json:"sessionId,required"`
	Hsenc             param.Opt[string]                                         `json:"_hsenc,omitzero"`
	ContactID         param.Opt[int64]                                          `json:"contactId,omitzero"`
	ContactUtk        param.Opt[string]                                         `json:"contactUtk,omitzero"`
	ExternalID        param.Opt[string]                                         `json:"externalId,omitzero"`
	MediaBridgeID     param.Opt[int64]                                          `json:"mediaBridgeId,omitzero"`
	MediaName         param.Opt[string]                                         `json:"mediaName,omitzero"`
	MediaURL          param.Opt[string]                                         `json:"mediaUrl,omitzero"`
	PageID            param.Opt[int64]                                          `json:"pageId,omitzero"`
	PageName          param.Opt[string]                                         `json:"pageName,omitzero"`
	PageURL           param.Opt[string]                                         `json:"pageUrl,omitzero"`
	paramObj
}

func (r MediaBridgeEventNewMediaPlayedPercentEventParams) MarshalJSON() (data []byte, err error) {
	type shadow MediaBridgeEventNewMediaPlayedPercentEventParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MediaBridgeEventNewMediaPlayedPercentEventParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeEventNewMediaPlayedPercentEventParamsMediaType string

const (
	MediaBridgeEventNewMediaPlayedPercentEventParamsMediaTypeVideo    MediaBridgeEventNewMediaPlayedPercentEventParamsMediaType = "VIDEO"
	MediaBridgeEventNewMediaPlayedPercentEventParamsMediaTypeAudio    MediaBridgeEventNewMediaPlayedPercentEventParamsMediaType = "AUDIO"
	MediaBridgeEventNewMediaPlayedPercentEventParamsMediaTypeDocument MediaBridgeEventNewMediaPlayedPercentEventParamsMediaType = "DOCUMENT"
	MediaBridgeEventNewMediaPlayedPercentEventParamsMediaTypeOther    MediaBridgeEventNewMediaPlayedPercentEventParamsMediaType = "OTHER"
	MediaBridgeEventNewMediaPlayedPercentEventParamsMediaTypeImage    MediaBridgeEventNewMediaPlayedPercentEventParamsMediaType = "IMAGE"
)
