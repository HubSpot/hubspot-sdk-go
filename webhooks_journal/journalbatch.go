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

// JournalBatchService contains methods and other services that help with
// interacting with the hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewJournalBatchService] method instead.
type JournalBatchService struct {
	options []option.RequestOption
}

// NewJournalBatchService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewJournalBatchService(opts ...option.RequestOption) (r JournalBatchService) {
	r = JournalBatchService{}
	r.options = opts
	return
}

// Execute a batch read operation on the webhooks journal for the specified date,
// 2026-03. This endpoint allows you to retrieve multiple entries from the webhooks
// journal in a single request, which can be useful for processing large amounts of
// data efficiently. Ensure that the request body is provided in the required
// format.
func (r *JournalBatchService) Get(ctx context.Context, params JournalBatchGetParams, opts ...option.RequestOption) (res *shared.BatchResponseJournalFetchResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "webhooks-journal/journal/2026-03/batch/read"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Retrieve the earliest batch of webhook journal entries for a specified count.
// This endpoint is useful for accessing historical webhook data in batches,
// allowing you to process or analyze older entries. The number of entries
// retrieved is determined by the count parameter.
func (r *JournalBatchService) GetEarliest(ctx context.Context, count int64, query JournalBatchGetEarliestParams, opts ...option.RequestOption) (res *shared.BatchResponseJournalFetchResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := fmt.Sprintf("webhooks-journal/journal/2026-03/batch/earliest/%v", count)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Retrieve a batch of webhook journal entries starting from a specified offset.
// This endpoint allows you to fetch a defined number of entries, which can be
// useful for processing large datasets in manageable chunks.
func (r *JournalBatchService) GetFromOffset(ctx context.Context, count int64, params JournalBatchGetFromOffsetParams, opts ...option.RequestOption) (res *shared.BatchResponseJournalFetchResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if params.Offset == "" {
		err = errors.New("missing required offset parameter")
		return nil, err
	}
	path := fmt.Sprintf("webhooks-journal/journal/2026-03/batch/%s/next/%v", url.PathEscape(params.Offset), count)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// Retrieve the latest batch of webhook journal entries up to the specified count.
// This endpoint is useful for fetching recent webhook data for analysis or
// processing. The count parameter determines the maximum number of entries to
// return.
func (r *JournalBatchService) GetLatest(ctx context.Context, count int64, query JournalBatchGetLatestParams, opts ...option.RequestOption) (res *shared.BatchResponseJournalFetchResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := fmt.Sprintf("webhooks-journal/journal/2026-03/batch/latest/%v", count)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

type JournalBatchGetParams struct {
	BatchInputString shared.BatchInputStringParam
	// An integer representing the ID of the portal installation for which the webhooks
	// journal data should be retrieved.
	InstallPortalID param.Opt[int64] `query:"installPortalId,omitzero" json:"-"`
	paramObj
}

func (r JournalBatchGetParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.BatchInputString)
}
func (r *JournalBatchGetParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [JournalBatchGetParams]'s query parameters as `url.Values`.
func (r JournalBatchGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type JournalBatchGetEarliestParams struct {
	// The ID of the portal installation. This is an integer value that specifies which
	// portal's data to access.
	InstallPortalID param.Opt[int64] `query:"installPortalId,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [JournalBatchGetEarliestParams]'s query parameters as
// `url.Values`.
func (r JournalBatchGetEarliestParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type JournalBatchGetFromOffsetParams struct {
	Offset string `path:"offset" api:"required" json:"-"`
	// The ID of the portal installation. This is an integer value.
	InstallPortalID param.Opt[int64] `query:"installPortalId,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [JournalBatchGetFromOffsetParams]'s query parameters as
// `url.Values`.
func (r JournalBatchGetFromOffsetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type JournalBatchGetLatestParams struct {
	// The ID of the portal installation. This is an integer value used to specify the
	// portal context for the request.
	InstallPortalID param.Opt[int64] `query:"installPortalId,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [JournalBatchGetLatestParams]'s query parameters as
// `url.Values`.
func (r JournalBatchGetLatestParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
