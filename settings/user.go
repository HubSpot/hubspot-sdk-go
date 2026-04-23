// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package settings

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/stainless-sdks/hubspot-sdk-go/internal/apijson"
	"github.com/stainless-sdks/hubspot-sdk-go/internal/apiquery"
	shimjson "github.com/stainless-sdks/hubspot-sdk-go/internal/encoding/json"
	"github.com/stainless-sdks/hubspot-sdk-go/internal/requestconfig"
	"github.com/stainless-sdks/hubspot-sdk-go/option"
	"github.com/stainless-sdks/hubspot-sdk-go/packages/pagination"
	"github.com/stainless-sdks/hubspot-sdk-go/packages/param"
	"github.com/stainless-sdks/hubspot-sdk-go/packages/respjson"
	"github.com/stainless-sdks/hubspot-sdk-go/shared"
)

// UserService contains methods and other services that help with interacting with
// the hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewUserService] method instead.
type UserService struct {
	options []option.RequestOption
}

// NewUserService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewUserService(opts ...option.RequestOption) (r UserService) {
	r = UserService{}
	r.options = opts
	return
}

func (r *UserService) New(ctx context.Context, body UserNewParams, opts ...option.RequestOption) (res *PublicUser, err error) {
	opts = slices.Concat(r.options, opts)
	path := "settings/users/2026-03"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Modifies a user identified by `userId`. `userId` refers to the user's ID by
// default, or optionally email as specified by the `IdProperty` query param.
func (r *UserService) Update(ctx context.Context, userID string, params UserUpdateParams, opts ...option.RequestOption) (res *PublicUser, err error) {
	opts = slices.Concat(r.options, opts)
	if userID == "" {
		err = errors.New("missing required userId parameter")
		return nil, err
	}
	path := fmt.Sprintf("settings/users/2026-03/%s", url.PathEscape(userID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return res, err
}

func (r *UserService) List(ctx context.Context, query UserListParams, opts ...option.RequestOption) (res *pagination.Page[PublicUser], err error) {
	var raw *http.Response
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "settings/users/2026-03"
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

func (r *UserService) ListAutoPaging(ctx context.Context, query UserListParams, opts ...option.RequestOption) *pagination.PageAutoPager[PublicUser] {
	return pagination.NewPageAutoPager(r.List(ctx, query, opts...))
}

// Removes a user identified by `userId`. `userId` refers to the user's ID by
// default, or optionally email as specified by the `IdProperty` query param.
func (r *UserService) Delete(ctx context.Context, userID string, body UserDeleteParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if userID == "" {
		err = errors.New("missing required userId parameter")
		return err
	}
	path := fmt.Sprintf("settings/users/2026-03/%s", url.PathEscape(userID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, nil, opts...)
	return err
}

// Retrieves a user identified by `userId`. `userId` refers to the user's ID by
// default, or optionally email as specified by the `IdProperty` query param.
func (r *UserService) Get(ctx context.Context, userID string, query UserGetParams, opts ...option.RequestOption) (res *PublicUser, err error) {
	opts = slices.Concat(r.options, opts)
	if userID == "" {
		err = errors.New("missing required userId parameter")
		return nil, err
	}
	path := fmt.Sprintf("settings/users/2026-03/%s", url.PathEscape(userID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Retrieves the roles on an account
func (r *UserService) ListRoles(ctx context.Context, opts ...option.RequestOption) (res *CollectionResponsePublicPermissionSetNoPaging, err error) {
	opts = slices.Concat(r.options, opts)
	path := "settings/users/2026-03/roles"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// View teams for this account
func (r *UserService) ListTeams(ctx context.Context, opts ...option.RequestOption) (res *CollectionResponsePublicTeamNoPaging, err error) {
	opts = slices.Concat(r.options, opts)
	path := "settings/users/2026-03/teams"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type CollectionResponsePublicPermissionSetNoPaging struct {
	Results []PublicPermissionSet `json:"results" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Results     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CollectionResponsePublicPermissionSetNoPaging) RawJSON() string { return r.JSON.raw }
func (r *CollectionResponsePublicPermissionSetNoPaging) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CollectionResponsePublicTeamNoPaging struct {
	Results []PublicTeam `json:"results" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Results     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CollectionResponsePublicTeamNoPaging) RawJSON() string { return r.JSON.raw }
func (r *CollectionResponsePublicTeamNoPaging) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CollectionResponsePublicUserForwardPaging struct {
	Results []PublicUser         `json:"results" api:"required"`
	Paging  shared.ForwardPaging `json:"paging"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Results     respjson.Field
		Paging      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CollectionResponsePublicUserForwardPaging) RawJSON() string { return r.JSON.raw }
func (r *CollectionResponsePublicUserForwardPaging) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PublicPermissionSet struct {
	// The role's unique ID
	ID string `json:"id" api:"required"`
	// The role's name
	Name string `json:"name" api:"required"`
	// Whether this role has a paid seat and requires the billing-write scope to
	// assign/unassign to users
	RequiresBillingWrite bool `json:"requiresBillingWrite" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                   respjson.Field
		Name                 respjson.Field
		RequiresBillingWrite respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PublicPermissionSet) RawJSON() string { return r.JSON.raw }
func (r *PublicPermissionSet) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PublicTeam struct {
	// The team's unique ID
	ID string `json:"id" api:"required"`
	// The team's name
	Name string `json:"name" api:"required"`
	// Secondary or additional members of this team
	SecondaryUserIDs []string `json:"secondaryUserIds" api:"required"`
	// Primary members of this team
	UserIDs []string `json:"userIds" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		Name             respjson.Field
		SecondaryUserIDs respjson.Field
		UserIDs          respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PublicTeam) RawJSON() string { return r.JSON.raw }
func (r *PublicTeam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PublicUser struct {
	// The user's unique ID.
	ID string `json:"id" api:"required"`
	// The user's email.
	Email string `json:"email" api:"required"`
	// A list of role IDs assigned to the user.
	RoleIDs []string `json:"roleIds" api:"required"`
	// Whether the user has super admin privileges.
	SuperAdmin bool `json:"superAdmin" api:"required"`
	// The user's first name.
	FirstName string `json:"firstName"`
	// The user's last name.
	LastName string `json:"lastName"`
	// The user's primary team
	PrimaryTeamID string `json:"primaryTeamId"`
	// The user's role.
	RoleID string `json:"roleId"`
	// The user's additional teams.
	SecondaryTeamIDs []string `json:"secondaryTeamIds"`
	// Whether a welcome email was sent to the user. This value will only be populated
	// in response to a provisioning request. Subsequent queries will be false.
	SendWelcomeEmail bool `json:"sendWelcomeEmail"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		Email            respjson.Field
		RoleIDs          respjson.Field
		SuperAdmin       respjson.Field
		FirstName        respjson.Field
		LastName         respjson.Field
		PrimaryTeamID    respjson.Field
		RoleID           respjson.Field
		SecondaryTeamIDs respjson.Field
		SendWelcomeEmail respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PublicUser) RawJSON() string { return r.JSON.raw }
func (r *PublicUser) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PublicUserUpdateParam struct {
	// The first name of the user.
	FirstName param.Opt[string] `json:"firstName,omitzero"`
	// The last name of the user.
	LastName param.Opt[string] `json:"lastName,omitzero"`
	// The user's primary team.
	PrimaryTeamID param.Opt[string] `json:"primaryTeamId,omitzero"`
	// The user's role.
	RoleID param.Opt[string] `json:"roleId,omitzero"`
	// The user's additional teams.
	SecondaryTeamIDs []string `json:"secondaryTeamIds,omitzero"`
	paramObj
}

func (r PublicUserUpdateParam) MarshalJSON() (data []byte, err error) {
	type shadow PublicUserUpdateParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PublicUserUpdateParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Email, SendWelcomeEmail are required.
type UserProvisionRequestParam struct {
	// The user's email.
	Email string `json:"email" api:"required"`
	// Whether to send a welcome email.
	SendWelcomeEmail bool `json:"sendWelcomeEmail" api:"required"`
	// The user's first name.
	FirstName param.Opt[string] `json:"firstName,omitzero"`
	// The user's last name.
	LastName param.Opt[string] `json:"lastName,omitzero"`
	// The user's primary team.
	PrimaryTeamID param.Opt[string] `json:"primaryTeamId,omitzero"`
	// The user's role.
	RoleID param.Opt[string] `json:"roleId,omitzero"`
	// The user's additional teams.
	SecondaryTeamIDs []string `json:"secondaryTeamIds,omitzero"`
	paramObj
}

func (r UserProvisionRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow UserProvisionRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UserProvisionRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserNewParams struct {
	UserProvisionRequest UserProvisionRequestParam
	paramObj
}

func (r UserNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.UserProvisionRequest)
}
func (r *UserNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserUpdateParams struct {
	PublicUserUpdate PublicUserUpdateParam
	// Any of "EMAIL", "USER_ID".
	IDProperty UserUpdateParamsIDProperty `query:"idProperty,omitzero" json:"-"`
	paramObj
}

func (r UserUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.PublicUserUpdate)
}
func (r *UserUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [UserUpdateParams]'s query parameters as `url.Values`.
func (r UserUpdateParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type UserUpdateParamsIDProperty string

const (
	UserUpdateParamsIDPropertyEmail  UserUpdateParamsIDProperty = "EMAIL"
	UserUpdateParamsIDPropertyUserID UserUpdateParamsIDProperty = "USER_ID"
)

type UserListParams struct {
	// The paging cursor token of the last successfully read resource will be returned
	// as the `paging.next.after` JSON property of a paged response containing more
	// results.
	After param.Opt[string] `query:"after,omitzero" json:"-"`
	// The maximum number of results to display per page.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [UserListParams]'s query parameters as `url.Values`.
func (r UserListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type UserDeleteParams struct {
	// Any of "EMAIL", "USER_ID".
	IDProperty UserDeleteParamsIDProperty `query:"idProperty,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [UserDeleteParams]'s query parameters as `url.Values`.
func (r UserDeleteParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type UserDeleteParamsIDProperty string

const (
	UserDeleteParamsIDPropertyEmail  UserDeleteParamsIDProperty = "EMAIL"
	UserDeleteParamsIDPropertyUserID UserDeleteParamsIDProperty = "USER_ID"
)

type UserGetParams struct {
	// Any of "EMAIL", "USER_ID".
	IDProperty UserGetParamsIDProperty `query:"idProperty,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [UserGetParams]'s query parameters as `url.Values`.
func (r UserGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type UserGetParamsIDProperty string

const (
	UserGetParamsIDPropertyEmail  UserGetParamsIDProperty = "EMAIL"
	UserGetParamsIDPropertyUserID UserGetParamsIDProperty = "USER_ID"
)
