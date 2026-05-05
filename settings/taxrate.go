// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package settings

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/HubSpot/hubspot-sdk-go/internal/apijson"
	"github.com/HubSpot/hubspot-sdk-go/internal/apiquery"
	"github.com/HubSpot/hubspot-sdk-go/internal/requestconfig"
	"github.com/HubSpot/hubspot-sdk-go/option"
	"github.com/HubSpot/hubspot-sdk-go/packages/pagination"
	"github.com/HubSpot/hubspot-sdk-go/packages/param"
	"github.com/HubSpot/hubspot-sdk-go/packages/respjson"
	"github.com/HubSpot/hubspot-sdk-go/shared"
)

// TaxRateService contains methods and other services that help with interacting
// with the hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTaxRateService] method instead.
type TaxRateService struct {
	options []option.RequestOption
}

// NewTaxRateService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewTaxRateService(opts ...option.RequestOption) (r TaxRateService) {
	r = TaxRateService{}
	r.options = opts
	return
}

// Retrieve a paginated list of all tax rates set up in the account tax rate
// library
func (r *TaxRateService) List(ctx context.Context, query TaxRateListParams, opts ...option.RequestOption) (res *pagination.Page[PublicTaxRateGroup], err error) {
	var raw *http.Response
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "tax-rates/2026-03/tax-rates"
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
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

// Retrieve a paginated list of all tax rates set up in the account tax rate
// library
func (r *TaxRateService) ListAutoPaging(ctx context.Context, query TaxRateListParams, opts ...option.RequestOption) *pagination.PageAutoPager[PublicTaxRateGroup] {
	return pagination.NewPageAutoPager(r.List(ctx, query, opts...))
}

// Retrieve a specific tax rate by its `taxRateGroupId`.
func (r *TaxRateService) Get(ctx context.Context, taxRateGroupID string, opts ...option.RequestOption) (res *PublicTaxRateGroup, err error) {
	opts = slices.Concat(r.options, opts)
	if taxRateGroupID == "" {
		err = errors.New("missing required taxRateGroupId parameter")
		return nil, err
	}
	path := fmt.Sprintf("tax-rates/2026-03/tax-rates/%s", url.PathEscape(taxRateGroupID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type CollectionResponsePublicTaxRateGroupForwardPaging struct {
	Results []PublicTaxRateGroup `json:"results" api:"required"`
	Paging  shared.ForwardPaging `json:"paging"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Results     respjson.Field
		Paging      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CollectionResponsePublicTaxRateGroupForwardPaging) RawJSON() string { return r.JSON.raw }
func (r *CollectionResponsePublicTaxRateGroupForwardPaging) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PublicTaxRateGroup struct {
	// The unique identifier for the tax rate.
	ID string `json:"id" api:"required"`
	// Indicates whether the tax rate group is currently active.
	Active bool `json:"active" api:"required"`
	// The date and time when the tax rate was created.
	CreatedAt time.Time `json:"createdAt" api:"required" format:"date-time"`
	// The display label for the tax rate.
	Label string `json:"label" api:"required"`
	// The name of the tax rate.
	Name string `json:"name" api:"required"`
	// The percentage rate applied.
	PercentageRate float64 `json:"percentageRate" api:"required"`
	// The date and time when the tax rate was last updated.
	UpdatedAt time.Time `json:"updatedAt" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID             respjson.Field
		Active         respjson.Field
		CreatedAt      respjson.Field
		Label          respjson.Field
		Name           respjson.Field
		PercentageRate respjson.Field
		UpdatedAt      respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PublicTaxRateGroup) RawJSON() string { return r.JSON.raw }
func (r *PublicTaxRateGroup) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TaxRateListParams struct {
	Active param.Opt[bool] `query:"active,omitzero" json:"-"`
	// The paging cursor token of the last successfully read resource will be returned
	// as the `paging.next.after` JSON property of a paged response containing more
	// results.
	After param.Opt[string] `query:"after,omitzero" json:"-"`
	// The maximum number of results to display per page.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [TaxRateListParams]'s query parameters as `url.Values`.
func (r TaxRateListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
