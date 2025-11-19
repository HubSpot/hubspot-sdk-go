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

// ObjectCommunicationBatchService contains methods and other services that help
// with interacting with the Hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewObjectCommunicationBatchService] method instead.
type ObjectCommunicationBatchService struct {
	Options []option.RequestOption
}

// NewObjectCommunicationBatchService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewObjectCommunicationBatchService(opts ...option.RequestOption) (r ObjectCommunicationBatchService) {
	r = ObjectCommunicationBatchService{}
	r.Options = opts
	return
}

// Create a batch of messages. The `inputs` array can contain a `properties` object
// to define property values for each message, along with an `associations` array
// to define
// [associations](https://developers.hubspot.com/docs/guides/api/crm/associations/associations-v4)
// with other CRM records.
func (r *ObjectCommunicationBatchService) New(ctx context.Context, body ObjectCommunicationBatchNewParams, opts ...option.RequestOption) (res *BatchResponseSimplePublicObject, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "crm/v3/objects/communications/batch/create"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Update a batch of messages by ID (`communicationId`) or unique property value
// (`idProperty`). Provided property values will be overwritten. Read-only and
// non-existent properties will result in an error. Properties values can be
// cleared by passing an empty string.
func (r *ObjectCommunicationBatchService) Update(ctx context.Context, body ObjectCommunicationBatchUpdateParams, opts ...option.RequestOption) (res *BatchResponseSimplePublicObject, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "crm/v3/objects/communications/batch/update"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Delete a batch of messages by ID. A deleted message can be restored within 90
// days of being deleted. Learn more about
// [restoring activity records](https://knowledge.hubspot.com/records/restore-deleted-activity-in-a-record).
func (r *ObjectCommunicationBatchService) Delete(ctx context.Context, body ObjectCommunicationBatchDeleteParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "crm/v3/objects/communications/batch/archive"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Retrieve a batch of messages by ID (`communicationId`) or unique property value
// (`idProperty`).
func (r *ObjectCommunicationBatchService) Get(ctx context.Context, params ObjectCommunicationBatchGetParams, opts ...option.RequestOption) (res *BatchResponseSimplePublicObject, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "crm/v3/objects/communications/batch/read"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Create or update records identified by a unique property value as specified by
// the `idProperty` query param. `idProperty` query param refers to a property
// whose values are unique for the object.
func (r *ObjectCommunicationBatchService) Upsert(ctx context.Context, body ObjectCommunicationBatchUpsertParams, opts ...option.RequestOption) (res *BatchResponseSimplePublicUpsertObject, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "crm/v3/objects/communications/batch/upsert"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type ObjectCommunicationBatchNewParams struct {
	BatchInputSimplePublicObjectBatchInputForCreate BatchInputSimplePublicObjectBatchInputForCreateParam
	paramObj
}

func (r ObjectCommunicationBatchNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.BatchInputSimplePublicObjectBatchInputForCreate)
}
func (r *ObjectCommunicationBatchNewParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.BatchInputSimplePublicObjectBatchInputForCreate)
}

type ObjectCommunicationBatchUpdateParams struct {
	BatchInputSimplePublicObjectBatchInput BatchInputSimplePublicObjectBatchInputParam
	paramObj
}

func (r ObjectCommunicationBatchUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.BatchInputSimplePublicObjectBatchInput)
}
func (r *ObjectCommunicationBatchUpdateParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.BatchInputSimplePublicObjectBatchInput)
}

type ObjectCommunicationBatchDeleteParams struct {
	BatchInputSimplePublicObjectID BatchInputSimplePublicObjectIDParam
	paramObj
}

func (r ObjectCommunicationBatchDeleteParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.BatchInputSimplePublicObjectID)
}
func (r *ObjectCommunicationBatchDeleteParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.BatchInputSimplePublicObjectID)
}

type ObjectCommunicationBatchGetParams struct {
	// Specifies the input for reading a batch of CRM objects, including arrays of
	// object IDs, requested property names (with optional history), and an optional
	// unique identifying property.
	BatchReadInputSimplePublicObjectID BatchReadInputSimplePublicObjectIDParam
	// Whether to return only results that have been archived.
	Archived param.Opt[bool] `query:"archived,omitzero" json:"-"`
	paramObj
}

func (r ObjectCommunicationBatchGetParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.BatchReadInputSimplePublicObjectID)
}
func (r *ObjectCommunicationBatchGetParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.BatchReadInputSimplePublicObjectID)
}

// URLQuery serializes [ObjectCommunicationBatchGetParams]'s query parameters as
// `url.Values`.
func (r ObjectCommunicationBatchGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ObjectCommunicationBatchUpsertParams struct {
	BatchInputSimplePublicObjectBatchInputUpsert BatchInputSimplePublicObjectBatchInputUpsertParam
	paramObj
}

func (r ObjectCommunicationBatchUpsertParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.BatchInputSimplePublicObjectBatchInputUpsert)
}
func (r *ObjectCommunicationBatchUpsertParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.BatchInputSimplePublicObjectBatchInputUpsert)
}
