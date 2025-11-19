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
	"github.com/stainless-sdks/hubspot-sdk-go/marketing"
	"github.com/stainless-sdks/hubspot-sdk-go/option"
	"github.com/stainless-sdks/hubspot-sdk-go/packages/pagination"
	"github.com/stainless-sdks/hubspot-sdk-go/packages/param"
	"github.com/stainless-sdks/hubspot-sdk-go/packages/respjson"
	"github.com/stainless-sdks/hubspot-sdk-go/shared"
)

// AssociationV4Service contains methods and other services that help with
// interacting with the Hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAssociationV4Service] method instead.
type AssociationV4Service struct {
	Options []option.RequestOption
	Batch   AssociationV4BatchService
	Report  AssociationV4ReportService
}

// NewAssociationV4Service generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAssociationV4Service(opts ...option.RequestOption) (r AssociationV4Service) {
	r = AssociationV4Service{}
	r.Options = opts
	r.Batch = NewAssociationV4BatchService(opts...)
	r.Report = NewAssociationV4ReportService(opts...)
	return
}

// Create the default (most generic) association type between two object types
func (r *AssociationV4Service) New(ctx context.Context, toObjectID string, body AssociationV4NewParams, opts ...option.RequestOption) (res *BatchResponsePublicDefaultAssociation, err error) {
	opts = slices.Concat(r.Options, opts)
	if body.FromObjectType == "" {
		err = errors.New("missing required fromObjectType parameter")
		return
	}
	if body.FromObjectID == "" {
		err = errors.New("missing required fromObjectId parameter")
		return
	}
	if body.ToObjectType == "" {
		err = errors.New("missing required toObjectType parameter")
		return
	}
	if toObjectID == "" {
		err = errors.New("missing required toObjectId parameter")
		return
	}
	path := fmt.Sprintf("crm/v4/objects/%s/%s/associations/default/%s/%s", body.FromObjectType, body.FromObjectID, body.ToObjectType, toObjectID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, nil, &res, opts...)
	return
}

