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
	"github.com/stainless-sdks/hubspot-sdk-go/packages/pagination"
	"github.com/stainless-sdks/hubspot-sdk-go/packages/param"
)

// ListMembershipService contains methods and other services that help with
// interacting with the Hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewListMembershipService] method instead.
type ListMembershipService struct {
	Options []option.RequestOption
}

// NewListMembershipService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewListMembershipService(opts ...option.RequestOption) (r ListMembershipService) {
	r = ListMembershipService{}
	r.Options = opts
	return
}

// Fetch the memberships of a list in order sorted by the `recordId` of the records
// in the list.
//
// The `recordId`s are sorted in _ascending_ order if an `after` offset or no
// offset is provided. If only a `before` offset is provided, then the records are
// sorted in _descending_ order.
//
// The `after` offset parameter will take precedence over the `before` offset in a
// case where both are provided.
func (r *ListMembershipService) List(ctx context.Context, listID string, query ListMembershipListParams, opts ...option.RequestOption) (res *pagination.Page[JoinTimeAndRecordID], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if listID == "" {
		err = errors.New("missing required listId parameter")
		return
	}
	path := fmt.Sprintf("crm/v3/lists/%s/memberships", listID)
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

// Fetch the memberships of a list in order sorted by the `recordId` of the records
// in the list.
//
// The `recordId`s are sorted in _ascending_ order if an `after` offset or no
// offset is provided. If only a `before` offset is provided, then the records are
// sorted in _descending_ order.
//
// The `after` offset parameter will take precedence over the `before` offset in a
// case where both are provided.
func (r *ListMembershipService) ListAutoPaging(ctx context.Context, listID string, query ListMembershipListParams, opts ...option.RequestOption) *pagination.PageAutoPager[JoinTimeAndRecordID] {
	return pagination.NewPageAutoPager(r.List(ctx, listID, query, opts...))
}

// Add the records provided to the list. Records that do not exist or that are
// already members of the list are ignored.
//
// This endpoint only works for lists that have a `processingType` of `MANUAL` or
// `SNAPSHOT`.
func (r *ListMembershipService) Add(ctx context.Context, listID string, body ListMembershipAddParams, opts ...option.RequestOption) (res *MembershipsUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if listID == "" {
		err = errors.New("missing required listId parameter")
		return
	}
	path := fmt.Sprintf("crm/v3/lists/%s/memberships/add", listID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Add all of the records from a _source list_ (specified by the `sourceListId`) to
// a _destination list_ (specified by the `listId`). Records that are already
// members of the _destination list_ will be ignored. The _destination_ and _source
// list_ IDs must be different. The _destination_ and _source lists_ must contain
// records of the same type (e.g. contacts, companies, etc.).
//
// This endpoint only works for _destination lists_ that have a `processingType` of
// `MANUAL` or `SNAPSHOT`. The _source list_ can have any `processingType`.
//
// This endpoint only supports a `sourceListId` for lists with less than 100,000
// memberships.
func (r *ListMembershipService) AddAllFromList(ctx context.Context, sourceListID string, body ListMembershipAddAllFromListParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if body.ListID == "" {
		err = errors.New("missing required listId parameter")
		return
	}
	if sourceListID == "" {
		err = errors.New("missing required sourceListId parameter")
		return
	}
	path := fmt.Sprintf("crm/v3/lists/%s/memberships/add-from/%s", body.ListID, sourceListID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, nil, nil, opts...)
	return
}

// Add and/or remove records that have already been created in the system to and/or
// from a list.
//
// This endpoint only works for lists that have a `processingType` of `MANUAL` or
// `SNAPSHOT`.
func (r *ListMembershipService) AddAndRemove(ctx context.Context, listID string, body ListMembershipAddAndRemoveParams, opts ...option.RequestOption) (res *MembershipsUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if listID == "" {
		err = errors.New("missing required listId parameter")
		return
	}
	path := fmt.Sprintf("crm/v3/lists/%s/memberships/add-and-remove", listID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// For given record provide lists this record is member of.
func (r *ListMembershipService) GetLists(ctx context.Context, recordID string, query ListMembershipGetListsParams, opts ...option.RequestOption) (res *APICollectionResponseRecordListMembershipNoPaging, err error) {
	opts = slices.Concat(r.Options, opts)
	if query.ObjectTypeID == "" {
		err = errors.New("missing required objectTypeId parameter")
		return
	}
	if recordID == "" {
		err = errors.New("missing required recordId parameter")
		return
	}
	path := fmt.Sprintf("crm/v3/lists/records/%s/%s/memberships", query.ObjectTypeID, recordID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Fetch the memberships of a list in order sorted by the time the records were
// added to the list.
//
// The `recordId`s are sorted in _ascending_ order if an `after` offset or no
// offset is provided. If only a `before` offset is provided, then the records are
// sorted in _descending_ order.
//
// The `after` offset parameter will take precedence over the `before` offset in a
// case where both are provided.
func (r *ListMembershipService) GetPageOrderedByAddedToListDate(ctx context.Context, listID string, query ListMembershipGetPageOrderedByAddedToListDateParams, opts ...option.RequestOption) (res *pagination.Page[JoinTimeAndRecordID], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if listID == "" {
		err = errors.New("missing required listId parameter")
		return
	}
	path := fmt.Sprintf("crm/v3/lists/%s/memberships/join-order", listID)
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

// Fetch the memberships of a list in order sorted by the time the records were
// added to the list.
//
// The `recordId`s are sorted in _ascending_ order if an `after` offset or no
// offset is provided. If only a `before` offset is provided, then the records are
// sorted in _descending_ order.
//
// The `after` offset parameter will take precedence over the `before` offset in a
// case where both are provided.
func (r *ListMembershipService) GetPageOrderedByAddedToListDateAutoPaging(ctx context.Context, listID string, query ListMembershipGetPageOrderedByAddedToListDateParams, opts ...option.RequestOption) *pagination.PageAutoPager[JoinTimeAndRecordID] {
	return pagination.NewPageAutoPager(r.GetPageOrderedByAddedToListDate(ctx, listID, query, opts...))
}

// Remove the records provided from the list. Records that do not exist or that are
// not members of the list are ignored.
//
// This endpoint only works for lists that have a `processingType` of `MANUAL` or
// `SNAPSHOT`.
func (r *ListMembershipService) Remove(ctx context.Context, listID string, body ListMembershipRemoveParams, opts ...option.RequestOption) (res *MembershipsUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if listID == "" {
		err = errors.New("missing required listId parameter")
		return
	}
	path := fmt.Sprintf("crm/v3/lists/%s/memberships/remove", listID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Remove **all** of the records from a list. **_Note:_** _The list is not
// deleted._
//
// This endpoint only works for lists that have a `processingType` of `MANUAL` or
// `SNAPSHOT`.
//
// This endpoint only supports lists that have less than 100,000 memberships.
func (r *ListMembershipService) RemoveAll(ctx context.Context, listID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if listID == "" {
		err = errors.New("missing required listId parameter")
		return
	}
	path := fmt.Sprintf("crm/v3/lists/%s/memberships", listID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

type ListMembershipListParams struct {
	// The paging offset token for the page that comes `after` the previously requested
	// records.
	//
	// If provided, then the records in the response will be the records following the
	// offset, sorted in _ascending_ order. Takes precedence over the `before` offset.
	After param.Opt[string] `query:"after,omitzero" json:"-"`
	// The paging offset token for the page that comes `before` the previously
	// requested records.
	//
	// If provided, then the records in the response will be the records preceding the
	// offset, sorted in _descending_ order.
	Before param.Opt[string] `query:"before,omitzero" json:"-"`
	// The number of records to return in the response. The maximum `limit` is 250.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ListMembershipListParams]'s query parameters as
// `url.Values`.
func (r ListMembershipListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ListMembershipAddParams struct {
	Body []string
	paramObj
}

func (r ListMembershipAddParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.Body)
}
func (r *ListMembershipAddParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.Body)
}

type ListMembershipAddAllFromListParams struct {
	ListID string `path:"listId,required" json:"-"`
	paramObj
}

type ListMembershipAddAndRemoveParams struct {
	// The IDs of the records to add and/or remove from a list.
	MembershipChangeRequest MembershipChangeRequestParam
	paramObj
}

func (r ListMembershipAddAndRemoveParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.MembershipChangeRequest)
}
func (r *ListMembershipAddAndRemoveParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.MembershipChangeRequest)
}

type ListMembershipGetListsParams struct {
	ObjectTypeID string `path:"objectTypeId,required" json:"-"`
	paramObj
}

type ListMembershipGetPageOrderedByAddedToListDateParams struct {
	// The paging offset token for the page that comes `after` the previously requested
	// records.
	//
	// If provided, then the records in the response will be the records following the
	// offset, sorted in _ascending_ order. Takes precedence over the `before` offset.
	After param.Opt[string] `query:"after,omitzero" json:"-"`
	// The paging offset token for the page that comes `before` the previously
	// requested records.
	//
	// If provided, then the records in the response will be the records preceding the
	// offset, sorted in _descending_ order.
	Before param.Opt[string] `query:"before,omitzero" json:"-"`
	// The number of records to return in the response. The maximum `limit` is 250.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ListMembershipGetPageOrderedByAddedToListDateParams]'s
// query parameters as `url.Values`.
func (r ListMembershipGetPageOrderedByAddedToListDateParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ListMembershipRemoveParams struct {
	Body []string
	paramObj
}

func (r ListMembershipRemoveParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.Body)
}
func (r *ListMembershipRemoveParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.Body)
}
