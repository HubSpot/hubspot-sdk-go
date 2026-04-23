// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package automation

import (
	"context"
	"errors"
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

// ActionRevisionService contains methods and other services that help with
// interacting with the hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewActionRevisionService] method instead.
type ActionRevisionService struct {
	options []option.RequestOption
}

// NewActionRevisionService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewActionRevisionService(opts ...option.RequestOption) (r ActionRevisionService) {
	r = ActionRevisionService{}
	r.options = opts
	return
}

// Retrieve the versions of a definition by ID.
func (r *ActionRevisionService) List(ctx context.Context, definitionID string, params ActionRevisionListParams, opts ...option.RequestOption) (res *pagination.Page[PublicActionRevision], err error) {
	var raw *http.Response
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if definitionID == "" {
		err = errors.New("missing required definitionId parameter")
		return nil, err
	}
	path := fmt.Sprintf("automation/actions/2026-03/%v/%s/revisions", params.AppID, url.PathEscape(definitionID))
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, params, &res, opts...)
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

// Retrieve the versions of a definition by ID.
func (r *ActionRevisionService) ListAutoPaging(ctx context.Context, definitionID string, params ActionRevisionListParams, opts ...option.RequestOption) *pagination.PageAutoPager[PublicActionRevision] {
	return pagination.NewPageAutoPager(r.List(ctx, definitionID, params, opts...))
}

// Retrieve a specific revision of a definition by revision ID.
func (r *ActionRevisionService) Get(ctx context.Context, revisionID string, query ActionRevisionGetParams, opts ...option.RequestOption) (res *PublicActionRevision, err error) {
	opts = slices.Concat(r.options, opts)
	if query.DefinitionID == "" {
		err = errors.New("missing required definitionId parameter")
		return nil, err
	}
	if revisionID == "" {
		err = errors.New("missing required revisionId parameter")
		return nil, err
	}
	path := fmt.Sprintf("automation/actions/2026-03/%v/%s/revisions/%s", query.AppID, url.PathEscape(query.DefinitionID), url.PathEscape(revisionID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type ActionRevisionListParams struct {
	AppID int64 `path:"appId" api:"required" json:"-"`
	// The paging cursor token of the last successfully read resource will be returned
	// as the `paging.next.after` JSON property of a paged response containing more
	// results.
	After param.Opt[string] `query:"after,omitzero" json:"-"`
	// The maximum number of results to display per page.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ActionRevisionListParams]'s query parameters as
// `url.Values`.
func (r ActionRevisionListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ActionRevisionGetParams struct {
	AppID        int64  `path:"appId" api:"required" json:"-"`
	DefinitionID string `path:"definitionId" api:"required" json:"-"`
	paramObj
}
