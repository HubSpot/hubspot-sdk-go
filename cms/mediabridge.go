// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cms

import (
	"github.com/stainless-sdks/hubspot-sdk-go/option"
)

// MediaBridgeService contains methods and other services that help with
// interacting with the Hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMediaBridgeService] method instead.
type MediaBridgeService struct {
	Options            []option.RequestOption
	Events             MediaBridgeEventService
	Groups             MediaBridgeGroupService
	IntegratorSettings MediaBridgeIntegratorSettingService
	Properties         MediaBridgePropertyService
	Schemas            MediaBridgeSchemaService
}

// NewMediaBridgeService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewMediaBridgeService(opts ...option.RequestOption) (r MediaBridgeService) {
	r = MediaBridgeService{}
	r.Options = opts
	r.Events = NewMediaBridgeEventService(opts...)
	r.Groups = NewMediaBridgeGroupService(opts...)
	r.IntegratorSettings = NewMediaBridgeIntegratorSettingService(opts...)
	r.Properties = NewMediaBridgePropertyService(opts...)
	r.Schemas = NewMediaBridgeSchemaService(opts...)
	return
}
