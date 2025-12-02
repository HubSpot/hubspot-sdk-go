// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package events

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

// Request body object for creating A/B tests.
//
// This is an alias to an internal type.
type AbTestCreateRequestVNextParam = shared.AbTestCreateRequestVNextParam

// This is an alias to an internal type.
type ActionResponse = shared.ActionResponse

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
const AssociationSpecAssociationCategoryHubspotDefined = shared.AssociationSpecAssociationCategoryHubspotDefined

// Equals "INTEGRATOR_DEFINED"
const AssociationSpecAssociationCategoryIntegratorDefined = shared.AssociationSpecAssociationCategoryIntegratorDefined

// Equals "USER_DEFINED"
const AssociationSpecAssociationCategoryUserDefined = shared.AssociationSpecAssociationCategoryUserDefined

// Defines the type, direction, and details of the relationship between two CRM
// objects.
//
// This is an alias to an internal type.
type AssociationSpecParam = shared.AssociationSpecParam

// This is an alias to an internal type.
type BatchInputPropertyCreateParam = shared.BatchInputPropertyCreateParam

// This is an alias to an internal type.
type BatchInputPropertyNameParam = shared.BatchInputPropertyNameParam

// This is an alias to an internal type.
type BatchInputPublicObjectIDParam = shared.BatchInputPublicObjectIDParam

// Wrapper for providing an array of strings as inputs.
//
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
type BatchResponseProperty = shared.BatchResponseProperty

// This is an alias to an internal type.
type BatchResponsePropertyStatus = shared.BatchResponsePropertyStatus

// Equals "CANCELED"
const BatchResponsePropertyStatusCanceled = shared.BatchResponsePropertyStatusCanceled

// Equals "COMPLETE"
const BatchResponsePropertyStatusComplete = shared.BatchResponsePropertyStatusComplete

// Equals "PENDING"
const BatchResponsePropertyStatusPending = shared.BatchResponsePropertyStatusPending

// Equals "PROCESSING"
const BatchResponsePropertyStatusProcessing = shared.BatchResponsePropertyStatusProcessing

// This is an alias to an internal type.
type Error = shared.Error

// This is an alias to an internal type.
type ErrorDetail = shared.ErrorDetail

// This is an alias to an internal type.
type ForwardPaging = shared.ForwardPaging

// HubDbTableRowV3Wrapper
//
// This is an alias to an internal type.
type HubDBTableRowV3Wrapper = shared.HubDBTableRowV3Wrapper

// Specifies the paging information needed to retrieve the next set of results in a
// paginated API response
//
// This is an alias to an internal type.
type NextPage = shared.NextPage

// This is an alias to an internal type.
type ObjectTypeDefinitionLabels = shared.ObjectTypeDefinitionLabels

// This is an alias to an internal type.
type ObjectTypeDefinitionLabelsParam = shared.ObjectTypeDefinitionLabelsParam

// The options available when a property is an enumeration
//
// This is an alias to an internal type.
type Option = shared.Option

// The options available when a property is an enumeration
//
// This is an alias to an internal type.
type OptionParam = shared.OptionParam

// This is an alias to an internal type.
type OptionInputParam = shared.OptionInputParam

// This is an alias to an internal type.
type Paging = shared.Paging

// specifies the paging information needed to retrieve the previous set of results
// in a paginated API response
//
// This is an alias to an internal type.
type PreviousPage = shared.PreviousPage

// Defines a property
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
type PropertyGroupCreateParam = shared.PropertyGroupCreateParam

// This is an alias to an internal type.
type PropertyGroupUpdateParam = shared.PropertyGroupUpdateParam

// This is an alias to an internal type.
type PropertyModificationMetadata = shared.PropertyModificationMetadata

// This is an alias to an internal type.
type PropertyNameParam = shared.PropertyNameParam

// This is an alias to an internal type.
type PublicAbsoluteComparativeTimestampRefineBy = shared.PublicAbsoluteComparativeTimestampRefineBy

// This is an alias to an internal type.
type PublicAbsoluteComparativeTimestampRefineByType = shared.PublicAbsoluteComparativeTimestampRefineByType

// Equals "ABSOLUTE_COMPARATIVE"
const PublicAbsoluteComparativeTimestampRefineByTypeAbsoluteComparative = shared.PublicAbsoluteComparativeTimestampRefineByTypeAbsoluteComparative

// This is an alias to an internal type.
type PublicAbsoluteComparativeTimestampRefineByParam = shared.PublicAbsoluteComparativeTimestampRefineByParam

// This is an alias to an internal type.
type PublicAbsoluteRangedTimestampRefineBy = shared.PublicAbsoluteRangedTimestampRefineBy

// This is an alias to an internal type.
type PublicAbsoluteRangedTimestampRefineByType = shared.PublicAbsoluteRangedTimestampRefineByType

// Equals "ABSOLUTE_RANGED"
const PublicAbsoluteRangedTimestampRefineByTypeAbsoluteRanged = shared.PublicAbsoluteRangedTimestampRefineByTypeAbsoluteRanged

// This is an alias to an internal type.
type PublicAbsoluteRangedTimestampRefineByParam = shared.PublicAbsoluteRangedTimestampRefineByParam

// This is an alias to an internal type.
type PublicAdsSearchFilter = shared.PublicAdsSearchFilter

// This is an alias to an internal type.
type PublicAdsSearchFilterFilterType = shared.PublicAdsSearchFilterFilterType

// Equals "ADS_SEARCH"
const PublicAdsSearchFilterFilterTypeAdsSearch = shared.PublicAdsSearchFilterFilterTypeAdsSearch

// This is an alias to an internal type.
type PublicAdsSearchFilterParam = shared.PublicAdsSearchFilterParam

// This is an alias to an internal type.
type PublicAdsTimeFilter = shared.PublicAdsTimeFilter

// This is an alias to an internal type.
type PublicAdsTimeFilterFilterType = shared.PublicAdsTimeFilterFilterType

// Equals "ADS_TIME"
const PublicAdsTimeFilterFilterTypeAdsTime = shared.PublicAdsTimeFilterFilterTypeAdsTime

// This is an alias to an internal type.
type PublicAdsTimeFilterPruningRefineByUnion = shared.PublicAdsTimeFilterPruningRefineByUnion

// This is an alias to an internal type.
type PublicAdsTimeFilterParam = shared.PublicAdsTimeFilterParam

// This is an alias to an internal type.
type PublicAdsTimeFilterPruningRefineByUnionParam = shared.PublicAdsTimeFilterPruningRefineByUnionParam

// This is an alias to an internal type.
type PublicAllHistoryRefineBy = shared.PublicAllHistoryRefineBy

// This is an alias to an internal type.
type PublicAllHistoryRefineByType = shared.PublicAllHistoryRefineByType

// Equals "ALL_HISTORY"
const PublicAllHistoryRefineByTypeAllHistory = shared.PublicAllHistoryRefineByTypeAllHistory

// This is an alias to an internal type.
type PublicAllHistoryRefineByParam = shared.PublicAllHistoryRefineByParam

// This is an alias to an internal type.
type PublicAllPropertyTypesOperation = shared.PublicAllPropertyTypesOperation

// This is an alias to an internal type.
type PublicAllPropertyTypesOperationOperationType = shared.PublicAllPropertyTypesOperationOperationType

// Equals "ALL_PROPERTY"
const PublicAllPropertyTypesOperationOperationTypeAllProperty = shared.PublicAllPropertyTypesOperationOperationTypeAllProperty

// This is an alias to an internal type.
type PublicAllPropertyTypesOperationParam = shared.PublicAllPropertyTypesOperationParam

// This is an alias to an internal type.
type PublicAndFilterBranch = shared.PublicAndFilterBranch

// This is an alias to an internal type.
type PublicAndFilterBranchFilterBranchUnion = shared.PublicAndFilterBranchFilterBranchUnion

