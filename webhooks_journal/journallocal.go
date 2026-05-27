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

// JournalLocalService contains methods and other services that help with
// interacting with the hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewJournalLocalService] method instead.
type JournalLocalService struct {
	options []option.RequestOption
	Batch   JournalLocalBatchService
}

// NewJournalLocalService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewJournalLocalService(opts ...option.RequestOption) (r JournalLocalService) {
	r = JournalLocalService{}
	r.options = opts
	r.Batch = NewJournalLocalBatchService(opts...)
	return
}

// Retrieve the earliest webhook journal entries for the specified portal. This
// endpoint can be used to access the oldest records available in the webhook
// journal, which may be useful for auditing or historical analysis.
func (r *JournalLocalService) GetEarliest(ctx context.Context, query JournalLocalGetEarliestParams, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "webhooks-journal/journal-local/2026-03/earliest"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Retrieve the latest entries from the webhooks journal for the specified portal.
// This endpoint is useful for accessing the most recent webhook events and their
// statuses, allowing you to monitor and debug webhook activity effectively.
func (r *JournalLocalService) GetLatest(ctx context.Context, query JournalLocalGetLatestParams, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "webhooks-journal/journal/2026-03/latest"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Retrieve the next set of webhook journal entries starting from a specified
// offset. This endpoint is useful for paginating through large sets of webhook
// data, allowing you to continue from where a previous request left off.
func (r *JournalLocalService) GetNextFromOffset(ctx context.Context, offset string, query JournalLocalGetNextFromOffsetParams, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if offset == "" {
		err = errors.New("missing required offset parameter")
		return nil, err
	}
	path := fmt.Sprintf("webhooks-journal/journal-local/2026-03/offset/%s/next", url.PathEscape(offset))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Retrieve the status of a specific webhook journal entry using its unique status
// ID. This endpoint is useful for monitoring the progress or outcome of webhook
// journal entries, allowing you to check if an entry is pending, in progress,
// completed, failed, or expired.
func (r *JournalLocalService) GetStatus(ctx context.Context, statusID string, opts ...option.RequestOption) (res *shared.SnapshotStatusResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if statusID == "" {
		err = errors.New("missing required statusId parameter")
		return nil, err
	}
	path := fmt.Sprintf("webhooks-journal/journal-local/2026-03/status/%s", statusID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type JournalLocalGetEarliestParams struct {
	// The ID of the portal for which to retrieve the earliest webhook journal entries.
	// This parameter is optional and should be an integer.
	InstallPortalID param.Opt[int64] `query:"installPortalId,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [JournalLocalGetEarliestParams]'s query parameters as
// `url.Values`.
func (r JournalLocalGetEarliestParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type JournalLocalGetLatestParams struct {
	// The unique identifier of the portal installation for which to retrieve the
	// latest journal entries. This parameter is optional and should be an integer.
	InstallPortalID param.Opt[int64] `query:"installPortalId,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [JournalLocalGetLatestParams]'s query parameters as
// `url.Values`.
func (r JournalLocalGetLatestParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type JournalLocalGetNextFromOffsetParams struct {
	// The ID of the portal installation to filter the webhook journal entries. This is
	// an integer value.
	InstallPortalID param.Opt[int64] `query:"installPortalId,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [JournalLocalGetNextFromOffsetParams]'s query parameters as
// `url.Values`.
func (r JournalLocalGetNextFromOffsetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
