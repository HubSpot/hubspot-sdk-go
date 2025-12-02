// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package scheduler

import (
	"encoding/json"
	"time"

	"github.com/stainless-sdks/hubspot-sdk-go/internal/apijson"
	"github.com/stainless-sdks/hubspot-sdk-go/option"
	"github.com/stainless-sdks/hubspot-sdk-go/packages/param"
	"github.com/stainless-sdks/hubspot-sdk-go/packages/respjson"
	"github.com/stainless-sdks/hubspot-sdk-go/shared"
)

// MeetingService contains methods and other services that help with interacting
// with the hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMeetingService] method instead.
type MeetingService struct {
	Options       []option.RequestOption
	Calendar      MeetingCalendarService
	MeetingsLinks MeetingMeetingsLinkService
}

// NewMeetingService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewMeetingService(opts ...option.RequestOption) (r MeetingService) {
	r = MeetingService{}
	r.Options = opts
	r.Calendar = NewMeetingCalendarService(opts...)
	r.MeetingsLinks = NewMeetingMeetingsLinkService(opts...)
	return
}

type CollectionResponseWithTotalExternalLinkMetadataForwardPaging struct {
	Results []ExternalLinkMetadata `json:"results,required"`
	Total   int64                  `json:"total,required"`
	Paging  shared.ForwardPaging   `json:"paging"`
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
func (r CollectionResponseWithTotalExternalLinkMetadataForwardPaging) RawJSON() string {
	return r.JSON.raw
}
func (r *CollectionResponseWithTotalExternalLinkMetadataForwardPaging) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties To, Types are required.
type ExternalAssociationCreateRequestParam struct {
	To    shared.PublicObjectIDParam    `json:"to,omitzero,required"`
	Types []shared.AssociationSpecParam `json:"types,omitzero,required"`
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
	Name  string `json:"name,required"`
	Value string `json:"value,required"`
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
	AllUsersBusyTimes []ExternalUserBusyTimes      `json:"allUsersBusyTimes,required"`
	CustomParams      ExternalMeetingsLinkSettings `json:"customParams,required"`
	IsOffline         bool                         `json:"isOffline,required"`
	LinkID            string                       `json:"linkId,required"`
	// Any of "GROUP_CALENDAR", "PERSONAL_LINK", "ROUND_ROBIN_CALENDAR".
	LinkType         ExternalBookingInfoLinkType `json:"linkType,required"`
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

type ExternalBookingInfoLinkType string

const (
	ExternalBookingInfoLinkTypeGroupCalendar      ExternalBookingInfoLinkType = "GROUP_CALENDAR"
	ExternalBookingInfoLinkTypePersonalLink       ExternalBookingInfoLinkType = "PERSONAL_LINK"
	ExternalBookingInfoLinkTypeRoundRobinCalendar ExternalBookingInfoLinkType = "ROUND_ROBIN_CALENDAR"
)

type ExternalBrandingMetadata struct {
	LogoAltText         string `json:"logoAltText,required"`
	ShowMarketingAd     bool   `json:"showMarketingAd,required"`
	ShowSalesAd         bool   `json:"showSalesAd,required"`
	Accent2Color        string `json:"accent2Color"`
	AccentColor         string `json:"accentColor"`
	CompanyAddressLine1 string `json:"companyAddressLine1"`
	CompanyAddressLine2 string `json:"companyAddressLine2"`
	CompanyAvatar       string `json:"companyAvatar"`
	CompanyCity         string `json:"companyCity"`
	CompanyCountry      string `json:"companyCountry"`
	CompanyDomain       string `json:"companyDomain"`
	CompanyName         string `json:"companyName"`
	CompanyState        string `json:"companyState"`
	CompanyZip          string `json:"companyZip"`
	LogoHeight          int64  `json:"logoHeight"`
	LogoURL             string `json:"logoUrl"`
	LogoWidth           int64  `json:"logoWidth"`
	PrimaryColor        string `json:"primaryColor"`
	SecondaryColor      string `json:"secondaryColor"`
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
// HsMeetingTitle, HsTimestamp, HubspotOwnerID are required.
type ExternalCalendarMeetingEventCreatePropertiesParam struct {
	HsMeetingEndTime       time.Time         `json:"hs_meeting_end_time,required" format:"date-time"`
	HsMeetingOutcome       string            `json:"hs_meeting_outcome,required"`
	HsMeetingStartTime     time.Time         `json:"hs_meeting_start_time,required" format:"date-time"`
	HsMeetingTitle         string            `json:"hs_meeting_title,required"`
	HsTimestamp            time.Time         `json:"hs_timestamp,required" format:"date-time"`
	HubspotOwnerID         string            `json:"hubspot_owner_id,required"`
	HsActivityType         param.Opt[string] `json:"hs_activity_type,omitzero"`
	HsInternalMeetingNotes param.Opt[string] `json:"hs_internal_meeting_notes,omitzero"`
	HsMeetingBody          param.Opt[string] `json:"hs_meeting_body,omitzero"`
	HsMeetingLocation      param.Opt[string] `json:"hs_meeting_location,omitzero"`
	HsMeetingLocationType  param.Opt[string] `json:"hs_meeting_location_type,omitzero"`
	HsAttachmentIDs        []string          `json:"hs_attachment_ids,omitzero"`
	HsAttendeeOwnerIDs     []string          `json:"hs_attendee_owner_ids,omitzero"`
	paramObj
}

func (r ExternalCalendarMeetingEventCreatePropertiesParam) MarshalJSON() (data []byte, err error) {
	type shadow ExternalCalendarMeetingEventCreatePropertiesParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExternalCalendarMeetingEventCreatePropertiesParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Associations, EmailReminderSchedule, Properties, Timezone are
// required.
type ExternalCalendarMeetingEventCreateRequestParam struct {
	Associations          []ExternalAssociationCreateRequestParam           `json:"associations,omitzero,required"`
	EmailReminderSchedule ExternalEmailReminderScheduleParam                `json:"emailReminderSchedule,omitzero,required"`
	Properties            ExternalCalendarMeetingEventCreatePropertiesParam `json:"properties,omitzero,required"`
	Timezone              string                                            `json:"timezone,required"`
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
	// Any of "ACADEMY", "ACCEPTANCE_TEST", "ACTIVITY_AUTO_ASSOCIATE", "ADS",
	// "AI_GROUP", "ANALYTICS", "API", "APPROVALS", "ASSISTS", "ASSOCIATIONS",
	// "AUTO_ASSOCIATE_BY_DOMAIN", "AUTOMATION_JOURNEY", "AUTOMATION_PLATFORM",
	// "AVATARS_SERVICE", "BATCH_UPDATE", "BCC_TO_CRM", "BEHAVIORAL_EVENTS",
	// "BET_ASSIGNMENT", "BET_CRM_CONNECTOR", "BIDEN", "BILLING", "BOT", "CALCULATED",
	// "CENTRAL_EXCHANGE_RATES", "CHATSPOT", "CLONE_OBJECTS", "COMMUNICATOR",
	// "COMPANIES", "COMPANY_FAMILIES", "COMPANY_INSIGHTS", "CONTACTS", "CONTACTS_WEB",
	// "CONTENT_MEMBERSHIP", "CONVERSATIONAL_ENRICHMENT", "CONVERSATIONS",
	// "CRM_PROCESSES_PLATFORM", "CRM_UI", "CRM_UI_BULK_ACTION", "CUSTOMER_AGENT",
	// "DATA_ENRICHMENT", "DATASET", "DEALS", "DEFAULT", "DELETE_OBJECTS", "EMAIL",
	// "EMAIL_INBOX_IMPORT", "EMAIL_INTEGRATION", "ENGAGEMENTS", "EXTENSION",
	// "FILE_MANAGER", "FLYWHEEL_PRODUCT_DATA_SYNC", "FORECASTING", "FORM",
	// "FORWARD_TO_CRM", "GMAIL_INTEGRATION", "GOALS", "HEISENBERG", "HELP_DESK",
	// "HELP_DESK_AI", "IMPORT", "INTEGRATION", "INTEGRATIONS_PLATFORM",
	// "INTEGRATIONS_SYNC", "INTENT", "INTERNAL_PROCESSING", "LEADIN",
	// "LEGAL_BASIS_REMEDIATION", "MARKET_SOURCING", "MARKETPLACE", "MEETINGS",
	// "MERGE_COMPANIES", "MERGE_CONTACTS", "MERGE_OBJECTS", "MICROAPPS", "MIGRATION",
	// "MOBILE_ANDROID", "MOBILE_IOS", "PAYMENTS", "PIPELINE_SETTINGS", "PLAYBOOKS",
	// "PORTAL_OBJECT_SYNC", "PORTAL_USER_ASSOCIATOR", "PRESENTATIONS",
	// "PRIMARY_AUTOMATION", "PROPERTY_RESTORE", "PROPERTY_SETTINGS",
	// "PROSPECTING_AGENT", "QUOTAS", "QUOTES", "RECYCLING_BIN", "RESTORE_OBJECTS",
	// "SALES", "SALES_MESSAGES", "SALESFORCE", "SEQUENCES", "SETTINGS", "SIDEKICK",
	// "SIGNALS", "SLACK_INTEGRATION", "SOCIAL", "SUCCESS", "TALLY", "TASK", "UNKNOWN",
	// "WAL_INCREMENTAL", "WORKFLOW_CONTACT_DELETE_ACTION", "WORKFLOWS".
	HsEngagementSource             ExternalCalendarMeetingEventResponsePropertiesHsEngagementSource `json:"hs_engagement_source,required"`
	HsEngagementSourceID           string                                                           `json:"hs_engagement_source_id,required"`
	HsMeetingEndTime               time.Time                                                        `json:"hs_meeting_end_time,required" format:"date-time"`
	HsMeetingOutcome               string                                                           `json:"hs_meeting_outcome,required"`
	HsMeetingStartTime             time.Time                                                        `json:"hs_meeting_start_time,required" format:"date-time"`
	HsMeetingTitle                 string                                                           `json:"hs_meeting_title,required"`
	HsTimestamp                    time.Time                                                        `json:"hs_timestamp,required" format:"date-time"`
	HsActivityType                 string                                                           `json:"hs_activity_type"`
	HsAttachmentIDs                []string                                                         `json:"hs_attachment_ids"`
	HsAttendeeOwnerIDs             []string                                                         `json:"hs_attendee_owner_ids"`
	HsIncludeDescriptionInReminder string                                                           `json:"hs_include_description_in_reminder"`
	HsInternalMeetingNotes         string                                                           `json:"hs_internal_meeting_notes"`
	HsMeetingBody                  string                                                           `json:"hs_meeting_body"`
	HsMeetingExternalURL           string                                                           `json:"hs_meeting_external_url"`
	HsMeetingLocation              string                                                           `json:"hs_meeting_location"`
	// Any of "ADDRESS", "CUSTOM", "PHONE".
	HsMeetingLocationType ExternalCalendarMeetingEventResponsePropertiesHsMeetingLocationType `json:"hs_meeting_location_type"`
	HsUniqueID            string                                                              `json:"hs_unique_id"`
	HubspotOwnerID        string                                                              `json:"hubspot_owner_id"`
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
		HubspotOwnerID                 respjson.Field
		ExtraFields                    map[string]respjson.Field
		raw                            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExternalCalendarMeetingEventResponseProperties) RawJSON() string { return r.JSON.raw }
func (r *ExternalCalendarMeetingEventResponseProperties) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExternalCalendarMeetingEventResponsePropertiesHsEngagementSource string

const (
	ExternalCalendarMeetingEventResponsePropertiesHsEngagementSourceAcademy                     ExternalCalendarMeetingEventResponsePropertiesHsEngagementSource = "ACADEMY"
	ExternalCalendarMeetingEventResponsePropertiesHsEngagementSourceAcceptanceTest              ExternalCalendarMeetingEventResponsePropertiesHsEngagementSource = "ACCEPTANCE_TEST"
	ExternalCalendarMeetingEventResponsePropertiesHsEngagementSourceActivityAutoAssociate       ExternalCalendarMeetingEventResponsePropertiesHsEngagementSource = "ACTIVITY_AUTO_ASSOCIATE"
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
	ExternalCalendarMeetingEventResponsePropertiesHsEngagementSourceCalculated                  ExternalCalendarMeetingEventResponsePropertiesHsEngagementSource = "CALCULATED"
	ExternalCalendarMeetingEventResponsePropertiesHsEngagementSourceCentralExchangeRates        ExternalCalendarMeetingEventResponsePropertiesHsEngagementSource = "CENTRAL_EXCHANGE_RATES"
	ExternalCalendarMeetingEventResponsePropertiesHsEngagementSourceChatspot                    ExternalCalendarMeetingEventResponsePropertiesHsEngagementSource = "CHATSPOT"
	ExternalCalendarMeetingEventResponsePropertiesHsEngagementSourceCloneObjects                ExternalCalendarMeetingEventResponsePropertiesHsEngagementSource = "CLONE_OBJECTS"
	ExternalCalendarMeetingEventResponsePropertiesHsEngagementSourceCommunicator                ExternalCalendarMeetingEventResponsePropertiesHsEngagementSource = "COMMUNICATOR"
	ExternalCalendarMeetingEventResponsePropertiesHsEngagementSourceCompanies                   ExternalCalendarMeetingEventResponsePropertiesHsEngagementSource = "COMPANIES"
	ExternalCalendarMeetingEventResponsePropertiesHsEngagementSourceCompanyFamilies             ExternalCalendarMeetingEventResponsePropertiesHsEngagementSource = "COMPANY_FAMILIES"
	ExternalCalendarMeetingEventResponsePropertiesHsEngagementSourceCompanyInsights             ExternalCalendarMeetingEventResponsePropertiesHsEngagementSource = "COMPANY_INSIGHTS"
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
	ExternalCalendarMeetingEventResponsePropertiesHsEngagementSourceDataset                     ExternalCalendarMeetingEventResponsePropertiesHsEngagementSource = "DATASET"
	ExternalCalendarMeetingEventResponsePropertiesHsEngagementSourceDeals                       ExternalCalendarMeetingEventResponsePropertiesHsEngagementSource = "DEALS"
	ExternalCalendarMeetingEventResponsePropertiesHsEngagementSourceDefault                     ExternalCalendarMeetingEventResponsePropertiesHsEngagementSource = "DEFAULT"
	ExternalCalendarMeetingEventResponsePropertiesHsEngagementSourceDeleteObjects               ExternalCalendarMeetingEventResponsePropertiesHsEngagementSource = "DELETE_OBJECTS"
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
	ExternalCalendarMeetingEventResponsePropertiesHsEngagementSourceMeetings                    ExternalCalendarMeetingEventResponsePropertiesHsEngagementSource = "MEETINGS"
	ExternalCalendarMeetingEventResponsePropertiesHsEngagementSourceMergeCompanies              ExternalCalendarMeetingEventResponsePropertiesHsEngagementSource = "MERGE_COMPANIES"
	ExternalCalendarMeetingEventResponsePropertiesHsEngagementSourceMergeContacts               ExternalCalendarMeetingEventResponsePropertiesHsEngagementSource = "MERGE_CONTACTS"
	ExternalCalendarMeetingEventResponsePropertiesHsEngagementSourceMergeObjects                ExternalCalendarMeetingEventResponsePropertiesHsEngagementSource = "MERGE_OBJECTS"
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
	ExternalCalendarMeetingEventResponsePropertiesHsEngagementSourcePropertyRestore             ExternalCalendarMeetingEventResponsePropertiesHsEngagementSource = "PROPERTY_RESTORE"
	ExternalCalendarMeetingEventResponsePropertiesHsEngagementSourcePropertySettings            ExternalCalendarMeetingEventResponsePropertiesHsEngagementSource = "PROPERTY_SETTINGS"
	ExternalCalendarMeetingEventResponsePropertiesHsEngagementSourceProspectingAgent            ExternalCalendarMeetingEventResponsePropertiesHsEngagementSource = "PROSPECTING_AGENT"
	ExternalCalendarMeetingEventResponsePropertiesHsEngagementSourceQuotas                      ExternalCalendarMeetingEventResponsePropertiesHsEngagementSource = "QUOTAS"
	ExternalCalendarMeetingEventResponsePropertiesHsEngagementSourceQuotes                      ExternalCalendarMeetingEventResponsePropertiesHsEngagementSource = "QUOTES"
	ExternalCalendarMeetingEventResponsePropertiesHsEngagementSourceRecyclingBin                ExternalCalendarMeetingEventResponsePropertiesHsEngagementSource = "RECYCLING_BIN"
	ExternalCalendarMeetingEventResponsePropertiesHsEngagementSourceRestoreObjects              ExternalCalendarMeetingEventResponsePropertiesHsEngagementSource = "RESTORE_OBJECTS"
	ExternalCalendarMeetingEventResponsePropertiesHsEngagementSourceSales                       ExternalCalendarMeetingEventResponsePropertiesHsEngagementSource = "SALES"
	ExternalCalendarMeetingEventResponsePropertiesHsEngagementSourceSalesMessages               ExternalCalendarMeetingEventResponsePropertiesHsEngagementSource = "SALES_MESSAGES"
	ExternalCalendarMeetingEventResponsePropertiesHsEngagementSourceSalesforce                  ExternalCalendarMeetingEventResponsePropertiesHsEngagementSource = "SALESFORCE"
	ExternalCalendarMeetingEventResponsePropertiesHsEngagementSourceSequences                   ExternalCalendarMeetingEventResponsePropertiesHsEngagementSource = "SEQUENCES"
	ExternalCalendarMeetingEventResponsePropertiesHsEngagementSourceSettings                    ExternalCalendarMeetingEventResponsePropertiesHsEngagementSource = "SETTINGS"
	ExternalCalendarMeetingEventResponsePropertiesHsEngagementSourceSidekick                    ExternalCalendarMeetingEventResponsePropertiesHsEngagementSource = "SIDEKICK"
	ExternalCalendarMeetingEventResponsePropertiesHsEngagementSourceSignals                     ExternalCalendarMeetingEventResponsePropertiesHsEngagementSource = "SIGNALS"
	ExternalCalendarMeetingEventResponsePropertiesHsEngagementSourceSlackIntegration            ExternalCalendarMeetingEventResponsePropertiesHsEngagementSource = "SLACK_INTEGRATION"
	ExternalCalendarMeetingEventResponsePropertiesHsEngagementSourceSocial                      ExternalCalendarMeetingEventResponsePropertiesHsEngagementSource = "SOCIAL"
	ExternalCalendarMeetingEventResponsePropertiesHsEngagementSourceSuccess                     ExternalCalendarMeetingEventResponsePropertiesHsEngagementSource = "SUCCESS"
	ExternalCalendarMeetingEventResponsePropertiesHsEngagementSourceTally                       ExternalCalendarMeetingEventResponsePropertiesHsEngagementSource = "TALLY"
	ExternalCalendarMeetingEventResponsePropertiesHsEngagementSourceTask                        ExternalCalendarMeetingEventResponsePropertiesHsEngagementSource = "TASK"
	ExternalCalendarMeetingEventResponsePropertiesHsEngagementSourceUnknown                     ExternalCalendarMeetingEventResponsePropertiesHsEngagementSource = "UNKNOWN"
	ExternalCalendarMeetingEventResponsePropertiesHsEngagementSourceWalIncremental              ExternalCalendarMeetingEventResponsePropertiesHsEngagementSource = "WAL_INCREMENTAL"
	ExternalCalendarMeetingEventResponsePropertiesHsEngagementSourceWorkflowContactDeleteAction ExternalCalendarMeetingEventResponsePropertiesHsEngagementSource = "WORKFLOW_CONTACT_DELETE_ACTION"
	ExternalCalendarMeetingEventResponsePropertiesHsEngagementSourceWorkflows                   ExternalCalendarMeetingEventResponsePropertiesHsEngagementSource = "WORKFLOWS"
)

type ExternalCalendarMeetingEventResponsePropertiesHsMeetingLocationType string

const (
	ExternalCalendarMeetingEventResponsePropertiesHsMeetingLocationTypeAddress ExternalCalendarMeetingEventResponsePropertiesHsMeetingLocationType = "ADDRESS"
	ExternalCalendarMeetingEventResponsePropertiesHsMeetingLocationTypeCustom  ExternalCalendarMeetingEventResponsePropertiesHsMeetingLocationType = "CUSTOM"
	ExternalCalendarMeetingEventResponsePropertiesHsMeetingLocationTypePhone   ExternalCalendarMeetingEventResponsePropertiesHsMeetingLocationType = "PHONE"
)

type ExternalCalenderMeetingEventResponse struct {
	ID            string                                         `json:"id,required"`
	CreatedAt     time.Time                                      `json:"createdAt,required" format:"date-time"`
	LastUpdatedAt time.Time                                      `json:"lastUpdatedAt,required" format:"date-time"`
	Properties    ExternalCalendarMeetingEventResponseProperties `json:"properties,required"`
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
	End   int64 `json:"end,required"`
	Start int64 `json:"start,required"`
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
	CommunicationTypeID string `json:"communicationTypeId,required"`
	Label               string `json:"label,required"`
	Required            bool   `json:"required,required"`
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
	Reminders                      []ExternalReminderParam `json:"reminders,omitzero,required"`
	ShouldIncludeInviteDescription bool                    `json:"shouldIncludeInviteDescription,required"`
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
	CanAddGuests  bool  `json:"canAddGuests,required"`
	MaxGuestCount int64 `json:"maxGuestCount,required"`
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
	CommunicationConsentCheckboxes      []ExternalCommunicationConsentCheckbox `json:"communicationConsentCheckboxes,required"`
	CommunicationConsentText            string                                 `json:"communicationConsentText,required"`
	IsLegitimateInterest                bool                                   `json:"isLegitimateInterest,required"`
	LegitimateInterestSubscriptionTypes []int64                                `json:"legitimateInterestSubscriptionTypes,required"`
	PrivacyPolicyText                   string                                 `json:"privacyPolicyText,required"`
	ProcessingConsentCheckboxLabel      string                                 `json:"processingConsentCheckboxLabel,required"`
	ProcessingConsentFooterText         string                                 `json:"processingConsentFooterText,required"`
	ProcessingConsentText               string                                 `json:"processingConsentText,required"`
	ProcessingConsentType               string                                 `json:"processingConsentType,required"`
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
	CommunicationTypeID string `json:"communicationTypeId,required"`
	Consented           bool   `json:"consented,required"`
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
	CommunicationTypeID string `json:"communicationTypeId,required"`
	Consented           bool   `json:"consented,required"`
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
	HasMore                    bool                                           `json:"hasMore,required"`
	LinkAvailabilityByDuration map[string]ExternalLinkAvailabilityForDuration `json:"linkAvailabilityByDuration,required"`
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
	AllUsersBusyTimes []ExternalUserBusyTimes  `json:"allUsersBusyTimes,required"`
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
	Availabilities        []ExternalMeetingAvailability `json:"availabilities,required"`
	MeetingDurationMillis int64                         `json:"meetingDurationMillis,required"`
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
	Avatar                    string `json:"avatar"`
	CompanyAvatar             string `json:"companyAvatar"`
	Headline                  string `json:"headline"`
	PublicDisplayAvatarOption string `json:"publicDisplayAvatarOption"`
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

type ExternalLinkFormField struct {
	FieldType  string           `json:"fieldType,required"`
	IsCustom   bool             `json:"isCustom,required"`
	IsRequired bool             `json:"isRequired,required"`
	Label      string           `json:"label,required"`
	Name       string           `json:"name,required"`
	Options    []ExternalOption `json:"options,required"`
	Type       string           `json:"type,required"`
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
	ID                   string    `json:"id,required"`
	CreatedAt            time.Time `json:"createdAt,required" format:"date-time"`
	DefaultLink          bool      `json:"defaultLink,required"`
	Link                 string    `json:"link,required"`
	OrganizerUserID      string    `json:"organizerUserId,required"`
	Slug                 string    `json:"slug,required"`
	Type                 string    `json:"type,required"`
	UserIDsOfLinkMembers []string  `json:"userIdsOfLinkMembers,required"`
	Name                 string    `json:"name"`
	UpdatedAt            time.Time `json:"updatedAt" format:"date-time"`
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

type ExternalMeetingAvailability struct {
	EndMillisUtc   int64 `json:"endMillisUtc,required"`
	StartMillisUtc int64 `json:"startMillisUtc,required"`
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
	Duration               int64                               `json:"duration,required"`
	Email                  string                              `json:"email,required"`
	FirstName              string                              `json:"firstName,required"`
	FormFields             []ExternalBookingFormFieldParam     `json:"formFields,omitzero,required"`
	LastName               string                              `json:"lastName,required"`
	LegalConsentResponses  []ExternalLegalConsentResponseParam `json:"legalConsentResponses,omitzero,required"`
	LikelyAvailableUserIDs []string                            `json:"likelyAvailableUserIds,omitzero,required"`
	Slug                   string                              `json:"slug,required"`
	StartTime              time.Time                           `json:"startTime,required" format:"date-time"`
	Locale                 param.Opt[string]                   `json:"locale,omitzero"`
	Timezone               param.Opt[string]                   `json:"timezone,omitzero"`
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
	BookingTimezone        string                         `json:"bookingTimezone,required"`
	CalendarEventID        string                         `json:"calendarEventId,required"`
	ContactID              string                         `json:"contactId,required"`
	Duration               int64                          `json:"duration,required"`
	End                    time.Time                      `json:"end,required" format:"date-time"`
	FormFields             []ExternalValidatedFormField   `json:"formFields,required"`
	GuestEmails            []string                       `json:"guestEmails,required"`
	IsOffline              bool                           `json:"isOffline,required"`
	LegalConsentResponses  []ExternalLegalConsentResponse `json:"legalConsentResponses,required"`
	Start                  time.Time                      `json:"start,required" format:"date-time"`
	Subject                string                         `json:"subject,required"`
	Locale                 string                         `json:"locale"`
	Location               string                         `json:"location"`
	WebConferenceMeetingID string                         `json:"webConferenceMeetingId"`
	WebConferenceURL       string                         `json:"webConferenceUrl"`
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
	Availability                map[string]ExternalClosedRange    `json:"availability,required"`
	Durations                   []int64                           `json:"durations,required"`
	FormFields                  []ExternalLinkFormField           `json:"formFields,required"`
	LegalConsentEnabled         bool                              `json:"legalConsentEnabled,required"`
	MeetingBufferTime           int64                             `json:"meetingBufferTime,required"`
	OwnerPrioritized            bool                              `json:"ownerPrioritized,required"`
	StartTimeIncrementMinutes   string                            `json:"startTimeIncrementMinutes,required"`
	WeeksToAdvertise            int64                             `json:"weeksToAdvertise,required"`
	CustomAvailabilityEndDate   int64                             `json:"customAvailabilityEndDate"`
	CustomAvailabilityStartDate int64                             `json:"customAvailabilityStartDate"`
	DisplayInfo                 ExternalLinkDisplayInfo           `json:"displayInfo"`
	GuestSettings               ExternalGuestSettings             `json:"guestSettings"`
	Language                    string                            `json:"language"`
	LegalConsentOptions         ExternalLegalConsentOptions       `json:"legalConsentOptions"`
	Locale                      string                            `json:"locale"`
	Location                    string                            `json:"location"`
	RedirectURL                 string                            `json:"redirectUrl"`
	WelcomeScreenInfo           ExternalMeetingsWelcomeScreenInfo `json:"welcomeScreenInfo"`
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

type ExternalMeetingsUser struct {
	ID               string              `json:"id,required"`
	CalendarProvider string              `json:"calendarProvider,required"`
	IsSalesStarter   bool                `json:"isSalesStarter,required"`
	UserID           string              `json:"userId,required"`
	UserProfile      ExternalUserProfile `json:"userProfile,required"`
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

type ExternalMeetingsWelcomeScreenInfo struct {
	Description       string `json:"description"`
	LogoURL           string `json:"logoUrl"`
	ShowWelcomeScreen bool   `json:"showWelcomeScreen"`
	Title             string `json:"title"`
	UseCompanyLogo    bool   `json:"useCompanyLogo"`
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
	Description  string  `json:"description,required"`
	DisplayOrder int64   `json:"displayOrder,required"`
	DoubleData   float64 `json:"doubleData,required"`
	Hidden       bool    `json:"hidden,required"`
	Label        string  `json:"label,required"`
	ReadOnly     bool    `json:"readOnly,required"`
	Value        string  `json:"value,required"`
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
	NumberOfTimeUnits int64  `json:"numberOfTimeUnits,required"`
	TimeUnit          string `json:"timeUnit,required"`
	paramObj
}

func (r ExternalReminderParam) MarshalJSON() (data []byte, err error) {
	type shadow ExternalReminderParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExternalReminderParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExternalTimeRange struct {
	End   int64 `json:"end,required"`
	Start int64 `json:"start,required"`
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
	BusyTimes    []ExternalTimeRange  `json:"busyTimes,required"`
	IsOffline    bool                 `json:"isOffline,required"`
	MeetingsUser ExternalMeetingsUser `json:"meetingsUser,required"`
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
	Email     string `json:"email,required"`
	FirstName string `json:"firstName"`
	FullName  string `json:"fullName"`
	LastName  string `json:"lastName"`
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
	IsCustom        bool   `json:"isCustom,required"`
	Label           string `json:"label,required"`
	Name            string `json:"name,required"`
	Value           string `json:"value,required"`
	FieldType       string `json:"fieldType"`
	TranslatedLabel string `json:"translatedLabel"`
	ValueLabel      string `json:"valueLabel"`
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
