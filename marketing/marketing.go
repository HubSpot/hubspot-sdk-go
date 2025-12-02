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
// with the hubspot API.
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
	// Any of "CANCELED", "COMPLETE", "PENDING", "PROCESSING".
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
	// Any of "ADDRESS_LIST_BOMBED", "ADDRESS_ONLY_ACCEPTED_ON_PROD",
	// "ADDRESS_OPTED_OUT", "BLOCKED_ADDRESS", "BLOCKED_DOMAIN", "CAMPAIGN_CANCELLED",
	// "CANCELLED_ABUSE", "CORRUPT_INPUT", "EMAIL_DISABLED", "EMAIL_UNCONFIRMED",
	// "GDPR_DOI_ENABLED", "GRAYMAIL_SUPPRESSED", "HUBL_LIMIT_EXCEEDED",
	// "IDEMPOTENT_FAIL", "IDEMPOTENT_IGNORE", "INVALID_FROM_ADDRESS",
	// "INVALID_TO_ADDRESS", "LOW_CONTACT_QUALITY_SCORE", "MISSING_CONTENT",
	// "MISSING_REQUIRED_PARAMETER", "MISSING_TEMPLATE_PROPERTIES", "MTA_IGNORE",
	// "NON_MARKETABLE_CONTACT", "PORTAL_AUTHENTICATION_FAILURE", "PORTAL_EXPIRED",
	// "PORTAL_MISSING_MARKETING_SCOPE", "PORTAL_NOT_AUTHORIZED_FOR_APPLICATION",
	// "PORTAL_OVER_LIMIT", "PORTAL_SUSPENDED", "PREVIOUS_SPAM", "PREVIOUSLY_BOUNCED",
	// "PREVIOUSLY_UNSUBSCRIBED_BRAND", "PREVIOUSLY_UNSUBSCRIBED_BUSINESS_UNIT",
	// "PREVIOUSLY_UNSUBSCRIBED_MESSAGE", "PREVIOUSLY_UNSUBSCRIBED_PORTAL",
	// "QUARANTINED_ADDRESS", "QUEUED", "RECIPIENT_FATIGUE_SUPPRESSED", "SENT",
	// "TEMPLATE_RENDER_EXCEPTION", "THROTTLED", "TOO_MANY_RECIPIENTS",
	// "UNCONFIGURED_SENDING_DOMAIN", "UNDELIVERABLE", "VALIDATION_FAILED".
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
	EmailSendStatusViewStatusCanceled   EmailSendStatusViewStatus = "CANCELED"
	EmailSendStatusViewStatusComplete   EmailSendStatusViewStatus = "COMPLETE"
	EmailSendStatusViewStatusPending    EmailSendStatusViewStatus = "PENDING"
	EmailSendStatusViewStatusProcessing EmailSendStatusViewStatus = "PROCESSING"
)

// Result of the send.
type EmailSendStatusViewSendResult string

