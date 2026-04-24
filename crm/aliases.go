// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package crm

import (
	"github.com/stainless-sdks/hubspot-sdk-go/internal/apierror"
	"github.com/stainless-sdks/hubspot-sdk-go/packages/param"
	"github.com/stainless-sdks/hubspot-sdk-go/shared"
)

// aliased to make [param.APIUnion] private when embedding
type paramUnion = param.APIUnion

// aliased to make [param.APIObject] private when embedding
type paramObj = param.APIObject

type Error = apierror.Error

// This is an alias to an internal type.
type AbTestCreateRequestVNextParam = shared.AbTestCreateRequestVNextParam

// This is an alias to an internal type.
type ActionResponse = shared.ActionResponse

// The current status of the action, with possible values: CANCELED, COMPLETE,
// PENDING, PROCESSING.
//
// This is an alias to an internal type.
type ActionResponseStatus = shared.ActionResponseStatus

// Equals "CANCELED"
const ActionResponseStatusCanceled = shared.ActionResponseStatusCanceled

// Equals "COMPLETE"
const ActionResponseStatusComplete = shared.ActionResponseStatusComplete

// Equals "PENDING"
const ActionResponseStatusPending = shared.ActionResponseStatusPending

// Equals "PROCESSING"
const ActionResponseStatusProcessing = shared.ActionResponseStatusProcessing

// The definition of an association
//
// This is an alias to an internal type.
type AssociationDefinition = shared.AssociationDefinition

// This is an alias to an internal type.
type AssociationDefinitionEggParam = shared.AssociationDefinitionEggParam

// Defines the type, direction, and details of the relationship between two CRM
// objects.
//
// This is an alias to an internal type.
type AssociationSpec = shared.AssociationSpec

// The category of the association, such as "HUBSPOT_DEFINED".
//
// This is an alias to an internal type.
type AssociationSpecAssociationCategory = shared.AssociationSpecAssociationCategory

// Equals "HUBSPOT_DEFINED"
const AssociationSpecAssociationCategoryHubSpotDefined = shared.AssociationSpecAssociationCategoryHubSpotDefined

// Equals "INTEGRATOR_DEFINED"
const AssociationSpecAssociationCategoryIntegratorDefined = shared.AssociationSpecAssociationCategoryIntegratorDefined

// Equals "USER_DEFINED"
const AssociationSpecAssociationCategoryUserDefined = shared.AssociationSpecAssociationCategoryUserDefined

// Equals "WORK"
const AssociationSpecAssociationCategoryWork = shared.AssociationSpecAssociationCategoryWork

// Defines the type, direction, and details of the relationship between two CRM
// objects.
//
// This is an alias to an internal type.
type AssociationSpecParam = shared.AssociationSpecParam

// A HubSpot property option
//
// This is an alias to an internal type.
type AutomationActionsOption = shared.AutomationActionsOption

// A HubSpot property option
//
// This is an alias to an internal type.
type AutomationActionsOptionParam = shared.AutomationActionsOptionParam

// This is an alias to an internal type.
type BatchInputPropertyCreateParam = shared.BatchInputPropertyCreateParam

// This is an alias to an internal type.
type BatchInputPropertyNameParam = shared.BatchInputPropertyNameParam

// This is an alias to an internal type.
type BatchInputPublicObjectIDParam = shared.BatchInputPublicObjectIDParam

// This is an alias to an internal type.
type BatchInputStringParam = shared.BatchInputStringParam

// This is an alias to an internal type.
type BatchReadInputPropertyNameParam = shared.BatchReadInputPropertyNameParam

// This is an alias to an internal type.
type BatchReadInputPropertyNameDataSensitivity = shared.BatchReadInputPropertyNameDataSensitivity

// Equals "highly_sensitive"
const BatchReadInputPropertyNameDataSensitivityHighlySensitive = shared.BatchReadInputPropertyNameDataSensitivityHighlySensitive

// Equals "non_sensitive"
const BatchReadInputPropertyNameDataSensitivityNonSensitive = shared.BatchReadInputPropertyNameDataSensitivityNonSensitive

// Equals "sensitive"
const BatchReadInputPropertyNameDataSensitivitySensitive = shared.BatchReadInputPropertyNameDataSensitivitySensitive

