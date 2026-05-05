// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package files

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/HubSpot/hubspot-sdk-go/internal/apijson"
	"github.com/HubSpot/hubspot-sdk-go/internal/apiquery"
	shimjson "github.com/HubSpot/hubspot-sdk-go/internal/encoding/json"
	"github.com/HubSpot/hubspot-sdk-go/internal/requestconfig"
	"github.com/HubSpot/hubspot-sdk-go/option"
	"github.com/HubSpot/hubspot-sdk-go/packages/pagination"
	"github.com/HubSpot/hubspot-sdk-go/packages/param"
)

// FolderService contains methods and other services that help with interacting
// with the hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewFolderService] method instead.
type FolderService struct {
	options []option.RequestOption
}

// NewFolderService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewFolderService(opts ...option.RequestOption) (r FolderService) {
	r = FolderService{}
	r.options = opts
	return
}

// Delete folder by ID.
func (r *FolderService) DeleteByID(ctx context.Context, folderID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if folderID == "" {
		err = errors.New("missing required folderId parameter")
		return err
	}
	path := fmt.Sprintf("files/2026-03/folders/%s", url.PathEscape(folderID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Delete a folder, identified by its path.
func (r *FolderService) DeleteByPath(ctx context.Context, folderPath string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if folderPath == "" {
		err = errors.New("missing required folderPath parameter")
		return err
	}
	path := fmt.Sprintf("files/2026-03/folders/%s", url.PathEscape(folderPath))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Retrieve a folder by its ID.
func (r *FolderService) GetByID(ctx context.Context, folderID string, query FolderGetByIDParams, opts ...option.RequestOption) (res *Folder, err error) {
	opts = slices.Concat(r.options, opts)
	if folderID == "" {
		err = errors.New("missing required folderId parameter")
		return nil, err
	}
	path := fmt.Sprintf("files/2026-03/folders/%s", url.PathEscape(folderID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Retrieve a folder, identified by its path.
func (r *FolderService) GetByPath(ctx context.Context, folderPath string, query FolderGetByPathParams, opts ...option.RequestOption) (res *Folder, err error) {
	opts = slices.Concat(r.options, opts)
	if folderPath == "" {
		err = errors.New("missing required folderPath parameter")
		return nil, err
	}
	path := fmt.Sprintf("files/2026-03/folders/%s", url.PathEscape(folderPath))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Check status of folder update. Folder updates happen asynchronously.
func (r *FolderService) GetUpdateAsyncStatus(ctx context.Context, taskID string, opts ...option.RequestOption) (res *FolderActionResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if taskID == "" {
		err = errors.New("missing required taskId parameter")
		return nil, err
	}
	path := fmt.Sprintf("files/2026-03/folders/update/async/tasks/%s/status", url.PathEscape(taskID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Search for folders. Does not contain hidden or archived folders.
func (r *FolderService) Search(ctx context.Context, query FolderSearchParams, opts ...option.RequestOption) (res *pagination.Page[Folder], err error) {
	var raw *http.Response
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "files/2026-03/folders/search"
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

// Search for folders. Does not contain hidden or archived folders.
func (r *FolderService) SearchAutoPaging(ctx context.Context, query FolderSearchParams, opts ...option.RequestOption) *pagination.PageAutoPager[Folder] {
	return pagination.NewPageAutoPager(r.Search(ctx, query, opts...))
}

// Update properties of folder by given ID. This action happens asynchronously and
// will update all of the folder's children as well.
func (r *FolderService) UpdateAsyncByID(ctx context.Context, body FolderUpdateAsyncByIDParams, opts ...option.RequestOption) (res *FolderUpdateTaskLocator, err error) {
	opts = slices.Concat(r.options, opts)
	path := "files/2026-03/folders/update/async"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Update a folder's properties, identified by folder ID.
func (r *FolderService) UpdateByID(ctx context.Context, folderID string, body FolderUpdateByIDParams, opts ...option.RequestOption) (res *Folder, err error) {
	opts = slices.Concat(r.options, opts)
	if folderID == "" {
		err = errors.New("missing required folderId parameter")
		return nil, err
	}
	path := fmt.Sprintf("files/2026-03/folders/%s", url.PathEscape(folderID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

type FolderGetByIDParams struct {
	Properties []string `query:"properties,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [FolderGetByIDParams]'s query parameters as `url.Values`.
func (r FolderGetByIDParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type FolderGetByPathParams struct {
	Properties []string `query:"properties,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [FolderGetByPathParams]'s query parameters as `url.Values`.
func (r FolderGetByPathParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type FolderSearchParams struct {
	// The paging cursor token of the last successfully read resource will be returned
	// as the `paging.next.after` JSON property of a paged response containing more
	// results.
	After        param.Opt[string]    `query:"after,omitzero" json:"-"`
	Before       param.Opt[string]    `query:"before,omitzero" json:"-"`
	CreatedAt    param.Opt[time.Time] `query:"createdAt,omitzero" format:"date-time" json:"-"`
	CreatedAtGte param.Opt[time.Time] `query:"createdAtGte,omitzero" format:"date-time" json:"-"`
	CreatedAtLte param.Opt[time.Time] `query:"createdAtLte,omitzero" format:"date-time" json:"-"`
	IDGte        param.Opt[int64]     `query:"idGte,omitzero" json:"-"`
	IDLte        param.Opt[int64]     `query:"idLte,omitzero" json:"-"`
	// The maximum number of results to display per page.
	Limit           param.Opt[int64]     `query:"limit,omitzero" json:"-"`
	Name            param.Opt[string]    `query:"name,omitzero" json:"-"`
	Path            param.Opt[string]    `query:"path,omitzero" json:"-"`
	UpdatedAt       param.Opt[time.Time] `query:"updatedAt,omitzero" format:"date-time" json:"-"`
	UpdatedAtGte    param.Opt[time.Time] `query:"updatedAtGte,omitzero" format:"date-time" json:"-"`
	UpdatedAtLte    param.Opt[time.Time] `query:"updatedAtLte,omitzero" format:"date-time" json:"-"`
	IDs             []int64              `query:"ids,omitzero" json:"-"`
	ParentFolderIDs []int64              `query:"parentFolderIds,omitzero" json:"-"`
	Properties      []string             `query:"properties,omitzero" json:"-"`
	Sort            []string             `query:"sort,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [FolderSearchParams]'s query parameters as `url.Values`.
func (r FolderSearchParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type FolderUpdateAsyncByIDParams struct {
	FolderUpdateInputWithID FolderUpdateInputWithIDParam
	paramObj
}

func (r FolderUpdateAsyncByIDParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.FolderUpdateInputWithID)
}
func (r *FolderUpdateAsyncByIDParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FolderUpdateByIDParams struct {
	FolderUpdateInput FolderUpdateInputParam
	paramObj
}

func (r FolderUpdateByIDParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.FolderUpdateInput)
}
func (r *FolderUpdateByIDParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
