// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package crm

import (
	"context"
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
	"github.com/stainless-sdks/hubspot-sdk-go/packages/param"
	"github.com/stainless-sdks/hubspot-sdk-go/packages/respjson"
	"github.com/stainless-sdks/hubspot-sdk-go/shared"
)

// ObjectSchemaService contains methods and other services that help with
// interacting with the hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewObjectSchemaService] method instead.
type ObjectSchemaService struct {
	options []option.RequestOption
	Batch   ObjectSchemaBatchService
}

// NewObjectSchemaService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewObjectSchemaService(opts ...option.RequestOption) (r ObjectSchemaService) {
	r = ObjectSchemaService{}
	r.options = opts
	r.Batch = NewObjectSchemaBatchService(opts...)
	return
}

// Create a new custom object schema by defining its properties and associations.
func (r *ObjectSchemaService) New(ctx context.Context, body ObjectSchemaNewParams, opts ...option.RequestOption) (res *ObjectSchema, err error) {
	opts = slices.Concat(r.options, opts)
	path := "crm-object-schemas/2026-03/schemas"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Update attributes of a custom object schema, such as properties and labels,
// using the object type ID or fully qualified name.
func (r *ObjectSchemaService) Update(ctx context.Context, objectType string, body ObjectSchemaUpdateParams, opts ...option.RequestOption) (res *shared.ObjectTypeDefinition, err error) {
	opts = slices.Concat(r.options, opts)
	if objectType == "" {
		err = errors.New("missing required objectType parameter")
		return nil, err
	}
	path := fmt.Sprintf("crm-object-schemas/2026-03/schemas/%s", url.PathEscape(objectType))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

// Retrieve all custom object schemas, with options to include property
// definitions, association definitions, and audit metadata.
func (r *ObjectSchemaService) List(ctx context.Context, query ObjectSchemaListParams, opts ...option.RequestOption) (res *CollectionResponseObjectSchemaNoPaging, err error) {
	opts = slices.Concat(r.options, opts)
	path := "crm-object-schemas/2026-03/schemas"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Remove a custom object schema from the account using its object type ID or fully
// qualified name.
func (r *ObjectSchemaService) Delete(ctx context.Context, objectType string, body ObjectSchemaDeleteParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if objectType == "" {
		err = errors.New("missing required objectType parameter")
		return err
	}
	path := fmt.Sprintf("crm-object-schemas/2026-03/schemas/%s", url.PathEscape(objectType))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, nil, opts...)
	return err
}

// Create a new association between the specified object type and another object
// type. This operation requires the definition of the association attributes, such
// as the primary and target object type IDs.
func (r *ObjectSchemaService) NewAssociation(ctx context.Context, objectType string, body ObjectSchemaNewAssociationParams, opts ...option.RequestOption) (res *shared.AssociationDefinition, err error) {
	opts = slices.Concat(r.options, opts)
	if objectType == "" {
		err = errors.New("missing required objectType parameter")
		return nil, err
	}
	path := fmt.Sprintf("crm-object-schemas/2026-03/schemas/%s/associations", url.PathEscape(objectType))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Remove an association between two object types identified by the association
// identifier and object type. This operation is irreversible and will permanently
// delete the specified association.
func (r *ObjectSchemaService) DeleteAssociation(ctx context.Context, associationIdentifier string, body ObjectSchemaDeleteAssociationParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if body.ObjectType == "" {
		err = errors.New("missing required objectType parameter")
		return err
	}
	if associationIdentifier == "" {
		err = errors.New("missing required associationIdentifier parameter")
		return err
	}
	path := fmt.Sprintf("crm-object-schemas/2026-03/schemas/%s/associations/%s", url.PathEscape(body.ObjectType), url.PathEscape(associationIdentifier))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Retrieve details of a custom object schema, including its properties and
// associations, using the object type ID or fully qualified name.
func (r *ObjectSchemaService) Get(ctx context.Context, objectType string, query ObjectSchemaGetParams, opts ...option.RequestOption) (res *ObjectSchema, err error) {
	opts = slices.Concat(r.options, opts)
	if objectType == "" {
		err = errors.New("missing required objectType parameter")
		return nil, err
	}
	path := fmt.Sprintf("crm-object-schemas/2026-03/schemas/%s", url.PathEscape(objectType))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

type CollectionResponseObjectSchemaNoPaging struct {
	Results []ObjectSchema `json:"results" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Results     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CollectionResponseObjectSchemaNoPaging) RawJSON() string { return r.JSON.raw }
func (r *CollectionResponseObjectSchemaNoPaging) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ObjectSchema struct {
	// A unique ID for this schema's object type. Will be defined as
	// {meta-type}-{unique ID}.
	ID                        string `json:"id" api:"required"`
	AllowsSensitiveProperties bool   `json:"allowsSensitiveProperties" api:"required"`
	Archived                  bool   `json:"archived" api:"required"`
	// Associations defined for a given object type.
	Associations []shared.AssociationDefinition `json:"associations" api:"required"`
	// An assigned unique ID for the object, including portal ID and object name.
	FullyQualifiedName string                            `json:"fullyQualifiedName" api:"required"`
	Labels             shared.ObjectTypeDefinitionLabels `json:"labels" api:"required"`
	// A unique name for the schema's object type.
	Name         string `json:"name" api:"required"`
	ObjectTypeID string `json:"objectTypeId" api:"required"`
	// Properties defined for this object type.
	Properties []shared.Property `json:"properties" api:"required"`
	// The names of properties that should be **required** when creating an object of
	// this type.
	RequiredProperties []string `json:"requiredProperties" api:"required"`
	// Names of properties that will be indexed for this object type in by HubSpot's
	// product search.
	SearchableProperties []string `json:"searchableProperties" api:"required"`
	// The names of secondary properties for this object. These will be displayed as
	// secondary on the HubSpot record page for this object type.
	SecondaryDisplayProperties []string `json:"secondaryDisplayProperties" api:"required"`
	// When the object schema was created.
	CreatedAt       time.Time `json:"createdAt" format:"date-time"`
	CreatedByUserID int64     `json:"createdByUserId"`
	Description     string    `json:"description"`
	// The name of the primary property for this object. This will be displayed as
	// primary on the HubSpot record page for this object type.
	PrimaryDisplayProperty string `json:"primaryDisplayProperty"`
	// When the object schema was last updated.
	UpdatedAt       time.Time `json:"updatedAt" format:"date-time"`
	UpdatedByUserID int64     `json:"updatedByUserId"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                         respjson.Field
		AllowsSensitiveProperties  respjson.Field
		Archived                   respjson.Field
		Associations               respjson.Field
		FullyQualifiedName         respjson.Field
		Labels                     respjson.Field
		Name                       respjson.Field
		ObjectTypeID               respjson.Field
		Properties                 respjson.Field
		RequiredProperties         respjson.Field
		SearchableProperties       respjson.Field
		SecondaryDisplayProperties respjson.Field
		CreatedAt                  respjson.Field
		CreatedByUserID            respjson.Field
		Description                respjson.Field
		PrimaryDisplayProperty     respjson.Field
		UpdatedAt                  respjson.Field
		UpdatedByUserID            respjson.Field
		ExtraFields                map[string]respjson.Field
		raw                        string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ObjectSchema) RawJSON() string { return r.JSON.raw }
func (r *ObjectSchema) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties IncludeAssociationDefinitions, IncludeAuditMetadata,
// IncludePropertyDefinitions, Inputs are required.
type ObjectSchemaBatchReadRequestParam struct {
	// Indicates whether to include association definitions in the response.
	IncludeAssociationDefinitions bool `json:"includeAssociationDefinitions" api:"required"`
	// Indicates whether to include audit metadata in the response.
	IncludeAuditMetadata bool `json:"includeAuditMetadata" api:"required"`
	// Indicates whether to include property definitions in the response.
	IncludePropertyDefinitions bool     `json:"includePropertyDefinitions" api:"required"`
	Inputs                     []string `json:"inputs,omitzero" api:"required"`
	paramObj
}

func (r ObjectSchemaBatchReadRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow ObjectSchemaBatchReadRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ObjectSchemaBatchReadRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties AllowsSensitiveProperties, AssociatedObjects, Labels, Name,
// Properties, RequiredProperties, SearchableProperties, SecondaryDisplayProperties
// are required.
type ObjectSchemaEggParam struct {
	// Determines if the object type can include properties that are marked as
	// sensitive.
	AllowsSensitiveProperties bool `json:"allowsSensitiveProperties" api:"required"`
	// Associations defined for this object type.
	AssociatedObjects []string                               `json:"associatedObjects,omitzero" api:"required"`
	Labels            shared.ObjectTypeDefinitionLabelsParam `json:"labels,omitzero" api:"required"`
	// A unique name for this object. For internal use only.
	Name string `json:"name" api:"required"`
	// Properties defined for this object type.
	Properties []ObjectTypePropertyCreateParam `json:"properties,omitzero" api:"required"`
	// The names of properties that should be **required** when creating an object of
	// this type.
	RequiredProperties []string `json:"requiredProperties,omitzero" api:"required"`
	// Names of properties that will be indexed for this object type in by HubSpot's
	// product search.
	SearchableProperties []string `json:"searchableProperties,omitzero" api:"required"`
	// The names of secondary properties for this object. These will be displayed as
	// secondary on the HubSpot record page for this object type.
	SecondaryDisplayProperties []string `json:"secondaryDisplayProperties,omitzero" api:"required"`
	// A brief explanation of the object type.
	Description param.Opt[string] `json:"description,omitzero"`
	// The name of the primary property for this object. This will be displayed as
	// primary on the HubSpot record page for this object type.
	PrimaryDisplayProperty param.Opt[string] `json:"primaryDisplayProperty,omitzero"`
	paramObj
}

func (r ObjectSchemaEggParam) MarshalJSON() (data []byte, err error) {
	type shadow ObjectSchemaEggParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ObjectSchemaEggParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties FieldType, Label, Name, Type are required.
type ObjectTypePropertyCreateParam struct {
	// Controls how the property appears in HubSpot.
	FieldType string `json:"fieldType" api:"required"`
	// A human-readable property label that will be shown in HubSpot.
	Label string `json:"label" api:"required"`
	// The internal property name, which must be used when referencing the property
	// from the API.
	Name string `json:"name" api:"required"`
	// The data type of the property.
	//
	// Any of "bool", "date", "datetime", "enumeration", "number", "phone_number",
	// "string".
	Type ObjectTypePropertyCreateType `json:"type,omitzero" api:"required"`
	// A description of the property that will be shown as help text in HubSpot.
	Description param.Opt[string] `json:"description,omitzero"`
	// The order that this property should be displayed in the HubSpot UI relative to
	// other properties for this object type. Properties are displayed in order
	// starting with the lowest positive integer value. A value of -1 will cause the
	// property to be displayed **after** any positive values.
	DisplayOrder param.Opt[int64] `json:"displayOrder,omitzero"`
	// Specifies the reference type for external options associated with the property.
	ExternalOptionsReferenceType param.Opt[string] `json:"externalOptionsReferenceType,omitzero"`
	// Whether the property can be used in a HubSpot form.
	FormField param.Opt[bool] `json:"formField,omitzero"`
	// The name of the group this property belongs to.
	GroupName param.Opt[string] `json:"groupName,omitzero"`
	// Whether or not the property's value must be unique. Once set, this can't be
	// changed.
	HasUniqueValue param.Opt[bool] `json:"hasUniqueValue,omitzero"`
	// Hidden options won't be shown in HubSpot.
	Hidden param.Opt[bool] `json:"hidden,omitzero"`
	// Defines the options this property will return, e.g. OWNER would return name of
	// users on the portal.
	ReferencedObjectType param.Opt[string] `json:"referencedObjectType,omitzero"`
	// Allow users to search for information entered to this field (limited to 3
	// properties)
	SearchableInGlobalSearch param.Opt[bool] `json:"searchableInGlobalSearch,omitzero"`
	// Whether the property will display the currency symbol in the HubSpot UI.
	ShowCurrencySymbol param.Opt[bool] `json:"showCurrencySymbol,omitzero"`
	// Controls how numeric properties are formatted in the HubSpot UI
	//
	// Any of "currency", "duration", "formatted", "percentage", "probability",
	// "unformatted".
	NumberDisplayHint ObjectTypePropertyCreateNumberDisplayHint `json:"numberDisplayHint,omitzero"`
	// A list of available options for the property. This field is only required for
	// enumerated properties.
	Options []shared.OptionInputParam `json:"options,omitzero"`
	// Controls how the property options will be sorted in the HubSpot UI.
	//
	// Any of "ALPHABETICAL", "DISPLAY_ORDER".
	OptionSortStrategy ObjectTypePropertyCreateOptionSortStrategy `json:"optionSortStrategy,omitzero"`
	// Controls how text properties are formatted in the HubSpot UI
	//
	// Any of "domain_name", "email", "ip_address", "multi_line", "phone_number",
	// "physical_address", "postal_code", "unformatted_single_line".
	TextDisplayHint ObjectTypePropertyCreateTextDisplayHint `json:"textDisplayHint,omitzero"`
	paramObj
}

func (r ObjectTypePropertyCreateParam) MarshalJSON() (data []byte, err error) {
	type shadow ObjectTypePropertyCreateParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ObjectTypePropertyCreateParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The data type of the property.
type ObjectTypePropertyCreateType string

const (
	ObjectTypePropertyCreateTypeBool        ObjectTypePropertyCreateType = "bool"
	ObjectTypePropertyCreateTypeDate        ObjectTypePropertyCreateType = "date"
	ObjectTypePropertyCreateTypeDatetime    ObjectTypePropertyCreateType = "datetime"
	ObjectTypePropertyCreateTypeEnumeration ObjectTypePropertyCreateType = "enumeration"
	ObjectTypePropertyCreateTypeNumber      ObjectTypePropertyCreateType = "number"
	ObjectTypePropertyCreateTypePhoneNumber ObjectTypePropertyCreateType = "phone_number"
	ObjectTypePropertyCreateTypeString      ObjectTypePropertyCreateType = "string"
)

// Controls how numeric properties are formatted in the HubSpot UI
type ObjectTypePropertyCreateNumberDisplayHint string

const (
	ObjectTypePropertyCreateNumberDisplayHintCurrency    ObjectTypePropertyCreateNumberDisplayHint = "currency"
	ObjectTypePropertyCreateNumberDisplayHintDuration    ObjectTypePropertyCreateNumberDisplayHint = "duration"
	ObjectTypePropertyCreateNumberDisplayHintFormatted   ObjectTypePropertyCreateNumberDisplayHint = "formatted"
	ObjectTypePropertyCreateNumberDisplayHintPercentage  ObjectTypePropertyCreateNumberDisplayHint = "percentage"
	ObjectTypePropertyCreateNumberDisplayHintProbability ObjectTypePropertyCreateNumberDisplayHint = "probability"
	ObjectTypePropertyCreateNumberDisplayHintUnformatted ObjectTypePropertyCreateNumberDisplayHint = "unformatted"
)

// Controls how the property options will be sorted in the HubSpot UI.
type ObjectTypePropertyCreateOptionSortStrategy string

const (
	ObjectTypePropertyCreateOptionSortStrategyAlphabetical ObjectTypePropertyCreateOptionSortStrategy = "ALPHABETICAL"
	ObjectTypePropertyCreateOptionSortStrategyDisplayOrder ObjectTypePropertyCreateOptionSortStrategy = "DISPLAY_ORDER"
)

// Controls how text properties are formatted in the HubSpot UI
type ObjectTypePropertyCreateTextDisplayHint string

const (
	ObjectTypePropertyCreateTextDisplayHintDomainName            ObjectTypePropertyCreateTextDisplayHint = "domain_name"
	ObjectTypePropertyCreateTextDisplayHintEmail                 ObjectTypePropertyCreateTextDisplayHint = "email"
	ObjectTypePropertyCreateTextDisplayHintIPAddress             ObjectTypePropertyCreateTextDisplayHint = "ip_address"
	ObjectTypePropertyCreateTextDisplayHintMultiLine             ObjectTypePropertyCreateTextDisplayHint = "multi_line"
	ObjectTypePropertyCreateTextDisplayHintPhoneNumber           ObjectTypePropertyCreateTextDisplayHint = "phone_number"
	ObjectTypePropertyCreateTextDisplayHintPhysicalAddress       ObjectTypePropertyCreateTextDisplayHint = "physical_address"
	ObjectTypePropertyCreateTextDisplayHintPostalCode            ObjectTypePropertyCreateTextDisplayHint = "postal_code"
	ObjectTypePropertyCreateTextDisplayHintUnformattedSingleLine ObjectTypePropertyCreateTextDisplayHint = "unformatted_single_line"
)

type ObjectSchemaNewParams struct {
	ObjectSchemaEgg ObjectSchemaEggParam
	paramObj
}

func (r ObjectSchemaNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.ObjectSchemaEgg)
}
func (r *ObjectSchemaNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ObjectSchemaUpdateParams struct {
	ObjectTypeDefinitionPatch shared.ObjectTypeDefinitionPatchParam
	paramObj
}

func (r ObjectSchemaUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.ObjectTypeDefinitionPatch)
}
func (r *ObjectSchemaUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ObjectSchemaListParams struct {
	// Whether to return only results that have been archived.
	Archived                      param.Opt[bool] `query:"archived,omitzero" json:"-"`
	IncludeAssociationDefinitions param.Opt[bool] `query:"includeAssociationDefinitions,omitzero" json:"-"`
	IncludeAuditMetadata          param.Opt[bool] `query:"includeAuditMetadata,omitzero" json:"-"`
	IncludePropertyDefinitions    param.Opt[bool] `query:"includePropertyDefinitions,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ObjectSchemaListParams]'s query parameters as `url.Values`.
func (r ObjectSchemaListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ObjectSchemaDeleteParams struct {
	// Whether to return only results that have been archived.
	Archived param.Opt[bool] `query:"archived,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ObjectSchemaDeleteParams]'s query parameters as
// `url.Values`.
func (r ObjectSchemaDeleteParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ObjectSchemaNewAssociationParams struct {
	AssociationDefinitionEgg shared.AssociationDefinitionEggParam
	paramObj
}

func (r ObjectSchemaNewAssociationParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.AssociationDefinitionEgg)
}
func (r *ObjectSchemaNewAssociationParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ObjectSchemaDeleteAssociationParams struct {
	ObjectType string `path:"objectType" api:"required" json:"-"`
	paramObj
}

type ObjectSchemaGetParams struct {
	IncludeAssociationDefinitions param.Opt[bool] `query:"includeAssociationDefinitions,omitzero" json:"-"`
	IncludeAuditMetadata          param.Opt[bool] `query:"includeAuditMetadata,omitzero" json:"-"`
	IncludePropertyDefinitions    param.Opt[bool] `query:"includePropertyDefinitions,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ObjectSchemaGetParams]'s query parameters as `url.Values`.
func (r ObjectSchemaGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
