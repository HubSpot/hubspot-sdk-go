// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package communication_preferences

import (
	"github.com/HubSpot/hubspot-sdk-go/option"
)

// StatusService contains methods and other services that help with interacting
// with the hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewStatusService] method instead.
type StatusService struct {
	options []option.RequestOption
	Batch   StatusBatchService
}

// NewStatusService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewStatusService(opts ...option.RequestOption) (r StatusService) {
	r = StatusService{}
	r.options = opts
	r.Batch = NewStatusBatchService(opts...)
	return
}
