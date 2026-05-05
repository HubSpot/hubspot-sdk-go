// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package communication_preferences_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/HubSpot/hubspot-sdk-go"
	"github.com/HubSpot/hubspot-sdk-go/communication_preferences"
	"github.com/HubSpot/hubspot-sdk-go/internal/testutil"
	"github.com/HubSpot/hubspot-sdk-go/option"
	"github.com/HubSpot/hubspot-sdk-go/shared"
)

func TestStatusBatchGetUnsubscribeAllStatusesWithOptionalParams(t *testing.T) {
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
	_, err := client.CommunicationPreferences.Statuses.Batch.GetUnsubscribeAllStatuses(context.TODO(), communication_preferences.StatusBatchGetUnsubscribeAllStatusesParams{
		Channel: communication_preferences.StatusBatchGetUnsubscribeAllStatusesParamsChannelEmail,
		BatchInputString: shared.BatchInputStringParam{
			Inputs: []string{"string"},
		},
		BusinessUnitID: hubspotsdk.Int(0),
	})
	if err != nil {
		var apierr *hubspotsdk.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestStatusBatchReadWithOptionalParams(t *testing.T) {
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
	_, err := client.CommunicationPreferences.Statuses.Batch.Read(context.TODO(), communication_preferences.StatusBatchReadParams{
		Channel: communication_preferences.StatusBatchReadParamsChannelEmail,
		BatchInputString: shared.BatchInputStringParam{
			Inputs: []string{"string"},
		},
		BusinessUnitID: hubspotsdk.Int(0),
	})
	if err != nil {
		var apierr *hubspotsdk.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestStatusBatchUnsubscribeAllWithOptionalParams(t *testing.T) {
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
	_, err := client.CommunicationPreferences.Statuses.Batch.UnsubscribeAll(context.TODO(), communication_preferences.StatusBatchUnsubscribeAllParams{
		Channel: communication_preferences.StatusBatchUnsubscribeAllParamsChannelEmail,
		BatchInputString: shared.BatchInputStringParam{
			Inputs: []string{"string"},
		},
		BusinessUnitID: hubspotsdk.Int(0),
		Verbose:        hubspotsdk.Bool(true),
	})
	if err != nil {
		var apierr *hubspotsdk.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestStatusBatchUpdateStatuses(t *testing.T) {
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
	_, err := client.CommunicationPreferences.Statuses.Batch.UpdateStatuses(context.TODO(), communication_preferences.StatusBatchUpdateStatusesParams{
		BatchInputPublicStatusRequest: communication_preferences.BatchInputPublicStatusRequestParam{
			Inputs: []communication_preferences.PublicStatusRequestParam{{
				Channel:               communication_preferences.PublicStatusRequestChannelEmail,
				StatusState:           communication_preferences.PublicStatusRequestStatusStateNotSpecified,
				SubscriberIDString:    "subscriberIdString",
				SubscriptionID:        0,
				LegalBasis:            communication_preferences.PublicStatusRequestLegalBasisConsentWithNotice,
				LegalBasisExplanation: hubspotsdk.String("legalBasisExplanation"),
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
