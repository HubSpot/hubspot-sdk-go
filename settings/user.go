// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package settings

import (
	"context"
	"encoding/json"
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
// the Hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewUserService] method instead.
type UserService struct {
	Options []option.RequestOption
}

// NewUserService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewUserService(opts ...option.RequestOption) (r UserService) {
	r = UserService{}
	r.Options = opts
	return
}

// New users will only have minimal permissions, which is contacts-base. A welcome
// email will prompt them to set a password and log in to HubSpot.
func (r *UserService) New(ctx context.Context, body UserNewParams, opts ...option.RequestOption) (res *PublicUser, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "settings/v3/users/"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Modifies a user identified by `userId`. `userId` refers to the user's ID by
// default, or optionally email as specified by the `IdProperty` query param.
func (r *UserService) Update(ctx context.Context, userID string, params UserUpdateParams, opts ...option.RequestOption) (res *PublicUser, err error) {
	opts = slices.Concat(r.Options, opts)
	if userID == "" {
		err = errors.New("missing required userId parameter")
		return
	}
	path := fmt.Sprintf("settings/v3/users/%s", userID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return
}

// Retrieves a list of users from an account
func (r *UserService) List(ctx context.Context, query UserListParams, opts ...option.RequestOption) (res *pagination.Page[PublicUser], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "settings/v3/users/"
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

// Retrieves a list of users from an account
func (r *UserService) ListAutoPaging(ctx context.Context, query UserListParams, opts ...option.RequestOption) *pagination.PageAutoPager[PublicUser] {
	return pagination.NewPageAutoPager(r.List(ctx, query, opts...))
}

// Removes a user identified by `userId`. `userId` refers to the user's ID by
// default, or optionally email as specified by the `IdProperty` query param.
func (r *UserService) Delete(ctx context.Context, userID string, body UserDeleteParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if userID == "" {
		err = errors.New("missing required userId parameter")
		return
	}
	path := fmt.Sprintf("settings/v3/users/%s", userID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, nil, opts...)
	return
}

// Retrieves a user identified by `userId`. `userId` refers to the user's ID by
// default, or optionally email as specified by the `IdProperty` query param.
func (r *UserService) Get(ctx context.Context, userID string, query UserGetParams, opts ...option.RequestOption) (res *PublicUser, err error) {
	opts = slices.Concat(r.Options, opts)
	if userID == "" {
		err = errors.New("missing required userId parameter")
		return
	}
	path := fmt.Sprintf("settings/v3/users/%s", userID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Retrieves the roles on an account
func (r *UserService) ListRoles(ctx context.Context, opts ...option.RequestOption) (res *CollectionResponsePublicPermissionSetNoPaging, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "settings/v3/users/roles"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// View teams for this account
func (r *UserService) ListTeams(ctx context.Context, opts ...option.RequestOption) (res *CollectionResponsePublicTeamNoPaging, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "settings/v3/users/teams"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type CollectionResponsePublicPermissionSetNoPaging struct {
	Results []PublicPermissionSet `json:"results,required"`
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
	Results []PublicTeam `json:"results,required"`
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
	Results []PublicUser         `json:"results,required"`
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

// A role that can be assigned to a user
type PublicPermissionSet struct {
	// The role's unique ID
	ID string `json:"id,required"`
	// The role's name
	Name string `json:"name,required"`
	// Whether this role has a paid seat and requires the billing-write scope to
	// assign/unassign to users
	RequiresBillingWrite bool `json:"requiresBillingWrite,required"`
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

// A team that can be assigned to a user
type PublicTeam struct {
	// The team's unique ID
	ID string `json:"id,required"`
	// The team's name
	Name string `json:"name,required"`
	// Secondary or additional members of this team
	SecondaryUserIDs []string `json:"secondaryUserIds,required"`
	// Primary members of this team
	UserIDs []string `json:"userIds,required"`
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

// A user
type PublicUser struct {
	// The user's unique ID
	ID string `json:"id,required"`
	// The user's email
	Email     string `json:"email,required"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	// The user's primary team
	PrimaryTeamID string `json:"primaryTeamId"`
	// The user's role
	RoleID  string   `json:"roleId"`
	RoleIDs []string `json:"roleIds"`
	// The user's additional teams
	SecondaryTeamIDs []string `json:"secondaryTeamIds"`
	SendWelcomeEmail bool     `json:"sendWelcomeEmail"`
	SuperAdmin       bool     `json:"superAdmin"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		Email            respjson.Field
		FirstName        respjson.Field
		LastName         respjson.Field
		PrimaryTeamID    respjson.Field
		RoleID           respjson.Field
		RoleIDs          respjson.Field
		SecondaryTeamIDs respjson.Field
		SendWelcomeEmail respjson.Field
		SuperAdmin       respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PublicUser) RawJSON() string { return r.JSON.raw }
func (r *PublicUser) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A user to update
type PublicUserUpdateParam struct {
	FirstName param.Opt[string] `json:"firstName,omitzero"`
	LastName  param.Opt[string] `json:"lastName,omitzero"`
	// The user's primary team
	PrimaryTeamID param.Opt[string] `json:"primaryTeamId,omitzero"`
	// The user's role
	RoleID param.Opt[string] `json:"roleId,omitzero"`
	// The user's additional teams
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

// A user creation request
//
// The property Email is required.
type UserProvisionRequestParam struct {
	// The created user's email
	Email     string            `json:"email,required"`
	FirstName param.Opt[string] `json:"firstName,omitzero"`
	LastName  param.Opt[string] `json:"lastName,omitzero"`
	// The user's primary team
	PrimaryTeamID param.Opt[string] `json:"primaryTeamId,omitzero"`
	// The user's role
	RoleID param.Opt[string] `json:"roleId,omitzero"`
	// Whether to send a welcome email
	SendWelcomeEmail param.Opt[bool] `json:"sendWelcomeEmail,omitzero"`
	// The user's additional teams
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
	// A user creation request
	UserProvisionRequest UserProvisionRequestParam
	paramObj
}

func (r UserNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.UserProvisionRequest)
}
func (r *UserNewParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.UserProvisionRequest)
}

type UserUpdateParams struct {
	// A user to update
	PublicUserUpdate PublicUserUpdateParam
	// The name of a property with unique user values. Valid values are
	// `USER_ID`(default) or `EMAIL`
	//
	// Any of "USER_ID", "EMAIL".
	IDProperty UserUpdateParamsIDProperty `query:"idProperty,omitzero" json:"-"`
	paramObj
}

func (r UserUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.PublicUserUpdate)
}
func (r *UserUpdateParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.PublicUserUpdate)
}

// URLQuery serializes [UserUpdateParams]'s query parameters as `url.Values`.
func (r UserUpdateParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// The name of a property with unique user values. Valid values are
// `USER_ID`(default) or `EMAIL`
type UserUpdateParamsIDProperty string

const (
	UserUpdateParamsIDPropertyUserID UserUpdateParamsIDProperty = "USER_ID"
	UserUpdateParamsIDPropertyEmail  UserUpdateParamsIDProperty = "EMAIL"
)

type UserListParams struct {
	// Results will display maximum 100 users per page. Additional results will be on
	// the next page.
	After param.Opt[string] `query:"after,omitzero" json:"-"`
	// The number of users to retrieve
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
	// The name of a property with unique user values. Valid values are
	// `USER_ID`(default) or `EMAIL`
	//
	// Any of "USER_ID", "EMAIL".
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

// The name of a property with unique user values. Valid values are
// `USER_ID`(default) or `EMAIL`
type UserDeleteParamsIDProperty string

const (
	UserDeleteParamsIDPropertyUserID UserDeleteParamsIDProperty = "USER_ID"
	UserDeleteParamsIDPropertyEmail  UserDeleteParamsIDProperty = "EMAIL"
)

type UserGetParams struct {
	// The name of a property with unique user values. Valid values are
	// `USER_ID`(default) or `EMAIL`
	//
	// Any of "USER_ID", "EMAIL".
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

// The name of a property with unique user values. Valid values are
// `USER_ID`(default) or `EMAIL`
type UserGetParamsIDProperty string

const (
	UserGetParamsIDPropertyUserID UserGetParamsIDProperty = "USER_ID"
	UserGetParamsIDPropertyEmail  UserGetParamsIDProperty = "EMAIL"
)
