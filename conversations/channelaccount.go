// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package conversations

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/stainless-sdks/hubspot-sdk-go/internal/apiquery"
	"github.com/stainless-sdks/hubspot-sdk-go/internal/requestconfig"
	"github.com/stainless-sdks/hubspot-sdk-go/option"
	"github.com/stainless-sdks/hubspot-sdk-go/packages/pagination"
	"github.com/stainless-sdks/hubspot-sdk-go/packages/param"
)

// ChannelAccountService contains methods and other services that help with
// interacting with the hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewChannelAccountService] method instead.
type ChannelAccountService struct {
	Options []option.RequestOption
}

// NewChannelAccountService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewChannelAccountService(opts ...option.RequestOption) (r ChannelAccountService) {
	r = ChannelAccountService{}
	r.Options = opts
	return
}

func (r *ChannelAccountService) List(ctx context.Context, query ChannelAccountListParams, opts ...option.RequestOption) (res *pagination.Page[PublicChannelAccount], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "conversations/v3/conversations/channel-accounts"
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

func (r *ChannelAccountService) ListAutoPaging(ctx context.Context, query ChannelAccountListParams, opts ...option.RequestOption) *pagination.PageAutoPager[PublicChannelAccount] {
	return pagination.NewPageAutoPager(r.List(ctx, query, opts...))
}

func (r *ChannelAccountService) Get(ctx context.Context, channelAccountID int64, query ChannelAccountGetParams, opts ...option.RequestOption) (res *PublicChannelAccount, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("conversations/v3/conversations/channel-accounts/%v", channelAccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type ChannelAccountListParams struct {
	After             param.Opt[string] `query:"after,omitzero" json:"-"`
	Archived          param.Opt[bool]   `query:"archived,omitzero" json:"-"`
	DefaultPageLength param.Opt[int64]  `query:"defaultPageLength,omitzero" json:"-"`
	Limit             param.Opt[int64]  `query:"limit,omitzero" json:"-"`
	ChannelID         []int64           `query:"channelId,omitzero" json:"-"`
	InboxID           []int64           `query:"inboxId,omitzero" json:"-"`
	Sort              []string          `query:"sort,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ChannelAccountListParams]'s query parameters as
// `url.Values`.
func (r ChannelAccountListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ChannelAccountGetParams struct {
	Archived param.Opt[bool] `query:"archived,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ChannelAccountGetParams]'s query parameters as
// `url.Values`.
func (r ChannelAccountGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
