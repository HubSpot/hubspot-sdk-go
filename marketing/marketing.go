// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package marketing

import (
	"time"

	"github.com/stainless-sdks/hubspot-sdk-go/internal/apijson"
	"github.com/stainless-sdks/hubspot-sdk-go/option"
	"github.com/stainless-sdks/hubspot-sdk-go/packages/param"
	"github.com/stainless-sdks/hubspot-sdk-go/packages/respjson"
)

// MarketingService contains methods and other services that help with interacting
// with the Hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMarketingService] method instead.
type MarketingService struct {
	Options       []option.RequestOption
	Campaigns     CampaignService
	Emails        EmailService
	Events        EventService
	Forms         FormService
	SingleSend    SingleSendService
	Subscriptions SubscriptionService
	Transactional TransactionalService
}

// NewMarketingService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewMarketingService(opts ...option.RequestOption) (r MarketingService) {
	r = MarketingService{}
	r.Options = opts
	r.Campaigns = NewCampaignService(opts...)
	r.Emails = NewEmailService(opts...)
	r.Events = NewEventService(opts...)
	r.Forms = NewFormService(opts...)
	r.SingleSend = NewSingleSendService(opts...)
	r.Subscriptions = NewSubscriptionService(opts...)
	r.Transactional = NewTransactionalService(opts...)
	return
}

