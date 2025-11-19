// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package crm

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"slices"

	shimjson "github.com/stainless-sdks/hubspot-sdk-go/internal/encoding/json"
	"github.com/stainless-sdks/hubspot-sdk-go/internal/requestconfig"
	"github.com/stainless-sdks/hubspot-sdk-go/option"
)

// PipelineStageService contains methods and other services that help with
// interacting with the Hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPipelineStageService] method instead.
type PipelineStageService struct {
	Options []option.RequestOption
}

// NewPipelineStageService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewPipelineStageService(opts ...option.RequestOption) (r PipelineStageService) {
	r = PipelineStageService{}
	r.Options = opts
	return
}

// Create a new stage associated with the pipeline identified by `{pipelineId}`.
// The entire stage object, including its unique ID, will be returned in the
// response.
func (r *PipelineStageService) New(ctx context.Context, pipelineID string, params PipelineStageNewParams, opts ...option.RequestOption) (res *PipelineStage, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.ObjectType == "" {
		err = errors.New("missing required objectType parameter")
		return
	}
	if pipelineID == "" {
		err = errors.New("missing required pipelineId parameter")
		return
	}
	path := fmt.Sprintf("crm/v3/pipelines/%s/%s/stages", params.ObjectType, pipelineID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Perform a partial update of the pipeline stage identified by `{stageId}`
// associated with the pipeline identified by `{pipelineId}`. Any properties not
// included in this update will keep their existing values. The updated stage will
// be returned in the response.
func (r *PipelineStageService) Update(ctx context.Context, stageID string, params PipelineStageUpdateParams, opts ...option.RequestOption) (res *PipelineStage, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.ObjectType == "" {
		err = errors.New("missing required objectType parameter")
		return
	}
	if params.PipelineID == "" {
		err = errors.New("missing required pipelineId parameter")
		return
	}
	if stageID == "" {
		err = errors.New("missing required stageId parameter")
		return
	}
	path := fmt.Sprintf("crm/v3/pipelines/%s/%s/stages/%s", params.ObjectType, params.PipelineID, stageID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return
}

// Return all the stages associated with the pipeline identified by `{pipelineId}`.
func (r *PipelineStageService) List(ctx context.Context, pipelineID string, query PipelineStageListParams, opts ...option.RequestOption) (res *CollectionResponsePipelineStageNoPaging, err error) {
	opts = slices.Concat(r.Options, opts)
	if query.ObjectType == "" {
		err = errors.New("missing required objectType parameter")
		return
	}
	if pipelineID == "" {
		err = errors.New("missing required pipelineId parameter")
		return
	}
	path := fmt.Sprintf("crm/v3/pipelines/%s/%s/stages", query.ObjectType, pipelineID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Delete the pipeline stage identified by `{stageId}` associated with the pipeline
// identified by `{pipelineId}`.
func (r *PipelineStageService) Delete(ctx context.Context, stageID string, body PipelineStageDeleteParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if body.ObjectType == "" {
		err = errors.New("missing required objectType parameter")
		return
	}
	if body.PipelineID == "" {
		err = errors.New("missing required pipelineId parameter")
		return
	}
	if stageID == "" {
		err = errors.New("missing required stageId parameter")
		return
	}
	path := fmt.Sprintf("crm/v3/pipelines/%s/%s/stages/%s", body.ObjectType, body.PipelineID, stageID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Return the stage identified by `{stageId}` associated with the pipeline
// identified by `{pipelineId}`.
func (r *PipelineStageService) Get(ctx context.Context, stageID string, query PipelineStageGetParams, opts ...option.RequestOption) (res *PipelineStage, err error) {
	opts = slices.Concat(r.Options, opts)
	if query.ObjectType == "" {
		err = errors.New("missing required objectType parameter")
		return
	}
	if query.PipelineID == "" {
		err = errors.New("missing required pipelineId parameter")
		return
	}
	if stageID == "" {
		err = errors.New("missing required stageId parameter")
		return
	}
	path := fmt.Sprintf("crm/v3/pipelines/%s/%s/stages/%s", query.ObjectType, query.PipelineID, stageID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Return a reverse chronological list of all mutations that have occurred on the
// pipeline stage identified by `{stageId}`.
func (r *PipelineStageService) GetAudit(ctx context.Context, stageID string, query PipelineStageGetAuditParams, opts ...option.RequestOption) (res *CollectionResponsePublicAuditInfoNoPaging, err error) {
	opts = slices.Concat(r.Options, opts)
	if query.ObjectType == "" {
		err = errors.New("missing required objectType parameter")
		return
	}
	if query.PipelineID == "" {
		err = errors.New("missing required pipelineId parameter")
		return
	}
	if stageID == "" {
		err = errors.New("missing required stageId parameter")
		return
	}
	path := fmt.Sprintf("crm/v3/pipelines/%s/%s/stages/%s/audit", query.ObjectType, query.PipelineID, stageID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Replace all the properties of an existing pipeline stage with the values
// provided. The updated stage will be returned in the response.
func (r *PipelineStageService) Replace(ctx context.Context, stageID string, params PipelineStageReplaceParams, opts ...option.RequestOption) (res *PipelineStage, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.ObjectType == "" {
		err = errors.New("missing required objectType parameter")
		return
	}
	if params.PipelineID == "" {
		err = errors.New("missing required pipelineId parameter")
		return
	}
	if stageID == "" {
		err = errors.New("missing required stageId parameter")
		return
	}
	path := fmt.Sprintf("crm/v3/pipelines/%s/%s/stages/%s", params.ObjectType, params.PipelineID, stageID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return
}

type PipelineStageNewParams struct {
	ObjectType string `path:"objectType,required" json:"-"`
	// An input used to create or replace a pipeline stage's definition.
	PipelineStageInput PipelineStageInputParam
	paramObj
}

func (r PipelineStageNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.PipelineStageInput)
}
func (r *PipelineStageNewParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.PipelineStageInput)
}

type PipelineStageUpdateParams struct {
	ObjectType string `path:"objectType,required" json:"-"`
	PipelineID string `path:"pipelineId,required" json:"-"`
	// An input used to update some properties on a pipeline definition.
	PipelineStagePatchInput PipelineStagePatchInputParam
	paramObj
}

func (r PipelineStageUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.PipelineStagePatchInput)
}
func (r *PipelineStageUpdateParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.PipelineStagePatchInput)
}

type PipelineStageListParams struct {
	ObjectType string `path:"objectType,required" json:"-"`
	paramObj
}

type PipelineStageDeleteParams struct {
	ObjectType string `path:"objectType,required" json:"-"`
	PipelineID string `path:"pipelineId,required" json:"-"`
	paramObj
}

type PipelineStageGetParams struct {
	ObjectType string `path:"objectType,required" json:"-"`
	PipelineID string `path:"pipelineId,required" json:"-"`
	paramObj
}

type PipelineStageGetAuditParams struct {
	ObjectType string `path:"objectType,required" json:"-"`
	PipelineID string `path:"pipelineId,required" json:"-"`
	paramObj
}

type PipelineStageReplaceParams struct {
	ObjectType string `path:"objectType,required" json:"-"`
	PipelineID string `path:"pipelineId,required" json:"-"`
	// An input used to create or replace a pipeline stage's definition.
	PipelineStageInput PipelineStageInputParam
	paramObj
}

func (r PipelineStageReplaceParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.PipelineStageInput)
}
func (r *PipelineStageReplaceParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.PipelineStageInput)
}