// Set association labels between two records.
func (r *AssociationV4Service) Update(ctx context.Context, toObjectID string, params AssociationV4UpdateParams, opts ...option.RequestOption) (res *CreatedResponseLabelsBetweenObjectPair, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.ObjectType == "" {
		err = errors.New("missing required objectType parameter")
		return
	}
	if params.ObjectID == "" {
		err = errors.New("missing required objectId parameter")
		return
	}
	if params.ToObjectType == "" {
		err = errors.New("missing required toObjectType parameter")
		return
	}
	if toObjectID == "" {
		err = errors.New("missing required toObjectId parameter")
		return
	}
	path := fmt.Sprintf("crm/v4/objects/%s/%s/associations/%s/%s", params.ObjectType, params.ObjectID, params.ToObjectType, toObjectID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return
}

// List all associations of an object by object type. Limit 500 per call.
func (r *AssociationV4Service) List(ctx context.Context, toObjectType string, params AssociationV4ListParams, opts ...option.RequestOption) (res *pagination.Page[MultiAssociatedObjectWithLabel], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.ObjectType == "" {
		err = errors.New("missing required objectType parameter")
		return
	}
	if params.ObjectID == "" {
		err = errors.New("missing required objectId parameter")
		return
	}
	if toObjectType == "" {
		err = errors.New("missing required toObjectType parameter")
		return
	}
	path := fmt.Sprintf("crm/v4/objects/%s/%s/associations/%s", params.ObjectType, params.ObjectID, toObjectType)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, params, &res, opts...)
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

// List all associations of an object by object type. Limit 500 per call.
func (r *AssociationV4Service) ListAutoPaging(ctx context.Context, toObjectType string, params AssociationV4ListParams, opts ...option.RequestOption) *pagination.PageAutoPager[MultiAssociatedObjectWithLabel] {
	return pagination.NewPageAutoPager(r.List(ctx, toObjectType, params, opts...))
}

// deletes all associations between two records.
func (r *AssociationV4Service) Delete(ctx context.Context, toObjectID string, body AssociationV4DeleteParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if body.ObjectType == "" {
		err = errors.New("missing required objectType parameter")
		return
	}
	if body.ObjectID == "" {
		err = errors.New("missing required objectId parameter")
		return
	}
	if body.ToObjectType == "" {
		err = errors.New("missing required toObjectType parameter")
		return
	}
	if toObjectID == "" {
		err = errors.New("missing required toObjectId parameter")
		return
	}
	path := fmt.Sprintf("crm/v4/objects/%s/%s/associations/%s/%s", body.ObjectType, body.ObjectID, body.ToObjectType, toObjectID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Defines the type, direction, and details of the relationship between two CRM
// objects.
type AssociationSpec1 struct {
	// The category of the association, such as "HUBSPOT_DEFINED".
	//
	// Any of "HUBSPOT_DEFINED", "USER_DEFINED", "INTEGRATOR_DEFINED".
	AssociationCategory AssociationSpec1AssociationCategory `json:"associationCategory,required"`
	// The ID representing the specific type of association.
	AssociationTypeID int64 `json:"associationTypeId,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AssociationCategory respjson.Field
		AssociationTypeID   respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AssociationSpec1) RawJSON() string { return r.JSON.raw }
func (r *AssociationSpec1) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this AssociationSpec1 to a AssociationSpec1Param.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// AssociationSpec1Param.Overrides()
func (r AssociationSpec1) ToParam() AssociationSpec1Param {
	return param.Override[AssociationSpec1Param](json.RawMessage(r.RawJSON()))
}

// The category of the association, such as "HUBSPOT_DEFINED".
type AssociationSpec1AssociationCategory string

const (
	AssociationSpec1AssociationCategoryHubspotDefined    AssociationSpec1AssociationCategory = "HUBSPOT_DEFINED"
	AssociationSpec1AssociationCategoryUserDefined       AssociationSpec1AssociationCategory = "USER_DEFINED"
	AssociationSpec1AssociationCategoryIntegratorDefined AssociationSpec1AssociationCategory = "INTEGRATOR_DEFINED"
)

// Defines the type, direction, and details of the relationship between two CRM
// objects.
//
// The properties AssociationCategory, AssociationTypeID are required.
type AssociationSpec1Param struct {
	// The category of the association, such as "HUBSPOT_DEFINED".
	//
	// Any of "HUBSPOT_DEFINED", "USER_DEFINED", "INTEGRATOR_DEFINED".
	AssociationCategory AssociationSpec1AssociationCategory `json:"associationCategory,omitzero,required"`
	// The ID representing the specific type of association.
	AssociationTypeID int64 `json:"associationTypeId,required"`
	paramObj
}

func (r AssociationSpec1Param) MarshalJSON() (data []byte, err error) {
	type shadow AssociationSpec1Param
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AssociationSpec1Param) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AssociationSpecWithLabel1 struct {
	// Any of "HUBSPOT_DEFINED", "USER_DEFINED", "INTEGRATOR_DEFINED".
	Category AssociationSpecWithLabel1Category `json:"category,required"`
	TypeID   int64                             `json:"typeId,required"`
	Label    string                            `json:"label"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Category    respjson.Field
		TypeID      respjson.Field
		Label       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AssociationSpecWithLabel1) RawJSON() string { return r.JSON.raw }
func (r *AssociationSpecWithLabel1) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AssociationSpecWithLabel1Category string

const (
	AssociationSpecWithLabel1CategoryHubspotDefined    AssociationSpecWithLabel1Category = "HUBSPOT_DEFINED"
	AssociationSpecWithLabel1CategoryUserDefined       AssociationSpecWithLabel1Category = "USER_DEFINED"
	AssociationSpecWithLabel1CategoryIntegratorDefined AssociationSpecWithLabel1Category = "INTEGRATOR_DEFINED"
)

// The property Inputs is required.
type BatchInputPublicAssociationMultiArchiveParam struct {
	Inputs []PublicAssociationMultiArchiveParam `json:"inputs,omitzero,required"`
	paramObj
}

func (r BatchInputPublicAssociationMultiArchiveParam) MarshalJSON() (data []byte, err error) {
	type shadow BatchInputPublicAssociationMultiArchiveParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BatchInputPublicAssociationMultiArchiveParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Inputs is required.
type BatchInputPublicAssociationMultiPostParam struct {
	Inputs []PublicAssociationMultiPostParam `json:"inputs,omitzero,required"`
	paramObj
}

func (r BatchInputPublicAssociationMultiPostParam) MarshalJSON() (data []byte, err error) {
	type shadow BatchInputPublicAssociationMultiPostParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BatchInputPublicAssociationMultiPostParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Inputs is required.
type BatchInputPublicDefaultAssociationMultiPostParam struct {
	Inputs []PublicDefaultAssociationMultiPostParam `json:"inputs,omitzero,required"`
	paramObj
}

func (r BatchInputPublicDefaultAssociationMultiPostParam) MarshalJSON() (data []byte, err error) {
	type shadow BatchInputPublicDefaultAssociationMultiPostParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BatchInputPublicDefaultAssociationMultiPostParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Inputs is required.
type BatchInputPublicFetchAssociationsBatchRequestParam struct {
	Inputs []PublicFetchAssociationsBatchRequestParam `json:"inputs,omitzero,required"`
	paramObj
}

func (r BatchInputPublicFetchAssociationsBatchRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow BatchInputPublicFetchAssociationsBatchRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BatchInputPublicFetchAssociationsBatchRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BatchResponseLabelsBetweenObjectPair struct {
	CompletedAt time.Time                 `json:"completedAt,required" format:"date-time"`
	Results     []LabelsBetweenObjectPair `json:"results,required"`
	StartedAt   time.Time                 `json:"startedAt,required" format:"date-time"`
	// Any of "PENDING", "PROCESSING", "CANCELED", "COMPLETE".
	Status      BatchResponseLabelsBetweenObjectPairStatus `json:"status,required"`
	Errors      []StandardError1                           `json:"errors"`
	Links       map[string]string                          `json:"links"`
	NumErrors   int64                                      `json:"numErrors"`
	RequestedAt time.Time                                  `json:"requestedAt" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CompletedAt respjson.Field
		Results     respjson.Field
		StartedAt   respjson.Field
		Status      respjson.Field
		Errors      respjson.Field
		Links       respjson.Field
		NumErrors   respjson.Field
		RequestedAt respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BatchResponseLabelsBetweenObjectPair) RawJSON() string { return r.JSON.raw }
func (r *BatchResponseLabelsBetweenObjectPair) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BatchResponseLabelsBetweenObjectPairStatus string

const (
	BatchResponseLabelsBetweenObjectPairStatusPending    BatchResponseLabelsBetweenObjectPairStatus = "PENDING"
	BatchResponseLabelsBetweenObjectPairStatusProcessing BatchResponseLabelsBetweenObjectPairStatus = "PROCESSING"
	BatchResponseLabelsBetweenObjectPairStatusCanceled   BatchResponseLabelsBetweenObjectPairStatus = "CANCELED"
	BatchResponseLabelsBetweenObjectPairStatusComplete   BatchResponseLabelsBetweenObjectPairStatus = "COMPLETE"
)

type BatchResponsePublicAssociationMultiWithLabel struct {
	CompletedAt time.Time                         `json:"completedAt,required" format:"date-time"`
	Results     []PublicAssociationMultiWithLabel `json:"results,required"`
	StartedAt   time.Time                         `json:"startedAt,required" format:"date-time"`
	// Any of "PENDING", "PROCESSING", "CANCELED", "COMPLETE".
	Status      BatchResponsePublicAssociationMultiWithLabelStatus `json:"status,required"`
	Errors      []StandardError1                                   `json:"errors"`
	Links       map[string]string                                  `json:"links"`
	NumErrors   int64                                              `json:"numErrors"`
	RequestedAt time.Time                                          `json:"requestedAt" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CompletedAt respjson.Field
		Results     respjson.Field
		StartedAt   respjson.Field
		Status      respjson.Field
		Errors      respjson.Field
		Links       respjson.Field
		NumErrors   respjson.Field
		RequestedAt respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BatchResponsePublicAssociationMultiWithLabel) RawJSON() string { return r.JSON.raw }
func (r *BatchResponsePublicAssociationMultiWithLabel) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BatchResponsePublicAssociationMultiWithLabelStatus string

const (
	BatchResponsePublicAssociationMultiWithLabelStatusPending    BatchResponsePublicAssociationMultiWithLabelStatus = "PENDING"
	BatchResponsePublicAssociationMultiWithLabelStatusProcessing BatchResponsePublicAssociationMultiWithLabelStatus = "PROCESSING"
	BatchResponsePublicAssociationMultiWithLabelStatusCanceled   BatchResponsePublicAssociationMultiWithLabelStatus = "CANCELED"
	BatchResponsePublicAssociationMultiWithLabelStatusComplete   BatchResponsePublicAssociationMultiWithLabelStatus = "COMPLETE"
)

type BatchResponseVoid struct {
	CompletedAt time.Time `json:"completedAt,required" format:"date-time"`
	Results     []any     `json:"results,required"`
	StartedAt   time.Time `json:"startedAt,required" format:"date-time"`
	// Any of "PENDING", "PROCESSING", "CANCELED", "COMPLETE".
	Status      BatchResponseVoidStatus `json:"status,required"`
	Errors      []StandardError1        `json:"errors"`
	Links       map[string]string       `json:"links"`
	NumErrors   int64                   `json:"numErrors"`
	RequestedAt time.Time               `json:"requestedAt" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CompletedAt respjson.Field
		Results     respjson.Field
		StartedAt   respjson.Field
		Status      respjson.Field
		Errors      respjson.Field
		Links       respjson.Field
		NumErrors   respjson.Field
		RequestedAt respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BatchResponseVoid) RawJSON() string { return r.JSON.raw }
func (r *BatchResponseVoid) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BatchResponseVoidStatus string

const (
	BatchResponseVoidStatusPending    BatchResponseVoidStatus = "PENDING"
	BatchResponseVoidStatusProcessing BatchResponseVoidStatus = "PROCESSING"
	BatchResponseVoidStatusCanceled   BatchResponseVoidStatus = "CANCELED"
	BatchResponseVoidStatusComplete   BatchResponseVoidStatus = "COMPLETE"
)

type DateTime struct {
	DateOnly      bool  `json:"dateOnly,required"`
	TimeZoneShift int64 `json:"timeZoneShift,required"`
	Value         int64 `json:"value,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DateOnly      respjson.Field
		TimeZoneShift respjson.Field
		Value         respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DateTime) RawJSON() string { return r.JSON.raw }
func (r *DateTime) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties From, To are required.
type PublicAssociationMultiArchiveParam struct {
	From shared.PublicObjectIDParam   `json:"from,omitzero,required"`
	To   []shared.PublicObjectIDParam `json:"to,omitzero,required"`
	paramObj
}

func (r PublicAssociationMultiArchiveParam) MarshalJSON() (data []byte, err error) {
	type shadow PublicAssociationMultiArchiveParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PublicAssociationMultiArchiveParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties From, To, Types are required.
type PublicAssociationMultiPostParam struct {
	From  shared.PublicObjectIDParam `json:"from,omitzero,required"`
	To    shared.PublicObjectIDParam `json:"to,omitzero,required"`
	Types []AssociationSpec1Param    `json:"types,omitzero,required"`
	paramObj
}

func (r PublicAssociationMultiPostParam) MarshalJSON() (data []byte, err error) {
	type shadow PublicAssociationMultiPostParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PublicAssociationMultiPostParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PublicAssociationMultiWithLabel struct {
	From shared.PublicObjectID            `json:"from,required"`
	To   []MultiAssociatedObjectWithLabel `json:"to,required"`
	// Contains information pagination of results.
	Paging marketing.Paging `json:"paging"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		From        respjson.Field
		To          respjson.Field
		Paging      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PublicAssociationMultiWithLabel) RawJSON() string { return r.JSON.raw }
func (r *PublicAssociationMultiWithLabel) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties From, To are required.
type PublicDefaultAssociationMultiPostParam struct {
	From shared.PublicObjectIDParam `json:"from,omitzero,required"`
	To   shared.PublicObjectIDParam `json:"to,omitzero,required"`
	paramObj
}

func (r PublicDefaultAssociationMultiPostParam) MarshalJSON() (data []byte, err error) {
	type shadow PublicDefaultAssociationMultiPostParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PublicDefaultAssociationMultiPostParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property ID is required.
type PublicFetchAssociationsBatchRequestParam struct {
	ID    string            `json:"id,required"`
	After param.Opt[string] `json:"after,omitzero"`
	paramObj
}

func (r PublicFetchAssociationsBatchRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow PublicFetchAssociationsBatchRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PublicFetchAssociationsBatchRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ReportCreationResponse struct {
	EnqueueTime DateTime `json:"enqueueTime,required"`
	UserEmail   string   `json:"userEmail,required"`
	UserID      int64    `json:"userId,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		EnqueueTime respjson.Field
		UserEmail   respjson.Field
		UserID      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ReportCreationResponse) RawJSON() string { return r.JSON.raw }
func (r *ReportCreationResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Ye olde error
type StandardError1 struct {
	// The main category of the error.
	Category string `json:"category,required"`
	// Additional context-specific information related to the error.
	Context map[string][]string `json:"context,required"`
	// The detailed error objects.
	Errors []shared.ErrorDetail `json:"errors,required"`
	// URLs linking to documentation or resources associated with the error.
	Links map[string]string `json:"links,required"`
	// A human-readable string describing the error and possible remediation steps.
	Message string `json:"message,required"`
	// The HTTP status code associated with the error.
	Status string `json:"status,required"`
	// A unique ID for the error instance.
	ID string `json:"id"`
	// A more specific error category within each main category.
	SubCategory any `json:"subCategory"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Category    respjson.Field
		Context     respjson.Field
		Errors      respjson.Field
		Links       respjson.Field
		Message     respjson.Field
		Status      respjson.Field
		ID          respjson.Field
		SubCategory respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r StandardError1) RawJSON() string { return r.JSON.raw }
func (r *StandardError1) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AssociationV4NewParams struct {
	FromObjectType string `path:"fromObjectType,required" json:"-"`
	FromObjectID   string `path:"fromObjectId,required" json:"-"`
	ToObjectType   string `path:"toObjectType,required" json:"-"`
	paramObj
}

type AssociationV4UpdateParams struct {
	ObjectType   string `path:"objectType,required" json:"-"`
	ObjectID     string `path:"objectId,required" json:"-"`
	ToObjectType string `path:"toObjectType,required" json:"-"`
	Body         []shared.AssociationSpecParam
	paramObj
}

func (r AssociationV4UpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.Body)
}
func (r *AssociationV4UpdateParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.Body)
}

type AssociationV4ListParams struct {
	ObjectType string `path:"objectType,required" json:"-"`
	ObjectID   string `path:"objectId,required" json:"-"`
	// The paging cursor token of the last successfully read resource will be returned
	// as the `paging.next.after` JSON property of a paged response containing more
	// results.
	After param.Opt[string] `query:"after,omitzero" json:"-"`
	// The maximum number of results to display per page.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [AssociationV4ListParams]'s query parameters as
// `url.Values`.
func (r AssociationV4ListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AssociationV4DeleteParams struct {
	ObjectType   string `path:"objectType,required" json:"-"`
	ObjectID     string `path:"objectId,required" json:"-"`
	ToObjectType string `path:"toObjectType,required" json:"-"`
	paramObj
}
