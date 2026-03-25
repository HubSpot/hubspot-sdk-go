// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package scheduler

import (
	"github.com/stainless-sdks/hubspot-sdk-go/internal/apierror"
	"github.com/stainless-sdks/hubspot-sdk-go/packages/param"
	"github.com/stainless-sdks/hubspot-sdk-go/shared"
)

// aliased to make [param.APIUnion] private when embedding
type paramUnion = param.APIUnion

// aliased to make [param.APIObject] private when embedding
type paramObj = param.APIObject

type Error = apierror.Error

// Defines the type, direction, and details of the relationship between two CRM
// objects.
//
// This is an alias to an internal type.
type AssociationSpecParam = shared.AssociationSpecParam

// The category of the association, such as "HUBSPOT_DEFINED".
//
// This is an alias to an internal type.
type AssociationSpecAssociationCategory = shared.AssociationSpecAssociationCategory

// Equals "HUBSPOT_DEFINED"
const AssociationSpecAssociationCategoryHubspotDefined = shared.AssociationSpecAssociationCategoryHubspotDefined

// Equals "INTEGRATOR_DEFINED"
const AssociationSpecAssociationCategoryIntegratorDefined = shared.AssociationSpecAssociationCategoryIntegratorDefined

// Equals "USER_DEFINED"
const AssociationSpecAssociationCategoryUserDefined = shared.AssociationSpecAssociationCategoryUserDefined

// Equals "WORK"
const AssociationSpecAssociationCategoryWork = shared.AssociationSpecAssociationCategoryWork

// This is an alias to an internal type.
type ErrorDetail = shared.ErrorDetail

// This is an alias to an internal type.
type ForwardPaging = shared.ForwardPaging

// Specifies the paging information needed to retrieve the next set of results in a
// paginated API response
//
// This is an alias to an internal type.
type NextPage = shared.NextPage

// This is an alias to an internal type.
type Paging = shared.Paging

// specifies the paging information needed to retrieve the previous set of results
// in a paginated API response
//
// This is an alias to an internal type.
type PreviousPage = shared.PreviousPage

// Contains the Id of a Public Object
//
// This is an alias to an internal type.
type PublicObjectIDParam = shared.PublicObjectIDParam

// Ye olde error
//
// This is an alias to an internal type.
type StandardError = shared.StandardError
