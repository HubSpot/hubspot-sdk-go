// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cms

import (
	"encoding/json"
	"time"

	"github.com/stainless-sdks/hubspot-sdk-go/crm"
	"github.com/stainless-sdks/hubspot-sdk-go/internal/apijson"
	"github.com/stainless-sdks/hubspot-sdk-go/option"
	"github.com/stainless-sdks/hubspot-sdk-go/packages/param"
	"github.com/stainless-sdks/hubspot-sdk-go/packages/respjson"
	"github.com/stainless-sdks/hubspot-sdk-go/shared"
)

// MediaBridgeService contains methods and other services that help with
// interacting with the hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMediaBridgeService] method instead.
type MediaBridgeService struct {
	Options            []option.RequestOption
	Events             MediaBridgeEventService
	Groups             MediaBridgeGroupService
	IntegratorSettings MediaBridgeIntegratorSettingService
	Properties         MediaBridgePropertyService
	Schemas            MediaBridgeSchemaService
}

// NewMediaBridgeService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewMediaBridgeService(opts ...option.RequestOption) (r MediaBridgeService) {
	r = MediaBridgeService{}
	r.Options = opts
	r.Events = NewMediaBridgeEventService(opts...)
	r.Groups = NewMediaBridgeGroupService(opts...)
	r.IntegratorSettings = NewMediaBridgeIntegratorSettingService(opts...)
	r.Properties = NewMediaBridgePropertyService(opts...)
	r.Schemas = NewMediaBridgeSchemaService(opts...)
	return
}

// The properties TotalPercentPlayed, TotalSecondsPlayed are required.
type AttentionSpanCalculatedValuesParam struct {
	TotalPercentPlayed float64 `json:"totalPercentPlayed,required"`
	TotalSecondsPlayed int64   `json:"totalSecondsPlayed,required"`
	paramObj
}