// This is an alias to an internal type.
type PublicAndFilterBranchFilterBranchType = shared.PublicAndFilterBranchFilterBranchType

// Equals "AND"
const PublicAndFilterBranchFilterBranchTypeAnd = shared.PublicAndFilterBranchFilterBranchTypeAnd

// This is an alias to an internal type.
type PublicAndFilterBranchFilterUnion = shared.PublicAndFilterBranchFilterUnion

// This is an alias to an internal type.
type PublicAndFilterBranchParam = shared.PublicAndFilterBranchParam

// This is an alias to an internal type.
type PublicAndFilterBranchFilterBranchUnionParam = shared.PublicAndFilterBranchFilterBranchUnionParam

// This is an alias to an internal type.
type PublicAndFilterBranchFilterUnionParam = shared.PublicAndFilterBranchFilterUnionParam

// This is an alias to an internal type.
type PublicAssociationFilterBranch = shared.PublicAssociationFilterBranch

// This is an alias to an internal type.
type PublicAssociationFilterBranchFilterBranchUnion = shared.PublicAssociationFilterBranchFilterBranchUnion

// This is an alias to an internal type.
type PublicAssociationFilterBranchFilterBranchType = shared.PublicAssociationFilterBranchFilterBranchType

// Equals "ASSOCIATION"
const PublicAssociationFilterBranchFilterBranchTypeAssociation = shared.PublicAssociationFilterBranchFilterBranchTypeAssociation

// This is an alias to an internal type.
type PublicAssociationFilterBranchFilterUnion = shared.PublicAssociationFilterBranchFilterUnion

// This is an alias to an internal type.
type PublicAssociationFilterBranchParam = shared.PublicAssociationFilterBranchParam

// This is an alias to an internal type.
type PublicAssociationFilterBranchFilterBranchUnionParam = shared.PublicAssociationFilterBranchFilterBranchUnionParam

// This is an alias to an internal type.
type PublicAssociationFilterBranchFilterUnionParam = shared.PublicAssociationFilterBranchFilterUnionParam

// This is an alias to an internal type.
type PublicAssociationInListFilter = shared.PublicAssociationInListFilter

// This is an alias to an internal type.
type PublicAssociationInListFilterCoalescingRefineByUnion = shared.PublicAssociationInListFilterCoalescingRefineByUnion

// This is an alias to an internal type.
type PublicAssociationInListFilterFilterType = shared.PublicAssociationInListFilterFilterType

// Equals "ASSOCIATION"
const PublicAssociationInListFilterFilterTypeAssociation = shared.PublicAssociationInListFilterFilterTypeAssociation

// This is an alias to an internal type.
type PublicAssociationInListFilterParam = shared.PublicAssociationInListFilterParam

// This is an alias to an internal type.
type PublicAssociationInListFilterCoalescingRefineByUnionParam = shared.PublicAssociationInListFilterCoalescingRefineByUnionParam

// This is an alias to an internal type.
type PublicBoolPropertyOperation = shared.PublicBoolPropertyOperation

// This is an alias to an internal type.
type PublicBoolPropertyOperationOperationType = shared.PublicBoolPropertyOperationOperationType

// Equals "BOOL"
const PublicBoolPropertyOperationOperationTypeBool = shared.PublicBoolPropertyOperationOperationTypeBool

// This is an alias to an internal type.
type PublicBoolPropertyOperationParam = shared.PublicBoolPropertyOperationParam

// This is an alias to an internal type.
type PublicCalendarDatePropertyOperation = shared.PublicCalendarDatePropertyOperation

// This is an alias to an internal type.
type PublicCalendarDatePropertyOperationOperationType = shared.PublicCalendarDatePropertyOperationOperationType

// Equals "CALENDAR_DATE"
const PublicCalendarDatePropertyOperationOperationTypeCalendarDate = shared.PublicCalendarDatePropertyOperationOperationTypeCalendarDate

// This is an alias to an internal type.
type PublicCalendarDatePropertyOperationFiscalYearStart = shared.PublicCalendarDatePropertyOperationFiscalYearStart

// Equals "APRIL"
const PublicCalendarDatePropertyOperationFiscalYearStartApril = shared.PublicCalendarDatePropertyOperationFiscalYearStartApril

// Equals "AUGUST"
const PublicCalendarDatePropertyOperationFiscalYearStartAugust = shared.PublicCalendarDatePropertyOperationFiscalYearStartAugust

// Equals "DECEMBER"
const PublicCalendarDatePropertyOperationFiscalYearStartDecember = shared.PublicCalendarDatePropertyOperationFiscalYearStartDecember

// Equals "FEBRUARY"
const PublicCalendarDatePropertyOperationFiscalYearStartFebruary = shared.PublicCalendarDatePropertyOperationFiscalYearStartFebruary

// Equals "JANUARY"
const PublicCalendarDatePropertyOperationFiscalYearStartJanuary = shared.PublicCalendarDatePropertyOperationFiscalYearStartJanuary

// Equals "JULY"
const PublicCalendarDatePropertyOperationFiscalYearStartJuly = shared.PublicCalendarDatePropertyOperationFiscalYearStartJuly

// Equals "JUNE"
const PublicCalendarDatePropertyOperationFiscalYearStartJune = shared.PublicCalendarDatePropertyOperationFiscalYearStartJune

// Equals "MARCH"
const PublicCalendarDatePropertyOperationFiscalYearStartMarch = shared.PublicCalendarDatePropertyOperationFiscalYearStartMarch

// Equals "MAY"
const PublicCalendarDatePropertyOperationFiscalYearStartMay = shared.PublicCalendarDatePropertyOperationFiscalYearStartMay

// Equals "NOVEMBER"
const PublicCalendarDatePropertyOperationFiscalYearStartNovember = shared.PublicCalendarDatePropertyOperationFiscalYearStartNovember

// Equals "OCTOBER"
const PublicCalendarDatePropertyOperationFiscalYearStartOctober = shared.PublicCalendarDatePropertyOperationFiscalYearStartOctober

// Equals "SEPTEMBER"
const PublicCalendarDatePropertyOperationFiscalYearStartSeptember = shared.PublicCalendarDatePropertyOperationFiscalYearStartSeptember

// This is an alias to an internal type.
type PublicCalendarDatePropertyOperationParam = shared.PublicCalendarDatePropertyOperationParam

// This is an alias to an internal type.
type PublicCampaignInfluencedFilter = shared.PublicCampaignInfluencedFilter

// This is an alias to an internal type.
type PublicCampaignInfluencedFilterFilterType = shared.PublicCampaignInfluencedFilterFilterType

// Equals "CAMPAIGN_INFLUENCED"
const PublicCampaignInfluencedFilterFilterTypeCampaignInfluenced = shared.PublicCampaignInfluencedFilterFilterTypeCampaignInfluenced

// This is an alias to an internal type.
type PublicCampaignInfluencedFilterParam = shared.PublicCampaignInfluencedFilterParam

// This is an alias to an internal type.
type PublicCommunicationSubscriptionFilter = shared.PublicCommunicationSubscriptionFilter

// This is an alias to an internal type.
type PublicCommunicationSubscriptionFilterFilterType = shared.PublicCommunicationSubscriptionFilterFilterType

// Equals "COMMUNICATION_SUBSCRIPTION"
const PublicCommunicationSubscriptionFilterFilterTypeCommunicationSubscription = shared.PublicCommunicationSubscriptionFilterFilterTypeCommunicationSubscription

// This is an alias to an internal type.
type PublicCommunicationSubscriptionFilterParam = shared.PublicCommunicationSubscriptionFilterParam

// This is an alias to an internal type.
type PublicComparativeDatePropertyOperation = shared.PublicComparativeDatePropertyOperation

// This is an alias to an internal type.
type PublicComparativeDatePropertyOperationOperationType = shared.PublicComparativeDatePropertyOperationOperationType

