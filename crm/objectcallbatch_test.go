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

func TestObjectCallBatchNew(t *testing.T) {
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
	_, err := client.Crm.Objects.Calls.Batch.New(context.TODO(), crm.ObjectCallBatchNewParams{
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

func TestObjectCallBatchUpdate(t *testing.T) {
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
	_, err := client.Crm.Objects.Calls.Batch.Update(context.TODO(), crm.ObjectCallBatchUpdateParams{
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

func TestObjectCallBatchDelete(t *testing.T) {
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
	err := client.Crm.Objects.Calls.Batch.Delete(context.TODO(), crm.ObjectCallBatchDeleteParams{
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

func TestObjectCallBatchGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Crm.Objects.Calls.Batch.Get(context.TODO(), crm.ObjectCallBatchGetParams{
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

func TestObjectCallBatchUpsert(t *testing.T) {
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
	_, err := client.Crm.Objects.Calls.Batch.Upsert(context.TODO(), crm.ObjectCallBatchUpsertParams{
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
