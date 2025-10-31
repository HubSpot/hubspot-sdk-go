// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package scheduler

import (
	"github.com/stainless-sdks/hubspot-sdk-go/option"
)

// SchedulerService contains methods and other services that help with interacting
// with the Hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSchedulerService] method instead.
type SchedulerService struct {
	Options  []option.RequestOption
	Meetings MeetingService
}

// NewSchedulerService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewSchedulerService(opts ...option.RequestOption) (r SchedulerService) {
	r = SchedulerService{}
	r.Options = opts
	r.Meetings = NewMeetingService(opts...)
	return
}
