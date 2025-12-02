// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package crm

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"
	"slices"

	"github.com/stainless-sdks/hubspot-sdk-go/internal/apiquery"
	shimjson "github.com/stainless-sdks/hubspot-sdk-go/internal/encoding/json"
	"github.com/stainless-sdks/hubspot-sdk-go/internal/requestconfig"
	"github.com/stainless-sdks/hubspot-sdk-go/option"
	"github.com/stainless-sdks/hubspot-sdk-go/packages/param"
)

// ListMappingService contains methods and other services that help with
// interacting with the hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewListMappingService] method instead.
type ListMappingService struct {
	Options []option.RequestOption
}

// NewListMappingService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewListMappingService(opts ...option.RequestOption) (r ListMappingService) {
	r = ListMappingService{}
	r.Options = opts
	return
}

// This API allows translation of a batch of legacy list id's to list id's. This
// allows for a maximum of 10,000 id's. This is a temporary API allowed for mapping
// old id's to new id's and will expire on May 30th, 2025.
func (r *ListMappingService) BatchNewIDMapping(ctx context.Context, body ListMappingBatchNewIDMappingParams, opts ...option.RequestOption) (res *PublicBatchMigrationMapping, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "crm/v3/lists/idmapping"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// This API allows translation of legacy list id to list id. This is a temporary
// API allowed for mapping old id's to new id's and will expire on May 30th, 2025.
func (r *ListMappingService) GetIDMapping(ctx context.Context, query ListMappingGetIDMappingParams, opts ...option.RequestOption) (res *PublicMigrationMapping, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "crm/v3/lists/idmapping"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type ListMappingBatchNewIDMappingParams struct {
	Body []string
	paramObj
}

func (r ListMappingBatchNewIDMappingParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.Body)
}
func (r *ListMappingBatchNewIDMappingParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.Body)
}

type ListMappingGetIDMappingParams struct {
	// The legacy list id from lists v1 API.
	LegacyListID param.Opt[string] `query:"legacyListId,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ListMappingGetIDMappingParams]'s query parameters as
// `url.Values`.
func (r ListMappingGetIDMappingParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