// Equals "COMPARATIVE_DATE"
const PublicComparativeDatePropertyOperationOperationTypeComparativeDate = shared.PublicComparativeDatePropertyOperationOperationTypeComparativeDate

// This is an alias to an internal type.
type PublicComparativeDatePropertyOperationParam = shared.PublicComparativeDatePropertyOperationParam

// This is an alias to an internal type.
type PublicComparativePropertyUpdatedOperation = shared.PublicComparativePropertyUpdatedOperation

// This is an alias to an internal type.
type PublicComparativePropertyUpdatedOperationOperationType = shared.PublicComparativePropertyUpdatedOperationOperationType

// Equals "COMPARATIVE_PROPERTY_UPDATED"
const PublicComparativePropertyUpdatedOperationOperationTypeComparativePropertyUpdated = shared.PublicComparativePropertyUpdatedOperationOperationTypeComparativePropertyUpdated

// This is an alias to an internal type.
type PublicComparativePropertyUpdatedOperationParam = shared.PublicComparativePropertyUpdatedOperationParam

// This is an alias to an internal type.
type PublicConstantFilter = shared.PublicConstantFilter

// This is an alias to an internal type.
type PublicConstantFilterFilterType = shared.PublicConstantFilterFilterType

// Equals "CONSTANT"
const PublicConstantFilterFilterTypeConstant = shared.PublicConstantFilterFilterTypeConstant

// This is an alias to an internal type.
type PublicConstantFilterParam = shared.PublicConstantFilterParam

// This is an alias to an internal type.
type PublicCtaAnalyticsFilter = shared.PublicCtaAnalyticsFilter

// This is an alias to an internal type.
type PublicCtaAnalyticsFilterFilterType = shared.PublicCtaAnalyticsFilterFilterType

// Equals "CTA"
const PublicCtaAnalyticsFilterFilterTypeCta = shared.PublicCtaAnalyticsFilterFilterTypeCta

// This is an alias to an internal type.
type PublicCtaAnalyticsFilterCoalescingRefineByUnion = shared.PublicCtaAnalyticsFilterCoalescingRefineByUnion

// This is an alias to an internal type.
type PublicCtaAnalyticsFilterPruningRefineByUnion = shared.PublicCtaAnalyticsFilterPruningRefineByUnion

// This is an alias to an internal type.
type PublicCtaAnalyticsFilterParam = shared.PublicCtaAnalyticsFilterParam

// This is an alias to an internal type.
type PublicCtaAnalyticsFilterCoalescingRefineByUnionParam = shared.PublicCtaAnalyticsFilterCoalescingRefineByUnionParam

// This is an alias to an internal type.
type PublicCtaAnalyticsFilterPruningRefineByUnionParam = shared.PublicCtaAnalyticsFilterPruningRefineByUnionParam

// This is an alias to an internal type.
type PublicDatePoint = shared.PublicDatePoint

// This is an alias to an internal type.
type PublicDatePointTimeType = shared.PublicDatePointTimeType

// Equals "DATE"
const PublicDatePointTimeTypeDate = shared.PublicDatePointTimeTypeDate

// This is an alias to an internal type.
type PublicDatePointParam = shared.PublicDatePointParam

// This is an alias to an internal type.
type PublicDatePropertyOperation = shared.PublicDatePropertyOperation

// This is an alias to an internal type.
type PublicDatePropertyOperationOperationType = shared.PublicDatePropertyOperationOperationType

// Equals "DATE"
const PublicDatePropertyOperationOperationTypeDate = shared.PublicDatePropertyOperationOperationTypeDate

// This is an alias to an internal type.
type PublicDatePropertyOperationParam = shared.PublicDatePropertyOperationParam

// This is an alias to an internal type.
type PublicDateTimePropertyOperation = shared.PublicDateTimePropertyOperation

// This is an alias to an internal type.
type PublicDateTimePropertyOperationOperationType = shared.PublicDateTimePropertyOperationOperationType

// Equals "DATETIME"
const PublicDateTimePropertyOperationOperationTypeDatetime = shared.PublicDateTimePropertyOperationOperationTypeDatetime

// This is an alias to an internal type.
type PublicDateTimePropertyOperationParam = shared.PublicDateTimePropertyOperationParam

// This is an alias to an internal type.
type PublicEmailEventFilter = shared.PublicEmailEventFilter

// This is an alias to an internal type.
type PublicEmailEventFilterFilterType = shared.PublicEmailEventFilterFilterType

// Equals "EMAIL_EVENT"
const PublicEmailEventFilterFilterTypeEmailEvent = shared.PublicEmailEventFilterFilterTypeEmailEvent

// This is an alias to an internal type.
type PublicEmailEventFilterOperator = shared.PublicEmailEventFilterOperator

// Equals "BOUNCED"
const PublicEmailEventFilterOperatorBounced = shared.PublicEmailEventFilterOperatorBounced

// Equals "LINK_CLICKED"
const PublicEmailEventFilterOperatorLinkClicked = shared.PublicEmailEventFilterOperatorLinkClicked

// Equals "MARKED_SPAM"
const PublicEmailEventFilterOperatorMarkedSpam = shared.PublicEmailEventFilterOperatorMarkedSpam

// Equals "OPENED"
const PublicEmailEventFilterOperatorOpened = shared.PublicEmailEventFilterOperatorOpened

// Equals "OPENED_BUT_LINK_NOT_CLICKED"
const PublicEmailEventFilterOperatorOpenedButLinkNotClicked = shared.PublicEmailEventFilterOperatorOpenedButLinkNotClicked

// Equals "OPENED_BUT_NOT_REPLIED"
const PublicEmailEventFilterOperatorOpenedButNotReplied = shared.PublicEmailEventFilterOperatorOpenedButNotReplied

// Equals "RECEIVED"
const PublicEmailEventFilterOperatorReceived = shared.PublicEmailEventFilterOperatorReceived

// Equals "RECEIVED_BUT_NOT_OPENED"
const PublicEmailEventFilterOperatorReceivedButNotOpened = shared.PublicEmailEventFilterOperatorReceivedButNotOpened

// Equals "REPLIED"
const PublicEmailEventFilterOperatorReplied = shared.PublicEmailEventFilterOperatorReplied

// Equals "SENT"
const PublicEmailEventFilterOperatorSent = shared.PublicEmailEventFilterOperatorSent

// Equals "SENT_BUT_LINK_NOT_CLICKED"
const PublicEmailEventFilterOperatorSentButLinkNotClicked = shared.PublicEmailEventFilterOperatorSentButLinkNotClicked

// Equals "SENT_BUT_NOT_RECEIVED"
const PublicEmailEventFilterOperatorSentButNotReceived = shared.PublicEmailEventFilterOperatorSentButNotReceived

// Equals "UNSUBSCRIBED"
const PublicEmailEventFilterOperatorUnsubscribed = shared.PublicEmailEventFilterOperatorUnsubscribed

// This is an alias to an internal type.
type PublicEmailEventFilterPruningRefineByUnion = shared.PublicEmailEventFilterPruningRefineByUnion

// This is an alias to an internal type.
type PublicEmailEventFilterParam = shared.PublicEmailEventFilterParam

// This is an alias to an internal type.
type PublicEmailEventFilterPruningRefineByUnionParam = shared.PublicEmailEventFilterPruningRefineByUnionParam

// This is an alias to an internal type.
type PublicEmailSubscriptionFilter = shared.PublicEmailSubscriptionFilter

// This is an alias to an internal type.
type PublicEmailSubscriptionFilterFilterType = shared.PublicEmailSubscriptionFilterFilterType

// Equals "EMAIL_SUBSCRIPTION"
const PublicEmailSubscriptionFilterFilterTypeEmailSubscription = shared.PublicEmailSubscriptionFilterFilterTypeEmailSubscription

// This is an alias to an internal type.
type PublicEmailSubscriptionFilterParam = shared.PublicEmailSubscriptionFilterParam

