// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cms

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/HubSpot/hubspot-sdk-go/internal/apijson"
	"github.com/HubSpot/hubspot-sdk-go/internal/apiquery"
	shimjson "github.com/HubSpot/hubspot-sdk-go/internal/encoding/json"
	"github.com/HubSpot/hubspot-sdk-go/internal/requestconfig"
	"github.com/HubSpot/hubspot-sdk-go/option"
	"github.com/HubSpot/hubspot-sdk-go/packages/pagination"
	"github.com/HubSpot/hubspot-sdk-go/packages/param"
	"github.com/HubSpot/hubspot-sdk-go/shared"
)

// PageFolderService contains methods and other services that help with interacting
// with the hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPageFolderService] method instead.
type PageFolderService struct {
	options []option.RequestOption
}

// NewPageFolderService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewPageFolderService(opts ...option.RequestOption) (r PageFolderService) {
	r = PageFolderService{}
	r.options = opts
	return
}

// Create a new folder for landing pages.
func (r *PageFolderService) New(ctx context.Context, body PageFolderNewParams, opts ...option.RequestOption) (res *ContentFolder, err error) {
	opts = slices.Concat(r.options, opts)
	path := "cms/pages/2026-03/landing-pages/folders"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Partially update a landing page folder, specified by the folder ID. You only
// need to specify the details values that you are modifying.
func (r *PageFolderService) Update(ctx context.Context, objectID string, params PageFolderUpdateParams, opts ...option.RequestOption) (res *ContentFolder, err error) {
	opts = slices.Concat(r.options, opts)
	if objectID == "" {
		err = errors.New("missing required objectId parameter")
		return nil, err
	}
	path := fmt.Sprintf("cms/pages/2026-03/landing-pages/folders/%s", url.PathEscape(objectID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return res, err
}

// Get the list of Landing Page Folders. Supports paging and filtering. This method
// would be useful for an integration that examined these models and used an
// external service to suggest edits.
func (r *PageFolderService) List(ctx context.Context, query PageFolderListParams, opts ...option.RequestOption) (res *pagination.Page[ContentFolder], err error) {
	var raw *http.Response
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "cms/pages/2026-03/landing-pages/folders"
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

// Get the list of Landing Page Folders. Supports paging and filtering. This method
// would be useful for an integration that examined these models and used an
// external service to suggest edits.
func (r *PageFolderService) ListAutoPaging(ctx context.Context, query PageFolderListParams, opts ...option.RequestOption) *pagination.PageAutoPager[ContentFolder] {
	return pagination.NewPageAutoPager(r.List(ctx, query, opts...))
}

// Delete a landing page folder, specified by its ID.
func (r *PageFolderService) Delete(ctx context.Context, objectID string, body PageFolderDeleteParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if objectID == "" {
		err = errors.New("missing required objectId parameter")
		return err
	}
	path := fmt.Sprintf("cms/pages/2026-03/landing-pages/folders/%s", url.PathEscape(objectID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, nil, opts...)
	return err
}

// Retrieve a batch of landing page folders as identified in the request body.
func (r *PageFolderService) BatchGet(ctx context.Context, params PageFolderBatchGetParams, opts ...option.RequestOption) (res *BatchResponseContentFolder, err error) {
	opts = slices.Concat(r.options, opts)
	path := "cms/pages/2026-03/landing-pages/folders/batch/read"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Retrieve a landing page folder, specified by its ID.
func (r *PageFolderService) Get(ctx context.Context, objectID string, query PageFolderGetParams, opts ...option.RequestOption) (res *ContentFolder, err error) {
	opts = slices.Concat(r.options, opts)
	if objectID == "" {
		err = errors.New("missing required objectId parameter")
		return nil, err
	}
	path := fmt.Sprintf("cms/pages/2026-03/landing-pages/folders/%s", url.PathEscape(objectID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Retrieve a previous version of a folder, specified by the folder ID and revision
// ID.
func (r *PageFolderService) GetRevision(ctx context.Context, revisionID string, query PageFolderGetRevisionParams, opts ...option.RequestOption) (res *ContentFolderVersion, err error) {
	opts = slices.Concat(r.options, opts)
	if query.ObjectID == "" {
		err = errors.New("missing required objectId parameter")
		return nil, err
	}
	if revisionID == "" {
		err = errors.New("missing required revisionId parameter")
		return nil, err
	}
	path := fmt.Sprintf("cms/pages/2026-03/landing-pages/folders/%s/revisions/%s", url.PathEscape(query.ObjectID), url.PathEscape(revisionID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Retrieves all the previous versions of a landing page folder.
func (r *PageFolderService) ListRevisions(ctx context.Context, objectID string, query PageFolderListRevisionsParams, opts ...option.RequestOption) (res *pagination.Page[ContentFolderVersion], err error) {
	var raw *http.Response
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if objectID == "" {
		err = errors.New("missing required objectId parameter")
		return nil, err
	}
	path := fmt.Sprintf("cms/pages/2026-03/landing-pages/folders/%s/revisions", url.PathEscape(objectID))
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

// Retrieves all the previous versions of a landing page folder.
func (r *PageFolderService) ListRevisionsAutoPaging(ctx context.Context, objectID string, query PageFolderListRevisionsParams, opts ...option.RequestOption) *pagination.PageAutoPager[ContentFolderVersion] {
	return pagination.NewPageAutoPager(r.ListRevisions(ctx, objectID, query, opts...))
}

// Takes a specified version of a landing page folder and restores it.
func (r *PageFolderService) RestoreRevision(ctx context.Context, revisionID string, body PageFolderRestoreRevisionParams, opts ...option.RequestOption) (res *ContentFolder, err error) {
	opts = slices.Concat(r.options, opts)
	if body.ObjectID == "" {
		err = errors.New("missing required objectId parameter")
		return nil, err
	}
	if revisionID == "" {
		err = errors.New("missing required revisionId parameter")
		return nil, err
	}
	path := fmt.Sprintf("cms/pages/2026-03/landing-pages/folders/%s/revisions/%s/restore", url.PathEscape(body.ObjectID), url.PathEscape(revisionID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

type PageFolderNewParams struct {
	ContentFolder ContentFolderParam
	paramObj
}

func (r PageFolderNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.ContentFolder)
}
func (r *PageFolderNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PageFolderUpdateParams struct {
	ContentFolder ContentFolderParam
	// Whether to return only results that have been archived.
	Archived param.Opt[bool] `query:"archived,omitzero" json:"-"`
	paramObj
}

func (r PageFolderUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.ContentFolder)
}
func (r *PageFolderUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [PageFolderUpdateParams]'s query parameters as `url.Values`.
func (r PageFolderUpdateParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type PageFolderListParams struct {
	// The paging cursor token of the last successfully read resource will be returned
	// as the `paging.next.after` JSON property of a paged response containing more
	// results.
	After param.Opt[string] `query:"after,omitzero" json:"-"`
	// Whether to return only results that have been archived.
	Archived param.Opt[bool] `query:"archived,omitzero" json:"-"`
	// Filter folders created after the specified date and time.
	CreatedAfter param.Opt[time.Time] `query:"createdAfter,omitzero" format:"date-time" json:"-"`
	// Filter folders by their exact creation date and time.
	CreatedAt     param.Opt[time.Time] `query:"createdAt,omitzero" format:"date-time" json:"-"`
	CreatedBefore param.Opt[time.Time] `query:"createdBefore,omitzero" format:"date-time" json:"-"`
	// The maximum number of results to display per page.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Specify a property to include in the response.
	Property      param.Opt[string]    `query:"property,omitzero" json:"-"`
	UpdatedAfter  param.Opt[time.Time] `query:"updatedAfter,omitzero" format:"date-time" json:"-"`
	UpdatedAt     param.Opt[time.Time] `query:"updatedAt,omitzero" format:"date-time" json:"-"`
	UpdatedBefore param.Opt[time.Time] `query:"updatedBefore,omitzero" format:"date-time" json:"-"`
	Sort          []string             `query:"sort,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [PageFolderListParams]'s query parameters as `url.Values`.
func (r PageFolderListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type PageFolderDeleteParams struct {
	// Whether to return only results that have been archived.
	Archived param.Opt[bool] `query:"archived,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [PageFolderDeleteParams]'s query parameters as `url.Values`.
func (r PageFolderDeleteParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type PageFolderBatchGetParams struct {
	BatchInputString shared.BatchInputStringParam
	// Whether to return only results that have been archived.
	Archived param.Opt[bool] `query:"archived,omitzero" json:"-"`
	paramObj
}

func (r PageFolderBatchGetParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.BatchInputString)
}
func (r *PageFolderBatchGetParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [PageFolderBatchGetParams]'s query parameters as
// `url.Values`.
func (r PageFolderBatchGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type PageFolderGetParams struct {
	// Whether to return only results that have been archived.
	Archived param.Opt[bool]   `query:"archived,omitzero" json:"-"`
	Property param.Opt[string] `query:"property,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [PageFolderGetParams]'s query parameters as `url.Values`.
func (r PageFolderGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type PageFolderGetRevisionParams struct {
	ObjectID string `path:"objectId" api:"required" json:"-"`
	paramObj
}

type PageFolderListRevisionsParams struct {
	// The paging cursor token of the last successfully read resource will be returned
	// as the `paging.next.after` JSON property of a paged response containing more
	// results.
	After  param.Opt[string] `query:"after,omitzero" json:"-"`
	Before param.Opt[string] `query:"before,omitzero" json:"-"`
	// The maximum number of results to display per page.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [PageFolderListRevisionsParams]'s query parameters as
// `url.Values`.
func (r PageFolderListRevisionsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type PageFolderRestoreRevisionParams struct {
	ObjectID string `path:"objectId" api:"required" json:"-"`
	paramObj
}
