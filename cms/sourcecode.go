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

	"github.com/stainless-sdks/hubspot-sdk-go/internal/apiform"
	"github.com/stainless-sdks/hubspot-sdk-go/internal/apijson"
	"github.com/stainless-sdks/hubspot-sdk-go/internal/apiquery"
	"github.com/stainless-sdks/hubspot-sdk-go/internal/requestconfig"
	"github.com/stainless-sdks/hubspot-sdk-go/option"
	"github.com/stainless-sdks/hubspot-sdk-go/packages/param"
	"github.com/stainless-sdks/hubspot-sdk-go/packages/respjson"
	"github.com/stainless-sdks/hubspot-sdk-go/shared"
)

// SourceCodeService contains methods and other services that help with interacting
// with the Hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSourceCodeService] method instead.
type SourceCodeService struct {
	Options []option.RequestOption
}

// NewSourceCodeService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewSourceCodeService(opts ...option.RequestOption) (r SourceCodeService) {
	r = SourceCodeService{}
	r.Options = opts
	return
}

// Creates a file at the specified path in the specified environment. Accepts
// multipart/form-data content type. Throws an error if a file already exists at
// the specified path.
//
// Deprecated: deprecated
func (r *SourceCodeService) New(ctx context.Context, path string, params SourceCodeNewParams, opts ...option.RequestOption) (res *SourceCodeNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.Environment == "" {
		err = errors.New("missing required environment parameter")
		return
	}
	if path == "" {
		err = errors.New("missing required path parameter")
		return
	}
	path := fmt.Sprintf("cms/v3/source-code/%s/content/%s", params.Environment, path)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Deletes the file at the specified path in the specified environment.
func (r *SourceCodeService) Delete(ctx context.Context, path string, body SourceCodeDeleteParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if body.Environment == "" {
		err = errors.New("missing required environment parameter")
		return
	}
	if path == "" {
		err = errors.New("missing required path parameter")
		return
	}
	path := fmt.Sprintf("cms/v3/source-code/%s/content/%s", body.Environment, path)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Extract a zip file in the developer file system. Extraction status can be
// checked with the `/extract/async/tasks/taskId/status` endpoint below.
func (r *SourceCodeService) ExtractAsync(ctx context.Context, body SourceCodeExtractAsyncParams, opts ...option.RequestOption) (res *shared.TaskLocator, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "cms/v3/source-code/extract/async"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Downloads the byte contents of the file at the specified path in the specified
// environment.
func (r *SourceCodeService) Get(ctx context.Context, path string, query SourceCodeGetParams, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/octet-stream")}, opts...)
	if query.Environment == "" {
		err = errors.New("missing required environment parameter")
		return
	}
	if path == "" {
		err = errors.New("missing required path parameter")
		return
	}
	path := fmt.Sprintf("cms/v3/source-code/%s/content/%s", query.Environment, path)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Get the status of an extraction by the `taskId` returned from the initial
// `extract/async` request.
func (r *SourceCodeService) GetExtractionStatus(ctx context.Context, taskID int64, opts ...option.RequestOption) (res *shared.ActionResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("cms/v3/source-code/extract/async/tasks/%v/status", taskID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Gets the metadata object for the file at the specified path in the specified
// environment.
func (r *SourceCodeService) GetMetadata(ctx context.Context, path string, params SourceCodeGetMetadataParams, opts ...option.RequestOption) (res *SourceCodeGetMetadataResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.Environment == "" {
		err = errors.New("missing required environment parameter")
		return
	}
	if path == "" {
		err = errors.New("missing required path parameter")
		return
	}
	path := fmt.Sprintf("cms/v3/source-code/%s/metadata/%s", params.Environment, path)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return
}

// Upserts a file at the specified path in the specified environment. Accepts
// multipart/form-data content type.
func (r *SourceCodeService) Upsert(ctx context.Context, path string, params SourceCodeUpsertParams, opts ...option.RequestOption) (res *SourceCodeUpsertResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.Environment == "" {
		err = errors.New("missing required environment parameter")
		return
	}
	if path == "" {
		err = errors.New("missing required path parameter")
		return
	}
	path := fmt.Sprintf("cms/v3/source-code/%s/content/%s", params.Environment, path)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return
}

// Validates the file contents passed to the endpoint given a specified path and
// environment. Accepts multipart/form-data content type.
func (r *SourceCodeService) Validate(ctx context.Context, path string, params SourceCodeValidateParams, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if params.Environment == "" {
		err = errors.New("missing required environment parameter")
		return
	}
	if path == "" {
		err = errors.New("missing required path parameter")
		return
	}
	path := fmt.Sprintf("cms/v3/source-code/%s/validate/%s", params.Environment, path)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

type SourceCodeNewResponse struct {
	// The path of the file in the CMS Developer File System.
	ID string `json:"id,required"`
	// Timestamp of when the object was first created.
	CreatedAt int64 `json:"createdAt,required"`
	// Determines whether or not this path points to a folder.
	Folder bool `json:"folder,required"`
	// The name of the file.
	Name string `json:"name,required"`
	// Timestamp of when the object was last updated.
	UpdatedAt int64 `json:"updatedAt,required"`
	// Timestamp of when the object was archived (deleted).
	ArchivedAt int64 `json:"archivedAt"`
	// If the object is a folder, contains the filenames of the files within the
	// folder.
	Children []string `json:"children"`
	Hash     string   `json:"hash"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		Folder      respjson.Field
		Name        respjson.Field
		UpdatedAt   respjson.Field
		ArchivedAt  respjson.Field
		Children    respjson.Field
		Hash        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SourceCodeNewResponse) RawJSON() string { return r.JSON.raw }
func (r *SourceCodeNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SourceCodeGetMetadataResponse struct {
	// The path of the file in the CMS Developer File System.
	ID string `json:"id,required"`
	// Timestamp of when the object was first created.
	CreatedAt int64 `json:"createdAt,required"`
	// Determines whether or not this path points to a folder.
	Folder bool `json:"folder,required"`
	// The name of the file.
	Name string `json:"name,required"`
	// Timestamp of when the object was last updated.
	UpdatedAt int64 `json:"updatedAt,required"`
	// Timestamp of when the object was archived (deleted).
	ArchivedAt int64 `json:"archivedAt"`
	// If the object is a folder, contains the filenames of the files within the
	// folder.
	Children []string `json:"children"`
	Hash     string   `json:"hash"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		Folder      respjson.Field
		Name        respjson.Field
		UpdatedAt   respjson.Field
		ArchivedAt  respjson.Field
		Children    respjson.Field
		Hash        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SourceCodeGetMetadataResponse) RawJSON() string { return r.JSON.raw }
func (r *SourceCodeGetMetadataResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SourceCodeUpsertResponse struct {
	// The path of the file in the CMS Developer File System.
	ID string `json:"id,required"`
	// Timestamp of when the object was first created.
	CreatedAt int64 `json:"createdAt,required"`
	// Determines whether or not this path points to a folder.
	Folder bool `json:"folder,required"`
	// The name of the file.
	Name string `json:"name,required"`
	// Timestamp of when the object was last updated.
	UpdatedAt int64 `json:"updatedAt,required"`
	// Timestamp of when the object was archived (deleted).
	ArchivedAt int64 `json:"archivedAt"`
	// If the object is a folder, contains the filenames of the files within the
	// folder.
	Children []string `json:"children"`
	Hash     string   `json:"hash"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		Folder      respjson.Field
		Name        respjson.Field
		UpdatedAt   respjson.Field
		ArchivedAt  respjson.Field
		Children    respjson.Field
		Hash        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SourceCodeUpsertResponse) RawJSON() string { return r.JSON.raw }
func (r *SourceCodeUpsertResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SourceCodeNewParams struct {
	Environment string    `path:"environment,required" json:"-"`
	File        io.Reader `json:"file,omitzero" format:"binary"`
	paramObj
}

func (r SourceCodeNewParams) MarshalMultipart() (data []byte, contentType string, err error) {
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

type SourceCodeDeleteParams struct {
	Environment string `path:"environment,required" json:"-"`
	paramObj
}

type SourceCodeExtractAsyncParams struct {
	Path string `json:"path,required"`
	paramObj
}

func (r SourceCodeExtractAsyncParams) MarshalJSON() (data []byte, err error) {
	type shadow SourceCodeExtractAsyncParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SourceCodeExtractAsyncParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SourceCodeGetParams struct {
	Environment string `path:"environment,required" json:"-"`
	paramObj
}

type SourceCodeGetMetadataParams struct {
	Environment string            `path:"environment,required" json:"-"`
	Properties  param.Opt[string] `query:"properties,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SourceCodeGetMetadataParams]'s query parameters as
// `url.Values`.
func (r SourceCodeGetMetadataParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SourceCodeUpsertParams struct {
	Environment string    `path:"environment,required" json:"-"`
	File        io.Reader `json:"file,omitzero" format:"binary"`
	paramObj
}

func (r SourceCodeUpsertParams) MarshalMultipart() (data []byte, contentType string, err error) {
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

type SourceCodeValidateParams struct {
	Environment string    `path:"environment,required" json:"-"`
	File        io.Reader `json:"file,omitzero" format:"binary"`
	paramObj
}

func (r SourceCodeValidateParams) MarshalMultipart() (data []byte, contentType string, err error) {
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
