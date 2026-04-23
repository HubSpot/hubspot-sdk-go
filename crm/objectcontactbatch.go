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

// ObjectContactBatchService contains methods and other services that help with
// interacting with the hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewObjectContactBatchService] method instead.
type ObjectContactBatchService struct {
	options []option.RequestOption
}

// NewObjectContactBatchService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewObjectContactBatchService(opts ...option.RequestOption) (r ObjectContactBatchService) {
	r = ObjectContactBatchService{}
	r.options = opts
	return
}

func (r *ObjectContactBatchService) New(ctx context.Context, body ObjectContactBatchNewParams, opts ...option.RequestOption) (res *BatchResponseSimplePublicObject, err error) {
	opts = slices.Concat(r.options, opts)
	path := "crm/objects/2026-03/contacts/batch/create"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

func (r *ObjectContactBatchService) Update(ctx context.Context, body ObjectContactBatchUpdateParams, opts ...option.RequestOption) (res *BatchResponseSimplePublicObject, err error) {
	opts = slices.Concat(r.options, opts)
	path := "crm/objects/2026-03/contacts/batch/update"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Archive a batch of contacts
func (r *ObjectContactBatchService) Delete(ctx context.Context, body ObjectContactBatchDeleteParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "crm/objects/2026-03/contacts/batch/archive"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return err
}

func (r *ObjectContactBatchService) Get(ctx context.Context, params ObjectContactBatchGetParams, opts ...option.RequestOption) (res *BatchResponseSimplePublicObject, err error) {
	opts = slices.Concat(r.options, opts)
	path := "crm/objects/2026-03/contacts/batch/read"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

func (r *ObjectContactBatchService) Upsert(ctx context.Context, body ObjectContactBatchUpsertParams, opts ...option.RequestOption) (res *BatchResponseSimplePublicUpsertObject, err error) {
	opts = slices.Concat(r.options, opts)
	path := "crm/objects/2026-03/contacts/batch/upsert"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type ObjectContactBatchNewParams struct {
	BatchInputSimplePublicObjectBatchInputForCreate BatchInputSimplePublicObjectBatchInputForCreateParam
	paramObj
}

func (r ObjectContactBatchNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.BatchInputSimplePublicObjectBatchInputForCreate)
}
func (r *ObjectContactBatchNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ObjectContactBatchUpdateParams struct {
	BatchInputSimplePublicObjectBatchInput BatchInputSimplePublicObjectBatchInputParam
	paramObj
}

func (r ObjectContactBatchUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.BatchInputSimplePublicObjectBatchInput)
}
func (r *ObjectContactBatchUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ObjectContactBatchDeleteParams struct {
	BatchInputSimplePublicObjectID BatchInputSimplePublicObjectIDParam
	paramObj
}

func (r ObjectContactBatchDeleteParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.BatchInputSimplePublicObjectID)
}
func (r *ObjectContactBatchDeleteParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ObjectContactBatchGetParams struct {
	// Specifies the input for reading a batch of CRM objects, including arrays of
	// object IDs, requested property names (with optional history), and an optional
	// unique identifying property.
	BatchReadInputSimplePublicObjectID BatchReadInputSimplePublicObjectIDParam
	// Whether to return only results that have been archived.
	Archived param.Opt[bool] `query:"archived,omitzero" json:"-"`
	paramObj
}

func (r ObjectContactBatchGetParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.BatchReadInputSimplePublicObjectID)
}
func (r *ObjectContactBatchGetParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [ObjectContactBatchGetParams]'s query parameters as
// `url.Values`.
func (r ObjectContactBatchGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ObjectContactBatchUpsertParams struct {
	BatchInputSimplePublicObjectBatchInputUpsert BatchInputSimplePublicObjectBatchInputUpsertParam
	paramObj
}

func (r ObjectContactBatchUpsertParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.BatchInputSimplePublicObjectBatchInputUpsert)
}
func (r *ObjectContactBatchUpsertParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
