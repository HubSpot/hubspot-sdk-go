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

// ObjectCartBatchService contains methods and other services that help with
// interacting with the hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewObjectCartBatchService] method instead.
type ObjectCartBatchService struct {
	options []option.RequestOption
}

// NewObjectCartBatchService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewObjectCartBatchService(opts ...option.RequestOption) (r ObjectCartBatchService) {
	r = ObjectCartBatchService{}
	r.options = opts
	return
}

// Create a batch of carts with specified properties and associations.
func (r *ObjectCartBatchService) New(ctx context.Context, body ObjectCartBatchNewParams, opts ...option.RequestOption) (res *BatchResponseSimplePublicObject, err error) {
	opts = slices.Concat(r.options, opts)
	path := "crm/objects/2026-03/carts/batch/create"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Update a batch of carts using their internal IDs or unique property values.
func (r *ObjectCartBatchService) Update(ctx context.Context, body ObjectCartBatchUpdateParams, opts ...option.RequestOption) (res *BatchResponseSimplePublicObject, err error) {
	opts = slices.Concat(r.options, opts)
	path := "crm/objects/2026-03/carts/batch/update"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Archive a batch of carts identified by their IDs.
func (r *ObjectCartBatchService) Delete(ctx context.Context, body ObjectCartBatchDeleteParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "crm/objects/2026-03/carts/batch/archive"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return err
}

// Retrieve carts by ID, or include the `idProperty` parameter to retrieve carts by
// a custom unique value property.
func (r *ObjectCartBatchService) Get(ctx context.Context, params ObjectCartBatchGetParams, opts ...option.RequestOption) (res *BatchResponseSimplePublicObject, err error) {
	opts = slices.Concat(r.options, opts)
	path := "crm/objects/2026-03/carts/batch/read"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Create or update records identified by a unique property value as specified by
// the `idProperty` query parameter.
func (r *ObjectCartBatchService) Upsert(ctx context.Context, body ObjectCartBatchUpsertParams, opts ...option.RequestOption) (res *BatchResponseSimplePublicUpsertObject, err error) {
	opts = slices.Concat(r.options, opts)
	path := "crm/objects/2026-03/carts/batch/upsert"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type ObjectCartBatchNewParams struct {
	BatchInputSimplePublicObjectBatchInputForCreate BatchInputSimplePublicObjectBatchInputForCreateParam
	paramObj
}

func (r ObjectCartBatchNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.BatchInputSimplePublicObjectBatchInputForCreate)
}
func (r *ObjectCartBatchNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ObjectCartBatchUpdateParams struct {
	BatchInputSimplePublicObjectBatchInput BatchInputSimplePublicObjectBatchInputParam
	paramObj
}

func (r ObjectCartBatchUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.BatchInputSimplePublicObjectBatchInput)
}
func (r *ObjectCartBatchUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ObjectCartBatchDeleteParams struct {
	BatchInputSimplePublicObjectID BatchInputSimplePublicObjectIDParam
	paramObj
}

func (r ObjectCartBatchDeleteParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.BatchInputSimplePublicObjectID)
}
func (r *ObjectCartBatchDeleteParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ObjectCartBatchGetParams struct {
	// Specifies the input for reading a batch of CRM objects, including arrays of
	// object IDs, requested property names (with optional history), and an optional
	// unique identifying property.
	BatchReadInputSimplePublicObjectID BatchReadInputSimplePublicObjectIDParam
	// Whether to return only results that have been archived.
	Archived param.Opt[bool] `query:"archived,omitzero" json:"-"`
	paramObj
}

func (r ObjectCartBatchGetParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.BatchReadInputSimplePublicObjectID)
}
func (r *ObjectCartBatchGetParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [ObjectCartBatchGetParams]'s query parameters as
// `url.Values`.
func (r ObjectCartBatchGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ObjectCartBatchUpsertParams struct {
	BatchInputSimplePublicObjectBatchInputUpsert BatchInputSimplePublicObjectBatchInputUpsertParam
	paramObj
}

func (r ObjectCartBatchUpsertParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.BatchInputSimplePublicObjectBatchInputUpsert)
}
func (r *ObjectCartBatchUpsertParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
