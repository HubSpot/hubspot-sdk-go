// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package files

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/stainless-sdks/hubspot-sdk-go/internal/apiquery"
	shimjson "github.com/stainless-sdks/hubspot-sdk-go/internal/encoding/json"
	"github.com/stainless-sdks/hubspot-sdk-go/internal/requestconfig"
	"github.com/stainless-sdks/hubspot-sdk-go/option"
	"github.com/stainless-sdks/hubspot-sdk-go/packages/pagination"
	"github.com/stainless-sdks/hubspot-sdk-go/packages/param"
)

// FolderService contains methods and other services that help with interacting
// with the Hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewFolderService] method instead.
type FolderService struct {
	Options []option.RequestOption
}

// NewFolderService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewFolderService(opts ...option.RequestOption) (r FolderService) {
	r = FolderService{}
	r.Options = opts
	return
}

// Creates a folder.
func (r *FolderService) New(ctx context.Context, body FolderNewParams, opts ...option.RequestOption) (res *Folder, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "files/v3/folders"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Delete folder by ID.
func (r *FolderService) DeleteByID(ctx context.Context, folderID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if folderID == "" {
		err = errors.New("missing required folderId parameter")
		return
	}
	path := fmt.Sprintf("files/v3/folders/%s", folderID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Delete a folder, identified by its path.
func (r *FolderService) DeleteByPath(ctx context.Context, folderPath string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if folderPath == "" {
		err = errors.New("missing required folderPath parameter")
		return
	}
	path := fmt.Sprintf("files/v3/folders/%s", folderPath)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Retrieve a folder by its ID.
func (r *FolderService) GetByID(ctx context.Context, folderID string, query FolderGetByIDParams, opts ...option.RequestOption) (res *Folder, err error) {
	opts = slices.Concat(r.Options, opts)
	if folderID == "" {
		err = errors.New("missing required folderId parameter")
		return
	}
	path := fmt.Sprintf("files/v3/folders/%s", folderID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Retrieve a folder, identified by its path.
func (r *FolderService) GetByPath(ctx context.Context, folderPath string, query FolderGetByPathParams, opts ...option.RequestOption) (res *Folder, err error) {
	opts = slices.Concat(r.Options, opts)
	if folderPath == "" {
		err = errors.New("missing required folderPath parameter")
		return
	}
	path := fmt.Sprintf("files/v3/folders/%s", folderPath)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Check status of folder update. Folder updates happen asynchronously.
func (r *FolderService) GetUpdateAsyncStatus(ctx context.Context, taskID string, opts ...option.RequestOption) (res *FolderActionResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if taskID == "" {
		err = errors.New("missing required taskId parameter")
		return
	}
	path := fmt.Sprintf("files/v3/folders/update/async/tasks/%s/status", taskID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Search for folders. Does not contain hidden or archived folders.
func (r *FolderService) Search(ctx context.Context, query FolderSearchParams, opts ...option.RequestOption) (res *pagination.Page[Folder], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "files/v3/folders/search"
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
	opts = slices.Concat(r.Options, opts)
	path := "files/v3/folders/update/async"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Update a folder's properties, identified by folder ID.
func (r *FolderService) UpdateByID(ctx context.Context, folderID string, body FolderUpdateByIDParams, opts ...option.RequestOption) (res *Folder, err error) {
	opts = slices.Concat(r.Options, opts)
	if folderID == "" {
		err = errors.New("missing required folderId parameter")
		return
	}
	path := fmt.Sprintf("files/v3/folders/%s", folderID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

type FolderNewParams struct {
	// Object for creating a folder.
	FolderInput FolderInputParam
	paramObj
}

func (r FolderNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.FolderInput)
}
func (r *FolderNewParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.FolderInput)
}

type FolderGetByIDParams struct {
	// Properties to set on returned folder.
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
	// Properties to set on returned folder.
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
	// Offset search results by this value. The default offset is 0 and the maximum
	// offset of items for a given search is 10,000. Narrow your search down if you are
	// reaching this limit.
	After  param.Opt[string] `query:"after,omitzero" json:"-"`
	Before param.Opt[string] `query:"before,omitzero" json:"-"`
	// Search folders by exact time of creation. Time must be epoch time in
	// milliseconds.
	CreatedAt param.Opt[time.Time] `query:"createdAt,omitzero" format:"date-time" json:"-"`
	// Search folders by greater than or equal to time of creation. Can be used with
	// createdAtLte to create a range.
	CreatedAtGte param.Opt[time.Time] `query:"createdAtGte,omitzero" format:"date-time" json:"-"`
	// Search folders by less than or equal to time of creation. Can be used with
	// createdAtGte to create a range.
	CreatedAtLte param.Opt[time.Time] `query:"createdAtLte,omitzero" format:"date-time" json:"-"`
	IDGte        param.Opt[int64]     `query:"idGte,omitzero" json:"-"`
	IDLte        param.Opt[int64]     `query:"idLte,omitzero" json:"-"`
	// Number of items to return. Default limit is 10, maximum limit is 100.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Search for folders containing the specified name.
	Name param.Opt[string] `query:"name,omitzero" json:"-"`
	// Search folders by path.
	Path param.Opt[string] `query:"path,omitzero" json:"-"`
	// Search folders by exact time of latest updated. Time must be epoch time in
	// milliseconds.
	UpdatedAt param.Opt[time.Time] `query:"updatedAt,omitzero" format:"date-time" json:"-"`
	// Search folders by greater than or equal to time of latest update. Can be used
	// with updatedAtLte to create a range.
	UpdatedAtGte param.Opt[time.Time] `query:"updatedAtGte,omitzero" format:"date-time" json:"-"`
	// Search folders by less than or equal to time of latest update. Can be used with
	// updatedAtGte to create a range.
	UpdatedAtLte param.Opt[time.Time] `query:"updatedAtLte,omitzero" format:"date-time" json:"-"`
	IDs          []int64              `query:"ids,omitzero" json:"-"`
	// Search folders with the given parent folderId.
	ParentFolderIDs []int64 `query:"parentFolderIds,omitzero" json:"-"`
	// Properties that should be included in the returned folders.
	Properties []string `query:"properties,omitzero" json:"-"`
	// Sort results by given property. For example -name sorts by name field
	// descending, name sorts by name field ascending.
	Sort []string `query:"sort,omitzero" json:"-"`
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
	return json.Unmarshal(data, &r.FolderUpdateInputWithID)
}

type FolderUpdateByIDParams struct {
	// Object for updating folders.
	FolderUpdateInput FolderUpdateInputParam
	paramObj
}

func (r FolderUpdateByIDParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.FolderUpdateInput)
}
func (r *FolderUpdateByIDParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.FolderUpdateInput)
}
