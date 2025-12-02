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

	"github.com/stainless-sdks/hubspot-sdk-go/internal/apiquery"
	shimjson "github.com/stainless-sdks/hubspot-sdk-go/internal/encoding/json"
	"github.com/stainless-sdks/hubspot-sdk-go/internal/requestconfig"
	"github.com/stainless-sdks/hubspot-sdk-go/option"
	"github.com/stainless-sdks/hubspot-sdk-go/packages/pagination"
	"github.com/stainless-sdks/hubspot-sdk-go/packages/param"
)

// ObjectPartnerClientService contains methods and other services that help with
// interacting with the hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewObjectPartnerClientService] method instead.
type ObjectPartnerClientService struct {
	Options      []option.RequestOption
	Associations ObjectPartnerClientAssociationService
	Batch        ObjectPartnerClientBatchService
}

// NewObjectPartnerClientService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewObjectPartnerClientService(opts ...option.RequestOption) (r ObjectPartnerClientService) {
	r = ObjectPartnerClientService{}
	r.Options = opts
	r.Associations = NewObjectPartnerClientAssociationService(opts...)
	r.Batch = NewObjectPartnerClientBatchService(opts...)
	return
}

func (r *ObjectPartnerClientService) Update(ctx context.Context, partnerClientID string, params ObjectPartnerClientUpdateParams, opts ...option.RequestOption) (res *SimplePublicObject, err error) {
	opts = slices.Concat(r.Options, opts)
	if partnerClientID == "" {
		err = errors.New("missing required partnerClientId parameter")
		return
	}
	path := fmt.Sprintf("crm/v3/objects/partner_clients/%s", partnerClientID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return
}

func (r *ObjectPartnerClientService) List(ctx context.Context, query ObjectPartnerClientListParams, opts ...option.RequestOption) (res *pagination.Page[SimplePublicObjectWithAssociations], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "crm/v3/objects/partner_clients"
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
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

func (r *ObjectPartnerClientService) ListAutoPaging(ctx context.Context, query ObjectPartnerClientListParams, opts ...option.RequestOption) *pagination.PageAutoPager[SimplePublicObjectWithAssociations] {
	return pagination.NewPageAutoPager(r.List(ctx, query, opts...))
}

func (r *ObjectPartnerClientService) Get(ctx context.Context, partnerClientID string, query ObjectPartnerClientGetParams, opts ...option.RequestOption) (res *SimplePublicObjectWithAssociations, err error) {
	opts = slices.Concat(r.Options, opts)
	if partnerClientID == "" {
		err = errors.New("missing required partnerClientId parameter")
		return
	}
	path := fmt.Sprintf("crm/v3/objects/partner_clients/%s", partnerClientID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

func (r *ObjectPartnerClientService) Search(ctx context.Context, body ObjectPartnerClientSearchParams, opts ...option.RequestOption) (res *CollectionResponseWithTotalSimplePublicObject, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "crm/v3/objects/partner_clients/search"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type ObjectPartnerClientUpdateParams struct {
	// Represents the input required to create or update a CRM object, containing an
	// object with property names and their corresponding values.
	SimplePublicObjectInput SimplePublicObjectInputParam
	IDProperty              param.Opt[string] `query:"idProperty,omitzero" json:"-"`
	paramObj
}

func (r ObjectPartnerClientUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.SimplePublicObjectInput)
}
func (r *ObjectPartnerClientUpdateParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.SimplePublicObjectInput)
}

// URLQuery serializes [ObjectPartnerClientUpdateParams]'s query parameters as
// `url.Values`.
func (r ObjectPartnerClientUpdateParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ObjectPartnerClientListParams struct {
	After                 param.Opt[string] `query:"after,omitzero" json:"-"`
	Archived              param.Opt[bool]   `query:"archived,omitzero" json:"-"`
	Limit                 param.Opt[int64]  `query:"limit,omitzero" json:"-"`
	Associations          []string          `query:"associations,omitzero" json:"-"`
	Properties            []string          `query:"properties,omitzero" json:"-"`
	PropertiesWithHistory []string          `query:"propertiesWithHistory,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ObjectPartnerClientListParams]'s query parameters as
// `url.Values`.
func (r ObjectPartnerClientListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ObjectPartnerClientGetParams struct {
	Archived              param.Opt[bool]   `query:"archived,omitzero" json:"-"`
	IDProperty            param.Opt[string] `query:"idProperty,omitzero" json:"-"`
	Associations          []string          `query:"associations,omitzero" json:"-"`
	Properties            []string          `query:"properties,omitzero" json:"-"`
	PropertiesWithHistory []string          `query:"propertiesWithHistory,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ObjectPartnerClientGetParams]'s query parameters as
// `url.Values`.
func (r ObjectPartnerClientGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ObjectPartnerClientSearchParams struct {
	// Describes a search request
	PublicObjectSearchRequest PublicObjectSearchRequestParam
	paramObj
}

func (r ObjectPartnerClientSearchParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.PublicObjectSearchRequest)
}
func (r *ObjectPartnerClientSearchParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.PublicObjectSearchRequest)
}
