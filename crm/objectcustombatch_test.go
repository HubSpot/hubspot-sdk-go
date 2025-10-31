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

func TestObjectCustomBatchNew(t *testing.T) {
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
	_, err := client.CRM.Objects.Custom.Batch.New(
		context.TODO(),
		"objectType",
		crm.ObjectCustomBatchNewParams{
			BatchInputSimplePublicObjectBatchInputForCreate: crm.BatchInputSimplePublicObjectBatchInputForCreateParam{
				Inputs: []crm.SimplePublicObjectBatchInputForCreateParam{{
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
					ObjectWriteTraceID: hubspotsdk.String("objectWriteTraceId"),
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

func TestObjectCustomBatchUpdate(t *testing.T) {
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
	_, err := client.CRM.Objects.Custom.Batch.Update(
		context.TODO(),
		"objectType",
		crm.ObjectCustomBatchUpdateParams{
			BatchInputSimplePublicObjectBatchInput: crm.BatchInputSimplePublicObjectBatchInputParam{
				Inputs: []crm.SimplePublicObjectBatchInputParam{{
					ID: "id",
					Properties: map[string]string{
						"foo": "string",
					},
					IDProperty:         hubspotsdk.String("my_unique_property_name"),
					ObjectWriteTraceID: hubspotsdk.String("objectWriteTraceId"),
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

func TestObjectCustomBatchDelete(t *testing.T) {
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
	err := client.CRM.Objects.Custom.Batch.Delete(
		context.TODO(),
		"objectType",
		crm.ObjectCustomBatchDeleteParams{
			BatchInputSimplePublicObjectID: crm.BatchInputSimplePublicObjectIDParam{
				Inputs: []crm.SimplePublicObjectIDParam{{
					ID: "id",
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

func TestObjectCustomBatchGetWithOptionalParams(t *testing.T) {
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
	_, err := client.CRM.Objects.Custom.Batch.Get(
		context.TODO(),
		"objectType",
		crm.ObjectCustomBatchGetParams{
			BatchReadInputSimplePublicObjectID: crm.BatchReadInputSimplePublicObjectIDParam{
				Inputs: []crm.SimplePublicObjectIDParam{{
					ID: "id",
				}},
				Properties:            []string{"string"},
				PropertiesWithHistory: []string{"string"},
				IDProperty:            hubspotsdk.String("idProperty"),
			},
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

func TestObjectCustomBatchUpsert(t *testing.T) {
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
	_, err := client.CRM.Objects.Custom.Batch.Upsert(
		context.TODO(),
		"objectType",
		crm.ObjectCustomBatchUpsertParams{
			BatchInputSimplePublicObjectBatchInputUpsert: crm.BatchInputSimplePublicObjectBatchInputUpsertParam{
				Inputs: []crm.SimplePublicObjectBatchInputUpsertParam{{
					ID: "id",
					Properties: map[string]string{
						"foo": "string",
					},
					IDProperty:         hubspotsdk.String("idProperty"),
					ObjectWriteTraceID: hubspotsdk.String("objectWriteTraceId"),
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
