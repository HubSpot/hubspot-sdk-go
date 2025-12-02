// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package crm

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"slices"

	shimjson "github.com/stainless-sdks/hubspot-sdk-go/internal/encoding/json"
	"github.com/stainless-sdks/hubspot-sdk-go/internal/requestconfig"
	"github.com/stainless-sdks/hubspot-sdk-go/option"
)

// TimelineTemplateService contains methods and other services that help with
// interacting with the hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTimelineTemplateService] method instead.
type TimelineTemplateService struct {
	Options []option.RequestOption
}

// NewTimelineTemplateService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewTimelineTemplateService(opts ...option.RequestOption) (r TimelineTemplateService) {
	r = TimelineTemplateService{}
	r.Options = opts
	return
}

// Event templates define the general structure for a custom timeline event, and
// enable you to send event data to HubSpot. A template includes formatted copy for
// its heading and details, as well as any custom property definitions. A single
// app can include up to 750 event templates.<br/><Warning>the `v1` and `v3`
// timeline events APIs are only available for app partners with existing `v1`/`v3`
// timeline events defined in their public app. <ul><li>If your app doesn't include
// any timeline events yet, requests to this endpoint will fail. Instead, you can
// get started on
// [latest version of the developer platform](/apps/developer-platform/build-apps/overview).
// Note that you'll need to request approval before you can define app events for
// your app. Learn more in the
// [app events overview](/apps/developer-platform/add-features/app-events/overview).</li><li>If
// your app includes a `v1`/`v3` timeline event, learn how to
// [migrate it to the developer platform](/apps/developer-platform/add-features/app-events/create-and-manage-event-types#migrate-an-existing-timeline-event-type).
// You don't need to request approval before migrating existing event
// types.</li></ul>If you're not an app partner, you can send custom event data to
// HubSpot using the
// [custom events API](/api-reference/events-manage-event-definitions-v3/guide).</Warning>
func (r *TimelineTemplateService) New(ctx context.Context, appID int64, body TimelineTemplateNewParams, opts ...option.RequestOption) (res *TimelineEventTemplate, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("integrators/timeline/v3/%v/event-templates", appID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Update an existing event template, specified by ID.
func (r *TimelineTemplateService) Update(ctx context.Context, eventTemplateID string, params TimelineTemplateUpdateParams, opts ...option.RequestOption) (res *TimelineEventTemplate, err error) {
	opts = slices.Concat(r.Options, opts)
	if eventTemplateID == "" {
		err = errors.New("missing required eventTemplateId parameter")
		return
	}
	path := fmt.Sprintf("integrators/timeline/v3/%v/event-templates/%s", params.AppID, eventTemplateID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return
}

// Retrieve all templates defined for an app.
func (r *TimelineTemplateService) List(ctx context.Context, appID int64, opts ...option.RequestOption) (res *CollectionResponseTimelineEventTemplateNoPaging, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("integrators/timeline/v3/%v/event-templates", appID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Delete an event type template by ID.
func (r *TimelineTemplateService) Delete(ctx context.Context, eventTemplateID string, body TimelineTemplateDeleteParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if eventTemplateID == "" {
		err = errors.New("missing required eventTemplateId parameter")
		return
	}
	path := fmt.Sprintf("integrators/timeline/v3/%v/event-templates/%s", body.AppID, eventTemplateID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Retrieve an event type template by ID.
func (r *TimelineTemplateService) Get(ctx context.Context, eventTemplateID string, query TimelineTemplateGetParams, opts ...option.RequestOption) (res *TimelineEventTemplate, err error) {
	opts = slices.Concat(r.Options, opts)
	if eventTemplateID == "" {
		err = errors.New("missing required eventTemplateId parameter")
		return
	}
	path := fmt.Sprintf("integrators/timeline/v3/%v/event-templates/%s", query.AppID, eventTemplateID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type TimelineTemplateNewParams struct {
	// State of the template definition being created.
	TimelineEventTemplateCreateRequest TimelineEventTemplateCreateRequestParam
	paramObj
}

func (r TimelineTemplateNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.TimelineEventTemplateCreateRequest)
}
func (r *TimelineTemplateNewParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.TimelineEventTemplateCreateRequest)
}

type TimelineTemplateUpdateParams struct {
	AppID int64 `path:"appId,required" json:"-"`
	// State of the template definition being updated.
	TimelineEventTemplateUpdateRequest TimelineEventTemplateUpdateRequestParam
	paramObj
}

func (r TimelineTemplateUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.TimelineEventTemplateUpdateRequest)
}
func (r *TimelineTemplateUpdateParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.TimelineEventTemplateUpdateRequest)
}

type TimelineTemplateDeleteParams struct {
	AppID int64 `path:"appId,required" json:"-"`
	paramObj
}

type TimelineTemplateGetParams struct {
	AppID int64 `path:"appId,required" json:"-"`
	paramObj
}
