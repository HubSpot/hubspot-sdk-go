// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package events_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/HubSpot/hubspot-sdk-go"
	"github.com/HubSpot/hubspot-sdk-go/events"
	"github.com/HubSpot/hubspot-sdk-go/internal/testutil"
	"github.com/HubSpot/hubspot-sdk-go/option"
	"github.com/HubSpot/hubspot-sdk-go/shared"
)

func TestDefinitionNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Events.Definitions.New(context.TODO(), events.DefinitionNewParams{
		ExternalBehavioralEventTypeDefinitionEgg: events.ExternalBehavioralEventTypeDefinitionEggParam{
			IncludeDefaultProperties: true,
			Label:                    "label",
			PropertyDefinitions: []events.ExternalBehavioralEventPropertyCreateParam{{
				Label:       "label",
				Type:        "type",
				Description: hubspotsdk.String("description"),
				Name:        hubspotsdk.String("name"),
				Options: []shared.OptionInputParam{{
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

func TestDefinitionUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Events.Definitions.Update(
		context.TODO(),
		"eventName",
		events.DefinitionUpdateParams{
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

func TestDefinitionListWithOptionalParams(t *testing.T) {
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
	_, err := client.Events.Definitions.List(context.TODO(), events.DefinitionListParams{
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

func TestDefinitionDelete(t *testing.T) {
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
	err := client.Events.Definitions.Delete(context.TODO(), "eventName")
	if err != nil {
		var apierr *hubspotsdk.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestDefinitionNewPropertyWithOptionalParams(t *testing.T) {
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
	_, err := client.Events.Definitions.NewProperty(
		context.TODO(),
		"eventName",
		events.DefinitionNewPropertyParams{
			ExternalBehavioralEventPropertyCreate: events.ExternalBehavioralEventPropertyCreateParam{
				Label:       "label",
				Type:        "type",
				Description: hubspotsdk.String("description"),
				Name:        hubspotsdk.String("name"),
				Options: []shared.OptionInputParam{{
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

func TestDefinitionDeleteProperty(t *testing.T) {
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
	err := client.Events.Definitions.DeleteProperty(
		context.TODO(),
		"propertyName",
		events.DefinitionDeletePropertyParams{
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

func TestDefinitionGet(t *testing.T) {
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
	_, err := client.Events.Definitions.Get(context.TODO(), "eventName")
	if err != nil {
		var apierr *hubspotsdk.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestDefinitionSendBatch(t *testing.T) {
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
	err := client.Events.Definitions.SendBatch(context.TODO(), events.DefinitionSendBatchParams{
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

func TestDefinitionUpdatePropertyWithOptionalParams(t *testing.T) {
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
	_, err := client.Events.Definitions.UpdateProperty(
		context.TODO(),
		"propertyName",
		events.DefinitionUpdatePropertyParams{
			EventName: "eventName",
			ExternalBehavioralEventPropertyDefinitionPatch: events.ExternalBehavioralEventPropertyDefinitionPatchParam{
				Description: hubspotsdk.String("description"),
				Label:       hubspotsdk.String("label"),
				Options: []shared.OptionInputParam{{
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
