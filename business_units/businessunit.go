// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package business_units

import (
	"github.com/HubSpot/hubspot-sdk-go/internal/apijson"
	"github.com/HubSpot/hubspot-sdk-go/option"
	"github.com/HubSpot/hubspot-sdk-go/packages/respjson"
)

// BusinessUnitService contains methods and other services that help with
// interacting with the hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBusinessUnitService] method instead.
type BusinessUnitService struct {
	options             []option.RequestOption
	BusinessUnitEntries BusinessUnitEntryService
}

// NewBusinessUnitService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewBusinessUnitService(opts ...option.RequestOption) (r BusinessUnitService) {
	r = BusinessUnitService{}
	r.options = opts
	r.BusinessUnitEntries = NewBusinessUnitEntryService(opts...)
	return
}

type CollectionResponsePublicBusinessUnitNoPaging struct {
	// The collection of Business Units
	Results []PublicBusinessUnit `json:"results" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Results     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CollectionResponsePublicBusinessUnitNoPaging) RawJSON() string { return r.JSON.raw }
func (r *CollectionResponsePublicBusinessUnitNoPaging) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PublicBusinessUnit struct {
	// The Business Unit's unique ID
	ID string `json:"id" api:"required"`
	// The Business Unit's name
	Name         string                         `json:"name" api:"required"`
	LogoMetadata PublicBusinessUnitLogoMetadata `json:"logoMetadata"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		Name         respjson.Field
		LogoMetadata respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PublicBusinessUnit) RawJSON() string { return r.JSON.raw }
func (r *PublicBusinessUnit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PublicBusinessUnitLogoMetadata struct {
	// The logo's alt text
	LogoAltText string `json:"logoAltText"`
	// The logo's url
	LogoURL string `json:"logoUrl"`
	// The logo's resized url
	ResizedURL string `json:"resizedUrl"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		LogoAltText respjson.Field
		LogoURL     respjson.Field
		ResizedURL  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PublicBusinessUnitLogoMetadata) RawJSON() string { return r.JSON.raw }
func (r *PublicBusinessUnitLogoMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
