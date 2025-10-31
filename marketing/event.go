// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package marketing

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/stainless-sdks/hubspot-sdk-go/internal/apijson"
	"github.com/stainless-sdks/hubspot-sdk-go/internal/apiquery"
	shimjson "github.com/stainless-sdks/hubspot-sdk-go/internal/encoding/json"
	"github.com/stainless-sdks/hubspot-sdk-go/internal/requestconfig"
	"github.com/stainless-sdks/hubspot-sdk-go/option"
	"github.com/stainless-sdks/hubspot-sdk-go/packages/pagination"
	"github.com/stainless-sdks/hubspot-sdk-go/packages/param"
	"github.com/stainless-sdks/hubspot-sdk-go/packages/respjson"
	"github.com/stainless-sdks/hubspot-sdk-go/shared"
)

// EventService contains methods and other services that help with interacting with
// the Hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEventService] method instead.
type EventService struct {
	Options        []option.RequestOption
	Associations   EventAssociationService
	Attendance     EventAttendanceService
	Participations EventParticipationService
	Settings       EventSettingService
}

// NewEventService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewEventService(opts ...option.RequestOption) (r EventService) {
	r = EventService{}
	r.Options = opts
	r.Associations = NewEventAssociationService(opts...)
	r.Attendance = NewEventAttendanceService(opts...)
	r.Participations = NewEventParticipationService(opts...)
	r.Settings = NewEventSettingService(opts...)
	return
}

