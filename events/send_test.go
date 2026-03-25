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

func TestSendNewEventDefinitionWithOptionalParams(t *testing.T) {
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
		option.WithAccessToken("pat-na1-xxxxxxxx-xxxx"),
	)
	_, err := client.Events.Send.NewEventDefinition(context.TODO(), events.SendNewEventDefinitionParams{
		ExternalBehavioralEventTypeDefinitionEgg: events.ExternalBehavioralEventTypeDefinitionEggParam{
			IncludeDefaultProperties: true,
			Label:                    "label",
			PropertyDefinitions: []events.ExternalBehavioralEventPropertyCreateParam{{
				Label:       "label",
				Type:        "type",
				Description: hubspotsdk.String("description"),
				Name:        hubspotsdk.String("name"),
				Options: []events.OptionInputParam{{
					DisplayOrder: 0,
					Hidden:       true,
					Label:        "label",
					Value:        "value",
					Description:  hubspotsdk.String("description"),
				}},
			}},
			CustomMatchingID: events.ExternalObjectResolutionMappingRequestParam{
				PrimaryObjectRule: events.ExternalPrimaryObjectResolutionRuleParam{
					EventPropertyName:        "eventPropertyName",
					TargetObjectPropertyName: "targetObjectPropertyName",
				},
			},
			Description:   hubspotsdk.String("description"),
			Name:          hubspotsdk.String("name"),
			PrimaryObject: hubspotsdk.String("primaryObject"),
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

func TestSendNewEventDefinitionPropertyWithOptionalParams(t *testing.T) {
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
		option.WithAccessToken("pat-na1-xxxxxxxx-xxxx"),
	)
	_, err := client.Events.Send.NewEventDefinitionProperty(
		context.TODO(),
		"eventName",
		events.SendNewEventDefinitionPropertyParams{
			ExternalBehavioralEventPropertyCreate: events.ExternalBehavioralEventPropertyCreateParam{
				Label:       "label",
				Type:        "type",
				Description: hubspotsdk.String("description"),
				Name:        hubspotsdk.String("name"),
				Options: []events.OptionInputParam{{
					DisplayOrder: 0,
					Hidden:       true,
					Label:        "label",
					Value:        "value",
					Description:  hubspotsdk.String("description"),
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

func TestSendDeleteEventDefinition(t *testing.T) {
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
		option.WithAccessToken("pat-na1-xxxxxxxx-xxxx"),
	)
	err := client.Events.Send.DeleteEventDefinition(context.TODO(), "eventName")
	if err != nil {
		var apierr *hubspotsdk.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSendDeleteEventDefinitionProperty(t *testing.T) {
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
		option.WithAccessToken("pat-na1-xxxxxxxx-xxxx"),
	)
	err := client.Events.Send.DeleteEventDefinitionProperty(
		context.TODO(),
		"propertyName",
		events.SendDeleteEventDefinitionPropertyParams{
			EventName: "eventName",
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

func TestSendGetEventDefinition(t *testing.T) {
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
		option.WithAccessToken("pat-na1-xxxxxxxx-xxxx"),
	)
	_, err := client.Events.Send.GetEventDefinition(context.TODO(), "eventName")
	if err != nil {
		var apierr *hubspotsdk.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSendListEventDefinitionsWithOptionalParams(t *testing.T) {
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
		option.WithAccessToken("pat-na1-xxxxxxxx-xxxx"),
	)
	_, err := client.Events.Send.ListEventDefinitions(context.TODO(), events.SendListEventDefinitionsParams{
		After:             hubspotsdk.String("after"),
		IncludeProperties: hubspotsdk.Bool(true),
		Limit:             hubspotsdk.Int(0),
		SearchString:      hubspotsdk.String("searchString"),
		SortOrder:         hubspotsdk.String("sortOrder"),
	})
	if err != nil {
		var apierr *hubspotsdk.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSendSendEventWithOptionalParams(t *testing.T) {
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
		option.WithAccessToken("pat-na1-xxxxxxxx-xxxx"),
	)
	err := client.Events.Send.SendEvent(context.TODO(), events.SendSendEventParams{
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

func TestSendSendEventBatch(t *testing.T) {
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
		option.WithAccessToken("pat-na1-xxxxxxxx-xxxx"),
	)
	err := client.Events.Send.SendEventBatch(context.TODO(), events.SendSendEventBatchParams{
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

func TestSendUpdateEventDefinitionWithOptionalParams(t *testing.T) {
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
		option.WithAccessToken("pat-na1-xxxxxxxx-xxxx"),
	)
	_, err := client.Events.Send.UpdateEventDefinition(
		context.TODO(),
		"eventName",
		events.SendUpdateEventDefinitionParams{
			ExternalBehavioralEventTypeDefinitionPatch: events.ExternalBehavioralEventTypeDefinitionPatchParam{
				Description: hubspotsdk.String("description"),
				Label:       hubspotsdk.String("label"),
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

func TestSendUpdateEventDefinitionPropertyWithOptionalParams(t *testing.T) {
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
		option.WithAccessToken("pat-na1-xxxxxxxx-xxxx"),
	)
	_, err := client.Events.Send.UpdateEventDefinitionProperty(
		context.TODO(),
		"propertyName",
		events.SendUpdateEventDefinitionPropertyParams{
			EventName: "eventName",
			ExternalBehavioralEventPropertyDefinitionPatch: events.ExternalBehavioralEventPropertyDefinitionPatchParam{
				Description: hubspotsdk.String("description"),
				Label:       hubspotsdk.String("label"),
				Options: []events.OptionInputParam{{
					DisplayOrder: 0,
					Hidden:       true,
					Label:        "label",
					Value:        "value",
					Description:  hubspotsdk.String("description"),
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
