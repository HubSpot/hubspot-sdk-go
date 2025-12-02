// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package marketing

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

// SubscriptionV4LinkService contains methods and other services that help with
// interacting with the hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSubscriptionV4LinkService] method instead.
type SubscriptionV4LinkService struct {
	Options []option.RequestOption
}

// NewSubscriptionV4LinkService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewSubscriptionV4LinkService(opts ...option.RequestOption) (r SubscriptionV4LinkService) {
	r = SubscriptionV4LinkService{}
	r.Options = opts
	return
}

func (r *SubscriptionV4LinkService) New(ctx context.Context, params SubscriptionV4LinkNewParams, opts ...option.RequestOption) (res *LinkGenerationResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "communication-preferences/v4/links/generate"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

type SubscriptionV4LinkNewParams struct {
	// Any of "EMAIL".
	Channel               SubscriptionV4LinkNewParamsChannel `query:"channel,omitzero,required" json:"-"`
	LinkGenerationRequest LinkGenerationRequestParam
	BusinessUnitID        param.Opt[int64] `query:"businessUnitId,omitzero" json:"-"`
	paramObj
}

func (r SubscriptionV4LinkNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.LinkGenerationRequest)
}
func (r *SubscriptionV4LinkNewParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.LinkGenerationRequest)
}

// URLQuery serializes [SubscriptionV4LinkNewParams]'s query parameters as
// `url.Values`.
func (r SubscriptionV4LinkNewParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SubscriptionV4LinkNewParamsChannel string

const (
	SubscriptionV4LinkNewParamsChannelEmail SubscriptionV4LinkNewParamsChannel = "EMAIL"
)
