// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package crm_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/stainless-sdks/hubspot-sdk-go"
	"github.com/stainless-sdks/hubspot-sdk-go/crm"
	"github.com/stainless-sdks/hubspot-sdk-go/internal/testutil"
	"github.com/stainless-sdks/hubspot-sdk-go/option"
)

func TestExtensionCardNew(t *testing.T) {
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
		option.WithAccessToken("pat-na1-xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"),
	)
	_, err := client.CRM.Extensions.Cards.New(
		context.TODO(),
		0,
		crm.ExtensionCardNewParams{
			CardCreateRequest: crm.CardCreateRequestParam{
				Actions: crm.CardActionsParam{
					BaseURLs: []string{"https://www.example.com/hubspot"},
				},
				Display: crm.CardDisplayBodyParam{
					Properties: []crm.CardDisplayPropertyParam{{
						DataType: crm.CardDisplayPropertyDataTypeString,
						Label:    "Pets Name",
						Name:     "pet_name",
						Options: []crm.DisplayOptionParam{{
							Label: "label",
							Name:  "name",
							Type:  crm.DisplayOptionTypeDefault,
						}},
					}},
				},
				Fetch: crm.CardFetchBodyParam{
					ObjectTypes: []crm.CardObjectTypeBodyParam{{
						Name:             crm.CardObjectTypeBodyNameContacts,
						PropertiesToSend: []string{"email", "firstname"},
					}},
					TargetURL:          "https://www.example.com/hubspot/target",
					CardType:           crm.CardFetchBodyCardTypeExternal,
					ServerlessFunction: hubspotsdk.String("serverlessFunction"),
				},
				Title: "PetSpot",
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

func TestExtensionCardUpdateWithOptionalParams(t *testing.T) {
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
		option.WithAccessToken("pat-na1-xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"),
	)
	_, err := client.CRM.Extensions.Cards.Update(
		context.TODO(),
		"cardId",
		crm.ExtensionCardUpdateParams{
			AppID: 0,
			CardPatchRequest: crm.CardPatchRequestParam{
				Actions: crm.CardActionsParam{
					BaseURLs: []string{"https://www.example.com/hubspot"},
				},
				Display: crm.CardDisplayBodyParam{
					Properties: []crm.CardDisplayPropertyParam{{
						DataType: crm.CardDisplayPropertyDataTypeString,
						Label:    "Pets Name",
						Name:     "pet_name",
						Options: []crm.DisplayOptionParam{{
							Label: "label",
							Name:  "name",
							Type:  crm.DisplayOptionTypeDefault,
						}},
					}},
				},
				Fetch: crm.CardFetchBodyPatchParam{
					ObjectTypes: []crm.CardObjectTypeBodyParam{{
						Name:             crm.CardObjectTypeBodyNameContacts,
						PropertiesToSend: []string{"email", "firstname"},
					}},
					CardType:           crm.CardFetchBodyPatchCardTypeExternal,
					ServerlessFunction: hubspotsdk.String("serverlessFunction"),
					TargetURL:          hubspotsdk.String("https://www.example.com/hubspot/target"),
				},
				Title: hubspotsdk.String("PetSpot"),
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

func TestExtensionCardList(t *testing.T) {
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
		option.WithAccessToken("pat-na1-xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"),
	)
	_, err := client.CRM.Extensions.Cards.List(context.TODO(), 0)
	if err != nil {
		var apierr *hubspotsdk.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestExtensionCardDelete(t *testing.T) {
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
		option.WithAccessToken("pat-na1-xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"),
	)
	err := client.CRM.Extensions.Cards.Delete(
		context.TODO(),
		"cardId",
		crm.ExtensionCardDeleteParams{
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

func TestExtensionCardGet(t *testing.T) {
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
		option.WithAccessToken("pat-na1-xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"),
	)
	_, err := client.CRM.Extensions.Cards.Get(
		context.TODO(),
		"cardId",
		crm.ExtensionCardGetParams{
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

func TestExtensionCardGetSampleResponse(t *testing.T) {
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
		option.WithAccessToken("pat-na1-xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"),
	)
	_, err := client.CRM.Extensions.Cards.GetSampleResponse(context.TODO())
	if err != nil {
		var apierr *hubspotsdk.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
