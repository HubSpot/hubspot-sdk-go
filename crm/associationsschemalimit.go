// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package crm

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/HubSpot/hubspot-sdk-go/internal/apijson"
	shimjson "github.com/HubSpot/hubspot-sdk-go/internal/encoding/json"
	"github.com/HubSpot/hubspot-sdk-go/internal/requestconfig"
	"github.com/HubSpot/hubspot-sdk-go/option"
)

// AssociationsSchemaLimitService contains methods and other services that help
// with interacting with the hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAssociationsSchemaLimitService] method instead.
type AssociationsSchemaLimitService struct {
	options []option.RequestOption
}

// NewAssociationsSchemaLimitService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAssociationsSchemaLimitService(opts ...option.RequestOption) (r AssociationsSchemaLimitService) {
	r = AssociationsSchemaLimitService{}
	r.options = opts
	return
}

// Retrieve all configured association limits between objects, which include
// details about how different CRM object types are associated with each other.
func (r *AssociationsSchemaLimitService) List(ctx context.Context, opts ...option.RequestOption) (res *CollectionResponsePublicAssociationDefinitionUserConfigurationNoPaging, err error) {
	opts = slices.Concat(r.options, opts)
	path := "crm/associations/2026-03/definitions/configurations/all"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Batch delete limits that have been defined for association types between two
// object types.
func (r *AssociationsSchemaLimitService) BatchDelete(ctx context.Context, toObjectType string, params AssociationsSchemaLimitBatchDeleteParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if params.FromObjectType == "" {
		err = errors.New("missing required fromObjectType parameter")
		return err
	}
	if toObjectType == "" {
		err = errors.New("missing required toObjectType parameter")
		return err
	}
	path := fmt.Sprintf("crm/associations/2026-03/definitions/configurations/%s/%s/batch/purge", url.PathEscape(params.FromObjectType), url.PathEscape(toObjectType))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, nil, opts...)
	return err
}

// Batch update association limits that have been configured between two object
// types.
func (r *AssociationsSchemaLimitService) BatchUpdate(ctx context.Context, toObjectType string, params AssociationsSchemaLimitBatchUpdateParams, opts ...option.RequestOption) (res *BatchResponsePublicAssociationDefinitionConfigurationUpdateResult, err error) {
	opts = slices.Concat(r.options, opts)
	if params.FromObjectType == "" {
		err = errors.New("missing required fromObjectType parameter")
		return nil, err
	}
	if toObjectType == "" {
		err = errors.New("missing required toObjectType parameter")
		return nil, err
	}
	path := fmt.Sprintf("crm/associations/2026-03/definitions/configurations/%s/%s/batch/update", url.PathEscape(params.FromObjectType), url.PathEscape(toObjectType))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Retrieve the configuration details for associations between two specified CRM
// object types. Use this endpoint to understand limits that have been set for
// specific association types.
func (r *AssociationsSchemaLimitService) GetByObjectTypes(ctx context.Context, toObjectType string, query AssociationsSchemaLimitGetByObjectTypesParams, opts ...option.RequestOption) (res *CollectionResponsePublicAssociationDefinitionUserConfigurationNoPaging, err error) {
	opts = slices.Concat(r.options, opts)
	if query.FromObjectType == "" {
		err = errors.New("missing required fromObjectType parameter")
		return nil, err
	}
	if toObjectType == "" {
		err = errors.New("missing required toObjectType parameter")
		return nil, err
	}
	path := fmt.Sprintf("crm/associations/2026-03/definitions/configurations/%s/%s", url.PathEscape(query.FromObjectType), url.PathEscape(toObjectType))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type AssociationsSchemaLimitBatchDeleteParams struct {
	FromObjectType                  string `path:"fromObjectType" api:"required" json:"-"`
	BatchInputPublicAssociationSpec BatchInputPublicAssociationSpecParam
	paramObj
}

func (r AssociationsSchemaLimitBatchDeleteParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.BatchInputPublicAssociationSpec)
}
func (r *AssociationsSchemaLimitBatchDeleteParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AssociationsSchemaLimitBatchUpdateParams struct {
	FromObjectType                                                  string `path:"fromObjectType" api:"required" json:"-"`
	BatchInputPublicAssociationDefinitionConfigurationUpdateRequest BatchInputPublicAssociationDefinitionConfigurationUpdateRequestParam
	paramObj
}

func (r AssociationsSchemaLimitBatchUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.BatchInputPublicAssociationDefinitionConfigurationUpdateRequest)
}
func (r *AssociationsSchemaLimitBatchUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AssociationsSchemaLimitGetByObjectTypesParams struct {
	FromObjectType string `path:"fromObjectType" api:"required" json:"-"`
	paramObj
}
