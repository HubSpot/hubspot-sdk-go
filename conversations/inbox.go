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

// InboxService contains methods and other services that help with interacting with
// the hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewInboxService] method instead.
type InboxService struct {
	Options []option.RequestOption
}

// NewInboxService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewInboxService(opts ...option.RequestOption) (r InboxService) {
	r = InboxService{}
	r.Options = opts
	return
}

func (r *InboxService) List(ctx context.Context, query InboxListParams, opts ...option.RequestOption) (res *pagination.Page[PublicInbox], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "conversations/v3/conversations/inboxes"
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

func (r *InboxService) ListAutoPaging(ctx context.Context, query InboxListParams, opts ...option.RequestOption) *pagination.PageAutoPager[PublicInbox] {
	return pagination.NewPageAutoPager(r.List(ctx, query, opts...))
}

func (r *InboxService) Get(ctx context.Context, inboxID int64, query InboxGetParams, opts ...option.RequestOption) (res *PublicInbox, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("conversations/v3/conversations/inboxes/%v", inboxID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type InboxListParams struct {
	After             param.Opt[string] `query:"after,omitzero" json:"-"`
	Archived          param.Opt[bool]   `query:"archived,omitzero" json:"-"`
	DefaultPageLength param.Opt[int64]  `query:"defaultPageLength,omitzero" json:"-"`
	Limit             param.Opt[int64]  `query:"limit,omitzero" json:"-"`
	Sort              []string          `query:"sort,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [InboxListParams]'s query parameters as `url.Values`.
func (r InboxListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type InboxGetParams struct {
	Archived param.Opt[bool] `query:"archived,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [InboxGetParams]'s query parameters as `url.Values`.
func (r InboxGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
