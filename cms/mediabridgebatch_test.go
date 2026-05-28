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

func TestMediaBridgeBatchNew(t *testing.T) {
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
	_, err := client.Cms.MediaBridge.Batch.New(
		context.TODO(),
		"objectType",
		cms.MediaBridgeBatchNewParams{
			AppID: 0,
			BatchInputPropertyCreate: shared.BatchInputPropertyCreateParam{
				Inputs: []shared.PropertyCreateParam{{
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
					TextDisplayHint:      shared.PropertyCreateTextDisplayHintDomainName,
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

func TestMediaBridgeBatchDelete(t *testing.T) {
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
	err := client.Cms.MediaBridge.Batch.Delete(
		context.TODO(),
		"objectType",
		cms.MediaBridgeBatchDeleteParams{
			AppID: 0,
			BatchInputPropertyName: shared.BatchInputPropertyNameParam{
				Inputs: []shared.PropertyNameParam{{
					Name: "name",
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

func TestMediaBridgeBatchGet(t *testing.T) {
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
	_, err := client.Cms.MediaBridge.Batch.Get(
		context.TODO(),
		"objectType",
		cms.MediaBridgeBatchGetParams{
			AppID: 0,
			BatchReadInputPropertyName: shared.BatchReadInputPropertyNameParam{
				Archived:        true,
				DataSensitivity: shared.BatchReadInputPropertyNameDataSensitivityHighlySensitive,
				Inputs: []shared.PropertyNameParam{{
					Name: "name",
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
