// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package crm

import (
	"context"
	"net/http"
	"slices"

	"github.com/stainless-sdks/hubspot-sdk-go/internal/apijson"
	shimjson "github.com/stainless-sdks/hubspot-sdk-go/internal/encoding/json"
	"github.com/stainless-sdks/hubspot-sdk-go/internal/requestconfig"
	"github.com/stainless-sdks/hubspot-sdk-go/option"
)

// TimelineBatchService contains methods and other services that help with
// interacting with the hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTimelineBatchService] method instead.
type TimelineBatchService struct {
	options []option.RequestOption
}

// NewTimelineBatchService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewTimelineBatchService(opts ...option.RequestOption) (r TimelineBatchService) {
	r = TimelineBatchService{}
	r.options = opts
	return
}

func (r *TimelineBatchService) New(ctx context.Context, body TimelineBatchNewParams, opts ...option.RequestOption) (res *BatchResponseAppEventOccurrence, err error) {
	opts = slices.Concat(r.options, opts)
	path := "integrators/timeline/2026-03/events/batch"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type TimelineBatchNewParams struct {
	BatchInputAppEventOccurrence BatchInputAppEventOccurrenceParam
	paramObj
}

func (r TimelineBatchNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.BatchInputAppEventOccurrence)
}
func (r *TimelineBatchNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
