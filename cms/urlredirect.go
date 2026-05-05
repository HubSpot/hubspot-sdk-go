// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cms

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

// URLRedirectService contains methods and other services that help with
// interacting with the hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewURLRedirectService] method instead.
type URLRedirectService struct {
	options []option.RequestOption
}

// NewURLRedirectService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewURLRedirectService(opts ...option.RequestOption) (r URLRedirectService) {
	r = URLRedirectService{}
	r.options = opts
	return
}

// Create a new URL redirect in your HubSpot account. This endpoint allows you to
// define a new URL mapping that redirects traffic from a specified route to a
// destination URL. This is useful for managing URL changes, handling outdated
// links, or creating short links.
func (r *URLRedirectService) New(ctx context.Context, body URLRedirectNewParams, opts ...option.RequestOption) (res *URLMapping, err error) {
	opts = slices.Concat(r.options, opts)
	path := "cms/url-redirects/2026-03"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Updates the settings for an existing URL redirect.
func (r *URLRedirectService) Update(ctx context.Context, urlRedirectID string, body URLRedirectUpdateParams, opts ...option.RequestOption) (res *URLMapping, err error) {
	opts = slices.Concat(r.options, opts)
	if urlRedirectID == "" {
		err = errors.New("missing required urlRedirectId parameter")
		return nil, err
	}
	path := fmt.Sprintf("cms/url-redirects/2026-03/%s", url.PathEscape(urlRedirectID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

// Retrieve a list of URL redirects configured in your HubSpot account. This
// endpoint allows you to filter redirects based on their creation or update
// timestamps, and sort the results. It supports pagination and can include
// archived redirects if specified.
func (r *URLRedirectService) List(ctx context.Context, query URLRedirectListParams, opts ...option.RequestOption) (res *pagination.Page[URLMapping], err error) {
	var raw *http.Response
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "cms/url-redirects/2026-03"
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

// Retrieve a list of URL redirects configured in your HubSpot account. This
// endpoint allows you to filter redirects based on their creation or update
// timestamps, and sort the results. It supports pagination and can include
// archived redirects if specified.
func (r *URLRedirectService) ListAutoPaging(ctx context.Context, query URLRedirectListParams, opts ...option.RequestOption) *pagination.PageAutoPager[URLMapping] {
	return pagination.NewPageAutoPager(r.List(ctx, query, opts...))
}

// Delete one existing redirect, so it is no longer mapped.
func (r *URLRedirectService) Delete(ctx context.Context, urlRedirectID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if urlRedirectID == "" {
		err = errors.New("missing required urlRedirectId parameter")
		return err
	}
	path := fmt.Sprintf("cms/url-redirects/2026-03/%s", url.PathEscape(urlRedirectID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Returns the details for a single existing URL redirect by ID.
func (r *URLRedirectService) Get(ctx context.Context, urlRedirectID string, opts ...option.RequestOption) (res *URLMapping, err error) {
	opts = slices.Concat(r.options, opts)
	if urlRedirectID == "" {
		err = errors.New("missing required urlRedirectId parameter")
		return nil, err
	}
	path := fmt.Sprintf("cms/url-redirects/2026-03/%s", url.PathEscape(urlRedirectID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type CollectionResponseWithTotalURLMappingForwardPaging struct {
	// An array of UrlMapping objects, each representing a specific URL mapping.
	Results []URLMapping `json:"results" api:"required"`
	// The total number of URL mappings available.
	Total  int64                `json:"total" api:"required"`
	Paging shared.ForwardPaging `json:"paging"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Results     respjson.Field
		Total       respjson.Field
		Paging      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CollectionResponseWithTotalURLMappingForwardPaging) RawJSON() string { return r.JSON.raw }
func (r *CollectionResponseWithTotalURLMappingForwardPaging) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type URLMapping struct {
	// The unique ID of this URL redirect.
	ID string `json:"id" api:"required"`
	// The date and time when the URL mapping was initially created.
	Created time.Time `json:"created" api:"required" format:"date-time"`
	// The destination URL, where the target URL should be redirected if it matches the
	// `routePrefix`.
	Destination string `json:"destination" api:"required"`
	// Whether the `routePrefix` should match on the entire URL, including the domain.
	IsMatchFullURL bool `json:"isMatchFullUrl" api:"required"`
	// Whether the `routePrefix` should match on the entire URL path, including the
	// query string.
	IsMatchQueryString bool `json:"isMatchQueryString" api:"required"`
	// Whether the URL redirect mapping should apply only if a live page on the URL
	// isn't found. If False, the URL redirect mapping will take precedence over any
	// existing page.
	IsOnlyAfterNotFound bool `json:"isOnlyAfterNotFound" api:"required"`
	// Whether the `routePrefix` should match based on pattern.
	IsPattern bool `json:"isPattern" api:"required"`
	// Whether the `routePrefix` should match both HTTP and HTTPS protocols.
	IsProtocolAgnostic bool `json:"isProtocolAgnostic" api:"required"`
	// Whether a trailing slash will be ignored.
	IsTrailingSlashOptional bool `json:"isTrailingSlashOptional" api:"required"`
	// Used to prioritize URL redirection. If a given URL matches more than one
	// redirect, the one with the **lower** precedence will be used.
	Precedence int64 `json:"precedence" api:"required"`
	// The type of redirect to create. Options include: 301 (permanent), 302
	// (temporary), or 305 (proxy). Find more details
	// [here](https://knowledge.hubspot.com/cos-general/how-to-redirect-a-hubspot-page).
	RedirectStyle int64 `json:"redirectStyle" api:"required"`
	// The target incoming URL, path, or pattern to match for redirection.
	RoutePrefix string `json:"routePrefix" api:"required"`
	// The date and time when the URL mapping was last modified.
	Updated time.Time `json:"updated" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                      respjson.Field
		Created                 respjson.Field
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
		Updated                 respjson.Field
		ExtraFields             map[string]respjson.Field
		raw                     string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r URLMapping) RawJSON() string { return r.JSON.raw }
func (r *URLMapping) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this URLMapping to a URLMappingParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// URLMappingParam.Overrides()
func (r URLMapping) ToParam() URLMappingParam {
	return param.Override[URLMappingParam](json.RawMessage(r.RawJSON()))
}

// The properties ID, Created, Destination, IsMatchFullURL, IsMatchQueryString,
// IsOnlyAfterNotFound, IsPattern, IsProtocolAgnostic, IsTrailingSlashOptional,
// Precedence, RedirectStyle, RoutePrefix, Updated are required.
type URLMappingParam struct {
	// The unique ID of this URL redirect.
	ID string `json:"id" api:"required"`
	// The date and time when the URL mapping was initially created.
	Created time.Time `json:"created" api:"required" format:"date-time"`
	// The destination URL, where the target URL should be redirected if it matches the
	// `routePrefix`.
	Destination string `json:"destination" api:"required"`
	// Whether the `routePrefix` should match on the entire URL, including the domain.
	IsMatchFullURL bool `json:"isMatchFullUrl" api:"required"`
	// Whether the `routePrefix` should match on the entire URL path, including the
	// query string.
	IsMatchQueryString bool `json:"isMatchQueryString" api:"required"`
	// Whether the URL redirect mapping should apply only if a live page on the URL
	// isn't found. If False, the URL redirect mapping will take precedence over any
	// existing page.
	IsOnlyAfterNotFound bool `json:"isOnlyAfterNotFound" api:"required"`
	// Whether the `routePrefix` should match based on pattern.
	IsPattern bool `json:"isPattern" api:"required"`
	// Whether the `routePrefix` should match both HTTP and HTTPS protocols.
	IsProtocolAgnostic bool `json:"isProtocolAgnostic" api:"required"`
	// Whether a trailing slash will be ignored.
	IsTrailingSlashOptional bool `json:"isTrailingSlashOptional" api:"required"`
	// Used to prioritize URL redirection. If a given URL matches more than one
	// redirect, the one with the **lower** precedence will be used.
	Precedence int64 `json:"precedence" api:"required"`
	// The type of redirect to create. Options include: 301 (permanent), 302
	// (temporary), or 305 (proxy). Find more details
	// [here](https://knowledge.hubspot.com/cos-general/how-to-redirect-a-hubspot-page).
	RedirectStyle int64 `json:"redirectStyle" api:"required"`
	// The target incoming URL, path, or pattern to match for redirection.
	RoutePrefix string `json:"routePrefix" api:"required"`
	// The date and time when the URL mapping was last modified.
	Updated time.Time `json:"updated" api:"required" format:"date-time"`
	paramObj
}

func (r URLMappingParam) MarshalJSON() (data []byte, err error) {
	type shadow URLMappingParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *URLMappingParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Destination, RedirectStyle, RoutePrefix are required.
type URLMappingCreateRequestBodyParam struct {
	// The destination URL, where the target URL should be redirected if it matches the
	// `routePrefix`.
	Destination string `json:"destination" api:"required"`
	// The type of redirect to create. Options include: 301 (permanent), 302
	// (temporary), or 305 (proxy). Find more details
	// [here](https://knowledge.hubspot.com/cos-general/how-to-redirect-a-hubspot-page).
	RedirectStyle int64 `json:"redirectStyle" api:"required"`
	// The target incoming URL, path, or pattern to match for redirection.
	RoutePrefix string `json:"routePrefix" api:"required"`
	// Whether the `routePrefix` should match on the entire URL, including the domain.
	IsMatchFullURL param.Opt[bool] `json:"isMatchFullUrl,omitzero"`
	// Whether the `routePrefix` should match on the entire URL path, including the
	// query string.
	IsMatchQueryString param.Opt[bool] `json:"isMatchQueryString,omitzero"`
	// Whether the URL redirect mapping should apply only if a live page on the URL
	// isn't found. If False, the URL redirect mapping will take precedence over any
	// existing page.
	IsOnlyAfterNotFound param.Opt[bool] `json:"isOnlyAfterNotFound,omitzero"`
	// Whether the `routePrefix` should match based on pattern.
	IsPattern param.Opt[bool] `json:"isPattern,omitzero"`
	// Whether the `routePrefix` should match both HTTP and HTTPS protocols.
	IsProtocolAgnostic param.Opt[bool] `json:"isProtocolAgnostic,omitzero"`
	// Whether a trailing slash will be ignored.
	IsTrailingSlashOptional param.Opt[bool] `json:"isTrailingSlashOptional,omitzero"`
	// Used to prioritize URL redirection. If a given URL matches more than one
	// redirect, the one with the **lower** precedence will be used.
	Precedence param.Opt[int64] `json:"precedence,omitzero"`
	paramObj
}

func (r URLMappingCreateRequestBodyParam) MarshalJSON() (data []byte, err error) {
	type shadow URLMappingCreateRequestBodyParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *URLMappingCreateRequestBodyParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type URLRedirectNewParams struct {
	URLMappingCreateRequestBody URLMappingCreateRequestBodyParam
	paramObj
}

func (r URLRedirectNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.URLMappingCreateRequestBody)
}
func (r *URLRedirectNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type URLRedirectUpdateParams struct {
	URLMapping URLMappingParam
	paramObj
}

func (r URLRedirectUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.URLMapping)
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
	Archived      param.Opt[bool]      `query:"archived,omitzero" json:"-"`
	CreatedAfter  param.Opt[time.Time] `query:"createdAfter,omitzero" format:"date-time" json:"-"`
	CreatedAt     param.Opt[time.Time] `query:"createdAt,omitzero" format:"date-time" json:"-"`
	CreatedBefore param.Opt[time.Time] `query:"createdBefore,omitzero" format:"date-time" json:"-"`
	// The maximum number of results to display per page.
	Limit         param.Opt[int64]     `query:"limit,omitzero" json:"-"`
	UpdatedAfter  param.Opt[time.Time] `query:"updatedAfter,omitzero" format:"date-time" json:"-"`
	UpdatedAt     param.Opt[time.Time] `query:"updatedAt,omitzero" format:"date-time" json:"-"`
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
