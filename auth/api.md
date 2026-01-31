# Auth

## OAuth

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/hubspot-sdk-go/auth">auth</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/hubspot-sdk-go/auth#AccessTokenInfoResponse">AccessTokenInfoResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/hubspot-sdk-go/auth">auth</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/hubspot-sdk-go/auth#RefreshTokenInfoResponse">RefreshTokenInfoResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/hubspot-sdk-go/auth">auth</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/hubspot-sdk-go/auth#SignedAccessToken">SignedAccessToken</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/hubspot-sdk-go/auth">auth</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/hubspot-sdk-go/auth#TokenResponseIf">TokenResponseIf</a>

Methods:

- <code title="post /oauth/v1/token">client.Auth.OAuth.<a href="https://pkg.go.dev/github.com/stainless-sdks/hubspot-sdk-go/auth#OAuthService.NewAccessToken">NewAccessToken</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/stainless-sdks/hubspot-sdk-go/auth">auth</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/hubspot-sdk-go/auth#OAuthNewAccessTokenParams">OAuthNewAccessTokenParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/hubspot-sdk-go/auth">auth</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/hubspot-sdk-go/auth#TokenResponseIf">TokenResponseIf</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /oauth/v1/refresh-tokens/{token}">client.Auth.OAuth.<a href="https://pkg.go.dev/github.com/stainless-sdks/hubspot-sdk-go/auth#OAuthService.DeleteRefreshToken">DeleteRefreshToken</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, token <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="get /oauth/v1/access-tokens/{token}">client.Auth.OAuth.<a href="https://pkg.go.dev/github.com/stainless-sdks/hubspot-sdk-go/auth#OAuthService.GetAccessToken">GetAccessToken</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, token <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/hubspot-sdk-go/auth">auth</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/hubspot-sdk-go/auth#AccessTokenInfoResponse">AccessTokenInfoResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /oauth/v1/refresh-tokens/{token}">client.Auth.OAuth.<a href="https://pkg.go.dev/github.com/stainless-sdks/hubspot-sdk-go/auth#OAuthService.GetRefreshToken">GetRefreshToken</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, token <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/hubspot-sdk-go/auth">auth</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/hubspot-sdk-go/auth#RefreshTokenInfoResponse">RefreshTokenInfoResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
