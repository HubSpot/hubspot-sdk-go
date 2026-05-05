// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package files

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

// FileAssetService contains methods and other services that help with interacting
// with the hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewFileAssetService] method instead.
type FileAssetService struct {
	options []option.RequestOption
}

// NewFileAssetService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewFileAssetService(opts ...option.RequestOption) (r FileAssetService) {
	r = FileAssetService{}
	r.options = opts
	return
}

// Creates a folder.
func (r *FileAssetService) New(ctx context.Context, body FileAssetNewParams, opts ...option.RequestOption) (res *Folder, err error) {
	opts = slices.Concat(r.options, opts)
	path := "files/2026-03/folders"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Update properties of file by ID.
func (r *FileAssetService) Update(ctx context.Context, fileID string, body FileAssetUpdateParams, opts ...option.RequestOption) (res *File, err error) {
	opts = slices.Concat(r.options, opts)
	if fileID == "" {
		err = errors.New("missing required fileId parameter")
		return nil, err
	}
	path := fmt.Sprintf("files/2026-03/files/%s", url.PathEscape(fileID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

// Delete a file by ID
func (r *FileAssetService) Delete(ctx context.Context, fileID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if fileID == "" {
		err = errors.New("missing required fileId parameter")
		return err
	}
	path := fmt.Sprintf("files/2026-03/files/%s", url.PathEscape(fileID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Delete a file in accordance with GDPR regulations.
func (r *FileAssetService) GdprDelete(ctx context.Context, fileID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if fileID == "" {
		err = errors.New("missing required fileId parameter")
		return err
	}
	path := fmt.Sprintf("files/2026-03/files/%s/gdpr-delete", url.PathEscape(fileID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Retrieve a file by its ID.
func (r *FileAssetService) Get(ctx context.Context, fileID string, query FileAssetGetParams, opts ...option.RequestOption) (res *File, err error) {
	opts = slices.Concat(r.options, opts)
	if fileID == "" {
		err = errors.New("missing required fileId parameter")
		return nil, err
	}
	path := fmt.Sprintf("files/2026-03/files/%s", url.PathEscape(fileID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Check the status of requested import.
func (r *FileAssetService) GetImportTaskStatus(ctx context.Context, taskID string, opts ...option.RequestOption) (res *FileActionResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if taskID == "" {
		err = errors.New("missing required taskId parameter")
		return nil, err
	}
	path := fmt.Sprintf("files/2026-03/files/import-from-url/async/tasks/%s/status", url.PathEscape(taskID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Generates signed URL that allows temporary access to a private file.
func (r *FileAssetService) GetSignedURL(ctx context.Context, fileID string, query FileAssetGetSignedURLParams, opts ...option.RequestOption) (res *SignedURL, err error) {
	opts = slices.Concat(r.options, opts)
	if fileID == "" {
		err = errors.New("missing required fileId parameter")
		return nil, err
	}
	path := fmt.Sprintf("files/2026-03/files/%s/signed-url", url.PathEscape(fileID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Asynchronously imports the file at the given URL into the file manager.
func (r *FileAssetService) ImportFromURLAsync(ctx context.Context, body FileAssetImportFromURLAsyncParams, opts ...option.RequestOption) (res *ImportFromURLTaskLocator, err error) {
	opts = slices.Concat(r.options, opts)
	path := "files/2026-03/files/import-from-url/async"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Replace existing file data with new file data. Can be used to change image
// content without having to upload a new file and update all references.
func (r *FileAssetService) Replace(ctx context.Context, fileID string, body FileAssetReplaceParams, opts ...option.RequestOption) (res *File, err error) {
	opts = slices.Concat(r.options, opts)
	if fileID == "" {
		err = errors.New("missing required fileId parameter")
		return nil, err
	}
	path := fmt.Sprintf("files/2026-03/files/%s", url.PathEscape(fileID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return res, err
}

// Search through files in the file manager. Does not display hidden or archived
// files.
func (r *FileAssetService) Search(ctx context.Context, query FileAssetSearchParams, opts ...option.RequestOption) (res *pagination.Page[File], err error) {
	var raw *http.Response
	opts = slices.Concat(r.options, opts)
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
func (r *FileAssetService) SearchAutoPaging(ctx context.Context, query FileAssetSearchParams, opts ...option.RequestOption) *pagination.PageAutoPager[File] {
	return pagination.NewPageAutoPager(r.Search(ctx, query, opts...))
}

// Upload a single file with content specified in request body.
func (r *FileAssetService) Upload(ctx context.Context, body FileAssetUploadParams, opts ...option.RequestOption) (res *File, err error) {
	opts = slices.Concat(r.options, opts)
	path := "files/2026-03/files"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type FileAssetNewParams struct {
	FolderInput FolderInputParam
	paramObj
}

func (r FileAssetNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.FolderInput)
}
func (r *FileAssetNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FileAssetUpdateParams struct {
	FileUpdateInput FileUpdateInputParam
	paramObj
}

func (r FileAssetUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.FileUpdateInput)
}
func (r *FileAssetUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FileAssetGetParams struct {
	Properties []string `query:"properties,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [FileAssetGetParams]'s query parameters as `url.Values`.
func (r FileAssetGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type FileAssetGetSignedURLParams struct {
	ExpirationSeconds param.Opt[int64] `query:"expirationSeconds,omitzero" json:"-"`
	Upscale           param.Opt[bool]  `query:"upscale,omitzero" json:"-"`
	// Any of "icon", "medium", "preview", "thumb".
	Size FileAssetGetSignedURLParamsSize `query:"size,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [FileAssetGetSignedURLParams]'s query parameters as
// `url.Values`.
func (r FileAssetGetSignedURLParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type FileAssetGetSignedURLParamsSize string

const (
	FileAssetGetSignedURLParamsSizeIcon    FileAssetGetSignedURLParamsSize = "icon"
	FileAssetGetSignedURLParamsSizeMedium  FileAssetGetSignedURLParamsSize = "medium"
	FileAssetGetSignedURLParamsSizePreview FileAssetGetSignedURLParamsSize = "preview"
	FileAssetGetSignedURLParamsSizeThumb   FileAssetGetSignedURLParamsSize = "thumb"
)

type FileAssetImportFromURLAsyncParams struct {
	ImportFromURLInput ImportFromURLInputParam
	paramObj
}

func (r FileAssetImportFromURLAsyncParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.ImportFromURLInput)
}
func (r *FileAssetImportFromURLAsyncParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FileAssetReplaceParams struct {
	CharsetHunch param.Opt[string] `json:"charsetHunch,omitzero"`
	Options      param.Opt[string] `json:"options,omitzero"`
	File         io.Reader         `json:"file,omitzero" format:"binary"`
	paramObj
}

func (r FileAssetReplaceParams) MarshalMultipart() (data []byte, contentType string, err error) {
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

type FileAssetSearchParams struct {
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

// URLQuery serializes [FileAssetSearchParams]'s query parameters as `url.Values`.
func (r FileAssetSearchParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type FileAssetUploadParams struct {
	CharsetHunch param.Opt[string] `json:"charsetHunch,omitzero"`
	FileName     param.Opt[string] `json:"fileName,omitzero"`
	FolderID     param.Opt[string] `json:"folderId,omitzero"`
	FolderPath   param.Opt[string] `json:"folderPath,omitzero"`
	Options      param.Opt[string] `json:"options,omitzero"`
	File         io.Reader         `json:"file,omitzero" format:"binary"`
	paramObj
}

func (r FileAssetUploadParams) MarshalMultipart() (data []byte, contentType string, err error) {
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
