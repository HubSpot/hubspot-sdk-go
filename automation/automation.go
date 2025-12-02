// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package automation

import (
	"github.com/stainless-sdks/hubspot-sdk-go/option"
)

// AutomationService contains methods and other services that help with interacting
// with the hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAutomationService] method instead.
type AutomationService struct {
	Options   []option.RequestOption
	Actions   ActionService
	Sequences SequenceService
	Workflows WorkflowService
}

// NewAutomationService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAutomationService(opts ...option.RequestOption) (r AutomationService) {
	r = AutomationService{}
	r.Options = opts
	r.Actions = NewActionService(opts...)
	r.Sequences = NewSequenceService(opts...)
	r.Workflows = NewWorkflowService(opts...)
	return
}
