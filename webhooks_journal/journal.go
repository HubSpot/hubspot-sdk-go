// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package webhooks_journal

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/HubSpot/hubspot-sdk-go/internal/apiquery"
	"github.com/HubSpot/hubspot-sdk-go/internal/requestconfig"
	"github.com/HubSpot/hubspot-sdk-go/option"
	"github.com/HubSpot/hubspot-sdk-go/packages/param"
	"github.com/HubSpot/hubspot-sdk-go/shared"
)

// JournalService contains methods and other services that help with interacting
// with the hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewJournalService] method instead.
type JournalService struct {
	options []option.RequestOption
	Batch   JournalBatchService
}

// NewJournalService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewJournalService(opts ...option.RequestOption) (r JournalService) {
	r = JournalService{}
	r.options = opts
	r.Batch = NewJournalBatchService(opts...)
	return
}

// Retrieve the earliest entry from the webhooks journal for the specified portal.
// This endpoint is useful for accessing the first recorded webhook event in the
// journal, which can be helpful for auditing or debugging purposes.
func (r *JournalService) GetEarliest(ctx context.Context, query JournalGetEarliestParams, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "webhooks-journal/journal/2026-03/earliest"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Retrieve the next set of entries from the webhooks journal starting from a
// specified offset. This endpoint is useful for paginating through journal entries
// to process or analyze webhook events sequentially.
func (r *JournalService) GetNextFromOffset(ctx context.Context, offset string, query JournalGetNextFromOffsetParams, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if offset == "" {
		err = errors.New("missing required offset parameter")
		return nil, err
	}
	path := fmt.Sprintf("webhooks-journal/journal/2026-03/offset/%s/next", url.PathEscape(offset))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Retrieve the status of a specific webhook journal entry using its unique status
// ID. This endpoint provides detailed information about the status, including
// whether it is pending, in progress, completed, failed, or expired. It is useful
// for monitoring and managing the state of webhook journal entries.
func (r *JournalService) GetStatus(ctx context.Context, statusID string, opts ...option.RequestOption) (res *shared.SnapshotStatusResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if statusID == "" {
		err = errors.New("missing required statusId parameter")
		return nil, err
	}
	path := fmt.Sprintf("webhooks-journal/journal/2026-03/status/%s", statusID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type JournalGetEarliestParams struct {
	// The ID of the portal installation to filter the journal entries by. This is an
	// integer value.
	InstallPortalID param.Opt[int64] `query:"installPortalId,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [JournalGetEarliestParams]'s query parameters as
// `url.Values`.
func (r JournalGetEarliestParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type JournalGetNextFromOffsetParams struct {
	// The ID of the portal where the webhooks are installed. This is an integer value.
	InstallPortalID param.Opt[int64] `query:"installPortalId,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [JournalGetNextFromOffsetParams]'s query parameters as
// `url.Values`.
func (r JournalGetNextFromOffsetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
