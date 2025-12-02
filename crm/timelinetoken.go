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

// TimelineTokenService contains methods and other services that help with
// interacting with the hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTimelineTokenService] method instead.
type TimelineTokenService struct {
	Options []option.RequestOption
}

// NewTimelineTokenService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewTimelineTokenService(opts ...option.RequestOption) (r TimelineTokenService) {
	r = TimelineTokenService{}
	r.Options = opts
	return
}

// Update an existing event type template with new tokens.
func (r *TimelineTokenService) New(ctx context.Context, eventTemplateID string, params TimelineTokenNewParams, opts ...option.RequestOption) (res *TimelineEventTemplateToken, err error) {
	opts = slices.Concat(r.Options, opts)
	if eventTemplateID == "" {
		err = errors.New("missing required eventTemplateId parameter")
		return
	}
	path := fmt.Sprintf("integrators/timeline/v3/%v/event-templates/%s/tokens", params.AppID, eventTemplateID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Update an event type template token, specified by token name.
func (r *TimelineTokenService) Update(ctx context.Context, tokenName string, params TimelineTokenUpdateParams, opts ...option.RequestOption) (res *TimelineEventTemplateToken, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.EventTemplateID == "" {
		err = errors.New("missing required eventTemplateId parameter")
		return
	}
	if tokenName == "" {
		err = errors.New("missing required tokenName parameter")
		return
	}
	path := fmt.Sprintf("integrators/timeline/v3/%v/event-templates/%s/tokens/%s", params.AppID, params.EventTemplateID, tokenName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return
}

// Delete an existing token from a specific event type template.
func (r *TimelineTokenService) Delete(ctx context.Context, tokenName string, body TimelineTokenDeleteParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if body.EventTemplateID == "" {
		err = errors.New("missing required eventTemplateId parameter")
		return
	}
	if tokenName == "" {
		err = errors.New("missing required tokenName parameter")
		return
	}
	path := fmt.Sprintf("integrators/timeline/v3/%v/event-templates/%s/tokens/%s", body.AppID, body.EventTemplateID, tokenName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

type TimelineTokenNewParams struct {
	AppID int64 `path:"appId,required" json:"-"`
	// State of the token definition.
	TimelineEventTemplateToken TimelineEventTemplateTokenParam
	paramObj
}

func (r TimelineTokenNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.TimelineEventTemplateToken)
}
func (r *TimelineTokenNewParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.TimelineEventTemplateToken)
}

type TimelineTokenUpdateParams struct {
	AppID           int64  `path:"appId,required" json:"-"`
	EventTemplateID string `path:"eventTemplateId,required" json:"-"`
	// State of the token definition for update requests.
	TimelineEventTemplateTokenUpdateRequest TimelineEventTemplateTokenUpdateRequestParam
	paramObj
}

func (r TimelineTokenUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.TimelineEventTemplateTokenUpdateRequest)
}
func (r *TimelineTokenUpdateParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.TimelineEventTemplateTokenUpdateRequest)
}

type TimelineTokenDeleteParams struct {
	AppID           int64  `path:"appId,required" json:"-"`
	EventTemplateID string `path:"eventTemplateId,required" json:"-"`
	paramObj
}
