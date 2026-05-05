// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package automation

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/HubSpot/hubspot-sdk-go/internal/apijson"
	"github.com/HubSpot/hubspot-sdk-go/internal/apiquery"
	shimjson "github.com/HubSpot/hubspot-sdk-go/internal/encoding/json"
	"github.com/HubSpot/hubspot-sdk-go/internal/requestconfig"
	"github.com/HubSpot/hubspot-sdk-go/option"
	"github.com/HubSpot/hubspot-sdk-go/packages/pagination"
	"github.com/HubSpot/hubspot-sdk-go/packages/param"
	"github.com/HubSpot/hubspot-sdk-go/packages/respjson"
	"github.com/HubSpot/hubspot-sdk-go/shared"
)

// SequenceService contains methods and other services that help with interacting
// with the hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSequenceService] method instead.
type SequenceService struct {
	options []option.RequestOption
}

// NewSequenceService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewSequenceService(opts ...option.RequestOption) (r SequenceService) {
	r = SequenceService{}
	r.options = opts
	return
}

// Retrieve a list of sequences available in your HubSpot account. This endpoint
// allows you to filter sequences by user ID and name, and supports pagination for
// large result sets. Use this endpoint to manage and review your sequences
// effectively.
func (r *SequenceService) List(ctx context.Context, query SequenceListParams, opts ...option.RequestOption) (res *pagination.Page[PublicSequenceLiteResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "automation/sequences/2026-03"
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// Retrieve a list of sequences available in your HubSpot account. This endpoint
// allows you to filter sequences by user ID and name, and supports pagination for
// large result sets. Use this endpoint to manage and review your sequences
// effectively.
func (r *SequenceService) ListAutoPaging(ctx context.Context, query SequenceListParams, opts ...option.RequestOption) *pagination.PageAutoPager[PublicSequenceLiteResponse] {
	return pagination.NewPageAutoPager(r.List(ctx, query, opts...))
}

// Enroll a contact into a sequence using the specified user ID and sequence
// details.
func (r *SequenceService) NewEnrollment(ctx context.Context, params SequenceNewEnrollmentParams, opts ...option.RequestOption) (res *PublicSequenceEnrollmentLiteResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "automation/sequences/2026-03/enrollments"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Retrieve details of a specific sequence by its ID.
func (r *SequenceService) Get(ctx context.Context, sequenceID string, query SequenceGetParams, opts ...option.RequestOption) (res *PublicSequenceResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if sequenceID == "" {
		err = errors.New("missing required sequenceId parameter")
		return nil, err
	}
	path := fmt.Sprintf("automation/sequences/2026-03/%s", url.PathEscape(sequenceID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Get the enrollment status of a contact in sequences by their contact ID.
func (r *SequenceService) GetEnrollmentByContactID(ctx context.Context, contactID string, opts ...option.RequestOption) (res *PublicSequenceEnrollmentResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if contactID == "" {
		err = errors.New("missing required contactId parameter")
		return nil, err
	}
	path := fmt.Sprintf("automation/sequences/2026-03/enrollments/contact/%s", url.PathEscape(contactID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type CollectionResponseWithTotalPublicSequenceLiteResponse struct {
	// An array of PublicSequenceLiteResponse objects, each representing a lightweight
	// version of a sequence.
	Results []PublicSequenceLiteResponse `json:"results" api:"required"`
	// An integer representing the total number of sequence items available.
	Total  int64         `json:"total" api:"required"`
	Paging shared.Paging `json:"paging"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Results     respjson.Field
		Total       respjson.Field
		Paging      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CollectionResponseWithTotalPublicSequenceLiteResponse) RawJSON() string { return r.JSON.raw }
func (r *CollectionResponseWithTotalPublicSequenceLiteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PublicEmailPatternResponse struct {
	// The unique identifier of the email pattern.
	ID string `json:"id" api:"required"`
	// The date and time when the email pattern was created.
	CreatedAt time.Time `json:"createdAt" api:"required" format:"date-time"`
	// The unique identifier of the email template associated with the pattern.
	TemplateID string `json:"templateId" api:"required"`
	// The date and time when the email pattern was last updated.
	UpdatedAt time.Time `json:"updatedAt" api:"required" format:"date-time"`
	// The order identifying the previous step to which the email thread is linked.
	ThreadEmailToStepOrder int64 `json:"threadEmailToStepOrder"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                     respjson.Field
		CreatedAt              respjson.Field
		TemplateID             respjson.Field
		UpdatedAt              respjson.Field
		ThreadEmailToStepOrder respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PublicEmailPatternResponse) RawJSON() string { return r.JSON.raw }
func (r *PublicEmailPatternResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PublicSequenceEnrollmentLiteResponse struct {
	// The unique identifier for the sequence enrollment.
	ID string `json:"id" api:"required"`
	// The date and time when the contact was enrolled in the sequence.
	EnrolledAt time.Time `json:"enrolledAt" api:"required" format:"date-time"`
	// The email address of the contact enrolled in the sequence.
	ToEmail string `json:"toEmail" api:"required"`
	// The date and time when the sequence enrollment was last updated.
	UpdatedAt time.Time `json:"updatedAt" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		EnrolledAt  respjson.Field
		ToEmail     respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PublicSequenceEnrollmentLiteResponse) RawJSON() string { return r.JSON.raw }
func (r *PublicSequenceEnrollmentLiteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties ContactID, SenderEmail, SequenceID are required.
type PublicSequenceEnrollmentRequestParam struct {
	// The unique identifier of the contact to be enrolled in the sequence.
	ContactID string `json:"contactId" api:"required"`
	// The email address of the sender enrolling the contact in the sequence.
	SenderEmail string `json:"senderEmail" api:"required"`
	// The unique identifier of the sequence in which the contact will be enrolled.
	SequenceID string `json:"sequenceId" api:"required"`
	// The alias email address used by the sender when enrolling the contact.
	SenderAliasAddress param.Opt[string] `json:"senderAliasAddress,omitzero"`
	paramObj
}

func (r PublicSequenceEnrollmentRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow PublicSequenceEnrollmentRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PublicSequenceEnrollmentRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PublicSequenceEnrollmentResponse struct {
	// The unique identifier for the sequence enrollment.
	ID string `json:"id" api:"required"`
	// The date and time when the contact was enrolled in the sequence.
	EnrolledAt time.Time `json:"enrolledAt" api:"required" format:"date-time"`
	// The identifier of the user who enrolled the contact in the sequence.
	EnrolledBy string `json:"enrolledBy" api:"required"`
	// The email address of the user who enrolled the contact in the sequence.
	EnrolledByEmail string `json:"enrolledByEmail" api:"required"`
	// The unique identifier of the sequence in which the contact is enrolled.
	SequenceID string `json:"sequenceId" api:"required"`
	// The name of the sequence in which the contact is enrolled.
	SequenceName string `json:"sequenceName" api:"required"`
	// The email address of the contact enrolled in the sequence.
	ToEmail string `json:"toEmail" api:"required"`
	// The date and time when the sequence enrollment was last updated.
	UpdatedAt time.Time `json:"updatedAt" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		EnrolledAt      respjson.Field
		EnrolledBy      respjson.Field
		EnrolledByEmail respjson.Field
		SequenceID      respjson.Field
		SequenceName    respjson.Field
		ToEmail         respjson.Field
		UpdatedAt       respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PublicSequenceEnrollmentResponse) RawJSON() string { return r.JSON.raw }
func (r *PublicSequenceEnrollmentResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PublicSequenceLiteResponse struct {
	// The unique identifier of the sequence.
	ID string `json:"id" api:"required"`
	// The date and time when the sequence was created.
	CreatedAt time.Time `json:"createdAt" api:"required" format:"date-time"`
	// The name of the sequence.
	Name string `json:"name" api:"required"`
	// The date and time when the sequence was last updated.
	UpdatedAt time.Time `json:"updatedAt" api:"required" format:"date-time"`
	// The ID of the user associated with the sequence.
	UserID string `json:"userId" api:"required"`
	// The ID of the folder containing the sequence.
	FolderID string `json:"folderId"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		Name        respjson.Field
		UpdatedAt   respjson.Field
		UserID      respjson.Field
		FolderID    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PublicSequenceLiteResponse) RawJSON() string { return r.JSON.raw }
func (r *PublicSequenceLiteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PublicSequenceResponse struct {
	// The unique identifier for the sequence.
	ID string `json:"id" api:"required"`
	// The date and time when the sequence was created.
	CreatedAt time.Time `json:"createdAt" api:"required" format:"date-time"`
	// An array of dependencies for the sequence steps, each represented as a
	// PublicSequenceStepDependencyResponse object.
	Dependencies []PublicSequenceStepDependencyResponse `json:"dependencies" api:"required"`
	// The name of the sequence.
	Name string `json:"name" api:"required"`
	// An array of steps included in the sequence, each represented by a
	// PublicSequenceStepResponse object.
	Steps []PublicSequenceStepResponse `json:"steps" api:"required"`
	// The date and time when the sequence was last updated.
	UpdatedAt time.Time `json:"updatedAt" api:"required" format:"date-time"`
	// The ID of the user associated with the sequence.
	UserID string `json:"userId" api:"required"`
	// The identifier of the folder containing the sequence.
	FolderID string                         `json:"folderId"`
	Settings PublicSequenceSettingsResponse `json:"settings"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		CreatedAt    respjson.Field
		Dependencies respjson.Field
		Name         respjson.Field
		Steps        respjson.Field
		UpdatedAt    respjson.Field
		UserID       respjson.Field
		FolderID     respjson.Field
		Settings     respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PublicSequenceResponse) RawJSON() string { return r.JSON.raw }
func (r *PublicSequenceResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PublicSequenceSettingsResponse struct {
	// The unique identifier for the sequence settings.
	ID string `json:"id" api:"required"`
	// The timestamp of when the sequence settings were created.
	CreatedAt time.Time `json:"createdAt" api:"required" format:"date-time"`
	// Specifies the days on which follow-up actions are allowed.
	//
	// Any of "BUSINESS_DAYS", "EVERYDAY", "WEEKDAYS_ONLY".
	EligibleFollowUpDays PublicSequenceSettingsResponseEligibleFollowUpDays `json:"eligibleFollowUpDays" api:"required"`
	// Indicates whether individual task reminders are enabled.
	IndividualTaskRemindersEnabled bool `json:"individualTaskRemindersEnabled" api:"required"`
	// (deprecated) Defines the unenrollment strategy, with accepted values being
	// ACCOUNT_BASED or LEAD_BASED. If ACCOUNT_BASED is used, all contacts associated
	// with the same company will be unenrolled if one contact meets any of the
	// unenrollment criteria.
	//
	// Any of "ACCOUNT_BASED", "LEAD_BASED".
	SellingStrategy PublicSequenceSettingsResponseSellingStrategy `json:"sellingStrategy" api:"required"`
	// Indicates the end minute of the time window during which automated emails can be
	// sent.
	SendWindowEndMinute int64 `json:"sendWindowEndMinute" api:"required"`
	// Indicates the start minute of the time window during which automated emails can
	// be sent.
	SendWindowStartMinute int64 `json:"sendWindowStartMinute" api:"required"`
	// Specifies the minute of day at which task reminders are triggered.
	TaskReminderMinute int64 `json:"taskReminderMinute" api:"required"`
	// The timestamp of when the sequence settings were last updated.
	UpdatedAt time.Time `json:"updatedAt" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                             respjson.Field
		CreatedAt                      respjson.Field
		EligibleFollowUpDays           respjson.Field
		IndividualTaskRemindersEnabled respjson.Field
		SellingStrategy                respjson.Field
		SendWindowEndMinute            respjson.Field
		SendWindowStartMinute          respjson.Field
		TaskReminderMinute             respjson.Field
		UpdatedAt                      respjson.Field
		ExtraFields                    map[string]respjson.Field
		raw                            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PublicSequenceSettingsResponse) RawJSON() string { return r.JSON.raw }
func (r *PublicSequenceSettingsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Specifies the days on which follow-up actions are allowed.
type PublicSequenceSettingsResponseEligibleFollowUpDays string

const (
	PublicSequenceSettingsResponseEligibleFollowUpDaysBusinessDays PublicSequenceSettingsResponseEligibleFollowUpDays = "BUSINESS_DAYS"
	PublicSequenceSettingsResponseEligibleFollowUpDaysEveryday     PublicSequenceSettingsResponseEligibleFollowUpDays = "EVERYDAY"
	PublicSequenceSettingsResponseEligibleFollowUpDaysWeekdaysOnly PublicSequenceSettingsResponseEligibleFollowUpDays = "WEEKDAYS_ONLY"
)

// (deprecated) Defines the unenrollment strategy, with accepted values being
// ACCOUNT_BASED or LEAD_BASED. If ACCOUNT_BASED is used, all contacts associated
// with the same company will be unenrolled if one contact meets any of the
// unenrollment criteria.
type PublicSequenceSettingsResponseSellingStrategy string

const (
	PublicSequenceSettingsResponseSellingStrategyAccountBased PublicSequenceSettingsResponseSellingStrategy = "ACCOUNT_BASED"
	PublicSequenceSettingsResponseSellingStrategyLeadBased    PublicSequenceSettingsResponseSellingStrategy = "LEAD_BASED"
)

type PublicSequenceStepDependencyResponse struct {
	// The unique identifier of the step dependency.
	ID string `json:"id" api:"required"`
	// The date and time when the step dependency was created.
	CreatedAt time.Time `json:"createdAt" api:"required" format:"date-time"`
	// The type of dependency between sequence steps with accepted values being
	// TASK_COMPLETION or MANUAL_PAUSE.
	//
	// Any of "MANUAL_PAUSE", "TASK_COMPLETION".
	DependencyType PublicSequenceStepDependencyResponseDependencyType `json:"dependencyType" api:"required"`
	// The unique identifier of the sequence step that is responsible for creating and
	// resolving this dependency.
	ReliesOnSequenceStepID string `json:"reliesOnSequenceStepId" api:"required"`
	// The order number of the step that is responsible for creating and resolving this
	// dependency.
	ReliesOnStepOrder int64 `json:"reliesOnStepOrder" api:"required"`
	// The unique identifier of the sequence step that requires this dependency.
	RequiredBySequenceStepID string `json:"requiredBySequenceStepId" api:"required"`
	// The order number of the step that requires this dependency.
	RequiredByStepOrder int64 `json:"requiredByStepOrder" api:"required"`
	// The date and time when the step dependency was last updated.
	UpdatedAt time.Time `json:"updatedAt" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                       respjson.Field
		CreatedAt                respjson.Field
		DependencyType           respjson.Field
		ReliesOnSequenceStepID   respjson.Field
		ReliesOnStepOrder        respjson.Field
		RequiredBySequenceStepID respjson.Field
		RequiredByStepOrder      respjson.Field
		UpdatedAt                respjson.Field
		ExtraFields              map[string]respjson.Field
		raw                      string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PublicSequenceStepDependencyResponse) RawJSON() string { return r.JSON.raw }
func (r *PublicSequenceStepDependencyResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The type of dependency between sequence steps with accepted values being
// TASK_COMPLETION or MANUAL_PAUSE.
type PublicSequenceStepDependencyResponseDependencyType string

const (
	PublicSequenceStepDependencyResponseDependencyTypeManualPause    PublicSequenceStepDependencyResponseDependencyType = "MANUAL_PAUSE"
	PublicSequenceStepDependencyResponseDependencyTypeTaskCompletion PublicSequenceStepDependencyResponseDependencyType = "TASK_COMPLETION"
)

type PublicSequenceStepResponse struct {
	// The unique identifier of the sequence step.
	ID string `json:"id" api:"required"`
	// The type of action to be performed in the sequence step.
	//
	// Any of "EMAIL", "FINISH_ENROLLMENT", "TASK".
	ActionType PublicSequenceStepResponseActionType `json:"actionType" api:"required"`
	// The date and time when the sequence step was created.
	CreatedAt time.Time `json:"createdAt" api:"required" format:"date-time"`
	// The delay in milliseconds before the sequence step is executed.
	DelayMillis int64 `json:"delayMillis" api:"required"`
	// The order of the step within the sequence.
	StepOrder int64 `json:"stepOrder" api:"required"`
	// The date and time when the sequence step was last updated.
	UpdatedAt    time.Time                  `json:"updatedAt" api:"required" format:"date-time"`
	EmailPattern PublicEmailPatternResponse `json:"emailPattern"`
	TaskPattern  PublicTaskPatternResponse  `json:"taskPattern"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		ActionType   respjson.Field
		CreatedAt    respjson.Field
		DelayMillis  respjson.Field
		StepOrder    respjson.Field
		UpdatedAt    respjson.Field
		EmailPattern respjson.Field
		TaskPattern  respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PublicSequenceStepResponse) RawJSON() string { return r.JSON.raw }
func (r *PublicSequenceStepResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The type of action to be performed in the sequence step.
type PublicSequenceStepResponseActionType string

const (
	PublicSequenceStepResponseActionTypeEmail            PublicSequenceStepResponseActionType = "EMAIL"
	PublicSequenceStepResponseActionTypeFinishEnrollment PublicSequenceStepResponseActionType = "FINISH_ENROLLMENT"
	PublicSequenceStepResponseActionTypeTask             PublicSequenceStepResponseActionType = "TASK"
)

type PublicTaskPatternResponse struct {
	// The unique identifier for the task pattern.
	ID string `json:"id" api:"required"`
	// The date and time when the task pattern was created.
	CreatedAt time.Time `json:"createdAt" api:"required" format:"date-time"`
	// The priority level assigned to the task.
	//
	// Any of "HIGH", "LOW", "MEDIUM", "NONE".
	TaskPriority PublicTaskPatternResponseTaskPriority `json:"taskPriority" api:"required"`
	// The type of task, such as an email or call.
	//
	// Any of "CALL", "EMAIL", "LINKED_IN_CONNECT", "LINKED_IN_MESSAGE", "MEETING",
	// "TODO".
	TaskType PublicTaskPatternResponseTaskType `json:"taskType" api:"required"`
	// The date and time when the task pattern was last updated.
	UpdatedAt time.Time `json:"updatedAt" api:"required" format:"date-time"`
	// Additional notes or comments associated with the task.
	Notes string `json:"notes"`
	// The identifier for the queue associated with the task.
	QueueID int64 `json:"queueId"`
	// The subject line of the task.
	Subject string `json:"subject"`
	// The identifier for the template used in the task.
	TemplateID int64 `json:"templateId"`
	// The order of the step to which the email thread is related.
	ThreadEmailToStepOrder int64 `json:"threadEmailToStepOrder"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                     respjson.Field
		CreatedAt              respjson.Field
		TaskPriority           respjson.Field
		TaskType               respjson.Field
		UpdatedAt              respjson.Field
		Notes                  respjson.Field
		QueueID                respjson.Field
		Subject                respjson.Field
		TemplateID             respjson.Field
		ThreadEmailToStepOrder respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PublicTaskPatternResponse) RawJSON() string { return r.JSON.raw }
func (r *PublicTaskPatternResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The priority level assigned to the task.
type PublicTaskPatternResponseTaskPriority string

const (
	PublicTaskPatternResponseTaskPriorityHigh   PublicTaskPatternResponseTaskPriority = "HIGH"
	PublicTaskPatternResponseTaskPriorityLow    PublicTaskPatternResponseTaskPriority = "LOW"
	PublicTaskPatternResponseTaskPriorityMedium PublicTaskPatternResponseTaskPriority = "MEDIUM"
	PublicTaskPatternResponseTaskPriorityNone   PublicTaskPatternResponseTaskPriority = "NONE"
)

// The type of task, such as an email or call.
type PublicTaskPatternResponseTaskType string

const (
	PublicTaskPatternResponseTaskTypeCall            PublicTaskPatternResponseTaskType = "CALL"
	PublicTaskPatternResponseTaskTypeEmail           PublicTaskPatternResponseTaskType = "EMAIL"
	PublicTaskPatternResponseTaskTypeLinkedInConnect PublicTaskPatternResponseTaskType = "LINKED_IN_CONNECT"
	PublicTaskPatternResponseTaskTypeLinkedInMessage PublicTaskPatternResponseTaskType = "LINKED_IN_MESSAGE"
	PublicTaskPatternResponseTaskTypeMeeting         PublicTaskPatternResponseTaskType = "MEETING"
	PublicTaskPatternResponseTaskTypeTodo            PublicTaskPatternResponseTaskType = "TODO"
)

type SequenceListParams struct {
	UserID string `query:"userId" api:"required" json:"-"`
	// The paging cursor token of the last successfully read resource will be returned
	// as the `paging.next.after` JSON property of a paged response containing more
	// results.
	After param.Opt[string] `query:"after,omitzero" json:"-"`
	// The maximum number of results to display per page.
	Limit param.Opt[int64]  `query:"limit,omitzero" json:"-"`
	Name  param.Opt[string] `query:"name,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SequenceListParams]'s query parameters as `url.Values`.
func (r SequenceListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SequenceNewEnrollmentParams struct {
	UserID                          string `query:"userId" api:"required" json:"-"`
	PublicSequenceEnrollmentRequest PublicSequenceEnrollmentRequestParam
	paramObj
}

func (r SequenceNewEnrollmentParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.PublicSequenceEnrollmentRequest)
}
func (r *SequenceNewEnrollmentParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [SequenceNewEnrollmentParams]'s query parameters as
// `url.Values`.
func (r SequenceNewEnrollmentParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SequenceGetParams struct {
	UserID string `query:"userId" api:"required" json:"-"`
	paramObj
}

// URLQuery serializes [SequenceGetParams]'s query parameters as `url.Values`.
func (r SequenceGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
