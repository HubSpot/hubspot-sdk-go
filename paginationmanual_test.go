// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package hubspotsdk_test

import (
	"context"
	"os"
	"testing"

	"github.com/stainless-sdks/hubspot-sdk-go"
	"github.com/stainless-sdks/hubspot-sdk-go/crm"
	"github.com/stainless-sdks/hubspot-sdk-go/internal/testutil"
	"github.com/stainless-sdks/hubspot-sdk-go/option"
)

func TestManualPagination(t *testing.T) {
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
	page, err := client.Crm.Objects.Contacts.List(context.TODO(), crm.ObjectContactListParams{
		Limit: hubspotsdk.Int(100),
	})
	if err != nil {
		t.Fatalf("err should be nil: %s", err.Error())
	}
	for _, contact := range page.Results {
		t.Logf("%+v\n", contact.ID)
	}
	// The mock server isn't going to give us real pagination
	page, err = page.GetNextPage()
	if err != nil {
		t.Fatalf("err should be nil: %s", err.Error())
	}
	if page != nil {
		for _, contact := range page.Results {
			t.Logf("%+v\n", contact.ID)
		}
	}
}
