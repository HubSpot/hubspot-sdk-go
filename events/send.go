// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package events

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/stainless-sdks/hubspot-sdk-go/internal/apijson"
	"github.com/stainless-sdks/hubspot-sdk-go/internal/apiquery"
	shimjson "github.com/stainless-sdks/hubspot-sdk-go/internal/encoding/json"
	"github.com/stainless-sdks/hubspot-sdk-go/internal/requestconfig"
	"github.com/stainless-sdks/hubspot-sdk-go/option"
	"github.com/stainless-sdks/hubspot-sdk-go/packages/pagination"
	"github.com/stainless-sdks/hubspot-sdk-go/packages/param"
	"github.com/stainless-sdks/hubspot-sdk-go/packages/respjson"
	"github.com/stainless-sdks/hubspot-sdk-go/shared"
)

// SendService contains methods and other services that help with interacting with
// the hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSendService] method instead.
type SendService struct {
	Options []option.RequestOption
}

// NewSendService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewSendService(opts ...option.RequestOption) (r SendService) {
	r = SendService{}
	r.Options = opts
	return
}

func (r *SendService) NewEventDefinition(ctx context.Context, body SendNewEventDefinitionParams, opts ...option.RequestOption) (res *ExternalBehavioralEventTypeDefinition, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "events/custom/2026-03/event-definitions"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

func (r *SendService) NewEventDefinitionProperty(ctx context.Context, eventName string, body SendNewEventDefinitionPropertyParams, opts ...option.RequestOption) (res *Property, err error) {
	opts = slices.Concat(r.Options, opts)
	if eventName == "" {
		err = errors.New("missing required eventName parameter")
		return nil, err
	}
	path := fmt.Sprintf("events/custom/2026-03/event-definitions/%s/property", eventName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

func (r *SendService) DeleteEventDefinition(ctx context.Context, eventName string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if eventName == "" {
		err = errors.New("missing required eventName parameter")
		return err
	}
	path := fmt.Sprintf("events/custom/2026-03/event-definitions/%s", eventName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

func (r *SendService) DeleteEventDefinitionProperty(ctx context.Context, propertyName string, body SendDeleteEventDefinitionPropertyParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if body.EventName == "" {
		err = errors.New("missing required eventName parameter")
		return err
	}
	if propertyName == "" {
		err = errors.New("missing required propertyName parameter")
		return err
	}
	path := fmt.Sprintf("events/custom/2026-03/event-definitions/%s/property/%s", body.EventName, propertyName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

func (r *SendService) GetEventDefinition(ctx context.Context, eventName string, opts ...option.RequestOption) (res *ExternalBehavioralEventTypeDefinition, err error) {
	opts = slices.Concat(r.Options, opts)
	if eventName == "" {
		err = errors.New("missing required eventName parameter")
		return nil, err
	}
	path := fmt.Sprintf("events/custom/2026-03/event-definitions/%s", eventName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

func (r *SendService) ListEventDefinitions(ctx context.Context, query SendListEventDefinitionsParams, opts ...option.RequestOption) (res *pagination.Page[ExternalBehavioralEventTypeDefinition], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "events/custom/2026-03/event-definitions"
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

func (r *SendService) ListEventDefinitionsAutoPaging(ctx context.Context, query SendListEventDefinitionsParams, opts ...option.RequestOption) *pagination.PageAutoPager[ExternalBehavioralEventTypeDefinition] {
	return pagination.NewPageAutoPager(r.ListEventDefinitions(ctx, query, opts...))
}

func (r *SendService) SendEvent(ctx context.Context, body SendSendEventParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "events/custom/2026-03/send"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return err
}

func (r *SendService) SendEventBatch(ctx context.Context, body SendSendEventBatchParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "events/custom/2026-03/send/batch"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return err
}

func (r *SendService) UpdateEventDefinition(ctx context.Context, eventName string, body SendUpdateEventDefinitionParams, opts ...option.RequestOption) (res *ExternalBehavioralEventTypeDefinition, err error) {
	opts = slices.Concat(r.Options, opts)
	if eventName == "" {
		err = errors.New("missing required eventName parameter")
		return nil, err
	}
	path := fmt.Sprintf("events/custom/2026-03/event-definitions/%s", eventName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

func (r *SendService) UpdateEventDefinitionProperty(ctx context.Context, propertyName string, params SendUpdateEventDefinitionPropertyParams, opts ...option.RequestOption) (res *Property, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.EventName == "" {
		err = errors.New("missing required eventName parameter")
		return nil, err
	}
	if propertyName == "" {
		err = errors.New("missing required propertyName parameter")
		return nil, err
	}
	path := fmt.Sprintf("events/custom/2026-03/event-definitions/%s/property/%s", params.EventName, propertyName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return res, err
}

type AbsoluteComparativeTimestampRefineBy struct {
	// Any of "AFTER", "BEFORE".
	Comparison AbsoluteComparativeTimestampRefineByComparison `json:"comparison" api:"required"`
	Timestamp  int64                                          `json:"timestamp" api:"required"`
	// Any of "AbsoluteComparativeTimestampRefineBy".
	Type AbsoluteComparativeTimestampRefineByType `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Comparison  respjson.Field
		Timestamp   respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AbsoluteComparativeTimestampRefineBy) RawJSON() string { return r.JSON.raw }
func (r *AbsoluteComparativeTimestampRefineBy) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AbsoluteComparativeTimestampRefineByComparison string

const (
	AbsoluteComparativeTimestampRefineByComparisonAfter  AbsoluteComparativeTimestampRefineByComparison = "AFTER"
	AbsoluteComparativeTimestampRefineByComparisonBefore AbsoluteComparativeTimestampRefineByComparison = "BEFORE"
)

type AbsoluteComparativeTimestampRefineByType string

const (
	AbsoluteComparativeTimestampRefineByTypeAbsoluteComparativeTimestampRefineBy AbsoluteComparativeTimestampRefineByType = "AbsoluteComparativeTimestampRefineBy"
)

type AbsoluteRangedTimestampRefineBy struct {
	LowerTimestamp int64 `json:"lowerTimestamp" api:"required"`
	// Any of "BETWEEN", "NOT_BETWEEN".
	RangeType AbsoluteRangedTimestampRefineByRangeType `json:"rangeType" api:"required"`
	// Any of "AbsoluteRangedTimestampRefineBy".
	Type           AbsoluteRangedTimestampRefineByType `json:"type" api:"required"`
	UpperTimestamp int64                               `json:"upperTimestamp" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		LowerTimestamp respjson.Field
		RangeType      respjson.Field
		Type           respjson.Field
		UpperTimestamp respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AbsoluteRangedTimestampRefineBy) RawJSON() string { return r.JSON.raw }
func (r *AbsoluteRangedTimestampRefineBy) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AbsoluteRangedTimestampRefineByRangeType string

const (
	AbsoluteRangedTimestampRefineByRangeTypeBetween    AbsoluteRangedTimestampRefineByRangeType = "BETWEEN"
	AbsoluteRangedTimestampRefineByRangeTypeNotBetween AbsoluteRangedTimestampRefineByRangeType = "NOT_BETWEEN"
)

type AbsoluteRangedTimestampRefineByType string

const (
	AbsoluteRangedTimestampRefineByTypeAbsoluteRangedTimestampRefineBy AbsoluteRangedTimestampRefineByType = "AbsoluteRangedTimestampRefineBy"
)

type AllHistoryRefineBy struct {
	// Any of "AllHistoryRefineBy".
	Type AllHistoryRefineByType `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AllHistoryRefineBy) RawJSON() string { return r.JSON.raw }
func (r *AllHistoryRefineBy) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AllHistoryRefineByType string

const (
	AllHistoryRefineByTypeAllHistoryRefineBy AllHistoryRefineByType = "AllHistoryRefineBy"
)

type AllPropertyTypesOperation struct {
	CoalescingRefineBy           AllPropertyTypesOperationCoalescingRefineByUnion `json:"coalescingRefineBy" api:"required"`
	IncludeObjectsWithNoValueSet bool                                             `json:"includeObjectsWithNoValueSet" api:"required"`
	OperationType                string                                           `json:"operationType" api:"required"`
	// Any of "IS_BLANK", "IS_KNOWN", "IS_NOT_BLANK", "IS_UNKNOWN".
	Operator     AllPropertyTypesOperationOperator `json:"operator" api:"required"`
	OperatorName string                            `json:"operatorName" api:"required"`
	// Any of "alltypes".
	PropertyType    AllPropertyTypesOperationPropertyType         `json:"propertyType" api:"required"`
	DefaultValue    string                                        `json:"defaultValue"`
	PruningRefineBy AllPropertyTypesOperationPruningRefineByUnion `json:"pruningRefineBy"`
	RenderSpec      string                                        `json:"renderSpec"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CoalescingRefineBy           respjson.Field
		IncludeObjectsWithNoValueSet respjson.Field
		OperationType                respjson.Field
		Operator                     respjson.Field
		OperatorName                 respjson.Field
		PropertyType                 respjson.Field
		DefaultValue                 respjson.Field
		PruningRefineBy              respjson.Field
		RenderSpec                   respjson.Field
		ExtraFields                  map[string]respjson.Field
		raw                          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AllPropertyTypesOperation) RawJSON() string { return r.JSON.raw }
func (r *AllPropertyTypesOperation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AllPropertyTypesOperationCoalescingRefineByUnion contains all possible
// properties and values from [NumOccurrencesRefineBy], [SetOccurrencesRefineBy].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type AllPropertyTypesOperationCoalescingRefineByUnion struct {
	Type string `json:"type"`
	// This field is from variant [NumOccurrencesRefineBy].
	MaxOccurrences int64 `json:"maxOccurrences"`
	// This field is from variant [NumOccurrencesRefineBy].
	MinOccurrences int64 `json:"minOccurrences"`
	// This field is from variant [SetOccurrencesRefineBy].
	SetType SetOccurrencesRefineBySetType `json:"setType"`
	JSON    struct {
		Type           respjson.Field
		MaxOccurrences respjson.Field
		MinOccurrences respjson.Field
		SetType        respjson.Field
		raw            string
	} `json:"-"`
}

func (u AllPropertyTypesOperationCoalescingRefineByUnion) AsNumOccurrencesRefineBy() (v NumOccurrencesRefineBy) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AllPropertyTypesOperationCoalescingRefineByUnion) AsSetOccurrencesRefineBy() (v SetOccurrencesRefineBy) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AllPropertyTypesOperationCoalescingRefineByUnion) RawJSON() string { return u.JSON.raw }

func (r *AllPropertyTypesOperationCoalescingRefineByUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AllPropertyTypesOperationOperator string

const (
	AllPropertyTypesOperationOperatorIsBlank    AllPropertyTypesOperationOperator = "IS_BLANK"
	AllPropertyTypesOperationOperatorIsKnown    AllPropertyTypesOperationOperator = "IS_KNOWN"
	AllPropertyTypesOperationOperatorIsNotBlank AllPropertyTypesOperationOperator = "IS_NOT_BLANK"
	AllPropertyTypesOperationOperatorIsUnknown  AllPropertyTypesOperationOperator = "IS_UNKNOWN"
)

type AllPropertyTypesOperationPropertyType string

const (
	AllPropertyTypesOperationPropertyTypeAlltypes AllPropertyTypesOperationPropertyType = "alltypes"
)

// AllPropertyTypesOperationPruningRefineByUnion contains all possible properties
// and values from [RelativeComparativeTimestampRefineBy],
// [RelativeRangedTimestampRefineBy], [AbsoluteComparativeTimestampRefineBy],
// [AbsoluteRangedTimestampRefineBy], [AllHistoryRefineBy], [TimePointOperation],
// [RangedTimeOperation].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type AllPropertyTypesOperationPruningRefineByUnion struct {
	Comparison string `json:"comparison"`
	// This field is from variant [RelativeComparativeTimestampRefineBy].
	TimeOffset TimeOffset `json:"timeOffset"`
	Type       string     `json:"type"`
	// This field is from variant [RelativeRangedTimestampRefineBy].
	LowerBoundOffset TimeOffset `json:"lowerBoundOffset"`
	RangeType        string     `json:"rangeType"`
	// This field is from variant [RelativeRangedTimestampRefineBy].
	UpperBoundOffset TimeOffset `json:"upperBoundOffset"`
	// This field is from variant [AbsoluteComparativeTimestampRefineBy].
	Timestamp int64 `json:"timestamp"`
	// This field is from variant [AbsoluteRangedTimestampRefineBy].
	LowerTimestamp int64 `json:"lowerTimestamp"`
	// This field is from variant [AbsoluteRangedTimestampRefineBy].
	UpperTimestamp int64 `json:"upperTimestamp"`
	// This field is from variant [TimePointOperation].
	EndpointBehavior             TimePointOperationEndpointBehavior `json:"endpointBehavior"`
	IncludeObjectsWithNoValueSet bool                               `json:"includeObjectsWithNoValueSet"`
	OperationType                string                             `json:"operationType"`
	Operator                     string                             `json:"operator"`
	OperatorName                 string                             `json:"operatorName"`
	PropertyParser               string                             `json:"propertyParser"`
	PropertyType                 string                             `json:"propertyType"`
	// This field is from variant [TimePointOperation].
	TimePoint    TimePointOperationTimePointUnion `json:"timePoint"`
	DefaultValue string                           `json:"defaultValue"`
	RenderSpec   string                           `json:"renderSpec"`
	// This field is from variant [RangedTimeOperation].
	LowerBoundEndpointBehavior RangedTimeOperationLowerBoundEndpointBehavior `json:"lowerBoundEndpointBehavior"`
	// This field is from variant [RangedTimeOperation].
	LowerBoundTimePoint RangedTimeOperationLowerBoundTimePointUnion `json:"lowerBoundTimePoint"`
	// This field is from variant [RangedTimeOperation].
	UpperBoundEndpointBehavior RangedTimeOperationUpperBoundEndpointBehavior `json:"upperBoundEndpointBehavior"`
	// This field is from variant [RangedTimeOperation].
	UpperBoundTimePoint RangedTimeOperationUpperBoundTimePointUnion `json:"upperBoundTimePoint"`
	JSON                struct {
		Comparison                   respjson.Field
		TimeOffset                   respjson.Field
		Type                         respjson.Field
		LowerBoundOffset             respjson.Field
		RangeType                    respjson.Field
		UpperBoundOffset             respjson.Field
		Timestamp                    respjson.Field
		LowerTimestamp               respjson.Field
		UpperTimestamp               respjson.Field
		EndpointBehavior             respjson.Field
		IncludeObjectsWithNoValueSet respjson.Field
		OperationType                respjson.Field
		Operator                     respjson.Field
		OperatorName                 respjson.Field
		PropertyParser               respjson.Field
		PropertyType                 respjson.Field
		TimePoint                    respjson.Field
		DefaultValue                 respjson.Field
		RenderSpec                   respjson.Field
		LowerBoundEndpointBehavior   respjson.Field
		LowerBoundTimePoint          respjson.Field
		UpperBoundEndpointBehavior   respjson.Field
		UpperBoundTimePoint          respjson.Field
		raw                          string
	} `json:"-"`
}

func (u AllPropertyTypesOperationPruningRefineByUnion) AsRelativeComparativeTimestampRefineBy() (v RelativeComparativeTimestampRefineBy) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AllPropertyTypesOperationPruningRefineByUnion) AsRelativeRangedTimestampRefineBy() (v RelativeRangedTimestampRefineBy) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AllPropertyTypesOperationPruningRefineByUnion) AsAbsoluteComparativeTimestampRefineBy() (v AbsoluteComparativeTimestampRefineBy) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AllPropertyTypesOperationPruningRefineByUnion) AsAbsoluteRangedTimestampRefineBy() (v AbsoluteRangedTimestampRefineBy) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AllPropertyTypesOperationPruningRefineByUnion) AsAllHistoryRefineBy() (v AllHistoryRefineBy) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AllPropertyTypesOperationPruningRefineByUnion) AsTimepoint() (v TimePointOperation) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AllPropertyTypesOperationPruningRefineByUnion) AsRangedtime() (v RangedTimeOperation) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AllPropertyTypesOperationPruningRefineByUnion) RawJSON() string { return u.JSON.raw }

func (r *AllPropertyTypesOperationPruningRefineByUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The definition of an association
type AssociationDefinition struct {
	// The unique ID of the associated object (e.g., a contact ID).
	ID int64 `json:"id" api:"required"`
	// Whether custom labels can be used in the association.
	AllowsCustomLabels bool `json:"allowsCustomLabels" api:"required"`
	// The cardinality from the source object's perspective, either "ONE_TO_ONE" or
	// "ONE_TO_MANY".
	//
	// Any of "ONE_TO_MANY", "ONE_TO_ONE".
	Cardinality AssociationDefinitionCardinality `json:"cardinality" api:"required"`
	// The error category
	//
	// Any of "HUBSPOT_DEFINED", "INTEGRATOR_DEFINED", "USER_DEFINED", "WORK".
	Category AssociationDefinitionCategory `json:"category" api:"required"`
	// The ID of the source object type (e.g., 0-1 for contacts).
	FromObjectTypeID string `json:"fromObjectTypeId" api:"required"`
	// Whether all potential linked objects are included in the association
	HasAllAssociatedObjects bool `json:"hasAllAssociatedObjects" api:"required"`
	// Whether deletions in the association should cause cascading deletes to linked
	// objects.
	HasCascadingDeletes bool `json:"hasCascadingDeletes" api:"required"`
	// Whether a user has set a limit for the number of source objects.
	HasUserEnforcedMaxFromObjectIDs bool `json:"hasUserEnforcedMaxFromObjectIds" api:"required"`
	// Whether a user has set a limit for the number of destination objects.
	HasUserEnforcedMaxToObjectIDs bool `json:"hasUserEnforcedMaxToObjectIds" api:"required"`
	// Whether the association is hidden or not.
	Hidden bool `json:"hidden" api:"required"`
	// Whether the reverse association can also support custom labels.
	InverseAllowsCustomLabels bool `json:"inverseAllowsCustomLabels" api:"required"`
	// The cardinality from the destination object's perspective, either "ONE_TO_ONE"
	// or "ONE_TO_MANY".
	//
	// Any of "ONE_TO_MANY", "ONE_TO_ONE".
	InverseCardinality AssociationDefinitionInverseCardinality `json:"inverseCardinality" api:"required"`
	// Whether all potential reverse linked objects are included in the association.
	InverseHasAllAssociatedObjects bool `json:"inverseHasAllAssociatedObjects" api:"required"`
	// The unique ID for the inverse side of the association.
	InverseID int64 `json:"inverseId" api:"required"`
	// The name used to describe the inverse relationship in this association
	InverseName string `json:"inverseName" api:"required"`
	IsDefault   bool   `json:"isDefault" api:"required"`
	// Whether the inverse association is considered primary.
	IsInversePrimary bool `json:"isInversePrimary" api:"required"`
	// Whether the association is the primary link between the entities involved.
	IsPrimary bool `json:"isPrimary" api:"required"`
	// The maximum number of source object IDs allowed in the association.
	MaxFromObjectIDs int64 `json:"maxFromObjectIds" api:"required"`
	// The maximum number of destination object IDs allowed in the association.
	MaxToObjectIDs int64 `json:"maxToObjectIds" api:"required"`
	// For labeled association types, the internal name of the association.
	Name string `json:"name" api:"required"`
	// A unique across-portal ID applied to the association.
	PortalUniqueIdentifier string `json:"portalUniqueIdentifier" api:"required"`
	ReadOnly               bool   `json:"readOnly" api:"required"`
	// The ID of the destination object type (e.g., 0-3 for deals).
	ToObjectTypeID string `json:"toObjectTypeId" api:"required"`
	// The name of the source object type (e.g,. "DEAL" or "QUOTE").
	//
	// Any of "ABANDONED_CART", "ACCEPTANCE_TEST", "AD", "AD_ACCOUNT", "AD_CAMPAIGN",
	// "AD_GROUP", "AI_FORECAST", "ALL_PAGES", "APPROVAL", "APPROVAL_STEP",
	// "ATTRIBUTION", "AUDIENCE", "AUTOMATION_JOURNEY", "AUTOMATION_PLATFORM_FLOW",
	// "AUTOMATION_PLATFORM_FLOW_ACTION", "BET_ALERT", "BET_DELIVERABLE_SERVICE",
	// "BLOG_LISTING_PAGE", "BLOG_POST", "CALL", "CAMPAIGN", "CAMPAIGN_BUDGET_ITEM",
	// "CAMPAIGN_SPEND_ITEM", "CAMPAIGN_STEP", "CAMPAIGN_TEMPLATE",
	// "CAMPAIGN_TEMPLATE_STEP", "CART", "CASE_STUDY", "CHATFLOW", "CLIP", "CMS_URL",
	// "COMBO_EVENT_CONFIGURATION", "COMMERCE_PAYMENT", "COMMUNICATION", "COMPANY",
	// "CONTACT", "CONTACT_CREATE_ATTRIBUTION", "CONTENT", "CONTENT_AUDIT",
	// "CONTENT_AUDIT_PAGE", "CONVERSATION", "CONVERSATION_INBOX",
	// "CONVERSATION_SESSION", "CRM_OBJECTS_DUMMY_TYPE", "CRM_PIPELINES_DUMMY_TYPE",
	// "CTA", "CTA_VARIANT", "DATA_PRIVACY_CONSENT", "DATA_SYNC_STATE", "DEAL",
	// "DEAL_CREATE_ATTRIBUTION", "DEAL_REGISTRATION", "DEAL_SPLIT", "DISCOUNT",
	// "DISCOUNT_CODE", "DISCOUNT_TEMPLATE", "EMAIL", "ENGAGEMENT", "EXPORT",
	// "EXTERNAL_WEB_URL", "FEE", "FEEDBACK_SUBMISSION", "FEEDBACK_SURVEY",
	// "FILE_MANAGER_FILE", "FILE_MANAGER_FOLDER", "FOLDER", "FORECAST", "FORM",
	// "FORM_SUBMISSION_INBOUNDDB", "GOAL_TARGET", "GOAL_TARGET_GROUP",
	// "GOAL_TEMPLATE", "GSC_PROPERTY", "HUB", "IMPORT", "INVOICE", "KEYWORD",
	// "KNOWLEDGE_ARTICLE", "LANDING_PAGE", "LEAD", "LINE_ITEM", "MARKETING_CALENDAR",
	// "MARKETING_CAMPAIGN_UTM", "MARKETING_EMAIL", "MARKETING_EVENT",
	// "MARKETING_EVENT_ATTENDANCE", "MARKETING_SMS", "MEDIA_BRIDGE", "MEETING_EVENT",
	// "MIC", "NOTE", "OBJECT_LIST", "ORDER", "OWNER", "PARTNER_ACCOUNT",
	// "PARTNER_CLIENT", "PARTNER_CLIENT_REVENUE", "PARTNER_SERVICE", "PAYMENT_LINK",
	// "PAYMENT_SCHEDULE", "PAYMENT_SCHEDULE_INSTALLMENT", "PERMISSIONS_TESTING",
	// "PLAYBOOK", "PLAYBOOK_QUESTION", "PLAYBOOK_SUBMISSION",
	// "PLAYBOOK_SUBMISSION_ANSWER", "PLAYLIST", "PLAYLIST_FOLDER", "PODCAST_EPISODE",
	// "PORTAL", "PORTAL_OBJECT_SYNC_MESSAGE", "POSTAL_MAIL", "PRIVACY_SCANNER_COOKIE",
	// "PRODUCT", "PRODUCT_OR_FOLDER", "PROPERTY_INFO",
	// "PROSPECTING_AGENT_CONTACT_ASSIGNMENT", "PUBLISHING_TASK",
	// "QUARANTINED_SUBMISSION", "QUOTA", "QUOTE", "QUOTE_FIELD", "QUOTE_MODULE",
	// "QUOTE_MODULE_FIELD", "QUOTE_TEMPLATE", "RESTORABLE_CRM_OBJECT", "ROSTER",
	// "ROSTER_MEMBER", "SALES_DOCUMENT", "SALES_TASK", "SALES_WORKLOAD",
	// "SALESFORCE_SYNC_ERROR", "SCHEDULING_PAGE", "SCHEMAS_BACKEND_TEST",
	// "SCORE_CONFIGURATION", "SEQUENCE", "SEQUENCE_ENROLLMENT", "SEQUENCE_STEP",
	// "SEQUENCE_STEP_ENROLLMENT", "SERVICE", "SITE_PAGE", "SNIPPET",
	// "SOCIAL_BROADCAST", "SOCIAL_CHANNEL", "SOCIAL_POST", "SOCIAL_PROFILE",
	// "SOX_PROTECTED_DUMMY_TYPE", "SOX_PROTECTED_TEST_TYPE", "SUBMISSION_TAG",
	// "SUBSCRIPTION", "TASK", "TASK_TEMPLATE", "TAX", "TEMPLATE", "TICKET", "UNKNOWN",
	// "UNSUBSCRIBE", "USER", "VIEW", "VIEW_BLOCK", "WEB_INTERACTIVE".
	FromObjectType AssociationDefinitionFromObjectType `json:"fromObjectType"`
	// Any of "DEFAULT", "INTERNAL", "USER_CONFIGURED".
	HiddenReason AssociationDefinitionHiddenReason `json:"hiddenReason"`
	// The label used to describe the reverse relationship in an association.
	InverseLabel string `json:"inverseLabel"`
	// The label given to an association.
	Label string `json:"label"`
	// The name of the destination object type (e.g,. "DEAL" or "QUOTE").
	//
	// Any of "ABANDONED_CART", "ACCEPTANCE_TEST", "AD", "AD_ACCOUNT", "AD_CAMPAIGN",
	// "AD_GROUP", "AI_FORECAST", "ALL_PAGES", "APPROVAL", "APPROVAL_STEP",
	// "ATTRIBUTION", "AUDIENCE", "AUTOMATION_JOURNEY", "AUTOMATION_PLATFORM_FLOW",
	// "AUTOMATION_PLATFORM_FLOW_ACTION", "BET_ALERT", "BET_DELIVERABLE_SERVICE",
	// "BLOG_LISTING_PAGE", "BLOG_POST", "CALL", "CAMPAIGN", "CAMPAIGN_BUDGET_ITEM",
	// "CAMPAIGN_SPEND_ITEM", "CAMPAIGN_STEP", "CAMPAIGN_TEMPLATE",
	// "CAMPAIGN_TEMPLATE_STEP", "CART", "CASE_STUDY", "CHATFLOW", "CLIP", "CMS_URL",
	// "COMBO_EVENT_CONFIGURATION", "COMMERCE_PAYMENT", "COMMUNICATION", "COMPANY",
	// "CONTACT", "CONTACT_CREATE_ATTRIBUTION", "CONTENT", "CONTENT_AUDIT",
	// "CONTENT_AUDIT_PAGE", "CONVERSATION", "CONVERSATION_INBOX",
	// "CONVERSATION_SESSION", "CRM_OBJECTS_DUMMY_TYPE", "CRM_PIPELINES_DUMMY_TYPE",
	// "CTA", "CTA_VARIANT", "DATA_PRIVACY_CONSENT", "DATA_SYNC_STATE", "DEAL",
	// "DEAL_CREATE_ATTRIBUTION", "DEAL_REGISTRATION", "DEAL_SPLIT", "DISCOUNT",
	// "DISCOUNT_CODE", "DISCOUNT_TEMPLATE", "EMAIL", "ENGAGEMENT", "EXPORT",
	// "EXTERNAL_WEB_URL", "FEE", "FEEDBACK_SUBMISSION", "FEEDBACK_SURVEY",
	// "FILE_MANAGER_FILE", "FILE_MANAGER_FOLDER", "FOLDER", "FORECAST", "FORM",
	// "FORM_SUBMISSION_INBOUNDDB", "GOAL_TARGET", "GOAL_TARGET_GROUP",
	// "GOAL_TEMPLATE", "GSC_PROPERTY", "HUB", "IMPORT", "INVOICE", "KEYWORD",
	// "KNOWLEDGE_ARTICLE", "LANDING_PAGE", "LEAD", "LINE_ITEM", "MARKETING_CALENDAR",
	// "MARKETING_CAMPAIGN_UTM", "MARKETING_EMAIL", "MARKETING_EVENT",
	// "MARKETING_EVENT_ATTENDANCE", "MARKETING_SMS", "MEDIA_BRIDGE", "MEETING_EVENT",
	// "MIC", "NOTE", "OBJECT_LIST", "ORDER", "OWNER", "PARTNER_ACCOUNT",
	// "PARTNER_CLIENT", "PARTNER_CLIENT_REVENUE", "PARTNER_SERVICE", "PAYMENT_LINK",
	// "PAYMENT_SCHEDULE", "PAYMENT_SCHEDULE_INSTALLMENT", "PERMISSIONS_TESTING",
	// "PLAYBOOK", "PLAYBOOK_QUESTION", "PLAYBOOK_SUBMISSION",
	// "PLAYBOOK_SUBMISSION_ANSWER", "PLAYLIST", "PLAYLIST_FOLDER", "PODCAST_EPISODE",
	// "PORTAL", "PORTAL_OBJECT_SYNC_MESSAGE", "POSTAL_MAIL", "PRIVACY_SCANNER_COOKIE",
	// "PRODUCT", "PRODUCT_OR_FOLDER", "PROPERTY_INFO",
	// "PROSPECTING_AGENT_CONTACT_ASSIGNMENT", "PUBLISHING_TASK",
	// "QUARANTINED_SUBMISSION", "QUOTA", "QUOTE", "QUOTE_FIELD", "QUOTE_MODULE",
	// "QUOTE_MODULE_FIELD", "QUOTE_TEMPLATE", "RESTORABLE_CRM_OBJECT", "ROSTER",
	// "ROSTER_MEMBER", "SALES_DOCUMENT", "SALES_TASK", "SALES_WORKLOAD",
	// "SALESFORCE_SYNC_ERROR", "SCHEDULING_PAGE", "SCHEMAS_BACKEND_TEST",
	// "SCORE_CONFIGURATION", "SEQUENCE", "SEQUENCE_ENROLLMENT", "SEQUENCE_STEP",
	// "SEQUENCE_STEP_ENROLLMENT", "SERVICE", "SITE_PAGE", "SNIPPET",
	// "SOCIAL_BROADCAST", "SOCIAL_CHANNEL", "SOCIAL_POST", "SOCIAL_PROFILE",
	// "SOX_PROTECTED_DUMMY_TYPE", "SOX_PROTECTED_TEST_TYPE", "SUBMISSION_TAG",
	// "SUBSCRIPTION", "TASK", "TASK_TEMPLATE", "TAX", "TEMPLATE", "TICKET", "UNKNOWN",
	// "UNSUBSCRIBE", "USER", "VIEW", "VIEW_BLOCK", "WEB_INTERACTIVE".
	ToObjectType AssociationDefinitionToObjectType `json:"toObjectType"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                              respjson.Field
		AllowsCustomLabels              respjson.Field
		Cardinality                     respjson.Field
		Category                        respjson.Field
		FromObjectTypeID                respjson.Field
		HasAllAssociatedObjects         respjson.Field
		HasCascadingDeletes             respjson.Field
		HasUserEnforcedMaxFromObjectIDs respjson.Field
		HasUserEnforcedMaxToObjectIDs   respjson.Field
		Hidden                          respjson.Field
		InverseAllowsCustomLabels       respjson.Field
		InverseCardinality              respjson.Field
		InverseHasAllAssociatedObjects  respjson.Field
		InverseID                       respjson.Field
		InverseName                     respjson.Field
		IsDefault                       respjson.Field
		IsInversePrimary                respjson.Field
		IsPrimary                       respjson.Field
		MaxFromObjectIDs                respjson.Field
		MaxToObjectIDs                  respjson.Field
		Name                            respjson.Field
		PortalUniqueIdentifier          respjson.Field
		ReadOnly                        respjson.Field
		ToObjectTypeID                  respjson.Field
		FromObjectType                  respjson.Field
		HiddenReason                    respjson.Field
		InverseLabel                    respjson.Field
		Label                           respjson.Field
		ToObjectType                    respjson.Field
		ExtraFields                     map[string]respjson.Field
		raw                             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AssociationDefinition) RawJSON() string { return r.JSON.raw }
func (r *AssociationDefinition) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The cardinality from the source object's perspective, either "ONE_TO_ONE" or
// "ONE_TO_MANY".
type AssociationDefinitionCardinality string

const (
	AssociationDefinitionCardinalityOneToMany AssociationDefinitionCardinality = "ONE_TO_MANY"
	AssociationDefinitionCardinalityOneToOne  AssociationDefinitionCardinality = "ONE_TO_ONE"
)

// The error category
type AssociationDefinitionCategory string

const (
	AssociationDefinitionCategoryHubspotDefined    AssociationDefinitionCategory = "HUBSPOT_DEFINED"
	AssociationDefinitionCategoryIntegratorDefined AssociationDefinitionCategory = "INTEGRATOR_DEFINED"
	AssociationDefinitionCategoryUserDefined       AssociationDefinitionCategory = "USER_DEFINED"
	AssociationDefinitionCategoryWork              AssociationDefinitionCategory = "WORK"
)

// The cardinality from the destination object's perspective, either "ONE_TO_ONE"
// or "ONE_TO_MANY".
type AssociationDefinitionInverseCardinality string

const (
	AssociationDefinitionInverseCardinalityOneToMany AssociationDefinitionInverseCardinality = "ONE_TO_MANY"
	AssociationDefinitionInverseCardinalityOneToOne  AssociationDefinitionInverseCardinality = "ONE_TO_ONE"
)

// The name of the source object type (e.g,. "DEAL" or "QUOTE").
type AssociationDefinitionFromObjectType string

const (
	AssociationDefinitionFromObjectTypeAbandonedCart                     AssociationDefinitionFromObjectType = "ABANDONED_CART"
	AssociationDefinitionFromObjectTypeAcceptanceTest                    AssociationDefinitionFromObjectType = "ACCEPTANCE_TEST"
	AssociationDefinitionFromObjectTypeAd                                AssociationDefinitionFromObjectType = "AD"
	AssociationDefinitionFromObjectTypeAdAccount                         AssociationDefinitionFromObjectType = "AD_ACCOUNT"
	AssociationDefinitionFromObjectTypeAdCampaign                        AssociationDefinitionFromObjectType = "AD_CAMPAIGN"
	AssociationDefinitionFromObjectTypeAdGroup                           AssociationDefinitionFromObjectType = "AD_GROUP"
	AssociationDefinitionFromObjectTypeAIForecast                        AssociationDefinitionFromObjectType = "AI_FORECAST"
	AssociationDefinitionFromObjectTypeAllPages                          AssociationDefinitionFromObjectType = "ALL_PAGES"
	AssociationDefinitionFromObjectTypeApproval                          AssociationDefinitionFromObjectType = "APPROVAL"
	AssociationDefinitionFromObjectTypeApprovalStep                      AssociationDefinitionFromObjectType = "APPROVAL_STEP"
	AssociationDefinitionFromObjectTypeAttribution                       AssociationDefinitionFromObjectType = "ATTRIBUTION"
	AssociationDefinitionFromObjectTypeAudience                          AssociationDefinitionFromObjectType = "AUDIENCE"
	AssociationDefinitionFromObjectTypeAutomationJourney                 AssociationDefinitionFromObjectType = "AUTOMATION_JOURNEY"
	AssociationDefinitionFromObjectTypeAutomationPlatformFlow            AssociationDefinitionFromObjectType = "AUTOMATION_PLATFORM_FLOW"
	AssociationDefinitionFromObjectTypeAutomationPlatformFlowAction      AssociationDefinitionFromObjectType = "AUTOMATION_PLATFORM_FLOW_ACTION"
	AssociationDefinitionFromObjectTypeBetAlert                          AssociationDefinitionFromObjectType = "BET_ALERT"
	AssociationDefinitionFromObjectTypeBetDeliverableService             AssociationDefinitionFromObjectType = "BET_DELIVERABLE_SERVICE"
	AssociationDefinitionFromObjectTypeBlogListingPage                   AssociationDefinitionFromObjectType = "BLOG_LISTING_PAGE"
	AssociationDefinitionFromObjectTypeBlogPost                          AssociationDefinitionFromObjectType = "BLOG_POST"
	AssociationDefinitionFromObjectTypeCall                              AssociationDefinitionFromObjectType = "CALL"
	AssociationDefinitionFromObjectTypeCampaign                          AssociationDefinitionFromObjectType = "CAMPAIGN"
	AssociationDefinitionFromObjectTypeCampaignBudgetItem                AssociationDefinitionFromObjectType = "CAMPAIGN_BUDGET_ITEM"
	AssociationDefinitionFromObjectTypeCampaignSpendItem                 AssociationDefinitionFromObjectType = "CAMPAIGN_SPEND_ITEM"
	AssociationDefinitionFromObjectTypeCampaignStep                      AssociationDefinitionFromObjectType = "CAMPAIGN_STEP"
	AssociationDefinitionFromObjectTypeCampaignTemplate                  AssociationDefinitionFromObjectType = "CAMPAIGN_TEMPLATE"
	AssociationDefinitionFromObjectTypeCampaignTemplateStep              AssociationDefinitionFromObjectType = "CAMPAIGN_TEMPLATE_STEP"
	AssociationDefinitionFromObjectTypeCart                              AssociationDefinitionFromObjectType = "CART"
	AssociationDefinitionFromObjectTypeCaseStudy                         AssociationDefinitionFromObjectType = "CASE_STUDY"
	AssociationDefinitionFromObjectTypeChatflow                          AssociationDefinitionFromObjectType = "CHATFLOW"
	AssociationDefinitionFromObjectTypeClip                              AssociationDefinitionFromObjectType = "CLIP"
	AssociationDefinitionFromObjectTypeCmsURL                            AssociationDefinitionFromObjectType = "CMS_URL"
	AssociationDefinitionFromObjectTypeComboEventConfiguration           AssociationDefinitionFromObjectType = "COMBO_EVENT_CONFIGURATION"
	AssociationDefinitionFromObjectTypeCommercePayment                   AssociationDefinitionFromObjectType = "COMMERCE_PAYMENT"
	AssociationDefinitionFromObjectTypeCommunication                     AssociationDefinitionFromObjectType = "COMMUNICATION"
	AssociationDefinitionFromObjectTypeCompany                           AssociationDefinitionFromObjectType = "COMPANY"
	AssociationDefinitionFromObjectTypeContact                           AssociationDefinitionFromObjectType = "CONTACT"
	AssociationDefinitionFromObjectTypeContactCreateAttribution          AssociationDefinitionFromObjectType = "CONTACT_CREATE_ATTRIBUTION"
	AssociationDefinitionFromObjectTypeContent                           AssociationDefinitionFromObjectType = "CONTENT"
	AssociationDefinitionFromObjectTypeContentAudit                      AssociationDefinitionFromObjectType = "CONTENT_AUDIT"
	AssociationDefinitionFromObjectTypeContentAuditPage                  AssociationDefinitionFromObjectType = "CONTENT_AUDIT_PAGE"
	AssociationDefinitionFromObjectTypeConversation                      AssociationDefinitionFromObjectType = "CONVERSATION"
	AssociationDefinitionFromObjectTypeConversationInbox                 AssociationDefinitionFromObjectType = "CONVERSATION_INBOX"
	AssociationDefinitionFromObjectTypeConversationSession               AssociationDefinitionFromObjectType = "CONVERSATION_SESSION"
	AssociationDefinitionFromObjectTypeCrmObjectsDummyType               AssociationDefinitionFromObjectType = "CRM_OBJECTS_DUMMY_TYPE"
	AssociationDefinitionFromObjectTypeCrmPipelinesDummyType             AssociationDefinitionFromObjectType = "CRM_PIPELINES_DUMMY_TYPE"
	AssociationDefinitionFromObjectTypeCta                               AssociationDefinitionFromObjectType = "CTA"
	AssociationDefinitionFromObjectTypeCtaVariant                        AssociationDefinitionFromObjectType = "CTA_VARIANT"
	AssociationDefinitionFromObjectTypeDataPrivacyConsent                AssociationDefinitionFromObjectType = "DATA_PRIVACY_CONSENT"
	AssociationDefinitionFromObjectTypeDataSyncState                     AssociationDefinitionFromObjectType = "DATA_SYNC_STATE"
	AssociationDefinitionFromObjectTypeDeal                              AssociationDefinitionFromObjectType = "DEAL"
	AssociationDefinitionFromObjectTypeDealCreateAttribution             AssociationDefinitionFromObjectType = "DEAL_CREATE_ATTRIBUTION"
	AssociationDefinitionFromObjectTypeDealRegistration                  AssociationDefinitionFromObjectType = "DEAL_REGISTRATION"
	AssociationDefinitionFromObjectTypeDealSplit                         AssociationDefinitionFromObjectType = "DEAL_SPLIT"
	AssociationDefinitionFromObjectTypeDiscount                          AssociationDefinitionFromObjectType = "DISCOUNT"
	AssociationDefinitionFromObjectTypeDiscountCode                      AssociationDefinitionFromObjectType = "DISCOUNT_CODE"
	AssociationDefinitionFromObjectTypeDiscountTemplate                  AssociationDefinitionFromObjectType = "DISCOUNT_TEMPLATE"
	AssociationDefinitionFromObjectTypeEmail                             AssociationDefinitionFromObjectType = "EMAIL"
	AssociationDefinitionFromObjectTypeEngagement                        AssociationDefinitionFromObjectType = "ENGAGEMENT"
	AssociationDefinitionFromObjectTypeExport                            AssociationDefinitionFromObjectType = "EXPORT"
	AssociationDefinitionFromObjectTypeExternalWebURL                    AssociationDefinitionFromObjectType = "EXTERNAL_WEB_URL"
	AssociationDefinitionFromObjectTypeFee                               AssociationDefinitionFromObjectType = "FEE"
	AssociationDefinitionFromObjectTypeFeedbackSubmission                AssociationDefinitionFromObjectType = "FEEDBACK_SUBMISSION"
	AssociationDefinitionFromObjectTypeFeedbackSurvey                    AssociationDefinitionFromObjectType = "FEEDBACK_SURVEY"
	AssociationDefinitionFromObjectTypeFileManagerFile                   AssociationDefinitionFromObjectType = "FILE_MANAGER_FILE"
	AssociationDefinitionFromObjectTypeFileManagerFolder                 AssociationDefinitionFromObjectType = "FILE_MANAGER_FOLDER"
	AssociationDefinitionFromObjectTypeFolder                            AssociationDefinitionFromObjectType = "FOLDER"
	AssociationDefinitionFromObjectTypeForecast                          AssociationDefinitionFromObjectType = "FORECAST"
	AssociationDefinitionFromObjectTypeForm                              AssociationDefinitionFromObjectType = "FORM"
	AssociationDefinitionFromObjectTypeFormSubmissionInbounddb           AssociationDefinitionFromObjectType = "FORM_SUBMISSION_INBOUNDDB"
	AssociationDefinitionFromObjectTypeGoalTarget                        AssociationDefinitionFromObjectType = "GOAL_TARGET"
	AssociationDefinitionFromObjectTypeGoalTargetGroup                   AssociationDefinitionFromObjectType = "GOAL_TARGET_GROUP"
	AssociationDefinitionFromObjectTypeGoalTemplate                      AssociationDefinitionFromObjectType = "GOAL_TEMPLATE"
	AssociationDefinitionFromObjectTypeGscProperty                       AssociationDefinitionFromObjectType = "GSC_PROPERTY"
	AssociationDefinitionFromObjectTypeHub                               AssociationDefinitionFromObjectType = "HUB"
	AssociationDefinitionFromObjectTypeImport                            AssociationDefinitionFromObjectType = "IMPORT"
	AssociationDefinitionFromObjectTypeInvoice                           AssociationDefinitionFromObjectType = "INVOICE"
	AssociationDefinitionFromObjectTypeKeyword                           AssociationDefinitionFromObjectType = "KEYWORD"
	AssociationDefinitionFromObjectTypeKnowledgeArticle                  AssociationDefinitionFromObjectType = "KNOWLEDGE_ARTICLE"
	AssociationDefinitionFromObjectTypeLandingPage                       AssociationDefinitionFromObjectType = "LANDING_PAGE"
	AssociationDefinitionFromObjectTypeLead                              AssociationDefinitionFromObjectType = "LEAD"
	AssociationDefinitionFromObjectTypeLineItem                          AssociationDefinitionFromObjectType = "LINE_ITEM"
	AssociationDefinitionFromObjectTypeMarketingCalendar                 AssociationDefinitionFromObjectType = "MARKETING_CALENDAR"
	AssociationDefinitionFromObjectTypeMarketingCampaignUtm              AssociationDefinitionFromObjectType = "MARKETING_CAMPAIGN_UTM"
	AssociationDefinitionFromObjectTypeMarketingEmail                    AssociationDefinitionFromObjectType = "MARKETING_EMAIL"
	AssociationDefinitionFromObjectTypeMarketingEvent                    AssociationDefinitionFromObjectType = "MARKETING_EVENT"
	AssociationDefinitionFromObjectTypeMarketingEventAttendance          AssociationDefinitionFromObjectType = "MARKETING_EVENT_ATTENDANCE"
	AssociationDefinitionFromObjectTypeMarketingSMS                      AssociationDefinitionFromObjectType = "MARKETING_SMS"
	AssociationDefinitionFromObjectTypeMediaBridge                       AssociationDefinitionFromObjectType = "MEDIA_BRIDGE"
	AssociationDefinitionFromObjectTypeMeetingEvent                      AssociationDefinitionFromObjectType = "MEETING_EVENT"
	AssociationDefinitionFromObjectTypeMic                               AssociationDefinitionFromObjectType = "MIC"
	AssociationDefinitionFromObjectTypeNote                              AssociationDefinitionFromObjectType = "NOTE"
	AssociationDefinitionFromObjectTypeObjectList                        AssociationDefinitionFromObjectType = "OBJECT_LIST"
	AssociationDefinitionFromObjectTypeOrder                             AssociationDefinitionFromObjectType = "ORDER"
	AssociationDefinitionFromObjectTypeOwner                             AssociationDefinitionFromObjectType = "OWNER"
	AssociationDefinitionFromObjectTypePartnerAccount                    AssociationDefinitionFromObjectType = "PARTNER_ACCOUNT"
	AssociationDefinitionFromObjectTypePartnerClient                     AssociationDefinitionFromObjectType = "PARTNER_CLIENT"
	AssociationDefinitionFromObjectTypePartnerClientRevenue              AssociationDefinitionFromObjectType = "PARTNER_CLIENT_REVENUE"
	AssociationDefinitionFromObjectTypePartnerService                    AssociationDefinitionFromObjectType = "PARTNER_SERVICE"
	AssociationDefinitionFromObjectTypePaymentLink                       AssociationDefinitionFromObjectType = "PAYMENT_LINK"
	AssociationDefinitionFromObjectTypePaymentSchedule                   AssociationDefinitionFromObjectType = "PAYMENT_SCHEDULE"
	AssociationDefinitionFromObjectTypePaymentScheduleInstallment        AssociationDefinitionFromObjectType = "PAYMENT_SCHEDULE_INSTALLMENT"
	AssociationDefinitionFromObjectTypePermissionsTesting                AssociationDefinitionFromObjectType = "PERMISSIONS_TESTING"
	AssociationDefinitionFromObjectTypePlaybook                          AssociationDefinitionFromObjectType = "PLAYBOOK"
	AssociationDefinitionFromObjectTypePlaybookQuestion                  AssociationDefinitionFromObjectType = "PLAYBOOK_QUESTION"
	AssociationDefinitionFromObjectTypePlaybookSubmission                AssociationDefinitionFromObjectType = "PLAYBOOK_SUBMISSION"
	AssociationDefinitionFromObjectTypePlaybookSubmissionAnswer          AssociationDefinitionFromObjectType = "PLAYBOOK_SUBMISSION_ANSWER"
	AssociationDefinitionFromObjectTypePlaylist                          AssociationDefinitionFromObjectType = "PLAYLIST"
	AssociationDefinitionFromObjectTypePlaylistFolder                    AssociationDefinitionFromObjectType = "PLAYLIST_FOLDER"
	AssociationDefinitionFromObjectTypePodcastEpisode                    AssociationDefinitionFromObjectType = "PODCAST_EPISODE"
	AssociationDefinitionFromObjectTypePortal                            AssociationDefinitionFromObjectType = "PORTAL"
	AssociationDefinitionFromObjectTypePortalObjectSyncMessage           AssociationDefinitionFromObjectType = "PORTAL_OBJECT_SYNC_MESSAGE"
	AssociationDefinitionFromObjectTypePostalMail                        AssociationDefinitionFromObjectType = "POSTAL_MAIL"
	AssociationDefinitionFromObjectTypePrivacyScannerCookie              AssociationDefinitionFromObjectType = "PRIVACY_SCANNER_COOKIE"
	AssociationDefinitionFromObjectTypeProduct                           AssociationDefinitionFromObjectType = "PRODUCT"
	AssociationDefinitionFromObjectTypeProductOrFolder                   AssociationDefinitionFromObjectType = "PRODUCT_OR_FOLDER"
	AssociationDefinitionFromObjectTypePropertyInfo                      AssociationDefinitionFromObjectType = "PROPERTY_INFO"
	AssociationDefinitionFromObjectTypeProspectingAgentContactAssignment AssociationDefinitionFromObjectType = "PROSPECTING_AGENT_CONTACT_ASSIGNMENT"
	AssociationDefinitionFromObjectTypePublishingTask                    AssociationDefinitionFromObjectType = "PUBLISHING_TASK"
	AssociationDefinitionFromObjectTypeQuarantinedSubmission             AssociationDefinitionFromObjectType = "QUARANTINED_SUBMISSION"
	AssociationDefinitionFromObjectTypeQuota                             AssociationDefinitionFromObjectType = "QUOTA"
	AssociationDefinitionFromObjectTypeQuote                             AssociationDefinitionFromObjectType = "QUOTE"
	AssociationDefinitionFromObjectTypeQuoteField                        AssociationDefinitionFromObjectType = "QUOTE_FIELD"
	AssociationDefinitionFromObjectTypeQuoteModule                       AssociationDefinitionFromObjectType = "QUOTE_MODULE"
	AssociationDefinitionFromObjectTypeQuoteModuleField                  AssociationDefinitionFromObjectType = "QUOTE_MODULE_FIELD"
	AssociationDefinitionFromObjectTypeQuoteTemplate                     AssociationDefinitionFromObjectType = "QUOTE_TEMPLATE"
	AssociationDefinitionFromObjectTypeRestorableCrmObject               AssociationDefinitionFromObjectType = "RESTORABLE_CRM_OBJECT"
	AssociationDefinitionFromObjectTypeRoster                            AssociationDefinitionFromObjectType = "ROSTER"
	AssociationDefinitionFromObjectTypeRosterMember                      AssociationDefinitionFromObjectType = "ROSTER_MEMBER"
	AssociationDefinitionFromObjectTypeSalesDocument                     AssociationDefinitionFromObjectType = "SALES_DOCUMENT"
	AssociationDefinitionFromObjectTypeSalesTask                         AssociationDefinitionFromObjectType = "SALES_TASK"
	AssociationDefinitionFromObjectTypeSalesWorkload                     AssociationDefinitionFromObjectType = "SALES_WORKLOAD"
	AssociationDefinitionFromObjectTypeSalesforceSyncError               AssociationDefinitionFromObjectType = "SALESFORCE_SYNC_ERROR"
	AssociationDefinitionFromObjectTypeSchedulingPage                    AssociationDefinitionFromObjectType = "SCHEDULING_PAGE"
	AssociationDefinitionFromObjectTypeSchemasBackendTest                AssociationDefinitionFromObjectType = "SCHEMAS_BACKEND_TEST"
	AssociationDefinitionFromObjectTypeScoreConfiguration                AssociationDefinitionFromObjectType = "SCORE_CONFIGURATION"
	AssociationDefinitionFromObjectTypeSequence                          AssociationDefinitionFromObjectType = "SEQUENCE"
	AssociationDefinitionFromObjectTypeSequenceEnrollment                AssociationDefinitionFromObjectType = "SEQUENCE_ENROLLMENT"
	AssociationDefinitionFromObjectTypeSequenceStep                      AssociationDefinitionFromObjectType = "SEQUENCE_STEP"
	AssociationDefinitionFromObjectTypeSequenceStepEnrollment            AssociationDefinitionFromObjectType = "SEQUENCE_STEP_ENROLLMENT"
	AssociationDefinitionFromObjectTypeService                           AssociationDefinitionFromObjectType = "SERVICE"
	AssociationDefinitionFromObjectTypeSitePage                          AssociationDefinitionFromObjectType = "SITE_PAGE"
	AssociationDefinitionFromObjectTypeSnippet                           AssociationDefinitionFromObjectType = "SNIPPET"
	AssociationDefinitionFromObjectTypeSocialBroadcast                   AssociationDefinitionFromObjectType = "SOCIAL_BROADCAST"
	AssociationDefinitionFromObjectTypeSocialChannel                     AssociationDefinitionFromObjectType = "SOCIAL_CHANNEL"
	AssociationDefinitionFromObjectTypeSocialPost                        AssociationDefinitionFromObjectType = "SOCIAL_POST"
	AssociationDefinitionFromObjectTypeSocialProfile                     AssociationDefinitionFromObjectType = "SOCIAL_PROFILE"
	AssociationDefinitionFromObjectTypeSoxProtectedDummyType             AssociationDefinitionFromObjectType = "SOX_PROTECTED_DUMMY_TYPE"
	AssociationDefinitionFromObjectTypeSoxProtectedTestType              AssociationDefinitionFromObjectType = "SOX_PROTECTED_TEST_TYPE"
	AssociationDefinitionFromObjectTypeSubmissionTag                     AssociationDefinitionFromObjectType = "SUBMISSION_TAG"
	AssociationDefinitionFromObjectTypeSubscription                      AssociationDefinitionFromObjectType = "SUBSCRIPTION"
	AssociationDefinitionFromObjectTypeTask                              AssociationDefinitionFromObjectType = "TASK"
	AssociationDefinitionFromObjectTypeTaskTemplate                      AssociationDefinitionFromObjectType = "TASK_TEMPLATE"
	AssociationDefinitionFromObjectTypeTax                               AssociationDefinitionFromObjectType = "TAX"
	AssociationDefinitionFromObjectTypeTemplate                          AssociationDefinitionFromObjectType = "TEMPLATE"
	AssociationDefinitionFromObjectTypeTicket                            AssociationDefinitionFromObjectType = "TICKET"
	AssociationDefinitionFromObjectTypeUnknown                           AssociationDefinitionFromObjectType = "UNKNOWN"
	AssociationDefinitionFromObjectTypeUnsubscribe                       AssociationDefinitionFromObjectType = "UNSUBSCRIBE"
	AssociationDefinitionFromObjectTypeUser                              AssociationDefinitionFromObjectType = "USER"
	AssociationDefinitionFromObjectTypeView                              AssociationDefinitionFromObjectType = "VIEW"
	AssociationDefinitionFromObjectTypeViewBlock                         AssociationDefinitionFromObjectType = "VIEW_BLOCK"
	AssociationDefinitionFromObjectTypeWebInteractive                    AssociationDefinitionFromObjectType = "WEB_INTERACTIVE"
)

type AssociationDefinitionHiddenReason string

const (
	AssociationDefinitionHiddenReasonDefault        AssociationDefinitionHiddenReason = "DEFAULT"
	AssociationDefinitionHiddenReasonInternal       AssociationDefinitionHiddenReason = "INTERNAL"
	AssociationDefinitionHiddenReasonUserConfigured AssociationDefinitionHiddenReason = "USER_CONFIGURED"
)

// The name of the destination object type (e.g,. "DEAL" or "QUOTE").
type AssociationDefinitionToObjectType string

const (
	AssociationDefinitionToObjectTypeAbandonedCart                     AssociationDefinitionToObjectType = "ABANDONED_CART"
	AssociationDefinitionToObjectTypeAcceptanceTest                    AssociationDefinitionToObjectType = "ACCEPTANCE_TEST"
	AssociationDefinitionToObjectTypeAd                                AssociationDefinitionToObjectType = "AD"
	AssociationDefinitionToObjectTypeAdAccount                         AssociationDefinitionToObjectType = "AD_ACCOUNT"
	AssociationDefinitionToObjectTypeAdCampaign                        AssociationDefinitionToObjectType = "AD_CAMPAIGN"
	AssociationDefinitionToObjectTypeAdGroup                           AssociationDefinitionToObjectType = "AD_GROUP"
	AssociationDefinitionToObjectTypeAIForecast                        AssociationDefinitionToObjectType = "AI_FORECAST"
	AssociationDefinitionToObjectTypeAllPages                          AssociationDefinitionToObjectType = "ALL_PAGES"
	AssociationDefinitionToObjectTypeApproval                          AssociationDefinitionToObjectType = "APPROVAL"
	AssociationDefinitionToObjectTypeApprovalStep                      AssociationDefinitionToObjectType = "APPROVAL_STEP"
	AssociationDefinitionToObjectTypeAttribution                       AssociationDefinitionToObjectType = "ATTRIBUTION"
	AssociationDefinitionToObjectTypeAudience                          AssociationDefinitionToObjectType = "AUDIENCE"
	AssociationDefinitionToObjectTypeAutomationJourney                 AssociationDefinitionToObjectType = "AUTOMATION_JOURNEY"
	AssociationDefinitionToObjectTypeAutomationPlatformFlow            AssociationDefinitionToObjectType = "AUTOMATION_PLATFORM_FLOW"
	AssociationDefinitionToObjectTypeAutomationPlatformFlowAction      AssociationDefinitionToObjectType = "AUTOMATION_PLATFORM_FLOW_ACTION"
	AssociationDefinitionToObjectTypeBetAlert                          AssociationDefinitionToObjectType = "BET_ALERT"
	AssociationDefinitionToObjectTypeBetDeliverableService             AssociationDefinitionToObjectType = "BET_DELIVERABLE_SERVICE"
	AssociationDefinitionToObjectTypeBlogListingPage                   AssociationDefinitionToObjectType = "BLOG_LISTING_PAGE"
	AssociationDefinitionToObjectTypeBlogPost                          AssociationDefinitionToObjectType = "BLOG_POST"
	AssociationDefinitionToObjectTypeCall                              AssociationDefinitionToObjectType = "CALL"
	AssociationDefinitionToObjectTypeCampaign                          AssociationDefinitionToObjectType = "CAMPAIGN"
	AssociationDefinitionToObjectTypeCampaignBudgetItem                AssociationDefinitionToObjectType = "CAMPAIGN_BUDGET_ITEM"
	AssociationDefinitionToObjectTypeCampaignSpendItem                 AssociationDefinitionToObjectType = "CAMPAIGN_SPEND_ITEM"
	AssociationDefinitionToObjectTypeCampaignStep                      AssociationDefinitionToObjectType = "CAMPAIGN_STEP"
	AssociationDefinitionToObjectTypeCampaignTemplate                  AssociationDefinitionToObjectType = "CAMPAIGN_TEMPLATE"
	AssociationDefinitionToObjectTypeCampaignTemplateStep              AssociationDefinitionToObjectType = "CAMPAIGN_TEMPLATE_STEP"
	AssociationDefinitionToObjectTypeCart                              AssociationDefinitionToObjectType = "CART"
	AssociationDefinitionToObjectTypeCaseStudy                         AssociationDefinitionToObjectType = "CASE_STUDY"
	AssociationDefinitionToObjectTypeChatflow                          AssociationDefinitionToObjectType = "CHATFLOW"
	AssociationDefinitionToObjectTypeClip                              AssociationDefinitionToObjectType = "CLIP"
	AssociationDefinitionToObjectTypeCmsURL                            AssociationDefinitionToObjectType = "CMS_URL"
	AssociationDefinitionToObjectTypeComboEventConfiguration           AssociationDefinitionToObjectType = "COMBO_EVENT_CONFIGURATION"
	AssociationDefinitionToObjectTypeCommercePayment                   AssociationDefinitionToObjectType = "COMMERCE_PAYMENT"
	AssociationDefinitionToObjectTypeCommunication                     AssociationDefinitionToObjectType = "COMMUNICATION"
	AssociationDefinitionToObjectTypeCompany                           AssociationDefinitionToObjectType = "COMPANY"
	AssociationDefinitionToObjectTypeContact                           AssociationDefinitionToObjectType = "CONTACT"
	AssociationDefinitionToObjectTypeContactCreateAttribution          AssociationDefinitionToObjectType = "CONTACT_CREATE_ATTRIBUTION"
	AssociationDefinitionToObjectTypeContent                           AssociationDefinitionToObjectType = "CONTENT"
	AssociationDefinitionToObjectTypeContentAudit                      AssociationDefinitionToObjectType = "CONTENT_AUDIT"
	AssociationDefinitionToObjectTypeContentAuditPage                  AssociationDefinitionToObjectType = "CONTENT_AUDIT_PAGE"
	AssociationDefinitionToObjectTypeConversation                      AssociationDefinitionToObjectType = "CONVERSATION"
	AssociationDefinitionToObjectTypeConversationInbox                 AssociationDefinitionToObjectType = "CONVERSATION_INBOX"
	AssociationDefinitionToObjectTypeConversationSession               AssociationDefinitionToObjectType = "CONVERSATION_SESSION"
	AssociationDefinitionToObjectTypeCrmObjectsDummyType               AssociationDefinitionToObjectType = "CRM_OBJECTS_DUMMY_TYPE"
	AssociationDefinitionToObjectTypeCrmPipelinesDummyType             AssociationDefinitionToObjectType = "CRM_PIPELINES_DUMMY_TYPE"
	AssociationDefinitionToObjectTypeCta                               AssociationDefinitionToObjectType = "CTA"
	AssociationDefinitionToObjectTypeCtaVariant                        AssociationDefinitionToObjectType = "CTA_VARIANT"
	AssociationDefinitionToObjectTypeDataPrivacyConsent                AssociationDefinitionToObjectType = "DATA_PRIVACY_CONSENT"
	AssociationDefinitionToObjectTypeDataSyncState                     AssociationDefinitionToObjectType = "DATA_SYNC_STATE"
	AssociationDefinitionToObjectTypeDeal                              AssociationDefinitionToObjectType = "DEAL"
	AssociationDefinitionToObjectTypeDealCreateAttribution             AssociationDefinitionToObjectType = "DEAL_CREATE_ATTRIBUTION"
	AssociationDefinitionToObjectTypeDealRegistration                  AssociationDefinitionToObjectType = "DEAL_REGISTRATION"
	AssociationDefinitionToObjectTypeDealSplit                         AssociationDefinitionToObjectType = "DEAL_SPLIT"
	AssociationDefinitionToObjectTypeDiscount                          AssociationDefinitionToObjectType = "DISCOUNT"
	AssociationDefinitionToObjectTypeDiscountCode                      AssociationDefinitionToObjectType = "DISCOUNT_CODE"
	AssociationDefinitionToObjectTypeDiscountTemplate                  AssociationDefinitionToObjectType = "DISCOUNT_TEMPLATE"
	AssociationDefinitionToObjectTypeEmail                             AssociationDefinitionToObjectType = "EMAIL"
	AssociationDefinitionToObjectTypeEngagement                        AssociationDefinitionToObjectType = "ENGAGEMENT"
	AssociationDefinitionToObjectTypeExport                            AssociationDefinitionToObjectType = "EXPORT"
	AssociationDefinitionToObjectTypeExternalWebURL                    AssociationDefinitionToObjectType = "EXTERNAL_WEB_URL"
	AssociationDefinitionToObjectTypeFee                               AssociationDefinitionToObjectType = "FEE"
	AssociationDefinitionToObjectTypeFeedbackSubmission                AssociationDefinitionToObjectType = "FEEDBACK_SUBMISSION"
	AssociationDefinitionToObjectTypeFeedbackSurvey                    AssociationDefinitionToObjectType = "FEEDBACK_SURVEY"
	AssociationDefinitionToObjectTypeFileManagerFile                   AssociationDefinitionToObjectType = "FILE_MANAGER_FILE"
	AssociationDefinitionToObjectTypeFileManagerFolder                 AssociationDefinitionToObjectType = "FILE_MANAGER_FOLDER"
	AssociationDefinitionToObjectTypeFolder                            AssociationDefinitionToObjectType = "FOLDER"
	AssociationDefinitionToObjectTypeForecast                          AssociationDefinitionToObjectType = "FORECAST"
	AssociationDefinitionToObjectTypeForm                              AssociationDefinitionToObjectType = "FORM"
	AssociationDefinitionToObjectTypeFormSubmissionInbounddb           AssociationDefinitionToObjectType = "FORM_SUBMISSION_INBOUNDDB"
	AssociationDefinitionToObjectTypeGoalTarget                        AssociationDefinitionToObjectType = "GOAL_TARGET"
	AssociationDefinitionToObjectTypeGoalTargetGroup                   AssociationDefinitionToObjectType = "GOAL_TARGET_GROUP"
	AssociationDefinitionToObjectTypeGoalTemplate                      AssociationDefinitionToObjectType = "GOAL_TEMPLATE"
	AssociationDefinitionToObjectTypeGscProperty                       AssociationDefinitionToObjectType = "GSC_PROPERTY"
	AssociationDefinitionToObjectTypeHub                               AssociationDefinitionToObjectType = "HUB"
	AssociationDefinitionToObjectTypeImport                            AssociationDefinitionToObjectType = "IMPORT"
	AssociationDefinitionToObjectTypeInvoice                           AssociationDefinitionToObjectType = "INVOICE"
	AssociationDefinitionToObjectTypeKeyword                           AssociationDefinitionToObjectType = "KEYWORD"
	AssociationDefinitionToObjectTypeKnowledgeArticle                  AssociationDefinitionToObjectType = "KNOWLEDGE_ARTICLE"
	AssociationDefinitionToObjectTypeLandingPage                       AssociationDefinitionToObjectType = "LANDING_PAGE"
	AssociationDefinitionToObjectTypeLead                              AssociationDefinitionToObjectType = "LEAD"
	AssociationDefinitionToObjectTypeLineItem                          AssociationDefinitionToObjectType = "LINE_ITEM"
	AssociationDefinitionToObjectTypeMarketingCalendar                 AssociationDefinitionToObjectType = "MARKETING_CALENDAR"
	AssociationDefinitionToObjectTypeMarketingCampaignUtm              AssociationDefinitionToObjectType = "MARKETING_CAMPAIGN_UTM"
	AssociationDefinitionToObjectTypeMarketingEmail                    AssociationDefinitionToObjectType = "MARKETING_EMAIL"
	AssociationDefinitionToObjectTypeMarketingEvent                    AssociationDefinitionToObjectType = "MARKETING_EVENT"
	AssociationDefinitionToObjectTypeMarketingEventAttendance          AssociationDefinitionToObjectType = "MARKETING_EVENT_ATTENDANCE"
	AssociationDefinitionToObjectTypeMarketingSMS                      AssociationDefinitionToObjectType = "MARKETING_SMS"
	AssociationDefinitionToObjectTypeMediaBridge                       AssociationDefinitionToObjectType = "MEDIA_BRIDGE"
	AssociationDefinitionToObjectTypeMeetingEvent                      AssociationDefinitionToObjectType = "MEETING_EVENT"
	AssociationDefinitionToObjectTypeMic                               AssociationDefinitionToObjectType = "MIC"
	AssociationDefinitionToObjectTypeNote                              AssociationDefinitionToObjectType = "NOTE"
	AssociationDefinitionToObjectTypeObjectList                        AssociationDefinitionToObjectType = "OBJECT_LIST"
	AssociationDefinitionToObjectTypeOrder                             AssociationDefinitionToObjectType = "ORDER"
	AssociationDefinitionToObjectTypeOwner                             AssociationDefinitionToObjectType = "OWNER"
	AssociationDefinitionToObjectTypePartnerAccount                    AssociationDefinitionToObjectType = "PARTNER_ACCOUNT"
	AssociationDefinitionToObjectTypePartnerClient                     AssociationDefinitionToObjectType = "PARTNER_CLIENT"
	AssociationDefinitionToObjectTypePartnerClientRevenue              AssociationDefinitionToObjectType = "PARTNER_CLIENT_REVENUE"
	AssociationDefinitionToObjectTypePartnerService                    AssociationDefinitionToObjectType = "PARTNER_SERVICE"
	AssociationDefinitionToObjectTypePaymentLink                       AssociationDefinitionToObjectType = "PAYMENT_LINK"
	AssociationDefinitionToObjectTypePaymentSchedule                   AssociationDefinitionToObjectType = "PAYMENT_SCHEDULE"
	AssociationDefinitionToObjectTypePaymentScheduleInstallment        AssociationDefinitionToObjectType = "PAYMENT_SCHEDULE_INSTALLMENT"
	AssociationDefinitionToObjectTypePermissionsTesting                AssociationDefinitionToObjectType = "PERMISSIONS_TESTING"
	AssociationDefinitionToObjectTypePlaybook                          AssociationDefinitionToObjectType = "PLAYBOOK"
	AssociationDefinitionToObjectTypePlaybookQuestion                  AssociationDefinitionToObjectType = "PLAYBOOK_QUESTION"
	AssociationDefinitionToObjectTypePlaybookSubmission                AssociationDefinitionToObjectType = "PLAYBOOK_SUBMISSION"
	AssociationDefinitionToObjectTypePlaybookSubmissionAnswer          AssociationDefinitionToObjectType = "PLAYBOOK_SUBMISSION_ANSWER"
	AssociationDefinitionToObjectTypePlaylist                          AssociationDefinitionToObjectType = "PLAYLIST"
	AssociationDefinitionToObjectTypePlaylistFolder                    AssociationDefinitionToObjectType = "PLAYLIST_FOLDER"
	AssociationDefinitionToObjectTypePodcastEpisode                    AssociationDefinitionToObjectType = "PODCAST_EPISODE"
	AssociationDefinitionToObjectTypePortal                            AssociationDefinitionToObjectType = "PORTAL"
	AssociationDefinitionToObjectTypePortalObjectSyncMessage           AssociationDefinitionToObjectType = "PORTAL_OBJECT_SYNC_MESSAGE"
	AssociationDefinitionToObjectTypePostalMail                        AssociationDefinitionToObjectType = "POSTAL_MAIL"
	AssociationDefinitionToObjectTypePrivacyScannerCookie              AssociationDefinitionToObjectType = "PRIVACY_SCANNER_COOKIE"
	AssociationDefinitionToObjectTypeProduct                           AssociationDefinitionToObjectType = "PRODUCT"
	AssociationDefinitionToObjectTypeProductOrFolder                   AssociationDefinitionToObjectType = "PRODUCT_OR_FOLDER"
	AssociationDefinitionToObjectTypePropertyInfo                      AssociationDefinitionToObjectType = "PROPERTY_INFO"
	AssociationDefinitionToObjectTypeProspectingAgentContactAssignment AssociationDefinitionToObjectType = "PROSPECTING_AGENT_CONTACT_ASSIGNMENT"
	AssociationDefinitionToObjectTypePublishingTask                    AssociationDefinitionToObjectType = "PUBLISHING_TASK"
	AssociationDefinitionToObjectTypeQuarantinedSubmission             AssociationDefinitionToObjectType = "QUARANTINED_SUBMISSION"
	AssociationDefinitionToObjectTypeQuota                             AssociationDefinitionToObjectType = "QUOTA"
	AssociationDefinitionToObjectTypeQuote                             AssociationDefinitionToObjectType = "QUOTE"
	AssociationDefinitionToObjectTypeQuoteField                        AssociationDefinitionToObjectType = "QUOTE_FIELD"
	AssociationDefinitionToObjectTypeQuoteModule                       AssociationDefinitionToObjectType = "QUOTE_MODULE"
	AssociationDefinitionToObjectTypeQuoteModuleField                  AssociationDefinitionToObjectType = "QUOTE_MODULE_FIELD"
	AssociationDefinitionToObjectTypeQuoteTemplate                     AssociationDefinitionToObjectType = "QUOTE_TEMPLATE"
	AssociationDefinitionToObjectTypeRestorableCrmObject               AssociationDefinitionToObjectType = "RESTORABLE_CRM_OBJECT"
	AssociationDefinitionToObjectTypeRoster                            AssociationDefinitionToObjectType = "ROSTER"
	AssociationDefinitionToObjectTypeRosterMember                      AssociationDefinitionToObjectType = "ROSTER_MEMBER"
	AssociationDefinitionToObjectTypeSalesDocument                     AssociationDefinitionToObjectType = "SALES_DOCUMENT"
	AssociationDefinitionToObjectTypeSalesTask                         AssociationDefinitionToObjectType = "SALES_TASK"
	AssociationDefinitionToObjectTypeSalesWorkload                     AssociationDefinitionToObjectType = "SALES_WORKLOAD"
	AssociationDefinitionToObjectTypeSalesforceSyncError               AssociationDefinitionToObjectType = "SALESFORCE_SYNC_ERROR"
	AssociationDefinitionToObjectTypeSchedulingPage                    AssociationDefinitionToObjectType = "SCHEDULING_PAGE"
	AssociationDefinitionToObjectTypeSchemasBackendTest                AssociationDefinitionToObjectType = "SCHEMAS_BACKEND_TEST"
	AssociationDefinitionToObjectTypeScoreConfiguration                AssociationDefinitionToObjectType = "SCORE_CONFIGURATION"
	AssociationDefinitionToObjectTypeSequence                          AssociationDefinitionToObjectType = "SEQUENCE"
	AssociationDefinitionToObjectTypeSequenceEnrollment                AssociationDefinitionToObjectType = "SEQUENCE_ENROLLMENT"
	AssociationDefinitionToObjectTypeSequenceStep                      AssociationDefinitionToObjectType = "SEQUENCE_STEP"
	AssociationDefinitionToObjectTypeSequenceStepEnrollment            AssociationDefinitionToObjectType = "SEQUENCE_STEP_ENROLLMENT"
	AssociationDefinitionToObjectTypeService                           AssociationDefinitionToObjectType = "SERVICE"
	AssociationDefinitionToObjectTypeSitePage                          AssociationDefinitionToObjectType = "SITE_PAGE"
	AssociationDefinitionToObjectTypeSnippet                           AssociationDefinitionToObjectType = "SNIPPET"
	AssociationDefinitionToObjectTypeSocialBroadcast                   AssociationDefinitionToObjectType = "SOCIAL_BROADCAST"
	AssociationDefinitionToObjectTypeSocialChannel                     AssociationDefinitionToObjectType = "SOCIAL_CHANNEL"
	AssociationDefinitionToObjectTypeSocialPost                        AssociationDefinitionToObjectType = "SOCIAL_POST"
	AssociationDefinitionToObjectTypeSocialProfile                     AssociationDefinitionToObjectType = "SOCIAL_PROFILE"
	AssociationDefinitionToObjectTypeSoxProtectedDummyType             AssociationDefinitionToObjectType = "SOX_PROTECTED_DUMMY_TYPE"
	AssociationDefinitionToObjectTypeSoxProtectedTestType              AssociationDefinitionToObjectType = "SOX_PROTECTED_TEST_TYPE"
	AssociationDefinitionToObjectTypeSubmissionTag                     AssociationDefinitionToObjectType = "SUBMISSION_TAG"
	AssociationDefinitionToObjectTypeSubscription                      AssociationDefinitionToObjectType = "SUBSCRIPTION"
	AssociationDefinitionToObjectTypeTask                              AssociationDefinitionToObjectType = "TASK"
	AssociationDefinitionToObjectTypeTaskTemplate                      AssociationDefinitionToObjectType = "TASK_TEMPLATE"
	AssociationDefinitionToObjectTypeTax                               AssociationDefinitionToObjectType = "TAX"
	AssociationDefinitionToObjectTypeTemplate                          AssociationDefinitionToObjectType = "TEMPLATE"
	AssociationDefinitionToObjectTypeTicket                            AssociationDefinitionToObjectType = "TICKET"
	AssociationDefinitionToObjectTypeUnknown                           AssociationDefinitionToObjectType = "UNKNOWN"
	AssociationDefinitionToObjectTypeUnsubscribe                       AssociationDefinitionToObjectType = "UNSUBSCRIBE"
	AssociationDefinitionToObjectTypeUser                              AssociationDefinitionToObjectType = "USER"
	AssociationDefinitionToObjectTypeView                              AssociationDefinitionToObjectType = "VIEW"
	AssociationDefinitionToObjectTypeViewBlock                         AssociationDefinitionToObjectType = "VIEW_BLOCK"
	AssociationDefinitionToObjectTypeWebInteractive                    AssociationDefinitionToObjectType = "WEB_INTERACTIVE"
)

// The property Inputs is required.
type BatchedBehavioralEventHTTPCompletionRequestParam struct {
	Inputs []BehavioralEventHTTPCompletionRequestParam `json:"inputs,omitzero" api:"required"`
	paramObj
}

func (r BatchedBehavioralEventHTTPCompletionRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow BatchedBehavioralEventHTTPCompletionRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BatchedBehavioralEventHTTPCompletionRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties EventName, Properties are required.
type BehavioralEventHTTPCompletionRequestParam struct {
	// The event's fully qualified name. This value (formatted as `pe{HubID}_{name}`)
	// can be retrieved through the
	// [event definitions API](https://developers.hubspot.com/docs/reference/api/analytics-and-events/custom-events/custom-event-definitions#get-%2Fevents%2Fv3%2Fevent-definitions)
	// or in
	// [HubSpot's UI](https://knowledge.hubspot.com/reports/create-custom-behavioral-events-with-the-code-wizard#find-internal-name).
	EventName string `json:"eventName" api:"required"`
	// The event properties to update. Takes the format of key-value pairs (property
	// internal name and property value). Learn more about
	// [HubSpot's default event properties](https://developers.hubspot.com/docs/guides/api/analytics-and-events/custom-events/custom-event-definitions#hubspot-s-default-event-properties).
	Properties map[string]string `json:"properties,omitzero" api:"required"`
	// The visitor's email address. Used for associating the event data with a CRM
	// record.
	Email param.Opt[string] `json:"email,omitzero"`
	// The ID of the record for which the event occurred (e.g., contact ID or visitor
	// ID).
	ObjectID param.Opt[string] `json:"objectId,omitzero"`
	// The time when this event occurred. If this isn't set, the current time will be
	// used.
	OccurredAt param.Opt[time.Time] `json:"occurredAt,omitzero" format:"date-time"`
	// The visitor's usertoken. Used for associating the event data with a CRM record.
	Utk param.Opt[string] `json:"utk,omitzero"`
	// Include a universally unique identifier to assign a unique ID to the event
	// occurrence. Can be useful for matching data between HubSpot and other external
	// systems.
	Uuid param.Opt[string] `json:"uuid,omitzero"`
	paramObj
}

func (r BehavioralEventHTTPCompletionRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow BehavioralEventHTTPCompletionRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BehavioralEventHTTPCompletionRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BehavioralEventTypeDefinitionLabels struct {
	Singular string `json:"singular" api:"required"`
	Plural   string `json:"plural"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Singular    respjson.Field
		Plural      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BehavioralEventTypeDefinitionLabels) RawJSON() string { return r.JSON.raw }
func (r *BehavioralEventTypeDefinitionLabels) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BoolPropertyOperation struct {
	IncludeObjectsWithNoValueSet bool   `json:"includeObjectsWithNoValueSet" api:"required"`
	OperationType                string `json:"operationType" api:"required"`
	// Any of "HAS_EVER_BEEN_EQUAL_TO", "HAS_NEVER_BEEN_EQUAL_TO", "IS_EQUAL_TO",
	// "IS_NOT_EQUAL_TO".
	Operator     BoolPropertyOperationOperator `json:"operator" api:"required"`
	OperatorName string                        `json:"operatorName" api:"required"`
	// Any of "bool".
	PropertyType BoolPropertyOperationPropertyType `json:"propertyType" api:"required"`
	Value        bool                              `json:"value" api:"required"`
	DefaultValue string                            `json:"defaultValue"`
	RenderSpec   string                            `json:"renderSpec"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		IncludeObjectsWithNoValueSet respjson.Field
		OperationType                respjson.Field
		Operator                     respjson.Field
		OperatorName                 respjson.Field
		PropertyType                 respjson.Field
		Value                        respjson.Field
		DefaultValue                 respjson.Field
		RenderSpec                   respjson.Field
		ExtraFields                  map[string]respjson.Field
		raw                          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BoolPropertyOperation) RawJSON() string { return r.JSON.raw }
func (r *BoolPropertyOperation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BoolPropertyOperationOperator string

const (
	BoolPropertyOperationOperatorHasEverBeenEqualTo  BoolPropertyOperationOperator = "HAS_EVER_BEEN_EQUAL_TO"
	BoolPropertyOperationOperatorHasNeverBeenEqualTo BoolPropertyOperationOperator = "HAS_NEVER_BEEN_EQUAL_TO"
	BoolPropertyOperationOperatorIsEqualTo           BoolPropertyOperationOperator = "IS_EQUAL_TO"
	BoolPropertyOperationOperatorIsNotEqualTo        BoolPropertyOperationOperator = "IS_NOT_EQUAL_TO"
)

type BoolPropertyOperationPropertyType string

const (
	BoolPropertyOperationPropertyTypeBool BoolPropertyOperationPropertyType = "bool"
)

type CalendarDatePropertyOperation struct {
	IncludeObjectsWithNoValueSet bool   `json:"includeObjectsWithNoValueSet" api:"required"`
	OperationType                string `json:"operationType" api:"required"`
	// Any of "IN_LAST_TIME_UNIT", "IN_NEXT_TIME_UNIT", "IN_THIS_TIME_UNIT",
	// "IN_THIS_TIME_UNIT_SO_FAR".
	Operator     CalendarDatePropertyOperationOperator `json:"operator" api:"required"`
	OperatorName string                                `json:"operatorName" api:"required"`
	// Any of "calendar-date".
	PropertyType CalendarDatePropertyOperationPropertyType `json:"propertyType" api:"required"`
	// Any of "DAY", "MONTH", "QUARTER", "WEEK", "YEAR".
	TimeUnit      CalendarDatePropertyOperationTimeUnit `json:"timeUnit" api:"required"`
	TimeUnitCount int64                                 `json:"timeUnitCount" api:"required"`
	UseFiscalYear bool                                  `json:"useFiscalYear" api:"required"`
	DefaultValue  string                                `json:"defaultValue"`
	// Any of "APRIL", "AUGUST", "DECEMBER", "FEBRUARY", "JANUARY", "JULY", "JUNE",
	// "MARCH", "MAY", "NOVEMBER", "OCTOBER", "SEPTEMBER".
	FiscalYearStart CalendarDatePropertyOperationFiscalYearStart `json:"fiscalYearStart"`
	RenderSpec      string                                       `json:"renderSpec"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		IncludeObjectsWithNoValueSet respjson.Field
		OperationType                respjson.Field
		Operator                     respjson.Field
		OperatorName                 respjson.Field
		PropertyType                 respjson.Field
		TimeUnit                     respjson.Field
		TimeUnitCount                respjson.Field
		UseFiscalYear                respjson.Field
		DefaultValue                 respjson.Field
		FiscalYearStart              respjson.Field
		RenderSpec                   respjson.Field
		ExtraFields                  map[string]respjson.Field
		raw                          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CalendarDatePropertyOperation) RawJSON() string { return r.JSON.raw }
func (r *CalendarDatePropertyOperation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CalendarDatePropertyOperationOperator string

const (
	CalendarDatePropertyOperationOperatorInLastTimeUnit      CalendarDatePropertyOperationOperator = "IN_LAST_TIME_UNIT"
	CalendarDatePropertyOperationOperatorInNextTimeUnit      CalendarDatePropertyOperationOperator = "IN_NEXT_TIME_UNIT"
	CalendarDatePropertyOperationOperatorInThisTimeUnit      CalendarDatePropertyOperationOperator = "IN_THIS_TIME_UNIT"
	CalendarDatePropertyOperationOperatorInThisTimeUnitSoFar CalendarDatePropertyOperationOperator = "IN_THIS_TIME_UNIT_SO_FAR"
)

type CalendarDatePropertyOperationPropertyType string

const (
	CalendarDatePropertyOperationPropertyTypeCalendarDate CalendarDatePropertyOperationPropertyType = "calendar-date"
)

type CalendarDatePropertyOperationTimeUnit string

const (
	CalendarDatePropertyOperationTimeUnitDay     CalendarDatePropertyOperationTimeUnit = "DAY"
	CalendarDatePropertyOperationTimeUnitMonth   CalendarDatePropertyOperationTimeUnit = "MONTH"
	CalendarDatePropertyOperationTimeUnitQuarter CalendarDatePropertyOperationTimeUnit = "QUARTER"
	CalendarDatePropertyOperationTimeUnitWeek    CalendarDatePropertyOperationTimeUnit = "WEEK"
	CalendarDatePropertyOperationTimeUnitYear    CalendarDatePropertyOperationTimeUnit = "YEAR"
)

type CalendarDatePropertyOperationFiscalYearStart string

const (
	CalendarDatePropertyOperationFiscalYearStartApril     CalendarDatePropertyOperationFiscalYearStart = "APRIL"
	CalendarDatePropertyOperationFiscalYearStartAugust    CalendarDatePropertyOperationFiscalYearStart = "AUGUST"
	CalendarDatePropertyOperationFiscalYearStartDecember  CalendarDatePropertyOperationFiscalYearStart = "DECEMBER"
	CalendarDatePropertyOperationFiscalYearStartFebruary  CalendarDatePropertyOperationFiscalYearStart = "FEBRUARY"
	CalendarDatePropertyOperationFiscalYearStartJanuary   CalendarDatePropertyOperationFiscalYearStart = "JANUARY"
	CalendarDatePropertyOperationFiscalYearStartJuly      CalendarDatePropertyOperationFiscalYearStart = "JULY"
	CalendarDatePropertyOperationFiscalYearStartJune      CalendarDatePropertyOperationFiscalYearStart = "JUNE"
	CalendarDatePropertyOperationFiscalYearStartMarch     CalendarDatePropertyOperationFiscalYearStart = "MARCH"
	CalendarDatePropertyOperationFiscalYearStartMay       CalendarDatePropertyOperationFiscalYearStart = "MAY"
	CalendarDatePropertyOperationFiscalYearStartNovember  CalendarDatePropertyOperationFiscalYearStart = "NOVEMBER"
	CalendarDatePropertyOperationFiscalYearStartOctober   CalendarDatePropertyOperationFiscalYearStart = "OCTOBER"
	CalendarDatePropertyOperationFiscalYearStartSeptember CalendarDatePropertyOperationFiscalYearStart = "SEPTEMBER"
)

type CollectionResponseWithTotalExternalBehavioralEventTypeDefinition struct {
	Results []ExternalBehavioralEventTypeDefinition `json:"results" api:"required"`
	Total   int64                                   `json:"total" api:"required"`
	Paging  shared.Paging                           `json:"paging"`
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
func (r CollectionResponseWithTotalExternalBehavioralEventTypeDefinition) RawJSON() string {
	return r.JSON.raw
}
func (r *CollectionResponseWithTotalExternalBehavioralEventTypeDefinition) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ComboEventRule struct {
	Count              int64            `json:"count" api:"required"`
	EventTypeID        string           `json:"eventTypeId" api:"required"`
	PropertyFilters    []PropertyFilter `json:"propertyFilters" api:"required"`
	LookbackWindowDays int64            `json:"lookbackWindowDays"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Count              respjson.Field
		EventTypeID        respjson.Field
		PropertyFilters    respjson.Field
		LookbackWindowDays respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ComboEventRule) RawJSON() string { return r.JSON.raw }
func (r *ComboEventRule) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ComboEventRuleBranch struct {
	ComposingRules []ComboEventRule `json:"composingRules" api:"required"`
	// Any of "AND", "OR".
	OperationType ComboEventRuleBranchOperationType `json:"operationType" api:"required"`
	RuleBranches  []ComboEventRuleBranch            `json:"ruleBranches" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ComposingRules respjson.Field
		OperationType  respjson.Field
		RuleBranches   respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ComboEventRuleBranch) RawJSON() string { return r.JSON.raw }
func (r *ComboEventRuleBranch) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ComboEventRuleBranchOperationType string

const (
	ComboEventRuleBranchOperationTypeAnd ComboEventRuleBranchOperationType = "AND"
	ComboEventRuleBranchOperationTypeOr  ComboEventRuleBranchOperationType = "OR"
)

type ComparativeBoolPropertyOperation struct {
	ComparisonPropertyName       string `json:"comparisonPropertyName" api:"required"`
	IncludeObjectsWithNoValueSet bool   `json:"includeObjectsWithNoValueSet" api:"required"`
	OperationType                string `json:"operationType" api:"required"`
	// Any of "IS_EQUAL_TO", "IS_NOT_EQUAL_TO".
	Operator     ComparativeBoolPropertyOperationOperator `json:"operator" api:"required"`
	OperatorName string                                   `json:"operatorName" api:"required"`
	// Any of "bool-comparative".
	PropertyType ComparativeBoolPropertyOperationPropertyType `json:"propertyType" api:"required"`
	DefaultValue string                                       `json:"defaultValue"`
	RenderSpec   string                                       `json:"renderSpec"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ComparisonPropertyName       respjson.Field
		IncludeObjectsWithNoValueSet respjson.Field
		OperationType                respjson.Field
		Operator                     respjson.Field
		OperatorName                 respjson.Field
		PropertyType                 respjson.Field
		DefaultValue                 respjson.Field
		RenderSpec                   respjson.Field
		ExtraFields                  map[string]respjson.Field
		raw                          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ComparativeBoolPropertyOperation) RawJSON() string { return r.JSON.raw }
func (r *ComparativeBoolPropertyOperation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ComparativeBoolPropertyOperationOperator string

const (
	ComparativeBoolPropertyOperationOperatorIsEqualTo    ComparativeBoolPropertyOperationOperator = "IS_EQUAL_TO"
	ComparativeBoolPropertyOperationOperatorIsNotEqualTo ComparativeBoolPropertyOperationOperator = "IS_NOT_EQUAL_TO"
)

type ComparativeBoolPropertyOperationPropertyType string

const (
	ComparativeBoolPropertyOperationPropertyTypeBoolComparative ComparativeBoolPropertyOperationPropertyType = "bool-comparative"
)

type ComparativeDatePropertyOperation struct {
	ComparisonPropertyName       string `json:"comparisonPropertyName" api:"required"`
	IncludeObjectsWithNoValueSet bool   `json:"includeObjectsWithNoValueSet" api:"required"`
	OperationType                string `json:"operationType" api:"required"`
	// Any of "IS_AFTER", "IS_BEFORE".
	Operator     ComparativeDatePropertyOperationOperator `json:"operator" api:"required"`
	OperatorName string                                   `json:"operatorName" api:"required"`
	// Any of "datetime-comparative".
	PropertyType           ComparativeDatePropertyOperationPropertyType `json:"propertyType" api:"required"`
	DefaultComparisonValue string                                       `json:"defaultComparisonValue"`
	DefaultValue           string                                       `json:"defaultValue"`
	RenderSpec             string                                       `json:"renderSpec"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ComparisonPropertyName       respjson.Field
		IncludeObjectsWithNoValueSet respjson.Field
		OperationType                respjson.Field
		Operator                     respjson.Field
		OperatorName                 respjson.Field
		PropertyType                 respjson.Field
		DefaultComparisonValue       respjson.Field
		DefaultValue                 respjson.Field
		RenderSpec                   respjson.Field
		ExtraFields                  map[string]respjson.Field
		raw                          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ComparativeDatePropertyOperation) RawJSON() string { return r.JSON.raw }
func (r *ComparativeDatePropertyOperation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ComparativeDatePropertyOperationOperator string

const (
	ComparativeDatePropertyOperationOperatorIsAfter  ComparativeDatePropertyOperationOperator = "IS_AFTER"
	ComparativeDatePropertyOperationOperatorIsBefore ComparativeDatePropertyOperationOperator = "IS_BEFORE"
)

type ComparativeDatePropertyOperationPropertyType string

const (
	ComparativeDatePropertyOperationPropertyTypeDatetimeComparative ComparativeDatePropertyOperationPropertyType = "datetime-comparative"
)

type ComparativeNumberPropertyOperation struct {
	ComparisonPropertyName       string `json:"comparisonPropertyName" api:"required"`
	IncludeObjectsWithNoValueSet bool   `json:"includeObjectsWithNoValueSet" api:"required"`
	OperationType                string `json:"operationType" api:"required"`
	// Any of "IS_EQUAL_TO", "IS_GREATER_THAN", "IS_GREATER_THAN_OR_EQUAL_TO",
	// "IS_LESS_THAN", "IS_LESS_THAN_OR_EQUAL_TO", "IS_NOT_EQUAL_TO".
	Operator     ComparativeNumberPropertyOperationOperator `json:"operator" api:"required"`
	OperatorName string                                     `json:"operatorName" api:"required"`
	// Any of "number-comparative".
	PropertyType ComparativeNumberPropertyOperationPropertyType `json:"propertyType" api:"required"`
	DefaultValue string                                         `json:"defaultValue"`
	RenderSpec   string                                         `json:"renderSpec"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ComparisonPropertyName       respjson.Field
		IncludeObjectsWithNoValueSet respjson.Field
		OperationType                respjson.Field
		Operator                     respjson.Field
		OperatorName                 respjson.Field
		PropertyType                 respjson.Field
		DefaultValue                 respjson.Field
		RenderSpec                   respjson.Field
		ExtraFields                  map[string]respjson.Field
		raw                          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ComparativeNumberPropertyOperation) RawJSON() string { return r.JSON.raw }
func (r *ComparativeNumberPropertyOperation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ComparativeNumberPropertyOperationOperator string

const (
	ComparativeNumberPropertyOperationOperatorIsEqualTo              ComparativeNumberPropertyOperationOperator = "IS_EQUAL_TO"
	ComparativeNumberPropertyOperationOperatorIsGreaterThan          ComparativeNumberPropertyOperationOperator = "IS_GREATER_THAN"
	ComparativeNumberPropertyOperationOperatorIsGreaterThanOrEqualTo ComparativeNumberPropertyOperationOperator = "IS_GREATER_THAN_OR_EQUAL_TO"
	ComparativeNumberPropertyOperationOperatorIsLessThan             ComparativeNumberPropertyOperationOperator = "IS_LESS_THAN"
	ComparativeNumberPropertyOperationOperatorIsLessThanOrEqualTo    ComparativeNumberPropertyOperationOperator = "IS_LESS_THAN_OR_EQUAL_TO"
	ComparativeNumberPropertyOperationOperatorIsNotEqualTo           ComparativeNumberPropertyOperationOperator = "IS_NOT_EQUAL_TO"
)

type ComparativeNumberPropertyOperationPropertyType string

const (
	ComparativeNumberPropertyOperationPropertyTypeNumberComparative ComparativeNumberPropertyOperationPropertyType = "number-comparative"
)

type ComparativePropertyUpdatedOperation struct {
	ComparisonPropertyName       string `json:"comparisonPropertyName" api:"required"`
	IncludeObjectsWithNoValueSet bool   `json:"includeObjectsWithNoValueSet" api:"required"`
	OperationType                string `json:"operationType" api:"required"`
	// Any of "IS_AFTER", "IS_BEFORE".
	Operator     ComparativePropertyUpdatedOperationOperator `json:"operator" api:"required"`
	OperatorName string                                      `json:"operatorName" api:"required"`
	// Any of "property-updated-comparative".
	PropertyType           ComparativePropertyUpdatedOperationPropertyType `json:"propertyType" api:"required"`
	DefaultComparisonValue string                                          `json:"defaultComparisonValue"`
	DefaultValue           string                                          `json:"defaultValue"`
	RenderSpec             string                                          `json:"renderSpec"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ComparisonPropertyName       respjson.Field
		IncludeObjectsWithNoValueSet respjson.Field
		OperationType                respjson.Field
		Operator                     respjson.Field
		OperatorName                 respjson.Field
		PropertyType                 respjson.Field
		DefaultComparisonValue       respjson.Field
		DefaultValue                 respjson.Field
		RenderSpec                   respjson.Field
		ExtraFields                  map[string]respjson.Field
		raw                          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ComparativePropertyUpdatedOperation) RawJSON() string { return r.JSON.raw }
func (r *ComparativePropertyUpdatedOperation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ComparativePropertyUpdatedOperationOperator string

const (
	ComparativePropertyUpdatedOperationOperatorIsAfter  ComparativePropertyUpdatedOperationOperator = "IS_AFTER"
	ComparativePropertyUpdatedOperationOperatorIsBefore ComparativePropertyUpdatedOperationOperator = "IS_BEFORE"
)

type ComparativePropertyUpdatedOperationPropertyType string

const (
	ComparativePropertyUpdatedOperationPropertyTypePropertyUpdatedComparative ComparativePropertyUpdatedOperationPropertyType = "property-updated-comparative"
)

type ComparativeStringPropertyOperation struct {
	ComparisonPropertyName       string `json:"comparisonPropertyName" api:"required"`
	IncludeObjectsWithNoValueSet bool   `json:"includeObjectsWithNoValueSet" api:"required"`
	OperationType                string `json:"operationType" api:"required"`
	// Any of "CONTAINS", "DOES_NOT_CONTAIN", "ENDS_WITH", "IS_EQUAL_TO",
	// "IS_NOT_EQUAL_TO", "STARTS_WITH".
	Operator     ComparativeStringPropertyOperationOperator `json:"operator" api:"required"`
	OperatorName string                                     `json:"operatorName" api:"required"`
	// Any of "string-comparative".
	PropertyType ComparativeStringPropertyOperationPropertyType `json:"propertyType" api:"required"`
	DefaultValue string                                         `json:"defaultValue"`
	RenderSpec   string                                         `json:"renderSpec"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ComparisonPropertyName       respjson.Field
		IncludeObjectsWithNoValueSet respjson.Field
		OperationType                respjson.Field
		Operator                     respjson.Field
		OperatorName                 respjson.Field
		PropertyType                 respjson.Field
		DefaultValue                 respjson.Field
		RenderSpec                   respjson.Field
		ExtraFields                  map[string]respjson.Field
		raw                          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ComparativeStringPropertyOperation) RawJSON() string { return r.JSON.raw }
func (r *ComparativeStringPropertyOperation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ComparativeStringPropertyOperationOperator string

const (
	ComparativeStringPropertyOperationOperatorContains       ComparativeStringPropertyOperationOperator = "CONTAINS"
	ComparativeStringPropertyOperationOperatorDoesNotContain ComparativeStringPropertyOperationOperator = "DOES_NOT_CONTAIN"
	ComparativeStringPropertyOperationOperatorEndsWith       ComparativeStringPropertyOperationOperator = "ENDS_WITH"
	ComparativeStringPropertyOperationOperatorIsEqualTo      ComparativeStringPropertyOperationOperator = "IS_EQUAL_TO"
	ComparativeStringPropertyOperationOperatorIsNotEqualTo   ComparativeStringPropertyOperationOperator = "IS_NOT_EQUAL_TO"
	ComparativeStringPropertyOperationOperatorStartsWith     ComparativeStringPropertyOperationOperator = "STARTS_WITH"
)

type ComparativeStringPropertyOperationPropertyType string

const (
	ComparativeStringPropertyOperationPropertyTypeStringComparative ComparativeStringPropertyOperationPropertyType = "string-comparative"
)

type DatePoint struct {
	Day   int64 `json:"day" api:"required"`
	Month int64 `json:"month" api:"required"`
	// Any of "DATE".
	TimeType DatePointTimeType `json:"timeType" api:"required"`
	// Any of "CUSTOM", "PORTAL", "USER".
	TimezoneSource DatePointTimezoneSource `json:"timezoneSource" api:"required"`
	Year           int64                   `json:"year" api:"required"`
	ZoneID         string                  `json:"zoneId" api:"required"`
	Hour           int64                   `json:"hour"`
	Millisecond    int64                   `json:"millisecond"`
	Minute         int64                   `json:"minute"`
	Second         int64                   `json:"second"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Day            respjson.Field
		Month          respjson.Field
		TimeType       respjson.Field
		TimezoneSource respjson.Field
		Year           respjson.Field
		ZoneID         respjson.Field
		Hour           respjson.Field
		Millisecond    respjson.Field
		Minute         respjson.Field
		Second         respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DatePoint) RawJSON() string { return r.JSON.raw }
func (r *DatePoint) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DatePointTimeType string

const (
	DatePointTimeTypeDate DatePointTimeType = "DATE"
)

type DatePointTimezoneSource string

const (
	DatePointTimezoneSourceCustom DatePointTimezoneSource = "CUSTOM"
	DatePointTimezoneSourcePortal DatePointTimezoneSource = "PORTAL"
	DatePointTimezoneSourceUser   DatePointTimezoneSource = "USER"
)

type DatePropertyOperation struct {
	Day                          int64 `json:"day" api:"required"`
	IncludeObjectsWithNoValueSet bool  `json:"includeObjectsWithNoValueSet" api:"required"`
	// Any of "APR", "AUG", "DEC", "FEB", "JAN", "JUL", "JUN", "MAR", "MAY", "NOV",
	// "OCT", "SEP".
	Month         DatePropertyOperationMonth `json:"month" api:"required"`
	OperationType string                     `json:"operationType" api:"required"`
	// Any of "AFTER", "BEFORE", "EQUAL".
	Operator     DatePropertyOperationOperator `json:"operator" api:"required"`
	OperatorName string                        `json:"operatorName" api:"required"`
	// Any of "date".
	PropertyType DatePropertyOperationPropertyType `json:"propertyType" api:"required"`
	Year         int64                             `json:"year" api:"required"`
	DefaultValue string                            `json:"defaultValue"`
	RenderSpec   string                            `json:"renderSpec"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Day                          respjson.Field
		IncludeObjectsWithNoValueSet respjson.Field
		Month                        respjson.Field
		OperationType                respjson.Field
		Operator                     respjson.Field
		OperatorName                 respjson.Field
		PropertyType                 respjson.Field
		Year                         respjson.Field
		DefaultValue                 respjson.Field
		RenderSpec                   respjson.Field
		ExtraFields                  map[string]respjson.Field
		raw                          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DatePropertyOperation) RawJSON() string { return r.JSON.raw }
func (r *DatePropertyOperation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DatePropertyOperationMonth string

const (
	DatePropertyOperationMonthApr DatePropertyOperationMonth = "APR"
	DatePropertyOperationMonthAug DatePropertyOperationMonth = "AUG"
	DatePropertyOperationMonthDec DatePropertyOperationMonth = "DEC"
	DatePropertyOperationMonthFeb DatePropertyOperationMonth = "FEB"
	DatePropertyOperationMonthJan DatePropertyOperationMonth = "JAN"
	DatePropertyOperationMonthJul DatePropertyOperationMonth = "JUL"
	DatePropertyOperationMonthJun DatePropertyOperationMonth = "JUN"
	DatePropertyOperationMonthMar DatePropertyOperationMonth = "MAR"
	DatePropertyOperationMonthMay DatePropertyOperationMonth = "MAY"
	DatePropertyOperationMonthNov DatePropertyOperationMonth = "NOV"
	DatePropertyOperationMonthOct DatePropertyOperationMonth = "OCT"
	DatePropertyOperationMonthSep DatePropertyOperationMonth = "SEP"
)

type DatePropertyOperationOperator string

const (
	DatePropertyOperationOperatorAfter  DatePropertyOperationOperator = "AFTER"
	DatePropertyOperationOperatorBefore DatePropertyOperationOperator = "BEFORE"
	DatePropertyOperationOperatorEqual  DatePropertyOperationOperator = "EQUAL"
)

type DatePropertyOperationPropertyType string

const (
	DatePropertyOperationPropertyTypeDate DatePropertyOperationPropertyType = "date"
)

type DateTimePropertyOperation struct {
	IncludeObjectsWithNoValueSet bool   `json:"includeObjectsWithNoValueSet" api:"required"`
	OperationType                string `json:"operationType" api:"required"`
	// Any of "IS_AFTER", "IS_AFTER_DATE", "IS_BEFORE", "IS_BEFORE_DATE",
	// "IS_EQUAL_TO".
	Operator     DateTimePropertyOperationOperator `json:"operator" api:"required"`
	OperatorName string                            `json:"operatorName" api:"required"`
	// Any of "datetime".
	PropertyType               DateTimePropertyOperationPropertyType `json:"propertyType" api:"required"`
	RequiresTimeZoneConversion bool                                  `json:"requiresTimeZoneConversion" api:"required"`
	Timestamp                  int64                                 `json:"timestamp" api:"required"`
	DefaultValue               string                                `json:"defaultValue"`
	RenderSpec                 string                                `json:"renderSpec"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		IncludeObjectsWithNoValueSet respjson.Field
		OperationType                respjson.Field
		Operator                     respjson.Field
		OperatorName                 respjson.Field
		PropertyType                 respjson.Field
		RequiresTimeZoneConversion   respjson.Field
		Timestamp                    respjson.Field
		DefaultValue                 respjson.Field
		RenderSpec                   respjson.Field
		ExtraFields                  map[string]respjson.Field
		raw                          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DateTimePropertyOperation) RawJSON() string { return r.JSON.raw }
func (r *DateTimePropertyOperation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DateTimePropertyOperationOperator string

const (
	DateTimePropertyOperationOperatorIsAfter      DateTimePropertyOperationOperator = "IS_AFTER"
	DateTimePropertyOperationOperatorIsAfterDate  DateTimePropertyOperationOperator = "IS_AFTER_DATE"
	DateTimePropertyOperationOperatorIsBefore     DateTimePropertyOperationOperator = "IS_BEFORE"
	DateTimePropertyOperationOperatorIsBeforeDate DateTimePropertyOperationOperator = "IS_BEFORE_DATE"
	DateTimePropertyOperationOperatorIsEqualTo    DateTimePropertyOperationOperator = "IS_EQUAL_TO"
)

type DateTimePropertyOperationPropertyType string

const (
	DateTimePropertyOperationPropertyTypeDatetime DateTimePropertyOperationPropertyType = "datetime"
)

type EnumerationPropertyOperation struct {
	IncludeObjectsWithNoValueSet bool   `json:"includeObjectsWithNoValueSet" api:"required"`
	OperationType                string `json:"operationType" api:"required"`
	// Any of "CONTAINS_ALL", "DOES_NOT_CONTAIN_ALL", "HAS_EVER_BEEN_ANY_OF",
	// "HAS_EVER_BEEN_EXACTLY", "HAS_EVER_CONTAINED_ALL", "HAS_NEVER_BEEN_ANY_OF",
	// "HAS_NEVER_BEEN_EXACTLY", "HAS_NEVER_CONTAINED_ALL", "IS_ANY_OF", "IS_EXACTLY",
	// "IS_NONE_OF", "IS_NOT_EXACTLY".
	Operator     EnumerationPropertyOperationOperator `json:"operator" api:"required"`
	OperatorName string                               `json:"operatorName" api:"required"`
	// Any of "enumeration".
	PropertyType EnumerationPropertyOperationPropertyType `json:"propertyType" api:"required"`
	Values       []string                                 `json:"values" api:"required"`
	DefaultValue string                                   `json:"defaultValue"`
	RenderSpec   string                                   `json:"renderSpec"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		IncludeObjectsWithNoValueSet respjson.Field
		OperationType                respjson.Field
		Operator                     respjson.Field
		OperatorName                 respjson.Field
		PropertyType                 respjson.Field
		Values                       respjson.Field
		DefaultValue                 respjson.Field
		RenderSpec                   respjson.Field
		ExtraFields                  map[string]respjson.Field
		raw                          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EnumerationPropertyOperation) RawJSON() string { return r.JSON.raw }
func (r *EnumerationPropertyOperation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EnumerationPropertyOperationOperator string

const (
	EnumerationPropertyOperationOperatorContainsAll          EnumerationPropertyOperationOperator = "CONTAINS_ALL"
	EnumerationPropertyOperationOperatorDoesNotContainAll    EnumerationPropertyOperationOperator = "DOES_NOT_CONTAIN_ALL"
	EnumerationPropertyOperationOperatorHasEverBeenAnyOf     EnumerationPropertyOperationOperator = "HAS_EVER_BEEN_ANY_OF"
	EnumerationPropertyOperationOperatorHasEverBeenExactly   EnumerationPropertyOperationOperator = "HAS_EVER_BEEN_EXACTLY"
	EnumerationPropertyOperationOperatorHasEverContainedAll  EnumerationPropertyOperationOperator = "HAS_EVER_CONTAINED_ALL"
	EnumerationPropertyOperationOperatorHasNeverBeenAnyOf    EnumerationPropertyOperationOperator = "HAS_NEVER_BEEN_ANY_OF"
	EnumerationPropertyOperationOperatorHasNeverBeenExactly  EnumerationPropertyOperationOperator = "HAS_NEVER_BEEN_EXACTLY"
	EnumerationPropertyOperationOperatorHasNeverContainedAll EnumerationPropertyOperationOperator = "HAS_NEVER_CONTAINED_ALL"
	EnumerationPropertyOperationOperatorIsAnyOf              EnumerationPropertyOperationOperator = "IS_ANY_OF"
	EnumerationPropertyOperationOperatorIsExactly            EnumerationPropertyOperationOperator = "IS_EXACTLY"
	EnumerationPropertyOperationOperatorIsNoneOf             EnumerationPropertyOperationOperator = "IS_NONE_OF"
	EnumerationPropertyOperationOperatorIsNotExactly         EnumerationPropertyOperationOperator = "IS_NOT_EXACTLY"
)

type EnumerationPropertyOperationPropertyType string

const (
	EnumerationPropertyOperationPropertyTypeEnumeration EnumerationPropertyOperationPropertyType = "enumeration"
)

// The properties Label, Type are required.
type ExternalBehavioralEventPropertyCreateParam struct {
	// Human readable label for the property. Used in HubSpot UI
	Label string `json:"label" api:"required"`
	// The data type of the property. Can be one of the following: [string, number,
	// enumeration, datetime]
	Type string `json:"type" api:"required"`
	// A description of the property that will be shown as help text in HubSpot.
	Description param.Opt[string] `json:"description,omitzero"`
	// Internal property name, which must be used when referencing the property from
	// the API
	Name param.Opt[string] `json:"name,omitzero"`
	// A list of available options for the property if it is an enumeration. NOTE: This
	// field is only applicable for enumerated properties.
	Options []OptionInputParam `json:"options,omitzero"`
	paramObj
}

func (r ExternalBehavioralEventPropertyCreateParam) MarshalJSON() (data []byte, err error) {
	type shadow ExternalBehavioralEventPropertyCreateParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExternalBehavioralEventPropertyCreateParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExternalBehavioralEventPropertyDefinitionPatchParam struct {
	// A description of the property that will be shown as help text in HubSpot.
	Description param.Opt[string] `json:"description,omitzero"`
	// Human readable label for the property. Used in HubSpot UI
	Label param.Opt[string] `json:"label,omitzero"`
	// A list of available options for the property if it is an enumeration. NOTE: This
	// field is only applicable for enumerated properties.
	Options []OptionInputParam `json:"options,omitzero"`
	paramObj
}

func (r ExternalBehavioralEventPropertyDefinitionPatchParam) MarshalJSON() (data []byte, err error) {
	type shadow ExternalBehavioralEventPropertyDefinitionPatchParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExternalBehavioralEventPropertyDefinitionPatchParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExternalBehavioralEventTypeDefinition struct {
	ID                 string                                  `json:"id" api:"required"`
	Archived           bool                                    `json:"archived" api:"required"`
	Associations       []AssociationDefinition                 `json:"associations" api:"required"`
	FullyQualifiedName string                                  `json:"fullyQualifiedName" api:"required"`
	Labels             BehavioralEventTypeDefinitionLabels     `json:"labels" api:"required"`
	Name               string                                  `json:"name" api:"required"`
	ObjectTypeID       string                                  `json:"objectTypeId" api:"required"`
	Properties         []Property                              `json:"properties" api:"required"`
	ComboEventRules    ComboEventRuleBranch                    `json:"comboEventRules"`
	CreatedAt          time.Time                               `json:"createdAt" format:"date-time"`
	CreatedUserID      int64                                   `json:"createdUserId"`
	CustomMatchingID   ExternalObjectResolutionMappingResponse `json:"customMatchingId"`
	Description        string                                  `json:"description"`
	PrimaryObject      string                                  `json:"primaryObject"`
	PrimaryObjectID    string                                  `json:"primaryObjectId"`
	// Any of "APP_EVENT", "AUTOCAPTURE_EVENT", "CLICKED_ELEMENT", "COMBO_EVENT",
	// "CUSTOM_SCRIPT", "IMPORT", "MANUAL", "PROPERTY_CHANGE", "VISITED_URL",
	// "WEBHOOK".
	TrackingType  ExternalBehavioralEventTypeDefinitionTrackingType `json:"trackingType"`
	UpdatedAt     time.Time                                         `json:"updatedAt" format:"date-time"`
	UpdatedUserID int64                                             `json:"updatedUserId"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                 respjson.Field
		Archived           respjson.Field
		Associations       respjson.Field
		FullyQualifiedName respjson.Field
		Labels             respjson.Field
		Name               respjson.Field
		ObjectTypeID       respjson.Field
		Properties         respjson.Field
		ComboEventRules    respjson.Field
		CreatedAt          respjson.Field
		CreatedUserID      respjson.Field
		CustomMatchingID   respjson.Field
		Description        respjson.Field
		PrimaryObject      respjson.Field
		PrimaryObjectID    respjson.Field
		TrackingType       respjson.Field
		UpdatedAt          respjson.Field
		UpdatedUserID      respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExternalBehavioralEventTypeDefinition) RawJSON() string { return r.JSON.raw }
func (r *ExternalBehavioralEventTypeDefinition) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExternalBehavioralEventTypeDefinitionTrackingType string

const (
	ExternalBehavioralEventTypeDefinitionTrackingTypeAppEvent         ExternalBehavioralEventTypeDefinitionTrackingType = "APP_EVENT"
	ExternalBehavioralEventTypeDefinitionTrackingTypeAutocaptureEvent ExternalBehavioralEventTypeDefinitionTrackingType = "AUTOCAPTURE_EVENT"
	ExternalBehavioralEventTypeDefinitionTrackingTypeClickedElement   ExternalBehavioralEventTypeDefinitionTrackingType = "CLICKED_ELEMENT"
	ExternalBehavioralEventTypeDefinitionTrackingTypeComboEvent       ExternalBehavioralEventTypeDefinitionTrackingType = "COMBO_EVENT"
	ExternalBehavioralEventTypeDefinitionTrackingTypeCustomScript     ExternalBehavioralEventTypeDefinitionTrackingType = "CUSTOM_SCRIPT"
	ExternalBehavioralEventTypeDefinitionTrackingTypeImport           ExternalBehavioralEventTypeDefinitionTrackingType = "IMPORT"
	ExternalBehavioralEventTypeDefinitionTrackingTypeManual           ExternalBehavioralEventTypeDefinitionTrackingType = "MANUAL"
	ExternalBehavioralEventTypeDefinitionTrackingTypePropertyChange   ExternalBehavioralEventTypeDefinitionTrackingType = "PROPERTY_CHANGE"
	ExternalBehavioralEventTypeDefinitionTrackingTypeVisitedURL       ExternalBehavioralEventTypeDefinitionTrackingType = "VISITED_URL"
	ExternalBehavioralEventTypeDefinitionTrackingTypeWebhook          ExternalBehavioralEventTypeDefinitionTrackingType = "WEBHOOK"
)

// The properties IncludeDefaultProperties, Label, PropertyDefinitions are
// required.
type ExternalBehavioralEventTypeDefinitionEggParam struct {
	IncludeDefaultProperties bool `json:"includeDefaultProperties" api:"required"`
	// Human readable label for the event for display in HubSpot's UI.
	Label string `json:"label" api:"required"`
	// List of custom properties on event
	PropertyDefinitions []ExternalBehavioralEventPropertyCreateParam `json:"propertyDefinitions,omitzero" api:"required"`
	// A description of the event that will be shown as help text in HubSpot.
	Description param.Opt[string] `json:"description,omitzero"`
	// Internal event name, which must be used when referencing the event from the API.
	// If a name is not supplied, one will be generated based on the label. The name
	// does not include the `pe<PORTAL_ID>_` prefix used when sending event
	// completions.
	Name param.Opt[string] `json:"name,omitzero"`
	// The object type to associate this event to. Can be one of `CONTACT`, `COMPANY`,
	// `DEAL`, `TICKET`. If no value is supplied, will default to `CONTACT`.
	PrimaryObject    param.Opt[string]                           `json:"primaryObject,omitzero"`
	CustomMatchingID ExternalObjectResolutionMappingRequestParam `json:"customMatchingId,omitzero"`
	paramObj
}

func (r ExternalBehavioralEventTypeDefinitionEggParam) MarshalJSON() (data []byte, err error) {
	type shadow ExternalBehavioralEventTypeDefinitionEggParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExternalBehavioralEventTypeDefinitionEggParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExternalBehavioralEventTypeDefinitionPatchParam struct {
	// A description of the event that will be shown as help text in HubSpot.
	Description param.Opt[string] `json:"description,omitzero"`
	// Human readable label for the event. Used in HubSpot UI
	Label param.Opt[string] `json:"label,omitzero"`
	paramObj
}

func (r ExternalBehavioralEventTypeDefinitionPatchParam) MarshalJSON() (data []byte, err error) {
	type shadow ExternalBehavioralEventTypeDefinitionPatchParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExternalBehavioralEventTypeDefinitionPatchParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property PrimaryObjectRule is required.
type ExternalObjectResolutionMappingRequestParam struct {
	PrimaryObjectRule ExternalPrimaryObjectResolutionRuleParam `json:"primaryObjectRule,omitzero" api:"required"`
	paramObj
}

func (r ExternalObjectResolutionMappingRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow ExternalObjectResolutionMappingRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExternalObjectResolutionMappingRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExternalObjectResolutionMappingResponse struct {
	PrimaryObjectRule ExternalPrimaryObjectResolutionRule `json:"primaryObjectRule" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PrimaryObjectRule respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExternalObjectResolutionMappingResponse) RawJSON() string { return r.JSON.raw }
func (r *ExternalObjectResolutionMappingResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExternalPrimaryObjectResolutionRule struct {
	EventPropertyName        string `json:"eventPropertyName" api:"required"`
	TargetObjectPropertyName string `json:"targetObjectPropertyName" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		EventPropertyName        respjson.Field
		TargetObjectPropertyName respjson.Field
		ExtraFields              map[string]respjson.Field
		raw                      string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExternalPrimaryObjectResolutionRule) RawJSON() string { return r.JSON.raw }
func (r *ExternalPrimaryObjectResolutionRule) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this ExternalPrimaryObjectResolutionRule to a
// ExternalPrimaryObjectResolutionRuleParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// ExternalPrimaryObjectResolutionRuleParam.Overrides()
func (r ExternalPrimaryObjectResolutionRule) ToParam() ExternalPrimaryObjectResolutionRuleParam {
	return param.Override[ExternalPrimaryObjectResolutionRuleParam](json.RawMessage(r.RawJSON()))
}

// The properties EventPropertyName, TargetObjectPropertyName are required.
type ExternalPrimaryObjectResolutionRuleParam struct {
	EventPropertyName        string `json:"eventPropertyName" api:"required"`
	TargetObjectPropertyName string `json:"targetObjectPropertyName" api:"required"`
	paramObj
}

func (r ExternalPrimaryObjectResolutionRuleParam) MarshalJSON() (data []byte, err error) {
	type shadow ExternalPrimaryObjectResolutionRuleParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExternalPrimaryObjectResolutionRuleParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FiscalQuarter struct {
	Day   int64 `json:"day" api:"required"`
	Month int64 `json:"month" api:"required"`
	// Any of "FISCAL_QUARTER".
	ReferenceType FiscalQuarterReferenceType `json:"referenceType" api:"required"`
	Hour          int64                      `json:"hour"`
	Millisecond   int64                      `json:"millisecond"`
	Minute        int64                      `json:"minute"`
	Second        int64                      `json:"second"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Day           respjson.Field
		Month         respjson.Field
		ReferenceType respjson.Field
		Hour          respjson.Field
		Millisecond   respjson.Field
		Minute        respjson.Field
		Second        respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FiscalQuarter) RawJSON() string { return r.JSON.raw }
func (r *FiscalQuarter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FiscalQuarterReferenceType string

const (
	FiscalQuarterReferenceTypeFiscalQuarter FiscalQuarterReferenceType = "FISCAL_QUARTER"
)

type FiscalYear struct {
	Day   int64 `json:"day" api:"required"`
	Month int64 `json:"month" api:"required"`
	// Any of "FISCAL_YEAR".
	ReferenceType FiscalYearReferenceType `json:"referenceType" api:"required"`
	Hour          int64                   `json:"hour"`
	Millisecond   int64                   `json:"millisecond"`
	Minute        int64                   `json:"minute"`
	Second        int64                   `json:"second"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Day           respjson.Field
		Month         respjson.Field
		ReferenceType respjson.Field
		Hour          respjson.Field
		Millisecond   respjson.Field
		Minute        respjson.Field
		Second        respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FiscalYear) RawJSON() string { return r.JSON.raw }
func (r *FiscalYear) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FiscalYearReferenceType string

const (
	FiscalYearReferenceTypeFiscalYear FiscalYearReferenceType = "FISCAL_YEAR"
)

type IndexOffset struct {
	Days         int64 `json:"days"`
	Hours        int64 `json:"hours"`
	Milliseconds int64 `json:"milliseconds"`
	Minutes      int64 `json:"minutes"`
	Months       int64 `json:"months"`
	Quarters     int64 `json:"quarters"`
	Seconds      int64 `json:"seconds"`
	Weeks        int64 `json:"weeks"`
	Years        int64 `json:"years"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Days         respjson.Field
		Hours        respjson.Field
		Milliseconds respjson.Field
		Minutes      respjson.Field
		Months       respjson.Field
		Quarters     respjson.Field
		Seconds      respjson.Field
		Weeks        respjson.Field
		Years        respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IndexOffset) RawJSON() string { return r.JSON.raw }
func (r *IndexOffset) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IndexedTimePoint struct {
	IndexReference IndexedTimePointIndexReferenceUnion `json:"indexReference" api:"required"`
	// Any of "INDEXED".
	TimeType IndexedTimePointTimeType `json:"timeType" api:"required"`
	// Any of "CUSTOM", "PORTAL", "USER".
	TimezoneSource            IndexedTimePointTimezoneSource `json:"timezoneSource" api:"required"`
	ZoneID                    string                         `json:"zoneId" api:"required"`
	Offset                    IndexOffset                    `json:"offset"`
	ShouldGenerateRefreshTime bool                           `json:"shouldGenerateRefreshTime"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		IndexReference            respjson.Field
		TimeType                  respjson.Field
		TimezoneSource            respjson.Field
		ZoneID                    respjson.Field
		Offset                    respjson.Field
		ShouldGenerateRefreshTime respjson.Field
		ExtraFields               map[string]respjson.Field
		raw                       string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IndexedTimePoint) RawJSON() string { return r.JSON.raw }
func (r *IndexedTimePoint) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// IndexedTimePointIndexReferenceUnion contains all possible properties and values
// from [NowReference], [TodayReference], [WeekReference], [MonthReference],
// [QuarterReference], [FiscalQuarter], [YearReference], [FiscalYear].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type IndexedTimePointIndexReferenceUnion struct {
	ReferenceType string `json:"referenceType"`
	Hour          int64  `json:"hour"`
	Millisecond   int64  `json:"millisecond"`
	Minute        int64  `json:"minute"`
	Second        int64  `json:"second"`
	// This field is from variant [WeekReference].
	DayOfWeek WeekReferenceDayOfWeek `json:"dayOfWeek"`
	Day       int64                  `json:"day"`
	Month     int64                  `json:"month"`
	JSON      struct {
		ReferenceType respjson.Field
		Hour          respjson.Field
		Millisecond   respjson.Field
		Minute        respjson.Field
		Second        respjson.Field
		DayOfWeek     respjson.Field
		Day           respjson.Field
		Month         respjson.Field
		raw           string
	} `json:"-"`
}

func (u IndexedTimePointIndexReferenceUnion) AsNow() (v NowReference) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u IndexedTimePointIndexReferenceUnion) AsToday() (v TodayReference) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u IndexedTimePointIndexReferenceUnion) AsWeek() (v WeekReference) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u IndexedTimePointIndexReferenceUnion) AsMonth() (v MonthReference) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u IndexedTimePointIndexReferenceUnion) AsQuarter() (v QuarterReference) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u IndexedTimePointIndexReferenceUnion) AsFiscalQuarter() (v FiscalQuarter) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u IndexedTimePointIndexReferenceUnion) AsYear() (v YearReference) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u IndexedTimePointIndexReferenceUnion) AsFiscalYear() (v FiscalYear) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u IndexedTimePointIndexReferenceUnion) RawJSON() string { return u.JSON.raw }

func (r *IndexedTimePointIndexReferenceUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IndexedTimePointTimeType string

const (
	IndexedTimePointTimeTypeIndexed IndexedTimePointTimeType = "INDEXED"
)

type IndexedTimePointTimezoneSource string

const (
	IndexedTimePointTimezoneSourceCustom IndexedTimePointTimezoneSource = "CUSTOM"
	IndexedTimePointTimezoneSourcePortal IndexedTimePointTimezoneSource = "PORTAL"
	IndexedTimePointTimezoneSourceUser   IndexedTimePointTimezoneSource = "USER"
)

type MonthReference struct {
	Day int64 `json:"day" api:"required"`
	// Any of "MONTH".
	ReferenceType MonthReferenceReferenceType `json:"referenceType" api:"required"`
	Hour          int64                       `json:"hour"`
	Millisecond   int64                       `json:"millisecond"`
	Minute        int64                       `json:"minute"`
	Second        int64                       `json:"second"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Day           respjson.Field
		ReferenceType respjson.Field
		Hour          respjson.Field
		Millisecond   respjson.Field
		Minute        respjson.Field
		Second        respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MonthReference) RawJSON() string { return r.JSON.raw }
func (r *MonthReference) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MonthReferenceReferenceType string

const (
	MonthReferenceReferenceTypeMonth MonthReferenceReferenceType = "MONTH"
)

type MultiStringPropertyOperation struct {
	CoalescingRefineBy           MultiStringPropertyOperationCoalescingRefineByUnion `json:"coalescingRefineBy" api:"required"`
	IncludeObjectsWithNoValueSet bool                                                `json:"includeObjectsWithNoValueSet" api:"required"`
	OperationType                string                                              `json:"operationType" api:"required"`
	// Any of "CONTAINS", "CONTAINS_EXACTLY", "DOES_NOT_CONTAIN",
	// "DOES_NOT_CONTAIN_EXACTLY", "ENDS_WITH", "IS_EQUAL_TO", "IS_NOT_EQUAL_TO",
	// "STARTS_WITH".
	Operator     MultiStringPropertyOperationOperator `json:"operator" api:"required"`
	OperatorName string                               `json:"operatorName" api:"required"`
	// Any of "multistring".
	PropertyType    MultiStringPropertyOperationPropertyType         `json:"propertyType" api:"required"`
	Values          []string                                         `json:"values" api:"required"`
	DefaultValue    string                                           `json:"defaultValue"`
	PruningRefineBy MultiStringPropertyOperationPruningRefineByUnion `json:"pruningRefineBy"`
	RenderSpec      string                                           `json:"renderSpec"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CoalescingRefineBy           respjson.Field
		IncludeObjectsWithNoValueSet respjson.Field
		OperationType                respjson.Field
		Operator                     respjson.Field
		OperatorName                 respjson.Field
		PropertyType                 respjson.Field
		Values                       respjson.Field
		DefaultValue                 respjson.Field
		PruningRefineBy              respjson.Field
		RenderSpec                   respjson.Field
		ExtraFields                  map[string]respjson.Field
		raw                          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MultiStringPropertyOperation) RawJSON() string { return r.JSON.raw }
func (r *MultiStringPropertyOperation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// MultiStringPropertyOperationCoalescingRefineByUnion contains all possible
// properties and values from [NumOccurrencesRefineBy], [SetOccurrencesRefineBy].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type MultiStringPropertyOperationCoalescingRefineByUnion struct {
	Type string `json:"type"`
	// This field is from variant [NumOccurrencesRefineBy].
	MaxOccurrences int64 `json:"maxOccurrences"`
	// This field is from variant [NumOccurrencesRefineBy].
	MinOccurrences int64 `json:"minOccurrences"`
	// This field is from variant [SetOccurrencesRefineBy].
	SetType SetOccurrencesRefineBySetType `json:"setType"`
	JSON    struct {
		Type           respjson.Field
		MaxOccurrences respjson.Field
		MinOccurrences respjson.Field
		SetType        respjson.Field
		raw            string
	} `json:"-"`
}

func (u MultiStringPropertyOperationCoalescingRefineByUnion) AsNumOccurrencesRefineBy() (v NumOccurrencesRefineBy) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MultiStringPropertyOperationCoalescingRefineByUnion) AsSetOccurrencesRefineBy() (v SetOccurrencesRefineBy) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u MultiStringPropertyOperationCoalescingRefineByUnion) RawJSON() string { return u.JSON.raw }

func (r *MultiStringPropertyOperationCoalescingRefineByUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MultiStringPropertyOperationOperator string

const (
	MultiStringPropertyOperationOperatorContains              MultiStringPropertyOperationOperator = "CONTAINS"
	MultiStringPropertyOperationOperatorContainsExactly       MultiStringPropertyOperationOperator = "CONTAINS_EXACTLY"
	MultiStringPropertyOperationOperatorDoesNotContain        MultiStringPropertyOperationOperator = "DOES_NOT_CONTAIN"
	MultiStringPropertyOperationOperatorDoesNotContainExactly MultiStringPropertyOperationOperator = "DOES_NOT_CONTAIN_EXACTLY"
	MultiStringPropertyOperationOperatorEndsWith              MultiStringPropertyOperationOperator = "ENDS_WITH"
	MultiStringPropertyOperationOperatorIsEqualTo             MultiStringPropertyOperationOperator = "IS_EQUAL_TO"
	MultiStringPropertyOperationOperatorIsNotEqualTo          MultiStringPropertyOperationOperator = "IS_NOT_EQUAL_TO"
	MultiStringPropertyOperationOperatorStartsWith            MultiStringPropertyOperationOperator = "STARTS_WITH"
)

type MultiStringPropertyOperationPropertyType string

const (
	MultiStringPropertyOperationPropertyTypeMultistring MultiStringPropertyOperationPropertyType = "multistring"
)

// MultiStringPropertyOperationPruningRefineByUnion contains all possible
// properties and values from [RelativeComparativeTimestampRefineBy],
// [RelativeRangedTimestampRefineBy], [AbsoluteComparativeTimestampRefineBy],
// [AbsoluteRangedTimestampRefineBy], [AllHistoryRefineBy], [TimePointOperation],
// [RangedTimeOperation].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type MultiStringPropertyOperationPruningRefineByUnion struct {
	Comparison string `json:"comparison"`
	// This field is from variant [RelativeComparativeTimestampRefineBy].
	TimeOffset TimeOffset `json:"timeOffset"`
	Type       string     `json:"type"`
	// This field is from variant [RelativeRangedTimestampRefineBy].
	LowerBoundOffset TimeOffset `json:"lowerBoundOffset"`
	RangeType        string     `json:"rangeType"`
	// This field is from variant [RelativeRangedTimestampRefineBy].
	UpperBoundOffset TimeOffset `json:"upperBoundOffset"`
	// This field is from variant [AbsoluteComparativeTimestampRefineBy].
	Timestamp int64 `json:"timestamp"`
	// This field is from variant [AbsoluteRangedTimestampRefineBy].
	LowerTimestamp int64 `json:"lowerTimestamp"`
	// This field is from variant [AbsoluteRangedTimestampRefineBy].
	UpperTimestamp int64 `json:"upperTimestamp"`
	// This field is from variant [TimePointOperation].
	EndpointBehavior             TimePointOperationEndpointBehavior `json:"endpointBehavior"`
	IncludeObjectsWithNoValueSet bool                               `json:"includeObjectsWithNoValueSet"`
	OperationType                string                             `json:"operationType"`
	Operator                     string                             `json:"operator"`
	OperatorName                 string                             `json:"operatorName"`
	PropertyParser               string                             `json:"propertyParser"`
	PropertyType                 string                             `json:"propertyType"`
	// This field is from variant [TimePointOperation].
	TimePoint    TimePointOperationTimePointUnion `json:"timePoint"`
	DefaultValue string                           `json:"defaultValue"`
	RenderSpec   string                           `json:"renderSpec"`
	// This field is from variant [RangedTimeOperation].
	LowerBoundEndpointBehavior RangedTimeOperationLowerBoundEndpointBehavior `json:"lowerBoundEndpointBehavior"`
	// This field is from variant [RangedTimeOperation].
	LowerBoundTimePoint RangedTimeOperationLowerBoundTimePointUnion `json:"lowerBoundTimePoint"`
	// This field is from variant [RangedTimeOperation].
	UpperBoundEndpointBehavior RangedTimeOperationUpperBoundEndpointBehavior `json:"upperBoundEndpointBehavior"`
	// This field is from variant [RangedTimeOperation].
	UpperBoundTimePoint RangedTimeOperationUpperBoundTimePointUnion `json:"upperBoundTimePoint"`
	JSON                struct {
		Comparison                   respjson.Field
		TimeOffset                   respjson.Field
		Type                         respjson.Field
		LowerBoundOffset             respjson.Field
		RangeType                    respjson.Field
		UpperBoundOffset             respjson.Field
		Timestamp                    respjson.Field
		LowerTimestamp               respjson.Field
		UpperTimestamp               respjson.Field
		EndpointBehavior             respjson.Field
		IncludeObjectsWithNoValueSet respjson.Field
		OperationType                respjson.Field
		Operator                     respjson.Field
		OperatorName                 respjson.Field
		PropertyParser               respjson.Field
		PropertyType                 respjson.Field
		TimePoint                    respjson.Field
		DefaultValue                 respjson.Field
		RenderSpec                   respjson.Field
		LowerBoundEndpointBehavior   respjson.Field
		LowerBoundTimePoint          respjson.Field
		UpperBoundEndpointBehavior   respjson.Field
		UpperBoundTimePoint          respjson.Field
		raw                          string
	} `json:"-"`
}

func (u MultiStringPropertyOperationPruningRefineByUnion) AsRelativeComparativeTimestampRefineBy() (v RelativeComparativeTimestampRefineBy) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MultiStringPropertyOperationPruningRefineByUnion) AsRelativeRangedTimestampRefineBy() (v RelativeRangedTimestampRefineBy) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MultiStringPropertyOperationPruningRefineByUnion) AsAbsoluteComparativeTimestampRefineBy() (v AbsoluteComparativeTimestampRefineBy) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MultiStringPropertyOperationPruningRefineByUnion) AsAbsoluteRangedTimestampRefineBy() (v AbsoluteRangedTimestampRefineBy) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MultiStringPropertyOperationPruningRefineByUnion) AsAllHistoryRefineBy() (v AllHistoryRefineBy) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MultiStringPropertyOperationPruningRefineByUnion) AsTimepoint() (v TimePointOperation) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MultiStringPropertyOperationPruningRefineByUnion) AsRangedtime() (v RangedTimeOperation) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u MultiStringPropertyOperationPruningRefineByUnion) RawJSON() string { return u.JSON.raw }

