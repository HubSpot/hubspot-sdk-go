// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package crm

import (
	"context"
	"net/http"
	"slices"

	"github.com/stainless-sdks/hubspot-sdk-go/internal/requestconfig"
	"github.com/stainless-sdks/hubspot-sdk-go/option"
)

// AppUninstallService contains methods and other services that help with
// interacting with the hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAppUninstallService] method instead.
type AppUninstallService struct {
	options []option.RequestOption
}

// NewAppUninstallService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAppUninstallService(opts ...option.RequestOption) (r AppUninstallService) {
	r = AppUninstallService{}
	r.options = opts
	return
}

// Use this endpoint to uninstall your app from a customer's HubSpot account. If
// successful, this endpoint will return a 204 and the customer will receive an
// email notification that the developer has uninstall the app from their account.
func (r *AppUninstallService) Uninstall(ctx context.Context, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "appinstalls/2026-03/external-install"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}
