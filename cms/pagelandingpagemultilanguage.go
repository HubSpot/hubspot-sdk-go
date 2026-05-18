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

// PageLandingPageMultiLanguageService contains methods and other services that
// help with interacting with the hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPageLandingPageMultiLanguageService] method instead.
type PageLandingPageMultiLanguageService struct {
	options []option.RequestOption
}

// NewPageLandingPageMultiLanguageService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewPageLandingPageMultiLanguageService(opts ...option.RequestOption) (r PageLandingPageMultiLanguageService) {
	r = PageLandingPageMultiLanguageService{}
	r.options = opts
	return
}

// Attach a landing page to a multi-language group.
func (r *PageLandingPageMultiLanguageService) AttachToLangGroup(ctx context.Context, body PageLandingPageMultiLanguageAttachToLangGroupParams, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "cms/pages/2026-03/landing-pages/multi-language/attach-to-lang-group"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Create a new language variation from an existing landing page. The variation
// will be a copy of the draft state of the source page. To preview the content,
// you can
// [retrieve the draft of the source landing page](/api-reference/latest/cms/pages/landing-pages/drafts/get-landing-page-draft).
func (r *PageLandingPageMultiLanguageService) NewLanguageVariation(ctx context.Context, body PageLandingPageMultiLanguageNewLanguageVariationParams, opts ...option.RequestOption) (res *PagesPage, err error) {
	opts = slices.Concat(r.options, opts)
	path := "cms/pages/2026-03/landing-pages/multi-language/create-language-variation"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Detach a landing page from a multi-language group.
func (r *PageLandingPageMultiLanguageService) DetachFromLangGroup(ctx context.Context, body PageLandingPageMultiLanguageDetachFromLangGroupParams, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "cms/pages/2026-03/landing-pages/multi-language/detach-from-lang-group"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Set a landing page as the primary language of a multi-language group.
func (r *PageLandingPageMultiLanguageService) SetNewLangPrimary(ctx context.Context, body PageLandingPageMultiLanguageSetNewLangPrimaryParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "cms/pages/2026-03/landing-pages/multi-language/set-new-lang-primary"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, nil, opts...)
	return err
}

// Explicitly set new languages for each landing page in a multi-language group.
func (r *PageLandingPageMultiLanguageService) UpdateLanguages(ctx context.Context, body PageLandingPageMultiLanguageUpdateLanguagesParams, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "cms/pages/2026-03/landing-pages/multi-language/update-languages"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type PageLandingPageMultiLanguageAttachToLangGroupParams struct {
	AttachToLangPrimaryRequestVNext AttachToLangPrimaryRequestVNextParam
	paramObj
}

func (r PageLandingPageMultiLanguageAttachToLangGroupParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.AttachToLangPrimaryRequestVNext)
}
func (r *PageLandingPageMultiLanguageAttachToLangGroupParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PageLandingPageMultiLanguageNewLanguageVariationParams struct {
	ContentLanguageCloneRequestVNext ContentLanguageCloneRequestVNextParam
	paramObj
}

func (r PageLandingPageMultiLanguageNewLanguageVariationParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.ContentLanguageCloneRequestVNext)
}
func (r *PageLandingPageMultiLanguageNewLanguageVariationParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PageLandingPageMultiLanguageDetachFromLangGroupParams struct {
	DetachFromLangGroupRequestVNext DetachFromLangGroupRequestVNextParam
	paramObj
}

func (r PageLandingPageMultiLanguageDetachFromLangGroupParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.DetachFromLangGroupRequestVNext)
}
func (r *PageLandingPageMultiLanguageDetachFromLangGroupParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PageLandingPageMultiLanguageSetNewLangPrimaryParams struct {
	SetNewLanguagePrimaryRequestVNext SetNewLanguagePrimaryRequestVNextParam
	paramObj
}

func (r PageLandingPageMultiLanguageSetNewLangPrimaryParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.SetNewLanguagePrimaryRequestVNext)
}
func (r *PageLandingPageMultiLanguageSetNewLangPrimaryParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PageLandingPageMultiLanguageUpdateLanguagesParams struct {
	UpdateLanguagesRequestVNext UpdateLanguagesRequestVNextParam
	paramObj
}

func (r PageLandingPageMultiLanguageUpdateLanguagesParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.UpdateLanguagesRequestVNext)
}
func (r *PageLandingPageMultiLanguageUpdateLanguagesParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
