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
	"github.com/stainless-sdks/hubspot-sdk-go/shared"
)

func TestPropertyNewWithOptionalParams(t *testing.T) {
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
	_, err := client.CRM.Properties.New(
		context.TODO(),
		"objectType",
		crm.PropertyNewParams{
			PropertyCreate: shared.PropertyCreateParam{
				FieldType:          shared.PropertyCreateFieldTypeBooleancheckbox,
				GroupName:          "groupName",
				Label:              "label",
				Name:               "name",
				Type:               shared.PropertyCreateTypeBool,
				CalculationFormula: hubspotsdk.String("calculationFormula"),
				DataSensitivity:    shared.PropertyCreateDataSensitivityNonSensitive,
				Description:        hubspotsdk.String("description"),
				DisplayOrder:       hubspotsdk.Int(0),
				ExternalOptions:    hubspotsdk.Bool(true),
				FormField:          hubspotsdk.Bool(true),
				HasUniqueValue:     hubspotsdk.Bool(true),
				Hidden:             hubspotsdk.Bool(true),
				Options: []shared.OptionInputParam{{
					DisplayOrder: 0,
					Hidden:       true,
					Label:        "label",
					Value:        "value",
					Description:  hubspotsdk.String("description"),
				}},
				ReferencedObjectType: hubspotsdk.String("referencedObjectType"),
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
	_, err := client.CRM.Properties.Update(
		context.TODO(),
		"propertyName",
		crm.PropertyUpdateParams{
			ObjectType: "objectType",
			PropertyUpdate: crm.PropertyUpdateParam{
				CalculationFormula: hubspotsdk.String("calculationFormula"),
				Description:        hubspotsdk.String("description"),
				DisplayOrder:       hubspotsdk.Int(2),
				FieldType:          crm.PropertyUpdateFieldTypeSelect,
				FormField:          hubspotsdk.Bool(true),
				GroupName:          hubspotsdk.String("contactinformation"),
				Hidden:             hubspotsdk.Bool(false),
				Label:              hubspotsdk.String("My Contact Property"),
				Options: []shared.OptionInputParam{{
					DisplayOrder: 1,
					Hidden:       false,
					Label:        "Option A",
					Value:        "A",
					Description:  hubspotsdk.String("Choice number one"),
				}, {
					DisplayOrder: 2,
					Hidden:       false,
					Label:        "Option B",
					Value:        "B",
					Description:  hubspotsdk.String("Choice number two"),
				}},
				Type: crm.PropertyUpdateTypeEnumeration,
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
	_, err := client.CRM.Properties.List(
		context.TODO(),
		"objectType",
		crm.PropertyListParams{
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

func TestPropertyDelete(t *testing.T) {
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
	err := client.CRM.Properties.Delete(
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
	_, err := client.CRM.Properties.Get(
		context.TODO(),
		"propertyName",
		crm.PropertyGetParams{
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
