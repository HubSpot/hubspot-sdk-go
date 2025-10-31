// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package crm

import (
	"github.com/stainless-sdks/hubspot-sdk-go/option"
)

// ExtensionService contains methods and other services that help with interacting
// with the Hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewExtensionService] method instead.
type ExtensionService struct {
	Options           []option.RequestOption
	Calling           ExtensionCallingService
	Cards             ExtensionCardService
	VideoConferencing ExtensionVideoConferencingService
	Videoconferencing ExtensionVideoconferencingService
}

// NewExtensionService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewExtensionService(opts ...option.RequestOption) (r ExtensionService) {
	r = ExtensionService{}
	r.Options = opts
	r.Calling = NewExtensionCallingService(opts...)
	r.Cards = NewExtensionCardService(opts...)
	r.VideoConferencing = NewExtensionVideoConferencingService(opts...)
	r.Videoconferencing = NewExtensionVideoconferencingService(opts...)
	return
}