// This is an alias to an internal type.
type CollectionResponsePropertyGroupNoPaging = shared.CollectionResponsePropertyGroupNoPaging

// This is an alias to an internal type.
type ErrorData = shared.ErrorData

// This is an alias to an internal type.
type ErrorDetail = shared.ErrorDetail

// This is an alias to an internal type.
type ForwardPaging = shared.ForwardPaging

// Specifies the paging information needed to retrieve the next set of results in a
// paginated API response
//
// This is an alias to an internal type.
type NextPage = shared.NextPage

// This is an alias to an internal type.
type ObjectTypeDefinition = shared.ObjectTypeDefinition

// This is an alias to an internal type.
type ObjectTypeDefinitionLabels = shared.ObjectTypeDefinitionLabels

// This is an alias to an internal type.
type ObjectTypeDefinitionLabelsParam = shared.ObjectTypeDefinitionLabelsParam

// This is an alias to an internal type.
type ObjectTypeDefinitionPatchParam = shared.ObjectTypeDefinitionPatchParam

// A HubSpot property option
//
// This is an alias to an internal type.
type Option = shared.Option

// This is an alias to an internal type.
type OptionInputParam = shared.OptionInputParam

// This is an alias to an internal type.
type Paging = shared.Paging

// specifies the paging information needed to retrieve the previous set of results
// in a paginated API response
//
// This is an alias to an internal type.
type PreviousPage = shared.PreviousPage

// A HubSpot property
//
// This is an alias to an internal type.
type Property = shared.Property

// Indicates the sensitivity level of the property, such as "non_sensitive",
// "sensitive", or "highly_sensitive".
//
// This is an alias to an internal type.
type PropertyDataSensitivity = shared.PropertyDataSensitivity

// Equals "highly_sensitive"
const PropertyDataSensitivityHighlySensitive = shared.PropertyDataSensitivityHighlySensitive

// Equals "non_sensitive"
const PropertyDataSensitivityNonSensitive = shared.PropertyDataSensitivityNonSensitive

// Equals "sensitive"
const PropertyDataSensitivitySensitive = shared.PropertyDataSensitivitySensitive

// Controls how date properties are displayed in the HubSpot UI, with options such
// as 'absolute', 'absolute_with_relative', 'time_since', and 'time_until'.
//
// This is an alias to an internal type.
type PropertyDateDisplayHint = shared.PropertyDateDisplayHint

// Equals "absolute"
const PropertyDateDisplayHintAbsolute = shared.PropertyDateDisplayHintAbsolute

// Equals "absolute_with_relative"
const PropertyDateDisplayHintAbsoluteWithRelative = shared.PropertyDateDisplayHintAbsoluteWithRelative

// Equals "time_since"
const PropertyDateDisplayHintTimeSince = shared.PropertyDateDisplayHintTimeSince

// Equals "time_until"
const PropertyDateDisplayHintTimeUntil = shared.PropertyDateDisplayHintTimeUntil

// Hint for how a number property is displayed and validated in HubSpot's UI. Can
// be: "unformatted", "formatted", "currency", "percentage", "duration", or
// "probability".
//
// This is an alias to an internal type.
type PropertyNumberDisplayHint = shared.PropertyNumberDisplayHint

// Equals "currency"
const PropertyNumberDisplayHintCurrency = shared.PropertyNumberDisplayHintCurrency

// Equals "duration"
const PropertyNumberDisplayHintDuration = shared.PropertyNumberDisplayHintDuration

// Equals "formatted"
const PropertyNumberDisplayHintFormatted = shared.PropertyNumberDisplayHintFormatted

// Equals "percentage"
const PropertyNumberDisplayHintPercentage = shared.PropertyNumberDisplayHintPercentage

// Equals "probability"
const PropertyNumberDisplayHintProbability = shared.PropertyNumberDisplayHintProbability

// Equals "unformatted"
const PropertyNumberDisplayHintUnformatted = shared.PropertyNumberDisplayHintUnformatted

// This is an alias to an internal type.
type PropertyCreateParam = shared.PropertyCreateParam

// This is an alias to an internal type.
type PropertyCreateFieldType = shared.PropertyCreateFieldType

// Equals "booleancheckbox"
const PropertyCreateFieldTypeBooleancheckbox = shared.PropertyCreateFieldTypeBooleancheckbox

