// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package webhooks_journal

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
	"github.com/HubSpot/hubspot-sdk-go/packages/param"
	"github.com/HubSpot/hubspot-sdk-go/shared"
)

// JournalLocalBatchService contains methods and other services that help with
// interacting with the hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewJournalLocalBatchService] method instead.
type JournalLocalBatchService struct {
	options []option.RequestOption
}

// NewJournalLocalBatchService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewJournalLocalBatchService(opts ...option.RequestOption) (r JournalLocalBatchService) {
	r = JournalLocalBatchService{}
	r.options = opts
	return
}

// Execute a batch read operation on the webhooks journal. This endpoint allows you
// to retrieve a batch of webhook journal entries by providing the necessary input
// data. It is useful for processing multiple records in a single request,
// streamlining data retrieval tasks.
func (r *JournalLocalBatchService) Get(ctx context.Context, params JournalLocalBatchGetParams, opts ...option.RequestOption) (res *shared.BatchResponseJournalFetchResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "webhooks-journal/journal-local/2026-03/batch/read"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Retrieve the earliest batch of webhook journal entries. This endpoint is useful
// for accessing the oldest available data in the webhook journal, allowing users
// to process or analyze historical webhook events. The number of entries to fetch
// is specified by the 'count' path parameter.
func (r *JournalLocalBatchService) GetEarliest(ctx context.Context, count int64, query JournalLocalBatchGetEarliestParams, opts ...option.RequestOption) (res *shared.BatchResponseJournalFetchResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := fmt.Sprintf("webhooks-journal/journal-local/2026-03/batch/earliest/%v", count)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Retrieve a batch of webhook journal entries starting from a specified offset.
// This endpoint is useful for paginating through large sets of webhook data. The
// number of entries returned is determined by the 'count' parameter.
func (r *JournalLocalBatchService) GetFromOffset(ctx context.Context, count int64, params JournalLocalBatchGetFromOffsetParams, opts ...option.RequestOption) (res *shared.BatchResponseJournalFetchResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if params.Offset == "" {
		err = errors.New("missing required offset parameter")
		return nil, err
	}
	path := fmt.Sprintf("webhooks-journal/journal-local/2026-03/batch/%s/next/%v", url.PathEscape(params.Offset), count)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

type JournalLocalBatchGetParams struct {
	BatchInputString shared.BatchInputStringParam
	// The ID of the portal where the webhooks are installed. This parameter is
	// optional and is used to specify the target portal for the operation.
	InstallPortalID param.Opt[int64] `query:"installPortalId,omitzero" json:"-"`
	paramObj
}

func (r JournalLocalBatchGetParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.BatchInputString)
}
func (r *JournalLocalBatchGetParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [JournalLocalBatchGetParams]'s query parameters as
// `url.Values`.
func (r JournalLocalBatchGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type JournalLocalBatchGetEarliestParams struct {
	// The ID of the portal installation to filter the webhook journal entries. This is
	// an optional integer parameter.
	InstallPortalID param.Opt[int64] `query:"installPortalId,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [JournalLocalBatchGetEarliestParams]'s query parameters as
// `url.Values`.
func (r JournalLocalBatchGetEarliestParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type JournalLocalBatchGetFromOffsetParams struct {
	Offset string `path:"offset" api:"required" json:"-"`
	// The ID of the portal where the webhooks are installed. This is an optional
	// parameter.
	InstallPortalID param.Opt[int64] `query:"installPortalId,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [JournalLocalBatchGetFromOffsetParams]'s query parameters as
// `url.Values`.
func (r JournalLocalBatchGetFromOffsetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
