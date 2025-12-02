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

// CurrencyCentralFxRateService contains methods and other services that help with
// interacting with the hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCurrencyCentralFxRateService] method instead.
type CurrencyCentralFxRateService struct {
	Options []option.RequestOption
}

// NewCurrencyCentralFxRateService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewCurrencyCentralFxRateService(opts ...option.RequestOption) (r CurrencyCentralFxRateService) {
	r = CurrencyCentralFxRateService{}
	r.Options = opts
	return
}

// Create a new currency with central exchange rates in the portal. Unsupported
// currencies cannot be added here.
func (r *CurrencyCentralFxRateService) NewCurrency(ctx context.Context, body CurrencyCentralFxRateNewCurrencyParams, opts ...option.RequestOption) (res *ExchangeRate, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "settings/v3/currencies/central-fx-rates/add-currency"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieve details on whether the central exchange rates feature is enabled for
// the portal.
func (r *CurrencyCentralFxRateService) GetInformation(ctx context.Context, opts ...option.RequestOption) (res *CentralExchangeRatesInformation, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "settings/v3/currencies/central-fx-rates/information"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Retrieve a list of currency codes that are not supported by the central exchange
// rates. Unsupported currencies will need to be manually updated.
func (r *CurrencyCentralFxRateService) GetUnsupportedCurrencies(ctx context.Context, opts ...option.RequestOption) (res *CollectionResponseCurrencyCodeInfoNoPaging, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "settings/v3/currencies/central-fx-rates/unsupported-currencies"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type CurrencyCentralFxRateNewCurrencyParams struct {
	CurrencyCreateRequest CurrencyCreateRequestParam
	paramObj
}

func (r CurrencyCentralFxRateNewCurrencyParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.CurrencyCreateRequest)
}
func (r *CurrencyCentralFxRateNewCurrencyParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.CurrencyCreateRequest)
}
