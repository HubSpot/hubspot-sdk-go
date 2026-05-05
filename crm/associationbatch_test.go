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

func TestAssociationBatchNew(t *testing.T) {
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
	_, err := client.Crm.Associations.Batch.New(
		context.TODO(),
		"toObjectId",
		crm.AssociationBatchNewParams{
			FromObjectType: "fromObjectType",
			FromObjectID:   "fromObjectId",
			ToObjectType:   "toObjectType",
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

func TestAssociationBatchDelete(t *testing.T) {
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
	err := client.Crm.Associations.Batch.Delete(
		context.TODO(),
		"toObjectType",
		crm.AssociationBatchDeleteParams{
			FromObjectType: "fromObjectType",
			BatchInputPublicAssociationMultiArchive: crm.BatchInputPublicAssociationMultiArchiveParam{
				Inputs: []crm.PublicAssociationMultiArchiveParam{{
					From: shared.PublicObjectIDParam{
						ID: "id",
					},
					To: []shared.PublicObjectIDParam{{
						ID: "id",
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

func TestAssociationBatchNewDefault(t *testing.T) {
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
	_, err := client.Crm.Associations.Batch.NewDefault(
		context.TODO(),
		"toObjectType",
		crm.AssociationBatchNewDefaultParams{
			FromObjectType: "fromObjectType",
			BatchInputPublicDefaultAssociationMultiPost: crm.BatchInputPublicDefaultAssociationMultiPostParam{
				Inputs: []crm.PublicDefaultAssociationMultiPostParam{{
					From: shared.PublicObjectIDParam{
						ID: "id",
					},
					To: shared.PublicObjectIDParam{
						ID: "id",
					},
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

func TestAssociationBatchDeleteLabels(t *testing.T) {
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
	err := client.Crm.Associations.Batch.DeleteLabels(
		context.TODO(),
		"toObjectType",
		crm.AssociationBatchDeleteLabelsParams{
			FromObjectType: "fromObjectType",
			BatchInputPublicAssociationMultiPost: crm.BatchInputPublicAssociationMultiPostParam{
				Inputs: []crm.PublicAssociationMultiPostParam{{
					From: shared.PublicObjectIDParam{
						ID: "id",
					},
					To: shared.PublicObjectIDParam{
						ID: "id",
					},
					Types: []shared.AssociationSpecParam{{
						AssociationCategory: shared.AssociationSpecAssociationCategoryHubSpotDefined,
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

func TestAssociationBatchGet(t *testing.T) {
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
	_, err := client.Crm.Associations.Batch.Get(
		context.TODO(),
		"toObjectType",
		crm.AssociationBatchGetParams{
			FromObjectType: "fromObjectType",
			BatchInputPublicFetchAssociationsBatchRequest: crm.BatchInputPublicFetchAssociationsBatchRequestParam{
				Inputs: []crm.PublicFetchAssociationsBatchRequestParam{{
					ID:    "id",
					After: hubspotsdk.String("after"),
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
