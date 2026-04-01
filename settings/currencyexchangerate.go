// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package settings

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

// CurrencyExchangeRateService contains methods and other services that help with
// interacting with the hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCurrencyExchangeRateService] method instead.
type CurrencyExchangeRateService struct {
	Options []option.RequestOption
	Batch   CurrencyExchangeRateBatchService
}

// NewCurrencyExchangeRateService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewCurrencyExchangeRateService(opts ...option.RequestOption) (r CurrencyExchangeRateService) {
	r = CurrencyExchangeRateService{}
	r.Options = opts
	r.Batch = NewCurrencyExchangeRateBatchService(opts...)
	return
}

func (r *CurrencyExchangeRateService) NewExchangeRate(ctx context.Context, body CurrencyExchangeRateNewExchangeRateParams, opts ...option.RequestOption) (res *ExchangeRate, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "settings/currencies/2026-03/exchange-rates"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

func (r *CurrencyExchangeRateService) GetExchangeRateByID(ctx context.Context, exchangeRateID string, opts ...option.RequestOption) (res *ExchangeRate, err error) {
	opts = slices.Concat(r.Options, opts)
	if exchangeRateID == "" {
		err = errors.New("missing required exchangeRateId parameter")
		return nil, err
	}
	path := fmt.Sprintf("settings/currencies/2026-03/exchange-rates/%s", exchangeRateID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

func (r *CurrencyExchangeRateService) ListCurrentExchangeRates(ctx context.Context, opts ...option.RequestOption) (res *CollectionResponseExchangeRateNoPaging, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "settings/currencies/2026-03/exchange-rates/current"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

func (r *CurrencyExchangeRateService) ListExchangeRates(ctx context.Context, query CurrencyExchangeRateListExchangeRatesParams, opts ...option.RequestOption) (res *pagination.Page[ExchangeRate], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "settings/currencies/2026-03/exchange-rates"
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

func (r *CurrencyExchangeRateService) ListExchangeRatesAutoPaging(ctx context.Context, query CurrencyExchangeRateListExchangeRatesParams, opts ...option.RequestOption) *pagination.PageAutoPager[ExchangeRate] {
	return pagination.NewPageAutoPager(r.ListExchangeRates(ctx, query, opts...))
}

func (r *CurrencyExchangeRateService) UpdateExchangeRate(ctx context.Context, exchangeRateID string, body CurrencyExchangeRateUpdateExchangeRateParams, opts ...option.RequestOption) (res *ExchangeRate, err error) {
	opts = slices.Concat(r.Options, opts)
	if exchangeRateID == "" {
		err = errors.New("missing required exchangeRateId parameter")
		return nil, err
	}
	path := fmt.Sprintf("settings/currencies/2026-03/exchange-rates/%s", exchangeRateID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

func (r *CurrencyExchangeRateService) UpdateVisibility(ctx context.Context, body CurrencyExchangeRateUpdateVisibilityParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "settings/currencies/2026-03/exchange-rates/update-visibility"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return err
}

type CurrencyExchangeRateNewExchangeRateParams struct {
	ExchangeRateCreateRequest ExchangeRateCreateRequestParam
	paramObj
}

func (r CurrencyExchangeRateNewExchangeRateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.ExchangeRateCreateRequest)
}
func (r *CurrencyExchangeRateNewExchangeRateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CurrencyExchangeRateListExchangeRatesParams struct {
	// The paging cursor token of the last successfully read resource will be returned
	// as the `paging.next.after` JSON property of a paged response containing more
	// results.
	After param.Opt[string] `query:"after,omitzero" json:"-"`
	// The maximum number of results to display per page.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
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
	FromCurrencyCode CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCode `query:"fromCurrencyCode,omitzero" json:"-"`
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
	ToCurrencyCode CurrencyExchangeRateListExchangeRatesParamsToCurrencyCode `query:"toCurrencyCode,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [CurrencyExchangeRateListExchangeRatesParams]'s query
// parameters as `url.Values`.
func (r CurrencyExchangeRateListExchangeRatesParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCode string

const (
	CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCodeAed CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCode = "AED"
	CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCodeAfn CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCode = "AFN"
	CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCodeAll CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCode = "ALL"
	CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCodeAmd CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCode = "AMD"
	CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCodeAng CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCode = "ANG"
	CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCodeAoa CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCode = "AOA"
	CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCodeArs CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCode = "ARS"
	CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCodeAud CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCode = "AUD"
	CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCodeAwg CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCode = "AWG"
	CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCodeAzn CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCode = "AZN"
	CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCodeBam CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCode = "BAM"
	CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCodeBbd CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCode = "BBD"
	CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCodeBdt CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCode = "BDT"
	CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCodeBgn CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCode = "BGN"
	CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCodeBhd CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCode = "BHD"
	CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCodeBif CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCode = "BIF"
	CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCodeBmd CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCode = "BMD"
	CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCodeBnd CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCode = "BND"
	CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCodeBob CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCode = "BOB"
	CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCodeBov CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCode = "BOV"
	CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCodeBrl CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCode = "BRL"
	CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCodeBsd CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCode = "BSD"
	CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCodeBtn CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCode = "BTN"
	CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCodeBwp CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCode = "BWP"
	CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCodeByn CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCode = "BYN"
	CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCodeBzd CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCode = "BZD"
	CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCodeCad CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCode = "CAD"
	CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCodeCdf CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCode = "CDF"
	CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCodeChe CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCode = "CHE"
	CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCodeChf CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCode = "CHF"
	CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCodeChw CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCode = "CHW"
	CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCodeClf CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCode = "CLF"
	CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCodeClp CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCode = "CLP"
	CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCodeCny CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCode = "CNY"
	CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCodeCop CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCode = "COP"
	CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCodeCou CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCode = "COU"
	CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCodeCrc CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCode = "CRC"
	CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCodeCuc CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCode = "CUC"
	CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCodeCup CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCode = "CUP"
	CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCodeCve CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCode = "CVE"
	CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCodeCzk CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCode = "CZK"
	CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCodeDjf CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCode = "DJF"
	CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCodeDkk CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCode = "DKK"
	CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCodeDop CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCode = "DOP"
	CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCodeDzd CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCode = "DZD"
	CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCodeEgp CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCode = "EGP"
	CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCodeErn CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCode = "ERN"
	CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCodeEtb CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCode = "ETB"
	CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCodeEur CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCode = "EUR"
	CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCodeFjd CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCode = "FJD"
	CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCodeFkp CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCode = "FKP"
	CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCodeGbp CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCode = "GBP"
	CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCodeGel CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCode = "GEL"
	CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCodeGhs CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCode = "GHS"
	CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCodeGip CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCode = "GIP"
	CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCodeGmd CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCode = "GMD"
	CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCodeGnf CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCode = "GNF"
	CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCodeGtq CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCode = "GTQ"
	CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCodeGyd CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCode = "GYD"
	CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCodeHkd CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCode = "HKD"
	CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCodeHnl CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCode = "HNL"
	CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCodeHrk CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCode = "HRK"
	CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCodeHtg CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCode = "HTG"
	CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCodeHuf CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCode = "HUF"
	CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCodeIdr CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCode = "IDR"
	CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCodeIls CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCode = "ILS"
	CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCodeInr CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCode = "INR"
	CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCodeIqd CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCode = "IQD"
	CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCodeIrr CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCode = "IRR"
	CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCodeIsk CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCode = "ISK"
	CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCodeJmd CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCode = "JMD"
	CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCodeJod CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCode = "JOD"
	CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCodeJpy CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCode = "JPY"
	CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCodeKes CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCode = "KES"
	CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCodeKgs CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCode = "KGS"
	CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCodeKhr CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCode = "KHR"
	CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCodeKmf CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCode = "KMF"
	CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCodeKpw CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCode = "KPW"
	CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCodeKrw CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCode = "KRW"
	CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCodeKwd CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCode = "KWD"
	CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCodeKyd CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCode = "KYD"
	CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCodeKzt CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCode = "KZT"
	CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCodeLak CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCode = "LAK"
	CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCodeLbp CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCode = "LBP"
	CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCodeLkr CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCode = "LKR"
	CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCodeLrd CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCode = "LRD"
	CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCodeLsl CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCode = "LSL"
	CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCodeLyd CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCode = "LYD"
	CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCodeMad CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCode = "MAD"
	CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCodeMdl CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCode = "MDL"
	CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCodeMga CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCode = "MGA"
	CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCodeMkd CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCode = "MKD"
	CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCodeMmk CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCode = "MMK"
	CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCodeMnt CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCode = "MNT"
	CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCodeMop CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCode = "MOP"
	CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCodeMru CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCode = "MRU"
	CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCodeMur CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCode = "MUR"
	CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCodeMvr CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCode = "MVR"
	CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCodeMwk CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCode = "MWK"
	CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCodeMxn CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCode = "MXN"
	CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCodeMxv CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCode = "MXV"
	CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCodeMyr CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCode = "MYR"
	CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCodeMzn CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCode = "MZN"
	CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCodeNad CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCode = "NAD"
	CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCodeNgn CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCode = "NGN"
	CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCodeNio CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCode = "NIO"
	CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCodeNok CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCode = "NOK"
	CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCodeNpr CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCode = "NPR"
	CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCodeNzd CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCode = "NZD"
	CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCodeOmr CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCode = "OMR"
	CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCodePab CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCode = "PAB"
	CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCodePen CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCode = "PEN"
	CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCodePgk CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCode = "PGK"
	CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCodePhp CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCode = "PHP"
	CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCodePkr CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCode = "PKR"
	CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCodePln CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCode = "PLN"
	CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCodePyg CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCode = "PYG"
	CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCodeQar CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCode = "QAR"
	CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCodeRon CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCode = "RON"
	CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCodeRsd CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCode = "RSD"
	CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCodeRub CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCode = "RUB"
	CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCodeRwf CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCode = "RWF"
	CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCodeSar CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCode = "SAR"
	CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCodeSbd CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCode = "SBD"
	CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCodeScr CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCode = "SCR"
	CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCodeSdg CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCode = "SDG"
	CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCodeSek CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCode = "SEK"
	CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCodeSgd CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCode = "SGD"
	CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCodeShp CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCode = "SHP"
	CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCodeSll CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCode = "SLL"
	CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCodeSos CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCode = "SOS"
	CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCodeSrd CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCode = "SRD"
	CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCodeSsp CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCode = "SSP"
	CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCodeStn CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCode = "STN"
	CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCodeSvc CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCode = "SVC"
	CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCodeSyp CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCode = "SYP"
	CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCodeSzl CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCode = "SZL"
	CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCodeThb CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCode = "THB"
	CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCodeTjs CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCode = "TJS"
	CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCodeTmt CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCode = "TMT"
	CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCodeTnd CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCode = "TND"
	CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCodeTop CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCode = "TOP"
	CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCodeTry CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCode = "TRY"
	CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCodeTtd CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCode = "TTD"
	CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCodeTwd CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCode = "TWD"
	CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCodeTzs CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCode = "TZS"
	CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCodeUah CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCode = "UAH"
	CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCodeUgx CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCode = "UGX"
	CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCodeUsd CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCode = "USD"
	CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCodeUsn CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCode = "USN"
	CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCodeUyi CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCode = "UYI"
	CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCodeUyu CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCode = "UYU"
	CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCodeUzs CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCode = "UZS"
	CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCodeVef CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCode = "VEF"
	CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCodeVnd CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCode = "VND"
	CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCodeVuv CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCode = "VUV"
	CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCodeWst CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCode = "WST"
	CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCodeXaf CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCode = "XAF"
	CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCodeXag CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCode = "XAG"
	CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCodeXau CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCode = "XAU"
	CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCodeXba CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCode = "XBA"
	CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCodeXbb CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCode = "XBB"
	CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCodeXbc CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCode = "XBC"
	CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCodeXbd CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCode = "XBD"
	CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCodeXcd CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCode = "XCD"
	CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCodeXdr CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCode = "XDR"
	CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCodeXof CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCode = "XOF"
	CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCodeXpd CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCode = "XPD"
	CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCodeXpf CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCode = "XPF"
	CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCodeXpt CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCode = "XPT"
	CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCodeXsu CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCode = "XSU"
	CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCodeXua CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCode = "XUA"
	CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCodeYer CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCode = "YER"
	CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCodeZar CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCode = "ZAR"
	CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCodeZmw CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCode = "ZMW"
	CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCodeZwl CurrencyExchangeRateListExchangeRatesParamsFromCurrencyCode = "ZWL"
)

type CurrencyExchangeRateListExchangeRatesParamsToCurrencyCode string

const (
	CurrencyExchangeRateListExchangeRatesParamsToCurrencyCodeAed CurrencyExchangeRateListExchangeRatesParamsToCurrencyCode = "AED"
	CurrencyExchangeRateListExchangeRatesParamsToCurrencyCodeAfn CurrencyExchangeRateListExchangeRatesParamsToCurrencyCode = "AFN"
	CurrencyExchangeRateListExchangeRatesParamsToCurrencyCodeAll CurrencyExchangeRateListExchangeRatesParamsToCurrencyCode = "ALL"
	CurrencyExchangeRateListExchangeRatesParamsToCurrencyCodeAmd CurrencyExchangeRateListExchangeRatesParamsToCurrencyCode = "AMD"
	CurrencyExchangeRateListExchangeRatesParamsToCurrencyCodeAng CurrencyExchangeRateListExchangeRatesParamsToCurrencyCode = "ANG"
	CurrencyExchangeRateListExchangeRatesParamsToCurrencyCodeAoa CurrencyExchangeRateListExchangeRatesParamsToCurrencyCode = "AOA"
	CurrencyExchangeRateListExchangeRatesParamsToCurrencyCodeArs CurrencyExchangeRateListExchangeRatesParamsToCurrencyCode = "ARS"
	CurrencyExchangeRateListExchangeRatesParamsToCurrencyCodeAud CurrencyExchangeRateListExchangeRatesParamsToCurrencyCode = "AUD"
	CurrencyExchangeRateListExchangeRatesParamsToCurrencyCodeAwg CurrencyExchangeRateListExchangeRatesParamsToCurrencyCode = "AWG"
	CurrencyExchangeRateListExchangeRatesParamsToCurrencyCodeAzn CurrencyExchangeRateListExchangeRatesParamsToCurrencyCode = "AZN"
	CurrencyExchangeRateListExchangeRatesParamsToCurrencyCodeBam CurrencyExchangeRateListExchangeRatesParamsToCurrencyCode = "BAM"
	CurrencyExchangeRateListExchangeRatesParamsToCurrencyCodeBbd CurrencyExchangeRateListExchangeRatesParamsToCurrencyCode = "BBD"
	CurrencyExchangeRateListExchangeRatesParamsToCurrencyCodeBdt CurrencyExchangeRateListExchangeRatesParamsToCurrencyCode = "BDT"
	CurrencyExchangeRateListExchangeRatesParamsToCurrencyCodeBgn CurrencyExchangeRateListExchangeRatesParamsToCurrencyCode = "BGN"
	CurrencyExchangeRateListExchangeRatesParamsToCurrencyCodeBhd CurrencyExchangeRateListExchangeRatesParamsToCurrencyCode = "BHD"
	CurrencyExchangeRateListExchangeRatesParamsToCurrencyCodeBif CurrencyExchangeRateListExchangeRatesParamsToCurrencyCode = "BIF"
	CurrencyExchangeRateListExchangeRatesParamsToCurrencyCodeBmd CurrencyExchangeRateListExchangeRatesParamsToCurrencyCode = "BMD"
	CurrencyExchangeRateListExchangeRatesParamsToCurrencyCodeBnd CurrencyExchangeRateListExchangeRatesParamsToCurrencyCode = "BND"
	CurrencyExchangeRateListExchangeRatesParamsToCurrencyCodeBob CurrencyExchangeRateListExchangeRatesParamsToCurrencyCode = "BOB"
	CurrencyExchangeRateListExchangeRatesParamsToCurrencyCodeBov CurrencyExchangeRateListExchangeRatesParamsToCurrencyCode = "BOV"
	CurrencyExchangeRateListExchangeRatesParamsToCurrencyCodeBrl CurrencyExchangeRateListExchangeRatesParamsToCurrencyCode = "BRL"
	CurrencyExchangeRateListExchangeRatesParamsToCurrencyCodeBsd CurrencyExchangeRateListExchangeRatesParamsToCurrencyCode = "BSD"
	CurrencyExchangeRateListExchangeRatesParamsToCurrencyCodeBtn CurrencyExchangeRateListExchangeRatesParamsToCurrencyCode = "BTN"
	CurrencyExchangeRateListExchangeRatesParamsToCurrencyCodeBwp CurrencyExchangeRateListExchangeRatesParamsToCurrencyCode = "BWP"
	CurrencyExchangeRateListExchangeRatesParamsToCurrencyCodeByn CurrencyExchangeRateListExchangeRatesParamsToCurrencyCode = "BYN"
	CurrencyExchangeRateListExchangeRatesParamsToCurrencyCodeBzd CurrencyExchangeRateListExchangeRatesParamsToCurrencyCode = "BZD"
	CurrencyExchangeRateListExchangeRatesParamsToCurrencyCodeCad CurrencyExchangeRateListExchangeRatesParamsToCurrencyCode = "CAD"
	CurrencyExchangeRateListExchangeRatesParamsToCurrencyCodeCdf CurrencyExchangeRateListExchangeRatesParamsToCurrencyCode = "CDF"
	CurrencyExchangeRateListExchangeRatesParamsToCurrencyCodeChe CurrencyExchangeRateListExchangeRatesParamsToCurrencyCode = "CHE"
	CurrencyExchangeRateListExchangeRatesParamsToCurrencyCodeChf CurrencyExchangeRateListExchangeRatesParamsToCurrencyCode = "CHF"
	CurrencyExchangeRateListExchangeRatesParamsToCurrencyCodeChw CurrencyExchangeRateListExchangeRatesParamsToCurrencyCode = "CHW"
	CurrencyExchangeRateListExchangeRatesParamsToCurrencyCodeClf CurrencyExchangeRateListExchangeRatesParamsToCurrencyCode = "CLF"
	CurrencyExchangeRateListExchangeRatesParamsToCurrencyCodeClp CurrencyExchangeRateListExchangeRatesParamsToCurrencyCode = "CLP"
	CurrencyExchangeRateListExchangeRatesParamsToCurrencyCodeCny CurrencyExchangeRateListExchangeRatesParamsToCurrencyCode = "CNY"
	CurrencyExchangeRateListExchangeRatesParamsToCurrencyCodeCop CurrencyExchangeRateListExchangeRatesParamsToCurrencyCode = "COP"
	CurrencyExchangeRateListExchangeRatesParamsToCurrencyCodeCou CurrencyExchangeRateListExchangeRatesParamsToCurrencyCode = "COU"
	CurrencyExchangeRateListExchangeRatesParamsToCurrencyCodeCrc CurrencyExchangeRateListExchangeRatesParamsToCurrencyCode = "CRC"
	CurrencyExchangeRateListExchangeRatesParamsToCurrencyCodeCuc CurrencyExchangeRateListExchangeRatesParamsToCurrencyCode = "CUC"
	CurrencyExchangeRateListExchangeRatesParamsToCurrencyCodeCup CurrencyExchangeRateListExchangeRatesParamsToCurrencyCode = "CUP"
	CurrencyExchangeRateListExchangeRatesParamsToCurrencyCodeCve CurrencyExchangeRateListExchangeRatesParamsToCurrencyCode = "CVE"
	CurrencyExchangeRateListExchangeRatesParamsToCurrencyCodeCzk CurrencyExchangeRateListExchangeRatesParamsToCurrencyCode = "CZK"
	CurrencyExchangeRateListExchangeRatesParamsToCurrencyCodeDjf CurrencyExchangeRateListExchangeRatesParamsToCurrencyCode = "DJF"
	CurrencyExchangeRateListExchangeRatesParamsToCurrencyCodeDkk CurrencyExchangeRateListExchangeRatesParamsToCurrencyCode = "DKK"
	CurrencyExchangeRateListExchangeRatesParamsToCurrencyCodeDop CurrencyExchangeRateListExchangeRatesParamsToCurrencyCode = "DOP"
	CurrencyExchangeRateListExchangeRatesParamsToCurrencyCodeDzd CurrencyExchangeRateListExchangeRatesParamsToCurrencyCode = "DZD"
	CurrencyExchangeRateListExchangeRatesParamsToCurrencyCodeEgp CurrencyExchangeRateListExchangeRatesParamsToCurrencyCode = "EGP"
	CurrencyExchangeRateListExchangeRatesParamsToCurrencyCodeErn CurrencyExchangeRateListExchangeRatesParamsToCurrencyCode = "ERN"
	CurrencyExchangeRateListExchangeRatesParamsToCurrencyCodeEtb CurrencyExchangeRateListExchangeRatesParamsToCurrencyCode = "ETB"
	CurrencyExchangeRateListExchangeRatesParamsToCurrencyCodeEur CurrencyExchangeRateListExchangeRatesParamsToCurrencyCode = "EUR"
	CurrencyExchangeRateListExchangeRatesParamsToCurrencyCodeFjd CurrencyExchangeRateListExchangeRatesParamsToCurrencyCode = "FJD"
	CurrencyExchangeRateListExchangeRatesParamsToCurrencyCodeFkp CurrencyExchangeRateListExchangeRatesParamsToCurrencyCode = "FKP"
	CurrencyExchangeRateListExchangeRatesParamsToCurrencyCodeGbp CurrencyExchangeRateListExchangeRatesParamsToCurrencyCode = "GBP"
	CurrencyExchangeRateListExchangeRatesParamsToCurrencyCodeGel CurrencyExchangeRateListExchangeRatesParamsToCurrencyCode = "GEL"
	CurrencyExchangeRateListExchangeRatesParamsToCurrencyCodeGhs CurrencyExchangeRateListExchangeRatesParamsToCurrencyCode = "GHS"
	CurrencyExchangeRateListExchangeRatesParamsToCurrencyCodeGip CurrencyExchangeRateListExchangeRatesParamsToCurrencyCode = "GIP"
	CurrencyExchangeRateListExchangeRatesParamsToCurrencyCodeGmd CurrencyExchangeRateListExchangeRatesParamsToCurrencyCode = "GMD"
	CurrencyExchangeRateListExchangeRatesParamsToCurrencyCodeGnf CurrencyExchangeRateListExchangeRatesParamsToCurrencyCode = "GNF"
	CurrencyExchangeRateListExchangeRatesParamsToCurrencyCodeGtq CurrencyExchangeRateListExchangeRatesParamsToCurrencyCode = "GTQ"
	CurrencyExchangeRateListExchangeRatesParamsToCurrencyCodeGyd CurrencyExchangeRateListExchangeRatesParamsToCurrencyCode = "GYD"
	CurrencyExchangeRateListExchangeRatesParamsToCurrencyCodeHkd CurrencyExchangeRateListExchangeRatesParamsToCurrencyCode = "HKD"
	CurrencyExchangeRateListExchangeRatesParamsToCurrencyCodeHnl CurrencyExchangeRateListExchangeRatesParamsToCurrencyCode = "HNL"
	CurrencyExchangeRateListExchangeRatesParamsToCurrencyCodeHrk CurrencyExchangeRateListExchangeRatesParamsToCurrencyCode = "HRK"
	CurrencyExchangeRateListExchangeRatesParamsToCurrencyCodeHtg CurrencyExchangeRateListExchangeRatesParamsToCurrencyCode = "HTG"
	CurrencyExchangeRateListExchangeRatesParamsToCurrencyCodeHuf CurrencyExchangeRateListExchangeRatesParamsToCurrencyCode = "HUF"
	CurrencyExchangeRateListExchangeRatesParamsToCurrencyCodeIdr CurrencyExchangeRateListExchangeRatesParamsToCurrencyCode = "IDR"
	CurrencyExchangeRateListExchangeRatesParamsToCurrencyCodeIls CurrencyExchangeRateListExchangeRatesParamsToCurrencyCode = "ILS"
	CurrencyExchangeRateListExchangeRatesParamsToCurrencyCodeInr CurrencyExchangeRateListExchangeRatesParamsToCurrencyCode = "INR"
	CurrencyExchangeRateListExchangeRatesParamsToCurrencyCodeIqd CurrencyExchangeRateListExchangeRatesParamsToCurrencyCode = "IQD"
	CurrencyExchangeRateListExchangeRatesParamsToCurrencyCodeIrr CurrencyExchangeRateListExchangeRatesParamsToCurrencyCode = "IRR"
	CurrencyExchangeRateListExchangeRatesParamsToCurrencyCodeIsk CurrencyExchangeRateListExchangeRatesParamsToCurrencyCode = "ISK"
	CurrencyExchangeRateListExchangeRatesParamsToCurrencyCodeJmd CurrencyExchangeRateListExchangeRatesParamsToCurrencyCode = "JMD"
	CurrencyExchangeRateListExchangeRatesParamsToCurrencyCodeJod CurrencyExchangeRateListExchangeRatesParamsToCurrencyCode = "JOD"
	CurrencyExchangeRateListExchangeRatesParamsToCurrencyCodeJpy CurrencyExchangeRateListExchangeRatesParamsToCurrencyCode = "JPY"
	CurrencyExchangeRateListExchangeRatesParamsToCurrencyCodeKes CurrencyExchangeRateListExchangeRatesParamsToCurrencyCode = "KES"
	CurrencyExchangeRateListExchangeRatesParamsToCurrencyCodeKgs CurrencyExchangeRateListExchangeRatesParamsToCurrencyCode = "KGS"
	CurrencyExchangeRateListExchangeRatesParamsToCurrencyCodeKhr CurrencyExchangeRateListExchangeRatesParamsToCurrencyCode = "KHR"
	CurrencyExchangeRateListExchangeRatesParamsToCurrencyCodeKmf CurrencyExchangeRateListExchangeRatesParamsToCurrencyCode = "KMF"
	CurrencyExchangeRateListExchangeRatesParamsToCurrencyCodeKpw CurrencyExchangeRateListExchangeRatesParamsToCurrencyCode = "KPW"
	CurrencyExchangeRateListExchangeRatesParamsToCurrencyCodeKrw CurrencyExchangeRateListExchangeRatesParamsToCurrencyCode = "KRW"
	CurrencyExchangeRateListExchangeRatesParamsToCurrencyCodeKwd CurrencyExchangeRateListExchangeRatesParamsToCurrencyCode = "KWD"
	CurrencyExchangeRateListExchangeRatesParamsToCurrencyCodeKyd CurrencyExchangeRateListExchangeRatesParamsToCurrencyCode = "KYD"
	CurrencyExchangeRateListExchangeRatesParamsToCurrencyCodeKzt CurrencyExchangeRateListExchangeRatesParamsToCurrencyCode = "KZT"
	CurrencyExchangeRateListExchangeRatesParamsToCurrencyCodeLak CurrencyExchangeRateListExchangeRatesParamsToCurrencyCode = "LAK"
	CurrencyExchangeRateListExchangeRatesParamsToCurrencyCodeLbp CurrencyExchangeRateListExchangeRatesParamsToCurrencyCode = "LBP"
	CurrencyExchangeRateListExchangeRatesParamsToCurrencyCodeLkr CurrencyExchangeRateListExchangeRatesParamsToCurrencyCode = "LKR"
	CurrencyExchangeRateListExchangeRatesParamsToCurrencyCodeLrd CurrencyExchangeRateListExchangeRatesParamsToCurrencyCode = "LRD"
	CurrencyExchangeRateListExchangeRatesParamsToCurrencyCodeLsl CurrencyExchangeRateListExchangeRatesParamsToCurrencyCode = "LSL"
	CurrencyExchangeRateListExchangeRatesParamsToCurrencyCodeLyd CurrencyExchangeRateListExchangeRatesParamsToCurrencyCode = "LYD"
	CurrencyExchangeRateListExchangeRatesParamsToCurrencyCodeMad CurrencyExchangeRateListExchangeRatesParamsToCurrencyCode = "MAD"
	CurrencyExchangeRateListExchangeRatesParamsToCurrencyCodeMdl CurrencyExchangeRateListExchangeRatesParamsToCurrencyCode = "MDL"
	CurrencyExchangeRateListExchangeRatesParamsToCurrencyCodeMga CurrencyExchangeRateListExchangeRatesParamsToCurrencyCode = "MGA"
	CurrencyExchangeRateListExchangeRatesParamsToCurrencyCodeMkd CurrencyExchangeRateListExchangeRatesParamsToCurrencyCode = "MKD"
	CurrencyExchangeRateListExchangeRatesParamsToCurrencyCodeMmk CurrencyExchangeRateListExchangeRatesParamsToCurrencyCode = "MMK"
	CurrencyExchangeRateListExchangeRatesParamsToCurrencyCodeMnt CurrencyExchangeRateListExchangeRatesParamsToCurrencyCode = "MNT"
	CurrencyExchangeRateListExchangeRatesParamsToCurrencyCodeMop CurrencyExchangeRateListExchangeRatesParamsToCurrencyCode = "MOP"
	CurrencyExchangeRateListExchangeRatesParamsToCurrencyCodeMru CurrencyExchangeRateListExchangeRatesParamsToCurrencyCode = "MRU"
	CurrencyExchangeRateListExchangeRatesParamsToCurrencyCodeMur CurrencyExchangeRateListExchangeRatesParamsToCurrencyCode = "MUR"
	CurrencyExchangeRateListExchangeRatesParamsToCurrencyCodeMvr CurrencyExchangeRateListExchangeRatesParamsToCurrencyCode = "MVR"
	CurrencyExchangeRateListExchangeRatesParamsToCurrencyCodeMwk CurrencyExchangeRateListExchangeRatesParamsToCurrencyCode = "MWK"
	CurrencyExchangeRateListExchangeRatesParamsToCurrencyCodeMxn CurrencyExchangeRateListExchangeRatesParamsToCurrencyCode = "MXN"
	CurrencyExchangeRateListExchangeRatesParamsToCurrencyCodeMxv CurrencyExchangeRateListExchangeRatesParamsToCurrencyCode = "MXV"
	CurrencyExchangeRateListExchangeRatesParamsToCurrencyCodeMyr CurrencyExchangeRateListExchangeRatesParamsToCurrencyCode = "MYR"
	CurrencyExchangeRateListExchangeRatesParamsToCurrencyCodeMzn CurrencyExchangeRateListExchangeRatesParamsToCurrencyCode = "MZN"
	CurrencyExchangeRateListExchangeRatesParamsToCurrencyCodeNad CurrencyExchangeRateListExchangeRatesParamsToCurrencyCode = "NAD"
	CurrencyExchangeRateListExchangeRatesParamsToCurrencyCodeNgn CurrencyExchangeRateListExchangeRatesParamsToCurrencyCode = "NGN"
	CurrencyExchangeRateListExchangeRatesParamsToCurrencyCodeNio CurrencyExchangeRateListExchangeRatesParamsToCurrencyCode = "NIO"
	CurrencyExchangeRateListExchangeRatesParamsToCurrencyCodeNok CurrencyExchangeRateListExchangeRatesParamsToCurrencyCode = "NOK"
	CurrencyExchangeRateListExchangeRatesParamsToCurrencyCodeNpr CurrencyExchangeRateListExchangeRatesParamsToCurrencyCode = "NPR"
	CurrencyExchangeRateListExchangeRatesParamsToCurrencyCodeNzd CurrencyExchangeRateListExchangeRatesParamsToCurrencyCode = "NZD"
	CurrencyExchangeRateListExchangeRatesParamsToCurrencyCodeOmr CurrencyExchangeRateListExchangeRatesParamsToCurrencyCode = "OMR"
	CurrencyExchangeRateListExchangeRatesParamsToCurrencyCodePab CurrencyExchangeRateListExchangeRatesParamsToCurrencyCode = "PAB"
	CurrencyExchangeRateListExchangeRatesParamsToCurrencyCodePen CurrencyExchangeRateListExchangeRatesParamsToCurrencyCode = "PEN"
	CurrencyExchangeRateListExchangeRatesParamsToCurrencyCodePgk CurrencyExchangeRateListExchangeRatesParamsToCurrencyCode = "PGK"
	CurrencyExchangeRateListExchangeRatesParamsToCurrencyCodePhp CurrencyExchangeRateListExchangeRatesParamsToCurrencyCode = "PHP"
	CurrencyExchangeRateListExchangeRatesParamsToCurrencyCodePkr CurrencyExchangeRateListExchangeRatesParamsToCurrencyCode = "PKR"
	CurrencyExchangeRateListExchangeRatesParamsToCurrencyCodePln CurrencyExchangeRateListExchangeRatesParamsToCurrencyCode = "PLN"
	CurrencyExchangeRateListExchangeRatesParamsToCurrencyCodePyg CurrencyExchangeRateListExchangeRatesParamsToCurrencyCode = "PYG"
	CurrencyExchangeRateListExchangeRatesParamsToCurrencyCodeQar CurrencyExchangeRateListExchangeRatesParamsToCurrencyCode = "QAR"
	CurrencyExchangeRateListExchangeRatesParamsToCurrencyCodeRon CurrencyExchangeRateListExchangeRatesParamsToCurrencyCode = "RON"
	CurrencyExchangeRateListExchangeRatesParamsToCurrencyCodeRsd CurrencyExchangeRateListExchangeRatesParamsToCurrencyCode = "RSD"
	CurrencyExchangeRateListExchangeRatesParamsToCurrencyCodeRub CurrencyExchangeRateListExchangeRatesParamsToCurrencyCode = "RUB"
	CurrencyExchangeRateListExchangeRatesParamsToCurrencyCodeRwf CurrencyExchangeRateListExchangeRatesParamsToCurrencyCode = "RWF"
	CurrencyExchangeRateListExchangeRatesParamsToCurrencyCodeSar CurrencyExchangeRateListExchangeRatesParamsToCurrencyCode = "SAR"
	CurrencyExchangeRateListExchangeRatesParamsToCurrencyCodeSbd CurrencyExchangeRateListExchangeRatesParamsToCurrencyCode = "SBD"
	CurrencyExchangeRateListExchangeRatesParamsToCurrencyCodeScr CurrencyExchangeRateListExchangeRatesParamsToCurrencyCode = "SCR"
	CurrencyExchangeRateListExchangeRatesParamsToCurrencyCodeSdg CurrencyExchangeRateListExchangeRatesParamsToCurrencyCode = "SDG"
	CurrencyExchangeRateListExchangeRatesParamsToCurrencyCodeSek CurrencyExchangeRateListExchangeRatesParamsToCurrencyCode = "SEK"
	CurrencyExchangeRateListExchangeRatesParamsToCurrencyCodeSgd CurrencyExchangeRateListExchangeRatesParamsToCurrencyCode = "SGD"
	CurrencyExchangeRateListExchangeRatesParamsToCurrencyCodeShp CurrencyExchangeRateListExchangeRatesParamsToCurrencyCode = "SHP"
	CurrencyExchangeRateListExchangeRatesParamsToCurrencyCodeSll CurrencyExchangeRateListExchangeRatesParamsToCurrencyCode = "SLL"
	CurrencyExchangeRateListExchangeRatesParamsToCurrencyCodeSos CurrencyExchangeRateListExchangeRatesParamsToCurrencyCode = "SOS"
	CurrencyExchangeRateListExchangeRatesParamsToCurrencyCodeSrd CurrencyExchangeRateListExchangeRatesParamsToCurrencyCode = "SRD"
	CurrencyExchangeRateListExchangeRatesParamsToCurrencyCodeSsp CurrencyExchangeRateListExchangeRatesParamsToCurrencyCode = "SSP"
	CurrencyExchangeRateListExchangeRatesParamsToCurrencyCodeStn CurrencyExchangeRateListExchangeRatesParamsToCurrencyCode = "STN"
	CurrencyExchangeRateListExchangeRatesParamsToCurrencyCodeSvc CurrencyExchangeRateListExchangeRatesParamsToCurrencyCode = "SVC"
	CurrencyExchangeRateListExchangeRatesParamsToCurrencyCodeSyp CurrencyExchangeRateListExchangeRatesParamsToCurrencyCode = "SYP"
	CurrencyExchangeRateListExchangeRatesParamsToCurrencyCodeSzl CurrencyExchangeRateListExchangeRatesParamsToCurrencyCode = "SZL"
	CurrencyExchangeRateListExchangeRatesParamsToCurrencyCodeThb CurrencyExchangeRateListExchangeRatesParamsToCurrencyCode = "THB"
	CurrencyExchangeRateListExchangeRatesParamsToCurrencyCodeTjs CurrencyExchangeRateListExchangeRatesParamsToCurrencyCode = "TJS"
	CurrencyExchangeRateListExchangeRatesParamsToCurrencyCodeTmt CurrencyExchangeRateListExchangeRatesParamsToCurrencyCode = "TMT"
	CurrencyExchangeRateListExchangeRatesParamsToCurrencyCodeTnd CurrencyExchangeRateListExchangeRatesParamsToCurrencyCode = "TND"
	CurrencyExchangeRateListExchangeRatesParamsToCurrencyCodeTop CurrencyExchangeRateListExchangeRatesParamsToCurrencyCode = "TOP"
	CurrencyExchangeRateListExchangeRatesParamsToCurrencyCodeTry CurrencyExchangeRateListExchangeRatesParamsToCurrencyCode = "TRY"
	CurrencyExchangeRateListExchangeRatesParamsToCurrencyCodeTtd CurrencyExchangeRateListExchangeRatesParamsToCurrencyCode = "TTD"
	CurrencyExchangeRateListExchangeRatesParamsToCurrencyCodeTwd CurrencyExchangeRateListExchangeRatesParamsToCurrencyCode = "TWD"
	CurrencyExchangeRateListExchangeRatesParamsToCurrencyCodeTzs CurrencyExchangeRateListExchangeRatesParamsToCurrencyCode = "TZS"
	CurrencyExchangeRateListExchangeRatesParamsToCurrencyCodeUah CurrencyExchangeRateListExchangeRatesParamsToCurrencyCode = "UAH"
	CurrencyExchangeRateListExchangeRatesParamsToCurrencyCodeUgx CurrencyExchangeRateListExchangeRatesParamsToCurrencyCode = "UGX"
	CurrencyExchangeRateListExchangeRatesParamsToCurrencyCodeUsd CurrencyExchangeRateListExchangeRatesParamsToCurrencyCode = "USD"
	CurrencyExchangeRateListExchangeRatesParamsToCurrencyCodeUsn CurrencyExchangeRateListExchangeRatesParamsToCurrencyCode = "USN"
	CurrencyExchangeRateListExchangeRatesParamsToCurrencyCodeUyi CurrencyExchangeRateListExchangeRatesParamsToCurrencyCode = "UYI"
	CurrencyExchangeRateListExchangeRatesParamsToCurrencyCodeUyu CurrencyExchangeRateListExchangeRatesParamsToCurrencyCode = "UYU"
	CurrencyExchangeRateListExchangeRatesParamsToCurrencyCodeUzs CurrencyExchangeRateListExchangeRatesParamsToCurrencyCode = "UZS"
	CurrencyExchangeRateListExchangeRatesParamsToCurrencyCodeVef CurrencyExchangeRateListExchangeRatesParamsToCurrencyCode = "VEF"
	CurrencyExchangeRateListExchangeRatesParamsToCurrencyCodeVnd CurrencyExchangeRateListExchangeRatesParamsToCurrencyCode = "VND"
	CurrencyExchangeRateListExchangeRatesParamsToCurrencyCodeVuv CurrencyExchangeRateListExchangeRatesParamsToCurrencyCode = "VUV"
	CurrencyExchangeRateListExchangeRatesParamsToCurrencyCodeWst CurrencyExchangeRateListExchangeRatesParamsToCurrencyCode = "WST"
	CurrencyExchangeRateListExchangeRatesParamsToCurrencyCodeXaf CurrencyExchangeRateListExchangeRatesParamsToCurrencyCode = "XAF"
	CurrencyExchangeRateListExchangeRatesParamsToCurrencyCodeXag CurrencyExchangeRateListExchangeRatesParamsToCurrencyCode = "XAG"
	CurrencyExchangeRateListExchangeRatesParamsToCurrencyCodeXau CurrencyExchangeRateListExchangeRatesParamsToCurrencyCode = "XAU"
	CurrencyExchangeRateListExchangeRatesParamsToCurrencyCodeXba CurrencyExchangeRateListExchangeRatesParamsToCurrencyCode = "XBA"
	CurrencyExchangeRateListExchangeRatesParamsToCurrencyCodeXbb CurrencyExchangeRateListExchangeRatesParamsToCurrencyCode = "XBB"
	CurrencyExchangeRateListExchangeRatesParamsToCurrencyCodeXbc CurrencyExchangeRateListExchangeRatesParamsToCurrencyCode = "XBC"
	CurrencyExchangeRateListExchangeRatesParamsToCurrencyCodeXbd CurrencyExchangeRateListExchangeRatesParamsToCurrencyCode = "XBD"
	CurrencyExchangeRateListExchangeRatesParamsToCurrencyCodeXcd CurrencyExchangeRateListExchangeRatesParamsToCurrencyCode = "XCD"
	CurrencyExchangeRateListExchangeRatesParamsToCurrencyCodeXdr CurrencyExchangeRateListExchangeRatesParamsToCurrencyCode = "XDR"
	CurrencyExchangeRateListExchangeRatesParamsToCurrencyCodeXof CurrencyExchangeRateListExchangeRatesParamsToCurrencyCode = "XOF"
	CurrencyExchangeRateListExchangeRatesParamsToCurrencyCodeXpd CurrencyExchangeRateListExchangeRatesParamsToCurrencyCode = "XPD"
	CurrencyExchangeRateListExchangeRatesParamsToCurrencyCodeXpf CurrencyExchangeRateListExchangeRatesParamsToCurrencyCode = "XPF"
	CurrencyExchangeRateListExchangeRatesParamsToCurrencyCodeXpt CurrencyExchangeRateListExchangeRatesParamsToCurrencyCode = "XPT"
	CurrencyExchangeRateListExchangeRatesParamsToCurrencyCodeXsu CurrencyExchangeRateListExchangeRatesParamsToCurrencyCode = "XSU"
	CurrencyExchangeRateListExchangeRatesParamsToCurrencyCodeXua CurrencyExchangeRateListExchangeRatesParamsToCurrencyCode = "XUA"
	CurrencyExchangeRateListExchangeRatesParamsToCurrencyCodeYer CurrencyExchangeRateListExchangeRatesParamsToCurrencyCode = "YER"
	CurrencyExchangeRateListExchangeRatesParamsToCurrencyCodeZar CurrencyExchangeRateListExchangeRatesParamsToCurrencyCode = "ZAR"
	CurrencyExchangeRateListExchangeRatesParamsToCurrencyCodeZmw CurrencyExchangeRateListExchangeRatesParamsToCurrencyCode = "ZMW"
	CurrencyExchangeRateListExchangeRatesParamsToCurrencyCodeZwl CurrencyExchangeRateListExchangeRatesParamsToCurrencyCode = "ZWL"
)

type CurrencyExchangeRateUpdateExchangeRateParams struct {
	ExchangeRateMultiplier ExchangeRateMultiplierParam
	paramObj
}

func (r CurrencyExchangeRateUpdateExchangeRateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.ExchangeRateMultiplier)
}
func (r *CurrencyExchangeRateUpdateExchangeRateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CurrencyExchangeRateUpdateVisibilityParams struct {
	CurrencyPairUpdate CurrencyPairUpdateParam
	paramObj
}

func (r CurrencyExchangeRateUpdateVisibilityParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.CurrencyPairUpdate)
}
func (r *CurrencyExchangeRateUpdateVisibilityParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
