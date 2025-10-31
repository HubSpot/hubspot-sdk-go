// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package account

import (
	"time"

	"github.com/stainless-sdks/hubspot-sdk-go/internal/apijson"
	"github.com/stainless-sdks/hubspot-sdk-go/marketing"
	"github.com/stainless-sdks/hubspot-sdk-go/option"
	"github.com/stainless-sdks/hubspot-sdk-go/packages/respjson"
)

// AccountService contains methods and other services that help with interacting
// with the Hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountService] method instead.
type AccountService struct {
	Options  []option.RequestOption
	Activity ActivityService
	Details  DetailService
	Usage    UsageService
}

// NewAccountService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewAccountService(opts ...option.RequestOption) (r AccountService) {
	r = AccountService{}
	r.Options = opts
	r.Activity = NewActivityService(opts...)
	r.Details = NewDetailService(opts...)
	r.Usage = NewUsageService(opts...)
	return
}

// API usage and limits information for a HubSpot account.
type APIUsage struct {
	// Indicates when the cache was last updated.
	CollectedAt time.Time `json:"collectedAt,required" format:"date-time"`
	// How many API calls an account has made for the current day.
	CurrentUsage int64 `json:"currentUsage,required"`
	// Status of fetching the information, including if the data came from the cache.
	//
	// Any of "SUCCESS", "TIMEOUT", "FAILURE", "CACHED", "NOTFOUND".
	FetchStatus APIUsageFetchStatus `json:"fetchStatus,required"`
	// Name of the limit type.
	Name string `json:"name,required"`
	// Limits by which a single integration can consume the HubSpot public APIs.
	UsageLimit int64 `json:"usageLimit,required"`
	// Time that the limit will reset.
	ResetsAt time.Time `json:"resetsAt" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CollectedAt  respjson.Field
		CurrentUsage respjson.Field
		FetchStatus  respjson.Field
		Name         respjson.Field
		UsageLimit   respjson.Field
		ResetsAt     respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r APIUsage) RawJSON() string { return r.JSON.raw }
func (r *APIUsage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Status of fetching the information, including if the data came from the cache.
type APIUsageFetchStatus string

const (
	APIUsageFetchStatusSuccess  APIUsageFetchStatus = "SUCCESS"
	APIUsageFetchStatusTimeout  APIUsageFetchStatus = "TIMEOUT"
	APIUsageFetchStatusFailure  APIUsageFetchStatus = "FAILURE"
	APIUsageFetchStatusCached   APIUsageFetchStatus = "CACHED"
	APIUsageFetchStatusNotfound APIUsageFetchStatus = "NOTFOUND"
)

type CollectionResponseAPIUsage struct {
	Results []APIUsage `json:"results,required"`
	// Contains information pagination of results.
	Paging marketing.Paging `json:"paging"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Results     respjson.Field
		Paging      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CollectionResponseAPIUsage) RawJSON() string { return r.JSON.raw }
func (r *CollectionResponseAPIUsage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PortalInformationResponse struct {
	// Any of "STANDARD", "DEVELOPER_TEST", "SANDBOX", "APP_DEVELOPER".
	AccountType           PortalInformationResponseAccountType `json:"accountType,required"`
	AdditionalCurrencies  []string                             `json:"additionalCurrencies,required"`
	CompanyCurrency       string                               `json:"companyCurrency,required"`
	DataHostingLocation   string                               `json:"dataHostingLocation,required"`
	PortalID              int64                                `json:"portalId,required"`
	TimeZone              string                               `json:"timeZone,required"`
	UiDomain              string                               `json:"uiDomain,required"`
	UtcOffset             string                               `json:"utcOffset,required"`
	UtcOffsetMilliseconds int64                                `json:"utcOffsetMilliseconds,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AccountType           respjson.Field
		AdditionalCurrencies  respjson.Field
		CompanyCurrency       respjson.Field
		DataHostingLocation   respjson.Field
		PortalID              respjson.Field
		TimeZone              respjson.Field
		UiDomain              respjson.Field
		UtcOffset             respjson.Field
		UtcOffsetMilliseconds respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PortalInformationResponse) RawJSON() string { return r.JSON.raw }
func (r *PortalInformationResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PortalInformationResponseAccountType string

const (
	PortalInformationResponseAccountTypeStandard      PortalInformationResponseAccountType = "STANDARD"
	PortalInformationResponseAccountTypeDeveloperTest PortalInformationResponseAccountType = "DEVELOPER_TEST"
	PortalInformationResponseAccountTypeSandbox       PortalInformationResponseAccountType = "SANDBOX"
	PortalInformationResponseAccountTypeAppDeveloper  PortalInformationResponseAccountType = "APP_DEVELOPER"
)
