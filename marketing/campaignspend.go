// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package marketing

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/HubSpot/hubspot-sdk-go/internal/apijson"
	shimjson "github.com/HubSpot/hubspot-sdk-go/internal/encoding/json"
	"github.com/HubSpot/hubspot-sdk-go/internal/requestconfig"
	"github.com/HubSpot/hubspot-sdk-go/option"
)

// CampaignSpendService contains methods and other services that help with
// interacting with the hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCampaignSpendService] method instead.
type CampaignSpendService struct {
	options []option.RequestOption
}

// NewCampaignSpendService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCampaignSpendService(opts ...option.RequestOption) (r CampaignSpendService) {
	r = CampaignSpendService{}
	r.options = opts
	return
}

// Create a new campaign spend item
func (r *CampaignSpendService) New(ctx context.Context, campaignGuid string, body CampaignSpendNewParams, opts ...option.RequestOption) (res *PublicSpendItem, err error) {
	opts = slices.Concat(r.options, opts)
	if campaignGuid == "" {
		err = errors.New("missing required campaignGuid parameter")
		return nil, err
	}
	path := fmt.Sprintf("marketing/campaigns/2026-03/%s/spend", url.PathEscape(campaignGuid))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Update a specific campaign spend item by ID
func (r *CampaignSpendService) Update(ctx context.Context, spendID int64, params CampaignSpendUpdateParams, opts ...option.RequestOption) (res *PublicSpendItem, err error) {
	opts = slices.Concat(r.options, opts)
	if params.CampaignGuid == "" {
		err = errors.New("missing required campaignGuid parameter")
		return nil, err
	}
	path := fmt.Sprintf("marketing/campaigns/2026-03/%s/spend/%v", url.PathEscape(params.CampaignGuid), spendID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return res, err
}

// Delete a specific campaign spend item by ID
func (r *CampaignSpendService) Delete(ctx context.Context, spendID int64, body CampaignSpendDeleteParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if body.CampaignGuid == "" {
		err = errors.New("missing required campaignGuid parameter")
		return err
	}
	path := fmt.Sprintf("marketing/campaigns/2026-03/%s/spend/%v", url.PathEscape(body.CampaignGuid), spendID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Read a campaign spend item by its spendId
func (r *CampaignSpendService) Get(ctx context.Context, spendID int64, query CampaignSpendGetParams, opts ...option.RequestOption) (res *PublicSpendItem, err error) {
	opts = slices.Concat(r.options, opts)
	if query.CampaignGuid == "" {
		err = errors.New("missing required campaignGuid parameter")
		return nil, err
	}
	path := fmt.Sprintf("marketing/campaigns/2026-03/%s/spend/%v", url.PathEscape(query.CampaignGuid), spendID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type CampaignSpendNewParams struct {
	PublicSpendItemInput PublicSpendItemInputParam
	paramObj
}

func (r CampaignSpendNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.PublicSpendItemInput)
}
func (r *CampaignSpendNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CampaignSpendUpdateParams struct {
	CampaignGuid         string `path:"campaignGuid" api:"required" json:"-"`
	PublicSpendItemInput PublicSpendItemInputParam
	paramObj
}

func (r CampaignSpendUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.PublicSpendItemInput)
}
func (r *CampaignSpendUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CampaignSpendDeleteParams struct {
	CampaignGuid string `path:"campaignGuid" api:"required" json:"-"`
	paramObj
}

type CampaignSpendGetParams struct {
	CampaignGuid string `path:"campaignGuid" api:"required" json:"-"`
	paramObj
}
