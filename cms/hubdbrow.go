// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cms

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/stainless-sdks/hubspot-sdk-go/internal/apiquery"
	shimjson "github.com/stainless-sdks/hubspot-sdk-go/internal/encoding/json"
	"github.com/stainless-sdks/hubspot-sdk-go/internal/requestconfig"
	"github.com/stainless-sdks/hubspot-sdk-go/option"
	"github.com/stainless-sdks/hubspot-sdk-go/packages/pagination"
	"github.com/stainless-sdks/hubspot-sdk-go/packages/param"
	"github.com/stainless-sdks/hubspot-sdk-go/shared"
)

// HubdbRowService contains methods and other services that help with interacting
// with the Hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewHubdbRowService] method instead.
type HubdbRowService struct {
	Options []option.RequestOption
	Batch   HubdbRowBatchService
}

// NewHubdbRowService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewHubdbRowService(opts ...option.RequestOption) (r HubdbRowService) {
	r = HubdbRowService{}
	r.Options = opts
	r.Batch = NewHubdbRowBatchService(opts...)
	return
}

// Add a new row to a HubDB table. New rows will be added to the draft version of
// the table. Use the `/publish` endpoint to push these changes to published
// version.
func (r *HubdbRowService) New(ctx context.Context, tableIDOrName string, body HubdbRowNewParams, opts ...option.RequestOption) (res *HubDBTableRowV3, err error) {
	opts = slices.Concat(r.Options, opts)
	if tableIDOrName == "" {
		err = errors.New("missing required tableIdOrName parameter")
		return
	}
	path := fmt.Sprintf("cms/v3/hubdb/tables/%s/rows", tableIDOrName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
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
func (r *HubdbRowService) List(ctx context.Context, tableIDOrName string, query HubdbRowListParams, opts ...option.RequestOption) (res *pagination.Page[shared.HubDBTableRowV3Wrapper], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if tableIDOrName == "" {
		err = errors.New("missing required tableIdOrName parameter")
		return
	}
	path := fmt.Sprintf("cms/v3/hubdb/tables/%s/rows", tableIDOrName)
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
func (r *HubdbRowService) ListAutoPaging(ctx context.Context, tableIDOrName string, query HubdbRowListParams, opts ...option.RequestOption) *pagination.PageAutoPager[shared.HubDBTableRowV3Wrapper] {
	return pagination.NewPageAutoPager(r.List(ctx, tableIDOrName, query, opts...))
}

// Clones a single row in the draft version of a table.
func (r *HubdbRowService) CloneDraft(ctx context.Context, rowID string, params HubdbRowCloneDraftParams, opts ...option.RequestOption) (res *HubDBTableRowV3, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.TableIDOrName == "" {
		err = errors.New("missing required tableIdOrName parameter")
		return
	}
	if rowID == "" {
		err = errors.New("missing required rowId parameter")
		return
	}
	path := fmt.Sprintf("cms/v3/hubdb/tables/%s/rows/%s/draft/clone", params.TableIDOrName, rowID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Permanently deletes a row from a table's draft version.
func (r *HubdbRowService) DeleteDraft(ctx context.Context, rowID string, body HubdbRowDeleteDraftParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if body.TableIDOrName == "" {
		err = errors.New("missing required tableIdOrName parameter")
		return
	}
	if rowID == "" {
		err = errors.New("missing required rowId parameter")
		return
	}
	path := fmt.Sprintf("cms/v3/hubdb/tables/%s/rows/%s/draft", body.TableIDOrName, rowID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Get a single row by ID from the published version of a table. **Note:** This
// endpoint can be accessed without any authentication, if the table is set to be
// allowed for public access.
func (r *HubdbRowService) Get(ctx context.Context, rowID string, params HubdbRowGetParams, opts ...option.RequestOption) (res *HubDBTableRowV3, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.TableIDOrName == "" {
		err = errors.New("missing required tableIdOrName parameter")
		return
	}
	if rowID == "" {
		err = errors.New("missing required rowId parameter")
		return
	}
	path := fmt.Sprintf("cms/v3/hubdb/tables/%s/rows/%s", params.TableIDOrName, rowID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return
}

// Get a single row by ID from a table's draft version.
func (r *HubdbRowService) GetDraft(ctx context.Context, rowID string, params HubdbRowGetDraftParams, opts ...option.RequestOption) (res *HubDBTableRowV3, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.TableIDOrName == "" {
		err = errors.New("missing required tableIdOrName parameter")
		return
	}
	if rowID == "" {
		err = errors.New("missing required rowId parameter")
		return
	}
	path := fmt.Sprintf("cms/v3/hubdb/tables/%s/rows/%s/draft", params.TableIDOrName, rowID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return
}

// Returns rows in the draft version of the specified table. Row results can be
// filtered and sorted. Filtering and sorting options will be sent as query
// parameters to the API request. For example, by adding the query parameters
// `column1__gt=5&sort=-column1`, API returns the rows with values for column
// `column1` greater than 5 and in the descending order of `column1` values. Refer
// to the
// [overview section](https://developers.hubspot.com/docs/api/cms/hubdb#filtering-and-sorting-table-rows)
// for detailed filtering and sorting options.
func (r *HubdbRowService) ListDraft(ctx context.Context, tableIDOrName string, query HubdbRowListDraftParams, opts ...option.RequestOption) (res *UnifiedCollectionResponseWithTotalBaseHubDBTableRowV3Union, err error) {
	opts = slices.Concat(r.Options, opts)
	if tableIDOrName == "" {
		err = errors.New("missing required tableIdOrName parameter")
		return
	}
	path := fmt.Sprintf("cms/v3/hubdb/tables/%s/rows/draft", tableIDOrName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Replace a single row in the draft version of a table. All column values must be
// specified. If a column has a value in the target table and this request doesn't
// define that value, it will be deleted. See the "Create a row" endpoint for
// instructions on how to format the JSON row definitions.
func (r *HubdbRowService) ReplaceDraft(ctx context.Context, rowID string, params HubdbRowReplaceDraftParams, opts ...option.RequestOption) (res *HubDBTableRowV3, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.TableIDOrName == "" {
		err = errors.New("missing required tableIdOrName parameter")
		return
	}
	if rowID == "" {
		err = errors.New("missing required rowId parameter")
		return
	}
	path := fmt.Sprintf("cms/v3/hubdb/tables/%s/rows/%s/draft", params.TableIDOrName, rowID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return
}

// Sparse updates a single row in the table's draft version. All the column values
// need not be specified. Only the columns or fields that needs to be modified can
// be specified. See the "Create a row" endpoint for instructions on how to format
// the JSON row definitions.
func (r *HubdbRowService) UpdateDraft(ctx context.Context, rowID string, params HubdbRowUpdateDraftParams, opts ...option.RequestOption) (res *HubDBTableRowV3, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.TableIDOrName == "" {
		err = errors.New("missing required tableIdOrName parameter")
		return
	}
	if rowID == "" {
		err = errors.New("missing required rowId parameter")
		return
	}
	path := fmt.Sprintf("cms/v3/hubdb/tables/%s/rows/%s/draft", params.TableIDOrName, rowID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return
}

type HubdbRowNewParams struct {
	HubDBTableRowV3Request HubDBTableRowV3RequestParam
	paramObj
}

func (r HubdbRowNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.HubDBTableRowV3Request)
}
func (r *HubdbRowNewParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.HubDBTableRowV3Request)
}

type HubdbRowListParams struct {
	// The cursor token value to get the next set of results. You can get this from the
	// `paging.next.after` JSON property of a paged response containing more results.
	After    param.Opt[string] `query:"after,omitzero" json:"-"`
	Archived param.Opt[bool]   `query:"archived,omitzero" json:"-"`
	// The maximum number of results to return. Default is `1000`.
	Limit  param.Opt[int64] `query:"limit,omitzero" json:"-"`
	Offset param.Opt[int64] `query:"offset,omitzero" json:"-"`
	// Specify the column names to get results containing only the required columns
	// instead of all column details.
	Properties []string `query:"properties,omitzero" json:"-"`
	// Specifies the column names to sort the results by. See the above description for
	// more details.
	Sort []string `query:"sort,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [HubdbRowListParams]'s query parameters as `url.Values`.
func (r HubdbRowListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type HubdbRowCloneDraftParams struct {
	TableIDOrName string            `path:"tableIdOrName,required" json:"-"`
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

type HubdbRowDeleteDraftParams struct {
	TableIDOrName string `path:"tableIdOrName,required" json:"-"`
	paramObj
}

type HubdbRowGetParams struct {
	TableIDOrName string          `path:"tableIdOrName,required" json:"-"`
	Archived      param.Opt[bool] `query:"archived,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [HubdbRowGetParams]'s query parameters as `url.Values`.
func (r HubdbRowGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type HubdbRowGetDraftParams struct {
	TableIDOrName string          `path:"tableIdOrName,required" json:"-"`
	Archived      param.Opt[bool] `query:"archived,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [HubdbRowGetDraftParams]'s query parameters as `url.Values`.
func (r HubdbRowGetDraftParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type HubdbRowListDraftParams struct {
	// The cursor token value to get the next set of results. You can get this from the
	// `paging.next.after` JSON property of a paged response containing more results.
	After    param.Opt[string] `query:"after,omitzero" json:"-"`
	Archived param.Opt[bool]   `query:"archived,omitzero" json:"-"`
	// The maximum number of results to return. Default is `1000`.
	Limit  param.Opt[int64] `query:"limit,omitzero" json:"-"`
	Offset param.Opt[int64] `query:"offset,omitzero" json:"-"`
	// Specify the column names to get results containing only the required columns
	// instead of all column details. If you want to include multiple columns in the
	// result, use this query param as many times.
	Properties []string `query:"properties,omitzero" json:"-"`
	// Specifies the column names to sort the results by.
	Sort []string `query:"sort,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [HubdbRowListDraftParams]'s query parameters as
// `url.Values`.
func (r HubdbRowListDraftParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type HubdbRowReplaceDraftParams struct {
	TableIDOrName          string `path:"tableIdOrName,required" json:"-"`
	HubDBTableRowV3Request HubDBTableRowV3RequestParam
	paramObj
}

func (r HubdbRowReplaceDraftParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.HubDBTableRowV3Request)
}
func (r *HubdbRowReplaceDraftParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.HubDBTableRowV3Request)
}

type HubdbRowUpdateDraftParams struct {
	TableIDOrName          string `path:"tableIdOrName,required" json:"-"`
	HubDBTableRowV3Request HubDBTableRowV3RequestParam
	paramObj
}

func (r HubdbRowUpdateDraftParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.HubDBTableRowV3Request)
}
func (r *HubdbRowUpdateDraftParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.HubDBTableRowV3Request)
}
