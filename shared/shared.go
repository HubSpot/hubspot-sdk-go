// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package shared

import (
	"encoding/json"
	"time"

	"github.com/HubSpot/hubspot-sdk-go/internal/apijson"
	"github.com/HubSpot/hubspot-sdk-go/packages/param"
	"github.com/HubSpot/hubspot-sdk-go/packages/respjson"
)

// aliased to make [param.APIUnion] private when embedding
type paramUnion = param.APIUnion

// aliased to make [param.APIObject] private when embedding
type paramObj = param.APIObject

// The properties ContentID, VariationName are required.
type AbTestCreateRequestVNextParam struct {
	// ID of the object to test.
	ContentID string `json:"contentId" api:"required"`
	// Name of A/B test variation.
	VariationName string `json:"variationName" api:"required"`
	paramObj
}

func (r AbTestCreateRequestVNextParam) MarshalJSON() (data []byte, err error) {
	type shadow AbTestCreateRequestVNextParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AbTestCreateRequestVNextParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ActionResponse struct {
	// The timestamp indicating when the action was completed.
	CompletedAt time.Time `json:"completedAt" api:"required" format:"date-time"`
	// The timestamp indicating when the action was started.
	StartedAt time.Time `json:"startedAt" api:"required" format:"date-time"`
	// The current status of the action, with possible values: CANCELED, COMPLETE,
	// PENDING, PROCESSING.
	//
	// Any of "CANCELED", "COMPLETE", "PENDING", "PROCESSING".
	Status ActionResponseStatus `json:"status" api:"required"`
	// A map of link names to associated URIs containing documentation about the error
	// or recommended remediation steps
	Links map[string]string `json:"links"`
	// The timestamp indicating when the action was requested.
	RequestedAt time.Time `json:"requestedAt" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CompletedAt respjson.Field
		StartedAt   respjson.Field
		Status      respjson.Field
		Links       respjson.Field
		RequestedAt respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ActionResponse) RawJSON() string { return r.JSON.raw }
func (r *ActionResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The current status of the action, with possible values: CANCELED, COMPLETE,
// PENDING, PROCESSING.
type ActionResponseStatus string

const (
	ActionResponseStatusCanceled   ActionResponseStatus = "CANCELED"
	ActionResponseStatusComplete   ActionResponseStatus = "COMPLETE"
	ActionResponseStatusPending    ActionResponseStatus = "PENDING"
	ActionResponseStatusProcessing ActionResponseStatus = "PROCESSING"
)

// The definition of an association
type AssociationDefinition struct {
	// The unique ID of the associated object (e.g., a contact ID).
	ID string `json:"id" api:"required"`
	// The ID of the source object type (e.g., 0-1 for contacts).
	FromObjectTypeID string `json:"fromObjectTypeId" api:"required"`
	// The ID of the destination object type (e.g., 0-3 for deals).
	ToObjectTypeID string `json:"toObjectTypeId" api:"required"`
	// The timestamp when the association was created, in ISO 8601 format.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// For labeled association types, the internal name of the association.
	Name string `json:"name"`
	// The timestamp when the last update was made to an association, in ISO 8601
	// format.
	UpdatedAt time.Time `json:"updatedAt" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		FromObjectTypeID respjson.Field
		ToObjectTypeID   respjson.Field
		CreatedAt        respjson.Field
		Name             respjson.Field
		UpdatedAt        respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AssociationDefinition) RawJSON() string { return r.JSON.raw }
func (r *AssociationDefinition) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties FromObjectTypeID, ToObjectTypeID are required.
type AssociationDefinitionEggParam struct {
	FromObjectTypeID string            `json:"fromObjectTypeId" api:"required"`
	ToObjectTypeID   string            `json:"toObjectTypeId" api:"required"`
	Name             param.Opt[string] `json:"name,omitzero"`
	paramObj
}

func (r AssociationDefinitionEggParam) MarshalJSON() (data []byte, err error) {
	type shadow AssociationDefinitionEggParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AssociationDefinitionEggParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Defines the type, direction, and details of the relationship between two CRM
// objects.
type AssociationSpec struct {
	// The category of the association, such as "HUBSPOT_DEFINED".
	//
	// Any of "HUBSPOT_DEFINED", "INTEGRATOR_DEFINED", "USER_DEFINED", "WORK".
	AssociationCategory AssociationSpecAssociationCategory `json:"associationCategory" api:"required"`
	// The ID representing the specific type of association.
	AssociationTypeID int64 `json:"associationTypeId" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AssociationCategory respjson.Field
		AssociationTypeID   respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AssociationSpec) RawJSON() string { return r.JSON.raw }
func (r *AssociationSpec) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this AssociationSpec to a AssociationSpecParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// AssociationSpecParam.Overrides()
func (r AssociationSpec) ToParam() AssociationSpecParam {
	return param.Override[AssociationSpecParam](json.RawMessage(r.RawJSON()))
}

// The category of the association, such as "HUBSPOT_DEFINED".
type AssociationSpecAssociationCategory string

const (
	AssociationSpecAssociationCategoryHubSpotDefined    AssociationSpecAssociationCategory = "HUBSPOT_DEFINED"
	AssociationSpecAssociationCategoryIntegratorDefined AssociationSpecAssociationCategory = "INTEGRATOR_DEFINED"
	AssociationSpecAssociationCategoryUserDefined       AssociationSpecAssociationCategory = "USER_DEFINED"
	AssociationSpecAssociationCategoryWork              AssociationSpecAssociationCategory = "WORK"
)

// Defines the type, direction, and details of the relationship between two CRM
// objects.
//
// The properties AssociationCategory, AssociationTypeID are required.
type AssociationSpecParam struct {
	// The category of the association, such as "HUBSPOT_DEFINED".
	//
	// Any of "HUBSPOT_DEFINED", "INTEGRATOR_DEFINED", "USER_DEFINED", "WORK".
	AssociationCategory AssociationSpecAssociationCategory `json:"associationCategory,omitzero" api:"required"`
	// The ID representing the specific type of association.
	AssociationTypeID int64 `json:"associationTypeId" api:"required"`
	paramObj
}

func (r AssociationSpecParam) MarshalJSON() (data []byte, err error) {
	type shadow AssociationSpecParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AssociationSpecParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A HubSpot property option
type AutomationActionsOption struct {
	// A description of the option.
	Description string `json:"description" api:"required"`
	// The position of the item relative to others in the list.
	DisplayOrder int64 `json:"displayOrder" api:"required"`
	// A numerical value associated with the option.
	DoubleData float64 `json:"doubleData" api:"required"`
	// Whether the option is displayed in HubSpot's UI.
	Hidden bool `json:"hidden" api:"required"`
	// A user-friendly label that identifies the option.
	Label string `json:"label" api:"required"`
	// Whether the option is read-only.
	ReadOnly bool `json:"readOnly" api:"required"`
	// The actual value of the option.
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
func (r AutomationActionsOption) RawJSON() string { return r.JSON.raw }
func (r *AutomationActionsOption) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this AutomationActionsOption to a AutomationActionsOptionParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// AutomationActionsOptionParam.Overrides()
func (r AutomationActionsOption) ToParam() AutomationActionsOptionParam {
	return param.Override[AutomationActionsOptionParam](json.RawMessage(r.RawJSON()))
}

// A HubSpot property option
//
// The properties Description, DisplayOrder, DoubleData, Hidden, Label, ReadOnly,
// Value are required.
type AutomationActionsOptionParam struct {
	// A description of the option.
	Description string `json:"description" api:"required"`
	// The position of the item relative to others in the list.
	DisplayOrder int64 `json:"displayOrder" api:"required"`
	// A numerical value associated with the option.
	DoubleData float64 `json:"doubleData" api:"required"`
	// Whether the option is displayed in HubSpot's UI.
	Hidden bool `json:"hidden" api:"required"`
	// A user-friendly label that identifies the option.
	Label string `json:"label" api:"required"`
	// Whether the option is read-only.
	ReadOnly bool `json:"readOnly" api:"required"`
	// The actual value of the option.
	Value string `json:"value" api:"required"`
	paramObj
}

func (r AutomationActionsOptionParam) MarshalJSON() (data []byte, err error) {
	type shadow AutomationActionsOptionParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AutomationActionsOptionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Inputs is required.
type BatchInputPropertyCreateParam struct {
	Inputs []PropertyCreateParam `json:"inputs,omitzero" api:"required"`
	paramObj
}

func (r BatchInputPropertyCreateParam) MarshalJSON() (data []byte, err error) {
	type shadow BatchInputPropertyCreateParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BatchInputPropertyCreateParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Inputs is required.
type BatchInputPropertyNameParam struct {
	Inputs []PropertyNameParam `json:"inputs,omitzero" api:"required"`
	paramObj
}

func (r BatchInputPropertyNameParam) MarshalJSON() (data []byte, err error) {
	type shadow BatchInputPropertyNameParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BatchInputPropertyNameParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Inputs is required.
type BatchInputPublicObjectIDParam struct {
	// An array of deal split inputs
	Inputs []PublicObjectIDParam `json:"inputs,omitzero" api:"required"`
	paramObj
}

func (r BatchInputPublicObjectIDParam) MarshalJSON() (data []byte, err error) {
	type shadow BatchInputPublicObjectIDParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BatchInputPublicObjectIDParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Inputs is required.
type BatchInputStringParam struct {
	// Strings to input.
	Inputs []string `json:"inputs,omitzero" api:"required"`
	paramObj
}

func (r BatchInputStringParam) MarshalJSON() (data []byte, err error) {
	type shadow BatchInputStringParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BatchInputStringParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Archived, DataSensitivity, Inputs are required.
type BatchReadInputPropertyNameParam struct {
	Archived bool `json:"archived" api:"required"`
	// Any of "highly_sensitive", "non_sensitive", "sensitive".
	DataSensitivity BatchReadInputPropertyNameDataSensitivity `json:"dataSensitivity,omitzero" api:"required"`
	Inputs          []PropertyNameParam                       `json:"inputs,omitzero" api:"required"`
	paramObj
}

func (r BatchReadInputPropertyNameParam) MarshalJSON() (data []byte, err error) {
	type shadow BatchReadInputPropertyNameParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BatchReadInputPropertyNameParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BatchReadInputPropertyNameDataSensitivity string

const (
	BatchReadInputPropertyNameDataSensitivityHighlySensitive BatchReadInputPropertyNameDataSensitivity = "highly_sensitive"
	BatchReadInputPropertyNameDataSensitivityNonSensitive    BatchReadInputPropertyNameDataSensitivity = "non_sensitive"
	BatchReadInputPropertyNameDataSensitivitySensitive       BatchReadInputPropertyNameDataSensitivity = "sensitive"
)

type CollectionResponsePropertyGroupNoPaging struct {
	Results []PropertyGroup `json:"results" api:"required"`
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

type ErrorData struct {
	// The error category
	Category string `json:"category" api:"required"`
	// A unique identifier for the request. Include this value with any error reports
	// or support tickets
	CorrelationID string `json:"correlationId" api:"required" format:"uuid"`
	// A human readable message describing the error along with remediation steps where
	// appropriate
	Message string `json:"message" api:"required"`
	// Context about the error condition
	Context map[string][]string `json:"context"`
	// further information about the error
	Errors []ErrorDetail `json:"errors"`
	// A map of link names to associated URIs containing documentation about the error
	// or recommended remediation steps
	Links map[string]string `json:"links"`
	// A specific category that contains more specific detail about the error
	SubCategory string `json:"subCategory"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Category      respjson.Field
		CorrelationID respjson.Field
		Message       respjson.Field
		Context       respjson.Field
		Errors        respjson.Field
		Links         respjson.Field
		SubCategory   respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ErrorData) RawJSON() string { return r.JSON.raw }
func (r *ErrorData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ErrorDetail struct {
	// A human readable message describing the error along with remediation steps where
	// appropriate
	Message string `json:"message" api:"required"`
	// The status code associated with the error detail
	Code string `json:"code"`
	// Context about the error condition
	Context map[string][]string `json:"context"`
	// The name of the field or parameter in which the error was found.
	In string `json:"in"`
	// A specific category that contains more specific detail about the error
	SubCategory string `json:"subCategory"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Message     respjson.Field
		Code        respjson.Field
		Context     respjson.Field
		In          respjson.Field
		SubCategory respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ErrorDetail) RawJSON() string { return r.JSON.raw }
func (r *ErrorDetail) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ForwardPaging struct {
	// Specifies the paging information needed to retrieve the next set of results in a
	// paginated API response
	Next NextPage `json:"next"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Next        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ForwardPaging) RawJSON() string { return r.JSON.raw }
func (r *ForwardPaging) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Specifies the paging information needed to retrieve the next set of results in a
// paginated API response
type NextPage struct {
	// A paging cursor token for retrieving subsequent pages.
	After string `json:"after" api:"required"`
	// A URL that can be used to retrieve the next page results.
	Link string `json:"link"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		After       respjson.Field
		Link        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NextPage) RawJSON() string { return r.JSON.raw }
func (r *NextPage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ObjectTypeDefinition struct {
	ID                         string                     `json:"id" api:"required"`
	AllowsSensitiveProperties  bool                       `json:"allowsSensitiveProperties" api:"required"`
	Archived                   bool                       `json:"archived" api:"required"`
	FullyQualifiedName         string                     `json:"fullyQualifiedName" api:"required"`
	Labels                     ObjectTypeDefinitionLabels `json:"labels" api:"required"`
	Name                       string                     `json:"name" api:"required"`
	ObjectTypeID               string                     `json:"objectTypeId" api:"required"`
	RequiredProperties         []string                   `json:"requiredProperties" api:"required"`
	SearchableProperties       []string                   `json:"searchableProperties" api:"required"`
	SecondaryDisplayProperties []string                   `json:"secondaryDisplayProperties" api:"required"`
	CreatedAt                  time.Time                  `json:"createdAt" format:"date-time"`
	Description                string                     `json:"description"`
	PortalID                   int64                      `json:"portalId"`
	PrimaryDisplayProperty     string                     `json:"primaryDisplayProperty"`
	UpdatedAt                  time.Time                  `json:"updatedAt" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                         respjson.Field
		AllowsSensitiveProperties  respjson.Field
		Archived                   respjson.Field
		FullyQualifiedName         respjson.Field
		Labels                     respjson.Field
		Name                       respjson.Field
		ObjectTypeID               respjson.Field
		RequiredProperties         respjson.Field
		SearchableProperties       respjson.Field
		SecondaryDisplayProperties respjson.Field
		CreatedAt                  respjson.Field
		Description                respjson.Field
		PortalID                   respjson.Field
		PrimaryDisplayProperty     respjson.Field
		UpdatedAt                  respjson.Field
		ExtraFields                map[string]respjson.Field
		raw                        string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ObjectTypeDefinition) RawJSON() string { return r.JSON.raw }
func (r *ObjectTypeDefinition) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ObjectTypeDefinitionLabels struct {
	Plural   string `json:"plural"`
	Singular string `json:"singular"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Plural      respjson.Field
		Singular    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ObjectTypeDefinitionLabels) RawJSON() string { return r.JSON.raw }
func (r *ObjectTypeDefinitionLabels) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this ObjectTypeDefinitionLabels to a
// ObjectTypeDefinitionLabelsParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// ObjectTypeDefinitionLabelsParam.Overrides()
func (r ObjectTypeDefinitionLabels) ToParam() ObjectTypeDefinitionLabelsParam {
	return param.Override[ObjectTypeDefinitionLabelsParam](json.RawMessage(r.RawJSON()))
}

type ObjectTypeDefinitionLabelsParam struct {
	Plural   param.Opt[string] `json:"plural,omitzero"`
	Singular param.Opt[string] `json:"singular,omitzero"`
	paramObj
}

func (r ObjectTypeDefinitionLabelsParam) MarshalJSON() (data []byte, err error) {
	type shadow ObjectTypeDefinitionLabelsParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ObjectTypeDefinitionLabelsParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property ClearDescription is required.
type ObjectTypeDefinitionPatchParam struct {
	ClearDescription           bool                            `json:"clearDescription" api:"required"`
	AllowsSensitiveProperties  param.Opt[bool]                 `json:"allowsSensitiveProperties,omitzero"`
	Description                param.Opt[string]               `json:"description,omitzero"`
	PrimaryDisplayProperty     param.Opt[string]               `json:"primaryDisplayProperty,omitzero"`
	Restorable                 param.Opt[bool]                 `json:"restorable,omitzero"`
	Labels                     ObjectTypeDefinitionLabelsParam `json:"labels,omitzero"`
	RequiredProperties         []string                        `json:"requiredProperties,omitzero"`
	SearchableProperties       []string                        `json:"searchableProperties,omitzero"`
	SecondaryDisplayProperties []string                        `json:"secondaryDisplayProperties,omitzero"`
	paramObj
}

func (r ObjectTypeDefinitionPatchParam) MarshalJSON() (data []byte, err error) {
	type shadow ObjectTypeDefinitionPatchParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ObjectTypeDefinitionPatchParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A HubSpot property option
type Option struct {
	// Hidden options will not be displayed in HubSpot.
	Hidden bool `json:"hidden" api:"required"`
	// A human-readable option label that will be shown in HubSpot.
	Label string `json:"label" api:"required"`
	// The internal value of the option, which must be used when setting the property
	// value through the API.
	Value string `json:"value" api:"required"`
	// A description of the option.
	Description string `json:"description"`
	// Options are displayed in order starting with the lowest positive integer value.
	// Values of -1 will cause the option to be displayed after any positive values.
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

// The properties DisplayOrder, Hidden, Label, Value are required.
type OptionInputParam struct {
	DisplayOrder int64             `json:"displayOrder" api:"required"`
	Hidden       bool              `json:"hidden" api:"required"`
	Label        string            `json:"label" api:"required"`
	Value        string            `json:"value" api:"required"`
	Description  param.Opt[string] `json:"description,omitzero"`
	paramObj
}

func (r OptionInputParam) MarshalJSON() (data []byte, err error) {
	type shadow OptionInputParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OptionInputParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Paging struct {
	// Specifies the paging information needed to retrieve the next set of results in a
	// paginated API response
	Next NextPage `json:"next"`
	// specifies the paging information needed to retrieve the previous set of results
	// in a paginated API response
	Prev PreviousPage `json:"prev"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Next        respjson.Field
		Prev        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Paging) RawJSON() string { return r.JSON.raw }
func (r *Paging) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// specifies the paging information needed to retrieve the previous set of results
// in a paginated API response
type PreviousPage struct {
	// A string token used to identify the position before the current page in the
	// pagination sequence.
	Before string `json:"before" api:"required"`
	// A URL string that provides a direct link to the previous page of results.
	Link string `json:"link"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Before      respjson.Field
		Link        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PreviousPage) RawJSON() string { return r.JSON.raw }
func (r *PreviousPage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A HubSpot property
type Property struct {
	// A description of the property that will be shown as help text in HubSpot.
	Description string `json:"description" api:"required"`
	// Controls how the property appears in HubSpot.
	FieldType string `json:"fieldType" api:"required"`
	// The name of the property group the property belongs to.
	GroupName string `json:"groupName" api:"required"`
	// A human-readable property label that will be shown in HubSpot.
	Label string `json:"label" api:"required"`
	// The internal property name, which must be used when referencing the property via
	// the API.
	Name string `json:"name" api:"required"`
	// A list of valid options for the property. This field is required for enumerated
	// properties, but will be empty for other property types.
	Options []Option `json:"options" api:"required"`
	// The property data type.
	Type string `json:"type" api:"required"`
	// Whether or not the property is archived.
	Archived bool `json:"archived"`
	// When the property was archived.
	ArchivedAt time.Time `json:"archivedAt" format:"date-time"`
	// For default properties, true indicates that the property is calculated by a
	// HubSpot process. It has no effect for custom properties.
	Calculated bool `json:"calculated"`
	// The formula used for calculated properties.
	CalculationFormula string `json:"calculationFormula"`
	// When the property was created
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// The internal ID of the user who created the property in HubSpot. This field may
	// not exist if the property was created outside of HubSpot.
	CreatedUserID string `json:"createdUserId"`
	// The name of the related currency property.
	CurrencyPropertyName string `json:"currencyPropertyName"`
	// Indicates the sensitivity level of the property, such as "non_sensitive",
	// "sensitive", or "highly_sensitive".
	//
	// Any of "highly_sensitive", "non_sensitive", "sensitive".
	DataSensitivity PropertyDataSensitivity `json:"dataSensitivity"`
	// Controls how date properties are displayed in the HubSpot UI, with options such
	// as 'absolute', 'absolute_with_relative', 'time_since', and 'time_until'.
	//
	// Any of "absolute", "absolute_with_relative", "time_since", "time_until".
	DateDisplayHint PropertyDateDisplayHint `json:"dateDisplayHint"`
	// The order that this property should be displayed in the HubSpot UI relative to
	// other properties for this object type. Properties are displayed in order
	// starting with the lowest positive integer value. A value of -1 will cause the
	// property to be displayed **after** any positive values.
	DisplayOrder int64 `json:"displayOrder"`
	// For default properties, true indicates that the options are stored externally to
	// the property settings.
	ExternalOptions bool `json:"externalOptions"`
	// Whether or not the property can be used in a HubSpot form.
	FormField bool `json:"formField"`
	// Whether or not the property's value must be unique. Once set, this can't be
	// changed.
	HasUniqueValue bool `json:"hasUniqueValue"`
	// Hidden options won't be shown in HubSpot.
	Hidden bool `json:"hidden"`
	// This will be true for default object properties built into HubSpot.
	HubSpotDefined       bool                         `json:"hubspotDefined"`
	ModificationMetadata PropertyModificationMetadata `json:"modificationMetadata"`
	// Hint for how a number property is displayed and validated in HubSpot's UI. Can
	// be: "unformatted", "formatted", "currency", "percentage", "duration", or
	// "probability".
	//
	// Any of "currency", "duration", "formatted", "percentage", "probability",
	// "unformatted".
	NumberDisplayHint PropertyNumberDisplayHint `json:"numberDisplayHint"`
	// If this property is related to other object(s), they'll be listed here.
	ReferencedObjectType string `json:"referencedObjectType"`
	// When sensitiveData is true, lists the type of sensitive data contained in the
	// property (e.g., "HIPAA").
	SensitiveDataCategories []string `json:"sensitiveDataCategories"`
	// Whether the property will display the currency symbol set in the account
	// settings.
	ShowCurrencySymbol bool `json:"showCurrencySymbol"`
	// When the object type was last updated.
	UpdatedAt time.Time `json:"updatedAt" format:"date-time"`
	// The internal user ID of the user who updated the property in HubSpot. This field
	// may not exist if the property was updated outside of HubSpot.
	UpdatedUserID string `json:"updatedUserId"`
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
		CurrencyPropertyName    respjson.Field
		DataSensitivity         respjson.Field
		DateDisplayHint         respjson.Field
		DisplayOrder            respjson.Field
		ExternalOptions         respjson.Field
		FormField               respjson.Field
		HasUniqueValue          respjson.Field
		Hidden                  respjson.Field
		HubSpotDefined          respjson.Field
		ModificationMetadata    respjson.Field
		NumberDisplayHint       respjson.Field
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

// Controls how date properties are displayed in the HubSpot UI, with options such
// as 'absolute', 'absolute_with_relative', 'time_since', and 'time_until'.
type PropertyDateDisplayHint string

const (
	PropertyDateDisplayHintAbsolute             PropertyDateDisplayHint = "absolute"
	PropertyDateDisplayHintAbsoluteWithRelative PropertyDateDisplayHint = "absolute_with_relative"
	PropertyDateDisplayHintTimeSince            PropertyDateDisplayHint = "time_since"
	PropertyDateDisplayHintTimeUntil            PropertyDateDisplayHint = "time_until"
)

// Hint for how a number property is displayed and validated in HubSpot's UI. Can
// be: "unformatted", "formatted", "currency", "percentage", "duration", or
// "probability".
type PropertyNumberDisplayHint string

const (
	PropertyNumberDisplayHintCurrency    PropertyNumberDisplayHint = "currency"
	PropertyNumberDisplayHintDuration    PropertyNumberDisplayHint = "duration"
	PropertyNumberDisplayHintFormatted   PropertyNumberDisplayHint = "formatted"
	PropertyNumberDisplayHintPercentage  PropertyNumberDisplayHint = "percentage"
	PropertyNumberDisplayHintProbability PropertyNumberDisplayHint = "probability"
	PropertyNumberDisplayHintUnformatted PropertyNumberDisplayHint = "unformatted"
)

// The properties FieldType, GroupName, Label, Name, Type are required.
type PropertyCreateParam struct {
	// Any of "booleancheckbox", "calculation_equation", "checkbox", "date", "file",
	// "html", "number", "phonenumber", "radio", "select", "text", "textarea".
	FieldType PropertyCreateFieldType `json:"fieldType,omitzero" api:"required"`
	GroupName string                  `json:"groupName" api:"required"`
	Label     string                  `json:"label" api:"required"`
	Name      string                  `json:"name" api:"required"`
	// Any of "bool", "date", "datetime", "enumeration", "number", "phone_number",
	// "string".
	Type                 PropertyCreateType `json:"type,omitzero" api:"required"`
	CalculationFormula   param.Opt[string]  `json:"calculationFormula,omitzero"`
	CurrencyPropertyName param.Opt[string]  `json:"currencyPropertyName,omitzero"`
	Description          param.Opt[string]  `json:"description,omitzero"`
	DisplayOrder         param.Opt[int64]   `json:"displayOrder,omitzero"`
	ExternalOptions      param.Opt[bool]    `json:"externalOptions,omitzero"`
	FormField            param.Opt[bool]    `json:"formField,omitzero"`
	HasUniqueValue       param.Opt[bool]    `json:"hasUniqueValue,omitzero"`
	Hidden               param.Opt[bool]    `json:"hidden,omitzero"`
	ReferencedObjectType param.Opt[string]  `json:"referencedObjectType,omitzero"`
	ShowCurrencySymbol   param.Opt[bool]    `json:"showCurrencySymbol,omitzero"`
	// Any of "highly_sensitive", "non_sensitive", "sensitive".
	DataSensitivity PropertyCreateDataSensitivity `json:"dataSensitivity,omitzero"`
	// Any of "currency", "duration", "formatted", "percentage", "probability",
	// "unformatted".
	NumberDisplayHint PropertyCreateNumberDisplayHint `json:"numberDisplayHint,omitzero"`
	Options           []OptionInputParam              `json:"options,omitzero"`
	paramObj
}

func (r PropertyCreateParam) MarshalJSON() (data []byte, err error) {
	type shadow PropertyCreateParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PropertyCreateParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PropertyCreateFieldType string

const (
	PropertyCreateFieldTypeBooleancheckbox     PropertyCreateFieldType = "booleancheckbox"
	PropertyCreateFieldTypeCalculationEquation PropertyCreateFieldType = "calculation_equation"
	PropertyCreateFieldTypeCheckbox            PropertyCreateFieldType = "checkbox"
	PropertyCreateFieldTypeDate                PropertyCreateFieldType = "date"
	PropertyCreateFieldTypeFile                PropertyCreateFieldType = "file"
	PropertyCreateFieldTypeHTML                PropertyCreateFieldType = "html"
	PropertyCreateFieldTypeNumber              PropertyCreateFieldType = "number"
	PropertyCreateFieldTypePhonenumber         PropertyCreateFieldType = "phonenumber"
	PropertyCreateFieldTypeRadio               PropertyCreateFieldType = "radio"
	PropertyCreateFieldTypeSelect              PropertyCreateFieldType = "select"
	PropertyCreateFieldTypeText                PropertyCreateFieldType = "text"
	PropertyCreateFieldTypeTextarea            PropertyCreateFieldType = "textarea"
)

type PropertyCreateType string

const (
	PropertyCreateTypeBool        PropertyCreateType = "bool"
	PropertyCreateTypeDate        PropertyCreateType = "date"
	PropertyCreateTypeDatetime    PropertyCreateType = "datetime"
	PropertyCreateTypeEnumeration PropertyCreateType = "enumeration"
	PropertyCreateTypeNumber      PropertyCreateType = "number"
	PropertyCreateTypePhoneNumber PropertyCreateType = "phone_number"
	PropertyCreateTypeString      PropertyCreateType = "string"
)

type PropertyCreateDataSensitivity string

const (
	PropertyCreateDataSensitivityHighlySensitive PropertyCreateDataSensitivity = "highly_sensitive"
	PropertyCreateDataSensitivityNonSensitive    PropertyCreateDataSensitivity = "non_sensitive"
	PropertyCreateDataSensitivitySensitive       PropertyCreateDataSensitivity = "sensitive"
)

type PropertyCreateNumberDisplayHint string

const (
	PropertyCreateNumberDisplayHintCurrency    PropertyCreateNumberDisplayHint = "currency"
	PropertyCreateNumberDisplayHintDuration    PropertyCreateNumberDisplayHint = "duration"
	PropertyCreateNumberDisplayHintFormatted   PropertyCreateNumberDisplayHint = "formatted"
	PropertyCreateNumberDisplayHintPercentage  PropertyCreateNumberDisplayHint = "percentage"
	PropertyCreateNumberDisplayHintProbability PropertyCreateNumberDisplayHint = "probability"
	PropertyCreateNumberDisplayHintUnformatted PropertyCreateNumberDisplayHint = "unformatted"
)

type PropertyGroup struct {
	Archived     bool   `json:"archived" api:"required"`
	Label        string `json:"label" api:"required"`
	Name         string `json:"name" api:"required"`
	DisplayOrder int64  `json:"displayOrder"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Archived     respjson.Field
		Label        respjson.Field
		Name         respjson.Field
		DisplayOrder respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PropertyGroup) RawJSON() string { return r.JSON.raw }
func (r *PropertyGroup) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Label, Name are required.
type PropertyGroupCreateParam struct {
	Label        string           `json:"label" api:"required"`
	Name         string           `json:"name" api:"required"`
	DisplayOrder param.Opt[int64] `json:"displayOrder,omitzero"`
	paramObj
}

func (r PropertyGroupCreateParam) MarshalJSON() (data []byte, err error) {
	type shadow PropertyGroupCreateParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PropertyGroupCreateParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PropertyGroupUpdateParam struct {
	DisplayOrder param.Opt[int64]  `json:"displayOrder,omitzero"`
	Label        param.Opt[string] `json:"label,omitzero"`
	paramObj
}

func (r PropertyGroupUpdateParam) MarshalJSON() (data []byte, err error) {
	type shadow PropertyGroupUpdateParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PropertyGroupUpdateParam) UnmarshalJSON(data []byte) error {
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

// The property Name is required.
type PropertyNameParam struct {
	Name string `json:"name" api:"required"`
	paramObj
}

func (r PropertyNameParam) MarshalJSON() (data []byte, err error) {
	type shadow PropertyNameParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PropertyNameParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Represents a single custom property of a marketing event, storing its name,
// value, metadata (like source, timestamp, and sensitivity), and related audit
// information for tracking changes.
type PropertyValue struct {
	// The sensitivity level of the property, such as "non_sensitive", "sensitive", and
	// "highly_sensitive".
	//
	// Any of "high", "none", "standard".
	DataSensitivity PropertyValueDataSensitivity `json:"dataSensitivity" api:"required"`
	// Whether the property value is encrypted.
	IsEncrypted bool `json:"isEncrypted" api:"required"`
	// Indicates if the value exceeds normal size limits.
	IsLargeValue bool `json:"isLargeValue" api:"required"`
	// The unique property name.
	Name string `json:"name" api:"required"`
	// When the value was persisted to database, in epoch milliseconds.
	PersistenceTimestamp int64 `json:"persistenceTimestamp" api:"required"`
	// A unique ID associated with this request.
	RequestID string `json:"requestId" api:"required"`
	// Whether the value was selected by a user.
	SelectedByUser bool `json:"selectedByUser" api:"required"`
	// The timestamp when the value was selected by a user, if applicable.
	SelectedByUserTimestamp int64 `json:"selectedByUserTimestamp" api:"required"`
	// The origin of the property value, such as "IMPORT" or "API".
	//
	// Any of "ACADEMY", "ACCEPTANCE_TEST", "ACTIVITY_AUTO_ASSOCIATE",
	// "ACTIVITY_LOG_REVERT", "ADS", "AI_GROUP", "ANALYTICS", "API", "APPROVALS",
	// "ASSISTS", "ASSOCIATIONS", "AUTO_ASSOCIATE_BY_DOMAIN", "AUTOMATION_JOURNEY",
	// "AUTOMATION_PLATFORM", "AVATARS_SERVICE", "BATCH_UPDATE", "BCC_TO_CRM",
	// "BEHAVIORAL_EVENTS", "BET_ASSIGNMENT", "BET_CRM_CONNECTOR", "BIDEN", "BILLING",
	// "BOT", "CALCULATED", "CENTRAL_EXCHANGE_RATES", "CHATSPOT", "CLONE_OBJECTS",
	// "COMMUNICATOR", "COMPANIES", "COMPANY_FAMILIES", "COMPANY_INSIGHTS",
	// "CONNECTED_ACCOUNT", "CONTACTS", "CONTACTS_WEB", "CONTENT_MEMBERSHIP",
	// "CONVERSATIONAL_ENRICHMENT", "CONVERSATIONS", "CRM_PROCESSES_PLATFORM",
	// "CRM_UI", "CRM_UI_BULK_ACTION", "CUSTOMER_AGENT", "DATA_ENRICHMENT",
	// "DATA_QUALITY", "DATASET", "DEALS", "DEFAULT", "DELETE_OBJECTS", "EMAIL",
	// "EMAIL_INBOX_IMPORT", "EMAIL_INTEGRATION", "ENGAGEMENTS", "EXTENSION",
	// "FILE_MANAGER", "FLYWHEEL_PRODUCT_DATA_SYNC", "FORECASTING", "FORM",
	// "FORWARD_TO_CRM", "GMAIL_INTEGRATION", "GOALS", "HEISENBERG", "HELP_DESK",
	// "HELP_DESK_AI", "IMPORT", "INTEGRATION", "INTEGRATIONS_PLATFORM",
	// "INTEGRATIONS_SYNC", "INTENT", "INTERNAL_PROCESSING", "LEADIN",
	// "LEGAL_BASIS_REMEDIATION", "MARKET_SOURCING", "MARKETPLACE", "MARKETS",
	// "MEETINGS", "MERGE_COMPANIES", "MERGE_CONTACTS", "MERGE_OBJECTS",
	// "MERGE_REVERT_OBJECTS", "MICROAPPS", "MIGRATION", "MOBILE_ANDROID",
	// "MOBILE_IOS", "PAYMENTS", "PIPELINE_SETTINGS", "PLAYBOOKS",
	// "PORTAL_OBJECT_SYNC", "PORTAL_USER_ASSOCIATOR", "PRESENTATIONS",
	// "PRIMARY_AUTOMATION", "PROPERTY_DEFAULT_VALUE", "PROPERTY_RESTORE",
	// "PROPERTY_SETTINGS", "PROSPECTING_AGENT", "QUOTAS", "QUOTES", "RECYCLING_BIN",
	// "RESTORE_OBJECTS", "REVENUE_PLATFORM", "SALES", "SALES_MESSAGES", "SALESFORCE",
	// "SEQUENCES", "SETTINGS", "SIDEKICK", "SIGNALS", "SLACK_INTEGRATION",
	// "SMART_DATA_CAPTURE", "SOCIAL", "SUCCESS", "TALLY", "TASK", "UNKNOWN",
	// "WAL_INCREMENTAL", "WORK_UI", "WORKFLOW_CONTACT_DELETE_ACTION", "WORKFLOWS".
	Source PropertyValueSource `json:"source" api:"required"`
	// The ID of the property source indicating where it was created.
	SourceID string `json:"sourceId" api:"required"`
	// A human-readable label.
	SourceLabel string `json:"sourceLabel" api:"required"`
	// Metadata providing additional context about the source.
	SourceMetadata           string `json:"sourceMetadata" api:"required"`
	SourceUpstreamDeployable string `json:"sourceUpstreamDeployable" api:"required"`
	// The unique identifier associated with the source.
	SourceVid []int64 `json:"sourceVid" api:"required"`
	// When the value was set, as a 64-bit integer.
	Timestamp int64 `json:"timestamp" api:"required"`
	// The unit of measurement or context for the value.
	Unit string `json:"unit" api:"required"`
	// The ID of the user who updated the property.
	UpdatedByUserID int64 `json:"updatedByUserId" api:"required"`
	// Flag indicating whether to use the timestamp field as the persistence timestamp.
	UseTimestampAsPersistenceTimestamp bool `json:"useTimestampAsPersistenceTimestamp" api:"required"`
	// The property value.
	Value string `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DataSensitivity                    respjson.Field
		IsEncrypted                        respjson.Field
		IsLargeValue                       respjson.Field
		Name                               respjson.Field
		PersistenceTimestamp               respjson.Field
		RequestID                          respjson.Field
		SelectedByUser                     respjson.Field
		SelectedByUserTimestamp            respjson.Field
		Source                             respjson.Field
		SourceID                           respjson.Field
		SourceLabel                        respjson.Field
		SourceMetadata                     respjson.Field
		SourceUpstreamDeployable           respjson.Field
		SourceVid                          respjson.Field
		Timestamp                          respjson.Field
		Unit                               respjson.Field
		UpdatedByUserID                    respjson.Field
		UseTimestampAsPersistenceTimestamp respjson.Field
		Value                              respjson.Field
		ExtraFields                        map[string]respjson.Field
		raw                                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PropertyValue) RawJSON() string { return r.JSON.raw }
func (r *PropertyValue) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this PropertyValue to a PropertyValueParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// PropertyValueParam.Overrides()
func (r PropertyValue) ToParam() PropertyValueParam {
	return param.Override[PropertyValueParam](json.RawMessage(r.RawJSON()))
}

// The sensitivity level of the property, such as "non_sensitive", "sensitive", and
// "highly_sensitive".
type PropertyValueDataSensitivity string

const (
	PropertyValueDataSensitivityHigh     PropertyValueDataSensitivity = "high"
	PropertyValueDataSensitivityNone     PropertyValueDataSensitivity = "none"
	PropertyValueDataSensitivityStandard PropertyValueDataSensitivity = "standard"
)

// The origin of the property value, such as "IMPORT" or "API".
type PropertyValueSource string

const (
	PropertyValueSourceAcademy                     PropertyValueSource = "ACADEMY"
	PropertyValueSourceAcceptanceTest              PropertyValueSource = "ACCEPTANCE_TEST"
	PropertyValueSourceActivityAutoAssociate       PropertyValueSource = "ACTIVITY_AUTO_ASSOCIATE"
	PropertyValueSourceActivityLogRevert           PropertyValueSource = "ACTIVITY_LOG_REVERT"
	PropertyValueSourceAds                         PropertyValueSource = "ADS"
	PropertyValueSourceAIGroup                     PropertyValueSource = "AI_GROUP"
	PropertyValueSourceAnalytics                   PropertyValueSource = "ANALYTICS"
	PropertyValueSourceAPI                         PropertyValueSource = "API"
	PropertyValueSourceApprovals                   PropertyValueSource = "APPROVALS"
	PropertyValueSourceAssists                     PropertyValueSource = "ASSISTS"
	PropertyValueSourceAssociations                PropertyValueSource = "ASSOCIATIONS"
	PropertyValueSourceAutoAssociateByDomain       PropertyValueSource = "AUTO_ASSOCIATE_BY_DOMAIN"
	PropertyValueSourceAutomationJourney           PropertyValueSource = "AUTOMATION_JOURNEY"
	PropertyValueSourceAutomationPlatform          PropertyValueSource = "AUTOMATION_PLATFORM"
	PropertyValueSourceAvatarsService              PropertyValueSource = "AVATARS_SERVICE"
	PropertyValueSourceBatchUpdate                 PropertyValueSource = "BATCH_UPDATE"
	PropertyValueSourceBccToCrm                    PropertyValueSource = "BCC_TO_CRM"
	PropertyValueSourceBehavioralEvents            PropertyValueSource = "BEHAVIORAL_EVENTS"
	PropertyValueSourceBetAssignment               PropertyValueSource = "BET_ASSIGNMENT"
	PropertyValueSourceBetCrmConnector             PropertyValueSource = "BET_CRM_CONNECTOR"
	PropertyValueSourceBiden                       PropertyValueSource = "BIDEN"
	PropertyValueSourceBilling                     PropertyValueSource = "BILLING"
	PropertyValueSourceBot                         PropertyValueSource = "BOT"
	PropertyValueSourceCalculated                  PropertyValueSource = "CALCULATED"
	PropertyValueSourceCentralExchangeRates        PropertyValueSource = "CENTRAL_EXCHANGE_RATES"
	PropertyValueSourceChatspot                    PropertyValueSource = "CHATSPOT"
	PropertyValueSourceCloneObjects                PropertyValueSource = "CLONE_OBJECTS"
	PropertyValueSourceCommunicator                PropertyValueSource = "COMMUNICATOR"
	PropertyValueSourceCompanies                   PropertyValueSource = "COMPANIES"
	PropertyValueSourceCompanyFamilies             PropertyValueSource = "COMPANY_FAMILIES"
	PropertyValueSourceCompanyInsights             PropertyValueSource = "COMPANY_INSIGHTS"
	PropertyValueSourceConnectedAccount            PropertyValueSource = "CONNECTED_ACCOUNT"
	PropertyValueSourceContacts                    PropertyValueSource = "CONTACTS"
	PropertyValueSourceContactsWeb                 PropertyValueSource = "CONTACTS_WEB"
	PropertyValueSourceContentMembership           PropertyValueSource = "CONTENT_MEMBERSHIP"
	PropertyValueSourceConversationalEnrichment    PropertyValueSource = "CONVERSATIONAL_ENRICHMENT"
	PropertyValueSourceConversations               PropertyValueSource = "CONVERSATIONS"
	PropertyValueSourceCrmProcessesPlatform        PropertyValueSource = "CRM_PROCESSES_PLATFORM"
	PropertyValueSourceCrmUi                       PropertyValueSource = "CRM_UI"
	PropertyValueSourceCrmUiBulkAction             PropertyValueSource = "CRM_UI_BULK_ACTION"
	PropertyValueSourceCustomerAgent               PropertyValueSource = "CUSTOMER_AGENT"
	PropertyValueSourceDataEnrichment              PropertyValueSource = "DATA_ENRICHMENT"
	PropertyValueSourceDataQuality                 PropertyValueSource = "DATA_QUALITY"
	PropertyValueSourceDataset                     PropertyValueSource = "DATASET"
	PropertyValueSourceDeals                       PropertyValueSource = "DEALS"
	PropertyValueSourceDefault                     PropertyValueSource = "DEFAULT"
	PropertyValueSourceDeleteObjects               PropertyValueSource = "DELETE_OBJECTS"
	PropertyValueSourceEmail                       PropertyValueSource = "EMAIL"
	PropertyValueSourceEmailInboxImport            PropertyValueSource = "EMAIL_INBOX_IMPORT"
	PropertyValueSourceEmailIntegration            PropertyValueSource = "EMAIL_INTEGRATION"
	PropertyValueSourceEngagements                 PropertyValueSource = "ENGAGEMENTS"
	PropertyValueSourceExtension                   PropertyValueSource = "EXTENSION"
	PropertyValueSourceFileManager                 PropertyValueSource = "FILE_MANAGER"
	PropertyValueSourceFlywheelProductDataSync     PropertyValueSource = "FLYWHEEL_PRODUCT_DATA_SYNC"
	PropertyValueSourceForecasting                 PropertyValueSource = "FORECASTING"
	PropertyValueSourceForm                        PropertyValueSource = "FORM"
	PropertyValueSourceForwardToCrm                PropertyValueSource = "FORWARD_TO_CRM"
	PropertyValueSourceGmailIntegration            PropertyValueSource = "GMAIL_INTEGRATION"
	PropertyValueSourceGoals                       PropertyValueSource = "GOALS"
	PropertyValueSourceHeisenberg                  PropertyValueSource = "HEISENBERG"
	PropertyValueSourceHelpDesk                    PropertyValueSource = "HELP_DESK"
	PropertyValueSourceHelpDeskAI                  PropertyValueSource = "HELP_DESK_AI"
	PropertyValueSourceImport                      PropertyValueSource = "IMPORT"
	PropertyValueSourceIntegration                 PropertyValueSource = "INTEGRATION"
	PropertyValueSourceIntegrationsPlatform        PropertyValueSource = "INTEGRATIONS_PLATFORM"
	PropertyValueSourceIntegrationsSync            PropertyValueSource = "INTEGRATIONS_SYNC"
	PropertyValueSourceIntent                      PropertyValueSource = "INTENT"
	PropertyValueSourceInternalProcessing          PropertyValueSource = "INTERNAL_PROCESSING"
	PropertyValueSourceLeadin                      PropertyValueSource = "LEADIN"
	PropertyValueSourceLegalBasisRemediation       PropertyValueSource = "LEGAL_BASIS_REMEDIATION"
	PropertyValueSourceMarketSourcing              PropertyValueSource = "MARKET_SOURCING"
	PropertyValueSourceMarketplace                 PropertyValueSource = "MARKETPLACE"
	PropertyValueSourceMarkets                     PropertyValueSource = "MARKETS"
	PropertyValueSourceMeetings                    PropertyValueSource = "MEETINGS"
	PropertyValueSourceMergeCompanies              PropertyValueSource = "MERGE_COMPANIES"
	PropertyValueSourceMergeContacts               PropertyValueSource = "MERGE_CONTACTS"
	PropertyValueSourceMergeObjects                PropertyValueSource = "MERGE_OBJECTS"
	PropertyValueSourceMergeRevertObjects          PropertyValueSource = "MERGE_REVERT_OBJECTS"
	PropertyValueSourceMicroapps                   PropertyValueSource = "MICROAPPS"
	PropertyValueSourceMigration                   PropertyValueSource = "MIGRATION"
	PropertyValueSourceMobileAndroid               PropertyValueSource = "MOBILE_ANDROID"
	PropertyValueSourceMobileIos                   PropertyValueSource = "MOBILE_IOS"
	PropertyValueSourcePayments                    PropertyValueSource = "PAYMENTS"
	PropertyValueSourcePipelineSettings            PropertyValueSource = "PIPELINE_SETTINGS"
	PropertyValueSourcePlaybooks                   PropertyValueSource = "PLAYBOOKS"
	PropertyValueSourcePortalObjectSync            PropertyValueSource = "PORTAL_OBJECT_SYNC"
	PropertyValueSourcePortalUserAssociator        PropertyValueSource = "PORTAL_USER_ASSOCIATOR"
	PropertyValueSourcePresentations               PropertyValueSource = "PRESENTATIONS"
	PropertyValueSourcePrimaryAutomation           PropertyValueSource = "PRIMARY_AUTOMATION"
	PropertyValueSourcePropertyDefaultValue        PropertyValueSource = "PROPERTY_DEFAULT_VALUE"
	PropertyValueSourcePropertyRestore             PropertyValueSource = "PROPERTY_RESTORE"
	PropertyValueSourcePropertySettings            PropertyValueSource = "PROPERTY_SETTINGS"
	PropertyValueSourceProspectingAgent            PropertyValueSource = "PROSPECTING_AGENT"
	PropertyValueSourceQuotas                      PropertyValueSource = "QUOTAS"
	PropertyValueSourceQuotes                      PropertyValueSource = "QUOTES"
	PropertyValueSourceRecyclingBin                PropertyValueSource = "RECYCLING_BIN"
	PropertyValueSourceRestoreObjects              PropertyValueSource = "RESTORE_OBJECTS"
	PropertyValueSourceRevenuePlatform             PropertyValueSource = "REVENUE_PLATFORM"
	PropertyValueSourceSales                       PropertyValueSource = "SALES"
	PropertyValueSourceSalesMessages               PropertyValueSource = "SALES_MESSAGES"
	PropertyValueSourceSalesforce                  PropertyValueSource = "SALESFORCE"
	PropertyValueSourceSequences                   PropertyValueSource = "SEQUENCES"
	PropertyValueSourceSettings                    PropertyValueSource = "SETTINGS"
	PropertyValueSourceSidekick                    PropertyValueSource = "SIDEKICK"
	PropertyValueSourceSignals                     PropertyValueSource = "SIGNALS"
	PropertyValueSourceSlackIntegration            PropertyValueSource = "SLACK_INTEGRATION"
	PropertyValueSourceSmartDataCapture            PropertyValueSource = "SMART_DATA_CAPTURE"
	PropertyValueSourceSocial                      PropertyValueSource = "SOCIAL"
	PropertyValueSourceSuccess                     PropertyValueSource = "SUCCESS"
	PropertyValueSourceTally                       PropertyValueSource = "TALLY"
	PropertyValueSourceTask                        PropertyValueSource = "TASK"
	PropertyValueSourceUnknown                     PropertyValueSource = "UNKNOWN"
	PropertyValueSourceWalIncremental              PropertyValueSource = "WAL_INCREMENTAL"
	PropertyValueSourceWorkUi                      PropertyValueSource = "WORK_UI"
	PropertyValueSourceWorkflowContactDeleteAction PropertyValueSource = "WORKFLOW_CONTACT_DELETE_ACTION"
	PropertyValueSourceWorkflows                   PropertyValueSource = "WORKFLOWS"
)

// Represents a single custom property of a marketing event, storing its name,
// value, metadata (like source, timestamp, and sensitivity), and related audit
// information for tracking changes.
//
// The properties DataSensitivity, IsEncrypted, IsLargeValue, Name,
// PersistenceTimestamp, RequestID, SelectedByUser, SelectedByUserTimestamp,
// Source, SourceID, SourceLabel, SourceMetadata, SourceUpstreamDeployable,
// SourceVid, Timestamp, Unit, UpdatedByUserID, UseTimestampAsPersistenceTimestamp,
// Value are required.
type PropertyValueParam struct {
	// The sensitivity level of the property, such as "non_sensitive", "sensitive", and
	// "highly_sensitive".
	//
	// Any of "high", "none", "standard".
	DataSensitivity PropertyValueDataSensitivity `json:"dataSensitivity,omitzero" api:"required"`
	// Whether the property value is encrypted.
	IsEncrypted bool `json:"isEncrypted" api:"required"`
	// Indicates if the value exceeds normal size limits.
	IsLargeValue bool `json:"isLargeValue" api:"required"`
	// The unique property name.
	Name string `json:"name" api:"required"`
	// When the value was persisted to database, in epoch milliseconds.
	PersistenceTimestamp int64 `json:"persistenceTimestamp" api:"required"`
	// A unique ID associated with this request.
	RequestID string `json:"requestId" api:"required"`
	// Whether the value was selected by a user.
	SelectedByUser bool `json:"selectedByUser" api:"required"`
	// The timestamp when the value was selected by a user, if applicable.
	SelectedByUserTimestamp int64 `json:"selectedByUserTimestamp" api:"required"`
	// The origin of the property value, such as "IMPORT" or "API".
	//
	// Any of "ACADEMY", "ACCEPTANCE_TEST", "ACTIVITY_AUTO_ASSOCIATE",
	// "ACTIVITY_LOG_REVERT", "ADS", "AI_GROUP", "ANALYTICS", "API", "APPROVALS",
	// "ASSISTS", "ASSOCIATIONS", "AUTO_ASSOCIATE_BY_DOMAIN", "AUTOMATION_JOURNEY",
	// "AUTOMATION_PLATFORM", "AVATARS_SERVICE", "BATCH_UPDATE", "BCC_TO_CRM",
	// "BEHAVIORAL_EVENTS", "BET_ASSIGNMENT", "BET_CRM_CONNECTOR", "BIDEN", "BILLING",
	// "BOT", "CALCULATED", "CENTRAL_EXCHANGE_RATES", "CHATSPOT", "CLONE_OBJECTS",
	// "COMMUNICATOR", "COMPANIES", "COMPANY_FAMILIES", "COMPANY_INSIGHTS",
	// "CONNECTED_ACCOUNT", "CONTACTS", "CONTACTS_WEB", "CONTENT_MEMBERSHIP",
	// "CONVERSATIONAL_ENRICHMENT", "CONVERSATIONS", "CRM_PROCESSES_PLATFORM",
	// "CRM_UI", "CRM_UI_BULK_ACTION", "CUSTOMER_AGENT", "DATA_ENRICHMENT",
	// "DATA_QUALITY", "DATASET", "DEALS", "DEFAULT", "DELETE_OBJECTS", "EMAIL",
	// "EMAIL_INBOX_IMPORT", "EMAIL_INTEGRATION", "ENGAGEMENTS", "EXTENSION",
	// "FILE_MANAGER", "FLYWHEEL_PRODUCT_DATA_SYNC", "FORECASTING", "FORM",
	// "FORWARD_TO_CRM", "GMAIL_INTEGRATION", "GOALS", "HEISENBERG", "HELP_DESK",
	// "HELP_DESK_AI", "IMPORT", "INTEGRATION", "INTEGRATIONS_PLATFORM",
	// "INTEGRATIONS_SYNC", "INTENT", "INTERNAL_PROCESSING", "LEADIN",
	// "LEGAL_BASIS_REMEDIATION", "MARKET_SOURCING", "MARKETPLACE", "MARKETS",
	// "MEETINGS", "MERGE_COMPANIES", "MERGE_CONTACTS", "MERGE_OBJECTS",
	// "MERGE_REVERT_OBJECTS", "MICROAPPS", "MIGRATION", "MOBILE_ANDROID",
	// "MOBILE_IOS", "PAYMENTS", "PIPELINE_SETTINGS", "PLAYBOOKS",
	// "PORTAL_OBJECT_SYNC", "PORTAL_USER_ASSOCIATOR", "PRESENTATIONS",
	// "PRIMARY_AUTOMATION", "PROPERTY_DEFAULT_VALUE", "PROPERTY_RESTORE",
	// "PROPERTY_SETTINGS", "PROSPECTING_AGENT", "QUOTAS", "QUOTES", "RECYCLING_BIN",
	// "RESTORE_OBJECTS", "REVENUE_PLATFORM", "SALES", "SALES_MESSAGES", "SALESFORCE",
	// "SEQUENCES", "SETTINGS", "SIDEKICK", "SIGNALS", "SLACK_INTEGRATION",
	// "SMART_DATA_CAPTURE", "SOCIAL", "SUCCESS", "TALLY", "TASK", "UNKNOWN",
	// "WAL_INCREMENTAL", "WORK_UI", "WORKFLOW_CONTACT_DELETE_ACTION", "WORKFLOWS".
	Source PropertyValueSource `json:"source,omitzero" api:"required"`
	// The ID of the property source indicating where it was created.
	SourceID string `json:"sourceId" api:"required"`
	// A human-readable label.
	SourceLabel string `json:"sourceLabel" api:"required"`
	// Metadata providing additional context about the source.
	SourceMetadata           string `json:"sourceMetadata" api:"required"`
	SourceUpstreamDeployable string `json:"sourceUpstreamDeployable" api:"required"`
	// The unique identifier associated with the source.
	SourceVid []int64 `json:"sourceVid,omitzero" api:"required"`
	// When the value was set, as a 64-bit integer.
	Timestamp int64 `json:"timestamp" api:"required"`
	// The unit of measurement or context for the value.
	Unit string `json:"unit" api:"required"`
	// The ID of the user who updated the property.
	UpdatedByUserID int64 `json:"updatedByUserId" api:"required"`
	// Flag indicating whether to use the timestamp field as the persistence timestamp.
	UseTimestampAsPersistenceTimestamp bool `json:"useTimestampAsPersistenceTimestamp" api:"required"`
	// The property value.
	Value string `json:"value" api:"required"`
	paramObj
}

func (r PropertyValueParam) MarshalJSON() (data []byte, err error) {
	type shadow PropertyValueParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PropertyValueParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Contains the Id of a Public Object
type PublicObjectID struct {
	// The unique identifier for the public object.
	ID string `json:"id" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PublicObjectID) RawJSON() string { return r.JSON.raw }
func (r *PublicObjectID) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this PublicObjectID to a PublicObjectIDParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// PublicObjectIDParam.Overrides()
func (r PublicObjectID) ToParam() PublicObjectIDParam {
	return param.Override[PublicObjectIDParam](json.RawMessage(r.RawJSON()))
}

// Contains the Id of a Public Object
//
// The property ID is required.
type PublicObjectIDParam struct {
	// The unique identifier for the public object.
	ID string `json:"id" api:"required"`
	paramObj
}

func (r PublicObjectIDParam) MarshalJSON() (data []byte, err error) {
	type shadow PublicObjectIDParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PublicObjectIDParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Ye olde error
type StandardError struct {
	// Error category.
	Category string `json:"category" api:"required"`
	// Error context.
	Context map[string][]string `json:"context" api:"required"`
	// List of error details.
	Errors []ErrorDetail `json:"errors" api:"required"`
	// Error links.
	Links map[string]string `json:"links" api:"required"`
	// Error message.
	Message string `json:"message" api:"required"`
	// Error status.
	Status string `json:"status" api:"required"`
	// Error ID.
	ID string `json:"id"`
	// Error subcategory.
	SubCategory any `json:"subCategory"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Category    respjson.Field
		Context     respjson.Field
		Errors      respjson.Field
		Links       respjson.Field
		Message     respjson.Field
		Status      respjson.Field
		ID          respjson.Field
		SubCategory respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r StandardError) RawJSON() string { return r.JSON.raw }
func (r *StandardError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TaskLocator struct {
	// The unique identifier for the task.
	ID string `json:"id" api:"required"`
	// A map of link names to associated URIs containing documentation about the error
	// or recommended remediation steps
	Links map[string]string `json:"links"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Links       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TaskLocator) RawJSON() string { return r.JSON.raw }
func (r *TaskLocator) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VersionUser struct {
	// The unique ID of the User.
	ID string `json:"id" api:"required"`
	// The email address of the user.
	Email string `json:"email" api:"required"`
	// The first and last name of the User.
	FullName string `json:"fullName" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Email       respjson.Field
		FullName    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VersionUser) RawJSON() string { return r.JSON.raw }
func (r *VersionUser) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
