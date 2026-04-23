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

// ObjectProductBatchService contains methods and other services that help with
// interacting with the hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewObjectProductBatchService] method instead.
type ObjectProductBatchService struct {
	options []option.RequestOption
}

// NewObjectProductBatchService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewObjectProductBatchService(opts ...option.RequestOption) (r ObjectProductBatchService) {
	r = ObjectProductBatchService{}
	r.options = opts
	return
}

// Create multiple products in a single request by specifying their properties, and
// receive a response containing the details of the created products.
func (r *ObjectProductBatchService) New(ctx context.Context, body ObjectProductBatchNewParams, opts ...option.RequestOption) (res *BatchResponseSimplePublicObject, err error) {
	opts = slices.Concat(r.options, opts)
	path := "crm/objects/2026-03/products/batch/create"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Update multiple products in a single request using their internal IDs or unique
// property values. This batch operation allows for efficient modifications of
// product records by specifying the properties to be updated. Ensure that the
// provided property values are correct, as read-only and non-existent properties
// will result in an error.
func (r *ObjectProductBatchService) Update(ctx context.Context, body ObjectProductBatchUpdateParams, opts ...option.RequestOption) (res *BatchResponseSimplePublicObject, err error) {
	opts = slices.Concat(r.options, opts)
	path := "crm/objects/2026-03/products/batch/update"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Archive multiple products at once by providing their IDs. This operation moves
// the specified products to the recycling bin, effectively removing them from
// active use without permanently deleting them.
func (r *ObjectProductBatchService) Delete(ctx context.Context, body ObjectProductBatchDeleteParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "crm/objects/2026-03/products/batch/archive"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return err
}

// Retrieve records by record ID or include the `idProperty` parameter to retrieve
// records by a custom unique value property.
func (r *ObjectProductBatchService) Get(ctx context.Context, params ObjectProductBatchGetParams, opts ...option.RequestOption) (res *BatchResponseSimplePublicObject, err error) {
	opts = slices.Concat(r.options, opts)
	path := "crm/objects/2026-03/products/batch/read"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Create or update records identified by a unique property value as specified by
// the `idProperty` query param. `idProperty` query param refers to a property
// whose values are unique for the object.
func (r *ObjectProductBatchService) Upsert(ctx context.Context, body ObjectProductBatchUpsertParams, opts ...option.RequestOption) (res *BatchResponseSimplePublicUpsertObject, err error) {
	opts = slices.Concat(r.options, opts)
	path := "crm/objects/2026-03/products/batch/upsert"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type ObjectProductBatchNewParams struct {
	BatchInputSimplePublicObjectBatchInputForCreate BatchInputSimplePublicObjectBatchInputForCreateParam
	paramObj
}

func (r ObjectProductBatchNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.BatchInputSimplePublicObjectBatchInputForCreate)
}
func (r *ObjectProductBatchNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ObjectProductBatchUpdateParams struct {
	BatchInputSimplePublicObjectBatchInput BatchInputSimplePublicObjectBatchInputParam
	paramObj
}

func (r ObjectProductBatchUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.BatchInputSimplePublicObjectBatchInput)
}
func (r *ObjectProductBatchUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ObjectProductBatchDeleteParams struct {
	BatchInputSimplePublicObjectID BatchInputSimplePublicObjectIDParam
	paramObj
}

func (r ObjectProductBatchDeleteParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.BatchInputSimplePublicObjectID)
}
func (r *ObjectProductBatchDeleteParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ObjectProductBatchGetParams struct {
	// Specifies the input for reading a batch of CRM objects, including arrays of
	// object IDs, requested property names (with optional history), and an optional
	// unique identifying property.
	BatchReadInputSimplePublicObjectID BatchReadInputSimplePublicObjectIDParam
	// Whether to return only results that have been archived.
	Archived param.Opt[bool] `query:"archived,omitzero" json:"-"`
	paramObj
}

func (r ObjectProductBatchGetParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.BatchReadInputSimplePublicObjectID)
}
func (r *ObjectProductBatchGetParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [ObjectProductBatchGetParams]'s query parameters as
// `url.Values`.
func (r ObjectProductBatchGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ObjectProductBatchUpsertParams struct {
	BatchInputSimplePublicObjectBatchInputUpsert BatchInputSimplePublicObjectBatchInputUpsertParam
	paramObj
}

func (r ObjectProductBatchUpsertParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.BatchInputSimplePublicObjectBatchInputUpsert)
}
func (r *ObjectProductBatchUpsertParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
