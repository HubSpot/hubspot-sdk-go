// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package conversations

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/stainless-sdks/hubspot-sdk-go/internal/apiquery"
	shimjson "github.com/stainless-sdks/hubspot-sdk-go/internal/encoding/json"
	"github.com/stainless-sdks/hubspot-sdk-go/internal/requestconfig"
	"github.com/stainless-sdks/hubspot-sdk-go/option"
	"github.com/stainless-sdks/hubspot-sdk-go/packages/pagination"
	"github.com/stainless-sdks/hubspot-sdk-go/packages/param"
)

// CustomChannelChannelAccountService contains methods and other services that help
// with interacting with the hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCustomChannelChannelAccountService] method instead.
type CustomChannelChannelAccountService struct {
	Options []option.RequestOption
}

// NewCustomChannelChannelAccountService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewCustomChannelChannelAccountService(opts ...option.RequestOption) (r CustomChannelChannelAccountService) {
	r = CustomChannelChannelAccountService{}
	r.Options = opts
	return
}

// Create a new account for a channel. Multiple accounts can communicate over a
// single channel using different delivery identifiers.
func (r *CustomChannelChannelAccountService) New(ctx context.Context, channelID int64, body CustomChannelChannelAccountNewParams, opts ...option.RequestOption) (res *PublicChannelAccount, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("conversations/v3/custom-channels/%v/channel-accounts", channelID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// This API is used to update the name of the channel account and it's isAuthorized
// status. Setting to isAuthorized flag to False disables the channel account.
func (r *CustomChannelChannelAccountService) Update(ctx context.Context, channelAccountID int64, params CustomChannelChannelAccountUpdateParams, opts ...option.RequestOption) (res *PublicChannelAccount, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("conversations/v3/custom-channels/%v/channel-accounts/%v", params.ChannelID, channelAccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return
}

// Retrieve a list of accounts for a custom channel.
func (r *CustomChannelChannelAccountService) List(ctx context.Context, channelID int64, query CustomChannelChannelAccountListParams, opts ...option.RequestOption) (res *pagination.Page[PublicChannelAccount], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := fmt.Sprintf("conversations/v3/custom-channels/%v/channel-accounts", channelID)
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

// Retrieve a list of accounts for a custom channel.
func (r *CustomChannelChannelAccountService) ListAutoPaging(ctx context.Context, channelID int64, query CustomChannelChannelAccountListParams, opts ...option.RequestOption) *pagination.PageAutoPager[PublicChannelAccount] {
	return pagination.NewPageAutoPager(r.List(ctx, channelID, query, opts...))
}

// Retrieve the details for a specific channel account. This contains all the
// metadata about your channel account, including its channel, associated inbox id,
// and delivery identifier information.
func (r *CustomChannelChannelAccountService) Get(ctx context.Context, channelAccountID int64, params CustomChannelChannelAccountGetParams, opts ...option.RequestOption) (res *PublicChannelAccount, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("conversations/v3/custom-channels/%v/channel-accounts/%v", params.ChannelID, channelAccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return
}

type CustomChannelChannelAccountNewParams struct {
	PublicChannelAccountEgg PublicChannelAccountEggParam
	paramObj
}

func (r CustomChannelChannelAccountNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.PublicChannelAccountEgg)
}
func (r *CustomChannelChannelAccountNewParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.PublicChannelAccountEgg)
}

type CustomChannelChannelAccountUpdateParams struct {
	ChannelID                         int64 `path:"channelId,required" json:"-"`
	PublicChannelAccountUpdateRequest PublicChannelAccountUpdateRequestParam
	paramObj
}

func (r CustomChannelChannelAccountUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.PublicChannelAccountUpdateRequest)
}
func (r *CustomChannelChannelAccountUpdateParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.PublicChannelAccountUpdateRequest)
}

type CustomChannelChannelAccountListParams struct {
	After                   param.Opt[string] `query:"after,omitzero" json:"-"`
	Archived                param.Opt[bool]   `query:"archived,omitzero" json:"-"`
	DefaultPageLength       param.Opt[int64]  `query:"defaultPageLength,omitzero" json:"-"`
	Limit                   param.Opt[int64]  `query:"limit,omitzero" json:"-"`
	DeliveryIdentifierType  []string          `query:"deliveryIdentifierType,omitzero" json:"-"`
	DeliveryIdentifierValue []string          `query:"deliveryIdentifierValue,omitzero" json:"-"`
	Sort                    []string          `query:"sort,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [CustomChannelChannelAccountListParams]'s query parameters
// as `url.Values`.
func (r CustomChannelChannelAccountListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type CustomChannelChannelAccountGetParams struct {
	ChannelID int64 `path:"channelId,required" json:"-"`
	// Filter results to include only archived or non-archived channel accounts.
	Archived param.Opt[bool] `query:"archived,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [CustomChannelChannelAccountGetParams]'s query parameters as
// `url.Values`.
func (r CustomChannelChannelAccountGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