const (
	EmailSendStatusViewSendResultAddressListBombed                  EmailSendStatusViewSendResult = "ADDRESS_LIST_BOMBED"
	EmailSendStatusViewSendResultAddressOnlyAcceptedOnProd          EmailSendStatusViewSendResult = "ADDRESS_ONLY_ACCEPTED_ON_PROD"
	EmailSendStatusViewSendResultAddressOptedOut                    EmailSendStatusViewSendResult = "ADDRESS_OPTED_OUT"
	EmailSendStatusViewSendResultBlockedAddress                     EmailSendStatusViewSendResult = "BLOCKED_ADDRESS"
	EmailSendStatusViewSendResultBlockedDomain                      EmailSendStatusViewSendResult = "BLOCKED_DOMAIN"
	EmailSendStatusViewSendResultCampaignCancelled                  EmailSendStatusViewSendResult = "CAMPAIGN_CANCELLED"
	EmailSendStatusViewSendResultCancelledAbuse                     EmailSendStatusViewSendResult = "CANCELLED_ABUSE"
	EmailSendStatusViewSendResultCorruptInput                       EmailSendStatusViewSendResult = "CORRUPT_INPUT"
	EmailSendStatusViewSendResultEmailDisabled                      EmailSendStatusViewSendResult = "EMAIL_DISABLED"
	EmailSendStatusViewSendResultEmailUnconfirmed                   EmailSendStatusViewSendResult = "EMAIL_UNCONFIRMED"
	EmailSendStatusViewSendResultGdprDoiEnabled                     EmailSendStatusViewSendResult = "GDPR_DOI_ENABLED"
	EmailSendStatusViewSendResultGraymailSuppressed                 EmailSendStatusViewSendResult = "GRAYMAIL_SUPPRESSED"
	EmailSendStatusViewSendResultHublLimitExceeded                  EmailSendStatusViewSendResult = "HUBL_LIMIT_EXCEEDED"
	EmailSendStatusViewSendResultIdempotentFail                     EmailSendStatusViewSendResult = "IDEMPOTENT_FAIL"
	EmailSendStatusViewSendResultIdempotentIgnore                   EmailSendStatusViewSendResult = "IDEMPOTENT_IGNORE"
	EmailSendStatusViewSendResultInvalidFromAddress                 EmailSendStatusViewSendResult = "INVALID_FROM_ADDRESS"
	EmailSendStatusViewSendResultInvalidToAddress                   EmailSendStatusViewSendResult = "INVALID_TO_ADDRESS"
	EmailSendStatusViewSendResultLowContactQualityScore             EmailSendStatusViewSendResult = "LOW_CONTACT_QUALITY_SCORE"
	EmailSendStatusViewSendResultMissingContent                     EmailSendStatusViewSendResult = "MISSING_CONTENT"
	EmailSendStatusViewSendResultMissingRequiredParameter           EmailSendStatusViewSendResult = "MISSING_REQUIRED_PARAMETER"
	EmailSendStatusViewSendResultMissingTemplateProperties          EmailSendStatusViewSendResult = "MISSING_TEMPLATE_PROPERTIES"
	EmailSendStatusViewSendResultMtaIgnore                          EmailSendStatusViewSendResult = "MTA_IGNORE"
	EmailSendStatusViewSendResultNonMarketableContact               EmailSendStatusViewSendResult = "NON_MARKETABLE_CONTACT"
	EmailSendStatusViewSendResultPortalAuthenticationFailure        EmailSendStatusViewSendResult = "PORTAL_AUTHENTICATION_FAILURE"
	EmailSendStatusViewSendResultPortalExpired                      EmailSendStatusViewSendResult = "PORTAL_EXPIRED"
	EmailSendStatusViewSendResultPortalMissingMarketingScope        EmailSendStatusViewSendResult = "PORTAL_MISSING_MARKETING_SCOPE"
	EmailSendStatusViewSendResultPortalNotAuthorizedForApplication  EmailSendStatusViewSendResult = "PORTAL_NOT_AUTHORIZED_FOR_APPLICATION"
	EmailSendStatusViewSendResultPortalOverLimit                    EmailSendStatusViewSendResult = "PORTAL_OVER_LIMIT"
	EmailSendStatusViewSendResultPortalSuspended                    EmailSendStatusViewSendResult = "PORTAL_SUSPENDED"
	EmailSendStatusViewSendResultPreviousSpam                       EmailSendStatusViewSendResult = "PREVIOUS_SPAM"
	EmailSendStatusViewSendResultPreviouslyBounced                  EmailSendStatusViewSendResult = "PREVIOUSLY_BOUNCED"
	EmailSendStatusViewSendResultPreviouslyUnsubscribedBrand        EmailSendStatusViewSendResult = "PREVIOUSLY_UNSUBSCRIBED_BRAND"
	EmailSendStatusViewSendResultPreviouslyUnsubscribedBusinessUnit EmailSendStatusViewSendResult = "PREVIOUSLY_UNSUBSCRIBED_BUSINESS_UNIT"
	EmailSendStatusViewSendResultPreviouslyUnsubscribedMessage      EmailSendStatusViewSendResult = "PREVIOUSLY_UNSUBSCRIBED_MESSAGE"
	EmailSendStatusViewSendResultPreviouslyUnsubscribedPortal       EmailSendStatusViewSendResult = "PREVIOUSLY_UNSUBSCRIBED_PORTAL"
	EmailSendStatusViewSendResultQuarantinedAddress                 EmailSendStatusViewSendResult = "QUARANTINED_ADDRESS"
	EmailSendStatusViewSendResultQueued                             EmailSendStatusViewSendResult = "QUEUED"
	EmailSendStatusViewSendResultRecipientFatigueSuppressed         EmailSendStatusViewSendResult = "RECIPIENT_FATIGUE_SUPPRESSED"
	EmailSendStatusViewSendResultSent                               EmailSendStatusViewSendResult = "SENT"
	EmailSendStatusViewSendResultTemplateRenderException            EmailSendStatusViewSendResult = "TEMPLATE_RENDER_EXCEPTION"
	EmailSendStatusViewSendResultThrottled                          EmailSendStatusViewSendResult = "THROTTLED"
	EmailSendStatusViewSendResultTooManyRecipients                  EmailSendStatusViewSendResult = "TOO_MANY_RECIPIENTS"
	EmailSendStatusViewSendResultUnconfiguredSendingDomain          EmailSendStatusViewSendResult = "UNCONFIGURED_SENDING_DOMAIN"
	EmailSendStatusViewSendResultUndeliverable                      EmailSendStatusViewSendResult = "UNDELIVERABLE"
	EmailSendStatusViewSendResultValidationFailed                   EmailSendStatusViewSendResult = "VALIDATION_FAILED"
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
