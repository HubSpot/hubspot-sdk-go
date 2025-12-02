// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cms

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
	"github.com/stainless-sdks/hubspot-sdk-go/shared"
)

// HubdbRowBatchService contains methods and other services that help with
// interacting with the hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewHubdbRowBatchService] method instead.
type HubdbRowBatchService struct {
	Options []option.RequestOption
}

// NewHubdbRowBatchService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewHubdbRowBatchService(opts ...option.RequestOption) (r HubdbRowBatchService) {
	r = HubdbRowBatchService{}
	r.Options = opts
	return
}

// Clones rows in the draft version of the specified table, given a set of row ids.
// Maximum of 100 row ids per call.
func (r *HubdbRowBatchService) CloneBatch(ctx context.Context, tableIDOrName string, body HubdbRowBatchCloneBatchParams, opts ...option.RequestOption) (res *BatchResponseHubDBTableRowV3, err error) {
	opts = slices.Concat(r.Options, opts)
	if tableIDOrName == "" {
		err = errors.New("missing required tableIdOrName parameter")
		return
	}
	path := fmt.Sprintf("cms/v3/hubdb/tables/%s/rows/draft/batch/clone", tableIDOrName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Creates rows in the draft version of the specified table, given an array of row
// objects. Maximum of 100 row object per call. See the overview section for more
// details with an example.
func (r *HubdbRowBatchService) NewBatch(ctx context.Context, tableIDOrName string, body HubdbRowBatchNewBatchParams, opts ...option.RequestOption) (res *BatchResponseHubDBTableRowV3, err error) {
	opts = slices.Concat(r.Options, opts)
	if tableIDOrName == "" {
		err = errors.New("missing required tableIdOrName parameter")
		return
	}
	path := fmt.Sprintf("cms/v3/hubdb/tables/%s/rows/draft/batch/create", tableIDOrName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Returns rows in the published version of the specified table, given a set of row
// IDs. **Note:** This endpoint can be accessed without any authentication if the
// table is set to be allowed for public access.
func (r *HubdbRowBatchService) GetBatch(ctx context.Context, tableIDOrName string, body HubdbRowBatchGetBatchParams, opts ...option.RequestOption) (res *BatchResponseHubDBTableRowV3, err error) {
	opts = slices.Concat(r.Options, opts)
	if tableIDOrName == "" {
		err = errors.New("missing required tableIdOrName parameter")
		return
	}
	path := fmt.Sprintf("cms/v3/hubdb/tables/%s/rows/batch/read", tableIDOrName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Returns rows in the draft version of the specified table, given a set of row
// IDs.
func (r *HubdbRowBatchService) GetDraftBatch(ctx context.Context, tableIDOrName string, body HubdbRowBatchGetDraftBatchParams, opts ...option.RequestOption) (res *BatchResponseHubDBTableRowV3, err error) {
	opts = slices.Concat(r.Options, opts)
	if tableIDOrName == "" {
		err = errors.New("missing required tableIdOrName parameter")
		return
	}
	path := fmt.Sprintf("cms/v3/hubdb/tables/%s/rows/draft/batch/read", tableIDOrName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Permanently deletes rows from the draft version of the table, given a set of row
// IDs. Maximum of 100 row IDs per call.
func (r *HubdbRowBatchService) PurgeBatch(ctx context.Context, tableIDOrName string, body HubdbRowBatchPurgeBatchParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if tableIDOrName == "" {
		err = errors.New("missing required tableIdOrName parameter")
		return
	}
	path := fmt.Sprintf("cms/v3/hubdb/tables/%s/rows/draft/batch/purge", tableIDOrName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Replaces multiple rows as a batch in the draft version of the table, with a
// maximum of 100 rows per call. See the endpoint
// `PUT /tables/{tableIdOrName}/rows/{rowId}/draft` for details on updating a
// single row.
func (r *HubdbRowBatchService) ReplaceBatch(ctx context.Context, tableIDOrName string, body HubdbRowBatchReplaceBatchParams, opts ...option.RequestOption) (res *BatchResponseHubDBTableRowV3, err error) {
	opts = slices.Concat(r.Options, opts)
	if tableIDOrName == "" {
		err = errors.New("missing required tableIdOrName parameter")
		return
	}
	path := fmt.Sprintf("cms/v3/hubdb/tables/%s/rows/draft/batch/replace", tableIDOrName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Updates multiple rows as a batch in the draft version of the table, with a
// maximum of 100 rows per call. See the endpoint
// `PATCH /tables/{tableIdOrName}/rows/{rowId}/draft` for details on updating a
// single row.
func (r *HubdbRowBatchService) UpdateBatch(ctx context.Context, tableIDOrName string, body HubdbRowBatchUpdateBatchParams, opts ...option.RequestOption) (res *BatchResponseHubDBTableRowV3, err error) {
	opts = slices.Concat(r.Options, opts)
	if tableIDOrName == "" {
		err = errors.New("missing required tableIdOrName parameter")
		return
	}
	path := fmt.Sprintf("cms/v3/hubdb/tables/%s/rows/draft/batch/update", tableIDOrName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type HubdbRowBatchCloneBatchParams struct {
	BatchInputHubDBTableRowBatchCloneRequest BatchInputHubDBTableRowBatchCloneRequestParam
	paramObj
}

func (r HubdbRowBatchCloneBatchParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.BatchInputHubDBTableRowBatchCloneRequest)
}
func (r *HubdbRowBatchCloneBatchParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.BatchInputHubDBTableRowBatchCloneRequest)
}

type HubdbRowBatchNewBatchParams struct {
	BatchInputHubDBTableRowV3Request BatchInputHubDBTableRowV3RequestParam
	paramObj
}

func (r HubdbRowBatchNewBatchParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.BatchInputHubDBTableRowV3Request)
}
func (r *HubdbRowBatchNewBatchParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.BatchInputHubDBTableRowV3Request)
}

type HubdbRowBatchGetBatchParams struct {
	// Wrapper for providing an array of strings as inputs.
	BatchInputString shared.BatchInputStringParam
	paramObj
}

func (r HubdbRowBatchGetBatchParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.BatchInputString)
}
func (r *HubdbRowBatchGetBatchParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.BatchInputString)
}

type HubdbRowBatchGetDraftBatchParams struct {
	// Wrapper for providing an array of strings as inputs.
	BatchInputString shared.BatchInputStringParam
	paramObj
}

func (r HubdbRowBatchGetDraftBatchParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.BatchInputString)
}
func (r *HubdbRowBatchGetDraftBatchParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.BatchInputString)
}

type HubdbRowBatchPurgeBatchParams struct {
	// Wrapper for providing an array of strings as inputs.
	BatchInputString shared.BatchInputStringParam
	paramObj
}

func (r HubdbRowBatchPurgeBatchParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.BatchInputString)
}
func (r *HubdbRowBatchPurgeBatchParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.BatchInputString)
}

type HubdbRowBatchReplaceBatchParams struct {
	BatchInputHubDBTableRowV3BatchUpdateRequest BatchInputHubDBTableRowV3BatchUpdateRequestParam
	paramObj
}

func (r HubdbRowBatchReplaceBatchParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.BatchInputHubDBTableRowV3BatchUpdateRequest)
}
func (r *HubdbRowBatchReplaceBatchParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.BatchInputHubDBTableRowV3BatchUpdateRequest)
}

type HubdbRowBatchUpdateBatchParams struct {
	BatchInputHubDBTableRowV3BatchUpdateRequest BatchInputHubDBTableRowV3BatchUpdateRequestParam
	paramObj
}

func (r HubdbRowBatchUpdateBatchParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.BatchInputHubDBTableRowV3BatchUpdateRequest)
}
func (r *HubdbRowBatchUpdateBatchParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.BatchInputHubDBTableRowV3BatchUpdateRequest)
}
