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

// AssociationsSchemaLabelService contains methods and other services that help
// with interacting with the hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAssociationsSchemaLabelService] method instead.
type AssociationsSchemaLabelService struct {
	options []option.RequestOption
}

// NewAssociationsSchemaLabelService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAssociationsSchemaLabelService(opts ...option.RequestOption) (r AssociationsSchemaLabelService) {
	r = AssociationsSchemaLabelService{}
	r.options = opts
	return
}

// Batch configure association limits between two object types.
func (r *AssociationsSchemaLabelService) BatchNew(ctx context.Context, toObjectType string, params AssociationsSchemaLabelBatchNewParams, opts ...option.RequestOption) (res *BatchResponsePublicAssociationDefinitionUserConfiguration, err error) {
	opts = slices.Concat(r.options, opts)
	if params.FromObjectType == "" {
		err = errors.New("missing required fromObjectType parameter")
		return nil, err
	}
	if toObjectType == "" {
		err = errors.New("missing required toObjectType parameter")
		return nil, err
	}
	path := fmt.Sprintf("crm/associations/2026-03/definitions/configurations/%s/%s/batch/create", url.PathEscape(params.FromObjectType), url.PathEscape(toObjectType))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Create a new label that describes the relationship between two specified CRM
// object types. This can help in categorizing and managing associations more
// effectively.
func (r *AssociationsSchemaLabelService) NewLabel(ctx context.Context, toObjectType string, params AssociationsSchemaLabelNewLabelParams, opts ...option.RequestOption) (res *CollectionResponseAssociationSpecWithLabelNoPaging, err error) {
	opts = slices.Concat(r.options, opts)
	if params.FromObjectType == "" {
		err = errors.New("missing required fromObjectType parameter")
		return nil, err
	}
	if toObjectType == "" {
		err = errors.New("missing required toObjectType parameter")
		return nil, err
	}
	path := fmt.Sprintf("crm/associations/2026-03/%s/%s/labels", url.PathEscape(params.FromObjectType), url.PathEscape(toObjectType))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Remove a specific label from the association between two CRM object types.
func (r *AssociationsSchemaLabelService) DeleteLabel(ctx context.Context, associationTypeID int64, body AssociationsSchemaLabelDeleteLabelParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if body.FromObjectType == "" {
		err = errors.New("missing required fromObjectType parameter")
		return err
	}
	if body.ToObjectType == "" {
		err = errors.New("missing required toObjectType parameter")
		return err
	}
	path := fmt.Sprintf("crm/associations/2026-03/%s/%s/labels/%v", url.PathEscape(body.FromObjectType), url.PathEscape(body.ToObjectType), associationTypeID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Retrieve all labels that describe the relationships between two specified CRM
// object types. These labels provide context about the nature of the associations.
func (r *AssociationsSchemaLabelService) ListLabels(ctx context.Context, toObjectType string, query AssociationsSchemaLabelListLabelsParams, opts ...option.RequestOption) (res *CollectionResponseAssociationSpecWithLabelNoPaging, err error) {
	opts = slices.Concat(r.options, opts)
	if query.FromObjectType == "" {
		err = errors.New("missing required fromObjectType parameter")
		return nil, err
	}
	if toObjectType == "" {
		err = errors.New("missing required toObjectType parameter")
		return nil, err
	}
	path := fmt.Sprintf("crm/associations/2026-03/%s/%s/labels", url.PathEscape(query.FromObjectType), url.PathEscape(toObjectType))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Update an existing label that describes the relationship between two specified
// CRM object types. This allows for modifications to existing association labels
// to better reflect the nature of the relationship.
func (r *AssociationsSchemaLabelService) UpdateLabel(ctx context.Context, toObjectType string, params AssociationsSchemaLabelUpdateLabelParams, opts ...option.RequestOption) (err error) {
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
	path := fmt.Sprintf("crm/associations/2026-03/%s/%s/labels", url.PathEscape(params.FromObjectType), url.PathEscape(toObjectType))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, nil, opts...)
	return err
}

type AssociationsSchemaLabelBatchNewParams struct {
	FromObjectType                                                  string `path:"fromObjectType" api:"required" json:"-"`
	BatchInputPublicAssociationDefinitionConfigurationCreateRequest BatchInputPublicAssociationDefinitionConfigurationCreateRequestParam
	paramObj
}

func (r AssociationsSchemaLabelBatchNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.BatchInputPublicAssociationDefinitionConfigurationCreateRequest)
}
func (r *AssociationsSchemaLabelBatchNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AssociationsSchemaLabelNewLabelParams struct {
	FromObjectType                           string `path:"fromObjectType" api:"required" json:"-"`
	PublicAssociationDefinitionCreateRequest PublicAssociationDefinitionCreateRequestParam
	paramObj
}

func (r AssociationsSchemaLabelNewLabelParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.PublicAssociationDefinitionCreateRequest)
}
func (r *AssociationsSchemaLabelNewLabelParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AssociationsSchemaLabelDeleteLabelParams struct {
	FromObjectType string `path:"fromObjectType" api:"required" json:"-"`
	ToObjectType   string `path:"toObjectType" api:"required" json:"-"`
	paramObj
}

type AssociationsSchemaLabelListLabelsParams struct {
	FromObjectType string `path:"fromObjectType" api:"required" json:"-"`
	paramObj
}

type AssociationsSchemaLabelUpdateLabelParams struct {
	FromObjectType                           string `path:"fromObjectType" api:"required" json:"-"`
	PublicAssociationDefinitionUpdateRequest PublicAssociationDefinitionUpdateRequestParam
	paramObj
}

func (r AssociationsSchemaLabelUpdateLabelParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.PublicAssociationDefinitionUpdateRequest)
}
func (r *AssociationsSchemaLabelUpdateLabelParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