// Describes the status of an email send request.
type EmailSendStatusView struct {
	// Status of the send request.
	//
	// Any of "PENDING", "PROCESSING", "CANCELED", "COMPLETE".
	Status EmailSendStatusViewStatus `json:"status,required"`
	// Identifier used to query the status of the send.
	StatusID string `json:"statusId,required"`
	// Time when the send was completed.
	CompletedAt time.Time `json:"completedAt" format:"date-time"`
	// The ID of a send event.
	EventID EventIDView `json:"eventId"`
	Message string      `json:"message"`
	// Time when the send was requested.
	RequestedAt time.Time `json:"requestedAt" format:"date-time"`
	// Result of the send.
	//
	// Any of "SENT", "IDEMPOTENT_IGNORE", "QUEUED", "IDEMPOTENT_FAIL", "THROTTLED",
	// "EMAIL_DISABLED", "PORTAL_SUSPENDED", "INVALID_TO_ADDRESS", "BLOCKED_DOMAIN",
	// "PREVIOUSLY_BOUNCED", "EMAIL_UNCONFIRMED", "PREVIOUS_SPAM",
	// "PREVIOUSLY_UNSUBSCRIBED_MESSAGE", "PREVIOUSLY_UNSUBSCRIBED_PORTAL",
	// "INVALID_FROM_ADDRESS", "CAMPAIGN_CANCELLED", "VALIDATION_FAILED", "MTA_IGNORE",
	// "BLOCKED_ADDRESS", "PORTAL_OVER_LIMIT", "PORTAL_EXPIRED",
	// "PORTAL_MISSING_MARKETING_SCOPE", "MISSING_TEMPLATE_PROPERTIES",
	// "MISSING_REQUIRED_PARAMETER", "PORTAL_AUTHENTICATION_FAILURE",
	// "MISSING_CONTENT", "CORRUPT_INPUT", "TEMPLATE_RENDER_EXCEPTION",
	// "GRAYMAIL_SUPPRESSED", "UNCONFIGURED_SENDING_DOMAIN", "UNDELIVERABLE",
	// "CANCELLED_ABUSE", "QUARANTINED_ADDRESS", "ADDRESS_ONLY_ACCEPTED_ON_PROD",
	// "PORTAL_NOT_AUTHORIZED_FOR_APPLICATION", "ADDRESS_LIST_BOMBED",
	// "ADDRESS_OPTED_OUT", "RECIPIENT_FATIGUE_SUPPRESSED", "TOO_MANY_RECIPIENTS",
	// "PREVIOUSLY_UNSUBSCRIBED_BRAND", "NON_MARKETABLE_CONTACT",
	// "PREVIOUSLY_UNSUBSCRIBED_BUSINESS_UNIT", "GDPR_DOI_ENABLED",
	// "HUBL_LIMIT_EXCEEDED", "LOW_CONTACT_QUALITY_SCORE".
	SendResult EmailSendStatusViewSendResult `json:"sendResult"`
	// Time when the send began processing.
	StartedAt time.Time `json:"startedAt" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Status      respjson.Field
		StatusID    respjson.Field
		CompletedAt respjson.Field
		EventID     respjson.Field
		Message     respjson.Field
		RequestedAt respjson.Field
		SendResult  respjson.Field
		StartedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmailSendStatusView) RawJSON() string { return r.JSON.raw }
func (r *EmailSendStatusView) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Status of the send request.
type EmailSendStatusViewStatus string

const (
	EmailSendStatusViewStatusPending    EmailSendStatusViewStatus = "PENDING"
	EmailSendStatusViewStatusProcessing EmailSendStatusViewStatus = "PROCESSING"
	EmailSendStatusViewStatusCanceled   EmailSendStatusViewStatus = "CANCELED"
	EmailSendStatusViewStatusComplete   EmailSendStatusViewStatus = "COMPLETE"
)

// Result of the send.
type EmailSendStatusViewSendResult string

const (
	EmailSendStatusViewSendResultSent                               EmailSendStatusViewSendResult = "SENT"
	EmailSendStatusViewSendResultIdempotentIgnore                   EmailSendStatusViewSendResult = "IDEMPOTENT_IGNORE"
	EmailSendStatusViewSendResultQueued                             EmailSendStatusViewSendResult = "QUEUED"
	EmailSendStatusViewSendResultIdempotentFail                     EmailSendStatusViewSendResult = "IDEMPOTENT_FAIL"
	EmailSendStatusViewSendResultThrottled                          EmailSendStatusViewSendResult = "THROTTLED"
	EmailSendStatusViewSendResultEmailDisabled                      EmailSendStatusViewSendResult = "EMAIL_DISABLED"
	EmailSendStatusViewSendResultPortalSuspended                    EmailSendStatusViewSendResult = "PORTAL_SUSPENDED"
	EmailSendStatusViewSendResultInvalidToAddress                   EmailSendStatusViewSendResult = "INVALID_TO_ADDRESS"
	EmailSendStatusViewSendResultBlockedDomain                      EmailSendStatusViewSendResult = "BLOCKED_DOMAIN"
	EmailSendStatusViewSendResultPreviouslyBounced                  EmailSendStatusViewSendResult = "PREVIOUSLY_BOUNCED"
	EmailSendStatusViewSendResultEmailUnconfirmed                   EmailSendStatusViewSendResult = "EMAIL_UNCONFIRMED"
	EmailSendStatusViewSendResultPreviousSpam                       EmailSendStatusViewSendResult = "PREVIOUS_SPAM"
	EmailSendStatusViewSendResultPreviouslyUnsubscribedMessage      EmailSendStatusViewSendResult = "PREVIOUSLY_UNSUBSCRIBED_MESSAGE"
	EmailSendStatusViewSendResultPreviouslyUnsubscribedPortal       EmailSendStatusViewSendResult = "PREVIOUSLY_UNSUBSCRIBED_PORTAL"
	EmailSendStatusViewSendResultInvalidFromAddress                 EmailSendStatusViewSendResult = "INVALID_FROM_ADDRESS"
	EmailSendStatusViewSendResultCampaignCancelled                  EmailSendStatusViewSendResult = "CAMPAIGN_CANCELLED"
	EmailSendStatusViewSendResultValidationFailed                   EmailSendStatusViewSendResult = "VALIDATION_FAILED"
	EmailSendStatusViewSendResultMtaIgnore                          EmailSendStatusViewSendResult = "MTA_IGNORE"
	EmailSendStatusViewSendResultBlockedAddress                     EmailSendStatusViewSendResult = "BLOCKED_ADDRESS"
	EmailSendStatusViewSendResultPortalOverLimit                    EmailSendStatusViewSendResult = "PORTAL_OVER_LIMIT"
	EmailSendStatusViewSendResultPortalExpired                      EmailSendStatusViewSendResult = "PORTAL_EXPIRED"
	EmailSendStatusViewSendResultPortalMissingMarketingScope        EmailSendStatusViewSendResult = "PORTAL_MISSING_MARKETING_SCOPE"
	EmailSendStatusViewSendResultMissingTemplateProperties          EmailSendStatusViewSendResult = "MISSING_TEMPLATE_PROPERTIES"
	EmailSendStatusViewSendResultMissingRequiredParameter           EmailSendStatusViewSendResult = "MISSING_REQUIRED_PARAMETER"
	EmailSendStatusViewSendResultPortalAuthenticationFailure        EmailSendStatusViewSendResult = "PORTAL_AUTHENTICATION_FAILURE"
	EmailSendStatusViewSendResultMissingContent                     EmailSendStatusViewSendResult = "MISSING_CONTENT"
	EmailSendStatusViewSendResultCorruptInput                       EmailSendStatusViewSendResult = "CORRUPT_INPUT"
	EmailSendStatusViewSendResultTemplateRenderException            EmailSendStatusViewSendResult = "TEMPLATE_RENDER_EXCEPTION"
	EmailSendStatusViewSendResultGraymailSuppressed                 EmailSendStatusViewSendResult = "GRAYMAIL_SUPPRESSED"
	EmailSendStatusViewSendResultUnconfiguredSendingDomain          EmailSendStatusViewSendResult = "UNCONFIGURED_SENDING_DOMAIN"
	EmailSendStatusViewSendResultUndeliverable                      EmailSendStatusViewSendResult = "UNDELIVERABLE"
	EmailSendStatusViewSendResultCancelledAbuse                     EmailSendStatusViewSendResult = "CANCELLED_ABUSE"
	EmailSendStatusViewSendResultQuarantinedAddress                 EmailSendStatusViewSendResult = "QUARANTINED_ADDRESS"
	EmailSendStatusViewSendResultAddressOnlyAcceptedOnProd          EmailSendStatusViewSendResult = "ADDRESS_ONLY_ACCEPTED_ON_PROD"
	EmailSendStatusViewSendResultPortalNotAuthorizedForApplication  EmailSendStatusViewSendResult = "PORTAL_NOT_AUTHORIZED_FOR_APPLICATION"
	EmailSendStatusViewSendResultAddressListBombed                  EmailSendStatusViewSendResult = "ADDRESS_LIST_BOMBED"
	EmailSendStatusViewSendResultAddressOptedOut                    EmailSendStatusViewSendResult = "ADDRESS_OPTED_OUT"
	EmailSendStatusViewSendResultRecipientFatigueSuppressed         EmailSendStatusViewSendResult = "RECIPIENT_FATIGUE_SUPPRESSED"
	EmailSendStatusViewSendResultTooManyRecipients                  EmailSendStatusViewSendResult = "TOO_MANY_RECIPIENTS"
	EmailSendStatusViewSendResultPreviouslyUnsubscribedBrand        EmailSendStatusViewSendResult = "PREVIOUSLY_UNSUBSCRIBED_BRAND"
	EmailSendStatusViewSendResultNonMarketableContact               EmailSendStatusViewSendResult = "NON_MARKETABLE_CONTACT"
	EmailSendStatusViewSendResultPreviouslyUnsubscribedBusinessUnit EmailSendStatusViewSendResult = "PREVIOUSLY_UNSUBSCRIBED_BUSINESS_UNIT"
	EmailSendStatusViewSendResultGdprDoiEnabled                     EmailSendStatusViewSendResult = "GDPR_DOI_ENABLED"
	EmailSendStatusViewSendResultHublLimitExceeded                  EmailSendStatusViewSendResult = "HUBL_LIMIT_EXCEEDED"
	EmailSendStatusViewSendResultLowContactQualityScore             EmailSendStatusViewSendResult = "LOW_CONTACT_QUALITY_SCORE"
)

// The ID of a send event.
type EventIDView struct {
	// Identifier of event.
	ID string `json:"id,required" format:"uuid"`
	// Time of event creation.
	Created time.Time `json:"created,required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Created     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EventIDView) RawJSON() string { return r.JSON.raw }
func (r *EventIDView) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A JSON object containing anything you want to override.
//
// The property To is required.
type PublicSingleSendEmailParam struct {
	// The recipient of the email.
	To string `json:"to,required"`
	// The From header for the email.
	From param.Opt[string] `json:"from,omitzero"`
	// ID for a particular send. No more than one email will be sent per sendId.
	SendID param.Opt[string] `json:"sendId,omitzero"`
	// List of email addresses to send as Bcc.
	Bcc []string `json:"bcc,omitzero"`
	// List of email addresses to send as Cc.
	Cc []string `json:"cc,omitzero"`
	// List of Reply-To header values for the email.
	ReplyTo []string `json:"replyTo,omitzero"`
	paramObj
}

func (r PublicSingleSendEmailParam) MarshalJSON() (data []byte, err error) {
	type shadow PublicSingleSendEmailParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PublicSingleSendEmailParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A request to send a single email asynchronously.
//
// The properties EmailID, Message are required.
type PublicSingleSendRequestEggParam struct {
	// The content ID for the email, which can be found in email tool UI.
	EmailID int64 `json:"emailId,required"`
	// A JSON object containing anything you want to override.
	Message PublicSingleSendEmailParam `json:"message,omitzero,required"`
	// The contactProperties field is a map of contact property values. Each contact
	// property value contains a name and value property. Each property will get set on
	// the contact record and will be visible in the template under {{ contact.NAME }}.
	// Use these properties when you want to set a contact property while you’re
	// sending the email. For example, when sending a receipt you may want to set a
	// last_paid_date property, as the sending of the receipt will have information
	// about the last payment.
	ContactProperties map[string]string `json:"contactProperties,omitzero"`
	// The customProperties field is a map of property values. Each property value
	// contains a name and value property. Each property will be visible in the
	// template under {{ custom.NAME }}. Note: Custom properties do not currently
	// support arrays. To provide a listing in an email, one workaround is to build an
	// HTML list (either with tables or ul) and specify it as a custom property.
	CustomProperties map[string]any `json:"customProperties,omitzero"`
	paramObj
}

func (r PublicSingleSendRequestEggParam) MarshalJSON() (data []byte, err error) {
	type shadow PublicSingleSendRequestEggParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PublicSingleSendRequestEggParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
