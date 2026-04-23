// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cms

import (
	"context"
	"net/http"
	"slices"

	"github.com/stainless-sdks/hubspot-sdk-go/internal/apijson"
	shimjson "github.com/stainless-sdks/hubspot-sdk-go/internal/encoding/json"
	"github.com/stainless-sdks/hubspot-sdk-go/internal/requestconfig"
	"github.com/stainless-sdks/hubspot-sdk-go/option"
)

// BlogPostMultiLanguageService contains methods and other services that help with
// interacting with the hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBlogPostMultiLanguageService] method instead.
type BlogPostMultiLanguageService struct {
	options []option.RequestOption
}

// NewBlogPostMultiLanguageService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewBlogPostMultiLanguageService(opts ...option.RequestOption) (r BlogPostMultiLanguageService) {
	r = BlogPostMultiLanguageService{}
	r.options = opts
	return
}

// Attach a blog post to a
// [multi-language group](https://developers.hubspot.com/docs/guides/cms/content/multi-language-content).
func (r *BlogPostMultiLanguageService) AttachToLangGroup(ctx context.Context, body BlogPostMultiLanguageAttachToLangGroupParams, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "cms/blogs/2026-03/posts/multi-language/attach-to-lang-group"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Create a new language variation from an existing blog post
func (r *BlogPostMultiLanguageService) NewLangVariation(ctx context.Context, body BlogPostMultiLanguageNewLangVariationParams, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "cms/blogs/2026-03/posts/multi-language/create-language-variation"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Detach a blog post from a
// [multi-language group](https://developers.hubspot.com/docs/guides/cms/content/multi-language-content).
func (r *BlogPostMultiLanguageService) DetachFromLangGroup(ctx context.Context, body BlogPostMultiLanguageDetachFromLangGroupParams, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "cms/blogs/2026-03/posts/multi-language/detach-from-lang-group"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Set the primary language of a
// [multi-language group](https://developers.hubspot.com/docs/guides/cms/content/multi-language-content)
// to the language of the provided post (specified as an ID in the request body)
func (r *BlogPostMultiLanguageService) SetLangPrimary(ctx context.Context, body BlogPostMultiLanguageSetLangPrimaryParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "cms/blogs/2026-03/posts/multi-language/set-new-lang-primary"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, nil, opts...)
	return err
}

// Explicitly set new languages for each post in a
// [multi-language group](https://developers.hubspot.com/docs/guides/cms/content/multi-language-content).
func (r *BlogPostMultiLanguageService) UpdateLangs(ctx context.Context, body BlogPostMultiLanguageUpdateLangsParams, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "cms/blogs/2026-03/posts/multi-language/update-languages"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type BlogPostMultiLanguageAttachToLangGroupParams struct {
	AttachToLangPrimaryRequestVNext AttachToLangPrimaryRequestVNextParam
	paramObj
}

func (r BlogPostMultiLanguageAttachToLangGroupParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.AttachToLangPrimaryRequestVNext)
}
func (r *BlogPostMultiLanguageAttachToLangGroupParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BlogPostMultiLanguageNewLangVariationParams struct {
	BlogPostLanguageCloneRequestVNext BlogPostLanguageCloneRequestVNextParam
	paramObj
}

func (r BlogPostMultiLanguageNewLangVariationParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.BlogPostLanguageCloneRequestVNext)
}
func (r *BlogPostMultiLanguageNewLangVariationParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BlogPostMultiLanguageDetachFromLangGroupParams struct {
	DetachFromLangGroupRequestVNext DetachFromLangGroupRequestVNextParam
	paramObj
}

func (r BlogPostMultiLanguageDetachFromLangGroupParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.DetachFromLangGroupRequestVNext)
}
func (r *BlogPostMultiLanguageDetachFromLangGroupParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BlogPostMultiLanguageSetLangPrimaryParams struct {
	SetNewLanguagePrimaryRequestVNext SetNewLanguagePrimaryRequestVNextParam
	paramObj
}

func (r BlogPostMultiLanguageSetLangPrimaryParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.SetNewLanguagePrimaryRequestVNext)
}
func (r *BlogPostMultiLanguageSetLangPrimaryParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BlogPostMultiLanguageUpdateLangsParams struct {
	UpdateLanguagesRequestVNext UpdateLanguagesRequestVNextParam
	paramObj
}

func (r BlogPostMultiLanguageUpdateLangsParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.UpdateLanguagesRequestVNext)
}
func (r *BlogPostMultiLanguageUpdateLangsParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