// This is an alias to an internal type.
type PublicEnumerationPropertyOperation = shared.PublicEnumerationPropertyOperation

// This is an alias to an internal type.
type PublicEnumerationPropertyOperationOperationType = shared.PublicEnumerationPropertyOperationOperationType

// Equals "ENUMERATION"
const PublicEnumerationPropertyOperationOperationTypeEnumeration = shared.PublicEnumerationPropertyOperationOperationTypeEnumeration

// This is an alias to an internal type.
type PublicEnumerationPropertyOperationParam = shared.PublicEnumerationPropertyOperationParam

// This is an alias to an internal type.
type PublicEventAnalyticsFilter = shared.PublicEventAnalyticsFilter

// This is an alias to an internal type.
type PublicEventAnalyticsFilterFilterType = shared.PublicEventAnalyticsFilterFilterType

// Equals "EVENT"
const PublicEventAnalyticsFilterFilterTypeEvent = shared.PublicEventAnalyticsFilterFilterTypeEvent

// This is an alias to an internal type.
type PublicEventAnalyticsFilterCoalescingRefineByUnion = shared.PublicEventAnalyticsFilterCoalescingRefineByUnion

// This is an alias to an internal type.
type PublicEventAnalyticsFilterPruningRefineByUnion = shared.PublicEventAnalyticsFilterPruningRefineByUnion

// This is an alias to an internal type.
type PublicEventAnalyticsFilterParam = shared.PublicEventAnalyticsFilterParam

// This is an alias to an internal type.
type PublicEventAnalyticsFilterCoalescingRefineByUnionParam = shared.PublicEventAnalyticsFilterCoalescingRefineByUnionParam

// This is an alias to an internal type.
type PublicEventAnalyticsFilterPruningRefineByUnionParam = shared.PublicEventAnalyticsFilterPruningRefineByUnionParam

// This is an alias to an internal type.
type PublicEventFilterMetadata = shared.PublicEventFilterMetadata

// This is an alias to an internal type.
type PublicEventFilterMetadataOperationUnion = shared.PublicEventFilterMetadataOperationUnion

// This is an alias to an internal type.
type PublicEventFilterMetadataParam = shared.PublicEventFilterMetadataParam

// This is an alias to an internal type.
type PublicEventFilterMetadataOperationUnionParam = shared.PublicEventFilterMetadataOperationUnionParam

// This is an alias to an internal type.
type PublicFiscalQuarterReference = shared.PublicFiscalQuarterReference

// This is an alias to an internal type.
type PublicFiscalQuarterReferenceReferenceType = shared.PublicFiscalQuarterReferenceReferenceType

// Equals "FISCAL_QUARTER"
const PublicFiscalQuarterReferenceReferenceTypeFiscalQuarter = shared.PublicFiscalQuarterReferenceReferenceTypeFiscalQuarter

// This is an alias to an internal type.
type PublicFiscalQuarterReferenceParam = shared.PublicFiscalQuarterReferenceParam

// This is an alias to an internal type.
type PublicFiscalYearReference = shared.PublicFiscalYearReference

// This is an alias to an internal type.
type PublicFiscalYearReferenceReferenceType = shared.PublicFiscalYearReferenceReferenceType

// Equals "FISCAL_YEAR"
const PublicFiscalYearReferenceReferenceTypeFiscalYear = shared.PublicFiscalYearReferenceReferenceTypeFiscalYear

// This is an alias to an internal type.
type PublicFiscalYearReferenceParam = shared.PublicFiscalYearReferenceParam

// This is an alias to an internal type.
type PublicFormSubmissionFilter = shared.PublicFormSubmissionFilter

// This is an alias to an internal type.
type PublicFormSubmissionFilterFilterType = shared.PublicFormSubmissionFilterFilterType

// Equals "FORM_SUBMISSION"
const PublicFormSubmissionFilterFilterTypeFormSubmission = shared.PublicFormSubmissionFilterFilterTypeFormSubmission

// This is an alias to an internal type.
type PublicFormSubmissionFilterOperator = shared.PublicFormSubmissionFilterOperator

// Equals "FILLED_OUT"
const PublicFormSubmissionFilterOperatorFilledOut = shared.PublicFormSubmissionFilterOperatorFilledOut

// Equals "NOT_FILLED_OUT"
const PublicFormSubmissionFilterOperatorNotFilledOut = shared.PublicFormSubmissionFilterOperatorNotFilledOut

// This is an alias to an internal type.
type PublicFormSubmissionFilterCoalescingRefineByUnion = shared.PublicFormSubmissionFilterCoalescingRefineByUnion

// This is an alias to an internal type.
type PublicFormSubmissionFilterPruningRefineByUnion = shared.PublicFormSubmissionFilterPruningRefineByUnion

// This is an alias to an internal type.
type PublicFormSubmissionFilterParam = shared.PublicFormSubmissionFilterParam

// This is an alias to an internal type.
type PublicFormSubmissionFilterCoalescingRefineByUnionParam = shared.PublicFormSubmissionFilterCoalescingRefineByUnionParam

// This is an alias to an internal type.
type PublicFormSubmissionFilterPruningRefineByUnionParam = shared.PublicFormSubmissionFilterPruningRefineByUnionParam

// This is an alias to an internal type.
type PublicFormSubmissionOnPageFilter = shared.PublicFormSubmissionOnPageFilter

// This is an alias to an internal type.
type PublicFormSubmissionOnPageFilterFilterType = shared.PublicFormSubmissionOnPageFilterFilterType

// Equals "FORM_SUBMISSION_ON_PAGE"
const PublicFormSubmissionOnPageFilterFilterTypeFormSubmissionOnPage = shared.PublicFormSubmissionOnPageFilterFilterTypeFormSubmissionOnPage

// This is an alias to an internal type.
type PublicFormSubmissionOnPageFilterOperator = shared.PublicFormSubmissionOnPageFilterOperator

// Equals "FILLED_OUT"
const PublicFormSubmissionOnPageFilterOperatorFilledOut = shared.PublicFormSubmissionOnPageFilterOperatorFilledOut

// Equals "NOT_FILLED_OUT"
const PublicFormSubmissionOnPageFilterOperatorNotFilledOut = shared.PublicFormSubmissionOnPageFilterOperatorNotFilledOut

// This is an alias to an internal type.
type PublicFormSubmissionOnPageFilterCoalescingRefineByUnion = shared.PublicFormSubmissionOnPageFilterCoalescingRefineByUnion

// This is an alias to an internal type.
type PublicFormSubmissionOnPageFilterPruningRefineByUnion = shared.PublicFormSubmissionOnPageFilterPruningRefineByUnion

// This is an alias to an internal type.
type PublicFormSubmissionOnPageFilterParam = shared.PublicFormSubmissionOnPageFilterParam

// This is an alias to an internal type.
type PublicFormSubmissionOnPageFilterCoalescingRefineByUnionParam = shared.PublicFormSubmissionOnPageFilterCoalescingRefineByUnionParam

// This is an alias to an internal type.
type PublicFormSubmissionOnPageFilterPruningRefineByUnionParam = shared.PublicFormSubmissionOnPageFilterPruningRefineByUnionParam

// This is an alias to an internal type.
type PublicInListFilter = shared.PublicInListFilter

// This is an alias to an internal type.
type PublicInListFilterFilterType = shared.PublicInListFilterFilterType

// Equals "IN_LIST"
const PublicInListFilterFilterTypeInList = shared.PublicInListFilterFilterTypeInList

// This is an alias to an internal type.
type PublicInListFilterParam = shared.PublicInListFilterParam

// This is an alias to an internal type.
type PublicInListFilterMetadata = shared.PublicInListFilterMetadata

// This is an alias to an internal type.
type PublicInListFilterMetadataParam = shared.PublicInListFilterMetadataParam

