// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cms_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/stainless-sdks/hubspot-sdk-go"
	"github.com/stainless-sdks/hubspot-sdk-go/cms"
	"github.com/stainless-sdks/hubspot-sdk-go/internal/testutil"
	"github.com/stainless-sdks/hubspot-sdk-go/option"
	"github.com/stainless-sdks/hubspot-sdk-go/shared"
)

func TestPageABTestNewLandingPageVariation(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := hubspotsdk.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAccessToken("My Access Token"),
	)
	_, err := client.Cms.Pages.ABTests.NewLandingPageVariation(context.TODO(), cms.PageABTestNewLandingPageVariationParams{
		AbTestCreateRequestVNext: shared.AbTestCreateRequestVNextParam{
			ContentID:     "contentId",
			VariationName: "variationName",
		},
	})
	if err != nil {
		var apierr *hubspotsdk.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPageABTestNewSitePageVariation(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := hubspotsdk.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAccessToken("My Access Token"),
	)
	_, err := client.Cms.Pages.ABTests.NewSitePageVariation(context.TODO(), cms.PageABTestNewSitePageVariationParams{
		AbTestCreateRequestVNext: shared.AbTestCreateRequestVNextParam{
			ContentID:     "contentId",
			VariationName: "variationName",
		},
	})
	if err != nil {
		var apierr *hubspotsdk.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPageABTestEndLandingPageTest(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := hubspotsdk.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAccessToken("My Access Token"),
	)
	err := client.Cms.Pages.ABTests.EndLandingPageTest(context.TODO(), cms.PageABTestEndLandingPageTestParams{
		AbTestEndRequestVNext: cms.AbTestEndRequestVNextParam{
			AbTestID: "abTestId",
			WinnerID: "winnerId",
		},
	})
	if err != nil {
		var apierr *hubspotsdk.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPageABTestEndSitePageTest(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := hubspotsdk.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAccessToken("My Access Token"),
	)
	err := client.Cms.Pages.ABTests.EndSitePageTest(context.TODO(), cms.PageABTestEndSitePageTestParams{
		AbTestEndRequestVNext: cms.AbTestEndRequestVNextParam{
			AbTestID: "abTestId",
			WinnerID: "winnerId",
		},
	})
	if err != nil {
		var apierr *hubspotsdk.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPageABTestRerunLandingPageTest(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := hubspotsdk.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAccessToken("My Access Token"),
	)
	err := client.Cms.Pages.ABTests.RerunLandingPageTest(context.TODO(), cms.PageABTestRerunLandingPageTestParams{
		AbTestRerunRequestVNext: cms.AbTestRerunRequestVNextParam{
			AbTestID:    "abTestId",
			VariationID: "variationId",
		},
	})
	if err != nil {
		var apierr *hubspotsdk.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPageABTestRerunSitePageTest(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := hubspotsdk.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAccessToken("My Access Token"),
	)
	err := client.Cms.Pages.ABTests.RerunSitePageTest(context.TODO(), cms.PageABTestRerunSitePageTestParams{
		AbTestRerunRequestVNext: cms.AbTestRerunRequestVNextParam{
			AbTestID:    "abTestId",
			VariationID: "variationId",
		},
	})
	if err != nil {
		var apierr *hubspotsdk.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
