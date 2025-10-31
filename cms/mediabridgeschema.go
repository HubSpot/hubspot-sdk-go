// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cms

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/stainless-sdks/hubspot-sdk-go/crm"
	"github.com/stainless-sdks/hubspot-sdk-go/internal/apijson"
	shimjson "github.com/stainless-sdks/hubspot-sdk-go/internal/encoding/json"
	"github.com/stainless-sdks/hubspot-sdk-go/internal/requestconfig"
	"github.com/stainless-sdks/hubspot-sdk-go/option"
	"github.com/stainless-sdks/hubspot-sdk-go/packages/respjson"
	"github.com/stainless-sdks/hubspot-sdk-go/shared"
)

// MediaBridgeSchemaService contains methods and other services that help with
// interacting with the Hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMediaBridgeSchemaService] method instead.
type MediaBridgeSchemaService struct {
	Options []option.RequestOption
}

// NewMediaBridgeSchemaService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewMediaBridgeSchemaService(opts ...option.RequestOption) (r MediaBridgeSchemaService) {
	r = MediaBridgeSchemaService{}
	r.Options = opts
	return
}

// Update the schema for an existing object type
func (r *MediaBridgeSchemaService) Update(ctx context.Context, objectType string, params MediaBridgeSchemaUpdateParams, opts ...option.RequestOption) (res *crm.ObjectTypeDefinition, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.AppID == "" {
		err = errors.New("missing required appId parameter")
		return
	}
	if objectType == "" {
		err = errors.New("missing required objectType parameter")
		return
	}
	path := fmt.Sprintf("media-bridge/v1/%s/schemas/%s", params.AppID, objectType)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return
}

