// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package marketing

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

// CampaignBudgetService contains methods and other services that help with
// interacting with the hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCampaignBudgetService] method instead.
type CampaignBudgetService struct {
	Options []option.RequestOption
}

// NewCampaignBudgetService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCampaignBudgetService(opts ...option.RequestOption) (r CampaignBudgetService) {
	r = CampaignBudgetService{}
	r.Options = opts
	return
}

// Add a new budget item to the campaign
func (r *CampaignBudgetService) New(ctx context.Context, campaignGuid string, body CampaignBudgetNewParams, opts ...option.RequestOption) (res *PublicBudgetItem, err error) {
	opts = slices.Concat(r.Options, opts)
	if campaignGuid == "" {
		err = errors.New("missing required campaignGuid parameter")
		return
	}
	path := fmt.Sprintf("marketing/v3/campaigns/%s/budget", campaignGuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Update a specific budget item by ID
func (r *CampaignBudgetService) Update(ctx context.Context, budgetID int64, params CampaignBudgetUpdateParams, opts ...option.RequestOption) (res *PublicBudgetItem, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.CampaignGuid == "" {
		err = errors.New("missing required campaignGuid parameter")
		return
	}
	path := fmt.Sprintf("marketing/v3/campaigns/%s/budget/%v", params.CampaignGuid, budgetID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return
}

// Delete a specific budget item by ID
func (r *CampaignBudgetService) Delete(ctx context.Context, budgetID int64, body CampaignBudgetDeleteParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if body.CampaignGuid == "" {
		err = errors.New("missing required campaignGuid parameter")
		return
	}
	path := fmt.Sprintf("marketing/v3/campaigns/%s/budget/%v", body.CampaignGuid, budgetID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Get a specific budget item by ID
func (r *CampaignBudgetService) Get(ctx context.Context, budgetID int64, query CampaignBudgetGetParams, opts ...option.RequestOption) (res *PublicBudgetItem, err error) {
	opts = slices.Concat(r.Options, opts)
	if query.CampaignGuid == "" {
		err = errors.New("missing required campaignGuid parameter")
		return
	}
	path := fmt.Sprintf("marketing/v3/campaigns/%s/budget/%v", query.CampaignGuid, budgetID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Retrieve detailed information about the budget and spend items for a specified
// campaign, including the total budget, total spend, and remaining budget. Budget
// and Spend items may be returned in any order, but the order field specifies
// their sequence based on the creation date. The item with order 0 is the oldest,
// and items with higher order values are newer
func (r *CampaignBudgetService) GetTotals(ctx context.Context, campaignGuid string, opts ...option.RequestOption) (res *PublicBudgetTotals, err error) {
	opts = slices.Concat(r.Options, opts)
	if campaignGuid == "" {
		err = errors.New("missing required campaignGuid parameter")
		return
	}
	path := fmt.Sprintf("marketing/v3/campaigns/%s/budget/totals", campaignGuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type CampaignBudgetNewParams struct {
	PublicBudgetItemInput PublicBudgetItemInputParam
	paramObj
}

func (r CampaignBudgetNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.PublicBudgetItemInput)
}
func (r *CampaignBudgetNewParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.PublicBudgetItemInput)
}

type CampaignBudgetUpdateParams struct {
	CampaignGuid          string `path:"campaignGuid,required" json:"-"`
	PublicBudgetItemInput PublicBudgetItemInputParam
	paramObj
}

func (r CampaignBudgetUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.PublicBudgetItemInput)
}
func (r *CampaignBudgetUpdateParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.PublicBudgetItemInput)
}

type CampaignBudgetDeleteParams struct {
	CampaignGuid string `path:"campaignGuid,required" json:"-"`
	paramObj
}

type CampaignBudgetGetParams struct {
	CampaignGuid string `path:"campaignGuid,required" json:"-"`
	paramObj
}
