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

	"github.com/stainless-sdks/hubspot-sdk-go/internal/apijson"
	"github.com/stainless-sdks/hubspot-sdk-go/internal/apiquery"
	shimjson "github.com/stainless-sdks/hubspot-sdk-go/internal/encoding/json"
	"github.com/stainless-sdks/hubspot-sdk-go/internal/requestconfig"
	"github.com/stainless-sdks/hubspot-sdk-go/marketing"
	"github.com/stainless-sdks/hubspot-sdk-go/option"
	"github.com/stainless-sdks/hubspot-sdk-go/packages/param"
	"github.com/stainless-sdks/hubspot-sdk-go/packages/respjson"
	"github.com/stainless-sdks/hubspot-sdk-go/shared"
)

// ListService contains methods and other services that help with interacting with
// the Hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewListService] method instead.
type ListService struct {
	Options     []option.RequestOption
	Folders     ListFolderService
	Mapping     ListMappingService
	Memberships ListMembershipService
}

// NewListService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewListService(opts ...option.RequestOption) (r ListService) {
	r = ListService{}
	r.Options = opts
	r.Folders = NewListFolderService(opts...)
	r.Mapping = NewListMappingService(opts...)
	r.Memberships = NewListMembershipService(opts...)
	return
}

// Create a new list with the provided object list definition.
func (r *ListService) New(ctx context.Context, body ListNewParams, opts ...option.RequestOption) (res *ListCreateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "crm/v3/lists/"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Fetch multiple lists in a single request by **ILS list ID**. The response will
// include the definitions of all lists that exist for the `listIds` provided.
func (r *ListService) List(ctx context.Context, query ListListParams, opts ...option.RequestOption) (res *ListsByIDResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "crm/v3/lists/"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Delete a list by **ILS list ID**. Lists deleted through this endpoint can be
// restored up to 90-days following the delete. After 90-days, the list is purged
// and can no longer be restored.
func (r *ListService) Delete(ctx context.Context, listID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if listID == "" {
		err = errors.New("missing required listId parameter")
		return
	}
	path := fmt.Sprintf("crm/v3/lists/%s", listID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Delete an existing scheduled conversion for a list.
func (r *ListService) DeleteScheduleConversion(ctx context.Context, listID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if listID == "" {
		err = errors.New("missing required listId parameter")
		return
	}
	path := fmt.Sprintf("crm/v3/lists/%s/schedule-conversion", listID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Fetch a single list by **ILS list ID**.
func (r *ListService) Get(ctx context.Context, listID string, query ListGetParams, opts ...option.RequestOption) (res *ListFetchResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if listID == "" {
		err = errors.New("missing required listId parameter")
		return
	}
	path := fmt.Sprintf("crm/v3/lists/%s", listID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Fetch a single list by list name and object type.
func (r *ListService) GetByObjectTypeIDAndName(ctx context.Context, listName string, params ListGetByObjectTypeIDAndNameParams, opts ...option.RequestOption) (res *ListFetchResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.ObjectTypeID == "" {
		err = errors.New("missing required objectTypeId parameter")
		return
	}
	if listName == "" {
		err = errors.New("missing required listName parameter")
		return
	}
	path := fmt.Sprintf("crm/v3/lists/object-type-id/%s/name/%s", params.ObjectTypeID, listName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return
}

// Retrieve the conversion details for a list. This can be used to check for an
// upcoming conversion, or to get the details of when a list was already converted.
func (r *ListService) GetScheduleConversion(ctx context.Context, listID string, opts ...option.RequestOption) (res *PublicListConversionResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if listID == "" {
		err = errors.New("missing required listId parameter")
		return
	}
	path := fmt.Sprintf("crm/v3/lists/%s/schedule-conversion", listID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Restore a previously deleted list by **ILS list ID**. Deleted lists are eligible
// to be restored up-to 90-days after the list has been deleted.
func (r *ListService) Restore(ctx context.Context, listID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if listID == "" {
		err = errors.New("missing required listId parameter")
		return
	}
	path := fmt.Sprintf("crm/v3/lists/%s/restore", listID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, nil, nil, opts...)
	return
}

// Schedule the conversion of an active list into a static list, or update the
// already scheduled conversion. This can be scheduled for a specific date or based
// on activity.
func (r *ListService) ScheduleConversion(ctx context.Context, listID string, body ListScheduleConversionParams, opts ...option.RequestOption) (res *PublicListConversionResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if listID == "" {
		err = errors.New("missing required listId parameter")
		return
	}
	path := fmt.Sprintf("crm/v3/lists/%s/schedule-conversion", listID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Search lists by list name or page through all lists by providing an empty
// `query` value.
func (r *ListService) Search(ctx context.Context, body ListSearchParams, opts ...option.RequestOption) (res *ListSearchResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "crm/v3/lists/search"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Update the filter branch definition of a `DYNAMIC` list. Once updated, the list
// memberships will be re-evaluated and updated to match the new definition.
func (r *ListService) UpdateFilters(ctx context.Context, listID string, params ListUpdateFiltersParams, opts ...option.RequestOption) (res *ListUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if listID == "" {
		err = errors.New("missing required listId parameter")
		return
	}
	path := fmt.Sprintf("crm/v3/lists/%s/update-list-filters", listID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return
}

// Update the name of a list. The name must be globally unique relative to all
// other public lists in the portal.
func (r *ListService) UpdateName(ctx context.Context, listID string, body ListUpdateNameParams, opts ...option.RequestOption) (res *ListUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if listID == "" {
		err = errors.New("missing required listId parameter")
		return
	}
	path := fmt.Sprintf("crm/v3/lists/%s/update-list-name", listID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

type APICollectionResponseJoinTimeAndRecordID struct {
	Results []JoinTimeAndRecordID `json:"results,required"`
	// Contains information pagination of results.
	Paging marketing.Paging `json:"paging"`
	Total  int64            `json:"total"`
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

type APICollectionResponseRecordListMembershipNoPaging struct {
	Results []RecordListMembership `json:"results,required"`
	Total   int64                  `json:"total"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Results     respjson.Field
		Total       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r APICollectionResponseRecordListMembershipNoPaging) RawJSON() string { return r.JSON.raw }
func (r *APICollectionResponseRecordListMembershipNoPaging) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type JoinTimeAndRecordID struct {
	MembershipTimestamp time.Time `json:"membershipTimestamp,required" format:"date-time"`
	RecordID            string    `json:"recordId,required"`
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

// The request object used when creating a new object list.
//
// The properties Name, ObjectTypeID, ProcessingType are required.
type ListCreateRequestParam struct {
	// The name of the list, which must be globally unique across all public lists in
	// the portal.
	Name string `json:"name,required"`
	// The object type ID of the type of objects that the list will store.
	ObjectTypeID string `json:"objectTypeId,required"`
	// The processing type of the list. One of: `SNAPSHOT`, `MANUAL`, or `DYNAMIC`.
	ProcessingType string `json:"processingType,required"`
	// The ID of the folder that the list should be created in. If left blank, then the
	// list will be created in the root of the list folder structure.
	ListFolderID param.Opt[int64] `json:"listFolderId,omitzero"`
	// The list of custom properties to tie to the list. Custom property name is the
	// key, the value is the value.
	CustomProperties   map[string]string                       `json:"customProperties,omitzero"`
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
	OfOr                  *shared.PublicOrFilterBranchParam                  `json:",omitzero,inline"`
	OfAnd                 *shared.PublicAndFilterBranchParam                 `json:",omitzero,inline"`
	OfNotAll              *shared.PublicNotAllFilterBranchParam              `json:",omitzero,inline"`
	OfNotAny              *shared.PublicNotAnyFilterBranchParam              `json:",omitzero,inline"`
	OfRestricted          *shared.PublicRestrictedFilterBranchParam          `json:",omitzero,inline"`
	OfUnifiedEvents       *shared.PublicUnifiedEventsFilterBranchParam       `json:",omitzero,inline"`
	OfPropertyAssociation *shared.PublicPropertyAssociationFilterBranchParam `json:",omitzero,inline"`
	OfAssociation         *shared.PublicAssociationFilterBranchParam         `json:",omitzero,inline"`
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

func (u *ListCreateRequestFilterBranchUnionParam) asAny() any {
	if !param.IsOmitted(u.OfOr) {
		return u.OfOr
	} else if !param.IsOmitted(u.OfAnd) {
		return u.OfAnd
	} else if !param.IsOmitted(u.OfNotAll) {
		return u.OfNotAll
	} else if !param.IsOmitted(u.OfNotAny) {
		return u.OfNotAny
	} else if !param.IsOmitted(u.OfRestricted) {
		return u.OfRestricted
	} else if !param.IsOmitted(u.OfUnifiedEvents) {
		return u.OfUnifiedEvents
	} else if !param.IsOmitted(u.OfPropertyAssociation) {
		return u.OfPropertyAssociation
	} else if !param.IsOmitted(u.OfAssociation) {
		return u.OfAssociation
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ListCreateRequestFilterBranchUnionParam) GetEventTypeID() *string {
	if vt := u.OfUnifiedEvents; vt != nil {
		return &vt.EventTypeID
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ListCreateRequestFilterBranchUnionParam) GetCoalescingRefineBy() *shared.PublicUnifiedEventsFilterBranchCoalescingRefineByUnionParam {
	if vt := u.OfUnifiedEvents; vt != nil {
		return &vt.CoalescingRefineBy
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ListCreateRequestFilterBranchUnionParam) GetPropertyWithObjectID() *string {
	if vt := u.OfPropertyAssociation; vt != nil {
		return &vt.PropertyWithObjectID
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ListCreateRequestFilterBranchUnionParam) GetAssociationCategory() *string {
	if vt := u.OfAssociation; vt != nil {
		return &vt.AssociationCategory
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ListCreateRequestFilterBranchUnionParam) GetAssociationTypeID() *int64 {
	if vt := u.OfAssociation; vt != nil {
		return &vt.AssociationTypeID
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ListCreateRequestFilterBranchUnionParam) GetFilterBranchOperator() *string {
	if vt := u.OfOr; vt != nil {
		return (*string)(&vt.FilterBranchOperator)
	} else if vt := u.OfAnd; vt != nil {
		return (*string)(&vt.FilterBranchOperator)
	} else if vt := u.OfNotAll; vt != nil {
		return (*string)(&vt.FilterBranchOperator)
	} else if vt := u.OfNotAny; vt != nil {
		return (*string)(&vt.FilterBranchOperator)
	} else if vt := u.OfRestricted; vt != nil {
		return (*string)(&vt.FilterBranchOperator)
	} else if vt := u.OfUnifiedEvents; vt != nil {
		return (*string)(&vt.FilterBranchOperator)
	} else if vt := u.OfPropertyAssociation; vt != nil {
		return (*string)(&vt.FilterBranchOperator)
	} else if vt := u.OfAssociation; vt != nil {
		return (*string)(&vt.FilterBranchOperator)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ListCreateRequestFilterBranchUnionParam) GetFilterBranchType() *string {
	if vt := u.OfOr; vt != nil {
		return (*string)(&vt.FilterBranchType)
	} else if vt := u.OfAnd; vt != nil {
		return (*string)(&vt.FilterBranchType)
	} else if vt := u.OfNotAll; vt != nil {
		return (*string)(&vt.FilterBranchType)
	} else if vt := u.OfNotAny; vt != nil {
		return (*string)(&vt.FilterBranchType)
	} else if vt := u.OfRestricted; vt != nil {
		return (*string)(&vt.FilterBranchType)
	} else if vt := u.OfUnifiedEvents; vt != nil {
		return (*string)(&vt.FilterBranchType)
	} else if vt := u.OfPropertyAssociation; vt != nil {
		return (*string)(&vt.FilterBranchType)
	} else if vt := u.OfAssociation; vt != nil {
		return (*string)(&vt.FilterBranchType)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ListCreateRequestFilterBranchUnionParam) GetOperator() *string {
	if vt := u.OfUnifiedEvents; vt != nil {
		return (*string)(&vt.Operator)
	} else if vt := u.OfPropertyAssociation; vt != nil {
		return (*string)(&vt.Operator)
	} else if vt := u.OfAssociation; vt != nil {
		return (*string)(&vt.Operator)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ListCreateRequestFilterBranchUnionParam) GetObjectTypeID() *string {
	if vt := u.OfPropertyAssociation; vt != nil {
		return (*string)(&vt.ObjectTypeID)
	} else if vt := u.OfAssociation; vt != nil {
		return (*string)(&vt.ObjectTypeID)
	}
	return nil
}

// Returns a subunion which exports methods to access subproperties
//
// Or use AsAny() to get the underlying value
func (u ListCreateRequestFilterBranchUnionParam) GetFilterBranches() (res listCreateRequestFilterBranchUnionParamFilterBranches) {
	if vt := u.OfOr; vt != nil {
		res.any = &vt.FilterBranches
	} else if vt := u.OfAnd; vt != nil {
		res.any = &vt.FilterBranches
	} else if vt := u.OfNotAll; vt != nil {
		res.any = &vt.FilterBranches
	} else if vt := u.OfNotAny; vt != nil {
		res.any = &vt.FilterBranches
	} else if vt := u.OfRestricted; vt != nil {
		res.any = &vt.FilterBranches
	} else if vt := u.OfUnifiedEvents; vt != nil {
		res.any = &vt.FilterBranches
	} else if vt := u.OfPropertyAssociation; vt != nil {
		res.any = &vt.FilterBranches
	} else if vt := u.OfAssociation; vt != nil {
		res.any = &vt.FilterBranches
	}
	return
}

// Can have the runtime types
// [_[]shared.PublicOrFilterBranchFilterBranchUnionParam],
// [_[]shared.PublicAndFilterBranchFilterBranchUnionParam],
// [_[]shared.PublicNotAllFilterBranchFilterBranchUnionParam],
// [_[]shared.PublicNotAnyFilterBranchFilterBranchUnionParam],
// [_[]shared.PublicRestrictedFilterBranchFilterBranchUnionParam],
// [_[]shared.PublicUnifiedEventsFilterBranchFilterBranchUnionParam],
// [_[]shared.PublicPropertyAssociationFilterBranchFilterBranchUnionParam],
// [_[]shared.PublicAssociationFilterBranchFilterBranchUnionParam]
type listCreateRequestFilterBranchUnionParamFilterBranches struct{ any }

// Use the following switch statement to get the type of the union:
//
//	switch u.AsAny().(type) {
//	case *[]shared.PublicOrFilterBranchFilterBranchUnionParam:
//	case *[]shared.PublicAndFilterBranchFilterBranchUnionParam:
//	case *[]shared.PublicNotAllFilterBranchFilterBranchUnionParam:
//	case *[]shared.PublicNotAnyFilterBranchFilterBranchUnionParam:
//	case *[]shared.PublicRestrictedFilterBranchFilterBranchUnionParam:
//	case *[]shared.PublicUnifiedEventsFilterBranchFilterBranchUnionParam:
//	case *[]shared.PublicPropertyAssociationFilterBranchFilterBranchUnionParam:
//	case *[]shared.PublicAssociationFilterBranchFilterBranchUnionParam:
//	default:
//	    fmt.Errorf("not present")
//	}
func (u listCreateRequestFilterBranchUnionParamFilterBranches) AsAny() any { return u.any }

// Returns a subunion which exports methods to access subproperties
//
// Or use AsAny() to get the underlying value
func (u ListCreateRequestFilterBranchUnionParam) GetFilters() (res listCreateRequestFilterBranchUnionParamFilters) {
	if vt := u.OfOr; vt != nil {
		res.any = &vt.Filters
	} else if vt := u.OfAnd; vt != nil {
		res.any = &vt.Filters
	} else if vt := u.OfNotAll; vt != nil {
		res.any = &vt.Filters
	} else if vt := u.OfNotAny; vt != nil {
		res.any = &vt.Filters
	} else if vt := u.OfRestricted; vt != nil {
		res.any = &vt.Filters
	} else if vt := u.OfUnifiedEvents; vt != nil {
		res.any = &vt.Filters
	} else if vt := u.OfPropertyAssociation; vt != nil {
		res.any = &vt.Filters
	} else if vt := u.OfAssociation; vt != nil {
		res.any = &vt.Filters
	}
	return
}

// Can have the runtime types [_[]shared.PublicOrFilterBranchFilterUnionParam],
// [_[]shared.PublicAndFilterBranchFilterUnionParam],
// [_[]shared.PublicNotAllFilterBranchFilterUnionParam],
// [_[]shared.PublicNotAnyFilterBranchFilterUnionParam],
// [_[]shared.PublicRestrictedFilterBranchFilterUnionParam],
// [_[]shared.PublicUnifiedEventsFilterBranchFilterUnionParam],
// [_[]shared.PublicPropertyAssociationFilterBranchFilterUnionParam],
// [_[]shared.PublicAssociationFilterBranchFilterUnionParam]
type listCreateRequestFilterBranchUnionParamFilters struct{ any }

// Use the following switch statement to get the type of the union:
//
//	switch u.AsAny().(type) {
//	case *[]shared.PublicOrFilterBranchFilterUnionParam:
//	case *[]shared.PublicAndFilterBranchFilterUnionParam:
//	case *[]shared.PublicNotAllFilterBranchFilterUnionParam:
//	case *[]shared.PublicNotAnyFilterBranchFilterUnionParam:
//	case *[]shared.PublicRestrictedFilterBranchFilterUnionParam:
//	case *[]shared.PublicUnifiedEventsFilterBranchFilterUnionParam:
//	case *[]shared.PublicPropertyAssociationFilterBranchFilterUnionParam:
//	case *[]shared.PublicAssociationFilterBranchFilterUnionParam:
//	default:
//	    fmt.Errorf("not present")
//	}
func (u listCreateRequestFilterBranchUnionParamFilters) AsAny() any { return u.any }

// The response for a list create request.
type ListCreateResponse struct {
	// An object list definition.
	List PublicObjectList `json:"list,required"`
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

// The response for a list fetch request.
type ListFetchResponse struct {
	// An object list definition.
	List PublicObjectList `json:"list,required"`
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

// The definition of the list filter branch update request.
//
// The property FilterBranch is required.
type ListFilterUpdateRequestParam struct {
	FilterBranch ListFilterUpdateRequestFilterBranchUnionParam `json:"filterBranch,omitzero,required"`
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
	OfOr                  *shared.PublicOrFilterBranchParam                  `json:",omitzero,inline"`
	OfAnd                 *shared.PublicAndFilterBranchParam                 `json:",omitzero,inline"`
	OfNotAll              *shared.PublicNotAllFilterBranchParam              `json:",omitzero,inline"`
	OfNotAny              *shared.PublicNotAnyFilterBranchParam              `json:",omitzero,inline"`
	OfRestricted          *shared.PublicRestrictedFilterBranchParam          `json:",omitzero,inline"`
	OfUnifiedEvents       *shared.PublicUnifiedEventsFilterBranchParam       `json:",omitzero,inline"`
	OfPropertyAssociation *shared.PublicPropertyAssociationFilterBranchParam `json:",omitzero,inline"`
	OfAssociation         *shared.PublicAssociationFilterBranchParam         `json:",omitzero,inline"`
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

func (u *ListFilterUpdateRequestFilterBranchUnionParam) asAny() any {
	if !param.IsOmitted(u.OfOr) {
		return u.OfOr
	} else if !param.IsOmitted(u.OfAnd) {
		return u.OfAnd
	} else if !param.IsOmitted(u.OfNotAll) {
		return u.OfNotAll
	} else if !param.IsOmitted(u.OfNotAny) {
		return u.OfNotAny
	} else if !param.IsOmitted(u.OfRestricted) {
		return u.OfRestricted
	} else if !param.IsOmitted(u.OfUnifiedEvents) {
		return u.OfUnifiedEvents
	} else if !param.IsOmitted(u.OfPropertyAssociation) {
		return u.OfPropertyAssociation
	} else if !param.IsOmitted(u.OfAssociation) {
		return u.OfAssociation
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ListFilterUpdateRequestFilterBranchUnionParam) GetEventTypeID() *string {
	if vt := u.OfUnifiedEvents; vt != nil {
		return &vt.EventTypeID
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ListFilterUpdateRequestFilterBranchUnionParam) GetCoalescingRefineBy() *shared.PublicUnifiedEventsFilterBranchCoalescingRefineByUnionParam {
	if vt := u.OfUnifiedEvents; vt != nil {
		return &vt.CoalescingRefineBy
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ListFilterUpdateRequestFilterBranchUnionParam) GetPropertyWithObjectID() *string {
	if vt := u.OfPropertyAssociation; vt != nil {
		return &vt.PropertyWithObjectID
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ListFilterUpdateRequestFilterBranchUnionParam) GetAssociationCategory() *string {
	if vt := u.OfAssociation; vt != nil {
		return &vt.AssociationCategory
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ListFilterUpdateRequestFilterBranchUnionParam) GetAssociationTypeID() *int64 {
	if vt := u.OfAssociation; vt != nil {
		return &vt.AssociationTypeID
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ListFilterUpdateRequestFilterBranchUnionParam) GetFilterBranchOperator() *string {
	if vt := u.OfOr; vt != nil {
		return (*string)(&vt.FilterBranchOperator)
	} else if vt := u.OfAnd; vt != nil {
		return (*string)(&vt.FilterBranchOperator)
	} else if vt := u.OfNotAll; vt != nil {
		return (*string)(&vt.FilterBranchOperator)
	} else if vt := u.OfNotAny; vt != nil {
		return (*string)(&vt.FilterBranchOperator)
	} else if vt := u.OfRestricted; vt != nil {
		return (*string)(&vt.FilterBranchOperator)
	} else if vt := u.OfUnifiedEvents; vt != nil {
		return (*string)(&vt.FilterBranchOperator)
	} else if vt := u.OfPropertyAssociation; vt != nil {
		return (*string)(&vt.FilterBranchOperator)
	} else if vt := u.OfAssociation; vt != nil {
		return (*string)(&vt.FilterBranchOperator)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ListFilterUpdateRequestFilterBranchUnionParam) GetFilterBranchType() *string {
	if vt := u.OfOr; vt != nil {
		return (*string)(&vt.FilterBranchType)
	} else if vt := u.OfAnd; vt != nil {
		return (*string)(&vt.FilterBranchType)
	} else if vt := u.OfNotAll; vt != nil {
		return (*string)(&vt.FilterBranchType)
	} else if vt := u.OfNotAny; vt != nil {
		return (*string)(&vt.FilterBranchType)
	} else if vt := u.OfRestricted; vt != nil {
		return (*string)(&vt.FilterBranchType)
	} else if vt := u.OfUnifiedEvents; vt != nil {
		return (*string)(&vt.FilterBranchType)
	} else if vt := u.OfPropertyAssociation; vt != nil {
		return (*string)(&vt.FilterBranchType)
	} else if vt := u.OfAssociation; vt != nil {
		return (*string)(&vt.FilterBranchType)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ListFilterUpdateRequestFilterBranchUnionParam) GetOperator() *string {
	if vt := u.OfUnifiedEvents; vt != nil {
		return (*string)(&vt.Operator)
	} else if vt := u.OfPropertyAssociation; vt != nil {
		return (*string)(&vt.Operator)
	} else if vt := u.OfAssociation; vt != nil {
		return (*string)(&vt.Operator)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ListFilterUpdateRequestFilterBranchUnionParam) GetObjectTypeID() *string {
	if vt := u.OfPropertyAssociation; vt != nil {
		return (*string)(&vt.ObjectTypeID)
	} else if vt := u.OfAssociation; vt != nil {
		return (*string)(&vt.ObjectTypeID)
	}
	return nil
}

// Returns a subunion which exports methods to access subproperties
//
// Or use AsAny() to get the underlying value
func (u ListFilterUpdateRequestFilterBranchUnionParam) GetFilterBranches() (res listFilterUpdateRequestFilterBranchUnionParamFilterBranches) {
	if vt := u.OfOr; vt != nil {
		res.any = &vt.FilterBranches
	} else if vt := u.OfAnd; vt != nil {
		res.any = &vt.FilterBranches
	} else if vt := u.OfNotAll; vt != nil {
		res.any = &vt.FilterBranches
	} else if vt := u.OfNotAny; vt != nil {
		res.any = &vt.FilterBranches
	} else if vt := u.OfRestricted; vt != nil {
		res.any = &vt.FilterBranches
	} else if vt := u.OfUnifiedEvents; vt != nil {
		res.any = &vt.FilterBranches
	} else if vt := u.OfPropertyAssociation; vt != nil {
		res.any = &vt.FilterBranches
	} else if vt := u.OfAssociation; vt != nil {
		res.any = &vt.FilterBranches
	}
	return
}

// Can have the runtime types
// [_[]shared.PublicOrFilterBranchFilterBranchUnionParam],
// [_[]shared.PublicAndFilterBranchFilterBranchUnionParam],
// [_[]shared.PublicNotAllFilterBranchFilterBranchUnionParam],
// [_[]shared.PublicNotAnyFilterBranchFilterBranchUnionParam],
// [_[]shared.PublicRestrictedFilterBranchFilterBranchUnionParam],
// [_[]shared.PublicUnifiedEventsFilterBranchFilterBranchUnionParam],
// [_[]shared.PublicPropertyAssociationFilterBranchFilterBranchUnionParam],
// [_[]shared.PublicAssociationFilterBranchFilterBranchUnionParam]
type listFilterUpdateRequestFilterBranchUnionParamFilterBranches struct{ any }

// Use the following switch statement to get the type of the union:
//
//	switch u.AsAny().(type) {
//	case *[]shared.PublicOrFilterBranchFilterBranchUnionParam:
//	case *[]shared.PublicAndFilterBranchFilterBranchUnionParam:
//	case *[]shared.PublicNotAllFilterBranchFilterBranchUnionParam:
//	case *[]shared.PublicNotAnyFilterBranchFilterBranchUnionParam:
//	case *[]shared.PublicRestrictedFilterBranchFilterBranchUnionParam:
//	case *[]shared.PublicUnifiedEventsFilterBranchFilterBranchUnionParam:
//	case *[]shared.PublicPropertyAssociationFilterBranchFilterBranchUnionParam:
//	case *[]shared.PublicAssociationFilterBranchFilterBranchUnionParam:
//	default:
//	    fmt.Errorf("not present")
//	}
func (u listFilterUpdateRequestFilterBranchUnionParamFilterBranches) AsAny() any { return u.any }

// Returns a subunion which exports methods to access subproperties
//
// Or use AsAny() to get the underlying value
func (u ListFilterUpdateRequestFilterBranchUnionParam) GetFilters() (res listFilterUpdateRequestFilterBranchUnionParamFilters) {
	if vt := u.OfOr; vt != nil {
		res.any = &vt.Filters
	} else if vt := u.OfAnd; vt != nil {
		res.any = &vt.Filters
	} else if vt := u.OfNotAll; vt != nil {
		res.any = &vt.Filters
	} else if vt := u.OfNotAny; vt != nil {
		res.any = &vt.Filters
	} else if vt := u.OfRestricted; vt != nil {
		res.any = &vt.Filters
	} else if vt := u.OfUnifiedEvents; vt != nil {
		res.any = &vt.Filters
	} else if vt := u.OfPropertyAssociation; vt != nil {
		res.any = &vt.Filters
	} else if vt := u.OfAssociation; vt != nil {
		res.any = &vt.Filters
	}
	return
}

// Can have the runtime types [_[]shared.PublicOrFilterBranchFilterUnionParam],
// [_[]shared.PublicAndFilterBranchFilterUnionParam],
// [_[]shared.PublicNotAllFilterBranchFilterUnionParam],
// [_[]shared.PublicNotAnyFilterBranchFilterUnionParam],
// [_[]shared.PublicRestrictedFilterBranchFilterUnionParam],
// [_[]shared.PublicUnifiedEventsFilterBranchFilterUnionParam],
// [_[]shared.PublicPropertyAssociationFilterBranchFilterUnionParam],
// [_[]shared.PublicAssociationFilterBranchFilterUnionParam]
type listFilterUpdateRequestFilterBranchUnionParamFilters struct{ any }

// Use the following switch statement to get the type of the union:
//
//	switch u.AsAny().(type) {
//	case *[]shared.PublicOrFilterBranchFilterUnionParam:
//	case *[]shared.PublicAndFilterBranchFilterUnionParam:
//	case *[]shared.PublicNotAllFilterBranchFilterUnionParam:
//	case *[]shared.PublicNotAnyFilterBranchFilterUnionParam:
//	case *[]shared.PublicRestrictedFilterBranchFilterUnionParam:
//	case *[]shared.PublicUnifiedEventsFilterBranchFilterUnionParam:
//	case *[]shared.PublicPropertyAssociationFilterBranchFilterUnionParam:
//	case *[]shared.PublicAssociationFilterBranchFilterUnionParam:
//	default:
//	    fmt.Errorf("not present")
//	}
func (u listFilterUpdateRequestFilterBranchUnionParamFilters) AsAny() any { return u.any }

// The property Name is required.
type ListFolderCreateRequestParam struct {
	// The name of the folder to be created.
	Name string `json:"name,required"`
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
	Folder PublicListFolder `json:"folder,required"`
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
	Folder PublicListFolder `json:"folder,required"`
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
	ListID string `json:"listId,required"`
	// The Id of folder to move the list to, the root folder is Id 0.
	NewFolderID string `json:"newFolderId,required"`
	paramObj
}

func (r ListMoveRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow ListMoveRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ListMoveRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The response object containing the lists found for a multi-list fetch.
type ListsByIDResponse struct {
	// The object list definitions.
	Lists []PublicObjectList `json:"lists,required"`
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

// The request object used for searching through lists.
type ListSearchRequestParam struct {
	// The number of lists to include in the response. Defaults to `20` if no value is
	// provided. The max `count` is `500`.
	Count param.Opt[int64] `json:"count,omitzero"`
	// Value used to paginate through lists. The `offset` provided in the response can
	// be used in the next request to fetch the next page of results. Defaults to `0`
	// if no offset is provided.
	Offset param.Opt[int64] `json:"offset,omitzero"`
	// The `query` that will be used to search for lists by list name. If no `query` is
	// provided, then the results will include all lists.
	Query param.Opt[string] `json:"query,omitzero"`
	Sort  param.Opt[string] `json:"sort,omitzero"`
	// The property names of any additional list properties to include in the response.
	// Properties that do not exist or that are empty for a particular list are not
	// included in the response.
	//
	// By default, all requests will fetch the following properties for each list:
	// `hs_list_size`, `hs_last_record_added_at`, `hs_last_record_removed_at`,
	// `hs_folder_name`, and `hs_list_reference_count`.
	AdditionalProperties []string `json:"additionalProperties,omitzero"`
	// The `listIds` that will be used to filter results by `listId`. If values are
	// provided, then the response will only include results that have a `listId` in
	// this array.
	//
	// If no value is provided, or if an empty list is provided, then the results will
	// not be filtered by `listId`.
	ListIDs []string `json:"listIds,omitzero"`
	// The `processingTypes` that will be used to filter results by `processingType`.
	// If values are provided, then the response will only include results that have a
	// `processingType` in this array.
	//
	// If no value is provided, or if an empty list is provided, then results will not
	// be filtered by `processingType`.
	//
	// Valid `processingTypes` are: `MANUAL`, `SNAPSHOT`, or `DYNAMIC`.
	ProcessingTypes []string `json:"processingTypes,omitzero"`
	paramObj
}

func (r ListSearchRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow ListSearchRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ListSearchRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The response object with the list search hits and additional information
// regarding pagination.
type ListSearchResponse struct {
	// Whether or not there are more results to page through.
	HasMore bool `json:"hasMore,required"`
	// The lists that matched the search criteria.
	Lists []PublicObjectListSearchResult `json:"lists,required"`
	// Value to be passed in a future request to paginate through list search results.
	Offset int64 `json:"offset,required"`
	// The total number of lists that match the search criteria.
	Total int64 `json:"total,required"`
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

// The updated definition of the list in response to a list update request.
type ListUpdateResponse struct {
	// An object list definition.
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

// The IDs of the records to add and/or remove from a list.
//
// The properties RecordIDsToAdd, RecordIDsToRemove are required.
type MembershipChangeRequestParam struct {
	RecordIDsToAdd    []string `json:"recordIdsToAdd,omitzero,required"`
	RecordIDsToRemove []string `json:"recordIdsToRemove,omitzero,required"`
	paramObj
}

func (r MembershipChangeRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow MembershipChangeRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MembershipChangeRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The IDs of the records that were `added`, `removed`, and/or found to be
// `missing` as a result of the membership update request.
type MembershipsUpdateResponse struct {
	// The IDs of the records that were `missing` (e.g. did not exist in the portal)
	// and so were not `added` or `removed`.
	RecordIDsMissing []string `json:"recordIdsMissing,required"`
	// The IDs of the records that were `removed` from the list.
	RecordIDsRemoved []string `json:"recordIdsRemoved,required"`
	RecordsIDsAdded  []string `json:"recordsIdsAdded,required"`
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

type PublicBatchMigrationMapping struct {
	LegacyListIDsToIDsMapping []PublicMigrationMapping `json:"legacyListIdsToIdsMapping,required"`
	// A list of legacy list ids that were passed in but not found. It will be empty if
	// no id's are missing
	MissingLegacyListIDs []string `json:"missingLegacyListIds,required"`
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

type PublicListConversionDate struct {
	// Any of "CONVERSION_DATE".
	ConversionType PublicListConversionDateConversionType `json:"conversionType,required"`
	Day            int64                                  `json:"day,required"`
	Month          int64                                  `json:"month,required"`
	Year           int64                                  `json:"year,required"`
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

type PublicListConversionDateConversionType string

const (
	PublicListConversionDateConversionTypeConversionDate PublicListConversionDateConversionType = "CONVERSION_DATE"
)

// The properties ConversionType, Day, Month, Year are required.
type PublicListConversionDateParam struct {
	// Any of "CONVERSION_DATE".
	ConversionType PublicListConversionDateConversionType `json:"conversionType,omitzero,required"`
	Day            int64                                  `json:"day,required"`
	Month          int64                                  `json:"month,required"`
	Year           int64                                  `json:"year,required"`
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
	// Any of "INACTIVITY".
	ConversionType PublicListConversionInactivityConversionType `json:"conversionType,required"`
	Offset         int64                                        `json:"offset,required"`
	// Any of "DAY", "WEEK", "MONTH".
	TimeUnit PublicListConversionInactivityTimeUnit `json:"timeUnit,required"`
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

type PublicListConversionInactivityConversionType string

const (
	PublicListConversionInactivityConversionTypeInactivity PublicListConversionInactivityConversionType = "INACTIVITY"
)

type PublicListConversionInactivityTimeUnit string

const (
	PublicListConversionInactivityTimeUnitDay   PublicListConversionInactivityTimeUnit = "DAY"
	PublicListConversionInactivityTimeUnitWeek  PublicListConversionInactivityTimeUnit = "WEEK"
	PublicListConversionInactivityTimeUnitMonth PublicListConversionInactivityTimeUnit = "MONTH"
)

// The properties ConversionType, Offset, TimeUnit are required.
type PublicListConversionInactivityParam struct {
	// Any of "INACTIVITY".
	ConversionType PublicListConversionInactivityConversionType `json:"conversionType,omitzero,required"`
	Offset         int64                                        `json:"offset,required"`
	// Any of "DAY", "WEEK", "MONTH".
	TimeUnit PublicListConversionInactivityTimeUnit `json:"timeUnit,omitzero,required"`
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
	ListID                  string                                                   `json:"listId,required"`
	ConvertedAt             time.Time                                                `json:"convertedAt" format:"date-time"`
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

func (u *PublicListConversionTimeUnionParam) asAny() any {
	if !param.IsOmitted(u.OfConversionDate) {
		return u.OfConversionDate
	} else if !param.IsOmitted(u.OfInactivity) {
		return u.OfInactivity
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u PublicListConversionTimeUnionParam) GetDay() *int64 {
	if vt := u.OfConversionDate; vt != nil {
		return &vt.Day
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u PublicListConversionTimeUnionParam) GetMonth() *int64 {
	if vt := u.OfConversionDate; vt != nil {
		return &vt.Month
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u PublicListConversionTimeUnionParam) GetYear() *int64 {
	if vt := u.OfConversionDate; vt != nil {
		return &vt.Year
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u PublicListConversionTimeUnionParam) GetOffset() *int64 {
	if vt := u.OfInactivity; vt != nil {
		return &vt.Offset
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u PublicListConversionTimeUnionParam) GetTimeUnit() *string {
	if vt := u.OfInactivity; vt != nil {
		return (*string)(&vt.TimeUnit)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u PublicListConversionTimeUnionParam) GetConversionType() *string {
	if vt := u.OfConversionDate; vt != nil {
		return (*string)(&vt.ConversionType)
	} else if vt := u.OfInactivity; vt != nil {
		return (*string)(&vt.ConversionType)
	}
	return nil
}

type PublicListFolder struct {
	// The Id of the folder.
	ID string `json:"id,required"`
	// An array of list Id's contained in this folder.
	ChildLists []int64            `json:"childLists,required"`
	ChildNodes []PublicListFolder `json:"childNodes,required"`
	// The Id of the folder this folder is in, the root folder is represented as 0.
	ParentFolderID string `json:"parentFolderId,required"`
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
	TeamsWithEditAccess []int64 `json:"teamsWithEditAccess,required"`
	UsersWithEditAccess []int64 `json:"usersWithEditAccess,required"`
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
	TeamsWithEditAccess []int64 `json:"teamsWithEditAccess,omitzero,required"`
	UsersWithEditAccess []int64 `json:"usersWithEditAccess,omitzero,required"`
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
	IncludeUnassigned bool  `json:"includeUnassigned"`
	MembershipTeamID  int64 `json:"membershipTeamId"`
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
	IncludeUnassigned param.Opt[bool]  `json:"includeUnassigned,omitzero"`
	MembershipTeamID  param.Opt[int64] `json:"membershipTeamId,omitzero"`
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
	LegacyListID string `json:"legacyListId,required"`
	// The V3 list id for the list
	ListID string `json:"listId,required"`
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

// An object list definition.
type PublicObjectList struct {
	// The **ILS ID** of the list.
	ListID string `json:"listId,required"`
	// The version of the list.
	ListVersion int64 `json:"listVersion,required"`
	// The name of the list.
	Name string `json:"name,required"`
	// The object type of the list.
	ObjectTypeID string `json:"objectTypeId,required"`
	// The processing status of the list.
	ProcessingStatus string `json:"processingStatus,required"`
	// The processing type of the list.
	ProcessingType string `json:"processingType,required"`
	// The time when the list was created.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// The ID of the user that created the list.
	CreatedByID string `json:"createdById"`
	// The time when the list was deleted.
	DeletedAt    time.Time                         `json:"deletedAt" format:"date-time"`
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
// from [shared.PublicOrFilterBranch], [shared.PublicAndFilterBranch],
// [shared.PublicNotAllFilterBranch], [shared.PublicNotAnyFilterBranch],
// [shared.PublicRestrictedFilterBranch], [shared.PublicUnifiedEventsFilterBranch],
// [shared.PublicPropertyAssociationFilterBranch],
// [shared.PublicAssociationFilterBranch].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type PublicObjectListFilterBranchUnion struct {
	// This field is a union of [[]shared.PublicOrFilterBranchFilterBranchUnion],
	// [[]shared.PublicAndFilterBranchFilterBranchUnion],
	// [[]shared.PublicNotAllFilterBranchFilterBranchUnion],
	// [[]shared.PublicNotAnyFilterBranchFilterBranchUnion],
	// [[]shared.PublicRestrictedFilterBranchFilterBranchUnion],
	// [[]shared.PublicUnifiedEventsFilterBranchFilterBranchUnion],
	// [[]shared.PublicPropertyAssociationFilterBranchFilterBranchUnion],
	// [[]shared.PublicAssociationFilterBranchFilterBranchUnion]
	FilterBranches       PublicObjectListFilterBranchUnionFilterBranches `json:"filterBranches"`
	FilterBranchOperator string                                          `json:"filterBranchOperator"`
	FilterBranchType     string                                          `json:"filterBranchType"`
	// This field is a union of [[]shared.PublicOrFilterBranchFilterUnion],
	// [[]shared.PublicAndFilterBranchFilterUnion],
	// [[]shared.PublicNotAllFilterBranchFilterUnion],
	// [[]shared.PublicNotAnyFilterBranchFilterUnion],
	// [[]shared.PublicRestrictedFilterBranchFilterUnion],
	// [[]shared.PublicUnifiedEventsFilterBranchFilterUnion],
	// [[]shared.PublicPropertyAssociationFilterBranchFilterUnion],
	// [[]shared.PublicAssociationFilterBranchFilterUnion]
	Filters PublicObjectListFilterBranchUnionFilters `json:"filters"`
	// This field is from variant [shared.PublicUnifiedEventsFilterBranch].
	EventTypeID string `json:"eventTypeId"`
	Operator    string `json:"operator"`
	// This field is from variant [shared.PublicUnifiedEventsFilterBranch].
	CoalescingRefineBy shared.PublicUnifiedEventsFilterBranchCoalescingRefineByUnion `json:"coalescingRefineBy"`
	ObjectTypeID       string                                                        `json:"objectTypeId"`
	// This field is from variant [shared.PublicPropertyAssociationFilterBranch].
	PropertyWithObjectID string `json:"propertyWithObjectId"`
	// This field is from variant [shared.PublicAssociationFilterBranch].
	AssociationCategory string `json:"associationCategory"`
	// This field is from variant [shared.PublicAssociationFilterBranch].
	AssociationTypeID int64 `json:"associationTypeId"`
	JSON              struct {
		FilterBranches       respjson.Field
		FilterBranchOperator respjson.Field
		FilterBranchType     respjson.Field
		Filters              respjson.Field
		EventTypeID          respjson.Field
		Operator             respjson.Field
		CoalescingRefineBy   respjson.Field
		ObjectTypeID         respjson.Field
		PropertyWithObjectID respjson.Field
		AssociationCategory  respjson.Field
		AssociationTypeID    respjson.Field
		raw                  string
	} `json:"-"`
}

func (u PublicObjectListFilterBranchUnion) AsOr() (v shared.PublicOrFilterBranch) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicObjectListFilterBranchUnion) AsAnd() (v shared.PublicAndFilterBranch) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicObjectListFilterBranchUnion) AsNotAll() (v shared.PublicNotAllFilterBranch) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicObjectListFilterBranchUnion) AsNotAny() (v shared.PublicNotAnyFilterBranch) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicObjectListFilterBranchUnion) AsRestricted() (v shared.PublicRestrictedFilterBranch) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicObjectListFilterBranchUnion) AsUnifiedEvents() (v shared.PublicUnifiedEventsFilterBranch) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicObjectListFilterBranchUnion) AsPropertyAssociation() (v shared.PublicPropertyAssociationFilterBranch) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicObjectListFilterBranchUnion) AsAssociation() (v shared.PublicAssociationFilterBranch) {
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
	// [[]shared.PublicOrFilterBranchFilterBranchUnion] instead of an object.
	OfPublicOrFilterBranchFilterBranches []shared.PublicOrFilterBranchFilterBranchUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]shared.PublicAndFilterBranchFilterBranchUnion] instead of an object.
	OfPublicAndFilterBranchFilterBranches []shared.PublicAndFilterBranchFilterBranchUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]shared.PublicNotAllFilterBranchFilterBranchUnion] instead of an object.
	OfPublicNotAllFilterBranchFilterBranches []shared.PublicNotAllFilterBranchFilterBranchUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]shared.PublicNotAnyFilterBranchFilterBranchUnion] instead of an object.
	OfPublicNotAnyFilterBranchFilterBranches []shared.PublicNotAnyFilterBranchFilterBranchUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]shared.PublicRestrictedFilterBranchFilterBranchUnion] instead of an object.
	OfPublicRestrictedFilterBranchFilterBranches []shared.PublicRestrictedFilterBranchFilterBranchUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]shared.PublicUnifiedEventsFilterBranchFilterBranchUnion] instead of an
	// object.
	OfPublicUnifiedEventsFilterBranchFilterBranches []shared.PublicUnifiedEventsFilterBranchFilterBranchUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]shared.PublicPropertyAssociationFilterBranchFilterBranchUnion] instead of an
	// object.
	OfPublicPropertyAssociationFilterBranchFilterBranches []shared.PublicPropertyAssociationFilterBranchFilterBranchUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]shared.PublicAssociationFilterBranchFilterBranchUnion] instead of an object.
	OfPublicAssociationFilterBranchFilterBranches []shared.PublicAssociationFilterBranchFilterBranchUnion `json:",inline"`
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
	// This field will be present if the value is a
	// [[]shared.PublicOrFilterBranchFilterUnion] instead of an object.
	OfPublicOrFilterBranchFilters []shared.PublicOrFilterBranchFilterUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]shared.PublicAndFilterBranchFilterUnion] instead of an object.
	OfPublicAndFilterBranchFilters []shared.PublicAndFilterBranchFilterUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]shared.PublicNotAllFilterBranchFilterUnion] instead of an object.
	OfPublicNotAllFilterBranchFilters []shared.PublicNotAllFilterBranchFilterUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]shared.PublicNotAnyFilterBranchFilterUnion] instead of an object.
	OfPublicNotAnyFilterBranchFilters []shared.PublicNotAnyFilterBranchFilterUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]shared.PublicRestrictedFilterBranchFilterUnion] instead of an object.
	OfPublicRestrictedFilterBranchFilters []shared.PublicRestrictedFilterBranchFilterUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]shared.PublicUnifiedEventsFilterBranchFilterUnion] instead of an object.
	OfPublicUnifiedEventsFilterBranchFilters []shared.PublicUnifiedEventsFilterBranchFilterUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]shared.PublicPropertyAssociationFilterBranchFilterUnion] instead of an
	// object.
	OfPublicPropertyAssociationFilterBranchFilters []shared.PublicPropertyAssociationFilterBranchFilterUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]shared.PublicAssociationFilterBranchFilterUnion] instead of an object.
	OfPublicAssociationFilterBranchFilters []shared.PublicAssociationFilterBranchFilterUnion `json:",inline"`
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
	// The name and value of any additional properties that exist for this list and
	// that were included in the search request.
	AdditionalProperties map[string]string `json:"additionalProperties,required"`
	// The **ILS ID** of the list.
	ListID string `json:"listId,required"`
	// The version of the list.
	ListVersion int64 `json:"listVersion,required"`
	// The name of the list.
	Name string `json:"name,required"`
	// The object type of the list.
	ObjectTypeID string `json:"objectTypeId,required"`
	// The processing status of the list.
	ProcessingStatus string `json:"processingStatus,required"`
	// The processing type of the list.
	ProcessingType string `json:"processingType,required"`
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
		AdditionalProperties respjson.Field
		ListID               respjson.Field
		ListVersion          respjson.Field
		Name                 respjson.Field
		ObjectTypeID         respjson.Field
		ProcessingStatus     respjson.Field
		ProcessingType       respjson.Field
		CreatedAt            respjson.Field
		CreatedByID          respjson.Field
		DeletedAt            respjson.Field
		FiltersUpdatedAt     respjson.Field
		UpdatedAt            respjson.Field
		UpdatedByID          respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PublicObjectListSearchResult) RawJSON() string { return r.JSON.raw }
func (r *PublicObjectListSearchResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Lists record is member of
type RecordListMembership struct {
	FirstAddedTimestamp time.Time `json:"firstAddedTimestamp,required" format:"date-time"`
	LastAddedTimestamp  time.Time `json:"lastAddedTimestamp,required" format:"date-time"`
	ListID              string    `json:"listId,required"`
	ListVersion         int64     `json:"listVersion,required"`
	IsPublicList        bool      `json:"isPublicList"`
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
	// The request object used when creating a new object list.
	ListCreateRequest ListCreateRequestParam
	paramObj
}

func (r ListNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.ListCreateRequest)
}
func (r *ListNewParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.ListCreateRequest)
}

type ListListParams struct {
	// A flag indicating whether or not the response object list definitions should
	// include a filter branch definition. By default, object list definitions will not
	// have their filter branch definitions included in the response.
	IncludeFilters param.Opt[bool] `query:"includeFilters,omitzero" json:"-"`
	// The **ILS IDs** of the lists to fetch.
	ListIDs []string `query:"listIds,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ListListParams]'s query parameters as `url.Values`.
func (r ListListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ListGetParams struct {
	// A flag indicating whether or not the response object list definition should
	// include a filter branch definition. By default, object list definitions will not
	// have their filter branch definitions included in the response.
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

type ListGetByObjectTypeIDAndNameParams struct {
	ObjectTypeID string `path:"objectTypeId,required" json:"-"`
	// A flag indicating whether or not the response object list definition should
	// include a filter branch definition. By default, object list definitions will not
	// have their filter branch definitions included in the response.
	IncludeFilters param.Opt[bool] `query:"includeFilters,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ListGetByObjectTypeIDAndNameParams]'s query parameters as
// `url.Values`.
func (r ListGetByObjectTypeIDAndNameParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ListScheduleConversionParams struct {
	PublicListConversionTime PublicListConversionTimeUnionParam
	paramObj
}

func (r ListScheduleConversionParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.PublicListConversionTime)
}
func (r *ListScheduleConversionParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.PublicListConversionTime)
}

type ListSearchParams struct {
	// The request object used for searching through lists.
	ListSearchRequest ListSearchRequestParam
	paramObj
}

func (r ListSearchParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.ListSearchRequest)
}
func (r *ListSearchParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.ListSearchRequest)
}

type ListUpdateFiltersParams struct {
	// The definition of the list filter branch update request.
	ListFilterUpdateRequest ListFilterUpdateRequestParam
	// A flag indicating whether or not the memberships added to the list as a result
	// of the filter change should be enrolled in workflows that are relevant to this
	// list.
	EnrollObjectsInWorkflows param.Opt[bool] `query:"enrollObjectsInWorkflows,omitzero" json:"-"`
	paramObj
}

func (r ListUpdateFiltersParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.ListFilterUpdateRequest)
}
func (r *ListUpdateFiltersParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.ListFilterUpdateRequest)
}

// URLQuery serializes [ListUpdateFiltersParams]'s query parameters as
// `url.Values`.
func (r ListUpdateFiltersParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ListUpdateNameParams struct {
	// A flag indicating whether or not the response object list definition should
	// include a filter branch definition. By default, object list definitions will not
	// have their filter branch definitions included in the response.
	IncludeFilters param.Opt[bool] `query:"includeFilters,omitzero" json:"-"`
	// The name to update the list to.
	ListName param.Opt[string] `query:"listName,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ListUpdateNameParams]'s query parameters as `url.Values`.
func (r ListUpdateNameParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
