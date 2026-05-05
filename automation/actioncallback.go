// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package automation

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/HubSpot/hubspot-sdk-go/internal/apijson"
	shimjson "github.com/HubSpot/hubspot-sdk-go/internal/encoding/json"
	"github.com/HubSpot/hubspot-sdk-go/internal/requestconfig"
	"github.com/HubSpot/hubspot-sdk-go/option"
)

// ActionCallbackService contains methods and other services that help with
// interacting with the hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewActionCallbackService] method instead.
type ActionCallbackService struct {
	options []option.RequestOption
}

// NewActionCallbackService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewActionCallbackService(opts ...option.RequestOption) (r ActionCallbackService) {
	r = ActionCallbackService{}
	r.options = opts
	return
}

// Complete a specific blocked action execution by ID.
func (r *ActionCallbackService) Complete(ctx context.Context, callbackID string, body ActionCallbackCompleteParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if callbackID == "" {
		err = errors.New("missing required callbackId parameter")
		return err
	}
	path := fmt.Sprintf("automation/actions/callbacks/2026-03/%s/complete", url.PathEscape(callbackID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return err
}

// Complete a batch of blocked action executions.
func (r *ActionCallbackService) CompleteBatch(ctx context.Context, body ActionCallbackCompleteBatchParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "automation/actions/callbacks/2026-03/complete"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return err
}

type ActionCallbackCompleteParams struct {
	CallbackCompletionRequest CallbackCompletionRequestParam
	paramObj
}

func (r ActionCallbackCompleteParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.CallbackCompletionRequest)
}
func (r *ActionCallbackCompleteParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ActionCallbackCompleteBatchParams struct {
	BatchInputCallbackCompletionBatchRequest BatchInputCallbackCompletionBatchRequestParam
	paramObj
}

func (r ActionCallbackCompleteBatchParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.BatchInputCallbackCompletionBatchRequest)
}
func (r *ActionCallbackCompleteBatchParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
