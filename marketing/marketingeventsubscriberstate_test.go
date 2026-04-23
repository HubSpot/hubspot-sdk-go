// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package marketing_test

import (
	"bytes"
	"context"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stainless-sdks/hubspot-sdk-go"
	"github.com/stainless-sdks/hubspot-sdk-go/marketing"
	"github.com/stainless-sdks/hubspot-sdk-go/option"
)

func TestMarketingEventSubscriberStateRecordByEmail(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte("abc"))
	}))
	defer server.Close()
	baseURL := server.URL
	client := hubspotsdk.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAccessToken("My Access Token"),
	)
	resp, err := client.Marketing.MarketingEvents.SubscriberState.RecordByEmail(
		context.TODO(),
		"subscriberState",
		marketing.MarketingEventSubscriberStateRecordByEmailParams{
			ExternalEventID:   "externalEventId",
			ExternalAccountID: "externalAccountId",
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
	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		var apierr *hubspotsdk.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
	if !bytes.Equal(b, []byte("abc")) {
		t.Fatalf("return value not %s: %s", "abc", b)
	}
}

func TestMarketingEventSubscriberStateRecordByID(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte("abc"))
	}))
	defer server.Close()
	baseURL := server.URL
	client := hubspotsdk.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAccessToken("My Access Token"),
	)
	resp, err := client.Marketing.MarketingEvents.SubscriberState.RecordByID(
		context.TODO(),
		"subscriberState",
		marketing.MarketingEventSubscriberStateRecordByIDParams{
			ExternalEventID:   "externalEventId",
			ExternalAccountID: "externalAccountId",
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
	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		var apierr *hubspotsdk.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
	if !bytes.Equal(b, []byte("abc")) {
		t.Fatalf("return value not %s: %s", "abc", b)
	}
}
