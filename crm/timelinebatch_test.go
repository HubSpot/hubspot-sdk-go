// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package crm_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/HubSpot/hubspot-sdk-go"
	"github.com/HubSpot/hubspot-sdk-go/crm"
	"github.com/HubSpot/hubspot-sdk-go/internal/testutil"
	"github.com/HubSpot/hubspot-sdk-go/option"
)

func TestTimelineBatchNew(t *testing.T) {
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
	_, err := client.Crm.Timeline.Batch.New(context.TODO(), crm.TimelineBatchNewParams{
		BatchInputAppEventOccurrence: crm.BatchInputAppEventOccurrenceParam{
			Inputs: []crm.AppEventOccurrenceParam{{
				ID:            "id",
				EventTypeName: "eventTypeName",
				Properties: map[string]string{
					"foo": "string",
				},
				Domain:                       hubspotsdk.String("domain"),
				Email:                        hubspotsdk.String("email"),
				ExtraData:                    map[string]any{},
				ObjectID:                     hubspotsdk.String("objectId"),
				ObjectTypeFullyQualifiedName: hubspotsdk.String("objectTypeFullyQualifiedName"),
				TimelineIFrame: crm.TimelineEventIFrameParam{
					HeaderLabel: "headerLabel",
					Height:      0,
					LinkLabel:   "linkLabel",
					URL:         "url",
					Width:       0,
				},
				Timestamp: hubspotsdk.Time(time.Now()),
				Utk:       hubspotsdk.String("utk"),
			}},
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
