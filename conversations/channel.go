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

// ChannelService contains methods and other services that help with interacting
// with the hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewChannelService] method instead.
type ChannelService struct {
	Options []option.RequestOption
}

// NewChannelService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewChannelService(opts ...option.RequestOption) (r ChannelService) {
	r = ChannelService{}
	r.Options = opts
	return
}

func (r *ChannelService) List(ctx context.Context, query ChannelListParams, opts ...option.RequestOption) (res *pagination.Page[PublicChannel], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "conversations/v3/conversations/channels"
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

func (r *ChannelService) ListAutoPaging(ctx context.Context, query ChannelListParams, opts ...option.RequestOption) *pagination.PageAutoPager[PublicChannel] {
	return pagination.NewPageAutoPager(r.List(ctx, query, opts...))
}

func (r *ChannelService) Get(ctx context.Context, channelID int64, opts ...option.RequestOption) (res *PublicChannel, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("conversations/v3/conversations/channels/%v", channelID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type ChannelListParams struct {
	After             param.Opt[string] `query:"after,omitzero" json:"-"`
	DefaultPageLength param.Opt[int64]  `query:"defaultPageLength,omitzero" json:"-"`
	Limit             param.Opt[int64]  `query:"limit,omitzero" json:"-"`
	Sort              []string          `query:"sort,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ChannelListParams]'s query parameters as `url.Values`.
func (r ChannelListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
