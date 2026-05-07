// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package marketing

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/HubSpot/hubspot-sdk-go/internal/apiquery"
	"github.com/HubSpot/hubspot-sdk-go/internal/requestconfig"
	"github.com/HubSpot/hubspot-sdk-go/option"
	"github.com/HubSpot/hubspot-sdk-go/packages/param"
)

// CampaignAssetService contains methods and other services that help with
// interacting with the hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCampaignAssetService] method instead.
type CampaignAssetService struct {
	options []option.RequestOption
}

// NewCampaignAssetService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCampaignAssetService(opts ...option.RequestOption) (r CampaignAssetService) {
	r = CampaignAssetService{}
	r.options = opts
	return
}

// Associate a specified asset with a campaign. Using the API, you can create
// associations for the following asset types: ads, blog posts, calls, case
// studies, CTAs, CTAs (legacy), external website pages, feedback surveys, forms,
// files, knowledge base articles, landing pages, marketing email, marketing
// events, meetings, playbooks, podcast episodes, sales documents, sales emails,
// sequences, SMS, social posts, static lists, videos, website pages, and
// workflows.
//
// For other asset types, it is recommended to manage your associations directly in
// the campaign tool in HubSpot.
func (r *CampaignAssetService) Update(ctx context.Context, assetID string, body CampaignAssetUpdateParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.options, opts)
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
	path := fmt.Sprintf("marketing/campaigns/2026-03/%s/assets/%s/%s", url.PathEscape(body.CampaignGuid), url.PathEscape(body.AssetType), url.PathEscape(assetID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, nil, nil, opts...)
	return err
}

// This endpoint lists all assets of the campaign by asset type. The assetType
// parameter is required, and each request can only fetch assets of a single type.
// Asset metrics can also be fetched along with the assets; they are available only
// if start and end dates are provided.
func (r *CampaignAssetService) List(ctx context.Context, assetType string, params CampaignAssetListParams, opts ...option.RequestOption) (res *CollectionResponsePublicCampaignAssetForwardPaging, err error) {
	opts = slices.Concat(r.options, opts)
	if params.CampaignGuid == "" {
		err = errors.New("missing required campaignGuid parameter")
		return nil, err
	}
	if assetType == "" {
		err = errors.New("missing required assetType parameter")
		return nil, err
	}
	path := fmt.Sprintf("marketing/campaigns/2026-03/%s/assets/%s", url.PathEscape(params.CampaignGuid), url.PathEscape(assetType))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// Disassociate a specified asset from a campaign. Using the API, you can remove
// associations for the following asset types: ads, blog posts, calls, case
// studies, CTAs, CTAs (legacy), external website pages, feedback surveys, forms,
// files, knowledge base articles, landing pages, marketing email, marketing
// events, meetings, playbooks, podcast episodes, sales documents, sales emails,
// sequences, SMS, social posts, static lists, videos, website pages, and
// workflows.
//
// For other asset types, it is recommended to manage your associations directly in
// the campaign tool in HubSpot.
func (r *CampaignAssetService) Delete(ctx context.Context, assetID string, body CampaignAssetDeleteParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.options, opts)
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
	path := fmt.Sprintf("marketing/campaigns/2026-03/%s/assets/%s/%s", url.PathEscape(body.CampaignGuid), url.PathEscape(body.AssetType), url.PathEscape(assetID))
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
	// End date to fetch asset metrics, formatted as YYYY-MM-DD. This date is used to
	// fetch the metrics associated with the assets for a specified period. If not
	// provided, no asset metrics will be fetched. Example: 2024-01-27
	EndDate param.Opt[string] `query:"endDate,omitzero" json:"-"`
	// The maximum number of results to display per page.
	Limit param.Opt[string] `query:"limit,omitzero" json:"-"`
	// Start date to fetch asset metrics, formatted as YYYY-MM-DD. This date is used to
	// fetch the metrics associated with the assets for a specified period. If not
	// provided, no asset metrics will be fetched. Example: 2023-01-20
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
