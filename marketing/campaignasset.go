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

// Associate a specified asset with a campaign. Using the API, you can create and
// remove associations for the following asset types: forms, static lists, external
// website pages, sequences, meetings, playbooks, feedback surveys, podcast
// episodes, sales documents, marketing emails, case studies, knowledge base
// articles, calls, and CTAs.
//
// For other asset types, it is recommended to manage your associations directly in
// the campaign tool in HubSpot.
func (r *CampaignAssetService) Update(ctx context.Context, assetID string, body CampaignAssetUpdateParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if body.CampaignGuid == "" {
		err = errors.New("missing required campaignGuid parameter")
		return
	}
	if body.AssetType == "" {
		err = errors.New("missing required assetType parameter")
		return
	}
	if assetID == "" {
		err = errors.New("missing required assetId parameter")
		return
	}
	path := fmt.Sprintf("marketing/v3/campaigns/%s/assets/%s/%s", body.CampaignGuid, body.AssetType, assetID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, nil, nil, opts...)
	return
}

// This endpoint lists all assets of the campaign by asset type. The assetType
// parameter is required, and each request can only fetch assets of a single type.
// Asset metrics can also be fetched along with the assets; they are available only
// if start and end dates are provided.
func (r *CampaignAssetService) List(ctx context.Context, assetType string, params CampaignAssetListParams, opts ...option.RequestOption) (res *CollectionResponsePublicCampaignAssetForwardPaging, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.CampaignGuid == "" {
		err = errors.New("missing required campaignGuid parameter")
		return
	}
	if assetType == "" {
		err = errors.New("missing required assetType parameter")
		return
	}
	path := fmt.Sprintf("marketing/v3/campaigns/%s/assets/%s", params.CampaignGuid, assetType)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return
}

// Disassociate a specified asset from a campaign. Important: Currently, only the
// following asset types can be associated and disassociated via the API: Forms,
// Static lists, External website pages
func (r *CampaignAssetService) Delete(ctx context.Context, assetID string, body CampaignAssetDeleteParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if body.CampaignGuid == "" {
		err = errors.New("missing required campaignGuid parameter")
		return
	}
	if body.AssetType == "" {
		err = errors.New("missing required assetType parameter")
		return
	}
	if assetID == "" {
		err = errors.New("missing required assetId parameter")
		return
	}
	path := fmt.Sprintf("marketing/v3/campaigns/%s/assets/%s/%s", body.CampaignGuid, body.AssetType, assetID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

type CampaignAssetUpdateParams struct {
	CampaignGuid string `path:"campaignGuid,required" json:"-"`
	AssetType    string `path:"assetType,required" json:"-"`
	paramObj
}

type CampaignAssetListParams struct {
	CampaignGuid string `path:"campaignGuid,required" json:"-"`
	// A cursor for pagination. If provided, the results will start after the given
	// cursor. Example: NTI1Cg%3D%3D
	After param.Opt[string] `query:"after,omitzero" json:"-"`
	// End date to fetch asset metrics, formatted as YYYY-MM-DD. This date is used to
	// fetch the metrics associated with the assets for a specified period. If not
	// provided, no asset metrics will be fetched.
	EndDate param.Opt[string] `query:"endDate,omitzero" json:"-"`
	// The maximum number of results to return. Default: 10
	Limit param.Opt[string] `query:"limit,omitzero" json:"-"`
	// Start date to fetch asset metrics, formatted as YYYY-MM-DD. This date is used to
	// fetch the metrics associated with the assets for a specified period. If not
	// provided, no asset metrics will be fetched.
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
	CampaignGuid string `path:"campaignGuid,required" json:"-"`
	AssetType    string `path:"assetType,required" json:"-"`
	paramObj
}
