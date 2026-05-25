// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package scheduler

import (
	"encoding/json"
	"time"

	"github.com/HubSpot/hubspot-sdk-go/internal/apijson"
	"github.com/HubSpot/hubspot-sdk-go/option"
	"github.com/HubSpot/hubspot-sdk-go/packages/param"
	"github.com/HubSpot/hubspot-sdk-go/packages/respjson"
	"github.com/HubSpot/hubspot-sdk-go/shared"
)

// MeetingService contains methods and other services that help with interacting
// with the hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMeetingService] method instead.
type MeetingService struct {
	options  []option.RequestOption
	Advanced MeetingAdvancedService
	Basic    MeetingBasicService
}

// NewMeetingService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewMeetingService(opts ...option.RequestOption) (r MeetingService) {
	r = MeetingService{}
	r.options = opts
	r.Advanced = NewMeetingAdvancedService(opts...)
	r.Basic = NewMeetingBasicService(opts...)
	return
}

type CollectionResponseWithTotalExternalLinkMetadata struct {
	Results []ExternalLinkMetadata `json:"results" api:"required"`
	Total   int64                  `json:"total" api:"required"`
	Paging  shared.Paging          `json:"paging"`
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
func (r CollectionResponseWithTotalExternalLinkMetadata) RawJSON() string { return r.JSON.raw }
func (r *CollectionResponseWithTotalExternalLinkMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties To, Types are required.
type ExternalAssociationCreateRequestParam struct {
	// Contains the Id of a Public Object
	To    shared.PublicObjectIDParam    `json:"to,omitzero" api:"required"`
	Types []shared.AssociationSpecParam `json:"types,omitzero" api:"required"`
	paramObj
}

func (r ExternalAssociationCreateRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow ExternalAssociationCreateRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExternalAssociationCreateRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Name, Value are required.
type ExternalBookingFormFieldParam struct {
	// The name of the form field.
	Name string `json:"name" api:"required"`
	// The value associated with the form field.
	Value string `json:"value" api:"required"`
	paramObj
}

func (r ExternalBookingFormFieldParam) MarshalJSON() (data []byte, err error) {
	type shadow ExternalBookingFormFieldParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExternalBookingFormFieldParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExternalBookingInfo struct {
	AllUsersBusyTimes []ExternalUserBusyTimes      `json:"allUsersBusyTimes" api:"required"`
	CustomParams      ExternalMeetingsLinkSettings `json:"customParams" api:"required"`
	// Whether the meeting was booked offline, meaning no associated calendar event was
	// created.
	IsOffline bool `json:"isOffline" api:"required"`
	// The unique identifier for the meeting link.
	LinkID string `json:"linkId" api:"required"`
	// The type of the meeting link. Accepted values are: GROUP_CALENDAR,
	// PERSONAL_LINK, ROUND_ROBIN_CALENDAR.
	//
	// Any of "GROUP_CALENDAR", "PERSONAL_LINK", "ROUND_ROBIN_CALENDAR".
	LinkType         ExternalBookingInfoLinkType `json:"linkType" api:"required"`
	BrandingMetadata ExternalBrandingMetadata    `json:"brandingMetadata"`
	LinkAvailability ExternalLinkAvailability    `json:"linkAvailability"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AllUsersBusyTimes respjson.Field
		CustomParams      respjson.Field
		IsOffline         respjson.Field
		LinkID            respjson.Field
		LinkType          respjson.Field
		BrandingMetadata  respjson.Field
		LinkAvailability  respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExternalBookingInfo) RawJSON() string { return r.JSON.raw }
func (r *ExternalBookingInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The type of the meeting link. Accepted values are: GROUP_CALENDAR,
// PERSONAL_LINK, ROUND_ROBIN_CALENDAR.
type ExternalBookingInfoLinkType string

const (
	ExternalBookingInfoLinkTypeGroupCalendar      ExternalBookingInfoLinkType = "GROUP_CALENDAR"
	ExternalBookingInfoLinkTypePersonalLink       ExternalBookingInfoLinkType = "PERSONAL_LINK"
	ExternalBookingInfoLinkTypeRoundRobinCalendar ExternalBookingInfoLinkType = "ROUND_ROBIN_CALENDAR"
)

type ExternalBrandingMetadata struct {
	// The alternative text for the current logo.
	LogoAltText string `json:"logoAltText" api:"required"`
	// Whether Hubspot Marketing ads are shown.
	ShowMarketingAd bool `json:"showMarketingAd" api:"required"`
	// Whether Hubspot Sales ads are shown.
	ShowSalesAd bool `json:"showSalesAd" api:"required"`
	// The secondary accent color used in branding.
	Accent2Color string `json:"accent2Color"`
	// The primary accent color used in branding.
	AccentColor string `json:"accentColor"`
	// The first line of the company's address.
	CompanyAddressLine1 string `json:"companyAddressLine1"`
	// The second line of the company's address.
	CompanyAddressLine2 string `json:"companyAddressLine2"`
	// The URL of the company's avatar image.
	CompanyAvatar string `json:"companyAvatar"`
	// The city where the company is located.
	CompanyCity string `json:"companyCity"`
	// The country where the company is located.
	CompanyCountry string `json:"companyCountry"`
	// The domain of the company's website.
	CompanyDomain string `json:"companyDomain"`
	// The name of the company.
	CompanyName string `json:"companyName"`
	// The state where the company is located.
	CompanyState string `json:"companyState"`
	// The ZIP code of the company's location.
	CompanyZip string `json:"companyZip"`
	// The height of the logo in pixels.
	LogoHeight int64 `json:"logoHeight"`
	// The URL of a custom logo image.
	LogoURL string `json:"logoUrl"`
	// The width of the logo in pixels.
	LogoWidth int64 `json:"logoWidth"`
	// The primary color used in branding.
	PrimaryColor string `json:"primaryColor"`
	// The secondary color used in branding.
	SecondaryColor string `json:"secondaryColor"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		LogoAltText         respjson.Field
		ShowMarketingAd     respjson.Field
		ShowSalesAd         respjson.Field
		Accent2Color        respjson.Field
		AccentColor         respjson.Field
		CompanyAddressLine1 respjson.Field
		CompanyAddressLine2 respjson.Field
		CompanyAvatar       respjson.Field
		CompanyCity         respjson.Field
		CompanyCountry      respjson.Field
		CompanyDomain       respjson.Field
		CompanyName         respjson.Field
		CompanyState        respjson.Field
		CompanyZip          respjson.Field
		LogoHeight          respjson.Field
		LogoURL             respjson.Field
		LogoWidth           respjson.Field
		PrimaryColor        respjson.Field
		SecondaryColor      respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExternalBrandingMetadata) RawJSON() string { return r.JSON.raw }
func (r *ExternalBrandingMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties HsMeetingEndTime, HsMeetingOutcome, HsMeetingStartTime,
// HsMeetingTitle, HsTimestamp, HubSpotOwnerID are required.
type ExternalCalendarMeetingEventCreatePropertiesParam struct {
	// The time that the meeting should end in ISO 8601 format.
	HsMeetingEndTime time.Time `json:"hs_meeting_end_time" api:"required" format:"date-time"`
	// The outcome of the meeting. Acceptable default values are: SCHEDULED, COMPLETED,
	// RESCHEDULED, NO_SHOW, CANCELED. This property can be changed to include
	// additional custom values.
	HsMeetingOutcome string `json:"hs_meeting_outcome" api:"required"`
	// The time that the meeting should start in ISO 8601 format.
	HsMeetingStartTime time.Time `json:"hs_meeting_start_time" api:"required" format:"date-time"`
	// The title of the meeting and calendar event.
	HsMeetingTitle string `json:"hs_meeting_title" api:"required"`
	// The time that the meeting should start in ISO 8601 format. This value should be
	// the same as `hs_meeting_start_time`.
	HsTimestamp time.Time `json:"hs_timestamp" api:"required" format:"date-time"`
	// The ownerId of the HubSpot user who will host the meeting.
	HubSpotOwnerID string `json:"hubspot_owner_id" api:"required"`
	// The activity type of the meeting. Acceptable values are based on portal defined
	// call and meeting types.
	HsActivityType param.Opt[string] `json:"hs_activity_type,omitzero"`
	// Internal notes related to the meeting.
	HsInternalMeetingNotes param.Opt[string] `json:"hs_internal_meeting_notes,omitzero"`
	// The description of the meeting and calendar event.
	HsMeetingBody param.Opt[string] `json:"hs_meeting_body,omitzero"`
	// The physical address, virtual location, or phone number where the meeting will
	// take place.
	HsMeetingLocation  param.Opt[string] `json:"hs_meeting_location,omitzero"`
	HsAttachmentIDs    []string          `json:"hs_attachment_ids,omitzero"`
	HsAttendeeOwnerIDs []string          `json:"hs_attendee_owner_ids,omitzero"`
	// The type of location for the meeting. Acceptable values are: ADDRESS, CUSTOM,
	// PHONE.
	//
	// Any of "ADDRESS", "CUSTOM", "PHONE".
	HsMeetingLocationType ExternalCalendarMeetingEventCreatePropertiesHsMeetingLocationType `json:"hs_meeting_location_type,omitzero"`
	paramObj
}

func (r ExternalCalendarMeetingEventCreatePropertiesParam) MarshalJSON() (data []byte, err error) {
	type shadow ExternalCalendarMeetingEventCreatePropertiesParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExternalCalendarMeetingEventCreatePropertiesParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The type of location for the meeting. Acceptable values are: ADDRESS, CUSTOM,
// PHONE.
type ExternalCalendarMeetingEventCreatePropertiesHsMeetingLocationType string

const (
	ExternalCalendarMeetingEventCreatePropertiesHsMeetingLocationTypeAddress ExternalCalendarMeetingEventCreatePropertiesHsMeetingLocationType = "ADDRESS"
	ExternalCalendarMeetingEventCreatePropertiesHsMeetingLocationTypeCustom  ExternalCalendarMeetingEventCreatePropertiesHsMeetingLocationType = "CUSTOM"
	ExternalCalendarMeetingEventCreatePropertiesHsMeetingLocationTypePhone   ExternalCalendarMeetingEventCreatePropertiesHsMeetingLocationType = "PHONE"
)

// The properties Associations, EmailReminderSchedule, Properties, Timezone are
// required.
type ExternalCalendarMeetingEventCreateRequestParam struct {
	Associations          []ExternalAssociationCreateRequestParam           `json:"associations,omitzero" api:"required"`
	EmailReminderSchedule ExternalEmailReminderScheduleParam                `json:"emailReminderSchedule,omitzero" api:"required"`
	Properties            ExternalCalendarMeetingEventCreatePropertiesParam `json:"properties,omitzero" api:"required"`
	// The timezone property that will be set on the meeting event.
	Timezone string `json:"timezone" api:"required"`
	paramObj
}

func (r ExternalCalendarMeetingEventCreateRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow ExternalCalendarMeetingEventCreateRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExternalCalendarMeetingEventCreateRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExternalCalendarMeetingEventResponseProperties struct {
	// The source of the engagement, will always be `MEETINGS`.
	//
	// Any of "ACADEMY", "ACCEPTANCE_TEST", "ACTIVITY_AUTO_ASSOCIATE",
	// "ACTIVITY_LOG_REVERT", "ADS", "AI_GROUP", "ANALYTICS", "API", "APPROVALS",
	// "ASSISTS", "ASSOCIATIONS", "AUTO_ASSOCIATE_BY_DOMAIN", "AUTOMATION_JOURNEY",
	// "AUTOMATION_PLATFORM", "AVATARS_SERVICE", "BATCH_UPDATE", "BCC_TO_CRM",
	// "BEHAVIORAL_EVENTS", "BET_ASSIGNMENT", "BET_CRM_CONNECTOR", "BIDEN", "BILLING",
	// "BOT", "BREEZE_AGENT", "CALCULATED", "CENTRAL_EXCHANGE_RATES", "CHATSPOT",
	// "CLONE_OBJECTS", "COMMUNICATOR", "COMPANIES", "COMPANY_FAMILIES",
	// "COMPANY_INSIGHTS", "CONNECTED_ACCOUNT", "CONTACTS", "CONTACTS_WEB",
	// "CONTENT_MEMBERSHIP", "CONVERSATIONAL_ENRICHMENT", "CONVERSATIONS",
	// "CRM_PROCESSES_PLATFORM", "CRM_UI", "CRM_UI_BULK_ACTION", "CUSTOMER_AGENT",
	// "DATA_ENRICHMENT", "DATA_QUALITY", "DATASET", "DEALS", "DEFAULT",
	// "DELETE_OBJECTS", "DI_WRITE_TO_CRM", "EMAIL", "EMAIL_INBOX_IMPORT",
	// "EMAIL_INTEGRATION", "ENGAGEMENTS", "EXTENSION", "FILE_MANAGER",
	// "FLYWHEEL_PRODUCT_DATA_SYNC", "FORECASTING", "FORM", "FORWARD_TO_CRM",
	// "GMAIL_INTEGRATION", "GOALS", "HEISENBERG", "HELP_DESK", "HELP_DESK_AI",
	// "IMPORT", "INTEGRATION", "INTEGRATIONS_PLATFORM", "INTEGRATIONS_SYNC", "INTENT",
	// "INTERNAL_PROCESSING", "LEADIN", "LEGAL_BASIS_REMEDIATION", "MARKET_SOURCING",
	// "MARKETPLACE", "MARKETS", "MEETINGS", "MERGE_COMPANIES", "MERGE_CONTACTS",
	// "MERGE_OBJECTS", "MERGE_REVERT_OBJECTS", "MICROAPPS", "MIGRATION",
	// "MOBILE_ANDROID", "MOBILE_IOS", "PAYMENTS", "PIPELINE_SETTINGS", "PLAYBOOKS",
	// "PORTAL_OBJECT_SYNC", "PORTAL_USER_ASSOCIATOR", "PRESENTATIONS",
	// "PRIMARY_AUTOMATION", "PROPERTY_DEFAULT_VALUE", "PROPERTY_RESTORE",
	// "PROPERTY_SETTINGS", "PROSPECTING_AGENT", "QUOTAS", "QUOTES", "RECYCLING_BIN",
	// "RESTORE_OBJECTS", "REVENUE_PLATFORM", "SALES", "SALES_MESSAGES", "SALESFORCE",
	// "SEQUENCES", "SETTINGS", "SIDEKICK", "SIGNALS", "SLACK_INTEGRATION",
	// "SMART_DATA_CAPTURE", "SOCIAL", "SUCCESS", "TALLY", "TASK", "UNKNOWN",
	// "WAL_INCREMENTAL", "WORK_UI", "WORKFLOW_CONTACT_DELETE_ACTION", "WORKFLOWS".
	HsEngagementSource ExternalCalendarMeetingEventResponsePropertiesHsEngagementSource `json:"hs_engagement_source" api:"required"`
	// The ID associated with the process created the engagement. Should always be
	// empty when creating meeting events through this API.
	HsEngagementSourceID string `json:"hs_engagement_source_id" api:"required"`
	// The end time of the meeting in ISO 8601 format.
	HsMeetingEndTime time.Time `json:"hs_meeting_end_time" api:"required" format:"date-time"`
	// The outcome of the meeting. Acceptable default values are: SCHEDULED, COMPLETED,
	// RESCHEDULED, NO_SHOW, CANCELED. This property can be changed to include
	// additional custom values.
	HsMeetingOutcome string `json:"hs_meeting_outcome" api:"required"`
	// The start time of the meeting in ISO 8601 format.
	HsMeetingStartTime time.Time `json:"hs_meeting_start_time" api:"required" format:"date-time"`
	// The title of the meeting and calendar event.
	HsMeetingTitle string `json:"hs_meeting_title" api:"required"`
	// The time that the meeting should start in ISO 8601 format. This value should be
	// the same as `hs_meeting_start_time`.
	HsTimestamp time.Time `json:"hs_timestamp" api:"required" format:"date-time"`
	// The activity type of the meeting. Acceptable values are based on portal defined
	// call and meeting types.
	HsActivityType     string   `json:"hs_activity_type"`
	HsAttachmentIDs    []string `json:"hs_attachment_ids"`
	HsAttendeeOwnerIDs []string `json:"hs_attendee_owner_ids"`
	// Whether to include the meeting description in the reminder.
	HsIncludeDescriptionInReminder string `json:"hs_include_description_in_reminder"`
	// Internal notes related to the meeting.
	HsInternalMeetingNotes string `json:"hs_internal_meeting_notes"`
	// The description of the meeting and calendar event.
	HsMeetingBody string `json:"hs_meeting_body"`
	// The calendar event URL for the meeting.
	HsMeetingExternalURL string `json:"hs_meeting_external_url"`
	// The physical address, virtual location, or phone number where the meeting will
	// take place.
	HsMeetingLocation string `json:"hs_meeting_location"`
	// The type of location for the meeting. Acceptable values are: ADDRESS, CUSTOM,
	// PHONE.
	//
	// Any of "ADDRESS", "CUSTOM", "PHONE".
	HsMeetingLocationType ExternalCalendarMeetingEventResponsePropertiesHsMeetingLocationType `json:"hs_meeting_location_type"`
	// The unique ID of the created calendar event.
	HsUniqueID string `json:"hs_unique_id"`
	// The owner ID of the HubSpot user hosting the meeting.
	HubSpotOwnerID string `json:"hubspot_owner_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		HsEngagementSource             respjson.Field
		HsEngagementSourceID           respjson.Field
		HsMeetingEndTime               respjson.Field
		HsMeetingOutcome               respjson.Field
		HsMeetingStartTime             respjson.Field
		HsMeetingTitle                 respjson.Field
		HsTimestamp                    respjson.Field
		HsActivityType                 respjson.Field
		HsAttachmentIDs                respjson.Field
		HsAttendeeOwnerIDs             respjson.Field
		HsIncludeDescriptionInReminder respjson.Field
		HsInternalMeetingNotes         respjson.Field
		HsMeetingBody                  respjson.Field
		HsMeetingExternalURL           respjson.Field
		HsMeetingLocation              respjson.Field
		HsMeetingLocationType          respjson.Field
		HsUniqueID                     respjson.Field
		HubSpotOwnerID                 respjson.Field
		ExtraFields                    map[string]respjson.Field
		raw                            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExternalCalendarMeetingEventResponseProperties) RawJSON() string { return r.JSON.raw }
func (r *ExternalCalendarMeetingEventResponseProperties) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The source of the engagement, will always be `MEETINGS`.
type ExternalCalendarMeetingEventResponsePropertiesHsEngagementSource string

const (
	ExternalCalendarMeetingEventResponsePropertiesHsEngagementSourceAcademy                     ExternalCalendarMeetingEventResponsePropertiesHsEngagementSource = "ACADEMY"
	ExternalCalendarMeetingEventResponsePropertiesHsEngagementSourceAcceptanceTest              ExternalCalendarMeetingEventResponsePropertiesHsEngagementSource = "ACCEPTANCE_TEST"
	ExternalCalendarMeetingEventResponsePropertiesHsEngagementSourceActivityAutoAssociate       ExternalCalendarMeetingEventResponsePropertiesHsEngagementSource = "ACTIVITY_AUTO_ASSOCIATE"
	ExternalCalendarMeetingEventResponsePropertiesHsEngagementSourceActivityLogRevert           ExternalCalendarMeetingEventResponsePropertiesHsEngagementSource = "ACTIVITY_LOG_REVERT"
	ExternalCalendarMeetingEventResponsePropertiesHsEngagementSourceAds                         ExternalCalendarMeetingEventResponsePropertiesHsEngagementSource = "ADS"
	ExternalCalendarMeetingEventResponsePropertiesHsEngagementSourceAIGroup                     ExternalCalendarMeetingEventResponsePropertiesHsEngagementSource = "AI_GROUP"
	ExternalCalendarMeetingEventResponsePropertiesHsEngagementSourceAnalytics                   ExternalCalendarMeetingEventResponsePropertiesHsEngagementSource = "ANALYTICS"
	ExternalCalendarMeetingEventResponsePropertiesHsEngagementSourceAPI                         ExternalCalendarMeetingEventResponsePropertiesHsEngagementSource = "API"
	ExternalCalendarMeetingEventResponsePropertiesHsEngagementSourceApprovals                   ExternalCalendarMeetingEventResponsePropertiesHsEngagementSource = "APPROVALS"
	ExternalCalendarMeetingEventResponsePropertiesHsEngagementSourceAssists                     ExternalCalendarMeetingEventResponsePropertiesHsEngagementSource = "ASSISTS"
	ExternalCalendarMeetingEventResponsePropertiesHsEngagementSourceAssociations                ExternalCalendarMeetingEventResponsePropertiesHsEngagementSource = "ASSOCIATIONS"
	ExternalCalendarMeetingEventResponsePropertiesHsEngagementSourceAutoAssociateByDomain       ExternalCalendarMeetingEventResponsePropertiesHsEngagementSource = "AUTO_ASSOCIATE_BY_DOMAIN"
	ExternalCalendarMeetingEventResponsePropertiesHsEngagementSourceAutomationJourney           ExternalCalendarMeetingEventResponsePropertiesHsEngagementSource = "AUTOMATION_JOURNEY"
	ExternalCalendarMeetingEventResponsePropertiesHsEngagementSourceAutomationPlatform          ExternalCalendarMeetingEventResponsePropertiesHsEngagementSource = "AUTOMATION_PLATFORM"
	ExternalCalendarMeetingEventResponsePropertiesHsEngagementSourceAvatarsService              ExternalCalendarMeetingEventResponsePropertiesHsEngagementSource = "AVATARS_SERVICE"
	ExternalCalendarMeetingEventResponsePropertiesHsEngagementSourceBatchUpdate                 ExternalCalendarMeetingEventResponsePropertiesHsEngagementSource = "BATCH_UPDATE"
	ExternalCalendarMeetingEventResponsePropertiesHsEngagementSourceBccToCrm                    ExternalCalendarMeetingEventResponsePropertiesHsEngagementSource = "BCC_TO_CRM"
	ExternalCalendarMeetingEventResponsePropertiesHsEngagementSourceBehavioralEvents            ExternalCalendarMeetingEventResponsePropertiesHsEngagementSource = "BEHAVIORAL_EVENTS"
	ExternalCalendarMeetingEventResponsePropertiesHsEngagementSourceBetAssignment               ExternalCalendarMeetingEventResponsePropertiesHsEngagementSource = "BET_ASSIGNMENT"
	ExternalCalendarMeetingEventResponsePropertiesHsEngagementSourceBetCrmConnector             ExternalCalendarMeetingEventResponsePropertiesHsEngagementSource = "BET_CRM_CONNECTOR"
	ExternalCalendarMeetingEventResponsePropertiesHsEngagementSourceBiden                       ExternalCalendarMeetingEventResponsePropertiesHsEngagementSource = "BIDEN"
	ExternalCalendarMeetingEventResponsePropertiesHsEngagementSourceBilling                     ExternalCalendarMeetingEventResponsePropertiesHsEngagementSource = "BILLING"
	ExternalCalendarMeetingEventResponsePropertiesHsEngagementSourceBot                         ExternalCalendarMeetingEventResponsePropertiesHsEngagementSource = "BOT"
	ExternalCalendarMeetingEventResponsePropertiesHsEngagementSourceBreezeAgent                 ExternalCalendarMeetingEventResponsePropertiesHsEngagementSource = "BREEZE_AGENT"
	ExternalCalendarMeetingEventResponsePropertiesHsEngagementSourceCalculated                  ExternalCalendarMeetingEventResponsePropertiesHsEngagementSource = "CALCULATED"
	ExternalCalendarMeetingEventResponsePropertiesHsEngagementSourceCentralExchangeRates        ExternalCalendarMeetingEventResponsePropertiesHsEngagementSource = "CENTRAL_EXCHANGE_RATES"
	ExternalCalendarMeetingEventResponsePropertiesHsEngagementSourceChatspot                    ExternalCalendarMeetingEventResponsePropertiesHsEngagementSource = "CHATSPOT"
	ExternalCalendarMeetingEventResponsePropertiesHsEngagementSourceCloneObjects                ExternalCalendarMeetingEventResponsePropertiesHsEngagementSource = "CLONE_OBJECTS"
	ExternalCalendarMeetingEventResponsePropertiesHsEngagementSourceCommunicator                ExternalCalendarMeetingEventResponsePropertiesHsEngagementSource = "COMMUNICATOR"
	ExternalCalendarMeetingEventResponsePropertiesHsEngagementSourceCompanies                   ExternalCalendarMeetingEventResponsePropertiesHsEngagementSource = "COMPANIES"
	ExternalCalendarMeetingEventResponsePropertiesHsEngagementSourceCompanyFamilies             ExternalCalendarMeetingEventResponsePropertiesHsEngagementSource = "COMPANY_FAMILIES"
	ExternalCalendarMeetingEventResponsePropertiesHsEngagementSourceCompanyInsights             ExternalCalendarMeetingEventResponsePropertiesHsEngagementSource = "COMPANY_INSIGHTS"
	ExternalCalendarMeetingEventResponsePropertiesHsEngagementSourceConnectedAccount            ExternalCalendarMeetingEventResponsePropertiesHsEngagementSource = "CONNECTED_ACCOUNT"
	ExternalCalendarMeetingEventResponsePropertiesHsEngagementSourceContacts                    ExternalCalendarMeetingEventResponsePropertiesHsEngagementSource = "CONTACTS"
	ExternalCalendarMeetingEventResponsePropertiesHsEngagementSourceContactsWeb                 ExternalCalendarMeetingEventResponsePropertiesHsEngagementSource = "CONTACTS_WEB"
	ExternalCalendarMeetingEventResponsePropertiesHsEngagementSourceContentMembership           ExternalCalendarMeetingEventResponsePropertiesHsEngagementSource = "CONTENT_MEMBERSHIP"
	ExternalCalendarMeetingEventResponsePropertiesHsEngagementSourceConversationalEnrichment    ExternalCalendarMeetingEventResponsePropertiesHsEngagementSource = "CONVERSATIONAL_ENRICHMENT"
	ExternalCalendarMeetingEventResponsePropertiesHsEngagementSourceConversations               ExternalCalendarMeetingEventResponsePropertiesHsEngagementSource = "CONVERSATIONS"
	ExternalCalendarMeetingEventResponsePropertiesHsEngagementSourceCrmProcessesPlatform        ExternalCalendarMeetingEventResponsePropertiesHsEngagementSource = "CRM_PROCESSES_PLATFORM"
	ExternalCalendarMeetingEventResponsePropertiesHsEngagementSourceCrmUi                       ExternalCalendarMeetingEventResponsePropertiesHsEngagementSource = "CRM_UI"
	ExternalCalendarMeetingEventResponsePropertiesHsEngagementSourceCrmUiBulkAction             ExternalCalendarMeetingEventResponsePropertiesHsEngagementSource = "CRM_UI_BULK_ACTION"
	ExternalCalendarMeetingEventResponsePropertiesHsEngagementSourceCustomerAgent               ExternalCalendarMeetingEventResponsePropertiesHsEngagementSource = "CUSTOMER_AGENT"
	ExternalCalendarMeetingEventResponsePropertiesHsEngagementSourceDataEnrichment              ExternalCalendarMeetingEventResponsePropertiesHsEngagementSource = "DATA_ENRICHMENT"
	ExternalCalendarMeetingEventResponsePropertiesHsEngagementSourceDataQuality                 ExternalCalendarMeetingEventResponsePropertiesHsEngagementSource = "DATA_QUALITY"
	ExternalCalendarMeetingEventResponsePropertiesHsEngagementSourceDataset                     ExternalCalendarMeetingEventResponsePropertiesHsEngagementSource = "DATASET"
	ExternalCalendarMeetingEventResponsePropertiesHsEngagementSourceDeals                       ExternalCalendarMeetingEventResponsePropertiesHsEngagementSource = "DEALS"
	ExternalCalendarMeetingEventResponsePropertiesHsEngagementSourceDefault                     ExternalCalendarMeetingEventResponsePropertiesHsEngagementSource = "DEFAULT"
	ExternalCalendarMeetingEventResponsePropertiesHsEngagementSourceDeleteObjects               ExternalCalendarMeetingEventResponsePropertiesHsEngagementSource = "DELETE_OBJECTS"
	ExternalCalendarMeetingEventResponsePropertiesHsEngagementSourceDiWriteToCrm                ExternalCalendarMeetingEventResponsePropertiesHsEngagementSource = "DI_WRITE_TO_CRM"
	ExternalCalendarMeetingEventResponsePropertiesHsEngagementSourceEmail                       ExternalCalendarMeetingEventResponsePropertiesHsEngagementSource = "EMAIL"
	ExternalCalendarMeetingEventResponsePropertiesHsEngagementSourceEmailInboxImport            ExternalCalendarMeetingEventResponsePropertiesHsEngagementSource = "EMAIL_INBOX_IMPORT"
	ExternalCalendarMeetingEventResponsePropertiesHsEngagementSourceEmailIntegration            ExternalCalendarMeetingEventResponsePropertiesHsEngagementSource = "EMAIL_INTEGRATION"
	ExternalCalendarMeetingEventResponsePropertiesHsEngagementSourceEngagements                 ExternalCalendarMeetingEventResponsePropertiesHsEngagementSource = "ENGAGEMENTS"
	ExternalCalendarMeetingEventResponsePropertiesHsEngagementSourceExtension                   ExternalCalendarMeetingEventResponsePropertiesHsEngagementSource = "EXTENSION"
	ExternalCalendarMeetingEventResponsePropertiesHsEngagementSourceFileManager                 ExternalCalendarMeetingEventResponsePropertiesHsEngagementSource = "FILE_MANAGER"
	ExternalCalendarMeetingEventResponsePropertiesHsEngagementSourceFlywheelProductDataSync     ExternalCalendarMeetingEventResponsePropertiesHsEngagementSource = "FLYWHEEL_PRODUCT_DATA_SYNC"
	ExternalCalendarMeetingEventResponsePropertiesHsEngagementSourceForecasting                 ExternalCalendarMeetingEventResponsePropertiesHsEngagementSource = "FORECASTING"
	ExternalCalendarMeetingEventResponsePropertiesHsEngagementSourceForm                        ExternalCalendarMeetingEventResponsePropertiesHsEngagementSource = "FORM"
	ExternalCalendarMeetingEventResponsePropertiesHsEngagementSourceForwardToCrm                ExternalCalendarMeetingEventResponsePropertiesHsEngagementSource = "FORWARD_TO_CRM"
	ExternalCalendarMeetingEventResponsePropertiesHsEngagementSourceGmailIntegration            ExternalCalendarMeetingEventResponsePropertiesHsEngagementSource = "GMAIL_INTEGRATION"
	ExternalCalendarMeetingEventResponsePropertiesHsEngagementSourceGoals                       ExternalCalendarMeetingEventResponsePropertiesHsEngagementSource = "GOALS"
	ExternalCalendarMeetingEventResponsePropertiesHsEngagementSourceHeisenberg                  ExternalCalendarMeetingEventResponsePropertiesHsEngagementSource = "HEISENBERG"
	ExternalCalendarMeetingEventResponsePropertiesHsEngagementSourceHelpDesk                    ExternalCalendarMeetingEventResponsePropertiesHsEngagementSource = "HELP_DESK"
	ExternalCalendarMeetingEventResponsePropertiesHsEngagementSourceHelpDeskAI                  ExternalCalendarMeetingEventResponsePropertiesHsEngagementSource = "HELP_DESK_AI"
	ExternalCalendarMeetingEventResponsePropertiesHsEngagementSourceImport                      ExternalCalendarMeetingEventResponsePropertiesHsEngagementSource = "IMPORT"
	ExternalCalendarMeetingEventResponsePropertiesHsEngagementSourceIntegration                 ExternalCalendarMeetingEventResponsePropertiesHsEngagementSource = "INTEGRATION"
	ExternalCalendarMeetingEventResponsePropertiesHsEngagementSourceIntegrationsPlatform        ExternalCalendarMeetingEventResponsePropertiesHsEngagementSource = "INTEGRATIONS_PLATFORM"
	ExternalCalendarMeetingEventResponsePropertiesHsEngagementSourceIntegrationsSync            ExternalCalendarMeetingEventResponsePropertiesHsEngagementSource = "INTEGRATIONS_SYNC"
	ExternalCalendarMeetingEventResponsePropertiesHsEngagementSourceIntent                      ExternalCalendarMeetingEventResponsePropertiesHsEngagementSource = "INTENT"
	ExternalCalendarMeetingEventResponsePropertiesHsEngagementSourceInternalProcessing          ExternalCalendarMeetingEventResponsePropertiesHsEngagementSource = "INTERNAL_PROCESSING"
	ExternalCalendarMeetingEventResponsePropertiesHsEngagementSourceLeadin                      ExternalCalendarMeetingEventResponsePropertiesHsEngagementSource = "LEADIN"
	ExternalCalendarMeetingEventResponsePropertiesHsEngagementSourceLegalBasisRemediation       ExternalCalendarMeetingEventResponsePropertiesHsEngagementSource = "LEGAL_BASIS_REMEDIATION"
	ExternalCalendarMeetingEventResponsePropertiesHsEngagementSourceMarketSourcing              ExternalCalendarMeetingEventResponsePropertiesHsEngagementSource = "MARKET_SOURCING"
	ExternalCalendarMeetingEventResponsePropertiesHsEngagementSourceMarketplace                 ExternalCalendarMeetingEventResponsePropertiesHsEngagementSource = "MARKETPLACE"
	ExternalCalendarMeetingEventResponsePropertiesHsEngagementSourceMarkets                     ExternalCalendarMeetingEventResponsePropertiesHsEngagementSource = "MARKETS"
	ExternalCalendarMeetingEventResponsePropertiesHsEngagementSourceMeetings                    ExternalCalendarMeetingEventResponsePropertiesHsEngagementSource = "MEETINGS"
	ExternalCalendarMeetingEventResponsePropertiesHsEngagementSourceMergeCompanies              ExternalCalendarMeetingEventResponsePropertiesHsEngagementSource = "MERGE_COMPANIES"
	ExternalCalendarMeetingEventResponsePropertiesHsEngagementSourceMergeContacts               ExternalCalendarMeetingEventResponsePropertiesHsEngagementSource = "MERGE_CONTACTS"
	ExternalCalendarMeetingEventResponsePropertiesHsEngagementSourceMergeObjects                ExternalCalendarMeetingEventResponsePropertiesHsEngagementSource = "MERGE_OBJECTS"
	ExternalCalendarMeetingEventResponsePropertiesHsEngagementSourceMergeRevertObjects          ExternalCalendarMeetingEventResponsePropertiesHsEngagementSource = "MERGE_REVERT_OBJECTS"
	ExternalCalendarMeetingEventResponsePropertiesHsEngagementSourceMicroapps                   ExternalCalendarMeetingEventResponsePropertiesHsEngagementSource = "MICROAPPS"
	ExternalCalendarMeetingEventResponsePropertiesHsEngagementSourceMigration                   ExternalCalendarMeetingEventResponsePropertiesHsEngagementSource = "MIGRATION"
	ExternalCalendarMeetingEventResponsePropertiesHsEngagementSourceMobileAndroid               ExternalCalendarMeetingEventResponsePropertiesHsEngagementSource = "MOBILE_ANDROID"
	ExternalCalendarMeetingEventResponsePropertiesHsEngagementSourceMobileIos                   ExternalCalendarMeetingEventResponsePropertiesHsEngagementSource = "MOBILE_IOS"
	ExternalCalendarMeetingEventResponsePropertiesHsEngagementSourcePayments                    ExternalCalendarMeetingEventResponsePropertiesHsEngagementSource = "PAYMENTS"
	ExternalCalendarMeetingEventResponsePropertiesHsEngagementSourcePipelineSettings            ExternalCalendarMeetingEventResponsePropertiesHsEngagementSource = "PIPELINE_SETTINGS"
	ExternalCalendarMeetingEventResponsePropertiesHsEngagementSourcePlaybooks                   ExternalCalendarMeetingEventResponsePropertiesHsEngagementSource = "PLAYBOOKS"
	ExternalCalendarMeetingEventResponsePropertiesHsEngagementSourcePortalObjectSync            ExternalCalendarMeetingEventResponsePropertiesHsEngagementSource = "PORTAL_OBJECT_SYNC"
	ExternalCalendarMeetingEventResponsePropertiesHsEngagementSourcePortalUserAssociator        ExternalCalendarMeetingEventResponsePropertiesHsEngagementSource = "PORTAL_USER_ASSOCIATOR"
	ExternalCalendarMeetingEventResponsePropertiesHsEngagementSourcePresentations               ExternalCalendarMeetingEventResponsePropertiesHsEngagementSource = "PRESENTATIONS"
	ExternalCalendarMeetingEventResponsePropertiesHsEngagementSourcePrimaryAutomation           ExternalCalendarMeetingEventResponsePropertiesHsEngagementSource = "PRIMARY_AUTOMATION"
	ExternalCalendarMeetingEventResponsePropertiesHsEngagementSourcePropertyDefaultValue        ExternalCalendarMeetingEventResponsePropertiesHsEngagementSource = "PROPERTY_DEFAULT_VALUE"
	ExternalCalendarMeetingEventResponsePropertiesHsEngagementSourcePropertyRestore             ExternalCalendarMeetingEventResponsePropertiesHsEngagementSource = "PROPERTY_RESTORE"
	ExternalCalendarMeetingEventResponsePropertiesHsEngagementSourcePropertySettings            ExternalCalendarMeetingEventResponsePropertiesHsEngagementSource = "PROPERTY_SETTINGS"
	ExternalCalendarMeetingEventResponsePropertiesHsEngagementSourceProspectingAgent            ExternalCalendarMeetingEventResponsePropertiesHsEngagementSource = "PROSPECTING_AGENT"
	ExternalCalendarMeetingEventResponsePropertiesHsEngagementSourceQuotas                      ExternalCalendarMeetingEventResponsePropertiesHsEngagementSource = "QUOTAS"
	ExternalCalendarMeetingEventResponsePropertiesHsEngagementSourceQuotes                      ExternalCalendarMeetingEventResponsePropertiesHsEngagementSource = "QUOTES"
	ExternalCalendarMeetingEventResponsePropertiesHsEngagementSourceRecyclingBin                ExternalCalendarMeetingEventResponsePropertiesHsEngagementSource = "RECYCLING_BIN"
	ExternalCalendarMeetingEventResponsePropertiesHsEngagementSourceRestoreObjects              ExternalCalendarMeetingEventResponsePropertiesHsEngagementSource = "RESTORE_OBJECTS"
	ExternalCalendarMeetingEventResponsePropertiesHsEngagementSourceRevenuePlatform             ExternalCalendarMeetingEventResponsePropertiesHsEngagementSource = "REVENUE_PLATFORM"
	ExternalCalendarMeetingEventResponsePropertiesHsEngagementSourceSales                       ExternalCalendarMeetingEventResponsePropertiesHsEngagementSource = "SALES"
	ExternalCalendarMeetingEventResponsePropertiesHsEngagementSourceSalesMessages               ExternalCalendarMeetingEventResponsePropertiesHsEngagementSource = "SALES_MESSAGES"
	ExternalCalendarMeetingEventResponsePropertiesHsEngagementSourceSalesforce                  ExternalCalendarMeetingEventResponsePropertiesHsEngagementSource = "SALESFORCE"
	ExternalCalendarMeetingEventResponsePropertiesHsEngagementSourceSequences                   ExternalCalendarMeetingEventResponsePropertiesHsEngagementSource = "SEQUENCES"
	ExternalCalendarMeetingEventResponsePropertiesHsEngagementSourceSettings                    ExternalCalendarMeetingEventResponsePropertiesHsEngagementSource = "SETTINGS"
	ExternalCalendarMeetingEventResponsePropertiesHsEngagementSourceSidekick                    ExternalCalendarMeetingEventResponsePropertiesHsEngagementSource = "SIDEKICK"
	ExternalCalendarMeetingEventResponsePropertiesHsEngagementSourceSignals                     ExternalCalendarMeetingEventResponsePropertiesHsEngagementSource = "SIGNALS"
	ExternalCalendarMeetingEventResponsePropertiesHsEngagementSourceSlackIntegration            ExternalCalendarMeetingEventResponsePropertiesHsEngagementSource = "SLACK_INTEGRATION"
	ExternalCalendarMeetingEventResponsePropertiesHsEngagementSourceSmartDataCapture            ExternalCalendarMeetingEventResponsePropertiesHsEngagementSource = "SMART_DATA_CAPTURE"
	ExternalCalendarMeetingEventResponsePropertiesHsEngagementSourceSocial                      ExternalCalendarMeetingEventResponsePropertiesHsEngagementSource = "SOCIAL"
	ExternalCalendarMeetingEventResponsePropertiesHsEngagementSourceSuccess                     ExternalCalendarMeetingEventResponsePropertiesHsEngagementSource = "SUCCESS"
	ExternalCalendarMeetingEventResponsePropertiesHsEngagementSourceTally                       ExternalCalendarMeetingEventResponsePropertiesHsEngagementSource = "TALLY"
	ExternalCalendarMeetingEventResponsePropertiesHsEngagementSourceTask                        ExternalCalendarMeetingEventResponsePropertiesHsEngagementSource = "TASK"
	ExternalCalendarMeetingEventResponsePropertiesHsEngagementSourceUnknown                     ExternalCalendarMeetingEventResponsePropertiesHsEngagementSource = "UNKNOWN"
	ExternalCalendarMeetingEventResponsePropertiesHsEngagementSourceWalIncremental              ExternalCalendarMeetingEventResponsePropertiesHsEngagementSource = "WAL_INCREMENTAL"
	ExternalCalendarMeetingEventResponsePropertiesHsEngagementSourceWorkUi                      ExternalCalendarMeetingEventResponsePropertiesHsEngagementSource = "WORK_UI"
	ExternalCalendarMeetingEventResponsePropertiesHsEngagementSourceWorkflowContactDeleteAction ExternalCalendarMeetingEventResponsePropertiesHsEngagementSource = "WORKFLOW_CONTACT_DELETE_ACTION"
	ExternalCalendarMeetingEventResponsePropertiesHsEngagementSourceWorkflows                   ExternalCalendarMeetingEventResponsePropertiesHsEngagementSource = "WORKFLOWS"
)

// The type of location for the meeting. Acceptable values are: ADDRESS, CUSTOM,
// PHONE.
type ExternalCalendarMeetingEventResponsePropertiesHsMeetingLocationType string

const (
	ExternalCalendarMeetingEventResponsePropertiesHsMeetingLocationTypeAddress ExternalCalendarMeetingEventResponsePropertiesHsMeetingLocationType = "ADDRESS"
	ExternalCalendarMeetingEventResponsePropertiesHsMeetingLocationTypeCustom  ExternalCalendarMeetingEventResponsePropertiesHsMeetingLocationType = "CUSTOM"
	ExternalCalendarMeetingEventResponsePropertiesHsMeetingLocationTypePhone   ExternalCalendarMeetingEventResponsePropertiesHsMeetingLocationType = "PHONE"
)

type ExternalCalenderMeetingEventResponse struct {
	// The unique identifier for the meeting event.
	ID string `json:"id" api:"required"`
	// The date and time when the meeting event was initially created, in ISO 8601
	// format.
	CreatedAt time.Time `json:"createdAt" api:"required" format:"date-time"`
	// The date and time when the meeting event was last updated, in ISO 8601 format.
	LastUpdatedAt time.Time                                      `json:"lastUpdatedAt" api:"required" format:"date-time"`
	Properties    ExternalCalendarMeetingEventResponseProperties `json:"properties" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		CreatedAt     respjson.Field
		LastUpdatedAt respjson.Field
		Properties    respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExternalCalenderMeetingEventResponse) RawJSON() string { return r.JSON.raw }
func (r *ExternalCalenderMeetingEventResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExternalClosedRange struct {
	// The end value of the closed range, represented as an integer.
	End int64 `json:"end" api:"required"`
	// The start value of the closed range, represented as an integer.
	Start int64 `json:"start" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		End         respjson.Field
		Start       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExternalClosedRange) RawJSON() string { return r.JSON.raw }
func (r *ExternalClosedRange) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExternalCommunicationConsentCheckbox struct {
	// The ID of the communication consent form being recorded.
	CommunicationTypeID string `json:"communicationTypeId" api:"required"`
	// The text label describing the consent being given.
	Label string `json:"label" api:"required"`
	// Whether the consent checkbox is required.
	Required bool `json:"required" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CommunicationTypeID respjson.Field
		Label               respjson.Field
		Required            respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExternalCommunicationConsentCheckbox) RawJSON() string { return r.JSON.raw }
func (r *ExternalCommunicationConsentCheckbox) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Reminders, ShouldIncludeInviteDescription are required.
type ExternalEmailReminderScheduleParam struct {
	Reminders []ExternalReminderParam `json:"reminders,omitzero" api:"required"`
	// Whether the invite description should be included in the reminder.
	ShouldIncludeInviteDescription bool `json:"shouldIncludeInviteDescription" api:"required"`
	paramObj
}

func (r ExternalEmailReminderScheduleParam) MarshalJSON() (data []byte, err error) {
	type shadow ExternalEmailReminderScheduleParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExternalEmailReminderScheduleParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExternalGuestSettings struct {
	// Indicates whether guests can be added to the meeting.
	CanAddGuests bool `json:"canAddGuests" api:"required"`
	// The maximum number of guests that can be added to the meeting.
	MaxGuestCount int64 `json:"maxGuestCount" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CanAddGuests  respjson.Field
		MaxGuestCount respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExternalGuestSettings) RawJSON() string { return r.JSON.raw }
func (r *ExternalGuestSettings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExternalLegalConsentOptions struct {
	CommunicationConsentCheckboxes []ExternalCommunicationConsentCheckbox `json:"communicationConsentCheckboxes" api:"required"`
	// The text that describes the consent for communication preferences.
	CommunicationConsentText string `json:"communicationConsentText" api:"required"`
	// Whether the legal basis for processing is legitimate interest.
	IsLegitimateInterest                bool    `json:"isLegitimateInterest" api:"required"`
	LegitimateInterestSubscriptionTypes []int64 `json:"legitimateInterestSubscriptionTypes" api:"required"`
	// The text that describes the data processing privacy policy.
	PrivacyPolicyText string `json:"privacyPolicyText" api:"required"`
	// The label for the checkbox used to obtain consent for data processing.
	ProcessingConsentCheckboxLabel string `json:"processingConsentCheckboxLabel" api:"required"`
	// The footer text accompanying the consent for data processing. This field is not
	// used by the meeting platform and will always be empty.
	ProcessingConsentFooterText string `json:"processingConsentFooterText" api:"required"`
	// The text that describes the consent for processing personal data.
	ProcessingConsentText string `json:"processingConsentText" api:"required"`
	// The type of consent required for processing. Accepted values are: IMPLICIT,
	// REQUIRED_CHECKBOX.
	//
	// Any of "IMPLICIT", "REQUIRED_CHECKBOX".
	ProcessingConsentType ExternalLegalConsentOptionsProcessingConsentType `json:"processingConsentType" api:"required"`
	// The legal basis for processing under legitimate interest. Accepted values are:
	// LEGITIMATE_INTEREST_PQL, LEGITIMATE_INTEREST_CLIENT, PERFORMANCE_OF_CONTRACT,
	// CONSENT_WITH_NOTICE, NON_GDPR, PROCESS_AND_STORE, LEGITIMATE_INTEREST_OTHER.
	//
	// Any of "CONSENT_WITH_NOTICE", "LEGITIMATE_INTEREST_CLIENT",
	// "LEGITIMATE_INTEREST_OTHER", "LEGITIMATE_INTEREST_PQL", "NON_GDPR",
	// "PERFORMANCE_OF_CONTRACT", "PROCESS_AND_STORE".
	LegitimateInterestLegalBasis ExternalLegalConsentOptionsLegitimateInterestLegalBasis `json:"legitimateInterestLegalBasis"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CommunicationConsentCheckboxes      respjson.Field
		CommunicationConsentText            respjson.Field
		IsLegitimateInterest                respjson.Field
		LegitimateInterestSubscriptionTypes respjson.Field
		PrivacyPolicyText                   respjson.Field
		ProcessingConsentCheckboxLabel      respjson.Field
		ProcessingConsentFooterText         respjson.Field
		ProcessingConsentText               respjson.Field
		ProcessingConsentType               respjson.Field
		LegitimateInterestLegalBasis        respjson.Field
		ExtraFields                         map[string]respjson.Field
		raw                                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExternalLegalConsentOptions) RawJSON() string { return r.JSON.raw }
func (r *ExternalLegalConsentOptions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The type of consent required for processing. Accepted values are: IMPLICIT,
// REQUIRED_CHECKBOX.
type ExternalLegalConsentOptionsProcessingConsentType string

const (
	ExternalLegalConsentOptionsProcessingConsentTypeImplicit         ExternalLegalConsentOptionsProcessingConsentType = "IMPLICIT"
	ExternalLegalConsentOptionsProcessingConsentTypeRequiredCheckbox ExternalLegalConsentOptionsProcessingConsentType = "REQUIRED_CHECKBOX"
)

// The legal basis for processing under legitimate interest. Accepted values are:
// LEGITIMATE_INTEREST_PQL, LEGITIMATE_INTEREST_CLIENT, PERFORMANCE_OF_CONTRACT,
// CONSENT_WITH_NOTICE, NON_GDPR, PROCESS_AND_STORE, LEGITIMATE_INTEREST_OTHER.
type ExternalLegalConsentOptionsLegitimateInterestLegalBasis string

const (
	ExternalLegalConsentOptionsLegitimateInterestLegalBasisConsentWithNotice        ExternalLegalConsentOptionsLegitimateInterestLegalBasis = "CONSENT_WITH_NOTICE"
	ExternalLegalConsentOptionsLegitimateInterestLegalBasisLegitimateInterestClient ExternalLegalConsentOptionsLegitimateInterestLegalBasis = "LEGITIMATE_INTEREST_CLIENT"
	ExternalLegalConsentOptionsLegitimateInterestLegalBasisLegitimateInterestOther  ExternalLegalConsentOptionsLegitimateInterestLegalBasis = "LEGITIMATE_INTEREST_OTHER"
	ExternalLegalConsentOptionsLegitimateInterestLegalBasisLegitimateInterestPql    ExternalLegalConsentOptionsLegitimateInterestLegalBasis = "LEGITIMATE_INTEREST_PQL"
	ExternalLegalConsentOptionsLegitimateInterestLegalBasisNonGdpr                  ExternalLegalConsentOptionsLegitimateInterestLegalBasis = "NON_GDPR"
	ExternalLegalConsentOptionsLegitimateInterestLegalBasisPerformanceOfContract    ExternalLegalConsentOptionsLegitimateInterestLegalBasis = "PERFORMANCE_OF_CONTRACT"
	ExternalLegalConsentOptionsLegitimateInterestLegalBasisProcessAndStore          ExternalLegalConsentOptionsLegitimateInterestLegalBasis = "PROCESS_AND_STORE"
)

type ExternalLegalConsentResponse struct {
	// The ID of communication consent form being recorded.
	CommunicationTypeID string `json:"communicationTypeId" api:"required"`
	// Whether the user has given consent for the specified communication type.
	Consented bool `json:"consented" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CommunicationTypeID respjson.Field
		Consented           respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExternalLegalConsentResponse) RawJSON() string { return r.JSON.raw }
func (r *ExternalLegalConsentResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this ExternalLegalConsentResponse to a
// ExternalLegalConsentResponseParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// ExternalLegalConsentResponseParam.Overrides()
func (r ExternalLegalConsentResponse) ToParam() ExternalLegalConsentResponseParam {
	return param.Override[ExternalLegalConsentResponseParam](json.RawMessage(r.RawJSON()))
}

// The properties CommunicationTypeID, Consented are required.
type ExternalLegalConsentResponseParam struct {
	// The ID of communication consent form being recorded.
	CommunicationTypeID string `json:"communicationTypeId" api:"required"`
	// Whether the user has given consent for the specified communication type.
	Consented bool `json:"consented" api:"required"`
	paramObj
}

func (r ExternalLegalConsentResponseParam) MarshalJSON() (data []byte, err error) {
	type shadow ExternalLegalConsentResponseParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExternalLegalConsentResponseParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExternalLinkAvailability struct {
	// Whether there are more available meeting times beyond the returned set.
	HasMore bool `json:"hasMore" api:"required"`
	// Available booking slots for the meeting, grouped by the duration.
	LinkAvailabilityByDuration map[string]ExternalLinkAvailabilityForDuration `json:"linkAvailabilityByDuration" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		HasMore                    respjson.Field
		LinkAvailabilityByDuration respjson.Field
		ExtraFields                map[string]respjson.Field
		raw                        string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExternalLinkAvailability) RawJSON() string { return r.JSON.raw }
func (r *ExternalLinkAvailability) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExternalLinkAvailabilityAndBusyTimes struct {
	AllUsersBusyTimes []ExternalUserBusyTimes  `json:"allUsersBusyTimes" api:"required"`
	LinkAvailability  ExternalLinkAvailability `json:"linkAvailability"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AllUsersBusyTimes respjson.Field
		LinkAvailability  respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExternalLinkAvailabilityAndBusyTimes) RawJSON() string { return r.JSON.raw }
func (r *ExternalLinkAvailabilityAndBusyTimes) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExternalLinkAvailabilityForDuration struct {
	Availabilities []ExternalMeetingAvailability `json:"availabilities" api:"required"`
	// The duration of the meeting in milliseconds.
	MeetingDurationMillis int64 `json:"meetingDurationMillis" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Availabilities        respjson.Field
		MeetingDurationMillis respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExternalLinkAvailabilityForDuration) RawJSON() string { return r.JSON.raw }
func (r *ExternalLinkAvailabilityForDuration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExternalLinkDisplayInfo struct {
	// The URL of the user's custom uploaded avatar image.
	Avatar string `json:"avatar"`
	// The URL of the company's avatar image.
	CompanyAvatar string `json:"companyAvatar"`
	// Deprecated field with no impact of link display info.
	Headline string `json:"headline"`
	// Option for determining which avatar to display on scheduling page. Accepted
	// values are: PROFILE_IMAGE, COMPANY_LOGO, CUSTOM_AVATAR,
	//
	// Any of "COMPANY_LOGO", "CUSTOM_AVATAR", "PROFILE_IMAGE".
	PublicDisplayAvatarOption ExternalLinkDisplayInfoPublicDisplayAvatarOption `json:"publicDisplayAvatarOption"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Avatar                    respjson.Field
		CompanyAvatar             respjson.Field
		Headline                  respjson.Field
		PublicDisplayAvatarOption respjson.Field
		ExtraFields               map[string]respjson.Field
		raw                       string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExternalLinkDisplayInfo) RawJSON() string { return r.JSON.raw }
func (r *ExternalLinkDisplayInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Option for determining which avatar to display on scheduling page. Accepted
// values are: PROFILE_IMAGE, COMPANY_LOGO, CUSTOM_AVATAR,
type ExternalLinkDisplayInfoPublicDisplayAvatarOption string

const (
	ExternalLinkDisplayInfoPublicDisplayAvatarOptionCompanyLogo  ExternalLinkDisplayInfoPublicDisplayAvatarOption = "COMPANY_LOGO"
	ExternalLinkDisplayInfoPublicDisplayAvatarOptionCustomAvatar ExternalLinkDisplayInfoPublicDisplayAvatarOption = "CUSTOM_AVATAR"
	ExternalLinkDisplayInfoPublicDisplayAvatarOptionProfileImage ExternalLinkDisplayInfoPublicDisplayAvatarOption = "PROFILE_IMAGE"
)

type ExternalLinkFormField struct {
	// The specific field type of the form field. Corresponds to property types (e.g.,
	// `select`, `radio`, `date`, etc)
	FieldType string `json:"fieldType" api:"required"`
	// Whether the form field is a custom field.
	IsCustom bool `json:"isCustom" api:"required"`
	// Whether the form field is mandatory.
	IsRequired bool `json:"isRequired" api:"required"`
	// The text label for the form field.
	Label string `json:"label" api:"required"`
	// The name identifier for the form field.
	Name    string           `json:"name" api:"required"`
	Options []ExternalOption `json:"options" api:"required"`
	// The data type of the form field accepts (e.g. `date`, `enumeration`, etc)
	Type string `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FieldType   respjson.Field
		IsCustom    respjson.Field
		IsRequired  respjson.Field
		Label       respjson.Field
		Name        respjson.Field
		Options     respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExternalLinkFormField) RawJSON() string { return r.JSON.raw }
func (r *ExternalLinkFormField) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExternalLinkMetadata struct {
	// The unique identifier for the meeting link.
	ID string `json:"id" api:"required"`
	// The Unix time in milliseconds when the meeting link was created.
	CreatedAt time.Time `json:"createdAt" api:"required" format:"date-time"`
	// Whether the meeting link is the user's default link.
	DefaultLink bool `json:"defaultLink" api:"required"`
	// The URL of the meeting link.
	Link string `json:"link" api:"required"`
	// The user ID of the meeting link's organizer.
	OrganizerUserID string `json:"organizerUserId" api:"required"`
	// The slug of the meeting link, located directly after the domain in the URL.
	Slug string `json:"slug" api:"required"`
	// The type of the external meeting link. Accepted values are: PERSONAL_LINK,
	// GROUP_CALENDAR, ROUND_ROBIN_CALENDAR.
	//
	// Any of "GROUP_CALENDAR", "PERSONAL_LINK", "ROUND_ROBIN_CALENDAR".
	Type                 ExternalLinkMetadataType `json:"type" api:"required"`
	UserIDsOfLinkMembers []string                 `json:"userIdsOfLinkMembers" api:"required"`
	// The name of the meeting link.
	Name string `json:"name"`
	// The Unix time in milliseconds when the meeting link was last updated.
	UpdatedAt time.Time `json:"updatedAt" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                   respjson.Field
		CreatedAt            respjson.Field
		DefaultLink          respjson.Field
		Link                 respjson.Field
		OrganizerUserID      respjson.Field
		Slug                 respjson.Field
		Type                 respjson.Field
		UserIDsOfLinkMembers respjson.Field
		Name                 respjson.Field
		UpdatedAt            respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExternalLinkMetadata) RawJSON() string { return r.JSON.raw }
func (r *ExternalLinkMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The type of the external meeting link. Accepted values are: PERSONAL_LINK,
// GROUP_CALENDAR, ROUND_ROBIN_CALENDAR.
type ExternalLinkMetadataType string

const (
	ExternalLinkMetadataTypeGroupCalendar      ExternalLinkMetadataType = "GROUP_CALENDAR"
	ExternalLinkMetadataTypePersonalLink       ExternalLinkMetadataType = "PERSONAL_LINK"
	ExternalLinkMetadataTypeRoundRobinCalendar ExternalLinkMetadataType = "ROUND_ROBIN_CALENDAR"
)

type ExternalMeetingAvailability struct {
	// The end time of the meeting availability, represented as Unix time in
	// milliseconds.
	EndMillisUtc int64 `json:"endMillisUtc" api:"required"`
	// The start time of the meeting availability, represented as Unix time in
	// milliseconds.
	StartMillisUtc int64 `json:"startMillisUtc" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		EndMillisUtc   respjson.Field
		StartMillisUtc respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExternalMeetingAvailability) RawJSON() string { return r.JSON.raw }
func (r *ExternalMeetingAvailability) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Duration, Email, FirstName, FormFields, LastName,
// LegalConsentResponses, LikelyAvailableUserIDs, Slug, StartTime are required.
type ExternalMeetingBookingParam struct {
	// The duration of the meeting in milliseconds.
	Duration int64 `json:"duration" api:"required"`
	// The email address of the person booking the meeting.
	Email string `json:"email" api:"required"`
	// The first name of the person booking the meeting.
	FirstName  string                          `json:"firstName" api:"required"`
	FormFields []ExternalBookingFormFieldParam `json:"formFields,omitzero" api:"required"`
	// The last name of the person booking the meeting.
	LastName               string                              `json:"lastName" api:"required"`
	LegalConsentResponses  []ExternalLegalConsentResponseParam `json:"legalConsentResponses,omitzero" api:"required"`
	LikelyAvailableUserIDs []string                            `json:"likelyAvailableUserIds,omitzero" api:"required"`
	// The unique path identifier for the meeting page.
	Slug string `json:"slug" api:"required"`
	// The date and time when the meeting is scheduled to start, in ISO 8601 format.
	StartTime time.Time `json:"startTime" api:"required" format:"date-time"`
	// The locale used for formatting dates and times in the meeting booking.
	Locale param.Opt[string] `json:"locale,omitzero"`
	// The timezone in which the meeting is scheduled.
	Timezone param.Opt[string] `json:"timezone,omitzero"`
	paramObj
}

func (r ExternalMeetingBookingParam) MarshalJSON() (data []byte, err error) {
	type shadow ExternalMeetingBookingParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExternalMeetingBookingParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExternalMeetingBookingResponse struct {
	// The timezone the meeting was booked from.
	BookingTimezone string `json:"bookingTimezone" api:"required"`
	// The unique identifier for the meeting's calendar event.
	CalendarEventID string `json:"calendarEventId" api:"required"`
	// The ID of the contact associated to the meeting.
	ContactID string `json:"contactId" api:"required"`
	// The duration of the meeting in milliseconds.
	Duration int64 `json:"duration" api:"required"`
	// The date and time when the meeting is scheduled to end, in ISO 8601 format.
	End         time.Time                    `json:"end" api:"required" format:"date-time"`
	FormFields  []ExternalValidatedFormField `json:"formFields" api:"required"`
	GuestEmails []string                     `json:"guestEmails" api:"required"`
	// Whether the meeting was booked offline and no associated calendar event was
	// created.
	IsOffline             bool                           `json:"isOffline" api:"required"`
	LegalConsentResponses []ExternalLegalConsentResponse `json:"legalConsentResponses" api:"required"`
	// The date and time when the meeting is scheduled to start, in ISO 8601 format.
	Start time.Time `json:"start" api:"required" format:"date-time"`
	// The title of the meeting and calendar event.
	Subject string `json:"subject" api:"required"`
	// The locale the meeting was booked with, used to determine date formatting in
	// calendar event description.
	Locale string `json:"locale"`
	// The physical or virtual location where the meeting will take place.
	Location string `json:"location"`
	// The unique identifier for the web conference meeting.
	WebConferenceMeetingID string `json:"webConferenceMeetingId"`
	// The URL for accessing the meeting's web conference.
	WebConferenceURL string `json:"webConferenceUrl"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BookingTimezone        respjson.Field
		CalendarEventID        respjson.Field
		ContactID              respjson.Field
		Duration               respjson.Field
		End                    respjson.Field
		FormFields             respjson.Field
		GuestEmails            respjson.Field
		IsOffline              respjson.Field
		LegalConsentResponses  respjson.Field
		Start                  respjson.Field
		Subject                respjson.Field
		Locale                 respjson.Field
		Location               respjson.Field
		WebConferenceMeetingID respjson.Field
		WebConferenceURL       respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExternalMeetingBookingResponse) RawJSON() string { return r.JSON.raw }
func (r *ExternalMeetingBookingResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExternalMeetingsLinkSettings struct {
	// An array containing the closed range availability for a meeting link. Closed
	// range times are provided as minute offsets from midnight (e.g., 540 corresponds
	// to 9am).
	Availability map[string]ExternalClosedRange `json:"availability" api:"required"`
	Durations    []int64                        `json:"durations" api:"required"`
	FormFields   []ExternalLinkFormField        `json:"formFields" api:"required"`
	// Whether the legal consent checkbox is displayed during meeting booking.
	LegalConsentEnabled bool `json:"legalConsentEnabled" api:"required"`
	// The minimum buffer time in milliseconds between consecutive meetings.
	MeetingBufferTime int64 `json:"meetingBufferTime" api:"required"`
	// Indicates whether the meeting owner is prioritized during booking. Only applies
	// to link types of ROUND_ROBIN.
	OwnerPrioritized bool `json:"ownerPrioritized" api:"required"`
	// The increment for available start times of meetings, spelt out as a word (e.g.
	// 15 minute increment corresponds to `FIFTEEN`). `MEETING_DURATION` is also a
	// valid value.
	//
	// Any of "FIFTEEN", "FIVE", "FORTY_FIVE", "MEETING_DURATION", "NINETY",
	// "ONE_HUNDRED_TWENTY", "SIXTY", "TEN", "THIRTY", "TWENTY".
	StartTimeIncrementMinutes ExternalMeetingsLinkSettingsStartTimeIncrementMinutes `json:"startTimeIncrementMinutes" api:"required"`
	// Legacy property that indicates the number of weeks in advance that availability
	// is advertised. May be outdated or superseded by other properties.
	WeeksToAdvertise int64 `json:"weeksToAdvertise" api:"required"`
	// The end date for a meeting link's custom availability window, represented as
	// Unix time in milliseconds.
	CustomAvailabilityEndDate int64 `json:"customAvailabilityEndDate"`
	// The start date for a meeting link's custom availability window, represented as
	// Unix time in milliseconds.
	CustomAvailabilityStartDate int64                   `json:"customAvailabilityStartDate"`
	DisplayInfo                 ExternalLinkDisplayInfo `json:"displayInfo"`
	GuestSettings               ExternalGuestSettings   `json:"guestSettings"`
	// The language setting used for the meeting link.
	Language            string                      `json:"language"`
	LegalConsentOptions ExternalLegalConsentOptions `json:"legalConsentOptions"`
	// The locale setting used for formatting dates and times in the meeting link.
	Locale string `json:"locale"`
	// The physical or virtual location where the meeting will take place.
	Location string `json:"location"`
	// The URL to redirect to after a meeting is booked.
	RedirectURL       string                            `json:"redirectUrl"`
	WelcomeScreenInfo ExternalMeetingsWelcomeScreenInfo `json:"welcomeScreenInfo"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Availability                respjson.Field
		Durations                   respjson.Field
		FormFields                  respjson.Field
		LegalConsentEnabled         respjson.Field
		MeetingBufferTime           respjson.Field
		OwnerPrioritized            respjson.Field
		StartTimeIncrementMinutes   respjson.Field
		WeeksToAdvertise            respjson.Field
		CustomAvailabilityEndDate   respjson.Field
		CustomAvailabilityStartDate respjson.Field
		DisplayInfo                 respjson.Field
		GuestSettings               respjson.Field
		Language                    respjson.Field
		LegalConsentOptions         respjson.Field
		Locale                      respjson.Field
		Location                    respjson.Field
		RedirectURL                 respjson.Field
		WelcomeScreenInfo           respjson.Field
		ExtraFields                 map[string]respjson.Field
		raw                         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExternalMeetingsLinkSettings) RawJSON() string { return r.JSON.raw }
func (r *ExternalMeetingsLinkSettings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The increment for available start times of meetings, spelt out as a word (e.g.
// 15 minute increment corresponds to `FIFTEEN`). `MEETING_DURATION` is also a
// valid value.
type ExternalMeetingsLinkSettingsStartTimeIncrementMinutes string

const (
	ExternalMeetingsLinkSettingsStartTimeIncrementMinutesFifteen          ExternalMeetingsLinkSettingsStartTimeIncrementMinutes = "FIFTEEN"
	ExternalMeetingsLinkSettingsStartTimeIncrementMinutesFive             ExternalMeetingsLinkSettingsStartTimeIncrementMinutes = "FIVE"
	ExternalMeetingsLinkSettingsStartTimeIncrementMinutesFortyFive        ExternalMeetingsLinkSettingsStartTimeIncrementMinutes = "FORTY_FIVE"
	ExternalMeetingsLinkSettingsStartTimeIncrementMinutesMeetingDuration  ExternalMeetingsLinkSettingsStartTimeIncrementMinutes = "MEETING_DURATION"
	ExternalMeetingsLinkSettingsStartTimeIncrementMinutesNinety           ExternalMeetingsLinkSettingsStartTimeIncrementMinutes = "NINETY"
	ExternalMeetingsLinkSettingsStartTimeIncrementMinutesOneHundredTwenty ExternalMeetingsLinkSettingsStartTimeIncrementMinutes = "ONE_HUNDRED_TWENTY"
	ExternalMeetingsLinkSettingsStartTimeIncrementMinutesSixty            ExternalMeetingsLinkSettingsStartTimeIncrementMinutes = "SIXTY"
	ExternalMeetingsLinkSettingsStartTimeIncrementMinutesTen              ExternalMeetingsLinkSettingsStartTimeIncrementMinutes = "TEN"
	ExternalMeetingsLinkSettingsStartTimeIncrementMinutesThirty           ExternalMeetingsLinkSettingsStartTimeIncrementMinutes = "THIRTY"
	ExternalMeetingsLinkSettingsStartTimeIncrementMinutesTwenty           ExternalMeetingsLinkSettingsStartTimeIncrementMinutes = "TWENTY"
)

type ExternalMeetingsUser struct {
	// The ID for the meetings user. This value is different than the userId.
	ID string `json:"id" api:"required"`
	// The calendar provider associated with the user. Accepted values are: GOOGLE,
	// OFFICE365, EXCHANGE, UNKNOWN.
	//
	// Any of "EXCHANGE", "GOOGLE", "OFFICE365", "UNKNOWN".
	CalendarProvider ExternalMeetingsUserCalendarProvider `json:"calendarProvider" api:"required"`
	// Whether the user has a sales starter seat.
	IsSalesStarter bool `json:"isSalesStarter" api:"required"`
	// The ID of the user.
	UserID      string              `json:"userId" api:"required"`
	UserProfile ExternalUserProfile `json:"userProfile" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		CalendarProvider respjson.Field
		IsSalesStarter   respjson.Field
		UserID           respjson.Field
		UserProfile      respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExternalMeetingsUser) RawJSON() string { return r.JSON.raw }
func (r *ExternalMeetingsUser) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The calendar provider associated with the user. Accepted values are: GOOGLE,
// OFFICE365, EXCHANGE, UNKNOWN.
type ExternalMeetingsUserCalendarProvider string

const (
	ExternalMeetingsUserCalendarProviderExchange  ExternalMeetingsUserCalendarProvider = "EXCHANGE"
	ExternalMeetingsUserCalendarProviderGoogle    ExternalMeetingsUserCalendarProvider = "GOOGLE"
	ExternalMeetingsUserCalendarProviderOffice365 ExternalMeetingsUserCalendarProvider = "OFFICE365"
	ExternalMeetingsUserCalendarProviderUnknown   ExternalMeetingsUserCalendarProvider = "UNKNOWN"
)

type ExternalMeetingsWelcomeScreenInfo struct {
	// A brief description displayed the welcome screen below the title.
	Description string `json:"description"`
	// The URL of the logo image to be displayed on the welcome screen, only used if
	// `useCompanyLogo` is false.
	LogoURL string `json:"logoUrl"`
	// Deprecated property. Value can be ignored but will always be false.
	ShowWelcomeScreen bool `json:"showWelcomeScreen"`
	// The main heading displayed on the welcome screen.
	Title string `json:"title"`
	// Whether the company's logo should be displayed on the welcome screen.
	UseCompanyLogo bool `json:"useCompanyLogo"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Description       respjson.Field
		LogoURL           respjson.Field
		ShowWelcomeScreen respjson.Field
		Title             respjson.Field
		UseCompanyLogo    respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExternalMeetingsWelcomeScreenInfo) RawJSON() string { return r.JSON.raw }
func (r *ExternalMeetingsWelcomeScreenInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExternalOption struct {
	// A brief description of the option.
	Description string `json:"description" api:"required"`
	// The order in which the option should be displayed.
	DisplayOrder int64 `json:"displayOrder" api:"required"`
	// Deprecated property. Will always be 0.
	DoubleData float64 `json:"doubleData" api:"required"`
	// Whether the option should be hidden from the user.
	Hidden bool `json:"hidden" api:"required"`
	// The text label for the option.
	Label string `json:"label" api:"required"`
	// Whether the option is read-only.
	ReadOnly bool `json:"readOnly" api:"required"`
	// The value associated with the option.
	Value string `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Description  respjson.Field
		DisplayOrder respjson.Field
		DoubleData   respjson.Field
		Hidden       respjson.Field
		Label        respjson.Field
		ReadOnly     respjson.Field
		Value        respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExternalOption) RawJSON() string { return r.JSON.raw }
func (r *ExternalOption) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties NumberOfTimeUnits, TimeUnit are required.
type ExternalReminderParam struct {
	// The number of timeUnits prior to the meeting start when the reminder will be
	// sent.
	NumberOfTimeUnits int64 `json:"numberOfTimeUnits" api:"required"`
	// Accepted values are: WEEKS, DAYS, HOURS, MINUTES.
	//
	// Any of "DAYS", "HOURS", "MINUTES", "WEEKS".
	TimeUnit ExternalReminderTimeUnit `json:"timeUnit,omitzero" api:"required"`
	paramObj
}

func (r ExternalReminderParam) MarshalJSON() (data []byte, err error) {
	type shadow ExternalReminderParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExternalReminderParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Accepted values are: WEEKS, DAYS, HOURS, MINUTES.
type ExternalReminderTimeUnit string

const (
	ExternalReminderTimeUnitDays    ExternalReminderTimeUnit = "DAYS"
	ExternalReminderTimeUnitHours   ExternalReminderTimeUnit = "HOURS"
	ExternalReminderTimeUnitMinutes ExternalReminderTimeUnit = "MINUTES"
	ExternalReminderTimeUnitWeeks   ExternalReminderTimeUnit = "WEEKS"
)

type ExternalTimeRange struct {
	// The end time of the time range, represented as Unix time in milliseconds.
	End int64 `json:"end" api:"required"`
	// The start time of the time range, represented as Unix time in milliseconds.
	Start int64 `json:"start" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		End         respjson.Field
		Start       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExternalTimeRange) RawJSON() string { return r.JSON.raw }
func (r *ExternalTimeRange) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExternalUserBusyTimes struct {
	BusyTimes []ExternalTimeRange `json:"busyTimes" api:"required"`
	// Whether the user is offline.
	IsOffline    bool                 `json:"isOffline" api:"required"`
	MeetingsUser ExternalMeetingsUser `json:"meetingsUser" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BusyTimes    respjson.Field
		IsOffline    respjson.Field
		MeetingsUser respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExternalUserBusyTimes) RawJSON() string { return r.JSON.raw }
func (r *ExternalUserBusyTimes) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExternalUserProfile struct {
	// The email address of the user.
	Email string `json:"email" api:"required"`
	// The first name of the user.
	FirstName string `json:"firstName"`
	// The full name of the user.
	FullName string `json:"fullName"`
	// The last name of the user.
	LastName string `json:"lastName"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Email       respjson.Field
		FirstName   respjson.Field
		FullName    respjson.Field
		LastName    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExternalUserProfile) RawJSON() string { return r.JSON.raw }
func (r *ExternalUserProfile) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExternalValidatedFormField struct {
	// Whether the form field is a custom field.
	IsCustom bool `json:"isCustom" api:"required"`
	// The text label associated with the form field.
	Label string `json:"label" api:"required"`
	// The name identifier for the form field, includes underscores in place of spaces
	// (e.g., the label `my form` is converted to `my_form`).
	Name string `json:"name" api:"required"`
	// The value associated with the form field.
	Value string `json:"value" api:"required"`
	// The specific input type of the form field. Corresponds to property types (e.g.,
	// `select`, `radio`, `date`, etc).
	FieldType string `json:"fieldType"`
	// The translated text label for the form field.
	TranslatedLabel string `json:"translatedLabel"`
	// The text label associated to a form field selection or option.
	ValueLabel string `json:"valueLabel"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		IsCustom        respjson.Field
		Label           respjson.Field
		Name            respjson.Field
		Value           respjson.Field
		FieldType       respjson.Field
		TranslatedLabel respjson.Field
		ValueLabel      respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExternalValidatedFormField) RawJSON() string { return r.JSON.raw }
func (r *ExternalValidatedFormField) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
