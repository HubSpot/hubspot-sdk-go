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
	"time"

	"github.com/HubSpot/hubspot-sdk-go/internal/apijson"
	"github.com/HubSpot/hubspot-sdk-go/internal/apiquery"
	shimjson "github.com/HubSpot/hubspot-sdk-go/internal/encoding/json"
	"github.com/HubSpot/hubspot-sdk-go/internal/requestconfig"
	"github.com/HubSpot/hubspot-sdk-go/option"
	"github.com/HubSpot/hubspot-sdk-go/packages/pagination"
	"github.com/HubSpot/hubspot-sdk-go/packages/param"
	"github.com/HubSpot/hubspot-sdk-go/packages/respjson"
	"github.com/HubSpot/hubspot-sdk-go/shared"
)

// ListService contains methods and other services that help with interacting with
// the hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewListService] method instead.
type ListService struct {
	options []option.RequestOption
}

// NewListService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewListService(opts ...option.RequestOption) (r ListService) {
	r = ListService{}
	r.options = opts
	return
}

func (r *ListService) New(ctx context.Context, body ListNewParams, opts ...option.RequestOption) (res *ListCreateResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "crm/lists/2026-03"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

func (r *ListService) List(ctx context.Context, query ListListParams, opts ...option.RequestOption) (res *ListsByIDResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "crm/lists/2026-03"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

func (r *ListService) Delete(ctx context.Context, listID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if listID == "" {
		err = errors.New("missing required listId parameter")
		return err
	}
	path := fmt.Sprintf("crm/lists/2026-03/%s", url.PathEscape(listID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

func (r *ListService) AddAndRemoveMemberships(ctx context.Context, listID string, body ListAddAndRemoveMembershipsParams, opts ...option.RequestOption) (res *MembershipsUpdateResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if listID == "" {
		err = errors.New("missing required listId parameter")
		return nil, err
	}
	path := fmt.Sprintf("crm/lists/2026-03/%s/memberships/add-and-remove", url.PathEscape(listID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return res, err
}

func (r *ListService) AddMemberships(ctx context.Context, listID string, body ListAddMembershipsParams, opts ...option.RequestOption) (res *MembershipsUpdateResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if listID == "" {
		err = errors.New("missing required listId parameter")
		return nil, err
	}
	path := fmt.Sprintf("crm/lists/2026-03/%s/memberships/add", url.PathEscape(listID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return res, err
}

func (r *ListService) AddMembershipsFrom(ctx context.Context, sourceListID string, body ListAddMembershipsFromParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if body.ListID == "" {
		err = errors.New("missing required listId parameter")
		return err
	}
	if sourceListID == "" {
		err = errors.New("missing required sourceListId parameter")
		return err
	}
	path := fmt.Sprintf("crm/lists/2026-03/%s/memberships/add-from/%s", url.PathEscape(body.ListID), url.PathEscape(sourceListID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, nil, nil, opts...)
	return err
}

func (r *ListService) BatchReadMemberships(ctx context.Context, body ListBatchReadMembershipsParams, opts ...option.RequestOption) (res *BatchResponseRecordIDWithMemberships, err error) {
	opts = slices.Concat(r.options, opts)
	path := "crm/lists/2026-03/records/memberships/batch/read"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

func (r *ListService) NewFolder(ctx context.Context, body ListNewFolderParams, opts ...option.RequestOption) (res *ListFolderCreateResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "crm/lists/2026-03/folders"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

func (r *ListService) NewIDMapping(ctx context.Context, body ListNewIDMappingParams, opts ...option.RequestOption) (res *PublicBatchMigrationMapping, err error) {
	opts = slices.Concat(r.options, opts)
	path := "crm/lists/2026-03/idmapping"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

func (r *ListService) DeleteFolder(ctx context.Context, folderID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if folderID == "" {
		err = errors.New("missing required folderId parameter")
		return err
	}
	path := fmt.Sprintf("crm/lists/2026-03/folders/%s", url.PathEscape(folderID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

func (r *ListService) DeleteMemberships(ctx context.Context, listID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if listID == "" {
		err = errors.New("missing required listId parameter")
		return err
	}
	path := fmt.Sprintf("crm/lists/2026-03/%s/memberships", url.PathEscape(listID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

func (r *ListService) Get(ctx context.Context, listID string, query ListGetParams, opts ...option.RequestOption) (res *ListFetchResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if listID == "" {
		err = errors.New("missing required listId parameter")
		return nil, err
	}
	path := fmt.Sprintf("crm/lists/2026-03/%s", url.PathEscape(listID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

func (r *ListService) GetByObjectTypeAndName(ctx context.Context, listName string, params ListGetByObjectTypeAndNameParams, opts ...option.RequestOption) (res *ListFetchResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if params.ObjectTypeID == "" {
		err = errors.New("missing required objectTypeId parameter")
		return nil, err
	}
	if listName == "" {
		err = errors.New("missing required listName parameter")
		return nil, err
	}
	path := fmt.Sprintf("crm/lists/2026-03/object-type-id/%s/name/%s", url.PathEscape(params.ObjectTypeID), url.PathEscape(listName))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

func (r *ListService) GetIDMapping(ctx context.Context, query ListGetIDMappingParams, opts ...option.RequestOption) (res *PublicMigrationMapping, err error) {
	opts = slices.Concat(r.options, opts)
	path := "crm/lists/2026-03/idmapping"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

func (r *ListService) GetMembershipsJoinOrder(ctx context.Context, listID string, query ListGetMembershipsJoinOrderParams, opts ...option.RequestOption) (res *pagination.Page[JoinTimeAndRecordID], err error) {
	var raw *http.Response
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if listID == "" {
		err = errors.New("missing required listId parameter")
		return nil, err
	}
	path := fmt.Sprintf("crm/lists/2026-03/%s/memberships/join-order", url.PathEscape(listID))
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

func (r *ListService) GetMembershipsJoinOrderAutoPaging(ctx context.Context, listID string, query ListGetMembershipsJoinOrderParams, opts ...option.RequestOption) *pagination.PageAutoPager[JoinTimeAndRecordID] {
	return pagination.NewPageAutoPager(r.GetMembershipsJoinOrder(ctx, listID, query, opts...))
}

func (r *ListService) GetRecordMemberships(ctx context.Context, recordID string, query ListGetRecordMembershipsParams, opts ...option.RequestOption) (res *APICollectionResponseRecordListMembership, err error) {
	opts = slices.Concat(r.options, opts)
	if query.ObjectTypeID == "" {
		err = errors.New("missing required objectTypeId parameter")
		return nil, err
	}
	if recordID == "" {
		err = errors.New("missing required recordId parameter")
		return nil, err
	}
	path := fmt.Sprintf("crm/lists/2026-03/records/%s/%s/memberships", url.PathEscape(query.ObjectTypeID), url.PathEscape(recordID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

func (r *ListService) GetScheduleConversion(ctx context.Context, listID string, opts ...option.RequestOption) (res *PublicListConversionResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if listID == "" {
		err = errors.New("missing required listId parameter")
		return nil, err
	}
	path := fmt.Sprintf("crm/lists/2026-03/%s/schedule-conversion", url.PathEscape(listID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

func (r *ListService) GetSizeAndEditsHistoryBetween(ctx context.Context, listID string, query ListGetSizeAndEditsHistoryBetweenParams, opts ...option.RequestOption) (res *ListSizeAndEditHistoryResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if listID == "" {
		err = errors.New("missing required listId parameter")
		return nil, err
	}
	path := fmt.Sprintf("crm/lists/2026-03/%s/size-and-edits-history/between", url.PathEscape(listID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

func (r *ListService) ListBySearch(ctx context.Context, body ListListBySearchParams, opts ...option.RequestOption) (res *ListSearchResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "crm/lists/2026-03/search"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

func (r *ListService) ListFolders(ctx context.Context, query ListListFoldersParams, opts ...option.RequestOption) (res *ListFolderFetchResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "crm/lists/2026-03/folders"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

func (r *ListService) ListMemberships(ctx context.Context, listID string, query ListListMembershipsParams, opts ...option.RequestOption) (res *pagination.Page[JoinTimeAndRecordID], err error) {
	var raw *http.Response
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if listID == "" {
		err = errors.New("missing required listId parameter")
		return nil, err
	}
	path := fmt.Sprintf("crm/lists/2026-03/%s/memberships", url.PathEscape(listID))
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

func (r *ListService) ListMembershipsAutoPaging(ctx context.Context, listID string, query ListListMembershipsParams, opts ...option.RequestOption) *pagination.PageAutoPager[JoinTimeAndRecordID] {
	return pagination.NewPageAutoPager(r.ListMemberships(ctx, listID, query, opts...))
}

func (r *ListService) MoveFolder(ctx context.Context, newParentFolderID string, body ListMoveFolderParams, opts ...option.RequestOption) (res *ListFolderFetchResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if body.FolderID == "" {
		err = errors.New("missing required folderId parameter")
		return nil, err
	}
	if newParentFolderID == "" {
		err = errors.New("missing required newParentFolderId parameter")
		return nil, err
	}
	path := fmt.Sprintf("crm/lists/2026-03/folders/%s/move/%s", url.PathEscape(body.FolderID), url.PathEscape(newParentFolderID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, nil, &res, opts...)
	return res, err
}

func (r *ListService) MoveList(ctx context.Context, body ListMoveListParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "crm/lists/2026-03/folders/move-list"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, nil, opts...)
	return err
}

func (r *ListService) RemoveMemberships(ctx context.Context, listID string, body ListRemoveMembershipsParams, opts ...option.RequestOption) (res *MembershipsUpdateResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if listID == "" {
		err = errors.New("missing required listId parameter")
		return nil, err
	}
	path := fmt.Sprintf("crm/lists/2026-03/%s/memberships/remove", url.PathEscape(listID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return res, err
}

func (r *ListService) RenameFolder(ctx context.Context, folderID string, body ListRenameFolderParams, opts ...option.RequestOption) (res *ListFolderFetchResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if folderID == "" {
		err = errors.New("missing required folderId parameter")
		return nil, err
	}
	path := fmt.Sprintf("crm/lists/2026-03/folders/%s/rename", url.PathEscape(folderID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return res, err
}

func (r *ListService) Restore(ctx context.Context, listID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if listID == "" {
		err = errors.New("missing required listId parameter")
		return err
	}
	path := fmt.Sprintf("crm/lists/2026-03/%s/restore", url.PathEscape(listID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, nil, nil, opts...)
	return err
}

func (r *ListService) ScheduleConversion(ctx context.Context, listID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if listID == "" {
		err = errors.New("missing required listId parameter")
		return err
	}
	path := fmt.Sprintf("crm/lists/2026-03/%s/schedule-conversion", url.PathEscape(listID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

func (r *ListService) UpdateListFilters(ctx context.Context, listID string, params ListUpdateListFiltersParams, opts ...option.RequestOption) (res *ListUpdateResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if listID == "" {
		err = errors.New("missing required listId parameter")
		return nil, err
	}
	path := fmt.Sprintf("crm/lists/2026-03/%s/update-list-filters", url.PathEscape(listID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return res, err
}

func (r *ListService) UpdateListName(ctx context.Context, listID string, body ListUpdateListNameParams, opts ...option.RequestOption) (res *ListUpdateResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if listID == "" {
		err = errors.New("missing required listId parameter")
		return nil, err
	}
	path := fmt.Sprintf("crm/lists/2026-03/%s/update-list-name", url.PathEscape(listID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return res, err
}

func (r *ListService) UpdateScheduleConversion(ctx context.Context, listID string, body ListUpdateScheduleConversionParams, opts ...option.RequestOption) (res *PublicListConversionResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if listID == "" {
		err = errors.New("missing required listId parameter")
		return nil, err
	}
	path := fmt.Sprintf("crm/lists/2026-03/%s/schedule-conversion", url.PathEscape(listID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return res, err
}

type APICollectionResponseJoinTimeAndRecordID struct {
	Results []JoinTimeAndRecordID `json:"results" api:"required"`
	Paging  shared.Paging         `json:"paging"`
	// The total number of records that match the query.
	Total int64 `json:"total"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Results     respjson.Field
		Paging      respjson.Field
		Total       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r APICollectionResponseJoinTimeAndRecordID) RawJSON() string { return r.JSON.raw }
func (r *APICollectionResponseJoinTimeAndRecordID) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type APICollectionResponseRecordListMembership struct {
	Results []RecordListMembership `json:"results" api:"required"`
	Paging  shared.Paging          `json:"paging"`
	Total   int64                  `json:"total"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Results     respjson.Field
		Paging      respjson.Field
		Total       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r APICollectionResponseRecordListMembership) RawJSON() string { return r.JSON.raw }
func (r *APICollectionResponseRecordListMembership) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Inputs is required.
type BatchInputRecordIDInputParam struct {
	Inputs []RecordIDInputParam `json:"inputs,omitzero" api:"required"`
	paramObj
}

func (r BatchInputRecordIDInputParam) MarshalJSON() (data []byte, err error) {
	type shadow BatchInputRecordIDInputParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BatchInputRecordIDInputParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BatchResponseRecordIDWithMemberships struct {
	CompletedAt time.Time                 `json:"completedAt" api:"required" format:"date-time"`
	Results     []RecordIDWithMemberships `json:"results" api:"required"`
	StartedAt   time.Time                 `json:"startedAt" api:"required" format:"date-time"`
	// Any of "CANCELED", "COMPLETE", "PENDING", "PROCESSING".
	Status      BatchResponseRecordIDWithMembershipsStatus `json:"status" api:"required"`
	Links       map[string]string                          `json:"links"`
	RequestedAt time.Time                                  `json:"requestedAt" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CompletedAt respjson.Field
		Results     respjson.Field
		StartedAt   respjson.Field
		Status      respjson.Field
		Links       respjson.Field
		RequestedAt respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BatchResponseRecordIDWithMemberships) RawJSON() string { return r.JSON.raw }
func (r *BatchResponseRecordIDWithMemberships) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BatchResponseRecordIDWithMembershipsStatus string

const (
	BatchResponseRecordIDWithMembershipsStatusCanceled   BatchResponseRecordIDWithMembershipsStatus = "CANCELED"
	BatchResponseRecordIDWithMembershipsStatusComplete   BatchResponseRecordIDWithMembershipsStatus = "COMPLETE"
	BatchResponseRecordIDWithMembershipsStatusPending    BatchResponseRecordIDWithMembershipsStatus = "PENDING"
	BatchResponseRecordIDWithMembershipsStatusProcessing BatchResponseRecordIDWithMembershipsStatus = "PROCESSING"
)

type JoinTimeAndRecordID struct {
	// The date and time when the record was added to the list.
	MembershipTimestamp time.Time `json:"membershipTimestamp" api:"required" format:"date-time"`
	// The unique identifier of the record.
	RecordID string `json:"recordId" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		MembershipTimestamp respjson.Field
		RecordID            respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r JoinTimeAndRecordID) RawJSON() string { return r.JSON.raw }
func (r *JoinTimeAndRecordID) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Name, ObjectTypeID, ProcessingType are required.
type ListCreateRequestParam struct {
	// The name of the list, which must be globally unique across all public lists in
	// the portal.
	Name string `json:"name" api:"required"`
	// The object type ID of the type of objects that the list will store.
	ObjectTypeID string `json:"objectTypeId" api:"required"`
	// The processing type of the list. One of: `SNAPSHOT`, `MANUAL`, or `DYNAMIC`.
	ProcessingType string `json:"processingType" api:"required"`
	// The ID of the folder that the list should be created in. If left blank, then the
	// list will be created in the root of the list folder structure.
	ListFolderID param.Opt[int64] `json:"listFolderId,omitzero"`
	// The list of custom properties to tie to the list. Custom property name is the
	// key, the value is the value.
	CustomProperties map[string]string `json:"customProperties,omitzero"`
	// Filter branch object containing filtering criteria for the list
	FilterBranch       ListCreateRequestFilterBranchUnionParam `json:"filterBranch,omitzero"`
	ListPermissions    PublicListPermissionsParam              `json:"listPermissions,omitzero"`
	MembershipSettings PublicMembershipSettingsParam           `json:"membershipSettings,omitzero"`
	paramObj
}

func (r ListCreateRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow ListCreateRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ListCreateRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ListCreateRequestFilterBranchUnionParam struct {
	OfOr                  *PublicOrFilterBranchParam                  `json:",omitzero,inline"`
	OfAnd                 *PublicAndFilterBranchParam                 `json:",omitzero,inline"`
	OfNotAll              *PublicNotAllFilterBranchParam              `json:",omitzero,inline"`
	OfNotAny              *PublicNotAnyFilterBranchParam              `json:",omitzero,inline"`
	OfRestricted          *PublicRestrictedFilterBranchParam          `json:",omitzero,inline"`
	OfUnifiedEvents       *PublicUnifiedEventsFilterBranchParam       `json:",omitzero,inline"`
	OfPropertyAssociation *PublicPropertyAssociationFilterBranchParam `json:",omitzero,inline"`
	OfAssociation         *PublicAssociationFilterBranchParam         `json:",omitzero,inline"`
	paramUnion
}

func (u ListCreateRequestFilterBranchUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfOr,
		u.OfAnd,
		u.OfNotAll,
		u.OfNotAny,
		u.OfRestricted,
		u.OfUnifiedEvents,
		u.OfPropertyAssociation,
		u.OfAssociation)
}
func (u *ListCreateRequestFilterBranchUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func init() {
	apijson.RegisterUnion[ListCreateRequestFilterBranchUnionParam](
		"filterBranchType",
		apijson.Discriminator[PublicOrFilterBranchParam]("OR"),
		apijson.Discriminator[PublicAndFilterBranchParam]("AND"),
		apijson.Discriminator[PublicNotAllFilterBranchParam]("NOT_ALL"),
		apijson.Discriminator[PublicNotAnyFilterBranchParam]("NOT_ANY"),
		apijson.Discriminator[PublicRestrictedFilterBranchParam]("RESTRICTED"),
		apijson.Discriminator[PublicUnifiedEventsFilterBranchParam]("UNIFIED_EVENTS"),
		apijson.Discriminator[PublicPropertyAssociationFilterBranchParam]("PROPERTY_ASSOCIATION"),
		apijson.Discriminator[PublicAssociationFilterBranchParam]("ASSOCIATION"),
	)
}

type ListCreateResponse struct {
	List PublicObjectList `json:"list" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		List        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ListCreateResponse) RawJSON() string { return r.JSON.raw }
func (r *ListCreateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ListFetchResponse struct {
	List PublicObjectList `json:"list" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		List        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ListFetchResponse) RawJSON() string { return r.JSON.raw }
func (r *ListFetchResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property FilterBranch is required.
type ListFilterUpdateRequestParam struct {
	// Updated filtering criteria for the list
	FilterBranch ListFilterUpdateRequestFilterBranchUnionParam `json:"filterBranch,omitzero" api:"required"`
	paramObj
}

func (r ListFilterUpdateRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow ListFilterUpdateRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ListFilterUpdateRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ListFilterUpdateRequestFilterBranchUnionParam struct {
	OfOr                  *PublicOrFilterBranchParam                  `json:",omitzero,inline"`
	OfAnd                 *PublicAndFilterBranchParam                 `json:",omitzero,inline"`
	OfNotAll              *PublicNotAllFilterBranchParam              `json:",omitzero,inline"`
	OfNotAny              *PublicNotAnyFilterBranchParam              `json:",omitzero,inline"`
	OfRestricted          *PublicRestrictedFilterBranchParam          `json:",omitzero,inline"`
	OfUnifiedEvents       *PublicUnifiedEventsFilterBranchParam       `json:",omitzero,inline"`
	OfPropertyAssociation *PublicPropertyAssociationFilterBranchParam `json:",omitzero,inline"`
	OfAssociation         *PublicAssociationFilterBranchParam         `json:",omitzero,inline"`
	paramUnion
}

func (u ListFilterUpdateRequestFilterBranchUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfOr,
		u.OfAnd,
		u.OfNotAll,
		u.OfNotAny,
		u.OfRestricted,
		u.OfUnifiedEvents,
		u.OfPropertyAssociation,
		u.OfAssociation)
}
func (u *ListFilterUpdateRequestFilterBranchUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func init() {
	apijson.RegisterUnion[ListFilterUpdateRequestFilterBranchUnionParam](
		"filterBranchType",
		apijson.Discriminator[PublicOrFilterBranchParam]("OR"),
		apijson.Discriminator[PublicAndFilterBranchParam]("AND"),
		apijson.Discriminator[PublicNotAllFilterBranchParam]("NOT_ALL"),
		apijson.Discriminator[PublicNotAnyFilterBranchParam]("NOT_ANY"),
		apijson.Discriminator[PublicRestrictedFilterBranchParam]("RESTRICTED"),
		apijson.Discriminator[PublicUnifiedEventsFilterBranchParam]("UNIFIED_EVENTS"),
		apijson.Discriminator[PublicPropertyAssociationFilterBranchParam]("PROPERTY_ASSOCIATION"),
		apijson.Discriminator[PublicAssociationFilterBranchParam]("ASSOCIATION"),
	)
}

// The property Name is required.
type ListFolderCreateRequestParam struct {
	// The name of the folder to be created.
	Name string `json:"name" api:"required"`
	// The folder this should be created in, if not specified will be created in the
	// root folder 0.
	ParentFolderID param.Opt[string] `json:"parentFolderId,omitzero"`
	paramObj
}

func (r ListFolderCreateRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow ListFolderCreateRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ListFolderCreateRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ListFolderCreateResponse struct {
	Folder PublicListFolder `json:"folder" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Folder      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ListFolderCreateResponse) RawJSON() string { return r.JSON.raw }
func (r *ListFolderCreateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ListFolderFetchResponse struct {
	Folder PublicListFolder `json:"folder" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Folder      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ListFolderFetchResponse) RawJSON() string { return r.JSON.raw }
func (r *ListFolderFetchResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties ListID, NewFolderID are required.
type ListMoveRequestParam struct {
	// The Id of the list to move.
	ListID string `json:"listId" api:"required"`
	// The Id of folder to move the list to, the root folder is Id 0.
	NewFolderID string `json:"newFolderId" api:"required"`
	paramObj
}

func (r ListMoveRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow ListMoveRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ListMoveRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties ListIDs, Offset, ProcessingTypes are required.
type ListSearchRequestParam struct {
	// ILS list ids to be included in search results. If not specified, all lists
	// matching other criteria will be included
	ListIDs []string `json:"listIds,omitzero" api:"required"`
	// Value used to paginate through lists. The `offset` provided in the response can
	// be used in the next request to fetch the next page of results. Defaults to `0`
	// if no offset is provided.
	Offset int64 `json:"offset" api:"required"`
	// List processing types to be included in search results. If not specified, all
	// lists with all processing types will be included.
	ProcessingTypes []string `json:"processingTypes,omitzero" api:"required"`
	// The number of lists to include in the response. Defaults to `20` if no value is
	// provided. The max `count` is `500`.
	Count        param.Opt[int64]  `json:"count,omitzero"`
	ObjectTypeID param.Opt[string] `json:"objectTypeId,omitzero"`
	// The `query` that will be used to search for lists by list name. If no `query` is
	// provided, then the results will include all lists.
	Query param.Opt[string] `json:"query,omitzero"`
	// Sort field and order
	Sort param.Opt[string] `json:"sort,omitzero"`
	// The property names of any additional list properties to include in the response.
	// Properties that do not exist or that are empty for a particular list are not
	// included in the response.
	//
	// By default, all requests will fetch the following properties for each list:
	// `hs_list_size`, `hs_last_record_added_at`, `hs_last_record_removed_at`,
	// `hs_folder_name`, and `hs_list_reference_count`.
	AdditionalFilterProperties []string `json:"additional_filter_properties,omitzero"`
	paramObj
}

func (r ListSearchRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow ListSearchRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ListSearchRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ListSearchResponse struct {
	// Whether or not there are more results to page through.
	HasMore bool `json:"hasMore" api:"required"`
	// The lists that matched the search criteria.
	Lists []PublicObjectListSearchResult `json:"lists" api:"required"`
	// Value to be passed in a future request to paginate through list search results.
	Offset int64 `json:"offset" api:"required"`
	// The total number of lists that match the search criteria.
	Total int64 `json:"total" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		HasMore     respjson.Field
		Lists       respjson.Field
		Offset      respjson.Field
		Total       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ListSearchResponse) RawJSON() string { return r.JSON.raw }
func (r *ListSearchResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ListSizeAndEditHistoryResponse struct {
	EditHistory []time.Time         `json:"editHistory" api:"required" format:"date-time"`
	SizeHistory []ListSizeDataPoint `json:"sizeHistory" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		EditHistory respjson.Field
		SizeHistory respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ListSizeAndEditHistoryResponse) RawJSON() string { return r.JSON.raw }
func (r *ListSizeAndEditHistoryResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ListSizeDataPoint struct {
	Size      int64     `json:"size" api:"required"`
	Timestamp time.Time `json:"timestamp" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Size        respjson.Field
		Timestamp   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ListSizeDataPoint) RawJSON() string { return r.JSON.raw }
func (r *ListSizeDataPoint) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ListUpdateResponse struct {
	UpdatedList PublicObjectList `json:"updatedList"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		UpdatedList respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ListUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *ListUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ListsByIDResponse struct {
	// The object list definitions.
	Lists []PublicObjectList `json:"lists" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Lists       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ListsByIDResponse) RawJSON() string { return r.JSON.raw }
func (r *ListsByIDResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties RecordIDsToAdd, RecordIDsToRemove are required.
type MembershipChangeRequestParam struct {
	RecordIDsToAdd    []string `json:"recordIdsToAdd,omitzero" api:"required"`
	RecordIDsToRemove []string `json:"recordIdsToRemove,omitzero" api:"required"`
	paramObj
}

func (r MembershipChangeRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow MembershipChangeRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MembershipChangeRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MembershipsUpdateResponse struct {
	// The IDs of the records that were `missing` (e.g. did not exist in the portal)
	// and so were not `added` or `removed`.
	RecordIDsMissing []string `json:"recordIdsMissing" api:"required"`
	// The IDs of the records that were `removed` from the list.
	RecordIDsRemoved []string `json:"recordIdsRemoved" api:"required"`
	RecordsIDsAdded  []string `json:"recordsIdsAdded" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		RecordIDsMissing respjson.Field
		RecordIDsRemoved respjson.Field
		RecordsIDsAdded  respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MembershipsUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *MembershipsUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PublicAbsoluteComparativeTimestampRefineBy struct {
	// Timestamp comparison options (BEFORE, AFTER)
	Comparison string `json:"comparison" api:"required"`
	// Timestamp to be used in refine by criteria
	Timestamp int64 `json:"timestamp" api:"required"`
	// type of refine by criteria (ABSOLUTE_COMPARATIVE)
	//
	// Any of "ABSOLUTE_COMPARATIVE".
	Type PublicAbsoluteComparativeTimestampRefineByType `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Comparison  respjson.Field
		Timestamp   respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PublicAbsoluteComparativeTimestampRefineBy) RawJSON() string { return r.JSON.raw }
func (r *PublicAbsoluteComparativeTimestampRefineBy) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this PublicAbsoluteComparativeTimestampRefineBy to a
// PublicAbsoluteComparativeTimestampRefineByParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// PublicAbsoluteComparativeTimestampRefineByParam.Overrides()
func (r PublicAbsoluteComparativeTimestampRefineBy) ToParam() PublicAbsoluteComparativeTimestampRefineByParam {
	return param.Override[PublicAbsoluteComparativeTimestampRefineByParam](json.RawMessage(r.RawJSON()))
}

// type of refine by criteria (ABSOLUTE_COMPARATIVE)
type PublicAbsoluteComparativeTimestampRefineByType string

const (
	PublicAbsoluteComparativeTimestampRefineByTypeAbsoluteComparative PublicAbsoluteComparativeTimestampRefineByType = "ABSOLUTE_COMPARATIVE"
)

// The properties Comparison, Timestamp, Type are required.
type PublicAbsoluteComparativeTimestampRefineByParam struct {
	// Timestamp comparison options (BEFORE, AFTER)
	Comparison string `json:"comparison" api:"required"`
	// Timestamp to be used in refine by criteria
	Timestamp int64 `json:"timestamp" api:"required"`
	// type of refine by criteria (ABSOLUTE_COMPARATIVE)
	//
	// Any of "ABSOLUTE_COMPARATIVE".
	Type PublicAbsoluteComparativeTimestampRefineByType `json:"type,omitzero" api:"required"`
	paramObj
}

func (r PublicAbsoluteComparativeTimestampRefineByParam) MarshalJSON() (data []byte, err error) {
	type shadow PublicAbsoluteComparativeTimestampRefineByParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PublicAbsoluteComparativeTimestampRefineByParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PublicAbsoluteRangedTimestampRefineBy struct {
	// Lower range timestamp of refinement criteria
	LowerTimestamp int64 `json:"lowerTimestamp" api:"required"`
	// Type of range of refinement critaria (BETWEEN, NOT_BETWEEN)
	RangeType string `json:"rangeType" api:"required"`
	// type of refine by criteria (ABSOLUTE_RANGED)
	//
	// Any of "ABSOLUTE_RANGED".
	Type PublicAbsoluteRangedTimestampRefineByType `json:"type" api:"required"`
	// Upper range timestamp of refinement criteria
	UpperTimestamp int64 `json:"upperTimestamp" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		LowerTimestamp respjson.Field
		RangeType      respjson.Field
		Type           respjson.Field
		UpperTimestamp respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PublicAbsoluteRangedTimestampRefineBy) RawJSON() string { return r.JSON.raw }
func (r *PublicAbsoluteRangedTimestampRefineBy) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this PublicAbsoluteRangedTimestampRefineBy to a
// PublicAbsoluteRangedTimestampRefineByParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// PublicAbsoluteRangedTimestampRefineByParam.Overrides()
func (r PublicAbsoluteRangedTimestampRefineBy) ToParam() PublicAbsoluteRangedTimestampRefineByParam {
	return param.Override[PublicAbsoluteRangedTimestampRefineByParam](json.RawMessage(r.RawJSON()))
}

// type of refine by criteria (ABSOLUTE_RANGED)
type PublicAbsoluteRangedTimestampRefineByType string

const (
	PublicAbsoluteRangedTimestampRefineByTypeAbsoluteRanged PublicAbsoluteRangedTimestampRefineByType = "ABSOLUTE_RANGED"
)

// The properties LowerTimestamp, RangeType, Type, UpperTimestamp are required.
type PublicAbsoluteRangedTimestampRefineByParam struct {
	// Lower range timestamp of refinement criteria
	LowerTimestamp int64 `json:"lowerTimestamp" api:"required"`
	// Type of range of refinement critaria (BETWEEN, NOT_BETWEEN)
	RangeType string `json:"rangeType" api:"required"`
	// type of refine by criteria (ABSOLUTE_RANGED)
	//
	// Any of "ABSOLUTE_RANGED".
	Type PublicAbsoluteRangedTimestampRefineByType `json:"type,omitzero" api:"required"`
	// Upper range timestamp of refinement criteria
	UpperTimestamp int64 `json:"upperTimestamp" api:"required"`
	paramObj
}

func (r PublicAbsoluteRangedTimestampRefineByParam) MarshalJSON() (data []byte, err error) {
	type shadow PublicAbsoluteRangedTimestampRefineByParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PublicAbsoluteRangedTimestampRefineByParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PublicAdsSearchFilter struct {
	// Ad network (ADWORDS, FACEBOOK, LINKEDIN, ALL)
	AdNetwork string `json:"adNetwork" api:"required"`
	// Type of ad entity (KEYWORD, ADGROUP, AD, CAMPAIGN)
	EntityType string `json:"entityType" api:"required"`
	// Type of the filter (ADS_SEARCH)
	//
	// Any of "ADS_SEARCH".
	FilterType PublicAdsSearchFilterFilterType `json:"filterType" api:"required"`
	// Operator to be applied (CONTAINS, IS_EQUAL_TO, ENDS_WITH, STARTS_WITH, IS_KNOWN)
	Operator    string   `json:"operator" api:"required"`
	SearchTerms []string `json:"searchTerms" api:"required"`
	// Search term to match an ad
	SearchTermType string `json:"searchTermType" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AdNetwork      respjson.Field
		EntityType     respjson.Field
		FilterType     respjson.Field
		Operator       respjson.Field
		SearchTerms    respjson.Field
		SearchTermType respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PublicAdsSearchFilter) RawJSON() string { return r.JSON.raw }
func (r *PublicAdsSearchFilter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this PublicAdsSearchFilter to a PublicAdsSearchFilterParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// PublicAdsSearchFilterParam.Overrides()
func (r PublicAdsSearchFilter) ToParam() PublicAdsSearchFilterParam {
	return param.Override[PublicAdsSearchFilterParam](json.RawMessage(r.RawJSON()))
}

// Type of the filter (ADS_SEARCH)
type PublicAdsSearchFilterFilterType string

const (
	PublicAdsSearchFilterFilterTypeAdsSearch PublicAdsSearchFilterFilterType = "ADS_SEARCH"
)

// The properties AdNetwork, EntityType, FilterType, Operator, SearchTerms,
// SearchTermType are required.
type PublicAdsSearchFilterParam struct {
	// Ad network (ADWORDS, FACEBOOK, LINKEDIN, ALL)
	AdNetwork string `json:"adNetwork" api:"required"`
	// Type of ad entity (KEYWORD, ADGROUP, AD, CAMPAIGN)
	EntityType string `json:"entityType" api:"required"`
	// Type of the filter (ADS_SEARCH)
	//
	// Any of "ADS_SEARCH".
	FilterType PublicAdsSearchFilterFilterType `json:"filterType,omitzero" api:"required"`
	// Operator to be applied (CONTAINS, IS_EQUAL_TO, ENDS_WITH, STARTS_WITH, IS_KNOWN)
	Operator    string   `json:"operator" api:"required"`
	SearchTerms []string `json:"searchTerms,omitzero" api:"required"`
	// Search term to match an ad
	SearchTermType string `json:"searchTermType" api:"required"`
	paramObj
}

func (r PublicAdsSearchFilterParam) MarshalJSON() (data []byte, err error) {
	type shadow PublicAdsSearchFilterParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PublicAdsSearchFilterParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PublicAdsTimeFilter struct {
	// Filter type (ADS_TIME)
	//
	// Any of "ADS_TIME".
	FilterType PublicAdsTimeFilterFilterType `json:"filterType" api:"required"`
	// Refinement criteria
	PruningRefineBy PublicAdsTimeFilterPruningRefineByUnion `json:"pruningRefineBy" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FilterType      respjson.Field
		PruningRefineBy respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PublicAdsTimeFilter) RawJSON() string { return r.JSON.raw }
func (r *PublicAdsTimeFilter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this PublicAdsTimeFilter to a PublicAdsTimeFilterParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// PublicAdsTimeFilterParam.Overrides()
func (r PublicAdsTimeFilter) ToParam() PublicAdsTimeFilterParam {
	return param.Override[PublicAdsTimeFilterParam](json.RawMessage(r.RawJSON()))
}

// Filter type (ADS_TIME)
type PublicAdsTimeFilterFilterType string

const (
	PublicAdsTimeFilterFilterTypeAdsTime PublicAdsTimeFilterFilterType = "ADS_TIME"
)

// PublicAdsTimeFilterPruningRefineByUnion contains all possible properties and
// values from [PublicNumOccurrencesRefineBy], [PublicSetOccurrencesRefineBy],
// [PublicRelativeComparativeTimestampRefineBy],
// [PublicRelativeRangedTimestampRefineBy],
// [PublicAbsoluteComparativeTimestampRefineBy],
// [PublicAbsoluteRangedTimestampRefineBy], [PublicAllHistoryRefineBy],
// [PublicTimePointOperation], [PublicRangedTimeOperation].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type PublicAdsTimeFilterPruningRefineByUnion struct {
	Type string `json:"type"`
	// This field is from variant [PublicNumOccurrencesRefineBy].
	MaxOccurrences int64 `json:"maxOccurrences"`
	// This field is from variant [PublicNumOccurrencesRefineBy].
	MinOccurrences int64 `json:"minOccurrences"`
	// This field is from variant [PublicSetOccurrencesRefineBy].
	SetType    string `json:"setType"`
	Comparison string `json:"comparison"`
	// This field is from variant [PublicRelativeComparativeTimestampRefineBy].
	TimeOffset PublicTimeOffset `json:"timeOffset"`
	// This field is from variant [PublicRelativeRangedTimestampRefineBy].
	LowerBoundOffset PublicTimeOffset `json:"lowerBoundOffset"`
	RangeType        string           `json:"rangeType"`
	// This field is from variant [PublicRelativeRangedTimestampRefineBy].
	UpperBoundOffset PublicTimeOffset `json:"upperBoundOffset"`
	// This field is from variant [PublicAbsoluteComparativeTimestampRefineBy].
	Timestamp int64 `json:"timestamp"`
	// This field is from variant [PublicAbsoluteRangedTimestampRefineBy].
	LowerTimestamp int64 `json:"lowerTimestamp"`
	// This field is from variant [PublicAbsoluteRangedTimestampRefineBy].
	UpperTimestamp               int64  `json:"upperTimestamp"`
	IncludeObjectsWithNoValueSet bool   `json:"includeObjectsWithNoValueSet"`
	OperationType                string `json:"operationType"`
	Operator                     string `json:"operator"`
	// This field is from variant [PublicTimePointOperation].
	TimePoint PublicTimePointOperationTimePointUnion `json:"timePoint"`
	// This field is from variant [PublicTimePointOperation].
	EndpointBehavior string `json:"endpointBehavior"`
	PropertyParser   string `json:"propertyParser"`
	// This field is from variant [PublicRangedTimeOperation].
	LowerBoundTimePoint PublicRangedTimeOperationLowerBoundTimePointUnion `json:"lowerBoundTimePoint"`
	// This field is from variant [PublicRangedTimeOperation].
	UpperBoundTimePoint PublicRangedTimeOperationUpperBoundTimePointUnion `json:"upperBoundTimePoint"`
	// This field is from variant [PublicRangedTimeOperation].
	LowerBoundEndpointBehavior string `json:"lowerBoundEndpointBehavior"`
	// This field is from variant [PublicRangedTimeOperation].
	UpperBoundEndpointBehavior string `json:"upperBoundEndpointBehavior"`
	JSON                       struct {
		Type                         respjson.Field
		MaxOccurrences               respjson.Field
		MinOccurrences               respjson.Field
		SetType                      respjson.Field
		Comparison                   respjson.Field
		TimeOffset                   respjson.Field
		LowerBoundOffset             respjson.Field
		RangeType                    respjson.Field
		UpperBoundOffset             respjson.Field
		Timestamp                    respjson.Field
		LowerTimestamp               respjson.Field
		UpperTimestamp               respjson.Field
		IncludeObjectsWithNoValueSet respjson.Field
		OperationType                respjson.Field
		Operator                     respjson.Field
		TimePoint                    respjson.Field
		EndpointBehavior             respjson.Field
		PropertyParser               respjson.Field
		LowerBoundTimePoint          respjson.Field
		UpperBoundTimePoint          respjson.Field
		LowerBoundEndpointBehavior   respjson.Field
		UpperBoundEndpointBehavior   respjson.Field
		raw                          string
	} `json:"-"`
}

func (u PublicAdsTimeFilterPruningRefineByUnion) AsNumOccurrences() (v PublicNumOccurrencesRefineBy) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicAdsTimeFilterPruningRefineByUnion) AsSetOccurrences() (v PublicSetOccurrencesRefineBy) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicAdsTimeFilterPruningRefineByUnion) AsRelativeComparative() (v PublicRelativeComparativeTimestampRefineBy) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicAdsTimeFilterPruningRefineByUnion) AsRelativeRanged() (v PublicRelativeRangedTimestampRefineBy) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicAdsTimeFilterPruningRefineByUnion) AsAbsoluteComparative() (v PublicAbsoluteComparativeTimestampRefineBy) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicAdsTimeFilterPruningRefineByUnion) AsAbsoluteRanged() (v PublicAbsoluteRangedTimestampRefineBy) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicAdsTimeFilterPruningRefineByUnion) AsAllHistory() (v PublicAllHistoryRefineBy) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicAdsTimeFilterPruningRefineByUnion) AsTimePoint() (v PublicTimePointOperation) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicAdsTimeFilterPruningRefineByUnion) AsTimeRanged() (v PublicRangedTimeOperation) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u PublicAdsTimeFilterPruningRefineByUnion) RawJSON() string { return u.JSON.raw }

func (r *PublicAdsTimeFilterPruningRefineByUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties FilterType, PruningRefineBy are required.
type PublicAdsTimeFilterParam struct {
	// Filter type (ADS_TIME)
	//
	// Any of "ADS_TIME".
	FilterType PublicAdsTimeFilterFilterType `json:"filterType,omitzero" api:"required"`
	// Refinement criteria
	PruningRefineBy PublicAdsTimeFilterPruningRefineByUnionParam `json:"pruningRefineBy,omitzero" api:"required"`
	paramObj
}

func (r PublicAdsTimeFilterParam) MarshalJSON() (data []byte, err error) {
	type shadow PublicAdsTimeFilterParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PublicAdsTimeFilterParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type PublicAdsTimeFilterPruningRefineByUnionParam struct {
	OfNumOccurrences      *PublicNumOccurrencesRefineByParam               `json:",omitzero,inline"`
	OfSetOccurrences      *PublicSetOccurrencesRefineByParam               `json:",omitzero,inline"`
	OfRelativeComparative *PublicRelativeComparativeTimestampRefineByParam `json:",omitzero,inline"`
	OfRelativeRanged      *PublicRelativeRangedTimestampRefineByParam      `json:",omitzero,inline"`
	OfAbsoluteComparative *PublicAbsoluteComparativeTimestampRefineByParam `json:",omitzero,inline"`
	OfAbsoluteRanged      *PublicAbsoluteRangedTimestampRefineByParam      `json:",omitzero,inline"`
	OfAllHistory          *PublicAllHistoryRefineByParam                   `json:",omitzero,inline"`
	OfTimePoint           *PublicTimePointOperationParam                   `json:",omitzero,inline"`
	OfTimeRanged          *PublicRangedTimeOperationParam                  `json:",omitzero,inline"`
	paramUnion
}

func (u PublicAdsTimeFilterPruningRefineByUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfNumOccurrences,
		u.OfSetOccurrences,
		u.OfRelativeComparative,
		u.OfRelativeRanged,
		u.OfAbsoluteComparative,
		u.OfAbsoluteRanged,
		u.OfAllHistory,
		u.OfTimePoint,
		u.OfTimeRanged)
}
func (u *PublicAdsTimeFilterPruningRefineByUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

type PublicAllHistoryRefineBy struct {
	// Type of refine by (ALL_HISTORY)
	//
	// Any of "ALL_HISTORY".
	Type PublicAllHistoryRefineByType `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PublicAllHistoryRefineBy) RawJSON() string { return r.JSON.raw }
func (r *PublicAllHistoryRefineBy) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this PublicAllHistoryRefineBy to a
// PublicAllHistoryRefineByParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// PublicAllHistoryRefineByParam.Overrides()
func (r PublicAllHistoryRefineBy) ToParam() PublicAllHistoryRefineByParam {
	return param.Override[PublicAllHistoryRefineByParam](json.RawMessage(r.RawJSON()))
}

// Type of refine by (ALL_HISTORY)
type PublicAllHistoryRefineByType string

const (
	PublicAllHistoryRefineByTypeAllHistory PublicAllHistoryRefineByType = "ALL_HISTORY"
)

// The property Type is required.
type PublicAllHistoryRefineByParam struct {
	// Type of refine by (ALL_HISTORY)
	//
	// Any of "ALL_HISTORY".
	Type PublicAllHistoryRefineByType `json:"type,omitzero" api:"required"`
	paramObj
}

func (r PublicAllHistoryRefineByParam) MarshalJSON() (data []byte, err error) {
	type shadow PublicAllHistoryRefineByParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PublicAllHistoryRefineByParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PublicAllPropertyTypesOperation struct {
	// Indication of whether objects with no value should be included
	IncludeObjectsWithNoValueSet bool `json:"includeObjectsWithNoValueSet" api:"required"`
	// Type of operation (ALL_PROPERTY)
	//
	// Any of "ALL_PROPERTY".
	OperationType PublicAllPropertyTypesOperationOperationType `json:"operationType" api:"required"`
	// Operator to be applied (IS_KNOWN, IS_UNKNOWN)
	Operator string `json:"operator" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		IncludeObjectsWithNoValueSet respjson.Field
		OperationType                respjson.Field
		Operator                     respjson.Field
		ExtraFields                  map[string]respjson.Field
		raw                          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PublicAllPropertyTypesOperation) RawJSON() string { return r.JSON.raw }
func (r *PublicAllPropertyTypesOperation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this PublicAllPropertyTypesOperation to a
// PublicAllPropertyTypesOperationParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// PublicAllPropertyTypesOperationParam.Overrides()
func (r PublicAllPropertyTypesOperation) ToParam() PublicAllPropertyTypesOperationParam {
	return param.Override[PublicAllPropertyTypesOperationParam](json.RawMessage(r.RawJSON()))
}

// Type of operation (ALL_PROPERTY)
type PublicAllPropertyTypesOperationOperationType string

const (
	PublicAllPropertyTypesOperationOperationTypeAllProperty PublicAllPropertyTypesOperationOperationType = "ALL_PROPERTY"
)

// The properties IncludeObjectsWithNoValueSet, OperationType, Operator are
// required.
type PublicAllPropertyTypesOperationParam struct {
	// Indication of whether objects with no value should be included
	IncludeObjectsWithNoValueSet bool `json:"includeObjectsWithNoValueSet" api:"required"`
	// Type of operation (ALL_PROPERTY)
	//
	// Any of "ALL_PROPERTY".
	OperationType PublicAllPropertyTypesOperationOperationType `json:"operationType,omitzero" api:"required"`
	// Operator to be applied (IS_KNOWN, IS_UNKNOWN)
	Operator string `json:"operator" api:"required"`
	paramObj
}

func (r PublicAllPropertyTypesOperationParam) MarshalJSON() (data []byte, err error) {
	type shadow PublicAllPropertyTypesOperationParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PublicAllPropertyTypesOperationParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PublicAndFilterBranch struct {
	FilterBranches []PublicAndFilterBranchFilterBranchUnion `json:"filterBranches" api:"required"`
	// Filter branch operator (AND)
	FilterBranchOperator string `json:"filterBranchOperator" api:"required"`
	// Type of filter branch (AND)
	//
	// Any of "AND".
	FilterBranchType PublicAndFilterBranchFilterBranchType `json:"filterBranchType" api:"required"`
	Filters          []PublicAndFilterBranchFilterUnion    `json:"filters" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FilterBranches       respjson.Field
		FilterBranchOperator respjson.Field
		FilterBranchType     respjson.Field
		Filters              respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PublicAndFilterBranch) RawJSON() string { return r.JSON.raw }
func (r *PublicAndFilterBranch) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this PublicAndFilterBranch to a PublicAndFilterBranchParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// PublicAndFilterBranchParam.Overrides()
func (r PublicAndFilterBranch) ToParam() PublicAndFilterBranchParam {
	return param.Override[PublicAndFilterBranchParam](json.RawMessage(r.RawJSON()))
}

// PublicAndFilterBranchFilterBranchUnion contains all possible properties and
// values from [PublicOrFilterBranch], [PublicAndFilterBranch],
// [PublicNotAllFilterBranch], [PublicNotAnyFilterBranch],
// [PublicRestrictedFilterBranch], [PublicUnifiedEventsFilterBranch],
// [PublicPropertyAssociationFilterBranch], [PublicAssociationFilterBranch].
//
// Use the [PublicAndFilterBranchFilterBranchUnion.AsAny] method to switch on the
// variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type PublicAndFilterBranchFilterBranchUnion struct {
	// This field is a union of [[]PublicOrFilterBranchFilterBranchUnion],
	// [[]PublicAndFilterBranchFilterBranchUnion],
	// [[]PublicNotAllFilterBranchFilterBranchUnion],
	// [[]PublicNotAnyFilterBranchFilterBranchUnion],
	// [[]PublicRestrictedFilterBranchFilterBranchUnion],
	// [[]PublicUnifiedEventsFilterBranchFilterBranchUnion],
	// [[]PublicPropertyAssociationFilterBranchFilterBranchUnion],
	// [[]PublicAssociationFilterBranchFilterBranchUnion]
	FilterBranches       PublicAndFilterBranchFilterBranchUnionFilterBranches `json:"filterBranches"`
	FilterBranchOperator string                                               `json:"filterBranchOperator"`
	// Any of "OR", "AND", "NOT_ALL", "NOT_ANY", "RESTRICTED", "UNIFIED_EVENTS",
	// "PROPERTY_ASSOCIATION", "ASSOCIATION".
	FilterBranchType string `json:"filterBranchType"`
	// This field is a union of [[]PublicOrFilterBranchFilterUnion],
	// [[]PublicAndFilterBranchFilterUnion], [[]PublicNotAllFilterBranchFilterUnion],
	// [[]PublicNotAnyFilterBranchFilterUnion],
	// [[]PublicRestrictedFilterBranchFilterUnion],
	// [[]PublicUnifiedEventsFilterBranchFilterUnion],
	// [[]PublicPropertyAssociationFilterBranchFilterUnion],
	// [[]PublicAssociationFilterBranchFilterUnion]
	Filters PublicAndFilterBranchFilterBranchUnionFilters `json:"filters"`
	// This field is from variant [PublicUnifiedEventsFilterBranch].
	EventTypeID string `json:"eventTypeId"`
	Operator    string `json:"operator"`
	// This field is from variant [PublicUnifiedEventsFilterBranch].
	CoalescingRefineBy PublicUnifiedEventsFilterBranchCoalescingRefineByUnion `json:"coalescingRefineBy"`
	// This field is from variant [PublicUnifiedEventsFilterBranch].
	PruningRefineBy PublicUnifiedEventsFilterBranchPruningRefineByUnion `json:"pruningRefineBy"`
	ObjectTypeID    string                                              `json:"objectTypeId"`
	// This field is from variant [PublicPropertyAssociationFilterBranch].
	PropertyWithObjectID string `json:"propertyWithObjectId"`
	// This field is from variant [PublicAssociationFilterBranch].
	AssociationCategory string `json:"associationCategory"`
	// This field is from variant [PublicAssociationFilterBranch].
	AssociationTypeID int64 `json:"associationTypeId"`
	JSON              struct {
		FilterBranches       respjson.Field
		FilterBranchOperator respjson.Field
		FilterBranchType     respjson.Field
		Filters              respjson.Field
		EventTypeID          respjson.Field
		Operator             respjson.Field
		CoalescingRefineBy   respjson.Field
		PruningRefineBy      respjson.Field
		ObjectTypeID         respjson.Field
		PropertyWithObjectID respjson.Field
		AssociationCategory  respjson.Field
		AssociationTypeID    respjson.Field
		raw                  string
	} `json:"-"`
}

// anyPublicAndFilterBranchFilterBranch is implemented by each variant of
// [PublicAndFilterBranchFilterBranchUnion] to add type safety for the return type
// of [PublicAndFilterBranchFilterBranchUnion.AsAny]
type anyPublicAndFilterBranchFilterBranch interface {
	implPublicAndFilterBranchFilterBranchUnion()
}

func (PublicOrFilterBranch) implPublicAndFilterBranchFilterBranchUnion()                  {}
func (PublicAndFilterBranch) implPublicAndFilterBranchFilterBranchUnion()                 {}
func (PublicNotAllFilterBranch) implPublicAndFilterBranchFilterBranchUnion()              {}
func (PublicNotAnyFilterBranch) implPublicAndFilterBranchFilterBranchUnion()              {}
func (PublicRestrictedFilterBranch) implPublicAndFilterBranchFilterBranchUnion()          {}
func (PublicUnifiedEventsFilterBranch) implPublicAndFilterBranchFilterBranchUnion()       {}
func (PublicPropertyAssociationFilterBranch) implPublicAndFilterBranchFilterBranchUnion() {}
func (PublicAssociationFilterBranch) implPublicAndFilterBranchFilterBranchUnion()         {}

// Use the following switch statement to find the correct variant
//
//	switch variant := PublicAndFilterBranchFilterBranchUnion.AsAny().(type) {
//	case crm.PublicOrFilterBranch:
//	case crm.PublicAndFilterBranch:
//	case crm.PublicNotAllFilterBranch:
//	case crm.PublicNotAnyFilterBranch:
//	case crm.PublicRestrictedFilterBranch:
//	case crm.PublicUnifiedEventsFilterBranch:
//	case crm.PublicPropertyAssociationFilterBranch:
//	case crm.PublicAssociationFilterBranch:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u PublicAndFilterBranchFilterBranchUnion) AsAny() anyPublicAndFilterBranchFilterBranch {
	switch u.FilterBranchType {
	case "OR":
		return u.AsOr()
	case "AND":
		return u.AsAnd()
	case "NOT_ALL":
		return u.AsNotAll()
	case "NOT_ANY":
		return u.AsNotAny()
	case "RESTRICTED":
		return u.AsRestricted()
	case "UNIFIED_EVENTS":
		return u.AsUnifiedEvents()
	case "PROPERTY_ASSOCIATION":
		return u.AsPropertyAssociation()
	case "ASSOCIATION":
		return u.AsAssociation()
	}
	return nil
}

func (u PublicAndFilterBranchFilterBranchUnion) AsOr() (v PublicOrFilterBranch) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicAndFilterBranchFilterBranchUnion) AsAnd() (v PublicAndFilterBranch) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicAndFilterBranchFilterBranchUnion) AsNotAll() (v PublicNotAllFilterBranch) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicAndFilterBranchFilterBranchUnion) AsNotAny() (v PublicNotAnyFilterBranch) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicAndFilterBranchFilterBranchUnion) AsRestricted() (v PublicRestrictedFilterBranch) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicAndFilterBranchFilterBranchUnion) AsUnifiedEvents() (v PublicUnifiedEventsFilterBranch) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicAndFilterBranchFilterBranchUnion) AsPropertyAssociation() (v PublicPropertyAssociationFilterBranch) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicAndFilterBranchFilterBranchUnion) AsAssociation() (v PublicAssociationFilterBranch) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u PublicAndFilterBranchFilterBranchUnion) RawJSON() string { return u.JSON.raw }

func (r *PublicAndFilterBranchFilterBranchUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PublicAndFilterBranchFilterBranchUnionFilterBranches is an implicit subunion of
// [PublicAndFilterBranchFilterBranchUnion].
// PublicAndFilterBranchFilterBranchUnionFilterBranches provides convenient access
// to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [PublicAndFilterBranchFilterBranchUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfPublicOrFilterBranchFilterBranches
// OfPublicAndFilterBranchFilterBranches OfPublicNotAllFilterBranchFilterBranches
// OfPublicNotAnyFilterBranchFilterBranches
// OfPublicRestrictedFilterBranchFilterBranches
// OfPublicUnifiedEventsFilterBranchFilterBranches
// OfPublicPropertyAssociationFilterBranchFilterBranches
// OfPublicAssociationFilterBranchFilterBranches]
type PublicAndFilterBranchFilterBranchUnionFilterBranches struct {
	// This field will be present if the value is a
	// [[]PublicOrFilterBranchFilterBranchUnion] instead of an object.
	OfPublicOrFilterBranchFilterBranches []PublicOrFilterBranchFilterBranchUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]PublicAndFilterBranchFilterBranchUnion] instead of an object.
	OfPublicAndFilterBranchFilterBranches []PublicAndFilterBranchFilterBranchUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]PublicNotAllFilterBranchFilterBranchUnion] instead of an object.
	OfPublicNotAllFilterBranchFilterBranches []PublicNotAllFilterBranchFilterBranchUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]PublicNotAnyFilterBranchFilterBranchUnion] instead of an object.
	OfPublicNotAnyFilterBranchFilterBranches []PublicNotAnyFilterBranchFilterBranchUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]PublicRestrictedFilterBranchFilterBranchUnion] instead of an object.
	OfPublicRestrictedFilterBranchFilterBranches []PublicRestrictedFilterBranchFilterBranchUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]PublicUnifiedEventsFilterBranchFilterBranchUnion] instead of an object.
	OfPublicUnifiedEventsFilterBranchFilterBranches []PublicUnifiedEventsFilterBranchFilterBranchUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]PublicPropertyAssociationFilterBranchFilterBranchUnion] instead of an object.
	OfPublicPropertyAssociationFilterBranchFilterBranches []PublicPropertyAssociationFilterBranchFilterBranchUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]PublicAssociationFilterBranchFilterBranchUnion] instead of an object.
	OfPublicAssociationFilterBranchFilterBranches []PublicAssociationFilterBranchFilterBranchUnion `json:",inline"`
	JSON                                          struct {
		OfPublicOrFilterBranchFilterBranches                  respjson.Field
		OfPublicAndFilterBranchFilterBranches                 respjson.Field
		OfPublicNotAllFilterBranchFilterBranches              respjson.Field
		OfPublicNotAnyFilterBranchFilterBranches              respjson.Field
		OfPublicRestrictedFilterBranchFilterBranches          respjson.Field
		OfPublicUnifiedEventsFilterBranchFilterBranches       respjson.Field
		OfPublicPropertyAssociationFilterBranchFilterBranches respjson.Field
		OfPublicAssociationFilterBranchFilterBranches         respjson.Field
		raw                                                   string
	} `json:"-"`
}

func (r *PublicAndFilterBranchFilterBranchUnionFilterBranches) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PublicAndFilterBranchFilterBranchUnionFilters is an implicit subunion of
// [PublicAndFilterBranchFilterBranchUnion].
// PublicAndFilterBranchFilterBranchUnionFilters provides convenient access to the
// sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [PublicAndFilterBranchFilterBranchUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfPublicOrFilterBranchFilters OfPublicAndFilterBranchFilters
// OfPublicNotAllFilterBranchFilters OfPublicNotAnyFilterBranchFilters
// OfPublicRestrictedFilterBranchFilters OfPublicUnifiedEventsFilterBranchFilters
// OfPublicPropertyAssociationFilterBranchFilters
// OfPublicAssociationFilterBranchFilters]
type PublicAndFilterBranchFilterBranchUnionFilters struct {
	// This field will be present if the value is a [[]PublicOrFilterBranchFilterUnion]
	// instead of an object.
	OfPublicOrFilterBranchFilters []PublicOrFilterBranchFilterUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]PublicAndFilterBranchFilterUnion] instead of an object.
	OfPublicAndFilterBranchFilters []PublicAndFilterBranchFilterUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]PublicNotAllFilterBranchFilterUnion] instead of an object.
	OfPublicNotAllFilterBranchFilters []PublicNotAllFilterBranchFilterUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]PublicNotAnyFilterBranchFilterUnion] instead of an object.
	OfPublicNotAnyFilterBranchFilters []PublicNotAnyFilterBranchFilterUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]PublicRestrictedFilterBranchFilterUnion] instead of an object.
	OfPublicRestrictedFilterBranchFilters []PublicRestrictedFilterBranchFilterUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]PublicUnifiedEventsFilterBranchFilterUnion] instead of an object.
	OfPublicUnifiedEventsFilterBranchFilters []PublicUnifiedEventsFilterBranchFilterUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]PublicPropertyAssociationFilterBranchFilterUnion] instead of an object.
	OfPublicPropertyAssociationFilterBranchFilters []PublicPropertyAssociationFilterBranchFilterUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]PublicAssociationFilterBranchFilterUnion] instead of an object.
	OfPublicAssociationFilterBranchFilters []PublicAssociationFilterBranchFilterUnion `json:",inline"`
	JSON                                   struct {
		OfPublicOrFilterBranchFilters                  respjson.Field
		OfPublicAndFilterBranchFilters                 respjson.Field
		OfPublicNotAllFilterBranchFilters              respjson.Field
		OfPublicNotAnyFilterBranchFilters              respjson.Field
		OfPublicRestrictedFilterBranchFilters          respjson.Field
		OfPublicUnifiedEventsFilterBranchFilters       respjson.Field
		OfPublicPropertyAssociationFilterBranchFilters respjson.Field
		OfPublicAssociationFilterBranchFilters         respjson.Field
		raw                                            string
	} `json:"-"`
}

func (r *PublicAndFilterBranchFilterBranchUnionFilters) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Type of filter branch (AND)
type PublicAndFilterBranchFilterBranchType string

const (
	PublicAndFilterBranchFilterBranchTypeAnd PublicAndFilterBranchFilterBranchType = "AND"
)

// PublicAndFilterBranchFilterUnion contains all possible properties and values
// from [PublicPropertyFilter], [PublicAssociationInListFilter],
// [PublicPageViewAnalyticsFilter], [PublicCtaAnalyticsFilter],
// [PublicEventAnalyticsFilter], [PublicFormSubmissionFilter],
// [PublicFormSubmissionOnPageFilter], [PublicIntegrationEventFilter],
// [PublicEmailSubscriptionFilter], [PublicCommunicationSubscriptionFilter],
// [PublicCampaignInfluencedFilter], [PublicSurveyMonkeyFilter],
// [PublicSurveyMonkeyValueFilter], [PublicWebinarFilter],
// [PublicEmailEventFilter], [PublicPrivacyAnalyticsFilter],
// [PublicAdsSearchFilter], [PublicAdsTimeFilter], [PublicInListFilter],
// [PublicNumAssociationsFilter], [PublicUnifiedEventsFilter],
// [PublicPropertyAssociationInListFilter], [PublicConstantFilter].
//
// Use the [PublicAndFilterBranchFilterUnion.AsAny] method to switch on the
// variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type PublicAndFilterBranchFilterUnion struct {
	// Any of "PROPERTY", "ASSOCIATION", "PAGE_VIEW", "CTA", "EVENT",
	// "FORM_SUBMISSION", "FORM_SUBMISSION_ON_PAGE", "INTEGRATION_EVENT",
	// "EMAIL_SUBSCRIPTION", "COMMUNICATION_SUBSCRIPTION", "CAMPAIGN_INFLUENCED",
	// "SURVEY_MONKEY", "SURVEY_MONKEY_VALUE", "WEBINAR", "EMAIL_EVENT", "PRIVACY",
	// "ADS_SEARCH", "ADS_TIME", "IN_LIST", "NUM_ASSOCIATIONS", "UNIFIED_EVENTS",
	// "PROPERTY_ASSOCIATION", "CONSTANT".
	FilterType string `json:"filterType"`
	// This field is from variant [PublicPropertyFilter].
	Operation PublicPropertyFilterOperationUnion `json:"operation"`
	// This field is from variant [PublicPropertyFilter].
	Property            string `json:"property"`
	AssociationCategory string `json:"associationCategory"`
	AssociationTypeID   int64  `json:"associationTypeId"`
	// This field is a union of [PublicAssociationInListFilterCoalescingRefineByUnion],
	// [PublicPageViewAnalyticsFilterCoalescingRefineByUnion],
	// [PublicCtaAnalyticsFilterCoalescingRefineByUnion],
	// [PublicEventAnalyticsFilterCoalescingRefineByUnion],
	// [PublicFormSubmissionFilterCoalescingRefineByUnion],
	// [PublicFormSubmissionOnPageFilterCoalescingRefineByUnion],
	// [PublicNumAssociationsFilterCoalescingRefineByUnion],
	// [PublicUnifiedEventsFilterCoalescingRefineByUnion],
	// [PublicPropertyAssociationInListFilterCoalescingRefineByUnion]
	CoalescingRefineBy PublicAndFilterBranchFilterUnionCoalescingRefineBy `json:"coalescingRefineBy"`
	ListID             string                                             `json:"listId"`
	Operator           string                                             `json:"operator"`
	// This field is from variant [PublicAssociationInListFilter].
	ToObjectType   string `json:"toObjectType"`
	ToObjectTypeID string `json:"toObjectTypeId"`
	// This field is from variant [PublicPageViewAnalyticsFilter].
	PageURL string `json:"pageUrl"`
	// This field is from variant [PublicPageViewAnalyticsFilter].
	EnableTracking bool `json:"enableTracking"`
	// This field is a union of [PublicPageViewAnalyticsFilterPruningRefineByUnion],
	// [PublicCtaAnalyticsFilterPruningRefineByUnion],
	// [PublicEventAnalyticsFilterPruningRefineByUnion],
	// [PublicFormSubmissionFilterPruningRefineByUnion],
	// [PublicFormSubmissionOnPageFilterPruningRefineByUnion],
	// [PublicEmailEventFilterPruningRefineByUnion],
	// [PublicAdsTimeFilterPruningRefineByUnion],
	// [PublicUnifiedEventsFilterPruningRefineByUnion]
	PruningRefineBy PublicAndFilterBranchFilterUnionPruningRefineBy `json:"pruningRefineBy"`
	// This field is from variant [PublicCtaAnalyticsFilter].
	CtaName string `json:"ctaName"`
	// This field is from variant [PublicEventAnalyticsFilter].
	EventID string `json:"eventId"`
	FormID  string `json:"formId"`
	// This field is from variant [PublicFormSubmissionOnPageFilter].
	PageID string `json:"pageId"`
	// This field is a union of [int64], [string]
	EventTypeID PublicAndFilterBranchFilterUnionEventTypeID `json:"eventTypeId"`
	FilterLines []PublicEventFilterMetadata                 `json:"filterLines"`
	// This field is from variant [PublicEmailSubscriptionFilter].
	AcceptedStatuses []string `json:"acceptedStatuses"`
	SubscriptionIDs  []string `json:"subscriptionIds"`
	SubscriptionType string   `json:"subscriptionType"`
	// This field is from variant [PublicCommunicationSubscriptionFilter].
	AcceptedOptStates []string `json:"acceptedOptStates"`
	// This field is from variant [PublicCommunicationSubscriptionFilter].
	Channel string `json:"channel"`
	// This field is from variant [PublicCommunicationSubscriptionFilter].
	BusinessUnitID string `json:"businessUnitId"`
	// This field is from variant [PublicCampaignInfluencedFilter].
	CampaignID string `json:"campaignId"`
	SurveyID   string `json:"surveyId"`
	// This field is from variant [PublicSurveyMonkeyValueFilter].
	SurveyQuestion string `json:"surveyQuestion"`
	// This field is from variant [PublicSurveyMonkeyValueFilter].
	ValueComparison PublicSurveyMonkeyValueFilterValueComparisonUnion `json:"valueComparison"`
	// This field is from variant [PublicSurveyMonkeyValueFilter].
	SurveyAnswerColID string `json:"surveyAnswerColId"`
	// This field is from variant [PublicSurveyMonkeyValueFilter].
	SurveyAnswerRowID string `json:"surveyAnswerRowId"`
	// This field is from variant [PublicWebinarFilter].
	WebinarID string `json:"webinarId"`
	// This field is from variant [PublicEmailEventFilter].
	AppID string `json:"appId"`
	// This field is from variant [PublicEmailEventFilter].
	EmailID string `json:"emailId"`
	// This field is from variant [PublicEmailEventFilter].
	Level string `json:"level"`
	// This field is from variant [PublicEmailEventFilter].
	ClickURL string `json:"clickUrl"`
	// This field is from variant [PublicPrivacyAnalyticsFilter].
	PrivacyName string `json:"privacyName"`
	// This field is from variant [PublicAdsSearchFilter].
	AdNetwork string `json:"adNetwork"`
	// This field is from variant [PublicAdsSearchFilter].
	EntityType string `json:"entityType"`
	// This field is from variant [PublicAdsSearchFilter].
	SearchTerms []string `json:"searchTerms"`
	// This field is from variant [PublicAdsSearchFilter].
	SearchTermType string `json:"searchTermType"`
	// This field is from variant [PublicInListFilter].
	Metadata PublicInListFilterMetadata `json:"metadata"`
	// This field is from variant [PublicPropertyAssociationInListFilter].
	PropertyWithObjectID string `json:"propertyWithObjectId"`
	// This field is from variant [PublicConstantFilter].
	ShouldAccept bool `json:"shouldAccept"`
	// This field is from variant [PublicConstantFilter].
	Source string `json:"source"`
	JSON   struct {
		FilterType           respjson.Field
		Operation            respjson.Field
		Property             respjson.Field
		AssociationCategory  respjson.Field
		AssociationTypeID    respjson.Field
		CoalescingRefineBy   respjson.Field
		ListID               respjson.Field
		Operator             respjson.Field
		ToObjectType         respjson.Field
		ToObjectTypeID       respjson.Field
		PageURL              respjson.Field
		EnableTracking       respjson.Field
		PruningRefineBy      respjson.Field
		CtaName              respjson.Field
		EventID              respjson.Field
		FormID               respjson.Field
		PageID               respjson.Field
		EventTypeID          respjson.Field
		FilterLines          respjson.Field
		AcceptedStatuses     respjson.Field
		SubscriptionIDs      respjson.Field
		SubscriptionType     respjson.Field
		AcceptedOptStates    respjson.Field
		Channel              respjson.Field
		BusinessUnitID       respjson.Field
		CampaignID           respjson.Field
		SurveyID             respjson.Field
		SurveyQuestion       respjson.Field
		ValueComparison      respjson.Field
		SurveyAnswerColID    respjson.Field
		SurveyAnswerRowID    respjson.Field
		WebinarID            respjson.Field
		AppID                respjson.Field
		EmailID              respjson.Field
		Level                respjson.Field
		ClickURL             respjson.Field
		PrivacyName          respjson.Field
		AdNetwork            respjson.Field
		EntityType           respjson.Field
		SearchTerms          respjson.Field
		SearchTermType       respjson.Field
		Metadata             respjson.Field
		PropertyWithObjectID respjson.Field
		ShouldAccept         respjson.Field
		Source               respjson.Field
		raw                  string
	} `json:"-"`
}

// anyPublicAndFilterBranchFilter is implemented by each variant of
// [PublicAndFilterBranchFilterUnion] to add type safety for the return type of
// [PublicAndFilterBranchFilterUnion.AsAny]
type anyPublicAndFilterBranchFilter interface {
	implPublicAndFilterBranchFilterUnion()
}

func (PublicPropertyFilter) implPublicAndFilterBranchFilterUnion()                  {}
func (PublicAssociationInListFilter) implPublicAndFilterBranchFilterUnion()         {}
func (PublicPageViewAnalyticsFilter) implPublicAndFilterBranchFilterUnion()         {}
func (PublicCtaAnalyticsFilter) implPublicAndFilterBranchFilterUnion()              {}
func (PublicEventAnalyticsFilter) implPublicAndFilterBranchFilterUnion()            {}
func (PublicFormSubmissionFilter) implPublicAndFilterBranchFilterUnion()            {}
func (PublicFormSubmissionOnPageFilter) implPublicAndFilterBranchFilterUnion()      {}
func (PublicIntegrationEventFilter) implPublicAndFilterBranchFilterUnion()          {}
func (PublicEmailSubscriptionFilter) implPublicAndFilterBranchFilterUnion()         {}
func (PublicCommunicationSubscriptionFilter) implPublicAndFilterBranchFilterUnion() {}
func (PublicCampaignInfluencedFilter) implPublicAndFilterBranchFilterUnion()        {}
func (PublicSurveyMonkeyFilter) implPublicAndFilterBranchFilterUnion()              {}
func (PublicSurveyMonkeyValueFilter) implPublicAndFilterBranchFilterUnion()         {}
func (PublicWebinarFilter) implPublicAndFilterBranchFilterUnion()                   {}
func (PublicEmailEventFilter) implPublicAndFilterBranchFilterUnion()                {}
func (PublicPrivacyAnalyticsFilter) implPublicAndFilterBranchFilterUnion()          {}
func (PublicAdsSearchFilter) implPublicAndFilterBranchFilterUnion()                 {}
func (PublicAdsTimeFilter) implPublicAndFilterBranchFilterUnion()                   {}
func (PublicInListFilter) implPublicAndFilterBranchFilterUnion()                    {}
func (PublicNumAssociationsFilter) implPublicAndFilterBranchFilterUnion()           {}
func (PublicUnifiedEventsFilter) implPublicAndFilterBranchFilterUnion()             {}
func (PublicPropertyAssociationInListFilter) implPublicAndFilterBranchFilterUnion() {}
func (PublicConstantFilter) implPublicAndFilterBranchFilterUnion()                  {}

// Use the following switch statement to find the correct variant
//
//	switch variant := PublicAndFilterBranchFilterUnion.AsAny().(type) {
//	case crm.PublicPropertyFilter:
//	case crm.PublicAssociationInListFilter:
//	case crm.PublicPageViewAnalyticsFilter:
//	case crm.PublicCtaAnalyticsFilter:
//	case crm.PublicEventAnalyticsFilter:
//	case crm.PublicFormSubmissionFilter:
//	case crm.PublicFormSubmissionOnPageFilter:
//	case crm.PublicIntegrationEventFilter:
//	case crm.PublicEmailSubscriptionFilter:
//	case crm.PublicCommunicationSubscriptionFilter:
//	case crm.PublicCampaignInfluencedFilter:
//	case crm.PublicSurveyMonkeyFilter:
//	case crm.PublicSurveyMonkeyValueFilter:
//	case crm.PublicWebinarFilter:
//	case crm.PublicEmailEventFilter:
//	case crm.PublicPrivacyAnalyticsFilter:
//	case crm.PublicAdsSearchFilter:
//	case crm.PublicAdsTimeFilter:
//	case crm.PublicInListFilter:
//	case crm.PublicNumAssociationsFilter:
//	case crm.PublicUnifiedEventsFilter:
//	case crm.PublicPropertyAssociationInListFilter:
//	case crm.PublicConstantFilter:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u PublicAndFilterBranchFilterUnion) AsAny() anyPublicAndFilterBranchFilter {
	switch u.FilterType {
	case "PROPERTY":
		return u.AsProperty()
	case "ASSOCIATION":
		return u.AsAssociation()
	case "PAGE_VIEW":
		return u.AsPageView()
	case "CTA":
		return u.AsCta()
	case "EVENT":
		return u.AsEvent()
	case "FORM_SUBMISSION":
		return u.AsFormSubmission()
	case "FORM_SUBMISSION_ON_PAGE":
		return u.AsFormSubmissionOnPage()
	case "INTEGRATION_EVENT":
		return u.AsIntegrationEvent()
	case "EMAIL_SUBSCRIPTION":
		return u.AsEmailSubscription()
	case "COMMUNICATION_SUBSCRIPTION":
		return u.AsCommunicationSubscription()
	case "CAMPAIGN_INFLUENCED":
		return u.AsCampaignInfluenced()
	case "SURVEY_MONKEY":
		return u.AsSurveyMonkey()
	case "SURVEY_MONKEY_VALUE":
		return u.AsSurveyMonkeyValue()
	case "WEBINAR":
		return u.AsWebinar()
	case "EMAIL_EVENT":
		return u.AsEmailEvent()
	case "PRIVACY":
		return u.AsPrivacy()
	case "ADS_SEARCH":
		return u.AsAdsSearch()
	case "ADS_TIME":
		return u.AsAdsTime()
	case "IN_LIST":
		return u.AsInList()
	case "NUM_ASSOCIATIONS":
		return u.AsNumAssociations()
	case "UNIFIED_EVENTS":
		return u.AsUnifiedEvents()
	case "PROPERTY_ASSOCIATION":
		return u.AsPropertyAssociation()
	case "CONSTANT":
		return u.AsConstant()
	}
	return nil
}

func (u PublicAndFilterBranchFilterUnion) AsProperty() (v PublicPropertyFilter) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicAndFilterBranchFilterUnion) AsAssociation() (v PublicAssociationInListFilter) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicAndFilterBranchFilterUnion) AsPageView() (v PublicPageViewAnalyticsFilter) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicAndFilterBranchFilterUnion) AsCta() (v PublicCtaAnalyticsFilter) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicAndFilterBranchFilterUnion) AsEvent() (v PublicEventAnalyticsFilter) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicAndFilterBranchFilterUnion) AsFormSubmission() (v PublicFormSubmissionFilter) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicAndFilterBranchFilterUnion) AsFormSubmissionOnPage() (v PublicFormSubmissionOnPageFilter) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicAndFilterBranchFilterUnion) AsIntegrationEvent() (v PublicIntegrationEventFilter) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicAndFilterBranchFilterUnion) AsEmailSubscription() (v PublicEmailSubscriptionFilter) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicAndFilterBranchFilterUnion) AsCommunicationSubscription() (v PublicCommunicationSubscriptionFilter) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicAndFilterBranchFilterUnion) AsCampaignInfluenced() (v PublicCampaignInfluencedFilter) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicAndFilterBranchFilterUnion) AsSurveyMonkey() (v PublicSurveyMonkeyFilter) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicAndFilterBranchFilterUnion) AsSurveyMonkeyValue() (v PublicSurveyMonkeyValueFilter) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicAndFilterBranchFilterUnion) AsWebinar() (v PublicWebinarFilter) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicAndFilterBranchFilterUnion) AsEmailEvent() (v PublicEmailEventFilter) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicAndFilterBranchFilterUnion) AsPrivacy() (v PublicPrivacyAnalyticsFilter) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicAndFilterBranchFilterUnion) AsAdsSearch() (v PublicAdsSearchFilter) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicAndFilterBranchFilterUnion) AsAdsTime() (v PublicAdsTimeFilter) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicAndFilterBranchFilterUnion) AsInList() (v PublicInListFilter) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicAndFilterBranchFilterUnion) AsNumAssociations() (v PublicNumAssociationsFilter) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicAndFilterBranchFilterUnion) AsUnifiedEvents() (v PublicUnifiedEventsFilter) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicAndFilterBranchFilterUnion) AsPropertyAssociation() (v PublicPropertyAssociationInListFilter) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicAndFilterBranchFilterUnion) AsConstant() (v PublicConstantFilter) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u PublicAndFilterBranchFilterUnion) RawJSON() string { return u.JSON.raw }

func (r *PublicAndFilterBranchFilterUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PublicAndFilterBranchFilterUnionCoalescingRefineBy is an implicit subunion of
// [PublicAndFilterBranchFilterUnion].
// PublicAndFilterBranchFilterUnionCoalescingRefineBy provides convenient access to
// the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [PublicAndFilterBranchFilterUnion].
type PublicAndFilterBranchFilterUnionCoalescingRefineBy struct {
	Type string `json:"type"`
	// This field is from variant
	// [PublicAssociationInListFilterCoalescingRefineByUnion],
	// [PublicPageViewAnalyticsFilterCoalescingRefineByUnion],
	// [PublicCtaAnalyticsFilterCoalescingRefineByUnion],
	// [PublicEventAnalyticsFilterCoalescingRefineByUnion],
	// [PublicFormSubmissionFilterCoalescingRefineByUnion],
	// [PublicFormSubmissionOnPageFilterCoalescingRefineByUnion],
	// [PublicNumAssociationsFilterCoalescingRefineByUnion],
	// [PublicUnifiedEventsFilterCoalescingRefineByUnion],
	// [PublicPropertyAssociationInListFilterCoalescingRefineByUnion].
	MaxOccurrences int64 `json:"maxOccurrences"`
	// This field is from variant
	// [PublicAssociationInListFilterCoalescingRefineByUnion],
	// [PublicPageViewAnalyticsFilterCoalescingRefineByUnion],
	// [PublicCtaAnalyticsFilterCoalescingRefineByUnion],
	// [PublicEventAnalyticsFilterCoalescingRefineByUnion],
	// [PublicFormSubmissionFilterCoalescingRefineByUnion],
	// [PublicFormSubmissionOnPageFilterCoalescingRefineByUnion],
	// [PublicNumAssociationsFilterCoalescingRefineByUnion],
	// [PublicUnifiedEventsFilterCoalescingRefineByUnion],
	// [PublicPropertyAssociationInListFilterCoalescingRefineByUnion].
	MinOccurrences int64 `json:"minOccurrences"`
	// This field is from variant
	// [PublicAssociationInListFilterCoalescingRefineByUnion],
	// [PublicPageViewAnalyticsFilterCoalescingRefineByUnion],
	// [PublicCtaAnalyticsFilterCoalescingRefineByUnion],
	// [PublicEventAnalyticsFilterCoalescingRefineByUnion],
	// [PublicFormSubmissionFilterCoalescingRefineByUnion],
	// [PublicFormSubmissionOnPageFilterCoalescingRefineByUnion],
	// [PublicNumAssociationsFilterCoalescingRefineByUnion],
	// [PublicUnifiedEventsFilterCoalescingRefineByUnion],
	// [PublicPropertyAssociationInListFilterCoalescingRefineByUnion].
	SetType    string `json:"setType"`
	Comparison string `json:"comparison"`
	// This field is from variant
	// [PublicAssociationInListFilterCoalescingRefineByUnion],
	// [PublicPageViewAnalyticsFilterCoalescingRefineByUnion],
	// [PublicCtaAnalyticsFilterCoalescingRefineByUnion],
	// [PublicEventAnalyticsFilterCoalescingRefineByUnion],
	// [PublicFormSubmissionFilterCoalescingRefineByUnion],
	// [PublicFormSubmissionOnPageFilterCoalescingRefineByUnion],
	// [PublicNumAssociationsFilterCoalescingRefineByUnion],
	// [PublicUnifiedEventsFilterCoalescingRefineByUnion],
	// [PublicPropertyAssociationInListFilterCoalescingRefineByUnion].
	TimeOffset PublicTimeOffset `json:"timeOffset"`
	// This field is from variant
	// [PublicAssociationInListFilterCoalescingRefineByUnion],
	// [PublicPageViewAnalyticsFilterCoalescingRefineByUnion],
	// [PublicCtaAnalyticsFilterCoalescingRefineByUnion],
	// [PublicEventAnalyticsFilterCoalescingRefineByUnion],
	// [PublicFormSubmissionFilterCoalescingRefineByUnion],
	// [PublicFormSubmissionOnPageFilterCoalescingRefineByUnion],
	// [PublicNumAssociationsFilterCoalescingRefineByUnion],
	// [PublicUnifiedEventsFilterCoalescingRefineByUnion],
	// [PublicPropertyAssociationInListFilterCoalescingRefineByUnion].
	LowerBoundOffset PublicTimeOffset `json:"lowerBoundOffset"`
	RangeType        string           `json:"rangeType"`
	// This field is from variant
	// [PublicAssociationInListFilterCoalescingRefineByUnion],
	// [PublicPageViewAnalyticsFilterCoalescingRefineByUnion],
	// [PublicCtaAnalyticsFilterCoalescingRefineByUnion],
	// [PublicEventAnalyticsFilterCoalescingRefineByUnion],
	// [PublicFormSubmissionFilterCoalescingRefineByUnion],
	// [PublicFormSubmissionOnPageFilterCoalescingRefineByUnion],
	// [PublicNumAssociationsFilterCoalescingRefineByUnion],
	// [PublicUnifiedEventsFilterCoalescingRefineByUnion],
	// [PublicPropertyAssociationInListFilterCoalescingRefineByUnion].
	UpperBoundOffset PublicTimeOffset `json:"upperBoundOffset"`
	// This field is from variant
	// [PublicAssociationInListFilterCoalescingRefineByUnion],
	// [PublicPageViewAnalyticsFilterCoalescingRefineByUnion],
	// [PublicCtaAnalyticsFilterCoalescingRefineByUnion],
	// [PublicEventAnalyticsFilterCoalescingRefineByUnion],
	// [PublicFormSubmissionFilterCoalescingRefineByUnion],
	// [PublicFormSubmissionOnPageFilterCoalescingRefineByUnion],
	// [PublicNumAssociationsFilterCoalescingRefineByUnion],
	// [PublicUnifiedEventsFilterCoalescingRefineByUnion],
	// [PublicPropertyAssociationInListFilterCoalescingRefineByUnion].
	Timestamp int64 `json:"timestamp"`
	// This field is from variant
	// [PublicAssociationInListFilterCoalescingRefineByUnion],
	// [PublicPageViewAnalyticsFilterCoalescingRefineByUnion],
	// [PublicCtaAnalyticsFilterCoalescingRefineByUnion],
	// [PublicEventAnalyticsFilterCoalescingRefineByUnion],
	// [PublicFormSubmissionFilterCoalescingRefineByUnion],
	// [PublicFormSubmissionOnPageFilterCoalescingRefineByUnion],
	// [PublicNumAssociationsFilterCoalescingRefineByUnion],
	// [PublicUnifiedEventsFilterCoalescingRefineByUnion],
	// [PublicPropertyAssociationInListFilterCoalescingRefineByUnion].
	LowerTimestamp int64 `json:"lowerTimestamp"`
	// This field is from variant
	// [PublicAssociationInListFilterCoalescingRefineByUnion],
	// [PublicPageViewAnalyticsFilterCoalescingRefineByUnion],
	// [PublicCtaAnalyticsFilterCoalescingRefineByUnion],
	// [PublicEventAnalyticsFilterCoalescingRefineByUnion],
	// [PublicFormSubmissionFilterCoalescingRefineByUnion],
	// [PublicFormSubmissionOnPageFilterCoalescingRefineByUnion],
	// [PublicNumAssociationsFilterCoalescingRefineByUnion],
	// [PublicUnifiedEventsFilterCoalescingRefineByUnion],
	// [PublicPropertyAssociationInListFilterCoalescingRefineByUnion].
	UpperTimestamp               int64  `json:"upperTimestamp"`
	IncludeObjectsWithNoValueSet bool   `json:"includeObjectsWithNoValueSet"`
	OperationType                string `json:"operationType"`
	Operator                     string `json:"operator"`
	// This field is from variant
	// [PublicAssociationInListFilterCoalescingRefineByUnion],
	// [PublicPageViewAnalyticsFilterCoalescingRefineByUnion],
	// [PublicCtaAnalyticsFilterCoalescingRefineByUnion],
	// [PublicEventAnalyticsFilterCoalescingRefineByUnion],
	// [PublicFormSubmissionFilterCoalescingRefineByUnion],
	// [PublicFormSubmissionOnPageFilterCoalescingRefineByUnion],
	// [PublicNumAssociationsFilterCoalescingRefineByUnion],
	// [PublicUnifiedEventsFilterCoalescingRefineByUnion],
	// [PublicPropertyAssociationInListFilterCoalescingRefineByUnion].
	TimePoint PublicTimePointOperationTimePointUnion `json:"timePoint"`
	// This field is from variant
	// [PublicAssociationInListFilterCoalescingRefineByUnion],
	// [PublicPageViewAnalyticsFilterCoalescingRefineByUnion],
	// [PublicCtaAnalyticsFilterCoalescingRefineByUnion],
	// [PublicEventAnalyticsFilterCoalescingRefineByUnion],
	// [PublicFormSubmissionFilterCoalescingRefineByUnion],
	// [PublicFormSubmissionOnPageFilterCoalescingRefineByUnion],
	// [PublicNumAssociationsFilterCoalescingRefineByUnion],
	// [PublicUnifiedEventsFilterCoalescingRefineByUnion],
	// [PublicPropertyAssociationInListFilterCoalescingRefineByUnion].
	EndpointBehavior string `json:"endpointBehavior"`
	PropertyParser   string `json:"propertyParser"`
	// This field is from variant
	// [PublicAssociationInListFilterCoalescingRefineByUnion],
	// [PublicPageViewAnalyticsFilterCoalescingRefineByUnion],
	// [PublicCtaAnalyticsFilterCoalescingRefineByUnion],
	// [PublicEventAnalyticsFilterCoalescingRefineByUnion],
	// [PublicFormSubmissionFilterCoalescingRefineByUnion],
	// [PublicFormSubmissionOnPageFilterCoalescingRefineByUnion],
	// [PublicNumAssociationsFilterCoalescingRefineByUnion],
	// [PublicUnifiedEventsFilterCoalescingRefineByUnion],
	// [PublicPropertyAssociationInListFilterCoalescingRefineByUnion].
	LowerBoundTimePoint PublicRangedTimeOperationLowerBoundTimePointUnion `json:"lowerBoundTimePoint"`
	// This field is from variant
	// [PublicAssociationInListFilterCoalescingRefineByUnion],
	// [PublicPageViewAnalyticsFilterCoalescingRefineByUnion],
	// [PublicCtaAnalyticsFilterCoalescingRefineByUnion],
	// [PublicEventAnalyticsFilterCoalescingRefineByUnion],
	// [PublicFormSubmissionFilterCoalescingRefineByUnion],
	// [PublicFormSubmissionOnPageFilterCoalescingRefineByUnion],
	// [PublicNumAssociationsFilterCoalescingRefineByUnion],
	// [PublicUnifiedEventsFilterCoalescingRefineByUnion],
	// [PublicPropertyAssociationInListFilterCoalescingRefineByUnion].
	UpperBoundTimePoint PublicRangedTimeOperationUpperBoundTimePointUnion `json:"upperBoundTimePoint"`
	// This field is from variant
	// [PublicAssociationInListFilterCoalescingRefineByUnion],
	// [PublicPageViewAnalyticsFilterCoalescingRefineByUnion],
	// [PublicCtaAnalyticsFilterCoalescingRefineByUnion],
	// [PublicEventAnalyticsFilterCoalescingRefineByUnion],
	// [PublicFormSubmissionFilterCoalescingRefineByUnion],
	// [PublicFormSubmissionOnPageFilterCoalescingRefineByUnion],
	// [PublicNumAssociationsFilterCoalescingRefineByUnion],
	// [PublicUnifiedEventsFilterCoalescingRefineByUnion],
	// [PublicPropertyAssociationInListFilterCoalescingRefineByUnion].
	LowerBoundEndpointBehavior string `json:"lowerBoundEndpointBehavior"`
	// This field is from variant
	// [PublicAssociationInListFilterCoalescingRefineByUnion],
	// [PublicPageViewAnalyticsFilterCoalescingRefineByUnion],
	// [PublicCtaAnalyticsFilterCoalescingRefineByUnion],
	// [PublicEventAnalyticsFilterCoalescingRefineByUnion],
	// [PublicFormSubmissionFilterCoalescingRefineByUnion],
	// [PublicFormSubmissionOnPageFilterCoalescingRefineByUnion],
	// [PublicNumAssociationsFilterCoalescingRefineByUnion],
	// [PublicUnifiedEventsFilterCoalescingRefineByUnion],
	// [PublicPropertyAssociationInListFilterCoalescingRefineByUnion].
	UpperBoundEndpointBehavior string `json:"upperBoundEndpointBehavior"`
	JSON                       struct {
		Type                         respjson.Field
		MaxOccurrences               respjson.Field
		MinOccurrences               respjson.Field
		SetType                      respjson.Field
		Comparison                   respjson.Field
		TimeOffset                   respjson.Field
		LowerBoundOffset             respjson.Field
		RangeType                    respjson.Field
		UpperBoundOffset             respjson.Field
		Timestamp                    respjson.Field
		LowerTimestamp               respjson.Field
		UpperTimestamp               respjson.Field
		IncludeObjectsWithNoValueSet respjson.Field
		OperationType                respjson.Field
		Operator                     respjson.Field
		TimePoint                    respjson.Field
		EndpointBehavior             respjson.Field
		PropertyParser               respjson.Field
		LowerBoundTimePoint          respjson.Field
		UpperBoundTimePoint          respjson.Field
		LowerBoundEndpointBehavior   respjson.Field
		UpperBoundEndpointBehavior   respjson.Field
		raw                          string
	} `json:"-"`
}

func (r *PublicAndFilterBranchFilterUnionCoalescingRefineBy) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PublicAndFilterBranchFilterUnionPruningRefineBy is an implicit subunion of
// [PublicAndFilterBranchFilterUnion].
// PublicAndFilterBranchFilterUnionPruningRefineBy provides convenient access to
// the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [PublicAndFilterBranchFilterUnion].
type PublicAndFilterBranchFilterUnionPruningRefineBy struct {
	Type string `json:"type"`
	// This field is from variant [PublicPageViewAnalyticsFilterPruningRefineByUnion],
	// [PublicCtaAnalyticsFilterPruningRefineByUnion],
	// [PublicEventAnalyticsFilterPruningRefineByUnion],
	// [PublicFormSubmissionFilterPruningRefineByUnion],
	// [PublicFormSubmissionOnPageFilterPruningRefineByUnion],
	// [PublicEmailEventFilterPruningRefineByUnion],
	// [PublicAdsTimeFilterPruningRefineByUnion],
	// [PublicUnifiedEventsFilterPruningRefineByUnion].
	MaxOccurrences int64 `json:"maxOccurrences"`
	// This field is from variant [PublicPageViewAnalyticsFilterPruningRefineByUnion],
	// [PublicCtaAnalyticsFilterPruningRefineByUnion],
	// [PublicEventAnalyticsFilterPruningRefineByUnion],
	// [PublicFormSubmissionFilterPruningRefineByUnion],
	// [PublicFormSubmissionOnPageFilterPruningRefineByUnion],
	// [PublicEmailEventFilterPruningRefineByUnion],
	// [PublicAdsTimeFilterPruningRefineByUnion],
	// [PublicUnifiedEventsFilterPruningRefineByUnion].
	MinOccurrences int64 `json:"minOccurrences"`
	// This field is from variant [PublicPageViewAnalyticsFilterPruningRefineByUnion],
	// [PublicCtaAnalyticsFilterPruningRefineByUnion],
	// [PublicEventAnalyticsFilterPruningRefineByUnion],
	// [PublicFormSubmissionFilterPruningRefineByUnion],
	// [PublicFormSubmissionOnPageFilterPruningRefineByUnion],
	// [PublicEmailEventFilterPruningRefineByUnion],
	// [PublicAdsTimeFilterPruningRefineByUnion],
	// [PublicUnifiedEventsFilterPruningRefineByUnion].
	SetType    string `json:"setType"`
	Comparison string `json:"comparison"`
	// This field is from variant [PublicPageViewAnalyticsFilterPruningRefineByUnion],
	// [PublicCtaAnalyticsFilterPruningRefineByUnion],
	// [PublicEventAnalyticsFilterPruningRefineByUnion],
	// [PublicFormSubmissionFilterPruningRefineByUnion],
	// [PublicFormSubmissionOnPageFilterPruningRefineByUnion],
	// [PublicEmailEventFilterPruningRefineByUnion],
	// [PublicAdsTimeFilterPruningRefineByUnion],
	// [PublicUnifiedEventsFilterPruningRefineByUnion].
	TimeOffset PublicTimeOffset `json:"timeOffset"`
	// This field is from variant [PublicPageViewAnalyticsFilterPruningRefineByUnion],
	// [PublicCtaAnalyticsFilterPruningRefineByUnion],
	// [PublicEventAnalyticsFilterPruningRefineByUnion],
	// [PublicFormSubmissionFilterPruningRefineByUnion],
	// [PublicFormSubmissionOnPageFilterPruningRefineByUnion],
	// [PublicEmailEventFilterPruningRefineByUnion],
	// [PublicAdsTimeFilterPruningRefineByUnion],
	// [PublicUnifiedEventsFilterPruningRefineByUnion].
	LowerBoundOffset PublicTimeOffset `json:"lowerBoundOffset"`
	RangeType        string           `json:"rangeType"`
	// This field is from variant [PublicPageViewAnalyticsFilterPruningRefineByUnion],
	// [PublicCtaAnalyticsFilterPruningRefineByUnion],
	// [PublicEventAnalyticsFilterPruningRefineByUnion],
	// [PublicFormSubmissionFilterPruningRefineByUnion],
	// [PublicFormSubmissionOnPageFilterPruningRefineByUnion],
	// [PublicEmailEventFilterPruningRefineByUnion],
	// [PublicAdsTimeFilterPruningRefineByUnion],
	// [PublicUnifiedEventsFilterPruningRefineByUnion].
	UpperBoundOffset PublicTimeOffset `json:"upperBoundOffset"`
	// This field is from variant [PublicPageViewAnalyticsFilterPruningRefineByUnion],
	// [PublicCtaAnalyticsFilterPruningRefineByUnion],
	// [PublicEventAnalyticsFilterPruningRefineByUnion],
	// [PublicFormSubmissionFilterPruningRefineByUnion],
	// [PublicFormSubmissionOnPageFilterPruningRefineByUnion],
	// [PublicEmailEventFilterPruningRefineByUnion],
	// [PublicAdsTimeFilterPruningRefineByUnion],
	// [PublicUnifiedEventsFilterPruningRefineByUnion].
	Timestamp int64 `json:"timestamp"`
	// This field is from variant [PublicPageViewAnalyticsFilterPruningRefineByUnion],
	// [PublicCtaAnalyticsFilterPruningRefineByUnion],
	// [PublicEventAnalyticsFilterPruningRefineByUnion],
	// [PublicFormSubmissionFilterPruningRefineByUnion],
	// [PublicFormSubmissionOnPageFilterPruningRefineByUnion],
	// [PublicEmailEventFilterPruningRefineByUnion],
	// [PublicAdsTimeFilterPruningRefineByUnion],
	// [PublicUnifiedEventsFilterPruningRefineByUnion].
	LowerTimestamp int64 `json:"lowerTimestamp"`
	// This field is from variant [PublicPageViewAnalyticsFilterPruningRefineByUnion],
	// [PublicCtaAnalyticsFilterPruningRefineByUnion],
	// [PublicEventAnalyticsFilterPruningRefineByUnion],
	// [PublicFormSubmissionFilterPruningRefineByUnion],
	// [PublicFormSubmissionOnPageFilterPruningRefineByUnion],
	// [PublicEmailEventFilterPruningRefineByUnion],
	// [PublicAdsTimeFilterPruningRefineByUnion],
	// [PublicUnifiedEventsFilterPruningRefineByUnion].
	UpperTimestamp               int64  `json:"upperTimestamp"`
	IncludeObjectsWithNoValueSet bool   `json:"includeObjectsWithNoValueSet"`
	OperationType                string `json:"operationType"`
	Operator                     string `json:"operator"`
	// This field is from variant [PublicPageViewAnalyticsFilterPruningRefineByUnion],
	// [PublicCtaAnalyticsFilterPruningRefineByUnion],
	// [PublicEventAnalyticsFilterPruningRefineByUnion],
	// [PublicFormSubmissionFilterPruningRefineByUnion],
	// [PublicFormSubmissionOnPageFilterPruningRefineByUnion],
	// [PublicEmailEventFilterPruningRefineByUnion],
	// [PublicAdsTimeFilterPruningRefineByUnion],
	// [PublicUnifiedEventsFilterPruningRefineByUnion].
	TimePoint PublicTimePointOperationTimePointUnion `json:"timePoint"`
	// This field is from variant [PublicPageViewAnalyticsFilterPruningRefineByUnion],
	// [PublicCtaAnalyticsFilterPruningRefineByUnion],
	// [PublicEventAnalyticsFilterPruningRefineByUnion],
	// [PublicFormSubmissionFilterPruningRefineByUnion],
	// [PublicFormSubmissionOnPageFilterPruningRefineByUnion],
	// [PublicEmailEventFilterPruningRefineByUnion],
	// [PublicAdsTimeFilterPruningRefineByUnion],
	// [PublicUnifiedEventsFilterPruningRefineByUnion].
	EndpointBehavior string `json:"endpointBehavior"`
	PropertyParser   string `json:"propertyParser"`
	// This field is from variant [PublicPageViewAnalyticsFilterPruningRefineByUnion],
	// [PublicCtaAnalyticsFilterPruningRefineByUnion],
	// [PublicEventAnalyticsFilterPruningRefineByUnion],
	// [PublicFormSubmissionFilterPruningRefineByUnion],
	// [PublicFormSubmissionOnPageFilterPruningRefineByUnion],
	// [PublicEmailEventFilterPruningRefineByUnion],
	// [PublicAdsTimeFilterPruningRefineByUnion],
	// [PublicUnifiedEventsFilterPruningRefineByUnion].
	LowerBoundTimePoint PublicRangedTimeOperationLowerBoundTimePointUnion `json:"lowerBoundTimePoint"`
	// This field is from variant [PublicPageViewAnalyticsFilterPruningRefineByUnion],
	// [PublicCtaAnalyticsFilterPruningRefineByUnion],
	// [PublicEventAnalyticsFilterPruningRefineByUnion],
	// [PublicFormSubmissionFilterPruningRefineByUnion],
	// [PublicFormSubmissionOnPageFilterPruningRefineByUnion],
	// [PublicEmailEventFilterPruningRefineByUnion],
	// [PublicAdsTimeFilterPruningRefineByUnion],
	// [PublicUnifiedEventsFilterPruningRefineByUnion].
	UpperBoundTimePoint PublicRangedTimeOperationUpperBoundTimePointUnion `json:"upperBoundTimePoint"`
	// This field is from variant [PublicPageViewAnalyticsFilterPruningRefineByUnion],
	// [PublicCtaAnalyticsFilterPruningRefineByUnion],
	// [PublicEventAnalyticsFilterPruningRefineByUnion],
	// [PublicFormSubmissionFilterPruningRefineByUnion],
	// [PublicFormSubmissionOnPageFilterPruningRefineByUnion],
	// [PublicEmailEventFilterPruningRefineByUnion],
	// [PublicAdsTimeFilterPruningRefineByUnion],
	// [PublicUnifiedEventsFilterPruningRefineByUnion].
	LowerBoundEndpointBehavior string `json:"lowerBoundEndpointBehavior"`
	// This field is from variant [PublicPageViewAnalyticsFilterPruningRefineByUnion],
	// [PublicCtaAnalyticsFilterPruningRefineByUnion],
	// [PublicEventAnalyticsFilterPruningRefineByUnion],
	// [PublicFormSubmissionFilterPruningRefineByUnion],
	// [PublicFormSubmissionOnPageFilterPruningRefineByUnion],
	// [PublicEmailEventFilterPruningRefineByUnion],
	// [PublicAdsTimeFilterPruningRefineByUnion],
	// [PublicUnifiedEventsFilterPruningRefineByUnion].
	UpperBoundEndpointBehavior string `json:"upperBoundEndpointBehavior"`
	JSON                       struct {
		Type                         respjson.Field
		MaxOccurrences               respjson.Field
		MinOccurrences               respjson.Field
		SetType                      respjson.Field
		Comparison                   respjson.Field
		TimeOffset                   respjson.Field
		LowerBoundOffset             respjson.Field
		RangeType                    respjson.Field
		UpperBoundOffset             respjson.Field
		Timestamp                    respjson.Field
		LowerTimestamp               respjson.Field
		UpperTimestamp               respjson.Field
		IncludeObjectsWithNoValueSet respjson.Field
		OperationType                respjson.Field
		Operator                     respjson.Field
		TimePoint                    respjson.Field
		EndpointBehavior             respjson.Field
		PropertyParser               respjson.Field
		LowerBoundTimePoint          respjson.Field
		UpperBoundTimePoint          respjson.Field
		LowerBoundEndpointBehavior   respjson.Field
		UpperBoundEndpointBehavior   respjson.Field
		raw                          string
	} `json:"-"`
}

func (r *PublicAndFilterBranchFilterUnionPruningRefineBy) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PublicAndFilterBranchFilterUnionEventTypeID is an implicit subunion of
// [PublicAndFilterBranchFilterUnion]. PublicAndFilterBranchFilterUnionEventTypeID
// provides convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [PublicAndFilterBranchFilterUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfInt OfString]
type PublicAndFilterBranchFilterUnionEventTypeID struct {
	// This field will be present if the value is a [int64] instead of an object.
	OfInt int64 `json:",inline"`
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	JSON     struct {
		OfInt    respjson.Field
		OfString respjson.Field
		raw      string
	} `json:"-"`
}

func (r *PublicAndFilterBranchFilterUnionEventTypeID) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties FilterBranches, FilterBranchOperator, FilterBranchType, Filters
// are required.
type PublicAndFilterBranchParam struct {
	FilterBranches []PublicAndFilterBranchFilterBranchUnionParam `json:"filterBranches,omitzero" api:"required"`
	// Filter branch operator (AND)
	FilterBranchOperator string `json:"filterBranchOperator" api:"required"`
	// Type of filter branch (AND)
	//
	// Any of "AND".
	FilterBranchType PublicAndFilterBranchFilterBranchType   `json:"filterBranchType,omitzero" api:"required"`
	Filters          []PublicAndFilterBranchFilterUnionParam `json:"filters,omitzero" api:"required"`
	paramObj
}

func (r PublicAndFilterBranchParam) MarshalJSON() (data []byte, err error) {
	type shadow PublicAndFilterBranchParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PublicAndFilterBranchParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type PublicAndFilterBranchFilterBranchUnionParam struct {
	OfOr                  *PublicOrFilterBranchParam                  `json:",omitzero,inline"`
	OfAnd                 *PublicAndFilterBranchParam                 `json:",omitzero,inline"`
	OfNotAll              *PublicNotAllFilterBranchParam              `json:",omitzero,inline"`
	OfNotAny              *PublicNotAnyFilterBranchParam              `json:",omitzero,inline"`
	OfRestricted          *PublicRestrictedFilterBranchParam          `json:",omitzero,inline"`
	OfUnifiedEvents       *PublicUnifiedEventsFilterBranchParam       `json:",omitzero,inline"`
	OfPropertyAssociation *PublicPropertyAssociationFilterBranchParam `json:",omitzero,inline"`
	OfAssociation         *PublicAssociationFilterBranchParam         `json:",omitzero,inline"`
	paramUnion
}

func (u PublicAndFilterBranchFilterBranchUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfOr,
		u.OfAnd,
		u.OfNotAll,
		u.OfNotAny,
		u.OfRestricted,
		u.OfUnifiedEvents,
		u.OfPropertyAssociation,
		u.OfAssociation)
}
func (u *PublicAndFilterBranchFilterBranchUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func init() {
	apijson.RegisterUnion[PublicAndFilterBranchFilterBranchUnionParam](
		"filterBranchType",
		apijson.Discriminator[PublicOrFilterBranchParam]("OR"),
		apijson.Discriminator[PublicAndFilterBranchParam]("AND"),
		apijson.Discriminator[PublicNotAllFilterBranchParam]("NOT_ALL"),
		apijson.Discriminator[PublicNotAnyFilterBranchParam]("NOT_ANY"),
		apijson.Discriminator[PublicRestrictedFilterBranchParam]("RESTRICTED"),
		apijson.Discriminator[PublicUnifiedEventsFilterBranchParam]("UNIFIED_EVENTS"),
		apijson.Discriminator[PublicPropertyAssociationFilterBranchParam]("PROPERTY_ASSOCIATION"),
		apijson.Discriminator[PublicAssociationFilterBranchParam]("ASSOCIATION"),
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type PublicAndFilterBranchFilterUnionParam struct {
	OfProperty                  *PublicPropertyFilterParam                  `json:",omitzero,inline"`
	OfAssociation               *PublicAssociationInListFilterParam         `json:",omitzero,inline"`
	OfPageView                  *PublicPageViewAnalyticsFilterParam         `json:",omitzero,inline"`
	OfCta                       *PublicCtaAnalyticsFilterParam              `json:",omitzero,inline"`
	OfEvent                     *PublicEventAnalyticsFilterParam            `json:",omitzero,inline"`
	OfFormSubmission            *PublicFormSubmissionFilterParam            `json:",omitzero,inline"`
	OfFormSubmissionOnPage      *PublicFormSubmissionOnPageFilterParam      `json:",omitzero,inline"`
	OfIntegrationEvent          *PublicIntegrationEventFilterParam          `json:",omitzero,inline"`
	OfEmailSubscription         *PublicEmailSubscriptionFilterParam         `json:",omitzero,inline"`
	OfCommunicationSubscription *PublicCommunicationSubscriptionFilterParam `json:",omitzero,inline"`
	OfCampaignInfluenced        *PublicCampaignInfluencedFilterParam        `json:",omitzero,inline"`
	OfSurveyMonkey              *PublicSurveyMonkeyFilterParam              `json:",omitzero,inline"`
	OfSurveyMonkeyValue         *PublicSurveyMonkeyValueFilterParam         `json:",omitzero,inline"`
	OfWebinar                   *PublicWebinarFilterParam                   `json:",omitzero,inline"`
	OfEmailEvent                *PublicEmailEventFilterParam                `json:",omitzero,inline"`
	OfPrivacy                   *PublicPrivacyAnalyticsFilterParam          `json:",omitzero,inline"`
	OfAdsSearch                 *PublicAdsSearchFilterParam                 `json:",omitzero,inline"`
	OfAdsTime                   *PublicAdsTimeFilterParam                   `json:",omitzero,inline"`
	OfInList                    *PublicInListFilterParam                    `json:",omitzero,inline"`
	OfNumAssociations           *PublicNumAssociationsFilterParam           `json:",omitzero,inline"`
	OfUnifiedEvents             *PublicUnifiedEventsFilterParam             `json:",omitzero,inline"`
	OfPropertyAssociation       *PublicPropertyAssociationInListFilterParam `json:",omitzero,inline"`
	OfConstant                  *PublicConstantFilterParam                  `json:",omitzero,inline"`
	paramUnion
}

func (u PublicAndFilterBranchFilterUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfProperty,
		u.OfAssociation,
		u.OfPageView,
		u.OfCta,
		u.OfEvent,
		u.OfFormSubmission,
		u.OfFormSubmissionOnPage,
		u.OfIntegrationEvent,
		u.OfEmailSubscription,
		u.OfCommunicationSubscription,
		u.OfCampaignInfluenced,
		u.OfSurveyMonkey,
		u.OfSurveyMonkeyValue,
		u.OfWebinar,
		u.OfEmailEvent,
		u.OfPrivacy,
		u.OfAdsSearch,
		u.OfAdsTime,
		u.OfInList,
		u.OfNumAssociations,
		u.OfUnifiedEvents,
		u.OfPropertyAssociation,
		u.OfConstant)
}
func (u *PublicAndFilterBranchFilterUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func init() {
	apijson.RegisterUnion[PublicAndFilterBranchFilterUnionParam](
		"filterType",
		apijson.Discriminator[PublicPropertyFilterParam]("PROPERTY"),
		apijson.Discriminator[PublicAssociationInListFilterParam]("ASSOCIATION"),
		apijson.Discriminator[PublicPageViewAnalyticsFilterParam]("PAGE_VIEW"),
		apijson.Discriminator[PublicCtaAnalyticsFilterParam]("CTA"),
		apijson.Discriminator[PublicEventAnalyticsFilterParam]("EVENT"),
		apijson.Discriminator[PublicFormSubmissionFilterParam]("FORM_SUBMISSION"),
		apijson.Discriminator[PublicFormSubmissionOnPageFilterParam]("FORM_SUBMISSION_ON_PAGE"),
		apijson.Discriminator[PublicIntegrationEventFilterParam]("INTEGRATION_EVENT"),
		apijson.Discriminator[PublicEmailSubscriptionFilterParam]("EMAIL_SUBSCRIPTION"),
		apijson.Discriminator[PublicCommunicationSubscriptionFilterParam]("COMMUNICATION_SUBSCRIPTION"),
		apijson.Discriminator[PublicCampaignInfluencedFilterParam]("CAMPAIGN_INFLUENCED"),
		apijson.Discriminator[PublicSurveyMonkeyFilterParam]("SURVEY_MONKEY"),
		apijson.Discriminator[PublicSurveyMonkeyValueFilterParam]("SURVEY_MONKEY_VALUE"),
		apijson.Discriminator[PublicWebinarFilterParam]("WEBINAR"),
		apijson.Discriminator[PublicEmailEventFilterParam]("EMAIL_EVENT"),
		apijson.Discriminator[PublicPrivacyAnalyticsFilterParam]("PRIVACY"),
		apijson.Discriminator[PublicAdsSearchFilterParam]("ADS_SEARCH"),
		apijson.Discriminator[PublicAdsTimeFilterParam]("ADS_TIME"),
		apijson.Discriminator[PublicInListFilterParam]("IN_LIST"),
		apijson.Discriminator[PublicNumAssociationsFilterParam]("NUM_ASSOCIATIONS"),
		apijson.Discriminator[PublicUnifiedEventsFilterParam]("UNIFIED_EVENTS"),
		apijson.Discriminator[PublicPropertyAssociationInListFilterParam]("PROPERTY_ASSOCIATION"),
		apijson.Discriminator[PublicConstantFilterParam]("CONSTANT"),
	)
}

type PublicAssociationFilterBranch struct {
	// Specifies the category of the association for the filter branch
	// (HUBSPOT_DEFINED, USER_DEFINED, INTEGRATOR_DEFINED, WORK).
	AssociationCategory string `json:"associationCategory" api:"required"`
	// Type id of the association
	AssociationTypeID int64                                            `json:"associationTypeId" api:"required"`
	FilterBranches    []PublicAssociationFilterBranchFilterBranchUnion `json:"filterBranches" api:"required"`
	// Filter branch operator (AND)
	FilterBranchOperator string `json:"filterBranchOperator" api:"required"`
	// Type of the filter branch (ASSOCIATION)
	//
	// Any of "ASSOCIATION".
	FilterBranchType PublicAssociationFilterBranchFilterBranchType `json:"filterBranchType" api:"required"`
	Filters          []PublicAssociationFilterBranchFilterUnion    `json:"filters" api:"required"`
	// The ID representing the type of object associated with the filter branch.
	ObjectTypeID string `json:"objectTypeId" api:"required"`
	// Defines the operation to be applied within the filter branch (IN_LIST,
	// NOT_IN_LIST).
	Operator string `json:"operator" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AssociationCategory  respjson.Field
		AssociationTypeID    respjson.Field
		FilterBranches       respjson.Field
		FilterBranchOperator respjson.Field
		FilterBranchType     respjson.Field
		Filters              respjson.Field
		ObjectTypeID         respjson.Field
		Operator             respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PublicAssociationFilterBranch) RawJSON() string { return r.JSON.raw }
func (r *PublicAssociationFilterBranch) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this PublicAssociationFilterBranch to a
// PublicAssociationFilterBranchParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// PublicAssociationFilterBranchParam.Overrides()
func (r PublicAssociationFilterBranch) ToParam() PublicAssociationFilterBranchParam {
	return param.Override[PublicAssociationFilterBranchParam](json.RawMessage(r.RawJSON()))
}

// PublicAssociationFilterBranchFilterBranchUnion contains all possible properties
// and values from [PublicOrFilterBranch], [PublicAndFilterBranch],
// [PublicNotAllFilterBranch], [PublicNotAnyFilterBranch],
// [PublicRestrictedFilterBranch], [PublicUnifiedEventsFilterBranch],
// [PublicPropertyAssociationFilterBranch], [PublicAssociationFilterBranch].
//
// Use the [PublicAssociationFilterBranchFilterBranchUnion.AsAny] method to switch
// on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type PublicAssociationFilterBranchFilterBranchUnion struct {
	// This field is a union of [[]PublicOrFilterBranchFilterBranchUnion],
	// [[]PublicAndFilterBranchFilterBranchUnion],
	// [[]PublicNotAllFilterBranchFilterBranchUnion],
	// [[]PublicNotAnyFilterBranchFilterBranchUnion],
	// [[]PublicRestrictedFilterBranchFilterBranchUnion],
	// [[]PublicUnifiedEventsFilterBranchFilterBranchUnion],
	// [[]PublicPropertyAssociationFilterBranchFilterBranchUnion],
	// [[]PublicAssociationFilterBranchFilterBranchUnion]
	FilterBranches       PublicAssociationFilterBranchFilterBranchUnionFilterBranches `json:"filterBranches"`
	FilterBranchOperator string                                                       `json:"filterBranchOperator"`
	// Any of "OR", "AND", "NOT_ALL", "NOT_ANY", "RESTRICTED", "UNIFIED_EVENTS",
	// "PROPERTY_ASSOCIATION", "ASSOCIATION".
	FilterBranchType string `json:"filterBranchType"`
	// This field is a union of [[]PublicOrFilterBranchFilterUnion],
	// [[]PublicAndFilterBranchFilterUnion], [[]PublicNotAllFilterBranchFilterUnion],
	// [[]PublicNotAnyFilterBranchFilterUnion],
	// [[]PublicRestrictedFilterBranchFilterUnion],
	// [[]PublicUnifiedEventsFilterBranchFilterUnion],
	// [[]PublicPropertyAssociationFilterBranchFilterUnion],
	// [[]PublicAssociationFilterBranchFilterUnion]
	Filters PublicAssociationFilterBranchFilterBranchUnionFilters `json:"filters"`
	// This field is from variant [PublicUnifiedEventsFilterBranch].
	EventTypeID string `json:"eventTypeId"`
	Operator    string `json:"operator"`
	// This field is from variant [PublicUnifiedEventsFilterBranch].
	CoalescingRefineBy PublicUnifiedEventsFilterBranchCoalescingRefineByUnion `json:"coalescingRefineBy"`
	// This field is from variant [PublicUnifiedEventsFilterBranch].
	PruningRefineBy PublicUnifiedEventsFilterBranchPruningRefineByUnion `json:"pruningRefineBy"`
	ObjectTypeID    string                                              `json:"objectTypeId"`
	// This field is from variant [PublicPropertyAssociationFilterBranch].
	PropertyWithObjectID string `json:"propertyWithObjectId"`
	// This field is from variant [PublicAssociationFilterBranch].
	AssociationCategory string `json:"associationCategory"`
	// This field is from variant [PublicAssociationFilterBranch].
	AssociationTypeID int64 `json:"associationTypeId"`
	JSON              struct {
		FilterBranches       respjson.Field
		FilterBranchOperator respjson.Field
		FilterBranchType     respjson.Field
		Filters              respjson.Field
		EventTypeID          respjson.Field
		Operator             respjson.Field
		CoalescingRefineBy   respjson.Field
		PruningRefineBy      respjson.Field
		ObjectTypeID         respjson.Field
		PropertyWithObjectID respjson.Field
		AssociationCategory  respjson.Field
		AssociationTypeID    respjson.Field
		raw                  string
	} `json:"-"`
}

// anyPublicAssociationFilterBranchFilterBranch is implemented by each variant of
// [PublicAssociationFilterBranchFilterBranchUnion] to add type safety for the
// return type of [PublicAssociationFilterBranchFilterBranchUnion.AsAny]
type anyPublicAssociationFilterBranchFilterBranch interface {
	implPublicAssociationFilterBranchFilterBranchUnion()
}

func (PublicOrFilterBranch) implPublicAssociationFilterBranchFilterBranchUnion()                  {}
func (PublicAndFilterBranch) implPublicAssociationFilterBranchFilterBranchUnion()                 {}
func (PublicNotAllFilterBranch) implPublicAssociationFilterBranchFilterBranchUnion()              {}
func (PublicNotAnyFilterBranch) implPublicAssociationFilterBranchFilterBranchUnion()              {}
func (PublicRestrictedFilterBranch) implPublicAssociationFilterBranchFilterBranchUnion()          {}
func (PublicUnifiedEventsFilterBranch) implPublicAssociationFilterBranchFilterBranchUnion()       {}
func (PublicPropertyAssociationFilterBranch) implPublicAssociationFilterBranchFilterBranchUnion() {}
func (PublicAssociationFilterBranch) implPublicAssociationFilterBranchFilterBranchUnion()         {}

// Use the following switch statement to find the correct variant
//
//	switch variant := PublicAssociationFilterBranchFilterBranchUnion.AsAny().(type) {
//	case crm.PublicOrFilterBranch:
//	case crm.PublicAndFilterBranch:
//	case crm.PublicNotAllFilterBranch:
//	case crm.PublicNotAnyFilterBranch:
//	case crm.PublicRestrictedFilterBranch:
//	case crm.PublicUnifiedEventsFilterBranch:
//	case crm.PublicPropertyAssociationFilterBranch:
//	case crm.PublicAssociationFilterBranch:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u PublicAssociationFilterBranchFilterBranchUnion) AsAny() anyPublicAssociationFilterBranchFilterBranch {
	switch u.FilterBranchType {
	case "OR":
		return u.AsOr()
	case "AND":
		return u.AsAnd()
	case "NOT_ALL":
		return u.AsNotAll()
	case "NOT_ANY":
		return u.AsNotAny()
	case "RESTRICTED":
		return u.AsRestricted()
	case "UNIFIED_EVENTS":
		return u.AsUnifiedEvents()
	case "PROPERTY_ASSOCIATION":
		return u.AsPropertyAssociation()
	case "ASSOCIATION":
		return u.AsAssociation()
	}
	return nil
}

func (u PublicAssociationFilterBranchFilterBranchUnion) AsOr() (v PublicOrFilterBranch) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicAssociationFilterBranchFilterBranchUnion) AsAnd() (v PublicAndFilterBranch) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicAssociationFilterBranchFilterBranchUnion) AsNotAll() (v PublicNotAllFilterBranch) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicAssociationFilterBranchFilterBranchUnion) AsNotAny() (v PublicNotAnyFilterBranch) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicAssociationFilterBranchFilterBranchUnion) AsRestricted() (v PublicRestrictedFilterBranch) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicAssociationFilterBranchFilterBranchUnion) AsUnifiedEvents() (v PublicUnifiedEventsFilterBranch) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicAssociationFilterBranchFilterBranchUnion) AsPropertyAssociation() (v PublicPropertyAssociationFilterBranch) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicAssociationFilterBranchFilterBranchUnion) AsAssociation() (v PublicAssociationFilterBranch) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u PublicAssociationFilterBranchFilterBranchUnion) RawJSON() string { return u.JSON.raw }

func (r *PublicAssociationFilterBranchFilterBranchUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PublicAssociationFilterBranchFilterBranchUnionFilterBranches is an implicit
// subunion of [PublicAssociationFilterBranchFilterBranchUnion].
// PublicAssociationFilterBranchFilterBranchUnionFilterBranches provides convenient
// access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [PublicAssociationFilterBranchFilterBranchUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfPublicOrFilterBranchFilterBranches
// OfPublicAndFilterBranchFilterBranches OfPublicNotAllFilterBranchFilterBranches
// OfPublicNotAnyFilterBranchFilterBranches
// OfPublicRestrictedFilterBranchFilterBranches
// OfPublicUnifiedEventsFilterBranchFilterBranches
// OfPublicPropertyAssociationFilterBranchFilterBranches
// OfPublicAssociationFilterBranchFilterBranches]
type PublicAssociationFilterBranchFilterBranchUnionFilterBranches struct {
	// This field will be present if the value is a
	// [[]PublicOrFilterBranchFilterBranchUnion] instead of an object.
	OfPublicOrFilterBranchFilterBranches []PublicOrFilterBranchFilterBranchUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]PublicAndFilterBranchFilterBranchUnion] instead of an object.
	OfPublicAndFilterBranchFilterBranches []PublicAndFilterBranchFilterBranchUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]PublicNotAllFilterBranchFilterBranchUnion] instead of an object.
	OfPublicNotAllFilterBranchFilterBranches []PublicNotAllFilterBranchFilterBranchUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]PublicNotAnyFilterBranchFilterBranchUnion] instead of an object.
	OfPublicNotAnyFilterBranchFilterBranches []PublicNotAnyFilterBranchFilterBranchUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]PublicRestrictedFilterBranchFilterBranchUnion] instead of an object.
	OfPublicRestrictedFilterBranchFilterBranches []PublicRestrictedFilterBranchFilterBranchUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]PublicUnifiedEventsFilterBranchFilterBranchUnion] instead of an object.
	OfPublicUnifiedEventsFilterBranchFilterBranches []PublicUnifiedEventsFilterBranchFilterBranchUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]PublicPropertyAssociationFilterBranchFilterBranchUnion] instead of an object.
	OfPublicPropertyAssociationFilterBranchFilterBranches []PublicPropertyAssociationFilterBranchFilterBranchUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]PublicAssociationFilterBranchFilterBranchUnion] instead of an object.
	OfPublicAssociationFilterBranchFilterBranches []PublicAssociationFilterBranchFilterBranchUnion `json:",inline"`
	JSON                                          struct {
		OfPublicOrFilterBranchFilterBranches                  respjson.Field
		OfPublicAndFilterBranchFilterBranches                 respjson.Field
		OfPublicNotAllFilterBranchFilterBranches              respjson.Field
		OfPublicNotAnyFilterBranchFilterBranches              respjson.Field
		OfPublicRestrictedFilterBranchFilterBranches          respjson.Field
		OfPublicUnifiedEventsFilterBranchFilterBranches       respjson.Field
		OfPublicPropertyAssociationFilterBranchFilterBranches respjson.Field
		OfPublicAssociationFilterBranchFilterBranches         respjson.Field
		raw                                                   string
	} `json:"-"`
}

func (r *PublicAssociationFilterBranchFilterBranchUnionFilterBranches) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PublicAssociationFilterBranchFilterBranchUnionFilters is an implicit subunion of
// [PublicAssociationFilterBranchFilterBranchUnion].
// PublicAssociationFilterBranchFilterBranchUnionFilters provides convenient access
// to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [PublicAssociationFilterBranchFilterBranchUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfPublicOrFilterBranchFilters OfPublicAndFilterBranchFilters
// OfPublicNotAllFilterBranchFilters OfPublicNotAnyFilterBranchFilters
// OfPublicRestrictedFilterBranchFilters OfPublicUnifiedEventsFilterBranchFilters
// OfPublicPropertyAssociationFilterBranchFilters
// OfPublicAssociationFilterBranchFilters]
type PublicAssociationFilterBranchFilterBranchUnionFilters struct {
	// This field will be present if the value is a [[]PublicOrFilterBranchFilterUnion]
	// instead of an object.
	OfPublicOrFilterBranchFilters []PublicOrFilterBranchFilterUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]PublicAndFilterBranchFilterUnion] instead of an object.
	OfPublicAndFilterBranchFilters []PublicAndFilterBranchFilterUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]PublicNotAllFilterBranchFilterUnion] instead of an object.
	OfPublicNotAllFilterBranchFilters []PublicNotAllFilterBranchFilterUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]PublicNotAnyFilterBranchFilterUnion] instead of an object.
	OfPublicNotAnyFilterBranchFilters []PublicNotAnyFilterBranchFilterUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]PublicRestrictedFilterBranchFilterUnion] instead of an object.
	OfPublicRestrictedFilterBranchFilters []PublicRestrictedFilterBranchFilterUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]PublicUnifiedEventsFilterBranchFilterUnion] instead of an object.
	OfPublicUnifiedEventsFilterBranchFilters []PublicUnifiedEventsFilterBranchFilterUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]PublicPropertyAssociationFilterBranchFilterUnion] instead of an object.
	OfPublicPropertyAssociationFilterBranchFilters []PublicPropertyAssociationFilterBranchFilterUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]PublicAssociationFilterBranchFilterUnion] instead of an object.
	OfPublicAssociationFilterBranchFilters []PublicAssociationFilterBranchFilterUnion `json:",inline"`
	JSON                                   struct {
		OfPublicOrFilterBranchFilters                  respjson.Field
		OfPublicAndFilterBranchFilters                 respjson.Field
		OfPublicNotAllFilterBranchFilters              respjson.Field
		OfPublicNotAnyFilterBranchFilters              respjson.Field
		OfPublicRestrictedFilterBranchFilters          respjson.Field
		OfPublicUnifiedEventsFilterBranchFilters       respjson.Field
		OfPublicPropertyAssociationFilterBranchFilters respjson.Field
		OfPublicAssociationFilterBranchFilters         respjson.Field
		raw                                            string
	} `json:"-"`
}

func (r *PublicAssociationFilterBranchFilterBranchUnionFilters) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Type of the filter branch (ASSOCIATION)
type PublicAssociationFilterBranchFilterBranchType string

const (
	PublicAssociationFilterBranchFilterBranchTypeAssociation PublicAssociationFilterBranchFilterBranchType = "ASSOCIATION"
)

// PublicAssociationFilterBranchFilterUnion contains all possible properties and
// values from [PublicPropertyFilter], [PublicAssociationInListFilter],
// [PublicPageViewAnalyticsFilter], [PublicCtaAnalyticsFilter],
// [PublicEventAnalyticsFilter], [PublicFormSubmissionFilter],
// [PublicFormSubmissionOnPageFilter], [PublicIntegrationEventFilter],
// [PublicEmailSubscriptionFilter], [PublicCommunicationSubscriptionFilter],
// [PublicCampaignInfluencedFilter], [PublicSurveyMonkeyFilter],
// [PublicSurveyMonkeyValueFilter], [PublicWebinarFilter],
// [PublicEmailEventFilter], [PublicPrivacyAnalyticsFilter],
// [PublicAdsSearchFilter], [PublicAdsTimeFilter], [PublicInListFilter],
// [PublicNumAssociationsFilter], [PublicUnifiedEventsFilter],
// [PublicPropertyAssociationInListFilter], [PublicConstantFilter].
//
// Use the [PublicAssociationFilterBranchFilterUnion.AsAny] method to switch on the
// variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type PublicAssociationFilterBranchFilterUnion struct {
	// Any of "PROPERTY", "ASSOCIATION", "PAGE_VIEW", "CTA", "EVENT",
	// "FORM_SUBMISSION", "FORM_SUBMISSION_ON_PAGE", "INTEGRATION_EVENT",
	// "EMAIL_SUBSCRIPTION", "COMMUNICATION_SUBSCRIPTION", "CAMPAIGN_INFLUENCED",
	// "SURVEY_MONKEY", "SURVEY_MONKEY_VALUE", "WEBINAR", "EMAIL_EVENT", "PRIVACY",
	// "ADS_SEARCH", "ADS_TIME", "IN_LIST", "NUM_ASSOCIATIONS", "UNIFIED_EVENTS",
	// "PROPERTY_ASSOCIATION", "CONSTANT".
	FilterType string `json:"filterType"`
	// This field is from variant [PublicPropertyFilter].
	Operation PublicPropertyFilterOperationUnion `json:"operation"`
	// This field is from variant [PublicPropertyFilter].
	Property            string `json:"property"`
	AssociationCategory string `json:"associationCategory"`
	AssociationTypeID   int64  `json:"associationTypeId"`
	// This field is a union of [PublicAssociationInListFilterCoalescingRefineByUnion],
	// [PublicPageViewAnalyticsFilterCoalescingRefineByUnion],
	// [PublicCtaAnalyticsFilterCoalescingRefineByUnion],
	// [PublicEventAnalyticsFilterCoalescingRefineByUnion],
	// [PublicFormSubmissionFilterCoalescingRefineByUnion],
	// [PublicFormSubmissionOnPageFilterCoalescingRefineByUnion],
	// [PublicNumAssociationsFilterCoalescingRefineByUnion],
	// [PublicUnifiedEventsFilterCoalescingRefineByUnion],
	// [PublicPropertyAssociationInListFilterCoalescingRefineByUnion]
	CoalescingRefineBy PublicAssociationFilterBranchFilterUnionCoalescingRefineBy `json:"coalescingRefineBy"`
	ListID             string                                                     `json:"listId"`
	Operator           string                                                     `json:"operator"`
	// This field is from variant [PublicAssociationInListFilter].
	ToObjectType   string `json:"toObjectType"`
	ToObjectTypeID string `json:"toObjectTypeId"`
	// This field is from variant [PublicPageViewAnalyticsFilter].
	PageURL string `json:"pageUrl"`
	// This field is from variant [PublicPageViewAnalyticsFilter].
	EnableTracking bool `json:"enableTracking"`
	// This field is a union of [PublicPageViewAnalyticsFilterPruningRefineByUnion],
	// [PublicCtaAnalyticsFilterPruningRefineByUnion],
	// [PublicEventAnalyticsFilterPruningRefineByUnion],
	// [PublicFormSubmissionFilterPruningRefineByUnion],
	// [PublicFormSubmissionOnPageFilterPruningRefineByUnion],
	// [PublicEmailEventFilterPruningRefineByUnion],
	// [PublicAdsTimeFilterPruningRefineByUnion],
	// [PublicUnifiedEventsFilterPruningRefineByUnion]
	PruningRefineBy PublicAssociationFilterBranchFilterUnionPruningRefineBy `json:"pruningRefineBy"`
	// This field is from variant [PublicCtaAnalyticsFilter].
	CtaName string `json:"ctaName"`
	// This field is from variant [PublicEventAnalyticsFilter].
	EventID string `json:"eventId"`
	FormID  string `json:"formId"`
	// This field is from variant [PublicFormSubmissionOnPageFilter].
	PageID string `json:"pageId"`
	// This field is a union of [int64], [string]
	EventTypeID PublicAssociationFilterBranchFilterUnionEventTypeID `json:"eventTypeId"`
	FilterLines []PublicEventFilterMetadata                         `json:"filterLines"`
	// This field is from variant [PublicEmailSubscriptionFilter].
	AcceptedStatuses []string `json:"acceptedStatuses"`
	SubscriptionIDs  []string `json:"subscriptionIds"`
	SubscriptionType string   `json:"subscriptionType"`
	// This field is from variant [PublicCommunicationSubscriptionFilter].
	AcceptedOptStates []string `json:"acceptedOptStates"`
	// This field is from variant [PublicCommunicationSubscriptionFilter].
	Channel string `json:"channel"`
	// This field is from variant [PublicCommunicationSubscriptionFilter].
	BusinessUnitID string `json:"businessUnitId"`
	// This field is from variant [PublicCampaignInfluencedFilter].
	CampaignID string `json:"campaignId"`
	SurveyID   string `json:"surveyId"`
	// This field is from variant [PublicSurveyMonkeyValueFilter].
	SurveyQuestion string `json:"surveyQuestion"`
	// This field is from variant [PublicSurveyMonkeyValueFilter].
	ValueComparison PublicSurveyMonkeyValueFilterValueComparisonUnion `json:"valueComparison"`
	// This field is from variant [PublicSurveyMonkeyValueFilter].
	SurveyAnswerColID string `json:"surveyAnswerColId"`
	// This field is from variant [PublicSurveyMonkeyValueFilter].
	SurveyAnswerRowID string `json:"surveyAnswerRowId"`
	// This field is from variant [PublicWebinarFilter].
	WebinarID string `json:"webinarId"`
	// This field is from variant [PublicEmailEventFilter].
	AppID string `json:"appId"`
	// This field is from variant [PublicEmailEventFilter].
	EmailID string `json:"emailId"`
	// This field is from variant [PublicEmailEventFilter].
	Level string `json:"level"`
	// This field is from variant [PublicEmailEventFilter].
	ClickURL string `json:"clickUrl"`
	// This field is from variant [PublicPrivacyAnalyticsFilter].
	PrivacyName string `json:"privacyName"`
	// This field is from variant [PublicAdsSearchFilter].
	AdNetwork string `json:"adNetwork"`
	// This field is from variant [PublicAdsSearchFilter].
	EntityType string `json:"entityType"`
	// This field is from variant [PublicAdsSearchFilter].
	SearchTerms []string `json:"searchTerms"`
	// This field is from variant [PublicAdsSearchFilter].
	SearchTermType string `json:"searchTermType"`
	// This field is from variant [PublicInListFilter].
	Metadata PublicInListFilterMetadata `json:"metadata"`
	// This field is from variant [PublicPropertyAssociationInListFilter].
	PropertyWithObjectID string `json:"propertyWithObjectId"`
	// This field is from variant [PublicConstantFilter].
	ShouldAccept bool `json:"shouldAccept"`
	// This field is from variant [PublicConstantFilter].
	Source string `json:"source"`
	JSON   struct {
		FilterType           respjson.Field
		Operation            respjson.Field
		Property             respjson.Field
		AssociationCategory  respjson.Field
		AssociationTypeID    respjson.Field
		CoalescingRefineBy   respjson.Field
		ListID               respjson.Field
		Operator             respjson.Field
		ToObjectType         respjson.Field
		ToObjectTypeID       respjson.Field
		PageURL              respjson.Field
		EnableTracking       respjson.Field
		PruningRefineBy      respjson.Field
		CtaName              respjson.Field
		EventID              respjson.Field
		FormID               respjson.Field
		PageID               respjson.Field
		EventTypeID          respjson.Field
		FilterLines          respjson.Field
		AcceptedStatuses     respjson.Field
		SubscriptionIDs      respjson.Field
		SubscriptionType     respjson.Field
		AcceptedOptStates    respjson.Field
		Channel              respjson.Field
		BusinessUnitID       respjson.Field
		CampaignID           respjson.Field
		SurveyID             respjson.Field
		SurveyQuestion       respjson.Field
		ValueComparison      respjson.Field
		SurveyAnswerColID    respjson.Field
		SurveyAnswerRowID    respjson.Field
		WebinarID            respjson.Field
		AppID                respjson.Field
		EmailID              respjson.Field
		Level                respjson.Field
		ClickURL             respjson.Field
		PrivacyName          respjson.Field
		AdNetwork            respjson.Field
		EntityType           respjson.Field
		SearchTerms          respjson.Field
		SearchTermType       respjson.Field
		Metadata             respjson.Field
		PropertyWithObjectID respjson.Field
		ShouldAccept         respjson.Field
		Source               respjson.Field
		raw                  string
	} `json:"-"`
}

// anyPublicAssociationFilterBranchFilter is implemented by each variant of
// [PublicAssociationFilterBranchFilterUnion] to add type safety for the return
// type of [PublicAssociationFilterBranchFilterUnion.AsAny]
type anyPublicAssociationFilterBranchFilter interface {
	implPublicAssociationFilterBranchFilterUnion()
}

func (PublicPropertyFilter) implPublicAssociationFilterBranchFilterUnion()                  {}
func (PublicAssociationInListFilter) implPublicAssociationFilterBranchFilterUnion()         {}
func (PublicPageViewAnalyticsFilter) implPublicAssociationFilterBranchFilterUnion()         {}
func (PublicCtaAnalyticsFilter) implPublicAssociationFilterBranchFilterUnion()              {}
func (PublicEventAnalyticsFilter) implPublicAssociationFilterBranchFilterUnion()            {}
func (PublicFormSubmissionFilter) implPublicAssociationFilterBranchFilterUnion()            {}
func (PublicFormSubmissionOnPageFilter) implPublicAssociationFilterBranchFilterUnion()      {}
func (PublicIntegrationEventFilter) implPublicAssociationFilterBranchFilterUnion()          {}
func (PublicEmailSubscriptionFilter) implPublicAssociationFilterBranchFilterUnion()         {}
func (PublicCommunicationSubscriptionFilter) implPublicAssociationFilterBranchFilterUnion() {}
func (PublicCampaignInfluencedFilter) implPublicAssociationFilterBranchFilterUnion()        {}
func (PublicSurveyMonkeyFilter) implPublicAssociationFilterBranchFilterUnion()              {}
func (PublicSurveyMonkeyValueFilter) implPublicAssociationFilterBranchFilterUnion()         {}
func (PublicWebinarFilter) implPublicAssociationFilterBranchFilterUnion()                   {}
func (PublicEmailEventFilter) implPublicAssociationFilterBranchFilterUnion()                {}
func (PublicPrivacyAnalyticsFilter) implPublicAssociationFilterBranchFilterUnion()          {}
func (PublicAdsSearchFilter) implPublicAssociationFilterBranchFilterUnion()                 {}
func (PublicAdsTimeFilter) implPublicAssociationFilterBranchFilterUnion()                   {}
func (PublicInListFilter) implPublicAssociationFilterBranchFilterUnion()                    {}
func (PublicNumAssociationsFilter) implPublicAssociationFilterBranchFilterUnion()           {}
func (PublicUnifiedEventsFilter) implPublicAssociationFilterBranchFilterUnion()             {}
func (PublicPropertyAssociationInListFilter) implPublicAssociationFilterBranchFilterUnion() {}
func (PublicConstantFilter) implPublicAssociationFilterBranchFilterUnion()                  {}

// Use the following switch statement to find the correct variant
//
//	switch variant := PublicAssociationFilterBranchFilterUnion.AsAny().(type) {
//	case crm.PublicPropertyFilter:
//	case crm.PublicAssociationInListFilter:
//	case crm.PublicPageViewAnalyticsFilter:
//	case crm.PublicCtaAnalyticsFilter:
//	case crm.PublicEventAnalyticsFilter:
//	case crm.PublicFormSubmissionFilter:
//	case crm.PublicFormSubmissionOnPageFilter:
//	case crm.PublicIntegrationEventFilter:
//	case crm.PublicEmailSubscriptionFilter:
//	case crm.PublicCommunicationSubscriptionFilter:
//	case crm.PublicCampaignInfluencedFilter:
//	case crm.PublicSurveyMonkeyFilter:
//	case crm.PublicSurveyMonkeyValueFilter:
//	case crm.PublicWebinarFilter:
//	case crm.PublicEmailEventFilter:
//	case crm.PublicPrivacyAnalyticsFilter:
//	case crm.PublicAdsSearchFilter:
//	case crm.PublicAdsTimeFilter:
//	case crm.PublicInListFilter:
//	case crm.PublicNumAssociationsFilter:
//	case crm.PublicUnifiedEventsFilter:
//	case crm.PublicPropertyAssociationInListFilter:
//	case crm.PublicConstantFilter:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u PublicAssociationFilterBranchFilterUnion) AsAny() anyPublicAssociationFilterBranchFilter {
	switch u.FilterType {
	case "PROPERTY":
		return u.AsProperty()
	case "ASSOCIATION":
		return u.AsAssociation()
	case "PAGE_VIEW":
		return u.AsPageView()
	case "CTA":
		return u.AsCta()
	case "EVENT":
		return u.AsEvent()
	case "FORM_SUBMISSION":
		return u.AsFormSubmission()
	case "FORM_SUBMISSION_ON_PAGE":
		return u.AsFormSubmissionOnPage()
	case "INTEGRATION_EVENT":
		return u.AsIntegrationEvent()
	case "EMAIL_SUBSCRIPTION":
		return u.AsEmailSubscription()
	case "COMMUNICATION_SUBSCRIPTION":
		return u.AsCommunicationSubscription()
	case "CAMPAIGN_INFLUENCED":
		return u.AsCampaignInfluenced()
	case "SURVEY_MONKEY":
		return u.AsSurveyMonkey()
	case "SURVEY_MONKEY_VALUE":
		return u.AsSurveyMonkeyValue()
	case "WEBINAR":
		return u.AsWebinar()
	case "EMAIL_EVENT":
		return u.AsEmailEvent()
	case "PRIVACY":
		return u.AsPrivacy()
	case "ADS_SEARCH":
		return u.AsAdsSearch()
	case "ADS_TIME":
		return u.AsAdsTime()
	case "IN_LIST":
		return u.AsInList()
	case "NUM_ASSOCIATIONS":
		return u.AsNumAssociations()
	case "UNIFIED_EVENTS":
		return u.AsUnifiedEvents()
	case "PROPERTY_ASSOCIATION":
		return u.AsPropertyAssociation()
	case "CONSTANT":
		return u.AsConstant()
	}
	return nil
}

func (u PublicAssociationFilterBranchFilterUnion) AsProperty() (v PublicPropertyFilter) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicAssociationFilterBranchFilterUnion) AsAssociation() (v PublicAssociationInListFilter) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicAssociationFilterBranchFilterUnion) AsPageView() (v PublicPageViewAnalyticsFilter) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicAssociationFilterBranchFilterUnion) AsCta() (v PublicCtaAnalyticsFilter) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicAssociationFilterBranchFilterUnion) AsEvent() (v PublicEventAnalyticsFilter) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicAssociationFilterBranchFilterUnion) AsFormSubmission() (v PublicFormSubmissionFilter) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicAssociationFilterBranchFilterUnion) AsFormSubmissionOnPage() (v PublicFormSubmissionOnPageFilter) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicAssociationFilterBranchFilterUnion) AsIntegrationEvent() (v PublicIntegrationEventFilter) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicAssociationFilterBranchFilterUnion) AsEmailSubscription() (v PublicEmailSubscriptionFilter) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicAssociationFilterBranchFilterUnion) AsCommunicationSubscription() (v PublicCommunicationSubscriptionFilter) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicAssociationFilterBranchFilterUnion) AsCampaignInfluenced() (v PublicCampaignInfluencedFilter) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicAssociationFilterBranchFilterUnion) AsSurveyMonkey() (v PublicSurveyMonkeyFilter) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicAssociationFilterBranchFilterUnion) AsSurveyMonkeyValue() (v PublicSurveyMonkeyValueFilter) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicAssociationFilterBranchFilterUnion) AsWebinar() (v PublicWebinarFilter) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicAssociationFilterBranchFilterUnion) AsEmailEvent() (v PublicEmailEventFilter) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicAssociationFilterBranchFilterUnion) AsPrivacy() (v PublicPrivacyAnalyticsFilter) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicAssociationFilterBranchFilterUnion) AsAdsSearch() (v PublicAdsSearchFilter) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicAssociationFilterBranchFilterUnion) AsAdsTime() (v PublicAdsTimeFilter) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicAssociationFilterBranchFilterUnion) AsInList() (v PublicInListFilter) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicAssociationFilterBranchFilterUnion) AsNumAssociations() (v PublicNumAssociationsFilter) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicAssociationFilterBranchFilterUnion) AsUnifiedEvents() (v PublicUnifiedEventsFilter) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicAssociationFilterBranchFilterUnion) AsPropertyAssociation() (v PublicPropertyAssociationInListFilter) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicAssociationFilterBranchFilterUnion) AsConstant() (v PublicConstantFilter) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u PublicAssociationFilterBranchFilterUnion) RawJSON() string { return u.JSON.raw }

func (r *PublicAssociationFilterBranchFilterUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PublicAssociationFilterBranchFilterUnionCoalescingRefineBy is an implicit
// subunion of [PublicAssociationFilterBranchFilterUnion].
// PublicAssociationFilterBranchFilterUnionCoalescingRefineBy provides convenient
// access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [PublicAssociationFilterBranchFilterUnion].
type PublicAssociationFilterBranchFilterUnionCoalescingRefineBy struct {
	Type string `json:"type"`
	// This field is from variant
	// [PublicAssociationInListFilterCoalescingRefineByUnion],
	// [PublicPageViewAnalyticsFilterCoalescingRefineByUnion],
	// [PublicCtaAnalyticsFilterCoalescingRefineByUnion],
	// [PublicEventAnalyticsFilterCoalescingRefineByUnion],
	// [PublicFormSubmissionFilterCoalescingRefineByUnion],
	// [PublicFormSubmissionOnPageFilterCoalescingRefineByUnion],
	// [PublicNumAssociationsFilterCoalescingRefineByUnion],
	// [PublicUnifiedEventsFilterCoalescingRefineByUnion],
	// [PublicPropertyAssociationInListFilterCoalescingRefineByUnion].
	MaxOccurrences int64 `json:"maxOccurrences"`
	// This field is from variant
	// [PublicAssociationInListFilterCoalescingRefineByUnion],
	// [PublicPageViewAnalyticsFilterCoalescingRefineByUnion],
	// [PublicCtaAnalyticsFilterCoalescingRefineByUnion],
	// [PublicEventAnalyticsFilterCoalescingRefineByUnion],
	// [PublicFormSubmissionFilterCoalescingRefineByUnion],
	// [PublicFormSubmissionOnPageFilterCoalescingRefineByUnion],
	// [PublicNumAssociationsFilterCoalescingRefineByUnion],
	// [PublicUnifiedEventsFilterCoalescingRefineByUnion],
	// [PublicPropertyAssociationInListFilterCoalescingRefineByUnion].
	MinOccurrences int64 `json:"minOccurrences"`
	// This field is from variant
	// [PublicAssociationInListFilterCoalescingRefineByUnion],
	// [PublicPageViewAnalyticsFilterCoalescingRefineByUnion],
	// [PublicCtaAnalyticsFilterCoalescingRefineByUnion],
	// [PublicEventAnalyticsFilterCoalescingRefineByUnion],
	// [PublicFormSubmissionFilterCoalescingRefineByUnion],
	// [PublicFormSubmissionOnPageFilterCoalescingRefineByUnion],
	// [PublicNumAssociationsFilterCoalescingRefineByUnion],
	// [PublicUnifiedEventsFilterCoalescingRefineByUnion],
	// [PublicPropertyAssociationInListFilterCoalescingRefineByUnion].
	SetType    string `json:"setType"`
	Comparison string `json:"comparison"`
	// This field is from variant
	// [PublicAssociationInListFilterCoalescingRefineByUnion],
	// [PublicPageViewAnalyticsFilterCoalescingRefineByUnion],
	// [PublicCtaAnalyticsFilterCoalescingRefineByUnion],
	// [PublicEventAnalyticsFilterCoalescingRefineByUnion],
	// [PublicFormSubmissionFilterCoalescingRefineByUnion],
	// [PublicFormSubmissionOnPageFilterCoalescingRefineByUnion],
	// [PublicNumAssociationsFilterCoalescingRefineByUnion],
	// [PublicUnifiedEventsFilterCoalescingRefineByUnion],
	// [PublicPropertyAssociationInListFilterCoalescingRefineByUnion].
	TimeOffset PublicTimeOffset `json:"timeOffset"`
	// This field is from variant
	// [PublicAssociationInListFilterCoalescingRefineByUnion],
	// [PublicPageViewAnalyticsFilterCoalescingRefineByUnion],
	// [PublicCtaAnalyticsFilterCoalescingRefineByUnion],
	// [PublicEventAnalyticsFilterCoalescingRefineByUnion],
	// [PublicFormSubmissionFilterCoalescingRefineByUnion],
	// [PublicFormSubmissionOnPageFilterCoalescingRefineByUnion],
	// [PublicNumAssociationsFilterCoalescingRefineByUnion],
	// [PublicUnifiedEventsFilterCoalescingRefineByUnion],
	// [PublicPropertyAssociationInListFilterCoalescingRefineByUnion].
	LowerBoundOffset PublicTimeOffset `json:"lowerBoundOffset"`
	RangeType        string           `json:"rangeType"`
	// This field is from variant
	// [PublicAssociationInListFilterCoalescingRefineByUnion],
	// [PublicPageViewAnalyticsFilterCoalescingRefineByUnion],
	// [PublicCtaAnalyticsFilterCoalescingRefineByUnion],
	// [PublicEventAnalyticsFilterCoalescingRefineByUnion],
	// [PublicFormSubmissionFilterCoalescingRefineByUnion],
	// [PublicFormSubmissionOnPageFilterCoalescingRefineByUnion],
	// [PublicNumAssociationsFilterCoalescingRefineByUnion],
	// [PublicUnifiedEventsFilterCoalescingRefineByUnion],
	// [PublicPropertyAssociationInListFilterCoalescingRefineByUnion].
	UpperBoundOffset PublicTimeOffset `json:"upperBoundOffset"`
	// This field is from variant
	// [PublicAssociationInListFilterCoalescingRefineByUnion],
	// [PublicPageViewAnalyticsFilterCoalescingRefineByUnion],
	// [PublicCtaAnalyticsFilterCoalescingRefineByUnion],
	// [PublicEventAnalyticsFilterCoalescingRefineByUnion],
	// [PublicFormSubmissionFilterCoalescingRefineByUnion],
	// [PublicFormSubmissionOnPageFilterCoalescingRefineByUnion],
	// [PublicNumAssociationsFilterCoalescingRefineByUnion],
	// [PublicUnifiedEventsFilterCoalescingRefineByUnion],
	// [PublicPropertyAssociationInListFilterCoalescingRefineByUnion].
	Timestamp int64 `json:"timestamp"`
	// This field is from variant
	// [PublicAssociationInListFilterCoalescingRefineByUnion],
	// [PublicPageViewAnalyticsFilterCoalescingRefineByUnion],
	// [PublicCtaAnalyticsFilterCoalescingRefineByUnion],
	// [PublicEventAnalyticsFilterCoalescingRefineByUnion],
	// [PublicFormSubmissionFilterCoalescingRefineByUnion],
	// [PublicFormSubmissionOnPageFilterCoalescingRefineByUnion],
	// [PublicNumAssociationsFilterCoalescingRefineByUnion],
	// [PublicUnifiedEventsFilterCoalescingRefineByUnion],
	// [PublicPropertyAssociationInListFilterCoalescingRefineByUnion].
	LowerTimestamp int64 `json:"lowerTimestamp"`
	// This field is from variant
	// [PublicAssociationInListFilterCoalescingRefineByUnion],
	// [PublicPageViewAnalyticsFilterCoalescingRefineByUnion],
	// [PublicCtaAnalyticsFilterCoalescingRefineByUnion],
	// [PublicEventAnalyticsFilterCoalescingRefineByUnion],
	// [PublicFormSubmissionFilterCoalescingRefineByUnion],
	// [PublicFormSubmissionOnPageFilterCoalescingRefineByUnion],
	// [PublicNumAssociationsFilterCoalescingRefineByUnion],
	// [PublicUnifiedEventsFilterCoalescingRefineByUnion],
	// [PublicPropertyAssociationInListFilterCoalescingRefineByUnion].
	UpperTimestamp               int64  `json:"upperTimestamp"`
	IncludeObjectsWithNoValueSet bool   `json:"includeObjectsWithNoValueSet"`
	OperationType                string `json:"operationType"`
	Operator                     string `json:"operator"`
	// This field is from variant
	// [PublicAssociationInListFilterCoalescingRefineByUnion],
	// [PublicPageViewAnalyticsFilterCoalescingRefineByUnion],
	// [PublicCtaAnalyticsFilterCoalescingRefineByUnion],
	// [PublicEventAnalyticsFilterCoalescingRefineByUnion],
	// [PublicFormSubmissionFilterCoalescingRefineByUnion],
	// [PublicFormSubmissionOnPageFilterCoalescingRefineByUnion],
	// [PublicNumAssociationsFilterCoalescingRefineByUnion],
	// [PublicUnifiedEventsFilterCoalescingRefineByUnion],
	// [PublicPropertyAssociationInListFilterCoalescingRefineByUnion].
	TimePoint PublicTimePointOperationTimePointUnion `json:"timePoint"`
	// This field is from variant
	// [PublicAssociationInListFilterCoalescingRefineByUnion],
	// [PublicPageViewAnalyticsFilterCoalescingRefineByUnion],
	// [PublicCtaAnalyticsFilterCoalescingRefineByUnion],
	// [PublicEventAnalyticsFilterCoalescingRefineByUnion],
	// [PublicFormSubmissionFilterCoalescingRefineByUnion],
	// [PublicFormSubmissionOnPageFilterCoalescingRefineByUnion],
	// [PublicNumAssociationsFilterCoalescingRefineByUnion],
	// [PublicUnifiedEventsFilterCoalescingRefineByUnion],
	// [PublicPropertyAssociationInListFilterCoalescingRefineByUnion].
	EndpointBehavior string `json:"endpointBehavior"`
	PropertyParser   string `json:"propertyParser"`
	// This field is from variant
	// [PublicAssociationInListFilterCoalescingRefineByUnion],
	// [PublicPageViewAnalyticsFilterCoalescingRefineByUnion],
	// [PublicCtaAnalyticsFilterCoalescingRefineByUnion],
	// [PublicEventAnalyticsFilterCoalescingRefineByUnion],
	// [PublicFormSubmissionFilterCoalescingRefineByUnion],
	// [PublicFormSubmissionOnPageFilterCoalescingRefineByUnion],
	// [PublicNumAssociationsFilterCoalescingRefineByUnion],
	// [PublicUnifiedEventsFilterCoalescingRefineByUnion],
	// [PublicPropertyAssociationInListFilterCoalescingRefineByUnion].
	LowerBoundTimePoint PublicRangedTimeOperationLowerBoundTimePointUnion `json:"lowerBoundTimePoint"`
	// This field is from variant
	// [PublicAssociationInListFilterCoalescingRefineByUnion],
	// [PublicPageViewAnalyticsFilterCoalescingRefineByUnion],
	// [PublicCtaAnalyticsFilterCoalescingRefineByUnion],
	// [PublicEventAnalyticsFilterCoalescingRefineByUnion],
	// [PublicFormSubmissionFilterCoalescingRefineByUnion],
	// [PublicFormSubmissionOnPageFilterCoalescingRefineByUnion],
	// [PublicNumAssociationsFilterCoalescingRefineByUnion],
	// [PublicUnifiedEventsFilterCoalescingRefineByUnion],
	// [PublicPropertyAssociationInListFilterCoalescingRefineByUnion].
	UpperBoundTimePoint PublicRangedTimeOperationUpperBoundTimePointUnion `json:"upperBoundTimePoint"`
	// This field is from variant
	// [PublicAssociationInListFilterCoalescingRefineByUnion],
	// [PublicPageViewAnalyticsFilterCoalescingRefineByUnion],
	// [PublicCtaAnalyticsFilterCoalescingRefineByUnion],
	// [PublicEventAnalyticsFilterCoalescingRefineByUnion],
	// [PublicFormSubmissionFilterCoalescingRefineByUnion],
	// [PublicFormSubmissionOnPageFilterCoalescingRefineByUnion],
	// [PublicNumAssociationsFilterCoalescingRefineByUnion],
	// [PublicUnifiedEventsFilterCoalescingRefineByUnion],
	// [PublicPropertyAssociationInListFilterCoalescingRefineByUnion].
	LowerBoundEndpointBehavior string `json:"lowerBoundEndpointBehavior"`
	// This field is from variant
	// [PublicAssociationInListFilterCoalescingRefineByUnion],
	// [PublicPageViewAnalyticsFilterCoalescingRefineByUnion],
	// [PublicCtaAnalyticsFilterCoalescingRefineByUnion],
	// [PublicEventAnalyticsFilterCoalescingRefineByUnion],
	// [PublicFormSubmissionFilterCoalescingRefineByUnion],
	// [PublicFormSubmissionOnPageFilterCoalescingRefineByUnion],
	// [PublicNumAssociationsFilterCoalescingRefineByUnion],
	// [PublicUnifiedEventsFilterCoalescingRefineByUnion],
	// [PublicPropertyAssociationInListFilterCoalescingRefineByUnion].
	UpperBoundEndpointBehavior string `json:"upperBoundEndpointBehavior"`
	JSON                       struct {
		Type                         respjson.Field
		MaxOccurrences               respjson.Field
		MinOccurrences               respjson.Field
		SetType                      respjson.Field
		Comparison                   respjson.Field
		TimeOffset                   respjson.Field
		LowerBoundOffset             respjson.Field
		RangeType                    respjson.Field
		UpperBoundOffset             respjson.Field
		Timestamp                    respjson.Field
		LowerTimestamp               respjson.Field
		UpperTimestamp               respjson.Field
		IncludeObjectsWithNoValueSet respjson.Field
		OperationType                respjson.Field
		Operator                     respjson.Field
		TimePoint                    respjson.Field
		EndpointBehavior             respjson.Field
		PropertyParser               respjson.Field
		LowerBoundTimePoint          respjson.Field
		UpperBoundTimePoint          respjson.Field
		LowerBoundEndpointBehavior   respjson.Field
		UpperBoundEndpointBehavior   respjson.Field
		raw                          string
	} `json:"-"`
}

func (r *PublicAssociationFilterBranchFilterUnionCoalescingRefineBy) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PublicAssociationFilterBranchFilterUnionPruningRefineBy is an implicit subunion
// of [PublicAssociationFilterBranchFilterUnion].
// PublicAssociationFilterBranchFilterUnionPruningRefineBy provides convenient
// access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [PublicAssociationFilterBranchFilterUnion].
type PublicAssociationFilterBranchFilterUnionPruningRefineBy struct {
	Type string `json:"type"`
	// This field is from variant [PublicPageViewAnalyticsFilterPruningRefineByUnion],
	// [PublicCtaAnalyticsFilterPruningRefineByUnion],
	// [PublicEventAnalyticsFilterPruningRefineByUnion],
	// [PublicFormSubmissionFilterPruningRefineByUnion],
	// [PublicFormSubmissionOnPageFilterPruningRefineByUnion],
	// [PublicEmailEventFilterPruningRefineByUnion],
	// [PublicAdsTimeFilterPruningRefineByUnion],
	// [PublicUnifiedEventsFilterPruningRefineByUnion].
	MaxOccurrences int64 `json:"maxOccurrences"`
	// This field is from variant [PublicPageViewAnalyticsFilterPruningRefineByUnion],
	// [PublicCtaAnalyticsFilterPruningRefineByUnion],
	// [PublicEventAnalyticsFilterPruningRefineByUnion],
	// [PublicFormSubmissionFilterPruningRefineByUnion],
	// [PublicFormSubmissionOnPageFilterPruningRefineByUnion],
	// [PublicEmailEventFilterPruningRefineByUnion],
	// [PublicAdsTimeFilterPruningRefineByUnion],
	// [PublicUnifiedEventsFilterPruningRefineByUnion].
	MinOccurrences int64 `json:"minOccurrences"`
	// This field is from variant [PublicPageViewAnalyticsFilterPruningRefineByUnion],
	// [PublicCtaAnalyticsFilterPruningRefineByUnion],
	// [PublicEventAnalyticsFilterPruningRefineByUnion],
	// [PublicFormSubmissionFilterPruningRefineByUnion],
	// [PublicFormSubmissionOnPageFilterPruningRefineByUnion],
	// [PublicEmailEventFilterPruningRefineByUnion],
	// [PublicAdsTimeFilterPruningRefineByUnion],
	// [PublicUnifiedEventsFilterPruningRefineByUnion].
	SetType    string `json:"setType"`
	Comparison string `json:"comparison"`
	// This field is from variant [PublicPageViewAnalyticsFilterPruningRefineByUnion],
	// [PublicCtaAnalyticsFilterPruningRefineByUnion],
	// [PublicEventAnalyticsFilterPruningRefineByUnion],
	// [PublicFormSubmissionFilterPruningRefineByUnion],
	// [PublicFormSubmissionOnPageFilterPruningRefineByUnion],
	// [PublicEmailEventFilterPruningRefineByUnion],
	// [PublicAdsTimeFilterPruningRefineByUnion],
	// [PublicUnifiedEventsFilterPruningRefineByUnion].
	TimeOffset PublicTimeOffset `json:"timeOffset"`
	// This field is from variant [PublicPageViewAnalyticsFilterPruningRefineByUnion],
	// [PublicCtaAnalyticsFilterPruningRefineByUnion],
	// [PublicEventAnalyticsFilterPruningRefineByUnion],
	// [PublicFormSubmissionFilterPruningRefineByUnion],
	// [PublicFormSubmissionOnPageFilterPruningRefineByUnion],
	// [PublicEmailEventFilterPruningRefineByUnion],
	// [PublicAdsTimeFilterPruningRefineByUnion],
	// [PublicUnifiedEventsFilterPruningRefineByUnion].
	LowerBoundOffset PublicTimeOffset `json:"lowerBoundOffset"`
	RangeType        string           `json:"rangeType"`
	// This field is from variant [PublicPageViewAnalyticsFilterPruningRefineByUnion],
	// [PublicCtaAnalyticsFilterPruningRefineByUnion],
	// [PublicEventAnalyticsFilterPruningRefineByUnion],
	// [PublicFormSubmissionFilterPruningRefineByUnion],
	// [PublicFormSubmissionOnPageFilterPruningRefineByUnion],
	// [PublicEmailEventFilterPruningRefineByUnion],
	// [PublicAdsTimeFilterPruningRefineByUnion],
	// [PublicUnifiedEventsFilterPruningRefineByUnion].
	UpperBoundOffset PublicTimeOffset `json:"upperBoundOffset"`
	// This field is from variant [PublicPageViewAnalyticsFilterPruningRefineByUnion],
	// [PublicCtaAnalyticsFilterPruningRefineByUnion],
	// [PublicEventAnalyticsFilterPruningRefineByUnion],
	// [PublicFormSubmissionFilterPruningRefineByUnion],
	// [PublicFormSubmissionOnPageFilterPruningRefineByUnion],
	// [PublicEmailEventFilterPruningRefineByUnion],
	// [PublicAdsTimeFilterPruningRefineByUnion],
	// [PublicUnifiedEventsFilterPruningRefineByUnion].
	Timestamp int64 `json:"timestamp"`
	// This field is from variant [PublicPageViewAnalyticsFilterPruningRefineByUnion],
	// [PublicCtaAnalyticsFilterPruningRefineByUnion],
	// [PublicEventAnalyticsFilterPruningRefineByUnion],
	// [PublicFormSubmissionFilterPruningRefineByUnion],
	// [PublicFormSubmissionOnPageFilterPruningRefineByUnion],
	// [PublicEmailEventFilterPruningRefineByUnion],
	// [PublicAdsTimeFilterPruningRefineByUnion],
	// [PublicUnifiedEventsFilterPruningRefineByUnion].
	LowerTimestamp int64 `json:"lowerTimestamp"`
	// This field is from variant [PublicPageViewAnalyticsFilterPruningRefineByUnion],
	// [PublicCtaAnalyticsFilterPruningRefineByUnion],
	// [PublicEventAnalyticsFilterPruningRefineByUnion],
	// [PublicFormSubmissionFilterPruningRefineByUnion],
	// [PublicFormSubmissionOnPageFilterPruningRefineByUnion],
	// [PublicEmailEventFilterPruningRefineByUnion],
	// [PublicAdsTimeFilterPruningRefineByUnion],
	// [PublicUnifiedEventsFilterPruningRefineByUnion].
	UpperTimestamp               int64  `json:"upperTimestamp"`
	IncludeObjectsWithNoValueSet bool   `json:"includeObjectsWithNoValueSet"`
	OperationType                string `json:"operationType"`
	Operator                     string `json:"operator"`
	// This field is from variant [PublicPageViewAnalyticsFilterPruningRefineByUnion],
	// [PublicCtaAnalyticsFilterPruningRefineByUnion],
	// [PublicEventAnalyticsFilterPruningRefineByUnion],
	// [PublicFormSubmissionFilterPruningRefineByUnion],
	// [PublicFormSubmissionOnPageFilterPruningRefineByUnion],
	// [PublicEmailEventFilterPruningRefineByUnion],
	// [PublicAdsTimeFilterPruningRefineByUnion],
	// [PublicUnifiedEventsFilterPruningRefineByUnion].
	TimePoint PublicTimePointOperationTimePointUnion `json:"timePoint"`
	// This field is from variant [PublicPageViewAnalyticsFilterPruningRefineByUnion],
	// [PublicCtaAnalyticsFilterPruningRefineByUnion],
	// [PublicEventAnalyticsFilterPruningRefineByUnion],
	// [PublicFormSubmissionFilterPruningRefineByUnion],
	// [PublicFormSubmissionOnPageFilterPruningRefineByUnion],
	// [PublicEmailEventFilterPruningRefineByUnion],
	// [PublicAdsTimeFilterPruningRefineByUnion],
	// [PublicUnifiedEventsFilterPruningRefineByUnion].
	EndpointBehavior string `json:"endpointBehavior"`
	PropertyParser   string `json:"propertyParser"`
	// This field is from variant [PublicPageViewAnalyticsFilterPruningRefineByUnion],
	// [PublicCtaAnalyticsFilterPruningRefineByUnion],
	// [PublicEventAnalyticsFilterPruningRefineByUnion],
	// [PublicFormSubmissionFilterPruningRefineByUnion],
	// [PublicFormSubmissionOnPageFilterPruningRefineByUnion],
	// [PublicEmailEventFilterPruningRefineByUnion],
	// [PublicAdsTimeFilterPruningRefineByUnion],
	// [PublicUnifiedEventsFilterPruningRefineByUnion].
	LowerBoundTimePoint PublicRangedTimeOperationLowerBoundTimePointUnion `json:"lowerBoundTimePoint"`
	// This field is from variant [PublicPageViewAnalyticsFilterPruningRefineByUnion],
	// [PublicCtaAnalyticsFilterPruningRefineByUnion],
	// [PublicEventAnalyticsFilterPruningRefineByUnion],
	// [PublicFormSubmissionFilterPruningRefineByUnion],
	// [PublicFormSubmissionOnPageFilterPruningRefineByUnion],
	// [PublicEmailEventFilterPruningRefineByUnion],
	// [PublicAdsTimeFilterPruningRefineByUnion],
	// [PublicUnifiedEventsFilterPruningRefineByUnion].
	UpperBoundTimePoint PublicRangedTimeOperationUpperBoundTimePointUnion `json:"upperBoundTimePoint"`
	// This field is from variant [PublicPageViewAnalyticsFilterPruningRefineByUnion],
	// [PublicCtaAnalyticsFilterPruningRefineByUnion],
	// [PublicEventAnalyticsFilterPruningRefineByUnion],
	// [PublicFormSubmissionFilterPruningRefineByUnion],
	// [PublicFormSubmissionOnPageFilterPruningRefineByUnion],
	// [PublicEmailEventFilterPruningRefineByUnion],
	// [PublicAdsTimeFilterPruningRefineByUnion],
	// [PublicUnifiedEventsFilterPruningRefineByUnion].
	LowerBoundEndpointBehavior string `json:"lowerBoundEndpointBehavior"`
	// This field is from variant [PublicPageViewAnalyticsFilterPruningRefineByUnion],
	// [PublicCtaAnalyticsFilterPruningRefineByUnion],
	// [PublicEventAnalyticsFilterPruningRefineByUnion],
	// [PublicFormSubmissionFilterPruningRefineByUnion],
	// [PublicFormSubmissionOnPageFilterPruningRefineByUnion],
	// [PublicEmailEventFilterPruningRefineByUnion],
	// [PublicAdsTimeFilterPruningRefineByUnion],
	// [PublicUnifiedEventsFilterPruningRefineByUnion].
	UpperBoundEndpointBehavior string `json:"upperBoundEndpointBehavior"`
	JSON                       struct {
		Type                         respjson.Field
		MaxOccurrences               respjson.Field
		MinOccurrences               respjson.Field
		SetType                      respjson.Field
		Comparison                   respjson.Field
		TimeOffset                   respjson.Field
		LowerBoundOffset             respjson.Field
		RangeType                    respjson.Field
		UpperBoundOffset             respjson.Field
		Timestamp                    respjson.Field
		LowerTimestamp               respjson.Field
		UpperTimestamp               respjson.Field
		IncludeObjectsWithNoValueSet respjson.Field
		OperationType                respjson.Field
		Operator                     respjson.Field
		TimePoint                    respjson.Field
		EndpointBehavior             respjson.Field
		PropertyParser               respjson.Field
		LowerBoundTimePoint          respjson.Field
		UpperBoundTimePoint          respjson.Field
		LowerBoundEndpointBehavior   respjson.Field
		UpperBoundEndpointBehavior   respjson.Field
		raw                          string
	} `json:"-"`
}

func (r *PublicAssociationFilterBranchFilterUnionPruningRefineBy) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PublicAssociationFilterBranchFilterUnionEventTypeID is an implicit subunion of
// [PublicAssociationFilterBranchFilterUnion].
// PublicAssociationFilterBranchFilterUnionEventTypeID provides convenient access
// to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [PublicAssociationFilterBranchFilterUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfInt OfString]
type PublicAssociationFilterBranchFilterUnionEventTypeID struct {
	// This field will be present if the value is a [int64] instead of an object.
	OfInt int64 `json:",inline"`
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	JSON     struct {
		OfInt    respjson.Field
		OfString respjson.Field
		raw      string
	} `json:"-"`
}

func (r *PublicAssociationFilterBranchFilterUnionEventTypeID) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties AssociationCategory, AssociationTypeID, FilterBranches,
// FilterBranchOperator, FilterBranchType, Filters, ObjectTypeID, Operator are
// required.
type PublicAssociationFilterBranchParam struct {
	// Specifies the category of the association for the filter branch
	// (HUBSPOT_DEFINED, USER_DEFINED, INTEGRATOR_DEFINED, WORK).
	AssociationCategory string `json:"associationCategory" api:"required"`
	// Type id of the association
	AssociationTypeID int64                                                 `json:"associationTypeId" api:"required"`
	FilterBranches    []PublicAssociationFilterBranchFilterBranchUnionParam `json:"filterBranches,omitzero" api:"required"`
	// Filter branch operator (AND)
	FilterBranchOperator string `json:"filterBranchOperator" api:"required"`
	// Type of the filter branch (ASSOCIATION)
	//
	// Any of "ASSOCIATION".
	FilterBranchType PublicAssociationFilterBranchFilterBranchType   `json:"filterBranchType,omitzero" api:"required"`
	Filters          []PublicAssociationFilterBranchFilterUnionParam `json:"filters,omitzero" api:"required"`
	// The ID representing the type of object associated with the filter branch.
	ObjectTypeID string `json:"objectTypeId" api:"required"`
	// Defines the operation to be applied within the filter branch (IN_LIST,
	// NOT_IN_LIST).
	Operator string `json:"operator" api:"required"`
	paramObj
}

func (r PublicAssociationFilterBranchParam) MarshalJSON() (data []byte, err error) {
	type shadow PublicAssociationFilterBranchParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PublicAssociationFilterBranchParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type PublicAssociationFilterBranchFilterBranchUnionParam struct {
	OfOr                  *PublicOrFilterBranchParam                  `json:",omitzero,inline"`
	OfAnd                 *PublicAndFilterBranchParam                 `json:",omitzero,inline"`
	OfNotAll              *PublicNotAllFilterBranchParam              `json:",omitzero,inline"`
	OfNotAny              *PublicNotAnyFilterBranchParam              `json:",omitzero,inline"`
	OfRestricted          *PublicRestrictedFilterBranchParam          `json:",omitzero,inline"`
	OfUnifiedEvents       *PublicUnifiedEventsFilterBranchParam       `json:",omitzero,inline"`
	OfPropertyAssociation *PublicPropertyAssociationFilterBranchParam `json:",omitzero,inline"`
	OfAssociation         *PublicAssociationFilterBranchParam         `json:",omitzero,inline"`
	paramUnion
}

func (u PublicAssociationFilterBranchFilterBranchUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfOr,
		u.OfAnd,
		u.OfNotAll,
		u.OfNotAny,
		u.OfRestricted,
		u.OfUnifiedEvents,
		u.OfPropertyAssociation,
		u.OfAssociation)
}
func (u *PublicAssociationFilterBranchFilterBranchUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func init() {
	apijson.RegisterUnion[PublicAssociationFilterBranchFilterBranchUnionParam](
		"filterBranchType",
		apijson.Discriminator[PublicOrFilterBranchParam]("OR"),
		apijson.Discriminator[PublicAndFilterBranchParam]("AND"),
		apijson.Discriminator[PublicNotAllFilterBranchParam]("NOT_ALL"),
		apijson.Discriminator[PublicNotAnyFilterBranchParam]("NOT_ANY"),
		apijson.Discriminator[PublicRestrictedFilterBranchParam]("RESTRICTED"),
		apijson.Discriminator[PublicUnifiedEventsFilterBranchParam]("UNIFIED_EVENTS"),
		apijson.Discriminator[PublicPropertyAssociationFilterBranchParam]("PROPERTY_ASSOCIATION"),
		apijson.Discriminator[PublicAssociationFilterBranchParam]("ASSOCIATION"),
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type PublicAssociationFilterBranchFilterUnionParam struct {
	OfProperty                  *PublicPropertyFilterParam                  `json:",omitzero,inline"`
	OfAssociation               *PublicAssociationInListFilterParam         `json:",omitzero,inline"`
	OfPageView                  *PublicPageViewAnalyticsFilterParam         `json:",omitzero,inline"`
	OfCta                       *PublicCtaAnalyticsFilterParam              `json:",omitzero,inline"`
	OfEvent                     *PublicEventAnalyticsFilterParam            `json:",omitzero,inline"`
	OfFormSubmission            *PublicFormSubmissionFilterParam            `json:",omitzero,inline"`
	OfFormSubmissionOnPage      *PublicFormSubmissionOnPageFilterParam      `json:",omitzero,inline"`
	OfIntegrationEvent          *PublicIntegrationEventFilterParam          `json:",omitzero,inline"`
	OfEmailSubscription         *PublicEmailSubscriptionFilterParam         `json:",omitzero,inline"`
	OfCommunicationSubscription *PublicCommunicationSubscriptionFilterParam `json:",omitzero,inline"`
	OfCampaignInfluenced        *PublicCampaignInfluencedFilterParam        `json:",omitzero,inline"`
	OfSurveyMonkey              *PublicSurveyMonkeyFilterParam              `json:",omitzero,inline"`
	OfSurveyMonkeyValue         *PublicSurveyMonkeyValueFilterParam         `json:",omitzero,inline"`
	OfWebinar                   *PublicWebinarFilterParam                   `json:",omitzero,inline"`
	OfEmailEvent                *PublicEmailEventFilterParam                `json:",omitzero,inline"`
	OfPrivacy                   *PublicPrivacyAnalyticsFilterParam          `json:",omitzero,inline"`
	OfAdsSearch                 *PublicAdsSearchFilterParam                 `json:",omitzero,inline"`
	OfAdsTime                   *PublicAdsTimeFilterParam                   `json:",omitzero,inline"`
	OfInList                    *PublicInListFilterParam                    `json:",omitzero,inline"`
	OfNumAssociations           *PublicNumAssociationsFilterParam           `json:",omitzero,inline"`
	OfUnifiedEvents             *PublicUnifiedEventsFilterParam             `json:",omitzero,inline"`
	OfPropertyAssociation       *PublicPropertyAssociationInListFilterParam `json:",omitzero,inline"`
	OfConstant                  *PublicConstantFilterParam                  `json:",omitzero,inline"`
	paramUnion
}

func (u PublicAssociationFilterBranchFilterUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfProperty,
		u.OfAssociation,
		u.OfPageView,
		u.OfCta,
		u.OfEvent,
		u.OfFormSubmission,
		u.OfFormSubmissionOnPage,
		u.OfIntegrationEvent,
		u.OfEmailSubscription,
		u.OfCommunicationSubscription,
		u.OfCampaignInfluenced,
		u.OfSurveyMonkey,
		u.OfSurveyMonkeyValue,
		u.OfWebinar,
		u.OfEmailEvent,
		u.OfPrivacy,
		u.OfAdsSearch,
		u.OfAdsTime,
		u.OfInList,
		u.OfNumAssociations,
		u.OfUnifiedEvents,
		u.OfPropertyAssociation,
		u.OfConstant)
}
func (u *PublicAssociationFilterBranchFilterUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func init() {
	apijson.RegisterUnion[PublicAssociationFilterBranchFilterUnionParam](
		"filterType",
		apijson.Discriminator[PublicPropertyFilterParam]("PROPERTY"),
		apijson.Discriminator[PublicAssociationInListFilterParam]("ASSOCIATION"),
		apijson.Discriminator[PublicPageViewAnalyticsFilterParam]("PAGE_VIEW"),
		apijson.Discriminator[PublicCtaAnalyticsFilterParam]("CTA"),
		apijson.Discriminator[PublicEventAnalyticsFilterParam]("EVENT"),
		apijson.Discriminator[PublicFormSubmissionFilterParam]("FORM_SUBMISSION"),
		apijson.Discriminator[PublicFormSubmissionOnPageFilterParam]("FORM_SUBMISSION_ON_PAGE"),
		apijson.Discriminator[PublicIntegrationEventFilterParam]("INTEGRATION_EVENT"),
		apijson.Discriminator[PublicEmailSubscriptionFilterParam]("EMAIL_SUBSCRIPTION"),
		apijson.Discriminator[PublicCommunicationSubscriptionFilterParam]("COMMUNICATION_SUBSCRIPTION"),
		apijson.Discriminator[PublicCampaignInfluencedFilterParam]("CAMPAIGN_INFLUENCED"),
		apijson.Discriminator[PublicSurveyMonkeyFilterParam]("SURVEY_MONKEY"),
		apijson.Discriminator[PublicSurveyMonkeyValueFilterParam]("SURVEY_MONKEY_VALUE"),
		apijson.Discriminator[PublicWebinarFilterParam]("WEBINAR"),
		apijson.Discriminator[PublicEmailEventFilterParam]("EMAIL_EVENT"),
		apijson.Discriminator[PublicPrivacyAnalyticsFilterParam]("PRIVACY"),
		apijson.Discriminator[PublicAdsSearchFilterParam]("ADS_SEARCH"),
		apijson.Discriminator[PublicAdsTimeFilterParam]("ADS_TIME"),
		apijson.Discriminator[PublicInListFilterParam]("IN_LIST"),
		apijson.Discriminator[PublicNumAssociationsFilterParam]("NUM_ASSOCIATIONS"),
		apijson.Discriminator[PublicUnifiedEventsFilterParam]("UNIFIED_EVENTS"),
		apijson.Discriminator[PublicPropertyAssociationInListFilterParam]("PROPERTY_ASSOCIATION"),
		apijson.Discriminator[PublicConstantFilterParam]("CONSTANT"),
	)
}

type PublicAssociationInListFilter struct {
	// Defines the category of the association, such as (HUBSPOT_DEFINED, USER_DEFINED,
	// INTEGRATOR_DEFINED, WORK).
	AssociationCategory string `json:"associationCategory" api:"required"`
	// The ID representing the type of association being filtered.
	AssociationTypeID int64 `json:"associationTypeId" api:"required"`
	// Specifies the criteria for refining the association filter.
	CoalescingRefineBy PublicAssociationInListFilterCoalescingRefineByUnion `json:"coalescingRefineBy" api:"required"`
	// Indicates the type of filter being applied, which is 'ASSOCIATION' by default.
	//
	// Any of "ASSOCIATION".
	FilterType PublicAssociationInListFilterFilterType `json:"filterType" api:"required"`
	// The ID of the list used in the association filter.
	ListID string `json:"listId" api:"required"`
	// Specifies the operation to be performed by the filter, such as 'IN_LIST' or
	// 'NOT_IN_LIST'.
	Operator string `json:"operator" api:"required"`
	// The type of object that the association filter is targeting.
	ToObjectType string `json:"toObjectType"`
	// The ID representing the type of object that the association filter is targeting.
	ToObjectTypeID string `json:"toObjectTypeId"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AssociationCategory respjson.Field
		AssociationTypeID   respjson.Field
		CoalescingRefineBy  respjson.Field
		FilterType          respjson.Field
		ListID              respjson.Field
		Operator            respjson.Field
		ToObjectType        respjson.Field
		ToObjectTypeID      respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PublicAssociationInListFilter) RawJSON() string { return r.JSON.raw }
func (r *PublicAssociationInListFilter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this PublicAssociationInListFilter to a
// PublicAssociationInListFilterParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// PublicAssociationInListFilterParam.Overrides()
func (r PublicAssociationInListFilter) ToParam() PublicAssociationInListFilterParam {
	return param.Override[PublicAssociationInListFilterParam](json.RawMessage(r.RawJSON()))
}

// PublicAssociationInListFilterCoalescingRefineByUnion contains all possible
// properties and values from [PublicNumOccurrencesRefineBy],
// [PublicSetOccurrencesRefineBy], [PublicRelativeComparativeTimestampRefineBy],
// [PublicRelativeRangedTimestampRefineBy],
// [PublicAbsoluteComparativeTimestampRefineBy],
// [PublicAbsoluteRangedTimestampRefineBy], [PublicAllHistoryRefineBy],
// [PublicTimePointOperation], [PublicRangedTimeOperation].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type PublicAssociationInListFilterCoalescingRefineByUnion struct {
	Type string `json:"type"`
	// This field is from variant [PublicNumOccurrencesRefineBy].
	MaxOccurrences int64 `json:"maxOccurrences"`
	// This field is from variant [PublicNumOccurrencesRefineBy].
	MinOccurrences int64 `json:"minOccurrences"`
	// This field is from variant [PublicSetOccurrencesRefineBy].
	SetType    string `json:"setType"`
	Comparison string `json:"comparison"`
	// This field is from variant [PublicRelativeComparativeTimestampRefineBy].
	TimeOffset PublicTimeOffset `json:"timeOffset"`
	// This field is from variant [PublicRelativeRangedTimestampRefineBy].
	LowerBoundOffset PublicTimeOffset `json:"lowerBoundOffset"`
	RangeType        string           `json:"rangeType"`
	// This field is from variant [PublicRelativeRangedTimestampRefineBy].
	UpperBoundOffset PublicTimeOffset `json:"upperBoundOffset"`
	// This field is from variant [PublicAbsoluteComparativeTimestampRefineBy].
	Timestamp int64 `json:"timestamp"`
	// This field is from variant [PublicAbsoluteRangedTimestampRefineBy].
	LowerTimestamp int64 `json:"lowerTimestamp"`
	// This field is from variant [PublicAbsoluteRangedTimestampRefineBy].
	UpperTimestamp               int64  `json:"upperTimestamp"`
	IncludeObjectsWithNoValueSet bool   `json:"includeObjectsWithNoValueSet"`
	OperationType                string `json:"operationType"`
	Operator                     string `json:"operator"`
	// This field is from variant [PublicTimePointOperation].
	TimePoint PublicTimePointOperationTimePointUnion `json:"timePoint"`
	// This field is from variant [PublicTimePointOperation].
	EndpointBehavior string `json:"endpointBehavior"`
	PropertyParser   string `json:"propertyParser"`
	// This field is from variant [PublicRangedTimeOperation].
	LowerBoundTimePoint PublicRangedTimeOperationLowerBoundTimePointUnion `json:"lowerBoundTimePoint"`
	// This field is from variant [PublicRangedTimeOperation].
	UpperBoundTimePoint PublicRangedTimeOperationUpperBoundTimePointUnion `json:"upperBoundTimePoint"`
	// This field is from variant [PublicRangedTimeOperation].
	LowerBoundEndpointBehavior string `json:"lowerBoundEndpointBehavior"`
	// This field is from variant [PublicRangedTimeOperation].
	UpperBoundEndpointBehavior string `json:"upperBoundEndpointBehavior"`
	JSON                       struct {
		Type                         respjson.Field
		MaxOccurrences               respjson.Field
		MinOccurrences               respjson.Field
		SetType                      respjson.Field
		Comparison                   respjson.Field
		TimeOffset                   respjson.Field
		LowerBoundOffset             respjson.Field
		RangeType                    respjson.Field
		UpperBoundOffset             respjson.Field
		Timestamp                    respjson.Field
		LowerTimestamp               respjson.Field
		UpperTimestamp               respjson.Field
		IncludeObjectsWithNoValueSet respjson.Field
		OperationType                respjson.Field
		Operator                     respjson.Field
		TimePoint                    respjson.Field
		EndpointBehavior             respjson.Field
		PropertyParser               respjson.Field
		LowerBoundTimePoint          respjson.Field
		UpperBoundTimePoint          respjson.Field
		LowerBoundEndpointBehavior   respjson.Field
		UpperBoundEndpointBehavior   respjson.Field
		raw                          string
	} `json:"-"`
}

func (u PublicAssociationInListFilterCoalescingRefineByUnion) AsNumOccurrences() (v PublicNumOccurrencesRefineBy) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicAssociationInListFilterCoalescingRefineByUnion) AsSetOccurrences() (v PublicSetOccurrencesRefineBy) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicAssociationInListFilterCoalescingRefineByUnion) AsRelativeComparative() (v PublicRelativeComparativeTimestampRefineBy) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicAssociationInListFilterCoalescingRefineByUnion) AsRelativeRanged() (v PublicRelativeRangedTimestampRefineBy) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicAssociationInListFilterCoalescingRefineByUnion) AsAbsoluteComparative() (v PublicAbsoluteComparativeTimestampRefineBy) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicAssociationInListFilterCoalescingRefineByUnion) AsAbsoluteRanged() (v PublicAbsoluteRangedTimestampRefineBy) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicAssociationInListFilterCoalescingRefineByUnion) AsAllHistory() (v PublicAllHistoryRefineBy) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicAssociationInListFilterCoalescingRefineByUnion) AsTimePoint() (v PublicTimePointOperation) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicAssociationInListFilterCoalescingRefineByUnion) AsTimeRanged() (v PublicRangedTimeOperation) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u PublicAssociationInListFilterCoalescingRefineByUnion) RawJSON() string { return u.JSON.raw }

func (r *PublicAssociationInListFilterCoalescingRefineByUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Indicates the type of filter being applied, which is 'ASSOCIATION' by default.
type PublicAssociationInListFilterFilterType string

const (
	PublicAssociationInListFilterFilterTypeAssociation PublicAssociationInListFilterFilterType = "ASSOCIATION"
)

// The properties AssociationCategory, AssociationTypeID, CoalescingRefineBy,
// FilterType, ListID, Operator are required.
type PublicAssociationInListFilterParam struct {
	// Defines the category of the association, such as (HUBSPOT_DEFINED, USER_DEFINED,
	// INTEGRATOR_DEFINED, WORK).
	AssociationCategory string `json:"associationCategory" api:"required"`
	// The ID representing the type of association being filtered.
	AssociationTypeID int64 `json:"associationTypeId" api:"required"`
	// Specifies the criteria for refining the association filter.
	CoalescingRefineBy PublicAssociationInListFilterCoalescingRefineByUnionParam `json:"coalescingRefineBy,omitzero" api:"required"`
	// Indicates the type of filter being applied, which is 'ASSOCIATION' by default.
	//
	// Any of "ASSOCIATION".
	FilterType PublicAssociationInListFilterFilterType `json:"filterType,omitzero" api:"required"`
	// The ID of the list used in the association filter.
	ListID string `json:"listId" api:"required"`
	// Specifies the operation to be performed by the filter, such as 'IN_LIST' or
	// 'NOT_IN_LIST'.
	Operator string `json:"operator" api:"required"`
	// The type of object that the association filter is targeting.
	ToObjectType param.Opt[string] `json:"toObjectType,omitzero"`
	// The ID representing the type of object that the association filter is targeting.
	ToObjectTypeID param.Opt[string] `json:"toObjectTypeId,omitzero"`
	paramObj
}

func (r PublicAssociationInListFilterParam) MarshalJSON() (data []byte, err error) {
	type shadow PublicAssociationInListFilterParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PublicAssociationInListFilterParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type PublicAssociationInListFilterCoalescingRefineByUnionParam struct {
	OfNumOccurrences      *PublicNumOccurrencesRefineByParam               `json:",omitzero,inline"`
	OfSetOccurrences      *PublicSetOccurrencesRefineByParam               `json:",omitzero,inline"`
	OfRelativeComparative *PublicRelativeComparativeTimestampRefineByParam `json:",omitzero,inline"`
	OfRelativeRanged      *PublicRelativeRangedTimestampRefineByParam      `json:",omitzero,inline"`
	OfAbsoluteComparative *PublicAbsoluteComparativeTimestampRefineByParam `json:",omitzero,inline"`
	OfAbsoluteRanged      *PublicAbsoluteRangedTimestampRefineByParam      `json:",omitzero,inline"`
	OfAllHistory          *PublicAllHistoryRefineByParam                   `json:",omitzero,inline"`
	OfTimePoint           *PublicTimePointOperationParam                   `json:",omitzero,inline"`
	OfTimeRanged          *PublicRangedTimeOperationParam                  `json:",omitzero,inline"`
	paramUnion
}

func (u PublicAssociationInListFilterCoalescingRefineByUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfNumOccurrences,
		u.OfSetOccurrences,
		u.OfRelativeComparative,
		u.OfRelativeRanged,
		u.OfAbsoluteComparative,
		u.OfAbsoluteRanged,
		u.OfAllHistory,
		u.OfTimePoint,
		u.OfTimeRanged)
}
func (u *PublicAssociationInListFilterCoalescingRefineByUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

type PublicBatchMigrationMapping struct {
	LegacyListIDsToIDsMapping []PublicMigrationMapping `json:"legacyListIdsToIdsMapping" api:"required"`
	// A list of legacy list ids that were passed in but not found. It will be empty if
	// no id's are missing
	MissingLegacyListIDs []string `json:"missingLegacyListIds" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		LegacyListIDsToIDsMapping respjson.Field
		MissingLegacyListIDs      respjson.Field
		ExtraFields               map[string]respjson.Field
		raw                       string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PublicBatchMigrationMapping) RawJSON() string { return r.JSON.raw }
func (r *PublicBatchMigrationMapping) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PublicBoolPropertyOperation struct {
	// Indicates whether objects with no value set for the property should be included
	// in the operation.
	IncludeObjectsWithNoValueSet bool `json:"includeObjectsWithNoValueSet" api:"required"`
	// Specifies the type of operation (BOOL).
	//
	// Any of "BOOL".
	OperationType PublicBoolPropertyOperationOperationType `json:"operationType" api:"required"`
	// Defines the operation to be applied in the boolean property operation
	// (IS_EQUAL_TO, IS_NOT_EQUAL_TO, HAS_EVER_BEEN_EQUAL_TO, HAS_NEVER_BEEN_EQUAL_TO).
	Operator string `json:"operator" api:"required"`
	// The boolean value to be used in the operation.
	Value bool `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		IncludeObjectsWithNoValueSet respjson.Field
		OperationType                respjson.Field
		Operator                     respjson.Field
		Value                        respjson.Field
		ExtraFields                  map[string]respjson.Field
		raw                          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PublicBoolPropertyOperation) RawJSON() string { return r.JSON.raw }
func (r *PublicBoolPropertyOperation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this PublicBoolPropertyOperation to a
// PublicBoolPropertyOperationParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// PublicBoolPropertyOperationParam.Overrides()
func (r PublicBoolPropertyOperation) ToParam() PublicBoolPropertyOperationParam {
	return param.Override[PublicBoolPropertyOperationParam](json.RawMessage(r.RawJSON()))
}

// Specifies the type of operation (BOOL).
type PublicBoolPropertyOperationOperationType string

const (
	PublicBoolPropertyOperationOperationTypeBool PublicBoolPropertyOperationOperationType = "BOOL"
)

// The properties IncludeObjectsWithNoValueSet, OperationType, Operator, Value are
// required.
type PublicBoolPropertyOperationParam struct {
	// Indicates whether objects with no value set for the property should be included
	// in the operation.
	IncludeObjectsWithNoValueSet bool `json:"includeObjectsWithNoValueSet" api:"required"`
	// Specifies the type of operation (BOOL).
	//
	// Any of "BOOL".
	OperationType PublicBoolPropertyOperationOperationType `json:"operationType,omitzero" api:"required"`
	// Defines the operation to be applied in the boolean property operation
	// (IS_EQUAL_TO, IS_NOT_EQUAL_TO, HAS_EVER_BEEN_EQUAL_TO, HAS_NEVER_BEEN_EQUAL_TO).
	Operator string `json:"operator" api:"required"`
	// The boolean value to be used in the operation.
	Value bool `json:"value" api:"required"`
	paramObj
}

func (r PublicBoolPropertyOperationParam) MarshalJSON() (data []byte, err error) {
	type shadow PublicBoolPropertyOperationParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PublicBoolPropertyOperationParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PublicCalendarDatePropertyOperation struct {
	// Indicates whether objects with no value set for the property should be included.
	IncludeObjectsWithNoValueSet bool `json:"includeObjectsWithNoValueSet" api:"required"`
	// The type of operation, which is (CALENDAR_DATE).
	//
	// Any of "CALENDAR_DATE".
	OperationType PublicCalendarDatePropertyOperationOperationType `json:"operationType" api:"required"`
	// Defines the operation to be applied to the calendar date property
	// (IN_THIS_TIME_UNIT, IN_THIS_TIME_UNIT_SO_FAR, IN_NEXT_TIME_UNIT,
	// IN_LAST_TIME_UNIT).
	Operator string `json:"operator" api:"required"`
	// The unit of time to be used in the operation (DAY, WEEK, MONTH, QUARTER, YEAR).
	TimeUnit string `json:"timeUnit" api:"required"`
	// The month in which the fiscal year starts.
	//
	// Any of "APRIL", "AUGUST", "DECEMBER", "FEBRUARY", "JANUARY", "JULY", "JUNE",
	// "MARCH", "MAY", "NOVEMBER", "OCTOBER", "SEPTEMBER".
	FiscalYearStart PublicCalendarDatePropertyOperationFiscalYearStart `json:"fiscalYearStart"`
	// The count of time units to be applied in the operation (1).
	TimeUnitCount int64 `json:"timeUnitCount"`
	// Specifies whether the fiscal year should be used in the operation.
	UseFiscalYear bool `json:"useFiscalYear"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		IncludeObjectsWithNoValueSet respjson.Field
		OperationType                respjson.Field
		Operator                     respjson.Field
		TimeUnit                     respjson.Field
		FiscalYearStart              respjson.Field
		TimeUnitCount                respjson.Field
		UseFiscalYear                respjson.Field
		ExtraFields                  map[string]respjson.Field
		raw                          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PublicCalendarDatePropertyOperation) RawJSON() string { return r.JSON.raw }
func (r *PublicCalendarDatePropertyOperation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this PublicCalendarDatePropertyOperation to a
// PublicCalendarDatePropertyOperationParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// PublicCalendarDatePropertyOperationParam.Overrides()
func (r PublicCalendarDatePropertyOperation) ToParam() PublicCalendarDatePropertyOperationParam {
	return param.Override[PublicCalendarDatePropertyOperationParam](json.RawMessage(r.RawJSON()))
}

// The type of operation, which is (CALENDAR_DATE).
type PublicCalendarDatePropertyOperationOperationType string

const (
	PublicCalendarDatePropertyOperationOperationTypeCalendarDate PublicCalendarDatePropertyOperationOperationType = "CALENDAR_DATE"
)

// The month in which the fiscal year starts.
type PublicCalendarDatePropertyOperationFiscalYearStart string

const (
	PublicCalendarDatePropertyOperationFiscalYearStartApril     PublicCalendarDatePropertyOperationFiscalYearStart = "APRIL"
	PublicCalendarDatePropertyOperationFiscalYearStartAugust    PublicCalendarDatePropertyOperationFiscalYearStart = "AUGUST"
	PublicCalendarDatePropertyOperationFiscalYearStartDecember  PublicCalendarDatePropertyOperationFiscalYearStart = "DECEMBER"
	PublicCalendarDatePropertyOperationFiscalYearStartFebruary  PublicCalendarDatePropertyOperationFiscalYearStart = "FEBRUARY"
	PublicCalendarDatePropertyOperationFiscalYearStartJanuary   PublicCalendarDatePropertyOperationFiscalYearStart = "JANUARY"
	PublicCalendarDatePropertyOperationFiscalYearStartJuly      PublicCalendarDatePropertyOperationFiscalYearStart = "JULY"
	PublicCalendarDatePropertyOperationFiscalYearStartJune      PublicCalendarDatePropertyOperationFiscalYearStart = "JUNE"
	PublicCalendarDatePropertyOperationFiscalYearStartMarch     PublicCalendarDatePropertyOperationFiscalYearStart = "MARCH"
	PublicCalendarDatePropertyOperationFiscalYearStartMay       PublicCalendarDatePropertyOperationFiscalYearStart = "MAY"
	PublicCalendarDatePropertyOperationFiscalYearStartNovember  PublicCalendarDatePropertyOperationFiscalYearStart = "NOVEMBER"
	PublicCalendarDatePropertyOperationFiscalYearStartOctober   PublicCalendarDatePropertyOperationFiscalYearStart = "OCTOBER"
	PublicCalendarDatePropertyOperationFiscalYearStartSeptember PublicCalendarDatePropertyOperationFiscalYearStart = "SEPTEMBER"
)

// The properties IncludeObjectsWithNoValueSet, OperationType, Operator, TimeUnit
// are required.
type PublicCalendarDatePropertyOperationParam struct {
	// Indicates whether objects with no value set for the property should be included.
	IncludeObjectsWithNoValueSet bool `json:"includeObjectsWithNoValueSet" api:"required"`
	// The type of operation, which is (CALENDAR_DATE).
	//
	// Any of "CALENDAR_DATE".
	OperationType PublicCalendarDatePropertyOperationOperationType `json:"operationType,omitzero" api:"required"`
	// Defines the operation to be applied to the calendar date property
	// (IN_THIS_TIME_UNIT, IN_THIS_TIME_UNIT_SO_FAR, IN_NEXT_TIME_UNIT,
	// IN_LAST_TIME_UNIT).
	Operator string `json:"operator" api:"required"`
	// The unit of time to be used in the operation (DAY, WEEK, MONTH, QUARTER, YEAR).
	TimeUnit string `json:"timeUnit" api:"required"`
	// The count of time units to be applied in the operation (1).
	TimeUnitCount param.Opt[int64] `json:"timeUnitCount,omitzero"`
	// Specifies whether the fiscal year should be used in the operation.
	UseFiscalYear param.Opt[bool] `json:"useFiscalYear,omitzero"`
	// The month in which the fiscal year starts.
	//
	// Any of "APRIL", "AUGUST", "DECEMBER", "FEBRUARY", "JANUARY", "JULY", "JUNE",
	// "MARCH", "MAY", "NOVEMBER", "OCTOBER", "SEPTEMBER".
	FiscalYearStart PublicCalendarDatePropertyOperationFiscalYearStart `json:"fiscalYearStart,omitzero"`
	paramObj
}

func (r PublicCalendarDatePropertyOperationParam) MarshalJSON() (data []byte, err error) {
	type shadow PublicCalendarDatePropertyOperationParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PublicCalendarDatePropertyOperationParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PublicCampaignInfluencedFilter struct {
	// The ID of the campaign that influences the filter.
	CampaignID string `json:"campaignId" api:"required"`
	// Indicates the type of filter (CAMPAIGN_INFLUENCED).
	//
	// Any of "CAMPAIGN_INFLUENCED".
	FilterType PublicCampaignInfluencedFilterFilterType `json:"filterType" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CampaignID  respjson.Field
		FilterType  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PublicCampaignInfluencedFilter) RawJSON() string { return r.JSON.raw }
func (r *PublicCampaignInfluencedFilter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this PublicCampaignInfluencedFilter to a
// PublicCampaignInfluencedFilterParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// PublicCampaignInfluencedFilterParam.Overrides()
func (r PublicCampaignInfluencedFilter) ToParam() PublicCampaignInfluencedFilterParam {
	return param.Override[PublicCampaignInfluencedFilterParam](json.RawMessage(r.RawJSON()))
}

// Indicates the type of filter (CAMPAIGN_INFLUENCED).
type PublicCampaignInfluencedFilterFilterType string

const (
	PublicCampaignInfluencedFilterFilterTypeCampaignInfluenced PublicCampaignInfluencedFilterFilterType = "CAMPAIGN_INFLUENCED"
)

// The properties CampaignID, FilterType are required.
type PublicCampaignInfluencedFilterParam struct {
	// The ID of the campaign that influences the filter.
	CampaignID string `json:"campaignId" api:"required"`
	// Indicates the type of filter (CAMPAIGN_INFLUENCED).
	//
	// Any of "CAMPAIGN_INFLUENCED".
	FilterType PublicCampaignInfluencedFilterFilterType `json:"filterType,omitzero" api:"required"`
	paramObj
}

func (r PublicCampaignInfluencedFilterParam) MarshalJSON() (data []byte, err error) {
	type shadow PublicCampaignInfluencedFilterParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PublicCampaignInfluencedFilterParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PublicCommunicationSubscriptionFilter struct {
	AcceptedOptStates []string `json:"acceptedOptStates" api:"required"`
	// Specifies the communication channel associated with the subscription filter
	// (EMAIL, WHATSAPP, SMS).
	Channel string `json:"channel" api:"required"`
	// Indicates the type of filter, which is (COMMUNICATION_SUBSCRIPTION)
	//
	// Any of "COMMUNICATION_SUBSCRIPTION".
	FilterType      PublicCommunicationSubscriptionFilterFilterType `json:"filterType" api:"required"`
	SubscriptionIDs []string                                        `json:"subscriptionIds" api:"required"`
	// Defines the type of subscription related to the filter (PORTAL_WIDE,
	// BUSINESS_UNIT_WIDE, INDIVIDUAL_SUBSCRIPTION)
	SubscriptionType string `json:"subscriptionType" api:"required"`
	// The ID of the business unit associated with the subscription filter.
	BusinessUnitID string `json:"businessUnitId"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AcceptedOptStates respjson.Field
		Channel           respjson.Field
		FilterType        respjson.Field
		SubscriptionIDs   respjson.Field
		SubscriptionType  respjson.Field
		BusinessUnitID    respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PublicCommunicationSubscriptionFilter) RawJSON() string { return r.JSON.raw }
func (r *PublicCommunicationSubscriptionFilter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this PublicCommunicationSubscriptionFilter to a
// PublicCommunicationSubscriptionFilterParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// PublicCommunicationSubscriptionFilterParam.Overrides()
func (r PublicCommunicationSubscriptionFilter) ToParam() PublicCommunicationSubscriptionFilterParam {
	return param.Override[PublicCommunicationSubscriptionFilterParam](json.RawMessage(r.RawJSON()))
}

// Indicates the type of filter, which is (COMMUNICATION_SUBSCRIPTION)
type PublicCommunicationSubscriptionFilterFilterType string

const (
	PublicCommunicationSubscriptionFilterFilterTypeCommunicationSubscription PublicCommunicationSubscriptionFilterFilterType = "COMMUNICATION_SUBSCRIPTION"
)

// The properties AcceptedOptStates, Channel, FilterType, SubscriptionIDs,
// SubscriptionType are required.
type PublicCommunicationSubscriptionFilterParam struct {
	AcceptedOptStates []string `json:"acceptedOptStates,omitzero" api:"required"`
	// Specifies the communication channel associated with the subscription filter
	// (EMAIL, WHATSAPP, SMS).
	Channel string `json:"channel" api:"required"`
	// Indicates the type of filter, which is (COMMUNICATION_SUBSCRIPTION)
	//
	// Any of "COMMUNICATION_SUBSCRIPTION".
	FilterType      PublicCommunicationSubscriptionFilterFilterType `json:"filterType,omitzero" api:"required"`
	SubscriptionIDs []string                                        `json:"subscriptionIds,omitzero" api:"required"`
	// Defines the type of subscription related to the filter (PORTAL_WIDE,
	// BUSINESS_UNIT_WIDE, INDIVIDUAL_SUBSCRIPTION)
	SubscriptionType string `json:"subscriptionType" api:"required"`
	// The ID of the business unit associated with the subscription filter.
	BusinessUnitID param.Opt[string] `json:"businessUnitId,omitzero"`
	paramObj
}

func (r PublicCommunicationSubscriptionFilterParam) MarshalJSON() (data []byte, err error) {
	type shadow PublicCommunicationSubscriptionFilterParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PublicCommunicationSubscriptionFilterParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PublicComparativeDatePropertyOperation struct {
	// The name of the property to compare against in the operation.
	ComparisonPropertyName string `json:"comparisonPropertyName" api:"required"`
	// Indicates whether objects with no value set for the property should be included
	// in the operation.
	IncludeObjectsWithNoValueSet bool `json:"includeObjectsWithNoValueSet" api:"required"`
	// The type of operation (COMPARATIVE_DATE).
	//
	// Any of "COMPARATIVE_DATE".
	OperationType PublicComparativeDatePropertyOperationOperationType `json:"operationType" api:"required"`
	// Defines the operation to be applied in the comparative date property operation
	// (IS_BEFORE, IS_AFTER).
	Operator string `json:"operator" api:"required"`
	// The default value used for comparison if the actual comparison property value is
	// not set.
	DefaultComparisonValue string `json:"defaultComparisonValue"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ComparisonPropertyName       respjson.Field
		IncludeObjectsWithNoValueSet respjson.Field
		OperationType                respjson.Field
		Operator                     respjson.Field
		DefaultComparisonValue       respjson.Field
		ExtraFields                  map[string]respjson.Field
		raw                          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PublicComparativeDatePropertyOperation) RawJSON() string { return r.JSON.raw }
func (r *PublicComparativeDatePropertyOperation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this PublicComparativeDatePropertyOperation to a
// PublicComparativeDatePropertyOperationParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// PublicComparativeDatePropertyOperationParam.Overrides()
func (r PublicComparativeDatePropertyOperation) ToParam() PublicComparativeDatePropertyOperationParam {
	return param.Override[PublicComparativeDatePropertyOperationParam](json.RawMessage(r.RawJSON()))
}

// The type of operation (COMPARATIVE_DATE).
type PublicComparativeDatePropertyOperationOperationType string

const (
	PublicComparativeDatePropertyOperationOperationTypeComparativeDate PublicComparativeDatePropertyOperationOperationType = "COMPARATIVE_DATE"
)

// The properties ComparisonPropertyName, IncludeObjectsWithNoValueSet,
// OperationType, Operator are required.
type PublicComparativeDatePropertyOperationParam struct {
	// The name of the property to compare against in the operation.
	ComparisonPropertyName string `json:"comparisonPropertyName" api:"required"`
	// Indicates whether objects with no value set for the property should be included
	// in the operation.
	IncludeObjectsWithNoValueSet bool `json:"includeObjectsWithNoValueSet" api:"required"`
	// The type of operation (COMPARATIVE_DATE).
	//
	// Any of "COMPARATIVE_DATE".
	OperationType PublicComparativeDatePropertyOperationOperationType `json:"operationType,omitzero" api:"required"`
	// Defines the operation to be applied in the comparative date property operation
	// (IS_BEFORE, IS_AFTER).
	Operator string `json:"operator" api:"required"`
	// The default value used for comparison if the actual comparison property value is
	// not set.
	DefaultComparisonValue param.Opt[string] `json:"defaultComparisonValue,omitzero"`
	paramObj
}

func (r PublicComparativeDatePropertyOperationParam) MarshalJSON() (data []byte, err error) {
	type shadow PublicComparativeDatePropertyOperationParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PublicComparativeDatePropertyOperationParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PublicComparativePropertyUpdatedOperation struct {
	// The name of the property to compare against in the operation.
	ComparisonPropertyName string `json:"comparisonPropertyName" api:"required"`
	// Indicates whether objects with no value set for the property should be included
	// in the operation.
	IncludeObjectsWithNoValueSet bool `json:"includeObjectsWithNoValueSet" api:"required"`
	// Specifies the type of operation (COMPARATIVE_PROPERTY_UPDATED).
	//
	// Any of "COMPARATIVE_PROPERTY_UPDATED".
	OperationType PublicComparativePropertyUpdatedOperationOperationType `json:"operationType" api:"required"`
	// Defines the operation to be applied, such as comparison operators (IS_BEFORE,
	// IS_AFTER).
	Operator string `json:"operator" api:"required"`
	// The default value used for comparison if the actual comparison property value is
	// not set.
	DefaultComparisonValue string `json:"defaultComparisonValue"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ComparisonPropertyName       respjson.Field
		IncludeObjectsWithNoValueSet respjson.Field
		OperationType                respjson.Field
		Operator                     respjson.Field
		DefaultComparisonValue       respjson.Field
		ExtraFields                  map[string]respjson.Field
		raw                          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PublicComparativePropertyUpdatedOperation) RawJSON() string { return r.JSON.raw }
func (r *PublicComparativePropertyUpdatedOperation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this PublicComparativePropertyUpdatedOperation to a
// PublicComparativePropertyUpdatedOperationParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// PublicComparativePropertyUpdatedOperationParam.Overrides()
func (r PublicComparativePropertyUpdatedOperation) ToParam() PublicComparativePropertyUpdatedOperationParam {
	return param.Override[PublicComparativePropertyUpdatedOperationParam](json.RawMessage(r.RawJSON()))
}

// Specifies the type of operation (COMPARATIVE_PROPERTY_UPDATED).
type PublicComparativePropertyUpdatedOperationOperationType string

const (
	PublicComparativePropertyUpdatedOperationOperationTypeComparativePropertyUpdated PublicComparativePropertyUpdatedOperationOperationType = "COMPARATIVE_PROPERTY_UPDATED"
)

// The properties ComparisonPropertyName, IncludeObjectsWithNoValueSet,
// OperationType, Operator are required.
type PublicComparativePropertyUpdatedOperationParam struct {
	// The name of the property to compare against in the operation.
	ComparisonPropertyName string `json:"comparisonPropertyName" api:"required"`
	// Indicates whether objects with no value set for the property should be included
	// in the operation.
	IncludeObjectsWithNoValueSet bool `json:"includeObjectsWithNoValueSet" api:"required"`
	// Specifies the type of operation (COMPARATIVE_PROPERTY_UPDATED).
	//
	// Any of "COMPARATIVE_PROPERTY_UPDATED".
	OperationType PublicComparativePropertyUpdatedOperationOperationType `json:"operationType,omitzero" api:"required"`
	// Defines the operation to be applied, such as comparison operators (IS_BEFORE,
	// IS_AFTER).
	Operator string `json:"operator" api:"required"`
	// The default value used for comparison if the actual comparison property value is
	// not set.
	DefaultComparisonValue param.Opt[string] `json:"defaultComparisonValue,omitzero"`
	paramObj
}

func (r PublicComparativePropertyUpdatedOperationParam) MarshalJSON() (data []byte, err error) {
	type shadow PublicComparativePropertyUpdatedOperationParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PublicComparativePropertyUpdatedOperationParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PublicConstantFilter struct {
	// Specifies the type of filter, which is (CONSTANT).
	//
	// Any of "CONSTANT".
	FilterType PublicConstantFilterFilterType `json:"filterType" api:"required"`
	// Indicates whether the filter should accept the condition.
	ShouldAccept bool `json:"shouldAccept" api:"required"`
	// Defines the source of the constant filter.
	Source string `json:"source"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FilterType   respjson.Field
		ShouldAccept respjson.Field
		Source       respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PublicConstantFilter) RawJSON() string { return r.JSON.raw }
func (r *PublicConstantFilter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this PublicConstantFilter to a PublicConstantFilterParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// PublicConstantFilterParam.Overrides()
func (r PublicConstantFilter) ToParam() PublicConstantFilterParam {
	return param.Override[PublicConstantFilterParam](json.RawMessage(r.RawJSON()))
}

// Specifies the type of filter, which is (CONSTANT).
type PublicConstantFilterFilterType string

const (
	PublicConstantFilterFilterTypeConstant PublicConstantFilterFilterType = "CONSTANT"
)

// The properties FilterType, ShouldAccept are required.
type PublicConstantFilterParam struct {
	// Specifies the type of filter, which is (CONSTANT).
	//
	// Any of "CONSTANT".
	FilterType PublicConstantFilterFilterType `json:"filterType,omitzero" api:"required"`
	// Indicates whether the filter should accept the condition.
	ShouldAccept bool `json:"shouldAccept" api:"required"`
	// Defines the source of the constant filter.
	Source param.Opt[string] `json:"source,omitzero"`
	paramObj
}

func (r PublicConstantFilterParam) MarshalJSON() (data []byte, err error) {
	type shadow PublicConstantFilterParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PublicConstantFilterParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PublicCtaAnalyticsFilter struct {
	// The name of the Call-to-Action (CTA) to be used in the filter.
	CtaName string `json:"ctaName" api:"required"`
	// Indicates the type of filter being applied, which is (CTA).
	//
	// Any of "CTA".
	FilterType PublicCtaAnalyticsFilterFilterType `json:"filterType" api:"required"`
	// Defines the operation to be applied within the filter (HAS_CLICKED_CTA,
	// HAS_NOT_CLICKED_CTA, HAS_OPENED_CTA, HAS_NOT_OPENED_CTA,
	// HAS_CLICKED_CTA_PLACEMENT, HAS_NOT_CLICKED_CTA_PLACEMENT,
	// HAS_OPENED_CTA_PLACEMENT, HAS_NOT_OPENED_CTA_PLACEMENT).
	Operator string `json:"operator" api:"required"`
	// Specifies the criteria for refining the filter by coalescing.
	CoalescingRefineBy PublicCtaAnalyticsFilterCoalescingRefineByUnion `json:"coalescingRefineBy"`
	// Specifies the criteria for refining the filter by pruning.
	PruningRefineBy PublicCtaAnalyticsFilterPruningRefineByUnion `json:"pruningRefineBy"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CtaName            respjson.Field
		FilterType         respjson.Field
		Operator           respjson.Field
		CoalescingRefineBy respjson.Field
		PruningRefineBy    respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PublicCtaAnalyticsFilter) RawJSON() string { return r.JSON.raw }
func (r *PublicCtaAnalyticsFilter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this PublicCtaAnalyticsFilter to a
// PublicCtaAnalyticsFilterParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// PublicCtaAnalyticsFilterParam.Overrides()
func (r PublicCtaAnalyticsFilter) ToParam() PublicCtaAnalyticsFilterParam {
	return param.Override[PublicCtaAnalyticsFilterParam](json.RawMessage(r.RawJSON()))
}

// Indicates the type of filter being applied, which is (CTA).
type PublicCtaAnalyticsFilterFilterType string

const (
	PublicCtaAnalyticsFilterFilterTypeCta PublicCtaAnalyticsFilterFilterType = "CTA"
)

// PublicCtaAnalyticsFilterCoalescingRefineByUnion contains all possible properties
// and values from [PublicNumOccurrencesRefineBy], [PublicSetOccurrencesRefineBy],
// [PublicRelativeComparativeTimestampRefineBy],
// [PublicRelativeRangedTimestampRefineBy],
// [PublicAbsoluteComparativeTimestampRefineBy],
// [PublicAbsoluteRangedTimestampRefineBy], [PublicAllHistoryRefineBy],
// [PublicTimePointOperation], [PublicRangedTimeOperation].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type PublicCtaAnalyticsFilterCoalescingRefineByUnion struct {
	Type string `json:"type"`
	// This field is from variant [PublicNumOccurrencesRefineBy].
	MaxOccurrences int64 `json:"maxOccurrences"`
	// This field is from variant [PublicNumOccurrencesRefineBy].
	MinOccurrences int64 `json:"minOccurrences"`
	// This field is from variant [PublicSetOccurrencesRefineBy].
	SetType    string `json:"setType"`
	Comparison string `json:"comparison"`
	// This field is from variant [PublicRelativeComparativeTimestampRefineBy].
	TimeOffset PublicTimeOffset `json:"timeOffset"`
	// This field is from variant [PublicRelativeRangedTimestampRefineBy].
	LowerBoundOffset PublicTimeOffset `json:"lowerBoundOffset"`
	RangeType        string           `json:"rangeType"`
	// This field is from variant [PublicRelativeRangedTimestampRefineBy].
	UpperBoundOffset PublicTimeOffset `json:"upperBoundOffset"`
	// This field is from variant [PublicAbsoluteComparativeTimestampRefineBy].
	Timestamp int64 `json:"timestamp"`
	// This field is from variant [PublicAbsoluteRangedTimestampRefineBy].
	LowerTimestamp int64 `json:"lowerTimestamp"`
	// This field is from variant [PublicAbsoluteRangedTimestampRefineBy].
	UpperTimestamp               int64  `json:"upperTimestamp"`
	IncludeObjectsWithNoValueSet bool   `json:"includeObjectsWithNoValueSet"`
	OperationType                string `json:"operationType"`
	Operator                     string `json:"operator"`
	// This field is from variant [PublicTimePointOperation].
	TimePoint PublicTimePointOperationTimePointUnion `json:"timePoint"`
	// This field is from variant [PublicTimePointOperation].
	EndpointBehavior string `json:"endpointBehavior"`
	PropertyParser   string `json:"propertyParser"`
	// This field is from variant [PublicRangedTimeOperation].
	LowerBoundTimePoint PublicRangedTimeOperationLowerBoundTimePointUnion `json:"lowerBoundTimePoint"`
	// This field is from variant [PublicRangedTimeOperation].
	UpperBoundTimePoint PublicRangedTimeOperationUpperBoundTimePointUnion `json:"upperBoundTimePoint"`
	// This field is from variant [PublicRangedTimeOperation].
	LowerBoundEndpointBehavior string `json:"lowerBoundEndpointBehavior"`
	// This field is from variant [PublicRangedTimeOperation].
	UpperBoundEndpointBehavior string `json:"upperBoundEndpointBehavior"`
	JSON                       struct {
		Type                         respjson.Field
		MaxOccurrences               respjson.Field
		MinOccurrences               respjson.Field
		SetType                      respjson.Field
		Comparison                   respjson.Field
		TimeOffset                   respjson.Field
		LowerBoundOffset             respjson.Field
		RangeType                    respjson.Field
		UpperBoundOffset             respjson.Field
		Timestamp                    respjson.Field
		LowerTimestamp               respjson.Field
		UpperTimestamp               respjson.Field
		IncludeObjectsWithNoValueSet respjson.Field
		OperationType                respjson.Field
		Operator                     respjson.Field
		TimePoint                    respjson.Field
		EndpointBehavior             respjson.Field
		PropertyParser               respjson.Field
		LowerBoundTimePoint          respjson.Field
		UpperBoundTimePoint          respjson.Field
		LowerBoundEndpointBehavior   respjson.Field
		UpperBoundEndpointBehavior   respjson.Field
		raw                          string
	} `json:"-"`
}

func (u PublicCtaAnalyticsFilterCoalescingRefineByUnion) AsNumOccurrences() (v PublicNumOccurrencesRefineBy) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicCtaAnalyticsFilterCoalescingRefineByUnion) AsSetOccurrences() (v PublicSetOccurrencesRefineBy) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicCtaAnalyticsFilterCoalescingRefineByUnion) AsRelativeComparative() (v PublicRelativeComparativeTimestampRefineBy) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicCtaAnalyticsFilterCoalescingRefineByUnion) AsRelativeRanged() (v PublicRelativeRangedTimestampRefineBy) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicCtaAnalyticsFilterCoalescingRefineByUnion) AsAbsoluteComparative() (v PublicAbsoluteComparativeTimestampRefineBy) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicCtaAnalyticsFilterCoalescingRefineByUnion) AsAbsoluteRanged() (v PublicAbsoluteRangedTimestampRefineBy) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicCtaAnalyticsFilterCoalescingRefineByUnion) AsAllHistory() (v PublicAllHistoryRefineBy) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicCtaAnalyticsFilterCoalescingRefineByUnion) AsTimePoint() (v PublicTimePointOperation) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicCtaAnalyticsFilterCoalescingRefineByUnion) AsTimeRanged() (v PublicRangedTimeOperation) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u PublicCtaAnalyticsFilterCoalescingRefineByUnion) RawJSON() string { return u.JSON.raw }

func (r *PublicCtaAnalyticsFilterCoalescingRefineByUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PublicCtaAnalyticsFilterPruningRefineByUnion contains all possible properties
// and values from [PublicNumOccurrencesRefineBy], [PublicSetOccurrencesRefineBy],
// [PublicRelativeComparativeTimestampRefineBy],
// [PublicRelativeRangedTimestampRefineBy],
// [PublicAbsoluteComparativeTimestampRefineBy],
// [PublicAbsoluteRangedTimestampRefineBy], [PublicAllHistoryRefineBy],
// [PublicTimePointOperation], [PublicRangedTimeOperation].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type PublicCtaAnalyticsFilterPruningRefineByUnion struct {
	Type string `json:"type"`
	// This field is from variant [PublicNumOccurrencesRefineBy].
	MaxOccurrences int64 `json:"maxOccurrences"`
	// This field is from variant [PublicNumOccurrencesRefineBy].
	MinOccurrences int64 `json:"minOccurrences"`
	// This field is from variant [PublicSetOccurrencesRefineBy].
	SetType    string `json:"setType"`
	Comparison string `json:"comparison"`
	// This field is from variant [PublicRelativeComparativeTimestampRefineBy].
	TimeOffset PublicTimeOffset `json:"timeOffset"`
	// This field is from variant [PublicRelativeRangedTimestampRefineBy].
	LowerBoundOffset PublicTimeOffset `json:"lowerBoundOffset"`
	RangeType        string           `json:"rangeType"`
	// This field is from variant [PublicRelativeRangedTimestampRefineBy].
	UpperBoundOffset PublicTimeOffset `json:"upperBoundOffset"`
	// This field is from variant [PublicAbsoluteComparativeTimestampRefineBy].
	Timestamp int64 `json:"timestamp"`
	// This field is from variant [PublicAbsoluteRangedTimestampRefineBy].
	LowerTimestamp int64 `json:"lowerTimestamp"`
	// This field is from variant [PublicAbsoluteRangedTimestampRefineBy].
	UpperTimestamp               int64  `json:"upperTimestamp"`
	IncludeObjectsWithNoValueSet bool   `json:"includeObjectsWithNoValueSet"`
	OperationType                string `json:"operationType"`
	Operator                     string `json:"operator"`
	// This field is from variant [PublicTimePointOperation].
	TimePoint PublicTimePointOperationTimePointUnion `json:"timePoint"`
	// This field is from variant [PublicTimePointOperation].
	EndpointBehavior string `json:"endpointBehavior"`
	PropertyParser   string `json:"propertyParser"`
	// This field is from variant [PublicRangedTimeOperation].
	LowerBoundTimePoint PublicRangedTimeOperationLowerBoundTimePointUnion `json:"lowerBoundTimePoint"`
	// This field is from variant [PublicRangedTimeOperation].
	UpperBoundTimePoint PublicRangedTimeOperationUpperBoundTimePointUnion `json:"upperBoundTimePoint"`
	// This field is from variant [PublicRangedTimeOperation].
	LowerBoundEndpointBehavior string `json:"lowerBoundEndpointBehavior"`
	// This field is from variant [PublicRangedTimeOperation].
	UpperBoundEndpointBehavior string `json:"upperBoundEndpointBehavior"`
	JSON                       struct {
		Type                         respjson.Field
		MaxOccurrences               respjson.Field
		MinOccurrences               respjson.Field
		SetType                      respjson.Field
		Comparison                   respjson.Field
		TimeOffset                   respjson.Field
		LowerBoundOffset             respjson.Field
		RangeType                    respjson.Field
		UpperBoundOffset             respjson.Field
		Timestamp                    respjson.Field
		LowerTimestamp               respjson.Field
		UpperTimestamp               respjson.Field
		IncludeObjectsWithNoValueSet respjson.Field
		OperationType                respjson.Field
		Operator                     respjson.Field
		TimePoint                    respjson.Field
		EndpointBehavior             respjson.Field
		PropertyParser               respjson.Field
		LowerBoundTimePoint          respjson.Field
		UpperBoundTimePoint          respjson.Field
		LowerBoundEndpointBehavior   respjson.Field
		UpperBoundEndpointBehavior   respjson.Field
		raw                          string
	} `json:"-"`
}

func (u PublicCtaAnalyticsFilterPruningRefineByUnion) AsNumOccurrences() (v PublicNumOccurrencesRefineBy) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicCtaAnalyticsFilterPruningRefineByUnion) AsSetOccurrences() (v PublicSetOccurrencesRefineBy) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicCtaAnalyticsFilterPruningRefineByUnion) AsRelativeComparative() (v PublicRelativeComparativeTimestampRefineBy) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicCtaAnalyticsFilterPruningRefineByUnion) AsRelativeRanged() (v PublicRelativeRangedTimestampRefineBy) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicCtaAnalyticsFilterPruningRefineByUnion) AsAbsoluteComparative() (v PublicAbsoluteComparativeTimestampRefineBy) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicCtaAnalyticsFilterPruningRefineByUnion) AsAbsoluteRanged() (v PublicAbsoluteRangedTimestampRefineBy) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicCtaAnalyticsFilterPruningRefineByUnion) AsAllHistory() (v PublicAllHistoryRefineBy) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicCtaAnalyticsFilterPruningRefineByUnion) AsTimePoint() (v PublicTimePointOperation) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicCtaAnalyticsFilterPruningRefineByUnion) AsTimeRanged() (v PublicRangedTimeOperation) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u PublicCtaAnalyticsFilterPruningRefineByUnion) RawJSON() string { return u.JSON.raw }

func (r *PublicCtaAnalyticsFilterPruningRefineByUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties CtaName, FilterType, Operator are required.
type PublicCtaAnalyticsFilterParam struct {
	// The name of the Call-to-Action (CTA) to be used in the filter.
	CtaName string `json:"ctaName" api:"required"`
	// Indicates the type of filter being applied, which is (CTA).
	//
	// Any of "CTA".
	FilterType PublicCtaAnalyticsFilterFilterType `json:"filterType,omitzero" api:"required"`
	// Defines the operation to be applied within the filter (HAS_CLICKED_CTA,
	// HAS_NOT_CLICKED_CTA, HAS_OPENED_CTA, HAS_NOT_OPENED_CTA,
	// HAS_CLICKED_CTA_PLACEMENT, HAS_NOT_CLICKED_CTA_PLACEMENT,
	// HAS_OPENED_CTA_PLACEMENT, HAS_NOT_OPENED_CTA_PLACEMENT).
	Operator string `json:"operator" api:"required"`
	// Specifies the criteria for refining the filter by coalescing.
	CoalescingRefineBy PublicCtaAnalyticsFilterCoalescingRefineByUnionParam `json:"coalescingRefineBy,omitzero"`
	// Specifies the criteria for refining the filter by pruning.
	PruningRefineBy PublicCtaAnalyticsFilterPruningRefineByUnionParam `json:"pruningRefineBy,omitzero"`
	paramObj
}

func (r PublicCtaAnalyticsFilterParam) MarshalJSON() (data []byte, err error) {
	type shadow PublicCtaAnalyticsFilterParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PublicCtaAnalyticsFilterParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type PublicCtaAnalyticsFilterCoalescingRefineByUnionParam struct {
	OfNumOccurrences      *PublicNumOccurrencesRefineByParam               `json:",omitzero,inline"`
	OfSetOccurrences      *PublicSetOccurrencesRefineByParam               `json:",omitzero,inline"`
	OfRelativeComparative *PublicRelativeComparativeTimestampRefineByParam `json:",omitzero,inline"`
	OfRelativeRanged      *PublicRelativeRangedTimestampRefineByParam      `json:",omitzero,inline"`
	OfAbsoluteComparative *PublicAbsoluteComparativeTimestampRefineByParam `json:",omitzero,inline"`
	OfAbsoluteRanged      *PublicAbsoluteRangedTimestampRefineByParam      `json:",omitzero,inline"`
	OfAllHistory          *PublicAllHistoryRefineByParam                   `json:",omitzero,inline"`
	OfTimePoint           *PublicTimePointOperationParam                   `json:",omitzero,inline"`
	OfTimeRanged          *PublicRangedTimeOperationParam                  `json:",omitzero,inline"`
	paramUnion
}

func (u PublicCtaAnalyticsFilterCoalescingRefineByUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfNumOccurrences,
		u.OfSetOccurrences,
		u.OfRelativeComparative,
		u.OfRelativeRanged,
		u.OfAbsoluteComparative,
		u.OfAbsoluteRanged,
		u.OfAllHistory,
		u.OfTimePoint,
		u.OfTimeRanged)
}
func (u *PublicCtaAnalyticsFilterCoalescingRefineByUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type PublicCtaAnalyticsFilterPruningRefineByUnionParam struct {
	OfNumOccurrences      *PublicNumOccurrencesRefineByParam               `json:",omitzero,inline"`
	OfSetOccurrences      *PublicSetOccurrencesRefineByParam               `json:",omitzero,inline"`
	OfRelativeComparative *PublicRelativeComparativeTimestampRefineByParam `json:",omitzero,inline"`
	OfRelativeRanged      *PublicRelativeRangedTimestampRefineByParam      `json:",omitzero,inline"`
	OfAbsoluteComparative *PublicAbsoluteComparativeTimestampRefineByParam `json:",omitzero,inline"`
	OfAbsoluteRanged      *PublicAbsoluteRangedTimestampRefineByParam      `json:",omitzero,inline"`
	OfAllHistory          *PublicAllHistoryRefineByParam                   `json:",omitzero,inline"`
	OfTimePoint           *PublicTimePointOperationParam                   `json:",omitzero,inline"`
	OfTimeRanged          *PublicRangedTimeOperationParam                  `json:",omitzero,inline"`
	paramUnion
}

func (u PublicCtaAnalyticsFilterPruningRefineByUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfNumOccurrences,
		u.OfSetOccurrences,
		u.OfRelativeComparative,
		u.OfRelativeRanged,
		u.OfAbsoluteComparative,
		u.OfAbsoluteRanged,
		u.OfAllHistory,
		u.OfTimePoint,
		u.OfTimeRanged)
}
func (u *PublicCtaAnalyticsFilterPruningRefineByUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

type PublicDatePoint struct {
	// The day component of the date.
	Day int64 `json:"day" api:"required"`
	// The month component of the date.
	Month int64 `json:"month" api:"required"`
	// Specifies the type of time (DATE).
	//
	// Any of "DATE".
	TimeType PublicDatePointTimeType `json:"timeType" api:"required"`
	// The year component of the date.
	Year int64 `json:"year" api:"required"`
	// The identifier for the time zone.
	ZoneID string `json:"zoneId" api:"required"`
	// The hour component of the time.
	Hour int64 `json:"hour"`
	// The millisecond component of the time.
	Millisecond int64 `json:"millisecond"`
	// The minute component of the time.
	Minute int64 `json:"minute"`
	// The second component of the time.
	Second int64 `json:"second"`
	// The source of the time zone information.
	TimezoneSource string `json:"timezoneSource"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Day            respjson.Field
		Month          respjson.Field
		TimeType       respjson.Field
		Year           respjson.Field
		ZoneID         respjson.Field
		Hour           respjson.Field
		Millisecond    respjson.Field
		Minute         respjson.Field
		Second         respjson.Field
		TimezoneSource respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PublicDatePoint) RawJSON() string { return r.JSON.raw }
func (r *PublicDatePoint) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this PublicDatePoint to a PublicDatePointParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// PublicDatePointParam.Overrides()
func (r PublicDatePoint) ToParam() PublicDatePointParam {
	return param.Override[PublicDatePointParam](json.RawMessage(r.RawJSON()))
}

// Specifies the type of time (DATE).
type PublicDatePointTimeType string

const (
	PublicDatePointTimeTypeDate PublicDatePointTimeType = "DATE"
)

// The properties Day, Month, TimeType, Year, ZoneID are required.
type PublicDatePointParam struct {
	// The day component of the date.
	Day int64 `json:"day" api:"required"`
	// The month component of the date.
	Month int64 `json:"month" api:"required"`
	// Specifies the type of time (DATE).
	//
	// Any of "DATE".
	TimeType PublicDatePointTimeType `json:"timeType,omitzero" api:"required"`
	// The year component of the date.
	Year int64 `json:"year" api:"required"`
	// The identifier for the time zone.
	ZoneID string `json:"zoneId" api:"required"`
	// The hour component of the time.
	Hour param.Opt[int64] `json:"hour,omitzero"`
	// The millisecond component of the time.
	Millisecond param.Opt[int64] `json:"millisecond,omitzero"`
	// The minute component of the time.
	Minute param.Opt[int64] `json:"minute,omitzero"`
	// The second component of the time.
	Second param.Opt[int64] `json:"second,omitzero"`
	// The source of the time zone information.
	TimezoneSource param.Opt[string] `json:"timezoneSource,omitzero"`
	paramObj
}

func (r PublicDatePointParam) MarshalJSON() (data []byte, err error) {
	type shadow PublicDatePointParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PublicDatePointParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PublicDatePropertyOperation struct {
	// The day of the month for the date operation.
	Day int64 `json:"day" api:"required"`
	// Indicates whether objects with no value set for the property should be included.
	IncludeObjectsWithNoValueSet bool `json:"includeObjectsWithNoValueSet" api:"required"`
	// The month for the date operation.
	Month string `json:"month" api:"required"`
	// Specifies the type of operation (DATE).
	//
	// Any of "DATE".
	OperationType PublicDatePropertyOperationOperationType `json:"operationType" api:"required"`
	// Defines the operation to be applied in the date property operation
	// (IS_LESS_THAN_X_DAYS_AGO, IS_MORE_THAN_X_DAYS_AGO, IS_LESS_THAN_X_DAYS_FROM_NOW,
	// IS_MORE_THAN_X_DAYS_FROM_NOW).
	Operator string `json:"operator" api:"required"`
	// The year for the date operation.
	Year int64 `json:"year" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Day                          respjson.Field
		IncludeObjectsWithNoValueSet respjson.Field
		Month                        respjson.Field
		OperationType                respjson.Field
		Operator                     respjson.Field
		Year                         respjson.Field
		ExtraFields                  map[string]respjson.Field
		raw                          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PublicDatePropertyOperation) RawJSON() string { return r.JSON.raw }
func (r *PublicDatePropertyOperation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this PublicDatePropertyOperation to a
// PublicDatePropertyOperationParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// PublicDatePropertyOperationParam.Overrides()
func (r PublicDatePropertyOperation) ToParam() PublicDatePropertyOperationParam {
	return param.Override[PublicDatePropertyOperationParam](json.RawMessage(r.RawJSON()))
}

// Specifies the type of operation (DATE).
type PublicDatePropertyOperationOperationType string

const (
	PublicDatePropertyOperationOperationTypeDate PublicDatePropertyOperationOperationType = "DATE"
)

// The properties Day, IncludeObjectsWithNoValueSet, Month, OperationType,
// Operator, Year are required.
type PublicDatePropertyOperationParam struct {
	// The day of the month for the date operation.
	Day int64 `json:"day" api:"required"`
	// Indicates whether objects with no value set for the property should be included.
	IncludeObjectsWithNoValueSet bool `json:"includeObjectsWithNoValueSet" api:"required"`
	// The month for the date operation.
	Month string `json:"month" api:"required"`
	// Specifies the type of operation (DATE).
	//
	// Any of "DATE".
	OperationType PublicDatePropertyOperationOperationType `json:"operationType,omitzero" api:"required"`
	// Defines the operation to be applied in the date property operation
	// (IS_LESS_THAN_X_DAYS_AGO, IS_MORE_THAN_X_DAYS_AGO, IS_LESS_THAN_X_DAYS_FROM_NOW,
	// IS_MORE_THAN_X_DAYS_FROM_NOW).
	Operator string `json:"operator" api:"required"`
	// The year for the date operation.
	Year int64 `json:"year" api:"required"`
	paramObj
}

func (r PublicDatePropertyOperationParam) MarshalJSON() (data []byte, err error) {
	type shadow PublicDatePropertyOperationParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PublicDatePropertyOperationParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PublicDateTimePropertyOperation struct {
	// Specifies whether objects without a set value should be included in the
	// operation.
	IncludeObjectsWithNoValueSet bool `json:"includeObjectsWithNoValueSet" api:"required"`
	// The type of operation (DATETIME).
	//
	// Any of "DATETIME".
	OperationType PublicDateTimePropertyOperationOperationType `json:"operationType" api:"required"`
	// Defines the operation to be applied, such as comparison operators (IS_BEFORE,
	// IS_AFTER).
	Operator string `json:"operator" api:"required"`
	// Indicates whether the timestamp requires conversion to a different time zone.
	RequiresTimeZoneConversion bool `json:"requiresTimeZoneConversion" api:"required"`
	// The specific point in time used in the operation.
	Timestamp int64 `json:"timestamp" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		IncludeObjectsWithNoValueSet respjson.Field
		OperationType                respjson.Field
		Operator                     respjson.Field
		RequiresTimeZoneConversion   respjson.Field
		Timestamp                    respjson.Field
		ExtraFields                  map[string]respjson.Field
		raw                          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PublicDateTimePropertyOperation) RawJSON() string { return r.JSON.raw }
func (r *PublicDateTimePropertyOperation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this PublicDateTimePropertyOperation to a
// PublicDateTimePropertyOperationParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// PublicDateTimePropertyOperationParam.Overrides()
func (r PublicDateTimePropertyOperation) ToParam() PublicDateTimePropertyOperationParam {
	return param.Override[PublicDateTimePropertyOperationParam](json.RawMessage(r.RawJSON()))
}

// The type of operation (DATETIME).
type PublicDateTimePropertyOperationOperationType string

const (
	PublicDateTimePropertyOperationOperationTypeDatetime PublicDateTimePropertyOperationOperationType = "DATETIME"
)

// The properties IncludeObjectsWithNoValueSet, OperationType, Operator,
// RequiresTimeZoneConversion, Timestamp are required.
type PublicDateTimePropertyOperationParam struct {
	// Specifies whether objects without a set value should be included in the
	// operation.
	IncludeObjectsWithNoValueSet bool `json:"includeObjectsWithNoValueSet" api:"required"`
	// The type of operation (DATETIME).
	//
	// Any of "DATETIME".
	OperationType PublicDateTimePropertyOperationOperationType `json:"operationType,omitzero" api:"required"`
	// Defines the operation to be applied, such as comparison operators (IS_BEFORE,
	// IS_AFTER).
	Operator string `json:"operator" api:"required"`
	// Indicates whether the timestamp requires conversion to a different time zone.
	RequiresTimeZoneConversion bool `json:"requiresTimeZoneConversion" api:"required"`
	// The specific point in time used in the operation.
	Timestamp int64 `json:"timestamp" api:"required"`
	paramObj
}

func (r PublicDateTimePropertyOperationParam) MarshalJSON() (data []byte, err error) {
	type shadow PublicDateTimePropertyOperationParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PublicDateTimePropertyOperationParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PublicEmailEventFilter struct {
	// The ID of the application associated with the email event filter.
	AppID string `json:"appId" api:"required"`
	// The ID of the email associated with the event filter.
	EmailID string `json:"emailId" api:"required"`
	// Indicates the type of filter (EMAIL_EVENT).
	//
	// Any of "EMAIL_EVENT".
	FilterType PublicEmailEventFilterFilterType `json:"filterType" api:"required"`
	// Specifies the level of the email event, such as EMAIL_API_CAMPAIGN_GROUP.
	Level string `json:"level" api:"required"`
	// Defines the operation to be applied within the filter (BOUNCED, LINK_CLICKED,
	// MARKED_SPAM, OPENED, OPENED_BUT_LINK_NOT_CLICKED, OPENED_BUT_NOT_REPLIED,
	// RECEIVED, RECEIVED_BUT_NOT_OPENED, REPLIED, SENT, SENT_BUT_LINK_NOT_CLICKED,
	// SENT_BUT_NOT_RECEIVED, UNSUBSCRIBED).
	//
	// Any of "BOUNCED", "LINK_CLICKED", "MARKED_SPAM", "OPENED",
	// "OPENED_BUT_LINK_NOT_CLICKED", "OPENED_BUT_NOT_REPLIED", "RECEIVED",
	// "RECEIVED_BUT_NOT_OPENED", "REPLIED", "SENT", "SENT_BUT_LINK_NOT_CLICKED",
	// "SENT_BUT_NOT_RECEIVED", "UNSUBSCRIBED".
	Operator PublicEmailEventFilterOperator `json:"operator" api:"required"`
	// The URL that was clicked in the email event.
	ClickURL string `json:"clickUrl"`
	// Specifies the criteria for refining the filter by pruning.
	PruningRefineBy PublicEmailEventFilterPruningRefineByUnion `json:"pruningRefineBy"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AppID           respjson.Field
		EmailID         respjson.Field
		FilterType      respjson.Field
		Level           respjson.Field
		Operator        respjson.Field
		ClickURL        respjson.Field
		PruningRefineBy respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PublicEmailEventFilter) RawJSON() string { return r.JSON.raw }
func (r *PublicEmailEventFilter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this PublicEmailEventFilter to a PublicEmailEventFilterParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// PublicEmailEventFilterParam.Overrides()
func (r PublicEmailEventFilter) ToParam() PublicEmailEventFilterParam {
	return param.Override[PublicEmailEventFilterParam](json.RawMessage(r.RawJSON()))
}

// Indicates the type of filter (EMAIL_EVENT).
type PublicEmailEventFilterFilterType string

const (
	PublicEmailEventFilterFilterTypeEmailEvent PublicEmailEventFilterFilterType = "EMAIL_EVENT"
)

// Defines the operation to be applied within the filter (BOUNCED, LINK_CLICKED,
// MARKED_SPAM, OPENED, OPENED_BUT_LINK_NOT_CLICKED, OPENED_BUT_NOT_REPLIED,
// RECEIVED, RECEIVED_BUT_NOT_OPENED, REPLIED, SENT, SENT_BUT_LINK_NOT_CLICKED,
// SENT_BUT_NOT_RECEIVED, UNSUBSCRIBED).
type PublicEmailEventFilterOperator string

const (
	PublicEmailEventFilterOperatorBounced                 PublicEmailEventFilterOperator = "BOUNCED"
	PublicEmailEventFilterOperatorLinkClicked             PublicEmailEventFilterOperator = "LINK_CLICKED"
	PublicEmailEventFilterOperatorMarkedSpam              PublicEmailEventFilterOperator = "MARKED_SPAM"
	PublicEmailEventFilterOperatorOpened                  PublicEmailEventFilterOperator = "OPENED"
	PublicEmailEventFilterOperatorOpenedButLinkNotClicked PublicEmailEventFilterOperator = "OPENED_BUT_LINK_NOT_CLICKED"
	PublicEmailEventFilterOperatorOpenedButNotReplied     PublicEmailEventFilterOperator = "OPENED_BUT_NOT_REPLIED"
	PublicEmailEventFilterOperatorReceived                PublicEmailEventFilterOperator = "RECEIVED"
	PublicEmailEventFilterOperatorReceivedButNotOpened    PublicEmailEventFilterOperator = "RECEIVED_BUT_NOT_OPENED"
	PublicEmailEventFilterOperatorReplied                 PublicEmailEventFilterOperator = "REPLIED"
	PublicEmailEventFilterOperatorSent                    PublicEmailEventFilterOperator = "SENT"
	PublicEmailEventFilterOperatorSentButLinkNotClicked   PublicEmailEventFilterOperator = "SENT_BUT_LINK_NOT_CLICKED"
	PublicEmailEventFilterOperatorSentButNotReceived      PublicEmailEventFilterOperator = "SENT_BUT_NOT_RECEIVED"
	PublicEmailEventFilterOperatorUnsubscribed            PublicEmailEventFilterOperator = "UNSUBSCRIBED"
)

// PublicEmailEventFilterPruningRefineByUnion contains all possible properties and
// values from [PublicNumOccurrencesRefineBy], [PublicSetOccurrencesRefineBy],
// [PublicRelativeComparativeTimestampRefineBy],
// [PublicRelativeRangedTimestampRefineBy],
// [PublicAbsoluteComparativeTimestampRefineBy],
// [PublicAbsoluteRangedTimestampRefineBy], [PublicAllHistoryRefineBy],
// [PublicTimePointOperation], [PublicRangedTimeOperation].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type PublicEmailEventFilterPruningRefineByUnion struct {
	Type string `json:"type"`
	// This field is from variant [PublicNumOccurrencesRefineBy].
	MaxOccurrences int64 `json:"maxOccurrences"`
	// This field is from variant [PublicNumOccurrencesRefineBy].
	MinOccurrences int64 `json:"minOccurrences"`
	// This field is from variant [PublicSetOccurrencesRefineBy].
	SetType    string `json:"setType"`
	Comparison string `json:"comparison"`
	// This field is from variant [PublicRelativeComparativeTimestampRefineBy].
	TimeOffset PublicTimeOffset `json:"timeOffset"`
	// This field is from variant [PublicRelativeRangedTimestampRefineBy].
	LowerBoundOffset PublicTimeOffset `json:"lowerBoundOffset"`
	RangeType        string           `json:"rangeType"`
	// This field is from variant [PublicRelativeRangedTimestampRefineBy].
	UpperBoundOffset PublicTimeOffset `json:"upperBoundOffset"`
	// This field is from variant [PublicAbsoluteComparativeTimestampRefineBy].
	Timestamp int64 `json:"timestamp"`
	// This field is from variant [PublicAbsoluteRangedTimestampRefineBy].
	LowerTimestamp int64 `json:"lowerTimestamp"`
	// This field is from variant [PublicAbsoluteRangedTimestampRefineBy].
	UpperTimestamp               int64  `json:"upperTimestamp"`
	IncludeObjectsWithNoValueSet bool   `json:"includeObjectsWithNoValueSet"`
	OperationType                string `json:"operationType"`
	Operator                     string `json:"operator"`
	// This field is from variant [PublicTimePointOperation].
	TimePoint PublicTimePointOperationTimePointUnion `json:"timePoint"`
	// This field is from variant [PublicTimePointOperation].
	EndpointBehavior string `json:"endpointBehavior"`
	PropertyParser   string `json:"propertyParser"`
	// This field is from variant [PublicRangedTimeOperation].
	LowerBoundTimePoint PublicRangedTimeOperationLowerBoundTimePointUnion `json:"lowerBoundTimePoint"`
	// This field is from variant [PublicRangedTimeOperation].
	UpperBoundTimePoint PublicRangedTimeOperationUpperBoundTimePointUnion `json:"upperBoundTimePoint"`
	// This field is from variant [PublicRangedTimeOperation].
	LowerBoundEndpointBehavior string `json:"lowerBoundEndpointBehavior"`
	// This field is from variant [PublicRangedTimeOperation].
	UpperBoundEndpointBehavior string `json:"upperBoundEndpointBehavior"`
	JSON                       struct {
		Type                         respjson.Field
		MaxOccurrences               respjson.Field
		MinOccurrences               respjson.Field
		SetType                      respjson.Field
		Comparison                   respjson.Field
		TimeOffset                   respjson.Field
		LowerBoundOffset             respjson.Field
		RangeType                    respjson.Field
		UpperBoundOffset             respjson.Field
		Timestamp                    respjson.Field
		LowerTimestamp               respjson.Field
		UpperTimestamp               respjson.Field
		IncludeObjectsWithNoValueSet respjson.Field
		OperationType                respjson.Field
		Operator                     respjson.Field
		TimePoint                    respjson.Field
		EndpointBehavior             respjson.Field
		PropertyParser               respjson.Field
		LowerBoundTimePoint          respjson.Field
		UpperBoundTimePoint          respjson.Field
		LowerBoundEndpointBehavior   respjson.Field
		UpperBoundEndpointBehavior   respjson.Field
		raw                          string
	} `json:"-"`
}

func (u PublicEmailEventFilterPruningRefineByUnion) AsNumOccurrences() (v PublicNumOccurrencesRefineBy) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicEmailEventFilterPruningRefineByUnion) AsSetOccurrences() (v PublicSetOccurrencesRefineBy) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicEmailEventFilterPruningRefineByUnion) AsRelativeComparative() (v PublicRelativeComparativeTimestampRefineBy) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicEmailEventFilterPruningRefineByUnion) AsRelativeRanged() (v PublicRelativeRangedTimestampRefineBy) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicEmailEventFilterPruningRefineByUnion) AsAbsoluteComparative() (v PublicAbsoluteComparativeTimestampRefineBy) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicEmailEventFilterPruningRefineByUnion) AsAbsoluteRanged() (v PublicAbsoluteRangedTimestampRefineBy) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicEmailEventFilterPruningRefineByUnion) AsAllHistory() (v PublicAllHistoryRefineBy) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicEmailEventFilterPruningRefineByUnion) AsTimePoint() (v PublicTimePointOperation) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicEmailEventFilterPruningRefineByUnion) AsTimeRanged() (v PublicRangedTimeOperation) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u PublicEmailEventFilterPruningRefineByUnion) RawJSON() string { return u.JSON.raw }

func (r *PublicEmailEventFilterPruningRefineByUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties AppID, EmailID, FilterType, Level, Operator are required.
type PublicEmailEventFilterParam struct {
	// The ID of the application associated with the email event filter.
	AppID string `json:"appId" api:"required"`
	// The ID of the email associated with the event filter.
	EmailID string `json:"emailId" api:"required"`
	// Indicates the type of filter (EMAIL_EVENT).
	//
	// Any of "EMAIL_EVENT".
	FilterType PublicEmailEventFilterFilterType `json:"filterType,omitzero" api:"required"`
	// Specifies the level of the email event, such as EMAIL_API_CAMPAIGN_GROUP.
	Level string `json:"level" api:"required"`
	// Defines the operation to be applied within the filter (BOUNCED, LINK_CLICKED,
	// MARKED_SPAM, OPENED, OPENED_BUT_LINK_NOT_CLICKED, OPENED_BUT_NOT_REPLIED,
	// RECEIVED, RECEIVED_BUT_NOT_OPENED, REPLIED, SENT, SENT_BUT_LINK_NOT_CLICKED,
	// SENT_BUT_NOT_RECEIVED, UNSUBSCRIBED).
	//
	// Any of "BOUNCED", "LINK_CLICKED", "MARKED_SPAM", "OPENED",
	// "OPENED_BUT_LINK_NOT_CLICKED", "OPENED_BUT_NOT_REPLIED", "RECEIVED",
	// "RECEIVED_BUT_NOT_OPENED", "REPLIED", "SENT", "SENT_BUT_LINK_NOT_CLICKED",
	// "SENT_BUT_NOT_RECEIVED", "UNSUBSCRIBED".
	Operator PublicEmailEventFilterOperator `json:"operator,omitzero" api:"required"`
	// The URL that was clicked in the email event.
	ClickURL param.Opt[string] `json:"clickUrl,omitzero"`
	// Specifies the criteria for refining the filter by pruning.
	PruningRefineBy PublicEmailEventFilterPruningRefineByUnionParam `json:"pruningRefineBy,omitzero"`
	paramObj
}

func (r PublicEmailEventFilterParam) MarshalJSON() (data []byte, err error) {
	type shadow PublicEmailEventFilterParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PublicEmailEventFilterParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type PublicEmailEventFilterPruningRefineByUnionParam struct {
	OfNumOccurrences      *PublicNumOccurrencesRefineByParam               `json:",omitzero,inline"`
	OfSetOccurrences      *PublicSetOccurrencesRefineByParam               `json:",omitzero,inline"`
	OfRelativeComparative *PublicRelativeComparativeTimestampRefineByParam `json:",omitzero,inline"`
	OfRelativeRanged      *PublicRelativeRangedTimestampRefineByParam      `json:",omitzero,inline"`
	OfAbsoluteComparative *PublicAbsoluteComparativeTimestampRefineByParam `json:",omitzero,inline"`
	OfAbsoluteRanged      *PublicAbsoluteRangedTimestampRefineByParam      `json:",omitzero,inline"`
	OfAllHistory          *PublicAllHistoryRefineByParam                   `json:",omitzero,inline"`
	OfTimePoint           *PublicTimePointOperationParam                   `json:",omitzero,inline"`
	OfTimeRanged          *PublicRangedTimeOperationParam                  `json:",omitzero,inline"`
	paramUnion
}

func (u PublicEmailEventFilterPruningRefineByUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfNumOccurrences,
		u.OfSetOccurrences,
		u.OfRelativeComparative,
		u.OfRelativeRanged,
		u.OfAbsoluteComparative,
		u.OfAbsoluteRanged,
		u.OfAllHistory,
		u.OfTimePoint,
		u.OfTimeRanged)
}
func (u *PublicEmailEventFilterPruningRefineByUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

type PublicEmailSubscriptionFilter struct {
	AcceptedStatuses []string `json:"acceptedStatuses" api:"required"`
	// Indicates the type of filter (EMAIL_SUBSCRIPTION).
	//
	// Any of "EMAIL_SUBSCRIPTION".
	FilterType      PublicEmailSubscriptionFilterFilterType `json:"filterType" api:"required"`
	SubscriptionIDs []string                                `json:"subscriptionIds" api:"required"`
	// The type of subscription related to the filter (PORTAL, BRAND, SUBSCRIPTION,
	// HARDBOUNCE, SPAMREPORT).
	SubscriptionType string `json:"subscriptionType"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AcceptedStatuses respjson.Field
		FilterType       respjson.Field
		SubscriptionIDs  respjson.Field
		SubscriptionType respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PublicEmailSubscriptionFilter) RawJSON() string { return r.JSON.raw }
func (r *PublicEmailSubscriptionFilter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this PublicEmailSubscriptionFilter to a
// PublicEmailSubscriptionFilterParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// PublicEmailSubscriptionFilterParam.Overrides()
func (r PublicEmailSubscriptionFilter) ToParam() PublicEmailSubscriptionFilterParam {
	return param.Override[PublicEmailSubscriptionFilterParam](json.RawMessage(r.RawJSON()))
}

// Indicates the type of filter (EMAIL_SUBSCRIPTION).
type PublicEmailSubscriptionFilterFilterType string

const (
	PublicEmailSubscriptionFilterFilterTypeEmailSubscription PublicEmailSubscriptionFilterFilterType = "EMAIL_SUBSCRIPTION"
)

// The properties AcceptedStatuses, FilterType, SubscriptionIDs are required.
type PublicEmailSubscriptionFilterParam struct {
	AcceptedStatuses []string `json:"acceptedStatuses,omitzero" api:"required"`
	// Indicates the type of filter (EMAIL_SUBSCRIPTION).
	//
	// Any of "EMAIL_SUBSCRIPTION".
	FilterType      PublicEmailSubscriptionFilterFilterType `json:"filterType,omitzero" api:"required"`
	SubscriptionIDs []string                                `json:"subscriptionIds,omitzero" api:"required"`
	// The type of subscription related to the filter (PORTAL, BRAND, SUBSCRIPTION,
	// HARDBOUNCE, SPAMREPORT).
	SubscriptionType param.Opt[string] `json:"subscriptionType,omitzero"`
	paramObj
}

func (r PublicEmailSubscriptionFilterParam) MarshalJSON() (data []byte, err error) {
	type shadow PublicEmailSubscriptionFilterParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PublicEmailSubscriptionFilterParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PublicEnumerationPropertyOperation struct {
	// Indicates whether objects with no value set for the property should be included
	// in the operation.
	IncludeObjectsWithNoValueSet bool `json:"includeObjectsWithNoValueSet" api:"required"`
	// Specifies the type of operation (ENUMERATION).
	//
	// Any of "ENUMERATION".
	OperationType PublicEnumerationPropertyOperationOperationType `json:"operationType" api:"required"`
	// Defines the operation to be applied in the enumeration property operation
	// (IS_ANY_OF, IS_NONE_OF, IS_EXACTLY, IS_NOT_EXACTLY, CONTAINS_ALL,
	// DOES_NOT_CONTAIN_ALL, HAS_EVER_BEEN_ANY_OF, HAS_NEVER_BEEN_ANY_OF,
	// HAS_EVER_BEEN_EXACTLY, HAS_NEVER_BEEN_EXACTLY, HAS_EVER_CONTAINED_ALL,
	// HAS_NEVER_CONTAINED_ALL).
	Operator string   `json:"operator" api:"required"`
	Values   []string `json:"values" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		IncludeObjectsWithNoValueSet respjson.Field
		OperationType                respjson.Field
		Operator                     respjson.Field
		Values                       respjson.Field
		ExtraFields                  map[string]respjson.Field
		raw                          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PublicEnumerationPropertyOperation) RawJSON() string { return r.JSON.raw }
func (r *PublicEnumerationPropertyOperation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this PublicEnumerationPropertyOperation to a
// PublicEnumerationPropertyOperationParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// PublicEnumerationPropertyOperationParam.Overrides()
func (r PublicEnumerationPropertyOperation) ToParam() PublicEnumerationPropertyOperationParam {
	return param.Override[PublicEnumerationPropertyOperationParam](json.RawMessage(r.RawJSON()))
}

// Specifies the type of operation (ENUMERATION).
type PublicEnumerationPropertyOperationOperationType string

const (
	PublicEnumerationPropertyOperationOperationTypeEnumeration PublicEnumerationPropertyOperationOperationType = "ENUMERATION"
)

// The properties IncludeObjectsWithNoValueSet, OperationType, Operator, Values are
// required.
type PublicEnumerationPropertyOperationParam struct {
	// Indicates whether objects with no value set for the property should be included
	// in the operation.
	IncludeObjectsWithNoValueSet bool `json:"includeObjectsWithNoValueSet" api:"required"`
	// Specifies the type of operation (ENUMERATION).
	//
	// Any of "ENUMERATION".
	OperationType PublicEnumerationPropertyOperationOperationType `json:"operationType,omitzero" api:"required"`
	// Defines the operation to be applied in the enumeration property operation
	// (IS_ANY_OF, IS_NONE_OF, IS_EXACTLY, IS_NOT_EXACTLY, CONTAINS_ALL,
	// DOES_NOT_CONTAIN_ALL, HAS_EVER_BEEN_ANY_OF, HAS_NEVER_BEEN_ANY_OF,
	// HAS_EVER_BEEN_EXACTLY, HAS_NEVER_BEEN_EXACTLY, HAS_EVER_CONTAINED_ALL,
	// HAS_NEVER_CONTAINED_ALL).
	Operator string   `json:"operator" api:"required"`
	Values   []string `json:"values,omitzero" api:"required"`
	paramObj
}

func (r PublicEnumerationPropertyOperationParam) MarshalJSON() (data []byte, err error) {
	type shadow PublicEnumerationPropertyOperationParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PublicEnumerationPropertyOperationParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PublicEventAnalyticsFilter struct {
	// The ID of the event to be used in the filter.
	EventID string `json:"eventId" api:"required"`
	// Indicates the type of filter being applied (EVENT).
	//
	// Any of "EVENT".
	FilterType PublicEventAnalyticsFilterFilterType `json:"filterType" api:"required"`
	// Defines the operation to be applied within the event filter (HAS_EVENT,
	// NOT_HAS_EVENT).
	Operator string `json:"operator" api:"required"`
	// Specifies the criteria for refining the event filter by coalescing.
	CoalescingRefineBy PublicEventAnalyticsFilterCoalescingRefineByUnion `json:"coalescingRefineBy"`
	// Specifies the criteria for refining the event filter by pruning.
	PruningRefineBy PublicEventAnalyticsFilterPruningRefineByUnion `json:"pruningRefineBy"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		EventID            respjson.Field
		FilterType         respjson.Field
		Operator           respjson.Field
		CoalescingRefineBy respjson.Field
		PruningRefineBy    respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PublicEventAnalyticsFilter) RawJSON() string { return r.JSON.raw }
func (r *PublicEventAnalyticsFilter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this PublicEventAnalyticsFilter to a
// PublicEventAnalyticsFilterParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// PublicEventAnalyticsFilterParam.Overrides()
func (r PublicEventAnalyticsFilter) ToParam() PublicEventAnalyticsFilterParam {
	return param.Override[PublicEventAnalyticsFilterParam](json.RawMessage(r.RawJSON()))
}

// Indicates the type of filter being applied (EVENT).
type PublicEventAnalyticsFilterFilterType string

const (
	PublicEventAnalyticsFilterFilterTypeEvent PublicEventAnalyticsFilterFilterType = "EVENT"
)

// PublicEventAnalyticsFilterCoalescingRefineByUnion contains all possible
// properties and values from [PublicNumOccurrencesRefineBy],
// [PublicSetOccurrencesRefineBy], [PublicRelativeComparativeTimestampRefineBy],
// [PublicRelativeRangedTimestampRefineBy],
// [PublicAbsoluteComparativeTimestampRefineBy],
// [PublicAbsoluteRangedTimestampRefineBy], [PublicAllHistoryRefineBy],
// [PublicTimePointOperation], [PublicRangedTimeOperation].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type PublicEventAnalyticsFilterCoalescingRefineByUnion struct {
	Type string `json:"type"`
	// This field is from variant [PublicNumOccurrencesRefineBy].
	MaxOccurrences int64 `json:"maxOccurrences"`
	// This field is from variant [PublicNumOccurrencesRefineBy].
	MinOccurrences int64 `json:"minOccurrences"`
	// This field is from variant [PublicSetOccurrencesRefineBy].
	SetType    string `json:"setType"`
	Comparison string `json:"comparison"`
	// This field is from variant [PublicRelativeComparativeTimestampRefineBy].
	TimeOffset PublicTimeOffset `json:"timeOffset"`
	// This field is from variant [PublicRelativeRangedTimestampRefineBy].
	LowerBoundOffset PublicTimeOffset `json:"lowerBoundOffset"`
	RangeType        string           `json:"rangeType"`
	// This field is from variant [PublicRelativeRangedTimestampRefineBy].
	UpperBoundOffset PublicTimeOffset `json:"upperBoundOffset"`
	// This field is from variant [PublicAbsoluteComparativeTimestampRefineBy].
	Timestamp int64 `json:"timestamp"`
	// This field is from variant [PublicAbsoluteRangedTimestampRefineBy].
	LowerTimestamp int64 `json:"lowerTimestamp"`
	// This field is from variant [PublicAbsoluteRangedTimestampRefineBy].
	UpperTimestamp               int64  `json:"upperTimestamp"`
	IncludeObjectsWithNoValueSet bool   `json:"includeObjectsWithNoValueSet"`
	OperationType                string `json:"operationType"`
	Operator                     string `json:"operator"`
	// This field is from variant [PublicTimePointOperation].
	TimePoint PublicTimePointOperationTimePointUnion `json:"timePoint"`
	// This field is from variant [PublicTimePointOperation].
	EndpointBehavior string `json:"endpointBehavior"`
	PropertyParser   string `json:"propertyParser"`
	// This field is from variant [PublicRangedTimeOperation].
	LowerBoundTimePoint PublicRangedTimeOperationLowerBoundTimePointUnion `json:"lowerBoundTimePoint"`
	// This field is from variant [PublicRangedTimeOperation].
	UpperBoundTimePoint PublicRangedTimeOperationUpperBoundTimePointUnion `json:"upperBoundTimePoint"`
	// This field is from variant [PublicRangedTimeOperation].
	LowerBoundEndpointBehavior string `json:"lowerBoundEndpointBehavior"`
	// This field is from variant [PublicRangedTimeOperation].
	UpperBoundEndpointBehavior string `json:"upperBoundEndpointBehavior"`
	JSON                       struct {
		Type                         respjson.Field
		MaxOccurrences               respjson.Field
		MinOccurrences               respjson.Field
		SetType                      respjson.Field
		Comparison                   respjson.Field
		TimeOffset                   respjson.Field
		LowerBoundOffset             respjson.Field
		RangeType                    respjson.Field
		UpperBoundOffset             respjson.Field
		Timestamp                    respjson.Field
		LowerTimestamp               respjson.Field
		UpperTimestamp               respjson.Field
		IncludeObjectsWithNoValueSet respjson.Field
		OperationType                respjson.Field
		Operator                     respjson.Field
		TimePoint                    respjson.Field
		EndpointBehavior             respjson.Field
		PropertyParser               respjson.Field
		LowerBoundTimePoint          respjson.Field
		UpperBoundTimePoint          respjson.Field
		LowerBoundEndpointBehavior   respjson.Field
		UpperBoundEndpointBehavior   respjson.Field
		raw                          string
	} `json:"-"`
}

func (u PublicEventAnalyticsFilterCoalescingRefineByUnion) AsNumOccurrences() (v PublicNumOccurrencesRefineBy) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicEventAnalyticsFilterCoalescingRefineByUnion) AsSetOccurrences() (v PublicSetOccurrencesRefineBy) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicEventAnalyticsFilterCoalescingRefineByUnion) AsRelativeComparative() (v PublicRelativeComparativeTimestampRefineBy) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicEventAnalyticsFilterCoalescingRefineByUnion) AsRelativeRanged() (v PublicRelativeRangedTimestampRefineBy) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicEventAnalyticsFilterCoalescingRefineByUnion) AsAbsoluteComparative() (v PublicAbsoluteComparativeTimestampRefineBy) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicEventAnalyticsFilterCoalescingRefineByUnion) AsAbsoluteRanged() (v PublicAbsoluteRangedTimestampRefineBy) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicEventAnalyticsFilterCoalescingRefineByUnion) AsAllHistory() (v PublicAllHistoryRefineBy) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicEventAnalyticsFilterCoalescingRefineByUnion) AsTimePoint() (v PublicTimePointOperation) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicEventAnalyticsFilterCoalescingRefineByUnion) AsTimeRanged() (v PublicRangedTimeOperation) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u PublicEventAnalyticsFilterCoalescingRefineByUnion) RawJSON() string { return u.JSON.raw }

func (r *PublicEventAnalyticsFilterCoalescingRefineByUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PublicEventAnalyticsFilterPruningRefineByUnion contains all possible properties
// and values from [PublicNumOccurrencesRefineBy], [PublicSetOccurrencesRefineBy],
// [PublicRelativeComparativeTimestampRefineBy],
// [PublicRelativeRangedTimestampRefineBy],
// [PublicAbsoluteComparativeTimestampRefineBy],
// [PublicAbsoluteRangedTimestampRefineBy], [PublicAllHistoryRefineBy],
// [PublicTimePointOperation], [PublicRangedTimeOperation].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type PublicEventAnalyticsFilterPruningRefineByUnion struct {
	Type string `json:"type"`
	// This field is from variant [PublicNumOccurrencesRefineBy].
	MaxOccurrences int64 `json:"maxOccurrences"`
	// This field is from variant [PublicNumOccurrencesRefineBy].
	MinOccurrences int64 `json:"minOccurrences"`
	// This field is from variant [PublicSetOccurrencesRefineBy].
	SetType    string `json:"setType"`
	Comparison string `json:"comparison"`
	// This field is from variant [PublicRelativeComparativeTimestampRefineBy].
	TimeOffset PublicTimeOffset `json:"timeOffset"`
	// This field is from variant [PublicRelativeRangedTimestampRefineBy].
	LowerBoundOffset PublicTimeOffset `json:"lowerBoundOffset"`
	RangeType        string           `json:"rangeType"`
	// This field is from variant [PublicRelativeRangedTimestampRefineBy].
	UpperBoundOffset PublicTimeOffset `json:"upperBoundOffset"`
	// This field is from variant [PublicAbsoluteComparativeTimestampRefineBy].
	Timestamp int64 `json:"timestamp"`
	// This field is from variant [PublicAbsoluteRangedTimestampRefineBy].
	LowerTimestamp int64 `json:"lowerTimestamp"`
	// This field is from variant [PublicAbsoluteRangedTimestampRefineBy].
	UpperTimestamp               int64  `json:"upperTimestamp"`
	IncludeObjectsWithNoValueSet bool   `json:"includeObjectsWithNoValueSet"`
	OperationType                string `json:"operationType"`
	Operator                     string `json:"operator"`
	// This field is from variant [PublicTimePointOperation].
	TimePoint PublicTimePointOperationTimePointUnion `json:"timePoint"`
	// This field is from variant [PublicTimePointOperation].
	EndpointBehavior string `json:"endpointBehavior"`
	PropertyParser   string `json:"propertyParser"`
	// This field is from variant [PublicRangedTimeOperation].
	LowerBoundTimePoint PublicRangedTimeOperationLowerBoundTimePointUnion `json:"lowerBoundTimePoint"`
	// This field is from variant [PublicRangedTimeOperation].
	UpperBoundTimePoint PublicRangedTimeOperationUpperBoundTimePointUnion `json:"upperBoundTimePoint"`
	// This field is from variant [PublicRangedTimeOperation].
	LowerBoundEndpointBehavior string `json:"lowerBoundEndpointBehavior"`
	// This field is from variant [PublicRangedTimeOperation].
	UpperBoundEndpointBehavior string `json:"upperBoundEndpointBehavior"`
	JSON                       struct {
		Type                         respjson.Field
		MaxOccurrences               respjson.Field
		MinOccurrences               respjson.Field
		SetType                      respjson.Field
		Comparison                   respjson.Field
		TimeOffset                   respjson.Field
		LowerBoundOffset             respjson.Field
		RangeType                    respjson.Field
		UpperBoundOffset             respjson.Field
		Timestamp                    respjson.Field
		LowerTimestamp               respjson.Field
		UpperTimestamp               respjson.Field
		IncludeObjectsWithNoValueSet respjson.Field
		OperationType                respjson.Field
		Operator                     respjson.Field
		TimePoint                    respjson.Field
		EndpointBehavior             respjson.Field
		PropertyParser               respjson.Field
		LowerBoundTimePoint          respjson.Field
		UpperBoundTimePoint          respjson.Field
		LowerBoundEndpointBehavior   respjson.Field
		UpperBoundEndpointBehavior   respjson.Field
		raw                          string
	} `json:"-"`
}

func (u PublicEventAnalyticsFilterPruningRefineByUnion) AsNumOccurrences() (v PublicNumOccurrencesRefineBy) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicEventAnalyticsFilterPruningRefineByUnion) AsSetOccurrences() (v PublicSetOccurrencesRefineBy) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicEventAnalyticsFilterPruningRefineByUnion) AsRelativeComparative() (v PublicRelativeComparativeTimestampRefineBy) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicEventAnalyticsFilterPruningRefineByUnion) AsRelativeRanged() (v PublicRelativeRangedTimestampRefineBy) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicEventAnalyticsFilterPruningRefineByUnion) AsAbsoluteComparative() (v PublicAbsoluteComparativeTimestampRefineBy) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicEventAnalyticsFilterPruningRefineByUnion) AsAbsoluteRanged() (v PublicAbsoluteRangedTimestampRefineBy) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicEventAnalyticsFilterPruningRefineByUnion) AsAllHistory() (v PublicAllHistoryRefineBy) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicEventAnalyticsFilterPruningRefineByUnion) AsTimePoint() (v PublicTimePointOperation) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicEventAnalyticsFilterPruningRefineByUnion) AsTimeRanged() (v PublicRangedTimeOperation) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u PublicEventAnalyticsFilterPruningRefineByUnion) RawJSON() string { return u.JSON.raw }

func (r *PublicEventAnalyticsFilterPruningRefineByUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties EventID, FilterType, Operator are required.
type PublicEventAnalyticsFilterParam struct {
	// The ID of the event to be used in the filter.
	EventID string `json:"eventId" api:"required"`
	// Indicates the type of filter being applied (EVENT).
	//
	// Any of "EVENT".
	FilterType PublicEventAnalyticsFilterFilterType `json:"filterType,omitzero" api:"required"`
	// Defines the operation to be applied within the event filter (HAS_EVENT,
	// NOT_HAS_EVENT).
	Operator string `json:"operator" api:"required"`
	// Specifies the criteria for refining the event filter by coalescing.
	CoalescingRefineBy PublicEventAnalyticsFilterCoalescingRefineByUnionParam `json:"coalescingRefineBy,omitzero"`
	// Specifies the criteria for refining the event filter by pruning.
	PruningRefineBy PublicEventAnalyticsFilterPruningRefineByUnionParam `json:"pruningRefineBy,omitzero"`
	paramObj
}

func (r PublicEventAnalyticsFilterParam) MarshalJSON() (data []byte, err error) {
	type shadow PublicEventAnalyticsFilterParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PublicEventAnalyticsFilterParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type PublicEventAnalyticsFilterCoalescingRefineByUnionParam struct {
	OfNumOccurrences      *PublicNumOccurrencesRefineByParam               `json:",omitzero,inline"`
	OfSetOccurrences      *PublicSetOccurrencesRefineByParam               `json:",omitzero,inline"`
	OfRelativeComparative *PublicRelativeComparativeTimestampRefineByParam `json:",omitzero,inline"`
	OfRelativeRanged      *PublicRelativeRangedTimestampRefineByParam      `json:",omitzero,inline"`
	OfAbsoluteComparative *PublicAbsoluteComparativeTimestampRefineByParam `json:",omitzero,inline"`
	OfAbsoluteRanged      *PublicAbsoluteRangedTimestampRefineByParam      `json:",omitzero,inline"`
	OfAllHistory          *PublicAllHistoryRefineByParam                   `json:",omitzero,inline"`
	OfTimePoint           *PublicTimePointOperationParam                   `json:",omitzero,inline"`
	OfTimeRanged          *PublicRangedTimeOperationParam                  `json:",omitzero,inline"`
	paramUnion
}

func (u PublicEventAnalyticsFilterCoalescingRefineByUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfNumOccurrences,
		u.OfSetOccurrences,
		u.OfRelativeComparative,
		u.OfRelativeRanged,
		u.OfAbsoluteComparative,
		u.OfAbsoluteRanged,
		u.OfAllHistory,
		u.OfTimePoint,
		u.OfTimeRanged)
}
func (u *PublicEventAnalyticsFilterCoalescingRefineByUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type PublicEventAnalyticsFilterPruningRefineByUnionParam struct {
	OfNumOccurrences      *PublicNumOccurrencesRefineByParam               `json:",omitzero,inline"`
	OfSetOccurrences      *PublicSetOccurrencesRefineByParam               `json:",omitzero,inline"`
	OfRelativeComparative *PublicRelativeComparativeTimestampRefineByParam `json:",omitzero,inline"`
	OfRelativeRanged      *PublicRelativeRangedTimestampRefineByParam      `json:",omitzero,inline"`
	OfAbsoluteComparative *PublicAbsoluteComparativeTimestampRefineByParam `json:",omitzero,inline"`
	OfAbsoluteRanged      *PublicAbsoluteRangedTimestampRefineByParam      `json:",omitzero,inline"`
	OfAllHistory          *PublicAllHistoryRefineByParam                   `json:",omitzero,inline"`
	OfTimePoint           *PublicTimePointOperationParam                   `json:",omitzero,inline"`
	OfTimeRanged          *PublicRangedTimeOperationParam                  `json:",omitzero,inline"`
	paramUnion
}

func (u PublicEventAnalyticsFilterPruningRefineByUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfNumOccurrences,
		u.OfSetOccurrences,
		u.OfRelativeComparative,
		u.OfRelativeRanged,
		u.OfAbsoluteComparative,
		u.OfAbsoluteRanged,
		u.OfAllHistory,
		u.OfTimePoint,
		u.OfTimeRanged)
}
func (u *PublicEventAnalyticsFilterPruningRefineByUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

type PublicEventFilterMetadata struct {
	// Defines the operation to be performed on the property
	Operation PublicEventFilterMetadataOperationUnion `json:"operation" api:"required"`
	// Specifies the property on which the operation is to be applied.
	Property string `json:"property" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Operation   respjson.Field
		Property    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PublicEventFilterMetadata) RawJSON() string { return r.JSON.raw }
func (r *PublicEventFilterMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this PublicEventFilterMetadata to a
// PublicEventFilterMetadataParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// PublicEventFilterMetadataParam.Overrides()
func (r PublicEventFilterMetadata) ToParam() PublicEventFilterMetadataParam {
	return param.Override[PublicEventFilterMetadataParam](json.RawMessage(r.RawJSON()))
}

// PublicEventFilterMetadataOperationUnion contains all possible properties and
// values from [PublicBoolPropertyOperation], [PublicNumberPropertyOperation],
// [PublicStringPropertyOperation], [PublicDateTimePropertyOperation],
// [PublicRangedDatePropertyOperation],
// [PublicComparativePropertyUpdatedOperation],
// [PublicComparativeDatePropertyOperation],
// [PublicRollingDateRangePropertyOperation],
// [PublicRollingPropertyUpdatedOperation], [PublicEnumerationPropertyOperation],
// [PublicAllPropertyTypesOperation], [PublicRangedNumberPropertyOperation],
// [PublicMultiStringPropertyOperation], [PublicDatePropertyOperation],
// [PublicCalendarDatePropertyOperation], [PublicTimePointOperation],
// [PublicRangedTimeOperation].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type PublicEventFilterMetadataOperationUnion struct {
	IncludeObjectsWithNoValueSet bool   `json:"includeObjectsWithNoValueSet"`
	OperationType                string `json:"operationType"`
	Operator                     string `json:"operator"`
	// This field is a union of [bool], [float64], [string]
	Value                      PublicEventFilterMetadataOperationUnionValue `json:"value"`
	RequiresTimeZoneConversion bool                                         `json:"requiresTimeZoneConversion"`
	// This field is from variant [PublicDateTimePropertyOperation].
	Timestamp              int64    `json:"timestamp"`
	LowerBound             int64    `json:"lowerBound"`
	UpperBound             int64    `json:"upperBound"`
	ComparisonPropertyName string   `json:"comparisonPropertyName"`
	DefaultComparisonValue string   `json:"defaultComparisonValue"`
	NumberOfDays           int64    `json:"numberOfDays"`
	Values                 []string `json:"values"`
	// This field is from variant [PublicDatePropertyOperation].
	Day int64 `json:"day"`
	// This field is from variant [PublicDatePropertyOperation].
	Month string `json:"month"`
	// This field is from variant [PublicDatePropertyOperation].
	Year int64 `json:"year"`
	// This field is from variant [PublicCalendarDatePropertyOperation].
	TimeUnit string `json:"timeUnit"`
	// This field is from variant [PublicCalendarDatePropertyOperation].
	FiscalYearStart PublicCalendarDatePropertyOperationFiscalYearStart `json:"fiscalYearStart"`
	// This field is from variant [PublicCalendarDatePropertyOperation].
	TimeUnitCount int64 `json:"timeUnitCount"`
	// This field is from variant [PublicCalendarDatePropertyOperation].
	UseFiscalYear bool `json:"useFiscalYear"`
	// This field is from variant [PublicTimePointOperation].
	TimePoint PublicTimePointOperationTimePointUnion `json:"timePoint"`
	Type      string                                 `json:"type"`
	// This field is from variant [PublicTimePointOperation].
	EndpointBehavior string `json:"endpointBehavior"`
	PropertyParser   string `json:"propertyParser"`
	// This field is from variant [PublicRangedTimeOperation].
	LowerBoundTimePoint PublicRangedTimeOperationLowerBoundTimePointUnion `json:"lowerBoundTimePoint"`
	// This field is from variant [PublicRangedTimeOperation].
	UpperBoundTimePoint PublicRangedTimeOperationUpperBoundTimePointUnion `json:"upperBoundTimePoint"`
	// This field is from variant [PublicRangedTimeOperation].
	LowerBoundEndpointBehavior string `json:"lowerBoundEndpointBehavior"`
	// This field is from variant [PublicRangedTimeOperation].
	UpperBoundEndpointBehavior string `json:"upperBoundEndpointBehavior"`
	JSON                       struct {
		IncludeObjectsWithNoValueSet respjson.Field
		OperationType                respjson.Field
		Operator                     respjson.Field
		Value                        respjson.Field
		RequiresTimeZoneConversion   respjson.Field
		Timestamp                    respjson.Field
		LowerBound                   respjson.Field
		UpperBound                   respjson.Field
		ComparisonPropertyName       respjson.Field
		DefaultComparisonValue       respjson.Field
		NumberOfDays                 respjson.Field
		Values                       respjson.Field
		Day                          respjson.Field
		Month                        respjson.Field
		Year                         respjson.Field
		TimeUnit                     respjson.Field
		FiscalYearStart              respjson.Field
		TimeUnitCount                respjson.Field
		UseFiscalYear                respjson.Field
		TimePoint                    respjson.Field
		Type                         respjson.Field
		EndpointBehavior             respjson.Field
		PropertyParser               respjson.Field
		LowerBoundTimePoint          respjson.Field
		UpperBoundTimePoint          respjson.Field
		LowerBoundEndpointBehavior   respjson.Field
		UpperBoundEndpointBehavior   respjson.Field
		raw                          string
	} `json:"-"`
}

func (u PublicEventFilterMetadataOperationUnion) AsBool() (v PublicBoolPropertyOperation) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicEventFilterMetadataOperationUnion) AsNumber() (v PublicNumberPropertyOperation) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicEventFilterMetadataOperationUnion) AsString() (v PublicStringPropertyOperation) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicEventFilterMetadataOperationUnion) AsDatetime() (v PublicDateTimePropertyOperation) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicEventFilterMetadataOperationUnion) AsRangedDate() (v PublicRangedDatePropertyOperation) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicEventFilterMetadataOperationUnion) AsComparativePropertyUpdated() (v PublicComparativePropertyUpdatedOperation) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicEventFilterMetadataOperationUnion) AsComparativeDate() (v PublicComparativeDatePropertyOperation) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicEventFilterMetadataOperationUnion) AsRollingDateRange() (v PublicRollingDateRangePropertyOperation) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicEventFilterMetadataOperationUnion) AsRollingPropertyUpdated() (v PublicRollingPropertyUpdatedOperation) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicEventFilterMetadataOperationUnion) AsEnumeration() (v PublicEnumerationPropertyOperation) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicEventFilterMetadataOperationUnion) AsAllProperty() (v PublicAllPropertyTypesOperation) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicEventFilterMetadataOperationUnion) AsNumberRanged() (v PublicRangedNumberPropertyOperation) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicEventFilterMetadataOperationUnion) AsMultistring() (v PublicMultiStringPropertyOperation) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicEventFilterMetadataOperationUnion) AsDate() (v PublicDatePropertyOperation) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicEventFilterMetadataOperationUnion) AsCalendarDate() (v PublicCalendarDatePropertyOperation) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicEventFilterMetadataOperationUnion) AsTimePoint() (v PublicTimePointOperation) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicEventFilterMetadataOperationUnion) AsTimeRanged() (v PublicRangedTimeOperation) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u PublicEventFilterMetadataOperationUnion) RawJSON() string { return u.JSON.raw }

func (r *PublicEventFilterMetadataOperationUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PublicEventFilterMetadataOperationUnionValue is an implicit subunion of
// [PublicEventFilterMetadataOperationUnion].
// PublicEventFilterMetadataOperationUnionValue provides convenient access to the
// sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [PublicEventFilterMetadataOperationUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfBool OfFloat OfString]
type PublicEventFilterMetadataOperationUnionValue struct {
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	JSON     struct {
		OfBool   respjson.Field
		OfFloat  respjson.Field
		OfString respjson.Field
		raw      string
	} `json:"-"`
}

func (r *PublicEventFilterMetadataOperationUnionValue) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Operation, Property are required.
type PublicEventFilterMetadataParam struct {
	// Defines the operation to be performed on the property
	Operation PublicEventFilterMetadataOperationUnionParam `json:"operation,omitzero" api:"required"`
	// Specifies the property on which the operation is to be applied.
	Property string `json:"property" api:"required"`
	paramObj
}

func (r PublicEventFilterMetadataParam) MarshalJSON() (data []byte, err error) {
	type shadow PublicEventFilterMetadataParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PublicEventFilterMetadataParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type PublicEventFilterMetadataOperationUnionParam struct {
	OfBool                       *PublicBoolPropertyOperationParam               `json:",omitzero,inline"`
	OfNumber                     *PublicNumberPropertyOperationParam             `json:",omitzero,inline"`
	OfString                     *PublicStringPropertyOperationParam             `json:",omitzero,inline"`
	OfDatetime                   *PublicDateTimePropertyOperationParam           `json:",omitzero,inline"`
	OfRangedDate                 *PublicRangedDatePropertyOperationParam         `json:",omitzero,inline"`
	OfComparativePropertyUpdated *PublicComparativePropertyUpdatedOperationParam `json:",omitzero,inline"`
	OfComparativeDate            *PublicComparativeDatePropertyOperationParam    `json:",omitzero,inline"`
	OfRollingDateRange           *PublicRollingDateRangePropertyOperationParam   `json:",omitzero,inline"`
	OfRollingPropertyUpdated     *PublicRollingPropertyUpdatedOperationParam     `json:",omitzero,inline"`
	OfEnumeration                *PublicEnumerationPropertyOperationParam        `json:",omitzero,inline"`
	OfAllProperty                *PublicAllPropertyTypesOperationParam           `json:",omitzero,inline"`
	OfNumberRanged               *PublicRangedNumberPropertyOperationParam       `json:",omitzero,inline"`
	OfMultistring                *PublicMultiStringPropertyOperationParam        `json:",omitzero,inline"`
	OfDate                       *PublicDatePropertyOperationParam               `json:",omitzero,inline"`
	OfCalendarDate               *PublicCalendarDatePropertyOperationParam       `json:",omitzero,inline"`
	OfTimePoint                  *PublicTimePointOperationParam                  `json:",omitzero,inline"`
	OfTimeRanged                 *PublicRangedTimeOperationParam                 `json:",omitzero,inline"`
	paramUnion
}

func (u PublicEventFilterMetadataOperationUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfBool,
		u.OfNumber,
		u.OfString,
		u.OfDatetime,
		u.OfRangedDate,
		u.OfComparativePropertyUpdated,
		u.OfComparativeDate,
		u.OfRollingDateRange,
		u.OfRollingPropertyUpdated,
		u.OfEnumeration,
		u.OfAllProperty,
		u.OfNumberRanged,
		u.OfMultistring,
		u.OfDate,
		u.OfCalendarDate,
		u.OfTimePoint,
		u.OfTimeRanged)
}
func (u *PublicEventFilterMetadataOperationUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

type PublicFiscalQuarterReference struct {
	// The day component of the fiscal quarter reference.
	Day int64 `json:"day" api:"required"`
	// The month component of the fiscal quarter reference.
	Month int64 `json:"month" api:"required"`
	// Indicates the type of reference (FISCAL_QUARTER).
	//
	// Any of "FISCAL_QUARTER".
	ReferenceType PublicFiscalQuarterReferenceReferenceType `json:"referenceType" api:"required"`
	// The hour component of the fiscal quarter reference.
	Hour int64 `json:"hour"`
	// The millisecond component of the fiscal quarter reference.
	Millisecond int64 `json:"millisecond"`
	// The minute component of the fiscal quarter reference.
	Minute int64 `json:"minute"`
	// The second component of the fiscal quarter reference.
	Second int64 `json:"second"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Day           respjson.Field
		Month         respjson.Field
		ReferenceType respjson.Field
		Hour          respjson.Field
		Millisecond   respjson.Field
		Minute        respjson.Field
		Second        respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PublicFiscalQuarterReference) RawJSON() string { return r.JSON.raw }
func (r *PublicFiscalQuarterReference) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this PublicFiscalQuarterReference to a
// PublicFiscalQuarterReferenceParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// PublicFiscalQuarterReferenceParam.Overrides()
func (r PublicFiscalQuarterReference) ToParam() PublicFiscalQuarterReferenceParam {
	return param.Override[PublicFiscalQuarterReferenceParam](json.RawMessage(r.RawJSON()))
}

// Indicates the type of reference (FISCAL_QUARTER).
type PublicFiscalQuarterReferenceReferenceType string

const (
	PublicFiscalQuarterReferenceReferenceTypeFiscalQuarter PublicFiscalQuarterReferenceReferenceType = "FISCAL_QUARTER"
)

// The properties Day, Month, ReferenceType are required.
type PublicFiscalQuarterReferenceParam struct {
	// The day component of the fiscal quarter reference.
	Day int64 `json:"day" api:"required"`
	// The month component of the fiscal quarter reference.
	Month int64 `json:"month" api:"required"`
	// Indicates the type of reference (FISCAL_QUARTER).
	//
	// Any of "FISCAL_QUARTER".
	ReferenceType PublicFiscalQuarterReferenceReferenceType `json:"referenceType,omitzero" api:"required"`
	// The hour component of the fiscal quarter reference.
	Hour param.Opt[int64] `json:"hour,omitzero"`
	// The millisecond component of the fiscal quarter reference.
	Millisecond param.Opt[int64] `json:"millisecond,omitzero"`
	// The minute component of the fiscal quarter reference.
	Minute param.Opt[int64] `json:"minute,omitzero"`
	// The second component of the fiscal quarter reference.
	Second param.Opt[int64] `json:"second,omitzero"`
	paramObj
}

func (r PublicFiscalQuarterReferenceParam) MarshalJSON() (data []byte, err error) {
	type shadow PublicFiscalQuarterReferenceParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PublicFiscalQuarterReferenceParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PublicFiscalYearReference struct {
	// The day component of the fiscal year reference.
	Day int64 `json:"day" api:"required"`
	// The month component of the fiscal year reference.
	Month int64 `json:"month" api:"required"`
	// Indicates the type of reference (FISCAL_YEAR).
	//
	// Any of "FISCAL_YEAR".
	ReferenceType PublicFiscalYearReferenceReferenceType `json:"referenceType" api:"required"`
	// The hour component of the fiscal year reference.
	Hour int64 `json:"hour"`
	// The millisecond component of the fiscal year reference.
	Millisecond int64 `json:"millisecond"`
	// The minute component of the fiscal year reference.
	Minute int64 `json:"minute"`
	// The second component of the fiscal year reference.
	Second int64 `json:"second"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Day           respjson.Field
		Month         respjson.Field
		ReferenceType respjson.Field
		Hour          respjson.Field
		Millisecond   respjson.Field
		Minute        respjson.Field
		Second        respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PublicFiscalYearReference) RawJSON() string { return r.JSON.raw }
func (r *PublicFiscalYearReference) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this PublicFiscalYearReference to a
// PublicFiscalYearReferenceParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// PublicFiscalYearReferenceParam.Overrides()
func (r PublicFiscalYearReference) ToParam() PublicFiscalYearReferenceParam {
	return param.Override[PublicFiscalYearReferenceParam](json.RawMessage(r.RawJSON()))
}

// Indicates the type of reference (FISCAL_YEAR).
type PublicFiscalYearReferenceReferenceType string

const (
	PublicFiscalYearReferenceReferenceTypeFiscalYear PublicFiscalYearReferenceReferenceType = "FISCAL_YEAR"
)

// The properties Day, Month, ReferenceType are required.
type PublicFiscalYearReferenceParam struct {
	// The day component of the fiscal year reference.
	Day int64 `json:"day" api:"required"`
	// The month component of the fiscal year reference.
	Month int64 `json:"month" api:"required"`
	// Indicates the type of reference (FISCAL_YEAR).
	//
	// Any of "FISCAL_YEAR".
	ReferenceType PublicFiscalYearReferenceReferenceType `json:"referenceType,omitzero" api:"required"`
	// The hour component of the fiscal year reference.
	Hour param.Opt[int64] `json:"hour,omitzero"`
	// The millisecond component of the fiscal year reference.
	Millisecond param.Opt[int64] `json:"millisecond,omitzero"`
	// The minute component of the fiscal year reference.
	Minute param.Opt[int64] `json:"minute,omitzero"`
	// The second component of the fiscal year reference.
	Second param.Opt[int64] `json:"second,omitzero"`
	paramObj
}

func (r PublicFiscalYearReferenceParam) MarshalJSON() (data []byte, err error) {
	type shadow PublicFiscalYearReferenceParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PublicFiscalYearReferenceParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PublicFormSubmissionFilter struct {
	// Indicates the type of filter (FORM_SUBMISSION).
	//
	// Any of "FORM_SUBMISSION".
	FilterType PublicFormSubmissionFilterFilterType `json:"filterType" api:"required"`
	// Specifies the operation to be performed (FILLED_OUT, NOT_FILLED_OUT).
	//
	// Any of "FILLED_OUT", "NOT_FILLED_OUT".
	Operator PublicFormSubmissionFilterOperator `json:"operator" api:"required"`
	// Specifies the criteria for refining the filter by coalescing.
	CoalescingRefineBy PublicFormSubmissionFilterCoalescingRefineByUnion `json:"coalescingRefineBy"`
	// The ID of the form used in the filter.
	FormID string `json:"formId"`
	// Specifies the criteria for refining the filter by pruning.
	PruningRefineBy PublicFormSubmissionFilterPruningRefineByUnion `json:"pruningRefineBy"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FilterType         respjson.Field
		Operator           respjson.Field
		CoalescingRefineBy respjson.Field
		FormID             respjson.Field
		PruningRefineBy    respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PublicFormSubmissionFilter) RawJSON() string { return r.JSON.raw }
func (r *PublicFormSubmissionFilter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this PublicFormSubmissionFilter to a
// PublicFormSubmissionFilterParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// PublicFormSubmissionFilterParam.Overrides()
func (r PublicFormSubmissionFilter) ToParam() PublicFormSubmissionFilterParam {
	return param.Override[PublicFormSubmissionFilterParam](json.RawMessage(r.RawJSON()))
}

// Indicates the type of filter (FORM_SUBMISSION).
type PublicFormSubmissionFilterFilterType string

const (
	PublicFormSubmissionFilterFilterTypeFormSubmission PublicFormSubmissionFilterFilterType = "FORM_SUBMISSION"
)

// Specifies the operation to be performed (FILLED_OUT, NOT_FILLED_OUT).
type PublicFormSubmissionFilterOperator string

const (
	PublicFormSubmissionFilterOperatorFilledOut    PublicFormSubmissionFilterOperator = "FILLED_OUT"
	PublicFormSubmissionFilterOperatorNotFilledOut PublicFormSubmissionFilterOperator = "NOT_FILLED_OUT"
)

// PublicFormSubmissionFilterCoalescingRefineByUnion contains all possible
// properties and values from [PublicNumOccurrencesRefineBy],
// [PublicSetOccurrencesRefineBy], [PublicRelativeComparativeTimestampRefineBy],
// [PublicRelativeRangedTimestampRefineBy],
// [PublicAbsoluteComparativeTimestampRefineBy],
// [PublicAbsoluteRangedTimestampRefineBy], [PublicAllHistoryRefineBy],
// [PublicTimePointOperation], [PublicRangedTimeOperation].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type PublicFormSubmissionFilterCoalescingRefineByUnion struct {
	Type string `json:"type"`
	// This field is from variant [PublicNumOccurrencesRefineBy].
	MaxOccurrences int64 `json:"maxOccurrences"`
	// This field is from variant [PublicNumOccurrencesRefineBy].
	MinOccurrences int64 `json:"minOccurrences"`
	// This field is from variant [PublicSetOccurrencesRefineBy].
	SetType    string `json:"setType"`
	Comparison string `json:"comparison"`
	// This field is from variant [PublicRelativeComparativeTimestampRefineBy].
	TimeOffset PublicTimeOffset `json:"timeOffset"`
	// This field is from variant [PublicRelativeRangedTimestampRefineBy].
	LowerBoundOffset PublicTimeOffset `json:"lowerBoundOffset"`
	RangeType        string           `json:"rangeType"`
	// This field is from variant [PublicRelativeRangedTimestampRefineBy].
	UpperBoundOffset PublicTimeOffset `json:"upperBoundOffset"`
	// This field is from variant [PublicAbsoluteComparativeTimestampRefineBy].
	Timestamp int64 `json:"timestamp"`
	// This field is from variant [PublicAbsoluteRangedTimestampRefineBy].
	LowerTimestamp int64 `json:"lowerTimestamp"`
	// This field is from variant [PublicAbsoluteRangedTimestampRefineBy].
	UpperTimestamp               int64  `json:"upperTimestamp"`
	IncludeObjectsWithNoValueSet bool   `json:"includeObjectsWithNoValueSet"`
	OperationType                string `json:"operationType"`
	Operator                     string `json:"operator"`
	// This field is from variant [PublicTimePointOperation].
	TimePoint PublicTimePointOperationTimePointUnion `json:"timePoint"`
	// This field is from variant [PublicTimePointOperation].
	EndpointBehavior string `json:"endpointBehavior"`
	PropertyParser   string `json:"propertyParser"`
	// This field is from variant [PublicRangedTimeOperation].
	LowerBoundTimePoint PublicRangedTimeOperationLowerBoundTimePointUnion `json:"lowerBoundTimePoint"`
	// This field is from variant [PublicRangedTimeOperation].
	UpperBoundTimePoint PublicRangedTimeOperationUpperBoundTimePointUnion `json:"upperBoundTimePoint"`
	// This field is from variant [PublicRangedTimeOperation].
	LowerBoundEndpointBehavior string `json:"lowerBoundEndpointBehavior"`
	// This field is from variant [PublicRangedTimeOperation].
	UpperBoundEndpointBehavior string `json:"upperBoundEndpointBehavior"`
	JSON                       struct {
		Type                         respjson.Field
		MaxOccurrences               respjson.Field
		MinOccurrences               respjson.Field
		SetType                      respjson.Field
		Comparison                   respjson.Field
		TimeOffset                   respjson.Field
		LowerBoundOffset             respjson.Field
		RangeType                    respjson.Field
		UpperBoundOffset             respjson.Field
		Timestamp                    respjson.Field
		LowerTimestamp               respjson.Field
		UpperTimestamp               respjson.Field
		IncludeObjectsWithNoValueSet respjson.Field
		OperationType                respjson.Field
		Operator                     respjson.Field
		TimePoint                    respjson.Field
		EndpointBehavior             respjson.Field
		PropertyParser               respjson.Field
		LowerBoundTimePoint          respjson.Field
		UpperBoundTimePoint          respjson.Field
		LowerBoundEndpointBehavior   respjson.Field
		UpperBoundEndpointBehavior   respjson.Field
		raw                          string
	} `json:"-"`
}

func (u PublicFormSubmissionFilterCoalescingRefineByUnion) AsNumOccurrences() (v PublicNumOccurrencesRefineBy) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicFormSubmissionFilterCoalescingRefineByUnion) AsSetOccurrences() (v PublicSetOccurrencesRefineBy) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicFormSubmissionFilterCoalescingRefineByUnion) AsRelativeComparative() (v PublicRelativeComparativeTimestampRefineBy) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicFormSubmissionFilterCoalescingRefineByUnion) AsRelativeRanged() (v PublicRelativeRangedTimestampRefineBy) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicFormSubmissionFilterCoalescingRefineByUnion) AsAbsoluteComparative() (v PublicAbsoluteComparativeTimestampRefineBy) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicFormSubmissionFilterCoalescingRefineByUnion) AsAbsoluteRanged() (v PublicAbsoluteRangedTimestampRefineBy) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicFormSubmissionFilterCoalescingRefineByUnion) AsAllHistory() (v PublicAllHistoryRefineBy) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicFormSubmissionFilterCoalescingRefineByUnion) AsTimePoint() (v PublicTimePointOperation) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicFormSubmissionFilterCoalescingRefineByUnion) AsTimeRanged() (v PublicRangedTimeOperation) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u PublicFormSubmissionFilterCoalescingRefineByUnion) RawJSON() string { return u.JSON.raw }

func (r *PublicFormSubmissionFilterCoalescingRefineByUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PublicFormSubmissionFilterPruningRefineByUnion contains all possible properties
// and values from [PublicNumOccurrencesRefineBy], [PublicSetOccurrencesRefineBy],
// [PublicRelativeComparativeTimestampRefineBy],
// [PublicRelativeRangedTimestampRefineBy],
// [PublicAbsoluteComparativeTimestampRefineBy],
// [PublicAbsoluteRangedTimestampRefineBy], [PublicAllHistoryRefineBy],
// [PublicTimePointOperation], [PublicRangedTimeOperation].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type PublicFormSubmissionFilterPruningRefineByUnion struct {
	Type string `json:"type"`
	// This field is from variant [PublicNumOccurrencesRefineBy].
	MaxOccurrences int64 `json:"maxOccurrences"`
	// This field is from variant [PublicNumOccurrencesRefineBy].
	MinOccurrences int64 `json:"minOccurrences"`
	// This field is from variant [PublicSetOccurrencesRefineBy].
	SetType    string `json:"setType"`
	Comparison string `json:"comparison"`
	// This field is from variant [PublicRelativeComparativeTimestampRefineBy].
	TimeOffset PublicTimeOffset `json:"timeOffset"`
	// This field is from variant [PublicRelativeRangedTimestampRefineBy].
	LowerBoundOffset PublicTimeOffset `json:"lowerBoundOffset"`
	RangeType        string           `json:"rangeType"`
	// This field is from variant [PublicRelativeRangedTimestampRefineBy].
	UpperBoundOffset PublicTimeOffset `json:"upperBoundOffset"`
	// This field is from variant [PublicAbsoluteComparativeTimestampRefineBy].
	Timestamp int64 `json:"timestamp"`
	// This field is from variant [PublicAbsoluteRangedTimestampRefineBy].
	LowerTimestamp int64 `json:"lowerTimestamp"`
	// This field is from variant [PublicAbsoluteRangedTimestampRefineBy].
	UpperTimestamp               int64  `json:"upperTimestamp"`
	IncludeObjectsWithNoValueSet bool   `json:"includeObjectsWithNoValueSet"`
	OperationType                string `json:"operationType"`
	Operator                     string `json:"operator"`
	// This field is from variant [PublicTimePointOperation].
	TimePoint PublicTimePointOperationTimePointUnion `json:"timePoint"`
	// This field is from variant [PublicTimePointOperation].
	EndpointBehavior string `json:"endpointBehavior"`
	PropertyParser   string `json:"propertyParser"`
	// This field is from variant [PublicRangedTimeOperation].
	LowerBoundTimePoint PublicRangedTimeOperationLowerBoundTimePointUnion `json:"lowerBoundTimePoint"`
	// This field is from variant [PublicRangedTimeOperation].
	UpperBoundTimePoint PublicRangedTimeOperationUpperBoundTimePointUnion `json:"upperBoundTimePoint"`
	// This field is from variant [PublicRangedTimeOperation].
	LowerBoundEndpointBehavior string `json:"lowerBoundEndpointBehavior"`
	// This field is from variant [PublicRangedTimeOperation].
	UpperBoundEndpointBehavior string `json:"upperBoundEndpointBehavior"`
	JSON                       struct {
		Type                         respjson.Field
		MaxOccurrences               respjson.Field
		MinOccurrences               respjson.Field
		SetType                      respjson.Field
		Comparison                   respjson.Field
		TimeOffset                   respjson.Field
		LowerBoundOffset             respjson.Field
		RangeType                    respjson.Field
		UpperBoundOffset             respjson.Field
		Timestamp                    respjson.Field
		LowerTimestamp               respjson.Field
		UpperTimestamp               respjson.Field
		IncludeObjectsWithNoValueSet respjson.Field
		OperationType                respjson.Field
		Operator                     respjson.Field
		TimePoint                    respjson.Field
		EndpointBehavior             respjson.Field
		PropertyParser               respjson.Field
		LowerBoundTimePoint          respjson.Field
		UpperBoundTimePoint          respjson.Field
		LowerBoundEndpointBehavior   respjson.Field
		UpperBoundEndpointBehavior   respjson.Field
		raw                          string
	} `json:"-"`
}

func (u PublicFormSubmissionFilterPruningRefineByUnion) AsNumOccurrences() (v PublicNumOccurrencesRefineBy) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicFormSubmissionFilterPruningRefineByUnion) AsSetOccurrences() (v PublicSetOccurrencesRefineBy) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicFormSubmissionFilterPruningRefineByUnion) AsRelativeComparative() (v PublicRelativeComparativeTimestampRefineBy) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicFormSubmissionFilterPruningRefineByUnion) AsRelativeRanged() (v PublicRelativeRangedTimestampRefineBy) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicFormSubmissionFilterPruningRefineByUnion) AsAbsoluteComparative() (v PublicAbsoluteComparativeTimestampRefineBy) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicFormSubmissionFilterPruningRefineByUnion) AsAbsoluteRanged() (v PublicAbsoluteRangedTimestampRefineBy) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicFormSubmissionFilterPruningRefineByUnion) AsAllHistory() (v PublicAllHistoryRefineBy) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicFormSubmissionFilterPruningRefineByUnion) AsTimePoint() (v PublicTimePointOperation) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicFormSubmissionFilterPruningRefineByUnion) AsTimeRanged() (v PublicRangedTimeOperation) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u PublicFormSubmissionFilterPruningRefineByUnion) RawJSON() string { return u.JSON.raw }

func (r *PublicFormSubmissionFilterPruningRefineByUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties FilterType, Operator are required.
type PublicFormSubmissionFilterParam struct {
	// Indicates the type of filter (FORM_SUBMISSION).
	//
	// Any of "FORM_SUBMISSION".
	FilterType PublicFormSubmissionFilterFilterType `json:"filterType,omitzero" api:"required"`
	// Specifies the operation to be performed (FILLED_OUT, NOT_FILLED_OUT).
	//
	// Any of "FILLED_OUT", "NOT_FILLED_OUT".
	Operator PublicFormSubmissionFilterOperator `json:"operator,omitzero" api:"required"`
	// The ID of the form used in the filter.
	FormID param.Opt[string] `json:"formId,omitzero"`
	// Specifies the criteria for refining the filter by coalescing.
	CoalescingRefineBy PublicFormSubmissionFilterCoalescingRefineByUnionParam `json:"coalescingRefineBy,omitzero"`
	// Specifies the criteria for refining the filter by pruning.
	PruningRefineBy PublicFormSubmissionFilterPruningRefineByUnionParam `json:"pruningRefineBy,omitzero"`
	paramObj
}

func (r PublicFormSubmissionFilterParam) MarshalJSON() (data []byte, err error) {
	type shadow PublicFormSubmissionFilterParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PublicFormSubmissionFilterParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type PublicFormSubmissionFilterCoalescingRefineByUnionParam struct {
	OfNumOccurrences      *PublicNumOccurrencesRefineByParam               `json:",omitzero,inline"`
	OfSetOccurrences      *PublicSetOccurrencesRefineByParam               `json:",omitzero,inline"`
	OfRelativeComparative *PublicRelativeComparativeTimestampRefineByParam `json:",omitzero,inline"`
	OfRelativeRanged      *PublicRelativeRangedTimestampRefineByParam      `json:",omitzero,inline"`
	OfAbsoluteComparative *PublicAbsoluteComparativeTimestampRefineByParam `json:",omitzero,inline"`
	OfAbsoluteRanged      *PublicAbsoluteRangedTimestampRefineByParam      `json:",omitzero,inline"`
	OfAllHistory          *PublicAllHistoryRefineByParam                   `json:",omitzero,inline"`
	OfTimePoint           *PublicTimePointOperationParam                   `json:",omitzero,inline"`
	OfTimeRanged          *PublicRangedTimeOperationParam                  `json:",omitzero,inline"`
	paramUnion
}

func (u PublicFormSubmissionFilterCoalescingRefineByUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfNumOccurrences,
		u.OfSetOccurrences,
		u.OfRelativeComparative,
		u.OfRelativeRanged,
		u.OfAbsoluteComparative,
		u.OfAbsoluteRanged,
		u.OfAllHistory,
		u.OfTimePoint,
		u.OfTimeRanged)
}
func (u *PublicFormSubmissionFilterCoalescingRefineByUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type PublicFormSubmissionFilterPruningRefineByUnionParam struct {
	OfNumOccurrences      *PublicNumOccurrencesRefineByParam               `json:",omitzero,inline"`
	OfSetOccurrences      *PublicSetOccurrencesRefineByParam               `json:",omitzero,inline"`
	OfRelativeComparative *PublicRelativeComparativeTimestampRefineByParam `json:",omitzero,inline"`
	OfRelativeRanged      *PublicRelativeRangedTimestampRefineByParam      `json:",omitzero,inline"`
	OfAbsoluteComparative *PublicAbsoluteComparativeTimestampRefineByParam `json:",omitzero,inline"`
	OfAbsoluteRanged      *PublicAbsoluteRangedTimestampRefineByParam      `json:",omitzero,inline"`
	OfAllHistory          *PublicAllHistoryRefineByParam                   `json:",omitzero,inline"`
	OfTimePoint           *PublicTimePointOperationParam                   `json:",omitzero,inline"`
	OfTimeRanged          *PublicRangedTimeOperationParam                  `json:",omitzero,inline"`
	paramUnion
}

func (u PublicFormSubmissionFilterPruningRefineByUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfNumOccurrences,
		u.OfSetOccurrences,
		u.OfRelativeComparative,
		u.OfRelativeRanged,
		u.OfAbsoluteComparative,
		u.OfAbsoluteRanged,
		u.OfAllHistory,
		u.OfTimePoint,
		u.OfTimeRanged)
}
func (u *PublicFormSubmissionFilterPruningRefineByUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

type PublicFormSubmissionOnPageFilter struct {
	// Indicates the type of filter (FORM_SUBMISSION_ON_PAGE).
	//
	// Any of "FORM_SUBMISSION_ON_PAGE".
	FilterType PublicFormSubmissionOnPageFilterFilterType `json:"filterType" api:"required"`
	// Specifies the operation to be applied (FILLED_OUT, NOT_FILLED_OUT).
	//
	// Any of "FILLED_OUT", "NOT_FILLED_OUT".
	Operator PublicFormSubmissionOnPageFilterOperator `json:"operator" api:"required"`
	// The ID of the page where the form submission occurred.
	PageID string `json:"pageId" api:"required"`
	// Defines the criteria for refining the filter by coalescing.
	CoalescingRefineBy PublicFormSubmissionOnPageFilterCoalescingRefineByUnion `json:"coalescingRefineBy"`
	// The ID of the form associated with the submission filter.
	FormID string `json:"formId"`
	// Specifies the criteria for refining the filter by pruning.
	PruningRefineBy PublicFormSubmissionOnPageFilterPruningRefineByUnion `json:"pruningRefineBy"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FilterType         respjson.Field
		Operator           respjson.Field
		PageID             respjson.Field
		CoalescingRefineBy respjson.Field
		FormID             respjson.Field
		PruningRefineBy    respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PublicFormSubmissionOnPageFilter) RawJSON() string { return r.JSON.raw }
func (r *PublicFormSubmissionOnPageFilter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this PublicFormSubmissionOnPageFilter to a
// PublicFormSubmissionOnPageFilterParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// PublicFormSubmissionOnPageFilterParam.Overrides()
func (r PublicFormSubmissionOnPageFilter) ToParam() PublicFormSubmissionOnPageFilterParam {
	return param.Override[PublicFormSubmissionOnPageFilterParam](json.RawMessage(r.RawJSON()))
}

// Indicates the type of filter (FORM_SUBMISSION_ON_PAGE).
type PublicFormSubmissionOnPageFilterFilterType string

const (
	PublicFormSubmissionOnPageFilterFilterTypeFormSubmissionOnPage PublicFormSubmissionOnPageFilterFilterType = "FORM_SUBMISSION_ON_PAGE"
)

// Specifies the operation to be applied (FILLED_OUT, NOT_FILLED_OUT).
type PublicFormSubmissionOnPageFilterOperator string

const (
	PublicFormSubmissionOnPageFilterOperatorFilledOut    PublicFormSubmissionOnPageFilterOperator = "FILLED_OUT"
	PublicFormSubmissionOnPageFilterOperatorNotFilledOut PublicFormSubmissionOnPageFilterOperator = "NOT_FILLED_OUT"
)

// PublicFormSubmissionOnPageFilterCoalescingRefineByUnion contains all possible
// properties and values from [PublicNumOccurrencesRefineBy],
// [PublicSetOccurrencesRefineBy], [PublicRelativeComparativeTimestampRefineBy],
// [PublicRelativeRangedTimestampRefineBy],
// [PublicAbsoluteComparativeTimestampRefineBy],
// [PublicAbsoluteRangedTimestampRefineBy], [PublicAllHistoryRefineBy],
// [PublicTimePointOperation], [PublicRangedTimeOperation].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type PublicFormSubmissionOnPageFilterCoalescingRefineByUnion struct {
	Type string `json:"type"`
	// This field is from variant [PublicNumOccurrencesRefineBy].
	MaxOccurrences int64 `json:"maxOccurrences"`
	// This field is from variant [PublicNumOccurrencesRefineBy].
	MinOccurrences int64 `json:"minOccurrences"`
	// This field is from variant [PublicSetOccurrencesRefineBy].
	SetType    string `json:"setType"`
	Comparison string `json:"comparison"`
	// This field is from variant [PublicRelativeComparativeTimestampRefineBy].
	TimeOffset PublicTimeOffset `json:"timeOffset"`
	// This field is from variant [PublicRelativeRangedTimestampRefineBy].
	LowerBoundOffset PublicTimeOffset `json:"lowerBoundOffset"`
	RangeType        string           `json:"rangeType"`
	// This field is from variant [PublicRelativeRangedTimestampRefineBy].
	UpperBoundOffset PublicTimeOffset `json:"upperBoundOffset"`
	// This field is from variant [PublicAbsoluteComparativeTimestampRefineBy].
	Timestamp int64 `json:"timestamp"`
	// This field is from variant [PublicAbsoluteRangedTimestampRefineBy].
	LowerTimestamp int64 `json:"lowerTimestamp"`
	// This field is from variant [PublicAbsoluteRangedTimestampRefineBy].
	UpperTimestamp               int64  `json:"upperTimestamp"`
	IncludeObjectsWithNoValueSet bool   `json:"includeObjectsWithNoValueSet"`
	OperationType                string `json:"operationType"`
	Operator                     string `json:"operator"`
	// This field is from variant [PublicTimePointOperation].
	TimePoint PublicTimePointOperationTimePointUnion `json:"timePoint"`
	// This field is from variant [PublicTimePointOperation].
	EndpointBehavior string `json:"endpointBehavior"`
	PropertyParser   string `json:"propertyParser"`
	// This field is from variant [PublicRangedTimeOperation].
	LowerBoundTimePoint PublicRangedTimeOperationLowerBoundTimePointUnion `json:"lowerBoundTimePoint"`
	// This field is from variant [PublicRangedTimeOperation].
	UpperBoundTimePoint PublicRangedTimeOperationUpperBoundTimePointUnion `json:"upperBoundTimePoint"`
	// This field is from variant [PublicRangedTimeOperation].
	LowerBoundEndpointBehavior string `json:"lowerBoundEndpointBehavior"`
	// This field is from variant [PublicRangedTimeOperation].
	UpperBoundEndpointBehavior string `json:"upperBoundEndpointBehavior"`
	JSON                       struct {
		Type                         respjson.Field
		MaxOccurrences               respjson.Field
		MinOccurrences               respjson.Field
		SetType                      respjson.Field
		Comparison                   respjson.Field
		TimeOffset                   respjson.Field
		LowerBoundOffset             respjson.Field
		RangeType                    respjson.Field
		UpperBoundOffset             respjson.Field
		Timestamp                    respjson.Field
		LowerTimestamp               respjson.Field
		UpperTimestamp               respjson.Field
		IncludeObjectsWithNoValueSet respjson.Field
		OperationType                respjson.Field
		Operator                     respjson.Field
		TimePoint                    respjson.Field
		EndpointBehavior             respjson.Field
		PropertyParser               respjson.Field
		LowerBoundTimePoint          respjson.Field
		UpperBoundTimePoint          respjson.Field
		LowerBoundEndpointBehavior   respjson.Field
		UpperBoundEndpointBehavior   respjson.Field
		raw                          string
	} `json:"-"`
}

func (u PublicFormSubmissionOnPageFilterCoalescingRefineByUnion) AsNumOccurrences() (v PublicNumOccurrencesRefineBy) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicFormSubmissionOnPageFilterCoalescingRefineByUnion) AsSetOccurrences() (v PublicSetOccurrencesRefineBy) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicFormSubmissionOnPageFilterCoalescingRefineByUnion) AsRelativeComparative() (v PublicRelativeComparativeTimestampRefineBy) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicFormSubmissionOnPageFilterCoalescingRefineByUnion) AsRelativeRanged() (v PublicRelativeRangedTimestampRefineBy) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicFormSubmissionOnPageFilterCoalescingRefineByUnion) AsAbsoluteComparative() (v PublicAbsoluteComparativeTimestampRefineBy) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicFormSubmissionOnPageFilterCoalescingRefineByUnion) AsAbsoluteRanged() (v PublicAbsoluteRangedTimestampRefineBy) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicFormSubmissionOnPageFilterCoalescingRefineByUnion) AsAllHistory() (v PublicAllHistoryRefineBy) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicFormSubmissionOnPageFilterCoalescingRefineByUnion) AsTimePoint() (v PublicTimePointOperation) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicFormSubmissionOnPageFilterCoalescingRefineByUnion) AsTimeRanged() (v PublicRangedTimeOperation) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u PublicFormSubmissionOnPageFilterCoalescingRefineByUnion) RawJSON() string { return u.JSON.raw }

func (r *PublicFormSubmissionOnPageFilterCoalescingRefineByUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PublicFormSubmissionOnPageFilterPruningRefineByUnion contains all possible
// properties and values from [PublicNumOccurrencesRefineBy],
// [PublicSetOccurrencesRefineBy], [PublicRelativeComparativeTimestampRefineBy],
// [PublicRelativeRangedTimestampRefineBy],
// [PublicAbsoluteComparativeTimestampRefineBy],
// [PublicAbsoluteRangedTimestampRefineBy], [PublicAllHistoryRefineBy],
// [PublicTimePointOperation], [PublicRangedTimeOperation].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type PublicFormSubmissionOnPageFilterPruningRefineByUnion struct {
	Type string `json:"type"`
	// This field is from variant [PublicNumOccurrencesRefineBy].
	MaxOccurrences int64 `json:"maxOccurrences"`
	// This field is from variant [PublicNumOccurrencesRefineBy].
	MinOccurrences int64 `json:"minOccurrences"`
	// This field is from variant [PublicSetOccurrencesRefineBy].
	SetType    string `json:"setType"`
	Comparison string `json:"comparison"`
	// This field is from variant [PublicRelativeComparativeTimestampRefineBy].
	TimeOffset PublicTimeOffset `json:"timeOffset"`
	// This field is from variant [PublicRelativeRangedTimestampRefineBy].
	LowerBoundOffset PublicTimeOffset `json:"lowerBoundOffset"`
	RangeType        string           `json:"rangeType"`
	// This field is from variant [PublicRelativeRangedTimestampRefineBy].
	UpperBoundOffset PublicTimeOffset `json:"upperBoundOffset"`
	// This field is from variant [PublicAbsoluteComparativeTimestampRefineBy].
	Timestamp int64 `json:"timestamp"`
	// This field is from variant [PublicAbsoluteRangedTimestampRefineBy].
	LowerTimestamp int64 `json:"lowerTimestamp"`
	// This field is from variant [PublicAbsoluteRangedTimestampRefineBy].
	UpperTimestamp               int64  `json:"upperTimestamp"`
	IncludeObjectsWithNoValueSet bool   `json:"includeObjectsWithNoValueSet"`
	OperationType                string `json:"operationType"`
	Operator                     string `json:"operator"`
	// This field is from variant [PublicTimePointOperation].
	TimePoint PublicTimePointOperationTimePointUnion `json:"timePoint"`
	// This field is from variant [PublicTimePointOperation].
	EndpointBehavior string `json:"endpointBehavior"`
	PropertyParser   string `json:"propertyParser"`
	// This field is from variant [PublicRangedTimeOperation].
	LowerBoundTimePoint PublicRangedTimeOperationLowerBoundTimePointUnion `json:"lowerBoundTimePoint"`
	// This field is from variant [PublicRangedTimeOperation].
	UpperBoundTimePoint PublicRangedTimeOperationUpperBoundTimePointUnion `json:"upperBoundTimePoint"`
	// This field is from variant [PublicRangedTimeOperation].
	LowerBoundEndpointBehavior string `json:"lowerBoundEndpointBehavior"`
	// This field is from variant [PublicRangedTimeOperation].
	UpperBoundEndpointBehavior string `json:"upperBoundEndpointBehavior"`
	JSON                       struct {
		Type                         respjson.Field
		MaxOccurrences               respjson.Field
		MinOccurrences               respjson.Field
		SetType                      respjson.Field
		Comparison                   respjson.Field
		TimeOffset                   respjson.Field
		LowerBoundOffset             respjson.Field
		RangeType                    respjson.Field
		UpperBoundOffset             respjson.Field
		Timestamp                    respjson.Field
		LowerTimestamp               respjson.Field
		UpperTimestamp               respjson.Field
		IncludeObjectsWithNoValueSet respjson.Field
		OperationType                respjson.Field
		Operator                     respjson.Field
		TimePoint                    respjson.Field
		EndpointBehavior             respjson.Field
		PropertyParser               respjson.Field
		LowerBoundTimePoint          respjson.Field
		UpperBoundTimePoint          respjson.Field
		LowerBoundEndpointBehavior   respjson.Field
		UpperBoundEndpointBehavior   respjson.Field
		raw                          string
	} `json:"-"`
}

func (u PublicFormSubmissionOnPageFilterPruningRefineByUnion) AsNumOccurrences() (v PublicNumOccurrencesRefineBy) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicFormSubmissionOnPageFilterPruningRefineByUnion) AsSetOccurrences() (v PublicSetOccurrencesRefineBy) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicFormSubmissionOnPageFilterPruningRefineByUnion) AsRelativeComparative() (v PublicRelativeComparativeTimestampRefineBy) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicFormSubmissionOnPageFilterPruningRefineByUnion) AsRelativeRanged() (v PublicRelativeRangedTimestampRefineBy) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicFormSubmissionOnPageFilterPruningRefineByUnion) AsAbsoluteComparative() (v PublicAbsoluteComparativeTimestampRefineBy) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicFormSubmissionOnPageFilterPruningRefineByUnion) AsAbsoluteRanged() (v PublicAbsoluteRangedTimestampRefineBy) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicFormSubmissionOnPageFilterPruningRefineByUnion) AsAllHistory() (v PublicAllHistoryRefineBy) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicFormSubmissionOnPageFilterPruningRefineByUnion) AsTimePoint() (v PublicTimePointOperation) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicFormSubmissionOnPageFilterPruningRefineByUnion) AsTimeRanged() (v PublicRangedTimeOperation) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u PublicFormSubmissionOnPageFilterPruningRefineByUnion) RawJSON() string { return u.JSON.raw }

func (r *PublicFormSubmissionOnPageFilterPruningRefineByUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties FilterType, Operator, PageID are required.
type PublicFormSubmissionOnPageFilterParam struct {
	// Indicates the type of filter (FORM_SUBMISSION_ON_PAGE).
	//
	// Any of "FORM_SUBMISSION_ON_PAGE".
	FilterType PublicFormSubmissionOnPageFilterFilterType `json:"filterType,omitzero" api:"required"`
	// Specifies the operation to be applied (FILLED_OUT, NOT_FILLED_OUT).
	//
	// Any of "FILLED_OUT", "NOT_FILLED_OUT".
	Operator PublicFormSubmissionOnPageFilterOperator `json:"operator,omitzero" api:"required"`
	// The ID of the page where the form submission occurred.
	PageID string `json:"pageId" api:"required"`
	// The ID of the form associated with the submission filter.
	FormID param.Opt[string] `json:"formId,omitzero"`
	// Defines the criteria for refining the filter by coalescing.
	CoalescingRefineBy PublicFormSubmissionOnPageFilterCoalescingRefineByUnionParam `json:"coalescingRefineBy,omitzero"`
	// Specifies the criteria for refining the filter by pruning.
	PruningRefineBy PublicFormSubmissionOnPageFilterPruningRefineByUnionParam `json:"pruningRefineBy,omitzero"`
	paramObj
}

func (r PublicFormSubmissionOnPageFilterParam) MarshalJSON() (data []byte, err error) {
	type shadow PublicFormSubmissionOnPageFilterParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PublicFormSubmissionOnPageFilterParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type PublicFormSubmissionOnPageFilterCoalescingRefineByUnionParam struct {
	OfNumOccurrences      *PublicNumOccurrencesRefineByParam               `json:",omitzero,inline"`
	OfSetOccurrences      *PublicSetOccurrencesRefineByParam               `json:",omitzero,inline"`
	OfRelativeComparative *PublicRelativeComparativeTimestampRefineByParam `json:",omitzero,inline"`
	OfRelativeRanged      *PublicRelativeRangedTimestampRefineByParam      `json:",omitzero,inline"`
	OfAbsoluteComparative *PublicAbsoluteComparativeTimestampRefineByParam `json:",omitzero,inline"`
	OfAbsoluteRanged      *PublicAbsoluteRangedTimestampRefineByParam      `json:",omitzero,inline"`
	OfAllHistory          *PublicAllHistoryRefineByParam                   `json:",omitzero,inline"`
	OfTimePoint           *PublicTimePointOperationParam                   `json:",omitzero,inline"`
	OfTimeRanged          *PublicRangedTimeOperationParam                  `json:",omitzero,inline"`
	paramUnion
}

func (u PublicFormSubmissionOnPageFilterCoalescingRefineByUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfNumOccurrences,
		u.OfSetOccurrences,
		u.OfRelativeComparative,
		u.OfRelativeRanged,
		u.OfAbsoluteComparative,
		u.OfAbsoluteRanged,
		u.OfAllHistory,
		u.OfTimePoint,
		u.OfTimeRanged)
}
func (u *PublicFormSubmissionOnPageFilterCoalescingRefineByUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type PublicFormSubmissionOnPageFilterPruningRefineByUnionParam struct {
	OfNumOccurrences      *PublicNumOccurrencesRefineByParam               `json:",omitzero,inline"`
	OfSetOccurrences      *PublicSetOccurrencesRefineByParam               `json:",omitzero,inline"`
	OfRelativeComparative *PublicRelativeComparativeTimestampRefineByParam `json:",omitzero,inline"`
	OfRelativeRanged      *PublicRelativeRangedTimestampRefineByParam      `json:",omitzero,inline"`
	OfAbsoluteComparative *PublicAbsoluteComparativeTimestampRefineByParam `json:",omitzero,inline"`
	OfAbsoluteRanged      *PublicAbsoluteRangedTimestampRefineByParam      `json:",omitzero,inline"`
	OfAllHistory          *PublicAllHistoryRefineByParam                   `json:",omitzero,inline"`
	OfTimePoint           *PublicTimePointOperationParam                   `json:",omitzero,inline"`
	OfTimeRanged          *PublicRangedTimeOperationParam                  `json:",omitzero,inline"`
	paramUnion
}

func (u PublicFormSubmissionOnPageFilterPruningRefineByUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfNumOccurrences,
		u.OfSetOccurrences,
		u.OfRelativeComparative,
		u.OfRelativeRanged,
		u.OfAbsoluteComparative,
		u.OfAbsoluteRanged,
		u.OfAllHistory,
		u.OfTimePoint,
		u.OfTimeRanged)
}
func (u *PublicFormSubmissionOnPageFilterPruningRefineByUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

type PublicInListFilter struct {
	// Indicates the type of filter being applied (IN_LIST).
	//
	// Any of "IN_LIST".
	FilterType PublicInListFilterFilterType `json:"filterType" api:"required"`
	// The ID of the list used in the association filter.
	ListID string `json:"listId" api:"required"`
	// Specifies the operation to be performed by the filter (IN_LIST, NOT_IN_LIST).
	Operator string                     `json:"operator" api:"required"`
	Metadata PublicInListFilterMetadata `json:"metadata"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FilterType  respjson.Field
		ListID      respjson.Field
		Operator    respjson.Field
		Metadata    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PublicInListFilter) RawJSON() string { return r.JSON.raw }
func (r *PublicInListFilter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this PublicInListFilter to a PublicInListFilterParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// PublicInListFilterParam.Overrides()
func (r PublicInListFilter) ToParam() PublicInListFilterParam {
	return param.Override[PublicInListFilterParam](json.RawMessage(r.RawJSON()))
}

// Indicates the type of filter being applied (IN_LIST).
type PublicInListFilterFilterType string

const (
	PublicInListFilterFilterTypeInList PublicInListFilterFilterType = "IN_LIST"
)

// The properties FilterType, ListID, Operator are required.
type PublicInListFilterParam struct {
	// Indicates the type of filter being applied (IN_LIST).
	//
	// Any of "IN_LIST".
	FilterType PublicInListFilterFilterType `json:"filterType,omitzero" api:"required"`
	// The ID of the list used in the association filter.
	ListID string `json:"listId" api:"required"`
	// Specifies the operation to be performed by the filter (IN_LIST, NOT_IN_LIST).
	Operator string                          `json:"operator" api:"required"`
	Metadata PublicInListFilterMetadataParam `json:"metadata,omitzero"`
	paramObj
}

func (r PublicInListFilterParam) MarshalJSON() (data []byte, err error) {
	type shadow PublicInListFilterParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PublicInListFilterParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PublicInListFilterMetadata struct {
	// The identifier for the filter metadata.
	ID string `json:"id" api:"required"`
	// Specifies the type of list for the filter (WORKFLOWS_ENROLLMENT,
	// WORKFLOWS_ACTIVE, WORKFLOWS_GOAL, WORKFLOWS_COMPLETED, IMPORT, DATASET,
	// DATASETS).
	InListType string `json:"inListType" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		InListType  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PublicInListFilterMetadata) RawJSON() string { return r.JSON.raw }
func (r *PublicInListFilterMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this PublicInListFilterMetadata to a
// PublicInListFilterMetadataParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// PublicInListFilterMetadataParam.Overrides()
func (r PublicInListFilterMetadata) ToParam() PublicInListFilterMetadataParam {
	return param.Override[PublicInListFilterMetadataParam](json.RawMessage(r.RawJSON()))
}

// The properties ID, InListType are required.
type PublicInListFilterMetadataParam struct {
	// The identifier for the filter metadata.
	ID string `json:"id" api:"required"`
	// Specifies the type of list for the filter (WORKFLOWS_ENROLLMENT,
	// WORKFLOWS_ACTIVE, WORKFLOWS_GOAL, WORKFLOWS_COMPLETED, IMPORT, DATASET,
	// DATASETS).
	InListType string `json:"inListType" api:"required"`
	paramObj
}

func (r PublicInListFilterMetadataParam) MarshalJSON() (data []byte, err error) {
	type shadow PublicInListFilterMetadataParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PublicInListFilterMetadataParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PublicIndexOffset struct {
	// The number of days to offset.
	Days int64 `json:"days"`
	// The number of hours to offset.
	Hours int64 `json:"hours"`
	// The number of milliseconds to offset.
	Milliseconds int64 `json:"milliseconds"`
	// The number of minutes to offset.
	Minutes int64 `json:"minutes"`
	// The number of months to offset.
	Months int64 `json:"months"`
	// The number of quarters to offset.
	Quarters int64 `json:"quarters"`
	// The number of seconds to offset.
	Seconds int64 `json:"seconds"`
	// The number of weeks to offset.
	Weeks int64 `json:"weeks"`
	// The number of years to offset.
	Years int64 `json:"years"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Days         respjson.Field
		Hours        respjson.Field
		Milliseconds respjson.Field
		Minutes      respjson.Field
		Months       respjson.Field
		Quarters     respjson.Field
		Seconds      respjson.Field
		Weeks        respjson.Field
		Years        respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PublicIndexOffset) RawJSON() string { return r.JSON.raw }
func (r *PublicIndexOffset) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this PublicIndexOffset to a PublicIndexOffsetParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// PublicIndexOffsetParam.Overrides()
func (r PublicIndexOffset) ToParam() PublicIndexOffsetParam {
	return param.Override[PublicIndexOffsetParam](json.RawMessage(r.RawJSON()))
}

type PublicIndexOffsetParam struct {
	// The number of days to offset.
	Days param.Opt[int64] `json:"days,omitzero"`
	// The number of hours to offset.
	Hours param.Opt[int64] `json:"hours,omitzero"`
	// The number of milliseconds to offset.
	Milliseconds param.Opt[int64] `json:"milliseconds,omitzero"`
	// The number of minutes to offset.
	Minutes param.Opt[int64] `json:"minutes,omitzero"`
	// The number of months to offset.
	Months param.Opt[int64] `json:"months,omitzero"`
	// The number of quarters to offset.
	Quarters param.Opt[int64] `json:"quarters,omitzero"`
	// The number of seconds to offset.
	Seconds param.Opt[int64] `json:"seconds,omitzero"`
	// The number of weeks to offset.
	Weeks param.Opt[int64] `json:"weeks,omitzero"`
	// The number of years to offset.
	Years param.Opt[int64] `json:"years,omitzero"`
	paramObj
}

func (r PublicIndexOffsetParam) MarshalJSON() (data []byte, err error) {
	type shadow PublicIndexOffsetParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PublicIndexOffsetParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PublicIndexedTimePoint struct {
	// Specifies the reference point in time for the indexed time point.
	IndexReference PublicIndexedTimePointIndexReferenceUnion `json:"indexReference" api:"required"`
	// Defines the type of time (INDEXED).
	//
	// Any of "INDEXED".
	TimeType PublicIndexedTimePointTimeType `json:"timeType" api:"required"`
	// Indicates the identifier for the time zone associated with the indexed time
	// point.
	ZoneID string            `json:"zoneId" api:"required"`
	Offset PublicIndexOffset `json:"offset"`
	// Specifies the source of the time zone information for the indexed time point
	// (CUSTOM, USER, PORTAL).
	TimezoneSource string `json:"timezoneSource"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		IndexReference respjson.Field
		TimeType       respjson.Field
		ZoneID         respjson.Field
		Offset         respjson.Field
		TimezoneSource respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PublicIndexedTimePoint) RawJSON() string { return r.JSON.raw }
func (r *PublicIndexedTimePoint) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this PublicIndexedTimePoint to a PublicIndexedTimePointParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// PublicIndexedTimePointParam.Overrides()
func (r PublicIndexedTimePoint) ToParam() PublicIndexedTimePointParam {
	return param.Override[PublicIndexedTimePointParam](json.RawMessage(r.RawJSON()))
}

// PublicIndexedTimePointIndexReferenceUnion contains all possible properties and
// values from [PublicNowReference], [PublicTodayReference], [PublicWeekReference],
// [PublicFiscalQuarterReference], [PublicFiscalYearReference],
// [PublicYearReference], [PublicQuarterReference], [PublicMonthReference].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type PublicIndexedTimePointIndexReferenceUnion struct {
	ReferenceType string `json:"referenceType"`
	Hour          int64  `json:"hour"`
	Millisecond   int64  `json:"millisecond"`
	Minute        int64  `json:"minute"`
	Second        int64  `json:"second"`
	// This field is from variant [PublicWeekReference].
	DayOfWeek PublicWeekReferenceDayOfWeek `json:"dayOfWeek"`
	Day       int64                        `json:"day"`
	Month     int64                        `json:"month"`
	JSON      struct {
		ReferenceType respjson.Field
		Hour          respjson.Field
		Millisecond   respjson.Field
		Minute        respjson.Field
		Second        respjson.Field
		DayOfWeek     respjson.Field
		Day           respjson.Field
		Month         respjson.Field
		raw           string
	} `json:"-"`
}

func (u PublicIndexedTimePointIndexReferenceUnion) AsNow() (v PublicNowReference) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicIndexedTimePointIndexReferenceUnion) AsToday() (v PublicTodayReference) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicIndexedTimePointIndexReferenceUnion) AsWeek() (v PublicWeekReference) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicIndexedTimePointIndexReferenceUnion) AsFiscalQuarter() (v PublicFiscalQuarterReference) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicIndexedTimePointIndexReferenceUnion) AsFiscalYear() (v PublicFiscalYearReference) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicIndexedTimePointIndexReferenceUnion) AsYear() (v PublicYearReference) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicIndexedTimePointIndexReferenceUnion) AsQuarter() (v PublicQuarterReference) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicIndexedTimePointIndexReferenceUnion) AsMonth() (v PublicMonthReference) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u PublicIndexedTimePointIndexReferenceUnion) RawJSON() string { return u.JSON.raw }

func (r *PublicIndexedTimePointIndexReferenceUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Defines the type of time (INDEXED).
type PublicIndexedTimePointTimeType string

const (
	PublicIndexedTimePointTimeTypeIndexed PublicIndexedTimePointTimeType = "INDEXED"
)

// The properties IndexReference, TimeType, ZoneID are required.
type PublicIndexedTimePointParam struct {
	// Specifies the reference point in time for the indexed time point.
	IndexReference PublicIndexedTimePointIndexReferenceUnionParam `json:"indexReference,omitzero" api:"required"`
	// Defines the type of time (INDEXED).
	//
	// Any of "INDEXED".
	TimeType PublicIndexedTimePointTimeType `json:"timeType,omitzero" api:"required"`
	// Indicates the identifier for the time zone associated with the indexed time
	// point.
	ZoneID string `json:"zoneId" api:"required"`
	// Specifies the source of the time zone information for the indexed time point
	// (CUSTOM, USER, PORTAL).
	TimezoneSource param.Opt[string]      `json:"timezoneSource,omitzero"`
	Offset         PublicIndexOffsetParam `json:"offset,omitzero"`
	paramObj
}

func (r PublicIndexedTimePointParam) MarshalJSON() (data []byte, err error) {
	type shadow PublicIndexedTimePointParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PublicIndexedTimePointParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type PublicIndexedTimePointIndexReferenceUnionParam struct {
	OfNow           *PublicNowReferenceParam           `json:",omitzero,inline"`
	OfToday         *PublicTodayReferenceParam         `json:",omitzero,inline"`
	OfWeek          *PublicWeekReferenceParam          `json:",omitzero,inline"`
	OfFiscalQuarter *PublicFiscalQuarterReferenceParam `json:",omitzero,inline"`
	OfFiscalYear    *PublicFiscalYearReferenceParam    `json:",omitzero,inline"`
	OfYear          *PublicYearReferenceParam          `json:",omitzero,inline"`
	OfQuarter       *PublicQuarterReferenceParam       `json:",omitzero,inline"`
	OfMonth         *PublicMonthReferenceParam         `json:",omitzero,inline"`
	paramUnion
}

func (u PublicIndexedTimePointIndexReferenceUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfNow,
		u.OfToday,
		u.OfWeek,
		u.OfFiscalQuarter,
		u.OfFiscalYear,
		u.OfYear,
		u.OfQuarter,
		u.OfMonth)
}
func (u *PublicIndexedTimePointIndexReferenceUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

type PublicIntegrationEventFilter struct {
	// The ID representing the type of event for the integration event filter.
	EventTypeID int64                       `json:"eventTypeId" api:"required"`
	FilterLines []PublicEventFilterMetadata `json:"filterLines" api:"required"`
	// Indicates the type of filter (INTEGRATION_EVENT).
	//
	// Any of "INTEGRATION_EVENT".
	FilterType PublicIntegrationEventFilterFilterType `json:"filterType" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		EventTypeID respjson.Field
		FilterLines respjson.Field
		FilterType  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PublicIntegrationEventFilter) RawJSON() string { return r.JSON.raw }
func (r *PublicIntegrationEventFilter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this PublicIntegrationEventFilter to a
// PublicIntegrationEventFilterParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// PublicIntegrationEventFilterParam.Overrides()
func (r PublicIntegrationEventFilter) ToParam() PublicIntegrationEventFilterParam {
	return param.Override[PublicIntegrationEventFilterParam](json.RawMessage(r.RawJSON()))
}

// Indicates the type of filter (INTEGRATION_EVENT).
type PublicIntegrationEventFilterFilterType string

const (
	PublicIntegrationEventFilterFilterTypeIntegrationEvent PublicIntegrationEventFilterFilterType = "INTEGRATION_EVENT"
)

// The properties EventTypeID, FilterLines, FilterType are required.
type PublicIntegrationEventFilterParam struct {
	// The ID representing the type of event for the integration event filter.
	EventTypeID int64                            `json:"eventTypeId" api:"required"`
	FilterLines []PublicEventFilterMetadataParam `json:"filterLines,omitzero" api:"required"`
	// Indicates the type of filter (INTEGRATION_EVENT).
	//
	// Any of "INTEGRATION_EVENT".
	FilterType PublicIntegrationEventFilterFilterType `json:"filterType,omitzero" api:"required"`
	paramObj
}

func (r PublicIntegrationEventFilterParam) MarshalJSON() (data []byte, err error) {
	type shadow PublicIntegrationEventFilterParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PublicIntegrationEventFilterParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PublicListConversionDate struct {
	// Specifies the type of conversion (CONVERSION_DATE).
	//
	// Any of "CONVERSION_DATE".
	ConversionType PublicListConversionDateConversionType `json:"conversionType" api:"required"`
	// The day component of the conversion date.
	Day int64 `json:"day" api:"required"`
	// The month component of the conversion date.
	Month int64 `json:"month" api:"required"`
	// The year component of the conversion date.
	Year int64 `json:"year" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ConversionType respjson.Field
		Day            respjson.Field
		Month          respjson.Field
		Year           respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PublicListConversionDate) RawJSON() string { return r.JSON.raw }
func (r *PublicListConversionDate) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this PublicListConversionDate to a
// PublicListConversionDateParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// PublicListConversionDateParam.Overrides()
func (r PublicListConversionDate) ToParam() PublicListConversionDateParam {
	return param.Override[PublicListConversionDateParam](json.RawMessage(r.RawJSON()))
}

// Specifies the type of conversion (CONVERSION_DATE).
type PublicListConversionDateConversionType string

const (
	PublicListConversionDateConversionTypeConversionDate PublicListConversionDateConversionType = "CONVERSION_DATE"
)

// The properties ConversionType, Day, Month, Year are required.
type PublicListConversionDateParam struct {
	// Specifies the type of conversion (CONVERSION_DATE).
	//
	// Any of "CONVERSION_DATE".
	ConversionType PublicListConversionDateConversionType `json:"conversionType,omitzero" api:"required"`
	// The day component of the conversion date.
	Day int64 `json:"day" api:"required"`
	// The month component of the conversion date.
	Month int64 `json:"month" api:"required"`
	// The year component of the conversion date.
	Year int64 `json:"year" api:"required"`
	paramObj
}

func (r PublicListConversionDateParam) MarshalJSON() (data []byte, err error) {
	type shadow PublicListConversionDateParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PublicListConversionDateParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PublicListConversionInactivity struct {
	// Specifies the type of conversion (INACTIVITY).
	//
	// Any of "INACTIVITY".
	ConversionType PublicListConversionInactivityConversionType `json:"conversionType" api:"required"`
	// Value used to paginate through lists. The `offset` provided in the response can
	// be used in the next request to fetch the next page of results. Defaults to `0`
	// if no offset is provided.
	Offset int64 `json:"offset" api:"required"`
	// The unit of time for the inactivity period, such as (DAY, MONTH, WEEK).
	//
	// Any of "DAY", "MONTH", "WEEK".
	TimeUnit PublicListConversionInactivityTimeUnit `json:"timeUnit" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ConversionType respjson.Field
		Offset         respjson.Field
		TimeUnit       respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PublicListConversionInactivity) RawJSON() string { return r.JSON.raw }
func (r *PublicListConversionInactivity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this PublicListConversionInactivity to a
// PublicListConversionInactivityParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// PublicListConversionInactivityParam.Overrides()
func (r PublicListConversionInactivity) ToParam() PublicListConversionInactivityParam {
	return param.Override[PublicListConversionInactivityParam](json.RawMessage(r.RawJSON()))
}

// Specifies the type of conversion (INACTIVITY).
type PublicListConversionInactivityConversionType string

const (
	PublicListConversionInactivityConversionTypeInactivity PublicListConversionInactivityConversionType = "INACTIVITY"
)

// The unit of time for the inactivity period, such as (DAY, MONTH, WEEK).
type PublicListConversionInactivityTimeUnit string

const (
	PublicListConversionInactivityTimeUnitDay   PublicListConversionInactivityTimeUnit = "DAY"
	PublicListConversionInactivityTimeUnitMonth PublicListConversionInactivityTimeUnit = "MONTH"
	PublicListConversionInactivityTimeUnitWeek  PublicListConversionInactivityTimeUnit = "WEEK"
)

// The properties ConversionType, Offset, TimeUnit are required.
type PublicListConversionInactivityParam struct {
	// Specifies the type of conversion (INACTIVITY).
	//
	// Any of "INACTIVITY".
	ConversionType PublicListConversionInactivityConversionType `json:"conversionType,omitzero" api:"required"`
	// Value used to paginate through lists. The `offset` provided in the response can
	// be used in the next request to fetch the next page of results. Defaults to `0`
	// if no offset is provided.
	Offset int64 `json:"offset" api:"required"`
	// The unit of time for the inactivity period, such as (DAY, MONTH, WEEK).
	//
	// Any of "DAY", "MONTH", "WEEK".
	TimeUnit PublicListConversionInactivityTimeUnit `json:"timeUnit,omitzero" api:"required"`
	paramObj
}

func (r PublicListConversionInactivityParam) MarshalJSON() (data []byte, err error) {
	type shadow PublicListConversionInactivityParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PublicListConversionInactivityParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PublicListConversionResponse struct {
	// The unique identifier of the list for which the conversion details are provided.
	ListID string `json:"listId" api:"required"`
	// The date and time when the list was converted.
	ConvertedAt time.Time `json:"convertedAt" format:"date-time"`
	// The scheduled time for the list conversion, which can be based on a specific
	// date or inactivity period.
	RequestedConversionTime PublicListConversionResponseRequestedConversionTimeUnion `json:"requestedConversionTime"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ListID                  respjson.Field
		ConvertedAt             respjson.Field
		RequestedConversionTime respjson.Field
		ExtraFields             map[string]respjson.Field
		raw                     string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PublicListConversionResponse) RawJSON() string { return r.JSON.raw }
func (r *PublicListConversionResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PublicListConversionResponseRequestedConversionTimeUnion contains all possible
// properties and values from [PublicListConversionDate],
// [PublicListConversionInactivity].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type PublicListConversionResponseRequestedConversionTimeUnion struct {
	ConversionType string `json:"conversionType"`
	// This field is from variant [PublicListConversionDate].
	Day int64 `json:"day"`
	// This field is from variant [PublicListConversionDate].
	Month int64 `json:"month"`
	// This field is from variant [PublicListConversionDate].
	Year int64 `json:"year"`
	// This field is from variant [PublicListConversionInactivity].
	Offset int64 `json:"offset"`
	// This field is from variant [PublicListConversionInactivity].
	TimeUnit PublicListConversionInactivityTimeUnit `json:"timeUnit"`
	JSON     struct {
		ConversionType respjson.Field
		Day            respjson.Field
		Month          respjson.Field
		Year           respjson.Field
		Offset         respjson.Field
		TimeUnit       respjson.Field
		raw            string
	} `json:"-"`
}

func (u PublicListConversionResponseRequestedConversionTimeUnion) AsConversionDate() (v PublicListConversionDate) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicListConversionResponseRequestedConversionTimeUnion) AsInactivity() (v PublicListConversionInactivity) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u PublicListConversionResponseRequestedConversionTimeUnion) RawJSON() string { return u.JSON.raw }

func (r *PublicListConversionResponseRequestedConversionTimeUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func PublicListConversionTimeParamOfInactivity(conversionType PublicListConversionInactivityConversionType, offset int64, timeUnit PublicListConversionInactivityTimeUnit) PublicListConversionTimeUnionParam {
	var variant PublicListConversionInactivityParam
	variant.ConversionType = conversionType
	variant.Offset = offset
	variant.TimeUnit = timeUnit
	return PublicListConversionTimeUnionParam{OfInactivity: &variant}
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type PublicListConversionTimeUnionParam struct {
	OfConversionDate *PublicListConversionDateParam       `json:",omitzero,inline"`
	OfInactivity     *PublicListConversionInactivityParam `json:",omitzero,inline"`
	paramUnion
}

func (u PublicListConversionTimeUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfConversionDate, u.OfInactivity)
}
func (u *PublicListConversionTimeUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

type PublicListFolder struct {
	// The Id of the folder.
	ID string `json:"id" api:"required"`
	// An array of list Id's contained in this folder.
	ChildLists []int64            `json:"childLists" api:"required"`
	ChildNodes []PublicListFolder `json:"childNodes" api:"required"`
	// The Id of the folder this folder is in, the root folder is represented as 0.
	ParentFolderID string `json:"parentFolderId" api:"required"`
	// The time the folder was created at.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// The name of the folder.
	Name string `json:"name"`
	// The time the folder was last updated at.
	UpdatedAt time.Time `json:"updatedAt" format:"date-time"`
	// The time that the contents of the folder was last updated at.
	UpdatedContentsAt time.Time `json:"updatedContentsAt" format:"date-time"`
	// The user Id of the owner of the folder.
	UserID int64 `json:"userId"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                respjson.Field
		ChildLists        respjson.Field
		ChildNodes        respjson.Field
		ParentFolderID    respjson.Field
		CreatedAt         respjson.Field
		Name              respjson.Field
		UpdatedAt         respjson.Field
		UpdatedContentsAt respjson.Field
		UserID            respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PublicListFolder) RawJSON() string { return r.JSON.raw }
func (r *PublicListFolder) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PublicListPermissions struct {
	TeamsWithEditAccess []int64 `json:"teamsWithEditAccess" api:"required"`
	UsersWithEditAccess []int64 `json:"usersWithEditAccess" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		TeamsWithEditAccess respjson.Field
		UsersWithEditAccess respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PublicListPermissions) RawJSON() string { return r.JSON.raw }
func (r *PublicListPermissions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this PublicListPermissions to a PublicListPermissionsParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// PublicListPermissionsParam.Overrides()
func (r PublicListPermissions) ToParam() PublicListPermissionsParam {
	return param.Override[PublicListPermissionsParam](json.RawMessage(r.RawJSON()))
}

// The properties TeamsWithEditAccess, UsersWithEditAccess are required.
type PublicListPermissionsParam struct {
	TeamsWithEditAccess []int64 `json:"teamsWithEditAccess,omitzero" api:"required"`
	UsersWithEditAccess []int64 `json:"usersWithEditAccess,omitzero" api:"required"`
	paramObj
}

func (r PublicListPermissionsParam) MarshalJSON() (data []byte, err error) {
	type shadow PublicListPermissionsParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PublicListPermissionsParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PublicMembershipSettings struct {
	// Indicates whether unassigned memberships should be included.
	IncludeUnassigned bool `json:"includeUnassigned"`
	// The ID of the team associated with the membership.
	MembershipTeamID int64 `json:"membershipTeamId"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		IncludeUnassigned respjson.Field
		MembershipTeamID  respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PublicMembershipSettings) RawJSON() string { return r.JSON.raw }
func (r *PublicMembershipSettings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this PublicMembershipSettings to a
// PublicMembershipSettingsParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// PublicMembershipSettingsParam.Overrides()
func (r PublicMembershipSettings) ToParam() PublicMembershipSettingsParam {
	return param.Override[PublicMembershipSettingsParam](json.RawMessage(r.RawJSON()))
}

type PublicMembershipSettingsParam struct {
	// Indicates whether unassigned memberships should be included.
	IncludeUnassigned param.Opt[bool] `json:"includeUnassigned,omitzero"`
	// The ID of the team associated with the membership.
	MembershipTeamID param.Opt[int64] `json:"membershipTeamId,omitzero"`
	paramObj
}

func (r PublicMembershipSettingsParam) MarshalJSON() (data []byte, err error) {
	type shadow PublicMembershipSettingsParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PublicMembershipSettingsParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PublicMigrationMapping struct {
	// The legacy list id for the list
	LegacyListID string `json:"legacyListId" api:"required"`
	// The V3 list id for the list
	ListID string `json:"listId" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		LegacyListID respjson.Field
		ListID       respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PublicMigrationMapping) RawJSON() string { return r.JSON.raw }
func (r *PublicMigrationMapping) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PublicMonthReference struct {
	// The day component of the month reference.
	Day int64 `json:"day" api:"required"`
	// Indicates the type of reference, (MONTH).
	//
	// Any of "MONTH".
	ReferenceType PublicMonthReferenceReferenceType `json:"referenceType" api:"required"`
	// The hour component of the month reference.
	Hour int64 `json:"hour"`
	// The millisecond component of the month reference.
	Millisecond int64 `json:"millisecond"`
	// The minute component of the month reference.
	Minute int64 `json:"minute"`
	// The second component of the month reference.
	Second int64 `json:"second"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Day           respjson.Field
		ReferenceType respjson.Field
		Hour          respjson.Field
		Millisecond   respjson.Field
		Minute        respjson.Field
		Second        respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PublicMonthReference) RawJSON() string { return r.JSON.raw }
func (r *PublicMonthReference) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this PublicMonthReference to a PublicMonthReferenceParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// PublicMonthReferenceParam.Overrides()
func (r PublicMonthReference) ToParam() PublicMonthReferenceParam {
	return param.Override[PublicMonthReferenceParam](json.RawMessage(r.RawJSON()))
}

// Indicates the type of reference, (MONTH).
type PublicMonthReferenceReferenceType string

const (
	PublicMonthReferenceReferenceTypeMonth PublicMonthReferenceReferenceType = "MONTH"
)

// The properties Day, ReferenceType are required.
type PublicMonthReferenceParam struct {
	// The day component of the month reference.
	Day int64 `json:"day" api:"required"`
	// Indicates the type of reference, (MONTH).
	//
	// Any of "MONTH".
	ReferenceType PublicMonthReferenceReferenceType `json:"referenceType,omitzero" api:"required"`
	// The hour component of the month reference.
	Hour param.Opt[int64] `json:"hour,omitzero"`
	// The millisecond component of the month reference.
	Millisecond param.Opt[int64] `json:"millisecond,omitzero"`
	// The minute component of the month reference.
	Minute param.Opt[int64] `json:"minute,omitzero"`
	// The second component of the month reference.
	Second param.Opt[int64] `json:"second,omitzero"`
	paramObj
}

func (r PublicMonthReferenceParam) MarshalJSON() (data []byte, err error) {
	type shadow PublicMonthReferenceParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PublicMonthReferenceParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PublicMultiStringPropertyOperation struct {
	// Indicates whether objects with no value set for the property should be included
	// in the operation.
	IncludeObjectsWithNoValueSet bool `json:"includeObjectsWithNoValueSet" api:"required"`
	// Specifies the type of operation (MULTISTRING).
	//
	// Any of "MULTISTRING".
	OperationType PublicMultiStringPropertyOperationOperationType `json:"operationType" api:"required"`
	// Defines the operation to be applied in the multi-string property operation
	// (IS_EQUAL_TO, IS_NOT_EQUAL_TO, CONTAINS, CONTAINS_EXACTLY, DOES_NOT_CONTAIN,
	// DOES_NOT_CONTAIN_EXACTLY, STARTS_WITH, ENDS_WITH).
	Operator string   `json:"operator" api:"required"`
	Values   []string `json:"values" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		IncludeObjectsWithNoValueSet respjson.Field
		OperationType                respjson.Field
		Operator                     respjson.Field
		Values                       respjson.Field
		ExtraFields                  map[string]respjson.Field
		raw                          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PublicMultiStringPropertyOperation) RawJSON() string { return r.JSON.raw }
func (r *PublicMultiStringPropertyOperation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this PublicMultiStringPropertyOperation to a
// PublicMultiStringPropertyOperationParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// PublicMultiStringPropertyOperationParam.Overrides()
func (r PublicMultiStringPropertyOperation) ToParam() PublicMultiStringPropertyOperationParam {
	return param.Override[PublicMultiStringPropertyOperationParam](json.RawMessage(r.RawJSON()))
}

// Specifies the type of operation (MULTISTRING).
type PublicMultiStringPropertyOperationOperationType string

const (
	PublicMultiStringPropertyOperationOperationTypeMultistring PublicMultiStringPropertyOperationOperationType = "MULTISTRING"
)

// The properties IncludeObjectsWithNoValueSet, OperationType, Operator, Values are
// required.
type PublicMultiStringPropertyOperationParam struct {
	// Indicates whether objects with no value set for the property should be included
	// in the operation.
	IncludeObjectsWithNoValueSet bool `json:"includeObjectsWithNoValueSet" api:"required"`
	// Specifies the type of operation (MULTISTRING).
	//
	// Any of "MULTISTRING".
	OperationType PublicMultiStringPropertyOperationOperationType `json:"operationType,omitzero" api:"required"`
	// Defines the operation to be applied in the multi-string property operation
	// (IS_EQUAL_TO, IS_NOT_EQUAL_TO, CONTAINS, CONTAINS_EXACTLY, DOES_NOT_CONTAIN,
	// DOES_NOT_CONTAIN_EXACTLY, STARTS_WITH, ENDS_WITH).
	Operator string   `json:"operator" api:"required"`
	Values   []string `json:"values,omitzero" api:"required"`
	paramObj
}

func (r PublicMultiStringPropertyOperationParam) MarshalJSON() (data []byte, err error) {
	type shadow PublicMultiStringPropertyOperationParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PublicMultiStringPropertyOperationParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PublicNotAllFilterBranch struct {
	FilterBranches []PublicNotAllFilterBranchFilterBranchUnion `json:"filterBranches" api:"required"`
	// The operator used to combine filters within the branch (NOT_ALL).
	FilterBranchOperator string `json:"filterBranchOperator" api:"required"`
	// The type of the filter branch (NOT_ALL).
	//
	// Any of "NOT_ALL".
	FilterBranchType PublicNotAllFilterBranchFilterBranchType `json:"filterBranchType" api:"required"`
	Filters          []PublicNotAllFilterBranchFilterUnion    `json:"filters" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FilterBranches       respjson.Field
		FilterBranchOperator respjson.Field
		FilterBranchType     respjson.Field
		Filters              respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PublicNotAllFilterBranch) RawJSON() string { return r.JSON.raw }
func (r *PublicNotAllFilterBranch) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this PublicNotAllFilterBranch to a
// PublicNotAllFilterBranchParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// PublicNotAllFilterBranchParam.Overrides()
func (r PublicNotAllFilterBranch) ToParam() PublicNotAllFilterBranchParam {
	return param.Override[PublicNotAllFilterBranchParam](json.RawMessage(r.RawJSON()))
}

// PublicNotAllFilterBranchFilterBranchUnion contains all possible properties and
// values from [PublicOrFilterBranch], [PublicAndFilterBranch],
// [PublicNotAllFilterBranch], [PublicNotAnyFilterBranch],
// [PublicRestrictedFilterBranch], [PublicUnifiedEventsFilterBranch],
// [PublicPropertyAssociationFilterBranch], [PublicAssociationFilterBranch].
//
// Use the [PublicNotAllFilterBranchFilterBranchUnion.AsAny] method to switch on
// the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type PublicNotAllFilterBranchFilterBranchUnion struct {
	// This field is a union of [[]PublicOrFilterBranchFilterBranchUnion],
	// [[]PublicAndFilterBranchFilterBranchUnion],
	// [[]PublicNotAllFilterBranchFilterBranchUnion],
	// [[]PublicNotAnyFilterBranchFilterBranchUnion],
	// [[]PublicRestrictedFilterBranchFilterBranchUnion],
	// [[]PublicUnifiedEventsFilterBranchFilterBranchUnion],
	// [[]PublicPropertyAssociationFilterBranchFilterBranchUnion],
	// [[]PublicAssociationFilterBranchFilterBranchUnion]
	FilterBranches       PublicNotAllFilterBranchFilterBranchUnionFilterBranches `json:"filterBranches"`
	FilterBranchOperator string                                                  `json:"filterBranchOperator"`
	// Any of "OR", "AND", "NOT_ALL", "NOT_ANY", "RESTRICTED", "UNIFIED_EVENTS",
	// "PROPERTY_ASSOCIATION", "ASSOCIATION".
	FilterBranchType string `json:"filterBranchType"`
	// This field is a union of [[]PublicOrFilterBranchFilterUnion],
	// [[]PublicAndFilterBranchFilterUnion], [[]PublicNotAllFilterBranchFilterUnion],
	// [[]PublicNotAnyFilterBranchFilterUnion],
	// [[]PublicRestrictedFilterBranchFilterUnion],
	// [[]PublicUnifiedEventsFilterBranchFilterUnion],
	// [[]PublicPropertyAssociationFilterBranchFilterUnion],
	// [[]PublicAssociationFilterBranchFilterUnion]
	Filters PublicNotAllFilterBranchFilterBranchUnionFilters `json:"filters"`
	// This field is from variant [PublicUnifiedEventsFilterBranch].
	EventTypeID string `json:"eventTypeId"`
	Operator    string `json:"operator"`
	// This field is from variant [PublicUnifiedEventsFilterBranch].
	CoalescingRefineBy PublicUnifiedEventsFilterBranchCoalescingRefineByUnion `json:"coalescingRefineBy"`
	// This field is from variant [PublicUnifiedEventsFilterBranch].
	PruningRefineBy PublicUnifiedEventsFilterBranchPruningRefineByUnion `json:"pruningRefineBy"`
	ObjectTypeID    string                                              `json:"objectTypeId"`
	// This field is from variant [PublicPropertyAssociationFilterBranch].
	PropertyWithObjectID string `json:"propertyWithObjectId"`
	// This field is from variant [PublicAssociationFilterBranch].
	AssociationCategory string `json:"associationCategory"`
	// This field is from variant [PublicAssociationFilterBranch].
	AssociationTypeID int64 `json:"associationTypeId"`
	JSON              struct {
		FilterBranches       respjson.Field
		FilterBranchOperator respjson.Field
		FilterBranchType     respjson.Field
		Filters              respjson.Field
		EventTypeID          respjson.Field
		Operator             respjson.Field
		CoalescingRefineBy   respjson.Field
		PruningRefineBy      respjson.Field
		ObjectTypeID         respjson.Field
		PropertyWithObjectID respjson.Field
		AssociationCategory  respjson.Field
		AssociationTypeID    respjson.Field
		raw                  string
	} `json:"-"`
}

// anyPublicNotAllFilterBranchFilterBranch is implemented by each variant of
// [PublicNotAllFilterBranchFilterBranchUnion] to add type safety for the return
// type of [PublicNotAllFilterBranchFilterBranchUnion.AsAny]
type anyPublicNotAllFilterBranchFilterBranch interface {
	implPublicNotAllFilterBranchFilterBranchUnion()
}

func (PublicOrFilterBranch) implPublicNotAllFilterBranchFilterBranchUnion()                  {}
func (PublicAndFilterBranch) implPublicNotAllFilterBranchFilterBranchUnion()                 {}
func (PublicNotAllFilterBranch) implPublicNotAllFilterBranchFilterBranchUnion()              {}
func (PublicNotAnyFilterBranch) implPublicNotAllFilterBranchFilterBranchUnion()              {}
func (PublicRestrictedFilterBranch) implPublicNotAllFilterBranchFilterBranchUnion()          {}
func (PublicUnifiedEventsFilterBranch) implPublicNotAllFilterBranchFilterBranchUnion()       {}
func (PublicPropertyAssociationFilterBranch) implPublicNotAllFilterBranchFilterBranchUnion() {}
func (PublicAssociationFilterBranch) implPublicNotAllFilterBranchFilterBranchUnion()         {}

// Use the following switch statement to find the correct variant
//
//	switch variant := PublicNotAllFilterBranchFilterBranchUnion.AsAny().(type) {
//	case crm.PublicOrFilterBranch:
//	case crm.PublicAndFilterBranch:
//	case crm.PublicNotAllFilterBranch:
//	case crm.PublicNotAnyFilterBranch:
//	case crm.PublicRestrictedFilterBranch:
//	case crm.PublicUnifiedEventsFilterBranch:
//	case crm.PublicPropertyAssociationFilterBranch:
//	case crm.PublicAssociationFilterBranch:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u PublicNotAllFilterBranchFilterBranchUnion) AsAny() anyPublicNotAllFilterBranchFilterBranch {
	switch u.FilterBranchType {
	case "OR":
		return u.AsOr()
	case "AND":
		return u.AsAnd()
	case "NOT_ALL":
		return u.AsNotAll()
	case "NOT_ANY":
		return u.AsNotAny()
	case "RESTRICTED":
		return u.AsRestricted()
	case "UNIFIED_EVENTS":
		return u.AsUnifiedEvents()
	case "PROPERTY_ASSOCIATION":
		return u.AsPropertyAssociation()
	case "ASSOCIATION":
		return u.AsAssociation()
	}
	return nil
}

func (u PublicNotAllFilterBranchFilterBranchUnion) AsOr() (v PublicOrFilterBranch) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicNotAllFilterBranchFilterBranchUnion) AsAnd() (v PublicAndFilterBranch) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicNotAllFilterBranchFilterBranchUnion) AsNotAll() (v PublicNotAllFilterBranch) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicNotAllFilterBranchFilterBranchUnion) AsNotAny() (v PublicNotAnyFilterBranch) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicNotAllFilterBranchFilterBranchUnion) AsRestricted() (v PublicRestrictedFilterBranch) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicNotAllFilterBranchFilterBranchUnion) AsUnifiedEvents() (v PublicUnifiedEventsFilterBranch) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicNotAllFilterBranchFilterBranchUnion) AsPropertyAssociation() (v PublicPropertyAssociationFilterBranch) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicNotAllFilterBranchFilterBranchUnion) AsAssociation() (v PublicAssociationFilterBranch) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u PublicNotAllFilterBranchFilterBranchUnion) RawJSON() string { return u.JSON.raw }

func (r *PublicNotAllFilterBranchFilterBranchUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PublicNotAllFilterBranchFilterBranchUnionFilterBranches is an implicit subunion
// of [PublicNotAllFilterBranchFilterBranchUnion].
// PublicNotAllFilterBranchFilterBranchUnionFilterBranches provides convenient
// access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [PublicNotAllFilterBranchFilterBranchUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfPublicOrFilterBranchFilterBranches
// OfPublicAndFilterBranchFilterBranches OfPublicNotAllFilterBranchFilterBranches
// OfPublicNotAnyFilterBranchFilterBranches
// OfPublicRestrictedFilterBranchFilterBranches
// OfPublicUnifiedEventsFilterBranchFilterBranches
// OfPublicPropertyAssociationFilterBranchFilterBranches
// OfPublicAssociationFilterBranchFilterBranches]
type PublicNotAllFilterBranchFilterBranchUnionFilterBranches struct {
	// This field will be present if the value is a
	// [[]PublicOrFilterBranchFilterBranchUnion] instead of an object.
	OfPublicOrFilterBranchFilterBranches []PublicOrFilterBranchFilterBranchUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]PublicAndFilterBranchFilterBranchUnion] instead of an object.
	OfPublicAndFilterBranchFilterBranches []PublicAndFilterBranchFilterBranchUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]PublicNotAllFilterBranchFilterBranchUnion] instead of an object.
	OfPublicNotAllFilterBranchFilterBranches []PublicNotAllFilterBranchFilterBranchUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]PublicNotAnyFilterBranchFilterBranchUnion] instead of an object.
	OfPublicNotAnyFilterBranchFilterBranches []PublicNotAnyFilterBranchFilterBranchUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]PublicRestrictedFilterBranchFilterBranchUnion] instead of an object.
	OfPublicRestrictedFilterBranchFilterBranches []PublicRestrictedFilterBranchFilterBranchUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]PublicUnifiedEventsFilterBranchFilterBranchUnion] instead of an object.
	OfPublicUnifiedEventsFilterBranchFilterBranches []PublicUnifiedEventsFilterBranchFilterBranchUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]PublicPropertyAssociationFilterBranchFilterBranchUnion] instead of an object.
	OfPublicPropertyAssociationFilterBranchFilterBranches []PublicPropertyAssociationFilterBranchFilterBranchUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]PublicAssociationFilterBranchFilterBranchUnion] instead of an object.
	OfPublicAssociationFilterBranchFilterBranches []PublicAssociationFilterBranchFilterBranchUnion `json:",inline"`
	JSON                                          struct {
		OfPublicOrFilterBranchFilterBranches                  respjson.Field
		OfPublicAndFilterBranchFilterBranches                 respjson.Field
		OfPublicNotAllFilterBranchFilterBranches              respjson.Field
		OfPublicNotAnyFilterBranchFilterBranches              respjson.Field
		OfPublicRestrictedFilterBranchFilterBranches          respjson.Field
		OfPublicUnifiedEventsFilterBranchFilterBranches       respjson.Field
		OfPublicPropertyAssociationFilterBranchFilterBranches respjson.Field
		OfPublicAssociationFilterBranchFilterBranches         respjson.Field
		raw                                                   string
	} `json:"-"`
}

func (r *PublicNotAllFilterBranchFilterBranchUnionFilterBranches) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PublicNotAllFilterBranchFilterBranchUnionFilters is an implicit subunion of
// [PublicNotAllFilterBranchFilterBranchUnion].
// PublicNotAllFilterBranchFilterBranchUnionFilters provides convenient access to
// the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [PublicNotAllFilterBranchFilterBranchUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfPublicOrFilterBranchFilters OfPublicAndFilterBranchFilters
// OfPublicNotAllFilterBranchFilters OfPublicNotAnyFilterBranchFilters
// OfPublicRestrictedFilterBranchFilters OfPublicUnifiedEventsFilterBranchFilters
// OfPublicPropertyAssociationFilterBranchFilters
// OfPublicAssociationFilterBranchFilters]
type PublicNotAllFilterBranchFilterBranchUnionFilters struct {
	// This field will be present if the value is a [[]PublicOrFilterBranchFilterUnion]
	// instead of an object.
	OfPublicOrFilterBranchFilters []PublicOrFilterBranchFilterUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]PublicAndFilterBranchFilterUnion] instead of an object.
	OfPublicAndFilterBranchFilters []PublicAndFilterBranchFilterUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]PublicNotAllFilterBranchFilterUnion] instead of an object.
	OfPublicNotAllFilterBranchFilters []PublicNotAllFilterBranchFilterUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]PublicNotAnyFilterBranchFilterUnion] instead of an object.
	OfPublicNotAnyFilterBranchFilters []PublicNotAnyFilterBranchFilterUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]PublicRestrictedFilterBranchFilterUnion] instead of an object.
	OfPublicRestrictedFilterBranchFilters []PublicRestrictedFilterBranchFilterUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]PublicUnifiedEventsFilterBranchFilterUnion] instead of an object.
	OfPublicUnifiedEventsFilterBranchFilters []PublicUnifiedEventsFilterBranchFilterUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]PublicPropertyAssociationFilterBranchFilterUnion] instead of an object.
	OfPublicPropertyAssociationFilterBranchFilters []PublicPropertyAssociationFilterBranchFilterUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]PublicAssociationFilterBranchFilterUnion] instead of an object.
	OfPublicAssociationFilterBranchFilters []PublicAssociationFilterBranchFilterUnion `json:",inline"`
	JSON                                   struct {
		OfPublicOrFilterBranchFilters                  respjson.Field
		OfPublicAndFilterBranchFilters                 respjson.Field
		OfPublicNotAllFilterBranchFilters              respjson.Field
		OfPublicNotAnyFilterBranchFilters              respjson.Field
		OfPublicRestrictedFilterBranchFilters          respjson.Field
		OfPublicUnifiedEventsFilterBranchFilters       respjson.Field
		OfPublicPropertyAssociationFilterBranchFilters respjson.Field
		OfPublicAssociationFilterBranchFilters         respjson.Field
		raw                                            string
	} `json:"-"`
}

func (r *PublicNotAllFilterBranchFilterBranchUnionFilters) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The type of the filter branch (NOT_ALL).
type PublicNotAllFilterBranchFilterBranchType string

const (
	PublicNotAllFilterBranchFilterBranchTypeNotAll PublicNotAllFilterBranchFilterBranchType = "NOT_ALL"
)

// PublicNotAllFilterBranchFilterUnion contains all possible properties and values
// from [PublicPropertyFilter], [PublicAssociationInListFilter],
// [PublicPageViewAnalyticsFilter], [PublicCtaAnalyticsFilter],
// [PublicEventAnalyticsFilter], [PublicFormSubmissionFilter],
// [PublicFormSubmissionOnPageFilter], [PublicIntegrationEventFilter],
// [PublicEmailSubscriptionFilter], [PublicCommunicationSubscriptionFilter],
// [PublicCampaignInfluencedFilter], [PublicSurveyMonkeyFilter],
// [PublicSurveyMonkeyValueFilter], [PublicWebinarFilter],
// [PublicEmailEventFilter], [PublicPrivacyAnalyticsFilter],
// [PublicAdsSearchFilter], [PublicAdsTimeFilter], [PublicInListFilter],
// [PublicNumAssociationsFilter], [PublicUnifiedEventsFilter],
// [PublicPropertyAssociationInListFilter], [PublicConstantFilter].
//
// Use the [PublicNotAllFilterBranchFilterUnion.AsAny] method to switch on the
// variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type PublicNotAllFilterBranchFilterUnion struct {
	// Any of "PROPERTY", "ASSOCIATION", "PAGE_VIEW", "CTA", "EVENT",
	// "FORM_SUBMISSION", "FORM_SUBMISSION_ON_PAGE", "INTEGRATION_EVENT",
	// "EMAIL_SUBSCRIPTION", "COMMUNICATION_SUBSCRIPTION", "CAMPAIGN_INFLUENCED",
	// "SURVEY_MONKEY", "SURVEY_MONKEY_VALUE", "WEBINAR", "EMAIL_EVENT", "PRIVACY",
	// "ADS_SEARCH", "ADS_TIME", "IN_LIST", "NUM_ASSOCIATIONS", "UNIFIED_EVENTS",
	// "PROPERTY_ASSOCIATION", "CONSTANT".
	FilterType string `json:"filterType"`
	// This field is from variant [PublicPropertyFilter].
	Operation PublicPropertyFilterOperationUnion `json:"operation"`
	// This field is from variant [PublicPropertyFilter].
	Property            string `json:"property"`
	AssociationCategory string `json:"associationCategory"`
	AssociationTypeID   int64  `json:"associationTypeId"`
	// This field is a union of [PublicAssociationInListFilterCoalescingRefineByUnion],
	// [PublicPageViewAnalyticsFilterCoalescingRefineByUnion],
	// [PublicCtaAnalyticsFilterCoalescingRefineByUnion],
	// [PublicEventAnalyticsFilterCoalescingRefineByUnion],
	// [PublicFormSubmissionFilterCoalescingRefineByUnion],
	// [PublicFormSubmissionOnPageFilterCoalescingRefineByUnion],
	// [PublicNumAssociationsFilterCoalescingRefineByUnion],
	// [PublicUnifiedEventsFilterCoalescingRefineByUnion],
	// [PublicPropertyAssociationInListFilterCoalescingRefineByUnion]
	CoalescingRefineBy PublicNotAllFilterBranchFilterUnionCoalescingRefineBy `json:"coalescingRefineBy"`
	ListID             string                                                `json:"listId"`
	Operator           string                                                `json:"operator"`
	// This field is from variant [PublicAssociationInListFilter].
	ToObjectType   string `json:"toObjectType"`
	ToObjectTypeID string `json:"toObjectTypeId"`
	// This field is from variant [PublicPageViewAnalyticsFilter].
	PageURL string `json:"pageUrl"`
	// This field is from variant [PublicPageViewAnalyticsFilter].
	EnableTracking bool `json:"enableTracking"`
	// This field is a union of [PublicPageViewAnalyticsFilterPruningRefineByUnion],
	// [PublicCtaAnalyticsFilterPruningRefineByUnion],
	// [PublicEventAnalyticsFilterPruningRefineByUnion],
	// [PublicFormSubmissionFilterPruningRefineByUnion],
	// [PublicFormSubmissionOnPageFilterPruningRefineByUnion],
	// [PublicEmailEventFilterPruningRefineByUnion],
	// [PublicAdsTimeFilterPruningRefineByUnion],
	// [PublicUnifiedEventsFilterPruningRefineByUnion]
	PruningRefineBy PublicNotAllFilterBranchFilterUnionPruningRefineBy `json:"pruningRefineBy"`
	// This field is from variant [PublicCtaAnalyticsFilter].
	CtaName string `json:"ctaName"`
	// This field is from variant [PublicEventAnalyticsFilter].
	EventID string `json:"eventId"`
	FormID  string `json:"formId"`
	// This field is from variant [PublicFormSubmissionOnPageFilter].
	PageID string `json:"pageId"`
	// This field is a union of [int64], [string]
	EventTypeID PublicNotAllFilterBranchFilterUnionEventTypeID `json:"eventTypeId"`
	FilterLines []PublicEventFilterMetadata                    `json:"filterLines"`
	// This field is from variant [PublicEmailSubscriptionFilter].
	AcceptedStatuses []string `json:"acceptedStatuses"`
	SubscriptionIDs  []string `json:"subscriptionIds"`
	SubscriptionType string   `json:"subscriptionType"`
	// This field is from variant [PublicCommunicationSubscriptionFilter].
	AcceptedOptStates []string `json:"acceptedOptStates"`
	// This field is from variant [PublicCommunicationSubscriptionFilter].
	Channel string `json:"channel"`
	// This field is from variant [PublicCommunicationSubscriptionFilter].
	BusinessUnitID string `json:"businessUnitId"`
	// This field is from variant [PublicCampaignInfluencedFilter].
	CampaignID string `json:"campaignId"`
	SurveyID   string `json:"surveyId"`
	// This field is from variant [PublicSurveyMonkeyValueFilter].
	SurveyQuestion string `json:"surveyQuestion"`
	// This field is from variant [PublicSurveyMonkeyValueFilter].
	ValueComparison PublicSurveyMonkeyValueFilterValueComparisonUnion `json:"valueComparison"`
	// This field is from variant [PublicSurveyMonkeyValueFilter].
	SurveyAnswerColID string `json:"surveyAnswerColId"`
	// This field is from variant [PublicSurveyMonkeyValueFilter].
	SurveyAnswerRowID string `json:"surveyAnswerRowId"`
	// This field is from variant [PublicWebinarFilter].
	WebinarID string `json:"webinarId"`
	// This field is from variant [PublicEmailEventFilter].
	AppID string `json:"appId"`
	// This field is from variant [PublicEmailEventFilter].
	EmailID string `json:"emailId"`
	// This field is from variant [PublicEmailEventFilter].
	Level string `json:"level"`
	// This field is from variant [PublicEmailEventFilter].
	ClickURL string `json:"clickUrl"`
	// This field is from variant [PublicPrivacyAnalyticsFilter].
	PrivacyName string `json:"privacyName"`
	// This field is from variant [PublicAdsSearchFilter].
	AdNetwork string `json:"adNetwork"`
	// This field is from variant [PublicAdsSearchFilter].
	EntityType string `json:"entityType"`
	// This field is from variant [PublicAdsSearchFilter].
	SearchTerms []string `json:"searchTerms"`
	// This field is from variant [PublicAdsSearchFilter].
	SearchTermType string `json:"searchTermType"`
	// This field is from variant [PublicInListFilter].
	Metadata PublicInListFilterMetadata `json:"metadata"`
	// This field is from variant [PublicPropertyAssociationInListFilter].
	PropertyWithObjectID string `json:"propertyWithObjectId"`
	// This field is from variant [PublicConstantFilter].
	ShouldAccept bool `json:"shouldAccept"`
	// This field is from variant [PublicConstantFilter].
	Source string `json:"source"`
	JSON   struct {
		FilterType           respjson.Field
		Operation            respjson.Field
		Property             respjson.Field
		AssociationCategory  respjson.Field
		AssociationTypeID    respjson.Field
		CoalescingRefineBy   respjson.Field
		ListID               respjson.Field
		Operator             respjson.Field
		ToObjectType         respjson.Field
		ToObjectTypeID       respjson.Field
		PageURL              respjson.Field
		EnableTracking       respjson.Field
		PruningRefineBy      respjson.Field
		CtaName              respjson.Field
		EventID              respjson.Field
		FormID               respjson.Field
		PageID               respjson.Field
		EventTypeID          respjson.Field
		FilterLines          respjson.Field
		AcceptedStatuses     respjson.Field
		SubscriptionIDs      respjson.Field
		SubscriptionType     respjson.Field
		AcceptedOptStates    respjson.Field
		Channel              respjson.Field
		BusinessUnitID       respjson.Field
		CampaignID           respjson.Field
		SurveyID             respjson.Field
		SurveyQuestion       respjson.Field
		ValueComparison      respjson.Field
		SurveyAnswerColID    respjson.Field
		SurveyAnswerRowID    respjson.Field
		WebinarID            respjson.Field
		AppID                respjson.Field
		EmailID              respjson.Field
		Level                respjson.Field
		ClickURL             respjson.Field
		PrivacyName          respjson.Field
		AdNetwork            respjson.Field
		EntityType           respjson.Field
		SearchTerms          respjson.Field
		SearchTermType       respjson.Field
		Metadata             respjson.Field
		PropertyWithObjectID respjson.Field
		ShouldAccept         respjson.Field
		Source               respjson.Field
		raw                  string
	} `json:"-"`
}

// anyPublicNotAllFilterBranchFilter is implemented by each variant of
// [PublicNotAllFilterBranchFilterUnion] to add type safety for the return type of
// [PublicNotAllFilterBranchFilterUnion.AsAny]
type anyPublicNotAllFilterBranchFilter interface {
	implPublicNotAllFilterBranchFilterUnion()
}

func (PublicPropertyFilter) implPublicNotAllFilterBranchFilterUnion()                  {}
func (PublicAssociationInListFilter) implPublicNotAllFilterBranchFilterUnion()         {}
func (PublicPageViewAnalyticsFilter) implPublicNotAllFilterBranchFilterUnion()         {}
func (PublicCtaAnalyticsFilter) implPublicNotAllFilterBranchFilterUnion()              {}
func (PublicEventAnalyticsFilter) implPublicNotAllFilterBranchFilterUnion()            {}
func (PublicFormSubmissionFilter) implPublicNotAllFilterBranchFilterUnion()            {}
func (PublicFormSubmissionOnPageFilter) implPublicNotAllFilterBranchFilterUnion()      {}
func (PublicIntegrationEventFilter) implPublicNotAllFilterBranchFilterUnion()          {}
func (PublicEmailSubscriptionFilter) implPublicNotAllFilterBranchFilterUnion()         {}
func (PublicCommunicationSubscriptionFilter) implPublicNotAllFilterBranchFilterUnion() {}
func (PublicCampaignInfluencedFilter) implPublicNotAllFilterBranchFilterUnion()        {}
func (PublicSurveyMonkeyFilter) implPublicNotAllFilterBranchFilterUnion()              {}
func (PublicSurveyMonkeyValueFilter) implPublicNotAllFilterBranchFilterUnion()         {}
func (PublicWebinarFilter) implPublicNotAllFilterBranchFilterUnion()                   {}
func (PublicEmailEventFilter) implPublicNotAllFilterBranchFilterUnion()                {}
func (PublicPrivacyAnalyticsFilter) implPublicNotAllFilterBranchFilterUnion()          {}
func (PublicAdsSearchFilter) implPublicNotAllFilterBranchFilterUnion()                 {}
func (PublicAdsTimeFilter) implPublicNotAllFilterBranchFilterUnion()                   {}
func (PublicInListFilter) implPublicNotAllFilterBranchFilterUnion()                    {}
func (PublicNumAssociationsFilter) implPublicNotAllFilterBranchFilterUnion()           {}
func (PublicUnifiedEventsFilter) implPublicNotAllFilterBranchFilterUnion()             {}
func (PublicPropertyAssociationInListFilter) implPublicNotAllFilterBranchFilterUnion() {}
func (PublicConstantFilter) implPublicNotAllFilterBranchFilterUnion()                  {}

// Use the following switch statement to find the correct variant
//
//	switch variant := PublicNotAllFilterBranchFilterUnion.AsAny().(type) {
//	case crm.PublicPropertyFilter:
//	case crm.PublicAssociationInListFilter:
//	case crm.PublicPageViewAnalyticsFilter:
//	case crm.PublicCtaAnalyticsFilter:
//	case crm.PublicEventAnalyticsFilter:
//	case crm.PublicFormSubmissionFilter:
//	case crm.PublicFormSubmissionOnPageFilter:
//	case crm.PublicIntegrationEventFilter:
//	case crm.PublicEmailSubscriptionFilter:
//	case crm.PublicCommunicationSubscriptionFilter:
//	case crm.PublicCampaignInfluencedFilter:
//	case crm.PublicSurveyMonkeyFilter:
//	case crm.PublicSurveyMonkeyValueFilter:
//	case crm.PublicWebinarFilter:
//	case crm.PublicEmailEventFilter:
//	case crm.PublicPrivacyAnalyticsFilter:
//	case crm.PublicAdsSearchFilter:
//	case crm.PublicAdsTimeFilter:
//	case crm.PublicInListFilter:
//	case crm.PublicNumAssociationsFilter:
//	case crm.PublicUnifiedEventsFilter:
//	case crm.PublicPropertyAssociationInListFilter:
//	case crm.PublicConstantFilter:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u PublicNotAllFilterBranchFilterUnion) AsAny() anyPublicNotAllFilterBranchFilter {
	switch u.FilterType {
	case "PROPERTY":
		return u.AsProperty()
	case "ASSOCIATION":
		return u.AsAssociation()
	case "PAGE_VIEW":
		return u.AsPageView()
	case "CTA":
		return u.AsCta()
	case "EVENT":
		return u.AsEvent()
	case "FORM_SUBMISSION":
		return u.AsFormSubmission()
	case "FORM_SUBMISSION_ON_PAGE":
		return u.AsFormSubmissionOnPage()
	case "INTEGRATION_EVENT":
		return u.AsIntegrationEvent()
	case "EMAIL_SUBSCRIPTION":
		return u.AsEmailSubscription()
	case "COMMUNICATION_SUBSCRIPTION":
		return u.AsCommunicationSubscription()
	case "CAMPAIGN_INFLUENCED":
		return u.AsCampaignInfluenced()
	case "SURVEY_MONKEY":
		return u.AsSurveyMonkey()
	case "SURVEY_MONKEY_VALUE":
		return u.AsSurveyMonkeyValue()
	case "WEBINAR":
		return u.AsWebinar()
	case "EMAIL_EVENT":
		return u.AsEmailEvent()
	case "PRIVACY":
		return u.AsPrivacy()
	case "ADS_SEARCH":
		return u.AsAdsSearch()
	case "ADS_TIME":
		return u.AsAdsTime()
	case "IN_LIST":
		return u.AsInList()
	case "NUM_ASSOCIATIONS":
		return u.AsNumAssociations()
	case "UNIFIED_EVENTS":
		return u.AsUnifiedEvents()
	case "PROPERTY_ASSOCIATION":
		return u.AsPropertyAssociation()
	case "CONSTANT":
		return u.AsConstant()
	}
	return nil
}

func (u PublicNotAllFilterBranchFilterUnion) AsProperty() (v PublicPropertyFilter) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicNotAllFilterBranchFilterUnion) AsAssociation() (v PublicAssociationInListFilter) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicNotAllFilterBranchFilterUnion) AsPageView() (v PublicPageViewAnalyticsFilter) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicNotAllFilterBranchFilterUnion) AsCta() (v PublicCtaAnalyticsFilter) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicNotAllFilterBranchFilterUnion) AsEvent() (v PublicEventAnalyticsFilter) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicNotAllFilterBranchFilterUnion) AsFormSubmission() (v PublicFormSubmissionFilter) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicNotAllFilterBranchFilterUnion) AsFormSubmissionOnPage() (v PublicFormSubmissionOnPageFilter) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicNotAllFilterBranchFilterUnion) AsIntegrationEvent() (v PublicIntegrationEventFilter) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicNotAllFilterBranchFilterUnion) AsEmailSubscription() (v PublicEmailSubscriptionFilter) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicNotAllFilterBranchFilterUnion) AsCommunicationSubscription() (v PublicCommunicationSubscriptionFilter) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicNotAllFilterBranchFilterUnion) AsCampaignInfluenced() (v PublicCampaignInfluencedFilter) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicNotAllFilterBranchFilterUnion) AsSurveyMonkey() (v PublicSurveyMonkeyFilter) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicNotAllFilterBranchFilterUnion) AsSurveyMonkeyValue() (v PublicSurveyMonkeyValueFilter) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicNotAllFilterBranchFilterUnion) AsWebinar() (v PublicWebinarFilter) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicNotAllFilterBranchFilterUnion) AsEmailEvent() (v PublicEmailEventFilter) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicNotAllFilterBranchFilterUnion) AsPrivacy() (v PublicPrivacyAnalyticsFilter) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicNotAllFilterBranchFilterUnion) AsAdsSearch() (v PublicAdsSearchFilter) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicNotAllFilterBranchFilterUnion) AsAdsTime() (v PublicAdsTimeFilter) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicNotAllFilterBranchFilterUnion) AsInList() (v PublicInListFilter) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicNotAllFilterBranchFilterUnion) AsNumAssociations() (v PublicNumAssociationsFilter) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicNotAllFilterBranchFilterUnion) AsUnifiedEvents() (v PublicUnifiedEventsFilter) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicNotAllFilterBranchFilterUnion) AsPropertyAssociation() (v PublicPropertyAssociationInListFilter) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicNotAllFilterBranchFilterUnion) AsConstant() (v PublicConstantFilter) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u PublicNotAllFilterBranchFilterUnion) RawJSON() string { return u.JSON.raw }

func (r *PublicNotAllFilterBranchFilterUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PublicNotAllFilterBranchFilterUnionCoalescingRefineBy is an implicit subunion of
// [PublicNotAllFilterBranchFilterUnion].
// PublicNotAllFilterBranchFilterUnionCoalescingRefineBy provides convenient access
// to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [PublicNotAllFilterBranchFilterUnion].
type PublicNotAllFilterBranchFilterUnionCoalescingRefineBy struct {
	Type string `json:"type"`
	// This field is from variant
	// [PublicAssociationInListFilterCoalescingRefineByUnion],
	// [PublicPageViewAnalyticsFilterCoalescingRefineByUnion],
	// [PublicCtaAnalyticsFilterCoalescingRefineByUnion],
	// [PublicEventAnalyticsFilterCoalescingRefineByUnion],
	// [PublicFormSubmissionFilterCoalescingRefineByUnion],
	// [PublicFormSubmissionOnPageFilterCoalescingRefineByUnion],
	// [PublicNumAssociationsFilterCoalescingRefineByUnion],
	// [PublicUnifiedEventsFilterCoalescingRefineByUnion],
	// [PublicPropertyAssociationInListFilterCoalescingRefineByUnion].
	MaxOccurrences int64 `json:"maxOccurrences"`
	// This field is from variant
	// [PublicAssociationInListFilterCoalescingRefineByUnion],
	// [PublicPageViewAnalyticsFilterCoalescingRefineByUnion],
	// [PublicCtaAnalyticsFilterCoalescingRefineByUnion],
	// [PublicEventAnalyticsFilterCoalescingRefineByUnion],
	// [PublicFormSubmissionFilterCoalescingRefineByUnion],
	// [PublicFormSubmissionOnPageFilterCoalescingRefineByUnion],
	// [PublicNumAssociationsFilterCoalescingRefineByUnion],
	// [PublicUnifiedEventsFilterCoalescingRefineByUnion],
	// [PublicPropertyAssociationInListFilterCoalescingRefineByUnion].
	MinOccurrences int64 `json:"minOccurrences"`
	// This field is from variant
	// [PublicAssociationInListFilterCoalescingRefineByUnion],
	// [PublicPageViewAnalyticsFilterCoalescingRefineByUnion],
	// [PublicCtaAnalyticsFilterCoalescingRefineByUnion],
	// [PublicEventAnalyticsFilterCoalescingRefineByUnion],
	// [PublicFormSubmissionFilterCoalescingRefineByUnion],
	// [PublicFormSubmissionOnPageFilterCoalescingRefineByUnion],
	// [PublicNumAssociationsFilterCoalescingRefineByUnion],
	// [PublicUnifiedEventsFilterCoalescingRefineByUnion],
	// [PublicPropertyAssociationInListFilterCoalescingRefineByUnion].
	SetType    string `json:"setType"`
	Comparison string `json:"comparison"`
	// This field is from variant
	// [PublicAssociationInListFilterCoalescingRefineByUnion],
	// [PublicPageViewAnalyticsFilterCoalescingRefineByUnion],
	// [PublicCtaAnalyticsFilterCoalescingRefineByUnion],
	// [PublicEventAnalyticsFilterCoalescingRefineByUnion],
	// [PublicFormSubmissionFilterCoalescingRefineByUnion],
	// [PublicFormSubmissionOnPageFilterCoalescingRefineByUnion],
	// [PublicNumAssociationsFilterCoalescingRefineByUnion],
	// [PublicUnifiedEventsFilterCoalescingRefineByUnion],
	// [PublicPropertyAssociationInListFilterCoalescingRefineByUnion].
	TimeOffset PublicTimeOffset `json:"timeOffset"`
	// This field is from variant
	// [PublicAssociationInListFilterCoalescingRefineByUnion],
	// [PublicPageViewAnalyticsFilterCoalescingRefineByUnion],
	// [PublicCtaAnalyticsFilterCoalescingRefineByUnion],
	// [PublicEventAnalyticsFilterCoalescingRefineByUnion],
	// [PublicFormSubmissionFilterCoalescingRefineByUnion],
	// [PublicFormSubmissionOnPageFilterCoalescingRefineByUnion],
	// [PublicNumAssociationsFilterCoalescingRefineByUnion],
	// [PublicUnifiedEventsFilterCoalescingRefineByUnion],
	// [PublicPropertyAssociationInListFilterCoalescingRefineByUnion].
	LowerBoundOffset PublicTimeOffset `json:"lowerBoundOffset"`
	RangeType        string           `json:"rangeType"`
	// This field is from variant
	// [PublicAssociationInListFilterCoalescingRefineByUnion],
	// [PublicPageViewAnalyticsFilterCoalescingRefineByUnion],
	// [PublicCtaAnalyticsFilterCoalescingRefineByUnion],
	// [PublicEventAnalyticsFilterCoalescingRefineByUnion],
	// [PublicFormSubmissionFilterCoalescingRefineByUnion],
	// [PublicFormSubmissionOnPageFilterCoalescingRefineByUnion],
	// [PublicNumAssociationsFilterCoalescingRefineByUnion],
	// [PublicUnifiedEventsFilterCoalescingRefineByUnion],
	// [PublicPropertyAssociationInListFilterCoalescingRefineByUnion].
	UpperBoundOffset PublicTimeOffset `json:"upperBoundOffset"`
	// This field is from variant
	// [PublicAssociationInListFilterCoalescingRefineByUnion],
	// [PublicPageViewAnalyticsFilterCoalescingRefineByUnion],
	// [PublicCtaAnalyticsFilterCoalescingRefineByUnion],
	// [PublicEventAnalyticsFilterCoalescingRefineByUnion],
	// [PublicFormSubmissionFilterCoalescingRefineByUnion],
	// [PublicFormSubmissionOnPageFilterCoalescingRefineByUnion],
	// [PublicNumAssociationsFilterCoalescingRefineByUnion],
	// [PublicUnifiedEventsFilterCoalescingRefineByUnion],
	// [PublicPropertyAssociationInListFilterCoalescingRefineByUnion].
	Timestamp int64 `json:"timestamp"`
	// This field is from variant
	// [PublicAssociationInListFilterCoalescingRefineByUnion],
	// [PublicPageViewAnalyticsFilterCoalescingRefineByUnion],
	// [PublicCtaAnalyticsFilterCoalescingRefineByUnion],
	// [PublicEventAnalyticsFilterCoalescingRefineByUnion],
	// [PublicFormSubmissionFilterCoalescingRefineByUnion],
	// [PublicFormSubmissionOnPageFilterCoalescingRefineByUnion],
	// [PublicNumAssociationsFilterCoalescingRefineByUnion],
	// [PublicUnifiedEventsFilterCoalescingRefineByUnion],
	// [PublicPropertyAssociationInListFilterCoalescingRefineByUnion].
	LowerTimestamp int64 `json:"lowerTimestamp"`
	// This field is from variant
	// [PublicAssociationInListFilterCoalescingRefineByUnion],
	// [PublicPageViewAnalyticsFilterCoalescingRefineByUnion],
	// [PublicCtaAnalyticsFilterCoalescingRefineByUnion],
	// [PublicEventAnalyticsFilterCoalescingRefineByUnion],
	// [PublicFormSubmissionFilterCoalescingRefineByUnion],
	// [PublicFormSubmissionOnPageFilterCoalescingRefineByUnion],
	// [PublicNumAssociationsFilterCoalescingRefineByUnion],
	// [PublicUnifiedEventsFilterCoalescingRefineByUnion],
	// [PublicPropertyAssociationInListFilterCoalescingRefineByUnion].
	UpperTimestamp               int64  `json:"upperTimestamp"`
	IncludeObjectsWithNoValueSet bool   `json:"includeObjectsWithNoValueSet"`
	OperationType                string `json:"operationType"`
	Operator                     string `json:"operator"`
	// This field is from variant
	// [PublicAssociationInListFilterCoalescingRefineByUnion],
	// [PublicPageViewAnalyticsFilterCoalescingRefineByUnion],
	// [PublicCtaAnalyticsFilterCoalescingRefineByUnion],
	// [PublicEventAnalyticsFilterCoalescingRefineByUnion],
	// [PublicFormSubmissionFilterCoalescingRefineByUnion],
	// [PublicFormSubmissionOnPageFilterCoalescingRefineByUnion],
	// [PublicNumAssociationsFilterCoalescingRefineByUnion],
	// [PublicUnifiedEventsFilterCoalescingRefineByUnion],
	// [PublicPropertyAssociationInListFilterCoalescingRefineByUnion].
	TimePoint PublicTimePointOperationTimePointUnion `json:"timePoint"`
	// This field is from variant
	// [PublicAssociationInListFilterCoalescingRefineByUnion],
	// [PublicPageViewAnalyticsFilterCoalescingRefineByUnion],
	// [PublicCtaAnalyticsFilterCoalescingRefineByUnion],
	// [PublicEventAnalyticsFilterCoalescingRefineByUnion],
	// [PublicFormSubmissionFilterCoalescingRefineByUnion],
	// [PublicFormSubmissionOnPageFilterCoalescingRefineByUnion],
	// [PublicNumAssociationsFilterCoalescingRefineByUnion],
	// [PublicUnifiedEventsFilterCoalescingRefineByUnion],
	// [PublicPropertyAssociationInListFilterCoalescingRefineByUnion].
	EndpointBehavior string `json:"endpointBehavior"`
	PropertyParser   string `json:"propertyParser"`
	// This field is from variant
	// [PublicAssociationInListFilterCoalescingRefineByUnion],
	// [PublicPageViewAnalyticsFilterCoalescingRefineByUnion],
	// [PublicCtaAnalyticsFilterCoalescingRefineByUnion],
	// [PublicEventAnalyticsFilterCoalescingRefineByUnion],
	// [PublicFormSubmissionFilterCoalescingRefineByUnion],
	// [PublicFormSubmissionOnPageFilterCoalescingRefineByUnion],
	// [PublicNumAssociationsFilterCoalescingRefineByUnion],
	// [PublicUnifiedEventsFilterCoalescingRefineByUnion],
	// [PublicPropertyAssociationInListFilterCoalescingRefineByUnion].
	LowerBoundTimePoint PublicRangedTimeOperationLowerBoundTimePointUnion `json:"lowerBoundTimePoint"`
	// This field is from variant
	// [PublicAssociationInListFilterCoalescingRefineByUnion],
	// [PublicPageViewAnalyticsFilterCoalescingRefineByUnion],
	// [PublicCtaAnalyticsFilterCoalescingRefineByUnion],
	// [PublicEventAnalyticsFilterCoalescingRefineByUnion],
	// [PublicFormSubmissionFilterCoalescingRefineByUnion],
	// [PublicFormSubmissionOnPageFilterCoalescingRefineByUnion],
	// [PublicNumAssociationsFilterCoalescingRefineByUnion],
	// [PublicUnifiedEventsFilterCoalescingRefineByUnion],
	// [PublicPropertyAssociationInListFilterCoalescingRefineByUnion].
	UpperBoundTimePoint PublicRangedTimeOperationUpperBoundTimePointUnion `json:"upperBoundTimePoint"`
	// This field is from variant
	// [PublicAssociationInListFilterCoalescingRefineByUnion],
	// [PublicPageViewAnalyticsFilterCoalescingRefineByUnion],
	// [PublicCtaAnalyticsFilterCoalescingRefineByUnion],
	// [PublicEventAnalyticsFilterCoalescingRefineByUnion],
	// [PublicFormSubmissionFilterCoalescingRefineByUnion],
	// [PublicFormSubmissionOnPageFilterCoalescingRefineByUnion],
	// [PublicNumAssociationsFilterCoalescingRefineByUnion],
	// [PublicUnifiedEventsFilterCoalescingRefineByUnion],
	// [PublicPropertyAssociationInListFilterCoalescingRefineByUnion].
	LowerBoundEndpointBehavior string `json:"lowerBoundEndpointBehavior"`
	// This field is from variant
	// [PublicAssociationInListFilterCoalescingRefineByUnion],
	// [PublicPageViewAnalyticsFilterCoalescingRefineByUnion],
	// [PublicCtaAnalyticsFilterCoalescingRefineByUnion],
	// [PublicEventAnalyticsFilterCoalescingRefineByUnion],
	// [PublicFormSubmissionFilterCoalescingRefineByUnion],
	// [PublicFormSubmissionOnPageFilterCoalescingRefineByUnion],
	// [PublicNumAssociationsFilterCoalescingRefineByUnion],
	// [PublicUnifiedEventsFilterCoalescingRefineByUnion],
	// [PublicPropertyAssociationInListFilterCoalescingRefineByUnion].
	UpperBoundEndpointBehavior string `json:"upperBoundEndpointBehavior"`
	JSON                       struct {
		Type                         respjson.Field
		MaxOccurrences               respjson.Field
		MinOccurrences               respjson.Field
		SetType                      respjson.Field
		Comparison                   respjson.Field
		TimeOffset                   respjson.Field
		LowerBoundOffset             respjson.Field
		RangeType                    respjson.Field
		UpperBoundOffset             respjson.Field
		Timestamp                    respjson.Field
		LowerTimestamp               respjson.Field
		UpperTimestamp               respjson.Field
		IncludeObjectsWithNoValueSet respjson.Field
		OperationType                respjson.Field
		Operator                     respjson.Field
		TimePoint                    respjson.Field
		EndpointBehavior             respjson.Field
		PropertyParser               respjson.Field
		LowerBoundTimePoint          respjson.Field
		UpperBoundTimePoint          respjson.Field
		LowerBoundEndpointBehavior   respjson.Field
		UpperBoundEndpointBehavior   respjson.Field
		raw                          string
	} `json:"-"`
}

func (r *PublicNotAllFilterBranchFilterUnionCoalescingRefineBy) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PublicNotAllFilterBranchFilterUnionPruningRefineBy is an implicit subunion of
// [PublicNotAllFilterBranchFilterUnion].
// PublicNotAllFilterBranchFilterUnionPruningRefineBy provides convenient access to
// the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [PublicNotAllFilterBranchFilterUnion].
type PublicNotAllFilterBranchFilterUnionPruningRefineBy struct {
	Type string `json:"type"`
	// This field is from variant [PublicPageViewAnalyticsFilterPruningRefineByUnion],
	// [PublicCtaAnalyticsFilterPruningRefineByUnion],
	// [PublicEventAnalyticsFilterPruningRefineByUnion],
	// [PublicFormSubmissionFilterPruningRefineByUnion],
	// [PublicFormSubmissionOnPageFilterPruningRefineByUnion],
	// [PublicEmailEventFilterPruningRefineByUnion],
	// [PublicAdsTimeFilterPruningRefineByUnion],
	// [PublicUnifiedEventsFilterPruningRefineByUnion].
	MaxOccurrences int64 `json:"maxOccurrences"`
	// This field is from variant [PublicPageViewAnalyticsFilterPruningRefineByUnion],
	// [PublicCtaAnalyticsFilterPruningRefineByUnion],
	// [PublicEventAnalyticsFilterPruningRefineByUnion],
	// [PublicFormSubmissionFilterPruningRefineByUnion],
	// [PublicFormSubmissionOnPageFilterPruningRefineByUnion],
	// [PublicEmailEventFilterPruningRefineByUnion],
	// [PublicAdsTimeFilterPruningRefineByUnion],
	// [PublicUnifiedEventsFilterPruningRefineByUnion].
	MinOccurrences int64 `json:"minOccurrences"`
	// This field is from variant [PublicPageViewAnalyticsFilterPruningRefineByUnion],
	// [PublicCtaAnalyticsFilterPruningRefineByUnion],
	// [PublicEventAnalyticsFilterPruningRefineByUnion],
	// [PublicFormSubmissionFilterPruningRefineByUnion],
	// [PublicFormSubmissionOnPageFilterPruningRefineByUnion],
	// [PublicEmailEventFilterPruningRefineByUnion],
	// [PublicAdsTimeFilterPruningRefineByUnion],
	// [PublicUnifiedEventsFilterPruningRefineByUnion].
	SetType    string `json:"setType"`
	Comparison string `json:"comparison"`
	// This field is from variant [PublicPageViewAnalyticsFilterPruningRefineByUnion],
	// [PublicCtaAnalyticsFilterPruningRefineByUnion],
	// [PublicEventAnalyticsFilterPruningRefineByUnion],
	// [PublicFormSubmissionFilterPruningRefineByUnion],
	// [PublicFormSubmissionOnPageFilterPruningRefineByUnion],
	// [PublicEmailEventFilterPruningRefineByUnion],
	// [PublicAdsTimeFilterPruningRefineByUnion],
	// [PublicUnifiedEventsFilterPruningRefineByUnion].
	TimeOffset PublicTimeOffset `json:"timeOffset"`
	// This field is from variant [PublicPageViewAnalyticsFilterPruningRefineByUnion],
	// [PublicCtaAnalyticsFilterPruningRefineByUnion],
	// [PublicEventAnalyticsFilterPruningRefineByUnion],
	// [PublicFormSubmissionFilterPruningRefineByUnion],
	// [PublicFormSubmissionOnPageFilterPruningRefineByUnion],
	// [PublicEmailEventFilterPruningRefineByUnion],
	// [PublicAdsTimeFilterPruningRefineByUnion],
	// [PublicUnifiedEventsFilterPruningRefineByUnion].
	LowerBoundOffset PublicTimeOffset `json:"lowerBoundOffset"`
	RangeType        string           `json:"rangeType"`
	// This field is from variant [PublicPageViewAnalyticsFilterPruningRefineByUnion],
	// [PublicCtaAnalyticsFilterPruningRefineByUnion],
	// [PublicEventAnalyticsFilterPruningRefineByUnion],
	// [PublicFormSubmissionFilterPruningRefineByUnion],
	// [PublicFormSubmissionOnPageFilterPruningRefineByUnion],
	// [PublicEmailEventFilterPruningRefineByUnion],
	// [PublicAdsTimeFilterPruningRefineByUnion],
	// [PublicUnifiedEventsFilterPruningRefineByUnion].
	UpperBoundOffset PublicTimeOffset `json:"upperBoundOffset"`
	// This field is from variant [PublicPageViewAnalyticsFilterPruningRefineByUnion],
	// [PublicCtaAnalyticsFilterPruningRefineByUnion],
	// [PublicEventAnalyticsFilterPruningRefineByUnion],
	// [PublicFormSubmissionFilterPruningRefineByUnion],
	// [PublicFormSubmissionOnPageFilterPruningRefineByUnion],
	// [PublicEmailEventFilterPruningRefineByUnion],
	// [PublicAdsTimeFilterPruningRefineByUnion],
	// [PublicUnifiedEventsFilterPruningRefineByUnion].
	Timestamp int64 `json:"timestamp"`
	// This field is from variant [PublicPageViewAnalyticsFilterPruningRefineByUnion],
	// [PublicCtaAnalyticsFilterPruningRefineByUnion],
	// [PublicEventAnalyticsFilterPruningRefineByUnion],
	// [PublicFormSubmissionFilterPruningRefineByUnion],
	// [PublicFormSubmissionOnPageFilterPruningRefineByUnion],
	// [PublicEmailEventFilterPruningRefineByUnion],
	// [PublicAdsTimeFilterPruningRefineByUnion],
	// [PublicUnifiedEventsFilterPruningRefineByUnion].
	LowerTimestamp int64 `json:"lowerTimestamp"`
	// This field is from variant [PublicPageViewAnalyticsFilterPruningRefineByUnion],
	// [PublicCtaAnalyticsFilterPruningRefineByUnion],
	// [PublicEventAnalyticsFilterPruningRefineByUnion],
	// [PublicFormSubmissionFilterPruningRefineByUnion],
	// [PublicFormSubmissionOnPageFilterPruningRefineByUnion],
	// [PublicEmailEventFilterPruningRefineByUnion],
	// [PublicAdsTimeFilterPruningRefineByUnion],
	// [PublicUnifiedEventsFilterPruningRefineByUnion].
	UpperTimestamp               int64  `json:"upperTimestamp"`
	IncludeObjectsWithNoValueSet bool   `json:"includeObjectsWithNoValueSet"`
	OperationType                string `json:"operationType"`
	Operator                     string `json:"operator"`
	// This field is from variant [PublicPageViewAnalyticsFilterPruningRefineByUnion],
	// [PublicCtaAnalyticsFilterPruningRefineByUnion],
	// [PublicEventAnalyticsFilterPruningRefineByUnion],
	// [PublicFormSubmissionFilterPruningRefineByUnion],
	// [PublicFormSubmissionOnPageFilterPruningRefineByUnion],
	// [PublicEmailEventFilterPruningRefineByUnion],
	// [PublicAdsTimeFilterPruningRefineByUnion],
	// [PublicUnifiedEventsFilterPruningRefineByUnion].
	TimePoint PublicTimePointOperationTimePointUnion `json:"timePoint"`
	// This field is from variant [PublicPageViewAnalyticsFilterPruningRefineByUnion],
	// [PublicCtaAnalyticsFilterPruningRefineByUnion],
	// [PublicEventAnalyticsFilterPruningRefineByUnion],
	// [PublicFormSubmissionFilterPruningRefineByUnion],
	// [PublicFormSubmissionOnPageFilterPruningRefineByUnion],
	// [PublicEmailEventFilterPruningRefineByUnion],
	// [PublicAdsTimeFilterPruningRefineByUnion],
	// [PublicUnifiedEventsFilterPruningRefineByUnion].
	EndpointBehavior string `json:"endpointBehavior"`
	PropertyParser   string `json:"propertyParser"`
	// This field is from variant [PublicPageViewAnalyticsFilterPruningRefineByUnion],
	// [PublicCtaAnalyticsFilterPruningRefineByUnion],
	// [PublicEventAnalyticsFilterPruningRefineByUnion],
	// [PublicFormSubmissionFilterPruningRefineByUnion],
	// [PublicFormSubmissionOnPageFilterPruningRefineByUnion],
	// [PublicEmailEventFilterPruningRefineByUnion],
	// [PublicAdsTimeFilterPruningRefineByUnion],
	// [PublicUnifiedEventsFilterPruningRefineByUnion].
	LowerBoundTimePoint PublicRangedTimeOperationLowerBoundTimePointUnion `json:"lowerBoundTimePoint"`
	// This field is from variant [PublicPageViewAnalyticsFilterPruningRefineByUnion],
	// [PublicCtaAnalyticsFilterPruningRefineByUnion],
	// [PublicEventAnalyticsFilterPruningRefineByUnion],
	// [PublicFormSubmissionFilterPruningRefineByUnion],
	// [PublicFormSubmissionOnPageFilterPruningRefineByUnion],
	// [PublicEmailEventFilterPruningRefineByUnion],
	// [PublicAdsTimeFilterPruningRefineByUnion],
	// [PublicUnifiedEventsFilterPruningRefineByUnion].
	UpperBoundTimePoint PublicRangedTimeOperationUpperBoundTimePointUnion `json:"upperBoundTimePoint"`
	// This field is from variant [PublicPageViewAnalyticsFilterPruningRefineByUnion],
	// [PublicCtaAnalyticsFilterPruningRefineByUnion],
	// [PublicEventAnalyticsFilterPruningRefineByUnion],
	// [PublicFormSubmissionFilterPruningRefineByUnion],
	// [PublicFormSubmissionOnPageFilterPruningRefineByUnion],
	// [PublicEmailEventFilterPruningRefineByUnion],
	// [PublicAdsTimeFilterPruningRefineByUnion],
	// [PublicUnifiedEventsFilterPruningRefineByUnion].
	LowerBoundEndpointBehavior string `json:"lowerBoundEndpointBehavior"`
	// This field is from variant [PublicPageViewAnalyticsFilterPruningRefineByUnion],
	// [PublicCtaAnalyticsFilterPruningRefineByUnion],
	// [PublicEventAnalyticsFilterPruningRefineByUnion],
	// [PublicFormSubmissionFilterPruningRefineByUnion],
	// [PublicFormSubmissionOnPageFilterPruningRefineByUnion],
	// [PublicEmailEventFilterPruningRefineByUnion],
	// [PublicAdsTimeFilterPruningRefineByUnion],
	// [PublicUnifiedEventsFilterPruningRefineByUnion].
	UpperBoundEndpointBehavior string `json:"upperBoundEndpointBehavior"`
	JSON                       struct {
		Type                         respjson.Field
		MaxOccurrences               respjson.Field
		MinOccurrences               respjson.Field
		SetType                      respjson.Field
		Comparison                   respjson.Field
		TimeOffset                   respjson.Field
		LowerBoundOffset             respjson.Field
		RangeType                    respjson.Field
		UpperBoundOffset             respjson.Field
		Timestamp                    respjson.Field
		LowerTimestamp               respjson.Field
		UpperTimestamp               respjson.Field
		IncludeObjectsWithNoValueSet respjson.Field
		OperationType                respjson.Field
		Operator                     respjson.Field
		TimePoint                    respjson.Field
		EndpointBehavior             respjson.Field
		PropertyParser               respjson.Field
		LowerBoundTimePoint          respjson.Field
		UpperBoundTimePoint          respjson.Field
		LowerBoundEndpointBehavior   respjson.Field
		UpperBoundEndpointBehavior   respjson.Field
		raw                          string
	} `json:"-"`
}

func (r *PublicNotAllFilterBranchFilterUnionPruningRefineBy) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PublicNotAllFilterBranchFilterUnionEventTypeID is an implicit subunion of
// [PublicNotAllFilterBranchFilterUnion].
// PublicNotAllFilterBranchFilterUnionEventTypeID provides convenient access to the
// sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [PublicNotAllFilterBranchFilterUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfInt OfString]
type PublicNotAllFilterBranchFilterUnionEventTypeID struct {
	// This field will be present if the value is a [int64] instead of an object.
	OfInt int64 `json:",inline"`
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	JSON     struct {
		OfInt    respjson.Field
		OfString respjson.Field
		raw      string
	} `json:"-"`
}

func (r *PublicNotAllFilterBranchFilterUnionEventTypeID) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties FilterBranches, FilterBranchOperator, FilterBranchType, Filters
// are required.
type PublicNotAllFilterBranchParam struct {
	FilterBranches []PublicNotAllFilterBranchFilterBranchUnionParam `json:"filterBranches,omitzero" api:"required"`
	// The operator used to combine filters within the branch (NOT_ALL).
	FilterBranchOperator string `json:"filterBranchOperator" api:"required"`
	// The type of the filter branch (NOT_ALL).
	//
	// Any of "NOT_ALL".
	FilterBranchType PublicNotAllFilterBranchFilterBranchType   `json:"filterBranchType,omitzero" api:"required"`
	Filters          []PublicNotAllFilterBranchFilterUnionParam `json:"filters,omitzero" api:"required"`
	paramObj
}

func (r PublicNotAllFilterBranchParam) MarshalJSON() (data []byte, err error) {
	type shadow PublicNotAllFilterBranchParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PublicNotAllFilterBranchParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type PublicNotAllFilterBranchFilterBranchUnionParam struct {
	OfOr                  *PublicOrFilterBranchParam                  `json:",omitzero,inline"`
	OfAnd                 *PublicAndFilterBranchParam                 `json:",omitzero,inline"`
	OfNotAll              *PublicNotAllFilterBranchParam              `json:",omitzero,inline"`
	OfNotAny              *PublicNotAnyFilterBranchParam              `json:",omitzero,inline"`
	OfRestricted          *PublicRestrictedFilterBranchParam          `json:",omitzero,inline"`
	OfUnifiedEvents       *PublicUnifiedEventsFilterBranchParam       `json:",omitzero,inline"`
	OfPropertyAssociation *PublicPropertyAssociationFilterBranchParam `json:",omitzero,inline"`
	OfAssociation         *PublicAssociationFilterBranchParam         `json:",omitzero,inline"`
	paramUnion
}

func (u PublicNotAllFilterBranchFilterBranchUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfOr,
		u.OfAnd,
		u.OfNotAll,
		u.OfNotAny,
		u.OfRestricted,
		u.OfUnifiedEvents,
		u.OfPropertyAssociation,
		u.OfAssociation)
}
func (u *PublicNotAllFilterBranchFilterBranchUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func init() {
	apijson.RegisterUnion[PublicNotAllFilterBranchFilterBranchUnionParam](
		"filterBranchType",
		apijson.Discriminator[PublicOrFilterBranchParam]("OR"),
		apijson.Discriminator[PublicAndFilterBranchParam]("AND"),
		apijson.Discriminator[PublicNotAllFilterBranchParam]("NOT_ALL"),
		apijson.Discriminator[PublicNotAnyFilterBranchParam]("NOT_ANY"),
		apijson.Discriminator[PublicRestrictedFilterBranchParam]("RESTRICTED"),
		apijson.Discriminator[PublicUnifiedEventsFilterBranchParam]("UNIFIED_EVENTS"),
		apijson.Discriminator[PublicPropertyAssociationFilterBranchParam]("PROPERTY_ASSOCIATION"),
		apijson.Discriminator[PublicAssociationFilterBranchParam]("ASSOCIATION"),
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type PublicNotAllFilterBranchFilterUnionParam struct {
	OfProperty                  *PublicPropertyFilterParam                  `json:",omitzero,inline"`
	OfAssociation               *PublicAssociationInListFilterParam         `json:",omitzero,inline"`
	OfPageView                  *PublicPageViewAnalyticsFilterParam         `json:",omitzero,inline"`
	OfCta                       *PublicCtaAnalyticsFilterParam              `json:",omitzero,inline"`
	OfEvent                     *PublicEventAnalyticsFilterParam            `json:",omitzero,inline"`
	OfFormSubmission            *PublicFormSubmissionFilterParam            `json:",omitzero,inline"`
	OfFormSubmissionOnPage      *PublicFormSubmissionOnPageFilterParam      `json:",omitzero,inline"`
	OfIntegrationEvent          *PublicIntegrationEventFilterParam          `json:",omitzero,inline"`
	OfEmailSubscription         *PublicEmailSubscriptionFilterParam         `json:",omitzero,inline"`
	OfCommunicationSubscription *PublicCommunicationSubscriptionFilterParam `json:",omitzero,inline"`
	OfCampaignInfluenced        *PublicCampaignInfluencedFilterParam        `json:",omitzero,inline"`
	OfSurveyMonkey              *PublicSurveyMonkeyFilterParam              `json:",omitzero,inline"`
	OfSurveyMonkeyValue         *PublicSurveyMonkeyValueFilterParam         `json:",omitzero,inline"`
	OfWebinar                   *PublicWebinarFilterParam                   `json:",omitzero,inline"`
	OfEmailEvent                *PublicEmailEventFilterParam                `json:",omitzero,inline"`
	OfPrivacy                   *PublicPrivacyAnalyticsFilterParam          `json:",omitzero,inline"`
	OfAdsSearch                 *PublicAdsSearchFilterParam                 `json:",omitzero,inline"`
	OfAdsTime                   *PublicAdsTimeFilterParam                   `json:",omitzero,inline"`
	OfInList                    *PublicInListFilterParam                    `json:",omitzero,inline"`
	OfNumAssociations           *PublicNumAssociationsFilterParam           `json:",omitzero,inline"`
	OfUnifiedEvents             *PublicUnifiedEventsFilterParam             `json:",omitzero,inline"`
	OfPropertyAssociation       *PublicPropertyAssociationInListFilterParam `json:",omitzero,inline"`
	OfConstant                  *PublicConstantFilterParam                  `json:",omitzero,inline"`
	paramUnion
}

func (u PublicNotAllFilterBranchFilterUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfProperty,
		u.OfAssociation,
		u.OfPageView,
		u.OfCta,
		u.OfEvent,
		u.OfFormSubmission,
		u.OfFormSubmissionOnPage,
		u.OfIntegrationEvent,
		u.OfEmailSubscription,
		u.OfCommunicationSubscription,
		u.OfCampaignInfluenced,
		u.OfSurveyMonkey,
		u.OfSurveyMonkeyValue,
		u.OfWebinar,
		u.OfEmailEvent,
		u.OfPrivacy,
		u.OfAdsSearch,
		u.OfAdsTime,
		u.OfInList,
		u.OfNumAssociations,
		u.OfUnifiedEvents,
		u.OfPropertyAssociation,
		u.OfConstant)
}
func (u *PublicNotAllFilterBranchFilterUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func init() {
	apijson.RegisterUnion[PublicNotAllFilterBranchFilterUnionParam](
		"filterType",
		apijson.Discriminator[PublicPropertyFilterParam]("PROPERTY"),
		apijson.Discriminator[PublicAssociationInListFilterParam]("ASSOCIATION"),
		apijson.Discriminator[PublicPageViewAnalyticsFilterParam]("PAGE_VIEW"),
		apijson.Discriminator[PublicCtaAnalyticsFilterParam]("CTA"),
		apijson.Discriminator[PublicEventAnalyticsFilterParam]("EVENT"),
		apijson.Discriminator[PublicFormSubmissionFilterParam]("FORM_SUBMISSION"),
		apijson.Discriminator[PublicFormSubmissionOnPageFilterParam]("FORM_SUBMISSION_ON_PAGE"),
		apijson.Discriminator[PublicIntegrationEventFilterParam]("INTEGRATION_EVENT"),
		apijson.Discriminator[PublicEmailSubscriptionFilterParam]("EMAIL_SUBSCRIPTION"),
		apijson.Discriminator[PublicCommunicationSubscriptionFilterParam]("COMMUNICATION_SUBSCRIPTION"),
		apijson.Discriminator[PublicCampaignInfluencedFilterParam]("CAMPAIGN_INFLUENCED"),
		apijson.Discriminator[PublicSurveyMonkeyFilterParam]("SURVEY_MONKEY"),
		apijson.Discriminator[PublicSurveyMonkeyValueFilterParam]("SURVEY_MONKEY_VALUE"),
		apijson.Discriminator[PublicWebinarFilterParam]("WEBINAR"),
		apijson.Discriminator[PublicEmailEventFilterParam]("EMAIL_EVENT"),
		apijson.Discriminator[PublicPrivacyAnalyticsFilterParam]("PRIVACY"),
		apijson.Discriminator[PublicAdsSearchFilterParam]("ADS_SEARCH"),
		apijson.Discriminator[PublicAdsTimeFilterParam]("ADS_TIME"),
		apijson.Discriminator[PublicInListFilterParam]("IN_LIST"),
		apijson.Discriminator[PublicNumAssociationsFilterParam]("NUM_ASSOCIATIONS"),
		apijson.Discriminator[PublicUnifiedEventsFilterParam]("UNIFIED_EVENTS"),
		apijson.Discriminator[PublicPropertyAssociationInListFilterParam]("PROPERTY_ASSOCIATION"),
		apijson.Discriminator[PublicConstantFilterParam]("CONSTANT"),
	)
}

type PublicNotAnyFilterBranch struct {
	FilterBranches []PublicNotAnyFilterBranchFilterBranchUnion `json:"filterBranches" api:"required"`
	// Specifies the logical operator used to combine filters within the branch
	// (NOT_ANY).
	FilterBranchOperator string `json:"filterBranchOperator" api:"required"`
	// Indicates the type of filter branch (NOT_ANY).
	//
	// Any of "NOT_ANY".
	FilterBranchType PublicNotAnyFilterBranchFilterBranchType `json:"filterBranchType" api:"required"`
	Filters          []PublicNotAnyFilterBranchFilterUnion    `json:"filters" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FilterBranches       respjson.Field
		FilterBranchOperator respjson.Field
		FilterBranchType     respjson.Field
		Filters              respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PublicNotAnyFilterBranch) RawJSON() string { return r.JSON.raw }
func (r *PublicNotAnyFilterBranch) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this PublicNotAnyFilterBranch to a
// PublicNotAnyFilterBranchParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// PublicNotAnyFilterBranchParam.Overrides()
func (r PublicNotAnyFilterBranch) ToParam() PublicNotAnyFilterBranchParam {
	return param.Override[PublicNotAnyFilterBranchParam](json.RawMessage(r.RawJSON()))
}

// PublicNotAnyFilterBranchFilterBranchUnion contains all possible properties and
// values from [PublicOrFilterBranch], [PublicAndFilterBranch],
// [PublicNotAllFilterBranch], [PublicNotAnyFilterBranch],
// [PublicRestrictedFilterBranch], [PublicUnifiedEventsFilterBranch],
// [PublicPropertyAssociationFilterBranch], [PublicAssociationFilterBranch].
//
// Use the [PublicNotAnyFilterBranchFilterBranchUnion.AsAny] method to switch on
// the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type PublicNotAnyFilterBranchFilterBranchUnion struct {
	// This field is a union of [[]PublicOrFilterBranchFilterBranchUnion],
	// [[]PublicAndFilterBranchFilterBranchUnion],
	// [[]PublicNotAllFilterBranchFilterBranchUnion],
	// [[]PublicNotAnyFilterBranchFilterBranchUnion],
	// [[]PublicRestrictedFilterBranchFilterBranchUnion],
	// [[]PublicUnifiedEventsFilterBranchFilterBranchUnion],
	// [[]PublicPropertyAssociationFilterBranchFilterBranchUnion],
	// [[]PublicAssociationFilterBranchFilterBranchUnion]
	FilterBranches       PublicNotAnyFilterBranchFilterBranchUnionFilterBranches `json:"filterBranches"`
	FilterBranchOperator string                                                  `json:"filterBranchOperator"`
	// Any of "OR", "AND", "NOT_ALL", "NOT_ANY", "RESTRICTED", "UNIFIED_EVENTS",
	// "PROPERTY_ASSOCIATION", "ASSOCIATION".
	FilterBranchType string `json:"filterBranchType"`
	// This field is a union of [[]PublicOrFilterBranchFilterUnion],
	// [[]PublicAndFilterBranchFilterUnion], [[]PublicNotAllFilterBranchFilterUnion],
	// [[]PublicNotAnyFilterBranchFilterUnion],
	// [[]PublicRestrictedFilterBranchFilterUnion],
	// [[]PublicUnifiedEventsFilterBranchFilterUnion],
	// [[]PublicPropertyAssociationFilterBranchFilterUnion],
	// [[]PublicAssociationFilterBranchFilterUnion]
	Filters PublicNotAnyFilterBranchFilterBranchUnionFilters `json:"filters"`
	// This field is from variant [PublicUnifiedEventsFilterBranch].
	EventTypeID string `json:"eventTypeId"`
	Operator    string `json:"operator"`
	// This field is from variant [PublicUnifiedEventsFilterBranch].
	CoalescingRefineBy PublicUnifiedEventsFilterBranchCoalescingRefineByUnion `json:"coalescingRefineBy"`
	// This field is from variant [PublicUnifiedEventsFilterBranch].
	PruningRefineBy PublicUnifiedEventsFilterBranchPruningRefineByUnion `json:"pruningRefineBy"`
	ObjectTypeID    string                                              `json:"objectTypeId"`
	// This field is from variant [PublicPropertyAssociationFilterBranch].
	PropertyWithObjectID string `json:"propertyWithObjectId"`
	// This field is from variant [PublicAssociationFilterBranch].
	AssociationCategory string `json:"associationCategory"`
	// This field is from variant [PublicAssociationFilterBranch].
	AssociationTypeID int64 `json:"associationTypeId"`
	JSON              struct {
		FilterBranches       respjson.Field
		FilterBranchOperator respjson.Field
		FilterBranchType     respjson.Field
		Filters              respjson.Field
		EventTypeID          respjson.Field
		Operator             respjson.Field
		CoalescingRefineBy   respjson.Field
		PruningRefineBy      respjson.Field
		ObjectTypeID         respjson.Field
		PropertyWithObjectID respjson.Field
		AssociationCategory  respjson.Field
		AssociationTypeID    respjson.Field
		raw                  string
	} `json:"-"`
}

// anyPublicNotAnyFilterBranchFilterBranch is implemented by each variant of
// [PublicNotAnyFilterBranchFilterBranchUnion] to add type safety for the return
// type of [PublicNotAnyFilterBranchFilterBranchUnion.AsAny]
type anyPublicNotAnyFilterBranchFilterBranch interface {
	implPublicNotAnyFilterBranchFilterBranchUnion()
}

func (PublicOrFilterBranch) implPublicNotAnyFilterBranchFilterBranchUnion()                  {}
func (PublicAndFilterBranch) implPublicNotAnyFilterBranchFilterBranchUnion()                 {}
func (PublicNotAllFilterBranch) implPublicNotAnyFilterBranchFilterBranchUnion()              {}
func (PublicNotAnyFilterBranch) implPublicNotAnyFilterBranchFilterBranchUnion()              {}
func (PublicRestrictedFilterBranch) implPublicNotAnyFilterBranchFilterBranchUnion()          {}
func (PublicUnifiedEventsFilterBranch) implPublicNotAnyFilterBranchFilterBranchUnion()       {}
func (PublicPropertyAssociationFilterBranch) implPublicNotAnyFilterBranchFilterBranchUnion() {}
func (PublicAssociationFilterBranch) implPublicNotAnyFilterBranchFilterBranchUnion()         {}

// Use the following switch statement to find the correct variant
//
//	switch variant := PublicNotAnyFilterBranchFilterBranchUnion.AsAny().(type) {
//	case crm.PublicOrFilterBranch:
//	case crm.PublicAndFilterBranch:
//	case crm.PublicNotAllFilterBranch:
//	case crm.PublicNotAnyFilterBranch:
//	case crm.PublicRestrictedFilterBranch:
//	case crm.PublicUnifiedEventsFilterBranch:
//	case crm.PublicPropertyAssociationFilterBranch:
//	case crm.PublicAssociationFilterBranch:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u PublicNotAnyFilterBranchFilterBranchUnion) AsAny() anyPublicNotAnyFilterBranchFilterBranch {
	switch u.FilterBranchType {
	case "OR":
		return u.AsOr()
	case "AND":
		return u.AsAnd()
	case "NOT_ALL":
		return u.AsNotAll()
	case "NOT_ANY":
		return u.AsNotAny()
	case "RESTRICTED":
		return u.AsRestricted()
	case "UNIFIED_EVENTS":
		return u.AsUnifiedEvents()
	case "PROPERTY_ASSOCIATION":
		return u.AsPropertyAssociation()
	case "ASSOCIATION":
		return u.AsAssociation()
	}
	return nil
}

func (u PublicNotAnyFilterBranchFilterBranchUnion) AsOr() (v PublicOrFilterBranch) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicNotAnyFilterBranchFilterBranchUnion) AsAnd() (v PublicAndFilterBranch) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicNotAnyFilterBranchFilterBranchUnion) AsNotAll() (v PublicNotAllFilterBranch) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicNotAnyFilterBranchFilterBranchUnion) AsNotAny() (v PublicNotAnyFilterBranch) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicNotAnyFilterBranchFilterBranchUnion) AsRestricted() (v PublicRestrictedFilterBranch) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicNotAnyFilterBranchFilterBranchUnion) AsUnifiedEvents() (v PublicUnifiedEventsFilterBranch) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicNotAnyFilterBranchFilterBranchUnion) AsPropertyAssociation() (v PublicPropertyAssociationFilterBranch) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicNotAnyFilterBranchFilterBranchUnion) AsAssociation() (v PublicAssociationFilterBranch) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u PublicNotAnyFilterBranchFilterBranchUnion) RawJSON() string { return u.JSON.raw }

func (r *PublicNotAnyFilterBranchFilterBranchUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PublicNotAnyFilterBranchFilterBranchUnionFilterBranches is an implicit subunion
// of [PublicNotAnyFilterBranchFilterBranchUnion].
// PublicNotAnyFilterBranchFilterBranchUnionFilterBranches provides convenient
// access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [PublicNotAnyFilterBranchFilterBranchUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfPublicOrFilterBranchFilterBranches
// OfPublicAndFilterBranchFilterBranches OfPublicNotAllFilterBranchFilterBranches
// OfPublicNotAnyFilterBranchFilterBranches
// OfPublicRestrictedFilterBranchFilterBranches
// OfPublicUnifiedEventsFilterBranchFilterBranches
// OfPublicPropertyAssociationFilterBranchFilterBranches
// OfPublicAssociationFilterBranchFilterBranches]
type PublicNotAnyFilterBranchFilterBranchUnionFilterBranches struct {
	// This field will be present if the value is a
	// [[]PublicOrFilterBranchFilterBranchUnion] instead of an object.
	OfPublicOrFilterBranchFilterBranches []PublicOrFilterBranchFilterBranchUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]PublicAndFilterBranchFilterBranchUnion] instead of an object.
	OfPublicAndFilterBranchFilterBranches []PublicAndFilterBranchFilterBranchUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]PublicNotAllFilterBranchFilterBranchUnion] instead of an object.
	OfPublicNotAllFilterBranchFilterBranches []PublicNotAllFilterBranchFilterBranchUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]PublicNotAnyFilterBranchFilterBranchUnion] instead of an object.
	OfPublicNotAnyFilterBranchFilterBranches []PublicNotAnyFilterBranchFilterBranchUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]PublicRestrictedFilterBranchFilterBranchUnion] instead of an object.
	OfPublicRestrictedFilterBranchFilterBranches []PublicRestrictedFilterBranchFilterBranchUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]PublicUnifiedEventsFilterBranchFilterBranchUnion] instead of an object.
	OfPublicUnifiedEventsFilterBranchFilterBranches []PublicUnifiedEventsFilterBranchFilterBranchUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]PublicPropertyAssociationFilterBranchFilterBranchUnion] instead of an object.
	OfPublicPropertyAssociationFilterBranchFilterBranches []PublicPropertyAssociationFilterBranchFilterBranchUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]PublicAssociationFilterBranchFilterBranchUnion] instead of an object.
	OfPublicAssociationFilterBranchFilterBranches []PublicAssociationFilterBranchFilterBranchUnion `json:",inline"`
	JSON                                          struct {
		OfPublicOrFilterBranchFilterBranches                  respjson.Field
		OfPublicAndFilterBranchFilterBranches                 respjson.Field
		OfPublicNotAllFilterBranchFilterBranches              respjson.Field
		OfPublicNotAnyFilterBranchFilterBranches              respjson.Field
		OfPublicRestrictedFilterBranchFilterBranches          respjson.Field
		OfPublicUnifiedEventsFilterBranchFilterBranches       respjson.Field
		OfPublicPropertyAssociationFilterBranchFilterBranches respjson.Field
		OfPublicAssociationFilterBranchFilterBranches         respjson.Field
		raw                                                   string
	} `json:"-"`
}

func (r *PublicNotAnyFilterBranchFilterBranchUnionFilterBranches) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PublicNotAnyFilterBranchFilterBranchUnionFilters is an implicit subunion of
// [PublicNotAnyFilterBranchFilterBranchUnion].
// PublicNotAnyFilterBranchFilterBranchUnionFilters provides convenient access to
// the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [PublicNotAnyFilterBranchFilterBranchUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfPublicOrFilterBranchFilters OfPublicAndFilterBranchFilters
// OfPublicNotAllFilterBranchFilters OfPublicNotAnyFilterBranchFilters
// OfPublicRestrictedFilterBranchFilters OfPublicUnifiedEventsFilterBranchFilters
// OfPublicPropertyAssociationFilterBranchFilters
// OfPublicAssociationFilterBranchFilters]
type PublicNotAnyFilterBranchFilterBranchUnionFilters struct {
	// This field will be present if the value is a [[]PublicOrFilterBranchFilterUnion]
	// instead of an object.
	OfPublicOrFilterBranchFilters []PublicOrFilterBranchFilterUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]PublicAndFilterBranchFilterUnion] instead of an object.
	OfPublicAndFilterBranchFilters []PublicAndFilterBranchFilterUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]PublicNotAllFilterBranchFilterUnion] instead of an object.
	OfPublicNotAllFilterBranchFilters []PublicNotAllFilterBranchFilterUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]PublicNotAnyFilterBranchFilterUnion] instead of an object.
	OfPublicNotAnyFilterBranchFilters []PublicNotAnyFilterBranchFilterUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]PublicRestrictedFilterBranchFilterUnion] instead of an object.
	OfPublicRestrictedFilterBranchFilters []PublicRestrictedFilterBranchFilterUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]PublicUnifiedEventsFilterBranchFilterUnion] instead of an object.
	OfPublicUnifiedEventsFilterBranchFilters []PublicUnifiedEventsFilterBranchFilterUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]PublicPropertyAssociationFilterBranchFilterUnion] instead of an object.
	OfPublicPropertyAssociationFilterBranchFilters []PublicPropertyAssociationFilterBranchFilterUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]PublicAssociationFilterBranchFilterUnion] instead of an object.
	OfPublicAssociationFilterBranchFilters []PublicAssociationFilterBranchFilterUnion `json:",inline"`
	JSON                                   struct {
		OfPublicOrFilterBranchFilters                  respjson.Field
		OfPublicAndFilterBranchFilters                 respjson.Field
		OfPublicNotAllFilterBranchFilters              respjson.Field
		OfPublicNotAnyFilterBranchFilters              respjson.Field
		OfPublicRestrictedFilterBranchFilters          respjson.Field
		OfPublicUnifiedEventsFilterBranchFilters       respjson.Field
		OfPublicPropertyAssociationFilterBranchFilters respjson.Field
		OfPublicAssociationFilterBranchFilters         respjson.Field
		raw                                            string
	} `json:"-"`
}

func (r *PublicNotAnyFilterBranchFilterBranchUnionFilters) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Indicates the type of filter branch (NOT_ANY).
type PublicNotAnyFilterBranchFilterBranchType string

const (
	PublicNotAnyFilterBranchFilterBranchTypeNotAny PublicNotAnyFilterBranchFilterBranchType = "NOT_ANY"
)

// PublicNotAnyFilterBranchFilterUnion contains all possible properties and values
// from [PublicPropertyFilter], [PublicAssociationInListFilter],
// [PublicPageViewAnalyticsFilter], [PublicCtaAnalyticsFilter],
// [PublicEventAnalyticsFilter], [PublicFormSubmissionFilter],
// [PublicFormSubmissionOnPageFilter], [PublicIntegrationEventFilter],
// [PublicEmailSubscriptionFilter], [PublicCommunicationSubscriptionFilter],
// [PublicCampaignInfluencedFilter], [PublicSurveyMonkeyFilter],
// [PublicSurveyMonkeyValueFilter], [PublicWebinarFilter],
// [PublicEmailEventFilter], [PublicPrivacyAnalyticsFilter],
// [PublicAdsSearchFilter], [PublicAdsTimeFilter], [PublicInListFilter],
// [PublicNumAssociationsFilter], [PublicUnifiedEventsFilter],
// [PublicPropertyAssociationInListFilter], [PublicConstantFilter].
//
// Use the [PublicNotAnyFilterBranchFilterUnion.AsAny] method to switch on the
// variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type PublicNotAnyFilterBranchFilterUnion struct {
	// Any of "PROPERTY", "ASSOCIATION", "PAGE_VIEW", "CTA", "EVENT",
	// "FORM_SUBMISSION", "FORM_SUBMISSION_ON_PAGE", "INTEGRATION_EVENT",
	// "EMAIL_SUBSCRIPTION", "COMMUNICATION_SUBSCRIPTION", "CAMPAIGN_INFLUENCED",
	// "SURVEY_MONKEY", "SURVEY_MONKEY_VALUE", "WEBINAR", "EMAIL_EVENT", "PRIVACY",
	// "ADS_SEARCH", "ADS_TIME", "IN_LIST", "NUM_ASSOCIATIONS", "UNIFIED_EVENTS",
	// "PROPERTY_ASSOCIATION", "CONSTANT".
	FilterType string `json:"filterType"`
	// This field is from variant [PublicPropertyFilter].
	Operation PublicPropertyFilterOperationUnion `json:"operation"`
	// This field is from variant [PublicPropertyFilter].
	Property            string `json:"property"`
	AssociationCategory string `json:"associationCategory"`
	AssociationTypeID   int64  `json:"associationTypeId"`
	// This field is a union of [PublicAssociationInListFilterCoalescingRefineByUnion],
	// [PublicPageViewAnalyticsFilterCoalescingRefineByUnion],
	// [PublicCtaAnalyticsFilterCoalescingRefineByUnion],
	// [PublicEventAnalyticsFilterCoalescingRefineByUnion],
	// [PublicFormSubmissionFilterCoalescingRefineByUnion],
	// [PublicFormSubmissionOnPageFilterCoalescingRefineByUnion],
	// [PublicNumAssociationsFilterCoalescingRefineByUnion],
	// [PublicUnifiedEventsFilterCoalescingRefineByUnion],
	// [PublicPropertyAssociationInListFilterCoalescingRefineByUnion]
	CoalescingRefineBy PublicNotAnyFilterBranchFilterUnionCoalescingRefineBy `json:"coalescingRefineBy"`
	ListID             string                                                `json:"listId"`
	Operator           string                                                `json:"operator"`
	// This field is from variant [PublicAssociationInListFilter].
	ToObjectType   string `json:"toObjectType"`
	ToObjectTypeID string `json:"toObjectTypeId"`
	// This field is from variant [PublicPageViewAnalyticsFilter].
	PageURL string `json:"pageUrl"`
	// This field is from variant [PublicPageViewAnalyticsFilter].
	EnableTracking bool `json:"enableTracking"`
	// This field is a union of [PublicPageViewAnalyticsFilterPruningRefineByUnion],
	// [PublicCtaAnalyticsFilterPruningRefineByUnion],
	// [PublicEventAnalyticsFilterPruningRefineByUnion],
	// [PublicFormSubmissionFilterPruningRefineByUnion],
	// [PublicFormSubmissionOnPageFilterPruningRefineByUnion],
	// [PublicEmailEventFilterPruningRefineByUnion],
	// [PublicAdsTimeFilterPruningRefineByUnion],
	// [PublicUnifiedEventsFilterPruningRefineByUnion]
	PruningRefineBy PublicNotAnyFilterBranchFilterUnionPruningRefineBy `json:"pruningRefineBy"`
	// This field is from variant [PublicCtaAnalyticsFilter].
	CtaName string `json:"ctaName"`
	// This field is from variant [PublicEventAnalyticsFilter].
	EventID string `json:"eventId"`
	FormID  string `json:"formId"`
	// This field is from variant [PublicFormSubmissionOnPageFilter].
	PageID string `json:"pageId"`
	// This field is a union of [int64], [string]
	EventTypeID PublicNotAnyFilterBranchFilterUnionEventTypeID `json:"eventTypeId"`
	FilterLines []PublicEventFilterMetadata                    `json:"filterLines"`
	// This field is from variant [PublicEmailSubscriptionFilter].
	AcceptedStatuses []string `json:"acceptedStatuses"`
	SubscriptionIDs  []string `json:"subscriptionIds"`
	SubscriptionType string   `json:"subscriptionType"`
	// This field is from variant [PublicCommunicationSubscriptionFilter].
	AcceptedOptStates []string `json:"acceptedOptStates"`
	// This field is from variant [PublicCommunicationSubscriptionFilter].
	Channel string `json:"channel"`
	// This field is from variant [PublicCommunicationSubscriptionFilter].
	BusinessUnitID string `json:"businessUnitId"`
	// This field is from variant [PublicCampaignInfluencedFilter].
	CampaignID string `json:"campaignId"`
	SurveyID   string `json:"surveyId"`
	// This field is from variant [PublicSurveyMonkeyValueFilter].
	SurveyQuestion string `json:"surveyQuestion"`
	// This field is from variant [PublicSurveyMonkeyValueFilter].
	ValueComparison PublicSurveyMonkeyValueFilterValueComparisonUnion `json:"valueComparison"`
	// This field is from variant [PublicSurveyMonkeyValueFilter].
	SurveyAnswerColID string `json:"surveyAnswerColId"`
	// This field is from variant [PublicSurveyMonkeyValueFilter].
	SurveyAnswerRowID string `json:"surveyAnswerRowId"`
	// This field is from variant [PublicWebinarFilter].
	WebinarID string `json:"webinarId"`
	// This field is from variant [PublicEmailEventFilter].
	AppID string `json:"appId"`
	// This field is from variant [PublicEmailEventFilter].
	EmailID string `json:"emailId"`
	// This field is from variant [PublicEmailEventFilter].
	Level string `json:"level"`
	// This field is from variant [PublicEmailEventFilter].
	ClickURL string `json:"clickUrl"`
	// This field is from variant [PublicPrivacyAnalyticsFilter].
	PrivacyName string `json:"privacyName"`
	// This field is from variant [PublicAdsSearchFilter].
	AdNetwork string `json:"adNetwork"`
	// This field is from variant [PublicAdsSearchFilter].
	EntityType string `json:"entityType"`
	// This field is from variant [PublicAdsSearchFilter].
	SearchTerms []string `json:"searchTerms"`
	// This field is from variant [PublicAdsSearchFilter].
	SearchTermType string `json:"searchTermType"`
	// This field is from variant [PublicInListFilter].
	Metadata PublicInListFilterMetadata `json:"metadata"`
	// This field is from variant [PublicPropertyAssociationInListFilter].
	PropertyWithObjectID string `json:"propertyWithObjectId"`
	// This field is from variant [PublicConstantFilter].
	ShouldAccept bool `json:"shouldAccept"`
	// This field is from variant [PublicConstantFilter].
	Source string `json:"source"`
	JSON   struct {
		FilterType           respjson.Field
		Operation            respjson.Field
		Property             respjson.Field
		AssociationCategory  respjson.Field
		AssociationTypeID    respjson.Field
		CoalescingRefineBy   respjson.Field
		ListID               respjson.Field
		Operator             respjson.Field
		ToObjectType         respjson.Field
		ToObjectTypeID       respjson.Field
		PageURL              respjson.Field
		EnableTracking       respjson.Field
		PruningRefineBy      respjson.Field
		CtaName              respjson.Field
		EventID              respjson.Field
		FormID               respjson.Field
		PageID               respjson.Field
		EventTypeID          respjson.Field
		FilterLines          respjson.Field
		AcceptedStatuses     respjson.Field
		SubscriptionIDs      respjson.Field
		SubscriptionType     respjson.Field
		AcceptedOptStates    respjson.Field
		Channel              respjson.Field
		BusinessUnitID       respjson.Field
		CampaignID           respjson.Field
		SurveyID             respjson.Field
		SurveyQuestion       respjson.Field
		ValueComparison      respjson.Field
		SurveyAnswerColID    respjson.Field
		SurveyAnswerRowID    respjson.Field
		WebinarID            respjson.Field
		AppID                respjson.Field
		EmailID              respjson.Field
		Level                respjson.Field
		ClickURL             respjson.Field
		PrivacyName          respjson.Field
		AdNetwork            respjson.Field
		EntityType           respjson.Field
		SearchTerms          respjson.Field
		SearchTermType       respjson.Field
		Metadata             respjson.Field
		PropertyWithObjectID respjson.Field
		ShouldAccept         respjson.Field
		Source               respjson.Field
		raw                  string
	} `json:"-"`
}

// anyPublicNotAnyFilterBranchFilter is implemented by each variant of
// [PublicNotAnyFilterBranchFilterUnion] to add type safety for the return type of
// [PublicNotAnyFilterBranchFilterUnion.AsAny]
type anyPublicNotAnyFilterBranchFilter interface {
	implPublicNotAnyFilterBranchFilterUnion()
}

func (PublicPropertyFilter) implPublicNotAnyFilterBranchFilterUnion()                  {}
func (PublicAssociationInListFilter) implPublicNotAnyFilterBranchFilterUnion()         {}
func (PublicPageViewAnalyticsFilter) implPublicNotAnyFilterBranchFilterUnion()         {}
func (PublicCtaAnalyticsFilter) implPublicNotAnyFilterBranchFilterUnion()              {}
func (PublicEventAnalyticsFilter) implPublicNotAnyFilterBranchFilterUnion()            {}
func (PublicFormSubmissionFilter) implPublicNotAnyFilterBranchFilterUnion()            {}
func (PublicFormSubmissionOnPageFilter) implPublicNotAnyFilterBranchFilterUnion()      {}
func (PublicIntegrationEventFilter) implPublicNotAnyFilterBranchFilterUnion()          {}
func (PublicEmailSubscriptionFilter) implPublicNotAnyFilterBranchFilterUnion()         {}
func (PublicCommunicationSubscriptionFilter) implPublicNotAnyFilterBranchFilterUnion() {}
func (PublicCampaignInfluencedFilter) implPublicNotAnyFilterBranchFilterUnion()        {}
func (PublicSurveyMonkeyFilter) implPublicNotAnyFilterBranchFilterUnion()              {}
func (PublicSurveyMonkeyValueFilter) implPublicNotAnyFilterBranchFilterUnion()         {}
func (PublicWebinarFilter) implPublicNotAnyFilterBranchFilterUnion()                   {}
func (PublicEmailEventFilter) implPublicNotAnyFilterBranchFilterUnion()                {}
func (PublicPrivacyAnalyticsFilter) implPublicNotAnyFilterBranchFilterUnion()          {}
func (PublicAdsSearchFilter) implPublicNotAnyFilterBranchFilterUnion()                 {}
func (PublicAdsTimeFilter) implPublicNotAnyFilterBranchFilterUnion()                   {}
func (PublicInListFilter) implPublicNotAnyFilterBranchFilterUnion()                    {}
func (PublicNumAssociationsFilter) implPublicNotAnyFilterBranchFilterUnion()           {}
func (PublicUnifiedEventsFilter) implPublicNotAnyFilterBranchFilterUnion()             {}
func (PublicPropertyAssociationInListFilter) implPublicNotAnyFilterBranchFilterUnion() {}
func (PublicConstantFilter) implPublicNotAnyFilterBranchFilterUnion()                  {}

// Use the following switch statement to find the correct variant
//
//	switch variant := PublicNotAnyFilterBranchFilterUnion.AsAny().(type) {
//	case crm.PublicPropertyFilter:
//	case crm.PublicAssociationInListFilter:
//	case crm.PublicPageViewAnalyticsFilter:
//	case crm.PublicCtaAnalyticsFilter:
//	case crm.PublicEventAnalyticsFilter:
//	case crm.PublicFormSubmissionFilter:
//	case crm.PublicFormSubmissionOnPageFilter:
//	case crm.PublicIntegrationEventFilter:
//	case crm.PublicEmailSubscriptionFilter:
//	case crm.PublicCommunicationSubscriptionFilter:
//	case crm.PublicCampaignInfluencedFilter:
//	case crm.PublicSurveyMonkeyFilter:
//	case crm.PublicSurveyMonkeyValueFilter:
//	case crm.PublicWebinarFilter:
//	case crm.PublicEmailEventFilter:
//	case crm.PublicPrivacyAnalyticsFilter:
//	case crm.PublicAdsSearchFilter:
//	case crm.PublicAdsTimeFilter:
//	case crm.PublicInListFilter:
//	case crm.PublicNumAssociationsFilter:
//	case crm.PublicUnifiedEventsFilter:
//	case crm.PublicPropertyAssociationInListFilter:
//	case crm.PublicConstantFilter:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u PublicNotAnyFilterBranchFilterUnion) AsAny() anyPublicNotAnyFilterBranchFilter {
	switch u.FilterType {
	case "PROPERTY":
		return u.AsProperty()
	case "ASSOCIATION":
		return u.AsAssociation()
	case "PAGE_VIEW":
		return u.AsPageView()
	case "CTA":
		return u.AsCta()
	case "EVENT":
		return u.AsEvent()
	case "FORM_SUBMISSION":
		return u.AsFormSubmission()
	case "FORM_SUBMISSION_ON_PAGE":
		return u.AsFormSubmissionOnPage()
	case "INTEGRATION_EVENT":
		return u.AsIntegrationEvent()
	case "EMAIL_SUBSCRIPTION":
		return u.AsEmailSubscription()
	case "COMMUNICATION_SUBSCRIPTION":
		return u.AsCommunicationSubscription()
	case "CAMPAIGN_INFLUENCED":
		return u.AsCampaignInfluenced()
	case "SURVEY_MONKEY":
		return u.AsSurveyMonkey()
	case "SURVEY_MONKEY_VALUE":
		return u.AsSurveyMonkeyValue()
	case "WEBINAR":
		return u.AsWebinar()
	case "EMAIL_EVENT":
		return u.AsEmailEvent()
	case "PRIVACY":
		return u.AsPrivacy()
	case "ADS_SEARCH":
		return u.AsAdsSearch()
	case "ADS_TIME":
		return u.AsAdsTime()
	case "IN_LIST":
		return u.AsInList()
	case "NUM_ASSOCIATIONS":
		return u.AsNumAssociations()
	case "UNIFIED_EVENTS":
		return u.AsUnifiedEvents()
	case "PROPERTY_ASSOCIATION":
		return u.AsPropertyAssociation()
	case "CONSTANT":
		return u.AsConstant()
	}
	return nil
}

func (u PublicNotAnyFilterBranchFilterUnion) AsProperty() (v PublicPropertyFilter) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicNotAnyFilterBranchFilterUnion) AsAssociation() (v PublicAssociationInListFilter) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicNotAnyFilterBranchFilterUnion) AsPageView() (v PublicPageViewAnalyticsFilter) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicNotAnyFilterBranchFilterUnion) AsCta() (v PublicCtaAnalyticsFilter) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicNotAnyFilterBranchFilterUnion) AsEvent() (v PublicEventAnalyticsFilter) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicNotAnyFilterBranchFilterUnion) AsFormSubmission() (v PublicFormSubmissionFilter) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicNotAnyFilterBranchFilterUnion) AsFormSubmissionOnPage() (v PublicFormSubmissionOnPageFilter) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicNotAnyFilterBranchFilterUnion) AsIntegrationEvent() (v PublicIntegrationEventFilter) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicNotAnyFilterBranchFilterUnion) AsEmailSubscription() (v PublicEmailSubscriptionFilter) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicNotAnyFilterBranchFilterUnion) AsCommunicationSubscription() (v PublicCommunicationSubscriptionFilter) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicNotAnyFilterBranchFilterUnion) AsCampaignInfluenced() (v PublicCampaignInfluencedFilter) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicNotAnyFilterBranchFilterUnion) AsSurveyMonkey() (v PublicSurveyMonkeyFilter) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicNotAnyFilterBranchFilterUnion) AsSurveyMonkeyValue() (v PublicSurveyMonkeyValueFilter) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicNotAnyFilterBranchFilterUnion) AsWebinar() (v PublicWebinarFilter) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicNotAnyFilterBranchFilterUnion) AsEmailEvent() (v PublicEmailEventFilter) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicNotAnyFilterBranchFilterUnion) AsPrivacy() (v PublicPrivacyAnalyticsFilter) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicNotAnyFilterBranchFilterUnion) AsAdsSearch() (v PublicAdsSearchFilter) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicNotAnyFilterBranchFilterUnion) AsAdsTime() (v PublicAdsTimeFilter) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicNotAnyFilterBranchFilterUnion) AsInList() (v PublicInListFilter) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicNotAnyFilterBranchFilterUnion) AsNumAssociations() (v PublicNumAssociationsFilter) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicNotAnyFilterBranchFilterUnion) AsUnifiedEvents() (v PublicUnifiedEventsFilter) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicNotAnyFilterBranchFilterUnion) AsPropertyAssociation() (v PublicPropertyAssociationInListFilter) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicNotAnyFilterBranchFilterUnion) AsConstant() (v PublicConstantFilter) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u PublicNotAnyFilterBranchFilterUnion) RawJSON() string { return u.JSON.raw }

func (r *PublicNotAnyFilterBranchFilterUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PublicNotAnyFilterBranchFilterUnionCoalescingRefineBy is an implicit subunion of
// [PublicNotAnyFilterBranchFilterUnion].
// PublicNotAnyFilterBranchFilterUnionCoalescingRefineBy provides convenient access
// to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [PublicNotAnyFilterBranchFilterUnion].
type PublicNotAnyFilterBranchFilterUnionCoalescingRefineBy struct {
	Type string `json:"type"`
	// This field is from variant
	// [PublicAssociationInListFilterCoalescingRefineByUnion],
	// [PublicPageViewAnalyticsFilterCoalescingRefineByUnion],
	// [PublicCtaAnalyticsFilterCoalescingRefineByUnion],
	// [PublicEventAnalyticsFilterCoalescingRefineByUnion],
	// [PublicFormSubmissionFilterCoalescingRefineByUnion],
	// [PublicFormSubmissionOnPageFilterCoalescingRefineByUnion],
	// [PublicNumAssociationsFilterCoalescingRefineByUnion],
	// [PublicUnifiedEventsFilterCoalescingRefineByUnion],
	// [PublicPropertyAssociationInListFilterCoalescingRefineByUnion].
	MaxOccurrences int64 `json:"maxOccurrences"`
	// This field is from variant
	// [PublicAssociationInListFilterCoalescingRefineByUnion],
	// [PublicPageViewAnalyticsFilterCoalescingRefineByUnion],
	// [PublicCtaAnalyticsFilterCoalescingRefineByUnion],
	// [PublicEventAnalyticsFilterCoalescingRefineByUnion],
	// [PublicFormSubmissionFilterCoalescingRefineByUnion],
	// [PublicFormSubmissionOnPageFilterCoalescingRefineByUnion],
	// [PublicNumAssociationsFilterCoalescingRefineByUnion],
	// [PublicUnifiedEventsFilterCoalescingRefineByUnion],
	// [PublicPropertyAssociationInListFilterCoalescingRefineByUnion].
	MinOccurrences int64 `json:"minOccurrences"`
	// This field is from variant
	// [PublicAssociationInListFilterCoalescingRefineByUnion],
	// [PublicPageViewAnalyticsFilterCoalescingRefineByUnion],
	// [PublicCtaAnalyticsFilterCoalescingRefineByUnion],
	// [PublicEventAnalyticsFilterCoalescingRefineByUnion],
	// [PublicFormSubmissionFilterCoalescingRefineByUnion],
	// [PublicFormSubmissionOnPageFilterCoalescingRefineByUnion],
	// [PublicNumAssociationsFilterCoalescingRefineByUnion],
	// [PublicUnifiedEventsFilterCoalescingRefineByUnion],
	// [PublicPropertyAssociationInListFilterCoalescingRefineByUnion].
	SetType    string `json:"setType"`
	Comparison string `json:"comparison"`
	// This field is from variant
	// [PublicAssociationInListFilterCoalescingRefineByUnion],
	// [PublicPageViewAnalyticsFilterCoalescingRefineByUnion],
	// [PublicCtaAnalyticsFilterCoalescingRefineByUnion],
	// [PublicEventAnalyticsFilterCoalescingRefineByUnion],
	// [PublicFormSubmissionFilterCoalescingRefineByUnion],
	// [PublicFormSubmissionOnPageFilterCoalescingRefineByUnion],
	// [PublicNumAssociationsFilterCoalescingRefineByUnion],
	// [PublicUnifiedEventsFilterCoalescingRefineByUnion],
	// [PublicPropertyAssociationInListFilterCoalescingRefineByUnion].
	TimeOffset PublicTimeOffset `json:"timeOffset"`
	// This field is from variant
	// [PublicAssociationInListFilterCoalescingRefineByUnion],
	// [PublicPageViewAnalyticsFilterCoalescingRefineByUnion],
	// [PublicCtaAnalyticsFilterCoalescingRefineByUnion],
	// [PublicEventAnalyticsFilterCoalescingRefineByUnion],
	// [PublicFormSubmissionFilterCoalescingRefineByUnion],
	// [PublicFormSubmissionOnPageFilterCoalescingRefineByUnion],
	// [PublicNumAssociationsFilterCoalescingRefineByUnion],
	// [PublicUnifiedEventsFilterCoalescingRefineByUnion],
	// [PublicPropertyAssociationInListFilterCoalescingRefineByUnion].
	LowerBoundOffset PublicTimeOffset `json:"lowerBoundOffset"`
	RangeType        string           `json:"rangeType"`
	// This field is from variant
	// [PublicAssociationInListFilterCoalescingRefineByUnion],
	// [PublicPageViewAnalyticsFilterCoalescingRefineByUnion],
	// [PublicCtaAnalyticsFilterCoalescingRefineByUnion],
	// [PublicEventAnalyticsFilterCoalescingRefineByUnion],
	// [PublicFormSubmissionFilterCoalescingRefineByUnion],
	// [PublicFormSubmissionOnPageFilterCoalescingRefineByUnion],
	// [PublicNumAssociationsFilterCoalescingRefineByUnion],
	// [PublicUnifiedEventsFilterCoalescingRefineByUnion],
	// [PublicPropertyAssociationInListFilterCoalescingRefineByUnion].
	UpperBoundOffset PublicTimeOffset `json:"upperBoundOffset"`
	// This field is from variant
	// [PublicAssociationInListFilterCoalescingRefineByUnion],
	// [PublicPageViewAnalyticsFilterCoalescingRefineByUnion],
	// [PublicCtaAnalyticsFilterCoalescingRefineByUnion],
	// [PublicEventAnalyticsFilterCoalescingRefineByUnion],
	// [PublicFormSubmissionFilterCoalescingRefineByUnion],
	// [PublicFormSubmissionOnPageFilterCoalescingRefineByUnion],
	// [PublicNumAssociationsFilterCoalescingRefineByUnion],
	// [PublicUnifiedEventsFilterCoalescingRefineByUnion],
	// [PublicPropertyAssociationInListFilterCoalescingRefineByUnion].
	Timestamp int64 `json:"timestamp"`
	// This field is from variant
	// [PublicAssociationInListFilterCoalescingRefineByUnion],
	// [PublicPageViewAnalyticsFilterCoalescingRefineByUnion],
	// [PublicCtaAnalyticsFilterCoalescingRefineByUnion],
	// [PublicEventAnalyticsFilterCoalescingRefineByUnion],
	// [PublicFormSubmissionFilterCoalescingRefineByUnion],
	// [PublicFormSubmissionOnPageFilterCoalescingRefineByUnion],
	// [PublicNumAssociationsFilterCoalescingRefineByUnion],
	// [PublicUnifiedEventsFilterCoalescingRefineByUnion],
	// [PublicPropertyAssociationInListFilterCoalescingRefineByUnion].
	LowerTimestamp int64 `json:"lowerTimestamp"`
	// This field is from variant
	// [PublicAssociationInListFilterCoalescingRefineByUnion],
	// [PublicPageViewAnalyticsFilterCoalescingRefineByUnion],
	// [PublicCtaAnalyticsFilterCoalescingRefineByUnion],
	// [PublicEventAnalyticsFilterCoalescingRefineByUnion],
	// [PublicFormSubmissionFilterCoalescingRefineByUnion],
	// [PublicFormSubmissionOnPageFilterCoalescingRefineByUnion],
	// [PublicNumAssociationsFilterCoalescingRefineByUnion],
	// [PublicUnifiedEventsFilterCoalescingRefineByUnion],
	// [PublicPropertyAssociationInListFilterCoalescingRefineByUnion].
	UpperTimestamp               int64  `json:"upperTimestamp"`
	IncludeObjectsWithNoValueSet bool   `json:"includeObjectsWithNoValueSet"`
	OperationType                string `json:"operationType"`
	Operator                     string `json:"operator"`
	// This field is from variant
	// [PublicAssociationInListFilterCoalescingRefineByUnion],
	// [PublicPageViewAnalyticsFilterCoalescingRefineByUnion],
	// [PublicCtaAnalyticsFilterCoalescingRefineByUnion],
	// [PublicEventAnalyticsFilterCoalescingRefineByUnion],
	// [PublicFormSubmissionFilterCoalescingRefineByUnion],
	// [PublicFormSubmissionOnPageFilterCoalescingRefineByUnion],
	// [PublicNumAssociationsFilterCoalescingRefineByUnion],
	// [PublicUnifiedEventsFilterCoalescingRefineByUnion],
	// [PublicPropertyAssociationInListFilterCoalescingRefineByUnion].
	TimePoint PublicTimePointOperationTimePointUnion `json:"timePoint"`
	// This field is from variant
	// [PublicAssociationInListFilterCoalescingRefineByUnion],
	// [PublicPageViewAnalyticsFilterCoalescingRefineByUnion],
	// [PublicCtaAnalyticsFilterCoalescingRefineByUnion],
	// [PublicEventAnalyticsFilterCoalescingRefineByUnion],
	// [PublicFormSubmissionFilterCoalescingRefineByUnion],
	// [PublicFormSubmissionOnPageFilterCoalescingRefineByUnion],
	// [PublicNumAssociationsFilterCoalescingRefineByUnion],
	// [PublicUnifiedEventsFilterCoalescingRefineByUnion],
	// [PublicPropertyAssociationInListFilterCoalescingRefineByUnion].
	EndpointBehavior string `json:"endpointBehavior"`
	PropertyParser   string `json:"propertyParser"`
	// This field is from variant
	// [PublicAssociationInListFilterCoalescingRefineByUnion],
	// [PublicPageViewAnalyticsFilterCoalescingRefineByUnion],
	// [PublicCtaAnalyticsFilterCoalescingRefineByUnion],
	// [PublicEventAnalyticsFilterCoalescingRefineByUnion],
	// [PublicFormSubmissionFilterCoalescingRefineByUnion],
	// [PublicFormSubmissionOnPageFilterCoalescingRefineByUnion],
	// [PublicNumAssociationsFilterCoalescingRefineByUnion],
	// [PublicUnifiedEventsFilterCoalescingRefineByUnion],
	// [PublicPropertyAssociationInListFilterCoalescingRefineByUnion].
	LowerBoundTimePoint PublicRangedTimeOperationLowerBoundTimePointUnion `json:"lowerBoundTimePoint"`
	// This field is from variant
	// [PublicAssociationInListFilterCoalescingRefineByUnion],
	// [PublicPageViewAnalyticsFilterCoalescingRefineByUnion],
	// [PublicCtaAnalyticsFilterCoalescingRefineByUnion],
	// [PublicEventAnalyticsFilterCoalescingRefineByUnion],
	// [PublicFormSubmissionFilterCoalescingRefineByUnion],
	// [PublicFormSubmissionOnPageFilterCoalescingRefineByUnion],
	// [PublicNumAssociationsFilterCoalescingRefineByUnion],
	// [PublicUnifiedEventsFilterCoalescingRefineByUnion],
	// [PublicPropertyAssociationInListFilterCoalescingRefineByUnion].
	UpperBoundTimePoint PublicRangedTimeOperationUpperBoundTimePointUnion `json:"upperBoundTimePoint"`
	// This field is from variant
	// [PublicAssociationInListFilterCoalescingRefineByUnion],
	// [PublicPageViewAnalyticsFilterCoalescingRefineByUnion],
	// [PublicCtaAnalyticsFilterCoalescingRefineByUnion],
	// [PublicEventAnalyticsFilterCoalescingRefineByUnion],
	// [PublicFormSubmissionFilterCoalescingRefineByUnion],
	// [PublicFormSubmissionOnPageFilterCoalescingRefineByUnion],
	// [PublicNumAssociationsFilterCoalescingRefineByUnion],
	// [PublicUnifiedEventsFilterCoalescingRefineByUnion],
	// [PublicPropertyAssociationInListFilterCoalescingRefineByUnion].
	LowerBoundEndpointBehavior string `json:"lowerBoundEndpointBehavior"`
	// This field is from variant
	// [PublicAssociationInListFilterCoalescingRefineByUnion],
	// [PublicPageViewAnalyticsFilterCoalescingRefineByUnion],
	// [PublicCtaAnalyticsFilterCoalescingRefineByUnion],
	// [PublicEventAnalyticsFilterCoalescingRefineByUnion],
	// [PublicFormSubmissionFilterCoalescingRefineByUnion],
	// [PublicFormSubmissionOnPageFilterCoalescingRefineByUnion],
	// [PublicNumAssociationsFilterCoalescingRefineByUnion],
	// [PublicUnifiedEventsFilterCoalescingRefineByUnion],
	// [PublicPropertyAssociationInListFilterCoalescingRefineByUnion].
	UpperBoundEndpointBehavior string `json:"upperBoundEndpointBehavior"`
	JSON                       struct {
		Type                         respjson.Field
		MaxOccurrences               respjson.Field
		MinOccurrences               respjson.Field
		SetType                      respjson.Field
		Comparison                   respjson.Field
		TimeOffset                   respjson.Field
		LowerBoundOffset             respjson.Field
		RangeType                    respjson.Field
		UpperBoundOffset             respjson.Field
		Timestamp                    respjson.Field
		LowerTimestamp               respjson.Field
		UpperTimestamp               respjson.Field
		IncludeObjectsWithNoValueSet respjson.Field
		OperationType                respjson.Field
		Operator                     respjson.Field
		TimePoint                    respjson.Field
		EndpointBehavior             respjson.Field
		PropertyParser               respjson.Field
		LowerBoundTimePoint          respjson.Field
		UpperBoundTimePoint          respjson.Field
		LowerBoundEndpointBehavior   respjson.Field
		UpperBoundEndpointBehavior   respjson.Field
		raw                          string
	} `json:"-"`
}

func (r *PublicNotAnyFilterBranchFilterUnionCoalescingRefineBy) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PublicNotAnyFilterBranchFilterUnionPruningRefineBy is an implicit subunion of
// [PublicNotAnyFilterBranchFilterUnion].
// PublicNotAnyFilterBranchFilterUnionPruningRefineBy provides convenient access to
// the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [PublicNotAnyFilterBranchFilterUnion].
type PublicNotAnyFilterBranchFilterUnionPruningRefineBy struct {
	Type string `json:"type"`
	// This field is from variant [PublicPageViewAnalyticsFilterPruningRefineByUnion],
	// [PublicCtaAnalyticsFilterPruningRefineByUnion],
	// [PublicEventAnalyticsFilterPruningRefineByUnion],
	// [PublicFormSubmissionFilterPruningRefineByUnion],
	// [PublicFormSubmissionOnPageFilterPruningRefineByUnion],
	// [PublicEmailEventFilterPruningRefineByUnion],
	// [PublicAdsTimeFilterPruningRefineByUnion],
	// [PublicUnifiedEventsFilterPruningRefineByUnion].
	MaxOccurrences int64 `json:"maxOccurrences"`
	// This field is from variant [PublicPageViewAnalyticsFilterPruningRefineByUnion],
	// [PublicCtaAnalyticsFilterPruningRefineByUnion],
	// [PublicEventAnalyticsFilterPruningRefineByUnion],
	// [PublicFormSubmissionFilterPruningRefineByUnion],
	// [PublicFormSubmissionOnPageFilterPruningRefineByUnion],
	// [PublicEmailEventFilterPruningRefineByUnion],
	// [PublicAdsTimeFilterPruningRefineByUnion],
	// [PublicUnifiedEventsFilterPruningRefineByUnion].
	MinOccurrences int64 `json:"minOccurrences"`
	// This field is from variant [PublicPageViewAnalyticsFilterPruningRefineByUnion],
	// [PublicCtaAnalyticsFilterPruningRefineByUnion],
	// [PublicEventAnalyticsFilterPruningRefineByUnion],
	// [PublicFormSubmissionFilterPruningRefineByUnion],
	// [PublicFormSubmissionOnPageFilterPruningRefineByUnion],
	// [PublicEmailEventFilterPruningRefineByUnion],
	// [PublicAdsTimeFilterPruningRefineByUnion],
	// [PublicUnifiedEventsFilterPruningRefineByUnion].
	SetType    string `json:"setType"`
	Comparison string `json:"comparison"`
	// This field is from variant [PublicPageViewAnalyticsFilterPruningRefineByUnion],
	// [PublicCtaAnalyticsFilterPruningRefineByUnion],
	// [PublicEventAnalyticsFilterPruningRefineByUnion],
	// [PublicFormSubmissionFilterPruningRefineByUnion],
	// [PublicFormSubmissionOnPageFilterPruningRefineByUnion],
	// [PublicEmailEventFilterPruningRefineByUnion],
	// [PublicAdsTimeFilterPruningRefineByUnion],
	// [PublicUnifiedEventsFilterPruningRefineByUnion].
	TimeOffset PublicTimeOffset `json:"timeOffset"`
	// This field is from variant [PublicPageViewAnalyticsFilterPruningRefineByUnion],
	// [PublicCtaAnalyticsFilterPruningRefineByUnion],
	// [PublicEventAnalyticsFilterPruningRefineByUnion],
	// [PublicFormSubmissionFilterPruningRefineByUnion],
	// [PublicFormSubmissionOnPageFilterPruningRefineByUnion],
	// [PublicEmailEventFilterPruningRefineByUnion],
	// [PublicAdsTimeFilterPruningRefineByUnion],
	// [PublicUnifiedEventsFilterPruningRefineByUnion].
	LowerBoundOffset PublicTimeOffset `json:"lowerBoundOffset"`
	RangeType        string           `json:"rangeType"`
	// This field is from variant [PublicPageViewAnalyticsFilterPruningRefineByUnion],
	// [PublicCtaAnalyticsFilterPruningRefineByUnion],
	// [PublicEventAnalyticsFilterPruningRefineByUnion],
	// [PublicFormSubmissionFilterPruningRefineByUnion],
	// [PublicFormSubmissionOnPageFilterPruningRefineByUnion],
	// [PublicEmailEventFilterPruningRefineByUnion],
	// [PublicAdsTimeFilterPruningRefineByUnion],
	// [PublicUnifiedEventsFilterPruningRefineByUnion].
	UpperBoundOffset PublicTimeOffset `json:"upperBoundOffset"`
	// This field is from variant [PublicPageViewAnalyticsFilterPruningRefineByUnion],
	// [PublicCtaAnalyticsFilterPruningRefineByUnion],
	// [PublicEventAnalyticsFilterPruningRefineByUnion],
	// [PublicFormSubmissionFilterPruningRefineByUnion],
	// [PublicFormSubmissionOnPageFilterPruningRefineByUnion],
	// [PublicEmailEventFilterPruningRefineByUnion],
	// [PublicAdsTimeFilterPruningRefineByUnion],
	// [PublicUnifiedEventsFilterPruningRefineByUnion].
	Timestamp int64 `json:"timestamp"`
	// This field is from variant [PublicPageViewAnalyticsFilterPruningRefineByUnion],
	// [PublicCtaAnalyticsFilterPruningRefineByUnion],
	// [PublicEventAnalyticsFilterPruningRefineByUnion],
	// [PublicFormSubmissionFilterPruningRefineByUnion],
	// [PublicFormSubmissionOnPageFilterPruningRefineByUnion],
	// [PublicEmailEventFilterPruningRefineByUnion],
	// [PublicAdsTimeFilterPruningRefineByUnion],
	// [PublicUnifiedEventsFilterPruningRefineByUnion].
	LowerTimestamp int64 `json:"lowerTimestamp"`
	// This field is from variant [PublicPageViewAnalyticsFilterPruningRefineByUnion],
	// [PublicCtaAnalyticsFilterPruningRefineByUnion],
	// [PublicEventAnalyticsFilterPruningRefineByUnion],
	// [PublicFormSubmissionFilterPruningRefineByUnion],
	// [PublicFormSubmissionOnPageFilterPruningRefineByUnion],
	// [PublicEmailEventFilterPruningRefineByUnion],
	// [PublicAdsTimeFilterPruningRefineByUnion],
	// [PublicUnifiedEventsFilterPruningRefineByUnion].
	UpperTimestamp               int64  `json:"upperTimestamp"`
	IncludeObjectsWithNoValueSet bool   `json:"includeObjectsWithNoValueSet"`
	OperationType                string `json:"operationType"`
	Operator                     string `json:"operator"`
	// This field is from variant [PublicPageViewAnalyticsFilterPruningRefineByUnion],
	// [PublicCtaAnalyticsFilterPruningRefineByUnion],
	// [PublicEventAnalyticsFilterPruningRefineByUnion],
	// [PublicFormSubmissionFilterPruningRefineByUnion],
	// [PublicFormSubmissionOnPageFilterPruningRefineByUnion],
	// [PublicEmailEventFilterPruningRefineByUnion],
	// [PublicAdsTimeFilterPruningRefineByUnion],
	// [PublicUnifiedEventsFilterPruningRefineByUnion].
	TimePoint PublicTimePointOperationTimePointUnion `json:"timePoint"`
	// This field is from variant [PublicPageViewAnalyticsFilterPruningRefineByUnion],
	// [PublicCtaAnalyticsFilterPruningRefineByUnion],
	// [PublicEventAnalyticsFilterPruningRefineByUnion],
	// [PublicFormSubmissionFilterPruningRefineByUnion],
	// [PublicFormSubmissionOnPageFilterPruningRefineByUnion],
	// [PublicEmailEventFilterPruningRefineByUnion],
	// [PublicAdsTimeFilterPruningRefineByUnion],
	// [PublicUnifiedEventsFilterPruningRefineByUnion].
	EndpointBehavior string `json:"endpointBehavior"`
	PropertyParser   string `json:"propertyParser"`
	// This field is from variant [PublicPageViewAnalyticsFilterPruningRefineByUnion],
	// [PublicCtaAnalyticsFilterPruningRefineByUnion],
	// [PublicEventAnalyticsFilterPruningRefineByUnion],
	// [PublicFormSubmissionFilterPruningRefineByUnion],
	// [PublicFormSubmissionOnPageFilterPruningRefineByUnion],
	// [PublicEmailEventFilterPruningRefineByUnion],
	// [PublicAdsTimeFilterPruningRefineByUnion],
	// [PublicUnifiedEventsFilterPruningRefineByUnion].
	LowerBoundTimePoint PublicRangedTimeOperationLowerBoundTimePointUnion `json:"lowerBoundTimePoint"`
	// This field is from variant [PublicPageViewAnalyticsFilterPruningRefineByUnion],
	// [PublicCtaAnalyticsFilterPruningRefineByUnion],
	// [PublicEventAnalyticsFilterPruningRefineByUnion],
	// [PublicFormSubmissionFilterPruningRefineByUnion],
	// [PublicFormSubmissionOnPageFilterPruningRefineByUnion],
	// [PublicEmailEventFilterPruningRefineByUnion],
	// [PublicAdsTimeFilterPruningRefineByUnion],
	// [PublicUnifiedEventsFilterPruningRefineByUnion].
	UpperBoundTimePoint PublicRangedTimeOperationUpperBoundTimePointUnion `json:"upperBoundTimePoint"`
	// This field is from variant [PublicPageViewAnalyticsFilterPruningRefineByUnion],
	// [PublicCtaAnalyticsFilterPruningRefineByUnion],
	// [PublicEventAnalyticsFilterPruningRefineByUnion],
	// [PublicFormSubmissionFilterPruningRefineByUnion],
	// [PublicFormSubmissionOnPageFilterPruningRefineByUnion],
	// [PublicEmailEventFilterPruningRefineByUnion],
	// [PublicAdsTimeFilterPruningRefineByUnion],
	// [PublicUnifiedEventsFilterPruningRefineByUnion].
	LowerBoundEndpointBehavior string `json:"lowerBoundEndpointBehavior"`
	// This field is from variant [PublicPageViewAnalyticsFilterPruningRefineByUnion],
	// [PublicCtaAnalyticsFilterPruningRefineByUnion],
	// [PublicEventAnalyticsFilterPruningRefineByUnion],
	// [PublicFormSubmissionFilterPruningRefineByUnion],
	// [PublicFormSubmissionOnPageFilterPruningRefineByUnion],
	// [PublicEmailEventFilterPruningRefineByUnion],
	// [PublicAdsTimeFilterPruningRefineByUnion],
	// [PublicUnifiedEventsFilterPruningRefineByUnion].
	UpperBoundEndpointBehavior string `json:"upperBoundEndpointBehavior"`
	JSON                       struct {
		Type                         respjson.Field
		MaxOccurrences               respjson.Field
		MinOccurrences               respjson.Field
		SetType                      respjson.Field
		Comparison                   respjson.Field
		TimeOffset                   respjson.Field
		LowerBoundOffset             respjson.Field
		RangeType                    respjson.Field
		UpperBoundOffset             respjson.Field
		Timestamp                    respjson.Field
		LowerTimestamp               respjson.Field
		UpperTimestamp               respjson.Field
		IncludeObjectsWithNoValueSet respjson.Field
		OperationType                respjson.Field
		Operator                     respjson.Field
		TimePoint                    respjson.Field
		EndpointBehavior             respjson.Field
		PropertyParser               respjson.Field
		LowerBoundTimePoint          respjson.Field
		UpperBoundTimePoint          respjson.Field
		LowerBoundEndpointBehavior   respjson.Field
		UpperBoundEndpointBehavior   respjson.Field
		raw                          string
	} `json:"-"`
}

func (r *PublicNotAnyFilterBranchFilterUnionPruningRefineBy) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PublicNotAnyFilterBranchFilterUnionEventTypeID is an implicit subunion of
// [PublicNotAnyFilterBranchFilterUnion].
// PublicNotAnyFilterBranchFilterUnionEventTypeID provides convenient access to the
// sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [PublicNotAnyFilterBranchFilterUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfInt OfString]
type PublicNotAnyFilterBranchFilterUnionEventTypeID struct {
	// This field will be present if the value is a [int64] instead of an object.
	OfInt int64 `json:",inline"`
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	JSON     struct {
		OfInt    respjson.Field
		OfString respjson.Field
		raw      string
	} `json:"-"`
}

func (r *PublicNotAnyFilterBranchFilterUnionEventTypeID) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties FilterBranches, FilterBranchOperator, FilterBranchType, Filters
// are required.
type PublicNotAnyFilterBranchParam struct {
	FilterBranches []PublicNotAnyFilterBranchFilterBranchUnionParam `json:"filterBranches,omitzero" api:"required"`
	// Specifies the logical operator used to combine filters within the branch
	// (NOT_ANY).
	FilterBranchOperator string `json:"filterBranchOperator" api:"required"`
	// Indicates the type of filter branch (NOT_ANY).
	//
	// Any of "NOT_ANY".
	FilterBranchType PublicNotAnyFilterBranchFilterBranchType   `json:"filterBranchType,omitzero" api:"required"`
	Filters          []PublicNotAnyFilterBranchFilterUnionParam `json:"filters,omitzero" api:"required"`
	paramObj
}

func (r PublicNotAnyFilterBranchParam) MarshalJSON() (data []byte, err error) {
	type shadow PublicNotAnyFilterBranchParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PublicNotAnyFilterBranchParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type PublicNotAnyFilterBranchFilterBranchUnionParam struct {
	OfOr                  *PublicOrFilterBranchParam                  `json:",omitzero,inline"`
	OfAnd                 *PublicAndFilterBranchParam                 `json:",omitzero,inline"`
	OfNotAll              *PublicNotAllFilterBranchParam              `json:",omitzero,inline"`
	OfNotAny              *PublicNotAnyFilterBranchParam              `json:",omitzero,inline"`
	OfRestricted          *PublicRestrictedFilterBranchParam          `json:",omitzero,inline"`
	OfUnifiedEvents       *PublicUnifiedEventsFilterBranchParam       `json:",omitzero,inline"`
	OfPropertyAssociation *PublicPropertyAssociationFilterBranchParam `json:",omitzero,inline"`
	OfAssociation         *PublicAssociationFilterBranchParam         `json:",omitzero,inline"`
	paramUnion
}

func (u PublicNotAnyFilterBranchFilterBranchUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfOr,
		u.OfAnd,
		u.OfNotAll,
		u.OfNotAny,
		u.OfRestricted,
		u.OfUnifiedEvents,
		u.OfPropertyAssociation,
		u.OfAssociation)
}
func (u *PublicNotAnyFilterBranchFilterBranchUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func init() {
	apijson.RegisterUnion[PublicNotAnyFilterBranchFilterBranchUnionParam](
		"filterBranchType",
		apijson.Discriminator[PublicOrFilterBranchParam]("OR"),
		apijson.Discriminator[PublicAndFilterBranchParam]("AND"),
		apijson.Discriminator[PublicNotAllFilterBranchParam]("NOT_ALL"),
		apijson.Discriminator[PublicNotAnyFilterBranchParam]("NOT_ANY"),
		apijson.Discriminator[PublicRestrictedFilterBranchParam]("RESTRICTED"),
		apijson.Discriminator[PublicUnifiedEventsFilterBranchParam]("UNIFIED_EVENTS"),
		apijson.Discriminator[PublicPropertyAssociationFilterBranchParam]("PROPERTY_ASSOCIATION"),
		apijson.Discriminator[PublicAssociationFilterBranchParam]("ASSOCIATION"),
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type PublicNotAnyFilterBranchFilterUnionParam struct {
	OfProperty                  *PublicPropertyFilterParam                  `json:",omitzero,inline"`
	OfAssociation               *PublicAssociationInListFilterParam         `json:",omitzero,inline"`
	OfPageView                  *PublicPageViewAnalyticsFilterParam         `json:",omitzero,inline"`
	OfCta                       *PublicCtaAnalyticsFilterParam              `json:",omitzero,inline"`
	OfEvent                     *PublicEventAnalyticsFilterParam            `json:",omitzero,inline"`
	OfFormSubmission            *PublicFormSubmissionFilterParam            `json:",omitzero,inline"`
	OfFormSubmissionOnPage      *PublicFormSubmissionOnPageFilterParam      `json:",omitzero,inline"`
	OfIntegrationEvent          *PublicIntegrationEventFilterParam          `json:",omitzero,inline"`
	OfEmailSubscription         *PublicEmailSubscriptionFilterParam         `json:",omitzero,inline"`
	OfCommunicationSubscription *PublicCommunicationSubscriptionFilterParam `json:",omitzero,inline"`
	OfCampaignInfluenced        *PublicCampaignInfluencedFilterParam        `json:",omitzero,inline"`
	OfSurveyMonkey              *PublicSurveyMonkeyFilterParam              `json:",omitzero,inline"`
	OfSurveyMonkeyValue         *PublicSurveyMonkeyValueFilterParam         `json:",omitzero,inline"`
	OfWebinar                   *PublicWebinarFilterParam                   `json:",omitzero,inline"`
	OfEmailEvent                *PublicEmailEventFilterParam                `json:",omitzero,inline"`
	OfPrivacy                   *PublicPrivacyAnalyticsFilterParam          `json:",omitzero,inline"`
	OfAdsSearch                 *PublicAdsSearchFilterParam                 `json:",omitzero,inline"`
	OfAdsTime                   *PublicAdsTimeFilterParam                   `json:",omitzero,inline"`
	OfInList                    *PublicInListFilterParam                    `json:",omitzero,inline"`
	OfNumAssociations           *PublicNumAssociationsFilterParam           `json:",omitzero,inline"`
	OfUnifiedEvents             *PublicUnifiedEventsFilterParam             `json:",omitzero,inline"`
	OfPropertyAssociation       *PublicPropertyAssociationInListFilterParam `json:",omitzero,inline"`
	OfConstant                  *PublicConstantFilterParam                  `json:",omitzero,inline"`
	paramUnion
}

func (u PublicNotAnyFilterBranchFilterUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfProperty,
		u.OfAssociation,
		u.OfPageView,
		u.OfCta,
		u.OfEvent,
		u.OfFormSubmission,
		u.OfFormSubmissionOnPage,
		u.OfIntegrationEvent,
		u.OfEmailSubscription,
		u.OfCommunicationSubscription,
		u.OfCampaignInfluenced,
		u.OfSurveyMonkey,
		u.OfSurveyMonkeyValue,
		u.OfWebinar,
		u.OfEmailEvent,
		u.OfPrivacy,
		u.OfAdsSearch,
		u.OfAdsTime,
		u.OfInList,
		u.OfNumAssociations,
		u.OfUnifiedEvents,
		u.OfPropertyAssociation,
		u.OfConstant)
}
func (u *PublicNotAnyFilterBranchFilterUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func init() {
	apijson.RegisterUnion[PublicNotAnyFilterBranchFilterUnionParam](
		"filterType",
		apijson.Discriminator[PublicPropertyFilterParam]("PROPERTY"),
		apijson.Discriminator[PublicAssociationInListFilterParam]("ASSOCIATION"),
		apijson.Discriminator[PublicPageViewAnalyticsFilterParam]("PAGE_VIEW"),
		apijson.Discriminator[PublicCtaAnalyticsFilterParam]("CTA"),
		apijson.Discriminator[PublicEventAnalyticsFilterParam]("EVENT"),
		apijson.Discriminator[PublicFormSubmissionFilterParam]("FORM_SUBMISSION"),
		apijson.Discriminator[PublicFormSubmissionOnPageFilterParam]("FORM_SUBMISSION_ON_PAGE"),
		apijson.Discriminator[PublicIntegrationEventFilterParam]("INTEGRATION_EVENT"),
		apijson.Discriminator[PublicEmailSubscriptionFilterParam]("EMAIL_SUBSCRIPTION"),
		apijson.Discriminator[PublicCommunicationSubscriptionFilterParam]("COMMUNICATION_SUBSCRIPTION"),
		apijson.Discriminator[PublicCampaignInfluencedFilterParam]("CAMPAIGN_INFLUENCED"),
		apijson.Discriminator[PublicSurveyMonkeyFilterParam]("SURVEY_MONKEY"),
		apijson.Discriminator[PublicSurveyMonkeyValueFilterParam]("SURVEY_MONKEY_VALUE"),
		apijson.Discriminator[PublicWebinarFilterParam]("WEBINAR"),
		apijson.Discriminator[PublicEmailEventFilterParam]("EMAIL_EVENT"),
		apijson.Discriminator[PublicPrivacyAnalyticsFilterParam]("PRIVACY"),
		apijson.Discriminator[PublicAdsSearchFilterParam]("ADS_SEARCH"),
		apijson.Discriminator[PublicAdsTimeFilterParam]("ADS_TIME"),
		apijson.Discriminator[PublicInListFilterParam]("IN_LIST"),
		apijson.Discriminator[PublicNumAssociationsFilterParam]("NUM_ASSOCIATIONS"),
		apijson.Discriminator[PublicUnifiedEventsFilterParam]("UNIFIED_EVENTS"),
		apijson.Discriminator[PublicPropertyAssociationInListFilterParam]("PROPERTY_ASSOCIATION"),
		apijson.Discriminator[PublicConstantFilterParam]("CONSTANT"),
	)
}

type PublicNowReference struct {
	// Indicates the type of reference (NOW).
	//
	// Any of "NOW".
	ReferenceType PublicNowReferenceReferenceType `json:"referenceType" api:"required"`
	// The hour component of the current time reference.
	Hour int64 `json:"hour"`
	// The millisecond component of the current time reference.
	Millisecond int64 `json:"millisecond"`
	// The minute component of the current time reference.
	Minute int64 `json:"minute"`
	// The second component of the current time reference.
	Second int64 `json:"second"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ReferenceType respjson.Field
		Hour          respjson.Field
		Millisecond   respjson.Field
		Minute        respjson.Field
		Second        respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PublicNowReference) RawJSON() string { return r.JSON.raw }
func (r *PublicNowReference) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this PublicNowReference to a PublicNowReferenceParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// PublicNowReferenceParam.Overrides()
func (r PublicNowReference) ToParam() PublicNowReferenceParam {
	return param.Override[PublicNowReferenceParam](json.RawMessage(r.RawJSON()))
}

// Indicates the type of reference (NOW).
type PublicNowReferenceReferenceType string

const (
	PublicNowReferenceReferenceTypeNow PublicNowReferenceReferenceType = "NOW"
)

// The property ReferenceType is required.
type PublicNowReferenceParam struct {
	// Indicates the type of reference (NOW).
	//
	// Any of "NOW".
	ReferenceType PublicNowReferenceReferenceType `json:"referenceType,omitzero" api:"required"`
	// The hour component of the current time reference.
	Hour param.Opt[int64] `json:"hour,omitzero"`
	// The millisecond component of the current time reference.
	Millisecond param.Opt[int64] `json:"millisecond,omitzero"`
	// The minute component of the current time reference.
	Minute param.Opt[int64] `json:"minute,omitzero"`
	// The second component of the current time reference.
	Second param.Opt[int64] `json:"second,omitzero"`
	paramObj
}

func (r PublicNowReferenceParam) MarshalJSON() (data []byte, err error) {
	type shadow PublicNowReferenceParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PublicNowReferenceParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PublicNumAssociationsFilter struct {
	// Defines the category of the association (HUBSPOT_DEFINED, USER_DEFINED,
	// INTEGRATOR_DEFINED, WORK).
	AssociationCategory string `json:"associationCategory" api:"required"`
	// The ID representing the type of association being filtered.
	AssociationTypeID int64 `json:"associationTypeId" api:"required"`
	// Specifies the criteria for refining the association filter.
	CoalescingRefineBy PublicNumAssociationsFilterCoalescingRefineByUnion `json:"coalescingRefineBy" api:"required"`
	// Indicates the type of filter being applied (NUM_ASSOCIATIONS).
	//
	// Any of "NUM_ASSOCIATIONS".
	FilterType PublicNumAssociationsFilterFilterType `json:"filterType" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AssociationCategory respjson.Field
		AssociationTypeID   respjson.Field
		CoalescingRefineBy  respjson.Field
		FilterType          respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PublicNumAssociationsFilter) RawJSON() string { return r.JSON.raw }
func (r *PublicNumAssociationsFilter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this PublicNumAssociationsFilter to a
// PublicNumAssociationsFilterParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// PublicNumAssociationsFilterParam.Overrides()
func (r PublicNumAssociationsFilter) ToParam() PublicNumAssociationsFilterParam {
	return param.Override[PublicNumAssociationsFilterParam](json.RawMessage(r.RawJSON()))
}

// PublicNumAssociationsFilterCoalescingRefineByUnion contains all possible
// properties and values from [PublicNumOccurrencesRefineBy],
// [PublicSetOccurrencesRefineBy], [PublicRelativeComparativeTimestampRefineBy],
// [PublicRelativeRangedTimestampRefineBy],
// [PublicAbsoluteComparativeTimestampRefineBy],
// [PublicAbsoluteRangedTimestampRefineBy], [PublicAllHistoryRefineBy],
// [PublicTimePointOperation], [PublicRangedTimeOperation].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type PublicNumAssociationsFilterCoalescingRefineByUnion struct {
	Type string `json:"type"`
	// This field is from variant [PublicNumOccurrencesRefineBy].
	MaxOccurrences int64 `json:"maxOccurrences"`
	// This field is from variant [PublicNumOccurrencesRefineBy].
	MinOccurrences int64 `json:"minOccurrences"`
	// This field is from variant [PublicSetOccurrencesRefineBy].
	SetType    string `json:"setType"`
	Comparison string `json:"comparison"`
	// This field is from variant [PublicRelativeComparativeTimestampRefineBy].
	TimeOffset PublicTimeOffset `json:"timeOffset"`
	// This field is from variant [PublicRelativeRangedTimestampRefineBy].
	LowerBoundOffset PublicTimeOffset `json:"lowerBoundOffset"`
	RangeType        string           `json:"rangeType"`
	// This field is from variant [PublicRelativeRangedTimestampRefineBy].
	UpperBoundOffset PublicTimeOffset `json:"upperBoundOffset"`
	// This field is from variant [PublicAbsoluteComparativeTimestampRefineBy].
	Timestamp int64 `json:"timestamp"`
	// This field is from variant [PublicAbsoluteRangedTimestampRefineBy].
	LowerTimestamp int64 `json:"lowerTimestamp"`
	// This field is from variant [PublicAbsoluteRangedTimestampRefineBy].
	UpperTimestamp               int64  `json:"upperTimestamp"`
	IncludeObjectsWithNoValueSet bool   `json:"includeObjectsWithNoValueSet"`
	OperationType                string `json:"operationType"`
	Operator                     string `json:"operator"`
	// This field is from variant [PublicTimePointOperation].
	TimePoint PublicTimePointOperationTimePointUnion `json:"timePoint"`
	// This field is from variant [PublicTimePointOperation].
	EndpointBehavior string `json:"endpointBehavior"`
	PropertyParser   string `json:"propertyParser"`
	// This field is from variant [PublicRangedTimeOperation].
	LowerBoundTimePoint PublicRangedTimeOperationLowerBoundTimePointUnion `json:"lowerBoundTimePoint"`
	// This field is from variant [PublicRangedTimeOperation].
	UpperBoundTimePoint PublicRangedTimeOperationUpperBoundTimePointUnion `json:"upperBoundTimePoint"`
	// This field is from variant [PublicRangedTimeOperation].
	LowerBoundEndpointBehavior string `json:"lowerBoundEndpointBehavior"`
	// This field is from variant [PublicRangedTimeOperation].
	UpperBoundEndpointBehavior string `json:"upperBoundEndpointBehavior"`
	JSON                       struct {
		Type                         respjson.Field
		MaxOccurrences               respjson.Field
		MinOccurrences               respjson.Field
		SetType                      respjson.Field
		Comparison                   respjson.Field
		TimeOffset                   respjson.Field
		LowerBoundOffset             respjson.Field
		RangeType                    respjson.Field
		UpperBoundOffset             respjson.Field
		Timestamp                    respjson.Field
		LowerTimestamp               respjson.Field
		UpperTimestamp               respjson.Field
		IncludeObjectsWithNoValueSet respjson.Field
		OperationType                respjson.Field
		Operator                     respjson.Field
		TimePoint                    respjson.Field
		EndpointBehavior             respjson.Field
		PropertyParser               respjson.Field
		LowerBoundTimePoint          respjson.Field
		UpperBoundTimePoint          respjson.Field
		LowerBoundEndpointBehavior   respjson.Field
		UpperBoundEndpointBehavior   respjson.Field
		raw                          string
	} `json:"-"`
}

func (u PublicNumAssociationsFilterCoalescingRefineByUnion) AsNumOccurrences() (v PublicNumOccurrencesRefineBy) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicNumAssociationsFilterCoalescingRefineByUnion) AsSetOccurrences() (v PublicSetOccurrencesRefineBy) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicNumAssociationsFilterCoalescingRefineByUnion) AsRelativeComparative() (v PublicRelativeComparativeTimestampRefineBy) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicNumAssociationsFilterCoalescingRefineByUnion) AsRelativeRanged() (v PublicRelativeRangedTimestampRefineBy) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicNumAssociationsFilterCoalescingRefineByUnion) AsAbsoluteComparative() (v PublicAbsoluteComparativeTimestampRefineBy) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicNumAssociationsFilterCoalescingRefineByUnion) AsAbsoluteRanged() (v PublicAbsoluteRangedTimestampRefineBy) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicNumAssociationsFilterCoalescingRefineByUnion) AsAllHistory() (v PublicAllHistoryRefineBy) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicNumAssociationsFilterCoalescingRefineByUnion) AsTimePoint() (v PublicTimePointOperation) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicNumAssociationsFilterCoalescingRefineByUnion) AsTimeRanged() (v PublicRangedTimeOperation) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u PublicNumAssociationsFilterCoalescingRefineByUnion) RawJSON() string { return u.JSON.raw }

func (r *PublicNumAssociationsFilterCoalescingRefineByUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Indicates the type of filter being applied (NUM_ASSOCIATIONS).
type PublicNumAssociationsFilterFilterType string

const (
	PublicNumAssociationsFilterFilterTypeNumAssociations PublicNumAssociationsFilterFilterType = "NUM_ASSOCIATIONS"
)

// The properties AssociationCategory, AssociationTypeID, CoalescingRefineBy,
// FilterType are required.
type PublicNumAssociationsFilterParam struct {
	// Defines the category of the association (HUBSPOT_DEFINED, USER_DEFINED,
	// INTEGRATOR_DEFINED, WORK).
	AssociationCategory string `json:"associationCategory" api:"required"`
	// The ID representing the type of association being filtered.
	AssociationTypeID int64 `json:"associationTypeId" api:"required"`
	// Specifies the criteria for refining the association filter.
	CoalescingRefineBy PublicNumAssociationsFilterCoalescingRefineByUnionParam `json:"coalescingRefineBy,omitzero" api:"required"`
	// Indicates the type of filter being applied (NUM_ASSOCIATIONS).
	//
	// Any of "NUM_ASSOCIATIONS".
	FilterType PublicNumAssociationsFilterFilterType `json:"filterType,omitzero" api:"required"`
	paramObj
}

func (r PublicNumAssociationsFilterParam) MarshalJSON() (data []byte, err error) {
	type shadow PublicNumAssociationsFilterParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PublicNumAssociationsFilterParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type PublicNumAssociationsFilterCoalescingRefineByUnionParam struct {
	OfNumOccurrences      *PublicNumOccurrencesRefineByParam               `json:",omitzero,inline"`
	OfSetOccurrences      *PublicSetOccurrencesRefineByParam               `json:",omitzero,inline"`
	OfRelativeComparative *PublicRelativeComparativeTimestampRefineByParam `json:",omitzero,inline"`
	OfRelativeRanged      *PublicRelativeRangedTimestampRefineByParam      `json:",omitzero,inline"`
	OfAbsoluteComparative *PublicAbsoluteComparativeTimestampRefineByParam `json:",omitzero,inline"`
	OfAbsoluteRanged      *PublicAbsoluteRangedTimestampRefineByParam      `json:",omitzero,inline"`
	OfAllHistory          *PublicAllHistoryRefineByParam                   `json:",omitzero,inline"`
	OfTimePoint           *PublicTimePointOperationParam                   `json:",omitzero,inline"`
	OfTimeRanged          *PublicRangedTimeOperationParam                  `json:",omitzero,inline"`
	paramUnion
}

func (u PublicNumAssociationsFilterCoalescingRefineByUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfNumOccurrences,
		u.OfSetOccurrences,
		u.OfRelativeComparative,
		u.OfRelativeRanged,
		u.OfAbsoluteComparative,
		u.OfAbsoluteRanged,
		u.OfAllHistory,
		u.OfTimePoint,
		u.OfTimeRanged)
}
func (u *PublicNumAssociationsFilterCoalescingRefineByUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

type PublicNumOccurrencesRefineBy struct {
	// The type of refinement (NUM_OCCURRENCES).
	//
	// Any of "NUM_OCCURRENCES".
	Type PublicNumOccurrencesRefineByType `json:"type" api:"required"`
	// The maximum number of occurrences allowed.
	MaxOccurrences int64 `json:"maxOccurrences"`
	// The minimum number of occurrences required.
	MinOccurrences int64 `json:"minOccurrences"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type           respjson.Field
		MaxOccurrences respjson.Field
		MinOccurrences respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PublicNumOccurrencesRefineBy) RawJSON() string { return r.JSON.raw }
func (r *PublicNumOccurrencesRefineBy) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this PublicNumOccurrencesRefineBy to a
// PublicNumOccurrencesRefineByParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// PublicNumOccurrencesRefineByParam.Overrides()
func (r PublicNumOccurrencesRefineBy) ToParam() PublicNumOccurrencesRefineByParam {
	return param.Override[PublicNumOccurrencesRefineByParam](json.RawMessage(r.RawJSON()))
}

// The type of refinement (NUM_OCCURRENCES).
type PublicNumOccurrencesRefineByType string

const (
	PublicNumOccurrencesRefineByTypeNumOccurrences PublicNumOccurrencesRefineByType = "NUM_OCCURRENCES"
)

// The property Type is required.
type PublicNumOccurrencesRefineByParam struct {
	// The type of refinement (NUM_OCCURRENCES).
	//
	// Any of "NUM_OCCURRENCES".
	Type PublicNumOccurrencesRefineByType `json:"type,omitzero" api:"required"`
	// The maximum number of occurrences allowed.
	MaxOccurrences param.Opt[int64] `json:"maxOccurrences,omitzero"`
	// The minimum number of occurrences required.
	MinOccurrences param.Opt[int64] `json:"minOccurrences,omitzero"`
	paramObj
}

func (r PublicNumOccurrencesRefineByParam) MarshalJSON() (data []byte, err error) {
	type shadow PublicNumOccurrencesRefineByParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PublicNumOccurrencesRefineByParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PublicNumberPropertyOperation struct {
	// Indicates whether objects with no value set for the property should be included
	// in the operation.
	IncludeObjectsWithNoValueSet bool `json:"includeObjectsWithNoValueSet" api:"required"`
	// Specifies the type of operation (NUMBER).
	//
	// Any of "NUMBER".
	OperationType PublicNumberPropertyOperationOperationType `json:"operationType" api:"required"`
	// Defines the operation to be applied in the number property operation
	// (IS_EQUAL_TO, IS_NOT_EQUAL_TO, IS_GREATER_THAN, IS_GREATER_THAN_OR_EQUAL_TO,
	// IS_LESS_THAN, IS_LESS_THAN_OR_EQUAL_TO, HAS_EVER_BEEN_EQUAL_TO,
	// HAS_NEVER_BEEN_EQUAL_TO).
	Operator string `json:"operator" api:"required"`
	// The numeric value to be used in the operation.
	Value float64 `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		IncludeObjectsWithNoValueSet respjson.Field
		OperationType                respjson.Field
		Operator                     respjson.Field
		Value                        respjson.Field
		ExtraFields                  map[string]respjson.Field
		raw                          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PublicNumberPropertyOperation) RawJSON() string { return r.JSON.raw }
func (r *PublicNumberPropertyOperation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this PublicNumberPropertyOperation to a
// PublicNumberPropertyOperationParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// PublicNumberPropertyOperationParam.Overrides()
func (r PublicNumberPropertyOperation) ToParam() PublicNumberPropertyOperationParam {
	return param.Override[PublicNumberPropertyOperationParam](json.RawMessage(r.RawJSON()))
}

// Specifies the type of operation (NUMBER).
type PublicNumberPropertyOperationOperationType string

const (
	PublicNumberPropertyOperationOperationTypeNumber PublicNumberPropertyOperationOperationType = "NUMBER"
)

// The properties IncludeObjectsWithNoValueSet, OperationType, Operator, Value are
// required.
type PublicNumberPropertyOperationParam struct {
	// Indicates whether objects with no value set for the property should be included
	// in the operation.
	IncludeObjectsWithNoValueSet bool `json:"includeObjectsWithNoValueSet" api:"required"`
	// Specifies the type of operation (NUMBER).
	//
	// Any of "NUMBER".
	OperationType PublicNumberPropertyOperationOperationType `json:"operationType,omitzero" api:"required"`
	// Defines the operation to be applied in the number property operation
	// (IS_EQUAL_TO, IS_NOT_EQUAL_TO, IS_GREATER_THAN, IS_GREATER_THAN_OR_EQUAL_TO,
	// IS_LESS_THAN, IS_LESS_THAN_OR_EQUAL_TO, HAS_EVER_BEEN_EQUAL_TO,
	// HAS_NEVER_BEEN_EQUAL_TO).
	Operator string `json:"operator" api:"required"`
	// The numeric value to be used in the operation.
	Value float64 `json:"value" api:"required"`
	paramObj
}

func (r PublicNumberPropertyOperationParam) MarshalJSON() (data []byte, err error) {
	type shadow PublicNumberPropertyOperationParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PublicNumberPropertyOperationParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PublicObjectList struct {
	// The **ILS ID** of the list.
	ListID string `json:"listId" api:"required"`
	// The version of the list.
	ListVersion int64 `json:"listVersion" api:"required"`
	// The name of the list.
	Name string `json:"name" api:"required"`
	// The object type of the list.
	ObjectTypeID string `json:"objectTypeId" api:"required"`
	// The processing status of the list.
	ProcessingStatus string `json:"processingStatus" api:"required"`
	// The processing type of the list.
	ProcessingType string `json:"processingType" api:"required"`
	// The time when the list was created.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// The ID of the user that created the list.
	CreatedByID string `json:"createdById"`
	// The time when the list was deleted.
	DeletedAt time.Time `json:"deletedAt" format:"date-time"`
	// Defines the filter criteria for the list, allowing for complex logical
	// operations and nested filter branches to determine list membership.
	FilterBranch PublicObjectListFilterBranchUnion `json:"filterBranch"`
	// The time when the filters for this list were last updated.
	FiltersUpdatedAt   time.Time                `json:"filtersUpdatedAt" format:"date-time"`
	ListPermissions    PublicListPermissions    `json:"listPermissions"`
	MembershipSettings PublicMembershipSettings `json:"membershipSettings"`
	// Size of the list
	Size int64 `json:"size"`
	// The time the list was last updated.
	UpdatedAt time.Time `json:"updatedAt" format:"date-time"`
	// The ID of the user that last updated the list.
	UpdatedByID string `json:"updatedById"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ListID             respjson.Field
		ListVersion        respjson.Field
		Name               respjson.Field
		ObjectTypeID       respjson.Field
		ProcessingStatus   respjson.Field
		ProcessingType     respjson.Field
		CreatedAt          respjson.Field
		CreatedByID        respjson.Field
		DeletedAt          respjson.Field
		FilterBranch       respjson.Field
		FiltersUpdatedAt   respjson.Field
		ListPermissions    respjson.Field
		MembershipSettings respjson.Field
		Size               respjson.Field
		UpdatedAt          respjson.Field
		UpdatedByID        respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PublicObjectList) RawJSON() string { return r.JSON.raw }
func (r *PublicObjectList) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PublicObjectListFilterBranchUnion contains all possible properties and values
// from [PublicOrFilterBranch], [PublicAndFilterBranch],
// [PublicNotAllFilterBranch], [PublicNotAnyFilterBranch],
// [PublicRestrictedFilterBranch], [PublicUnifiedEventsFilterBranch],
// [PublicPropertyAssociationFilterBranch], [PublicAssociationFilterBranch].
//
// Use the [PublicObjectListFilterBranchUnion.AsAny] method to switch on the
// variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type PublicObjectListFilterBranchUnion struct {
	// This field is a union of [[]PublicOrFilterBranchFilterBranchUnion],
	// [[]PublicAndFilterBranchFilterBranchUnion],
	// [[]PublicNotAllFilterBranchFilterBranchUnion],
	// [[]PublicNotAnyFilterBranchFilterBranchUnion],
	// [[]PublicRestrictedFilterBranchFilterBranchUnion],
	// [[]PublicUnifiedEventsFilterBranchFilterBranchUnion],
	// [[]PublicPropertyAssociationFilterBranchFilterBranchUnion],
	// [[]PublicAssociationFilterBranchFilterBranchUnion]
	FilterBranches       PublicObjectListFilterBranchUnionFilterBranches `json:"filterBranches"`
	FilterBranchOperator string                                          `json:"filterBranchOperator"`
	// Any of "OR", "AND", "NOT_ALL", "NOT_ANY", "RESTRICTED", "UNIFIED_EVENTS",
	// "PROPERTY_ASSOCIATION", "ASSOCIATION".
	FilterBranchType string `json:"filterBranchType"`
	// This field is a union of [[]PublicOrFilterBranchFilterUnion],
	// [[]PublicAndFilterBranchFilterUnion], [[]PublicNotAllFilterBranchFilterUnion],
	// [[]PublicNotAnyFilterBranchFilterUnion],
	// [[]PublicRestrictedFilterBranchFilterUnion],
	// [[]PublicUnifiedEventsFilterBranchFilterUnion],
	// [[]PublicPropertyAssociationFilterBranchFilterUnion],
	// [[]PublicAssociationFilterBranchFilterUnion]
	Filters PublicObjectListFilterBranchUnionFilters `json:"filters"`
	// This field is from variant [PublicUnifiedEventsFilterBranch].
	EventTypeID string `json:"eventTypeId"`
	Operator    string `json:"operator"`
	// This field is from variant [PublicUnifiedEventsFilterBranch].
	CoalescingRefineBy PublicUnifiedEventsFilterBranchCoalescingRefineByUnion `json:"coalescingRefineBy"`
	// This field is from variant [PublicUnifiedEventsFilterBranch].
	PruningRefineBy PublicUnifiedEventsFilterBranchPruningRefineByUnion `json:"pruningRefineBy"`
	ObjectTypeID    string                                              `json:"objectTypeId"`
	// This field is from variant [PublicPropertyAssociationFilterBranch].
	PropertyWithObjectID string `json:"propertyWithObjectId"`
	// This field is from variant [PublicAssociationFilterBranch].
	AssociationCategory string `json:"associationCategory"`
	// This field is from variant [PublicAssociationFilterBranch].
	AssociationTypeID int64 `json:"associationTypeId"`
	JSON              struct {
		FilterBranches       respjson.Field
		FilterBranchOperator respjson.Field
		FilterBranchType     respjson.Field
		Filters              respjson.Field
		EventTypeID          respjson.Field
		Operator             respjson.Field
		CoalescingRefineBy   respjson.Field
		PruningRefineBy      respjson.Field
		ObjectTypeID         respjson.Field
		PropertyWithObjectID respjson.Field
		AssociationCategory  respjson.Field
		AssociationTypeID    respjson.Field
		raw                  string
	} `json:"-"`
}

// anyPublicObjectListFilterBranch is implemented by each variant of
// [PublicObjectListFilterBranchUnion] to add type safety for the return type of
// [PublicObjectListFilterBranchUnion.AsAny]
type anyPublicObjectListFilterBranch interface {
	implPublicObjectListFilterBranchUnion()
}

func (PublicOrFilterBranch) implPublicObjectListFilterBranchUnion()                  {}
func (PublicAndFilterBranch) implPublicObjectListFilterBranchUnion()                 {}
func (PublicNotAllFilterBranch) implPublicObjectListFilterBranchUnion()              {}
func (PublicNotAnyFilterBranch) implPublicObjectListFilterBranchUnion()              {}
func (PublicRestrictedFilterBranch) implPublicObjectListFilterBranchUnion()          {}
func (PublicUnifiedEventsFilterBranch) implPublicObjectListFilterBranchUnion()       {}
func (PublicPropertyAssociationFilterBranch) implPublicObjectListFilterBranchUnion() {}
func (PublicAssociationFilterBranch) implPublicObjectListFilterBranchUnion()         {}

// Use the following switch statement to find the correct variant
//
//	switch variant := PublicObjectListFilterBranchUnion.AsAny().(type) {
//	case crm.PublicOrFilterBranch:
//	case crm.PublicAndFilterBranch:
//	case crm.PublicNotAllFilterBranch:
//	case crm.PublicNotAnyFilterBranch:
//	case crm.PublicRestrictedFilterBranch:
//	case crm.PublicUnifiedEventsFilterBranch:
//	case crm.PublicPropertyAssociationFilterBranch:
//	case crm.PublicAssociationFilterBranch:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u PublicObjectListFilterBranchUnion) AsAny() anyPublicObjectListFilterBranch {
	switch u.FilterBranchType {
	case "OR":
		return u.AsOr()
	case "AND":
		return u.AsAnd()
	case "NOT_ALL":
		return u.AsNotAll()
	case "NOT_ANY":
		return u.AsNotAny()
	case "RESTRICTED":
		return u.AsRestricted()
	case "UNIFIED_EVENTS":
		return u.AsUnifiedEvents()
	case "PROPERTY_ASSOCIATION":
		return u.AsPropertyAssociation()
	case "ASSOCIATION":
		return u.AsAssociation()
	}
	return nil
}

func (u PublicObjectListFilterBranchUnion) AsOr() (v PublicOrFilterBranch) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicObjectListFilterBranchUnion) AsAnd() (v PublicAndFilterBranch) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicObjectListFilterBranchUnion) AsNotAll() (v PublicNotAllFilterBranch) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicObjectListFilterBranchUnion) AsNotAny() (v PublicNotAnyFilterBranch) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicObjectListFilterBranchUnion) AsRestricted() (v PublicRestrictedFilterBranch) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicObjectListFilterBranchUnion) AsUnifiedEvents() (v PublicUnifiedEventsFilterBranch) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicObjectListFilterBranchUnion) AsPropertyAssociation() (v PublicPropertyAssociationFilterBranch) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicObjectListFilterBranchUnion) AsAssociation() (v PublicAssociationFilterBranch) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u PublicObjectListFilterBranchUnion) RawJSON() string { return u.JSON.raw }

func (r *PublicObjectListFilterBranchUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PublicObjectListFilterBranchUnionFilterBranches is an implicit subunion of
// [PublicObjectListFilterBranchUnion].
// PublicObjectListFilterBranchUnionFilterBranches provides convenient access to
// the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [PublicObjectListFilterBranchUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfPublicOrFilterBranchFilterBranches
// OfPublicAndFilterBranchFilterBranches OfPublicNotAllFilterBranchFilterBranches
// OfPublicNotAnyFilterBranchFilterBranches
// OfPublicRestrictedFilterBranchFilterBranches
// OfPublicUnifiedEventsFilterBranchFilterBranches
// OfPublicPropertyAssociationFilterBranchFilterBranches
// OfPublicAssociationFilterBranchFilterBranches]
type PublicObjectListFilterBranchUnionFilterBranches struct {
	// This field will be present if the value is a
	// [[]PublicOrFilterBranchFilterBranchUnion] instead of an object.
	OfPublicOrFilterBranchFilterBranches []PublicOrFilterBranchFilterBranchUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]PublicAndFilterBranchFilterBranchUnion] instead of an object.
	OfPublicAndFilterBranchFilterBranches []PublicAndFilterBranchFilterBranchUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]PublicNotAllFilterBranchFilterBranchUnion] instead of an object.
	OfPublicNotAllFilterBranchFilterBranches []PublicNotAllFilterBranchFilterBranchUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]PublicNotAnyFilterBranchFilterBranchUnion] instead of an object.
	OfPublicNotAnyFilterBranchFilterBranches []PublicNotAnyFilterBranchFilterBranchUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]PublicRestrictedFilterBranchFilterBranchUnion] instead of an object.
	OfPublicRestrictedFilterBranchFilterBranches []PublicRestrictedFilterBranchFilterBranchUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]PublicUnifiedEventsFilterBranchFilterBranchUnion] instead of an object.
	OfPublicUnifiedEventsFilterBranchFilterBranches []PublicUnifiedEventsFilterBranchFilterBranchUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]PublicPropertyAssociationFilterBranchFilterBranchUnion] instead of an object.
	OfPublicPropertyAssociationFilterBranchFilterBranches []PublicPropertyAssociationFilterBranchFilterBranchUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]PublicAssociationFilterBranchFilterBranchUnion] instead of an object.
	OfPublicAssociationFilterBranchFilterBranches []PublicAssociationFilterBranchFilterBranchUnion `json:",inline"`
	JSON                                          struct {
		OfPublicOrFilterBranchFilterBranches                  respjson.Field
		OfPublicAndFilterBranchFilterBranches                 respjson.Field
		OfPublicNotAllFilterBranchFilterBranches              respjson.Field
		OfPublicNotAnyFilterBranchFilterBranches              respjson.Field
		OfPublicRestrictedFilterBranchFilterBranches          respjson.Field
		OfPublicUnifiedEventsFilterBranchFilterBranches       respjson.Field
		OfPublicPropertyAssociationFilterBranchFilterBranches respjson.Field
		OfPublicAssociationFilterBranchFilterBranches         respjson.Field
		raw                                                   string
	} `json:"-"`
}

func (r *PublicObjectListFilterBranchUnionFilterBranches) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PublicObjectListFilterBranchUnionFilters is an implicit subunion of
// [PublicObjectListFilterBranchUnion]. PublicObjectListFilterBranchUnionFilters
// provides convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [PublicObjectListFilterBranchUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfPublicOrFilterBranchFilters OfPublicAndFilterBranchFilters
// OfPublicNotAllFilterBranchFilters OfPublicNotAnyFilterBranchFilters
// OfPublicRestrictedFilterBranchFilters OfPublicUnifiedEventsFilterBranchFilters
// OfPublicPropertyAssociationFilterBranchFilters
// OfPublicAssociationFilterBranchFilters]
type PublicObjectListFilterBranchUnionFilters struct {
	// This field will be present if the value is a [[]PublicOrFilterBranchFilterUnion]
	// instead of an object.
	OfPublicOrFilterBranchFilters []PublicOrFilterBranchFilterUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]PublicAndFilterBranchFilterUnion] instead of an object.
	OfPublicAndFilterBranchFilters []PublicAndFilterBranchFilterUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]PublicNotAllFilterBranchFilterUnion] instead of an object.
	OfPublicNotAllFilterBranchFilters []PublicNotAllFilterBranchFilterUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]PublicNotAnyFilterBranchFilterUnion] instead of an object.
	OfPublicNotAnyFilterBranchFilters []PublicNotAnyFilterBranchFilterUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]PublicRestrictedFilterBranchFilterUnion] instead of an object.
	OfPublicRestrictedFilterBranchFilters []PublicRestrictedFilterBranchFilterUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]PublicUnifiedEventsFilterBranchFilterUnion] instead of an object.
	OfPublicUnifiedEventsFilterBranchFilters []PublicUnifiedEventsFilterBranchFilterUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]PublicPropertyAssociationFilterBranchFilterUnion] instead of an object.
	OfPublicPropertyAssociationFilterBranchFilters []PublicPropertyAssociationFilterBranchFilterUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]PublicAssociationFilterBranchFilterUnion] instead of an object.
	OfPublicAssociationFilterBranchFilters []PublicAssociationFilterBranchFilterUnion `json:",inline"`
	JSON                                   struct {
		OfPublicOrFilterBranchFilters                  respjson.Field
		OfPublicAndFilterBranchFilters                 respjson.Field
		OfPublicNotAllFilterBranchFilters              respjson.Field
		OfPublicNotAnyFilterBranchFilters              respjson.Field
		OfPublicRestrictedFilterBranchFilters          respjson.Field
		OfPublicUnifiedEventsFilterBranchFilters       respjson.Field
		OfPublicPropertyAssociationFilterBranchFilters respjson.Field
		OfPublicAssociationFilterBranchFilters         respjson.Field
		raw                                            string
	} `json:"-"`
}

func (r *PublicObjectListFilterBranchUnionFilters) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PublicObjectListSearchResult struct {
	// The **ILS ID** of the list.
	ListID string `json:"listId" api:"required"`
	// The version of the list.
	ListVersion int64 `json:"listVersion" api:"required"`
	// The name of the list.
	Name string `json:"name" api:"required"`
	// The object type of the list.
	ObjectTypeID string `json:"objectTypeId" api:"required"`
	// The processing status of the list.
	ProcessingStatus string `json:"processingStatus" api:"required"`
	// The processing type of the list.
	ProcessingType string `json:"processingType" api:"required"`
	// The name and value of any additional properties that exist for this list and
	// that were included in the search request.
	AdditionalFilterProperties map[string]string `json:"additional_filter_properties"`
	// The time when the list was created.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// The ID of the user that created the list.
	CreatedByID string `json:"createdById"`
	// The time when the list was deleted.
	DeletedAt time.Time `json:"deletedAt" format:"date-time"`
	// The time when the filters for this list were last updated.
	FiltersUpdatedAt time.Time `json:"filtersUpdatedAt" format:"date-time"`
	// The time the list was last updated.
	UpdatedAt time.Time `json:"updatedAt" format:"date-time"`
	// The ID of the user that last updated the list.
	UpdatedByID string `json:"updatedById"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ListID                     respjson.Field
		ListVersion                respjson.Field
		Name                       respjson.Field
		ObjectTypeID               respjson.Field
		ProcessingStatus           respjson.Field
		ProcessingType             respjson.Field
		AdditionalFilterProperties respjson.Field
		CreatedAt                  respjson.Field
		CreatedByID                respjson.Field
		DeletedAt                  respjson.Field
		FiltersUpdatedAt           respjson.Field
		UpdatedAt                  respjson.Field
		UpdatedByID                respjson.Field
		ExtraFields                map[string]respjson.Field
		raw                        string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PublicObjectListSearchResult) RawJSON() string { return r.JSON.raw }
func (r *PublicObjectListSearchResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PublicOrFilterBranch struct {
	FilterBranches []PublicOrFilterBranchFilterBranchUnion `json:"filterBranches" api:"required"`
	// The logical operator used to combine the filters within the branch (OR).
	FilterBranchOperator string `json:"filterBranchOperator" api:"required"`
	// The type of the filter branch (OR).
	//
	// Any of "OR".
	FilterBranchType PublicOrFilterBranchFilterBranchType `json:"filterBranchType" api:"required"`
	Filters          []PublicOrFilterBranchFilterUnion    `json:"filters" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FilterBranches       respjson.Field
		FilterBranchOperator respjson.Field
		FilterBranchType     respjson.Field
		Filters              respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PublicOrFilterBranch) RawJSON() string { return r.JSON.raw }
func (r *PublicOrFilterBranch) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this PublicOrFilterBranch to a PublicOrFilterBranchParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// PublicOrFilterBranchParam.Overrides()
func (r PublicOrFilterBranch) ToParam() PublicOrFilterBranchParam {
	return param.Override[PublicOrFilterBranchParam](json.RawMessage(r.RawJSON()))
}

// PublicOrFilterBranchFilterBranchUnion contains all possible properties and
// values from [PublicOrFilterBranch], [PublicAndFilterBranch],
// [PublicNotAllFilterBranch], [PublicNotAnyFilterBranch],
// [PublicRestrictedFilterBranch], [PublicUnifiedEventsFilterBranch],
// [PublicPropertyAssociationFilterBranch], [PublicAssociationFilterBranch].
//
// Use the [PublicOrFilterBranchFilterBranchUnion.AsAny] method to switch on the
// variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type PublicOrFilterBranchFilterBranchUnion struct {
	// This field is a union of [[]PublicOrFilterBranchFilterBranchUnion],
	// [[]PublicAndFilterBranchFilterBranchUnion],
	// [[]PublicNotAllFilterBranchFilterBranchUnion],
	// [[]PublicNotAnyFilterBranchFilterBranchUnion],
	// [[]PublicRestrictedFilterBranchFilterBranchUnion],
	// [[]PublicUnifiedEventsFilterBranchFilterBranchUnion],
	// [[]PublicPropertyAssociationFilterBranchFilterBranchUnion],
	// [[]PublicAssociationFilterBranchFilterBranchUnion]
	FilterBranches       PublicOrFilterBranchFilterBranchUnionFilterBranches `json:"filterBranches"`
	FilterBranchOperator string                                              `json:"filterBranchOperator"`
	// Any of "OR", "AND", "NOT_ALL", "NOT_ANY", "RESTRICTED", "UNIFIED_EVENTS",
	// "PROPERTY_ASSOCIATION", "ASSOCIATION".
	FilterBranchType string `json:"filterBranchType"`
	// This field is a union of [[]PublicOrFilterBranchFilterUnion],
	// [[]PublicAndFilterBranchFilterUnion], [[]PublicNotAllFilterBranchFilterUnion],
	// [[]PublicNotAnyFilterBranchFilterUnion],
	// [[]PublicRestrictedFilterBranchFilterUnion],
	// [[]PublicUnifiedEventsFilterBranchFilterUnion],
	// [[]PublicPropertyAssociationFilterBranchFilterUnion],
	// [[]PublicAssociationFilterBranchFilterUnion]
	Filters PublicOrFilterBranchFilterBranchUnionFilters `json:"filters"`
	// This field is from variant [PublicUnifiedEventsFilterBranch].
	EventTypeID string `json:"eventTypeId"`
	Operator    string `json:"operator"`
	// This field is from variant [PublicUnifiedEventsFilterBranch].
	CoalescingRefineBy PublicUnifiedEventsFilterBranchCoalescingRefineByUnion `json:"coalescingRefineBy"`
	// This field is from variant [PublicUnifiedEventsFilterBranch].
	PruningRefineBy PublicUnifiedEventsFilterBranchPruningRefineByUnion `json:"pruningRefineBy"`
	ObjectTypeID    string                                              `json:"objectTypeId"`
	// This field is from variant [PublicPropertyAssociationFilterBranch].
	PropertyWithObjectID string `json:"propertyWithObjectId"`
	// This field is from variant [PublicAssociationFilterBranch].
	AssociationCategory string `json:"associationCategory"`
	// This field is from variant [PublicAssociationFilterBranch].
	AssociationTypeID int64 `json:"associationTypeId"`
	JSON              struct {
		FilterBranches       respjson.Field
		FilterBranchOperator respjson.Field
		FilterBranchType     respjson.Field
		Filters              respjson.Field
		EventTypeID          respjson.Field
		Operator             respjson.Field
		CoalescingRefineBy   respjson.Field
		PruningRefineBy      respjson.Field
		ObjectTypeID         respjson.Field
		PropertyWithObjectID respjson.Field
		AssociationCategory  respjson.Field
		AssociationTypeID    respjson.Field
		raw                  string
	} `json:"-"`
}

// anyPublicOrFilterBranchFilterBranch is implemented by each variant of
// [PublicOrFilterBranchFilterBranchUnion] to add type safety for the return type
// of [PublicOrFilterBranchFilterBranchUnion.AsAny]
type anyPublicOrFilterBranchFilterBranch interface {
	implPublicOrFilterBranchFilterBranchUnion()
}

func (PublicOrFilterBranch) implPublicOrFilterBranchFilterBranchUnion()                  {}
func (PublicAndFilterBranch) implPublicOrFilterBranchFilterBranchUnion()                 {}
func (PublicNotAllFilterBranch) implPublicOrFilterBranchFilterBranchUnion()              {}
func (PublicNotAnyFilterBranch) implPublicOrFilterBranchFilterBranchUnion()              {}
func (PublicRestrictedFilterBranch) implPublicOrFilterBranchFilterBranchUnion()          {}
func (PublicUnifiedEventsFilterBranch) implPublicOrFilterBranchFilterBranchUnion()       {}
func (PublicPropertyAssociationFilterBranch) implPublicOrFilterBranchFilterBranchUnion() {}
func (PublicAssociationFilterBranch) implPublicOrFilterBranchFilterBranchUnion()         {}

// Use the following switch statement to find the correct variant
//
//	switch variant := PublicOrFilterBranchFilterBranchUnion.AsAny().(type) {
//	case crm.PublicOrFilterBranch:
//	case crm.PublicAndFilterBranch:
//	case crm.PublicNotAllFilterBranch:
//	case crm.PublicNotAnyFilterBranch:
//	case crm.PublicRestrictedFilterBranch:
//	case crm.PublicUnifiedEventsFilterBranch:
//	case crm.PublicPropertyAssociationFilterBranch:
//	case crm.PublicAssociationFilterBranch:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u PublicOrFilterBranchFilterBranchUnion) AsAny() anyPublicOrFilterBranchFilterBranch {
	switch u.FilterBranchType {
	case "OR":
		return u.AsOr()
	case "AND":
		return u.AsAnd()
	case "NOT_ALL":
		return u.AsNotAll()
	case "NOT_ANY":
		return u.AsNotAny()
	case "RESTRICTED":
		return u.AsRestricted()
	case "UNIFIED_EVENTS":
		return u.AsUnifiedEvents()
	case "PROPERTY_ASSOCIATION":
		return u.AsPropertyAssociation()
	case "ASSOCIATION":
		return u.AsAssociation()
	}
	return nil
}

func (u PublicOrFilterBranchFilterBranchUnion) AsOr() (v PublicOrFilterBranch) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicOrFilterBranchFilterBranchUnion) AsAnd() (v PublicAndFilterBranch) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicOrFilterBranchFilterBranchUnion) AsNotAll() (v PublicNotAllFilterBranch) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicOrFilterBranchFilterBranchUnion) AsNotAny() (v PublicNotAnyFilterBranch) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicOrFilterBranchFilterBranchUnion) AsRestricted() (v PublicRestrictedFilterBranch) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicOrFilterBranchFilterBranchUnion) AsUnifiedEvents() (v PublicUnifiedEventsFilterBranch) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicOrFilterBranchFilterBranchUnion) AsPropertyAssociation() (v PublicPropertyAssociationFilterBranch) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicOrFilterBranchFilterBranchUnion) AsAssociation() (v PublicAssociationFilterBranch) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u PublicOrFilterBranchFilterBranchUnion) RawJSON() string { return u.JSON.raw }

func (r *PublicOrFilterBranchFilterBranchUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PublicOrFilterBranchFilterBranchUnionFilterBranches is an implicit subunion of
// [PublicOrFilterBranchFilterBranchUnion].
// PublicOrFilterBranchFilterBranchUnionFilterBranches provides convenient access
// to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [PublicOrFilterBranchFilterBranchUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfPublicOrFilterBranchFilterBranches
// OfPublicAndFilterBranchFilterBranches OfPublicNotAllFilterBranchFilterBranches
// OfPublicNotAnyFilterBranchFilterBranches
// OfPublicRestrictedFilterBranchFilterBranches
// OfPublicUnifiedEventsFilterBranchFilterBranches
// OfPublicPropertyAssociationFilterBranchFilterBranches
// OfPublicAssociationFilterBranchFilterBranches]
type PublicOrFilterBranchFilterBranchUnionFilterBranches struct {
	// This field will be present if the value is a
	// [[]PublicOrFilterBranchFilterBranchUnion] instead of an object.
	OfPublicOrFilterBranchFilterBranches []PublicOrFilterBranchFilterBranchUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]PublicAndFilterBranchFilterBranchUnion] instead of an object.
	OfPublicAndFilterBranchFilterBranches []PublicAndFilterBranchFilterBranchUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]PublicNotAllFilterBranchFilterBranchUnion] instead of an object.
	OfPublicNotAllFilterBranchFilterBranches []PublicNotAllFilterBranchFilterBranchUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]PublicNotAnyFilterBranchFilterBranchUnion] instead of an object.
	OfPublicNotAnyFilterBranchFilterBranches []PublicNotAnyFilterBranchFilterBranchUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]PublicRestrictedFilterBranchFilterBranchUnion] instead of an object.
	OfPublicRestrictedFilterBranchFilterBranches []PublicRestrictedFilterBranchFilterBranchUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]PublicUnifiedEventsFilterBranchFilterBranchUnion] instead of an object.
	OfPublicUnifiedEventsFilterBranchFilterBranches []PublicUnifiedEventsFilterBranchFilterBranchUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]PublicPropertyAssociationFilterBranchFilterBranchUnion] instead of an object.
	OfPublicPropertyAssociationFilterBranchFilterBranches []PublicPropertyAssociationFilterBranchFilterBranchUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]PublicAssociationFilterBranchFilterBranchUnion] instead of an object.
	OfPublicAssociationFilterBranchFilterBranches []PublicAssociationFilterBranchFilterBranchUnion `json:",inline"`
	JSON                                          struct {
		OfPublicOrFilterBranchFilterBranches                  respjson.Field
		OfPublicAndFilterBranchFilterBranches                 respjson.Field
		OfPublicNotAllFilterBranchFilterBranches              respjson.Field
		OfPublicNotAnyFilterBranchFilterBranches              respjson.Field
		OfPublicRestrictedFilterBranchFilterBranches          respjson.Field
		OfPublicUnifiedEventsFilterBranchFilterBranches       respjson.Field
		OfPublicPropertyAssociationFilterBranchFilterBranches respjson.Field
		OfPublicAssociationFilterBranchFilterBranches         respjson.Field
		raw                                                   string
	} `json:"-"`
}

func (r *PublicOrFilterBranchFilterBranchUnionFilterBranches) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PublicOrFilterBranchFilterBranchUnionFilters is an implicit subunion of
// [PublicOrFilterBranchFilterBranchUnion].
// PublicOrFilterBranchFilterBranchUnionFilters provides convenient access to the
// sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [PublicOrFilterBranchFilterBranchUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfPublicOrFilterBranchFilters OfPublicAndFilterBranchFilters
// OfPublicNotAllFilterBranchFilters OfPublicNotAnyFilterBranchFilters
// OfPublicRestrictedFilterBranchFilters OfPublicUnifiedEventsFilterBranchFilters
// OfPublicPropertyAssociationFilterBranchFilters
// OfPublicAssociationFilterBranchFilters]
type PublicOrFilterBranchFilterBranchUnionFilters struct {
	// This field will be present if the value is a [[]PublicOrFilterBranchFilterUnion]
	// instead of an object.
	OfPublicOrFilterBranchFilters []PublicOrFilterBranchFilterUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]PublicAndFilterBranchFilterUnion] instead of an object.
	OfPublicAndFilterBranchFilters []PublicAndFilterBranchFilterUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]PublicNotAllFilterBranchFilterUnion] instead of an object.
	OfPublicNotAllFilterBranchFilters []PublicNotAllFilterBranchFilterUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]PublicNotAnyFilterBranchFilterUnion] instead of an object.
	OfPublicNotAnyFilterBranchFilters []PublicNotAnyFilterBranchFilterUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]PublicRestrictedFilterBranchFilterUnion] instead of an object.
	OfPublicRestrictedFilterBranchFilters []PublicRestrictedFilterBranchFilterUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]PublicUnifiedEventsFilterBranchFilterUnion] instead of an object.
	OfPublicUnifiedEventsFilterBranchFilters []PublicUnifiedEventsFilterBranchFilterUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]PublicPropertyAssociationFilterBranchFilterUnion] instead of an object.
	OfPublicPropertyAssociationFilterBranchFilters []PublicPropertyAssociationFilterBranchFilterUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]PublicAssociationFilterBranchFilterUnion] instead of an object.
	OfPublicAssociationFilterBranchFilters []PublicAssociationFilterBranchFilterUnion `json:",inline"`
	JSON                                   struct {
		OfPublicOrFilterBranchFilters                  respjson.Field
		OfPublicAndFilterBranchFilters                 respjson.Field
		OfPublicNotAllFilterBranchFilters              respjson.Field
		OfPublicNotAnyFilterBranchFilters              respjson.Field
		OfPublicRestrictedFilterBranchFilters          respjson.Field
		OfPublicUnifiedEventsFilterBranchFilters       respjson.Field
		OfPublicPropertyAssociationFilterBranchFilters respjson.Field
		OfPublicAssociationFilterBranchFilters         respjson.Field
		raw                                            string
	} `json:"-"`
}

func (r *PublicOrFilterBranchFilterBranchUnionFilters) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The type of the filter branch (OR).
type PublicOrFilterBranchFilterBranchType string

const (
	PublicOrFilterBranchFilterBranchTypeOr PublicOrFilterBranchFilterBranchType = "OR"
)

// PublicOrFilterBranchFilterUnion contains all possible properties and values from
// [PublicPropertyFilter], [PublicAssociationInListFilter],
// [PublicPageViewAnalyticsFilter], [PublicCtaAnalyticsFilter],
// [PublicEventAnalyticsFilter], [PublicFormSubmissionFilter],
// [PublicFormSubmissionOnPageFilter], [PublicIntegrationEventFilter],
// [PublicEmailSubscriptionFilter], [PublicCommunicationSubscriptionFilter],
// [PublicCampaignInfluencedFilter], [PublicSurveyMonkeyFilter],
// [PublicSurveyMonkeyValueFilter], [PublicWebinarFilter],
// [PublicEmailEventFilter], [PublicPrivacyAnalyticsFilter],
// [PublicAdsSearchFilter], [PublicAdsTimeFilter], [PublicInListFilter],
// [PublicNumAssociationsFilter], [PublicUnifiedEventsFilter],
// [PublicPropertyAssociationInListFilter], [PublicConstantFilter].
//
// Use the [PublicOrFilterBranchFilterUnion.AsAny] method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type PublicOrFilterBranchFilterUnion struct {
	// Any of "PROPERTY", "ASSOCIATION", "PAGE_VIEW", "CTA", "EVENT",
	// "FORM_SUBMISSION", "FORM_SUBMISSION_ON_PAGE", "INTEGRATION_EVENT",
	// "EMAIL_SUBSCRIPTION", "COMMUNICATION_SUBSCRIPTION", "CAMPAIGN_INFLUENCED",
	// "SURVEY_MONKEY", "SURVEY_MONKEY_VALUE", "WEBINAR", "EMAIL_EVENT", "PRIVACY",
	// "ADS_SEARCH", "ADS_TIME", "IN_LIST", "NUM_ASSOCIATIONS", "UNIFIED_EVENTS",
	// "PROPERTY_ASSOCIATION", "CONSTANT".
	FilterType string `json:"filterType"`
	// This field is from variant [PublicPropertyFilter].
	Operation PublicPropertyFilterOperationUnion `json:"operation"`
	// This field is from variant [PublicPropertyFilter].
	Property            string `json:"property"`
	AssociationCategory string `json:"associationCategory"`
	AssociationTypeID   int64  `json:"associationTypeId"`
	// This field is a union of [PublicAssociationInListFilterCoalescingRefineByUnion],
	// [PublicPageViewAnalyticsFilterCoalescingRefineByUnion],
	// [PublicCtaAnalyticsFilterCoalescingRefineByUnion],
	// [PublicEventAnalyticsFilterCoalescingRefineByUnion],
	// [PublicFormSubmissionFilterCoalescingRefineByUnion],
	// [PublicFormSubmissionOnPageFilterCoalescingRefineByUnion],
	// [PublicNumAssociationsFilterCoalescingRefineByUnion],
	// [PublicUnifiedEventsFilterCoalescingRefineByUnion],
	// [PublicPropertyAssociationInListFilterCoalescingRefineByUnion]
	CoalescingRefineBy PublicOrFilterBranchFilterUnionCoalescingRefineBy `json:"coalescingRefineBy"`
	ListID             string                                            `json:"listId"`
	Operator           string                                            `json:"operator"`
	// This field is from variant [PublicAssociationInListFilter].
	ToObjectType   string `json:"toObjectType"`
	ToObjectTypeID string `json:"toObjectTypeId"`
	// This field is from variant [PublicPageViewAnalyticsFilter].
	PageURL string `json:"pageUrl"`
	// This field is from variant [PublicPageViewAnalyticsFilter].
	EnableTracking bool `json:"enableTracking"`
	// This field is a union of [PublicPageViewAnalyticsFilterPruningRefineByUnion],
	// [PublicCtaAnalyticsFilterPruningRefineByUnion],
	// [PublicEventAnalyticsFilterPruningRefineByUnion],
	// [PublicFormSubmissionFilterPruningRefineByUnion],
	// [PublicFormSubmissionOnPageFilterPruningRefineByUnion],
	// [PublicEmailEventFilterPruningRefineByUnion],
	// [PublicAdsTimeFilterPruningRefineByUnion],
	// [PublicUnifiedEventsFilterPruningRefineByUnion]
	PruningRefineBy PublicOrFilterBranchFilterUnionPruningRefineBy `json:"pruningRefineBy"`
	// This field is from variant [PublicCtaAnalyticsFilter].
	CtaName string `json:"ctaName"`
	// This field is from variant [PublicEventAnalyticsFilter].
	EventID string `json:"eventId"`
	FormID  string `json:"formId"`
	// This field is from variant [PublicFormSubmissionOnPageFilter].
	PageID string `json:"pageId"`
	// This field is a union of [int64], [string]
	EventTypeID PublicOrFilterBranchFilterUnionEventTypeID `json:"eventTypeId"`
	FilterLines []PublicEventFilterMetadata                `json:"filterLines"`
	// This field is from variant [PublicEmailSubscriptionFilter].
	AcceptedStatuses []string `json:"acceptedStatuses"`
	SubscriptionIDs  []string `json:"subscriptionIds"`
	SubscriptionType string   `json:"subscriptionType"`
	// This field is from variant [PublicCommunicationSubscriptionFilter].
	AcceptedOptStates []string `json:"acceptedOptStates"`
	// This field is from variant [PublicCommunicationSubscriptionFilter].
	Channel string `json:"channel"`
	// This field is from variant [PublicCommunicationSubscriptionFilter].
	BusinessUnitID string `json:"businessUnitId"`
	// This field is from variant [PublicCampaignInfluencedFilter].
	CampaignID string `json:"campaignId"`
	SurveyID   string `json:"surveyId"`
	// This field is from variant [PublicSurveyMonkeyValueFilter].
	SurveyQuestion string `json:"surveyQuestion"`
	// This field is from variant [PublicSurveyMonkeyValueFilter].
	ValueComparison PublicSurveyMonkeyValueFilterValueComparisonUnion `json:"valueComparison"`
	// This field is from variant [PublicSurveyMonkeyValueFilter].
	SurveyAnswerColID string `json:"surveyAnswerColId"`
	// This field is from variant [PublicSurveyMonkeyValueFilter].
	SurveyAnswerRowID string `json:"surveyAnswerRowId"`
	// This field is from variant [PublicWebinarFilter].
	WebinarID string `json:"webinarId"`
	// This field is from variant [PublicEmailEventFilter].
	AppID string `json:"appId"`
	// This field is from variant [PublicEmailEventFilter].
	EmailID string `json:"emailId"`
	// This field is from variant [PublicEmailEventFilter].
	Level string `json:"level"`
	// This field is from variant [PublicEmailEventFilter].
	ClickURL string `json:"clickUrl"`
	// This field is from variant [PublicPrivacyAnalyticsFilter].
	PrivacyName string `json:"privacyName"`
	// This field is from variant [PublicAdsSearchFilter].
	AdNetwork string `json:"adNetwork"`
	// This field is from variant [PublicAdsSearchFilter].
	EntityType string `json:"entityType"`
	// This field is from variant [PublicAdsSearchFilter].
	SearchTerms []string `json:"searchTerms"`
	// This field is from variant [PublicAdsSearchFilter].
	SearchTermType string `json:"searchTermType"`
	// This field is from variant [PublicInListFilter].
	Metadata PublicInListFilterMetadata `json:"metadata"`
	// This field is from variant [PublicPropertyAssociationInListFilter].
	PropertyWithObjectID string `json:"propertyWithObjectId"`
	// This field is from variant [PublicConstantFilter].
	ShouldAccept bool `json:"shouldAccept"`
	// This field is from variant [PublicConstantFilter].
	Source string `json:"source"`
	JSON   struct {
		FilterType           respjson.Field
		Operation            respjson.Field
		Property             respjson.Field
		AssociationCategory  respjson.Field
		AssociationTypeID    respjson.Field
		CoalescingRefineBy   respjson.Field
		ListID               respjson.Field
		Operator             respjson.Field
		ToObjectType         respjson.Field
		ToObjectTypeID       respjson.Field
		PageURL              respjson.Field
		EnableTracking       respjson.Field
		PruningRefineBy      respjson.Field
		CtaName              respjson.Field
		EventID              respjson.Field
		FormID               respjson.Field
		PageID               respjson.Field
		EventTypeID          respjson.Field
		FilterLines          respjson.Field
		AcceptedStatuses     respjson.Field
		SubscriptionIDs      respjson.Field
		SubscriptionType     respjson.Field
		AcceptedOptStates    respjson.Field
		Channel              respjson.Field
		BusinessUnitID       respjson.Field
		CampaignID           respjson.Field
		SurveyID             respjson.Field
		SurveyQuestion       respjson.Field
		ValueComparison      respjson.Field
		SurveyAnswerColID    respjson.Field
		SurveyAnswerRowID    respjson.Field
		WebinarID            respjson.Field
		AppID                respjson.Field
		EmailID              respjson.Field
		Level                respjson.Field
		ClickURL             respjson.Field
		PrivacyName          respjson.Field
		AdNetwork            respjson.Field
		EntityType           respjson.Field
		SearchTerms          respjson.Field
		SearchTermType       respjson.Field
		Metadata             respjson.Field
		PropertyWithObjectID respjson.Field
		ShouldAccept         respjson.Field
		Source               respjson.Field
		raw                  string
	} `json:"-"`
}

// anyPublicOrFilterBranchFilter is implemented by each variant of
// [PublicOrFilterBranchFilterUnion] to add type safety for the return type of
// [PublicOrFilterBranchFilterUnion.AsAny]
type anyPublicOrFilterBranchFilter interface {
	implPublicOrFilterBranchFilterUnion()
}

func (PublicPropertyFilter) implPublicOrFilterBranchFilterUnion()                  {}
func (PublicAssociationInListFilter) implPublicOrFilterBranchFilterUnion()         {}
func (PublicPageViewAnalyticsFilter) implPublicOrFilterBranchFilterUnion()         {}
func (PublicCtaAnalyticsFilter) implPublicOrFilterBranchFilterUnion()              {}
func (PublicEventAnalyticsFilter) implPublicOrFilterBranchFilterUnion()            {}
func (PublicFormSubmissionFilter) implPublicOrFilterBranchFilterUnion()            {}
func (PublicFormSubmissionOnPageFilter) implPublicOrFilterBranchFilterUnion()      {}
func (PublicIntegrationEventFilter) implPublicOrFilterBranchFilterUnion()          {}
func (PublicEmailSubscriptionFilter) implPublicOrFilterBranchFilterUnion()         {}
func (PublicCommunicationSubscriptionFilter) implPublicOrFilterBranchFilterUnion() {}
func (PublicCampaignInfluencedFilter) implPublicOrFilterBranchFilterUnion()        {}
func (PublicSurveyMonkeyFilter) implPublicOrFilterBranchFilterUnion()              {}
func (PublicSurveyMonkeyValueFilter) implPublicOrFilterBranchFilterUnion()         {}
func (PublicWebinarFilter) implPublicOrFilterBranchFilterUnion()                   {}
func (PublicEmailEventFilter) implPublicOrFilterBranchFilterUnion()                {}
func (PublicPrivacyAnalyticsFilter) implPublicOrFilterBranchFilterUnion()          {}
func (PublicAdsSearchFilter) implPublicOrFilterBranchFilterUnion()                 {}
func (PublicAdsTimeFilter) implPublicOrFilterBranchFilterUnion()                   {}
func (PublicInListFilter) implPublicOrFilterBranchFilterUnion()                    {}
func (PublicNumAssociationsFilter) implPublicOrFilterBranchFilterUnion()           {}
func (PublicUnifiedEventsFilter) implPublicOrFilterBranchFilterUnion()             {}
func (PublicPropertyAssociationInListFilter) implPublicOrFilterBranchFilterUnion() {}
func (PublicConstantFilter) implPublicOrFilterBranchFilterUnion()                  {}

// Use the following switch statement to find the correct variant
//
//	switch variant := PublicOrFilterBranchFilterUnion.AsAny().(type) {
//	case crm.PublicPropertyFilter:
//	case crm.PublicAssociationInListFilter:
//	case crm.PublicPageViewAnalyticsFilter:
//	case crm.PublicCtaAnalyticsFilter:
//	case crm.PublicEventAnalyticsFilter:
//	case crm.PublicFormSubmissionFilter:
//	case crm.PublicFormSubmissionOnPageFilter:
//	case crm.PublicIntegrationEventFilter:
//	case crm.PublicEmailSubscriptionFilter:
//	case crm.PublicCommunicationSubscriptionFilter:
//	case crm.PublicCampaignInfluencedFilter:
//	case crm.PublicSurveyMonkeyFilter:
//	case crm.PublicSurveyMonkeyValueFilter:
//	case crm.PublicWebinarFilter:
//	case crm.PublicEmailEventFilter:
//	case crm.PublicPrivacyAnalyticsFilter:
//	case crm.PublicAdsSearchFilter:
//	case crm.PublicAdsTimeFilter:
//	case crm.PublicInListFilter:
//	case crm.PublicNumAssociationsFilter:
//	case crm.PublicUnifiedEventsFilter:
//	case crm.PublicPropertyAssociationInListFilter:
//	case crm.PublicConstantFilter:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u PublicOrFilterBranchFilterUnion) AsAny() anyPublicOrFilterBranchFilter {
	switch u.FilterType {
	case "PROPERTY":
		return u.AsProperty()
	case "ASSOCIATION":
		return u.AsAssociation()
	case "PAGE_VIEW":
		return u.AsPageView()
	case "CTA":
		return u.AsCta()
	case "EVENT":
		return u.AsEvent()
	case "FORM_SUBMISSION":
		return u.AsFormSubmission()
	case "FORM_SUBMISSION_ON_PAGE":
		return u.AsFormSubmissionOnPage()
	case "INTEGRATION_EVENT":
		return u.AsIntegrationEvent()
	case "EMAIL_SUBSCRIPTION":
		return u.AsEmailSubscription()
	case "COMMUNICATION_SUBSCRIPTION":
		return u.AsCommunicationSubscription()
	case "CAMPAIGN_INFLUENCED":
		return u.AsCampaignInfluenced()
	case "SURVEY_MONKEY":
		return u.AsSurveyMonkey()
	case "SURVEY_MONKEY_VALUE":
		return u.AsSurveyMonkeyValue()
	case "WEBINAR":
		return u.AsWebinar()
	case "EMAIL_EVENT":
		return u.AsEmailEvent()
	case "PRIVACY":
		return u.AsPrivacy()
	case "ADS_SEARCH":
		return u.AsAdsSearch()
	case "ADS_TIME":
		return u.AsAdsTime()
	case "IN_LIST":
		return u.AsInList()
	case "NUM_ASSOCIATIONS":
		return u.AsNumAssociations()
	case "UNIFIED_EVENTS":
		return u.AsUnifiedEvents()
	case "PROPERTY_ASSOCIATION":
		return u.AsPropertyAssociation()
	case "CONSTANT":
		return u.AsConstant()
	}
	return nil
}

func (u PublicOrFilterBranchFilterUnion) AsProperty() (v PublicPropertyFilter) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicOrFilterBranchFilterUnion) AsAssociation() (v PublicAssociationInListFilter) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicOrFilterBranchFilterUnion) AsPageView() (v PublicPageViewAnalyticsFilter) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicOrFilterBranchFilterUnion) AsCta() (v PublicCtaAnalyticsFilter) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicOrFilterBranchFilterUnion) AsEvent() (v PublicEventAnalyticsFilter) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicOrFilterBranchFilterUnion) AsFormSubmission() (v PublicFormSubmissionFilter) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicOrFilterBranchFilterUnion) AsFormSubmissionOnPage() (v PublicFormSubmissionOnPageFilter) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicOrFilterBranchFilterUnion) AsIntegrationEvent() (v PublicIntegrationEventFilter) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicOrFilterBranchFilterUnion) AsEmailSubscription() (v PublicEmailSubscriptionFilter) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicOrFilterBranchFilterUnion) AsCommunicationSubscription() (v PublicCommunicationSubscriptionFilter) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicOrFilterBranchFilterUnion) AsCampaignInfluenced() (v PublicCampaignInfluencedFilter) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicOrFilterBranchFilterUnion) AsSurveyMonkey() (v PublicSurveyMonkeyFilter) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicOrFilterBranchFilterUnion) AsSurveyMonkeyValue() (v PublicSurveyMonkeyValueFilter) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicOrFilterBranchFilterUnion) AsWebinar() (v PublicWebinarFilter) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicOrFilterBranchFilterUnion) AsEmailEvent() (v PublicEmailEventFilter) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicOrFilterBranchFilterUnion) AsPrivacy() (v PublicPrivacyAnalyticsFilter) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicOrFilterBranchFilterUnion) AsAdsSearch() (v PublicAdsSearchFilter) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicOrFilterBranchFilterUnion) AsAdsTime() (v PublicAdsTimeFilter) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicOrFilterBranchFilterUnion) AsInList() (v PublicInListFilter) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicOrFilterBranchFilterUnion) AsNumAssociations() (v PublicNumAssociationsFilter) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicOrFilterBranchFilterUnion) AsUnifiedEvents() (v PublicUnifiedEventsFilter) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicOrFilterBranchFilterUnion) AsPropertyAssociation() (v PublicPropertyAssociationInListFilter) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicOrFilterBranchFilterUnion) AsConstant() (v PublicConstantFilter) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u PublicOrFilterBranchFilterUnion) RawJSON() string { return u.JSON.raw }

func (r *PublicOrFilterBranchFilterUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PublicOrFilterBranchFilterUnionCoalescingRefineBy is an implicit subunion of
// [PublicOrFilterBranchFilterUnion].
// PublicOrFilterBranchFilterUnionCoalescingRefineBy provides convenient access to
// the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [PublicOrFilterBranchFilterUnion].
type PublicOrFilterBranchFilterUnionCoalescingRefineBy struct {
	Type string `json:"type"`
	// This field is from variant
	// [PublicAssociationInListFilterCoalescingRefineByUnion],
	// [PublicPageViewAnalyticsFilterCoalescingRefineByUnion],
	// [PublicCtaAnalyticsFilterCoalescingRefineByUnion],
	// [PublicEventAnalyticsFilterCoalescingRefineByUnion],
	// [PublicFormSubmissionFilterCoalescingRefineByUnion],
	// [PublicFormSubmissionOnPageFilterCoalescingRefineByUnion],
	// [PublicNumAssociationsFilterCoalescingRefineByUnion],
	// [PublicUnifiedEventsFilterCoalescingRefineByUnion],
	// [PublicPropertyAssociationInListFilterCoalescingRefineByUnion].
	MaxOccurrences int64 `json:"maxOccurrences"`
	// This field is from variant
	// [PublicAssociationInListFilterCoalescingRefineByUnion],
	// [PublicPageViewAnalyticsFilterCoalescingRefineByUnion],
	// [PublicCtaAnalyticsFilterCoalescingRefineByUnion],
	// [PublicEventAnalyticsFilterCoalescingRefineByUnion],
	// [PublicFormSubmissionFilterCoalescingRefineByUnion],
	// [PublicFormSubmissionOnPageFilterCoalescingRefineByUnion],
	// [PublicNumAssociationsFilterCoalescingRefineByUnion],
	// [PublicUnifiedEventsFilterCoalescingRefineByUnion],
	// [PublicPropertyAssociationInListFilterCoalescingRefineByUnion].
	MinOccurrences int64 `json:"minOccurrences"`
	// This field is from variant
	// [PublicAssociationInListFilterCoalescingRefineByUnion],
	// [PublicPageViewAnalyticsFilterCoalescingRefineByUnion],
	// [PublicCtaAnalyticsFilterCoalescingRefineByUnion],
	// [PublicEventAnalyticsFilterCoalescingRefineByUnion],
	// [PublicFormSubmissionFilterCoalescingRefineByUnion],
	// [PublicFormSubmissionOnPageFilterCoalescingRefineByUnion],
	// [PublicNumAssociationsFilterCoalescingRefineByUnion],
	// [PublicUnifiedEventsFilterCoalescingRefineByUnion],
	// [PublicPropertyAssociationInListFilterCoalescingRefineByUnion].
	SetType    string `json:"setType"`
	Comparison string `json:"comparison"`
	// This field is from variant
	// [PublicAssociationInListFilterCoalescingRefineByUnion],
	// [PublicPageViewAnalyticsFilterCoalescingRefineByUnion],
	// [PublicCtaAnalyticsFilterCoalescingRefineByUnion],
	// [PublicEventAnalyticsFilterCoalescingRefineByUnion],
	// [PublicFormSubmissionFilterCoalescingRefineByUnion],
	// [PublicFormSubmissionOnPageFilterCoalescingRefineByUnion],
	// [PublicNumAssociationsFilterCoalescingRefineByUnion],
	// [PublicUnifiedEventsFilterCoalescingRefineByUnion],
	// [PublicPropertyAssociationInListFilterCoalescingRefineByUnion].
	TimeOffset PublicTimeOffset `json:"timeOffset"`
	// This field is from variant
	// [PublicAssociationInListFilterCoalescingRefineByUnion],
	// [PublicPageViewAnalyticsFilterCoalescingRefineByUnion],
	// [PublicCtaAnalyticsFilterCoalescingRefineByUnion],
	// [PublicEventAnalyticsFilterCoalescingRefineByUnion],
	// [PublicFormSubmissionFilterCoalescingRefineByUnion],
	// [PublicFormSubmissionOnPageFilterCoalescingRefineByUnion],
	// [PublicNumAssociationsFilterCoalescingRefineByUnion],
	// [PublicUnifiedEventsFilterCoalescingRefineByUnion],
	// [PublicPropertyAssociationInListFilterCoalescingRefineByUnion].
	LowerBoundOffset PublicTimeOffset `json:"lowerBoundOffset"`
	RangeType        string           `json:"rangeType"`
	// This field is from variant
	// [PublicAssociationInListFilterCoalescingRefineByUnion],
	// [PublicPageViewAnalyticsFilterCoalescingRefineByUnion],
	// [PublicCtaAnalyticsFilterCoalescingRefineByUnion],
	// [PublicEventAnalyticsFilterCoalescingRefineByUnion],
	// [PublicFormSubmissionFilterCoalescingRefineByUnion],
	// [PublicFormSubmissionOnPageFilterCoalescingRefineByUnion],
	// [PublicNumAssociationsFilterCoalescingRefineByUnion],
	// [PublicUnifiedEventsFilterCoalescingRefineByUnion],
	// [PublicPropertyAssociationInListFilterCoalescingRefineByUnion].
	UpperBoundOffset PublicTimeOffset `json:"upperBoundOffset"`
	// This field is from variant
	// [PublicAssociationInListFilterCoalescingRefineByUnion],
	// [PublicPageViewAnalyticsFilterCoalescingRefineByUnion],
	// [PublicCtaAnalyticsFilterCoalescingRefineByUnion],
	// [PublicEventAnalyticsFilterCoalescingRefineByUnion],
	// [PublicFormSubmissionFilterCoalescingRefineByUnion],
	// [PublicFormSubmissionOnPageFilterCoalescingRefineByUnion],
	// [PublicNumAssociationsFilterCoalescingRefineByUnion],
	// [PublicUnifiedEventsFilterCoalescingRefineByUnion],
	// [PublicPropertyAssociationInListFilterCoalescingRefineByUnion].
	Timestamp int64 `json:"timestamp"`
	// This field is from variant
	// [PublicAssociationInListFilterCoalescingRefineByUnion],
	// [PublicPageViewAnalyticsFilterCoalescingRefineByUnion],
	// [PublicCtaAnalyticsFilterCoalescingRefineByUnion],
	// [PublicEventAnalyticsFilterCoalescingRefineByUnion],
	// [PublicFormSubmissionFilterCoalescingRefineByUnion],
	// [PublicFormSubmissionOnPageFilterCoalescingRefineByUnion],
	// [PublicNumAssociationsFilterCoalescingRefineByUnion],
	// [PublicUnifiedEventsFilterCoalescingRefineByUnion],
	// [PublicPropertyAssociationInListFilterCoalescingRefineByUnion].
	LowerTimestamp int64 `json:"lowerTimestamp"`
	// This field is from variant
	// [PublicAssociationInListFilterCoalescingRefineByUnion],
	// [PublicPageViewAnalyticsFilterCoalescingRefineByUnion],
	// [PublicCtaAnalyticsFilterCoalescingRefineByUnion],
	// [PublicEventAnalyticsFilterCoalescingRefineByUnion],
	// [PublicFormSubmissionFilterCoalescingRefineByUnion],
	// [PublicFormSubmissionOnPageFilterCoalescingRefineByUnion],
	// [PublicNumAssociationsFilterCoalescingRefineByUnion],
	// [PublicUnifiedEventsFilterCoalescingRefineByUnion],
	// [PublicPropertyAssociationInListFilterCoalescingRefineByUnion].
	UpperTimestamp               int64  `json:"upperTimestamp"`
	IncludeObjectsWithNoValueSet bool   `json:"includeObjectsWithNoValueSet"`
	OperationType                string `json:"operationType"`
	Operator                     string `json:"operator"`
	// This field is from variant
	// [PublicAssociationInListFilterCoalescingRefineByUnion],
	// [PublicPageViewAnalyticsFilterCoalescingRefineByUnion],
	// [PublicCtaAnalyticsFilterCoalescingRefineByUnion],
	// [PublicEventAnalyticsFilterCoalescingRefineByUnion],
	// [PublicFormSubmissionFilterCoalescingRefineByUnion],
	// [PublicFormSubmissionOnPageFilterCoalescingRefineByUnion],
	// [PublicNumAssociationsFilterCoalescingRefineByUnion],
	// [PublicUnifiedEventsFilterCoalescingRefineByUnion],
	// [PublicPropertyAssociationInListFilterCoalescingRefineByUnion].
	TimePoint PublicTimePointOperationTimePointUnion `json:"timePoint"`
	// This field is from variant
	// [PublicAssociationInListFilterCoalescingRefineByUnion],
	// [PublicPageViewAnalyticsFilterCoalescingRefineByUnion],
	// [PublicCtaAnalyticsFilterCoalescingRefineByUnion],
	// [PublicEventAnalyticsFilterCoalescingRefineByUnion],
	// [PublicFormSubmissionFilterCoalescingRefineByUnion],
	// [PublicFormSubmissionOnPageFilterCoalescingRefineByUnion],
	// [PublicNumAssociationsFilterCoalescingRefineByUnion],
	// [PublicUnifiedEventsFilterCoalescingRefineByUnion],
	// [PublicPropertyAssociationInListFilterCoalescingRefineByUnion].
	EndpointBehavior string `json:"endpointBehavior"`
	PropertyParser   string `json:"propertyParser"`
	// This field is from variant
	// [PublicAssociationInListFilterCoalescingRefineByUnion],
	// [PublicPageViewAnalyticsFilterCoalescingRefineByUnion],
	// [PublicCtaAnalyticsFilterCoalescingRefineByUnion],
	// [PublicEventAnalyticsFilterCoalescingRefineByUnion],
	// [PublicFormSubmissionFilterCoalescingRefineByUnion],
	// [PublicFormSubmissionOnPageFilterCoalescingRefineByUnion],
	// [PublicNumAssociationsFilterCoalescingRefineByUnion],
	// [PublicUnifiedEventsFilterCoalescingRefineByUnion],
	// [PublicPropertyAssociationInListFilterCoalescingRefineByUnion].
	LowerBoundTimePoint PublicRangedTimeOperationLowerBoundTimePointUnion `json:"lowerBoundTimePoint"`
	// This field is from variant
	// [PublicAssociationInListFilterCoalescingRefineByUnion],
	// [PublicPageViewAnalyticsFilterCoalescingRefineByUnion],
	// [PublicCtaAnalyticsFilterCoalescingRefineByUnion],
	// [PublicEventAnalyticsFilterCoalescingRefineByUnion],
	// [PublicFormSubmissionFilterCoalescingRefineByUnion],
	// [PublicFormSubmissionOnPageFilterCoalescingRefineByUnion],
	// [PublicNumAssociationsFilterCoalescingRefineByUnion],
	// [PublicUnifiedEventsFilterCoalescingRefineByUnion],
	// [PublicPropertyAssociationInListFilterCoalescingRefineByUnion].
	UpperBoundTimePoint PublicRangedTimeOperationUpperBoundTimePointUnion `json:"upperBoundTimePoint"`
	// This field is from variant
	// [PublicAssociationInListFilterCoalescingRefineByUnion],
	// [PublicPageViewAnalyticsFilterCoalescingRefineByUnion],
	// [PublicCtaAnalyticsFilterCoalescingRefineByUnion],
	// [PublicEventAnalyticsFilterCoalescingRefineByUnion],
	// [PublicFormSubmissionFilterCoalescingRefineByUnion],
	// [PublicFormSubmissionOnPageFilterCoalescingRefineByUnion],
	// [PublicNumAssociationsFilterCoalescingRefineByUnion],
	// [PublicUnifiedEventsFilterCoalescingRefineByUnion],
	// [PublicPropertyAssociationInListFilterCoalescingRefineByUnion].
	LowerBoundEndpointBehavior string `json:"lowerBoundEndpointBehavior"`
	// This field is from variant
	// [PublicAssociationInListFilterCoalescingRefineByUnion],
	// [PublicPageViewAnalyticsFilterCoalescingRefineByUnion],
	// [PublicCtaAnalyticsFilterCoalescingRefineByUnion],
	// [PublicEventAnalyticsFilterCoalescingRefineByUnion],
	// [PublicFormSubmissionFilterCoalescingRefineByUnion],
	// [PublicFormSubmissionOnPageFilterCoalescingRefineByUnion],
	// [PublicNumAssociationsFilterCoalescingRefineByUnion],
	// [PublicUnifiedEventsFilterCoalescingRefineByUnion],
	// [PublicPropertyAssociationInListFilterCoalescingRefineByUnion].
	UpperBoundEndpointBehavior string `json:"upperBoundEndpointBehavior"`
	JSON                       struct {
		Type                         respjson.Field
		MaxOccurrences               respjson.Field
		MinOccurrences               respjson.Field
		SetType                      respjson.Field
		Comparison                   respjson.Field
		TimeOffset                   respjson.Field
		LowerBoundOffset             respjson.Field
		RangeType                    respjson.Field
		UpperBoundOffset             respjson.Field
		Timestamp                    respjson.Field
		LowerTimestamp               respjson.Field
		UpperTimestamp               respjson.Field
		IncludeObjectsWithNoValueSet respjson.Field
		OperationType                respjson.Field
		Operator                     respjson.Field
		TimePoint                    respjson.Field
		EndpointBehavior             respjson.Field
		PropertyParser               respjson.Field
		LowerBoundTimePoint          respjson.Field
		UpperBoundTimePoint          respjson.Field
		LowerBoundEndpointBehavior   respjson.Field
		UpperBoundEndpointBehavior   respjson.Field
		raw                          string
	} `json:"-"`
}

func (r *PublicOrFilterBranchFilterUnionCoalescingRefineBy) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PublicOrFilterBranchFilterUnionPruningRefineBy is an implicit subunion of
// [PublicOrFilterBranchFilterUnion].
// PublicOrFilterBranchFilterUnionPruningRefineBy provides convenient access to the
// sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [PublicOrFilterBranchFilterUnion].
type PublicOrFilterBranchFilterUnionPruningRefineBy struct {
	Type string `json:"type"`
	// This field is from variant [PublicPageViewAnalyticsFilterPruningRefineByUnion],
	// [PublicCtaAnalyticsFilterPruningRefineByUnion],
	// [PublicEventAnalyticsFilterPruningRefineByUnion],
	// [PublicFormSubmissionFilterPruningRefineByUnion],
	// [PublicFormSubmissionOnPageFilterPruningRefineByUnion],
	// [PublicEmailEventFilterPruningRefineByUnion],
	// [PublicAdsTimeFilterPruningRefineByUnion],
	// [PublicUnifiedEventsFilterPruningRefineByUnion].
	MaxOccurrences int64 `json:"maxOccurrences"`
	// This field is from variant [PublicPageViewAnalyticsFilterPruningRefineByUnion],
	// [PublicCtaAnalyticsFilterPruningRefineByUnion],
	// [PublicEventAnalyticsFilterPruningRefineByUnion],
	// [PublicFormSubmissionFilterPruningRefineByUnion],
	// [PublicFormSubmissionOnPageFilterPruningRefineByUnion],
	// [PublicEmailEventFilterPruningRefineByUnion],
	// [PublicAdsTimeFilterPruningRefineByUnion],
	// [PublicUnifiedEventsFilterPruningRefineByUnion].
	MinOccurrences int64 `json:"minOccurrences"`
	// This field is from variant [PublicPageViewAnalyticsFilterPruningRefineByUnion],
	// [PublicCtaAnalyticsFilterPruningRefineByUnion],
	// [PublicEventAnalyticsFilterPruningRefineByUnion],
	// [PublicFormSubmissionFilterPruningRefineByUnion],
	// [PublicFormSubmissionOnPageFilterPruningRefineByUnion],
	// [PublicEmailEventFilterPruningRefineByUnion],
	// [PublicAdsTimeFilterPruningRefineByUnion],
	// [PublicUnifiedEventsFilterPruningRefineByUnion].
	SetType    string `json:"setType"`
	Comparison string `json:"comparison"`
	// This field is from variant [PublicPageViewAnalyticsFilterPruningRefineByUnion],
	// [PublicCtaAnalyticsFilterPruningRefineByUnion],
	// [PublicEventAnalyticsFilterPruningRefineByUnion],
	// [PublicFormSubmissionFilterPruningRefineByUnion],
	// [PublicFormSubmissionOnPageFilterPruningRefineByUnion],
	// [PublicEmailEventFilterPruningRefineByUnion],
	// [PublicAdsTimeFilterPruningRefineByUnion],
	// [PublicUnifiedEventsFilterPruningRefineByUnion].
	TimeOffset PublicTimeOffset `json:"timeOffset"`
	// This field is from variant [PublicPageViewAnalyticsFilterPruningRefineByUnion],
	// [PublicCtaAnalyticsFilterPruningRefineByUnion],
	// [PublicEventAnalyticsFilterPruningRefineByUnion],
	// [PublicFormSubmissionFilterPruningRefineByUnion],
	// [PublicFormSubmissionOnPageFilterPruningRefineByUnion],
	// [PublicEmailEventFilterPruningRefineByUnion],
	// [PublicAdsTimeFilterPruningRefineByUnion],
	// [PublicUnifiedEventsFilterPruningRefineByUnion].
	LowerBoundOffset PublicTimeOffset `json:"lowerBoundOffset"`
	RangeType        string           `json:"rangeType"`
	// This field is from variant [PublicPageViewAnalyticsFilterPruningRefineByUnion],
	// [PublicCtaAnalyticsFilterPruningRefineByUnion],
	// [PublicEventAnalyticsFilterPruningRefineByUnion],
	// [PublicFormSubmissionFilterPruningRefineByUnion],
	// [PublicFormSubmissionOnPageFilterPruningRefineByUnion],
	// [PublicEmailEventFilterPruningRefineByUnion],
	// [PublicAdsTimeFilterPruningRefineByUnion],
	// [PublicUnifiedEventsFilterPruningRefineByUnion].
	UpperBoundOffset PublicTimeOffset `json:"upperBoundOffset"`
	// This field is from variant [PublicPageViewAnalyticsFilterPruningRefineByUnion],
	// [PublicCtaAnalyticsFilterPruningRefineByUnion],
	// [PublicEventAnalyticsFilterPruningRefineByUnion],
	// [PublicFormSubmissionFilterPruningRefineByUnion],
	// [PublicFormSubmissionOnPageFilterPruningRefineByUnion],
	// [PublicEmailEventFilterPruningRefineByUnion],
	// [PublicAdsTimeFilterPruningRefineByUnion],
	// [PublicUnifiedEventsFilterPruningRefineByUnion].
	Timestamp int64 `json:"timestamp"`
	// This field is from variant [PublicPageViewAnalyticsFilterPruningRefineByUnion],
	// [PublicCtaAnalyticsFilterPruningRefineByUnion],
	// [PublicEventAnalyticsFilterPruningRefineByUnion],
	// [PublicFormSubmissionFilterPruningRefineByUnion],
	// [PublicFormSubmissionOnPageFilterPruningRefineByUnion],
	// [PublicEmailEventFilterPruningRefineByUnion],
	// [PublicAdsTimeFilterPruningRefineByUnion],
	// [PublicUnifiedEventsFilterPruningRefineByUnion].
	LowerTimestamp int64 `json:"lowerTimestamp"`
	// This field is from variant [PublicPageViewAnalyticsFilterPruningRefineByUnion],
	// [PublicCtaAnalyticsFilterPruningRefineByUnion],
	// [PublicEventAnalyticsFilterPruningRefineByUnion],
	// [PublicFormSubmissionFilterPruningRefineByUnion],
	// [PublicFormSubmissionOnPageFilterPruningRefineByUnion],
	// [PublicEmailEventFilterPruningRefineByUnion],
	// [PublicAdsTimeFilterPruningRefineByUnion],
	// [PublicUnifiedEventsFilterPruningRefineByUnion].
	UpperTimestamp               int64  `json:"upperTimestamp"`
	IncludeObjectsWithNoValueSet bool   `json:"includeObjectsWithNoValueSet"`
	OperationType                string `json:"operationType"`
	Operator                     string `json:"operator"`
	// This field is from variant [PublicPageViewAnalyticsFilterPruningRefineByUnion],
	// [PublicCtaAnalyticsFilterPruningRefineByUnion],
	// [PublicEventAnalyticsFilterPruningRefineByUnion],
	// [PublicFormSubmissionFilterPruningRefineByUnion],
	// [PublicFormSubmissionOnPageFilterPruningRefineByUnion],
	// [PublicEmailEventFilterPruningRefineByUnion],
	// [PublicAdsTimeFilterPruningRefineByUnion],
	// [PublicUnifiedEventsFilterPruningRefineByUnion].
	TimePoint PublicTimePointOperationTimePointUnion `json:"timePoint"`
	// This field is from variant [PublicPageViewAnalyticsFilterPruningRefineByUnion],
	// [PublicCtaAnalyticsFilterPruningRefineByUnion],
	// [PublicEventAnalyticsFilterPruningRefineByUnion],
	// [PublicFormSubmissionFilterPruningRefineByUnion],
	// [PublicFormSubmissionOnPageFilterPruningRefineByUnion],
	// [PublicEmailEventFilterPruningRefineByUnion],
	// [PublicAdsTimeFilterPruningRefineByUnion],
	// [PublicUnifiedEventsFilterPruningRefineByUnion].
	EndpointBehavior string `json:"endpointBehavior"`
	PropertyParser   string `json:"propertyParser"`
	// This field is from variant [PublicPageViewAnalyticsFilterPruningRefineByUnion],
	// [PublicCtaAnalyticsFilterPruningRefineByUnion],
	// [PublicEventAnalyticsFilterPruningRefineByUnion],
	// [PublicFormSubmissionFilterPruningRefineByUnion],
	// [PublicFormSubmissionOnPageFilterPruningRefineByUnion],
	// [PublicEmailEventFilterPruningRefineByUnion],
	// [PublicAdsTimeFilterPruningRefineByUnion],
	// [PublicUnifiedEventsFilterPruningRefineByUnion].
	LowerBoundTimePoint PublicRangedTimeOperationLowerBoundTimePointUnion `json:"lowerBoundTimePoint"`
	// This field is from variant [PublicPageViewAnalyticsFilterPruningRefineByUnion],
	// [PublicCtaAnalyticsFilterPruningRefineByUnion],
	// [PublicEventAnalyticsFilterPruningRefineByUnion],
	// [PublicFormSubmissionFilterPruningRefineByUnion],
	// [PublicFormSubmissionOnPageFilterPruningRefineByUnion],
	// [PublicEmailEventFilterPruningRefineByUnion],
	// [PublicAdsTimeFilterPruningRefineByUnion],
	// [PublicUnifiedEventsFilterPruningRefineByUnion].
	UpperBoundTimePoint PublicRangedTimeOperationUpperBoundTimePointUnion `json:"upperBoundTimePoint"`
	// This field is from variant [PublicPageViewAnalyticsFilterPruningRefineByUnion],
	// [PublicCtaAnalyticsFilterPruningRefineByUnion],
	// [PublicEventAnalyticsFilterPruningRefineByUnion],
	// [PublicFormSubmissionFilterPruningRefineByUnion],
	// [PublicFormSubmissionOnPageFilterPruningRefineByUnion],
	// [PublicEmailEventFilterPruningRefineByUnion],
	// [PublicAdsTimeFilterPruningRefineByUnion],
	// [PublicUnifiedEventsFilterPruningRefineByUnion].
	LowerBoundEndpointBehavior string `json:"lowerBoundEndpointBehavior"`
	// This field is from variant [PublicPageViewAnalyticsFilterPruningRefineByUnion],
	// [PublicCtaAnalyticsFilterPruningRefineByUnion],
	// [PublicEventAnalyticsFilterPruningRefineByUnion],
	// [PublicFormSubmissionFilterPruningRefineByUnion],
	// [PublicFormSubmissionOnPageFilterPruningRefineByUnion],
	// [PublicEmailEventFilterPruningRefineByUnion],
	// [PublicAdsTimeFilterPruningRefineByUnion],
	// [PublicUnifiedEventsFilterPruningRefineByUnion].
	UpperBoundEndpointBehavior string `json:"upperBoundEndpointBehavior"`
	JSON                       struct {
		Type                         respjson.Field
		MaxOccurrences               respjson.Field
		MinOccurrences               respjson.Field
		SetType                      respjson.Field
		Comparison                   respjson.Field
		TimeOffset                   respjson.Field
		LowerBoundOffset             respjson.Field
		RangeType                    respjson.Field
		UpperBoundOffset             respjson.Field
		Timestamp                    respjson.Field
		LowerTimestamp               respjson.Field
		UpperTimestamp               respjson.Field
		IncludeObjectsWithNoValueSet respjson.Field
		OperationType                respjson.Field
		Operator                     respjson.Field
		TimePoint                    respjson.Field
		EndpointBehavior             respjson.Field
		PropertyParser               respjson.Field
		LowerBoundTimePoint          respjson.Field
		UpperBoundTimePoint          respjson.Field
		LowerBoundEndpointBehavior   respjson.Field
		UpperBoundEndpointBehavior   respjson.Field
		raw                          string
	} `json:"-"`
}

func (r *PublicOrFilterBranchFilterUnionPruningRefineBy) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PublicOrFilterBranchFilterUnionEventTypeID is an implicit subunion of
// [PublicOrFilterBranchFilterUnion]. PublicOrFilterBranchFilterUnionEventTypeID
// provides convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [PublicOrFilterBranchFilterUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfInt OfString]
type PublicOrFilterBranchFilterUnionEventTypeID struct {
	// This field will be present if the value is a [int64] instead of an object.
	OfInt int64 `json:",inline"`
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	JSON     struct {
		OfInt    respjson.Field
		OfString respjson.Field
		raw      string
	} `json:"-"`
}

func (r *PublicOrFilterBranchFilterUnionEventTypeID) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties FilterBranches, FilterBranchOperator, FilterBranchType, Filters
// are required.
type PublicOrFilterBranchParam struct {
	FilterBranches []PublicOrFilterBranchFilterBranchUnionParam `json:"filterBranches,omitzero" api:"required"`
	// The logical operator used to combine the filters within the branch (OR).
	FilterBranchOperator string `json:"filterBranchOperator" api:"required"`
	// The type of the filter branch (OR).
	//
	// Any of "OR".
	FilterBranchType PublicOrFilterBranchFilterBranchType   `json:"filterBranchType,omitzero" api:"required"`
	Filters          []PublicOrFilterBranchFilterUnionParam `json:"filters,omitzero" api:"required"`
	paramObj
}

func (r PublicOrFilterBranchParam) MarshalJSON() (data []byte, err error) {
	type shadow PublicOrFilterBranchParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PublicOrFilterBranchParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type PublicOrFilterBranchFilterBranchUnionParam struct {
	OfOr                  *PublicOrFilterBranchParam                  `json:",omitzero,inline"`
	OfAnd                 *PublicAndFilterBranchParam                 `json:",omitzero,inline"`
	OfNotAll              *PublicNotAllFilterBranchParam              `json:",omitzero,inline"`
	OfNotAny              *PublicNotAnyFilterBranchParam              `json:",omitzero,inline"`
	OfRestricted          *PublicRestrictedFilterBranchParam          `json:",omitzero,inline"`
	OfUnifiedEvents       *PublicUnifiedEventsFilterBranchParam       `json:",omitzero,inline"`
	OfPropertyAssociation *PublicPropertyAssociationFilterBranchParam `json:",omitzero,inline"`
	OfAssociation         *PublicAssociationFilterBranchParam         `json:",omitzero,inline"`
	paramUnion
}

func (u PublicOrFilterBranchFilterBranchUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfOr,
		u.OfAnd,
		u.OfNotAll,
		u.OfNotAny,
		u.OfRestricted,
		u.OfUnifiedEvents,
		u.OfPropertyAssociation,
		u.OfAssociation)
}
func (u *PublicOrFilterBranchFilterBranchUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func init() {
	apijson.RegisterUnion[PublicOrFilterBranchFilterBranchUnionParam](
		"filterBranchType",
		apijson.Discriminator[PublicOrFilterBranchParam]("OR"),
		apijson.Discriminator[PublicAndFilterBranchParam]("AND"),
		apijson.Discriminator[PublicNotAllFilterBranchParam]("NOT_ALL"),
		apijson.Discriminator[PublicNotAnyFilterBranchParam]("NOT_ANY"),
		apijson.Discriminator[PublicRestrictedFilterBranchParam]("RESTRICTED"),
		apijson.Discriminator[PublicUnifiedEventsFilterBranchParam]("UNIFIED_EVENTS"),
		apijson.Discriminator[PublicPropertyAssociationFilterBranchParam]("PROPERTY_ASSOCIATION"),
		apijson.Discriminator[PublicAssociationFilterBranchParam]("ASSOCIATION"),
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type PublicOrFilterBranchFilterUnionParam struct {
	OfProperty                  *PublicPropertyFilterParam                  `json:",omitzero,inline"`
	OfAssociation               *PublicAssociationInListFilterParam         `json:",omitzero,inline"`
	OfPageView                  *PublicPageViewAnalyticsFilterParam         `json:",omitzero,inline"`
	OfCta                       *PublicCtaAnalyticsFilterParam              `json:",omitzero,inline"`
	OfEvent                     *PublicEventAnalyticsFilterParam            `json:",omitzero,inline"`
	OfFormSubmission            *PublicFormSubmissionFilterParam            `json:",omitzero,inline"`
	OfFormSubmissionOnPage      *PublicFormSubmissionOnPageFilterParam      `json:",omitzero,inline"`
	OfIntegrationEvent          *PublicIntegrationEventFilterParam          `json:",omitzero,inline"`
	OfEmailSubscription         *PublicEmailSubscriptionFilterParam         `json:",omitzero,inline"`
	OfCommunicationSubscription *PublicCommunicationSubscriptionFilterParam `json:",omitzero,inline"`
	OfCampaignInfluenced        *PublicCampaignInfluencedFilterParam        `json:",omitzero,inline"`
	OfSurveyMonkey              *PublicSurveyMonkeyFilterParam              `json:",omitzero,inline"`
	OfSurveyMonkeyValue         *PublicSurveyMonkeyValueFilterParam         `json:",omitzero,inline"`
	OfWebinar                   *PublicWebinarFilterParam                   `json:",omitzero,inline"`
	OfEmailEvent                *PublicEmailEventFilterParam                `json:",omitzero,inline"`
	OfPrivacy                   *PublicPrivacyAnalyticsFilterParam          `json:",omitzero,inline"`
	OfAdsSearch                 *PublicAdsSearchFilterParam                 `json:",omitzero,inline"`
	OfAdsTime                   *PublicAdsTimeFilterParam                   `json:",omitzero,inline"`
	OfInList                    *PublicInListFilterParam                    `json:",omitzero,inline"`
	OfNumAssociations           *PublicNumAssociationsFilterParam           `json:",omitzero,inline"`
	OfUnifiedEvents             *PublicUnifiedEventsFilterParam             `json:",omitzero,inline"`
	OfPropertyAssociation       *PublicPropertyAssociationInListFilterParam `json:",omitzero,inline"`
	OfConstant                  *PublicConstantFilterParam                  `json:",omitzero,inline"`
	paramUnion
}

func (u PublicOrFilterBranchFilterUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfProperty,
		u.OfAssociation,
		u.OfPageView,
		u.OfCta,
		u.OfEvent,
		u.OfFormSubmission,
		u.OfFormSubmissionOnPage,
		u.OfIntegrationEvent,
		u.OfEmailSubscription,
		u.OfCommunicationSubscription,
		u.OfCampaignInfluenced,
		u.OfSurveyMonkey,
		u.OfSurveyMonkeyValue,
		u.OfWebinar,
		u.OfEmailEvent,
		u.OfPrivacy,
		u.OfAdsSearch,
		u.OfAdsTime,
		u.OfInList,
		u.OfNumAssociations,
		u.OfUnifiedEvents,
		u.OfPropertyAssociation,
		u.OfConstant)
}
func (u *PublicOrFilterBranchFilterUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func init() {
	apijson.RegisterUnion[PublicOrFilterBranchFilterUnionParam](
		"filterType",
		apijson.Discriminator[PublicPropertyFilterParam]("PROPERTY"),
		apijson.Discriminator[PublicAssociationInListFilterParam]("ASSOCIATION"),
		apijson.Discriminator[PublicPageViewAnalyticsFilterParam]("PAGE_VIEW"),
		apijson.Discriminator[PublicCtaAnalyticsFilterParam]("CTA"),
		apijson.Discriminator[PublicEventAnalyticsFilterParam]("EVENT"),
		apijson.Discriminator[PublicFormSubmissionFilterParam]("FORM_SUBMISSION"),
		apijson.Discriminator[PublicFormSubmissionOnPageFilterParam]("FORM_SUBMISSION_ON_PAGE"),
		apijson.Discriminator[PublicIntegrationEventFilterParam]("INTEGRATION_EVENT"),
		apijson.Discriminator[PublicEmailSubscriptionFilterParam]("EMAIL_SUBSCRIPTION"),
		apijson.Discriminator[PublicCommunicationSubscriptionFilterParam]("COMMUNICATION_SUBSCRIPTION"),
		apijson.Discriminator[PublicCampaignInfluencedFilterParam]("CAMPAIGN_INFLUENCED"),
		apijson.Discriminator[PublicSurveyMonkeyFilterParam]("SURVEY_MONKEY"),
		apijson.Discriminator[PublicSurveyMonkeyValueFilterParam]("SURVEY_MONKEY_VALUE"),
		apijson.Discriminator[PublicWebinarFilterParam]("WEBINAR"),
		apijson.Discriminator[PublicEmailEventFilterParam]("EMAIL_EVENT"),
		apijson.Discriminator[PublicPrivacyAnalyticsFilterParam]("PRIVACY"),
		apijson.Discriminator[PublicAdsSearchFilterParam]("ADS_SEARCH"),
		apijson.Discriminator[PublicAdsTimeFilterParam]("ADS_TIME"),
		apijson.Discriminator[PublicInListFilterParam]("IN_LIST"),
		apijson.Discriminator[PublicNumAssociationsFilterParam]("NUM_ASSOCIATIONS"),
		apijson.Discriminator[PublicUnifiedEventsFilterParam]("UNIFIED_EVENTS"),
		apijson.Discriminator[PublicPropertyAssociationInListFilterParam]("PROPERTY_ASSOCIATION"),
		apijson.Discriminator[PublicConstantFilterParam]("CONSTANT"),
	)
}

type PublicPageViewAnalyticsFilter struct {
	// Indicates the type of filter being applied (PAGE_VIEW).
	//
	// Any of "PAGE_VIEW".
	FilterType PublicPageViewAnalyticsFilterFilterType `json:"filterType" api:"required"`
	// Defines the operation to be applied within the filter (HAS_PAGEVIEW_EQ,
	// HAS_PAGEVIEW_CONTAINS, HAS_PAGEVIEW_MATCHES_REGEX, NOT_HAS_PAGEVIEW_EQ,
	// NOT_HAS_PAGEVIEW_CONTAINS).
	Operator string `json:"operator" api:"required"`
	// The URL of the page to be used in the filter.
	PageURL string `json:"pageUrl" api:"required"`
	// Specifies the criteria for refining the filter by coalescing.
	CoalescingRefineBy PublicPageViewAnalyticsFilterCoalescingRefineByUnion `json:"coalescingRefineBy"`
	// Indicates whether tracking is enabled for the page view.
	EnableTracking bool `json:"enableTracking"`
	// Specifies the criteria for refining the filter by pruning.
	PruningRefineBy PublicPageViewAnalyticsFilterPruningRefineByUnion `json:"pruningRefineBy"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FilterType         respjson.Field
		Operator           respjson.Field
		PageURL            respjson.Field
		CoalescingRefineBy respjson.Field
		EnableTracking     respjson.Field
		PruningRefineBy    respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PublicPageViewAnalyticsFilter) RawJSON() string { return r.JSON.raw }
func (r *PublicPageViewAnalyticsFilter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this PublicPageViewAnalyticsFilter to a
// PublicPageViewAnalyticsFilterParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// PublicPageViewAnalyticsFilterParam.Overrides()
func (r PublicPageViewAnalyticsFilter) ToParam() PublicPageViewAnalyticsFilterParam {
	return param.Override[PublicPageViewAnalyticsFilterParam](json.RawMessage(r.RawJSON()))
}

// Indicates the type of filter being applied (PAGE_VIEW).
type PublicPageViewAnalyticsFilterFilterType string

const (
	PublicPageViewAnalyticsFilterFilterTypePageView PublicPageViewAnalyticsFilterFilterType = "PAGE_VIEW"
)

// PublicPageViewAnalyticsFilterCoalescingRefineByUnion contains all possible
// properties and values from [PublicNumOccurrencesRefineBy],
// [PublicSetOccurrencesRefineBy], [PublicRelativeComparativeTimestampRefineBy],
// [PublicRelativeRangedTimestampRefineBy],
// [PublicAbsoluteComparativeTimestampRefineBy],
// [PublicAbsoluteRangedTimestampRefineBy], [PublicAllHistoryRefineBy],
// [PublicTimePointOperation], [PublicRangedTimeOperation].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type PublicPageViewAnalyticsFilterCoalescingRefineByUnion struct {
	Type string `json:"type"`
	// This field is from variant [PublicNumOccurrencesRefineBy].
	MaxOccurrences int64 `json:"maxOccurrences"`
	// This field is from variant [PublicNumOccurrencesRefineBy].
	MinOccurrences int64 `json:"minOccurrences"`
	// This field is from variant [PublicSetOccurrencesRefineBy].
	SetType    string `json:"setType"`
	Comparison string `json:"comparison"`
	// This field is from variant [PublicRelativeComparativeTimestampRefineBy].
	TimeOffset PublicTimeOffset `json:"timeOffset"`
	// This field is from variant [PublicRelativeRangedTimestampRefineBy].
	LowerBoundOffset PublicTimeOffset `json:"lowerBoundOffset"`
	RangeType        string           `json:"rangeType"`
	// This field is from variant [PublicRelativeRangedTimestampRefineBy].
	UpperBoundOffset PublicTimeOffset `json:"upperBoundOffset"`
	// This field is from variant [PublicAbsoluteComparativeTimestampRefineBy].
	Timestamp int64 `json:"timestamp"`
	// This field is from variant [PublicAbsoluteRangedTimestampRefineBy].
	LowerTimestamp int64 `json:"lowerTimestamp"`
	// This field is from variant [PublicAbsoluteRangedTimestampRefineBy].
	UpperTimestamp               int64  `json:"upperTimestamp"`
	IncludeObjectsWithNoValueSet bool   `json:"includeObjectsWithNoValueSet"`
	OperationType                string `json:"operationType"`
	Operator                     string `json:"operator"`
	// This field is from variant [PublicTimePointOperation].
	TimePoint PublicTimePointOperationTimePointUnion `json:"timePoint"`
	// This field is from variant [PublicTimePointOperation].
	EndpointBehavior string `json:"endpointBehavior"`
	PropertyParser   string `json:"propertyParser"`
	// This field is from variant [PublicRangedTimeOperation].
	LowerBoundTimePoint PublicRangedTimeOperationLowerBoundTimePointUnion `json:"lowerBoundTimePoint"`
	// This field is from variant [PublicRangedTimeOperation].
	UpperBoundTimePoint PublicRangedTimeOperationUpperBoundTimePointUnion `json:"upperBoundTimePoint"`
	// This field is from variant [PublicRangedTimeOperation].
	LowerBoundEndpointBehavior string `json:"lowerBoundEndpointBehavior"`
	// This field is from variant [PublicRangedTimeOperation].
	UpperBoundEndpointBehavior string `json:"upperBoundEndpointBehavior"`
	JSON                       struct {
		Type                         respjson.Field
		MaxOccurrences               respjson.Field
		MinOccurrences               respjson.Field
		SetType                      respjson.Field
		Comparison                   respjson.Field
		TimeOffset                   respjson.Field
		LowerBoundOffset             respjson.Field
		RangeType                    respjson.Field
		UpperBoundOffset             respjson.Field
		Timestamp                    respjson.Field
		LowerTimestamp               respjson.Field
		UpperTimestamp               respjson.Field
		IncludeObjectsWithNoValueSet respjson.Field
		OperationType                respjson.Field
		Operator                     respjson.Field
		TimePoint                    respjson.Field
		EndpointBehavior             respjson.Field
		PropertyParser               respjson.Field
		LowerBoundTimePoint          respjson.Field
		UpperBoundTimePoint          respjson.Field
		LowerBoundEndpointBehavior   respjson.Field
		UpperBoundEndpointBehavior   respjson.Field
		raw                          string
	} `json:"-"`
}

func (u PublicPageViewAnalyticsFilterCoalescingRefineByUnion) AsNumOccurrences() (v PublicNumOccurrencesRefineBy) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicPageViewAnalyticsFilterCoalescingRefineByUnion) AsSetOccurrences() (v PublicSetOccurrencesRefineBy) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicPageViewAnalyticsFilterCoalescingRefineByUnion) AsRelativeComparative() (v PublicRelativeComparativeTimestampRefineBy) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicPageViewAnalyticsFilterCoalescingRefineByUnion) AsRelativeRanged() (v PublicRelativeRangedTimestampRefineBy) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicPageViewAnalyticsFilterCoalescingRefineByUnion) AsAbsoluteComparative() (v PublicAbsoluteComparativeTimestampRefineBy) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicPageViewAnalyticsFilterCoalescingRefineByUnion) AsAbsoluteRanged() (v PublicAbsoluteRangedTimestampRefineBy) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicPageViewAnalyticsFilterCoalescingRefineByUnion) AsAllHistory() (v PublicAllHistoryRefineBy) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicPageViewAnalyticsFilterCoalescingRefineByUnion) AsTimePoint() (v PublicTimePointOperation) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicPageViewAnalyticsFilterCoalescingRefineByUnion) AsTimeRanged() (v PublicRangedTimeOperation) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u PublicPageViewAnalyticsFilterCoalescingRefineByUnion) RawJSON() string { return u.JSON.raw }

func (r *PublicPageViewAnalyticsFilterCoalescingRefineByUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PublicPageViewAnalyticsFilterPruningRefineByUnion contains all possible
// properties and values from [PublicNumOccurrencesRefineBy],
// [PublicSetOccurrencesRefineBy], [PublicRelativeComparativeTimestampRefineBy],
// [PublicRelativeRangedTimestampRefineBy],
// [PublicAbsoluteComparativeTimestampRefineBy],
// [PublicAbsoluteRangedTimestampRefineBy], [PublicAllHistoryRefineBy],
// [PublicTimePointOperation], [PublicRangedTimeOperation].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type PublicPageViewAnalyticsFilterPruningRefineByUnion struct {
	Type string `json:"type"`
	// This field is from variant [PublicNumOccurrencesRefineBy].
	MaxOccurrences int64 `json:"maxOccurrences"`
	// This field is from variant [PublicNumOccurrencesRefineBy].
	MinOccurrences int64 `json:"minOccurrences"`
	// This field is from variant [PublicSetOccurrencesRefineBy].
	SetType    string `json:"setType"`
	Comparison string `json:"comparison"`
	// This field is from variant [PublicRelativeComparativeTimestampRefineBy].
	TimeOffset PublicTimeOffset `json:"timeOffset"`
	// This field is from variant [PublicRelativeRangedTimestampRefineBy].
	LowerBoundOffset PublicTimeOffset `json:"lowerBoundOffset"`
	RangeType        string           `json:"rangeType"`
	// This field is from variant [PublicRelativeRangedTimestampRefineBy].
	UpperBoundOffset PublicTimeOffset `json:"upperBoundOffset"`
	// This field is from variant [PublicAbsoluteComparativeTimestampRefineBy].
	Timestamp int64 `json:"timestamp"`
	// This field is from variant [PublicAbsoluteRangedTimestampRefineBy].
	LowerTimestamp int64 `json:"lowerTimestamp"`
	// This field is from variant [PublicAbsoluteRangedTimestampRefineBy].
	UpperTimestamp               int64  `json:"upperTimestamp"`
	IncludeObjectsWithNoValueSet bool   `json:"includeObjectsWithNoValueSet"`
	OperationType                string `json:"operationType"`
	Operator                     string `json:"operator"`
	// This field is from variant [PublicTimePointOperation].
	TimePoint PublicTimePointOperationTimePointUnion `json:"timePoint"`
	// This field is from variant [PublicTimePointOperation].
	EndpointBehavior string `json:"endpointBehavior"`
	PropertyParser   string `json:"propertyParser"`
	// This field is from variant [PublicRangedTimeOperation].
	LowerBoundTimePoint PublicRangedTimeOperationLowerBoundTimePointUnion `json:"lowerBoundTimePoint"`
	// This field is from variant [PublicRangedTimeOperation].
	UpperBoundTimePoint PublicRangedTimeOperationUpperBoundTimePointUnion `json:"upperBoundTimePoint"`
	// This field is from variant [PublicRangedTimeOperation].
	LowerBoundEndpointBehavior string `json:"lowerBoundEndpointBehavior"`
	// This field is from variant [PublicRangedTimeOperation].
	UpperBoundEndpointBehavior string `json:"upperBoundEndpointBehavior"`
	JSON                       struct {
		Type                         respjson.Field
		MaxOccurrences               respjson.Field
		MinOccurrences               respjson.Field
		SetType                      respjson.Field
		Comparison                   respjson.Field
		TimeOffset                   respjson.Field
		LowerBoundOffset             respjson.Field
		RangeType                    respjson.Field
		UpperBoundOffset             respjson.Field
		Timestamp                    respjson.Field
		LowerTimestamp               respjson.Field
		UpperTimestamp               respjson.Field
		IncludeObjectsWithNoValueSet respjson.Field
		OperationType                respjson.Field
		Operator                     respjson.Field
		TimePoint                    respjson.Field
		EndpointBehavior             respjson.Field
		PropertyParser               respjson.Field
		LowerBoundTimePoint          respjson.Field
		UpperBoundTimePoint          respjson.Field
		LowerBoundEndpointBehavior   respjson.Field
		UpperBoundEndpointBehavior   respjson.Field
		raw                          string
	} `json:"-"`
}

func (u PublicPageViewAnalyticsFilterPruningRefineByUnion) AsNumOccurrences() (v PublicNumOccurrencesRefineBy) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicPageViewAnalyticsFilterPruningRefineByUnion) AsSetOccurrences() (v PublicSetOccurrencesRefineBy) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicPageViewAnalyticsFilterPruningRefineByUnion) AsRelativeComparative() (v PublicRelativeComparativeTimestampRefineBy) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicPageViewAnalyticsFilterPruningRefineByUnion) AsRelativeRanged() (v PublicRelativeRangedTimestampRefineBy) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicPageViewAnalyticsFilterPruningRefineByUnion) AsAbsoluteComparative() (v PublicAbsoluteComparativeTimestampRefineBy) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicPageViewAnalyticsFilterPruningRefineByUnion) AsAbsoluteRanged() (v PublicAbsoluteRangedTimestampRefineBy) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicPageViewAnalyticsFilterPruningRefineByUnion) AsAllHistory() (v PublicAllHistoryRefineBy) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicPageViewAnalyticsFilterPruningRefineByUnion) AsTimePoint() (v PublicTimePointOperation) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicPageViewAnalyticsFilterPruningRefineByUnion) AsTimeRanged() (v PublicRangedTimeOperation) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u PublicPageViewAnalyticsFilterPruningRefineByUnion) RawJSON() string { return u.JSON.raw }

func (r *PublicPageViewAnalyticsFilterPruningRefineByUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties FilterType, Operator, PageURL are required.
type PublicPageViewAnalyticsFilterParam struct {
	// Indicates the type of filter being applied (PAGE_VIEW).
	//
	// Any of "PAGE_VIEW".
	FilterType PublicPageViewAnalyticsFilterFilterType `json:"filterType,omitzero" api:"required"`
	// Defines the operation to be applied within the filter (HAS_PAGEVIEW_EQ,
	// HAS_PAGEVIEW_CONTAINS, HAS_PAGEVIEW_MATCHES_REGEX, NOT_HAS_PAGEVIEW_EQ,
	// NOT_HAS_PAGEVIEW_CONTAINS).
	Operator string `json:"operator" api:"required"`
	// The URL of the page to be used in the filter.
	PageURL string `json:"pageUrl" api:"required"`
	// Indicates whether tracking is enabled for the page view.
	EnableTracking param.Opt[bool] `json:"enableTracking,omitzero"`
	// Specifies the criteria for refining the filter by coalescing.
	CoalescingRefineBy PublicPageViewAnalyticsFilterCoalescingRefineByUnionParam `json:"coalescingRefineBy,omitzero"`
	// Specifies the criteria for refining the filter by pruning.
	PruningRefineBy PublicPageViewAnalyticsFilterPruningRefineByUnionParam `json:"pruningRefineBy,omitzero"`
	paramObj
}

func (r PublicPageViewAnalyticsFilterParam) MarshalJSON() (data []byte, err error) {
	type shadow PublicPageViewAnalyticsFilterParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PublicPageViewAnalyticsFilterParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type PublicPageViewAnalyticsFilterCoalescingRefineByUnionParam struct {
	OfNumOccurrences      *PublicNumOccurrencesRefineByParam               `json:",omitzero,inline"`
	OfSetOccurrences      *PublicSetOccurrencesRefineByParam               `json:",omitzero,inline"`
	OfRelativeComparative *PublicRelativeComparativeTimestampRefineByParam `json:",omitzero,inline"`
	OfRelativeRanged      *PublicRelativeRangedTimestampRefineByParam      `json:",omitzero,inline"`
	OfAbsoluteComparative *PublicAbsoluteComparativeTimestampRefineByParam `json:",omitzero,inline"`
	OfAbsoluteRanged      *PublicAbsoluteRangedTimestampRefineByParam      `json:",omitzero,inline"`
	OfAllHistory          *PublicAllHistoryRefineByParam                   `json:",omitzero,inline"`
	OfTimePoint           *PublicTimePointOperationParam                   `json:",omitzero,inline"`
	OfTimeRanged          *PublicRangedTimeOperationParam                  `json:",omitzero,inline"`
	paramUnion
}

func (u PublicPageViewAnalyticsFilterCoalescingRefineByUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfNumOccurrences,
		u.OfSetOccurrences,
		u.OfRelativeComparative,
		u.OfRelativeRanged,
		u.OfAbsoluteComparative,
		u.OfAbsoluteRanged,
		u.OfAllHistory,
		u.OfTimePoint,
		u.OfTimeRanged)
}
func (u *PublicPageViewAnalyticsFilterCoalescingRefineByUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type PublicPageViewAnalyticsFilterPruningRefineByUnionParam struct {
	OfNumOccurrences      *PublicNumOccurrencesRefineByParam               `json:",omitzero,inline"`
	OfSetOccurrences      *PublicSetOccurrencesRefineByParam               `json:",omitzero,inline"`
	OfRelativeComparative *PublicRelativeComparativeTimestampRefineByParam `json:",omitzero,inline"`
	OfRelativeRanged      *PublicRelativeRangedTimestampRefineByParam      `json:",omitzero,inline"`
	OfAbsoluteComparative *PublicAbsoluteComparativeTimestampRefineByParam `json:",omitzero,inline"`
	OfAbsoluteRanged      *PublicAbsoluteRangedTimestampRefineByParam      `json:",omitzero,inline"`
	OfAllHistory          *PublicAllHistoryRefineByParam                   `json:",omitzero,inline"`
	OfTimePoint           *PublicTimePointOperationParam                   `json:",omitzero,inline"`
	OfTimeRanged          *PublicRangedTimeOperationParam                  `json:",omitzero,inline"`
	paramUnion
}

func (u PublicPageViewAnalyticsFilterPruningRefineByUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfNumOccurrences,
		u.OfSetOccurrences,
		u.OfRelativeComparative,
		u.OfRelativeRanged,
		u.OfAbsoluteComparative,
		u.OfAbsoluteRanged,
		u.OfAllHistory,
		u.OfTimePoint,
		u.OfTimeRanged)
}
func (u *PublicPageViewAnalyticsFilterPruningRefineByUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

type PublicPrivacyAnalyticsFilter struct {
	// Specifies the type of filter (PRIVACY).
	//
	// Any of "PRIVACY".
	FilterType PublicPrivacyAnalyticsFilterFilterType `json:"filterType" api:"required"`
	// Defines the operation to be applied within the filter (PRIVACY_CONSENT_GRANTED,
	// PRIVACY_CONSENT_NOT_GRANTED).
	Operator string `json:"operator" api:"required"`
	// The name of the privacy setting used in the filter.
	PrivacyName string `json:"privacyName" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FilterType  respjson.Field
		Operator    respjson.Field
		PrivacyName respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PublicPrivacyAnalyticsFilter) RawJSON() string { return r.JSON.raw }
func (r *PublicPrivacyAnalyticsFilter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this PublicPrivacyAnalyticsFilter to a
// PublicPrivacyAnalyticsFilterParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// PublicPrivacyAnalyticsFilterParam.Overrides()
func (r PublicPrivacyAnalyticsFilter) ToParam() PublicPrivacyAnalyticsFilterParam {
	return param.Override[PublicPrivacyAnalyticsFilterParam](json.RawMessage(r.RawJSON()))
}

// Specifies the type of filter (PRIVACY).
type PublicPrivacyAnalyticsFilterFilterType string

const (
	PublicPrivacyAnalyticsFilterFilterTypePrivacy PublicPrivacyAnalyticsFilterFilterType = "PRIVACY"
)

// The properties FilterType, Operator, PrivacyName are required.
type PublicPrivacyAnalyticsFilterParam struct {
	// Specifies the type of filter (PRIVACY).
	//
	// Any of "PRIVACY".
	FilterType PublicPrivacyAnalyticsFilterFilterType `json:"filterType,omitzero" api:"required"`
	// Defines the operation to be applied within the filter (PRIVACY_CONSENT_GRANTED,
	// PRIVACY_CONSENT_NOT_GRANTED).
	Operator string `json:"operator" api:"required"`
	// The name of the privacy setting used in the filter.
	PrivacyName string `json:"privacyName" api:"required"`
	paramObj
}

func (r PublicPrivacyAnalyticsFilterParam) MarshalJSON() (data []byte, err error) {
	type shadow PublicPrivacyAnalyticsFilterParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PublicPrivacyAnalyticsFilterParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PublicPropertyAssociationFilterBranch struct {
	FilterBranches []PublicPropertyAssociationFilterBranchFilterBranchUnion `json:"filterBranches" api:"required"`
	// The logical operator used to combine filters within the branch.
	FilterBranchOperator string `json:"filterBranchOperator" api:"required"`
	// The type of the filter branch (PROPERTY_ASSOCIATION).
	//
	// Any of "PROPERTY_ASSOCIATION".
	FilterBranchType PublicPropertyAssociationFilterBranchFilterBranchType `json:"filterBranchType" api:"required"`
	Filters          []PublicPropertyAssociationFilterBranchFilterUnion    `json:"filters" api:"required"`
	// The ID representing the type of object associated with the filter branch.
	ObjectTypeID string `json:"objectTypeId" api:"required"`
	// Defines the operation to be applied within the filter branch (IN_LIST,
	// NOT_IN_LIST).
	Operator string `json:"operator" api:"required"`
	// The property that is associated with the object ID in the filter branch.
	PropertyWithObjectID string `json:"propertyWithObjectId" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FilterBranches       respjson.Field
		FilterBranchOperator respjson.Field
		FilterBranchType     respjson.Field
		Filters              respjson.Field
		ObjectTypeID         respjson.Field
		Operator             respjson.Field
		PropertyWithObjectID respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PublicPropertyAssociationFilterBranch) RawJSON() string { return r.JSON.raw }
func (r *PublicPropertyAssociationFilterBranch) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this PublicPropertyAssociationFilterBranch to a
// PublicPropertyAssociationFilterBranchParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// PublicPropertyAssociationFilterBranchParam.Overrides()
func (r PublicPropertyAssociationFilterBranch) ToParam() PublicPropertyAssociationFilterBranchParam {
	return param.Override[PublicPropertyAssociationFilterBranchParam](json.RawMessage(r.RawJSON()))
}

// PublicPropertyAssociationFilterBranchFilterBranchUnion contains all possible
// properties and values from [PublicOrFilterBranch], [PublicAndFilterBranch],
// [PublicNotAllFilterBranch], [PublicNotAnyFilterBranch],
// [PublicRestrictedFilterBranch], [PublicUnifiedEventsFilterBranch],
// [PublicPropertyAssociationFilterBranch], [PublicAssociationFilterBranch].
//
// Use the [PublicPropertyAssociationFilterBranchFilterBranchUnion.AsAny] method to
// switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type PublicPropertyAssociationFilterBranchFilterBranchUnion struct {
	// This field is a union of [[]PublicOrFilterBranchFilterBranchUnion],
	// [[]PublicAndFilterBranchFilterBranchUnion],
	// [[]PublicNotAllFilterBranchFilterBranchUnion],
	// [[]PublicNotAnyFilterBranchFilterBranchUnion],
	// [[]PublicRestrictedFilterBranchFilterBranchUnion],
	// [[]PublicUnifiedEventsFilterBranchFilterBranchUnion],
	// [[]PublicPropertyAssociationFilterBranchFilterBranchUnion],
	// [[]PublicAssociationFilterBranchFilterBranchUnion]
	FilterBranches       PublicPropertyAssociationFilterBranchFilterBranchUnionFilterBranches `json:"filterBranches"`
	FilterBranchOperator string                                                               `json:"filterBranchOperator"`
	// Any of "OR", "AND", "NOT_ALL", "NOT_ANY", "RESTRICTED", "UNIFIED_EVENTS",
	// "PROPERTY_ASSOCIATION", "ASSOCIATION".
	FilterBranchType string `json:"filterBranchType"`
	// This field is a union of [[]PublicOrFilterBranchFilterUnion],
	// [[]PublicAndFilterBranchFilterUnion], [[]PublicNotAllFilterBranchFilterUnion],
	// [[]PublicNotAnyFilterBranchFilterUnion],
	// [[]PublicRestrictedFilterBranchFilterUnion],
	// [[]PublicUnifiedEventsFilterBranchFilterUnion],
	// [[]PublicPropertyAssociationFilterBranchFilterUnion],
	// [[]PublicAssociationFilterBranchFilterUnion]
	Filters PublicPropertyAssociationFilterBranchFilterBranchUnionFilters `json:"filters"`
	// This field is from variant [PublicUnifiedEventsFilterBranch].
	EventTypeID string `json:"eventTypeId"`
	Operator    string `json:"operator"`
	// This field is from variant [PublicUnifiedEventsFilterBranch].
	CoalescingRefineBy PublicUnifiedEventsFilterBranchCoalescingRefineByUnion `json:"coalescingRefineBy"`
	// This field is from variant [PublicUnifiedEventsFilterBranch].
	PruningRefineBy PublicUnifiedEventsFilterBranchPruningRefineByUnion `json:"pruningRefineBy"`
	ObjectTypeID    string                                              `json:"objectTypeId"`
	// This field is from variant [PublicPropertyAssociationFilterBranch].
	PropertyWithObjectID string `json:"propertyWithObjectId"`
	// This field is from variant [PublicAssociationFilterBranch].
	AssociationCategory string `json:"associationCategory"`
	// This field is from variant [PublicAssociationFilterBranch].
	AssociationTypeID int64 `json:"associationTypeId"`
	JSON              struct {
		FilterBranches       respjson.Field
		FilterBranchOperator respjson.Field
		FilterBranchType     respjson.Field
		Filters              respjson.Field
		EventTypeID          respjson.Field
		Operator             respjson.Field
		CoalescingRefineBy   respjson.Field
		PruningRefineBy      respjson.Field
		ObjectTypeID         respjson.Field
		PropertyWithObjectID respjson.Field
		AssociationCategory  respjson.Field
		AssociationTypeID    respjson.Field
		raw                  string
	} `json:"-"`
}

// anyPublicPropertyAssociationFilterBranchFilterBranch is implemented by each
// variant of [PublicPropertyAssociationFilterBranchFilterBranchUnion] to add type
// safety for the return type of
// [PublicPropertyAssociationFilterBranchFilterBranchUnion.AsAny]
type anyPublicPropertyAssociationFilterBranchFilterBranch interface {
	implPublicPropertyAssociationFilterBranchFilterBranchUnion()
}

func (PublicOrFilterBranch) implPublicPropertyAssociationFilterBranchFilterBranchUnion()            {}
func (PublicAndFilterBranch) implPublicPropertyAssociationFilterBranchFilterBranchUnion()           {}
func (PublicNotAllFilterBranch) implPublicPropertyAssociationFilterBranchFilterBranchUnion()        {}
func (PublicNotAnyFilterBranch) implPublicPropertyAssociationFilterBranchFilterBranchUnion()        {}
func (PublicRestrictedFilterBranch) implPublicPropertyAssociationFilterBranchFilterBranchUnion()    {}
func (PublicUnifiedEventsFilterBranch) implPublicPropertyAssociationFilterBranchFilterBranchUnion() {}
func (PublicPropertyAssociationFilterBranch) implPublicPropertyAssociationFilterBranchFilterBranchUnion() {
}
func (PublicAssociationFilterBranch) implPublicPropertyAssociationFilterBranchFilterBranchUnion() {}

// Use the following switch statement to find the correct variant
//
//	switch variant := PublicPropertyAssociationFilterBranchFilterBranchUnion.AsAny().(type) {
//	case crm.PublicOrFilterBranch:
//	case crm.PublicAndFilterBranch:
//	case crm.PublicNotAllFilterBranch:
//	case crm.PublicNotAnyFilterBranch:
//	case crm.PublicRestrictedFilterBranch:
//	case crm.PublicUnifiedEventsFilterBranch:
//	case crm.PublicPropertyAssociationFilterBranch:
//	case crm.PublicAssociationFilterBranch:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u PublicPropertyAssociationFilterBranchFilterBranchUnion) AsAny() anyPublicPropertyAssociationFilterBranchFilterBranch {
	switch u.FilterBranchType {
	case "OR":
		return u.AsOr()
	case "AND":
		return u.AsAnd()
	case "NOT_ALL":
		return u.AsNotAll()
	case "NOT_ANY":
		return u.AsNotAny()
	case "RESTRICTED":
		return u.AsRestricted()
	case "UNIFIED_EVENTS":
		return u.AsUnifiedEvents()
	case "PROPERTY_ASSOCIATION":
		return u.AsPropertyAssociation()
	case "ASSOCIATION":
		return u.AsAssociation()
	}
	return nil
}

func (u PublicPropertyAssociationFilterBranchFilterBranchUnion) AsOr() (v PublicOrFilterBranch) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicPropertyAssociationFilterBranchFilterBranchUnion) AsAnd() (v PublicAndFilterBranch) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicPropertyAssociationFilterBranchFilterBranchUnion) AsNotAll() (v PublicNotAllFilterBranch) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicPropertyAssociationFilterBranchFilterBranchUnion) AsNotAny() (v PublicNotAnyFilterBranch) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicPropertyAssociationFilterBranchFilterBranchUnion) AsRestricted() (v PublicRestrictedFilterBranch) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicPropertyAssociationFilterBranchFilterBranchUnion) AsUnifiedEvents() (v PublicUnifiedEventsFilterBranch) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicPropertyAssociationFilterBranchFilterBranchUnion) AsPropertyAssociation() (v PublicPropertyAssociationFilterBranch) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicPropertyAssociationFilterBranchFilterBranchUnion) AsAssociation() (v PublicAssociationFilterBranch) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u PublicPropertyAssociationFilterBranchFilterBranchUnion) RawJSON() string { return u.JSON.raw }

func (r *PublicPropertyAssociationFilterBranchFilterBranchUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PublicPropertyAssociationFilterBranchFilterBranchUnionFilterBranches is an
// implicit subunion of [PublicPropertyAssociationFilterBranchFilterBranchUnion].
// PublicPropertyAssociationFilterBranchFilterBranchUnionFilterBranches provides
// convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [PublicPropertyAssociationFilterBranchFilterBranchUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfPublicOrFilterBranchFilterBranches
// OfPublicAndFilterBranchFilterBranches OfPublicNotAllFilterBranchFilterBranches
// OfPublicNotAnyFilterBranchFilterBranches
// OfPublicRestrictedFilterBranchFilterBranches
// OfPublicUnifiedEventsFilterBranchFilterBranches
// OfPublicPropertyAssociationFilterBranchFilterBranches
// OfPublicAssociationFilterBranchFilterBranches]
type PublicPropertyAssociationFilterBranchFilterBranchUnionFilterBranches struct {
	// This field will be present if the value is a
	// [[]PublicOrFilterBranchFilterBranchUnion] instead of an object.
	OfPublicOrFilterBranchFilterBranches []PublicOrFilterBranchFilterBranchUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]PublicAndFilterBranchFilterBranchUnion] instead of an object.
	OfPublicAndFilterBranchFilterBranches []PublicAndFilterBranchFilterBranchUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]PublicNotAllFilterBranchFilterBranchUnion] instead of an object.
	OfPublicNotAllFilterBranchFilterBranches []PublicNotAllFilterBranchFilterBranchUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]PublicNotAnyFilterBranchFilterBranchUnion] instead of an object.
	OfPublicNotAnyFilterBranchFilterBranches []PublicNotAnyFilterBranchFilterBranchUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]PublicRestrictedFilterBranchFilterBranchUnion] instead of an object.
	OfPublicRestrictedFilterBranchFilterBranches []PublicRestrictedFilterBranchFilterBranchUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]PublicUnifiedEventsFilterBranchFilterBranchUnion] instead of an object.
	OfPublicUnifiedEventsFilterBranchFilterBranches []PublicUnifiedEventsFilterBranchFilterBranchUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]PublicPropertyAssociationFilterBranchFilterBranchUnion] instead of an object.
	OfPublicPropertyAssociationFilterBranchFilterBranches []PublicPropertyAssociationFilterBranchFilterBranchUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]PublicAssociationFilterBranchFilterBranchUnion] instead of an object.
	OfPublicAssociationFilterBranchFilterBranches []PublicAssociationFilterBranchFilterBranchUnion `json:",inline"`
	JSON                                          struct {
		OfPublicOrFilterBranchFilterBranches                  respjson.Field
		OfPublicAndFilterBranchFilterBranches                 respjson.Field
		OfPublicNotAllFilterBranchFilterBranches              respjson.Field
		OfPublicNotAnyFilterBranchFilterBranches              respjson.Field
		OfPublicRestrictedFilterBranchFilterBranches          respjson.Field
		OfPublicUnifiedEventsFilterBranchFilterBranches       respjson.Field
		OfPublicPropertyAssociationFilterBranchFilterBranches respjson.Field
		OfPublicAssociationFilterBranchFilterBranches         respjson.Field
		raw                                                   string
	} `json:"-"`
}

func (r *PublicPropertyAssociationFilterBranchFilterBranchUnionFilterBranches) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PublicPropertyAssociationFilterBranchFilterBranchUnionFilters is an implicit
// subunion of [PublicPropertyAssociationFilterBranchFilterBranchUnion].
// PublicPropertyAssociationFilterBranchFilterBranchUnionFilters provides
// convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [PublicPropertyAssociationFilterBranchFilterBranchUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfPublicOrFilterBranchFilters OfPublicAndFilterBranchFilters
// OfPublicNotAllFilterBranchFilters OfPublicNotAnyFilterBranchFilters
// OfPublicRestrictedFilterBranchFilters OfPublicUnifiedEventsFilterBranchFilters
// OfPublicPropertyAssociationFilterBranchFilters
// OfPublicAssociationFilterBranchFilters]
type PublicPropertyAssociationFilterBranchFilterBranchUnionFilters struct {
	// This field will be present if the value is a [[]PublicOrFilterBranchFilterUnion]
	// instead of an object.
	OfPublicOrFilterBranchFilters []PublicOrFilterBranchFilterUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]PublicAndFilterBranchFilterUnion] instead of an object.
	OfPublicAndFilterBranchFilters []PublicAndFilterBranchFilterUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]PublicNotAllFilterBranchFilterUnion] instead of an object.
	OfPublicNotAllFilterBranchFilters []PublicNotAllFilterBranchFilterUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]PublicNotAnyFilterBranchFilterUnion] instead of an object.
	OfPublicNotAnyFilterBranchFilters []PublicNotAnyFilterBranchFilterUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]PublicRestrictedFilterBranchFilterUnion] instead of an object.
	OfPublicRestrictedFilterBranchFilters []PublicRestrictedFilterBranchFilterUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]PublicUnifiedEventsFilterBranchFilterUnion] instead of an object.
	OfPublicUnifiedEventsFilterBranchFilters []PublicUnifiedEventsFilterBranchFilterUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]PublicPropertyAssociationFilterBranchFilterUnion] instead of an object.
	OfPublicPropertyAssociationFilterBranchFilters []PublicPropertyAssociationFilterBranchFilterUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]PublicAssociationFilterBranchFilterUnion] instead of an object.
	OfPublicAssociationFilterBranchFilters []PublicAssociationFilterBranchFilterUnion `json:",inline"`
	JSON                                   struct {
		OfPublicOrFilterBranchFilters                  respjson.Field
		OfPublicAndFilterBranchFilters                 respjson.Field
		OfPublicNotAllFilterBranchFilters              respjson.Field
		OfPublicNotAnyFilterBranchFilters              respjson.Field
		OfPublicRestrictedFilterBranchFilters          respjson.Field
		OfPublicUnifiedEventsFilterBranchFilters       respjson.Field
		OfPublicPropertyAssociationFilterBranchFilters respjson.Field
		OfPublicAssociationFilterBranchFilters         respjson.Field
		raw                                            string
	} `json:"-"`
}

func (r *PublicPropertyAssociationFilterBranchFilterBranchUnionFilters) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The type of the filter branch (PROPERTY_ASSOCIATION).
type PublicPropertyAssociationFilterBranchFilterBranchType string

const (
	PublicPropertyAssociationFilterBranchFilterBranchTypePropertyAssociation PublicPropertyAssociationFilterBranchFilterBranchType = "PROPERTY_ASSOCIATION"
)

// PublicPropertyAssociationFilterBranchFilterUnion contains all possible
// properties and values from [PublicPropertyFilter],
// [PublicAssociationInListFilter], [PublicPageViewAnalyticsFilter],
// [PublicCtaAnalyticsFilter], [PublicEventAnalyticsFilter],
// [PublicFormSubmissionFilter], [PublicFormSubmissionOnPageFilter],
// [PublicIntegrationEventFilter], [PublicEmailSubscriptionFilter],
// [PublicCommunicationSubscriptionFilter], [PublicCampaignInfluencedFilter],
// [PublicSurveyMonkeyFilter], [PublicSurveyMonkeyValueFilter],
// [PublicWebinarFilter], [PublicEmailEventFilter], [PublicPrivacyAnalyticsFilter],
// [PublicAdsSearchFilter], [PublicAdsTimeFilter], [PublicInListFilter],
// [PublicNumAssociationsFilter], [PublicUnifiedEventsFilter],
// [PublicPropertyAssociationInListFilter], [PublicConstantFilter].
//
// Use the [PublicPropertyAssociationFilterBranchFilterUnion.AsAny] method to
// switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type PublicPropertyAssociationFilterBranchFilterUnion struct {
	// Any of "PROPERTY", "ASSOCIATION", "PAGE_VIEW", "CTA", "EVENT",
	// "FORM_SUBMISSION", "FORM_SUBMISSION_ON_PAGE", "INTEGRATION_EVENT",
	// "EMAIL_SUBSCRIPTION", "COMMUNICATION_SUBSCRIPTION", "CAMPAIGN_INFLUENCED",
	// "SURVEY_MONKEY", "SURVEY_MONKEY_VALUE", "WEBINAR", "EMAIL_EVENT", "PRIVACY",
	// "ADS_SEARCH", "ADS_TIME", "IN_LIST", "NUM_ASSOCIATIONS", "UNIFIED_EVENTS",
	// "PROPERTY_ASSOCIATION", "CONSTANT".
	FilterType string `json:"filterType"`
	// This field is from variant [PublicPropertyFilter].
	Operation PublicPropertyFilterOperationUnion `json:"operation"`
	// This field is from variant [PublicPropertyFilter].
	Property            string `json:"property"`
	AssociationCategory string `json:"associationCategory"`
	AssociationTypeID   int64  `json:"associationTypeId"`
	// This field is a union of [PublicAssociationInListFilterCoalescingRefineByUnion],
	// [PublicPageViewAnalyticsFilterCoalescingRefineByUnion],
	// [PublicCtaAnalyticsFilterCoalescingRefineByUnion],
	// [PublicEventAnalyticsFilterCoalescingRefineByUnion],
	// [PublicFormSubmissionFilterCoalescingRefineByUnion],
	// [PublicFormSubmissionOnPageFilterCoalescingRefineByUnion],
	// [PublicNumAssociationsFilterCoalescingRefineByUnion],
	// [PublicUnifiedEventsFilterCoalescingRefineByUnion],
	// [PublicPropertyAssociationInListFilterCoalescingRefineByUnion]
	CoalescingRefineBy PublicPropertyAssociationFilterBranchFilterUnionCoalescingRefineBy `json:"coalescingRefineBy"`
	ListID             string                                                             `json:"listId"`
	Operator           string                                                             `json:"operator"`
	// This field is from variant [PublicAssociationInListFilter].
	ToObjectType   string `json:"toObjectType"`
	ToObjectTypeID string `json:"toObjectTypeId"`
	// This field is from variant [PublicPageViewAnalyticsFilter].
	PageURL string `json:"pageUrl"`
	// This field is from variant [PublicPageViewAnalyticsFilter].
	EnableTracking bool `json:"enableTracking"`
	// This field is a union of [PublicPageViewAnalyticsFilterPruningRefineByUnion],
	// [PublicCtaAnalyticsFilterPruningRefineByUnion],
	// [PublicEventAnalyticsFilterPruningRefineByUnion],
	// [PublicFormSubmissionFilterPruningRefineByUnion],
	// [PublicFormSubmissionOnPageFilterPruningRefineByUnion],
	// [PublicEmailEventFilterPruningRefineByUnion],
	// [PublicAdsTimeFilterPruningRefineByUnion],
	// [PublicUnifiedEventsFilterPruningRefineByUnion]
	PruningRefineBy PublicPropertyAssociationFilterBranchFilterUnionPruningRefineBy `json:"pruningRefineBy"`
	// This field is from variant [PublicCtaAnalyticsFilter].
	CtaName string `json:"ctaName"`
	// This field is from variant [PublicEventAnalyticsFilter].
	EventID string `json:"eventId"`
	FormID  string `json:"formId"`
	// This field is from variant [PublicFormSubmissionOnPageFilter].
	PageID string `json:"pageId"`
	// This field is a union of [int64], [string]
	EventTypeID PublicPropertyAssociationFilterBranchFilterUnionEventTypeID `json:"eventTypeId"`
	FilterLines []PublicEventFilterMetadata                                 `json:"filterLines"`
	// This field is from variant [PublicEmailSubscriptionFilter].
	AcceptedStatuses []string `json:"acceptedStatuses"`
	SubscriptionIDs  []string `json:"subscriptionIds"`
	SubscriptionType string   `json:"subscriptionType"`
	// This field is from variant [PublicCommunicationSubscriptionFilter].
	AcceptedOptStates []string `json:"acceptedOptStates"`
	// This field is from variant [PublicCommunicationSubscriptionFilter].
	Channel string `json:"channel"`
	// This field is from variant [PublicCommunicationSubscriptionFilter].
	BusinessUnitID string `json:"businessUnitId"`
	// This field is from variant [PublicCampaignInfluencedFilter].
	CampaignID string `json:"campaignId"`
	SurveyID   string `json:"surveyId"`
	// This field is from variant [PublicSurveyMonkeyValueFilter].
	SurveyQuestion string `json:"surveyQuestion"`
	// This field is from variant [PublicSurveyMonkeyValueFilter].
	ValueComparison PublicSurveyMonkeyValueFilterValueComparisonUnion `json:"valueComparison"`
	// This field is from variant [PublicSurveyMonkeyValueFilter].
	SurveyAnswerColID string `json:"surveyAnswerColId"`
	// This field is from variant [PublicSurveyMonkeyValueFilter].
	SurveyAnswerRowID string `json:"surveyAnswerRowId"`
	// This field is from variant [PublicWebinarFilter].
	WebinarID string `json:"webinarId"`
	// This field is from variant [PublicEmailEventFilter].
	AppID string `json:"appId"`
	// This field is from variant [PublicEmailEventFilter].
	EmailID string `json:"emailId"`
	// This field is from variant [PublicEmailEventFilter].
	Level string `json:"level"`
	// This field is from variant [PublicEmailEventFilter].
	ClickURL string `json:"clickUrl"`
	// This field is from variant [PublicPrivacyAnalyticsFilter].
	PrivacyName string `json:"privacyName"`
	// This field is from variant [PublicAdsSearchFilter].
	AdNetwork string `json:"adNetwork"`
	// This field is from variant [PublicAdsSearchFilter].
	EntityType string `json:"entityType"`
	// This field is from variant [PublicAdsSearchFilter].
	SearchTerms []string `json:"searchTerms"`
	// This field is from variant [PublicAdsSearchFilter].
	SearchTermType string `json:"searchTermType"`
	// This field is from variant [PublicInListFilter].
	Metadata PublicInListFilterMetadata `json:"metadata"`
	// This field is from variant [PublicPropertyAssociationInListFilter].
	PropertyWithObjectID string `json:"propertyWithObjectId"`
	// This field is from variant [PublicConstantFilter].
	ShouldAccept bool `json:"shouldAccept"`
	// This field is from variant [PublicConstantFilter].
	Source string `json:"source"`
	JSON   struct {
		FilterType           respjson.Field
		Operation            respjson.Field
		Property             respjson.Field
		AssociationCategory  respjson.Field
		AssociationTypeID    respjson.Field
		CoalescingRefineBy   respjson.Field
		ListID               respjson.Field
		Operator             respjson.Field
		ToObjectType         respjson.Field
		ToObjectTypeID       respjson.Field
		PageURL              respjson.Field
		EnableTracking       respjson.Field
		PruningRefineBy      respjson.Field
		CtaName              respjson.Field
		EventID              respjson.Field
		FormID               respjson.Field
		PageID               respjson.Field
		EventTypeID          respjson.Field
		FilterLines          respjson.Field
		AcceptedStatuses     respjson.Field
		SubscriptionIDs      respjson.Field
		SubscriptionType     respjson.Field
		AcceptedOptStates    respjson.Field
		Channel              respjson.Field
		BusinessUnitID       respjson.Field
		CampaignID           respjson.Field
		SurveyID             respjson.Field
		SurveyQuestion       respjson.Field
		ValueComparison      respjson.Field
		SurveyAnswerColID    respjson.Field
		SurveyAnswerRowID    respjson.Field
		WebinarID            respjson.Field
		AppID                respjson.Field
		EmailID              respjson.Field
		Level                respjson.Field
		ClickURL             respjson.Field
		PrivacyName          respjson.Field
		AdNetwork            respjson.Field
		EntityType           respjson.Field
		SearchTerms          respjson.Field
		SearchTermType       respjson.Field
		Metadata             respjson.Field
		PropertyWithObjectID respjson.Field
		ShouldAccept         respjson.Field
		Source               respjson.Field
		raw                  string
	} `json:"-"`
}

// anyPublicPropertyAssociationFilterBranchFilter is implemented by each variant of
// [PublicPropertyAssociationFilterBranchFilterUnion] to add type safety for the
// return type of [PublicPropertyAssociationFilterBranchFilterUnion.AsAny]
type anyPublicPropertyAssociationFilterBranchFilter interface {
	implPublicPropertyAssociationFilterBranchFilterUnion()
}

func (PublicPropertyFilter) implPublicPropertyAssociationFilterBranchFilterUnion()                  {}
func (PublicAssociationInListFilter) implPublicPropertyAssociationFilterBranchFilterUnion()         {}
func (PublicPageViewAnalyticsFilter) implPublicPropertyAssociationFilterBranchFilterUnion()         {}
func (PublicCtaAnalyticsFilter) implPublicPropertyAssociationFilterBranchFilterUnion()              {}
func (PublicEventAnalyticsFilter) implPublicPropertyAssociationFilterBranchFilterUnion()            {}
func (PublicFormSubmissionFilter) implPublicPropertyAssociationFilterBranchFilterUnion()            {}
func (PublicFormSubmissionOnPageFilter) implPublicPropertyAssociationFilterBranchFilterUnion()      {}
func (PublicIntegrationEventFilter) implPublicPropertyAssociationFilterBranchFilterUnion()          {}
func (PublicEmailSubscriptionFilter) implPublicPropertyAssociationFilterBranchFilterUnion()         {}
func (PublicCommunicationSubscriptionFilter) implPublicPropertyAssociationFilterBranchFilterUnion() {}
func (PublicCampaignInfluencedFilter) implPublicPropertyAssociationFilterBranchFilterUnion()        {}
func (PublicSurveyMonkeyFilter) implPublicPropertyAssociationFilterBranchFilterUnion()              {}
func (PublicSurveyMonkeyValueFilter) implPublicPropertyAssociationFilterBranchFilterUnion()         {}
func (PublicWebinarFilter) implPublicPropertyAssociationFilterBranchFilterUnion()                   {}
func (PublicEmailEventFilter) implPublicPropertyAssociationFilterBranchFilterUnion()                {}
func (PublicPrivacyAnalyticsFilter) implPublicPropertyAssociationFilterBranchFilterUnion()          {}
func (PublicAdsSearchFilter) implPublicPropertyAssociationFilterBranchFilterUnion()                 {}
func (PublicAdsTimeFilter) implPublicPropertyAssociationFilterBranchFilterUnion()                   {}
func (PublicInListFilter) implPublicPropertyAssociationFilterBranchFilterUnion()                    {}
func (PublicNumAssociationsFilter) implPublicPropertyAssociationFilterBranchFilterUnion()           {}
func (PublicUnifiedEventsFilter) implPublicPropertyAssociationFilterBranchFilterUnion()             {}
func (PublicPropertyAssociationInListFilter) implPublicPropertyAssociationFilterBranchFilterUnion() {}
func (PublicConstantFilter) implPublicPropertyAssociationFilterBranchFilterUnion()                  {}

// Use the following switch statement to find the correct variant
//
//	switch variant := PublicPropertyAssociationFilterBranchFilterUnion.AsAny().(type) {
//	case crm.PublicPropertyFilter:
//	case crm.PublicAssociationInListFilter:
//	case crm.PublicPageViewAnalyticsFilter:
//	case crm.PublicCtaAnalyticsFilter:
//	case crm.PublicEventAnalyticsFilter:
//	case crm.PublicFormSubmissionFilter:
//	case crm.PublicFormSubmissionOnPageFilter:
//	case crm.PublicIntegrationEventFilter:
//	case crm.PublicEmailSubscriptionFilter:
//	case crm.PublicCommunicationSubscriptionFilter:
//	case crm.PublicCampaignInfluencedFilter:
//	case crm.PublicSurveyMonkeyFilter:
//	case crm.PublicSurveyMonkeyValueFilter:
//	case crm.PublicWebinarFilter:
//	case crm.PublicEmailEventFilter:
//	case crm.PublicPrivacyAnalyticsFilter:
//	case crm.PublicAdsSearchFilter:
//	case crm.PublicAdsTimeFilter:
//	case crm.PublicInListFilter:
//	case crm.PublicNumAssociationsFilter:
//	case crm.PublicUnifiedEventsFilter:
//	case crm.PublicPropertyAssociationInListFilter:
//	case crm.PublicConstantFilter:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u PublicPropertyAssociationFilterBranchFilterUnion) AsAny() anyPublicPropertyAssociationFilterBranchFilter {
	switch u.FilterType {
	case "PROPERTY":
		return u.AsProperty()
	case "ASSOCIATION":
		return u.AsAssociation()
	case "PAGE_VIEW":
		return u.AsPageView()
	case "CTA":
		return u.AsCta()
	case "EVENT":
		return u.AsEvent()
	case "FORM_SUBMISSION":
		return u.AsFormSubmission()
	case "FORM_SUBMISSION_ON_PAGE":
		return u.AsFormSubmissionOnPage()
	case "INTEGRATION_EVENT":
		return u.AsIntegrationEvent()
	case "EMAIL_SUBSCRIPTION":
		return u.AsEmailSubscription()
	case "COMMUNICATION_SUBSCRIPTION":
		return u.AsCommunicationSubscription()
	case "CAMPAIGN_INFLUENCED":
		return u.AsCampaignInfluenced()
	case "SURVEY_MONKEY":
		return u.AsSurveyMonkey()
	case "SURVEY_MONKEY_VALUE":
		return u.AsSurveyMonkeyValue()
	case "WEBINAR":
		return u.AsWebinar()
	case "EMAIL_EVENT":
		return u.AsEmailEvent()
	case "PRIVACY":
		return u.AsPrivacy()
	case "ADS_SEARCH":
		return u.AsAdsSearch()
	case "ADS_TIME":
		return u.AsAdsTime()
	case "IN_LIST":
		return u.AsInList()
	case "NUM_ASSOCIATIONS":
		return u.AsNumAssociations()
	case "UNIFIED_EVENTS":
		return u.AsUnifiedEvents()
	case "PROPERTY_ASSOCIATION":
		return u.AsPropertyAssociation()
	case "CONSTANT":
		return u.AsConstant()
	}
	return nil
}

func (u PublicPropertyAssociationFilterBranchFilterUnion) AsProperty() (v PublicPropertyFilter) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicPropertyAssociationFilterBranchFilterUnion) AsAssociation() (v PublicAssociationInListFilter) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicPropertyAssociationFilterBranchFilterUnion) AsPageView() (v PublicPageViewAnalyticsFilter) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicPropertyAssociationFilterBranchFilterUnion) AsCta() (v PublicCtaAnalyticsFilter) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicPropertyAssociationFilterBranchFilterUnion) AsEvent() (v PublicEventAnalyticsFilter) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicPropertyAssociationFilterBranchFilterUnion) AsFormSubmission() (v PublicFormSubmissionFilter) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicPropertyAssociationFilterBranchFilterUnion) AsFormSubmissionOnPage() (v PublicFormSubmissionOnPageFilter) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicPropertyAssociationFilterBranchFilterUnion) AsIntegrationEvent() (v PublicIntegrationEventFilter) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicPropertyAssociationFilterBranchFilterUnion) AsEmailSubscription() (v PublicEmailSubscriptionFilter) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicPropertyAssociationFilterBranchFilterUnion) AsCommunicationSubscription() (v PublicCommunicationSubscriptionFilter) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicPropertyAssociationFilterBranchFilterUnion) AsCampaignInfluenced() (v PublicCampaignInfluencedFilter) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicPropertyAssociationFilterBranchFilterUnion) AsSurveyMonkey() (v PublicSurveyMonkeyFilter) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicPropertyAssociationFilterBranchFilterUnion) AsSurveyMonkeyValue() (v PublicSurveyMonkeyValueFilter) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicPropertyAssociationFilterBranchFilterUnion) AsWebinar() (v PublicWebinarFilter) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicPropertyAssociationFilterBranchFilterUnion) AsEmailEvent() (v PublicEmailEventFilter) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicPropertyAssociationFilterBranchFilterUnion) AsPrivacy() (v PublicPrivacyAnalyticsFilter) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicPropertyAssociationFilterBranchFilterUnion) AsAdsSearch() (v PublicAdsSearchFilter) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicPropertyAssociationFilterBranchFilterUnion) AsAdsTime() (v PublicAdsTimeFilter) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicPropertyAssociationFilterBranchFilterUnion) AsInList() (v PublicInListFilter) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicPropertyAssociationFilterBranchFilterUnion) AsNumAssociations() (v PublicNumAssociationsFilter) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicPropertyAssociationFilterBranchFilterUnion) AsUnifiedEvents() (v PublicUnifiedEventsFilter) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicPropertyAssociationFilterBranchFilterUnion) AsPropertyAssociation() (v PublicPropertyAssociationInListFilter) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicPropertyAssociationFilterBranchFilterUnion) AsConstant() (v PublicConstantFilter) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u PublicPropertyAssociationFilterBranchFilterUnion) RawJSON() string { return u.JSON.raw }

func (r *PublicPropertyAssociationFilterBranchFilterUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PublicPropertyAssociationFilterBranchFilterUnionCoalescingRefineBy is an
// implicit subunion of [PublicPropertyAssociationFilterBranchFilterUnion].
// PublicPropertyAssociationFilterBranchFilterUnionCoalescingRefineBy provides
// convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [PublicPropertyAssociationFilterBranchFilterUnion].
type PublicPropertyAssociationFilterBranchFilterUnionCoalescingRefineBy struct {
	Type string `json:"type"`
	// This field is from variant
	// [PublicAssociationInListFilterCoalescingRefineByUnion],
	// [PublicPageViewAnalyticsFilterCoalescingRefineByUnion],
	// [PublicCtaAnalyticsFilterCoalescingRefineByUnion],
	// [PublicEventAnalyticsFilterCoalescingRefineByUnion],
	// [PublicFormSubmissionFilterCoalescingRefineByUnion],
	// [PublicFormSubmissionOnPageFilterCoalescingRefineByUnion],
	// [PublicNumAssociationsFilterCoalescingRefineByUnion],
	// [PublicUnifiedEventsFilterCoalescingRefineByUnion],
	// [PublicPropertyAssociationInListFilterCoalescingRefineByUnion].
	MaxOccurrences int64 `json:"maxOccurrences"`
	// This field is from variant
	// [PublicAssociationInListFilterCoalescingRefineByUnion],
	// [PublicPageViewAnalyticsFilterCoalescingRefineByUnion],
	// [PublicCtaAnalyticsFilterCoalescingRefineByUnion],
	// [PublicEventAnalyticsFilterCoalescingRefineByUnion],
	// [PublicFormSubmissionFilterCoalescingRefineByUnion],
	// [PublicFormSubmissionOnPageFilterCoalescingRefineByUnion],
	// [PublicNumAssociationsFilterCoalescingRefineByUnion],
	// [PublicUnifiedEventsFilterCoalescingRefineByUnion],
	// [PublicPropertyAssociationInListFilterCoalescingRefineByUnion].
	MinOccurrences int64 `json:"minOccurrences"`
	// This field is from variant
	// [PublicAssociationInListFilterCoalescingRefineByUnion],
	// [PublicPageViewAnalyticsFilterCoalescingRefineByUnion],
	// [PublicCtaAnalyticsFilterCoalescingRefineByUnion],
	// [PublicEventAnalyticsFilterCoalescingRefineByUnion],
	// [PublicFormSubmissionFilterCoalescingRefineByUnion],
	// [PublicFormSubmissionOnPageFilterCoalescingRefineByUnion],
	// [PublicNumAssociationsFilterCoalescingRefineByUnion],
	// [PublicUnifiedEventsFilterCoalescingRefineByUnion],
	// [PublicPropertyAssociationInListFilterCoalescingRefineByUnion].
	SetType    string `json:"setType"`
	Comparison string `json:"comparison"`
	// This field is from variant
	// [PublicAssociationInListFilterCoalescingRefineByUnion],
	// [PublicPageViewAnalyticsFilterCoalescingRefineByUnion],
	// [PublicCtaAnalyticsFilterCoalescingRefineByUnion],
	// [PublicEventAnalyticsFilterCoalescingRefineByUnion],
	// [PublicFormSubmissionFilterCoalescingRefineByUnion],
	// [PublicFormSubmissionOnPageFilterCoalescingRefineByUnion],
	// [PublicNumAssociationsFilterCoalescingRefineByUnion],
	// [PublicUnifiedEventsFilterCoalescingRefineByUnion],
	// [PublicPropertyAssociationInListFilterCoalescingRefineByUnion].
	TimeOffset PublicTimeOffset `json:"timeOffset"`
	// This field is from variant
	// [PublicAssociationInListFilterCoalescingRefineByUnion],
	// [PublicPageViewAnalyticsFilterCoalescingRefineByUnion],
	// [PublicCtaAnalyticsFilterCoalescingRefineByUnion],
	// [PublicEventAnalyticsFilterCoalescingRefineByUnion],
	// [PublicFormSubmissionFilterCoalescingRefineByUnion],
	// [PublicFormSubmissionOnPageFilterCoalescingRefineByUnion],
	// [PublicNumAssociationsFilterCoalescingRefineByUnion],
	// [PublicUnifiedEventsFilterCoalescingRefineByUnion],
	// [PublicPropertyAssociationInListFilterCoalescingRefineByUnion].
	LowerBoundOffset PublicTimeOffset `json:"lowerBoundOffset"`
	RangeType        string           `json:"rangeType"`
	// This field is from variant
	// [PublicAssociationInListFilterCoalescingRefineByUnion],
	// [PublicPageViewAnalyticsFilterCoalescingRefineByUnion],
	// [PublicCtaAnalyticsFilterCoalescingRefineByUnion],
	// [PublicEventAnalyticsFilterCoalescingRefineByUnion],
	// [PublicFormSubmissionFilterCoalescingRefineByUnion],
	// [PublicFormSubmissionOnPageFilterCoalescingRefineByUnion],
	// [PublicNumAssociationsFilterCoalescingRefineByUnion],
	// [PublicUnifiedEventsFilterCoalescingRefineByUnion],
	// [PublicPropertyAssociationInListFilterCoalescingRefineByUnion].
	UpperBoundOffset PublicTimeOffset `json:"upperBoundOffset"`
	// This field is from variant
	// [PublicAssociationInListFilterCoalescingRefineByUnion],
	// [PublicPageViewAnalyticsFilterCoalescingRefineByUnion],
	// [PublicCtaAnalyticsFilterCoalescingRefineByUnion],
	// [PublicEventAnalyticsFilterCoalescingRefineByUnion],
	// [PublicFormSubmissionFilterCoalescingRefineByUnion],
	// [PublicFormSubmissionOnPageFilterCoalescingRefineByUnion],
	// [PublicNumAssociationsFilterCoalescingRefineByUnion],
	// [PublicUnifiedEventsFilterCoalescingRefineByUnion],
	// [PublicPropertyAssociationInListFilterCoalescingRefineByUnion].
	Timestamp int64 `json:"timestamp"`
	// This field is from variant
	// [PublicAssociationInListFilterCoalescingRefineByUnion],
	// [PublicPageViewAnalyticsFilterCoalescingRefineByUnion],
	// [PublicCtaAnalyticsFilterCoalescingRefineByUnion],
	// [PublicEventAnalyticsFilterCoalescingRefineByUnion],
	// [PublicFormSubmissionFilterCoalescingRefineByUnion],
	// [PublicFormSubmissionOnPageFilterCoalescingRefineByUnion],
	// [PublicNumAssociationsFilterCoalescingRefineByUnion],
	// [PublicUnifiedEventsFilterCoalescingRefineByUnion],
	// [PublicPropertyAssociationInListFilterCoalescingRefineByUnion].
	LowerTimestamp int64 `json:"lowerTimestamp"`
	// This field is from variant
	// [PublicAssociationInListFilterCoalescingRefineByUnion],
	// [PublicPageViewAnalyticsFilterCoalescingRefineByUnion],
	// [PublicCtaAnalyticsFilterCoalescingRefineByUnion],
	// [PublicEventAnalyticsFilterCoalescingRefineByUnion],
	// [PublicFormSubmissionFilterCoalescingRefineByUnion],
	// [PublicFormSubmissionOnPageFilterCoalescingRefineByUnion],
	// [PublicNumAssociationsFilterCoalescingRefineByUnion],
	// [PublicUnifiedEventsFilterCoalescingRefineByUnion],
	// [PublicPropertyAssociationInListFilterCoalescingRefineByUnion].
	UpperTimestamp               int64  `json:"upperTimestamp"`
	IncludeObjectsWithNoValueSet bool   `json:"includeObjectsWithNoValueSet"`
	OperationType                string `json:"operationType"`
	Operator                     string `json:"operator"`
	// This field is from variant
	// [PublicAssociationInListFilterCoalescingRefineByUnion],
	// [PublicPageViewAnalyticsFilterCoalescingRefineByUnion],
	// [PublicCtaAnalyticsFilterCoalescingRefineByUnion],
	// [PublicEventAnalyticsFilterCoalescingRefineByUnion],
	// [PublicFormSubmissionFilterCoalescingRefineByUnion],
	// [PublicFormSubmissionOnPageFilterCoalescingRefineByUnion],
	// [PublicNumAssociationsFilterCoalescingRefineByUnion],
	// [PublicUnifiedEventsFilterCoalescingRefineByUnion],
	// [PublicPropertyAssociationInListFilterCoalescingRefineByUnion].
	TimePoint PublicTimePointOperationTimePointUnion `json:"timePoint"`
	// This field is from variant
	// [PublicAssociationInListFilterCoalescingRefineByUnion],
	// [PublicPageViewAnalyticsFilterCoalescingRefineByUnion],
	// [PublicCtaAnalyticsFilterCoalescingRefineByUnion],
	// [PublicEventAnalyticsFilterCoalescingRefineByUnion],
	// [PublicFormSubmissionFilterCoalescingRefineByUnion],
	// [PublicFormSubmissionOnPageFilterCoalescingRefineByUnion],
	// [PublicNumAssociationsFilterCoalescingRefineByUnion],
	// [PublicUnifiedEventsFilterCoalescingRefineByUnion],
	// [PublicPropertyAssociationInListFilterCoalescingRefineByUnion].
	EndpointBehavior string `json:"endpointBehavior"`
	PropertyParser   string `json:"propertyParser"`
	// This field is from variant
	// [PublicAssociationInListFilterCoalescingRefineByUnion],
	// [PublicPageViewAnalyticsFilterCoalescingRefineByUnion],
	// [PublicCtaAnalyticsFilterCoalescingRefineByUnion],
	// [PublicEventAnalyticsFilterCoalescingRefineByUnion],
	// [PublicFormSubmissionFilterCoalescingRefineByUnion],
	// [PublicFormSubmissionOnPageFilterCoalescingRefineByUnion],
	// [PublicNumAssociationsFilterCoalescingRefineByUnion],
	// [PublicUnifiedEventsFilterCoalescingRefineByUnion],
	// [PublicPropertyAssociationInListFilterCoalescingRefineByUnion].
	LowerBoundTimePoint PublicRangedTimeOperationLowerBoundTimePointUnion `json:"lowerBoundTimePoint"`
	// This field is from variant
	// [PublicAssociationInListFilterCoalescingRefineByUnion],
	// [PublicPageViewAnalyticsFilterCoalescingRefineByUnion],
	// [PublicCtaAnalyticsFilterCoalescingRefineByUnion],
	// [PublicEventAnalyticsFilterCoalescingRefineByUnion],
	// [PublicFormSubmissionFilterCoalescingRefineByUnion],
	// [PublicFormSubmissionOnPageFilterCoalescingRefineByUnion],
	// [PublicNumAssociationsFilterCoalescingRefineByUnion],
	// [PublicUnifiedEventsFilterCoalescingRefineByUnion],
	// [PublicPropertyAssociationInListFilterCoalescingRefineByUnion].
	UpperBoundTimePoint PublicRangedTimeOperationUpperBoundTimePointUnion `json:"upperBoundTimePoint"`
	// This field is from variant
	// [PublicAssociationInListFilterCoalescingRefineByUnion],
	// [PublicPageViewAnalyticsFilterCoalescingRefineByUnion],
	// [PublicCtaAnalyticsFilterCoalescingRefineByUnion],
	// [PublicEventAnalyticsFilterCoalescingRefineByUnion],
	// [PublicFormSubmissionFilterCoalescingRefineByUnion],
	// [PublicFormSubmissionOnPageFilterCoalescingRefineByUnion],
	// [PublicNumAssociationsFilterCoalescingRefineByUnion],
	// [PublicUnifiedEventsFilterCoalescingRefineByUnion],
	// [PublicPropertyAssociationInListFilterCoalescingRefineByUnion].
	LowerBoundEndpointBehavior string `json:"lowerBoundEndpointBehavior"`
	// This field is from variant
	// [PublicAssociationInListFilterCoalescingRefineByUnion],
	// [PublicPageViewAnalyticsFilterCoalescingRefineByUnion],
	// [PublicCtaAnalyticsFilterCoalescingRefineByUnion],
	// [PublicEventAnalyticsFilterCoalescingRefineByUnion],
	// [PublicFormSubmissionFilterCoalescingRefineByUnion],
	// [PublicFormSubmissionOnPageFilterCoalescingRefineByUnion],
	// [PublicNumAssociationsFilterCoalescingRefineByUnion],
	// [PublicUnifiedEventsFilterCoalescingRefineByUnion],
	// [PublicPropertyAssociationInListFilterCoalescingRefineByUnion].
	UpperBoundEndpointBehavior string `json:"upperBoundEndpointBehavior"`
	JSON                       struct {
		Type                         respjson.Field
		MaxOccurrences               respjson.Field
		MinOccurrences               respjson.Field
		SetType                      respjson.Field
		Comparison                   respjson.Field
		TimeOffset                   respjson.Field
		LowerBoundOffset             respjson.Field
		RangeType                    respjson.Field
		UpperBoundOffset             respjson.Field
		Timestamp                    respjson.Field
		LowerTimestamp               respjson.Field
		UpperTimestamp               respjson.Field
		IncludeObjectsWithNoValueSet respjson.Field
		OperationType                respjson.Field
		Operator                     respjson.Field
		TimePoint                    respjson.Field
		EndpointBehavior             respjson.Field
		PropertyParser               respjson.Field
		LowerBoundTimePoint          respjson.Field
		UpperBoundTimePoint          respjson.Field
		LowerBoundEndpointBehavior   respjson.Field
		UpperBoundEndpointBehavior   respjson.Field
		raw                          string
	} `json:"-"`
}

func (r *PublicPropertyAssociationFilterBranchFilterUnionCoalescingRefineBy) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PublicPropertyAssociationFilterBranchFilterUnionPruningRefineBy is an implicit
// subunion of [PublicPropertyAssociationFilterBranchFilterUnion].
// PublicPropertyAssociationFilterBranchFilterUnionPruningRefineBy provides
// convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [PublicPropertyAssociationFilterBranchFilterUnion].
type PublicPropertyAssociationFilterBranchFilterUnionPruningRefineBy struct {
	Type string `json:"type"`
	// This field is from variant [PublicPageViewAnalyticsFilterPruningRefineByUnion],
	// [PublicCtaAnalyticsFilterPruningRefineByUnion],
	// [PublicEventAnalyticsFilterPruningRefineByUnion],
	// [PublicFormSubmissionFilterPruningRefineByUnion],
	// [PublicFormSubmissionOnPageFilterPruningRefineByUnion],
	// [PublicEmailEventFilterPruningRefineByUnion],
	// [PublicAdsTimeFilterPruningRefineByUnion],
	// [PublicUnifiedEventsFilterPruningRefineByUnion].
	MaxOccurrences int64 `json:"maxOccurrences"`
	// This field is from variant [PublicPageViewAnalyticsFilterPruningRefineByUnion],
	// [PublicCtaAnalyticsFilterPruningRefineByUnion],
	// [PublicEventAnalyticsFilterPruningRefineByUnion],
	// [PublicFormSubmissionFilterPruningRefineByUnion],
	// [PublicFormSubmissionOnPageFilterPruningRefineByUnion],
	// [PublicEmailEventFilterPruningRefineByUnion],
	// [PublicAdsTimeFilterPruningRefineByUnion],
	// [PublicUnifiedEventsFilterPruningRefineByUnion].
	MinOccurrences int64 `json:"minOccurrences"`
	// This field is from variant [PublicPageViewAnalyticsFilterPruningRefineByUnion],
	// [PublicCtaAnalyticsFilterPruningRefineByUnion],
	// [PublicEventAnalyticsFilterPruningRefineByUnion],
	// [PublicFormSubmissionFilterPruningRefineByUnion],
	// [PublicFormSubmissionOnPageFilterPruningRefineByUnion],
	// [PublicEmailEventFilterPruningRefineByUnion],
	// [PublicAdsTimeFilterPruningRefineByUnion],
	// [PublicUnifiedEventsFilterPruningRefineByUnion].
	SetType    string `json:"setType"`
	Comparison string `json:"comparison"`
	// This field is from variant [PublicPageViewAnalyticsFilterPruningRefineByUnion],
	// [PublicCtaAnalyticsFilterPruningRefineByUnion],
	// [PublicEventAnalyticsFilterPruningRefineByUnion],
	// [PublicFormSubmissionFilterPruningRefineByUnion],
	// [PublicFormSubmissionOnPageFilterPruningRefineByUnion],
	// [PublicEmailEventFilterPruningRefineByUnion],
	// [PublicAdsTimeFilterPruningRefineByUnion],
	// [PublicUnifiedEventsFilterPruningRefineByUnion].
	TimeOffset PublicTimeOffset `json:"timeOffset"`
	// This field is from variant [PublicPageViewAnalyticsFilterPruningRefineByUnion],
	// [PublicCtaAnalyticsFilterPruningRefineByUnion],
	// [PublicEventAnalyticsFilterPruningRefineByUnion],
	// [PublicFormSubmissionFilterPruningRefineByUnion],
	// [PublicFormSubmissionOnPageFilterPruningRefineByUnion],
	// [PublicEmailEventFilterPruningRefineByUnion],
	// [PublicAdsTimeFilterPruningRefineByUnion],
	// [PublicUnifiedEventsFilterPruningRefineByUnion].
	LowerBoundOffset PublicTimeOffset `json:"lowerBoundOffset"`
	RangeType        string           `json:"rangeType"`
	// This field is from variant [PublicPageViewAnalyticsFilterPruningRefineByUnion],
	// [PublicCtaAnalyticsFilterPruningRefineByUnion],
	// [PublicEventAnalyticsFilterPruningRefineByUnion],
	// [PublicFormSubmissionFilterPruningRefineByUnion],
	// [PublicFormSubmissionOnPageFilterPruningRefineByUnion],
	// [PublicEmailEventFilterPruningRefineByUnion],
	// [PublicAdsTimeFilterPruningRefineByUnion],
	// [PublicUnifiedEventsFilterPruningRefineByUnion].
	UpperBoundOffset PublicTimeOffset `json:"upperBoundOffset"`
	// This field is from variant [PublicPageViewAnalyticsFilterPruningRefineByUnion],
	// [PublicCtaAnalyticsFilterPruningRefineByUnion],
	// [PublicEventAnalyticsFilterPruningRefineByUnion],
	// [PublicFormSubmissionFilterPruningRefineByUnion],
	// [PublicFormSubmissionOnPageFilterPruningRefineByUnion],
	// [PublicEmailEventFilterPruningRefineByUnion],
	// [PublicAdsTimeFilterPruningRefineByUnion],
	// [PublicUnifiedEventsFilterPruningRefineByUnion].
	Timestamp int64 `json:"timestamp"`
	// This field is from variant [PublicPageViewAnalyticsFilterPruningRefineByUnion],
	// [PublicCtaAnalyticsFilterPruningRefineByUnion],
	// [PublicEventAnalyticsFilterPruningRefineByUnion],
	// [PublicFormSubmissionFilterPruningRefineByUnion],
	// [PublicFormSubmissionOnPageFilterPruningRefineByUnion],
	// [PublicEmailEventFilterPruningRefineByUnion],
	// [PublicAdsTimeFilterPruningRefineByUnion],
	// [PublicUnifiedEventsFilterPruningRefineByUnion].
	LowerTimestamp int64 `json:"lowerTimestamp"`
	// This field is from variant [PublicPageViewAnalyticsFilterPruningRefineByUnion],
	// [PublicCtaAnalyticsFilterPruningRefineByUnion],
	// [PublicEventAnalyticsFilterPruningRefineByUnion],
	// [PublicFormSubmissionFilterPruningRefineByUnion],
	// [PublicFormSubmissionOnPageFilterPruningRefineByUnion],
	// [PublicEmailEventFilterPruningRefineByUnion],
	// [PublicAdsTimeFilterPruningRefineByUnion],
	// [PublicUnifiedEventsFilterPruningRefineByUnion].
	UpperTimestamp               int64  `json:"upperTimestamp"`
	IncludeObjectsWithNoValueSet bool   `json:"includeObjectsWithNoValueSet"`
	OperationType                string `json:"operationType"`
	Operator                     string `json:"operator"`
	// This field is from variant [PublicPageViewAnalyticsFilterPruningRefineByUnion],
	// [PublicCtaAnalyticsFilterPruningRefineByUnion],
	// [PublicEventAnalyticsFilterPruningRefineByUnion],
	// [PublicFormSubmissionFilterPruningRefineByUnion],
	// [PublicFormSubmissionOnPageFilterPruningRefineByUnion],
	// [PublicEmailEventFilterPruningRefineByUnion],
	// [PublicAdsTimeFilterPruningRefineByUnion],
	// [PublicUnifiedEventsFilterPruningRefineByUnion].
	TimePoint PublicTimePointOperationTimePointUnion `json:"timePoint"`
	// This field is from variant [PublicPageViewAnalyticsFilterPruningRefineByUnion],
	// [PublicCtaAnalyticsFilterPruningRefineByUnion],
	// [PublicEventAnalyticsFilterPruningRefineByUnion],
	// [PublicFormSubmissionFilterPruningRefineByUnion],
	// [PublicFormSubmissionOnPageFilterPruningRefineByUnion],
	// [PublicEmailEventFilterPruningRefineByUnion],
	// [PublicAdsTimeFilterPruningRefineByUnion],
	// [PublicUnifiedEventsFilterPruningRefineByUnion].
	EndpointBehavior string `json:"endpointBehavior"`
	PropertyParser   string `json:"propertyParser"`
	// This field is from variant [PublicPageViewAnalyticsFilterPruningRefineByUnion],
	// [PublicCtaAnalyticsFilterPruningRefineByUnion],
	// [PublicEventAnalyticsFilterPruningRefineByUnion],
	// [PublicFormSubmissionFilterPruningRefineByUnion],
	// [PublicFormSubmissionOnPageFilterPruningRefineByUnion],
	// [PublicEmailEventFilterPruningRefineByUnion],
	// [PublicAdsTimeFilterPruningRefineByUnion],
	// [PublicUnifiedEventsFilterPruningRefineByUnion].
	LowerBoundTimePoint PublicRangedTimeOperationLowerBoundTimePointUnion `json:"lowerBoundTimePoint"`
	// This field is from variant [PublicPageViewAnalyticsFilterPruningRefineByUnion],
	// [PublicCtaAnalyticsFilterPruningRefineByUnion],
	// [PublicEventAnalyticsFilterPruningRefineByUnion],
	// [PublicFormSubmissionFilterPruningRefineByUnion],
	// [PublicFormSubmissionOnPageFilterPruningRefineByUnion],
	// [PublicEmailEventFilterPruningRefineByUnion],
	// [PublicAdsTimeFilterPruningRefineByUnion],
	// [PublicUnifiedEventsFilterPruningRefineByUnion].
	UpperBoundTimePoint PublicRangedTimeOperationUpperBoundTimePointUnion `json:"upperBoundTimePoint"`
	// This field is from variant [PublicPageViewAnalyticsFilterPruningRefineByUnion],
	// [PublicCtaAnalyticsFilterPruningRefineByUnion],
	// [PublicEventAnalyticsFilterPruningRefineByUnion],
	// [PublicFormSubmissionFilterPruningRefineByUnion],
	// [PublicFormSubmissionOnPageFilterPruningRefineByUnion],
	// [PublicEmailEventFilterPruningRefineByUnion],
	// [PublicAdsTimeFilterPruningRefineByUnion],
	// [PublicUnifiedEventsFilterPruningRefineByUnion].
	LowerBoundEndpointBehavior string `json:"lowerBoundEndpointBehavior"`
	// This field is from variant [PublicPageViewAnalyticsFilterPruningRefineByUnion],
	// [PublicCtaAnalyticsFilterPruningRefineByUnion],
	// [PublicEventAnalyticsFilterPruningRefineByUnion],
	// [PublicFormSubmissionFilterPruningRefineByUnion],
	// [PublicFormSubmissionOnPageFilterPruningRefineByUnion],
	// [PublicEmailEventFilterPruningRefineByUnion],
	// [PublicAdsTimeFilterPruningRefineByUnion],
	// [PublicUnifiedEventsFilterPruningRefineByUnion].
	UpperBoundEndpointBehavior string `json:"upperBoundEndpointBehavior"`
	JSON                       struct {
		Type                         respjson.Field
		MaxOccurrences               respjson.Field
		MinOccurrences               respjson.Field
		SetType                      respjson.Field
		Comparison                   respjson.Field
		TimeOffset                   respjson.Field
		LowerBoundOffset             respjson.Field
		RangeType                    respjson.Field
		UpperBoundOffset             respjson.Field
		Timestamp                    respjson.Field
		LowerTimestamp               respjson.Field
		UpperTimestamp               respjson.Field
		IncludeObjectsWithNoValueSet respjson.Field
		OperationType                respjson.Field
		Operator                     respjson.Field
		TimePoint                    respjson.Field
		EndpointBehavior             respjson.Field
		PropertyParser               respjson.Field
		LowerBoundTimePoint          respjson.Field
		UpperBoundTimePoint          respjson.Field
		LowerBoundEndpointBehavior   respjson.Field
		UpperBoundEndpointBehavior   respjson.Field
		raw                          string
	} `json:"-"`
}

func (r *PublicPropertyAssociationFilterBranchFilterUnionPruningRefineBy) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PublicPropertyAssociationFilterBranchFilterUnionEventTypeID is an implicit
// subunion of [PublicPropertyAssociationFilterBranchFilterUnion].
// PublicPropertyAssociationFilterBranchFilterUnionEventTypeID provides convenient
// access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [PublicPropertyAssociationFilterBranchFilterUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfInt OfString]
type PublicPropertyAssociationFilterBranchFilterUnionEventTypeID struct {
	// This field will be present if the value is a [int64] instead of an object.
	OfInt int64 `json:",inline"`
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	JSON     struct {
		OfInt    respjson.Field
		OfString respjson.Field
		raw      string
	} `json:"-"`
}

func (r *PublicPropertyAssociationFilterBranchFilterUnionEventTypeID) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties FilterBranches, FilterBranchOperator, FilterBranchType, Filters,
// ObjectTypeID, Operator, PropertyWithObjectID are required.
type PublicPropertyAssociationFilterBranchParam struct {
	FilterBranches []PublicPropertyAssociationFilterBranchFilterBranchUnionParam `json:"filterBranches,omitzero" api:"required"`
	// The logical operator used to combine filters within the branch.
	FilterBranchOperator string `json:"filterBranchOperator" api:"required"`
	// The type of the filter branch (PROPERTY_ASSOCIATION).
	//
	// Any of "PROPERTY_ASSOCIATION".
	FilterBranchType PublicPropertyAssociationFilterBranchFilterBranchType   `json:"filterBranchType,omitzero" api:"required"`
	Filters          []PublicPropertyAssociationFilterBranchFilterUnionParam `json:"filters,omitzero" api:"required"`
	// The ID representing the type of object associated with the filter branch.
	ObjectTypeID string `json:"objectTypeId" api:"required"`
	// Defines the operation to be applied within the filter branch (IN_LIST,
	// NOT_IN_LIST).
	Operator string `json:"operator" api:"required"`
	// The property that is associated with the object ID in the filter branch.
	PropertyWithObjectID string `json:"propertyWithObjectId" api:"required"`
	paramObj
}

func (r PublicPropertyAssociationFilterBranchParam) MarshalJSON() (data []byte, err error) {
	type shadow PublicPropertyAssociationFilterBranchParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PublicPropertyAssociationFilterBranchParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type PublicPropertyAssociationFilterBranchFilterBranchUnionParam struct {
	OfOr                  *PublicOrFilterBranchParam                  `json:",omitzero,inline"`
	OfAnd                 *PublicAndFilterBranchParam                 `json:",omitzero,inline"`
	OfNotAll              *PublicNotAllFilterBranchParam              `json:",omitzero,inline"`
	OfNotAny              *PublicNotAnyFilterBranchParam              `json:",omitzero,inline"`
	OfRestricted          *PublicRestrictedFilterBranchParam          `json:",omitzero,inline"`
	OfUnifiedEvents       *PublicUnifiedEventsFilterBranchParam       `json:",omitzero,inline"`
	OfPropertyAssociation *PublicPropertyAssociationFilterBranchParam `json:",omitzero,inline"`
	OfAssociation         *PublicAssociationFilterBranchParam         `json:",omitzero,inline"`
	paramUnion
}

func (u PublicPropertyAssociationFilterBranchFilterBranchUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfOr,
		u.OfAnd,
		u.OfNotAll,
		u.OfNotAny,
		u.OfRestricted,
		u.OfUnifiedEvents,
		u.OfPropertyAssociation,
		u.OfAssociation)
}
func (u *PublicPropertyAssociationFilterBranchFilterBranchUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func init() {
	apijson.RegisterUnion[PublicPropertyAssociationFilterBranchFilterBranchUnionParam](
		"filterBranchType",
		apijson.Discriminator[PublicOrFilterBranchParam]("OR"),
		apijson.Discriminator[PublicAndFilterBranchParam]("AND"),
		apijson.Discriminator[PublicNotAllFilterBranchParam]("NOT_ALL"),
		apijson.Discriminator[PublicNotAnyFilterBranchParam]("NOT_ANY"),
		apijson.Discriminator[PublicRestrictedFilterBranchParam]("RESTRICTED"),
		apijson.Discriminator[PublicUnifiedEventsFilterBranchParam]("UNIFIED_EVENTS"),
		apijson.Discriminator[PublicPropertyAssociationFilterBranchParam]("PROPERTY_ASSOCIATION"),
		apijson.Discriminator[PublicAssociationFilterBranchParam]("ASSOCIATION"),
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type PublicPropertyAssociationFilterBranchFilterUnionParam struct {
	OfProperty                  *PublicPropertyFilterParam                  `json:",omitzero,inline"`
	OfAssociation               *PublicAssociationInListFilterParam         `json:",omitzero,inline"`
	OfPageView                  *PublicPageViewAnalyticsFilterParam         `json:",omitzero,inline"`
	OfCta                       *PublicCtaAnalyticsFilterParam              `json:",omitzero,inline"`
	OfEvent                     *PublicEventAnalyticsFilterParam            `json:",omitzero,inline"`
	OfFormSubmission            *PublicFormSubmissionFilterParam            `json:",omitzero,inline"`
	OfFormSubmissionOnPage      *PublicFormSubmissionOnPageFilterParam      `json:",omitzero,inline"`
	OfIntegrationEvent          *PublicIntegrationEventFilterParam          `json:",omitzero,inline"`
	OfEmailSubscription         *PublicEmailSubscriptionFilterParam         `json:",omitzero,inline"`
	OfCommunicationSubscription *PublicCommunicationSubscriptionFilterParam `json:",omitzero,inline"`
	OfCampaignInfluenced        *PublicCampaignInfluencedFilterParam        `json:",omitzero,inline"`
	OfSurveyMonkey              *PublicSurveyMonkeyFilterParam              `json:",omitzero,inline"`
	OfSurveyMonkeyValue         *PublicSurveyMonkeyValueFilterParam         `json:",omitzero,inline"`
	OfWebinar                   *PublicWebinarFilterParam                   `json:",omitzero,inline"`
	OfEmailEvent                *PublicEmailEventFilterParam                `json:",omitzero,inline"`
	OfPrivacy                   *PublicPrivacyAnalyticsFilterParam          `json:",omitzero,inline"`
	OfAdsSearch                 *PublicAdsSearchFilterParam                 `json:",omitzero,inline"`
	OfAdsTime                   *PublicAdsTimeFilterParam                   `json:",omitzero,inline"`
	OfInList                    *PublicInListFilterParam                    `json:",omitzero,inline"`
	OfNumAssociations           *PublicNumAssociationsFilterParam           `json:",omitzero,inline"`
	OfUnifiedEvents             *PublicUnifiedEventsFilterParam             `json:",omitzero,inline"`
	OfPropertyAssociation       *PublicPropertyAssociationInListFilterParam `json:",omitzero,inline"`
	OfConstant                  *PublicConstantFilterParam                  `json:",omitzero,inline"`
	paramUnion
}

func (u PublicPropertyAssociationFilterBranchFilterUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfProperty,
		u.OfAssociation,
		u.OfPageView,
		u.OfCta,
		u.OfEvent,
		u.OfFormSubmission,
		u.OfFormSubmissionOnPage,
		u.OfIntegrationEvent,
		u.OfEmailSubscription,
		u.OfCommunicationSubscription,
		u.OfCampaignInfluenced,
		u.OfSurveyMonkey,
		u.OfSurveyMonkeyValue,
		u.OfWebinar,
		u.OfEmailEvent,
		u.OfPrivacy,
		u.OfAdsSearch,
		u.OfAdsTime,
		u.OfInList,
		u.OfNumAssociations,
		u.OfUnifiedEvents,
		u.OfPropertyAssociation,
		u.OfConstant)
}
func (u *PublicPropertyAssociationFilterBranchFilterUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func init() {
	apijson.RegisterUnion[PublicPropertyAssociationFilterBranchFilterUnionParam](
		"filterType",
		apijson.Discriminator[PublicPropertyFilterParam]("PROPERTY"),
		apijson.Discriminator[PublicAssociationInListFilterParam]("ASSOCIATION"),
		apijson.Discriminator[PublicPageViewAnalyticsFilterParam]("PAGE_VIEW"),
		apijson.Discriminator[PublicCtaAnalyticsFilterParam]("CTA"),
		apijson.Discriminator[PublicEventAnalyticsFilterParam]("EVENT"),
		apijson.Discriminator[PublicFormSubmissionFilterParam]("FORM_SUBMISSION"),
		apijson.Discriminator[PublicFormSubmissionOnPageFilterParam]("FORM_SUBMISSION_ON_PAGE"),
		apijson.Discriminator[PublicIntegrationEventFilterParam]("INTEGRATION_EVENT"),
		apijson.Discriminator[PublicEmailSubscriptionFilterParam]("EMAIL_SUBSCRIPTION"),
		apijson.Discriminator[PublicCommunicationSubscriptionFilterParam]("COMMUNICATION_SUBSCRIPTION"),
		apijson.Discriminator[PublicCampaignInfluencedFilterParam]("CAMPAIGN_INFLUENCED"),
		apijson.Discriminator[PublicSurveyMonkeyFilterParam]("SURVEY_MONKEY"),
		apijson.Discriminator[PublicSurveyMonkeyValueFilterParam]("SURVEY_MONKEY_VALUE"),
		apijson.Discriminator[PublicWebinarFilterParam]("WEBINAR"),
		apijson.Discriminator[PublicEmailEventFilterParam]("EMAIL_EVENT"),
		apijson.Discriminator[PublicPrivacyAnalyticsFilterParam]("PRIVACY"),
		apijson.Discriminator[PublicAdsSearchFilterParam]("ADS_SEARCH"),
		apijson.Discriminator[PublicAdsTimeFilterParam]("ADS_TIME"),
		apijson.Discriminator[PublicInListFilterParam]("IN_LIST"),
		apijson.Discriminator[PublicNumAssociationsFilterParam]("NUM_ASSOCIATIONS"),
		apijson.Discriminator[PublicUnifiedEventsFilterParam]("UNIFIED_EVENTS"),
		apijson.Discriminator[PublicPropertyAssociationInListFilterParam]("PROPERTY_ASSOCIATION"),
		apijson.Discriminator[PublicConstantFilterParam]("CONSTANT"),
	)
}

type PublicPropertyAssociationInListFilter struct {
	// Specifies the criteria for refining the filter by coalescing.
	CoalescingRefineBy PublicPropertyAssociationInListFilterCoalescingRefineByUnion `json:"coalescingRefineBy" api:"required"`
	// Indicates the type of filter being applied (PROPERTY_ASSOCIATION).
	//
	// Any of "PROPERTY_ASSOCIATION".
	FilterType PublicPropertyAssociationInListFilterFilterType `json:"filterType" api:"required"`
	// The ID of the list used in the property association filter.
	ListID string `json:"listId" api:"required"`
	// Defines the operation to be applied by the filter (IN_LIST, NOT_IN_LIST).
	Operator string `json:"operator" api:"required"`
	// The property associated with the object ID in the filter.
	PropertyWithObjectID string `json:"propertyWithObjectId" api:"required"`
	// The ID representing the type of object that the property association filter is
	// targeting.
	ToObjectTypeID string `json:"toObjectTypeId"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CoalescingRefineBy   respjson.Field
		FilterType           respjson.Field
		ListID               respjson.Field
		Operator             respjson.Field
		PropertyWithObjectID respjson.Field
		ToObjectTypeID       respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PublicPropertyAssociationInListFilter) RawJSON() string { return r.JSON.raw }
func (r *PublicPropertyAssociationInListFilter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this PublicPropertyAssociationInListFilter to a
// PublicPropertyAssociationInListFilterParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// PublicPropertyAssociationInListFilterParam.Overrides()
func (r PublicPropertyAssociationInListFilter) ToParam() PublicPropertyAssociationInListFilterParam {
	return param.Override[PublicPropertyAssociationInListFilterParam](json.RawMessage(r.RawJSON()))
}

// PublicPropertyAssociationInListFilterCoalescingRefineByUnion contains all
// possible properties and values from [PublicNumOccurrencesRefineBy],
// [PublicSetOccurrencesRefineBy], [PublicRelativeComparativeTimestampRefineBy],
// [PublicRelativeRangedTimestampRefineBy],
// [PublicAbsoluteComparativeTimestampRefineBy],
// [PublicAbsoluteRangedTimestampRefineBy], [PublicAllHistoryRefineBy],
// [PublicTimePointOperation], [PublicRangedTimeOperation].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type PublicPropertyAssociationInListFilterCoalescingRefineByUnion struct {
	Type string `json:"type"`
	// This field is from variant [PublicNumOccurrencesRefineBy].
	MaxOccurrences int64 `json:"maxOccurrences"`
	// This field is from variant [PublicNumOccurrencesRefineBy].
	MinOccurrences int64 `json:"minOccurrences"`
	// This field is from variant [PublicSetOccurrencesRefineBy].
	SetType    string `json:"setType"`
	Comparison string `json:"comparison"`
	// This field is from variant [PublicRelativeComparativeTimestampRefineBy].
	TimeOffset PublicTimeOffset `json:"timeOffset"`
	// This field is from variant [PublicRelativeRangedTimestampRefineBy].
	LowerBoundOffset PublicTimeOffset `json:"lowerBoundOffset"`
	RangeType        string           `json:"rangeType"`
	// This field is from variant [PublicRelativeRangedTimestampRefineBy].
	UpperBoundOffset PublicTimeOffset `json:"upperBoundOffset"`
	// This field is from variant [PublicAbsoluteComparativeTimestampRefineBy].
	Timestamp int64 `json:"timestamp"`
	// This field is from variant [PublicAbsoluteRangedTimestampRefineBy].
	LowerTimestamp int64 `json:"lowerTimestamp"`
	// This field is from variant [PublicAbsoluteRangedTimestampRefineBy].
	UpperTimestamp               int64  `json:"upperTimestamp"`
	IncludeObjectsWithNoValueSet bool   `json:"includeObjectsWithNoValueSet"`
	OperationType                string `json:"operationType"`
	Operator                     string `json:"operator"`
	// This field is from variant [PublicTimePointOperation].
	TimePoint PublicTimePointOperationTimePointUnion `json:"timePoint"`
	// This field is from variant [PublicTimePointOperation].
	EndpointBehavior string `json:"endpointBehavior"`
	PropertyParser   string `json:"propertyParser"`
	// This field is from variant [PublicRangedTimeOperation].
	LowerBoundTimePoint PublicRangedTimeOperationLowerBoundTimePointUnion `json:"lowerBoundTimePoint"`
	// This field is from variant [PublicRangedTimeOperation].
	UpperBoundTimePoint PublicRangedTimeOperationUpperBoundTimePointUnion `json:"upperBoundTimePoint"`
	// This field is from variant [PublicRangedTimeOperation].
	LowerBoundEndpointBehavior string `json:"lowerBoundEndpointBehavior"`
	// This field is from variant [PublicRangedTimeOperation].
	UpperBoundEndpointBehavior string `json:"upperBoundEndpointBehavior"`
	JSON                       struct {
		Type                         respjson.Field
		MaxOccurrences               respjson.Field
		MinOccurrences               respjson.Field
		SetType                      respjson.Field
		Comparison                   respjson.Field
		TimeOffset                   respjson.Field
		LowerBoundOffset             respjson.Field
		RangeType                    respjson.Field
		UpperBoundOffset             respjson.Field
		Timestamp                    respjson.Field
		LowerTimestamp               respjson.Field
		UpperTimestamp               respjson.Field
		IncludeObjectsWithNoValueSet respjson.Field
		OperationType                respjson.Field
		Operator                     respjson.Field
		TimePoint                    respjson.Field
		EndpointBehavior             respjson.Field
		PropertyParser               respjson.Field
		LowerBoundTimePoint          respjson.Field
		UpperBoundTimePoint          respjson.Field
		LowerBoundEndpointBehavior   respjson.Field
		UpperBoundEndpointBehavior   respjson.Field
		raw                          string
	} `json:"-"`
}

func (u PublicPropertyAssociationInListFilterCoalescingRefineByUnion) AsNumOccurrences() (v PublicNumOccurrencesRefineBy) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicPropertyAssociationInListFilterCoalescingRefineByUnion) AsSetOccurrences() (v PublicSetOccurrencesRefineBy) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicPropertyAssociationInListFilterCoalescingRefineByUnion) AsRelativeComparative() (v PublicRelativeComparativeTimestampRefineBy) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicPropertyAssociationInListFilterCoalescingRefineByUnion) AsRelativeRanged() (v PublicRelativeRangedTimestampRefineBy) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicPropertyAssociationInListFilterCoalescingRefineByUnion) AsAbsoluteComparative() (v PublicAbsoluteComparativeTimestampRefineBy) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicPropertyAssociationInListFilterCoalescingRefineByUnion) AsAbsoluteRanged() (v PublicAbsoluteRangedTimestampRefineBy) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicPropertyAssociationInListFilterCoalescingRefineByUnion) AsAllHistory() (v PublicAllHistoryRefineBy) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicPropertyAssociationInListFilterCoalescingRefineByUnion) AsTimePoint() (v PublicTimePointOperation) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicPropertyAssociationInListFilterCoalescingRefineByUnion) AsTimeRanged() (v PublicRangedTimeOperation) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u PublicPropertyAssociationInListFilterCoalescingRefineByUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *PublicPropertyAssociationInListFilterCoalescingRefineByUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Indicates the type of filter being applied (PROPERTY_ASSOCIATION).
type PublicPropertyAssociationInListFilterFilterType string

const (
	PublicPropertyAssociationInListFilterFilterTypePropertyAssociation PublicPropertyAssociationInListFilterFilterType = "PROPERTY_ASSOCIATION"
)

// The properties CoalescingRefineBy, FilterType, ListID, Operator,
// PropertyWithObjectID are required.
type PublicPropertyAssociationInListFilterParam struct {
	// Specifies the criteria for refining the filter by coalescing.
	CoalescingRefineBy PublicPropertyAssociationInListFilterCoalescingRefineByUnionParam `json:"coalescingRefineBy,omitzero" api:"required"`
	// Indicates the type of filter being applied (PROPERTY_ASSOCIATION).
	//
	// Any of "PROPERTY_ASSOCIATION".
	FilterType PublicPropertyAssociationInListFilterFilterType `json:"filterType,omitzero" api:"required"`
	// The ID of the list used in the property association filter.
	ListID string `json:"listId" api:"required"`
	// Defines the operation to be applied by the filter (IN_LIST, NOT_IN_LIST).
	Operator string `json:"operator" api:"required"`
	// The property associated with the object ID in the filter.
	PropertyWithObjectID string `json:"propertyWithObjectId" api:"required"`
	// The ID representing the type of object that the property association filter is
	// targeting.
	ToObjectTypeID param.Opt[string] `json:"toObjectTypeId,omitzero"`
	paramObj
}

func (r PublicPropertyAssociationInListFilterParam) MarshalJSON() (data []byte, err error) {
	type shadow PublicPropertyAssociationInListFilterParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PublicPropertyAssociationInListFilterParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type PublicPropertyAssociationInListFilterCoalescingRefineByUnionParam struct {
	OfNumOccurrences      *PublicNumOccurrencesRefineByParam               `json:",omitzero,inline"`
	OfSetOccurrences      *PublicSetOccurrencesRefineByParam               `json:",omitzero,inline"`
	OfRelativeComparative *PublicRelativeComparativeTimestampRefineByParam `json:",omitzero,inline"`
	OfRelativeRanged      *PublicRelativeRangedTimestampRefineByParam      `json:",omitzero,inline"`
	OfAbsoluteComparative *PublicAbsoluteComparativeTimestampRefineByParam `json:",omitzero,inline"`
	OfAbsoluteRanged      *PublicAbsoluteRangedTimestampRefineByParam      `json:",omitzero,inline"`
	OfAllHistory          *PublicAllHistoryRefineByParam                   `json:",omitzero,inline"`
	OfTimePoint           *PublicTimePointOperationParam                   `json:",omitzero,inline"`
	OfTimeRanged          *PublicRangedTimeOperationParam                  `json:",omitzero,inline"`
	paramUnion
}

func (u PublicPropertyAssociationInListFilterCoalescingRefineByUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfNumOccurrences,
		u.OfSetOccurrences,
		u.OfRelativeComparative,
		u.OfRelativeRanged,
		u.OfAbsoluteComparative,
		u.OfAbsoluteRanged,
		u.OfAllHistory,
		u.OfTimePoint,
		u.OfTimeRanged)
}
func (u *PublicPropertyAssociationInListFilterCoalescingRefineByUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

type PublicPropertyFilter struct {
	// Indicates that the filter (PROPERTY).
	//
	// Any of "PROPERTY".
	FilterType PublicPropertyFilterFilterType `json:"filterType" api:"required"`
	// Defines the operation to be performed on the property, such as comparison or
	// value matching.
	Operation PublicPropertyFilterOperationUnion `json:"operation" api:"required"`
	// Specifies the name of the property that the filter is applied to.
	Property string `json:"property" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FilterType  respjson.Field
		Operation   respjson.Field
		Property    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PublicPropertyFilter) RawJSON() string { return r.JSON.raw }
func (r *PublicPropertyFilter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this PublicPropertyFilter to a PublicPropertyFilterParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// PublicPropertyFilterParam.Overrides()
func (r PublicPropertyFilter) ToParam() PublicPropertyFilterParam {
	return param.Override[PublicPropertyFilterParam](json.RawMessage(r.RawJSON()))
}

// Indicates that the filter (PROPERTY).
type PublicPropertyFilterFilterType string

const (
	PublicPropertyFilterFilterTypeProperty PublicPropertyFilterFilterType = "PROPERTY"
)

// PublicPropertyFilterOperationUnion contains all possible properties and values
// from [PublicBoolPropertyOperation], [PublicNumberPropertyOperation],
// [PublicStringPropertyOperation], [PublicDateTimePropertyOperation],
// [PublicRangedDatePropertyOperation],
// [PublicComparativePropertyUpdatedOperation],
// [PublicComparativeDatePropertyOperation],
// [PublicRollingDateRangePropertyOperation],
// [PublicRollingPropertyUpdatedOperation], [PublicEnumerationPropertyOperation],
// [PublicAllPropertyTypesOperation], [PublicRangedNumberPropertyOperation],
// [PublicMultiStringPropertyOperation], [PublicDatePropertyOperation],
// [PublicCalendarDatePropertyOperation], [PublicTimePointOperation],
// [PublicRangedTimeOperation].
//
// Use the [PublicPropertyFilterOperationUnion.AsAny] method to switch on the
// variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type PublicPropertyFilterOperationUnion struct {
	IncludeObjectsWithNoValueSet bool `json:"includeObjectsWithNoValueSet"`
	// Any of "BOOL", "NUMBER", "STRING", "DATETIME", "RANGED_DATE",
	// "COMPARATIVE_PROPERTY_UPDATED", "COMPARATIVE_DATE", "ROLLING_DATE_RANGE",
	// "ROLLING_PROPERTY_UPDATED", "ENUMERATION", "ALL_PROPERTY", "NUMBER_RANGED",
	// "MULTISTRING", "DATE", "CALENDAR_DATE", "TIME_POINT", "TIME_RANGED".
	OperationType string `json:"operationType"`
	Operator      string `json:"operator"`
	// This field is a union of [bool], [float64], [string]
	Value                      PublicPropertyFilterOperationUnionValue `json:"value"`
	RequiresTimeZoneConversion bool                                    `json:"requiresTimeZoneConversion"`
	// This field is from variant [PublicDateTimePropertyOperation].
	Timestamp              int64    `json:"timestamp"`
	LowerBound             int64    `json:"lowerBound"`
	UpperBound             int64    `json:"upperBound"`
	ComparisonPropertyName string   `json:"comparisonPropertyName"`
	DefaultComparisonValue string   `json:"defaultComparisonValue"`
	NumberOfDays           int64    `json:"numberOfDays"`
	Values                 []string `json:"values"`
	// This field is from variant [PublicDatePropertyOperation].
	Day int64 `json:"day"`
	// This field is from variant [PublicDatePropertyOperation].
	Month string `json:"month"`
	// This field is from variant [PublicDatePropertyOperation].
	Year int64 `json:"year"`
	// This field is from variant [PublicCalendarDatePropertyOperation].
	TimeUnit string `json:"timeUnit"`
	// This field is from variant [PublicCalendarDatePropertyOperation].
	FiscalYearStart PublicCalendarDatePropertyOperationFiscalYearStart `json:"fiscalYearStart"`
	// This field is from variant [PublicCalendarDatePropertyOperation].
	TimeUnitCount int64 `json:"timeUnitCount"`
	// This field is from variant [PublicCalendarDatePropertyOperation].
	UseFiscalYear bool `json:"useFiscalYear"`
	// This field is from variant [PublicTimePointOperation].
	TimePoint PublicTimePointOperationTimePointUnion `json:"timePoint"`
	Type      string                                 `json:"type"`
	// This field is from variant [PublicTimePointOperation].
	EndpointBehavior string `json:"endpointBehavior"`
	PropertyParser   string `json:"propertyParser"`
	// This field is from variant [PublicRangedTimeOperation].
	LowerBoundTimePoint PublicRangedTimeOperationLowerBoundTimePointUnion `json:"lowerBoundTimePoint"`
	// This field is from variant [PublicRangedTimeOperation].
	UpperBoundTimePoint PublicRangedTimeOperationUpperBoundTimePointUnion `json:"upperBoundTimePoint"`
	// This field is from variant [PublicRangedTimeOperation].
	LowerBoundEndpointBehavior string `json:"lowerBoundEndpointBehavior"`
	// This field is from variant [PublicRangedTimeOperation].
	UpperBoundEndpointBehavior string `json:"upperBoundEndpointBehavior"`
	JSON                       struct {
		IncludeObjectsWithNoValueSet respjson.Field
		OperationType                respjson.Field
		Operator                     respjson.Field
		Value                        respjson.Field
		RequiresTimeZoneConversion   respjson.Field
		Timestamp                    respjson.Field
		LowerBound                   respjson.Field
		UpperBound                   respjson.Field
		ComparisonPropertyName       respjson.Field
		DefaultComparisonValue       respjson.Field
		NumberOfDays                 respjson.Field
		Values                       respjson.Field
		Day                          respjson.Field
		Month                        respjson.Field
		Year                         respjson.Field
		TimeUnit                     respjson.Field
		FiscalYearStart              respjson.Field
		TimeUnitCount                respjson.Field
		UseFiscalYear                respjson.Field
		TimePoint                    respjson.Field
		Type                         respjson.Field
		EndpointBehavior             respjson.Field
		PropertyParser               respjson.Field
		LowerBoundTimePoint          respjson.Field
		UpperBoundTimePoint          respjson.Field
		LowerBoundEndpointBehavior   respjson.Field
		UpperBoundEndpointBehavior   respjson.Field
		raw                          string
	} `json:"-"`
}

// anyPublicPropertyFilterOperation is implemented by each variant of
// [PublicPropertyFilterOperationUnion] to add type safety for the return type of
// [PublicPropertyFilterOperationUnion.AsAny]
type anyPublicPropertyFilterOperation interface {
	implPublicPropertyFilterOperationUnion()
}

func (PublicBoolPropertyOperation) implPublicPropertyFilterOperationUnion()               {}
func (PublicNumberPropertyOperation) implPublicPropertyFilterOperationUnion()             {}
func (PublicStringPropertyOperation) implPublicPropertyFilterOperationUnion()             {}
func (PublicDateTimePropertyOperation) implPublicPropertyFilterOperationUnion()           {}
func (PublicRangedDatePropertyOperation) implPublicPropertyFilterOperationUnion()         {}
func (PublicComparativePropertyUpdatedOperation) implPublicPropertyFilterOperationUnion() {}
func (PublicComparativeDatePropertyOperation) implPublicPropertyFilterOperationUnion()    {}
func (PublicRollingDateRangePropertyOperation) implPublicPropertyFilterOperationUnion()   {}
func (PublicRollingPropertyUpdatedOperation) implPublicPropertyFilterOperationUnion()     {}
func (PublicEnumerationPropertyOperation) implPublicPropertyFilterOperationUnion()        {}
func (PublicAllPropertyTypesOperation) implPublicPropertyFilterOperationUnion()           {}
func (PublicRangedNumberPropertyOperation) implPublicPropertyFilterOperationUnion()       {}
func (PublicMultiStringPropertyOperation) implPublicPropertyFilterOperationUnion()        {}
func (PublicDatePropertyOperation) implPublicPropertyFilterOperationUnion()               {}
func (PublicCalendarDatePropertyOperation) implPublicPropertyFilterOperationUnion()       {}
func (PublicTimePointOperation) implPublicPropertyFilterOperationUnion()                  {}
func (PublicRangedTimeOperation) implPublicPropertyFilterOperationUnion()                 {}

// Use the following switch statement to find the correct variant
//
//	switch variant := PublicPropertyFilterOperationUnion.AsAny().(type) {
//	case crm.PublicBoolPropertyOperation:
//	case crm.PublicNumberPropertyOperation:
//	case crm.PublicStringPropertyOperation:
//	case crm.PublicDateTimePropertyOperation:
//	case crm.PublicRangedDatePropertyOperation:
//	case crm.PublicComparativePropertyUpdatedOperation:
//	case crm.PublicComparativeDatePropertyOperation:
//	case crm.PublicRollingDateRangePropertyOperation:
//	case crm.PublicRollingPropertyUpdatedOperation:
//	case crm.PublicEnumerationPropertyOperation:
//	case crm.PublicAllPropertyTypesOperation:
//	case crm.PublicRangedNumberPropertyOperation:
//	case crm.PublicMultiStringPropertyOperation:
//	case crm.PublicDatePropertyOperation:
//	case crm.PublicCalendarDatePropertyOperation:
//	case crm.PublicTimePointOperation:
//	case crm.PublicRangedTimeOperation:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u PublicPropertyFilterOperationUnion) AsAny() anyPublicPropertyFilterOperation {
	switch u.OperationType {
	case "BOOL":
		return u.AsBool()
	case "NUMBER":
		return u.AsNumber()
	case "STRING":
		return u.AsString()
	case "DATETIME":
		return u.AsDatetime()
	case "RANGED_DATE":
		return u.AsRangedDate()
	case "COMPARATIVE_PROPERTY_UPDATED":
		return u.AsComparativePropertyUpdated()
	case "COMPARATIVE_DATE":
		return u.AsComparativeDate()
	case "ROLLING_DATE_RANGE":
		return u.AsRollingDateRange()
	case "ROLLING_PROPERTY_UPDATED":
		return u.AsRollingPropertyUpdated()
	case "ENUMERATION":
		return u.AsEnumeration()
	case "ALL_PROPERTY":
		return u.AsAllProperty()
	case "NUMBER_RANGED":
		return u.AsNumberRanged()
	case "MULTISTRING":
		return u.AsMultistring()
	case "DATE":
		return u.AsDate()
	case "CALENDAR_DATE":
		return u.AsCalendarDate()
	case "TIME_POINT":
		return u.AsTimePoint()
	case "TIME_RANGED":
		return u.AsTimeRanged()
	}
	return nil
}

func (u PublicPropertyFilterOperationUnion) AsBool() (v PublicBoolPropertyOperation) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicPropertyFilterOperationUnion) AsNumber() (v PublicNumberPropertyOperation) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicPropertyFilterOperationUnion) AsString() (v PublicStringPropertyOperation) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicPropertyFilterOperationUnion) AsDatetime() (v PublicDateTimePropertyOperation) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicPropertyFilterOperationUnion) AsRangedDate() (v PublicRangedDatePropertyOperation) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicPropertyFilterOperationUnion) AsComparativePropertyUpdated() (v PublicComparativePropertyUpdatedOperation) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicPropertyFilterOperationUnion) AsComparativeDate() (v PublicComparativeDatePropertyOperation) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicPropertyFilterOperationUnion) AsRollingDateRange() (v PublicRollingDateRangePropertyOperation) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicPropertyFilterOperationUnion) AsRollingPropertyUpdated() (v PublicRollingPropertyUpdatedOperation) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicPropertyFilterOperationUnion) AsEnumeration() (v PublicEnumerationPropertyOperation) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicPropertyFilterOperationUnion) AsAllProperty() (v PublicAllPropertyTypesOperation) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicPropertyFilterOperationUnion) AsNumberRanged() (v PublicRangedNumberPropertyOperation) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicPropertyFilterOperationUnion) AsMultistring() (v PublicMultiStringPropertyOperation) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicPropertyFilterOperationUnion) AsDate() (v PublicDatePropertyOperation) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicPropertyFilterOperationUnion) AsCalendarDate() (v PublicCalendarDatePropertyOperation) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicPropertyFilterOperationUnion) AsTimePoint() (v PublicTimePointOperation) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicPropertyFilterOperationUnion) AsTimeRanged() (v PublicRangedTimeOperation) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u PublicPropertyFilterOperationUnion) RawJSON() string { return u.JSON.raw }

func (r *PublicPropertyFilterOperationUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PublicPropertyFilterOperationUnionValue is an implicit subunion of
// [PublicPropertyFilterOperationUnion]. PublicPropertyFilterOperationUnionValue
// provides convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [PublicPropertyFilterOperationUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfBool OfFloat OfString]
type PublicPropertyFilterOperationUnionValue struct {
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	JSON     struct {
		OfBool   respjson.Field
		OfFloat  respjson.Field
		OfString respjson.Field
		raw      string
	} `json:"-"`
}

func (r *PublicPropertyFilterOperationUnionValue) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties FilterType, Operation, Property are required.
type PublicPropertyFilterParam struct {
	// Indicates that the filter (PROPERTY).
	//
	// Any of "PROPERTY".
	FilterType PublicPropertyFilterFilterType `json:"filterType,omitzero" api:"required"`
	// Defines the operation to be performed on the property, such as comparison or
	// value matching.
	Operation PublicPropertyFilterOperationUnionParam `json:"operation,omitzero" api:"required"`
	// Specifies the name of the property that the filter is applied to.
	Property string `json:"property" api:"required"`
	paramObj
}

func (r PublicPropertyFilterParam) MarshalJSON() (data []byte, err error) {
	type shadow PublicPropertyFilterParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PublicPropertyFilterParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type PublicPropertyFilterOperationUnionParam struct {
	OfBool                       *PublicBoolPropertyOperationParam               `json:",omitzero,inline"`
	OfNumber                     *PublicNumberPropertyOperationParam             `json:",omitzero,inline"`
	OfString                     *PublicStringPropertyOperationParam             `json:",omitzero,inline"`
	OfDatetime                   *PublicDateTimePropertyOperationParam           `json:",omitzero,inline"`
	OfRangedDate                 *PublicRangedDatePropertyOperationParam         `json:",omitzero,inline"`
	OfComparativePropertyUpdated *PublicComparativePropertyUpdatedOperationParam `json:",omitzero,inline"`
	OfComparativeDate            *PublicComparativeDatePropertyOperationParam    `json:",omitzero,inline"`
	OfRollingDateRange           *PublicRollingDateRangePropertyOperationParam   `json:",omitzero,inline"`
	OfRollingPropertyUpdated     *PublicRollingPropertyUpdatedOperationParam     `json:",omitzero,inline"`
	OfEnumeration                *PublicEnumerationPropertyOperationParam        `json:",omitzero,inline"`
	OfAllProperty                *PublicAllPropertyTypesOperationParam           `json:",omitzero,inline"`
	OfNumberRanged               *PublicRangedNumberPropertyOperationParam       `json:",omitzero,inline"`
	OfMultistring                *PublicMultiStringPropertyOperationParam        `json:",omitzero,inline"`
	OfDate                       *PublicDatePropertyOperationParam               `json:",omitzero,inline"`
	OfCalendarDate               *PublicCalendarDatePropertyOperationParam       `json:",omitzero,inline"`
	OfTimePoint                  *PublicTimePointOperationParam                  `json:",omitzero,inline"`
	OfTimeRanged                 *PublicRangedTimeOperationParam                 `json:",omitzero,inline"`
	paramUnion
}

func (u PublicPropertyFilterOperationUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfBool,
		u.OfNumber,
		u.OfString,
		u.OfDatetime,
		u.OfRangedDate,
		u.OfComparativePropertyUpdated,
		u.OfComparativeDate,
		u.OfRollingDateRange,
		u.OfRollingPropertyUpdated,
		u.OfEnumeration,
		u.OfAllProperty,
		u.OfNumberRanged,
		u.OfMultistring,
		u.OfDate,
		u.OfCalendarDate,
		u.OfTimePoint,
		u.OfTimeRanged)
}
func (u *PublicPropertyFilterOperationUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func init() {
	apijson.RegisterUnion[PublicPropertyFilterOperationUnionParam](
		"operationType",
		apijson.Discriminator[PublicBoolPropertyOperationParam]("BOOL"),
		apijson.Discriminator[PublicNumberPropertyOperationParam]("NUMBER"),
		apijson.Discriminator[PublicStringPropertyOperationParam]("STRING"),
		apijson.Discriminator[PublicDateTimePropertyOperationParam]("DATETIME"),
		apijson.Discriminator[PublicRangedDatePropertyOperationParam]("RANGED_DATE"),
		apijson.Discriminator[PublicComparativePropertyUpdatedOperationParam]("COMPARATIVE_PROPERTY_UPDATED"),
		apijson.Discriminator[PublicComparativeDatePropertyOperationParam]("COMPARATIVE_DATE"),
		apijson.Discriminator[PublicRollingDateRangePropertyOperationParam]("ROLLING_DATE_RANGE"),
		apijson.Discriminator[PublicRollingPropertyUpdatedOperationParam]("ROLLING_PROPERTY_UPDATED"),
		apijson.Discriminator[PublicEnumerationPropertyOperationParam]("ENUMERATION"),
		apijson.Discriminator[PublicAllPropertyTypesOperationParam]("ALL_PROPERTY"),
		apijson.Discriminator[PublicRangedNumberPropertyOperationParam]("NUMBER_RANGED"),
		apijson.Discriminator[PublicMultiStringPropertyOperationParam]("MULTISTRING"),
		apijson.Discriminator[PublicDatePropertyOperationParam]("DATE"),
		apijson.Discriminator[PublicCalendarDatePropertyOperationParam]("CALENDAR_DATE"),
		apijson.Discriminator[PublicTimePointOperationParam]("TIME_POINT"),
		apijson.Discriminator[PublicRangedTimeOperationParam]("TIME_RANGED"),
	)
}

type PublicPropertyReferencedTime struct {
	// Specifies the name of the property that the time reference is applied to.
	Property string `json:"property" api:"required"`
	// Specifies the type of reference for the property (VALUE, UPDATED_AT,
	// ANNIVERSARY, VALUE_WITH_ZONE_SAME_LOCAL_CONVERSION,
	// ANNIVERSARY_WITH_ZONE_SAME_LOCAL_CONVERSION).
	ReferenceType string `json:"referenceType" api:"required"`
	// Defines the type of time (PROPERTY_REFERENCED).
	//
	// Any of "PROPERTY_REFERENCED".
	TimeType PublicPropertyReferencedTimeTimeType `json:"timeType" api:"required"`
	// Indicates the identifier for the time zone associated with the property.
	ZoneID string `json:"zoneId" api:"required"`
	// Specifies the source of the time zone information for the property (CUSTOM,
	// USER, PORTAL).
	TimezoneSource string `json:"timezoneSource"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Property       respjson.Field
		ReferenceType  respjson.Field
		TimeType       respjson.Field
		ZoneID         respjson.Field
		TimezoneSource respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PublicPropertyReferencedTime) RawJSON() string { return r.JSON.raw }
func (r *PublicPropertyReferencedTime) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this PublicPropertyReferencedTime to a
// PublicPropertyReferencedTimeParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// PublicPropertyReferencedTimeParam.Overrides()
func (r PublicPropertyReferencedTime) ToParam() PublicPropertyReferencedTimeParam {
	return param.Override[PublicPropertyReferencedTimeParam](json.RawMessage(r.RawJSON()))
}

// Defines the type of time (PROPERTY_REFERENCED).
type PublicPropertyReferencedTimeTimeType string

const (
	PublicPropertyReferencedTimeTimeTypePropertyReferenced PublicPropertyReferencedTimeTimeType = "PROPERTY_REFERENCED"
)

// The properties Property, ReferenceType, TimeType, ZoneID are required.
type PublicPropertyReferencedTimeParam struct {
	// Specifies the name of the property that the time reference is applied to.
	Property string `json:"property" api:"required"`
	// Specifies the type of reference for the property (VALUE, UPDATED_AT,
	// ANNIVERSARY, VALUE_WITH_ZONE_SAME_LOCAL_CONVERSION,
	// ANNIVERSARY_WITH_ZONE_SAME_LOCAL_CONVERSION).
	ReferenceType string `json:"referenceType" api:"required"`
	// Defines the type of time (PROPERTY_REFERENCED).
	//
	// Any of "PROPERTY_REFERENCED".
	TimeType PublicPropertyReferencedTimeTimeType `json:"timeType,omitzero" api:"required"`
	// Indicates the identifier for the time zone associated with the property.
	ZoneID string `json:"zoneId" api:"required"`
	// Specifies the source of the time zone information for the property (CUSTOM,
	// USER, PORTAL).
	TimezoneSource param.Opt[string] `json:"timezoneSource,omitzero"`
	paramObj
}

func (r PublicPropertyReferencedTimeParam) MarshalJSON() (data []byte, err error) {
	type shadow PublicPropertyReferencedTimeParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PublicPropertyReferencedTimeParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PublicQuarterReference struct {
	// The day component of the quarter reference.
	Day int64 `json:"day" api:"required"`
	// The month component of the quarter reference.
	Month int64 `json:"month" api:"required"`
	// Indicates the type of reference (QUARTER).
	//
	// Any of "QUARTER".
	ReferenceType PublicQuarterReferenceReferenceType `json:"referenceType" api:"required"`
	// The hour component of the quarter reference.
	Hour int64 `json:"hour"`
	// The millisecond component of the quarter reference.
	Millisecond int64 `json:"millisecond"`
	// The minute component of the quarter reference.
	Minute int64 `json:"minute"`
	// The second component of the quarter reference.
	Second int64 `json:"second"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Day           respjson.Field
		Month         respjson.Field
		ReferenceType respjson.Field
		Hour          respjson.Field
		Millisecond   respjson.Field
		Minute        respjson.Field
		Second        respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PublicQuarterReference) RawJSON() string { return r.JSON.raw }
func (r *PublicQuarterReference) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this PublicQuarterReference to a PublicQuarterReferenceParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// PublicQuarterReferenceParam.Overrides()
func (r PublicQuarterReference) ToParam() PublicQuarterReferenceParam {
	return param.Override[PublicQuarterReferenceParam](json.RawMessage(r.RawJSON()))
}

// Indicates the type of reference (QUARTER).
type PublicQuarterReferenceReferenceType string

const (
	PublicQuarterReferenceReferenceTypeQuarter PublicQuarterReferenceReferenceType = "QUARTER"
)

// The properties Day, Month, ReferenceType are required.
type PublicQuarterReferenceParam struct {
	// The day component of the quarter reference.
	Day int64 `json:"day" api:"required"`
	// The month component of the quarter reference.
	Month int64 `json:"month" api:"required"`
	// Indicates the type of reference (QUARTER).
	//
	// Any of "QUARTER".
	ReferenceType PublicQuarterReferenceReferenceType `json:"referenceType,omitzero" api:"required"`
	// The hour component of the quarter reference.
	Hour param.Opt[int64] `json:"hour,omitzero"`
	// The millisecond component of the quarter reference.
	Millisecond param.Opt[int64] `json:"millisecond,omitzero"`
	// The minute component of the quarter reference.
	Minute param.Opt[int64] `json:"minute,omitzero"`
	// The second component of the quarter reference.
	Second param.Opt[int64] `json:"second,omitzero"`
	paramObj
}

func (r PublicQuarterReferenceParam) MarshalJSON() (data []byte, err error) {
	type shadow PublicQuarterReferenceParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PublicQuarterReferenceParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PublicRangedDatePropertyOperation struct {
	// Specifies whether objects without a set value should be included in the
	// operation.
	IncludeObjectsWithNoValueSet bool `json:"includeObjectsWithNoValueSet" api:"required"`
	// The lower limit of the date range for the operation.
	LowerBound int64 `json:"lowerBound" api:"required"`
	// Specifies the type of operation (RANGED_DATE).
	//
	// Any of "RANGED_DATE".
	OperationType PublicRangedDatePropertyOperationOperationType `json:"operationType" api:"required"`
	// Defines the operation to be applied in the ranged date property operation
	// (IS_BETWEEN, IS_NOT_BETWEEN).
	Operator string `json:"operator" api:"required"`
	// Indicates whether the operation requires conversion to a different time zone.
	RequiresTimeZoneConversion bool `json:"requiresTimeZoneConversion" api:"required"`
	// The upper limit of the date range for the operation.
	UpperBound int64 `json:"upperBound" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		IncludeObjectsWithNoValueSet respjson.Field
		LowerBound                   respjson.Field
		OperationType                respjson.Field
		Operator                     respjson.Field
		RequiresTimeZoneConversion   respjson.Field
		UpperBound                   respjson.Field
		ExtraFields                  map[string]respjson.Field
		raw                          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PublicRangedDatePropertyOperation) RawJSON() string { return r.JSON.raw }
func (r *PublicRangedDatePropertyOperation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this PublicRangedDatePropertyOperation to a
// PublicRangedDatePropertyOperationParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// PublicRangedDatePropertyOperationParam.Overrides()
func (r PublicRangedDatePropertyOperation) ToParam() PublicRangedDatePropertyOperationParam {
	return param.Override[PublicRangedDatePropertyOperationParam](json.RawMessage(r.RawJSON()))
}

// Specifies the type of operation (RANGED_DATE).
type PublicRangedDatePropertyOperationOperationType string

const (
	PublicRangedDatePropertyOperationOperationTypeRangedDate PublicRangedDatePropertyOperationOperationType = "RANGED_DATE"
)

// The properties IncludeObjectsWithNoValueSet, LowerBound, OperationType,
// Operator, RequiresTimeZoneConversion, UpperBound are required.
type PublicRangedDatePropertyOperationParam struct {
	// Specifies whether objects without a set value should be included in the
	// operation.
	IncludeObjectsWithNoValueSet bool `json:"includeObjectsWithNoValueSet" api:"required"`
	// The lower limit of the date range for the operation.
	LowerBound int64 `json:"lowerBound" api:"required"`
	// Specifies the type of operation (RANGED_DATE).
	//
	// Any of "RANGED_DATE".
	OperationType PublicRangedDatePropertyOperationOperationType `json:"operationType,omitzero" api:"required"`
	// Defines the operation to be applied in the ranged date property operation
	// (IS_BETWEEN, IS_NOT_BETWEEN).
	Operator string `json:"operator" api:"required"`
	// Indicates whether the operation requires conversion to a different time zone.
	RequiresTimeZoneConversion bool `json:"requiresTimeZoneConversion" api:"required"`
	// The upper limit of the date range for the operation.
	UpperBound int64 `json:"upperBound" api:"required"`
	paramObj
}

func (r PublicRangedDatePropertyOperationParam) MarshalJSON() (data []byte, err error) {
	type shadow PublicRangedDatePropertyOperationParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PublicRangedDatePropertyOperationParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PublicRangedNumberPropertyOperation struct {
	// Indicates whether objects with no value set for the property should be included
	// in the operation.
	IncludeObjectsWithNoValueSet bool `json:"includeObjectsWithNoValueSet" api:"required"`
	// The lower limit of the number range for the operation.
	LowerBound int64 `json:"lowerBound" api:"required"`
	// Specifies the type of operation (NUMBER_RANGED).
	//
	// Any of "NUMBER_RANGED".
	OperationType PublicRangedNumberPropertyOperationOperationType `json:"operationType" api:"required"`
	// Defines the operation to be applied in the ranged number property operation
	// (IS_BETWEEN, IS_NOT_BETWEEN).
	Operator string `json:"operator" api:"required"`
	// The upper limit of the number range for the operation.
	UpperBound int64 `json:"upperBound" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		IncludeObjectsWithNoValueSet respjson.Field
		LowerBound                   respjson.Field
		OperationType                respjson.Field
		Operator                     respjson.Field
		UpperBound                   respjson.Field
		ExtraFields                  map[string]respjson.Field
		raw                          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PublicRangedNumberPropertyOperation) RawJSON() string { return r.JSON.raw }
func (r *PublicRangedNumberPropertyOperation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this PublicRangedNumberPropertyOperation to a
// PublicRangedNumberPropertyOperationParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// PublicRangedNumberPropertyOperationParam.Overrides()
func (r PublicRangedNumberPropertyOperation) ToParam() PublicRangedNumberPropertyOperationParam {
	return param.Override[PublicRangedNumberPropertyOperationParam](json.RawMessage(r.RawJSON()))
}

// Specifies the type of operation (NUMBER_RANGED).
type PublicRangedNumberPropertyOperationOperationType string

const (
	PublicRangedNumberPropertyOperationOperationTypeNumberRanged PublicRangedNumberPropertyOperationOperationType = "NUMBER_RANGED"
)

// The properties IncludeObjectsWithNoValueSet, LowerBound, OperationType,
// Operator, UpperBound are required.
type PublicRangedNumberPropertyOperationParam struct {
	// Indicates whether objects with no value set for the property should be included
	// in the operation.
	IncludeObjectsWithNoValueSet bool `json:"includeObjectsWithNoValueSet" api:"required"`
	// The lower limit of the number range for the operation.
	LowerBound int64 `json:"lowerBound" api:"required"`
	// Specifies the type of operation (NUMBER_RANGED).
	//
	// Any of "NUMBER_RANGED".
	OperationType PublicRangedNumberPropertyOperationOperationType `json:"operationType,omitzero" api:"required"`
	// Defines the operation to be applied in the ranged number property operation
	// (IS_BETWEEN, IS_NOT_BETWEEN).
	Operator string `json:"operator" api:"required"`
	// The upper limit of the number range for the operation.
	UpperBound int64 `json:"upperBound" api:"required"`
	paramObj
}

func (r PublicRangedNumberPropertyOperationParam) MarshalJSON() (data []byte, err error) {
	type shadow PublicRangedNumberPropertyOperationParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PublicRangedNumberPropertyOperationParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PublicRangedTimeOperation struct {
	// Indicates whether objects with no value set for the property should be included
	// in the operation.
	IncludeObjectsWithNoValueSet bool `json:"includeObjectsWithNoValueSet" api:"required"`
	// Defines the lower bound time point for the operation.
	LowerBoundTimePoint PublicRangedTimeOperationLowerBoundTimePointUnion `json:"lowerBoundTimePoint" api:"required"`
	// Specifies the type of operation (TIME_RANGED).
	//
	// Any of "TIME_RANGED".
	OperationType PublicRangedTimeOperationOperationType `json:"operationType" api:"required"`
	// Defines the operation to be applied within the time range (IS_BETWEEN,
	// IS_NOT_BETWEEN).
	Operator string `json:"operator" api:"required"`
	// Specifies the type of operation (TIME_RANGED).
	//
	// Any of "TIME_RANGED".
	Type PublicRangedTimeOperationType `json:"type" api:"required"`
	// Defines the upper bound time point for the operation.
	UpperBoundTimePoint PublicRangedTimeOperationUpperBoundTimePointUnion `json:"upperBoundTimePoint" api:"required"`
	// Describes the behavior at the lower bound endpoint of the time range.
	LowerBoundEndpointBehavior string `json:"lowerBoundEndpointBehavior"`
	// Specifies the parser used for the property in the operation.
	PropertyParser string `json:"propertyParser"`
	// Describes the behavior at the upper bound endpoint of the time range.
	UpperBoundEndpointBehavior string `json:"upperBoundEndpointBehavior"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		IncludeObjectsWithNoValueSet respjson.Field
		LowerBoundTimePoint          respjson.Field
		OperationType                respjson.Field
		Operator                     respjson.Field
		Type                         respjson.Field
		UpperBoundTimePoint          respjson.Field
		LowerBoundEndpointBehavior   respjson.Field
		PropertyParser               respjson.Field
		UpperBoundEndpointBehavior   respjson.Field
		ExtraFields                  map[string]respjson.Field
		raw                          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PublicRangedTimeOperation) RawJSON() string { return r.JSON.raw }
func (r *PublicRangedTimeOperation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this PublicRangedTimeOperation to a
// PublicRangedTimeOperationParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// PublicRangedTimeOperationParam.Overrides()
func (r PublicRangedTimeOperation) ToParam() PublicRangedTimeOperationParam {
	return param.Override[PublicRangedTimeOperationParam](json.RawMessage(r.RawJSON()))
}

// PublicRangedTimeOperationLowerBoundTimePointUnion contains all possible
// properties and values from [PublicDatePoint], [PublicIndexedTimePoint],
// [PublicPropertyReferencedTime].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type PublicRangedTimeOperationLowerBoundTimePointUnion struct {
	// This field is from variant [PublicDatePoint].
	Day int64 `json:"day"`
	// This field is from variant [PublicDatePoint].
	Month    int64  `json:"month"`
	TimeType string `json:"timeType"`
	// This field is from variant [PublicDatePoint].
	Year   int64  `json:"year"`
	ZoneID string `json:"zoneId"`
	// This field is from variant [PublicDatePoint].
	Hour int64 `json:"hour"`
	// This field is from variant [PublicDatePoint].
	Millisecond int64 `json:"millisecond"`
	// This field is from variant [PublicDatePoint].
	Minute int64 `json:"minute"`
	// This field is from variant [PublicDatePoint].
	Second         int64  `json:"second"`
	TimezoneSource string `json:"timezoneSource"`
	// This field is from variant [PublicIndexedTimePoint].
	IndexReference PublicIndexedTimePointIndexReferenceUnion `json:"indexReference"`
	// This field is from variant [PublicIndexedTimePoint].
	Offset PublicIndexOffset `json:"offset"`
	// This field is from variant [PublicPropertyReferencedTime].
	Property string `json:"property"`
	// This field is from variant [PublicPropertyReferencedTime].
	ReferenceType string `json:"referenceType"`
	JSON          struct {
		Day            respjson.Field
		Month          respjson.Field
		TimeType       respjson.Field
		Year           respjson.Field
		ZoneID         respjson.Field
		Hour           respjson.Field
		Millisecond    respjson.Field
		Minute         respjson.Field
		Second         respjson.Field
		TimezoneSource respjson.Field
		IndexReference respjson.Field
		Offset         respjson.Field
		Property       respjson.Field
		ReferenceType  respjson.Field
		raw            string
	} `json:"-"`
}

func (u PublicRangedTimeOperationLowerBoundTimePointUnion) AsDate() (v PublicDatePoint) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicRangedTimeOperationLowerBoundTimePointUnion) AsIndexed() (v PublicIndexedTimePoint) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicRangedTimeOperationLowerBoundTimePointUnion) AsPropertyReferenced() (v PublicPropertyReferencedTime) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u PublicRangedTimeOperationLowerBoundTimePointUnion) RawJSON() string { return u.JSON.raw }

func (r *PublicRangedTimeOperationLowerBoundTimePointUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Specifies the type of operation (TIME_RANGED).
type PublicRangedTimeOperationOperationType string

const (
	PublicRangedTimeOperationOperationTypeTimeRanged PublicRangedTimeOperationOperationType = "TIME_RANGED"
)

// Specifies the type of operation (TIME_RANGED).
type PublicRangedTimeOperationType string

const (
	PublicRangedTimeOperationTypeTimeRanged PublicRangedTimeOperationType = "TIME_RANGED"
)

// PublicRangedTimeOperationUpperBoundTimePointUnion contains all possible
// properties and values from [PublicDatePoint], [PublicIndexedTimePoint],
// [PublicPropertyReferencedTime].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type PublicRangedTimeOperationUpperBoundTimePointUnion struct {
	// This field is from variant [PublicDatePoint].
	Day int64 `json:"day"`
	// This field is from variant [PublicDatePoint].
	Month    int64  `json:"month"`
	TimeType string `json:"timeType"`
	// This field is from variant [PublicDatePoint].
	Year   int64  `json:"year"`
	ZoneID string `json:"zoneId"`
	// This field is from variant [PublicDatePoint].
	Hour int64 `json:"hour"`
	// This field is from variant [PublicDatePoint].
	Millisecond int64 `json:"millisecond"`
	// This field is from variant [PublicDatePoint].
	Minute int64 `json:"minute"`
	// This field is from variant [PublicDatePoint].
	Second         int64  `json:"second"`
	TimezoneSource string `json:"timezoneSource"`
	// This field is from variant [PublicIndexedTimePoint].
	IndexReference PublicIndexedTimePointIndexReferenceUnion `json:"indexReference"`
	// This field is from variant [PublicIndexedTimePoint].
	Offset PublicIndexOffset `json:"offset"`
	// This field is from variant [PublicPropertyReferencedTime].
	Property string `json:"property"`
	// This field is from variant [PublicPropertyReferencedTime].
	ReferenceType string `json:"referenceType"`
	JSON          struct {
		Day            respjson.Field
		Month          respjson.Field
		TimeType       respjson.Field
		Year           respjson.Field
		ZoneID         respjson.Field
		Hour           respjson.Field
		Millisecond    respjson.Field
		Minute         respjson.Field
		Second         respjson.Field
		TimezoneSource respjson.Field
		IndexReference respjson.Field
		Offset         respjson.Field
		Property       respjson.Field
		ReferenceType  respjson.Field
		raw            string
	} `json:"-"`
}

func (u PublicRangedTimeOperationUpperBoundTimePointUnion) AsDate() (v PublicDatePoint) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicRangedTimeOperationUpperBoundTimePointUnion) AsIndexed() (v PublicIndexedTimePoint) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicRangedTimeOperationUpperBoundTimePointUnion) AsPropertyReferenced() (v PublicPropertyReferencedTime) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u PublicRangedTimeOperationUpperBoundTimePointUnion) RawJSON() string { return u.JSON.raw }

func (r *PublicRangedTimeOperationUpperBoundTimePointUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties IncludeObjectsWithNoValueSet, LowerBoundTimePoint, OperationType,
// Operator, Type, UpperBoundTimePoint are required.
type PublicRangedTimeOperationParam struct {
	// Indicates whether objects with no value set for the property should be included
	// in the operation.
	IncludeObjectsWithNoValueSet bool `json:"includeObjectsWithNoValueSet" api:"required"`
	// Defines the lower bound time point for the operation.
	LowerBoundTimePoint PublicRangedTimeOperationLowerBoundTimePointUnionParam `json:"lowerBoundTimePoint,omitzero" api:"required"`
	// Specifies the type of operation (TIME_RANGED).
	//
	// Any of "TIME_RANGED".
	OperationType PublicRangedTimeOperationOperationType `json:"operationType,omitzero" api:"required"`
	// Defines the operation to be applied within the time range (IS_BETWEEN,
	// IS_NOT_BETWEEN).
	Operator string `json:"operator" api:"required"`
	// Specifies the type of operation (TIME_RANGED).
	//
	// Any of "TIME_RANGED".
	Type PublicRangedTimeOperationType `json:"type,omitzero" api:"required"`
	// Defines the upper bound time point for the operation.
	UpperBoundTimePoint PublicRangedTimeOperationUpperBoundTimePointUnionParam `json:"upperBoundTimePoint,omitzero" api:"required"`
	// Describes the behavior at the lower bound endpoint of the time range.
	LowerBoundEndpointBehavior param.Opt[string] `json:"lowerBoundEndpointBehavior,omitzero"`
	// Specifies the parser used for the property in the operation.
	PropertyParser param.Opt[string] `json:"propertyParser,omitzero"`
	// Describes the behavior at the upper bound endpoint of the time range.
	UpperBoundEndpointBehavior param.Opt[string] `json:"upperBoundEndpointBehavior,omitzero"`
	paramObj
}

func (r PublicRangedTimeOperationParam) MarshalJSON() (data []byte, err error) {
	type shadow PublicRangedTimeOperationParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PublicRangedTimeOperationParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type PublicRangedTimeOperationLowerBoundTimePointUnionParam struct {
	OfDate               *PublicDatePointParam              `json:",omitzero,inline"`
	OfIndexed            *PublicIndexedTimePointParam       `json:",omitzero,inline"`
	OfPropertyReferenced *PublicPropertyReferencedTimeParam `json:",omitzero,inline"`
	paramUnion
}

func (u PublicRangedTimeOperationLowerBoundTimePointUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfDate, u.OfIndexed, u.OfPropertyReferenced)
}
func (u *PublicRangedTimeOperationLowerBoundTimePointUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type PublicRangedTimeOperationUpperBoundTimePointUnionParam struct {
	OfDate               *PublicDatePointParam              `json:",omitzero,inline"`
	OfIndexed            *PublicIndexedTimePointParam       `json:",omitzero,inline"`
	OfPropertyReferenced *PublicPropertyReferencedTimeParam `json:",omitzero,inline"`
	paramUnion
}

func (u PublicRangedTimeOperationUpperBoundTimePointUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfDate, u.OfIndexed, u.OfPropertyReferenced)
}
func (u *PublicRangedTimeOperationUpperBoundTimePointUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

type PublicRelativeComparativeTimestampRefineBy struct {
	// Defines the comparison operation to be used in the refinement (BEFORE, AFTER).
	Comparison string           `json:"comparison" api:"required"`
	TimeOffset PublicTimeOffset `json:"timeOffset" api:"required"`
	// Specifies the type of refinement, (RELATIVE_COMPARATIVE).
	//
	// Any of "RELATIVE_COMPARATIVE".
	Type PublicRelativeComparativeTimestampRefineByType `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Comparison  respjson.Field
		TimeOffset  respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PublicRelativeComparativeTimestampRefineBy) RawJSON() string { return r.JSON.raw }
func (r *PublicRelativeComparativeTimestampRefineBy) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this PublicRelativeComparativeTimestampRefineBy to a
// PublicRelativeComparativeTimestampRefineByParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// PublicRelativeComparativeTimestampRefineByParam.Overrides()
func (r PublicRelativeComparativeTimestampRefineBy) ToParam() PublicRelativeComparativeTimestampRefineByParam {
	return param.Override[PublicRelativeComparativeTimestampRefineByParam](json.RawMessage(r.RawJSON()))
}

// Specifies the type of refinement, (RELATIVE_COMPARATIVE).
type PublicRelativeComparativeTimestampRefineByType string

const (
	PublicRelativeComparativeTimestampRefineByTypeRelativeComparative PublicRelativeComparativeTimestampRefineByType = "RELATIVE_COMPARATIVE"
)

// The properties Comparison, TimeOffset, Type are required.
type PublicRelativeComparativeTimestampRefineByParam struct {
	// Defines the comparison operation to be used in the refinement (BEFORE, AFTER).
	Comparison string                `json:"comparison" api:"required"`
	TimeOffset PublicTimeOffsetParam `json:"timeOffset,omitzero" api:"required"`
	// Specifies the type of refinement, (RELATIVE_COMPARATIVE).
	//
	// Any of "RELATIVE_COMPARATIVE".
	Type PublicRelativeComparativeTimestampRefineByType `json:"type,omitzero" api:"required"`
	paramObj
}

func (r PublicRelativeComparativeTimestampRefineByParam) MarshalJSON() (data []byte, err error) {
	type shadow PublicRelativeComparativeTimestampRefineByParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PublicRelativeComparativeTimestampRefineByParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PublicRelativeRangedTimestampRefineBy struct {
	LowerBoundOffset PublicTimeOffset `json:"lowerBoundOffset" api:"required"`
	// Specifies the type of range for the refinement criteria (BETWEEN, NOT_BETWEEN).
	RangeType string `json:"rangeType" api:"required"`
	// Indicates the type of refinement (RELATIVE_RANGED).
	//
	// Any of "RELATIVE_RANGED".
	Type             PublicRelativeRangedTimestampRefineByType `json:"type" api:"required"`
	UpperBoundOffset PublicTimeOffset                          `json:"upperBoundOffset" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		LowerBoundOffset respjson.Field
		RangeType        respjson.Field
		Type             respjson.Field
		UpperBoundOffset respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PublicRelativeRangedTimestampRefineBy) RawJSON() string { return r.JSON.raw }
func (r *PublicRelativeRangedTimestampRefineBy) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this PublicRelativeRangedTimestampRefineBy to a
// PublicRelativeRangedTimestampRefineByParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// PublicRelativeRangedTimestampRefineByParam.Overrides()
func (r PublicRelativeRangedTimestampRefineBy) ToParam() PublicRelativeRangedTimestampRefineByParam {
	return param.Override[PublicRelativeRangedTimestampRefineByParam](json.RawMessage(r.RawJSON()))
}

// Indicates the type of refinement (RELATIVE_RANGED).
type PublicRelativeRangedTimestampRefineByType string

const (
	PublicRelativeRangedTimestampRefineByTypeRelativeRanged PublicRelativeRangedTimestampRefineByType = "RELATIVE_RANGED"
)

// The properties LowerBoundOffset, RangeType, Type, UpperBoundOffset are required.
type PublicRelativeRangedTimestampRefineByParam struct {
	LowerBoundOffset PublicTimeOffsetParam `json:"lowerBoundOffset,omitzero" api:"required"`
	// Specifies the type of range for the refinement criteria (BETWEEN, NOT_BETWEEN).
	RangeType string `json:"rangeType" api:"required"`
	// Indicates the type of refinement (RELATIVE_RANGED).
	//
	// Any of "RELATIVE_RANGED".
	Type             PublicRelativeRangedTimestampRefineByType `json:"type,omitzero" api:"required"`
	UpperBoundOffset PublicTimeOffsetParam                     `json:"upperBoundOffset,omitzero" api:"required"`
	paramObj
}

func (r PublicRelativeRangedTimestampRefineByParam) MarshalJSON() (data []byte, err error) {
	type shadow PublicRelativeRangedTimestampRefineByParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PublicRelativeRangedTimestampRefineByParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PublicRestrictedFilterBranch struct {
	FilterBranches []PublicRestrictedFilterBranchFilterBranchUnion `json:"filterBranches" api:"required"`
	// The logical operator used to combine filters within the restricted filter
	// branch.
	FilterBranchOperator string `json:"filterBranchOperator" api:"required"`
	// Specifies the type of the filter branch (RESTRICTED).
	//
	// Any of "RESTRICTED".
	FilterBranchType PublicRestrictedFilterBranchFilterBranchType `json:"filterBranchType" api:"required"`
	Filters          []PublicRestrictedFilterBranchFilterUnion    `json:"filters" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FilterBranches       respjson.Field
		FilterBranchOperator respjson.Field
		FilterBranchType     respjson.Field
		Filters              respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PublicRestrictedFilterBranch) RawJSON() string { return r.JSON.raw }
func (r *PublicRestrictedFilterBranch) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this PublicRestrictedFilterBranch to a
// PublicRestrictedFilterBranchParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// PublicRestrictedFilterBranchParam.Overrides()
func (r PublicRestrictedFilterBranch) ToParam() PublicRestrictedFilterBranchParam {
	return param.Override[PublicRestrictedFilterBranchParam](json.RawMessage(r.RawJSON()))
}

// PublicRestrictedFilterBranchFilterBranchUnion contains all possible properties
// and values from [PublicOrFilterBranch], [PublicAndFilterBranch],
// [PublicNotAllFilterBranch], [PublicNotAnyFilterBranch],
// [PublicRestrictedFilterBranch], [PublicUnifiedEventsFilterBranch],
// [PublicPropertyAssociationFilterBranch], [PublicAssociationFilterBranch].
//
// Use the [PublicRestrictedFilterBranchFilterBranchUnion.AsAny] method to switch
// on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type PublicRestrictedFilterBranchFilterBranchUnion struct {
	// This field is a union of [[]PublicOrFilterBranchFilterBranchUnion],
	// [[]PublicAndFilterBranchFilterBranchUnion],
	// [[]PublicNotAllFilterBranchFilterBranchUnion],
	// [[]PublicNotAnyFilterBranchFilterBranchUnion],
	// [[]PublicRestrictedFilterBranchFilterBranchUnion],
	// [[]PublicUnifiedEventsFilterBranchFilterBranchUnion],
	// [[]PublicPropertyAssociationFilterBranchFilterBranchUnion],
	// [[]PublicAssociationFilterBranchFilterBranchUnion]
	FilterBranches       PublicRestrictedFilterBranchFilterBranchUnionFilterBranches `json:"filterBranches"`
	FilterBranchOperator string                                                      `json:"filterBranchOperator"`
	// Any of "OR", "AND", "NOT_ALL", "NOT_ANY", "RESTRICTED", "UNIFIED_EVENTS",
	// "PROPERTY_ASSOCIATION", "ASSOCIATION".
	FilterBranchType string `json:"filterBranchType"`
	// This field is a union of [[]PublicOrFilterBranchFilterUnion],
	// [[]PublicAndFilterBranchFilterUnion], [[]PublicNotAllFilterBranchFilterUnion],
	// [[]PublicNotAnyFilterBranchFilterUnion],
	// [[]PublicRestrictedFilterBranchFilterUnion],
	// [[]PublicUnifiedEventsFilterBranchFilterUnion],
	// [[]PublicPropertyAssociationFilterBranchFilterUnion],
	// [[]PublicAssociationFilterBranchFilterUnion]
	Filters PublicRestrictedFilterBranchFilterBranchUnionFilters `json:"filters"`
	// This field is from variant [PublicUnifiedEventsFilterBranch].
	EventTypeID string `json:"eventTypeId"`
	Operator    string `json:"operator"`
	// This field is from variant [PublicUnifiedEventsFilterBranch].
	CoalescingRefineBy PublicUnifiedEventsFilterBranchCoalescingRefineByUnion `json:"coalescingRefineBy"`
	// This field is from variant [PublicUnifiedEventsFilterBranch].
	PruningRefineBy PublicUnifiedEventsFilterBranchPruningRefineByUnion `json:"pruningRefineBy"`
	ObjectTypeID    string                                              `json:"objectTypeId"`
	// This field is from variant [PublicPropertyAssociationFilterBranch].
	PropertyWithObjectID string `json:"propertyWithObjectId"`
	// This field is from variant [PublicAssociationFilterBranch].
	AssociationCategory string `json:"associationCategory"`
	// This field is from variant [PublicAssociationFilterBranch].
	AssociationTypeID int64 `json:"associationTypeId"`
	JSON              struct {
		FilterBranches       respjson.Field
		FilterBranchOperator respjson.Field
		FilterBranchType     respjson.Field
		Filters              respjson.Field
		EventTypeID          respjson.Field
		Operator             respjson.Field
		CoalescingRefineBy   respjson.Field
		PruningRefineBy      respjson.Field
		ObjectTypeID         respjson.Field
		PropertyWithObjectID respjson.Field
		AssociationCategory  respjson.Field
		AssociationTypeID    respjson.Field
		raw                  string
	} `json:"-"`
}

// anyPublicRestrictedFilterBranchFilterBranch is implemented by each variant of
// [PublicRestrictedFilterBranchFilterBranchUnion] to add type safety for the
// return type of [PublicRestrictedFilterBranchFilterBranchUnion.AsAny]
type anyPublicRestrictedFilterBranchFilterBranch interface {
	implPublicRestrictedFilterBranchFilterBranchUnion()
}

func (PublicOrFilterBranch) implPublicRestrictedFilterBranchFilterBranchUnion()                  {}
func (PublicAndFilterBranch) implPublicRestrictedFilterBranchFilterBranchUnion()                 {}
func (PublicNotAllFilterBranch) implPublicRestrictedFilterBranchFilterBranchUnion()              {}
func (PublicNotAnyFilterBranch) implPublicRestrictedFilterBranchFilterBranchUnion()              {}
func (PublicRestrictedFilterBranch) implPublicRestrictedFilterBranchFilterBranchUnion()          {}
func (PublicUnifiedEventsFilterBranch) implPublicRestrictedFilterBranchFilterBranchUnion()       {}
func (PublicPropertyAssociationFilterBranch) implPublicRestrictedFilterBranchFilterBranchUnion() {}
func (PublicAssociationFilterBranch) implPublicRestrictedFilterBranchFilterBranchUnion()         {}

// Use the following switch statement to find the correct variant
//
//	switch variant := PublicRestrictedFilterBranchFilterBranchUnion.AsAny().(type) {
//	case crm.PublicOrFilterBranch:
//	case crm.PublicAndFilterBranch:
//	case crm.PublicNotAllFilterBranch:
//	case crm.PublicNotAnyFilterBranch:
//	case crm.PublicRestrictedFilterBranch:
//	case crm.PublicUnifiedEventsFilterBranch:
//	case crm.PublicPropertyAssociationFilterBranch:
//	case crm.PublicAssociationFilterBranch:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u PublicRestrictedFilterBranchFilterBranchUnion) AsAny() anyPublicRestrictedFilterBranchFilterBranch {
	switch u.FilterBranchType {
	case "OR":
		return u.AsOr()
	case "AND":
		return u.AsAnd()
	case "NOT_ALL":
		return u.AsNotAll()
	case "NOT_ANY":
		return u.AsNotAny()
	case "RESTRICTED":
		return u.AsRestricted()
	case "UNIFIED_EVENTS":
		return u.AsUnifiedEvents()
	case "PROPERTY_ASSOCIATION":
		return u.AsPropertyAssociation()
	case "ASSOCIATION":
		return u.AsAssociation()
	}
	return nil
}

func (u PublicRestrictedFilterBranchFilterBranchUnion) AsOr() (v PublicOrFilterBranch) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicRestrictedFilterBranchFilterBranchUnion) AsAnd() (v PublicAndFilterBranch) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicRestrictedFilterBranchFilterBranchUnion) AsNotAll() (v PublicNotAllFilterBranch) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicRestrictedFilterBranchFilterBranchUnion) AsNotAny() (v PublicNotAnyFilterBranch) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicRestrictedFilterBranchFilterBranchUnion) AsRestricted() (v PublicRestrictedFilterBranch) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicRestrictedFilterBranchFilterBranchUnion) AsUnifiedEvents() (v PublicUnifiedEventsFilterBranch) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicRestrictedFilterBranchFilterBranchUnion) AsPropertyAssociation() (v PublicPropertyAssociationFilterBranch) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicRestrictedFilterBranchFilterBranchUnion) AsAssociation() (v PublicAssociationFilterBranch) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u PublicRestrictedFilterBranchFilterBranchUnion) RawJSON() string { return u.JSON.raw }

func (r *PublicRestrictedFilterBranchFilterBranchUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PublicRestrictedFilterBranchFilterBranchUnionFilterBranches is an implicit
// subunion of [PublicRestrictedFilterBranchFilterBranchUnion].
// PublicRestrictedFilterBranchFilterBranchUnionFilterBranches provides convenient
// access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [PublicRestrictedFilterBranchFilterBranchUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfPublicOrFilterBranchFilterBranches
// OfPublicAndFilterBranchFilterBranches OfPublicNotAllFilterBranchFilterBranches
// OfPublicNotAnyFilterBranchFilterBranches
// OfPublicRestrictedFilterBranchFilterBranches
// OfPublicUnifiedEventsFilterBranchFilterBranches
// OfPublicPropertyAssociationFilterBranchFilterBranches
// OfPublicAssociationFilterBranchFilterBranches]
type PublicRestrictedFilterBranchFilterBranchUnionFilterBranches struct {
	// This field will be present if the value is a
	// [[]PublicOrFilterBranchFilterBranchUnion] instead of an object.
	OfPublicOrFilterBranchFilterBranches []PublicOrFilterBranchFilterBranchUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]PublicAndFilterBranchFilterBranchUnion] instead of an object.
	OfPublicAndFilterBranchFilterBranches []PublicAndFilterBranchFilterBranchUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]PublicNotAllFilterBranchFilterBranchUnion] instead of an object.
	OfPublicNotAllFilterBranchFilterBranches []PublicNotAllFilterBranchFilterBranchUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]PublicNotAnyFilterBranchFilterBranchUnion] instead of an object.
	OfPublicNotAnyFilterBranchFilterBranches []PublicNotAnyFilterBranchFilterBranchUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]PublicRestrictedFilterBranchFilterBranchUnion] instead of an object.
	OfPublicRestrictedFilterBranchFilterBranches []PublicRestrictedFilterBranchFilterBranchUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]PublicUnifiedEventsFilterBranchFilterBranchUnion] instead of an object.
	OfPublicUnifiedEventsFilterBranchFilterBranches []PublicUnifiedEventsFilterBranchFilterBranchUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]PublicPropertyAssociationFilterBranchFilterBranchUnion] instead of an object.
	OfPublicPropertyAssociationFilterBranchFilterBranches []PublicPropertyAssociationFilterBranchFilterBranchUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]PublicAssociationFilterBranchFilterBranchUnion] instead of an object.
	OfPublicAssociationFilterBranchFilterBranches []PublicAssociationFilterBranchFilterBranchUnion `json:",inline"`
	JSON                                          struct {
		OfPublicOrFilterBranchFilterBranches                  respjson.Field
		OfPublicAndFilterBranchFilterBranches                 respjson.Field
		OfPublicNotAllFilterBranchFilterBranches              respjson.Field
		OfPublicNotAnyFilterBranchFilterBranches              respjson.Field
		OfPublicRestrictedFilterBranchFilterBranches          respjson.Field
		OfPublicUnifiedEventsFilterBranchFilterBranches       respjson.Field
		OfPublicPropertyAssociationFilterBranchFilterBranches respjson.Field
		OfPublicAssociationFilterBranchFilterBranches         respjson.Field
		raw                                                   string
	} `json:"-"`
}

func (r *PublicRestrictedFilterBranchFilterBranchUnionFilterBranches) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PublicRestrictedFilterBranchFilterBranchUnionFilters is an implicit subunion of
// [PublicRestrictedFilterBranchFilterBranchUnion].
// PublicRestrictedFilterBranchFilterBranchUnionFilters provides convenient access
// to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [PublicRestrictedFilterBranchFilterBranchUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfPublicOrFilterBranchFilters OfPublicAndFilterBranchFilters
// OfPublicNotAllFilterBranchFilters OfPublicNotAnyFilterBranchFilters
// OfPublicRestrictedFilterBranchFilters OfPublicUnifiedEventsFilterBranchFilters
// OfPublicPropertyAssociationFilterBranchFilters
// OfPublicAssociationFilterBranchFilters]
type PublicRestrictedFilterBranchFilterBranchUnionFilters struct {
	// This field will be present if the value is a [[]PublicOrFilterBranchFilterUnion]
	// instead of an object.
	OfPublicOrFilterBranchFilters []PublicOrFilterBranchFilterUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]PublicAndFilterBranchFilterUnion] instead of an object.
	OfPublicAndFilterBranchFilters []PublicAndFilterBranchFilterUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]PublicNotAllFilterBranchFilterUnion] instead of an object.
	OfPublicNotAllFilterBranchFilters []PublicNotAllFilterBranchFilterUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]PublicNotAnyFilterBranchFilterUnion] instead of an object.
	OfPublicNotAnyFilterBranchFilters []PublicNotAnyFilterBranchFilterUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]PublicRestrictedFilterBranchFilterUnion] instead of an object.
	OfPublicRestrictedFilterBranchFilters []PublicRestrictedFilterBranchFilterUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]PublicUnifiedEventsFilterBranchFilterUnion] instead of an object.
	OfPublicUnifiedEventsFilterBranchFilters []PublicUnifiedEventsFilterBranchFilterUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]PublicPropertyAssociationFilterBranchFilterUnion] instead of an object.
	OfPublicPropertyAssociationFilterBranchFilters []PublicPropertyAssociationFilterBranchFilterUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]PublicAssociationFilterBranchFilterUnion] instead of an object.
	OfPublicAssociationFilterBranchFilters []PublicAssociationFilterBranchFilterUnion `json:",inline"`
	JSON                                   struct {
		OfPublicOrFilterBranchFilters                  respjson.Field
		OfPublicAndFilterBranchFilters                 respjson.Field
		OfPublicNotAllFilterBranchFilters              respjson.Field
		OfPublicNotAnyFilterBranchFilters              respjson.Field
		OfPublicRestrictedFilterBranchFilters          respjson.Field
		OfPublicUnifiedEventsFilterBranchFilters       respjson.Field
		OfPublicPropertyAssociationFilterBranchFilters respjson.Field
		OfPublicAssociationFilterBranchFilters         respjson.Field
		raw                                            string
	} `json:"-"`
}

func (r *PublicRestrictedFilterBranchFilterBranchUnionFilters) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Specifies the type of the filter branch (RESTRICTED).
type PublicRestrictedFilterBranchFilterBranchType string

const (
	PublicRestrictedFilterBranchFilterBranchTypeRestricted PublicRestrictedFilterBranchFilterBranchType = "RESTRICTED"
)

// PublicRestrictedFilterBranchFilterUnion contains all possible properties and
// values from [PublicPropertyFilter], [PublicAssociationInListFilter],
// [PublicPageViewAnalyticsFilter], [PublicCtaAnalyticsFilter],
// [PublicEventAnalyticsFilter], [PublicFormSubmissionFilter],
// [PublicFormSubmissionOnPageFilter], [PublicIntegrationEventFilter],
// [PublicEmailSubscriptionFilter], [PublicCommunicationSubscriptionFilter],
// [PublicCampaignInfluencedFilter], [PublicSurveyMonkeyFilter],
// [PublicSurveyMonkeyValueFilter], [PublicWebinarFilter],
// [PublicEmailEventFilter], [PublicPrivacyAnalyticsFilter],
// [PublicAdsSearchFilter], [PublicAdsTimeFilter], [PublicInListFilter],
// [PublicNumAssociationsFilter], [PublicUnifiedEventsFilter],
// [PublicPropertyAssociationInListFilter], [PublicConstantFilter].
//
// Use the [PublicRestrictedFilterBranchFilterUnion.AsAny] method to switch on the
// variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type PublicRestrictedFilterBranchFilterUnion struct {
	// Any of "PROPERTY", "ASSOCIATION", "PAGE_VIEW", "CTA", "EVENT",
	// "FORM_SUBMISSION", "FORM_SUBMISSION_ON_PAGE", "INTEGRATION_EVENT",
	// "EMAIL_SUBSCRIPTION", "COMMUNICATION_SUBSCRIPTION", "CAMPAIGN_INFLUENCED",
	// "SURVEY_MONKEY", "SURVEY_MONKEY_VALUE", "WEBINAR", "EMAIL_EVENT", "PRIVACY",
	// "ADS_SEARCH", "ADS_TIME", "IN_LIST", "NUM_ASSOCIATIONS", "UNIFIED_EVENTS",
	// "PROPERTY_ASSOCIATION", "CONSTANT".
	FilterType string `json:"filterType"`
	// This field is from variant [PublicPropertyFilter].
	Operation PublicPropertyFilterOperationUnion `json:"operation"`
	// This field is from variant [PublicPropertyFilter].
	Property            string `json:"property"`
	AssociationCategory string `json:"associationCategory"`
	AssociationTypeID   int64  `json:"associationTypeId"`
	// This field is a union of [PublicAssociationInListFilterCoalescingRefineByUnion],
	// [PublicPageViewAnalyticsFilterCoalescingRefineByUnion],
	// [PublicCtaAnalyticsFilterCoalescingRefineByUnion],
	// [PublicEventAnalyticsFilterCoalescingRefineByUnion],
	// [PublicFormSubmissionFilterCoalescingRefineByUnion],
	// [PublicFormSubmissionOnPageFilterCoalescingRefineByUnion],
	// [PublicNumAssociationsFilterCoalescingRefineByUnion],
	// [PublicUnifiedEventsFilterCoalescingRefineByUnion],
	// [PublicPropertyAssociationInListFilterCoalescingRefineByUnion]
	CoalescingRefineBy PublicRestrictedFilterBranchFilterUnionCoalescingRefineBy `json:"coalescingRefineBy"`
	ListID             string                                                    `json:"listId"`
	Operator           string                                                    `json:"operator"`
	// This field is from variant [PublicAssociationInListFilter].
	ToObjectType   string `json:"toObjectType"`
	ToObjectTypeID string `json:"toObjectTypeId"`
	// This field is from variant [PublicPageViewAnalyticsFilter].
	PageURL string `json:"pageUrl"`
	// This field is from variant [PublicPageViewAnalyticsFilter].
	EnableTracking bool `json:"enableTracking"`
	// This field is a union of [PublicPageViewAnalyticsFilterPruningRefineByUnion],
	// [PublicCtaAnalyticsFilterPruningRefineByUnion],
	// [PublicEventAnalyticsFilterPruningRefineByUnion],
	// [PublicFormSubmissionFilterPruningRefineByUnion],
	// [PublicFormSubmissionOnPageFilterPruningRefineByUnion],
	// [PublicEmailEventFilterPruningRefineByUnion],
	// [PublicAdsTimeFilterPruningRefineByUnion],
	// [PublicUnifiedEventsFilterPruningRefineByUnion]
	PruningRefineBy PublicRestrictedFilterBranchFilterUnionPruningRefineBy `json:"pruningRefineBy"`
	// This field is from variant [PublicCtaAnalyticsFilter].
	CtaName string `json:"ctaName"`
	// This field is from variant [PublicEventAnalyticsFilter].
	EventID string `json:"eventId"`
	FormID  string `json:"formId"`
	// This field is from variant [PublicFormSubmissionOnPageFilter].
	PageID string `json:"pageId"`
	// This field is a union of [int64], [string]
	EventTypeID PublicRestrictedFilterBranchFilterUnionEventTypeID `json:"eventTypeId"`
	FilterLines []PublicEventFilterMetadata                        `json:"filterLines"`
	// This field is from variant [PublicEmailSubscriptionFilter].
	AcceptedStatuses []string `json:"acceptedStatuses"`
	SubscriptionIDs  []string `json:"subscriptionIds"`
	SubscriptionType string   `json:"subscriptionType"`
	// This field is from variant [PublicCommunicationSubscriptionFilter].
	AcceptedOptStates []string `json:"acceptedOptStates"`
	// This field is from variant [PublicCommunicationSubscriptionFilter].
	Channel string `json:"channel"`
	// This field is from variant [PublicCommunicationSubscriptionFilter].
	BusinessUnitID string `json:"businessUnitId"`
	// This field is from variant [PublicCampaignInfluencedFilter].
	CampaignID string `json:"campaignId"`
	SurveyID   string `json:"surveyId"`
	// This field is from variant [PublicSurveyMonkeyValueFilter].
	SurveyQuestion string `json:"surveyQuestion"`
	// This field is from variant [PublicSurveyMonkeyValueFilter].
	ValueComparison PublicSurveyMonkeyValueFilterValueComparisonUnion `json:"valueComparison"`
	// This field is from variant [PublicSurveyMonkeyValueFilter].
	SurveyAnswerColID string `json:"surveyAnswerColId"`
	// This field is from variant [PublicSurveyMonkeyValueFilter].
	SurveyAnswerRowID string `json:"surveyAnswerRowId"`
	// This field is from variant [PublicWebinarFilter].
	WebinarID string `json:"webinarId"`
	// This field is from variant [PublicEmailEventFilter].
	AppID string `json:"appId"`
	// This field is from variant [PublicEmailEventFilter].
	EmailID string `json:"emailId"`
	// This field is from variant [PublicEmailEventFilter].
	Level string `json:"level"`
	// This field is from variant [PublicEmailEventFilter].
	ClickURL string `json:"clickUrl"`
	// This field is from variant [PublicPrivacyAnalyticsFilter].
	PrivacyName string `json:"privacyName"`
	// This field is from variant [PublicAdsSearchFilter].
	AdNetwork string `json:"adNetwork"`
	// This field is from variant [PublicAdsSearchFilter].
	EntityType string `json:"entityType"`
	// This field is from variant [PublicAdsSearchFilter].
	SearchTerms []string `json:"searchTerms"`
	// This field is from variant [PublicAdsSearchFilter].
	SearchTermType string `json:"searchTermType"`
	// This field is from variant [PublicInListFilter].
	Metadata PublicInListFilterMetadata `json:"metadata"`
	// This field is from variant [PublicPropertyAssociationInListFilter].
	PropertyWithObjectID string `json:"propertyWithObjectId"`
	// This field is from variant [PublicConstantFilter].
	ShouldAccept bool `json:"shouldAccept"`
	// This field is from variant [PublicConstantFilter].
	Source string `json:"source"`
	JSON   struct {
		FilterType           respjson.Field
		Operation            respjson.Field
		Property             respjson.Field
		AssociationCategory  respjson.Field
		AssociationTypeID    respjson.Field
		CoalescingRefineBy   respjson.Field
		ListID               respjson.Field
		Operator             respjson.Field
		ToObjectType         respjson.Field
		ToObjectTypeID       respjson.Field
		PageURL              respjson.Field
		EnableTracking       respjson.Field
		PruningRefineBy      respjson.Field
		CtaName              respjson.Field
		EventID              respjson.Field
		FormID               respjson.Field
		PageID               respjson.Field
		EventTypeID          respjson.Field
		FilterLines          respjson.Field
		AcceptedStatuses     respjson.Field
		SubscriptionIDs      respjson.Field
		SubscriptionType     respjson.Field
		AcceptedOptStates    respjson.Field
		Channel              respjson.Field
		BusinessUnitID       respjson.Field
		CampaignID           respjson.Field
		SurveyID             respjson.Field
		SurveyQuestion       respjson.Field
		ValueComparison      respjson.Field
		SurveyAnswerColID    respjson.Field
		SurveyAnswerRowID    respjson.Field
		WebinarID            respjson.Field
		AppID                respjson.Field
		EmailID              respjson.Field
		Level                respjson.Field
		ClickURL             respjson.Field
		PrivacyName          respjson.Field
		AdNetwork            respjson.Field
		EntityType           respjson.Field
		SearchTerms          respjson.Field
		SearchTermType       respjson.Field
		Metadata             respjson.Field
		PropertyWithObjectID respjson.Field
		ShouldAccept         respjson.Field
		Source               respjson.Field
		raw                  string
	} `json:"-"`
}

// anyPublicRestrictedFilterBranchFilter is implemented by each variant of
// [PublicRestrictedFilterBranchFilterUnion] to add type safety for the return type
// of [PublicRestrictedFilterBranchFilterUnion.AsAny]
type anyPublicRestrictedFilterBranchFilter interface {
	implPublicRestrictedFilterBranchFilterUnion()
}

func (PublicPropertyFilter) implPublicRestrictedFilterBranchFilterUnion()                  {}
func (PublicAssociationInListFilter) implPublicRestrictedFilterBranchFilterUnion()         {}
func (PublicPageViewAnalyticsFilter) implPublicRestrictedFilterBranchFilterUnion()         {}
func (PublicCtaAnalyticsFilter) implPublicRestrictedFilterBranchFilterUnion()              {}
func (PublicEventAnalyticsFilter) implPublicRestrictedFilterBranchFilterUnion()            {}
func (PublicFormSubmissionFilter) implPublicRestrictedFilterBranchFilterUnion()            {}
func (PublicFormSubmissionOnPageFilter) implPublicRestrictedFilterBranchFilterUnion()      {}
func (PublicIntegrationEventFilter) implPublicRestrictedFilterBranchFilterUnion()          {}
func (PublicEmailSubscriptionFilter) implPublicRestrictedFilterBranchFilterUnion()         {}
func (PublicCommunicationSubscriptionFilter) implPublicRestrictedFilterBranchFilterUnion() {}
func (PublicCampaignInfluencedFilter) implPublicRestrictedFilterBranchFilterUnion()        {}
func (PublicSurveyMonkeyFilter) implPublicRestrictedFilterBranchFilterUnion()              {}
func (PublicSurveyMonkeyValueFilter) implPublicRestrictedFilterBranchFilterUnion()         {}
func (PublicWebinarFilter) implPublicRestrictedFilterBranchFilterUnion()                   {}
func (PublicEmailEventFilter) implPublicRestrictedFilterBranchFilterUnion()                {}
func (PublicPrivacyAnalyticsFilter) implPublicRestrictedFilterBranchFilterUnion()          {}
func (PublicAdsSearchFilter) implPublicRestrictedFilterBranchFilterUnion()                 {}
func (PublicAdsTimeFilter) implPublicRestrictedFilterBranchFilterUnion()                   {}
func (PublicInListFilter) implPublicRestrictedFilterBranchFilterUnion()                    {}
func (PublicNumAssociationsFilter) implPublicRestrictedFilterBranchFilterUnion()           {}
func (PublicUnifiedEventsFilter) implPublicRestrictedFilterBranchFilterUnion()             {}
func (PublicPropertyAssociationInListFilter) implPublicRestrictedFilterBranchFilterUnion() {}
func (PublicConstantFilter) implPublicRestrictedFilterBranchFilterUnion()                  {}

// Use the following switch statement to find the correct variant
//
//	switch variant := PublicRestrictedFilterBranchFilterUnion.AsAny().(type) {
//	case crm.PublicPropertyFilter:
//	case crm.PublicAssociationInListFilter:
//	case crm.PublicPageViewAnalyticsFilter:
//	case crm.PublicCtaAnalyticsFilter:
//	case crm.PublicEventAnalyticsFilter:
//	case crm.PublicFormSubmissionFilter:
//	case crm.PublicFormSubmissionOnPageFilter:
//	case crm.PublicIntegrationEventFilter:
//	case crm.PublicEmailSubscriptionFilter:
//	case crm.PublicCommunicationSubscriptionFilter:
//	case crm.PublicCampaignInfluencedFilter:
//	case crm.PublicSurveyMonkeyFilter:
//	case crm.PublicSurveyMonkeyValueFilter:
//	case crm.PublicWebinarFilter:
//	case crm.PublicEmailEventFilter:
//	case crm.PublicPrivacyAnalyticsFilter:
//	case crm.PublicAdsSearchFilter:
//	case crm.PublicAdsTimeFilter:
//	case crm.PublicInListFilter:
//	case crm.PublicNumAssociationsFilter:
//	case crm.PublicUnifiedEventsFilter:
//	case crm.PublicPropertyAssociationInListFilter:
//	case crm.PublicConstantFilter:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u PublicRestrictedFilterBranchFilterUnion) AsAny() anyPublicRestrictedFilterBranchFilter {
	switch u.FilterType {
	case "PROPERTY":
		return u.AsProperty()
	case "ASSOCIATION":
		return u.AsAssociation()
	case "PAGE_VIEW":
		return u.AsPageView()
	case "CTA":
		return u.AsCta()
	case "EVENT":
		return u.AsEvent()
	case "FORM_SUBMISSION":
		return u.AsFormSubmission()
	case "FORM_SUBMISSION_ON_PAGE":
		return u.AsFormSubmissionOnPage()
	case "INTEGRATION_EVENT":
		return u.AsIntegrationEvent()
	case "EMAIL_SUBSCRIPTION":
		return u.AsEmailSubscription()
	case "COMMUNICATION_SUBSCRIPTION":
		return u.AsCommunicationSubscription()
	case "CAMPAIGN_INFLUENCED":
		return u.AsCampaignInfluenced()
	case "SURVEY_MONKEY":
		return u.AsSurveyMonkey()
	case "SURVEY_MONKEY_VALUE":
		return u.AsSurveyMonkeyValue()
	case "WEBINAR":
		return u.AsWebinar()
	case "EMAIL_EVENT":
		return u.AsEmailEvent()
	case "PRIVACY":
		return u.AsPrivacy()
	case "ADS_SEARCH":
		return u.AsAdsSearch()
	case "ADS_TIME":
		return u.AsAdsTime()
	case "IN_LIST":
		return u.AsInList()
	case "NUM_ASSOCIATIONS":
		return u.AsNumAssociations()
	case "UNIFIED_EVENTS":
		return u.AsUnifiedEvents()
	case "PROPERTY_ASSOCIATION":
		return u.AsPropertyAssociation()
	case "CONSTANT":
		return u.AsConstant()
	}
	return nil
}

func (u PublicRestrictedFilterBranchFilterUnion) AsProperty() (v PublicPropertyFilter) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicRestrictedFilterBranchFilterUnion) AsAssociation() (v PublicAssociationInListFilter) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicRestrictedFilterBranchFilterUnion) AsPageView() (v PublicPageViewAnalyticsFilter) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicRestrictedFilterBranchFilterUnion) AsCta() (v PublicCtaAnalyticsFilter) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicRestrictedFilterBranchFilterUnion) AsEvent() (v PublicEventAnalyticsFilter) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicRestrictedFilterBranchFilterUnion) AsFormSubmission() (v PublicFormSubmissionFilter) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicRestrictedFilterBranchFilterUnion) AsFormSubmissionOnPage() (v PublicFormSubmissionOnPageFilter) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicRestrictedFilterBranchFilterUnion) AsIntegrationEvent() (v PublicIntegrationEventFilter) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicRestrictedFilterBranchFilterUnion) AsEmailSubscription() (v PublicEmailSubscriptionFilter) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicRestrictedFilterBranchFilterUnion) AsCommunicationSubscription() (v PublicCommunicationSubscriptionFilter) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicRestrictedFilterBranchFilterUnion) AsCampaignInfluenced() (v PublicCampaignInfluencedFilter) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicRestrictedFilterBranchFilterUnion) AsSurveyMonkey() (v PublicSurveyMonkeyFilter) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicRestrictedFilterBranchFilterUnion) AsSurveyMonkeyValue() (v PublicSurveyMonkeyValueFilter) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicRestrictedFilterBranchFilterUnion) AsWebinar() (v PublicWebinarFilter) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicRestrictedFilterBranchFilterUnion) AsEmailEvent() (v PublicEmailEventFilter) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicRestrictedFilterBranchFilterUnion) AsPrivacy() (v PublicPrivacyAnalyticsFilter) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicRestrictedFilterBranchFilterUnion) AsAdsSearch() (v PublicAdsSearchFilter) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicRestrictedFilterBranchFilterUnion) AsAdsTime() (v PublicAdsTimeFilter) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicRestrictedFilterBranchFilterUnion) AsInList() (v PublicInListFilter) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicRestrictedFilterBranchFilterUnion) AsNumAssociations() (v PublicNumAssociationsFilter) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicRestrictedFilterBranchFilterUnion) AsUnifiedEvents() (v PublicUnifiedEventsFilter) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicRestrictedFilterBranchFilterUnion) AsPropertyAssociation() (v PublicPropertyAssociationInListFilter) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicRestrictedFilterBranchFilterUnion) AsConstant() (v PublicConstantFilter) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u PublicRestrictedFilterBranchFilterUnion) RawJSON() string { return u.JSON.raw }

func (r *PublicRestrictedFilterBranchFilterUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PublicRestrictedFilterBranchFilterUnionCoalescingRefineBy is an implicit
// subunion of [PublicRestrictedFilterBranchFilterUnion].
// PublicRestrictedFilterBranchFilterUnionCoalescingRefineBy provides convenient
// access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [PublicRestrictedFilterBranchFilterUnion].
type PublicRestrictedFilterBranchFilterUnionCoalescingRefineBy struct {
	Type string `json:"type"`
	// This field is from variant
	// [PublicAssociationInListFilterCoalescingRefineByUnion],
	// [PublicPageViewAnalyticsFilterCoalescingRefineByUnion],
	// [PublicCtaAnalyticsFilterCoalescingRefineByUnion],
	// [PublicEventAnalyticsFilterCoalescingRefineByUnion],
	// [PublicFormSubmissionFilterCoalescingRefineByUnion],
	// [PublicFormSubmissionOnPageFilterCoalescingRefineByUnion],
	// [PublicNumAssociationsFilterCoalescingRefineByUnion],
	// [PublicUnifiedEventsFilterCoalescingRefineByUnion],
	// [PublicPropertyAssociationInListFilterCoalescingRefineByUnion].
	MaxOccurrences int64 `json:"maxOccurrences"`
	// This field is from variant
	// [PublicAssociationInListFilterCoalescingRefineByUnion],
	// [PublicPageViewAnalyticsFilterCoalescingRefineByUnion],
	// [PublicCtaAnalyticsFilterCoalescingRefineByUnion],
	// [PublicEventAnalyticsFilterCoalescingRefineByUnion],
	// [PublicFormSubmissionFilterCoalescingRefineByUnion],
	// [PublicFormSubmissionOnPageFilterCoalescingRefineByUnion],
	// [PublicNumAssociationsFilterCoalescingRefineByUnion],
	// [PublicUnifiedEventsFilterCoalescingRefineByUnion],
	// [PublicPropertyAssociationInListFilterCoalescingRefineByUnion].
	MinOccurrences int64 `json:"minOccurrences"`
	// This field is from variant
	// [PublicAssociationInListFilterCoalescingRefineByUnion],
	// [PublicPageViewAnalyticsFilterCoalescingRefineByUnion],
	// [PublicCtaAnalyticsFilterCoalescingRefineByUnion],
	// [PublicEventAnalyticsFilterCoalescingRefineByUnion],
	// [PublicFormSubmissionFilterCoalescingRefineByUnion],
	// [PublicFormSubmissionOnPageFilterCoalescingRefineByUnion],
	// [PublicNumAssociationsFilterCoalescingRefineByUnion],
	// [PublicUnifiedEventsFilterCoalescingRefineByUnion],
	// [PublicPropertyAssociationInListFilterCoalescingRefineByUnion].
	SetType    string `json:"setType"`
	Comparison string `json:"comparison"`
	// This field is from variant
	// [PublicAssociationInListFilterCoalescingRefineByUnion],
	// [PublicPageViewAnalyticsFilterCoalescingRefineByUnion],
	// [PublicCtaAnalyticsFilterCoalescingRefineByUnion],
	// [PublicEventAnalyticsFilterCoalescingRefineByUnion],
	// [PublicFormSubmissionFilterCoalescingRefineByUnion],
	// [PublicFormSubmissionOnPageFilterCoalescingRefineByUnion],
	// [PublicNumAssociationsFilterCoalescingRefineByUnion],
	// [PublicUnifiedEventsFilterCoalescingRefineByUnion],
	// [PublicPropertyAssociationInListFilterCoalescingRefineByUnion].
	TimeOffset PublicTimeOffset `json:"timeOffset"`
	// This field is from variant
	// [PublicAssociationInListFilterCoalescingRefineByUnion],
	// [PublicPageViewAnalyticsFilterCoalescingRefineByUnion],
	// [PublicCtaAnalyticsFilterCoalescingRefineByUnion],
	// [PublicEventAnalyticsFilterCoalescingRefineByUnion],
	// [PublicFormSubmissionFilterCoalescingRefineByUnion],
	// [PublicFormSubmissionOnPageFilterCoalescingRefineByUnion],
	// [PublicNumAssociationsFilterCoalescingRefineByUnion],
	// [PublicUnifiedEventsFilterCoalescingRefineByUnion],
	// [PublicPropertyAssociationInListFilterCoalescingRefineByUnion].
	LowerBoundOffset PublicTimeOffset `json:"lowerBoundOffset"`
	RangeType        string           `json:"rangeType"`
	// This field is from variant
	// [PublicAssociationInListFilterCoalescingRefineByUnion],
	// [PublicPageViewAnalyticsFilterCoalescingRefineByUnion],
	// [PublicCtaAnalyticsFilterCoalescingRefineByUnion],
	// [PublicEventAnalyticsFilterCoalescingRefineByUnion],
	// [PublicFormSubmissionFilterCoalescingRefineByUnion],
	// [PublicFormSubmissionOnPageFilterCoalescingRefineByUnion],
	// [PublicNumAssociationsFilterCoalescingRefineByUnion],
	// [PublicUnifiedEventsFilterCoalescingRefineByUnion],
	// [PublicPropertyAssociationInListFilterCoalescingRefineByUnion].
	UpperBoundOffset PublicTimeOffset `json:"upperBoundOffset"`
	// This field is from variant
	// [PublicAssociationInListFilterCoalescingRefineByUnion],
	// [PublicPageViewAnalyticsFilterCoalescingRefineByUnion],
	// [PublicCtaAnalyticsFilterCoalescingRefineByUnion],
	// [PublicEventAnalyticsFilterCoalescingRefineByUnion],
	// [PublicFormSubmissionFilterCoalescingRefineByUnion],
	// [PublicFormSubmissionOnPageFilterCoalescingRefineByUnion],
	// [PublicNumAssociationsFilterCoalescingRefineByUnion],
	// [PublicUnifiedEventsFilterCoalescingRefineByUnion],
	// [PublicPropertyAssociationInListFilterCoalescingRefineByUnion].
	Timestamp int64 `json:"timestamp"`
	// This field is from variant
	// [PublicAssociationInListFilterCoalescingRefineByUnion],
	// [PublicPageViewAnalyticsFilterCoalescingRefineByUnion],
	// [PublicCtaAnalyticsFilterCoalescingRefineByUnion],
	// [PublicEventAnalyticsFilterCoalescingRefineByUnion],
	// [PublicFormSubmissionFilterCoalescingRefineByUnion],
	// [PublicFormSubmissionOnPageFilterCoalescingRefineByUnion],
	// [PublicNumAssociationsFilterCoalescingRefineByUnion],
	// [PublicUnifiedEventsFilterCoalescingRefineByUnion],
	// [PublicPropertyAssociationInListFilterCoalescingRefineByUnion].
	LowerTimestamp int64 `json:"lowerTimestamp"`
	// This field is from variant
	// [PublicAssociationInListFilterCoalescingRefineByUnion],
	// [PublicPageViewAnalyticsFilterCoalescingRefineByUnion],
	// [PublicCtaAnalyticsFilterCoalescingRefineByUnion],
	// [PublicEventAnalyticsFilterCoalescingRefineByUnion],
	// [PublicFormSubmissionFilterCoalescingRefineByUnion],
	// [PublicFormSubmissionOnPageFilterCoalescingRefineByUnion],
	// [PublicNumAssociationsFilterCoalescingRefineByUnion],
	// [PublicUnifiedEventsFilterCoalescingRefineByUnion],
	// [PublicPropertyAssociationInListFilterCoalescingRefineByUnion].
	UpperTimestamp               int64  `json:"upperTimestamp"`
	IncludeObjectsWithNoValueSet bool   `json:"includeObjectsWithNoValueSet"`
	OperationType                string `json:"operationType"`
	Operator                     string `json:"operator"`
	// This field is from variant
	// [PublicAssociationInListFilterCoalescingRefineByUnion],
	// [PublicPageViewAnalyticsFilterCoalescingRefineByUnion],
	// [PublicCtaAnalyticsFilterCoalescingRefineByUnion],
	// [PublicEventAnalyticsFilterCoalescingRefineByUnion],
	// [PublicFormSubmissionFilterCoalescingRefineByUnion],
	// [PublicFormSubmissionOnPageFilterCoalescingRefineByUnion],
	// [PublicNumAssociationsFilterCoalescingRefineByUnion],
	// [PublicUnifiedEventsFilterCoalescingRefineByUnion],
	// [PublicPropertyAssociationInListFilterCoalescingRefineByUnion].
	TimePoint PublicTimePointOperationTimePointUnion `json:"timePoint"`
	// This field is from variant
	// [PublicAssociationInListFilterCoalescingRefineByUnion],
	// [PublicPageViewAnalyticsFilterCoalescingRefineByUnion],
	// [PublicCtaAnalyticsFilterCoalescingRefineByUnion],
	// [PublicEventAnalyticsFilterCoalescingRefineByUnion],
	// [PublicFormSubmissionFilterCoalescingRefineByUnion],
	// [PublicFormSubmissionOnPageFilterCoalescingRefineByUnion],
	// [PublicNumAssociationsFilterCoalescingRefineByUnion],
	// [PublicUnifiedEventsFilterCoalescingRefineByUnion],
	// [PublicPropertyAssociationInListFilterCoalescingRefineByUnion].
	EndpointBehavior string `json:"endpointBehavior"`
	PropertyParser   string `json:"propertyParser"`
	// This field is from variant
	// [PublicAssociationInListFilterCoalescingRefineByUnion],
	// [PublicPageViewAnalyticsFilterCoalescingRefineByUnion],
	// [PublicCtaAnalyticsFilterCoalescingRefineByUnion],
	// [PublicEventAnalyticsFilterCoalescingRefineByUnion],
	// [PublicFormSubmissionFilterCoalescingRefineByUnion],
	// [PublicFormSubmissionOnPageFilterCoalescingRefineByUnion],
	// [PublicNumAssociationsFilterCoalescingRefineByUnion],
	// [PublicUnifiedEventsFilterCoalescingRefineByUnion],
	// [PublicPropertyAssociationInListFilterCoalescingRefineByUnion].
	LowerBoundTimePoint PublicRangedTimeOperationLowerBoundTimePointUnion `json:"lowerBoundTimePoint"`
	// This field is from variant
	// [PublicAssociationInListFilterCoalescingRefineByUnion],
	// [PublicPageViewAnalyticsFilterCoalescingRefineByUnion],
	// [PublicCtaAnalyticsFilterCoalescingRefineByUnion],
	// [PublicEventAnalyticsFilterCoalescingRefineByUnion],
	// [PublicFormSubmissionFilterCoalescingRefineByUnion],
	// [PublicFormSubmissionOnPageFilterCoalescingRefineByUnion],
	// [PublicNumAssociationsFilterCoalescingRefineByUnion],
	// [PublicUnifiedEventsFilterCoalescingRefineByUnion],
	// [PublicPropertyAssociationInListFilterCoalescingRefineByUnion].
	UpperBoundTimePoint PublicRangedTimeOperationUpperBoundTimePointUnion `json:"upperBoundTimePoint"`
	// This field is from variant
	// [PublicAssociationInListFilterCoalescingRefineByUnion],
	// [PublicPageViewAnalyticsFilterCoalescingRefineByUnion],
	// [PublicCtaAnalyticsFilterCoalescingRefineByUnion],
	// [PublicEventAnalyticsFilterCoalescingRefineByUnion],
	// [PublicFormSubmissionFilterCoalescingRefineByUnion],
	// [PublicFormSubmissionOnPageFilterCoalescingRefineByUnion],
	// [PublicNumAssociationsFilterCoalescingRefineByUnion],
	// [PublicUnifiedEventsFilterCoalescingRefineByUnion],
	// [PublicPropertyAssociationInListFilterCoalescingRefineByUnion].
	LowerBoundEndpointBehavior string `json:"lowerBoundEndpointBehavior"`
	// This field is from variant
	// [PublicAssociationInListFilterCoalescingRefineByUnion],
	// [PublicPageViewAnalyticsFilterCoalescingRefineByUnion],
	// [PublicCtaAnalyticsFilterCoalescingRefineByUnion],
	// [PublicEventAnalyticsFilterCoalescingRefineByUnion],
	// [PublicFormSubmissionFilterCoalescingRefineByUnion],
	// [PublicFormSubmissionOnPageFilterCoalescingRefineByUnion],
	// [PublicNumAssociationsFilterCoalescingRefineByUnion],
	// [PublicUnifiedEventsFilterCoalescingRefineByUnion],
	// [PublicPropertyAssociationInListFilterCoalescingRefineByUnion].
	UpperBoundEndpointBehavior string `json:"upperBoundEndpointBehavior"`
	JSON                       struct {
		Type                         respjson.Field
		MaxOccurrences               respjson.Field
		MinOccurrences               respjson.Field
		SetType                      respjson.Field
		Comparison                   respjson.Field
		TimeOffset                   respjson.Field
		LowerBoundOffset             respjson.Field
		RangeType                    respjson.Field
		UpperBoundOffset             respjson.Field
		Timestamp                    respjson.Field
		LowerTimestamp               respjson.Field
		UpperTimestamp               respjson.Field
		IncludeObjectsWithNoValueSet respjson.Field
		OperationType                respjson.Field
		Operator                     respjson.Field
		TimePoint                    respjson.Field
		EndpointBehavior             respjson.Field
		PropertyParser               respjson.Field
		LowerBoundTimePoint          respjson.Field
		UpperBoundTimePoint          respjson.Field
		LowerBoundEndpointBehavior   respjson.Field
		UpperBoundEndpointBehavior   respjson.Field
		raw                          string
	} `json:"-"`
}

func (r *PublicRestrictedFilterBranchFilterUnionCoalescingRefineBy) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PublicRestrictedFilterBranchFilterUnionPruningRefineBy is an implicit subunion
// of [PublicRestrictedFilterBranchFilterUnion].
// PublicRestrictedFilterBranchFilterUnionPruningRefineBy provides convenient
// access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [PublicRestrictedFilterBranchFilterUnion].
type PublicRestrictedFilterBranchFilterUnionPruningRefineBy struct {
	Type string `json:"type"`
	// This field is from variant [PublicPageViewAnalyticsFilterPruningRefineByUnion],
	// [PublicCtaAnalyticsFilterPruningRefineByUnion],
	// [PublicEventAnalyticsFilterPruningRefineByUnion],
	// [PublicFormSubmissionFilterPruningRefineByUnion],
	// [PublicFormSubmissionOnPageFilterPruningRefineByUnion],
	// [PublicEmailEventFilterPruningRefineByUnion],
	// [PublicAdsTimeFilterPruningRefineByUnion],
	// [PublicUnifiedEventsFilterPruningRefineByUnion].
	MaxOccurrences int64 `json:"maxOccurrences"`
	// This field is from variant [PublicPageViewAnalyticsFilterPruningRefineByUnion],
	// [PublicCtaAnalyticsFilterPruningRefineByUnion],
	// [PublicEventAnalyticsFilterPruningRefineByUnion],
	// [PublicFormSubmissionFilterPruningRefineByUnion],
	// [PublicFormSubmissionOnPageFilterPruningRefineByUnion],
	// [PublicEmailEventFilterPruningRefineByUnion],
	// [PublicAdsTimeFilterPruningRefineByUnion],
	// [PublicUnifiedEventsFilterPruningRefineByUnion].
	MinOccurrences int64 `json:"minOccurrences"`
	// This field is from variant [PublicPageViewAnalyticsFilterPruningRefineByUnion],
	// [PublicCtaAnalyticsFilterPruningRefineByUnion],
	// [PublicEventAnalyticsFilterPruningRefineByUnion],
	// [PublicFormSubmissionFilterPruningRefineByUnion],
	// [PublicFormSubmissionOnPageFilterPruningRefineByUnion],
	// [PublicEmailEventFilterPruningRefineByUnion],
	// [PublicAdsTimeFilterPruningRefineByUnion],
	// [PublicUnifiedEventsFilterPruningRefineByUnion].
	SetType    string `json:"setType"`
	Comparison string `json:"comparison"`
	// This field is from variant [PublicPageViewAnalyticsFilterPruningRefineByUnion],
	// [PublicCtaAnalyticsFilterPruningRefineByUnion],
	// [PublicEventAnalyticsFilterPruningRefineByUnion],
	// [PublicFormSubmissionFilterPruningRefineByUnion],
	// [PublicFormSubmissionOnPageFilterPruningRefineByUnion],
	// [PublicEmailEventFilterPruningRefineByUnion],
	// [PublicAdsTimeFilterPruningRefineByUnion],
	// [PublicUnifiedEventsFilterPruningRefineByUnion].
	TimeOffset PublicTimeOffset `json:"timeOffset"`
	// This field is from variant [PublicPageViewAnalyticsFilterPruningRefineByUnion],
	// [PublicCtaAnalyticsFilterPruningRefineByUnion],
	// [PublicEventAnalyticsFilterPruningRefineByUnion],
	// [PublicFormSubmissionFilterPruningRefineByUnion],
	// [PublicFormSubmissionOnPageFilterPruningRefineByUnion],
	// [PublicEmailEventFilterPruningRefineByUnion],
	// [PublicAdsTimeFilterPruningRefineByUnion],
	// [PublicUnifiedEventsFilterPruningRefineByUnion].
	LowerBoundOffset PublicTimeOffset `json:"lowerBoundOffset"`
	RangeType        string           `json:"rangeType"`
	// This field is from variant [PublicPageViewAnalyticsFilterPruningRefineByUnion],
	// [PublicCtaAnalyticsFilterPruningRefineByUnion],
	// [PublicEventAnalyticsFilterPruningRefineByUnion],
	// [PublicFormSubmissionFilterPruningRefineByUnion],
	// [PublicFormSubmissionOnPageFilterPruningRefineByUnion],
	// [PublicEmailEventFilterPruningRefineByUnion],
	// [PublicAdsTimeFilterPruningRefineByUnion],
	// [PublicUnifiedEventsFilterPruningRefineByUnion].
	UpperBoundOffset PublicTimeOffset `json:"upperBoundOffset"`
	// This field is from variant [PublicPageViewAnalyticsFilterPruningRefineByUnion],
	// [PublicCtaAnalyticsFilterPruningRefineByUnion],
	// [PublicEventAnalyticsFilterPruningRefineByUnion],
	// [PublicFormSubmissionFilterPruningRefineByUnion],
	// [PublicFormSubmissionOnPageFilterPruningRefineByUnion],
	// [PublicEmailEventFilterPruningRefineByUnion],
	// [PublicAdsTimeFilterPruningRefineByUnion],
	// [PublicUnifiedEventsFilterPruningRefineByUnion].
	Timestamp int64 `json:"timestamp"`
	// This field is from variant [PublicPageViewAnalyticsFilterPruningRefineByUnion],
	// [PublicCtaAnalyticsFilterPruningRefineByUnion],
	// [PublicEventAnalyticsFilterPruningRefineByUnion],
	// [PublicFormSubmissionFilterPruningRefineByUnion],
	// [PublicFormSubmissionOnPageFilterPruningRefineByUnion],
	// [PublicEmailEventFilterPruningRefineByUnion],
	// [PublicAdsTimeFilterPruningRefineByUnion],
	// [PublicUnifiedEventsFilterPruningRefineByUnion].
	LowerTimestamp int64 `json:"lowerTimestamp"`
	// This field is from variant [PublicPageViewAnalyticsFilterPruningRefineByUnion],
	// [PublicCtaAnalyticsFilterPruningRefineByUnion],
	// [PublicEventAnalyticsFilterPruningRefineByUnion],
	// [PublicFormSubmissionFilterPruningRefineByUnion],
	// [PublicFormSubmissionOnPageFilterPruningRefineByUnion],
	// [PublicEmailEventFilterPruningRefineByUnion],
	// [PublicAdsTimeFilterPruningRefineByUnion],
	// [PublicUnifiedEventsFilterPruningRefineByUnion].
	UpperTimestamp               int64  `json:"upperTimestamp"`
	IncludeObjectsWithNoValueSet bool   `json:"includeObjectsWithNoValueSet"`
	OperationType                string `json:"operationType"`
	Operator                     string `json:"operator"`
	// This field is from variant [PublicPageViewAnalyticsFilterPruningRefineByUnion],
	// [PublicCtaAnalyticsFilterPruningRefineByUnion],
	// [PublicEventAnalyticsFilterPruningRefineByUnion],
	// [PublicFormSubmissionFilterPruningRefineByUnion],
	// [PublicFormSubmissionOnPageFilterPruningRefineByUnion],
	// [PublicEmailEventFilterPruningRefineByUnion],
	// [PublicAdsTimeFilterPruningRefineByUnion],
	// [PublicUnifiedEventsFilterPruningRefineByUnion].
	TimePoint PublicTimePointOperationTimePointUnion `json:"timePoint"`
	// This field is from variant [PublicPageViewAnalyticsFilterPruningRefineByUnion],
	// [PublicCtaAnalyticsFilterPruningRefineByUnion],
	// [PublicEventAnalyticsFilterPruningRefineByUnion],
	// [PublicFormSubmissionFilterPruningRefineByUnion],
	// [PublicFormSubmissionOnPageFilterPruningRefineByUnion],
	// [PublicEmailEventFilterPruningRefineByUnion],
	// [PublicAdsTimeFilterPruningRefineByUnion],
	// [PublicUnifiedEventsFilterPruningRefineByUnion].
	EndpointBehavior string `json:"endpointBehavior"`
	PropertyParser   string `json:"propertyParser"`
	// This field is from variant [PublicPageViewAnalyticsFilterPruningRefineByUnion],
	// [PublicCtaAnalyticsFilterPruningRefineByUnion],
	// [PublicEventAnalyticsFilterPruningRefineByUnion],
	// [PublicFormSubmissionFilterPruningRefineByUnion],
	// [PublicFormSubmissionOnPageFilterPruningRefineByUnion],
	// [PublicEmailEventFilterPruningRefineByUnion],
	// [PublicAdsTimeFilterPruningRefineByUnion],
	// [PublicUnifiedEventsFilterPruningRefineByUnion].
	LowerBoundTimePoint PublicRangedTimeOperationLowerBoundTimePointUnion `json:"lowerBoundTimePoint"`
	// This field is from variant [PublicPageViewAnalyticsFilterPruningRefineByUnion],
	// [PublicCtaAnalyticsFilterPruningRefineByUnion],
	// [PublicEventAnalyticsFilterPruningRefineByUnion],
	// [PublicFormSubmissionFilterPruningRefineByUnion],
	// [PublicFormSubmissionOnPageFilterPruningRefineByUnion],
	// [PublicEmailEventFilterPruningRefineByUnion],
	// [PublicAdsTimeFilterPruningRefineByUnion],
	// [PublicUnifiedEventsFilterPruningRefineByUnion].
	UpperBoundTimePoint PublicRangedTimeOperationUpperBoundTimePointUnion `json:"upperBoundTimePoint"`
	// This field is from variant [PublicPageViewAnalyticsFilterPruningRefineByUnion],
	// [PublicCtaAnalyticsFilterPruningRefineByUnion],
	// [PublicEventAnalyticsFilterPruningRefineByUnion],
	// [PublicFormSubmissionFilterPruningRefineByUnion],
	// [PublicFormSubmissionOnPageFilterPruningRefineByUnion],
	// [PublicEmailEventFilterPruningRefineByUnion],
	// [PublicAdsTimeFilterPruningRefineByUnion],
	// [PublicUnifiedEventsFilterPruningRefineByUnion].
	LowerBoundEndpointBehavior string `json:"lowerBoundEndpointBehavior"`
	// This field is from variant [PublicPageViewAnalyticsFilterPruningRefineByUnion],
	// [PublicCtaAnalyticsFilterPruningRefineByUnion],
	// [PublicEventAnalyticsFilterPruningRefineByUnion],
	// [PublicFormSubmissionFilterPruningRefineByUnion],
	// [PublicFormSubmissionOnPageFilterPruningRefineByUnion],
	// [PublicEmailEventFilterPruningRefineByUnion],
	// [PublicAdsTimeFilterPruningRefineByUnion],
	// [PublicUnifiedEventsFilterPruningRefineByUnion].
	UpperBoundEndpointBehavior string `json:"upperBoundEndpointBehavior"`
	JSON                       struct {
		Type                         respjson.Field
		MaxOccurrences               respjson.Field
		MinOccurrences               respjson.Field
		SetType                      respjson.Field
		Comparison                   respjson.Field
		TimeOffset                   respjson.Field
		LowerBoundOffset             respjson.Field
		RangeType                    respjson.Field
		UpperBoundOffset             respjson.Field
		Timestamp                    respjson.Field
		LowerTimestamp               respjson.Field
		UpperTimestamp               respjson.Field
		IncludeObjectsWithNoValueSet respjson.Field
		OperationType                respjson.Field
		Operator                     respjson.Field
		TimePoint                    respjson.Field
		EndpointBehavior             respjson.Field
		PropertyParser               respjson.Field
		LowerBoundTimePoint          respjson.Field
		UpperBoundTimePoint          respjson.Field
		LowerBoundEndpointBehavior   respjson.Field
		UpperBoundEndpointBehavior   respjson.Field
		raw                          string
	} `json:"-"`
}

func (r *PublicRestrictedFilterBranchFilterUnionPruningRefineBy) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PublicRestrictedFilterBranchFilterUnionEventTypeID is an implicit subunion of
// [PublicRestrictedFilterBranchFilterUnion].
// PublicRestrictedFilterBranchFilterUnionEventTypeID provides convenient access to
// the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [PublicRestrictedFilterBranchFilterUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfInt OfString]
type PublicRestrictedFilterBranchFilterUnionEventTypeID struct {
	// This field will be present if the value is a [int64] instead of an object.
	OfInt int64 `json:",inline"`
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	JSON     struct {
		OfInt    respjson.Field
		OfString respjson.Field
		raw      string
	} `json:"-"`
}

func (r *PublicRestrictedFilterBranchFilterUnionEventTypeID) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties FilterBranches, FilterBranchOperator, FilterBranchType, Filters
// are required.
type PublicRestrictedFilterBranchParam struct {
	FilterBranches []PublicRestrictedFilterBranchFilterBranchUnionParam `json:"filterBranches,omitzero" api:"required"`
	// The logical operator used to combine filters within the restricted filter
	// branch.
	FilterBranchOperator string `json:"filterBranchOperator" api:"required"`
	// Specifies the type of the filter branch (RESTRICTED).
	//
	// Any of "RESTRICTED".
	FilterBranchType PublicRestrictedFilterBranchFilterBranchType   `json:"filterBranchType,omitzero" api:"required"`
	Filters          []PublicRestrictedFilterBranchFilterUnionParam `json:"filters,omitzero" api:"required"`
	paramObj
}

func (r PublicRestrictedFilterBranchParam) MarshalJSON() (data []byte, err error) {
	type shadow PublicRestrictedFilterBranchParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PublicRestrictedFilterBranchParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type PublicRestrictedFilterBranchFilterBranchUnionParam struct {
	OfOr                  *PublicOrFilterBranchParam                  `json:",omitzero,inline"`
	OfAnd                 *PublicAndFilterBranchParam                 `json:",omitzero,inline"`
	OfNotAll              *PublicNotAllFilterBranchParam              `json:",omitzero,inline"`
	OfNotAny              *PublicNotAnyFilterBranchParam              `json:",omitzero,inline"`
	OfRestricted          *PublicRestrictedFilterBranchParam          `json:",omitzero,inline"`
	OfUnifiedEvents       *PublicUnifiedEventsFilterBranchParam       `json:",omitzero,inline"`
	OfPropertyAssociation *PublicPropertyAssociationFilterBranchParam `json:",omitzero,inline"`
	OfAssociation         *PublicAssociationFilterBranchParam         `json:",omitzero,inline"`
	paramUnion
}

func (u PublicRestrictedFilterBranchFilterBranchUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfOr,
		u.OfAnd,
		u.OfNotAll,
		u.OfNotAny,
		u.OfRestricted,
		u.OfUnifiedEvents,
		u.OfPropertyAssociation,
		u.OfAssociation)
}
func (u *PublicRestrictedFilterBranchFilterBranchUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func init() {
	apijson.RegisterUnion[PublicRestrictedFilterBranchFilterBranchUnionParam](
		"filterBranchType",
		apijson.Discriminator[PublicOrFilterBranchParam]("OR"),
		apijson.Discriminator[PublicAndFilterBranchParam]("AND"),
		apijson.Discriminator[PublicNotAllFilterBranchParam]("NOT_ALL"),
		apijson.Discriminator[PublicNotAnyFilterBranchParam]("NOT_ANY"),
		apijson.Discriminator[PublicRestrictedFilterBranchParam]("RESTRICTED"),
		apijson.Discriminator[PublicUnifiedEventsFilterBranchParam]("UNIFIED_EVENTS"),
		apijson.Discriminator[PublicPropertyAssociationFilterBranchParam]("PROPERTY_ASSOCIATION"),
		apijson.Discriminator[PublicAssociationFilterBranchParam]("ASSOCIATION"),
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type PublicRestrictedFilterBranchFilterUnionParam struct {
	OfProperty                  *PublicPropertyFilterParam                  `json:",omitzero,inline"`
	OfAssociation               *PublicAssociationInListFilterParam         `json:",omitzero,inline"`
	OfPageView                  *PublicPageViewAnalyticsFilterParam         `json:",omitzero,inline"`
	OfCta                       *PublicCtaAnalyticsFilterParam              `json:",omitzero,inline"`
	OfEvent                     *PublicEventAnalyticsFilterParam            `json:",omitzero,inline"`
	OfFormSubmission            *PublicFormSubmissionFilterParam            `json:",omitzero,inline"`
	OfFormSubmissionOnPage      *PublicFormSubmissionOnPageFilterParam      `json:",omitzero,inline"`
	OfIntegrationEvent          *PublicIntegrationEventFilterParam          `json:",omitzero,inline"`
	OfEmailSubscription         *PublicEmailSubscriptionFilterParam         `json:",omitzero,inline"`
	OfCommunicationSubscription *PublicCommunicationSubscriptionFilterParam `json:",omitzero,inline"`
	OfCampaignInfluenced        *PublicCampaignInfluencedFilterParam        `json:",omitzero,inline"`
	OfSurveyMonkey              *PublicSurveyMonkeyFilterParam              `json:",omitzero,inline"`
	OfSurveyMonkeyValue         *PublicSurveyMonkeyValueFilterParam         `json:",omitzero,inline"`
	OfWebinar                   *PublicWebinarFilterParam                   `json:",omitzero,inline"`
	OfEmailEvent                *PublicEmailEventFilterParam                `json:",omitzero,inline"`
	OfPrivacy                   *PublicPrivacyAnalyticsFilterParam          `json:",omitzero,inline"`
	OfAdsSearch                 *PublicAdsSearchFilterParam                 `json:",omitzero,inline"`
	OfAdsTime                   *PublicAdsTimeFilterParam                   `json:",omitzero,inline"`
	OfInList                    *PublicInListFilterParam                    `json:",omitzero,inline"`
	OfNumAssociations           *PublicNumAssociationsFilterParam           `json:",omitzero,inline"`
	OfUnifiedEvents             *PublicUnifiedEventsFilterParam             `json:",omitzero,inline"`
	OfPropertyAssociation       *PublicPropertyAssociationInListFilterParam `json:",omitzero,inline"`
	OfConstant                  *PublicConstantFilterParam                  `json:",omitzero,inline"`
	paramUnion
}

func (u PublicRestrictedFilterBranchFilterUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfProperty,
		u.OfAssociation,
		u.OfPageView,
		u.OfCta,
		u.OfEvent,
		u.OfFormSubmission,
		u.OfFormSubmissionOnPage,
		u.OfIntegrationEvent,
		u.OfEmailSubscription,
		u.OfCommunicationSubscription,
		u.OfCampaignInfluenced,
		u.OfSurveyMonkey,
		u.OfSurveyMonkeyValue,
		u.OfWebinar,
		u.OfEmailEvent,
		u.OfPrivacy,
		u.OfAdsSearch,
		u.OfAdsTime,
		u.OfInList,
		u.OfNumAssociations,
		u.OfUnifiedEvents,
		u.OfPropertyAssociation,
		u.OfConstant)
}
func (u *PublicRestrictedFilterBranchFilterUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func init() {
	apijson.RegisterUnion[PublicRestrictedFilterBranchFilterUnionParam](
		"filterType",
		apijson.Discriminator[PublicPropertyFilterParam]("PROPERTY"),
		apijson.Discriminator[PublicAssociationInListFilterParam]("ASSOCIATION"),
		apijson.Discriminator[PublicPageViewAnalyticsFilterParam]("PAGE_VIEW"),
		apijson.Discriminator[PublicCtaAnalyticsFilterParam]("CTA"),
		apijson.Discriminator[PublicEventAnalyticsFilterParam]("EVENT"),
		apijson.Discriminator[PublicFormSubmissionFilterParam]("FORM_SUBMISSION"),
		apijson.Discriminator[PublicFormSubmissionOnPageFilterParam]("FORM_SUBMISSION_ON_PAGE"),
		apijson.Discriminator[PublicIntegrationEventFilterParam]("INTEGRATION_EVENT"),
		apijson.Discriminator[PublicEmailSubscriptionFilterParam]("EMAIL_SUBSCRIPTION"),
		apijson.Discriminator[PublicCommunicationSubscriptionFilterParam]("COMMUNICATION_SUBSCRIPTION"),
		apijson.Discriminator[PublicCampaignInfluencedFilterParam]("CAMPAIGN_INFLUENCED"),
		apijson.Discriminator[PublicSurveyMonkeyFilterParam]("SURVEY_MONKEY"),
		apijson.Discriminator[PublicSurveyMonkeyValueFilterParam]("SURVEY_MONKEY_VALUE"),
		apijson.Discriminator[PublicWebinarFilterParam]("WEBINAR"),
		apijson.Discriminator[PublicEmailEventFilterParam]("EMAIL_EVENT"),
		apijson.Discriminator[PublicPrivacyAnalyticsFilterParam]("PRIVACY"),
		apijson.Discriminator[PublicAdsSearchFilterParam]("ADS_SEARCH"),
		apijson.Discriminator[PublicAdsTimeFilterParam]("ADS_TIME"),
		apijson.Discriminator[PublicInListFilterParam]("IN_LIST"),
		apijson.Discriminator[PublicNumAssociationsFilterParam]("NUM_ASSOCIATIONS"),
		apijson.Discriminator[PublicUnifiedEventsFilterParam]("UNIFIED_EVENTS"),
		apijson.Discriminator[PublicPropertyAssociationInListFilterParam]("PROPERTY_ASSOCIATION"),
		apijson.Discriminator[PublicConstantFilterParam]("CONSTANT"),
	)
}

type PublicRollingDateRangePropertyOperation struct {
	// Indicates whether objects with no value set for the property should be included
	// in the operation.
	IncludeObjectsWithNoValueSet bool `json:"includeObjectsWithNoValueSet" api:"required"`
	// The number of days to be considered in the rolling date range operation.
	NumberOfDays int64 `json:"numberOfDays" api:"required"`
	// Specifies the type of operation (ROLLING_DATE_RANGE).
	//
	// Any of "ROLLING_DATE_RANGE".
	OperationType PublicRollingDateRangePropertyOperationOperationType `json:"operationType" api:"required"`
	// Defines the operation to be applied within the rolling date range property
	// operation (IS_LESS_THAN_X_DAYS_AGO, IS_MORE_THAN_X_DAYS_AGO,
	// IS_LESS_THAN_X_DAYS_FROM_NOW, IS_MORE_THAN_X_DAYS_FROM_NOW).
	Operator string `json:"operator" api:"required"`
	// Specifies whether the operation requires conversion to a different time zone.
	RequiresTimeZoneConversion bool `json:"requiresTimeZoneConversion" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		IncludeObjectsWithNoValueSet respjson.Field
		NumberOfDays                 respjson.Field
		OperationType                respjson.Field
		Operator                     respjson.Field
		RequiresTimeZoneConversion   respjson.Field
		ExtraFields                  map[string]respjson.Field
		raw                          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PublicRollingDateRangePropertyOperation) RawJSON() string { return r.JSON.raw }
func (r *PublicRollingDateRangePropertyOperation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this PublicRollingDateRangePropertyOperation to a
// PublicRollingDateRangePropertyOperationParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// PublicRollingDateRangePropertyOperationParam.Overrides()
func (r PublicRollingDateRangePropertyOperation) ToParam() PublicRollingDateRangePropertyOperationParam {
	return param.Override[PublicRollingDateRangePropertyOperationParam](json.RawMessage(r.RawJSON()))
}

// Specifies the type of operation (ROLLING_DATE_RANGE).
type PublicRollingDateRangePropertyOperationOperationType string

const (
	PublicRollingDateRangePropertyOperationOperationTypeRollingDateRange PublicRollingDateRangePropertyOperationOperationType = "ROLLING_DATE_RANGE"
)

// The properties IncludeObjectsWithNoValueSet, NumberOfDays, OperationType,
// Operator, RequiresTimeZoneConversion are required.
type PublicRollingDateRangePropertyOperationParam struct {
	// Indicates whether objects with no value set for the property should be included
	// in the operation.
	IncludeObjectsWithNoValueSet bool `json:"includeObjectsWithNoValueSet" api:"required"`
	// The number of days to be considered in the rolling date range operation.
	NumberOfDays int64 `json:"numberOfDays" api:"required"`
	// Specifies the type of operation (ROLLING_DATE_RANGE).
	//
	// Any of "ROLLING_DATE_RANGE".
	OperationType PublicRollingDateRangePropertyOperationOperationType `json:"operationType,omitzero" api:"required"`
	// Defines the operation to be applied within the rolling date range property
	// operation (IS_LESS_THAN_X_DAYS_AGO, IS_MORE_THAN_X_DAYS_AGO,
	// IS_LESS_THAN_X_DAYS_FROM_NOW, IS_MORE_THAN_X_DAYS_FROM_NOW).
	Operator string `json:"operator" api:"required"`
	// Specifies whether the operation requires conversion to a different time zone.
	RequiresTimeZoneConversion bool `json:"requiresTimeZoneConversion" api:"required"`
	paramObj
}

func (r PublicRollingDateRangePropertyOperationParam) MarshalJSON() (data []byte, err error) {
	type shadow PublicRollingDateRangePropertyOperationParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PublicRollingDateRangePropertyOperationParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PublicRollingPropertyUpdatedOperation struct {
	// Indicates whether objects with no value set for the property should be included
	// in the operation.
	IncludeObjectsWithNoValueSet bool `json:"includeObjectsWithNoValueSet" api:"required"`
	// The number of days to be considered in the rolling property updated operation.
	NumberOfDays int64 `json:"numberOfDays" api:"required"`
	// Specifies the type of operation (ROLLING_PROPERTY_UPDATED).
	//
	// Any of "ROLLING_PROPERTY_UPDATED".
	OperationType PublicRollingPropertyUpdatedOperationOperationType `json:"operationType" api:"required"`
	// Defines the operation to be applied within the rolling property updated
	// operation (UPDATED_IN_LAST_X_DAYS, NOT_UPDATED_IN_LAST_X_DAYS).
	Operator string `json:"operator" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		IncludeObjectsWithNoValueSet respjson.Field
		NumberOfDays                 respjson.Field
		OperationType                respjson.Field
		Operator                     respjson.Field
		ExtraFields                  map[string]respjson.Field
		raw                          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PublicRollingPropertyUpdatedOperation) RawJSON() string { return r.JSON.raw }
func (r *PublicRollingPropertyUpdatedOperation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this PublicRollingPropertyUpdatedOperation to a
// PublicRollingPropertyUpdatedOperationParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// PublicRollingPropertyUpdatedOperationParam.Overrides()
func (r PublicRollingPropertyUpdatedOperation) ToParam() PublicRollingPropertyUpdatedOperationParam {
	return param.Override[PublicRollingPropertyUpdatedOperationParam](json.RawMessage(r.RawJSON()))
}

// Specifies the type of operation (ROLLING_PROPERTY_UPDATED).
type PublicRollingPropertyUpdatedOperationOperationType string

const (
	PublicRollingPropertyUpdatedOperationOperationTypeRollingPropertyUpdated PublicRollingPropertyUpdatedOperationOperationType = "ROLLING_PROPERTY_UPDATED"
)

// The properties IncludeObjectsWithNoValueSet, NumberOfDays, OperationType,
// Operator are required.
type PublicRollingPropertyUpdatedOperationParam struct {
	// Indicates whether objects with no value set for the property should be included
	// in the operation.
	IncludeObjectsWithNoValueSet bool `json:"includeObjectsWithNoValueSet" api:"required"`
	// The number of days to be considered in the rolling property updated operation.
	NumberOfDays int64 `json:"numberOfDays" api:"required"`
	// Specifies the type of operation (ROLLING_PROPERTY_UPDATED).
	//
	// Any of "ROLLING_PROPERTY_UPDATED".
	OperationType PublicRollingPropertyUpdatedOperationOperationType `json:"operationType,omitzero" api:"required"`
	// Defines the operation to be applied within the rolling property updated
	// operation (UPDATED_IN_LAST_X_DAYS, NOT_UPDATED_IN_LAST_X_DAYS).
	Operator string `json:"operator" api:"required"`
	paramObj
}

func (r PublicRollingPropertyUpdatedOperationParam) MarshalJSON() (data []byte, err error) {
	type shadow PublicRollingPropertyUpdatedOperationParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PublicRollingPropertyUpdatedOperationParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PublicSetOccurrencesRefineBy struct {
	// Indicates the specific set type used in the refinement (ALL, ALL_INCLUDE_EMPTY,
	// ANY, NONE, NONE_EXCLUDE_EMPTY, ANY_INCLUDE_EMPTY).
	SetType string `json:"setType" api:"required"`
	// Specifies the type of refinement (SET_OCCURRENCES).
	//
	// Any of "SET_OCCURRENCES".
	Type PublicSetOccurrencesRefineByType `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		SetType     respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PublicSetOccurrencesRefineBy) RawJSON() string { return r.JSON.raw }
func (r *PublicSetOccurrencesRefineBy) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this PublicSetOccurrencesRefineBy to a
// PublicSetOccurrencesRefineByParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// PublicSetOccurrencesRefineByParam.Overrides()
func (r PublicSetOccurrencesRefineBy) ToParam() PublicSetOccurrencesRefineByParam {
	return param.Override[PublicSetOccurrencesRefineByParam](json.RawMessage(r.RawJSON()))
}

// Specifies the type of refinement (SET_OCCURRENCES).
type PublicSetOccurrencesRefineByType string

const (
	PublicSetOccurrencesRefineByTypeSetOccurrences PublicSetOccurrencesRefineByType = "SET_OCCURRENCES"
)

// The properties SetType, Type are required.
type PublicSetOccurrencesRefineByParam struct {
	// Indicates the specific set type used in the refinement (ALL, ALL_INCLUDE_EMPTY,
	// ANY, NONE, NONE_EXCLUDE_EMPTY, ANY_INCLUDE_EMPTY).
	SetType string `json:"setType" api:"required"`
	// Specifies the type of refinement (SET_OCCURRENCES).
	//
	// Any of "SET_OCCURRENCES".
	Type PublicSetOccurrencesRefineByType `json:"type,omitzero" api:"required"`
	paramObj
}

func (r PublicSetOccurrencesRefineByParam) MarshalJSON() (data []byte, err error) {
	type shadow PublicSetOccurrencesRefineByParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PublicSetOccurrencesRefineByParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PublicStringPropertyOperation struct {
	// Indicates whether objects with no value set for the property should be included
	// in the operation.
	IncludeObjectsWithNoValueSet bool `json:"includeObjectsWithNoValueSet" api:"required"`
	// Specifies the type of operation (STRING).
	//
	// Any of "STRING".
	OperationType PublicStringPropertyOperationOperationType `json:"operationType" api:"required"`
	// Defines the operation to be applied in the string property operation
	// ()IS_EQUAL_TO, IS_NOT_EQUAL_TO, CONTAINS, DOES_NOT_CONTAIN, STARTS_WITH,
	// ENDS_WITH, HAS_EVER_BEEN_EQUAL_TO, HAS_NEVER_BEEN_EQUAL_TO, HAS_EVER_CONTAINED,
	// HAS_NEVER_CONTAINED).
	Operator string `json:"operator" api:"required"`
	// The string value to be used in the operation.
	Value string `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		IncludeObjectsWithNoValueSet respjson.Field
		OperationType                respjson.Field
		Operator                     respjson.Field
		Value                        respjson.Field
		ExtraFields                  map[string]respjson.Field
		raw                          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PublicStringPropertyOperation) RawJSON() string { return r.JSON.raw }
func (r *PublicStringPropertyOperation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this PublicStringPropertyOperation to a
// PublicStringPropertyOperationParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// PublicStringPropertyOperationParam.Overrides()
func (r PublicStringPropertyOperation) ToParam() PublicStringPropertyOperationParam {
	return param.Override[PublicStringPropertyOperationParam](json.RawMessage(r.RawJSON()))
}

// Specifies the type of operation (STRING).
type PublicStringPropertyOperationOperationType string

const (
	PublicStringPropertyOperationOperationTypeString PublicStringPropertyOperationOperationType = "STRING"
)

// The properties IncludeObjectsWithNoValueSet, OperationType, Operator, Value are
// required.
type PublicStringPropertyOperationParam struct {
	// Indicates whether objects with no value set for the property should be included
	// in the operation.
	IncludeObjectsWithNoValueSet bool `json:"includeObjectsWithNoValueSet" api:"required"`
	// Specifies the type of operation (STRING).
	//
	// Any of "STRING".
	OperationType PublicStringPropertyOperationOperationType `json:"operationType,omitzero" api:"required"`
	// Defines the operation to be applied in the string property operation
	// ()IS_EQUAL_TO, IS_NOT_EQUAL_TO, CONTAINS, DOES_NOT_CONTAIN, STARTS_WITH,
	// ENDS_WITH, HAS_EVER_BEEN_EQUAL_TO, HAS_NEVER_BEEN_EQUAL_TO, HAS_EVER_CONTAINED,
	// HAS_NEVER_CONTAINED).
	Operator string `json:"operator" api:"required"`
	// The string value to be used in the operation.
	Value string `json:"value" api:"required"`
	paramObj
}

func (r PublicStringPropertyOperationParam) MarshalJSON() (data []byte, err error) {
	type shadow PublicStringPropertyOperationParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PublicStringPropertyOperationParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PublicSurveyMonkeyFilter struct {
	// Indicates the type of filter being applied (SURVEY_MONKEY).
	//
	// Any of "SURVEY_MONKEY".
	FilterType PublicSurveyMonkeyFilterFilterType `json:"filterType" api:"required"`
	// Specifies the operation to be performed by the filter (HAS_RESPONDED_TO_SURVEY,
	// HAS_NOT_RESPONDED_TO_SURVEY).
	Operator string `json:"operator" api:"required"`
	// The ID of the survey associated with the filter.
	SurveyID string `json:"surveyId" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FilterType  respjson.Field
		Operator    respjson.Field
		SurveyID    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PublicSurveyMonkeyFilter) RawJSON() string { return r.JSON.raw }
func (r *PublicSurveyMonkeyFilter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this PublicSurveyMonkeyFilter to a
// PublicSurveyMonkeyFilterParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// PublicSurveyMonkeyFilterParam.Overrides()
func (r PublicSurveyMonkeyFilter) ToParam() PublicSurveyMonkeyFilterParam {
	return param.Override[PublicSurveyMonkeyFilterParam](json.RawMessage(r.RawJSON()))
}

// Indicates the type of filter being applied (SURVEY_MONKEY).
type PublicSurveyMonkeyFilterFilterType string

const (
	PublicSurveyMonkeyFilterFilterTypeSurveyMonkey PublicSurveyMonkeyFilterFilterType = "SURVEY_MONKEY"
)

// The properties FilterType, Operator, SurveyID are required.
type PublicSurveyMonkeyFilterParam struct {
	// Indicates the type of filter being applied (SURVEY_MONKEY).
	//
	// Any of "SURVEY_MONKEY".
	FilterType PublicSurveyMonkeyFilterFilterType `json:"filterType,omitzero" api:"required"`
	// Specifies the operation to be performed by the filter (HAS_RESPONDED_TO_SURVEY,
	// HAS_NOT_RESPONDED_TO_SURVEY).
	Operator string `json:"operator" api:"required"`
	// The ID of the survey associated with the filter.
	SurveyID string `json:"surveyId" api:"required"`
	paramObj
}

func (r PublicSurveyMonkeyFilterParam) MarshalJSON() (data []byte, err error) {
	type shadow PublicSurveyMonkeyFilterParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PublicSurveyMonkeyFilterParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PublicSurveyMonkeyValueFilter struct {
	// Defines the type of filter (SURVEY_MONKEY_VALUE).
	//
	// Any of "SURVEY_MONKEY_VALUE".
	FilterType PublicSurveyMonkeyValueFilterFilterType `json:"filterType" api:"required"`
	// Defines the operation to be applied within the filter
	// (HAS_ANSWERED_SURVEY_QUESTION_WITH_VALUE).
	Operator string `json:"operator" api:"required"`
	// The ID of the survey used in the filter.
	SurveyID string `json:"surveyId" api:"required"`
	// The question from the survey used in the filter.
	SurveyQuestion string `json:"surveyQuestion" api:"required"`
	// Specifies the operation used to compare the survey answer value.
	ValueComparison PublicSurveyMonkeyValueFilterValueComparisonUnion `json:"valueComparison" api:"required"`
	// The column ID of the survey answer used in the filter.
	SurveyAnswerColID string `json:"surveyAnswerColId"`
	// The row ID of the survey answer used in the filter.
	SurveyAnswerRowID string `json:"surveyAnswerRowId"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FilterType        respjson.Field
		Operator          respjson.Field
		SurveyID          respjson.Field
		SurveyQuestion    respjson.Field
		ValueComparison   respjson.Field
		SurveyAnswerColID respjson.Field
		SurveyAnswerRowID respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PublicSurveyMonkeyValueFilter) RawJSON() string { return r.JSON.raw }
func (r *PublicSurveyMonkeyValueFilter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this PublicSurveyMonkeyValueFilter to a
// PublicSurveyMonkeyValueFilterParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// PublicSurveyMonkeyValueFilterParam.Overrides()
func (r PublicSurveyMonkeyValueFilter) ToParam() PublicSurveyMonkeyValueFilterParam {
	return param.Override[PublicSurveyMonkeyValueFilterParam](json.RawMessage(r.RawJSON()))
}

// Defines the type of filter (SURVEY_MONKEY_VALUE).
type PublicSurveyMonkeyValueFilterFilterType string

const (
	PublicSurveyMonkeyValueFilterFilterTypeSurveyMonkeyValue PublicSurveyMonkeyValueFilterFilterType = "SURVEY_MONKEY_VALUE"
)

// PublicSurveyMonkeyValueFilterValueComparisonUnion contains all possible
// properties and values from [PublicBoolPropertyOperation],
// [PublicNumberPropertyOperation], [PublicStringPropertyOperation],
// [PublicDateTimePropertyOperation], [PublicRangedDatePropertyOperation],
// [PublicComparativePropertyUpdatedOperation],
// [PublicComparativeDatePropertyOperation],
// [PublicRollingDateRangePropertyOperation],
// [PublicRollingPropertyUpdatedOperation], [PublicEnumerationPropertyOperation],
// [PublicAllPropertyTypesOperation], [PublicRangedNumberPropertyOperation],
// [PublicMultiStringPropertyOperation], [PublicDatePropertyOperation],
// [PublicCalendarDatePropertyOperation], [PublicTimePointOperation],
// [PublicRangedTimeOperation].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type PublicSurveyMonkeyValueFilterValueComparisonUnion struct {
	IncludeObjectsWithNoValueSet bool   `json:"includeObjectsWithNoValueSet"`
	OperationType                string `json:"operationType"`
	Operator                     string `json:"operator"`
	// This field is a union of [bool], [float64], [string]
	Value                      PublicSurveyMonkeyValueFilterValueComparisonUnionValue `json:"value"`
	RequiresTimeZoneConversion bool                                                   `json:"requiresTimeZoneConversion"`
	// This field is from variant [PublicDateTimePropertyOperation].
	Timestamp              int64    `json:"timestamp"`
	LowerBound             int64    `json:"lowerBound"`
	UpperBound             int64    `json:"upperBound"`
	ComparisonPropertyName string   `json:"comparisonPropertyName"`
	DefaultComparisonValue string   `json:"defaultComparisonValue"`
	NumberOfDays           int64    `json:"numberOfDays"`
	Values                 []string `json:"values"`
	// This field is from variant [PublicDatePropertyOperation].
	Day int64 `json:"day"`
	// This field is from variant [PublicDatePropertyOperation].
	Month string `json:"month"`
	// This field is from variant [PublicDatePropertyOperation].
	Year int64 `json:"year"`
	// This field is from variant [PublicCalendarDatePropertyOperation].
	TimeUnit string `json:"timeUnit"`
	// This field is from variant [PublicCalendarDatePropertyOperation].
	FiscalYearStart PublicCalendarDatePropertyOperationFiscalYearStart `json:"fiscalYearStart"`
	// This field is from variant [PublicCalendarDatePropertyOperation].
	TimeUnitCount int64 `json:"timeUnitCount"`
	// This field is from variant [PublicCalendarDatePropertyOperation].
	UseFiscalYear bool `json:"useFiscalYear"`
	// This field is from variant [PublicTimePointOperation].
	TimePoint PublicTimePointOperationTimePointUnion `json:"timePoint"`
	Type      string                                 `json:"type"`
	// This field is from variant [PublicTimePointOperation].
	EndpointBehavior string `json:"endpointBehavior"`
	PropertyParser   string `json:"propertyParser"`
	// This field is from variant [PublicRangedTimeOperation].
	LowerBoundTimePoint PublicRangedTimeOperationLowerBoundTimePointUnion `json:"lowerBoundTimePoint"`
	// This field is from variant [PublicRangedTimeOperation].
	UpperBoundTimePoint PublicRangedTimeOperationUpperBoundTimePointUnion `json:"upperBoundTimePoint"`
	// This field is from variant [PublicRangedTimeOperation].
	LowerBoundEndpointBehavior string `json:"lowerBoundEndpointBehavior"`
	// This field is from variant [PublicRangedTimeOperation].
	UpperBoundEndpointBehavior string `json:"upperBoundEndpointBehavior"`
	JSON                       struct {
		IncludeObjectsWithNoValueSet respjson.Field
		OperationType                respjson.Field
		Operator                     respjson.Field
		Value                        respjson.Field
		RequiresTimeZoneConversion   respjson.Field
		Timestamp                    respjson.Field
		LowerBound                   respjson.Field
		UpperBound                   respjson.Field
		ComparisonPropertyName       respjson.Field
		DefaultComparisonValue       respjson.Field
		NumberOfDays                 respjson.Field
		Values                       respjson.Field
		Day                          respjson.Field
		Month                        respjson.Field
		Year                         respjson.Field
		TimeUnit                     respjson.Field
		FiscalYearStart              respjson.Field
		TimeUnitCount                respjson.Field
		UseFiscalYear                respjson.Field
		TimePoint                    respjson.Field
		Type                         respjson.Field
		EndpointBehavior             respjson.Field
		PropertyParser               respjson.Field
		LowerBoundTimePoint          respjson.Field
		UpperBoundTimePoint          respjson.Field
		LowerBoundEndpointBehavior   respjson.Field
		UpperBoundEndpointBehavior   respjson.Field
		raw                          string
	} `json:"-"`
}

func (u PublicSurveyMonkeyValueFilterValueComparisonUnion) AsBool() (v PublicBoolPropertyOperation) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicSurveyMonkeyValueFilterValueComparisonUnion) AsNumber() (v PublicNumberPropertyOperation) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicSurveyMonkeyValueFilterValueComparisonUnion) AsString() (v PublicStringPropertyOperation) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicSurveyMonkeyValueFilterValueComparisonUnion) AsDatetime() (v PublicDateTimePropertyOperation) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicSurveyMonkeyValueFilterValueComparisonUnion) AsRangedDate() (v PublicRangedDatePropertyOperation) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicSurveyMonkeyValueFilterValueComparisonUnion) AsComparativePropertyUpdated() (v PublicComparativePropertyUpdatedOperation) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicSurveyMonkeyValueFilterValueComparisonUnion) AsComparativeDate() (v PublicComparativeDatePropertyOperation) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicSurveyMonkeyValueFilterValueComparisonUnion) AsRollingDateRange() (v PublicRollingDateRangePropertyOperation) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicSurveyMonkeyValueFilterValueComparisonUnion) AsRollingPropertyUpdated() (v PublicRollingPropertyUpdatedOperation) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicSurveyMonkeyValueFilterValueComparisonUnion) AsEnumeration() (v PublicEnumerationPropertyOperation) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicSurveyMonkeyValueFilterValueComparisonUnion) AsAllProperty() (v PublicAllPropertyTypesOperation) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicSurveyMonkeyValueFilterValueComparisonUnion) AsNumberRanged() (v PublicRangedNumberPropertyOperation) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicSurveyMonkeyValueFilterValueComparisonUnion) AsMultistring() (v PublicMultiStringPropertyOperation) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicSurveyMonkeyValueFilterValueComparisonUnion) AsDate() (v PublicDatePropertyOperation) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicSurveyMonkeyValueFilterValueComparisonUnion) AsCalendarDate() (v PublicCalendarDatePropertyOperation) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicSurveyMonkeyValueFilterValueComparisonUnion) AsTimePoint() (v PublicTimePointOperation) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicSurveyMonkeyValueFilterValueComparisonUnion) AsTimeRanged() (v PublicRangedTimeOperation) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u PublicSurveyMonkeyValueFilterValueComparisonUnion) RawJSON() string { return u.JSON.raw }

func (r *PublicSurveyMonkeyValueFilterValueComparisonUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PublicSurveyMonkeyValueFilterValueComparisonUnionValue is an implicit subunion
// of [PublicSurveyMonkeyValueFilterValueComparisonUnion].
// PublicSurveyMonkeyValueFilterValueComparisonUnionValue provides convenient
// access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [PublicSurveyMonkeyValueFilterValueComparisonUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfBool OfFloat OfString]
type PublicSurveyMonkeyValueFilterValueComparisonUnionValue struct {
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	JSON     struct {
		OfBool   respjson.Field
		OfFloat  respjson.Field
		OfString respjson.Field
		raw      string
	} `json:"-"`
}

func (r *PublicSurveyMonkeyValueFilterValueComparisonUnionValue) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties FilterType, Operator, SurveyID, SurveyQuestion, ValueComparison
// are required.
type PublicSurveyMonkeyValueFilterParam struct {
	// Defines the type of filter (SURVEY_MONKEY_VALUE).
	//
	// Any of "SURVEY_MONKEY_VALUE".
	FilterType PublicSurveyMonkeyValueFilterFilterType `json:"filterType,omitzero" api:"required"`
	// Defines the operation to be applied within the filter
	// (HAS_ANSWERED_SURVEY_QUESTION_WITH_VALUE).
	Operator string `json:"operator" api:"required"`
	// The ID of the survey used in the filter.
	SurveyID string `json:"surveyId" api:"required"`
	// The question from the survey used in the filter.
	SurveyQuestion string `json:"surveyQuestion" api:"required"`
	// Specifies the operation used to compare the survey answer value.
	ValueComparison PublicSurveyMonkeyValueFilterValueComparisonUnionParam `json:"valueComparison,omitzero" api:"required"`
	// The column ID of the survey answer used in the filter.
	SurveyAnswerColID param.Opt[string] `json:"surveyAnswerColId,omitzero"`
	// The row ID of the survey answer used in the filter.
	SurveyAnswerRowID param.Opt[string] `json:"surveyAnswerRowId,omitzero"`
	paramObj
}

func (r PublicSurveyMonkeyValueFilterParam) MarshalJSON() (data []byte, err error) {
	type shadow PublicSurveyMonkeyValueFilterParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PublicSurveyMonkeyValueFilterParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type PublicSurveyMonkeyValueFilterValueComparisonUnionParam struct {
	OfBool                       *PublicBoolPropertyOperationParam               `json:",omitzero,inline"`
	OfNumber                     *PublicNumberPropertyOperationParam             `json:",omitzero,inline"`
	OfString                     *PublicStringPropertyOperationParam             `json:",omitzero,inline"`
	OfDatetime                   *PublicDateTimePropertyOperationParam           `json:",omitzero,inline"`
	OfRangedDate                 *PublicRangedDatePropertyOperationParam         `json:",omitzero,inline"`
	OfComparativePropertyUpdated *PublicComparativePropertyUpdatedOperationParam `json:",omitzero,inline"`
	OfComparativeDate            *PublicComparativeDatePropertyOperationParam    `json:",omitzero,inline"`
	OfRollingDateRange           *PublicRollingDateRangePropertyOperationParam   `json:",omitzero,inline"`
	OfRollingPropertyUpdated     *PublicRollingPropertyUpdatedOperationParam     `json:",omitzero,inline"`
	OfEnumeration                *PublicEnumerationPropertyOperationParam        `json:",omitzero,inline"`
	OfAllProperty                *PublicAllPropertyTypesOperationParam           `json:",omitzero,inline"`
	OfNumberRanged               *PublicRangedNumberPropertyOperationParam       `json:",omitzero,inline"`
	OfMultistring                *PublicMultiStringPropertyOperationParam        `json:",omitzero,inline"`
	OfDate                       *PublicDatePropertyOperationParam               `json:",omitzero,inline"`
	OfCalendarDate               *PublicCalendarDatePropertyOperationParam       `json:",omitzero,inline"`
	OfTimePoint                  *PublicTimePointOperationParam                  `json:",omitzero,inline"`
	OfTimeRanged                 *PublicRangedTimeOperationParam                 `json:",omitzero,inline"`
	paramUnion
}

func (u PublicSurveyMonkeyValueFilterValueComparisonUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfBool,
		u.OfNumber,
		u.OfString,
		u.OfDatetime,
		u.OfRangedDate,
		u.OfComparativePropertyUpdated,
		u.OfComparativeDate,
		u.OfRollingDateRange,
		u.OfRollingPropertyUpdated,
		u.OfEnumeration,
		u.OfAllProperty,
		u.OfNumberRanged,
		u.OfMultistring,
		u.OfDate,
		u.OfCalendarDate,
		u.OfTimePoint,
		u.OfTimeRanged)
}
func (u *PublicSurveyMonkeyValueFilterValueComparisonUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

type PublicTimeOffset struct {
	// The numerical value representing the quantity of the time offset.
	Amount int64 `json:"amount" api:"required"`
	// Indicates the direction of the time offset, such as forward or backward.
	OffsetDirection string `json:"offsetDirection" api:"required"`
	// Specifies the unit of time for the offset, such as days, hours, or minutes.
	TimeUnit string `json:"timeUnit" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Amount          respjson.Field
		OffsetDirection respjson.Field
		TimeUnit        respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PublicTimeOffset) RawJSON() string { return r.JSON.raw }
func (r *PublicTimeOffset) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this PublicTimeOffset to a PublicTimeOffsetParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// PublicTimeOffsetParam.Overrides()
func (r PublicTimeOffset) ToParam() PublicTimeOffsetParam {
	return param.Override[PublicTimeOffsetParam](json.RawMessage(r.RawJSON()))
}

// The properties Amount, OffsetDirection, TimeUnit are required.
type PublicTimeOffsetParam struct {
	// The numerical value representing the quantity of the time offset.
	Amount int64 `json:"amount" api:"required"`
	// Indicates the direction of the time offset, such as forward or backward.
	OffsetDirection string `json:"offsetDirection" api:"required"`
	// Specifies the unit of time for the offset, such as days, hours, or minutes.
	TimeUnit string `json:"timeUnit" api:"required"`
	paramObj
}

func (r PublicTimeOffsetParam) MarshalJSON() (data []byte, err error) {
	type shadow PublicTimeOffsetParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PublicTimeOffsetParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PublicTimePointOperation struct {
	// Indicates whether objects with no value set for the property should be included
	// in the operation.
	IncludeObjectsWithNoValueSet bool `json:"includeObjectsWithNoValueSet" api:"required"`
	// Specifies the type of operation (TIME_POINT).
	//
	// Any of "TIME_POINT".
	OperationType PublicTimePointOperationOperationType `json:"operationType" api:"required"`
	// Specifies the operation to be applied within the time point operation
	// (IS_BEFORE, IS_AFTER).
	Operator string `json:"operator" api:"required"`
	// Defines the specific point in time for the operation, which can be a date,
	// indexed time, or property-referenced time.
	TimePoint PublicTimePointOperationTimePointUnion `json:"timePoint" api:"required"`
	// Defines the type of operation being performed.
	Type string `json:"type" api:"required"`
	// Describes the behavior at the endpoint of the time point operation.
	EndpointBehavior string `json:"endpointBehavior"`
	// Specifies the parser used for interpreting the property in the operation.
	PropertyParser string `json:"propertyParser"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		IncludeObjectsWithNoValueSet respjson.Field
		OperationType                respjson.Field
		Operator                     respjson.Field
		TimePoint                    respjson.Field
		Type                         respjson.Field
		EndpointBehavior             respjson.Field
		PropertyParser               respjson.Field
		ExtraFields                  map[string]respjson.Field
		raw                          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PublicTimePointOperation) RawJSON() string { return r.JSON.raw }
func (r *PublicTimePointOperation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this PublicTimePointOperation to a
// PublicTimePointOperationParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// PublicTimePointOperationParam.Overrides()
func (r PublicTimePointOperation) ToParam() PublicTimePointOperationParam {
	return param.Override[PublicTimePointOperationParam](json.RawMessage(r.RawJSON()))
}

// Specifies the type of operation (TIME_POINT).
type PublicTimePointOperationOperationType string

const (
	PublicTimePointOperationOperationTypeTimePoint PublicTimePointOperationOperationType = "TIME_POINT"
)

// PublicTimePointOperationTimePointUnion contains all possible properties and
// values from [PublicDatePoint], [PublicIndexedTimePoint],
// [PublicPropertyReferencedTime].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type PublicTimePointOperationTimePointUnion struct {
	// This field is from variant [PublicDatePoint].
	Day int64 `json:"day"`
	// This field is from variant [PublicDatePoint].
	Month    int64  `json:"month"`
	TimeType string `json:"timeType"`
	// This field is from variant [PublicDatePoint].
	Year   int64  `json:"year"`
	ZoneID string `json:"zoneId"`
	// This field is from variant [PublicDatePoint].
	Hour int64 `json:"hour"`
	// This field is from variant [PublicDatePoint].
	Millisecond int64 `json:"millisecond"`
	// This field is from variant [PublicDatePoint].
	Minute int64 `json:"minute"`
	// This field is from variant [PublicDatePoint].
	Second         int64  `json:"second"`
	TimezoneSource string `json:"timezoneSource"`
	// This field is from variant [PublicIndexedTimePoint].
	IndexReference PublicIndexedTimePointIndexReferenceUnion `json:"indexReference"`
	// This field is from variant [PublicIndexedTimePoint].
	Offset PublicIndexOffset `json:"offset"`
	// This field is from variant [PublicPropertyReferencedTime].
	Property string `json:"property"`
	// This field is from variant [PublicPropertyReferencedTime].
	ReferenceType string `json:"referenceType"`
	JSON          struct {
		Day            respjson.Field
		Month          respjson.Field
		TimeType       respjson.Field
		Year           respjson.Field
		ZoneID         respjson.Field
		Hour           respjson.Field
		Millisecond    respjson.Field
		Minute         respjson.Field
		Second         respjson.Field
		TimezoneSource respjson.Field
		IndexReference respjson.Field
		Offset         respjson.Field
		Property       respjson.Field
		ReferenceType  respjson.Field
		raw            string
	} `json:"-"`
}

func (u PublicTimePointOperationTimePointUnion) AsDate() (v PublicDatePoint) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicTimePointOperationTimePointUnion) AsIndexed() (v PublicIndexedTimePoint) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicTimePointOperationTimePointUnion) AsPropertyReferenced() (v PublicPropertyReferencedTime) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u PublicTimePointOperationTimePointUnion) RawJSON() string { return u.JSON.raw }

func (r *PublicTimePointOperationTimePointUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties IncludeObjectsWithNoValueSet, OperationType, Operator, TimePoint,
// Type are required.
type PublicTimePointOperationParam struct {
	// Indicates whether objects with no value set for the property should be included
	// in the operation.
	IncludeObjectsWithNoValueSet bool `json:"includeObjectsWithNoValueSet" api:"required"`
	// Specifies the type of operation (TIME_POINT).
	//
	// Any of "TIME_POINT".
	OperationType PublicTimePointOperationOperationType `json:"operationType,omitzero" api:"required"`
	// Specifies the operation to be applied within the time point operation
	// (IS_BEFORE, IS_AFTER).
	Operator string `json:"operator" api:"required"`
	// Defines the specific point in time for the operation, which can be a date,
	// indexed time, or property-referenced time.
	TimePoint PublicTimePointOperationTimePointUnionParam `json:"timePoint,omitzero" api:"required"`
	// Defines the type of operation being performed.
	Type string `json:"type" api:"required"`
	// Describes the behavior at the endpoint of the time point operation.
	EndpointBehavior param.Opt[string] `json:"endpointBehavior,omitzero"`
	// Specifies the parser used for interpreting the property in the operation.
	PropertyParser param.Opt[string] `json:"propertyParser,omitzero"`
	paramObj
}

func (r PublicTimePointOperationParam) MarshalJSON() (data []byte, err error) {
	type shadow PublicTimePointOperationParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PublicTimePointOperationParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type PublicTimePointOperationTimePointUnionParam struct {
	OfDate               *PublicDatePointParam              `json:",omitzero,inline"`
	OfIndexed            *PublicIndexedTimePointParam       `json:",omitzero,inline"`
	OfPropertyReferenced *PublicPropertyReferencedTimeParam `json:",omitzero,inline"`
	paramUnion
}

func (u PublicTimePointOperationTimePointUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfDate, u.OfIndexed, u.OfPropertyReferenced)
}
func (u *PublicTimePointOperationTimePointUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

type PublicTodayReference struct {
	// Indicates the type of reference (TODAY).
	//
	// Any of "TODAY".
	ReferenceType PublicTodayReferenceReferenceType `json:"referenceType" api:"required"`
	// The hour component of the current day reference.
	Hour int64 `json:"hour"`
	// The millisecond component of the current day reference.
	Millisecond int64 `json:"millisecond"`
	// The minute component of the current day reference.
	Minute int64 `json:"minute"`
	// The second component of the current day reference.
	Second int64 `json:"second"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ReferenceType respjson.Field
		Hour          respjson.Field
		Millisecond   respjson.Field
		Minute        respjson.Field
		Second        respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PublicTodayReference) RawJSON() string { return r.JSON.raw }
func (r *PublicTodayReference) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this PublicTodayReference to a PublicTodayReferenceParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// PublicTodayReferenceParam.Overrides()
func (r PublicTodayReference) ToParam() PublicTodayReferenceParam {
	return param.Override[PublicTodayReferenceParam](json.RawMessage(r.RawJSON()))
}

// Indicates the type of reference (TODAY).
type PublicTodayReferenceReferenceType string

const (
	PublicTodayReferenceReferenceTypeToday PublicTodayReferenceReferenceType = "TODAY"
)

// The property ReferenceType is required.
type PublicTodayReferenceParam struct {
	// Indicates the type of reference (TODAY).
	//
	// Any of "TODAY".
	ReferenceType PublicTodayReferenceReferenceType `json:"referenceType,omitzero" api:"required"`
	// The hour component of the current day reference.
	Hour param.Opt[int64] `json:"hour,omitzero"`
	// The millisecond component of the current day reference.
	Millisecond param.Opt[int64] `json:"millisecond,omitzero"`
	// The minute component of the current day reference.
	Minute param.Opt[int64] `json:"minute,omitzero"`
	// The second component of the current day reference.
	Second param.Opt[int64] `json:"second,omitzero"`
	paramObj
}

func (r PublicTodayReferenceParam) MarshalJSON() (data []byte, err error) {
	type shadow PublicTodayReferenceParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PublicTodayReferenceParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PublicUnifiedEventsFilter struct {
	FilterLines []PublicEventFilterMetadata `json:"filterLines" api:"required"`
	// Indicates the type of filter being applied (UNIFIED_EVENTS).
	//
	// Any of "UNIFIED_EVENTS".
	FilterType PublicUnifiedEventsFilterFilterType `json:"filterType" api:"required"`
	// Specifies the criteria for refining the filter by coalescing.
	CoalescingRefineBy PublicUnifiedEventsFilterCoalescingRefineByUnion `json:"coalescingRefineBy"`
	// The identifier for the type of event in the unified events filter.
	EventTypeID string `json:"eventTypeId"`
	// Specifies the criteria for refining the filter by pruning.
	PruningRefineBy PublicUnifiedEventsFilterPruningRefineByUnion `json:"pruningRefineBy"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FilterLines        respjson.Field
		FilterType         respjson.Field
		CoalescingRefineBy respjson.Field
		EventTypeID        respjson.Field
		PruningRefineBy    respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PublicUnifiedEventsFilter) RawJSON() string { return r.JSON.raw }
func (r *PublicUnifiedEventsFilter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this PublicUnifiedEventsFilter to a
// PublicUnifiedEventsFilterParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// PublicUnifiedEventsFilterParam.Overrides()
func (r PublicUnifiedEventsFilter) ToParam() PublicUnifiedEventsFilterParam {
	return param.Override[PublicUnifiedEventsFilterParam](json.RawMessage(r.RawJSON()))
}

// Indicates the type of filter being applied (UNIFIED_EVENTS).
type PublicUnifiedEventsFilterFilterType string

const (
	PublicUnifiedEventsFilterFilterTypeUnifiedEvents PublicUnifiedEventsFilterFilterType = "UNIFIED_EVENTS"
)

// PublicUnifiedEventsFilterCoalescingRefineByUnion contains all possible
// properties and values from [PublicNumOccurrencesRefineBy],
// [PublicSetOccurrencesRefineBy], [PublicRelativeComparativeTimestampRefineBy],
// [PublicRelativeRangedTimestampRefineBy],
// [PublicAbsoluteComparativeTimestampRefineBy],
// [PublicAbsoluteRangedTimestampRefineBy], [PublicAllHistoryRefineBy],
// [PublicTimePointOperation], [PublicRangedTimeOperation].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type PublicUnifiedEventsFilterCoalescingRefineByUnion struct {
	Type string `json:"type"`
	// This field is from variant [PublicNumOccurrencesRefineBy].
	MaxOccurrences int64 `json:"maxOccurrences"`
	// This field is from variant [PublicNumOccurrencesRefineBy].
	MinOccurrences int64 `json:"minOccurrences"`
	// This field is from variant [PublicSetOccurrencesRefineBy].
	SetType    string `json:"setType"`
	Comparison string `json:"comparison"`
	// This field is from variant [PublicRelativeComparativeTimestampRefineBy].
	TimeOffset PublicTimeOffset `json:"timeOffset"`
	// This field is from variant [PublicRelativeRangedTimestampRefineBy].
	LowerBoundOffset PublicTimeOffset `json:"lowerBoundOffset"`
	RangeType        string           `json:"rangeType"`
	// This field is from variant [PublicRelativeRangedTimestampRefineBy].
	UpperBoundOffset PublicTimeOffset `json:"upperBoundOffset"`
	// This field is from variant [PublicAbsoluteComparativeTimestampRefineBy].
	Timestamp int64 `json:"timestamp"`
	// This field is from variant [PublicAbsoluteRangedTimestampRefineBy].
	LowerTimestamp int64 `json:"lowerTimestamp"`
	// This field is from variant [PublicAbsoluteRangedTimestampRefineBy].
	UpperTimestamp               int64  `json:"upperTimestamp"`
	IncludeObjectsWithNoValueSet bool   `json:"includeObjectsWithNoValueSet"`
	OperationType                string `json:"operationType"`
	Operator                     string `json:"operator"`
	// This field is from variant [PublicTimePointOperation].
	TimePoint PublicTimePointOperationTimePointUnion `json:"timePoint"`
	// This field is from variant [PublicTimePointOperation].
	EndpointBehavior string `json:"endpointBehavior"`
	PropertyParser   string `json:"propertyParser"`
	// This field is from variant [PublicRangedTimeOperation].
	LowerBoundTimePoint PublicRangedTimeOperationLowerBoundTimePointUnion `json:"lowerBoundTimePoint"`
	// This field is from variant [PublicRangedTimeOperation].
	UpperBoundTimePoint PublicRangedTimeOperationUpperBoundTimePointUnion `json:"upperBoundTimePoint"`
	// This field is from variant [PublicRangedTimeOperation].
	LowerBoundEndpointBehavior string `json:"lowerBoundEndpointBehavior"`
	// This field is from variant [PublicRangedTimeOperation].
	UpperBoundEndpointBehavior string `json:"upperBoundEndpointBehavior"`
	JSON                       struct {
		Type                         respjson.Field
		MaxOccurrences               respjson.Field
		MinOccurrences               respjson.Field
		SetType                      respjson.Field
		Comparison                   respjson.Field
		TimeOffset                   respjson.Field
		LowerBoundOffset             respjson.Field
		RangeType                    respjson.Field
		UpperBoundOffset             respjson.Field
		Timestamp                    respjson.Field
		LowerTimestamp               respjson.Field
		UpperTimestamp               respjson.Field
		IncludeObjectsWithNoValueSet respjson.Field
		OperationType                respjson.Field
		Operator                     respjson.Field
		TimePoint                    respjson.Field
		EndpointBehavior             respjson.Field
		PropertyParser               respjson.Field
		LowerBoundTimePoint          respjson.Field
		UpperBoundTimePoint          respjson.Field
		LowerBoundEndpointBehavior   respjson.Field
		UpperBoundEndpointBehavior   respjson.Field
		raw                          string
	} `json:"-"`
}

func (u PublicUnifiedEventsFilterCoalescingRefineByUnion) AsNumOccurrences() (v PublicNumOccurrencesRefineBy) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicUnifiedEventsFilterCoalescingRefineByUnion) AsSetOccurrences() (v PublicSetOccurrencesRefineBy) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicUnifiedEventsFilterCoalescingRefineByUnion) AsRelativeComparative() (v PublicRelativeComparativeTimestampRefineBy) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicUnifiedEventsFilterCoalescingRefineByUnion) AsRelativeRanged() (v PublicRelativeRangedTimestampRefineBy) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicUnifiedEventsFilterCoalescingRefineByUnion) AsAbsoluteComparative() (v PublicAbsoluteComparativeTimestampRefineBy) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicUnifiedEventsFilterCoalescingRefineByUnion) AsAbsoluteRanged() (v PublicAbsoluteRangedTimestampRefineBy) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicUnifiedEventsFilterCoalescingRefineByUnion) AsAllHistory() (v PublicAllHistoryRefineBy) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicUnifiedEventsFilterCoalescingRefineByUnion) AsTimePoint() (v PublicTimePointOperation) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicUnifiedEventsFilterCoalescingRefineByUnion) AsTimeRanged() (v PublicRangedTimeOperation) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u PublicUnifiedEventsFilterCoalescingRefineByUnion) RawJSON() string { return u.JSON.raw }

func (r *PublicUnifiedEventsFilterCoalescingRefineByUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PublicUnifiedEventsFilterPruningRefineByUnion contains all possible properties
// and values from [PublicNumOccurrencesRefineBy], [PublicSetOccurrencesRefineBy],
// [PublicRelativeComparativeTimestampRefineBy],
// [PublicRelativeRangedTimestampRefineBy],
// [PublicAbsoluteComparativeTimestampRefineBy],
// [PublicAbsoluteRangedTimestampRefineBy], [PublicAllHistoryRefineBy],
// [PublicTimePointOperation], [PublicRangedTimeOperation].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type PublicUnifiedEventsFilterPruningRefineByUnion struct {
	Type string `json:"type"`
	// This field is from variant [PublicNumOccurrencesRefineBy].
	MaxOccurrences int64 `json:"maxOccurrences"`
	// This field is from variant [PublicNumOccurrencesRefineBy].
	MinOccurrences int64 `json:"minOccurrences"`
	// This field is from variant [PublicSetOccurrencesRefineBy].
	SetType    string `json:"setType"`
	Comparison string `json:"comparison"`
	// This field is from variant [PublicRelativeComparativeTimestampRefineBy].
	TimeOffset PublicTimeOffset `json:"timeOffset"`
	// This field is from variant [PublicRelativeRangedTimestampRefineBy].
	LowerBoundOffset PublicTimeOffset `json:"lowerBoundOffset"`
	RangeType        string           `json:"rangeType"`
	// This field is from variant [PublicRelativeRangedTimestampRefineBy].
	UpperBoundOffset PublicTimeOffset `json:"upperBoundOffset"`
	// This field is from variant [PublicAbsoluteComparativeTimestampRefineBy].
	Timestamp int64 `json:"timestamp"`
	// This field is from variant [PublicAbsoluteRangedTimestampRefineBy].
	LowerTimestamp int64 `json:"lowerTimestamp"`
	// This field is from variant [PublicAbsoluteRangedTimestampRefineBy].
	UpperTimestamp               int64  `json:"upperTimestamp"`
	IncludeObjectsWithNoValueSet bool   `json:"includeObjectsWithNoValueSet"`
	OperationType                string `json:"operationType"`
	Operator                     string `json:"operator"`
	// This field is from variant [PublicTimePointOperation].
	TimePoint PublicTimePointOperationTimePointUnion `json:"timePoint"`
	// This field is from variant [PublicTimePointOperation].
	EndpointBehavior string `json:"endpointBehavior"`
	PropertyParser   string `json:"propertyParser"`
	// This field is from variant [PublicRangedTimeOperation].
	LowerBoundTimePoint PublicRangedTimeOperationLowerBoundTimePointUnion `json:"lowerBoundTimePoint"`
	// This field is from variant [PublicRangedTimeOperation].
	UpperBoundTimePoint PublicRangedTimeOperationUpperBoundTimePointUnion `json:"upperBoundTimePoint"`
	// This field is from variant [PublicRangedTimeOperation].
	LowerBoundEndpointBehavior string `json:"lowerBoundEndpointBehavior"`
	// This field is from variant [PublicRangedTimeOperation].
	UpperBoundEndpointBehavior string `json:"upperBoundEndpointBehavior"`
	JSON                       struct {
		Type                         respjson.Field
		MaxOccurrences               respjson.Field
		MinOccurrences               respjson.Field
		SetType                      respjson.Field
		Comparison                   respjson.Field
		TimeOffset                   respjson.Field
		LowerBoundOffset             respjson.Field
		RangeType                    respjson.Field
		UpperBoundOffset             respjson.Field
		Timestamp                    respjson.Field
		LowerTimestamp               respjson.Field
		UpperTimestamp               respjson.Field
		IncludeObjectsWithNoValueSet respjson.Field
		OperationType                respjson.Field
		Operator                     respjson.Field
		TimePoint                    respjson.Field
		EndpointBehavior             respjson.Field
		PropertyParser               respjson.Field
		LowerBoundTimePoint          respjson.Field
		UpperBoundTimePoint          respjson.Field
		LowerBoundEndpointBehavior   respjson.Field
		UpperBoundEndpointBehavior   respjson.Field
		raw                          string
	} `json:"-"`
}

func (u PublicUnifiedEventsFilterPruningRefineByUnion) AsNumOccurrences() (v PublicNumOccurrencesRefineBy) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicUnifiedEventsFilterPruningRefineByUnion) AsSetOccurrences() (v PublicSetOccurrencesRefineBy) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicUnifiedEventsFilterPruningRefineByUnion) AsRelativeComparative() (v PublicRelativeComparativeTimestampRefineBy) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicUnifiedEventsFilterPruningRefineByUnion) AsRelativeRanged() (v PublicRelativeRangedTimestampRefineBy) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicUnifiedEventsFilterPruningRefineByUnion) AsAbsoluteComparative() (v PublicAbsoluteComparativeTimestampRefineBy) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicUnifiedEventsFilterPruningRefineByUnion) AsAbsoluteRanged() (v PublicAbsoluteRangedTimestampRefineBy) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicUnifiedEventsFilterPruningRefineByUnion) AsAllHistory() (v PublicAllHistoryRefineBy) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicUnifiedEventsFilterPruningRefineByUnion) AsTimePoint() (v PublicTimePointOperation) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicUnifiedEventsFilterPruningRefineByUnion) AsTimeRanged() (v PublicRangedTimeOperation) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u PublicUnifiedEventsFilterPruningRefineByUnion) RawJSON() string { return u.JSON.raw }

func (r *PublicUnifiedEventsFilterPruningRefineByUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties FilterLines, FilterType are required.
type PublicUnifiedEventsFilterParam struct {
	FilterLines []PublicEventFilterMetadataParam `json:"filterLines,omitzero" api:"required"`
	// Indicates the type of filter being applied (UNIFIED_EVENTS).
	//
	// Any of "UNIFIED_EVENTS".
	FilterType PublicUnifiedEventsFilterFilterType `json:"filterType,omitzero" api:"required"`
	// The identifier for the type of event in the unified events filter.
	EventTypeID param.Opt[string] `json:"eventTypeId,omitzero"`
	// Specifies the criteria for refining the filter by coalescing.
	CoalescingRefineBy PublicUnifiedEventsFilterCoalescingRefineByUnionParam `json:"coalescingRefineBy,omitzero"`
	// Specifies the criteria for refining the filter by pruning.
	PruningRefineBy PublicUnifiedEventsFilterPruningRefineByUnionParam `json:"pruningRefineBy,omitzero"`
	paramObj
}

func (r PublicUnifiedEventsFilterParam) MarshalJSON() (data []byte, err error) {
	type shadow PublicUnifiedEventsFilterParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PublicUnifiedEventsFilterParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type PublicUnifiedEventsFilterCoalescingRefineByUnionParam struct {
	OfNumOccurrences      *PublicNumOccurrencesRefineByParam               `json:",omitzero,inline"`
	OfSetOccurrences      *PublicSetOccurrencesRefineByParam               `json:",omitzero,inline"`
	OfRelativeComparative *PublicRelativeComparativeTimestampRefineByParam `json:",omitzero,inline"`
	OfRelativeRanged      *PublicRelativeRangedTimestampRefineByParam      `json:",omitzero,inline"`
	OfAbsoluteComparative *PublicAbsoluteComparativeTimestampRefineByParam `json:",omitzero,inline"`
	OfAbsoluteRanged      *PublicAbsoluteRangedTimestampRefineByParam      `json:",omitzero,inline"`
	OfAllHistory          *PublicAllHistoryRefineByParam                   `json:",omitzero,inline"`
	OfTimePoint           *PublicTimePointOperationParam                   `json:",omitzero,inline"`
	OfTimeRanged          *PublicRangedTimeOperationParam                  `json:",omitzero,inline"`
	paramUnion
}

func (u PublicUnifiedEventsFilterCoalescingRefineByUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfNumOccurrences,
		u.OfSetOccurrences,
		u.OfRelativeComparative,
		u.OfRelativeRanged,
		u.OfAbsoluteComparative,
		u.OfAbsoluteRanged,
		u.OfAllHistory,
		u.OfTimePoint,
		u.OfTimeRanged)
}
func (u *PublicUnifiedEventsFilterCoalescingRefineByUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type PublicUnifiedEventsFilterPruningRefineByUnionParam struct {
	OfNumOccurrences      *PublicNumOccurrencesRefineByParam               `json:",omitzero,inline"`
	OfSetOccurrences      *PublicSetOccurrencesRefineByParam               `json:",omitzero,inline"`
	OfRelativeComparative *PublicRelativeComparativeTimestampRefineByParam `json:",omitzero,inline"`
	OfRelativeRanged      *PublicRelativeRangedTimestampRefineByParam      `json:",omitzero,inline"`
	OfAbsoluteComparative *PublicAbsoluteComparativeTimestampRefineByParam `json:",omitzero,inline"`
	OfAbsoluteRanged      *PublicAbsoluteRangedTimestampRefineByParam      `json:",omitzero,inline"`
	OfAllHistory          *PublicAllHistoryRefineByParam                   `json:",omitzero,inline"`
	OfTimePoint           *PublicTimePointOperationParam                   `json:",omitzero,inline"`
	OfTimeRanged          *PublicRangedTimeOperationParam                  `json:",omitzero,inline"`
	paramUnion
}

func (u PublicUnifiedEventsFilterPruningRefineByUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfNumOccurrences,
		u.OfSetOccurrences,
		u.OfRelativeComparative,
		u.OfRelativeRanged,
		u.OfAbsoluteComparative,
		u.OfAbsoluteRanged,
		u.OfAllHistory,
		u.OfTimePoint,
		u.OfTimeRanged)
}
func (u *PublicUnifiedEventsFilterPruningRefineByUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

type PublicUnifiedEventsFilterBranch struct {
	// The identifier for the type of event associated with the filter branch.
	EventTypeID    string                                             `json:"eventTypeId" api:"required"`
	FilterBranches []PublicUnifiedEventsFilterBranchFilterBranchUnion `json:"filterBranches" api:"required"`
	// The logical operator used to combine filters within the branch (AND).
	FilterBranchOperator string `json:"filterBranchOperator" api:"required"`
	// The type of the filter branch (UNIFIED_EVENTS).
	//
	// Any of "UNIFIED_EVENTS".
	FilterBranchType PublicUnifiedEventsFilterBranchFilterBranchType `json:"filterBranchType" api:"required"`
	Filters          []PublicUnifiedEventsFilterBranchFilterUnion    `json:"filters" api:"required"`
	// Defines the operation to be applied within the filter branch (HAS_COMPLETED,
	// HAS_NOT_COMPLETED).
	//
	// Any of "HAS_COMPLETED", "HAS_NOT_COMPLETED".
	Operator PublicUnifiedEventsFilterBranchOperator `json:"operator" api:"required"`
	// Specifies the criteria for refining the filter by coalescing.
	CoalescingRefineBy PublicUnifiedEventsFilterBranchCoalescingRefineByUnion `json:"coalescingRefineBy"`
	PruningRefineBy    PublicUnifiedEventsFilterBranchPruningRefineByUnion    `json:"pruningRefineBy"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		EventTypeID          respjson.Field
		FilterBranches       respjson.Field
		FilterBranchOperator respjson.Field
		FilterBranchType     respjson.Field
		Filters              respjson.Field
		Operator             respjson.Field
		CoalescingRefineBy   respjson.Field
		PruningRefineBy      respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PublicUnifiedEventsFilterBranch) RawJSON() string { return r.JSON.raw }
func (r *PublicUnifiedEventsFilterBranch) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this PublicUnifiedEventsFilterBranch to a
// PublicUnifiedEventsFilterBranchParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// PublicUnifiedEventsFilterBranchParam.Overrides()
func (r PublicUnifiedEventsFilterBranch) ToParam() PublicUnifiedEventsFilterBranchParam {
	return param.Override[PublicUnifiedEventsFilterBranchParam](json.RawMessage(r.RawJSON()))
}

// PublicUnifiedEventsFilterBranchFilterBranchUnion contains all possible
// properties and values from [PublicOrFilterBranch], [PublicAndFilterBranch],
// [PublicNotAllFilterBranch], [PublicNotAnyFilterBranch],
// [PublicRestrictedFilterBranch], [PublicUnifiedEventsFilterBranch],
// [PublicPropertyAssociationFilterBranch], [PublicAssociationFilterBranch].
//
// Use the [PublicUnifiedEventsFilterBranchFilterBranchUnion.AsAny] method to
// switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type PublicUnifiedEventsFilterBranchFilterBranchUnion struct {
	// This field is a union of [[]PublicOrFilterBranchFilterBranchUnion],
	// [[]PublicAndFilterBranchFilterBranchUnion],
	// [[]PublicNotAllFilterBranchFilterBranchUnion],
	// [[]PublicNotAnyFilterBranchFilterBranchUnion],
	// [[]PublicRestrictedFilterBranchFilterBranchUnion],
	// [[]PublicUnifiedEventsFilterBranchFilterBranchUnion],
	// [[]PublicPropertyAssociationFilterBranchFilterBranchUnion],
	// [[]PublicAssociationFilterBranchFilterBranchUnion]
	FilterBranches       PublicUnifiedEventsFilterBranchFilterBranchUnionFilterBranches `json:"filterBranches"`
	FilterBranchOperator string                                                         `json:"filterBranchOperator"`
	// Any of "OR", "AND", "NOT_ALL", "NOT_ANY", "RESTRICTED", "UNIFIED_EVENTS",
	// "PROPERTY_ASSOCIATION", "ASSOCIATION".
	FilterBranchType string `json:"filterBranchType"`
	// This field is a union of [[]PublicOrFilterBranchFilterUnion],
	// [[]PublicAndFilterBranchFilterUnion], [[]PublicNotAllFilterBranchFilterUnion],
	// [[]PublicNotAnyFilterBranchFilterUnion],
	// [[]PublicRestrictedFilterBranchFilterUnion],
	// [[]PublicUnifiedEventsFilterBranchFilterUnion],
	// [[]PublicPropertyAssociationFilterBranchFilterUnion],
	// [[]PublicAssociationFilterBranchFilterUnion]
	Filters PublicUnifiedEventsFilterBranchFilterBranchUnionFilters `json:"filters"`
	// This field is from variant [PublicUnifiedEventsFilterBranch].
	EventTypeID string `json:"eventTypeId"`
	Operator    string `json:"operator"`
	// This field is from variant [PublicUnifiedEventsFilterBranch].
	CoalescingRefineBy PublicUnifiedEventsFilterBranchCoalescingRefineByUnion `json:"coalescingRefineBy"`
	// This field is from variant [PublicUnifiedEventsFilterBranch].
	PruningRefineBy PublicUnifiedEventsFilterBranchPruningRefineByUnion `json:"pruningRefineBy"`
	ObjectTypeID    string                                              `json:"objectTypeId"`
	// This field is from variant [PublicPropertyAssociationFilterBranch].
	PropertyWithObjectID string `json:"propertyWithObjectId"`
	// This field is from variant [PublicAssociationFilterBranch].
	AssociationCategory string `json:"associationCategory"`
	// This field is from variant [PublicAssociationFilterBranch].
	AssociationTypeID int64 `json:"associationTypeId"`
	JSON              struct {
		FilterBranches       respjson.Field
		FilterBranchOperator respjson.Field
		FilterBranchType     respjson.Field
		Filters              respjson.Field
		EventTypeID          respjson.Field
		Operator             respjson.Field
		CoalescingRefineBy   respjson.Field
		PruningRefineBy      respjson.Field
		ObjectTypeID         respjson.Field
		PropertyWithObjectID respjson.Field
		AssociationCategory  respjson.Field
		AssociationTypeID    respjson.Field
		raw                  string
	} `json:"-"`
}

// anyPublicUnifiedEventsFilterBranchFilterBranch is implemented by each variant of
// [PublicUnifiedEventsFilterBranchFilterBranchUnion] to add type safety for the
// return type of [PublicUnifiedEventsFilterBranchFilterBranchUnion.AsAny]
type anyPublicUnifiedEventsFilterBranchFilterBranch interface {
	implPublicUnifiedEventsFilterBranchFilterBranchUnion()
}

func (PublicOrFilterBranch) implPublicUnifiedEventsFilterBranchFilterBranchUnion()                  {}
func (PublicAndFilterBranch) implPublicUnifiedEventsFilterBranchFilterBranchUnion()                 {}
func (PublicNotAllFilterBranch) implPublicUnifiedEventsFilterBranchFilterBranchUnion()              {}
func (PublicNotAnyFilterBranch) implPublicUnifiedEventsFilterBranchFilterBranchUnion()              {}
func (PublicRestrictedFilterBranch) implPublicUnifiedEventsFilterBranchFilterBranchUnion()          {}
func (PublicUnifiedEventsFilterBranch) implPublicUnifiedEventsFilterBranchFilterBranchUnion()       {}
func (PublicPropertyAssociationFilterBranch) implPublicUnifiedEventsFilterBranchFilterBranchUnion() {}
func (PublicAssociationFilterBranch) implPublicUnifiedEventsFilterBranchFilterBranchUnion()         {}

// Use the following switch statement to find the correct variant
//
//	switch variant := PublicUnifiedEventsFilterBranchFilterBranchUnion.AsAny().(type) {
//	case crm.PublicOrFilterBranch:
//	case crm.PublicAndFilterBranch:
//	case crm.PublicNotAllFilterBranch:
//	case crm.PublicNotAnyFilterBranch:
//	case crm.PublicRestrictedFilterBranch:
//	case crm.PublicUnifiedEventsFilterBranch:
//	case crm.PublicPropertyAssociationFilterBranch:
//	case crm.PublicAssociationFilterBranch:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u PublicUnifiedEventsFilterBranchFilterBranchUnion) AsAny() anyPublicUnifiedEventsFilterBranchFilterBranch {
	switch u.FilterBranchType {
	case "OR":
		return u.AsOr()
	case "AND":
		return u.AsAnd()
	case "NOT_ALL":
		return u.AsNotAll()
	case "NOT_ANY":
		return u.AsNotAny()
	case "RESTRICTED":
		return u.AsRestricted()
	case "UNIFIED_EVENTS":
		return u.AsUnifiedEvents()
	case "PROPERTY_ASSOCIATION":
		return u.AsPropertyAssociation()
	case "ASSOCIATION":
		return u.AsAssociation()
	}
	return nil
}

func (u PublicUnifiedEventsFilterBranchFilterBranchUnion) AsOr() (v PublicOrFilterBranch) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicUnifiedEventsFilterBranchFilterBranchUnion) AsAnd() (v PublicAndFilterBranch) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicUnifiedEventsFilterBranchFilterBranchUnion) AsNotAll() (v PublicNotAllFilterBranch) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicUnifiedEventsFilterBranchFilterBranchUnion) AsNotAny() (v PublicNotAnyFilterBranch) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicUnifiedEventsFilterBranchFilterBranchUnion) AsRestricted() (v PublicRestrictedFilterBranch) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicUnifiedEventsFilterBranchFilterBranchUnion) AsUnifiedEvents() (v PublicUnifiedEventsFilterBranch) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicUnifiedEventsFilterBranchFilterBranchUnion) AsPropertyAssociation() (v PublicPropertyAssociationFilterBranch) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicUnifiedEventsFilterBranchFilterBranchUnion) AsAssociation() (v PublicAssociationFilterBranch) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u PublicUnifiedEventsFilterBranchFilterBranchUnion) RawJSON() string { return u.JSON.raw }

func (r *PublicUnifiedEventsFilterBranchFilterBranchUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PublicUnifiedEventsFilterBranchFilterBranchUnionFilterBranches is an implicit
// subunion of [PublicUnifiedEventsFilterBranchFilterBranchUnion].
// PublicUnifiedEventsFilterBranchFilterBranchUnionFilterBranches provides
// convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [PublicUnifiedEventsFilterBranchFilterBranchUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfPublicOrFilterBranchFilterBranches
// OfPublicAndFilterBranchFilterBranches OfPublicNotAllFilterBranchFilterBranches
// OfPublicNotAnyFilterBranchFilterBranches
// OfPublicRestrictedFilterBranchFilterBranches
// OfPublicUnifiedEventsFilterBranchFilterBranches
// OfPublicPropertyAssociationFilterBranchFilterBranches
// OfPublicAssociationFilterBranchFilterBranches]
type PublicUnifiedEventsFilterBranchFilterBranchUnionFilterBranches struct {
	// This field will be present if the value is a
	// [[]PublicOrFilterBranchFilterBranchUnion] instead of an object.
	OfPublicOrFilterBranchFilterBranches []PublicOrFilterBranchFilterBranchUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]PublicAndFilterBranchFilterBranchUnion] instead of an object.
	OfPublicAndFilterBranchFilterBranches []PublicAndFilterBranchFilterBranchUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]PublicNotAllFilterBranchFilterBranchUnion] instead of an object.
	OfPublicNotAllFilterBranchFilterBranches []PublicNotAllFilterBranchFilterBranchUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]PublicNotAnyFilterBranchFilterBranchUnion] instead of an object.
	OfPublicNotAnyFilterBranchFilterBranches []PublicNotAnyFilterBranchFilterBranchUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]PublicRestrictedFilterBranchFilterBranchUnion] instead of an object.
	OfPublicRestrictedFilterBranchFilterBranches []PublicRestrictedFilterBranchFilterBranchUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]PublicUnifiedEventsFilterBranchFilterBranchUnion] instead of an object.
	OfPublicUnifiedEventsFilterBranchFilterBranches []PublicUnifiedEventsFilterBranchFilterBranchUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]PublicPropertyAssociationFilterBranchFilterBranchUnion] instead of an object.
	OfPublicPropertyAssociationFilterBranchFilterBranches []PublicPropertyAssociationFilterBranchFilterBranchUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]PublicAssociationFilterBranchFilterBranchUnion] instead of an object.
	OfPublicAssociationFilterBranchFilterBranches []PublicAssociationFilterBranchFilterBranchUnion `json:",inline"`
	JSON                                          struct {
		OfPublicOrFilterBranchFilterBranches                  respjson.Field
		OfPublicAndFilterBranchFilterBranches                 respjson.Field
		OfPublicNotAllFilterBranchFilterBranches              respjson.Field
		OfPublicNotAnyFilterBranchFilterBranches              respjson.Field
		OfPublicRestrictedFilterBranchFilterBranches          respjson.Field
		OfPublicUnifiedEventsFilterBranchFilterBranches       respjson.Field
		OfPublicPropertyAssociationFilterBranchFilterBranches respjson.Field
		OfPublicAssociationFilterBranchFilterBranches         respjson.Field
		raw                                                   string
	} `json:"-"`
}

func (r *PublicUnifiedEventsFilterBranchFilterBranchUnionFilterBranches) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PublicUnifiedEventsFilterBranchFilterBranchUnionFilters is an implicit subunion
// of [PublicUnifiedEventsFilterBranchFilterBranchUnion].
// PublicUnifiedEventsFilterBranchFilterBranchUnionFilters provides convenient
// access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [PublicUnifiedEventsFilterBranchFilterBranchUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfPublicOrFilterBranchFilters OfPublicAndFilterBranchFilters
// OfPublicNotAllFilterBranchFilters OfPublicNotAnyFilterBranchFilters
// OfPublicRestrictedFilterBranchFilters OfPublicUnifiedEventsFilterBranchFilters
// OfPublicPropertyAssociationFilterBranchFilters
// OfPublicAssociationFilterBranchFilters]
type PublicUnifiedEventsFilterBranchFilterBranchUnionFilters struct {
	// This field will be present if the value is a [[]PublicOrFilterBranchFilterUnion]
	// instead of an object.
	OfPublicOrFilterBranchFilters []PublicOrFilterBranchFilterUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]PublicAndFilterBranchFilterUnion] instead of an object.
	OfPublicAndFilterBranchFilters []PublicAndFilterBranchFilterUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]PublicNotAllFilterBranchFilterUnion] instead of an object.
	OfPublicNotAllFilterBranchFilters []PublicNotAllFilterBranchFilterUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]PublicNotAnyFilterBranchFilterUnion] instead of an object.
	OfPublicNotAnyFilterBranchFilters []PublicNotAnyFilterBranchFilterUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]PublicRestrictedFilterBranchFilterUnion] instead of an object.
	OfPublicRestrictedFilterBranchFilters []PublicRestrictedFilterBranchFilterUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]PublicUnifiedEventsFilterBranchFilterUnion] instead of an object.
	OfPublicUnifiedEventsFilterBranchFilters []PublicUnifiedEventsFilterBranchFilterUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]PublicPropertyAssociationFilterBranchFilterUnion] instead of an object.
	OfPublicPropertyAssociationFilterBranchFilters []PublicPropertyAssociationFilterBranchFilterUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]PublicAssociationFilterBranchFilterUnion] instead of an object.
	OfPublicAssociationFilterBranchFilters []PublicAssociationFilterBranchFilterUnion `json:",inline"`
	JSON                                   struct {
		OfPublicOrFilterBranchFilters                  respjson.Field
		OfPublicAndFilterBranchFilters                 respjson.Field
		OfPublicNotAllFilterBranchFilters              respjson.Field
		OfPublicNotAnyFilterBranchFilters              respjson.Field
		OfPublicRestrictedFilterBranchFilters          respjson.Field
		OfPublicUnifiedEventsFilterBranchFilters       respjson.Field
		OfPublicPropertyAssociationFilterBranchFilters respjson.Field
		OfPublicAssociationFilterBranchFilters         respjson.Field
		raw                                            string
	} `json:"-"`
}

func (r *PublicUnifiedEventsFilterBranchFilterBranchUnionFilters) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The type of the filter branch (UNIFIED_EVENTS).
type PublicUnifiedEventsFilterBranchFilterBranchType string

const (
	PublicUnifiedEventsFilterBranchFilterBranchTypeUnifiedEvents PublicUnifiedEventsFilterBranchFilterBranchType = "UNIFIED_EVENTS"
)

// PublicUnifiedEventsFilterBranchFilterUnion contains all possible properties and
// values from [PublicPropertyFilter], [PublicAssociationInListFilter],
// [PublicPageViewAnalyticsFilter], [PublicCtaAnalyticsFilter],
// [PublicEventAnalyticsFilter], [PublicFormSubmissionFilter],
// [PublicFormSubmissionOnPageFilter], [PublicIntegrationEventFilter],
// [PublicEmailSubscriptionFilter], [PublicCommunicationSubscriptionFilter],
// [PublicCampaignInfluencedFilter], [PublicSurveyMonkeyFilter],
// [PublicSurveyMonkeyValueFilter], [PublicWebinarFilter],
// [PublicEmailEventFilter], [PublicPrivacyAnalyticsFilter],
// [PublicAdsSearchFilter], [PublicAdsTimeFilter], [PublicInListFilter],
// [PublicNumAssociationsFilter], [PublicUnifiedEventsFilter],
// [PublicPropertyAssociationInListFilter], [PublicConstantFilter].
//
// Use the [PublicUnifiedEventsFilterBranchFilterUnion.AsAny] method to switch on
// the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type PublicUnifiedEventsFilterBranchFilterUnion struct {
	// Any of "PROPERTY", "ASSOCIATION", "PAGE_VIEW", "CTA", "EVENT",
	// "FORM_SUBMISSION", "FORM_SUBMISSION_ON_PAGE", "INTEGRATION_EVENT",
	// "EMAIL_SUBSCRIPTION", "COMMUNICATION_SUBSCRIPTION", "CAMPAIGN_INFLUENCED",
	// "SURVEY_MONKEY", "SURVEY_MONKEY_VALUE", "WEBINAR", "EMAIL_EVENT", "PRIVACY",
	// "ADS_SEARCH", "ADS_TIME", "IN_LIST", "NUM_ASSOCIATIONS", "UNIFIED_EVENTS",
	// "PROPERTY_ASSOCIATION", "CONSTANT".
	FilterType string `json:"filterType"`
	// This field is from variant [PublicPropertyFilter].
	Operation PublicPropertyFilterOperationUnion `json:"operation"`
	// This field is from variant [PublicPropertyFilter].
	Property            string `json:"property"`
	AssociationCategory string `json:"associationCategory"`
	AssociationTypeID   int64  `json:"associationTypeId"`
	// This field is a union of [PublicAssociationInListFilterCoalescingRefineByUnion],
	// [PublicPageViewAnalyticsFilterCoalescingRefineByUnion],
	// [PublicCtaAnalyticsFilterCoalescingRefineByUnion],
	// [PublicEventAnalyticsFilterCoalescingRefineByUnion],
	// [PublicFormSubmissionFilterCoalescingRefineByUnion],
	// [PublicFormSubmissionOnPageFilterCoalescingRefineByUnion],
	// [PublicNumAssociationsFilterCoalescingRefineByUnion],
	// [PublicUnifiedEventsFilterCoalescingRefineByUnion],
	// [PublicPropertyAssociationInListFilterCoalescingRefineByUnion]
	CoalescingRefineBy PublicUnifiedEventsFilterBranchFilterUnionCoalescingRefineBy `json:"coalescingRefineBy"`
	ListID             string                                                       `json:"listId"`
	Operator           string                                                       `json:"operator"`
	// This field is from variant [PublicAssociationInListFilter].
	ToObjectType   string `json:"toObjectType"`
	ToObjectTypeID string `json:"toObjectTypeId"`
	// This field is from variant [PublicPageViewAnalyticsFilter].
	PageURL string `json:"pageUrl"`
	// This field is from variant [PublicPageViewAnalyticsFilter].
	EnableTracking bool `json:"enableTracking"`
	// This field is a union of [PublicPageViewAnalyticsFilterPruningRefineByUnion],
	// [PublicCtaAnalyticsFilterPruningRefineByUnion],
	// [PublicEventAnalyticsFilterPruningRefineByUnion],
	// [PublicFormSubmissionFilterPruningRefineByUnion],
	// [PublicFormSubmissionOnPageFilterPruningRefineByUnion],
	// [PublicEmailEventFilterPruningRefineByUnion],
	// [PublicAdsTimeFilterPruningRefineByUnion],
	// [PublicUnifiedEventsFilterPruningRefineByUnion]
	PruningRefineBy PublicUnifiedEventsFilterBranchFilterUnionPruningRefineBy `json:"pruningRefineBy"`
	// This field is from variant [PublicCtaAnalyticsFilter].
	CtaName string `json:"ctaName"`
	// This field is from variant [PublicEventAnalyticsFilter].
	EventID string `json:"eventId"`
	FormID  string `json:"formId"`
	// This field is from variant [PublicFormSubmissionOnPageFilter].
	PageID string `json:"pageId"`
	// This field is a union of [int64], [string]
	EventTypeID PublicUnifiedEventsFilterBranchFilterUnionEventTypeID `json:"eventTypeId"`
	FilterLines []PublicEventFilterMetadata                           `json:"filterLines"`
	// This field is from variant [PublicEmailSubscriptionFilter].
	AcceptedStatuses []string `json:"acceptedStatuses"`
	SubscriptionIDs  []string `json:"subscriptionIds"`
	SubscriptionType string   `json:"subscriptionType"`
	// This field is from variant [PublicCommunicationSubscriptionFilter].
	AcceptedOptStates []string `json:"acceptedOptStates"`
	// This field is from variant [PublicCommunicationSubscriptionFilter].
	Channel string `json:"channel"`
	// This field is from variant [PublicCommunicationSubscriptionFilter].
	BusinessUnitID string `json:"businessUnitId"`
	// This field is from variant [PublicCampaignInfluencedFilter].
	CampaignID string `json:"campaignId"`
	SurveyID   string `json:"surveyId"`
	// This field is from variant [PublicSurveyMonkeyValueFilter].
	SurveyQuestion string `json:"surveyQuestion"`
	// This field is from variant [PublicSurveyMonkeyValueFilter].
	ValueComparison PublicSurveyMonkeyValueFilterValueComparisonUnion `json:"valueComparison"`
	// This field is from variant [PublicSurveyMonkeyValueFilter].
	SurveyAnswerColID string `json:"surveyAnswerColId"`
	// This field is from variant [PublicSurveyMonkeyValueFilter].
	SurveyAnswerRowID string `json:"surveyAnswerRowId"`
	// This field is from variant [PublicWebinarFilter].
	WebinarID string `json:"webinarId"`
	// This field is from variant [PublicEmailEventFilter].
	AppID string `json:"appId"`
	// This field is from variant [PublicEmailEventFilter].
	EmailID string `json:"emailId"`
	// This field is from variant [PublicEmailEventFilter].
	Level string `json:"level"`
	// This field is from variant [PublicEmailEventFilter].
	ClickURL string `json:"clickUrl"`
	// This field is from variant [PublicPrivacyAnalyticsFilter].
	PrivacyName string `json:"privacyName"`
	// This field is from variant [PublicAdsSearchFilter].
	AdNetwork string `json:"adNetwork"`
	// This field is from variant [PublicAdsSearchFilter].
	EntityType string `json:"entityType"`
	// This field is from variant [PublicAdsSearchFilter].
	SearchTerms []string `json:"searchTerms"`
	// This field is from variant [PublicAdsSearchFilter].
	SearchTermType string `json:"searchTermType"`
	// This field is from variant [PublicInListFilter].
	Metadata PublicInListFilterMetadata `json:"metadata"`
	// This field is from variant [PublicPropertyAssociationInListFilter].
	PropertyWithObjectID string `json:"propertyWithObjectId"`
	// This field is from variant [PublicConstantFilter].
	ShouldAccept bool `json:"shouldAccept"`
	// This field is from variant [PublicConstantFilter].
	Source string `json:"source"`
	JSON   struct {
		FilterType           respjson.Field
		Operation            respjson.Field
		Property             respjson.Field
		AssociationCategory  respjson.Field
		AssociationTypeID    respjson.Field
		CoalescingRefineBy   respjson.Field
		ListID               respjson.Field
		Operator             respjson.Field
		ToObjectType         respjson.Field
		ToObjectTypeID       respjson.Field
		PageURL              respjson.Field
		EnableTracking       respjson.Field
		PruningRefineBy      respjson.Field
		CtaName              respjson.Field
		EventID              respjson.Field
		FormID               respjson.Field
		PageID               respjson.Field
		EventTypeID          respjson.Field
		FilterLines          respjson.Field
		AcceptedStatuses     respjson.Field
		SubscriptionIDs      respjson.Field
		SubscriptionType     respjson.Field
		AcceptedOptStates    respjson.Field
		Channel              respjson.Field
		BusinessUnitID       respjson.Field
		CampaignID           respjson.Field
		SurveyID             respjson.Field
		SurveyQuestion       respjson.Field
		ValueComparison      respjson.Field
		SurveyAnswerColID    respjson.Field
		SurveyAnswerRowID    respjson.Field
		WebinarID            respjson.Field
		AppID                respjson.Field
		EmailID              respjson.Field
		Level                respjson.Field
		ClickURL             respjson.Field
		PrivacyName          respjson.Field
		AdNetwork            respjson.Field
		EntityType           respjson.Field
		SearchTerms          respjson.Field
		SearchTermType       respjson.Field
		Metadata             respjson.Field
		PropertyWithObjectID respjson.Field
		ShouldAccept         respjson.Field
		Source               respjson.Field
		raw                  string
	} `json:"-"`
}

// anyPublicUnifiedEventsFilterBranchFilter is implemented by each variant of
// [PublicUnifiedEventsFilterBranchFilterUnion] to add type safety for the return
// type of [PublicUnifiedEventsFilterBranchFilterUnion.AsAny]
type anyPublicUnifiedEventsFilterBranchFilter interface {
	implPublicUnifiedEventsFilterBranchFilterUnion()
}

func (PublicPropertyFilter) implPublicUnifiedEventsFilterBranchFilterUnion()                  {}
func (PublicAssociationInListFilter) implPublicUnifiedEventsFilterBranchFilterUnion()         {}
func (PublicPageViewAnalyticsFilter) implPublicUnifiedEventsFilterBranchFilterUnion()         {}
func (PublicCtaAnalyticsFilter) implPublicUnifiedEventsFilterBranchFilterUnion()              {}
func (PublicEventAnalyticsFilter) implPublicUnifiedEventsFilterBranchFilterUnion()            {}
func (PublicFormSubmissionFilter) implPublicUnifiedEventsFilterBranchFilterUnion()            {}
func (PublicFormSubmissionOnPageFilter) implPublicUnifiedEventsFilterBranchFilterUnion()      {}
func (PublicIntegrationEventFilter) implPublicUnifiedEventsFilterBranchFilterUnion()          {}
func (PublicEmailSubscriptionFilter) implPublicUnifiedEventsFilterBranchFilterUnion()         {}
func (PublicCommunicationSubscriptionFilter) implPublicUnifiedEventsFilterBranchFilterUnion() {}
func (PublicCampaignInfluencedFilter) implPublicUnifiedEventsFilterBranchFilterUnion()        {}
func (PublicSurveyMonkeyFilter) implPublicUnifiedEventsFilterBranchFilterUnion()              {}
func (PublicSurveyMonkeyValueFilter) implPublicUnifiedEventsFilterBranchFilterUnion()         {}
func (PublicWebinarFilter) implPublicUnifiedEventsFilterBranchFilterUnion()                   {}
func (PublicEmailEventFilter) implPublicUnifiedEventsFilterBranchFilterUnion()                {}
func (PublicPrivacyAnalyticsFilter) implPublicUnifiedEventsFilterBranchFilterUnion()          {}
func (PublicAdsSearchFilter) implPublicUnifiedEventsFilterBranchFilterUnion()                 {}
func (PublicAdsTimeFilter) implPublicUnifiedEventsFilterBranchFilterUnion()                   {}
func (PublicInListFilter) implPublicUnifiedEventsFilterBranchFilterUnion()                    {}
func (PublicNumAssociationsFilter) implPublicUnifiedEventsFilterBranchFilterUnion()           {}
func (PublicUnifiedEventsFilter) implPublicUnifiedEventsFilterBranchFilterUnion()             {}
func (PublicPropertyAssociationInListFilter) implPublicUnifiedEventsFilterBranchFilterUnion() {}
func (PublicConstantFilter) implPublicUnifiedEventsFilterBranchFilterUnion()                  {}

// Use the following switch statement to find the correct variant
//
//	switch variant := PublicUnifiedEventsFilterBranchFilterUnion.AsAny().(type) {
//	case crm.PublicPropertyFilter:
//	case crm.PublicAssociationInListFilter:
//	case crm.PublicPageViewAnalyticsFilter:
//	case crm.PublicCtaAnalyticsFilter:
//	case crm.PublicEventAnalyticsFilter:
//	case crm.PublicFormSubmissionFilter:
//	case crm.PublicFormSubmissionOnPageFilter:
//	case crm.PublicIntegrationEventFilter:
//	case crm.PublicEmailSubscriptionFilter:
//	case crm.PublicCommunicationSubscriptionFilter:
//	case crm.PublicCampaignInfluencedFilter:
//	case crm.PublicSurveyMonkeyFilter:
//	case crm.PublicSurveyMonkeyValueFilter:
//	case crm.PublicWebinarFilter:
//	case crm.PublicEmailEventFilter:
//	case crm.PublicPrivacyAnalyticsFilter:
//	case crm.PublicAdsSearchFilter:
//	case crm.PublicAdsTimeFilter:
//	case crm.PublicInListFilter:
//	case crm.PublicNumAssociationsFilter:
//	case crm.PublicUnifiedEventsFilter:
//	case crm.PublicPropertyAssociationInListFilter:
//	case crm.PublicConstantFilter:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u PublicUnifiedEventsFilterBranchFilterUnion) AsAny() anyPublicUnifiedEventsFilterBranchFilter {
	switch u.FilterType {
	case "PROPERTY":
		return u.AsProperty()
	case "ASSOCIATION":
		return u.AsAssociation()
	case "PAGE_VIEW":
		return u.AsPageView()
	case "CTA":
		return u.AsCta()
	case "EVENT":
		return u.AsEvent()
	case "FORM_SUBMISSION":
		return u.AsFormSubmission()
	case "FORM_SUBMISSION_ON_PAGE":
		return u.AsFormSubmissionOnPage()
	case "INTEGRATION_EVENT":
		return u.AsIntegrationEvent()
	case "EMAIL_SUBSCRIPTION":
		return u.AsEmailSubscription()
	case "COMMUNICATION_SUBSCRIPTION":
		return u.AsCommunicationSubscription()
	case "CAMPAIGN_INFLUENCED":
		return u.AsCampaignInfluenced()
	case "SURVEY_MONKEY":
		return u.AsSurveyMonkey()
	case "SURVEY_MONKEY_VALUE":
		return u.AsSurveyMonkeyValue()
	case "WEBINAR":
		return u.AsWebinar()
	case "EMAIL_EVENT":
		return u.AsEmailEvent()
	case "PRIVACY":
		return u.AsPrivacy()
	case "ADS_SEARCH":
		return u.AsAdsSearch()
	case "ADS_TIME":
		return u.AsAdsTime()
	case "IN_LIST":
		return u.AsInList()
	case "NUM_ASSOCIATIONS":
		return u.AsNumAssociations()
	case "UNIFIED_EVENTS":
		return u.AsUnifiedEvents()
	case "PROPERTY_ASSOCIATION":
		return u.AsPropertyAssociation()
	case "CONSTANT":
		return u.AsConstant()
	}
	return nil
}

func (u PublicUnifiedEventsFilterBranchFilterUnion) AsProperty() (v PublicPropertyFilter) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicUnifiedEventsFilterBranchFilterUnion) AsAssociation() (v PublicAssociationInListFilter) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicUnifiedEventsFilterBranchFilterUnion) AsPageView() (v PublicPageViewAnalyticsFilter) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicUnifiedEventsFilterBranchFilterUnion) AsCta() (v PublicCtaAnalyticsFilter) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicUnifiedEventsFilterBranchFilterUnion) AsEvent() (v PublicEventAnalyticsFilter) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicUnifiedEventsFilterBranchFilterUnion) AsFormSubmission() (v PublicFormSubmissionFilter) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicUnifiedEventsFilterBranchFilterUnion) AsFormSubmissionOnPage() (v PublicFormSubmissionOnPageFilter) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicUnifiedEventsFilterBranchFilterUnion) AsIntegrationEvent() (v PublicIntegrationEventFilter) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicUnifiedEventsFilterBranchFilterUnion) AsEmailSubscription() (v PublicEmailSubscriptionFilter) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicUnifiedEventsFilterBranchFilterUnion) AsCommunicationSubscription() (v PublicCommunicationSubscriptionFilter) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicUnifiedEventsFilterBranchFilterUnion) AsCampaignInfluenced() (v PublicCampaignInfluencedFilter) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicUnifiedEventsFilterBranchFilterUnion) AsSurveyMonkey() (v PublicSurveyMonkeyFilter) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicUnifiedEventsFilterBranchFilterUnion) AsSurveyMonkeyValue() (v PublicSurveyMonkeyValueFilter) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicUnifiedEventsFilterBranchFilterUnion) AsWebinar() (v PublicWebinarFilter) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicUnifiedEventsFilterBranchFilterUnion) AsEmailEvent() (v PublicEmailEventFilter) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicUnifiedEventsFilterBranchFilterUnion) AsPrivacy() (v PublicPrivacyAnalyticsFilter) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicUnifiedEventsFilterBranchFilterUnion) AsAdsSearch() (v PublicAdsSearchFilter) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicUnifiedEventsFilterBranchFilterUnion) AsAdsTime() (v PublicAdsTimeFilter) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicUnifiedEventsFilterBranchFilterUnion) AsInList() (v PublicInListFilter) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicUnifiedEventsFilterBranchFilterUnion) AsNumAssociations() (v PublicNumAssociationsFilter) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicUnifiedEventsFilterBranchFilterUnion) AsUnifiedEvents() (v PublicUnifiedEventsFilter) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicUnifiedEventsFilterBranchFilterUnion) AsPropertyAssociation() (v PublicPropertyAssociationInListFilter) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicUnifiedEventsFilterBranchFilterUnion) AsConstant() (v PublicConstantFilter) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u PublicUnifiedEventsFilterBranchFilterUnion) RawJSON() string { return u.JSON.raw }

func (r *PublicUnifiedEventsFilterBranchFilterUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PublicUnifiedEventsFilterBranchFilterUnionCoalescingRefineBy is an implicit
// subunion of [PublicUnifiedEventsFilterBranchFilterUnion].
// PublicUnifiedEventsFilterBranchFilterUnionCoalescingRefineBy provides convenient
// access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [PublicUnifiedEventsFilterBranchFilterUnion].
type PublicUnifiedEventsFilterBranchFilterUnionCoalescingRefineBy struct {
	Type string `json:"type"`
	// This field is from variant
	// [PublicAssociationInListFilterCoalescingRefineByUnion],
	// [PublicPageViewAnalyticsFilterCoalescingRefineByUnion],
	// [PublicCtaAnalyticsFilterCoalescingRefineByUnion],
	// [PublicEventAnalyticsFilterCoalescingRefineByUnion],
	// [PublicFormSubmissionFilterCoalescingRefineByUnion],
	// [PublicFormSubmissionOnPageFilterCoalescingRefineByUnion],
	// [PublicNumAssociationsFilterCoalescingRefineByUnion],
	// [PublicUnifiedEventsFilterCoalescingRefineByUnion],
	// [PublicPropertyAssociationInListFilterCoalescingRefineByUnion].
	MaxOccurrences int64 `json:"maxOccurrences"`
	// This field is from variant
	// [PublicAssociationInListFilterCoalescingRefineByUnion],
	// [PublicPageViewAnalyticsFilterCoalescingRefineByUnion],
	// [PublicCtaAnalyticsFilterCoalescingRefineByUnion],
	// [PublicEventAnalyticsFilterCoalescingRefineByUnion],
	// [PublicFormSubmissionFilterCoalescingRefineByUnion],
	// [PublicFormSubmissionOnPageFilterCoalescingRefineByUnion],
	// [PublicNumAssociationsFilterCoalescingRefineByUnion],
	// [PublicUnifiedEventsFilterCoalescingRefineByUnion],
	// [PublicPropertyAssociationInListFilterCoalescingRefineByUnion].
	MinOccurrences int64 `json:"minOccurrences"`
	// This field is from variant
	// [PublicAssociationInListFilterCoalescingRefineByUnion],
	// [PublicPageViewAnalyticsFilterCoalescingRefineByUnion],
	// [PublicCtaAnalyticsFilterCoalescingRefineByUnion],
	// [PublicEventAnalyticsFilterCoalescingRefineByUnion],
	// [PublicFormSubmissionFilterCoalescingRefineByUnion],
	// [PublicFormSubmissionOnPageFilterCoalescingRefineByUnion],
	// [PublicNumAssociationsFilterCoalescingRefineByUnion],
	// [PublicUnifiedEventsFilterCoalescingRefineByUnion],
	// [PublicPropertyAssociationInListFilterCoalescingRefineByUnion].
	SetType    string `json:"setType"`
	Comparison string `json:"comparison"`
	// This field is from variant
	// [PublicAssociationInListFilterCoalescingRefineByUnion],
	// [PublicPageViewAnalyticsFilterCoalescingRefineByUnion],
	// [PublicCtaAnalyticsFilterCoalescingRefineByUnion],
	// [PublicEventAnalyticsFilterCoalescingRefineByUnion],
	// [PublicFormSubmissionFilterCoalescingRefineByUnion],
	// [PublicFormSubmissionOnPageFilterCoalescingRefineByUnion],
	// [PublicNumAssociationsFilterCoalescingRefineByUnion],
	// [PublicUnifiedEventsFilterCoalescingRefineByUnion],
	// [PublicPropertyAssociationInListFilterCoalescingRefineByUnion].
	TimeOffset PublicTimeOffset `json:"timeOffset"`
	// This field is from variant
	// [PublicAssociationInListFilterCoalescingRefineByUnion],
	// [PublicPageViewAnalyticsFilterCoalescingRefineByUnion],
	// [PublicCtaAnalyticsFilterCoalescingRefineByUnion],
	// [PublicEventAnalyticsFilterCoalescingRefineByUnion],
	// [PublicFormSubmissionFilterCoalescingRefineByUnion],
	// [PublicFormSubmissionOnPageFilterCoalescingRefineByUnion],
	// [PublicNumAssociationsFilterCoalescingRefineByUnion],
	// [PublicUnifiedEventsFilterCoalescingRefineByUnion],
	// [PublicPropertyAssociationInListFilterCoalescingRefineByUnion].
	LowerBoundOffset PublicTimeOffset `json:"lowerBoundOffset"`
	RangeType        string           `json:"rangeType"`
	// This field is from variant
	// [PublicAssociationInListFilterCoalescingRefineByUnion],
	// [PublicPageViewAnalyticsFilterCoalescingRefineByUnion],
	// [PublicCtaAnalyticsFilterCoalescingRefineByUnion],
	// [PublicEventAnalyticsFilterCoalescingRefineByUnion],
	// [PublicFormSubmissionFilterCoalescingRefineByUnion],
	// [PublicFormSubmissionOnPageFilterCoalescingRefineByUnion],
	// [PublicNumAssociationsFilterCoalescingRefineByUnion],
	// [PublicUnifiedEventsFilterCoalescingRefineByUnion],
	// [PublicPropertyAssociationInListFilterCoalescingRefineByUnion].
	UpperBoundOffset PublicTimeOffset `json:"upperBoundOffset"`
	// This field is from variant
	// [PublicAssociationInListFilterCoalescingRefineByUnion],
	// [PublicPageViewAnalyticsFilterCoalescingRefineByUnion],
	// [PublicCtaAnalyticsFilterCoalescingRefineByUnion],
	// [PublicEventAnalyticsFilterCoalescingRefineByUnion],
	// [PublicFormSubmissionFilterCoalescingRefineByUnion],
	// [PublicFormSubmissionOnPageFilterCoalescingRefineByUnion],
	// [PublicNumAssociationsFilterCoalescingRefineByUnion],
	// [PublicUnifiedEventsFilterCoalescingRefineByUnion],
	// [PublicPropertyAssociationInListFilterCoalescingRefineByUnion].
	Timestamp int64 `json:"timestamp"`
	// This field is from variant
	// [PublicAssociationInListFilterCoalescingRefineByUnion],
	// [PublicPageViewAnalyticsFilterCoalescingRefineByUnion],
	// [PublicCtaAnalyticsFilterCoalescingRefineByUnion],
	// [PublicEventAnalyticsFilterCoalescingRefineByUnion],
	// [PublicFormSubmissionFilterCoalescingRefineByUnion],
	// [PublicFormSubmissionOnPageFilterCoalescingRefineByUnion],
	// [PublicNumAssociationsFilterCoalescingRefineByUnion],
	// [PublicUnifiedEventsFilterCoalescingRefineByUnion],
	// [PublicPropertyAssociationInListFilterCoalescingRefineByUnion].
	LowerTimestamp int64 `json:"lowerTimestamp"`
	// This field is from variant
	// [PublicAssociationInListFilterCoalescingRefineByUnion],
	// [PublicPageViewAnalyticsFilterCoalescingRefineByUnion],
	// [PublicCtaAnalyticsFilterCoalescingRefineByUnion],
	// [PublicEventAnalyticsFilterCoalescingRefineByUnion],
	// [PublicFormSubmissionFilterCoalescingRefineByUnion],
	// [PublicFormSubmissionOnPageFilterCoalescingRefineByUnion],
	// [PublicNumAssociationsFilterCoalescingRefineByUnion],
	// [PublicUnifiedEventsFilterCoalescingRefineByUnion],
	// [PublicPropertyAssociationInListFilterCoalescingRefineByUnion].
	UpperTimestamp               int64  `json:"upperTimestamp"`
	IncludeObjectsWithNoValueSet bool   `json:"includeObjectsWithNoValueSet"`
	OperationType                string `json:"operationType"`
	Operator                     string `json:"operator"`
	// This field is from variant
	// [PublicAssociationInListFilterCoalescingRefineByUnion],
	// [PublicPageViewAnalyticsFilterCoalescingRefineByUnion],
	// [PublicCtaAnalyticsFilterCoalescingRefineByUnion],
	// [PublicEventAnalyticsFilterCoalescingRefineByUnion],
	// [PublicFormSubmissionFilterCoalescingRefineByUnion],
	// [PublicFormSubmissionOnPageFilterCoalescingRefineByUnion],
	// [PublicNumAssociationsFilterCoalescingRefineByUnion],
	// [PublicUnifiedEventsFilterCoalescingRefineByUnion],
	// [PublicPropertyAssociationInListFilterCoalescingRefineByUnion].
	TimePoint PublicTimePointOperationTimePointUnion `json:"timePoint"`
	// This field is from variant
	// [PublicAssociationInListFilterCoalescingRefineByUnion],
	// [PublicPageViewAnalyticsFilterCoalescingRefineByUnion],
	// [PublicCtaAnalyticsFilterCoalescingRefineByUnion],
	// [PublicEventAnalyticsFilterCoalescingRefineByUnion],
	// [PublicFormSubmissionFilterCoalescingRefineByUnion],
	// [PublicFormSubmissionOnPageFilterCoalescingRefineByUnion],
	// [PublicNumAssociationsFilterCoalescingRefineByUnion],
	// [PublicUnifiedEventsFilterCoalescingRefineByUnion],
	// [PublicPropertyAssociationInListFilterCoalescingRefineByUnion].
	EndpointBehavior string `json:"endpointBehavior"`
	PropertyParser   string `json:"propertyParser"`
	// This field is from variant
	// [PublicAssociationInListFilterCoalescingRefineByUnion],
	// [PublicPageViewAnalyticsFilterCoalescingRefineByUnion],
	// [PublicCtaAnalyticsFilterCoalescingRefineByUnion],
	// [PublicEventAnalyticsFilterCoalescingRefineByUnion],
	// [PublicFormSubmissionFilterCoalescingRefineByUnion],
	// [PublicFormSubmissionOnPageFilterCoalescingRefineByUnion],
	// [PublicNumAssociationsFilterCoalescingRefineByUnion],
	// [PublicUnifiedEventsFilterCoalescingRefineByUnion],
	// [PublicPropertyAssociationInListFilterCoalescingRefineByUnion].
	LowerBoundTimePoint PublicRangedTimeOperationLowerBoundTimePointUnion `json:"lowerBoundTimePoint"`
	// This field is from variant
	// [PublicAssociationInListFilterCoalescingRefineByUnion],
	// [PublicPageViewAnalyticsFilterCoalescingRefineByUnion],
	// [PublicCtaAnalyticsFilterCoalescingRefineByUnion],
	// [PublicEventAnalyticsFilterCoalescingRefineByUnion],
	// [PublicFormSubmissionFilterCoalescingRefineByUnion],
	// [PublicFormSubmissionOnPageFilterCoalescingRefineByUnion],
	// [PublicNumAssociationsFilterCoalescingRefineByUnion],
	// [PublicUnifiedEventsFilterCoalescingRefineByUnion],
	// [PublicPropertyAssociationInListFilterCoalescingRefineByUnion].
	UpperBoundTimePoint PublicRangedTimeOperationUpperBoundTimePointUnion `json:"upperBoundTimePoint"`
	// This field is from variant
	// [PublicAssociationInListFilterCoalescingRefineByUnion],
	// [PublicPageViewAnalyticsFilterCoalescingRefineByUnion],
	// [PublicCtaAnalyticsFilterCoalescingRefineByUnion],
	// [PublicEventAnalyticsFilterCoalescingRefineByUnion],
	// [PublicFormSubmissionFilterCoalescingRefineByUnion],
	// [PublicFormSubmissionOnPageFilterCoalescingRefineByUnion],
	// [PublicNumAssociationsFilterCoalescingRefineByUnion],
	// [PublicUnifiedEventsFilterCoalescingRefineByUnion],
	// [PublicPropertyAssociationInListFilterCoalescingRefineByUnion].
	LowerBoundEndpointBehavior string `json:"lowerBoundEndpointBehavior"`
	// This field is from variant
	// [PublicAssociationInListFilterCoalescingRefineByUnion],
	// [PublicPageViewAnalyticsFilterCoalescingRefineByUnion],
	// [PublicCtaAnalyticsFilterCoalescingRefineByUnion],
	// [PublicEventAnalyticsFilterCoalescingRefineByUnion],
	// [PublicFormSubmissionFilterCoalescingRefineByUnion],
	// [PublicFormSubmissionOnPageFilterCoalescingRefineByUnion],
	// [PublicNumAssociationsFilterCoalescingRefineByUnion],
	// [PublicUnifiedEventsFilterCoalescingRefineByUnion],
	// [PublicPropertyAssociationInListFilterCoalescingRefineByUnion].
	UpperBoundEndpointBehavior string `json:"upperBoundEndpointBehavior"`
	JSON                       struct {
		Type                         respjson.Field
		MaxOccurrences               respjson.Field
		MinOccurrences               respjson.Field
		SetType                      respjson.Field
		Comparison                   respjson.Field
		TimeOffset                   respjson.Field
		LowerBoundOffset             respjson.Field
		RangeType                    respjson.Field
		UpperBoundOffset             respjson.Field
		Timestamp                    respjson.Field
		LowerTimestamp               respjson.Field
		UpperTimestamp               respjson.Field
		IncludeObjectsWithNoValueSet respjson.Field
		OperationType                respjson.Field
		Operator                     respjson.Field
		TimePoint                    respjson.Field
		EndpointBehavior             respjson.Field
		PropertyParser               respjson.Field
		LowerBoundTimePoint          respjson.Field
		UpperBoundTimePoint          respjson.Field
		LowerBoundEndpointBehavior   respjson.Field
		UpperBoundEndpointBehavior   respjson.Field
		raw                          string
	} `json:"-"`
}

func (r *PublicUnifiedEventsFilterBranchFilterUnionCoalescingRefineBy) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PublicUnifiedEventsFilterBranchFilterUnionPruningRefineBy is an implicit
// subunion of [PublicUnifiedEventsFilterBranchFilterUnion].
// PublicUnifiedEventsFilterBranchFilterUnionPruningRefineBy provides convenient
// access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [PublicUnifiedEventsFilterBranchFilterUnion].
type PublicUnifiedEventsFilterBranchFilterUnionPruningRefineBy struct {
	Type string `json:"type"`
	// This field is from variant [PublicPageViewAnalyticsFilterPruningRefineByUnion],
	// [PublicCtaAnalyticsFilterPruningRefineByUnion],
	// [PublicEventAnalyticsFilterPruningRefineByUnion],
	// [PublicFormSubmissionFilterPruningRefineByUnion],
	// [PublicFormSubmissionOnPageFilterPruningRefineByUnion],
	// [PublicEmailEventFilterPruningRefineByUnion],
	// [PublicAdsTimeFilterPruningRefineByUnion],
	// [PublicUnifiedEventsFilterPruningRefineByUnion].
	MaxOccurrences int64 `json:"maxOccurrences"`
	// This field is from variant [PublicPageViewAnalyticsFilterPruningRefineByUnion],
	// [PublicCtaAnalyticsFilterPruningRefineByUnion],
	// [PublicEventAnalyticsFilterPruningRefineByUnion],
	// [PublicFormSubmissionFilterPruningRefineByUnion],
	// [PublicFormSubmissionOnPageFilterPruningRefineByUnion],
	// [PublicEmailEventFilterPruningRefineByUnion],
	// [PublicAdsTimeFilterPruningRefineByUnion],
	// [PublicUnifiedEventsFilterPruningRefineByUnion].
	MinOccurrences int64 `json:"minOccurrences"`
	// This field is from variant [PublicPageViewAnalyticsFilterPruningRefineByUnion],
	// [PublicCtaAnalyticsFilterPruningRefineByUnion],
	// [PublicEventAnalyticsFilterPruningRefineByUnion],
	// [PublicFormSubmissionFilterPruningRefineByUnion],
	// [PublicFormSubmissionOnPageFilterPruningRefineByUnion],
	// [PublicEmailEventFilterPruningRefineByUnion],
	// [PublicAdsTimeFilterPruningRefineByUnion],
	// [PublicUnifiedEventsFilterPruningRefineByUnion].
	SetType    string `json:"setType"`
	Comparison string `json:"comparison"`
	// This field is from variant [PublicPageViewAnalyticsFilterPruningRefineByUnion],
	// [PublicCtaAnalyticsFilterPruningRefineByUnion],
	// [PublicEventAnalyticsFilterPruningRefineByUnion],
	// [PublicFormSubmissionFilterPruningRefineByUnion],
	// [PublicFormSubmissionOnPageFilterPruningRefineByUnion],
	// [PublicEmailEventFilterPruningRefineByUnion],
	// [PublicAdsTimeFilterPruningRefineByUnion],
	// [PublicUnifiedEventsFilterPruningRefineByUnion].
	TimeOffset PublicTimeOffset `json:"timeOffset"`
	// This field is from variant [PublicPageViewAnalyticsFilterPruningRefineByUnion],
	// [PublicCtaAnalyticsFilterPruningRefineByUnion],
	// [PublicEventAnalyticsFilterPruningRefineByUnion],
	// [PublicFormSubmissionFilterPruningRefineByUnion],
	// [PublicFormSubmissionOnPageFilterPruningRefineByUnion],
	// [PublicEmailEventFilterPruningRefineByUnion],
	// [PublicAdsTimeFilterPruningRefineByUnion],
	// [PublicUnifiedEventsFilterPruningRefineByUnion].
	LowerBoundOffset PublicTimeOffset `json:"lowerBoundOffset"`
	RangeType        string           `json:"rangeType"`
	// This field is from variant [PublicPageViewAnalyticsFilterPruningRefineByUnion],
	// [PublicCtaAnalyticsFilterPruningRefineByUnion],
	// [PublicEventAnalyticsFilterPruningRefineByUnion],
	// [PublicFormSubmissionFilterPruningRefineByUnion],
	// [PublicFormSubmissionOnPageFilterPruningRefineByUnion],
	// [PublicEmailEventFilterPruningRefineByUnion],
	// [PublicAdsTimeFilterPruningRefineByUnion],
	// [PublicUnifiedEventsFilterPruningRefineByUnion].
	UpperBoundOffset PublicTimeOffset `json:"upperBoundOffset"`
	// This field is from variant [PublicPageViewAnalyticsFilterPruningRefineByUnion],
	// [PublicCtaAnalyticsFilterPruningRefineByUnion],
	// [PublicEventAnalyticsFilterPruningRefineByUnion],
	// [PublicFormSubmissionFilterPruningRefineByUnion],
	// [PublicFormSubmissionOnPageFilterPruningRefineByUnion],
	// [PublicEmailEventFilterPruningRefineByUnion],
	// [PublicAdsTimeFilterPruningRefineByUnion],
	// [PublicUnifiedEventsFilterPruningRefineByUnion].
	Timestamp int64 `json:"timestamp"`
	// This field is from variant [PublicPageViewAnalyticsFilterPruningRefineByUnion],
	// [PublicCtaAnalyticsFilterPruningRefineByUnion],
	// [PublicEventAnalyticsFilterPruningRefineByUnion],
	// [PublicFormSubmissionFilterPruningRefineByUnion],
	// [PublicFormSubmissionOnPageFilterPruningRefineByUnion],
	// [PublicEmailEventFilterPruningRefineByUnion],
	// [PublicAdsTimeFilterPruningRefineByUnion],
	// [PublicUnifiedEventsFilterPruningRefineByUnion].
	LowerTimestamp int64 `json:"lowerTimestamp"`
	// This field is from variant [PublicPageViewAnalyticsFilterPruningRefineByUnion],
	// [PublicCtaAnalyticsFilterPruningRefineByUnion],
	// [PublicEventAnalyticsFilterPruningRefineByUnion],
	// [PublicFormSubmissionFilterPruningRefineByUnion],
	// [PublicFormSubmissionOnPageFilterPruningRefineByUnion],
	// [PublicEmailEventFilterPruningRefineByUnion],
	// [PublicAdsTimeFilterPruningRefineByUnion],
	// [PublicUnifiedEventsFilterPruningRefineByUnion].
	UpperTimestamp               int64  `json:"upperTimestamp"`
	IncludeObjectsWithNoValueSet bool   `json:"includeObjectsWithNoValueSet"`
	OperationType                string `json:"operationType"`
	Operator                     string `json:"operator"`
	// This field is from variant [PublicPageViewAnalyticsFilterPruningRefineByUnion],
	// [PublicCtaAnalyticsFilterPruningRefineByUnion],
	// [PublicEventAnalyticsFilterPruningRefineByUnion],
	// [PublicFormSubmissionFilterPruningRefineByUnion],
	// [PublicFormSubmissionOnPageFilterPruningRefineByUnion],
	// [PublicEmailEventFilterPruningRefineByUnion],
	// [PublicAdsTimeFilterPruningRefineByUnion],
	// [PublicUnifiedEventsFilterPruningRefineByUnion].
	TimePoint PublicTimePointOperationTimePointUnion `json:"timePoint"`
	// This field is from variant [PublicPageViewAnalyticsFilterPruningRefineByUnion],
	// [PublicCtaAnalyticsFilterPruningRefineByUnion],
	// [PublicEventAnalyticsFilterPruningRefineByUnion],
	// [PublicFormSubmissionFilterPruningRefineByUnion],
	// [PublicFormSubmissionOnPageFilterPruningRefineByUnion],
	// [PublicEmailEventFilterPruningRefineByUnion],
	// [PublicAdsTimeFilterPruningRefineByUnion],
	// [PublicUnifiedEventsFilterPruningRefineByUnion].
	EndpointBehavior string `json:"endpointBehavior"`
	PropertyParser   string `json:"propertyParser"`
	// This field is from variant [PublicPageViewAnalyticsFilterPruningRefineByUnion],
	// [PublicCtaAnalyticsFilterPruningRefineByUnion],
	// [PublicEventAnalyticsFilterPruningRefineByUnion],
	// [PublicFormSubmissionFilterPruningRefineByUnion],
	// [PublicFormSubmissionOnPageFilterPruningRefineByUnion],
	// [PublicEmailEventFilterPruningRefineByUnion],
	// [PublicAdsTimeFilterPruningRefineByUnion],
	// [PublicUnifiedEventsFilterPruningRefineByUnion].
	LowerBoundTimePoint PublicRangedTimeOperationLowerBoundTimePointUnion `json:"lowerBoundTimePoint"`
	// This field is from variant [PublicPageViewAnalyticsFilterPruningRefineByUnion],
	// [PublicCtaAnalyticsFilterPruningRefineByUnion],
	// [PublicEventAnalyticsFilterPruningRefineByUnion],
	// [PublicFormSubmissionFilterPruningRefineByUnion],
	// [PublicFormSubmissionOnPageFilterPruningRefineByUnion],
	// [PublicEmailEventFilterPruningRefineByUnion],
	// [PublicAdsTimeFilterPruningRefineByUnion],
	// [PublicUnifiedEventsFilterPruningRefineByUnion].
	UpperBoundTimePoint PublicRangedTimeOperationUpperBoundTimePointUnion `json:"upperBoundTimePoint"`
	// This field is from variant [PublicPageViewAnalyticsFilterPruningRefineByUnion],
	// [PublicCtaAnalyticsFilterPruningRefineByUnion],
	// [PublicEventAnalyticsFilterPruningRefineByUnion],
	// [PublicFormSubmissionFilterPruningRefineByUnion],
	// [PublicFormSubmissionOnPageFilterPruningRefineByUnion],
	// [PublicEmailEventFilterPruningRefineByUnion],
	// [PublicAdsTimeFilterPruningRefineByUnion],
	// [PublicUnifiedEventsFilterPruningRefineByUnion].
	LowerBoundEndpointBehavior string `json:"lowerBoundEndpointBehavior"`
	// This field is from variant [PublicPageViewAnalyticsFilterPruningRefineByUnion],
	// [PublicCtaAnalyticsFilterPruningRefineByUnion],
	// [PublicEventAnalyticsFilterPruningRefineByUnion],
	// [PublicFormSubmissionFilterPruningRefineByUnion],
	// [PublicFormSubmissionOnPageFilterPruningRefineByUnion],
	// [PublicEmailEventFilterPruningRefineByUnion],
	// [PublicAdsTimeFilterPruningRefineByUnion],
	// [PublicUnifiedEventsFilterPruningRefineByUnion].
	UpperBoundEndpointBehavior string `json:"upperBoundEndpointBehavior"`
	JSON                       struct {
		Type                         respjson.Field
		MaxOccurrences               respjson.Field
		MinOccurrences               respjson.Field
		SetType                      respjson.Field
		Comparison                   respjson.Field
		TimeOffset                   respjson.Field
		LowerBoundOffset             respjson.Field
		RangeType                    respjson.Field
		UpperBoundOffset             respjson.Field
		Timestamp                    respjson.Field
		LowerTimestamp               respjson.Field
		UpperTimestamp               respjson.Field
		IncludeObjectsWithNoValueSet respjson.Field
		OperationType                respjson.Field
		Operator                     respjson.Field
		TimePoint                    respjson.Field
		EndpointBehavior             respjson.Field
		PropertyParser               respjson.Field
		LowerBoundTimePoint          respjson.Field
		UpperBoundTimePoint          respjson.Field
		LowerBoundEndpointBehavior   respjson.Field
		UpperBoundEndpointBehavior   respjson.Field
		raw                          string
	} `json:"-"`
}

func (r *PublicUnifiedEventsFilterBranchFilterUnionPruningRefineBy) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PublicUnifiedEventsFilterBranchFilterUnionEventTypeID is an implicit subunion of
// [PublicUnifiedEventsFilterBranchFilterUnion].
// PublicUnifiedEventsFilterBranchFilterUnionEventTypeID provides convenient access
// to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [PublicUnifiedEventsFilterBranchFilterUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfInt OfString]
type PublicUnifiedEventsFilterBranchFilterUnionEventTypeID struct {
	// This field will be present if the value is a [int64] instead of an object.
	OfInt int64 `json:",inline"`
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	JSON     struct {
		OfInt    respjson.Field
		OfString respjson.Field
		raw      string
	} `json:"-"`
}

func (r *PublicUnifiedEventsFilterBranchFilterUnionEventTypeID) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Defines the operation to be applied within the filter branch (HAS_COMPLETED,
// HAS_NOT_COMPLETED).
type PublicUnifiedEventsFilterBranchOperator string

const (
	PublicUnifiedEventsFilterBranchOperatorHasCompleted    PublicUnifiedEventsFilterBranchOperator = "HAS_COMPLETED"
	PublicUnifiedEventsFilterBranchOperatorHasNotCompleted PublicUnifiedEventsFilterBranchOperator = "HAS_NOT_COMPLETED"
)

// PublicUnifiedEventsFilterBranchCoalescingRefineByUnion contains all possible
// properties and values from [PublicNumOccurrencesRefineBy],
// [PublicSetOccurrencesRefineBy], [PublicRelativeComparativeTimestampRefineBy],
// [PublicRelativeRangedTimestampRefineBy],
// [PublicAbsoluteComparativeTimestampRefineBy],
// [PublicAbsoluteRangedTimestampRefineBy], [PublicAllHistoryRefineBy],
// [PublicTimePointOperation], [PublicRangedTimeOperation].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type PublicUnifiedEventsFilterBranchCoalescingRefineByUnion struct {
	Type string `json:"type"`
	// This field is from variant [PublicNumOccurrencesRefineBy].
	MaxOccurrences int64 `json:"maxOccurrences"`
	// This field is from variant [PublicNumOccurrencesRefineBy].
	MinOccurrences int64 `json:"minOccurrences"`
	// This field is from variant [PublicSetOccurrencesRefineBy].
	SetType    string `json:"setType"`
	Comparison string `json:"comparison"`
	// This field is from variant [PublicRelativeComparativeTimestampRefineBy].
	TimeOffset PublicTimeOffset `json:"timeOffset"`
	// This field is from variant [PublicRelativeRangedTimestampRefineBy].
	LowerBoundOffset PublicTimeOffset `json:"lowerBoundOffset"`
	RangeType        string           `json:"rangeType"`
	// This field is from variant [PublicRelativeRangedTimestampRefineBy].
	UpperBoundOffset PublicTimeOffset `json:"upperBoundOffset"`
	// This field is from variant [PublicAbsoluteComparativeTimestampRefineBy].
	Timestamp int64 `json:"timestamp"`
	// This field is from variant [PublicAbsoluteRangedTimestampRefineBy].
	LowerTimestamp int64 `json:"lowerTimestamp"`
	// This field is from variant [PublicAbsoluteRangedTimestampRefineBy].
	UpperTimestamp               int64  `json:"upperTimestamp"`
	IncludeObjectsWithNoValueSet bool   `json:"includeObjectsWithNoValueSet"`
	OperationType                string `json:"operationType"`
	Operator                     string `json:"operator"`
	// This field is from variant [PublicTimePointOperation].
	TimePoint PublicTimePointOperationTimePointUnion `json:"timePoint"`
	// This field is from variant [PublicTimePointOperation].
	EndpointBehavior string `json:"endpointBehavior"`
	PropertyParser   string `json:"propertyParser"`
	// This field is from variant [PublicRangedTimeOperation].
	LowerBoundTimePoint PublicRangedTimeOperationLowerBoundTimePointUnion `json:"lowerBoundTimePoint"`
	// This field is from variant [PublicRangedTimeOperation].
	UpperBoundTimePoint PublicRangedTimeOperationUpperBoundTimePointUnion `json:"upperBoundTimePoint"`
	// This field is from variant [PublicRangedTimeOperation].
	LowerBoundEndpointBehavior string `json:"lowerBoundEndpointBehavior"`
	// This field is from variant [PublicRangedTimeOperation].
	UpperBoundEndpointBehavior string `json:"upperBoundEndpointBehavior"`
	JSON                       struct {
		Type                         respjson.Field
		MaxOccurrences               respjson.Field
		MinOccurrences               respjson.Field
		SetType                      respjson.Field
		Comparison                   respjson.Field
		TimeOffset                   respjson.Field
		LowerBoundOffset             respjson.Field
		RangeType                    respjson.Field
		UpperBoundOffset             respjson.Field
		Timestamp                    respjson.Field
		LowerTimestamp               respjson.Field
		UpperTimestamp               respjson.Field
		IncludeObjectsWithNoValueSet respjson.Field
		OperationType                respjson.Field
		Operator                     respjson.Field
		TimePoint                    respjson.Field
		EndpointBehavior             respjson.Field
		PropertyParser               respjson.Field
		LowerBoundTimePoint          respjson.Field
		UpperBoundTimePoint          respjson.Field
		LowerBoundEndpointBehavior   respjson.Field
		UpperBoundEndpointBehavior   respjson.Field
		raw                          string
	} `json:"-"`
}

func (u PublicUnifiedEventsFilterBranchCoalescingRefineByUnion) AsNumOccurrences() (v PublicNumOccurrencesRefineBy) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicUnifiedEventsFilterBranchCoalescingRefineByUnion) AsSetOccurrences() (v PublicSetOccurrencesRefineBy) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicUnifiedEventsFilterBranchCoalescingRefineByUnion) AsRelativeComparative() (v PublicRelativeComparativeTimestampRefineBy) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicUnifiedEventsFilterBranchCoalescingRefineByUnion) AsRelativeRanged() (v PublicRelativeRangedTimestampRefineBy) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicUnifiedEventsFilterBranchCoalescingRefineByUnion) AsAbsoluteComparative() (v PublicAbsoluteComparativeTimestampRefineBy) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicUnifiedEventsFilterBranchCoalescingRefineByUnion) AsAbsoluteRanged() (v PublicAbsoluteRangedTimestampRefineBy) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicUnifiedEventsFilterBranchCoalescingRefineByUnion) AsAllHistory() (v PublicAllHistoryRefineBy) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicUnifiedEventsFilterBranchCoalescingRefineByUnion) AsTimePoint() (v PublicTimePointOperation) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicUnifiedEventsFilterBranchCoalescingRefineByUnion) AsTimeRanged() (v PublicRangedTimeOperation) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u PublicUnifiedEventsFilterBranchCoalescingRefineByUnion) RawJSON() string { return u.JSON.raw }

func (r *PublicUnifiedEventsFilterBranchCoalescingRefineByUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PublicUnifiedEventsFilterBranchPruningRefineByUnion contains all possible
// properties and values from [PublicNumOccurrencesRefineBy],
// [PublicSetOccurrencesRefineBy], [PublicRelativeComparativeTimestampRefineBy],
// [PublicRelativeRangedTimestampRefineBy],
// [PublicAbsoluteComparativeTimestampRefineBy],
// [PublicAbsoluteRangedTimestampRefineBy], [PublicAllHistoryRefineBy],
// [PublicTimePointOperation], [PublicRangedTimeOperation].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type PublicUnifiedEventsFilterBranchPruningRefineByUnion struct {
	Type string `json:"type"`
	// This field is from variant [PublicNumOccurrencesRefineBy].
	MaxOccurrences int64 `json:"maxOccurrences"`
	// This field is from variant [PublicNumOccurrencesRefineBy].
	MinOccurrences int64 `json:"minOccurrences"`
	// This field is from variant [PublicSetOccurrencesRefineBy].
	SetType    string `json:"setType"`
	Comparison string `json:"comparison"`
	// This field is from variant [PublicRelativeComparativeTimestampRefineBy].
	TimeOffset PublicTimeOffset `json:"timeOffset"`
	// This field is from variant [PublicRelativeRangedTimestampRefineBy].
	LowerBoundOffset PublicTimeOffset `json:"lowerBoundOffset"`
	RangeType        string           `json:"rangeType"`
	// This field is from variant [PublicRelativeRangedTimestampRefineBy].
	UpperBoundOffset PublicTimeOffset `json:"upperBoundOffset"`
	// This field is from variant [PublicAbsoluteComparativeTimestampRefineBy].
	Timestamp int64 `json:"timestamp"`
	// This field is from variant [PublicAbsoluteRangedTimestampRefineBy].
	LowerTimestamp int64 `json:"lowerTimestamp"`
	// This field is from variant [PublicAbsoluteRangedTimestampRefineBy].
	UpperTimestamp               int64  `json:"upperTimestamp"`
	IncludeObjectsWithNoValueSet bool   `json:"includeObjectsWithNoValueSet"`
	OperationType                string `json:"operationType"`
	Operator                     string `json:"operator"`
	// This field is from variant [PublicTimePointOperation].
	TimePoint PublicTimePointOperationTimePointUnion `json:"timePoint"`
	// This field is from variant [PublicTimePointOperation].
	EndpointBehavior string `json:"endpointBehavior"`
	PropertyParser   string `json:"propertyParser"`
	// This field is from variant [PublicRangedTimeOperation].
	LowerBoundTimePoint PublicRangedTimeOperationLowerBoundTimePointUnion `json:"lowerBoundTimePoint"`
	// This field is from variant [PublicRangedTimeOperation].
	UpperBoundTimePoint PublicRangedTimeOperationUpperBoundTimePointUnion `json:"upperBoundTimePoint"`
	// This field is from variant [PublicRangedTimeOperation].
	LowerBoundEndpointBehavior string `json:"lowerBoundEndpointBehavior"`
	// This field is from variant [PublicRangedTimeOperation].
	UpperBoundEndpointBehavior string `json:"upperBoundEndpointBehavior"`
	JSON                       struct {
		Type                         respjson.Field
		MaxOccurrences               respjson.Field
		MinOccurrences               respjson.Field
		SetType                      respjson.Field
		Comparison                   respjson.Field
		TimeOffset                   respjson.Field
		LowerBoundOffset             respjson.Field
		RangeType                    respjson.Field
		UpperBoundOffset             respjson.Field
		Timestamp                    respjson.Field
		LowerTimestamp               respjson.Field
		UpperTimestamp               respjson.Field
		IncludeObjectsWithNoValueSet respjson.Field
		OperationType                respjson.Field
		Operator                     respjson.Field
		TimePoint                    respjson.Field
		EndpointBehavior             respjson.Field
		PropertyParser               respjson.Field
		LowerBoundTimePoint          respjson.Field
		UpperBoundTimePoint          respjson.Field
		LowerBoundEndpointBehavior   respjson.Field
		UpperBoundEndpointBehavior   respjson.Field
		raw                          string
	} `json:"-"`
}

func (u PublicUnifiedEventsFilterBranchPruningRefineByUnion) AsNumOccurrences() (v PublicNumOccurrencesRefineBy) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicUnifiedEventsFilterBranchPruningRefineByUnion) AsSetOccurrences() (v PublicSetOccurrencesRefineBy) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicUnifiedEventsFilterBranchPruningRefineByUnion) AsRelativeComparative() (v PublicRelativeComparativeTimestampRefineBy) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicUnifiedEventsFilterBranchPruningRefineByUnion) AsRelativeRanged() (v PublicRelativeRangedTimestampRefineBy) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicUnifiedEventsFilterBranchPruningRefineByUnion) AsAbsoluteComparative() (v PublicAbsoluteComparativeTimestampRefineBy) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicUnifiedEventsFilterBranchPruningRefineByUnion) AsAbsoluteRanged() (v PublicAbsoluteRangedTimestampRefineBy) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicUnifiedEventsFilterBranchPruningRefineByUnion) AsAllHistory() (v PublicAllHistoryRefineBy) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicUnifiedEventsFilterBranchPruningRefineByUnion) AsTimePoint() (v PublicTimePointOperation) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicUnifiedEventsFilterBranchPruningRefineByUnion) AsTimeRanged() (v PublicRangedTimeOperation) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u PublicUnifiedEventsFilterBranchPruningRefineByUnion) RawJSON() string { return u.JSON.raw }

func (r *PublicUnifiedEventsFilterBranchPruningRefineByUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties EventTypeID, FilterBranches, FilterBranchOperator,
// FilterBranchType, Filters, Operator are required.
type PublicUnifiedEventsFilterBranchParam struct {
	// The identifier for the type of event associated with the filter branch.
	EventTypeID    string                                                  `json:"eventTypeId" api:"required"`
	FilterBranches []PublicUnifiedEventsFilterBranchFilterBranchUnionParam `json:"filterBranches,omitzero" api:"required"`
	// The logical operator used to combine filters within the branch (AND).
	FilterBranchOperator string `json:"filterBranchOperator" api:"required"`
	// The type of the filter branch (UNIFIED_EVENTS).
	//
	// Any of "UNIFIED_EVENTS".
	FilterBranchType PublicUnifiedEventsFilterBranchFilterBranchType   `json:"filterBranchType,omitzero" api:"required"`
	Filters          []PublicUnifiedEventsFilterBranchFilterUnionParam `json:"filters,omitzero" api:"required"`
	// Defines the operation to be applied within the filter branch (HAS_COMPLETED,
	// HAS_NOT_COMPLETED).
	//
	// Any of "HAS_COMPLETED", "HAS_NOT_COMPLETED".
	Operator PublicUnifiedEventsFilterBranchOperator `json:"operator,omitzero" api:"required"`
	// Specifies the criteria for refining the filter by coalescing.
	CoalescingRefineBy PublicUnifiedEventsFilterBranchCoalescingRefineByUnionParam `json:"coalescingRefineBy,omitzero"`
	PruningRefineBy    PublicUnifiedEventsFilterBranchPruningRefineByUnionParam    `json:"pruningRefineBy,omitzero"`
	paramObj
}

func (r PublicUnifiedEventsFilterBranchParam) MarshalJSON() (data []byte, err error) {
	type shadow PublicUnifiedEventsFilterBranchParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PublicUnifiedEventsFilterBranchParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type PublicUnifiedEventsFilterBranchFilterBranchUnionParam struct {
	OfOr                  *PublicOrFilterBranchParam                  `json:",omitzero,inline"`
	OfAnd                 *PublicAndFilterBranchParam                 `json:",omitzero,inline"`
	OfNotAll              *PublicNotAllFilterBranchParam              `json:",omitzero,inline"`
	OfNotAny              *PublicNotAnyFilterBranchParam              `json:",omitzero,inline"`
	OfRestricted          *PublicRestrictedFilterBranchParam          `json:",omitzero,inline"`
	OfUnifiedEvents       *PublicUnifiedEventsFilterBranchParam       `json:",omitzero,inline"`
	OfPropertyAssociation *PublicPropertyAssociationFilterBranchParam `json:",omitzero,inline"`
	OfAssociation         *PublicAssociationFilterBranchParam         `json:",omitzero,inline"`
	paramUnion
}

func (u PublicUnifiedEventsFilterBranchFilterBranchUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfOr,
		u.OfAnd,
		u.OfNotAll,
		u.OfNotAny,
		u.OfRestricted,
		u.OfUnifiedEvents,
		u.OfPropertyAssociation,
		u.OfAssociation)
}
func (u *PublicUnifiedEventsFilterBranchFilterBranchUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func init() {
	apijson.RegisterUnion[PublicUnifiedEventsFilterBranchFilterBranchUnionParam](
		"filterBranchType",
		apijson.Discriminator[PublicOrFilterBranchParam]("OR"),
		apijson.Discriminator[PublicAndFilterBranchParam]("AND"),
		apijson.Discriminator[PublicNotAllFilterBranchParam]("NOT_ALL"),
		apijson.Discriminator[PublicNotAnyFilterBranchParam]("NOT_ANY"),
		apijson.Discriminator[PublicRestrictedFilterBranchParam]("RESTRICTED"),
		apijson.Discriminator[PublicUnifiedEventsFilterBranchParam]("UNIFIED_EVENTS"),
		apijson.Discriminator[PublicPropertyAssociationFilterBranchParam]("PROPERTY_ASSOCIATION"),
		apijson.Discriminator[PublicAssociationFilterBranchParam]("ASSOCIATION"),
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type PublicUnifiedEventsFilterBranchFilterUnionParam struct {
	OfProperty                  *PublicPropertyFilterParam                  `json:",omitzero,inline"`
	OfAssociation               *PublicAssociationInListFilterParam         `json:",omitzero,inline"`
	OfPageView                  *PublicPageViewAnalyticsFilterParam         `json:",omitzero,inline"`
	OfCta                       *PublicCtaAnalyticsFilterParam              `json:",omitzero,inline"`
	OfEvent                     *PublicEventAnalyticsFilterParam            `json:",omitzero,inline"`
	OfFormSubmission            *PublicFormSubmissionFilterParam            `json:",omitzero,inline"`
	OfFormSubmissionOnPage      *PublicFormSubmissionOnPageFilterParam      `json:",omitzero,inline"`
	OfIntegrationEvent          *PublicIntegrationEventFilterParam          `json:",omitzero,inline"`
	OfEmailSubscription         *PublicEmailSubscriptionFilterParam         `json:",omitzero,inline"`
	OfCommunicationSubscription *PublicCommunicationSubscriptionFilterParam `json:",omitzero,inline"`
	OfCampaignInfluenced        *PublicCampaignInfluencedFilterParam        `json:",omitzero,inline"`
	OfSurveyMonkey              *PublicSurveyMonkeyFilterParam              `json:",omitzero,inline"`
	OfSurveyMonkeyValue         *PublicSurveyMonkeyValueFilterParam         `json:",omitzero,inline"`
	OfWebinar                   *PublicWebinarFilterParam                   `json:",omitzero,inline"`
	OfEmailEvent                *PublicEmailEventFilterParam                `json:",omitzero,inline"`
	OfPrivacy                   *PublicPrivacyAnalyticsFilterParam          `json:",omitzero,inline"`
	OfAdsSearch                 *PublicAdsSearchFilterParam                 `json:",omitzero,inline"`
	OfAdsTime                   *PublicAdsTimeFilterParam                   `json:",omitzero,inline"`
	OfInList                    *PublicInListFilterParam                    `json:",omitzero,inline"`
	OfNumAssociations           *PublicNumAssociationsFilterParam           `json:",omitzero,inline"`
	OfUnifiedEvents             *PublicUnifiedEventsFilterParam             `json:",omitzero,inline"`
	OfPropertyAssociation       *PublicPropertyAssociationInListFilterParam `json:",omitzero,inline"`
	OfConstant                  *PublicConstantFilterParam                  `json:",omitzero,inline"`
	paramUnion
}

func (u PublicUnifiedEventsFilterBranchFilterUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfProperty,
		u.OfAssociation,
		u.OfPageView,
		u.OfCta,
		u.OfEvent,
		u.OfFormSubmission,
		u.OfFormSubmissionOnPage,
		u.OfIntegrationEvent,
		u.OfEmailSubscription,
		u.OfCommunicationSubscription,
		u.OfCampaignInfluenced,
		u.OfSurveyMonkey,
		u.OfSurveyMonkeyValue,
		u.OfWebinar,
		u.OfEmailEvent,
		u.OfPrivacy,
		u.OfAdsSearch,
		u.OfAdsTime,
		u.OfInList,
		u.OfNumAssociations,
		u.OfUnifiedEvents,
		u.OfPropertyAssociation,
		u.OfConstant)
}
func (u *PublicUnifiedEventsFilterBranchFilterUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func init() {
	apijson.RegisterUnion[PublicUnifiedEventsFilterBranchFilterUnionParam](
		"filterType",
		apijson.Discriminator[PublicPropertyFilterParam]("PROPERTY"),
		apijson.Discriminator[PublicAssociationInListFilterParam]("ASSOCIATION"),
		apijson.Discriminator[PublicPageViewAnalyticsFilterParam]("PAGE_VIEW"),
		apijson.Discriminator[PublicCtaAnalyticsFilterParam]("CTA"),
		apijson.Discriminator[PublicEventAnalyticsFilterParam]("EVENT"),
		apijson.Discriminator[PublicFormSubmissionFilterParam]("FORM_SUBMISSION"),
		apijson.Discriminator[PublicFormSubmissionOnPageFilterParam]("FORM_SUBMISSION_ON_PAGE"),
		apijson.Discriminator[PublicIntegrationEventFilterParam]("INTEGRATION_EVENT"),
		apijson.Discriminator[PublicEmailSubscriptionFilterParam]("EMAIL_SUBSCRIPTION"),
		apijson.Discriminator[PublicCommunicationSubscriptionFilterParam]("COMMUNICATION_SUBSCRIPTION"),
		apijson.Discriminator[PublicCampaignInfluencedFilterParam]("CAMPAIGN_INFLUENCED"),
		apijson.Discriminator[PublicSurveyMonkeyFilterParam]("SURVEY_MONKEY"),
		apijson.Discriminator[PublicSurveyMonkeyValueFilterParam]("SURVEY_MONKEY_VALUE"),
		apijson.Discriminator[PublicWebinarFilterParam]("WEBINAR"),
		apijson.Discriminator[PublicEmailEventFilterParam]("EMAIL_EVENT"),
		apijson.Discriminator[PublicPrivacyAnalyticsFilterParam]("PRIVACY"),
		apijson.Discriminator[PublicAdsSearchFilterParam]("ADS_SEARCH"),
		apijson.Discriminator[PublicAdsTimeFilterParam]("ADS_TIME"),
		apijson.Discriminator[PublicInListFilterParam]("IN_LIST"),
		apijson.Discriminator[PublicNumAssociationsFilterParam]("NUM_ASSOCIATIONS"),
		apijson.Discriminator[PublicUnifiedEventsFilterParam]("UNIFIED_EVENTS"),
		apijson.Discriminator[PublicPropertyAssociationInListFilterParam]("PROPERTY_ASSOCIATION"),
		apijson.Discriminator[PublicConstantFilterParam]("CONSTANT"),
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type PublicUnifiedEventsFilterBranchCoalescingRefineByUnionParam struct {
	OfNumOccurrences      *PublicNumOccurrencesRefineByParam               `json:",omitzero,inline"`
	OfSetOccurrences      *PublicSetOccurrencesRefineByParam               `json:",omitzero,inline"`
	OfRelativeComparative *PublicRelativeComparativeTimestampRefineByParam `json:",omitzero,inline"`
	OfRelativeRanged      *PublicRelativeRangedTimestampRefineByParam      `json:",omitzero,inline"`
	OfAbsoluteComparative *PublicAbsoluteComparativeTimestampRefineByParam `json:",omitzero,inline"`
	OfAbsoluteRanged      *PublicAbsoluteRangedTimestampRefineByParam      `json:",omitzero,inline"`
	OfAllHistory          *PublicAllHistoryRefineByParam                   `json:",omitzero,inline"`
	OfTimePoint           *PublicTimePointOperationParam                   `json:",omitzero,inline"`
	OfTimeRanged          *PublicRangedTimeOperationParam                  `json:",omitzero,inline"`
	paramUnion
}

func (u PublicUnifiedEventsFilterBranchCoalescingRefineByUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfNumOccurrences,
		u.OfSetOccurrences,
		u.OfRelativeComparative,
		u.OfRelativeRanged,
		u.OfAbsoluteComparative,
		u.OfAbsoluteRanged,
		u.OfAllHistory,
		u.OfTimePoint,
		u.OfTimeRanged)
}
func (u *PublicUnifiedEventsFilterBranchCoalescingRefineByUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type PublicUnifiedEventsFilterBranchPruningRefineByUnionParam struct {
	OfNumOccurrences      *PublicNumOccurrencesRefineByParam               `json:",omitzero,inline"`
	OfSetOccurrences      *PublicSetOccurrencesRefineByParam               `json:",omitzero,inline"`
	OfRelativeComparative *PublicRelativeComparativeTimestampRefineByParam `json:",omitzero,inline"`
	OfRelativeRanged      *PublicRelativeRangedTimestampRefineByParam      `json:",omitzero,inline"`
	OfAbsoluteComparative *PublicAbsoluteComparativeTimestampRefineByParam `json:",omitzero,inline"`
	OfAbsoluteRanged      *PublicAbsoluteRangedTimestampRefineByParam      `json:",omitzero,inline"`
	OfAllHistory          *PublicAllHistoryRefineByParam                   `json:",omitzero,inline"`
	OfTimePoint           *PublicTimePointOperationParam                   `json:",omitzero,inline"`
	OfTimeRanged          *PublicRangedTimeOperationParam                  `json:",omitzero,inline"`
	paramUnion
}

func (u PublicUnifiedEventsFilterBranchPruningRefineByUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfNumOccurrences,
		u.OfSetOccurrences,
		u.OfRelativeComparative,
		u.OfRelativeRanged,
		u.OfAbsoluteComparative,
		u.OfAbsoluteRanged,
		u.OfAllHistory,
		u.OfTimePoint,
		u.OfTimeRanged)
}
func (u *PublicUnifiedEventsFilterBranchPruningRefineByUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

type PublicWebinarFilter struct {
	// Indicates the type of filter, (WEBINAR).
	//
	// Any of "WEBINAR".
	FilterType PublicWebinarFilterFilterType `json:"filterType" api:"required"`
	// Specifies the operation to be performed by the filter (HAS_WEBINAR_REGISTRATION,
	// NOT_HAS_WEBINAR_REGISTRATION, HAS_WEBINAR_ATTENDANCE,
	// NOT_HAS_WEBINAR_ATTENDANCE).
	Operator string `json:"operator" api:"required"`
	// The ID of the webinar associated with the filter.
	WebinarID string `json:"webinarId"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FilterType  respjson.Field
		Operator    respjson.Field
		WebinarID   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PublicWebinarFilter) RawJSON() string { return r.JSON.raw }
func (r *PublicWebinarFilter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this PublicWebinarFilter to a PublicWebinarFilterParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// PublicWebinarFilterParam.Overrides()
func (r PublicWebinarFilter) ToParam() PublicWebinarFilterParam {
	return param.Override[PublicWebinarFilterParam](json.RawMessage(r.RawJSON()))
}

// Indicates the type of filter, (WEBINAR).
type PublicWebinarFilterFilterType string

const (
	PublicWebinarFilterFilterTypeWebinar PublicWebinarFilterFilterType = "WEBINAR"
)

// The properties FilterType, Operator are required.
type PublicWebinarFilterParam struct {
	// Indicates the type of filter, (WEBINAR).
	//
	// Any of "WEBINAR".
	FilterType PublicWebinarFilterFilterType `json:"filterType,omitzero" api:"required"`
	// Specifies the operation to be performed by the filter (HAS_WEBINAR_REGISTRATION,
	// NOT_HAS_WEBINAR_REGISTRATION, HAS_WEBINAR_ATTENDANCE,
	// NOT_HAS_WEBINAR_ATTENDANCE).
	Operator string `json:"operator" api:"required"`
	// The ID of the webinar associated with the filter.
	WebinarID param.Opt[string] `json:"webinarId,omitzero"`
	paramObj
}

func (r PublicWebinarFilterParam) MarshalJSON() (data []byte, err error) {
	type shadow PublicWebinarFilterParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PublicWebinarFilterParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PublicWeekReference struct {
	// The day of the week (SUNDAY, MONDAY, TUESDAY, WEDNESDAY, THURSDAY, FRIDAY,
	// SATURDAY).
	//
	// Any of "FRIDAY", "MONDAY", "SATURDAY", "SUNDAY", "THURSDAY", "TUESDAY",
	// "WEDNESDAY".
	DayOfWeek PublicWeekReferenceDayOfWeek `json:"dayOfWeek" api:"required"`
	// Indicates the type of reference (WEEK).
	//
	// Any of "WEEK".
	ReferenceType PublicWeekReferenceReferenceType `json:"referenceType" api:"required"`
	// The hour component of the week reference.
	Hour int64 `json:"hour"`
	// The millisecond component of the week reference.
	Millisecond int64 `json:"millisecond"`
	// The minute component of the week reference.
	Minute int64 `json:"minute"`
	// The second component of the week reference.
	Second int64 `json:"second"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DayOfWeek     respjson.Field
		ReferenceType respjson.Field
		Hour          respjson.Field
		Millisecond   respjson.Field
		Minute        respjson.Field
		Second        respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PublicWeekReference) RawJSON() string { return r.JSON.raw }
func (r *PublicWeekReference) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this PublicWeekReference to a PublicWeekReferenceParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// PublicWeekReferenceParam.Overrides()
func (r PublicWeekReference) ToParam() PublicWeekReferenceParam {
	return param.Override[PublicWeekReferenceParam](json.RawMessage(r.RawJSON()))
}

// The day of the week (SUNDAY, MONDAY, TUESDAY, WEDNESDAY, THURSDAY, FRIDAY,
// SATURDAY).
type PublicWeekReferenceDayOfWeek string

const (
	PublicWeekReferenceDayOfWeekFriday    PublicWeekReferenceDayOfWeek = "FRIDAY"
	PublicWeekReferenceDayOfWeekMonday    PublicWeekReferenceDayOfWeek = "MONDAY"
	PublicWeekReferenceDayOfWeekSaturday  PublicWeekReferenceDayOfWeek = "SATURDAY"
	PublicWeekReferenceDayOfWeekSunday    PublicWeekReferenceDayOfWeek = "SUNDAY"
	PublicWeekReferenceDayOfWeekThursday  PublicWeekReferenceDayOfWeek = "THURSDAY"
	PublicWeekReferenceDayOfWeekTuesday   PublicWeekReferenceDayOfWeek = "TUESDAY"
	PublicWeekReferenceDayOfWeekWednesday PublicWeekReferenceDayOfWeek = "WEDNESDAY"
)

// Indicates the type of reference (WEEK).
type PublicWeekReferenceReferenceType string

const (
	PublicWeekReferenceReferenceTypeWeek PublicWeekReferenceReferenceType = "WEEK"
)

// The properties DayOfWeek, ReferenceType are required.
type PublicWeekReferenceParam struct {
	// The day of the week (SUNDAY, MONDAY, TUESDAY, WEDNESDAY, THURSDAY, FRIDAY,
	// SATURDAY).
	//
	// Any of "FRIDAY", "MONDAY", "SATURDAY", "SUNDAY", "THURSDAY", "TUESDAY",
	// "WEDNESDAY".
	DayOfWeek PublicWeekReferenceDayOfWeek `json:"dayOfWeek,omitzero" api:"required"`
	// Indicates the type of reference (WEEK).
	//
	// Any of "WEEK".
	ReferenceType PublicWeekReferenceReferenceType `json:"referenceType,omitzero" api:"required"`
	// The hour component of the week reference.
	Hour param.Opt[int64] `json:"hour,omitzero"`
	// The millisecond component of the week reference.
	Millisecond param.Opt[int64] `json:"millisecond,omitzero"`
	// The minute component of the week reference.
	Minute param.Opt[int64] `json:"minute,omitzero"`
	// The second component of the week reference.
	Second param.Opt[int64] `json:"second,omitzero"`
	paramObj
}

func (r PublicWeekReferenceParam) MarshalJSON() (data []byte, err error) {
	type shadow PublicWeekReferenceParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PublicWeekReferenceParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PublicYearReference struct {
	// The day component of the year reference.
	Day int64 `json:"day" api:"required"`
	// The month component of the year reference.
	Month int64 `json:"month" api:"required"`
	// Indicates the type of reference (YEAR).
	//
	// Any of "YEAR".
	ReferenceType PublicYearReferenceReferenceType `json:"referenceType" api:"required"`
	// The hour component of the year reference.
	Hour int64 `json:"hour"`
	// The millisecond component of the year reference.
	Millisecond int64 `json:"millisecond"`
	// The minute component of the year reference.
	Minute int64 `json:"minute"`
	// The second component of the year reference.
	Second int64 `json:"second"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Day           respjson.Field
		Month         respjson.Field
		ReferenceType respjson.Field
		Hour          respjson.Field
		Millisecond   respjson.Field
		Minute        respjson.Field
		Second        respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PublicYearReference) RawJSON() string { return r.JSON.raw }
func (r *PublicYearReference) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this PublicYearReference to a PublicYearReferenceParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// PublicYearReferenceParam.Overrides()
func (r PublicYearReference) ToParam() PublicYearReferenceParam {
	return param.Override[PublicYearReferenceParam](json.RawMessage(r.RawJSON()))
}

// Indicates the type of reference (YEAR).
type PublicYearReferenceReferenceType string

const (
	PublicYearReferenceReferenceTypeYear PublicYearReferenceReferenceType = "YEAR"
)

// The properties Day, Month, ReferenceType are required.
type PublicYearReferenceParam struct {
	// The day component of the year reference.
	Day int64 `json:"day" api:"required"`
	// The month component of the year reference.
	Month int64 `json:"month" api:"required"`
	// Indicates the type of reference (YEAR).
	//
	// Any of "YEAR".
	ReferenceType PublicYearReferenceReferenceType `json:"referenceType,omitzero" api:"required"`
	// The hour component of the year reference.
	Hour param.Opt[int64] `json:"hour,omitzero"`
	// The millisecond component of the year reference.
	Millisecond param.Opt[int64] `json:"millisecond,omitzero"`
	// The minute component of the year reference.
	Minute param.Opt[int64] `json:"minute,omitzero"`
	// The second component of the year reference.
	Second param.Opt[int64] `json:"second,omitzero"`
	paramObj
}

func (r PublicYearReferenceParam) MarshalJSON() (data []byte, err error) {
	type shadow PublicYearReferenceParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PublicYearReferenceParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties ObjectTypeID, RecordID are required.
type RecordIDInputParam struct {
	ObjectTypeID string `json:"objectTypeId" api:"required"`
	RecordID     string `json:"recordId" api:"required"`
	paramObj
}

func (r RecordIDInputParam) MarshalJSON() (data []byte, err error) {
	type shadow RecordIDInputParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RecordIDInputParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RecordIDWithMemberships struct {
	ObjectTypeID          string                 `json:"objectTypeId" api:"required"`
	RecordID              string                 `json:"recordId" api:"required"`
	RecordListMemberships []RecordListMembership `json:"recordListMemberships" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ObjectTypeID          respjson.Field
		RecordID              respjson.Field
		RecordListMemberships respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RecordIDWithMemberships) RawJSON() string { return r.JSON.raw }
func (r *RecordIDWithMemberships) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RecordListMembership struct {
	// The timestamp when the record was first added to the list.
	FirstAddedTimestamp time.Time `json:"firstAddedTimestamp" api:"required" format:"date-time"`
	// The timestamp when the record was last added to the list.
	LastAddedTimestamp time.Time `json:"lastAddedTimestamp" api:"required" format:"date-time"`
	// The unique identifier of the list.
	ListID string `json:"listId" api:"required"`
	// The version number of the list.
	ListVersion int64 `json:"listVersion" api:"required"`
	// Indicates whether the list is public.
	IsPublicList bool `json:"isPublicList"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FirstAddedTimestamp respjson.Field
		LastAddedTimestamp  respjson.Field
		ListID              respjson.Field
		ListVersion         respjson.Field
		IsPublicList        respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RecordListMembership) RawJSON() string { return r.JSON.raw }
func (r *RecordListMembership) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ListNewParams struct {
	ListCreateRequest ListCreateRequestParam
	paramObj
}

func (r ListNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.ListCreateRequest)
}
func (r *ListNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ListListParams struct {
	IncludeFilters param.Opt[bool] `query:"includeFilters,omitzero" json:"-"`
	ListIDs        []string        `query:"listIds,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ListListParams]'s query parameters as `url.Values`.
func (r ListListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ListAddAndRemoveMembershipsParams struct {
	MembershipChangeRequest MembershipChangeRequestParam
	paramObj
}

func (r ListAddAndRemoveMembershipsParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.MembershipChangeRequest)
}
func (r *ListAddAndRemoveMembershipsParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ListAddMembershipsParams struct {
	Body []string
	paramObj
}

func (r ListAddMembershipsParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.Body)
}
func (r *ListAddMembershipsParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ListAddMembershipsFromParams struct {
	ListID string `path:"listId" api:"required" json:"-"`
	paramObj
}

type ListBatchReadMembershipsParams struct {
	BatchInputRecordIDInput BatchInputRecordIDInputParam
	paramObj
}

func (r ListBatchReadMembershipsParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.BatchInputRecordIDInput)
}
func (r *ListBatchReadMembershipsParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ListNewFolderParams struct {
	ListFolderCreateRequest ListFolderCreateRequestParam
	paramObj
}

func (r ListNewFolderParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.ListFolderCreateRequest)
}
func (r *ListNewFolderParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ListNewIDMappingParams struct {
	Body []string
	paramObj
}

func (r ListNewIDMappingParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.Body)
}
func (r *ListNewIDMappingParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ListGetParams struct {
	IncludeFilters param.Opt[bool] `query:"includeFilters,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ListGetParams]'s query parameters as `url.Values`.
func (r ListGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ListGetByObjectTypeAndNameParams struct {
	ObjectTypeID   string          `path:"objectTypeId" api:"required" json:"-"`
	IncludeFilters param.Opt[bool] `query:"includeFilters,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ListGetByObjectTypeAndNameParams]'s query parameters as
// `url.Values`.
func (r ListGetByObjectTypeAndNameParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ListGetIDMappingParams struct {
	LegacyListID param.Opt[string] `query:"legacyListId,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ListGetIDMappingParams]'s query parameters as `url.Values`.
func (r ListGetIDMappingParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ListGetMembershipsJoinOrderParams struct {
	// The paging cursor token of the last successfully read resource will be returned
	// as the `paging.next.after` JSON property of a paged response containing more
	// results.
	After  param.Opt[string] `query:"after,omitzero" json:"-"`
	Before param.Opt[string] `query:"before,omitzero" json:"-"`
	// The maximum number of results to display per page.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ListGetMembershipsJoinOrderParams]'s query parameters as
// `url.Values`.
func (r ListGetMembershipsJoinOrderParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ListGetRecordMembershipsParams struct {
	ObjectTypeID string `path:"objectTypeId" api:"required" json:"-"`
	paramObj
}

type ListGetSizeAndEditsHistoryBetweenParams struct {
	EndDate   param.Opt[string] `query:"endDate,omitzero" json:"-"`
	StartDate param.Opt[string] `query:"startDate,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ListGetSizeAndEditsHistoryBetweenParams]'s query parameters
// as `url.Values`.
func (r ListGetSizeAndEditsHistoryBetweenParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ListListBySearchParams struct {
	ListSearchRequest ListSearchRequestParam
	paramObj
}

func (r ListListBySearchParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.ListSearchRequest)
}
func (r *ListListBySearchParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ListListFoldersParams struct {
	FolderID param.Opt[string] `query:"folderId,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ListListFoldersParams]'s query parameters as `url.Values`.
func (r ListListFoldersParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ListListMembershipsParams struct {
	// The paging cursor token of the last successfully read resource will be returned
	// as the `paging.next.after` JSON property of a paged response containing more
	// results.
	After  param.Opt[string] `query:"after,omitzero" json:"-"`
	Before param.Opt[string] `query:"before,omitzero" json:"-"`
	// The maximum number of results to display per page.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ListListMembershipsParams]'s query parameters as
// `url.Values`.
func (r ListListMembershipsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ListMoveFolderParams struct {
	FolderID string `path:"folderId" api:"required" json:"-"`
	paramObj
}

type ListMoveListParams struct {
	ListMoveRequest ListMoveRequestParam
	paramObj
}

func (r ListMoveListParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.ListMoveRequest)
}
func (r *ListMoveListParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ListRemoveMembershipsParams struct {
	Body []string
	paramObj
}

func (r ListRemoveMembershipsParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.Body)
}
func (r *ListRemoveMembershipsParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ListRenameFolderParams struct {
	NewFolderName param.Opt[string] `query:"newFolderName,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ListRenameFolderParams]'s query parameters as `url.Values`.
func (r ListRenameFolderParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ListUpdateListFiltersParams struct {
	ListFilterUpdateRequest  ListFilterUpdateRequestParam
	EnrollObjectsInWorkflows param.Opt[bool] `query:"enrollObjectsInWorkflows,omitzero" json:"-"`
	paramObj
}

func (r ListUpdateListFiltersParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.ListFilterUpdateRequest)
}
func (r *ListUpdateListFiltersParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [ListUpdateListFiltersParams]'s query parameters as
// `url.Values`.
func (r ListUpdateListFiltersParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ListUpdateListNameParams struct {
	IncludeFilters param.Opt[bool]   `query:"includeFilters,omitzero" json:"-"`
	ListName       param.Opt[string] `query:"listName,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ListUpdateListNameParams]'s query parameters as
// `url.Values`.
func (r ListUpdateListNameParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ListUpdateScheduleConversionParams struct {
	PublicListConversionTime PublicListConversionTimeUnionParam
	paramObj
}

func (r ListUpdateScheduleConversionParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.PublicListConversionTime)
}
func (r *ListUpdateScheduleConversionParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
