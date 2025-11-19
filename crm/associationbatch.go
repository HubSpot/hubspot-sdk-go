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
	"github.com/stainless-sdks/hubspot-sdk-go/shared"
)

// AssociationBatchService contains methods and other services that help with
// interacting with the Hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAssociationBatchService] method instead.
type AssociationBatchService struct {
	Options []option.RequestOption
}

// NewAssociationBatchService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAssociationBatchService(opts ...option.RequestOption) (r AssociationBatchService) {
	r = AssociationBatchService{}
	r.Options = opts
	return
}

func (r *AssociationBatchService) New(ctx context.Context, toObjectType string, params AssociationBatchNewParams, opts ...option.RequestOption) (res *BatchResponsePublicAssociation, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.FromObjectType == "" {
		err = errors.New("missing required fromObjectType parameter")
		return
	}
	if toObjectType == "" {
		err = errors.New("missing required toObjectType parameter")
		return
	}
	path := fmt.Sprintf("crm/v3/associations/%s/%s/batch/create", params.FromObjectType, toObjectType)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

func (r *AssociationBatchService) Delete(ctx context.Context, toObjectType string, params AssociationBatchDeleteParams, opts ...option.RequestOption) (err error) {
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
	path := fmt.Sprintf("crm/v3/associations/%s/%s/batch/archive", params.FromObjectType, toObjectType)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, nil, opts...)
	return
}

func (r *AssociationBatchService) Get(ctx context.Context, toObjectType string, params AssociationBatchGetParams, opts ...option.RequestOption) (res *BatchResponsePublicAssociationMulti, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.FromObjectType == "" {
		err = errors.New("missing required fromObjectType parameter")
		return
	}
	if toObjectType == "" {
		err = errors.New("missing required toObjectType parameter")
		return
	}
	path := fmt.Sprintf("crm/v3/associations/%s/%s/batch/read", params.FromObjectType, toObjectType)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

type AssociationBatchNewParams struct {
	FromObjectType              string `path:"fromObjectType,required" json:"-"`
	BatchInputPublicAssociation BatchInputPublicAssociationParam
	paramObj
}

func (r AssociationBatchNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.BatchInputPublicAssociation)
}
func (r *AssociationBatchNewParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.BatchInputPublicAssociation)
}

type AssociationBatchDeleteParams struct {
	FromObjectType              string `path:"fromObjectType,required" json:"-"`
	BatchInputPublicAssociation BatchInputPublicAssociationParam
	paramObj
}

func (r AssociationBatchDeleteParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.BatchInputPublicAssociation)
}
func (r *AssociationBatchDeleteParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.BatchInputPublicAssociation)
}

type AssociationBatchGetParams struct {
	FromObjectType           string `path:"fromObjectType,required" json:"-"`
	BatchInputPublicObjectID shared.BatchInputPublicObjectIDParam
	paramObj
}

func (r AssociationBatchGetParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.BatchInputPublicObjectID)
}
func (r *AssociationBatchGetParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.BatchInputPublicObjectID)
}
