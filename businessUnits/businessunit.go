// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package businessUnits

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/stainless-sdks/hubspot-sdk-go/internal/apijson"
	"github.com/stainless-sdks/hubspot-sdk-go/internal/apiquery"
	"github.com/stainless-sdks/hubspot-sdk-go/internal/requestconfig"
	"github.com/stainless-sdks/hubspot-sdk-go/marketing"
	"github.com/stainless-sdks/hubspot-sdk-go/option"
	"github.com/stainless-sdks/hubspot-sdk-go/packages/respjson"
)

// BusinessUnitService contains methods and other services that help with
// interacting with the Hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBusinessUnitService] method instead.
type BusinessUnitService struct {
	Options []option.RequestOption
}

// NewBusinessUnitService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewBusinessUnitService(opts ...option.RequestOption) (r BusinessUnitService) {
	r = BusinessUnitService{}
	r.Options = opts
	return
}

// Get Business Units identified by `userId`. The `userId` refers to the user’s ID.
func (r *BusinessUnitService) GetByUserID(ctx context.Context, userID string, query BusinessUnitGetByUserIDParams, opts ...option.RequestOption) (res *BusinessUnitGetByUserIDResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if userID == "" {
		err = errors.New("missing required userId parameter")
		return
	}
	path := fmt.Sprintf("business-units/v3/business-units/user/%s", userID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// A response object containing a collection of Business Units
type BusinessUnitGetByUserIDResponse struct {
	// The collection of Business Units
	Results []marketing.PublicBusinessUnit `json:"results,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Results     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BusinessUnitGetByUserIDResponse) RawJSON() string { return r.JSON.raw }
func (r *BusinessUnitGetByUserIDResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BusinessUnitGetByUserIDParams struct {
	// The names of Business Units to retrieve. If empty or not provided, then all
	// associated Business Units will be returned.
	Name []string `query:"name,omitzero" json:"-"`
	// The names of properties to optionally include in the response body. The only
	// valid value is `logoMetadata`.
	Properties []string `query:"properties,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [BusinessUnitGetByUserIDParams]'s query parameters as
// `url.Values`.
func (r BusinessUnitGetByUserIDParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
