// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package business_units

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/HubSpot/hubspot-sdk-go/internal/apiquery"
	"github.com/HubSpot/hubspot-sdk-go/internal/requestconfig"
	"github.com/HubSpot/hubspot-sdk-go/option"
)

// BusinessUnitEntryService contains methods and other services that help with
// interacting with the hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBusinessUnitEntryService] method instead.
type BusinessUnitEntryService struct {
	options []option.RequestOption
}

// NewBusinessUnitEntryService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewBusinessUnitEntryService(opts ...option.RequestOption) (r BusinessUnitEntryService) {
	r = BusinessUnitEntryService{}
	r.options = opts
	return
}

// Retrieve the brands that a specific user can access.
func (r *BusinessUnitEntryService) GetByUserID(ctx context.Context, userID string, query BusinessUnitEntryGetByUserIDParams, opts ...option.RequestOption) (res *CollectionResponsePublicBusinessUnitNoPaging, err error) {
	opts = slices.Concat(r.options, opts)
	if userID == "" {
		err = errors.New("missing required userId parameter")
		return nil, err
	}
	path := fmt.Sprintf("business-units/public/2026-03/business-units/user/%s", url.PathEscape(userID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

type BusinessUnitEntryGetByUserIDParams struct {
	Name       []string `query:"name,omitzero" json:"-"`
	Properties []string `query:"properties,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [BusinessUnitEntryGetByUserIDParams]'s query parameters as
// `url.Values`.
func (r BusinessUnitEntryGetByUserIDParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
