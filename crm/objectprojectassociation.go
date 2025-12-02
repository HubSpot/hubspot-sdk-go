// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package crm

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/stainless-sdks/hubspot-sdk-go/internal/apiquery"
	"github.com/stainless-sdks/hubspot-sdk-go/internal/requestconfig"
	"github.com/stainless-sdks/hubspot-sdk-go/option"
	"github.com/stainless-sdks/hubspot-sdk-go/packages/pagination"
	"github.com/stainless-sdks/hubspot-sdk-go/packages/param"
)

// ObjectProjectAssociationService contains methods and other services that help
// with interacting with the hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewObjectProjectAssociationService] method instead.
type ObjectProjectAssociationService struct {
	Options []option.RequestOption
}

// NewObjectProjectAssociationService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewObjectProjectAssociationService(opts ...option.RequestOption) (r ObjectProjectAssociationService) {
	r = ObjectProjectAssociationService{}
	r.Options = opts
	return
}

func (r *ObjectProjectAssociationService) Update(ctx context.Context, associationType string, body ObjectProjectAssociationUpdateParams, opts ...option.RequestOption) (res *SimplePublicObjectWithAssociations, err error) {
	opts = slices.Concat(r.Options, opts)
	if body.ProjectID == "" {
		err = errors.New("missing required projectId parameter")
		return
	}
	if body.ToObjectType == "" {
		err = errors.New("missing required toObjectType parameter")
		return
	}
	if body.ToObjectID == "" {
		err = errors.New("missing required toObjectId parameter")
		return
	}
	if associationType == "" {
		err = errors.New("missing required associationType parameter")
		return
	}
	path := fmt.Sprintf("crm/objects/v3/projects/%s/associations/%s/%s/%s", body.ProjectID, body.ToObjectType, body.ToObjectID, associationType)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, nil, &res, opts...)
	return
}

func (r *ObjectProjectAssociationService) List(ctx context.Context, toObjectType string, params ObjectProjectAssociationListParams, opts ...option.RequestOption) (res *pagination.Page[AssociatedID], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.ProjectID == "" {
		err = errors.New("missing required projectId parameter")
		return
	}
	if toObjectType == "" {
		err = errors.New("missing required toObjectType parameter")
		return
	}
	path := fmt.Sprintf("crm/objects/v3/projects/%s/associations/%s", params.ProjectID, toObjectType)
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

func (r *ObjectProjectAssociationService) ListAutoPaging(ctx context.Context, toObjectType string, params ObjectProjectAssociationListParams, opts ...option.RequestOption) *pagination.PageAutoPager[AssociatedID] {
	return pagination.NewPageAutoPager(r.List(ctx, toObjectType, params, opts...))
}

func (r *ObjectProjectAssociationService) Delete(ctx context.Context, associationType string, body ObjectProjectAssociationDeleteParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if body.ProjectID == "" {
		err = errors.New("missing required projectId parameter")
		return
	}
	if body.ToObjectType == "" {
		err = errors.New("missing required toObjectType parameter")
		return
	}
	if body.ToObjectID == "" {
		err = errors.New("missing required toObjectId parameter")
		return
	}
	if associationType == "" {
		err = errors.New("missing required associationType parameter")
		return
	}
	path := fmt.Sprintf("crm/objects/v3/projects/%s/associations/%s/%s/%s", body.ProjectID, body.ToObjectType, body.ToObjectID, associationType)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

type ObjectProjectAssociationUpdateParams struct {
	ProjectID    string `path:"projectId,required" json:"-"`
	ToObjectType string `path:"toObjectType,required" json:"-"`
	ToObjectID   string `path:"toObjectId,required" json:"-"`
	paramObj
}

type ObjectProjectAssociationListParams struct {
	ProjectID string            `path:"projectId,required" json:"-"`
	After     param.Opt[string] `query:"after,omitzero" json:"-"`
	IncludeFa param.Opt[bool]   `query:"includeFA,omitzero" json:"-"`
	Limit     param.Opt[int64]  `query:"limit,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ObjectProjectAssociationListParams]'s query parameters as
// `url.Values`.
func (r ObjectProjectAssociationListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ObjectProjectAssociationDeleteParams struct {
	ProjectID    string `path:"projectId,required" json:"-"`
	ToObjectType string `path:"toObjectType,required" json:"-"`
	ToObjectID   string `path:"toObjectId,required" json:"-"`
	paramObj
}
