// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package crm

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"
	"slices"

	"github.com/stainless-sdks/hubspot-sdk-go/internal/apiquery"
	shimjson "github.com/stainless-sdks/hubspot-sdk-go/internal/encoding/json"
	"github.com/stainless-sdks/hubspot-sdk-go/internal/requestconfig"
	"github.com/stainless-sdks/hubspot-sdk-go/option"
	"github.com/stainless-sdks/hubspot-sdk-go/packages/param"
)

// ObjectCartBatchService contains methods and other services that help with
// interacting with the Hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewObjectCartBatchService] method instead.
type ObjectCartBatchService struct {
	Options []option.RequestOption
}

// NewObjectCartBatchService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewObjectCartBatchService(opts ...option.RequestOption) (r ObjectCartBatchService) {
	r = ObjectCartBatchService{}
	r.Options = opts
	return
}

// Create a batch of carts
func (r *ObjectCartBatchService) New(ctx context.Context, body ObjectCartBatchNewParams, opts ...option.RequestOption) (res *BatchResponseSimplePublicObject, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "crm/v3/objects/carts/batch/create"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Update a batch of carts by internal ID, or unique property values
func (r *ObjectCartBatchService) Update(ctx context.Context, body ObjectCartBatchUpdateParams, opts ...option.RequestOption) (res *BatchResponseSimplePublicObject, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "crm/v3/objects/carts/batch/update"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Archive a batch of carts by ID
func (r *ObjectCartBatchService) Delete(ctx context.Context, body ObjectCartBatchDeleteParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "crm/v3/objects/carts/batch/archive"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Retrieve records by record ID or include the `idProperty` parameter to retrieve
// records by a custom unique value property.
func (r *ObjectCartBatchService) Get(ctx context.Context, params ObjectCartBatchGetParams, opts ...option.RequestOption) (res *BatchResponseSimplePublicObject, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "crm/v3/objects/carts/batch/read"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Create or update records identified by a unique property value as specified by
// the `idProperty` query param. `idProperty` query param refers to a property
// whose values are unique for the object.
func (r *ObjectCartBatchService) Upsert(ctx context.Context, body ObjectCartBatchUpsertParams, opts ...option.RequestOption) (res *BatchResponseSimplePublicUpsertObject, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "crm/v3/objects/carts/batch/upsert"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type ObjectCartBatchNewParams struct {
	BatchInputSimplePublicObjectBatchInputForCreate BatchInputSimplePublicObjectBatchInputForCreateParam
	paramObj
}

func (r ObjectCartBatchNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.BatchInputSimplePublicObjectBatchInputForCreate)
}
func (r *ObjectCartBatchNewParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.BatchInputSimplePublicObjectBatchInputForCreate)
}

type ObjectCartBatchUpdateParams struct {
	BatchInputSimplePublicObjectBatchInput BatchInputSimplePublicObjectBatchInputParam
	paramObj
}

func (r ObjectCartBatchUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.BatchInputSimplePublicObjectBatchInput)
}
func (r *ObjectCartBatchUpdateParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.BatchInputSimplePublicObjectBatchInput)
}

type ObjectCartBatchDeleteParams struct {
	BatchInputSimplePublicObjectID BatchInputSimplePublicObjectIDParam
	paramObj
}

func (r ObjectCartBatchDeleteParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.BatchInputSimplePublicObjectID)
}
func (r *ObjectCartBatchDeleteParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.BatchInputSimplePublicObjectID)
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
	return json.Unmarshal(data, &r.BatchReadInputSimplePublicObjectID)
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
	return json.Unmarshal(data, &r.BatchInputSimplePublicObjectBatchInputUpsert)
}
