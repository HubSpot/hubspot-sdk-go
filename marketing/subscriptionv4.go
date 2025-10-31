// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package marketing

import (
	"time"

	"github.com/stainless-sdks/hubspot-sdk-go/internal/apijson"
	"github.com/stainless-sdks/hubspot-sdk-go/option"
	"github.com/stainless-sdks/hubspot-sdk-go/packages/param"
	"github.com/stainless-sdks/hubspot-sdk-go/packages/respjson"
	"github.com/stainless-sdks/hubspot-sdk-go/shared"
)

// SubscriptionV4Service contains methods and other services that help with
// interacting with the Hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSubscriptionV4Service] method instead.
type SubscriptionV4Service struct {
	Options     []option.RequestOption
	Definitions SubscriptionV4DefinitionService
	Links       SubscriptionV4LinkService
	Statuses    SubscriptionV4StatusService
}

// NewSubscriptionV4Service generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewSubscriptionV4Service(opts ...option.RequestOption) (r SubscriptionV4Service) {
	r = SubscriptionV4Service{}
	r.Options = opts
	r.Definitions = NewSubscriptionV4DefinitionService(opts...)
	r.Links = NewSubscriptionV4LinkService(opts...)
	r.Statuses = NewSubscriptionV4StatusService(opts...)
	return
}

type ActionResponseWithResultsPublicStatus struct {
	// The date and time when the operation was completed.
	CompletedAt time.Time `json:"completedAt,required" format:"date-time"`
	// An array of results from the operation.
	Results []PublicStatus `json:"results,required"`
	// The date and time when the operation started.
	StartedAt time.Time `json:"startedAt,required" format:"date-time"`
	// Indicates the current status of the operation, with possible values: PENDING,
	// PROCESSING, CANCELED, COMPLETE.
	//
	// Any of "PENDING", "PROCESSING", "CANCELED", "COMPLETE".
	Status ActionResponseWithResultsPublicStatusStatus `json:"status,required"`
	// A list of errors that occurred during the operation.
	Errors []shared.StandardError `json:"errors"`
	// Contains URLs related to the response, such as documentation or resources.
	Links map[string]string `json:"links"`
	// The number of errors that occurred during the operation.
	NumErrors int64 `json:"numErrors"`
	// The date and time when the request was made.
	RequestedAt time.Time `json:"requestedAt" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CompletedAt respjson.Field
		Results     respjson.Field
		StartedAt   respjson.Field
		Status      respjson.Field
		Errors      respjson.Field
		Links       respjson.Field
		NumErrors   respjson.Field
		RequestedAt respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ActionResponseWithResultsPublicStatus) RawJSON() string { return r.JSON.raw }
func (r *ActionResponseWithResultsPublicStatus) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Indicates the current status of the operation, with possible values: PENDING,
// PROCESSING, CANCELED, COMPLETE.
type ActionResponseWithResultsPublicStatusStatus string

const (
	ActionResponseWithResultsPublicStatusStatusPending    ActionResponseWithResultsPublicStatusStatus = "PENDING"
	ActionResponseWithResultsPublicStatusStatusProcessing ActionResponseWithResultsPublicStatusStatus = "PROCESSING"
	ActionResponseWithResultsPublicStatusStatusCanceled   ActionResponseWithResultsPublicStatusStatus = "CANCELED"
	ActionResponseWithResultsPublicStatusStatusComplete   ActionResponseWithResultsPublicStatusStatus = "COMPLETE"
)

type ActionResponseWithResultsPublicWideStatus struct {
	// The date and time when the operation was completed.
	CompletedAt time.Time `json:"completedAt,required" format:"date-time"`
	// An array containing the results of the operation.
	Results []PublicWideStatus `json:"results,required"`
	// The date and time when the operation started.
	StartedAt time.Time `json:"startedAt,required" format:"date-time"`
	// The current status of the operation, which can be PENDING, PROCESSING, CANCELED,
	// or COMPLETE.
	//
	// Any of "PENDING", "PROCESSING", "CANCELED", "COMPLETE".
	Status ActionResponseWithResultsPublicWideStatusStatus `json:"status,required"`
	// An array of error objects detailing any issues encountered during the operation.
	Errors []shared.StandardError `json:"errors"`
	// An object containing related links, where each key is a link name and each value
	// is a URL.
	Links map[string]string `json:"links"`
	// The number of errors encountered during the operation.
	NumErrors int64 `json:"numErrors"`
	// The date and time when the request was made.
	RequestedAt time.Time `json:"requestedAt" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CompletedAt respjson.Field
		Results     respjson.Field
		StartedAt   respjson.Field
		Status      respjson.Field
		Errors      respjson.Field
		Links       respjson.Field
		NumErrors   respjson.Field
		RequestedAt respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ActionResponseWithResultsPublicWideStatus) RawJSON() string { return r.JSON.raw }
