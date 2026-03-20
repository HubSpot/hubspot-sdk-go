// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package crm

import (
	"github.com/stainless-sdks/hubspot-sdk-go/option"
)

// CrmService contains methods and other services that help with interacting with
// the hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCrmService] method instead.
type CrmService struct {
	Options []option.RequestOption
	Objects ObjectService
}

// NewCrmService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewCrmService(opts ...option.RequestOption) (r CrmService) {
	r = CrmService{}
	r.Options = opts
	r.Objects = NewObjectService(opts...)
	return
}
