// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package marketing

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/stainless-sdks/hubspot-sdk-go/internal/apijson"
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

// Add a new budget item to the specified campaign. This operation allows you to
// allocate a budget for a campaign by specifying the necessary details in the
// request body.
func (r *CampaignBudgetService) New(ctx context.Context, campaignGuid string, body CampaignBudgetNewParams, opts ...option.RequestOption) (res *PublicBudgetItem, err error) {
	opts = slices.Concat(r.Options, opts)
	if campaignGuid == "" {
		err = errors.New("missing required campaignGuid parameter")
		return nil, err
	}
	path := fmt.Sprintf("marketing/campaigns/2026-03/%s/budget", campaignGuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Update a specific budget item by its ID within a marketing campaign. This
// operation allows you to modify the details of a budget item, such as its amount,
// name, or order, ensuring that your campaign's financial records are accurate and
// up-to-date.
func (r *CampaignBudgetService) Update(ctx context.Context, budgetID int64, params CampaignBudgetUpdateParams, opts ...option.RequestOption) (res *PublicBudgetItem, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.CampaignGuid == "" {
		err = errors.New("missing required campaignGuid parameter")
		return nil, err
	}
	path := fmt.Sprintf("marketing/campaigns/2026-03/%s/budget/%v", params.CampaignGuid, budgetID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return res, err
}

// Delete a specific budget item from a campaign using its unique ID. This
// operation removes the budget item from the campaign's budget list, ensuring it
// is no longer considered in budget calculations.
func (r *CampaignBudgetService) Delete(ctx context.Context, budgetID int64, body CampaignBudgetDeleteParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if body.CampaignGuid == "" {
		err = errors.New("missing required campaignGuid parameter")
		return err
	}
	path := fmt.Sprintf("marketing/campaigns/2026-03/%s/budget/%v", body.CampaignGuid, budgetID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Retrieve a specific budget item by its ID for a given campaign. This endpoint is
// useful for accessing detailed information about a particular budget item
// associated with a marketing campaign.
func (r *CampaignBudgetService) Get(ctx context.Context, budgetID int64, query CampaignBudgetGetParams, opts ...option.RequestOption) (res *PublicBudgetItem, err error) {
	opts = slices.Concat(r.Options, opts)
	if query.CampaignGuid == "" {
		err = errors.New("missing required campaignGuid parameter")
		return nil, err
	}
	path := fmt.Sprintf("marketing/campaigns/2026-03/%s/budget/%v", query.CampaignGuid, budgetID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Retrieve budget and spending items along with their totals for a specific
// campaign. This endpoint provides insights into the financial allocations and
// expenditures associated with the campaign, helping users to manage and analyze
// campaign budgets effectively.
func (r *CampaignBudgetService) GetTotals(ctx context.Context, campaignGuid string, opts ...option.RequestOption) (res *PublicBudgetTotals, err error) {
	opts = slices.Concat(r.Options, opts)
	if campaignGuid == "" {
		err = errors.New("missing required campaignGuid parameter")
		return nil, err
	}
	path := fmt.Sprintf("marketing/campaigns/2026-03/%s/budget/totals", campaignGuid)
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
