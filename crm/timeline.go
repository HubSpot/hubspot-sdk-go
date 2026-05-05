// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package crm

import (
	"context"
	"encoding/json"
	"net/http"
	"slices"
	"time"

	"github.com/HubSpot/hubspot-sdk-go/internal/apijson"
	shimjson "github.com/HubSpot/hubspot-sdk-go/internal/encoding/json"
	"github.com/HubSpot/hubspot-sdk-go/internal/requestconfig"
	"github.com/HubSpot/hubspot-sdk-go/option"
	"github.com/HubSpot/hubspot-sdk-go/packages/param"
	"github.com/HubSpot/hubspot-sdk-go/packages/respjson"
	"github.com/HubSpot/hubspot-sdk-go/shared"
)

// TimelineService contains methods and other services that help with interacting
// with the hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTimelineService] method instead.
type TimelineService struct {
	options []option.RequestOption
	Batch   TimelineBatchService
}

// NewTimelineService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewTimelineService(opts ...option.RequestOption) (r TimelineService) {
	r = TimelineService{}
	r.options = opts
	r.Batch = NewTimelineBatchService(opts...)
	return
}

// Send a single instance of event data to a specified event type.
func (r *TimelineService) NewEvent(ctx context.Context, body TimelineNewEventParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "integrators/timeline/2026-03/events"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return err
}

