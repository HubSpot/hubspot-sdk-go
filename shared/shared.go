// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package shared

import (
	"github.com/stainless-sdks/hubspot-sdk-go/internal/apijson"
	"github.com/stainless-sdks/hubspot-sdk-go/packages/param"
	"github.com/stainless-sdks/hubspot-sdk-go/packages/respjson"
)

// aliased to make [param.APIUnion] private when embedding
type paramUnion = param.APIUnion

// aliased to make [param.APIObject] private when embedding
type paramObj = param.APIObject

type ErrorDetail struct {
	// A human readable message describing the error along with remediation steps where
	// appropriate
	Message string `json:"message" api:"required"`
	// The status code associated with the error detail
	Code string `json:"code"`
	// Context about the error condition
	Context map[string][]string `json:"context"`
	// The name of the field or parameter in which the error was found.
	In string `json:"in"`
	// A specific category that contains more specific detail about the error.
	SubCategory string `json:"subCategory"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Message     respjson.Field
		Code        respjson.Field
		Context     respjson.Field
		In          respjson.Field
		SubCategory respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ErrorDetail) RawJSON() string { return r.JSON.raw }
func (r *ErrorDetail) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ForwardPaging struct {
	// Specifies the paging information needed to retrieve the next set of results in a
	// paginated API response
	Next NextPage `json:"next"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Next        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ForwardPaging) RawJSON() string { return r.JSON.raw }
func (r *ForwardPaging) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Specifies the paging information needed to retrieve the next set of results in a
// paginated API response
type NextPage struct {
	// A paging cursor token for retrieving subsequent pages.
	After string `json:"after" api:"required"`
	// A URL that can be used to retrieve the next page results.
	Link string `json:"link"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		After       respjson.Field
		Link        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NextPage) RawJSON() string { return r.JSON.raw }
func (r *NextPage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
