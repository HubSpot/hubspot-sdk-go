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

	"github.com/stainless-sdks/hubspot-sdk-go/internal/apijson"
	"github.com/stainless-sdks/hubspot-sdk-go/internal/apiquery"
	shimjson "github.com/stainless-sdks/hubspot-sdk-go/internal/encoding/json"
	"github.com/stainless-sdks/hubspot-sdk-go/internal/requestconfig"
	"github.com/stainless-sdks/hubspot-sdk-go/option"
	"github.com/stainless-sdks/hubspot-sdk-go/packages/pagination"
	"github.com/stainless-sdks/hubspot-sdk-go/packages/param"
)

// FileService contains methods and other services that help with interacting with
// the hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewFileService] method instead.
type FileService struct {
	Options []option.RequestOption
}

// NewFileService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewFileService(opts ...option.RequestOption) (r FileService) {
	r = FileService{}
	r.Options = opts
	return
}

// Check the status of requested import.
func (r *FileService) GetImportTaskStatus(ctx context.Context, taskID string, opts ...option.RequestOption) (res *FileActionResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if taskID == "" {
		err = errors.New("missing required taskId parameter")
		return nil, err
	}
	path := fmt.Sprintf("files/2026-03/files/import-from-url/async/tasks/%s/status", taskID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Asynchronously imports the file at the given URL into the file manager.
func (r *FileService) ImportFromURLAsync(ctx context.Context, body FileImportFromURLAsyncParams, opts ...option.RequestOption) (res *ImportFromURLTaskLocator, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "files/2026-03/files/import-from-url/async"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Search through files in the file manager. Does not display hidden or archived
// files.
func (r *FileService) Search(ctx context.Context, query FileSearchParams, opts ...option.RequestOption) (res *pagination.Page[File], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "files/2026-03/files/search"
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

// Search through files in the file manager. Does not display hidden or archived
// files.
func (r *FileService) SearchAutoPaging(ctx context.Context, query FileSearchParams, opts ...option.RequestOption) *pagination.PageAutoPager[File] {
	return pagination.NewPageAutoPager(r.Search(ctx, query, opts...))
}

type FileImportFromURLAsyncParams struct {
	ImportFromURLInput ImportFromURLInputParam
	paramObj
}

func (r FileImportFromURLAsyncParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.ImportFromURLInput)
}
func (r *FileImportFromURLAsyncParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FileSearchParams struct {
	// The paging cursor token of the last successfully read resource will be returned
	// as the `paging.next.after` JSON property of a paged response containing more
	// results.
	After                 param.Opt[string]    `query:"after,omitzero" json:"-"`
	AllowsAnonymousAccess param.Opt[bool]      `query:"allowsAnonymousAccess,omitzero" json:"-"`
	Before                param.Opt[string]    `query:"before,omitzero" json:"-"`
	CreatedAt             param.Opt[time.Time] `query:"createdAt,omitzero" format:"date-time" json:"-"`
	CreatedAtGte          param.Opt[time.Time] `query:"createdAtGte,omitzero" format:"date-time" json:"-"`
	CreatedAtLte          param.Opt[time.Time] `query:"createdAtLte,omitzero" format:"date-time" json:"-"`
	Encoding              param.Opt[string]    `query:"encoding,omitzero" json:"-"`
	ExpiresAt             param.Opt[time.Time] `query:"expiresAt,omitzero" format:"date-time" json:"-"`
	ExpiresAtGte          param.Opt[time.Time] `query:"expiresAtGte,omitzero" format:"date-time" json:"-"`
	ExpiresAtLte          param.Opt[time.Time] `query:"expiresAtLte,omitzero" format:"date-time" json:"-"`
	Extension             param.Opt[string]    `query:"extension,omitzero" json:"-"`
	FileMd5               param.Opt[string]    `query:"fileMd5,omitzero" json:"-"`
	Height                param.Opt[int64]     `query:"height,omitzero" json:"-"`
	HeightGte             param.Opt[int64]     `query:"heightGte,omitzero" json:"-"`
	HeightLte             param.Opt[int64]     `query:"heightLte,omitzero" json:"-"`
	IDGte                 param.Opt[int64]     `query:"idGte,omitzero" json:"-"`
	IDLte                 param.Opt[int64]     `query:"idLte,omitzero" json:"-"`
	IsUsableInContent     param.Opt[bool]      `query:"isUsableInContent,omitzero" json:"-"`
	// The maximum number of results to display per page.
	Limit           param.Opt[int64]     `query:"limit,omitzero" json:"-"`
	Name            param.Opt[string]    `query:"name,omitzero" json:"-"`
	Path            param.Opt[string]    `query:"path,omitzero" json:"-"`
	Size            param.Opt[int64]     `query:"size,omitzero" json:"-"`
	SizeGte         param.Opt[int64]     `query:"sizeGte,omitzero" json:"-"`
	SizeLte         param.Opt[int64]     `query:"sizeLte,omitzero" json:"-"`
	Type            param.Opt[string]    `query:"type,omitzero" json:"-"`
	UpdatedAt       param.Opt[time.Time] `query:"updatedAt,omitzero" format:"date-time" json:"-"`
	UpdatedAtGte    param.Opt[time.Time] `query:"updatedAtGte,omitzero" format:"date-time" json:"-"`
	UpdatedAtLte    param.Opt[time.Time] `query:"updatedAtLte,omitzero" format:"date-time" json:"-"`
	URL             param.Opt[string]    `query:"url,omitzero" json:"-"`
	Width           param.Opt[int64]     `query:"width,omitzero" json:"-"`
	WidthGte        param.Opt[int64]     `query:"widthGte,omitzero" json:"-"`
	WidthLte        param.Opt[int64]     `query:"widthLte,omitzero" json:"-"`
	IDs             []int64              `query:"ids,omitzero" json:"-"`
	ParentFolderIDs []int64              `query:"parentFolderIds,omitzero" json:"-"`
	Properties      []string             `query:"properties,omitzero" json:"-"`
	Sort            []string             `query:"sort,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [FileSearchParams]'s query parameters as `url.Values`.
func (r FileSearchParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
