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

// ObjectPartnerServiceAssociationService contains methods and other services that
// help with interacting with the Hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewObjectPartnerServiceAssociationService] method instead.
type ObjectPartnerServiceAssociationService struct {
	Options []option.RequestOption
}

// NewObjectPartnerServiceAssociationService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewObjectPartnerServiceAssociationService(opts ...option.RequestOption) (r ObjectPartnerServiceAssociationService) {
	r = ObjectPartnerServiceAssociationService{}
	r.Options = opts
	return
}

// Associate a partner service with another object
func (r *ObjectPartnerServiceAssociationService) Update(ctx context.Context, associationType string, body ObjectPartnerServiceAssociationUpdateParams, opts ...option.RequestOption) (res *SimplePublicObjectWithAssociations, err error) {
	opts = slices.Concat(r.Options, opts)
	if body.PartnerServiceID == "" {
		err = errors.New("missing required partnerServiceId parameter")
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
	path := fmt.Sprintf("crm/v3/objects/partner_services/%s/associations/%s/%s/%s", body.PartnerServiceID, body.ToObjectType, body.ToObjectID, associationType)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, nil, &res, opts...)
	return
}

// List associations of a partner service by type
func (r *ObjectPartnerServiceAssociationService) List(ctx context.Context, toObjectType string, params ObjectPartnerServiceAssociationListParams, opts ...option.RequestOption) (res *pagination.Page[AssociatedID], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.PartnerServiceID == "" {
		err = errors.New("missing required partnerServiceId parameter")
		return
	}
	if toObjectType == "" {
		err = errors.New("missing required toObjectType parameter")
		return
	}
	path := fmt.Sprintf("crm/v3/objects/partner_services/%s/associations/%s", params.PartnerServiceID, toObjectType)
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

// List associations of a partner service by type
func (r *ObjectPartnerServiceAssociationService) ListAutoPaging(ctx context.Context, toObjectType string, params ObjectPartnerServiceAssociationListParams, opts ...option.RequestOption) *pagination.PageAutoPager[AssociatedID] {
	return pagination.NewPageAutoPager(r.List(ctx, toObjectType, params, opts...))
}

// Remove an association between two partner services
func (r *ObjectPartnerServiceAssociationService) Delete(ctx context.Context, associationType string, body ObjectPartnerServiceAssociationDeleteParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if body.PartnerServiceID == "" {
		err = errors.New("missing required partnerServiceId parameter")
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
	path := fmt.Sprintf("crm/v3/objects/partner_services/%s/associations/%s/%s/%s", body.PartnerServiceID, body.ToObjectType, body.ToObjectID, associationType)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

type ObjectPartnerServiceAssociationUpdateParams struct {
	PartnerServiceID string `path:"partnerServiceId,required" json:"-"`
	ToObjectType     string `path:"toObjectType,required" json:"-"`
	ToObjectID       string `path:"toObjectId,required" json:"-"`
	paramObj
}

type ObjectPartnerServiceAssociationListParams struct {
	PartnerServiceID string `path:"partnerServiceId,required" json:"-"`
	// The paging cursor token of the last successfully read resource will be returned
	// as the `paging.next.after` JSON property of a paged response containing more
	// results.
	After     param.Opt[string] `query:"after,omitzero" json:"-"`
	IncludeFa param.Opt[bool]   `query:"includeFA,omitzero" json:"-"`
	// The maximum number of results to display per page.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ObjectPartnerServiceAssociationListParams]'s query
// parameters as `url.Values`.
func (r ObjectPartnerServiceAssociationListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ObjectPartnerServiceAssociationDeleteParams struct {
	PartnerServiceID string `path:"partnerServiceId,required" json:"-"`
	ToObjectType     string `path:"toObjectType,required" json:"-"`
	ToObjectID       string `path:"toObjectId,required" json:"-"`
	paramObj
}
