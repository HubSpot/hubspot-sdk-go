// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cms

import (
	"context"
	"fmt"
	"net/http"
	"slices"

	"github.com/HubSpot/hubspot-sdk-go/internal/apijson"
	shimjson "github.com/HubSpot/hubspot-sdk-go/internal/encoding/json"
	"github.com/HubSpot/hubspot-sdk-go/internal/requestconfig"
	"github.com/HubSpot/hubspot-sdk-go/option"
	"github.com/HubSpot/hubspot-sdk-go/packages/param"
	"github.com/HubSpot/hubspot-sdk-go/shared"
)

// SourceCodeService contains methods and other services that help with interacting
// with the hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSourceCodeService] method instead.
type SourceCodeService struct {
	options []option.RequestOption
}

// NewSourceCodeService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewSourceCodeService(opts ...option.RequestOption) (r SourceCodeService) {
	r = SourceCodeService{}
	r.options = opts
	return
}

// Extract a zip file in the developer file system. Extraction status can be
// checked with the `/extract/async/tasks/taskId/status` endpoint below.
func (r *SourceCodeService) ExtractAsync(ctx context.Context, body SourceCodeExtractAsyncParams, opts ...option.RequestOption) (res *shared.TaskLocator, err error) {
	opts = slices.Concat(r.options, opts)
	path := "cms/source-code/2026-03/extract/async"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Get the status of an extraction by the `taskId` returned from the initial
// `extract/async` request.
func (r *SourceCodeService) GetExtractionStatus(ctx context.Context, taskID int64, opts ...option.RequestOption) (res *shared.ActionResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := fmt.Sprintf("cms/source-code/2026-03/extract/async/tasks/%v/status", taskID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// The property Path is required.
type FileExtractRequestParam struct {
	// The file system location where the zip file is to be extracted.
	Path string `json:"path" api:"required"`
	paramObj
}

func (r FileExtractRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow FileExtractRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FileExtractRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SourceCodeExtractAsyncParams struct {
	FileExtractRequest FileExtractRequestParam
	paramObj
}

func (r SourceCodeExtractAsyncParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.FileExtractRequest)
}
func (r *SourceCodeExtractAsyncParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
