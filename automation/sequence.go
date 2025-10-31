// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package automation

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"
	"time"

	"github.com/stainless-sdks/hubspot-sdk-go/internal/apijson"
	"github.com/stainless-sdks/hubspot-sdk-go/internal/requestconfig"
	"github.com/stainless-sdks/hubspot-sdk-go/option"
	"github.com/stainless-sdks/hubspot-sdk-go/packages/param"
	"github.com/stainless-sdks/hubspot-sdk-go/packages/respjson"
	"github.com/stainless-sdks/hubspot-sdk-go/shared"
)

// SequenceService contains methods and other services that help with interacting
// with the Hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSequenceService] method instead.
type SequenceService struct {
	Options     []option.RequestOption
	Enrollments SequenceEnrollmentService
}

// NewSequenceService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewSequenceService(opts ...option.RequestOption) (r SequenceService) {
	r = SequenceService{}
	r.Options = opts
	r.Enrollments = NewSequenceEnrollmentService(opts...)
	return
}

// Retrieve a list of sequences that belong to a specific user.
func (r *SequenceService) List(ctx context.Context, opts ...option.RequestOption) (res *CollectionResponseWithTotalPublicSequenceLiteResponseForwardPaging, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "automation/v4/sequences/"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Retrieve details of a specific sequence by its ID.
func (r *SequenceService) Get(ctx context.Context, sequenceID string, opts ...option.RequestOption) (res *PublicSequenceResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if sequenceID == "" {
		err = errors.New("missing required sequenceId parameter")
		return
	}
	path := fmt.Sprintf("automation/v4/sequences/%s", sequenceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type CollectionResponseWithTotalPublicSequenceLiteResponseForwardPaging struct {
	Results []PublicSequenceLiteResponse `json:"results,required"`
	Total   int64                        `json:"total,required"`
	Paging  shared.ForwardPaging         `json:"paging"`
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
func (r CollectionResponseWithTotalPublicSequenceLiteResponseForwardPaging) RawJSON() string {
	return r.JSON.raw
}
func (r *CollectionResponseWithTotalPublicSequenceLiteResponseForwardPaging) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmailSettingsResponse struct {
	// Any of "ALL", "NONE".
	Criteria EmailSettingsResponseCriteria `json:"criteria,required"`
	// Any of "LEAD_BASED", "ACCOUNT_BASED".
	SellingStrategy EmailSettingsResponseSellingStrategy `json:"sellingStrategy,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Criteria        respjson.Field
		SellingStrategy respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmailSettingsResponse) RawJSON() string { return r.JSON.raw }
func (r *EmailSettingsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmailSettingsResponseCriteria string

const (
	EmailSettingsResponseCriteriaAll  EmailSettingsResponseCriteria = "ALL"
	EmailSettingsResponseCriteriaNone EmailSettingsResponseCriteria = "NONE"
)

type EmailSettingsResponseSellingStrategy string

const (
	EmailSettingsResponseSellingStrategyLeadBased    EmailSettingsResponseSellingStrategy = "LEAD_BASED"
	EmailSettingsResponseSellingStrategyAccountBased EmailSettingsResponseSellingStrategy = "ACCOUNT_BASED"
)

type MeetingSettingsResponse struct {
	// Any of "ALL", "NONE".
	Criteria MeetingSettingsResponseCriteria `json:"criteria,required"`
	// Any of "LEAD_BASED", "ACCOUNT_BASED".
	SellingStrategy MeetingSettingsResponseSellingStrategy `json:"sellingStrategy,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Criteria        respjson.Field
		SellingStrategy respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MeetingSettingsResponse) RawJSON() string { return r.JSON.raw }
func (r *MeetingSettingsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MeetingSettingsResponseCriteria string

const (
	MeetingSettingsResponseCriteriaAll  MeetingSettingsResponseCriteria = "ALL"
	MeetingSettingsResponseCriteriaNone MeetingSettingsResponseCriteria = "NONE"
)

type MeetingSettingsResponseSellingStrategy string

const (
	MeetingSettingsResponseSellingStrategyLeadBased    MeetingSettingsResponseSellingStrategy = "LEAD_BASED"
	MeetingSettingsResponseSellingStrategyAccountBased MeetingSettingsResponseSellingStrategy = "ACCOUNT_BASED"
)

type PublicEmailPatternResponse struct {
	ID                     string    `json:"id,required"`
	CreatedAt              time.Time `json:"createdAt,required" format:"date-time"`
	TemplateID             string    `json:"templateId,required"`
	UpdatedAt              time.Time `json:"updatedAt,required" format:"date-time"`
	ThreadEmailToStepOrder int64     `json:"threadEmailToStepOrder"`
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
	ID         string    `json:"id,required"`
	EnrolledAt time.Time `json:"enrolledAt,required" format:"date-time"`
	ToEmail    string    `json:"toEmail,required"`
	UpdatedAt  time.Time `json:"updatedAt,required" format:"date-time"`
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
	ContactID          string            `json:"contactId,required"`
	SenderEmail        string            `json:"senderEmail,required"`
	SequenceID         string            `json:"sequenceId,required"`
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
	ID              string    `json:"id,required"`
	EnrolledAt      time.Time `json:"enrolledAt,required" format:"date-time"`
	EnrolledBy      string    `json:"enrolledBy,required"`
	EnrolledByEmail string    `json:"enrolledByEmail,required"`
	SequenceID      string    `json:"sequenceId,required"`
	SequenceName    string    `json:"sequenceName,required"`
	ToEmail         string    `json:"toEmail,required"`
	UpdatedAt       time.Time `json:"updatedAt,required" format:"date-time"`
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
	ID        string    `json:"id,required"`
	CreatedAt time.Time `json:"createdAt,required" format:"date-time"`
	Name      string    `json:"name,required"`
	UpdatedAt time.Time `json:"updatedAt,required" format:"date-time"`
	UserID    string    `json:"userId,required"`
	FolderID  string    `json:"folderId"`
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
	ID           string                                 `json:"id,required"`
	CreatedAt    time.Time                              `json:"createdAt,required" format:"date-time"`
	Dependencies []PublicSequenceStepDependencyResponse `json:"dependencies,required"`
	Name         string                                 `json:"name,required"`
	Steps        []PublicSequenceStepResponse           `json:"steps,required"`
	UpdatedAt    time.Time                              `json:"updatedAt,required" format:"date-time"`
	UserID       string                                 `json:"userId,required"`
	FolderID     string                                 `json:"folderId"`
	Settings     PublicSequenceSettingsResponse         `json:"settings"`
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
	ID                             string                       `json:"id,required"`
	CreatedAt                      time.Time                    `json:"createdAt,required" format:"date-time"`
	EligibleFollowUpDays           string                       `json:"eligibleFollowUpDays,required"`
	IndividualTaskRemindersEnabled bool                         `json:"individualTaskRemindersEnabled,required"`
	SellingStrategy                string                       `json:"sellingStrategy,required"`
	SendWindowEndMinute            int64                        `json:"sendWindowEndMinute,required"`
	SendWindowStartMinute          int64                        `json:"sendWindowStartMinute,required"`
	TaskReminderMinute             int64                        `json:"taskReminderMinute,required"`
	UpdatedAt                      time.Time                    `json:"updatedAt,required" format:"date-time"`
	UnenrollmentSettings           UnenrollmentSettingsResponse `json:"unenrollmentSettings"`
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
		UnenrollmentSettings           respjson.Field
		ExtraFields                    map[string]respjson.Field
		raw                            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PublicSequenceSettingsResponse) RawJSON() string { return r.JSON.raw }
func (r *PublicSequenceSettingsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PublicSequenceStepDependencyResponse struct {
	ID                       string    `json:"id,required"`
	CreatedAt                time.Time `json:"createdAt,required" format:"date-time"`
	DependencyType           string    `json:"dependencyType,required"`
	ReliesOnSequenceStepID   string    `json:"reliesOnSequenceStepId,required"`
	ReliesOnStepOrder        int64     `json:"reliesOnStepOrder,required"`
	RequiredBySequenceStepID string    `json:"requiredBySequenceStepId,required"`
	RequiredByStepOrder      int64     `json:"requiredByStepOrder,required"`
	UpdatedAt                time.Time `json:"updatedAt,required" format:"date-time"`
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

type PublicSequenceStepResponse struct {
	ID           string                     `json:"id,required"`
	ActionType   string                     `json:"actionType,required"`
	CreatedAt    time.Time                  `json:"createdAt,required" format:"date-time"`
	DelayMillis  int64                      `json:"delayMillis,required"`
	StepOrder    int64                      `json:"stepOrder,required"`
	UpdatedAt    time.Time                  `json:"updatedAt,required" format:"date-time"`
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

type PublicTaskPatternResponse struct {
	ID                     string    `json:"id,required"`
	CreatedAt              time.Time `json:"createdAt,required" format:"date-time"`
	TaskPriority           string    `json:"taskPriority,required"`
	TaskType               string    `json:"taskType,required"`
	UpdatedAt              time.Time `json:"updatedAt,required" format:"date-time"`
	Notes                  string    `json:"notes"`
	QueueID                int64     `json:"queueId"`
	Subject                string    `json:"subject"`
	TemplateID             int64     `json:"templateId"`
	ThreadEmailToStepOrder int64     `json:"threadEmailToStepOrder"`
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

type UnenrollmentSettingsResponse struct {
	EmailSettings   EmailSettingsResponse   `json:"emailSettings,required"`
	MeetingSettings MeetingSettingsResponse `json:"meetingSettings,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		EmailSettings   respjson.Field
		MeetingSettings respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UnenrollmentSettingsResponse) RawJSON() string { return r.JSON.raw }
func (r *UnenrollmentSettingsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
