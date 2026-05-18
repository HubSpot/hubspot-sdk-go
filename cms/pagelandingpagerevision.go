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

// PageLandingPageRevisionService contains methods and other services that help
// with interacting with the hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPageLandingPageRevisionService] method instead.
type PageLandingPageRevisionService struct {
	options []option.RequestOption
}

// NewPageLandingPageRevisionService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewPageLandingPageRevisionService(opts ...option.RequestOption) (r PageLandingPageRevisionService) {
	r = PageLandingPageRevisionService{}
	r.options = opts
	return
}

// Retrieve a previous version of a landing page, specified by page ID and revision
// ID.
func (r *PageLandingPageRevisionService) GetLandingPageRevision(ctx context.Context, revisionID string, query PageLandingPageRevisionGetLandingPageRevisionParams, opts ...option.RequestOption) (res *PageVersion, err error) {
	opts = slices.Concat(r.options, opts)
	if query.ObjectID == "" {
		err = errors.New("missing required objectId parameter")
		return nil, err
	}
	if revisionID == "" {
		err = errors.New("missing required revisionId parameter")
		return nil, err
	}
	path := fmt.Sprintf("cms/pages/2026-03/landing-pages/%s/revisions/%s", url.PathEscape(query.ObjectID), url.PathEscape(revisionID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Retrieve all the previous versions of a landing page, specified by page ID.
func (r *PageLandingPageRevisionService) ListLandingPageRevisions(ctx context.Context, objectID string, query PageLandingPageRevisionListLandingPageRevisionsParams, opts ...option.RequestOption) (res *pagination.Page[PageVersion], err error) {
	var raw *http.Response
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if objectID == "" {
		err = errors.New("missing required objectId parameter")
		return nil, err
	}
	path := fmt.Sprintf("cms/pages/2026-03/landing-pages/%s/revisions", url.PathEscape(objectID))
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

// Retrieve all the previous versions of a landing page, specified by page ID.
func (r *PageLandingPageRevisionService) ListLandingPageRevisionsAutoPaging(ctx context.Context, objectID string, query PageLandingPageRevisionListLandingPageRevisionsParams, opts ...option.RequestOption) *pagination.PageAutoPager[PageVersion] {
	return pagination.NewPageAutoPager(r.ListLandingPageRevisions(ctx, objectID, query, opts...))
}

// Restores a previous version of a landing page, specified by page ID and revision
// ID.
func (r *PageLandingPageRevisionService) RestoreLandingPageRevision(ctx context.Context, revisionID string, body PageLandingPageRevisionRestoreLandingPageRevisionParams, opts ...option.RequestOption) (res *PagesPage, err error) {
	opts = slices.Concat(r.options, opts)
	if body.ObjectID == "" {
		err = errors.New("missing required objectId parameter")
		return nil, err
	}
	if revisionID == "" {
		err = errors.New("missing required revisionId parameter")
		return nil, err
	}
	path := fmt.Sprintf("cms/pages/2026-03/landing-pages/%s/revisions/%s/restore", url.PathEscape(body.ObjectID), url.PathEscape(revisionID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

// Specify a previous version of a landing page to set as the page draft.
func (r *PageLandingPageRevisionService) RestoreLandingPageRevisionToDraft(ctx context.Context, revisionID int64, body PageLandingPageRevisionRestoreLandingPageRevisionToDraftParams, opts ...option.RequestOption) (res *PagesPage, err error) {
	opts = slices.Concat(r.options, opts)
	if body.ObjectID == "" {
		err = errors.New("missing required objectId parameter")
		return nil, err
	}
	path := fmt.Sprintf("cms/pages/2026-03/landing-pages/%s/revisions/%v/restore-to-draft", url.PathEscape(body.ObjectID), revisionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

type PageLandingPageRevisionGetLandingPageRevisionParams struct {
	ObjectID string `path:"objectId" api:"required" json:"-"`
	paramObj
}

type PageLandingPageRevisionListLandingPageRevisionsParams struct {
	// The paging cursor token of the last successfully read resource will be returned
	// as the `paging.next.after` JSON property of a paged response containing more
	// results.
	After  param.Opt[string] `query:"after,omitzero" json:"-"`
	Before param.Opt[string] `query:"before,omitzero" json:"-"`
	// The maximum number of results to display per page.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [PageLandingPageRevisionListLandingPageRevisionsParams]'s
// query parameters as `url.Values`.
func (r PageLandingPageRevisionListLandingPageRevisionsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type PageLandingPageRevisionRestoreLandingPageRevisionParams struct {
	ObjectID string `path:"objectId" api:"required" json:"-"`
	paramObj
}

type PageLandingPageRevisionRestoreLandingPageRevisionToDraftParams struct {
	ObjectID string `path:"objectId" api:"required" json:"-"`
	paramObj
}
