// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package scheduler_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/stainless-sdks/hubspot-sdk-go"
	"github.com/stainless-sdks/hubspot-sdk-go/internal/testutil"
	"github.com/stainless-sdks/hubspot-sdk-go/option"
	"github.com/stainless-sdks/hubspot-sdk-go/scheduler"
)

func TestMeetingMeetingsLinkListWithOptionalParams(t *testing.T) {
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
		option.WithAccessToken("pat-na1-xxxxxxxx-xxxx"),
	)
	_, err := client.Scheduler.Meetings.MeetingsLinks.List(context.TODO(), scheduler.MeetingMeetingsLinkListParams{
		After:           hubspotsdk.String("after"),
		Limit:           hubspotsdk.Int(0),
		Name:            hubspotsdk.String("name"),
		OrganizerUserID: hubspotsdk.String("organizerUserId"),
		Type:            hubspotsdk.String("type"),
	})
	if err != nil {
		var apierr *hubspotsdk.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestMeetingMeetingsLinkBookWithOptionalParams(t *testing.T) {
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
		option.WithAccessToken("pat-na1-xxxxxxxx-xxxx"),
	)
	_, err := client.Scheduler.Meetings.MeetingsLinks.Book(context.TODO(), scheduler.MeetingMeetingsLinkBookParams{
		ExternalMeetingBooking: scheduler.ExternalMeetingBookingParam{
			Duration:  0,
			Email:     "email",
			FirstName: "firstName",
			FormFields: []scheduler.ExternalBookingFormFieldParam{{
				Name:  "name",
				Value: "value",
			}},
			LastName: "lastName",
			LegalConsentResponses: []scheduler.ExternalLegalConsentResponseParam{{
				CommunicationTypeID: "communicationTypeId",
				Consented:           true,
			}},
			LikelyAvailableUserIDs: []string{"string"},
			Slug:                   "slug",
			StartTime:              time.Now(),
			Locale:                 hubspotsdk.String("locale"),
			Timezone:               hubspotsdk.String("timezone"),
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

func TestMeetingMeetingsLinkGetAvailabilityBySlugWithOptionalParams(t *testing.T) {
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
		option.WithAccessToken("pat-na1-xxxxxxxx-xxxx"),
	)
	_, err := client.Scheduler.Meetings.MeetingsLinks.GetAvailabilityBySlug(
		context.TODO(),
		"slug",
		scheduler.MeetingMeetingsLinkGetAvailabilityBySlugParams{
			Timezone:    "timezone",
			MonthOffset: hubspotsdk.Int(0),
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

func TestMeetingMeetingsLinkGetBookingInfoBySlug(t *testing.T) {
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
		option.WithAccessToken("pat-na1-xxxxxxxx-xxxx"),
	)
	_, err := client.Scheduler.Meetings.MeetingsLinks.GetBookingInfoBySlug(
		context.TODO(),
		"slug",
		scheduler.MeetingMeetingsLinkGetBookingInfoBySlugParams{
			Timezone: "timezone",
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
