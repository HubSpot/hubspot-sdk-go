// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package crm

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/stainless-sdks/hubspot-sdk-go/internal/apijson"
	shimjson "github.com/stainless-sdks/hubspot-sdk-go/internal/encoding/json"
	"github.com/stainless-sdks/hubspot-sdk-go/internal/requestconfig"
	"github.com/stainless-sdks/hubspot-sdk-go/option"
	"github.com/stainless-sdks/hubspot-sdk-go/packages/param"
	"github.com/stainless-sdks/hubspot-sdk-go/packages/respjson"
)

// PropertyValidationService contains methods and other services that help with
// interacting with the hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPropertyValidationService] method instead.
type PropertyValidationService struct {
	Options []option.RequestOption
}

// NewPropertyValidationService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewPropertyValidationService(opts ...option.RequestOption) (r PropertyValidationService) {
	r = PropertyValidationService{}
	r.Options = opts
	return
}

// Read all properties with validation rules for a given object.
func (r *PropertyValidationService) List(ctx context.Context, objectTypeID string, opts ...option.RequestOption) (res *CollectionResponsePublicPropertyValidationRuleMapNoPaging, err error) {
	opts = slices.Concat(r.Options, opts)
	if objectTypeID == "" {
		err = errors.New("missing required objectTypeId parameter")
		return
	}
	path := fmt.Sprintf("crm/v3/property-validations/%s", objectTypeID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update a specific validation rule for a property identified by its name and rule
// type.
func (r *PropertyValidationService) CrmV3PropertyValidationsObjectTypeIDPropertyNameRuleTypeRuleType(ctx context.Context, ruleType PropertyValidationCrmV3PropertyValidationsObjectTypeIDPropertyNameRuleTypeRuleTypeParamsRuleType, params PropertyValidationCrmV3PropertyValidationsObjectTypeIDPropertyNameRuleTypeRuleTypeParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if params.ObjectTypeID == "" {
		err = errors.New("missing required objectTypeId parameter")
		return
	}
	if params.PropertyName == "" {
		err = errors.New("missing required propertyName parameter")
		return
	}
	path := fmt.Sprintf("crm/v3/property-validations/%s/%s/rule-type/%v", params.ObjectTypeID, params.PropertyName, ruleType)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, nil, opts...)
	return
}

// Read a property's validation rules identified by {propertyName}.
func (r *PropertyValidationService) Get(ctx context.Context, propertyName string, query PropertyValidationGetParams, opts ...option.RequestOption) (res *CollectionResponsePublicPropertyValidationRuleNoPaging, err error) {
	opts = slices.Concat(r.Options, opts)
	if query.ObjectTypeID == "" {
		err = errors.New("missing required objectTypeId parameter")
		return
	}
	if propertyName == "" {
		err = errors.New("missing required propertyName parameter")
		return
	}
	path := fmt.Sprintf("crm/v3/property-validations/%s/%s", query.ObjectTypeID, propertyName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type CollectionResponsePublicPropertyValidationRuleMapNoPaging struct {
	// Collection of properties with their validation rules. Each item maps a property
	// name to its configured validation rules for the specified object type.
	Results []PublicPropertyValidationRuleMap `json:"results,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Results     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CollectionResponsePublicPropertyValidationRuleMapNoPaging) RawJSON() string {
	return r.JSON.raw
}
func (r *CollectionResponsePublicPropertyValidationRuleMapNoPaging) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CollectionResponsePublicPropertyValidationRuleNoPaging struct {
	// Collection of validation rules configured for the specified property. Each rule
	// defines a constraint that property values must satisfy (e.g., format
	// requirements, length limits, allowed values).
	Results []PublicPropertyValidationRule `json:"results,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Results     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CollectionResponsePublicPropertyValidationRuleNoPaging) RawJSON() string { return r.JSON.raw }
func (r *CollectionResponsePublicPropertyValidationRuleNoPaging) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PublicPropertyValidationRule struct {
	// A list of arguments that define the specific conditions or parameters for the
	// validation rule.
	RuleArguments []string `json:"ruleArguments,required"`
	// The category of validation applied to the property, such as FORMAT,
	// ALPHANUMERIC, or MAX_LENGTH.
	//
	// Any of "AFTER_DATETIME_DURATION", "AFTER_DURATION", "ALPHANUMERIC",
	// "BEFORE_DATETIME_DURATION", "BEFORE_DURATION", "DAYS_OF_WEEK", "DECIMAL",
	// "DOMAIN", "EMAIL", "EMAIL_ALLOWED_DOMAINS", "EMAIL_BLOCKED_DOMAINS", "END_DATE",
	// "END_DATETIME", "FORMAT", "MAX_LENGTH", "MAX_NUMBER", "MIN_LENGTH",
	// "MIN_NUMBER", "PHONE_NUMBER_WITH_EXPLICIT_COUNTRY_CODE", "REGEX",
	// "SPECIAL_CHARACTERS", "START_DATE", "START_DATETIME", "URL",
	// "URL_ALLOWED_DOMAINS", "URL_BLOCKED_DOMAINS", "WHITESPACE".
	RuleType PublicPropertyValidationRuleRuleType `json:"ruleType,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		RuleArguments respjson.Field
		RuleType      respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PublicPropertyValidationRule) RawJSON() string { return r.JSON.raw }
func (r *PublicPropertyValidationRule) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The category of validation applied to the property, such as FORMAT,
// ALPHANUMERIC, or MAX_LENGTH.
type PublicPropertyValidationRuleRuleType string

const (
	PublicPropertyValidationRuleRuleTypeAfterDatetimeDuration              PublicPropertyValidationRuleRuleType = "AFTER_DATETIME_DURATION"
	PublicPropertyValidationRuleRuleTypeAfterDuration                      PublicPropertyValidationRuleRuleType = "AFTER_DURATION"
	PublicPropertyValidationRuleRuleTypeAlphanumeric                       PublicPropertyValidationRuleRuleType = "ALPHANUMERIC"
	PublicPropertyValidationRuleRuleTypeBeforeDatetimeDuration             PublicPropertyValidationRuleRuleType = "BEFORE_DATETIME_DURATION"
	PublicPropertyValidationRuleRuleTypeBeforeDuration                     PublicPropertyValidationRuleRuleType = "BEFORE_DURATION"
	PublicPropertyValidationRuleRuleTypeDaysOfWeek                         PublicPropertyValidationRuleRuleType = "DAYS_OF_WEEK"
	PublicPropertyValidationRuleRuleTypeDecimal                            PublicPropertyValidationRuleRuleType = "DECIMAL"
	PublicPropertyValidationRuleRuleTypeDomain                             PublicPropertyValidationRuleRuleType = "DOMAIN"
	PublicPropertyValidationRuleRuleTypeEmail                              PublicPropertyValidationRuleRuleType = "EMAIL"
	PublicPropertyValidationRuleRuleTypeEmailAllowedDomains                PublicPropertyValidationRuleRuleType = "EMAIL_ALLOWED_DOMAINS"
	PublicPropertyValidationRuleRuleTypeEmailBlockedDomains                PublicPropertyValidationRuleRuleType = "EMAIL_BLOCKED_DOMAINS"
	PublicPropertyValidationRuleRuleTypeEndDate                            PublicPropertyValidationRuleRuleType = "END_DATE"
	PublicPropertyValidationRuleRuleTypeEndDatetime                        PublicPropertyValidationRuleRuleType = "END_DATETIME"
	PublicPropertyValidationRuleRuleTypeFormat                             PublicPropertyValidationRuleRuleType = "FORMAT"
	PublicPropertyValidationRuleRuleTypeMaxLength                          PublicPropertyValidationRuleRuleType = "MAX_LENGTH"
	PublicPropertyValidationRuleRuleTypeMaxNumber                          PublicPropertyValidationRuleRuleType = "MAX_NUMBER"
	PublicPropertyValidationRuleRuleTypeMinLength                          PublicPropertyValidationRuleRuleType = "MIN_LENGTH"
	PublicPropertyValidationRuleRuleTypeMinNumber                          PublicPropertyValidationRuleRuleType = "MIN_NUMBER"
	PublicPropertyValidationRuleRuleTypePhoneNumberWithExplicitCountryCode PublicPropertyValidationRuleRuleType = "PHONE_NUMBER_WITH_EXPLICIT_COUNTRY_CODE"
	PublicPropertyValidationRuleRuleTypeRegex                              PublicPropertyValidationRuleRuleType = "REGEX"
	PublicPropertyValidationRuleRuleTypeSpecialCharacters                  PublicPropertyValidationRuleRuleType = "SPECIAL_CHARACTERS"
	PublicPropertyValidationRuleRuleTypeStartDate                          PublicPropertyValidationRuleRuleType = "START_DATE"
	PublicPropertyValidationRuleRuleTypeStartDatetime                      PublicPropertyValidationRuleRuleType = "START_DATETIME"
	PublicPropertyValidationRuleRuleTypeURL                                PublicPropertyValidationRuleRuleType = "URL"
	PublicPropertyValidationRuleRuleTypeURLAllowedDomains                  PublicPropertyValidationRuleRuleType = "URL_ALLOWED_DOMAINS"
	PublicPropertyValidationRuleRuleTypeURLBlockedDomains                  PublicPropertyValidationRuleRuleType = "URL_BLOCKED_DOMAINS"
	PublicPropertyValidationRuleRuleTypeWhitespace                         PublicPropertyValidationRuleRuleType = "WHITESPACE"
)

type PublicPropertyValidationRuleMap struct {
	// The name of the property for which validation rules are defined.
	PropertyName string `json:"propertyName,required"`
	// A list of validation rules applicable to the property.
	PropertyValidationRules []PublicPropertyValidationRule `json:"propertyValidationRules,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PropertyName            respjson.Field
		PropertyValidationRules respjson.Field
		ExtraFields             map[string]respjson.Field
		raw                     string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PublicPropertyValidationRuleMap) RawJSON() string { return r.JSON.raw }
func (r *PublicPropertyValidationRuleMap) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property RuleArguments is required.
type PublicPropertyValidationRuleUpdateParam struct {
	// A list of arguments that define the constraints for the validation rule.
	RuleArguments []string `json:"ruleArguments,omitzero,required"`
	paramObj
}

func (r PublicPropertyValidationRuleUpdateParam) MarshalJSON() (data []byte, err error) {
	type shadow PublicPropertyValidationRuleUpdateParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PublicPropertyValidationRuleUpdateParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PropertyValidationCrmV3PropertyValidationsObjectTypeIDPropertyNameRuleTypeRuleTypeParams struct {
	ObjectTypeID                       string `path:"objectTypeId,required" json:"-"`
	PropertyName                       string `path:"propertyName,required" json:"-"`
	PublicPropertyValidationRuleUpdate PublicPropertyValidationRuleUpdateParam
	paramObj
}

func (r PropertyValidationCrmV3PropertyValidationsObjectTypeIDPropertyNameRuleTypeRuleTypeParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.PublicPropertyValidationRuleUpdate)
}
func (r *PropertyValidationCrmV3PropertyValidationsObjectTypeIDPropertyNameRuleTypeRuleTypeParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.PublicPropertyValidationRuleUpdate)
}

type PropertyValidationCrmV3PropertyValidationsObjectTypeIDPropertyNameRuleTypeRuleTypeParamsRuleType string

const (
	PropertyValidationCrmV3PropertyValidationsObjectTypeIDPropertyNameRuleTypeRuleTypeParamsRuleTypeAfterDatetimeDuration              PropertyValidationCrmV3PropertyValidationsObjectTypeIDPropertyNameRuleTypeRuleTypeParamsRuleType = "AFTER_DATETIME_DURATION"
	PropertyValidationCrmV3PropertyValidationsObjectTypeIDPropertyNameRuleTypeRuleTypeParamsRuleTypeAfterDuration                      PropertyValidationCrmV3PropertyValidationsObjectTypeIDPropertyNameRuleTypeRuleTypeParamsRuleType = "AFTER_DURATION"
	PropertyValidationCrmV3PropertyValidationsObjectTypeIDPropertyNameRuleTypeRuleTypeParamsRuleTypeAlphanumeric                       PropertyValidationCrmV3PropertyValidationsObjectTypeIDPropertyNameRuleTypeRuleTypeParamsRuleType = "ALPHANUMERIC"
	PropertyValidationCrmV3PropertyValidationsObjectTypeIDPropertyNameRuleTypeRuleTypeParamsRuleTypeBeforeDatetimeDuration             PropertyValidationCrmV3PropertyValidationsObjectTypeIDPropertyNameRuleTypeRuleTypeParamsRuleType = "BEFORE_DATETIME_DURATION"
	PropertyValidationCrmV3PropertyValidationsObjectTypeIDPropertyNameRuleTypeRuleTypeParamsRuleTypeBeforeDuration                     PropertyValidationCrmV3PropertyValidationsObjectTypeIDPropertyNameRuleTypeRuleTypeParamsRuleType = "BEFORE_DURATION"
	PropertyValidationCrmV3PropertyValidationsObjectTypeIDPropertyNameRuleTypeRuleTypeParamsRuleTypeDaysOfWeek                         PropertyValidationCrmV3PropertyValidationsObjectTypeIDPropertyNameRuleTypeRuleTypeParamsRuleType = "DAYS_OF_WEEK"
	PropertyValidationCrmV3PropertyValidationsObjectTypeIDPropertyNameRuleTypeRuleTypeParamsRuleTypeDecimal                            PropertyValidationCrmV3PropertyValidationsObjectTypeIDPropertyNameRuleTypeRuleTypeParamsRuleType = "DECIMAL"
	PropertyValidationCrmV3PropertyValidationsObjectTypeIDPropertyNameRuleTypeRuleTypeParamsRuleTypeDomain                             PropertyValidationCrmV3PropertyValidationsObjectTypeIDPropertyNameRuleTypeRuleTypeParamsRuleType = "DOMAIN"
	PropertyValidationCrmV3PropertyValidationsObjectTypeIDPropertyNameRuleTypeRuleTypeParamsRuleTypeEmail                              PropertyValidationCrmV3PropertyValidationsObjectTypeIDPropertyNameRuleTypeRuleTypeParamsRuleType = "EMAIL"
	PropertyValidationCrmV3PropertyValidationsObjectTypeIDPropertyNameRuleTypeRuleTypeParamsRuleTypeEmailAllowedDomains                PropertyValidationCrmV3PropertyValidationsObjectTypeIDPropertyNameRuleTypeRuleTypeParamsRuleType = "EMAIL_ALLOWED_DOMAINS"
	PropertyValidationCrmV3PropertyValidationsObjectTypeIDPropertyNameRuleTypeRuleTypeParamsRuleTypeEmailBlockedDomains                PropertyValidationCrmV3PropertyValidationsObjectTypeIDPropertyNameRuleTypeRuleTypeParamsRuleType = "EMAIL_BLOCKED_DOMAINS"
	PropertyValidationCrmV3PropertyValidationsObjectTypeIDPropertyNameRuleTypeRuleTypeParamsRuleTypeEndDate                            PropertyValidationCrmV3PropertyValidationsObjectTypeIDPropertyNameRuleTypeRuleTypeParamsRuleType = "END_DATE"
	PropertyValidationCrmV3PropertyValidationsObjectTypeIDPropertyNameRuleTypeRuleTypeParamsRuleTypeEndDatetime                        PropertyValidationCrmV3PropertyValidationsObjectTypeIDPropertyNameRuleTypeRuleTypeParamsRuleType = "END_DATETIME"
	PropertyValidationCrmV3PropertyValidationsObjectTypeIDPropertyNameRuleTypeRuleTypeParamsRuleTypeFormat                             PropertyValidationCrmV3PropertyValidationsObjectTypeIDPropertyNameRuleTypeRuleTypeParamsRuleType = "FORMAT"
	PropertyValidationCrmV3PropertyValidationsObjectTypeIDPropertyNameRuleTypeRuleTypeParamsRuleTypeMaxLength                          PropertyValidationCrmV3PropertyValidationsObjectTypeIDPropertyNameRuleTypeRuleTypeParamsRuleType = "MAX_LENGTH"
	PropertyValidationCrmV3PropertyValidationsObjectTypeIDPropertyNameRuleTypeRuleTypeParamsRuleTypeMaxNumber                          PropertyValidationCrmV3PropertyValidationsObjectTypeIDPropertyNameRuleTypeRuleTypeParamsRuleType = "MAX_NUMBER"
	PropertyValidationCrmV3PropertyValidationsObjectTypeIDPropertyNameRuleTypeRuleTypeParamsRuleTypeMinLength                          PropertyValidationCrmV3PropertyValidationsObjectTypeIDPropertyNameRuleTypeRuleTypeParamsRuleType = "MIN_LENGTH"
	PropertyValidationCrmV3PropertyValidationsObjectTypeIDPropertyNameRuleTypeRuleTypeParamsRuleTypeMinNumber                          PropertyValidationCrmV3PropertyValidationsObjectTypeIDPropertyNameRuleTypeRuleTypeParamsRuleType = "MIN_NUMBER"
	PropertyValidationCrmV3PropertyValidationsObjectTypeIDPropertyNameRuleTypeRuleTypeParamsRuleTypePhoneNumberWithExplicitCountryCode PropertyValidationCrmV3PropertyValidationsObjectTypeIDPropertyNameRuleTypeRuleTypeParamsRuleType = "PHONE_NUMBER_WITH_EXPLICIT_COUNTRY_CODE"
	PropertyValidationCrmV3PropertyValidationsObjectTypeIDPropertyNameRuleTypeRuleTypeParamsRuleTypeRegex                              PropertyValidationCrmV3PropertyValidationsObjectTypeIDPropertyNameRuleTypeRuleTypeParamsRuleType = "REGEX"
	PropertyValidationCrmV3PropertyValidationsObjectTypeIDPropertyNameRuleTypeRuleTypeParamsRuleTypeSpecialCharacters                  PropertyValidationCrmV3PropertyValidationsObjectTypeIDPropertyNameRuleTypeRuleTypeParamsRuleType = "SPECIAL_CHARACTERS"
	PropertyValidationCrmV3PropertyValidationsObjectTypeIDPropertyNameRuleTypeRuleTypeParamsRuleTypeStartDate                          PropertyValidationCrmV3PropertyValidationsObjectTypeIDPropertyNameRuleTypeRuleTypeParamsRuleType = "START_DATE"
	PropertyValidationCrmV3PropertyValidationsObjectTypeIDPropertyNameRuleTypeRuleTypeParamsRuleTypeStartDatetime                      PropertyValidationCrmV3PropertyValidationsObjectTypeIDPropertyNameRuleTypeRuleTypeParamsRuleType = "START_DATETIME"
	PropertyValidationCrmV3PropertyValidationsObjectTypeIDPropertyNameRuleTypeRuleTypeParamsRuleTypeURL                                PropertyValidationCrmV3PropertyValidationsObjectTypeIDPropertyNameRuleTypeRuleTypeParamsRuleType = "URL"
	PropertyValidationCrmV3PropertyValidationsObjectTypeIDPropertyNameRuleTypeRuleTypeParamsRuleTypeURLAllowedDomains                  PropertyValidationCrmV3PropertyValidationsObjectTypeIDPropertyNameRuleTypeRuleTypeParamsRuleType = "URL_ALLOWED_DOMAINS"
	PropertyValidationCrmV3PropertyValidationsObjectTypeIDPropertyNameRuleTypeRuleTypeParamsRuleTypeURLBlockedDomains                  PropertyValidationCrmV3PropertyValidationsObjectTypeIDPropertyNameRuleTypeRuleTypeParamsRuleType = "URL_BLOCKED_DOMAINS"
	PropertyValidationCrmV3PropertyValidationsObjectTypeIDPropertyNameRuleTypeRuleTypeParamsRuleTypeWhitespace                         PropertyValidationCrmV3PropertyValidationsObjectTypeIDPropertyNameRuleTypeRuleTypeParamsRuleType = "WHITESPACE"
)

type PropertyValidationGetParams struct {
	ObjectTypeID string `path:"objectTypeId,required" json:"-"`
	paramObj
}
