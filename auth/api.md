# Auth

## OAuth

Response Types:

- <a href="https://pkg.go.dev/github.com/HubSpot/hubspot-sdk-go/auth">auth</a>.<a href="https://pkg.go.dev/github.com/HubSpot/hubspot-sdk-go/auth#AccessTokenResponse">AccessTokenResponse</a>
- <a href="https://pkg.go.dev/github.com/HubSpot/hubspot-sdk-go/auth">auth</a>.<a href="https://pkg.go.dev/github.com/HubSpot/hubspot-sdk-go/auth#ClientCredentialsTokenResponse">ClientCredentialsTokenResponse</a>
- <a href="https://pkg.go.dev/github.com/HubSpot/hubspot-sdk-go/auth">auth</a>.<a href="https://pkg.go.dev/github.com/HubSpot/hubspot-sdk-go/auth#PublicAccessTokenInfoResponse">PublicAccessTokenInfoResponse</a>
- <a href="https://pkg.go.dev/github.com/HubSpot/hubspot-sdk-go/auth">auth</a>.<a href="https://pkg.go.dev/github.com/HubSpot/hubspot-sdk-go/auth#PublicRefreshTokenInfoResponse">PublicRefreshTokenInfoResponse</a>
- <a href="https://pkg.go.dev/github.com/HubSpot/hubspot-sdk-go/auth">auth</a>.<a href="https://pkg.go.dev/github.com/HubSpot/hubspot-sdk-go/auth#SignedAccessToken">SignedAccessToken</a>
- <a href="https://pkg.go.dev/github.com/HubSpot/hubspot-sdk-go/auth">auth</a>.<a href="https://pkg.go.dev/github.com/HubSpot/hubspot-sdk-go/auth#TokenInfoResponseBaseIfUnion">TokenInfoResponseBaseIfUnion</a>
- <a href="https://pkg.go.dev/github.com/HubSpot/hubspot-sdk-go/auth">auth</a>.<a href="https://pkg.go.dev/github.com/HubSpot/hubspot-sdk-go/auth#TokenResponseIfUnion">TokenResponseIfUnion</a>

Methods:

- <code title="post /oauth/2026-03/token">client.Auth.OAuth.<a href="https://pkg.go.dev/github.com/HubSpot/hubspot-sdk-go/auth#OAuthService.NewToken">NewToken</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/HubSpot/hubspot-sdk-go/auth">auth</a>.<a href="https://pkg.go.dev/github.com/HubSpot/hubspot-sdk-go/auth#OAuthNewTokenParams">OAuthNewTokenParams</a>) (\*<a href="https://pkg.go.dev/github.com/HubSpot/hubspot-sdk-go/auth">auth</a>.<a href="https://pkg.go.dev/github.com/HubSpot/hubspot-sdk-go/auth#TokenResponseIfUnion">TokenResponseIfUnion</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /oauth/2026-03/token/introspect">client.Auth.OAuth.<a href="https://pkg.go.dev/github.com/HubSpot/hubspot-sdk-go/auth#OAuthService.IntrospectToken">IntrospectToken</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/HubSpot/hubspot-sdk-go/auth">auth</a>.<a href="https://pkg.go.dev/github.com/HubSpot/hubspot-sdk-go/auth#OAuthIntrospectTokenParams">OAuthIntrospectTokenParams</a>) (\*<a href="https://pkg.go.dev/github.com/HubSpot/hubspot-sdk-go/auth">auth</a>.<a href="https://pkg.go.dev/github.com/HubSpot/hubspot-sdk-go/auth#TokenInfoResponseBaseIfUnion">TokenInfoResponseBaseIfUnion</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /oauth/2026-03/token/revoke">client.Auth.OAuth.<a href="https://pkg.go.dev/github.com/HubSpot/hubspot-sdk-go/auth#OAuthService.RevokeToken">RevokeToken</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/HubSpot/hubspot-sdk-go/auth">auth</a>.<a href="https://pkg.go.dev/github.com/HubSpot/hubspot-sdk-go/auth#OAuthRevokeTokenParams">OAuthRevokeTokenParams</a>) (\*http.Response, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