func (r *MultiStringPropertyOperationPruningRefineByUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NowReference struct {
	// Any of "NOW".
	ReferenceType NowReferenceReferenceType `json:"referenceType" api:"required"`
	Hour          int64                     `json:"hour"`
	Millisecond   int64                     `json:"millisecond"`
	Minute        int64                     `json:"minute"`
	Second        int64                     `json:"second"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ReferenceType respjson.Field
		Hour          respjson.Field
		Millisecond   respjson.Field
		Minute        respjson.Field
		Second        respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NowReference) RawJSON() string { return r.JSON.raw }
func (r *NowReference) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NowReferenceReferenceType string

const (
	NowReferenceReferenceTypeNow NowReferenceReferenceType = "NOW"
)

type NumOccurrencesRefineBy struct {
	// Any of "NumOccurrencesRefineBy".
	Type           NumOccurrencesRefineByType `json:"type" api:"required"`
	MaxOccurrences int64                      `json:"maxOccurrences"`
	MinOccurrences int64                      `json:"minOccurrences"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type           respjson.Field
		MaxOccurrences respjson.Field
		MinOccurrences respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NumOccurrencesRefineBy) RawJSON() string { return r.JSON.raw }
func (r *NumOccurrencesRefineBy) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NumOccurrencesRefineByType string

const (
	NumOccurrencesRefineByTypeNumOccurrencesRefineBy NumOccurrencesRefineByType = "NumOccurrencesRefineBy"
)

type NumberPropertyOperation struct {
	IncludeObjectsWithNoValueSet bool   `json:"includeObjectsWithNoValueSet" api:"required"`
	OperationType                string `json:"operationType" api:"required"`
	// Any of "HAS_EVER_BEEN_EQUAL_TO", "HAS_NEVER_BEEN_EQUAL_TO", "IS_EQUAL_TO",
	// "IS_GREATER_THAN", "IS_GREATER_THAN_OR_EQUAL_TO", "IS_LESS_THAN",
	// "IS_LESS_THAN_OR_EQUAL_TO", "IS_NOT_EQUAL_TO".
	Operator     NumberPropertyOperationOperator `json:"operator" api:"required"`
	OperatorName string                          `json:"operatorName" api:"required"`
	// Any of "number".
	PropertyType NumberPropertyOperationPropertyType `json:"propertyType" api:"required"`
	Value        float64                             `json:"value" api:"required"`
	DefaultValue string                              `json:"defaultValue"`
	RenderSpec   string                              `json:"renderSpec"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		IncludeObjectsWithNoValueSet respjson.Field
		OperationType                respjson.Field
		Operator                     respjson.Field
		OperatorName                 respjson.Field
		PropertyType                 respjson.Field
		Value                        respjson.Field
		DefaultValue                 respjson.Field
		RenderSpec                   respjson.Field
		ExtraFields                  map[string]respjson.Field
		raw                          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NumberPropertyOperation) RawJSON() string { return r.JSON.raw }
func (r *NumberPropertyOperation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NumberPropertyOperationOperator string

const (
	NumberPropertyOperationOperatorHasEverBeenEqualTo     NumberPropertyOperationOperator = "HAS_EVER_BEEN_EQUAL_TO"
	NumberPropertyOperationOperatorHasNeverBeenEqualTo    NumberPropertyOperationOperator = "HAS_NEVER_BEEN_EQUAL_TO"
	NumberPropertyOperationOperatorIsEqualTo              NumberPropertyOperationOperator = "IS_EQUAL_TO"
	NumberPropertyOperationOperatorIsGreaterThan          NumberPropertyOperationOperator = "IS_GREATER_THAN"
	NumberPropertyOperationOperatorIsGreaterThanOrEqualTo NumberPropertyOperationOperator = "IS_GREATER_THAN_OR_EQUAL_TO"
	NumberPropertyOperationOperatorIsLessThan             NumberPropertyOperationOperator = "IS_LESS_THAN"
	NumberPropertyOperationOperatorIsLessThanOrEqualTo    NumberPropertyOperationOperator = "IS_LESS_THAN_OR_EQUAL_TO"
	NumberPropertyOperationOperatorIsNotEqualTo           NumberPropertyOperationOperator = "IS_NOT_EQUAL_TO"
)

type NumberPropertyOperationPropertyType string

const (
	NumberPropertyOperationPropertyTypeNumber NumberPropertyOperationPropertyType = "number"
)

// A HubSpot property option
type Option struct {
	// Whether the option is displayed in HubSpot's UI.
	Hidden bool `json:"hidden" api:"required"`
	// A user-friendly label that identifies the option.
	Label string `json:"label" api:"required"`
	// The actual value of the option.
	Value string `json:"value" api:"required"`
	// A description of the option.
	Description string `json:"description"`
	// The position of the item relative to others in the list.
	DisplayOrder int64 `json:"displayOrder"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Hidden       respjson.Field
		Label        respjson.Field
		Value        respjson.Field
		Description  respjson.Field
		DisplayOrder respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Option) RawJSON() string { return r.JSON.raw }
func (r *Option) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this Option to a OptionParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// OptionParam.Overrides()
func (r Option) ToParam() OptionParam {
	return param.Override[OptionParam](json.RawMessage(r.RawJSON()))
}

// A HubSpot property option
//
// The properties Hidden, Label, Value are required.
type OptionParam struct {
	// Whether the option is displayed in HubSpot's UI.
	Hidden bool `json:"hidden" api:"required"`
	// A user-friendly label that identifies the option.
	Label string `json:"label" api:"required"`
	// The actual value of the option.
	Value string `json:"value" api:"required"`
	// A description of the option.
	Description param.Opt[string] `json:"description,omitzero"`
	// The position of the item relative to others in the list.
	DisplayOrder param.Opt[int64] `json:"displayOrder,omitzero"`
	paramObj
}

func (r OptionParam) MarshalJSON() (data []byte, err error) {
	type shadow OptionParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OptionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties DisplayOrder, Hidden, Label, Value are required.
type OptionInputParam struct {
	DisplayOrder int64 `json:"displayOrder" api:"required"`
	Hidden       bool  `json:"hidden" api:"required"`
	// null
	Label string `json:"label" api:"required"`
	// null
	Value string `json:"value" api:"required"`
	// null
	Description param.Opt[string] `json:"description,omitzero"`
	paramObj
}

func (r OptionInputParam) MarshalJSON() (data []byte, err error) {
	type shadow OptionInputParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OptionInputParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A HubSpot property
type Property struct {
	// A summary of the property's purpose.
	Description string `json:"description" api:"required"`
	// Determines how the property will appear in HubSpot's UI or on a form. Learn more
	// in the properties API guide.
	FieldType string `json:"fieldType" api:"required"`
	// The name of the group to which the property is assigned.
	GroupName string `json:"groupName" api:"required"`
	// The display label for the property.
	Label string `json:"label" api:"required"`
	// The internal name for the property.
	Name string `json:"name" api:"required"`
	// A list of valid options for the property. This field is required for enumerated
	// properties.
	Options []Option `json:"options" api:"required"`
	// The data type of the property, such as string or number.
	Type string `json:"type" api:"required"`
	// Whether the property is archived.
	Archived bool `json:"archived"`
	// The timestamp when the property was archived, in ISO 8601 format.
	ArchivedAt time.Time `json:"archivedAt" format:"date-time"`
	// Whether the property is a calculated field.
	Calculated bool `json:"calculated"`
	// The formula used for calculated properties.
	CalculationFormula string `json:"calculationFormula"`
	// The timestamp when the property was created, in ISO 8601 format.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// The ID of the user who created the property.
	CreatedUserID string `json:"createdUserId"`
	// Indicates the sensitivity level of the property, such as "non_sensitive",
	// "sensitive", or "highly_sensitive".
	//
	// Any of "highly_sensitive", "non_sensitive", "sensitive".
	DataSensitivity PropertyDataSensitivity `json:"dataSensitivity"`
	// Any of "absolute", "absolute_with_relative", "time_since", "time_until".
	DateDisplayHint PropertyDateDisplayHint `json:"dateDisplayHint"`
	// The position of the item relative to others in the list.
	DisplayOrder int64 `json:"displayOrder"`
	// Applicable only for enumeration type properties. Should be set to true with a
	// 'referencedObjectType' of 'OWNER'. Otherwise false.
	ExternalOptions bool `json:"externalOptions"`
	// Whether the property can appear on forms.
	FormField bool `json:"formField"`
	// Whether the property is a unique identifier property.
	HasUniqueValue bool `json:"hasUniqueValue"`
	// Whether or not the property will be hidden from the HubSpot UI. It's recommended
	// that this be set to false for custom properties.
	Hidden bool `json:"hidden"`
	// A boolean value set to true for HubSpot default properties.
	HubspotDefined       bool                         `json:"hubspotDefined"`
	ModificationMetadata PropertyModificationMetadata `json:"modificationMetadata"`
	// Deprecated. Use externalOptionsReferenceType instead.
	ReferencedObjectType string `json:"referencedObjectType"`
	// When sensitiveData is true, lists the type of sensitive data contained in the
	// property (e.g., "HIPAA").
	SensitiveDataCategories []string `json:"sensitiveDataCategories"`
	// Whether to show the currency symbol in HubSpot's UI.
	ShowCurrencySymbol bool `json:"showCurrencySymbol"`
	// The timestamp when the property was last updated, in ISO 8601 format.
	UpdatedAt     time.Time `json:"updatedAt" format:"date-time"`
	UpdatedUserID string    `json:"updatedUserId"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Description             respjson.Field
		FieldType               respjson.Field
		GroupName               respjson.Field
		Label                   respjson.Field
		Name                    respjson.Field
		Options                 respjson.Field
		Type                    respjson.Field
		Archived                respjson.Field
		ArchivedAt              respjson.Field
		Calculated              respjson.Field
		CalculationFormula      respjson.Field
		CreatedAt               respjson.Field
		CreatedUserID           respjson.Field
		DataSensitivity         respjson.Field
		DateDisplayHint         respjson.Field
		DisplayOrder            respjson.Field
		ExternalOptions         respjson.Field
		FormField               respjson.Field
		HasUniqueValue          respjson.Field
		Hidden                  respjson.Field
		HubspotDefined          respjson.Field
		ModificationMetadata    respjson.Field
		ReferencedObjectType    respjson.Field
		SensitiveDataCategories respjson.Field
		ShowCurrencySymbol      respjson.Field
		UpdatedAt               respjson.Field
		UpdatedUserID           respjson.Field
		ExtraFields             map[string]respjson.Field
		raw                     string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Property) RawJSON() string { return r.JSON.raw }
func (r *Property) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Indicates the sensitivity level of the property, such as "non_sensitive",
// "sensitive", or "highly_sensitive".
type PropertyDataSensitivity string

const (
	PropertyDataSensitivityHighlySensitive PropertyDataSensitivity = "highly_sensitive"
	PropertyDataSensitivityNonSensitive    PropertyDataSensitivity = "non_sensitive"
	PropertyDataSensitivitySensitive       PropertyDataSensitivity = "sensitive"
)

type PropertyDateDisplayHint string

const (
	PropertyDateDisplayHintAbsolute             PropertyDateDisplayHint = "absolute"
	PropertyDateDisplayHintAbsoluteWithRelative PropertyDateDisplayHint = "absolute_with_relative"
	PropertyDateDisplayHintTimeSince            PropertyDateDisplayHint = "time_since"
	PropertyDateDisplayHintTimeUntil            PropertyDateDisplayHint = "time_until"
)

type PropertyFilter struct {
	// Any of "PROPERTY".
	FilterType        PropertyFilterFilterType     `json:"filterType" api:"required"`
	Operation         PropertyFilterOperationUnion `json:"operation" api:"required"`
	Property          string                       `json:"property" api:"required"`
	Context           PropertyFilterContext        `json:"context"`
	FilterInsightsID  int64                        `json:"filterInsightsId"`
	FrameworkFilterID int64                        `json:"frameworkFilterId"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FilterType        respjson.Field
		Operation         respjson.Field
		Property          respjson.Field
		Context           respjson.Field
		FilterInsightsID  respjson.Field
		FrameworkFilterID respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PropertyFilter) RawJSON() string { return r.JSON.raw }
func (r *PropertyFilter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PropertyFilterFilterType string

const (
	PropertyFilterFilterTypeProperty PropertyFilterFilterType = "PROPERTY"
)

// PropertyFilterOperationUnion contains all possible properties and values from
// [BoolPropertyOperation], [NumberPropertyOperation], [StringPropertyOperation],
// [DateTimePropertyOperation], [RangedDatePropertyOperation],
// [ComparativeDatePropertyOperation], [ComparativeBoolPropertyOperation],
// [ComparativeNumberPropertyOperation], [ComparativeStringPropertyOperation],
// [ComparativePropertyUpdatedOperation], [RollingDateRangePropertyOperation],
// [RollingPropertyUpdatedOperation], [EnumerationPropertyOperation],
// [AllPropertyTypesOperation], [RangedNumberPropertyOperation],
// [MultiStringPropertyOperation], [DatePropertyOperation],
// [CalendarDatePropertyOperation], [TimePointOperation], [RangedTimeOperation],
// [RegexPropertyOperation].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type PropertyFilterOperationUnion struct {
	IncludeObjectsWithNoValueSet bool   `json:"includeObjectsWithNoValueSet"`
	OperationType                string `json:"operationType"`
	Operator                     string `json:"operator"`
	OperatorName                 string `json:"operatorName"`
	PropertyType                 string `json:"propertyType"`
	// This field is a union of [bool], [float64], [string]
	Value                      PropertyFilterOperationUnionValue `json:"value"`
	DefaultValue               string                            `json:"defaultValue"`
	RenderSpec                 string                            `json:"renderSpec"`
	RequiresTimeZoneConversion bool                              `json:"requiresTimeZoneConversion"`
	// This field is from variant [DateTimePropertyOperation].
	Timestamp int64 `json:"timestamp"`
	// This field is from variant [RangedDatePropertyOperation].
	LowerBoundTimestamp int64 `json:"lowerBoundTimestamp"`
	// This field is from variant [RangedDatePropertyOperation].
	UpperBoundTimestamp    int64    `json:"upperBoundTimestamp"`
	ComparisonPropertyName string   `json:"comparisonPropertyName"`
	DefaultComparisonValue string   `json:"defaultComparisonValue"`
	NumberOfDays           int64    `json:"numberOfDays"`
	Values                 []string `json:"values"`
	// This field is a union of [AllPropertyTypesOperationCoalescingRefineByUnion],
	// [MultiStringPropertyOperationCoalescingRefineByUnion]
	CoalescingRefineBy PropertyFilterOperationUnionCoalescingRefineBy `json:"coalescingRefineBy"`
	// This field is a union of [AllPropertyTypesOperationPruningRefineByUnion],
	// [MultiStringPropertyOperationPruningRefineByUnion]
	PruningRefineBy PropertyFilterOperationUnionPruningRefineBy `json:"pruningRefineBy"`
	// This field is from variant [RangedNumberPropertyOperation].
	LowerBound float64 `json:"lowerBound"`
	// This field is from variant [RangedNumberPropertyOperation].
	UpperBound float64 `json:"upperBound"`
	// This field is from variant [DatePropertyOperation].
	Day int64 `json:"day"`
	// This field is from variant [DatePropertyOperation].
	Month DatePropertyOperationMonth `json:"month"`
	// This field is from variant [DatePropertyOperation].
	Year int64 `json:"year"`
	// This field is from variant [CalendarDatePropertyOperation].
	TimeUnit CalendarDatePropertyOperationTimeUnit `json:"timeUnit"`
	// This field is from variant [CalendarDatePropertyOperation].
	TimeUnitCount int64 `json:"timeUnitCount"`
	// This field is from variant [CalendarDatePropertyOperation].
	UseFiscalYear bool `json:"useFiscalYear"`
	// This field is from variant [CalendarDatePropertyOperation].
	FiscalYearStart CalendarDatePropertyOperationFiscalYearStart `json:"fiscalYearStart"`
	// This field is from variant [TimePointOperation].
	EndpointBehavior TimePointOperationEndpointBehavior `json:"endpointBehavior"`
	PropertyParser   string                             `json:"propertyParser"`
	// This field is from variant [TimePointOperation].
	TimePoint TimePointOperationTimePointUnion `json:"timePoint"`
	Type      string                           `json:"type"`
	// This field is from variant [RangedTimeOperation].
	LowerBoundEndpointBehavior RangedTimeOperationLowerBoundEndpointBehavior `json:"lowerBoundEndpointBehavior"`
	// This field is from variant [RangedTimeOperation].
	LowerBoundTimePoint RangedTimeOperationLowerBoundTimePointUnion `json:"lowerBoundTimePoint"`
	// This field is from variant [RangedTimeOperation].
	UpperBoundEndpointBehavior RangedTimeOperationUpperBoundEndpointBehavior `json:"upperBoundEndpointBehavior"`
	// This field is from variant [RangedTimeOperation].
	UpperBoundTimePoint RangedTimeOperationUpperBoundTimePointUnion `json:"upperBoundTimePoint"`
	// This field is from variant [RegexPropertyOperation].
	CaseSensitive bool `json:"caseSensitive"`
	// This field is from variant [RegexPropertyOperation].
	Pattern string `json:"pattern"`
	JSON    struct {
		IncludeObjectsWithNoValueSet respjson.Field
		OperationType                respjson.Field
		Operator                     respjson.Field
		OperatorName                 respjson.Field
		PropertyType                 respjson.Field
		Value                        respjson.Field
		DefaultValue                 respjson.Field
		RenderSpec                   respjson.Field
		RequiresTimeZoneConversion   respjson.Field
		Timestamp                    respjson.Field
		LowerBoundTimestamp          respjson.Field
		UpperBoundTimestamp          respjson.Field
		ComparisonPropertyName       respjson.Field
		DefaultComparisonValue       respjson.Field
		NumberOfDays                 respjson.Field
		Values                       respjson.Field
		CoalescingRefineBy           respjson.Field
		PruningRefineBy              respjson.Field
		LowerBound                   respjson.Field
		UpperBound                   respjson.Field
		Day                          respjson.Field
		Month                        respjson.Field
		Year                         respjson.Field
		TimeUnit                     respjson.Field
		TimeUnitCount                respjson.Field
		UseFiscalYear                respjson.Field
		FiscalYearStart              respjson.Field
		EndpointBehavior             respjson.Field
		PropertyParser               respjson.Field
		TimePoint                    respjson.Field
		Type                         respjson.Field
		LowerBoundEndpointBehavior   respjson.Field
		LowerBoundTimePoint          respjson.Field
		UpperBoundEndpointBehavior   respjson.Field
		UpperBoundTimePoint          respjson.Field
		CaseSensitive                respjson.Field
		Pattern                      respjson.Field
		raw                          string
	} `json:"-"`
}

func (u PropertyFilterOperationUnion) AsBool() (v BoolPropertyOperation) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PropertyFilterOperationUnion) AsNumber() (v NumberPropertyOperation) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PropertyFilterOperationUnion) AsString() (v StringPropertyOperation) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PropertyFilterOperationUnion) AsDatetime() (v DateTimePropertyOperation) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PropertyFilterOperationUnion) AsDatetimeRanged() (v RangedDatePropertyOperation) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PropertyFilterOperationUnion) AsDatetimeComparative() (v ComparativeDatePropertyOperation) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PropertyFilterOperationUnion) AsBoolComparative() (v ComparativeBoolPropertyOperation) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PropertyFilterOperationUnion) AsNumberComparative() (v ComparativeNumberPropertyOperation) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PropertyFilterOperationUnion) AsStringComparative() (v ComparativeStringPropertyOperation) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PropertyFilterOperationUnion) AsPropertyUpdatedComparative() (v ComparativePropertyUpdatedOperation) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PropertyFilterOperationUnion) AsDatetimeRolling() (v RollingDateRangePropertyOperation) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PropertyFilterOperationUnion) AsRollingPropertyUpdated() (v RollingPropertyUpdatedOperation) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PropertyFilterOperationUnion) AsEnumeration() (v EnumerationPropertyOperation) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PropertyFilterOperationUnion) AsAlltypes() (v AllPropertyTypesOperation) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PropertyFilterOperationUnion) AsNumberRanged() (v RangedNumberPropertyOperation) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PropertyFilterOperationUnion) AsMultistring() (v MultiStringPropertyOperation) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PropertyFilterOperationUnion) AsDate() (v DatePropertyOperation) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PropertyFilterOperationUnion) AsCalendarDate() (v CalendarDatePropertyOperation) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PropertyFilterOperationUnion) AsTimepoint() (v TimePointOperation) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PropertyFilterOperationUnion) AsRangedtime() (v RangedTimeOperation) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PropertyFilterOperationUnion) AsRegex() (v RegexPropertyOperation) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u PropertyFilterOperationUnion) RawJSON() string { return u.JSON.raw }

