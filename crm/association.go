// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package crm

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
	"github.com/HubSpot/hubspot-sdk-go/packages/respjson"
	"github.com/HubSpot/hubspot-sdk-go/shared"
)

// AssociationService contains methods and other services that help with
// interacting with the hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAssociationService] method instead.
type AssociationService struct {
	options []option.RequestOption
	Batch   AssociationBatchService
}

// NewAssociationService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAssociationService(opts ...option.RequestOption) (r AssociationService) {
	r = AssociationService{}
	r.options = opts
	r.Batch = NewAssociationBatchService(opts...)
	return
}

// Retrieve all associations between a specific record and an object type. Limit
// 500 per call.
func (r *AssociationService) List(ctx context.Context, toObjectType string, params AssociationListParams, opts ...option.RequestOption) (res *pagination.Page[MultiAssociatedObjectWithLabel], err error) {
	var raw *http.Response
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.ObjectType == "" {
		err = errors.New("missing required objectType parameter")
		return nil, err
	}
	if params.ObjectID == "" {
		err = errors.New("missing required objectId parameter")
		return nil, err
	}
	if toObjectType == "" {
		err = errors.New("missing required toObjectType parameter")
		return nil, err
	}
	path := fmt.Sprintf("crm/objects/2026-03/%s/%s/associations/%s", url.PathEscape(params.ObjectType), url.PathEscape(params.ObjectID), url.PathEscape(toObjectType))
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

// Retrieve all associations between a specific record and an object type. Limit
// 500 per call.
func (r *AssociationService) ListAutoPaging(ctx context.Context, toObjectType string, params AssociationListParams, opts ...option.RequestOption) *pagination.PageAutoPager[MultiAssociatedObjectWithLabel] {
	return pagination.NewPageAutoPager(r.List(ctx, toObjectType, params, opts...))
}

func (r *AssociationService) Delete(ctx context.Context, toObjectID string, body AssociationDeleteParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if body.ObjectType == "" {
		err = errors.New("missing required objectType parameter")
		return err
	}
	if body.ObjectID == "" {
		err = errors.New("missing required objectId parameter")
		return err
	}
	if body.ToObjectType == "" {
		err = errors.New("missing required toObjectType parameter")
		return err
	}
	if toObjectID == "" {
		err = errors.New("missing required toObjectId parameter")
		return err
	}
	path := fmt.Sprintf("crm/objects/2026-03/%s/%s/associations/%s/%s", url.PathEscape(body.ObjectType), url.PathEscape(body.ObjectID), url.PathEscape(body.ToObjectType), url.PathEscape(toObjectID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Requests a report of all objects in the portal which have a high usage of
// associations
func (r *AssociationService) RequestHighUsageReport(ctx context.Context, userID int64, opts ...option.RequestOption) (res *ReportCreationResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := fmt.Sprintf("crm/associations/2026-03/usage/high-usage-report/%v", userID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

func (r *AssociationService) Search(ctx context.Context, objectType string, body AssociationSearchParams, opts ...option.RequestOption) (res *CollectionResponseWithTotalSimplePublicObject, err error) {
	opts = slices.Concat(r.options, opts)
	if objectType == "" {
		err = errors.New("missing required objectType parameter")
		return nil, err
	}
	path := fmt.Sprintf("crm/objects/2026-03/%s/search", url.PathEscape(objectType))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

func (r *AssociationService) UpdateAssociationLabels(ctx context.Context, toObjectID string, params AssociationUpdateAssociationLabelsParams, opts ...option.RequestOption) (res *LabelsBetweenObjectPair, err error) {
	opts = slices.Concat(r.options, opts)
	if params.ObjectType == "" {
		err = errors.New("missing required objectType parameter")
		return nil, err
	}
	if params.ObjectID == "" {
		err = errors.New("missing required objectId parameter")
		return nil, err
	}
	if params.ToObjectType == "" {
		err = errors.New("missing required toObjectType parameter")
		return nil, err
	}
	if toObjectID == "" {
		err = errors.New("missing required toObjectId parameter")
		return nil, err
	}
	path := fmt.Sprintf("crm/objects/2026-03/%s/%s/associations/%s/%s", url.PathEscape(params.ObjectType), url.PathEscape(params.ObjectID), url.PathEscape(params.ToObjectType), url.PathEscape(toObjectID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return res, err
}

// The property Inputs is required.
type BatchInputPublicAssociationMultiArchiveParam struct {
	Inputs []PublicAssociationMultiArchiveParam `json:"inputs,omitzero" api:"required"`
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
	Inputs []PublicAssociationMultiPostParam `json:"inputs,omitzero" api:"required"`
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
	Inputs []PublicDefaultAssociationMultiPostParam `json:"inputs,omitzero" api:"required"`
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
	Inputs []PublicFetchAssociationsBatchRequestParam `json:"inputs,omitzero" api:"required"`
	paramObj
}

func (r BatchInputPublicFetchAssociationsBatchRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow BatchInputPublicFetchAssociationsBatchRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BatchInputPublicFetchAssociationsBatchRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BatchResponsePublicAssociationMultiWithLabel struct {
	// The timestamp when the batch processing was completed, in ISO 8601 format.
	CompletedAt time.Time                         `json:"completedAt" api:"required" format:"date-time"`
	Results     []PublicAssociationMultiWithLabel `json:"results" api:"required"`
	// The timestamp when the batch processing began, in ISO 8601 format.
	StartedAt time.Time `json:"startedAt" api:"required" format:"date-time"`
	// The status of the batch processing request: "PENDING", "PROCESSING", "CANCELED",
	// or "COMPLETE".
	//
	// Any of "CANCELED", "COMPLETE", "PENDING", "PROCESSING".
	Status BatchResponsePublicAssociationMultiWithLabelStatus `json:"status" api:"required"`
	// An object containing relevant links related to the batch request.
	Links map[string]string `json:"links"`
	// The timestamp when the batch request was initially made, in ISO 8601 format.
	RequestedAt time.Time `json:"requestedAt" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CompletedAt respjson.Field
		Results     respjson.Field
		StartedAt   respjson.Field
		Status      respjson.Field
		Links       respjson.Field
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

// The status of the batch processing request: "PENDING", "PROCESSING", "CANCELED",
// or "COMPLETE".
type BatchResponsePublicAssociationMultiWithLabelStatus string

const (
	BatchResponsePublicAssociationMultiWithLabelStatusCanceled   BatchResponsePublicAssociationMultiWithLabelStatus = "CANCELED"
	BatchResponsePublicAssociationMultiWithLabelStatusComplete   BatchResponsePublicAssociationMultiWithLabelStatus = "COMPLETE"
	BatchResponsePublicAssociationMultiWithLabelStatusPending    BatchResponsePublicAssociationMultiWithLabelStatus = "PENDING"
	BatchResponsePublicAssociationMultiWithLabelStatusProcessing BatchResponsePublicAssociationMultiWithLabelStatus = "PROCESSING"
)

type DateTime struct {
	// Indicates whether the DateTime value represents only a date without a time
	// component.
	DateOnly bool `json:"dateOnly" api:"required"`
	// The integer value representing the shift in minutes from UTC for the DateTime
	// value.
	TimeZoneShift int64 `json:"timeZoneShift" api:"required"`
	// The integer value representing a specific point in time.
	Value int64 `json:"value" api:"required"`
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
	// Contains the Id of a Public Object
	From shared.PublicObjectIDParam   `json:"from,omitzero" api:"required"`
	To   []shared.PublicObjectIDParam `json:"to,omitzero" api:"required"`
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
	// Contains the Id of a Public Object
	From shared.PublicObjectIDParam `json:"from,omitzero" api:"required"`
	// Contains the Id of a Public Object
	To    shared.PublicObjectIDParam    `json:"to,omitzero" api:"required"`
	Types []shared.AssociationSpecParam `json:"types,omitzero" api:"required"`
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
	// Contains the Id of a Public Object
	From   shared.PublicObjectID            `json:"from" api:"required"`
	To     []MultiAssociatedObjectWithLabel `json:"to" api:"required"`
	Paging shared.Paging                    `json:"paging"`
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
	// Contains the Id of a Public Object
	From shared.PublicObjectIDParam `json:"from,omitzero" api:"required"`
	// Contains the Id of a Public Object
	To shared.PublicObjectIDParam `json:"to,omitzero" api:"required"`
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
	// The unique identifier for the object whose associations are being fetched.
	ID string `json:"id" api:"required"`
	// A paging cursor token used to retrieve the next set of results in a paginated
	// response.
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
	EnqueueTime DateTime `json:"enqueueTime" api:"required"`
	// Email of the user
	UserEmail string `json:"userEmail" api:"required"`
	// ID of the user
	UserID int64 `json:"userId" api:"required"`
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

type AssociationListParams struct {
	ObjectType string `path:"objectType" api:"required" json:"-"`
	ObjectID   string `path:"objectId" api:"required" json:"-"`
	// The paging cursor token of the last successfully read resource will be returned
	// as the `paging.next.after` JSON property of a paged response containing more
	// results.
	After param.Opt[string] `query:"after,omitzero" json:"-"`
	// The maximum number of results to display per page.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [AssociationListParams]'s query parameters as `url.Values`.
func (r AssociationListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AssociationDeleteParams struct {
	ObjectType   string `path:"objectType" api:"required" json:"-"`
	ObjectID     string `path:"objectId" api:"required" json:"-"`
	ToObjectType string `path:"toObjectType" api:"required" json:"-"`
	paramObj
}

type AssociationSearchParams struct {
	// Describes a search request
	PublicObjectSearchRequest PublicObjectSearchRequestParam
	paramObj
}

func (r AssociationSearchParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.PublicObjectSearchRequest)
}
func (r *AssociationSearchParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AssociationUpdateAssociationLabelsParams struct {
	ObjectType   string `path:"objectType" api:"required" json:"-"`
	ObjectID     string `path:"objectId" api:"required" json:"-"`
	ToObjectType string `path:"toObjectType" api:"required" json:"-"`
	Body         []shared.AssociationSpecParam
	paramObj
}

func (r AssociationUpdateAssociationLabelsParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.Body)
}
func (r *AssociationUpdateAssociationLabelsParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
