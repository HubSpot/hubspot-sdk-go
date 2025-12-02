// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package crm_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/stainless-sdks/hubspot-sdk-go"
	"github.com/stainless-sdks/hubspot-sdk-go/crm"
	"github.com/stainless-sdks/hubspot-sdk-go/internal/testutil"
	"github.com/stainless-sdks/hubspot-sdk-go/option"
)

func TestTimelineEventNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Crm.Timeline.Events.New(context.TODO(), crm.TimelineEventNewParams{
		TimelineEvent: crm.TimelineEventParam{
			EventTemplateID: "1001298",
			Tokens: map[string]string{
				"petAge":   "string",
				"petColor": "black",
				"petName":  "Art3mis",
			},
			ID:     hubspotsdk.String("id"),
			Domain: hubspotsdk.String("domain"),
			Email:  hubspotsdk.String("art3mis-pup@petspot.com"),
			ExtraData: map[string]any{
				"questions": []undefined{
					map[string]any{
						"answer":   "Bark!",
						"question": "Who's a good girl?",
					},
					map[string]any{
						"answer":   "Woof!",
						"question": "Do you wanna go on a walk?",
					},
				},
			},
			ObjectID: hubspotsdk.String("objectId"),
			TimelineIFrame: crm.TimelineEventIFrameParam{
				HeaderLabel: "Art3mis dog",
				Height:      400,
				LinkLabel:   "View Art3mis",
				URL:         "https://my.petspot.com/pets/Art3mis",
				Width:       600,
			},
			Timestamp: hubspotsdk.Time(time.Now()),
			Utk:       hubspotsdk.String("utk"),
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

func TestTimelineEventBatchNew(t *testing.T) {
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
	err := client.Crm.Timeline.Events.BatchNew(context.TODO(), crm.TimelineEventBatchNewParams{
		BatchInputTimelineEvent: crm.BatchInputTimelineEventParam{
			Inputs: []crm.TimelineEventParam{{
				EventTemplateID: "1001298",
				Tokens: map[string]string{
					"petAge":   "string",
					"petColor": "black",
					"petName":  "Art3mis",
				},
				ID:     hubspotsdk.String("id"),
				Domain: hubspotsdk.String("domain"),
				Email:  hubspotsdk.String("art3mis-pup@petspot.com"),
				ExtraData: map[string]any{
					"questions": []undefined{
						map[string]any{
							"answer":   "Bark!",
							"question": "Who's a good girl?",
						},
						map[string]any{
							"answer":   "Woof!",
							"question": "Do you wanna go on a walk?",
						},
					},
				},
				ObjectID: hubspotsdk.String("objectId"),
				TimelineIFrame: crm.TimelineEventIFrameParam{
					HeaderLabel: "Art3mis dog",
					Height:      400,
					LinkLabel:   "View Art3mis",
					URL:         "https://my.petspot.com/pets/Art3mis",
					Width:       600,
				},
				Timestamp: hubspotsdk.Time(time.Now()),
				Utk:       hubspotsdk.String("utk"),
			}, {
				EventTemplateID: "1001298",
				Tokens: map[string]string{
					"petAge":   "string",
					"petColor": "yellow",
					"petName":  "Pocket",
				},
				ID:     hubspotsdk.String("id"),
				Domain: hubspotsdk.String("domain"),
				Email:  hubspotsdk.String("pocket-tiger@petspot.com"),
				ExtraData: map[string]any{
					"questions": []undefined{
						map[string]any{
							"answer":   "Purr...",
							"question": "Who's a good kitty?",
						},
						map[string]any{
							"answer":   "Meow!",
							"question": "Will you stop playing with that?",
						},
					},
				},
				ObjectID: hubspotsdk.String("objectId"),
				TimelineIFrame: crm.TimelineEventIFrameParam{
					HeaderLabel: "Pocket Tiger",
					Height:      400,
					LinkLabel:   "View Pocket",
					URL:         "https://my.petspot.com/pets/Pocket",
					Width:       600,
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

func TestTimelineEventGet(t *testing.T) {
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
	_, err := client.Crm.Timeline.Events.Get(
		context.TODO(),
		"eventId",
		crm.TimelineEventGetParams{
			EventTemplateID: "eventTemplateId",
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

func TestTimelineEventGetDetail(t *testing.T) {
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
	_, err := client.Crm.Timeline.Events.GetDetail(
		context.TODO(),
		"eventId",
		crm.TimelineEventGetDetailParams{
			EventTemplateID: "eventTemplateId",
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