// Creates a new marketing event in HubSpot
func (r *EventService) New(ctx context.Context, body EventNewParams, opts ...option.RequestOption) (res *MarketingEventDefaultResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "marketing/v3/marketing-events/events"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Updates the details of an existing Marketing Event identified by its objectId,
// if it exists.
func (r *EventService) Update(ctx context.Context, objectID string, body EventUpdateParams, opts ...option.RequestOption) (res *MarketingEventPublicDefaultResponseV2, err error) {
	opts = slices.Concat(r.Options, opts)
	if objectID == "" {
		err = errors.New("missing required objectId parameter")
		return
	}
	path := fmt.Sprintf("marketing/v3/marketing-events/%s", objectID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Returns all Marketing Events available on the portal, along with their
// properties, regardless of whether they were created manually or through the
// application.
//
// The marketing events returned by this endpoint are sorted by objectId.
func (r *EventService) List(ctx context.Context, query EventListParams, opts ...option.RequestOption) (res *pagination.Page[MarketingEventPublicReadResponseV2], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "marketing/v3/marketing-events/"
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// Returns all Marketing Events available on the portal, along with their
// properties, regardless of whether they were created manually or through the
// application.
//
// The marketing events returned by this endpoint are sorted by objectId.
func (r *EventService) ListAutoPaging(ctx context.Context, query EventListParams, opts ...option.RequestOption) *pagination.PageAutoPager[MarketingEventPublicReadResponseV2] {
	return pagination.NewPageAutoPager(r.List(ctx, query, opts...))
}

// Deletes the existing Marketing Event with the specified objectId, if it exists.
func (r *EventService) Delete(ctx context.Context, objectID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if objectID == "" {
		err = errors.New("missing required objectId parameter")
		return
	}
	path := fmt.Sprintf("marketing/v3/marketing-events/%s", objectID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Mark a marketing event as cancelled.
func (r *EventService) CancelByExternalEventID(ctx context.Context, externalEventID string, body EventCancelByExternalEventIDParams, opts ...option.RequestOption) (res *MarketingEventDefaultResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if externalEventID == "" {
		err = errors.New("missing required externalEventId parameter")
		return
	}
	path := fmt.Sprintf("marketing/v3/marketing-events/events/%s/cancel", externalEventID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Mark a marketing event as completed
func (r *EventService) CompleteByExternalEventID(ctx context.Context, externalEventID string, params EventCompleteByExternalEventIDParams, opts ...option.RequestOption) (res *MarketingEventDefaultResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if externalEventID == "" {
		err = errors.New("missing required externalEventId parameter")
		return
	}
	path := fmt.Sprintf("marketing/v3/marketing-events/events/%s/complete", externalEventID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Deletes multiple Marketing Events from the portal based on their objectId, if
// they exist.
//
// Responses: 204: Returned if all specified Marketing Events were successfully
// deleted. 207: Returned if some objectIds did not correspond to any existing
// Marketing Events.
func (r *EventService) DeleteBatch(ctx context.Context, body EventDeleteBatchParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "marketing/v3/marketing-events/batch/archive"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Deletes multiple Marketing Events based on externalAccountId, externalEventId,
// and appId.
//
// Only Marketing Events created by the same apps will be deleted; events from
// other apps cannot be removed by this endpoint.
func (r *EventService) DeleteBatchByExternalEventID(ctx context.Context, body EventDeleteBatchByExternalEventIDParams, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "marketing/v3/marketing-events/events/delete"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Deletes the existing Marketing Event with the specified externalAccountId,
// externalEventId, if it exists.
//
// Only Marketing Events created by the same app can be deleted.
func (r *EventService) DeleteByExternalEventID(ctx context.Context, externalEventID string, body EventDeleteByExternalEventIDParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if externalEventID == "" {
		err = errors.New("missing required externalEventId parameter")
		return
	}
	path := fmt.Sprintf("marketing/v3/marketing-events/events/%s", externalEventID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, nil, opts...)
	return
}

// Returns the details of a Marketing Event with the specified objectId, if it
// exists.
func (r *EventService) Get(ctx context.Context, objectID string, opts ...option.RequestOption) (res *MarketingEventPublicReadResponseV2, err error) {
	opts = slices.Concat(r.Options, opts)
	if objectID == "" {
		err = errors.New("missing required objectId parameter")
		return
	}
	path := fmt.Sprintf("marketing/v3/marketing-events/%s", objectID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Returns the details of a Marketing Event with the specified externalAccountId,
// externalEventId, if it exists.
//
// Only Marketing Events created by the same app making the request can be
// retrieved.
func (r *EventService) GetByExternalEventID(ctx context.Context, externalEventID string, query EventGetByExternalEventIDParams, opts ...option.RequestOption) (res *MarketingEventPublicReadResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if externalEventID == "" {
		err = errors.New("missing required externalEventId parameter")
		return
	}
	path := fmt.Sprintf("marketing/v3/marketing-events/events/%s", externalEventID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Retrieves Marketing Events where the externalEventId matches the value provided
// in the request, limited to events created by the app making the request.
//
// Marketing Events created by other apps will not be included in the results.
func (r *EventService) SearchByExternalEventID(ctx context.Context, query EventSearchByExternalEventIDParams, opts ...option.RequestOption) (res *CollectionResponseSearchPublicResponseWrapperNoPaging, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "marketing/v3/marketing-events/events/search"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// This endpoint searches the portal for all Marketing Events whose externalEventId
// matches the value provided in the request.
//
// It retrieves the objectId and additional event details for each matching
// Marketing Event.
//
// Since multiple Marketing Events can have the same externalEventId, the endpoint
// returns all matching results.
//
// Note: Marketing Events become searchable by externalEventId a few minutes after
// creation.
func (r *EventService) SearchIdentifiersByExternalEventID(ctx context.Context, externalEventID string, opts ...option.RequestOption) (res *CollectionResponseWithTotalMarketingEventIdentifiersResponseNoPaging, err error) {
	opts = slices.Concat(r.Options, opts)
	if externalEventID == "" {
		err = errors.New("missing required externalEventId parameter")
		return
	}
	path := fmt.Sprintf("marketing/v3/marketing-events/%s/identifiers", externalEventID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates multiple Marketing Events on the portal based on their objectId, if they
// exist.
func (r *EventService) UpdateBatch(ctx context.Context, body EventUpdateBatchParams, opts ...option.RequestOption) (res *BatchResponseMarketingEventPublicDefaultResponseV2, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "marketing/v3/marketing-events/batch/update"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Updates the details of an existing Marketing Event identified by its
// externalAccountId, externalEventId if it exists.
//
// Only Marketing Events created by the same app can be updated.
func (r *EventService) UpdateByExternalEventID(ctx context.Context, externalEventID string, params EventUpdateByExternalEventIDParams, opts ...option.RequestOption) (res *MarketingEventPublicDefaultResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if externalEventID == "" {
		err = errors.New("missing required externalEventId parameter")
		return
	}
	path := fmt.Sprintf("marketing/v3/marketing-events/events/%s", externalEventID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return
}

// Upserts multiple Marketing Events. If a Marketing Event with the specified ID
// already exists, it will be updated; otherwise, a new event will be created.
//
// Only Marketing Events originally created by the same app can be updated.
func (r *EventService) UpsertBatch(ctx context.Context, body EventUpsertBatchParams, opts ...option.RequestOption) (res *BatchResponseMarketingEventPublicDefaultResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "marketing/v3/marketing-events/events/upsert"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Upserts a marketing event If there is an existing marketing event with the
// specified ID, it will be updated; otherwise a new event will be created.
func (r *EventService) UpsertByExternalEventID(ctx context.Context, externalEventID string, body EventUpsertByExternalEventIDParams, opts ...option.RequestOption) (res *MarketingEventPublicDefaultResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if externalEventID == "" {
		err = errors.New("missing required externalEventId parameter")
		return
	}
	path := fmt.Sprintf("marketing/v3/marketing-events/events/%s", externalEventID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Record a subscriber state between multiple HubSpot contacts and a marketing
// event, using contact email addresses. Note that the contact must already exist
// in HubSpot; a contact will not be created. The contactProperties field is used
// only when creating a new contact. These properties will not update existing
// contacts.
func (r *EventService) UpsertSubscriberStateByEmail(ctx context.Context, subscriberState string, params EventUpsertSubscriberStateByEmailParams, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if params.ExternalEventID == "" {
		err = errors.New("missing required externalEventId parameter")
		return
	}
	if subscriberState == "" {
		err = errors.New("missing required subscriberState parameter")
		return
	}
	path := fmt.Sprintf("marketing/v3/marketing-events/events/%s/%s/email-upsert", params.ExternalEventID, subscriberState)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Record a subscriber state between multiple HubSpot contacts and a marketing
// event, using HubSpot contact IDs. Note that the contact must already exist in
// HubSpot; a contact will not be created.
func (r *EventService) UpsertSubscriberStateByID(ctx context.Context, subscriberState string, params EventUpsertSubscriberStateByIDParams, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if params.ExternalEventID == "" {
		err = errors.New("missing required externalEventId parameter")
		return
	}
	if subscriberState == "" {
		err = errors.New("missing required subscriberState parameter")
		return
	}
	path := fmt.Sprintf("marketing/v3/marketing-events/events/%s/%s/upsert", params.ExternalEventID, subscriberState)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

type AppInfo struct {
	ID   string `json:"id,required"`
	Name string `json:"name,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AppInfo) RawJSON() string { return r.JSON.raw }
func (r *AppInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AttendanceCounters struct {
	Attended   int64 `json:"attended,required"`
	Cancelled  int64 `json:"cancelled,required"`
	NoShows    int64 `json:"noShows,required"`
	Registered int64 `json:"registered,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Attended    respjson.Field
		Cancelled   respjson.Field
		NoShows     respjson.Field
		Registered  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AttendanceCounters) RawJSON() string { return r.JSON.raw }
func (r *AttendanceCounters) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Inputs is required.
type BatchInputMarketingEventCreateRequestParams struct {
	Inputs []MarketingEventCreateRequestParams `json:"inputs,omitzero,required"`
	paramObj
}

func (r BatchInputMarketingEventCreateRequestParams) MarshalJSON() (data []byte, err error) {
	type shadow BatchInputMarketingEventCreateRequestParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BatchInputMarketingEventCreateRequestParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Inputs is required.
type BatchInputMarketingEventEmailSubscriberParam struct {
	// List of marketing event details to create or update
	Inputs []MarketingEventEmailSubscriberParam `json:"inputs,omitzero,required"`
	paramObj
}

func (r BatchInputMarketingEventEmailSubscriberParam) MarshalJSON() (data []byte, err error) {
	type shadow BatchInputMarketingEventEmailSubscriberParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BatchInputMarketingEventEmailSubscriberParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Inputs is required.
type BatchInputMarketingEventExternalUniqueIdentifierParam struct {
	Inputs []MarketingEventExternalUniqueIdentifierParam `json:"inputs,omitzero,required"`
	paramObj
}

func (r BatchInputMarketingEventExternalUniqueIdentifierParam) MarshalJSON() (data []byte, err error) {
	type shadow BatchInputMarketingEventExternalUniqueIdentifierParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BatchInputMarketingEventExternalUniqueIdentifierParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Inputs is required.
type BatchInputMarketingEventPublicObjectIDDeleteRequestParam struct {
	Inputs []MarketingEventPublicObjectIDDeleteRequestParam `json:"inputs,omitzero,required"`
	paramObj
}

func (r BatchInputMarketingEventPublicObjectIDDeleteRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow BatchInputMarketingEventPublicObjectIDDeleteRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BatchInputMarketingEventPublicObjectIDDeleteRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Inputs is required.
type BatchInputMarketingEventPublicUpdateRequestFullV2Param struct {
	Inputs []MarketingEventPublicUpdateRequestFullV2Param `json:"inputs,omitzero,required"`
	paramObj
}

func (r BatchInputMarketingEventPublicUpdateRequestFullV2Param) MarshalJSON() (data []byte, err error) {
	type shadow BatchInputMarketingEventPublicUpdateRequestFullV2Param
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BatchInputMarketingEventPublicUpdateRequestFullV2Param) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Inputs is required.
type BatchInputMarketingEventSubscriberParam struct {
	// List of HubSpot contacts to subscribe to the marketing event
	Inputs []MarketingEventSubscriberParam `json:"inputs,omitzero,required"`
	paramObj
}

func (r BatchInputMarketingEventSubscriberParam) MarshalJSON() (data []byte, err error) {
	type shadow BatchInputMarketingEventSubscriberParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BatchInputMarketingEventSubscriberParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BatchResponseMarketingEventPublicDefaultResponse struct {
	CompletedAt time.Time                             `json:"completedAt,required" format:"date-time"`
	Results     []MarketingEventPublicDefaultResponse `json:"results,required"`
	StartedAt   time.Time                             `json:"startedAt,required" format:"date-time"`
	// Any of "PENDING", "PROCESSING", "CANCELED", "COMPLETE".
	Status      BatchResponseMarketingEventPublicDefaultResponseStatus `json:"status,required"`
	Errors      []shared.StandardError                                 `json:"errors"`
	Links       map[string]string                                      `json:"links"`
	NumErrors   int64                                                  `json:"numErrors"`
	RequestedAt time.Time                                              `json:"requestedAt" format:"date-time"`
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
func (r BatchResponseMarketingEventPublicDefaultResponse) RawJSON() string { return r.JSON.raw }
func (r *BatchResponseMarketingEventPublicDefaultResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BatchResponseMarketingEventPublicDefaultResponseStatus string

const (
	BatchResponseMarketingEventPublicDefaultResponseStatusPending    BatchResponseMarketingEventPublicDefaultResponseStatus = "PENDING"
	BatchResponseMarketingEventPublicDefaultResponseStatusProcessing BatchResponseMarketingEventPublicDefaultResponseStatus = "PROCESSING"
	BatchResponseMarketingEventPublicDefaultResponseStatusCanceled   BatchResponseMarketingEventPublicDefaultResponseStatus = "CANCELED"
	BatchResponseMarketingEventPublicDefaultResponseStatusComplete   BatchResponseMarketingEventPublicDefaultResponseStatus = "COMPLETE"
)

type BatchResponseMarketingEventPublicDefaultResponseV2 struct {
	CompletedAt time.Time                               `json:"completedAt,required" format:"date-time"`
	Results     []MarketingEventPublicDefaultResponseV2 `json:"results,required"`
	StartedAt   time.Time                               `json:"startedAt,required" format:"date-time"`
	// Any of "PENDING", "PROCESSING", "CANCELED", "COMPLETE".
	Status      BatchResponseMarketingEventPublicDefaultResponseV2Status `json:"status,required"`
	Links       map[string]string                                        `json:"links"`
	RequestedAt time.Time                                                `json:"requestedAt" format:"date-time"`
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
func (r BatchResponseMarketingEventPublicDefaultResponseV2) RawJSON() string { return r.JSON.raw }
func (r *BatchResponseMarketingEventPublicDefaultResponseV2) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BatchResponseMarketingEventPublicDefaultResponseV2Status string

const (
	BatchResponseMarketingEventPublicDefaultResponseV2StatusPending    BatchResponseMarketingEventPublicDefaultResponseV2Status = "PENDING"
	BatchResponseMarketingEventPublicDefaultResponseV2StatusProcessing BatchResponseMarketingEventPublicDefaultResponseV2Status = "PROCESSING"
	BatchResponseMarketingEventPublicDefaultResponseV2StatusCanceled   BatchResponseMarketingEventPublicDefaultResponseV2Status = "CANCELED"
	BatchResponseMarketingEventPublicDefaultResponseV2StatusComplete   BatchResponseMarketingEventPublicDefaultResponseV2Status = "COMPLETE"
)

type BatchResponseSubscriberEmailResponse struct {
	CompletedAt time.Time                 `json:"completedAt,required" format:"date-time"`
	Results     []SubscriberEmailResponse `json:"results,required"`
	StartedAt   time.Time                 `json:"startedAt,required" format:"date-time"`
	// Any of "PENDING", "PROCESSING", "CANCELED", "COMPLETE".
	Status      BatchResponseSubscriberEmailResponseStatus `json:"status,required"`
	Errors      []shared.StandardError                     `json:"errors"`
	Links       map[string]string                          `json:"links"`
	NumErrors   int64                                      `json:"numErrors"`
	RequestedAt time.Time                                  `json:"requestedAt" format:"date-time"`
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
func (r BatchResponseSubscriberEmailResponse) RawJSON() string { return r.JSON.raw }
func (r *BatchResponseSubscriberEmailResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BatchResponseSubscriberEmailResponseStatus string

const (
	BatchResponseSubscriberEmailResponseStatusPending    BatchResponseSubscriberEmailResponseStatus = "PENDING"
	BatchResponseSubscriberEmailResponseStatusProcessing BatchResponseSubscriberEmailResponseStatus = "PROCESSING"
	BatchResponseSubscriberEmailResponseStatusCanceled   BatchResponseSubscriberEmailResponseStatus = "CANCELED"
	BatchResponseSubscriberEmailResponseStatusComplete   BatchResponseSubscriberEmailResponseStatus = "COMPLETE"
)

type BatchResponseSubscriberVidResponse struct {
	CompletedAt time.Time               `json:"completedAt,required" format:"date-time"`
	Results     []SubscriberVidResponse `json:"results,required"`
	StartedAt   time.Time               `json:"startedAt,required" format:"date-time"`
	// Any of "PENDING", "PROCESSING", "CANCELED", "COMPLETE".
	Status      BatchResponseSubscriberVidResponseStatus `json:"status,required"`
	Errors      []shared.StandardError                   `json:"errors"`
	Links       map[string]string                        `json:"links"`
	NumErrors   int64                                    `json:"numErrors"`
	RequestedAt time.Time                                `json:"requestedAt" format:"date-time"`
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
func (r BatchResponseSubscriberVidResponse) RawJSON() string { return r.JSON.raw }
func (r *BatchResponseSubscriberVidResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BatchResponseSubscriberVidResponseStatus string

const (
	BatchResponseSubscriberVidResponseStatusPending    BatchResponseSubscriberVidResponseStatus = "PENDING"
	BatchResponseSubscriberVidResponseStatusProcessing BatchResponseSubscriberVidResponseStatus = "PROCESSING"
	BatchResponseSubscriberVidResponseStatusCanceled   BatchResponseSubscriberVidResponseStatus = "CANCELED"
	BatchResponseSubscriberVidResponseStatusComplete   BatchResponseSubscriberVidResponseStatus = "COMPLETE"
)

type CollectionResponseMarketingEventPublicReadResponseV2ForwardPaging struct {
	Results []MarketingEventPublicReadResponseV2 `json:"results,required"`
	Paging  shared.ForwardPaging                 `json:"paging"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Results     respjson.Field
		Paging      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CollectionResponseMarketingEventPublicReadResponseV2ForwardPaging) RawJSON() string {
	return r.JSON.raw
}
func (r *CollectionResponseMarketingEventPublicReadResponseV2ForwardPaging) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CollectionResponseSearchPublicResponseWrapperNoPaging struct {
	Results []SearchPublicResponseWrapper `json:"results,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Results     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CollectionResponseSearchPublicResponseWrapperNoPaging) RawJSON() string { return r.JSON.raw }
func (r *CollectionResponseSearchPublicResponseWrapperNoPaging) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CollectionResponseWithTotalMarketingEventIdentifiersResponseNoPaging struct {
	Results []MarketingEventIdentifiersResponse `json:"results,required"`
	Total   int64                               `json:"total,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Results     respjson.Field
		Total       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CollectionResponseWithTotalMarketingEventIdentifiersResponseNoPaging) RawJSON() string {
	return r.JSON.raw
}
func (r *CollectionResponseWithTotalMarketingEventIdentifiersResponseNoPaging) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CollectionResponseWithTotalParticipationBreakdownForwardPaging struct {
	Results []ParticipationBreakdown `json:"results,required"`
	Total   int64                    `json:"total,required"`
	Paging  shared.ForwardPaging     `json:"paging"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Results     respjson.Field
		Total       respjson.Field
		Paging      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CollectionResponseWithTotalParticipationBreakdownForwardPaging) RawJSON() string {
	return r.JSON.raw
}
func (r *CollectionResponseWithTotalParticipationBreakdownForwardPaging) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CollectionResponseWithTotalPublicListNoPaging struct {
	Results []PublicList `json:"results,required"`
	Total   int64        `json:"total,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Results     respjson.Field
		Total       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CollectionResponseWithTotalPublicListNoPaging) RawJSON() string { return r.JSON.raw }
func (r *CollectionResponseWithTotalPublicListNoPaging) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ContactAssociation struct {
	ContactID string `json:"contactId,required"`
	Email     string `json:"email,required"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ContactID   respjson.Field
		Email       respjson.Field
		Firstname   respjson.Field
		Lastname    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ContactAssociation) RawJSON() string { return r.JSON.raw }
func (r *ContactAssociation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CRMPropertyWrapper struct {
	Name  string `json:"name,required"`
	Value string `json:"value,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CRMPropertyWrapper) RawJSON() string { return r.JSON.raw }
func (r *CRMPropertyWrapper) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EventDetailSettings struct {
	// The id of the application the settings are for
	AppID int64 `json:"appId,required"`
	// The url that will be used to fetch marketing event details by id
	EventDetailsURL string `json:"eventDetailsUrl,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AppID           respjson.Field
		EventDetailsURL respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EventDetailSettings) RawJSON() string { return r.JSON.raw }
func (r *EventDetailSettings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property EventDetailsURL is required.
type EventDetailSettingsURLParam struct {
	// The url that will be used to fetch marketing event details by id. Must contain a
	// `%s` character sequence that will be substituted with the event id. For example:
	// `https://my.event.app/events/%s`
	EventDetailsURL string `json:"eventDetailsUrl,required"`
	paramObj
}

func (r EventDetailSettingsURLParam) MarshalJSON() (data []byte, err error) {
	type shadow EventDetailSettingsURLParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EventDetailSettingsURLParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MarketingEventAssociation struct {
	MarketingEventID  string `json:"marketingEventId,required"`
	Name              string `json:"name,required"`
	ExternalAccountID string `json:"externalAccountId"`
	ExternalEventID   string `json:"externalEventId"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		MarketingEventID  respjson.Field
		Name              respjson.Field
		ExternalAccountID respjson.Field
		ExternalEventID   respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MarketingEventAssociation) RawJSON() string { return r.JSON.raw }
func (r *MarketingEventAssociation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties EndDateTime, StartDateTime are required.
type MarketingEventCompleteRequestParams struct {
	EndDateTime   time.Time `json:"endDateTime,required" format:"date-time"`
	StartDateTime time.Time `json:"startDateTime,required" format:"date-time"`
	paramObj
}

func (r MarketingEventCompleteRequestParams) MarshalJSON() (data []byte, err error) {
	type shadow MarketingEventCompleteRequestParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MarketingEventCompleteRequestParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties EventName, EventOrganizer, ExternalAccountID, ExternalEventID are
// required.
type MarketingEventCreateRequestParams struct {
	// The name of the marketing event.
	EventName string `json:"eventName,required"`
	// The name of the organizer of the marketing event.
	EventOrganizer string `json:"eventOrganizer,required"`
	// The accountId that is associated with this marketing event in the external event
	// application.
	ExternalAccountID string `json:"externalAccountId,required"`
	// The id of the marketing event in the external event application.
	ExternalEventID string `json:"externalEventId,required"`
	// The end date and time of the marketing event.
	EndDateTime param.Opt[time.Time] `json:"endDateTime,omitzero" format:"date-time"`
	// Indicates if the marketing event has been cancelled. Defaults to `false`
	EventCancelled param.Opt[bool] `json:"eventCancelled,omitzero"`
	EventCompleted param.Opt[bool] `json:"eventCompleted,omitzero"`
	// The description of the marketing event.
	EventDescription param.Opt[string] `json:"eventDescription,omitzero"`
	// Describes what type of event this is. For example: `WEBINAR`, `CONFERENCE`,
	// `WORKSHOP`
	EventType param.Opt[string] `json:"eventType,omitzero"`
	// A URL in the external event application where the marketing event can be
	// managed.
	EventURL param.Opt[string] `json:"eventUrl,omitzero"`
	// The start date and time of the marketing event.
	StartDateTime param.Opt[time.Time] `json:"startDateTime,omitzero" format:"date-time"`
	// A list of PropertyValues. These can be whatever kind of property names and
	// values you want. However, they must already exist on the HubSpot account's
	// definition of the MarketingEvent Object. If they don't they will be filtered out
	// and not set. In order to do this you'll need to create a new PropertyGroup on
	// the HubSpot account's MarketingEvent object for your specific app and create the
	// Custom Property you want to track on that HubSpot account. Do not create any new
	// default properties on the MarketingEvent object as that will apply to all
	// HubSpot accounts.
	CustomProperties []PropertyValueParam `json:"customProperties,omitzero"`
	paramObj
}

func (r MarketingEventCreateRequestParams) MarshalJSON() (data []byte, err error) {
	type shadow MarketingEventCreateRequestParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MarketingEventCreateRequestParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MarketingEventDefaultResponse struct {
	// The name of the marketing event.
	EventName string `json:"eventName,required"`
	// The name of the organizer of the marketing event.
	EventOrganizer string `json:"eventOrganizer,required"`
	// A list of PropertyValues. These can be whatever kind of property names and
	// values you want. However, they must already exist on the HubSpot account's
	// definition of the MarketingEvent Object. If they don't they will be filtered out
	// and not set. In order to do this you'll need to create a new PropertyGroup on
	// the HubSpot account's MarketingEvent object for your specific app and create the
	// Custom Property you want to track on that HubSpot account. Do not create any new
	// default properties on the MarketingEvent object as that will apply to all
	// HubSpot accounts.
	CustomProperties []PropertyValue `json:"customProperties"`
	// The end date and time of the marketing event.
	EndDateTime time.Time `json:"endDateTime" format:"date-time"`
	// Indicates if the marketing event has been cancelled.
	EventCancelled bool `json:"eventCancelled"`
	EventCompleted bool `json:"eventCompleted"`
	// The description of the marketing event.
	EventDescription string `json:"eventDescription"`
	// The type of the marketing event.
	EventType string `json:"eventType"`
	// The URL in the external event application where the marketing event can be
	// managed.
	EventURL string `json:"eventUrl"`
	ObjectID string `json:"objectId"`
	// The start date and time of the marketing event.
	StartDateTime time.Time `json:"startDateTime" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		EventName        respjson.Field
		EventOrganizer   respjson.Field
		CustomProperties respjson.Field
		EndDateTime      respjson.Field
		EventCancelled   respjson.Field
		EventCompleted   respjson.Field
		EventDescription respjson.Field
		EventType        respjson.Field
		EventURL         respjson.Field
		ObjectID         respjson.Field
		StartDateTime    respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MarketingEventDefaultResponse) RawJSON() string { return r.JSON.raw }
func (r *MarketingEventDefaultResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Email, InteractionDateTime are required.
type MarketingEventEmailSubscriberParam struct {
	// The email address of the contact in HubSpot to associate with the event.
	Email string `json:"email,required"`
	// Timestamp in milliseconds at which the contact subscribed to the event.
	InteractionDateTime int64             `json:"interactionDateTime,required"`
	ContactProperties   map[string]string `json:"contactProperties,omitzero"`
	Properties          map[string]string `json:"properties,omitzero"`
	paramObj
}

func (r MarketingEventEmailSubscriberParam) MarshalJSON() (data []byte, err error) {
	type shadow MarketingEventEmailSubscriberParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MarketingEventEmailSubscriberParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties AppID, ExternalAccountID, ExternalEventID are required.
type MarketingEventExternalUniqueIdentifierParam struct {
	// The id of the application that created the marketing event in HubSpot.
	AppID int64 `json:"appId,required"`
	// The accountId that is associated with this marketing event in the external event
	// application.
	ExternalAccountID string `json:"externalAccountId,required"`
	// The id of the marketing event in the external event application.
	ExternalEventID string `json:"externalEventId,required"`
	paramObj
}

func (r MarketingEventExternalUniqueIdentifierParam) MarshalJSON() (data []byte, err error) {
	type shadow MarketingEventExternalUniqueIdentifierParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MarketingEventExternalUniqueIdentifierParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MarketingEventIdentifiersResponse struct {
	ExternalEventID    string  `json:"externalEventId,required"`
	MarketingEventName string  `json:"marketingEventName,required"`
	ObjectID           string  `json:"objectId,required"`
	AppInfo            AppInfo `json:"appInfo"`
	ExternalAccountID  string  `json:"externalAccountId"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ExternalEventID    respjson.Field
		MarketingEventName respjson.Field
		ObjectID           respjson.Field
		AppInfo            respjson.Field
		ExternalAccountID  respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MarketingEventIdentifiersResponse) RawJSON() string { return r.JSON.raw }
func (r *MarketingEventIdentifiersResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MarketingEventPublicDefaultResponse struct {
	ID        string    `json:"id,required"`
	CreatedAt time.Time `json:"createdAt,required" format:"date-time"`
	// The name of the marketing event.
	EventName string `json:"eventName,required"`
	// The name of the organizer of the marketing event.
	EventOrganizer string    `json:"eventOrganizer,required"`
	UpdatedAt      time.Time `json:"updatedAt,required" format:"date-time"`
	// A list of PropertyValues. These can be whatever kind of property names and
	// values you want. However, they must already exist on the HubSpot account's
	// definition of the MarketingEvent Object. If they don't they will be filtered out
	// and not set. In order to do this you'll need to create a new PropertyGroup on
	// the HubSpot account's MarketingEvent object for your specific app and create the
	// Custom Property you want to track on that HubSpot account. Do not create any new
	// default properties on the MarketingEvent object as that will apply to all
	// HubSpot accounts.
	CustomProperties []PropertyValue `json:"customProperties"`
	// The end date and time of the marketing event.
	EndDateTime time.Time `json:"endDateTime" format:"date-time"`
	// Indicates if the marketing event has been cancelled.
	EventCancelled bool `json:"eventCancelled"`
	EventCompleted bool `json:"eventCompleted"`
	// The description of the marketing event.
	EventDescription string `json:"eventDescription"`
	// The type of the marketing event.
	EventType string `json:"eventType"`
	// A URL in the external event application where the marketing event can be
	// managed.
	EventURL string `json:"eventUrl"`
	ObjectID string `json:"objectId"`
	// The start date and time of the marketing event.
	StartDateTime time.Time `json:"startDateTime" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		CreatedAt        respjson.Field
		EventName        respjson.Field
		EventOrganizer   respjson.Field
		UpdatedAt        respjson.Field
		CustomProperties respjson.Field
		EndDateTime      respjson.Field
		EventCancelled   respjson.Field
		EventCompleted   respjson.Field
		EventDescription respjson.Field
		EventType        respjson.Field
		EventURL         respjson.Field
		ObjectID         respjson.Field
		StartDateTime    respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MarketingEventPublicDefaultResponse) RawJSON() string { return r.JSON.raw }
func (r *MarketingEventPublicDefaultResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MarketingEventPublicDefaultResponseV2 struct {
	CreatedAt        time.Time            `json:"createdAt,required" format:"date-time"`
	CustomProperties []CRMPropertyWrapper `json:"customProperties,required"`
	EventName        string               `json:"eventName,required"`
	ObjectID         string               `json:"objectId,required"`
	UpdatedAt        time.Time            `json:"updatedAt,required" format:"date-time"`
	AppInfo          AppInfo              `json:"appInfo"`
	EndDateTime      time.Time            `json:"endDateTime" format:"date-time"`
	EventCancelled   bool                 `json:"eventCancelled"`
	EventCompleted   bool                 `json:"eventCompleted"`
	EventDescription string               `json:"eventDescription"`
	EventOrganizer   string               `json:"eventOrganizer"`
	EventType        string               `json:"eventType"`
	EventURL         string               `json:"eventUrl"`
	StartDateTime    time.Time            `json:"startDateTime" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreatedAt        respjson.Field
		CustomProperties respjson.Field
		EventName        respjson.Field
		ObjectID         respjson.Field
		UpdatedAt        respjson.Field
		AppInfo          respjson.Field
		EndDateTime      respjson.Field
		EventCancelled   respjson.Field
		EventCompleted   respjson.Field
		EventDescription respjson.Field
		EventOrganizer   respjson.Field
		EventType        respjson.Field
		EventURL         respjson.Field
		StartDateTime    respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MarketingEventPublicDefaultResponseV2) RawJSON() string { return r.JSON.raw }
func (r *MarketingEventPublicDefaultResponseV2) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property ObjectID is required.
type MarketingEventPublicObjectIDDeleteRequestParam struct {
	ObjectID string `json:"objectId,required"`
	paramObj
}

func (r MarketingEventPublicObjectIDDeleteRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow MarketingEventPublicObjectIDDeleteRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MarketingEventPublicObjectIDDeleteRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MarketingEventPublicReadResponse struct {
	ID string `json:"id,required"`
	// The number of HubSpot contacts that attended this marketing event.
	Attendees int64 `json:"attendees,required"`
	// The number of HubSpot contacts that registered for this marketing event, but
	// later cancelled their registration.
	Cancellations int64     `json:"cancellations,required"`
	CreatedAt     time.Time `json:"createdAt,required" format:"date-time"`
	// The name of the marketing event.
	EventName string `json:"eventName,required"`
	// The name of the organizer of the marketing event.
	EventOrganizer string `json:"eventOrganizer,required"`
	// The id of the marketing event in the external event application.
	ExternalEventID string `json:"externalEventId,required"`
	// The number of HubSpot contacts that registered for this marketing event, but did
	// not attend. This field only had a value when the event is over.
	NoShows int64 `json:"noShows,required"`
	// The number of HubSpot contacts that registered for this marketing event.
	Registrants int64     `json:"registrants,required"`
	UpdatedAt   time.Time `json:"updatedAt,required" format:"date-time"`
	// A list of PropertyValues. These can be whatever kind of property names and
	// values you want. However, they must already exist on the HubSpot account's
	// definition of the MarketingEvent Object. If they don't they will be filtered out
	// and not set. In order to do this you'll need to create a new PropertyGroup on
	// the HubSpot account's MarketingEvent object for your specific app and create the
	// Custom Property you want to track on that HubSpot account. Do not create any new
	// default properties on the MarketingEvent object as that will apply to all
	// HubSpot accounts.
	CustomProperties []PropertyValue `json:"customProperties"`
	// The end date and time of the marketing event.
	EndDateTime time.Time `json:"endDateTime" format:"date-time"`
	// Indicates if the marketing event has been cancelled.
	EventCancelled bool `json:"eventCancelled"`
	EventCompleted bool `json:"eventCompleted"`
	// The description of the marketing event.
	EventDescription string `json:"eventDescription"`
	// The type of the marketing event.
	EventType string `json:"eventType"`
	// A URL in the external event application where the marketing event can be
	// managed.
	EventURL string `json:"eventUrl"`
	ObjectID string `json:"objectId"`
	// The start date and time of the marketing event.
	StartDateTime time.Time `json:"startDateTime" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		Attendees        respjson.Field
		Cancellations    respjson.Field
		CreatedAt        respjson.Field
		EventName        respjson.Field
		EventOrganizer   respjson.Field
		ExternalEventID  respjson.Field
		NoShows          respjson.Field
		Registrants      respjson.Field
		UpdatedAt        respjson.Field
		CustomProperties respjson.Field
		EndDateTime      respjson.Field
		EventCancelled   respjson.Field
		EventCompleted   respjson.Field
		EventDescription respjson.Field
		EventType        respjson.Field
		EventURL         respjson.Field
		ObjectID         respjson.Field
		StartDateTime    respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MarketingEventPublicReadResponse) RawJSON() string { return r.JSON.raw }
func (r *MarketingEventPublicReadResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MarketingEventPublicReadResponseV2 struct {
	CreatedAt        time.Time            `json:"createdAt,required" format:"date-time"`
	CustomProperties []CRMPropertyWrapper `json:"customProperties,required"`
	EventName        string               `json:"eventName,required"`
	ObjectID         string               `json:"objectId,required"`
	UpdatedAt        time.Time            `json:"updatedAt,required" format:"date-time"`
	AppInfo          AppInfo              `json:"appInfo"`
	Attendees        int64                `json:"attendees"`
	Cancellations    int64                `json:"cancellations"`
	EndDateTime      time.Time            `json:"endDateTime" format:"date-time"`
	EventCancelled   bool                 `json:"eventCancelled"`
	EventCompleted   bool                 `json:"eventCompleted"`
	EventDescription string               `json:"eventDescription"`
	EventOrganizer   string               `json:"eventOrganizer"`
	EventStatus      string               `json:"eventStatus"`
	EventType        string               `json:"eventType"`
	EventURL         string               `json:"eventUrl"`
	ExternalEventID  string               `json:"externalEventId"`
	NoShows          int64                `json:"noShows"`
	Registrants      int64                `json:"registrants"`
	StartDateTime    time.Time            `json:"startDateTime" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreatedAt        respjson.Field
		CustomProperties respjson.Field
		EventName        respjson.Field
		ObjectID         respjson.Field
		UpdatedAt        respjson.Field
		AppInfo          respjson.Field
		Attendees        respjson.Field
		Cancellations    respjson.Field
		EndDateTime      respjson.Field
		EventCancelled   respjson.Field
		EventCompleted   respjson.Field
		EventDescription respjson.Field
		EventOrganizer   respjson.Field
		EventStatus      respjson.Field
		EventType        respjson.Field
		EventURL         respjson.Field
		ExternalEventID  respjson.Field
		NoShows          respjson.Field
		Registrants      respjson.Field
		StartDateTime    respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MarketingEventPublicReadResponseV2) RawJSON() string { return r.JSON.raw }
func (r *MarketingEventPublicReadResponseV2) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties CustomProperties, ObjectID are required.
type MarketingEventPublicUpdateRequestFullV2Param struct {
	CustomProperties []PropertyValueParam `json:"customProperties,omitzero,required"`
	ObjectID         string               `json:"objectId,required"`
	EndDateTime      param.Opt[time.Time] `json:"endDateTime,omitzero" format:"date-time"`
	EventCancelled   param.Opt[bool]      `json:"eventCancelled,omitzero"`
	EventDescription param.Opt[string]    `json:"eventDescription,omitzero"`
	EventName        param.Opt[string]    `json:"eventName,omitzero"`
	EventOrganizer   param.Opt[string]    `json:"eventOrganizer,omitzero"`
	EventType        param.Opt[string]    `json:"eventType,omitzero"`
	EventURL         param.Opt[string]    `json:"eventUrl,omitzero"`
	StartDateTime    param.Opt[time.Time] `json:"startDateTime,omitzero" format:"date-time"`
	paramObj
}

func (r MarketingEventPublicUpdateRequestFullV2Param) MarshalJSON() (data []byte, err error) {
	type shadow MarketingEventPublicUpdateRequestFullV2Param
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MarketingEventPublicUpdateRequestFullV2Param) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property CustomProperties is required.
type MarketingEventPublicUpdateRequestV2Param struct {
	CustomProperties []PropertyValueParam `json:"customProperties,omitzero,required"`
	EndDateTime      param.Opt[time.Time] `json:"endDateTime,omitzero" format:"date-time"`
	EventCancelled   param.Opt[bool]      `json:"eventCancelled,omitzero"`
	EventDescription param.Opt[string]    `json:"eventDescription,omitzero"`
	EventName        param.Opt[string]    `json:"eventName,omitzero"`
	EventOrganizer   param.Opt[string]    `json:"eventOrganizer,omitzero"`
	EventType        param.Opt[string]    `json:"eventType,omitzero"`
	EventURL         param.Opt[string]    `json:"eventUrl,omitzero"`
	StartDateTime    param.Opt[time.Time] `json:"startDateTime,omitzero" format:"date-time"`
	paramObj
}

func (r MarketingEventPublicUpdateRequestV2Param) MarshalJSON() (data []byte, err error) {
	type shadow MarketingEventPublicUpdateRequestV2Param
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MarketingEventPublicUpdateRequestV2Param) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property InteractionDateTime is required.
type MarketingEventSubscriberParam struct {
	// Timestamp in milliseconds at which the contact subscribed to the event.
	InteractionDateTime int64             `json:"interactionDateTime,required"`
	Vid                 param.Opt[int64]  `json:"vid,omitzero"`
	Properties          map[string]string `json:"properties,omitzero"`
	paramObj
}

func (r MarketingEventSubscriberParam) MarshalJSON() (data []byte, err error) {
	type shadow MarketingEventSubscriberParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MarketingEventSubscriberParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MarketingEventUpdateRequestParams struct {
	// The end date and time of the marketing event.
	EndDateTime param.Opt[time.Time] `json:"endDateTime,omitzero" format:"date-time"`
	// Indicates if the marketing event has been cancelled. Defaults to `false`
	EventCancelled param.Opt[bool] `json:"eventCancelled,omitzero"`
	EventCompleted param.Opt[bool] `json:"eventCompleted,omitzero"`
	// The description of the marketing event.
	EventDescription param.Opt[string] `json:"eventDescription,omitzero"`
	// The name of the marketing event.
	EventName param.Opt[string] `json:"eventName,omitzero"`
	// The name of the organizer of the marketing event.
	EventOrganizer param.Opt[string] `json:"eventOrganizer,omitzero"`
	// Describes what type of event this is. For example: `WEBINAR`, `CONFERENCE`,
	// `WORKSHOP`
	EventType param.Opt[string] `json:"eventType,omitzero"`
	// A URL in the external event application where the marketing event can be
	// managed.
	EventURL param.Opt[string] `json:"eventUrl,omitzero"`
	// The start date and time of the marketing event.
	StartDateTime param.Opt[time.Time] `json:"startDateTime,omitzero" format:"date-time"`
	// A list of PropertyValues. These can be whatever kind of property names and
	// values you want. However, they must already exist on the HubSpot account's
	// definition of the MarketingEvent Object. If they don't they will be filtered out
	// and not set. In order to do this you'll need to create a new PropertyGroup on
	// the HubSpot account's MarketingEvent object for your specific app and create the
	// Custom Property you want to track on that HubSpot account. Do not create any new
	// default properties on the MarketingEvent object as that will apply to all
	// HubSpot accounts.
	CustomProperties []PropertyValueParam `json:"customProperties,omitzero"`
	paramObj
}

func (r MarketingEventUpdateRequestParams) MarshalJSON() (data []byte, err error) {
	type shadow MarketingEventUpdateRequestParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MarketingEventUpdateRequestParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ParticipationAssociations struct {
	Contact        ContactAssociation        `json:"contact,required"`
	MarketingEvent MarketingEventAssociation `json:"marketingEvent,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Contact        respjson.Field
		MarketingEvent respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ParticipationAssociations) RawJSON() string { return r.JSON.raw }
func (r *ParticipationAssociations) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ParticipationBreakdown struct {
	ID           string                    `json:"id,required"`
	Associations ParticipationAssociations `json:"associations,required"`
	CreatedAt    time.Time                 `json:"createdAt,required" format:"date-time"`
	Properties   ParticipationProperties   `json:"properties,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		Associations respjson.Field
		CreatedAt    respjson.Field
		Properties   respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ParticipationBreakdown) RawJSON() string { return r.JSON.raw }
func (r *ParticipationBreakdown) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ParticipationProperties struct {
	// Any of "REGISTERED", "ATTENDED", "CANCELLED", "EMPTY", "NO_SHOW".
	AttendanceState           ParticipationPropertiesAttendanceState `json:"attendanceState,required"`
	OccurredAt                int64                                  `json:"occurredAt,required"`
	AttendanceDurationSeconds int64                                  `json:"attendanceDurationSeconds"`
	AttendancePercentage      string                                 `json:"attendancePercentage"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AttendanceState           respjson.Field
		OccurredAt                respjson.Field
		AttendanceDurationSeconds respjson.Field
		AttendancePercentage      respjson.Field
		ExtraFields               map[string]respjson.Field
		raw                       string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ParticipationProperties) RawJSON() string { return r.JSON.raw }
func (r *ParticipationProperties) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ParticipationPropertiesAttendanceState string

const (
	ParticipationPropertiesAttendanceStateRegistered ParticipationPropertiesAttendanceState = "REGISTERED"
	ParticipationPropertiesAttendanceStateAttended   ParticipationPropertiesAttendanceState = "ATTENDED"
	ParticipationPropertiesAttendanceStateCancelled  ParticipationPropertiesAttendanceState = "CANCELLED"
	ParticipationPropertiesAttendanceStateEmpty      ParticipationPropertiesAttendanceState = "EMPTY"
	ParticipationPropertiesAttendanceStateNoShow     ParticipationPropertiesAttendanceState = "NO_SHOW"
)

// Represents a single custom property of a marketing event, storing its name,
// value, metadata (like source, timestamp, and sensitivity), and related audit
// information for tracking changes.
type PropertyValue struct {
	// Name of custom property
	Name                     string `json:"name,required"`
	SourceUpstreamDeployable string `json:"sourceUpstreamDeployable,required"`
	// Custom property value
	Value string `json:"value,required"`
	// The sensitivity level of the property, such as "non_sensitive", "sensitive", and
	// "highly_sensitive".
	//
	// Any of "none", "standard", "high".
	DataSensitivity PropertyValueDataSensitivity `json:"dataSensitivity"`
	// Whether the property value is encrypted.
	IsEncrypted          bool  `json:"isEncrypted"`
	IsLargeValue         bool  `json:"isLargeValue"`
	PersistenceTimestamp int64 `json:"persistenceTimestamp"`
	// A unique ID associated with this request.
	RequestID string `json:"requestId"`
	// Whether the value was selected by a user.
	SelectedByUser bool `json:"selectedByUser"`
	// The timestamp when the value was selected by a user, if applicable.
	SelectedByUserTimestamp int64 `json:"selectedByUserTimestamp"`
	// The origin of the property value, such as "IMPORT" or "API".
	//
	// Any of "UNKNOWN", "IMPORT", "API", "FORM", "ANALYTICS", "MIGRATION",
	// "SALESFORCE", "INTEGRATION", "CONTACTS_WEB", "WAL_INCREMENTAL", "TASK", "EMAIL",
	// "WORKFLOWS", "CALCULATED", "SOCIAL", "BATCH_UPDATE", "SIGNALS", "BIDEN",
	// "DEFAULT", "COMPANIES", "DEALS", "ASSISTS", "PRESENTATIONS", "TALLY",
	// "SIDEKICK", "CRM_UI", "MERGE_CONTACTS", "PORTAL_USER_ASSOCIATOR",
	// "INTEGRATIONS_PLATFORM", "BCC_TO_CRM", "FORWARD_TO_CRM", "ENGAGEMENTS", "SALES",
	// "HEISENBERG", "LEADIN", "GMAIL_INTEGRATION", "ACADEMY", "SALES_MESSAGES",
	// "AVATARS_SERVICE", "MERGE_COMPANIES", "SEQUENCES", "COMPANY_FAMILIES",
	// "MOBILE_IOS", "MOBILE_ANDROID", "CONTACTS", "ASSOCIATIONS", "EXTENSION",
	// "SUCCESS", "BOT", "INTEGRATIONS_SYNC", "AUTOMATION_PLATFORM", "CONVERSATIONS",
	// "EMAIL_INTEGRATION", "CONTENT_MEMBERSHIP", "QUOTES", "BET_ASSIGNMENT", "QUOTAS",
	// "BET_CRM_CONNECTOR", "MEETINGS", "MERGE_OBJECTS", "RECYCLING_BIN", "ADS",
	// "AI_GROUP", "COMMUNICATOR", "SETTINGS", "PROPERTY_SETTINGS",
	// "PIPELINE_SETTINGS", "COMPANY_INSIGHTS", "BEHAVIORAL_EVENTS", "PAYMENTS",
	// "GOALS", "PORTAL_OBJECT_SYNC", "APPROVALS", "FILE_MANAGER", "MARKETPLACE",
	// "INTERNAL_PROCESSING", "FORECASTING", "SLACK_INTEGRATION", "CRM_UI_BULK_ACTION",
	// "WORKFLOW_CONTACT_DELETE_ACTION", "ACCEPTANCE_TEST", "PLAYBOOKS", "CHATSPOT",
	// "FLYWHEEL_PRODUCT_DATA_SYNC", "HELP_DESK", "BILLING", "DATA_ENRICHMENT",
	// "AUTOMATION_JOURNEY", "MICROAPPS", "INTENT", "PROSPECTING_AGENT",
	// "CENTRAL_EXCHANGE_RATES", "HELP_DESK_AI", "CONVERSATIONAL_ENRICHMENT",
	// "CRM_PROCESSES_PLATFORM", "CLONE_OBJECTS", "MARKET_SOURCING", "DATASET",
	// "PROPERTY_RESTORE".
	Source PropertyValueSource `json:"source"`
	// The ID of the property source indicating where it was created.
	SourceID string `json:"sourceId"`
	// A human-readable label.
	SourceLabel string `json:"sourceLabel"`
	// Source metadata encoded as a base64 string. For example: `ZXhhbXBsZSBzdHJpbmc=`
	SourceMetadata string `json:"sourceMetadata"`
	// The unique identifier associated with the source.
	SourceVid []int64 `json:"sourceVid"`
	// When the value was set, as a 64-bit integer.
	Timestamp int64 `json:"timestamp"`
	// The unit of measurement or context for the value.
	Unit string `json:"unit"`
	// The ID of the user who updated the property.
	UpdatedByUserID                    int64 `json:"updatedByUserId"`
	UseTimestampAsPersistenceTimestamp bool  `json:"useTimestampAsPersistenceTimestamp"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name                               respjson.Field
		SourceUpstreamDeployable           respjson.Field
		Value                              respjson.Field
		DataSensitivity                    respjson.Field
		IsEncrypted                        respjson.Field
		IsLargeValue                       respjson.Field
		PersistenceTimestamp               respjson.Field
		RequestID                          respjson.Field
		SelectedByUser                     respjson.Field
		SelectedByUserTimestamp            respjson.Field
		Source                             respjson.Field
		SourceID                           respjson.Field
		SourceLabel                        respjson.Field
		SourceMetadata                     respjson.Field
		SourceVid                          respjson.Field
		Timestamp                          respjson.Field
		Unit                               respjson.Field
		UpdatedByUserID                    respjson.Field
		UseTimestampAsPersistenceTimestamp respjson.Field
		ExtraFields                        map[string]respjson.Field
		raw                                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PropertyValue) RawJSON() string { return r.JSON.raw }
func (r *PropertyValue) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this PropertyValue to a PropertyValueParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// PropertyValueParam.Overrides()
func (r PropertyValue) ToParam() PropertyValueParam {
	return param.Override[PropertyValueParam](json.RawMessage(r.RawJSON()))
}

// The sensitivity level of the property, such as "non_sensitive", "sensitive", and
// "highly_sensitive".
type PropertyValueDataSensitivity string

const (
	PropertyValueDataSensitivityNone     PropertyValueDataSensitivity = "none"
	PropertyValueDataSensitivityStandard PropertyValueDataSensitivity = "standard"
	PropertyValueDataSensitivityHigh     PropertyValueDataSensitivity = "high"
)

// The origin of the property value, such as "IMPORT" or "API".
type PropertyValueSource string

const (
	PropertyValueSourceUnknown                     PropertyValueSource = "UNKNOWN"
	PropertyValueSourceImport                      PropertyValueSource = "IMPORT"
	PropertyValueSourceAPI                         PropertyValueSource = "API"
	PropertyValueSourceForm                        PropertyValueSource = "FORM"
	PropertyValueSourceAnalytics                   PropertyValueSource = "ANALYTICS"
	PropertyValueSourceMigration                   PropertyValueSource = "MIGRATION"
	PropertyValueSourceSalesforce                  PropertyValueSource = "SALESFORCE"
	PropertyValueSourceIntegration                 PropertyValueSource = "INTEGRATION"
	PropertyValueSourceContactsWeb                 PropertyValueSource = "CONTACTS_WEB"
	PropertyValueSourceWalIncremental              PropertyValueSource = "WAL_INCREMENTAL"
	PropertyValueSourceTask                        PropertyValueSource = "TASK"
	PropertyValueSourceEmail                       PropertyValueSource = "EMAIL"
	PropertyValueSourceWorkflows                   PropertyValueSource = "WORKFLOWS"
	PropertyValueSourceCalculated                  PropertyValueSource = "CALCULATED"
	PropertyValueSourceSocial                      PropertyValueSource = "SOCIAL"
	PropertyValueSourceBatchUpdate                 PropertyValueSource = "BATCH_UPDATE"
	PropertyValueSourceSignals                     PropertyValueSource = "SIGNALS"
	PropertyValueSourceBiden                       PropertyValueSource = "BIDEN"
	PropertyValueSourceDefault                     PropertyValueSource = "DEFAULT"
	PropertyValueSourceCompanies                   PropertyValueSource = "COMPANIES"
	PropertyValueSourceDeals                       PropertyValueSource = "DEALS"
	PropertyValueSourceAssists                     PropertyValueSource = "ASSISTS"
	PropertyValueSourcePresentations               PropertyValueSource = "PRESENTATIONS"
	PropertyValueSourceTally                       PropertyValueSource = "TALLY"
	PropertyValueSourceSidekick                    PropertyValueSource = "SIDEKICK"
	PropertyValueSourceCRMUi                       PropertyValueSource = "CRM_UI"
	PropertyValueSourceMergeContacts               PropertyValueSource = "MERGE_CONTACTS"
	PropertyValueSourcePortalUserAssociator        PropertyValueSource = "PORTAL_USER_ASSOCIATOR"
	PropertyValueSourceIntegrationsPlatform        PropertyValueSource = "INTEGRATIONS_PLATFORM"
	PropertyValueSourceBccToCRM                    PropertyValueSource = "BCC_TO_CRM"
	PropertyValueSourceForwardToCRM                PropertyValueSource = "FORWARD_TO_CRM"
	PropertyValueSourceEngagements                 PropertyValueSource = "ENGAGEMENTS"
	PropertyValueSourceSales                       PropertyValueSource = "SALES"
	PropertyValueSourceHeisenberg                  PropertyValueSource = "HEISENBERG"
	PropertyValueSourceLeadin                      PropertyValueSource = "LEADIN"
	PropertyValueSourceGmailIntegration            PropertyValueSource = "GMAIL_INTEGRATION"
	PropertyValueSourceAcademy                     PropertyValueSource = "ACADEMY"
	PropertyValueSourceSalesMessages               PropertyValueSource = "SALES_MESSAGES"
	PropertyValueSourceAvatarsService              PropertyValueSource = "AVATARS_SERVICE"
	PropertyValueSourceMergeCompanies              PropertyValueSource = "MERGE_COMPANIES"
	PropertyValueSourceSequences                   PropertyValueSource = "SEQUENCES"
	PropertyValueSourceCompanyFamilies             PropertyValueSource = "COMPANY_FAMILIES"
	PropertyValueSourceMobileIos                   PropertyValueSource = "MOBILE_IOS"
	PropertyValueSourceMobileAndroid               PropertyValueSource = "MOBILE_ANDROID"
	PropertyValueSourceContacts                    PropertyValueSource = "CONTACTS"
	PropertyValueSourceAssociations                PropertyValueSource = "ASSOCIATIONS"
	PropertyValueSourceExtension                   PropertyValueSource = "EXTENSION"
	PropertyValueSourceSuccess                     PropertyValueSource = "SUCCESS"
	PropertyValueSourceBot                         PropertyValueSource = "BOT"
	PropertyValueSourceIntegrationsSync            PropertyValueSource = "INTEGRATIONS_SYNC"
	PropertyValueSourceAutomationPlatform          PropertyValueSource = "AUTOMATION_PLATFORM"
	PropertyValueSourceConversations               PropertyValueSource = "CONVERSATIONS"
	PropertyValueSourceEmailIntegration            PropertyValueSource = "EMAIL_INTEGRATION"
	PropertyValueSourceContentMembership           PropertyValueSource = "CONTENT_MEMBERSHIP"
	PropertyValueSourceQuotes                      PropertyValueSource = "QUOTES"
	PropertyValueSourceBetAssignment               PropertyValueSource = "BET_ASSIGNMENT"
	PropertyValueSourceQuotas                      PropertyValueSource = "QUOTAS"
	PropertyValueSourceBetCRMConnector             PropertyValueSource = "BET_CRM_CONNECTOR"
	PropertyValueSourceMeetings                    PropertyValueSource = "MEETINGS"
	PropertyValueSourceMergeObjects                PropertyValueSource = "MERGE_OBJECTS"
	PropertyValueSourceRecyclingBin                PropertyValueSource = "RECYCLING_BIN"
	PropertyValueSourceAds                         PropertyValueSource = "ADS"
	PropertyValueSourceAIGroup                     PropertyValueSource = "AI_GROUP"
	PropertyValueSourceCommunicator                PropertyValueSource = "COMMUNICATOR"
	PropertyValueSourceSettings                    PropertyValueSource = "SETTINGS"
	PropertyValueSourcePropertySettings            PropertyValueSource = "PROPERTY_SETTINGS"
	PropertyValueSourcePipelineSettings            PropertyValueSource = "PIPELINE_SETTINGS"
	PropertyValueSourceCompanyInsights             PropertyValueSource = "COMPANY_INSIGHTS"
	PropertyValueSourceBehavioralEvents            PropertyValueSource = "BEHAVIORAL_EVENTS"
	PropertyValueSourcePayments                    PropertyValueSource = "PAYMENTS"
	PropertyValueSourceGoals                       PropertyValueSource = "GOALS"
	PropertyValueSourcePortalObjectSync            PropertyValueSource = "PORTAL_OBJECT_SYNC"
	PropertyValueSourceApprovals                   PropertyValueSource = "APPROVALS"
	PropertyValueSourceFileManager                 PropertyValueSource = "FILE_MANAGER"
	PropertyValueSourceMarketplace                 PropertyValueSource = "MARKETPLACE"
	PropertyValueSourceInternalProcessing          PropertyValueSource = "INTERNAL_PROCESSING"
	PropertyValueSourceForecasting                 PropertyValueSource = "FORECASTING"
	PropertyValueSourceSlackIntegration            PropertyValueSource = "SLACK_INTEGRATION"
	PropertyValueSourceCRMUiBulkAction             PropertyValueSource = "CRM_UI_BULK_ACTION"
	PropertyValueSourceWorkflowContactDeleteAction PropertyValueSource = "WORKFLOW_CONTACT_DELETE_ACTION"
	PropertyValueSourceAcceptanceTest              PropertyValueSource = "ACCEPTANCE_TEST"
	PropertyValueSourcePlaybooks                   PropertyValueSource = "PLAYBOOKS"
	PropertyValueSourceChatspot                    PropertyValueSource = "CHATSPOT"
	PropertyValueSourceFlywheelProductDataSync     PropertyValueSource = "FLYWHEEL_PRODUCT_DATA_SYNC"
	PropertyValueSourceHelpDesk                    PropertyValueSource = "HELP_DESK"
	PropertyValueSourceBilling                     PropertyValueSource = "BILLING"
	PropertyValueSourceDataEnrichment              PropertyValueSource = "DATA_ENRICHMENT"
	PropertyValueSourceAutomationJourney           PropertyValueSource = "AUTOMATION_JOURNEY"
	PropertyValueSourceMicroapps                   PropertyValueSource = "MICROAPPS"
	PropertyValueSourceIntent                      PropertyValueSource = "INTENT"
	PropertyValueSourceProspectingAgent            PropertyValueSource = "PROSPECTING_AGENT"
	PropertyValueSourceCentralExchangeRates        PropertyValueSource = "CENTRAL_EXCHANGE_RATES"
	PropertyValueSourceHelpDeskAI                  PropertyValueSource = "HELP_DESK_AI"
	PropertyValueSourceConversationalEnrichment    PropertyValueSource = "CONVERSATIONAL_ENRICHMENT"
	PropertyValueSourceCRMProcessesPlatform        PropertyValueSource = "CRM_PROCESSES_PLATFORM"
	PropertyValueSourceCloneObjects                PropertyValueSource = "CLONE_OBJECTS"
	PropertyValueSourceMarketSourcing              PropertyValueSource = "MARKET_SOURCING"
	PropertyValueSourceDataset                     PropertyValueSource = "DATASET"
	PropertyValueSourcePropertyRestore             PropertyValueSource = "PROPERTY_RESTORE"
)

// Represents a single custom property of a marketing event, storing its name,
// value, metadata (like source, timestamp, and sensitivity), and related audit
// information for tracking changes.
//
// The properties Name, SourceUpstreamDeployable, Value are required.
type PropertyValueParam struct {
	// Name of custom property
	Name                     string `json:"name,required"`
	SourceUpstreamDeployable string `json:"sourceUpstreamDeployable,required"`
	// Custom property value
	Value string `json:"value,required"`
	// Whether the property value is encrypted.
	IsEncrypted          param.Opt[bool]  `json:"isEncrypted,omitzero"`
	IsLargeValue         param.Opt[bool]  `json:"isLargeValue,omitzero"`
	PersistenceTimestamp param.Opt[int64] `json:"persistenceTimestamp,omitzero"`
	// A unique ID associated with this request.
	RequestID param.Opt[string] `json:"requestId,omitzero"`
	// Whether the value was selected by a user.
	SelectedByUser param.Opt[bool] `json:"selectedByUser,omitzero"`
	// The timestamp when the value was selected by a user, if applicable.
	SelectedByUserTimestamp param.Opt[int64] `json:"selectedByUserTimestamp,omitzero"`
	// The ID of the property source indicating where it was created.
	SourceID param.Opt[string] `json:"sourceId,omitzero"`
	// A human-readable label.
	SourceLabel param.Opt[string] `json:"sourceLabel,omitzero"`
	// Source metadata encoded as a base64 string. For example: `ZXhhbXBsZSBzdHJpbmc=`
	SourceMetadata param.Opt[string] `json:"sourceMetadata,omitzero"`
	// When the value was set, as a 64-bit integer.
	Timestamp param.Opt[int64] `json:"timestamp,omitzero"`
	// The unit of measurement or context for the value.
	Unit param.Opt[string] `json:"unit,omitzero"`
	// The ID of the user who updated the property.
	UpdatedByUserID                    param.Opt[int64] `json:"updatedByUserId,omitzero"`
	UseTimestampAsPersistenceTimestamp param.Opt[bool]  `json:"useTimestampAsPersistenceTimestamp,omitzero"`
	// The sensitivity level of the property, such as "non_sensitive", "sensitive", and
	// "highly_sensitive".
	//
	// Any of "none", "standard", "high".
	DataSensitivity PropertyValueDataSensitivity `json:"dataSensitivity,omitzero"`
	// The origin of the property value, such as "IMPORT" or "API".
	//
	// Any of "UNKNOWN", "IMPORT", "API", "FORM", "ANALYTICS", "MIGRATION",
	// "SALESFORCE", "INTEGRATION", "CONTACTS_WEB", "WAL_INCREMENTAL", "TASK", "EMAIL",
	// "WORKFLOWS", "CALCULATED", "SOCIAL", "BATCH_UPDATE", "SIGNALS", "BIDEN",
	// "DEFAULT", "COMPANIES", "DEALS", "ASSISTS", "PRESENTATIONS", "TALLY",
	// "SIDEKICK", "CRM_UI", "MERGE_CONTACTS", "PORTAL_USER_ASSOCIATOR",
	// "INTEGRATIONS_PLATFORM", "BCC_TO_CRM", "FORWARD_TO_CRM", "ENGAGEMENTS", "SALES",
	// "HEISENBERG", "LEADIN", "GMAIL_INTEGRATION", "ACADEMY", "SALES_MESSAGES",
	// "AVATARS_SERVICE", "MERGE_COMPANIES", "SEQUENCES", "COMPANY_FAMILIES",
	// "MOBILE_IOS", "MOBILE_ANDROID", "CONTACTS", "ASSOCIATIONS", "EXTENSION",
	// "SUCCESS", "BOT", "INTEGRATIONS_SYNC", "AUTOMATION_PLATFORM", "CONVERSATIONS",
	// "EMAIL_INTEGRATION", "CONTENT_MEMBERSHIP", "QUOTES", "BET_ASSIGNMENT", "QUOTAS",
	// "BET_CRM_CONNECTOR", "MEETINGS", "MERGE_OBJECTS", "RECYCLING_BIN", "ADS",
	// "AI_GROUP", "COMMUNICATOR", "SETTINGS", "PROPERTY_SETTINGS",
	// "PIPELINE_SETTINGS", "COMPANY_INSIGHTS", "BEHAVIORAL_EVENTS", "PAYMENTS",
	// "GOALS", "PORTAL_OBJECT_SYNC", "APPROVALS", "FILE_MANAGER", "MARKETPLACE",
	// "INTERNAL_PROCESSING", "FORECASTING", "SLACK_INTEGRATION", "CRM_UI_BULK_ACTION",
	// "WORKFLOW_CONTACT_DELETE_ACTION", "ACCEPTANCE_TEST", "PLAYBOOKS", "CHATSPOT",
	// "FLYWHEEL_PRODUCT_DATA_SYNC", "HELP_DESK", "BILLING", "DATA_ENRICHMENT",
	// "AUTOMATION_JOURNEY", "MICROAPPS", "INTENT", "PROSPECTING_AGENT",
	// "CENTRAL_EXCHANGE_RATES", "HELP_DESK_AI", "CONVERSATIONAL_ENRICHMENT",
	// "CRM_PROCESSES_PLATFORM", "CLONE_OBJECTS", "MARKET_SOURCING", "DATASET",
	// "PROPERTY_RESTORE".
	Source PropertyValueSource `json:"source,omitzero"`
	// The unique identifier associated with the source.
	SourceVid []int64 `json:"sourceVid,omitzero"`
	paramObj
}

func (r PropertyValueParam) MarshalJSON() (data []byte, err error) {
	type shadow PropertyValueParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PropertyValueParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PublicList struct {
	ListID           string    `json:"listId,required"`
	ListVersion      int64     `json:"listVersion,required"`
	Name             string    `json:"name,required"`
	ObjectTypeID     string    `json:"objectTypeId,required"`
	ProcessingStatus string    `json:"processingStatus,required"`
	ProcessingType   string    `json:"processingType,required"`
	CreatedAt        time.Time `json:"createdAt" format:"date-time"`
	CreatedByID      string    `json:"createdById"`
	DeletedAt        time.Time `json:"deletedAt" format:"date-time"`
	FiltersUpdatedAt time.Time `json:"filtersUpdatedAt" format:"date-time"`
	Size             int64     `json:"size"`
	UpdatedAt        time.Time `json:"updatedAt" format:"date-time"`
	UpdatedByID      string    `json:"updatedById"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ListID           respjson.Field
		ListVersion      respjson.Field
		Name             respjson.Field
		ObjectTypeID     respjson.Field
		ProcessingStatus respjson.Field
		ProcessingType   respjson.Field
		CreatedAt        respjson.Field
		CreatedByID      respjson.Field
		DeletedAt        respjson.Field
		FiltersUpdatedAt respjson.Field
		Size             respjson.Field
		UpdatedAt        respjson.Field
		UpdatedByID      respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PublicList) RawJSON() string { return r.JSON.raw }
func (r *PublicList) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SearchPublicResponseWrapper struct {
	AppID             int64  `json:"appId,required"`
	ExternalAccountID string `json:"externalAccountId,required"`
	ExternalEventID   string `json:"externalEventId,required"`
	ObjectID          string `json:"objectId,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AppID             respjson.Field
		ExternalAccountID respjson.Field
		ExternalEventID   respjson.Field
		ObjectID          respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SearchPublicResponseWrapper) RawJSON() string { return r.JSON.raw }
func (r *SearchPublicResponseWrapper) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SubscriberEmailResponse struct {
	Email string `json:"email,required"`
	Vid   int64  `json:"vid,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Email       respjson.Field
		Vid         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SubscriberEmailResponse) RawJSON() string { return r.JSON.raw }
func (r *SubscriberEmailResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SubscriberVidResponse struct {
	Vid int64 `json:"vid,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Vid         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SubscriberVidResponse) RawJSON() string { return r.JSON.raw }
func (r *SubscriberVidResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EventNewParams struct {
	MarketingEventCreateRequestParams MarketingEventCreateRequestParams
	paramObj
}

func (r EventNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.MarketingEventCreateRequestParams)
}
func (r *EventNewParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.MarketingEventCreateRequestParams)
}

type EventUpdateParams struct {
	MarketingEventPublicUpdateRequestV2 MarketingEventPublicUpdateRequestV2Param
	paramObj
}

func (r EventUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.MarketingEventPublicUpdateRequestV2)
}
func (r *EventUpdateParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.MarketingEventPublicUpdateRequestV2)
}

type EventListParams struct {
	// The cursor indicating the position of the last retrieved item.
	After param.Opt[string] `query:"after,omitzero" json:"-"`
	// The limit for response size. The default value is 10, the max number is 100
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [EventListParams]'s query parameters as `url.Values`.
func (r EventListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type EventCancelByExternalEventIDParams struct {
	// The accountId that is associated with this marketing event in the external event
	// application
	ExternalAccountID string `query:"externalAccountId,required" json:"-"`
	paramObj
}

// URLQuery serializes [EventCancelByExternalEventIDParams]'s query parameters as
// `url.Values`.
func (r EventCancelByExternalEventIDParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type EventCompleteByExternalEventIDParams struct {
	// The accountId that is associated with this marketing event in the external event
	// application.
	ExternalAccountID                   string `query:"externalAccountId,required" json:"-"`
	MarketingEventCompleteRequestParams MarketingEventCompleteRequestParams
	paramObj
}

func (r EventCompleteByExternalEventIDParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.MarketingEventCompleteRequestParams)
}
func (r *EventCompleteByExternalEventIDParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.MarketingEventCompleteRequestParams)
}

// URLQuery serializes [EventCompleteByExternalEventIDParams]'s query parameters as
// `url.Values`.
func (r EventCompleteByExternalEventIDParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type EventDeleteBatchParams struct {
	BatchInputMarketingEventPublicObjectIDDeleteRequest BatchInputMarketingEventPublicObjectIDDeleteRequestParam
	paramObj
}

func (r EventDeleteBatchParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.BatchInputMarketingEventPublicObjectIDDeleteRequest)
}
func (r *EventDeleteBatchParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.BatchInputMarketingEventPublicObjectIDDeleteRequest)
}

type EventDeleteBatchByExternalEventIDParams struct {
	BatchInputMarketingEventExternalUniqueIdentifier BatchInputMarketingEventExternalUniqueIdentifierParam
	paramObj
}

func (r EventDeleteBatchByExternalEventIDParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.BatchInputMarketingEventExternalUniqueIdentifier)
}
func (r *EventDeleteBatchByExternalEventIDParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.BatchInputMarketingEventExternalUniqueIdentifier)
}

type EventDeleteByExternalEventIDParams struct {
	// The accountId that is associated with this marketing event in the external event
	// application
	ExternalAccountID string `query:"externalAccountId,required" json:"-"`
	paramObj
}

// URLQuery serializes [EventDeleteByExternalEventIDParams]'s query parameters as
// `url.Values`.
func (r EventDeleteByExternalEventIDParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type EventGetByExternalEventIDParams struct {
	// The accountId that is associated with this marketing event in the external event
	// application
	ExternalAccountID string `query:"externalAccountId,required" json:"-"`
	paramObj
}

// URLQuery serializes [EventGetByExternalEventIDParams]'s query parameters as
// `url.Values`.
func (r EventGetByExternalEventIDParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type EventSearchByExternalEventIDParams struct {
	// The id of the marketing event in the external event application
	// (externalEventId)
	Q string `query:"q,required" json:"-"`
	paramObj
}

// URLQuery serializes [EventSearchByExternalEventIDParams]'s query parameters as
// `url.Values`.
func (r EventSearchByExternalEventIDParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type EventUpdateBatchParams struct {
	BatchInputMarketingEventPublicUpdateRequestFullV2 BatchInputMarketingEventPublicUpdateRequestFullV2Param
	paramObj
}

func (r EventUpdateBatchParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.BatchInputMarketingEventPublicUpdateRequestFullV2)
}
func (r *EventUpdateBatchParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.BatchInputMarketingEventPublicUpdateRequestFullV2)
}

type EventUpdateByExternalEventIDParams struct {
	// The accountId that is associated with this marketing event in the external event
	// application
	ExternalAccountID                 string `query:"externalAccountId,required" json:"-"`
	MarketingEventUpdateRequestParams MarketingEventUpdateRequestParams
	paramObj
}

func (r EventUpdateByExternalEventIDParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.MarketingEventUpdateRequestParams)
}
func (r *EventUpdateByExternalEventIDParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.MarketingEventUpdateRequestParams)
}

// URLQuery serializes [EventUpdateByExternalEventIDParams]'s query parameters as
// `url.Values`.
func (r EventUpdateByExternalEventIDParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type EventUpsertBatchParams struct {
	BatchInputMarketingEventCreateRequestParams BatchInputMarketingEventCreateRequestParams
	paramObj
}

func (r EventUpsertBatchParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.BatchInputMarketingEventCreateRequestParams)
}
func (r *EventUpsertBatchParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.BatchInputMarketingEventCreateRequestParams)
}

type EventUpsertByExternalEventIDParams struct {
	MarketingEventCreateRequestParams MarketingEventCreateRequestParams
	paramObj
}

func (r EventUpsertByExternalEventIDParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.MarketingEventCreateRequestParams)
}
func (r *EventUpsertByExternalEventIDParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.MarketingEventCreateRequestParams)
}

type EventUpsertSubscriberStateByEmailParams struct {
	ExternalEventID string `path:"externalEventId,required" json:"-"`
	// The accountId that is associated with this marketing event in the external event
	// application
	ExternalAccountID                       string `query:"externalAccountId,required" json:"-"`
	BatchInputMarketingEventEmailSubscriber BatchInputMarketingEventEmailSubscriberParam
	paramObj
}

func (r EventUpsertSubscriberStateByEmailParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.BatchInputMarketingEventEmailSubscriber)
}
func (r *EventUpsertSubscriberStateByEmailParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.BatchInputMarketingEventEmailSubscriber)
}

// URLQuery serializes [EventUpsertSubscriberStateByEmailParams]'s query parameters
// as `url.Values`.
func (r EventUpsertSubscriberStateByEmailParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type EventUpsertSubscriberStateByIDParams struct {
	ExternalEventID string `path:"externalEventId,required" json:"-"`
	// The accountId that is associated with this marketing event in the external event
	// application
	ExternalAccountID                  string `query:"externalAccountId,required" json:"-"`
	BatchInputMarketingEventSubscriber BatchInputMarketingEventSubscriberParam
	paramObj
}

func (r EventUpsertSubscriberStateByIDParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.BatchInputMarketingEventSubscriber)
}
func (r *EventUpsertSubscriberStateByIDParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.BatchInputMarketingEventSubscriber)
}

// URLQuery serializes [EventUpsertSubscriberStateByIDParams]'s query parameters as
// `url.Values`.
func (r EventUpsertSubscriberStateByIDParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
