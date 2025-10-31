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
)

func TestSiteSearchGetIndexedDataWithOptionalParams(t *testing.T) {
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := hubspotsdk.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAccessToken("pat-na1-xxxxxxxx-xxxx"),
	)
	_, err := client.Cms.SiteSearch.GetIndexedData(
		context.TODO(),
		"contentId",
		cms.SiteSearchGetIndexedDataParams{
			Type: cms.SiteSearchGetIndexedDataParamsTypeLandingPage,
		},
	)
	if err != nil {
		var apierr *hubspotsdk.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSiteSearchSearchWithOptionalParams(t *testing.T) {
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := hubspotsdk.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAccessToken("pat-na1-xxxxxxxx-xxxx"),
	)
	_, err := client.Cms.SiteSearch.Search(context.TODO(), cms.SiteSearchSearchParams{
		Autocomplete:    hubspotsdk.Bool(true),
		BoostLimit:      hubspotsdk.Float(0),
		BoostRecent:     hubspotsdk.String("boostRecent"),
		Domain:          []string{"string"},
		GroupID:         []int64{0},
		HubdbQuery:      hubspotsdk.String("hubdbQuery"),
		Language:        cms.SiteSearchSearchParamsLanguageAf,
		Length:          cms.SiteSearchSearchParamsLengthShort,
		Limit:           hubspotsdk.Int(0),
		MatchPrefix:     hubspotsdk.Bool(true),
		Offset:          hubspotsdk.Int(0),
		PathPrefix:      []string{"string"},
		PopularityBoost: hubspotsdk.Float(0),
		Property:        []string{"string"},
		Q:               hubspotsdk.String("q"),
		TableID:         hubspotsdk.Int(0),
		Type:            []string{"LANDING_PAGE"},
	})
	if err != nil {
		var apierr *hubspotsdk.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
