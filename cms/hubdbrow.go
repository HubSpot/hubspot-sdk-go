// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cms

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/HubSpot/hubspot-sdk-go/internal/apijson"
	"github.com/HubSpot/hubspot-sdk-go/internal/apiquery"
	shimjson "github.com/HubSpot/hubspot-sdk-go/internal/encoding/json"
	"github.com/HubSpot/hubspot-sdk-go/internal/requestconfig"
	"github.com/HubSpot/hubspot-sdk-go/option"
	"github.com/HubSpot/hubspot-sdk-go/packages/param"
	"github.com/HubSpot/hubspot-sdk-go/shared"
)

// HubdbRowService contains methods and other services that help with interacting
// with the hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewHubdbRowService] method instead.
type HubdbRowService struct {
	options []option.RequestOption
}

// NewHubdbRowService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewHubdbRowService(opts ...option.RequestOption) (r HubdbRowService) {
	r = HubdbRowService{}
	r.options = opts
	return
}

// Add a new row to a HubDB table. New rows will be added to the draft version of
// the table. Use the `/publish` endpoint to push these changes to published
// version.
func (r *HubdbRowService) New(ctx context.Context, tableIDOrName string, body HubdbRowNewParams, opts ...option.RequestOption) (res *HubDBTableRowV3, err error) {
	opts = slices.Concat(r.options, opts)
	if tableIDOrName == "" {
		err = errors.New("missing required tableIdOrName parameter")
		return nil, err
	}
	path := fmt.Sprintf("cms/hubdb/2026-03/tables/%s/rows", url.PathEscape(tableIDOrName))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Returns a set of rows in the published version of the specified table. Row
// results can be filtered and sorted. Filtering and sorting options will be sent
// as query parameters to the API request. For example, by adding the query
// parameters `column1__gt=5&sort=-column1`, API returns the rows with values for
// column `column1` greater than 5 and in the descending order of `column1` values.
// Refer to the
// [overview section](https://developers.hubspot.com/docs/api/cms/hubdb#filtering-and-sorting-table-rows)
// for detailed filtering and sorting options. **Note:** This endpoint can be
// accessed without any authentication, if the table is set to be allowed for
// public access.
func (r *HubdbRowService) List(ctx context.Context, tableIDOrName string, query HubdbRowListParams, opts ...option.RequestOption) (res *UnifiedCollectionResponseWithTotalBaseHubDBTableRowV3Union, err error) {
	opts = slices.Concat(r.options, opts)
	if tableIDOrName == "" {
		err = errors.New("missing required tableIdOrName parameter")
		return nil, err
	}
	path := fmt.Sprintf("cms/hubdb/2026-03/tables/%s/rows", url.PathEscape(tableIDOrName))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Clones rows in the draft version of the specified table, given a set of row ids.
// Maximum of 100 row ids per call.
func (r *HubdbRowService) CloneBatch(ctx context.Context, tableIDOrName string, body HubdbRowCloneBatchParams, opts ...option.RequestOption) (res *BatchResponseHubDBTableRowV3, err error) {
	opts = slices.Concat(r.options, opts)
	if tableIDOrName == "" {
		err = errors.New("missing required tableIdOrName parameter")
		return nil, err
	}
	path := fmt.Sprintf("cms/hubdb/2026-03/tables/%s/rows/draft/batch/clone", url.PathEscape(tableIDOrName))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Clones a single row in the draft version of a table.
func (r *HubdbRowService) CloneDraft(ctx context.Context, rowID string, params HubdbRowCloneDraftParams, opts ...option.RequestOption) (res *HubDBTableRowV3, err error) {
	opts = slices.Concat(r.options, opts)
	if params.TableIDOrName == "" {
		err = errors.New("missing required tableIdOrName parameter")
		return nil, err
	}
	if rowID == "" {
		err = errors.New("missing required rowId parameter")
		return nil, err
	}
	path := fmt.Sprintf("cms/hubdb/2026-03/tables/%s/rows/%s/draft/clone", url.PathEscape(params.TableIDOrName), url.PathEscape(rowID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Creates rows in the draft version of the specified table, given an array of row
// objects. Maximum of 100 row object per call. See the overview section for more
// details with an example.
func (r *HubdbRowService) NewBatch(ctx context.Context, tableIDOrName string, body HubdbRowNewBatchParams, opts ...option.RequestOption) (res *BatchResponseHubDBTableRowV3, err error) {
	opts = slices.Concat(r.options, opts)
	if tableIDOrName == "" {
		err = errors.New("missing required tableIdOrName parameter")
		return nil, err
	}
	path := fmt.Sprintf("cms/hubdb/2026-03/tables/%s/rows/draft/batch/create", url.PathEscape(tableIDOrName))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Permanently deletes a row from a table's draft version.
func (r *HubdbRowService) DeleteDraft(ctx context.Context, rowID string, body HubdbRowDeleteDraftParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if body.TableIDOrName == "" {
		err = errors.New("missing required tableIdOrName parameter")
		return err
	}
	if rowID == "" {
		err = errors.New("missing required rowId parameter")
		return err
	}
	path := fmt.Sprintf("cms/hubdb/2026-03/tables/%s/rows/%s/draft", url.PathEscape(body.TableIDOrName), url.PathEscape(rowID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Get a single row by ID from the published version of a table. **Note:** This
// endpoint can be accessed without any authentication, if the table is set to be
// allowed for public access.
func (r *HubdbRowService) Get(ctx context.Context, rowID string, params HubdbRowGetParams, opts ...option.RequestOption) (res *HubDBTableRowV3, err error) {
	opts = slices.Concat(r.options, opts)
	if params.TableIDOrName == "" {
		err = errors.New("missing required tableIdOrName parameter")
		return nil, err
	}
	if rowID == "" {
		err = errors.New("missing required rowId parameter")
		return nil, err
	}
	path := fmt.Sprintf("cms/hubdb/2026-03/tables/%s/rows/%s", url.PathEscape(params.TableIDOrName), url.PathEscape(rowID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// Returns rows in the published version of the specified table, given a set of row
// IDs. **Note:** This endpoint can be accessed without any authentication if the
// table is set to be allowed for public access.
func (r *HubdbRowService) GetBatch(ctx context.Context, tableIDOrName string, body HubdbRowGetBatchParams, opts ...option.RequestOption) (res *BatchResponseHubDBTableRowV3, err error) {
	opts = slices.Concat(r.options, opts)
	if tableIDOrName == "" {
		err = errors.New("missing required tableIdOrName parameter")
		return nil, err
	}
	path := fmt.Sprintf("cms/hubdb/2026-03/tables/%s/rows/batch/read", url.PathEscape(tableIDOrName))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Get a single row by ID from a table's draft version.
func (r *HubdbRowService) GetDraft(ctx context.Context, rowID string, params HubdbRowGetDraftParams, opts ...option.RequestOption) (res *HubDBTableRowV3, err error) {
	opts = slices.Concat(r.options, opts)
	if params.TableIDOrName == "" {
		err = errors.New("missing required tableIdOrName parameter")
		return nil, err
	}
	if rowID == "" {
		err = errors.New("missing required rowId parameter")
		return nil, err
	}
	path := fmt.Sprintf("cms/hubdb/2026-03/tables/%s/rows/%s/draft", url.PathEscape(params.TableIDOrName), url.PathEscape(rowID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// Returns rows in the draft version of the specified table, given a set of row
// IDs.
func (r *HubdbRowService) GetDraftBatch(ctx context.Context, tableIDOrName string, body HubdbRowGetDraftBatchParams, opts ...option.RequestOption) (res *BatchResponseHubDBTableRowV3, err error) {
	opts = slices.Concat(r.options, opts)
	if tableIDOrName == "" {
		err = errors.New("missing required tableIdOrName parameter")
		return nil, err
	}
	path := fmt.Sprintf("cms/hubdb/2026-03/tables/%s/rows/draft/batch/read", url.PathEscape(tableIDOrName))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Permanently delete rows from the draft version of a table, given a set of row
// IDs. Maximum of 100 row IDs per call.
func (r *HubdbRowService) PurgeBatch(ctx context.Context, tableIDOrName string, body HubdbRowPurgeBatchParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if tableIDOrName == "" {
		err = errors.New("missing required tableIdOrName parameter")
		return err
	}
	path := fmt.Sprintf("cms/hubdb/2026-03/tables/%s/rows/draft/batch/purge", url.PathEscape(tableIDOrName))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return err
}

// Replaces multiple rows as a batch in the draft version of the table, with a
// maximum of 100 rows per call. See the endpoint
// `PUT /tables/{tableIdOrName}/rows/{rowId}/draft` for details on updating a
// single row.
func (r *HubdbRowService) ReplaceBatch(ctx context.Context, tableIDOrName string, body HubdbRowReplaceBatchParams, opts ...option.RequestOption) (res *BatchResponseHubDBTableRowV3, err error) {
	opts = slices.Concat(r.options, opts)
	if tableIDOrName == "" {
		err = errors.New("missing required tableIdOrName parameter")
		return nil, err
	}
	path := fmt.Sprintf("cms/hubdb/2026-03/tables/%s/rows/draft/batch/replace", url.PathEscape(tableIDOrName))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Replace a single row in the draft version of a table. All column values must be
// specified. If a column has a value in the target table and this request doesn't
// define that value, it will be deleted. See the "Create a row" endpoint for
// instructions on how to format the JSON row definitions.
func (r *HubdbRowService) ReplaceDraft(ctx context.Context, rowID string, params HubdbRowReplaceDraftParams, opts ...option.RequestOption) (res *HubDBTableRowV3, err error) {
	opts = slices.Concat(r.options, opts)
	if params.TableIDOrName == "" {
		err = errors.New("missing required tableIdOrName parameter")
		return nil, err
	}
	if rowID == "" {
		err = errors.New("missing required rowId parameter")
		return nil, err
	}
	path := fmt.Sprintf("cms/hubdb/2026-03/tables/%s/rows/%s/draft", url.PathEscape(params.TableIDOrName), url.PathEscape(rowID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return res, err
}

// Updates multiple rows as a batch in the draft version of the table, with a
// maximum of 100 rows per call. See the endpoint
// `PATCH /tables/{tableIdOrName}/rows/{rowId}/draft` for details on updating a
// single row.
func (r *HubdbRowService) UpdateBatch(ctx context.Context, tableIDOrName string, body HubdbRowUpdateBatchParams, opts ...option.RequestOption) (res *BatchResponseHubDBTableRowV3, err error) {
	opts = slices.Concat(r.options, opts)
	if tableIDOrName == "" {
		err = errors.New("missing required tableIdOrName parameter")
		return nil, err
	}
	path := fmt.Sprintf("cms/hubdb/2026-03/tables/%s/rows/draft/batch/update", url.PathEscape(tableIDOrName))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Partially update a single row in the table's draft version. All the column
// values need not be specified. Only the columns or fields that needs to be
// modified can be specified. See the "Create a row" endpoint for instructions on
// how to format the JSON row definitions.
func (r *HubdbRowService) UpdateDraft(ctx context.Context, rowID string, params HubdbRowUpdateDraftParams, opts ...option.RequestOption) (res *HubDBTableRowV3, err error) {
	opts = slices.Concat(r.options, opts)
	if params.TableIDOrName == "" {
		err = errors.New("missing required tableIdOrName parameter")
		return nil, err
	}
	if rowID == "" {
		err = errors.New("missing required rowId parameter")
		return nil, err
	}
	path := fmt.Sprintf("cms/hubdb/2026-03/tables/%s/rows/%s/draft", url.PathEscape(params.TableIDOrName), url.PathEscape(rowID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return res, err
}

type HubdbRowNewParams struct {
	HubDBTableRowV3Request HubDBTableRowV3RequestParam
	paramObj
}

func (r HubdbRowNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.HubDBTableRowV3Request)
}
func (r *HubdbRowNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type HubdbRowListParams struct {
	// The paging cursor token of the last successfully read resource will be returned
	// as the `paging.next.after` JSON property of a paged response containing more
	// results.
	After param.Opt[string] `query:"after,omitzero" json:"-"`
	// Whether to return only results that have been archived.
	Archived param.Opt[bool] `query:"archived,omitzero" json:"-"`
	// The maximum number of results to display per page.
	Limit      param.Opt[int64] `query:"limit,omitzero" json:"-"`
	Offset     param.Opt[int64] `query:"offset,omitzero" json:"-"`
	Properties []string         `query:"properties,omitzero" json:"-"`
	Sort       []string         `query:"sort,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [HubdbRowListParams]'s query parameters as `url.Values`.
func (r HubdbRowListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type HubdbRowCloneBatchParams struct {
	BatchInputHubDBTableRowBatchCloneRequest BatchInputHubDBTableRowBatchCloneRequestParam
	paramObj
}

func (r HubdbRowCloneBatchParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.BatchInputHubDBTableRowBatchCloneRequest)
}
func (r *HubdbRowCloneBatchParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type HubdbRowCloneDraftParams struct {
	TableIDOrName string            `path:"tableIdOrName" api:"required" json:"-"`
	Name          param.Opt[string] `query:"name,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [HubdbRowCloneDraftParams]'s query parameters as
// `url.Values`.
func (r HubdbRowCloneDraftParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type HubdbRowNewBatchParams struct {
	BatchInputHubDBTableRowV3Request BatchInputHubDBTableRowV3RequestParam
	paramObj
}

func (r HubdbRowNewBatchParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.BatchInputHubDBTableRowV3Request)
}
func (r *HubdbRowNewBatchParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type HubdbRowDeleteDraftParams struct {
	TableIDOrName string `path:"tableIdOrName" api:"required" json:"-"`
	paramObj
}

type HubdbRowGetParams struct {
	TableIDOrName string `path:"tableIdOrName" api:"required" json:"-"`
	// Whether to return only results that have been archived.
	Archived param.Opt[bool] `query:"archived,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [HubdbRowGetParams]'s query parameters as `url.Values`.
func (r HubdbRowGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type HubdbRowGetBatchParams struct {
	BatchInputString shared.BatchInputStringParam
	paramObj
}

func (r HubdbRowGetBatchParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.BatchInputString)
}
func (r *HubdbRowGetBatchParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type HubdbRowGetDraftParams struct {
	TableIDOrName string `path:"tableIdOrName" api:"required" json:"-"`
	// Whether to return only results that have been archived.
	Archived param.Opt[bool] `query:"archived,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [HubdbRowGetDraftParams]'s query parameters as `url.Values`.
func (r HubdbRowGetDraftParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type HubdbRowGetDraftBatchParams struct {
	BatchInputString shared.BatchInputStringParam
	paramObj
}

func (r HubdbRowGetDraftBatchParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.BatchInputString)
}
func (r *HubdbRowGetDraftBatchParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type HubdbRowPurgeBatchParams struct {
	BatchInputString shared.BatchInputStringParam
	paramObj
}

func (r HubdbRowPurgeBatchParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.BatchInputString)
}
func (r *HubdbRowPurgeBatchParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type HubdbRowReplaceBatchParams struct {
	BatchInputHubDBTableRowV3BatchUpdateRequest BatchInputHubDBTableRowV3BatchUpdateRequestParam
	paramObj
}

func (r HubdbRowReplaceBatchParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.BatchInputHubDBTableRowV3BatchUpdateRequest)
}
func (r *HubdbRowReplaceBatchParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type HubdbRowReplaceDraftParams struct {
	TableIDOrName          string `path:"tableIdOrName" api:"required" json:"-"`
	HubDBTableRowV3Request HubDBTableRowV3RequestParam
	paramObj
}

func (r HubdbRowReplaceDraftParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.HubDBTableRowV3Request)
}
func (r *HubdbRowReplaceDraftParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type HubdbRowUpdateBatchParams struct {
	BatchInputHubDBTableRowV3BatchUpdateRequest BatchInputHubDBTableRowV3BatchUpdateRequestParam
	paramObj
}

func (r HubdbRowUpdateBatchParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.BatchInputHubDBTableRowV3BatchUpdateRequest)
}
func (r *HubdbRowUpdateBatchParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type HubdbRowUpdateDraftParams struct {
	TableIDOrName          string `path:"tableIdOrName" api:"required" json:"-"`
	HubDBTableRowV3Request HubDBTableRowV3RequestParam
	paramObj
}

func (r HubdbRowUpdateDraftParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.HubDBTableRowV3Request)
}
func (r *HubdbRowUpdateDraftParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
