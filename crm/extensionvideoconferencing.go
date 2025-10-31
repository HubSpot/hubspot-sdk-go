// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package crm

import (
	"github.com/stainless-sdks/hubspot-sdk-go/option"
)

// ExtensionVideoConferencingService contains methods and other services that help
// with interacting with the Hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewExtensionVideoConferencingService] method instead.
type ExtensionVideoConferencingService struct {
	Options  []option.RequestOption
	Settings ExtensionVideoConferencingSettingService
}

// NewExtensionVideoConferencingService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewExtensionVideoConferencingService(opts ...option.RequestOption) (r ExtensionVideoConferencingService) {
	r = ExtensionVideoConferencingService{}
	r.Options = opts
	r.Settings = NewExtensionVideoConferencingSettingService(opts...)
	return
}
