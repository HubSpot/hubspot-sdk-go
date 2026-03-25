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

// Defines the type, direction, and details of the relationship between two CRM
// objects.
//
// The properties AssociationCategory, AssociationTypeID are required.
type AssociationSpecParam struct {
	// The category of the association, such as "HUBSPOT_DEFINED".
	//
	// Any of "HUBSPOT_DEFINED", "INTEGRATOR_DEFINED", "USER_DEFINED", "WORK".
	AssociationCategory AssociationSpecAssociationCategory `json:"associationCategory,omitzero" api:"required"`
	// The ID representing the specific type of association.
	AssociationTypeID int64 `json:"associationTypeId" api:"required"`
	paramObj
}

func (r AssociationSpecParam) MarshalJSON() (data []byte, err error) {
	type shadow AssociationSpecParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AssociationSpecParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The category of the association, such as "HUBSPOT_DEFINED".
type AssociationSpecAssociationCategory string

const (
	AssociationSpecAssociationCategoryHubspotDefined    AssociationSpecAssociationCategory = "HUBSPOT_DEFINED"
	AssociationSpecAssociationCategoryIntegratorDefined AssociationSpecAssociationCategory = "INTEGRATOR_DEFINED"
	AssociationSpecAssociationCategoryUserDefined       AssociationSpecAssociationCategory = "USER_DEFINED"
	AssociationSpecAssociationCategoryWork              AssociationSpecAssociationCategory = "WORK"
)

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

type Paging struct {
	// Specifies the paging information needed to retrieve the next set of results in a
	// paginated API response
	Next NextPage `json:"next"`
	// specifies the paging information needed to retrieve the previous set of results
	// in a paginated API response
	Prev PreviousPage `json:"prev"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Next        respjson.Field
		Prev        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Paging) RawJSON() string { return r.JSON.raw }
func (r *Paging) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// specifies the paging information needed to retrieve the previous set of results
// in a paginated API response
type PreviousPage struct {
	// A paging cursor token for retrieving previous pages.
	Before string `json:"before" api:"required"`
	// A URL that can be used to retrieve the previous pages' results.
	Link string `json:"link"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Before      respjson.Field
		Link        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PreviousPage) RawJSON() string { return r.JSON.raw }
func (r *PreviousPage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Contains the Id of a Public Object
//
// The property ID is required.
type PublicObjectIDParam struct {
	// The unique ID of the object.
	ID string `json:"id" api:"required"`
	paramObj
}

func (r PublicObjectIDParam) MarshalJSON() (data []byte, err error) {
	type shadow PublicObjectIDParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PublicObjectIDParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Ye olde error
type StandardError struct {
	// Error category.
	Category string `json:"category" api:"required"`
	// Error context.
	Context map[string][]string `json:"context" api:"required"`
	// List of error details.
	Errors []ErrorDetail `json:"errors" api:"required"`
	// Error links.
	Links map[string]string `json:"links" api:"required"`
	// Error message.
	Message string `json:"message" api:"required"`
	// Error status.
	Status string `json:"status" api:"required"`
	// Error ID.
	ID string `json:"id"`
	// Error subcategory.
	SubCategory any `json:"subCategory"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Category    respjson.Field
		Context     respjson.Field
		Errors      respjson.Field
		Links       respjson.Field
		Message     respjson.Field
		Status      respjson.Field
		ID          respjson.Field
		SubCategory respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r StandardError) RawJSON() string { return r.JSON.raw }
func (r *StandardError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
