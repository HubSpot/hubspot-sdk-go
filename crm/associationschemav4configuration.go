// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package crm

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"slices"

	shimjson "github.com/stainless-sdks/hubspot-sdk-go/internal/encoding/json"
	"github.com/stainless-sdks/hubspot-sdk-go/internal/requestconfig"
	"github.com/stainless-sdks/hubspot-sdk-go/option"
)

// AssociationSchemaV4ConfigurationService contains methods and other services that
// help with interacting with the hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAssociationSchemaV4ConfigurationService] method instead.
type AssociationSchemaV4ConfigurationService struct {
	Options []option.RequestOption
}

// NewAssociationSchemaV4ConfigurationService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAssociationSchemaV4ConfigurationService(opts ...option.RequestOption) (r AssociationSchemaV4ConfigurationService) {
	r = AssociationSchemaV4ConfigurationService{}
	r.Options = opts
	return
}

func (r *AssociationSchemaV4ConfigurationService) List(ctx context.Context, opts ...option.RequestOption) (res *CollectionResponsePublicAssociationDefinitionUserConfiguration, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "crm/associations/v4/definitions/configurations/all"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

func (r *AssociationSchemaV4ConfigurationService) BatchNew(ctx context.Context, toObjectType string, params AssociationSchemaV4ConfigurationBatchNewParams, opts ...option.RequestOption) (res *BatchResponsePublicAssociationDefinitionUserConfiguration, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.FromObjectType == "" {
		err = errors.New("missing required fromObjectType parameter")
		return
	}
	if toObjectType == "" {
		err = errors.New("missing required toObjectType parameter")
		return
	}
	path := fmt.Sprintf("crm/associations/v4/definitions/configurations/%s/%s/batch/create", params.FromObjectType, toObjectType)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

func (r *AssociationSchemaV4ConfigurationService) BatchDelete(ctx context.Context, toObjectType string, params AssociationSchemaV4ConfigurationBatchDeleteParams, opts ...option.RequestOption) (res *BatchResponseVoid, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.FromObjectType == "" {
		err = errors.New("missing required fromObjectType parameter")
		return
	}
	if toObjectType == "" {
		err = errors.New("missing required toObjectType parameter")
		return
	}
	path := fmt.Sprintf("crm/associations/v4/definitions/configurations/%s/%s/batch/purge", params.FromObjectType, toObjectType)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

func (r *AssociationSchemaV4ConfigurationService) BatchUpdate(ctx context.Context, toObjectType string, params AssociationSchemaV4ConfigurationBatchUpdateParams, opts ...option.RequestOption) (res *BatchResponsePublicAssociationDefinitionConfigurationUpdateResult, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.FromObjectType == "" {
		err = errors.New("missing required fromObjectType parameter")
		return
	}
	if toObjectType == "" {
		err = errors.New("missing required toObjectType parameter")
		return
	}
	path := fmt.Sprintf("crm/associations/v4/definitions/configurations/%s/%s/batch/update", params.FromObjectType, toObjectType)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

func (r *AssociationSchemaV4ConfigurationService) GetByObjectTypes(ctx context.Context, toObjectType string, query AssociationSchemaV4ConfigurationGetByObjectTypesParams, opts ...option.RequestOption) (res *CollectionResponsePublicAssociationDefinitionUserConfiguration, err error) {
	opts = slices.Concat(r.Options, opts)
	if query.FromObjectType == "" {
		err = errors.New("missing required fromObjectType parameter")
		return
	}
	if toObjectType == "" {
		err = errors.New("missing required toObjectType parameter")
		return
	}
	path := fmt.Sprintf("crm/associations/v4/definitions/configurations/%s/%s", query.FromObjectType, toObjectType)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type AssociationSchemaV4ConfigurationBatchNewParams struct {
	FromObjectType                                                  string `path:"fromObjectType,required" json:"-"`
	BatchInputPublicAssociationDefinitionConfigurationCreateRequest BatchInputPublicAssociationDefinitionConfigurationCreateRequestParam
	paramObj
}

func (r AssociationSchemaV4ConfigurationBatchNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.BatchInputPublicAssociationDefinitionConfigurationCreateRequest)
}
func (r *AssociationSchemaV4ConfigurationBatchNewParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.BatchInputPublicAssociationDefinitionConfigurationCreateRequest)
}

type AssociationSchemaV4ConfigurationBatchDeleteParams struct {
	FromObjectType                  string `path:"fromObjectType,required" json:"-"`
	BatchInputPublicAssociationSpec BatchInputPublicAssociationSpecParam
	paramObj
}

func (r AssociationSchemaV4ConfigurationBatchDeleteParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.BatchInputPublicAssociationSpec)
}
func (r *AssociationSchemaV4ConfigurationBatchDeleteParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.BatchInputPublicAssociationSpec)
}

type AssociationSchemaV4ConfigurationBatchUpdateParams struct {
	FromObjectType                                                  string `path:"fromObjectType,required" json:"-"`
	BatchInputPublicAssociationDefinitionConfigurationUpdateRequest BatchInputPublicAssociationDefinitionConfigurationUpdateRequestParam
	paramObj
}

func (r AssociationSchemaV4ConfigurationBatchUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.BatchInputPublicAssociationDefinitionConfigurationUpdateRequest)
}
func (r *AssociationSchemaV4ConfigurationBatchUpdateParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.BatchInputPublicAssociationDefinitionConfigurationUpdateRequest)
}

type AssociationSchemaV4ConfigurationGetByObjectTypesParams struct {
	FromObjectType string `path:"fromObjectType,required" json:"-"`
	paramObj
}
