// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package marketing_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/stainless-sdks/hubspot-sdk-go"
	"github.com/stainless-sdks/hubspot-sdk-go/internal/testutil"
	"github.com/stainless-sdks/hubspot-sdk-go/marketing"
	"github.com/stainless-sdks/hubspot-sdk-go/option"
)

func TestSingleSendNew(t *testing.T) {
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
	_, err := client.Marketing.SingleSend.New(context.TODO(), marketing.SingleSendNewParams{
		PublicSingleSendRequestEgg: marketing.PublicSingleSendRequestEggParam{
			ContactProperties: map[string]string{
				"foo": "string",
			},
			CustomProperties: map[string]any{
				"foo": map[string]any{},
			},
			EmailID: 0,
			Message: marketing.PublicSingleSendEmailParam{
				Bcc:     []string{"string"},
				Cc:      []string{"string"},
				ReplyTo: []string{"string"},
				From:    hubspotsdk.String("from"),
				SendID:  hubspotsdk.String("sendId"),
				To:      hubspotsdk.String("to"),
			},
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
