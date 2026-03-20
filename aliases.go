// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package hubspotsdk

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

// This is an alias to an internal type.
type ErrorDetail = shared.ErrorDetail

// This is an alias to an internal type.
type ForwardPaging = shared.ForwardPaging

// Specifies the paging information needed to retrieve the next set of results in a
// paginated API response
//
// This is an alias to an internal type.
type NextPage = shared.NextPage
