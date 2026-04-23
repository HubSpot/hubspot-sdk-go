// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package conversations

import (
	"github.com/stainless-sdks/hubspot-sdk-go/option"
)

// ConversationService contains methods and other services that help with
// interacting with the hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewConversationService] method instead.
type ConversationService struct {
	options               []option.RequestOption
	CustomChannels        CustomChannelService
	VisitorIdentification VisitorIdentificationService
}

// NewConversationService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewConversationService(opts ...option.RequestOption) (r ConversationService) {
	r = ConversationService{}
	r.options = opts
	r.CustomChannels = NewCustomChannelService(opts...)
	r.VisitorIdentification = NewVisitorIdentificationService(opts...)
	return
}
