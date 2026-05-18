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

// PageSitePageAbTestService contains methods and other services that help with
// interacting with the hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPageSitePageAbTestService] method instead.
type PageSitePageAbTestService struct {
	options []option.RequestOption
}

// NewPageSitePageAbTestService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewPageSitePageAbTestService(opts ...option.RequestOption) (r PageSitePageAbTestService) {
	r = PageSitePageAbTestService{}
	r.options = opts
	return
}

// Create a new A/B test variation based on the information provided in the request
// body.
func (r *PageSitePageAbTestService) NewSitePageVariation(ctx context.Context, body PageSitePageAbTestNewSitePageVariationParams, opts ...option.RequestOption) (res *PagesPage, err error) {
	opts = slices.Concat(r.options, opts)
	path := "cms/pages/2026-03/site-pages/ab-test/create-variation"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// End an active A/B test and designate a winner.
func (r *PageSitePageAbTestService) EndSitePageTest(ctx context.Context, body PageSitePageAbTestEndSitePageTestParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "cms/pages/2026-03/site-pages/ab-test/end"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return err
}

// Rerun a previous A/B test.
func (r *PageSitePageAbTestService) RerunSitePageTest(ctx context.Context, body PageSitePageAbTestRerunSitePageTestParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "cms/pages/2026-03/site-pages/ab-test/rerun"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return err
}

type PageSitePageAbTestNewSitePageVariationParams struct {
	AbTestCreateRequestVNext shared.AbTestCreateRequestVNextParam
	paramObj
}

func (r PageSitePageAbTestNewSitePageVariationParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.AbTestCreateRequestVNext)
}
func (r *PageSitePageAbTestNewSitePageVariationParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PageSitePageAbTestEndSitePageTestParams struct {
	AbTestEndRequestVNext AbTestEndRequestVNextParam
	paramObj
}

func (r PageSitePageAbTestEndSitePageTestParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.AbTestEndRequestVNext)
}
func (r *PageSitePageAbTestEndSitePageTestParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PageSitePageAbTestRerunSitePageTestParams struct {
	AbTestRerunRequestVNext AbTestRerunRequestVNextParam
	paramObj
}

func (r PageSitePageAbTestRerunSitePageTestParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.AbTestRerunRequestVNext)
}
func (r *PageSitePageAbTestRerunSitePageTestParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
