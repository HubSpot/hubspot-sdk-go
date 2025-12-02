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

type AbsoluteValue struct {
	// Any of "ABSOLUTE_VALUE".
	Operator     AbsoluteValueOperator `json:"operator,required"`
	Inputs       []ExpressionUnion     `json:"inputs"`
	PropertyName string                `json:"propertyName"`
	Value        float64               `json:"value"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Operator     respjson.Field
		Inputs       respjson.Field
		PropertyName respjson.Field
		Value        respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AbsoluteValue) RawJSON() string { return r.JSON.raw }
func (r *AbsoluteValue) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AbsoluteValueOperator string

const (
	AbsoluteValueOperatorAbsoluteValue AbsoluteValueOperator = "ABSOLUTE_VALUE"
)

type AddNumbers struct {
	EnclosedInParentheses bool `json:"enclosedInParentheses,required"`
	// Any of "ADD_NUMBERS".
	Operator     AddNumbersOperator `json:"operator,required"`
	Inputs       []ExpressionUnion  `json:"inputs"`
	PropertyName string             `json:"propertyName"`
	Value        float64            `json:"value"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		EnclosedInParentheses respjson.Field
		Operator              respjson.Field
		Inputs                respjson.Field
		PropertyName          respjson.Field
		Value                 respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AddNumbers) RawJSON() string { return r.JSON.raw }
func (r *AddNumbers) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AddNumbersOperator string

const (
	AddNumbersOperatorAddNumbers AddNumbersOperator = "ADD_NUMBERS"
)

type AddTime struct {
	// Any of "ADD_TIME".
	Operator      AddTimeOperator   `json:"operator,required"`
	StringToCheck ExpressionUnion   `json:"stringToCheck,required"`
	Inputs        []ExpressionUnion `json:"inputs"`
	PropertyName  string            `json:"propertyName"`
	Value         float64           `json:"value"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Operator      respjson.Field
		StringToCheck respjson.Field
		Inputs        respjson.Field
		PropertyName  respjson.Field
		Value         respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AddTime) RawJSON() string { return r.JSON.raw }
func (r *AddTime) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AddTimeOperator string

const (
	AddTimeOperatorAddTime AddTimeOperator = "ADD_TIME"
)

type And struct {
	EnclosedInParentheses bool `json:"enclosedInParentheses,required"`
	// Any of "AND".
	Operator     AndOperator       `json:"operator,required"`
	Inputs       []ExpressionUnion `json:"inputs"`
	PropertyName string            `json:"propertyName"`
	Value        bool              `json:"value"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		EnclosedInParentheses respjson.Field
		Operator              respjson.Field
		Inputs                respjson.Field
		PropertyName          respjson.Field
		Value                 respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r And) RawJSON() string { return r.JSON.raw }
func (r *And) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AndOperator string

const (
	AndOperatorAnd AndOperator = "AND"
)

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

type BeginsWith struct {
	// Any of "BEGINS_WITH".
	Operator      BeginsWithOperator `json:"operator,required"`
	StringToCheck ExpressionUnion    `json:"stringToCheck,required"`
	Inputs        []ExpressionUnion  `json:"inputs"`
	PropertyName  string             `json:"propertyName"`
	Value         bool               `json:"value"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Operator      respjson.Field
		StringToCheck respjson.Field
		Inputs        respjson.Field
		PropertyName  respjson.Field
		Value         respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BeginsWith) RawJSON() string { return r.JSON.raw }
func (r *BeginsWith) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BeginsWithOperator string

const (
	BeginsWithOperatorBeginsWith BeginsWithOperator = "BEGINS_WITH"
)

type BooleanPropertyVariable struct {
	// Any of "BOOLEAN_PROPERTY_VARIABLE".
	Operator     BooleanPropertyVariableOperator `json:"operator,required"`
	Inputs       []ExpressionUnion               `json:"inputs"`
	PropertyName string                          `json:"propertyName"`
	Value        bool                            `json:"value"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Operator     respjson.Field
		Inputs       respjson.Field
		PropertyName respjson.Field
		Value        respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BooleanPropertyVariable) RawJSON() string { return r.JSON.raw }
func (r *BooleanPropertyVariable) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BooleanPropertyVariableOperator string

const (
	BooleanPropertyVariableOperatorBooleanPropertyVariable BooleanPropertyVariableOperator = "BOOLEAN_PROPERTY_VARIABLE"
)

type BooleanTargetPropertyVariable struct {
	// Any of "BOOLEAN_TARGET_PROPERTY_VARIABLE".
	Operator     BooleanTargetPropertyVariableOperator `json:"operator,required"`
	Inputs       []ExpressionUnion                     `json:"inputs"`
	PropertyName string                                `json:"propertyName"`
	Value        bool                                  `json:"value"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Operator     respjson.Field
		Inputs       respjson.Field
		PropertyName respjson.Field
		Value        respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BooleanTargetPropertyVariable) RawJSON() string { return r.JSON.raw }
func (r *BooleanTargetPropertyVariable) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BooleanTargetPropertyVariableOperator string

const (
	BooleanTargetPropertyVariableOperatorBooleanTargetPropertyVariable BooleanTargetPropertyVariableOperator = "BOOLEAN_TARGET_PROPERTY_VARIABLE"
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

type ConcatStrings struct {
	// Any of "CONCAT_STRINGS".
	Operator     ConcatStringsOperator `json:"operator,required"`
	Inputs       []ExpressionUnion     `json:"inputs"`
	PropertyName string                `json:"propertyName"`
	Value        string                `json:"value"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Operator     respjson.Field
		Inputs       respjson.Field
		PropertyName respjson.Field
		Value        respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConcatStrings) RawJSON() string { return r.JSON.raw }
