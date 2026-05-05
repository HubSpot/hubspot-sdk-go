// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package auth

import (
	"context"
	"encoding/json"
	"net/http"
	"slices"

	"github.com/HubSpot/hubspot-sdk-go/internal/apijson"
	"github.com/HubSpot/hubspot-sdk-go/internal/requestconfig"
	"github.com/HubSpot/hubspot-sdk-go/option"
	"github.com/HubSpot/hubspot-sdk-go/packages/param"
	"github.com/HubSpot/hubspot-sdk-go/packages/respjson"
)

// OAuthService contains methods and other services that help with interacting with
// the hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewOAuthService] method instead.
type OAuthService struct {
	options []option.RequestOption
}

// NewOAuthService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewOAuthService(opts ...option.RequestOption) (r OAuthService) {
	r = OAuthService{}
	r.options = opts
	return
}

// Authenticates a client and returns access and refresh tokens.
func (r *OAuthService) NewToken(ctx context.Context, body OAuthNewTokenParams, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "oauth/2026-03/token"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Returns validity and metadata for access and refresh tokens.
func (r *OAuthService) IntrospectToken(ctx context.Context, body OAuthIntrospectTokenParams, opts ...option.RequestOption) (res *TokenInfoResponseBaseIfUnion, err error) {
	opts = slices.Concat(r.options, opts)
	path := "oauth/2026-03/token/introspect"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Deletes/Revokes provided Refresh Token
func (r *OAuthService) RevokeToken(ctx context.Context, body OAuthRevokeTokenParams, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "oauth/2026-03/token/revoke"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type PublicAccessTokenInfoResponse struct {
	Token                 string            `json:"token" api:"required"`
	Active                bool              `json:"active" api:"required"`
	AppID                 int64             `json:"app_id" api:"required"`
	ClientID              string            `json:"client_id" api:"required"`
	ExpiresIn             int64             `json:"expires_in" api:"required"`
	HubID                 int64             `json:"hub_id" api:"required"`
	IsPrivateDistribution bool              `json:"is_private_distribution" api:"required"`
	Scopes                []string          `json:"scopes" api:"required"`
	SignedAccessToken     SignedAccessToken `json:"signed_access_token" api:"required"`
	TokenType             string            `json:"token_type" api:"required"`
	// Any of "access_token".
	TokenUse  PublicAccessTokenInfoResponseTokenUse `json:"token_use" api:"required"`
	UserID    int64                                 `json:"user_id" api:"required"`
	HubDomain string                                `json:"hub_domain"`
	User      string                                `json:"user"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Token                 respjson.Field
		Active                respjson.Field
		AppID                 respjson.Field
		ClientID              respjson.Field
		ExpiresIn             respjson.Field
		HubID                 respjson.Field
		IsPrivateDistribution respjson.Field
		Scopes                respjson.Field
		SignedAccessToken     respjson.Field
		TokenType             respjson.Field
		TokenUse              respjson.Field
		UserID                respjson.Field
		HubDomain             respjson.Field
		User                  respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PublicAccessTokenInfoResponse) RawJSON() string { return r.JSON.raw }
func (r *PublicAccessTokenInfoResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PublicAccessTokenInfoResponseTokenUse string

const (
	PublicAccessTokenInfoResponseTokenUseAccessToken PublicAccessTokenInfoResponseTokenUse = "access_token"
)

type PublicRefreshTokenInfoResponse struct {
	Token     string   `json:"token" api:"required"`
	Active    bool     `json:"active" api:"required"`
	AppID     int64    `json:"app_id" api:"required"`
	ClientID  string   `json:"client_id" api:"required"`
	HubID     int64    `json:"hub_id" api:"required"`
	Scopes    []string `json:"scopes" api:"required"`
	TokenType string   `json:"token_type" api:"required"`
	// Any of "refresh_token".
	TokenUse  PublicRefreshTokenInfoResponseTokenUse `json:"token_use" api:"required"`
	UserID    int64                                  `json:"user_id" api:"required"`
	HubDomain string                                 `json:"hub_domain"`
	User      string                                 `json:"user"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Token       respjson.Field
		Active      respjson.Field
		AppID       respjson.Field
		ClientID    respjson.Field
		HubID       respjson.Field
		Scopes      respjson.Field
		TokenType   respjson.Field
		TokenUse    respjson.Field
		UserID      respjson.Field
		HubDomain   respjson.Field
		User        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PublicRefreshTokenInfoResponse) RawJSON() string { return r.JSON.raw }
func (r *PublicRefreshTokenInfoResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PublicRefreshTokenInfoResponseTokenUse string

const (
	PublicRefreshTokenInfoResponseTokenUseRefreshToken PublicRefreshTokenInfoResponseTokenUse = "refresh_token"
)

type SignedAccessToken struct {
	AppID                     int64  `json:"appId" api:"required"`
	ExpiresAt                 int64  `json:"expiresAt" api:"required"`
	HubID                     int64  `json:"hubId" api:"required"`
	Hublet                    string `json:"hublet" api:"required"`
	InstallingUserID          int64  `json:"installingUserId" api:"required"`
	IsPrivateDistribution     bool   `json:"isPrivateDistribution" api:"required"`
	IsServiceAccount          bool   `json:"isServiceAccount" api:"required"`
	IsUserLevel               bool   `json:"isUserLevel" api:"required"`
	NewSignature              string `json:"newSignature" api:"required"`
	Scopes                    string `json:"scopes" api:"required"`
	ScopeToScopeGroupPks      string `json:"scopeToScopeGroupPks" api:"required"`
	Signature                 string `json:"signature" api:"required"`
	TrialScopes               string `json:"trialScopes" api:"required"`
	TrialScopeToScopeGroupPks string `json:"trialScopeToScopeGroupPks" api:"required"`
	UserID                    int64  `json:"userId" api:"required"`
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

// TokenInfoResponseBaseIfUnion contains all possible properties and values from
// [PublicAccessTokenInfoResponse], [PublicRefreshTokenInfoResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type TokenInfoResponseBaseIfUnion struct {
	Token    string `json:"token"`
	Active   bool   `json:"active"`
	AppID    int64  `json:"app_id"`
	ClientID string `json:"client_id"`
	// This field is from variant [PublicAccessTokenInfoResponse].
	ExpiresIn int64 `json:"expires_in"`
	HubID     int64 `json:"hub_id"`
	// This field is from variant [PublicAccessTokenInfoResponse].
	IsPrivateDistribution bool     `json:"is_private_distribution"`
	Scopes                []string `json:"scopes"`
	// This field is from variant [PublicAccessTokenInfoResponse].
	SignedAccessToken SignedAccessToken `json:"signed_access_token"`
	TokenType         string            `json:"token_type"`
	TokenUse          string            `json:"token_use"`
	UserID            int64             `json:"user_id"`
	HubDomain         string            `json:"hub_domain"`
	User              string            `json:"user"`
	JSON              struct {
		Token                 respjson.Field
		Active                respjson.Field
		AppID                 respjson.Field
		ClientID              respjson.Field
		ExpiresIn             respjson.Field
		HubID                 respjson.Field
		IsPrivateDistribution respjson.Field
		Scopes                respjson.Field
		SignedAccessToken     respjson.Field
		TokenType             respjson.Field
		TokenUse              respjson.Field
		UserID                respjson.Field
		HubDomain             respjson.Field
		User                  respjson.Field
		raw                   string
	} `json:"-"`
}

func (u TokenInfoResponseBaseIfUnion) AsPublicAccessTokenInfoResponse() (v PublicAccessTokenInfoResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TokenInfoResponseBaseIfUnion) AsPublicRefreshTokenInfoResponse() (v PublicRefreshTokenInfoResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u TokenInfoResponseBaseIfUnion) RawJSON() string { return u.JSON.raw }

func (r *TokenInfoResponseBaseIfUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OAuthNewTokenParams struct {
	ClientID     param.Opt[string] `json:"client_id,omitzero"`
	ClientSecret param.Opt[string] `json:"client_secret,omitzero"`
	Code         param.Opt[string] `json:"code,omitzero"`
	CodeVerifier param.Opt[string] `json:"code_verifier,omitzero"`
	RedirectUri  param.Opt[string] `json:"redirect_uri,omitzero"`
	RefreshToken param.Opt[string] `json:"refresh_token,omitzero"`
	Scope        param.Opt[string] `json:"scope,omitzero"`
	// Any of "authorization_code", "refresh_token".
	GrantType OAuthNewTokenParamsGrantType `json:"grant_type,omitzero"`
	paramObj
}

func (r OAuthNewTokenParams) MarshalJSON() (data []byte, err error) {
	type shadow OAuthNewTokenParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OAuthNewTokenParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OAuthNewTokenParamsGrantType string

const (
	OAuthNewTokenParamsGrantTypeAuthorizationCode OAuthNewTokenParamsGrantType = "authorization_code"
	OAuthNewTokenParamsGrantTypeRefreshToken      OAuthNewTokenParamsGrantType = "refresh_token"
)

type OAuthIntrospectTokenParams struct {
	Token         param.Opt[string] `json:"token,omitzero"`
	ClientID      param.Opt[string] `json:"client_id,omitzero"`
	ClientSecret  param.Opt[string] `json:"client_secret,omitzero"`
	TokenTypeHint param.Opt[string] `json:"token_type_hint,omitzero"`
	paramObj
}

func (r OAuthIntrospectTokenParams) MarshalJSON() (data []byte, err error) {
	type shadow OAuthIntrospectTokenParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OAuthIntrospectTokenParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OAuthRevokeTokenParams struct {
	Token         param.Opt[string] `json:"token,omitzero"`
	ClientID      param.Opt[string] `json:"client_id,omitzero"`
	ClientSecret  param.Opt[string] `json:"client_secret,omitzero"`
	TokenTypeHint param.Opt[string] `json:"token_type_hint,omitzero"`
	paramObj
}

func (r OAuthRevokeTokenParams) MarshalJSON() (data []byte, err error) {
	type shadow OAuthRevokeTokenParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OAuthRevokeTokenParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
