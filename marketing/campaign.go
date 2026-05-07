// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package marketing

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/HubSpot/hubspot-sdk-go/internal/apijson"
	"github.com/HubSpot/hubspot-sdk-go/internal/apiquery"
	shimjson "github.com/HubSpot/hubspot-sdk-go/internal/encoding/json"
	"github.com/HubSpot/hubspot-sdk-go/internal/requestconfig"
	"github.com/HubSpot/hubspot-sdk-go/option"
	"github.com/HubSpot/hubspot-sdk-go/packages/pagination"
	"github.com/HubSpot/hubspot-sdk-go/packages/param"
	"github.com/HubSpot/hubspot-sdk-go/packages/respjson"
	"github.com/HubSpot/hubspot-sdk-go/shared"
)

// CampaignService contains methods and other services that help with interacting
// with the hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCampaignService] method instead.
type CampaignService struct {
	options []option.RequestOption
	Assets  CampaignAssetService
	Batch   CampaignBatchService
	Budget  CampaignBudgetService
	Metrics CampaignMetricService
	Spend   CampaignSpendService
}

// NewCampaignService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCampaignService(opts ...option.RequestOption) (r CampaignService) {
	r = CampaignService{}
	r.options = opts
	r.Assets = NewCampaignAssetService(opts...)
	r.Batch = NewCampaignBatchService(opts...)
	r.Budget = NewCampaignBudgetService(opts...)
	r.Metrics = NewCampaignMetricService(opts...)
	r.Spend = NewCampaignSpendService(opts...)
	return
}

