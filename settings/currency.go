// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package settings

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"slices"
	"time"

	"github.com/stainless-sdks/hubspot-sdk-go/internal/apijson"
	shimjson "github.com/stainless-sdks/hubspot-sdk-go/internal/encoding/json"
	"github.com/stainless-sdks/hubspot-sdk-go/internal/requestconfig"
	"github.com/stainless-sdks/hubspot-sdk-go/option"
	"github.com/stainless-sdks/hubspot-sdk-go/packages/param"
	"github.com/stainless-sdks/hubspot-sdk-go/packages/respjson"
	"github.com/stainless-sdks/hubspot-sdk-go/shared"
)

// CurrencyService contains methods and other services that help with interacting
// with the Hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCurrencyService] method instead.
type CurrencyService struct {
	Options        []option.RequestOption
	CentralFxRates CurrencyCentralFxRateService
}

// NewCurrencyService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCurrencyService(opts ...option.RequestOption) (r CurrencyService) {
	r = CurrencyService{}
	r.Options = opts
	r.CentralFxRates = NewCurrencyCentralFxRateService(opts...)
	return
}

// Create multiple exchange rates in a single request.
func (r *CurrencyService) BatchNew(ctx context.Context, body CurrencyBatchNewParams, opts ...option.RequestOption) (res *BatchResponseExchangeRate, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "settings/v3/currencies/exchange-rates/batch/create"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieve the details of multiple exchange rates in a single request, specified
// by their IDs.
func (r *CurrencyService) BatchGet(ctx context.Context, body CurrencyBatchGetParams, opts ...option.RequestOption) (res *BatchResponseExchangeRate, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "settings/v3/currencies/exchange-rates/batch/read"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Update the conversion rates for multiple exchange rates in a batch operation.
func (r *CurrencyService) BatchUpdate(ctx context.Context, body CurrencyBatchUpdateParams, opts ...option.RequestOption) (res *BatchResponseExchangeRate, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "settings/v3/currencies/exchange-rates/batch/update"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Create a new exchange rate with specified conversion rate and currency codes.
func (r *CurrencyService) NewExchangeRate(ctx context.Context, body CurrencyNewExchangeRateParams, opts ...option.RequestOption) (res *ExchangeRate, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "settings/v3/currencies/exchange-rates"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get the details for the company currency. The company currency is used in deal
// totals, reports, and the default currency for new deals.
func (r *CurrencyService) GetCompanyCurrency(ctx context.Context, opts ...option.RequestOption) (res *CompanyCurrency, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "settings/v3/currencies/company-currency"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Retrieve the details for a specific exchange rate specified by its ID.
func (r *CurrencyService) GetExchangeRateByID(ctx context.Context, exchangeRateID string, opts ...option.RequestOption) (res *ExchangeRate, err error) {
	opts = slices.Concat(r.Options, opts)
	if exchangeRateID == "" {
		err = errors.New("missing required exchangeRateId parameter")
		return
	}
	path := fmt.Sprintf("settings/v3/currencies/exchange-rates/%s", exchangeRateID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Retrieve a list of all available currency codes and their names.
func (r *CurrencyService) ListCodes(ctx context.Context, opts ...option.RequestOption) (res *CollectionResponseCurrencyCodeInfoNoPaging, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "settings/v3/currencies/codes"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Retrieve all current exchange rates for all currency pairs.
func (r *CurrencyService) ListCurrentExchangeRates(ctx context.Context, opts ...option.RequestOption) (res *CollectionResponseExchangeRateNoPaging, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "settings/v3/currencies/exchange-rates/current"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Get a list of exchange rates
func (r *CurrencyService) ListExchangeRates(ctx context.Context, opts ...option.RequestOption) (res *CollectionResponseExchangeRateForwardPaging, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "settings/v3/currencies/exchange-rates"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Set or update the primary company currency.
func (r *CurrencyService) UpdateCompanyCurrency(ctx context.Context, body CurrencyUpdateCompanyCurrencyParams, opts ...option.RequestOption) (res *CompanyCurrency, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "settings/v3/currencies/company-currency"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Update an existing conversion rate, specified by its ID.
func (r *CurrencyService) UpdateExchangeRate(ctx context.Context, exchangeRateID string, body CurrencyUpdateExchangeRateParams, opts ...option.RequestOption) (res *ExchangeRate, err error) {
	opts = slices.Concat(r.Options, opts)
	if exchangeRateID == "" {
		err = errors.New("missing required exchangeRateId parameter")
		return
	}
	path := fmt.Sprintf("settings/v3/currencies/exchange-rates/%s", exchangeRateID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Change the visibility setting for a currency pair. This will hide or display a
// currency pair for users in the HubSpot app.
func (r *CurrencyService) UpdateVisibility(ctx context.Context, body CurrencyUpdateVisibilityParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "settings/v3/currencies/exchange-rates/update-visibility"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// The property Inputs is required.
type BatchInputExchangeRateCreateRequestParam struct {
	Inputs []ExchangeRateCreateRequestParam `json:"inputs,omitzero,required"`
	paramObj
}

func (r BatchInputExchangeRateCreateRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow BatchInputExchangeRateCreateRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BatchInputExchangeRateCreateRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Inputs is required.
type BatchInputExchangeRateUpdateRequestParam struct {
	Inputs []ExchangeRateUpdateRequestParam `json:"inputs,omitzero,required"`
	paramObj
}

func (r BatchInputExchangeRateUpdateRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow BatchInputExchangeRateUpdateRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BatchInputExchangeRateUpdateRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BatchResponseExchangeRate struct {
	CompletedAt time.Time      `json:"completedAt,required" format:"date-time"`
	Results     []ExchangeRate `json:"results,required"`
	StartedAt   time.Time      `json:"startedAt,required" format:"date-time"`
	// Any of "PENDING", "PROCESSING", "CANCELED", "COMPLETE".
	Status      BatchResponseExchangeRateStatus `json:"status,required"`
	Links       map[string]string               `json:"links"`
	RequestedAt time.Time                       `json:"requestedAt" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CompletedAt respjson.Field
		Results     respjson.Field
		StartedAt   respjson.Field
		Status      respjson.Field
		Links       respjson.Field
		RequestedAt respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BatchResponseExchangeRate) RawJSON() string { return r.JSON.raw }
func (r *BatchResponseExchangeRate) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BatchResponseExchangeRateStatus string

const (
	BatchResponseExchangeRateStatusPending    BatchResponseExchangeRateStatus = "PENDING"
	BatchResponseExchangeRateStatusProcessing BatchResponseExchangeRateStatus = "PROCESSING"
	BatchResponseExchangeRateStatusCanceled   BatchResponseExchangeRateStatus = "CANCELED"
	BatchResponseExchangeRateStatusComplete   BatchResponseExchangeRateStatus = "COMPLETE"
)

type CentralExchangeRatesInformation struct {
	CentralExchangeRatesEnabled bool `json:"centralExchangeRatesEnabled,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CentralExchangeRatesEnabled respjson.Field
		ExtraFields                 map[string]respjson.Field
		raw                         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CentralExchangeRatesInformation) RawJSON() string { return r.JSON.raw }
func (r *CentralExchangeRatesInformation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CollectionResponseCurrencyCodeInfoNoPaging struct {
	Results []CurrencyCodeInfo `json:"results,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Results     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CollectionResponseCurrencyCodeInfoNoPaging) RawJSON() string { return r.JSON.raw }
func (r *CollectionResponseCurrencyCodeInfoNoPaging) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CollectionResponseExchangeRateForwardPaging struct {
	Results []ExchangeRate       `json:"results,required"`
	Paging  shared.ForwardPaging `json:"paging"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Results     respjson.Field
		Paging      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CollectionResponseExchangeRateForwardPaging) RawJSON() string { return r.JSON.raw }
func (r *CollectionResponseExchangeRateForwardPaging) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CollectionResponseExchangeRateNoPaging struct {
	Results []ExchangeRate `json:"results,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Results     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CollectionResponseExchangeRateNoPaging) RawJSON() string { return r.JSON.raw }
func (r *CollectionResponseExchangeRateNoPaging) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CompanyCurrency struct {
	ID        string    `json:"id,required"`
	CreatedAt time.Time `json:"createdAt,required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CompanyCurrency) RawJSON() string { return r.JSON.raw }
func (r *CompanyCurrency) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property CurrencyCode is required.
type CompanyCurrencyUpdateRequestParam struct {
	// Any of "AED", "AFN", "ALL", "AMD", "ANG", "AOA", "ARS", "AUD", "AWG", "AZN",
	// "BAM", "BBD", "BDT", "BGN", "BHD", "BIF", "BMD", "BND", "BOB", "BOV", "BRL",
	// "BSD", "BTN", "BWP", "BYN", "BZD", "CAD", "CDF", "CHE", "CHF", "CHW", "CLF",
	// "CLP", "CNY", "COP", "COU", "CRC", "CUC", "CUP", "CVE", "CZK", "DJF", "DKK",
	// "DOP", "DZD", "EGP", "ERN", "ETB", "EUR", "FJD", "FKP", "GBP", "GEL", "GHS",
	// "GIP", "GMD", "GNF", "GTQ", "GYD", "HKD", "HNL", "HRK", "HTG", "HUF", "IDR",
	// "ILS", "INR", "IQD", "IRR", "ISK", "JMD", "JOD", "JPY", "KES", "KGS", "KHR",
	// "KMF", "KPW", "KRW", "KWD", "KYD", "KZT", "LAK", "LBP", "LKR", "LRD", "LSL",
	// "LYD", "MAD", "MDL", "MGA", "MKD", "MMK", "MNT", "MOP", "MRU", "MUR", "MVR",
	// "MWK", "MXN", "MXV", "MYR", "MZN", "NAD", "NGN", "NIO", "NOK", "NPR", "NZD",
	// "OMR", "PAB", "PEN", "PGK", "PHP", "PKR", "PLN", "PYG", "QAR", "RON", "RSD",
	// "RUB", "RWF", "SAR", "SBD", "SCR", "SDG", "SEK", "SGD", "SHP", "SLL", "SOS",
	// "SRD", "SSP", "STN", "SVC", "SYP", "SZL", "THB", "TJS", "TMT", "TND", "TOP",
	// "TRY", "TTD", "TWD", "TZS", "UAH", "UGX", "USD", "USN", "UYI", "UYU", "UZS",
	// "VEF", "VND", "VUV", "WST", "XAF", "XAG", "XAU", "XBA", "XBB", "XBC", "XBD",
	// "XCD", "XDR", "XOF", "XPD", "XPF", "XPT", "XSU", "XUA", "YER", "ZAR", "ZMW",
	// "ZWL".
	CurrencyCode CompanyCurrencyUpdateRequestCurrencyCode `json:"currencyCode,omitzero,required"`
	paramObj
}

func (r CompanyCurrencyUpdateRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow CompanyCurrencyUpdateRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CompanyCurrencyUpdateRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CompanyCurrencyUpdateRequestCurrencyCode string

const (
	CompanyCurrencyUpdateRequestCurrencyCodeAed CompanyCurrencyUpdateRequestCurrencyCode = "AED"
	CompanyCurrencyUpdateRequestCurrencyCodeAfn CompanyCurrencyUpdateRequestCurrencyCode = "AFN"
	CompanyCurrencyUpdateRequestCurrencyCodeAll CompanyCurrencyUpdateRequestCurrencyCode = "ALL"
	CompanyCurrencyUpdateRequestCurrencyCodeAmd CompanyCurrencyUpdateRequestCurrencyCode = "AMD"
	CompanyCurrencyUpdateRequestCurrencyCodeAng CompanyCurrencyUpdateRequestCurrencyCode = "ANG"
	CompanyCurrencyUpdateRequestCurrencyCodeAoa CompanyCurrencyUpdateRequestCurrencyCode = "AOA"
	CompanyCurrencyUpdateRequestCurrencyCodeArs CompanyCurrencyUpdateRequestCurrencyCode = "ARS"
	CompanyCurrencyUpdateRequestCurrencyCodeAud CompanyCurrencyUpdateRequestCurrencyCode = "AUD"
	CompanyCurrencyUpdateRequestCurrencyCodeAwg CompanyCurrencyUpdateRequestCurrencyCode = "AWG"
	CompanyCurrencyUpdateRequestCurrencyCodeAzn CompanyCurrencyUpdateRequestCurrencyCode = "AZN"
	CompanyCurrencyUpdateRequestCurrencyCodeBam CompanyCurrencyUpdateRequestCurrencyCode = "BAM"
	CompanyCurrencyUpdateRequestCurrencyCodeBbd CompanyCurrencyUpdateRequestCurrencyCode = "BBD"
	CompanyCurrencyUpdateRequestCurrencyCodeBdt CompanyCurrencyUpdateRequestCurrencyCode = "BDT"
	CompanyCurrencyUpdateRequestCurrencyCodeBgn CompanyCurrencyUpdateRequestCurrencyCode = "BGN"
	CompanyCurrencyUpdateRequestCurrencyCodeBhd CompanyCurrencyUpdateRequestCurrencyCode = "BHD"
	CompanyCurrencyUpdateRequestCurrencyCodeBif CompanyCurrencyUpdateRequestCurrencyCode = "BIF"
	CompanyCurrencyUpdateRequestCurrencyCodeBmd CompanyCurrencyUpdateRequestCurrencyCode = "BMD"
	CompanyCurrencyUpdateRequestCurrencyCodeBnd CompanyCurrencyUpdateRequestCurrencyCode = "BND"
	CompanyCurrencyUpdateRequestCurrencyCodeBob CompanyCurrencyUpdateRequestCurrencyCode = "BOB"
	CompanyCurrencyUpdateRequestCurrencyCodeBov CompanyCurrencyUpdateRequestCurrencyCode = "BOV"
	CompanyCurrencyUpdateRequestCurrencyCodeBrl CompanyCurrencyUpdateRequestCurrencyCode = "BRL"
	CompanyCurrencyUpdateRequestCurrencyCodeBsd CompanyCurrencyUpdateRequestCurrencyCode = "BSD"
	CompanyCurrencyUpdateRequestCurrencyCodeBtn CompanyCurrencyUpdateRequestCurrencyCode = "BTN"
	CompanyCurrencyUpdateRequestCurrencyCodeBwp CompanyCurrencyUpdateRequestCurrencyCode = "BWP"
	CompanyCurrencyUpdateRequestCurrencyCodeByn CompanyCurrencyUpdateRequestCurrencyCode = "BYN"
	CompanyCurrencyUpdateRequestCurrencyCodeBzd CompanyCurrencyUpdateRequestCurrencyCode = "BZD"
	CompanyCurrencyUpdateRequestCurrencyCodeCad CompanyCurrencyUpdateRequestCurrencyCode = "CAD"
	CompanyCurrencyUpdateRequestCurrencyCodeCdf CompanyCurrencyUpdateRequestCurrencyCode = "CDF"
	CompanyCurrencyUpdateRequestCurrencyCodeChe CompanyCurrencyUpdateRequestCurrencyCode = "CHE"
	CompanyCurrencyUpdateRequestCurrencyCodeChf CompanyCurrencyUpdateRequestCurrencyCode = "CHF"
	CompanyCurrencyUpdateRequestCurrencyCodeChw CompanyCurrencyUpdateRequestCurrencyCode = "CHW"
	CompanyCurrencyUpdateRequestCurrencyCodeClf CompanyCurrencyUpdateRequestCurrencyCode = "CLF"
	CompanyCurrencyUpdateRequestCurrencyCodeClp CompanyCurrencyUpdateRequestCurrencyCode = "CLP"
	CompanyCurrencyUpdateRequestCurrencyCodeCny CompanyCurrencyUpdateRequestCurrencyCode = "CNY"
	CompanyCurrencyUpdateRequestCurrencyCodeCop CompanyCurrencyUpdateRequestCurrencyCode = "COP"
	CompanyCurrencyUpdateRequestCurrencyCodeCou CompanyCurrencyUpdateRequestCurrencyCode = "COU"
	CompanyCurrencyUpdateRequestCurrencyCodeCrc CompanyCurrencyUpdateRequestCurrencyCode = "CRC"
	CompanyCurrencyUpdateRequestCurrencyCodeCuc CompanyCurrencyUpdateRequestCurrencyCode = "CUC"
	CompanyCurrencyUpdateRequestCurrencyCodeCup CompanyCurrencyUpdateRequestCurrencyCode = "CUP"
	CompanyCurrencyUpdateRequestCurrencyCodeCve CompanyCurrencyUpdateRequestCurrencyCode = "CVE"
	CompanyCurrencyUpdateRequestCurrencyCodeCzk CompanyCurrencyUpdateRequestCurrencyCode = "CZK"
	CompanyCurrencyUpdateRequestCurrencyCodeDjf CompanyCurrencyUpdateRequestCurrencyCode = "DJF"
	CompanyCurrencyUpdateRequestCurrencyCodeDkk CompanyCurrencyUpdateRequestCurrencyCode = "DKK"
	CompanyCurrencyUpdateRequestCurrencyCodeDop CompanyCurrencyUpdateRequestCurrencyCode = "DOP"
	CompanyCurrencyUpdateRequestCurrencyCodeDzd CompanyCurrencyUpdateRequestCurrencyCode = "DZD"
	CompanyCurrencyUpdateRequestCurrencyCodeEgp CompanyCurrencyUpdateRequestCurrencyCode = "EGP"
	CompanyCurrencyUpdateRequestCurrencyCodeErn CompanyCurrencyUpdateRequestCurrencyCode = "ERN"
	CompanyCurrencyUpdateRequestCurrencyCodeEtb CompanyCurrencyUpdateRequestCurrencyCode = "ETB"
	CompanyCurrencyUpdateRequestCurrencyCodeEur CompanyCurrencyUpdateRequestCurrencyCode = "EUR"
	CompanyCurrencyUpdateRequestCurrencyCodeFjd CompanyCurrencyUpdateRequestCurrencyCode = "FJD"
	CompanyCurrencyUpdateRequestCurrencyCodeFkp CompanyCurrencyUpdateRequestCurrencyCode = "FKP"
	CompanyCurrencyUpdateRequestCurrencyCodeGbp CompanyCurrencyUpdateRequestCurrencyCode = "GBP"
	CompanyCurrencyUpdateRequestCurrencyCodeGel CompanyCurrencyUpdateRequestCurrencyCode = "GEL"
	CompanyCurrencyUpdateRequestCurrencyCodeGhs CompanyCurrencyUpdateRequestCurrencyCode = "GHS"
	CompanyCurrencyUpdateRequestCurrencyCodeGip CompanyCurrencyUpdateRequestCurrencyCode = "GIP"
	CompanyCurrencyUpdateRequestCurrencyCodeGmd CompanyCurrencyUpdateRequestCurrencyCode = "GMD"
	CompanyCurrencyUpdateRequestCurrencyCodeGnf CompanyCurrencyUpdateRequestCurrencyCode = "GNF"
	CompanyCurrencyUpdateRequestCurrencyCodeGtq CompanyCurrencyUpdateRequestCurrencyCode = "GTQ"
	CompanyCurrencyUpdateRequestCurrencyCodeGyd CompanyCurrencyUpdateRequestCurrencyCode = "GYD"
	CompanyCurrencyUpdateRequestCurrencyCodeHkd CompanyCurrencyUpdateRequestCurrencyCode = "HKD"
	CompanyCurrencyUpdateRequestCurrencyCodeHnl CompanyCurrencyUpdateRequestCurrencyCode = "HNL"
	CompanyCurrencyUpdateRequestCurrencyCodeHrk CompanyCurrencyUpdateRequestCurrencyCode = "HRK"
	CompanyCurrencyUpdateRequestCurrencyCodeHtg CompanyCurrencyUpdateRequestCurrencyCode = "HTG"
	CompanyCurrencyUpdateRequestCurrencyCodeHuf CompanyCurrencyUpdateRequestCurrencyCode = "HUF"
	CompanyCurrencyUpdateRequestCurrencyCodeIdr CompanyCurrencyUpdateRequestCurrencyCode = "IDR"
	CompanyCurrencyUpdateRequestCurrencyCodeIls CompanyCurrencyUpdateRequestCurrencyCode = "ILS"
	CompanyCurrencyUpdateRequestCurrencyCodeInr CompanyCurrencyUpdateRequestCurrencyCode = "INR"
	CompanyCurrencyUpdateRequestCurrencyCodeIqd CompanyCurrencyUpdateRequestCurrencyCode = "IQD"
	CompanyCurrencyUpdateRequestCurrencyCodeIrr CompanyCurrencyUpdateRequestCurrencyCode = "IRR"
	CompanyCurrencyUpdateRequestCurrencyCodeIsk CompanyCurrencyUpdateRequestCurrencyCode = "ISK"
	CompanyCurrencyUpdateRequestCurrencyCodeJmd CompanyCurrencyUpdateRequestCurrencyCode = "JMD"
	CompanyCurrencyUpdateRequestCurrencyCodeJod CompanyCurrencyUpdateRequestCurrencyCode = "JOD"
	CompanyCurrencyUpdateRequestCurrencyCodeJpy CompanyCurrencyUpdateRequestCurrencyCode = "JPY"
	CompanyCurrencyUpdateRequestCurrencyCodeKes CompanyCurrencyUpdateRequestCurrencyCode = "KES"
	CompanyCurrencyUpdateRequestCurrencyCodeKgs CompanyCurrencyUpdateRequestCurrencyCode = "KGS"
	CompanyCurrencyUpdateRequestCurrencyCodeKhr CompanyCurrencyUpdateRequestCurrencyCode = "KHR"
	CompanyCurrencyUpdateRequestCurrencyCodeKmf CompanyCurrencyUpdateRequestCurrencyCode = "KMF"
	CompanyCurrencyUpdateRequestCurrencyCodeKpw CompanyCurrencyUpdateRequestCurrencyCode = "KPW"
	CompanyCurrencyUpdateRequestCurrencyCodeKrw CompanyCurrencyUpdateRequestCurrencyCode = "KRW"
	CompanyCurrencyUpdateRequestCurrencyCodeKwd CompanyCurrencyUpdateRequestCurrencyCode = "KWD"
	CompanyCurrencyUpdateRequestCurrencyCodeKyd CompanyCurrencyUpdateRequestCurrencyCode = "KYD"
	CompanyCurrencyUpdateRequestCurrencyCodeKzt CompanyCurrencyUpdateRequestCurrencyCode = "KZT"
	CompanyCurrencyUpdateRequestCurrencyCodeLak CompanyCurrencyUpdateRequestCurrencyCode = "LAK"
	CompanyCurrencyUpdateRequestCurrencyCodeLbp CompanyCurrencyUpdateRequestCurrencyCode = "LBP"
	CompanyCurrencyUpdateRequestCurrencyCodeLkr CompanyCurrencyUpdateRequestCurrencyCode = "LKR"
	CompanyCurrencyUpdateRequestCurrencyCodeLrd CompanyCurrencyUpdateRequestCurrencyCode = "LRD"
	CompanyCurrencyUpdateRequestCurrencyCodeLsl CompanyCurrencyUpdateRequestCurrencyCode = "LSL"
	CompanyCurrencyUpdateRequestCurrencyCodeLyd CompanyCurrencyUpdateRequestCurrencyCode = "LYD"
	CompanyCurrencyUpdateRequestCurrencyCodeMad CompanyCurrencyUpdateRequestCurrencyCode = "MAD"
	CompanyCurrencyUpdateRequestCurrencyCodeMdl CompanyCurrencyUpdateRequestCurrencyCode = "MDL"
	CompanyCurrencyUpdateRequestCurrencyCodeMga CompanyCurrencyUpdateRequestCurrencyCode = "MGA"
	CompanyCurrencyUpdateRequestCurrencyCodeMkd CompanyCurrencyUpdateRequestCurrencyCode = "MKD"
	CompanyCurrencyUpdateRequestCurrencyCodeMmk CompanyCurrencyUpdateRequestCurrencyCode = "MMK"
	CompanyCurrencyUpdateRequestCurrencyCodeMnt CompanyCurrencyUpdateRequestCurrencyCode = "MNT"
	CompanyCurrencyUpdateRequestCurrencyCodeMop CompanyCurrencyUpdateRequestCurrencyCode = "MOP"
	CompanyCurrencyUpdateRequestCurrencyCodeMru CompanyCurrencyUpdateRequestCurrencyCode = "MRU"
	CompanyCurrencyUpdateRequestCurrencyCodeMur CompanyCurrencyUpdateRequestCurrencyCode = "MUR"
	CompanyCurrencyUpdateRequestCurrencyCodeMvr CompanyCurrencyUpdateRequestCurrencyCode = "MVR"
	CompanyCurrencyUpdateRequestCurrencyCodeMwk CompanyCurrencyUpdateRequestCurrencyCode = "MWK"
	CompanyCurrencyUpdateRequestCurrencyCodeMxn CompanyCurrencyUpdateRequestCurrencyCode = "MXN"
	CompanyCurrencyUpdateRequestCurrencyCodeMxv CompanyCurrencyUpdateRequestCurrencyCode = "MXV"
	CompanyCurrencyUpdateRequestCurrencyCodeMyr CompanyCurrencyUpdateRequestCurrencyCode = "MYR"
	CompanyCurrencyUpdateRequestCurrencyCodeMzn CompanyCurrencyUpdateRequestCurrencyCode = "MZN"
	CompanyCurrencyUpdateRequestCurrencyCodeNad CompanyCurrencyUpdateRequestCurrencyCode = "NAD"
	CompanyCurrencyUpdateRequestCurrencyCodeNgn CompanyCurrencyUpdateRequestCurrencyCode = "NGN"
	CompanyCurrencyUpdateRequestCurrencyCodeNio CompanyCurrencyUpdateRequestCurrencyCode = "NIO"
	CompanyCurrencyUpdateRequestCurrencyCodeNok CompanyCurrencyUpdateRequestCurrencyCode = "NOK"
	CompanyCurrencyUpdateRequestCurrencyCodeNpr CompanyCurrencyUpdateRequestCurrencyCode = "NPR"
	CompanyCurrencyUpdateRequestCurrencyCodeNzd CompanyCurrencyUpdateRequestCurrencyCode = "NZD"
	CompanyCurrencyUpdateRequestCurrencyCodeOmr CompanyCurrencyUpdateRequestCurrencyCode = "OMR"
	CompanyCurrencyUpdateRequestCurrencyCodePab CompanyCurrencyUpdateRequestCurrencyCode = "PAB"
	CompanyCurrencyUpdateRequestCurrencyCodePen CompanyCurrencyUpdateRequestCurrencyCode = "PEN"
	CompanyCurrencyUpdateRequestCurrencyCodePgk CompanyCurrencyUpdateRequestCurrencyCode = "PGK"
	CompanyCurrencyUpdateRequestCurrencyCodePhp CompanyCurrencyUpdateRequestCurrencyCode = "PHP"
	CompanyCurrencyUpdateRequestCurrencyCodePkr CompanyCurrencyUpdateRequestCurrencyCode = "PKR"
	CompanyCurrencyUpdateRequestCurrencyCodePln CompanyCurrencyUpdateRequestCurrencyCode = "PLN"
	CompanyCurrencyUpdateRequestCurrencyCodePyg CompanyCurrencyUpdateRequestCurrencyCode = "PYG"
	CompanyCurrencyUpdateRequestCurrencyCodeQar CompanyCurrencyUpdateRequestCurrencyCode = "QAR"
	CompanyCurrencyUpdateRequestCurrencyCodeRon CompanyCurrencyUpdateRequestCurrencyCode = "RON"
	CompanyCurrencyUpdateRequestCurrencyCodeRsd CompanyCurrencyUpdateRequestCurrencyCode = "RSD"
	CompanyCurrencyUpdateRequestCurrencyCodeRub CompanyCurrencyUpdateRequestCurrencyCode = "RUB"
	CompanyCurrencyUpdateRequestCurrencyCodeRwf CompanyCurrencyUpdateRequestCurrencyCode = "RWF"
	CompanyCurrencyUpdateRequestCurrencyCodeSar CompanyCurrencyUpdateRequestCurrencyCode = "SAR"
	CompanyCurrencyUpdateRequestCurrencyCodeSbd CompanyCurrencyUpdateRequestCurrencyCode = "SBD"
	CompanyCurrencyUpdateRequestCurrencyCodeScr CompanyCurrencyUpdateRequestCurrencyCode = "SCR"
	CompanyCurrencyUpdateRequestCurrencyCodeSdg CompanyCurrencyUpdateRequestCurrencyCode = "SDG"
	CompanyCurrencyUpdateRequestCurrencyCodeSek CompanyCurrencyUpdateRequestCurrencyCode = "SEK"
	CompanyCurrencyUpdateRequestCurrencyCodeSgd CompanyCurrencyUpdateRequestCurrencyCode = "SGD"
	CompanyCurrencyUpdateRequestCurrencyCodeShp CompanyCurrencyUpdateRequestCurrencyCode = "SHP"
	CompanyCurrencyUpdateRequestCurrencyCodeSll CompanyCurrencyUpdateRequestCurrencyCode = "SLL"
	CompanyCurrencyUpdateRequestCurrencyCodeSos CompanyCurrencyUpdateRequestCurrencyCode = "SOS"
	CompanyCurrencyUpdateRequestCurrencyCodeSrd CompanyCurrencyUpdateRequestCurrencyCode = "SRD"
	CompanyCurrencyUpdateRequestCurrencyCodeSsp CompanyCurrencyUpdateRequestCurrencyCode = "SSP"
	CompanyCurrencyUpdateRequestCurrencyCodeStn CompanyCurrencyUpdateRequestCurrencyCode = "STN"
	CompanyCurrencyUpdateRequestCurrencyCodeSvc CompanyCurrencyUpdateRequestCurrencyCode = "SVC"
	CompanyCurrencyUpdateRequestCurrencyCodeSyp CompanyCurrencyUpdateRequestCurrencyCode = "SYP"
	CompanyCurrencyUpdateRequestCurrencyCodeSzl CompanyCurrencyUpdateRequestCurrencyCode = "SZL"
	CompanyCurrencyUpdateRequestCurrencyCodeThb CompanyCurrencyUpdateRequestCurrencyCode = "THB"
	CompanyCurrencyUpdateRequestCurrencyCodeTjs CompanyCurrencyUpdateRequestCurrencyCode = "TJS"
	CompanyCurrencyUpdateRequestCurrencyCodeTmt CompanyCurrencyUpdateRequestCurrencyCode = "TMT"
	CompanyCurrencyUpdateRequestCurrencyCodeTnd CompanyCurrencyUpdateRequestCurrencyCode = "TND"
	CompanyCurrencyUpdateRequestCurrencyCodeTop CompanyCurrencyUpdateRequestCurrencyCode = "TOP"
	CompanyCurrencyUpdateRequestCurrencyCodeTry CompanyCurrencyUpdateRequestCurrencyCode = "TRY"
	CompanyCurrencyUpdateRequestCurrencyCodeTtd CompanyCurrencyUpdateRequestCurrencyCode = "TTD"
	CompanyCurrencyUpdateRequestCurrencyCodeTwd CompanyCurrencyUpdateRequestCurrencyCode = "TWD"
	CompanyCurrencyUpdateRequestCurrencyCodeTzs CompanyCurrencyUpdateRequestCurrencyCode = "TZS"
	CompanyCurrencyUpdateRequestCurrencyCodeUah CompanyCurrencyUpdateRequestCurrencyCode = "UAH"
	CompanyCurrencyUpdateRequestCurrencyCodeUgx CompanyCurrencyUpdateRequestCurrencyCode = "UGX"
	CompanyCurrencyUpdateRequestCurrencyCodeUsd CompanyCurrencyUpdateRequestCurrencyCode = "USD"
	CompanyCurrencyUpdateRequestCurrencyCodeUsn CompanyCurrencyUpdateRequestCurrencyCode = "USN"
	CompanyCurrencyUpdateRequestCurrencyCodeUyi CompanyCurrencyUpdateRequestCurrencyCode = "UYI"
	CompanyCurrencyUpdateRequestCurrencyCodeUyu CompanyCurrencyUpdateRequestCurrencyCode = "UYU"
	CompanyCurrencyUpdateRequestCurrencyCodeUzs CompanyCurrencyUpdateRequestCurrencyCode = "UZS"
	CompanyCurrencyUpdateRequestCurrencyCodeVef CompanyCurrencyUpdateRequestCurrencyCode = "VEF"
	CompanyCurrencyUpdateRequestCurrencyCodeVnd CompanyCurrencyUpdateRequestCurrencyCode = "VND"
	CompanyCurrencyUpdateRequestCurrencyCodeVuv CompanyCurrencyUpdateRequestCurrencyCode = "VUV"
	CompanyCurrencyUpdateRequestCurrencyCodeWst CompanyCurrencyUpdateRequestCurrencyCode = "WST"
	CompanyCurrencyUpdateRequestCurrencyCodeXaf CompanyCurrencyUpdateRequestCurrencyCode = "XAF"
	CompanyCurrencyUpdateRequestCurrencyCodeXag CompanyCurrencyUpdateRequestCurrencyCode = "XAG"
	CompanyCurrencyUpdateRequestCurrencyCodeXau CompanyCurrencyUpdateRequestCurrencyCode = "XAU"
	CompanyCurrencyUpdateRequestCurrencyCodeXba CompanyCurrencyUpdateRequestCurrencyCode = "XBA"
	CompanyCurrencyUpdateRequestCurrencyCodeXbb CompanyCurrencyUpdateRequestCurrencyCode = "XBB"
	CompanyCurrencyUpdateRequestCurrencyCodeXbc CompanyCurrencyUpdateRequestCurrencyCode = "XBC"
	CompanyCurrencyUpdateRequestCurrencyCodeXbd CompanyCurrencyUpdateRequestCurrencyCode = "XBD"
	CompanyCurrencyUpdateRequestCurrencyCodeXcd CompanyCurrencyUpdateRequestCurrencyCode = "XCD"
	CompanyCurrencyUpdateRequestCurrencyCodeXdr CompanyCurrencyUpdateRequestCurrencyCode = "XDR"
	CompanyCurrencyUpdateRequestCurrencyCodeXof CompanyCurrencyUpdateRequestCurrencyCode = "XOF"
	CompanyCurrencyUpdateRequestCurrencyCodeXpd CompanyCurrencyUpdateRequestCurrencyCode = "XPD"
	CompanyCurrencyUpdateRequestCurrencyCodeXpf CompanyCurrencyUpdateRequestCurrencyCode = "XPF"
	CompanyCurrencyUpdateRequestCurrencyCodeXpt CompanyCurrencyUpdateRequestCurrencyCode = "XPT"
	CompanyCurrencyUpdateRequestCurrencyCodeXsu CompanyCurrencyUpdateRequestCurrencyCode = "XSU"
	CompanyCurrencyUpdateRequestCurrencyCodeXua CompanyCurrencyUpdateRequestCurrencyCode = "XUA"
	CompanyCurrencyUpdateRequestCurrencyCodeYer CompanyCurrencyUpdateRequestCurrencyCode = "YER"
	CompanyCurrencyUpdateRequestCurrencyCodeZar CompanyCurrencyUpdateRequestCurrencyCode = "ZAR"
	CompanyCurrencyUpdateRequestCurrencyCodeZmw CompanyCurrencyUpdateRequestCurrencyCode = "ZMW"
	CompanyCurrencyUpdateRequestCurrencyCodeZwl CompanyCurrencyUpdateRequestCurrencyCode = "ZWL"
)

type CurrencyCodeInfo struct {
	CurrencyCode string `json:"currencyCode,required"`
	CurrencyName string `json:"currencyName,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CurrencyCode respjson.Field
		CurrencyName respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CurrencyCodeInfo) RawJSON() string { return r.JSON.raw }
func (r *CurrencyCodeInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property CurrencyCode is required.
type CurrencyCreateRequestParam struct {
	// Any of "AED", "AFN", "ALL", "AMD", "ANG", "AOA", "ARS", "AUD", "AWG", "AZN",
	// "BAM", "BBD", "BDT", "BGN", "BHD", "BIF", "BMD", "BND", "BOB", "BOV", "BRL",
	// "BSD", "BTN", "BWP", "BYN", "BZD", "CAD", "CDF", "CHE", "CHF", "CHW", "CLF",
	// "CLP", "CNY", "COP", "COU", "CRC", "CUC", "CUP", "CVE", "CZK", "DJF", "DKK",
	// "DOP", "DZD", "EGP", "ERN", "ETB", "EUR", "FJD", "FKP", "GBP", "GEL", "GHS",
	// "GIP", "GMD", "GNF", "GTQ", "GYD", "HKD", "HNL", "HRK", "HTG", "HUF", "IDR",
	// "ILS", "INR", "IQD", "IRR", "ISK", "JMD", "JOD", "JPY", "KES", "KGS", "KHR",
	// "KMF", "KPW", "KRW", "KWD", "KYD", "KZT", "LAK", "LBP", "LKR", "LRD", "LSL",
	// "LYD", "MAD", "MDL", "MGA", "MKD", "MMK", "MNT", "MOP", "MRU", "MUR", "MVR",
	// "MWK", "MXN", "MXV", "MYR", "MZN", "NAD", "NGN", "NIO", "NOK", "NPR", "NZD",
	// "OMR", "PAB", "PEN", "PGK", "PHP", "PKR", "PLN", "PYG", "QAR", "RON", "RSD",
	// "RUB", "RWF", "SAR", "SBD", "SCR", "SDG", "SEK", "SGD", "SHP", "SLL", "SOS",
	// "SRD", "SSP", "STN", "SVC", "SYP", "SZL", "THB", "TJS", "TMT", "TND", "TOP",
	// "TRY", "TTD", "TWD", "TZS", "UAH", "UGX", "USD", "USN", "UYI", "UYU", "UZS",
	// "VEF", "VND", "VUV", "WST", "XAF", "XAG", "XAU", "XBA", "XBB", "XBC", "XBD",
	// "XCD", "XDR", "XOF", "XPD", "XPF", "XPT", "XSU", "XUA", "YER", "ZAR", "ZMW",
	// "ZWL".
	CurrencyCode CurrencyCreateRequestCurrencyCode `json:"currencyCode,omitzero,required"`
	paramObj
}

func (r CurrencyCreateRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow CurrencyCreateRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CurrencyCreateRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CurrencyCreateRequestCurrencyCode string

const (
	CurrencyCreateRequestCurrencyCodeAed CurrencyCreateRequestCurrencyCode = "AED"
	CurrencyCreateRequestCurrencyCodeAfn CurrencyCreateRequestCurrencyCode = "AFN"
	CurrencyCreateRequestCurrencyCodeAll CurrencyCreateRequestCurrencyCode = "ALL"
	CurrencyCreateRequestCurrencyCodeAmd CurrencyCreateRequestCurrencyCode = "AMD"
	CurrencyCreateRequestCurrencyCodeAng CurrencyCreateRequestCurrencyCode = "ANG"
	CurrencyCreateRequestCurrencyCodeAoa CurrencyCreateRequestCurrencyCode = "AOA"
	CurrencyCreateRequestCurrencyCodeArs CurrencyCreateRequestCurrencyCode = "ARS"
	CurrencyCreateRequestCurrencyCodeAud CurrencyCreateRequestCurrencyCode = "AUD"
	CurrencyCreateRequestCurrencyCodeAwg CurrencyCreateRequestCurrencyCode = "AWG"
	CurrencyCreateRequestCurrencyCodeAzn CurrencyCreateRequestCurrencyCode = "AZN"
	CurrencyCreateRequestCurrencyCodeBam CurrencyCreateRequestCurrencyCode = "BAM"
	CurrencyCreateRequestCurrencyCodeBbd CurrencyCreateRequestCurrencyCode = "BBD"
	CurrencyCreateRequestCurrencyCodeBdt CurrencyCreateRequestCurrencyCode = "BDT"
	CurrencyCreateRequestCurrencyCodeBgn CurrencyCreateRequestCurrencyCode = "BGN"
	CurrencyCreateRequestCurrencyCodeBhd CurrencyCreateRequestCurrencyCode = "BHD"
	CurrencyCreateRequestCurrencyCodeBif CurrencyCreateRequestCurrencyCode = "BIF"
	CurrencyCreateRequestCurrencyCodeBmd CurrencyCreateRequestCurrencyCode = "BMD"
	CurrencyCreateRequestCurrencyCodeBnd CurrencyCreateRequestCurrencyCode = "BND"
	CurrencyCreateRequestCurrencyCodeBob CurrencyCreateRequestCurrencyCode = "BOB"
	CurrencyCreateRequestCurrencyCodeBov CurrencyCreateRequestCurrencyCode = "BOV"
	CurrencyCreateRequestCurrencyCodeBrl CurrencyCreateRequestCurrencyCode = "BRL"
	CurrencyCreateRequestCurrencyCodeBsd CurrencyCreateRequestCurrencyCode = "BSD"
	CurrencyCreateRequestCurrencyCodeBtn CurrencyCreateRequestCurrencyCode = "BTN"
	CurrencyCreateRequestCurrencyCodeBwp CurrencyCreateRequestCurrencyCode = "BWP"
	CurrencyCreateRequestCurrencyCodeByn CurrencyCreateRequestCurrencyCode = "BYN"
	CurrencyCreateRequestCurrencyCodeBzd CurrencyCreateRequestCurrencyCode = "BZD"
	CurrencyCreateRequestCurrencyCodeCad CurrencyCreateRequestCurrencyCode = "CAD"
	CurrencyCreateRequestCurrencyCodeCdf CurrencyCreateRequestCurrencyCode = "CDF"
	CurrencyCreateRequestCurrencyCodeChe CurrencyCreateRequestCurrencyCode = "CHE"
	CurrencyCreateRequestCurrencyCodeChf CurrencyCreateRequestCurrencyCode = "CHF"
	CurrencyCreateRequestCurrencyCodeChw CurrencyCreateRequestCurrencyCode = "CHW"
	CurrencyCreateRequestCurrencyCodeClf CurrencyCreateRequestCurrencyCode = "CLF"
	CurrencyCreateRequestCurrencyCodeClp CurrencyCreateRequestCurrencyCode = "CLP"
	CurrencyCreateRequestCurrencyCodeCny CurrencyCreateRequestCurrencyCode = "CNY"
	CurrencyCreateRequestCurrencyCodeCop CurrencyCreateRequestCurrencyCode = "COP"
	CurrencyCreateRequestCurrencyCodeCou CurrencyCreateRequestCurrencyCode = "COU"
	CurrencyCreateRequestCurrencyCodeCrc CurrencyCreateRequestCurrencyCode = "CRC"
	CurrencyCreateRequestCurrencyCodeCuc CurrencyCreateRequestCurrencyCode = "CUC"
	CurrencyCreateRequestCurrencyCodeCup CurrencyCreateRequestCurrencyCode = "CUP"
	CurrencyCreateRequestCurrencyCodeCve CurrencyCreateRequestCurrencyCode = "CVE"
	CurrencyCreateRequestCurrencyCodeCzk CurrencyCreateRequestCurrencyCode = "CZK"
	CurrencyCreateRequestCurrencyCodeDjf CurrencyCreateRequestCurrencyCode = "DJF"
	CurrencyCreateRequestCurrencyCodeDkk CurrencyCreateRequestCurrencyCode = "DKK"
	CurrencyCreateRequestCurrencyCodeDop CurrencyCreateRequestCurrencyCode = "DOP"
	CurrencyCreateRequestCurrencyCodeDzd CurrencyCreateRequestCurrencyCode = "DZD"
	CurrencyCreateRequestCurrencyCodeEgp CurrencyCreateRequestCurrencyCode = "EGP"
	CurrencyCreateRequestCurrencyCodeErn CurrencyCreateRequestCurrencyCode = "ERN"
	CurrencyCreateRequestCurrencyCodeEtb CurrencyCreateRequestCurrencyCode = "ETB"
	CurrencyCreateRequestCurrencyCodeEur CurrencyCreateRequestCurrencyCode = "EUR"
	CurrencyCreateRequestCurrencyCodeFjd CurrencyCreateRequestCurrencyCode = "FJD"
	CurrencyCreateRequestCurrencyCodeFkp CurrencyCreateRequestCurrencyCode = "FKP"
	CurrencyCreateRequestCurrencyCodeGbp CurrencyCreateRequestCurrencyCode = "GBP"
	CurrencyCreateRequestCurrencyCodeGel CurrencyCreateRequestCurrencyCode = "GEL"
	CurrencyCreateRequestCurrencyCodeGhs CurrencyCreateRequestCurrencyCode = "GHS"
	CurrencyCreateRequestCurrencyCodeGip CurrencyCreateRequestCurrencyCode = "GIP"
	CurrencyCreateRequestCurrencyCodeGmd CurrencyCreateRequestCurrencyCode = "GMD"
	CurrencyCreateRequestCurrencyCodeGnf CurrencyCreateRequestCurrencyCode = "GNF"
	CurrencyCreateRequestCurrencyCodeGtq CurrencyCreateRequestCurrencyCode = "GTQ"
	CurrencyCreateRequestCurrencyCodeGyd CurrencyCreateRequestCurrencyCode = "GYD"
	CurrencyCreateRequestCurrencyCodeHkd CurrencyCreateRequestCurrencyCode = "HKD"
	CurrencyCreateRequestCurrencyCodeHnl CurrencyCreateRequestCurrencyCode = "HNL"
	CurrencyCreateRequestCurrencyCodeHrk CurrencyCreateRequestCurrencyCode = "HRK"
	CurrencyCreateRequestCurrencyCodeHtg CurrencyCreateRequestCurrencyCode = "HTG"
	CurrencyCreateRequestCurrencyCodeHuf CurrencyCreateRequestCurrencyCode = "HUF"
	CurrencyCreateRequestCurrencyCodeIdr CurrencyCreateRequestCurrencyCode = "IDR"
	CurrencyCreateRequestCurrencyCodeIls CurrencyCreateRequestCurrencyCode = "ILS"
	CurrencyCreateRequestCurrencyCodeInr CurrencyCreateRequestCurrencyCode = "INR"
	CurrencyCreateRequestCurrencyCodeIqd CurrencyCreateRequestCurrencyCode = "IQD"
	CurrencyCreateRequestCurrencyCodeIrr CurrencyCreateRequestCurrencyCode = "IRR"
	CurrencyCreateRequestCurrencyCodeIsk CurrencyCreateRequestCurrencyCode = "ISK"
	CurrencyCreateRequestCurrencyCodeJmd CurrencyCreateRequestCurrencyCode = "JMD"
	CurrencyCreateRequestCurrencyCodeJod CurrencyCreateRequestCurrencyCode = "JOD"
	CurrencyCreateRequestCurrencyCodeJpy CurrencyCreateRequestCurrencyCode = "JPY"
	CurrencyCreateRequestCurrencyCodeKes CurrencyCreateRequestCurrencyCode = "KES"
	CurrencyCreateRequestCurrencyCodeKgs CurrencyCreateRequestCurrencyCode = "KGS"
	CurrencyCreateRequestCurrencyCodeKhr CurrencyCreateRequestCurrencyCode = "KHR"
	CurrencyCreateRequestCurrencyCodeKmf CurrencyCreateRequestCurrencyCode = "KMF"
	CurrencyCreateRequestCurrencyCodeKpw CurrencyCreateRequestCurrencyCode = "KPW"
	CurrencyCreateRequestCurrencyCodeKrw CurrencyCreateRequestCurrencyCode = "KRW"
	CurrencyCreateRequestCurrencyCodeKwd CurrencyCreateRequestCurrencyCode = "KWD"
	CurrencyCreateRequestCurrencyCodeKyd CurrencyCreateRequestCurrencyCode = "KYD"
	CurrencyCreateRequestCurrencyCodeKzt CurrencyCreateRequestCurrencyCode = "KZT"
	CurrencyCreateRequestCurrencyCodeLak CurrencyCreateRequestCurrencyCode = "LAK"
	CurrencyCreateRequestCurrencyCodeLbp CurrencyCreateRequestCurrencyCode = "LBP"
	CurrencyCreateRequestCurrencyCodeLkr CurrencyCreateRequestCurrencyCode = "LKR"
	CurrencyCreateRequestCurrencyCodeLrd CurrencyCreateRequestCurrencyCode = "LRD"
	CurrencyCreateRequestCurrencyCodeLsl CurrencyCreateRequestCurrencyCode = "LSL"
	CurrencyCreateRequestCurrencyCodeLyd CurrencyCreateRequestCurrencyCode = "LYD"
	CurrencyCreateRequestCurrencyCodeMad CurrencyCreateRequestCurrencyCode = "MAD"
	CurrencyCreateRequestCurrencyCodeMdl CurrencyCreateRequestCurrencyCode = "MDL"
	CurrencyCreateRequestCurrencyCodeMga CurrencyCreateRequestCurrencyCode = "MGA"
	CurrencyCreateRequestCurrencyCodeMkd CurrencyCreateRequestCurrencyCode = "MKD"
	CurrencyCreateRequestCurrencyCodeMmk CurrencyCreateRequestCurrencyCode = "MMK"
	CurrencyCreateRequestCurrencyCodeMnt CurrencyCreateRequestCurrencyCode = "MNT"
	CurrencyCreateRequestCurrencyCodeMop CurrencyCreateRequestCurrencyCode = "MOP"
	CurrencyCreateRequestCurrencyCodeMru CurrencyCreateRequestCurrencyCode = "MRU"
	CurrencyCreateRequestCurrencyCodeMur CurrencyCreateRequestCurrencyCode = "MUR"
	CurrencyCreateRequestCurrencyCodeMvr CurrencyCreateRequestCurrencyCode = "MVR"
	CurrencyCreateRequestCurrencyCodeMwk CurrencyCreateRequestCurrencyCode = "MWK"
	CurrencyCreateRequestCurrencyCodeMxn CurrencyCreateRequestCurrencyCode = "MXN"
	CurrencyCreateRequestCurrencyCodeMxv CurrencyCreateRequestCurrencyCode = "MXV"
	CurrencyCreateRequestCurrencyCodeMyr CurrencyCreateRequestCurrencyCode = "MYR"
	CurrencyCreateRequestCurrencyCodeMzn CurrencyCreateRequestCurrencyCode = "MZN"
	CurrencyCreateRequestCurrencyCodeNad CurrencyCreateRequestCurrencyCode = "NAD"
	CurrencyCreateRequestCurrencyCodeNgn CurrencyCreateRequestCurrencyCode = "NGN"
	CurrencyCreateRequestCurrencyCodeNio CurrencyCreateRequestCurrencyCode = "NIO"
	CurrencyCreateRequestCurrencyCodeNok CurrencyCreateRequestCurrencyCode = "NOK"
	CurrencyCreateRequestCurrencyCodeNpr CurrencyCreateRequestCurrencyCode = "NPR"
	CurrencyCreateRequestCurrencyCodeNzd CurrencyCreateRequestCurrencyCode = "NZD"
	CurrencyCreateRequestCurrencyCodeOmr CurrencyCreateRequestCurrencyCode = "OMR"
	CurrencyCreateRequestCurrencyCodePab CurrencyCreateRequestCurrencyCode = "PAB"
	CurrencyCreateRequestCurrencyCodePen CurrencyCreateRequestCurrencyCode = "PEN"
	CurrencyCreateRequestCurrencyCodePgk CurrencyCreateRequestCurrencyCode = "PGK"
	CurrencyCreateRequestCurrencyCodePhp CurrencyCreateRequestCurrencyCode = "PHP"
	CurrencyCreateRequestCurrencyCodePkr CurrencyCreateRequestCurrencyCode = "PKR"
	CurrencyCreateRequestCurrencyCodePln CurrencyCreateRequestCurrencyCode = "PLN"
	CurrencyCreateRequestCurrencyCodePyg CurrencyCreateRequestCurrencyCode = "PYG"
	CurrencyCreateRequestCurrencyCodeQar CurrencyCreateRequestCurrencyCode = "QAR"
	CurrencyCreateRequestCurrencyCodeRon CurrencyCreateRequestCurrencyCode = "RON"
	CurrencyCreateRequestCurrencyCodeRsd CurrencyCreateRequestCurrencyCode = "RSD"
	CurrencyCreateRequestCurrencyCodeRub CurrencyCreateRequestCurrencyCode = "RUB"
	CurrencyCreateRequestCurrencyCodeRwf CurrencyCreateRequestCurrencyCode = "RWF"
	CurrencyCreateRequestCurrencyCodeSar CurrencyCreateRequestCurrencyCode = "SAR"
	CurrencyCreateRequestCurrencyCodeSbd CurrencyCreateRequestCurrencyCode = "SBD"
	CurrencyCreateRequestCurrencyCodeScr CurrencyCreateRequestCurrencyCode = "SCR"
	CurrencyCreateRequestCurrencyCodeSdg CurrencyCreateRequestCurrencyCode = "SDG"
	CurrencyCreateRequestCurrencyCodeSek CurrencyCreateRequestCurrencyCode = "SEK"
	CurrencyCreateRequestCurrencyCodeSgd CurrencyCreateRequestCurrencyCode = "SGD"
	CurrencyCreateRequestCurrencyCodeShp CurrencyCreateRequestCurrencyCode = "SHP"
	CurrencyCreateRequestCurrencyCodeSll CurrencyCreateRequestCurrencyCode = "SLL"
	CurrencyCreateRequestCurrencyCodeSos CurrencyCreateRequestCurrencyCode = "SOS"
	CurrencyCreateRequestCurrencyCodeSrd CurrencyCreateRequestCurrencyCode = "SRD"
	CurrencyCreateRequestCurrencyCodeSsp CurrencyCreateRequestCurrencyCode = "SSP"
	CurrencyCreateRequestCurrencyCodeStn CurrencyCreateRequestCurrencyCode = "STN"
	CurrencyCreateRequestCurrencyCodeSvc CurrencyCreateRequestCurrencyCode = "SVC"
	CurrencyCreateRequestCurrencyCodeSyp CurrencyCreateRequestCurrencyCode = "SYP"
	CurrencyCreateRequestCurrencyCodeSzl CurrencyCreateRequestCurrencyCode = "SZL"
	CurrencyCreateRequestCurrencyCodeThb CurrencyCreateRequestCurrencyCode = "THB"
	CurrencyCreateRequestCurrencyCodeTjs CurrencyCreateRequestCurrencyCode = "TJS"
	CurrencyCreateRequestCurrencyCodeTmt CurrencyCreateRequestCurrencyCode = "TMT"
	CurrencyCreateRequestCurrencyCodeTnd CurrencyCreateRequestCurrencyCode = "TND"
	CurrencyCreateRequestCurrencyCodeTop CurrencyCreateRequestCurrencyCode = "TOP"
	CurrencyCreateRequestCurrencyCodeTry CurrencyCreateRequestCurrencyCode = "TRY"
	CurrencyCreateRequestCurrencyCodeTtd CurrencyCreateRequestCurrencyCode = "TTD"
	CurrencyCreateRequestCurrencyCodeTwd CurrencyCreateRequestCurrencyCode = "TWD"
	CurrencyCreateRequestCurrencyCodeTzs CurrencyCreateRequestCurrencyCode = "TZS"
	CurrencyCreateRequestCurrencyCodeUah CurrencyCreateRequestCurrencyCode = "UAH"
	CurrencyCreateRequestCurrencyCodeUgx CurrencyCreateRequestCurrencyCode = "UGX"
	CurrencyCreateRequestCurrencyCodeUsd CurrencyCreateRequestCurrencyCode = "USD"
	CurrencyCreateRequestCurrencyCodeUsn CurrencyCreateRequestCurrencyCode = "USN"
	CurrencyCreateRequestCurrencyCodeUyi CurrencyCreateRequestCurrencyCode = "UYI"
	CurrencyCreateRequestCurrencyCodeUyu CurrencyCreateRequestCurrencyCode = "UYU"
	CurrencyCreateRequestCurrencyCodeUzs CurrencyCreateRequestCurrencyCode = "UZS"
	CurrencyCreateRequestCurrencyCodeVef CurrencyCreateRequestCurrencyCode = "VEF"
	CurrencyCreateRequestCurrencyCodeVnd CurrencyCreateRequestCurrencyCode = "VND"
	CurrencyCreateRequestCurrencyCodeVuv CurrencyCreateRequestCurrencyCode = "VUV"
	CurrencyCreateRequestCurrencyCodeWst CurrencyCreateRequestCurrencyCode = "WST"
	CurrencyCreateRequestCurrencyCodeXaf CurrencyCreateRequestCurrencyCode = "XAF"
	CurrencyCreateRequestCurrencyCodeXag CurrencyCreateRequestCurrencyCode = "XAG"
	CurrencyCreateRequestCurrencyCodeXau CurrencyCreateRequestCurrencyCode = "XAU"
	CurrencyCreateRequestCurrencyCodeXba CurrencyCreateRequestCurrencyCode = "XBA"
	CurrencyCreateRequestCurrencyCodeXbb CurrencyCreateRequestCurrencyCode = "XBB"
	CurrencyCreateRequestCurrencyCodeXbc CurrencyCreateRequestCurrencyCode = "XBC"
	CurrencyCreateRequestCurrencyCodeXbd CurrencyCreateRequestCurrencyCode = "XBD"
	CurrencyCreateRequestCurrencyCodeXcd CurrencyCreateRequestCurrencyCode = "XCD"
	CurrencyCreateRequestCurrencyCodeXdr CurrencyCreateRequestCurrencyCode = "XDR"
	CurrencyCreateRequestCurrencyCodeXof CurrencyCreateRequestCurrencyCode = "XOF"
	CurrencyCreateRequestCurrencyCodeXpd CurrencyCreateRequestCurrencyCode = "XPD"
	CurrencyCreateRequestCurrencyCodeXpf CurrencyCreateRequestCurrencyCode = "XPF"
	CurrencyCreateRequestCurrencyCodeXpt CurrencyCreateRequestCurrencyCode = "XPT"
	CurrencyCreateRequestCurrencyCodeXsu CurrencyCreateRequestCurrencyCode = "XSU"
	CurrencyCreateRequestCurrencyCodeXua CurrencyCreateRequestCurrencyCode = "XUA"
	CurrencyCreateRequestCurrencyCodeYer CurrencyCreateRequestCurrencyCode = "YER"
	CurrencyCreateRequestCurrencyCodeZar CurrencyCreateRequestCurrencyCode = "ZAR"
	CurrencyCreateRequestCurrencyCodeZmw CurrencyCreateRequestCurrencyCode = "ZMW"
	CurrencyCreateRequestCurrencyCodeZwl CurrencyCreateRequestCurrencyCode = "ZWL"
)

// The properties FromCurrencyCode, ToCurrencyCode, VisibleInUi are required.
type CurrencyPairUpdateParam struct {
	// Any of "AED", "AFN", "ALL", "AMD", "ANG", "AOA", "ARS", "AUD", "AWG", "AZN",
	// "BAM", "BBD", "BDT", "BGN", "BHD", "BIF", "BMD", "BND", "BOB", "BOV", "BRL",
	// "BSD", "BTN", "BWP", "BYN", "BZD", "CAD", "CDF", "CHE", "CHF", "CHW", "CLF",
	// "CLP", "CNY", "COP", "COU", "CRC", "CUC", "CUP", "CVE", "CZK", "DJF", "DKK",
	// "DOP", "DZD", "EGP", "ERN", "ETB", "EUR", "FJD", "FKP", "GBP", "GEL", "GHS",
	// "GIP", "GMD", "GNF", "GTQ", "GYD", "HKD", "HNL", "HRK", "HTG", "HUF", "IDR",
	// "ILS", "INR", "IQD", "IRR", "ISK", "JMD", "JOD", "JPY", "KES", "KGS", "KHR",
	// "KMF", "KPW", "KRW", "KWD", "KYD", "KZT", "LAK", "LBP", "LKR", "LRD", "LSL",
	// "LYD", "MAD", "MDL", "MGA", "MKD", "MMK", "MNT", "MOP", "MRU", "MUR", "MVR",
	// "MWK", "MXN", "MXV", "MYR", "MZN", "NAD", "NGN", "NIO", "NOK", "NPR", "NZD",
	// "OMR", "PAB", "PEN", "PGK", "PHP", "PKR", "PLN", "PYG", "QAR", "RON", "RSD",
	// "RUB", "RWF", "SAR", "SBD", "SCR", "SDG", "SEK", "SGD", "SHP", "SLL", "SOS",
	// "SRD", "SSP", "STN", "SVC", "SYP", "SZL", "THB", "TJS", "TMT", "TND", "TOP",
	// "TRY", "TTD", "TWD", "TZS", "UAH", "UGX", "USD", "USN", "UYI", "UYU", "UZS",
	// "VEF", "VND", "VUV", "WST", "XAF", "XAG", "XAU", "XBA", "XBB", "XBC", "XBD",
	// "XCD", "XDR", "XOF", "XPD", "XPF", "XPT", "XSU", "XUA", "YER", "ZAR", "ZMW",
	// "ZWL".
	FromCurrencyCode CurrencyPairUpdateFromCurrencyCode `json:"fromCurrencyCode,omitzero,required"`
	// Any of "AED", "AFN", "ALL", "AMD", "ANG", "AOA", "ARS", "AUD", "AWG", "AZN",
	// "BAM", "BBD", "BDT", "BGN", "BHD", "BIF", "BMD", "BND", "BOB", "BOV", "BRL",
	// "BSD", "BTN", "BWP", "BYN", "BZD", "CAD", "CDF", "CHE", "CHF", "CHW", "CLF",
	// "CLP", "CNY", "COP", "COU", "CRC", "CUC", "CUP", "CVE", "CZK", "DJF", "DKK",
	// "DOP", "DZD", "EGP", "ERN", "ETB", "EUR", "FJD", "FKP", "GBP", "GEL", "GHS",
	// "GIP", "GMD", "GNF", "GTQ", "GYD", "HKD", "HNL", "HRK", "HTG", "HUF", "IDR",
	// "ILS", "INR", "IQD", "IRR", "ISK", "JMD", "JOD", "JPY", "KES", "KGS", "KHR",
	// "KMF", "KPW", "KRW", "KWD", "KYD", "KZT", "LAK", "LBP", "LKR", "LRD", "LSL",
	// "LYD", "MAD", "MDL", "MGA", "MKD", "MMK", "MNT", "MOP", "MRU", "MUR", "MVR",
	// "MWK", "MXN", "MXV", "MYR", "MZN", "NAD", "NGN", "NIO", "NOK", "NPR", "NZD",
	// "OMR", "PAB", "PEN", "PGK", "PHP", "PKR", "PLN", "PYG", "QAR", "RON", "RSD",
	// "RUB", "RWF", "SAR", "SBD", "SCR", "SDG", "SEK", "SGD", "SHP", "SLL", "SOS",
	// "SRD", "SSP", "STN", "SVC", "SYP", "SZL", "THB", "TJS", "TMT", "TND", "TOP",
	// "TRY", "TTD", "TWD", "TZS", "UAH", "UGX", "USD", "USN", "UYI", "UYU", "UZS",
	// "VEF", "VND", "VUV", "WST", "XAF", "XAG", "XAU", "XBA", "XBB", "XBC", "XBD",
	// "XCD", "XDR", "XOF", "XPD", "XPF", "XPT", "XSU", "XUA", "YER", "ZAR", "ZMW",
	// "ZWL".
	ToCurrencyCode CurrencyPairUpdateToCurrencyCode `json:"toCurrencyCode,omitzero,required"`
	VisibleInUi    bool                             `json:"visibleInUI,required"`
	paramObj
}

func (r CurrencyPairUpdateParam) MarshalJSON() (data []byte, err error) {
	type shadow CurrencyPairUpdateParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CurrencyPairUpdateParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CurrencyPairUpdateFromCurrencyCode string

const (
	CurrencyPairUpdateFromCurrencyCodeAed CurrencyPairUpdateFromCurrencyCode = "AED"
	CurrencyPairUpdateFromCurrencyCodeAfn CurrencyPairUpdateFromCurrencyCode = "AFN"
	CurrencyPairUpdateFromCurrencyCodeAll CurrencyPairUpdateFromCurrencyCode = "ALL"
	CurrencyPairUpdateFromCurrencyCodeAmd CurrencyPairUpdateFromCurrencyCode = "AMD"
	CurrencyPairUpdateFromCurrencyCodeAng CurrencyPairUpdateFromCurrencyCode = "ANG"
	CurrencyPairUpdateFromCurrencyCodeAoa CurrencyPairUpdateFromCurrencyCode = "AOA"
	CurrencyPairUpdateFromCurrencyCodeArs CurrencyPairUpdateFromCurrencyCode = "ARS"
	CurrencyPairUpdateFromCurrencyCodeAud CurrencyPairUpdateFromCurrencyCode = "AUD"
	CurrencyPairUpdateFromCurrencyCodeAwg CurrencyPairUpdateFromCurrencyCode = "AWG"
	CurrencyPairUpdateFromCurrencyCodeAzn CurrencyPairUpdateFromCurrencyCode = "AZN"
	CurrencyPairUpdateFromCurrencyCodeBam CurrencyPairUpdateFromCurrencyCode = "BAM"
	CurrencyPairUpdateFromCurrencyCodeBbd CurrencyPairUpdateFromCurrencyCode = "BBD"
	CurrencyPairUpdateFromCurrencyCodeBdt CurrencyPairUpdateFromCurrencyCode = "BDT"
	CurrencyPairUpdateFromCurrencyCodeBgn CurrencyPairUpdateFromCurrencyCode = "BGN"
	CurrencyPairUpdateFromCurrencyCodeBhd CurrencyPairUpdateFromCurrencyCode = "BHD"
	CurrencyPairUpdateFromCurrencyCodeBif CurrencyPairUpdateFromCurrencyCode = "BIF"
	CurrencyPairUpdateFromCurrencyCodeBmd CurrencyPairUpdateFromCurrencyCode = "BMD"
	CurrencyPairUpdateFromCurrencyCodeBnd CurrencyPairUpdateFromCurrencyCode = "BND"
	CurrencyPairUpdateFromCurrencyCodeBob CurrencyPairUpdateFromCurrencyCode = "BOB"
	CurrencyPairUpdateFromCurrencyCodeBov CurrencyPairUpdateFromCurrencyCode = "BOV"
	CurrencyPairUpdateFromCurrencyCodeBrl CurrencyPairUpdateFromCurrencyCode = "BRL"
	CurrencyPairUpdateFromCurrencyCodeBsd CurrencyPairUpdateFromCurrencyCode = "BSD"
	CurrencyPairUpdateFromCurrencyCodeBtn CurrencyPairUpdateFromCurrencyCode = "BTN"
	CurrencyPairUpdateFromCurrencyCodeBwp CurrencyPairUpdateFromCurrencyCode = "BWP"
	CurrencyPairUpdateFromCurrencyCodeByn CurrencyPairUpdateFromCurrencyCode = "BYN"
	CurrencyPairUpdateFromCurrencyCodeBzd CurrencyPairUpdateFromCurrencyCode = "BZD"
	CurrencyPairUpdateFromCurrencyCodeCad CurrencyPairUpdateFromCurrencyCode = "CAD"
	CurrencyPairUpdateFromCurrencyCodeCdf CurrencyPairUpdateFromCurrencyCode = "CDF"
	CurrencyPairUpdateFromCurrencyCodeChe CurrencyPairUpdateFromCurrencyCode = "CHE"
	CurrencyPairUpdateFromCurrencyCodeChf CurrencyPairUpdateFromCurrencyCode = "CHF"
	CurrencyPairUpdateFromCurrencyCodeChw CurrencyPairUpdateFromCurrencyCode = "CHW"
	CurrencyPairUpdateFromCurrencyCodeClf CurrencyPairUpdateFromCurrencyCode = "CLF"
	CurrencyPairUpdateFromCurrencyCodeClp CurrencyPairUpdateFromCurrencyCode = "CLP"
	CurrencyPairUpdateFromCurrencyCodeCny CurrencyPairUpdateFromCurrencyCode = "CNY"
	CurrencyPairUpdateFromCurrencyCodeCop CurrencyPairUpdateFromCurrencyCode = "COP"
	CurrencyPairUpdateFromCurrencyCodeCou CurrencyPairUpdateFromCurrencyCode = "COU"
	CurrencyPairUpdateFromCurrencyCodeCrc CurrencyPairUpdateFromCurrencyCode = "CRC"
	CurrencyPairUpdateFromCurrencyCodeCuc CurrencyPairUpdateFromCurrencyCode = "CUC"
	CurrencyPairUpdateFromCurrencyCodeCup CurrencyPairUpdateFromCurrencyCode = "CUP"
	CurrencyPairUpdateFromCurrencyCodeCve CurrencyPairUpdateFromCurrencyCode = "CVE"
	CurrencyPairUpdateFromCurrencyCodeCzk CurrencyPairUpdateFromCurrencyCode = "CZK"
	CurrencyPairUpdateFromCurrencyCodeDjf CurrencyPairUpdateFromCurrencyCode = "DJF"
	CurrencyPairUpdateFromCurrencyCodeDkk CurrencyPairUpdateFromCurrencyCode = "DKK"
	CurrencyPairUpdateFromCurrencyCodeDop CurrencyPairUpdateFromCurrencyCode = "DOP"
	CurrencyPairUpdateFromCurrencyCodeDzd CurrencyPairUpdateFromCurrencyCode = "DZD"
	CurrencyPairUpdateFromCurrencyCodeEgp CurrencyPairUpdateFromCurrencyCode = "EGP"
	CurrencyPairUpdateFromCurrencyCodeErn CurrencyPairUpdateFromCurrencyCode = "ERN"
	CurrencyPairUpdateFromCurrencyCodeEtb CurrencyPairUpdateFromCurrencyCode = "ETB"
	CurrencyPairUpdateFromCurrencyCodeEur CurrencyPairUpdateFromCurrencyCode = "EUR"
	CurrencyPairUpdateFromCurrencyCodeFjd CurrencyPairUpdateFromCurrencyCode = "FJD"
	CurrencyPairUpdateFromCurrencyCodeFkp CurrencyPairUpdateFromCurrencyCode = "FKP"
	CurrencyPairUpdateFromCurrencyCodeGbp CurrencyPairUpdateFromCurrencyCode = "GBP"
	CurrencyPairUpdateFromCurrencyCodeGel CurrencyPairUpdateFromCurrencyCode = "GEL"
	CurrencyPairUpdateFromCurrencyCodeGhs CurrencyPairUpdateFromCurrencyCode = "GHS"
	CurrencyPairUpdateFromCurrencyCodeGip CurrencyPairUpdateFromCurrencyCode = "GIP"
	CurrencyPairUpdateFromCurrencyCodeGmd CurrencyPairUpdateFromCurrencyCode = "GMD"
	CurrencyPairUpdateFromCurrencyCodeGnf CurrencyPairUpdateFromCurrencyCode = "GNF"
	CurrencyPairUpdateFromCurrencyCodeGtq CurrencyPairUpdateFromCurrencyCode = "GTQ"
	CurrencyPairUpdateFromCurrencyCodeGyd CurrencyPairUpdateFromCurrencyCode = "GYD"
	CurrencyPairUpdateFromCurrencyCodeHkd CurrencyPairUpdateFromCurrencyCode = "HKD"
	CurrencyPairUpdateFromCurrencyCodeHnl CurrencyPairUpdateFromCurrencyCode = "HNL"
	CurrencyPairUpdateFromCurrencyCodeHrk CurrencyPairUpdateFromCurrencyCode = "HRK"
	CurrencyPairUpdateFromCurrencyCodeHtg CurrencyPairUpdateFromCurrencyCode = "HTG"
	CurrencyPairUpdateFromCurrencyCodeHuf CurrencyPairUpdateFromCurrencyCode = "HUF"
	CurrencyPairUpdateFromCurrencyCodeIdr CurrencyPairUpdateFromCurrencyCode = "IDR"
	CurrencyPairUpdateFromCurrencyCodeIls CurrencyPairUpdateFromCurrencyCode = "ILS"
	CurrencyPairUpdateFromCurrencyCodeInr CurrencyPairUpdateFromCurrencyCode = "INR"
	CurrencyPairUpdateFromCurrencyCodeIqd CurrencyPairUpdateFromCurrencyCode = "IQD"
	CurrencyPairUpdateFromCurrencyCodeIrr CurrencyPairUpdateFromCurrencyCode = "IRR"
	CurrencyPairUpdateFromCurrencyCodeIsk CurrencyPairUpdateFromCurrencyCode = "ISK"
	CurrencyPairUpdateFromCurrencyCodeJmd CurrencyPairUpdateFromCurrencyCode = "JMD"
	CurrencyPairUpdateFromCurrencyCodeJod CurrencyPairUpdateFromCurrencyCode = "JOD"
	CurrencyPairUpdateFromCurrencyCodeJpy CurrencyPairUpdateFromCurrencyCode = "JPY"
	CurrencyPairUpdateFromCurrencyCodeKes CurrencyPairUpdateFromCurrencyCode = "KES"
	CurrencyPairUpdateFromCurrencyCodeKgs CurrencyPairUpdateFromCurrencyCode = "KGS"
	CurrencyPairUpdateFromCurrencyCodeKhr CurrencyPairUpdateFromCurrencyCode = "KHR"
	CurrencyPairUpdateFromCurrencyCodeKmf CurrencyPairUpdateFromCurrencyCode = "KMF"
	CurrencyPairUpdateFromCurrencyCodeKpw CurrencyPairUpdateFromCurrencyCode = "KPW"
	CurrencyPairUpdateFromCurrencyCodeKrw CurrencyPairUpdateFromCurrencyCode = "KRW"
	CurrencyPairUpdateFromCurrencyCodeKwd CurrencyPairUpdateFromCurrencyCode = "KWD"
	CurrencyPairUpdateFromCurrencyCodeKyd CurrencyPairUpdateFromCurrencyCode = "KYD"
	CurrencyPairUpdateFromCurrencyCodeKzt CurrencyPairUpdateFromCurrencyCode = "KZT"
	CurrencyPairUpdateFromCurrencyCodeLak CurrencyPairUpdateFromCurrencyCode = "LAK"
	CurrencyPairUpdateFromCurrencyCodeLbp CurrencyPairUpdateFromCurrencyCode = "LBP"
	CurrencyPairUpdateFromCurrencyCodeLkr CurrencyPairUpdateFromCurrencyCode = "LKR"
	CurrencyPairUpdateFromCurrencyCodeLrd CurrencyPairUpdateFromCurrencyCode = "LRD"
	CurrencyPairUpdateFromCurrencyCodeLsl CurrencyPairUpdateFromCurrencyCode = "LSL"
	CurrencyPairUpdateFromCurrencyCodeLyd CurrencyPairUpdateFromCurrencyCode = "LYD"
	CurrencyPairUpdateFromCurrencyCodeMad CurrencyPairUpdateFromCurrencyCode = "MAD"
	CurrencyPairUpdateFromCurrencyCodeMdl CurrencyPairUpdateFromCurrencyCode = "MDL"
	CurrencyPairUpdateFromCurrencyCodeMga CurrencyPairUpdateFromCurrencyCode = "MGA"
	CurrencyPairUpdateFromCurrencyCodeMkd CurrencyPairUpdateFromCurrencyCode = "MKD"
	CurrencyPairUpdateFromCurrencyCodeMmk CurrencyPairUpdateFromCurrencyCode = "MMK"
	CurrencyPairUpdateFromCurrencyCodeMnt CurrencyPairUpdateFromCurrencyCode = "MNT"
	CurrencyPairUpdateFromCurrencyCodeMop CurrencyPairUpdateFromCurrencyCode = "MOP"
	CurrencyPairUpdateFromCurrencyCodeMru CurrencyPairUpdateFromCurrencyCode = "MRU"
	CurrencyPairUpdateFromCurrencyCodeMur CurrencyPairUpdateFromCurrencyCode = "MUR"
	CurrencyPairUpdateFromCurrencyCodeMvr CurrencyPairUpdateFromCurrencyCode = "MVR"
	CurrencyPairUpdateFromCurrencyCodeMwk CurrencyPairUpdateFromCurrencyCode = "MWK"
	CurrencyPairUpdateFromCurrencyCodeMxn CurrencyPairUpdateFromCurrencyCode = "MXN"
	CurrencyPairUpdateFromCurrencyCodeMxv CurrencyPairUpdateFromCurrencyCode = "MXV"
	CurrencyPairUpdateFromCurrencyCodeMyr CurrencyPairUpdateFromCurrencyCode = "MYR"
	CurrencyPairUpdateFromCurrencyCodeMzn CurrencyPairUpdateFromCurrencyCode = "MZN"
	CurrencyPairUpdateFromCurrencyCodeNad CurrencyPairUpdateFromCurrencyCode = "NAD"
	CurrencyPairUpdateFromCurrencyCodeNgn CurrencyPairUpdateFromCurrencyCode = "NGN"
	CurrencyPairUpdateFromCurrencyCodeNio CurrencyPairUpdateFromCurrencyCode = "NIO"
	CurrencyPairUpdateFromCurrencyCodeNok CurrencyPairUpdateFromCurrencyCode = "NOK"
	CurrencyPairUpdateFromCurrencyCodeNpr CurrencyPairUpdateFromCurrencyCode = "NPR"
	CurrencyPairUpdateFromCurrencyCodeNzd CurrencyPairUpdateFromCurrencyCode = "NZD"
	CurrencyPairUpdateFromCurrencyCodeOmr CurrencyPairUpdateFromCurrencyCode = "OMR"
	CurrencyPairUpdateFromCurrencyCodePab CurrencyPairUpdateFromCurrencyCode = "PAB"
	CurrencyPairUpdateFromCurrencyCodePen CurrencyPairUpdateFromCurrencyCode = "PEN"
	CurrencyPairUpdateFromCurrencyCodePgk CurrencyPairUpdateFromCurrencyCode = "PGK"
	CurrencyPairUpdateFromCurrencyCodePhp CurrencyPairUpdateFromCurrencyCode = "PHP"
	CurrencyPairUpdateFromCurrencyCodePkr CurrencyPairUpdateFromCurrencyCode = "PKR"
	CurrencyPairUpdateFromCurrencyCodePln CurrencyPairUpdateFromCurrencyCode = "PLN"
	CurrencyPairUpdateFromCurrencyCodePyg CurrencyPairUpdateFromCurrencyCode = "PYG"
	CurrencyPairUpdateFromCurrencyCodeQar CurrencyPairUpdateFromCurrencyCode = "QAR"
	CurrencyPairUpdateFromCurrencyCodeRon CurrencyPairUpdateFromCurrencyCode = "RON"
	CurrencyPairUpdateFromCurrencyCodeRsd CurrencyPairUpdateFromCurrencyCode = "RSD"
	CurrencyPairUpdateFromCurrencyCodeRub CurrencyPairUpdateFromCurrencyCode = "RUB"
	CurrencyPairUpdateFromCurrencyCodeRwf CurrencyPairUpdateFromCurrencyCode = "RWF"
	CurrencyPairUpdateFromCurrencyCodeSar CurrencyPairUpdateFromCurrencyCode = "SAR"
	CurrencyPairUpdateFromCurrencyCodeSbd CurrencyPairUpdateFromCurrencyCode = "SBD"
	CurrencyPairUpdateFromCurrencyCodeScr CurrencyPairUpdateFromCurrencyCode = "SCR"
	CurrencyPairUpdateFromCurrencyCodeSdg CurrencyPairUpdateFromCurrencyCode = "SDG"
	CurrencyPairUpdateFromCurrencyCodeSek CurrencyPairUpdateFromCurrencyCode = "SEK"
	CurrencyPairUpdateFromCurrencyCodeSgd CurrencyPairUpdateFromCurrencyCode = "SGD"
	CurrencyPairUpdateFromCurrencyCodeShp CurrencyPairUpdateFromCurrencyCode = "SHP"
	CurrencyPairUpdateFromCurrencyCodeSll CurrencyPairUpdateFromCurrencyCode = "SLL"
	CurrencyPairUpdateFromCurrencyCodeSos CurrencyPairUpdateFromCurrencyCode = "SOS"
	CurrencyPairUpdateFromCurrencyCodeSrd CurrencyPairUpdateFromCurrencyCode = "SRD"
	CurrencyPairUpdateFromCurrencyCodeSsp CurrencyPairUpdateFromCurrencyCode = "SSP"
	CurrencyPairUpdateFromCurrencyCodeStn CurrencyPairUpdateFromCurrencyCode = "STN"
	CurrencyPairUpdateFromCurrencyCodeSvc CurrencyPairUpdateFromCurrencyCode = "SVC"
	CurrencyPairUpdateFromCurrencyCodeSyp CurrencyPairUpdateFromCurrencyCode = "SYP"
	CurrencyPairUpdateFromCurrencyCodeSzl CurrencyPairUpdateFromCurrencyCode = "SZL"
	CurrencyPairUpdateFromCurrencyCodeThb CurrencyPairUpdateFromCurrencyCode = "THB"
	CurrencyPairUpdateFromCurrencyCodeTjs CurrencyPairUpdateFromCurrencyCode = "TJS"
	CurrencyPairUpdateFromCurrencyCodeTmt CurrencyPairUpdateFromCurrencyCode = "TMT"
	CurrencyPairUpdateFromCurrencyCodeTnd CurrencyPairUpdateFromCurrencyCode = "TND"
	CurrencyPairUpdateFromCurrencyCodeTop CurrencyPairUpdateFromCurrencyCode = "TOP"
	CurrencyPairUpdateFromCurrencyCodeTry CurrencyPairUpdateFromCurrencyCode = "TRY"
	CurrencyPairUpdateFromCurrencyCodeTtd CurrencyPairUpdateFromCurrencyCode = "TTD"
	CurrencyPairUpdateFromCurrencyCodeTwd CurrencyPairUpdateFromCurrencyCode = "TWD"
	CurrencyPairUpdateFromCurrencyCodeTzs CurrencyPairUpdateFromCurrencyCode = "TZS"
	CurrencyPairUpdateFromCurrencyCodeUah CurrencyPairUpdateFromCurrencyCode = "UAH"
	CurrencyPairUpdateFromCurrencyCodeUgx CurrencyPairUpdateFromCurrencyCode = "UGX"
	CurrencyPairUpdateFromCurrencyCodeUsd CurrencyPairUpdateFromCurrencyCode = "USD"
	CurrencyPairUpdateFromCurrencyCodeUsn CurrencyPairUpdateFromCurrencyCode = "USN"
	CurrencyPairUpdateFromCurrencyCodeUyi CurrencyPairUpdateFromCurrencyCode = "UYI"
	CurrencyPairUpdateFromCurrencyCodeUyu CurrencyPairUpdateFromCurrencyCode = "UYU"
	CurrencyPairUpdateFromCurrencyCodeUzs CurrencyPairUpdateFromCurrencyCode = "UZS"
	CurrencyPairUpdateFromCurrencyCodeVef CurrencyPairUpdateFromCurrencyCode = "VEF"
	CurrencyPairUpdateFromCurrencyCodeVnd CurrencyPairUpdateFromCurrencyCode = "VND"
	CurrencyPairUpdateFromCurrencyCodeVuv CurrencyPairUpdateFromCurrencyCode = "VUV"
	CurrencyPairUpdateFromCurrencyCodeWst CurrencyPairUpdateFromCurrencyCode = "WST"
	CurrencyPairUpdateFromCurrencyCodeXaf CurrencyPairUpdateFromCurrencyCode = "XAF"
	CurrencyPairUpdateFromCurrencyCodeXag CurrencyPairUpdateFromCurrencyCode = "XAG"
	CurrencyPairUpdateFromCurrencyCodeXau CurrencyPairUpdateFromCurrencyCode = "XAU"
	CurrencyPairUpdateFromCurrencyCodeXba CurrencyPairUpdateFromCurrencyCode = "XBA"
	CurrencyPairUpdateFromCurrencyCodeXbb CurrencyPairUpdateFromCurrencyCode = "XBB"
	CurrencyPairUpdateFromCurrencyCodeXbc CurrencyPairUpdateFromCurrencyCode = "XBC"
	CurrencyPairUpdateFromCurrencyCodeXbd CurrencyPairUpdateFromCurrencyCode = "XBD"
	CurrencyPairUpdateFromCurrencyCodeXcd CurrencyPairUpdateFromCurrencyCode = "XCD"
	CurrencyPairUpdateFromCurrencyCodeXdr CurrencyPairUpdateFromCurrencyCode = "XDR"
	CurrencyPairUpdateFromCurrencyCodeXof CurrencyPairUpdateFromCurrencyCode = "XOF"
	CurrencyPairUpdateFromCurrencyCodeXpd CurrencyPairUpdateFromCurrencyCode = "XPD"
	CurrencyPairUpdateFromCurrencyCodeXpf CurrencyPairUpdateFromCurrencyCode = "XPF"
	CurrencyPairUpdateFromCurrencyCodeXpt CurrencyPairUpdateFromCurrencyCode = "XPT"
	CurrencyPairUpdateFromCurrencyCodeXsu CurrencyPairUpdateFromCurrencyCode = "XSU"
	CurrencyPairUpdateFromCurrencyCodeXua CurrencyPairUpdateFromCurrencyCode = "XUA"
	CurrencyPairUpdateFromCurrencyCodeYer CurrencyPairUpdateFromCurrencyCode = "YER"
	CurrencyPairUpdateFromCurrencyCodeZar CurrencyPairUpdateFromCurrencyCode = "ZAR"
	CurrencyPairUpdateFromCurrencyCodeZmw CurrencyPairUpdateFromCurrencyCode = "ZMW"
	CurrencyPairUpdateFromCurrencyCodeZwl CurrencyPairUpdateFromCurrencyCode = "ZWL"
)

type CurrencyPairUpdateToCurrencyCode string

const (
	CurrencyPairUpdateToCurrencyCodeAed CurrencyPairUpdateToCurrencyCode = "AED"
	CurrencyPairUpdateToCurrencyCodeAfn CurrencyPairUpdateToCurrencyCode = "AFN"
	CurrencyPairUpdateToCurrencyCodeAll CurrencyPairUpdateToCurrencyCode = "ALL"
	CurrencyPairUpdateToCurrencyCodeAmd CurrencyPairUpdateToCurrencyCode = "AMD"
	CurrencyPairUpdateToCurrencyCodeAng CurrencyPairUpdateToCurrencyCode = "ANG"
	CurrencyPairUpdateToCurrencyCodeAoa CurrencyPairUpdateToCurrencyCode = "AOA"
	CurrencyPairUpdateToCurrencyCodeArs CurrencyPairUpdateToCurrencyCode = "ARS"
	CurrencyPairUpdateToCurrencyCodeAud CurrencyPairUpdateToCurrencyCode = "AUD"
	CurrencyPairUpdateToCurrencyCodeAwg CurrencyPairUpdateToCurrencyCode = "AWG"
	CurrencyPairUpdateToCurrencyCodeAzn CurrencyPairUpdateToCurrencyCode = "AZN"
	CurrencyPairUpdateToCurrencyCodeBam CurrencyPairUpdateToCurrencyCode = "BAM"
	CurrencyPairUpdateToCurrencyCodeBbd CurrencyPairUpdateToCurrencyCode = "BBD"
	CurrencyPairUpdateToCurrencyCodeBdt CurrencyPairUpdateToCurrencyCode = "BDT"
	CurrencyPairUpdateToCurrencyCodeBgn CurrencyPairUpdateToCurrencyCode = "BGN"
	CurrencyPairUpdateToCurrencyCodeBhd CurrencyPairUpdateToCurrencyCode = "BHD"
	CurrencyPairUpdateToCurrencyCodeBif CurrencyPairUpdateToCurrencyCode = "BIF"
	CurrencyPairUpdateToCurrencyCodeBmd CurrencyPairUpdateToCurrencyCode = "BMD"
	CurrencyPairUpdateToCurrencyCodeBnd CurrencyPairUpdateToCurrencyCode = "BND"
	CurrencyPairUpdateToCurrencyCodeBob CurrencyPairUpdateToCurrencyCode = "BOB"
	CurrencyPairUpdateToCurrencyCodeBov CurrencyPairUpdateToCurrencyCode = "BOV"
	CurrencyPairUpdateToCurrencyCodeBrl CurrencyPairUpdateToCurrencyCode = "BRL"
	CurrencyPairUpdateToCurrencyCodeBsd CurrencyPairUpdateToCurrencyCode = "BSD"
	CurrencyPairUpdateToCurrencyCodeBtn CurrencyPairUpdateToCurrencyCode = "BTN"
	CurrencyPairUpdateToCurrencyCodeBwp CurrencyPairUpdateToCurrencyCode = "BWP"
	CurrencyPairUpdateToCurrencyCodeByn CurrencyPairUpdateToCurrencyCode = "BYN"
	CurrencyPairUpdateToCurrencyCodeBzd CurrencyPairUpdateToCurrencyCode = "BZD"
	CurrencyPairUpdateToCurrencyCodeCad CurrencyPairUpdateToCurrencyCode = "CAD"
	CurrencyPairUpdateToCurrencyCodeCdf CurrencyPairUpdateToCurrencyCode = "CDF"
	CurrencyPairUpdateToCurrencyCodeChe CurrencyPairUpdateToCurrencyCode = "CHE"
	CurrencyPairUpdateToCurrencyCodeChf CurrencyPairUpdateToCurrencyCode = "CHF"
	CurrencyPairUpdateToCurrencyCodeChw CurrencyPairUpdateToCurrencyCode = "CHW"
	CurrencyPairUpdateToCurrencyCodeClf CurrencyPairUpdateToCurrencyCode = "CLF"
	CurrencyPairUpdateToCurrencyCodeClp CurrencyPairUpdateToCurrencyCode = "CLP"
	CurrencyPairUpdateToCurrencyCodeCny CurrencyPairUpdateToCurrencyCode = "CNY"
	CurrencyPairUpdateToCurrencyCodeCop CurrencyPairUpdateToCurrencyCode = "COP"
	CurrencyPairUpdateToCurrencyCodeCou CurrencyPairUpdateToCurrencyCode = "COU"
	CurrencyPairUpdateToCurrencyCodeCrc CurrencyPairUpdateToCurrencyCode = "CRC"
	CurrencyPairUpdateToCurrencyCodeCuc CurrencyPairUpdateToCurrencyCode = "CUC"
	CurrencyPairUpdateToCurrencyCodeCup CurrencyPairUpdateToCurrencyCode = "CUP"
	CurrencyPairUpdateToCurrencyCodeCve CurrencyPairUpdateToCurrencyCode = "CVE"
	CurrencyPairUpdateToCurrencyCodeCzk CurrencyPairUpdateToCurrencyCode = "CZK"
	CurrencyPairUpdateToCurrencyCodeDjf CurrencyPairUpdateToCurrencyCode = "DJF"
	CurrencyPairUpdateToCurrencyCodeDkk CurrencyPairUpdateToCurrencyCode = "DKK"
	CurrencyPairUpdateToCurrencyCodeDop CurrencyPairUpdateToCurrencyCode = "DOP"
	CurrencyPairUpdateToCurrencyCodeDzd CurrencyPairUpdateToCurrencyCode = "DZD"
	CurrencyPairUpdateToCurrencyCodeEgp CurrencyPairUpdateToCurrencyCode = "EGP"
	CurrencyPairUpdateToCurrencyCodeErn CurrencyPairUpdateToCurrencyCode = "ERN"
	CurrencyPairUpdateToCurrencyCodeEtb CurrencyPairUpdateToCurrencyCode = "ETB"
	CurrencyPairUpdateToCurrencyCodeEur CurrencyPairUpdateToCurrencyCode = "EUR"
	CurrencyPairUpdateToCurrencyCodeFjd CurrencyPairUpdateToCurrencyCode = "FJD"
	CurrencyPairUpdateToCurrencyCodeFkp CurrencyPairUpdateToCurrencyCode = "FKP"
	CurrencyPairUpdateToCurrencyCodeGbp CurrencyPairUpdateToCurrencyCode = "GBP"
	CurrencyPairUpdateToCurrencyCodeGel CurrencyPairUpdateToCurrencyCode = "GEL"
	CurrencyPairUpdateToCurrencyCodeGhs CurrencyPairUpdateToCurrencyCode = "GHS"
	CurrencyPairUpdateToCurrencyCodeGip CurrencyPairUpdateToCurrencyCode = "GIP"
	CurrencyPairUpdateToCurrencyCodeGmd CurrencyPairUpdateToCurrencyCode = "GMD"
	CurrencyPairUpdateToCurrencyCodeGnf CurrencyPairUpdateToCurrencyCode = "GNF"
	CurrencyPairUpdateToCurrencyCodeGtq CurrencyPairUpdateToCurrencyCode = "GTQ"
	CurrencyPairUpdateToCurrencyCodeGyd CurrencyPairUpdateToCurrencyCode = "GYD"
	CurrencyPairUpdateToCurrencyCodeHkd CurrencyPairUpdateToCurrencyCode = "HKD"
	CurrencyPairUpdateToCurrencyCodeHnl CurrencyPairUpdateToCurrencyCode = "HNL"
	CurrencyPairUpdateToCurrencyCodeHrk CurrencyPairUpdateToCurrencyCode = "HRK"
	CurrencyPairUpdateToCurrencyCodeHtg CurrencyPairUpdateToCurrencyCode = "HTG"
	CurrencyPairUpdateToCurrencyCodeHuf CurrencyPairUpdateToCurrencyCode = "HUF"
	CurrencyPairUpdateToCurrencyCodeIdr CurrencyPairUpdateToCurrencyCode = "IDR"
	CurrencyPairUpdateToCurrencyCodeIls CurrencyPairUpdateToCurrencyCode = "ILS"
	CurrencyPairUpdateToCurrencyCodeInr CurrencyPairUpdateToCurrencyCode = "INR"
	CurrencyPairUpdateToCurrencyCodeIqd CurrencyPairUpdateToCurrencyCode = "IQD"
	CurrencyPairUpdateToCurrencyCodeIrr CurrencyPairUpdateToCurrencyCode = "IRR"
	CurrencyPairUpdateToCurrencyCodeIsk CurrencyPairUpdateToCurrencyCode = "ISK"
	CurrencyPairUpdateToCurrencyCodeJmd CurrencyPairUpdateToCurrencyCode = "JMD"
	CurrencyPairUpdateToCurrencyCodeJod CurrencyPairUpdateToCurrencyCode = "JOD"
	CurrencyPairUpdateToCurrencyCodeJpy CurrencyPairUpdateToCurrencyCode = "JPY"
	CurrencyPairUpdateToCurrencyCodeKes CurrencyPairUpdateToCurrencyCode = "KES"
	CurrencyPairUpdateToCurrencyCodeKgs CurrencyPairUpdateToCurrencyCode = "KGS"
	CurrencyPairUpdateToCurrencyCodeKhr CurrencyPairUpdateToCurrencyCode = "KHR"
	CurrencyPairUpdateToCurrencyCodeKmf CurrencyPairUpdateToCurrencyCode = "KMF"
	CurrencyPairUpdateToCurrencyCodeKpw CurrencyPairUpdateToCurrencyCode = "KPW"
	CurrencyPairUpdateToCurrencyCodeKrw CurrencyPairUpdateToCurrencyCode = "KRW"
	CurrencyPairUpdateToCurrencyCodeKwd CurrencyPairUpdateToCurrencyCode = "KWD"
	CurrencyPairUpdateToCurrencyCodeKyd CurrencyPairUpdateToCurrencyCode = "KYD"
	CurrencyPairUpdateToCurrencyCodeKzt CurrencyPairUpdateToCurrencyCode = "KZT"
	CurrencyPairUpdateToCurrencyCodeLak CurrencyPairUpdateToCurrencyCode = "LAK"
	CurrencyPairUpdateToCurrencyCodeLbp CurrencyPairUpdateToCurrencyCode = "LBP"
	CurrencyPairUpdateToCurrencyCodeLkr CurrencyPairUpdateToCurrencyCode = "LKR"
	CurrencyPairUpdateToCurrencyCodeLrd CurrencyPairUpdateToCurrencyCode = "LRD"
	CurrencyPairUpdateToCurrencyCodeLsl CurrencyPairUpdateToCurrencyCode = "LSL"
	CurrencyPairUpdateToCurrencyCodeLyd CurrencyPairUpdateToCurrencyCode = "LYD"
	CurrencyPairUpdateToCurrencyCodeMad CurrencyPairUpdateToCurrencyCode = "MAD"
	CurrencyPairUpdateToCurrencyCodeMdl CurrencyPairUpdateToCurrencyCode = "MDL"
	CurrencyPairUpdateToCurrencyCodeMga CurrencyPairUpdateToCurrencyCode = "MGA"
	CurrencyPairUpdateToCurrencyCodeMkd CurrencyPairUpdateToCurrencyCode = "MKD"
	CurrencyPairUpdateToCurrencyCodeMmk CurrencyPairUpdateToCurrencyCode = "MMK"
	CurrencyPairUpdateToCurrencyCodeMnt CurrencyPairUpdateToCurrencyCode = "MNT"
	CurrencyPairUpdateToCurrencyCodeMop CurrencyPairUpdateToCurrencyCode = "MOP"
	CurrencyPairUpdateToCurrencyCodeMru CurrencyPairUpdateToCurrencyCode = "MRU"
	CurrencyPairUpdateToCurrencyCodeMur CurrencyPairUpdateToCurrencyCode = "MUR"
	CurrencyPairUpdateToCurrencyCodeMvr CurrencyPairUpdateToCurrencyCode = "MVR"
	CurrencyPairUpdateToCurrencyCodeMwk CurrencyPairUpdateToCurrencyCode = "MWK"
	CurrencyPairUpdateToCurrencyCodeMxn CurrencyPairUpdateToCurrencyCode = "MXN"
	CurrencyPairUpdateToCurrencyCodeMxv CurrencyPairUpdateToCurrencyCode = "MXV"
	CurrencyPairUpdateToCurrencyCodeMyr CurrencyPairUpdateToCurrencyCode = "MYR"
	CurrencyPairUpdateToCurrencyCodeMzn CurrencyPairUpdateToCurrencyCode = "MZN"
	CurrencyPairUpdateToCurrencyCodeNad CurrencyPairUpdateToCurrencyCode = "NAD"
	CurrencyPairUpdateToCurrencyCodeNgn CurrencyPairUpdateToCurrencyCode = "NGN"
	CurrencyPairUpdateToCurrencyCodeNio CurrencyPairUpdateToCurrencyCode = "NIO"
	CurrencyPairUpdateToCurrencyCodeNok CurrencyPairUpdateToCurrencyCode = "NOK"
	CurrencyPairUpdateToCurrencyCodeNpr CurrencyPairUpdateToCurrencyCode = "NPR"
	CurrencyPairUpdateToCurrencyCodeNzd CurrencyPairUpdateToCurrencyCode = "NZD"
	CurrencyPairUpdateToCurrencyCodeOmr CurrencyPairUpdateToCurrencyCode = "OMR"
	CurrencyPairUpdateToCurrencyCodePab CurrencyPairUpdateToCurrencyCode = "PAB"
	CurrencyPairUpdateToCurrencyCodePen CurrencyPairUpdateToCurrencyCode = "PEN"
	CurrencyPairUpdateToCurrencyCodePgk CurrencyPairUpdateToCurrencyCode = "PGK"
	CurrencyPairUpdateToCurrencyCodePhp CurrencyPairUpdateToCurrencyCode = "PHP"
	CurrencyPairUpdateToCurrencyCodePkr CurrencyPairUpdateToCurrencyCode = "PKR"
	CurrencyPairUpdateToCurrencyCodePln CurrencyPairUpdateToCurrencyCode = "PLN"
	CurrencyPairUpdateToCurrencyCodePyg CurrencyPairUpdateToCurrencyCode = "PYG"
	CurrencyPairUpdateToCurrencyCodeQar CurrencyPairUpdateToCurrencyCode = "QAR"
	CurrencyPairUpdateToCurrencyCodeRon CurrencyPairUpdateToCurrencyCode = "RON"
	CurrencyPairUpdateToCurrencyCodeRsd CurrencyPairUpdateToCurrencyCode = "RSD"
	CurrencyPairUpdateToCurrencyCodeRub CurrencyPairUpdateToCurrencyCode = "RUB"
	CurrencyPairUpdateToCurrencyCodeRwf CurrencyPairUpdateToCurrencyCode = "RWF"
	CurrencyPairUpdateToCurrencyCodeSar CurrencyPairUpdateToCurrencyCode = "SAR"
	CurrencyPairUpdateToCurrencyCodeSbd CurrencyPairUpdateToCurrencyCode = "SBD"
	CurrencyPairUpdateToCurrencyCodeScr CurrencyPairUpdateToCurrencyCode = "SCR"
	CurrencyPairUpdateToCurrencyCodeSdg CurrencyPairUpdateToCurrencyCode = "SDG"
	CurrencyPairUpdateToCurrencyCodeSek CurrencyPairUpdateToCurrencyCode = "SEK"
	CurrencyPairUpdateToCurrencyCodeSgd CurrencyPairUpdateToCurrencyCode = "SGD"
	CurrencyPairUpdateToCurrencyCodeShp CurrencyPairUpdateToCurrencyCode = "SHP"
	CurrencyPairUpdateToCurrencyCodeSll CurrencyPairUpdateToCurrencyCode = "SLL"
	CurrencyPairUpdateToCurrencyCodeSos CurrencyPairUpdateToCurrencyCode = "SOS"
	CurrencyPairUpdateToCurrencyCodeSrd CurrencyPairUpdateToCurrencyCode = "SRD"
	CurrencyPairUpdateToCurrencyCodeSsp CurrencyPairUpdateToCurrencyCode = "SSP"
	CurrencyPairUpdateToCurrencyCodeStn CurrencyPairUpdateToCurrencyCode = "STN"
	CurrencyPairUpdateToCurrencyCodeSvc CurrencyPairUpdateToCurrencyCode = "SVC"
	CurrencyPairUpdateToCurrencyCodeSyp CurrencyPairUpdateToCurrencyCode = "SYP"
	CurrencyPairUpdateToCurrencyCodeSzl CurrencyPairUpdateToCurrencyCode = "SZL"
	CurrencyPairUpdateToCurrencyCodeThb CurrencyPairUpdateToCurrencyCode = "THB"
	CurrencyPairUpdateToCurrencyCodeTjs CurrencyPairUpdateToCurrencyCode = "TJS"
	CurrencyPairUpdateToCurrencyCodeTmt CurrencyPairUpdateToCurrencyCode = "TMT"
	CurrencyPairUpdateToCurrencyCodeTnd CurrencyPairUpdateToCurrencyCode = "TND"
	CurrencyPairUpdateToCurrencyCodeTop CurrencyPairUpdateToCurrencyCode = "TOP"
	CurrencyPairUpdateToCurrencyCodeTry CurrencyPairUpdateToCurrencyCode = "TRY"
	CurrencyPairUpdateToCurrencyCodeTtd CurrencyPairUpdateToCurrencyCode = "TTD"
	CurrencyPairUpdateToCurrencyCodeTwd CurrencyPairUpdateToCurrencyCode = "TWD"
	CurrencyPairUpdateToCurrencyCodeTzs CurrencyPairUpdateToCurrencyCode = "TZS"
	CurrencyPairUpdateToCurrencyCodeUah CurrencyPairUpdateToCurrencyCode = "UAH"
	CurrencyPairUpdateToCurrencyCodeUgx CurrencyPairUpdateToCurrencyCode = "UGX"
	CurrencyPairUpdateToCurrencyCodeUsd CurrencyPairUpdateToCurrencyCode = "USD"
	CurrencyPairUpdateToCurrencyCodeUsn CurrencyPairUpdateToCurrencyCode = "USN"
	CurrencyPairUpdateToCurrencyCodeUyi CurrencyPairUpdateToCurrencyCode = "UYI"
	CurrencyPairUpdateToCurrencyCodeUyu CurrencyPairUpdateToCurrencyCode = "UYU"
	CurrencyPairUpdateToCurrencyCodeUzs CurrencyPairUpdateToCurrencyCode = "UZS"
	CurrencyPairUpdateToCurrencyCodeVef CurrencyPairUpdateToCurrencyCode = "VEF"
	CurrencyPairUpdateToCurrencyCodeVnd CurrencyPairUpdateToCurrencyCode = "VND"
	CurrencyPairUpdateToCurrencyCodeVuv CurrencyPairUpdateToCurrencyCode = "VUV"
	CurrencyPairUpdateToCurrencyCodeWst CurrencyPairUpdateToCurrencyCode = "WST"
	CurrencyPairUpdateToCurrencyCodeXaf CurrencyPairUpdateToCurrencyCode = "XAF"
	CurrencyPairUpdateToCurrencyCodeXag CurrencyPairUpdateToCurrencyCode = "XAG"
	CurrencyPairUpdateToCurrencyCodeXau CurrencyPairUpdateToCurrencyCode = "XAU"
	CurrencyPairUpdateToCurrencyCodeXba CurrencyPairUpdateToCurrencyCode = "XBA"
	CurrencyPairUpdateToCurrencyCodeXbb CurrencyPairUpdateToCurrencyCode = "XBB"
	CurrencyPairUpdateToCurrencyCodeXbc CurrencyPairUpdateToCurrencyCode = "XBC"
	CurrencyPairUpdateToCurrencyCodeXbd CurrencyPairUpdateToCurrencyCode = "XBD"
	CurrencyPairUpdateToCurrencyCodeXcd CurrencyPairUpdateToCurrencyCode = "XCD"
	CurrencyPairUpdateToCurrencyCodeXdr CurrencyPairUpdateToCurrencyCode = "XDR"
	CurrencyPairUpdateToCurrencyCodeXof CurrencyPairUpdateToCurrencyCode = "XOF"
	CurrencyPairUpdateToCurrencyCodeXpd CurrencyPairUpdateToCurrencyCode = "XPD"
	CurrencyPairUpdateToCurrencyCodeXpf CurrencyPairUpdateToCurrencyCode = "XPF"
	CurrencyPairUpdateToCurrencyCodeXpt CurrencyPairUpdateToCurrencyCode = "XPT"
	CurrencyPairUpdateToCurrencyCodeXsu CurrencyPairUpdateToCurrencyCode = "XSU"
	CurrencyPairUpdateToCurrencyCodeXua CurrencyPairUpdateToCurrencyCode = "XUA"
	CurrencyPairUpdateToCurrencyCodeYer CurrencyPairUpdateToCurrencyCode = "YER"
	CurrencyPairUpdateToCurrencyCodeZar CurrencyPairUpdateToCurrencyCode = "ZAR"
	CurrencyPairUpdateToCurrencyCodeZmw CurrencyPairUpdateToCurrencyCode = "ZMW"
	CurrencyPairUpdateToCurrencyCodeZwl CurrencyPairUpdateToCurrencyCode = "ZWL"
)

type ExchangeRate struct {
	ID             string    `json:"id,required"`
	ConversionRate float64   `json:"conversionRate,required"`
	CreatedAt      time.Time `json:"createdAt,required" format:"date-time"`
	EffectiveAt    time.Time `json:"effectiveAt,required" format:"date-time"`
	// Any of "AED", "AFN", "ALL", "AMD", "ANG", "AOA", "ARS", "AUD", "AWG", "AZN",
	// "BAM", "BBD", "BDT", "BGN", "BHD", "BIF", "BMD", "BND", "BOB", "BOV", "BRL",
	// "BSD", "BTN", "BWP", "BYN", "BZD", "CAD", "CDF", "CHE", "CHF", "CHW", "CLF",
	// "CLP", "CNY", "COP", "COU", "CRC", "CUC", "CUP", "CVE", "CZK", "DJF", "DKK",
	// "DOP", "DZD", "EGP", "ERN", "ETB", "EUR", "FJD", "FKP", "GBP", "GEL", "GHS",
	// "GIP", "GMD", "GNF", "GTQ", "GYD", "HKD", "HNL", "HRK", "HTG", "HUF", "IDR",
	// "ILS", "INR", "IQD", "IRR", "ISK", "JMD", "JOD", "JPY", "KES", "KGS", "KHR",
	// "KMF", "KPW", "KRW", "KWD", "KYD", "KZT", "LAK", "LBP", "LKR", "LRD", "LSL",
	// "LYD", "MAD", "MDL", "MGA", "MKD", "MMK", "MNT", "MOP", "MRU", "MUR", "MVR",
	// "MWK", "MXN", "MXV", "MYR", "MZN", "NAD", "NGN", "NIO", "NOK", "NPR", "NZD",
	// "OMR", "PAB", "PEN", "PGK", "PHP", "PKR", "PLN", "PYG", "QAR", "RON", "RSD",
	// "RUB", "RWF", "SAR", "SBD", "SCR", "SDG", "SEK", "SGD", "SHP", "SLL", "SOS",
	// "SRD", "SSP", "STN", "SVC", "SYP", "SZL", "THB", "TJS", "TMT", "TND", "TOP",
	// "TRY", "TTD", "TWD", "TZS", "UAH", "UGX", "USD", "USN", "UYI", "UYU", "UZS",
	// "VEF", "VND", "VUV", "WST", "XAF", "XAG", "XAU", "XBA", "XBB", "XBC", "XBD",
	// "XCD", "XDR", "XOF", "XPD", "XPF", "XPT", "XSU", "XUA", "YER", "ZAR", "ZMW",
	// "ZWL".
	FromCurrencyCode ExchangeRateFromCurrencyCode `json:"fromCurrencyCode,required"`
	// Any of "AED", "AFN", "ALL", "AMD", "ANG", "AOA", "ARS", "AUD", "AWG", "AZN",
	// "BAM", "BBD", "BDT", "BGN", "BHD", "BIF", "BMD", "BND", "BOB", "BOV", "BRL",
	// "BSD", "BTN", "BWP", "BYN", "BZD", "CAD", "CDF", "CHE", "CHF", "CHW", "CLF",
	// "CLP", "CNY", "COP", "COU", "CRC", "CUC", "CUP", "CVE", "CZK", "DJF", "DKK",
	// "DOP", "DZD", "EGP", "ERN", "ETB", "EUR", "FJD", "FKP", "GBP", "GEL", "GHS",
	// "GIP", "GMD", "GNF", "GTQ", "GYD", "HKD", "HNL", "HRK", "HTG", "HUF", "IDR",
	// "ILS", "INR", "IQD", "IRR", "ISK", "JMD", "JOD", "JPY", "KES", "KGS", "KHR",
	// "KMF", "KPW", "KRW", "KWD", "KYD", "KZT", "LAK", "LBP", "LKR", "LRD", "LSL",
	// "LYD", "MAD", "MDL", "MGA", "MKD", "MMK", "MNT", "MOP", "MRU", "MUR", "MVR",
	// "MWK", "MXN", "MXV", "MYR", "MZN", "NAD", "NGN", "NIO", "NOK", "NPR", "NZD",
	// "OMR", "PAB", "PEN", "PGK", "PHP", "PKR", "PLN", "PYG", "QAR", "RON", "RSD",
	// "RUB", "RWF", "SAR", "SBD", "SCR", "SDG", "SEK", "SGD", "SHP", "SLL", "SOS",
	// "SRD", "SSP", "STN", "SVC", "SYP", "SZL", "THB", "TJS", "TMT", "TND", "TOP",
	// "TRY", "TTD", "TWD", "TZS", "UAH", "UGX", "USD", "USN", "UYI", "UYU", "UZS",
	// "VEF", "VND", "VUV", "WST", "XAF", "XAG", "XAU", "XBA", "XBB", "XBC", "XBD",
	// "XCD", "XDR", "XOF", "XPD", "XPF", "XPT", "XSU", "XUA", "YER", "ZAR", "ZMW",
	// "ZWL".
	ToCurrencyCode ExchangeRateToCurrencyCode `json:"toCurrencyCode,required"`
	UpdatedAt      time.Time                  `json:"updatedAt,required" format:"date-time"`
	VisibleInUi    bool                       `json:"visibleInUI,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		ConversionRate   respjson.Field
		CreatedAt        respjson.Field
		EffectiveAt      respjson.Field
		FromCurrencyCode respjson.Field
		ToCurrencyCode   respjson.Field
		UpdatedAt        respjson.Field
		VisibleInUi      respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExchangeRate) RawJSON() string { return r.JSON.raw }
func (r *ExchangeRate) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExchangeRateFromCurrencyCode string

const (
	ExchangeRateFromCurrencyCodeAed ExchangeRateFromCurrencyCode = "AED"
	ExchangeRateFromCurrencyCodeAfn ExchangeRateFromCurrencyCode = "AFN"
	ExchangeRateFromCurrencyCodeAll ExchangeRateFromCurrencyCode = "ALL"
	ExchangeRateFromCurrencyCodeAmd ExchangeRateFromCurrencyCode = "AMD"
	ExchangeRateFromCurrencyCodeAng ExchangeRateFromCurrencyCode = "ANG"
	ExchangeRateFromCurrencyCodeAoa ExchangeRateFromCurrencyCode = "AOA"
	ExchangeRateFromCurrencyCodeArs ExchangeRateFromCurrencyCode = "ARS"
	ExchangeRateFromCurrencyCodeAud ExchangeRateFromCurrencyCode = "AUD"
	ExchangeRateFromCurrencyCodeAwg ExchangeRateFromCurrencyCode = "AWG"
	ExchangeRateFromCurrencyCodeAzn ExchangeRateFromCurrencyCode = "AZN"
	ExchangeRateFromCurrencyCodeBam ExchangeRateFromCurrencyCode = "BAM"
	ExchangeRateFromCurrencyCodeBbd ExchangeRateFromCurrencyCode = "BBD"
	ExchangeRateFromCurrencyCodeBdt ExchangeRateFromCurrencyCode = "BDT"
	ExchangeRateFromCurrencyCodeBgn ExchangeRateFromCurrencyCode = "BGN"
	ExchangeRateFromCurrencyCodeBhd ExchangeRateFromCurrencyCode = "BHD"
	ExchangeRateFromCurrencyCodeBif ExchangeRateFromCurrencyCode = "BIF"
	ExchangeRateFromCurrencyCodeBmd ExchangeRateFromCurrencyCode = "BMD"
	ExchangeRateFromCurrencyCodeBnd ExchangeRateFromCurrencyCode = "BND"
	ExchangeRateFromCurrencyCodeBob ExchangeRateFromCurrencyCode = "BOB"
	ExchangeRateFromCurrencyCodeBov ExchangeRateFromCurrencyCode = "BOV"
	ExchangeRateFromCurrencyCodeBrl ExchangeRateFromCurrencyCode = "BRL"
	ExchangeRateFromCurrencyCodeBsd ExchangeRateFromCurrencyCode = "BSD"
	ExchangeRateFromCurrencyCodeBtn ExchangeRateFromCurrencyCode = "BTN"
	ExchangeRateFromCurrencyCodeBwp ExchangeRateFromCurrencyCode = "BWP"
	ExchangeRateFromCurrencyCodeByn ExchangeRateFromCurrencyCode = "BYN"
	ExchangeRateFromCurrencyCodeBzd ExchangeRateFromCurrencyCode = "BZD"
	ExchangeRateFromCurrencyCodeCad ExchangeRateFromCurrencyCode = "CAD"
	ExchangeRateFromCurrencyCodeCdf ExchangeRateFromCurrencyCode = "CDF"
	ExchangeRateFromCurrencyCodeChe ExchangeRateFromCurrencyCode = "CHE"
	ExchangeRateFromCurrencyCodeChf ExchangeRateFromCurrencyCode = "CHF"
	ExchangeRateFromCurrencyCodeChw ExchangeRateFromCurrencyCode = "CHW"
	ExchangeRateFromCurrencyCodeClf ExchangeRateFromCurrencyCode = "CLF"
	ExchangeRateFromCurrencyCodeClp ExchangeRateFromCurrencyCode = "CLP"
	ExchangeRateFromCurrencyCodeCny ExchangeRateFromCurrencyCode = "CNY"
	ExchangeRateFromCurrencyCodeCop ExchangeRateFromCurrencyCode = "COP"
	ExchangeRateFromCurrencyCodeCou ExchangeRateFromCurrencyCode = "COU"
	ExchangeRateFromCurrencyCodeCrc ExchangeRateFromCurrencyCode = "CRC"
	ExchangeRateFromCurrencyCodeCuc ExchangeRateFromCurrencyCode = "CUC"
	ExchangeRateFromCurrencyCodeCup ExchangeRateFromCurrencyCode = "CUP"
	ExchangeRateFromCurrencyCodeCve ExchangeRateFromCurrencyCode = "CVE"
	ExchangeRateFromCurrencyCodeCzk ExchangeRateFromCurrencyCode = "CZK"
	ExchangeRateFromCurrencyCodeDjf ExchangeRateFromCurrencyCode = "DJF"
	ExchangeRateFromCurrencyCodeDkk ExchangeRateFromCurrencyCode = "DKK"
	ExchangeRateFromCurrencyCodeDop ExchangeRateFromCurrencyCode = "DOP"
	ExchangeRateFromCurrencyCodeDzd ExchangeRateFromCurrencyCode = "DZD"
	ExchangeRateFromCurrencyCodeEgp ExchangeRateFromCurrencyCode = "EGP"
	ExchangeRateFromCurrencyCodeErn ExchangeRateFromCurrencyCode = "ERN"
	ExchangeRateFromCurrencyCodeEtb ExchangeRateFromCurrencyCode = "ETB"
	ExchangeRateFromCurrencyCodeEur ExchangeRateFromCurrencyCode = "EUR"
	ExchangeRateFromCurrencyCodeFjd ExchangeRateFromCurrencyCode = "FJD"
	ExchangeRateFromCurrencyCodeFkp ExchangeRateFromCurrencyCode = "FKP"
	ExchangeRateFromCurrencyCodeGbp ExchangeRateFromCurrencyCode = "GBP"
	ExchangeRateFromCurrencyCodeGel ExchangeRateFromCurrencyCode = "GEL"
	ExchangeRateFromCurrencyCodeGhs ExchangeRateFromCurrencyCode = "GHS"
	ExchangeRateFromCurrencyCodeGip ExchangeRateFromCurrencyCode = "GIP"
	ExchangeRateFromCurrencyCodeGmd ExchangeRateFromCurrencyCode = "GMD"
	ExchangeRateFromCurrencyCodeGnf ExchangeRateFromCurrencyCode = "GNF"
	ExchangeRateFromCurrencyCodeGtq ExchangeRateFromCurrencyCode = "GTQ"
	ExchangeRateFromCurrencyCodeGyd ExchangeRateFromCurrencyCode = "GYD"
	ExchangeRateFromCurrencyCodeHkd ExchangeRateFromCurrencyCode = "HKD"
	ExchangeRateFromCurrencyCodeHnl ExchangeRateFromCurrencyCode = "HNL"
	ExchangeRateFromCurrencyCodeHrk ExchangeRateFromCurrencyCode = "HRK"
	ExchangeRateFromCurrencyCodeHtg ExchangeRateFromCurrencyCode = "HTG"
	ExchangeRateFromCurrencyCodeHuf ExchangeRateFromCurrencyCode = "HUF"
	ExchangeRateFromCurrencyCodeIdr ExchangeRateFromCurrencyCode = "IDR"
	ExchangeRateFromCurrencyCodeIls ExchangeRateFromCurrencyCode = "ILS"
	ExchangeRateFromCurrencyCodeInr ExchangeRateFromCurrencyCode = "INR"
	ExchangeRateFromCurrencyCodeIqd ExchangeRateFromCurrencyCode = "IQD"
	ExchangeRateFromCurrencyCodeIrr ExchangeRateFromCurrencyCode = "IRR"
	ExchangeRateFromCurrencyCodeIsk ExchangeRateFromCurrencyCode = "ISK"
	ExchangeRateFromCurrencyCodeJmd ExchangeRateFromCurrencyCode = "JMD"
	ExchangeRateFromCurrencyCodeJod ExchangeRateFromCurrencyCode = "JOD"
	ExchangeRateFromCurrencyCodeJpy ExchangeRateFromCurrencyCode = "JPY"
	ExchangeRateFromCurrencyCodeKes ExchangeRateFromCurrencyCode = "KES"
	ExchangeRateFromCurrencyCodeKgs ExchangeRateFromCurrencyCode = "KGS"
	ExchangeRateFromCurrencyCodeKhr ExchangeRateFromCurrencyCode = "KHR"
	ExchangeRateFromCurrencyCodeKmf ExchangeRateFromCurrencyCode = "KMF"
	ExchangeRateFromCurrencyCodeKpw ExchangeRateFromCurrencyCode = "KPW"
	ExchangeRateFromCurrencyCodeKrw ExchangeRateFromCurrencyCode = "KRW"
	ExchangeRateFromCurrencyCodeKwd ExchangeRateFromCurrencyCode = "KWD"
	ExchangeRateFromCurrencyCodeKyd ExchangeRateFromCurrencyCode = "KYD"
	ExchangeRateFromCurrencyCodeKzt ExchangeRateFromCurrencyCode = "KZT"
	ExchangeRateFromCurrencyCodeLak ExchangeRateFromCurrencyCode = "LAK"
	ExchangeRateFromCurrencyCodeLbp ExchangeRateFromCurrencyCode = "LBP"
	ExchangeRateFromCurrencyCodeLkr ExchangeRateFromCurrencyCode = "LKR"
	ExchangeRateFromCurrencyCodeLrd ExchangeRateFromCurrencyCode = "LRD"
	ExchangeRateFromCurrencyCodeLsl ExchangeRateFromCurrencyCode = "LSL"
	ExchangeRateFromCurrencyCodeLyd ExchangeRateFromCurrencyCode = "LYD"
	ExchangeRateFromCurrencyCodeMad ExchangeRateFromCurrencyCode = "MAD"
	ExchangeRateFromCurrencyCodeMdl ExchangeRateFromCurrencyCode = "MDL"
	ExchangeRateFromCurrencyCodeMga ExchangeRateFromCurrencyCode = "MGA"
	ExchangeRateFromCurrencyCodeMkd ExchangeRateFromCurrencyCode = "MKD"
	ExchangeRateFromCurrencyCodeMmk ExchangeRateFromCurrencyCode = "MMK"
	ExchangeRateFromCurrencyCodeMnt ExchangeRateFromCurrencyCode = "MNT"
	ExchangeRateFromCurrencyCodeMop ExchangeRateFromCurrencyCode = "MOP"
	ExchangeRateFromCurrencyCodeMru ExchangeRateFromCurrencyCode = "MRU"
	ExchangeRateFromCurrencyCodeMur ExchangeRateFromCurrencyCode = "MUR"
	ExchangeRateFromCurrencyCodeMvr ExchangeRateFromCurrencyCode = "MVR"
	ExchangeRateFromCurrencyCodeMwk ExchangeRateFromCurrencyCode = "MWK"
	ExchangeRateFromCurrencyCodeMxn ExchangeRateFromCurrencyCode = "MXN"
	ExchangeRateFromCurrencyCodeMxv ExchangeRateFromCurrencyCode = "MXV"
	ExchangeRateFromCurrencyCodeMyr ExchangeRateFromCurrencyCode = "MYR"
	ExchangeRateFromCurrencyCodeMzn ExchangeRateFromCurrencyCode = "MZN"
	ExchangeRateFromCurrencyCodeNad ExchangeRateFromCurrencyCode = "NAD"
	ExchangeRateFromCurrencyCodeNgn ExchangeRateFromCurrencyCode = "NGN"
	ExchangeRateFromCurrencyCodeNio ExchangeRateFromCurrencyCode = "NIO"
	ExchangeRateFromCurrencyCodeNok ExchangeRateFromCurrencyCode = "NOK"
	ExchangeRateFromCurrencyCodeNpr ExchangeRateFromCurrencyCode = "NPR"
	ExchangeRateFromCurrencyCodeNzd ExchangeRateFromCurrencyCode = "NZD"
	ExchangeRateFromCurrencyCodeOmr ExchangeRateFromCurrencyCode = "OMR"
	ExchangeRateFromCurrencyCodePab ExchangeRateFromCurrencyCode = "PAB"
	ExchangeRateFromCurrencyCodePen ExchangeRateFromCurrencyCode = "PEN"
	ExchangeRateFromCurrencyCodePgk ExchangeRateFromCurrencyCode = "PGK"
	ExchangeRateFromCurrencyCodePhp ExchangeRateFromCurrencyCode = "PHP"
	ExchangeRateFromCurrencyCodePkr ExchangeRateFromCurrencyCode = "PKR"
	ExchangeRateFromCurrencyCodePln ExchangeRateFromCurrencyCode = "PLN"
	ExchangeRateFromCurrencyCodePyg ExchangeRateFromCurrencyCode = "PYG"
	ExchangeRateFromCurrencyCodeQar ExchangeRateFromCurrencyCode = "QAR"
	ExchangeRateFromCurrencyCodeRon ExchangeRateFromCurrencyCode = "RON"
	ExchangeRateFromCurrencyCodeRsd ExchangeRateFromCurrencyCode = "RSD"
	ExchangeRateFromCurrencyCodeRub ExchangeRateFromCurrencyCode = "RUB"
	ExchangeRateFromCurrencyCodeRwf ExchangeRateFromCurrencyCode = "RWF"
	ExchangeRateFromCurrencyCodeSar ExchangeRateFromCurrencyCode = "SAR"
	ExchangeRateFromCurrencyCodeSbd ExchangeRateFromCurrencyCode = "SBD"
	ExchangeRateFromCurrencyCodeScr ExchangeRateFromCurrencyCode = "SCR"
	ExchangeRateFromCurrencyCodeSdg ExchangeRateFromCurrencyCode = "SDG"
	ExchangeRateFromCurrencyCodeSek ExchangeRateFromCurrencyCode = "SEK"
	ExchangeRateFromCurrencyCodeSgd ExchangeRateFromCurrencyCode = "SGD"
	ExchangeRateFromCurrencyCodeShp ExchangeRateFromCurrencyCode = "SHP"
	ExchangeRateFromCurrencyCodeSll ExchangeRateFromCurrencyCode = "SLL"
	ExchangeRateFromCurrencyCodeSos ExchangeRateFromCurrencyCode = "SOS"
	ExchangeRateFromCurrencyCodeSrd ExchangeRateFromCurrencyCode = "SRD"
	ExchangeRateFromCurrencyCodeSsp ExchangeRateFromCurrencyCode = "SSP"
	ExchangeRateFromCurrencyCodeStn ExchangeRateFromCurrencyCode = "STN"
	ExchangeRateFromCurrencyCodeSvc ExchangeRateFromCurrencyCode = "SVC"
	ExchangeRateFromCurrencyCodeSyp ExchangeRateFromCurrencyCode = "SYP"
	ExchangeRateFromCurrencyCodeSzl ExchangeRateFromCurrencyCode = "SZL"
	ExchangeRateFromCurrencyCodeThb ExchangeRateFromCurrencyCode = "THB"
	ExchangeRateFromCurrencyCodeTjs ExchangeRateFromCurrencyCode = "TJS"
	ExchangeRateFromCurrencyCodeTmt ExchangeRateFromCurrencyCode = "TMT"
	ExchangeRateFromCurrencyCodeTnd ExchangeRateFromCurrencyCode = "TND"
	ExchangeRateFromCurrencyCodeTop ExchangeRateFromCurrencyCode = "TOP"
	ExchangeRateFromCurrencyCodeTry ExchangeRateFromCurrencyCode = "TRY"
	ExchangeRateFromCurrencyCodeTtd ExchangeRateFromCurrencyCode = "TTD"
	ExchangeRateFromCurrencyCodeTwd ExchangeRateFromCurrencyCode = "TWD"
	ExchangeRateFromCurrencyCodeTzs ExchangeRateFromCurrencyCode = "TZS"
	ExchangeRateFromCurrencyCodeUah ExchangeRateFromCurrencyCode = "UAH"
	ExchangeRateFromCurrencyCodeUgx ExchangeRateFromCurrencyCode = "UGX"
	ExchangeRateFromCurrencyCodeUsd ExchangeRateFromCurrencyCode = "USD"
	ExchangeRateFromCurrencyCodeUsn ExchangeRateFromCurrencyCode = "USN"
	ExchangeRateFromCurrencyCodeUyi ExchangeRateFromCurrencyCode = "UYI"
	ExchangeRateFromCurrencyCodeUyu ExchangeRateFromCurrencyCode = "UYU"
	ExchangeRateFromCurrencyCodeUzs ExchangeRateFromCurrencyCode = "UZS"
	ExchangeRateFromCurrencyCodeVef ExchangeRateFromCurrencyCode = "VEF"
	ExchangeRateFromCurrencyCodeVnd ExchangeRateFromCurrencyCode = "VND"
	ExchangeRateFromCurrencyCodeVuv ExchangeRateFromCurrencyCode = "VUV"
	ExchangeRateFromCurrencyCodeWst ExchangeRateFromCurrencyCode = "WST"
	ExchangeRateFromCurrencyCodeXaf ExchangeRateFromCurrencyCode = "XAF"
	ExchangeRateFromCurrencyCodeXag ExchangeRateFromCurrencyCode = "XAG"
	ExchangeRateFromCurrencyCodeXau ExchangeRateFromCurrencyCode = "XAU"
	ExchangeRateFromCurrencyCodeXba ExchangeRateFromCurrencyCode = "XBA"
	ExchangeRateFromCurrencyCodeXbb ExchangeRateFromCurrencyCode = "XBB"
	ExchangeRateFromCurrencyCodeXbc ExchangeRateFromCurrencyCode = "XBC"
	ExchangeRateFromCurrencyCodeXbd ExchangeRateFromCurrencyCode = "XBD"
	ExchangeRateFromCurrencyCodeXcd ExchangeRateFromCurrencyCode = "XCD"
	ExchangeRateFromCurrencyCodeXdr ExchangeRateFromCurrencyCode = "XDR"
	ExchangeRateFromCurrencyCodeXof ExchangeRateFromCurrencyCode = "XOF"
	ExchangeRateFromCurrencyCodeXpd ExchangeRateFromCurrencyCode = "XPD"
	ExchangeRateFromCurrencyCodeXpf ExchangeRateFromCurrencyCode = "XPF"
	ExchangeRateFromCurrencyCodeXpt ExchangeRateFromCurrencyCode = "XPT"
	ExchangeRateFromCurrencyCodeXsu ExchangeRateFromCurrencyCode = "XSU"
	ExchangeRateFromCurrencyCodeXua ExchangeRateFromCurrencyCode = "XUA"
	ExchangeRateFromCurrencyCodeYer ExchangeRateFromCurrencyCode = "YER"
	ExchangeRateFromCurrencyCodeZar ExchangeRateFromCurrencyCode = "ZAR"
	ExchangeRateFromCurrencyCodeZmw ExchangeRateFromCurrencyCode = "ZMW"
	ExchangeRateFromCurrencyCodeZwl ExchangeRateFromCurrencyCode = "ZWL"
)

type ExchangeRateToCurrencyCode string

const (
	ExchangeRateToCurrencyCodeAed ExchangeRateToCurrencyCode = "AED"
	ExchangeRateToCurrencyCodeAfn ExchangeRateToCurrencyCode = "AFN"
	ExchangeRateToCurrencyCodeAll ExchangeRateToCurrencyCode = "ALL"
	ExchangeRateToCurrencyCodeAmd ExchangeRateToCurrencyCode = "AMD"
	ExchangeRateToCurrencyCodeAng ExchangeRateToCurrencyCode = "ANG"
	ExchangeRateToCurrencyCodeAoa ExchangeRateToCurrencyCode = "AOA"
	ExchangeRateToCurrencyCodeArs ExchangeRateToCurrencyCode = "ARS"
	ExchangeRateToCurrencyCodeAud ExchangeRateToCurrencyCode = "AUD"
	ExchangeRateToCurrencyCodeAwg ExchangeRateToCurrencyCode = "AWG"
	ExchangeRateToCurrencyCodeAzn ExchangeRateToCurrencyCode = "AZN"
	ExchangeRateToCurrencyCodeBam ExchangeRateToCurrencyCode = "BAM"
	ExchangeRateToCurrencyCodeBbd ExchangeRateToCurrencyCode = "BBD"
	ExchangeRateToCurrencyCodeBdt ExchangeRateToCurrencyCode = "BDT"
	ExchangeRateToCurrencyCodeBgn ExchangeRateToCurrencyCode = "BGN"
	ExchangeRateToCurrencyCodeBhd ExchangeRateToCurrencyCode = "BHD"
	ExchangeRateToCurrencyCodeBif ExchangeRateToCurrencyCode = "BIF"
	ExchangeRateToCurrencyCodeBmd ExchangeRateToCurrencyCode = "BMD"
	ExchangeRateToCurrencyCodeBnd ExchangeRateToCurrencyCode = "BND"
	ExchangeRateToCurrencyCodeBob ExchangeRateToCurrencyCode = "BOB"
	ExchangeRateToCurrencyCodeBov ExchangeRateToCurrencyCode = "BOV"
	ExchangeRateToCurrencyCodeBrl ExchangeRateToCurrencyCode = "BRL"
	ExchangeRateToCurrencyCodeBsd ExchangeRateToCurrencyCode = "BSD"
	ExchangeRateToCurrencyCodeBtn ExchangeRateToCurrencyCode = "BTN"
	ExchangeRateToCurrencyCodeBwp ExchangeRateToCurrencyCode = "BWP"
	ExchangeRateToCurrencyCodeByn ExchangeRateToCurrencyCode = "BYN"
	ExchangeRateToCurrencyCodeBzd ExchangeRateToCurrencyCode = "BZD"
	ExchangeRateToCurrencyCodeCad ExchangeRateToCurrencyCode = "CAD"
	ExchangeRateToCurrencyCodeCdf ExchangeRateToCurrencyCode = "CDF"
	ExchangeRateToCurrencyCodeChe ExchangeRateToCurrencyCode = "CHE"
	ExchangeRateToCurrencyCodeChf ExchangeRateToCurrencyCode = "CHF"
	ExchangeRateToCurrencyCodeChw ExchangeRateToCurrencyCode = "CHW"
	ExchangeRateToCurrencyCodeClf ExchangeRateToCurrencyCode = "CLF"
	ExchangeRateToCurrencyCodeClp ExchangeRateToCurrencyCode = "CLP"
	ExchangeRateToCurrencyCodeCny ExchangeRateToCurrencyCode = "CNY"
	ExchangeRateToCurrencyCodeCop ExchangeRateToCurrencyCode = "COP"
	ExchangeRateToCurrencyCodeCou ExchangeRateToCurrencyCode = "COU"
	ExchangeRateToCurrencyCodeCrc ExchangeRateToCurrencyCode = "CRC"
	ExchangeRateToCurrencyCodeCuc ExchangeRateToCurrencyCode = "CUC"
	ExchangeRateToCurrencyCodeCup ExchangeRateToCurrencyCode = "CUP"
	ExchangeRateToCurrencyCodeCve ExchangeRateToCurrencyCode = "CVE"
	ExchangeRateToCurrencyCodeCzk ExchangeRateToCurrencyCode = "CZK"
	ExchangeRateToCurrencyCodeDjf ExchangeRateToCurrencyCode = "DJF"
	ExchangeRateToCurrencyCodeDkk ExchangeRateToCurrencyCode = "DKK"
	ExchangeRateToCurrencyCodeDop ExchangeRateToCurrencyCode = "DOP"
	ExchangeRateToCurrencyCodeDzd ExchangeRateToCurrencyCode = "DZD"
	ExchangeRateToCurrencyCodeEgp ExchangeRateToCurrencyCode = "EGP"
	ExchangeRateToCurrencyCodeErn ExchangeRateToCurrencyCode = "ERN"
	ExchangeRateToCurrencyCodeEtb ExchangeRateToCurrencyCode = "ETB"
	ExchangeRateToCurrencyCodeEur ExchangeRateToCurrencyCode = "EUR"
	ExchangeRateToCurrencyCodeFjd ExchangeRateToCurrencyCode = "FJD"
	ExchangeRateToCurrencyCodeFkp ExchangeRateToCurrencyCode = "FKP"
	ExchangeRateToCurrencyCodeGbp ExchangeRateToCurrencyCode = "GBP"
	ExchangeRateToCurrencyCodeGel ExchangeRateToCurrencyCode = "GEL"
	ExchangeRateToCurrencyCodeGhs ExchangeRateToCurrencyCode = "GHS"
	ExchangeRateToCurrencyCodeGip ExchangeRateToCurrencyCode = "GIP"
	ExchangeRateToCurrencyCodeGmd ExchangeRateToCurrencyCode = "GMD"
	ExchangeRateToCurrencyCodeGnf ExchangeRateToCurrencyCode = "GNF"
	ExchangeRateToCurrencyCodeGtq ExchangeRateToCurrencyCode = "GTQ"
	ExchangeRateToCurrencyCodeGyd ExchangeRateToCurrencyCode = "GYD"
	ExchangeRateToCurrencyCodeHkd ExchangeRateToCurrencyCode = "HKD"
	ExchangeRateToCurrencyCodeHnl ExchangeRateToCurrencyCode = "HNL"
	ExchangeRateToCurrencyCodeHrk ExchangeRateToCurrencyCode = "HRK"
	ExchangeRateToCurrencyCodeHtg ExchangeRateToCurrencyCode = "HTG"
	ExchangeRateToCurrencyCodeHuf ExchangeRateToCurrencyCode = "HUF"
	ExchangeRateToCurrencyCodeIdr ExchangeRateToCurrencyCode = "IDR"
	ExchangeRateToCurrencyCodeIls ExchangeRateToCurrencyCode = "ILS"
	ExchangeRateToCurrencyCodeInr ExchangeRateToCurrencyCode = "INR"
	ExchangeRateToCurrencyCodeIqd ExchangeRateToCurrencyCode = "IQD"
	ExchangeRateToCurrencyCodeIrr ExchangeRateToCurrencyCode = "IRR"
	ExchangeRateToCurrencyCodeIsk ExchangeRateToCurrencyCode = "ISK"
	ExchangeRateToCurrencyCodeJmd ExchangeRateToCurrencyCode = "JMD"
	ExchangeRateToCurrencyCodeJod ExchangeRateToCurrencyCode = "JOD"
	ExchangeRateToCurrencyCodeJpy ExchangeRateToCurrencyCode = "JPY"
	ExchangeRateToCurrencyCodeKes ExchangeRateToCurrencyCode = "KES"
	ExchangeRateToCurrencyCodeKgs ExchangeRateToCurrencyCode = "KGS"
	ExchangeRateToCurrencyCodeKhr ExchangeRateToCurrencyCode = "KHR"
	ExchangeRateToCurrencyCodeKmf ExchangeRateToCurrencyCode = "KMF"
	ExchangeRateToCurrencyCodeKpw ExchangeRateToCurrencyCode = "KPW"
	ExchangeRateToCurrencyCodeKrw ExchangeRateToCurrencyCode = "KRW"
	ExchangeRateToCurrencyCodeKwd ExchangeRateToCurrencyCode = "KWD"
	ExchangeRateToCurrencyCodeKyd ExchangeRateToCurrencyCode = "KYD"
	ExchangeRateToCurrencyCodeKzt ExchangeRateToCurrencyCode = "KZT"
	ExchangeRateToCurrencyCodeLak ExchangeRateToCurrencyCode = "LAK"
	ExchangeRateToCurrencyCodeLbp ExchangeRateToCurrencyCode = "LBP"
	ExchangeRateToCurrencyCodeLkr ExchangeRateToCurrencyCode = "LKR"
	ExchangeRateToCurrencyCodeLrd ExchangeRateToCurrencyCode = "LRD"
	ExchangeRateToCurrencyCodeLsl ExchangeRateToCurrencyCode = "LSL"
	ExchangeRateToCurrencyCodeLyd ExchangeRateToCurrencyCode = "LYD"
	ExchangeRateToCurrencyCodeMad ExchangeRateToCurrencyCode = "MAD"
	ExchangeRateToCurrencyCodeMdl ExchangeRateToCurrencyCode = "MDL"
	ExchangeRateToCurrencyCodeMga ExchangeRateToCurrencyCode = "MGA"
	ExchangeRateToCurrencyCodeMkd ExchangeRateToCurrencyCode = "MKD"
	ExchangeRateToCurrencyCodeMmk ExchangeRateToCurrencyCode = "MMK"
	ExchangeRateToCurrencyCodeMnt ExchangeRateToCurrencyCode = "MNT"
	ExchangeRateToCurrencyCodeMop ExchangeRateToCurrencyCode = "MOP"
	ExchangeRateToCurrencyCodeMru ExchangeRateToCurrencyCode = "MRU"
	ExchangeRateToCurrencyCodeMur ExchangeRateToCurrencyCode = "MUR"
	ExchangeRateToCurrencyCodeMvr ExchangeRateToCurrencyCode = "MVR"
	ExchangeRateToCurrencyCodeMwk ExchangeRateToCurrencyCode = "MWK"
	ExchangeRateToCurrencyCodeMxn ExchangeRateToCurrencyCode = "MXN"
	ExchangeRateToCurrencyCodeMxv ExchangeRateToCurrencyCode = "MXV"
	ExchangeRateToCurrencyCodeMyr ExchangeRateToCurrencyCode = "MYR"
	ExchangeRateToCurrencyCodeMzn ExchangeRateToCurrencyCode = "MZN"
	ExchangeRateToCurrencyCodeNad ExchangeRateToCurrencyCode = "NAD"
	ExchangeRateToCurrencyCodeNgn ExchangeRateToCurrencyCode = "NGN"
	ExchangeRateToCurrencyCodeNio ExchangeRateToCurrencyCode = "NIO"
	ExchangeRateToCurrencyCodeNok ExchangeRateToCurrencyCode = "NOK"
	ExchangeRateToCurrencyCodeNpr ExchangeRateToCurrencyCode = "NPR"
	ExchangeRateToCurrencyCodeNzd ExchangeRateToCurrencyCode = "NZD"
	ExchangeRateToCurrencyCodeOmr ExchangeRateToCurrencyCode = "OMR"
	ExchangeRateToCurrencyCodePab ExchangeRateToCurrencyCode = "PAB"
	ExchangeRateToCurrencyCodePen ExchangeRateToCurrencyCode = "PEN"
	ExchangeRateToCurrencyCodePgk ExchangeRateToCurrencyCode = "PGK"
	ExchangeRateToCurrencyCodePhp ExchangeRateToCurrencyCode = "PHP"
	ExchangeRateToCurrencyCodePkr ExchangeRateToCurrencyCode = "PKR"
	ExchangeRateToCurrencyCodePln ExchangeRateToCurrencyCode = "PLN"
	ExchangeRateToCurrencyCodePyg ExchangeRateToCurrencyCode = "PYG"
	ExchangeRateToCurrencyCodeQar ExchangeRateToCurrencyCode = "QAR"
	ExchangeRateToCurrencyCodeRon ExchangeRateToCurrencyCode = "RON"
	ExchangeRateToCurrencyCodeRsd ExchangeRateToCurrencyCode = "RSD"
	ExchangeRateToCurrencyCodeRub ExchangeRateToCurrencyCode = "RUB"
	ExchangeRateToCurrencyCodeRwf ExchangeRateToCurrencyCode = "RWF"
	ExchangeRateToCurrencyCodeSar ExchangeRateToCurrencyCode = "SAR"
	ExchangeRateToCurrencyCodeSbd ExchangeRateToCurrencyCode = "SBD"
	ExchangeRateToCurrencyCodeScr ExchangeRateToCurrencyCode = "SCR"
	ExchangeRateToCurrencyCodeSdg ExchangeRateToCurrencyCode = "SDG"
	ExchangeRateToCurrencyCodeSek ExchangeRateToCurrencyCode = "SEK"
	ExchangeRateToCurrencyCodeSgd ExchangeRateToCurrencyCode = "SGD"
	ExchangeRateToCurrencyCodeShp ExchangeRateToCurrencyCode = "SHP"
	ExchangeRateToCurrencyCodeSll ExchangeRateToCurrencyCode = "SLL"
	ExchangeRateToCurrencyCodeSos ExchangeRateToCurrencyCode = "SOS"
	ExchangeRateToCurrencyCodeSrd ExchangeRateToCurrencyCode = "SRD"
	ExchangeRateToCurrencyCodeSsp ExchangeRateToCurrencyCode = "SSP"
	ExchangeRateToCurrencyCodeStn ExchangeRateToCurrencyCode = "STN"
	ExchangeRateToCurrencyCodeSvc ExchangeRateToCurrencyCode = "SVC"
	ExchangeRateToCurrencyCodeSyp ExchangeRateToCurrencyCode = "SYP"
	ExchangeRateToCurrencyCodeSzl ExchangeRateToCurrencyCode = "SZL"
	ExchangeRateToCurrencyCodeThb ExchangeRateToCurrencyCode = "THB"
	ExchangeRateToCurrencyCodeTjs ExchangeRateToCurrencyCode = "TJS"
	ExchangeRateToCurrencyCodeTmt ExchangeRateToCurrencyCode = "TMT"
	ExchangeRateToCurrencyCodeTnd ExchangeRateToCurrencyCode = "TND"
	ExchangeRateToCurrencyCodeTop ExchangeRateToCurrencyCode = "TOP"
	ExchangeRateToCurrencyCodeTry ExchangeRateToCurrencyCode = "TRY"
	ExchangeRateToCurrencyCodeTtd ExchangeRateToCurrencyCode = "TTD"
	ExchangeRateToCurrencyCodeTwd ExchangeRateToCurrencyCode = "TWD"
	ExchangeRateToCurrencyCodeTzs ExchangeRateToCurrencyCode = "TZS"
	ExchangeRateToCurrencyCodeUah ExchangeRateToCurrencyCode = "UAH"
	ExchangeRateToCurrencyCodeUgx ExchangeRateToCurrencyCode = "UGX"
	ExchangeRateToCurrencyCodeUsd ExchangeRateToCurrencyCode = "USD"
	ExchangeRateToCurrencyCodeUsn ExchangeRateToCurrencyCode = "USN"
	ExchangeRateToCurrencyCodeUyi ExchangeRateToCurrencyCode = "UYI"
	ExchangeRateToCurrencyCodeUyu ExchangeRateToCurrencyCode = "UYU"
	ExchangeRateToCurrencyCodeUzs ExchangeRateToCurrencyCode = "UZS"
	ExchangeRateToCurrencyCodeVef ExchangeRateToCurrencyCode = "VEF"
	ExchangeRateToCurrencyCodeVnd ExchangeRateToCurrencyCode = "VND"
	ExchangeRateToCurrencyCodeVuv ExchangeRateToCurrencyCode = "VUV"
	ExchangeRateToCurrencyCodeWst ExchangeRateToCurrencyCode = "WST"
	ExchangeRateToCurrencyCodeXaf ExchangeRateToCurrencyCode = "XAF"
	ExchangeRateToCurrencyCodeXag ExchangeRateToCurrencyCode = "XAG"
	ExchangeRateToCurrencyCodeXau ExchangeRateToCurrencyCode = "XAU"
	ExchangeRateToCurrencyCodeXba ExchangeRateToCurrencyCode = "XBA"
	ExchangeRateToCurrencyCodeXbb ExchangeRateToCurrencyCode = "XBB"
	ExchangeRateToCurrencyCodeXbc ExchangeRateToCurrencyCode = "XBC"
	ExchangeRateToCurrencyCodeXbd ExchangeRateToCurrencyCode = "XBD"
	ExchangeRateToCurrencyCodeXcd ExchangeRateToCurrencyCode = "XCD"
	ExchangeRateToCurrencyCodeXdr ExchangeRateToCurrencyCode = "XDR"
	ExchangeRateToCurrencyCodeXof ExchangeRateToCurrencyCode = "XOF"
	ExchangeRateToCurrencyCodeXpd ExchangeRateToCurrencyCode = "XPD"
	ExchangeRateToCurrencyCodeXpf ExchangeRateToCurrencyCode = "XPF"
	ExchangeRateToCurrencyCodeXpt ExchangeRateToCurrencyCode = "XPT"
	ExchangeRateToCurrencyCodeXsu ExchangeRateToCurrencyCode = "XSU"
	ExchangeRateToCurrencyCodeXua ExchangeRateToCurrencyCode = "XUA"
	ExchangeRateToCurrencyCodeYer ExchangeRateToCurrencyCode = "YER"
	ExchangeRateToCurrencyCodeZar ExchangeRateToCurrencyCode = "ZAR"
	ExchangeRateToCurrencyCodeZmw ExchangeRateToCurrencyCode = "ZMW"
	ExchangeRateToCurrencyCodeZwl ExchangeRateToCurrencyCode = "ZWL"
)

// The properties ConversionRate, FromCurrencyCode are required.
type ExchangeRateCreateRequestParam struct {
	ConversionRate float64 `json:"conversionRate,required"`
	// Any of "AED", "AFN", "ALL", "AMD", "ANG", "AOA", "ARS", "AUD", "AWG", "AZN",
	// "BAM", "BBD", "BDT", "BGN", "BHD", "BIF", "BMD", "BND", "BOB", "BOV", "BRL",
	// "BSD", "BTN", "BWP", "BYN", "BZD", "CAD", "CDF", "CHE", "CHF", "CHW", "CLF",
	// "CLP", "CNY", "COP", "COU", "CRC", "CUC", "CUP", "CVE", "CZK", "DJF", "DKK",
	// "DOP", "DZD", "EGP", "ERN", "ETB", "EUR", "FJD", "FKP", "GBP", "GEL", "GHS",
	// "GIP", "GMD", "GNF", "GTQ", "GYD", "HKD", "HNL", "HRK", "HTG", "HUF", "IDR",
	// "ILS", "INR", "IQD", "IRR", "ISK", "JMD", "JOD", "JPY", "KES", "KGS", "KHR",
	// "KMF", "KPW", "KRW", "KWD", "KYD", "KZT", "LAK", "LBP", "LKR", "LRD", "LSL",
	// "LYD", "MAD", "MDL", "MGA", "MKD", "MMK", "MNT", "MOP", "MRU", "MUR", "MVR",
	// "MWK", "MXN", "MXV", "MYR", "MZN", "NAD", "NGN", "NIO", "NOK", "NPR", "NZD",
	// "OMR", "PAB", "PEN", "PGK", "PHP", "PKR", "PLN", "PYG", "QAR", "RON", "RSD",
	// "RUB", "RWF", "SAR", "SBD", "SCR", "SDG", "SEK", "SGD", "SHP", "SLL", "SOS",
	// "SRD", "SSP", "STN", "SVC", "SYP", "SZL", "THB", "TJS", "TMT", "TND", "TOP",
	// "TRY", "TTD", "TWD", "TZS", "UAH", "UGX", "USD", "USN", "UYI", "UYU", "UZS",
	// "VEF", "VND", "VUV", "WST", "XAF", "XAG", "XAU", "XBA", "XBB", "XBC", "XBD",
	// "XCD", "XDR", "XOF", "XPD", "XPF", "XPT", "XSU", "XUA", "YER", "ZAR", "ZMW",
	// "ZWL".
	FromCurrencyCode ExchangeRateCreateRequestFromCurrencyCode `json:"fromCurrencyCode,omitzero,required"`
	EffectiveAt      param.Opt[time.Time]                      `json:"effectiveAt,omitzero" format:"date-time"`
	paramObj
}

func (r ExchangeRateCreateRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow ExchangeRateCreateRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExchangeRateCreateRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExchangeRateCreateRequestFromCurrencyCode string

const (
	ExchangeRateCreateRequestFromCurrencyCodeAed ExchangeRateCreateRequestFromCurrencyCode = "AED"
	ExchangeRateCreateRequestFromCurrencyCodeAfn ExchangeRateCreateRequestFromCurrencyCode = "AFN"
	ExchangeRateCreateRequestFromCurrencyCodeAll ExchangeRateCreateRequestFromCurrencyCode = "ALL"
	ExchangeRateCreateRequestFromCurrencyCodeAmd ExchangeRateCreateRequestFromCurrencyCode = "AMD"
	ExchangeRateCreateRequestFromCurrencyCodeAng ExchangeRateCreateRequestFromCurrencyCode = "ANG"
	ExchangeRateCreateRequestFromCurrencyCodeAoa ExchangeRateCreateRequestFromCurrencyCode = "AOA"
	ExchangeRateCreateRequestFromCurrencyCodeArs ExchangeRateCreateRequestFromCurrencyCode = "ARS"
	ExchangeRateCreateRequestFromCurrencyCodeAud ExchangeRateCreateRequestFromCurrencyCode = "AUD"
	ExchangeRateCreateRequestFromCurrencyCodeAwg ExchangeRateCreateRequestFromCurrencyCode = "AWG"
	ExchangeRateCreateRequestFromCurrencyCodeAzn ExchangeRateCreateRequestFromCurrencyCode = "AZN"
	ExchangeRateCreateRequestFromCurrencyCodeBam ExchangeRateCreateRequestFromCurrencyCode = "BAM"
	ExchangeRateCreateRequestFromCurrencyCodeBbd ExchangeRateCreateRequestFromCurrencyCode = "BBD"
	ExchangeRateCreateRequestFromCurrencyCodeBdt ExchangeRateCreateRequestFromCurrencyCode = "BDT"
	ExchangeRateCreateRequestFromCurrencyCodeBgn ExchangeRateCreateRequestFromCurrencyCode = "BGN"
	ExchangeRateCreateRequestFromCurrencyCodeBhd ExchangeRateCreateRequestFromCurrencyCode = "BHD"
	ExchangeRateCreateRequestFromCurrencyCodeBif ExchangeRateCreateRequestFromCurrencyCode = "BIF"
	ExchangeRateCreateRequestFromCurrencyCodeBmd ExchangeRateCreateRequestFromCurrencyCode = "BMD"
	ExchangeRateCreateRequestFromCurrencyCodeBnd ExchangeRateCreateRequestFromCurrencyCode = "BND"
	ExchangeRateCreateRequestFromCurrencyCodeBob ExchangeRateCreateRequestFromCurrencyCode = "BOB"
	ExchangeRateCreateRequestFromCurrencyCodeBov ExchangeRateCreateRequestFromCurrencyCode = "BOV"
	ExchangeRateCreateRequestFromCurrencyCodeBrl ExchangeRateCreateRequestFromCurrencyCode = "BRL"
	ExchangeRateCreateRequestFromCurrencyCodeBsd ExchangeRateCreateRequestFromCurrencyCode = "BSD"
	ExchangeRateCreateRequestFromCurrencyCodeBtn ExchangeRateCreateRequestFromCurrencyCode = "BTN"
	ExchangeRateCreateRequestFromCurrencyCodeBwp ExchangeRateCreateRequestFromCurrencyCode = "BWP"
	ExchangeRateCreateRequestFromCurrencyCodeByn ExchangeRateCreateRequestFromCurrencyCode = "BYN"
	ExchangeRateCreateRequestFromCurrencyCodeBzd ExchangeRateCreateRequestFromCurrencyCode = "BZD"
	ExchangeRateCreateRequestFromCurrencyCodeCad ExchangeRateCreateRequestFromCurrencyCode = "CAD"
	ExchangeRateCreateRequestFromCurrencyCodeCdf ExchangeRateCreateRequestFromCurrencyCode = "CDF"
	ExchangeRateCreateRequestFromCurrencyCodeChe ExchangeRateCreateRequestFromCurrencyCode = "CHE"
	ExchangeRateCreateRequestFromCurrencyCodeChf ExchangeRateCreateRequestFromCurrencyCode = "CHF"
	ExchangeRateCreateRequestFromCurrencyCodeChw ExchangeRateCreateRequestFromCurrencyCode = "CHW"
	ExchangeRateCreateRequestFromCurrencyCodeClf ExchangeRateCreateRequestFromCurrencyCode = "CLF"
	ExchangeRateCreateRequestFromCurrencyCodeClp ExchangeRateCreateRequestFromCurrencyCode = "CLP"
	ExchangeRateCreateRequestFromCurrencyCodeCny ExchangeRateCreateRequestFromCurrencyCode = "CNY"
	ExchangeRateCreateRequestFromCurrencyCodeCop ExchangeRateCreateRequestFromCurrencyCode = "COP"
	ExchangeRateCreateRequestFromCurrencyCodeCou ExchangeRateCreateRequestFromCurrencyCode = "COU"
	ExchangeRateCreateRequestFromCurrencyCodeCrc ExchangeRateCreateRequestFromCurrencyCode = "CRC"
	ExchangeRateCreateRequestFromCurrencyCodeCuc ExchangeRateCreateRequestFromCurrencyCode = "CUC"
	ExchangeRateCreateRequestFromCurrencyCodeCup ExchangeRateCreateRequestFromCurrencyCode = "CUP"
	ExchangeRateCreateRequestFromCurrencyCodeCve ExchangeRateCreateRequestFromCurrencyCode = "CVE"
	ExchangeRateCreateRequestFromCurrencyCodeCzk ExchangeRateCreateRequestFromCurrencyCode = "CZK"
	ExchangeRateCreateRequestFromCurrencyCodeDjf ExchangeRateCreateRequestFromCurrencyCode = "DJF"
	ExchangeRateCreateRequestFromCurrencyCodeDkk ExchangeRateCreateRequestFromCurrencyCode = "DKK"
	ExchangeRateCreateRequestFromCurrencyCodeDop ExchangeRateCreateRequestFromCurrencyCode = "DOP"
	ExchangeRateCreateRequestFromCurrencyCodeDzd ExchangeRateCreateRequestFromCurrencyCode = "DZD"
	ExchangeRateCreateRequestFromCurrencyCodeEgp ExchangeRateCreateRequestFromCurrencyCode = "EGP"
	ExchangeRateCreateRequestFromCurrencyCodeErn ExchangeRateCreateRequestFromCurrencyCode = "ERN"
	ExchangeRateCreateRequestFromCurrencyCodeEtb ExchangeRateCreateRequestFromCurrencyCode = "ETB"
	ExchangeRateCreateRequestFromCurrencyCodeEur ExchangeRateCreateRequestFromCurrencyCode = "EUR"
	ExchangeRateCreateRequestFromCurrencyCodeFjd ExchangeRateCreateRequestFromCurrencyCode = "FJD"
	ExchangeRateCreateRequestFromCurrencyCodeFkp ExchangeRateCreateRequestFromCurrencyCode = "FKP"
	ExchangeRateCreateRequestFromCurrencyCodeGbp ExchangeRateCreateRequestFromCurrencyCode = "GBP"
	ExchangeRateCreateRequestFromCurrencyCodeGel ExchangeRateCreateRequestFromCurrencyCode = "GEL"
	ExchangeRateCreateRequestFromCurrencyCodeGhs ExchangeRateCreateRequestFromCurrencyCode = "GHS"
	ExchangeRateCreateRequestFromCurrencyCodeGip ExchangeRateCreateRequestFromCurrencyCode = "GIP"
	ExchangeRateCreateRequestFromCurrencyCodeGmd ExchangeRateCreateRequestFromCurrencyCode = "GMD"
	ExchangeRateCreateRequestFromCurrencyCodeGnf ExchangeRateCreateRequestFromCurrencyCode = "GNF"
	ExchangeRateCreateRequestFromCurrencyCodeGtq ExchangeRateCreateRequestFromCurrencyCode = "GTQ"
	ExchangeRateCreateRequestFromCurrencyCodeGyd ExchangeRateCreateRequestFromCurrencyCode = "GYD"
	ExchangeRateCreateRequestFromCurrencyCodeHkd ExchangeRateCreateRequestFromCurrencyCode = "HKD"
	ExchangeRateCreateRequestFromCurrencyCodeHnl ExchangeRateCreateRequestFromCurrencyCode = "HNL"
	ExchangeRateCreateRequestFromCurrencyCodeHrk ExchangeRateCreateRequestFromCurrencyCode = "HRK"
	ExchangeRateCreateRequestFromCurrencyCodeHtg ExchangeRateCreateRequestFromCurrencyCode = "HTG"
	ExchangeRateCreateRequestFromCurrencyCodeHuf ExchangeRateCreateRequestFromCurrencyCode = "HUF"
	ExchangeRateCreateRequestFromCurrencyCodeIdr ExchangeRateCreateRequestFromCurrencyCode = "IDR"
	ExchangeRateCreateRequestFromCurrencyCodeIls ExchangeRateCreateRequestFromCurrencyCode = "ILS"
	ExchangeRateCreateRequestFromCurrencyCodeInr ExchangeRateCreateRequestFromCurrencyCode = "INR"
	ExchangeRateCreateRequestFromCurrencyCodeIqd ExchangeRateCreateRequestFromCurrencyCode = "IQD"
	ExchangeRateCreateRequestFromCurrencyCodeIrr ExchangeRateCreateRequestFromCurrencyCode = "IRR"
	ExchangeRateCreateRequestFromCurrencyCodeIsk ExchangeRateCreateRequestFromCurrencyCode = "ISK"
	ExchangeRateCreateRequestFromCurrencyCodeJmd ExchangeRateCreateRequestFromCurrencyCode = "JMD"
	ExchangeRateCreateRequestFromCurrencyCodeJod ExchangeRateCreateRequestFromCurrencyCode = "JOD"
	ExchangeRateCreateRequestFromCurrencyCodeJpy ExchangeRateCreateRequestFromCurrencyCode = "JPY"
	ExchangeRateCreateRequestFromCurrencyCodeKes ExchangeRateCreateRequestFromCurrencyCode = "KES"
	ExchangeRateCreateRequestFromCurrencyCodeKgs ExchangeRateCreateRequestFromCurrencyCode = "KGS"
	ExchangeRateCreateRequestFromCurrencyCodeKhr ExchangeRateCreateRequestFromCurrencyCode = "KHR"
	ExchangeRateCreateRequestFromCurrencyCodeKmf ExchangeRateCreateRequestFromCurrencyCode = "KMF"
	ExchangeRateCreateRequestFromCurrencyCodeKpw ExchangeRateCreateRequestFromCurrencyCode = "KPW"
	ExchangeRateCreateRequestFromCurrencyCodeKrw ExchangeRateCreateRequestFromCurrencyCode = "KRW"
	ExchangeRateCreateRequestFromCurrencyCodeKwd ExchangeRateCreateRequestFromCurrencyCode = "KWD"
	ExchangeRateCreateRequestFromCurrencyCodeKyd ExchangeRateCreateRequestFromCurrencyCode = "KYD"
	ExchangeRateCreateRequestFromCurrencyCodeKzt ExchangeRateCreateRequestFromCurrencyCode = "KZT"
	ExchangeRateCreateRequestFromCurrencyCodeLak ExchangeRateCreateRequestFromCurrencyCode = "LAK"
	ExchangeRateCreateRequestFromCurrencyCodeLbp ExchangeRateCreateRequestFromCurrencyCode = "LBP"
	ExchangeRateCreateRequestFromCurrencyCodeLkr ExchangeRateCreateRequestFromCurrencyCode = "LKR"
	ExchangeRateCreateRequestFromCurrencyCodeLrd ExchangeRateCreateRequestFromCurrencyCode = "LRD"
	ExchangeRateCreateRequestFromCurrencyCodeLsl ExchangeRateCreateRequestFromCurrencyCode = "LSL"
	ExchangeRateCreateRequestFromCurrencyCodeLyd ExchangeRateCreateRequestFromCurrencyCode = "LYD"
	ExchangeRateCreateRequestFromCurrencyCodeMad ExchangeRateCreateRequestFromCurrencyCode = "MAD"
	ExchangeRateCreateRequestFromCurrencyCodeMdl ExchangeRateCreateRequestFromCurrencyCode = "MDL"
	ExchangeRateCreateRequestFromCurrencyCodeMga ExchangeRateCreateRequestFromCurrencyCode = "MGA"
	ExchangeRateCreateRequestFromCurrencyCodeMkd ExchangeRateCreateRequestFromCurrencyCode = "MKD"
	ExchangeRateCreateRequestFromCurrencyCodeMmk ExchangeRateCreateRequestFromCurrencyCode = "MMK"
	ExchangeRateCreateRequestFromCurrencyCodeMnt ExchangeRateCreateRequestFromCurrencyCode = "MNT"
	ExchangeRateCreateRequestFromCurrencyCodeMop ExchangeRateCreateRequestFromCurrencyCode = "MOP"
	ExchangeRateCreateRequestFromCurrencyCodeMru ExchangeRateCreateRequestFromCurrencyCode = "MRU"
	ExchangeRateCreateRequestFromCurrencyCodeMur ExchangeRateCreateRequestFromCurrencyCode = "MUR"
	ExchangeRateCreateRequestFromCurrencyCodeMvr ExchangeRateCreateRequestFromCurrencyCode = "MVR"
	ExchangeRateCreateRequestFromCurrencyCodeMwk ExchangeRateCreateRequestFromCurrencyCode = "MWK"
	ExchangeRateCreateRequestFromCurrencyCodeMxn ExchangeRateCreateRequestFromCurrencyCode = "MXN"
	ExchangeRateCreateRequestFromCurrencyCodeMxv ExchangeRateCreateRequestFromCurrencyCode = "MXV"
	ExchangeRateCreateRequestFromCurrencyCodeMyr ExchangeRateCreateRequestFromCurrencyCode = "MYR"
	ExchangeRateCreateRequestFromCurrencyCodeMzn ExchangeRateCreateRequestFromCurrencyCode = "MZN"
	ExchangeRateCreateRequestFromCurrencyCodeNad ExchangeRateCreateRequestFromCurrencyCode = "NAD"
	ExchangeRateCreateRequestFromCurrencyCodeNgn ExchangeRateCreateRequestFromCurrencyCode = "NGN"
	ExchangeRateCreateRequestFromCurrencyCodeNio ExchangeRateCreateRequestFromCurrencyCode = "NIO"
	ExchangeRateCreateRequestFromCurrencyCodeNok ExchangeRateCreateRequestFromCurrencyCode = "NOK"
	ExchangeRateCreateRequestFromCurrencyCodeNpr ExchangeRateCreateRequestFromCurrencyCode = "NPR"
	ExchangeRateCreateRequestFromCurrencyCodeNzd ExchangeRateCreateRequestFromCurrencyCode = "NZD"
	ExchangeRateCreateRequestFromCurrencyCodeOmr ExchangeRateCreateRequestFromCurrencyCode = "OMR"
	ExchangeRateCreateRequestFromCurrencyCodePab ExchangeRateCreateRequestFromCurrencyCode = "PAB"
	ExchangeRateCreateRequestFromCurrencyCodePen ExchangeRateCreateRequestFromCurrencyCode = "PEN"
	ExchangeRateCreateRequestFromCurrencyCodePgk ExchangeRateCreateRequestFromCurrencyCode = "PGK"
	ExchangeRateCreateRequestFromCurrencyCodePhp ExchangeRateCreateRequestFromCurrencyCode = "PHP"
	ExchangeRateCreateRequestFromCurrencyCodePkr ExchangeRateCreateRequestFromCurrencyCode = "PKR"
	ExchangeRateCreateRequestFromCurrencyCodePln ExchangeRateCreateRequestFromCurrencyCode = "PLN"
	ExchangeRateCreateRequestFromCurrencyCodePyg ExchangeRateCreateRequestFromCurrencyCode = "PYG"
	ExchangeRateCreateRequestFromCurrencyCodeQar ExchangeRateCreateRequestFromCurrencyCode = "QAR"
	ExchangeRateCreateRequestFromCurrencyCodeRon ExchangeRateCreateRequestFromCurrencyCode = "RON"
	ExchangeRateCreateRequestFromCurrencyCodeRsd ExchangeRateCreateRequestFromCurrencyCode = "RSD"
	ExchangeRateCreateRequestFromCurrencyCodeRub ExchangeRateCreateRequestFromCurrencyCode = "RUB"
	ExchangeRateCreateRequestFromCurrencyCodeRwf ExchangeRateCreateRequestFromCurrencyCode = "RWF"
	ExchangeRateCreateRequestFromCurrencyCodeSar ExchangeRateCreateRequestFromCurrencyCode = "SAR"
	ExchangeRateCreateRequestFromCurrencyCodeSbd ExchangeRateCreateRequestFromCurrencyCode = "SBD"
	ExchangeRateCreateRequestFromCurrencyCodeScr ExchangeRateCreateRequestFromCurrencyCode = "SCR"
	ExchangeRateCreateRequestFromCurrencyCodeSdg ExchangeRateCreateRequestFromCurrencyCode = "SDG"
	ExchangeRateCreateRequestFromCurrencyCodeSek ExchangeRateCreateRequestFromCurrencyCode = "SEK"
	ExchangeRateCreateRequestFromCurrencyCodeSgd ExchangeRateCreateRequestFromCurrencyCode = "SGD"
	ExchangeRateCreateRequestFromCurrencyCodeShp ExchangeRateCreateRequestFromCurrencyCode = "SHP"
	ExchangeRateCreateRequestFromCurrencyCodeSll ExchangeRateCreateRequestFromCurrencyCode = "SLL"
	ExchangeRateCreateRequestFromCurrencyCodeSos ExchangeRateCreateRequestFromCurrencyCode = "SOS"
	ExchangeRateCreateRequestFromCurrencyCodeSrd ExchangeRateCreateRequestFromCurrencyCode = "SRD"
	ExchangeRateCreateRequestFromCurrencyCodeSsp ExchangeRateCreateRequestFromCurrencyCode = "SSP"
	ExchangeRateCreateRequestFromCurrencyCodeStn ExchangeRateCreateRequestFromCurrencyCode = "STN"
	ExchangeRateCreateRequestFromCurrencyCodeSvc ExchangeRateCreateRequestFromCurrencyCode = "SVC"
	ExchangeRateCreateRequestFromCurrencyCodeSyp ExchangeRateCreateRequestFromCurrencyCode = "SYP"
	ExchangeRateCreateRequestFromCurrencyCodeSzl ExchangeRateCreateRequestFromCurrencyCode = "SZL"
	ExchangeRateCreateRequestFromCurrencyCodeThb ExchangeRateCreateRequestFromCurrencyCode = "THB"
	ExchangeRateCreateRequestFromCurrencyCodeTjs ExchangeRateCreateRequestFromCurrencyCode = "TJS"
	ExchangeRateCreateRequestFromCurrencyCodeTmt ExchangeRateCreateRequestFromCurrencyCode = "TMT"
	ExchangeRateCreateRequestFromCurrencyCodeTnd ExchangeRateCreateRequestFromCurrencyCode = "TND"
	ExchangeRateCreateRequestFromCurrencyCodeTop ExchangeRateCreateRequestFromCurrencyCode = "TOP"
	ExchangeRateCreateRequestFromCurrencyCodeTry ExchangeRateCreateRequestFromCurrencyCode = "TRY"
	ExchangeRateCreateRequestFromCurrencyCodeTtd ExchangeRateCreateRequestFromCurrencyCode = "TTD"
	ExchangeRateCreateRequestFromCurrencyCodeTwd ExchangeRateCreateRequestFromCurrencyCode = "TWD"
	ExchangeRateCreateRequestFromCurrencyCodeTzs ExchangeRateCreateRequestFromCurrencyCode = "TZS"
	ExchangeRateCreateRequestFromCurrencyCodeUah ExchangeRateCreateRequestFromCurrencyCode = "UAH"
	ExchangeRateCreateRequestFromCurrencyCodeUgx ExchangeRateCreateRequestFromCurrencyCode = "UGX"
	ExchangeRateCreateRequestFromCurrencyCodeUsd ExchangeRateCreateRequestFromCurrencyCode = "USD"
	ExchangeRateCreateRequestFromCurrencyCodeUsn ExchangeRateCreateRequestFromCurrencyCode = "USN"
	ExchangeRateCreateRequestFromCurrencyCodeUyi ExchangeRateCreateRequestFromCurrencyCode = "UYI"
	ExchangeRateCreateRequestFromCurrencyCodeUyu ExchangeRateCreateRequestFromCurrencyCode = "UYU"
	ExchangeRateCreateRequestFromCurrencyCodeUzs ExchangeRateCreateRequestFromCurrencyCode = "UZS"
	ExchangeRateCreateRequestFromCurrencyCodeVef ExchangeRateCreateRequestFromCurrencyCode = "VEF"
	ExchangeRateCreateRequestFromCurrencyCodeVnd ExchangeRateCreateRequestFromCurrencyCode = "VND"
	ExchangeRateCreateRequestFromCurrencyCodeVuv ExchangeRateCreateRequestFromCurrencyCode = "VUV"
	ExchangeRateCreateRequestFromCurrencyCodeWst ExchangeRateCreateRequestFromCurrencyCode = "WST"
	ExchangeRateCreateRequestFromCurrencyCodeXaf ExchangeRateCreateRequestFromCurrencyCode = "XAF"
	ExchangeRateCreateRequestFromCurrencyCodeXag ExchangeRateCreateRequestFromCurrencyCode = "XAG"
	ExchangeRateCreateRequestFromCurrencyCodeXau ExchangeRateCreateRequestFromCurrencyCode = "XAU"
	ExchangeRateCreateRequestFromCurrencyCodeXba ExchangeRateCreateRequestFromCurrencyCode = "XBA"
	ExchangeRateCreateRequestFromCurrencyCodeXbb ExchangeRateCreateRequestFromCurrencyCode = "XBB"
	ExchangeRateCreateRequestFromCurrencyCodeXbc ExchangeRateCreateRequestFromCurrencyCode = "XBC"
	ExchangeRateCreateRequestFromCurrencyCodeXbd ExchangeRateCreateRequestFromCurrencyCode = "XBD"
	ExchangeRateCreateRequestFromCurrencyCodeXcd ExchangeRateCreateRequestFromCurrencyCode = "XCD"
	ExchangeRateCreateRequestFromCurrencyCodeXdr ExchangeRateCreateRequestFromCurrencyCode = "XDR"
	ExchangeRateCreateRequestFromCurrencyCodeXof ExchangeRateCreateRequestFromCurrencyCode = "XOF"
	ExchangeRateCreateRequestFromCurrencyCodeXpd ExchangeRateCreateRequestFromCurrencyCode = "XPD"
	ExchangeRateCreateRequestFromCurrencyCodeXpf ExchangeRateCreateRequestFromCurrencyCode = "XPF"
	ExchangeRateCreateRequestFromCurrencyCodeXpt ExchangeRateCreateRequestFromCurrencyCode = "XPT"
	ExchangeRateCreateRequestFromCurrencyCodeXsu ExchangeRateCreateRequestFromCurrencyCode = "XSU"
	ExchangeRateCreateRequestFromCurrencyCodeXua ExchangeRateCreateRequestFromCurrencyCode = "XUA"
	ExchangeRateCreateRequestFromCurrencyCodeYer ExchangeRateCreateRequestFromCurrencyCode = "YER"
	ExchangeRateCreateRequestFromCurrencyCodeZar ExchangeRateCreateRequestFromCurrencyCode = "ZAR"
	ExchangeRateCreateRequestFromCurrencyCodeZmw ExchangeRateCreateRequestFromCurrencyCode = "ZMW"
	ExchangeRateCreateRequestFromCurrencyCodeZwl ExchangeRateCreateRequestFromCurrencyCode = "ZWL"
)

// The property ConversionRate is required.
type ExchangeRateMultiplierParam struct {
	ConversionRate float64              `json:"conversionRate,required"`
	EffectiveAt    param.Opt[time.Time] `json:"effectiveAt,omitzero" format:"date-time"`
	paramObj
}

func (r ExchangeRateMultiplierParam) MarshalJSON() (data []byte, err error) {
	type shadow ExchangeRateMultiplierParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExchangeRateMultiplierParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties ID, ConversionRate are required.
type ExchangeRateUpdateRequestParam struct {
	ID             string               `json:"id,required"`
	ConversionRate float64              `json:"conversionRate,required"`
	EffectiveAt    param.Opt[time.Time] `json:"effectiveAt,omitzero" format:"date-time"`
	paramObj
}

func (r ExchangeRateUpdateRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow ExchangeRateUpdateRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExchangeRateUpdateRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CurrencyBatchNewParams struct {
	BatchInputExchangeRateCreateRequest BatchInputExchangeRateCreateRequestParam
	paramObj
}

func (r CurrencyBatchNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.BatchInputExchangeRateCreateRequest)
}
func (r *CurrencyBatchNewParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.BatchInputExchangeRateCreateRequest)
}

type CurrencyBatchGetParams struct {
	BatchInputPublicObjectID shared.BatchInputPublicObjectIDParam
	paramObj
}

func (r CurrencyBatchGetParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.BatchInputPublicObjectID)
}
func (r *CurrencyBatchGetParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.BatchInputPublicObjectID)
}

type CurrencyBatchUpdateParams struct {
	BatchInputExchangeRateUpdateRequest BatchInputExchangeRateUpdateRequestParam
	paramObj
}

func (r CurrencyBatchUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.BatchInputExchangeRateUpdateRequest)
}
func (r *CurrencyBatchUpdateParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.BatchInputExchangeRateUpdateRequest)
}

type CurrencyNewExchangeRateParams struct {
	ExchangeRateCreateRequest ExchangeRateCreateRequestParam
	paramObj
}

func (r CurrencyNewExchangeRateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.ExchangeRateCreateRequest)
}
func (r *CurrencyNewExchangeRateParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.ExchangeRateCreateRequest)
}

type CurrencyUpdateCompanyCurrencyParams struct {
	CompanyCurrencyUpdateRequest CompanyCurrencyUpdateRequestParam
	paramObj
}

func (r CurrencyUpdateCompanyCurrencyParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.CompanyCurrencyUpdateRequest)
}
func (r *CurrencyUpdateCompanyCurrencyParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.CompanyCurrencyUpdateRequest)
}

type CurrencyUpdateExchangeRateParams struct {
	ExchangeRateMultiplier ExchangeRateMultiplierParam
	paramObj
}

func (r CurrencyUpdateExchangeRateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.ExchangeRateMultiplier)
}
func (r *CurrencyUpdateExchangeRateParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.ExchangeRateMultiplier)
}

type CurrencyUpdateVisibilityParams struct {
	CurrencyPairUpdate CurrencyPairUpdateParam
	paramObj
}

func (r CurrencyUpdateVisibilityParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.CurrencyPairUpdate)
}
func (r *CurrencyUpdateVisibilityParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.CurrencyPairUpdate)
}
