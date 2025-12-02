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

// AssociationSchemaV4DefinitionService contains methods and other services that
// help with interacting with the hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAssociationSchemaV4DefinitionService] method instead.
type AssociationSchemaV4DefinitionService struct {
	Options []option.RequestOption
}

// NewAssociationSchemaV4DefinitionService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAssociationSchemaV4DefinitionService(opts ...option.RequestOption) (r AssociationSchemaV4DefinitionService) {
	r = AssociationSchemaV4DefinitionService{}
	r.Options = opts
	return
}

func (r *AssociationSchemaV4DefinitionService) NewLabel(ctx context.Context, toObjectType string, params AssociationSchemaV4DefinitionNewLabelParams, opts ...option.RequestOption) (res *CollectionResponseAssociationSpecWithLabel, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.FromObjectType == "" {
		err = errors.New("missing required fromObjectType parameter")
		return
	}
	if toObjectType == "" {
		err = errors.New("missing required toObjectType parameter")
		return
	}
	path := fmt.Sprintf("crm/associations/v4/%s/%s/labels", params.FromObjectType, toObjectType)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

func (r *AssociationSchemaV4DefinitionService) DeleteLabel(ctx context.Context, associationTypeID int64, body AssociationSchemaV4DefinitionDeleteLabelParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if body.FromObjectType == "" {
		err = errors.New("missing required fromObjectType parameter")
		return
	}
	if body.ToObjectType == "" {
		err = errors.New("missing required toObjectType parameter")
		return
	}
	path := fmt.Sprintf("crm/associations/v4/%s/%s/labels/%v", body.FromObjectType, body.ToObjectType, associationTypeID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

func (r *AssociationSchemaV4DefinitionService) ListLabels(ctx context.Context, toObjectType string, query AssociationSchemaV4DefinitionListLabelsParams, opts ...option.RequestOption) (res *CollectionResponseAssociationSpecWithLabel, err error) {
	opts = slices.Concat(r.Options, opts)
	if query.FromObjectType == "" {
		err = errors.New("missing required fromObjectType parameter")
		return
	}
	if toObjectType == "" {
		err = errors.New("missing required toObjectType parameter")
		return
	}
	path := fmt.Sprintf("crm/associations/v4/%s/%s/labels", query.FromObjectType, toObjectType)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

func (r *AssociationSchemaV4DefinitionService) UpdateLabel(ctx context.Context, toObjectType string, params AssociationSchemaV4DefinitionUpdateLabelParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if params.FromObjectType == "" {
		err = errors.New("missing required fromObjectType parameter")
		return
	}
	if toObjectType == "" {
		err = errors.New("missing required toObjectType parameter")
		return
	}
	path := fmt.Sprintf("crm/associations/v4/%s/%s/labels", params.FromObjectType, toObjectType)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, nil, opts...)
	return
}

type AssociationSchemaV4DefinitionNewLabelParams struct {
	FromObjectType                           string `path:"fromObjectType,required" json:"-"`
	PublicAssociationDefinitionCreateRequest PublicAssociationDefinitionCreateRequestParam
	paramObj
}

func (r AssociationSchemaV4DefinitionNewLabelParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.PublicAssociationDefinitionCreateRequest)
}
func (r *AssociationSchemaV4DefinitionNewLabelParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.PublicAssociationDefinitionCreateRequest)
}

type AssociationSchemaV4DefinitionDeleteLabelParams struct {
	FromObjectType string `path:"fromObjectType,required" json:"-"`
	ToObjectType   string `path:"toObjectType,required" json:"-"`
	paramObj
}

type AssociationSchemaV4DefinitionListLabelsParams struct {
	FromObjectType string `path:"fromObjectType,required" json:"-"`
	paramObj
}

type AssociationSchemaV4DefinitionUpdateLabelParams struct {
	FromObjectType                           string `path:"fromObjectType,required" json:"-"`
	PublicAssociationDefinitionUpdateRequest PublicAssociationDefinitionUpdateRequestParam
	paramObj
}

func (r AssociationSchemaV4DefinitionUpdateLabelParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.PublicAssociationDefinitionUpdateRequest)
}
func (r *AssociationSchemaV4DefinitionUpdateLabelParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.PublicAssociationDefinitionUpdateRequest)
}
