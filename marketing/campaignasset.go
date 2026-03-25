// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package marketing

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
	"github.com/stainless-sdks/hubspot-sdk-go/packages/param"
)

// CampaignAssetService contains methods and other services that help with
// interacting with the hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCampaignAssetService] method instead.
type CampaignAssetService struct {
	Options []option.RequestOption
}

// NewCampaignAssetService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCampaignAssetService(opts ...option.RequestOption) (r CampaignAssetService) {
	r = CampaignAssetService{}
	r.Options = opts
	return
}

// Associate an asset with a specific campaign in your HubSpot account. This
// operation allows you to link an asset of a specified type and ID to a campaign,
// facilitating better organization and tracking of campaign resources.
func (r *CampaignAssetService) Update(ctx context.Context, assetID string, body CampaignAssetUpdateParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if body.CampaignGuid == "" {
		err = errors.New("missing required campaignGuid parameter")
		return err
	}
	if body.AssetType == "" {
		err = errors.New("missing required assetType parameter")
		return err
	}
	if assetID == "" {
		err = errors.New("missing required assetId parameter")
		return err
	}
	path := fmt.Sprintf("marketing/campaigns/2026-03/%s/assets/%s/%s", body.CampaignGuid, body.AssetType, assetID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, nil, nil, opts...)
	return err
}

// List all assets of a specified campaign by asset type. This endpoint allows you
// to retrieve assets associated with a campaign, filtered by the type of asset. It
// supports pagination and date filtering to manage and refine the results.
func (r *CampaignAssetService) List(ctx context.Context, assetType string, params CampaignAssetListParams, opts ...option.RequestOption) (res *CollectionResponsePublicCampaignAssetForwardPaging, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.CampaignGuid == "" {
		err = errors.New("missing required campaignGuid parameter")
		return nil, err
	}
	if assetType == "" {
		err = errors.New("missing required assetType parameter")
		return nil, err
	}
	path := fmt.Sprintf("marketing/campaigns/2026-03/%s/assets/%s", params.CampaignGuid, assetType)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// Disassociate an asset from a specific campaign. This operation removes the
// association between the specified asset and campaign, effectively detaching the
// asset from the campaign's context.
func (r *CampaignAssetService) Delete(ctx context.Context, assetID string, body CampaignAssetDeleteParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if body.CampaignGuid == "" {
		err = errors.New("missing required campaignGuid parameter")
		return err
	}
	if body.AssetType == "" {
		err = errors.New("missing required assetType parameter")
		return err
	}
	if assetID == "" {
		err = errors.New("missing required assetId parameter")
		return err
	}
	path := fmt.Sprintf("marketing/campaigns/2026-03/%s/assets/%s/%s", body.CampaignGuid, body.AssetType, assetID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

type CampaignAssetUpdateParams struct {
	CampaignGuid string `path:"campaignGuid" api:"required" json:"-"`
	AssetType    string `path:"assetType" api:"required" json:"-"`
	paramObj
}

type CampaignAssetListParams struct {
	CampaignGuid string `path:"campaignGuid" api:"required" json:"-"`
	// The paging cursor token of the last successfully read resource will be returned
	// as the `paging.next.after` JSON property of a paged response containing more
	// results.
	After param.Opt[string] `query:"after,omitzero" json:"-"`
	// The end date for filtering assets, in YYYY-MM-DD format.
	EndDate param.Opt[string] `query:"endDate,omitzero" json:"-"`
	// The maximum number of results to display per page.
	Limit param.Opt[string] `query:"limit,omitzero" json:"-"`
	// The start date for filtering assets, in YYYY-MM-DD format.
	StartDate param.Opt[string] `query:"startDate,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [CampaignAssetListParams]'s query parameters as
// `url.Values`.
func (r CampaignAssetListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type CampaignAssetDeleteParams struct {
	CampaignGuid string `path:"campaignGuid" api:"required" json:"-"`
	AssetType    string `path:"assetType" api:"required" json:"-"`
	paramObj
}
