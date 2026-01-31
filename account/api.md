# Account

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/hubspot-sdk-go/account">account</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/hubspot-sdk-go/account#APIUsage">APIUsage</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/hubspot-sdk-go/account">account</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/hubspot-sdk-go/account#CollectionResponseAPIUsage">CollectionResponseAPIUsage</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/hubspot-sdk-go/account">account</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/hubspot-sdk-go/account#PortalInformationResponse">PortalInformationResponse</a>

## Activity

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/hubspot-sdk-go/account">account</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/hubspot-sdk-go/account#ActingUser">ActingUser</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/hubspot-sdk-go/account">account</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/hubspot-sdk-go/account#CollectionResponseHydratedCriticalActionForwardPaging">CollectionResponseHydratedCriticalActionForwardPaging</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/hubspot-sdk-go/account">account</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/hubspot-sdk-go/account#CollectionResponsePublicAPIUserActionEventForwardPaging">CollectionResponsePublicAPIUserActionEventForwardPaging</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/hubspot-sdk-go/account">account</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/hubspot-sdk-go/account#CollectionResponsePublicLoginAuditForwardPaging">CollectionResponsePublicLoginAuditForwardPaging</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/hubspot-sdk-go/account">account</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/hubspot-sdk-go/account#HydratedCriticalAction">HydratedCriticalAction</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/hubspot-sdk-go/account">account</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/hubspot-sdk-go/account#PublicAPIUserActionEvent">PublicAPIUserActionEvent</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/hubspot-sdk-go/account">account</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/hubspot-sdk-go/account#PublicLoginAudit">PublicLoginAudit</a>

Methods:

- <code title="get /account-info/v3/activity/audit-logs">client.Account.Activity.<a href="https://pkg.go.dev/github.com/stainless-sdks/hubspot-sdk-go/account#ActivityService.ListAuditLogs">ListAuditLogs</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/hubspot-sdk-go/account">account</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/hubspot-sdk-go/account#ActivityListAuditLogsParams">ActivityListAuditLogsParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/hubspot-sdk-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/hubspot-sdk-go/packages/pagination#Page">Page</a>[<a href="https://pkg.go.dev/github.com/stainless-sdks/hubspot-sdk-go/account">account</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/hubspot-sdk-go/account#PublicAPIUserActionEvent">PublicAPIUserActionEvent</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /account-info/v3/activity/login">client.Account.Activity.<a href="https://pkg.go.dev/github.com/stainless-sdks/hubspot-sdk-go/account#ActivityService.ListLoginActivities">ListLoginActivities</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/hubspot-sdk-go/account">account</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/hubspot-sdk-go/account#ActivityListLoginActivitiesParams">ActivityListLoginActivitiesParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/hubspot-sdk-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/hubspot-sdk-go/packages/pagination#Page">Page</a>[<a href="https://pkg.go.dev/github.com/stainless-sdks/hubspot-sdk-go/account">account</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/hubspot-sdk-go/account#PublicLoginAudit">PublicLoginAudit</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /account-info/v3/activity/security">client.Account.Activity.<a href="https://pkg.go.dev/github.com/stainless-sdks/hubspot-sdk-go/account#ActivityService.ListSecurityActivities">ListSecurityActivities</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/hubspot-sdk-go/account">account</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/hubspot-sdk-go/account#ActivityListSecurityActivitiesParams">ActivityListSecurityActivitiesParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/hubspot-sdk-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/hubspot-sdk-go/packages/pagination#Page">Page</a>[<a href="https://pkg.go.dev/github.com/stainless-sdks/hubspot-sdk-go/account">account</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/hubspot-sdk-go/account#HydratedCriticalAction">HydratedCriticalAction</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Details

Methods:

- <code title="get /account-info/v3/details">client.Account.Details.<a href="https://pkg.go.dev/github.com/stainless-sdks/hubspot-sdk-go/account#DetailService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/hubspot-sdk-go/account">account</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/hubspot-sdk-go/account#PortalInformationResponse">PortalInformationResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Usage

Methods:

- <code title="get /account-info/v3/api-usage/daily/private-apps">client.Account.Usage.<a href="https://pkg.go.dev/github.com/stainless-sdks/hubspot-sdk-go/account#UsageService.GetDailyPrivateAppsUsage">GetDailyPrivateAppsUsage</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/hubspot-sdk-go/account">account</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/hubspot-sdk-go/account#CollectionResponseAPIUsage">CollectionResponseAPIUsage</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