func (r *PropertyFilterOperationUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PropertyFilterOperationUnionValue is an implicit subunion of
// [PropertyFilterOperationUnion]. PropertyFilterOperationUnionValue provides
// convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [PropertyFilterOperationUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfBool OfFloat OfString]
type PropertyFilterOperationUnionValue struct {
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	JSON     struct {
		OfBool   respjson.Field
		OfFloat  respjson.Field
		OfString respjson.Field
		raw      string
	} `json:"-"`
}

func (r *PropertyFilterOperationUnionValue) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PropertyFilterOperationUnionCoalescingRefineBy is an implicit subunion of
// [PropertyFilterOperationUnion]. PropertyFilterOperationUnionCoalescingRefineBy
// provides convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [PropertyFilterOperationUnion].
type PropertyFilterOperationUnionCoalescingRefineBy struct {
	Type string `json:"type"`
	// This field is from variant [AllPropertyTypesOperationCoalescingRefineByUnion],
	// [MultiStringPropertyOperationCoalescingRefineByUnion].
	MaxOccurrences int64 `json:"maxOccurrences"`
	// This field is from variant [AllPropertyTypesOperationCoalescingRefineByUnion],
	// [MultiStringPropertyOperationCoalescingRefineByUnion].
	MinOccurrences int64 `json:"minOccurrences"`
	// This field is from variant [AllPropertyTypesOperationCoalescingRefineByUnion],
	// [MultiStringPropertyOperationCoalescingRefineByUnion].
	SetType SetOccurrencesRefineBySetType `json:"setType"`
	JSON    struct {
		Type           respjson.Field
		MaxOccurrences respjson.Field
		MinOccurrences respjson.Field
		SetType        respjson.Field
		raw            string
	} `json:"-"`
}

func (r *PropertyFilterOperationUnionCoalescingRefineBy) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PropertyFilterOperationUnionPruningRefineBy is an implicit subunion of
// [PropertyFilterOperationUnion]. PropertyFilterOperationUnionPruningRefineBy
// provides convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [PropertyFilterOperationUnion].
type PropertyFilterOperationUnionPruningRefineBy struct {
	Comparison string `json:"comparison"`
	// This field is from variant [AllPropertyTypesOperationPruningRefineByUnion],
	// [MultiStringPropertyOperationPruningRefineByUnion].
	TimeOffset TimeOffset `json:"timeOffset"`
	Type       string     `json:"type"`
	// This field is from variant [AllPropertyTypesOperationPruningRefineByUnion],
	// [MultiStringPropertyOperationPruningRefineByUnion].
	LowerBoundOffset TimeOffset `json:"lowerBoundOffset"`
	RangeType        string     `json:"rangeType"`
	// This field is from variant [AllPropertyTypesOperationPruningRefineByUnion],
	// [MultiStringPropertyOperationPruningRefineByUnion].
	UpperBoundOffset TimeOffset `json:"upperBoundOffset"`
	// This field is from variant [AllPropertyTypesOperationPruningRefineByUnion],
	// [MultiStringPropertyOperationPruningRefineByUnion].
	Timestamp int64 `json:"timestamp"`
	// This field is from variant [AllPropertyTypesOperationPruningRefineByUnion],
	// [MultiStringPropertyOperationPruningRefineByUnion].
	LowerTimestamp int64 `json:"lowerTimestamp"`
	// This field is from variant [AllPropertyTypesOperationPruningRefineByUnion],
	// [MultiStringPropertyOperationPruningRefineByUnion].
	UpperTimestamp int64 `json:"upperTimestamp"`
	// This field is from variant [AllPropertyTypesOperationPruningRefineByUnion],
	// [MultiStringPropertyOperationPruningRefineByUnion].
	EndpointBehavior             TimePointOperationEndpointBehavior `json:"endpointBehavior"`
	IncludeObjectsWithNoValueSet bool                               `json:"includeObjectsWithNoValueSet"`
	OperationType                string                             `json:"operationType"`
	Operator                     string                             `json:"operator"`
	OperatorName                 string                             `json:"operatorName"`
	PropertyParser               string                             `json:"propertyParser"`
	PropertyType                 string                             `json:"propertyType"`
	// This field is from variant [AllPropertyTypesOperationPruningRefineByUnion],
	// [MultiStringPropertyOperationPruningRefineByUnion].
	TimePoint    TimePointOperationTimePointUnion `json:"timePoint"`
	DefaultValue string                           `json:"defaultValue"`
	RenderSpec   string                           `json:"renderSpec"`
	// This field is from variant [AllPropertyTypesOperationPruningRefineByUnion],
	// [MultiStringPropertyOperationPruningRefineByUnion].
	LowerBoundEndpointBehavior RangedTimeOperationLowerBoundEndpointBehavior `json:"lowerBoundEndpointBehavior"`
	// This field is from variant [AllPropertyTypesOperationPruningRefineByUnion],
	// [MultiStringPropertyOperationPruningRefineByUnion].
	LowerBoundTimePoint RangedTimeOperationLowerBoundTimePointUnion `json:"lowerBoundTimePoint"`
	// This field is from variant [AllPropertyTypesOperationPruningRefineByUnion],
	// [MultiStringPropertyOperationPruningRefineByUnion].
	UpperBoundEndpointBehavior RangedTimeOperationUpperBoundEndpointBehavior `json:"upperBoundEndpointBehavior"`
	// This field is from variant [AllPropertyTypesOperationPruningRefineByUnion],
	// [MultiStringPropertyOperationPruningRefineByUnion].
	UpperBoundTimePoint RangedTimeOperationUpperBoundTimePointUnion `json:"upperBoundTimePoint"`
	JSON                struct {
		Comparison                   respjson.Field
		TimeOffset                   respjson.Field
		Type                         respjson.Field
		LowerBoundOffset             respjson.Field
		RangeType                    respjson.Field
		UpperBoundOffset             respjson.Field
		Timestamp                    respjson.Field
		LowerTimestamp               respjson.Field
		UpperTimestamp               respjson.Field
		EndpointBehavior             respjson.Field
		IncludeObjectsWithNoValueSet respjson.Field
		OperationType                respjson.Field
		Operator                     respjson.Field
		OperatorName                 respjson.Field
		PropertyParser               respjson.Field
		PropertyType                 respjson.Field
		TimePoint                    respjson.Field
		DefaultValue                 respjson.Field
		RenderSpec                   respjson.Field
		LowerBoundEndpointBehavior   respjson.Field
		LowerBoundTimePoint          respjson.Field
		UpperBoundEndpointBehavior   respjson.Field
		UpperBoundTimePoint          respjson.Field
		raw                          string
	} `json:"-"`
}

func (r *PropertyFilterOperationUnionPruningRefineBy) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PropertyFilterContext struct {
	ObjectTypeID string `json:"objectTypeId" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ObjectTypeID respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PropertyFilterContext) RawJSON() string { return r.JSON.raw }
func (r *PropertyFilterContext) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PropertyModificationMetadata struct {
	Archivable         bool `json:"archivable" api:"required"`
	ReadOnlyDefinition bool `json:"readOnlyDefinition" api:"required"`
	ReadOnlyValue      bool `json:"readOnlyValue" api:"required"`
	ReadOnlyOptions    bool `json:"readOnlyOptions"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Archivable         respjson.Field
		ReadOnlyDefinition respjson.Field
		ReadOnlyValue      respjson.Field
		ReadOnlyOptions    respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PropertyModificationMetadata) RawJSON() string { return r.JSON.raw }
func (r *PropertyModificationMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PropertyReferencedTime struct {
	Property string `json:"property" api:"required"`
	// Any of "ANNIVERSARY", "ANNIVERSARY_WITH_ZONE_SAME_LOCAL_CONVERSION",
	// "UPDATED_AT", "VALUE", "VALUE_WITH_ZONE_SAME_LOCAL_CONVERSION".
	ReferenceType PropertyReferencedTimeReferenceType `json:"referenceType" api:"required"`
	// Any of "PROPERTY_REFERENCED".
	TimeType PropertyReferencedTimeTimeType `json:"timeType" api:"required"`
	// Any of "CUSTOM", "PORTAL", "USER".
	TimezoneSource PropertyReferencedTimeTimezoneSource `json:"timezoneSource" api:"required"`
	ZoneID         string                               `json:"zoneId" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Property       respjson.Field
		ReferenceType  respjson.Field
		TimeType       respjson.Field
		TimezoneSource respjson.Field
		ZoneID         respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PropertyReferencedTime) RawJSON() string { return r.JSON.raw }
func (r *PropertyReferencedTime) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PropertyReferencedTimeReferenceType string

const (
	PropertyReferencedTimeReferenceTypeAnniversary                            PropertyReferencedTimeReferenceType = "ANNIVERSARY"
	PropertyReferencedTimeReferenceTypeAnniversaryWithZoneSameLocalConversion PropertyReferencedTimeReferenceType = "ANNIVERSARY_WITH_ZONE_SAME_LOCAL_CONVERSION"
	PropertyReferencedTimeReferenceTypeUpdatedAt                              PropertyReferencedTimeReferenceType = "UPDATED_AT"
	PropertyReferencedTimeReferenceTypeValue                                  PropertyReferencedTimeReferenceType = "VALUE"
	PropertyReferencedTimeReferenceTypeValueWithZoneSameLocalConversion       PropertyReferencedTimeReferenceType = "VALUE_WITH_ZONE_SAME_LOCAL_CONVERSION"
)

type PropertyReferencedTimeTimeType string

const (
	PropertyReferencedTimeTimeTypePropertyReferenced PropertyReferencedTimeTimeType = "PROPERTY_REFERENCED"
)

type PropertyReferencedTimeTimezoneSource string

const (
	PropertyReferencedTimeTimezoneSourceCustom PropertyReferencedTimeTimezoneSource = "CUSTOM"
	PropertyReferencedTimeTimezoneSourcePortal PropertyReferencedTimeTimezoneSource = "PORTAL"
	PropertyReferencedTimeTimezoneSourceUser   PropertyReferencedTimeTimezoneSource = "USER"
)

type QuarterReference struct {
	Day   int64 `json:"day" api:"required"`
	Month int64 `json:"month" api:"required"`
	// Any of "QUARTER".
	ReferenceType QuarterReferenceReferenceType `json:"referenceType" api:"required"`
	Hour          int64                         `json:"hour"`
	Millisecond   int64                         `json:"millisecond"`
	Minute        int64                         `json:"minute"`
	Second        int64                         `json:"second"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Day           respjson.Field
		Month         respjson.Field
		ReferenceType respjson.Field
		Hour          respjson.Field
		Millisecond   respjson.Field
		Minute        respjson.Field
		Second        respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r QuarterReference) RawJSON() string { return r.JSON.raw }
func (r *QuarterReference) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type QuarterReferenceReferenceType string

const (
	QuarterReferenceReferenceTypeQuarter QuarterReferenceReferenceType = "QUARTER"
)

type RangedDatePropertyOperation struct {
	IncludeObjectsWithNoValueSet bool   `json:"includeObjectsWithNoValueSet" api:"required"`
	LowerBoundTimestamp          int64  `json:"lowerBoundTimestamp" api:"required"`
	OperationType                string `json:"operationType" api:"required"`
	// Any of "IS_BETWEEN", "IS_NOT_BETWEEN".
	Operator     RangedDatePropertyOperationOperator `json:"operator" api:"required"`
	OperatorName string                              `json:"operatorName" api:"required"`
	// Any of "datetime-ranged".
	PropertyType               RangedDatePropertyOperationPropertyType `json:"propertyType" api:"required"`
	RequiresTimeZoneConversion bool                                    `json:"requiresTimeZoneConversion" api:"required"`
	UpperBoundTimestamp        int64                                   `json:"upperBoundTimestamp" api:"required"`
	DefaultValue               string                                  `json:"defaultValue"`
	RenderSpec                 string                                  `json:"renderSpec"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		IncludeObjectsWithNoValueSet respjson.Field
		LowerBoundTimestamp          respjson.Field
		OperationType                respjson.Field
		Operator                     respjson.Field
		OperatorName                 respjson.Field
		PropertyType                 respjson.Field
		RequiresTimeZoneConversion   respjson.Field
		UpperBoundTimestamp          respjson.Field
		DefaultValue                 respjson.Field
		RenderSpec                   respjson.Field
		ExtraFields                  map[string]respjson.Field
		raw                          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RangedDatePropertyOperation) RawJSON() string { return r.JSON.raw }
func (r *RangedDatePropertyOperation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RangedDatePropertyOperationOperator string

const (
	RangedDatePropertyOperationOperatorIsBetween    RangedDatePropertyOperationOperator = "IS_BETWEEN"
	RangedDatePropertyOperationOperatorIsNotBetween RangedDatePropertyOperationOperator = "IS_NOT_BETWEEN"
)

type RangedDatePropertyOperationPropertyType string

const (
	RangedDatePropertyOperationPropertyTypeDatetimeRanged RangedDatePropertyOperationPropertyType = "datetime-ranged"
)

type RangedNumberPropertyOperation struct {
	IncludeObjectsWithNoValueSet bool    `json:"includeObjectsWithNoValueSet" api:"required"`
	LowerBound                   float64 `json:"lowerBound" api:"required"`
	OperationType                string  `json:"operationType" api:"required"`
	// Any of "IS_BETWEEN", "IS_NOT_BETWEEN".
	Operator     RangedNumberPropertyOperationOperator `json:"operator" api:"required"`
	OperatorName string                                `json:"operatorName" api:"required"`
	// Any of "number-ranged".
	PropertyType RangedNumberPropertyOperationPropertyType `json:"propertyType" api:"required"`
	UpperBound   float64                                   `json:"upperBound" api:"required"`
	DefaultValue string                                    `json:"defaultValue"`
	RenderSpec   string                                    `json:"renderSpec"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		IncludeObjectsWithNoValueSet respjson.Field
		LowerBound                   respjson.Field
		OperationType                respjson.Field
		Operator                     respjson.Field
		OperatorName                 respjson.Field
		PropertyType                 respjson.Field
		UpperBound                   respjson.Field
		DefaultValue                 respjson.Field
		RenderSpec                   respjson.Field
		ExtraFields                  map[string]respjson.Field
		raw                          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RangedNumberPropertyOperation) RawJSON() string { return r.JSON.raw }
func (r *RangedNumberPropertyOperation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RangedNumberPropertyOperationOperator string

const (
	RangedNumberPropertyOperationOperatorIsBetween    RangedNumberPropertyOperationOperator = "IS_BETWEEN"
	RangedNumberPropertyOperationOperatorIsNotBetween RangedNumberPropertyOperationOperator = "IS_NOT_BETWEEN"
)

type RangedNumberPropertyOperationPropertyType string

const (
	RangedNumberPropertyOperationPropertyTypeNumberRanged RangedNumberPropertyOperationPropertyType = "number-ranged"
)

type RangedTimeOperation struct {
	IncludeObjectsWithNoValueSet bool `json:"includeObjectsWithNoValueSet" api:"required"`
	// Any of "EXCLUSIVE", "INCLUSIVE".
	LowerBoundEndpointBehavior RangedTimeOperationLowerBoundEndpointBehavior `json:"lowerBoundEndpointBehavior" api:"required"`
	LowerBoundTimePoint        RangedTimeOperationLowerBoundTimePointUnion   `json:"lowerBoundTimePoint" api:"required"`
	OperationType              string                                        `json:"operationType" api:"required"`
	// Any of "IS_BETWEEN", "IS_NOT_BETWEEN".
	Operator     RangedTimeOperationOperator `json:"operator" api:"required"`
	OperatorName string                      `json:"operatorName" api:"required"`
	// Any of "ANNIVERSARY", "ANNIVERSARY_WITH_ZONE_SAME_LOCAL_CONVERSION",
	// "UPDATED_AT", "VALUE", "VALUE_WITH_ZONE_SAME_LOCAL_CONVERSION".
	PropertyParser RangedTimeOperationPropertyParser `json:"propertyParser" api:"required"`
	// Any of "rangedtime".
	PropertyType RangedTimeOperationPropertyType `json:"propertyType" api:"required"`
	Type         string                          `json:"type" api:"required"`
	// Any of "EXCLUSIVE", "INCLUSIVE".
	UpperBoundEndpointBehavior RangedTimeOperationUpperBoundEndpointBehavior `json:"upperBoundEndpointBehavior" api:"required"`
	UpperBoundTimePoint        RangedTimeOperationUpperBoundTimePointUnion   `json:"upperBoundTimePoint" api:"required"`
	DefaultValue               string                                        `json:"defaultValue"`
	RenderSpec                 string                                        `json:"renderSpec"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		IncludeObjectsWithNoValueSet respjson.Field
		LowerBoundEndpointBehavior   respjson.Field
		LowerBoundTimePoint          respjson.Field
		OperationType                respjson.Field
		Operator                     respjson.Field
		OperatorName                 respjson.Field
		PropertyParser               respjson.Field
		PropertyType                 respjson.Field
		Type                         respjson.Field
		UpperBoundEndpointBehavior   respjson.Field
		UpperBoundTimePoint          respjson.Field
		DefaultValue                 respjson.Field
		RenderSpec                   respjson.Field
		ExtraFields                  map[string]respjson.Field
		raw                          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RangedTimeOperation) RawJSON() string { return r.JSON.raw }
func (r *RangedTimeOperation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RangedTimeOperationLowerBoundEndpointBehavior string

const (
	RangedTimeOperationLowerBoundEndpointBehaviorExclusive RangedTimeOperationLowerBoundEndpointBehavior = "EXCLUSIVE"
	RangedTimeOperationLowerBoundEndpointBehaviorInclusive RangedTimeOperationLowerBoundEndpointBehavior = "INCLUSIVE"
)

// RangedTimeOperationLowerBoundTimePointUnion contains all possible properties and
// values from [DatePoint], [IndexedTimePoint], [PropertyReferencedTime].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type RangedTimeOperationLowerBoundTimePointUnion struct {
	// This field is from variant [DatePoint].
	Day int64 `json:"day"`
	// This field is from variant [DatePoint].
	Month          int64  `json:"month"`
	TimeType       string `json:"timeType"`
	TimezoneSource string `json:"timezoneSource"`
	// This field is from variant [DatePoint].
	Year   int64  `json:"year"`
	ZoneID string `json:"zoneId"`
	// This field is from variant [DatePoint].
	Hour int64 `json:"hour"`
	// This field is from variant [DatePoint].
	Millisecond int64 `json:"millisecond"`
	// This field is from variant [DatePoint].
	Minute int64 `json:"minute"`
	// This field is from variant [DatePoint].
	Second int64 `json:"second"`
	// This field is from variant [IndexedTimePoint].
	IndexReference IndexedTimePointIndexReferenceUnion `json:"indexReference"`
	// This field is from variant [IndexedTimePoint].
	Offset IndexOffset `json:"offset"`
	// This field is from variant [IndexedTimePoint].
	ShouldGenerateRefreshTime bool `json:"shouldGenerateRefreshTime"`
	// This field is from variant [PropertyReferencedTime].
	Property string `json:"property"`
	// This field is from variant [PropertyReferencedTime].
	ReferenceType PropertyReferencedTimeReferenceType `json:"referenceType"`
	JSON          struct {
		Day                       respjson.Field
		Month                     respjson.Field
		TimeType                  respjson.Field
		TimezoneSource            respjson.Field
		Year                      respjson.Field
		ZoneID                    respjson.Field
		Hour                      respjson.Field
		Millisecond               respjson.Field
		Minute                    respjson.Field
		Second                    respjson.Field
		IndexReference            respjson.Field
		Offset                    respjson.Field
		ShouldGenerateRefreshTime respjson.Field
		Property                  respjson.Field
		ReferenceType             respjson.Field
		raw                       string
	} `json:"-"`
}

func (u RangedTimeOperationLowerBoundTimePointUnion) AsDate() (v DatePoint) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u RangedTimeOperationLowerBoundTimePointUnion) AsIndexed() (v IndexedTimePoint) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u RangedTimeOperationLowerBoundTimePointUnion) AsPropertyReferenced() (v PropertyReferencedTime) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u RangedTimeOperationLowerBoundTimePointUnion) RawJSON() string { return u.JSON.raw }

func (r *RangedTimeOperationLowerBoundTimePointUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RangedTimeOperationOperator string

const (
	RangedTimeOperationOperatorIsBetween    RangedTimeOperationOperator = "IS_BETWEEN"
	RangedTimeOperationOperatorIsNotBetween RangedTimeOperationOperator = "IS_NOT_BETWEEN"
)

type RangedTimeOperationPropertyParser string

const (
	RangedTimeOperationPropertyParserAnniversary                            RangedTimeOperationPropertyParser = "ANNIVERSARY"
	RangedTimeOperationPropertyParserAnniversaryWithZoneSameLocalConversion RangedTimeOperationPropertyParser = "ANNIVERSARY_WITH_ZONE_SAME_LOCAL_CONVERSION"
	RangedTimeOperationPropertyParserUpdatedAt                              RangedTimeOperationPropertyParser = "UPDATED_AT"
	RangedTimeOperationPropertyParserValue                                  RangedTimeOperationPropertyParser = "VALUE"
	RangedTimeOperationPropertyParserValueWithZoneSameLocalConversion       RangedTimeOperationPropertyParser = "VALUE_WITH_ZONE_SAME_LOCAL_CONVERSION"
)

type RangedTimeOperationPropertyType string

const (
	RangedTimeOperationPropertyTypeRangedtime RangedTimeOperationPropertyType = "rangedtime"
)

type RangedTimeOperationUpperBoundEndpointBehavior string

const (
	RangedTimeOperationUpperBoundEndpointBehaviorExclusive RangedTimeOperationUpperBoundEndpointBehavior = "EXCLUSIVE"
	RangedTimeOperationUpperBoundEndpointBehaviorInclusive RangedTimeOperationUpperBoundEndpointBehavior = "INCLUSIVE"
)

// RangedTimeOperationUpperBoundTimePointUnion contains all possible properties and
// values from [DatePoint], [IndexedTimePoint], [PropertyReferencedTime].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type RangedTimeOperationUpperBoundTimePointUnion struct {
	// This field is from variant [DatePoint].
	Day int64 `json:"day"`
	// This field is from variant [DatePoint].
	Month          int64  `json:"month"`
	TimeType       string `json:"timeType"`
	TimezoneSource string `json:"timezoneSource"`
	// This field is from variant [DatePoint].
	Year   int64  `json:"year"`
	ZoneID string `json:"zoneId"`
	// This field is from variant [DatePoint].
	Hour int64 `json:"hour"`
	// This field is from variant [DatePoint].
	Millisecond int64 `json:"millisecond"`
	// This field is from variant [DatePoint].
	Minute int64 `json:"minute"`
	// This field is from variant [DatePoint].
	Second int64 `json:"second"`
	// This field is from variant [IndexedTimePoint].
	IndexReference IndexedTimePointIndexReferenceUnion `json:"indexReference"`
	// This field is from variant [IndexedTimePoint].
	Offset IndexOffset `json:"offset"`
	// This field is from variant [IndexedTimePoint].
	ShouldGenerateRefreshTime bool `json:"shouldGenerateRefreshTime"`
	// This field is from variant [PropertyReferencedTime].
	Property string `json:"property"`
	// This field is from variant [PropertyReferencedTime].
	ReferenceType PropertyReferencedTimeReferenceType `json:"referenceType"`
	JSON          struct {
		Day                       respjson.Field
		Month                     respjson.Field
		TimeType                  respjson.Field
		TimezoneSource            respjson.Field
		Year                      respjson.Field
		ZoneID                    respjson.Field
		Hour                      respjson.Field
		Millisecond               respjson.Field
		Minute                    respjson.Field
		Second                    respjson.Field
		IndexReference            respjson.Field
		Offset                    respjson.Field
		ShouldGenerateRefreshTime respjson.Field
		Property                  respjson.Field
		ReferenceType             respjson.Field
		raw                       string
	} `json:"-"`
}

func (u RangedTimeOperationUpperBoundTimePointUnion) AsDate() (v DatePoint) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u RangedTimeOperationUpperBoundTimePointUnion) AsIndexed() (v IndexedTimePoint) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u RangedTimeOperationUpperBoundTimePointUnion) AsPropertyReferenced() (v PropertyReferencedTime) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u RangedTimeOperationUpperBoundTimePointUnion) RawJSON() string { return u.JSON.raw }

func (r *RangedTimeOperationUpperBoundTimePointUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RegexPropertyOperation struct {
	CaseSensitive                bool   `json:"caseSensitive" api:"required"`
	IncludeObjectsWithNoValueSet bool   `json:"includeObjectsWithNoValueSet" api:"required"`
	OperationType                string `json:"operationType" api:"required"`
	// Any of "DOES_NOT_MATCH_REGEX", "MATCHES_REGEX".
	Operator     RegexPropertyOperationOperator `json:"operator" api:"required"`
	OperatorName string                         `json:"operatorName" api:"required"`
	Pattern      string                         `json:"pattern" api:"required"`
	// Any of "regex".
	PropertyType RegexPropertyOperationPropertyType `json:"propertyType" api:"required"`
	DefaultValue string                             `json:"defaultValue"`
	RenderSpec   string                             `json:"renderSpec"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CaseSensitive                respjson.Field
		IncludeObjectsWithNoValueSet respjson.Field
		OperationType                respjson.Field
		Operator                     respjson.Field
		OperatorName                 respjson.Field
		Pattern                      respjson.Field
		PropertyType                 respjson.Field
		DefaultValue                 respjson.Field
		RenderSpec                   respjson.Field
		ExtraFields                  map[string]respjson.Field
		raw                          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RegexPropertyOperation) RawJSON() string { return r.JSON.raw }
func (r *RegexPropertyOperation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RegexPropertyOperationOperator string

const (
	RegexPropertyOperationOperatorDoesNotMatchRegex RegexPropertyOperationOperator = "DOES_NOT_MATCH_REGEX"
	RegexPropertyOperationOperatorMatchesRegex      RegexPropertyOperationOperator = "MATCHES_REGEX"
)

type RegexPropertyOperationPropertyType string

const (
	RegexPropertyOperationPropertyTypeRegex RegexPropertyOperationPropertyType = "regex"
)

type RelativeComparativeTimestampRefineBy struct {
	// Any of "AFTER", "BEFORE".
	Comparison RelativeComparativeTimestampRefineByComparison `json:"comparison" api:"required"`
	TimeOffset TimeOffset                                     `json:"timeOffset" api:"required"`
	// Any of "RelativeComparativeTimestampRefineBy".
	Type RelativeComparativeTimestampRefineByType `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Comparison  respjson.Field
		TimeOffset  respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RelativeComparativeTimestampRefineBy) RawJSON() string { return r.JSON.raw }
func (r *RelativeComparativeTimestampRefineBy) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RelativeComparativeTimestampRefineByComparison string

const (
	RelativeComparativeTimestampRefineByComparisonAfter  RelativeComparativeTimestampRefineByComparison = "AFTER"
	RelativeComparativeTimestampRefineByComparisonBefore RelativeComparativeTimestampRefineByComparison = "BEFORE"
)

type RelativeComparativeTimestampRefineByType string

const (
	RelativeComparativeTimestampRefineByTypeRelativeComparativeTimestampRefineBy RelativeComparativeTimestampRefineByType = "RelativeComparativeTimestampRefineBy"
)

type RelativeRangedTimestampRefineBy struct {
	LowerBoundOffset TimeOffset `json:"lowerBoundOffset" api:"required"`
	// Any of "BETWEEN", "NOT_BETWEEN".
	RangeType RelativeRangedTimestampRefineByRangeType `json:"rangeType" api:"required"`
	// Any of "RelativeRangedTimestampRefineBy".
	Type             RelativeRangedTimestampRefineByType `json:"type" api:"required"`
	UpperBoundOffset TimeOffset                          `json:"upperBoundOffset" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		LowerBoundOffset respjson.Field
		RangeType        respjson.Field
		Type             respjson.Field
		UpperBoundOffset respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RelativeRangedTimestampRefineBy) RawJSON() string { return r.JSON.raw }
func (r *RelativeRangedTimestampRefineBy) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RelativeRangedTimestampRefineByRangeType string

const (
	RelativeRangedTimestampRefineByRangeTypeBetween    RelativeRangedTimestampRefineByRangeType = "BETWEEN"
	RelativeRangedTimestampRefineByRangeTypeNotBetween RelativeRangedTimestampRefineByRangeType = "NOT_BETWEEN"
)

type RelativeRangedTimestampRefineByType string

const (
	RelativeRangedTimestampRefineByTypeRelativeRangedTimestampRefineBy RelativeRangedTimestampRefineByType = "RelativeRangedTimestampRefineBy"
)

type RollingDateRangePropertyOperation struct {
	IncludeObjectsWithNoValueSet bool   `json:"includeObjectsWithNoValueSet" api:"required"`
	NumberOfDays                 int64  `json:"numberOfDays" api:"required"`
	OperationType                string `json:"operationType" api:"required"`
	// Any of "IS_LESS_THAN_X_DAYS_AGO", "IS_LESS_THAN_X_DAYS_FROM_NOW",
	// "IS_MORE_THAN_X_DAYS_AGO", "IS_MORE_THAN_X_DAYS_FROM_NOW".
	Operator     RollingDateRangePropertyOperationOperator `json:"operator" api:"required"`
	OperatorName string                                    `json:"operatorName" api:"required"`
	// Any of "datetime-rolling".
	PropertyType               RollingDateRangePropertyOperationPropertyType `json:"propertyType" api:"required"`
	RequiresTimeZoneConversion bool                                          `json:"requiresTimeZoneConversion" api:"required"`
	DefaultValue               string                                        `json:"defaultValue"`
	RenderSpec                 string                                        `json:"renderSpec"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		IncludeObjectsWithNoValueSet respjson.Field
		NumberOfDays                 respjson.Field
		OperationType                respjson.Field
		Operator                     respjson.Field
		OperatorName                 respjson.Field
		PropertyType                 respjson.Field
		RequiresTimeZoneConversion   respjson.Field
		DefaultValue                 respjson.Field
		RenderSpec                   respjson.Field
		ExtraFields                  map[string]respjson.Field
		raw                          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RollingDateRangePropertyOperation) RawJSON() string { return r.JSON.raw }
func (r *RollingDateRangePropertyOperation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RollingDateRangePropertyOperationOperator string

const (
	RollingDateRangePropertyOperationOperatorIsLessThanXDaysAgo     RollingDateRangePropertyOperationOperator = "IS_LESS_THAN_X_DAYS_AGO"
	RollingDateRangePropertyOperationOperatorIsLessThanXDaysFromNow RollingDateRangePropertyOperationOperator = "IS_LESS_THAN_X_DAYS_FROM_NOW"
	RollingDateRangePropertyOperationOperatorIsMoreThanXDaysAgo     RollingDateRangePropertyOperationOperator = "IS_MORE_THAN_X_DAYS_AGO"
	RollingDateRangePropertyOperationOperatorIsMoreThanXDaysFromNow RollingDateRangePropertyOperationOperator = "IS_MORE_THAN_X_DAYS_FROM_NOW"
)

type RollingDateRangePropertyOperationPropertyType string

const (
	RollingDateRangePropertyOperationPropertyTypeDatetimeRolling RollingDateRangePropertyOperationPropertyType = "datetime-rolling"
)

type RollingPropertyUpdatedOperation struct {
	IncludeObjectsWithNoValueSet bool   `json:"includeObjectsWithNoValueSet" api:"required"`
	NumberOfDays                 int64  `json:"numberOfDays" api:"required"`
	OperationType                string `json:"operationType" api:"required"`
	// Any of "NOT_UPDATED_IN_LAST_X_DAYS", "UPDATED_IN_LAST_X_DAYS".
	Operator     RollingPropertyUpdatedOperationOperator `json:"operator" api:"required"`
	OperatorName string                                  `json:"operatorName" api:"required"`
	// Any of "rolling-property-updated".
	PropertyType RollingPropertyUpdatedOperationPropertyType `json:"propertyType" api:"required"`
	DefaultValue string                                      `json:"defaultValue"`
	RenderSpec   string                                      `json:"renderSpec"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		IncludeObjectsWithNoValueSet respjson.Field
		NumberOfDays                 respjson.Field
		OperationType                respjson.Field
		Operator                     respjson.Field
		OperatorName                 respjson.Field
		PropertyType                 respjson.Field
		DefaultValue                 respjson.Field
		RenderSpec                   respjson.Field
		ExtraFields                  map[string]respjson.Field
		raw                          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RollingPropertyUpdatedOperation) RawJSON() string { return r.JSON.raw }
func (r *RollingPropertyUpdatedOperation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RollingPropertyUpdatedOperationOperator string

const (
	RollingPropertyUpdatedOperationOperatorNotUpdatedInLastXDays RollingPropertyUpdatedOperationOperator = "NOT_UPDATED_IN_LAST_X_DAYS"
	RollingPropertyUpdatedOperationOperatorUpdatedInLastXDays    RollingPropertyUpdatedOperationOperator = "UPDATED_IN_LAST_X_DAYS"
)

type RollingPropertyUpdatedOperationPropertyType string

const (
	RollingPropertyUpdatedOperationPropertyTypeRollingPropertyUpdated RollingPropertyUpdatedOperationPropertyType = "rolling-property-updated"
)

type SetOccurrencesRefineBy struct {
	// Any of "ALL", "ALL_INCLUDE_EMPTY", "ANY", "ANY_INCLUDE_EMPTY", "NONE",
	// "NONE_EXCLUDE_EMPTY".
	SetType SetOccurrencesRefineBySetType `json:"setType" api:"required"`
	// Any of "SetOccurrencesRefineBy".
	Type SetOccurrencesRefineByType `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		SetType     respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SetOccurrencesRefineBy) RawJSON() string { return r.JSON.raw }
func (r *SetOccurrencesRefineBy) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SetOccurrencesRefineBySetType string

const (
	SetOccurrencesRefineBySetTypeAll              SetOccurrencesRefineBySetType = "ALL"
	SetOccurrencesRefineBySetTypeAllIncludeEmpty  SetOccurrencesRefineBySetType = "ALL_INCLUDE_EMPTY"
	SetOccurrencesRefineBySetTypeAny              SetOccurrencesRefineBySetType = "ANY"
	SetOccurrencesRefineBySetTypeAnyIncludeEmpty  SetOccurrencesRefineBySetType = "ANY_INCLUDE_EMPTY"
	SetOccurrencesRefineBySetTypeNone             SetOccurrencesRefineBySetType = "NONE"
	SetOccurrencesRefineBySetTypeNoneExcludeEmpty SetOccurrencesRefineBySetType = "NONE_EXCLUDE_EMPTY"
)

type SetOccurrencesRefineByType string

const (
	SetOccurrencesRefineByTypeSetOccurrencesRefineBy SetOccurrencesRefineByType = "SetOccurrencesRefineBy"
)

type StringPropertyOperation struct {
	IncludeObjectsWithNoValueSet bool   `json:"includeObjectsWithNoValueSet" api:"required"`
	OperationType                string `json:"operationType" api:"required"`
	// Any of "CONTAINS", "DOES_NOT_CONTAIN", "ENDS_WITH", "HAS_EVER_BEEN_EQUAL_TO",
	// "HAS_EVER_CONTAINED", "HAS_NEVER_BEEN_EQUAL_TO", "HAS_NEVER_CONTAINED",
	// "IS_EQUAL_TO", "IS_NOT_EQUAL_TO", "STARTS_WITH".
	Operator     StringPropertyOperationOperator `json:"operator" api:"required"`
	OperatorName string                          `json:"operatorName" api:"required"`
	// Any of "string".
	PropertyType StringPropertyOperationPropertyType `json:"propertyType" api:"required"`
	Value        string                              `json:"value" api:"required"`
	DefaultValue string                              `json:"defaultValue"`
	RenderSpec   string                              `json:"renderSpec"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		IncludeObjectsWithNoValueSet respjson.Field
		OperationType                respjson.Field
		Operator                     respjson.Field
		OperatorName                 respjson.Field
		PropertyType                 respjson.Field
		Value                        respjson.Field
		DefaultValue                 respjson.Field
		RenderSpec                   respjson.Field
		ExtraFields                  map[string]respjson.Field
		raw                          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r StringPropertyOperation) RawJSON() string { return r.JSON.raw }
func (r *StringPropertyOperation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type StringPropertyOperationOperator string

const (
	StringPropertyOperationOperatorContains            StringPropertyOperationOperator = "CONTAINS"
	StringPropertyOperationOperatorDoesNotContain      StringPropertyOperationOperator = "DOES_NOT_CONTAIN"
	StringPropertyOperationOperatorEndsWith            StringPropertyOperationOperator = "ENDS_WITH"
	StringPropertyOperationOperatorHasEverBeenEqualTo  StringPropertyOperationOperator = "HAS_EVER_BEEN_EQUAL_TO"
	StringPropertyOperationOperatorHasEverContained    StringPropertyOperationOperator = "HAS_EVER_CONTAINED"
	StringPropertyOperationOperatorHasNeverBeenEqualTo StringPropertyOperationOperator = "HAS_NEVER_BEEN_EQUAL_TO"
	StringPropertyOperationOperatorHasNeverContained   StringPropertyOperationOperator = "HAS_NEVER_CONTAINED"
	StringPropertyOperationOperatorIsEqualTo           StringPropertyOperationOperator = "IS_EQUAL_TO"
	StringPropertyOperationOperatorIsNotEqualTo        StringPropertyOperationOperator = "IS_NOT_EQUAL_TO"
	StringPropertyOperationOperatorStartsWith          StringPropertyOperationOperator = "STARTS_WITH"
)

type StringPropertyOperationPropertyType string

const (
	StringPropertyOperationPropertyTypeString StringPropertyOperationPropertyType = "string"
)

type TimeOffset struct {
	Amount int64 `json:"amount" api:"required"`
	// Any of "FUTURE", "PAST".
	OffsetDirection TimeOffsetOffsetDirection `json:"offsetDirection" api:"required"`
	// Any of "DAYS", "HOURS", "MINUTES", "WEEKS".
	TimeUnit TimeOffsetTimeUnit `json:"timeUnit" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Amount          respjson.Field
		OffsetDirection respjson.Field
		TimeUnit        respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TimeOffset) RawJSON() string { return r.JSON.raw }
func (r *TimeOffset) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TimeOffsetOffsetDirection string

const (
	TimeOffsetOffsetDirectionFuture TimeOffsetOffsetDirection = "FUTURE"
	TimeOffsetOffsetDirectionPast   TimeOffsetOffsetDirection = "PAST"
)

type TimeOffsetTimeUnit string

const (
	TimeOffsetTimeUnitDays    TimeOffsetTimeUnit = "DAYS"
	TimeOffsetTimeUnitHours   TimeOffsetTimeUnit = "HOURS"
	TimeOffsetTimeUnitMinutes TimeOffsetTimeUnit = "MINUTES"
	TimeOffsetTimeUnitWeeks   TimeOffsetTimeUnit = "WEEKS"
)

type TimePointOperation struct {
	// Any of "EXCLUSIVE", "INCLUSIVE".
	EndpointBehavior             TimePointOperationEndpointBehavior `json:"endpointBehavior" api:"required"`
	IncludeObjectsWithNoValueSet bool                               `json:"includeObjectsWithNoValueSet" api:"required"`
	OperationType                string                             `json:"operationType" api:"required"`
	// Any of "IS_AFTER", "IS_BEFORE".
	Operator     TimePointOperationOperator `json:"operator" api:"required"`
	OperatorName string                     `json:"operatorName" api:"required"`
	// Any of "ANNIVERSARY", "ANNIVERSARY_WITH_ZONE_SAME_LOCAL_CONVERSION",
	// "UPDATED_AT", "VALUE", "VALUE_WITH_ZONE_SAME_LOCAL_CONVERSION".
	PropertyParser TimePointOperationPropertyParser `json:"propertyParser" api:"required"`
	// Any of "timepoint".
	PropertyType TimePointOperationPropertyType   `json:"propertyType" api:"required"`
	TimePoint    TimePointOperationTimePointUnion `json:"timePoint" api:"required"`
	Type         string                           `json:"type" api:"required"`
	DefaultValue string                           `json:"defaultValue"`
	RenderSpec   string                           `json:"renderSpec"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		EndpointBehavior             respjson.Field
		IncludeObjectsWithNoValueSet respjson.Field
		OperationType                respjson.Field
		Operator                     respjson.Field
		OperatorName                 respjson.Field
		PropertyParser               respjson.Field
		PropertyType                 respjson.Field
		TimePoint                    respjson.Field
		Type                         respjson.Field
		DefaultValue                 respjson.Field
		RenderSpec                   respjson.Field
		ExtraFields                  map[string]respjson.Field
		raw                          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TimePointOperation) RawJSON() string { return r.JSON.raw }
func (r *TimePointOperation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TimePointOperationEndpointBehavior string

const (
	TimePointOperationEndpointBehaviorExclusive TimePointOperationEndpointBehavior = "EXCLUSIVE"
	TimePointOperationEndpointBehaviorInclusive TimePointOperationEndpointBehavior = "INCLUSIVE"
)

type TimePointOperationOperator string

const (
	TimePointOperationOperatorIsAfter  TimePointOperationOperator = "IS_AFTER"
	TimePointOperationOperatorIsBefore TimePointOperationOperator = "IS_BEFORE"
)

type TimePointOperationPropertyParser string

const (
	TimePointOperationPropertyParserAnniversary                            TimePointOperationPropertyParser = "ANNIVERSARY"
	TimePointOperationPropertyParserAnniversaryWithZoneSameLocalConversion TimePointOperationPropertyParser = "ANNIVERSARY_WITH_ZONE_SAME_LOCAL_CONVERSION"
	TimePointOperationPropertyParserUpdatedAt                              TimePointOperationPropertyParser = "UPDATED_AT"
	TimePointOperationPropertyParserValue                                  TimePointOperationPropertyParser = "VALUE"
	TimePointOperationPropertyParserValueWithZoneSameLocalConversion       TimePointOperationPropertyParser = "VALUE_WITH_ZONE_SAME_LOCAL_CONVERSION"
)

type TimePointOperationPropertyType string

const (
	TimePointOperationPropertyTypeTimepoint TimePointOperationPropertyType = "timepoint"
)

// TimePointOperationTimePointUnion contains all possible properties and values
// from [DatePoint], [IndexedTimePoint], [PropertyReferencedTime].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type TimePointOperationTimePointUnion struct {
	// This field is from variant [DatePoint].
	Day int64 `json:"day"`
	// This field is from variant [DatePoint].
	Month          int64  `json:"month"`
	TimeType       string `json:"timeType"`
	TimezoneSource string `json:"timezoneSource"`
	// This field is from variant [DatePoint].
	Year   int64  `json:"year"`
	ZoneID string `json:"zoneId"`
	// This field is from variant [DatePoint].
	Hour int64 `json:"hour"`
	// This field is from variant [DatePoint].
	Millisecond int64 `json:"millisecond"`
	// This field is from variant [DatePoint].
	Minute int64 `json:"minute"`
	// This field is from variant [DatePoint].
	Second int64 `json:"second"`
	// This field is from variant [IndexedTimePoint].
	IndexReference IndexedTimePointIndexReferenceUnion `json:"indexReference"`
	// This field is from variant [IndexedTimePoint].
	Offset IndexOffset `json:"offset"`
	// This field is from variant [IndexedTimePoint].
	ShouldGenerateRefreshTime bool `json:"shouldGenerateRefreshTime"`
	// This field is from variant [PropertyReferencedTime].
	Property string `json:"property"`
	// This field is from variant [PropertyReferencedTime].
	ReferenceType PropertyReferencedTimeReferenceType `json:"referenceType"`
	JSON          struct {
		Day                       respjson.Field
		Month                     respjson.Field
		TimeType                  respjson.Field
		TimezoneSource            respjson.Field
		Year                      respjson.Field
		ZoneID                    respjson.Field
		Hour                      respjson.Field
		Millisecond               respjson.Field
		Minute                    respjson.Field
		Second                    respjson.Field
		IndexReference            respjson.Field
		Offset                    respjson.Field
		ShouldGenerateRefreshTime respjson.Field
		Property                  respjson.Field
		ReferenceType             respjson.Field
		raw                       string
	} `json:"-"`
}

func (u TimePointOperationTimePointUnion) AsDate() (v DatePoint) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TimePointOperationTimePointUnion) AsIndexed() (v IndexedTimePoint) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TimePointOperationTimePointUnion) AsPropertyReferenced() (v PropertyReferencedTime) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u TimePointOperationTimePointUnion) RawJSON() string { return u.JSON.raw }

func (r *TimePointOperationTimePointUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TodayReference struct {
	// Any of "TODAY".
	ReferenceType TodayReferenceReferenceType `json:"referenceType" api:"required"`
	Hour          int64                       `json:"hour"`
	Millisecond   int64                       `json:"millisecond"`
	Minute        int64                       `json:"minute"`
	Second        int64                       `json:"second"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ReferenceType respjson.Field
		Hour          respjson.Field
		Millisecond   respjson.Field
		Minute        respjson.Field
		Second        respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TodayReference) RawJSON() string { return r.JSON.raw }
func (r *TodayReference) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TodayReferenceReferenceType string

const (
	TodayReferenceReferenceTypeToday TodayReferenceReferenceType = "TODAY"
)

type WeekReference struct {
	// Any of "FRIDAY", "MONDAY", "SATURDAY", "SUNDAY", "THURSDAY", "TUESDAY",
	// "WEDNESDAY".
	DayOfWeek WeekReferenceDayOfWeek `json:"dayOfWeek" api:"required"`
	// Any of "WEEK".
	ReferenceType WeekReferenceReferenceType `json:"referenceType" api:"required"`
	Hour          int64                      `json:"hour"`
	Millisecond   int64                      `json:"millisecond"`
	Minute        int64                      `json:"minute"`
	Second        int64                      `json:"second"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DayOfWeek     respjson.Field
		ReferenceType respjson.Field
		Hour          respjson.Field
		Millisecond   respjson.Field
		Minute        respjson.Field
		Second        respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WeekReference) RawJSON() string { return r.JSON.raw }
func (r *WeekReference) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WeekReferenceDayOfWeek string

const (
	WeekReferenceDayOfWeekFriday    WeekReferenceDayOfWeek = "FRIDAY"
	WeekReferenceDayOfWeekMonday    WeekReferenceDayOfWeek = "MONDAY"
	WeekReferenceDayOfWeekSaturday  WeekReferenceDayOfWeek = "SATURDAY"
	WeekReferenceDayOfWeekSunday    WeekReferenceDayOfWeek = "SUNDAY"
	WeekReferenceDayOfWeekThursday  WeekReferenceDayOfWeek = "THURSDAY"
	WeekReferenceDayOfWeekTuesday   WeekReferenceDayOfWeek = "TUESDAY"
	WeekReferenceDayOfWeekWednesday WeekReferenceDayOfWeek = "WEDNESDAY"
)

type WeekReferenceReferenceType string

const (
	WeekReferenceReferenceTypeWeek WeekReferenceReferenceType = "WEEK"
)

type YearReference struct {
	Day   int64 `json:"day" api:"required"`
	Month int64 `json:"month" api:"required"`
	// Any of "YEAR".
	ReferenceType YearReferenceReferenceType `json:"referenceType" api:"required"`
	Hour          int64                      `json:"hour"`
	Millisecond   int64                      `json:"millisecond"`
	Minute        int64                      `json:"minute"`
	Second        int64                      `json:"second"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Day           respjson.Field
		Month         respjson.Field
		ReferenceType respjson.Field
		Hour          respjson.Field
		Millisecond   respjson.Field
		Minute        respjson.Field
		Second        respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r YearReference) RawJSON() string { return r.JSON.raw }
func (r *YearReference) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type YearReferenceReferenceType string

const (
	YearReferenceReferenceTypeYear YearReferenceReferenceType = "YEAR"
)

type SendNewEventDefinitionParams struct {
	ExternalBehavioralEventTypeDefinitionEgg ExternalBehavioralEventTypeDefinitionEggParam
	paramObj
}

func (r SendNewEventDefinitionParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.ExternalBehavioralEventTypeDefinitionEgg)
}
func (r *SendNewEventDefinitionParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.ExternalBehavioralEventTypeDefinitionEgg)
}

type SendNewEventDefinitionPropertyParams struct {
	ExternalBehavioralEventPropertyCreate ExternalBehavioralEventPropertyCreateParam
	paramObj
}

func (r SendNewEventDefinitionPropertyParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.ExternalBehavioralEventPropertyCreate)
}
func (r *SendNewEventDefinitionPropertyParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.ExternalBehavioralEventPropertyCreate)
}

type SendDeleteEventDefinitionPropertyParams struct {
	EventName string `path:"eventName" api:"required" json:"-"`
	paramObj
}

type SendListEventDefinitionsParams struct {
	// The paging cursor token of the last successfully read resource will be returned
	// as the `paging.next.after` JSON property of a paged response containing more
	// results.
	After             param.Opt[string] `query:"after,omitzero" json:"-"`
	IncludeProperties param.Opt[bool]   `query:"includeProperties,omitzero" json:"-"`
	// The maximum number of results to display per page.
	Limit        param.Opt[int64]  `query:"limit,omitzero" json:"-"`
	SearchString param.Opt[string] `query:"searchString,omitzero" json:"-"`
	SortOrder    param.Opt[string] `query:"sortOrder,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SendListEventDefinitionsParams]'s query parameters as
// `url.Values`.
func (r SendListEventDefinitionsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SendSendEventParams struct {
	BehavioralEventHTTPCompletionRequest BehavioralEventHTTPCompletionRequestParam
	paramObj
}

func (r SendSendEventParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.BehavioralEventHTTPCompletionRequest)
}
func (r *SendSendEventParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.BehavioralEventHTTPCompletionRequest)
}

type SendSendEventBatchParams struct {
	BatchedBehavioralEventHTTPCompletionRequest BatchedBehavioralEventHTTPCompletionRequestParam
	paramObj
}

func (r SendSendEventBatchParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.BatchedBehavioralEventHTTPCompletionRequest)
}
func (r *SendSendEventBatchParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.BatchedBehavioralEventHTTPCompletionRequest)
}

type SendUpdateEventDefinitionParams struct {
	ExternalBehavioralEventTypeDefinitionPatch ExternalBehavioralEventTypeDefinitionPatchParam
	paramObj
}

func (r SendUpdateEventDefinitionParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.ExternalBehavioralEventTypeDefinitionPatch)
}
func (r *SendUpdateEventDefinitionParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.ExternalBehavioralEventTypeDefinitionPatch)
}

type SendUpdateEventDefinitionPropertyParams struct {
	EventName                                      string `path:"eventName" api:"required" json:"-"`
	ExternalBehavioralEventPropertyDefinitionPatch ExternalBehavioralEventPropertyDefinitionPatchParam
	paramObj
}

func (r SendUpdateEventDefinitionPropertyParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.ExternalBehavioralEventPropertyDefinitionPatch)
}
func (r *SendUpdateEventDefinitionPropertyParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.ExternalBehavioralEventPropertyDefinitionPatch)
}
