// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package crm

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/stainless-sdks/hubspot-sdk-go/internal/apijson"
	"github.com/stainless-sdks/hubspot-sdk-go/internal/apiquery"
	shimjson "github.com/stainless-sdks/hubspot-sdk-go/internal/encoding/json"
	"github.com/stainless-sdks/hubspot-sdk-go/internal/requestconfig"
	"github.com/stainless-sdks/hubspot-sdk-go/option"
	"github.com/stainless-sdks/hubspot-sdk-go/packages/pagination"
	"github.com/stainless-sdks/hubspot-sdk-go/packages/param"
)

// ObjectInvoiceService contains methods and other services that help with
// interacting with the hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewObjectInvoiceService] method instead.
type ObjectInvoiceService struct {
	options []option.RequestOption
	Batch   ObjectInvoiceBatchService
}

// NewObjectInvoiceService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewObjectInvoiceService(opts ...option.RequestOption) (r ObjectInvoiceService) {
	r = ObjectInvoiceService{}
	r.options = opts
	r.Batch = NewObjectInvoiceBatchService(opts...)
	return
}

// Create a invoice with the given properties and return a copy of the object,
// including the ID. Documentation and examples for creating standard invoices is
// provided.
func (r *ObjectInvoiceService) New(ctx context.Context, body ObjectInvoiceNewParams, opts ...option.RequestOption) (res *SimplePublicObject, err error) {
	opts = slices.Concat(r.options, opts)
	path := "crm/objects/2026-03/invoices"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Perform a partial update of an Object identified by `{invoiceId}`or optionally a
// unique property value as specified by the `idProperty` query param.
// `{invoiceId}` refers to the internal object ID by default, and the `idProperty`
// query param refers to a property whose values are unique for the object.
// Provided property values will be overwritten. Read-only and non-existent
// properties will result in an error. Properties values can be cleared by passing
// an empty string.
func (r *ObjectInvoiceService) Update(ctx context.Context, invoiceID string, params ObjectInvoiceUpdateParams, opts ...option.RequestOption) (res *SimplePublicObject, err error) {
	opts = slices.Concat(r.options, opts)
	if invoiceID == "" {
		err = errors.New("missing required invoiceId parameter")
		return nil, err
	}
	path := fmt.Sprintf("crm/objects/2026-03/invoices/%s", url.PathEscape(invoiceID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return res, err
}

// Read a page of invoices. Control what is returned via the `properties` query
// param.
func (r *ObjectInvoiceService) List(ctx context.Context, query ObjectInvoiceListParams, opts ...option.RequestOption) (res *pagination.Page[SimplePublicObjectWithAssociations], err error) {
	var raw *http.Response
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "crm/objects/2026-03/invoices"
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

// Read a page of invoices. Control what is returned via the `properties` query
// param.
func (r *ObjectInvoiceService) ListAutoPaging(ctx context.Context, query ObjectInvoiceListParams, opts ...option.RequestOption) *pagination.PageAutoPager[SimplePublicObjectWithAssociations] {
	return pagination.NewPageAutoPager(r.List(ctx, query, opts...))
}

// Move an Object identified by `{invoiceId}` to the recycling bin.
func (r *ObjectInvoiceService) Delete(ctx context.Context, invoiceID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if invoiceID == "" {
		err = errors.New("missing required invoiceId parameter")
		return err
	}
	path := fmt.Sprintf("crm/objects/2026-03/invoices/%s", url.PathEscape(invoiceID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Read an Object identified by `{invoiceId}`. `{invoiceId}` refers to the internal
// object ID by default, or optionally any unique property value as specified by
// the `idProperty` query param. Control what is returned via the `properties`
// query param.
func (r *ObjectInvoiceService) Get(ctx context.Context, invoiceID string, query ObjectInvoiceGetParams, opts ...option.RequestOption) (res *SimplePublicObjectWithAssociations, err error) {
	opts = slices.Concat(r.options, opts)
	if invoiceID == "" {
		err = errors.New("missing required invoiceId parameter")
		return nil, err
	}
	path := fmt.Sprintf("crm/objects/2026-03/invoices/%s", url.PathEscape(invoiceID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Execute a search for invoices based on filter criteria, sorting options, and
// properties to include in the response. This endpoint supports pagination and
// allows for complex queries using multiple filter groups.
func (r *ObjectInvoiceService) Search(ctx context.Context, body ObjectInvoiceSearchParams, opts ...option.RequestOption) (res *CollectionResponseWithTotalSimplePublicObject, err error) {
	opts = slices.Concat(r.options, opts)
	path := "crm/objects/2026-03/invoices/search"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type ObjectInvoiceNewParams struct {
	// Is the input object used to create a new CRM object, containing the properties
	// to be set and optional associations to link the new record with other CRM
	// objects.
	SimplePublicObjectInputForCreate SimplePublicObjectInputForCreateParam
	paramObj
}

func (r ObjectInvoiceNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.SimplePublicObjectInputForCreate)
}
func (r *ObjectInvoiceNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ObjectInvoiceUpdateParams struct {
	// Represents the input required to create or update a CRM object, containing an
	// object with property names and their corresponding values.
	SimplePublicObjectInput SimplePublicObjectInputParam
	// The name of a property whose values are unique for this object type
	IDProperty param.Opt[string] `query:"idProperty,omitzero" json:"-"`
	paramObj
}

func (r ObjectInvoiceUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.SimplePublicObjectInput)
}
func (r *ObjectInvoiceUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [ObjectInvoiceUpdateParams]'s query parameters as
// `url.Values`.
func (r ObjectInvoiceUpdateParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ObjectInvoiceListParams struct {
	// The paging cursor token of the last successfully read resource will be returned
	// as the `paging.next.after` JSON property of a paged response containing more
	// results.
	After param.Opt[string] `query:"after,omitzero" json:"-"`
	// Whether to return only results that have been archived.
	Archived param.Opt[bool] `query:"archived,omitzero" json:"-"`
	// The maximum number of results to display per page.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// A comma separated list of object types to retrieve associated IDs for. If any of
	// the specified associations do not exist, they will be ignored.
	Associations []string `query:"associations,omitzero" json:"-"`
	// A comma separated list of the properties to be returned in the response. If any
	// of the specified properties are not present on the requested object(s), they
	// will be ignored.
	Properties []string `query:"properties,omitzero" json:"-"`
	// A comma separated list of the properties to be returned along with their history
	// of previous values. If any of the specified properties are not present on the
	// requested object(s), they will be ignored. Usage of this parameter will reduce
	// the maximum number of objects that can be read by a single request.
	PropertiesWithHistory []string `query:"propertiesWithHistory,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ObjectInvoiceListParams]'s query parameters as
// `url.Values`.
func (r ObjectInvoiceListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ObjectInvoiceGetParams struct {
	// Whether to return only results that have been archived.
	Archived param.Opt[bool] `query:"archived,omitzero" json:"-"`
	// The name of a property whose values are unique for this object type
	IDProperty param.Opt[string] `query:"idProperty,omitzero" json:"-"`
	// A comma separated list of object types to retrieve associated IDs for. If any of
	// the specified associations do not exist, they will be ignored.
	Associations []string `query:"associations,omitzero" json:"-"`
	// A comma separated list of the properties to be returned in the response. If any
	// of the specified properties are not present on the requested object(s), they
	// will be ignored.
	Properties []string `query:"properties,omitzero" json:"-"`
	// A comma separated list of the properties to be returned along with their history
	// of previous values. If any of the specified properties are not present on the
	// requested object(s), they will be ignored.
	PropertiesWithHistory []string `query:"propertiesWithHistory,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ObjectInvoiceGetParams]'s query parameters as `url.Values`.
func (r ObjectInvoiceGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ObjectInvoiceSearchParams struct {
	// Describes a search request
	PublicObjectSearchRequest PublicObjectSearchRequestParam
	paramObj
}

func (r ObjectInvoiceSearchParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.PublicObjectSearchRequest)
}
func (r *ObjectInvoiceSearchParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