// This is an alias to an internal type.
type PublicIndexOffset = shared.PublicIndexOffset

// This is an alias to an internal type.
type PublicIndexOffsetParam = shared.PublicIndexOffsetParam

// This is an alias to an internal type.
type PublicIndexedTimePoint = shared.PublicIndexedTimePoint

// This is an alias to an internal type.
type PublicIndexedTimePointIndexReferenceUnion = shared.PublicIndexedTimePointIndexReferenceUnion

// This is an alias to an internal type.
type PublicIndexedTimePointTimeType = shared.PublicIndexedTimePointTimeType

// Equals "INDEXED"
const PublicIndexedTimePointTimeTypeIndexed = shared.PublicIndexedTimePointTimeTypeIndexed

// This is an alias to an internal type.
type PublicIndexedTimePointParam = shared.PublicIndexedTimePointParam

// This is an alias to an internal type.
type PublicIndexedTimePointIndexReferenceUnionParam = shared.PublicIndexedTimePointIndexReferenceUnionParam

// This is an alias to an internal type.
type PublicIntegrationEventFilter = shared.PublicIntegrationEventFilter

// This is an alias to an internal type.
type PublicIntegrationEventFilterFilterType = shared.PublicIntegrationEventFilterFilterType

// Equals "INTEGRATION_EVENT"
const PublicIntegrationEventFilterFilterTypeIntegrationEvent = shared.PublicIntegrationEventFilterFilterTypeIntegrationEvent

// This is an alias to an internal type.
type PublicIntegrationEventFilterParam = shared.PublicIntegrationEventFilterParam

// This is an alias to an internal type.
type PublicMonthReference = shared.PublicMonthReference

// This is an alias to an internal type.
type PublicMonthReferenceReferenceType = shared.PublicMonthReferenceReferenceType

// Equals "MONTH"
const PublicMonthReferenceReferenceTypeMonth = shared.PublicMonthReferenceReferenceTypeMonth

// This is an alias to an internal type.
type PublicMonthReferenceParam = shared.PublicMonthReferenceParam

// This is an alias to an internal type.
type PublicMultiStringPropertyOperation = shared.PublicMultiStringPropertyOperation

// This is an alias to an internal type.
type PublicMultiStringPropertyOperationOperationType = shared.PublicMultiStringPropertyOperationOperationType

// Equals "MULTISTRING"
const PublicMultiStringPropertyOperationOperationTypeMultistring = shared.PublicMultiStringPropertyOperationOperationTypeMultistring

// This is an alias to an internal type.
type PublicMultiStringPropertyOperationParam = shared.PublicMultiStringPropertyOperationParam

// This is an alias to an internal type.
type PublicNotAllFilterBranch = shared.PublicNotAllFilterBranch

// This is an alias to an internal type.
type PublicNotAllFilterBranchFilterBranchUnion = shared.PublicNotAllFilterBranchFilterBranchUnion

// This is an alias to an internal type.
type PublicNotAllFilterBranchFilterBranchType = shared.PublicNotAllFilterBranchFilterBranchType

// Equals "NOT_ALL"
const PublicNotAllFilterBranchFilterBranchTypeNotAll = shared.PublicNotAllFilterBranchFilterBranchTypeNotAll

// This is an alias to an internal type.
type PublicNotAllFilterBranchFilterUnion = shared.PublicNotAllFilterBranchFilterUnion

// This is an alias to an internal type.
type PublicNotAllFilterBranchParam = shared.PublicNotAllFilterBranchParam

// This is an alias to an internal type.
type PublicNotAllFilterBranchFilterBranchUnionParam = shared.PublicNotAllFilterBranchFilterBranchUnionParam

// This is an alias to an internal type.
type PublicNotAllFilterBranchFilterUnionParam = shared.PublicNotAllFilterBranchFilterUnionParam

// This is an alias to an internal type.
type PublicNotAnyFilterBranch = shared.PublicNotAnyFilterBranch

// This is an alias to an internal type.
type PublicNotAnyFilterBranchFilterBranchUnion = shared.PublicNotAnyFilterBranchFilterBranchUnion

// This is an alias to an internal type.
type PublicNotAnyFilterBranchFilterBranchType = shared.PublicNotAnyFilterBranchFilterBranchType

// Equals "NOT_ANY"
const PublicNotAnyFilterBranchFilterBranchTypeNotAny = shared.PublicNotAnyFilterBranchFilterBranchTypeNotAny

// This is an alias to an internal type.
type PublicNotAnyFilterBranchFilterUnion = shared.PublicNotAnyFilterBranchFilterUnion

// This is an alias to an internal type.
type PublicNotAnyFilterBranchParam = shared.PublicNotAnyFilterBranchParam

// This is an alias to an internal type.
type PublicNotAnyFilterBranchFilterBranchUnionParam = shared.PublicNotAnyFilterBranchFilterBranchUnionParam

// This is an alias to an internal type.
type PublicNotAnyFilterBranchFilterUnionParam = shared.PublicNotAnyFilterBranchFilterUnionParam

// This is an alias to an internal type.
type PublicNowReference = shared.PublicNowReference

// This is an alias to an internal type.
type PublicNowReferenceReferenceType = shared.PublicNowReferenceReferenceType

// Equals "NOW"
const PublicNowReferenceReferenceTypeNow = shared.PublicNowReferenceReferenceTypeNow

// This is an alias to an internal type.
type PublicNowReferenceParam = shared.PublicNowReferenceParam

// This is an alias to an internal type.
type PublicNumAssociationsFilter = shared.PublicNumAssociationsFilter

// This is an alias to an internal type.
type PublicNumAssociationsFilterCoalescingRefineByUnion = shared.PublicNumAssociationsFilterCoalescingRefineByUnion

// This is an alias to an internal type.
type PublicNumAssociationsFilterFilterType = shared.PublicNumAssociationsFilterFilterType

// Equals "NUM_ASSOCIATIONS"
const PublicNumAssociationsFilterFilterTypeNumAssociations = shared.PublicNumAssociationsFilterFilterTypeNumAssociations

// This is an alias to an internal type.
type PublicNumAssociationsFilterParam = shared.PublicNumAssociationsFilterParam

// This is an alias to an internal type.
type PublicNumAssociationsFilterCoalescingRefineByUnionParam = shared.PublicNumAssociationsFilterCoalescingRefineByUnionParam

// This is an alias to an internal type.
type PublicNumOccurrencesRefineBy = shared.PublicNumOccurrencesRefineBy

// This is an alias to an internal type.
type PublicNumOccurrencesRefineByType = shared.PublicNumOccurrencesRefineByType

// Equals "NUM_OCCURRENCES"
const PublicNumOccurrencesRefineByTypeNumOccurrences = shared.PublicNumOccurrencesRefineByTypeNumOccurrences

// This is an alias to an internal type.
type PublicNumOccurrencesRefineByParam = shared.PublicNumOccurrencesRefineByParam

// This is an alias to an internal type.
type PublicNumberPropertyOperation = shared.PublicNumberPropertyOperation

// This is an alias to an internal type.
type PublicNumberPropertyOperationOperationType = shared.PublicNumberPropertyOperationOperationType

// Equals "NUMBER"
const PublicNumberPropertyOperationOperationTypeNumber = shared.PublicNumberPropertyOperationOperationTypeNumber

// This is an alias to an internal type.
type PublicNumberPropertyOperationParam = shared.PublicNumberPropertyOperationParam

// This is an alias to an internal type.
type PublicObjectID = shared.PublicObjectID

// This is an alias to an internal type.
type PublicObjectIDParam = shared.PublicObjectIDParam

// This is an alias to an internal type.
type PublicOrFilterBranch = shared.PublicOrFilterBranch

// This is an alias to an internal type.
type PublicOrFilterBranchFilterBranchUnion = shared.PublicOrFilterBranchFilterBranchUnion

// This is an alias to an internal type.
type PublicOrFilterBranchFilterBranchType = shared.PublicOrFilterBranchFilterBranchType

