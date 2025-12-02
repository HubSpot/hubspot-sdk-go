// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package auth

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/stainless-sdks/hubspot-sdk-go/internal/apijson"
	"github.com/stainless-sdks/hubspot-sdk-go/internal/apiquery"
	"github.com/stainless-sdks/hubspot-sdk-go/internal/requestconfig"
	"github.com/stainless-sdks/hubspot-sdk-go/option"
	"github.com/stainless-sdks/hubspot-sdk-go/packages/param"
	"github.com/stainless-sdks/hubspot-sdk-go/packages/respjson"
)

// OAuthService contains methods and other services that help with interacting with
// the hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewOAuthService] method instead.
type OAuthService struct {
	Options []option.RequestOption
}

// NewOAuthService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewOAuthService(opts ...option.RequestOption) (r OAuthService) {
	r = OAuthService{}
	r.Options = opts
	return
}

// Use a
// [previously obtained refresh token](#get-oauth-2.0-access-and-refresh-tokens) to
// generate a new access token.
//
// Access tokens are short lived. You can check the `expires_in` parameter when
// generating an access token to determine its lifetime (in seconds). If you need
// offline access to HubSpot data, store the refresh token you get when
// [initiating your OAuth integration](https://developers.hubspot.com/docs/guides/api/app-management/oauth-tokens#initiating-oauth-access)
// and use it to generate a new access token once the initial one expires.
//
// Note: HubSpot access tokens will fluctuate in size as the information that's
// encoded in them changes over time. It's recommended to allow for tokens to be up
// to 300 characters to account for any potential changes.
func (r *OAuthService) NewAccessToken(ctx context.Context, params OAuthNewAccessTokenParams, opts ...option.RequestOption) (res *TokenResponseIf, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "oauth/v1/token"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Delete a refresh token, typically after a user uninstalls your app. Access
// tokens generated with the refresh token will not be affected.
//
// This will not uninstall the application from HubSpot or inhibit data syncing
// between an account and the app.
//
// Deprecated: deprecated
func (r *OAuthService) DeleteRefreshToken(ctx context.Context, token string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if token == "" {
		err = errors.New("missing required token parameter")
		return
	}
	path := fmt.Sprintf("oauth/v1/refresh-tokens/%s", token)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Retrieve a token's metadata, including the email address of the user that the
// token was created for and the ID of the account it's associated with.
//
// Note: HubSpot access tokens will fluctuate in size as the information that's
// encoded in them changes over time. It's recommended to allow for tokens to be up
// to 300 characters to account for any potential changes.
//
// Deprecated: deprecated
func (r *OAuthService) GetAccessToken(ctx context.Context, token string, opts ...option.RequestOption) (res *AccessTokenInfoResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if token == "" {
		err = errors.New("missing required token parameter")
		return
	}
	path := fmt.Sprintf("oauth/v1/access-tokens/%s", token)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Retrieve a refresh token's metadata, including the email address of the user
// that the token was created for and the ID of the account it's associated with.
// Learn more about
// [refresh tokens](https://developers.hubspot.com/docs/guides/api/app-management/oauth-tokens#generate-initial-access-and-refresh-tokens).
//
// Deprecated: deprecated
func (r *OAuthService) GetRefreshToken(ctx context.Context, token string, opts ...option.RequestOption) (res *RefreshTokenInfoResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if token == "" {
		err = errors.New("missing required token parameter")
		return
	}
	path := fmt.Sprintf("oauth/v1/refresh-tokens/%s", token)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type AccessTokenInfoResponse struct {
	Token                 string            `json:"token,required"`
	AppID                 int64             `json:"app_id,required"`
	ExpiresIn             int64             `json:"expires_in,required"`
	HubID                 int64             `json:"hub_id,required"`
	Scopes                []string          `json:"scopes,required"`
	TokenType             string            `json:"token_type,required"`
	UserID                int64             `json:"user_id,required"`
	HubDomain             string            `json:"hub_domain"`
	IsPrivateDistribution bool              `json:"is_private_distribution"`
	SignedAccessToken     SignedAccessToken `json:"signed_access_token"`
	User                  string            `json:"user"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Token                 respjson.Field
		AppID                 respjson.Field
		ExpiresIn             respjson.Field
		HubID                 respjson.Field
		Scopes                respjson.Field
		TokenType             respjson.Field
		UserID                respjson.Field
		HubDomain             respjson.Field
		IsPrivateDistribution respjson.Field
		SignedAccessToken     respjson.Field
		User                  respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AccessTokenInfoResponse) RawJSON() string { return r.JSON.raw }
func (r *AccessTokenInfoResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RefreshTokenInfoResponse struct {
	Token     string   `json:"token,required"`
	ClientID  string   `json:"client_id,required"`
	HubID     int64    `json:"hub_id,required"`
	Scopes    []string `json:"scopes,required"`
	TokenType string   `json:"token_type,required"`
	UserID    int64    `json:"user_id,required"`
	HubDomain string   `json:"hub_domain"`
	User      string   `json:"user"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Token       respjson.Field
		ClientID    respjson.Field
		HubID       respjson.Field
		Scopes      respjson.Field
		TokenType   respjson.Field
		UserID      respjson.Field
		HubDomain   respjson.Field
		User        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RefreshTokenInfoResponse) RawJSON() string { return r.JSON.raw }
func (r *RefreshTokenInfoResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SignedAccessToken struct {
	AppID                     int64  `json:"appId,required"`
	ExpiresAt                 int64  `json:"expiresAt,required"`
	HubID                     int64  `json:"hubId,required"`
	Hublet                    string `json:"hublet,required"`
	InstallingUserID          int64  `json:"installingUserId,required"`
	IsPrivateDistribution     bool   `json:"isPrivateDistribution,required"`
	IsServiceAccount          bool   `json:"isServiceAccount,required"`
	IsUserLevel               bool   `json:"isUserLevel,required"`
	NewSignature              string `json:"newSignature,required"`
	Scopes                    string `json:"scopes,required"`
	ScopeToScopeGroupPks      string `json:"scopeToScopeGroupPks,required"`
	Signature                 string `json:"signature,required"`
	TrialScopes               string `json:"trialScopes,required"`
	TrialScopeToScopeGroupPks string `json:"trialScopeToScopeGroupPks,required"`
	UserID                    int64  `json:"userId,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AppID                     respjson.Field
		ExpiresAt                 respjson.Field
		HubID                     respjson.Field
		Hublet                    respjson.Field
		InstallingUserID          respjson.Field
		IsPrivateDistribution     respjson.Field
		IsServiceAccount          respjson.Field
		IsUserLevel               respjson.Field
		NewSignature              respjson.Field
		Scopes                    respjson.Field
		ScopeToScopeGroupPks      respjson.Field
		Signature                 respjson.Field
		TrialScopes               respjson.Field
		TrialScopeToScopeGroupPks respjson.Field
		UserID                    respjson.Field
		ExtraFields               map[string]respjson.Field
		raw                       string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SignedAccessToken) RawJSON() string { return r.JSON.raw }
func (r *SignedAccessToken) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TokenResponseIf struct {
	AccessToken string   `json:"accessToken"`
	ExpiresIn   int64    `json:"expiresIn"`
	HubID       int64    `json:"hubId"`
	IDToken     string   `json:"idToken"`
	Scopes      []string `json:"scopes"`
	TokenType   string   `json:"tokenType"`
	UserID      int64    `json:"userId"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AccessToken respjson.Field
		ExpiresIn   respjson.Field
		HubID       respjson.Field
		IDToken     respjson.Field
		Scopes      respjson.Field
		TokenType   respjson.Field
		UserID      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TokenResponseIf) RawJSON() string { return r.JSON.raw }
func (r *TokenResponseIf) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OAuthNewAccessTokenParams struct {
	QueryClientSecret param.Opt[string] `query:"client_secret,omitzero" json:"-"`
	QueryRefreshToken param.Opt[string] `query:"refresh_token,omitzero" json:"-"`
	ClientID          param.Opt[string] `json:"client_id,omitzero"`
	BodyClientSecret  param.Opt[string] `json:"client_secret,omitzero"`
	Code              param.Opt[string] `json:"code,omitzero"`
	CodeVerifier      param.Opt[string] `json:"code_verifier,omitzero"`
	RedirectUri       param.Opt[string] `json:"redirect_uri,omitzero"`
	BodyRefreshToken  param.Opt[string] `json:"refresh_token,omitzero"`
	Scope             param.Opt[string] `json:"scope,omitzero"`
	// Any of "authorization_code", "client_credentials", "refresh_token".
	GrantType OAuthNewAccessTokenParamsGrantType `json:"grant_type,omitzero"`
	paramObj
}

func (r OAuthNewAccessTokenParams) MarshalJSON() (data []byte, err error) {
	type shadow OAuthNewAccessTokenParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OAuthNewAccessTokenParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [OAuthNewAccessTokenParams]'s query parameters as
// `url.Values`.
func (r OAuthNewAccessTokenParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type OAuthNewAccessTokenParamsGrantType string

const (
	OAuthNewAccessTokenParamsGrantTypeAuthorizationCode OAuthNewAccessTokenParamsGrantType = "authorization_code"
	OAuthNewAccessTokenParamsGrantTypeClientCredentials OAuthNewAccessTokenParamsGrantType = "client_credentials"
	OAuthNewAccessTokenParamsGrantTypeRefreshToken      OAuthNewAccessTokenParamsGrantType = "refresh_token"
)
