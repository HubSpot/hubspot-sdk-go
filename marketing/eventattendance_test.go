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

func TestEventAttendanceNewByEventIDAndContactID(t *testing.T) {
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
	_, err := client.Marketing.Events.Attendance.NewByEventIDAndContactID(
		context.TODO(),
		"subscriberState",
		marketing.EventAttendanceNewByEventIDAndContactIDParams{
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

func TestEventAttendanceNewByEventIDAndEmail(t *testing.T) {
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
	_, err := client.Marketing.Events.Attendance.NewByEventIDAndEmail(
		context.TODO(),
		"subscriberState",
		marketing.EventAttendanceNewByEventIDAndEmailParams{
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

func TestEventAttendanceNewByExternalEventIDAndContactIDWithOptionalParams(t *testing.T) {
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
	_, err := client.Marketing.Events.Attendance.NewByExternalEventIDAndContactID(
		context.TODO(),
		"subscriberState",
		marketing.EventAttendanceNewByExternalEventIDAndContactIDParams{
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

func TestEventAttendanceNewByExternalEventIDAndEmailWithOptionalParams(t *testing.T) {
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
	_, err := client.Marketing.Events.Attendance.NewByExternalEventIDAndEmail(
		context.TODO(),
		"subscriberState",
		marketing.EventAttendanceNewByExternalEventIDAndEmailParams{
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
