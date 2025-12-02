// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package conversations

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/stainless-sdks/hubspot-sdk-go/internal/apiquery"
	shimjson "github.com/stainless-sdks/hubspot-sdk-go/internal/encoding/json"
	"github.com/stainless-sdks/hubspot-sdk-go/internal/requestconfig"
	"github.com/stainless-sdks/hubspot-sdk-go/option"
	"github.com/stainless-sdks/hubspot-sdk-go/packages/pagination"
	"github.com/stainless-sdks/hubspot-sdk-go/packages/param"
)

// ThreadService contains methods and other services that help with interacting
// with the hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewThreadService] method instead.
type ThreadService struct {
	Options []option.RequestOption
}

// NewThreadService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewThreadService(opts ...option.RequestOption) (r ThreadService) {
	r = ThreadService{}
	r.Options = opts
	return
}

func (r *ThreadService) Update(ctx context.Context, threadID int64, params ThreadUpdateParams, opts ...option.RequestOption) (res *PublicThread, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("conversations/v3/conversations/threads/%v", threadID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return
}

func (r *ThreadService) List(ctx context.Context, query ThreadListParams, opts ...option.RequestOption) (res *pagination.Page[PublicThread], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "conversations/v3/conversations/threads"
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

func (r *ThreadService) ListAutoPaging(ctx context.Context, query ThreadListParams, opts ...option.RequestOption) *pagination.PageAutoPager[PublicThread] {
	return pagination.NewPageAutoPager(r.List(ctx, query, opts...))
}

func (r *ThreadService) Delete(ctx context.Context, threadID int64, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := fmt.Sprintf("conversations/v3/conversations/threads/%v", threadID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

func (r *ThreadService) Get(ctx context.Context, threadID int64, query ThreadGetParams, opts ...option.RequestOption) (res *PublicThread, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("conversations/v3/conversations/threads/%v", threadID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type ThreadUpdateParams struct {
	PublicThreadUpdateRequest PublicThreadUpdateRequestParam
	Archived                  param.Opt[bool] `query:"archived,omitzero" json:"-"`
	paramObj
}

func (r ThreadUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.PublicThreadUpdateRequest)
}
func (r *ThreadUpdateParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.PublicThreadUpdateRequest)
}

// URLQuery serializes [ThreadUpdateParams]'s query parameters as `url.Values`.
func (r ThreadUpdateParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ThreadListParams struct {
	After                       param.Opt[string]    `query:"after,omitzero" json:"-"`
	Archived                    param.Opt[bool]      `query:"archived,omitzero" json:"-"`
	AssociatedContactID         param.Opt[int64]     `query:"associatedContactId,omitzero" json:"-"`
	LatestMessageTimestampAfter param.Opt[time.Time] `query:"latestMessageTimestampAfter,omitzero" format:"date-time" json:"-"`
	Limit                       param.Opt[int64]     `query:"limit,omitzero" json:"-"`
	Property                    param.Opt[string]    `query:"property,omitzero" json:"-"`
	ThreadStatus                param.Opt[string]    `query:"threadStatus,omitzero" json:"-"`
	// Any of "TICKET".
	Association []string `query:"association,omitzero" json:"-"`
	InboxID     []int64  `query:"inboxId,omitzero" json:"-"`
	Sort        []string `query:"sort,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ThreadListParams]'s query parameters as `url.Values`.
func (r ThreadListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ThreadGetParams struct {
	Archived param.Opt[bool]   `query:"archived,omitzero" json:"-"`
	Property param.Opt[string] `query:"property,omitzero" json:"-"`
	// Any of "TICKET".
	Association []string `query:"association,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ThreadGetParams]'s query parameters as `url.Values`.
func (r ThreadGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
