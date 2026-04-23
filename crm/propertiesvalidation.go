// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package crm

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/stainless-sdks/hubspot-sdk-go/internal/apijson"
	shimjson "github.com/stainless-sdks/hubspot-sdk-go/internal/encoding/json"
	"github.com/stainless-sdks/hubspot-sdk-go/internal/requestconfig"
	"github.com/stainless-sdks/hubspot-sdk-go/option"
	"github.com/stainless-sdks/hubspot-sdk-go/packages/param"
	"github.com/stainless-sdks/hubspot-sdk-go/packages/respjson"
)

// PropertiesValidationService contains methods and other services that help with
// interacting with the hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPropertiesValidationService] method instead.
type PropertiesValidationService struct {
	options []option.RequestOption
}

// NewPropertiesValidationService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewPropertiesValidationService(opts ...option.RequestOption) (r PropertiesValidationService) {
	r = PropertiesValidationService{}
	r.options = opts
	return
}

// Read all properties with validation rules for a given object.
func (r *PropertiesValidationService) GetByObjectTypeID(ctx context.Context, objectTypeID string, opts ...option.RequestOption) (res *CollectionResponsePublicPropertyValidationRuleMapNoPaging, err error) {
	opts = slices.Concat(r.options, opts)
	if objectTypeID == "" {
		err = errors.New("missing required objectTypeId parameter")
		return nil, err
	}
	path := fmt.Sprintf("crm/property-validations/2026-03/%s", url.PathEscape(objectTypeID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Read a property's validation rules identified by {propertyName}.
func (r *PropertiesValidationService) GetByObjectTypeIDAndPropertyName(ctx context.Context, propertyName string, query PropertiesValidationGetByObjectTypeIDAndPropertyNameParams, opts ...option.RequestOption) (res *CollectionResponsePublicPropertyValidationRuleNoPaging, err error) {
	opts = slices.Concat(r.options, opts)
	if query.ObjectTypeID == "" {
		err = errors.New("missing required objectTypeId parameter")
		return nil, err
	}
	if propertyName == "" {
		err = errors.New("missing required propertyName parameter")
		return nil, err
	}
	path := fmt.Sprintf("crm/property-validations/2026-03/%s/%s", url.PathEscape(query.ObjectTypeID), url.PathEscape(propertyName))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Retrieve a specific validation rule for a property identified by its name and
// rule type.
func (r *PropertiesValidationService) GetByObjectTypeIDPropertyNameAndRuleType(ctx context.Context, ruleType PropertiesValidationGetByObjectTypeIDPropertyNameAndRuleTypeParamsRuleType, query PropertiesValidationGetByObjectTypeIDPropertyNameAndRuleTypeParams, opts ...option.RequestOption) (res *PublicPropertyValidationRule, err error) {
	opts = slices.Concat(r.options, opts)
	if query.ObjectTypeID == "" {
		err = errors.New("missing required objectTypeId parameter")
		return nil, err
	}
	if query.PropertyName == "" {
		err = errors.New("missing required propertyName parameter")
		return nil, err
	}
	path := fmt.Sprintf("crm/property-validations/2026-03/%s/%s/rule-type/%v", url.PathEscape(query.ObjectTypeID), url.PathEscape(query.PropertyName), ruleType)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Update a specific validation rule for a property identified by its name and rule
// type.
func (r *PropertiesValidationService) UpdateByObjectTypeIDPropertyNameAndRuleType(ctx context.Context, ruleType PropertiesValidationUpdateByObjectTypeIDPropertyNameAndRuleTypeParamsRuleType, params PropertiesValidationUpdateByObjectTypeIDPropertyNameAndRuleTypeParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if params.ObjectTypeID == "" {
		err = errors.New("missing required objectTypeId parameter")
		return err
	}
	if params.PropertyName == "" {
		err = errors.New("missing required propertyName parameter")
		return err
	}
	path := fmt.Sprintf("crm/property-validations/2026-03/%s/%s/rule-type/%v", url.PathEscape(params.ObjectTypeID), url.PathEscape(params.PropertyName), ruleType)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, nil, opts...)
	return err
}

type CollectionResponsePublicPropertyValidationRuleMapNoPaging struct {
	// Collection of properties with their validation rules. Each item maps a property
	// name to its configured validation rules for the specified object type.
	Results []PublicPropertyValidationRuleMap `json:"results" api:"required"`
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
	Results []PublicPropertyValidationRule `json:"results" api:"required"`
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
	RuleArguments []string `json:"ruleArguments" api:"required"`
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
	RuleType                 PublicPropertyValidationRuleRuleType `json:"ruleType" api:"required"`
	ShouldApplyNormalization bool                                 `json:"shouldApplyNormalization"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		RuleArguments            respjson.Field
		RuleType                 respjson.Field
		ShouldApplyNormalization respjson.Field
		ExtraFields              map[string]respjson.Field
		raw                      string
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
	PropertyName string `json:"propertyName" api:"required"`
	// A list of validation rules applicable to the property.
	PropertyValidationRules []PublicPropertyValidationRule `json:"propertyValidationRules" api:"required"`
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
	RuleArguments            []string        `json:"ruleArguments,omitzero" api:"required"`
	ShouldApplyNormalization param.Opt[bool] `json:"shouldApplyNormalization,omitzero"`
	paramObj
}

func (r PublicPropertyValidationRuleUpdateParam) MarshalJSON() (data []byte, err error) {
	type shadow PublicPropertyValidationRuleUpdateParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PublicPropertyValidationRuleUpdateParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PropertiesValidationGetByObjectTypeIDAndPropertyNameParams struct {
	ObjectTypeID string `path:"objectTypeId" api:"required" json:"-"`
	paramObj
}

type PropertiesValidationGetByObjectTypeIDPropertyNameAndRuleTypeParams struct {
	ObjectTypeID string `path:"objectTypeId" api:"required" json:"-"`
	PropertyName string `path:"propertyName" api:"required" json:"-"`
	paramObj
}

type PropertiesValidationGetByObjectTypeIDPropertyNameAndRuleTypeParamsRuleType string

const (
	PropertiesValidationGetByObjectTypeIDPropertyNameAndRuleTypeParamsRuleTypeAfterDatetimeDuration              PropertiesValidationGetByObjectTypeIDPropertyNameAndRuleTypeParamsRuleType = "AFTER_DATETIME_DURATION"
	PropertiesValidationGetByObjectTypeIDPropertyNameAndRuleTypeParamsRuleTypeAfterDuration                      PropertiesValidationGetByObjectTypeIDPropertyNameAndRuleTypeParamsRuleType = "AFTER_DURATION"
	PropertiesValidationGetByObjectTypeIDPropertyNameAndRuleTypeParamsRuleTypeAlphanumeric                       PropertiesValidationGetByObjectTypeIDPropertyNameAndRuleTypeParamsRuleType = "ALPHANUMERIC"
	PropertiesValidationGetByObjectTypeIDPropertyNameAndRuleTypeParamsRuleTypeBeforeDatetimeDuration             PropertiesValidationGetByObjectTypeIDPropertyNameAndRuleTypeParamsRuleType = "BEFORE_DATETIME_DURATION"
	PropertiesValidationGetByObjectTypeIDPropertyNameAndRuleTypeParamsRuleTypeBeforeDuration                     PropertiesValidationGetByObjectTypeIDPropertyNameAndRuleTypeParamsRuleType = "BEFORE_DURATION"
	PropertiesValidationGetByObjectTypeIDPropertyNameAndRuleTypeParamsRuleTypeDaysOfWeek                         PropertiesValidationGetByObjectTypeIDPropertyNameAndRuleTypeParamsRuleType = "DAYS_OF_WEEK"
	PropertiesValidationGetByObjectTypeIDPropertyNameAndRuleTypeParamsRuleTypeDecimal                            PropertiesValidationGetByObjectTypeIDPropertyNameAndRuleTypeParamsRuleType = "DECIMAL"
	PropertiesValidationGetByObjectTypeIDPropertyNameAndRuleTypeParamsRuleTypeDomain                             PropertiesValidationGetByObjectTypeIDPropertyNameAndRuleTypeParamsRuleType = "DOMAIN"
	PropertiesValidationGetByObjectTypeIDPropertyNameAndRuleTypeParamsRuleTypeEmail                              PropertiesValidationGetByObjectTypeIDPropertyNameAndRuleTypeParamsRuleType = "EMAIL"
	PropertiesValidationGetByObjectTypeIDPropertyNameAndRuleTypeParamsRuleTypeEmailAllowedDomains                PropertiesValidationGetByObjectTypeIDPropertyNameAndRuleTypeParamsRuleType = "EMAIL_ALLOWED_DOMAINS"
	PropertiesValidationGetByObjectTypeIDPropertyNameAndRuleTypeParamsRuleTypeEmailBlockedDomains                PropertiesValidationGetByObjectTypeIDPropertyNameAndRuleTypeParamsRuleType = "EMAIL_BLOCKED_DOMAINS"
	PropertiesValidationGetByObjectTypeIDPropertyNameAndRuleTypeParamsRuleTypeEndDate                            PropertiesValidationGetByObjectTypeIDPropertyNameAndRuleTypeParamsRuleType = "END_DATE"
	PropertiesValidationGetByObjectTypeIDPropertyNameAndRuleTypeParamsRuleTypeEndDatetime                        PropertiesValidationGetByObjectTypeIDPropertyNameAndRuleTypeParamsRuleType = "END_DATETIME"
	PropertiesValidationGetByObjectTypeIDPropertyNameAndRuleTypeParamsRuleTypeFormat                             PropertiesValidationGetByObjectTypeIDPropertyNameAndRuleTypeParamsRuleType = "FORMAT"
	PropertiesValidationGetByObjectTypeIDPropertyNameAndRuleTypeParamsRuleTypeMaxLength                          PropertiesValidationGetByObjectTypeIDPropertyNameAndRuleTypeParamsRuleType = "MAX_LENGTH"
	PropertiesValidationGetByObjectTypeIDPropertyNameAndRuleTypeParamsRuleTypeMaxNumber                          PropertiesValidationGetByObjectTypeIDPropertyNameAndRuleTypeParamsRuleType = "MAX_NUMBER"
	PropertiesValidationGetByObjectTypeIDPropertyNameAndRuleTypeParamsRuleTypeMinLength                          PropertiesValidationGetByObjectTypeIDPropertyNameAndRuleTypeParamsRuleType = "MIN_LENGTH"
	PropertiesValidationGetByObjectTypeIDPropertyNameAndRuleTypeParamsRuleTypeMinNumber                          PropertiesValidationGetByObjectTypeIDPropertyNameAndRuleTypeParamsRuleType = "MIN_NUMBER"
	PropertiesValidationGetByObjectTypeIDPropertyNameAndRuleTypeParamsRuleTypePhoneNumberWithExplicitCountryCode PropertiesValidationGetByObjectTypeIDPropertyNameAndRuleTypeParamsRuleType = "PHONE_NUMBER_WITH_EXPLICIT_COUNTRY_CODE"
	PropertiesValidationGetByObjectTypeIDPropertyNameAndRuleTypeParamsRuleTypeRegex                              PropertiesValidationGetByObjectTypeIDPropertyNameAndRuleTypeParamsRuleType = "REGEX"
	PropertiesValidationGetByObjectTypeIDPropertyNameAndRuleTypeParamsRuleTypeSpecialCharacters                  PropertiesValidationGetByObjectTypeIDPropertyNameAndRuleTypeParamsRuleType = "SPECIAL_CHARACTERS"
	PropertiesValidationGetByObjectTypeIDPropertyNameAndRuleTypeParamsRuleTypeStartDate                          PropertiesValidationGetByObjectTypeIDPropertyNameAndRuleTypeParamsRuleType = "START_DATE"
	PropertiesValidationGetByObjectTypeIDPropertyNameAndRuleTypeParamsRuleTypeStartDatetime                      PropertiesValidationGetByObjectTypeIDPropertyNameAndRuleTypeParamsRuleType = "START_DATETIME"
	PropertiesValidationGetByObjectTypeIDPropertyNameAndRuleTypeParamsRuleTypeURL                                PropertiesValidationGetByObjectTypeIDPropertyNameAndRuleTypeParamsRuleType = "URL"
	PropertiesValidationGetByObjectTypeIDPropertyNameAndRuleTypeParamsRuleTypeURLAllowedDomains                  PropertiesValidationGetByObjectTypeIDPropertyNameAndRuleTypeParamsRuleType = "URL_ALLOWED_DOMAINS"
	PropertiesValidationGetByObjectTypeIDPropertyNameAndRuleTypeParamsRuleTypeURLBlockedDomains                  PropertiesValidationGetByObjectTypeIDPropertyNameAndRuleTypeParamsRuleType = "URL_BLOCKED_DOMAINS"
	PropertiesValidationGetByObjectTypeIDPropertyNameAndRuleTypeParamsRuleTypeWhitespace                         PropertiesValidationGetByObjectTypeIDPropertyNameAndRuleTypeParamsRuleType = "WHITESPACE"
)

type PropertiesValidationUpdateByObjectTypeIDPropertyNameAndRuleTypeParams struct {
	ObjectTypeID                       string `path:"objectTypeId" api:"required" json:"-"`
	PropertyName                       string `path:"propertyName" api:"required" json:"-"`
	PublicPropertyValidationRuleUpdate PublicPropertyValidationRuleUpdateParam
	paramObj
}

func (r PropertiesValidationUpdateByObjectTypeIDPropertyNameAndRuleTypeParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.PublicPropertyValidationRuleUpdate)
}
func (r *PropertiesValidationUpdateByObjectTypeIDPropertyNameAndRuleTypeParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PropertiesValidationUpdateByObjectTypeIDPropertyNameAndRuleTypeParamsRuleType string

const (
	PropertiesValidationUpdateByObjectTypeIDPropertyNameAndRuleTypeParamsRuleTypeAfterDatetimeDuration              PropertiesValidationUpdateByObjectTypeIDPropertyNameAndRuleTypeParamsRuleType = "AFTER_DATETIME_DURATION"
	PropertiesValidationUpdateByObjectTypeIDPropertyNameAndRuleTypeParamsRuleTypeAfterDuration                      PropertiesValidationUpdateByObjectTypeIDPropertyNameAndRuleTypeParamsRuleType = "AFTER_DURATION"
	PropertiesValidationUpdateByObjectTypeIDPropertyNameAndRuleTypeParamsRuleTypeAlphanumeric                       PropertiesValidationUpdateByObjectTypeIDPropertyNameAndRuleTypeParamsRuleType = "ALPHANUMERIC"
	PropertiesValidationUpdateByObjectTypeIDPropertyNameAndRuleTypeParamsRuleTypeBeforeDatetimeDuration             PropertiesValidationUpdateByObjectTypeIDPropertyNameAndRuleTypeParamsRuleType = "BEFORE_DATETIME_DURATION"
	PropertiesValidationUpdateByObjectTypeIDPropertyNameAndRuleTypeParamsRuleTypeBeforeDuration                     PropertiesValidationUpdateByObjectTypeIDPropertyNameAndRuleTypeParamsRuleType = "BEFORE_DURATION"
	PropertiesValidationUpdateByObjectTypeIDPropertyNameAndRuleTypeParamsRuleTypeDaysOfWeek                         PropertiesValidationUpdateByObjectTypeIDPropertyNameAndRuleTypeParamsRuleType = "DAYS_OF_WEEK"
	PropertiesValidationUpdateByObjectTypeIDPropertyNameAndRuleTypeParamsRuleTypeDecimal                            PropertiesValidationUpdateByObjectTypeIDPropertyNameAndRuleTypeParamsRuleType = "DECIMAL"
	PropertiesValidationUpdateByObjectTypeIDPropertyNameAndRuleTypeParamsRuleTypeDomain                             PropertiesValidationUpdateByObjectTypeIDPropertyNameAndRuleTypeParamsRuleType = "DOMAIN"
	PropertiesValidationUpdateByObjectTypeIDPropertyNameAndRuleTypeParamsRuleTypeEmail                              PropertiesValidationUpdateByObjectTypeIDPropertyNameAndRuleTypeParamsRuleType = "EMAIL"
	PropertiesValidationUpdateByObjectTypeIDPropertyNameAndRuleTypeParamsRuleTypeEmailAllowedDomains                PropertiesValidationUpdateByObjectTypeIDPropertyNameAndRuleTypeParamsRuleType = "EMAIL_ALLOWED_DOMAINS"
	PropertiesValidationUpdateByObjectTypeIDPropertyNameAndRuleTypeParamsRuleTypeEmailBlockedDomains                PropertiesValidationUpdateByObjectTypeIDPropertyNameAndRuleTypeParamsRuleType = "EMAIL_BLOCKED_DOMAINS"
	PropertiesValidationUpdateByObjectTypeIDPropertyNameAndRuleTypeParamsRuleTypeEndDate                            PropertiesValidationUpdateByObjectTypeIDPropertyNameAndRuleTypeParamsRuleType = "END_DATE"
	PropertiesValidationUpdateByObjectTypeIDPropertyNameAndRuleTypeParamsRuleTypeEndDatetime                        PropertiesValidationUpdateByObjectTypeIDPropertyNameAndRuleTypeParamsRuleType = "END_DATETIME"
	PropertiesValidationUpdateByObjectTypeIDPropertyNameAndRuleTypeParamsRuleTypeFormat                             PropertiesValidationUpdateByObjectTypeIDPropertyNameAndRuleTypeParamsRuleType = "FORMAT"
	PropertiesValidationUpdateByObjectTypeIDPropertyNameAndRuleTypeParamsRuleTypeMaxLength                          PropertiesValidationUpdateByObjectTypeIDPropertyNameAndRuleTypeParamsRuleType = "MAX_LENGTH"
	PropertiesValidationUpdateByObjectTypeIDPropertyNameAndRuleTypeParamsRuleTypeMaxNumber                          PropertiesValidationUpdateByObjectTypeIDPropertyNameAndRuleTypeParamsRuleType = "MAX_NUMBER"
	PropertiesValidationUpdateByObjectTypeIDPropertyNameAndRuleTypeParamsRuleTypeMinLength                          PropertiesValidationUpdateByObjectTypeIDPropertyNameAndRuleTypeParamsRuleType = "MIN_LENGTH"
	PropertiesValidationUpdateByObjectTypeIDPropertyNameAndRuleTypeParamsRuleTypeMinNumber                          PropertiesValidationUpdateByObjectTypeIDPropertyNameAndRuleTypeParamsRuleType = "MIN_NUMBER"
	PropertiesValidationUpdateByObjectTypeIDPropertyNameAndRuleTypeParamsRuleTypePhoneNumberWithExplicitCountryCode PropertiesValidationUpdateByObjectTypeIDPropertyNameAndRuleTypeParamsRuleType = "PHONE_NUMBER_WITH_EXPLICIT_COUNTRY_CODE"
	PropertiesValidationUpdateByObjectTypeIDPropertyNameAndRuleTypeParamsRuleTypeRegex                              PropertiesValidationUpdateByObjectTypeIDPropertyNameAndRuleTypeParamsRuleType = "REGEX"
	PropertiesValidationUpdateByObjectTypeIDPropertyNameAndRuleTypeParamsRuleTypeSpecialCharacters                  PropertiesValidationUpdateByObjectTypeIDPropertyNameAndRuleTypeParamsRuleType = "SPECIAL_CHARACTERS"
	PropertiesValidationUpdateByObjectTypeIDPropertyNameAndRuleTypeParamsRuleTypeStartDate                          PropertiesValidationUpdateByObjectTypeIDPropertyNameAndRuleTypeParamsRuleType = "START_DATE"
	PropertiesValidationUpdateByObjectTypeIDPropertyNameAndRuleTypeParamsRuleTypeStartDatetime                      PropertiesValidationUpdateByObjectTypeIDPropertyNameAndRuleTypeParamsRuleType = "START_DATETIME"
	PropertiesValidationUpdateByObjectTypeIDPropertyNameAndRuleTypeParamsRuleTypeURL                                PropertiesValidationUpdateByObjectTypeIDPropertyNameAndRuleTypeParamsRuleType = "URL"
	PropertiesValidationUpdateByObjectTypeIDPropertyNameAndRuleTypeParamsRuleTypeURLAllowedDomains                  PropertiesValidationUpdateByObjectTypeIDPropertyNameAndRuleTypeParamsRuleType = "URL_ALLOWED_DOMAINS"
	PropertiesValidationUpdateByObjectTypeIDPropertyNameAndRuleTypeParamsRuleTypeURLBlockedDomains                  PropertiesValidationUpdateByObjectTypeIDPropertyNameAndRuleTypeParamsRuleType = "URL_BLOCKED_DOMAINS"
	PropertiesValidationUpdateByObjectTypeIDPropertyNameAndRuleTypeParamsRuleTypeWhitespace                         PropertiesValidationUpdateByObjectTypeIDPropertyNameAndRuleTypeParamsRuleType = "WHITESPACE"
)
