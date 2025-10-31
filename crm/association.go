// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package crm

import (
	"encoding/json"
	"time"

	"github.com/stainless-sdks/hubspot-sdk-go/internal/apijson"
	"github.com/stainless-sdks/hubspot-sdk-go/marketing"
	"github.com/stainless-sdks/hubspot-sdk-go/option"
	"github.com/stainless-sdks/hubspot-sdk-go/packages/param"
	"github.com/stainless-sdks/hubspot-sdk-go/packages/respjson"
	"github.com/stainless-sdks/hubspot-sdk-go/shared"
)

// AssociationService contains methods and other services that help with
// interacting with the Hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAssociationService] method instead.
type AssociationService struct {
	Options []option.RequestOption
	Batch   AssociationBatchService
	Schema  AssociationSchemaService
	V4      AssociationV4Service
}

// NewAssociationService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAssociationService(opts ...option.RequestOption) (r AssociationService) {
	r = AssociationService{}
	r.Options = opts
	r.Batch = NewAssociationBatchService(opts...)
	r.Schema = NewAssociationSchemaService(opts...)
	r.V4 = NewAssociationV4Service(opts...)
	return
}

// The property Inputs is required.
type BatchInputPublicAssociationParam struct {
	Inputs []PublicAssociationParam `json:"inputs,omitzero,required"`
	paramObj
}

func (r BatchInputPublicAssociationParam) MarshalJSON() (data []byte, err error) {
	type shadow BatchInputPublicAssociationParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BatchInputPublicAssociationParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BatchResponsePublicAssociation struct {
	CompletedAt time.Time           `json:"completedAt,required" format:"date-time"`
	Results     []PublicAssociation `json:"results,required"`
	StartedAt   time.Time           `json:"startedAt,required" format:"date-time"`
	// Any of "PENDING", "PROCESSING", "CANCELED", "COMPLETE".
	Status      BatchResponsePublicAssociationStatus `json:"status,required"`
	Errors      []shared.StandardError               `json:"errors"`
	Links       map[string]string                    `json:"links"`
	NumErrors   int64                                `json:"numErrors"`
	RequestedAt time.Time                            `json:"requestedAt" format:"date-time"`
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
func (r BatchResponsePublicAssociation) RawJSON() string { return r.JSON.raw }
func (r *BatchResponsePublicAssociation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BatchResponsePublicAssociationStatus string

const (
	BatchResponsePublicAssociationStatusPending    BatchResponsePublicAssociationStatus = "PENDING"
	BatchResponsePublicAssociationStatusProcessing BatchResponsePublicAssociationStatus = "PROCESSING"
	BatchResponsePublicAssociationStatusCanceled   BatchResponsePublicAssociationStatus = "CANCELED"
	BatchResponsePublicAssociationStatusComplete   BatchResponsePublicAssociationStatus = "COMPLETE"
)

type BatchResponsePublicAssociationMulti struct {
	CompletedAt time.Time                `json:"completedAt,required" format:"date-time"`
	Results     []PublicAssociationMulti `json:"results,required"`
	StartedAt   time.Time                `json:"startedAt,required" format:"date-time"`
	// Any of "PENDING", "PROCESSING", "CANCELED", "COMPLETE".
	Status      BatchResponsePublicAssociationMultiStatus `json:"status,required"`
	Errors      []shared.StandardError                    `json:"errors"`
	Links       map[string]string                         `json:"links"`
	NumErrors   int64                                     `json:"numErrors"`
	RequestedAt time.Time                                 `json:"requestedAt" format:"date-time"`
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
func (r BatchResponsePublicAssociationMulti) RawJSON() string { return r.JSON.raw }
func (r *BatchResponsePublicAssociationMulti) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BatchResponsePublicAssociationMultiStatus string

const (
	BatchResponsePublicAssociationMultiStatusPending    BatchResponsePublicAssociationMultiStatus = "PENDING"
	BatchResponsePublicAssociationMultiStatusProcessing BatchResponsePublicAssociationMultiStatus = "PROCESSING"
	BatchResponsePublicAssociationMultiStatusCanceled   BatchResponsePublicAssociationMultiStatus = "CANCELED"
	BatchResponsePublicAssociationMultiStatusComplete   BatchResponsePublicAssociationMultiStatus = "COMPLETE"
)

type PublicAssociation struct {
	From shared.PublicObjectID `json:"from,required"`
	To   shared.PublicObjectID `json:"to,required"`
	Type string                `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		From        respjson.Field
		To          respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PublicAssociation) RawJSON() string { return r.JSON.raw }
func (r *PublicAssociation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this PublicAssociation to a PublicAssociationParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// PublicAssociationParam.Overrides()
func (r PublicAssociation) ToParam() PublicAssociationParam {
	return param.Override[PublicAssociationParam](json.RawMessage(r.RawJSON()))
}

// The properties From, To, Type are required.
type PublicAssociationParam struct {
	From shared.PublicObjectIDParam `json:"from,omitzero,required"`
	To   shared.PublicObjectIDParam `json:"to,omitzero,required"`
	Type string                     `json:"type,required"`
	paramObj
}

func (r PublicAssociationParam) MarshalJSON() (data []byte, err error) {
	type shadow PublicAssociationParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PublicAssociationParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PublicAssociationMulti struct {
	From shared.PublicObjectID `json:"from,required"`
	// The IDs of objects that are associated with the object identified by the ID in
	// 'from'.
	To []AssociatedID `json:"to,required"`
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
func (r PublicAssociationMulti) RawJSON() string { return r.JSON.raw }
func (r *PublicAssociationMulti) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
