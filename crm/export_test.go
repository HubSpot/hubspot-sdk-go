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

func TestExportNewAsyncWithOptionalParams(t *testing.T) {
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
	_, err := client.Crm.Exports.NewAsync(context.TODO(), crm.ExportNewAsyncParams{
		PublicExportRequest: crm.PublicExportRequestUnionParam{
			OfView: &crm.PublicExportViewRequestParam{
				AssociatedObjectType:        []string{"string"},
				ExportInternalValuesOptions: []string{"NAMES"},
				ExportName:                  "exportName",
				ExportType:                  crm.PublicExportViewRequestExportTypeView,
				Format:                      crm.PublicExportViewRequestFormatCsv,
				IncludeLabeledAssociations:  true,
				IncludePrimaryDisplayPropertyForAssociatedObjects: true,
				Language:         crm.PublicExportViewRequestLanguageAfZa,
				ObjectProperties: []string{"string"},
				ObjectType:       "objectType",
				OverrideAssociatedObjectsPerDefinitionPerRowLimit: true,
				PublicCrmSearchRequest: crm.PublicCrmSearchRequestParam{
					FilterGroups: []crm.FilterGroupParam{{
						Filters: []crm.CrmFilterParam{{
							Operator:     crm.CrmFilterOperatorBetween,
							PropertyName: "propertyName",
							HighValue:    hubspotsdk.String("highValue"),
							Value:        hubspotsdk.String("value"),
							Values:       []string{"string"},
						}},
					}},
					Filters: []crm.CrmFilterParam{{
						Operator:     crm.CrmFilterOperatorBetween,
						PropertyName: "propertyName",
						HighValue:    hubspotsdk.String("highValue"),
						Value:        hubspotsdk.String("value"),
						Values:       []string{"string"},
					}},
					Sorts: []string{"string"},
					Query: hubspotsdk.String("query"),
				},
			},
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

func TestExportGet(t *testing.T) {
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
	_, err := client.Crm.Exports.Get(context.TODO(), 0)
	if err != nil {
		var apierr *hubspotsdk.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestExportGetStatus(t *testing.T) {
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
	_, err := client.Crm.Exports.GetStatus(context.TODO(), 0)
	if err != nil {
		var apierr *hubspotsdk.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
