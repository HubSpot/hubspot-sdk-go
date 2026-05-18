// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cms

import (
	"context"
	"net/http"
	"slices"

	"github.com/HubSpot/hubspot-sdk-go/internal/apijson"
	shimjson "github.com/HubSpot/hubspot-sdk-go/internal/encoding/json"
	"github.com/HubSpot/hubspot-sdk-go/internal/requestconfig"
	"github.com/HubSpot/hubspot-sdk-go/option"
)

// PageSitePageMultiLanguageService contains methods and other services that help
// with interacting with the hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPageSitePageMultiLanguageService] method instead.
type PageSitePageMultiLanguageService struct {
	options []option.RequestOption
}

// NewPageSitePageMultiLanguageService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewPageSitePageMultiLanguageService(opts ...option.RequestOption) (r PageSitePageMultiLanguageService) {
	r = PageSitePageMultiLanguageService{}
	r.options = opts
	return
}

// Attach a site page to a multi-language group.
func (r *PageSitePageMultiLanguageService) AttachToLangGroup(ctx context.Context, body PageSitePageMultiLanguageAttachToLangGroupParams, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "cms/pages/2026-03/site-pages/multi-language/attach-to-lang-group"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Create a new language variation from an existing website page. The variation
// will be a copy of the draft state of the source page. To preview the content,
// you can
// [retrieve the draft of the source website page](/api-reference/latest/cms/pages/website-pages/drafts/get-website-page-draft).
func (r *PageSitePageMultiLanguageService) NewLanguageVariation(ctx context.Context, body PageSitePageMultiLanguageNewLanguageVariationParams, opts ...option.RequestOption) (res *PagesPage, err error) {
	opts = slices.Concat(r.options, opts)
	path := "cms/pages/2026-03/site-pages/multi-language/create-language-variation"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Detach a website page from a multi-language group.
func (r *PageSitePageMultiLanguageService) DetachFromLangGroup(ctx context.Context, body PageSitePageMultiLanguageDetachFromLangGroupParams, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "cms/pages/2026-03/site-pages/multi-language/detach-from-lang-group"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Set a site page as the primary language of a multi-language group.
func (r *PageSitePageMultiLanguageService) SetNewLangPrimary(ctx context.Context, body PageSitePageMultiLanguageSetNewLangPrimaryParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "cms/pages/2026-03/site-pages/multi-language/set-new-lang-primary"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, nil, opts...)
	return err
}

// Explicitly set new languages for each site page in a multi-language group.
func (r *PageSitePageMultiLanguageService) UpdateLanguages(ctx context.Context, body PageSitePageMultiLanguageUpdateLanguagesParams, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "cms/pages/2026-03/site-pages/multi-language/update-languages"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type PageSitePageMultiLanguageAttachToLangGroupParams struct {
	AttachToLangPrimaryRequestVNext AttachToLangPrimaryRequestVNextParam
	paramObj
}

func (r PageSitePageMultiLanguageAttachToLangGroupParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.AttachToLangPrimaryRequestVNext)
}
func (r *PageSitePageMultiLanguageAttachToLangGroupParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PageSitePageMultiLanguageNewLanguageVariationParams struct {
	ContentLanguageCloneRequestVNext ContentLanguageCloneRequestVNextParam
	paramObj
}

func (r PageSitePageMultiLanguageNewLanguageVariationParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.ContentLanguageCloneRequestVNext)
}
func (r *PageSitePageMultiLanguageNewLanguageVariationParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PageSitePageMultiLanguageDetachFromLangGroupParams struct {
	DetachFromLangGroupRequestVNext DetachFromLangGroupRequestVNextParam
	paramObj
}

func (r PageSitePageMultiLanguageDetachFromLangGroupParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.DetachFromLangGroupRequestVNext)
}
func (r *PageSitePageMultiLanguageDetachFromLangGroupParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PageSitePageMultiLanguageSetNewLangPrimaryParams struct {
	SetNewLanguagePrimaryRequestVNext SetNewLanguagePrimaryRequestVNextParam
	paramObj
}

func (r PageSitePageMultiLanguageSetNewLangPrimaryParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.SetNewLanguagePrimaryRequestVNext)
}
func (r *PageSitePageMultiLanguageSetNewLangPrimaryParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PageSitePageMultiLanguageUpdateLanguagesParams struct {
	UpdateLanguagesRequestVNext UpdateLanguagesRequestVNextParam
	paramObj
}

func (r PageSitePageMultiLanguageUpdateLanguagesParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.UpdateLanguagesRequestVNext)
}
func (r *PageSitePageMultiLanguageUpdateLanguagesParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
