// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package crm_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/HubSpot/hubspot-sdk-go"
	"github.com/HubSpot/hubspot-sdk-go/crm"
	"github.com/HubSpot/hubspot-sdk-go/internal/testutil"
	"github.com/HubSpot/hubspot-sdk-go/option"
)

func TestExtensionCardsDevNew(t *testing.T) {
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
	_, err := client.Crm.Extensions.CardsDev.New(
		context.TODO(),
		0,
		crm.ExtensionCardsDevNewParams{
			CardCreateRequest: crm.CardCreateRequestParam{
				Actions: crm.CardActionsParam{
					BaseURLs: []string{"string"},
				},
				Display: crm.CardDisplayBodyParam{
					Properties: []crm.CardDisplayPropertyParam{{
						DataType: crm.CardDisplayPropertyDataTypeBoolean,
						Label:    "label",
						Name:     "name",
						Options: []crm.DisplayOptionParam{{
							Label: "label",
							Name:  "name",
							Type:  crm.DisplayOptionTypeDanger,
						}},
					}},
				},
				Fetch: crm.CardFetchBodyParam{
					CardType: crm.CardFetchBodyCardTypeExternal,
					ObjectTypes: []crm.CardObjectTypeBodyParam{{
						Name:             crm.CardObjectTypeBodyNameCompanies,
						PropertiesToSend: []string{"string"},
					}},
					TargetURL:          "targetUrl",
					ServerlessFunction: hubspotsdk.String("serverlessFunction"),
				},
				Title: "title",
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

func TestExtensionCardsDevUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Crm.Extensions.CardsDev.Update(
		context.TODO(),
		"cardId",
		crm.ExtensionCardsDevUpdateParams{
			AppID: 0,
			CardPatchRequest: crm.CardPatchRequestParam{
				Actions: crm.CardActionsParam{
					BaseURLs: []string{"string"},
				},
				Display: crm.CardDisplayBodyParam{
					Properties: []crm.CardDisplayPropertyParam{{
						DataType: crm.CardDisplayPropertyDataTypeBoolean,
						Label:    "label",
						Name:     "name",
						Options: []crm.DisplayOptionParam{{
							Label: "label",
							Name:  "name",
							Type:  crm.DisplayOptionTypeDanger,
						}},
					}},
				},
				Fetch: crm.CardFetchBodyPatchParam{
					ObjectTypes: []crm.CardObjectTypeBodyParam{{
						Name:             crm.CardObjectTypeBodyNameCompanies,
						PropertiesToSend: []string{"string"},
					}},
					CardType:           crm.CardFetchBodyPatchCardTypeExternal,
					ServerlessFunction: hubspotsdk.String("serverlessFunction"),
					TargetURL:          hubspotsdk.String("targetUrl"),
				},
				Title: hubspotsdk.String("title"),
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

func TestExtensionCardsDevDelete(t *testing.T) {
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
	err := client.Crm.Extensions.CardsDev.Delete(
		context.TODO(),
		"cardId",
		crm.ExtensionCardsDevDeleteParams{
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

func TestExtensionCardsDevGet(t *testing.T) {
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
	_, err := client.Crm.Extensions.CardsDev.Get(context.TODO(), 0)
	if err != nil {
		var apierr *hubspotsdk.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestExtensionCardsDevGetByID(t *testing.T) {
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
	_, err := client.Crm.Extensions.CardsDev.GetByID(
		context.TODO(),
		"cardId",
		crm.ExtensionCardsDevGetByIDParams{
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

func TestExtensionCardsDevGetSampleResponse(t *testing.T) {
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
	_, err := client.Crm.Extensions.CardsDev.GetSampleResponse(context.TODO())
	if err != nil {
		var apierr *hubspotsdk.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestExtensionCardsDevMigrateViewsWithOptionalParams(t *testing.T) {
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
	_, err := client.Crm.Extensions.CardsDev.MigrateViews(
		context.TODO(),
		0,
		crm.ExtensionCardsDevMigrateViewsParams{
			CardMigrateViewsRequest: crm.CardMigrateViewsRequestParam{
				AllowDuplicateAppCardIDs: true,
				AppCardID:                0,
				LegacyCrmCardID:          0,
				HelpdeskAppCardID:        hubspotsdk.Int(0),
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
