// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package communication_preferences_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/stainless-sdks/hubspot-sdk-go"
	"github.com/stainless-sdks/hubspot-sdk-go/communication_preferences"
	"github.com/stainless-sdks/hubspot-sdk-go/internal/testutil"
	"github.com/stainless-sdks/hubspot-sdk-go/option"
)

func TestCommunicationPreferenceGenerateLinksWithOptionalParams(t *testing.T) {
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
	_, err := client.CommunicationPreferences.GenerateLinks(context.TODO(), communication_preferences.CommunicationPreferenceGenerateLinksParams{
		Channel: communication_preferences.CommunicationPreferenceGenerateLinksParamsChannelEmail,
		LinkGenerationRequest: communication_preferences.LinkGenerationRequestParam{
			SubscriberIDString: "subscriberIdString",
			Language:           hubspotsdk.String("language"),
			SubscriptionID:     hubspotsdk.Int(0),
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

func TestCommunicationPreferenceGetStatusesWithOptionalParams(t *testing.T) {
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
	_, err := client.CommunicationPreferences.GetStatuses(
		context.TODO(),
		"subscriberIdString",
		communication_preferences.CommunicationPreferenceGetStatusesParams{
			Channel:        communication_preferences.CommunicationPreferenceGetStatusesParamsChannelEmail,
			BusinessUnitID: hubspotsdk.Int(0),
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

func TestCommunicationPreferenceGetUnsubscribeAllStatusWithOptionalParams(t *testing.T) {
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
	_, err := client.CommunicationPreferences.GetUnsubscribeAllStatus(
		context.TODO(),
		"subscriberIdString",
		communication_preferences.CommunicationPreferenceGetUnsubscribeAllStatusParams{
			Channel:        communication_preferences.CommunicationPreferenceGetUnsubscribeAllStatusParamsChannelEmail,
			BusinessUnitID: hubspotsdk.Int(0),
			Verbose:        hubspotsdk.Bool(true),
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

func TestCommunicationPreferenceUnsubscribeAllWithOptionalParams(t *testing.T) {
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
	_, err := client.CommunicationPreferences.UnsubscribeAll(
		context.TODO(),
		"subscriberIdString",
		communication_preferences.CommunicationPreferenceUnsubscribeAllParams{
			Channel:        communication_preferences.CommunicationPreferenceUnsubscribeAllParamsChannelEmail,
			BusinessUnitID: hubspotsdk.Int(0),
			Verbose:        hubspotsdk.Bool(true),
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

func TestCommunicationPreferenceUpdateStatusWithOptionalParams(t *testing.T) {
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
	_, err := client.CommunicationPreferences.UpdateStatus(
		context.TODO(),
		"subscriberIdString",
		communication_preferences.CommunicationPreferenceUpdateStatusParams{
			PartialPublicStatusRequest: communication_preferences.PartialPublicStatusRequestParam{
				Channel:               communication_preferences.PartialPublicStatusRequestChannelEmail,
				StatusState:           communication_preferences.PartialPublicStatusRequestStatusStateNotSpecified,
				SubscriptionID:        0,
				LegalBasis:            communication_preferences.PartialPublicStatusRequestLegalBasisConsentWithNotice,
				LegalBasisExplanation: hubspotsdk.String("legalBasisExplanation"),
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
