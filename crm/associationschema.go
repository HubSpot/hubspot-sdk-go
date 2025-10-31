// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package crm

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/stainless-sdks/hubspot-sdk-go/internal/apijson"
	"github.com/stainless-sdks/hubspot-sdk-go/internal/requestconfig"
	"github.com/stainless-sdks/hubspot-sdk-go/option"
	"github.com/stainless-sdks/hubspot-sdk-go/packages/respjson"
)

// AssociationSchemaService contains methods and other services that help with
// interacting with the Hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAssociationSchemaService] method instead.
type AssociationSchemaService struct {
	Options []option.RequestOption
	V4      AssociationSchemaV4Service
}

// NewAssociationSchemaService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAssociationSchemaService(opts ...option.RequestOption) (r AssociationSchemaService) {
	r = AssociationSchemaService{}
	r.Options = opts
	r.V4 = NewAssociationSchemaV4Service(opts...)
	return
}

// List all the valid association types available between two object types
func (r *AssociationSchemaService) List(ctx context.Context, toObjectType string, query AssociationSchemaListParams, opts ...option.RequestOption) (res *CollectionResponsePublicAssociationDefinitionNoPaging, err error) {
	opts = slices.Concat(r.Options, opts)
	if query.FromObjectType == "" {
		err = errors.New("missing required fromObjectType parameter")
		return
	}
	if toObjectType == "" {
		err = errors.New("missing required toObjectType parameter")
		return
	}
	path := fmt.Sprintf("crm/v3/associations/%s/%s/types", query.FromObjectType, toObjectType)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type CollectionResponsePublicAssociationDefinitionNoPaging struct {
	Results []PublicAssociationDefinition `json:"results,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Results     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CollectionResponsePublicAssociationDefinitionNoPaging) RawJSON() string { return r.JSON.raw }
func (r *CollectionResponsePublicAssociationDefinitionNoPaging) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PublicAssociationDefinition struct {
	ID   string `json:"id,required"`
	Name string `json:"name,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PublicAssociationDefinition) RawJSON() string { return r.JSON.raw }
func (r *PublicAssociationDefinition) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AssociationSchemaListParams struct {
	FromObjectType string `path:"fromObjectType,required" json:"-"`
	paramObj
}
