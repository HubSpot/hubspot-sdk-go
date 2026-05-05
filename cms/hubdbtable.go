// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cms

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/HubSpot/hubspot-sdk-go/internal/apiform"
	"github.com/HubSpot/hubspot-sdk-go/internal/apijson"
	"github.com/HubSpot/hubspot-sdk-go/internal/apiquery"
	shimjson "github.com/HubSpot/hubspot-sdk-go/internal/encoding/json"
	"github.com/HubSpot/hubspot-sdk-go/internal/requestconfig"
	"github.com/HubSpot/hubspot-sdk-go/option"
	"github.com/HubSpot/hubspot-sdk-go/packages/pagination"
	"github.com/HubSpot/hubspot-sdk-go/packages/param"
)

// HubdbTableService contains methods and other services that help with interacting
// with the hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewHubdbTableService] method instead.
type HubdbTableService struct {
	options []option.RequestOption
}

// NewHubdbTableService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewHubdbTableService(opts ...option.RequestOption) (r HubdbTableService) {
	r = HubdbTableService{}
	r.options = opts
	return
}

// Creates a new draft HubDB table given a JSON schema. The table name and label
// should be unique for each account.
func (r *HubdbTableService) New(ctx context.Context, body HubdbTableNewParams, opts ...option.RequestOption) (res *HubDBTableV3, err error) {
	opts = slices.Concat(r.options, opts)
	path := "cms/hubdb/2026-03/tables"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Returns the details for the published version of each table defined in an
// account, including column definitions.
func (r *HubdbTableService) List(ctx context.Context, query HubdbTableListParams, opts ...option.RequestOption) (res *pagination.Page[HubDBTableV3], err error) {
	var raw *http.Response
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "cms/hubdb/2026-03/tables"
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

// Returns the details for the published version of each table defined in an
// account, including column definitions.
func (r *HubdbTableService) ListAutoPaging(ctx context.Context, query HubdbTableListParams, opts ...option.RequestOption) *pagination.PageAutoPager[HubDBTableV3] {
	return pagination.NewPageAutoPager(r.List(ctx, query, opts...))
}

// Archive (soft delete) an existing HubDB table. This archives both the published
// and draft versions.
func (r *HubdbTableService) Delete(ctx context.Context, tableIDOrName string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if tableIDOrName == "" {
		err = errors.New("missing required tableIdOrName parameter")
		return err
	}
	path := fmt.Sprintf("cms/hubdb/2026-03/tables/%s", url.PathEscape(tableIDOrName))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Clone an existing HubDB table. The `newName` and `newLabel` of the new table can
// be sent as JSON in the request body. This will create the cloned table as a
// draft.
func (r *HubdbTableService) CloneDraft(ctx context.Context, tableIDOrName string, body HubdbTableCloneDraftParams, opts ...option.RequestOption) (res *HubDBTableV3, err error) {
	opts = slices.Concat(r.options, opts)
	if tableIDOrName == "" {
		err = errors.New("missing required tableIdOrName parameter")
		return nil, err
	}
	path := fmt.Sprintf("cms/hubdb/2026-03/tables/%s/draft/clone", url.PathEscape(tableIDOrName))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Delete a specific version of a table
func (r *HubdbTableService) DeleteVersion(ctx context.Context, versionID int64, body HubdbTableDeleteVersionParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if body.TableIDOrName == "" {
		err = errors.New("missing required tableIdOrName parameter")
		return err
	}
	path := fmt.Sprintf("cms/hubdb/2026-03/tables/%s/versions/%v", url.PathEscape(body.TableIDOrName), versionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Exports the published version of a table in a specified format.
func (r *HubdbTableService) Export(ctx context.Context, tableIDOrName string, query HubdbTableExportParams, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/vnd.ms-excel")}, opts...)
	if tableIDOrName == "" {
		err = errors.New("missing required tableIdOrName parameter")
		return nil, err
	}
	path := fmt.Sprintf("cms/hubdb/2026-03/tables/%s/export", url.PathEscape(tableIDOrName))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Exports the draft version of a table to CSV / EXCEL format.
func (r *HubdbTableService) ExportDraft(ctx context.Context, tableIDOrName string, query HubdbTableExportDraftParams, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/vnd.ms-excel")}, opts...)
	if tableIDOrName == "" {
		err = errors.New("missing required tableIdOrName parameter")
		return nil, err
	}
	path := fmt.Sprintf("cms/hubdb/2026-03/tables/%s/draft/export", url.PathEscape(tableIDOrName))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Returns the details for the published version of the specified table. This will
// include the definitions for the columns in the table and the number of rows in
// the table.
//
// **Note:** This endpoint can be accessed without any authentication if the table
// is set to be allowed for public access. To do so, you'll need to include the
// HubSpot account ID in a `portalId` query parameter.
func (r *HubdbTableService) Get(ctx context.Context, tableIDOrName string, query HubdbTableGetParams, opts ...option.RequestOption) (res *HubDBTableV3, err error) {
	opts = slices.Concat(r.options, opts)
	if tableIDOrName == "" {
		err = errors.New("missing required tableIdOrName parameter")
		return nil, err
	}
	path := fmt.Sprintf("cms/hubdb/2026-03/tables/%s", url.PathEscape(tableIDOrName))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Get the details for the draft version of a specific HubDB table. This will
// include the definitions for the columns in the table and the number of rows in
// the table.
func (r *HubdbTableService) GetDraft(ctx context.Context, tableIDOrName string, query HubdbTableGetDraftParams, opts ...option.RequestOption) (res *HubDBTableV3, err error) {
	opts = slices.Concat(r.options, opts)
	if tableIDOrName == "" {
		err = errors.New("missing required tableIdOrName parameter")
		return nil, err
	}
	path := fmt.Sprintf("cms/hubdb/2026-03/tables/%s/draft", url.PathEscape(tableIDOrName))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Import the contents of a CSV file into an existing HubDB table. The data will
// always be imported into the draft version of the table. Use the `/publish`
// endpoint to push these changes to the published version. This endpoint takes a
// multi-part POST request. The first part will be a set of JSON-formatted options
// for the import and you can specify this with the name as `config`. The second
// part will be the CSV file you want to import and you can specify this with the
// name as `file`. Refer the
// [overview section](https://developers.hubspot.com/docs/api/cms/hubdb#importing-tables)
// to check the details and format of the JSON-formatted options for the import.
func (r *HubdbTableService) ImportDraft(ctx context.Context, tableIDOrName string, body HubdbTableImportDraftParams, opts ...option.RequestOption) (res *ImportResult, err error) {
	opts = slices.Concat(r.options, opts)
	if tableIDOrName == "" {
		err = errors.New("missing required tableIdOrName parameter")
		return nil, err
	}
	path := fmt.Sprintf("cms/hubdb/2026-03/tables/%s/draft/import", url.PathEscape(tableIDOrName))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Returns the details for each draft table defined in the specified account,
// including column definitions.
func (r *HubdbTableService) ListDraft(ctx context.Context, query HubdbTableListDraftParams, opts ...option.RequestOption) (res *pagination.Page[HubDBTableV3], err error) {
	var raw *http.Response
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "cms/hubdb/2026-03/tables/draft"
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

// Returns the details for each draft table defined in the specified account,
// including column definitions.
func (r *HubdbTableService) ListDraftAutoPaging(ctx context.Context, query HubdbTableListDraftParams, opts ...option.RequestOption) *pagination.PageAutoPager[HubDBTableV3] {
	return pagination.NewPageAutoPager(r.ListDraft(ctx, query, opts...))
}

// Publishes the table by copying the data and table schema changes from draft
// version to the published version, meaning any website pages using data from the
// table will be updated.
func (r *HubdbTableService) PublishDraft(ctx context.Context, tableIDOrName string, body HubdbTablePublishDraftParams, opts ...option.RequestOption) (res *HubDBTableV3, err error) {
	opts = slices.Concat(r.options, opts)
	if tableIDOrName == "" {
		err = errors.New("missing required tableIdOrName parameter")
		return nil, err
	}
	path := fmt.Sprintf("cms/hubdb/2026-03/tables/%s/draft/publish", url.PathEscape(tableIDOrName))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Replaces the data in the draft version of the table with values from the
// published version. Any unpublished changes in the draft will be lost after this
// call is made.
func (r *HubdbTableService) ResetDraft(ctx context.Context, tableIDOrName string, body HubdbTableResetDraftParams, opts ...option.RequestOption) (res *HubDBTableV3, err error) {
	opts = slices.Concat(r.options, opts)
	if tableIDOrName == "" {
		err = errors.New("missing required tableIdOrName parameter")
		return nil, err
	}
	path := fmt.Sprintf("cms/hubdb/2026-03/tables/%s/draft/reset", url.PathEscape(tableIDOrName))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Unpublishes the table, meaning any website pages using data from the table will
// not render any data.
func (r *HubdbTableService) Unpublish(ctx context.Context, tableIDOrName string, body HubdbTableUnpublishParams, opts ...option.RequestOption) (res *HubDBTableV3, err error) {
	opts = slices.Concat(r.options, opts)
	if tableIDOrName == "" {
		err = errors.New("missing required tableIdOrName parameter")
		return nil, err
	}
	path := fmt.Sprintf("cms/hubdb/2026-03/tables/%s/unpublish", url.PathEscape(tableIDOrName))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Update an existing HubDB table. You can use this endpoint to add or remove
// columns to the table as well as restore an archived table. Tables updated using
// the endpoint will only modify the draft verion of the table. Use the `/publish`
// endpoint to push all the changes to the published version. To restore a table,
// include the query parameter `archived=true` and `"archived": false` in the json
// body. **Note:** You need to include all the columns in the input when you are
// adding/removing/updating a column. If you do not include an already existing
// column in the request, it will be deleted.
func (r *HubdbTableService) UpdateDraft(ctx context.Context, tableIDOrName string, params HubdbTableUpdateDraftParams, opts ...option.RequestOption) (res *HubDBTableV3, err error) {
	opts = slices.Concat(r.options, opts)
	if tableIDOrName == "" {
		err = errors.New("missing required tableIdOrName parameter")
		return nil, err
	}
	path := fmt.Sprintf("cms/hubdb/2026-03/tables/%s/draft", url.PathEscape(tableIDOrName))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return res, err
}

type HubdbTableNewParams struct {
	HubDBTableV3Request HubDBTableV3RequestParam
	paramObj
}

func (r HubdbTableNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.HubDBTableV3Request)
}
func (r *HubdbTableNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type HubdbTableListParams struct {
	// The paging cursor token of the last successfully read resource will be returned
	// as the `paging.next.after` JSON property of a paged response containing more
	// results.
	After param.Opt[string] `query:"after,omitzero" json:"-"`
	// Whether to return only results that have been archived.
	Archived             param.Opt[bool]      `query:"archived,omitzero" json:"-"`
	ContentType          param.Opt[string]    `query:"contentType,omitzero" json:"-"`
	CreatedAfter         param.Opt[time.Time] `query:"createdAfter,omitzero" format:"date-time" json:"-"`
	CreatedAt            param.Opt[time.Time] `query:"createdAt,omitzero" format:"date-time" json:"-"`
	CreatedBefore        param.Opt[time.Time] `query:"createdBefore,omitzero" format:"date-time" json:"-"`
	IsGetLocalizedSchema param.Opt[bool]      `query:"isGetLocalizedSchema,omitzero" json:"-"`
	// The maximum number of results to display per page.
	Limit         param.Opt[int64]     `query:"limit,omitzero" json:"-"`
	UpdatedAfter  param.Opt[time.Time] `query:"updatedAfter,omitzero" format:"date-time" json:"-"`
	UpdatedAt     param.Opt[time.Time] `query:"updatedAt,omitzero" format:"date-time" json:"-"`
	UpdatedBefore param.Opt[time.Time] `query:"updatedBefore,omitzero" format:"date-time" json:"-"`
	Sort          []string             `query:"sort,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [HubdbTableListParams]'s query parameters as `url.Values`.
func (r HubdbTableListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type HubdbTableCloneDraftParams struct {
	HubDBTableCloneRequest HubDBTableCloneRequestParam
	paramObj
}

func (r HubdbTableCloneDraftParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.HubDBTableCloneRequest)
}
func (r *HubdbTableCloneDraftParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type HubdbTableDeleteVersionParams struct {
	TableIDOrName string `path:"tableIdOrName" api:"required" json:"-"`
	paramObj
}

type HubdbTableExportParams struct {
	Format param.Opt[string] `query:"format,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [HubdbTableExportParams]'s query parameters as `url.Values`.
func (r HubdbTableExportParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type HubdbTableExportDraftParams struct {
	Format param.Opt[string] `query:"format,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [HubdbTableExportDraftParams]'s query parameters as
// `url.Values`.
func (r HubdbTableExportDraftParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type HubdbTableGetParams struct {
	// Whether to return only results that have been archived.
	Archived             param.Opt[bool] `query:"archived,omitzero" json:"-"`
	IncludeForeignIDs    param.Opt[bool] `query:"includeForeignIds,omitzero" json:"-"`
	IsGetLocalizedSchema param.Opt[bool] `query:"isGetLocalizedSchema,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [HubdbTableGetParams]'s query parameters as `url.Values`.
func (r HubdbTableGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type HubdbTableGetDraftParams struct {
	// Whether to return only results that have been archived.
	Archived             param.Opt[bool] `query:"archived,omitzero" json:"-"`
	IncludeForeignIDs    param.Opt[bool] `query:"includeForeignIds,omitzero" json:"-"`
	IsGetLocalizedSchema param.Opt[bool] `query:"isGetLocalizedSchema,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [HubdbTableGetDraftParams]'s query parameters as
// `url.Values`.
func (r HubdbTableGetDraftParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type HubdbTableImportDraftParams struct {
	Config param.Opt[string] `json:"config,omitzero"`
	File   io.Reader         `json:"file,omitzero" format:"binary"`
	paramObj
}

func (r HubdbTableImportDraftParams) MarshalMultipart() (data []byte, contentType string, err error) {
	buf := bytes.NewBuffer(nil)
	writer := multipart.NewWriter(buf)
	err = apiform.MarshalRoot(r, writer)
	if err == nil {
		err = apiform.WriteExtras(writer, r.ExtraFields())
	}
	if err != nil {
		writer.Close()
		return nil, "", err
	}
	err = writer.Close()
	if err != nil {
		return nil, "", err
	}
	return buf.Bytes(), writer.FormDataContentType(), nil
}

type HubdbTableListDraftParams struct {
	// The paging cursor token of the last successfully read resource will be returned
	// as the `paging.next.after` JSON property of a paged response containing more
	// results.
	After param.Opt[string] `query:"after,omitzero" json:"-"`
	// Whether to return only results that have been archived.
	Archived             param.Opt[bool]      `query:"archived,omitzero" json:"-"`
	ContentType          param.Opt[string]    `query:"contentType,omitzero" json:"-"`
	CreatedAfter         param.Opt[time.Time] `query:"createdAfter,omitzero" format:"date-time" json:"-"`
	CreatedAt            param.Opt[time.Time] `query:"createdAt,omitzero" format:"date-time" json:"-"`
	CreatedBefore        param.Opt[time.Time] `query:"createdBefore,omitzero" format:"date-time" json:"-"`
	IsGetLocalizedSchema param.Opt[bool]      `query:"isGetLocalizedSchema,omitzero" json:"-"`
	// The maximum number of results to display per page.
	Limit         param.Opt[int64]     `query:"limit,omitzero" json:"-"`
	UpdatedAfter  param.Opt[time.Time] `query:"updatedAfter,omitzero" format:"date-time" json:"-"`
	UpdatedAt     param.Opt[time.Time] `query:"updatedAt,omitzero" format:"date-time" json:"-"`
	UpdatedBefore param.Opt[time.Time] `query:"updatedBefore,omitzero" format:"date-time" json:"-"`
	Sort          []string             `query:"sort,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [HubdbTableListDraftParams]'s query parameters as
// `url.Values`.
func (r HubdbTableListDraftParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type HubdbTablePublishDraftParams struct {
	IncludeForeignIDs param.Opt[bool] `query:"includeForeignIds,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [HubdbTablePublishDraftParams]'s query parameters as
// `url.Values`.
func (r HubdbTablePublishDraftParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type HubdbTableResetDraftParams struct {
	IncludeForeignIDs param.Opt[bool] `query:"includeForeignIds,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [HubdbTableResetDraftParams]'s query parameters as
// `url.Values`.
func (r HubdbTableResetDraftParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type HubdbTableUnpublishParams struct {
	IncludeForeignIDs param.Opt[bool] `query:"includeForeignIds,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [HubdbTableUnpublishParams]'s query parameters as
// `url.Values`.
func (r HubdbTableUnpublishParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type HubdbTableUpdateDraftParams struct {
	HubDBTableV3Request HubDBTableV3RequestParam
	// Whether to return only results that have been archived.
	Archived             param.Opt[bool] `query:"archived,omitzero" json:"-"`
	IncludeForeignIDs    param.Opt[bool] `query:"includeForeignIds,omitzero" json:"-"`
	IsGetLocalizedSchema param.Opt[bool] `query:"isGetLocalizedSchema,omitzero" json:"-"`
	paramObj
}

func (r HubdbTableUpdateDraftParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.HubDBTableV3Request)
}
func (r *HubdbTableUpdateDraftParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [HubdbTableUpdateDraftParams]'s query parameters as
// `url.Values`.
func (r HubdbTableUpdateDraftParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