func (r *ConcatStrings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConcatStringsOperator string

const (
	ConcatStringsOperatorConcatStrings ConcatStringsOperator = "CONCAT_STRINGS"
)

type ConstantBoolean struct {
	// Any of "CONSTANT_BOOLEAN".
	Operator     ConstantBooleanOperator `json:"operator,required"`
	Inputs       []ExpressionUnion       `json:"inputs"`
	PropertyName string                  `json:"propertyName"`
	Value        bool                    `json:"value"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Operator     respjson.Field
		Inputs       respjson.Field
		PropertyName respjson.Field
		Value        respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConstantBoolean) RawJSON() string { return r.JSON.raw }
func (r *ConstantBoolean) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConstantBooleanOperator string

const (
	ConstantBooleanOperatorConstantBoolean ConstantBooleanOperator = "CONSTANT_BOOLEAN"
)

type ConstantNumber struct {
	// Any of "CONSTANT_NUMBER".
	Operator     ConstantNumberOperator `json:"operator,required"`
	Inputs       []ExpressionUnion      `json:"inputs"`
	PropertyName string                 `json:"propertyName"`
	Value        float64                `json:"value"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Operator     respjson.Field
		Inputs       respjson.Field
		PropertyName respjson.Field
		Value        respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConstantNumber) RawJSON() string { return r.JSON.raw }
func (r *ConstantNumber) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConstantNumberOperator string

const (
	ConstantNumberOperatorConstantNumber ConstantNumberOperator = "CONSTANT_NUMBER"
)

type ConstantString struct {
	// Any of "CONSTANT_STRING".
	Operator     ConstantStringOperator `json:"operator,required"`
	Inputs       []ExpressionUnion      `json:"inputs"`
	PropertyName string                 `json:"propertyName"`
	Value        string                 `json:"value"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Operator     respjson.Field
		Inputs       respjson.Field
		PropertyName respjson.Field
		Value        respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConstantString) RawJSON() string { return r.JSON.raw }
func (r *ConstantString) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConstantStringOperator string

const (
	ConstantStringOperatorConstantString ConstantStringOperator = "CONSTANT_STRING"
)

type Contains struct {
	// Any of "CONTAINS".
	Operator      ContainsOperator  `json:"operator,required"`
	StringToCheck ExpressionUnion   `json:"stringToCheck,required"`
	Inputs        []ExpressionUnion `json:"inputs"`
	PropertyName  string            `json:"propertyName"`
	Value         bool              `json:"value"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Operator      respjson.Field
		StringToCheck respjson.Field
		Inputs        respjson.Field
		PropertyName  respjson.Field
		Value         respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Contains) RawJSON() string { return r.JSON.raw }
func (r *Contains) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ContainsOperator string

const (
	ContainsOperatorContains ContainsOperator = "CONTAINS"
)

type Date struct {
	// Any of "DATE".
	Operator     DateOperator      `json:"operator,required"`
	Inputs       []ExpressionUnion `json:"inputs"`
	PropertyName string            `json:"propertyName"`
	Value        float64           `json:"value"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Operator     respjson.Field
		Inputs       respjson.Field
		PropertyName respjson.Field
		Value        respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Date) RawJSON() string { return r.JSON.raw }
func (r *Date) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DateOperator string

const (
	DateOperatorDate DateOperator = "DATE"
)

type DatedExchangeRate struct {
	// Any of "DATED_EXCHANGE_RATE".
	Operator     DatedExchangeRateOperator `json:"operator,required"`
	Inputs       []ExpressionUnion         `json:"inputs"`
	PropertyName string                    `json:"propertyName"`
	Value        float64                   `json:"value"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Operator     respjson.Field
		Inputs       respjson.Field
		PropertyName respjson.Field
		Value        respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DatedExchangeRate) RawJSON() string { return r.JSON.raw }
func (r *DatedExchangeRate) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DatedExchangeRateOperator string

const (
	DatedExchangeRateOperatorDatedExchangeRate DatedExchangeRateOperator = "DATED_EXCHANGE_RATE"
)

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

type DivideNumbers struct {
	EnclosedInParentheses bool `json:"enclosedInParentheses,required"`
	// Any of "DIVIDE_NUMBERS".
	Operator     DivideNumbersOperator `json:"operator,required"`
	Inputs       []ExpressionUnion     `json:"inputs"`
	PropertyName string                `json:"propertyName"`
	Value        float64               `json:"value"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		EnclosedInParentheses respjson.Field
		Operator              respjson.Field
		Inputs                respjson.Field
		PropertyName          respjson.Field
		Value                 respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DivideNumbers) RawJSON() string { return r.JSON.raw }
func (r *DivideNumbers) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DivideNumbersOperator string

const (
	DivideNumbersOperatorDivideNumbers DivideNumbersOperator = "DIVIDE_NUMBERS"
)

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

type Euler struct {
	// Any of "EULER".
	Operator     EulerOperator     `json:"operator,required"`
	Inputs       []ExpressionUnion `json:"inputs"`
	PropertyName string            `json:"propertyName"`
	Value        float64           `json:"value"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Operator     respjson.Field
		Inputs       respjson.Field
		PropertyName respjson.Field
		Value        respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Euler) RawJSON() string { return r.JSON.raw }
func (r *Euler) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EulerOperator string

const (
	EulerOperatorEuler EulerOperator = "EULER"
)

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

// ExpressionUnion contains all possible properties and values from
// [ConstantBoolean], [ConstantNumber], [ConstantString],
// [BooleanPropertyVariable], [StringPropertyVariable], [NumberPropertyVariable],
// [TimestampOfPropertyVariable], [BooleanTargetPropertyVariable],
// [StringTargetPropertyVariable], [NumberTargetPropertyVariable],
// [TimestampOfTargetPropertyVariable], [AddNumbers], [SubtractNumbers],
// [MultiplyNumbers], [DivideNumbers], [RoundDownNumbers], [RoundUpNumbers],
// [RoundNearestNumbers], [UpperCase], [LowerCase], [ConcatStrings], [Contains],
// [BeginsWith], [NumberToString], [ParseNumber], [FetchExchangeRate],
// [FetchCurrencyDecimalPlaces], [FetchSingleCurrencyPortalCurrency],
// [DatedExchangeRate], [PipelineProbability], [MaxNumbers], [MinNumbers],
// [LessThan], [LessThanOrEqual], [MoreThan], [MoreThanOrEqual], [NumberEquals],
// [StringEquals], [IsPipelineStageClosed], [Not], [Date], [Month], [Year], [Now],
// [TimeBetween], [PeriodToMonths], [PeriodToWeeks], [And], [Or], [Xor],
// [IfString], [IfNumber], [IfBoolean], [IsPresent], [HasEmailReply],
// [HasPlainTextEmailReply], [ExtractMostRecentEmailReplyHTML],
// [ExtractMostRecentEmailReplyText], [ExtractMostRecentPlainTextEmailReply],
// [SetContainsString], [IsEngagementType], [FormatFullName], [AbsoluteValue],
// [SquareRoot], [Power], [Substring], [Euler], [StringLength], [AddTime],
// [SubtractTime].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ExpressionUnion struct {
	Operator     string            `json:"operator"`
	Inputs       []ExpressionUnion `json:"inputs"`
	PropertyName string            `json:"propertyName"`
	// This field is a union of [bool], [float64], [string], [bool], [string],
	// [float64], [string], [bool], [string], [float64], [string], [float64],
	// [float64], [float64], [float64], [float64], [float64], [float64], [string],
	// [string], [string], [bool], [bool], [string], [float64], [float64], [float64],
	// [string], [float64], [float64], [float64], [float64], [bool], [bool], [bool],
	// [bool], [bool], [bool], [bool], [bool], [float64], [float64], [float64],
	// [float64], [float64], [float64], [float64], [bool], [bool], [bool], [string],
	// [float64], [bool], [bool], [bool], [bool], [string], [string], [string], [bool],
	// [bool], [string], [float64], [float64], [float64], [string], [float64],
	// [float64], [float64], [float64]
	Value                 ExpressionUnionValue `json:"value"`
	EnclosedInParentheses bool                 `json:"enclosedInParentheses"`
	// This field is from variant [Contains].
	StringToCheck ExpressionUnion `json:"stringToCheck"`
	// This field is from variant [IfString].
	IfExpression ExpressionUnion `json:"ifExpression"`
	// This field is from variant [IfString].
	ElseExpression ExpressionUnion `json:"elseExpression"`
	// This field is from variant [IsPresent].
	ExpressionToEvaluate ExpressionUnion `json:"expressionToEvaluate"`
	JSON                 struct {
		Operator              respjson.Field
		Inputs                respjson.Field
		PropertyName          respjson.Field
		Value                 respjson.Field
		EnclosedInParentheses respjson.Field
		StringToCheck         respjson.Field
		IfExpression          respjson.Field
		ElseExpression        respjson.Field
		ExpressionToEvaluate  respjson.Field
		raw                   string
	} `json:"-"`
}

