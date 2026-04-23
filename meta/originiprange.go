// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package meta

import (
	"context"
	"net/http"
	"net/url"
	"slices"

	"github.com/stainless-sdks/hubspot-sdk-go/internal/apiquery"
	"github.com/stainless-sdks/hubspot-sdk-go/internal/requestconfig"
	"github.com/stainless-sdks/hubspot-sdk-go/option"
)

// OriginIPRangeService contains methods and other services that help with
// interacting with the hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewOriginIPRangeService] method instead.
type OriginIPRangeService struct {
	options []option.RequestOption
}

// NewOriginIPRangeService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewOriginIPRangeService(opts ...option.RequestOption) (r OriginIPRangeService) {
	r = OriginIPRangeService{}
	r.options = opts
	return
}

// Retrieve a collection of IP ranges associated with specific services and
// directions, such as `EMAIL`, `API`, `DNS`, or `WEB_SCRAPING`. The response
// includes details like CIDR notation, description, and the direction of IP
// traffic.
func (r *OriginIPRangeService) List(ctx context.Context, query OriginIPRangeListParams, opts ...option.RequestOption) (res *CollectionResponseIPRangeNoPaging, err error) {
	opts = slices.Concat(r.options, opts)
	path := "meta/network-origins/2026-03/ip-ranges"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Retrieve a simplified list of IP ranges for specified services and directions in
// plain text format. This endpoint provides a straightforward representation of IP
// ranges without additional metadata.
func (r *OriginIPRangeService) ListSimple(ctx context.Context, query OriginIPRangeListSimpleParams, opts ...option.RequestOption) (res *string, err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/plain")}, opts...)
	path := "meta/network-origins/2026-03/ip-ranges/simple"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

type OriginIPRangeListParams struct {
	// Any of "INGRESS", "EGRESS".
	Direction []string `query:"direction,omitzero" json:"-"`
	// Any of "EMAIL", "API", "DNS", "WEB_SCRAPING", "TEST_SERVICE".
	Service []string `query:"service,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [OriginIPRangeListParams]'s query parameters as
// `url.Values`.
func (r OriginIPRangeListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type OriginIPRangeListSimpleParams struct {
	// Any of "INGRESS", "EGRESS".
	Direction []string `query:"direction,omitzero" json:"-"`
	// Any of "EMAIL", "API", "DNS", "WEB_SCRAPING", "TEST_SERVICE".
	Service []string `query:"service,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [OriginIPRangeListSimpleParams]'s query parameters as
// `url.Values`.
func (r OriginIPRangeListSimpleParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
