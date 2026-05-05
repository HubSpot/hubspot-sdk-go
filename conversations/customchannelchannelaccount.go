// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package conversations

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/HubSpot/hubspot-sdk-go/internal/apijson"
	"github.com/HubSpot/hubspot-sdk-go/internal/apiquery"
	shimjson "github.com/HubSpot/hubspot-sdk-go/internal/encoding/json"
	"github.com/HubSpot/hubspot-sdk-go/internal/requestconfig"
	"github.com/HubSpot/hubspot-sdk-go/option"
	"github.com/HubSpot/hubspot-sdk-go/packages/pagination"
	"github.com/HubSpot/hubspot-sdk-go/packages/param"
)

// CustomChannelChannelAccountService contains methods and other services that help
// with interacting with the hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCustomChannelChannelAccountService] method instead.
type CustomChannelChannelAccountService struct {
	options []option.RequestOption
}

// NewCustomChannelChannelAccountService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewCustomChannelChannelAccountService(opts ...option.RequestOption) (r CustomChannelChannelAccountService) {
	r = CustomChannelChannelAccountService{}
	r.options = opts
	return
}

// Create a new account for a channel. Multiple accounts can communicate over a
// single channel using different delivery identifiers.
func (r *CustomChannelChannelAccountService) New(ctx context.Context, channelID int64, body CustomChannelChannelAccountNewParams, opts ...option.RequestOption) (res *PublicChannelAccount, err error) {
	opts = slices.Concat(r.options, opts)
	path := fmt.Sprintf("conversations/custom-channels/2026-03/%v/channel-accounts", channelID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// This API is used to update the name of the channel account and it's isAuthorized
// status. Setting to isAuthorized flag to False disables the channel account.
func (r *CustomChannelChannelAccountService) Update(ctx context.Context, channelAccountID int64, params CustomChannelChannelAccountUpdateParams, opts ...option.RequestOption) (res *PublicChannelAccount, err error) {
	opts = slices.Concat(r.options, opts)
	path := fmt.Sprintf("conversations/custom-channels/2026-03/%v/channel-accounts/%v", params.ChannelID, channelAccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return res, err
}

// Retrieve a list of accounts for a custom channel.
func (r *CustomChannelChannelAccountService) List(ctx context.Context, channelID int64, query CustomChannelChannelAccountListParams, opts ...option.RequestOption) (res *pagination.Page[PublicChannelAccount], err error) {
	var raw *http.Response
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := fmt.Sprintf("conversations/custom-channels/2026-03/%v/channel-accounts", channelID)
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

// Update a channel account staging token's account name and delivery identifier.
// This information will be applied to the channel account created from this
// staging token. This is used for public apps.
func (r *CustomChannelChannelAccountService) UpdateStagingToken(ctx context.Context, accountToken string, params CustomChannelChannelAccountUpdateStagingTokenParams, opts ...option.RequestOption) (res *PublicChannelAccountStagingToken, err error) {
	opts = slices.Concat(r.options, opts)
	if accountToken == "" {
		err = errors.New("missing required accountToken parameter")
		return nil, err
	}
	path := fmt.Sprintf("conversations/custom-channels/2026-03/%v/channel-account-staging-tokens/%s", params.ChannelID, url.PathEscape(accountToken))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return res, err
}

type CustomChannelChannelAccountNewParams struct {
	PublicChannelAccountEgg PublicChannelAccountEggParam
	paramObj
}

func (r CustomChannelChannelAccountNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.PublicChannelAccountEgg)
}
func (r *CustomChannelChannelAccountNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CustomChannelChannelAccountUpdateParams struct {
	ChannelID                         int64 `path:"channelId" api:"required" json:"-"`
	PublicChannelAccountUpdateRequest PublicChannelAccountUpdateRequestParam
	paramObj
}

func (r CustomChannelChannelAccountUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.PublicChannelAccountUpdateRequest)
}
func (r *CustomChannelChannelAccountUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CustomChannelChannelAccountListParams struct {
	// The paging cursor token of the last successfully read resource will be returned
	// as the `paging.next.after` JSON property of a paged response containing more
	// results.
	After param.Opt[string] `query:"after,omitzero" json:"-"`
	// Whether to return only results that have been archived.
	Archived          param.Opt[bool]  `query:"archived,omitzero" json:"-"`
	DefaultPageLength param.Opt[int64] `query:"defaultPageLength,omitzero" json:"-"`
	// The maximum number of results to display per page.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Any of "HS_EMAIL_ADDRESS", "HS_PHONE_NUMBER", "HS_SHORT_CODE",
	// "CHANNEL_SPECIFIC_OPAQUE_ID".
	DeliveryIdentifierType  []string `query:"deliveryIdentifierType,omitzero" json:"-"`
	DeliveryIdentifierValue []string `query:"deliveryIdentifierValue,omitzero" json:"-"`
	Sort                    []string `query:"sort,omitzero" json:"-"`
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

type CustomChannelChannelAccountUpdateStagingTokenParams struct {
	ChannelID                                     int64 `path:"channelId" api:"required" json:"-"`
	PublicChannelAccountStagingTokenUpdateRequest PublicChannelAccountStagingTokenUpdateRequestParam
	paramObj
}

func (r CustomChannelChannelAccountUpdateStagingTokenParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.PublicChannelAccountStagingTokenUpdateRequest)
}
func (r *CustomChannelChannelAccountUpdateStagingTokenParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
