// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package conversations

import (
	"context"
	"encoding/json"
	"errors"
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

// MessageService contains methods and other services that help with interacting
// with the hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMessageService] method instead.
type MessageService struct {
	Options []option.RequestOption
}

// NewMessageService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewMessageService(opts ...option.RequestOption) (r MessageService) {
	r = MessageService{}
	r.Options = opts
	return
}

func (r *MessageService) New(ctx context.Context, threadID int64, body MessageNewParams, opts ...option.RequestOption) (res *PublicMessageUnion, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("conversations/v3/conversations/threads/%v/messages", threadID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

func (r *MessageService) List(ctx context.Context, threadID int64, query MessageListParams, opts ...option.RequestOption) (res *pagination.Page[PublicMessageUnion], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := fmt.Sprintf("conversations/v3/conversations/threads/%v/messages", threadID)
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

func (r *MessageService) ListAutoPaging(ctx context.Context, threadID int64, query MessageListParams, opts ...option.RequestOption) *pagination.PageAutoPager[PublicMessageUnion] {
	return pagination.NewPageAutoPager(r.List(ctx, threadID, query, opts...))
}

func (r *MessageService) Get(ctx context.Context, messageID string, params MessageGetParams, opts ...option.RequestOption) (res *PublicMessageUnion, err error) {
	opts = slices.Concat(r.Options, opts)
	if messageID == "" {
		err = errors.New("missing required messageId parameter")
		return
	}
	path := fmt.Sprintf("conversations/v3/conversations/threads/%v/messages/%s", params.ThreadID, messageID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return
}

func (r *MessageService) GetOriginalContent(ctx context.Context, messageID string, params MessageGetOriginalContentParams, opts ...option.RequestOption) (res *PublicMessageContent, err error) {
	opts = slices.Concat(r.Options, opts)
	if messageID == "" {
		err = errors.New("missing required messageId parameter")
		return
	}
	path := fmt.Sprintf("conversations/v3/conversations/threads/%v/messages/%s/original-content", params.ThreadID, messageID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return
}

type MessageNewParams struct {
	PublicMessageEgg PublicMessageEggUnionParam
	paramObj
}

func (r MessageNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.PublicMessageEgg)
}
func (r *MessageNewParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.PublicMessageEgg)
}

type MessageListParams struct {
	After    param.Opt[string] `query:"after,omitzero" json:"-"`
	Archived param.Opt[bool]   `query:"archived,omitzero" json:"-"`
	Limit    param.Opt[int64]  `query:"limit,omitzero" json:"-"`
	Property param.Opt[string] `query:"property,omitzero" json:"-"`
	Sort     []string          `query:"sort,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [MessageListParams]'s query parameters as `url.Values`.
func (r MessageListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type MessageGetParams struct {
	ThreadID int64             `path:"threadId,required" json:"-"`
	Property param.Opt[string] `query:"property,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [MessageGetParams]'s query parameters as `url.Values`.
func (r MessageGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type MessageGetOriginalContentParams struct {
	ThreadID int64             `path:"threadId,required" json:"-"`
	Property param.Opt[string] `query:"property,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [MessageGetOriginalContentParams]'s query parameters as
// `url.Values`.
func (r MessageGetOriginalContentParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