func (u ExpressionUnion) AsConstantBoolean() (v ConstantBoolean) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ExpressionUnion) AsConstantNumber() (v ConstantNumber) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ExpressionUnion) AsConstantString() (v ConstantString) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ExpressionUnion) AsBooleanPropertyVariable() (v BooleanPropertyVariable) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ExpressionUnion) AsStringPropertyVariable() (v StringPropertyVariable) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ExpressionUnion) AsNumberPropertyVariable() (v NumberPropertyVariable) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ExpressionUnion) AsTimestampOfPropertyVariable() (v TimestampOfPropertyVariable) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ExpressionUnion) AsBooleanTargetPropertyVariable() (v BooleanTargetPropertyVariable) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ExpressionUnion) AsStringTargetPropertyVariable() (v StringTargetPropertyVariable) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ExpressionUnion) AsNumberTargetPropertyVariable() (v NumberTargetPropertyVariable) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ExpressionUnion) AsTimestampOfTargetPropertyVariable() (v TimestampOfTargetPropertyVariable) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ExpressionUnion) AsAddNumbers() (v AddNumbers) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ExpressionUnion) AsSubtractNumbers() (v SubtractNumbers) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ExpressionUnion) AsMultiplyNumbers() (v MultiplyNumbers) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ExpressionUnion) AsDivideNumbers() (v DivideNumbers) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ExpressionUnion) AsRoundDown() (v RoundDownNumbers) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ExpressionUnion) AsRoundUp() (v RoundUpNumbers) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ExpressionUnion) AsRoundNearest() (v RoundNearestNumbers) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ExpressionUnion) AsUpperCase() (v UpperCase) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ExpressionUnion) AsLowerCase() (v LowerCase) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ExpressionUnion) AsConcatStrings() (v ConcatStrings) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ExpressionUnion) AsContains() (v Contains) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ExpressionUnion) AsBeginsWith() (v BeginsWith) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ExpressionUnion) AsNumberToString() (v NumberToString) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ExpressionUnion) AsParseNumber() (v ParseNumber) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ExpressionUnion) AsFetchExchangeRate() (v FetchExchangeRate) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ExpressionUnion) AsFetchCurrencyDecimalPlaces() (v FetchCurrencyDecimalPlaces) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ExpressionUnion) AsFetchSingleCurrencyPortalCurrency() (v FetchSingleCurrencyPortalCurrency) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ExpressionUnion) AsDatedExchangeRate() (v DatedExchangeRate) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ExpressionUnion) AsPipelineProbability() (v PipelineProbability) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ExpressionUnion) AsMaxNumbers() (v MaxNumbers) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ExpressionUnion) AsMinNumbers() (v MinNumbers) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ExpressionUnion) AsLessThan() (v LessThan) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ExpressionUnion) AsLessThanOrEqual() (v LessThanOrEqual) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ExpressionUnion) AsMoreThan() (v MoreThan) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ExpressionUnion) AsMoreThanOrEqual() (v MoreThanOrEqual) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ExpressionUnion) AsNumberEquals() (v NumberEquals) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ExpressionUnion) AsStringEquals() (v StringEquals) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ExpressionUnion) AsIsPipelineStageClosed() (v IsPipelineStageClosed) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ExpressionUnion) AsNot() (v Not) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ExpressionUnion) AsDate() (v Date) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ExpressionUnion) AsMonth() (v Month) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ExpressionUnion) AsYear() (v Year) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ExpressionUnion) AsNow() (v Now) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ExpressionUnion) AsTimeBetween() (v TimeBetween) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ExpressionUnion) AsPeriodToMonths() (v PeriodToMonths) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ExpressionUnion) AsPeriodToWeeks() (v PeriodToWeeks) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ExpressionUnion) AsAnd() (v And) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ExpressionUnion) AsOr() (v Or) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ExpressionUnion) AsXor() (v Xor) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ExpressionUnion) AsIfString() (v IfString) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ExpressionUnion) AsIfNumber() (v IfNumber) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ExpressionUnion) AsIfBoolean() (v IfBoolean) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ExpressionUnion) AsIsPresent() (v IsPresent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ExpressionUnion) AsHasEmailReply() (v HasEmailReply) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ExpressionUnion) AsHasPlainTextEmailReply() (v HasPlainTextEmailReply) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ExpressionUnion) AsExtractMostRecentEmailReplyHTML() (v ExtractMostRecentEmailReplyHTML) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ExpressionUnion) AsExtractMostRecentEmailReplyText() (v ExtractMostRecentEmailReplyText) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ExpressionUnion) AsExtractMostRecentPlainTextEmailReply() (v ExtractMostRecentPlainTextEmailReply) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ExpressionUnion) AsSetContainsString() (v SetContainsString) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ExpressionUnion) AsIsEngagementType() (v IsEngagementType) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ExpressionUnion) AsFormatFullName() (v FormatFullName) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ExpressionUnion) AsAbsoluteValue() (v AbsoluteValue) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ExpressionUnion) AsSquareRoot() (v SquareRoot) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ExpressionUnion) AsPower() (v Power) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ExpressionUnion) AsSubstring() (v Substring) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ExpressionUnion) AsEuler() (v Euler) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ExpressionUnion) AsStringLength() (v StringLength) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ExpressionUnion) AsAddTime() (v AddTime) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ExpressionUnion) AsSubtractTime() (v SubtractTime) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ExpressionUnion) RawJSON() string { return u.JSON.raw }

