// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cms

import (
	"context"
	"encoding/json"
	"net/http"
	"slices"

	shimjson "github.com/stainless-sdks/hubspot-sdk-go/internal/encoding/json"
	"github.com/stainless-sdks/hubspot-sdk-go/internal/requestconfig"
	"github.com/stainless-sdks/hubspot-sdk-go/option"
)

// BlogSettingMultiLanguageService contains methods and other services that help
// with interacting with the hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBlogSettingMultiLanguageService] method instead.
type BlogSettingMultiLanguageService struct {
	Options []option.RequestOption
}

// NewBlogSettingMultiLanguageService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewBlogSettingMultiLanguageService(opts ...option.RequestOption) (r BlogSettingMultiLanguageService) {
	r = BlogSettingMultiLanguageService{}
	r.Options = opts
	return
}

func (r *BlogSettingMultiLanguageService) AttachToLangGroup(ctx context.Context, body BlogSettingMultiLanguageAttachToLangGroupParams, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "cms/blog-settings/2026-03/settings/multi-language/attach-to-lang-group"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

func (r *BlogSettingMultiLanguageService) NewLanguageVariation(ctx context.Context, body BlogSettingMultiLanguageNewLanguageVariationParams, opts ...option.RequestOption) (res *Blog, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "cms/blog-settings/2026-03/settings/multi-language/create-language-variation"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

func (r *BlogSettingMultiLanguageService) DetachFromLangGroup(ctx context.Context, body BlogSettingMultiLanguageDetachFromLangGroupParams, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "cms/blog-settings/2026-03/settings/multi-language/detach-from-lang-group"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

func (r *BlogSettingMultiLanguageService) SetNewLangPrimary(ctx context.Context, body BlogSettingMultiLanguageSetNewLangPrimaryParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "cms/blog-settings/2026-03/settings/multi-language/set-new-lang-primary"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, nil, opts...)
	return err
}

func (r *BlogSettingMultiLanguageService) UpdateLanguages(ctx context.Context, body BlogSettingMultiLanguageUpdateLanguagesParams, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "cms/blog-settings/2026-03/settings/multi-language/update-languages"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type BlogSettingMultiLanguageAttachToLangGroupParams struct {
	AttachToLangPrimaryRequestVNext AttachToLangPrimaryRequestVNextParam
	paramObj
}

func (r BlogSettingMultiLanguageAttachToLangGroupParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.AttachToLangPrimaryRequestVNext)
}
func (r *BlogSettingMultiLanguageAttachToLangGroupParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.AttachToLangPrimaryRequestVNext)
}

type BlogSettingMultiLanguageNewLanguageVariationParams struct {
	BlogLanguageCloneRequestVNext BlogLanguageCloneRequestVNextParam
	paramObj
}

func (r BlogSettingMultiLanguageNewLanguageVariationParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.BlogLanguageCloneRequestVNext)
}
func (r *BlogSettingMultiLanguageNewLanguageVariationParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.BlogLanguageCloneRequestVNext)
}

type BlogSettingMultiLanguageDetachFromLangGroupParams struct {
	DetachFromLangGroupRequestVNext DetachFromLangGroupRequestVNextParam
	paramObj
}

func (r BlogSettingMultiLanguageDetachFromLangGroupParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.DetachFromLangGroupRequestVNext)
}
func (r *BlogSettingMultiLanguageDetachFromLangGroupParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.DetachFromLangGroupRequestVNext)
}

type BlogSettingMultiLanguageSetNewLangPrimaryParams struct {
	SetNewLanguagePrimaryRequestVNext SetNewLanguagePrimaryRequestVNextParam
	paramObj
}

func (r BlogSettingMultiLanguageSetNewLangPrimaryParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.SetNewLanguagePrimaryRequestVNext)
}
func (r *BlogSettingMultiLanguageSetNewLangPrimaryParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.SetNewLanguagePrimaryRequestVNext)
}

type BlogSettingMultiLanguageUpdateLanguagesParams struct {
	UpdateLanguagesRequestVNext UpdateLanguagesRequestVNextParam
	paramObj
}

func (r BlogSettingMultiLanguageUpdateLanguagesParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.UpdateLanguagesRequestVNext)
}
func (r *BlogSettingMultiLanguageUpdateLanguagesParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.UpdateLanguagesRequestVNext)
}