func (r *ActionResponseWithResultsPublicWideStatus) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The current status of the operation, which can be PENDING, PROCESSING, CANCELED,
// or COMPLETE.
type ActionResponseWithResultsPublicWideStatusStatus string

const (
	ActionResponseWithResultsPublicWideStatusStatusPending    ActionResponseWithResultsPublicWideStatusStatus = "PENDING"
	ActionResponseWithResultsPublicWideStatusStatusProcessing ActionResponseWithResultsPublicWideStatusStatus = "PROCESSING"
	ActionResponseWithResultsPublicWideStatusStatusCanceled   ActionResponseWithResultsPublicWideStatusStatus = "CANCELED"
	ActionResponseWithResultsPublicWideStatusStatusComplete   ActionResponseWithResultsPublicWideStatusStatus = "COMPLETE"
)

type ActionResponseWithResultsSubscriptionDefinition struct {
	// The date and time when the operation was completed.
	CompletedAt time.Time `json:"completedAt,required" format:"date-time"`
	// An array containing the results of the operation.
	Results []SubscriptionDefinition `json:"results,required"`
	// The date and time when the operation started.
	StartedAt time.Time `json:"startedAt,required" format:"date-time"`
	// The current status of the operation, which can be PENDING, PROCESSING, CANCELED,
	// or COMPLETE.
	//
	// Any of "PENDING", "PROCESSING", "CANCELED", "COMPLETE".
	Status ActionResponseWithResultsSubscriptionDefinitionStatus `json:"status,required"`
	// An array of errors that occurred during the operation.
	Errors []shared.StandardError `json:"errors"`
	// A collection of related links associated with the operation.
	Links map[string]string `json:"links"`
	// The number of errors encountered during the operation.
	NumErrors int64 `json:"numErrors"`
	// The date and time when the operation was requested.
	RequestedAt time.Time `json:"requestedAt" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CompletedAt respjson.Field
		Results     respjson.Field
		StartedAt   respjson.Field
		Status      respjson.Field
		Errors      respjson.Field
		Links       respjson.Field
		NumErrors   respjson.Field
		RequestedAt respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ActionResponseWithResultsSubscriptionDefinition) RawJSON() string { return r.JSON.raw }
func (r *ActionResponseWithResultsSubscriptionDefinition) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The current status of the operation, which can be PENDING, PROCESSING, CANCELED,
// or COMPLETE.
type ActionResponseWithResultsSubscriptionDefinitionStatus string

const (
	ActionResponseWithResultsSubscriptionDefinitionStatusPending    ActionResponseWithResultsSubscriptionDefinitionStatus = "PENDING"
	ActionResponseWithResultsSubscriptionDefinitionStatusProcessing ActionResponseWithResultsSubscriptionDefinitionStatus = "PROCESSING"
	ActionResponseWithResultsSubscriptionDefinitionStatusCanceled   ActionResponseWithResultsSubscriptionDefinitionStatus = "CANCELED"
	ActionResponseWithResultsSubscriptionDefinitionStatusComplete   ActionResponseWithResultsSubscriptionDefinitionStatus = "COMPLETE"
)

// The property Inputs is required.
type BatchInputPublicStatusRequestParam struct {
	Inputs []PublicStatusRequestParam `json:"inputs,omitzero,required"`
	paramObj
}

func (r BatchInputPublicStatusRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow BatchInputPublicStatusRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BatchInputPublicStatusRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BatchResponsePublicBulkOptOutFromAllResponse struct {
	// The date and time when the bulk opt-out operation was completed.
	CompletedAt time.Time `json:"completedAt,required" format:"date-time"`
	// An array containing the results of the bulk opt-out from all communications
	// operation.
	Results []PublicBulkOptOutFromAllResponse `json:"results,required"`
	// The date and time when the bulk opt-out operation began.
	StartedAt time.Time `json:"startedAt,required" format:"date-time"`
	// The current status of the bulk opt-out operation, which can be PENDING,
	// PROCESSING, CANCELED, or COMPLETE.
	//
	// Any of "PENDING", "PROCESSING", "CANCELED", "COMPLETE".
	Status BatchResponsePublicBulkOptOutFromAllResponseStatus `json:"status,required"`
	// An array of error objects detailing any issues encountered during the bulk
	// opt-out operation.
	Errors []shared.StandardError `json:"errors"`
	// A collection of URLs linking to related resources or documentation.
	Links map[string]string `json:"links"`
	// The total number of errors encountered during the bulk opt-out operation.
	NumErrors int64 `json:"numErrors"`
	// The date and time when the bulk opt-out request was made.
	RequestedAt time.Time `json:"requestedAt" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CompletedAt respjson.Field
		Results     respjson.Field
		StartedAt   respjson.Field
		Status      respjson.Field
		Errors      respjson.Field
		Links       respjson.Field
		NumErrors   respjson.Field
		RequestedAt respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BatchResponsePublicBulkOptOutFromAllResponse) RawJSON() string { return r.JSON.raw }
