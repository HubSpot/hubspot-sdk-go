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
	"github.com/stainless-sdks/hubspot-sdk-go/shared"
)

// PageABTestService contains methods and other services that help with interacting
// with the hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPageABTestService] method instead.
type PageABTestService struct {
	options []option.RequestOption
}

// NewPageABTestService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewPageABTestService(opts ...option.RequestOption) (r PageABTestService) {
	r = PageABTestService{}
	r.options = opts
	return
}

// Create a new A/B test variation based on the information provided in the request
// body.
func (r *PageABTestService) NewLandingPageVariation(ctx context.Context, body PageABTestNewLandingPageVariationParams, opts ...option.RequestOption) (res *PageData, err error) {
	opts = slices.Concat(r.options, opts)
	path := "cms/pages/2026-03/landing-pages/ab-test/create-variation"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Create a new A/B test variation based on the information provided in the request
// body.
func (r *PageABTestService) NewSitePageVariation(ctx context.Context, body PageABTestNewSitePageVariationParams, opts ...option.RequestOption) (res *PageData, err error) {
	opts = slices.Concat(r.options, opts)
	path := "cms/pages/2026-03/site-pages/ab-test/create-variation"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// End an active A/B test and designate a winner.
func (r *PageABTestService) EndLandingPageTest(ctx context.Context, body PageABTestEndLandingPageTestParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "cms/pages/2026-03/landing-pages/ab-test/end"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return err
}

// End an active A/B test and designate a winner.
func (r *PageABTestService) EndSitePageTest(ctx context.Context, body PageABTestEndSitePageTestParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "cms/pages/2026-03/site-pages/ab-test/end"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return err
}

// Rerun a previous A/B test.
func (r *PageABTestService) RerunLandingPageTest(ctx context.Context, body PageABTestRerunLandingPageTestParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "cms/pages/2026-03/landing-pages/ab-test/rerun"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return err
}

// Rerun a previous A/B test.
func (r *PageABTestService) RerunSitePageTest(ctx context.Context, body PageABTestRerunSitePageTestParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "cms/pages/2026-03/site-pages/ab-test/rerun"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return err
}

type PageABTestNewLandingPageVariationParams struct {
	AbTestCreateRequestVNext shared.AbTestCreateRequestVNextParam
	paramObj
}

func (r PageABTestNewLandingPageVariationParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.AbTestCreateRequestVNext)
}
func (r *PageABTestNewLandingPageVariationParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PageABTestNewSitePageVariationParams struct {
	AbTestCreateRequestVNext shared.AbTestCreateRequestVNextParam
	paramObj
}

func (r PageABTestNewSitePageVariationParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.AbTestCreateRequestVNext)
}
func (r *PageABTestNewSitePageVariationParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PageABTestEndLandingPageTestParams struct {
	AbTestEndRequestVNext AbTestEndRequestVNextParam
	paramObj
}

func (r PageABTestEndLandingPageTestParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.AbTestEndRequestVNext)
}
func (r *PageABTestEndLandingPageTestParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PageABTestEndSitePageTestParams struct {
	AbTestEndRequestVNext AbTestEndRequestVNextParam
	paramObj
}

func (r PageABTestEndSitePageTestParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.AbTestEndRequestVNext)
}
func (r *PageABTestEndSitePageTestParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PageABTestRerunLandingPageTestParams struct {
	AbTestRerunRequestVNext AbTestRerunRequestVNextParam
	paramObj
}

func (r PageABTestRerunLandingPageTestParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.AbTestRerunRequestVNext)
}
func (r *PageABTestRerunLandingPageTestParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PageABTestRerunSitePageTestParams struct {
	AbTestRerunRequestVNext AbTestRerunRequestVNextParam
	paramObj
}

func (r PageABTestRerunSitePageTestParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.AbTestRerunRequestVNext)
}
func (r *PageABTestRerunSitePageTestParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
