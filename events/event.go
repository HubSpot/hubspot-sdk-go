// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package events

import (
	"context"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/stainless-sdks/hubspot-sdk-go/internal/apijson"
	"github.com/stainless-sdks/hubspot-sdk-go/internal/apiquery"
	"github.com/stainless-sdks/hubspot-sdk-go/internal/requestconfig"
	"github.com/stainless-sdks/hubspot-sdk-go/option"
	"github.com/stainless-sdks/hubspot-sdk-go/packages/pagination"
	"github.com/stainless-sdks/hubspot-sdk-go/packages/param"
	"github.com/stainless-sdks/hubspot-sdk-go/packages/respjson"
	"github.com/stainless-sdks/hubspot-sdk-go/shared"
)

// EventService contains methods and other services that help with interacting with
// the hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEventService] method instead.
type EventService struct {
	Options          []option.RequestOption
	EventDefinitions EventDefinitionService
	Send             SendService
}

// NewEventService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewEventService(opts ...option.RequestOption) (r EventService) {
	r = EventService{}
	r.Options = opts
	r.EventDefinitions = NewEventDefinitionService(opts...)
	r.Send = NewSendService(opts...)
	return
}

// Retrieve instances of event completion data. For example, retrieve all event
// completions associated with a specific contact.
func (r *EventService) List(ctx context.Context, query EventListParams, opts ...option.RequestOption) (res *pagination.Page[ExternalUnifiedEvent], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "events/v3/events/"
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

// Retrieve instances of event completion data. For example, retrieve all event
// completions associated with a specific contact.
func (r *EventService) ListAutoPaging(ctx context.Context, query EventListParams, opts ...option.RequestOption) *pagination.PageAutoPager[ExternalUnifiedEvent] {
	return pagination.NewPageAutoPager(r.List(ctx, query, opts...))
}

// This endpoint returns a list of event type names which are visible to you. You
// may use these event type names to query the API for specific event instances of
// a desired type.
//
// Note: the `get_types` method is only supported in the Python SDK version
// `12.0.0-beta.1` or later.
func (r *EventService) ListEventTypes(ctx context.Context, opts ...option.RequestOption) (res *VisibleExternalEventTypeNames, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "events/v3/events/event-types"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type CollectionResponseExternalUnifiedEvent struct {
	Results []ExternalUnifiedEvent `json:"results,required"`
	Paging  shared.Paging          `json:"paging"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Results     respjson.Field
		Paging      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CollectionResponseExternalUnifiedEvent) RawJSON() string { return r.JSON.raw }
func (r *CollectionResponseExternalUnifiedEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExternalUnifiedEvent struct {
	// A unique identifier for the event.
	ID string `json:"id,required"`
	// The format of the `eventType` string is `ae{appId}_{eventTypeLabel}`,
	// `pe{portalId}_{eventTypeLabel}`, or just `e_{eventTypeLabel}` for HubSpot
	// events.
	EventType string `json:"eventType,required"`
	// The objectId of the object which did the event.
	ObjectID string `json:"objectId,required"`
	// The objectType for the object which did the event.
	ObjectType string `json:"objectType,required"`
	// An ISO 8601 timestamp when the event occurred.
	OccurredAt time.Time `json:"occurredAt,required" format:"date-time"`
	// A key-value map of event-specific properties. The available properties depend on
	// the event type definition.
	Properties map[string]string `json:"properties,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		EventType   respjson.Field
		ObjectID    respjson.Field
		ObjectType  respjson.Field
		OccurredAt  respjson.Field
		Properties  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExternalUnifiedEvent) RawJSON() string { return r.JSON.raw }
func (r *ExternalUnifiedEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VisibleExternalEventTypeNames struct {
	// List of event type names.
	EventTypes []string `json:"eventTypes,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		EventTypes  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VisibleExternalEventTypeNames) RawJSON() string { return r.JSON.raw }
func (r *VisibleExternalEventTypeNames) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EventListParams struct {
	// The paging cursor token of the last successfully read resource will be returned
	// as the `paging.next.after` JSON property of a paged response containing more
	// results.
	After  param.Opt[string] `query:"after,omitzero" json:"-"`
	Before param.Opt[string] `query:"before,omitzero" json:"-"`
	// The event type name. You can retrieve available event types using the
	// [event types endpoint](#get-%2Fevents%2Fv3%2Fevents%2Fevent-types).
	EventType param.Opt[string] `query:"eventType,omitzero" json:"-"`
	// The maximum number of results to display per page.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// The ID of the CRM Object to filter event instances on. When including this
	// parameter, you must also include the `objectType` parameter.
	ObjectID param.Opt[int64] `query:"objectId,omitzero" json:"-"`
	// The type of CRM object to filter event instances on (e.g., `contact`). To
	// retrieve event data for a specific CRM record, include the additional `objectId`
	// query parameter (below).
	ObjectType param.Opt[string] `query:"objectType,omitzero" json:"-"`
	// Filter for event data that occurred after a specific datetime.
	OccurredAfter param.Opt[time.Time] `query:"occurredAfter,omitzero" format:"date-time" json:"-"`
	// Filter for event data that occurred before a specific datetime.
	OccurredBefore param.Opt[time.Time] `query:"occurredBefore,omitzero" format:"date-time" json:"-"`
	// ID of an event instance. IDs are 1:1 with event instances. If you provide this
	// filter and additional filters, the other filters must match the values on the
	// event instance to yield results.
	ID             []string                      `query:"id,omitzero" json:"-"`
	ObjectProperty EventListParamsObjectProperty `query:"objectProperty,omitzero" json:"-"`
	Property       EventListParamsProperty       `query:"property,omitzero" json:"-"`
	// Sort direction based on the timestamp of the event instance, `ASCENDING` or
	// `DESCENDING`.
	Sort []string `query:"sort,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [EventListParams]'s query parameters as `url.Values`.
func (r EventListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type EventListParamsObjectProperty struct {
	// Instead of retrieving event data for a specific object by its ID, you can
	// specify a unique identifier property. For contacts, you can use the `email`
	// property. (e.g., `objectProperty.email=name@domain.com`).
	Propname any `query:"{propname},omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [EventListParamsObjectProperty]'s query parameters as
// `url.Values`.
func (r EventListParamsObjectProperty) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type EventListParamsProperty struct {
	// Filter for event completions that contain a specific value for an event property
	// (e.g., `property.hs_city=portland`). For properties values with spaces, replaces
	// spaces with `%20` or `+` (e.g., `property.hs_city=new+york`).
	Propname any `query:"{propname},omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [EventListParamsProperty]'s query parameters as
// `url.Values`.
func (r EventListParamsProperty) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
