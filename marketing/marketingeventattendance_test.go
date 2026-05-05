// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package marketing_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/HubSpot/hubspot-sdk-go"
	"github.com/HubSpot/hubspot-sdk-go/internal/testutil"
	"github.com/HubSpot/hubspot-sdk-go/marketing"
	"github.com/HubSpot/hubspot-sdk-go/option"
)

func TestMarketingEventAttendanceNewByEventIDAndContactID(t *testing.T) {
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
	_, err := client.Marketing.MarketingEvents.Attendance.NewByEventIDAndContactID(
		context.TODO(),
		"subscriberState",
		marketing.MarketingEventAttendanceNewByEventIDAndContactIDParams{
			ObjectID: "objectId",
			BatchInputMarketingEventSubscriber: marketing.BatchInputMarketingEventSubscriberParam{
				Inputs: []marketing.MarketingEventSubscriberParam{{
					InteractionDateTime: 0,
					Properties: map[string]string{
						"foo": "string",
					},
					Vid: 0,
				}},
			},
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

func TestMarketingEventAttendanceNewByEventIDAndEmail(t *testing.T) {
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
	_, err := client.Marketing.MarketingEvents.Attendance.NewByEventIDAndEmail(
		context.TODO(),
		"subscriberState",
		marketing.MarketingEventAttendanceNewByEventIDAndEmailParams{
			ObjectID: "objectId",
			BatchInputMarketingEventEmailSubscriber: marketing.BatchInputMarketingEventEmailSubscriberParam{
				Inputs: []marketing.MarketingEventEmailSubscriberParam{{
					ContactProperties: map[string]string{
						"foo": "string",
					},
					Email:               "email",
					InteractionDateTime: 0,
					Properties: map[string]string{
						"foo": "string",
					},
				}},
			},
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

func TestMarketingEventAttendanceNewByExternalEventIDAndContactIDWithOptionalParams(t *testing.T) {
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
	_, err := client.Marketing.MarketingEvents.Attendance.NewByExternalEventIDAndContactID(
		context.TODO(),
		"subscriberState",
		marketing.MarketingEventAttendanceNewByExternalEventIDAndContactIDParams{
			ExternalEventID: "externalEventId",
			BatchInputMarketingEventSubscriber: marketing.BatchInputMarketingEventSubscriberParam{
				Inputs: []marketing.MarketingEventSubscriberParam{{
					InteractionDateTime: 0,
					Properties: map[string]string{
						"foo": "string",
					},
					Vid: 0,
				}},
			},
			ExternalAccountID: hubspotsdk.String("externalAccountId"),
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

func TestMarketingEventAttendanceNewByExternalEventIDAndEmailWithOptionalParams(t *testing.T) {
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
	_, err := client.Marketing.MarketingEvents.Attendance.NewByExternalEventIDAndEmail(
		context.TODO(),
		"subscriberState",
		marketing.MarketingEventAttendanceNewByExternalEventIDAndEmailParams{
			ExternalEventID: "externalEventId",
			BatchInputMarketingEventEmailSubscriber: marketing.BatchInputMarketingEventEmailSubscriberParam{
				Inputs: []marketing.MarketingEventEmailSubscriberParam{{
					ContactProperties: map[string]string{
						"foo": "string",
					},
					Email:               "email",
					InteractionDateTime: 0,
					Properties: map[string]string{
						"foo": "string",
					},
				}},
			},
			ExternalAccountID: hubspotsdk.String("externalAccountId"),
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
