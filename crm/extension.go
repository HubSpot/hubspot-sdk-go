// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package crm

import (
	"github.com/HubSpot/hubspot-sdk-go/option"
)

// ExtensionService contains methods and other services that help with interacting
// with the hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewExtensionService] method instead.
type ExtensionService struct {
	options           []option.RequestOption
	Calling           ExtensionCallingService
	CardsDev          ExtensionCardsDevService
	VideoConferencing ExtensionVideoConferencingService
}

// NewExtensionService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewExtensionService(opts ...option.RequestOption) (r ExtensionService) {
	r = ExtensionService{}
	r.options = opts
	r.Calling = NewExtensionCallingService(opts...)
	r.CardsDev = NewExtensionCardsDevService(opts...)
	r.VideoConferencing = NewExtensionVideoConferencingService(opts...)
	return
}