func (r *TimelineService) NewProjectType(ctx context.Context, body TimelineNewProjectTypeParams, opts ...option.RequestOption) (res *AppEventResolutionResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "integrators/timeline/2026-03/types/projects"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type AppEventOccurrence struct {
	ID                           string              `json:"id" api:"required"`
	EventTypeName                string              `json:"eventTypeName" api:"required"`
	Properties                   map[string]string   `json:"properties" api:"required"`
	Domain                       string              `json:"domain"`
	Email                        string              `json:"email"`
	ExtraData                    any                 `json:"extraData"`
	ObjectID                     string              `json:"objectId"`
	ObjectTypeFullyQualifiedName string              `json:"objectTypeFullyQualifiedName"`
	TimelineIFrame               TimelineEventIFrame `json:"timelineIFrame"`
	Timestamp                    time.Time           `json:"timestamp" format:"date-time"`
	Utk                          string              `json:"utk"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                           respjson.Field
		EventTypeName                respjson.Field
		Properties                   respjson.Field
		Domain                       respjson.Field
		Email                        respjson.Field
		ExtraData                    respjson.Field
		ObjectID                     respjson.Field
		ObjectTypeFullyQualifiedName respjson.Field
		TimelineIFrame               respjson.Field
		Timestamp                    respjson.Field
		Utk                          respjson.Field
		ExtraFields                  map[string]respjson.Field
		raw                          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AppEventOccurrence) RawJSON() string { return r.JSON.raw }
func (r *AppEventOccurrence) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this AppEventOccurrence to a AppEventOccurrenceParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// AppEventOccurrenceParam.Overrides()
func (r AppEventOccurrence) ToParam() AppEventOccurrenceParam {
	return param.Override[AppEventOccurrenceParam](json.RawMessage(r.RawJSON()))
}

// The properties ID, EventTypeName, Properties are required.
type AppEventOccurrenceParam struct {
	ID                           string                   `json:"id" api:"required"`
	EventTypeName                string                   `json:"eventTypeName" api:"required"`
	Properties                   map[string]string        `json:"properties,omitzero" api:"required"`
	Domain                       param.Opt[string]        `json:"domain,omitzero"`
	Email                        param.Opt[string]        `json:"email,omitzero"`
	ObjectID                     param.Opt[string]        `json:"objectId,omitzero"`
	ObjectTypeFullyQualifiedName param.Opt[string]        `json:"objectTypeFullyQualifiedName,omitzero"`
	Timestamp                    param.Opt[time.Time]     `json:"timestamp,omitzero" format:"date-time"`
	Utk                          param.Opt[string]        `json:"utk,omitzero"`
	ExtraData                    any                      `json:"extraData,omitzero"`
	TimelineIFrame               TimelineEventIFrameParam `json:"timelineIFrame,omitzero"`
	paramObj
}

func (r AppEventOccurrenceParam) MarshalJSON() (data []byte, err error) {
	type shadow AppEventOccurrenceParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AppEventOccurrenceParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AppEventResolutionResponse struct {
	DeveloperQualifiedSymbol DeveloperQualifiedSymbol `json:"developerQualifiedSymbol" api:"required"`
	FullyQualifiedName       string                   `json:"fullyQualifiedName" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DeveloperQualifiedSymbol respjson.Field
		FullyQualifiedName       respjson.Field
		ExtraFields              map[string]respjson.Field
		raw                      string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AppEventResolutionResponse) RawJSON() string { return r.JSON.raw }
func (r *AppEventResolutionResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Inputs is required.
type BatchInputAppEventOccurrenceParam struct {
	Inputs []AppEventOccurrenceParam `json:"inputs,omitzero" api:"required"`
	paramObj
}

func (r BatchInputAppEventOccurrenceParam) MarshalJSON() (data []byte, err error) {
	type shadow BatchInputAppEventOccurrenceParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BatchInputAppEventOccurrenceParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BatchResponseAppEventOccurrence struct {
	CompletedAt time.Time            `json:"completedAt" api:"required" format:"date-time"`
	Results     []AppEventOccurrence `json:"results" api:"required"`
	StartedAt   time.Time            `json:"startedAt" api:"required" format:"date-time"`
	// Any of "CANCELED", "COMPLETE", "PENDING", "PROCESSING".
	Status      BatchResponseAppEventOccurrenceStatus `json:"status" api:"required"`
	Errors      []shared.StandardError                `json:"errors"`
	Links       map[string]string                     `json:"links"`
	NumErrors   int64                                 `json:"numErrors"`
	RequestedAt time.Time                             `json:"requestedAt" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CompletedAt respjson.Field
		Results     respjson.Field
		StartedAt   respjson.Field
		Status      respjson.Field
		Errors      respjson.Field
		Links       respjson.Field
		NumErrors   respjson.Field
		RequestedAt respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BatchResponseAppEventOccurrence) RawJSON() string { return r.JSON.raw }
func (r *BatchResponseAppEventOccurrence) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BatchResponseAppEventOccurrenceStatus string

const (
	BatchResponseAppEventOccurrenceStatusCanceled   BatchResponseAppEventOccurrenceStatus = "CANCELED"
	BatchResponseAppEventOccurrenceStatusComplete   BatchResponseAppEventOccurrenceStatus = "COMPLETE"
	BatchResponseAppEventOccurrenceStatusPending    BatchResponseAppEventOccurrenceStatus = "PENDING"
	BatchResponseAppEventOccurrenceStatusProcessing BatchResponseAppEventOccurrenceStatus = "PROCESSING"
)

type DeveloperQualifiedSymbol struct {
	DeveloperSymbol string `json:"developerSymbol" api:"required"`
	ProjectName     string `json:"projectName" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DeveloperSymbol respjson.Field
		ProjectName     respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DeveloperQualifiedSymbol) RawJSON() string { return r.JSON.raw }
func (r *DeveloperQualifiedSymbol) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties DeveloperSymbol, ProjectName are required.
type ExternalAppEventResolutionRequestParam struct {
	DeveloperSymbol string `json:"developerSymbol" api:"required"`
	ProjectName     string `json:"projectName" api:"required"`
	paramObj
}

func (r ExternalAppEventResolutionRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow ExternalAppEventResolutionRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExternalAppEventResolutionRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TimelineEventIFrame struct {
	// The label of the modal window that displays the iframe contents.
	HeaderLabel string `json:"headerLabel" api:"required"`
	// The height of the modal window in pixels.
	Height int64 `json:"height" api:"required"`
	// The text displaying the link that will display the iframe.
	LinkLabel string `json:"linkLabel" api:"required"`
	// The URI of the iframe contents.
	URL string `json:"url" api:"required"`
	// The width of the modal window in pixels.
	Width int64 `json:"width" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		HeaderLabel respjson.Field
		Height      respjson.Field
		LinkLabel   respjson.Field
		URL         respjson.Field
		Width       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TimelineEventIFrame) RawJSON() string { return r.JSON.raw }
func (r *TimelineEventIFrame) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this TimelineEventIFrame to a TimelineEventIFrameParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// TimelineEventIFrameParam.Overrides()
func (r TimelineEventIFrame) ToParam() TimelineEventIFrameParam {
	return param.Override[TimelineEventIFrameParam](json.RawMessage(r.RawJSON()))
}

// The properties HeaderLabel, Height, LinkLabel, URL, Width are required.
type TimelineEventIFrameParam struct {
	// The label of the modal window that displays the iframe contents.
	HeaderLabel string `json:"headerLabel" api:"required"`
	// The height of the modal window in pixels.
	Height int64 `json:"height" api:"required"`
	// The text displaying the link that will display the iframe.
	LinkLabel string `json:"linkLabel" api:"required"`
	// The URI of the iframe contents.
	URL string `json:"url" api:"required"`
	// The width of the modal window in pixels.
	Width int64 `json:"width" api:"required"`
	paramObj
}

func (r TimelineEventIFrameParam) MarshalJSON() (data []byte, err error) {
	type shadow TimelineEventIFrameParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TimelineEventIFrameParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TimelineNewEventParams struct {
	AppEventOccurrence AppEventOccurrenceParam
	paramObj
}

func (r TimelineNewEventParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.AppEventOccurrence)
}
func (r *TimelineNewEventParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TimelineNewProjectTypeParams struct {
	ExternalAppEventResolutionRequest ExternalAppEventResolutionRequestParam
	paramObj
}

func (r TimelineNewProjectTypeParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.ExternalAppEventResolutionRequest)
}
func (r *TimelineNewProjectTypeParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
