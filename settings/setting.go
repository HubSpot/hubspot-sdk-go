// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package settings

import (
	"github.com/stainless-sdks/hubspot-sdk-go/option"
)

// SettingService contains methods and other services that help with interacting
// with the hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSettingService] method instead.
type SettingService struct {
	Options    []option.RequestOption
	Currencies CurrencyService
	TaxRates   TaxRateService
	Users      UserService
}

// NewSettingService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewSettingService(opts ...option.RequestOption) (r SettingService) {
	r = SettingService{}
	r.Options = opts
	r.Currencies = NewCurrencyService(opts...)
	r.TaxRates = NewTaxRateService(opts...)
	r.Users = NewUserService(opts...)
	return
}
