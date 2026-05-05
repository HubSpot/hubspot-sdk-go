// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package meta

import (
	"github.com/HubSpot/hubspot-sdk-go/option"
)

// MetaService contains methods and other services that help with interacting with
// the hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMetaService] method instead.
type MetaService struct {
	options []option.RequestOption
	Origins OriginService
}

// NewMetaService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewMetaService(opts ...option.RequestOption) (r MetaService) {
	r = MetaService{}
	r.options = opts
	r.Origins = NewOriginService(opts...)
	return
}
