// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package crm

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/stainless-sdks/hubspot-sdk-go/internal/apijson"
	"github.com/stainless-sdks/hubspot-sdk-go/internal/requestconfig"
	"github.com/stainless-sdks/hubspot-sdk-go/option"
	"github.com/stainless-sdks/hubspot-sdk-go/packages/respjson"
)

// PropertyValidationService contains methods and other services that help with
// interacting with the Hubspot API.
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
func (r *PropertyValidationService) List(ctx context.Context, objectTypeID string, opts ...option.RequestOption) (res *PropertyValidationListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if objectTypeID == "" {
		err = errors.New("missing required objectTypeId parameter")
		return
	}
	path := fmt.Sprintf("crm/v3/property-validations/%s", objectTypeID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Read a property's validation rules identified by {propertyName}.
func (r *PropertyValidationService) Get(ctx context.Context, propertyName string, query PropertyValidationGetParams, opts ...option.RequestOption) (res *PropertyValidationGetResponse, err error) {
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

type PropertyValidationListResponse struct {
	Results []PropertyValidationListResponseResult `json:"results,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Results     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PropertyValidationListResponse) RawJSON() string { return r.JSON.raw }
func (r *PropertyValidationListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PropertyValidationListResponseResult struct {
	PropertyName            string                                                       `json:"propertyName,required"`
	PropertyValidationRules []PropertyValidationListResponseResultPropertyValidationRule `json:"propertyValidationRules,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PropertyName            respjson.Field
		PropertyValidationRules respjson.Field
		ExtraFields             map[string]respjson.Field
		raw                     string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PropertyValidationListResponseResult) RawJSON() string { return r.JSON.raw }
func (r *PropertyValidationListResponseResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PropertyValidationListResponseResultPropertyValidationRule struct {
	RuleArguments []string `json:"ruleArguments,required"`
	RuleType      string   `json:"ruleType,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		RuleArguments respjson.Field
		RuleType      respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PropertyValidationListResponseResultPropertyValidationRule) RawJSON() string {
	return r.JSON.raw
}
func (r *PropertyValidationListResponseResultPropertyValidationRule) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PropertyValidationGetResponse struct {
	Results []PropertyValidationGetResponseResult `json:"results,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Results     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PropertyValidationGetResponse) RawJSON() string { return r.JSON.raw }
func (r *PropertyValidationGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PropertyValidationGetResponseResult struct {
	RuleArguments []string `json:"ruleArguments,required"`
	RuleType      string   `json:"ruleType,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		RuleArguments respjson.Field
		RuleType      respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PropertyValidationGetResponseResult) RawJSON() string { return r.JSON.raw }
func (r *PropertyValidationGetResponseResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PropertyValidationGetParams struct {
	ObjectTypeID string `path:"objectTypeId,required" json:"-"`
	paramObj
}
