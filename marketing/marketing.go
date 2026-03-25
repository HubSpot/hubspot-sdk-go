// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package marketing

import (
	"github.com/stainless-sdks/hubspot-sdk-go/option"
)

// MarketingService contains methods and other services that help with interacting
// with the hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMarketingService] method instead.
type MarketingService struct {
	Options   []option.RequestOption
	Campaigns CampaignService
}

// NewMarketingService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewMarketingService(opts ...option.RequestOption) (r MarketingService) {
	r = MarketingService{}
	r.Options = opts
	r.Campaigns = NewCampaignService(opts...)
	return
}
