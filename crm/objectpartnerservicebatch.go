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

// ObjectPartnerServiceBatchService contains methods and other services that help
// with interacting with the hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewObjectPartnerServiceBatchService] method instead.
type ObjectPartnerServiceBatchService struct {
	options []option.RequestOption
}

// NewObjectPartnerServiceBatchService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewObjectPartnerServiceBatchService(opts ...option.RequestOption) (r ObjectPartnerServiceBatchService) {
	r = ObjectPartnerServiceBatchService{}
	r.options = opts
	return
}

// Update multiple partner services using their internal IDs or unique property
// values. This operation allows for batch processing of updates, ensuring
// efficient synchronization of service data between HubSpot and other systems.
func (r *ObjectPartnerServiceBatchService) Update(ctx context.Context, body ObjectPartnerServiceBatchUpdateParams, opts ...option.RequestOption) (res *BatchResponseSimplePublicObject, err error) {
	opts = slices.Concat(r.options, opts)
	path := "crm/objects/2026-03/partner_services/batch/update"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Retrieve records by record ID or include the `idProperty` parameter to retrieve
// records by a custom unique value property.
func (r *ObjectPartnerServiceBatchService) Get(ctx context.Context, params ObjectPartnerServiceBatchGetParams, opts ...option.RequestOption) (res *BatchResponseSimplePublicObject, err error) {
	opts = slices.Concat(r.options, opts)
	path := "crm/objects/2026-03/partner_services/batch/read"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

type ObjectPartnerServiceBatchUpdateParams struct {
	BatchInputSimplePublicObjectBatchInput BatchInputSimplePublicObjectBatchInputParam
	paramObj
}

func (r ObjectPartnerServiceBatchUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.BatchInputSimplePublicObjectBatchInput)
}
func (r *ObjectPartnerServiceBatchUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ObjectPartnerServiceBatchGetParams struct {
	// Specifies the input for reading a batch of CRM objects, including arrays of
	// object IDs, requested property names (with optional history), and an optional
	// unique identifying property.
	BatchReadInputSimplePublicObjectID BatchReadInputSimplePublicObjectIDParam
	// Whether to return only results that have been archived.
	Archived param.Opt[bool] `query:"archived,omitzero" json:"-"`
	paramObj
}

func (r ObjectPartnerServiceBatchGetParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.BatchReadInputSimplePublicObjectID)
}
func (r *ObjectPartnerServiceBatchGetParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [ObjectPartnerServiceBatchGetParams]'s query parameters as
// `url.Values`.
func (r ObjectPartnerServiceBatchGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