func (r *BatchResponsePublicBulkOptOutFromAllResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The current status of the bulk opt-out operation, which can be PENDING,
// PROCESSING, CANCELED, or COMPLETE.
type BatchResponsePublicBulkOptOutFromAllResponseStatus string

const (
	BatchResponsePublicBulkOptOutFromAllResponseStatusPending    BatchResponsePublicBulkOptOutFromAllResponseStatus = "PENDING"
	BatchResponsePublicBulkOptOutFromAllResponseStatusProcessing BatchResponsePublicBulkOptOutFromAllResponseStatus = "PROCESSING"
	BatchResponsePublicBulkOptOutFromAllResponseStatusCanceled   BatchResponsePublicBulkOptOutFromAllResponseStatus = "CANCELED"
	BatchResponsePublicBulkOptOutFromAllResponseStatusComplete   BatchResponsePublicBulkOptOutFromAllResponseStatus = "COMPLETE"
)

type BatchResponsePublicStatus struct {
	// The date and time when the batch operation was completed.
	CompletedAt time.Time `json:"completedAt,required" format:"date-time"`
	// An array containing the results of the batch operation.
	Results []PublicStatus `json:"results,required"`
	// The date and time when the batch operation started.
	StartedAt time.Time `json:"startedAt,required" format:"date-time"`
	// The current status of the batch operation, which can be PENDING, PROCESSING,
	// CANCELED, or COMPLETE.
	//
	// Any of "PENDING", "PROCESSING", "CANCELED", "COMPLETE".
	Status BatchResponsePublicStatusStatus `json:"status,required"`
	// An array of error objects detailing any issues encountered.
	Errors []shared.StandardError `json:"errors"`
	// URLs linking to related resources or documentation.
	Links map[string]string `json:"links"`
	// The number of errors encountered during the batch operation.
	NumErrors int64 `json:"numErrors"`
	// The date and time when the request was made.
	RequestedAt time.Time `json:"requestedAt" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CompletedAt respjson.Field
		Results     respjson.Field
		StartedAt   respjson.Field
		Status      respjson.Field
		Errors      respjson.Field
		Links       respjson.Field
		NumErrors   respjson.Field
		RequestedAt respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BatchResponsePublicStatus) RawJSON() string { return r.JSON.raw }
func (r *BatchResponsePublicStatus) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The current status of the batch operation, which can be PENDING, PROCESSING,
// CANCELED, or COMPLETE.
type BatchResponsePublicStatusStatus string

const (
	BatchResponsePublicStatusStatusPending    BatchResponsePublicStatusStatus = "PENDING"
	BatchResponsePublicStatusStatusProcessing BatchResponsePublicStatusStatus = "PROCESSING"
	BatchResponsePublicStatusStatusCanceled   BatchResponsePublicStatusStatus = "CANCELED"
	BatchResponsePublicStatusStatusComplete   BatchResponsePublicStatusStatus = "COMPLETE"
)

type BatchResponsePublicStatusBulkResponse struct {
	// The date and time when the batch process was completed.
	CompletedAt time.Time `json:"completedAt,required" format:"date-time"`
	// The array of results from the batch process, each containing subscription status
	// information.
	Results []PublicStatusBulkResponse `json:"results,required"`
	// The date and time when the batch process began.
	StartedAt time.Time `json:"startedAt,required" format:"date-time"`
	// The current status of the batch process, with possible values: PENDING,
	// PROCESSING, CANCELED, COMPLETE.
	//
	// Any of "PENDING", "PROCESSING", "CANCELED", "COMPLETE".
	Status BatchResponsePublicStatusBulkResponseStatus `json:"status,required"`
	// A collection of related links associated with the batch response.
	Links map[string]string `json:"links"`
	// The date and time when the batch request was made.
	RequestedAt time.Time `json:"requestedAt" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CompletedAt respjson.Field
		Results     respjson.Field
		StartedAt   respjson.Field
		Status      respjson.Field
		Links       respjson.Field
		RequestedAt respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BatchResponsePublicStatusBulkResponse) RawJSON() string { return r.JSON.raw }
func (r *BatchResponsePublicStatusBulkResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The current status of the batch process, with possible values: PENDING,
// PROCESSING, CANCELED, COMPLETE.
type BatchResponsePublicStatusBulkResponseStatus string

const (
	BatchResponsePublicStatusBulkResponseStatusPending    BatchResponsePublicStatusBulkResponseStatus = "PENDING"
	BatchResponsePublicStatusBulkResponseStatusProcessing BatchResponsePublicStatusBulkResponseStatus = "PROCESSING"
	BatchResponsePublicStatusBulkResponseStatusCanceled   BatchResponsePublicStatusBulkResponseStatus = "CANCELED"
	BatchResponsePublicStatusBulkResponseStatusComplete   BatchResponsePublicStatusBulkResponseStatus = "COMPLETE"
)

type BatchResponsePublicWideStatusBulkResponse struct {
	// The date and time when the batch process was completed.
	CompletedAt time.Time `json:"completedAt,required" format:"date-time"`
	// The array of results from the batch process, each containing subscription status
	// information.
	Results []PublicWideStatusBulkResponse `json:"results,required"`
	// The date and time when the batch process began.
	StartedAt time.Time `json:"startedAt,required" format:"date-time"`
	// The current status of the batch process, with possible values: PENDING,
	// PROCESSING, CANCELED, COMPLETE.
	//
	// Any of "PENDING", "PROCESSING", "CANCELED", "COMPLETE".
	Status BatchResponsePublicWideStatusBulkResponseStatus `json:"status,required"`
	// A collection of related links associated with the batch response.
	Links map[string]string `json:"links"`
	// The date and time when the batch request was made.
	RequestedAt time.Time `json:"requestedAt" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CompletedAt respjson.Field
		Results     respjson.Field
		StartedAt   respjson.Field
		Status      respjson.Field
		Links       respjson.Field
		RequestedAt respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BatchResponsePublicWideStatusBulkResponse) RawJSON() string { return r.JSON.raw }
func (r *BatchResponsePublicWideStatusBulkResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The current status of the batch process, with possible values: PENDING,
// PROCESSING, CANCELED, COMPLETE.
type BatchResponsePublicWideStatusBulkResponseStatus string

const (
	BatchResponsePublicWideStatusBulkResponseStatusPending    BatchResponsePublicWideStatusBulkResponseStatus = "PENDING"
	BatchResponsePublicWideStatusBulkResponseStatusProcessing BatchResponsePublicWideStatusBulkResponseStatus = "PROCESSING"
	BatchResponsePublicWideStatusBulkResponseStatusCanceled   BatchResponsePublicWideStatusBulkResponseStatus = "CANCELED"
	BatchResponsePublicWideStatusBulkResponseStatusComplete   BatchResponsePublicWideStatusBulkResponseStatus = "COMPLETE"
)

// The property SubscriberIDString is required.
type LinkGenerationRequestParam struct {
	SubscriberIDString string            `json:"subscriberIdString,required"`
	Language           param.Opt[string] `json:"language,omitzero"`
	SubscriptionID     param.Opt[int64]  `json:"subscriptionId,omitzero"`
	paramObj
}

func (r LinkGenerationRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow LinkGenerationRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LinkGenerationRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LinkGenerationResponse struct {
	ManagePreferencesURL string `json:"managePreferencesUrl,required"`
	SubscriberIDString   string `json:"subscriberIdString,required"`
	UnsubscribeAllURL    string `json:"unsubscribeAllUrl,required"`
	UnsubscribeSingleURL string `json:"unsubscribeSingleUrl"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ManagePreferencesURL respjson.Field
		SubscriberIDString   respjson.Field
		UnsubscribeAllURL    respjson.Field
		UnsubscribeSingleURL respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LinkGenerationResponse) RawJSON() string { return r.JSON.raw }
func (r *LinkGenerationResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Channel, StatusState, SubscriptionID are required.
type PartialPublicStatusRequestParam struct {
	// The type of communication channel, with 'EMAIL' as the only supported option.
	//
	// Any of "EMAIL".
	Channel PartialPublicStatusRequestChannel `json:"channel,omitzero,required"`
	// The current subscription status of the contact, which can be 'SUBSCRIBED',
	// 'UNSUBSCRIBED', or 'NOT_SPECIFIED'.
	//
	// Any of "SUBSCRIBED", "UNSUBSCRIBED", "NOT_SPECIFIED".
	StatusState PartialPublicStatusRequestStatusState `json:"statusState,omitzero,required"`
	// The unique identifier of the subscription to be updated.
	SubscriptionID int64 `json:"subscriptionId,required"`
	// An explanation for the legal basis used for communication.
	LegalBasisExplanation param.Opt[string] `json:"legalBasisExplanation,omitzero"`
	// The legal basis for communication, with options including
	// 'LEGITIMATE_INTEREST_PQL', 'LEGITIMATE_INTEREST_CLIENT',
	// 'PERFORMANCE_OF_CONTRACT', 'CONSENT_WITH_NOTICE', 'NON_GDPR',
	// 'PROCESS_AND_STORE', and 'LEGITIMATE_INTEREST_OTHER'.
	//
	// Any of "LEGITIMATE_INTEREST_PQL", "LEGITIMATE_INTEREST_CLIENT",
	// "PERFORMANCE_OF_CONTRACT", "CONSENT_WITH_NOTICE", "NON_GDPR",
	// "PROCESS_AND_STORE", "LEGITIMATE_INTEREST_OTHER".
	LegalBasis PartialPublicStatusRequestLegalBasis `json:"legalBasis,omitzero"`
	paramObj
}

func (r PartialPublicStatusRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow PartialPublicStatusRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PartialPublicStatusRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The type of communication channel, with 'EMAIL' as the only supported option.
type PartialPublicStatusRequestChannel string

const (
	PartialPublicStatusRequestChannelEmail PartialPublicStatusRequestChannel = "EMAIL"
)

// The current subscription status of the contact, which can be 'SUBSCRIBED',
// 'UNSUBSCRIBED', or 'NOT_SPECIFIED'.
type PartialPublicStatusRequestStatusState string

const (
	PartialPublicStatusRequestStatusStateSubscribed   PartialPublicStatusRequestStatusState = "SUBSCRIBED"
	PartialPublicStatusRequestStatusStateUnsubscribed PartialPublicStatusRequestStatusState = "UNSUBSCRIBED"
	PartialPublicStatusRequestStatusStateNotSpecified PartialPublicStatusRequestStatusState = "NOT_SPECIFIED"
)

// The legal basis for communication, with options including
// 'LEGITIMATE_INTEREST_PQL', 'LEGITIMATE_INTEREST_CLIENT',
// 'PERFORMANCE_OF_CONTRACT', 'CONSENT_WITH_NOTICE', 'NON_GDPR',
// 'PROCESS_AND_STORE', and 'LEGITIMATE_INTEREST_OTHER'.
type PartialPublicStatusRequestLegalBasis string

const (
	PartialPublicStatusRequestLegalBasisLegitimateInterestPql    PartialPublicStatusRequestLegalBasis = "LEGITIMATE_INTEREST_PQL"
	PartialPublicStatusRequestLegalBasisLegitimateInterestClient PartialPublicStatusRequestLegalBasis = "LEGITIMATE_INTEREST_CLIENT"
	PartialPublicStatusRequestLegalBasisPerformanceOfContract    PartialPublicStatusRequestLegalBasis = "PERFORMANCE_OF_CONTRACT"
	PartialPublicStatusRequestLegalBasisConsentWithNotice        PartialPublicStatusRequestLegalBasis = "CONSENT_WITH_NOTICE"
	PartialPublicStatusRequestLegalBasisNonGdpr                  PartialPublicStatusRequestLegalBasis = "NON_GDPR"
	PartialPublicStatusRequestLegalBasisProcessAndStore          PartialPublicStatusRequestLegalBasis = "PROCESS_AND_STORE"
	PartialPublicStatusRequestLegalBasisLegitimateInterestOther  PartialPublicStatusRequestLegalBasis = "LEGITIMATE_INTEREST_OTHER"
)

type PublicBulkOptOutFromAllResponse struct {
	// The email address of the contact.
	SubscriberIDString string `json:"subscriberIdString,required"`
	// An array of subscription status objects for the contact.
	Statuses []PublicStatus `json:"statuses"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		SubscriberIDString respjson.Field
		Statuses           respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PublicBulkOptOutFromAllResponse) RawJSON() string { return r.JSON.raw }
func (r *PublicBulkOptOutFromAllResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PublicStatus struct {
	// The type of communication channel, with 'EMAIL' as the only supported option.
	//
	// Any of "EMAIL".
	Channel PublicStatusChannel `json:"channel,required"`
	// The origin or method through which the subscription status was set.
	Source string `json:"source,required"`
	// The current subscription status of the contact, which can be 'SUBSCRIBED',
	// 'UNSUBSCRIBED', or 'NOT_SPECIFIED'.
	//
	// Any of "SUBSCRIBED", "UNSUBSCRIBED", "NOT_SPECIFIED".
	Status PublicStatusStatus `json:"status,required"`
	// The contact's email address.
	SubscriberIDString string `json:"subscriberIdString,required"`
	// The unique identifier of the subscription.
	SubscriptionID int64 `json:"subscriptionId,required"`
	// The date and time when the subscription status was last updated.
	Timestamp time.Time `json:"timestamp,required" format:"date-time"`
	// The ID of the business unit associated with the subscription.
	BusinessUnitID int64 `json:"businessUnitId"`
	// The legal basis for communication, with options including
	// 'LEGITIMATE_INTEREST_PQL', 'LEGITIMATE_INTEREST_CLIENT',
	// 'PERFORMANCE_OF_CONTRACT', 'CONSENT_WITH_NOTICE', 'NON_GDPR',
	// 'PROCESS_AND_STORE', and 'LEGITIMATE_INTEREST_OTHER'.
	//
	// Any of "LEGITIMATE_INTEREST_PQL", "LEGITIMATE_INTEREST_CLIENT",
	// "PERFORMANCE_OF_CONTRACT", "CONSENT_WITH_NOTICE", "NON_GDPR",
	// "PROCESS_AND_STORE", "LEGITIMATE_INTEREST_OTHER".
	LegalBasis PublicStatusLegalBasis `json:"legalBasis"`
	// An explanation for the legal basis used for communication.
	LegalBasisExplanation string `json:"legalBasisExplanation"`
	// The reason for the successful change in subscription status, such as
	// 'RESUBSCRIBE_OCCURRED' or 'NO_STATUS_CHANGE'.
	//
	// Any of "RESUBSCRIBE_OCCURRED", "NO_STATUS_CHANGE",
	// "UNSUBSCRIBE_FROM_ALL_OCCURRED", "REQUESTED_CHANGE_OCCURRED".
	SetStatusSuccessReason PublicStatusSetStatusSuccessReason `json:"setStatusSuccessReason"`
	// The name of the subscription.
	SubscriptionName string `json:"subscriptionName"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Channel                respjson.Field
		Source                 respjson.Field
		Status                 respjson.Field
		SubscriberIDString     respjson.Field
		SubscriptionID         respjson.Field
		Timestamp              respjson.Field
		BusinessUnitID         respjson.Field
		LegalBasis             respjson.Field
		LegalBasisExplanation  respjson.Field
		SetStatusSuccessReason respjson.Field
		SubscriptionName       respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PublicStatus) RawJSON() string { return r.JSON.raw }
func (r *PublicStatus) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The type of communication channel, with 'EMAIL' as the only supported option.
type PublicStatusChannel string

const (
	PublicStatusChannelEmail PublicStatusChannel = "EMAIL"
)

// The current subscription status of the contact, which can be 'SUBSCRIBED',
// 'UNSUBSCRIBED', or 'NOT_SPECIFIED'.
type PublicStatusStatus string

const (
	PublicStatusStatusSubscribed   PublicStatusStatus = "SUBSCRIBED"
	PublicStatusStatusUnsubscribed PublicStatusStatus = "UNSUBSCRIBED"
	PublicStatusStatusNotSpecified PublicStatusStatus = "NOT_SPECIFIED"
)

// The legal basis for communication, with options including
// 'LEGITIMATE_INTEREST_PQL', 'LEGITIMATE_INTEREST_CLIENT',
// 'PERFORMANCE_OF_CONTRACT', 'CONSENT_WITH_NOTICE', 'NON_GDPR',
// 'PROCESS_AND_STORE', and 'LEGITIMATE_INTEREST_OTHER'.
type PublicStatusLegalBasis string

const (
	PublicStatusLegalBasisLegitimateInterestPql    PublicStatusLegalBasis = "LEGITIMATE_INTEREST_PQL"
	PublicStatusLegalBasisLegitimateInterestClient PublicStatusLegalBasis = "LEGITIMATE_INTEREST_CLIENT"
	PublicStatusLegalBasisPerformanceOfContract    PublicStatusLegalBasis = "PERFORMANCE_OF_CONTRACT"
	PublicStatusLegalBasisConsentWithNotice        PublicStatusLegalBasis = "CONSENT_WITH_NOTICE"
	PublicStatusLegalBasisNonGdpr                  PublicStatusLegalBasis = "NON_GDPR"
	PublicStatusLegalBasisProcessAndStore          PublicStatusLegalBasis = "PROCESS_AND_STORE"
	PublicStatusLegalBasisLegitimateInterestOther  PublicStatusLegalBasis = "LEGITIMATE_INTEREST_OTHER"
)

// The reason for the successful change in subscription status, such as
// 'RESUBSCRIBE_OCCURRED' or 'NO_STATUS_CHANGE'.
type PublicStatusSetStatusSuccessReason string

const (
	PublicStatusSetStatusSuccessReasonResubscribeOccurred        PublicStatusSetStatusSuccessReason = "RESUBSCRIBE_OCCURRED"
	PublicStatusSetStatusSuccessReasonNoStatusChange             PublicStatusSetStatusSuccessReason = "NO_STATUS_CHANGE"
	PublicStatusSetStatusSuccessReasonUnsubscribeFromAllOccurred PublicStatusSetStatusSuccessReason = "UNSUBSCRIBE_FROM_ALL_OCCURRED"
	PublicStatusSetStatusSuccessReasonRequestedChangeOccurred    PublicStatusSetStatusSuccessReason = "REQUESTED_CHANGE_OCCURRED"
)

type PublicStatusBulkResponse struct {
	// An array of subscription status objects for the contact.
	Statuses []PublicStatus `json:"statuses,required"`
	// The email address of the contact.
	SubscriberIDString string `json:"subscriberIdString,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Statuses           respjson.Field
		SubscriberIDString respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PublicStatusBulkResponse) RawJSON() string { return r.JSON.raw }
func (r *PublicStatusBulkResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Channel, StatusState, SubscriberIDString, SubscriptionID are
// required.
type PublicStatusRequestParam struct {
	// The type of communication channel. Currently, only `EMAIL` is supported.
	//
	// Any of "EMAIL".
	Channel PublicStatusRequestChannel `json:"channel,omitzero,required"`
	// The status of the contact's subscription.
	//
	// Any of "SUBSCRIBED", "UNSUBSCRIBED", "NOT_SPECIFIED".
	StatusState PublicStatusRequestStatusState `json:"statusState,omitzero,required"`
	// The contact's email address.
	SubscriberIDString string `json:"subscriberIdString,required"`
	// The ID of the subscription to update.
	SubscriptionID int64 `json:"subscriptionId,required"`
	// The explanation for the legal basis.
	LegalBasisExplanation param.Opt[string] `json:"legalBasisExplanation,omitzero"`
	// The legal basis for communication.
	//
	// Any of "LEGITIMATE_INTEREST_PQL", "LEGITIMATE_INTEREST_CLIENT",
	// "PERFORMANCE_OF_CONTRACT", "CONSENT_WITH_NOTICE", "NON_GDPR",
	// "PROCESS_AND_STORE", "LEGITIMATE_INTEREST_OTHER".
	LegalBasis PublicStatusRequestLegalBasis `json:"legalBasis,omitzero"`
	paramObj
}

func (r PublicStatusRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow PublicStatusRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PublicStatusRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The type of communication channel. Currently, only `EMAIL` is supported.
type PublicStatusRequestChannel string

const (
	PublicStatusRequestChannelEmail PublicStatusRequestChannel = "EMAIL"
)

// The status of the contact's subscription.
type PublicStatusRequestStatusState string

const (
	PublicStatusRequestStatusStateSubscribed   PublicStatusRequestStatusState = "SUBSCRIBED"
	PublicStatusRequestStatusStateUnsubscribed PublicStatusRequestStatusState = "UNSUBSCRIBED"
	PublicStatusRequestStatusStateNotSpecified PublicStatusRequestStatusState = "NOT_SPECIFIED"
)

// The legal basis for communication.
type PublicStatusRequestLegalBasis string

const (
	PublicStatusRequestLegalBasisLegitimateInterestPql    PublicStatusRequestLegalBasis = "LEGITIMATE_INTEREST_PQL"
	PublicStatusRequestLegalBasisLegitimateInterestClient PublicStatusRequestLegalBasis = "LEGITIMATE_INTEREST_CLIENT"
	PublicStatusRequestLegalBasisPerformanceOfContract    PublicStatusRequestLegalBasis = "PERFORMANCE_OF_CONTRACT"
	PublicStatusRequestLegalBasisConsentWithNotice        PublicStatusRequestLegalBasis = "CONSENT_WITH_NOTICE"
	PublicStatusRequestLegalBasisNonGdpr                  PublicStatusRequestLegalBasis = "NON_GDPR"
	PublicStatusRequestLegalBasisProcessAndStore          PublicStatusRequestLegalBasis = "PROCESS_AND_STORE"
	PublicStatusRequestLegalBasisLegitimateInterestOther  PublicStatusRequestLegalBasis = "LEGITIMATE_INTEREST_OTHER"
)

type PublicWideStatus struct {
	// The type of communication channel, with 'EMAIL' as the only supported option.
	//
	// Any of "EMAIL".
	Channel PublicWideStatusChannel `json:"channel,required"`
	// The subscription status of the contact, which can be 'SUBSCRIBED',
	// 'UNSUBSCRIBED', or 'NOT_SPECIFIED'.
	//
	// Any of "SUBSCRIBED", "UNSUBSCRIBED", "NOT_SPECIFIED".
	Status PublicWideStatusStatus `json:"status,required"`
	// The email address of the contact.
	SubscriberIDString string `json:"subscriberIdString,required"`
	// The date and time when the status was recorded.
	Timestamp time.Time `json:"timestamp,required" format:"date-time"`
	// The type of wide status, which can be 'PORTAL_WIDE' or 'BUSINESS_UNIT_WIDE'.
	//
	// Any of "PORTAL_WIDE", "BUSINESS_UNIT_WIDE".
	WideStatusType PublicWideStatusWideStatusType `json:"wideStatusType,required"`
	// The ID of the business unit associated with the status.
	BusinessUnitID int64 `json:"businessUnitId"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Channel            respjson.Field
		Status             respjson.Field
		SubscriberIDString respjson.Field
		Timestamp          respjson.Field
		WideStatusType     respjson.Field
		BusinessUnitID     respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PublicWideStatus) RawJSON() string { return r.JSON.raw }
func (r *PublicWideStatus) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The type of communication channel, with 'EMAIL' as the only supported option.
type PublicWideStatusChannel string

const (
	PublicWideStatusChannelEmail PublicWideStatusChannel = "EMAIL"
)

// The subscription status of the contact, which can be 'SUBSCRIBED',
// 'UNSUBSCRIBED', or 'NOT_SPECIFIED'.
type PublicWideStatusStatus string

const (
	PublicWideStatusStatusSubscribed   PublicWideStatusStatus = "SUBSCRIBED"
	PublicWideStatusStatusUnsubscribed PublicWideStatusStatus = "UNSUBSCRIBED"
	PublicWideStatusStatusNotSpecified PublicWideStatusStatus = "NOT_SPECIFIED"
)

// The type of wide status, which can be 'PORTAL_WIDE' or 'BUSINESS_UNIT_WIDE'.
type PublicWideStatusWideStatusType string

const (
	PublicWideStatusWideStatusTypePortalWide       PublicWideStatusWideStatusType = "PORTAL_WIDE"
	PublicWideStatusWideStatusTypeBusinessUnitWide PublicWideStatusWideStatusType = "BUSINESS_UNIT_WIDE"
)

type PublicWideStatusBulkResponse struct {
	// The contact's email address.
	SubscriberIDString string `json:"subscriberIdString,required"`
	// An array containing the wide status results for the operation.
	WideStatuses []PublicWideStatus `json:"wideStatuses,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		SubscriberIDString respjson.Field
		WideStatuses       respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PublicWideStatusBulkResponse) RawJSON() string { return r.JSON.raw }
func (r *PublicWideStatusBulkResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
