// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package communication_preferences

import (
	"context"
	"net/http"
	"net/url"
	"slices"

	"github.com/stainless-sdks/hubspot-sdk-go/internal/apiquery"
	"github.com/stainless-sdks/hubspot-sdk-go/internal/requestconfig"
	"github.com/stainless-sdks/hubspot-sdk-go/option"
	"github.com/stainless-sdks/hubspot-sdk-go/packages/param"
)

// DefinitionService contains methods and other services that help with interacting
// with the hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDefinitionService] method instead.
type DefinitionService struct {
	options []option.RequestOption
}

// NewDefinitionService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewDefinitionService(opts ...option.RequestOption) (r DefinitionService) {
	r = DefinitionService{}
	r.options = opts
	return
}

// Get a list of subscription status definitions from the account.
func (r *DefinitionService) List(ctx context.Context, query DefinitionListParams, opts ...option.RequestOption) (res *ActionResponseWithResultsSubscriptionDefinition, err error) {
	opts = slices.Concat(r.options, opts)
	path := "communication-preferences/2026-03/definitions"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

type DefinitionListParams struct {
	// An integer representing the ID of the business unit for which to retrieve
	// subscription definitions.
	BusinessUnitID param.Opt[int64] `query:"businessUnitId,omitzero" json:"-"`
	// A boolean indicating whether to include translations of the communication
	// preferences definitions in the response.
	IncludeTranslations param.Opt[bool] `query:"includeTranslations,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [DefinitionListParams]'s query parameters as `url.Values`.
func (r DefinitionListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
