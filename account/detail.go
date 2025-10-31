// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package account

import (
	"context"
	"net/http"
	"slices"

	"github.com/stainless-sdks/hubspot-sdk-go/internal/requestconfig"
	"github.com/stainless-sdks/hubspot-sdk-go/option"
)

// DetailService contains methods and other services that help with interacting
// with the Hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDetailService] method instead.
type DetailService struct {
	Options []option.RequestOption
}

// NewDetailService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewDetailService(opts ...option.RequestOption) (r DetailService) {
	r = DetailService{}
	r.Options = opts
	return
}

// Retrieve account details such as the account type, time zone, currencies, and
// data hosting location.
func (r *DetailService) Get(ctx context.Context, opts ...option.RequestOption) (res *PortalInformationResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "account-info/v3/details"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}
