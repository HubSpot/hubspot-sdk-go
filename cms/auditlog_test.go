// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cms_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/stainless-sdks/hubspot-sdk-go"
	"github.com/stainless-sdks/hubspot-sdk-go/cms"
	"github.com/stainless-sdks/hubspot-sdk-go/internal/testutil"
	"github.com/stainless-sdks/hubspot-sdk-go/option"
)

func TestAuditLogListWithOptionalParams(t *testing.T) {
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
	_, err := client.Cms.AuditLogs.List(context.TODO(), cms.AuditLogListParams{
		After:      hubspotsdk.String("after"),
		Before:     hubspotsdk.String("before"),
		EventType:  []string{"string"},
		Limit:      hubspotsdk.Int(0),
		ObjectID:   []string{"string"},
		ObjectType: []string{"string"},
		Sort:       []string{"string"},
		UserID:     []string{"string"},
	})
	if err != nil {
		var apierr *hubspotsdk.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAuditLogExportWithOptionalParams(t *testing.T) {
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
	err := client.Cms.AuditLogs.Export(context.TODO(), cms.AuditLogExportParams{
		CmsAuditLoggingExportSettings: cms.CmsAuditLoggingExportSettingsParam{
			Email:                           "email",
			Format:                          cms.CmsAuditLoggingExportSettingsFormatCsv,
			PortalID:                        0,
			RecipientUserIDs:                []int64{0},
			ShouldMarkExportFileAsSensitive: true,
			Type:                            "type",
			Filters: cms.CmsAuditLoggingExportFiltersParam{
				ObjectType: []string{"string"},
			},
			Partition:    hubspotsdk.Int(0),
			UserID:       hubspotsdk.Int(0),
			UserTimeZone: hubspotsdk.String("userTimeZone"),
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
