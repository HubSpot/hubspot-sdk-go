// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package files

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/stainless-sdks/hubspot-sdk-go/internal/apiform"
	"github.com/stainless-sdks/hubspot-sdk-go/internal/apiquery"
	shimjson "github.com/stainless-sdks/hubspot-sdk-go/internal/encoding/json"
	"github.com/stainless-sdks/hubspot-sdk-go/internal/requestconfig"
	"github.com/stainless-sdks/hubspot-sdk-go/option"
	"github.com/stainless-sdks/hubspot-sdk-go/packages/pagination"
	"github.com/stainless-sdks/hubspot-sdk-go/packages/param"
)

// FileOperationService contains methods and other services that help with
// interacting with the hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewFileOperationService] method instead.
type FileOperationService struct {
	Options []option.RequestOption
}

// NewFileOperationService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewFileOperationService(opts ...option.RequestOption) (r FileOperationService) {
	r = FileOperationService{}
	r.Options = opts
	return
}

// Update properties of file by ID.
func (r *FileOperationService) Update(ctx context.Context, fileID string, body FileOperationUpdateParams, opts ...option.RequestOption) (res *File, err error) {
	opts = slices.Concat(r.Options, opts)
	if fileID == "" {
		err = errors.New("missing required fileId parameter")
		return
	}
	path := fmt.Sprintf("files/v3/files/%s", fileID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Delete a file by ID
func (r *FileOperationService) Delete(ctx context.Context, fileID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if fileID == "" {
		err = errors.New("missing required fileId parameter")
		return
	}
	path := fmt.Sprintf("files/v3/files/%s", fileID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Delete a file in accordance with GDPR regulations.
func (r *FileOperationService) GdprDelete(ctx context.Context, fileID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if fileID == "" {
		err = errors.New("missing required fileId parameter")
		return
	}
	path := fmt.Sprintf("files/v3/files/%s/gdpr-delete", fileID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Retrieve a file by its ID.
func (r *FileOperationService) Get(ctx context.Context, fileID string, query FileOperationGetParams, opts ...option.RequestOption) (res *File, err error) {
	opts = slices.Concat(r.Options, opts)
	if fileID == "" {
		err = errors.New("missing required fileId parameter")
		return
	}
	path := fmt.Sprintf("files/v3/files/%s", fileID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Retrieve a file by its path.
func (r *FileOperationService) GetByPath(ctx context.Context, filePath string, query FileOperationGetByPathParams, opts ...option.RequestOption) (res *FileStat, err error) {
	opts = slices.Concat(r.Options, opts)
	if filePath == "" {
		err = errors.New("missing required file_path parameter")
		return
	}
	path := fmt.Sprintf("files/v3/files/stat/%s", filePath)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Check the status of requested import.
func (r *FileOperationService) GetImportTaskStatus(ctx context.Context, taskID string, opts ...option.RequestOption) (res *FileActionResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if taskID == "" {
		err = errors.New("missing required taskId parameter")
		return
	}
	path := fmt.Sprintf("files/v3/files/import-from-url/async/tasks/%s/status", taskID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Generates signed URL that allows temporary access to a private file.
func (r *FileOperationService) GetSignedURL(ctx context.Context, fileID string, query FileOperationGetSignedURLParams, opts ...option.RequestOption) (res *SignedURL, err error) {
	opts = slices.Concat(r.Options, opts)
	if fileID == "" {
		err = errors.New("missing required fileId parameter")
		return
	}
	path := fmt.Sprintf("files/v3/files/%s/signed-url", fileID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Asynchronously imports the file at the given URL into the file manager.
func (r *FileOperationService) ImportFromURLAsync(ctx context.Context, body FileOperationImportFromURLAsyncParams, opts ...option.RequestOption) (res *ImportFromURLTaskLocator, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "files/v3/files/import-from-url/async"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Replace existing file data with new file data. Can be used to change image
// content without having to upload a new file and update all references.
func (r *FileOperationService) Replace(ctx context.Context, fileID string, body FileOperationReplaceParams, opts ...option.RequestOption) (res *File, err error) {
	opts = slices.Concat(r.Options, opts)
	if fileID == "" {
		err = errors.New("missing required fileId parameter")
		return
	}
	path := fmt.Sprintf("files/v3/files/%s", fileID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Search through files in the file manager. Does not display hidden or archived
// files.
func (r *FileOperationService) Search(ctx context.Context, query FileOperationSearchParams, opts ...option.RequestOption) (res *pagination.Page[File], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "files/v3/files/search"
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
func (r *FileOperationService) SearchAutoPaging(ctx context.Context, query FileOperationSearchParams, opts ...option.RequestOption) *pagination.PageAutoPager[File] {
	return pagination.NewPageAutoPager(r.Search(ctx, query, opts...))
}

// Upload a single file with content specified in request body.
func (r *FileOperationService) Upload(ctx context.Context, body FileOperationUploadParams, opts ...option.RequestOption) (res *File, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "files/v3/files"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type FileOperationUpdateParams struct {
	// Object for updating files.
	FileUpdateInput FileUpdateInputParam
	paramObj
}

func (r FileOperationUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.FileUpdateInput)
}
func (r *FileOperationUpdateParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.FileUpdateInput)
}

type FileOperationGetParams struct {
	Properties []string `query:"properties,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [FileOperationGetParams]'s query parameters as `url.Values`.
func (r FileOperationGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type FileOperationGetByPathParams struct {
	// Properties to return in the response.
	Properties []string `query:"properties,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [FileOperationGetByPathParams]'s query parameters as
// `url.Values`.
func (r FileOperationGetByPathParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type FileOperationGetSignedURLParams struct {
	// How long in seconds the link will provide access to the file.
	ExpirationSeconds param.Opt[int64] `query:"expirationSeconds,omitzero" json:"-"`
	// If size is provided, this will upscale the image to fit the size dimensions.
	Upscale param.Opt[bool] `query:"upscale,omitzero" json:"-"`
	// For image files. This will resize the image to the desired size before sharing.
	// Does not affect the original file, just the file served by this signed URL.
	//
	// Any of "icon", "medium", "preview", "thumb".
	Size FileOperationGetSignedURLParamsSize `query:"size,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [FileOperationGetSignedURLParams]'s query parameters as
// `url.Values`.
func (r FileOperationGetSignedURLParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// For image files. This will resize the image to the desired size before sharing.
// Does not affect the original file, just the file served by this signed URL.
type FileOperationGetSignedURLParamsSize string

const (
	FileOperationGetSignedURLParamsSizeIcon    FileOperationGetSignedURLParamsSize = "icon"
	FileOperationGetSignedURLParamsSizeMedium  FileOperationGetSignedURLParamsSize = "medium"
	FileOperationGetSignedURLParamsSizePreview FileOperationGetSignedURLParamsSize = "preview"
	FileOperationGetSignedURLParamsSizeThumb   FileOperationGetSignedURLParamsSize = "thumb"
)

type FileOperationImportFromURLAsyncParams struct {
	ImportFromURLInput ImportFromURLInputParam
	paramObj
}

func (r FileOperationImportFromURLAsyncParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.ImportFromURLInput)
}
func (r *FileOperationImportFromURLAsyncParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.ImportFromURLInput)
}

type FileOperationReplaceParams struct {
	// Character set of given file data.
	CharsetHunch param.Opt[string] `json:"charsetHunch,omitzero"`
	// JSON string representing FileReplaceOptions. Includes options to set the access
	// and expiresAt properties, which will automatically update when the file is
	// replaced.
	Options param.Opt[string] `json:"options,omitzero"`
	// File data that will replace existing file in the file manager.
	File io.Reader `json:"file,omitzero" format:"binary"`
	paramObj
}

func (r FileOperationReplaceParams) MarshalMultipart() (data []byte, contentType string, err error) {
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

type FileOperationSearchParams struct {
	// Offset search results by this value. The default offset is 0 and the maximum
	// offset of items for a given search is 10,000. Narrow your search down if you are
	// reaching this limit.
	After param.Opt[string] `query:"after,omitzero" json:"-"`
	// Search files by access. If `true`, will show only public files. If `false`, will
	// show only private files.
	AllowsAnonymousAccess param.Opt[bool]   `query:"allowsAnonymousAccess,omitzero" json:"-"`
	Before                param.Opt[string] `query:"before,omitzero" json:"-"`
	// Search files by time of creation.
	CreatedAt param.Opt[time.Time] `query:"createdAt,omitzero" format:"date-time" json:"-"`
	// Search files by greater than or equal to time of creation. Can be used with
	// `createdAtLte` to create a range.
	CreatedAtGte param.Opt[time.Time] `query:"createdAtGte,omitzero" format:"date-time" json:"-"`
	// Search files by less than or equal to time of creation. Can be used with
	// `createdAtGte` to create a range.
	CreatedAtLte param.Opt[time.Time] `query:"createdAtLte,omitzero" format:"date-time" json:"-"`
	// Search files by specified encoding.
	Encoding param.Opt[string] `query:"encoding,omitzero" json:"-"`
	// Search files by exact expires time. Time must be epoch time in milliseconds.
	ExpiresAt param.Opt[time.Time] `query:"expiresAt,omitzero" format:"date-time" json:"-"`
	// Search files by greater than or equal to expires time. Can be used with
	// `expiresAtLte` to create a range.
	ExpiresAtGte param.Opt[time.Time] `query:"expiresAtGte,omitzero" format:"date-time" json:"-"`
	// Search files by less than or equal to expires time. Can be used with
	// `expiresAtGte` to create a range.
	ExpiresAtLte param.Opt[time.Time] `query:"expiresAtLte,omitzero" format:"date-time" json:"-"`
	// Search files by given extension.
	Extension param.Opt[string] `query:"extension,omitzero" json:"-"`
	// Search files by a specific md5 hash.
	FileMd5 param.Opt[string] `query:"fileMd5,omitzero" json:"-"`
	// Search files by height of image or video.
	Height param.Opt[int64] `query:"height,omitzero" json:"-"`
	// Search files by greater than or equal to height of image or video. Can be used
	// with `heightLte` to create a range.
	HeightGte param.Opt[int64] `query:"heightGte,omitzero" json:"-"`
	// Search files by less than or equal to height of image or video. Can be used with
	// `heightGte` to create a range.
	HeightLte param.Opt[int64] `query:"heightLte,omitzero" json:"-"`
	IDGte     param.Opt[int64] `query:"idGte,omitzero" json:"-"`
	IDLte     param.Opt[int64] `query:"idLte,omitzero" json:"-"`
	// If `true`, shows files that have been marked to be used in new content. If
	// `false`, shows files that should not be used in new content.
	IsUsableInContent param.Opt[bool] `query:"isUsableInContent,omitzero" json:"-"`
	// Number of items to return. Default limit is 10, maximum limit is 100.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Search for files containing the given name.
	Name param.Opt[string] `query:"name,omitzero" json:"-"`
	// Search files by path.
	Path param.Opt[string] `query:"path,omitzero" json:"-"`
	// Search files by exact file size in bytes.
	Size param.Opt[int64] `query:"size,omitzero" json:"-"`
	// Search files by greater than or equal to file size. Can be used with `sizeLte`
	// to create a range.
	SizeGte param.Opt[int64] `query:"sizeGte,omitzero" json:"-"`
	// Search files by less than or equal to file size. Can be used with `sizeGte` to
	// create a range.
	SizeLte param.Opt[int64] `query:"sizeLte,omitzero" json:"-"`
	// Filter by provided file type.
	Type param.Opt[string] `query:"type,omitzero" json:"-"`
	// Search files by time of latest updated.
	UpdatedAt param.Opt[time.Time] `query:"updatedAt,omitzero" format:"date-time" json:"-"`
	// Search files by greater than or equal to time of latest update. Can be used with
	// `updatedAtLte` to create a range.
	UpdatedAtGte param.Opt[time.Time] `query:"updatedAtGte,omitzero" format:"date-time" json:"-"`
	// Search files by less than or equal to time of latest update. Can be used with
	// `updatedAtGte` to create a range.
	UpdatedAtLte param.Opt[time.Time] `query:"updatedAtLte,omitzero" format:"date-time" json:"-"`
	// Search by file URL.
	URL param.Opt[string] `query:"url,omitzero" json:"-"`
	// Search files by width of image or video.
	Width param.Opt[int64] `query:"width,omitzero" json:"-"`
	// Search files by greater than or equal to width of image or video. Can be used
	// with `widthLte` to create a range.
	WidthGte param.Opt[int64] `query:"widthGte,omitzero" json:"-"`
	// Search files by less than or equal to width of image or video. Can be used with
	// `widthGte` to create a range.
	WidthLte param.Opt[int64] `query:"widthLte,omitzero" json:"-"`
	// Search by a list of file IDs.
	IDs []int64 `query:"ids,omitzero" json:"-"`
	// Search files within given `folderId`.
	ParentFolderIDs []int64 `query:"parentFolderIds,omitzero" json:"-"`
	// A list of file properties to return.
	Properties []string `query:"properties,omitzero" json:"-"`
	// Sort files by a given field.
	Sort []string `query:"sort,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [FileOperationSearchParams]'s query parameters as
// `url.Values`.
func (r FileOperationSearchParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type FileOperationUploadParams struct {
	// Character set of the uploaded file.
	CharsetHunch param.Opt[string] `json:"charsetHunch,omitzero"`
	// Desired name for the uploaded file.
	FileName param.Opt[string] `json:"fileName,omitzero"`
	// Either 'folderId' or 'folderPath' is required. folderId is the ID of the folder
	// the file will be uploaded to.
	FolderID param.Opt[string] `json:"folderId,omitzero"`
	// Either 'folderPath' or 'folderId' is required. This field represents the
	// destination folder path for the uploaded file. If a path doesn't exist, the
	// system will try to create one.
	FolderPath param.Opt[string] `json:"folderPath,omitzero"`
	// JSON string representing FileUploadOptions.
	Options param.Opt[string] `json:"options,omitzero"`
	// File to be uploaded.
	File io.Reader `json:"file,omitzero" format:"binary"`
	paramObj
}

func (r FileOperationUploadParams) MarshalMultipart() (data []byte, contentType string, err error) {
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