// Equals "calculation_equation"
const PropertyCreateFieldTypeCalculationEquation = shared.PropertyCreateFieldTypeCalculationEquation

// Equals "checkbox"
const PropertyCreateFieldTypeCheckbox = shared.PropertyCreateFieldTypeCheckbox

// Equals "date"
const PropertyCreateFieldTypeDate = shared.PropertyCreateFieldTypeDate

// Equals "file"
const PropertyCreateFieldTypeFile = shared.PropertyCreateFieldTypeFile

// Equals "html"
const PropertyCreateFieldTypeHTML = shared.PropertyCreateFieldTypeHTML

// Equals "number"
const PropertyCreateFieldTypeNumber = shared.PropertyCreateFieldTypeNumber

// Equals "phonenumber"
const PropertyCreateFieldTypePhonenumber = shared.PropertyCreateFieldTypePhonenumber

// Equals "radio"
const PropertyCreateFieldTypeRadio = shared.PropertyCreateFieldTypeRadio

// Equals "select"
const PropertyCreateFieldTypeSelect = shared.PropertyCreateFieldTypeSelect

// Equals "text"
const PropertyCreateFieldTypeText = shared.PropertyCreateFieldTypeText

// Equals "textarea"
const PropertyCreateFieldTypeTextarea = shared.PropertyCreateFieldTypeTextarea

// This is an alias to an internal type.
type PropertyCreateType = shared.PropertyCreateType

// Equals "bool"
const PropertyCreateTypeBool = shared.PropertyCreateTypeBool

// Equals "date"
const PropertyCreateTypeDate = shared.PropertyCreateTypeDate

// Equals "datetime"
const PropertyCreateTypeDatetime = shared.PropertyCreateTypeDatetime

// Equals "enumeration"
const PropertyCreateTypeEnumeration = shared.PropertyCreateTypeEnumeration

// Equals "number"
const PropertyCreateTypeNumber = shared.PropertyCreateTypeNumber

// Equals "phone_number"
const PropertyCreateTypePhoneNumber = shared.PropertyCreateTypePhoneNumber

// Equals "string"
const PropertyCreateTypeString = shared.PropertyCreateTypeString

// This is an alias to an internal type.
type PropertyCreateDataSensitivity = shared.PropertyCreateDataSensitivity

// Equals "highly_sensitive"
const PropertyCreateDataSensitivityHighlySensitive = shared.PropertyCreateDataSensitivityHighlySensitive

// Equals "non_sensitive"
const PropertyCreateDataSensitivityNonSensitive = shared.PropertyCreateDataSensitivityNonSensitive

// Equals "sensitive"
const PropertyCreateDataSensitivitySensitive = shared.PropertyCreateDataSensitivitySensitive

// This is an alias to an internal type.
type PropertyCreateNumberDisplayHint = shared.PropertyCreateNumberDisplayHint

// Equals "currency"
const PropertyCreateNumberDisplayHintCurrency = shared.PropertyCreateNumberDisplayHintCurrency

// Equals "duration"
const PropertyCreateNumberDisplayHintDuration = shared.PropertyCreateNumberDisplayHintDuration

// Equals "formatted"
const PropertyCreateNumberDisplayHintFormatted = shared.PropertyCreateNumberDisplayHintFormatted

// Equals "percentage"
const PropertyCreateNumberDisplayHintPercentage = shared.PropertyCreateNumberDisplayHintPercentage

// Equals "probability"
const PropertyCreateNumberDisplayHintProbability = shared.PropertyCreateNumberDisplayHintProbability

// Equals "unformatted"
const PropertyCreateNumberDisplayHintUnformatted = shared.PropertyCreateNumberDisplayHintUnformatted

// This is an alias to an internal type.
type PropertyGroup = shared.PropertyGroup

// This is an alias to an internal type.
type PropertyGroupCreateParam = shared.PropertyGroupCreateParam

// This is an alias to an internal type.
type PropertyGroupUpdateParam = shared.PropertyGroupUpdateParam

// This is an alias to an internal type.
type PropertyModificationMetadata = shared.PropertyModificationMetadata

// This is an alias to an internal type.
type PropertyNameParam = shared.PropertyNameParam

