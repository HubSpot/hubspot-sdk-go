// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package communication_preferences

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
	"github.com/HubSpot/hubspot-sdk-go/packages/param"
	"github.com/HubSpot/hubspot-sdk-go/packages/respjson"
	"github.com/HubSpot/hubspot-sdk-go/shared"
)

// CommunicationPreferenceService contains methods and other services that help
// with interacting with the hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCommunicationPreferenceService] method instead.
type CommunicationPreferenceService struct {
	options     []option.RequestOption
	Definitions DefinitionService
	Statuses    StatusService
}

// NewCommunicationPreferenceService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewCommunicationPreferenceService(opts ...option.RequestOption) (r CommunicationPreferenceService) {
	r = CommunicationPreferenceService{}
	r.options = opts
	r.Definitions = NewDefinitionService(opts...)
	r.Statuses = NewStatusService(opts...)
	return
}

// Generate communication preference links for a subscriber. This endpoint allows
// you to create URLs for managing preferences and unsubscribing, tailored to a
// specific subscriber. It is useful for integrating communication preference
// management into your applications.
func (r *CommunicationPreferenceService) GenerateLinks(ctx context.Context, params CommunicationPreferenceGenerateLinksParams, opts ...option.RequestOption) (res *LinkGenerationResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "communication-preferences/2026-03/links/generate"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Retrieve a contact's current email subscription preferences.
func (r *CommunicationPreferenceService) GetStatuses(ctx context.Context, subscriberIDString string, query CommunicationPreferenceGetStatusesParams, opts ...option.RequestOption) (res *ActionResponseWithResultsPublicStatus, err error) {
	opts = slices.Concat(r.options, opts)
	if subscriberIDString == "" {
		err = errors.New("missing required subscriberIdString parameter")
		return nil, err
	}
	path := fmt.Sprintf("communication-preferences/2026-03/statuses/%s", url.PathEscape(subscriberIDString))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Check whether a contact has unsubscribed from all email subscriptions. If a
// contact has not opted out of all communications, the response `results` array
// will be empty.
func (r *CommunicationPreferenceService) GetUnsubscribeAllStatus(ctx context.Context, subscriberIDString string, query CommunicationPreferenceGetUnsubscribeAllStatusParams, opts ...option.RequestOption) (res *ActionResponseWithResultsPublicWideStatus, err error) {
	opts = slices.Concat(r.options, opts)
	if subscriberIDString == "" {
		err = errors.New("missing required subscriberIdString parameter")
		return nil, err
	}
	path := fmt.Sprintf("communication-preferences/2026-03/statuses/%s/unsubscribe-all", url.PathEscape(subscriberIDString))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Unsubscribe a contact from all email subscriptions.
func (r *CommunicationPreferenceService) UnsubscribeAll(ctx context.Context, subscriberIDString string, body CommunicationPreferenceUnsubscribeAllParams, opts ...option.RequestOption) (res *ActionResponseWithResultsPublicStatus, err error) {
	opts = slices.Concat(r.options, opts)
	if subscriberIDString == "" {
		err = errors.New("missing required subscriberIdString parameter")
		return nil, err
	}
	path := fmt.Sprintf("communication-preferences/2026-03/statuses/%s/unsubscribe-all", url.PathEscape(subscriberIDString))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Set the subscription status of a specific contact.
func (r *CommunicationPreferenceService) UpdateStatus(ctx context.Context, subscriberIDString string, body CommunicationPreferenceUpdateStatusParams, opts ...option.RequestOption) (res *ActionResponseWithResultsPublicStatus, err error) {
	opts = slices.Concat(r.options, opts)
	if subscriberIDString == "" {
		err = errors.New("missing required subscriberIdString parameter")
		return nil, err
	}
	path := fmt.Sprintf("communication-preferences/2026-03/statuses/%s", url.PathEscape(subscriberIDString))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type ActionResponseWithResultsPublicStatus struct {
	// The date and time when the operation was completed.
	CompletedAt time.Time `json:"completedAt" api:"required" format:"date-time"`
	// An array of results from the operation.
	Results []PublicStatus `json:"results" api:"required"`
	// The date and time when the operation started.
	StartedAt time.Time `json:"startedAt" api:"required" format:"date-time"`
	// Indicates the current status of the operation, with possible values: PENDING,
	// PROCESSING, CANCELED, COMPLETE.
	//
	// Any of "CANCELED", "COMPLETE", "PENDING", "PROCESSING".
	Status ActionResponseWithResultsPublicStatusStatus `json:"status" api:"required"`
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
	ActionResponseWithResultsPublicStatusStatusCanceled   ActionResponseWithResultsPublicStatusStatus = "CANCELED"
	ActionResponseWithResultsPublicStatusStatusComplete   ActionResponseWithResultsPublicStatusStatus = "COMPLETE"
	ActionResponseWithResultsPublicStatusStatusPending    ActionResponseWithResultsPublicStatusStatus = "PENDING"
	ActionResponseWithResultsPublicStatusStatusProcessing ActionResponseWithResultsPublicStatusStatus = "PROCESSING"
)

type ActionResponseWithResultsPublicWideStatus struct {
	// The date and time when the operation was completed.
	CompletedAt time.Time `json:"completedAt" api:"required" format:"date-time"`
	// An array containing the results of the operation.
	Results []PublicWideStatus `json:"results" api:"required"`
	// The date and time when the operation started.
	StartedAt time.Time `json:"startedAt" api:"required" format:"date-time"`
	// The current status of the operation, which can be PENDING, PROCESSING, CANCELED,
	// or COMPLETE.
	//
	// Any of "CANCELED", "COMPLETE", "PENDING", "PROCESSING".
	Status ActionResponseWithResultsPublicWideStatusStatus `json:"status" api:"required"`
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
	ActionResponseWithResultsPublicWideStatusStatusCanceled   ActionResponseWithResultsPublicWideStatusStatus = "CANCELED"
	ActionResponseWithResultsPublicWideStatusStatusComplete   ActionResponseWithResultsPublicWideStatusStatus = "COMPLETE"
	ActionResponseWithResultsPublicWideStatusStatusPending    ActionResponseWithResultsPublicWideStatusStatus = "PENDING"
	ActionResponseWithResultsPublicWideStatusStatusProcessing ActionResponseWithResultsPublicWideStatusStatus = "PROCESSING"
)

type ActionResponseWithResultsSubscriptionDefinition struct {
	// The date and time when the operation was completed.
	CompletedAt time.Time `json:"completedAt" api:"required" format:"date-time"`
	// An array containing the results of the operation.
	Results []SubscriptionDefinition `json:"results" api:"required"`
	// The date and time when the operation started.
	StartedAt time.Time `json:"startedAt" api:"required" format:"date-time"`
	// The current status of the operation, which can be PENDING, PROCESSING, CANCELED,
	// or COMPLETE.
	//
	// Any of "CANCELED", "COMPLETE", "PENDING", "PROCESSING".
	Status ActionResponseWithResultsSubscriptionDefinitionStatus `json:"status" api:"required"`
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
	ActionResponseWithResultsSubscriptionDefinitionStatusCanceled   ActionResponseWithResultsSubscriptionDefinitionStatus = "CANCELED"
	ActionResponseWithResultsSubscriptionDefinitionStatusComplete   ActionResponseWithResultsSubscriptionDefinitionStatus = "COMPLETE"
	ActionResponseWithResultsSubscriptionDefinitionStatusPending    ActionResponseWithResultsSubscriptionDefinitionStatus = "PENDING"
	ActionResponseWithResultsSubscriptionDefinitionStatusProcessing ActionResponseWithResultsSubscriptionDefinitionStatus = "PROCESSING"
)

// The property Inputs is required.
type BatchInputPublicStatusRequestParam struct {
	// An array of PublicStatusRequest objects, each representing a subscription status
	// update request. This property is required.
	Inputs []PublicStatusRequestParam `json:"inputs,omitzero" api:"required"`
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
	CompletedAt time.Time `json:"completedAt" api:"required" format:"date-time"`
	// An array containing the results of the bulk opt-out from all communications
	// operation.
	Results []PublicBulkOptOutFromAllResponse `json:"results" api:"required"`
	// The date and time when the bulk opt-out operation began.
	StartedAt time.Time `json:"startedAt" api:"required" format:"date-time"`
	// The current status of the bulk opt-out operation, which can be PENDING,
	// PROCESSING, CANCELED, or COMPLETE.
	//
	// Any of "CANCELED", "COMPLETE", "PENDING", "PROCESSING".
	Status BatchResponsePublicBulkOptOutFromAllResponseStatus `json:"status" api:"required"`
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
	BatchResponsePublicBulkOptOutFromAllResponseStatusCanceled   BatchResponsePublicBulkOptOutFromAllResponseStatus = "CANCELED"
	BatchResponsePublicBulkOptOutFromAllResponseStatusComplete   BatchResponsePublicBulkOptOutFromAllResponseStatus = "COMPLETE"
	BatchResponsePublicBulkOptOutFromAllResponseStatusPending    BatchResponsePublicBulkOptOutFromAllResponseStatus = "PENDING"
	BatchResponsePublicBulkOptOutFromAllResponseStatusProcessing BatchResponsePublicBulkOptOutFromAllResponseStatus = "PROCESSING"
)

type BatchResponsePublicStatus struct {
	// The date and time when the batch operation was completed.
	CompletedAt time.Time `json:"completedAt" api:"required" format:"date-time"`
	// An array containing the results of the batch operation.
	Results []PublicStatus `json:"results" api:"required"`
	// The date and time when the batch operation started.
	StartedAt time.Time `json:"startedAt" api:"required" format:"date-time"`
	// The current status of the batch operation, which can be PENDING, PROCESSING,
	// CANCELED, or COMPLETE.
	//
	// Any of "CANCELED", "COMPLETE", "PENDING", "PROCESSING".
	Status BatchResponsePublicStatusStatus `json:"status" api:"required"`
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
	BatchResponsePublicStatusStatusCanceled   BatchResponsePublicStatusStatus = "CANCELED"
	BatchResponsePublicStatusStatusComplete   BatchResponsePublicStatusStatus = "COMPLETE"
	BatchResponsePublicStatusStatusPending    BatchResponsePublicStatusStatus = "PENDING"
	BatchResponsePublicStatusStatusProcessing BatchResponsePublicStatusStatus = "PROCESSING"
)

type BatchResponsePublicStatusBulkResponse struct {
	// The date and time when the batch process was completed.
	CompletedAt time.Time `json:"completedAt" api:"required" format:"date-time"`
	// The array of results from the batch process, each containing subscription status
	// information.
	Results []PublicStatusBulkResponse `json:"results" api:"required"`
	// The date and time when the batch process began.
	StartedAt time.Time `json:"startedAt" api:"required" format:"date-time"`
	// The current status of the batch process, with possible values: PENDING,
	// PROCESSING, CANCELED, COMPLETE.
	//
	// Any of "CANCELED", "COMPLETE", "PENDING", "PROCESSING".
	Status BatchResponsePublicStatusBulkResponseStatus `json:"status" api:"required"`
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
	BatchResponsePublicStatusBulkResponseStatusCanceled   BatchResponsePublicStatusBulkResponseStatus = "CANCELED"
	BatchResponsePublicStatusBulkResponseStatusComplete   BatchResponsePublicStatusBulkResponseStatus = "COMPLETE"
	BatchResponsePublicStatusBulkResponseStatusPending    BatchResponsePublicStatusBulkResponseStatus = "PENDING"
	BatchResponsePublicStatusBulkResponseStatusProcessing BatchResponsePublicStatusBulkResponseStatus = "PROCESSING"
)

type BatchResponsePublicWideStatusBulkResponse struct {
	// The date and time when the batch process was completed.
	CompletedAt time.Time `json:"completedAt" api:"required" format:"date-time"`
	// The array of results from the batch process, each containing subscription status
	// information.
	Results []PublicWideStatusBulkResponse `json:"results" api:"required"`
	// The date and time when the batch process began.
	StartedAt time.Time `json:"startedAt" api:"required" format:"date-time"`
	// The current status of the batch process, with possible values: PENDING,
	// PROCESSING, CANCELED, COMPLETE.
	//
	// Any of "CANCELED", "COMPLETE", "PENDING", "PROCESSING".
	Status BatchResponsePublicWideStatusBulkResponseStatus `json:"status" api:"required"`
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
	BatchResponsePublicWideStatusBulkResponseStatusCanceled   BatchResponsePublicWideStatusBulkResponseStatus = "CANCELED"
	BatchResponsePublicWideStatusBulkResponseStatusComplete   BatchResponsePublicWideStatusBulkResponseStatus = "COMPLETE"
	BatchResponsePublicWideStatusBulkResponseStatusPending    BatchResponsePublicWideStatusBulkResponseStatus = "PENDING"
	BatchResponsePublicWideStatusBulkResponseStatusProcessing BatchResponsePublicWideStatusBulkResponseStatus = "PROCESSING"
)

// The property SubscriberIDString is required.
type LinkGenerationRequestParam struct {
	// A string representing the unique identifier of the subscriber. This property is
	// required.
	SubscriberIDString string `json:"subscriberIdString" api:"required"`
	// The language in which the generated link should be presented, represented as a
	// string.
	Language param.Opt[string] `json:"language,omitzero"`
	// The unique identifier for the subscription, represented as an integer in int64
	// format.
	SubscriptionID param.Opt[int64] `json:"subscriptionId,omitzero"`
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
	// The URL where the subscriber can manage their communication preferences.
	ManagePreferencesURL string `json:"managePreferencesUrl" api:"required"`
	// A string representing the unique identifier of the subscriber.
	SubscriberIDString string `json:"subscriberIdString" api:"required"`
	// A string containing the URL for unsubscribing the subscriber from all
	// communications.
	UnsubscribeAllURL string `json:"unsubscribeAllUrl" api:"required"`
	// A string containing the URL to unsubscribe the subscriber from a single
	// communication.
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
	Channel PartialPublicStatusRequestChannel `json:"channel,omitzero" api:"required"`
	// The current subscription status of the contact, which can be 'SUBSCRIBED',
	// 'UNSUBSCRIBED', or 'NOT_SPECIFIED'.
	//
	// Any of "NOT_SPECIFIED", "SUBSCRIBED", "UNSUBSCRIBED".
	StatusState PartialPublicStatusRequestStatusState `json:"statusState,omitzero" api:"required"`
	// The unique identifier of the subscription to be updated.
	SubscriptionID int64 `json:"subscriptionId" api:"required"`
	// An explanation for the legal basis used for communication.
	LegalBasisExplanation param.Opt[string] `json:"legalBasisExplanation,omitzero"`
	// The legal basis for communication, with options including
	// 'LEGITIMATE_INTEREST_PQL', 'LEGITIMATE_INTEREST_CLIENT',
	// 'PERFORMANCE_OF_CONTRACT', 'CONSENT_WITH_NOTICE', 'NON_GDPR',
	// 'PROCESS_AND_STORE', and 'LEGITIMATE_INTEREST_OTHER'.
	//
	// Any of "CONSENT_WITH_NOTICE", "LEGITIMATE_INTEREST_CLIENT",
	// "LEGITIMATE_INTEREST_OTHER", "LEGITIMATE_INTEREST_PQL", "NON_GDPR",
	// "PERFORMANCE_OF_CONTRACT", "PROCESS_AND_STORE".
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
	PartialPublicStatusRequestStatusStateNotSpecified PartialPublicStatusRequestStatusState = "NOT_SPECIFIED"
	PartialPublicStatusRequestStatusStateSubscribed   PartialPublicStatusRequestStatusState = "SUBSCRIBED"
	PartialPublicStatusRequestStatusStateUnsubscribed PartialPublicStatusRequestStatusState = "UNSUBSCRIBED"
)

// The legal basis for communication, with options including
// 'LEGITIMATE_INTEREST_PQL', 'LEGITIMATE_INTEREST_CLIENT',
// 'PERFORMANCE_OF_CONTRACT', 'CONSENT_WITH_NOTICE', 'NON_GDPR',
// 'PROCESS_AND_STORE', and 'LEGITIMATE_INTEREST_OTHER'.
type PartialPublicStatusRequestLegalBasis string

const (
	PartialPublicStatusRequestLegalBasisConsentWithNotice        PartialPublicStatusRequestLegalBasis = "CONSENT_WITH_NOTICE"
	PartialPublicStatusRequestLegalBasisLegitimateInterestClient PartialPublicStatusRequestLegalBasis = "LEGITIMATE_INTEREST_CLIENT"
	PartialPublicStatusRequestLegalBasisLegitimateInterestOther  PartialPublicStatusRequestLegalBasis = "LEGITIMATE_INTEREST_OTHER"
	PartialPublicStatusRequestLegalBasisLegitimateInterestPql    PartialPublicStatusRequestLegalBasis = "LEGITIMATE_INTEREST_PQL"
	PartialPublicStatusRequestLegalBasisNonGdpr                  PartialPublicStatusRequestLegalBasis = "NON_GDPR"
	PartialPublicStatusRequestLegalBasisPerformanceOfContract    PartialPublicStatusRequestLegalBasis = "PERFORMANCE_OF_CONTRACT"
	PartialPublicStatusRequestLegalBasisProcessAndStore          PartialPublicStatusRequestLegalBasis = "PROCESS_AND_STORE"
)

type PublicBulkOptOutFromAllResponse struct {
	// The email address of the contact.
	SubscriberIDString string `json:"subscriberIdString" api:"required"`
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
	Channel PublicStatusChannel `json:"channel" api:"required"`
	// The origin or method through which the subscription status was set.
	Source string `json:"source" api:"required"`
	// The current subscription status of the contact, which can be 'SUBSCRIBED',
	// 'UNSUBSCRIBED', or 'NOT_SPECIFIED'.
	//
	// Any of "NOT_SPECIFIED", "SUBSCRIBED", "UNSUBSCRIBED".
	Status PublicStatusStatus `json:"status" api:"required"`
	// The contact's email address.
	SubscriberIDString string `json:"subscriberIdString" api:"required"`
	// The unique identifier of the subscription.
	SubscriptionID int64 `json:"subscriptionId" api:"required"`
	// The date and time when the subscription status was last updated.
	Timestamp time.Time `json:"timestamp" api:"required" format:"date-time"`
	// The ID of the business unit associated with the subscription.
	BusinessUnitID int64 `json:"businessUnitId"`
	// The legal basis for communication, with options including
	// 'LEGITIMATE_INTEREST_PQL', 'LEGITIMATE_INTEREST_CLIENT',
	// 'PERFORMANCE_OF_CONTRACT', 'CONSENT_WITH_NOTICE', 'NON_GDPR',
	// 'PROCESS_AND_STORE', and 'LEGITIMATE_INTEREST_OTHER'.
	//
	// Any of "CONSENT_WITH_NOTICE", "LEGITIMATE_INTEREST_CLIENT",
	// "LEGITIMATE_INTEREST_OTHER", "LEGITIMATE_INTEREST_PQL", "NON_GDPR",
	// "PERFORMANCE_OF_CONTRACT", "PROCESS_AND_STORE".
	LegalBasis PublicStatusLegalBasis `json:"legalBasis"`
	// An explanation for the legal basis used for communication.
	LegalBasisExplanation string `json:"legalBasisExplanation"`
	// The reason for the successful change in subscription status, such as
	// 'RESUBSCRIBE_OCCURRED' or 'NO_STATUS_CHANGE'.
	//
	// Any of "NO_STATUS_CHANGE", "REQUESTED_CHANGE_OCCURRED", "RESUBSCRIBE_OCCURRED",
	// "UNSUBSCRIBE_FROM_ALL_OCCURRED".
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
	PublicStatusStatusNotSpecified PublicStatusStatus = "NOT_SPECIFIED"
	PublicStatusStatusSubscribed   PublicStatusStatus = "SUBSCRIBED"
	PublicStatusStatusUnsubscribed PublicStatusStatus = "UNSUBSCRIBED"
)

// The legal basis for communication, with options including
// 'LEGITIMATE_INTEREST_PQL', 'LEGITIMATE_INTEREST_CLIENT',
// 'PERFORMANCE_OF_CONTRACT', 'CONSENT_WITH_NOTICE', 'NON_GDPR',
// 'PROCESS_AND_STORE', and 'LEGITIMATE_INTEREST_OTHER'.
type PublicStatusLegalBasis string

const (
	PublicStatusLegalBasisConsentWithNotice        PublicStatusLegalBasis = "CONSENT_WITH_NOTICE"
	PublicStatusLegalBasisLegitimateInterestClient PublicStatusLegalBasis = "LEGITIMATE_INTEREST_CLIENT"
	PublicStatusLegalBasisLegitimateInterestOther  PublicStatusLegalBasis = "LEGITIMATE_INTEREST_OTHER"
	PublicStatusLegalBasisLegitimateInterestPql    PublicStatusLegalBasis = "LEGITIMATE_INTEREST_PQL"
	PublicStatusLegalBasisNonGdpr                  PublicStatusLegalBasis = "NON_GDPR"
	PublicStatusLegalBasisPerformanceOfContract    PublicStatusLegalBasis = "PERFORMANCE_OF_CONTRACT"
	PublicStatusLegalBasisProcessAndStore          PublicStatusLegalBasis = "PROCESS_AND_STORE"
)

// The reason for the successful change in subscription status, such as
// 'RESUBSCRIBE_OCCURRED' or 'NO_STATUS_CHANGE'.
type PublicStatusSetStatusSuccessReason string

const (
	PublicStatusSetStatusSuccessReasonNoStatusChange             PublicStatusSetStatusSuccessReason = "NO_STATUS_CHANGE"
	PublicStatusSetStatusSuccessReasonRequestedChangeOccurred    PublicStatusSetStatusSuccessReason = "REQUESTED_CHANGE_OCCURRED"
	PublicStatusSetStatusSuccessReasonResubscribeOccurred        PublicStatusSetStatusSuccessReason = "RESUBSCRIBE_OCCURRED"
	PublicStatusSetStatusSuccessReasonUnsubscribeFromAllOccurred PublicStatusSetStatusSuccessReason = "UNSUBSCRIBE_FROM_ALL_OCCURRED"
)

type PublicStatusBulkResponse struct {
	// An array of subscription status objects for the contact.
	Statuses []PublicStatus `json:"statuses" api:"required"`
	// The email address of the contact.
	SubscriberIDString string `json:"subscriberIdString" api:"required"`
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
	Channel PublicStatusRequestChannel `json:"channel,omitzero" api:"required"`
	// The status of the contact's subscription.
	//
	// Any of "NOT_SPECIFIED", "SUBSCRIBED", "UNSUBSCRIBED".
	StatusState PublicStatusRequestStatusState `json:"statusState,omitzero" api:"required"`
	// The contact's email address.
	SubscriberIDString string `json:"subscriberIdString" api:"required"`
	// The ID of the subscription to update.
	SubscriptionID int64 `json:"subscriptionId" api:"required"`
	// The explanation for the legal basis.
	LegalBasisExplanation param.Opt[string] `json:"legalBasisExplanation,omitzero"`
	// The legal basis for communication.
	//
	// Any of "CONSENT_WITH_NOTICE", "LEGITIMATE_INTEREST_CLIENT",
	// "LEGITIMATE_INTEREST_OTHER", "LEGITIMATE_INTEREST_PQL", "NON_GDPR",
	// "PERFORMANCE_OF_CONTRACT", "PROCESS_AND_STORE".
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
	PublicStatusRequestStatusStateNotSpecified PublicStatusRequestStatusState = "NOT_SPECIFIED"
	PublicStatusRequestStatusStateSubscribed   PublicStatusRequestStatusState = "SUBSCRIBED"
	PublicStatusRequestStatusStateUnsubscribed PublicStatusRequestStatusState = "UNSUBSCRIBED"
)

// The legal basis for communication.
type PublicStatusRequestLegalBasis string

const (
	PublicStatusRequestLegalBasisConsentWithNotice        PublicStatusRequestLegalBasis = "CONSENT_WITH_NOTICE"
	PublicStatusRequestLegalBasisLegitimateInterestClient PublicStatusRequestLegalBasis = "LEGITIMATE_INTEREST_CLIENT"
	PublicStatusRequestLegalBasisLegitimateInterestOther  PublicStatusRequestLegalBasis = "LEGITIMATE_INTEREST_OTHER"
	PublicStatusRequestLegalBasisLegitimateInterestPql    PublicStatusRequestLegalBasis = "LEGITIMATE_INTEREST_PQL"
	PublicStatusRequestLegalBasisNonGdpr                  PublicStatusRequestLegalBasis = "NON_GDPR"
	PublicStatusRequestLegalBasisPerformanceOfContract    PublicStatusRequestLegalBasis = "PERFORMANCE_OF_CONTRACT"
	PublicStatusRequestLegalBasisProcessAndStore          PublicStatusRequestLegalBasis = "PROCESS_AND_STORE"
)

type PublicSubscriptionTranslation struct {
	// The timestamp indicating when the subscription translation was created.
	CreatedAt int64 `json:"createdAt" api:"required"`
	// A text description of the subscription translation.
	Description string `json:"description" api:"required"`
	// The code representing the language of the subscription translation.
	LanguageCode string `json:"languageCode" api:"required"`
	// The name of the subscription translation.
	Name string `json:"name" api:"required"`
	// The unique identifier for the subscription associated with the translation.
	SubscriptionID int64 `json:"subscriptionId" api:"required"`
	// The timestamp indicating when the subscription translation was last updated.
	UpdatedAt int64 `json:"updatedAt" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreatedAt      respjson.Field
		Description    respjson.Field
		LanguageCode   respjson.Field
		Name           respjson.Field
		SubscriptionID respjson.Field
		UpdatedAt      respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PublicSubscriptionTranslation) RawJSON() string { return r.JSON.raw }
func (r *PublicSubscriptionTranslation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PublicWideStatus struct {
	// The type of communication channel, with 'EMAIL' as the only supported option.
	//
	// Any of "EMAIL".
	Channel PublicWideStatusChannel `json:"channel" api:"required"`
	// The subscription status of the contact, which can be 'SUBSCRIBED',
	// 'UNSUBSCRIBED', or 'NOT_SPECIFIED'.
	//
	// Any of "NOT_SPECIFIED", "SUBSCRIBED", "UNSUBSCRIBED".
	Status PublicWideStatusStatus `json:"status" api:"required"`
	// The email address of the contact.
	SubscriberIDString string `json:"subscriberIdString" api:"required"`
	// The date and time when the status was recorded.
	Timestamp time.Time `json:"timestamp" api:"required" format:"date-time"`
	// The type of wide status, which can be 'PORTAL_WIDE' or 'BUSINESS_UNIT_WIDE'.
	//
	// Any of "BUSINESS_UNIT_WIDE", "PORTAL_WIDE".
	WideStatusType PublicWideStatusWideStatusType `json:"wideStatusType" api:"required"`
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
	PublicWideStatusStatusNotSpecified PublicWideStatusStatus = "NOT_SPECIFIED"
	PublicWideStatusStatusSubscribed   PublicWideStatusStatus = "SUBSCRIBED"
	PublicWideStatusStatusUnsubscribed PublicWideStatusStatus = "UNSUBSCRIBED"
)

// The type of wide status, which can be 'PORTAL_WIDE' or 'BUSINESS_UNIT_WIDE'.
type PublicWideStatusWideStatusType string

const (
	PublicWideStatusWideStatusTypeBusinessUnitWide PublicWideStatusWideStatusType = "BUSINESS_UNIT_WIDE"
	PublicWideStatusWideStatusTypePortalWide       PublicWideStatusWideStatusType = "PORTAL_WIDE"
)

type PublicWideStatusBulkResponse struct {
	// The contact's email address.
	SubscriberIDString string `json:"subscriberIdString" api:"required"`
	// An array containing the wide status results for the operation.
	WideStatuses []PublicWideStatus `json:"wideStatuses" api:"required"`
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

type SubscriptionDefinition struct {
	// The unique identifier for the subscription.
	ID string `json:"id" api:"required"`
	// The date and time when the subscription was created.
	CreatedAt time.Time `json:"createdAt" api:"required" format:"date-time"`
	// A description of the subscription.
	Description string `json:"description" api:"required"`
	// Indicates whether the subscription is active.
	IsActive bool `json:"isActive" api:"required"`
	// Indicates whether the subscription is the default option.
	IsDefault bool `json:"isDefault" api:"required"`
	// Indicates whether the subscription is internal.
	IsInternal bool `json:"isInternal" api:"required"`
	// The name of the subscription.
	Name string `json:"name" api:"required"`
	// The date and time when the subscription was last updated.
	UpdatedAt time.Time `json:"updatedAt" api:"required" format:"date-time"`
	// The ID of the business unit associated with the subscription.
	BusinessUnitID int64 `json:"businessUnitId"`
	// The method of communication for the subscription.
	CommunicationMethod string `json:"communicationMethod"`
	// The purpose of the subscription.
	Purpose string `json:"purpose"`
	// A list of translations associated with the subscription.
	SubscriptionTranslations []PublicSubscriptionTranslation `json:"subscriptionTranslations"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                       respjson.Field
		CreatedAt                respjson.Field
		Description              respjson.Field
		IsActive                 respjson.Field
		IsDefault                respjson.Field
		IsInternal               respjson.Field
		Name                     respjson.Field
		UpdatedAt                respjson.Field
		BusinessUnitID           respjson.Field
		CommunicationMethod      respjson.Field
		Purpose                  respjson.Field
		SubscriptionTranslations respjson.Field
		ExtraFields              map[string]respjson.Field
		raw                      string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SubscriptionDefinition) RawJSON() string { return r.JSON.raw }
func (r *SubscriptionDefinition) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CommunicationPreferenceGenerateLinksParams struct {
	// The communication channel for which the links are generated. Must be 'EMAIL'.
	//
	// Any of "EMAIL".
	Channel               CommunicationPreferenceGenerateLinksParamsChannel `query:"channel,omitzero" api:"required" json:"-"`
	LinkGenerationRequest LinkGenerationRequestParam
	// The identifier of the business unit. Defaults to 0 if not specified.
	BusinessUnitID param.Opt[int64] `query:"businessUnitId,omitzero" json:"-"`
	paramObj
}

func (r CommunicationPreferenceGenerateLinksParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.LinkGenerationRequest)
}
func (r *CommunicationPreferenceGenerateLinksParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [CommunicationPreferenceGenerateLinksParams]'s query
// parameters as `url.Values`.
func (r CommunicationPreferenceGenerateLinksParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// The communication channel for which the links are generated. Must be 'EMAIL'.
type CommunicationPreferenceGenerateLinksParamsChannel string

const (
	CommunicationPreferenceGenerateLinksParamsChannelEmail CommunicationPreferenceGenerateLinksParamsChannel = "EMAIL"
)

type CommunicationPreferenceGetStatusesParams struct {
	// The communication channel for which the subscription status is being retrieved.
	// This parameter is required and currently supports only 'EMAIL'.
	//
	// Any of "EMAIL".
	Channel CommunicationPreferenceGetStatusesParamsChannel `query:"channel,omitzero" api:"required" json:"-"`
	// The ID of the business unit to filter the subscription status by. This is an
	// optional parameter.
	BusinessUnitID param.Opt[int64] `query:"businessUnitId,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [CommunicationPreferenceGetStatusesParams]'s query
// parameters as `url.Values`.
func (r CommunicationPreferenceGetStatusesParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// The communication channel for which the subscription status is being retrieved.
// This parameter is required and currently supports only 'EMAIL'.
type CommunicationPreferenceGetStatusesParamsChannel string

const (
	CommunicationPreferenceGetStatusesParamsChannelEmail CommunicationPreferenceGetStatusesParamsChannel = "EMAIL"
)

type CommunicationPreferenceGetUnsubscribeAllStatusParams struct {
	// The communication channel from which to unsubscribe the subscriber. This is a
	// required parameter and must be 'EMAIL'.
	//
	// Any of "EMAIL".
	Channel CommunicationPreferenceGetUnsubscribeAllStatusParamsChannel `query:"channel,omitzero" api:"required" json:"-"`
	// The ID of the business unit to which the subscriber belongs. This is an optional
	// parameter.
	BusinessUnitID param.Opt[int64] `query:"businessUnitId,omitzero" json:"-"`
	// A boolean indicating whether to include detailed information in the response.
	// Defaults to false.
	Verbose param.Opt[bool] `query:"verbose,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [CommunicationPreferenceGetUnsubscribeAllStatusParams]'s
// query parameters as `url.Values`.
func (r CommunicationPreferenceGetUnsubscribeAllStatusParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// The communication channel from which to unsubscribe the subscriber. This is a
// required parameter and must be 'EMAIL'.
type CommunicationPreferenceGetUnsubscribeAllStatusParamsChannel string

const (
	CommunicationPreferenceGetUnsubscribeAllStatusParamsChannelEmail CommunicationPreferenceGetUnsubscribeAllStatusParamsChannel = "EMAIL"
)

type CommunicationPreferenceUnsubscribeAllParams struct {
	// The communication channel to unsubscribe from. Must be 'EMAIL'.
	//
	// Any of "EMAIL".
	Channel CommunicationPreferenceUnsubscribeAllParamsChannel `query:"channel,omitzero" api:"required" json:"-"`
	// The ID of the business unit associated with the request. This is an optional
	// integer parameter.
	BusinessUnitID param.Opt[int64] `query:"businessUnitId,omitzero" json:"-"`
	// A boolean indicating whether to include detailed information in the response.
	// Defaults to false.
	Verbose param.Opt[bool] `query:"verbose,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [CommunicationPreferenceUnsubscribeAllParams]'s query
// parameters as `url.Values`.
func (r CommunicationPreferenceUnsubscribeAllParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// The communication channel to unsubscribe from. Must be 'EMAIL'.
type CommunicationPreferenceUnsubscribeAllParamsChannel string

const (
	CommunicationPreferenceUnsubscribeAllParamsChannelEmail CommunicationPreferenceUnsubscribeAllParamsChannel = "EMAIL"
)

type CommunicationPreferenceUpdateStatusParams struct {
	PartialPublicStatusRequest PartialPublicStatusRequestParam
	paramObj
}

func (r CommunicationPreferenceUpdateStatusParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.PartialPublicStatusRequest)
}
func (r *CommunicationPreferenceUpdateStatusParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
