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

func TestExtensionCallingRecordingSettingNew(t *testing.T) {
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
	_, err := client.CRM.Extensions.Calling.RecordingSettings.New(
		context.TODO(),
		0,
		crm.ExtensionCallingRecordingSettingNewParams{
			RecordingSettingsRequest: crm.RecordingSettingsRequestParam{
				URLToRetrieveAuthedRecording: "urlToRetrieveAuthedRecording",
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

func TestExtensionCallingRecordingSettingUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.CRM.Extensions.Calling.RecordingSettings.Update(
		context.TODO(),
		0,
		crm.ExtensionCallingRecordingSettingUpdateParams{
			RecordingSettingsPatchRequest: crm.RecordingSettingsPatchRequestParam{
				URLToRetrieveAuthedRecording: hubspotsdk.String("urlToRetrieveAuthedRecording"),
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

func TestExtensionCallingRecordingSettingGet(t *testing.T) {
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
	_, err := client.CRM.Extensions.Calling.RecordingSettings.Get(context.TODO(), 0)
	if err != nil {
		var apierr *hubspotsdk.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestExtensionCallingRecordingSettingMarkReady(t *testing.T) {
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
	err := client.CRM.Extensions.Calling.RecordingSettings.MarkReady(context.TODO(), crm.ExtensionCallingRecordingSettingMarkReadyParams{
		MarkRecordingAsReadyRequest: crm.MarkRecordingAsReadyRequestParam{
			EngagementID: 0,
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