func (r AttentionSpanCalculatedValuesParam) MarshalJSON() (data []byte, err error) {
	type shadow AttentionSpanCalculatedValuesParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AttentionSpanCalculatedValuesParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AttentionSpanEvent struct {
	// The ID of the contact in HubSpot’s system that consumed the media. This can be
	// fetched using HubSpot's Get contact by usertoken (utk) API. The API also
	// supports supplying a usertoken, and will handle converting this into a contact
	// ID automatically.
	ContactID                    int64  `json:"contactId,required"`
	MediaBridgeID                int64  `json:"mediaBridgeId,required"`
	MediaBridgeObjectCoordinates string `json:"mediaBridgeObjectCoordinates,required"`
	MediaBridgeObjectTypeID      string `json:"mediaBridgeObjectTypeId,required"`
	MediaName                    string `json:"mediaName,required"`
	// Any of "AUDIO", "DOCUMENT", "IMAGE", "OTHER", "VIDEO".
	MediaType AttentionSpanEventMediaType `json:"mediaType,required"`
	// The timestamp at which this event occurred, in milliseconds since the epoch.
	OccurredTimestamp int64  `json:"occurredTimestamp,required"`
	PercentRange      string `json:"percentRange,required"`
	// The ID of the HubSpot account.
	PortalID   int64  `json:"portalId,required"`
	ProviderID int64  `json:"providerId,required"`
	SessionID  string `json:"sessionId,required"`
	// The percent of the media that the user consumed. Providers may calculate this
	// differently depending on how they consider repeated views of the same portion of
	// media. For this reason, the API will not attempt to validate totalPercentWatched
	// against the attention span information for the event. If it is missing, HubSpot
	// will calculate this from the attention span map as follows: (number of spans
	// with a value of 1 or more)/(Total number of spans).
	TotalPercentPlayed float64 `json:"totalPercentPlayed,required"`
	MediaURL           string  `json:"mediaUrl"`
	// The ID of the page, if hosted on HubSpot. Required for HubSpot pages.
	PageID int64 `json:"pageId"`
	// The name of the page. Required if the page is not hosted on HubSpot.
	PageName              string `json:"pageName"`
	PageObjectCoordinates string `json:"pageObjectCoordinates"`
	// The URL of the page that an event happened on. Required if the page is not
	// hosted on HubSpot.
	PageURL string `json:"pageUrl"`
	// This is the raw data which provides the most granular data about spans of the
	// media, and how many times each span was consumed by the user. For example, for a
	// 10 second video where each second is a span, if a visitor watches the first 5
	// seconds of the video, then restarts the video and watches the first 2 seconds
	// again, the resulting `rawDataString` would be
	// `“0=2;1=2;2=1;3=1;4=1;5=0;6=0;7=0;8=0;9=0;”`.
	RawData string `json:"rawData"`
	// The seconds that a user spent consuming the media. The media bridge calculates
	// this as `totalPercentPlayed`\*`mediaDuration`. If a provider would like this to
	// be calculated differently, they can provide the pre-calculated value when they
	// create the event.
	TotalSecondsPlayed int64 `json:"totalSecondsPlayed"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ContactID                    respjson.Field
		MediaBridgeID                respjson.Field
		MediaBridgeObjectCoordinates respjson.Field
		MediaBridgeObjectTypeID      respjson.Field
		MediaName                    respjson.Field
		MediaType                    respjson.Field
		OccurredTimestamp            respjson.Field
		PercentRange                 respjson.Field
		PortalID                     respjson.Field
		ProviderID                   respjson.Field
		SessionID                    respjson.Field
		TotalPercentPlayed           respjson.Field
		MediaURL                     respjson.Field
		PageID                       respjson.Field
		PageName                     respjson.Field
		PageObjectCoordinates        respjson.Field
		PageURL                      respjson.Field
		RawData                      respjson.Field
		TotalSecondsPlayed           respjson.Field
		ExtraFields                  map[string]respjson.Field
		raw                          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AttentionSpanEvent) RawJSON() string { return r.JSON.raw }
func (r *AttentionSpanEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AttentionSpanEventMediaType string

const (
	AttentionSpanEventMediaTypeAudio    AttentionSpanEventMediaType = "AUDIO"
	AttentionSpanEventMediaTypeDocument AttentionSpanEventMediaType = "DOCUMENT"
	AttentionSpanEventMediaTypeImage    AttentionSpanEventMediaType = "IMAGE"
	AttentionSpanEventMediaTypeOther    AttentionSpanEventMediaType = "OTHER"
	AttentionSpanEventMediaTypeVideo    AttentionSpanEventMediaType = "VIDEO"
)

// The properties MediaType, OccurredTimestamp, RawDataMap, SessionID are required.
type AttentionSpanEventRequestParam struct {
	// Any of "AUDIO", "DOCUMENT", "IMAGE", "OTHER", "VIDEO".
	MediaType         AttentionSpanEventRequestMediaType `json:"mediaType,omitzero,required"`
	OccurredTimestamp int64                              `json:"occurredTimestamp,required"`
	RawDataMap        map[string]int64                   `json:"rawDataMap,omitzero,required"`
	SessionID         string                             `json:"sessionId,required"`
	Hsenc             param.Opt[string]                  `json:"_hsenc,omitzero"`
	ContactID         param.Opt[int64]                   `json:"contactId,omitzero"`
	ContactUtk        param.Opt[string]                  `json:"contactUtk,omitzero"`
	ExternalID        param.Opt[string]                  `json:"externalId,omitzero"`
	MediaBridgeID     param.Opt[int64]                   `json:"mediaBridgeId,omitzero"`
	MediaName         param.Opt[string]                  `json:"mediaName,omitzero"`
	MediaURL          param.Opt[string]                  `json:"mediaUrl,omitzero"`
	PageID            param.Opt[int64]                   `json:"pageId,omitzero"`
	PageName          param.Opt[string]                  `json:"pageName,omitzero"`
	PageURL           param.Opt[string]                  `json:"pageUrl,omitzero"`
	RawDataString     param.Opt[string]                  `json:"rawDataString,omitzero"`
	DerivedValues     AttentionSpanCalculatedValuesParam `json:"derivedValues,omitzero"`
	paramObj
}

func (r AttentionSpanEventRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow AttentionSpanEventRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AttentionSpanEventRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AttentionSpanEventRequestMediaType string

const (
	AttentionSpanEventRequestMediaTypeAudio    AttentionSpanEventRequestMediaType = "AUDIO"
	AttentionSpanEventRequestMediaTypeDocument AttentionSpanEventRequestMediaType = "DOCUMENT"
	AttentionSpanEventRequestMediaTypeImage    AttentionSpanEventRequestMediaType = "IMAGE"
	AttentionSpanEventRequestMediaTypeOther    AttentionSpanEventRequestMediaType = "OTHER"
	AttentionSpanEventRequestMediaTypeVideo    AttentionSpanEventRequestMediaType = "VIDEO"
)

type BulkIntegratorObjectCreationResponse struct {
	CreatedObjects map[string]IntegratorObjectCreationResponse `json:"createdObjects,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreatedObjects respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BulkIntegratorObjectCreationResponse) RawJSON() string { return r.JSON.raw }
func (r *BulkIntegratorObjectCreationResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CaseChangeTestExtensionData struct {
	Mood string `json:"mood,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Mood        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CaseChangeTestExtensionData) RawJSON() string { return r.JSON.raw }
func (r *CaseChangeTestExtensionData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CollectionResponsePropertyGroupNoPaging struct {
	Results []crm.PropertyGroup `json:"results,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Results     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CollectionResponsePropertyGroupNoPaging) RawJSON() string { return r.JSON.raw }
func (r *CollectionResponsePropertyGroupNoPaging) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CollectionResponsePropertyNoPaging struct {
	Results []Property1 `json:"results,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Results     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CollectionResponsePropertyNoPaging) RawJSON() string { return r.JSON.raw }
func (r *CollectionResponsePropertyNoPaging) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DefaultRequirements struct {
	Gates []string `json:"gates,required"`
	// Any of "AND", "OR".
	Operator   DefaultRequirementsOperator `json:"operator,required"`
	ScopeNames []string                    `json:"scopeNames,required"`
	Settings   []string                    `json:"settings,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Gates       respjson.Field
		Operator    respjson.Field
		ScopeNames  respjson.Field
		Settings    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DefaultRequirements) RawJSON() string { return r.JSON.raw }
func (r *DefaultRequirements) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DefaultRequirementsOperator string

const (
	DefaultRequirementsOperatorAnd DefaultRequirementsOperator = "AND"
	DefaultRequirementsOperatorOr  DefaultRequirementsOperator = "OR"
)

type DefinitionSource struct {
	Type string `json:"type,required"`
	Name string `json:"name"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DefinitionSource) RawJSON() string { return r.JSON.raw }
func (r *DefinitionSource) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Endpoints struct {
	Discovery bool     `json:"discovery,required"`
	Schemes   []string `json:"schemes,required"`
	URL       string   `json:"url,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Discovery   respjson.Field
		Schemes     respjson.Field
		URL         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Endpoints) RawJSON() string { return r.JSON.raw }
func (r *Endpoints) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this Endpoints to a EndpointsParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// EndpointsParam.Overrides()
func (r Endpoints) ToParam() EndpointsParam {
	return param.Override[EndpointsParam](json.RawMessage(r.RawJSON()))
}

// The properties Discovery, Schemes, URL are required.
type EndpointsParam struct {
	Discovery bool     `json:"discovery,required"`
	Schemes   []string `json:"schemes,omitzero,required"`
	URL       string   `json:"url,required"`
	paramObj
}

func (r EndpointsParam) MarshalJSON() (data []byte, err error) {
	type shadow EndpointsParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EndpointsParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EventVisibilityChange struct {
	// Any of "ALL", "ATTENTION_SPAN", "MEDIA_PLAYS", "MEDIA_PLAYS_PERCENT".
	EventType       EventVisibilityChangeEventType `json:"eventType,required"`
	UpdatedAt       int64                          `json:"updatedAt,required"`
	ShowInReporting bool                           `json:"showInReporting"`
	ShowInTimeline  bool                           `json:"showInTimeline"`
	ShowInWorkflows bool                           `json:"showInWorkflows"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		EventType       respjson.Field
		UpdatedAt       respjson.Field
		ShowInReporting respjson.Field
		ShowInTimeline  respjson.Field
		ShowInWorkflows respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EventVisibilityChange) RawJSON() string { return r.JSON.raw }
func (r *EventVisibilityChange) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this EventVisibilityChange to a EventVisibilityChangeParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// EventVisibilityChangeParam.Overrides()
func (r EventVisibilityChange) ToParam() EventVisibilityChangeParam {
	return param.Override[EventVisibilityChangeParam](json.RawMessage(r.RawJSON()))
}

type EventVisibilityChangeEventType string

const (
	EventVisibilityChangeEventTypeAll               EventVisibilityChangeEventType = "ALL"
	EventVisibilityChangeEventTypeAttentionSpan     EventVisibilityChangeEventType = "ATTENTION_SPAN"
	EventVisibilityChangeEventTypeMediaPlays        EventVisibilityChangeEventType = "MEDIA_PLAYS"
	EventVisibilityChangeEventTypeMediaPlaysPercent EventVisibilityChangeEventType = "MEDIA_PLAYS_PERCENT"
)

// The properties EventType, UpdatedAt are required.
type EventVisibilityChangeParam struct {
	// Any of "ALL", "ATTENTION_SPAN", "MEDIA_PLAYS", "MEDIA_PLAYS_PERCENT".
	EventType       EventVisibilityChangeEventType `json:"eventType,omitzero,required"`
	UpdatedAt       int64                          `json:"updatedAt,required"`
	ShowInReporting param.Opt[bool]                `json:"showInReporting,omitzero"`
	ShowInTimeline  param.Opt[bool]                `json:"showInTimeline,omitzero"`
	ShowInWorkflows param.Opt[bool]                `json:"showInWorkflows,omitzero"`
	paramObj
}

func (r EventVisibilityChangeParam) MarshalJSON() (data []byte, err error) {
	type shadow EventVisibilityChangeParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EventVisibilityChangeParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EventVisibilityResponse struct {
	CreatedAt          time.Time               `json:"createdAt,required" format:"date-time"`
	VisibilitySettings []EventVisibilityChange `json:"visibilitySettings,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreatedAt          respjson.Field
		VisibilitySettings respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EventVisibilityResponse) RawJSON() string { return r.JSON.raw }
func (r *EventVisibilityResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExtensionData struct {
	ExtensionStatusMap                  map[string]string                   `json:"extensionStatusMap,required"`
	Tags                                []string                            `json:"tags,required"`
	CaseChangeTestExtensionData         CaseChangeTestExtensionData         `json:"caseChangeTestExtensionData"`
	OptionDecoratorsExtensionData       OptionDecoratorsExtensionData       `json:"optionDecoratorsExtensionData"`
	RequiredPropertiesExtensionData     RequiredPropertiesExtensionData     `json:"requiredPropertiesExtensionData"`
	SoftRequiredPropertiesExtensionData SoftRequiredPropertiesExtensionData `json:"softRequiredPropertiesExtensionData"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ExtensionStatusMap                  respjson.Field
		Tags                                respjson.Field
		CaseChangeTestExtensionData         respjson.Field
		OptionDecoratorsExtensionData       respjson.Field
		RequiredPropertiesExtensionData     respjson.Field
		SoftRequiredPropertiesExtensionData respjson.Field
		ExtraFields                         map[string]respjson.Field
		raw                                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExtensionData) RawJSON() string { return r.JSON.raw }
func (r *ExtensionData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExternalOptionsMetaData struct {
	Filter              FilteringMetaData `json:"filter"`
	RelatedObjectTypeID string            `json:"relatedObjectTypeId"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Filter              respjson.Field
		RelatedObjectTypeID respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExternalOptionsMetaData) RawJSON() string { return r.JSON.raw }
func (r *ExternalOptionsMetaData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FieldLevelPermission struct {
	AccessLevel string `json:"accessLevel,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AccessLevel respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FieldLevelPermission) RawJSON() string { return r.JSON.raw }
func (r *FieldLevelPermission) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FilteringMetaData struct {
	IncludeUnconfirmedUsers bool     `json:"includeUnconfirmedUsers,required"`
	PipelineIDs             []string `json:"pipelineIds,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		IncludeUnconfirmedUsers respjson.Field
		PipelineIDs             respjson.Field
		ExtraFields             map[string]respjson.Field
		raw                     string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FilteringMetaData) RawJSON() string { return r.JSON.raw }
func (r *FilteringMetaData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Group struct {
	Deleted          bool   `json:"deleted,required"`
	DisplayName      string `json:"displayName,required"`
	DisplayOrder     int64  `json:"displayOrder,required"`
	FulcrumPortalID  int64  `json:"fulcrumPortalId,required"`
	FulcrumTimestamp int64  `json:"fulcrumTimestamp,required"`
	HubspotDefined   bool   `json:"hubspotDefined,required"`
	Name             string `json:"name,required"`
	PortalID         int64  `json:"portalId,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Deleted          respjson.Field
		DisplayName      respjson.Field
		DisplayOrder     respjson.Field
		FulcrumPortalID  respjson.Field
		FulcrumTimestamp respjson.Field
		HubspotDefined   respjson.Field
		Name             respjson.Field
		PortalID         respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Group) RawJSON() string { return r.JSON.raw }
func (r *Group) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GroupView struct {
	DisplayName      string `json:"displayName,required"`
	DisplayOrder     int64  `json:"displayOrder,required"`
	FulcrumPortalID  int64  `json:"fulcrumPortalId,required"`
	FulcrumTimestamp int64  `json:"fulcrumTimestamp,required"`
	HubspotDefined   bool   `json:"hubspotDefined,required"`
	Name             string `json:"name,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DisplayName      respjson.Field
		DisplayOrder     respjson.Field
		FulcrumPortalID  respjson.Field
		FulcrumTimestamp respjson.Field
		HubspotDefined   respjson.Field
		Name             respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GroupView) RawJSON() string { return r.JSON.raw }
func (r *GroupView) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InboundDBObjectType struct {
	ID                          int64    `json:"id,required"`
	AllowsSensitiveProperties   bool     `json:"allowsSensitiveProperties,required"`
	CreateDatePropertyName      string   `json:"createDatePropertyName,required"`
	DefaultSearchPropertyNames  []string `json:"defaultSearchPropertyNames,required"`
	Deleted                     bool     `json:"deleted,required"`
	FullyQualifiedName          string   `json:"fullyQualifiedName,required"`
	HasCustomProperties         bool     `json:"hasCustomProperties,required"`
	HasDefaultProperties        bool     `json:"hasDefaultProperties,required"`
	HasExternalObjectIDs        bool     `json:"hasExternalObjectIds,required"`
	HasOwners                   bool     `json:"hasOwners,required"`
	HasPipelines                bool     `json:"hasPipelines,required"`
	IndexedForFiltersAndReports bool     `json:"indexedForFiltersAndReports,required"`
	LastModifiedPropertyName    string   `json:"lastModifiedPropertyName,required"`
	// Any of "CMS_HUBDB", "HUBSPOT", "HUBSPOT_EVENT", "INTEGRATION",
	// "INTEGRATION_EVENT", "PORTAL_SPECIFIC", "PORTAL_SPECIFIC_EVENT".
	MetaType                           InboundDBObjectTypeMetaType `json:"metaType,required"`
	MetaTypeID                         int64                       `json:"metaTypeId,required"`
	Name                               string                      `json:"name,required"`
	ObjectTypeID                       string                      `json:"objectTypeId,required"`
	PermissioningType                  string                      `json:"permissioningType,required"`
	PipelinePropertyName               string                      `json:"pipelinePropertyName,required"`
	PipelineStagePropertyName          string                      `json:"pipelineStagePropertyName,required"`
	RequiredProperties                 []string                    `json:"requiredProperties,required"`
	Restorable                         bool                        `json:"restorable,required"`
	ScopeMappings                      []ScopeMapping              `json:"scopeMappings,required"`
	SecondaryDisplayLabelPropertyNames []string                    `json:"secondaryDisplayLabelPropertyNames,required"`
	AccessScopeName                    string                      `json:"accessScopeName"`
	CreatedAt                          int64                       `json:"createdAt"`
	Description                        string                      `json:"description"`
	IntegrationAppID                   int64                       `json:"integrationAppId"`
	JanusGroup                         string                      `json:"janusGroup"`
	OwnerPortalID                      int64                       `json:"ownerPortalId"`
	PipelineCloseDatePropertyName      string                      `json:"pipelineCloseDatePropertyName"`
	PipelineTimeToClosePropertyName    string                      `json:"pipelineTimeToClosePropertyName"`
	PluralForm                         string                      `json:"pluralForm"`
	PrimaryDisplayLabelPropertyName    string                      `json:"primaryDisplayLabelPropertyName"`
	ReadScopeName                      string                      `json:"readScopeName"`
	SingularForm                       string                      `json:"singularForm"`
	Status                             string                      `json:"status"`
	Visibility                         string                      `json:"visibility"`
	WriteScopeName                     string                      `json:"writeScopeName"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                                 respjson.Field
		AllowsSensitiveProperties          respjson.Field
		CreateDatePropertyName             respjson.Field
		DefaultSearchPropertyNames         respjson.Field
		Deleted                            respjson.Field
		FullyQualifiedName                 respjson.Field
		HasCustomProperties                respjson.Field
		HasDefaultProperties               respjson.Field
		HasExternalObjectIDs               respjson.Field
		HasOwners                          respjson.Field
		HasPipelines                       respjson.Field
		IndexedForFiltersAndReports        respjson.Field
		LastModifiedPropertyName           respjson.Field
		MetaType                           respjson.Field
		MetaTypeID                         respjson.Field
		Name                               respjson.Field
		ObjectTypeID                       respjson.Field
		PermissioningType                  respjson.Field
		PipelinePropertyName               respjson.Field
		PipelineStagePropertyName          respjson.Field
		RequiredProperties                 respjson.Field
		Restorable                         respjson.Field
		ScopeMappings                      respjson.Field
		SecondaryDisplayLabelPropertyNames respjson.Field
		AccessScopeName                    respjson.Field
		CreatedAt                          respjson.Field
		Description                        respjson.Field
		IntegrationAppID                   respjson.Field
		JanusGroup                         respjson.Field
		OwnerPortalID                      respjson.Field
		PipelineCloseDatePropertyName      respjson.Field
		PipelineTimeToClosePropertyName    respjson.Field
		PluralForm                         respjson.Field
		PrimaryDisplayLabelPropertyName    respjson.Field
		ReadScopeName                      respjson.Field
		SingularForm                       respjson.Field
		Status                             respjson.Field
		Visibility                         respjson.Field
		WriteScopeName                     respjson.Field
		ExtraFields                        map[string]respjson.Field
		raw                                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InboundDBObjectType) RawJSON() string { return r.JSON.raw }
func (r *InboundDBObjectType) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InboundDBObjectTypeMetaType string

const (
	InboundDBObjectTypeMetaTypeCmsHubdb            InboundDBObjectTypeMetaType = "CMS_HUBDB"
	InboundDBObjectTypeMetaTypeHubspot             InboundDBObjectTypeMetaType = "HUBSPOT"
	InboundDBObjectTypeMetaTypeHubspotEvent        InboundDBObjectTypeMetaType = "HUBSPOT_EVENT"
	InboundDBObjectTypeMetaTypeIntegration         InboundDBObjectTypeMetaType = "INTEGRATION"
	InboundDBObjectTypeMetaTypeIntegrationEvent    InboundDBObjectTypeMetaType = "INTEGRATION_EVENT"
	InboundDBObjectTypeMetaTypePortalSpecific      InboundDBObjectTypeMetaType = "PORTAL_SPECIFIC"
	InboundDBObjectTypeMetaTypePortalSpecificEvent InboundDBObjectTypeMetaType = "PORTAL_SPECIFIC_EVENT"
)

type IntegratorOEmbedDomainModel struct {
	ID        int64     `json:"id,required"`
	AppID     int64     `json:"appId,required"`
	CreatedAt int64     `json:"createdAt,required"`
	DeletedAt int64     `json:"deletedAt,required"`
	Endpoints Endpoints `json:"endpoints,required"`
	PortalID  int64     `json:"portalId,required"`
	UpdatedAt int64     `json:"updatedAt,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		AppID       respjson.Field
		CreatedAt   respjson.Field
		DeletedAt   respjson.Field
		Endpoints   respjson.Field
		PortalID    respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IntegratorOEmbedDomainModel) RawJSON() string { return r.JSON.raw }
func (r *IntegratorOEmbedDomainModel) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Endpoints is required.
type IntegratorOEmbedDomainRequestParam struct {
	Endpoints EndpointsParam   `json:"endpoints,omitzero,required"`
	PortalID  param.Opt[int64] `json:"portalId,omitzero"`
	paramObj
}

func (r IntegratorOEmbedDomainRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow IntegratorOEmbedDomainRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *IntegratorOEmbedDomainRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property MediaTypes is required.
type IntegratorObjectCreationRequestParam struct {
	// Any of "VIDEO", "AUDIO", "DOCUMENT", "OTHER", "IMAGE".
	MediaTypes []string `json:"mediaTypes,omitzero,required"`
	paramObj
}

func (r IntegratorObjectCreationRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow IntegratorObjectCreationRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *IntegratorObjectCreationRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IntegratorObjectCreationResponse struct {
	ObjectType     InboundDBObjectType  `json:"objectType,required"`
	Properties     []PropertyDefinition `json:"properties,required"`
	PropertyGroups []Group              `json:"propertyGroups,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ObjectType     respjson.Field
		Properties     respjson.Field
		PropertyGroups respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IntegratorObjectCreationResponse) RawJSON() string { return r.JSON.raw }
func (r *IntegratorObjectCreationResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgePropertyUpdateParam struct {
	CalculationFormula param.Opt[string] `json:"calculationFormula,omitzero"`
	Description        param.Opt[string] `json:"description,omitzero"`
	DisplayOrder       param.Opt[int64]  `json:"displayOrder,omitzero"`
	FormField          param.Opt[bool]   `json:"formField,omitzero"`
	GroupName          param.Opt[string] `json:"groupName,omitzero"`
	HasUniqueValue     param.Opt[bool]   `json:"hasUniqueValue,omitzero"`
	Hidden             param.Opt[bool]   `json:"hidden,omitzero"`
	Label              param.Opt[string] `json:"label,omitzero"`
	// Any of "booleancheckbox", "calculation_equation", "checkbox", "date", "file",
	// "html", "number", "phonenumber", "radio", "select", "text", "textarea".
	FieldType MediaBridgePropertyUpdateFieldType `json:"fieldType,omitzero"`
	Options   []shared.OptionInputParam          `json:"options,omitzero"`
	// Any of "bool", "date", "datetime", "enumeration", "number", "phone_number",
	// "string".
	Type MediaBridgePropertyUpdateType `json:"type,omitzero"`
	paramObj
}

func (r MediaBridgePropertyUpdateParam) MarshalJSON() (data []byte, err error) {
	type shadow MediaBridgePropertyUpdateParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MediaBridgePropertyUpdateParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgePropertyUpdateFieldType string

const (
	MediaBridgePropertyUpdateFieldTypeBooleancheckbox     MediaBridgePropertyUpdateFieldType = "booleancheckbox"
	MediaBridgePropertyUpdateFieldTypeCalculationEquation MediaBridgePropertyUpdateFieldType = "calculation_equation"
	MediaBridgePropertyUpdateFieldTypeCheckbox            MediaBridgePropertyUpdateFieldType = "checkbox"
	MediaBridgePropertyUpdateFieldTypeDate                MediaBridgePropertyUpdateFieldType = "date"
	MediaBridgePropertyUpdateFieldTypeFile                MediaBridgePropertyUpdateFieldType = "file"
	MediaBridgePropertyUpdateFieldTypeHTML                MediaBridgePropertyUpdateFieldType = "html"
	MediaBridgePropertyUpdateFieldTypeNumber              MediaBridgePropertyUpdateFieldType = "number"
	MediaBridgePropertyUpdateFieldTypePhonenumber         MediaBridgePropertyUpdateFieldType = "phonenumber"
	MediaBridgePropertyUpdateFieldTypeRadio               MediaBridgePropertyUpdateFieldType = "radio"
	MediaBridgePropertyUpdateFieldTypeSelect              MediaBridgePropertyUpdateFieldType = "select"
	MediaBridgePropertyUpdateFieldTypeText                MediaBridgePropertyUpdateFieldType = "text"
	MediaBridgePropertyUpdateFieldTypeTextarea            MediaBridgePropertyUpdateFieldType = "textarea"
)

type MediaBridgePropertyUpdateType string

const (
	MediaBridgePropertyUpdateTypeBool        MediaBridgePropertyUpdateType = "bool"
	MediaBridgePropertyUpdateTypeDate        MediaBridgePropertyUpdateType = "date"
	MediaBridgePropertyUpdateTypeDatetime    MediaBridgePropertyUpdateType = "datetime"
	MediaBridgePropertyUpdateTypeEnumeration MediaBridgePropertyUpdateType = "enumeration"
	MediaBridgePropertyUpdateTypeNumber      MediaBridgePropertyUpdateType = "number"
	MediaBridgePropertyUpdateTypePhoneNumber MediaBridgePropertyUpdateType = "phone_number"
	MediaBridgePropertyUpdateTypeString      MediaBridgePropertyUpdateType = "string"
)

// The property UpdatedAt is required.
type MediaBridgeProviderPartialParam struct {
	UpdatedAt int64             `json:"updatedAt,required"`
	Name      param.Opt[string] `json:"name,omitzero"`
	paramObj
}

func (r MediaBridgeProviderPartialParam) MarshalJSON() (data []byte, err error) {
	type shadow MediaBridgeProviderPartialParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MediaBridgeProviderPartialParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeProviderRegistrationResponse struct {
	AppID int64  `json:"appId,required"`
	Name  string `json:"name,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AppID       respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MediaBridgeProviderRegistrationResponse) RawJSON() string { return r.JSON.raw }
func (r *MediaBridgeProviderRegistrationResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaPlayedEvent struct {
	ContactID                    int64  `json:"contactId,required"`
	MediaBridgeID                int64  `json:"mediaBridgeId,required"`
	MediaBridgeObjectCoordinates string `json:"mediaBridgeObjectCoordinates,required"`
	MediaBridgeObjectTypeID      string `json:"mediaBridgeObjectTypeId,required"`
	MediaName                    string `json:"mediaName,required"`
	// Any of "AUDIO", "DOCUMENT", "IMAGE", "OTHER", "VIDEO".
	MediaType         MediaPlayedEventMediaType `json:"mediaType,required"`
	OccurredTimestamp int64                     `json:"occurredTimestamp,required"`
	PortalID          int64                     `json:"portalId,required"`
	ProviderID        int64                     `json:"providerId,required"`
	SessionID         string                    `json:"sessionId,required"`
	// Any of "STARTED", "VIEWED".
	State                 MediaPlayedEventState `json:"state,required"`
	IframeURL             string                `json:"iframeUrl"`
	MediaURL              string                `json:"mediaUrl"`
	PageID                int64                 `json:"pageId"`
	PageName              string                `json:"pageName"`
	PageObjectCoordinates string                `json:"pageObjectCoordinates"`
	PageURL               string                `json:"pageUrl"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ContactID                    respjson.Field
		MediaBridgeID                respjson.Field
		MediaBridgeObjectCoordinates respjson.Field
		MediaBridgeObjectTypeID      respjson.Field
		MediaName                    respjson.Field
		MediaType                    respjson.Field
		OccurredTimestamp            respjson.Field
		PortalID                     respjson.Field
		ProviderID                   respjson.Field
		SessionID                    respjson.Field
		State                        respjson.Field
		IframeURL                    respjson.Field
		MediaURL                     respjson.Field
		PageID                       respjson.Field
		PageName                     respjson.Field
		PageObjectCoordinates        respjson.Field
		PageURL                      respjson.Field
		ExtraFields                  map[string]respjson.Field
		raw                          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MediaPlayedEvent) RawJSON() string { return r.JSON.raw }
func (r *MediaPlayedEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaPlayedEventMediaType string

const (
	MediaPlayedEventMediaTypeAudio    MediaPlayedEventMediaType = "AUDIO"
	MediaPlayedEventMediaTypeDocument MediaPlayedEventMediaType = "DOCUMENT"
	MediaPlayedEventMediaTypeImage    MediaPlayedEventMediaType = "IMAGE"
	MediaPlayedEventMediaTypeOther    MediaPlayedEventMediaType = "OTHER"
	MediaPlayedEventMediaTypeVideo    MediaPlayedEventMediaType = "VIDEO"
)

type MediaPlayedEventState string

const (
	MediaPlayedEventStateStarted MediaPlayedEventState = "STARTED"
	MediaPlayedEventStateViewed  MediaPlayedEventState = "VIEWED"
)

// The properties MediaType, OccurredTimestamp, SessionID, State are required.
type MediaPlayedEventRequestParam struct {
	// Any of "AUDIO", "DOCUMENT", "IMAGE", "OTHER", "VIDEO".
	MediaType         MediaPlayedEventRequestMediaType `json:"mediaType,omitzero,required"`
	OccurredTimestamp int64                            `json:"occurredTimestamp,required"`
	SessionID         string                           `json:"sessionId,required"`
	// Any of "STARTED", "VIEWED".
	State         MediaPlayedEventRequestState `json:"state,omitzero,required"`
	Hsenc         param.Opt[string]            `json:"_hsenc,omitzero"`
	ContactID     param.Opt[int64]             `json:"contactId,omitzero"`
	ContactUtk    param.Opt[string]            `json:"contactUtk,omitzero"`
	ExternalID    param.Opt[string]            `json:"externalId,omitzero"`
	IframeURL     param.Opt[string]            `json:"iframeUrl,omitzero"`
	MediaBridgeID param.Opt[int64]             `json:"mediaBridgeId,omitzero"`
	MediaName     param.Opt[string]            `json:"mediaName,omitzero"`
	MediaURL      param.Opt[string]            `json:"mediaUrl,omitzero"`
	PageID        param.Opt[int64]             `json:"pageId,omitzero"`
	PageName      param.Opt[string]            `json:"pageName,omitzero"`
	PageURL       param.Opt[string]            `json:"pageUrl,omitzero"`
	paramObj
}

func (r MediaPlayedEventRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow MediaPlayedEventRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MediaPlayedEventRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaPlayedEventRequestMediaType string

const (
	MediaPlayedEventRequestMediaTypeAudio    MediaPlayedEventRequestMediaType = "AUDIO"
	MediaPlayedEventRequestMediaTypeDocument MediaPlayedEventRequestMediaType = "DOCUMENT"
	MediaPlayedEventRequestMediaTypeImage    MediaPlayedEventRequestMediaType = "IMAGE"
	MediaPlayedEventRequestMediaTypeOther    MediaPlayedEventRequestMediaType = "OTHER"
	MediaPlayedEventRequestMediaTypeVideo    MediaPlayedEventRequestMediaType = "VIDEO"
)

type MediaPlayedEventRequestState string

const (
	MediaPlayedEventRequestStateStarted MediaPlayedEventRequestState = "STARTED"
	MediaPlayedEventRequestStateViewed  MediaPlayedEventRequestState = "VIEWED"
)

type MediaPlayedPercentageEvent struct {
	// The ID of the contact in HubSpot’s system that consumed the media. This can be
	// fetched using HubSpot's Get contact by usertoken (utk) API. The API also
	// supports supplying a usertoken, and will handle converting this into a contact
	// ID automatically.
	ContactID                    int64  `json:"contactId,required"`
	MediaBridgeID                int64  `json:"mediaBridgeId,required"`
	MediaBridgeObjectCoordinates string `json:"mediaBridgeObjectCoordinates,required"`
	MediaBridgeObjectTypeID      string `json:"mediaBridgeObjectTypeId,required"`
	MediaName                    string `json:"mediaName,required"`
	// Any of "AUDIO", "DOCUMENT", "IMAGE", "OTHER", "VIDEO".
	MediaType         MediaPlayedPercentageEventMediaType `json:"mediaType,required"`
	OccurredTimestamp int64                               `json:"occurredTimestamp,required"`
	PlayedPercent     int64                               `json:"playedPercent,required"`
	// The ID of the HubSpot account.
	PortalID   int64  `json:"portalId,required"`
	ProviderID int64  `json:"providerId,required"`
	SessionID  string `json:"sessionId,required"`
	MediaURL   string `json:"mediaUrl"`
	// The content ID of the page that an event happened on, for HubSpot pages.
	// Required if the page is a HubSpot page.
	PageID int64 `json:"pageId"`
	// The name or title of the page that an event happened on. Required for
	// non-HubSpot pages.
	PageName              string `json:"pageName"`
	PageObjectCoordinates string `json:"pageObjectCoordinates"`
	// The URL of the page that an event happened on. Required for non-HubSpot pages.
	PageURL string `json:"pageUrl"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ContactID                    respjson.Field
		MediaBridgeID                respjson.Field
		MediaBridgeObjectCoordinates respjson.Field
		MediaBridgeObjectTypeID      respjson.Field
		MediaName                    respjson.Field
		MediaType                    respjson.Field
		OccurredTimestamp            respjson.Field
		PlayedPercent                respjson.Field
		PortalID                     respjson.Field
		ProviderID                   respjson.Field
		SessionID                    respjson.Field
		MediaURL                     respjson.Field
		PageID                       respjson.Field
		PageName                     respjson.Field
		PageObjectCoordinates        respjson.Field
		PageURL                      respjson.Field
		ExtraFields                  map[string]respjson.Field
		raw                          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MediaPlayedPercentageEvent) RawJSON() string { return r.JSON.raw }
func (r *MediaPlayedPercentageEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaPlayedPercentageEventMediaType string

const (
	MediaPlayedPercentageEventMediaTypeAudio    MediaPlayedPercentageEventMediaType = "AUDIO"
	MediaPlayedPercentageEventMediaTypeDocument MediaPlayedPercentageEventMediaType = "DOCUMENT"
	MediaPlayedPercentageEventMediaTypeImage    MediaPlayedPercentageEventMediaType = "IMAGE"
	MediaPlayedPercentageEventMediaTypeOther    MediaPlayedPercentageEventMediaType = "OTHER"
	MediaPlayedPercentageEventMediaTypeVideo    MediaPlayedPercentageEventMediaType = "VIDEO"
)

// The properties MediaType, OccurredTimestamp, PlayedPercent, SessionID are
// required.
type MediaPlayedPercentageEventRequestParam struct {
	// Any of "AUDIO", "DOCUMENT", "IMAGE", "OTHER", "VIDEO".
	MediaType         MediaPlayedPercentageEventRequestMediaType `json:"mediaType,omitzero,required"`
	OccurredTimestamp int64                                      `json:"occurredTimestamp,required"`
	PlayedPercent     int64                                      `json:"playedPercent,required"`
	SessionID         string                                     `json:"sessionId,required"`
	Hsenc             param.Opt[string]                          `json:"_hsenc,omitzero"`
	ContactID         param.Opt[int64]                           `json:"contactId,omitzero"`
	ContactUtk        param.Opt[string]                          `json:"contactUtk,omitzero"`
	ExternalID        param.Opt[string]                          `json:"externalId,omitzero"`
	MediaBridgeID     param.Opt[int64]                           `json:"mediaBridgeId,omitzero"`
	MediaName         param.Opt[string]                          `json:"mediaName,omitzero"`
	MediaURL          param.Opt[string]                          `json:"mediaUrl,omitzero"`
	PageID            param.Opt[int64]                           `json:"pageId,omitzero"`
	PageName          param.Opt[string]                          `json:"pageName,omitzero"`
	PageURL           param.Opt[string]                          `json:"pageUrl,omitzero"`
	paramObj
}

func (r MediaPlayedPercentageEventRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow MediaPlayedPercentageEventRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MediaPlayedPercentageEventRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaPlayedPercentageEventRequestMediaType string

const (
	MediaPlayedPercentageEventRequestMediaTypeAudio    MediaPlayedPercentageEventRequestMediaType = "AUDIO"
	MediaPlayedPercentageEventRequestMediaTypeDocument MediaPlayedPercentageEventRequestMediaType = "DOCUMENT"
	MediaPlayedPercentageEventRequestMediaTypeImage    MediaPlayedPercentageEventRequestMediaType = "IMAGE"
	MediaPlayedPercentageEventRequestMediaTypeOther    MediaPlayedPercentageEventRequestMediaType = "OTHER"
	MediaPlayedPercentageEventRequestMediaTypeVideo    MediaPlayedPercentageEventRequestMediaType = "VIDEO"
)

type OEmbedDomainsCollectionResponse struct {
	Results    []IntegratorOEmbedDomainModel `json:"results,required"`
	TotalCount int64                         `json:"totalCount"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Results     respjson.Field
		TotalCount  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OEmbedDomainsCollectionResponse) RawJSON() string { return r.JSON.raw }
func (r *OEmbedDomainsCollectionResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ObjectDefinitionResponse struct {
	ObjectTypeID   string               `json:"objectTypeId,required"`
	ObjectTypeName string               `json:"objectTypeName,required"`
	Properties     []PropertyDefinition `json:"properties,required"`
	PropertyGroups []GroupView          `json:"propertyGroups,required"`
	Schema         InboundDBObjectType  `json:"schema"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ObjectTypeID   respjson.Field
		ObjectTypeName respjson.Field
		Properties     respjson.Field
		PropertyGroups respjson.Field
		Schema         respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ObjectDefinitionResponse) RawJSON() string { return r.JSON.raw }
func (r *ObjectDefinitionResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Option1 struct {
	Hidden       bool   `json:"hidden,required"`
	Label        string `json:"label,required"`
	Value        string `json:"value,required"`
	Description  string `json:"description"`
	DisplayOrder int64  `json:"displayOrder"`
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
func (r Option1) RawJSON() string { return r.JSON.raw }
func (r *Option1) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OptionDecorations struct {
	Color string `json:"color,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Color       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OptionDecorations) RawJSON() string { return r.JSON.raw }
func (r *OptionDecorations) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OptionDecoratorsExtensionData struct {
	OptionDecorators     map[string]OptionDecorations `json:"optionDecorators,required"`
	OptionDecoratorStyle string                       `json:"optionDecoratorStyle,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		OptionDecorators     respjson.Field
		OptionDecoratorStyle respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OptionDecoratorsExtensionData) RawJSON() string { return r.JSON.raw }
func (r *OptionDecoratorsExtensionData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Property1 struct {
	Description        string    `json:"description,required"`
	FieldType          string    `json:"fieldType,required"`
	GroupName          string    `json:"groupName,required"`
	Label              string    `json:"label,required"`
	Name               string    `json:"name,required"`
	Options            []Option1 `json:"options,required"`
	Type               string    `json:"type,required"`
	Archived           bool      `json:"archived"`
	ArchivedAt         time.Time `json:"archivedAt" format:"date-time"`
	Calculated         bool      `json:"calculated"`
	CalculationFormula string    `json:"calculationFormula"`
	CreatedAt          time.Time `json:"createdAt" format:"date-time"`
	CreatedUserID      string    `json:"createdUserId"`
	// Any of "highly_sensitive", "non_sensitive", "sensitive".
	DataSensitivity Property1DataSensitivity `json:"dataSensitivity"`
	// Any of "absolute", "absolute_with_relative", "time_since", "time_until".
	DateDisplayHint         Property1DateDisplayHint            `json:"dateDisplayHint"`
	DisplayOrder            int64                               `json:"displayOrder"`
	ExternalOptions         bool                                `json:"externalOptions"`
	FormField               bool                                `json:"formField"`
	HasUniqueValue          bool                                `json:"hasUniqueValue"`
	Hidden                  bool                                `json:"hidden"`
	HubspotDefined          bool                                `json:"hubspotDefined"`
	ModificationMetadata    shared.PropertyModificationMetadata `json:"modificationMetadata"`
	ReferencedObjectType    string                              `json:"referencedObjectType"`
	SensitiveDataCategories []string                            `json:"sensitiveDataCategories"`
	ShowCurrencySymbol      bool                                `json:"showCurrencySymbol"`
	UpdatedAt               time.Time                           `json:"updatedAt" format:"date-time"`
	UpdatedUserID           string                              `json:"updatedUserId"`
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
func (r Property1) RawJSON() string { return r.JSON.raw }
func (r *Property1) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Property1DataSensitivity string

const (
	Property1DataSensitivityHighlySensitive Property1DataSensitivity = "highly_sensitive"
	Property1DataSensitivityNonSensitive    Property1DataSensitivity = "non_sensitive"
	Property1DataSensitivitySensitive       Property1DataSensitivity = "sensitive"
)

type Property1DateDisplayHint string

const (
	Property1DateDisplayHintAbsolute             Property1DateDisplayHint = "absolute"
	Property1DateDisplayHintAbsoluteWithRelative Property1DateDisplayHint = "absolute_with_relative"
	Property1DateDisplayHintTimeSince            Property1DateDisplayHint = "time_since"
	Property1DateDisplayHintTimeUntil            Property1DateDisplayHint = "time_until"
)

type PropertyDefinition struct {
	ObjectTypeID string `json:"objectTypeId,required"`
	// Defines a property
	Property                 shared.Property          `json:"property,required"`
	CalculationExpression    map[string]any           `json:"calculationExpression"`
	CalculationFormula       string                   `json:"calculationFormula"`
	DefinitionSource         PropertyDefinitionSource `json:"definitionSource"`
	ExtensionData            ExtensionData            `json:"extensionData"`
	ExternalOptionsMetaData  ExternalOptionsMetaData  `json:"externalOptionsMetaData"`
	FulcrumPortalID          int64                    `json:"fulcrumPortalId"`
	FulcrumTimestamp         int64                    `json:"fulcrumTimestamp"`
	JanusGroup               string                   `json:"janusGroup"`
	Permission               FieldLevelPermission     `json:"permission"`
	PropertyDefinitionSource DefinitionSource         `json:"propertyDefinitionSource"`
	PropertyRequirements     DefaultRequirements      `json:"propertyRequirements"`
	RollupExpression         RollupExpression         `json:"rollupExpression"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ObjectTypeID             respjson.Field
		Property                 respjson.Field
		CalculationExpression    respjson.Field
		CalculationFormula       respjson.Field
		DefinitionSource         respjson.Field
		ExtensionData            respjson.Field
		ExternalOptionsMetaData  respjson.Field
		FulcrumPortalID          respjson.Field
		FulcrumTimestamp         respjson.Field
		JanusGroup               respjson.Field
		Permission               respjson.Field
		PropertyDefinitionSource respjson.Field
		PropertyRequirements     respjson.Field
		RollupExpression         respjson.Field
		ExtraFields              map[string]respjson.Field
		raw                      string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PropertyDefinition) RawJSON() string { return r.JSON.raw }
func (r *PropertyDefinition) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PropertyDefinitionSource struct {
	// Any of "GLOBAL", "HAVEN_BRANCH", "OBJECT_TYPE", "PORTAL".
	Type PropertyDefinitionSourceType `json:"type,required"`
	Name string                       `json:"name"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PropertyDefinitionSource) RawJSON() string { return r.JSON.raw }
func (r *PropertyDefinitionSource) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PropertyDefinitionSourceType string

const (
	PropertyDefinitionSourceTypeGlobal      PropertyDefinitionSourceType = "GLOBAL"
	PropertyDefinitionSourceTypeHavenBranch PropertyDefinitionSourceType = "HAVEN_BRANCH"
	PropertyDefinitionSourceTypeObjectType  PropertyDefinitionSourceType = "OBJECT_TYPE"
	PropertyDefinitionSourceTypePortal      PropertyDefinitionSourceType = "PORTAL"
)

type RequiredPropertiesExtensionData struct {
	IsRequiredProperty bool `json:"isRequiredProperty,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		IsRequiredProperty respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RequiredPropertiesExtensionData) RawJSON() string { return r.JSON.raw }
func (r *RequiredPropertiesExtensionData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RollupExpression struct {
	AssociationTypes            []shared.AssociationSpec `json:"associationTypes,required"`
	RollupOperator              string                   `json:"rollupOperator,required"`
	SourceObjectTypeID          string                   `json:"sourceObjectTypeId,required"`
	SourcePropertyName          string                   `json:"sourcePropertyName,required"`
	ConditionalExpression       map[string]any           `json:"conditionalExpression"`
	ConditionalFormula          string                   `json:"conditionalFormula"`
	EmptyRollupValue            string                   `json:"emptyRollupValue"`
	SourceCompareByPropertyName string                   `json:"sourceCompareByPropertyName"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AssociationTypes            respjson.Field
		RollupOperator              respjson.Field
		SourceObjectTypeID          respjson.Field
		SourcePropertyName          respjson.Field
		ConditionalExpression       respjson.Field
		ConditionalFormula          respjson.Field
		EmptyRollupValue            respjson.Field
		SourceCompareByPropertyName respjson.Field
		ExtraFields                 map[string]respjson.Field
		raw                         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RollupExpression) RawJSON() string { return r.JSON.raw }
func (r *RollupExpression) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ScopeMapping struct {
	AccessLevel   string `json:"accessLevel,required"`
	RequestAction string `json:"requestAction,required"`
	ScopeName     string `json:"scopeName,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AccessLevel   respjson.Field
		RequestAction respjson.Field
		ScopeName     respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ScopeMapping) RawJSON() string { return r.JSON.raw }
func (r *ScopeMapping) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SoftRequiredPropertiesExtensionData struct {
	IsSoftRequiredProperty bool `json:"isSoftRequiredProperty,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		IsSoftRequiredProperty respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SoftRequiredPropertiesExtensionData) RawJSON() string { return r.JSON.raw }
func (r *SoftRequiredPropertiesExtensionData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
