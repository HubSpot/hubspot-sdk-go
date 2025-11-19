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

// CampaignSpendService contains methods and other services that help with
// interacting with the Hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCampaignSpendService] method instead.
type CampaignSpendService struct {
	Options []option.RequestOption
}

// NewCampaignSpendService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCampaignSpendService(opts ...option.RequestOption) (r CampaignSpendService) {
	r = CampaignSpendService{}
	r.Options = opts
	return
}

// Create a new campaign spend item
func (r *CampaignSpendService) New(ctx context.Context, campaignGuid string, body CampaignSpendNewParams, opts ...option.RequestOption) (res *PublicSpendItem, err error) {
	opts = slices.Concat(r.Options, opts)
	if campaignGuid == "" {
		err = errors.New("missing required campaignGuid parameter")
		return
	}
	path := fmt.Sprintf("marketing/v3/campaigns/%s/spend", campaignGuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Update a specific campaign spend item by ID
func (r *CampaignSpendService) Update(ctx context.Context, spendID int64, params CampaignSpendUpdateParams, opts ...option.RequestOption) (res *PublicSpendItem, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.CampaignGuid == "" {
		err = errors.New("missing required campaignGuid parameter")
		return
	}
	path := fmt.Sprintf("marketing/v3/campaigns/%s/spend/%v", params.CampaignGuid, spendID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return
}

// Delete a specific campaign spend item by ID
func (r *CampaignSpendService) Delete(ctx context.Context, spendID int64, body CampaignSpendDeleteParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if body.CampaignGuid == "" {
		err = errors.New("missing required campaignGuid parameter")
		return
	}
	path := fmt.Sprintf("marketing/v3/campaigns/%s/spend/%v", body.CampaignGuid, spendID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Read a campaign spend item by its spendId
func (r *CampaignSpendService) Get(ctx context.Context, spendID int64, query CampaignSpendGetParams, opts ...option.RequestOption) (res *PublicSpendItem, err error) {
	opts = slices.Concat(r.Options, opts)
	if query.CampaignGuid == "" {
		err = errors.New("missing required campaignGuid parameter")
		return
	}
	path := fmt.Sprintf("marketing/v3/campaigns/%s/spend/%v", query.CampaignGuid, spendID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type CampaignSpendNewParams struct {
	PublicSpendItemInput PublicSpendItemInputParam
	paramObj
}

func (r CampaignSpendNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.PublicSpendItemInput)
}
func (r *CampaignSpendNewParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.PublicSpendItemInput)
}

type CampaignSpendUpdateParams struct {
	CampaignGuid         string `path:"campaignGuid,required" json:"-"`
	PublicSpendItemInput PublicSpendItemInputParam
	paramObj
}

func (r CampaignSpendUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.PublicSpendItemInput)
}
func (r *CampaignSpendUpdateParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.PublicSpendItemInput)
}

type CampaignSpendDeleteParams struct {
	CampaignGuid string `path:"campaignGuid,required" json:"-"`
	paramObj
}

type CampaignSpendGetParams struct {
	CampaignGuid string `path:"campaignGuid,required" json:"-"`
	paramObj
}
