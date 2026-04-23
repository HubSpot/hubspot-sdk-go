// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package marketing

import (
	"context"
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

// MarketingEventService contains methods and other services that help with
// interacting with the hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMarketingEventService] method instead.
type MarketingEventService struct {
	options          []option.RequestOption
	Attendance       MarketingEventAttendanceService
	Events           MarketingEventEventService
	ListAssociations MarketingEventListAssociationService
	Participations   MarketingEventParticipationService
	Settings         MarketingEventSettingService
	SubscriberState  MarketingEventSubscriberStateService
}

// NewMarketingEventService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewMarketingEventService(opts ...option.RequestOption) (r MarketingEventService) {
	r = MarketingEventService{}
	r.options = opts
	r.Attendance = NewMarketingEventAttendanceService(opts...)
	r.Events = NewMarketingEventEventService(opts...)
	r.ListAssociations = NewMarketingEventListAssociationService(opts...)
	r.Participations = NewMarketingEventParticipationService(opts...)
	r.Settings = NewMarketingEventSettingService(opts...)
	r.SubscriberState = NewMarketingEventSubscriberStateService(opts...)
	return
}

// Creates a new marketing event in HubSpot
func (r *MarketingEventService) New(ctx context.Context, body MarketingEventNewParams, opts ...option.RequestOption) (res *MarketingEventDefaultResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "marketing/marketing-events/2026-03/events"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Updates the details of an existing Marketing Event identified by its objectId,
// if it exists.
func (r *MarketingEventService) Update(ctx context.Context, objectID string, body MarketingEventUpdateParams, opts ...option.RequestOption) (res *MarketingEventPublicDefaultResponseV2, err error) {
	opts = slices.Concat(r.options, opts)
	if objectID == "" {
		err = errors.New("missing required objectId parameter")
		return nil, err
	}
	path := fmt.Sprintf("marketing/marketing-events/2026-03/%s", url.PathEscape(objectID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

func (r *MarketingEventService) List(ctx context.Context, query MarketingEventListParams, opts ...option.RequestOption) (res *pagination.Page[MarketingEventPublicReadResponseV2], err error) {
	var raw *http.Response
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "marketing/marketing-events/2026-03"
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

func (r *MarketingEventService) ListAutoPaging(ctx context.Context, query MarketingEventListParams, opts ...option.RequestOption) *pagination.PageAutoPager[MarketingEventPublicReadResponseV2] {
	return pagination.NewPageAutoPager(r.List(ctx, query, opts...))
}

// Deletes the existing Marketing Event with the specified objectId, if it exists.
func (r *MarketingEventService) Delete(ctx context.Context, objectID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if objectID == "" {
		err = errors.New("missing required objectId parameter")
		return err
	}
	path := fmt.Sprintf("marketing/marketing-events/2026-03/%s", url.PathEscape(objectID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Deletes multiple Marketing Events from the portal based on their objectId, if
// they exist.
//
// Responses: 204: Returned if all specified Marketing Events were successfully
// deleted. 207: Returned if some objectIds did not correspond to any existing
// Marketing Events.
func (r *MarketingEventService) DeleteBatch(ctx context.Context, body MarketingEventDeleteBatchParams, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "marketing/marketing-events/2026-03/batch/archive"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Deletes multiple Marketing Events based on externalAccountId, externalEventId,
// and appId.
//
// Only Marketing Events created by the same apps will be deleted; events from
// other apps cannot be removed by this endpoint.
func (r *MarketingEventService) DeleteBatchByExternalEventID(ctx context.Context, body MarketingEventDeleteBatchByExternalEventIDParams, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "marketing/marketing-events/2026-03/events/delete"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Deletes the existing Marketing Event with the specified externalAccountId,
// externalEventId, if it exists.
//
// Only Marketing Events created by the same app can be deleted.
func (r *MarketingEventService) DeleteByExternalEventID(ctx context.Context, externalEventID string, body MarketingEventDeleteByExternalEventIDParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if externalEventID == "" {
		err = errors.New("missing required externalEventId parameter")
		return err
	}
	path := fmt.Sprintf("marketing/marketing-events/2026-03/events/%s", url.PathEscape(externalEventID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, nil, opts...)
	return err
}

// Returns the details of a Marketing Event with the specified objectId, if it
// exists.
func (r *MarketingEventService) Get(ctx context.Context, objectID string, opts ...option.RequestOption) (res *MarketingEventPublicReadResponseV2, err error) {
	opts = slices.Concat(r.options, opts)
	if objectID == "" {
		err = errors.New("missing required objectId parameter")
		return nil, err
	}
	path := fmt.Sprintf("marketing/marketing-events/2026-03/%s", url.PathEscape(objectID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Returns the details of a Marketing Event with the specified externalAccountId,
// externalEventId, if it exists.
//
// Only Marketing Events created by the same app making the request can be
// retrieved.
func (r *MarketingEventService) GetByExternalEventID(ctx context.Context, externalEventID string, query MarketingEventGetByExternalEventIDParams, opts ...option.RequestOption) (res *MarketingEventPublicReadResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if externalEventID == "" {
		err = errors.New("missing required externalEventId parameter")
		return nil, err
	}
	path := fmt.Sprintf("marketing/marketing-events/2026-03/events/%s", url.PathEscape(externalEventID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Retrieves Marketing Events where the externalEventId matches the value provided
// in the request, limited to events created by the app making the request.
//
// Marketing Events created by other apps will not be included in the results.
func (r *MarketingEventService) SearchByExternalEventID(ctx context.Context, query MarketingEventSearchByExternalEventIDParams, opts ...option.RequestOption) (res *CollectionResponseSearchPublicResponseWrapperNoPaging, err error) {
	opts = slices.Concat(r.options, opts)
	path := "marketing/marketing-events/2026-03/events/search"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
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
func (r *MarketingEventService) SearchIdentifiersByExternalEventID(ctx context.Context, externalEventID string, opts ...option.RequestOption) (res *CollectionResponseWithTotalMarketingEventIdentifiersResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if externalEventID == "" {
		err = errors.New("missing required externalEventId parameter")
		return nil, err
	}
	path := fmt.Sprintf("marketing/marketing-events/2026-03/%s/identifiers", url.PathEscape(externalEventID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Updates multiple Marketing Events on the portal based on their objectId, if they
// exist.
func (r *MarketingEventService) UpdateBatch(ctx context.Context, body MarketingEventUpdateBatchParams, opts ...option.RequestOption) (res *BatchResponseMarketingEventPublicDefaultResponseV2, err error) {
	opts = slices.Concat(r.options, opts)
	path := "marketing/marketing-events/2026-03/batch/update"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Updates the details of an existing Marketing Event identified by its
// externalAccountId, externalEventId if it exists.
//
// Only Marketing Events created by the same app can be updated.
func (r *MarketingEventService) UpdateByExternalEventID(ctx context.Context, externalEventID string, params MarketingEventUpdateByExternalEventIDParams, opts ...option.RequestOption) (res *MarketingEventPublicDefaultResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if externalEventID == "" {
		err = errors.New("missing required externalEventId parameter")
		return nil, err
	}
	path := fmt.Sprintf("marketing/marketing-events/2026-03/events/%s", url.PathEscape(externalEventID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return res, err
}

// Upserts multiple Marketing Events. If a Marketing Event with the specified ID
// already exists, it will be updated; otherwise, a new event will be created.
//
// Only Marketing Events originally created by the same app can be updated.
func (r *MarketingEventService) UpsertBatch(ctx context.Context, body MarketingEventUpsertBatchParams, opts ...option.RequestOption) (res *BatchResponseMarketingEventPublicDefaultResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "marketing/marketing-events/2026-03/events/upsert"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Upserts a marketing event If there is an existing marketing event with the
// specified ID, it will be updated; otherwise a new event will be created.
func (r *MarketingEventService) UpsertByExternalEventID(ctx context.Context, externalEventID string, body MarketingEventUpsertByExternalEventIDParams, opts ...option.RequestOption) (res *MarketingEventPublicDefaultResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if externalEventID == "" {
		err = errors.New("missing required externalEventId parameter")
		return nil, err
	}
	path := fmt.Sprintf("marketing/marketing-events/2026-03/events/%s", url.PathEscape(externalEventID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return res, err
}

type AppInfo struct {
	// The ID of the application
	ID string `json:"id" api:"required"`
	// The name of the application
	Name string `json:"name" api:"required"`
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
	// Number of attended contact records of a marketing event
	Attended int64 `json:"attended" api:"required"`
	// Number of cancelled contact records of a marketing event
	Cancelled int64 `json:"cancelled" api:"required"`
	// Number of no-show contact records of a marketing event
	NoShows int64 `json:"noShows" api:"required"`
	// Number of registered contact records of a marketing event
	Registered int64 `json:"registered" api:"required"`
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
	Inputs []MarketingEventCreateRequestParams `json:"inputs,omitzero" api:"required"`
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
	Inputs []MarketingEventEmailSubscriberParam `json:"inputs,omitzero" api:"required"`
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
	Inputs []MarketingEventExternalUniqueIdentifierParam `json:"inputs,omitzero" api:"required"`
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
	Inputs []MarketingEventPublicObjectIDDeleteRequestParam `json:"inputs,omitzero" api:"required"`
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
	Inputs []MarketingEventPublicUpdateRequestFullV2Param `json:"inputs,omitzero" api:"required"`
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
	Inputs []MarketingEventSubscriberParam `json:"inputs,omitzero" api:"required"`
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
	// Timestamp of when the request was completed.
	CompletedAt time.Time                             `json:"completedAt" api:"required" format:"date-time"`
	Results     []MarketingEventPublicDefaultResponse `json:"results" api:"required"`
	// Timestamp of when the request started processing.
	StartedAt time.Time `json:"startedAt" api:"required" format:"date-time"`
	// Status of the response.
	//
	// Any of "CANCELED", "COMPLETE", "PENDING", "PROCESSING".
	Status BatchResponseMarketingEventPublicDefaultResponseStatus `json:"status" api:"required"`
	Errors []shared.StandardError                                 `json:"errors"`
	// Result of the request.
	Links map[string]string `json:"links"`
	// The number of errors that occurred during the request.
	NumErrors int64 `json:"numErrors"`
	// Timestamp of when the request was sent.
	RequestedAt time.Time `json:"requestedAt" format:"date-time"`
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

// Status of the response.
type BatchResponseMarketingEventPublicDefaultResponseStatus string

const (
	BatchResponseMarketingEventPublicDefaultResponseStatusCanceled   BatchResponseMarketingEventPublicDefaultResponseStatus = "CANCELED"
	BatchResponseMarketingEventPublicDefaultResponseStatusComplete   BatchResponseMarketingEventPublicDefaultResponseStatus = "COMPLETE"
	BatchResponseMarketingEventPublicDefaultResponseStatusPending    BatchResponseMarketingEventPublicDefaultResponseStatus = "PENDING"
	BatchResponseMarketingEventPublicDefaultResponseStatusProcessing BatchResponseMarketingEventPublicDefaultResponseStatus = "PROCESSING"
)

type BatchResponseMarketingEventPublicDefaultResponseV2 struct {
	// Timestamp of when the request was processed.
	CompletedAt time.Time                               `json:"completedAt" api:"required" format:"date-time"`
	Results     []MarketingEventPublicDefaultResponseV2 `json:"results" api:"required"`
	// Timestamp of when the request started processing.
	StartedAt time.Time `json:"startedAt" api:"required" format:"date-time"`
	// The status of the response.
	//
	// Any of "CANCELED", "COMPLETE", "PENDING", "PROCESSING".
	Status BatchResponseMarketingEventPublicDefaultResponseV2Status `json:"status" api:"required"`
	// Result object of the request.
	Links map[string]string `json:"links"`
	// Timestamp of when the request was sent.
	RequestedAt time.Time `json:"requestedAt" format:"date-time"`
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

// The status of the response.
type BatchResponseMarketingEventPublicDefaultResponseV2Status string

const (
	BatchResponseMarketingEventPublicDefaultResponseV2StatusCanceled   BatchResponseMarketingEventPublicDefaultResponseV2Status = "CANCELED"
	BatchResponseMarketingEventPublicDefaultResponseV2StatusComplete   BatchResponseMarketingEventPublicDefaultResponseV2Status = "COMPLETE"
	BatchResponseMarketingEventPublicDefaultResponseV2StatusPending    BatchResponseMarketingEventPublicDefaultResponseV2Status = "PENDING"
	BatchResponseMarketingEventPublicDefaultResponseV2StatusProcessing BatchResponseMarketingEventPublicDefaultResponseV2Status = "PROCESSING"
)

type BatchResponseSubscriberEmailResponse struct {
	// Timestamp that represents when the request finished processing
	CompletedAt time.Time                 `json:"completedAt" api:"required" format:"date-time"`
	Results     []SubscriberEmailResponse `json:"results" api:"required"`
	// Timestamp that represents when the request started processing
	StartedAt time.Time `json:"startedAt" api:"required" format:"date-time"`
	// The status of the request processing
	//
	// Any of "CANCELED", "COMPLETE", "PENDING", "PROCESSING".
	Status BatchResponseSubscriberEmailResponseStatus `json:"status" api:"required"`
	Errors []shared.StandardError                     `json:"errors"`
	// Result of the request
	Links map[string]string `json:"links"`
	// The number of errors that occurred during the processing
	NumErrors int64 `json:"numErrors"`
	// Timestamp that represents when the request was made
	RequestedAt time.Time `json:"requestedAt" format:"date-time"`
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

// The status of the request processing
type BatchResponseSubscriberEmailResponseStatus string

const (
	BatchResponseSubscriberEmailResponseStatusCanceled   BatchResponseSubscriberEmailResponseStatus = "CANCELED"
	BatchResponseSubscriberEmailResponseStatusComplete   BatchResponseSubscriberEmailResponseStatus = "COMPLETE"
	BatchResponseSubscriberEmailResponseStatusPending    BatchResponseSubscriberEmailResponseStatus = "PENDING"
	BatchResponseSubscriberEmailResponseStatusProcessing BatchResponseSubscriberEmailResponseStatus = "PROCESSING"
)

type BatchResponseSubscriberVidResponse struct {
	// Timestamp that represents when the request finished processing
	CompletedAt time.Time               `json:"completedAt" api:"required" format:"date-time"`
	Results     []SubscriberVidResponse `json:"results" api:"required"`
	// Timestamp that represents when the request started processing
	StartedAt time.Time `json:"startedAt" api:"required" format:"date-time"`
	// The status of the request processing
	//
	// Any of "CANCELED", "COMPLETE", "PENDING", "PROCESSING".
	Status BatchResponseSubscriberVidResponseStatus `json:"status" api:"required"`
	Errors []shared.StandardError                   `json:"errors"`
	// Result of the request
	Links map[string]string `json:"links"`
	// The number of errors that occurred during the processing
	NumErrors int64 `json:"numErrors"`
	// Timestamp that represents when the request was made
	RequestedAt time.Time `json:"requestedAt" format:"date-time"`
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

// The status of the request processing
type BatchResponseSubscriberVidResponseStatus string

const (
	BatchResponseSubscriberVidResponseStatusCanceled   BatchResponseSubscriberVidResponseStatus = "CANCELED"
	BatchResponseSubscriberVidResponseStatusComplete   BatchResponseSubscriberVidResponseStatus = "COMPLETE"
	BatchResponseSubscriberVidResponseStatusPending    BatchResponseSubscriberVidResponseStatus = "PENDING"
	BatchResponseSubscriberVidResponseStatusProcessing BatchResponseSubscriberVidResponseStatus = "PROCESSING"
)

type CollectionResponseMarketingEventPublicReadResponseV2ForwardPaging struct {
	Results []MarketingEventPublicReadResponseV2 `json:"results" api:"required"`
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
	Results []SearchPublicResponseWrapper `json:"results" api:"required"`
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

type CollectionResponseWithTotalMarketingEventIdentifiersResponse struct {
	Results []MarketingEventIdentifiersResponse `json:"results" api:"required"`
	Total   int64                               `json:"total" api:"required"`
	Paging  shared.Paging                       `json:"paging"`
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
func (r CollectionResponseWithTotalMarketingEventIdentifiersResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *CollectionResponseWithTotalMarketingEventIdentifiersResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CollectionResponseWithTotalParticipationBreakdown struct {
	Results []ParticipationBreakdown `json:"results" api:"required"`
	Total   int64                    `json:"total" api:"required"`
	Paging  shared.Paging            `json:"paging"`
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
func (r CollectionResponseWithTotalParticipationBreakdown) RawJSON() string { return r.JSON.raw }
func (r *CollectionResponseWithTotalParticipationBreakdown) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CollectionResponseWithTotalPublicList struct {
	Results []PublicList  `json:"results" api:"required"`
	Total   int64         `json:"total" api:"required"`
	Paging  shared.Paging `json:"paging"`
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
func (r CollectionResponseWithTotalPublicList) RawJSON() string { return r.JSON.raw }
func (r *CollectionResponseWithTotalPublicList) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ContactAssociation struct {
	// The internal ID of the contact in HubSpot
	ContactID string `json:"contactId" api:"required"`
	// The email of the contact in HubSpot
	Email string `json:"email" api:"required"`
	// The first name of the contact in HubSpot
	Firstname string `json:"firstname"`
	// The last name of the contact in HubSpot
	Lastname string `json:"lastname"`
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

type CrmPropertyWrapper struct {
	// The name of the property in the CRM
	Name string `json:"name" api:"required"`
	// The value of the property in the CRM
	Value string `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CrmPropertyWrapper) RawJSON() string { return r.JSON.raw }
func (r *CrmPropertyWrapper) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EventDetailSettings struct {
	// The id of the application the settings are for
	AppID int64 `json:"appId" api:"required"`
	// The url that will be used to fetch marketing event details by id
	EventDetailsURL string `json:"eventDetailsUrl" api:"required"`
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
	EventDetailsURL string `json:"eventDetailsUrl" api:"required"`
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
	// The internal ID of the marketing event in HubSpot
	MarketingEventID string `json:"marketingEventId" api:"required"`
	// The name of the marketing event in HubSpot
	Name string `json:"name" api:"required"`
	// The account ID that is associated with this marketing event in the external
	// event application
	ExternalAccountID string `json:"externalAccountId"`
	// The event ID that is associated with this marketing event in the external event
	// application
	ExternalEventID string `json:"externalEventId"`
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
	// The end date and time of the marketing event in ISO 8601 format
	EndDateTime time.Time `json:"endDateTime" api:"required" format:"date-time"`
	// The start date and time of the marketing event in ISO 8601 format
	StartDateTime time.Time `json:"startDateTime" api:"required" format:"date-time"`
	paramObj
}

func (r MarketingEventCompleteRequestParams) MarshalJSON() (data []byte, err error) {
	type shadow MarketingEventCompleteRequestParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MarketingEventCompleteRequestParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties CustomProperties, EventName, EventOrganizer, ExternalAccountID,
// ExternalEventID are required.
type MarketingEventCreateRequestParams struct {
	// A list of PropertyValues. These can be whatever kind of property names and
	// values you want. However, they must already exist on the HubSpot account's
	// definition of the MarketingEvent Object. If they don't they will be filtered out
	// and not set. In order to do this you'll need to create a new PropertyGroup on
	// the HubSpot account's MarketingEvent object for your specific app and create the
	// Custom Property you want to track on that HubSpot account. Do not create any new
	// default properties on the MarketingEvent object as that will apply to all
	// HubSpot accounts.
	CustomProperties []shared.PropertyValueParam `json:"customProperties,omitzero" api:"required"`
	// The name of the marketing event.
	EventName string `json:"eventName" api:"required"`
	// The name of the organizer of the marketing event.
	EventOrganizer string `json:"eventOrganizer" api:"required"`
	// The accountId that is associated with this marketing event in the external event
	// application.
	ExternalAccountID string `json:"externalAccountId" api:"required"`
	// The id of the marketing event in the external event application.
	ExternalEventID string `json:"externalEventId" api:"required"`
	// The end date and time of the marketing event.
	EndDateTime param.Opt[time.Time] `json:"endDateTime,omitzero" format:"date-time"`
	// Indicates if the marketing event has been cancelled. Defaults to `false`
	EventCancelled param.Opt[bool] `json:"eventCancelled,omitzero"`
	// Indicates if the marketing event has been completed. Defaults to `false`
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
	// A list of PropertyValues. These can be whatever kind of property names and
	// values you want. However, they must already exist on the HubSpot account's
	// definition of the MarketingEvent Object. If they don't they will be filtered out
	// and not set. In order to do this you'll need to create a new PropertyGroup on
	// the HubSpot account's MarketingEvent object for your specific app and create the
	// Custom Property you want to track on that HubSpot account. Do not create any new
	// default properties on the MarketingEvent object as that will apply to all
	// HubSpot accounts.
	CustomProperties []shared.PropertyValue `json:"customProperties" api:"required"`
	// The name of the marketing event.
	EventName string `json:"eventName" api:"required"`
	// The name of the organizer of the marketing event.
	EventOrganizer string `json:"eventOrganizer" api:"required"`
	// The end date and time of the marketing event.
	EndDateTime time.Time `json:"endDateTime" format:"date-time"`
	// Indicates if the marketing event has been cancelled.
	EventCancelled bool `json:"eventCancelled"`
	// Indicates if the marketing event has been completed.
	EventCompleted bool `json:"eventCompleted"`
	// The description of the marketing event.
	EventDescription string `json:"eventDescription"`
	// The type of the marketing event.
	EventType string `json:"eventType"`
	// The URL in the external event application where the marketing event can be
	// managed.
	EventURL string `json:"eventUrl"`
	// The ID of the marketing event CRM object
	ObjectID string `json:"objectId"`
	// The start date and time of the marketing event.
	StartDateTime time.Time `json:"startDateTime" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CustomProperties respjson.Field
		EventName        respjson.Field
		EventOrganizer   respjson.Field
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

// The properties ContactProperties, Email, InteractionDateTime, Properties are
// required.
type MarketingEventEmailSubscriberParam struct {
	// The key-value set that contains properties of the contact.
	ContactProperties map[string]string `json:"contactProperties,omitzero" api:"required"`
	// The email address of the contact in HubSpot to associate with the event.
	Email string `json:"email" api:"required"`
	// Timestamp in milliseconds at which the contact subscribed to the event.
	InteractionDateTime int64 `json:"interactionDateTime" api:"required"`
	// The key-value set that contains properties of the marketing event.
	Properties map[string]string `json:"properties,omitzero" api:"required"`
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
	AppID int64 `json:"appId" api:"required"`
	// The accountId that is associated with this marketing event in the external event
	// application.
	ExternalAccountID string `json:"externalAccountId" api:"required"`
	// The id of the marketing event in the external event application.
	ExternalEventID string `json:"externalEventId" api:"required"`
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
	// The ID that is associated with this marketing event in the external event
	// application
	ExternalEventID string `json:"externalEventId" api:"required"`
	// The name of the marketing event
	MarketingEventName string `json:"marketingEventName" api:"required"`
	// The internal ID of the marketing event in HubSpot CRM
	ObjectID string  `json:"objectId" api:"required"`
	AppInfo  AppInfo `json:"appInfo"`
	// The accountId that is associated with this marketing event in the external event
	// application
	ExternalAccountID string `json:"externalAccountId"`
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
	// The ID of the marketing event CRM object.
	ID string `json:"id" api:"required"`
	// The creation date and time of the marketing event.
	CreatedAt time.Time `json:"createdAt" api:"required" format:"date-time"`
	// A list of PropertyValues. These can be whatever kind of property names and
	// values you want. However, they must already exist on the HubSpot account's
	// definition of the MarketingEvent Object. If they don't they will be filtered out
	// and not set. In order to do this you'll need to create a new PropertyGroup on
	// the HubSpot account's MarketingEvent object for your specific app and create the
	// Custom Property you want to track on that HubSpot account. Do not create any new
	// default properties on the MarketingEvent object as that will apply to all
	// HubSpot accounts.
	CustomProperties []shared.PropertyValue `json:"customProperties" api:"required"`
	// The name of the marketing event.
	EventName string `json:"eventName" api:"required"`
	// The name of the organizer of the marketing event.
	EventOrganizer string `json:"eventOrganizer" api:"required"`
	// The update date and time of the marketing event.
	UpdatedAt time.Time `json:"updatedAt" api:"required" format:"date-time"`
	// The end date and time of the marketing event.
	EndDateTime time.Time `json:"endDateTime" format:"date-time"`
	// Indicates if the marketing event has been cancelled.
	EventCancelled bool `json:"eventCancelled"`
	// Indicates if the marketing event has been completed.
	EventCompleted bool `json:"eventCompleted"`
	// The description of the marketing event.
	EventDescription string `json:"eventDescription"`
	// The type of the marketing event.
	EventType string `json:"eventType"`
	// A URL in the external event application where the marketing event can be
	// managed.
	EventURL string `json:"eventUrl"`
	// The ID of the marketing event CRM object.
	ObjectID string `json:"objectId"`
	// The start date and time of the marketing event.
	StartDateTime time.Time `json:"startDateTime" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		CreatedAt        respjson.Field
		CustomProperties respjson.Field
		EventName        respjson.Field
		EventOrganizer   respjson.Field
		UpdatedAt        respjson.Field
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
	// The creation date and time of the marketing event
	CreatedAt        time.Time            `json:"createdAt" api:"required" format:"date-time"`
	CustomProperties []CrmPropertyWrapper `json:"customProperties" api:"required"`
	// The name of the marketing event
	EventName string `json:"eventName" api:"required"`
	// The internal ID of the marketing event in HubSpot
	ObjectID string `json:"objectId" api:"required"`
	// The update date and time of the marketing event
	UpdatedAt time.Time `json:"updatedAt" api:"required" format:"date-time"`
	AppInfo   AppInfo   `json:"appInfo"`
	// The end date and time of the marketing event
	EndDateTime time.Time `json:"endDateTime" format:"date-time"`
	// Indicates if the marketing event has been cancelled
	EventCancelled bool `json:"eventCancelled"`
	// Indicates if the marketing event has been completed
	EventCompleted bool `json:"eventCompleted"`
	// The description of the marketing event
	EventDescription string `json:"eventDescription"`
	// The name of the organizer of the marketing event
	EventOrganizer string `json:"eventOrganizer"`
	// The type of the marketing event
	EventType string `json:"eventType"`
	// A URL in the external event application where the marketing event can be managed
	EventURL string `json:"eventUrl"`
	// The start date and time of the marketing event
	StartDateTime time.Time `json:"startDateTime" format:"date-time"`
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
	// The internal ID of the marketing event in HubSpot
	ObjectID string `json:"objectId" api:"required"`
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
	// The ID of the marketing event CRM object.
	ID string `json:"id" api:"required"`
	// The number of HubSpot contacts that attended this marketing event.
	Attendees int64 `json:"attendees" api:"required"`
	// The number of HubSpot contacts that registered for this marketing event, but
	// later cancelled their registration.
	Cancellations int64 `json:"cancellations" api:"required"`
	// The creation date and time of the marketing event.
	CreatedAt time.Time `json:"createdAt" api:"required" format:"date-time"`
	// A list of PropertyValues. These can be whatever kind of property names and
	// values you want. However, they must already exist on the HubSpot account's
	// definition of the MarketingEvent Object. If they don't they will be filtered out
	// and not set. In order to do this you'll need to create a new PropertyGroup on
	// the HubSpot account's MarketingEvent object for your specific app and create the
	// Custom Property you want to track on that HubSpot account. Do not create any new
	// default properties on the MarketingEvent object as that will apply to all
	// HubSpot accounts.
	CustomProperties []shared.PropertyValue `json:"customProperties" api:"required"`
	// The name of the marketing event.
	EventName string `json:"eventName" api:"required"`
	// The name of the organizer of the marketing event.
	EventOrganizer string `json:"eventOrganizer" api:"required"`
	// The id of the marketing event in the external event application.
	ExternalEventID string `json:"externalEventId" api:"required"`
	// The number of HubSpot contacts that registered for this marketing event, but did
	// not attend. This field only had a value when the event is over.
	NoShows int64 `json:"noShows" api:"required"`
	// The number of HubSpot contacts that registered for this marketing event.
	Registrants int64 `json:"registrants" api:"required"`
	// The update date and time of the marketing event.
	UpdatedAt time.Time `json:"updatedAt" api:"required" format:"date-time"`
	// The end date and time of the marketing event.
	EndDateTime time.Time `json:"endDateTime" format:"date-time"`
	// Indicates if the marketing event has been cancelled.
	EventCancelled bool `json:"eventCancelled"`
	// Indicates if the marketing event has been completed.
	EventCompleted bool `json:"eventCompleted"`
	// The description of the marketing event.
	EventDescription string `json:"eventDescription"`
	// The type of the marketing event.
	EventType string `json:"eventType"`
	// A URL in the external event application where the marketing event can be
	// managed.
	EventURL string `json:"eventUrl"`
	// The ID of the marketing event CRM object.
	ObjectID string `json:"objectId"`
	// The start date and time of the marketing event.
	StartDateTime time.Time `json:"startDateTime" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		Attendees        respjson.Field
		Cancellations    respjson.Field
		CreatedAt        respjson.Field
		CustomProperties respjson.Field
		EventName        respjson.Field
		EventOrganizer   respjson.Field
		ExternalEventID  respjson.Field
		NoShows          respjson.Field
		Registrants      respjson.Field
		UpdatedAt        respjson.Field
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
	// The creation date and time of the marketing event
	CreatedAt        time.Time            `json:"createdAt" api:"required" format:"date-time"`
	CustomProperties []CrmPropertyWrapper `json:"customProperties" api:"required"`
	// The name of the marketing event
	EventName string `json:"eventName" api:"required"`
	// The internal ID of the marketing event in HubSpot
	ObjectID string `json:"objectId" api:"required"`
	// The update date and time of the marketing event
	UpdatedAt time.Time `json:"updatedAt" api:"required" format:"date-time"`
	AppInfo   AppInfo   `json:"appInfo"`
	// Number of attended contact records of a marketing event
	Attendees int64 `json:"attendees"`
	// Number of cancelled contact records of a marketing event
	Cancellations int64 `json:"cancellations"`
	// The end date and time of the marketing event
	EndDateTime time.Time `json:"endDateTime" format:"date-time"`
	// Indicates if the marketing event has been cancelled
	EventCancelled bool `json:"eventCancelled"`
	// Indicates if the marketing event has been completed
	EventCompleted bool `json:"eventCompleted"`
	// The description of the marketing event
	EventDescription string `json:"eventDescription"`
	// The name of the organizer of the marketing event
	EventOrganizer string `json:"eventOrganizer"`
	// The status of the marketing event
	EventStatus   string `json:"eventStatus"`
	EventStatusV2 string `json:"eventStatusV2"`
	// The type of the marketing event
	EventType string `json:"eventType"`
	// A URL in the external event application where the marketing event can be managed
	EventURL string `json:"eventUrl"`
	// The ID that is associated with this marketing event in the external event
	// application
	ExternalEventID string `json:"externalEventId"`
	// Number of no-show contact records of a marketing event
	NoShows int64 `json:"noShows"`
	// Number of registered contact records of a marketing event
	Registrants int64 `json:"registrants"`
	// The start date and time of the marketing event
	StartDateTime time.Time `json:"startDateTime" format:"date-time"`
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
		EventStatusV2    respjson.Field
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
	CustomProperties []shared.PropertyValueParam `json:"customProperties,omitzero" api:"required"`
	// The internal ID of the marketing event in HubSpot
	ObjectID string `json:"objectId" api:"required"`
	// The end date and time of the marketing event
	EndDateTime param.Opt[time.Time] `json:"endDateTime,omitzero" format:"date-time"`
	// Indicates if the marketing event has been cancelled
	EventCancelled param.Opt[bool] `json:"eventCancelled,omitzero"`
	// The description of the marketing event
	EventDescription param.Opt[string] `json:"eventDescription,omitzero"`
	// The name of the marketing event
	EventName param.Opt[string] `json:"eventName,omitzero"`
	// The name of the organizer of the marketing event
	EventOrganizer param.Opt[string] `json:"eventOrganizer,omitzero"`
	// The type of the marketing event
	EventType param.Opt[string] `json:"eventType,omitzero"`
	// A URL in the external event application where the marketing event can be managed
	EventURL param.Opt[string] `json:"eventUrl,omitzero"`
	// The start date and time of the marketing event
	StartDateTime param.Opt[time.Time] `json:"startDateTime,omitzero" format:"date-time"`
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
	CustomProperties []shared.PropertyValueParam `json:"customProperties,omitzero" api:"required"`
	// The end date and time of the marketing event
	EndDateTime param.Opt[time.Time] `json:"endDateTime,omitzero" format:"date-time"`
	// Indicates if the marketing event has been cancelled
	EventCancelled param.Opt[bool] `json:"eventCancelled,omitzero"`
	// The description of the marketing event
	EventDescription param.Opt[string] `json:"eventDescription,omitzero"`
	// The name of the marketing event
	EventName param.Opt[string] `json:"eventName,omitzero"`
	// The name of the organizer of the marketing event
	EventOrganizer param.Opt[string] `json:"eventOrganizer,omitzero"`
	// The type of the marketing event
	EventType param.Opt[string] `json:"eventType,omitzero"`
	// A URL in the external event application where the marketing event can be managed
	EventURL param.Opt[string] `json:"eventUrl,omitzero"`
	// The start date and time of the marketing event
	StartDateTime param.Opt[time.Time] `json:"startDateTime,omitzero" format:"date-time"`
	paramObj
}

func (r MarketingEventPublicUpdateRequestV2Param) MarshalJSON() (data []byte, err error) {
	type shadow MarketingEventPublicUpdateRequestV2Param
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MarketingEventPublicUpdateRequestV2Param) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties InteractionDateTime, Properties, Vid are required.
type MarketingEventSubscriberParam struct {
	// Timestamp in milliseconds at which the contact subscribed to the event.
	InteractionDateTime int64 `json:"interactionDateTime" api:"required"`
	// The key-value set of the properties of the contact
	Properties map[string]string `json:"properties,omitzero" api:"required"`
	// The ID of the contact in HubSpot
	Vid int64 `json:"vid" api:"required"`
	paramObj
}

func (r MarketingEventSubscriberParam) MarshalJSON() (data []byte, err error) {
	type shadow MarketingEventSubscriberParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MarketingEventSubscriberParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property CustomProperties is required.
type MarketingEventUpdateRequestParams struct {
	// A list of PropertyValues. These can be whatever kind of property names and
	// values you want. However, they must already exist on the HubSpot account's
	// definition of the MarketingEvent Object. If they don't they will be filtered out
	// and not set. In order to do this you'll need to create a new PropertyGroup on
	// the HubSpot account's MarketingEvent object for your specific app and create the
	// Custom Property you want to track on that HubSpot account. Do not create any new
	// default properties on the MarketingEvent object as that will apply to all
	// HubSpot accounts.
	CustomProperties []shared.PropertyValueParam `json:"customProperties,omitzero" api:"required"`
	// The end date and time of the marketing event.
	EndDateTime param.Opt[time.Time] `json:"endDateTime,omitzero" format:"date-time"`
	// Indicates if the marketing event has been cancelled. Defaults to `false`
	EventCancelled param.Opt[bool] `json:"eventCancelled,omitzero"`
	// Indicates if the marketing event has been completed. Defaults to `false`
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
	Contact        ContactAssociation        `json:"contact" api:"required"`
	MarketingEvent MarketingEventAssociation `json:"marketingEvent" api:"required"`
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
	// The internal ID of the target marketing event
	ID           string                    `json:"id" api:"required"`
	Associations ParticipationAssociations `json:"associations" api:"required"`
	// The creation time and date of the target marketing event
	CreatedAt  time.Time               `json:"createdAt" api:"required" format:"date-time"`
	Properties ParticipationProperties `json:"properties" api:"required"`
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
	// The state of the participation
	//
	// Any of "ATTENDED", "CANCELLED", "EMPTY", "NO_SHOW", "REGISTERED".
	AttendanceState ParticipationPropertiesAttendanceState `json:"attendanceState" api:"required"`
	// Timestamp of when the participation occurred
	OccurredAt int64 `json:"occurredAt" api:"required"`
	// The number of seconds the participation lasted
	AttendanceDurationSeconds int64 `json:"attendanceDurationSeconds"`
	// Percentage of the participation duration relative to the event duration
	AttendancePercentage string `json:"attendancePercentage"`
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

// The state of the participation
type ParticipationPropertiesAttendanceState string

const (
	ParticipationPropertiesAttendanceStateAttended   ParticipationPropertiesAttendanceState = "ATTENDED"
	ParticipationPropertiesAttendanceStateCancelled  ParticipationPropertiesAttendanceState = "CANCELLED"
	ParticipationPropertiesAttendanceStateEmpty      ParticipationPropertiesAttendanceState = "EMPTY"
	ParticipationPropertiesAttendanceStateNoShow     ParticipationPropertiesAttendanceState = "NO_SHOW"
	ParticipationPropertiesAttendanceStateRegistered ParticipationPropertiesAttendanceState = "REGISTERED"
)

type PublicList struct {
	// An internal ID of the list
	ListID string `json:"listId" api:"required"`
	// A number that represents a version of the list
	ListVersion int64 `json:"listVersion" api:"required"`
	// The name of the list
	Name string `json:"name" api:"required"`
	// The internal ID of the object type of the list
	ObjectTypeID string `json:"objectTypeId" api:"required"`
	// Represents the current processing status of the list
	ProcessingStatus string `json:"processingStatus" api:"required"`
	// Processing type of the list
	ProcessingType string `json:"processingType" api:"required"`
	// Timestamp of the creation of the list
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// The ID of the user who created the list
	CreatedByID string `json:"createdById"`
	// Timestamp of the deletion of the list
	DeletedAt time.Time `json:"deletedAt" format:"date-time"`
	// Timestamp of the last update of the list filters
	FiltersUpdatedAt time.Time `json:"filtersUpdatedAt" format:"date-time"`
	// The size of the result list
	Size int64 `json:"size"`
	// Timestamp of the last update of the list
	UpdatedAt time.Time `json:"updatedAt" format:"date-time"`
	// The ID of the user who last updated the list
	UpdatedByID string `json:"updatedById"`
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
	// The ID of the source application of the marketing event
	AppID int64 `json:"appId" api:"required"`
	// The account ID associated with this marketing event in the external application
	ExternalAccountID string `json:"externalAccountId" api:"required"`
	// The ID of the marketing event in the external event application
	ExternalEventID string `json:"externalEventId" api:"required"`
	// The internal ID of the marketing event in HubSpot
	ObjectID string `json:"objectId" api:"required"`
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
	// The email of the contact
	Email string `json:"email" api:"required"`
	// The internal ID of the contact
	Vid int64 `json:"vid" api:"required"`
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
	// The internal ID of the contact
	Vid int64 `json:"vid" api:"required"`
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

type MarketingEventNewParams struct {
	MarketingEventCreateRequestParams MarketingEventCreateRequestParams
	paramObj
}

func (r MarketingEventNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.MarketingEventCreateRequestParams)
}
func (r *MarketingEventNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MarketingEventUpdateParams struct {
	MarketingEventPublicUpdateRequestV2 MarketingEventPublicUpdateRequestV2Param
	paramObj
}

func (r MarketingEventUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.MarketingEventPublicUpdateRequestV2)
}
func (r *MarketingEventUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MarketingEventListParams struct {
	// The cursor indicating the position of the last retrieved item.
	After param.Opt[string] `query:"after,omitzero" json:"-"`
	// The limit for response size. The default value is 10, the max number is 100
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [MarketingEventListParams]'s query parameters as
// `url.Values`.
func (r MarketingEventListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type MarketingEventDeleteBatchParams struct {
	BatchInputMarketingEventPublicObjectIDDeleteRequest BatchInputMarketingEventPublicObjectIDDeleteRequestParam
	paramObj
}

func (r MarketingEventDeleteBatchParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.BatchInputMarketingEventPublicObjectIDDeleteRequest)
}
func (r *MarketingEventDeleteBatchParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MarketingEventDeleteBatchByExternalEventIDParams struct {
	BatchInputMarketingEventExternalUniqueIdentifier BatchInputMarketingEventExternalUniqueIdentifierParam
	paramObj
}

func (r MarketingEventDeleteBatchByExternalEventIDParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.BatchInputMarketingEventExternalUniqueIdentifier)
}
func (r *MarketingEventDeleteBatchByExternalEventIDParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MarketingEventDeleteByExternalEventIDParams struct {
	ExternalAccountID string `query:"externalAccountId" api:"required" json:"-"`
	paramObj
}

// URLQuery serializes [MarketingEventDeleteByExternalEventIDParams]'s query
// parameters as `url.Values`.
func (r MarketingEventDeleteByExternalEventIDParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type MarketingEventGetByExternalEventIDParams struct {
	ExternalAccountID string `query:"externalAccountId" api:"required" json:"-"`
	paramObj
}

// URLQuery serializes [MarketingEventGetByExternalEventIDParams]'s query
// parameters as `url.Values`.
func (r MarketingEventGetByExternalEventIDParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type MarketingEventSearchByExternalEventIDParams struct {
	Q string `query:"q" api:"required" json:"-"`
	paramObj
}

// URLQuery serializes [MarketingEventSearchByExternalEventIDParams]'s query
// parameters as `url.Values`.
func (r MarketingEventSearchByExternalEventIDParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type MarketingEventUpdateBatchParams struct {
	BatchInputMarketingEventPublicUpdateRequestFullV2 BatchInputMarketingEventPublicUpdateRequestFullV2Param
	paramObj
}

func (r MarketingEventUpdateBatchParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.BatchInputMarketingEventPublicUpdateRequestFullV2)
}
func (r *MarketingEventUpdateBatchParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MarketingEventUpdateByExternalEventIDParams struct {
	ExternalAccountID                 string `query:"externalAccountId" api:"required" json:"-"`
	MarketingEventUpdateRequestParams MarketingEventUpdateRequestParams
	paramObj
}

func (r MarketingEventUpdateByExternalEventIDParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.MarketingEventUpdateRequestParams)
}
func (r *MarketingEventUpdateByExternalEventIDParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [MarketingEventUpdateByExternalEventIDParams]'s query
// parameters as `url.Values`.
func (r MarketingEventUpdateByExternalEventIDParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type MarketingEventUpsertBatchParams struct {
	BatchInputMarketingEventCreateRequestParams BatchInputMarketingEventCreateRequestParams
	paramObj
}

func (r MarketingEventUpsertBatchParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.BatchInputMarketingEventCreateRequestParams)
}
func (r *MarketingEventUpsertBatchParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MarketingEventUpsertByExternalEventIDParams struct {
	MarketingEventCreateRequestParams MarketingEventCreateRequestParams
	paramObj
}

func (r MarketingEventUpsertByExternalEventIDParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.MarketingEventCreateRequestParams)
}
func (r *MarketingEventUpsertByExternalEventIDParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
