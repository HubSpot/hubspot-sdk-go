// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cms

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"slices"
	"time"

	"github.com/stainless-sdks/hubspot-sdk-go/internal/apijson"
	"github.com/stainless-sdks/hubspot-sdk-go/internal/requestconfig"
	"github.com/stainless-sdks/hubspot-sdk-go/option"
	"github.com/stainless-sdks/hubspot-sdk-go/packages/param"
	"github.com/stainless-sdks/hubspot-sdk-go/packages/respjson"
	"github.com/stainless-sdks/hubspot-sdk-go/shared"
)

// MediaBridgeIntegratorSettingService contains methods and other services that
// help with interacting with the Hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMediaBridgeIntegratorSettingService] method instead.
type MediaBridgeIntegratorSettingService struct {
	Options []option.RequestOption
}

// NewMediaBridgeIntegratorSettingService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewMediaBridgeIntegratorSettingService(opts ...option.RequestOption) (r MediaBridgeIntegratorSettingService) {
	r = MediaBridgeIntegratorSettingService{}
	r.Options = opts
	return
}

// Create a new media object type
func (r *MediaBridgeIntegratorSettingService) NewObjectDefinition(ctx context.Context, appID string, body MediaBridgeIntegratorSettingNewObjectDefinitionParams, opts ...option.RequestOption) (res *MediaBridgeIntegratorSettingNewObjectDefinitionResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if appID == "" {
		err = errors.New("missing required appId parameter")
		return
	}
	path := fmt.Sprintf("media-bridge/v1/%s/settings/object-definitions", appID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Set up a new oEmbed domain for your media bridge app.
func (r *MediaBridgeIntegratorSettingService) NewOembedDomain(ctx context.Context, appID string, body MediaBridgeIntegratorSettingNewOembedDomainParams, opts ...option.RequestOption) (res *MediaBridgeIntegratorSettingNewOembedDomainResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if appID == "" {
		err = errors.New("missing required appId parameter")
		return
	}
	path := fmt.Sprintf("media-bridge/v1/%s/settings/oembed-domains", appID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Delete an existing oEmbed domain.
func (r *MediaBridgeIntegratorSettingService) DeleteOembedDomain(ctx context.Context, appID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if appID == "" {
		err = errors.New("missing required appId parameter")
		return
	}
	path := fmt.Sprintf("media-bridge/v1/%s/settings/oembed-domains", appID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Get the visibility settings for media bridge events for your apps.
func (r *MediaBridgeIntegratorSettingService) GetEventVisibilitySettings(ctx context.Context, appID string, opts ...option.RequestOption) (res *MediaBridgeIntegratorSettingGetEventVisibilitySettingsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if appID == "" {
		err = errors.New("missing required appId parameter")
		return
	}
	path := fmt.Sprintf("media-bridge/v1/%s/settings/event-visibility", appID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Get the existing objects types that belong to the specified media type.
func (r *MediaBridgeIntegratorSettingService) GetObjectDefinitionsByMediaType(ctx context.Context, mediaType string, query MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeParams, opts ...option.RequestOption) (res *MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if query.AppID == "" {
		err = errors.New("missing required appId parameter")
		return
	}
	if mediaType == "" {
		err = errors.New("missing required mediaType parameter")
		return
	}
	path := fmt.Sprintf("media-bridge/v1/%s/settings/object-definitions/%s", query.AppID, mediaType)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Get the details for an existing oEmbed domain.
func (r *MediaBridgeIntegratorSettingService) GetOembedDomain(ctx context.Context, oEmbedDomainID string, query MediaBridgeIntegratorSettingGetOembedDomainParams, opts ...option.RequestOption) (res *MediaBridgeIntegratorSettingGetOembedDomainResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if query.AppID == "" {
		err = errors.New("missing required appId parameter")
		return
	}
	if oEmbedDomainID == "" {
		err = errors.New("missing required oEmbedDomainId parameter")
		return
	}
	path := fmt.Sprintf("media-bridge/v1/%s/settings/oembed-domains/%s", query.AppID, oEmbedDomainID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Get the details for existing oEmbed domains for your app
func (r *MediaBridgeIntegratorSettingService) ListOembedDomains(ctx context.Context, appID string, opts ...option.RequestOption) (res *MediaBridgeIntegratorSettingListOembedDomainsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if appID == "" {
		err = errors.New("missing required appId parameter")
		return
	}
	path := fmt.Sprintf("media-bridge/v1/%s/settings/oembed-domains", appID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Register the name that your app will display when a user is selecting media
// bridge items.
//
// Deprecated: deprecated
func (r *MediaBridgeIntegratorSettingService) RegisterAppName(ctx context.Context, appID string, body MediaBridgeIntegratorSettingRegisterAppNameParams, opts ...option.RequestOption) (res *MediaBridgeIntegratorSettingRegisterAppNameResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if appID == "" {
		err = errors.New("missing required appId parameter")
		return
	}
	path := fmt.Sprintf("media-bridge/v1/%s/settings/register", appID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Update the name that your app will display when a user is selecting media bridge
// items.
func (r *MediaBridgeIntegratorSettingService) UpdateAppName(ctx context.Context, appID string, body MediaBridgeIntegratorSettingUpdateAppNameParams, opts ...option.RequestOption) (res *MediaBridgeIntegratorSettingUpdateAppNameResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if appID == "" {
		err = errors.New("missing required appId parameter")
		return
	}
	path := fmt.Sprintf("media-bridge/v1/%s/settings", appID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Set the visibility settings for media bridge events created by your app.
func (r *MediaBridgeIntegratorSettingService) UpdateEventVisibilitySettings(ctx context.Context, appID string, body MediaBridgeIntegratorSettingUpdateEventVisibilitySettingsParams, opts ...option.RequestOption) (res *MediaBridgeIntegratorSettingUpdateEventVisibilitySettingsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if appID == "" {
		err = errors.New("missing required appId parameter")
		return
	}
	path := fmt.Sprintf("media-bridge/v1/%s/settings/event-visibility", appID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Update an existing oEmbed domain.
func (r *MediaBridgeIntegratorSettingService) UpdateOembedDomain(ctx context.Context, oEmbedDomainID string, params MediaBridgeIntegratorSettingUpdateOembedDomainParams, opts ...option.RequestOption) (res *MediaBridgeIntegratorSettingUpdateOembedDomainResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.AppID == "" {
		err = errors.New("missing required appId parameter")
		return
	}
	if oEmbedDomainID == "" {
		err = errors.New("missing required oEmbedDomainId parameter")
		return
	}
	path := fmt.Sprintf("media-bridge/v1/%s/settings/oembed-domains/%s", params.AppID, oEmbedDomainID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return
}

type MediaBridgeIntegratorSettingNewObjectDefinitionResponse struct {
	CreatedObjects map[string]MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObject `json:"createdObjects,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreatedObjects respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MediaBridgeIntegratorSettingNewObjectDefinitionResponse) RawJSON() string { return r.JSON.raw }
func (r *MediaBridgeIntegratorSettingNewObjectDefinitionResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObject struct {
	ObjectType     MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectObjectType      `json:"objectType,required"`
	Properties     []MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectProperty      `json:"properties,required"`
	PropertyGroups []MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyGroup `json:"propertyGroups,required"`
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
func (r MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObject) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectObjectType struct {
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
	// Any of "HUBSPOT", "INTEGRATION", "PORTAL_SPECIFIC", "CMS_HUBDB",
	// "HUBSPOT_EVENT", "INTEGRATION_EVENT", "PORTAL_SPECIFIC_EVENT".
	MetaType                           string                                                                                       `json:"metaType,required"`
	MetaTypeID                         int64                                                                                        `json:"metaTypeId,required"`
	Name                               string                                                                                       `json:"name,required"`
	ObjectTypeID                       string                                                                                       `json:"objectTypeId,required"`
	PermissioningType                  string                                                                                       `json:"permissioningType,required"`
	PipelinePropertyName               string                                                                                       `json:"pipelinePropertyName,required"`
	PipelineStagePropertyName          string                                                                                       `json:"pipelineStagePropertyName,required"`
	RequiredProperties                 []string                                                                                     `json:"requiredProperties,required"`
	Restorable                         bool                                                                                         `json:"restorable,required"`
	ScopeMappings                      []MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectObjectTypeScopeMapping `json:"scopeMappings,required"`
	SecondaryDisplayLabelPropertyNames []string                                                                                     `json:"secondaryDisplayLabelPropertyNames,required"`
	AccessScopeName                    string                                                                                       `json:"accessScopeName"`
	CreatedAt                          int64                                                                                        `json:"createdAt"`
	Description                        string                                                                                       `json:"description"`
	IntegrationAppID                   int64                                                                                        `json:"integrationAppId"`
	JanusGroup                         string                                                                                       `json:"janusGroup"`
	OwnerPortalID                      int64                                                                                        `json:"ownerPortalId"`
	PipelineCloseDatePropertyName      string                                                                                       `json:"pipelineCloseDatePropertyName"`
	PipelineTimeToClosePropertyName    string                                                                                       `json:"pipelineTimeToClosePropertyName"`
	PluralForm                         string                                                                                       `json:"pluralForm"`
	PrimaryDisplayLabelPropertyName    string                                                                                       `json:"primaryDisplayLabelPropertyName"`
	ReadScopeName                      string                                                                                       `json:"readScopeName"`
	SingularForm                       string                                                                                       `json:"singularForm"`
	Status                             string                                                                                       `json:"status"`
	Visibility                         string                                                                                       `json:"visibility"`
	WriteScopeName                     string                                                                                       `json:"writeScopeName"`
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
func (r MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectObjectType) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectObjectType) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectObjectTypeScopeMapping struct {
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
func (r MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectObjectTypeScopeMapping) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectObjectTypeScopeMapping) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectProperty struct {
	ObjectTypeID string `json:"objectTypeId,required"`
	// Defines a property
	Property                 shared.Property                                                                                        `json:"property,required"`
	CalculationExpression    MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionUnion `json:"calculationExpression"`
	CalculationFormula       string                                                                                                 `json:"calculationFormula"`
	DefinitionSource         MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyDefinitionSource           `json:"definitionSource"`
	ExtensionData            MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyExtensionData              `json:"extensionData"`
	ExternalOptionsMetaData  MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyExternalOptionsMetaData    `json:"externalOptionsMetaData"`
	FulcrumPortalID          int64                                                                                                  `json:"fulcrumPortalId"`
	FulcrumTimestamp         int64                                                                                                  `json:"fulcrumTimestamp"`
	JanusGroup               string                                                                                                 `json:"janusGroup"`
	Permission               MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyPermission                 `json:"permission"`
	PropertyDefinitionSource MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyPropertyDefinitionSource   `json:"propertyDefinitionSource"`
	PropertyRequirements     MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyPropertyRequirements       `json:"propertyRequirements"`
	RollupExpression         MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpression           `json:"rollupExpression"`
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
func (r MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectProperty) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectProperty) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionUnion
// contains all possible properties and values from
// [MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionConstantBoolean],
// [MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionConstantNumber],
// [MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionConstantString],
// [MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionBooleanPropertyVariable],
// [MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionStringPropertyVariable],
// [MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionNumberPropertyVariable],
// [MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionTimestampOfPropertyVariable],
// [MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionBooleanTargetPropertyVariable],
// [MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionStringTargetPropertyVariable],
// [MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionNumberTargetPropertyVariable],
// [MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionTimestampOfTargetPropertyVariable],
// [MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionAddNumbers],
// [MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionSubtractNumbers],
// [MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionMultiplyNumbers],
// [MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionDivideNumbers],
// [MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionRoundDown],
// [MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionRoundUp],
// [MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionRoundNearest],
// [MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionUpperCase],
// [MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionLowerCase],
// [MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionConcatStrings],
// [MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionContains],
// [MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionBeginsWith],
// [MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionNumberToString],
// [MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionParseNumber],
// [MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionFetchExchangeRate],
// [MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionFetchCurrencyDecimalPlaces],
// [MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionFetchSingleCurrencyPortalCurrency],
// [MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionDatedExchangeRate],
// [MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionPipelineProbability],
// [MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionMaxNumbers],
// [MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionMinNumbers],
// [MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionLessThan],
// [MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionLessThanOrEqual],
// [MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionMoreThan],
// [MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionMoreThanOrEqual],
// [MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionNumberEquals],
// [MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionStringEquals],
// [MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionIsPipelineStageClosed],
// [MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionNot],
// [MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionDate],
// [MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionMonth],
// [MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionYear],
// [MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionNow],
// [MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionTimeBetween],
// [MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionPeriodToMonths],
// [MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionPeriodToWeeks],
// [MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionAnd],
// [MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionOr],
// [MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionXor],
// [MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionIfString],
// [MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionIfNumber],
// [MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionIfBoolean],
// [MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionIsPresent],
// [MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionHasEmailReply],
// [MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionHasPlainTextEmailReply],
// [MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionExtractMostRecentEmailReplyHTML],
// [MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionExtractMostRecentEmailReplyText],
// [MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionExtractMostRecentPlainTextEmailReply],
// [MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionSetContainsString],
// [MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionIsEngagementType],
// [MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionFormatFullName],
// [MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionAbsoluteValue],
// [MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionSquareRoot],
// [MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionPower],
// [MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionSubstring],
// [MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionEuler],
// [MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionStringLength],
// [MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionAddTime],
// [MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionSubtractTime].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionUnion struct {
	Operator     string `json:"operator"`
	Inputs       []any  `json:"inputs"`
	PropertyName string `json:"propertyName"`
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
	Value                 MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionUnionValue `json:"value"`
	EnclosedInParentheses bool                                                                                                        `json:"enclosedInParentheses"`
	StringToCheck         any                                                                                                         `json:"stringToCheck"`
	IfExpression          any                                                                                                         `json:"ifExpression"`
	ElseExpression        any                                                                                                         `json:"elseExpression"`
	// This field is from variant
	// [MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionIsPresent].
	ExpressionToEvaluate any `json:"expressionToEvaluate"`
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

func (u MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionUnion) AsConstantBoolean() (v MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionConstantBoolean) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionUnion) AsConstantNumber() (v MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionConstantNumber) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionUnion) AsConstantString() (v MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionConstantString) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionUnion) AsBooleanPropertyVariable() (v MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionBooleanPropertyVariable) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionUnion) AsStringPropertyVariable() (v MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionStringPropertyVariable) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionUnion) AsNumberPropertyVariable() (v MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionNumberPropertyVariable) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionUnion) AsTimestampOfPropertyVariable() (v MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionTimestampOfPropertyVariable) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionUnion) AsBooleanTargetPropertyVariable() (v MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionBooleanTargetPropertyVariable) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionUnion) AsStringTargetPropertyVariable() (v MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionStringTargetPropertyVariable) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionUnion) AsNumberTargetPropertyVariable() (v MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionNumberTargetPropertyVariable) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionUnion) AsTimestampOfTargetPropertyVariable() (v MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionTimestampOfTargetPropertyVariable) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionUnion) AsAddNumbers() (v MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionAddNumbers) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionUnion) AsSubtractNumbers() (v MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionSubtractNumbers) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionUnion) AsMultiplyNumbers() (v MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionMultiplyNumbers) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionUnion) AsDivideNumbers() (v MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionDivideNumbers) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionUnion) AsRoundDown() (v MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionRoundDown) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionUnion) AsRoundUp() (v MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionRoundUp) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionUnion) AsRoundNearest() (v MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionRoundNearest) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionUnion) AsUpperCase() (v MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionUpperCase) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionUnion) AsLowerCase() (v MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionLowerCase) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionUnion) AsConcatStrings() (v MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionConcatStrings) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionUnion) AsContains() (v MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionContains) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionUnion) AsBeginsWith() (v MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionBeginsWith) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionUnion) AsNumberToString() (v MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionNumberToString) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionUnion) AsParseNumber() (v MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionParseNumber) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionUnion) AsFetchExchangeRate() (v MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionFetchExchangeRate) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionUnion) AsFetchCurrencyDecimalPlaces() (v MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionFetchCurrencyDecimalPlaces) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionUnion) AsFetchSingleCurrencyPortalCurrency() (v MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionFetchSingleCurrencyPortalCurrency) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionUnion) AsDatedExchangeRate() (v MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionDatedExchangeRate) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionUnion) AsPipelineProbability() (v MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionPipelineProbability) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionUnion) AsMaxNumbers() (v MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionMaxNumbers) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionUnion) AsMinNumbers() (v MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionMinNumbers) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionUnion) AsLessThan() (v MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionLessThan) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionUnion) AsLessThanOrEqual() (v MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionLessThanOrEqual) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionUnion) AsMoreThan() (v MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionMoreThan) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionUnion) AsMoreThanOrEqual() (v MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionMoreThanOrEqual) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionUnion) AsNumberEquals() (v MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionNumberEquals) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionUnion) AsStringEquals() (v MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionStringEquals) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionUnion) AsIsPipelineStageClosed() (v MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionIsPipelineStageClosed) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionUnion) AsNot() (v MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionNot) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionUnion) AsDate() (v MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionDate) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionUnion) AsMonth() (v MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionMonth) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionUnion) AsYear() (v MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionYear) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionUnion) AsNow() (v MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionNow) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionUnion) AsTimeBetween() (v MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionTimeBetween) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionUnion) AsPeriodToMonths() (v MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionPeriodToMonths) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionUnion) AsPeriodToWeeks() (v MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionPeriodToWeeks) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionUnion) AsAnd() (v MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionAnd) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionUnion) AsOr() (v MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionOr) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionUnion) AsXor() (v MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionXor) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionUnion) AsIfString() (v MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionIfString) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionUnion) AsIfNumber() (v MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionIfNumber) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionUnion) AsIfBoolean() (v MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionIfBoolean) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionUnion) AsIsPresent() (v MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionIsPresent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionUnion) AsHasEmailReply() (v MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionHasEmailReply) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionUnion) AsHasPlainTextEmailReply() (v MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionHasPlainTextEmailReply) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionUnion) AsExtractMostRecentEmailReplyHTML() (v MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionExtractMostRecentEmailReplyHTML) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionUnion) AsExtractMostRecentEmailReplyText() (v MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionExtractMostRecentEmailReplyText) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionUnion) AsExtractMostRecentPlainTextEmailReply() (v MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionExtractMostRecentPlainTextEmailReply) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionUnion) AsSetContainsString() (v MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionSetContainsString) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionUnion) AsIsEngagementType() (v MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionIsEngagementType) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionUnion) AsFormatFullName() (v MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionFormatFullName) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionUnion) AsAbsoluteValue() (v MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionAbsoluteValue) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionUnion) AsSquareRoot() (v MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionSquareRoot) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionUnion) AsPower() (v MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionPower) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionUnion) AsSubstring() (v MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionSubstring) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionUnion) AsEuler() (v MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionEuler) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionUnion) AsStringLength() (v MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionStringLength) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionUnion) AsAddTime() (v MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionAddTime) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionUnion) AsSubtractTime() (v MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionSubtractTime) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionUnionValue
// is an implicit subunion of
// [MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionUnion].
// MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionUnionValue
// provides convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfBool OfFloat OfString]
type MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionUnionValue struct {
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

func (r *MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionUnionValue) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionConstantBoolean struct {
	// Any of "CONSTANT_BOOLEAN".
	Operator     string `json:"operator,required"`
	Inputs       []any  `json:"inputs"`
	PropertyName string `json:"propertyName"`
	Value        bool   `json:"value"`
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
func (r MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionConstantBoolean) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionConstantBoolean) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionConstantNumber struct {
	// Any of "CONSTANT_NUMBER".
	Operator     string  `json:"operator,required"`
	Inputs       []any   `json:"inputs"`
	PropertyName string  `json:"propertyName"`
	Value        float64 `json:"value"`
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
func (r MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionConstantNumber) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionConstantNumber) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionConstantString struct {
	// Any of "CONSTANT_STRING".
	Operator     string `json:"operator,required"`
	Inputs       []any  `json:"inputs"`
	PropertyName string `json:"propertyName"`
	Value        string `json:"value"`
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
func (r MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionConstantString) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionConstantString) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionBooleanPropertyVariable struct {
	// Any of "BOOLEAN_PROPERTY_VARIABLE".
	Operator     string `json:"operator,required"`
	Inputs       []any  `json:"inputs"`
	PropertyName string `json:"propertyName"`
	Value        bool   `json:"value"`
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
func (r MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionBooleanPropertyVariable) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionBooleanPropertyVariable) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionStringPropertyVariable struct {
	// Any of "STRING_PROPERTY_VARIABLE".
	Operator     string `json:"operator,required"`
	Inputs       []any  `json:"inputs"`
	PropertyName string `json:"propertyName"`
	Value        string `json:"value"`
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
func (r MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionStringPropertyVariable) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionStringPropertyVariable) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionNumberPropertyVariable struct {
	// Any of "NUMBER_PROPERTY_VARIABLE".
	Operator     string  `json:"operator,required"`
	Inputs       []any   `json:"inputs"`
	PropertyName string  `json:"propertyName"`
	Value        float64 `json:"value"`
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
func (r MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionNumberPropertyVariable) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionNumberPropertyVariable) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionTimestampOfPropertyVariable struct {
	// Any of "TIMESTAMP_OF_PROPERTY_VARIABLE".
	Operator     string `json:"operator,required"`
	Inputs       []any  `json:"inputs"`
	PropertyName string `json:"propertyName"`
	Value        string `json:"value"`
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
func (r MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionTimestampOfPropertyVariable) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionTimestampOfPropertyVariable) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionBooleanTargetPropertyVariable struct {
	// Any of "BOOLEAN_TARGET_PROPERTY_VARIABLE".
	Operator     string `json:"operator,required"`
	Inputs       []any  `json:"inputs"`
	PropertyName string `json:"propertyName"`
	Value        bool   `json:"value"`
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
func (r MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionBooleanTargetPropertyVariable) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionBooleanTargetPropertyVariable) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionStringTargetPropertyVariable struct {
	// Any of "STRING_TARGET_PROPERTY_VARIABLE".
	Operator     string `json:"operator,required"`
	Inputs       []any  `json:"inputs"`
	PropertyName string `json:"propertyName"`
	Value        string `json:"value"`
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
func (r MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionStringTargetPropertyVariable) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionStringTargetPropertyVariable) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionNumberTargetPropertyVariable struct {
	// Any of "NUMBER_TARGET_PROPERTY_VARIABLE".
	Operator     string  `json:"operator,required"`
	Inputs       []any   `json:"inputs"`
	PropertyName string  `json:"propertyName"`
	Value        float64 `json:"value"`
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
func (r MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionNumberTargetPropertyVariable) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionNumberTargetPropertyVariable) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionTimestampOfTargetPropertyVariable struct {
	// Any of "TIMESTAMP_OF_TARGET_PROPERTY_VARIABLE".
	Operator     string `json:"operator,required"`
	Inputs       []any  `json:"inputs"`
	PropertyName string `json:"propertyName"`
	Value        string `json:"value"`
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
func (r MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionTimestampOfTargetPropertyVariable) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionTimestampOfTargetPropertyVariable) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionAddNumbers struct {
	EnclosedInParentheses bool `json:"enclosedInParentheses,required"`
	// Any of "ADD_NUMBERS".
	Operator     string  `json:"operator,required"`
	Inputs       []any   `json:"inputs"`
	PropertyName string  `json:"propertyName"`
	Value        float64 `json:"value"`
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
func (r MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionAddNumbers) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionAddNumbers) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionSubtractNumbers struct {
	EnclosedInParentheses bool `json:"enclosedInParentheses,required"`
	// Any of "SUBTRACT_NUMBERS".
	Operator     string  `json:"operator,required"`
	Inputs       []any   `json:"inputs"`
	PropertyName string  `json:"propertyName"`
	Value        float64 `json:"value"`
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
func (r MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionSubtractNumbers) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionSubtractNumbers) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionMultiplyNumbers struct {
	EnclosedInParentheses bool `json:"enclosedInParentheses,required"`
	// Any of "MULTIPLY_NUMBERS".
	Operator     string  `json:"operator,required"`
	Inputs       []any   `json:"inputs"`
	PropertyName string  `json:"propertyName"`
	Value        float64 `json:"value"`
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
func (r MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionMultiplyNumbers) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionMultiplyNumbers) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionDivideNumbers struct {
	EnclosedInParentheses bool `json:"enclosedInParentheses,required"`
	// Any of "DIVIDE_NUMBERS".
	Operator     string  `json:"operator,required"`
	Inputs       []any   `json:"inputs"`
	PropertyName string  `json:"propertyName"`
	Value        float64 `json:"value"`
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
func (r MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionDivideNumbers) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionDivideNumbers) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionRoundDown struct {
	// Any of "ROUND_DOWN".
	Operator     string  `json:"operator,required"`
	Inputs       []any   `json:"inputs"`
	PropertyName string  `json:"propertyName"`
	Value        float64 `json:"value"`
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
func (r MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionRoundDown) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionRoundDown) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionRoundUp struct {
	// Any of "ROUND_UP".
	Operator     string  `json:"operator,required"`
	Inputs       []any   `json:"inputs"`
	PropertyName string  `json:"propertyName"`
	Value        float64 `json:"value"`
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
func (r MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionRoundUp) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionRoundUp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionRoundNearest struct {
	// Any of "ROUND_NEAREST".
	Operator     string  `json:"operator,required"`
	Inputs       []any   `json:"inputs"`
	PropertyName string  `json:"propertyName"`
	Value        float64 `json:"value"`
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
func (r MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionRoundNearest) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionRoundNearest) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionUpperCase struct {
	// Any of "UPPER_CASE".
	Operator     string `json:"operator,required"`
	Inputs       []any  `json:"inputs"`
	PropertyName string `json:"propertyName"`
	Value        string `json:"value"`
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
func (r MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionUpperCase) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionUpperCase) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionLowerCase struct {
	// Any of "LOWER_CASE".
	Operator     string `json:"operator,required"`
	Inputs       []any  `json:"inputs"`
	PropertyName string `json:"propertyName"`
	Value        string `json:"value"`
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
func (r MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionLowerCase) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionLowerCase) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionConcatStrings struct {
	// Any of "CONCAT_STRINGS".
	Operator     string `json:"operator,required"`
	Inputs       []any  `json:"inputs"`
	PropertyName string `json:"propertyName"`
	Value        string `json:"value"`
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
func (r MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionConcatStrings) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionConcatStrings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionContains struct {
	// Any of "CONTAINS".
	Operator      string `json:"operator,required"`
	StringToCheck any    `json:"stringToCheck,required"`
	Inputs        []any  `json:"inputs"`
	PropertyName  string `json:"propertyName"`
	Value         bool   `json:"value"`
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
func (r MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionContains) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionContains) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionBeginsWith struct {
	// Any of "BEGINS_WITH".
	Operator      string `json:"operator,required"`
	StringToCheck any    `json:"stringToCheck,required"`
	Inputs        []any  `json:"inputs"`
	PropertyName  string `json:"propertyName"`
	Value         bool   `json:"value"`
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
func (r MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionBeginsWith) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionBeginsWith) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionNumberToString struct {
	// Any of "NUMBER_TO_STRING".
	Operator     string `json:"operator,required"`
	Inputs       []any  `json:"inputs"`
	PropertyName string `json:"propertyName"`
	Value        string `json:"value"`
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
func (r MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionNumberToString) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionNumberToString) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionParseNumber struct {
	// Any of "PARSE_NUMBER".
	Operator     string  `json:"operator,required"`
	Inputs       []any   `json:"inputs"`
	PropertyName string  `json:"propertyName"`
	Value        float64 `json:"value"`
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
func (r MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionParseNumber) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionParseNumber) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionFetchExchangeRate struct {
	// Any of "FETCH_EXCHANGE_RATE".
	Operator     string  `json:"operator,required"`
	Inputs       []any   `json:"inputs"`
	PropertyName string  `json:"propertyName"`
	Value        float64 `json:"value"`
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
func (r MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionFetchExchangeRate) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionFetchExchangeRate) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionFetchCurrencyDecimalPlaces struct {
	// Any of "FETCH_CURRENCY_DECIMAL_PLACES".
	Operator     string  `json:"operator,required"`
	Inputs       []any   `json:"inputs"`
	PropertyName string  `json:"propertyName"`
	Value        float64 `json:"value"`
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
func (r MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionFetchCurrencyDecimalPlaces) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionFetchCurrencyDecimalPlaces) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionFetchSingleCurrencyPortalCurrency struct {
	// Any of "FETCH_SINGLE_CURRENCY_PORTAL_CURRENCY".
	Operator     string `json:"operator,required"`
	Inputs       []any  `json:"inputs"`
	PropertyName string `json:"propertyName"`
	Value        string `json:"value"`
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
func (r MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionFetchSingleCurrencyPortalCurrency) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionFetchSingleCurrencyPortalCurrency) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionDatedExchangeRate struct {
	// Any of "DATED_EXCHANGE_RATE".
	Operator     string  `json:"operator,required"`
	Inputs       []any   `json:"inputs"`
	PropertyName string  `json:"propertyName"`
	Value        float64 `json:"value"`
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
func (r MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionDatedExchangeRate) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionDatedExchangeRate) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionPipelineProbability struct {
	// Any of "PIPELINE_PROBABILITY".
	Operator     string  `json:"operator,required"`
	Inputs       []any   `json:"inputs"`
	PropertyName string  `json:"propertyName"`
	Value        float64 `json:"value"`
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
func (r MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionPipelineProbability) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionPipelineProbability) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionMaxNumbers struct {
	// Any of "MAX_NUMBERS".
	Operator     string  `json:"operator,required"`
	Inputs       []any   `json:"inputs"`
	PropertyName string  `json:"propertyName"`
	Value        float64 `json:"value"`
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
func (r MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionMaxNumbers) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionMaxNumbers) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionMinNumbers struct {
	// Any of "MIN_NUMBERS".
	Operator     string  `json:"operator,required"`
	Inputs       []any   `json:"inputs"`
	PropertyName string  `json:"propertyName"`
	Value        float64 `json:"value"`
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
func (r MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionMinNumbers) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionMinNumbers) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionLessThan struct {
	// Any of "LESS_THAN".
	Operator     string `json:"operator,required"`
	Inputs       []any  `json:"inputs"`
	PropertyName string `json:"propertyName"`
	Value        bool   `json:"value"`
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
func (r MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionLessThan) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionLessThan) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionLessThanOrEqual struct {
	// Any of "LESS_THAN_OR_EQUAL".
	Operator     string `json:"operator,required"`
	Inputs       []any  `json:"inputs"`
	PropertyName string `json:"propertyName"`
	Value        bool   `json:"value"`
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
func (r MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionLessThanOrEqual) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionLessThanOrEqual) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionMoreThan struct {
	// Any of "MORE_THAN".
	Operator     string `json:"operator,required"`
	Inputs       []any  `json:"inputs"`
	PropertyName string `json:"propertyName"`
	Value        bool   `json:"value"`
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
func (r MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionMoreThan) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionMoreThan) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionMoreThanOrEqual struct {
	// Any of "MORE_THAN_OR_EQUAL".
	Operator     string `json:"operator,required"`
	Inputs       []any  `json:"inputs"`
	PropertyName string `json:"propertyName"`
	Value        bool   `json:"value"`
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
func (r MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionMoreThanOrEqual) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionMoreThanOrEqual) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionNumberEquals struct {
	// Any of "NUMBER_EQUALS".
	Operator     string `json:"operator,required"`
	Inputs       []any  `json:"inputs"`
	PropertyName string `json:"propertyName"`
	Value        bool   `json:"value"`
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
func (r MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionNumberEquals) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionNumberEquals) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionStringEquals struct {
	// Any of "STRING_EQUALS".
	Operator     string `json:"operator,required"`
	Inputs       []any  `json:"inputs"`
	PropertyName string `json:"propertyName"`
	Value        bool   `json:"value"`
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
func (r MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionStringEquals) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionStringEquals) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionIsPipelineStageClosed struct {
	// Any of "IS_PIPELINE_STAGE_CLOSED".
	Operator     string `json:"operator,required"`
	Inputs       []any  `json:"inputs"`
	PropertyName string `json:"propertyName"`
	Value        bool   `json:"value"`
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
func (r MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionIsPipelineStageClosed) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionIsPipelineStageClosed) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionNot struct {
	// Any of "NOT".
	Operator     string `json:"operator,required"`
	Inputs       []any  `json:"inputs"`
	PropertyName string `json:"propertyName"`
	Value        bool   `json:"value"`
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
func (r MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionNot) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionNot) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionDate struct {
	// Any of "DATE".
	Operator     string  `json:"operator,required"`
	Inputs       []any   `json:"inputs"`
	PropertyName string  `json:"propertyName"`
	Value        float64 `json:"value"`
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
func (r MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionDate) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionDate) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionMonth struct {
	// Any of "MONTH".
	Operator     string  `json:"operator,required"`
	Inputs       []any   `json:"inputs"`
	PropertyName string  `json:"propertyName"`
	Value        float64 `json:"value"`
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
func (r MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionMonth) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionMonth) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionYear struct {
	// Any of "YEAR".
	Operator     string  `json:"operator,required"`
	Inputs       []any   `json:"inputs"`
	PropertyName string  `json:"propertyName"`
	Value        float64 `json:"value"`
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
func (r MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionYear) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionYear) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionNow struct {
	// Any of "NOW".
	Operator     string  `json:"operator,required"`
	Inputs       []any   `json:"inputs"`
	PropertyName string  `json:"propertyName"`
	Value        float64 `json:"value"`
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
func (r MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionNow) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionNow) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionTimeBetween struct {
	// Any of "TIME_BETWEEN".
	Operator     string  `json:"operator,required"`
	Inputs       []any   `json:"inputs"`
	PropertyName string  `json:"propertyName"`
	Value        float64 `json:"value"`
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
func (r MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionTimeBetween) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionTimeBetween) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionPeriodToMonths struct {
	// Any of "PERIOD_TO_MONTHS".
	Operator     string  `json:"operator,required"`
	Inputs       []any   `json:"inputs"`
	PropertyName string  `json:"propertyName"`
	Value        float64 `json:"value"`
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
func (r MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionPeriodToMonths) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionPeriodToMonths) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionPeriodToWeeks struct {
	// Any of "PERIOD_TO_WEEKS".
	Operator     string  `json:"operator,required"`
	Inputs       []any   `json:"inputs"`
	PropertyName string  `json:"propertyName"`
	Value        float64 `json:"value"`
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
func (r MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionPeriodToWeeks) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionPeriodToWeeks) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionAnd struct {
	EnclosedInParentheses bool `json:"enclosedInParentheses,required"`
	// Any of "AND".
	Operator     string `json:"operator,required"`
	Inputs       []any  `json:"inputs"`
	PropertyName string `json:"propertyName"`
	Value        bool   `json:"value"`
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
func (r MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionAnd) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionAnd) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionOr struct {
	EnclosedInParentheses bool `json:"enclosedInParentheses,required"`
	// Any of "OR".
	Operator     string `json:"operator,required"`
	Inputs       []any  `json:"inputs"`
	PropertyName string `json:"propertyName"`
	Value        bool   `json:"value"`
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
func (r MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionOr) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionOr) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionXor struct {
	EnclosedInParentheses bool `json:"enclosedInParentheses,required"`
	// Any of "XOR".
	Operator     string `json:"operator,required"`
	Inputs       []any  `json:"inputs"`
	PropertyName string `json:"propertyName"`
	Value        bool   `json:"value"`
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
func (r MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionXor) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionXor) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionIfString struct {
	EnclosedInParentheses bool `json:"enclosedInParentheses,required"`
	IfExpression          any  `json:"ifExpression,required"`
	// Any of "IF_STRING".
	Operator       string `json:"operator,required"`
	ElseExpression any    `json:"elseExpression"`
	Inputs         []any  `json:"inputs"`
	PropertyName   string `json:"propertyName"`
	Value          string `json:"value"`
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
func (r MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionIfString) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionIfString) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionIfNumber struct {
	EnclosedInParentheses bool `json:"enclosedInParentheses,required"`
	IfExpression          any  `json:"ifExpression,required"`
	// Any of "IF_NUMBER".
	Operator       string  `json:"operator,required"`
	ElseExpression any     `json:"elseExpression"`
	Inputs         []any   `json:"inputs"`
	PropertyName   string  `json:"propertyName"`
	Value          float64 `json:"value"`
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
func (r MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionIfNumber) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionIfNumber) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionIfBoolean struct {
	EnclosedInParentheses bool `json:"enclosedInParentheses,required"`
	IfExpression          any  `json:"ifExpression,required"`
	// Any of "IF_BOOLEAN".
	Operator       string `json:"operator,required"`
	ElseExpression any    `json:"elseExpression"`
	Inputs         []any  `json:"inputs"`
	PropertyName   string `json:"propertyName"`
	Value          bool   `json:"value"`
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
func (r MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionIfBoolean) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionIfBoolean) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionIsPresent struct {
	ExpressionToEvaluate any `json:"expressionToEvaluate,required"`
	// Any of "IS_PRESENT".
	Operator     string `json:"operator,required"`
	Inputs       []any  `json:"inputs"`
	PropertyName string `json:"propertyName"`
	Value        bool   `json:"value"`
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
func (r MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionIsPresent) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionIsPresent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionHasEmailReply struct {
	// Any of "HAS_EMAIL_REPLY".
	Operator     string `json:"operator,required"`
	Inputs       []any  `json:"inputs"`
	PropertyName string `json:"propertyName"`
	Value        bool   `json:"value"`
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
func (r MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionHasEmailReply) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionHasEmailReply) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionHasPlainTextEmailReply struct {
	// Any of "HAS_PLAIN_TEXT_EMAIL_REPLY".
	Operator     string `json:"operator,required"`
	Inputs       []any  `json:"inputs"`
	PropertyName string `json:"propertyName"`
	Value        bool   `json:"value"`
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
func (r MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionHasPlainTextEmailReply) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionHasPlainTextEmailReply) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionExtractMostRecentEmailReplyHTML struct {
	// Any of "EXTRACT_MOST_RECENT_EMAIL_REPLY_HTML".
	Operator     string `json:"operator,required"`
	Inputs       []any  `json:"inputs"`
	PropertyName string `json:"propertyName"`
	Value        string `json:"value"`
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
func (r MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionExtractMostRecentEmailReplyHTML) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionExtractMostRecentEmailReplyHTML) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionExtractMostRecentEmailReplyText struct {
	// Any of "EXTRACT_MOST_RECENT_EMAIL_REPLY_TEXT".
	Operator     string `json:"operator,required"`
	Inputs       []any  `json:"inputs"`
	PropertyName string `json:"propertyName"`
	Value        string `json:"value"`
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
func (r MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionExtractMostRecentEmailReplyText) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionExtractMostRecentEmailReplyText) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionExtractMostRecentPlainTextEmailReply struct {
	// Any of "EXTRACT_MOST_RECENT_PLAIN_TEXT_EMAIL_REPLY".
	Operator     string `json:"operator,required"`
	Inputs       []any  `json:"inputs"`
	PropertyName string `json:"propertyName"`
	Value        string `json:"value"`
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
func (r MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionExtractMostRecentPlainTextEmailReply) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionExtractMostRecentPlainTextEmailReply) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionSetContainsString struct {
	// Any of "SET_CONTAINS_STRING".
	Operator      string `json:"operator,required"`
	StringToCheck any    `json:"stringToCheck,required"`
	Inputs        []any  `json:"inputs"`
	PropertyName  string `json:"propertyName"`
	Value         bool   `json:"value"`
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
func (r MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionSetContainsString) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionSetContainsString) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionIsEngagementType struct {
	// Any of "IS_ENGAGEMENT_TYPE".
	Operator     string `json:"operator,required"`
	Inputs       []any  `json:"inputs"`
	PropertyName string `json:"propertyName"`
	Value        bool   `json:"value"`
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
func (r MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionIsEngagementType) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionIsEngagementType) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionFormatFullName struct {
	// Any of "FORMAT_FULL_NAME".
	Operator     string `json:"operator,required"`
	Inputs       []any  `json:"inputs"`
	PropertyName string `json:"propertyName"`
	Value        string `json:"value"`
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
func (r MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionFormatFullName) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionFormatFullName) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionAbsoluteValue struct {
	// Any of "ABSOLUTE_VALUE".
	Operator     string  `json:"operator,required"`
	Inputs       []any   `json:"inputs"`
	PropertyName string  `json:"propertyName"`
	Value        float64 `json:"value"`
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
func (r MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionAbsoluteValue) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionAbsoluteValue) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionSquareRoot struct {
	// Any of "SQUARE_ROOT".
	Operator     string  `json:"operator,required"`
	Inputs       []any   `json:"inputs"`
	PropertyName string  `json:"propertyName"`
	Value        float64 `json:"value"`
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
func (r MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionSquareRoot) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionSquareRoot) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionPower struct {
	// Any of "POWER".
	Operator     string  `json:"operator,required"`
	Inputs       []any   `json:"inputs"`
	PropertyName string  `json:"propertyName"`
	Value        float64 `json:"value"`
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
func (r MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionPower) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionPower) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionSubstring struct {
	// Any of "SUBSTRING".
	Operator      string `json:"operator,required"`
	StringToCheck any    `json:"stringToCheck,required"`
	Inputs        []any  `json:"inputs"`
	PropertyName  string `json:"propertyName"`
	Value         string `json:"value"`
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
func (r MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionSubstring) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionSubstring) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionEuler struct {
	// Any of "EULER".
	Operator     string  `json:"operator,required"`
	Inputs       []any   `json:"inputs"`
	PropertyName string  `json:"propertyName"`
	Value        float64 `json:"value"`
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
func (r MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionEuler) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionEuler) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionStringLength struct {
	// Any of "STRING_LENGTH".
	Operator     string  `json:"operator,required"`
	Inputs       []any   `json:"inputs"`
	PropertyName string  `json:"propertyName"`
	Value        float64 `json:"value"`
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
func (r MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionStringLength) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionStringLength) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionAddTime struct {
	// Any of "ADD_TIME".
	Operator      string  `json:"operator,required"`
	StringToCheck any     `json:"stringToCheck,required"`
	Inputs        []any   `json:"inputs"`
	PropertyName  string  `json:"propertyName"`
	Value         float64 `json:"value"`
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
func (r MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionAddTime) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionAddTime) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionSubtractTime struct {
	// Any of "SUBTRACT_TIME".
	Operator      string  `json:"operator,required"`
	StringToCheck any     `json:"stringToCheck,required"`
	Inputs        []any   `json:"inputs"`
	PropertyName  string  `json:"propertyName"`
	Value         float64 `json:"value"`
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
func (r MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionSubtractTime) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyCalculationExpressionSubtractTime) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyDefinitionSource struct {
	// Any of "GLOBAL", "OBJECT_TYPE", "HAVEN_BRANCH", "PORTAL".
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
func (r MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyDefinitionSource) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyDefinitionSource) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyExtensionData struct {
	ExtensionStatusMap                  map[string]string                                                                                                            `json:"extensionStatusMap,required"`
	Tags                                []string                                                                                                                     `json:"tags,required"`
	CaseChangeTestExtensionData         MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyExtensionDataCaseChangeTestExtensionData         `json:"caseChangeTestExtensionData"`
	OptionDecoratorsExtensionData       MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyExtensionDataOptionDecoratorsExtensionData       `json:"optionDecoratorsExtensionData"`
	RequiredPropertiesExtensionData     MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyExtensionDataRequiredPropertiesExtensionData     `json:"requiredPropertiesExtensionData"`
	SoftRequiredPropertiesExtensionData MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyExtensionDataSoftRequiredPropertiesExtensionData `json:"softRequiredPropertiesExtensionData"`
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
func (r MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyExtensionData) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyExtensionData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyExtensionDataCaseChangeTestExtensionData struct {
	Mood string `json:"mood,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Mood        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyExtensionDataCaseChangeTestExtensionData) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyExtensionDataCaseChangeTestExtensionData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyExtensionDataOptionDecoratorsExtensionData struct {
	OptionDecorators     map[string]MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyExtensionDataOptionDecoratorsExtensionDataOptionDecorator `json:"optionDecorators,required"`
	OptionDecoratorStyle string                                                                                                                                           `json:"optionDecoratorStyle,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		OptionDecorators     respjson.Field
		OptionDecoratorStyle respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyExtensionDataOptionDecoratorsExtensionData) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyExtensionDataOptionDecoratorsExtensionData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyExtensionDataOptionDecoratorsExtensionDataOptionDecorator struct {
	Color string `json:"color,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Color       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyExtensionDataOptionDecoratorsExtensionDataOptionDecorator) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyExtensionDataOptionDecoratorsExtensionDataOptionDecorator) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyExtensionDataRequiredPropertiesExtensionData struct {
	IsRequiredProperty bool `json:"isRequiredProperty,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		IsRequiredProperty respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyExtensionDataRequiredPropertiesExtensionData) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyExtensionDataRequiredPropertiesExtensionData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyExtensionDataSoftRequiredPropertiesExtensionData struct {
	IsSoftRequiredProperty bool `json:"isSoftRequiredProperty,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		IsSoftRequiredProperty respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyExtensionDataSoftRequiredPropertiesExtensionData) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyExtensionDataSoftRequiredPropertiesExtensionData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyExternalOptionsMetaData struct {
	Filter              MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyExternalOptionsMetaDataFilter `json:"filter"`
	RelatedObjectTypeID string                                                                                                    `json:"relatedObjectTypeId"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Filter              respjson.Field
		RelatedObjectTypeID respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyExternalOptionsMetaData) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyExternalOptionsMetaData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyExternalOptionsMetaDataFilter struct {
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
func (r MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyExternalOptionsMetaDataFilter) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyExternalOptionsMetaDataFilter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyPermission struct {
	AccessLevel string `json:"accessLevel,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AccessLevel respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyPermission) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyPermission) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyPropertyDefinitionSource struct {
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
func (r MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyPropertyDefinitionSource) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyPropertyDefinitionSource) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyPropertyRequirements struct {
	Gates []string `json:"gates,required"`
	// Any of "AND", "OR".
	Operator   string   `json:"operator,required"`
	ScopeNames []string `json:"scopeNames,required"`
	Settings   []string `json:"settings,required"`
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
func (r MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyPropertyRequirements) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyPropertyRequirements) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpression struct {
	AssociationTypes            []shared.AssociationSpec                                                                                               `json:"associationTypes,required"`
	RollupOperator              string                                                                                                                 `json:"rollupOperator,required"`
	SourceObjectTypeID          string                                                                                                                 `json:"sourceObjectTypeId,required"`
	SourcePropertyName          string                                                                                                                 `json:"sourcePropertyName,required"`
	ConditionalExpression       MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionUnion `json:"conditionalExpression"`
	ConditionalFormula          string                                                                                                                 `json:"conditionalFormula"`
	EmptyRollupValue            string                                                                                                                 `json:"emptyRollupValue"`
	SourceCompareByPropertyName string                                                                                                                 `json:"sourceCompareByPropertyName"`
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
func (r MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpression) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpression) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionUnion
// contains all possible properties and values from
// [MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionConstantBoolean],
// [MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionConstantNumber],
// [MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionConstantString],
// [MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionBooleanPropertyVariable],
// [MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionStringPropertyVariable],
// [MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionNumberPropertyVariable],
// [MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionTimestampOfPropertyVariable],
// [MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionBooleanTargetPropertyVariable],
// [MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionStringTargetPropertyVariable],
// [MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionNumberTargetPropertyVariable],
// [MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionTimestampOfTargetPropertyVariable],
// [MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionAddNumbers],
// [MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionSubtractNumbers],
// [MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionMultiplyNumbers],
// [MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionDivideNumbers],
// [MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionRoundDown],
// [MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionRoundUp],
// [MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionRoundNearest],
// [MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionUpperCase],
// [MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionLowerCase],
// [MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionConcatStrings],
// [MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionContains],
// [MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionBeginsWith],
// [MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionNumberToString],
// [MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionParseNumber],
// [MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionFetchExchangeRate],
// [MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionFetchCurrencyDecimalPlaces],
// [MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionFetchSingleCurrencyPortalCurrency],
// [MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionDatedExchangeRate],
// [MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionPipelineProbability],
// [MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionMaxNumbers],
// [MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionMinNumbers],
// [MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionLessThan],
// [MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionLessThanOrEqual],
// [MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionMoreThan],
// [MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionMoreThanOrEqual],
// [MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionNumberEquals],
// [MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionStringEquals],
// [MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionIsPipelineStageClosed],
// [MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionNot],
// [MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionDate],
// [MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionMonth],
// [MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionYear],
// [MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionNow],
// [MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionTimeBetween],
// [MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionPeriodToMonths],
// [MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionPeriodToWeeks],
// [MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionAnd],
// [MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionOr],
// [MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionXor],
// [MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionIfString],
// [MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionIfNumber],
// [MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionIfBoolean],
// [MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionIsPresent],
// [MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionHasEmailReply],
// [MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionHasPlainTextEmailReply],
// [MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionExtractMostRecentEmailReplyHTML],
// [MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionExtractMostRecentEmailReplyText],
// [MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionExtractMostRecentPlainTextEmailReply],
// [MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionSetContainsString],
// [MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionIsEngagementType],
// [MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionFormatFullName],
// [MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionAbsoluteValue],
// [MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionSquareRoot],
// [MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionPower],
// [MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionSubstring],
// [MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionEuler],
// [MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionStringLength],
// [MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionAddTime],
// [MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionSubtractTime].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionUnion struct {
	Operator     string `json:"operator"`
	Inputs       []any  `json:"inputs"`
	PropertyName string `json:"propertyName"`
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
	Value                 MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionUnionValue `json:"value"`
	EnclosedInParentheses bool                                                                                                                        `json:"enclosedInParentheses"`
	StringToCheck         any                                                                                                                         `json:"stringToCheck"`
	IfExpression          any                                                                                                                         `json:"ifExpression"`
	ElseExpression        any                                                                                                                         `json:"elseExpression"`
	// This field is from variant
	// [MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionIsPresent].
	ExpressionToEvaluate any `json:"expressionToEvaluate"`
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

func (u MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionUnion) AsConstantBoolean() (v MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionConstantBoolean) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionUnion) AsConstantNumber() (v MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionConstantNumber) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionUnion) AsConstantString() (v MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionConstantString) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionUnion) AsBooleanPropertyVariable() (v MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionBooleanPropertyVariable) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionUnion) AsStringPropertyVariable() (v MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionStringPropertyVariable) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionUnion) AsNumberPropertyVariable() (v MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionNumberPropertyVariable) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionUnion) AsTimestampOfPropertyVariable() (v MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionTimestampOfPropertyVariable) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionUnion) AsBooleanTargetPropertyVariable() (v MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionBooleanTargetPropertyVariable) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionUnion) AsStringTargetPropertyVariable() (v MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionStringTargetPropertyVariable) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionUnion) AsNumberTargetPropertyVariable() (v MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionNumberTargetPropertyVariable) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionUnion) AsTimestampOfTargetPropertyVariable() (v MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionTimestampOfTargetPropertyVariable) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionUnion) AsAddNumbers() (v MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionAddNumbers) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionUnion) AsSubtractNumbers() (v MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionSubtractNumbers) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionUnion) AsMultiplyNumbers() (v MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionMultiplyNumbers) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionUnion) AsDivideNumbers() (v MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionDivideNumbers) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionUnion) AsRoundDown() (v MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionRoundDown) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionUnion) AsRoundUp() (v MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionRoundUp) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionUnion) AsRoundNearest() (v MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionRoundNearest) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionUnion) AsUpperCase() (v MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionUpperCase) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionUnion) AsLowerCase() (v MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionLowerCase) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionUnion) AsConcatStrings() (v MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionConcatStrings) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionUnion) AsContains() (v MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionContains) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionUnion) AsBeginsWith() (v MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionBeginsWith) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionUnion) AsNumberToString() (v MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionNumberToString) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionUnion) AsParseNumber() (v MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionParseNumber) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionUnion) AsFetchExchangeRate() (v MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionFetchExchangeRate) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionUnion) AsFetchCurrencyDecimalPlaces() (v MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionFetchCurrencyDecimalPlaces) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionUnion) AsFetchSingleCurrencyPortalCurrency() (v MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionFetchSingleCurrencyPortalCurrency) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionUnion) AsDatedExchangeRate() (v MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionDatedExchangeRate) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionUnion) AsPipelineProbability() (v MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionPipelineProbability) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionUnion) AsMaxNumbers() (v MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionMaxNumbers) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionUnion) AsMinNumbers() (v MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionMinNumbers) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionUnion) AsLessThan() (v MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionLessThan) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionUnion) AsLessThanOrEqual() (v MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionLessThanOrEqual) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionUnion) AsMoreThan() (v MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionMoreThan) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionUnion) AsMoreThanOrEqual() (v MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionMoreThanOrEqual) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionUnion) AsNumberEquals() (v MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionNumberEquals) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionUnion) AsStringEquals() (v MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionStringEquals) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionUnion) AsIsPipelineStageClosed() (v MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionIsPipelineStageClosed) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionUnion) AsNot() (v MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionNot) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionUnion) AsDate() (v MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionDate) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionUnion) AsMonth() (v MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionMonth) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionUnion) AsYear() (v MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionYear) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionUnion) AsNow() (v MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionNow) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionUnion) AsTimeBetween() (v MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionTimeBetween) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionUnion) AsPeriodToMonths() (v MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionPeriodToMonths) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionUnion) AsPeriodToWeeks() (v MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionPeriodToWeeks) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionUnion) AsAnd() (v MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionAnd) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionUnion) AsOr() (v MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionOr) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionUnion) AsXor() (v MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionXor) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionUnion) AsIfString() (v MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionIfString) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionUnion) AsIfNumber() (v MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionIfNumber) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionUnion) AsIfBoolean() (v MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionIfBoolean) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionUnion) AsIsPresent() (v MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionIsPresent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionUnion) AsHasEmailReply() (v MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionHasEmailReply) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionUnion) AsHasPlainTextEmailReply() (v MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionHasPlainTextEmailReply) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionUnion) AsExtractMostRecentEmailReplyHTML() (v MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionExtractMostRecentEmailReplyHTML) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionUnion) AsExtractMostRecentEmailReplyText() (v MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionExtractMostRecentEmailReplyText) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionUnion) AsExtractMostRecentPlainTextEmailReply() (v MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionExtractMostRecentPlainTextEmailReply) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionUnion) AsSetContainsString() (v MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionSetContainsString) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionUnion) AsIsEngagementType() (v MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionIsEngagementType) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionUnion) AsFormatFullName() (v MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionFormatFullName) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionUnion) AsAbsoluteValue() (v MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionAbsoluteValue) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionUnion) AsSquareRoot() (v MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionSquareRoot) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionUnion) AsPower() (v MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionPower) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionUnion) AsSubstring() (v MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionSubstring) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionUnion) AsEuler() (v MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionEuler) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionUnion) AsStringLength() (v MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionStringLength) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionUnion) AsAddTime() (v MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionAddTime) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionUnion) AsSubtractTime() (v MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionSubtractTime) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionUnionValue
// is an implicit subunion of
// [MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionUnion].
// MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionUnionValue
// provides convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfBool OfFloat OfString]
type MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionUnionValue struct {
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

func (r *MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionUnionValue) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionConstantBoolean struct {
	// Any of "CONSTANT_BOOLEAN".
	Operator     string `json:"operator,required"`
	Inputs       []any  `json:"inputs"`
	PropertyName string `json:"propertyName"`
	Value        bool   `json:"value"`
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
func (r MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionConstantBoolean) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionConstantBoolean) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionConstantNumber struct {
	// Any of "CONSTANT_NUMBER".
	Operator     string  `json:"operator,required"`
	Inputs       []any   `json:"inputs"`
	PropertyName string  `json:"propertyName"`
	Value        float64 `json:"value"`
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
func (r MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionConstantNumber) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionConstantNumber) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionConstantString struct {
	// Any of "CONSTANT_STRING".
	Operator     string `json:"operator,required"`
	Inputs       []any  `json:"inputs"`
	PropertyName string `json:"propertyName"`
	Value        string `json:"value"`
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
func (r MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionConstantString) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionConstantString) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionBooleanPropertyVariable struct {
	// Any of "BOOLEAN_PROPERTY_VARIABLE".
	Operator     string `json:"operator,required"`
	Inputs       []any  `json:"inputs"`
	PropertyName string `json:"propertyName"`
	Value        bool   `json:"value"`
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
func (r MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionBooleanPropertyVariable) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionBooleanPropertyVariable) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionStringPropertyVariable struct {
	// Any of "STRING_PROPERTY_VARIABLE".
	Operator     string `json:"operator,required"`
	Inputs       []any  `json:"inputs"`
	PropertyName string `json:"propertyName"`
	Value        string `json:"value"`
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
func (r MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionStringPropertyVariable) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionStringPropertyVariable) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionNumberPropertyVariable struct {
	// Any of "NUMBER_PROPERTY_VARIABLE".
	Operator     string  `json:"operator,required"`
	Inputs       []any   `json:"inputs"`
	PropertyName string  `json:"propertyName"`
	Value        float64 `json:"value"`
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
func (r MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionNumberPropertyVariable) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionNumberPropertyVariable) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionTimestampOfPropertyVariable struct {
	// Any of "TIMESTAMP_OF_PROPERTY_VARIABLE".
	Operator     string `json:"operator,required"`
	Inputs       []any  `json:"inputs"`
	PropertyName string `json:"propertyName"`
	Value        string `json:"value"`
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
func (r MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionTimestampOfPropertyVariable) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionTimestampOfPropertyVariable) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionBooleanTargetPropertyVariable struct {
	// Any of "BOOLEAN_TARGET_PROPERTY_VARIABLE".
	Operator     string `json:"operator,required"`
	Inputs       []any  `json:"inputs"`
	PropertyName string `json:"propertyName"`
	Value        bool   `json:"value"`
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
func (r MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionBooleanTargetPropertyVariable) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionBooleanTargetPropertyVariable) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionStringTargetPropertyVariable struct {
	// Any of "STRING_TARGET_PROPERTY_VARIABLE".
	Operator     string `json:"operator,required"`
	Inputs       []any  `json:"inputs"`
	PropertyName string `json:"propertyName"`
	Value        string `json:"value"`
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
func (r MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionStringTargetPropertyVariable) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionStringTargetPropertyVariable) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionNumberTargetPropertyVariable struct {
	// Any of "NUMBER_TARGET_PROPERTY_VARIABLE".
	Operator     string  `json:"operator,required"`
	Inputs       []any   `json:"inputs"`
	PropertyName string  `json:"propertyName"`
	Value        float64 `json:"value"`
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
func (r MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionNumberTargetPropertyVariable) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionNumberTargetPropertyVariable) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionTimestampOfTargetPropertyVariable struct {
	// Any of "TIMESTAMP_OF_TARGET_PROPERTY_VARIABLE".
	Operator     string `json:"operator,required"`
	Inputs       []any  `json:"inputs"`
	PropertyName string `json:"propertyName"`
	Value        string `json:"value"`
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
func (r MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionTimestampOfTargetPropertyVariable) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionTimestampOfTargetPropertyVariable) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionAddNumbers struct {
	EnclosedInParentheses bool `json:"enclosedInParentheses,required"`
	// Any of "ADD_NUMBERS".
	Operator     string  `json:"operator,required"`
	Inputs       []any   `json:"inputs"`
	PropertyName string  `json:"propertyName"`
	Value        float64 `json:"value"`
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
func (r MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionAddNumbers) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionAddNumbers) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionSubtractNumbers struct {
	EnclosedInParentheses bool `json:"enclosedInParentheses,required"`
	// Any of "SUBTRACT_NUMBERS".
	Operator     string  `json:"operator,required"`
	Inputs       []any   `json:"inputs"`
	PropertyName string  `json:"propertyName"`
	Value        float64 `json:"value"`
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
func (r MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionSubtractNumbers) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionSubtractNumbers) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionMultiplyNumbers struct {
	EnclosedInParentheses bool `json:"enclosedInParentheses,required"`
	// Any of "MULTIPLY_NUMBERS".
	Operator     string  `json:"operator,required"`
	Inputs       []any   `json:"inputs"`
	PropertyName string  `json:"propertyName"`
	Value        float64 `json:"value"`
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
func (r MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionMultiplyNumbers) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionMultiplyNumbers) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionDivideNumbers struct {
	EnclosedInParentheses bool `json:"enclosedInParentheses,required"`
	// Any of "DIVIDE_NUMBERS".
	Operator     string  `json:"operator,required"`
	Inputs       []any   `json:"inputs"`
	PropertyName string  `json:"propertyName"`
	Value        float64 `json:"value"`
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
func (r MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionDivideNumbers) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionDivideNumbers) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionRoundDown struct {
	// Any of "ROUND_DOWN".
	Operator     string  `json:"operator,required"`
	Inputs       []any   `json:"inputs"`
	PropertyName string  `json:"propertyName"`
	Value        float64 `json:"value"`
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
func (r MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionRoundDown) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionRoundDown) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionRoundUp struct {
	// Any of "ROUND_UP".
	Operator     string  `json:"operator,required"`
	Inputs       []any   `json:"inputs"`
	PropertyName string  `json:"propertyName"`
	Value        float64 `json:"value"`
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
func (r MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionRoundUp) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionRoundUp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionRoundNearest struct {
	// Any of "ROUND_NEAREST".
	Operator     string  `json:"operator,required"`
	Inputs       []any   `json:"inputs"`
	PropertyName string  `json:"propertyName"`
	Value        float64 `json:"value"`
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
func (r MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionRoundNearest) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionRoundNearest) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionUpperCase struct {
	// Any of "UPPER_CASE".
	Operator     string `json:"operator,required"`
	Inputs       []any  `json:"inputs"`
	PropertyName string `json:"propertyName"`
	Value        string `json:"value"`
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
func (r MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionUpperCase) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionUpperCase) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionLowerCase struct {
	// Any of "LOWER_CASE".
	Operator     string `json:"operator,required"`
	Inputs       []any  `json:"inputs"`
	PropertyName string `json:"propertyName"`
	Value        string `json:"value"`
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
func (r MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionLowerCase) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionLowerCase) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionConcatStrings struct {
	// Any of "CONCAT_STRINGS".
	Operator     string `json:"operator,required"`
	Inputs       []any  `json:"inputs"`
	PropertyName string `json:"propertyName"`
	Value        string `json:"value"`
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
func (r MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionConcatStrings) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionConcatStrings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionContains struct {
	// Any of "CONTAINS".
	Operator      string `json:"operator,required"`
	StringToCheck any    `json:"stringToCheck,required"`
	Inputs        []any  `json:"inputs"`
	PropertyName  string `json:"propertyName"`
	Value         bool   `json:"value"`
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
func (r MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionContains) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionContains) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionBeginsWith struct {
	// Any of "BEGINS_WITH".
	Operator      string `json:"operator,required"`
	StringToCheck any    `json:"stringToCheck,required"`
	Inputs        []any  `json:"inputs"`
	PropertyName  string `json:"propertyName"`
	Value         bool   `json:"value"`
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
func (r MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionBeginsWith) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionBeginsWith) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionNumberToString struct {
	// Any of "NUMBER_TO_STRING".
	Operator     string `json:"operator,required"`
	Inputs       []any  `json:"inputs"`
	PropertyName string `json:"propertyName"`
	Value        string `json:"value"`
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
func (r MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionNumberToString) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionNumberToString) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionParseNumber struct {
	// Any of "PARSE_NUMBER".
	Operator     string  `json:"operator,required"`
	Inputs       []any   `json:"inputs"`
	PropertyName string  `json:"propertyName"`
	Value        float64 `json:"value"`
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
func (r MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionParseNumber) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionParseNumber) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionFetchExchangeRate struct {
	// Any of "FETCH_EXCHANGE_RATE".
	Operator     string  `json:"operator,required"`
	Inputs       []any   `json:"inputs"`
	PropertyName string  `json:"propertyName"`
	Value        float64 `json:"value"`
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
func (r MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionFetchExchangeRate) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionFetchExchangeRate) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionFetchCurrencyDecimalPlaces struct {
	// Any of "FETCH_CURRENCY_DECIMAL_PLACES".
	Operator     string  `json:"operator,required"`
	Inputs       []any   `json:"inputs"`
	PropertyName string  `json:"propertyName"`
	Value        float64 `json:"value"`
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
func (r MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionFetchCurrencyDecimalPlaces) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionFetchCurrencyDecimalPlaces) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionFetchSingleCurrencyPortalCurrency struct {
	// Any of "FETCH_SINGLE_CURRENCY_PORTAL_CURRENCY".
	Operator     string `json:"operator,required"`
	Inputs       []any  `json:"inputs"`
	PropertyName string `json:"propertyName"`
	Value        string `json:"value"`
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
func (r MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionFetchSingleCurrencyPortalCurrency) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionFetchSingleCurrencyPortalCurrency) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionDatedExchangeRate struct {
	// Any of "DATED_EXCHANGE_RATE".
	Operator     string  `json:"operator,required"`
	Inputs       []any   `json:"inputs"`
	PropertyName string  `json:"propertyName"`
	Value        float64 `json:"value"`
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
func (r MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionDatedExchangeRate) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionDatedExchangeRate) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionPipelineProbability struct {
	// Any of "PIPELINE_PROBABILITY".
	Operator     string  `json:"operator,required"`
	Inputs       []any   `json:"inputs"`
	PropertyName string  `json:"propertyName"`
	Value        float64 `json:"value"`
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
func (r MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionPipelineProbability) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionPipelineProbability) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionMaxNumbers struct {
	// Any of "MAX_NUMBERS".
	Operator     string  `json:"operator,required"`
	Inputs       []any   `json:"inputs"`
	PropertyName string  `json:"propertyName"`
	Value        float64 `json:"value"`
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
func (r MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionMaxNumbers) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionMaxNumbers) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionMinNumbers struct {
	// Any of "MIN_NUMBERS".
	Operator     string  `json:"operator,required"`
	Inputs       []any   `json:"inputs"`
	PropertyName string  `json:"propertyName"`
	Value        float64 `json:"value"`
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
func (r MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionMinNumbers) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionMinNumbers) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionLessThan struct {
	// Any of "LESS_THAN".
	Operator     string `json:"operator,required"`
	Inputs       []any  `json:"inputs"`
	PropertyName string `json:"propertyName"`
	Value        bool   `json:"value"`
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
func (r MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionLessThan) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionLessThan) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionLessThanOrEqual struct {
	// Any of "LESS_THAN_OR_EQUAL".
	Operator     string `json:"operator,required"`
	Inputs       []any  `json:"inputs"`
	PropertyName string `json:"propertyName"`
	Value        bool   `json:"value"`
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
func (r MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionLessThanOrEqual) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionLessThanOrEqual) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionMoreThan struct {
	// Any of "MORE_THAN".
	Operator     string `json:"operator,required"`
	Inputs       []any  `json:"inputs"`
	PropertyName string `json:"propertyName"`
	Value        bool   `json:"value"`
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
func (r MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionMoreThan) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionMoreThan) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionMoreThanOrEqual struct {
	// Any of "MORE_THAN_OR_EQUAL".
	Operator     string `json:"operator,required"`
	Inputs       []any  `json:"inputs"`
	PropertyName string `json:"propertyName"`
	Value        bool   `json:"value"`
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
func (r MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionMoreThanOrEqual) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionMoreThanOrEqual) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionNumberEquals struct {
	// Any of "NUMBER_EQUALS".
	Operator     string `json:"operator,required"`
	Inputs       []any  `json:"inputs"`
	PropertyName string `json:"propertyName"`
	Value        bool   `json:"value"`
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
func (r MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionNumberEquals) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionNumberEquals) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionStringEquals struct {
	// Any of "STRING_EQUALS".
	Operator     string `json:"operator,required"`
	Inputs       []any  `json:"inputs"`
	PropertyName string `json:"propertyName"`
	Value        bool   `json:"value"`
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
func (r MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionStringEquals) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionStringEquals) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionIsPipelineStageClosed struct {
	// Any of "IS_PIPELINE_STAGE_CLOSED".
	Operator     string `json:"operator,required"`
	Inputs       []any  `json:"inputs"`
	PropertyName string `json:"propertyName"`
	Value        bool   `json:"value"`
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
func (r MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionIsPipelineStageClosed) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionIsPipelineStageClosed) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionNot struct {
	// Any of "NOT".
	Operator     string `json:"operator,required"`
	Inputs       []any  `json:"inputs"`
	PropertyName string `json:"propertyName"`
	Value        bool   `json:"value"`
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
func (r MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionNot) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionNot) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionDate struct {
	// Any of "DATE".
	Operator     string  `json:"operator,required"`
	Inputs       []any   `json:"inputs"`
	PropertyName string  `json:"propertyName"`
	Value        float64 `json:"value"`
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
func (r MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionDate) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionDate) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionMonth struct {
	// Any of "MONTH".
	Operator     string  `json:"operator,required"`
	Inputs       []any   `json:"inputs"`
	PropertyName string  `json:"propertyName"`
	Value        float64 `json:"value"`
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
func (r MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionMonth) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionMonth) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionYear struct {
	// Any of "YEAR".
	Operator     string  `json:"operator,required"`
	Inputs       []any   `json:"inputs"`
	PropertyName string  `json:"propertyName"`
	Value        float64 `json:"value"`
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
func (r MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionYear) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionYear) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionNow struct {
	// Any of "NOW".
	Operator     string  `json:"operator,required"`
	Inputs       []any   `json:"inputs"`
	PropertyName string  `json:"propertyName"`
	Value        float64 `json:"value"`
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
func (r MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionNow) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionNow) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionTimeBetween struct {
	// Any of "TIME_BETWEEN".
	Operator     string  `json:"operator,required"`
	Inputs       []any   `json:"inputs"`
	PropertyName string  `json:"propertyName"`
	Value        float64 `json:"value"`
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
func (r MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionTimeBetween) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionTimeBetween) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionPeriodToMonths struct {
	// Any of "PERIOD_TO_MONTHS".
	Operator     string  `json:"operator,required"`
	Inputs       []any   `json:"inputs"`
	PropertyName string  `json:"propertyName"`
	Value        float64 `json:"value"`
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
func (r MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionPeriodToMonths) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionPeriodToMonths) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionPeriodToWeeks struct {
	// Any of "PERIOD_TO_WEEKS".
	Operator     string  `json:"operator,required"`
	Inputs       []any   `json:"inputs"`
	PropertyName string  `json:"propertyName"`
	Value        float64 `json:"value"`
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
func (r MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionPeriodToWeeks) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionPeriodToWeeks) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionAnd struct {
	EnclosedInParentheses bool `json:"enclosedInParentheses,required"`
	// Any of "AND".
	Operator     string `json:"operator,required"`
	Inputs       []any  `json:"inputs"`
	PropertyName string `json:"propertyName"`
	Value        bool   `json:"value"`
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
func (r MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionAnd) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionAnd) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionOr struct {
	EnclosedInParentheses bool `json:"enclosedInParentheses,required"`
	// Any of "OR".
	Operator     string `json:"operator,required"`
	Inputs       []any  `json:"inputs"`
	PropertyName string `json:"propertyName"`
	Value        bool   `json:"value"`
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
func (r MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionOr) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionOr) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionXor struct {
	EnclosedInParentheses bool `json:"enclosedInParentheses,required"`
	// Any of "XOR".
	Operator     string `json:"operator,required"`
	Inputs       []any  `json:"inputs"`
	PropertyName string `json:"propertyName"`
	Value        bool   `json:"value"`
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
func (r MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionXor) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionXor) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionIfString struct {
	EnclosedInParentheses bool `json:"enclosedInParentheses,required"`
	IfExpression          any  `json:"ifExpression,required"`
	// Any of "IF_STRING".
	Operator       string `json:"operator,required"`
	ElseExpression any    `json:"elseExpression"`
	Inputs         []any  `json:"inputs"`
	PropertyName   string `json:"propertyName"`
	Value          string `json:"value"`
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
func (r MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionIfString) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionIfString) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionIfNumber struct {
	EnclosedInParentheses bool `json:"enclosedInParentheses,required"`
	IfExpression          any  `json:"ifExpression,required"`
	// Any of "IF_NUMBER".
	Operator       string  `json:"operator,required"`
	ElseExpression any     `json:"elseExpression"`
	Inputs         []any   `json:"inputs"`
	PropertyName   string  `json:"propertyName"`
	Value          float64 `json:"value"`
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
func (r MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionIfNumber) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionIfNumber) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionIfBoolean struct {
	EnclosedInParentheses bool `json:"enclosedInParentheses,required"`
	IfExpression          any  `json:"ifExpression,required"`
	// Any of "IF_BOOLEAN".
	Operator       string `json:"operator,required"`
	ElseExpression any    `json:"elseExpression"`
	Inputs         []any  `json:"inputs"`
	PropertyName   string `json:"propertyName"`
	Value          bool   `json:"value"`
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
func (r MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionIfBoolean) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionIfBoolean) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionIsPresent struct {
	ExpressionToEvaluate any `json:"expressionToEvaluate,required"`
	// Any of "IS_PRESENT".
	Operator     string `json:"operator,required"`
	Inputs       []any  `json:"inputs"`
	PropertyName string `json:"propertyName"`
	Value        bool   `json:"value"`
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
func (r MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionIsPresent) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionIsPresent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionHasEmailReply struct {
	// Any of "HAS_EMAIL_REPLY".
	Operator     string `json:"operator,required"`
	Inputs       []any  `json:"inputs"`
	PropertyName string `json:"propertyName"`
	Value        bool   `json:"value"`
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
func (r MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionHasEmailReply) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionHasEmailReply) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionHasPlainTextEmailReply struct {
	// Any of "HAS_PLAIN_TEXT_EMAIL_REPLY".
	Operator     string `json:"operator,required"`
	Inputs       []any  `json:"inputs"`
	PropertyName string `json:"propertyName"`
	Value        bool   `json:"value"`
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
func (r MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionHasPlainTextEmailReply) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionHasPlainTextEmailReply) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionExtractMostRecentEmailReplyHTML struct {
	// Any of "EXTRACT_MOST_RECENT_EMAIL_REPLY_HTML".
	Operator     string `json:"operator,required"`
	Inputs       []any  `json:"inputs"`
	PropertyName string `json:"propertyName"`
	Value        string `json:"value"`
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
func (r MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionExtractMostRecentEmailReplyHTML) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionExtractMostRecentEmailReplyHTML) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionExtractMostRecentEmailReplyText struct {
	// Any of "EXTRACT_MOST_RECENT_EMAIL_REPLY_TEXT".
	Operator     string `json:"operator,required"`
	Inputs       []any  `json:"inputs"`
	PropertyName string `json:"propertyName"`
	Value        string `json:"value"`
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
func (r MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionExtractMostRecentEmailReplyText) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionExtractMostRecentEmailReplyText) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionExtractMostRecentPlainTextEmailReply struct {
	// Any of "EXTRACT_MOST_RECENT_PLAIN_TEXT_EMAIL_REPLY".
	Operator     string `json:"operator,required"`
	Inputs       []any  `json:"inputs"`
	PropertyName string `json:"propertyName"`
	Value        string `json:"value"`
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
func (r MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionExtractMostRecentPlainTextEmailReply) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionExtractMostRecentPlainTextEmailReply) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionSetContainsString struct {
	// Any of "SET_CONTAINS_STRING".
	Operator      string `json:"operator,required"`
	StringToCheck any    `json:"stringToCheck,required"`
	Inputs        []any  `json:"inputs"`
	PropertyName  string `json:"propertyName"`
	Value         bool   `json:"value"`
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
func (r MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionSetContainsString) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionSetContainsString) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionIsEngagementType struct {
	// Any of "IS_ENGAGEMENT_TYPE".
	Operator     string `json:"operator,required"`
	Inputs       []any  `json:"inputs"`
	PropertyName string `json:"propertyName"`
	Value        bool   `json:"value"`
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
func (r MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionIsEngagementType) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionIsEngagementType) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionFormatFullName struct {
	// Any of "FORMAT_FULL_NAME".
	Operator     string `json:"operator,required"`
	Inputs       []any  `json:"inputs"`
	PropertyName string `json:"propertyName"`
	Value        string `json:"value"`
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
func (r MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionFormatFullName) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionFormatFullName) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionAbsoluteValue struct {
	// Any of "ABSOLUTE_VALUE".
	Operator     string  `json:"operator,required"`
	Inputs       []any   `json:"inputs"`
	PropertyName string  `json:"propertyName"`
	Value        float64 `json:"value"`
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
func (r MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionAbsoluteValue) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionAbsoluteValue) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionSquareRoot struct {
	// Any of "SQUARE_ROOT".
	Operator     string  `json:"operator,required"`
	Inputs       []any   `json:"inputs"`
	PropertyName string  `json:"propertyName"`
	Value        float64 `json:"value"`
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
func (r MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionSquareRoot) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionSquareRoot) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionPower struct {
	// Any of "POWER".
	Operator     string  `json:"operator,required"`
	Inputs       []any   `json:"inputs"`
	PropertyName string  `json:"propertyName"`
	Value        float64 `json:"value"`
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
func (r MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionPower) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionPower) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionSubstring struct {
	// Any of "SUBSTRING".
	Operator      string `json:"operator,required"`
	StringToCheck any    `json:"stringToCheck,required"`
	Inputs        []any  `json:"inputs"`
	PropertyName  string `json:"propertyName"`
	Value         string `json:"value"`
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
func (r MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionSubstring) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionSubstring) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionEuler struct {
	// Any of "EULER".
	Operator     string  `json:"operator,required"`
	Inputs       []any   `json:"inputs"`
	PropertyName string  `json:"propertyName"`
	Value        float64 `json:"value"`
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
func (r MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionEuler) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionEuler) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionStringLength struct {
	// Any of "STRING_LENGTH".
	Operator     string  `json:"operator,required"`
	Inputs       []any   `json:"inputs"`
	PropertyName string  `json:"propertyName"`
	Value        float64 `json:"value"`
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
func (r MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionStringLength) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionStringLength) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionAddTime struct {
	// Any of "ADD_TIME".
	Operator      string  `json:"operator,required"`
	StringToCheck any     `json:"stringToCheck,required"`
	Inputs        []any   `json:"inputs"`
	PropertyName  string  `json:"propertyName"`
	Value         float64 `json:"value"`
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
func (r MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionAddTime) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionAddTime) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionSubtractTime struct {
	// Any of "SUBTRACT_TIME".
	Operator      string  `json:"operator,required"`
	StringToCheck any     `json:"stringToCheck,required"`
	Inputs        []any   `json:"inputs"`
	PropertyName  string  `json:"propertyName"`
	Value         float64 `json:"value"`
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
func (r MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionSubtractTime) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyRollupExpressionConditionalExpressionSubtractTime) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyGroup struct {
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
func (r MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyGroup) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingNewObjectDefinitionResponseCreatedObjectPropertyGroup) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingNewOembedDomainResponse struct {
	ID        int64                                                        `json:"id,required"`
	AppID     int64                                                        `json:"appId,required"`
	CreatedAt int64                                                        `json:"createdAt,required"`
	DeletedAt int64                                                        `json:"deletedAt,required"`
	Endpoints MediaBridgeIntegratorSettingNewOembedDomainResponseEndpoints `json:"endpoints,required"`
	PortalID  int64                                                        `json:"portalId,required"`
	UpdatedAt int64                                                        `json:"updatedAt,required"`
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
func (r MediaBridgeIntegratorSettingNewOembedDomainResponse) RawJSON() string { return r.JSON.raw }
func (r *MediaBridgeIntegratorSettingNewOembedDomainResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingNewOembedDomainResponseEndpoints struct {
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
func (r MediaBridgeIntegratorSettingNewOembedDomainResponseEndpoints) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingNewOembedDomainResponseEndpoints) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingGetEventVisibilitySettingsResponse struct {
	CreatedAt          time.Time                                                                         `json:"createdAt,required" format:"date-time"`
	VisibilitySettings []MediaBridgeIntegratorSettingGetEventVisibilitySettingsResponseVisibilitySetting `json:"visibilitySettings,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreatedAt          respjson.Field
		VisibilitySettings respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MediaBridgeIntegratorSettingGetEventVisibilitySettingsResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingGetEventVisibilitySettingsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingGetEventVisibilitySettingsResponseVisibilitySetting struct {
	// Any of "ALL", "MEDIA_PLAYS", "MEDIA_PLAYS_PERCENT", "ATTENTION_SPAN".
	EventType       string `json:"eventType,required"`
	UpdatedAt       int64  `json:"updatedAt,required"`
	ShowInReporting bool   `json:"showInReporting"`
	ShowInTimeline  bool   `json:"showInTimeline"`
	ShowInWorkflows bool   `json:"showInWorkflows"`
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
func (r MediaBridgeIntegratorSettingGetEventVisibilitySettingsResponseVisibilitySetting) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingGetEventVisibilitySettingsResponseVisibilitySetting) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponse struct {
	ObjectTypeID   string                                                                             `json:"objectTypeId,required"`
	ObjectTypeName string                                                                             `json:"objectTypeName,required"`
	Properties     []MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponseProperty      `json:"properties,required"`
	PropertyGroups []MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyGroup `json:"propertyGroups,required"`
	Schema         MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponseSchema          `json:"schema"`
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
func (r MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponseProperty struct {
	ObjectTypeID string `json:"objectTypeId,required"`
	// Defines a property
	Property                 shared.Property                                                                                       `json:"property,required"`
	CalculationExpression    MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionUnion `json:"calculationExpression"`
	CalculationFormula       string                                                                                                `json:"calculationFormula"`
	DefinitionSource         MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyDefinitionSource           `json:"definitionSource"`
	ExtensionData            MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyExtensionData              `json:"extensionData"`
	ExternalOptionsMetaData  MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyExternalOptionsMetaData    `json:"externalOptionsMetaData"`
	FulcrumPortalID          int64                                                                                                 `json:"fulcrumPortalId"`
	FulcrumTimestamp         int64                                                                                                 `json:"fulcrumTimestamp"`
	JanusGroup               string                                                                                                `json:"janusGroup"`
	Permission               MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyPermission                 `json:"permission"`
	PropertyDefinitionSource MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyPropertyDefinitionSource   `json:"propertyDefinitionSource"`
	PropertyRequirements     MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyPropertyRequirements       `json:"propertyRequirements"`
	RollupExpression         MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpression           `json:"rollupExpression"`
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
func (r MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponseProperty) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponseProperty) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionUnion
// contains all possible properties and values from
// [MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionConstantBoolean],
// [MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionConstantNumber],
// [MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionConstantString],
// [MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionBooleanPropertyVariable],
// [MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionStringPropertyVariable],
// [MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionNumberPropertyVariable],
// [MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionTimestampOfPropertyVariable],
// [MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionBooleanTargetPropertyVariable],
// [MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionStringTargetPropertyVariable],
// [MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionNumberTargetPropertyVariable],
// [MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionTimestampOfTargetPropertyVariable],
// [MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionAddNumbers],
// [MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionSubtractNumbers],
// [MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionMultiplyNumbers],
// [MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionDivideNumbers],
// [MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionRoundDown],
// [MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionRoundUp],
// [MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionRoundNearest],
// [MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionUpperCase],
// [MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionLowerCase],
// [MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionConcatStrings],
// [MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionContains],
// [MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionBeginsWith],
// [MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionNumberToString],
// [MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionParseNumber],
// [MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionFetchExchangeRate],
// [MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionFetchCurrencyDecimalPlaces],
// [MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionFetchSingleCurrencyPortalCurrency],
// [MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionDatedExchangeRate],
// [MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionPipelineProbability],
// [MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionMaxNumbers],
// [MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionMinNumbers],
// [MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionLessThan],
// [MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionLessThanOrEqual],
// [MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionMoreThan],
// [MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionMoreThanOrEqual],
// [MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionNumberEquals],
// [MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionStringEquals],
// [MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionIsPipelineStageClosed],
// [MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionNot],
// [MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionDate],
// [MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionMonth],
// [MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionYear],
// [MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionNow],
// [MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionTimeBetween],
// [MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionPeriodToMonths],
// [MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionPeriodToWeeks],
// [MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionAnd],
// [MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionOr],
// [MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionXor],
// [MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionIfString],
// [MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionIfNumber],
// [MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionIfBoolean],
// [MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionIsPresent],
// [MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionHasEmailReply],
// [MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionHasPlainTextEmailReply],
// [MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionExtractMostRecentEmailReplyHTML],
// [MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionExtractMostRecentEmailReplyText],
// [MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionExtractMostRecentPlainTextEmailReply],
// [MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionSetContainsString],
// [MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionIsEngagementType],
// [MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionFormatFullName],
// [MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionAbsoluteValue],
// [MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionSquareRoot],
// [MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionPower],
// [MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionSubstring],
// [MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionEuler],
// [MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionStringLength],
// [MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionAddTime],
// [MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionSubtractTime].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionUnion struct {
	Operator     string `json:"operator"`
	Inputs       []any  `json:"inputs"`
	PropertyName string `json:"propertyName"`
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
	Value                 MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionUnionValue `json:"value"`
	EnclosedInParentheses bool                                                                                                       `json:"enclosedInParentheses"`
	StringToCheck         any                                                                                                        `json:"stringToCheck"`
	IfExpression          any                                                                                                        `json:"ifExpression"`
	ElseExpression        any                                                                                                        `json:"elseExpression"`
	// This field is from variant
	// [MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionIsPresent].
	ExpressionToEvaluate any `json:"expressionToEvaluate"`
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

func (u MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionUnion) AsConstantBoolean() (v MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionConstantBoolean) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionUnion) AsConstantNumber() (v MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionConstantNumber) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionUnion) AsConstantString() (v MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionConstantString) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionUnion) AsBooleanPropertyVariable() (v MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionBooleanPropertyVariable) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionUnion) AsStringPropertyVariable() (v MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionStringPropertyVariable) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionUnion) AsNumberPropertyVariable() (v MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionNumberPropertyVariable) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionUnion) AsTimestampOfPropertyVariable() (v MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionTimestampOfPropertyVariable) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionUnion) AsBooleanTargetPropertyVariable() (v MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionBooleanTargetPropertyVariable) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionUnion) AsStringTargetPropertyVariable() (v MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionStringTargetPropertyVariable) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionUnion) AsNumberTargetPropertyVariable() (v MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionNumberTargetPropertyVariable) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionUnion) AsTimestampOfTargetPropertyVariable() (v MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionTimestampOfTargetPropertyVariable) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionUnion) AsAddNumbers() (v MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionAddNumbers) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionUnion) AsSubtractNumbers() (v MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionSubtractNumbers) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionUnion) AsMultiplyNumbers() (v MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionMultiplyNumbers) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionUnion) AsDivideNumbers() (v MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionDivideNumbers) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionUnion) AsRoundDown() (v MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionRoundDown) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionUnion) AsRoundUp() (v MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionRoundUp) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionUnion) AsRoundNearest() (v MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionRoundNearest) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionUnion) AsUpperCase() (v MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionUpperCase) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionUnion) AsLowerCase() (v MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionLowerCase) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionUnion) AsConcatStrings() (v MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionConcatStrings) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionUnion) AsContains() (v MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionContains) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionUnion) AsBeginsWith() (v MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionBeginsWith) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionUnion) AsNumberToString() (v MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionNumberToString) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionUnion) AsParseNumber() (v MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionParseNumber) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionUnion) AsFetchExchangeRate() (v MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionFetchExchangeRate) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionUnion) AsFetchCurrencyDecimalPlaces() (v MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionFetchCurrencyDecimalPlaces) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionUnion) AsFetchSingleCurrencyPortalCurrency() (v MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionFetchSingleCurrencyPortalCurrency) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionUnion) AsDatedExchangeRate() (v MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionDatedExchangeRate) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionUnion) AsPipelineProbability() (v MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionPipelineProbability) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionUnion) AsMaxNumbers() (v MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionMaxNumbers) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionUnion) AsMinNumbers() (v MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionMinNumbers) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionUnion) AsLessThan() (v MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionLessThan) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionUnion) AsLessThanOrEqual() (v MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionLessThanOrEqual) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionUnion) AsMoreThan() (v MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionMoreThan) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionUnion) AsMoreThanOrEqual() (v MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionMoreThanOrEqual) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionUnion) AsNumberEquals() (v MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionNumberEquals) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionUnion) AsStringEquals() (v MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionStringEquals) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionUnion) AsIsPipelineStageClosed() (v MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionIsPipelineStageClosed) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionUnion) AsNot() (v MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionNot) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionUnion) AsDate() (v MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionDate) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionUnion) AsMonth() (v MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionMonth) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionUnion) AsYear() (v MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionYear) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionUnion) AsNow() (v MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionNow) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionUnion) AsTimeBetween() (v MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionTimeBetween) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionUnion) AsPeriodToMonths() (v MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionPeriodToMonths) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionUnion) AsPeriodToWeeks() (v MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionPeriodToWeeks) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionUnion) AsAnd() (v MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionAnd) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionUnion) AsOr() (v MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionOr) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionUnion) AsXor() (v MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionXor) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionUnion) AsIfString() (v MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionIfString) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionUnion) AsIfNumber() (v MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionIfNumber) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionUnion) AsIfBoolean() (v MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionIfBoolean) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionUnion) AsIsPresent() (v MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionIsPresent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionUnion) AsHasEmailReply() (v MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionHasEmailReply) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionUnion) AsHasPlainTextEmailReply() (v MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionHasPlainTextEmailReply) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionUnion) AsExtractMostRecentEmailReplyHTML() (v MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionExtractMostRecentEmailReplyHTML) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionUnion) AsExtractMostRecentEmailReplyText() (v MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionExtractMostRecentEmailReplyText) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionUnion) AsExtractMostRecentPlainTextEmailReply() (v MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionExtractMostRecentPlainTextEmailReply) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionUnion) AsSetContainsString() (v MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionSetContainsString) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionUnion) AsIsEngagementType() (v MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionIsEngagementType) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionUnion) AsFormatFullName() (v MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionFormatFullName) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionUnion) AsAbsoluteValue() (v MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionAbsoluteValue) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionUnion) AsSquareRoot() (v MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionSquareRoot) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionUnion) AsPower() (v MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionPower) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionUnion) AsSubstring() (v MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionSubstring) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionUnion) AsEuler() (v MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionEuler) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionUnion) AsStringLength() (v MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionStringLength) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionUnion) AsAddTime() (v MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionAddTime) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionUnion) AsSubtractTime() (v MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionSubtractTime) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionUnionValue
// is an implicit subunion of
// [MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionUnion].
// MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionUnionValue
// provides convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfBool OfFloat OfString]
type MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionUnionValue struct {
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

func (r *MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionUnionValue) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionConstantBoolean struct {
	// Any of "CONSTANT_BOOLEAN".
	Operator     string `json:"operator,required"`
	Inputs       []any  `json:"inputs"`
	PropertyName string `json:"propertyName"`
	Value        bool   `json:"value"`
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
func (r MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionConstantBoolean) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionConstantBoolean) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionConstantNumber struct {
	// Any of "CONSTANT_NUMBER".
	Operator     string  `json:"operator,required"`
	Inputs       []any   `json:"inputs"`
	PropertyName string  `json:"propertyName"`
	Value        float64 `json:"value"`
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
func (r MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionConstantNumber) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionConstantNumber) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionConstantString struct {
	// Any of "CONSTANT_STRING".
	Operator     string `json:"operator,required"`
	Inputs       []any  `json:"inputs"`
	PropertyName string `json:"propertyName"`
	Value        string `json:"value"`
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
func (r MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionConstantString) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionConstantString) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionBooleanPropertyVariable struct {
	// Any of "BOOLEAN_PROPERTY_VARIABLE".
	Operator     string `json:"operator,required"`
	Inputs       []any  `json:"inputs"`
	PropertyName string `json:"propertyName"`
	Value        bool   `json:"value"`
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
func (r MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionBooleanPropertyVariable) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionBooleanPropertyVariable) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionStringPropertyVariable struct {
	// Any of "STRING_PROPERTY_VARIABLE".
	Operator     string `json:"operator,required"`
	Inputs       []any  `json:"inputs"`
	PropertyName string `json:"propertyName"`
	Value        string `json:"value"`
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
func (r MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionStringPropertyVariable) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionStringPropertyVariable) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionNumberPropertyVariable struct {
	// Any of "NUMBER_PROPERTY_VARIABLE".
	Operator     string  `json:"operator,required"`
	Inputs       []any   `json:"inputs"`
	PropertyName string  `json:"propertyName"`
	Value        float64 `json:"value"`
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
func (r MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionNumberPropertyVariable) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionNumberPropertyVariable) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionTimestampOfPropertyVariable struct {
	// Any of "TIMESTAMP_OF_PROPERTY_VARIABLE".
	Operator     string `json:"operator,required"`
	Inputs       []any  `json:"inputs"`
	PropertyName string `json:"propertyName"`
	Value        string `json:"value"`
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
func (r MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionTimestampOfPropertyVariable) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionTimestampOfPropertyVariable) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionBooleanTargetPropertyVariable struct {
	// Any of "BOOLEAN_TARGET_PROPERTY_VARIABLE".
	Operator     string `json:"operator,required"`
	Inputs       []any  `json:"inputs"`
	PropertyName string `json:"propertyName"`
	Value        bool   `json:"value"`
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
func (r MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionBooleanTargetPropertyVariable) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionBooleanTargetPropertyVariable) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionStringTargetPropertyVariable struct {
	// Any of "STRING_TARGET_PROPERTY_VARIABLE".
	Operator     string `json:"operator,required"`
	Inputs       []any  `json:"inputs"`
	PropertyName string `json:"propertyName"`
	Value        string `json:"value"`
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
func (r MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionStringTargetPropertyVariable) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionStringTargetPropertyVariable) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionNumberTargetPropertyVariable struct {
	// Any of "NUMBER_TARGET_PROPERTY_VARIABLE".
	Operator     string  `json:"operator,required"`
	Inputs       []any   `json:"inputs"`
	PropertyName string  `json:"propertyName"`
	Value        float64 `json:"value"`
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
func (r MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionNumberTargetPropertyVariable) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionNumberTargetPropertyVariable) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionTimestampOfTargetPropertyVariable struct {
	// Any of "TIMESTAMP_OF_TARGET_PROPERTY_VARIABLE".
	Operator     string `json:"operator,required"`
	Inputs       []any  `json:"inputs"`
	PropertyName string `json:"propertyName"`
	Value        string `json:"value"`
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
func (r MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionTimestampOfTargetPropertyVariable) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionTimestampOfTargetPropertyVariable) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionAddNumbers struct {
	EnclosedInParentheses bool `json:"enclosedInParentheses,required"`
	// Any of "ADD_NUMBERS".
	Operator     string  `json:"operator,required"`
	Inputs       []any   `json:"inputs"`
	PropertyName string  `json:"propertyName"`
	Value        float64 `json:"value"`
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
func (r MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionAddNumbers) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionAddNumbers) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionSubtractNumbers struct {
	EnclosedInParentheses bool `json:"enclosedInParentheses,required"`
	// Any of "SUBTRACT_NUMBERS".
	Operator     string  `json:"operator,required"`
	Inputs       []any   `json:"inputs"`
	PropertyName string  `json:"propertyName"`
	Value        float64 `json:"value"`
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
func (r MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionSubtractNumbers) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionSubtractNumbers) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionMultiplyNumbers struct {
	EnclosedInParentheses bool `json:"enclosedInParentheses,required"`
	// Any of "MULTIPLY_NUMBERS".
	Operator     string  `json:"operator,required"`
	Inputs       []any   `json:"inputs"`
	PropertyName string  `json:"propertyName"`
	Value        float64 `json:"value"`
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
func (r MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionMultiplyNumbers) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionMultiplyNumbers) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionDivideNumbers struct {
	EnclosedInParentheses bool `json:"enclosedInParentheses,required"`
	// Any of "DIVIDE_NUMBERS".
	Operator     string  `json:"operator,required"`
	Inputs       []any   `json:"inputs"`
	PropertyName string  `json:"propertyName"`
	Value        float64 `json:"value"`
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
func (r MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionDivideNumbers) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionDivideNumbers) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionRoundDown struct {
	// Any of "ROUND_DOWN".
	Operator     string  `json:"operator,required"`
	Inputs       []any   `json:"inputs"`
	PropertyName string  `json:"propertyName"`
	Value        float64 `json:"value"`
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
func (r MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionRoundDown) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionRoundDown) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionRoundUp struct {
	// Any of "ROUND_UP".
	Operator     string  `json:"operator,required"`
	Inputs       []any   `json:"inputs"`
	PropertyName string  `json:"propertyName"`
	Value        float64 `json:"value"`
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
func (r MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionRoundUp) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionRoundUp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionRoundNearest struct {
	// Any of "ROUND_NEAREST".
	Operator     string  `json:"operator,required"`
	Inputs       []any   `json:"inputs"`
	PropertyName string  `json:"propertyName"`
	Value        float64 `json:"value"`
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
func (r MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionRoundNearest) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionRoundNearest) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionUpperCase struct {
	// Any of "UPPER_CASE".
	Operator     string `json:"operator,required"`
	Inputs       []any  `json:"inputs"`
	PropertyName string `json:"propertyName"`
	Value        string `json:"value"`
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
func (r MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionUpperCase) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionUpperCase) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionLowerCase struct {
	// Any of "LOWER_CASE".
	Operator     string `json:"operator,required"`
	Inputs       []any  `json:"inputs"`
	PropertyName string `json:"propertyName"`
	Value        string `json:"value"`
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
func (r MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionLowerCase) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionLowerCase) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionConcatStrings struct {
	// Any of "CONCAT_STRINGS".
	Operator     string `json:"operator,required"`
	Inputs       []any  `json:"inputs"`
	PropertyName string `json:"propertyName"`
	Value        string `json:"value"`
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
func (r MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionConcatStrings) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionConcatStrings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionContains struct {
	// Any of "CONTAINS".
	Operator      string `json:"operator,required"`
	StringToCheck any    `json:"stringToCheck,required"`
	Inputs        []any  `json:"inputs"`
	PropertyName  string `json:"propertyName"`
	Value         bool   `json:"value"`
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
func (r MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionContains) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionContains) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionBeginsWith struct {
	// Any of "BEGINS_WITH".
	Operator      string `json:"operator,required"`
	StringToCheck any    `json:"stringToCheck,required"`
	Inputs        []any  `json:"inputs"`
	PropertyName  string `json:"propertyName"`
	Value         bool   `json:"value"`
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
func (r MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionBeginsWith) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionBeginsWith) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionNumberToString struct {
	// Any of "NUMBER_TO_STRING".
	Operator     string `json:"operator,required"`
	Inputs       []any  `json:"inputs"`
	PropertyName string `json:"propertyName"`
	Value        string `json:"value"`
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
func (r MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionNumberToString) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionNumberToString) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionParseNumber struct {
	// Any of "PARSE_NUMBER".
	Operator     string  `json:"operator,required"`
	Inputs       []any   `json:"inputs"`
	PropertyName string  `json:"propertyName"`
	Value        float64 `json:"value"`
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
func (r MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionParseNumber) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionParseNumber) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionFetchExchangeRate struct {
	// Any of "FETCH_EXCHANGE_RATE".
	Operator     string  `json:"operator,required"`
	Inputs       []any   `json:"inputs"`
	PropertyName string  `json:"propertyName"`
	Value        float64 `json:"value"`
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
func (r MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionFetchExchangeRate) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionFetchExchangeRate) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionFetchCurrencyDecimalPlaces struct {
	// Any of "FETCH_CURRENCY_DECIMAL_PLACES".
	Operator     string  `json:"operator,required"`
	Inputs       []any   `json:"inputs"`
	PropertyName string  `json:"propertyName"`
	Value        float64 `json:"value"`
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
func (r MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionFetchCurrencyDecimalPlaces) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionFetchCurrencyDecimalPlaces) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionFetchSingleCurrencyPortalCurrency struct {
	// Any of "FETCH_SINGLE_CURRENCY_PORTAL_CURRENCY".
	Operator     string `json:"operator,required"`
	Inputs       []any  `json:"inputs"`
	PropertyName string `json:"propertyName"`
	Value        string `json:"value"`
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
func (r MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionFetchSingleCurrencyPortalCurrency) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionFetchSingleCurrencyPortalCurrency) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionDatedExchangeRate struct {
	// Any of "DATED_EXCHANGE_RATE".
	Operator     string  `json:"operator,required"`
	Inputs       []any   `json:"inputs"`
	PropertyName string  `json:"propertyName"`
	Value        float64 `json:"value"`
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
func (r MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionDatedExchangeRate) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionDatedExchangeRate) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionPipelineProbability struct {
	// Any of "PIPELINE_PROBABILITY".
	Operator     string  `json:"operator,required"`
	Inputs       []any   `json:"inputs"`
	PropertyName string  `json:"propertyName"`
	Value        float64 `json:"value"`
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
func (r MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionPipelineProbability) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionPipelineProbability) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionMaxNumbers struct {
	// Any of "MAX_NUMBERS".
	Operator     string  `json:"operator,required"`
	Inputs       []any   `json:"inputs"`
	PropertyName string  `json:"propertyName"`
	Value        float64 `json:"value"`
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
func (r MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionMaxNumbers) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionMaxNumbers) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionMinNumbers struct {
	// Any of "MIN_NUMBERS".
	Operator     string  `json:"operator,required"`
	Inputs       []any   `json:"inputs"`
	PropertyName string  `json:"propertyName"`
	Value        float64 `json:"value"`
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
func (r MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionMinNumbers) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionMinNumbers) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionLessThan struct {
	// Any of "LESS_THAN".
	Operator     string `json:"operator,required"`
	Inputs       []any  `json:"inputs"`
	PropertyName string `json:"propertyName"`
	Value        bool   `json:"value"`
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
func (r MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionLessThan) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionLessThan) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionLessThanOrEqual struct {
	// Any of "LESS_THAN_OR_EQUAL".
	Operator     string `json:"operator,required"`
	Inputs       []any  `json:"inputs"`
	PropertyName string `json:"propertyName"`
	Value        bool   `json:"value"`
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
func (r MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionLessThanOrEqual) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionLessThanOrEqual) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionMoreThan struct {
	// Any of "MORE_THAN".
	Operator     string `json:"operator,required"`
	Inputs       []any  `json:"inputs"`
	PropertyName string `json:"propertyName"`
	Value        bool   `json:"value"`
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
func (r MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionMoreThan) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionMoreThan) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionMoreThanOrEqual struct {
	// Any of "MORE_THAN_OR_EQUAL".
	Operator     string `json:"operator,required"`
	Inputs       []any  `json:"inputs"`
	PropertyName string `json:"propertyName"`
	Value        bool   `json:"value"`
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
func (r MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionMoreThanOrEqual) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionMoreThanOrEqual) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionNumberEquals struct {
	// Any of "NUMBER_EQUALS".
	Operator     string `json:"operator,required"`
	Inputs       []any  `json:"inputs"`
	PropertyName string `json:"propertyName"`
	Value        bool   `json:"value"`
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
func (r MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionNumberEquals) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionNumberEquals) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionStringEquals struct {
	// Any of "STRING_EQUALS".
	Operator     string `json:"operator,required"`
	Inputs       []any  `json:"inputs"`
	PropertyName string `json:"propertyName"`
	Value        bool   `json:"value"`
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
func (r MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionStringEquals) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionStringEquals) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionIsPipelineStageClosed struct {
	// Any of "IS_PIPELINE_STAGE_CLOSED".
	Operator     string `json:"operator,required"`
	Inputs       []any  `json:"inputs"`
	PropertyName string `json:"propertyName"`
	Value        bool   `json:"value"`
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
func (r MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionIsPipelineStageClosed) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionIsPipelineStageClosed) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionNot struct {
	// Any of "NOT".
	Operator     string `json:"operator,required"`
	Inputs       []any  `json:"inputs"`
	PropertyName string `json:"propertyName"`
	Value        bool   `json:"value"`
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
func (r MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionNot) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionNot) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionDate struct {
	// Any of "DATE".
	Operator     string  `json:"operator,required"`
	Inputs       []any   `json:"inputs"`
	PropertyName string  `json:"propertyName"`
	Value        float64 `json:"value"`
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
func (r MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionDate) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionDate) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionMonth struct {
	// Any of "MONTH".
	Operator     string  `json:"operator,required"`
	Inputs       []any   `json:"inputs"`
	PropertyName string  `json:"propertyName"`
	Value        float64 `json:"value"`
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
func (r MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionMonth) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionMonth) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionYear struct {
	// Any of "YEAR".
	Operator     string  `json:"operator,required"`
	Inputs       []any   `json:"inputs"`
	PropertyName string  `json:"propertyName"`
	Value        float64 `json:"value"`
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
func (r MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionYear) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionYear) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionNow struct {
	// Any of "NOW".
	Operator     string  `json:"operator,required"`
	Inputs       []any   `json:"inputs"`
	PropertyName string  `json:"propertyName"`
	Value        float64 `json:"value"`
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
func (r MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionNow) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionNow) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionTimeBetween struct {
	// Any of "TIME_BETWEEN".
	Operator     string  `json:"operator,required"`
	Inputs       []any   `json:"inputs"`
	PropertyName string  `json:"propertyName"`
	Value        float64 `json:"value"`
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
func (r MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionTimeBetween) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionTimeBetween) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionPeriodToMonths struct {
	// Any of "PERIOD_TO_MONTHS".
	Operator     string  `json:"operator,required"`
	Inputs       []any   `json:"inputs"`
	PropertyName string  `json:"propertyName"`
	Value        float64 `json:"value"`
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
func (r MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionPeriodToMonths) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionPeriodToMonths) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionPeriodToWeeks struct {
	// Any of "PERIOD_TO_WEEKS".
	Operator     string  `json:"operator,required"`
	Inputs       []any   `json:"inputs"`
	PropertyName string  `json:"propertyName"`
	Value        float64 `json:"value"`
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
func (r MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionPeriodToWeeks) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionPeriodToWeeks) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionAnd struct {
	EnclosedInParentheses bool `json:"enclosedInParentheses,required"`
	// Any of "AND".
	Operator     string `json:"operator,required"`
	Inputs       []any  `json:"inputs"`
	PropertyName string `json:"propertyName"`
	Value        bool   `json:"value"`
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
func (r MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionAnd) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionAnd) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionOr struct {
	EnclosedInParentheses bool `json:"enclosedInParentheses,required"`
	// Any of "OR".
	Operator     string `json:"operator,required"`
	Inputs       []any  `json:"inputs"`
	PropertyName string `json:"propertyName"`
	Value        bool   `json:"value"`
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
func (r MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionOr) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionOr) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionXor struct {
	EnclosedInParentheses bool `json:"enclosedInParentheses,required"`
	// Any of "XOR".
	Operator     string `json:"operator,required"`
	Inputs       []any  `json:"inputs"`
	PropertyName string `json:"propertyName"`
	Value        bool   `json:"value"`
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
func (r MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionXor) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionXor) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionIfString struct {
	EnclosedInParentheses bool `json:"enclosedInParentheses,required"`
	IfExpression          any  `json:"ifExpression,required"`
	// Any of "IF_STRING".
	Operator       string `json:"operator,required"`
	ElseExpression any    `json:"elseExpression"`
	Inputs         []any  `json:"inputs"`
	PropertyName   string `json:"propertyName"`
	Value          string `json:"value"`
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
func (r MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionIfString) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionIfString) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionIfNumber struct {
	EnclosedInParentheses bool `json:"enclosedInParentheses,required"`
	IfExpression          any  `json:"ifExpression,required"`
	// Any of "IF_NUMBER".
	Operator       string  `json:"operator,required"`
	ElseExpression any     `json:"elseExpression"`
	Inputs         []any   `json:"inputs"`
	PropertyName   string  `json:"propertyName"`
	Value          float64 `json:"value"`
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
func (r MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionIfNumber) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionIfNumber) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionIfBoolean struct {
	EnclosedInParentheses bool `json:"enclosedInParentheses,required"`
	IfExpression          any  `json:"ifExpression,required"`
	// Any of "IF_BOOLEAN".
	Operator       string `json:"operator,required"`
	ElseExpression any    `json:"elseExpression"`
	Inputs         []any  `json:"inputs"`
	PropertyName   string `json:"propertyName"`
	Value          bool   `json:"value"`
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
func (r MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionIfBoolean) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionIfBoolean) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionIsPresent struct {
	ExpressionToEvaluate any `json:"expressionToEvaluate,required"`
	// Any of "IS_PRESENT".
	Operator     string `json:"operator,required"`
	Inputs       []any  `json:"inputs"`
	PropertyName string `json:"propertyName"`
	Value        bool   `json:"value"`
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
func (r MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionIsPresent) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionIsPresent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionHasEmailReply struct {
	// Any of "HAS_EMAIL_REPLY".
	Operator     string `json:"operator,required"`
	Inputs       []any  `json:"inputs"`
	PropertyName string `json:"propertyName"`
	Value        bool   `json:"value"`
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
func (r MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionHasEmailReply) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionHasEmailReply) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionHasPlainTextEmailReply struct {
	// Any of "HAS_PLAIN_TEXT_EMAIL_REPLY".
	Operator     string `json:"operator,required"`
	Inputs       []any  `json:"inputs"`
	PropertyName string `json:"propertyName"`
	Value        bool   `json:"value"`
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
func (r MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionHasPlainTextEmailReply) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionHasPlainTextEmailReply) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionExtractMostRecentEmailReplyHTML struct {
	// Any of "EXTRACT_MOST_RECENT_EMAIL_REPLY_HTML".
	Operator     string `json:"operator,required"`
	Inputs       []any  `json:"inputs"`
	PropertyName string `json:"propertyName"`
	Value        string `json:"value"`
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
func (r MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionExtractMostRecentEmailReplyHTML) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionExtractMostRecentEmailReplyHTML) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionExtractMostRecentEmailReplyText struct {
	// Any of "EXTRACT_MOST_RECENT_EMAIL_REPLY_TEXT".
	Operator     string `json:"operator,required"`
	Inputs       []any  `json:"inputs"`
	PropertyName string `json:"propertyName"`
	Value        string `json:"value"`
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
func (r MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionExtractMostRecentEmailReplyText) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionExtractMostRecentEmailReplyText) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionExtractMostRecentPlainTextEmailReply struct {
	// Any of "EXTRACT_MOST_RECENT_PLAIN_TEXT_EMAIL_REPLY".
	Operator     string `json:"operator,required"`
	Inputs       []any  `json:"inputs"`
	PropertyName string `json:"propertyName"`
	Value        string `json:"value"`
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
func (r MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionExtractMostRecentPlainTextEmailReply) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionExtractMostRecentPlainTextEmailReply) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionSetContainsString struct {
	// Any of "SET_CONTAINS_STRING".
	Operator      string `json:"operator,required"`
	StringToCheck any    `json:"stringToCheck,required"`
	Inputs        []any  `json:"inputs"`
	PropertyName  string `json:"propertyName"`
	Value         bool   `json:"value"`
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
func (r MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionSetContainsString) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionSetContainsString) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionIsEngagementType struct {
	// Any of "IS_ENGAGEMENT_TYPE".
	Operator     string `json:"operator,required"`
	Inputs       []any  `json:"inputs"`
	PropertyName string `json:"propertyName"`
	Value        bool   `json:"value"`
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
func (r MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionIsEngagementType) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionIsEngagementType) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionFormatFullName struct {
	// Any of "FORMAT_FULL_NAME".
	Operator     string `json:"operator,required"`
	Inputs       []any  `json:"inputs"`
	PropertyName string `json:"propertyName"`
	Value        string `json:"value"`
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
func (r MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionFormatFullName) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionFormatFullName) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionAbsoluteValue struct {
	// Any of "ABSOLUTE_VALUE".
	Operator     string  `json:"operator,required"`
	Inputs       []any   `json:"inputs"`
	PropertyName string  `json:"propertyName"`
	Value        float64 `json:"value"`
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
func (r MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionAbsoluteValue) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionAbsoluteValue) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionSquareRoot struct {
	// Any of "SQUARE_ROOT".
	Operator     string  `json:"operator,required"`
	Inputs       []any   `json:"inputs"`
	PropertyName string  `json:"propertyName"`
	Value        float64 `json:"value"`
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
func (r MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionSquareRoot) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionSquareRoot) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionPower struct {
	// Any of "POWER".
	Operator     string  `json:"operator,required"`
	Inputs       []any   `json:"inputs"`
	PropertyName string  `json:"propertyName"`
	Value        float64 `json:"value"`
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
func (r MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionPower) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionPower) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionSubstring struct {
	// Any of "SUBSTRING".
	Operator      string `json:"operator,required"`
	StringToCheck any    `json:"stringToCheck,required"`
	Inputs        []any  `json:"inputs"`
	PropertyName  string `json:"propertyName"`
	Value         string `json:"value"`
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
func (r MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionSubstring) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionSubstring) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionEuler struct {
	// Any of "EULER".
	Operator     string  `json:"operator,required"`
	Inputs       []any   `json:"inputs"`
	PropertyName string  `json:"propertyName"`
	Value        float64 `json:"value"`
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
func (r MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionEuler) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionEuler) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionStringLength struct {
	// Any of "STRING_LENGTH".
	Operator     string  `json:"operator,required"`
	Inputs       []any   `json:"inputs"`
	PropertyName string  `json:"propertyName"`
	Value        float64 `json:"value"`
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
func (r MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionStringLength) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionStringLength) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionAddTime struct {
	// Any of "ADD_TIME".
	Operator      string  `json:"operator,required"`
	StringToCheck any     `json:"stringToCheck,required"`
	Inputs        []any   `json:"inputs"`
	PropertyName  string  `json:"propertyName"`
	Value         float64 `json:"value"`
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
func (r MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionAddTime) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionAddTime) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionSubtractTime struct {
	// Any of "SUBTRACT_TIME".
	Operator      string  `json:"operator,required"`
	StringToCheck any     `json:"stringToCheck,required"`
	Inputs        []any   `json:"inputs"`
	PropertyName  string  `json:"propertyName"`
	Value         float64 `json:"value"`
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
func (r MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionSubtractTime) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyCalculationExpressionSubtractTime) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyDefinitionSource struct {
	// Any of "GLOBAL", "OBJECT_TYPE", "HAVEN_BRANCH", "PORTAL".
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
func (r MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyDefinitionSource) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyDefinitionSource) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyExtensionData struct {
	ExtensionStatusMap                  map[string]string                                                                                                           `json:"extensionStatusMap,required"`
	Tags                                []string                                                                                                                    `json:"tags,required"`
	CaseChangeTestExtensionData         MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyExtensionDataCaseChangeTestExtensionData         `json:"caseChangeTestExtensionData"`
	OptionDecoratorsExtensionData       MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyExtensionDataOptionDecoratorsExtensionData       `json:"optionDecoratorsExtensionData"`
	RequiredPropertiesExtensionData     MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyExtensionDataRequiredPropertiesExtensionData     `json:"requiredPropertiesExtensionData"`
	SoftRequiredPropertiesExtensionData MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyExtensionDataSoftRequiredPropertiesExtensionData `json:"softRequiredPropertiesExtensionData"`
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
func (r MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyExtensionData) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyExtensionData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyExtensionDataCaseChangeTestExtensionData struct {
	Mood string `json:"mood,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Mood        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyExtensionDataCaseChangeTestExtensionData) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyExtensionDataCaseChangeTestExtensionData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyExtensionDataOptionDecoratorsExtensionData struct {
	OptionDecorators     map[string]MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyExtensionDataOptionDecoratorsExtensionDataOptionDecorator `json:"optionDecorators,required"`
	OptionDecoratorStyle string                                                                                                                                          `json:"optionDecoratorStyle,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		OptionDecorators     respjson.Field
		OptionDecoratorStyle respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyExtensionDataOptionDecoratorsExtensionData) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyExtensionDataOptionDecoratorsExtensionData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyExtensionDataOptionDecoratorsExtensionDataOptionDecorator struct {
	Color string `json:"color,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Color       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyExtensionDataOptionDecoratorsExtensionDataOptionDecorator) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyExtensionDataOptionDecoratorsExtensionDataOptionDecorator) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyExtensionDataRequiredPropertiesExtensionData struct {
	IsRequiredProperty bool `json:"isRequiredProperty,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		IsRequiredProperty respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyExtensionDataRequiredPropertiesExtensionData) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyExtensionDataRequiredPropertiesExtensionData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyExtensionDataSoftRequiredPropertiesExtensionData struct {
	IsSoftRequiredProperty bool `json:"isSoftRequiredProperty,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		IsSoftRequiredProperty respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyExtensionDataSoftRequiredPropertiesExtensionData) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyExtensionDataSoftRequiredPropertiesExtensionData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyExternalOptionsMetaData struct {
	Filter              MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyExternalOptionsMetaDataFilter `json:"filter"`
	RelatedObjectTypeID string                                                                                                   `json:"relatedObjectTypeId"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Filter              respjson.Field
		RelatedObjectTypeID respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyExternalOptionsMetaData) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyExternalOptionsMetaData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyExternalOptionsMetaDataFilter struct {
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
func (r MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyExternalOptionsMetaDataFilter) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyExternalOptionsMetaDataFilter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyPermission struct {
	AccessLevel string `json:"accessLevel,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AccessLevel respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyPermission) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyPermission) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyPropertyDefinitionSource struct {
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
func (r MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyPropertyDefinitionSource) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyPropertyDefinitionSource) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyPropertyRequirements struct {
	Gates []string `json:"gates,required"`
	// Any of "AND", "OR".
	Operator   string   `json:"operator,required"`
	ScopeNames []string `json:"scopeNames,required"`
	Settings   []string `json:"settings,required"`
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
func (r MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyPropertyRequirements) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyPropertyRequirements) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpression struct {
	AssociationTypes            []shared.AssociationSpec                                                                                              `json:"associationTypes,required"`
	RollupOperator              string                                                                                                                `json:"rollupOperator,required"`
	SourceObjectTypeID          string                                                                                                                `json:"sourceObjectTypeId,required"`
	SourcePropertyName          string                                                                                                                `json:"sourcePropertyName,required"`
	ConditionalExpression       MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionUnion `json:"conditionalExpression"`
	ConditionalFormula          string                                                                                                                `json:"conditionalFormula"`
	EmptyRollupValue            string                                                                                                                `json:"emptyRollupValue"`
	SourceCompareByPropertyName string                                                                                                                `json:"sourceCompareByPropertyName"`
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
func (r MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpression) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpression) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionUnion
// contains all possible properties and values from
// [MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionConstantBoolean],
// [MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionConstantNumber],
// [MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionConstantString],
// [MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionBooleanPropertyVariable],
// [MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionStringPropertyVariable],
// [MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionNumberPropertyVariable],
// [MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionTimestampOfPropertyVariable],
// [MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionBooleanTargetPropertyVariable],
// [MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionStringTargetPropertyVariable],
// [MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionNumberTargetPropertyVariable],
// [MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionTimestampOfTargetPropertyVariable],
// [MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionAddNumbers],
// [MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionSubtractNumbers],
// [MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionMultiplyNumbers],
// [MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionDivideNumbers],
// [MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionRoundDown],
// [MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionRoundUp],
// [MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionRoundNearest],
// [MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionUpperCase],
// [MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionLowerCase],
// [MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionConcatStrings],
// [MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionContains],
// [MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionBeginsWith],
// [MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionNumberToString],
// [MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionParseNumber],
// [MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionFetchExchangeRate],
// [MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionFetchCurrencyDecimalPlaces],
// [MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionFetchSingleCurrencyPortalCurrency],
// [MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionDatedExchangeRate],
// [MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionPipelineProbability],
// [MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionMaxNumbers],
// [MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionMinNumbers],
// [MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionLessThan],
// [MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionLessThanOrEqual],
// [MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionMoreThan],
// [MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionMoreThanOrEqual],
// [MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionNumberEquals],
// [MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionStringEquals],
// [MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionIsPipelineStageClosed],
// [MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionNot],
// [MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionDate],
// [MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionMonth],
// [MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionYear],
// [MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionNow],
// [MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionTimeBetween],
// [MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionPeriodToMonths],
// [MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionPeriodToWeeks],
// [MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionAnd],
// [MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionOr],
// [MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionXor],
// [MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionIfString],
// [MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionIfNumber],
// [MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionIfBoolean],
// [MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionIsPresent],
// [MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionHasEmailReply],
// [MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionHasPlainTextEmailReply],
// [MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionExtractMostRecentEmailReplyHTML],
// [MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionExtractMostRecentEmailReplyText],
// [MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionExtractMostRecentPlainTextEmailReply],
// [MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionSetContainsString],
// [MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionIsEngagementType],
// [MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionFormatFullName],
// [MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionAbsoluteValue],
// [MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionSquareRoot],
// [MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionPower],
// [MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionSubstring],
// [MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionEuler],
// [MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionStringLength],
// [MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionAddTime],
// [MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionSubtractTime].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionUnion struct {
	Operator     string `json:"operator"`
	Inputs       []any  `json:"inputs"`
	PropertyName string `json:"propertyName"`
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
	Value                 MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionUnionValue `json:"value"`
	EnclosedInParentheses bool                                                                                                                       `json:"enclosedInParentheses"`
	StringToCheck         any                                                                                                                        `json:"stringToCheck"`
	IfExpression          any                                                                                                                        `json:"ifExpression"`
	ElseExpression        any                                                                                                                        `json:"elseExpression"`
	// This field is from variant
	// [MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionIsPresent].
	ExpressionToEvaluate any `json:"expressionToEvaluate"`
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

func (u MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionUnion) AsConstantBoolean() (v MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionConstantBoolean) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionUnion) AsConstantNumber() (v MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionConstantNumber) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionUnion) AsConstantString() (v MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionConstantString) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionUnion) AsBooleanPropertyVariable() (v MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionBooleanPropertyVariable) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionUnion) AsStringPropertyVariable() (v MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionStringPropertyVariable) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionUnion) AsNumberPropertyVariable() (v MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionNumberPropertyVariable) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionUnion) AsTimestampOfPropertyVariable() (v MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionTimestampOfPropertyVariable) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionUnion) AsBooleanTargetPropertyVariable() (v MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionBooleanTargetPropertyVariable) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionUnion) AsStringTargetPropertyVariable() (v MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionStringTargetPropertyVariable) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionUnion) AsNumberTargetPropertyVariable() (v MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionNumberTargetPropertyVariable) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionUnion) AsTimestampOfTargetPropertyVariable() (v MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionTimestampOfTargetPropertyVariable) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionUnion) AsAddNumbers() (v MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionAddNumbers) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionUnion) AsSubtractNumbers() (v MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionSubtractNumbers) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionUnion) AsMultiplyNumbers() (v MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionMultiplyNumbers) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionUnion) AsDivideNumbers() (v MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionDivideNumbers) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionUnion) AsRoundDown() (v MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionRoundDown) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionUnion) AsRoundUp() (v MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionRoundUp) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionUnion) AsRoundNearest() (v MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionRoundNearest) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionUnion) AsUpperCase() (v MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionUpperCase) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionUnion) AsLowerCase() (v MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionLowerCase) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionUnion) AsConcatStrings() (v MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionConcatStrings) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionUnion) AsContains() (v MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionContains) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionUnion) AsBeginsWith() (v MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionBeginsWith) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionUnion) AsNumberToString() (v MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionNumberToString) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionUnion) AsParseNumber() (v MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionParseNumber) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionUnion) AsFetchExchangeRate() (v MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionFetchExchangeRate) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionUnion) AsFetchCurrencyDecimalPlaces() (v MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionFetchCurrencyDecimalPlaces) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionUnion) AsFetchSingleCurrencyPortalCurrency() (v MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionFetchSingleCurrencyPortalCurrency) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionUnion) AsDatedExchangeRate() (v MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionDatedExchangeRate) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionUnion) AsPipelineProbability() (v MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionPipelineProbability) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionUnion) AsMaxNumbers() (v MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionMaxNumbers) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionUnion) AsMinNumbers() (v MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionMinNumbers) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionUnion) AsLessThan() (v MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionLessThan) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionUnion) AsLessThanOrEqual() (v MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionLessThanOrEqual) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionUnion) AsMoreThan() (v MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionMoreThan) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionUnion) AsMoreThanOrEqual() (v MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionMoreThanOrEqual) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionUnion) AsNumberEquals() (v MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionNumberEquals) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionUnion) AsStringEquals() (v MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionStringEquals) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionUnion) AsIsPipelineStageClosed() (v MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionIsPipelineStageClosed) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionUnion) AsNot() (v MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionNot) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionUnion) AsDate() (v MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionDate) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionUnion) AsMonth() (v MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionMonth) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionUnion) AsYear() (v MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionYear) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionUnion) AsNow() (v MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionNow) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionUnion) AsTimeBetween() (v MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionTimeBetween) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionUnion) AsPeriodToMonths() (v MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionPeriodToMonths) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionUnion) AsPeriodToWeeks() (v MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionPeriodToWeeks) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionUnion) AsAnd() (v MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionAnd) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionUnion) AsOr() (v MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionOr) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionUnion) AsXor() (v MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionXor) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionUnion) AsIfString() (v MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionIfString) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionUnion) AsIfNumber() (v MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionIfNumber) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionUnion) AsIfBoolean() (v MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionIfBoolean) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionUnion) AsIsPresent() (v MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionIsPresent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionUnion) AsHasEmailReply() (v MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionHasEmailReply) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionUnion) AsHasPlainTextEmailReply() (v MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionHasPlainTextEmailReply) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionUnion) AsExtractMostRecentEmailReplyHTML() (v MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionExtractMostRecentEmailReplyHTML) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionUnion) AsExtractMostRecentEmailReplyText() (v MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionExtractMostRecentEmailReplyText) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionUnion) AsExtractMostRecentPlainTextEmailReply() (v MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionExtractMostRecentPlainTextEmailReply) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionUnion) AsSetContainsString() (v MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionSetContainsString) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionUnion) AsIsEngagementType() (v MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionIsEngagementType) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionUnion) AsFormatFullName() (v MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionFormatFullName) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionUnion) AsAbsoluteValue() (v MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionAbsoluteValue) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionUnion) AsSquareRoot() (v MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionSquareRoot) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionUnion) AsPower() (v MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionPower) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionUnion) AsSubstring() (v MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionSubstring) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionUnion) AsEuler() (v MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionEuler) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionUnion) AsStringLength() (v MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionStringLength) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionUnion) AsAddTime() (v MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionAddTime) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionUnion) AsSubtractTime() (v MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionSubtractTime) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionUnionValue
// is an implicit subunion of
// [MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionUnion].
// MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionUnionValue
// provides convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfBool OfFloat OfString]
type MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionUnionValue struct {
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

func (r *MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionUnionValue) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionConstantBoolean struct {
	// Any of "CONSTANT_BOOLEAN".
	Operator     string `json:"operator,required"`
	Inputs       []any  `json:"inputs"`
	PropertyName string `json:"propertyName"`
	Value        bool   `json:"value"`
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
func (r MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionConstantBoolean) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionConstantBoolean) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionConstantNumber struct {
	// Any of "CONSTANT_NUMBER".
	Operator     string  `json:"operator,required"`
	Inputs       []any   `json:"inputs"`
	PropertyName string  `json:"propertyName"`
	Value        float64 `json:"value"`
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
func (r MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionConstantNumber) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionConstantNumber) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionConstantString struct {
	// Any of "CONSTANT_STRING".
	Operator     string `json:"operator,required"`
	Inputs       []any  `json:"inputs"`
	PropertyName string `json:"propertyName"`
	Value        string `json:"value"`
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
func (r MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionConstantString) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionConstantString) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionBooleanPropertyVariable struct {
	// Any of "BOOLEAN_PROPERTY_VARIABLE".
	Operator     string `json:"operator,required"`
	Inputs       []any  `json:"inputs"`
	PropertyName string `json:"propertyName"`
	Value        bool   `json:"value"`
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
func (r MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionBooleanPropertyVariable) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionBooleanPropertyVariable) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionStringPropertyVariable struct {
	// Any of "STRING_PROPERTY_VARIABLE".
	Operator     string `json:"operator,required"`
	Inputs       []any  `json:"inputs"`
	PropertyName string `json:"propertyName"`
	Value        string `json:"value"`
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
func (r MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionStringPropertyVariable) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionStringPropertyVariable) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionNumberPropertyVariable struct {
	// Any of "NUMBER_PROPERTY_VARIABLE".
	Operator     string  `json:"operator,required"`
	Inputs       []any   `json:"inputs"`
	PropertyName string  `json:"propertyName"`
	Value        float64 `json:"value"`
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
func (r MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionNumberPropertyVariable) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionNumberPropertyVariable) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionTimestampOfPropertyVariable struct {
	// Any of "TIMESTAMP_OF_PROPERTY_VARIABLE".
	Operator     string `json:"operator,required"`
	Inputs       []any  `json:"inputs"`
	PropertyName string `json:"propertyName"`
	Value        string `json:"value"`
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
func (r MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionTimestampOfPropertyVariable) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionTimestampOfPropertyVariable) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionBooleanTargetPropertyVariable struct {
	// Any of "BOOLEAN_TARGET_PROPERTY_VARIABLE".
	Operator     string `json:"operator,required"`
	Inputs       []any  `json:"inputs"`
	PropertyName string `json:"propertyName"`
	Value        bool   `json:"value"`
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
func (r MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionBooleanTargetPropertyVariable) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionBooleanTargetPropertyVariable) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionStringTargetPropertyVariable struct {
	// Any of "STRING_TARGET_PROPERTY_VARIABLE".
	Operator     string `json:"operator,required"`
	Inputs       []any  `json:"inputs"`
	PropertyName string `json:"propertyName"`
	Value        string `json:"value"`
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
func (r MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionStringTargetPropertyVariable) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionStringTargetPropertyVariable) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionNumberTargetPropertyVariable struct {
	// Any of "NUMBER_TARGET_PROPERTY_VARIABLE".
	Operator     string  `json:"operator,required"`
	Inputs       []any   `json:"inputs"`
	PropertyName string  `json:"propertyName"`
	Value        float64 `json:"value"`
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
func (r MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionNumberTargetPropertyVariable) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionNumberTargetPropertyVariable) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionTimestampOfTargetPropertyVariable struct {
	// Any of "TIMESTAMP_OF_TARGET_PROPERTY_VARIABLE".
	Operator     string `json:"operator,required"`
	Inputs       []any  `json:"inputs"`
	PropertyName string `json:"propertyName"`
	Value        string `json:"value"`
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
func (r MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionTimestampOfTargetPropertyVariable) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionTimestampOfTargetPropertyVariable) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionAddNumbers struct {
	EnclosedInParentheses bool `json:"enclosedInParentheses,required"`
	// Any of "ADD_NUMBERS".
	Operator     string  `json:"operator,required"`
	Inputs       []any   `json:"inputs"`
	PropertyName string  `json:"propertyName"`
	Value        float64 `json:"value"`
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
func (r MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionAddNumbers) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionAddNumbers) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionSubtractNumbers struct {
	EnclosedInParentheses bool `json:"enclosedInParentheses,required"`
	// Any of "SUBTRACT_NUMBERS".
	Operator     string  `json:"operator,required"`
	Inputs       []any   `json:"inputs"`
	PropertyName string  `json:"propertyName"`
	Value        float64 `json:"value"`
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
func (r MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionSubtractNumbers) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionSubtractNumbers) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionMultiplyNumbers struct {
	EnclosedInParentheses bool `json:"enclosedInParentheses,required"`
	// Any of "MULTIPLY_NUMBERS".
	Operator     string  `json:"operator,required"`
	Inputs       []any   `json:"inputs"`
	PropertyName string  `json:"propertyName"`
	Value        float64 `json:"value"`
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
func (r MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionMultiplyNumbers) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionMultiplyNumbers) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionDivideNumbers struct {
	EnclosedInParentheses bool `json:"enclosedInParentheses,required"`
	// Any of "DIVIDE_NUMBERS".
	Operator     string  `json:"operator,required"`
	Inputs       []any   `json:"inputs"`
	PropertyName string  `json:"propertyName"`
	Value        float64 `json:"value"`
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
func (r MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionDivideNumbers) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionDivideNumbers) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionRoundDown struct {
	// Any of "ROUND_DOWN".
	Operator     string  `json:"operator,required"`
	Inputs       []any   `json:"inputs"`
	PropertyName string  `json:"propertyName"`
	Value        float64 `json:"value"`
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
func (r MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionRoundDown) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionRoundDown) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionRoundUp struct {
	// Any of "ROUND_UP".
	Operator     string  `json:"operator,required"`
	Inputs       []any   `json:"inputs"`
	PropertyName string  `json:"propertyName"`
	Value        float64 `json:"value"`
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
func (r MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionRoundUp) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionRoundUp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionRoundNearest struct {
	// Any of "ROUND_NEAREST".
	Operator     string  `json:"operator,required"`
	Inputs       []any   `json:"inputs"`
	PropertyName string  `json:"propertyName"`
	Value        float64 `json:"value"`
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
func (r MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionRoundNearest) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionRoundNearest) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionUpperCase struct {
	// Any of "UPPER_CASE".
	Operator     string `json:"operator,required"`
	Inputs       []any  `json:"inputs"`
	PropertyName string `json:"propertyName"`
	Value        string `json:"value"`
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
func (r MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionUpperCase) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionUpperCase) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionLowerCase struct {
	// Any of "LOWER_CASE".
	Operator     string `json:"operator,required"`
	Inputs       []any  `json:"inputs"`
	PropertyName string `json:"propertyName"`
	Value        string `json:"value"`
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
func (r MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionLowerCase) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionLowerCase) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionConcatStrings struct {
	// Any of "CONCAT_STRINGS".
	Operator     string `json:"operator,required"`
	Inputs       []any  `json:"inputs"`
	PropertyName string `json:"propertyName"`
	Value        string `json:"value"`
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
func (r MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionConcatStrings) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionConcatStrings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionContains struct {
	// Any of "CONTAINS".
	Operator      string `json:"operator,required"`
	StringToCheck any    `json:"stringToCheck,required"`
	Inputs        []any  `json:"inputs"`
	PropertyName  string `json:"propertyName"`
	Value         bool   `json:"value"`
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
func (r MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionContains) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionContains) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionBeginsWith struct {
	// Any of "BEGINS_WITH".
	Operator      string `json:"operator,required"`
	StringToCheck any    `json:"stringToCheck,required"`
	Inputs        []any  `json:"inputs"`
	PropertyName  string `json:"propertyName"`
	Value         bool   `json:"value"`
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
func (r MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionBeginsWith) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionBeginsWith) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionNumberToString struct {
	// Any of "NUMBER_TO_STRING".
	Operator     string `json:"operator,required"`
	Inputs       []any  `json:"inputs"`
	PropertyName string `json:"propertyName"`
	Value        string `json:"value"`
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
func (r MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionNumberToString) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionNumberToString) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionParseNumber struct {
	// Any of "PARSE_NUMBER".
	Operator     string  `json:"operator,required"`
	Inputs       []any   `json:"inputs"`
	PropertyName string  `json:"propertyName"`
	Value        float64 `json:"value"`
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
func (r MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionParseNumber) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionParseNumber) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionFetchExchangeRate struct {
	// Any of "FETCH_EXCHANGE_RATE".
	Operator     string  `json:"operator,required"`
	Inputs       []any   `json:"inputs"`
	PropertyName string  `json:"propertyName"`
	Value        float64 `json:"value"`
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
func (r MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionFetchExchangeRate) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionFetchExchangeRate) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionFetchCurrencyDecimalPlaces struct {
	// Any of "FETCH_CURRENCY_DECIMAL_PLACES".
	Operator     string  `json:"operator,required"`
	Inputs       []any   `json:"inputs"`
	PropertyName string  `json:"propertyName"`
	Value        float64 `json:"value"`
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
func (r MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionFetchCurrencyDecimalPlaces) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionFetchCurrencyDecimalPlaces) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionFetchSingleCurrencyPortalCurrency struct {
	// Any of "FETCH_SINGLE_CURRENCY_PORTAL_CURRENCY".
	Operator     string `json:"operator,required"`
	Inputs       []any  `json:"inputs"`
	PropertyName string `json:"propertyName"`
	Value        string `json:"value"`
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
func (r MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionFetchSingleCurrencyPortalCurrency) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionFetchSingleCurrencyPortalCurrency) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionDatedExchangeRate struct {
	// Any of "DATED_EXCHANGE_RATE".
	Operator     string  `json:"operator,required"`
	Inputs       []any   `json:"inputs"`
	PropertyName string  `json:"propertyName"`
	Value        float64 `json:"value"`
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
func (r MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionDatedExchangeRate) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionDatedExchangeRate) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionPipelineProbability struct {
	// Any of "PIPELINE_PROBABILITY".
	Operator     string  `json:"operator,required"`
	Inputs       []any   `json:"inputs"`
	PropertyName string  `json:"propertyName"`
	Value        float64 `json:"value"`
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
func (r MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionPipelineProbability) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionPipelineProbability) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionMaxNumbers struct {
	// Any of "MAX_NUMBERS".
	Operator     string  `json:"operator,required"`
	Inputs       []any   `json:"inputs"`
	PropertyName string  `json:"propertyName"`
	Value        float64 `json:"value"`
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
func (r MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionMaxNumbers) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionMaxNumbers) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionMinNumbers struct {
	// Any of "MIN_NUMBERS".
	Operator     string  `json:"operator,required"`
	Inputs       []any   `json:"inputs"`
	PropertyName string  `json:"propertyName"`
	Value        float64 `json:"value"`
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
func (r MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionMinNumbers) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionMinNumbers) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionLessThan struct {
	// Any of "LESS_THAN".
	Operator     string `json:"operator,required"`
	Inputs       []any  `json:"inputs"`
	PropertyName string `json:"propertyName"`
	Value        bool   `json:"value"`
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
func (r MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionLessThan) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionLessThan) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionLessThanOrEqual struct {
	// Any of "LESS_THAN_OR_EQUAL".
	Operator     string `json:"operator,required"`
	Inputs       []any  `json:"inputs"`
	PropertyName string `json:"propertyName"`
	Value        bool   `json:"value"`
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
func (r MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionLessThanOrEqual) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionLessThanOrEqual) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionMoreThan struct {
	// Any of "MORE_THAN".
	Operator     string `json:"operator,required"`
	Inputs       []any  `json:"inputs"`
	PropertyName string `json:"propertyName"`
	Value        bool   `json:"value"`
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
func (r MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionMoreThan) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionMoreThan) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionMoreThanOrEqual struct {
	// Any of "MORE_THAN_OR_EQUAL".
	Operator     string `json:"operator,required"`
	Inputs       []any  `json:"inputs"`
	PropertyName string `json:"propertyName"`
	Value        bool   `json:"value"`
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
func (r MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionMoreThanOrEqual) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionMoreThanOrEqual) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionNumberEquals struct {
	// Any of "NUMBER_EQUALS".
	Operator     string `json:"operator,required"`
	Inputs       []any  `json:"inputs"`
	PropertyName string `json:"propertyName"`
	Value        bool   `json:"value"`
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
func (r MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionNumberEquals) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionNumberEquals) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionStringEquals struct {
	// Any of "STRING_EQUALS".
	Operator     string `json:"operator,required"`
	Inputs       []any  `json:"inputs"`
	PropertyName string `json:"propertyName"`
	Value        bool   `json:"value"`
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
func (r MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionStringEquals) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionStringEquals) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionIsPipelineStageClosed struct {
	// Any of "IS_PIPELINE_STAGE_CLOSED".
	Operator     string `json:"operator,required"`
	Inputs       []any  `json:"inputs"`
	PropertyName string `json:"propertyName"`
	Value        bool   `json:"value"`
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
func (r MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionIsPipelineStageClosed) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionIsPipelineStageClosed) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionNot struct {
	// Any of "NOT".
	Operator     string `json:"operator,required"`
	Inputs       []any  `json:"inputs"`
	PropertyName string `json:"propertyName"`
	Value        bool   `json:"value"`
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
func (r MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionNot) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionNot) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionDate struct {
	// Any of "DATE".
	Operator     string  `json:"operator,required"`
	Inputs       []any   `json:"inputs"`
	PropertyName string  `json:"propertyName"`
	Value        float64 `json:"value"`
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
func (r MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionDate) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionDate) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionMonth struct {
	// Any of "MONTH".
	Operator     string  `json:"operator,required"`
	Inputs       []any   `json:"inputs"`
	PropertyName string  `json:"propertyName"`
	Value        float64 `json:"value"`
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
func (r MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionMonth) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionMonth) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionYear struct {
	// Any of "YEAR".
	Operator     string  `json:"operator,required"`
	Inputs       []any   `json:"inputs"`
	PropertyName string  `json:"propertyName"`
	Value        float64 `json:"value"`
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
func (r MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionYear) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionYear) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionNow struct {
	// Any of "NOW".
	Operator     string  `json:"operator,required"`
	Inputs       []any   `json:"inputs"`
	PropertyName string  `json:"propertyName"`
	Value        float64 `json:"value"`
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
func (r MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionNow) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionNow) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionTimeBetween struct {
	// Any of "TIME_BETWEEN".
	Operator     string  `json:"operator,required"`
	Inputs       []any   `json:"inputs"`
	PropertyName string  `json:"propertyName"`
	Value        float64 `json:"value"`
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
func (r MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionTimeBetween) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionTimeBetween) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionPeriodToMonths struct {
	// Any of "PERIOD_TO_MONTHS".
	Operator     string  `json:"operator,required"`
	Inputs       []any   `json:"inputs"`
	PropertyName string  `json:"propertyName"`
	Value        float64 `json:"value"`
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
func (r MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionPeriodToMonths) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionPeriodToMonths) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionPeriodToWeeks struct {
	// Any of "PERIOD_TO_WEEKS".
	Operator     string  `json:"operator,required"`
	Inputs       []any   `json:"inputs"`
	PropertyName string  `json:"propertyName"`
	Value        float64 `json:"value"`
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
func (r MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionPeriodToWeeks) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionPeriodToWeeks) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionAnd struct {
	EnclosedInParentheses bool `json:"enclosedInParentheses,required"`
	// Any of "AND".
	Operator     string `json:"operator,required"`
	Inputs       []any  `json:"inputs"`
	PropertyName string `json:"propertyName"`
	Value        bool   `json:"value"`
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
func (r MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionAnd) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionAnd) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionOr struct {
	EnclosedInParentheses bool `json:"enclosedInParentheses,required"`
	// Any of "OR".
	Operator     string `json:"operator,required"`
	Inputs       []any  `json:"inputs"`
	PropertyName string `json:"propertyName"`
	Value        bool   `json:"value"`
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
func (r MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionOr) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionOr) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionXor struct {
	EnclosedInParentheses bool `json:"enclosedInParentheses,required"`
	// Any of "XOR".
	Operator     string `json:"operator,required"`
	Inputs       []any  `json:"inputs"`
	PropertyName string `json:"propertyName"`
	Value        bool   `json:"value"`
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
func (r MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionXor) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionXor) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionIfString struct {
	EnclosedInParentheses bool `json:"enclosedInParentheses,required"`
	IfExpression          any  `json:"ifExpression,required"`
	// Any of "IF_STRING".
	Operator       string `json:"operator,required"`
	ElseExpression any    `json:"elseExpression"`
	Inputs         []any  `json:"inputs"`
	PropertyName   string `json:"propertyName"`
	Value          string `json:"value"`
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
func (r MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionIfString) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionIfString) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionIfNumber struct {
	EnclosedInParentheses bool `json:"enclosedInParentheses,required"`
	IfExpression          any  `json:"ifExpression,required"`
	// Any of "IF_NUMBER".
	Operator       string  `json:"operator,required"`
	ElseExpression any     `json:"elseExpression"`
	Inputs         []any   `json:"inputs"`
	PropertyName   string  `json:"propertyName"`
	Value          float64 `json:"value"`
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
func (r MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionIfNumber) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionIfNumber) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionIfBoolean struct {
	EnclosedInParentheses bool `json:"enclosedInParentheses,required"`
	IfExpression          any  `json:"ifExpression,required"`
	// Any of "IF_BOOLEAN".
	Operator       string `json:"operator,required"`
	ElseExpression any    `json:"elseExpression"`
	Inputs         []any  `json:"inputs"`
	PropertyName   string `json:"propertyName"`
	Value          bool   `json:"value"`
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
func (r MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionIfBoolean) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionIfBoolean) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionIsPresent struct {
	ExpressionToEvaluate any `json:"expressionToEvaluate,required"`
	// Any of "IS_PRESENT".
	Operator     string `json:"operator,required"`
	Inputs       []any  `json:"inputs"`
	PropertyName string `json:"propertyName"`
	Value        bool   `json:"value"`
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
func (r MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionIsPresent) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionIsPresent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionHasEmailReply struct {
	// Any of "HAS_EMAIL_REPLY".
	Operator     string `json:"operator,required"`
	Inputs       []any  `json:"inputs"`
	PropertyName string `json:"propertyName"`
	Value        bool   `json:"value"`
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
func (r MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionHasEmailReply) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionHasEmailReply) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionHasPlainTextEmailReply struct {
	// Any of "HAS_PLAIN_TEXT_EMAIL_REPLY".
	Operator     string `json:"operator,required"`
	Inputs       []any  `json:"inputs"`
	PropertyName string `json:"propertyName"`
	Value        bool   `json:"value"`
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
func (r MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionHasPlainTextEmailReply) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionHasPlainTextEmailReply) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionExtractMostRecentEmailReplyHTML struct {
	// Any of "EXTRACT_MOST_RECENT_EMAIL_REPLY_HTML".
	Operator     string `json:"operator,required"`
	Inputs       []any  `json:"inputs"`
	PropertyName string `json:"propertyName"`
	Value        string `json:"value"`
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
func (r MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionExtractMostRecentEmailReplyHTML) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionExtractMostRecentEmailReplyHTML) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionExtractMostRecentEmailReplyText struct {
	// Any of "EXTRACT_MOST_RECENT_EMAIL_REPLY_TEXT".
	Operator     string `json:"operator,required"`
	Inputs       []any  `json:"inputs"`
	PropertyName string `json:"propertyName"`
	Value        string `json:"value"`
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
func (r MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionExtractMostRecentEmailReplyText) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionExtractMostRecentEmailReplyText) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionExtractMostRecentPlainTextEmailReply struct {
	// Any of "EXTRACT_MOST_RECENT_PLAIN_TEXT_EMAIL_REPLY".
	Operator     string `json:"operator,required"`
	Inputs       []any  `json:"inputs"`
	PropertyName string `json:"propertyName"`
	Value        string `json:"value"`
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
func (r MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionExtractMostRecentPlainTextEmailReply) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionExtractMostRecentPlainTextEmailReply) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionSetContainsString struct {
	// Any of "SET_CONTAINS_STRING".
	Operator      string `json:"operator,required"`
	StringToCheck any    `json:"stringToCheck,required"`
	Inputs        []any  `json:"inputs"`
	PropertyName  string `json:"propertyName"`
	Value         bool   `json:"value"`
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
func (r MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionSetContainsString) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionSetContainsString) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionIsEngagementType struct {
	// Any of "IS_ENGAGEMENT_TYPE".
	Operator     string `json:"operator,required"`
	Inputs       []any  `json:"inputs"`
	PropertyName string `json:"propertyName"`
	Value        bool   `json:"value"`
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
func (r MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionIsEngagementType) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionIsEngagementType) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionFormatFullName struct {
	// Any of "FORMAT_FULL_NAME".
	Operator     string `json:"operator,required"`
	Inputs       []any  `json:"inputs"`
	PropertyName string `json:"propertyName"`
	Value        string `json:"value"`
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
func (r MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionFormatFullName) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionFormatFullName) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionAbsoluteValue struct {
	// Any of "ABSOLUTE_VALUE".
	Operator     string  `json:"operator,required"`
	Inputs       []any   `json:"inputs"`
	PropertyName string  `json:"propertyName"`
	Value        float64 `json:"value"`
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
func (r MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionAbsoluteValue) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionAbsoluteValue) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionSquareRoot struct {
	// Any of "SQUARE_ROOT".
	Operator     string  `json:"operator,required"`
	Inputs       []any   `json:"inputs"`
	PropertyName string  `json:"propertyName"`
	Value        float64 `json:"value"`
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
func (r MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionSquareRoot) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionSquareRoot) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionPower struct {
	// Any of "POWER".
	Operator     string  `json:"operator,required"`
	Inputs       []any   `json:"inputs"`
	PropertyName string  `json:"propertyName"`
	Value        float64 `json:"value"`
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
func (r MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionPower) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionPower) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionSubstring struct {
	// Any of "SUBSTRING".
	Operator      string `json:"operator,required"`
	StringToCheck any    `json:"stringToCheck,required"`
	Inputs        []any  `json:"inputs"`
	PropertyName  string `json:"propertyName"`
	Value         string `json:"value"`
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
func (r MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionSubstring) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionSubstring) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionEuler struct {
	// Any of "EULER".
	Operator     string  `json:"operator,required"`
	Inputs       []any   `json:"inputs"`
	PropertyName string  `json:"propertyName"`
	Value        float64 `json:"value"`
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
func (r MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionEuler) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionEuler) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionStringLength struct {
	// Any of "STRING_LENGTH".
	Operator     string  `json:"operator,required"`
	Inputs       []any   `json:"inputs"`
	PropertyName string  `json:"propertyName"`
	Value        float64 `json:"value"`
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
func (r MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionStringLength) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionStringLength) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionAddTime struct {
	// Any of "ADD_TIME".
	Operator      string  `json:"operator,required"`
	StringToCheck any     `json:"stringToCheck,required"`
	Inputs        []any   `json:"inputs"`
	PropertyName  string  `json:"propertyName"`
	Value         float64 `json:"value"`
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
func (r MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionAddTime) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionAddTime) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionSubtractTime struct {
	// Any of "SUBTRACT_TIME".
	Operator      string  `json:"operator,required"`
	StringToCheck any     `json:"stringToCheck,required"`
	Inputs        []any   `json:"inputs"`
	PropertyName  string  `json:"propertyName"`
	Value         float64 `json:"value"`
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
func (r MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionSubtractTime) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyRollupExpressionConditionalExpressionSubtractTime) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyGroup struct {
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
func (r MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyGroup) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponsePropertyGroup) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponseSchema struct {
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
	// Any of "HUBSPOT", "INTEGRATION", "PORTAL_SPECIFIC", "CMS_HUBDB",
	// "HUBSPOT_EVENT", "INTEGRATION_EVENT", "PORTAL_SPECIFIC_EVENT".
	MetaType                           string                                                                                  `json:"metaType,required"`
	MetaTypeID                         int64                                                                                   `json:"metaTypeId,required"`
	Name                               string                                                                                  `json:"name,required"`
	ObjectTypeID                       string                                                                                  `json:"objectTypeId,required"`
	PermissioningType                  string                                                                                  `json:"permissioningType,required"`
	PipelinePropertyName               string                                                                                  `json:"pipelinePropertyName,required"`
	PipelineStagePropertyName          string                                                                                  `json:"pipelineStagePropertyName,required"`
	RequiredProperties                 []string                                                                                `json:"requiredProperties,required"`
	Restorable                         bool                                                                                    `json:"restorable,required"`
	ScopeMappings                      []MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponseSchemaScopeMapping `json:"scopeMappings,required"`
	SecondaryDisplayLabelPropertyNames []string                                                                                `json:"secondaryDisplayLabelPropertyNames,required"`
	AccessScopeName                    string                                                                                  `json:"accessScopeName"`
	CreatedAt                          int64                                                                                   `json:"createdAt"`
	Description                        string                                                                                  `json:"description"`
	IntegrationAppID                   int64                                                                                   `json:"integrationAppId"`
	JanusGroup                         string                                                                                  `json:"janusGroup"`
	OwnerPortalID                      int64                                                                                   `json:"ownerPortalId"`
	PipelineCloseDatePropertyName      string                                                                                  `json:"pipelineCloseDatePropertyName"`
	PipelineTimeToClosePropertyName    string                                                                                  `json:"pipelineTimeToClosePropertyName"`
	PluralForm                         string                                                                                  `json:"pluralForm"`
	PrimaryDisplayLabelPropertyName    string                                                                                  `json:"primaryDisplayLabelPropertyName"`
	ReadScopeName                      string                                                                                  `json:"readScopeName"`
	SingularForm                       string                                                                                  `json:"singularForm"`
	Status                             string                                                                                  `json:"status"`
	Visibility                         string                                                                                  `json:"visibility"`
	WriteScopeName                     string                                                                                  `json:"writeScopeName"`
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
func (r MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponseSchema) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponseSchema) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponseSchemaScopeMapping struct {
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
func (r MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponseSchemaScopeMapping) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeResponseSchemaScopeMapping) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingGetOembedDomainResponse struct {
	ID        int64                                                        `json:"id,required"`
	AppID     int64                                                        `json:"appId,required"`
	CreatedAt int64                                                        `json:"createdAt,required"`
	DeletedAt int64                                                        `json:"deletedAt,required"`
	Endpoints MediaBridgeIntegratorSettingGetOembedDomainResponseEndpoints `json:"endpoints,required"`
	PortalID  int64                                                        `json:"portalId,required"`
	UpdatedAt int64                                                        `json:"updatedAt,required"`
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
func (r MediaBridgeIntegratorSettingGetOembedDomainResponse) RawJSON() string { return r.JSON.raw }
func (r *MediaBridgeIntegratorSettingGetOembedDomainResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingGetOembedDomainResponseEndpoints struct {
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
func (r MediaBridgeIntegratorSettingGetOembedDomainResponseEndpoints) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingGetOembedDomainResponseEndpoints) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingListOembedDomainsResponse struct {
	Results    []MediaBridgeIntegratorSettingListOembedDomainsResponseResult `json:"results,required"`
	TotalCount int64                                                         `json:"totalCount"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Results     respjson.Field
		TotalCount  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MediaBridgeIntegratorSettingListOembedDomainsResponse) RawJSON() string { return r.JSON.raw }
func (r *MediaBridgeIntegratorSettingListOembedDomainsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingListOembedDomainsResponseResult struct {
	ID        int64                                                                `json:"id,required"`
	AppID     int64                                                                `json:"appId,required"`
	CreatedAt int64                                                                `json:"createdAt,required"`
	DeletedAt int64                                                                `json:"deletedAt,required"`
	Endpoints MediaBridgeIntegratorSettingListOembedDomainsResponseResultEndpoints `json:"endpoints,required"`
	PortalID  int64                                                                `json:"portalId,required"`
	UpdatedAt int64                                                                `json:"updatedAt,required"`
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
func (r MediaBridgeIntegratorSettingListOembedDomainsResponseResult) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingListOembedDomainsResponseResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingListOembedDomainsResponseResultEndpoints struct {
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
func (r MediaBridgeIntegratorSettingListOembedDomainsResponseResultEndpoints) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingListOembedDomainsResponseResultEndpoints) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingRegisterAppNameResponse struct {
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
func (r MediaBridgeIntegratorSettingRegisterAppNameResponse) RawJSON() string { return r.JSON.raw }
func (r *MediaBridgeIntegratorSettingRegisterAppNameResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingUpdateAppNameResponse struct {
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
func (r MediaBridgeIntegratorSettingUpdateAppNameResponse) RawJSON() string { return r.JSON.raw }
func (r *MediaBridgeIntegratorSettingUpdateAppNameResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingUpdateEventVisibilitySettingsResponse struct {
	// Any of "ALL", "MEDIA_PLAYS", "MEDIA_PLAYS_PERCENT", "ATTENTION_SPAN".
	EventType       MediaBridgeIntegratorSettingUpdateEventVisibilitySettingsResponseEventType `json:"eventType,required"`
	UpdatedAt       int64                                                                      `json:"updatedAt,required"`
	ShowInReporting bool                                                                       `json:"showInReporting"`
	ShowInTimeline  bool                                                                       `json:"showInTimeline"`
	ShowInWorkflows bool                                                                       `json:"showInWorkflows"`
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
func (r MediaBridgeIntegratorSettingUpdateEventVisibilitySettingsResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingUpdateEventVisibilitySettingsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingUpdateEventVisibilitySettingsResponseEventType string

const (
	MediaBridgeIntegratorSettingUpdateEventVisibilitySettingsResponseEventTypeAll               MediaBridgeIntegratorSettingUpdateEventVisibilitySettingsResponseEventType = "ALL"
	MediaBridgeIntegratorSettingUpdateEventVisibilitySettingsResponseEventTypeMediaPlays        MediaBridgeIntegratorSettingUpdateEventVisibilitySettingsResponseEventType = "MEDIA_PLAYS"
	MediaBridgeIntegratorSettingUpdateEventVisibilitySettingsResponseEventTypeMediaPlaysPercent MediaBridgeIntegratorSettingUpdateEventVisibilitySettingsResponseEventType = "MEDIA_PLAYS_PERCENT"
	MediaBridgeIntegratorSettingUpdateEventVisibilitySettingsResponseEventTypeAttentionSpan     MediaBridgeIntegratorSettingUpdateEventVisibilitySettingsResponseEventType = "ATTENTION_SPAN"
)

type MediaBridgeIntegratorSettingUpdateOembedDomainResponse struct {
	ID        int64                                                           `json:"id,required"`
	AppID     int64                                                           `json:"appId,required"`
	CreatedAt int64                                                           `json:"createdAt,required"`
	DeletedAt int64                                                           `json:"deletedAt,required"`
	Endpoints MediaBridgeIntegratorSettingUpdateOembedDomainResponseEndpoints `json:"endpoints,required"`
	PortalID  int64                                                           `json:"portalId,required"`
	UpdatedAt int64                                                           `json:"updatedAt,required"`
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
func (r MediaBridgeIntegratorSettingUpdateOembedDomainResponse) RawJSON() string { return r.JSON.raw }
func (r *MediaBridgeIntegratorSettingUpdateOembedDomainResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingUpdateOembedDomainResponseEndpoints struct {
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
func (r MediaBridgeIntegratorSettingUpdateOembedDomainResponseEndpoints) RawJSON() string {
	return r.JSON.raw
}
func (r *MediaBridgeIntegratorSettingUpdateOembedDomainResponseEndpoints) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingNewObjectDefinitionParams struct {
	// Any of "VIDEO", "AUDIO", "DOCUMENT", "OTHER", "IMAGE".
	MediaTypes []string `json:"mediaTypes,omitzero,required"`
	paramObj
}

func (r MediaBridgeIntegratorSettingNewObjectDefinitionParams) MarshalJSON() (data []byte, err error) {
	type shadow MediaBridgeIntegratorSettingNewObjectDefinitionParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MediaBridgeIntegratorSettingNewObjectDefinitionParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingNewOembedDomainParams struct {
	Endpoints MediaBridgeIntegratorSettingNewOembedDomainParamsEndpoints `json:"endpoints,omitzero,required"`
	PortalID  param.Opt[int64]                                           `json:"portalId,omitzero"`
	paramObj
}

func (r MediaBridgeIntegratorSettingNewOembedDomainParams) MarshalJSON() (data []byte, err error) {
	type shadow MediaBridgeIntegratorSettingNewOembedDomainParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MediaBridgeIntegratorSettingNewOembedDomainParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Discovery, Schemes, URL are required.
type MediaBridgeIntegratorSettingNewOembedDomainParamsEndpoints struct {
	Discovery bool     `json:"discovery,required"`
	Schemes   []string `json:"schemes,omitzero,required"`
	URL       string   `json:"url,required"`
	paramObj
}

func (r MediaBridgeIntegratorSettingNewOembedDomainParamsEndpoints) MarshalJSON() (data []byte, err error) {
	type shadow MediaBridgeIntegratorSettingNewOembedDomainParamsEndpoints
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MediaBridgeIntegratorSettingNewOembedDomainParamsEndpoints) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeParams struct {
	AppID string `path:"appId,required" json:"-"`
	paramObj
}

type MediaBridgeIntegratorSettingGetOembedDomainParams struct {
	AppID string `path:"appId,required" json:"-"`
	paramObj
}

type MediaBridgeIntegratorSettingRegisterAppNameParams struct {
	UpdatedAt int64             `json:"updatedAt,required"`
	Name      param.Opt[string] `json:"name,omitzero"`
	paramObj
}

func (r MediaBridgeIntegratorSettingRegisterAppNameParams) MarshalJSON() (data []byte, err error) {
	type shadow MediaBridgeIntegratorSettingRegisterAppNameParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MediaBridgeIntegratorSettingRegisterAppNameParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingUpdateAppNameParams struct {
	UpdatedAt int64             `json:"updatedAt,required"`
	Name      param.Opt[string] `json:"name,omitzero"`
	paramObj
}

func (r MediaBridgeIntegratorSettingUpdateAppNameParams) MarshalJSON() (data []byte, err error) {
	type shadow MediaBridgeIntegratorSettingUpdateAppNameParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MediaBridgeIntegratorSettingUpdateAppNameParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingUpdateEventVisibilitySettingsParams struct {
	// Any of "ALL", "MEDIA_PLAYS", "MEDIA_PLAYS_PERCENT", "ATTENTION_SPAN".
	EventType       MediaBridgeIntegratorSettingUpdateEventVisibilitySettingsParamsEventType `json:"eventType,omitzero,required"`
	UpdatedAt       int64                                                                    `json:"updatedAt,required"`
	ShowInReporting param.Opt[bool]                                                          `json:"showInReporting,omitzero"`
	ShowInTimeline  param.Opt[bool]                                                          `json:"showInTimeline,omitzero"`
	ShowInWorkflows param.Opt[bool]                                                          `json:"showInWorkflows,omitzero"`
	paramObj
}

func (r MediaBridgeIntegratorSettingUpdateEventVisibilitySettingsParams) MarshalJSON() (data []byte, err error) {
	type shadow MediaBridgeIntegratorSettingUpdateEventVisibilitySettingsParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MediaBridgeIntegratorSettingUpdateEventVisibilitySettingsParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgeIntegratorSettingUpdateEventVisibilitySettingsParamsEventType string

const (
	MediaBridgeIntegratorSettingUpdateEventVisibilitySettingsParamsEventTypeAll               MediaBridgeIntegratorSettingUpdateEventVisibilitySettingsParamsEventType = "ALL"
	MediaBridgeIntegratorSettingUpdateEventVisibilitySettingsParamsEventTypeMediaPlays        MediaBridgeIntegratorSettingUpdateEventVisibilitySettingsParamsEventType = "MEDIA_PLAYS"
	MediaBridgeIntegratorSettingUpdateEventVisibilitySettingsParamsEventTypeMediaPlaysPercent MediaBridgeIntegratorSettingUpdateEventVisibilitySettingsParamsEventType = "MEDIA_PLAYS_PERCENT"
	MediaBridgeIntegratorSettingUpdateEventVisibilitySettingsParamsEventTypeAttentionSpan     MediaBridgeIntegratorSettingUpdateEventVisibilitySettingsParamsEventType = "ATTENTION_SPAN"
)

type MediaBridgeIntegratorSettingUpdateOembedDomainParams struct {
	AppID     string                                                        `path:"appId,required" json:"-"`
	Endpoints MediaBridgeIntegratorSettingUpdateOembedDomainParamsEndpoints `json:"endpoints,omitzero,required"`
	PortalID  param.Opt[int64]                                              `json:"portalId,omitzero"`
	paramObj
}

func (r MediaBridgeIntegratorSettingUpdateOembedDomainParams) MarshalJSON() (data []byte, err error) {
	type shadow MediaBridgeIntegratorSettingUpdateOembedDomainParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MediaBridgeIntegratorSettingUpdateOembedDomainParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Discovery, Schemes, URL are required.
type MediaBridgeIntegratorSettingUpdateOembedDomainParamsEndpoints struct {
	Discovery bool     `json:"discovery,required"`
	Schemes   []string `json:"schemes,omitzero,required"`
	URL       string   `json:"url,required"`
	paramObj
}

func (r MediaBridgeIntegratorSettingUpdateOembedDomainParamsEndpoints) MarshalJSON() (data []byte, err error) {
	type shadow MediaBridgeIntegratorSettingUpdateOembedDomainParamsEndpoints
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MediaBridgeIntegratorSettingUpdateOembedDomainParamsEndpoints) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
