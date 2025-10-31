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
	"github.com/stainless-sdks/hubspot-sdk-go/shared"
)

func TestMeetingCalendarNew(t *testing.T) {
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
	_, err := client.Scheduler.Meetings.Calendar.New(context.TODO(), scheduler.MeetingCalendarNewParams{
		ExternalCalendarMeetingEventCreateRequest: scheduler.ExternalCalendarMeetingEventCreateRequestParam{
			Associations: []scheduler.ExternalAssociationCreateRequestParam{{
				To: shared.PublicObjectIDParam{
					ID: "37295",
				},
				Types: []shared.AssociationSpecParam{{
					AssociationCategory: shared.AssociationSpecAssociationCategoryHubspotDefined,
					AssociationTypeID:   0,
				}},
			}},
			EmailReminderSchedule: scheduler.ExternalEmailReminderScheduleParam{
				Reminders: []scheduler.ExternalReminderParam{{
					NumberOfTimeUnits: 0,
					TimeUnit:          "timeUnit",
				}},
				ShouldIncludeInviteDescription: true,
			},
			Properties: scheduler.ExternalCalendarMeetingEventCreatePropertiesParam{
				HsMeetingEndTime:       time.Now(),
				HsMeetingOutcome:       "hs_meeting_outcome",
				HsMeetingStartTime:     time.Now(),
				HsMeetingTitle:         "hs_meeting_title",
				HsTimestamp:            time.Now(),
				HsActivityType:         hubspotsdk.String("hs_activity_type"),
				HsAttachmentIDs:        []string{"string"},
				HsAttendeeOwnerIDs:     []string{"string"},
				HsInternalMeetingNotes: hubspotsdk.String("hs_internal_meeting_notes"),
				HsMeetingBody:          hubspotsdk.String("hs_meeting_body"),
				HsMeetingLocation:      hubspotsdk.String("hs_meeting_location"),
				HsMeetingLocationType:  hubspotsdk.String("hs_meeting_location_type"),
				HubspotOwnerID:         hubspotsdk.String("hubspot_owner_id"),
			},
			Timezone: "timezone",
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
