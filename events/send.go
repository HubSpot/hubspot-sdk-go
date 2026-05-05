// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package events

import (
	"context"
	"net/http"
	"slices"

	"github.com/HubSpot/hubspot-sdk-go/internal/apijson"
	shimjson "github.com/HubSpot/hubspot-sdk-go/internal/encoding/json"
	"github.com/HubSpot/hubspot-sdk-go/internal/requestconfig"
	"github.com/HubSpot/hubspot-sdk-go/option"
)

// SendService contains methods and other services that help with interacting with
// the hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSendService] method instead.
type SendService struct {
	options []option.RequestOption
}

// NewSendService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewSendService(opts ...option.RequestOption) (r SendService) {
	r = SendService{}
	r.options = opts
	return
}

// Send multiple event occurrences at once.
func (r *SendService) BatchSend(ctx context.Context, body SendBatchSendParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "events/2026-03/send/batch"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return err
}

// Send data for a single custom event occurrence.
func (r *SendService) Send(ctx context.Context, body SendSendParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "events/2026-03/send"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return err
}

type SendBatchSendParams struct {
	BatchedBehavioralEventHTTPCompletionRequest BatchedBehavioralEventHTTPCompletionRequestParam
	paramObj
}

func (r SendBatchSendParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.BatchedBehavioralEventHTTPCompletionRequest)
}
func (r *SendBatchSendParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SendSendParams struct {
	BehavioralEventHTTPCompletionRequest BehavioralEventHTTPCompletionRequestParam
	paramObj
}

func (r SendSendParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.BehavioralEventHTTPCompletionRequest)
}
func (r *SendSendParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
