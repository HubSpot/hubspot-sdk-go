// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cms

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
	"github.com/HubSpot/hubspot-sdk-go/packages/pagination"
	"github.com/HubSpot/hubspot-sdk-go/packages/param"
)

// PageSitePageRevisionService contains methods and other services that help with
// interacting with the hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPageSitePageRevisionService] method instead.
type PageSitePageRevisionService struct {
	options []option.RequestOption
}

// NewPageSitePageRevisionService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewPageSitePageRevisionService(opts ...option.RequestOption) (r PageSitePageRevisionService) {
	r = PageSitePageRevisionService{}
	r.options = opts
	return
}

// Retrieve a previous version of a website page by the revision ID.
func (r *PageSitePageRevisionService) GetSitePageRevision(ctx context.Context, revisionID string, query PageSitePageRevisionGetSitePageRevisionParams, opts ...option.RequestOption) (res *PageVersion, err error) {
	opts = slices.Concat(r.options, opts)
	if query.ObjectID == "" {
		err = errors.New("missing required objectId parameter")
		return nil, err
	}
	if revisionID == "" {
		err = errors.New("missing required revisionId parameter")
		return nil, err
	}
	path := fmt.Sprintf("cms/pages/2026-03/site-pages/%s/revisions/%s", url.PathEscape(query.ObjectID), url.PathEscape(revisionID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Retrieves all the previous versions of a website page, specified by page ID.
func (r *PageSitePageRevisionService) ListSitePageRevisions(ctx context.Context, objectID string, query PageSitePageRevisionListSitePageRevisionsParams, opts ...option.RequestOption) (res *pagination.Page[PageVersion], err error) {
	var raw *http.Response
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if objectID == "" {
		err = errors.New("missing required objectId parameter")
		return nil, err
	}
	path := fmt.Sprintf("cms/pages/2026-03/site-pages/%s/revisions", url.PathEscape(objectID))
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

// Retrieves all the previous versions of a website page, specified by page ID.
func (r *PageSitePageRevisionService) ListSitePageRevisionsAutoPaging(ctx context.Context, objectID string, query PageSitePageRevisionListSitePageRevisionsParams, opts ...option.RequestOption) *pagination.PageAutoPager[PageVersion] {
	return pagination.NewPageAutoPager(r.ListSitePageRevisions(ctx, objectID, query, opts...))
}

// Restores a website page to a previous version, specified by page ID and version
// ID.
func (r *PageSitePageRevisionService) RestoreSitePageRevision(ctx context.Context, revisionID string, body PageSitePageRevisionRestoreSitePageRevisionParams, opts ...option.RequestOption) (res *PagesPage, err error) {
	opts = slices.Concat(r.options, opts)
	if body.ObjectID == "" {
		err = errors.New("missing required objectId parameter")
		return nil, err
	}
	if revisionID == "" {
		err = errors.New("missing required revisionId parameter")
		return nil, err
	}
	path := fmt.Sprintf("cms/pages/2026-03/site-pages/%s/revisions/%s/restore", url.PathEscape(body.ObjectID), url.PathEscape(revisionID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

// Takes a specified version of a website page and sets it as the new draft version
// of the page.
func (r *PageSitePageRevisionService) RestoreSitePageRevisionToDraft(ctx context.Context, revisionID int64, body PageSitePageRevisionRestoreSitePageRevisionToDraftParams, opts ...option.RequestOption) (res *PagesPage, err error) {
	opts = slices.Concat(r.options, opts)
	if body.ObjectID == "" {
		err = errors.New("missing required objectId parameter")
		return nil, err
	}
	path := fmt.Sprintf("cms/pages/2026-03/site-pages/%s/revisions/%v/restore-to-draft", url.PathEscape(body.ObjectID), revisionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

type PageSitePageRevisionGetSitePageRevisionParams struct {
	ObjectID string `path:"objectId" api:"required" json:"-"`
	paramObj
}

type PageSitePageRevisionListSitePageRevisionsParams struct {
	// The paging cursor token of the last successfully read resource will be returned
	// as the `paging.next.after` JSON property of a paged response containing more
	// results.
	After  param.Opt[string] `query:"after,omitzero" json:"-"`
	Before param.Opt[string] `query:"before,omitzero" json:"-"`
	// The maximum number of results to display per page.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [PageSitePageRevisionListSitePageRevisionsParams]'s query
// parameters as `url.Values`.
func (r PageSitePageRevisionListSitePageRevisionsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type PageSitePageRevisionRestoreSitePageRevisionParams struct {
	ObjectID string `path:"objectId" api:"required" json:"-"`
	paramObj
}

type PageSitePageRevisionRestoreSitePageRevisionToDraftParams struct {
	ObjectID string `path:"objectId" api:"required" json:"-"`
	paramObj
}
