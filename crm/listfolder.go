// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package crm

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/stainless-sdks/hubspot-sdk-go/internal/apiquery"
	shimjson "github.com/stainless-sdks/hubspot-sdk-go/internal/encoding/json"
	"github.com/stainless-sdks/hubspot-sdk-go/internal/requestconfig"
	"github.com/stainless-sdks/hubspot-sdk-go/option"
	"github.com/stainless-sdks/hubspot-sdk-go/packages/param"
)

// ListFolderService contains methods and other services that help with interacting
// with the Hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewListFolderService] method instead.
type ListFolderService struct {
	Options []option.RequestOption
}

// NewListFolderService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewListFolderService(opts ...option.RequestOption) (r ListFolderService) {
	r = ListFolderService{}
	r.Options = opts
	return
}

// Creates a folder with the given information.
func (r *ListFolderService) New(ctx context.Context, body ListFolderNewParams, opts ...option.RequestOption) (res *ListFolderCreateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "crm/v3/lists/folders"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Deletes the folder with the given Id.
func (r *ListFolderService) Delete(ctx context.Context, folderID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if folderID == "" {
		err = errors.New("missing required folderId parameter")
		return
	}
	path := fmt.Sprintf("crm/v3/lists/folders/%s", folderID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Retrieves a folder and recursively includes all folders via the childNodes
// attribute. The child lists field will be empty in all child nodes. Only the
// folder retrieved will include the child lists in that folder.
func (r *ListFolderService) Get(ctx context.Context, query ListFolderGetParams, opts ...option.RequestOption) (res *ListFolderFetchResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "crm/v3/lists/folders"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// This moves the folder from its current location to a new location. It updates
// the parent of this folder to the new Id given.
func (r *ListFolderService) Move(ctx context.Context, newParentFolderID string, body ListFolderMoveParams, opts ...option.RequestOption) (res *ListFolderFetchResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if body.FolderID == "" {
		err = errors.New("missing required folderId parameter")
		return
	}
	if newParentFolderID == "" {
		err = errors.New("missing required newParentFolderId parameter")
		return
	}
	path := fmt.Sprintf("crm/v3/lists/folders/%s/move/%s", body.FolderID, newParentFolderID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, nil, &res, opts...)
	return
}

// Given a list and a folder, the list will be moved to that folder.
func (r *ListFolderService) MoveList(ctx context.Context, body ListFolderMoveListParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "crm/v3/lists/folders/move-list"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, nil, opts...)
	return
}

// Renames the given folderId with a new name.
func (r *ListFolderService) Rename(ctx context.Context, folderID string, body ListFolderRenameParams, opts ...option.RequestOption) (res *ListFolderFetchResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if folderID == "" {
		err = errors.New("missing required folderId parameter")
		return
	}
	path := fmt.Sprintf("crm/v3/lists/folders/%s/rename", folderID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

type ListFolderNewParams struct {
	ListFolderCreateRequest ListFolderCreateRequestParam
	paramObj
}

func (r ListFolderNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.ListFolderCreateRequest)
}
func (r *ListFolderNewParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.ListFolderCreateRequest)
}

type ListFolderGetParams struct {
	// The Id of the folder to retrieve.
	FolderID param.Opt[string] `query:"folderId,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ListFolderGetParams]'s query parameters as `url.Values`.
func (r ListFolderGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ListFolderMoveParams struct {
	FolderID string `path:"folderId,required" json:"-"`
	paramObj
}

type ListFolderMoveListParams struct {
	ListMoveRequest ListMoveRequestParam
	paramObj
}

func (r ListFolderMoveListParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.ListMoveRequest)
}
func (r *ListFolderMoveListParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.ListMoveRequest)
}

type ListFolderRenameParams struct {
	// The new name of the folder.
	NewFolderName param.Opt[string] `query:"newFolderName,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ListFolderRenameParams]'s query parameters as `url.Values`.
func (r ListFolderRenameParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
