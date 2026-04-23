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

// OccurrenceService contains methods and other services that help with interacting
// with the hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewOccurrenceService] method instead.
type OccurrenceService struct {
	options []option.RequestOption
}

// NewOccurrenceService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewOccurrenceService(opts ...option.RequestOption) (r OccurrenceService) {
	r = OccurrenceService{}
	r.options = opts
	return
}

// Retrieve event occurrences for the specified time frame. This endpoint allows
// filtering by various parameters such as object type, event type, and occurrence
// time. It supports pagination and sorting of results.
func (r *OccurrenceService) List(ctx context.Context, query OccurrenceListParams, opts ...option.RequestOption) (res *pagination.Page[ExternalUnifiedEvent], err error) {
	var raw *http.Response
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "events/event-occurrences/2026-03"
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

// Retrieve event occurrences for the specified time frame. This endpoint allows
// filtering by various parameters such as object type, event type, and occurrence
// time. It supports pagination and sorting of results.
func (r *OccurrenceService) ListAutoPaging(ctx context.Context, query OccurrenceListParams, opts ...option.RequestOption) *pagination.PageAutoPager[ExternalUnifiedEvent] {
	return pagination.NewPageAutoPager(r.List(ctx, query, opts...))
}

// Retrieve a list of event type names. You may use these event types to query the
// API for event occurrences of a desired type.
//
// Note: the `get_types` method is only supported in the Python SDK version
// `12.0.0-beta.1` or later.
func (r *OccurrenceService) ListEventTypes(ctx context.Context, opts ...option.RequestOption) (res *VisibleExternalEventTypeNames, err error) {
	opts = slices.Concat(r.options, opts)
	path := "events/event-occurrences/2026-03/event-types"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type CollectionResponseExternalUnifiedEvent struct {
	// An array of ExternalUnifiedEvent objects, each representing an individual event
	// with its associated details.
	Results []ExternalUnifiedEvent `json:"results" api:"required"`
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
	ID string `json:"id" api:"required"`
	// The format of the `eventType` string is `ae{appId}_{eventTypeLabel}`,
	// `pe{portalId}_{eventTypeLabel}`, or just `e_{eventTypeLabel}` for HubSpot
	// events.
	EventType string `json:"eventType" api:"required"`
	// The objectId of the object which did the event.
	ObjectID string `json:"objectId" api:"required"`
	// The objectType for the object which did the event.
	ObjectType string `json:"objectType" api:"required"`
	// An ISO 8601 timestamp when the event occurred.
	OccurredAt time.Time `json:"occurredAt" api:"required" format:"date-time"`
	// A key-value map of event-specific properties. The available properties depend on
	// the event type definition.
	Properties map[string]string `json:"properties" api:"required"`
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
	EventTypes []string `json:"eventTypes" api:"required"`
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

type OccurrenceListParams struct {
	// The paging cursor token of the last successfully read resource will be returned
	// as the `paging.next.after` JSON property of a paged response containing more
	// results.
	After     param.Opt[string] `query:"after,omitzero" json:"-"`
	Before    param.Opt[string] `query:"before,omitzero" json:"-"`
	EventType param.Opt[string] `query:"eventType,omitzero" json:"-"`
	// The maximum number of results to display per page.
	Limit          param.Opt[int64]                   `query:"limit,omitzero" json:"-"`
	ObjectID       param.Opt[int64]                   `query:"objectId,omitzero" json:"-"`
	ObjectType     param.Opt[string]                  `query:"objectType,omitzero" json:"-"`
	OccurredAfter  param.Opt[time.Time]               `query:"occurredAfter,omitzero" format:"date-time" json:"-"`
	OccurredBefore param.Opt[time.Time]               `query:"occurredBefore,omitzero" format:"date-time" json:"-"`
	ID             []string                           `query:"id,omitzero" json:"-"`
	ObjectProperty OccurrenceListParamsObjectProperty `query:"objectProperty,omitzero" json:"-"`
	Properties     []string                           `query:"properties,omitzero" json:"-"`
	Property       OccurrenceListParamsProperty       `query:"property,omitzero" json:"-"`
	Sort           []string                           `query:"sort,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [OccurrenceListParams]'s query parameters as `url.Values`.
func (r OccurrenceListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type OccurrenceListParamsObjectProperty struct {
	Propname any `query:"{propname},omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [OccurrenceListParamsObjectProperty]'s query parameters as
// `url.Values`.
func (r OccurrenceListParamsObjectProperty) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type OccurrenceListParamsProperty struct {
	Propname any `query:"{propname},omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [OccurrenceListParamsProperty]'s query parameters as
// `url.Values`.
func (r OccurrenceListParamsProperty) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
