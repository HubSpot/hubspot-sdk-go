// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package crm

import (
	"context"
	"net/http"
	"net/url"
	"slices"

	"github.com/stainless-sdks/hubspot-sdk-go/internal/apijson"
	"github.com/stainless-sdks/hubspot-sdk-go/internal/apiquery"
	shimjson "github.com/stainless-sdks/hubspot-sdk-go/internal/encoding/json"
	"github.com/stainless-sdks/hubspot-sdk-go/internal/requestconfig"
	"github.com/stainless-sdks/hubspot-sdk-go/option"
	"github.com/stainless-sdks/hubspot-sdk-go/packages/param"
)

// ObjectCommercePaymentBatchService contains methods and other services that help
// with interacting with the hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewObjectCommercePaymentBatchService] method instead.
type ObjectCommercePaymentBatchService struct {
	options []option.RequestOption
}

// NewObjectCommercePaymentBatchService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewObjectCommercePaymentBatchService(opts ...option.RequestOption) (r ObjectCommercePaymentBatchService) {
	r = ObjectCommercePaymentBatchService{}
	r.options = opts
	return
}

// Create multiple commerce payment records in a single request, returning the
// details of each created payment, including their unique IDs.
func (r *ObjectCommercePaymentBatchService) New(ctx context.Context, body ObjectCommercePaymentBatchNewParams, opts ...option.RequestOption) (res *BatchResponseSimplePublicObject, err error) {
	opts = slices.Concat(r.options, opts)
	path := "crm/objects/2026-03/commerce_payments/batch/create"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Update multiple commerce payment records using their internal IDs or unique
// property values. This operation allows you to modify existing payment records in
// bulk by providing a list of records with their respective IDs and updated
// property values.
func (r *ObjectCommercePaymentBatchService) Update(ctx context.Context, body ObjectCommercePaymentBatchUpdateParams, opts ...option.RequestOption) (res *BatchResponseSimplePublicObject, err error) {
	opts = slices.Concat(r.options, opts)
	path := "crm/objects/2026-03/commerce_payments/batch/update"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Archive a batch of commerce payments by their IDs. This operation moves the
// specified payments to the archive, making them inactive in the system.
func (r *ObjectCommercePaymentBatchService) Delete(ctx context.Context, body ObjectCommercePaymentBatchDeleteParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "crm/objects/2026-03/commerce_payments/batch/archive"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return err
}

// Retrieve records by record ID or include the `idProperty` parameter to retrieve
// records by a custom unique value property.
func (r *ObjectCommercePaymentBatchService) Get(ctx context.Context, params ObjectCommercePaymentBatchGetParams, opts ...option.RequestOption) (res *BatchResponseSimplePublicObject, err error) {
	opts = slices.Concat(r.options, opts)
	path := "crm/objects/2026-03/commerce_payments/batch/read"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Create or update records identified by a unique property value as specified by
// the `idProperty` query param. `idProperty` query param refers to a property
// whose values are unique for the object.
func (r *ObjectCommercePaymentBatchService) Upsert(ctx context.Context, body ObjectCommercePaymentBatchUpsertParams, opts ...option.RequestOption) (res *BatchResponseSimplePublicUpsertObject, err error) {
	opts = slices.Concat(r.options, opts)
	path := "crm/objects/2026-03/commerce_payments/batch/upsert"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type ObjectCommercePaymentBatchNewParams struct {
	BatchInputSimplePublicObjectBatchInputForCreate BatchInputSimplePublicObjectBatchInputForCreateParam
	paramObj
}

func (r ObjectCommercePaymentBatchNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.BatchInputSimplePublicObjectBatchInputForCreate)
}
func (r *ObjectCommercePaymentBatchNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ObjectCommercePaymentBatchUpdateParams struct {
	BatchInputSimplePublicObjectBatchInput BatchInputSimplePublicObjectBatchInputParam
	paramObj
}

func (r ObjectCommercePaymentBatchUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.BatchInputSimplePublicObjectBatchInput)
}
func (r *ObjectCommercePaymentBatchUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ObjectCommercePaymentBatchDeleteParams struct {
	BatchInputSimplePublicObjectID BatchInputSimplePublicObjectIDParam
	paramObj
}

func (r ObjectCommercePaymentBatchDeleteParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.BatchInputSimplePublicObjectID)
}
func (r *ObjectCommercePaymentBatchDeleteParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ObjectCommercePaymentBatchGetParams struct {
	// Specifies the input for reading a batch of CRM objects, including arrays of
	// object IDs, requested property names (with optional history), and an optional
	// unique identifying property.
	BatchReadInputSimplePublicObjectID BatchReadInputSimplePublicObjectIDParam
	// Whether to return only results that have been archived.
	Archived param.Opt[bool] `query:"archived,omitzero" json:"-"`
	paramObj
}

func (r ObjectCommercePaymentBatchGetParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.BatchReadInputSimplePublicObjectID)
}
func (r *ObjectCommercePaymentBatchGetParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [ObjectCommercePaymentBatchGetParams]'s query parameters as
// `url.Values`.
func (r ObjectCommercePaymentBatchGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ObjectCommercePaymentBatchUpsertParams struct {
	BatchInputSimplePublicObjectBatchInputUpsert BatchInputSimplePublicObjectBatchInputUpsertParam
	paramObj
}

func (r ObjectCommercePaymentBatchUpsertParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.BatchInputSimplePublicObjectBatchInputUpsert)
}
func (r *ObjectCommercePaymentBatchUpsertParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
