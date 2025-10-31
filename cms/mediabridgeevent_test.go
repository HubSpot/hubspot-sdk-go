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

func TestMediaBridgeEventNewAttentionSpanEventWithOptionalParams(t *testing.T) {
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
	_, err := client.Cms.MediaBridge.Events.NewAttentionSpanEvent(context.TODO(), cms.MediaBridgeEventNewAttentionSpanEventParams{
		MediaType:         cms.MediaBridgeEventNewAttentionSpanEventParamsMediaTypeVideo,
		OccurredTimestamp: 0,
		RawDataMap: map[string]int64{
			"foo": 0,
		},
		SessionID:  "sessionId",
		Hsenc:      hubspotsdk.String("_hsenc"),
		ContactID:  hubspotsdk.Int(0),
		ContactUtk: hubspotsdk.String("contactUtk"),
		DerivedValues: cms.MediaBridgeEventNewAttentionSpanEventParamsDerivedValues{
			TotalPercentPlayed: 0,
			TotalSecondsPlayed: 0,
		},
		ExternalID:    hubspotsdk.String("externalId"),
		MediaBridgeID: hubspotsdk.Int(0),
		MediaName:     hubspotsdk.String("mediaName"),
		MediaURL:      hubspotsdk.String("mediaUrl"),
		PageID:        hubspotsdk.Int(0),
		PageName:      hubspotsdk.String("pageName"),
		PageURL:       hubspotsdk.String("pageUrl"),
		RawDataString: hubspotsdk.String("rawDataString"),
	})
	if err != nil {
		var apierr *hubspotsdk.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestMediaBridgeEventNewMediaPlayedEventWithOptionalParams(t *testing.T) {
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
	_, err := client.Cms.MediaBridge.Events.NewMediaPlayedEvent(context.TODO(), cms.MediaBridgeEventNewMediaPlayedEventParams{
		MediaType:         cms.MediaBridgeEventNewMediaPlayedEventParamsMediaTypeVideo,
		OccurredTimestamp: 0,
		SessionID:         "sessionId",
		State:             cms.MediaBridgeEventNewMediaPlayedEventParamsStateStarted,
		Hsenc:             hubspotsdk.String("_hsenc"),
		ContactID:         hubspotsdk.Int(0),
		ContactUtk:        hubspotsdk.String("contactUtk"),
		ExternalID:        hubspotsdk.String("externalId"),
		IframeURL:         hubspotsdk.String("iframeUrl"),
		MediaBridgeID:     hubspotsdk.Int(0),
		MediaName:         hubspotsdk.String("mediaName"),
		MediaURL:          hubspotsdk.String("mediaUrl"),
		PageID:            hubspotsdk.Int(0),
		PageName:          hubspotsdk.String("pageName"),
		PageURL:           hubspotsdk.String("pageUrl"),
	})
	if err != nil {
		var apierr *hubspotsdk.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestMediaBridgeEventNewMediaPlayedPercentEventWithOptionalParams(t *testing.T) {
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
	_, err := client.Cms.MediaBridge.Events.NewMediaPlayedPercentEvent(context.TODO(), cms.MediaBridgeEventNewMediaPlayedPercentEventParams{
		MediaType:         cms.MediaBridgeEventNewMediaPlayedPercentEventParamsMediaTypeVideo,
		OccurredTimestamp: 0,
		PlayedPercent:     0,
		SessionID:         "sessionId",
		Hsenc:             hubspotsdk.String("_hsenc"),
		ContactID:         hubspotsdk.Int(0),
		ContactUtk:        hubspotsdk.String("contactUtk"),
		ExternalID:        hubspotsdk.String("externalId"),
		MediaBridgeID:     hubspotsdk.Int(0),
		MediaName:         hubspotsdk.String("mediaName"),
		MediaURL:          hubspotsdk.String("mediaUrl"),
		PageID:            hubspotsdk.Int(0),
		PageName:          hubspotsdk.String("pageName"),
		PageURL:           hubspotsdk.String("pageUrl"),
	})
	if err != nil {
		var apierr *hubspotsdk.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