// Equals "OR"
const PublicOrFilterBranchFilterBranchTypeOr = shared.PublicOrFilterBranchFilterBranchTypeOr

// This is an alias to an internal type.
type PublicOrFilterBranchFilterUnion = shared.PublicOrFilterBranchFilterUnion

// This is an alias to an internal type.
type PublicOrFilterBranchParam = shared.PublicOrFilterBranchParam

// This is an alias to an internal type.
type PublicOrFilterBranchFilterBranchUnionParam = shared.PublicOrFilterBranchFilterBranchUnionParam

// This is an alias to an internal type.
type PublicOrFilterBranchFilterUnionParam = shared.PublicOrFilterBranchFilterUnionParam

// This is an alias to an internal type.
type PublicPageViewAnalyticsFilter = shared.PublicPageViewAnalyticsFilter

// This is an alias to an internal type.
type PublicPageViewAnalyticsFilterFilterType = shared.PublicPageViewAnalyticsFilterFilterType

// Equals "PAGE_VIEW"
const PublicPageViewAnalyticsFilterFilterTypePageView = shared.PublicPageViewAnalyticsFilterFilterTypePageView

// This is an alias to an internal type.
type PublicPageViewAnalyticsFilterCoalescingRefineByUnion = shared.PublicPageViewAnalyticsFilterCoalescingRefineByUnion

// This is an alias to an internal type.
type PublicPageViewAnalyticsFilterPruningRefineByUnion = shared.PublicPageViewAnalyticsFilterPruningRefineByUnion

// This is an alias to an internal type.
type PublicPageViewAnalyticsFilterParam = shared.PublicPageViewAnalyticsFilterParam

// This is an alias to an internal type.
type PublicPageViewAnalyticsFilterCoalescingRefineByUnionParam = shared.PublicPageViewAnalyticsFilterCoalescingRefineByUnionParam

// This is an alias to an internal type.
type PublicPageViewAnalyticsFilterPruningRefineByUnionParam = shared.PublicPageViewAnalyticsFilterPruningRefineByUnionParam

// This is an alias to an internal type.
type PublicPrivacyAnalyticsFilter = shared.PublicPrivacyAnalyticsFilter

// This is an alias to an internal type.
type PublicPrivacyAnalyticsFilterFilterType = shared.PublicPrivacyAnalyticsFilterFilterType

// Equals "PRIVACY"
const PublicPrivacyAnalyticsFilterFilterTypePrivacy = shared.PublicPrivacyAnalyticsFilterFilterTypePrivacy

// This is an alias to an internal type.
type PublicPrivacyAnalyticsFilterParam = shared.PublicPrivacyAnalyticsFilterParam

// This is an alias to an internal type.
type PublicPropertyAssociationFilterBranch = shared.PublicPropertyAssociationFilterBranch

// This is an alias to an internal type.
type PublicPropertyAssociationFilterBranchFilterBranchUnion = shared.PublicPropertyAssociationFilterBranchFilterBranchUnion

// This is an alias to an internal type.
type PublicPropertyAssociationFilterBranchFilterBranchType = shared.PublicPropertyAssociationFilterBranchFilterBranchType

// Equals "PROPERTY_ASSOCIATION"
const PublicPropertyAssociationFilterBranchFilterBranchTypePropertyAssociation = shared.PublicPropertyAssociationFilterBranchFilterBranchTypePropertyAssociation

// This is an alias to an internal type.
type PublicPropertyAssociationFilterBranchFilterUnion = shared.PublicPropertyAssociationFilterBranchFilterUnion

// This is an alias to an internal type.
type PublicPropertyAssociationFilterBranchParam = shared.PublicPropertyAssociationFilterBranchParam

// This is an alias to an internal type.
type PublicPropertyAssociationFilterBranchFilterBranchUnionParam = shared.PublicPropertyAssociationFilterBranchFilterBranchUnionParam

// This is an alias to an internal type.
type PublicPropertyAssociationFilterBranchFilterUnionParam = shared.PublicPropertyAssociationFilterBranchFilterUnionParam

// This is an alias to an internal type.
type PublicPropertyAssociationInListFilter = shared.PublicPropertyAssociationInListFilter

// This is an alias to an internal type.
type PublicPropertyAssociationInListFilterCoalescingRefineByUnion = shared.PublicPropertyAssociationInListFilterCoalescingRefineByUnion

// This is an alias to an internal type.
type PublicPropertyAssociationInListFilterFilterType = shared.PublicPropertyAssociationInListFilterFilterType

// Equals "PROPERTY_ASSOCIATION"
const PublicPropertyAssociationInListFilterFilterTypePropertyAssociation = shared.PublicPropertyAssociationInListFilterFilterTypePropertyAssociation

// This is an alias to an internal type.
type PublicPropertyAssociationInListFilterParam = shared.PublicPropertyAssociationInListFilterParam

// This is an alias to an internal type.
type PublicPropertyAssociationInListFilterCoalescingRefineByUnionParam = shared.PublicPropertyAssociationInListFilterCoalescingRefineByUnionParam

// This is an alias to an internal type.
type PublicPropertyFilter = shared.PublicPropertyFilter

// This is an alias to an internal type.
type PublicPropertyFilterFilterType = shared.PublicPropertyFilterFilterType

// Equals "PROPERTY"
const PublicPropertyFilterFilterTypeProperty = shared.PublicPropertyFilterFilterTypeProperty

// This is an alias to an internal type.
type PublicPropertyFilterOperationUnion = shared.PublicPropertyFilterOperationUnion

// This is an alias to an internal type.
type PublicPropertyFilterParam = shared.PublicPropertyFilterParam

// This is an alias to an internal type.
type PublicPropertyFilterOperationUnionParam = shared.PublicPropertyFilterOperationUnionParam

// This is an alias to an internal type.
type PublicPropertyReferencedTime = shared.PublicPropertyReferencedTime

// This is an alias to an internal type.
type PublicPropertyReferencedTimeTimeType = shared.PublicPropertyReferencedTimeTimeType

// Equals "PROPERTY_REFERENCED"
const PublicPropertyReferencedTimeTimeTypePropertyReferenced = shared.PublicPropertyReferencedTimeTimeTypePropertyReferenced

// This is an alias to an internal type.
type PublicPropertyReferencedTimeParam = shared.PublicPropertyReferencedTimeParam

// This is an alias to an internal type.
type PublicQuarterReference = shared.PublicQuarterReference

// This is an alias to an internal type.
type PublicQuarterReferenceReferenceType = shared.PublicQuarterReferenceReferenceType

// Equals "QUARTER"
const PublicQuarterReferenceReferenceTypeQuarter = shared.PublicQuarterReferenceReferenceTypeQuarter

// This is an alias to an internal type.
type PublicQuarterReferenceParam = shared.PublicQuarterReferenceParam

// This is an alias to an internal type.
type PublicRangedDatePropertyOperation = shared.PublicRangedDatePropertyOperation

// This is an alias to an internal type.
type PublicRangedDatePropertyOperationOperationType = shared.PublicRangedDatePropertyOperationOperationType

// Equals "RANGED_DATE"
const PublicRangedDatePropertyOperationOperationTypeRangedDate = shared.PublicRangedDatePropertyOperationOperationTypeRangedDate

// This is an alias to an internal type.
type PublicRangedDatePropertyOperationParam = shared.PublicRangedDatePropertyOperationParam

// This is an alias to an internal type.
type PublicRangedNumberPropertyOperation = shared.PublicRangedNumberPropertyOperation

// This is an alias to an internal type.
type PublicRangedNumberPropertyOperationOperationType = shared.PublicRangedNumberPropertyOperationOperationType

// Equals "NUMBER_RANGED"
const PublicRangedNumberPropertyOperationOperationTypeNumberRanged = shared.PublicRangedNumberPropertyOperationOperationTypeNumberRanged

