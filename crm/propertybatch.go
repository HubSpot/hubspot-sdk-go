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
	"github.com/stainless-sdks/hubspot-sdk-go/packages/param"
	"github.com/stainless-sdks/hubspot-sdk-go/shared"
)

// PropertyBatchService contains methods and other services that help with
// interacting with the hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPropertyBatchService] method instead.
type PropertyBatchService struct {
	options []option.RequestOption
}

// NewPropertyBatchService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewPropertyBatchService(opts ...option.RequestOption) (r PropertyBatchService) {
	r = PropertyBatchService{}
	r.options = opts
	return
}

// Create a batch of properties using the same rules as when creating an individual
// property.
func (r *PropertyBatchService) New(ctx context.Context, objectType string, body PropertyBatchNewParams, opts ...option.RequestOption) (res *BatchResponseProperty, err error) {
	opts = slices.Concat(r.options, opts)
	if objectType == "" {
		err = errors.New("missing required objectType parameter")
		return nil, err
	}
	path := fmt.Sprintf("crm/properties/2026-03/%s/batch/create", url.PathEscape(objectType))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Archive a provided list of properties. This method will return a 204 No Content
// response on success regardless of the initial state of the property (e.g.
// active, already archived, non-existent).
func (r *PropertyBatchService) Delete(ctx context.Context, objectType string, body PropertyBatchDeleteParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if objectType == "" {
		err = errors.New("missing required objectType parameter")
		return err
	}
	path := fmt.Sprintf("crm/properties/2026-03/%s/batch/archive", url.PathEscape(objectType))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return err
}

// Read a provided list of properties.
func (r *PropertyBatchService) Get(ctx context.Context, objectType string, params PropertyBatchGetParams, opts ...option.RequestOption) (res *BatchResponseProperty, err error) {
	opts = slices.Concat(r.options, opts)
	if objectType == "" {
		err = errors.New("missing required objectType parameter")
		return nil, err
	}
	path := fmt.Sprintf("crm/properties/2026-03/%s/batch/read", url.PathEscape(objectType))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

type PropertyBatchNewParams struct {
	BatchInputPropertyCreate shared.BatchInputPropertyCreateParam
	paramObj
}

func (r PropertyBatchNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.BatchInputPropertyCreate)
}
func (r *PropertyBatchNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PropertyBatchDeleteParams struct {
	BatchInputPropertyName shared.BatchInputPropertyNameParam
	paramObj
}

func (r PropertyBatchDeleteParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.BatchInputPropertyName)
}
func (r *PropertyBatchDeleteParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PropertyBatchGetParams struct {
	BatchReadInputPropertyName shared.BatchReadInputPropertyNameParam
	Locale                     param.Opt[string] `query:"locale,omitzero" json:"-"`
	paramObj
}

func (r PropertyBatchGetParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.BatchReadInputPropertyName)
}
func (r *PropertyBatchGetParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [PropertyBatchGetParams]'s query parameters as `url.Values`.
func (r PropertyBatchGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
