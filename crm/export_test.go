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

func TestExportNewWithOptionalParams(t *testing.T) {
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
	_, err := client.CRM.Exports.New(context.TODO(), crm.ExportNewParams{
		PublicExportRequest: crm.PublicExportRequestUnionParam{
			OfPublicExportViewRequest: &crm.PublicExportViewRequestParam{
				ExportInternalValuesOptions: []string{"NAMES"},
				ExportName:                  "exportName",
				ExportType:                  crm.PublicExportViewRequestExportTypeView,
				Format:                      crm.PublicExportViewRequestFormatXls,
				Language:                    crm.PublicExportViewRequestLanguageEn,
				ObjectProperties:            []string{"string"},
				ObjectType:                  "objectType",
				OverrideAssociatedObjectsPerDefinitionPerRowLimit: true,
				AssociatedObjectType:                              hubspotsdk.String("associatedObjectType"),
				PublicCRMSearchRequest: crm.PublicCRMSearchRequestParam{
					Filters: []crm.FilterParam{{
						Operator:     crm.FilterOperatorEq,
						PropertyName: "",
						HighValue:    hubspotsdk.String(""),
						Value:        hubspotsdk.String(""),
						Values:       []string{"string"},
					}},
					Query: "query",
					Sorts: []string{"string"},
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

func TestExportGetStatus(t *testing.T) {
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
	_, err := client.CRM.Exports.GetStatus(context.TODO(), 0)
	if err != nil {
		var apierr *hubspotsdk.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
