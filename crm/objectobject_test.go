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

func TestObjectObjectNewWithOptionalParams(t *testing.T) {
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
	_, err := client.CRM.Objects.Objects.New(
		context.TODO(),
		"objectType",
		crm.ObjectObjectNewParams{
			SimplePublicObjectInputForCreate: crm.SimplePublicObjectInputForCreateParam{
				Properties: map[string]string{
					"foo": "string",
				},
				Associations: []crm.PublicAssociationsForObjectParam{{
					To: shared.PublicObjectIDParam{
						ID: "37295",
					},
					Types: []shared.AssociationSpecParam{{
						AssociationCategory: shared.AssociationSpecAssociationCategoryHubspotDefined,
						AssociationTypeID:   0,
					}},
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

func TestObjectObjectUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.CRM.Objects.Objects.Update(
		context.TODO(),
		"objectId",
		crm.ObjectObjectUpdateParams{
			ObjectType: "objectType",
			SimplePublicObjectInput: crm.SimplePublicObjectInputParam{
				Properties: map[string]string{
					"property_checkbox":            "false",
					"property_date":                "1572480000000",
					"property_dropdown":            "choice_b",
					"property_multiple_checkboxes": "chocolate;strawberry",
					"property_number":              "17",
					"property_radio":               "option_1",
					"property_string":              "value",
				},
			},
			IDProperty: hubspotsdk.String("idProperty"),
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

func TestObjectObjectListWithOptionalParams(t *testing.T) {
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
	_, err := client.CRM.Objects.Objects.List(
		context.TODO(),
		"objectType",
		crm.ObjectObjectListParams{
			After:                 hubspotsdk.String("after"),
			Archived:              hubspotsdk.Bool(true),
			Associations:          []string{"string"},
			Limit:                 hubspotsdk.Int(0),
			Properties:            []string{"string"},
			PropertiesWithHistory: []string{"string"},
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

func TestObjectObjectDelete(t *testing.T) {
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
	err := client.CRM.Objects.Objects.Delete(
		context.TODO(),
		"objectId",
		crm.ObjectObjectDeleteParams{
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

func TestObjectObjectGetWithOptionalParams(t *testing.T) {
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
	_, err := client.CRM.Objects.Objects.Get(
		context.TODO(),
		"objectId",
		crm.ObjectObjectGetParams{
			ObjectType:            "objectType",
			Archived:              hubspotsdk.Bool(true),
			Associations:          []string{"string"},
			IDProperty:            hubspotsdk.String("idProperty"),
			Properties:            []string{"string"},
			PropertiesWithHistory: []string{"string"},
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

func TestObjectObjectSearchWithOptionalParams(t *testing.T) {
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
	_, err := client.CRM.Objects.Objects.Search(
		context.TODO(),
		"objectType",
		crm.ObjectObjectSearchParams{
			PublicObjectSearchRequest: crm.PublicObjectSearchRequestParam{
				After: hubspotsdk.String("after"),
				FilterGroups: []crm.FilterGroupParam{{
					Filters: []crm.FilterParam{{
						Operator:     crm.FilterOperatorEq,
						PropertyName: "",
						HighValue:    hubspotsdk.String(""),
						Value:        hubspotsdk.String(""),
						Values:       []string{"string"},
					}},
				}},
				Limit:      hubspotsdk.Int(0),
				Properties: []string{"string"},
				Query:      hubspotsdk.String("query"),
				Sorts:      []string{"string"},
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
