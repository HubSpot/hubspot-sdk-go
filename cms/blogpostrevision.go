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
	"github.com/HubSpot/hubspot-sdk-go/packages/param"
)

// BlogPostRevisionService contains methods and other services that help with
// interacting with the hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBlogPostRevisionService] method instead.
type BlogPostRevisionService struct {
	options []option.RequestOption
}

// NewBlogPostRevisionService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewBlogPostRevisionService(opts ...option.RequestOption) (r BlogPostRevisionService) {
	r = BlogPostRevisionService{}
	r.options = opts
	return
}

// Retrieve a previous version of a blog post.
func (r *BlogPostRevisionService) GetPreviousVersion(ctx context.Context, revisionID string, query BlogPostRevisionGetPreviousVersionParams, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if query.ObjectID == "" {
		err = errors.New("missing required objectId parameter")
		return nil, err
	}
	if revisionID == "" {
		err = errors.New("missing required revisionId parameter")
		return nil, err
	}
	path := fmt.Sprintf("cms/blogs/2026-03/posts/%s/revisions/%s", url.PathEscape(query.ObjectID), url.PathEscape(revisionID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Retrieve all the previous versions of a blog post.
func (r *BlogPostRevisionService) GetPreviousVersions(ctx context.Context, objectID string, query BlogPostRevisionGetPreviousVersionsParams, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if objectID == "" {
		err = errors.New("missing required objectId parameter")
		return nil, err
	}
	path := fmt.Sprintf("cms/blogs/2026-03/posts/%s/revisions", url.PathEscape(objectID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Restores a blog post to one of its previous versions.
func (r *BlogPostRevisionService) RestorePreviousVersion(ctx context.Context, revisionID string, body BlogPostRevisionRestorePreviousVersionParams, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if body.ObjectID == "" {
		err = errors.New("missing required objectId parameter")
		return nil, err
	}
	if revisionID == "" {
		err = errors.New("missing required revisionId parameter")
		return nil, err
	}
	path := fmt.Sprintf("cms/blogs/2026-03/posts/%s/revisions/%s/restore", url.PathEscape(body.ObjectID), url.PathEscape(revisionID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

// Takes a specified version of a blog post, sets it as the new draft version of
// the blog post.
func (r *BlogPostRevisionService) RestorePreviousVersionToDraft(ctx context.Context, revisionID int64, body BlogPostRevisionRestorePreviousVersionToDraftParams, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if body.ObjectID == "" {
		err = errors.New("missing required objectId parameter")
		return nil, err
	}
	path := fmt.Sprintf("cms/blogs/2026-03/posts/%s/revisions/%v/restore-to-draft", url.PathEscape(body.ObjectID), revisionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

type BlogPostRevisionGetPreviousVersionParams struct {
	ObjectID string `path:"objectId" api:"required" json:"-"`
	paramObj
}

type BlogPostRevisionGetPreviousVersionsParams struct {
	// The paging cursor token of the last successfully read resource will be returned
	// as the `paging.next.after` JSON property of a paged response containing more
	// results.
	After  param.Opt[string] `query:"after,omitzero" json:"-"`
	Before param.Opt[string] `query:"before,omitzero" json:"-"`
	// The maximum number of results to display per page.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [BlogPostRevisionGetPreviousVersionsParams]'s query
// parameters as `url.Values`.
func (r BlogPostRevisionGetPreviousVersionsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type BlogPostRevisionRestorePreviousVersionParams struct {
	ObjectID string `path:"objectId" api:"required" json:"-"`
	paramObj
}

type BlogPostRevisionRestorePreviousVersionToDraftParams struct {
	ObjectID string `path:"objectId" api:"required" json:"-"`
	paramObj
}
