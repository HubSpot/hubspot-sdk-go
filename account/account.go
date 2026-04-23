// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package account

import (
	"context"
	"net/http"
	"slices"
	"time"

	"github.com/stainless-sdks/hubspot-sdk-go/internal/apijson"
	"github.com/stainless-sdks/hubspot-sdk-go/internal/requestconfig"
	"github.com/stainless-sdks/hubspot-sdk-go/option"
	"github.com/stainless-sdks/hubspot-sdk-go/packages/respjson"
)

// AccountService contains methods and other services that help with interacting
// with the hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountService] method instead.
type AccountService struct {
	options  []option.RequestOption
	Activity ActivityService
}

// NewAccountService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewAccountService(opts ...option.RequestOption) (r AccountService) {
	r = AccountService{}
	r.options = opts
	r.Activity = NewActivityService(opts...)
	return
}

// Retrieve account details such as the account type, time zone, currencies, and
// data hosting location.
func (r *AccountService) Get(ctx context.Context, opts ...option.RequestOption) (res *PortalInformationResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "account-info/2026-03/details"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Retrieve the daily API usage for private apps in the account, along with
// information about usage limits.
func (r *AccountService) GetDailyPrivateAppsUsage(ctx context.Context, opts ...option.RequestOption) (res *CollectionResponseAPIUsageNoPaging, err error) {
	opts = slices.Concat(r.options, opts)
	path := "account-info/2026-03/api-usage/daily/private-apps"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type APIUsage struct {
	// Indicates when the cache was last updated.
	CollectedAt time.Time `json:"collectedAt" api:"required" format:"date-time"`
	// How many API calls an account has made for the current day.
	CurrentUsage int64 `json:"currentUsage" api:"required"`
	// Status of fetching the information, including if the data came from the cache.
	//
	// Any of "CACHED", "FAILURE", "NOTFOUND", "SUCCESS", "TIMEOUT".
	FetchStatus APIUsageFetchStatus `json:"fetchStatus" api:"required"`
	// Name of the limit type.
	Name string `json:"name" api:"required"`
	// Limits by which a single integration can consume the HubSpot public APIs.
	UsageLimit int64 `json:"usageLimit" api:"required"`
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
	APIUsageFetchStatusCached   APIUsageFetchStatus = "CACHED"
	APIUsageFetchStatusFailure  APIUsageFetchStatus = "FAILURE"
	APIUsageFetchStatusNotfound APIUsageFetchStatus = "NOTFOUND"
	APIUsageFetchStatusSuccess  APIUsageFetchStatus = "SUCCESS"
	APIUsageFetchStatusTimeout  APIUsageFetchStatus = "TIMEOUT"
)

type CollectionResponseAPIUsageNoPaging struct {
	Results []APIUsage `json:"results" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Results     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CollectionResponseAPIUsageNoPaging) RawJSON() string { return r.JSON.raw }
func (r *CollectionResponseAPIUsageNoPaging) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PortalInformationResponse struct {
	// The type of account, such as APP_DEVELOPER, DEVELOPER_TEST, SANDBOX, or
	// STANDARD.
	//
	// Any of "APP_DEVELOPER", "DEVELOPER_TEST", "SANDBOX", "STANDARD".
	AccountType          PortalInformationResponseAccountType `json:"accountType" api:"required"`
	AdditionalCurrencies []string                             `json:"additionalCurrencies" api:"required"`
	// The primary currency used by the company.
	CompanyCurrency string `json:"companyCurrency" api:"required"`
	// The location where the account's data is hosted.
	DataHostingLocation string `json:"dataHostingLocation" api:"required"`
	// The unique identifier for the HubSpot account.
	PortalID int64 `json:"portalId" api:"required"`
	// The time zone in which the account operates.
	TimeZone string `json:"timeZone" api:"required"`
	// The domain used for accessing the HubSpot user interface.
	UiDomain string `json:"uiDomain" api:"required"`
	// The time zone offset from UTC in hours and minutes.
	UtcOffset string `json:"utcOffset" api:"required"`
	// The time zone offset from UTC in milliseconds.
	UtcOffsetMilliseconds int64 `json:"utcOffsetMilliseconds" api:"required"`
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

// The type of account, such as APP_DEVELOPER, DEVELOPER_TEST, SANDBOX, or
// STANDARD.
type PortalInformationResponseAccountType string

const (
	PortalInformationResponseAccountTypeAppDeveloper  PortalInformationResponseAccountType = "APP_DEVELOPER"
	PortalInformationResponseAccountTypeDeveloperTest PortalInformationResponseAccountType = "DEVELOPER_TEST"
	PortalInformationResponseAccountTypeSandbox       PortalInformationResponseAccountType = "SANDBOX"
	PortalInformationResponseAccountTypeStandard      PortalInformationResponseAccountType = "STANDARD"
)
