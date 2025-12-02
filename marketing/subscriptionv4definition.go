// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package marketing

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

// SubscriptionV4DefinitionService contains methods and other services that help
// with interacting with the hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSubscriptionV4DefinitionService] method instead.
type SubscriptionV4DefinitionService struct {
	Options []option.RequestOption
}

// NewSubscriptionV4DefinitionService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewSubscriptionV4DefinitionService(opts ...option.RequestOption) (r SubscriptionV4DefinitionService) {
	r = SubscriptionV4DefinitionService{}
	r.Options = opts
	return
}

// Get a list of subscription status definitions from the account.
func (r *SubscriptionV4DefinitionService) List(ctx context.Context, query SubscriptionV4DefinitionListParams, opts ...option.RequestOption) (res *ActionResponseWithResultsSubscriptionDefinition, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "communication-preferences/v4/definitions"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type SubscriptionV4DefinitionListParams struct {
	// If you have the
	// [business unit add-on](https://developers.hubspot.com/beta-docs/guides/api/settings/business-units-api),
	// include this parameter to filter results by business unit ID. The default
	// Account business unit will always use `0`.
	BusinessUnitID param.Opt[int64] `query:"businessUnitId,omitzero" json:"-"`
	// Set to `true` to return subscription translations associated with each
	// definition.
	IncludeTranslations param.Opt[bool] `query:"includeTranslations,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SubscriptionV4DefinitionListParams]'s query parameters as
// `url.Values`.
func (r SubscriptionV4DefinitionListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
