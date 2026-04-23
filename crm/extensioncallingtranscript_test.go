// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package crm_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/stainless-sdks/hubspot-sdk-go"
	"github.com/stainless-sdks/hubspot-sdk-go/crm"
	"github.com/stainless-sdks/hubspot-sdk-go/internal/testutil"
	"github.com/stainless-sdks/hubspot-sdk-go/option"
)

func TestExtensionCallingTranscriptNew(t *testing.T) {
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
	_, err := client.Crm.Extensions.Calling.Transcripts.New(context.TODO(), crm.ExtensionCallingTranscriptNewParams{
		TranscriptCreateRequest: crm.TranscriptCreateRequestParam{
			EngagementID: 0,
			TranscriptCreateUtterances: []crm.TranscriptCreateUtteranceParam{{
				EndTimeMillis: 0,
				Speaker: crm.SpeakerParam{
					ID:    "id",
					Name:  "name",
					Email: hubspotsdk.String("email"),
				},
				StartTimeMillis: 0,
				Text:            "text",
				LanguageCode:    hubspotsdk.String("languageCode"),
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

func TestExtensionCallingTranscriptDelete(t *testing.T) {
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
	err := client.Crm.Extensions.Calling.Transcripts.Delete(context.TODO(), "transcriptId")
	if err != nil {
		var apierr *hubspotsdk.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestExtensionCallingTranscriptNewInboundCallWithOptionalParams(t *testing.T) {
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
	_, err := client.Crm.Extensions.Calling.Transcripts.NewInboundCall(context.TODO(), crm.ExtensionCallingTranscriptNewInboundCallParams{
		CompletedThirdPartyCallRequest: crm.CompletedThirdPartyCallRequestParam{
			CreateEngagement: true,
			EngagementProperties: map[string]string{
				"foo": "string",
			},
			ExternalCallID:  "externalCallId",
			FinalCallStatus: crm.CompletedThirdPartyCallRequestFinalCallStatusBusy,
			FromNumber: crm.FormattedPhoneNumberParam{
				E164Number:      "e164Number",
				PhoneNumberType: crm.FormattedPhoneNumberPhoneNumberTypeFixedLine,
				Extension:       hubspotsdk.String("extension"),
			},
			PotentialRecipientUserIDs: []int64{0},
			ToNumber: crm.FormattedPhoneNumberParam{
				E164Number:      "e164Number",
				PhoneNumberType: crm.FormattedPhoneNumberPhoneNumberTypeFixedLine,
				Extension:       hubspotsdk.String("extension"),
			},
			CallStartedTimestamp: hubspotsdk.Time(time.Now()),
			DurationSeconds:      hubspotsdk.Int(0),
			UserID:               hubspotsdk.Int(0),
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

func TestExtensionCallingTranscriptGet(t *testing.T) {
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
	_, err := client.Crm.Extensions.Calling.Transcripts.Get(context.TODO(), "transcriptId")
	if err != nil {
		var apierr *hubspotsdk.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
