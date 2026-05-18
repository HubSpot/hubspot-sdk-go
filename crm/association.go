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
	"github.com/HubSpot/hubspot-sdk-go/internal/apiquery"
	shimjson "github.com/HubSpot/hubspot-sdk-go/internal/encoding/json"
	"github.com/HubSpot/hubspot-sdk-go/internal/requestconfig"
	"github.com/HubSpot/hubspot-sdk-go/option"
	"github.com/HubSpot/hubspot-sdk-go/packages/pagination"
	"github.com/HubSpot/hubspot-sdk-go/packages/param"
	"github.com/HubSpot/hubspot-sdk-go/shared"
)

// AssociationService contains methods and other services that help with
// interacting with the hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAssociationService] method instead.
type AssociationService struct {
	options []option.RequestOption
	Batch   AssociationBatchService
}

// NewAssociationService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAssociationService(opts ...option.RequestOption) (r AssociationService) {
	r = AssociationService{}
	r.options = opts
	r.Batch = NewAssociationBatchService(opts...)
	return
}

// Create the default (most generic) association type between two object types
func (r *AssociationService) New(ctx context.Context, toObjectID string, body AssociationNewParams, opts ...option.RequestOption) (res *BatchResponsePublicDefaultAssociation, err error) {
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

// Retrieve all associations between a specific record and an object type. Limit
// 500 per call.
func (r *AssociationService) List(ctx context.Context, toObjectType string, params AssociationListParams, opts ...option.RequestOption) (res *pagination.Page[MultiAssociatedObjectWithLabel], err error) {
	var raw *http.Response
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.ObjectType == "" {
		err = errors.New("missing required objectType parameter")
		return nil, err
	}
	if params.ObjectID == "" {
		err = errors.New("missing required objectId parameter")
		return nil, err
	}
	if toObjectType == "" {
		err = errors.New("missing required toObjectType parameter")
		return nil, err
	}
	path := fmt.Sprintf("crm/objects/2026-03/%s/%s/associations/%s", url.PathEscape(params.ObjectType), url.PathEscape(params.ObjectID), url.PathEscape(toObjectType))
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, params, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// Retrieve all associations between a specific record and an object type. Limit
// 500 per call.
func (r *AssociationService) ListAutoPaging(ctx context.Context, toObjectType string, params AssociationListParams, opts ...option.RequestOption) *pagination.PageAutoPager[MultiAssociatedObjectWithLabel] {
	return pagination.NewPageAutoPager(r.List(ctx, toObjectType, params, opts...))
}

// deletes all associations between two records.
func (r *AssociationService) Delete(ctx context.Context, toObjectID string, body AssociationDeleteParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if body.ObjectType == "" {
		err = errors.New("missing required objectType parameter")
		return err
	}
	if body.ObjectID == "" {
		err = errors.New("missing required objectId parameter")
		return err
	}
	if body.ToObjectType == "" {
		err = errors.New("missing required toObjectType parameter")
		return err
	}
	if toObjectID == "" {
		err = errors.New("missing required toObjectId parameter")
		return err
	}
	path := fmt.Sprintf("crm/objects/2026-03/%s/%s/associations/%s/%s", url.PathEscape(body.ObjectType), url.PathEscape(body.ObjectID), url.PathEscape(body.ToObjectType), url.PathEscape(toObjectID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Requests a report of all objects in the portal which have a high usage of
// associations
func (r *AssociationService) RequestHighUsageReport(ctx context.Context, userID int64, opts ...option.RequestOption) (res *ReportCreationResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := fmt.Sprintf("crm/associations/2026-03/usage/high-usage-report/%v", userID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

func (r *AssociationService) Search(ctx context.Context, objectType string, body AssociationSearchParams, opts ...option.RequestOption) (res *CollectionResponseWithTotalSimplePublicObject, err error) {
	opts = slices.Concat(r.options, opts)
	if objectType == "" {
		err = errors.New("missing required objectType parameter")
		return nil, err
	}
	path := fmt.Sprintf("crm/objects/2026-03/%s/search", url.PathEscape(objectType))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Set association labels between two records.
func (r *AssociationService) UpdateLabels(ctx context.Context, toObjectID string, params AssociationUpdateLabelsParams, opts ...option.RequestOption) (res *LabelsBetweenObjectPair, err error) {
	opts = slices.Concat(r.options, opts)
	if params.ObjectType == "" {
		err = errors.New("missing required objectType parameter")
		return nil, err
	}
	if params.ObjectID == "" {
		err = errors.New("missing required objectId parameter")
		return nil, err
	}
	if params.ToObjectType == "" {
		err = errors.New("missing required toObjectType parameter")
		return nil, err
	}
	if toObjectID == "" {
		err = errors.New("missing required toObjectId parameter")
		return nil, err
	}
	path := fmt.Sprintf("crm/objects/2026-03/%s/%s/associations/%s/%s", url.PathEscape(params.ObjectType), url.PathEscape(params.ObjectID), url.PathEscape(params.ToObjectType), url.PathEscape(toObjectID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return res, err
}

type AssociationNewParams struct {
	FromObjectType string `path:"fromObjectType" api:"required" json:"-"`
	FromObjectID   string `path:"fromObjectId" api:"required" json:"-"`
	ToObjectType   string `path:"toObjectType" api:"required" json:"-"`
	paramObj
}

type AssociationListParams struct {
	ObjectType string `path:"objectType" api:"required" json:"-"`
	ObjectID   string `path:"objectId" api:"required" json:"-"`
	// The paging cursor token of the last successfully read resource will be returned
	// as the `paging.next.after` JSON property of a paged response containing more
	// results.
	After param.Opt[string] `query:"after,omitzero" json:"-"`
	// The maximum number of results to display per page.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [AssociationListParams]'s query parameters as `url.Values`.
func (r AssociationListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AssociationDeleteParams struct {
	ObjectType   string `path:"objectType" api:"required" json:"-"`
	ObjectID     string `path:"objectId" api:"required" json:"-"`
	ToObjectType string `path:"toObjectType" api:"required" json:"-"`
	paramObj
}

type AssociationSearchParams struct {
	// Describes a search request
	PublicObjectSearchRequest PublicObjectSearchRequestParam
	paramObj
}

func (r AssociationSearchParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.PublicObjectSearchRequest)
}
func (r *AssociationSearchParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AssociationUpdateLabelsParams struct {
	ObjectType   string `path:"objectType" api:"required" json:"-"`
	ObjectID     string `path:"objectId" api:"required" json:"-"`
	ToObjectType string `path:"toObjectType" api:"required" json:"-"`
	Body         []shared.AssociationSpecParam
	paramObj
}

func (r AssociationUpdateLabelsParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.Body)
}
func (r *AssociationUpdateLabelsParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