// Represents a single custom property of a marketing event, storing its name,
// value, metadata (like source, timestamp, and sensitivity), and related audit
// information for tracking changes.
//
// This is an alias to an internal type.
type PropertyValue = shared.PropertyValue

// The sensitivity level of the property, such as "non_sensitive", "sensitive", and
// "highly_sensitive".
//
// This is an alias to an internal type.
type PropertyValueDataSensitivity = shared.PropertyValueDataSensitivity

// Equals "high"
const PropertyValueDataSensitivityHigh = shared.PropertyValueDataSensitivityHigh

// Equals "none"
const PropertyValueDataSensitivityNone = shared.PropertyValueDataSensitivityNone

// Equals "standard"
const PropertyValueDataSensitivityStandard = shared.PropertyValueDataSensitivityStandard

// The origin of the property value, such as "IMPORT" or "API".
//
// This is an alias to an internal type.
type PropertyValueSource = shared.PropertyValueSource

// Equals "ACADEMY"
const PropertyValueSourceAcademy = shared.PropertyValueSourceAcademy

// Equals "ACCEPTANCE_TEST"
const PropertyValueSourceAcceptanceTest = shared.PropertyValueSourceAcceptanceTest

// Equals "ACTIVITY_AUTO_ASSOCIATE"
const PropertyValueSourceActivityAutoAssociate = shared.PropertyValueSourceActivityAutoAssociate

// Equals "ACTIVITY_LOG_REVERT"
const PropertyValueSourceActivityLogRevert = shared.PropertyValueSourceActivityLogRevert

// Equals "ADS"
const PropertyValueSourceAds = shared.PropertyValueSourceAds

// Equals "AI_GROUP"
const PropertyValueSourceAIGroup = shared.PropertyValueSourceAIGroup

// Equals "ANALYTICS"
const PropertyValueSourceAnalytics = shared.PropertyValueSourceAnalytics

// Equals "API"
const PropertyValueSourceAPI = shared.PropertyValueSourceAPI

// Equals "APPROVALS"
const PropertyValueSourceApprovals = shared.PropertyValueSourceApprovals

// Equals "ASSISTS"
const PropertyValueSourceAssists = shared.PropertyValueSourceAssists

// Equals "ASSOCIATIONS"
const PropertyValueSourceAssociations = shared.PropertyValueSourceAssociations

// Equals "AUTO_ASSOCIATE_BY_DOMAIN"
const PropertyValueSourceAutoAssociateByDomain = shared.PropertyValueSourceAutoAssociateByDomain

// Equals "AUTOMATION_JOURNEY"
const PropertyValueSourceAutomationJourney = shared.PropertyValueSourceAutomationJourney

// Equals "AUTOMATION_PLATFORM"
const PropertyValueSourceAutomationPlatform = shared.PropertyValueSourceAutomationPlatform

// Equals "AVATARS_SERVICE"
const PropertyValueSourceAvatarsService = shared.PropertyValueSourceAvatarsService

// Equals "BATCH_UPDATE"
const PropertyValueSourceBatchUpdate = shared.PropertyValueSourceBatchUpdate

// Equals "BCC_TO_CRM"
const PropertyValueSourceBccToCrm = shared.PropertyValueSourceBccToCrm

// Equals "BEHAVIORAL_EVENTS"
const PropertyValueSourceBehavioralEvents = shared.PropertyValueSourceBehavioralEvents

// Equals "BET_ASSIGNMENT"
const PropertyValueSourceBetAssignment = shared.PropertyValueSourceBetAssignment

// Equals "BET_CRM_CONNECTOR"
const PropertyValueSourceBetCrmConnector = shared.PropertyValueSourceBetCrmConnector

// Equals "BIDEN"
const PropertyValueSourceBiden = shared.PropertyValueSourceBiden

// Equals "BILLING"
const PropertyValueSourceBilling = shared.PropertyValueSourceBilling

// Equals "BOT"
const PropertyValueSourceBot = shared.PropertyValueSourceBot

// Equals "CALCULATED"
const PropertyValueSourceCalculated = shared.PropertyValueSourceCalculated

// Equals "CENTRAL_EXCHANGE_RATES"
const PropertyValueSourceCentralExchangeRates = shared.PropertyValueSourceCentralExchangeRates

