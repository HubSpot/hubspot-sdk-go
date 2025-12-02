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

	"github.com/stainless-sdks/hubspot-sdk-go/internal/apijson"
	"github.com/stainless-sdks/hubspot-sdk-go/internal/apiquery"
	"github.com/stainless-sdks/hubspot-sdk-go/internal/requestconfig"
	"github.com/stainless-sdks/hubspot-sdk-go/option"
	"github.com/stainless-sdks/hubspot-sdk-go/packages/pagination"
	"github.com/stainless-sdks/hubspot-sdk-go/packages/param"
	"github.com/stainless-sdks/hubspot-sdk-go/packages/respjson"
	"github.com/stainless-sdks/hubspot-sdk-go/shared"
)

// TaxRateService contains methods and other services that help with interacting
// with the hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTaxRateService] method instead.
type TaxRateService struct {
	Options []option.RequestOption
}

// NewTaxRateService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewTaxRateService(opts ...option.RequestOption) (r TaxRateService) {
	r = TaxRateService{}
	r.Options = opts
	return
}

// Retrieve a paginated list of all tax rates set up in the account tax rate
// library
func (r *TaxRateService) List(ctx context.Context, query TaxRateListParams, opts ...option.RequestOption) (res *pagination.Page[PublicTaxRateGroup], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "tax-rates/v1/tax-rates"
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
	opts = slices.Concat(r.Options, opts)
	if taxRateGroupID == "" {
		err = errors.New("missing required taxRateGroupId parameter")
		return
	}
	path := fmt.Sprintf("tax-rates/v1/tax-rates/%s", taxRateGroupID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type CollectionResponsePublicTaxRateGroupForwardPaging struct {
	Results []PublicTaxRateGroup `json:"results,required"`
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
	ID             string    `json:"id,required"`
	Active         bool      `json:"active,required"`
	CreatedAt      time.Time `json:"createdAt,required" format:"date-time"`
	Label          string    `json:"label,required"`
	Name           string    `json:"name,required"`
	PercentageRate float64   `json:"percentageRate,required"`
	UpdatedAt      time.Time `json:"updatedAt,required" format:"date-time"`
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
	// Include inactive rates.
	Active param.Opt[bool] `query:"active,omitzero" json:"-"`
	// The paging cursor token of the last successfully read resource will be returned
	// as the paging.next.after JSON property of a paged response containing more
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
