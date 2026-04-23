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
	options         []option.RequestOption
	Campaigns       CampaignService
	Emails          EmailService
	MarketingEvents MarketingEventService
	SingleSend      SingleSendService
	Transactional   TransactionalService
}

// NewMarketingService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewMarketingService(opts ...option.RequestOption) (r MarketingService) {
	r = MarketingService{}
	r.options = opts
	r.Campaigns = NewCampaignService(opts...)
	r.Emails = NewEmailService(opts...)
	r.MarketingEvents = NewMarketingEventService(opts...)
	r.SingleSend = NewSingleSendService(opts...)
	r.Transactional = NewTransactionalService(opts...)
	return
}

type EmailSendStatusView struct {
	// Status of the send request.
	//
	// Any of "CANCELED", "COMPLETE", "PENDING", "PROCESSING".
	Status EmailSendStatusViewStatus `json:"status" api:"required"`
	// Identifier used to query the status of the send.
	StatusID string `json:"statusId" api:"required"`
	// Time when the send was completed.
	CompletedAt time.Time   `json:"completedAt" format:"date-time"`
	EventID     EventIDView `json:"eventId"`
	// A human readable message describing the error along with remediation steps where
	// appropriate
	Message string `json:"message"`
	// Time when the send was requested.
	RequestedAt time.Time `json:"requestedAt" format:"date-time"`
	// Result of the send.
	//
	// Any of "ADDRESS_LIST_BOMBED", "ADDRESS_ONLY_ACCEPTED_ON_PROD",
	// "ADDRESS_OPTED_OUT", "ATTACHMENT_DOWNLOAD_QUEUE_FULL", "BLOCKED_ADDRESS",
	// "BLOCKED_DOMAIN", "BRAND_RECIPIENT_FATIGUE_SUPPRESSED", "CAMPAIGN_CANCELLED",
	// "CANCELLED_ABUSE", "CORRUPT_INPUT", "EMAIL_DISABLED", "EMAIL_UNCONFIRMED",
	// "GDPR_DOI_ENABLED", "GRAYMAIL_SUPPRESSED", "HUBL_LIMIT_EXCEEDED",
	// "IDEMPOTENT_FAIL", "IDEMPOTENT_IGNORE", "INVALID_APP_ID_ATTRIBUTION",
	// "INVALID_FROM_ADDRESS", "INVALID_TO_ADDRESS", "LOW_CONTACT_QUALITY_SCORE",
	// "MISSING_CONTENT", "MISSING_REQUIRED_PARAMETER", "MISSING_TEMPLATE_PROPERTIES",
	// "MTA_IGNORE", "NON_MARKETABLE_CONTACT", "PORTAL_AUTHENTICATION_FAILURE",
	// "PORTAL_EXPIRED", "PORTAL_MISSING_MARKETING_SCOPE",
	// "PORTAL_NOT_AUTHORIZED_FOR_APPLICATION", "PORTAL_OVER_LIMIT",
	// "PORTAL_SUSPENDED", "PREVIOUS_SPAM", "PREVIOUSLY_BOUNCED",
	// "PREVIOUSLY_UNSUBSCRIBED_BRAND", "PREVIOUSLY_UNSUBSCRIBED_BUSINESS_UNIT",
	// "PREVIOUSLY_UNSUBSCRIBED_MESSAGE", "PREVIOUSLY_UNSUBSCRIBED_PORTAL",
	// "QUARANTINED_ADDRESS", "QUEUED", "RECIPIENT_FATIGUE_SUPPRESSED", "SENT",
	// "TEMPLATE_RENDER_EXCEPTION", "THROTTLED", "TOO_MANY_RECIPIENTS",
	// "UBB_GOVERNANCE_MISSING", "UNCONFIGURED_SENDING_DOMAIN", "UNDELIVERABLE",
	// "VALIDATION_FAILED".
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
	EmailSendStatusViewSendResultAttachmentDownloadQueueFull        EmailSendStatusViewSendResult = "ATTACHMENT_DOWNLOAD_QUEUE_FULL"
	EmailSendStatusViewSendResultBlockedAddress                     EmailSendStatusViewSendResult = "BLOCKED_ADDRESS"
	EmailSendStatusViewSendResultBlockedDomain                      EmailSendStatusViewSendResult = "BLOCKED_DOMAIN"
	EmailSendStatusViewSendResultBrandRecipientFatigueSuppressed    EmailSendStatusViewSendResult = "BRAND_RECIPIENT_FATIGUE_SUPPRESSED"
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
	EmailSendStatusViewSendResultInvalidAppIDAttribution            EmailSendStatusViewSendResult = "INVALID_APP_ID_ATTRIBUTION"
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
	EmailSendStatusViewSendResultUbbGovernanceMissing               EmailSendStatusViewSendResult = "UBB_GOVERNANCE_MISSING"
	EmailSendStatusViewSendResultUnconfiguredSendingDomain          EmailSendStatusViewSendResult = "UNCONFIGURED_SENDING_DOMAIN"
	EmailSendStatusViewSendResultUndeliverable                      EmailSendStatusViewSendResult = "UNDELIVERABLE"
	EmailSendStatusViewSendResultValidationFailed                   EmailSendStatusViewSendResult = "VALIDATION_FAILED"
)