// Equals "CHATSPOT"
const PropertyValueSourceChatspot = shared.PropertyValueSourceChatspot

// Equals "CLONE_OBJECTS"
const PropertyValueSourceCloneObjects = shared.PropertyValueSourceCloneObjects

// Equals "COMMUNICATOR"
const PropertyValueSourceCommunicator = shared.PropertyValueSourceCommunicator

// Equals "COMPANIES"
const PropertyValueSourceCompanies = shared.PropertyValueSourceCompanies

// Equals "COMPANY_FAMILIES"
const PropertyValueSourceCompanyFamilies = shared.PropertyValueSourceCompanyFamilies

// Equals "COMPANY_INSIGHTS"
const PropertyValueSourceCompanyInsights = shared.PropertyValueSourceCompanyInsights

// Equals "CONNECTED_ACCOUNT"
const PropertyValueSourceConnectedAccount = shared.PropertyValueSourceConnectedAccount

// Equals "CONTACTS"
const PropertyValueSourceContacts = shared.PropertyValueSourceContacts

// Equals "CONTACTS_WEB"
const PropertyValueSourceContactsWeb = shared.PropertyValueSourceContactsWeb

// Equals "CONTENT_MEMBERSHIP"
const PropertyValueSourceContentMembership = shared.PropertyValueSourceContentMembership

// Equals "CONVERSATIONAL_ENRICHMENT"
const PropertyValueSourceConversationalEnrichment = shared.PropertyValueSourceConversationalEnrichment

// Equals "CONVERSATIONS"
const PropertyValueSourceConversations = shared.PropertyValueSourceConversations

// Equals "CRM_PROCESSES_PLATFORM"
const PropertyValueSourceCrmProcessesPlatform = shared.PropertyValueSourceCrmProcessesPlatform

// Equals "CRM_UI"
const PropertyValueSourceCrmUi = shared.PropertyValueSourceCrmUi

// Equals "CRM_UI_BULK_ACTION"
const PropertyValueSourceCrmUiBulkAction = shared.PropertyValueSourceCrmUiBulkAction

// Equals "CUSTOMER_AGENT"
const PropertyValueSourceCustomerAgent = shared.PropertyValueSourceCustomerAgent

// Equals "DATA_ENRICHMENT"
const PropertyValueSourceDataEnrichment = shared.PropertyValueSourceDataEnrichment

// Equals "DATA_QUALITY"
const PropertyValueSourceDataQuality = shared.PropertyValueSourceDataQuality

// Equals "DATASET"
const PropertyValueSourceDataset = shared.PropertyValueSourceDataset

// Equals "DEALS"
const PropertyValueSourceDeals = shared.PropertyValueSourceDeals

// Equals "DEFAULT"
const PropertyValueSourceDefault = shared.PropertyValueSourceDefault

// Equals "DELETE_OBJECTS"
const PropertyValueSourceDeleteObjects = shared.PropertyValueSourceDeleteObjects

// Equals "EMAIL"
const PropertyValueSourceEmail = shared.PropertyValueSourceEmail

// Equals "EMAIL_INBOX_IMPORT"
const PropertyValueSourceEmailInboxImport = shared.PropertyValueSourceEmailInboxImport

// Equals "EMAIL_INTEGRATION"
const PropertyValueSourceEmailIntegration = shared.PropertyValueSourceEmailIntegration

// Equals "ENGAGEMENTS"
const PropertyValueSourceEngagements = shared.PropertyValueSourceEngagements

// Equals "EXTENSION"
const PropertyValueSourceExtension = shared.PropertyValueSourceExtension

// Equals "FILE_MANAGER"
const PropertyValueSourceFileManager = shared.PropertyValueSourceFileManager

// Equals "FLYWHEEL_PRODUCT_DATA_SYNC"
const PropertyValueSourceFlywheelProductDataSync = shared.PropertyValueSourceFlywheelProductDataSync

// Equals "FORECASTING"
const PropertyValueSourceForecasting = shared.PropertyValueSourceForecasting

// Equals "FORM"
const PropertyValueSourceForm = shared.PropertyValueSourceForm

// Equals "FORWARD_TO_CRM"
const PropertyValueSourceForwardToCrm = shared.PropertyValueSourceForwardToCrm

