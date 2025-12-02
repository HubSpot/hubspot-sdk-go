// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package crm

import (
	"github.com/stainless-sdks/hubspot-sdk-go/internal/apijson"
	"github.com/stainless-sdks/hubspot-sdk-go/option"
	"github.com/stainless-sdks/hubspot-sdk-go/packages/respjson"
)

// ObjectLibraryService contains methods and other services that help with
// interacting with the hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewObjectLibraryService] method instead.
type ObjectLibraryService struct {
	Options    []option.RequestOption
	Enablement ObjectLibraryEnablementService
}

// NewObjectLibraryService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewObjectLibraryService(opts ...option.RequestOption) (r ObjectLibraryService) {
	r = ObjectLibraryService{}
	r.Options = opts
	r.Enablement = NewObjectLibraryEnablementService(opts...)
	return
}

type ObjectTypeEnablementPublicResponse struct {
	// Whether the object type is enabled or not
	Enablement bool `json:"enablement,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Enablement  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ObjectTypeEnablementPublicResponse) RawJSON() string { return r.JSON.raw }
func (r *ObjectTypeEnablementPublicResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PortalObjectTypeEnablementPublicResponse struct {
	// A map of objectTypeId to whether that object type is enabled or not
	EnablementByObjectTypeID map[string]bool `json:"enablementByObjectTypeId,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		EnablementByObjectTypeID respjson.Field
		ExtraFields              map[string]respjson.Field
		raw                      string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PortalObjectTypeEnablementPublicResponse) RawJSON() string { return r.JSON.raw }
func (r *PortalObjectTypeEnablementPublicResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