// This is an alias to an internal type.
type PublicRangedNumberPropertyOperationParam = shared.PublicRangedNumberPropertyOperationParam

// This is an alias to an internal type.
type PublicRangedTimeOperation = shared.PublicRangedTimeOperation

// This is an alias to an internal type.
type PublicRangedTimeOperationLowerBoundTimePointUnion = shared.PublicRangedTimeOperationLowerBoundTimePointUnion

// This is an alias to an internal type.
type PublicRangedTimeOperationType = shared.PublicRangedTimeOperationType

// Equals "TIME_RANGED"
const PublicRangedTimeOperationTypeTimeRanged = shared.PublicRangedTimeOperationTypeTimeRanged

// This is an alias to an internal type.
type PublicRangedTimeOperationUpperBoundTimePointUnion = shared.PublicRangedTimeOperationUpperBoundTimePointUnion

// This is an alias to an internal type.
type PublicRangedTimeOperationParam = shared.PublicRangedTimeOperationParam

// This is an alias to an internal type.
type PublicRangedTimeOperationLowerBoundTimePointUnionParam = shared.PublicRangedTimeOperationLowerBoundTimePointUnionParam

// This is an alias to an internal type.
type PublicRangedTimeOperationUpperBoundTimePointUnionParam = shared.PublicRangedTimeOperationUpperBoundTimePointUnionParam

// This is an alias to an internal type.
type PublicRelativeComparativeTimestampRefineBy = shared.PublicRelativeComparativeTimestampRefineBy

// This is an alias to an internal type.
type PublicRelativeComparativeTimestampRefineByType = shared.PublicRelativeComparativeTimestampRefineByType

// Equals "RELATIVE_COMPARATIVE"
const PublicRelativeComparativeTimestampRefineByTypeRelativeComparative = shared.PublicRelativeComparativeTimestampRefineByTypeRelativeComparative

// This is an alias to an internal type.
type PublicRelativeComparativeTimestampRefineByParam = shared.PublicRelativeComparativeTimestampRefineByParam

// This is an alias to an internal type.
type PublicRelativeRangedTimestampRefineBy = shared.PublicRelativeRangedTimestampRefineBy

// This is an alias to an internal type.
type PublicRelativeRangedTimestampRefineByType = shared.PublicRelativeRangedTimestampRefineByType

// Equals "RELATIVE_RANGED"
const PublicRelativeRangedTimestampRefineByTypeRelativeRanged = shared.PublicRelativeRangedTimestampRefineByTypeRelativeRanged

// This is an alias to an internal type.
type PublicRelativeRangedTimestampRefineByParam = shared.PublicRelativeRangedTimestampRefineByParam

// This is an alias to an internal type.
type PublicRestrictedFilterBranch = shared.PublicRestrictedFilterBranch

// This is an alias to an internal type.
type PublicRestrictedFilterBranchFilterBranchUnion = shared.PublicRestrictedFilterBranchFilterBranchUnion

// This is an alias to an internal type.
type PublicRestrictedFilterBranchFilterBranchType = shared.PublicRestrictedFilterBranchFilterBranchType

// Equals "RESTRICTED"
const PublicRestrictedFilterBranchFilterBranchTypeRestricted = shared.PublicRestrictedFilterBranchFilterBranchTypeRestricted

// This is an alias to an internal type.
type PublicRestrictedFilterBranchFilterUnion = shared.PublicRestrictedFilterBranchFilterUnion

// This is an alias to an internal type.
type PublicRestrictedFilterBranchParam = shared.PublicRestrictedFilterBranchParam

// This is an alias to an internal type.
type PublicRestrictedFilterBranchFilterBranchUnionParam = shared.PublicRestrictedFilterBranchFilterBranchUnionParam

// This is an alias to an internal type.
type PublicRestrictedFilterBranchFilterUnionParam = shared.PublicRestrictedFilterBranchFilterUnionParam

// This is an alias to an internal type.
type PublicRollingDateRangePropertyOperation = shared.PublicRollingDateRangePropertyOperation

// This is an alias to an internal type.
type PublicRollingDateRangePropertyOperationOperationType = shared.PublicRollingDateRangePropertyOperationOperationType

// Equals "ROLLING_DATE_RANGE"
const PublicRollingDateRangePropertyOperationOperationTypeRollingDateRange = shared.PublicRollingDateRangePropertyOperationOperationTypeRollingDateRange

// This is an alias to an internal type.
type PublicRollingDateRangePropertyOperationParam = shared.PublicRollingDateRangePropertyOperationParam

// This is an alias to an internal type.
type PublicRollingPropertyUpdatedOperation = shared.PublicRollingPropertyUpdatedOperation

// This is an alias to an internal type.
type PublicRollingPropertyUpdatedOperationOperationType = shared.PublicRollingPropertyUpdatedOperationOperationType

// Equals "ROLLING_PROPERTY_UPDATED"
const PublicRollingPropertyUpdatedOperationOperationTypeRollingPropertyUpdated = shared.PublicRollingPropertyUpdatedOperationOperationTypeRollingPropertyUpdated

// This is an alias to an internal type.
type PublicRollingPropertyUpdatedOperationParam = shared.PublicRollingPropertyUpdatedOperationParam

// This is an alias to an internal type.
type PublicSetOccurrencesRefineBy = shared.PublicSetOccurrencesRefineBy

// This is an alias to an internal type.
type PublicSetOccurrencesRefineByType = shared.PublicSetOccurrencesRefineByType

// Equals "SET_OCCURRENCES"
const PublicSetOccurrencesRefineByTypeSetOccurrences = shared.PublicSetOccurrencesRefineByTypeSetOccurrences

// This is an alias to an internal type.
type PublicSetOccurrencesRefineByParam = shared.PublicSetOccurrencesRefineByParam

// This is an alias to an internal type.
type PublicStringPropertyOperation = shared.PublicStringPropertyOperation

// This is an alias to an internal type.
type PublicStringPropertyOperationOperationType = shared.PublicStringPropertyOperationOperationType

// Equals "STRING"
const PublicStringPropertyOperationOperationTypeString = shared.PublicStringPropertyOperationOperationTypeString

// This is an alias to an internal type.
type PublicStringPropertyOperationParam = shared.PublicStringPropertyOperationParam

// This is an alias to an internal type.
type PublicSurveyMonkeyFilter = shared.PublicSurveyMonkeyFilter

// This is an alias to an internal type.
type PublicSurveyMonkeyFilterFilterType = shared.PublicSurveyMonkeyFilterFilterType

// Equals "SURVEY_MONKEY"
const PublicSurveyMonkeyFilterFilterTypeSurveyMonkey = shared.PublicSurveyMonkeyFilterFilterTypeSurveyMonkey

// This is an alias to an internal type.
type PublicSurveyMonkeyFilterParam = shared.PublicSurveyMonkeyFilterParam

// This is an alias to an internal type.
type PublicSurveyMonkeyValueFilter = shared.PublicSurveyMonkeyValueFilter

// This is an alias to an internal type.
type PublicSurveyMonkeyValueFilterFilterType = shared.PublicSurveyMonkeyValueFilterFilterType

// Equals "SURVEY_MONKEY_VALUE"
const PublicSurveyMonkeyValueFilterFilterTypeSurveyMonkeyValue = shared.PublicSurveyMonkeyValueFilterFilterTypeSurveyMonkeyValue

// This is an alias to an internal type.
type PublicSurveyMonkeyValueFilterValueComparisonUnion = shared.PublicSurveyMonkeyValueFilterValueComparisonUnion

// This is an alias to an internal type.
type PublicSurveyMonkeyValueFilterParam = shared.PublicSurveyMonkeyValueFilterParam

// This is an alias to an internal type.
type PublicSurveyMonkeyValueFilterValueComparisonUnionParam = shared.PublicSurveyMonkeyValueFilterValueComparisonUnionParam

