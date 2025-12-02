// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package account

import (
	"context"
	"net/http"
	"slices"

	"github.com/stainless-sdks/hubspot-sdk-go/internal/requestconfig"
	"github.com/stainless-sdks/hubspot-sdk-go/option"
)

// UsageService contains methods and other services that help with interacting with
// the hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewUsageService] method instead.
type UsageService struct {
	Options []option.RequestOption
}

// NewUsageService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewUsageService(opts ...option.RequestOption) (r UsageService) {
	r = UsageService{}
	r.Options = opts
	return
}

// Retrieve the daily API usage for private apps in the account, along with
// information about usage limits.
func (r *UsageService) GetDailyPrivateAppsUsage(ctx context.Context, opts ...option.RequestOption) (res *CollectionResponseAPIUsage, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "account-info/v3/api-usage/daily/private-apps"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}
