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

func TestSendBatchSend(t *testing.T) {
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
	err := client.Events.Send.BatchSend(context.TODO(), events.SendBatchSendParams{
		BatchedBehavioralEventHTTPCompletionRequest: events.BatchedBehavioralEventHTTPCompletionRequestParam{
			Inputs: []events.BehavioralEventHTTPCompletionRequestParam{{
				EventName: "eventName",
				Properties: map[string]string{
					"foo": "string",
				},
				Email:      hubspotsdk.String("email"),
				ObjectID:   hubspotsdk.String("objectId"),
				OccurredAt: hubspotsdk.Time(time.Now()),
				Utk:        hubspotsdk.String("utk"),
				Uuid:       hubspotsdk.String("uuid"),
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

func TestSendSendWithOptionalParams(t *testing.T) {
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
	err := client.Events.Send.Send(context.TODO(), events.SendSendParams{
		BehavioralEventHTTPCompletionRequest: events.BehavioralEventHTTPCompletionRequestParam{
			EventName: "eventName",
			Properties: map[string]string{
				"foo": "string",
			},
			Email:      hubspotsdk.String("email"),
			ObjectID:   hubspotsdk.String("objectId"),
			OccurredAt: hubspotsdk.Time(time.Now()),
			Utk:        hubspotsdk.String("utk"),
			Uuid:       hubspotsdk.String("uuid"),
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
