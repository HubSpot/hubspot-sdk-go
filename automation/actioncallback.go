// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package automation

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

// ActionCallbackService contains methods and other services that help with
// interacting with the Hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewActionCallbackService] method instead.
type ActionCallbackService struct {
	Options []option.RequestOption
}

// NewActionCallbackService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewActionCallbackService(opts ...option.RequestOption) (r ActionCallbackService) {
	r = ActionCallbackService{}
	r.Options = opts
	return
}

// Complete a specific blocked action execution by ID.
func (r *ActionCallbackService) Complete(ctx context.Context, callbackID string, body ActionCallbackCompleteParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if callbackID == "" {
		err = errors.New("missing required callbackId parameter")
		return
	}
	path := fmt.Sprintf("automation/v4/actions/callbacks/%s/complete", callbackID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Complete a batch of blocked action executions.
func (r *ActionCallbackService) CompleteBatch(ctx context.Context, body ActionCallbackCompleteBatchParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "automation/v4/actions/callbacks/complete"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

type ActionCallbackCompleteParams struct {
	CallbackCompletionRequest CallbackCompletionRequestParam
	paramObj
}

func (r ActionCallbackCompleteParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.CallbackCompletionRequest)
}
func (r *ActionCallbackCompleteParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.CallbackCompletionRequest)
}

type ActionCallbackCompleteBatchParams struct {
	BatchInputCallbackCompletionBatchRequest BatchInputCallbackCompletionBatchRequestParam
	paramObj
}

func (r ActionCallbackCompleteBatchParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.BatchInputCallbackCompletionBatchRequest)
}
func (r *ActionCallbackCompleteBatchParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.BatchInputCallbackCompletionBatchRequest)
}