// Get the schemas for all object types.
func (r *MediaBridgeSchemaService) List(ctx context.Context, appID string, opts ...option.RequestOption) (res *shared.CollectionResponseObjectSchemaNoPaging, err error) {
	opts = slices.Concat(r.Options, opts)
	if appID == "" {
		err = errors.New("missing required appId parameter")
		return
	}
	path := fmt.Sprintf("media-bridge/v1/%s/schemas", appID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Create a new association definition for the specified object type.
func (r *MediaBridgeSchemaService) NewAssociation(ctx context.Context, objectType string, params MediaBridgeSchemaNewAssociationParams, opts ...option.RequestOption) (res *MediaBridgeSchemaNewAssociationResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.AppID == "" {
		err = errors.New("missing required appId parameter")
		return
	}
	if objectType == "" {
		err = errors.New("missing required objectType parameter")
		return
	}
	path := fmt.Sprintf("media-bridge/v1/%s/schemas/%s/associations", params.AppID, objectType)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Delete an existing association definition for an object type.
func (r *MediaBridgeSchemaService) DeleteAssociation(ctx context.Context, associationID string, body MediaBridgeSchemaDeleteAssociationParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if body.AppID == "" {
		err = errors.New("missing required appId parameter")
		return
	}
	if body.ObjectType == "" {
		err = errors.New("missing required objectType parameter")
		return
	}
	if associationID == "" {
		err = errors.New("missing required associationId parameter")
		return
	}
	path := fmt.Sprintf("media-bridge/v1/%s/schemas/%s/associations/%s", body.AppID, body.ObjectType, associationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Get the schema for a specified object type.
func (r *MediaBridgeSchemaService) Get(ctx context.Context, objectType string, query MediaBridgeSchemaGetParams, opts ...option.RequestOption) (res *crm.ObjectSchema, err error) {
	opts = slices.Concat(r.Options, opts)
	if query.AppID == "" {
		err = errors.New("missing required appId parameter")
		return
	}
	if objectType == "" {
		err = errors.New("missing required objectType parameter")
		return
	}
	path := fmt.Sprintf("media-bridge/v1/%s/schemas/%s", query.AppID, objectType)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// The definition of an association
type MediaBridgeSchemaNewAssociationResponse struct {
	// The unique ID of the associated object (e.g., a contact ID).
	ID int64 `json:"id,required"`
	// Whether custom labels can be used in the association.
	AllowsCustomLabels bool `json:"allowsCustomLabels,required"`
	// The cardinality from the source object's perspective, either "ONE_TO_ONE" or
	// "ONE_TO_MANY".
	//
	// Any of "ONE_TO_ONE", "ONE_TO_MANY".
	Cardinality MediaBridgeSchemaNewAssociationResponseCardinality `json:"cardinality,required"`
	// The category of the association. Can be: "HUBSPOT_DEFINED", "USER_DEFINED", or
	// "INTEGRATOR_DEFINED"
	//
	// Any of "HUBSPOT_DEFINED", "USER_DEFINED", "INTEGRATOR_DEFINED".
	Category MediaBridgeSchemaNewAssociationResponseCategory `json:"category,required"`
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
	InverseCardinality MediaBridgeSchemaNewAssociationResponseInverseCardinality `json:"inverseCardinality,required"`
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
	FromObjectType MediaBridgeSchemaNewAssociationResponseFromObjectType `json:"fromObjectType"`
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
	ToObjectType MediaBridgeSchemaNewAssociationResponseToObjectType `json:"toObjectType"`
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
func (r MediaBridgeSchemaNewAssociationResponse) RawJSON() string { return r.JSON.raw }
func (r *MediaBridgeSchemaNewAssociationResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The cardinality from the source object's perspective, either "ONE_TO_ONE" or
// "ONE_TO_MANY".
type MediaBridgeSchemaNewAssociationResponseCardinality string

const (
	MediaBridgeSchemaNewAssociationResponseCardinalityOneToOne  MediaBridgeSchemaNewAssociationResponseCardinality = "ONE_TO_ONE"
	MediaBridgeSchemaNewAssociationResponseCardinalityOneToMany MediaBridgeSchemaNewAssociationResponseCardinality = "ONE_TO_MANY"
)

// The category of the association. Can be: "HUBSPOT_DEFINED", "USER_DEFINED", or
// "INTEGRATOR_DEFINED"
type MediaBridgeSchemaNewAssociationResponseCategory string

const (
	MediaBridgeSchemaNewAssociationResponseCategoryHubspotDefined    MediaBridgeSchemaNewAssociationResponseCategory = "HUBSPOT_DEFINED"
	MediaBridgeSchemaNewAssociationResponseCategoryUserDefined       MediaBridgeSchemaNewAssociationResponseCategory = "USER_DEFINED"
	MediaBridgeSchemaNewAssociationResponseCategoryIntegratorDefined MediaBridgeSchemaNewAssociationResponseCategory = "INTEGRATOR_DEFINED"
)

// The cardinality from the destination object's perspective, either "ONE_TO_ONE"
// or "ONE_TO_MANY".
type MediaBridgeSchemaNewAssociationResponseInverseCardinality string

const (
	MediaBridgeSchemaNewAssociationResponseInverseCardinalityOneToOne  MediaBridgeSchemaNewAssociationResponseInverseCardinality = "ONE_TO_ONE"
	MediaBridgeSchemaNewAssociationResponseInverseCardinalityOneToMany MediaBridgeSchemaNewAssociationResponseInverseCardinality = "ONE_TO_MANY"
)

// The name of the source object type (e.g,. "DEAL" or "QUOTE").
type MediaBridgeSchemaNewAssociationResponseFromObjectType string

const (
	MediaBridgeSchemaNewAssociationResponseFromObjectTypeContact                      MediaBridgeSchemaNewAssociationResponseFromObjectType = "CONTACT"
	MediaBridgeSchemaNewAssociationResponseFromObjectTypeCompany                      MediaBridgeSchemaNewAssociationResponseFromObjectType = "COMPANY"
	MediaBridgeSchemaNewAssociationResponseFromObjectTypeDeal                         MediaBridgeSchemaNewAssociationResponseFromObjectType = "DEAL"
	MediaBridgeSchemaNewAssociationResponseFromObjectTypeEngagement                   MediaBridgeSchemaNewAssociationResponseFromObjectType = "ENGAGEMENT"
	MediaBridgeSchemaNewAssociationResponseFromObjectTypeTicket                       MediaBridgeSchemaNewAssociationResponseFromObjectType = "TICKET"
	MediaBridgeSchemaNewAssociationResponseFromObjectTypeOwner                        MediaBridgeSchemaNewAssociationResponseFromObjectType = "OWNER"
	MediaBridgeSchemaNewAssociationResponseFromObjectTypeProduct                      MediaBridgeSchemaNewAssociationResponseFromObjectType = "PRODUCT"
	MediaBridgeSchemaNewAssociationResponseFromObjectTypeLineItem                     MediaBridgeSchemaNewAssociationResponseFromObjectType = "LINE_ITEM"
	MediaBridgeSchemaNewAssociationResponseFromObjectTypeBetDeliverableService        MediaBridgeSchemaNewAssociationResponseFromObjectType = "BET_DELIVERABLE_SERVICE"
	MediaBridgeSchemaNewAssociationResponseFromObjectTypeContent                      MediaBridgeSchemaNewAssociationResponseFromObjectType = "CONTENT"
	MediaBridgeSchemaNewAssociationResponseFromObjectTypeConversation                 MediaBridgeSchemaNewAssociationResponseFromObjectType = "CONVERSATION"
	MediaBridgeSchemaNewAssociationResponseFromObjectTypeBetAlert                     MediaBridgeSchemaNewAssociationResponseFromObjectType = "BET_ALERT"
	MediaBridgeSchemaNewAssociationResponseFromObjectTypePortal                       MediaBridgeSchemaNewAssociationResponseFromObjectType = "PORTAL"
	MediaBridgeSchemaNewAssociationResponseFromObjectTypeQuote                        MediaBridgeSchemaNewAssociationResponseFromObjectType = "QUOTE"
	MediaBridgeSchemaNewAssociationResponseFromObjectTypeFormSubmissionInbounddb      MediaBridgeSchemaNewAssociationResponseFromObjectType = "FORM_SUBMISSION_INBOUNDDB"
	MediaBridgeSchemaNewAssociationResponseFromObjectTypeQuota                        MediaBridgeSchemaNewAssociationResponseFromObjectType = "QUOTA"
	MediaBridgeSchemaNewAssociationResponseFromObjectTypeUnsubscribe                  MediaBridgeSchemaNewAssociationResponseFromObjectType = "UNSUBSCRIBE"
	MediaBridgeSchemaNewAssociationResponseFromObjectTypeCommunication                MediaBridgeSchemaNewAssociationResponseFromObjectType = "COMMUNICATION"
	MediaBridgeSchemaNewAssociationResponseFromObjectTypeFeedbackSubmission           MediaBridgeSchemaNewAssociationResponseFromObjectType = "FEEDBACK_SUBMISSION"
	MediaBridgeSchemaNewAssociationResponseFromObjectTypeAttribution                  MediaBridgeSchemaNewAssociationResponseFromObjectType = "ATTRIBUTION"
	MediaBridgeSchemaNewAssociationResponseFromObjectTypeSalesforceSyncError          MediaBridgeSchemaNewAssociationResponseFromObjectType = "SALESFORCE_SYNC_ERROR"
	MediaBridgeSchemaNewAssociationResponseFromObjectTypeRestorableCRMObject          MediaBridgeSchemaNewAssociationResponseFromObjectType = "RESTORABLE_CRM_OBJECT"
	MediaBridgeSchemaNewAssociationResponseFromObjectTypeHub                          MediaBridgeSchemaNewAssociationResponseFromObjectType = "HUB"
	MediaBridgeSchemaNewAssociationResponseFromObjectTypeLandingPage                  MediaBridgeSchemaNewAssociationResponseFromObjectType = "LANDING_PAGE"
	MediaBridgeSchemaNewAssociationResponseFromObjectTypeProductOrFolder              MediaBridgeSchemaNewAssociationResponseFromObjectType = "PRODUCT_OR_FOLDER"
	MediaBridgeSchemaNewAssociationResponseFromObjectTypeTask                         MediaBridgeSchemaNewAssociationResponseFromObjectType = "TASK"
	MediaBridgeSchemaNewAssociationResponseFromObjectTypeForm                         MediaBridgeSchemaNewAssociationResponseFromObjectType = "FORM"
	MediaBridgeSchemaNewAssociationResponseFromObjectTypeMarketingEmail               MediaBridgeSchemaNewAssociationResponseFromObjectType = "MARKETING_EMAIL"
	MediaBridgeSchemaNewAssociationResponseFromObjectTypeAdAccount                    MediaBridgeSchemaNewAssociationResponseFromObjectType = "AD_ACCOUNT"
	MediaBridgeSchemaNewAssociationResponseFromObjectTypeAdCampaign                   MediaBridgeSchemaNewAssociationResponseFromObjectType = "AD_CAMPAIGN"
	MediaBridgeSchemaNewAssociationResponseFromObjectTypeAdGroup                      MediaBridgeSchemaNewAssociationResponseFromObjectType = "AD_GROUP"
	MediaBridgeSchemaNewAssociationResponseFromObjectTypeAd                           MediaBridgeSchemaNewAssociationResponseFromObjectType = "AD"
	MediaBridgeSchemaNewAssociationResponseFromObjectTypeKeyword                      MediaBridgeSchemaNewAssociationResponseFromObjectType = "KEYWORD"
	MediaBridgeSchemaNewAssociationResponseFromObjectTypeCampaign                     MediaBridgeSchemaNewAssociationResponseFromObjectType = "CAMPAIGN"
	MediaBridgeSchemaNewAssociationResponseFromObjectTypeSocialChannel                MediaBridgeSchemaNewAssociationResponseFromObjectType = "SOCIAL_CHANNEL"
	MediaBridgeSchemaNewAssociationResponseFromObjectTypeSocialPost                   MediaBridgeSchemaNewAssociationResponseFromObjectType = "SOCIAL_POST"
	MediaBridgeSchemaNewAssociationResponseFromObjectTypeSitePage                     MediaBridgeSchemaNewAssociationResponseFromObjectType = "SITE_PAGE"
	MediaBridgeSchemaNewAssociationResponseFromObjectTypeBlogPost                     MediaBridgeSchemaNewAssociationResponseFromObjectType = "BLOG_POST"
	MediaBridgeSchemaNewAssociationResponseFromObjectTypeImport                       MediaBridgeSchemaNewAssociationResponseFromObjectType = "IMPORT"
	MediaBridgeSchemaNewAssociationResponseFromObjectTypeExport                       MediaBridgeSchemaNewAssociationResponseFromObjectType = "EXPORT"
	MediaBridgeSchemaNewAssociationResponseFromObjectTypeCta                          MediaBridgeSchemaNewAssociationResponseFromObjectType = "CTA"
	MediaBridgeSchemaNewAssociationResponseFromObjectTypeTaskTemplate                 MediaBridgeSchemaNewAssociationResponseFromObjectType = "TASK_TEMPLATE"
	MediaBridgeSchemaNewAssociationResponseFromObjectTypeAutomationPlatformFlow       MediaBridgeSchemaNewAssociationResponseFromObjectType = "AUTOMATION_PLATFORM_FLOW"
	MediaBridgeSchemaNewAssociationResponseFromObjectTypeObjectList                   MediaBridgeSchemaNewAssociationResponseFromObjectType = "OBJECT_LIST"
	MediaBridgeSchemaNewAssociationResponseFromObjectTypeNote                         MediaBridgeSchemaNewAssociationResponseFromObjectType = "NOTE"
	MediaBridgeSchemaNewAssociationResponseFromObjectTypeMeetingEvent                 MediaBridgeSchemaNewAssociationResponseFromObjectType = "MEETING_EVENT"
	MediaBridgeSchemaNewAssociationResponseFromObjectTypeCall                         MediaBridgeSchemaNewAssociationResponseFromObjectType = "CALL"
	MediaBridgeSchemaNewAssociationResponseFromObjectTypeEmail                        MediaBridgeSchemaNewAssociationResponseFromObjectType = "EMAIL"
	MediaBridgeSchemaNewAssociationResponseFromObjectTypePublishingTask               MediaBridgeSchemaNewAssociationResponseFromObjectType = "PUBLISHING_TASK"
	MediaBridgeSchemaNewAssociationResponseFromObjectTypeConversationSession          MediaBridgeSchemaNewAssociationResponseFromObjectType = "CONVERSATION_SESSION"
	MediaBridgeSchemaNewAssociationResponseFromObjectTypeContactCreateAttribution     MediaBridgeSchemaNewAssociationResponseFromObjectType = "CONTACT_CREATE_ATTRIBUTION"
	MediaBridgeSchemaNewAssociationResponseFromObjectTypeInvoice                      MediaBridgeSchemaNewAssociationResponseFromObjectType = "INVOICE"
	MediaBridgeSchemaNewAssociationResponseFromObjectTypeMarketingEvent               MediaBridgeSchemaNewAssociationResponseFromObjectType = "MARKETING_EVENT"
	MediaBridgeSchemaNewAssociationResponseFromObjectTypeConversationInbox            MediaBridgeSchemaNewAssociationResponseFromObjectType = "CONVERSATION_INBOX"
	MediaBridgeSchemaNewAssociationResponseFromObjectTypeChatflow                     MediaBridgeSchemaNewAssociationResponseFromObjectType = "CHATFLOW"
	MediaBridgeSchemaNewAssociationResponseFromObjectTypeMediaBridge                  MediaBridgeSchemaNewAssociationResponseFromObjectType = "MEDIA_BRIDGE"
	MediaBridgeSchemaNewAssociationResponseFromObjectTypeSequence                     MediaBridgeSchemaNewAssociationResponseFromObjectType = "SEQUENCE"
	MediaBridgeSchemaNewAssociationResponseFromObjectTypeSequenceStep                 MediaBridgeSchemaNewAssociationResponseFromObjectType = "SEQUENCE_STEP"
	MediaBridgeSchemaNewAssociationResponseFromObjectTypeForecast                     MediaBridgeSchemaNewAssociationResponseFromObjectType = "FORECAST"
	MediaBridgeSchemaNewAssociationResponseFromObjectTypeSnippet                      MediaBridgeSchemaNewAssociationResponseFromObjectType = "SNIPPET"
	MediaBridgeSchemaNewAssociationResponseFromObjectTypeTemplate                     MediaBridgeSchemaNewAssociationResponseFromObjectType = "TEMPLATE"
	MediaBridgeSchemaNewAssociationResponseFromObjectTypeDealCreateAttribution        MediaBridgeSchemaNewAssociationResponseFromObjectType = "DEAL_CREATE_ATTRIBUTION"
	MediaBridgeSchemaNewAssociationResponseFromObjectTypeQuoteTemplate                MediaBridgeSchemaNewAssociationResponseFromObjectType = "QUOTE_TEMPLATE"
	MediaBridgeSchemaNewAssociationResponseFromObjectTypeQuoteModule                  MediaBridgeSchemaNewAssociationResponseFromObjectType = "QUOTE_MODULE"
	MediaBridgeSchemaNewAssociationResponseFromObjectTypeQuoteModuleField             MediaBridgeSchemaNewAssociationResponseFromObjectType = "QUOTE_MODULE_FIELD"
	MediaBridgeSchemaNewAssociationResponseFromObjectTypeQuoteField                   MediaBridgeSchemaNewAssociationResponseFromObjectType = "QUOTE_FIELD"
	MediaBridgeSchemaNewAssociationResponseFromObjectTypeSequenceEnrollment           MediaBridgeSchemaNewAssociationResponseFromObjectType = "SEQUENCE_ENROLLMENT"
	MediaBridgeSchemaNewAssociationResponseFromObjectTypeSubscription                 MediaBridgeSchemaNewAssociationResponseFromObjectType = "SUBSCRIPTION"
	MediaBridgeSchemaNewAssociationResponseFromObjectTypeAcceptanceTest               MediaBridgeSchemaNewAssociationResponseFromObjectType = "ACCEPTANCE_TEST"
	MediaBridgeSchemaNewAssociationResponseFromObjectTypeSocialBroadcast              MediaBridgeSchemaNewAssociationResponseFromObjectType = "SOCIAL_BROADCAST"
	MediaBridgeSchemaNewAssociationResponseFromObjectTypeDealSplit                    MediaBridgeSchemaNewAssociationResponseFromObjectType = "DEAL_SPLIT"
	MediaBridgeSchemaNewAssociationResponseFromObjectTypeDealRegistration             MediaBridgeSchemaNewAssociationResponseFromObjectType = "DEAL_REGISTRATION"
	MediaBridgeSchemaNewAssociationResponseFromObjectTypeGoalTarget                   MediaBridgeSchemaNewAssociationResponseFromObjectType = "GOAL_TARGET"
	MediaBridgeSchemaNewAssociationResponseFromObjectTypeGoalTargetGroup              MediaBridgeSchemaNewAssociationResponseFromObjectType = "GOAL_TARGET_GROUP"
	MediaBridgeSchemaNewAssociationResponseFromObjectTypePortalObjectSyncMessage      MediaBridgeSchemaNewAssociationResponseFromObjectType = "PORTAL_OBJECT_SYNC_MESSAGE"
	MediaBridgeSchemaNewAssociationResponseFromObjectTypeFileManagerFile              MediaBridgeSchemaNewAssociationResponseFromObjectType = "FILE_MANAGER_FILE"
	MediaBridgeSchemaNewAssociationResponseFromObjectTypeFileManagerFolder            MediaBridgeSchemaNewAssociationResponseFromObjectType = "FILE_MANAGER_FOLDER"
	MediaBridgeSchemaNewAssociationResponseFromObjectTypeSequenceStepEnrollment       MediaBridgeSchemaNewAssociationResponseFromObjectType = "SEQUENCE_STEP_ENROLLMENT"
	MediaBridgeSchemaNewAssociationResponseFromObjectTypeApproval                     MediaBridgeSchemaNewAssociationResponseFromObjectType = "APPROVAL"
	MediaBridgeSchemaNewAssociationResponseFromObjectTypeApprovalStep                 MediaBridgeSchemaNewAssociationResponseFromObjectType = "APPROVAL_STEP"
	MediaBridgeSchemaNewAssociationResponseFromObjectTypeCtaVariant                   MediaBridgeSchemaNewAssociationResponseFromObjectType = "CTA_VARIANT"
	MediaBridgeSchemaNewAssociationResponseFromObjectTypeSalesDocument                MediaBridgeSchemaNewAssociationResponseFromObjectType = "SALES_DOCUMENT"
	MediaBridgeSchemaNewAssociationResponseFromObjectTypeDiscount                     MediaBridgeSchemaNewAssociationResponseFromObjectType = "DISCOUNT"
	MediaBridgeSchemaNewAssociationResponseFromObjectTypeFee                          MediaBridgeSchemaNewAssociationResponseFromObjectType = "FEE"
	MediaBridgeSchemaNewAssociationResponseFromObjectTypeTax                          MediaBridgeSchemaNewAssociationResponseFromObjectType = "TAX"
	MediaBridgeSchemaNewAssociationResponseFromObjectTypeMarketingCalendar            MediaBridgeSchemaNewAssociationResponseFromObjectType = "MARKETING_CALENDAR"
	MediaBridgeSchemaNewAssociationResponseFromObjectTypePermissionsTesting           MediaBridgeSchemaNewAssociationResponseFromObjectType = "PERMISSIONS_TESTING"
	MediaBridgeSchemaNewAssociationResponseFromObjectTypePrivacyScannerCookie         MediaBridgeSchemaNewAssociationResponseFromObjectType = "PRIVACY_SCANNER_COOKIE"
	MediaBridgeSchemaNewAssociationResponseFromObjectTypeDataSyncState                MediaBridgeSchemaNewAssociationResponseFromObjectType = "DATA_SYNC_STATE"
	MediaBridgeSchemaNewAssociationResponseFromObjectTypeWebInteractive               MediaBridgeSchemaNewAssociationResponseFromObjectType = "WEB_INTERACTIVE"
	MediaBridgeSchemaNewAssociationResponseFromObjectTypePlaybook                     MediaBridgeSchemaNewAssociationResponseFromObjectType = "PLAYBOOK"
	MediaBridgeSchemaNewAssociationResponseFromObjectTypeFolder                       MediaBridgeSchemaNewAssociationResponseFromObjectType = "FOLDER"
	MediaBridgeSchemaNewAssociationResponseFromObjectTypePlaybookQuestion             MediaBridgeSchemaNewAssociationResponseFromObjectType = "PLAYBOOK_QUESTION"
	MediaBridgeSchemaNewAssociationResponseFromObjectTypePlaybookSubmission           MediaBridgeSchemaNewAssociationResponseFromObjectType = "PLAYBOOK_SUBMISSION"
	MediaBridgeSchemaNewAssociationResponseFromObjectTypePlaybookSubmissionAnswer     MediaBridgeSchemaNewAssociationResponseFromObjectType = "PLAYBOOK_SUBMISSION_ANSWER"
	MediaBridgeSchemaNewAssociationResponseFromObjectTypeCommercePayment              MediaBridgeSchemaNewAssociationResponseFromObjectType = "COMMERCE_PAYMENT"
	MediaBridgeSchemaNewAssociationResponseFromObjectTypeGscProperty                  MediaBridgeSchemaNewAssociationResponseFromObjectType = "GSC_PROPERTY"
	MediaBridgeSchemaNewAssociationResponseFromObjectTypeSoxProtectedDummyType        MediaBridgeSchemaNewAssociationResponseFromObjectType = "SOX_PROTECTED_DUMMY_TYPE"
	MediaBridgeSchemaNewAssociationResponseFromObjectTypeBlogListingPage              MediaBridgeSchemaNewAssociationResponseFromObjectType = "BLOG_LISTING_PAGE"
	MediaBridgeSchemaNewAssociationResponseFromObjectTypeQuarantinedSubmission        MediaBridgeSchemaNewAssociationResponseFromObjectType = "QUARANTINED_SUBMISSION"
	MediaBridgeSchemaNewAssociationResponseFromObjectTypePaymentSchedule              MediaBridgeSchemaNewAssociationResponseFromObjectType = "PAYMENT_SCHEDULE"
	MediaBridgeSchemaNewAssociationResponseFromObjectTypePaymentScheduleInstallment   MediaBridgeSchemaNewAssociationResponseFromObjectType = "PAYMENT_SCHEDULE_INSTALLMENT"
	MediaBridgeSchemaNewAssociationResponseFromObjectTypeMarketingCampaignUtm         MediaBridgeSchemaNewAssociationResponseFromObjectType = "MARKETING_CAMPAIGN_UTM"
	MediaBridgeSchemaNewAssociationResponseFromObjectTypeDiscountTemplate             MediaBridgeSchemaNewAssociationResponseFromObjectType = "DISCOUNT_TEMPLATE"
	MediaBridgeSchemaNewAssociationResponseFromObjectTypeDiscountCode                 MediaBridgeSchemaNewAssociationResponseFromObjectType = "DISCOUNT_CODE"
	MediaBridgeSchemaNewAssociationResponseFromObjectTypeFeedbackSurvey               MediaBridgeSchemaNewAssociationResponseFromObjectType = "FEEDBACK_SURVEY"
	MediaBridgeSchemaNewAssociationResponseFromObjectTypeCmsURL                       MediaBridgeSchemaNewAssociationResponseFromObjectType = "CMS_URL"
	MediaBridgeSchemaNewAssociationResponseFromObjectTypeSalesTask                    MediaBridgeSchemaNewAssociationResponseFromObjectType = "SALES_TASK"
	MediaBridgeSchemaNewAssociationResponseFromObjectTypeSalesWorkload                MediaBridgeSchemaNewAssociationResponseFromObjectType = "SALES_WORKLOAD"
	MediaBridgeSchemaNewAssociationResponseFromObjectTypeUser                         MediaBridgeSchemaNewAssociationResponseFromObjectType = "USER"
	MediaBridgeSchemaNewAssociationResponseFromObjectTypePostalMail                   MediaBridgeSchemaNewAssociationResponseFromObjectType = "POSTAL_MAIL"
	MediaBridgeSchemaNewAssociationResponseFromObjectTypeSchemasBackendTest           MediaBridgeSchemaNewAssociationResponseFromObjectType = "SCHEMAS_BACKEND_TEST"
	MediaBridgeSchemaNewAssociationResponseFromObjectTypePaymentLink                  MediaBridgeSchemaNewAssociationResponseFromObjectType = "PAYMENT_LINK"
	MediaBridgeSchemaNewAssociationResponseFromObjectTypeSubmissionTag                MediaBridgeSchemaNewAssociationResponseFromObjectType = "SUBMISSION_TAG"
	MediaBridgeSchemaNewAssociationResponseFromObjectTypeCampaignStep                 MediaBridgeSchemaNewAssociationResponseFromObjectType = "CAMPAIGN_STEP"
	MediaBridgeSchemaNewAssociationResponseFromObjectTypeSchedulingPage               MediaBridgeSchemaNewAssociationResponseFromObjectType = "SCHEDULING_PAGE"
	MediaBridgeSchemaNewAssociationResponseFromObjectTypeSoxProtectedTestType         MediaBridgeSchemaNewAssociationResponseFromObjectType = "SOX_PROTECTED_TEST_TYPE"
	MediaBridgeSchemaNewAssociationResponseFromObjectTypeOrder                        MediaBridgeSchemaNewAssociationResponseFromObjectType = "ORDER"
	MediaBridgeSchemaNewAssociationResponseFromObjectTypeMarketingSMS                 MediaBridgeSchemaNewAssociationResponseFromObjectType = "MARKETING_SMS"
	MediaBridgeSchemaNewAssociationResponseFromObjectTypePartnerAccount               MediaBridgeSchemaNewAssociationResponseFromObjectType = "PARTNER_ACCOUNT"
	MediaBridgeSchemaNewAssociationResponseFromObjectTypeCampaignTemplate             MediaBridgeSchemaNewAssociationResponseFromObjectType = "CAMPAIGN_TEMPLATE"
	MediaBridgeSchemaNewAssociationResponseFromObjectTypeCampaignTemplateStep         MediaBridgeSchemaNewAssociationResponseFromObjectType = "CAMPAIGN_TEMPLATE_STEP"
	MediaBridgeSchemaNewAssociationResponseFromObjectTypePlaylist                     MediaBridgeSchemaNewAssociationResponseFromObjectType = "PLAYLIST"
	MediaBridgeSchemaNewAssociationResponseFromObjectTypeClip                         MediaBridgeSchemaNewAssociationResponseFromObjectType = "CLIP"
	MediaBridgeSchemaNewAssociationResponseFromObjectTypeCampaignBudgetItem           MediaBridgeSchemaNewAssociationResponseFromObjectType = "CAMPAIGN_BUDGET_ITEM"
	MediaBridgeSchemaNewAssociationResponseFromObjectTypeCampaignSpendItem            MediaBridgeSchemaNewAssociationResponseFromObjectType = "CAMPAIGN_SPEND_ITEM"
	MediaBridgeSchemaNewAssociationResponseFromObjectTypeMic                          MediaBridgeSchemaNewAssociationResponseFromObjectType = "MIC"
	MediaBridgeSchemaNewAssociationResponseFromObjectTypeContentAudit                 MediaBridgeSchemaNewAssociationResponseFromObjectType = "CONTENT_AUDIT"
	MediaBridgeSchemaNewAssociationResponseFromObjectTypeContentAuditPage             MediaBridgeSchemaNewAssociationResponseFromObjectType = "CONTENT_AUDIT_PAGE"
	MediaBridgeSchemaNewAssociationResponseFromObjectTypePlaylistFolder               MediaBridgeSchemaNewAssociationResponseFromObjectType = "PLAYLIST_FOLDER"
	MediaBridgeSchemaNewAssociationResponseFromObjectTypeLead                         MediaBridgeSchemaNewAssociationResponseFromObjectType = "LEAD"
	MediaBridgeSchemaNewAssociationResponseFromObjectTypeAbandonedCart                MediaBridgeSchemaNewAssociationResponseFromObjectType = "ABANDONED_CART"
	MediaBridgeSchemaNewAssociationResponseFromObjectTypeExternalWebURL               MediaBridgeSchemaNewAssociationResponseFromObjectType = "EXTERNAL_WEB_URL"
	MediaBridgeSchemaNewAssociationResponseFromObjectTypeView                         MediaBridgeSchemaNewAssociationResponseFromObjectType = "VIEW"
	MediaBridgeSchemaNewAssociationResponseFromObjectTypeViewBlock                    MediaBridgeSchemaNewAssociationResponseFromObjectType = "VIEW_BLOCK"
	MediaBridgeSchemaNewAssociationResponseFromObjectTypeRoster                       MediaBridgeSchemaNewAssociationResponseFromObjectType = "ROSTER"
	MediaBridgeSchemaNewAssociationResponseFromObjectTypeCart                         MediaBridgeSchemaNewAssociationResponseFromObjectType = "CART"
	MediaBridgeSchemaNewAssociationResponseFromObjectTypeAutomationPlatformFlowAction MediaBridgeSchemaNewAssociationResponseFromObjectType = "AUTOMATION_PLATFORM_FLOW_ACTION"
	MediaBridgeSchemaNewAssociationResponseFromObjectTypeSocialProfile                MediaBridgeSchemaNewAssociationResponseFromObjectType = "SOCIAL_PROFILE"
	MediaBridgeSchemaNewAssociationResponseFromObjectTypePartnerClient                MediaBridgeSchemaNewAssociationResponseFromObjectType = "PARTNER_CLIENT"
	MediaBridgeSchemaNewAssociationResponseFromObjectTypeRosterMember                 MediaBridgeSchemaNewAssociationResponseFromObjectType = "ROSTER_MEMBER"
	MediaBridgeSchemaNewAssociationResponseFromObjectTypeMarketingEventAttendance     MediaBridgeSchemaNewAssociationResponseFromObjectType = "MARKETING_EVENT_ATTENDANCE"
	MediaBridgeSchemaNewAssociationResponseFromObjectTypeAllPages                     MediaBridgeSchemaNewAssociationResponseFromObjectType = "ALL_PAGES"
	MediaBridgeSchemaNewAssociationResponseFromObjectTypeAIForecast                   MediaBridgeSchemaNewAssociationResponseFromObjectType = "AI_FORECAST"
	MediaBridgeSchemaNewAssociationResponseFromObjectTypeCRMPipelinesDummyType        MediaBridgeSchemaNewAssociationResponseFromObjectType = "CRM_PIPELINES_DUMMY_TYPE"
	MediaBridgeSchemaNewAssociationResponseFromObjectTypeKnowledgeArticle             MediaBridgeSchemaNewAssociationResponseFromObjectType = "KNOWLEDGE_ARTICLE"
	MediaBridgeSchemaNewAssociationResponseFromObjectTypePropertyInfo                 MediaBridgeSchemaNewAssociationResponseFromObjectType = "PROPERTY_INFO"
	MediaBridgeSchemaNewAssociationResponseFromObjectTypeDataPrivacyConsent           MediaBridgeSchemaNewAssociationResponseFromObjectType = "DATA_PRIVACY_CONSENT"
	MediaBridgeSchemaNewAssociationResponseFromObjectTypeGoalTemplate                 MediaBridgeSchemaNewAssociationResponseFromObjectType = "GOAL_TEMPLATE"
	MediaBridgeSchemaNewAssociationResponseFromObjectTypeScoreConfiguration           MediaBridgeSchemaNewAssociationResponseFromObjectType = "SCORE_CONFIGURATION"
	MediaBridgeSchemaNewAssociationResponseFromObjectTypeAudience                     MediaBridgeSchemaNewAssociationResponseFromObjectType = "AUDIENCE"
	MediaBridgeSchemaNewAssociationResponseFromObjectTypePartnerClientRevenue         MediaBridgeSchemaNewAssociationResponseFromObjectType = "PARTNER_CLIENT_REVENUE"
	MediaBridgeSchemaNewAssociationResponseFromObjectTypeAutomationJourney            MediaBridgeSchemaNewAssociationResponseFromObjectType = "AUTOMATION_JOURNEY"
	MediaBridgeSchemaNewAssociationResponseFromObjectTypeComboEventConfiguration      MediaBridgeSchemaNewAssociationResponseFromObjectType = "COMBO_EVENT_CONFIGURATION"
	MediaBridgeSchemaNewAssociationResponseFromObjectTypeCRMObjectsDummyType          MediaBridgeSchemaNewAssociationResponseFromObjectType = "CRM_OBJECTS_DUMMY_TYPE"
	MediaBridgeSchemaNewAssociationResponseFromObjectTypeCaseStudy                    MediaBridgeSchemaNewAssociationResponseFromObjectType = "CASE_STUDY"
	MediaBridgeSchemaNewAssociationResponseFromObjectTypeService                      MediaBridgeSchemaNewAssociationResponseFromObjectType = "SERVICE"
	MediaBridgeSchemaNewAssociationResponseFromObjectTypePodcastEpisode               MediaBridgeSchemaNewAssociationResponseFromObjectType = "PODCAST_EPISODE"
	MediaBridgeSchemaNewAssociationResponseFromObjectTypePartnerService               MediaBridgeSchemaNewAssociationResponseFromObjectType = "PARTNER_SERVICE"
	MediaBridgeSchemaNewAssociationResponseFromObjectTypeUnknown                      MediaBridgeSchemaNewAssociationResponseFromObjectType = "UNKNOWN"
)

// The name of the destination object type (e.g,. "DEAL" or "QUOTE").
type MediaBridgeSchemaNewAssociationResponseToObjectType string

const (
	MediaBridgeSchemaNewAssociationResponseToObjectTypeContact                      MediaBridgeSchemaNewAssociationResponseToObjectType = "CONTACT"
	MediaBridgeSchemaNewAssociationResponseToObjectTypeCompany                      MediaBridgeSchemaNewAssociationResponseToObjectType = "COMPANY"
	MediaBridgeSchemaNewAssociationResponseToObjectTypeDeal                         MediaBridgeSchemaNewAssociationResponseToObjectType = "DEAL"
	MediaBridgeSchemaNewAssociationResponseToObjectTypeEngagement                   MediaBridgeSchemaNewAssociationResponseToObjectType = "ENGAGEMENT"
	MediaBridgeSchemaNewAssociationResponseToObjectTypeTicket                       MediaBridgeSchemaNewAssociationResponseToObjectType = "TICKET"
	MediaBridgeSchemaNewAssociationResponseToObjectTypeOwner                        MediaBridgeSchemaNewAssociationResponseToObjectType = "OWNER"
	MediaBridgeSchemaNewAssociationResponseToObjectTypeProduct                      MediaBridgeSchemaNewAssociationResponseToObjectType = "PRODUCT"
	MediaBridgeSchemaNewAssociationResponseToObjectTypeLineItem                     MediaBridgeSchemaNewAssociationResponseToObjectType = "LINE_ITEM"
	MediaBridgeSchemaNewAssociationResponseToObjectTypeBetDeliverableService        MediaBridgeSchemaNewAssociationResponseToObjectType = "BET_DELIVERABLE_SERVICE"
	MediaBridgeSchemaNewAssociationResponseToObjectTypeContent                      MediaBridgeSchemaNewAssociationResponseToObjectType = "CONTENT"
	MediaBridgeSchemaNewAssociationResponseToObjectTypeConversation                 MediaBridgeSchemaNewAssociationResponseToObjectType = "CONVERSATION"
	MediaBridgeSchemaNewAssociationResponseToObjectTypeBetAlert                     MediaBridgeSchemaNewAssociationResponseToObjectType = "BET_ALERT"
	MediaBridgeSchemaNewAssociationResponseToObjectTypePortal                       MediaBridgeSchemaNewAssociationResponseToObjectType = "PORTAL"
	MediaBridgeSchemaNewAssociationResponseToObjectTypeQuote                        MediaBridgeSchemaNewAssociationResponseToObjectType = "QUOTE"
	MediaBridgeSchemaNewAssociationResponseToObjectTypeFormSubmissionInbounddb      MediaBridgeSchemaNewAssociationResponseToObjectType = "FORM_SUBMISSION_INBOUNDDB"
	MediaBridgeSchemaNewAssociationResponseToObjectTypeQuota                        MediaBridgeSchemaNewAssociationResponseToObjectType = "QUOTA"
	MediaBridgeSchemaNewAssociationResponseToObjectTypeUnsubscribe                  MediaBridgeSchemaNewAssociationResponseToObjectType = "UNSUBSCRIBE"
	MediaBridgeSchemaNewAssociationResponseToObjectTypeCommunication                MediaBridgeSchemaNewAssociationResponseToObjectType = "COMMUNICATION"
	MediaBridgeSchemaNewAssociationResponseToObjectTypeFeedbackSubmission           MediaBridgeSchemaNewAssociationResponseToObjectType = "FEEDBACK_SUBMISSION"
	MediaBridgeSchemaNewAssociationResponseToObjectTypeAttribution                  MediaBridgeSchemaNewAssociationResponseToObjectType = "ATTRIBUTION"
	MediaBridgeSchemaNewAssociationResponseToObjectTypeSalesforceSyncError          MediaBridgeSchemaNewAssociationResponseToObjectType = "SALESFORCE_SYNC_ERROR"
	MediaBridgeSchemaNewAssociationResponseToObjectTypeRestorableCRMObject          MediaBridgeSchemaNewAssociationResponseToObjectType = "RESTORABLE_CRM_OBJECT"
	MediaBridgeSchemaNewAssociationResponseToObjectTypeHub                          MediaBridgeSchemaNewAssociationResponseToObjectType = "HUB"
	MediaBridgeSchemaNewAssociationResponseToObjectTypeLandingPage                  MediaBridgeSchemaNewAssociationResponseToObjectType = "LANDING_PAGE"
	MediaBridgeSchemaNewAssociationResponseToObjectTypeProductOrFolder              MediaBridgeSchemaNewAssociationResponseToObjectType = "PRODUCT_OR_FOLDER"
	MediaBridgeSchemaNewAssociationResponseToObjectTypeTask                         MediaBridgeSchemaNewAssociationResponseToObjectType = "TASK"
	MediaBridgeSchemaNewAssociationResponseToObjectTypeForm                         MediaBridgeSchemaNewAssociationResponseToObjectType = "FORM"
	MediaBridgeSchemaNewAssociationResponseToObjectTypeMarketingEmail               MediaBridgeSchemaNewAssociationResponseToObjectType = "MARKETING_EMAIL"
	MediaBridgeSchemaNewAssociationResponseToObjectTypeAdAccount                    MediaBridgeSchemaNewAssociationResponseToObjectType = "AD_ACCOUNT"
	MediaBridgeSchemaNewAssociationResponseToObjectTypeAdCampaign                   MediaBridgeSchemaNewAssociationResponseToObjectType = "AD_CAMPAIGN"
	MediaBridgeSchemaNewAssociationResponseToObjectTypeAdGroup                      MediaBridgeSchemaNewAssociationResponseToObjectType = "AD_GROUP"
	MediaBridgeSchemaNewAssociationResponseToObjectTypeAd                           MediaBridgeSchemaNewAssociationResponseToObjectType = "AD"
	MediaBridgeSchemaNewAssociationResponseToObjectTypeKeyword                      MediaBridgeSchemaNewAssociationResponseToObjectType = "KEYWORD"
	MediaBridgeSchemaNewAssociationResponseToObjectTypeCampaign                     MediaBridgeSchemaNewAssociationResponseToObjectType = "CAMPAIGN"
	MediaBridgeSchemaNewAssociationResponseToObjectTypeSocialChannel                MediaBridgeSchemaNewAssociationResponseToObjectType = "SOCIAL_CHANNEL"
	MediaBridgeSchemaNewAssociationResponseToObjectTypeSocialPost                   MediaBridgeSchemaNewAssociationResponseToObjectType = "SOCIAL_POST"
	MediaBridgeSchemaNewAssociationResponseToObjectTypeSitePage                     MediaBridgeSchemaNewAssociationResponseToObjectType = "SITE_PAGE"
	MediaBridgeSchemaNewAssociationResponseToObjectTypeBlogPost                     MediaBridgeSchemaNewAssociationResponseToObjectType = "BLOG_POST"
	MediaBridgeSchemaNewAssociationResponseToObjectTypeImport                       MediaBridgeSchemaNewAssociationResponseToObjectType = "IMPORT"
	MediaBridgeSchemaNewAssociationResponseToObjectTypeExport                       MediaBridgeSchemaNewAssociationResponseToObjectType = "EXPORT"
	MediaBridgeSchemaNewAssociationResponseToObjectTypeCta                          MediaBridgeSchemaNewAssociationResponseToObjectType = "CTA"
	MediaBridgeSchemaNewAssociationResponseToObjectTypeTaskTemplate                 MediaBridgeSchemaNewAssociationResponseToObjectType = "TASK_TEMPLATE"
	MediaBridgeSchemaNewAssociationResponseToObjectTypeAutomationPlatformFlow       MediaBridgeSchemaNewAssociationResponseToObjectType = "AUTOMATION_PLATFORM_FLOW"
	MediaBridgeSchemaNewAssociationResponseToObjectTypeObjectList                   MediaBridgeSchemaNewAssociationResponseToObjectType = "OBJECT_LIST"
	MediaBridgeSchemaNewAssociationResponseToObjectTypeNote                         MediaBridgeSchemaNewAssociationResponseToObjectType = "NOTE"
	MediaBridgeSchemaNewAssociationResponseToObjectTypeMeetingEvent                 MediaBridgeSchemaNewAssociationResponseToObjectType = "MEETING_EVENT"
	MediaBridgeSchemaNewAssociationResponseToObjectTypeCall                         MediaBridgeSchemaNewAssociationResponseToObjectType = "CALL"
	MediaBridgeSchemaNewAssociationResponseToObjectTypeEmail                        MediaBridgeSchemaNewAssociationResponseToObjectType = "EMAIL"
	MediaBridgeSchemaNewAssociationResponseToObjectTypePublishingTask               MediaBridgeSchemaNewAssociationResponseToObjectType = "PUBLISHING_TASK"
	MediaBridgeSchemaNewAssociationResponseToObjectTypeConversationSession          MediaBridgeSchemaNewAssociationResponseToObjectType = "CONVERSATION_SESSION"
	MediaBridgeSchemaNewAssociationResponseToObjectTypeContactCreateAttribution     MediaBridgeSchemaNewAssociationResponseToObjectType = "CONTACT_CREATE_ATTRIBUTION"
	MediaBridgeSchemaNewAssociationResponseToObjectTypeInvoice                      MediaBridgeSchemaNewAssociationResponseToObjectType = "INVOICE"
	MediaBridgeSchemaNewAssociationResponseToObjectTypeMarketingEvent               MediaBridgeSchemaNewAssociationResponseToObjectType = "MARKETING_EVENT"
	MediaBridgeSchemaNewAssociationResponseToObjectTypeConversationInbox            MediaBridgeSchemaNewAssociationResponseToObjectType = "CONVERSATION_INBOX"
	MediaBridgeSchemaNewAssociationResponseToObjectTypeChatflow                     MediaBridgeSchemaNewAssociationResponseToObjectType = "CHATFLOW"
	MediaBridgeSchemaNewAssociationResponseToObjectTypeMediaBridge                  MediaBridgeSchemaNewAssociationResponseToObjectType = "MEDIA_BRIDGE"
	MediaBridgeSchemaNewAssociationResponseToObjectTypeSequence                     MediaBridgeSchemaNewAssociationResponseToObjectType = "SEQUENCE"
	MediaBridgeSchemaNewAssociationResponseToObjectTypeSequenceStep                 MediaBridgeSchemaNewAssociationResponseToObjectType = "SEQUENCE_STEP"
	MediaBridgeSchemaNewAssociationResponseToObjectTypeForecast                     MediaBridgeSchemaNewAssociationResponseToObjectType = "FORECAST"
	MediaBridgeSchemaNewAssociationResponseToObjectTypeSnippet                      MediaBridgeSchemaNewAssociationResponseToObjectType = "SNIPPET"
	MediaBridgeSchemaNewAssociationResponseToObjectTypeTemplate                     MediaBridgeSchemaNewAssociationResponseToObjectType = "TEMPLATE"
	MediaBridgeSchemaNewAssociationResponseToObjectTypeDealCreateAttribution        MediaBridgeSchemaNewAssociationResponseToObjectType = "DEAL_CREATE_ATTRIBUTION"
	MediaBridgeSchemaNewAssociationResponseToObjectTypeQuoteTemplate                MediaBridgeSchemaNewAssociationResponseToObjectType = "QUOTE_TEMPLATE"
	MediaBridgeSchemaNewAssociationResponseToObjectTypeQuoteModule                  MediaBridgeSchemaNewAssociationResponseToObjectType = "QUOTE_MODULE"
	MediaBridgeSchemaNewAssociationResponseToObjectTypeQuoteModuleField             MediaBridgeSchemaNewAssociationResponseToObjectType = "QUOTE_MODULE_FIELD"
	MediaBridgeSchemaNewAssociationResponseToObjectTypeQuoteField                   MediaBridgeSchemaNewAssociationResponseToObjectType = "QUOTE_FIELD"
	MediaBridgeSchemaNewAssociationResponseToObjectTypeSequenceEnrollment           MediaBridgeSchemaNewAssociationResponseToObjectType = "SEQUENCE_ENROLLMENT"
	MediaBridgeSchemaNewAssociationResponseToObjectTypeSubscription                 MediaBridgeSchemaNewAssociationResponseToObjectType = "SUBSCRIPTION"
	MediaBridgeSchemaNewAssociationResponseToObjectTypeAcceptanceTest               MediaBridgeSchemaNewAssociationResponseToObjectType = "ACCEPTANCE_TEST"
	MediaBridgeSchemaNewAssociationResponseToObjectTypeSocialBroadcast              MediaBridgeSchemaNewAssociationResponseToObjectType = "SOCIAL_BROADCAST"
	MediaBridgeSchemaNewAssociationResponseToObjectTypeDealSplit                    MediaBridgeSchemaNewAssociationResponseToObjectType = "DEAL_SPLIT"
	MediaBridgeSchemaNewAssociationResponseToObjectTypeDealRegistration             MediaBridgeSchemaNewAssociationResponseToObjectType = "DEAL_REGISTRATION"
	MediaBridgeSchemaNewAssociationResponseToObjectTypeGoalTarget                   MediaBridgeSchemaNewAssociationResponseToObjectType = "GOAL_TARGET"
	MediaBridgeSchemaNewAssociationResponseToObjectTypeGoalTargetGroup              MediaBridgeSchemaNewAssociationResponseToObjectType = "GOAL_TARGET_GROUP"
	MediaBridgeSchemaNewAssociationResponseToObjectTypePortalObjectSyncMessage      MediaBridgeSchemaNewAssociationResponseToObjectType = "PORTAL_OBJECT_SYNC_MESSAGE"
	MediaBridgeSchemaNewAssociationResponseToObjectTypeFileManagerFile              MediaBridgeSchemaNewAssociationResponseToObjectType = "FILE_MANAGER_FILE"
	MediaBridgeSchemaNewAssociationResponseToObjectTypeFileManagerFolder            MediaBridgeSchemaNewAssociationResponseToObjectType = "FILE_MANAGER_FOLDER"
	MediaBridgeSchemaNewAssociationResponseToObjectTypeSequenceStepEnrollment       MediaBridgeSchemaNewAssociationResponseToObjectType = "SEQUENCE_STEP_ENROLLMENT"
	MediaBridgeSchemaNewAssociationResponseToObjectTypeApproval                     MediaBridgeSchemaNewAssociationResponseToObjectType = "APPROVAL"
	MediaBridgeSchemaNewAssociationResponseToObjectTypeApprovalStep                 MediaBridgeSchemaNewAssociationResponseToObjectType = "APPROVAL_STEP"
	MediaBridgeSchemaNewAssociationResponseToObjectTypeCtaVariant                   MediaBridgeSchemaNewAssociationResponseToObjectType = "CTA_VARIANT"
	MediaBridgeSchemaNewAssociationResponseToObjectTypeSalesDocument                MediaBridgeSchemaNewAssociationResponseToObjectType = "SALES_DOCUMENT"
	MediaBridgeSchemaNewAssociationResponseToObjectTypeDiscount                     MediaBridgeSchemaNewAssociationResponseToObjectType = "DISCOUNT"
	MediaBridgeSchemaNewAssociationResponseToObjectTypeFee                          MediaBridgeSchemaNewAssociationResponseToObjectType = "FEE"
	MediaBridgeSchemaNewAssociationResponseToObjectTypeTax                          MediaBridgeSchemaNewAssociationResponseToObjectType = "TAX"
	MediaBridgeSchemaNewAssociationResponseToObjectTypeMarketingCalendar            MediaBridgeSchemaNewAssociationResponseToObjectType = "MARKETING_CALENDAR"
	MediaBridgeSchemaNewAssociationResponseToObjectTypePermissionsTesting           MediaBridgeSchemaNewAssociationResponseToObjectType = "PERMISSIONS_TESTING"
	MediaBridgeSchemaNewAssociationResponseToObjectTypePrivacyScannerCookie         MediaBridgeSchemaNewAssociationResponseToObjectType = "PRIVACY_SCANNER_COOKIE"
	MediaBridgeSchemaNewAssociationResponseToObjectTypeDataSyncState                MediaBridgeSchemaNewAssociationResponseToObjectType = "DATA_SYNC_STATE"
	MediaBridgeSchemaNewAssociationResponseToObjectTypeWebInteractive               MediaBridgeSchemaNewAssociationResponseToObjectType = "WEB_INTERACTIVE"
	MediaBridgeSchemaNewAssociationResponseToObjectTypePlaybook                     MediaBridgeSchemaNewAssociationResponseToObjectType = "PLAYBOOK"
	MediaBridgeSchemaNewAssociationResponseToObjectTypeFolder                       MediaBridgeSchemaNewAssociationResponseToObjectType = "FOLDER"
	MediaBridgeSchemaNewAssociationResponseToObjectTypePlaybookQuestion             MediaBridgeSchemaNewAssociationResponseToObjectType = "PLAYBOOK_QUESTION"
	MediaBridgeSchemaNewAssociationResponseToObjectTypePlaybookSubmission           MediaBridgeSchemaNewAssociationResponseToObjectType = "PLAYBOOK_SUBMISSION"
	MediaBridgeSchemaNewAssociationResponseToObjectTypePlaybookSubmissionAnswer     MediaBridgeSchemaNewAssociationResponseToObjectType = "PLAYBOOK_SUBMISSION_ANSWER"
	MediaBridgeSchemaNewAssociationResponseToObjectTypeCommercePayment              MediaBridgeSchemaNewAssociationResponseToObjectType = "COMMERCE_PAYMENT"
	MediaBridgeSchemaNewAssociationResponseToObjectTypeGscProperty                  MediaBridgeSchemaNewAssociationResponseToObjectType = "GSC_PROPERTY"
	MediaBridgeSchemaNewAssociationResponseToObjectTypeSoxProtectedDummyType        MediaBridgeSchemaNewAssociationResponseToObjectType = "SOX_PROTECTED_DUMMY_TYPE"
	MediaBridgeSchemaNewAssociationResponseToObjectTypeBlogListingPage              MediaBridgeSchemaNewAssociationResponseToObjectType = "BLOG_LISTING_PAGE"
	MediaBridgeSchemaNewAssociationResponseToObjectTypeQuarantinedSubmission        MediaBridgeSchemaNewAssociationResponseToObjectType = "QUARANTINED_SUBMISSION"
	MediaBridgeSchemaNewAssociationResponseToObjectTypePaymentSchedule              MediaBridgeSchemaNewAssociationResponseToObjectType = "PAYMENT_SCHEDULE"
	MediaBridgeSchemaNewAssociationResponseToObjectTypePaymentScheduleInstallment   MediaBridgeSchemaNewAssociationResponseToObjectType = "PAYMENT_SCHEDULE_INSTALLMENT"
	MediaBridgeSchemaNewAssociationResponseToObjectTypeMarketingCampaignUtm         MediaBridgeSchemaNewAssociationResponseToObjectType = "MARKETING_CAMPAIGN_UTM"
	MediaBridgeSchemaNewAssociationResponseToObjectTypeDiscountTemplate             MediaBridgeSchemaNewAssociationResponseToObjectType = "DISCOUNT_TEMPLATE"
	MediaBridgeSchemaNewAssociationResponseToObjectTypeDiscountCode                 MediaBridgeSchemaNewAssociationResponseToObjectType = "DISCOUNT_CODE"
	MediaBridgeSchemaNewAssociationResponseToObjectTypeFeedbackSurvey               MediaBridgeSchemaNewAssociationResponseToObjectType = "FEEDBACK_SURVEY"
	MediaBridgeSchemaNewAssociationResponseToObjectTypeCmsURL                       MediaBridgeSchemaNewAssociationResponseToObjectType = "CMS_URL"
	MediaBridgeSchemaNewAssociationResponseToObjectTypeSalesTask                    MediaBridgeSchemaNewAssociationResponseToObjectType = "SALES_TASK"
	MediaBridgeSchemaNewAssociationResponseToObjectTypeSalesWorkload                MediaBridgeSchemaNewAssociationResponseToObjectType = "SALES_WORKLOAD"
	MediaBridgeSchemaNewAssociationResponseToObjectTypeUser                         MediaBridgeSchemaNewAssociationResponseToObjectType = "USER"
	MediaBridgeSchemaNewAssociationResponseToObjectTypePostalMail                   MediaBridgeSchemaNewAssociationResponseToObjectType = "POSTAL_MAIL"
	MediaBridgeSchemaNewAssociationResponseToObjectTypeSchemasBackendTest           MediaBridgeSchemaNewAssociationResponseToObjectType = "SCHEMAS_BACKEND_TEST"
	MediaBridgeSchemaNewAssociationResponseToObjectTypePaymentLink                  MediaBridgeSchemaNewAssociationResponseToObjectType = "PAYMENT_LINK"
	MediaBridgeSchemaNewAssociationResponseToObjectTypeSubmissionTag                MediaBridgeSchemaNewAssociationResponseToObjectType = "SUBMISSION_TAG"
	MediaBridgeSchemaNewAssociationResponseToObjectTypeCampaignStep                 MediaBridgeSchemaNewAssociationResponseToObjectType = "CAMPAIGN_STEP"
	MediaBridgeSchemaNewAssociationResponseToObjectTypeSchedulingPage               MediaBridgeSchemaNewAssociationResponseToObjectType = "SCHEDULING_PAGE"
	MediaBridgeSchemaNewAssociationResponseToObjectTypeSoxProtectedTestType         MediaBridgeSchemaNewAssociationResponseToObjectType = "SOX_PROTECTED_TEST_TYPE"
	MediaBridgeSchemaNewAssociationResponseToObjectTypeOrder                        MediaBridgeSchemaNewAssociationResponseToObjectType = "ORDER"
	MediaBridgeSchemaNewAssociationResponseToObjectTypeMarketingSMS                 MediaBridgeSchemaNewAssociationResponseToObjectType = "MARKETING_SMS"
	MediaBridgeSchemaNewAssociationResponseToObjectTypePartnerAccount               MediaBridgeSchemaNewAssociationResponseToObjectType = "PARTNER_ACCOUNT"
	MediaBridgeSchemaNewAssociationResponseToObjectTypeCampaignTemplate             MediaBridgeSchemaNewAssociationResponseToObjectType = "CAMPAIGN_TEMPLATE"
	MediaBridgeSchemaNewAssociationResponseToObjectTypeCampaignTemplateStep         MediaBridgeSchemaNewAssociationResponseToObjectType = "CAMPAIGN_TEMPLATE_STEP"
	MediaBridgeSchemaNewAssociationResponseToObjectTypePlaylist                     MediaBridgeSchemaNewAssociationResponseToObjectType = "PLAYLIST"
	MediaBridgeSchemaNewAssociationResponseToObjectTypeClip                         MediaBridgeSchemaNewAssociationResponseToObjectType = "CLIP"
	MediaBridgeSchemaNewAssociationResponseToObjectTypeCampaignBudgetItem           MediaBridgeSchemaNewAssociationResponseToObjectType = "CAMPAIGN_BUDGET_ITEM"
	MediaBridgeSchemaNewAssociationResponseToObjectTypeCampaignSpendItem            MediaBridgeSchemaNewAssociationResponseToObjectType = "CAMPAIGN_SPEND_ITEM"
	MediaBridgeSchemaNewAssociationResponseToObjectTypeMic                          MediaBridgeSchemaNewAssociationResponseToObjectType = "MIC"
	MediaBridgeSchemaNewAssociationResponseToObjectTypeContentAudit                 MediaBridgeSchemaNewAssociationResponseToObjectType = "CONTENT_AUDIT"
	MediaBridgeSchemaNewAssociationResponseToObjectTypeContentAuditPage             MediaBridgeSchemaNewAssociationResponseToObjectType = "CONTENT_AUDIT_PAGE"
	MediaBridgeSchemaNewAssociationResponseToObjectTypePlaylistFolder               MediaBridgeSchemaNewAssociationResponseToObjectType = "PLAYLIST_FOLDER"
	MediaBridgeSchemaNewAssociationResponseToObjectTypeLead                         MediaBridgeSchemaNewAssociationResponseToObjectType = "LEAD"
	MediaBridgeSchemaNewAssociationResponseToObjectTypeAbandonedCart                MediaBridgeSchemaNewAssociationResponseToObjectType = "ABANDONED_CART"
	MediaBridgeSchemaNewAssociationResponseToObjectTypeExternalWebURL               MediaBridgeSchemaNewAssociationResponseToObjectType = "EXTERNAL_WEB_URL"
	MediaBridgeSchemaNewAssociationResponseToObjectTypeView                         MediaBridgeSchemaNewAssociationResponseToObjectType = "VIEW"
	MediaBridgeSchemaNewAssociationResponseToObjectTypeViewBlock                    MediaBridgeSchemaNewAssociationResponseToObjectType = "VIEW_BLOCK"
	MediaBridgeSchemaNewAssociationResponseToObjectTypeRoster                       MediaBridgeSchemaNewAssociationResponseToObjectType = "ROSTER"
	MediaBridgeSchemaNewAssociationResponseToObjectTypeCart                         MediaBridgeSchemaNewAssociationResponseToObjectType = "CART"
	MediaBridgeSchemaNewAssociationResponseToObjectTypeAutomationPlatformFlowAction MediaBridgeSchemaNewAssociationResponseToObjectType = "AUTOMATION_PLATFORM_FLOW_ACTION"
	MediaBridgeSchemaNewAssociationResponseToObjectTypeSocialProfile                MediaBridgeSchemaNewAssociationResponseToObjectType = "SOCIAL_PROFILE"
	MediaBridgeSchemaNewAssociationResponseToObjectTypePartnerClient                MediaBridgeSchemaNewAssociationResponseToObjectType = "PARTNER_CLIENT"
	MediaBridgeSchemaNewAssociationResponseToObjectTypeRosterMember                 MediaBridgeSchemaNewAssociationResponseToObjectType = "ROSTER_MEMBER"
	MediaBridgeSchemaNewAssociationResponseToObjectTypeMarketingEventAttendance     MediaBridgeSchemaNewAssociationResponseToObjectType = "MARKETING_EVENT_ATTENDANCE"
	MediaBridgeSchemaNewAssociationResponseToObjectTypeAllPages                     MediaBridgeSchemaNewAssociationResponseToObjectType = "ALL_PAGES"
	MediaBridgeSchemaNewAssociationResponseToObjectTypeAIForecast                   MediaBridgeSchemaNewAssociationResponseToObjectType = "AI_FORECAST"
	MediaBridgeSchemaNewAssociationResponseToObjectTypeCRMPipelinesDummyType        MediaBridgeSchemaNewAssociationResponseToObjectType = "CRM_PIPELINES_DUMMY_TYPE"
	MediaBridgeSchemaNewAssociationResponseToObjectTypeKnowledgeArticle             MediaBridgeSchemaNewAssociationResponseToObjectType = "KNOWLEDGE_ARTICLE"
	MediaBridgeSchemaNewAssociationResponseToObjectTypePropertyInfo                 MediaBridgeSchemaNewAssociationResponseToObjectType = "PROPERTY_INFO"
	MediaBridgeSchemaNewAssociationResponseToObjectTypeDataPrivacyConsent           MediaBridgeSchemaNewAssociationResponseToObjectType = "DATA_PRIVACY_CONSENT"
	MediaBridgeSchemaNewAssociationResponseToObjectTypeGoalTemplate                 MediaBridgeSchemaNewAssociationResponseToObjectType = "GOAL_TEMPLATE"
	MediaBridgeSchemaNewAssociationResponseToObjectTypeScoreConfiguration           MediaBridgeSchemaNewAssociationResponseToObjectType = "SCORE_CONFIGURATION"
	MediaBridgeSchemaNewAssociationResponseToObjectTypeAudience                     MediaBridgeSchemaNewAssociationResponseToObjectType = "AUDIENCE"
	MediaBridgeSchemaNewAssociationResponseToObjectTypePartnerClientRevenue         MediaBridgeSchemaNewAssociationResponseToObjectType = "PARTNER_CLIENT_REVENUE"
	MediaBridgeSchemaNewAssociationResponseToObjectTypeAutomationJourney            MediaBridgeSchemaNewAssociationResponseToObjectType = "AUTOMATION_JOURNEY"
	MediaBridgeSchemaNewAssociationResponseToObjectTypeComboEventConfiguration      MediaBridgeSchemaNewAssociationResponseToObjectType = "COMBO_EVENT_CONFIGURATION"
	MediaBridgeSchemaNewAssociationResponseToObjectTypeCRMObjectsDummyType          MediaBridgeSchemaNewAssociationResponseToObjectType = "CRM_OBJECTS_DUMMY_TYPE"
	MediaBridgeSchemaNewAssociationResponseToObjectTypeCaseStudy                    MediaBridgeSchemaNewAssociationResponseToObjectType = "CASE_STUDY"
	MediaBridgeSchemaNewAssociationResponseToObjectTypeService                      MediaBridgeSchemaNewAssociationResponseToObjectType = "SERVICE"
	MediaBridgeSchemaNewAssociationResponseToObjectTypePodcastEpisode               MediaBridgeSchemaNewAssociationResponseToObjectType = "PODCAST_EPISODE"
	MediaBridgeSchemaNewAssociationResponseToObjectTypePartnerService               MediaBridgeSchemaNewAssociationResponseToObjectType = "PARTNER_SERVICE"
	MediaBridgeSchemaNewAssociationResponseToObjectTypeUnknown                      MediaBridgeSchemaNewAssociationResponseToObjectType = "UNKNOWN"
)

type MediaBridgeSchemaUpdateParams struct {
	AppID string `path:"appId,required" json:"-"`
	// Defines attributes to update on an object type.
	ObjectTypeDefinitionPatch crm.ObjectTypeDefinitionPatchParam
	paramObj
}

func (r MediaBridgeSchemaUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.ObjectTypeDefinitionPatch)
}
func (r *MediaBridgeSchemaUpdateParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.ObjectTypeDefinitionPatch)
}

type MediaBridgeSchemaNewAssociationParams struct {
	AppID                    string `path:"appId,required" json:"-"`
	AssociationDefinitionEgg shared.AssociationDefinitionEggParam
	paramObj
}

func (r MediaBridgeSchemaNewAssociationParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.AssociationDefinitionEgg)
}
func (r *MediaBridgeSchemaNewAssociationParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.AssociationDefinitionEgg)
}

type MediaBridgeSchemaDeleteAssociationParams struct {
	AppID      string `path:"appId,required" json:"-"`
	ObjectType string `path:"objectType,required" json:"-"`
	paramObj
}

type MediaBridgeSchemaGetParams struct {
	AppID string `path:"appId,required" json:"-"`
	paramObj
}
