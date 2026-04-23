// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package crm

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/stainless-sdks/hubspot-sdk-go/internal/apijson"
	shimjson "github.com/stainless-sdks/hubspot-sdk-go/internal/encoding/json"
	"github.com/stainless-sdks/hubspot-sdk-go/internal/requestconfig"
	"github.com/stainless-sdks/hubspot-sdk-go/option"
)

// AssociationBatchService contains methods and other services that help with
// interacting with the hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAssociationBatchService] method instead.
type AssociationBatchService struct {
	options []option.RequestOption
}

// NewAssociationBatchService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAssociationBatchService(opts ...option.RequestOption) (r AssociationBatchService) {
	r = AssociationBatchService{}
	r.options = opts
	return
}

func (r *AssociationBatchService) New(ctx context.Context, toObjectID string, body AssociationBatchNewParams, opts ...option.RequestOption) (res *BatchResponsePublicDefaultAssociation, err error) {
	opts = slices.Concat(r.options, opts)
	if body.FromObjectType == "" {
		err = errors.New("missing required fromObjectType parameter")
		return nil, err
	}
	if body.FromObjectID == "" {
		err = errors.New("missing required fromObjectId parameter")
		return nil, err
	}
	if body.ToObjectType == "" {
		err = errors.New("missing required toObjectType parameter")
		return nil, err
	}
	if toObjectID == "" {
		err = errors.New("missing required toObjectId parameter")
		return nil, err
	}
	path := fmt.Sprintf("crm/objects/2026-03/%s/%s/associations/default/%s/%s", url.PathEscape(body.FromObjectType), url.PathEscape(body.FromObjectID), url.PathEscape(body.ToObjectType), url.PathEscape(toObjectID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, nil, &res, opts...)
	return res, err
}

// Batch delete associations for objects
func (r *AssociationBatchService) Delete(ctx context.Context, toObjectType string, params AssociationBatchDeleteParams, opts ...option.RequestOption) (err error) {
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
	path := fmt.Sprintf("crm/associations/2026-03/%s/%s/batch/archive", url.PathEscape(params.FromObjectType), url.PathEscape(toObjectType))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, nil, opts...)
	return err
}

// Create the default (most generic) association type between two object types
func (r *AssociationBatchService) NewDefault(ctx context.Context, toObjectType string, params AssociationBatchNewDefaultParams, opts ...option.RequestOption) (res *BatchResponsePublicDefaultAssociation, err error) {
	opts = slices.Concat(r.options, opts)
	if params.FromObjectType == "" {
		err = errors.New("missing required fromObjectType parameter")
		return nil, err
	}
	if toObjectType == "" {
		err = errors.New("missing required toObjectType parameter")
		return nil, err
	}
	path := fmt.Sprintf("crm/associations/2026-03/%s/%s/batch/associate/default", url.PathEscape(params.FromObjectType), url.PathEscape(toObjectType))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Batch delete specific association labels for objects. Deleting an unlabeled
// association will also delete all labeled associations between those two objects
func (r *AssociationBatchService) DeleteLabels(ctx context.Context, toObjectType string, params AssociationBatchDeleteLabelsParams, opts ...option.RequestOption) (err error) {
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
	path := fmt.Sprintf("crm/associations/2026-03/%s/%s/batch/labels/archive", url.PathEscape(params.FromObjectType), url.PathEscape(toObjectType))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, nil, opts...)
	return err
}

// Batch read associations for objects to specific object type. The 'after' field
// in a returned paging object can be added alongside the 'id' to retrieve the next
// page of associations from that objectId. The 'link' field is deprecated and
// should be ignored. Note: The 'paging' field will only be present if there are
// more pages and absent otherwise.
func (r *AssociationBatchService) Get(ctx context.Context, toObjectType string, params AssociationBatchGetParams, opts ...option.RequestOption) (res *BatchResponsePublicAssociationMultiWithLabel, err error) {
	opts = slices.Concat(r.options, opts)
	if params.FromObjectType == "" {
		err = errors.New("missing required fromObjectType parameter")
		return nil, err
	}
	if toObjectType == "" {
		err = errors.New("missing required toObjectType parameter")
		return nil, err
	}
	path := fmt.Sprintf("crm/associations/2026-03/%s/%s/batch/read", url.PathEscape(params.FromObjectType), url.PathEscape(toObjectType))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

type AssociationBatchNewParams struct {
	FromObjectType string `path:"fromObjectType" api:"required" json:"-"`
	FromObjectID   string `path:"fromObjectId" api:"required" json:"-"`
	ToObjectType   string `path:"toObjectType" api:"required" json:"-"`
	paramObj
}

type AssociationBatchDeleteParams struct {
	FromObjectType                          string `path:"fromObjectType" api:"required" json:"-"`
	BatchInputPublicAssociationMultiArchive BatchInputPublicAssociationMultiArchiveParam
	paramObj
}

func (r AssociationBatchDeleteParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.BatchInputPublicAssociationMultiArchive)
}
func (r *AssociationBatchDeleteParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AssociationBatchNewDefaultParams struct {
	FromObjectType                              string `path:"fromObjectType" api:"required" json:"-"`
	BatchInputPublicDefaultAssociationMultiPost BatchInputPublicDefaultAssociationMultiPostParam
	paramObj
}

func (r AssociationBatchNewDefaultParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.BatchInputPublicDefaultAssociationMultiPost)
}
func (r *AssociationBatchNewDefaultParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AssociationBatchDeleteLabelsParams struct {
	FromObjectType                       string `path:"fromObjectType" api:"required" json:"-"`
	BatchInputPublicAssociationMultiPost BatchInputPublicAssociationMultiPostParam
	paramObj
}

func (r AssociationBatchDeleteLabelsParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.BatchInputPublicAssociationMultiPost)
}
func (r *AssociationBatchDeleteLabelsParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AssociationBatchGetParams struct {
	FromObjectType                                string `path:"fromObjectType" api:"required" json:"-"`
	BatchInputPublicFetchAssociationsBatchRequest BatchInputPublicFetchAssociationsBatchRequestParam
	paramObj
}

func (r AssociationBatchGetParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.BatchInputPublicFetchAssociationsBatchRequest)
}
func (r *AssociationBatchGetParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
