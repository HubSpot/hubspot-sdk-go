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
	"github.com/HubSpot/hubspot-sdk-go/shared"
)

func TestPropertyNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Crm.Properties.New(
		context.TODO(),
		"objectType",
		crm.PropertyNewParams{
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

func TestPropertyUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Crm.Properties.Update(
		context.TODO(),
		"propertyName",
		crm.PropertyUpdateParams{
			ObjectType: "objectType",
			PropertyUpdate: crm.PropertyUpdateParam{
				CalculationFormula:   hubspotsdk.String("calculationFormula"),
				CurrencyPropertyName: hubspotsdk.String("currencyPropertyName"),
				Description:          hubspotsdk.String("description"),
				DisplayOrder:         hubspotsdk.Int(0),
				FieldType:            crm.PropertyUpdateFieldTypeBooleancheckbox,
				FormField:            hubspotsdk.Bool(true),
				GroupName:            hubspotsdk.String("groupName"),
				Hidden:               hubspotsdk.Bool(true),
				Label:                hubspotsdk.String("label"),
				NumberDisplayHint:    crm.PropertyUpdateNumberDisplayHintCurrency,
				Options: []shared.OptionInputParam{{
					DisplayOrder: 0,
					Hidden:       true,
					Label:        "label",
					Value:        "value",
					Description:  hubspotsdk.String("description"),
				}},
				ShowCurrencySymbol: hubspotsdk.Bool(true),
				Type:               crm.PropertyUpdateTypeBool,
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

func TestPropertyListWithOptionalParams(t *testing.T) {
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
	_, err := client.Crm.Properties.List(
		context.TODO(),
		"objectType",
		crm.PropertyListParams{
			Archived:        hubspotsdk.Bool(true),
			DataSensitivity: crm.PropertyListParamsDataSensitivityHighlySensitive,
			Locale:          hubspotsdk.String("locale"),
			Properties:      hubspotsdk.String("properties"),
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

func TestPropertyDelete(t *testing.T) {
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
	err := client.Crm.Properties.Delete(
		context.TODO(),
		"propertyName",
		crm.PropertyDeleteParams{
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

func TestPropertyGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Crm.Properties.Get(
		context.TODO(),
		"propertyName",
		crm.PropertyGetParams{
			ObjectType:      "objectType",
			Archived:        hubspotsdk.Bool(true),
			DataSensitivity: crm.PropertyGetParamsDataSensitivityHighlySensitive,
			Locale:          hubspotsdk.String("locale"),
			Properties:      hubspotsdk.String("properties"),
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
