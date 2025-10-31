// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package settings

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"
	"time"

	"github.com/stainless-sdks/hubspot-sdk-go/internal/apijson"
	"github.com/stainless-sdks/hubspot-sdk-go/internal/requestconfig"
	"github.com/stainless-sdks/hubspot-sdk-go/option"
	"github.com/stainless-sdks/hubspot-sdk-go/packages/respjson"
	"github.com/stainless-sdks/hubspot-sdk-go/shared"
)

// TaxRateService contains methods and other services that help with interacting
// with the Hubspot API.
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
func (r *TaxRateService) List(ctx context.Context, opts ...option.RequestOption) (res *TaxRateListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "tax-rates/v1/tax-rates"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Retrieve a specific tax rate by its `taxRateGroupId`.
func (r *TaxRateService) Get(ctx context.Context, taxRateGroupID string, opts ...option.RequestOption) (res *TaxRateGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if taxRateGroupID == "" {
		err = errors.New("missing required taxRateGroupId parameter")
		return
	}
	path := fmt.Sprintf("tax-rates/v1/tax-rates/%s", taxRateGroupID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type TaxRateListResponse struct {
	Results []TaxRateListResponseResult `json:"results,required"`
	Paging  shared.ForwardPaging        `json:"paging"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Results     respjson.Field
		Paging      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TaxRateListResponse) RawJSON() string { return r.JSON.raw }
func (r *TaxRateListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TaxRateListResponseResult struct {
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
func (r TaxRateListResponseResult) RawJSON() string { return r.JSON.raw }
func (r *TaxRateListResponseResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TaxRateGetResponse struct {
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
func (r TaxRateGetResponse) RawJSON() string { return r.JSON.raw }
func (r *TaxRateGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
