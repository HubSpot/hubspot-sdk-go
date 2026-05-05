// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cms_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/HubSpot/hubspot-sdk-go"
	"github.com/HubSpot/hubspot-sdk-go/cms"
	"github.com/HubSpot/hubspot-sdk-go/internal/testutil"
	"github.com/HubSpot/hubspot-sdk-go/option"
	"github.com/HubSpot/hubspot-sdk-go/shared"
)

func TestMediaBridgeNewAssociationWithOptionalParams(t *testing.T) {
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
	_, err := client.Cms.MediaBridge.NewAssociation(
		context.TODO(),
		"objectType",
		cms.MediaBridgeNewAssociationParams{
			AppID: 0,
			AssociationDefinitionEgg: shared.AssociationDefinitionEggParam{
				FromObjectTypeID: "fromObjectTypeId",
				ToObjectTypeID:   "toObjectTypeId",
				Name:             hubspotsdk.String("name"),
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

func TestMediaBridgeNewAttentionSpanEventWithOptionalParams(t *testing.T) {
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
	_, err := client.Cms.MediaBridge.NewAttentionSpanEvent(context.TODO(), cms.MediaBridgeNewAttentionSpanEventParams{
		AttentionSpanEventRequest: cms.AttentionSpanEventRequestParam{
			MediaType:         cms.AttentionSpanEventRequestMediaTypeAudio,
			OccurredTimestamp: 0,
			RawDataMap: map[string]int64{
				"foo": 0,
			},
			SessionID:  "sessionId",
			Hsenc:      hubspotsdk.String("_hsenc"),
			ContactID:  hubspotsdk.Int(0),
			ContactUtk: hubspotsdk.String("contactUtk"),
			DerivedValues: cms.AttentionSpanCalculatedValuesParam{
				TotalPercentPlayed: 0,
				TotalSecondsPlayed: 0,
			},
			ExternalID:          hubspotsdk.String("externalId"),
			ExternalPlayContext: cms.AttentionSpanEventRequestExternalPlayContextEmail,
			MediaBridgeID:       hubspotsdk.Int(0),
			MediaName:           hubspotsdk.String("mediaName"),
			MediaURL:            hubspotsdk.String("mediaUrl"),
			PageID:              hubspotsdk.Int(0),
			PageName:            hubspotsdk.String("pageName"),
			PageURL:             hubspotsdk.String("pageUrl"),
			RawDataString:       hubspotsdk.String("rawDataString"),
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

func TestMediaBridgeNewMediaPlayedEventWithOptionalParams(t *testing.T) {
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
	_, err := client.Cms.MediaBridge.NewMediaPlayedEvent(context.TODO(), cms.MediaBridgeNewMediaPlayedEventParams{
		MediaPlayedEventRequest: cms.MediaPlayedEventRequestParam{
			MediaType:           cms.MediaPlayedEventRequestMediaTypeAudio,
			OccurredTimestamp:   0,
			SessionID:           "sessionId",
			State:               cms.MediaPlayedEventRequestStateStarted,
			Hsenc:               hubspotsdk.String("_hsenc"),
			ContactID:           hubspotsdk.Int(0),
			ContactUtk:          hubspotsdk.String("contactUtk"),
			ExternalID:          hubspotsdk.String("externalId"),
			ExternalPlayContext: cms.MediaPlayedEventRequestExternalPlayContextEmail,
			IframeURL:           hubspotsdk.String("iframeUrl"),
			MediaBridgeID:       hubspotsdk.Int(0),
			MediaName:           hubspotsdk.String("mediaName"),
			MediaURL:            hubspotsdk.String("mediaUrl"),
			PageID:              hubspotsdk.Int(0),
			PageName:            hubspotsdk.String("pageName"),
			PageURL:             hubspotsdk.String("pageUrl"),
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

func TestMediaBridgeNewMediaPlayedPercentEventWithOptionalParams(t *testing.T) {
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
	_, err := client.Cms.MediaBridge.NewMediaPlayedPercentEvent(context.TODO(), cms.MediaBridgeNewMediaPlayedPercentEventParams{
		MediaPlayedPercentageEventRequest: cms.MediaPlayedPercentageEventRequestParam{
			MediaType:           cms.MediaPlayedPercentageEventRequestMediaTypeAudio,
			OccurredTimestamp:   0,
			PlayedPercent:       0,
			SessionID:           "sessionId",
			Hsenc:               hubspotsdk.String("_hsenc"),
			ContactID:           hubspotsdk.Int(0),
			ContactUtk:          hubspotsdk.String("contactUtk"),
			ExternalID:          hubspotsdk.String("externalId"),
			ExternalPlayContext: cms.MediaPlayedPercentageEventRequestExternalPlayContextEmail,
			MediaBridgeID:       hubspotsdk.Int(0),
			MediaName:           hubspotsdk.String("mediaName"),
			MediaURL:            hubspotsdk.String("mediaUrl"),
			PageID:              hubspotsdk.Int(0),
			PageName:            hubspotsdk.String("pageName"),
			PageURL:             hubspotsdk.String("pageUrl"),
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

func TestMediaBridgeNewObjectType(t *testing.T) {
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
	_, err := client.Cms.MediaBridge.NewObjectType(
		context.TODO(),
		0,
		cms.MediaBridgeNewObjectTypeParams{
			IntegratorObjectCreationRequest: cms.IntegratorObjectCreationRequestParam{
				MediaTypes: []string{"VIDEO"},
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

func TestMediaBridgeNewOembedDomainWithOptionalParams(t *testing.T) {
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
	_, err := client.Cms.MediaBridge.NewOembedDomain(
		context.TODO(),
		0,
		cms.MediaBridgeNewOembedDomainParams{
			IntegratorOEmbedDomainRequest: cms.IntegratorOEmbedDomainRequestParam{
				Endpoints: cms.EndpointsParam{
					Discovery: true,
					Schemes:   []string{"string"},
					URL:       "url",
				},
				PortalID: hubspotsdk.Int(0),
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

func TestMediaBridgeNewPropertyWithOptionalParams(t *testing.T) {
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
	_, err := client.Cms.MediaBridge.NewProperty(
		context.TODO(),
		"objectType",
		cms.MediaBridgeNewPropertyParams{
			AppID: 0,
			PropertyCreate: shared.PropertyCreateParam{
				FieldType:            shared.PropertyCreateFieldTypeBooleancheckbox,
				GroupName:            "groupName",
				Label:                "label",
				Name:                 "name",
				Type:                 shared.PropertyCreateTypeBool,
				CalculationFormula:   hubspotsdk.String("calculationFormula"),
				CurrencyPropertyName: hubspotsdk.String("currencyPropertyName"),
				DataSensitivity:      shared.PropertyCreateDataSensitivityHighlySensitive,
				Description:          hubspotsdk.String("description"),
				DisplayOrder:         hubspotsdk.Int(0),
				ExternalOptions:      hubspotsdk.Bool(true),
				FormField:            hubspotsdk.Bool(true),
				HasUniqueValue:       hubspotsdk.Bool(true),
				Hidden:               hubspotsdk.Bool(true),
				NumberDisplayHint:    shared.PropertyCreateNumberDisplayHintCurrency,
				Options: []shared.OptionInputParam{{
					DisplayOrder: 0,
					Hidden:       true,
					Label:        "label",
					Value:        "value",
					Description:  hubspotsdk.String("description"),
				}},
				ReferencedObjectType: hubspotsdk.String("referencedObjectType"),
				ShowCurrencySymbol:   hubspotsdk.Bool(true),
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

func TestMediaBridgeNewPropertyGroupWithOptionalParams(t *testing.T) {
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
	_, err := client.Cms.MediaBridge.NewPropertyGroup(
		context.TODO(),
		"objectType",
		cms.MediaBridgeNewPropertyGroupParams{
			AppID: 0,
			PropertyGroupCreate: shared.PropertyGroupCreateParam{
				Label:        "label",
				Name:         "name",
				DisplayOrder: hubspotsdk.Int(0),
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

func TestMediaBridgeNewVideoAssociationDefinition(t *testing.T) {
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
	_, err := client.Cms.MediaBridge.NewVideoAssociationDefinition(context.TODO(), 0)
	if err != nil {
		var apierr *hubspotsdk.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestMediaBridgeDeleteAssociation(t *testing.T) {
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
	err := client.Cms.MediaBridge.DeleteAssociation(
		context.TODO(),
		"associationId",
		cms.MediaBridgeDeleteAssociationParams{
			AppID:      0,
			ObjectType: "objectType",
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

func TestMediaBridgeDeleteOembedDomainWithOptionalParams(t *testing.T) {
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
	err := client.Cms.MediaBridge.DeleteOembedDomain(
		context.TODO(),
		0,
		cms.MediaBridgeDeleteOembedDomainParams{
			ID:             hubspotsdk.Int(0),
			DomainPortalID: hubspotsdk.Int(0),
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

func TestMediaBridgeDeleteProperty(t *testing.T) {
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
	err := client.Cms.MediaBridge.DeleteProperty(
		context.TODO(),
		"propertyName",
		cms.MediaBridgeDeletePropertyParams{
			AppID:      0,
			ObjectType: "objectType",
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

func TestMediaBridgeDeletePropertyGroup(t *testing.T) {
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
	err := client.Cms.MediaBridge.DeletePropertyGroup(
		context.TODO(),
		"groupName",
		cms.MediaBridgeDeletePropertyGroupParams{
			AppID:      0,
			ObjectType: "objectType",
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

func TestMediaBridgeGetEventVisibilitySettings(t *testing.T) {
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
	_, err := client.Cms.MediaBridge.GetEventVisibilitySettings(context.TODO(), 0)
	if err != nil {
		var apierr *hubspotsdk.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestMediaBridgeGetOembedDomain(t *testing.T) {
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
	_, err := client.Cms.MediaBridge.GetOembedDomain(
		context.TODO(),
		"oEmbedDomainId",
		cms.MediaBridgeGetOembedDomainParams{
			AppID: 0,
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

func TestMediaBridgeGetPropertyWithOptionalParams(t *testing.T) {
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
	_, err := client.Cms.MediaBridge.GetProperty(
		context.TODO(),
		"propertyName",
		cms.MediaBridgeGetPropertyParams{
			AppID:      0,
			ObjectType: "objectType",
			Archived:   hubspotsdk.Bool(true),
			Properties: hubspotsdk.String("properties"),
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

func TestMediaBridgeGetPropertyGroup(t *testing.T) {
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
	_, err := client.Cms.MediaBridge.GetPropertyGroup(
		context.TODO(),
		"groupName",
		cms.MediaBridgeGetPropertyGroupParams{
			AppID:      0,
			ObjectType: "objectType",
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

func TestMediaBridgeGetSchema(t *testing.T) {
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
	_, err := client.Cms.MediaBridge.GetSchema(
		context.TODO(),
		"objectType",
		cms.MediaBridgeGetSchemaParams{
			AppID: 0,
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

func TestMediaBridgeListObjectTypesByMediaTypeWithOptionalParams(t *testing.T) {
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
	_, err := client.Cms.MediaBridge.ListObjectTypesByMediaType(
		context.TODO(),
		cms.MediaBridgeListObjectTypesByMediaTypeParamsMediaTypeAudio,
		cms.MediaBridgeListObjectTypesByMediaTypeParams{
			AppID:                 0,
			IncludeFullDefinition: hubspotsdk.Bool(true),
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

func TestMediaBridgeListOembedDomainsWithOptionalParams(t *testing.T) {
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
	_, err := client.Cms.MediaBridge.ListOembedDomains(
		context.TODO(),
		0,
		cms.MediaBridgeListOembedDomainsParams{
			DomainPortalID: hubspotsdk.Int(0),
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

func TestMediaBridgeListPropertiesWithOptionalParams(t *testing.T) {
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
	_, err := client.Cms.MediaBridge.ListProperties(
		context.TODO(),
		"objectType",
		cms.MediaBridgeListPropertiesParams{
			AppID:      0,
			Archived:   hubspotsdk.Bool(true),
			Properties: hubspotsdk.String("properties"),
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

func TestMediaBridgeListPropertyGroups(t *testing.T) {
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
	_, err := client.Cms.MediaBridge.ListPropertyGroups(
		context.TODO(),
		"objectType",
		cms.MediaBridgeListPropertyGroupsParams{
			AppID: 0,
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

func TestMediaBridgeListSchemasWithOptionalParams(t *testing.T) {
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
	_, err := client.Cms.MediaBridge.ListSchemas(
		context.TODO(),
		0,
		cms.MediaBridgeListSchemasParams{
			Archived: hubspotsdk.Bool(true),
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

func TestMediaBridgeRegisterAppNameWithOptionalParams(t *testing.T) {
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
	_, err := client.Cms.MediaBridge.RegisterAppName(
		context.TODO(),
		0,
		cms.MediaBridgeRegisterAppNameParams{
			MediaBridgeProviderPartial: cms.MediaBridgeProviderPartialParam{
				UpdatedAt:               0,
				AllowImportOnDisconnect: hubspotsdk.Bool(true),
				ModuleName:              hubspotsdk.String("moduleName"),
				Name:                    hubspotsdk.String("name"),
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

func TestMediaBridgeUpdateEventVisibilitySettingsWithOptionalParams(t *testing.T) {
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
	_, err := client.Cms.MediaBridge.UpdateEventVisibilitySettings(
		context.TODO(),
		0,
		cms.MediaBridgeUpdateEventVisibilitySettingsParams{
			EventVisibilityChange: cms.EventVisibilityChangeParam{
				EventType:       cms.EventVisibilityChangeEventTypeAll,
				UpdatedAt:       0,
				ShowInReporting: hubspotsdk.Bool(true),
				ShowInTimeline:  hubspotsdk.Bool(true),
				ShowInWorkflows: hubspotsdk.Bool(true),
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

func TestMediaBridgeUpdateOembedDomainWithOptionalParams(t *testing.T) {
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
	_, err := client.Cms.MediaBridge.UpdateOembedDomain(
		context.TODO(),
		"oEmbedDomainId",
		cms.MediaBridgeUpdateOembedDomainParams{
			AppID: 0,
			IntegratorOEmbedDomainRequest: cms.IntegratorOEmbedDomainRequestParam{
				Endpoints: cms.EndpointsParam{
					Discovery: true,
					Schemes:   []string{"string"},
					URL:       "url",
				},
				PortalID: hubspotsdk.Int(0),
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

func TestMediaBridgeUpdatePropertyWithOptionalParams(t *testing.T) {
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
	_, err := client.Cms.MediaBridge.UpdateProperty(
		context.TODO(),
		"propertyName",
		cms.MediaBridgeUpdatePropertyParams{
			AppID:      0,
			ObjectType: "objectType",
			MediaBridgePropertyUpdate: cms.MediaBridgePropertyUpdateParam{
				CalculationFormula:   hubspotsdk.String("calculationFormula"),
				CurrencyPropertyName: hubspotsdk.String("currencyPropertyName"),
				Description:          hubspotsdk.String("description"),
				DisplayOrder:         hubspotsdk.Int(0),
				FieldType:            cms.MediaBridgePropertyUpdateFieldTypeBooleancheckbox,
				FormField:            hubspotsdk.Bool(true),
				GroupName:            hubspotsdk.String("groupName"),
				HasUniqueValue:       hubspotsdk.Bool(true),
				Hidden:               hubspotsdk.Bool(true),
				Label:                hubspotsdk.String("label"),
				NumberDisplayHint:    cms.MediaBridgePropertyUpdateNumberDisplayHintCurrency,
				Options: []shared.OptionInputParam{{
					DisplayOrder: 0,
					Hidden:       true,
					Label:        "label",
					Value:        "value",
					Description:  hubspotsdk.String("description"),
				}},
				ShowCurrencySymbol: hubspotsdk.Bool(true),
				Type:               cms.MediaBridgePropertyUpdateTypeBool,
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

func TestMediaBridgeUpdatePropertyGroupWithOptionalParams(t *testing.T) {
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
	_, err := client.Cms.MediaBridge.UpdatePropertyGroup(
		context.TODO(),
		"groupName",
		cms.MediaBridgeUpdatePropertyGroupParams{
			AppID:      0,
			ObjectType: "objectType",
			PropertyGroupUpdate: shared.PropertyGroupUpdateParam{
				DisplayOrder: hubspotsdk.Int(0),
				Label:        hubspotsdk.String("label"),
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

func TestMediaBridgeUpdateSchemaWithOptionalParams(t *testing.T) {
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
	_, err := client.Cms.MediaBridge.UpdateSchema(
		context.TODO(),
		"objectType",
		cms.MediaBridgeUpdateSchemaParams{
			AppID: 0,
			ObjectTypeDefinitionPatch: shared.ObjectTypeDefinitionPatchParam{
				ClearDescription:          true,
				AllowsSensitiveProperties: hubspotsdk.Bool(true),
				Description:               hubspotsdk.String("description"),
				Labels: shared.ObjectTypeDefinitionLabelsParam{
					Plural:   hubspotsdk.String("plural"),
					Singular: hubspotsdk.String("singular"),
				},
				PrimaryDisplayProperty:     hubspotsdk.String("primaryDisplayProperty"),
				RequiredProperties:         []string{"string"},
				Restorable:                 hubspotsdk.Bool(true),
				SearchableProperties:       []string{"string"},
				SecondaryDisplayProperties: []string{"string"},
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

func TestMediaBridgeUpdateSettingsWithOptionalParams(t *testing.T) {
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
	_, err := client.Cms.MediaBridge.UpdateSettings(
		context.TODO(),
		0,
		cms.MediaBridgeUpdateSettingsParams{
			MediaBridgeProviderPartial: cms.MediaBridgeProviderPartialParam{
				UpdatedAt:               0,
				AllowImportOnDisconnect: hubspotsdk.Bool(true),
				ModuleName:              hubspotsdk.String("moduleName"),
				Name:                    hubspotsdk.String("name"),
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
