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

// PageMultiLanguageService contains methods and other services that help with
// interacting with the hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPageMultiLanguageService] method instead.
type PageMultiLanguageService struct {
	options []option.RequestOption
}

// NewPageMultiLanguageService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewPageMultiLanguageService(opts ...option.RequestOption) (r PageMultiLanguageService) {
	r = PageMultiLanguageService{}
	r.options = opts
	return
}

// Attach a site page to a multi-language group.
func (r *PageMultiLanguageService) AttachToLangGroup(ctx context.Context, body PageMultiLanguageAttachToLangGroupParams, opts ...option.RequestOption) (res *http.Response, err error) {
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
func (r *PageMultiLanguageService) NewLanguageVariation(ctx context.Context, body PageMultiLanguageNewLanguageVariationParams, opts ...option.RequestOption) (res *PageData, err error) {
	opts = slices.Concat(r.options, opts)
	path := "cms/pages/2026-03/site-pages/multi-language/create-language-variation"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Detach a website page from a multi-language group.
func (r *PageMultiLanguageService) DetachFromLangGroup(ctx context.Context, body PageMultiLanguageDetachFromLangGroupParams, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "cms/pages/2026-03/site-pages/multi-language/detach-from-lang-group"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Set a site page as the primary language of a multi-language group.
func (r *PageMultiLanguageService) SetNewLangPrimary(ctx context.Context, body PageMultiLanguageSetNewLangPrimaryParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "cms/pages/2026-03/site-pages/multi-language/set-new-lang-primary"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, nil, opts...)
	return err
}

// Explicitly set new languages for each site page in a multi-language group.
func (r *PageMultiLanguageService) UpdateLanguages(ctx context.Context, body PageMultiLanguageUpdateLanguagesParams, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "cms/pages/2026-03/site-pages/multi-language/update-languages"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type PageMultiLanguageAttachToLangGroupParams struct {
	AttachToLangPrimaryRequestVNext AttachToLangPrimaryRequestVNextParam
	paramObj
}

func (r PageMultiLanguageAttachToLangGroupParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.AttachToLangPrimaryRequestVNext)
}
func (r *PageMultiLanguageAttachToLangGroupParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PageMultiLanguageNewLanguageVariationParams struct {
	ContentLanguageCloneRequestVNext ContentLanguageCloneRequestVNextParam
	paramObj
}

func (r PageMultiLanguageNewLanguageVariationParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.ContentLanguageCloneRequestVNext)
}
func (r *PageMultiLanguageNewLanguageVariationParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PageMultiLanguageDetachFromLangGroupParams struct {
	DetachFromLangGroupRequestVNext DetachFromLangGroupRequestVNextParam
	paramObj
}

func (r PageMultiLanguageDetachFromLangGroupParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.DetachFromLangGroupRequestVNext)
}
func (r *PageMultiLanguageDetachFromLangGroupParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PageMultiLanguageSetNewLangPrimaryParams struct {
	SetNewLanguagePrimaryRequestVNext SetNewLanguagePrimaryRequestVNextParam
	paramObj
}

func (r PageMultiLanguageSetNewLangPrimaryParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.SetNewLanguagePrimaryRequestVNext)
}
func (r *PageMultiLanguageSetNewLangPrimaryParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PageMultiLanguageUpdateLanguagesParams struct {
	UpdateLanguagesRequestVNext UpdateLanguagesRequestVNextParam
	paramObj
}

func (r PageMultiLanguageUpdateLanguagesParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.UpdateLanguagesRequestVNext)
}
func (r *PageMultiLanguageUpdateLanguagesParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
