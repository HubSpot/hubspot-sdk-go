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

// CampaignBudgetService contains methods and other services that help with
// interacting with the hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCampaignBudgetService] method instead.
type CampaignBudgetService struct {
	options []option.RequestOption
}

// NewCampaignBudgetService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCampaignBudgetService(opts ...option.RequestOption) (r CampaignBudgetService) {
	r = CampaignBudgetService{}
	r.options = opts
	return
}

// Add a new budget item to the campaign
func (r *CampaignBudgetService) New(ctx context.Context, campaignGuid string, body CampaignBudgetNewParams, opts ...option.RequestOption) (res *PublicBudgetItem, err error) {
	opts = slices.Concat(r.options, opts)
	if campaignGuid == "" {
		err = errors.New("missing required campaignGuid parameter")
		return nil, err
	}
	path := fmt.Sprintf("marketing/campaigns/2026-03/%s/budget", url.PathEscape(campaignGuid))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Update a specific budget item by ID
func (r *CampaignBudgetService) Update(ctx context.Context, budgetID int64, params CampaignBudgetUpdateParams, opts ...option.RequestOption) (res *PublicBudgetItem, err error) {
	opts = slices.Concat(r.options, opts)
	if params.CampaignGuid == "" {
		err = errors.New("missing required campaignGuid parameter")
		return nil, err
	}
	path := fmt.Sprintf("marketing/campaigns/2026-03/%s/budget/%v", url.PathEscape(params.CampaignGuid), budgetID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return res, err
}

// Delete a specific budget item by ID
func (r *CampaignBudgetService) Delete(ctx context.Context, budgetID int64, body CampaignBudgetDeleteParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if body.CampaignGuid == "" {
		err = errors.New("missing required campaignGuid parameter")
		return err
	}
	path := fmt.Sprintf("marketing/campaigns/2026-03/%s/budget/%v", url.PathEscape(body.CampaignGuid), budgetID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Get a specific budget item by ID
func (r *CampaignBudgetService) Get(ctx context.Context, budgetID int64, query CampaignBudgetGetParams, opts ...option.RequestOption) (res *PublicBudgetItem, err error) {
	opts = slices.Concat(r.options, opts)
	if query.CampaignGuid == "" {
		err = errors.New("missing required campaignGuid parameter")
		return nil, err
	}
	path := fmt.Sprintf("marketing/campaigns/2026-03/%s/budget/%v", url.PathEscape(query.CampaignGuid), budgetID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Retrieve detailed information about the budget and spend items for a specified
// campaign, including the total budget, total spend, and remaining budget. Budget
// and Spend items may be returned in any order, but the order field specifies
// their sequence based on the creation date. The item with order 0 is the oldest,
// and items with higher order values are newer
func (r *CampaignBudgetService) GetTotals(ctx context.Context, campaignGuid string, opts ...option.RequestOption) (res *PublicBudgetTotals, err error) {
	opts = slices.Concat(r.options, opts)
	if campaignGuid == "" {
		err = errors.New("missing required campaignGuid parameter")
		return nil, err
	}
	path := fmt.Sprintf("marketing/campaigns/2026-03/%s/budget/totals", url.PathEscape(campaignGuid))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type CampaignBudgetNewParams struct {
	PublicBudgetItemInput PublicBudgetItemInputParam
	paramObj
}

func (r CampaignBudgetNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.PublicBudgetItemInput)
}
func (r *CampaignBudgetNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CampaignBudgetUpdateParams struct {
	CampaignGuid          string `path:"campaignGuid" api:"required" json:"-"`
	PublicBudgetItemInput PublicBudgetItemInputParam
	paramObj
}

func (r CampaignBudgetUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.PublicBudgetItemInput)
}
func (r *CampaignBudgetUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CampaignBudgetDeleteParams struct {
	CampaignGuid string `path:"campaignGuid" api:"required" json:"-"`
	paramObj
}

type CampaignBudgetGetParams struct {
	CampaignGuid string `path:"campaignGuid" api:"required" json:"-"`
	paramObj
}
