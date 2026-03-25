// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package meta

import (
	"github.com/stainless-sdks/hubspot-sdk-go/internal/apijson"
	"github.com/stainless-sdks/hubspot-sdk-go/option"
	"github.com/stainless-sdks/hubspot-sdk-go/packages/respjson"
)

// OriginService contains methods and other services that help with interacting
// with the hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewOriginService] method instead.
type OriginService struct {
	Options  []option.RequestOption
	IPRanges OriginIPRangeService
}

// NewOriginService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewOriginService(opts ...option.RequestOption) (r OriginService) {
	r = OriginService{}
	r.Options = opts
	r.IPRanges = NewOriginIPRangeService(opts...)
	return
}

type CollectionResponseIPRangeNoPaging struct {
	// An array of IpRange objects, each representing a specific IP range with
	// associated details such as CIDR, direction, service, and description.
	Results []IPRange `json:"results" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Results     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CollectionResponseIPRangeNoPaging) RawJSON() string { return r.JSON.raw }
func (r *CollectionResponseIPRangeNoPaging) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IPRange struct {
	// The CIDR notation representing the IP range.
	Cidr string `json:"cidr" api:"required"`
	// A description of the IP range.
	Description string `json:"description" api:"required"`
	// The direction of the IP traffic, which can be INGRESS or EGRESS.
	//
	// Any of "EGRESS", "INGRESS".
	Direction IPRangeDirection `json:"direction" api:"required"`
	// The service associated with the IP range, such as EMAIL, API, DNS, or
	// WEB_SCRAPING.
	//
	// Any of "API", "DNS", "EMAIL", "WEB_SCRAPING".
	Service IPRangeService `json:"service" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Cidr        respjson.Field
		Description respjson.Field
		Direction   respjson.Field
		Service     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IPRange) RawJSON() string { return r.JSON.raw }
func (r *IPRange) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The direction of the IP traffic, which can be INGRESS or EGRESS.
type IPRangeDirection string

const (
	IPRangeDirectionEgress  IPRangeDirection = "EGRESS"
	IPRangeDirectionIngress IPRangeDirection = "INGRESS"
)

// The service associated with the IP range, such as EMAIL, API, DNS, or
// WEB_SCRAPING.
type IPRangeService string

const (
	IPRangeServiceAPI         IPRangeService = "API"
	IPRangeServiceDNS         IPRangeService = "DNS"
	IPRangeServiceEmail       IPRangeService = "EMAIL"
	IPRangeServiceWebScraping IPRangeService = "WEB_SCRAPING"
)
