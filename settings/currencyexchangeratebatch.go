// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package settings

import (
	"context"
	"encoding/json"
	"net/http"
	"slices"

	shimjson "github.com/stainless-sdks/hubspot-sdk-go/internal/encoding/json"
	"github.com/stainless-sdks/hubspot-sdk-go/internal/requestconfig"
	"github.com/stainless-sdks/hubspot-sdk-go/option"
)

// CurrencyExchangeRateBatchService contains methods and other services that help
// with interacting with the hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCurrencyExchangeRateBatchService] method instead.
type CurrencyExchangeRateBatchService struct {
	Options []option.RequestOption
}

// NewCurrencyExchangeRateBatchService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewCurrencyExchangeRateBatchService(opts ...option.RequestOption) (r CurrencyExchangeRateBatchService) {
	r = CurrencyExchangeRateBatchService{}
	r.Options = opts
	return
}

func (r *CurrencyExchangeRateBatchService) New(ctx context.Context, body CurrencyExchangeRateBatchNewParams, opts ...option.RequestOption) (res *BatchResponseExchangeRate, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "settings/currencies/2026-03/exchange-rates/batch/create"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

func (r *CurrencyExchangeRateBatchService) Update(ctx context.Context, body CurrencyExchangeRateBatchUpdateParams, opts ...option.RequestOption) (res *BatchResponseExchangeRate, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "settings/currencies/2026-03/exchange-rates/batch/update"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

func (r *CurrencyExchangeRateBatchService) Get(ctx context.Context, body CurrencyExchangeRateBatchGetParams, opts ...option.RequestOption) (res *BatchResponseExchangeRate, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "settings/currencies/2026-03/exchange-rates/batch/read"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type CurrencyExchangeRateBatchNewParams struct {
	BatchInputExchangeRateCreateRequest BatchInputExchangeRateCreateRequestParam
	paramObj
}

func (r CurrencyExchangeRateBatchNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.BatchInputExchangeRateCreateRequest)
}
func (r *CurrencyExchangeRateBatchNewParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.BatchInputExchangeRateCreateRequest)
}

type CurrencyExchangeRateBatchUpdateParams struct {
	BatchInputExchangeRateUpdateRequest BatchInputExchangeRateUpdateRequestParam
	paramObj
}

func (r CurrencyExchangeRateBatchUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.BatchInputExchangeRateUpdateRequest)
}
func (r *CurrencyExchangeRateBatchUpdateParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.BatchInputExchangeRateUpdateRequest)
}

type CurrencyExchangeRateBatchGetParams struct {
	BatchInputPublicObjectID BatchInputPublicObjectIDParam
	paramObj
}

func (r CurrencyExchangeRateBatchGetParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.BatchInputPublicObjectID)
}
func (r *CurrencyExchangeRateBatchGetParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.BatchInputPublicObjectID)
}