// This is an alias to an internal type.
type PublicTimeOffset = shared.PublicTimeOffset

// This is an alias to an internal type.
type PublicTimeOffsetParam = shared.PublicTimeOffsetParam

// This is an alias to an internal type.
type PublicTimePointOperation = shared.PublicTimePointOperation

// This is an alias to an internal type.
type PublicTimePointOperationOperationType = shared.PublicTimePointOperationOperationType

// Equals "TIME_POINT"
const PublicTimePointOperationOperationTypeTimePoint = shared.PublicTimePointOperationOperationTypeTimePoint

// This is an alias to an internal type.
type PublicTimePointOperationTimePointUnion = shared.PublicTimePointOperationTimePointUnion

// This is an alias to an internal type.
type PublicTimePointOperationParam = shared.PublicTimePointOperationParam

// This is an alias to an internal type.
type PublicTimePointOperationTimePointUnionParam = shared.PublicTimePointOperationTimePointUnionParam

// This is an alias to an internal type.
type PublicTodayReference = shared.PublicTodayReference

// This is an alias to an internal type.
type PublicTodayReferenceReferenceType = shared.PublicTodayReferenceReferenceType

// Equals "TODAY"
const PublicTodayReferenceReferenceTypeToday = shared.PublicTodayReferenceReferenceTypeToday

// This is an alias to an internal type.
type PublicTodayReferenceParam = shared.PublicTodayReferenceParam

// This is an alias to an internal type.
type PublicUnifiedEventsFilter = shared.PublicUnifiedEventsFilter

// This is an alias to an internal type.
type PublicUnifiedEventsFilterFilterType = shared.PublicUnifiedEventsFilterFilterType

// Equals "UNIFIED_EVENTS"
const PublicUnifiedEventsFilterFilterTypeUnifiedEvents = shared.PublicUnifiedEventsFilterFilterTypeUnifiedEvents

// This is an alias to an internal type.
type PublicUnifiedEventsFilterCoalescingRefineByUnion = shared.PublicUnifiedEventsFilterCoalescingRefineByUnion

// This is an alias to an internal type.
type PublicUnifiedEventsFilterPruningRefineByUnion = shared.PublicUnifiedEventsFilterPruningRefineByUnion

// This is an alias to an internal type.
type PublicUnifiedEventsFilterParam = shared.PublicUnifiedEventsFilterParam

// This is an alias to an internal type.
type PublicUnifiedEventsFilterCoalescingRefineByUnionParam = shared.PublicUnifiedEventsFilterCoalescingRefineByUnionParam

// This is an alias to an internal type.
type PublicUnifiedEventsFilterPruningRefineByUnionParam = shared.PublicUnifiedEventsFilterPruningRefineByUnionParam

// This is an alias to an internal type.
type PublicUnifiedEventsFilterBranch = shared.PublicUnifiedEventsFilterBranch

// This is an alias to an internal type.
type PublicUnifiedEventsFilterBranchFilterBranchUnion = shared.PublicUnifiedEventsFilterBranchFilterBranchUnion

// This is an alias to an internal type.
type PublicUnifiedEventsFilterBranchFilterBranchType = shared.PublicUnifiedEventsFilterBranchFilterBranchType

// Equals "UNIFIED_EVENTS"
const PublicUnifiedEventsFilterBranchFilterBranchTypeUnifiedEvents = shared.PublicUnifiedEventsFilterBranchFilterBranchTypeUnifiedEvents

// This is an alias to an internal type.
type PublicUnifiedEventsFilterBranchFilterUnion = shared.PublicUnifiedEventsFilterBranchFilterUnion

// This is an alias to an internal type.
type PublicUnifiedEventsFilterBranchOperator = shared.PublicUnifiedEventsFilterBranchOperator

// Equals "HAS_COMPLETED"
const PublicUnifiedEventsFilterBranchOperatorHasCompleted = shared.PublicUnifiedEventsFilterBranchOperatorHasCompleted

// Equals "HAS_NOT_COMPLETED"
const PublicUnifiedEventsFilterBranchOperatorHasNotCompleted = shared.PublicUnifiedEventsFilterBranchOperatorHasNotCompleted

// This is an alias to an internal type.
type PublicUnifiedEventsFilterBranchCoalescingRefineByUnion = shared.PublicUnifiedEventsFilterBranchCoalescingRefineByUnion

// This is an alias to an internal type.
type PublicUnifiedEventsFilterBranchParam = shared.PublicUnifiedEventsFilterBranchParam

// This is an alias to an internal type.
type PublicUnifiedEventsFilterBranchFilterBranchUnionParam = shared.PublicUnifiedEventsFilterBranchFilterBranchUnionParam

// This is an alias to an internal type.
type PublicUnifiedEventsFilterBranchFilterUnionParam = shared.PublicUnifiedEventsFilterBranchFilterUnionParam

// This is an alias to an internal type.
type PublicUnifiedEventsFilterBranchCoalescingRefineByUnionParam = shared.PublicUnifiedEventsFilterBranchCoalescingRefineByUnionParam

// This is an alias to an internal type.
type PublicWebinarFilter = shared.PublicWebinarFilter

// This is an alias to an internal type.
type PublicWebinarFilterFilterType = shared.PublicWebinarFilterFilterType

// Equals "WEBINAR"
const PublicWebinarFilterFilterTypeWebinar = shared.PublicWebinarFilterFilterTypeWebinar

// This is an alias to an internal type.
type PublicWebinarFilterParam = shared.PublicWebinarFilterParam

// This is an alias to an internal type.
type PublicWeekReference = shared.PublicWeekReference

// This is an alias to an internal type.
type PublicWeekReferenceDayOfWeek = shared.PublicWeekReferenceDayOfWeek

// Equals "FRIDAY"
const PublicWeekReferenceDayOfWeekFriday = shared.PublicWeekReferenceDayOfWeekFriday

// Equals "MONDAY"
const PublicWeekReferenceDayOfWeekMonday = shared.PublicWeekReferenceDayOfWeekMonday

// Equals "SATURDAY"
const PublicWeekReferenceDayOfWeekSaturday = shared.PublicWeekReferenceDayOfWeekSaturday

// Equals "SUNDAY"
const PublicWeekReferenceDayOfWeekSunday = shared.PublicWeekReferenceDayOfWeekSunday

// Equals "THURSDAY"
const PublicWeekReferenceDayOfWeekThursday = shared.PublicWeekReferenceDayOfWeekThursday

// Equals "TUESDAY"
const PublicWeekReferenceDayOfWeekTuesday = shared.PublicWeekReferenceDayOfWeekTuesday

// Equals "WEDNESDAY"
const PublicWeekReferenceDayOfWeekWednesday = shared.PublicWeekReferenceDayOfWeekWednesday

// This is an alias to an internal type.
type PublicWeekReferenceReferenceType = shared.PublicWeekReferenceReferenceType

// Equals "WEEK"
const PublicWeekReferenceReferenceTypeWeek = shared.PublicWeekReferenceReferenceTypeWeek

// This is an alias to an internal type.
type PublicWeekReferenceParam = shared.PublicWeekReferenceParam

// This is an alias to an internal type.
type PublicYearReference = shared.PublicYearReference

// This is an alias to an internal type.
type PublicYearReferenceReferenceType = shared.PublicYearReferenceReferenceType

// Equals "YEAR"
const PublicYearReferenceReferenceTypeYear = shared.PublicYearReferenceReferenceTypeYear

// This is an alias to an internal type.
type PublicYearReferenceParam = shared.PublicYearReferenceParam

// Ye olde error
//
// This is an alias to an internal type.
type StandardError = shared.StandardError

// This is an alias to an internal type.
type TaskLocator = shared.TaskLocator

// Model definition for a version user. Contains addition information about the
// user who created a version.
//
// This is an alias to an internal type.
type VersionUser = shared.VersionUser
