// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package events_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/stainless-sdks/hubspot-sdk-go"
	"github.com/stainless-sdks/hubspot-sdk-go/events"
	"github.com/stainless-sdks/hubspot-sdk-go/internal/testutil"
	"github.com/stainless-sdks/hubspot-sdk-go/option"
)

func TestOccurrenceListWithOptionalParams(t *testing.T) {
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
	_, err := client.Events.Occurrences.List(context.TODO(), events.OccurrenceListParams{
		ID:        []string{"string"},
		After:     hubspotsdk.String("after"),
		Before:    hubspotsdk.String("before"),
		EventType: hubspotsdk.String("eventType"),
		Limit:     hubspotsdk.Int(0),
		ObjectID:  hubspotsdk.Int(0),
		ObjectProperty: events.OccurrenceListParamsObjectProperty{
			Propname: map[string]any{},
		},
		ObjectType:     hubspotsdk.String("objectType"),
		OccurredAfter:  hubspotsdk.Time(time.Now()),
		OccurredBefore: hubspotsdk.Time(time.Now()),
		Properties:     []string{"string"},
		Property: events.OccurrenceListParamsProperty{
			Propname: map[string]any{},
		},
		Sort: []string{"string"},
	})
	if err != nil {
		var apierr *hubspotsdk.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestOccurrenceListEventTypes(t *testing.T) {
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
	_, err := client.Events.Occurrences.ListEventTypes(context.TODO())
	if err != nil {
		var apierr *hubspotsdk.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