// Equals "GMAIL_INTEGRATION"
const PropertyValueSourceGmailIntegration = shared.PropertyValueSourceGmailIntegration

// Equals "GOALS"
const PropertyValueSourceGoals = shared.PropertyValueSourceGoals

// Equals "HEISENBERG"
const PropertyValueSourceHeisenberg = shared.PropertyValueSourceHeisenberg

// Equals "HELP_DESK"
const PropertyValueSourceHelpDesk = shared.PropertyValueSourceHelpDesk

// Equals "HELP_DESK_AI"
const PropertyValueSourceHelpDeskAI = shared.PropertyValueSourceHelpDeskAI

// Equals "IMPORT"
const PropertyValueSourceImport = shared.PropertyValueSourceImport

// Equals "INTEGRATION"
const PropertyValueSourceIntegration = shared.PropertyValueSourceIntegration

// Equals "INTEGRATIONS_PLATFORM"
const PropertyValueSourceIntegrationsPlatform = shared.PropertyValueSourceIntegrationsPlatform

// Equals "INTEGRATIONS_SYNC"
const PropertyValueSourceIntegrationsSync = shared.PropertyValueSourceIntegrationsSync

// Equals "INTENT"
const PropertyValueSourceIntent = shared.PropertyValueSourceIntent

// Equals "INTERNAL_PROCESSING"
const PropertyValueSourceInternalProcessing = shared.PropertyValueSourceInternalProcessing

// Equals "LEADIN"
const PropertyValueSourceLeadin = shared.PropertyValueSourceLeadin

// Equals "LEGAL_BASIS_REMEDIATION"
const PropertyValueSourceLegalBasisRemediation = shared.PropertyValueSourceLegalBasisRemediation

// Equals "MARKET_SOURCING"
const PropertyValueSourceMarketSourcing = shared.PropertyValueSourceMarketSourcing

// Equals "MARKETPLACE"
const PropertyValueSourceMarketplace = shared.PropertyValueSourceMarketplace

// Equals "MARKETS"
const PropertyValueSourceMarkets = shared.PropertyValueSourceMarkets

// Equals "MEETINGS"
const PropertyValueSourceMeetings = shared.PropertyValueSourceMeetings

// Equals "MERGE_COMPANIES"
const PropertyValueSourceMergeCompanies = shared.PropertyValueSourceMergeCompanies

// Equals "MERGE_CONTACTS"
const PropertyValueSourceMergeContacts = shared.PropertyValueSourceMergeContacts

// Equals "MERGE_OBJECTS"
const PropertyValueSourceMergeObjects = shared.PropertyValueSourceMergeObjects

// Equals "MERGE_REVERT_OBJECTS"
const PropertyValueSourceMergeRevertObjects = shared.PropertyValueSourceMergeRevertObjects

// Equals "MICROAPPS"
const PropertyValueSourceMicroapps = shared.PropertyValueSourceMicroapps

// Equals "MIGRATION"
const PropertyValueSourceMigration = shared.PropertyValueSourceMigration

// Equals "MOBILE_ANDROID"
const PropertyValueSourceMobileAndroid = shared.PropertyValueSourceMobileAndroid

// Equals "MOBILE_IOS"
const PropertyValueSourceMobileIos = shared.PropertyValueSourceMobileIos

// Equals "PAYMENTS"
const PropertyValueSourcePayments = shared.PropertyValueSourcePayments

// Equals "PIPELINE_SETTINGS"
const PropertyValueSourcePipelineSettings = shared.PropertyValueSourcePipelineSettings

// Equals "PLAYBOOKS"
const PropertyValueSourcePlaybooks = shared.PropertyValueSourcePlaybooks

// Equals "PORTAL_OBJECT_SYNC"
const PropertyValueSourcePortalObjectSync = shared.PropertyValueSourcePortalObjectSync

// Equals "PORTAL_USER_ASSOCIATOR"
const PropertyValueSourcePortalUserAssociator = shared.PropertyValueSourcePortalUserAssociator

// Equals "PRESENTATIONS"
const PropertyValueSourcePresentations = shared.PropertyValueSourcePresentations

// Equals "PRIMARY_AUTOMATION"
const PropertyValueSourcePrimaryAutomation = shared.PropertyValueSourcePrimaryAutomation

