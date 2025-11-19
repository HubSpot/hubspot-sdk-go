// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cms

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/stainless-sdks/hubspot-sdk-go/internal/apijson"
	"github.com/stainless-sdks/hubspot-sdk-go/internal/apiquery"
	"github.com/stainless-sdks/hubspot-sdk-go/internal/requestconfig"
	"github.com/stainless-sdks/hubspot-sdk-go/option"
	"github.com/stainless-sdks/hubspot-sdk-go/packages/pagination"
	"github.com/stainless-sdks/hubspot-sdk-go/packages/param"
	"github.com/stainless-sdks/hubspot-sdk-go/packages/respjson"
)

// URLRedirectService contains methods and other services that help with
// interacting with the Hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewURLRedirectService] method instead.
type URLRedirectService struct {
	Options []option.RequestOption
}

// NewURLRedirectService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewURLRedirectService(opts ...option.RequestOption) (r URLRedirectService) {
	r = URLRedirectService{}
	r.Options = opts
	return
}

// Creates and configures a new URL redirect.
func (r *URLRedirectService) New(ctx context.Context, body URLRedirectNewParams, opts ...option.RequestOption) (res *URLRedirectNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "cms/v3/url-redirects/"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Updates the settings for an existing URL redirect.
func (r *URLRedirectService) Update(ctx context.Context, urlRedirectID string, body URLRedirectUpdateParams, opts ...option.RequestOption) (res *URLRedirectUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if urlRedirectID == "" {
		err = errors.New("missing required urlRedirectId parameter")
		return
	}
	path := fmt.Sprintf("cms/v3/url-redirects/%s", urlRedirectID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Returns all existing URL redirects. Results can be limited and filtered by
// creation or updated date.
func (r *URLRedirectService) List(ctx context.Context, query URLRedirectListParams, opts ...option.RequestOption) (res *pagination.Page[URLRedirectListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "cms/v3/url-redirects/"
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

// Returns all existing URL redirects. Results can be limited and filtered by
// creation or updated date.
func (r *URLRedirectService) ListAutoPaging(ctx context.Context, query URLRedirectListParams, opts ...option.RequestOption) *pagination.PageAutoPager[URLRedirectListResponse] {
	return pagination.NewPageAutoPager(r.List(ctx, query, opts...))
}

// Delete one existing redirect, so it is no longer mapped.
func (r *URLRedirectService) Delete(ctx context.Context, urlRedirectID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if urlRedirectID == "" {
		err = errors.New("missing required urlRedirectId parameter")
		return
	}
	path := fmt.Sprintf("cms/v3/url-redirects/%s", urlRedirectID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Returns the details for a single existing URL redirect by ID.
func (r *URLRedirectService) Get(ctx context.Context, urlRedirectID string, opts ...option.RequestOption) (res *URLRedirectGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if urlRedirectID == "" {
		err = errors.New("missing required urlRedirectId parameter")
		return
	}
	path := fmt.Sprintf("cms/v3/url-redirects/%s", urlRedirectID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type URLRedirectNewResponse struct {
	// The unique ID of this URL redirect.
	ID string `json:"id,required"`
	// The destination URL, where the target URL should be redirected if it matches the
	// `routePrefix`.
	Destination string `json:"destination,required"`
	// Whether the `routePrefix` should match on the entire URL, including the domain.
	IsMatchFullURL bool `json:"isMatchFullUrl,required"`
	// Whether the `routePrefix` should match on the entire URL path, including the
	// query string.
	IsMatchQueryString bool `json:"isMatchQueryString,required"`
	// Whether the URL redirect mapping should apply only if a live page on the URL
	// isn't found. If False, the URL redirect mapping will take precedence over any
	// existing page.
	IsOnlyAfterNotFound bool `json:"isOnlyAfterNotFound,required"`
	// Whether the `routePrefix` should match based on pattern.
	IsPattern bool `json:"isPattern,required"`
	// Whether the `routePrefix` should match both HTTP and HTTPS protocols.
	IsProtocolAgnostic bool `json:"isProtocolAgnostic,required"`
	// Whether a trailing slash will be ignored.
	IsTrailingSlashOptional bool `json:"isTrailingSlashOptional,required"`
	// Used to prioritize URL redirection. If a given URL matches more than one
	// redirect, the one with the **lower** precedence will be used.
	Precedence int64 `json:"precedence,required"`
	// The type of redirect to create. Options include: 301 (permanent), 302
	// (temporary), or 305 (proxy). Find more details
	// [here](https://knowledge.hubspot.com/cos-general/how-to-redirect-a-hubspot-page).
	RedirectStyle int64 `json:"redirectStyle,required"`
	// The target incoming URL, path, or pattern to match for redirection.
	RoutePrefix string    `json:"routePrefix,required"`
	Created     time.Time `json:"created" format:"date-time"`
	Updated     time.Time `json:"updated" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                      respjson.Field
		Destination             respjson.Field
		IsMatchFullURL          respjson.Field
		IsMatchQueryString      respjson.Field
		IsOnlyAfterNotFound     respjson.Field
		IsPattern               respjson.Field
		IsProtocolAgnostic      respjson.Field
		IsTrailingSlashOptional respjson.Field
		Precedence              respjson.Field
		RedirectStyle           respjson.Field
		RoutePrefix             respjson.Field
		Created                 respjson.Field
		Updated                 respjson.Field
		ExtraFields             map[string]respjson.Field
		raw                     string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r URLRedirectNewResponse) RawJSON() string { return r.JSON.raw }
func (r *URLRedirectNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type URLRedirectUpdateResponse struct {
	// The unique ID of this URL redirect.
	ID string `json:"id,required"`
	// The destination URL, where the target URL should be redirected if it matches the
	// `routePrefix`.
	Destination string `json:"destination,required"`
	// Whether the `routePrefix` should match on the entire URL, including the domain.
	IsMatchFullURL bool `json:"isMatchFullUrl,required"`
	// Whether the `routePrefix` should match on the entire URL path, including the
	// query string.
	IsMatchQueryString bool `json:"isMatchQueryString,required"`
	// Whether the URL redirect mapping should apply only if a live page on the URL
	// isn't found. If False, the URL redirect mapping will take precedence over any
	// existing page.
	IsOnlyAfterNotFound bool `json:"isOnlyAfterNotFound,required"`
	// Whether the `routePrefix` should match based on pattern.
	IsPattern bool `json:"isPattern,required"`
	// Whether the `routePrefix` should match both HTTP and HTTPS protocols.
	IsProtocolAgnostic bool `json:"isProtocolAgnostic,required"`
	// Whether a trailing slash will be ignored.
	IsTrailingSlashOptional bool `json:"isTrailingSlashOptional,required"`
	// Used to prioritize URL redirection. If a given URL matches more than one
	// redirect, the one with the **lower** precedence will be used.
	Precedence int64 `json:"precedence,required"`
	// The type of redirect to create. Options include: 301 (permanent), 302
	// (temporary), or 305 (proxy). Find more details
	// [here](https://knowledge.hubspot.com/cos-general/how-to-redirect-a-hubspot-page).
	RedirectStyle int64 `json:"redirectStyle,required"`
	// The target incoming URL, path, or pattern to match for redirection.
	RoutePrefix string    `json:"routePrefix,required"`
	Created     time.Time `json:"created" format:"date-time"`
	Updated     time.Time `json:"updated" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                      respjson.Field
		Destination             respjson.Field
		IsMatchFullURL          respjson.Field
		IsMatchQueryString      respjson.Field
		IsOnlyAfterNotFound     respjson.Field
		IsPattern               respjson.Field
		IsProtocolAgnostic      respjson.Field
		IsTrailingSlashOptional respjson.Field
		Precedence              respjson.Field
		RedirectStyle           respjson.Field
		RoutePrefix             respjson.Field
		Created                 respjson.Field
		Updated                 respjson.Field
		ExtraFields             map[string]respjson.Field
		raw                     string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r URLRedirectUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *URLRedirectUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type URLRedirectListResponse struct {
	// The unique ID of this URL redirect.
	ID string `json:"id,required"`
	// The destination URL, where the target URL should be redirected if it matches the
	// `routePrefix`.
	Destination string `json:"destination,required"`
	// Whether the `routePrefix` should match on the entire URL, including the domain.
	IsMatchFullURL bool `json:"isMatchFullUrl,required"`
	// Whether the `routePrefix` should match on the entire URL path, including the
	// query string.
	IsMatchQueryString bool `json:"isMatchQueryString,required"`
	// Whether the URL redirect mapping should apply only if a live page on the URL
	// isn't found. If False, the URL redirect mapping will take precedence over any
	// existing page.
	IsOnlyAfterNotFound bool `json:"isOnlyAfterNotFound,required"`
	// Whether the `routePrefix` should match based on pattern.
	IsPattern bool `json:"isPattern,required"`
	// Whether the `routePrefix` should match both HTTP and HTTPS protocols.
	IsProtocolAgnostic bool `json:"isProtocolAgnostic,required"`
	// Whether a trailing slash will be ignored.
	IsTrailingSlashOptional bool `json:"isTrailingSlashOptional,required"`
	// Used to prioritize URL redirection. If a given URL matches more than one
	// redirect, the one with the **lower** precedence will be used.
	Precedence int64 `json:"precedence,required"`
	// The type of redirect to create. Options include: 301 (permanent), 302
	// (temporary), or 305 (proxy). Find more details
	// [here](https://knowledge.hubspot.com/cos-general/how-to-redirect-a-hubspot-page).
	RedirectStyle int64 `json:"redirectStyle,required"`
	// The target incoming URL, path, or pattern to match for redirection.
	RoutePrefix string    `json:"routePrefix,required"`
	Created     time.Time `json:"created" format:"date-time"`
	Updated     time.Time `json:"updated" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                      respjson.Field
		Destination             respjson.Field
		IsMatchFullURL          respjson.Field
		IsMatchQueryString      respjson.Field
		IsOnlyAfterNotFound     respjson.Field
		IsPattern               respjson.Field
		IsProtocolAgnostic      respjson.Field
		IsTrailingSlashOptional respjson.Field
		Precedence              respjson.Field
		RedirectStyle           respjson.Field
		RoutePrefix             respjson.Field
		Created                 respjson.Field
		Updated                 respjson.Field
		ExtraFields             map[string]respjson.Field
		raw                     string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r URLRedirectListResponse) RawJSON() string { return r.JSON.raw }
func (r *URLRedirectListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type URLRedirectGetResponse struct {
	// The unique ID of this URL redirect.
	ID string `json:"id,required"`
	// The destination URL, where the target URL should be redirected if it matches the
	// `routePrefix`.
	Destination string `json:"destination,required"`
	// Whether the `routePrefix` should match on the entire URL, including the domain.
	IsMatchFullURL bool `json:"isMatchFullUrl,required"`
	// Whether the `routePrefix` should match on the entire URL path, including the
	// query string.
	IsMatchQueryString bool `json:"isMatchQueryString,required"`
	// Whether the URL redirect mapping should apply only if a live page on the URL
	// isn't found. If False, the URL redirect mapping will take precedence over any
	// existing page.
	IsOnlyAfterNotFound bool `json:"isOnlyAfterNotFound,required"`
	// Whether the `routePrefix` should match based on pattern.
	IsPattern bool `json:"isPattern,required"`
	// Whether the `routePrefix` should match both HTTP and HTTPS protocols.
	IsProtocolAgnostic bool `json:"isProtocolAgnostic,required"`
	// Whether a trailing slash will be ignored.
	IsTrailingSlashOptional bool `json:"isTrailingSlashOptional,required"`
	// Used to prioritize URL redirection. If a given URL matches more than one
	// redirect, the one with the **lower** precedence will be used.
	Precedence int64 `json:"precedence,required"`
	// The type of redirect to create. Options include: 301 (permanent), 302
	// (temporary), or 305 (proxy). Find more details
	// [here](https://knowledge.hubspot.com/cos-general/how-to-redirect-a-hubspot-page).
	RedirectStyle int64 `json:"redirectStyle,required"`
	// The target incoming URL, path, or pattern to match for redirection.
	RoutePrefix string    `json:"routePrefix,required"`
	Created     time.Time `json:"created" format:"date-time"`
	Updated     time.Time `json:"updated" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                      respjson.Field
		Destination             respjson.Field
		IsMatchFullURL          respjson.Field
		IsMatchQueryString      respjson.Field
		IsOnlyAfterNotFound     respjson.Field
		IsPattern               respjson.Field
		IsProtocolAgnostic      respjson.Field
		IsTrailingSlashOptional respjson.Field
		Precedence              respjson.Field
		RedirectStyle           respjson.Field
		RoutePrefix             respjson.Field
		Created                 respjson.Field
		Updated                 respjson.Field
		ExtraFields             map[string]respjson.Field
		raw                     string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r URLRedirectGetResponse) RawJSON() string { return r.JSON.raw }
func (r *URLRedirectGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type URLRedirectNewParams struct {
	Destination             string           `json:"destination,required"`
	RedirectStyle           int64            `json:"redirectStyle,required"`
	RoutePrefix             string           `json:"routePrefix,required"`
	IsMatchFullURL          param.Opt[bool]  `json:"isMatchFullUrl,omitzero"`
	IsMatchQueryString      param.Opt[bool]  `json:"isMatchQueryString,omitzero"`
	IsOnlyAfterNotFound     param.Opt[bool]  `json:"isOnlyAfterNotFound,omitzero"`
	IsPattern               param.Opt[bool]  `json:"isPattern,omitzero"`
	IsProtocolAgnostic      param.Opt[bool]  `json:"isProtocolAgnostic,omitzero"`
	IsTrailingSlashOptional param.Opt[bool]  `json:"isTrailingSlashOptional,omitzero"`
	Precedence              param.Opt[int64] `json:"precedence,omitzero"`
	paramObj
}

func (r URLRedirectNewParams) MarshalJSON() (data []byte, err error) {
	type shadow URLRedirectNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *URLRedirectNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type URLRedirectUpdateParams struct {
	// The unique ID of this URL redirect.
	ID string `json:"id,required"`
	// The destination URL, where the target URL should be redirected if it matches the
	// `routePrefix`.
	Destination string `json:"destination,required"`
	// Whether the `routePrefix` should match on the entire URL, including the domain.
	IsMatchFullURL bool `json:"isMatchFullUrl,required"`
	// Whether the `routePrefix` should match on the entire URL path, including the
	// query string.
	IsMatchQueryString bool `json:"isMatchQueryString,required"`
	// Whether the URL redirect mapping should apply only if a live page on the URL
	// isn't found. If False, the URL redirect mapping will take precedence over any
	// existing page.
	IsOnlyAfterNotFound bool `json:"isOnlyAfterNotFound,required"`
	// Whether the `routePrefix` should match based on pattern.
	IsPattern bool `json:"isPattern,required"`
	// Whether the `routePrefix` should match both HTTP and HTTPS protocols.
	IsProtocolAgnostic bool `json:"isProtocolAgnostic,required"`
	// Whether a trailing slash will be ignored.
	IsTrailingSlashOptional bool `json:"isTrailingSlashOptional,required"`
	// Used to prioritize URL redirection. If a given URL matches more than one
	// redirect, the one with the **lower** precedence will be used.
	Precedence int64 `json:"precedence,required"`
	// The type of redirect to create. Options include: 301 (permanent), 302
	// (temporary), or 305 (proxy). Find more details
	// [here](https://knowledge.hubspot.com/cos-general/how-to-redirect-a-hubspot-page).
	RedirectStyle int64 `json:"redirectStyle,required"`
	// The target incoming URL, path, or pattern to match for redirection.
	RoutePrefix string               `json:"routePrefix,required"`
	Created     param.Opt[time.Time] `json:"created,omitzero" format:"date-time"`
	Updated     param.Opt[time.Time] `json:"updated,omitzero" format:"date-time"`
	paramObj
}

func (r URLRedirectUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow URLRedirectUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *URLRedirectUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type URLRedirectListParams struct {
	// The paging cursor token of the last successfully read resource will be returned
	// as the `paging.next.after` JSON property of a paged response containing more
	// results.
	After param.Opt[string] `query:"after,omitzero" json:"-"`
	// Whether to return only results that have been archived.
	Archived param.Opt[bool] `query:"archived,omitzero" json:"-"`
	// Only return redirects created after this date.
	CreatedAfter param.Opt[time.Time] `query:"createdAfter,omitzero" format:"date-time" json:"-"`
	// Only return redirects created on exactly this date.
	CreatedAt param.Opt[time.Time] `query:"createdAt,omitzero" format:"date-time" json:"-"`
	// Only return redirects created before this date.
	CreatedBefore param.Opt[time.Time] `query:"createdBefore,omitzero" format:"date-time" json:"-"`
	// Maximum number of result per page
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Only return redirects last updated after this date.
	UpdatedAfter param.Opt[time.Time] `query:"updatedAfter,omitzero" format:"date-time" json:"-"`
	// Only return redirects last updated on exactly this date.
	UpdatedAt param.Opt[time.Time] `query:"updatedAt,omitzero" format:"date-time" json:"-"`
	// Only return redirects last updated before this date.
	UpdatedBefore param.Opt[time.Time] `query:"updatedBefore,omitzero" format:"date-time" json:"-"`
	Sort          []string             `query:"sort,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [URLRedirectListParams]'s query parameters as `url.Values`.
func (r URLRedirectListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
