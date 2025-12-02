// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package crm

import (
	"context"
	"encoding/json"
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
	"github.com/stainless-sdks/hubspot-sdk-go/packages/param"
	"github.com/stainless-sdks/hubspot-sdk-go/packages/respjson"
)

// PipelineService contains methods and other services that help with interacting
// with the hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPipelineService] method instead.
type PipelineService struct {
	Options []option.RequestOption
	Stages  PipelineStageService
}

// NewPipelineService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewPipelineService(opts ...option.RequestOption) (r PipelineService) {
	r = PipelineService{}
	r.Options = opts
	r.Stages = NewPipelineStageService(opts...)
	return
}

// Create a new pipeline with the provided property values. The entire pipeline
// object, including its unique ID, will be returned in the response.
func (r *PipelineService) New(ctx context.Context, objectType string, body PipelineNewParams, opts ...option.RequestOption) (res *Pipeline, err error) {
	opts = slices.Concat(r.Options, opts)
	if objectType == "" {
		err = errors.New("missing required objectType parameter")
		return
	}
	path := fmt.Sprintf("crm/v3/pipelines/%s", objectType)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Perform a partial update of the pipeline identified by `{pipelineId}`. The
// updated pipeline will be returned in the response.
func (r *PipelineService) Update(ctx context.Context, pipelineID string, params PipelineUpdateParams, opts ...option.RequestOption) (res *Pipeline, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.ObjectType == "" {
		err = errors.New("missing required objectType parameter")
		return
	}
	if pipelineID == "" {
		err = errors.New("missing required pipelineId parameter")
		return
	}
	path := fmt.Sprintf("crm/v3/pipelines/%s/%s", params.ObjectType, pipelineID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return
}

// Return all pipelines for the object type specified by `{objectType}`.
func (r *PipelineService) List(ctx context.Context, objectType string, opts ...option.RequestOption) (res *CollectionResponsePipelineNoPaging, err error) {
	opts = slices.Concat(r.Options, opts)
	if objectType == "" {
		err = errors.New("missing required objectType parameter")
		return
	}
	path := fmt.Sprintf("crm/v3/pipelines/%s", objectType)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Delete a pipeline identified by its unique pipelineId
func (r *PipelineService) Delete(ctx context.Context, pipelineID string, params PipelineDeleteParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if params.ObjectType == "" {
		err = errors.New("missing required objectType parameter")
		return
	}
	if pipelineID == "" {
		err = errors.New("missing required pipelineId parameter")
		return
	}
	path := fmt.Sprintf("crm/v3/pipelines/%s/%s", params.ObjectType, pipelineID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, params, nil, opts...)
	return
}

// Return a single pipeline object identified by its unique `{pipelineId}`.
func (r *PipelineService) Get(ctx context.Context, pipelineID string, query PipelineGetParams, opts ...option.RequestOption) (res *Pipeline, err error) {
	opts = slices.Concat(r.Options, opts)
	if query.ObjectType == "" {
		err = errors.New("missing required objectType parameter")
		return
	}
	if pipelineID == "" {
		err = errors.New("missing required pipelineId parameter")
		return
	}
	path := fmt.Sprintf("crm/v3/pipelines/%s/%s", query.ObjectType, pipelineID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Return a reverse chronological list of all mutations that have occurred on the
// pipeline identified by `{pipelineId}`.
func (r *PipelineService) GetAudit(ctx context.Context, pipelineID string, query PipelineGetAuditParams, opts ...option.RequestOption) (res *CollectionResponsePublicAuditInfoNoPaging, err error) {
	opts = slices.Concat(r.Options, opts)
	if query.ObjectType == "" {
		err = errors.New("missing required objectType parameter")
		return
	}
	if pipelineID == "" {
		err = errors.New("missing required pipelineId parameter")
		return
	}
	path := fmt.Sprintf("crm/v3/pipelines/%s/%s/audit", query.ObjectType, pipelineID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Replace all properties of an existing pipeline with the provided values.
func (r *PipelineService) Replace(ctx context.Context, pipelineID string, params PipelineReplaceParams, opts ...option.RequestOption) (res *Pipeline, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.ObjectType == "" {
		err = errors.New("missing required objectType parameter")
		return
	}
	if pipelineID == "" {
		err = errors.New("missing required pipelineId parameter")
		return
	}
	path := fmt.Sprintf("crm/v3/pipelines/%s/%s", params.ObjectType, pipelineID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return
}

type CollectionResponsePipelineNoPaging struct {
	Results []Pipeline `json:"results,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Results     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CollectionResponsePipelineNoPaging) RawJSON() string { return r.JSON.raw }
func (r *CollectionResponsePipelineNoPaging) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CollectionResponsePipelineStageNoPaging struct {
	Results []PipelineStage `json:"results,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Results     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CollectionResponsePipelineStageNoPaging) RawJSON() string { return r.JSON.raw }
func (r *CollectionResponsePipelineStageNoPaging) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CollectionResponsePublicAuditInfoNoPaging struct {
	Results []PublicAuditInfo `json:"results,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Results     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CollectionResponsePublicAuditInfoNoPaging) RawJSON() string { return r.JSON.raw }
func (r *CollectionResponsePublicAuditInfoNoPaging) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A pipeline definition.
type Pipeline struct {
	// A unique identifier generated by HubSpot that can be used to retrieve and update
	// the pipeline.
	ID string `json:"id,required"`
	// Whether the pipeline is archived.
	Archived bool `json:"archived,required"`
	// The date the pipeline was created. The default pipelines will have createdAt
	// = 0.
	CreatedAt time.Time `json:"createdAt,required" format:"date-time"`
	// The order for displaying this pipeline. If two pipelines have a matching
	// `displayOrder`, they will be sorted alphabetically by label.
	DisplayOrder int64 `json:"displayOrder,required"`
	// A unique label used to organize pipelines in HubSpot's UI
	Label string `json:"label,required"`
	// The stages associated with the pipeline. They can be retrieved and updated via
	// the pipeline stages endpoints.
	Stages []PipelineStage `json:"stages,required"`
	// The date the pipeline was last updated.
	UpdatedAt time.Time `json:"updatedAt,required" format:"date-time"`
	// The date the pipeline was archived. `archivedAt` will only be present if the
	// pipeline is archived.
	ArchivedAt time.Time `json:"archivedAt" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		Archived     respjson.Field
		CreatedAt    respjson.Field
		DisplayOrder respjson.Field
		Label        respjson.Field
		Stages       respjson.Field
		UpdatedAt    respjson.Field
		ArchivedAt   respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Pipeline) RawJSON() string { return r.JSON.raw }
func (r *Pipeline) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An input used to create or replace a pipeline's definition.
//
// The properties DisplayOrder, Label, Stages are required.
type PipelineInputParam struct {
	// The order for displaying this pipeline. If two pipelines have a matching
	// `displayOrder`, they will be sorted alphabetically by label.
	DisplayOrder int64 `json:"displayOrder,required"`
	// A unique label used to organize pipelines in HubSpot's UI
	Label string `json:"label,required"`
	// Pipeline stage inputs used to create the new or replacement pipeline.
	Stages []PipelineStageInputParam `json:"stages,omitzero,required"`
	paramObj
}

func (r PipelineInputParam) MarshalJSON() (data []byte, err error) {
	type shadow PipelineInputParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PipelineInputParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An input used to update some properties on a pipeline definition.
type PipelinePatchInputParam struct {
	// Whether the pipeline is archived. This property should only be provided when
	// restoring an archived pipeline. If it's provided in any other call, the request
	// will fail and a `400 Bad Request` will be returned.
	Archived param.Opt[bool] `json:"archived,omitzero"`
	// The order for displaying this pipeline. If two pipelines have a matching
	// `displayOrder`, they will be sorted alphabetically by label.
	DisplayOrder param.Opt[int64] `json:"displayOrder,omitzero"`
	// A unique label used to organize pipelines in HubSpot's UI
	Label param.Opt[string] `json:"label,omitzero"`
	paramObj
}

func (r PipelinePatchInputParam) MarshalJSON() (data []byte, err error) {
	type shadow PipelinePatchInputParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PipelinePatchInputParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A pipeline stage definition.
type PipelineStage struct {
	// A unique identifier generated by HubSpot that can be used to retrieve and update
	// the pipeline stage.
	ID string `json:"id,required"`
	// Whether the pipeline is archived.
	Archived bool `json:"archived,required"`
	// The date the pipeline stage was created. The stages on default pipelines will
	// have createdAt = 0.
	CreatedAt time.Time `json:"createdAt,required" format:"date-time"`
	// The order for displaying this pipeline stage. If two pipeline stages have a
	// matching `displayOrder`, they will be sorted alphabetically by label.
	DisplayOrder int64 `json:"displayOrder,required"`
	// A label used to organize pipeline stages in HubSpot's UI. Each pipeline stage's
	// label must be unique within that pipeline.
	Label string `json:"label,required"`
	// A JSON object containing properties that are not present on all object
	// pipelines.
	//
	// For `deals` pipelines, the `probability` field is required
	// (`{ "probability": 0.5 }`), and represents the likelihood a deal will close.
	// Possible values are between 0.0 and 1.0 in increments of 0.1.
	//
	// For `tickets` pipelines, the `ticketState` field is optional
	// (`{ "ticketState": "OPEN" }`), and represents whether the ticket remains open or
	// has been closed by a member of your Support team. Possible values are `OPEN` or
	// `CLOSED`.
	Metadata map[string]string `json:"metadata,required"`
	// The date the pipeline stage was last updated.
	UpdatedAt time.Time `json:"updatedAt,required" format:"date-time"`
	// The date the pipeline was archived. `archivedAt` will only be present if the
	// pipeline is archived.
	ArchivedAt time.Time `json:"archivedAt" format:"date-time"`
	// Defines the level of write access for the pipeline stage, with possible values
	// being CRM_PERMISSIONS_ENFORCEMENT, READ_ONLY, or INTERNAL_ONLY.
	//
	// Any of "CRM_PERMISSIONS_ENFORCEMENT", "INTERNAL_ONLY", "READ_ONLY".
	WritePermissions PipelineStageWritePermissions `json:"writePermissions"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		Archived         respjson.Field
		CreatedAt        respjson.Field
		DisplayOrder     respjson.Field
		Label            respjson.Field
		Metadata         respjson.Field
		UpdatedAt        respjson.Field
		ArchivedAt       respjson.Field
		WritePermissions respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PipelineStage) RawJSON() string { return r.JSON.raw }
func (r *PipelineStage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Defines the level of write access for the pipeline stage, with possible values
// being CRM_PERMISSIONS_ENFORCEMENT, READ_ONLY, or INTERNAL_ONLY.
type PipelineStageWritePermissions string

const (
	PipelineStageWritePermissionsCrmPermissionsEnforcement PipelineStageWritePermissions = "CRM_PERMISSIONS_ENFORCEMENT"
	PipelineStageWritePermissionsInternalOnly              PipelineStageWritePermissions = "INTERNAL_ONLY"
	PipelineStageWritePermissionsReadOnly                  PipelineStageWritePermissions = "READ_ONLY"
)

// An input used to create or replace a pipeline stage's definition.
//
// The properties DisplayOrder, Label, Metadata are required.
type PipelineStageInputParam struct {
	// The order for displaying this pipeline stage. If two pipeline stages have a
	// matching `displayOrder`, they will be sorted alphabetically by label.
	DisplayOrder int64 `json:"displayOrder,required"`
	// A label used to organize pipeline stages in HubSpot's UI. Each pipeline stage's
	// label must be unique within that pipeline.
	Label string `json:"label,required"`
	// A JSON object containing properties that are not present on all object
	// pipelines.
	//
	// For `deals` pipelines, the `probability` field is required
	// (`{ "probability": 0.5 }`), and represents the likelihood a deal will close.
	// Possible values are between 0.0 and 1.0 in increments of 0.1.
	//
	// For `tickets` pipelines, the `ticketState` field is optional
	// (`{ "ticketState": "OPEN" }`), and represents whether the ticket remains open or
	// has been closed by a member of your Support team. Possible values are `OPEN` or
	// `CLOSED`.
	Metadata map[string]string `json:"metadata,omitzero,required"`
	paramObj
}

func (r PipelineStageInputParam) MarshalJSON() (data []byte, err error) {
	type shadow PipelineStageInputParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PipelineStageInputParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An input used to update some properties on a pipeline definition.
//
// The property Metadata is required.
type PipelineStagePatchInputParam struct {
	// A JSON object containing properties that are not present on all object
	// pipelines.
	//
	// For `deals` pipelines, the `probability` field is required
	// (`{ "probability": 0.5 }`), and represents the likelihood a deal will close.
	// Possible values are between 0.0 and 1.0 in increments of 0.1.
	//
	// For `tickets` pipelines, the `ticketState` field is optional
	// (`{ "ticketState": "OPEN" }`), and represents whether the ticket remains open or
	// has been closed by a member of your Support team. Possible values are `OPEN` or
	// `CLOSED`.
	Metadata map[string]string `json:"metadata,omitzero,required"`
	// Whether the pipeline is archived.
	Archived param.Opt[bool] `json:"archived,omitzero"`
	// The order for displaying this pipeline stage. If two pipeline stages have a
	// matching `displayOrder`, they will be sorted alphabetically by label.
	DisplayOrder param.Opt[int64] `json:"displayOrder,omitzero"`
	// A label used to organize pipeline stages in HubSpot's UI. Each pipeline stage's
	// label must be unique within that pipeline.
	Label param.Opt[string] `json:"label,omitzero"`
	paramObj
}

func (r PipelineStagePatchInputParam) MarshalJSON() (data []byte, err error) {
	type shadow PipelineStagePatchInputParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PipelineStagePatchInputParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PublicAuditInfo struct {
	// The action performed that triggered the audit event.
	Action string `json:"action,required"`
	// A unique string identifier for the audit event.
	Identifier string `json:"identifier,required"`
	// The unique identifier for the HubSpot portal where the audit event occurred.
	PortalID int64 `json:"portalId,required"`
	// The ID of the user who initiated the audit event.
	FromUserID int64 `json:"fromUserId"`
	// A descriptive message related to the audit event.
	Message string `json:"message"`
	// An object containing the raw data associated with the audit event.
	RawObject any `json:"rawObject"`
	// The date and time when the audit event took place.
	Timestamp time.Time `json:"timestamp" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Action      respjson.Field
		Identifier  respjson.Field
		PortalID    respjson.Field
		FromUserID  respjson.Field
		Message     respjson.Field
		RawObject   respjson.Field
		Timestamp   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PublicAuditInfo) RawJSON() string { return r.JSON.raw }
func (r *PublicAuditInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PipelineNewParams struct {
	// An input used to create or replace a pipeline's definition.
	PipelineInput PipelineInputParam
	paramObj
}

func (r PipelineNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.PipelineInput)
}
func (r *PipelineNewParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.PipelineInput)
}

type PipelineUpdateParams struct {
	ObjectType string `path:"objectType,required" json:"-"`
	// An input used to update some properties on a pipeline definition.
	PipelinePatchInput PipelinePatchInputParam
	// Indicates whether to validate deal stage usages before deleting the pipeline.
	ValidateDealStageUsagesBeforeDelete param.Opt[bool] `query:"validateDealStageUsagesBeforeDelete,omitzero" json:"-"`
	// Indicates whether to validate references before deleting the pipeline.
	ValidateReferencesBeforeDelete param.Opt[bool] `query:"validateReferencesBeforeDelete,omitzero" json:"-"`
	paramObj
}

func (r PipelineUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.PipelinePatchInput)
}
func (r *PipelineUpdateParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.PipelinePatchInput)
}

// URLQuery serializes [PipelineUpdateParams]'s query parameters as `url.Values`.
func (r PipelineUpdateParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type PipelineDeleteParams struct {
	ObjectType string `path:"objectType,required" json:"-"`
	// Indicates whether to validate deal stage usages before deleting the pipeline.
	ValidateDealStageUsagesBeforeDelete param.Opt[bool] `query:"validateDealStageUsagesBeforeDelete,omitzero" json:"-"`
	// Indicates whether to validate references before deleting the pipeline.
	ValidateReferencesBeforeDelete param.Opt[bool] `query:"validateReferencesBeforeDelete,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [PipelineDeleteParams]'s query parameters as `url.Values`.
func (r PipelineDeleteParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type PipelineGetParams struct {
	ObjectType string `path:"objectType,required" json:"-"`
	paramObj
}

type PipelineGetAuditParams struct {
	ObjectType string `path:"objectType,required" json:"-"`
	paramObj
}

type PipelineReplaceParams struct {
	ObjectType string `path:"objectType,required" json:"-"`
	// An input used to create or replace a pipeline's definition.
	PipelineInput PipelineInputParam
	// Indicates whether to validate deal stage usages before deleting the pipeline.
	ValidateDealStageUsagesBeforeDelete param.Opt[bool] `query:"validateDealStageUsagesBeforeDelete,omitzero" json:"-"`
	// Indicates whether to validate references before deleting the pipeline.
	ValidateReferencesBeforeDelete param.Opt[bool] `query:"validateReferencesBeforeDelete,omitzero" json:"-"`
	paramObj
}

func (r PipelineReplaceParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.PipelineInput)
}
func (r *PipelineReplaceParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.PipelineInput)
}

// URLQuery serializes [PipelineReplaceParams]'s query parameters as `url.Values`.
func (r PipelineReplaceParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
