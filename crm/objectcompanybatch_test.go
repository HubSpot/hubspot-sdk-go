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

func TestObjectCompanyBatchNew(t *testing.T) {
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
	_, err := client.Crm.Objects.Companies.Batch.New(context.TODO(), crm.ObjectCompanyBatchNewParams{
		BatchInputSimplePublicObjectBatchInputForCreate: crm.BatchInputSimplePublicObjectBatchInputForCreateParam{
			Inputs: []crm.SimplePublicObjectBatchInputForCreateParam{{
				Associations: []crm.PublicAssociationsForObjectParam{{
					To: shared.PublicObjectIDParam{
						ID: "id",
					},
					Types: []shared.AssociationSpecParam{{
						AssociationCategory: shared.AssociationSpecAssociationCategoryHubSpotDefined,
						AssociationTypeID:   0,
					}},
				}},
				Properties: map[string]string{
					"foo": "string",
				},
				ObjectWriteTraceID: hubspotsdk.String("objectWriteTraceId"),
			}},
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

func TestObjectCompanyBatchUpdate(t *testing.T) {
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
	_, err := client.Crm.Objects.Companies.Batch.Update(context.TODO(), crm.ObjectCompanyBatchUpdateParams{
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
	})
	if err != nil {
		var apierr *hubspotsdk.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestObjectCompanyBatchDelete(t *testing.T) {
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
	err := client.Crm.Objects.Companies.Batch.Delete(context.TODO(), crm.ObjectCompanyBatchDeleteParams{
		BatchInputSimplePublicObjectID: crm.BatchInputSimplePublicObjectIDParam{
			Inputs: []crm.SimplePublicObjectIDParam{{
				ID: "430001",
			}},
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

func TestObjectCompanyBatchGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Crm.Objects.Companies.Batch.Get(context.TODO(), crm.ObjectCompanyBatchGetParams{
		BatchReadInputSimplePublicObjectID: crm.BatchReadInputSimplePublicObjectIDParam{
			Inputs: []crm.SimplePublicObjectIDParam{{
				ID: "430001",
			}},
			Properties:            []string{"string"},
			PropertiesWithHistory: []string{"string"},
			IDProperty:            hubspotsdk.String("idProperty"),
		},
		Archived: hubspotsdk.Bool(true),
	})
	if err != nil {
		var apierr *hubspotsdk.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestObjectCompanyBatchUpsert(t *testing.T) {
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
	_, err := client.Crm.Objects.Companies.Batch.Upsert(context.TODO(), crm.ObjectCompanyBatchUpsertParams{
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
	})
	if err != nil {
		var apierr *hubspotsdk.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