// Create a campaign with the specified properties and receive a copy of the
// campaign object, including its ID. Note that the 'hs_goal' property is
// deprecated and will be ignored if provided.
func (r *CampaignService) New(ctx context.Context, body CampaignNewParams, opts ...option.RequestOption) (res *PublicCampaign, err error) {
	opts = slices.Concat(r.options, opts)
	path := "marketing/campaigns/2026-03"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Perform a partial update of a campaign identified by the specified campaignGuid.
// Provided property values will be overwritten. Read-only and non-existent
// properties will cause 400 error. If an empty string is passed for any property
// in the Batch Update, it will reset that property's value.
func (r *CampaignService) Update(ctx context.Context, campaignGuid string, body CampaignUpdateParams, opts ...option.RequestOption) (res *PublicCampaign, err error) {
	opts = slices.Concat(r.options, opts)
	if campaignGuid == "" {
		err = errors.New("missing required campaignGuid parameter")
		return nil, err
	}
	path := fmt.Sprintf("marketing/campaigns/2026-03/%s", url.PathEscape(campaignGuid))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

// Retrieve a paginated list of campaigns from your HubSpot account. This endpoint
// allows you to specify sorting, pagination, and filtering options to tailor the
// results to your needs.
func (r *CampaignService) List(ctx context.Context, query CampaignListParams, opts ...option.RequestOption) (res *pagination.Page[PublicCampaign], err error) {
	var raw *http.Response
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "marketing/campaigns/2026-03"
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

// Retrieve a paginated list of campaigns from your HubSpot account. This endpoint
// allows you to specify sorting, pagination, and filtering options to tailor the
// results to your needs.
func (r *CampaignService) ListAutoPaging(ctx context.Context, query CampaignListParams, opts ...option.RequestOption) *pagination.PageAutoPager[PublicCampaign] {
	return pagination.NewPageAutoPager(r.List(ctx, query, opts...))
}

// Delete a specified campaign from the system. This call will return a 204 No
// Content response regardless of whether the campaignGuid provided corresponds to
// an existing campaign or not.
func (r *CampaignService) Delete(ctx context.Context, campaignGuid string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if campaignGuid == "" {
		err = errors.New("missing required campaignGuid parameter")
		return err
	}
	path := fmt.Sprintf("marketing/campaigns/2026-03/%s", url.PathEscape(campaignGuid))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Get a campaign identified by a specific campaignGuid with the given properties.
// Along with the campaign information, it also returns information about assets.
// Depending on the query parameters used, this can also be used to return
// information about the corresponding assets' metrics. Metrics are available only
// if startDate and endDate are provided.
func (r *CampaignService) Get(ctx context.Context, campaignGuid string, query CampaignGetParams, opts ...option.RequestOption) (res *PublicCampaignWithAssets, err error) {
	opts = slices.Concat(r.options, opts)
	if campaignGuid == "" {
		err = errors.New("missing required campaignGuid parameter")
		return nil, err
	}
	path := fmt.Sprintf("marketing/campaigns/2026-03/%s", url.PathEscape(campaignGuid))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// The property Inputs is required.
type BatchInputPublicCampaignBatchUpdateItemParam struct {
	// An array of PublicCampaignBatchUpdateItem objects, each containing the ID and
	// properties to update for a specific campaign.
	Inputs []PublicCampaignBatchUpdateItemParam `json:"inputs,omitzero" api:"required"`
	paramObj
}

func (r BatchInputPublicCampaignBatchUpdateItemParam) MarshalJSON() (data []byte, err error) {
	type shadow BatchInputPublicCampaignBatchUpdateItemParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BatchInputPublicCampaignBatchUpdateItemParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Inputs is required.
type BatchInputPublicCampaignDeleteInputParam struct {
	// An array of PublicCampaignDeleteInput objects, each specifying a campaign to be
	// deleted. Each object must include the campaign's unique identifier.
	Inputs []PublicCampaignDeleteInputParam `json:"inputs,omitzero" api:"required"`
	paramObj
}

func (r BatchInputPublicCampaignDeleteInputParam) MarshalJSON() (data []byte, err error) {
	type shadow BatchInputPublicCampaignDeleteInputParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BatchInputPublicCampaignDeleteInputParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Inputs is required.
type BatchInputPublicCampaignInputParam struct {
	// An array of PublicCampaignInput objects, each representing the properties of a
	// campaign to be created in the batch. This property is required.
	Inputs []PublicCampaignInputParam `json:"inputs,omitzero" api:"required"`
	paramObj
}

func (r BatchInputPublicCampaignInputParam) MarshalJSON() (data []byte, err error) {
	type shadow BatchInputPublicCampaignInputParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BatchInputPublicCampaignInputParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Inputs is required.
type BatchInputPublicCampaignReadInputParam struct {
	// An array of PublicCampaignReadInput objects, each containing the ID of a
	// campaign to be read. This property is required.
	Inputs []PublicCampaignReadInputParam `json:"inputs,omitzero" api:"required"`
	paramObj
}

func (r BatchInputPublicCampaignReadInputParam) MarshalJSON() (data []byte, err error) {
	type shadow BatchInputPublicCampaignReadInputParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BatchInputPublicCampaignReadInputParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BatchResponsePublicCampaign struct {
	// The date and time when the batch operation was completed, formatted as a
	// date-time string.
	CompletedAt time.Time `json:"completedAt" api:"required" format:"date-time"`
	// An array of results from the batch operation, each item representing a public
	// campaign.
	Results []PublicCampaign `json:"results" api:"required"`
	// The date and time when the batch operation started, formatted as a date-time
	// string.
	StartedAt time.Time `json:"startedAt" api:"required" format:"date-time"`
	// The current status of the batch operation, with possible values: CANCELED,
	// COMPLETE, PENDING, PROCESSING.
	//
	// Any of "CANCELED", "COMPLETE", "PENDING", "PROCESSING".
	Status BatchResponsePublicCampaignStatus `json:"status" api:"required"`
	// A map of related links associated with the batch operation.
	Links map[string]string `json:"links"`
	// The date and time when the batch operation was requested, formatted as a
	// date-time string.
	RequestedAt time.Time `json:"requestedAt" format:"date-time"`
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
func (r BatchResponsePublicCampaign) RawJSON() string { return r.JSON.raw }
func (r *BatchResponsePublicCampaign) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The current status of the batch operation, with possible values: CANCELED,
// COMPLETE, PENDING, PROCESSING.
type BatchResponsePublicCampaignStatus string

const (
	BatchResponsePublicCampaignStatusCanceled   BatchResponsePublicCampaignStatus = "CANCELED"
	BatchResponsePublicCampaignStatusComplete   BatchResponsePublicCampaignStatus = "COMPLETE"
	BatchResponsePublicCampaignStatusPending    BatchResponsePublicCampaignStatus = "PENDING"
	BatchResponsePublicCampaignStatusProcessing BatchResponsePublicCampaignStatus = "PROCESSING"
)

type BatchResponsePublicCampaignWithAssets struct {
	// The timestamp when the batch request processing was completed.
	CompletedAt time.Time `json:"completedAt" api:"required" format:"date-time"`
	// An array of results from the batch operation, each representing a public
	// campaign with assets.
	Results []PublicCampaignWithAssets `json:"results" api:"required"`
	// The timestamp when the processing of the batch request began.
	StartedAt time.Time `json:"startedAt" api:"required" format:"date-time"`
	// The current processing status of the batch operation, with possible values:
	// CANCELED, COMPLETE, PENDING, PROCESSING.
	//
	// Any of "CANCELED", "COMPLETE", "PENDING", "PROCESSING".
	Status BatchResponsePublicCampaignWithAssetsStatus `json:"status" api:"required"`
	// A collection of URLs linking to related resources or documentation.
	Links map[string]string `json:"links"`
	// The timestamp when the batch request was initially made.
	RequestedAt time.Time `json:"requestedAt" format:"date-time"`
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
func (r BatchResponsePublicCampaignWithAssets) RawJSON() string { return r.JSON.raw }
func (r *BatchResponsePublicCampaignWithAssets) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The current processing status of the batch operation, with possible values:
// CANCELED, COMPLETE, PENDING, PROCESSING.
type BatchResponsePublicCampaignWithAssetsStatus string

const (
	BatchResponsePublicCampaignWithAssetsStatusCanceled   BatchResponsePublicCampaignWithAssetsStatus = "CANCELED"
	BatchResponsePublicCampaignWithAssetsStatusComplete   BatchResponsePublicCampaignWithAssetsStatus = "COMPLETE"
	BatchResponsePublicCampaignWithAssetsStatusPending    BatchResponsePublicCampaignWithAssetsStatus = "PENDING"
	BatchResponsePublicCampaignWithAssetsStatusProcessing BatchResponsePublicCampaignWithAssetsStatus = "PROCESSING"
)

type CollectionResponseContactReferenceForwardPaging struct {
	// An array of contact references, each containing an identifier for a contact.
	Results []ContactReference   `json:"results" api:"required"`
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
func (r CollectionResponseContactReferenceForwardPaging) RawJSON() string { return r.JSON.raw }
func (r *CollectionResponseContactReferenceForwardPaging) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CollectionResponsePublicCampaignAsset struct {
	// An array of public campaign assets. Each item in the array is an object
	// representing a campaign asset.
	Results []PublicCampaignAsset `json:"results" api:"required"`
	Paging  shared.Paging         `json:"paging"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Results     respjson.Field
		Paging      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CollectionResponsePublicCampaignAsset) RawJSON() string { return r.JSON.raw }
func (r *CollectionResponsePublicCampaignAsset) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CollectionResponsePublicCampaignAssetForwardPaging struct {
	// An array of public campaign assets. Each item in the array is a reference to a
	// PublicCampaignAsset object.
	Results []PublicCampaignAsset `json:"results" api:"required"`
	Paging  shared.ForwardPaging  `json:"paging"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Results     respjson.Field
		Paging      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CollectionResponsePublicCampaignAssetForwardPaging) RawJSON() string { return r.JSON.raw }
func (r *CollectionResponsePublicCampaignAssetForwardPaging) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CollectionResponseWithTotalPublicCampaign struct {
	// An array of PublicCampaign objects, each representing a campaign with its
	// associated properties.
	Results []PublicCampaign `json:"results" api:"required"`
	// An integer representing the total number of public campaigns available.
	Total  int64         `json:"total" api:"required"`
	Paging shared.Paging `json:"paging"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Results     respjson.Field
		Total       respjson.Field
		Paging      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CollectionResponseWithTotalPublicCampaign) RawJSON() string { return r.JSON.raw }
func (r *CollectionResponseWithTotalPublicCampaign) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ContactReference struct {
	// Unique identifier for the contact.
	ID string `json:"id" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ContactReference) RawJSON() string { return r.JSON.raw }
func (r *ContactReference) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MetricsCounters struct {
	// The number of contacts influenced by the campaign.
	InfluencedContacts int64 `json:"influencedContacts" api:"required"`
	// The number of new contacts attributed to the campaign as the first touchpoint.
	NewContactsFirstTouch int64 `json:"newContactsFirstTouch" api:"required"`
	// The number of new contacts attributed to the campaign as the last touchpoint.
	NewContactsLastTouch int64 `json:"newContactsLastTouch" api:"required"`
	// The total number of sessions generated by the campaign.
	Sessions int64 `json:"sessions" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		InfluencedContacts    respjson.Field
		NewContactsFirstTouch respjson.Field
		NewContactsLastTouch  respjson.Field
		Sessions              respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MetricsCounters) RawJSON() string { return r.JSON.raw }
func (r *MetricsCounters) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PublicBudgetItem struct {
	// The unique identifier for the budget item.
	ID string `json:"id" api:"required"`
	// The monetary amount allocated for the budget item.
	Amount float64 `json:"amount" api:"required"`
	// The timestamp when the budget item was created.
	CreatedAt int64 `json:"createdAt" api:"required"`
	// The name of the budget item.
	Name string `json:"name" api:"required"`
	// The order of the budget item, indicating its sequence based on creation date.
	Order int64 `json:"order" api:"required"`
	// The timestamp when the budget item was last updated.
	UpdatedAt int64 `json:"updatedAt" api:"required"`
	// A description of the budget item.
	Description string `json:"description"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Amount      respjson.Field
		CreatedAt   respjson.Field
		Name        respjson.Field
		Order       respjson.Field
		UpdatedAt   respjson.Field
		Description respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PublicBudgetItem) RawJSON() string { return r.JSON.raw }
func (r *PublicBudgetItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Amount, Name, Order are required.
type PublicBudgetItemInputParam struct {
	// The monetary value assigned to the budget item.
	Amount float64 `json:"amount" api:"required"`
	// The name of the budget item.
	Name string `json:"name" api:"required"`
	// The sequence number indicating the order of the budget item.
	Order int64 `json:"order" api:"required"`
	// A detailed explanation or notes about the budget item.
	Description param.Opt[string] `json:"description,omitzero"`
	paramObj
}

func (r PublicBudgetItemInputParam) MarshalJSON() (data []byte, err error) {
	type shadow PublicBudgetItemInputParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PublicBudgetItemInputParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PublicBudgetTotals struct {
	// An array of budget items associated with the campaign. Each item is represented
	// by a PublicBudgetItem object.
	BudgetItems []PublicBudgetItem `json:"budgetItems" api:"required"`
	// The currency code used for the budget and spend amounts, following ISO 4217
	// standards.
	//
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
	CurrencyCode PublicBudgetTotalsCurrencyCode `json:"currencyCode" api:"required"`
	// An array of spend items associated with the campaign. Each item is represented
	// by a PublicSpendItem object.
	SpendItems []PublicSpendItem `json:"spendItems" api:"required"`
	// The total budget allocated for the campaign.
	BudgetTotal float64 `json:"budgetTotal"`
	// The remaining budget available for the campaign after accounting for all spend
	// items.
	RemainingBudget float64 `json:"remainingBudget"`
	// The total amount spent across all spend items in the campaign.
	SpendTotal float64 `json:"spendTotal"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BudgetItems     respjson.Field
		CurrencyCode    respjson.Field
		SpendItems      respjson.Field
		BudgetTotal     respjson.Field
		RemainingBudget respjson.Field
		SpendTotal      respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PublicBudgetTotals) RawJSON() string { return r.JSON.raw }
func (r *PublicBudgetTotals) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The currency code used for the budget and spend amounts, following ISO 4217
// standards.
type PublicBudgetTotalsCurrencyCode string

const (
	PublicBudgetTotalsCurrencyCodeAed PublicBudgetTotalsCurrencyCode = "AED"
	PublicBudgetTotalsCurrencyCodeAfn PublicBudgetTotalsCurrencyCode = "AFN"
	PublicBudgetTotalsCurrencyCodeAll PublicBudgetTotalsCurrencyCode = "ALL"
	PublicBudgetTotalsCurrencyCodeAmd PublicBudgetTotalsCurrencyCode = "AMD"
	PublicBudgetTotalsCurrencyCodeAng PublicBudgetTotalsCurrencyCode = "ANG"
	PublicBudgetTotalsCurrencyCodeAoa PublicBudgetTotalsCurrencyCode = "AOA"
	PublicBudgetTotalsCurrencyCodeArs PublicBudgetTotalsCurrencyCode = "ARS"
	PublicBudgetTotalsCurrencyCodeAud PublicBudgetTotalsCurrencyCode = "AUD"
	PublicBudgetTotalsCurrencyCodeAwg PublicBudgetTotalsCurrencyCode = "AWG"
	PublicBudgetTotalsCurrencyCodeAzn PublicBudgetTotalsCurrencyCode = "AZN"
	PublicBudgetTotalsCurrencyCodeBam PublicBudgetTotalsCurrencyCode = "BAM"
	PublicBudgetTotalsCurrencyCodeBbd PublicBudgetTotalsCurrencyCode = "BBD"
	PublicBudgetTotalsCurrencyCodeBdt PublicBudgetTotalsCurrencyCode = "BDT"
	PublicBudgetTotalsCurrencyCodeBgn PublicBudgetTotalsCurrencyCode = "BGN"
	PublicBudgetTotalsCurrencyCodeBhd PublicBudgetTotalsCurrencyCode = "BHD"
	PublicBudgetTotalsCurrencyCodeBif PublicBudgetTotalsCurrencyCode = "BIF"
	PublicBudgetTotalsCurrencyCodeBmd PublicBudgetTotalsCurrencyCode = "BMD"
	PublicBudgetTotalsCurrencyCodeBnd PublicBudgetTotalsCurrencyCode = "BND"
	PublicBudgetTotalsCurrencyCodeBob PublicBudgetTotalsCurrencyCode = "BOB"
	PublicBudgetTotalsCurrencyCodeBov PublicBudgetTotalsCurrencyCode = "BOV"
	PublicBudgetTotalsCurrencyCodeBrl PublicBudgetTotalsCurrencyCode = "BRL"
	PublicBudgetTotalsCurrencyCodeBsd PublicBudgetTotalsCurrencyCode = "BSD"
	PublicBudgetTotalsCurrencyCodeBtn PublicBudgetTotalsCurrencyCode = "BTN"
	PublicBudgetTotalsCurrencyCodeBwp PublicBudgetTotalsCurrencyCode = "BWP"
	PublicBudgetTotalsCurrencyCodeByn PublicBudgetTotalsCurrencyCode = "BYN"
	PublicBudgetTotalsCurrencyCodeBzd PublicBudgetTotalsCurrencyCode = "BZD"
	PublicBudgetTotalsCurrencyCodeCad PublicBudgetTotalsCurrencyCode = "CAD"
	PublicBudgetTotalsCurrencyCodeCdf PublicBudgetTotalsCurrencyCode = "CDF"
	PublicBudgetTotalsCurrencyCodeChe PublicBudgetTotalsCurrencyCode = "CHE"
	PublicBudgetTotalsCurrencyCodeChf PublicBudgetTotalsCurrencyCode = "CHF"
	PublicBudgetTotalsCurrencyCodeChw PublicBudgetTotalsCurrencyCode = "CHW"
	PublicBudgetTotalsCurrencyCodeClf PublicBudgetTotalsCurrencyCode = "CLF"
	PublicBudgetTotalsCurrencyCodeClp PublicBudgetTotalsCurrencyCode = "CLP"
	PublicBudgetTotalsCurrencyCodeCny PublicBudgetTotalsCurrencyCode = "CNY"
	PublicBudgetTotalsCurrencyCodeCop PublicBudgetTotalsCurrencyCode = "COP"
	PublicBudgetTotalsCurrencyCodeCou PublicBudgetTotalsCurrencyCode = "COU"
	PublicBudgetTotalsCurrencyCodeCrc PublicBudgetTotalsCurrencyCode = "CRC"
	PublicBudgetTotalsCurrencyCodeCuc PublicBudgetTotalsCurrencyCode = "CUC"
	PublicBudgetTotalsCurrencyCodeCup PublicBudgetTotalsCurrencyCode = "CUP"
	PublicBudgetTotalsCurrencyCodeCve PublicBudgetTotalsCurrencyCode = "CVE"
	PublicBudgetTotalsCurrencyCodeCzk PublicBudgetTotalsCurrencyCode = "CZK"
	PublicBudgetTotalsCurrencyCodeDjf PublicBudgetTotalsCurrencyCode = "DJF"
	PublicBudgetTotalsCurrencyCodeDkk PublicBudgetTotalsCurrencyCode = "DKK"
	PublicBudgetTotalsCurrencyCodeDop PublicBudgetTotalsCurrencyCode = "DOP"
	PublicBudgetTotalsCurrencyCodeDzd PublicBudgetTotalsCurrencyCode = "DZD"
	PublicBudgetTotalsCurrencyCodeEgp PublicBudgetTotalsCurrencyCode = "EGP"
	PublicBudgetTotalsCurrencyCodeErn PublicBudgetTotalsCurrencyCode = "ERN"
	PublicBudgetTotalsCurrencyCodeEtb PublicBudgetTotalsCurrencyCode = "ETB"
	PublicBudgetTotalsCurrencyCodeEur PublicBudgetTotalsCurrencyCode = "EUR"
	PublicBudgetTotalsCurrencyCodeFjd PublicBudgetTotalsCurrencyCode = "FJD"
	PublicBudgetTotalsCurrencyCodeFkp PublicBudgetTotalsCurrencyCode = "FKP"
	PublicBudgetTotalsCurrencyCodeGbp PublicBudgetTotalsCurrencyCode = "GBP"
	PublicBudgetTotalsCurrencyCodeGel PublicBudgetTotalsCurrencyCode = "GEL"
	PublicBudgetTotalsCurrencyCodeGhs PublicBudgetTotalsCurrencyCode = "GHS"
	PublicBudgetTotalsCurrencyCodeGip PublicBudgetTotalsCurrencyCode = "GIP"
	PublicBudgetTotalsCurrencyCodeGmd PublicBudgetTotalsCurrencyCode = "GMD"
	PublicBudgetTotalsCurrencyCodeGnf PublicBudgetTotalsCurrencyCode = "GNF"
	PublicBudgetTotalsCurrencyCodeGtq PublicBudgetTotalsCurrencyCode = "GTQ"
	PublicBudgetTotalsCurrencyCodeGyd PublicBudgetTotalsCurrencyCode = "GYD"
	PublicBudgetTotalsCurrencyCodeHkd PublicBudgetTotalsCurrencyCode = "HKD"
	PublicBudgetTotalsCurrencyCodeHnl PublicBudgetTotalsCurrencyCode = "HNL"
	PublicBudgetTotalsCurrencyCodeHrk PublicBudgetTotalsCurrencyCode = "HRK"
	PublicBudgetTotalsCurrencyCodeHtg PublicBudgetTotalsCurrencyCode = "HTG"
	PublicBudgetTotalsCurrencyCodeHuf PublicBudgetTotalsCurrencyCode = "HUF"
	PublicBudgetTotalsCurrencyCodeIdr PublicBudgetTotalsCurrencyCode = "IDR"
	PublicBudgetTotalsCurrencyCodeIls PublicBudgetTotalsCurrencyCode = "ILS"
	PublicBudgetTotalsCurrencyCodeInr PublicBudgetTotalsCurrencyCode = "INR"
	PublicBudgetTotalsCurrencyCodeIqd PublicBudgetTotalsCurrencyCode = "IQD"
	PublicBudgetTotalsCurrencyCodeIrr PublicBudgetTotalsCurrencyCode = "IRR"
	PublicBudgetTotalsCurrencyCodeIsk PublicBudgetTotalsCurrencyCode = "ISK"
	PublicBudgetTotalsCurrencyCodeJmd PublicBudgetTotalsCurrencyCode = "JMD"
	PublicBudgetTotalsCurrencyCodeJod PublicBudgetTotalsCurrencyCode = "JOD"
	PublicBudgetTotalsCurrencyCodeJpy PublicBudgetTotalsCurrencyCode = "JPY"
	PublicBudgetTotalsCurrencyCodeKes PublicBudgetTotalsCurrencyCode = "KES"
	PublicBudgetTotalsCurrencyCodeKgs PublicBudgetTotalsCurrencyCode = "KGS"
	PublicBudgetTotalsCurrencyCodeKhr PublicBudgetTotalsCurrencyCode = "KHR"
	PublicBudgetTotalsCurrencyCodeKmf PublicBudgetTotalsCurrencyCode = "KMF"
	PublicBudgetTotalsCurrencyCodeKpw PublicBudgetTotalsCurrencyCode = "KPW"
	PublicBudgetTotalsCurrencyCodeKrw PublicBudgetTotalsCurrencyCode = "KRW"
	PublicBudgetTotalsCurrencyCodeKwd PublicBudgetTotalsCurrencyCode = "KWD"
	PublicBudgetTotalsCurrencyCodeKyd PublicBudgetTotalsCurrencyCode = "KYD"
	PublicBudgetTotalsCurrencyCodeKzt PublicBudgetTotalsCurrencyCode = "KZT"
	PublicBudgetTotalsCurrencyCodeLak PublicBudgetTotalsCurrencyCode = "LAK"
	PublicBudgetTotalsCurrencyCodeLbp PublicBudgetTotalsCurrencyCode = "LBP"
	PublicBudgetTotalsCurrencyCodeLkr PublicBudgetTotalsCurrencyCode = "LKR"
	PublicBudgetTotalsCurrencyCodeLrd PublicBudgetTotalsCurrencyCode = "LRD"
	PublicBudgetTotalsCurrencyCodeLsl PublicBudgetTotalsCurrencyCode = "LSL"
	PublicBudgetTotalsCurrencyCodeLyd PublicBudgetTotalsCurrencyCode = "LYD"
	PublicBudgetTotalsCurrencyCodeMad PublicBudgetTotalsCurrencyCode = "MAD"
	PublicBudgetTotalsCurrencyCodeMdl PublicBudgetTotalsCurrencyCode = "MDL"
	PublicBudgetTotalsCurrencyCodeMga PublicBudgetTotalsCurrencyCode = "MGA"
	PublicBudgetTotalsCurrencyCodeMkd PublicBudgetTotalsCurrencyCode = "MKD"
	PublicBudgetTotalsCurrencyCodeMmk PublicBudgetTotalsCurrencyCode = "MMK"
	PublicBudgetTotalsCurrencyCodeMnt PublicBudgetTotalsCurrencyCode = "MNT"
	PublicBudgetTotalsCurrencyCodeMop PublicBudgetTotalsCurrencyCode = "MOP"
	PublicBudgetTotalsCurrencyCodeMru PublicBudgetTotalsCurrencyCode = "MRU"
	PublicBudgetTotalsCurrencyCodeMur PublicBudgetTotalsCurrencyCode = "MUR"
	PublicBudgetTotalsCurrencyCodeMvr PublicBudgetTotalsCurrencyCode = "MVR"
	PublicBudgetTotalsCurrencyCodeMwk PublicBudgetTotalsCurrencyCode = "MWK"
	PublicBudgetTotalsCurrencyCodeMxn PublicBudgetTotalsCurrencyCode = "MXN"
	PublicBudgetTotalsCurrencyCodeMxv PublicBudgetTotalsCurrencyCode = "MXV"
	PublicBudgetTotalsCurrencyCodeMyr PublicBudgetTotalsCurrencyCode = "MYR"
	PublicBudgetTotalsCurrencyCodeMzn PublicBudgetTotalsCurrencyCode = "MZN"
	PublicBudgetTotalsCurrencyCodeNad PublicBudgetTotalsCurrencyCode = "NAD"
	PublicBudgetTotalsCurrencyCodeNgn PublicBudgetTotalsCurrencyCode = "NGN"
	PublicBudgetTotalsCurrencyCodeNio PublicBudgetTotalsCurrencyCode = "NIO"
	PublicBudgetTotalsCurrencyCodeNok PublicBudgetTotalsCurrencyCode = "NOK"
	PublicBudgetTotalsCurrencyCodeNpr PublicBudgetTotalsCurrencyCode = "NPR"
	PublicBudgetTotalsCurrencyCodeNzd PublicBudgetTotalsCurrencyCode = "NZD"
	PublicBudgetTotalsCurrencyCodeOmr PublicBudgetTotalsCurrencyCode = "OMR"
	PublicBudgetTotalsCurrencyCodePab PublicBudgetTotalsCurrencyCode = "PAB"
	PublicBudgetTotalsCurrencyCodePen PublicBudgetTotalsCurrencyCode = "PEN"
	PublicBudgetTotalsCurrencyCodePgk PublicBudgetTotalsCurrencyCode = "PGK"
	PublicBudgetTotalsCurrencyCodePhp PublicBudgetTotalsCurrencyCode = "PHP"
	PublicBudgetTotalsCurrencyCodePkr PublicBudgetTotalsCurrencyCode = "PKR"
	PublicBudgetTotalsCurrencyCodePln PublicBudgetTotalsCurrencyCode = "PLN"
	PublicBudgetTotalsCurrencyCodePyg PublicBudgetTotalsCurrencyCode = "PYG"
	PublicBudgetTotalsCurrencyCodeQar PublicBudgetTotalsCurrencyCode = "QAR"
	PublicBudgetTotalsCurrencyCodeRon PublicBudgetTotalsCurrencyCode = "RON"
	PublicBudgetTotalsCurrencyCodeRsd PublicBudgetTotalsCurrencyCode = "RSD"
	PublicBudgetTotalsCurrencyCodeRub PublicBudgetTotalsCurrencyCode = "RUB"
	PublicBudgetTotalsCurrencyCodeRwf PublicBudgetTotalsCurrencyCode = "RWF"
	PublicBudgetTotalsCurrencyCodeSar PublicBudgetTotalsCurrencyCode = "SAR"
	PublicBudgetTotalsCurrencyCodeSbd PublicBudgetTotalsCurrencyCode = "SBD"
	PublicBudgetTotalsCurrencyCodeScr PublicBudgetTotalsCurrencyCode = "SCR"
	PublicBudgetTotalsCurrencyCodeSdg PublicBudgetTotalsCurrencyCode = "SDG"
	PublicBudgetTotalsCurrencyCodeSek PublicBudgetTotalsCurrencyCode = "SEK"
	PublicBudgetTotalsCurrencyCodeSgd PublicBudgetTotalsCurrencyCode = "SGD"
	PublicBudgetTotalsCurrencyCodeShp PublicBudgetTotalsCurrencyCode = "SHP"
	PublicBudgetTotalsCurrencyCodeSll PublicBudgetTotalsCurrencyCode = "SLL"
	PublicBudgetTotalsCurrencyCodeSos PublicBudgetTotalsCurrencyCode = "SOS"
	PublicBudgetTotalsCurrencyCodeSrd PublicBudgetTotalsCurrencyCode = "SRD"
	PublicBudgetTotalsCurrencyCodeSsp PublicBudgetTotalsCurrencyCode = "SSP"
	PublicBudgetTotalsCurrencyCodeStn PublicBudgetTotalsCurrencyCode = "STN"
	PublicBudgetTotalsCurrencyCodeSvc PublicBudgetTotalsCurrencyCode = "SVC"
	PublicBudgetTotalsCurrencyCodeSyp PublicBudgetTotalsCurrencyCode = "SYP"
	PublicBudgetTotalsCurrencyCodeSzl PublicBudgetTotalsCurrencyCode = "SZL"
	PublicBudgetTotalsCurrencyCodeThb PublicBudgetTotalsCurrencyCode = "THB"
	PublicBudgetTotalsCurrencyCodeTjs PublicBudgetTotalsCurrencyCode = "TJS"
	PublicBudgetTotalsCurrencyCodeTmt PublicBudgetTotalsCurrencyCode = "TMT"
	PublicBudgetTotalsCurrencyCodeTnd PublicBudgetTotalsCurrencyCode = "TND"
	PublicBudgetTotalsCurrencyCodeTop PublicBudgetTotalsCurrencyCode = "TOP"
	PublicBudgetTotalsCurrencyCodeTry PublicBudgetTotalsCurrencyCode = "TRY"
	PublicBudgetTotalsCurrencyCodeTtd PublicBudgetTotalsCurrencyCode = "TTD"
	PublicBudgetTotalsCurrencyCodeTwd PublicBudgetTotalsCurrencyCode = "TWD"
	PublicBudgetTotalsCurrencyCodeTzs PublicBudgetTotalsCurrencyCode = "TZS"
	PublicBudgetTotalsCurrencyCodeUah PublicBudgetTotalsCurrencyCode = "UAH"
	PublicBudgetTotalsCurrencyCodeUgx PublicBudgetTotalsCurrencyCode = "UGX"
	PublicBudgetTotalsCurrencyCodeUsd PublicBudgetTotalsCurrencyCode = "USD"
	PublicBudgetTotalsCurrencyCodeUsn PublicBudgetTotalsCurrencyCode = "USN"
	PublicBudgetTotalsCurrencyCodeUyi PublicBudgetTotalsCurrencyCode = "UYI"
	PublicBudgetTotalsCurrencyCodeUyu PublicBudgetTotalsCurrencyCode = "UYU"
	PublicBudgetTotalsCurrencyCodeUzs PublicBudgetTotalsCurrencyCode = "UZS"
	PublicBudgetTotalsCurrencyCodeVef PublicBudgetTotalsCurrencyCode = "VEF"
	PublicBudgetTotalsCurrencyCodeVnd PublicBudgetTotalsCurrencyCode = "VND"
	PublicBudgetTotalsCurrencyCodeVuv PublicBudgetTotalsCurrencyCode = "VUV"
	PublicBudgetTotalsCurrencyCodeWst PublicBudgetTotalsCurrencyCode = "WST"
	PublicBudgetTotalsCurrencyCodeXaf PublicBudgetTotalsCurrencyCode = "XAF"
	PublicBudgetTotalsCurrencyCodeXag PublicBudgetTotalsCurrencyCode = "XAG"
	PublicBudgetTotalsCurrencyCodeXau PublicBudgetTotalsCurrencyCode = "XAU"
	PublicBudgetTotalsCurrencyCodeXba PublicBudgetTotalsCurrencyCode = "XBA"
	PublicBudgetTotalsCurrencyCodeXbb PublicBudgetTotalsCurrencyCode = "XBB"
	PublicBudgetTotalsCurrencyCodeXbc PublicBudgetTotalsCurrencyCode = "XBC"
	PublicBudgetTotalsCurrencyCodeXbd PublicBudgetTotalsCurrencyCode = "XBD"
	PublicBudgetTotalsCurrencyCodeXcd PublicBudgetTotalsCurrencyCode = "XCD"
	PublicBudgetTotalsCurrencyCodeXdr PublicBudgetTotalsCurrencyCode = "XDR"
	PublicBudgetTotalsCurrencyCodeXof PublicBudgetTotalsCurrencyCode = "XOF"
	PublicBudgetTotalsCurrencyCodeXpd PublicBudgetTotalsCurrencyCode = "XPD"
	PublicBudgetTotalsCurrencyCodeXpf PublicBudgetTotalsCurrencyCode = "XPF"
	PublicBudgetTotalsCurrencyCodeXpt PublicBudgetTotalsCurrencyCode = "XPT"
	PublicBudgetTotalsCurrencyCodeXsu PublicBudgetTotalsCurrencyCode = "XSU"
	PublicBudgetTotalsCurrencyCodeXua PublicBudgetTotalsCurrencyCode = "XUA"
	PublicBudgetTotalsCurrencyCodeYer PublicBudgetTotalsCurrencyCode = "YER"
	PublicBudgetTotalsCurrencyCodeZar PublicBudgetTotalsCurrencyCode = "ZAR"
	PublicBudgetTotalsCurrencyCodeZmw PublicBudgetTotalsCurrencyCode = "ZMW"
	PublicBudgetTotalsCurrencyCodeZwl PublicBudgetTotalsCurrencyCode = "ZWL"
)

type PublicBusinessUnit struct {
	// The unique identifier for the business unit, represented as an integer.
	ID int64 `json:"id" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PublicBusinessUnit) RawJSON() string { return r.JSON.raw }
func (r *PublicBusinessUnit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PublicCampaign struct {
	// The unique identifier for the campaign.
	ID string `json:"id" api:"required"`
	// An array of business units associated with the campaign, each represented by a
	// PublicBusinessUnit object.
	BusinessUnits []PublicBusinessUnit `json:"businessUnits" api:"required"`
	// The date and time when the campaign was created, formatted as a date-time
	// string.
	CreatedAt time.Time `json:"createdAt" api:"required" format:"date-time"`
	// A collection of key-value pairs representing additional properties of the
	// campaign.
	Properties map[string]string `json:"properties" api:"required"`
	// The date and time when the campaign was last updated, formatted as a date-time
	// string.
	UpdatedAt time.Time `json:"updatedAt" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		BusinessUnits respjson.Field
		CreatedAt     respjson.Field
		Properties    respjson.Field
		UpdatedAt     respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PublicCampaign) RawJSON() string { return r.JSON.raw }
func (r *PublicCampaign) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PublicCampaignAsset struct {
	// The unique identifier for the campaign asset.
	ID string `json:"id" api:"required"`
	// A collection of metrics associated with the campaign asset, represented as
	// key-value pairs.
	Metrics map[string]float64 `json:"metrics"`
	// The name of the campaign asset.
	Name string `json:"name"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Metrics     respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PublicCampaignAsset) RawJSON() string { return r.JSON.raw }
func (r *PublicCampaignAsset) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties ID, Properties are required.
type PublicCampaignBatchUpdateItemParam struct {
	// The unique identifier for the campaign to be updated.
	ID string `json:"id" api:"required"`
	// A set of key-value pairs representing the properties to be updated for the
	// campaign.
	Properties map[string]string `json:"properties,omitzero" api:"required"`
	paramObj
}

func (r PublicCampaignBatchUpdateItemParam) MarshalJSON() (data []byte, err error) {
	type shadow PublicCampaignBatchUpdateItemParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PublicCampaignBatchUpdateItemParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property ID is required.
type PublicCampaignDeleteInputParam struct {
	// The unique identifier for the campaign to be deleted.
	ID string `json:"id" api:"required"`
	paramObj
}

func (r PublicCampaignDeleteInputParam) MarshalJSON() (data []byte, err error) {
	type shadow PublicCampaignDeleteInputParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PublicCampaignDeleteInputParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Properties is required.
type PublicCampaignInputParam struct {
	// A collection of key-value pairs representing the properties of the campaign.
	// Each key is a property name, and the corresponding value is the property's
	// value.
	Properties map[string]string `json:"properties,omitzero" api:"required"`
	paramObj
}

func (r PublicCampaignInputParam) MarshalJSON() (data []byte, err error) {
	type shadow PublicCampaignInputParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PublicCampaignInputParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property ID is required.
type PublicCampaignReadInputParam struct {
	// The unique identifier for a campaign.
	ID string `json:"id" api:"required"`
	paramObj
}

func (r PublicCampaignReadInputParam) MarshalJSON() (data []byte, err error) {
	type shadow PublicCampaignReadInputParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PublicCampaignReadInputParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PublicCampaignWithAssets struct {
	// The unique identifier for the campaign.
	ID string `json:"id" api:"required"`
	// Contains the assets associated with the campaign, each represented as a
	// collection of campaign assets.
	Assets map[string]CollectionResponsePublicCampaignAsset `json:"assets" api:"required"`
	// An array of business units associated with the campaign, each represented by a
	// PublicBusinessUnit object.
	BusinessUnits []PublicBusinessUnit `json:"businessUnits" api:"required"`
	// The date and time when the campaign was created, formatted as a date-time
	// string.
	CreatedAt time.Time `json:"createdAt" api:"required" format:"date-time"`
	// A map of key-value pairs representing the properties of the campaign.
	Properties map[string]string `json:"properties" api:"required"`
	// The date and time when the campaign was last updated, formatted as a date-time
	// string.
	UpdatedAt time.Time `json:"updatedAt" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		Assets        respjson.Field
		BusinessUnits respjson.Field
		CreatedAt     respjson.Field
		Properties    respjson.Field
		UpdatedAt     respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PublicCampaignWithAssets) RawJSON() string { return r.JSON.raw }
func (r *PublicCampaignWithAssets) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PublicSpendItem struct {
	// Unique identifier for the spend item.
	ID string `json:"id" api:"required"`
	// The monetary value associated with the spend item.
	Amount float64 `json:"amount" api:"required"`
	// The timestamp indicating when the spend item was created.
	CreatedAt int64 `json:"createdAt" api:"required"`
	// The name assigned to the spend item.
	Name string `json:"name" api:"required"`
	// The sequence order of the spend item, where 0 is the oldest.
	Order int64 `json:"order" api:"required"`
	// The timestamp indicating when the spend item was last updated.
	UpdatedAt int64 `json:"updatedAt" api:"required"`
	// A detailed explanation or notes about the spend item.
	Description string `json:"description"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Amount      respjson.Field
		CreatedAt   respjson.Field
		Name        respjson.Field
		Order       respjson.Field
		UpdatedAt   respjson.Field
		Description respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PublicSpendItem) RawJSON() string { return r.JSON.raw }
func (r *PublicSpendItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Amount, Name, Order are required.
type PublicSpendItemInputParam struct {
	// The monetary value of the spend item.
	Amount float64 `json:"amount" api:"required"`
	// The name of the spend item.
	Name string `json:"name" api:"required"`
	// The sequence number indicating the order of the spend item.
	Order int64 `json:"order" api:"required"`
	// A brief description of the spend item.
	Description param.Opt[string] `json:"description,omitzero"`
	paramObj
}

func (r PublicSpendItemInputParam) MarshalJSON() (data []byte, err error) {
	type shadow PublicSpendItemInputParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PublicSpendItemInputParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RevenueAttributionAggregate struct {
	// The number of contacts attributed to the campaign.
	ContactsNumber int64 `json:"contactsNumber"`
	// The currency code used for the revenue attribution, with accepted values
	// including AED, AFN, ALL, and others.
	//
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
	CurrencyCode RevenueAttributionAggregateCurrencyCode `json:"currencyCode"`
	// The total amount of deals attributed to the campaign.
	DealAmount float64 `json:"dealAmount"`
	// The number of deals attributed to the campaign.
	DealsNumber int64 `json:"dealsNumber"`
	// The total revenue amount attributed to the campaign.
	RevenueAmount float64 `json:"revenueAmount"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ContactsNumber respjson.Field
		CurrencyCode   respjson.Field
		DealAmount     respjson.Field
		DealsNumber    respjson.Field
		RevenueAmount  respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RevenueAttributionAggregate) RawJSON() string { return r.JSON.raw }
func (r *RevenueAttributionAggregate) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The currency code used for the revenue attribution, with accepted values
// including AED, AFN, ALL, and others.
type RevenueAttributionAggregateCurrencyCode string

const (
	RevenueAttributionAggregateCurrencyCodeAed RevenueAttributionAggregateCurrencyCode = "AED"
	RevenueAttributionAggregateCurrencyCodeAfn RevenueAttributionAggregateCurrencyCode = "AFN"
	RevenueAttributionAggregateCurrencyCodeAll RevenueAttributionAggregateCurrencyCode = "ALL"
	RevenueAttributionAggregateCurrencyCodeAmd RevenueAttributionAggregateCurrencyCode = "AMD"
	RevenueAttributionAggregateCurrencyCodeAng RevenueAttributionAggregateCurrencyCode = "ANG"
	RevenueAttributionAggregateCurrencyCodeAoa RevenueAttributionAggregateCurrencyCode = "AOA"
	RevenueAttributionAggregateCurrencyCodeArs RevenueAttributionAggregateCurrencyCode = "ARS"
	RevenueAttributionAggregateCurrencyCodeAud RevenueAttributionAggregateCurrencyCode = "AUD"
	RevenueAttributionAggregateCurrencyCodeAwg RevenueAttributionAggregateCurrencyCode = "AWG"
	RevenueAttributionAggregateCurrencyCodeAzn RevenueAttributionAggregateCurrencyCode = "AZN"
	RevenueAttributionAggregateCurrencyCodeBam RevenueAttributionAggregateCurrencyCode = "BAM"
	RevenueAttributionAggregateCurrencyCodeBbd RevenueAttributionAggregateCurrencyCode = "BBD"
	RevenueAttributionAggregateCurrencyCodeBdt RevenueAttributionAggregateCurrencyCode = "BDT"
	RevenueAttributionAggregateCurrencyCodeBgn RevenueAttributionAggregateCurrencyCode = "BGN"
	RevenueAttributionAggregateCurrencyCodeBhd RevenueAttributionAggregateCurrencyCode = "BHD"
	RevenueAttributionAggregateCurrencyCodeBif RevenueAttributionAggregateCurrencyCode = "BIF"
	RevenueAttributionAggregateCurrencyCodeBmd RevenueAttributionAggregateCurrencyCode = "BMD"
	RevenueAttributionAggregateCurrencyCodeBnd RevenueAttributionAggregateCurrencyCode = "BND"
	RevenueAttributionAggregateCurrencyCodeBob RevenueAttributionAggregateCurrencyCode = "BOB"
	RevenueAttributionAggregateCurrencyCodeBov RevenueAttributionAggregateCurrencyCode = "BOV"
	RevenueAttributionAggregateCurrencyCodeBrl RevenueAttributionAggregateCurrencyCode = "BRL"
	RevenueAttributionAggregateCurrencyCodeBsd RevenueAttributionAggregateCurrencyCode = "BSD"
	RevenueAttributionAggregateCurrencyCodeBtn RevenueAttributionAggregateCurrencyCode = "BTN"
	RevenueAttributionAggregateCurrencyCodeBwp RevenueAttributionAggregateCurrencyCode = "BWP"
	RevenueAttributionAggregateCurrencyCodeByn RevenueAttributionAggregateCurrencyCode = "BYN"
	RevenueAttributionAggregateCurrencyCodeBzd RevenueAttributionAggregateCurrencyCode = "BZD"
	RevenueAttributionAggregateCurrencyCodeCad RevenueAttributionAggregateCurrencyCode = "CAD"
	RevenueAttributionAggregateCurrencyCodeCdf RevenueAttributionAggregateCurrencyCode = "CDF"
	RevenueAttributionAggregateCurrencyCodeChe RevenueAttributionAggregateCurrencyCode = "CHE"
	RevenueAttributionAggregateCurrencyCodeChf RevenueAttributionAggregateCurrencyCode = "CHF"
	RevenueAttributionAggregateCurrencyCodeChw RevenueAttributionAggregateCurrencyCode = "CHW"
	RevenueAttributionAggregateCurrencyCodeClf RevenueAttributionAggregateCurrencyCode = "CLF"
	RevenueAttributionAggregateCurrencyCodeClp RevenueAttributionAggregateCurrencyCode = "CLP"
	RevenueAttributionAggregateCurrencyCodeCny RevenueAttributionAggregateCurrencyCode = "CNY"
	RevenueAttributionAggregateCurrencyCodeCop RevenueAttributionAggregateCurrencyCode = "COP"
	RevenueAttributionAggregateCurrencyCodeCou RevenueAttributionAggregateCurrencyCode = "COU"
	RevenueAttributionAggregateCurrencyCodeCrc RevenueAttributionAggregateCurrencyCode = "CRC"
	RevenueAttributionAggregateCurrencyCodeCuc RevenueAttributionAggregateCurrencyCode = "CUC"
	RevenueAttributionAggregateCurrencyCodeCup RevenueAttributionAggregateCurrencyCode = "CUP"
	RevenueAttributionAggregateCurrencyCodeCve RevenueAttributionAggregateCurrencyCode = "CVE"
	RevenueAttributionAggregateCurrencyCodeCzk RevenueAttributionAggregateCurrencyCode = "CZK"
	RevenueAttributionAggregateCurrencyCodeDjf RevenueAttributionAggregateCurrencyCode = "DJF"
	RevenueAttributionAggregateCurrencyCodeDkk RevenueAttributionAggregateCurrencyCode = "DKK"
	RevenueAttributionAggregateCurrencyCodeDop RevenueAttributionAggregateCurrencyCode = "DOP"
	RevenueAttributionAggregateCurrencyCodeDzd RevenueAttributionAggregateCurrencyCode = "DZD"
	RevenueAttributionAggregateCurrencyCodeEgp RevenueAttributionAggregateCurrencyCode = "EGP"
	RevenueAttributionAggregateCurrencyCodeErn RevenueAttributionAggregateCurrencyCode = "ERN"
	RevenueAttributionAggregateCurrencyCodeEtb RevenueAttributionAggregateCurrencyCode = "ETB"
	RevenueAttributionAggregateCurrencyCodeEur RevenueAttributionAggregateCurrencyCode = "EUR"
	RevenueAttributionAggregateCurrencyCodeFjd RevenueAttributionAggregateCurrencyCode = "FJD"
	RevenueAttributionAggregateCurrencyCodeFkp RevenueAttributionAggregateCurrencyCode = "FKP"
	RevenueAttributionAggregateCurrencyCodeGbp RevenueAttributionAggregateCurrencyCode = "GBP"
	RevenueAttributionAggregateCurrencyCodeGel RevenueAttributionAggregateCurrencyCode = "GEL"
	RevenueAttributionAggregateCurrencyCodeGhs RevenueAttributionAggregateCurrencyCode = "GHS"
	RevenueAttributionAggregateCurrencyCodeGip RevenueAttributionAggregateCurrencyCode = "GIP"
	RevenueAttributionAggregateCurrencyCodeGmd RevenueAttributionAggregateCurrencyCode = "GMD"
	RevenueAttributionAggregateCurrencyCodeGnf RevenueAttributionAggregateCurrencyCode = "GNF"
	RevenueAttributionAggregateCurrencyCodeGtq RevenueAttributionAggregateCurrencyCode = "GTQ"
	RevenueAttributionAggregateCurrencyCodeGyd RevenueAttributionAggregateCurrencyCode = "GYD"
	RevenueAttributionAggregateCurrencyCodeHkd RevenueAttributionAggregateCurrencyCode = "HKD"
	RevenueAttributionAggregateCurrencyCodeHnl RevenueAttributionAggregateCurrencyCode = "HNL"
	RevenueAttributionAggregateCurrencyCodeHrk RevenueAttributionAggregateCurrencyCode = "HRK"
	RevenueAttributionAggregateCurrencyCodeHtg RevenueAttributionAggregateCurrencyCode = "HTG"
	RevenueAttributionAggregateCurrencyCodeHuf RevenueAttributionAggregateCurrencyCode = "HUF"
	RevenueAttributionAggregateCurrencyCodeIdr RevenueAttributionAggregateCurrencyCode = "IDR"
	RevenueAttributionAggregateCurrencyCodeIls RevenueAttributionAggregateCurrencyCode = "ILS"
	RevenueAttributionAggregateCurrencyCodeInr RevenueAttributionAggregateCurrencyCode = "INR"
	RevenueAttributionAggregateCurrencyCodeIqd RevenueAttributionAggregateCurrencyCode = "IQD"
	RevenueAttributionAggregateCurrencyCodeIrr RevenueAttributionAggregateCurrencyCode = "IRR"
	RevenueAttributionAggregateCurrencyCodeIsk RevenueAttributionAggregateCurrencyCode = "ISK"
	RevenueAttributionAggregateCurrencyCodeJmd RevenueAttributionAggregateCurrencyCode = "JMD"
	RevenueAttributionAggregateCurrencyCodeJod RevenueAttributionAggregateCurrencyCode = "JOD"
	RevenueAttributionAggregateCurrencyCodeJpy RevenueAttributionAggregateCurrencyCode = "JPY"
	RevenueAttributionAggregateCurrencyCodeKes RevenueAttributionAggregateCurrencyCode = "KES"
	RevenueAttributionAggregateCurrencyCodeKgs RevenueAttributionAggregateCurrencyCode = "KGS"
	RevenueAttributionAggregateCurrencyCodeKhr RevenueAttributionAggregateCurrencyCode = "KHR"
	RevenueAttributionAggregateCurrencyCodeKmf RevenueAttributionAggregateCurrencyCode = "KMF"
	RevenueAttributionAggregateCurrencyCodeKpw RevenueAttributionAggregateCurrencyCode = "KPW"
	RevenueAttributionAggregateCurrencyCodeKrw RevenueAttributionAggregateCurrencyCode = "KRW"
	RevenueAttributionAggregateCurrencyCodeKwd RevenueAttributionAggregateCurrencyCode = "KWD"
	RevenueAttributionAggregateCurrencyCodeKyd RevenueAttributionAggregateCurrencyCode = "KYD"
	RevenueAttributionAggregateCurrencyCodeKzt RevenueAttributionAggregateCurrencyCode = "KZT"
	RevenueAttributionAggregateCurrencyCodeLak RevenueAttributionAggregateCurrencyCode = "LAK"
	RevenueAttributionAggregateCurrencyCodeLbp RevenueAttributionAggregateCurrencyCode = "LBP"
	RevenueAttributionAggregateCurrencyCodeLkr RevenueAttributionAggregateCurrencyCode = "LKR"
	RevenueAttributionAggregateCurrencyCodeLrd RevenueAttributionAggregateCurrencyCode = "LRD"
	RevenueAttributionAggregateCurrencyCodeLsl RevenueAttributionAggregateCurrencyCode = "LSL"
	RevenueAttributionAggregateCurrencyCodeLyd RevenueAttributionAggregateCurrencyCode = "LYD"
	RevenueAttributionAggregateCurrencyCodeMad RevenueAttributionAggregateCurrencyCode = "MAD"
	RevenueAttributionAggregateCurrencyCodeMdl RevenueAttributionAggregateCurrencyCode = "MDL"
	RevenueAttributionAggregateCurrencyCodeMga RevenueAttributionAggregateCurrencyCode = "MGA"
	RevenueAttributionAggregateCurrencyCodeMkd RevenueAttributionAggregateCurrencyCode = "MKD"
	RevenueAttributionAggregateCurrencyCodeMmk RevenueAttributionAggregateCurrencyCode = "MMK"
	RevenueAttributionAggregateCurrencyCodeMnt RevenueAttributionAggregateCurrencyCode = "MNT"
	RevenueAttributionAggregateCurrencyCodeMop RevenueAttributionAggregateCurrencyCode = "MOP"
	RevenueAttributionAggregateCurrencyCodeMru RevenueAttributionAggregateCurrencyCode = "MRU"
	RevenueAttributionAggregateCurrencyCodeMur RevenueAttributionAggregateCurrencyCode = "MUR"
	RevenueAttributionAggregateCurrencyCodeMvr RevenueAttributionAggregateCurrencyCode = "MVR"
	RevenueAttributionAggregateCurrencyCodeMwk RevenueAttributionAggregateCurrencyCode = "MWK"
	RevenueAttributionAggregateCurrencyCodeMxn RevenueAttributionAggregateCurrencyCode = "MXN"
	RevenueAttributionAggregateCurrencyCodeMxv RevenueAttributionAggregateCurrencyCode = "MXV"
	RevenueAttributionAggregateCurrencyCodeMyr RevenueAttributionAggregateCurrencyCode = "MYR"
	RevenueAttributionAggregateCurrencyCodeMzn RevenueAttributionAggregateCurrencyCode = "MZN"
	RevenueAttributionAggregateCurrencyCodeNad RevenueAttributionAggregateCurrencyCode = "NAD"
	RevenueAttributionAggregateCurrencyCodeNgn RevenueAttributionAggregateCurrencyCode = "NGN"
	RevenueAttributionAggregateCurrencyCodeNio RevenueAttributionAggregateCurrencyCode = "NIO"
	RevenueAttributionAggregateCurrencyCodeNok RevenueAttributionAggregateCurrencyCode = "NOK"
	RevenueAttributionAggregateCurrencyCodeNpr RevenueAttributionAggregateCurrencyCode = "NPR"
	RevenueAttributionAggregateCurrencyCodeNzd RevenueAttributionAggregateCurrencyCode = "NZD"
	RevenueAttributionAggregateCurrencyCodeOmr RevenueAttributionAggregateCurrencyCode = "OMR"
	RevenueAttributionAggregateCurrencyCodePab RevenueAttributionAggregateCurrencyCode = "PAB"
	RevenueAttributionAggregateCurrencyCodePen RevenueAttributionAggregateCurrencyCode = "PEN"
	RevenueAttributionAggregateCurrencyCodePgk RevenueAttributionAggregateCurrencyCode = "PGK"
	RevenueAttributionAggregateCurrencyCodePhp RevenueAttributionAggregateCurrencyCode = "PHP"
	RevenueAttributionAggregateCurrencyCodePkr RevenueAttributionAggregateCurrencyCode = "PKR"
	RevenueAttributionAggregateCurrencyCodePln RevenueAttributionAggregateCurrencyCode = "PLN"
	RevenueAttributionAggregateCurrencyCodePyg RevenueAttributionAggregateCurrencyCode = "PYG"
	RevenueAttributionAggregateCurrencyCodeQar RevenueAttributionAggregateCurrencyCode = "QAR"
	RevenueAttributionAggregateCurrencyCodeRon RevenueAttributionAggregateCurrencyCode = "RON"
	RevenueAttributionAggregateCurrencyCodeRsd RevenueAttributionAggregateCurrencyCode = "RSD"
	RevenueAttributionAggregateCurrencyCodeRub RevenueAttributionAggregateCurrencyCode = "RUB"
	RevenueAttributionAggregateCurrencyCodeRwf RevenueAttributionAggregateCurrencyCode = "RWF"
	RevenueAttributionAggregateCurrencyCodeSar RevenueAttributionAggregateCurrencyCode = "SAR"
	RevenueAttributionAggregateCurrencyCodeSbd RevenueAttributionAggregateCurrencyCode = "SBD"
	RevenueAttributionAggregateCurrencyCodeScr RevenueAttributionAggregateCurrencyCode = "SCR"
	RevenueAttributionAggregateCurrencyCodeSdg RevenueAttributionAggregateCurrencyCode = "SDG"
	RevenueAttributionAggregateCurrencyCodeSek RevenueAttributionAggregateCurrencyCode = "SEK"
	RevenueAttributionAggregateCurrencyCodeSgd RevenueAttributionAggregateCurrencyCode = "SGD"
	RevenueAttributionAggregateCurrencyCodeShp RevenueAttributionAggregateCurrencyCode = "SHP"
	RevenueAttributionAggregateCurrencyCodeSll RevenueAttributionAggregateCurrencyCode = "SLL"
	RevenueAttributionAggregateCurrencyCodeSos RevenueAttributionAggregateCurrencyCode = "SOS"
	RevenueAttributionAggregateCurrencyCodeSrd RevenueAttributionAggregateCurrencyCode = "SRD"
	RevenueAttributionAggregateCurrencyCodeSsp RevenueAttributionAggregateCurrencyCode = "SSP"
	RevenueAttributionAggregateCurrencyCodeStn RevenueAttributionAggregateCurrencyCode = "STN"
	RevenueAttributionAggregateCurrencyCodeSvc RevenueAttributionAggregateCurrencyCode = "SVC"
	RevenueAttributionAggregateCurrencyCodeSyp RevenueAttributionAggregateCurrencyCode = "SYP"
	RevenueAttributionAggregateCurrencyCodeSzl RevenueAttributionAggregateCurrencyCode = "SZL"
	RevenueAttributionAggregateCurrencyCodeThb RevenueAttributionAggregateCurrencyCode = "THB"
	RevenueAttributionAggregateCurrencyCodeTjs RevenueAttributionAggregateCurrencyCode = "TJS"
	RevenueAttributionAggregateCurrencyCodeTmt RevenueAttributionAggregateCurrencyCode = "TMT"
	RevenueAttributionAggregateCurrencyCodeTnd RevenueAttributionAggregateCurrencyCode = "TND"
	RevenueAttributionAggregateCurrencyCodeTop RevenueAttributionAggregateCurrencyCode = "TOP"
	RevenueAttributionAggregateCurrencyCodeTry RevenueAttributionAggregateCurrencyCode = "TRY"
	RevenueAttributionAggregateCurrencyCodeTtd RevenueAttributionAggregateCurrencyCode = "TTD"
	RevenueAttributionAggregateCurrencyCodeTwd RevenueAttributionAggregateCurrencyCode = "TWD"
	RevenueAttributionAggregateCurrencyCodeTzs RevenueAttributionAggregateCurrencyCode = "TZS"
	RevenueAttributionAggregateCurrencyCodeUah RevenueAttributionAggregateCurrencyCode = "UAH"
	RevenueAttributionAggregateCurrencyCodeUgx RevenueAttributionAggregateCurrencyCode = "UGX"
	RevenueAttributionAggregateCurrencyCodeUsd RevenueAttributionAggregateCurrencyCode = "USD"
	RevenueAttributionAggregateCurrencyCodeUsn RevenueAttributionAggregateCurrencyCode = "USN"
	RevenueAttributionAggregateCurrencyCodeUyi RevenueAttributionAggregateCurrencyCode = "UYI"
	RevenueAttributionAggregateCurrencyCodeUyu RevenueAttributionAggregateCurrencyCode = "UYU"
	RevenueAttributionAggregateCurrencyCodeUzs RevenueAttributionAggregateCurrencyCode = "UZS"
	RevenueAttributionAggregateCurrencyCodeVef RevenueAttributionAggregateCurrencyCode = "VEF"
	RevenueAttributionAggregateCurrencyCodeVnd RevenueAttributionAggregateCurrencyCode = "VND"
	RevenueAttributionAggregateCurrencyCodeVuv RevenueAttributionAggregateCurrencyCode = "VUV"
	RevenueAttributionAggregateCurrencyCodeWst RevenueAttributionAggregateCurrencyCode = "WST"
	RevenueAttributionAggregateCurrencyCodeXaf RevenueAttributionAggregateCurrencyCode = "XAF"
	RevenueAttributionAggregateCurrencyCodeXag RevenueAttributionAggregateCurrencyCode = "XAG"
	RevenueAttributionAggregateCurrencyCodeXau RevenueAttributionAggregateCurrencyCode = "XAU"
	RevenueAttributionAggregateCurrencyCodeXba RevenueAttributionAggregateCurrencyCode = "XBA"
	RevenueAttributionAggregateCurrencyCodeXbb RevenueAttributionAggregateCurrencyCode = "XBB"
	RevenueAttributionAggregateCurrencyCodeXbc RevenueAttributionAggregateCurrencyCode = "XBC"
	RevenueAttributionAggregateCurrencyCodeXbd RevenueAttributionAggregateCurrencyCode = "XBD"
	RevenueAttributionAggregateCurrencyCodeXcd RevenueAttributionAggregateCurrencyCode = "XCD"
	RevenueAttributionAggregateCurrencyCodeXdr RevenueAttributionAggregateCurrencyCode = "XDR"
	RevenueAttributionAggregateCurrencyCodeXof RevenueAttributionAggregateCurrencyCode = "XOF"
	RevenueAttributionAggregateCurrencyCodeXpd RevenueAttributionAggregateCurrencyCode = "XPD"
	RevenueAttributionAggregateCurrencyCodeXpf RevenueAttributionAggregateCurrencyCode = "XPF"
	RevenueAttributionAggregateCurrencyCodeXpt RevenueAttributionAggregateCurrencyCode = "XPT"
	RevenueAttributionAggregateCurrencyCodeXsu RevenueAttributionAggregateCurrencyCode = "XSU"
	RevenueAttributionAggregateCurrencyCodeXua RevenueAttributionAggregateCurrencyCode = "XUA"
	RevenueAttributionAggregateCurrencyCodeYer RevenueAttributionAggregateCurrencyCode = "YER"
	RevenueAttributionAggregateCurrencyCodeZar RevenueAttributionAggregateCurrencyCode = "ZAR"
	RevenueAttributionAggregateCurrencyCodeZmw RevenueAttributionAggregateCurrencyCode = "ZMW"
	RevenueAttributionAggregateCurrencyCodeZwl RevenueAttributionAggregateCurrencyCode = "ZWL"
)

type CampaignNewParams struct {
	PublicCampaignInput PublicCampaignInputParam
	paramObj
}

func (r CampaignNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.PublicCampaignInput)
}
func (r *CampaignNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CampaignUpdateParams struct {
	PublicCampaignInput PublicCampaignInputParam
	paramObj
}

func (r CampaignUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.PublicCampaignInput)
}
func (r *CampaignUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CampaignListParams struct {
	// The paging cursor token of the last successfully read resource will be returned
	// as the `paging.next.after` JSON property of a paged response containing more
	// results.
	After param.Opt[string] `query:"after,omitzero" json:"-"`
	// The maximum number of results to display per page.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Filter campaigns by name. Optional.
	Name param.Opt[string] `query:"name,omitzero" json:"-"`
	// The property to sort results by. Optional.
	Sort param.Opt[string] `query:"sort,omitzero" json:"-"`
	// A comma-separated list of properties to include in the response.
	//
	//	Unrecognized properties are ignored. Optional. Example:
	//	hs_name, hs_budget,hs_notes
	Properties []string `query:"properties,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [CampaignListParams]'s query parameters as `url.Values`.
func (r CampaignListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type CampaignGetParams struct {
	// The end date for fetching asset metrics, in YYYY-MM-DD format. Optional.
	// Example: 2000-01-27
	EndDate param.Opt[string] `query:"endDate,omitzero" json:"-"`
	// The start date for fetching asset metrics, in YYYY-MM-DD format.
	// Optional. Example: 2000-01-20
	StartDate param.Opt[string] `query:"startDate,omitzero" json:"-"`
	// A comma-separated list of properties to include in the response.
	//
	//	Unrecognized properties are ignored. Optional. Example: hs_name,hs_budget, hs_notes
	Properties []string `query:"properties,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [CampaignGetParams]'s query parameters as `url.Values`.
func (r CampaignGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
