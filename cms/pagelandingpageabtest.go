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
	"github.com/HubSpot/hubspot-sdk-go/shared"
)

// PageLandingPageAbTestService contains methods and other services that help with
// interacting with the hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPageLandingPageAbTestService] method instead.
type PageLandingPageAbTestService struct {
	options []option.RequestOption
}

// NewPageLandingPageAbTestService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewPageLandingPageAbTestService(opts ...option.RequestOption) (r PageLandingPageAbTestService) {
	r = PageLandingPageAbTestService{}
	r.options = opts
	return
}

// Create a new A/B test variation based on the information provided in the request
// body.
func (r *PageLandingPageAbTestService) NewLandingPageVariation(ctx context.Context, body PageLandingPageAbTestNewLandingPageVariationParams, opts ...option.RequestOption) (res *PagesPage, err error) {
	opts = slices.Concat(r.options, opts)
	path := "cms/pages/2026-03/landing-pages/ab-test/create-variation"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// End an active A/B test and designate a winner.
func (r *PageLandingPageAbTestService) EndLandingPageTest(ctx context.Context, body PageLandingPageAbTestEndLandingPageTestParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "cms/pages/2026-03/landing-pages/ab-test/end"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return err
}

// Rerun a previous A/B test.
func (r *PageLandingPageAbTestService) RerunLandingPageTest(ctx context.Context, body PageLandingPageAbTestRerunLandingPageTestParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "cms/pages/2026-03/landing-pages/ab-test/rerun"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return err
}

type PageLandingPageAbTestNewLandingPageVariationParams struct {
	AbTestCreateRequestVNext shared.AbTestCreateRequestVNextParam
	paramObj
}

func (r PageLandingPageAbTestNewLandingPageVariationParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.AbTestCreateRequestVNext)
}
func (r *PageLandingPageAbTestNewLandingPageVariationParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PageLandingPageAbTestEndLandingPageTestParams struct {
	AbTestEndRequestVNext AbTestEndRequestVNextParam
	paramObj
}

func (r PageLandingPageAbTestEndLandingPageTestParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.AbTestEndRequestVNext)
}
func (r *PageLandingPageAbTestEndLandingPageTestParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PageLandingPageAbTestRerunLandingPageTestParams struct {
	AbTestRerunRequestVNext AbTestRerunRequestVNextParam
	paramObj
}

func (r PageLandingPageAbTestRerunLandingPageTestParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.AbTestRerunRequestVNext)
}
func (r *PageLandingPageAbTestRerunLandingPageTestParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