type EventIDView struct {
	// Identifier of event.
	ID string `json:"id" api:"required" format:"uuid"`
	// Time of event creation.
	Created time.Time `json:"created" api:"required" format:"date-time"`
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

// The properties Bcc, Cc, ReplyTo are required.
type PublicSingleSendEmailParam struct {
	// List of email addresses to send as Bcc.
	Bcc []string `json:"bcc,omitzero" api:"required"`
	// List of email addresses to send as Cc.
	Cc []string `json:"cc,omitzero" api:"required"`
	// List of Reply-To header values for the email.
	ReplyTo []string `json:"replyTo,omitzero" api:"required"`
	// The From header for the email.
	From param.Opt[string] `json:"from,omitzero"`
	// ID for a particular send. No more than one email will be sent per sendId.
	SendID param.Opt[string] `json:"sendId,omitzero"`
	// The recipient of the email.
	To param.Opt[string] `json:"to,omitzero"`
	paramObj
}

func (r PublicSingleSendEmailParam) MarshalJSON() (data []byte, err error) {
	type shadow PublicSingleSendEmailParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PublicSingleSendEmailParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties ContactProperties, CustomProperties, EmailID, Message are
// required.
type PublicSingleSendRequestEggParam struct {
	// The contactProperties field is a map of contact property values. Each contact
	// property value contains a name and value property. Each property will get set on
	// the contact record and will be visible in the template under {{ contact.NAME }}.
	// Use these properties when you want to set a contact property while you’re
	// sending the email. For example, when sending a receipt you may want to set a
	// last_paid_date property, as the sending of the receipt will have information
	// about the last payment.
	ContactProperties map[string]string `json:"contactProperties,omitzero" api:"required"`
	// The customProperties field is a map of property values. Each property value
	// contains a name and value property. Each property will be visible in the
	// template under {{ custom.NAME }}. Note: Custom properties do not currently
	// support arrays. To provide a listing in an email, one workaround is to build an
	// HTML list (either with tables or ul) and specify it as a custom property.
	CustomProperties map[string]any `json:"customProperties,omitzero" api:"required"`
	// The content ID for the email, which can be found in email tool UI.
	EmailID int64                      `json:"emailId" api:"required"`
	Message PublicSingleSendEmailParam `json:"message,omitzero" api:"required"`
	paramObj
}

func (r PublicSingleSendRequestEggParam) MarshalJSON() (data []byte, err error) {
	type shadow PublicSingleSendRequestEggParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PublicSingleSendRequestEggParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
