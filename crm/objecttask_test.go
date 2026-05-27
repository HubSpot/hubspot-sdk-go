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

func TestObjectTaskNew(t *testing.T) {
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
	_, err := client.Crm.Objects.Tasks.New(context.TODO(), crm.ObjectTaskNewParams{
		SimplePublicObjectInputForCreate: crm.SimplePublicObjectInputForCreateParam{
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

func TestObjectTaskUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Crm.Objects.Tasks.Update(
		context.TODO(),
		"taskId",
		crm.ObjectTaskUpdateParams{
			SimplePublicObjectInput: crm.SimplePublicObjectInputParam{
				Properties: map[string]string{
					"foo": "string",
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

func TestObjectTaskListWithOptionalParams(t *testing.T) {
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
	_, err := client.Crm.Objects.Tasks.List(context.TODO(), crm.ObjectTaskListParams{
		After:                 hubspotsdk.String("after"),
		Archived:              hubspotsdk.Bool(true),
		Associations:          []string{"string"},
		Limit:                 hubspotsdk.Int(0),
		Properties:            []string{"string"},
		PropertiesWithHistory: []string{"string"},
	})
	if err != nil {
		var apierr *hubspotsdk.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestObjectTaskDelete(t *testing.T) {
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
	err := client.Crm.Objects.Tasks.Delete(context.TODO(), "taskId")
	if err != nil {
		var apierr *hubspotsdk.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestObjectTaskGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Crm.Objects.Tasks.Get(
		context.TODO(),
		"taskId",
		crm.ObjectTaskGetParams{
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

func TestObjectTaskSearchWithOptionalParams(t *testing.T) {
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
	_, err := client.Crm.Objects.Tasks.Search(context.TODO(), crm.ObjectTaskSearchParams{
		PublicObjectSearchRequest: crm.PublicObjectSearchRequestParam{
			After: "after",
			FilterGroups: []crm.FilterGroupParam{{
				Filters: []crm.CrmFilterParam{{
					Operator:     crm.CrmFilterOperatorBetween,
					PropertyName: "propertyName",
					HighValue:    hubspotsdk.String("highValue"),
					Value:        hubspotsdk.String("value"),
					Values:       []string{"string"},
				}},
			}},
			Limit:      0,
			Properties: []string{"string"},
			Sorts:      []string{"string"},
			Query:      hubspotsdk.String("query"),
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
