// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package crm

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
	"github.com/stainless-sdks/hubspot-sdk-go/packages/param"
	"github.com/stainless-sdks/hubspot-sdk-go/packages/respjson"
	"github.com/stainless-sdks/hubspot-sdk-go/shared"
)

// ObjectSchemaService contains methods and other services that help with
// interacting with the Hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewObjectSchemaService] method instead.
type ObjectSchemaService struct {
	Options []option.RequestOption
}

// NewObjectSchemaService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewObjectSchemaService(opts ...option.RequestOption) (r ObjectSchemaService) {
	r = ObjectSchemaService{}
	r.Options = opts
	return
}

// Define a new object schema, along with custom properties and associations. The
// entire object schema, including its object type ID, properties, and associations
// will be returned in the response.
func (r *ObjectSchemaService) New(ctx context.Context, body ObjectSchemaNewParams, opts ...option.RequestOption) (res *ObjectSchema, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "crm-object-schemas/v3/schemas"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Update the details for an existing object schema.
func (r *ObjectSchemaService) Update(ctx context.Context, objectType string, body ObjectSchemaUpdateParams, opts ...option.RequestOption) (res *ObjectTypeDefinition, err error) {
	opts = slices.Concat(r.Options, opts)
	if objectType == "" {
		err = errors.New("missing required objectType parameter")
		return
	}
	path := fmt.Sprintf("crm-object-schemas/v3/schemas/%s", objectType)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Returns all object schemas that have been defined for your account.
func (r *ObjectSchemaService) List(ctx context.Context, query ObjectSchemaListParams, opts ...option.RequestOption) (res *shared.CollectionResponseObjectSchemaNoPaging, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "crm-object-schemas/v3/schemas"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Deletes a schema. Any existing records of this schema must be deleted **first**.
// Otherwise this call will fail.
func (r *ObjectSchemaService) Delete(ctx context.Context, objectType string, body ObjectSchemaDeleteParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if objectType == "" {
		err = errors.New("missing required objectType parameter")
		return
	}
	path := fmt.Sprintf("crm-object-schemas/v3/schemas/%s", objectType)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, nil, opts...)
	return
}

// Defines a new association between the primary schema's object type and other
// object types.
func (r *ObjectSchemaService) NewAssociation(ctx context.Context, objectType string, body ObjectSchemaNewAssociationParams, opts ...option.RequestOption) (res *ObjectSchemaNewAssociationResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if objectType == "" {
		err = errors.New("missing required objectType parameter")
		return
	}
	path := fmt.Sprintf("crm-object-schemas/v3/schemas/%s/associations", objectType)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Removes an existing association from a schema.
func (r *ObjectSchemaService) DeleteAssociation(ctx context.Context, associationIdentifier string, body ObjectSchemaDeleteAssociationParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if body.ObjectType == "" {
		err = errors.New("missing required objectType parameter")
		return
	}
	if associationIdentifier == "" {
		err = errors.New("missing required associationIdentifier parameter")
		return
	}
	path := fmt.Sprintf("crm-object-schemas/v3/schemas/%s/associations/%s", body.ObjectType, associationIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Returns an existing object schema.
func (r *ObjectSchemaService) Get(ctx context.Context, objectType string, opts ...option.RequestOption) (res *ObjectSchema, err error) {
	opts = slices.Concat(r.Options, opts)
	if objectType == "" {
		err = errors.New("missing required objectType parameter")
		return
	}
	path := fmt.Sprintf("crm-object-schemas/v3/schemas/%s", objectType)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Defines an object schema, including its properties and associations.
type ObjectSchema struct {
	// A unique ID for this schema's object type. Will be defined as
	// {meta-type}-{unique ID}.
	ID string `json:"id,required"`
	// Associations defined for a given object type.
	Associations []ObjectSchemaAssociation         `json:"associations,required"`
	Labels       shared.ObjectTypeDefinitionLabels `json:"labels,required"`
	// A unique name for the schema's object type.
	Name string `json:"name,required"`
	// Properties defined for this object type.
	Properties []shared.Property `json:"properties,required"`
	// The names of properties that should be **required** when creating an object of
	// this type.
	RequiredProperties []string `json:"requiredProperties,required"`
	Archived           bool     `json:"archived"`
	// When the object schema was created.
	CreatedAt       time.Time `json:"createdAt" format:"date-time"`
	CreatedByUserID int64     `json:"createdByUserId"`
	Description     string    `json:"description"`
	// An assigned unique ID for the object, including portal ID and object name.
	FullyQualifiedName string `json:"fullyQualifiedName"`
	ObjectTypeID       string `json:"objectTypeId"`
	// The name of the primary property for this object. This will be displayed as
	// primary on the HubSpot record page for this object type.
	PrimaryDisplayProperty string `json:"primaryDisplayProperty"`
	// Names of properties that will be indexed for this object type in by HubSpot's
	// product search.
	SearchableProperties []string `json:"searchableProperties"`
	// The names of secondary properties for this object. These will be displayed as
	// secondary on the HubSpot record page for this object type.
	SecondaryDisplayProperties []string `json:"secondaryDisplayProperties"`
	// When the object schema was last updated.
	UpdatedAt       time.Time `json:"updatedAt" format:"date-time"`
	UpdatedByUserID int64     `json:"updatedByUserId"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                         respjson.Field
		Associations               respjson.Field
		Labels                     respjson.Field
		Name                       respjson.Field
		Properties                 respjson.Field
		RequiredProperties         respjson.Field
		Archived                   respjson.Field
		CreatedAt                  respjson.Field
		CreatedByUserID            respjson.Field
		Description                respjson.Field
		FullyQualifiedName         respjson.Field
		ObjectTypeID               respjson.Field
		PrimaryDisplayProperty     respjson.Field
		SearchableProperties       respjson.Field
		SecondaryDisplayProperties respjson.Field
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

// The definition of an association
type ObjectSchemaAssociation struct {
	// The unique ID of the associated object (e.g., a contact ID).
	ID int64 `json:"id,required"`
	// Whether custom labels can be used in the association.
	AllowsCustomLabels bool `json:"allowsCustomLabels,required"`
	// The cardinality from the source object's perspective, either "ONE_TO_ONE" or
	// "ONE_TO_MANY".
	//
	// Any of "ONE_TO_ONE", "ONE_TO_MANY".
	Cardinality string `json:"cardinality,required"`
	// The category of the association. Can be: "HUBSPOT_DEFINED", "USER_DEFINED", or
	// "INTEGRATOR_DEFINED"
	//
	// Any of "HUBSPOT_DEFINED", "USER_DEFINED", "INTEGRATOR_DEFINED".
	Category string `json:"category,required"`
	// The ID of the source object type (e.g., 0-1 for contacts).
	FromObjectTypeID string `json:"fromObjectTypeId,required"`
	// Whether all potential linked objects are included in the association
	HasAllAssociatedObjects bool `json:"hasAllAssociatedObjects,required"`
	// Whether deletions in the association should cause cascading deletes to linked
	// objects.
	HasCascadingDeletes bool `json:"hasCascadingDeletes,required"`
	// Whether a user has set a limit for the number of source objects.
	HasUserEnforcedMaxFromObjectIDs bool `json:"hasUserEnforcedMaxFromObjectIds,required"`
	// Whether a user has set a limit for the number of destination objects.
	HasUserEnforcedMaxToObjectIDs bool `json:"hasUserEnforcedMaxToObjectIds,required"`
	// Whether the association is hidden or not.
	Hidden bool `json:"hidden,required"`
	// Whether the reverse association can also support custom labels.
	InverseAllowsCustomLabels bool `json:"inverseAllowsCustomLabels,required"`
	// The cardinality from the destination object's perspective, either "ONE_TO_ONE"
	// or "ONE_TO_MANY".
	//
	// Any of "ONE_TO_ONE", "ONE_TO_MANY".
	InverseCardinality string `json:"inverseCardinality,required"`
	// Whether all potential reverse linked objects are included in the association.
	InverseHasAllAssociatedObjects bool `json:"inverseHasAllAssociatedObjects,required"`
	// The unique ID for the inverse side of the association.
	InverseID int64 `json:"inverseId,required"`
	// The name used to describe the inverse relationship in this association
	InverseName string `json:"inverseName,required"`
	// Whether the inverse association is considered primary.
	IsInversePrimary bool `json:"isInversePrimary,required"`
	// Whether the association is the primary link between the entities involved.
	IsPrimary bool `json:"isPrimary,required"`
	// The maximum number of source object IDs allowed in the association.
	MaxFromObjectIDs int64 `json:"maxFromObjectIds,required"`
	// The maximum number of destination object IDs allowed in the association.
	MaxToObjectIDs int64 `json:"maxToObjectIds,required"`
	// For labeled association types, the internal name of the association.
	Name string `json:"name,required"`
	// A unique across-portal ID applied to the association.
	PortalUniqueIdentifier string `json:"portalUniqueIdentifier,required"`
	// The ID of the destination object type (e.g., 0-3 for deals).
	ToObjectTypeID string `json:"toObjectTypeId,required"`
	// The name of the source object type (e.g,. "DEAL" or "QUOTE").
	//
	// Any of "CONTACT", "COMPANY", "DEAL", "ENGAGEMENT", "TICKET", "OWNER", "PRODUCT",
	// "LINE_ITEM", "BET_DELIVERABLE_SERVICE", "CONTENT", "CONVERSATION", "BET_ALERT",
	// "PORTAL", "QUOTE", "FORM_SUBMISSION_INBOUNDDB", "QUOTA", "UNSUBSCRIBE",
	// "COMMUNICATION", "FEEDBACK_SUBMISSION", "ATTRIBUTION", "SALESFORCE_SYNC_ERROR",
	// "RESTORABLE_CRM_OBJECT", "HUB", "LANDING_PAGE", "PRODUCT_OR_FOLDER", "TASK",
	// "FORM", "MARKETING_EMAIL", "AD_ACCOUNT", "AD_CAMPAIGN", "AD_GROUP", "AD",
	// "KEYWORD", "CAMPAIGN", "SOCIAL_CHANNEL", "SOCIAL_POST", "SITE_PAGE",
	// "BLOG_POST", "IMPORT", "EXPORT", "CTA", "TASK_TEMPLATE",
	// "AUTOMATION_PLATFORM_FLOW", "OBJECT_LIST", "NOTE", "MEETING_EVENT", "CALL",
	// "EMAIL", "PUBLISHING_TASK", "CONVERSATION_SESSION",
	// "CONTACT_CREATE_ATTRIBUTION", "INVOICE", "MARKETING_EVENT",
	// "CONVERSATION_INBOX", "CHATFLOW", "MEDIA_BRIDGE", "SEQUENCE", "SEQUENCE_STEP",
	// "FORECAST", "SNIPPET", "TEMPLATE", "DEAL_CREATE_ATTRIBUTION", "QUOTE_TEMPLATE",
	// "QUOTE_MODULE", "QUOTE_MODULE_FIELD", "QUOTE_FIELD", "SEQUENCE_ENROLLMENT",
	// "SUBSCRIPTION", "ACCEPTANCE_TEST", "SOCIAL_BROADCAST", "DEAL_SPLIT",
	// "DEAL_REGISTRATION", "GOAL_TARGET", "GOAL_TARGET_GROUP",
	// "PORTAL_OBJECT_SYNC_MESSAGE", "FILE_MANAGER_FILE", "FILE_MANAGER_FOLDER",
	// "SEQUENCE_STEP_ENROLLMENT", "APPROVAL", "APPROVAL_STEP", "CTA_VARIANT",
	// "SALES_DOCUMENT", "DISCOUNT", "FEE", "TAX", "MARKETING_CALENDAR",
	// "PERMISSIONS_TESTING", "PRIVACY_SCANNER_COOKIE", "DATA_SYNC_STATE",
	// "WEB_INTERACTIVE", "PLAYBOOK", "FOLDER", "PLAYBOOK_QUESTION",
	// "PLAYBOOK_SUBMISSION", "PLAYBOOK_SUBMISSION_ANSWER", "COMMERCE_PAYMENT",
	// "GSC_PROPERTY", "SOX_PROTECTED_DUMMY_TYPE", "BLOG_LISTING_PAGE",
	// "QUARANTINED_SUBMISSION", "PAYMENT_SCHEDULE", "PAYMENT_SCHEDULE_INSTALLMENT",
	// "MARKETING_CAMPAIGN_UTM", "DISCOUNT_TEMPLATE", "DISCOUNT_CODE",
	// "FEEDBACK_SURVEY", "CMS_URL", "SALES_TASK", "SALES_WORKLOAD", "USER",
	// "POSTAL_MAIL", "SCHEMAS_BACKEND_TEST", "PAYMENT_LINK", "SUBMISSION_TAG",
	// "CAMPAIGN_STEP", "SCHEDULING_PAGE", "SOX_PROTECTED_TEST_TYPE", "ORDER",
	// "MARKETING_SMS", "PARTNER_ACCOUNT", "CAMPAIGN_TEMPLATE",
	// "CAMPAIGN_TEMPLATE_STEP", "PLAYLIST", "CLIP", "CAMPAIGN_BUDGET_ITEM",
	// "CAMPAIGN_SPEND_ITEM", "MIC", "CONTENT_AUDIT", "CONTENT_AUDIT_PAGE",
	// "PLAYLIST_FOLDER", "LEAD", "ABANDONED_CART", "EXTERNAL_WEB_URL", "VIEW",
	// "VIEW_BLOCK", "ROSTER", "CART", "AUTOMATION_PLATFORM_FLOW_ACTION",
	// "SOCIAL_PROFILE", "PARTNER_CLIENT", "ROSTER_MEMBER",
	// "MARKETING_EVENT_ATTENDANCE", "ALL_PAGES", "AI_FORECAST",
	// "CRM_PIPELINES_DUMMY_TYPE", "KNOWLEDGE_ARTICLE", "PROPERTY_INFO",
	// "DATA_PRIVACY_CONSENT", "GOAL_TEMPLATE", "SCORE_CONFIGURATION", "AUDIENCE",
	// "PARTNER_CLIENT_REVENUE", "AUTOMATION_JOURNEY", "COMBO_EVENT_CONFIGURATION",
	// "CRM_OBJECTS_DUMMY_TYPE", "CASE_STUDY", "SERVICE", "PODCAST_EPISODE",
	// "PARTNER_SERVICE", "UNKNOWN".
	FromObjectType string `json:"fromObjectType"`
	// The label used to describe the reverse relationship in an association.
	InverseLabel string `json:"inverseLabel"`
	// The label given to an association.
	Label string `json:"label"`
	// The name of the destination object type (e.g,. "DEAL" or "QUOTE").
	//
	// Any of "CONTACT", "COMPANY", "DEAL", "ENGAGEMENT", "TICKET", "OWNER", "PRODUCT",
	// "LINE_ITEM", "BET_DELIVERABLE_SERVICE", "CONTENT", "CONVERSATION", "BET_ALERT",
	// "PORTAL", "QUOTE", "FORM_SUBMISSION_INBOUNDDB", "QUOTA", "UNSUBSCRIBE",
	// "COMMUNICATION", "FEEDBACK_SUBMISSION", "ATTRIBUTION", "SALESFORCE_SYNC_ERROR",
	// "RESTORABLE_CRM_OBJECT", "HUB", "LANDING_PAGE", "PRODUCT_OR_FOLDER", "TASK",
	// "FORM", "MARKETING_EMAIL", "AD_ACCOUNT", "AD_CAMPAIGN", "AD_GROUP", "AD",
	// "KEYWORD", "CAMPAIGN", "SOCIAL_CHANNEL", "SOCIAL_POST", "SITE_PAGE",
	// "BLOG_POST", "IMPORT", "EXPORT", "CTA", "TASK_TEMPLATE",
	// "AUTOMATION_PLATFORM_FLOW", "OBJECT_LIST", "NOTE", "MEETING_EVENT", "CALL",
	// "EMAIL", "PUBLISHING_TASK", "CONVERSATION_SESSION",
	// "CONTACT_CREATE_ATTRIBUTION", "INVOICE", "MARKETING_EVENT",
	// "CONVERSATION_INBOX", "CHATFLOW", "MEDIA_BRIDGE", "SEQUENCE", "SEQUENCE_STEP",
	// "FORECAST", "SNIPPET", "TEMPLATE", "DEAL_CREATE_ATTRIBUTION", "QUOTE_TEMPLATE",
	// "QUOTE_MODULE", "QUOTE_MODULE_FIELD", "QUOTE_FIELD", "SEQUENCE_ENROLLMENT",
	// "SUBSCRIPTION", "ACCEPTANCE_TEST", "SOCIAL_BROADCAST", "DEAL_SPLIT",
	// "DEAL_REGISTRATION", "GOAL_TARGET", "GOAL_TARGET_GROUP",
	// "PORTAL_OBJECT_SYNC_MESSAGE", "FILE_MANAGER_FILE", "FILE_MANAGER_FOLDER",
	// "SEQUENCE_STEP_ENROLLMENT", "APPROVAL", "APPROVAL_STEP", "CTA_VARIANT",
	// "SALES_DOCUMENT", "DISCOUNT", "FEE", "TAX", "MARKETING_CALENDAR",
	// "PERMISSIONS_TESTING", "PRIVACY_SCANNER_COOKIE", "DATA_SYNC_STATE",
	// "WEB_INTERACTIVE", "PLAYBOOK", "FOLDER", "PLAYBOOK_QUESTION",
	// "PLAYBOOK_SUBMISSION", "PLAYBOOK_SUBMISSION_ANSWER", "COMMERCE_PAYMENT",
	// "GSC_PROPERTY", "SOX_PROTECTED_DUMMY_TYPE", "BLOG_LISTING_PAGE",
	// "QUARANTINED_SUBMISSION", "PAYMENT_SCHEDULE", "PAYMENT_SCHEDULE_INSTALLMENT",
	// "MARKETING_CAMPAIGN_UTM", "DISCOUNT_TEMPLATE", "DISCOUNT_CODE",
	// "FEEDBACK_SURVEY", "CMS_URL", "SALES_TASK", "SALES_WORKLOAD", "USER",
	// "POSTAL_MAIL", "SCHEMAS_BACKEND_TEST", "PAYMENT_LINK", "SUBMISSION_TAG",
	// "CAMPAIGN_STEP", "SCHEDULING_PAGE", "SOX_PROTECTED_TEST_TYPE", "ORDER",
	// "MARKETING_SMS", "PARTNER_ACCOUNT", "CAMPAIGN_TEMPLATE",
	// "CAMPAIGN_TEMPLATE_STEP", "PLAYLIST", "CLIP", "CAMPAIGN_BUDGET_ITEM",
	// "CAMPAIGN_SPEND_ITEM", "MIC", "CONTENT_AUDIT", "CONTENT_AUDIT_PAGE",
	// "PLAYLIST_FOLDER", "LEAD", "ABANDONED_CART", "EXTERNAL_WEB_URL", "VIEW",
	// "VIEW_BLOCK", "ROSTER", "CART", "AUTOMATION_PLATFORM_FLOW_ACTION",
	// "SOCIAL_PROFILE", "PARTNER_CLIENT", "ROSTER_MEMBER",
	// "MARKETING_EVENT_ATTENDANCE", "ALL_PAGES", "AI_FORECAST",
	// "CRM_PIPELINES_DUMMY_TYPE", "KNOWLEDGE_ARTICLE", "PROPERTY_INFO",
	// "DATA_PRIVACY_CONSENT", "GOAL_TEMPLATE", "SCORE_CONFIGURATION", "AUDIENCE",
	// "PARTNER_CLIENT_REVENUE", "AUTOMATION_JOURNEY", "COMBO_EVENT_CONFIGURATION",
	// "CRM_OBJECTS_DUMMY_TYPE", "CASE_STUDY", "SERVICE", "PODCAST_EPISODE",
	// "PARTNER_SERVICE", "UNKNOWN".
	ToObjectType string `json:"toObjectType"`
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
		IsInversePrimary                respjson.Field
		IsPrimary                       respjson.Field
		MaxFromObjectIDs                respjson.Field
		MaxToObjectIDs                  respjson.Field
		Name                            respjson.Field
		PortalUniqueIdentifier          respjson.Field
		ToObjectTypeID                  respjson.Field
		FromObjectType                  respjson.Field
		InverseLabel                    respjson.Field
		Label                           respjson.Field
		ToObjectType                    respjson.Field
		ExtraFields                     map[string]respjson.Field
		raw                             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ObjectSchemaAssociation) RawJSON() string { return r.JSON.raw }
func (r *ObjectSchemaAssociation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Defines a new object type, its properties, and associations.
//
// The properties AssociatedObjects, Labels, Name, Properties, RequiredProperties
// are required.
type ObjectSchemaEggParam struct {
	// Associations defined for this object type.
	AssociatedObjects []string                               `json:"associatedObjects,omitzero,required"`
	Labels            shared.ObjectTypeDefinitionLabelsParam `json:"labels,omitzero,required"`
	// A unique name for this object. For internal use only.
	Name string `json:"name,required"`
	// Properties defined for this object type.
	Properties []ObjectTypePropertyCreateParam `json:"properties,omitzero,required"`
	// The names of properties that should be **required** when creating an object of
	// this type.
	RequiredProperties []string          `json:"requiredProperties,omitzero,required"`
	Description        param.Opt[string] `json:"description,omitzero"`
	// The name of the primary property for this object. This will be displayed as
	// primary on the HubSpot record page for this object type.
	PrimaryDisplayProperty param.Opt[string] `json:"primaryDisplayProperty,omitzero"`
	// Names of properties that will be indexed for this object type in by HubSpot's
	// product search.
	SearchableProperties []string `json:"searchableProperties,omitzero"`
	// The names of secondary properties for this object. These will be displayed as
	// secondary on the HubSpot record page for this object type.
	SecondaryDisplayProperties []string `json:"secondaryDisplayProperties,omitzero"`
	paramObj
}

func (r ObjectSchemaEggParam) MarshalJSON() (data []byte, err error) {
	type shadow ObjectSchemaEggParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ObjectSchemaEggParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Defines an object type.
type ObjectTypeDefinition struct {
	// A unique ID for this object type. Will be defined as {meta-type}-{unique ID}.
	ID     string                            `json:"id,required"`
	Labels shared.ObjectTypeDefinitionLabels `json:"labels,required"`
	// A unique name for this object. For internal use only.
	Name string `json:"name,required"`
	// The names of properties that should be **required** when creating an object of
	// this type.
	RequiredProperties []string `json:"requiredProperties,required"`
	Archived           bool     `json:"archived"`
	// When the object type was created.
	CreatedAt          time.Time `json:"createdAt" format:"date-time"`
	Description        string    `json:"description"`
	FullyQualifiedName string    `json:"fullyQualifiedName"`
	ObjectTypeID       string    `json:"objectTypeId"`
	// The ID of the account that this object type is specific to.
	PortalID int64 `json:"portalId"`
	// The name of the primary property for this object. This will be displayed as
	// primary on the HubSpot record page for this object type.
	PrimaryDisplayProperty string `json:"primaryDisplayProperty"`
	// Names of properties that will be indexed for this object type in by HubSpot's
	// product search.
	SearchableProperties []string `json:"searchableProperties"`
	// The names of secondary properties for this object. These will be displayed as
	// secondary on the HubSpot record page for this object type.
	SecondaryDisplayProperties []string `json:"secondaryDisplayProperties"`
	// When the object type was last updated.
	UpdatedAt time.Time `json:"updatedAt" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                         respjson.Field
		Labels                     respjson.Field
		Name                       respjson.Field
		RequiredProperties         respjson.Field
		Archived                   respjson.Field
		CreatedAt                  respjson.Field
		Description                respjson.Field
		FullyQualifiedName         respjson.Field
		ObjectTypeID               respjson.Field
		PortalID                   respjson.Field
		PrimaryDisplayProperty     respjson.Field
		SearchableProperties       respjson.Field
		SecondaryDisplayProperties respjson.Field
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

// Defines attributes to update on an object type.
type ObjectTypeDefinitionPatchParam struct {
	ClearDescription param.Opt[bool]   `json:"clearDescription,omitzero"`
	Description      param.Opt[string] `json:"description,omitzero"`
	// The name of the primary property for this object. This will be displayed as
	// primary on the HubSpot record page for this object type.
	PrimaryDisplayProperty param.Opt[string]                      `json:"primaryDisplayProperty,omitzero"`
	Restorable             param.Opt[bool]                        `json:"restorable,omitzero"`
	Labels                 shared.ObjectTypeDefinitionLabelsParam `json:"labels,omitzero"`
	// The names of properties that should be **required** when creating an object of
	// this type.
	RequiredProperties []string `json:"requiredProperties,omitzero"`
	// Names of properties that will be indexed for this object type in by HubSpot's
	// product search.
	SearchableProperties []string `json:"searchableProperties,omitzero"`
	// The names of secondary properties for this object. These will be displayed as
	// secondary on the HubSpot record page for this object type.
	SecondaryDisplayProperties []string `json:"secondaryDisplayProperties,omitzero"`
	paramObj
}

func (r ObjectTypeDefinitionPatchParam) MarshalJSON() (data []byte, err error) {
	type shadow ObjectTypeDefinitionPatchParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ObjectTypeDefinitionPatchParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Defines a property to create.
//
// The properties FieldType, Label, Name, Type are required.
type ObjectTypePropertyCreateParam struct {
	// Controls how the property appears in HubSpot.
	FieldType string `json:"fieldType,required"`
	// A human-readable property label that will be shown in HubSpot.
	Label string `json:"label,required"`
	// The internal property name, which must be used when referencing the property
	// from the API.
	Name string `json:"name,required"`
	// The data type of the property.
	//
	// Any of "string", "number", "date", "datetime", "enumeration", "bool".
	Type ObjectTypePropertyCreateType `json:"type,omitzero,required"`
	// A description of the property that will be shown as help text in HubSpot.
	Description param.Opt[string] `json:"description,omitzero"`
	// The order that this property should be displayed in the HubSpot UI relative to
	// other properties for this object type. Properties are displayed in order
	// starting with the lowest positive integer value. A value of -1 will cause the
	// property to be displayed **after** any positive values.
	DisplayOrder param.Opt[int64] `json:"displayOrder,omitzero"`
	// Whether the property can be used in a HubSpot form.
	FormField param.Opt[bool] `json:"formField,omitzero"`
	// The name of the group this property belongs to.
	GroupName param.Opt[string] `json:"groupName,omitzero"`
	// Whether or not the property's value must be unique. Once set, this can't be
	// changed.
	HasUniqueValue param.Opt[bool] `json:"hasUniqueValue,omitzero"`
	Hidden         param.Opt[bool] `json:"hidden,omitzero"`
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
	// Any of "unformatted", "formatted", "currency", "percentage", "duration",
	// "probability".
	NumberDisplayHint ObjectTypePropertyCreateNumberDisplayHint `json:"numberDisplayHint,omitzero"`
	// A list of available options for the property. This field is only required for
	// enumerated properties.
	Options []shared.OptionInputParam `json:"options,omitzero"`
	// Controls how the property options will be sorted in the HubSpot UI.
	//
	// Any of "DISPLAY_ORDER", "ALPHABETICAL".
	OptionSortStrategy ObjectTypePropertyCreateOptionSortStrategy `json:"optionSortStrategy,omitzero"`
	// Controls how text properties are formatted in the HubSpot UI
	//
	// Any of "unformatted_single_line", "multi_line", "email", "phone_number",
	// "domain_name", "ip_address", "physical_address", "postal_code".
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
	ObjectTypePropertyCreateTypeString      ObjectTypePropertyCreateType = "string"
	ObjectTypePropertyCreateTypeNumber      ObjectTypePropertyCreateType = "number"
	ObjectTypePropertyCreateTypeDate        ObjectTypePropertyCreateType = "date"
	ObjectTypePropertyCreateTypeDatetime    ObjectTypePropertyCreateType = "datetime"
	ObjectTypePropertyCreateTypeEnumeration ObjectTypePropertyCreateType = "enumeration"
	ObjectTypePropertyCreateTypeBool        ObjectTypePropertyCreateType = "bool"
)

// Controls how numeric properties are formatted in the HubSpot UI
type ObjectTypePropertyCreateNumberDisplayHint string

const (
	ObjectTypePropertyCreateNumberDisplayHintUnformatted ObjectTypePropertyCreateNumberDisplayHint = "unformatted"
	ObjectTypePropertyCreateNumberDisplayHintFormatted   ObjectTypePropertyCreateNumberDisplayHint = "formatted"
	ObjectTypePropertyCreateNumberDisplayHintCurrency    ObjectTypePropertyCreateNumberDisplayHint = "currency"
	ObjectTypePropertyCreateNumberDisplayHintPercentage  ObjectTypePropertyCreateNumberDisplayHint = "percentage"
	ObjectTypePropertyCreateNumberDisplayHintDuration    ObjectTypePropertyCreateNumberDisplayHint = "duration"
	ObjectTypePropertyCreateNumberDisplayHintProbability ObjectTypePropertyCreateNumberDisplayHint = "probability"
)

// Controls how the property options will be sorted in the HubSpot UI.
type ObjectTypePropertyCreateOptionSortStrategy string

const (
	ObjectTypePropertyCreateOptionSortStrategyDisplayOrder ObjectTypePropertyCreateOptionSortStrategy = "DISPLAY_ORDER"
	ObjectTypePropertyCreateOptionSortStrategyAlphabetical ObjectTypePropertyCreateOptionSortStrategy = "ALPHABETICAL"
)

// Controls how text properties are formatted in the HubSpot UI
type ObjectTypePropertyCreateTextDisplayHint string

const (
	ObjectTypePropertyCreateTextDisplayHintUnformattedSingleLine ObjectTypePropertyCreateTextDisplayHint = "unformatted_single_line"
	ObjectTypePropertyCreateTextDisplayHintMultiLine             ObjectTypePropertyCreateTextDisplayHint = "multi_line"
	ObjectTypePropertyCreateTextDisplayHintEmail                 ObjectTypePropertyCreateTextDisplayHint = "email"
	ObjectTypePropertyCreateTextDisplayHintPhoneNumber           ObjectTypePropertyCreateTextDisplayHint = "phone_number"
	ObjectTypePropertyCreateTextDisplayHintDomainName            ObjectTypePropertyCreateTextDisplayHint = "domain_name"
	ObjectTypePropertyCreateTextDisplayHintIPAddress             ObjectTypePropertyCreateTextDisplayHint = "ip_address"
	ObjectTypePropertyCreateTextDisplayHintPhysicalAddress       ObjectTypePropertyCreateTextDisplayHint = "physical_address"
	ObjectTypePropertyCreateTextDisplayHintPostalCode            ObjectTypePropertyCreateTextDisplayHint = "postal_code"
)

// The definition of an association
type ObjectSchemaNewAssociationResponse struct {
	// The unique ID of the associated object (e.g., a contact ID).
	ID int64 `json:"id,required"`
	// Whether custom labels can be used in the association.
	AllowsCustomLabels bool `json:"allowsCustomLabels,required"`
	// The cardinality from the source object's perspective, either "ONE_TO_ONE" or
	// "ONE_TO_MANY".
	//
	// Any of "ONE_TO_ONE", "ONE_TO_MANY".
	Cardinality ObjectSchemaNewAssociationResponseCardinality `json:"cardinality,required"`
	// The category of the association. Can be: "HUBSPOT_DEFINED", "USER_DEFINED", or
	// "INTEGRATOR_DEFINED"
	//
	// Any of "HUBSPOT_DEFINED", "USER_DEFINED", "INTEGRATOR_DEFINED".
	Category ObjectSchemaNewAssociationResponseCategory `json:"category,required"`
	// The ID of the source object type (e.g., 0-1 for contacts).
	FromObjectTypeID string `json:"fromObjectTypeId,required"`
	// Whether all potential linked objects are included in the association
	HasAllAssociatedObjects bool `json:"hasAllAssociatedObjects,required"`
	// Whether deletions in the association should cause cascading deletes to linked
	// objects.
	HasCascadingDeletes bool `json:"hasCascadingDeletes,required"`
	// Whether a user has set a limit for the number of source objects.
	HasUserEnforcedMaxFromObjectIDs bool `json:"hasUserEnforcedMaxFromObjectIds,required"`
	// Whether a user has set a limit for the number of destination objects.
	HasUserEnforcedMaxToObjectIDs bool `json:"hasUserEnforcedMaxToObjectIds,required"`
	// Whether the association is hidden or not.
	Hidden bool `json:"hidden,required"`
	// Whether the reverse association can also support custom labels.
	InverseAllowsCustomLabels bool `json:"inverseAllowsCustomLabels,required"`
	// The cardinality from the destination object's perspective, either "ONE_TO_ONE"
	// or "ONE_TO_MANY".
	//
	// Any of "ONE_TO_ONE", "ONE_TO_MANY".
	InverseCardinality ObjectSchemaNewAssociationResponseInverseCardinality `json:"inverseCardinality,required"`
	// Whether all potential reverse linked objects are included in the association.
	InverseHasAllAssociatedObjects bool `json:"inverseHasAllAssociatedObjects,required"`
	// The unique ID for the inverse side of the association.
	InverseID int64 `json:"inverseId,required"`
	// The name used to describe the inverse relationship in this association
	InverseName string `json:"inverseName,required"`
	// Whether the inverse association is considered primary.
	IsInversePrimary bool `json:"isInversePrimary,required"`
	// Whether the association is the primary link between the entities involved.
	IsPrimary bool `json:"isPrimary,required"`
	// The maximum number of source object IDs allowed in the association.
	MaxFromObjectIDs int64 `json:"maxFromObjectIds,required"`
	// The maximum number of destination object IDs allowed in the association.
	MaxToObjectIDs int64 `json:"maxToObjectIds,required"`
	// For labeled association types, the internal name of the association.
	Name string `json:"name,required"`
	// A unique across-portal ID applied to the association.
	PortalUniqueIdentifier string `json:"portalUniqueIdentifier,required"`
	// The ID of the destination object type (e.g., 0-3 for deals).
	ToObjectTypeID string `json:"toObjectTypeId,required"`
	// The name of the source object type (e.g,. "DEAL" or "QUOTE").
	//
	// Any of "CONTACT", "COMPANY", "DEAL", "ENGAGEMENT", "TICKET", "OWNER", "PRODUCT",
	// "LINE_ITEM", "BET_DELIVERABLE_SERVICE", "CONTENT", "CONVERSATION", "BET_ALERT",
	// "PORTAL", "QUOTE", "FORM_SUBMISSION_INBOUNDDB", "QUOTA", "UNSUBSCRIBE",
	// "COMMUNICATION", "FEEDBACK_SUBMISSION", "ATTRIBUTION", "SALESFORCE_SYNC_ERROR",
	// "RESTORABLE_CRM_OBJECT", "HUB", "LANDING_PAGE", "PRODUCT_OR_FOLDER", "TASK",
	// "FORM", "MARKETING_EMAIL", "AD_ACCOUNT", "AD_CAMPAIGN", "AD_GROUP", "AD",
	// "KEYWORD", "CAMPAIGN", "SOCIAL_CHANNEL", "SOCIAL_POST", "SITE_PAGE",
	// "BLOG_POST", "IMPORT", "EXPORT", "CTA", "TASK_TEMPLATE",
	// "AUTOMATION_PLATFORM_FLOW", "OBJECT_LIST", "NOTE", "MEETING_EVENT", "CALL",
	// "EMAIL", "PUBLISHING_TASK", "CONVERSATION_SESSION",
	// "CONTACT_CREATE_ATTRIBUTION", "INVOICE", "MARKETING_EVENT",
	// "CONVERSATION_INBOX", "CHATFLOW", "MEDIA_BRIDGE", "SEQUENCE", "SEQUENCE_STEP",
	// "FORECAST", "SNIPPET", "TEMPLATE", "DEAL_CREATE_ATTRIBUTION", "QUOTE_TEMPLATE",
	// "QUOTE_MODULE", "QUOTE_MODULE_FIELD", "QUOTE_FIELD", "SEQUENCE_ENROLLMENT",
	// "SUBSCRIPTION", "ACCEPTANCE_TEST", "SOCIAL_BROADCAST", "DEAL_SPLIT",
	// "DEAL_REGISTRATION", "GOAL_TARGET", "GOAL_TARGET_GROUP",
	// "PORTAL_OBJECT_SYNC_MESSAGE", "FILE_MANAGER_FILE", "FILE_MANAGER_FOLDER",
	// "SEQUENCE_STEP_ENROLLMENT", "APPROVAL", "APPROVAL_STEP", "CTA_VARIANT",
	// "SALES_DOCUMENT", "DISCOUNT", "FEE", "TAX", "MARKETING_CALENDAR",
	// "PERMISSIONS_TESTING", "PRIVACY_SCANNER_COOKIE", "DATA_SYNC_STATE",
	// "WEB_INTERACTIVE", "PLAYBOOK", "FOLDER", "PLAYBOOK_QUESTION",
	// "PLAYBOOK_SUBMISSION", "PLAYBOOK_SUBMISSION_ANSWER", "COMMERCE_PAYMENT",
	// "GSC_PROPERTY", "SOX_PROTECTED_DUMMY_TYPE", "BLOG_LISTING_PAGE",
	// "QUARANTINED_SUBMISSION", "PAYMENT_SCHEDULE", "PAYMENT_SCHEDULE_INSTALLMENT",
	// "MARKETING_CAMPAIGN_UTM", "DISCOUNT_TEMPLATE", "DISCOUNT_CODE",
	// "FEEDBACK_SURVEY", "CMS_URL", "SALES_TASK", "SALES_WORKLOAD", "USER",
	// "POSTAL_MAIL", "SCHEMAS_BACKEND_TEST", "PAYMENT_LINK", "SUBMISSION_TAG",
	// "CAMPAIGN_STEP", "SCHEDULING_PAGE", "SOX_PROTECTED_TEST_TYPE", "ORDER",
	// "MARKETING_SMS", "PARTNER_ACCOUNT", "CAMPAIGN_TEMPLATE",
	// "CAMPAIGN_TEMPLATE_STEP", "PLAYLIST", "CLIP", "CAMPAIGN_BUDGET_ITEM",
	// "CAMPAIGN_SPEND_ITEM", "MIC", "CONTENT_AUDIT", "CONTENT_AUDIT_PAGE",
	// "PLAYLIST_FOLDER", "LEAD", "ABANDONED_CART", "EXTERNAL_WEB_URL", "VIEW",
	// "VIEW_BLOCK", "ROSTER", "CART", "AUTOMATION_PLATFORM_FLOW_ACTION",
	// "SOCIAL_PROFILE", "PARTNER_CLIENT", "ROSTER_MEMBER",
	// "MARKETING_EVENT_ATTENDANCE", "ALL_PAGES", "AI_FORECAST",
	// "CRM_PIPELINES_DUMMY_TYPE", "KNOWLEDGE_ARTICLE", "PROPERTY_INFO",
	// "DATA_PRIVACY_CONSENT", "GOAL_TEMPLATE", "SCORE_CONFIGURATION", "AUDIENCE",
	// "PARTNER_CLIENT_REVENUE", "AUTOMATION_JOURNEY", "COMBO_EVENT_CONFIGURATION",
	// "CRM_OBJECTS_DUMMY_TYPE", "CASE_STUDY", "SERVICE", "PODCAST_EPISODE",
	// "PARTNER_SERVICE", "UNKNOWN".
	FromObjectType ObjectSchemaNewAssociationResponseFromObjectType `json:"fromObjectType"`
	// The label used to describe the reverse relationship in an association.
	InverseLabel string `json:"inverseLabel"`
	// The label given to an association.
	Label string `json:"label"`
	// The name of the destination object type (e.g,. "DEAL" or "QUOTE").
	//
	// Any of "CONTACT", "COMPANY", "DEAL", "ENGAGEMENT", "TICKET", "OWNER", "PRODUCT",
	// "LINE_ITEM", "BET_DELIVERABLE_SERVICE", "CONTENT", "CONVERSATION", "BET_ALERT",
	// "PORTAL", "QUOTE", "FORM_SUBMISSION_INBOUNDDB", "QUOTA", "UNSUBSCRIBE",
	// "COMMUNICATION", "FEEDBACK_SUBMISSION", "ATTRIBUTION", "SALESFORCE_SYNC_ERROR",
	// "RESTORABLE_CRM_OBJECT", "HUB", "LANDING_PAGE", "PRODUCT_OR_FOLDER", "TASK",
	// "FORM", "MARKETING_EMAIL", "AD_ACCOUNT", "AD_CAMPAIGN", "AD_GROUP", "AD",
	// "KEYWORD", "CAMPAIGN", "SOCIAL_CHANNEL", "SOCIAL_POST", "SITE_PAGE",
	// "BLOG_POST", "IMPORT", "EXPORT", "CTA", "TASK_TEMPLATE",
	// "AUTOMATION_PLATFORM_FLOW", "OBJECT_LIST", "NOTE", "MEETING_EVENT", "CALL",
	// "EMAIL", "PUBLISHING_TASK", "CONVERSATION_SESSION",
	// "CONTACT_CREATE_ATTRIBUTION", "INVOICE", "MARKETING_EVENT",
	// "CONVERSATION_INBOX", "CHATFLOW", "MEDIA_BRIDGE", "SEQUENCE", "SEQUENCE_STEP",
	// "FORECAST", "SNIPPET", "TEMPLATE", "DEAL_CREATE_ATTRIBUTION", "QUOTE_TEMPLATE",
	// "QUOTE_MODULE", "QUOTE_MODULE_FIELD", "QUOTE_FIELD", "SEQUENCE_ENROLLMENT",
	// "SUBSCRIPTION", "ACCEPTANCE_TEST", "SOCIAL_BROADCAST", "DEAL_SPLIT",
	// "DEAL_REGISTRATION", "GOAL_TARGET", "GOAL_TARGET_GROUP",
	// "PORTAL_OBJECT_SYNC_MESSAGE", "FILE_MANAGER_FILE", "FILE_MANAGER_FOLDER",
	// "SEQUENCE_STEP_ENROLLMENT", "APPROVAL", "APPROVAL_STEP", "CTA_VARIANT",
	// "SALES_DOCUMENT", "DISCOUNT", "FEE", "TAX", "MARKETING_CALENDAR",
	// "PERMISSIONS_TESTING", "PRIVACY_SCANNER_COOKIE", "DATA_SYNC_STATE",
	// "WEB_INTERACTIVE", "PLAYBOOK", "FOLDER", "PLAYBOOK_QUESTION",
	// "PLAYBOOK_SUBMISSION", "PLAYBOOK_SUBMISSION_ANSWER", "COMMERCE_PAYMENT",
	// "GSC_PROPERTY", "SOX_PROTECTED_DUMMY_TYPE", "BLOG_LISTING_PAGE",
	// "QUARANTINED_SUBMISSION", "PAYMENT_SCHEDULE", "PAYMENT_SCHEDULE_INSTALLMENT",
	// "MARKETING_CAMPAIGN_UTM", "DISCOUNT_TEMPLATE", "DISCOUNT_CODE",
	// "FEEDBACK_SURVEY", "CMS_URL", "SALES_TASK", "SALES_WORKLOAD", "USER",
	// "POSTAL_MAIL", "SCHEMAS_BACKEND_TEST", "PAYMENT_LINK", "SUBMISSION_TAG",
	// "CAMPAIGN_STEP", "SCHEDULING_PAGE", "SOX_PROTECTED_TEST_TYPE", "ORDER",
	// "MARKETING_SMS", "PARTNER_ACCOUNT", "CAMPAIGN_TEMPLATE",
	// "CAMPAIGN_TEMPLATE_STEP", "PLAYLIST", "CLIP", "CAMPAIGN_BUDGET_ITEM",
	// "CAMPAIGN_SPEND_ITEM", "MIC", "CONTENT_AUDIT", "CONTENT_AUDIT_PAGE",
	// "PLAYLIST_FOLDER", "LEAD", "ABANDONED_CART", "EXTERNAL_WEB_URL", "VIEW",
	// "VIEW_BLOCK", "ROSTER", "CART", "AUTOMATION_PLATFORM_FLOW_ACTION",
	// "SOCIAL_PROFILE", "PARTNER_CLIENT", "ROSTER_MEMBER",
	// "MARKETING_EVENT_ATTENDANCE", "ALL_PAGES", "AI_FORECAST",
	// "CRM_PIPELINES_DUMMY_TYPE", "KNOWLEDGE_ARTICLE", "PROPERTY_INFO",
	// "DATA_PRIVACY_CONSENT", "GOAL_TEMPLATE", "SCORE_CONFIGURATION", "AUDIENCE",
	// "PARTNER_CLIENT_REVENUE", "AUTOMATION_JOURNEY", "COMBO_EVENT_CONFIGURATION",
	// "CRM_OBJECTS_DUMMY_TYPE", "CASE_STUDY", "SERVICE", "PODCAST_EPISODE",
	// "PARTNER_SERVICE", "UNKNOWN".
	ToObjectType ObjectSchemaNewAssociationResponseToObjectType `json:"toObjectType"`
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
		IsInversePrimary                respjson.Field
		IsPrimary                       respjson.Field
		MaxFromObjectIDs                respjson.Field
		MaxToObjectIDs                  respjson.Field
		Name                            respjson.Field
		PortalUniqueIdentifier          respjson.Field
		ToObjectTypeID                  respjson.Field
		FromObjectType                  respjson.Field
		InverseLabel                    respjson.Field
		Label                           respjson.Field
		ToObjectType                    respjson.Field
		ExtraFields                     map[string]respjson.Field
		raw                             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ObjectSchemaNewAssociationResponse) RawJSON() string { return r.JSON.raw }
func (r *ObjectSchemaNewAssociationResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The cardinality from the source object's perspective, either "ONE_TO_ONE" or
// "ONE_TO_MANY".
type ObjectSchemaNewAssociationResponseCardinality string

const (
	ObjectSchemaNewAssociationResponseCardinalityOneToOne  ObjectSchemaNewAssociationResponseCardinality = "ONE_TO_ONE"
	ObjectSchemaNewAssociationResponseCardinalityOneToMany ObjectSchemaNewAssociationResponseCardinality = "ONE_TO_MANY"
)

// The category of the association. Can be: "HUBSPOT_DEFINED", "USER_DEFINED", or
// "INTEGRATOR_DEFINED"
type ObjectSchemaNewAssociationResponseCategory string

const (
	ObjectSchemaNewAssociationResponseCategoryHubspotDefined    ObjectSchemaNewAssociationResponseCategory = "HUBSPOT_DEFINED"
	ObjectSchemaNewAssociationResponseCategoryUserDefined       ObjectSchemaNewAssociationResponseCategory = "USER_DEFINED"
	ObjectSchemaNewAssociationResponseCategoryIntegratorDefined ObjectSchemaNewAssociationResponseCategory = "INTEGRATOR_DEFINED"
)

// The cardinality from the destination object's perspective, either "ONE_TO_ONE"
// or "ONE_TO_MANY".
type ObjectSchemaNewAssociationResponseInverseCardinality string

const (
	ObjectSchemaNewAssociationResponseInverseCardinalityOneToOne  ObjectSchemaNewAssociationResponseInverseCardinality = "ONE_TO_ONE"
	ObjectSchemaNewAssociationResponseInverseCardinalityOneToMany ObjectSchemaNewAssociationResponseInverseCardinality = "ONE_TO_MANY"
)

// The name of the source object type (e.g,. "DEAL" or "QUOTE").
type ObjectSchemaNewAssociationResponseFromObjectType string

const (
	ObjectSchemaNewAssociationResponseFromObjectTypeContact                      ObjectSchemaNewAssociationResponseFromObjectType = "CONTACT"
	ObjectSchemaNewAssociationResponseFromObjectTypeCompany                      ObjectSchemaNewAssociationResponseFromObjectType = "COMPANY"
	ObjectSchemaNewAssociationResponseFromObjectTypeDeal                         ObjectSchemaNewAssociationResponseFromObjectType = "DEAL"
	ObjectSchemaNewAssociationResponseFromObjectTypeEngagement                   ObjectSchemaNewAssociationResponseFromObjectType = "ENGAGEMENT"
	ObjectSchemaNewAssociationResponseFromObjectTypeTicket                       ObjectSchemaNewAssociationResponseFromObjectType = "TICKET"
	ObjectSchemaNewAssociationResponseFromObjectTypeOwner                        ObjectSchemaNewAssociationResponseFromObjectType = "OWNER"
	ObjectSchemaNewAssociationResponseFromObjectTypeProduct                      ObjectSchemaNewAssociationResponseFromObjectType = "PRODUCT"
	ObjectSchemaNewAssociationResponseFromObjectTypeLineItem                     ObjectSchemaNewAssociationResponseFromObjectType = "LINE_ITEM"
	ObjectSchemaNewAssociationResponseFromObjectTypeBetDeliverableService        ObjectSchemaNewAssociationResponseFromObjectType = "BET_DELIVERABLE_SERVICE"
	ObjectSchemaNewAssociationResponseFromObjectTypeContent                      ObjectSchemaNewAssociationResponseFromObjectType = "CONTENT"
	ObjectSchemaNewAssociationResponseFromObjectTypeConversation                 ObjectSchemaNewAssociationResponseFromObjectType = "CONVERSATION"
	ObjectSchemaNewAssociationResponseFromObjectTypeBetAlert                     ObjectSchemaNewAssociationResponseFromObjectType = "BET_ALERT"
	ObjectSchemaNewAssociationResponseFromObjectTypePortal                       ObjectSchemaNewAssociationResponseFromObjectType = "PORTAL"
	ObjectSchemaNewAssociationResponseFromObjectTypeQuote                        ObjectSchemaNewAssociationResponseFromObjectType = "QUOTE"
	ObjectSchemaNewAssociationResponseFromObjectTypeFormSubmissionInbounddb      ObjectSchemaNewAssociationResponseFromObjectType = "FORM_SUBMISSION_INBOUNDDB"
	ObjectSchemaNewAssociationResponseFromObjectTypeQuota                        ObjectSchemaNewAssociationResponseFromObjectType = "QUOTA"
	ObjectSchemaNewAssociationResponseFromObjectTypeUnsubscribe                  ObjectSchemaNewAssociationResponseFromObjectType = "UNSUBSCRIBE"
	ObjectSchemaNewAssociationResponseFromObjectTypeCommunication                ObjectSchemaNewAssociationResponseFromObjectType = "COMMUNICATION"
	ObjectSchemaNewAssociationResponseFromObjectTypeFeedbackSubmission           ObjectSchemaNewAssociationResponseFromObjectType = "FEEDBACK_SUBMISSION"
	ObjectSchemaNewAssociationResponseFromObjectTypeAttribution                  ObjectSchemaNewAssociationResponseFromObjectType = "ATTRIBUTION"
	ObjectSchemaNewAssociationResponseFromObjectTypeSalesforceSyncError          ObjectSchemaNewAssociationResponseFromObjectType = "SALESFORCE_SYNC_ERROR"
	ObjectSchemaNewAssociationResponseFromObjectTypeRestorableCRMObject          ObjectSchemaNewAssociationResponseFromObjectType = "RESTORABLE_CRM_OBJECT"
	ObjectSchemaNewAssociationResponseFromObjectTypeHub                          ObjectSchemaNewAssociationResponseFromObjectType = "HUB"
	ObjectSchemaNewAssociationResponseFromObjectTypeLandingPage                  ObjectSchemaNewAssociationResponseFromObjectType = "LANDING_PAGE"
	ObjectSchemaNewAssociationResponseFromObjectTypeProductOrFolder              ObjectSchemaNewAssociationResponseFromObjectType = "PRODUCT_OR_FOLDER"
	ObjectSchemaNewAssociationResponseFromObjectTypeTask                         ObjectSchemaNewAssociationResponseFromObjectType = "TASK"
	ObjectSchemaNewAssociationResponseFromObjectTypeForm                         ObjectSchemaNewAssociationResponseFromObjectType = "FORM"
	ObjectSchemaNewAssociationResponseFromObjectTypeMarketingEmail               ObjectSchemaNewAssociationResponseFromObjectType = "MARKETING_EMAIL"
	ObjectSchemaNewAssociationResponseFromObjectTypeAdAccount                    ObjectSchemaNewAssociationResponseFromObjectType = "AD_ACCOUNT"
	ObjectSchemaNewAssociationResponseFromObjectTypeAdCampaign                   ObjectSchemaNewAssociationResponseFromObjectType = "AD_CAMPAIGN"
	ObjectSchemaNewAssociationResponseFromObjectTypeAdGroup                      ObjectSchemaNewAssociationResponseFromObjectType = "AD_GROUP"
	ObjectSchemaNewAssociationResponseFromObjectTypeAd                           ObjectSchemaNewAssociationResponseFromObjectType = "AD"
	ObjectSchemaNewAssociationResponseFromObjectTypeKeyword                      ObjectSchemaNewAssociationResponseFromObjectType = "KEYWORD"
	ObjectSchemaNewAssociationResponseFromObjectTypeCampaign                     ObjectSchemaNewAssociationResponseFromObjectType = "CAMPAIGN"
	ObjectSchemaNewAssociationResponseFromObjectTypeSocialChannel                ObjectSchemaNewAssociationResponseFromObjectType = "SOCIAL_CHANNEL"
	ObjectSchemaNewAssociationResponseFromObjectTypeSocialPost                   ObjectSchemaNewAssociationResponseFromObjectType = "SOCIAL_POST"
	ObjectSchemaNewAssociationResponseFromObjectTypeSitePage                     ObjectSchemaNewAssociationResponseFromObjectType = "SITE_PAGE"
	ObjectSchemaNewAssociationResponseFromObjectTypeBlogPost                     ObjectSchemaNewAssociationResponseFromObjectType = "BLOG_POST"
	ObjectSchemaNewAssociationResponseFromObjectTypeImport                       ObjectSchemaNewAssociationResponseFromObjectType = "IMPORT"
	ObjectSchemaNewAssociationResponseFromObjectTypeExport                       ObjectSchemaNewAssociationResponseFromObjectType = "EXPORT"
	ObjectSchemaNewAssociationResponseFromObjectTypeCta                          ObjectSchemaNewAssociationResponseFromObjectType = "CTA"
	ObjectSchemaNewAssociationResponseFromObjectTypeTaskTemplate                 ObjectSchemaNewAssociationResponseFromObjectType = "TASK_TEMPLATE"
	ObjectSchemaNewAssociationResponseFromObjectTypeAutomationPlatformFlow       ObjectSchemaNewAssociationResponseFromObjectType = "AUTOMATION_PLATFORM_FLOW"
	ObjectSchemaNewAssociationResponseFromObjectTypeObjectList                   ObjectSchemaNewAssociationResponseFromObjectType = "OBJECT_LIST"
	ObjectSchemaNewAssociationResponseFromObjectTypeNote                         ObjectSchemaNewAssociationResponseFromObjectType = "NOTE"
	ObjectSchemaNewAssociationResponseFromObjectTypeMeetingEvent                 ObjectSchemaNewAssociationResponseFromObjectType = "MEETING_EVENT"
	ObjectSchemaNewAssociationResponseFromObjectTypeCall                         ObjectSchemaNewAssociationResponseFromObjectType = "CALL"
	ObjectSchemaNewAssociationResponseFromObjectTypeEmail                        ObjectSchemaNewAssociationResponseFromObjectType = "EMAIL"
	ObjectSchemaNewAssociationResponseFromObjectTypePublishingTask               ObjectSchemaNewAssociationResponseFromObjectType = "PUBLISHING_TASK"
	ObjectSchemaNewAssociationResponseFromObjectTypeConversationSession          ObjectSchemaNewAssociationResponseFromObjectType = "CONVERSATION_SESSION"
	ObjectSchemaNewAssociationResponseFromObjectTypeContactCreateAttribution     ObjectSchemaNewAssociationResponseFromObjectType = "CONTACT_CREATE_ATTRIBUTION"
	ObjectSchemaNewAssociationResponseFromObjectTypeInvoice                      ObjectSchemaNewAssociationResponseFromObjectType = "INVOICE"
	ObjectSchemaNewAssociationResponseFromObjectTypeMarketingEvent               ObjectSchemaNewAssociationResponseFromObjectType = "MARKETING_EVENT"
	ObjectSchemaNewAssociationResponseFromObjectTypeConversationInbox            ObjectSchemaNewAssociationResponseFromObjectType = "CONVERSATION_INBOX"
	ObjectSchemaNewAssociationResponseFromObjectTypeChatflow                     ObjectSchemaNewAssociationResponseFromObjectType = "CHATFLOW"
	ObjectSchemaNewAssociationResponseFromObjectTypeMediaBridge                  ObjectSchemaNewAssociationResponseFromObjectType = "MEDIA_BRIDGE"
	ObjectSchemaNewAssociationResponseFromObjectTypeSequence                     ObjectSchemaNewAssociationResponseFromObjectType = "SEQUENCE"
	ObjectSchemaNewAssociationResponseFromObjectTypeSequenceStep                 ObjectSchemaNewAssociationResponseFromObjectType = "SEQUENCE_STEP"
	ObjectSchemaNewAssociationResponseFromObjectTypeForecast                     ObjectSchemaNewAssociationResponseFromObjectType = "FORECAST"
	ObjectSchemaNewAssociationResponseFromObjectTypeSnippet                      ObjectSchemaNewAssociationResponseFromObjectType = "SNIPPET"
	ObjectSchemaNewAssociationResponseFromObjectTypeTemplate                     ObjectSchemaNewAssociationResponseFromObjectType = "TEMPLATE"
	ObjectSchemaNewAssociationResponseFromObjectTypeDealCreateAttribution        ObjectSchemaNewAssociationResponseFromObjectType = "DEAL_CREATE_ATTRIBUTION"
	ObjectSchemaNewAssociationResponseFromObjectTypeQuoteTemplate                ObjectSchemaNewAssociationResponseFromObjectType = "QUOTE_TEMPLATE"
	ObjectSchemaNewAssociationResponseFromObjectTypeQuoteModule                  ObjectSchemaNewAssociationResponseFromObjectType = "QUOTE_MODULE"
	ObjectSchemaNewAssociationResponseFromObjectTypeQuoteModuleField             ObjectSchemaNewAssociationResponseFromObjectType = "QUOTE_MODULE_FIELD"
	ObjectSchemaNewAssociationResponseFromObjectTypeQuoteField                   ObjectSchemaNewAssociationResponseFromObjectType = "QUOTE_FIELD"
	ObjectSchemaNewAssociationResponseFromObjectTypeSequenceEnrollment           ObjectSchemaNewAssociationResponseFromObjectType = "SEQUENCE_ENROLLMENT"
	ObjectSchemaNewAssociationResponseFromObjectTypeSubscription                 ObjectSchemaNewAssociationResponseFromObjectType = "SUBSCRIPTION"
	ObjectSchemaNewAssociationResponseFromObjectTypeAcceptanceTest               ObjectSchemaNewAssociationResponseFromObjectType = "ACCEPTANCE_TEST"
	ObjectSchemaNewAssociationResponseFromObjectTypeSocialBroadcast              ObjectSchemaNewAssociationResponseFromObjectType = "SOCIAL_BROADCAST"
	ObjectSchemaNewAssociationResponseFromObjectTypeDealSplit                    ObjectSchemaNewAssociationResponseFromObjectType = "DEAL_SPLIT"
	ObjectSchemaNewAssociationResponseFromObjectTypeDealRegistration             ObjectSchemaNewAssociationResponseFromObjectType = "DEAL_REGISTRATION"
	ObjectSchemaNewAssociationResponseFromObjectTypeGoalTarget                   ObjectSchemaNewAssociationResponseFromObjectType = "GOAL_TARGET"
	ObjectSchemaNewAssociationResponseFromObjectTypeGoalTargetGroup              ObjectSchemaNewAssociationResponseFromObjectType = "GOAL_TARGET_GROUP"
	ObjectSchemaNewAssociationResponseFromObjectTypePortalObjectSyncMessage      ObjectSchemaNewAssociationResponseFromObjectType = "PORTAL_OBJECT_SYNC_MESSAGE"
	ObjectSchemaNewAssociationResponseFromObjectTypeFileManagerFile              ObjectSchemaNewAssociationResponseFromObjectType = "FILE_MANAGER_FILE"
	ObjectSchemaNewAssociationResponseFromObjectTypeFileManagerFolder            ObjectSchemaNewAssociationResponseFromObjectType = "FILE_MANAGER_FOLDER"
	ObjectSchemaNewAssociationResponseFromObjectTypeSequenceStepEnrollment       ObjectSchemaNewAssociationResponseFromObjectType = "SEQUENCE_STEP_ENROLLMENT"
	ObjectSchemaNewAssociationResponseFromObjectTypeApproval                     ObjectSchemaNewAssociationResponseFromObjectType = "APPROVAL"
	ObjectSchemaNewAssociationResponseFromObjectTypeApprovalStep                 ObjectSchemaNewAssociationResponseFromObjectType = "APPROVAL_STEP"
	ObjectSchemaNewAssociationResponseFromObjectTypeCtaVariant                   ObjectSchemaNewAssociationResponseFromObjectType = "CTA_VARIANT"
	ObjectSchemaNewAssociationResponseFromObjectTypeSalesDocument                ObjectSchemaNewAssociationResponseFromObjectType = "SALES_DOCUMENT"
	ObjectSchemaNewAssociationResponseFromObjectTypeDiscount                     ObjectSchemaNewAssociationResponseFromObjectType = "DISCOUNT"
	ObjectSchemaNewAssociationResponseFromObjectTypeFee                          ObjectSchemaNewAssociationResponseFromObjectType = "FEE"
	ObjectSchemaNewAssociationResponseFromObjectTypeTax                          ObjectSchemaNewAssociationResponseFromObjectType = "TAX"
	ObjectSchemaNewAssociationResponseFromObjectTypeMarketingCalendar            ObjectSchemaNewAssociationResponseFromObjectType = "MARKETING_CALENDAR"
	ObjectSchemaNewAssociationResponseFromObjectTypePermissionsTesting           ObjectSchemaNewAssociationResponseFromObjectType = "PERMISSIONS_TESTING"
	ObjectSchemaNewAssociationResponseFromObjectTypePrivacyScannerCookie         ObjectSchemaNewAssociationResponseFromObjectType = "PRIVACY_SCANNER_COOKIE"
	ObjectSchemaNewAssociationResponseFromObjectTypeDataSyncState                ObjectSchemaNewAssociationResponseFromObjectType = "DATA_SYNC_STATE"
	ObjectSchemaNewAssociationResponseFromObjectTypeWebInteractive               ObjectSchemaNewAssociationResponseFromObjectType = "WEB_INTERACTIVE"
	ObjectSchemaNewAssociationResponseFromObjectTypePlaybook                     ObjectSchemaNewAssociationResponseFromObjectType = "PLAYBOOK"
	ObjectSchemaNewAssociationResponseFromObjectTypeFolder                       ObjectSchemaNewAssociationResponseFromObjectType = "FOLDER"
	ObjectSchemaNewAssociationResponseFromObjectTypePlaybookQuestion             ObjectSchemaNewAssociationResponseFromObjectType = "PLAYBOOK_QUESTION"
	ObjectSchemaNewAssociationResponseFromObjectTypePlaybookSubmission           ObjectSchemaNewAssociationResponseFromObjectType = "PLAYBOOK_SUBMISSION"
	ObjectSchemaNewAssociationResponseFromObjectTypePlaybookSubmissionAnswer     ObjectSchemaNewAssociationResponseFromObjectType = "PLAYBOOK_SUBMISSION_ANSWER"
	ObjectSchemaNewAssociationResponseFromObjectTypeCommercePayment              ObjectSchemaNewAssociationResponseFromObjectType = "COMMERCE_PAYMENT"
	ObjectSchemaNewAssociationResponseFromObjectTypeGscProperty                  ObjectSchemaNewAssociationResponseFromObjectType = "GSC_PROPERTY"
	ObjectSchemaNewAssociationResponseFromObjectTypeSoxProtectedDummyType        ObjectSchemaNewAssociationResponseFromObjectType = "SOX_PROTECTED_DUMMY_TYPE"
	ObjectSchemaNewAssociationResponseFromObjectTypeBlogListingPage              ObjectSchemaNewAssociationResponseFromObjectType = "BLOG_LISTING_PAGE"
	ObjectSchemaNewAssociationResponseFromObjectTypeQuarantinedSubmission        ObjectSchemaNewAssociationResponseFromObjectType = "QUARANTINED_SUBMISSION"
	ObjectSchemaNewAssociationResponseFromObjectTypePaymentSchedule              ObjectSchemaNewAssociationResponseFromObjectType = "PAYMENT_SCHEDULE"
	ObjectSchemaNewAssociationResponseFromObjectTypePaymentScheduleInstallment   ObjectSchemaNewAssociationResponseFromObjectType = "PAYMENT_SCHEDULE_INSTALLMENT"
	ObjectSchemaNewAssociationResponseFromObjectTypeMarketingCampaignUtm         ObjectSchemaNewAssociationResponseFromObjectType = "MARKETING_CAMPAIGN_UTM"
	ObjectSchemaNewAssociationResponseFromObjectTypeDiscountTemplate             ObjectSchemaNewAssociationResponseFromObjectType = "DISCOUNT_TEMPLATE"
	ObjectSchemaNewAssociationResponseFromObjectTypeDiscountCode                 ObjectSchemaNewAssociationResponseFromObjectType = "DISCOUNT_CODE"
	ObjectSchemaNewAssociationResponseFromObjectTypeFeedbackSurvey               ObjectSchemaNewAssociationResponseFromObjectType = "FEEDBACK_SURVEY"
	ObjectSchemaNewAssociationResponseFromObjectTypeCmsURL                       ObjectSchemaNewAssociationResponseFromObjectType = "CMS_URL"
	ObjectSchemaNewAssociationResponseFromObjectTypeSalesTask                    ObjectSchemaNewAssociationResponseFromObjectType = "SALES_TASK"
	ObjectSchemaNewAssociationResponseFromObjectTypeSalesWorkload                ObjectSchemaNewAssociationResponseFromObjectType = "SALES_WORKLOAD"
	ObjectSchemaNewAssociationResponseFromObjectTypeUser                         ObjectSchemaNewAssociationResponseFromObjectType = "USER"
	ObjectSchemaNewAssociationResponseFromObjectTypePostalMail                   ObjectSchemaNewAssociationResponseFromObjectType = "POSTAL_MAIL"
	ObjectSchemaNewAssociationResponseFromObjectTypeSchemasBackendTest           ObjectSchemaNewAssociationResponseFromObjectType = "SCHEMAS_BACKEND_TEST"
	ObjectSchemaNewAssociationResponseFromObjectTypePaymentLink                  ObjectSchemaNewAssociationResponseFromObjectType = "PAYMENT_LINK"
	ObjectSchemaNewAssociationResponseFromObjectTypeSubmissionTag                ObjectSchemaNewAssociationResponseFromObjectType = "SUBMISSION_TAG"
	ObjectSchemaNewAssociationResponseFromObjectTypeCampaignStep                 ObjectSchemaNewAssociationResponseFromObjectType = "CAMPAIGN_STEP"
	ObjectSchemaNewAssociationResponseFromObjectTypeSchedulingPage               ObjectSchemaNewAssociationResponseFromObjectType = "SCHEDULING_PAGE"
	ObjectSchemaNewAssociationResponseFromObjectTypeSoxProtectedTestType         ObjectSchemaNewAssociationResponseFromObjectType = "SOX_PROTECTED_TEST_TYPE"
	ObjectSchemaNewAssociationResponseFromObjectTypeOrder                        ObjectSchemaNewAssociationResponseFromObjectType = "ORDER"
	ObjectSchemaNewAssociationResponseFromObjectTypeMarketingSMS                 ObjectSchemaNewAssociationResponseFromObjectType = "MARKETING_SMS"
	ObjectSchemaNewAssociationResponseFromObjectTypePartnerAccount               ObjectSchemaNewAssociationResponseFromObjectType = "PARTNER_ACCOUNT"
	ObjectSchemaNewAssociationResponseFromObjectTypeCampaignTemplate             ObjectSchemaNewAssociationResponseFromObjectType = "CAMPAIGN_TEMPLATE"
	ObjectSchemaNewAssociationResponseFromObjectTypeCampaignTemplateStep         ObjectSchemaNewAssociationResponseFromObjectType = "CAMPAIGN_TEMPLATE_STEP"
	ObjectSchemaNewAssociationResponseFromObjectTypePlaylist                     ObjectSchemaNewAssociationResponseFromObjectType = "PLAYLIST"
	ObjectSchemaNewAssociationResponseFromObjectTypeClip                         ObjectSchemaNewAssociationResponseFromObjectType = "CLIP"
	ObjectSchemaNewAssociationResponseFromObjectTypeCampaignBudgetItem           ObjectSchemaNewAssociationResponseFromObjectType = "CAMPAIGN_BUDGET_ITEM"
	ObjectSchemaNewAssociationResponseFromObjectTypeCampaignSpendItem            ObjectSchemaNewAssociationResponseFromObjectType = "CAMPAIGN_SPEND_ITEM"
	ObjectSchemaNewAssociationResponseFromObjectTypeMic                          ObjectSchemaNewAssociationResponseFromObjectType = "MIC"
	ObjectSchemaNewAssociationResponseFromObjectTypeContentAudit                 ObjectSchemaNewAssociationResponseFromObjectType = "CONTENT_AUDIT"
	ObjectSchemaNewAssociationResponseFromObjectTypeContentAuditPage             ObjectSchemaNewAssociationResponseFromObjectType = "CONTENT_AUDIT_PAGE"
	ObjectSchemaNewAssociationResponseFromObjectTypePlaylistFolder               ObjectSchemaNewAssociationResponseFromObjectType = "PLAYLIST_FOLDER"
	ObjectSchemaNewAssociationResponseFromObjectTypeLead                         ObjectSchemaNewAssociationResponseFromObjectType = "LEAD"
	ObjectSchemaNewAssociationResponseFromObjectTypeAbandonedCart                ObjectSchemaNewAssociationResponseFromObjectType = "ABANDONED_CART"
	ObjectSchemaNewAssociationResponseFromObjectTypeExternalWebURL               ObjectSchemaNewAssociationResponseFromObjectType = "EXTERNAL_WEB_URL"
	ObjectSchemaNewAssociationResponseFromObjectTypeView                         ObjectSchemaNewAssociationResponseFromObjectType = "VIEW"
	ObjectSchemaNewAssociationResponseFromObjectTypeViewBlock                    ObjectSchemaNewAssociationResponseFromObjectType = "VIEW_BLOCK"
	ObjectSchemaNewAssociationResponseFromObjectTypeRoster                       ObjectSchemaNewAssociationResponseFromObjectType = "ROSTER"
	ObjectSchemaNewAssociationResponseFromObjectTypeCart                         ObjectSchemaNewAssociationResponseFromObjectType = "CART"
	ObjectSchemaNewAssociationResponseFromObjectTypeAutomationPlatformFlowAction ObjectSchemaNewAssociationResponseFromObjectType = "AUTOMATION_PLATFORM_FLOW_ACTION"
	ObjectSchemaNewAssociationResponseFromObjectTypeSocialProfile                ObjectSchemaNewAssociationResponseFromObjectType = "SOCIAL_PROFILE"
	ObjectSchemaNewAssociationResponseFromObjectTypePartnerClient                ObjectSchemaNewAssociationResponseFromObjectType = "PARTNER_CLIENT"
	ObjectSchemaNewAssociationResponseFromObjectTypeRosterMember                 ObjectSchemaNewAssociationResponseFromObjectType = "ROSTER_MEMBER"
	ObjectSchemaNewAssociationResponseFromObjectTypeMarketingEventAttendance     ObjectSchemaNewAssociationResponseFromObjectType = "MARKETING_EVENT_ATTENDANCE"
	ObjectSchemaNewAssociationResponseFromObjectTypeAllPages                     ObjectSchemaNewAssociationResponseFromObjectType = "ALL_PAGES"
	ObjectSchemaNewAssociationResponseFromObjectTypeAIForecast                   ObjectSchemaNewAssociationResponseFromObjectType = "AI_FORECAST"
	ObjectSchemaNewAssociationResponseFromObjectTypeCRMPipelinesDummyType        ObjectSchemaNewAssociationResponseFromObjectType = "CRM_PIPELINES_DUMMY_TYPE"
	ObjectSchemaNewAssociationResponseFromObjectTypeKnowledgeArticle             ObjectSchemaNewAssociationResponseFromObjectType = "KNOWLEDGE_ARTICLE"
	ObjectSchemaNewAssociationResponseFromObjectTypePropertyInfo                 ObjectSchemaNewAssociationResponseFromObjectType = "PROPERTY_INFO"
	ObjectSchemaNewAssociationResponseFromObjectTypeDataPrivacyConsent           ObjectSchemaNewAssociationResponseFromObjectType = "DATA_PRIVACY_CONSENT"
	ObjectSchemaNewAssociationResponseFromObjectTypeGoalTemplate                 ObjectSchemaNewAssociationResponseFromObjectType = "GOAL_TEMPLATE"
	ObjectSchemaNewAssociationResponseFromObjectTypeScoreConfiguration           ObjectSchemaNewAssociationResponseFromObjectType = "SCORE_CONFIGURATION"
	ObjectSchemaNewAssociationResponseFromObjectTypeAudience                     ObjectSchemaNewAssociationResponseFromObjectType = "AUDIENCE"
	ObjectSchemaNewAssociationResponseFromObjectTypePartnerClientRevenue         ObjectSchemaNewAssociationResponseFromObjectType = "PARTNER_CLIENT_REVENUE"
	ObjectSchemaNewAssociationResponseFromObjectTypeAutomationJourney            ObjectSchemaNewAssociationResponseFromObjectType = "AUTOMATION_JOURNEY"
	ObjectSchemaNewAssociationResponseFromObjectTypeComboEventConfiguration      ObjectSchemaNewAssociationResponseFromObjectType = "COMBO_EVENT_CONFIGURATION"
	ObjectSchemaNewAssociationResponseFromObjectTypeCRMObjectsDummyType          ObjectSchemaNewAssociationResponseFromObjectType = "CRM_OBJECTS_DUMMY_TYPE"
	ObjectSchemaNewAssociationResponseFromObjectTypeCaseStudy                    ObjectSchemaNewAssociationResponseFromObjectType = "CASE_STUDY"
	ObjectSchemaNewAssociationResponseFromObjectTypeService                      ObjectSchemaNewAssociationResponseFromObjectType = "SERVICE"
	ObjectSchemaNewAssociationResponseFromObjectTypePodcastEpisode               ObjectSchemaNewAssociationResponseFromObjectType = "PODCAST_EPISODE"
	ObjectSchemaNewAssociationResponseFromObjectTypePartnerService               ObjectSchemaNewAssociationResponseFromObjectType = "PARTNER_SERVICE"
	ObjectSchemaNewAssociationResponseFromObjectTypeUnknown                      ObjectSchemaNewAssociationResponseFromObjectType = "UNKNOWN"
)

// The name of the destination object type (e.g,. "DEAL" or "QUOTE").
type ObjectSchemaNewAssociationResponseToObjectType string

const (
	ObjectSchemaNewAssociationResponseToObjectTypeContact                      ObjectSchemaNewAssociationResponseToObjectType = "CONTACT"
	ObjectSchemaNewAssociationResponseToObjectTypeCompany                      ObjectSchemaNewAssociationResponseToObjectType = "COMPANY"
	ObjectSchemaNewAssociationResponseToObjectTypeDeal                         ObjectSchemaNewAssociationResponseToObjectType = "DEAL"
	ObjectSchemaNewAssociationResponseToObjectTypeEngagement                   ObjectSchemaNewAssociationResponseToObjectType = "ENGAGEMENT"
	ObjectSchemaNewAssociationResponseToObjectTypeTicket                       ObjectSchemaNewAssociationResponseToObjectType = "TICKET"
	ObjectSchemaNewAssociationResponseToObjectTypeOwner                        ObjectSchemaNewAssociationResponseToObjectType = "OWNER"
	ObjectSchemaNewAssociationResponseToObjectTypeProduct                      ObjectSchemaNewAssociationResponseToObjectType = "PRODUCT"
	ObjectSchemaNewAssociationResponseToObjectTypeLineItem                     ObjectSchemaNewAssociationResponseToObjectType = "LINE_ITEM"
	ObjectSchemaNewAssociationResponseToObjectTypeBetDeliverableService        ObjectSchemaNewAssociationResponseToObjectType = "BET_DELIVERABLE_SERVICE"
	ObjectSchemaNewAssociationResponseToObjectTypeContent                      ObjectSchemaNewAssociationResponseToObjectType = "CONTENT"
	ObjectSchemaNewAssociationResponseToObjectTypeConversation                 ObjectSchemaNewAssociationResponseToObjectType = "CONVERSATION"
	ObjectSchemaNewAssociationResponseToObjectTypeBetAlert                     ObjectSchemaNewAssociationResponseToObjectType = "BET_ALERT"
	ObjectSchemaNewAssociationResponseToObjectTypePortal                       ObjectSchemaNewAssociationResponseToObjectType = "PORTAL"
	ObjectSchemaNewAssociationResponseToObjectTypeQuote                        ObjectSchemaNewAssociationResponseToObjectType = "QUOTE"
	ObjectSchemaNewAssociationResponseToObjectTypeFormSubmissionInbounddb      ObjectSchemaNewAssociationResponseToObjectType = "FORM_SUBMISSION_INBOUNDDB"
	ObjectSchemaNewAssociationResponseToObjectTypeQuota                        ObjectSchemaNewAssociationResponseToObjectType = "QUOTA"
	ObjectSchemaNewAssociationResponseToObjectTypeUnsubscribe                  ObjectSchemaNewAssociationResponseToObjectType = "UNSUBSCRIBE"
	ObjectSchemaNewAssociationResponseToObjectTypeCommunication                ObjectSchemaNewAssociationResponseToObjectType = "COMMUNICATION"
	ObjectSchemaNewAssociationResponseToObjectTypeFeedbackSubmission           ObjectSchemaNewAssociationResponseToObjectType = "FEEDBACK_SUBMISSION"
	ObjectSchemaNewAssociationResponseToObjectTypeAttribution                  ObjectSchemaNewAssociationResponseToObjectType = "ATTRIBUTION"
	ObjectSchemaNewAssociationResponseToObjectTypeSalesforceSyncError          ObjectSchemaNewAssociationResponseToObjectType = "SALESFORCE_SYNC_ERROR"
	ObjectSchemaNewAssociationResponseToObjectTypeRestorableCRMObject          ObjectSchemaNewAssociationResponseToObjectType = "RESTORABLE_CRM_OBJECT"
	ObjectSchemaNewAssociationResponseToObjectTypeHub                          ObjectSchemaNewAssociationResponseToObjectType = "HUB"
	ObjectSchemaNewAssociationResponseToObjectTypeLandingPage                  ObjectSchemaNewAssociationResponseToObjectType = "LANDING_PAGE"
	ObjectSchemaNewAssociationResponseToObjectTypeProductOrFolder              ObjectSchemaNewAssociationResponseToObjectType = "PRODUCT_OR_FOLDER"
	ObjectSchemaNewAssociationResponseToObjectTypeTask                         ObjectSchemaNewAssociationResponseToObjectType = "TASK"
	ObjectSchemaNewAssociationResponseToObjectTypeForm                         ObjectSchemaNewAssociationResponseToObjectType = "FORM"
	ObjectSchemaNewAssociationResponseToObjectTypeMarketingEmail               ObjectSchemaNewAssociationResponseToObjectType = "MARKETING_EMAIL"
	ObjectSchemaNewAssociationResponseToObjectTypeAdAccount                    ObjectSchemaNewAssociationResponseToObjectType = "AD_ACCOUNT"
	ObjectSchemaNewAssociationResponseToObjectTypeAdCampaign                   ObjectSchemaNewAssociationResponseToObjectType = "AD_CAMPAIGN"
	ObjectSchemaNewAssociationResponseToObjectTypeAdGroup                      ObjectSchemaNewAssociationResponseToObjectType = "AD_GROUP"
	ObjectSchemaNewAssociationResponseToObjectTypeAd                           ObjectSchemaNewAssociationResponseToObjectType = "AD"
	ObjectSchemaNewAssociationResponseToObjectTypeKeyword                      ObjectSchemaNewAssociationResponseToObjectType = "KEYWORD"
	ObjectSchemaNewAssociationResponseToObjectTypeCampaign                     ObjectSchemaNewAssociationResponseToObjectType = "CAMPAIGN"
	ObjectSchemaNewAssociationResponseToObjectTypeSocialChannel                ObjectSchemaNewAssociationResponseToObjectType = "SOCIAL_CHANNEL"
	ObjectSchemaNewAssociationResponseToObjectTypeSocialPost                   ObjectSchemaNewAssociationResponseToObjectType = "SOCIAL_POST"
	ObjectSchemaNewAssociationResponseToObjectTypeSitePage                     ObjectSchemaNewAssociationResponseToObjectType = "SITE_PAGE"
	ObjectSchemaNewAssociationResponseToObjectTypeBlogPost                     ObjectSchemaNewAssociationResponseToObjectType = "BLOG_POST"
	ObjectSchemaNewAssociationResponseToObjectTypeImport                       ObjectSchemaNewAssociationResponseToObjectType = "IMPORT"
	ObjectSchemaNewAssociationResponseToObjectTypeExport                       ObjectSchemaNewAssociationResponseToObjectType = "EXPORT"
	ObjectSchemaNewAssociationResponseToObjectTypeCta                          ObjectSchemaNewAssociationResponseToObjectType = "CTA"
	ObjectSchemaNewAssociationResponseToObjectTypeTaskTemplate                 ObjectSchemaNewAssociationResponseToObjectType = "TASK_TEMPLATE"
	ObjectSchemaNewAssociationResponseToObjectTypeAutomationPlatformFlow       ObjectSchemaNewAssociationResponseToObjectType = "AUTOMATION_PLATFORM_FLOW"
	ObjectSchemaNewAssociationResponseToObjectTypeObjectList                   ObjectSchemaNewAssociationResponseToObjectType = "OBJECT_LIST"
	ObjectSchemaNewAssociationResponseToObjectTypeNote                         ObjectSchemaNewAssociationResponseToObjectType = "NOTE"
	ObjectSchemaNewAssociationResponseToObjectTypeMeetingEvent                 ObjectSchemaNewAssociationResponseToObjectType = "MEETING_EVENT"
	ObjectSchemaNewAssociationResponseToObjectTypeCall                         ObjectSchemaNewAssociationResponseToObjectType = "CALL"
	ObjectSchemaNewAssociationResponseToObjectTypeEmail                        ObjectSchemaNewAssociationResponseToObjectType = "EMAIL"
	ObjectSchemaNewAssociationResponseToObjectTypePublishingTask               ObjectSchemaNewAssociationResponseToObjectType = "PUBLISHING_TASK"
	ObjectSchemaNewAssociationResponseToObjectTypeConversationSession          ObjectSchemaNewAssociationResponseToObjectType = "CONVERSATION_SESSION"
	ObjectSchemaNewAssociationResponseToObjectTypeContactCreateAttribution     ObjectSchemaNewAssociationResponseToObjectType = "CONTACT_CREATE_ATTRIBUTION"
	ObjectSchemaNewAssociationResponseToObjectTypeInvoice                      ObjectSchemaNewAssociationResponseToObjectType = "INVOICE"
	ObjectSchemaNewAssociationResponseToObjectTypeMarketingEvent               ObjectSchemaNewAssociationResponseToObjectType = "MARKETING_EVENT"
	ObjectSchemaNewAssociationResponseToObjectTypeConversationInbox            ObjectSchemaNewAssociationResponseToObjectType = "CONVERSATION_INBOX"
	ObjectSchemaNewAssociationResponseToObjectTypeChatflow                     ObjectSchemaNewAssociationResponseToObjectType = "CHATFLOW"
	ObjectSchemaNewAssociationResponseToObjectTypeMediaBridge                  ObjectSchemaNewAssociationResponseToObjectType = "MEDIA_BRIDGE"
	ObjectSchemaNewAssociationResponseToObjectTypeSequence                     ObjectSchemaNewAssociationResponseToObjectType = "SEQUENCE"
	ObjectSchemaNewAssociationResponseToObjectTypeSequenceStep                 ObjectSchemaNewAssociationResponseToObjectType = "SEQUENCE_STEP"
	ObjectSchemaNewAssociationResponseToObjectTypeForecast                     ObjectSchemaNewAssociationResponseToObjectType = "FORECAST"
	ObjectSchemaNewAssociationResponseToObjectTypeSnippet                      ObjectSchemaNewAssociationResponseToObjectType = "SNIPPET"
	ObjectSchemaNewAssociationResponseToObjectTypeTemplate                     ObjectSchemaNewAssociationResponseToObjectType = "TEMPLATE"
	ObjectSchemaNewAssociationResponseToObjectTypeDealCreateAttribution        ObjectSchemaNewAssociationResponseToObjectType = "DEAL_CREATE_ATTRIBUTION"
	ObjectSchemaNewAssociationResponseToObjectTypeQuoteTemplate                ObjectSchemaNewAssociationResponseToObjectType = "QUOTE_TEMPLATE"
	ObjectSchemaNewAssociationResponseToObjectTypeQuoteModule                  ObjectSchemaNewAssociationResponseToObjectType = "QUOTE_MODULE"
	ObjectSchemaNewAssociationResponseToObjectTypeQuoteModuleField             ObjectSchemaNewAssociationResponseToObjectType = "QUOTE_MODULE_FIELD"
	ObjectSchemaNewAssociationResponseToObjectTypeQuoteField                   ObjectSchemaNewAssociationResponseToObjectType = "QUOTE_FIELD"
	ObjectSchemaNewAssociationResponseToObjectTypeSequenceEnrollment           ObjectSchemaNewAssociationResponseToObjectType = "SEQUENCE_ENROLLMENT"
	ObjectSchemaNewAssociationResponseToObjectTypeSubscription                 ObjectSchemaNewAssociationResponseToObjectType = "SUBSCRIPTION"
	ObjectSchemaNewAssociationResponseToObjectTypeAcceptanceTest               ObjectSchemaNewAssociationResponseToObjectType = "ACCEPTANCE_TEST"
	ObjectSchemaNewAssociationResponseToObjectTypeSocialBroadcast              ObjectSchemaNewAssociationResponseToObjectType = "SOCIAL_BROADCAST"
	ObjectSchemaNewAssociationResponseToObjectTypeDealSplit                    ObjectSchemaNewAssociationResponseToObjectType = "DEAL_SPLIT"
	ObjectSchemaNewAssociationResponseToObjectTypeDealRegistration             ObjectSchemaNewAssociationResponseToObjectType = "DEAL_REGISTRATION"
	ObjectSchemaNewAssociationResponseToObjectTypeGoalTarget                   ObjectSchemaNewAssociationResponseToObjectType = "GOAL_TARGET"
	ObjectSchemaNewAssociationResponseToObjectTypeGoalTargetGroup              ObjectSchemaNewAssociationResponseToObjectType = "GOAL_TARGET_GROUP"
	ObjectSchemaNewAssociationResponseToObjectTypePortalObjectSyncMessage      ObjectSchemaNewAssociationResponseToObjectType = "PORTAL_OBJECT_SYNC_MESSAGE"
	ObjectSchemaNewAssociationResponseToObjectTypeFileManagerFile              ObjectSchemaNewAssociationResponseToObjectType = "FILE_MANAGER_FILE"
	ObjectSchemaNewAssociationResponseToObjectTypeFileManagerFolder            ObjectSchemaNewAssociationResponseToObjectType = "FILE_MANAGER_FOLDER"
	ObjectSchemaNewAssociationResponseToObjectTypeSequenceStepEnrollment       ObjectSchemaNewAssociationResponseToObjectType = "SEQUENCE_STEP_ENROLLMENT"
	ObjectSchemaNewAssociationResponseToObjectTypeApproval                     ObjectSchemaNewAssociationResponseToObjectType = "APPROVAL"
	ObjectSchemaNewAssociationResponseToObjectTypeApprovalStep                 ObjectSchemaNewAssociationResponseToObjectType = "APPROVAL_STEP"
	ObjectSchemaNewAssociationResponseToObjectTypeCtaVariant                   ObjectSchemaNewAssociationResponseToObjectType = "CTA_VARIANT"
	ObjectSchemaNewAssociationResponseToObjectTypeSalesDocument                ObjectSchemaNewAssociationResponseToObjectType = "SALES_DOCUMENT"
	ObjectSchemaNewAssociationResponseToObjectTypeDiscount                     ObjectSchemaNewAssociationResponseToObjectType = "DISCOUNT"
	ObjectSchemaNewAssociationResponseToObjectTypeFee                          ObjectSchemaNewAssociationResponseToObjectType = "FEE"
	ObjectSchemaNewAssociationResponseToObjectTypeTax                          ObjectSchemaNewAssociationResponseToObjectType = "TAX"
	ObjectSchemaNewAssociationResponseToObjectTypeMarketingCalendar            ObjectSchemaNewAssociationResponseToObjectType = "MARKETING_CALENDAR"
	ObjectSchemaNewAssociationResponseToObjectTypePermissionsTesting           ObjectSchemaNewAssociationResponseToObjectType = "PERMISSIONS_TESTING"
	ObjectSchemaNewAssociationResponseToObjectTypePrivacyScannerCookie         ObjectSchemaNewAssociationResponseToObjectType = "PRIVACY_SCANNER_COOKIE"
	ObjectSchemaNewAssociationResponseToObjectTypeDataSyncState                ObjectSchemaNewAssociationResponseToObjectType = "DATA_SYNC_STATE"
	ObjectSchemaNewAssociationResponseToObjectTypeWebInteractive               ObjectSchemaNewAssociationResponseToObjectType = "WEB_INTERACTIVE"
	ObjectSchemaNewAssociationResponseToObjectTypePlaybook                     ObjectSchemaNewAssociationResponseToObjectType = "PLAYBOOK"
	ObjectSchemaNewAssociationResponseToObjectTypeFolder                       ObjectSchemaNewAssociationResponseToObjectType = "FOLDER"
	ObjectSchemaNewAssociationResponseToObjectTypePlaybookQuestion             ObjectSchemaNewAssociationResponseToObjectType = "PLAYBOOK_QUESTION"
	ObjectSchemaNewAssociationResponseToObjectTypePlaybookSubmission           ObjectSchemaNewAssociationResponseToObjectType = "PLAYBOOK_SUBMISSION"
	ObjectSchemaNewAssociationResponseToObjectTypePlaybookSubmissionAnswer     ObjectSchemaNewAssociationResponseToObjectType = "PLAYBOOK_SUBMISSION_ANSWER"
	ObjectSchemaNewAssociationResponseToObjectTypeCommercePayment              ObjectSchemaNewAssociationResponseToObjectType = "COMMERCE_PAYMENT"
	ObjectSchemaNewAssociationResponseToObjectTypeGscProperty                  ObjectSchemaNewAssociationResponseToObjectType = "GSC_PROPERTY"
	ObjectSchemaNewAssociationResponseToObjectTypeSoxProtectedDummyType        ObjectSchemaNewAssociationResponseToObjectType = "SOX_PROTECTED_DUMMY_TYPE"
	ObjectSchemaNewAssociationResponseToObjectTypeBlogListingPage              ObjectSchemaNewAssociationResponseToObjectType = "BLOG_LISTING_PAGE"
	ObjectSchemaNewAssociationResponseToObjectTypeQuarantinedSubmission        ObjectSchemaNewAssociationResponseToObjectType = "QUARANTINED_SUBMISSION"
	ObjectSchemaNewAssociationResponseToObjectTypePaymentSchedule              ObjectSchemaNewAssociationResponseToObjectType = "PAYMENT_SCHEDULE"
	ObjectSchemaNewAssociationResponseToObjectTypePaymentScheduleInstallment   ObjectSchemaNewAssociationResponseToObjectType = "PAYMENT_SCHEDULE_INSTALLMENT"
	ObjectSchemaNewAssociationResponseToObjectTypeMarketingCampaignUtm         ObjectSchemaNewAssociationResponseToObjectType = "MARKETING_CAMPAIGN_UTM"
	ObjectSchemaNewAssociationResponseToObjectTypeDiscountTemplate             ObjectSchemaNewAssociationResponseToObjectType = "DISCOUNT_TEMPLATE"
	ObjectSchemaNewAssociationResponseToObjectTypeDiscountCode                 ObjectSchemaNewAssociationResponseToObjectType = "DISCOUNT_CODE"
	ObjectSchemaNewAssociationResponseToObjectTypeFeedbackSurvey               ObjectSchemaNewAssociationResponseToObjectType = "FEEDBACK_SURVEY"
	ObjectSchemaNewAssociationResponseToObjectTypeCmsURL                       ObjectSchemaNewAssociationResponseToObjectType = "CMS_URL"
	ObjectSchemaNewAssociationResponseToObjectTypeSalesTask                    ObjectSchemaNewAssociationResponseToObjectType = "SALES_TASK"
	ObjectSchemaNewAssociationResponseToObjectTypeSalesWorkload                ObjectSchemaNewAssociationResponseToObjectType = "SALES_WORKLOAD"
	ObjectSchemaNewAssociationResponseToObjectTypeUser                         ObjectSchemaNewAssociationResponseToObjectType = "USER"
	ObjectSchemaNewAssociationResponseToObjectTypePostalMail                   ObjectSchemaNewAssociationResponseToObjectType = "POSTAL_MAIL"
	ObjectSchemaNewAssociationResponseToObjectTypeSchemasBackendTest           ObjectSchemaNewAssociationResponseToObjectType = "SCHEMAS_BACKEND_TEST"
	ObjectSchemaNewAssociationResponseToObjectTypePaymentLink                  ObjectSchemaNewAssociationResponseToObjectType = "PAYMENT_LINK"
	ObjectSchemaNewAssociationResponseToObjectTypeSubmissionTag                ObjectSchemaNewAssociationResponseToObjectType = "SUBMISSION_TAG"
	ObjectSchemaNewAssociationResponseToObjectTypeCampaignStep                 ObjectSchemaNewAssociationResponseToObjectType = "CAMPAIGN_STEP"
	ObjectSchemaNewAssociationResponseToObjectTypeSchedulingPage               ObjectSchemaNewAssociationResponseToObjectType = "SCHEDULING_PAGE"
	ObjectSchemaNewAssociationResponseToObjectTypeSoxProtectedTestType         ObjectSchemaNewAssociationResponseToObjectType = "SOX_PROTECTED_TEST_TYPE"
	ObjectSchemaNewAssociationResponseToObjectTypeOrder                        ObjectSchemaNewAssociationResponseToObjectType = "ORDER"
	ObjectSchemaNewAssociationResponseToObjectTypeMarketingSMS                 ObjectSchemaNewAssociationResponseToObjectType = "MARKETING_SMS"
	ObjectSchemaNewAssociationResponseToObjectTypePartnerAccount               ObjectSchemaNewAssociationResponseToObjectType = "PARTNER_ACCOUNT"
	ObjectSchemaNewAssociationResponseToObjectTypeCampaignTemplate             ObjectSchemaNewAssociationResponseToObjectType = "CAMPAIGN_TEMPLATE"
	ObjectSchemaNewAssociationResponseToObjectTypeCampaignTemplateStep         ObjectSchemaNewAssociationResponseToObjectType = "CAMPAIGN_TEMPLATE_STEP"
	ObjectSchemaNewAssociationResponseToObjectTypePlaylist                     ObjectSchemaNewAssociationResponseToObjectType = "PLAYLIST"
	ObjectSchemaNewAssociationResponseToObjectTypeClip                         ObjectSchemaNewAssociationResponseToObjectType = "CLIP"
	ObjectSchemaNewAssociationResponseToObjectTypeCampaignBudgetItem           ObjectSchemaNewAssociationResponseToObjectType = "CAMPAIGN_BUDGET_ITEM"
	ObjectSchemaNewAssociationResponseToObjectTypeCampaignSpendItem            ObjectSchemaNewAssociationResponseToObjectType = "CAMPAIGN_SPEND_ITEM"
	ObjectSchemaNewAssociationResponseToObjectTypeMic                          ObjectSchemaNewAssociationResponseToObjectType = "MIC"
	ObjectSchemaNewAssociationResponseToObjectTypeContentAudit                 ObjectSchemaNewAssociationResponseToObjectType = "CONTENT_AUDIT"
	ObjectSchemaNewAssociationResponseToObjectTypeContentAuditPage             ObjectSchemaNewAssociationResponseToObjectType = "CONTENT_AUDIT_PAGE"
	ObjectSchemaNewAssociationResponseToObjectTypePlaylistFolder               ObjectSchemaNewAssociationResponseToObjectType = "PLAYLIST_FOLDER"
	ObjectSchemaNewAssociationResponseToObjectTypeLead                         ObjectSchemaNewAssociationResponseToObjectType = "LEAD"
	ObjectSchemaNewAssociationResponseToObjectTypeAbandonedCart                ObjectSchemaNewAssociationResponseToObjectType = "ABANDONED_CART"
	ObjectSchemaNewAssociationResponseToObjectTypeExternalWebURL               ObjectSchemaNewAssociationResponseToObjectType = "EXTERNAL_WEB_URL"
	ObjectSchemaNewAssociationResponseToObjectTypeView                         ObjectSchemaNewAssociationResponseToObjectType = "VIEW"
	ObjectSchemaNewAssociationResponseToObjectTypeViewBlock                    ObjectSchemaNewAssociationResponseToObjectType = "VIEW_BLOCK"
	ObjectSchemaNewAssociationResponseToObjectTypeRoster                       ObjectSchemaNewAssociationResponseToObjectType = "ROSTER"
	ObjectSchemaNewAssociationResponseToObjectTypeCart                         ObjectSchemaNewAssociationResponseToObjectType = "CART"
	ObjectSchemaNewAssociationResponseToObjectTypeAutomationPlatformFlowAction ObjectSchemaNewAssociationResponseToObjectType = "AUTOMATION_PLATFORM_FLOW_ACTION"
	ObjectSchemaNewAssociationResponseToObjectTypeSocialProfile                ObjectSchemaNewAssociationResponseToObjectType = "SOCIAL_PROFILE"
	ObjectSchemaNewAssociationResponseToObjectTypePartnerClient                ObjectSchemaNewAssociationResponseToObjectType = "PARTNER_CLIENT"
	ObjectSchemaNewAssociationResponseToObjectTypeRosterMember                 ObjectSchemaNewAssociationResponseToObjectType = "ROSTER_MEMBER"
	ObjectSchemaNewAssociationResponseToObjectTypeMarketingEventAttendance     ObjectSchemaNewAssociationResponseToObjectType = "MARKETING_EVENT_ATTENDANCE"
	ObjectSchemaNewAssociationResponseToObjectTypeAllPages                     ObjectSchemaNewAssociationResponseToObjectType = "ALL_PAGES"
	ObjectSchemaNewAssociationResponseToObjectTypeAIForecast                   ObjectSchemaNewAssociationResponseToObjectType = "AI_FORECAST"
	ObjectSchemaNewAssociationResponseToObjectTypeCRMPipelinesDummyType        ObjectSchemaNewAssociationResponseToObjectType = "CRM_PIPELINES_DUMMY_TYPE"
	ObjectSchemaNewAssociationResponseToObjectTypeKnowledgeArticle             ObjectSchemaNewAssociationResponseToObjectType = "KNOWLEDGE_ARTICLE"
	ObjectSchemaNewAssociationResponseToObjectTypePropertyInfo                 ObjectSchemaNewAssociationResponseToObjectType = "PROPERTY_INFO"
	ObjectSchemaNewAssociationResponseToObjectTypeDataPrivacyConsent           ObjectSchemaNewAssociationResponseToObjectType = "DATA_PRIVACY_CONSENT"
	ObjectSchemaNewAssociationResponseToObjectTypeGoalTemplate                 ObjectSchemaNewAssociationResponseToObjectType = "GOAL_TEMPLATE"
	ObjectSchemaNewAssociationResponseToObjectTypeScoreConfiguration           ObjectSchemaNewAssociationResponseToObjectType = "SCORE_CONFIGURATION"
	ObjectSchemaNewAssociationResponseToObjectTypeAudience                     ObjectSchemaNewAssociationResponseToObjectType = "AUDIENCE"
	ObjectSchemaNewAssociationResponseToObjectTypePartnerClientRevenue         ObjectSchemaNewAssociationResponseToObjectType = "PARTNER_CLIENT_REVENUE"
	ObjectSchemaNewAssociationResponseToObjectTypeAutomationJourney            ObjectSchemaNewAssociationResponseToObjectType = "AUTOMATION_JOURNEY"
	ObjectSchemaNewAssociationResponseToObjectTypeComboEventConfiguration      ObjectSchemaNewAssociationResponseToObjectType = "COMBO_EVENT_CONFIGURATION"
	ObjectSchemaNewAssociationResponseToObjectTypeCRMObjectsDummyType          ObjectSchemaNewAssociationResponseToObjectType = "CRM_OBJECTS_DUMMY_TYPE"
	ObjectSchemaNewAssociationResponseToObjectTypeCaseStudy                    ObjectSchemaNewAssociationResponseToObjectType = "CASE_STUDY"
	ObjectSchemaNewAssociationResponseToObjectTypeService                      ObjectSchemaNewAssociationResponseToObjectType = "SERVICE"
	ObjectSchemaNewAssociationResponseToObjectTypePodcastEpisode               ObjectSchemaNewAssociationResponseToObjectType = "PODCAST_EPISODE"
	ObjectSchemaNewAssociationResponseToObjectTypePartnerService               ObjectSchemaNewAssociationResponseToObjectType = "PARTNER_SERVICE"
	ObjectSchemaNewAssociationResponseToObjectTypeUnknown                      ObjectSchemaNewAssociationResponseToObjectType = "UNKNOWN"
)

type ObjectSchemaNewParams struct {
	// Defines a new object type, its properties, and associations.
	ObjectSchemaEgg ObjectSchemaEggParam
	paramObj
}

func (r ObjectSchemaNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.ObjectSchemaEgg)
}
func (r *ObjectSchemaNewParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.ObjectSchemaEgg)
}

type ObjectSchemaUpdateParams struct {
	// Defines attributes to update on an object type.
	ObjectTypeDefinitionPatch ObjectTypeDefinitionPatchParam
	paramObj
}

func (r ObjectSchemaUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.ObjectTypeDefinitionPatch)
}
func (r *ObjectSchemaUpdateParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.ObjectTypeDefinitionPatch)
}

type ObjectSchemaListParams struct {
	// Whether to return only results that have been archived.
	Archived param.Opt[bool] `query:"archived,omitzero" json:"-"`
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
	return json.Unmarshal(data, &r.AssociationDefinitionEgg)
}

type ObjectSchemaDeleteAssociationParams struct {
	ObjectType string `path:"objectType,required" json:"-"`
	paramObj
}
