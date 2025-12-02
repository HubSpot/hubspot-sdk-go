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

// AssociationV4BatchService contains methods and other services that help with
// interacting with the hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAssociationV4BatchService] method instead.
type AssociationV4BatchService struct {
	Options []option.RequestOption
}

// NewAssociationV4BatchService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAssociationV4BatchService(opts ...option.RequestOption) (r AssociationV4BatchService) {
	r = AssociationV4BatchService{}
	r.Options = opts
	return
}

// Batch create associations for objects
func (r *AssociationV4BatchService) New(ctx context.Context, toObjectType string, params AssociationV4BatchNewParams, opts ...option.RequestOption) (res *BatchResponseLabelsBetweenObjectPair, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.FromObjectType == "" {
		err = errors.New("missing required fromObjectType parameter")
		return
	}
	if toObjectType == "" {
		err = errors.New("missing required toObjectType parameter")
		return
	}
	path := fmt.Sprintf("crm/v4/associations/%s/%s/batch/create", params.FromObjectType, toObjectType)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Batch delete associations for objects
func (r *AssociationV4BatchService) Delete(ctx context.Context, toObjectType string, params AssociationV4BatchDeleteParams, opts ...option.RequestOption) (res *BatchResponseVoid, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.FromObjectType == "" {
		err = errors.New("missing required fromObjectType parameter")
		return
	}
	if toObjectType == "" {
		err = errors.New("missing required toObjectType parameter")
		return
	}
	path := fmt.Sprintf("crm/v4/associations/%s/%s/batch/archive", params.FromObjectType, toObjectType)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Create the default (most generic) association type between two object types
func (r *AssociationV4BatchService) NewDefault(ctx context.Context, toObjectType string, params AssociationV4BatchNewDefaultParams, opts ...option.RequestOption) (res *BatchResponsePublicDefaultAssociation, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.FromObjectType == "" {
		err = errors.New("missing required fromObjectType parameter")
		return
	}
	if toObjectType == "" {
		err = errors.New("missing required toObjectType parameter")
		return
	}
	path := fmt.Sprintf("crm/v4/associations/%s/%s/batch/associate/default", params.FromObjectType, toObjectType)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Batch delete specific association labels for objects. Deleting an unlabeled
// association will also delete all labeled associations between those two objects
func (r *AssociationV4BatchService) DeleteLabels(ctx context.Context, toObjectType string, params AssociationV4BatchDeleteLabelsParams, opts ...option.RequestOption) (res *BatchResponseVoid, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.FromObjectType == "" {
		err = errors.New("missing required fromObjectType parameter")
		return
	}
	if toObjectType == "" {
		err = errors.New("missing required toObjectType parameter")
		return
	}
	path := fmt.Sprintf("crm/v4/associations/%s/%s/batch/labels/archive", params.FromObjectType, toObjectType)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Batch read associations for objects to specific object type. The 'after' field
// in a returned paging object can be added alongside the 'id' to retrieve the next
// page of associations from that objectId. The 'link' field is deprecated and
// should be ignored. Note: The 'paging' field will only be present if there are
// more pages and absent otherwise.
func (r *AssociationV4BatchService) Get(ctx context.Context, toObjectType string, params AssociationV4BatchGetParams, opts ...option.RequestOption) (res *BatchResponsePublicAssociationMultiWithLabel, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.FromObjectType == "" {
		err = errors.New("missing required fromObjectType parameter")
		return
	}
	if toObjectType == "" {
		err = errors.New("missing required toObjectType parameter")
		return
	}
	path := fmt.Sprintf("crm/v4/associations/%s/%s/batch/read", params.FromObjectType, toObjectType)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

type AssociationV4BatchNewParams struct {
	FromObjectType                       string `path:"fromObjectType,required" json:"-"`
	BatchInputPublicAssociationMultiPost BatchInputPublicAssociationMultiPostParam
	paramObj
}

func (r AssociationV4BatchNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.BatchInputPublicAssociationMultiPost)
}
func (r *AssociationV4BatchNewParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.BatchInputPublicAssociationMultiPost)
}

type AssociationV4BatchDeleteParams struct {
	FromObjectType                          string `path:"fromObjectType,required" json:"-"`
	BatchInputPublicAssociationMultiArchive BatchInputPublicAssociationMultiArchiveParam
	paramObj
}

func (r AssociationV4BatchDeleteParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.BatchInputPublicAssociationMultiArchive)
}
func (r *AssociationV4BatchDeleteParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.BatchInputPublicAssociationMultiArchive)
}

type AssociationV4BatchNewDefaultParams struct {
	FromObjectType                              string `path:"fromObjectType,required" json:"-"`
	BatchInputPublicDefaultAssociationMultiPost BatchInputPublicDefaultAssociationMultiPostParam
	paramObj
}

func (r AssociationV4BatchNewDefaultParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.BatchInputPublicDefaultAssociationMultiPost)
}
func (r *AssociationV4BatchNewDefaultParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.BatchInputPublicDefaultAssociationMultiPost)
}

type AssociationV4BatchDeleteLabelsParams struct {
	FromObjectType                       string `path:"fromObjectType,required" json:"-"`
	BatchInputPublicAssociationMultiPost BatchInputPublicAssociationMultiPostParam
	paramObj
}

func (r AssociationV4BatchDeleteLabelsParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.BatchInputPublicAssociationMultiPost)
}
func (r *AssociationV4BatchDeleteLabelsParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.BatchInputPublicAssociationMultiPost)
}

type AssociationV4BatchGetParams struct {
	FromObjectType                                string `path:"fromObjectType,required" json:"-"`
	BatchInputPublicFetchAssociationsBatchRequest BatchInputPublicFetchAssociationsBatchRequestParam
	paramObj
}

func (r AssociationV4BatchGetParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.BatchInputPublicFetchAssociationsBatchRequest)
}
func (r *AssociationV4BatchGetParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.BatchInputPublicFetchAssociationsBatchRequest)
}
