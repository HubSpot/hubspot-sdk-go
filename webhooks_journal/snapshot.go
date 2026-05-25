// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package webhooks_journal

import (
	"context"
	"net/http"
	"slices"

	"github.com/HubSpot/hubspot-sdk-go/internal/apijson"
	shimjson "github.com/HubSpot/hubspot-sdk-go/internal/encoding/json"
	"github.com/HubSpot/hubspot-sdk-go/internal/requestconfig"
	"github.com/HubSpot/hubspot-sdk-go/option"
	"github.com/HubSpot/hubspot-sdk-go/shared"
)

// SnapshotService contains methods and other services that help with interacting
// with the hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSnapshotService] method instead.
type SnapshotService struct {
	options []option.RequestOption
}

// NewSnapshotService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewSnapshotService(opts ...option.RequestOption) (r SnapshotService) {
	r = SnapshotService{}
	r.options = opts
	return
}

// Create a batch of CRM object snapshots in HubSpot. This endpoint is used to
// capture the current state of specified CRM objects for later reference or
// analysis. It requires a JSON payload containing the details of the CRM objects
// to snapshot. This operation is exempt from daily and ten-secondly rate limits.
func (r *SnapshotService) New(ctx context.Context, body SnapshotNewParams, opts ...option.RequestOption) (res *shared.CrmObjectSnapshotBatchResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "webhooks-journal/snapshots/2026-03/crm"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type SnapshotNewParams struct {
	CrmObjectSnapshotBatchRequest shared.CrmObjectSnapshotBatchRequestParam
	paramObj
}

func (r SnapshotNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.CrmObjectSnapshotBatchRequest)
}
func (r *SnapshotNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