func (r *ExpressionUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ExpressionUnionValue is an implicit subunion of [ExpressionUnion].
// ExpressionUnionValue provides convenient access to the sub-properties of the
// union.
//
// For type safety it is recommended to directly use a variant of the
// [ExpressionUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfBool OfFloat OfString]
type ExpressionUnionValue struct {
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

func (r *ExpressionUnionValue) UnmarshalJSON(data []byte) error {
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

type ExtractMostRecentEmailReplyHTML struct {
	// Any of "EXTRACT_MOST_RECENT_EMAIL_REPLY_HTML".
	Operator     ExtractMostRecentEmailReplyHTMLOperator `json:"operator,required"`
	Inputs       []ExpressionUnion                       `json:"inputs"`
	PropertyName string                                  `json:"propertyName"`
	Value        string                                  `json:"value"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Operator     respjson.Field
		Inputs       respjson.Field
		PropertyName respjson.Field
		Value        respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExtractMostRecentEmailReplyHTML) RawJSON() string { return r.JSON.raw }
func (r *ExtractMostRecentEmailReplyHTML) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExtractMostRecentEmailReplyHTMLOperator string

const (
	ExtractMostRecentEmailReplyHTMLOperatorExtractMostRecentEmailReplyHTML ExtractMostRecentEmailReplyHTMLOperator = "EXTRACT_MOST_RECENT_EMAIL_REPLY_HTML"
)

type ExtractMostRecentEmailReplyText struct {
	// Any of "EXTRACT_MOST_RECENT_EMAIL_REPLY_TEXT".
	Operator     ExtractMostRecentEmailReplyTextOperator `json:"operator,required"`
	Inputs       []ExpressionUnion                       `json:"inputs"`
	PropertyName string                                  `json:"propertyName"`
	Value        string                                  `json:"value"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Operator     respjson.Field
		Inputs       respjson.Field
		PropertyName respjson.Field
		Value        respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExtractMostRecentEmailReplyText) RawJSON() string { return r.JSON.raw }
func (r *ExtractMostRecentEmailReplyText) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExtractMostRecentEmailReplyTextOperator string

const (
	ExtractMostRecentEmailReplyTextOperatorExtractMostRecentEmailReplyText ExtractMostRecentEmailReplyTextOperator = "EXTRACT_MOST_RECENT_EMAIL_REPLY_TEXT"
)

type ExtractMostRecentPlainTextEmailReply struct {
	// Any of "EXTRACT_MOST_RECENT_PLAIN_TEXT_EMAIL_REPLY".
	Operator     ExtractMostRecentPlainTextEmailReplyOperator `json:"operator,required"`
	Inputs       []ExpressionUnion                            `json:"inputs"`
	PropertyName string                                       `json:"propertyName"`
	Value        string                                       `json:"value"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Operator     respjson.Field
		Inputs       respjson.Field
		PropertyName respjson.Field
		Value        respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExtractMostRecentPlainTextEmailReply) RawJSON() string { return r.JSON.raw }
func (r *ExtractMostRecentPlainTextEmailReply) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExtractMostRecentPlainTextEmailReplyOperator string

const (
	ExtractMostRecentPlainTextEmailReplyOperatorExtractMostRecentPlainTextEmailReply ExtractMostRecentPlainTextEmailReplyOperator = "EXTRACT_MOST_RECENT_PLAIN_TEXT_EMAIL_REPLY"
)

type FetchCurrencyDecimalPlaces struct {
	// Any of "FETCH_CURRENCY_DECIMAL_PLACES".
	Operator     FetchCurrencyDecimalPlacesOperator `json:"operator,required"`
	Inputs       []ExpressionUnion                  `json:"inputs"`
	PropertyName string                             `json:"propertyName"`
	Value        float64                            `json:"value"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Operator     respjson.Field
		Inputs       respjson.Field
		PropertyName respjson.Field
		Value        respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FetchCurrencyDecimalPlaces) RawJSON() string { return r.JSON.raw }
func (r *FetchCurrencyDecimalPlaces) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FetchCurrencyDecimalPlacesOperator string

const (
	FetchCurrencyDecimalPlacesOperatorFetchCurrencyDecimalPlaces FetchCurrencyDecimalPlacesOperator = "FETCH_CURRENCY_DECIMAL_PLACES"
)

type FetchExchangeRate struct {
	// Any of "FETCH_EXCHANGE_RATE".
	Operator     FetchExchangeRateOperator `json:"operator,required"`
	Inputs       []ExpressionUnion         `json:"inputs"`
	PropertyName string                    `json:"propertyName"`
	Value        float64                   `json:"value"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Operator     respjson.Field
		Inputs       respjson.Field
		PropertyName respjson.Field
		Value        respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FetchExchangeRate) RawJSON() string { return r.JSON.raw }
func (r *FetchExchangeRate) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FetchExchangeRateOperator string

const (
	FetchExchangeRateOperatorFetchExchangeRate FetchExchangeRateOperator = "FETCH_EXCHANGE_RATE"
)

type FetchSingleCurrencyPortalCurrency struct {
	// Any of "FETCH_SINGLE_CURRENCY_PORTAL_CURRENCY".
	Operator     FetchSingleCurrencyPortalCurrencyOperator `json:"operator,required"`
	Inputs       []ExpressionUnion                         `json:"inputs"`
	PropertyName string                                    `json:"propertyName"`
	Value        string                                    `json:"value"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Operator     respjson.Field
		Inputs       respjson.Field
		PropertyName respjson.Field
		Value        respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FetchSingleCurrencyPortalCurrency) RawJSON() string { return r.JSON.raw }
func (r *FetchSingleCurrencyPortalCurrency) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FetchSingleCurrencyPortalCurrencyOperator string

const (
	FetchSingleCurrencyPortalCurrencyOperatorFetchSingleCurrencyPortalCurrency FetchSingleCurrencyPortalCurrencyOperator = "FETCH_SINGLE_CURRENCY_PORTAL_CURRENCY"
)

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

type FormatFullName struct {
	// Any of "FORMAT_FULL_NAME".
	Operator     FormatFullNameOperator `json:"operator,required"`
	Inputs       []ExpressionUnion      `json:"inputs"`
	PropertyName string                 `json:"propertyName"`
	Value        string                 `json:"value"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Operator     respjson.Field
		Inputs       respjson.Field
		PropertyName respjson.Field
		Value        respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FormatFullName) RawJSON() string { return r.JSON.raw }
func (r *FormatFullName) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FormatFullNameOperator string

const (
	FormatFullNameOperatorFormatFullName FormatFullNameOperator = "FORMAT_FULL_NAME"
)

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

type HasEmailReply struct {
	// Any of "HAS_EMAIL_REPLY".
	Operator     HasEmailReplyOperator `json:"operator,required"`
	Inputs       []ExpressionUnion     `json:"inputs"`
	PropertyName string                `json:"propertyName"`
	Value        bool                  `json:"value"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Operator     respjson.Field
		Inputs       respjson.Field
		PropertyName respjson.Field
		Value        respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r HasEmailReply) RawJSON() string { return r.JSON.raw }
func (r *HasEmailReply) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type HasEmailReplyOperator string

const (
	HasEmailReplyOperatorHasEmailReply HasEmailReplyOperator = "HAS_EMAIL_REPLY"
)

type HasPlainTextEmailReply struct {
	// Any of "HAS_PLAIN_TEXT_EMAIL_REPLY".
	Operator     HasPlainTextEmailReplyOperator `json:"operator,required"`
	Inputs       []ExpressionUnion              `json:"inputs"`
	PropertyName string                         `json:"propertyName"`
	Value        bool                           `json:"value"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Operator     respjson.Field
		Inputs       respjson.Field
		PropertyName respjson.Field
		Value        respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r HasPlainTextEmailReply) RawJSON() string { return r.JSON.raw }
func (r *HasPlainTextEmailReply) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type HasPlainTextEmailReplyOperator string

const (
	HasPlainTextEmailReplyOperatorHasPlainTextEmailReply HasPlainTextEmailReplyOperator = "HAS_PLAIN_TEXT_EMAIL_REPLY"
)

type IfBoolean struct {
	EnclosedInParentheses bool            `json:"enclosedInParentheses,required"`
	IfExpression          ExpressionUnion `json:"ifExpression,required"`
	// Any of "IF_BOOLEAN".
	Operator       IfBooleanOperator `json:"operator,required"`
	ElseExpression ExpressionUnion   `json:"elseExpression"`
	Inputs         []ExpressionUnion `json:"inputs"`
	PropertyName   string            `json:"propertyName"`
	Value          bool              `json:"value"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		EnclosedInParentheses respjson.Field
		IfExpression          respjson.Field
		Operator              respjson.Field
		ElseExpression        respjson.Field
		Inputs                respjson.Field
		PropertyName          respjson.Field
		Value                 respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IfBoolean) RawJSON() string { return r.JSON.raw }
func (r *IfBoolean) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IfBooleanOperator string

const (
	IfBooleanOperatorIfBoolean IfBooleanOperator = "IF_BOOLEAN"
)

type IfNumber struct {
	EnclosedInParentheses bool            `json:"enclosedInParentheses,required"`
	IfExpression          ExpressionUnion `json:"ifExpression,required"`
	// Any of "IF_NUMBER".
	Operator       IfNumberOperator  `json:"operator,required"`
	ElseExpression ExpressionUnion   `json:"elseExpression"`
	Inputs         []ExpressionUnion `json:"inputs"`
	PropertyName   string            `json:"propertyName"`
	Value          float64           `json:"value"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		EnclosedInParentheses respjson.Field
		IfExpression          respjson.Field
		Operator              respjson.Field
		ElseExpression        respjson.Field
		Inputs                respjson.Field
		PropertyName          respjson.Field
		Value                 respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IfNumber) RawJSON() string { return r.JSON.raw }
func (r *IfNumber) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IfNumberOperator string

const (
	IfNumberOperatorIfNumber IfNumberOperator = "IF_NUMBER"
)

type IfString struct {
	EnclosedInParentheses bool            `json:"enclosedInParentheses,required"`
	IfExpression          ExpressionUnion `json:"ifExpression,required"`
	// Any of "IF_STRING".
	Operator       IfStringOperator  `json:"operator,required"`
	ElseExpression ExpressionUnion   `json:"elseExpression"`
	Inputs         []ExpressionUnion `json:"inputs"`
	PropertyName   string            `json:"propertyName"`
	Value          string            `json:"value"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		EnclosedInParentheses respjson.Field
		IfExpression          respjson.Field
		Operator              respjson.Field
		ElseExpression        respjson.Field
		Inputs                respjson.Field
		PropertyName          respjson.Field
		Value                 respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IfString) RawJSON() string { return r.JSON.raw }
func (r *IfString) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IfStringOperator string

const (
	IfStringOperatorIfString IfStringOperator = "IF_STRING"
)

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

type IsEngagementType struct {
	// Any of "IS_ENGAGEMENT_TYPE".
	Operator     IsEngagementTypeOperator `json:"operator,required"`
	Inputs       []ExpressionUnion        `json:"inputs"`
	PropertyName string                   `json:"propertyName"`
	Value        bool                     `json:"value"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Operator     respjson.Field
		Inputs       respjson.Field
		PropertyName respjson.Field
		Value        respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IsEngagementType) RawJSON() string { return r.JSON.raw }
func (r *IsEngagementType) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IsEngagementTypeOperator string

const (
	IsEngagementTypeOperatorIsEngagementType IsEngagementTypeOperator = "IS_ENGAGEMENT_TYPE"
)

type IsPipelineStageClosed struct {
	// Any of "IS_PIPELINE_STAGE_CLOSED".
	Operator     IsPipelineStageClosedOperator `json:"operator,required"`
	Inputs       []ExpressionUnion             `json:"inputs"`
	PropertyName string                        `json:"propertyName"`
	Value        bool                          `json:"value"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Operator     respjson.Field
		Inputs       respjson.Field
		PropertyName respjson.Field
		Value        respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IsPipelineStageClosed) RawJSON() string { return r.JSON.raw }
func (r *IsPipelineStageClosed) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IsPipelineStageClosedOperator string

const (
	IsPipelineStageClosedOperatorIsPipelineStageClosed IsPipelineStageClosedOperator = "IS_PIPELINE_STAGE_CLOSED"
)

type IsPresent struct {
	ExpressionToEvaluate ExpressionUnion `json:"expressionToEvaluate,required"`
	// Any of "IS_PRESENT".
	Operator     IsPresentOperator `json:"operator,required"`
	Inputs       []ExpressionUnion `json:"inputs"`
	PropertyName string            `json:"propertyName"`
	Value        bool              `json:"value"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ExpressionToEvaluate respjson.Field
		Operator             respjson.Field
		Inputs               respjson.Field
		PropertyName         respjson.Field
		Value                respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IsPresent) RawJSON() string { return r.JSON.raw }
func (r *IsPresent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IsPresentOperator string

const (
	IsPresentOperatorIsPresent IsPresentOperator = "IS_PRESENT"
)

type LessThan struct {
	// Any of "LESS_THAN".
	Operator     LessThanOperator  `json:"operator,required"`
	Inputs       []ExpressionUnion `json:"inputs"`
	PropertyName string            `json:"propertyName"`
	Value        bool              `json:"value"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Operator     respjson.Field
		Inputs       respjson.Field
		PropertyName respjson.Field
		Value        respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LessThan) RawJSON() string { return r.JSON.raw }
func (r *LessThan) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LessThanOperator string

const (
	LessThanOperatorLessThan LessThanOperator = "LESS_THAN"
)

type LessThanOrEqual struct {
	// Any of "LESS_THAN_OR_EQUAL".
	Operator     LessThanOrEqualOperator `json:"operator,required"`
	Inputs       []ExpressionUnion       `json:"inputs"`
	PropertyName string                  `json:"propertyName"`
	Value        bool                    `json:"value"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Operator     respjson.Field
		Inputs       respjson.Field
		PropertyName respjson.Field
		Value        respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LessThanOrEqual) RawJSON() string { return r.JSON.raw }
func (r *LessThanOrEqual) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LessThanOrEqualOperator string

const (
	LessThanOrEqualOperatorLessThanOrEqual LessThanOrEqualOperator = "LESS_THAN_OR_EQUAL"
)

type LowerCase struct {
	// Any of "LOWER_CASE".
	Operator     LowerCaseOperator `json:"operator,required"`
	Inputs       []ExpressionUnion `json:"inputs"`
	PropertyName string            `json:"propertyName"`
	Value        string            `json:"value"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Operator     respjson.Field
		Inputs       respjson.Field
		PropertyName respjson.Field
		Value        respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LowerCase) RawJSON() string { return r.JSON.raw }
func (r *LowerCase) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LowerCaseOperator string

const (
	LowerCaseOperatorLowerCase LowerCaseOperator = "LOWER_CASE"
)

type MaxNumbers struct {
	// Any of "MAX_NUMBERS".
	Operator     MaxNumbersOperator `json:"operator,required"`
	Inputs       []ExpressionUnion  `json:"inputs"`
	PropertyName string             `json:"propertyName"`
	Value        float64            `json:"value"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Operator     respjson.Field
		Inputs       respjson.Field
		PropertyName respjson.Field
		Value        respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MaxNumbers) RawJSON() string { return r.JSON.raw }
func (r *MaxNumbers) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MaxNumbersOperator string

const (
	MaxNumbersOperatorMaxNumbers MaxNumbersOperator = "MAX_NUMBERS"
)

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

type MinNumbers struct {
	// Any of "MIN_NUMBERS".
	Operator     MinNumbersOperator `json:"operator,required"`
	Inputs       []ExpressionUnion  `json:"inputs"`
	PropertyName string             `json:"propertyName"`
	Value        float64            `json:"value"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Operator     respjson.Field
		Inputs       respjson.Field
		PropertyName respjson.Field
		Value        respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MinNumbers) RawJSON() string { return r.JSON.raw }
func (r *MinNumbers) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MinNumbersOperator string

const (
	MinNumbersOperatorMinNumbers MinNumbersOperator = "MIN_NUMBERS"
)

type Month struct {
	// Any of "MONTH".
	Operator     MonthOperator     `json:"operator,required"`
	Inputs       []ExpressionUnion `json:"inputs"`
	PropertyName string            `json:"propertyName"`
	Value        float64           `json:"value"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Operator     respjson.Field
		Inputs       respjson.Field
		PropertyName respjson.Field
		Value        respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Month) RawJSON() string { return r.JSON.raw }
func (r *Month) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MonthOperator string

const (
	MonthOperatorMonth MonthOperator = "MONTH"
)

type MoreThan struct {
	// Any of "MORE_THAN".
	Operator     MoreThanOperator  `json:"operator,required"`
	Inputs       []ExpressionUnion `json:"inputs"`
	PropertyName string            `json:"propertyName"`
	Value        bool              `json:"value"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Operator     respjson.Field
		Inputs       respjson.Field
		PropertyName respjson.Field
		Value        respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MoreThan) RawJSON() string { return r.JSON.raw }
func (r *MoreThan) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MoreThanOperator string

const (
	MoreThanOperatorMoreThan MoreThanOperator = "MORE_THAN"
)

type MoreThanOrEqual struct {
	// Any of "MORE_THAN_OR_EQUAL".
	Operator     MoreThanOrEqualOperator `json:"operator,required"`
	Inputs       []ExpressionUnion       `json:"inputs"`
	PropertyName string                  `json:"propertyName"`
	Value        bool                    `json:"value"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Operator     respjson.Field
		Inputs       respjson.Field
		PropertyName respjson.Field
		Value        respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MoreThanOrEqual) RawJSON() string { return r.JSON.raw }
func (r *MoreThanOrEqual) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MoreThanOrEqualOperator string

const (
	MoreThanOrEqualOperatorMoreThanOrEqual MoreThanOrEqualOperator = "MORE_THAN_OR_EQUAL"
)

type MultiplyNumbers struct {
	EnclosedInParentheses bool `json:"enclosedInParentheses,required"`
	// Any of "MULTIPLY_NUMBERS".
	Operator     MultiplyNumbersOperator `json:"operator,required"`
	Inputs       []ExpressionUnion       `json:"inputs"`
	PropertyName string                  `json:"propertyName"`
	Value        float64                 `json:"value"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		EnclosedInParentheses respjson.Field
		Operator              respjson.Field
		Inputs                respjson.Field
		PropertyName          respjson.Field
		Value                 respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MultiplyNumbers) RawJSON() string { return r.JSON.raw }
func (r *MultiplyNumbers) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MultiplyNumbersOperator string

const (
	MultiplyNumbersOperatorMultiplyNumbers MultiplyNumbersOperator = "MULTIPLY_NUMBERS"
)

type Not struct {
	// Any of "NOT".
	Operator     NotOperator       `json:"operator,required"`
	Inputs       []ExpressionUnion `json:"inputs"`
	PropertyName string            `json:"propertyName"`
	Value        bool              `json:"value"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Operator     respjson.Field
		Inputs       respjson.Field
		PropertyName respjson.Field
		Value        respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Not) RawJSON() string { return r.JSON.raw }
func (r *Not) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NotOperator string

const (
	NotOperatorNot NotOperator = "NOT"
)

type Now struct {
	// Any of "NOW".
	Operator     NowOperator       `json:"operator,required"`
	Inputs       []ExpressionUnion `json:"inputs"`
	PropertyName string            `json:"propertyName"`
	Value        float64           `json:"value"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Operator     respjson.Field
		Inputs       respjson.Field
		PropertyName respjson.Field
		Value        respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Now) RawJSON() string { return r.JSON.raw }
func (r *Now) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NowOperator string

const (
	NowOperatorNow NowOperator = "NOW"
)

type NumberEquals struct {
	// Any of "NUMBER_EQUALS".
	Operator     NumberEqualsOperator `json:"operator,required"`
	Inputs       []ExpressionUnion    `json:"inputs"`
	PropertyName string               `json:"propertyName"`
	Value        bool                 `json:"value"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Operator     respjson.Field
		Inputs       respjson.Field
		PropertyName respjson.Field
		Value        respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NumberEquals) RawJSON() string { return r.JSON.raw }
func (r *NumberEquals) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NumberEqualsOperator string

const (
	NumberEqualsOperatorNumberEquals NumberEqualsOperator = "NUMBER_EQUALS"
)

type NumberPropertyVariable struct {
	// Any of "NUMBER_PROPERTY_VARIABLE".
	Operator     NumberPropertyVariableOperator `json:"operator,required"`
	Inputs       []ExpressionUnion              `json:"inputs"`
	PropertyName string                         `json:"propertyName"`
	Value        float64                        `json:"value"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Operator     respjson.Field
		Inputs       respjson.Field
		PropertyName respjson.Field
		Value        respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NumberPropertyVariable) RawJSON() string { return r.JSON.raw }
func (r *NumberPropertyVariable) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NumberPropertyVariableOperator string

const (
	NumberPropertyVariableOperatorNumberPropertyVariable NumberPropertyVariableOperator = "NUMBER_PROPERTY_VARIABLE"
)

type NumberTargetPropertyVariable struct {
	// Any of "NUMBER_TARGET_PROPERTY_VARIABLE".
	Operator     NumberTargetPropertyVariableOperator `json:"operator,required"`
	Inputs       []ExpressionUnion                    `json:"inputs"`
	PropertyName string                               `json:"propertyName"`
	Value        float64                              `json:"value"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Operator     respjson.Field
		Inputs       respjson.Field
		PropertyName respjson.Field
		Value        respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NumberTargetPropertyVariable) RawJSON() string { return r.JSON.raw }
func (r *NumberTargetPropertyVariable) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NumberTargetPropertyVariableOperator string

const (
	NumberTargetPropertyVariableOperatorNumberTargetPropertyVariable NumberTargetPropertyVariableOperator = "NUMBER_TARGET_PROPERTY_VARIABLE"
)

type NumberToString struct {
	// Any of "NUMBER_TO_STRING".
	Operator     NumberToStringOperator `json:"operator,required"`
	Inputs       []ExpressionUnion      `json:"inputs"`
	PropertyName string                 `json:"propertyName"`
	Value        string                 `json:"value"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Operator     respjson.Field
		Inputs       respjson.Field
		PropertyName respjson.Field
		Value        respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NumberToString) RawJSON() string { return r.JSON.raw }
func (r *NumberToString) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NumberToStringOperator string

const (
	NumberToStringOperatorNumberToString NumberToStringOperator = "NUMBER_TO_STRING"
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

type Or struct {
	EnclosedInParentheses bool `json:"enclosedInParentheses,required"`
	// Any of "OR".
	Operator     OrOperator        `json:"operator,required"`
	Inputs       []ExpressionUnion `json:"inputs"`
	PropertyName string            `json:"propertyName"`
	Value        bool              `json:"value"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		EnclosedInParentheses respjson.Field
		Operator              respjson.Field
		Inputs                respjson.Field
		PropertyName          respjson.Field
		Value                 respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Or) RawJSON() string { return r.JSON.raw }
func (r *Or) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OrOperator string

const (
	OrOperatorOr OrOperator = "OR"
)

type ParseNumber struct {
	// Any of "PARSE_NUMBER".
	Operator     ParseNumberOperator `json:"operator,required"`
	Inputs       []ExpressionUnion   `json:"inputs"`
	PropertyName string              `json:"propertyName"`
	Value        float64             `json:"value"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Operator     respjson.Field
		Inputs       respjson.Field
		PropertyName respjson.Field
		Value        respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ParseNumber) RawJSON() string { return r.JSON.raw }
func (r *ParseNumber) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ParseNumberOperator string

const (
	ParseNumberOperatorParseNumber ParseNumberOperator = "PARSE_NUMBER"
)

type PeriodToMonths struct {
	// Any of "PERIOD_TO_MONTHS".
	Operator     PeriodToMonthsOperator `json:"operator,required"`
	Inputs       []ExpressionUnion      `json:"inputs"`
	PropertyName string                 `json:"propertyName"`
	Value        float64                `json:"value"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Operator     respjson.Field
		Inputs       respjson.Field
		PropertyName respjson.Field
		Value        respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PeriodToMonths) RawJSON() string { return r.JSON.raw }
func (r *PeriodToMonths) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PeriodToMonthsOperator string

const (
	PeriodToMonthsOperatorPeriodToMonths PeriodToMonthsOperator = "PERIOD_TO_MONTHS"
)

type PeriodToWeeks struct {
	// Any of "PERIOD_TO_WEEKS".
	Operator     PeriodToWeeksOperator `json:"operator,required"`
	Inputs       []ExpressionUnion     `json:"inputs"`
	PropertyName string                `json:"propertyName"`
	Value        float64               `json:"value"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Operator     respjson.Field
		Inputs       respjson.Field
		PropertyName respjson.Field
		Value        respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PeriodToWeeks) RawJSON() string { return r.JSON.raw }
func (r *PeriodToWeeks) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PeriodToWeeksOperator string

const (
	PeriodToWeeksOperatorPeriodToWeeks PeriodToWeeksOperator = "PERIOD_TO_WEEKS"
)

type PipelineProbability struct {
	// Any of "PIPELINE_PROBABILITY".
	Operator     PipelineProbabilityOperator `json:"operator,required"`
	Inputs       []ExpressionUnion           `json:"inputs"`
	PropertyName string                      `json:"propertyName"`
	Value        float64                     `json:"value"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Operator     respjson.Field
		Inputs       respjson.Field
		PropertyName respjson.Field
		Value        respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PipelineProbability) RawJSON() string { return r.JSON.raw }
func (r *PipelineProbability) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PipelineProbabilityOperator string

const (
	PipelineProbabilityOperatorPipelineProbability PipelineProbabilityOperator = "PIPELINE_PROBABILITY"
)

type Power struct {
	// Any of "POWER".
	Operator     PowerOperator     `json:"operator,required"`
	Inputs       []ExpressionUnion `json:"inputs"`
	PropertyName string            `json:"propertyName"`
	Value        float64           `json:"value"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Operator     respjson.Field
		Inputs       respjson.Field
		PropertyName respjson.Field
		Value        respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Power) RawJSON() string { return r.JSON.raw }
func (r *Power) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PowerOperator string

const (
	PowerOperatorPower PowerOperator = "POWER"
)

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
	CalculationExpression    ExpressionUnion          `json:"calculationExpression"`
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
	ConditionalExpression       ExpressionUnion          `json:"conditionalExpression"`
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

type RoundDownNumbers struct {
	// Any of "ROUND_DOWN".
	Operator     RoundDownNumbersOperator `json:"operator,required"`
	Inputs       []ExpressionUnion        `json:"inputs"`
	PropertyName string                   `json:"propertyName"`
	Value        float64                  `json:"value"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Operator     respjson.Field
		Inputs       respjson.Field
		PropertyName respjson.Field
		Value        respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RoundDownNumbers) RawJSON() string { return r.JSON.raw }
func (r *RoundDownNumbers) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RoundDownNumbersOperator string

const (
	RoundDownNumbersOperatorRoundDown RoundDownNumbersOperator = "ROUND_DOWN"
)

type RoundNearestNumbers struct {
	// Any of "ROUND_NEAREST".
	Operator     RoundNearestNumbersOperator `json:"operator,required"`
	Inputs       []ExpressionUnion           `json:"inputs"`
	PropertyName string                      `json:"propertyName"`
	Value        float64                     `json:"value"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Operator     respjson.Field
		Inputs       respjson.Field
		PropertyName respjson.Field
		Value        respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RoundNearestNumbers) RawJSON() string { return r.JSON.raw }
func (r *RoundNearestNumbers) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RoundNearestNumbersOperator string

const (
	RoundNearestNumbersOperatorRoundNearest RoundNearestNumbersOperator = "ROUND_NEAREST"
)

type RoundUpNumbers struct {
	// Any of "ROUND_UP".
	Operator     RoundUpNumbersOperator `json:"operator,required"`
	Inputs       []ExpressionUnion      `json:"inputs"`
	PropertyName string                 `json:"propertyName"`
	Value        float64                `json:"value"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Operator     respjson.Field
		Inputs       respjson.Field
		PropertyName respjson.Field
		Value        respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RoundUpNumbers) RawJSON() string { return r.JSON.raw }
func (r *RoundUpNumbers) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RoundUpNumbersOperator string

const (
	RoundUpNumbersOperatorRoundUp RoundUpNumbersOperator = "ROUND_UP"
)

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

type SetContainsString struct {
	// Any of "SET_CONTAINS_STRING".
	Operator      SetContainsStringOperator `json:"operator,required"`
	StringToCheck ExpressionUnion           `json:"stringToCheck,required"`
	Inputs        []ExpressionUnion         `json:"inputs"`
	PropertyName  string                    `json:"propertyName"`
	Value         bool                      `json:"value"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Operator      respjson.Field
		StringToCheck respjson.Field
		Inputs        respjson.Field
		PropertyName  respjson.Field
		Value         respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SetContainsString) RawJSON() string { return r.JSON.raw }
func (r *SetContainsString) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SetContainsStringOperator string

const (
	SetContainsStringOperatorSetContainsString SetContainsStringOperator = "SET_CONTAINS_STRING"
)

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

type SquareRoot struct {
	// Any of "SQUARE_ROOT".
	Operator     SquareRootOperator `json:"operator,required"`
	Inputs       []ExpressionUnion  `json:"inputs"`
	PropertyName string             `json:"propertyName"`
	Value        float64            `json:"value"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Operator     respjson.Field
		Inputs       respjson.Field
		PropertyName respjson.Field
		Value        respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SquareRoot) RawJSON() string { return r.JSON.raw }
func (r *SquareRoot) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SquareRootOperator string

const (
	SquareRootOperatorSquareRoot SquareRootOperator = "SQUARE_ROOT"
)

type StringEquals struct {
	// Any of "STRING_EQUALS".
	Operator     StringEqualsOperator `json:"operator,required"`
	Inputs       []ExpressionUnion    `json:"inputs"`
	PropertyName string               `json:"propertyName"`
	Value        bool                 `json:"value"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Operator     respjson.Field
		Inputs       respjson.Field
		PropertyName respjson.Field
		Value        respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r StringEquals) RawJSON() string { return r.JSON.raw }
func (r *StringEquals) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type StringEqualsOperator string

const (
	StringEqualsOperatorStringEquals StringEqualsOperator = "STRING_EQUALS"
)

type StringLength struct {
	// Any of "STRING_LENGTH".
	Operator     StringLengthOperator `json:"operator,required"`
	Inputs       []ExpressionUnion    `json:"inputs"`
	PropertyName string               `json:"propertyName"`
	Value        float64              `json:"value"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Operator     respjson.Field
		Inputs       respjson.Field
		PropertyName respjson.Field
		Value        respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r StringLength) RawJSON() string { return r.JSON.raw }
func (r *StringLength) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type StringLengthOperator string

const (
	StringLengthOperatorStringLength StringLengthOperator = "STRING_LENGTH"
)

type StringPropertyVariable struct {
	// Any of "STRING_PROPERTY_VARIABLE".
	Operator     StringPropertyVariableOperator `json:"operator,required"`
	Inputs       []ExpressionUnion              `json:"inputs"`
	PropertyName string                         `json:"propertyName"`
	Value        string                         `json:"value"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Operator     respjson.Field
		Inputs       respjson.Field
		PropertyName respjson.Field
		Value        respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r StringPropertyVariable) RawJSON() string { return r.JSON.raw }
func (r *StringPropertyVariable) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type StringPropertyVariableOperator string

const (
	StringPropertyVariableOperatorStringPropertyVariable StringPropertyVariableOperator = "STRING_PROPERTY_VARIABLE"
)

type StringTargetPropertyVariable struct {
	// Any of "STRING_TARGET_PROPERTY_VARIABLE".
	Operator     StringTargetPropertyVariableOperator `json:"operator,required"`
	Inputs       []ExpressionUnion                    `json:"inputs"`
	PropertyName string                               `json:"propertyName"`
	Value        string                               `json:"value"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Operator     respjson.Field
		Inputs       respjson.Field
		PropertyName respjson.Field
		Value        respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r StringTargetPropertyVariable) RawJSON() string { return r.JSON.raw }
func (r *StringTargetPropertyVariable) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type StringTargetPropertyVariableOperator string

const (
	StringTargetPropertyVariableOperatorStringTargetPropertyVariable StringTargetPropertyVariableOperator = "STRING_TARGET_PROPERTY_VARIABLE"
)

type Substring struct {
	// Any of "SUBSTRING".
	Operator      SubstringOperator `json:"operator,required"`
	StringToCheck ExpressionUnion   `json:"stringToCheck,required"`
	Inputs        []ExpressionUnion `json:"inputs"`
	PropertyName  string            `json:"propertyName"`
	Value         string            `json:"value"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Operator      respjson.Field
		StringToCheck respjson.Field
		Inputs        respjson.Field
		PropertyName  respjson.Field
		Value         respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Substring) RawJSON() string { return r.JSON.raw }
func (r *Substring) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SubstringOperator string

const (
	SubstringOperatorSubstring SubstringOperator = "SUBSTRING"
)

type SubtractNumbers struct {
	EnclosedInParentheses bool `json:"enclosedInParentheses,required"`
	// Any of "SUBTRACT_NUMBERS".
	Operator     SubtractNumbersOperator `json:"operator,required"`
	Inputs       []ExpressionUnion       `json:"inputs"`
	PropertyName string                  `json:"propertyName"`
	Value        float64                 `json:"value"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		EnclosedInParentheses respjson.Field
		Operator              respjson.Field
		Inputs                respjson.Field
		PropertyName          respjson.Field
		Value                 respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SubtractNumbers) RawJSON() string { return r.JSON.raw }
func (r *SubtractNumbers) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SubtractNumbersOperator string

const (
	SubtractNumbersOperatorSubtractNumbers SubtractNumbersOperator = "SUBTRACT_NUMBERS"
)

type SubtractTime struct {
	// Any of "SUBTRACT_TIME".
	Operator      SubtractTimeOperator `json:"operator,required"`
	StringToCheck ExpressionUnion      `json:"stringToCheck,required"`
	Inputs        []ExpressionUnion    `json:"inputs"`
	PropertyName  string               `json:"propertyName"`
	Value         float64              `json:"value"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Operator      respjson.Field
		StringToCheck respjson.Field
		Inputs        respjson.Field
		PropertyName  respjson.Field
		Value         respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SubtractTime) RawJSON() string { return r.JSON.raw }
func (r *SubtractTime) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SubtractTimeOperator string

const (
	SubtractTimeOperatorSubtractTime SubtractTimeOperator = "SUBTRACT_TIME"
)

type TimeBetween struct {
	// Any of "TIME_BETWEEN".
	Operator     TimeBetweenOperator `json:"operator,required"`
	Inputs       []ExpressionUnion   `json:"inputs"`
	PropertyName string              `json:"propertyName"`
	Value        float64             `json:"value"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Operator     respjson.Field
		Inputs       respjson.Field
		PropertyName respjson.Field
		Value        respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TimeBetween) RawJSON() string { return r.JSON.raw }
func (r *TimeBetween) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TimeBetweenOperator string

const (
	TimeBetweenOperatorTimeBetween TimeBetweenOperator = "TIME_BETWEEN"
)

type TimestampOfPropertyVariable struct {
	// Any of "TIMESTAMP_OF_PROPERTY_VARIABLE".
	Operator     TimestampOfPropertyVariableOperator `json:"operator,required"`
	Inputs       []ExpressionUnion                   `json:"inputs"`
	PropertyName string                              `json:"propertyName"`
	Value        string                              `json:"value"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Operator     respjson.Field
		Inputs       respjson.Field
		PropertyName respjson.Field
		Value        respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TimestampOfPropertyVariable) RawJSON() string { return r.JSON.raw }
func (r *TimestampOfPropertyVariable) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TimestampOfPropertyVariableOperator string

const (
	TimestampOfPropertyVariableOperatorTimestampOfPropertyVariable TimestampOfPropertyVariableOperator = "TIMESTAMP_OF_PROPERTY_VARIABLE"
)

type TimestampOfTargetPropertyVariable struct {
	// Any of "TIMESTAMP_OF_TARGET_PROPERTY_VARIABLE".
	Operator     TimestampOfTargetPropertyVariableOperator `json:"operator,required"`
	Inputs       []ExpressionUnion                         `json:"inputs"`
	PropertyName string                                    `json:"propertyName"`
	Value        string                                    `json:"value"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Operator     respjson.Field
		Inputs       respjson.Field
		PropertyName respjson.Field
		Value        respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TimestampOfTargetPropertyVariable) RawJSON() string { return r.JSON.raw }
func (r *TimestampOfTargetPropertyVariable) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TimestampOfTargetPropertyVariableOperator string

const (
	TimestampOfTargetPropertyVariableOperatorTimestampOfTargetPropertyVariable TimestampOfTargetPropertyVariableOperator = "TIMESTAMP_OF_TARGET_PROPERTY_VARIABLE"
)

type UpperCase struct {
	// Any of "UPPER_CASE".
	Operator     UpperCaseOperator `json:"operator,required"`
	Inputs       []ExpressionUnion `json:"inputs"`
	PropertyName string            `json:"propertyName"`
	Value        string            `json:"value"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Operator     respjson.Field
		Inputs       respjson.Field
		PropertyName respjson.Field
		Value        respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UpperCase) RawJSON() string { return r.JSON.raw }
func (r *UpperCase) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UpperCaseOperator string

const (
	UpperCaseOperatorUpperCase UpperCaseOperator = "UPPER_CASE"
)

type Xor struct {
	EnclosedInParentheses bool `json:"enclosedInParentheses,required"`
	// Any of "XOR".
	Operator     XorOperator       `json:"operator,required"`
	Inputs       []ExpressionUnion `json:"inputs"`
	PropertyName string            `json:"propertyName"`
	Value        bool              `json:"value"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		EnclosedInParentheses respjson.Field
		Operator              respjson.Field
		Inputs                respjson.Field
		PropertyName          respjson.Field
		Value                 respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Xor) RawJSON() string { return r.JSON.raw }
func (r *Xor) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type XorOperator string

const (
	XorOperatorXor XorOperator = "XOR"
)

type Year struct {
	// Any of "YEAR".
	Operator     YearOperator      `json:"operator,required"`
	Inputs       []ExpressionUnion `json:"inputs"`
	PropertyName string            `json:"propertyName"`
	Value        float64           `json:"value"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Operator     respjson.Field
		Inputs       respjson.Field
		PropertyName respjson.Field
		Value        respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Year) RawJSON() string { return r.JSON.raw }
func (r *Year) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type YearOperator string

const (
	YearOperatorYear YearOperator = "YEAR"
)