// Equals "PROPERTY_DEFAULT_VALUE"
const PropertyValueSourcePropertyDefaultValue = shared.PropertyValueSourcePropertyDefaultValue

// Equals "PROPERTY_RESTORE"
const PropertyValueSourcePropertyRestore = shared.PropertyValueSourcePropertyRestore

// Equals "PROPERTY_SETTINGS"
const PropertyValueSourcePropertySettings = shared.PropertyValueSourcePropertySettings

// Equals "PROSPECTING_AGENT"
const PropertyValueSourceProspectingAgent = shared.PropertyValueSourceProspectingAgent

// Equals "QUOTAS"
const PropertyValueSourceQuotas = shared.PropertyValueSourceQuotas

// Equals "QUOTES"
const PropertyValueSourceQuotes = shared.PropertyValueSourceQuotes

// Equals "RECYCLING_BIN"
const PropertyValueSourceRecyclingBin = shared.PropertyValueSourceRecyclingBin

// Equals "RESTORE_OBJECTS"
const PropertyValueSourceRestoreObjects = shared.PropertyValueSourceRestoreObjects

// Equals "REVENUE_PLATFORM"
const PropertyValueSourceRevenuePlatform = shared.PropertyValueSourceRevenuePlatform

// Equals "SALES"
const PropertyValueSourceSales = shared.PropertyValueSourceSales

// Equals "SALES_MESSAGES"
const PropertyValueSourceSalesMessages = shared.PropertyValueSourceSalesMessages

// Equals "SALESFORCE"
const PropertyValueSourceSalesforce = shared.PropertyValueSourceSalesforce

// Equals "SEQUENCES"
const PropertyValueSourceSequences = shared.PropertyValueSourceSequences

// Equals "SETTINGS"
const PropertyValueSourceSettings = shared.PropertyValueSourceSettings

// Equals "SIDEKICK"
const PropertyValueSourceSidekick = shared.PropertyValueSourceSidekick

// Equals "SIGNALS"
const PropertyValueSourceSignals = shared.PropertyValueSourceSignals

// Equals "SLACK_INTEGRATION"
const PropertyValueSourceSlackIntegration = shared.PropertyValueSourceSlackIntegration

// Equals "SMART_DATA_CAPTURE"
const PropertyValueSourceSmartDataCapture = shared.PropertyValueSourceSmartDataCapture

// Equals "SOCIAL"
const PropertyValueSourceSocial = shared.PropertyValueSourceSocial

// Equals "SUCCESS"
const PropertyValueSourceSuccess = shared.PropertyValueSourceSuccess

// Equals "TALLY"
const PropertyValueSourceTally = shared.PropertyValueSourceTally

// Equals "TASK"
const PropertyValueSourceTask = shared.PropertyValueSourceTask

// Equals "UNKNOWN"
const PropertyValueSourceUnknown = shared.PropertyValueSourceUnknown

// Equals "WAL_INCREMENTAL"
const PropertyValueSourceWalIncremental = shared.PropertyValueSourceWalIncremental

// Equals "WORK_UI"
const PropertyValueSourceWorkUi = shared.PropertyValueSourceWorkUi

// Equals "WORKFLOW_CONTACT_DELETE_ACTION"
const PropertyValueSourceWorkflowContactDeleteAction = shared.PropertyValueSourceWorkflowContactDeleteAction

// Equals "WORKFLOWS"
const PropertyValueSourceWorkflows = shared.PropertyValueSourceWorkflows

// Represents a single custom property of a marketing event, storing its name,
// value, metadata (like source, timestamp, and sensitivity), and related audit
// information for tracking changes.
//
// This is an alias to an internal type.
type PropertyValueParam = shared.PropertyValueParam

// Contains the Id of a Public Object
//
// This is an alias to an internal type.
type PublicObjectID = shared.PublicObjectID

// Contains the Id of a Public Object
//
// This is an alias to an internal type.
type PublicObjectIDParam = shared.PublicObjectIDParam

// Ye olde error
//
// This is an alias to an internal type.
type StandardError = shared.StandardError

// This is an alias to an internal type.
type TaskLocator = shared.TaskLocator

// This is an alias to an internal type.
type VersionUser = shared.VersionUser
