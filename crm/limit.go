// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package crm

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/HubSpot/hubspot-sdk-go/internal/apijson"
	"github.com/HubSpot/hubspot-sdk-go/internal/apiquery"
	"github.com/HubSpot/hubspot-sdk-go/internal/requestconfig"
	"github.com/HubSpot/hubspot-sdk-go/option"
	"github.com/HubSpot/hubspot-sdk-go/packages/param"
	"github.com/HubSpot/hubspot-sdk-go/packages/respjson"
)

// LimitService contains methods and other services that help with interacting with
// the hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewLimitService] method instead.
type LimitService struct {
	options []option.RequestOption
}

// NewLimitService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewLimitService(opts ...option.RequestOption) (r LimitService) {
	r = LimitService{}
	r.options = opts
	return
}

// Returns limits and usage for custom association labels
func (r *LimitService) GetAssociationLabelLimits(ctx context.Context, query LimitGetAssociationLabelLimitsParams, opts ...option.RequestOption) (res *CollectionResponseAssociationLabelLimitResponseNoPaging, err error) {
	opts = slices.Concat(r.options, opts)
	path := "crm/limits/2026-03/associations/labels"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Returns records approaching or at association limits between two objects
func (r *LimitService) GetAssociationRecordsLimitsByObjectType(ctx context.Context, toObjectTypeID string, query LimitGetAssociationRecordsLimitsByObjectTypeParams, opts ...option.RequestOption) (res *AssociationRecordLimitResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if query.FromObjectTypeID == "" {
		err = errors.New("missing required fromObjectTypeId parameter")
		return nil, err
	}
	if toObjectTypeID == "" {
		err = errors.New("missing required toObjectTypeId parameter")
		return nil, err
	}
	path := fmt.Sprintf("crm/limits/2026-03/associations/records/%s/%s", url.PathEscape(query.FromObjectTypeID), url.PathEscape(toObjectTypeID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Returns objects with records approaching or at association limits
func (r *LimitService) GetAssociationRecordsLimitsFromObjects(ctx context.Context, opts ...option.RequestOption) (res *CollectionResponseObjectTypeNearOrAtAssociationLimitNoPaging, err error) {
	opts = slices.Concat(r.options, opts)
	path := "crm/limits/2026-03/associations/records/from"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Returns objects for which the from object has records approaching or at
// association limits
func (r *LimitService) GetAssociationRecordsLimitsToObjects(ctx context.Context, fromObjectTypeID string, opts ...option.RequestOption) (res *CollectionResponseObjectTypeNearOrAtAssociationLimitNoPaging, err error) {
	opts = slices.Concat(r.options, opts)
	if fromObjectTypeID == "" {
		err = errors.New("missing required fromObjectTypeId parameter")
		return nil, err
	}
	path := fmt.Sprintf("crm/limits/2026-03/associations/records/%s/to", url.PathEscape(fromObjectTypeID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Returns overall limit and per object usage for calculated properties
func (r *LimitService) GetCalculatedPropertyLimits(ctx context.Context, opts ...option.RequestOption) (res *CalculatedPropertyLimitResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "crm/limits/2026-03/calculated-properties"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Returns limits and usage for custom object schemas
func (r *LimitService) GetCustomObjectTypeLimits(ctx context.Context, opts ...option.RequestOption) (res *CustomObjectLimitResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "crm/limits/2026-03/custom-object-types"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Returns limits and usage per object for custom properties
func (r *LimitService) GetCustomPropertyLimits(ctx context.Context, opts ...option.RequestOption) (res *CustomPropertyLimitResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "crm/limits/2026-03/custom-properties"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Returns limits and usage per object for pipelines
func (r *LimitService) GetPipelineLimits(ctx context.Context, opts ...option.RequestOption) (res *PipelineLimitResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "crm/limits/2026-03/pipelines"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Returns limits and usage per object for records
func (r *LimitService) GetRecordLimits(ctx context.Context, opts ...option.RequestOption) (res *RecordLimitResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "crm/limits/2026-03/records"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type AssociationLabelLimitResponse struct {
	// A list of all association labels.
	AllLabels      []string                   `json:"allLabels" api:"required"`
	FromObjectType LimitsObjectTypeDefinition `json:"fromObjectType" api:"required"`
	// The maximum number of association labels allowed.
	Limit int64 `json:"limit" api:"required"`
	// The percentage of the association label limit that has been used.
	Percentage   float64                    `json:"percentage" api:"required"`
	ToObjectType LimitsObjectTypeDefinition `json:"toObjectType" api:"required"`
	// The current number of association labels used.
	Usage int64 `json:"usage" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AllLabels      respjson.Field
		FromObjectType respjson.Field
		Limit          respjson.Field
		Percentage     respjson.Field
		ToObjectType   respjson.Field
		Usage          respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AssociationLabelLimitResponse) RawJSON() string { return r.JSON.raw }
func (r *AssociationLabelLimitResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AssociationRecordLimitResponse struct {
	AtLimitFromRecordSamples []AtLimitRecordSample `json:"atLimitFromRecordSamples" api:"required"`
	// The maximum number of associations allowed for records.
	Limit                      int64                   `json:"limit" api:"required"`
	NearLimitFromRecordSamples []NearLimitRecordSample `json:"nearLimitFromRecordSamples" api:"required"`
	// The total number of records that have reached their association limit.
	TotalRecordsAtLimit int64 `json:"totalRecordsAtLimit" api:"required"`
	// The total number of records that are approaching their association limit.
	TotalRecordsNearLimit int64 `json:"totalRecordsNearLimit" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AtLimitFromRecordSamples   respjson.Field
		Limit                      respjson.Field
		NearLimitFromRecordSamples respjson.Field
		TotalRecordsAtLimit        respjson.Field
		TotalRecordsNearLimit      respjson.Field
		ExtraFields                map[string]respjson.Field
		raw                        string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AssociationRecordLimitResponse) RawJSON() string { return r.JSON.raw }
func (r *AssociationRecordLimitResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AtLimitRecordSample struct {
	// The label associated with a record that is at its limit.
	Label string `json:"label" api:"required"`
	// The objectId of the object that is at its limit.
	ObjectID int64 `json:"objectId" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Label       respjson.Field
		ObjectID    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AtLimitRecordSample) RawJSON() string { return r.JSON.raw }
func (r *AtLimitRecordSample) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CalculatedPropertyLimitResponse struct {
	ByObjectType []UsageForObjectType `json:"byObjectType" api:"required"`
	// The maximum number of calculated properties allowed.
	OverallLimit int64 `json:"overallLimit" api:"required"`
	// The percentage of the overall limit that is currently being used for calculated
	// properties.
	OverallPercentage float64 `json:"overallPercentage" api:"required"`
	// The total number of calculated properties currently in use.
	OverallUsage int64 `json:"overallUsage" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ByObjectType      respjson.Field
		OverallLimit      respjson.Field
		OverallPercentage respjson.Field
		OverallUsage      respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CalculatedPropertyLimitResponse) RawJSON() string { return r.JSON.raw }
func (r *CalculatedPropertyLimitResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CollectionResponseAssociationLabelLimitResponseNoPaging struct {
	Results []AssociationLabelLimitResponse `json:"results" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Results     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CollectionResponseAssociationLabelLimitResponseNoPaging) RawJSON() string { return r.JSON.raw }
func (r *CollectionResponseAssociationLabelLimitResponseNoPaging) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CollectionResponseObjectTypeNearOrAtAssociationLimitNoPaging struct {
	Results []ObjectTypeNearOrAtAssociationLimit `json:"results" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Results     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CollectionResponseObjectTypeNearOrAtAssociationLimitNoPaging) RawJSON() string {
	return r.JSON.raw
}
func (r *CollectionResponseObjectTypeNearOrAtAssociationLimitNoPaging) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CustomObjectLimitResponse struct {
	// The maximum number of custom objects allowed.
	Limit int64 `json:"limit" api:"required"`
	// The percentage of the custom object limit that is currently used.
	Percentage float64 `json:"percentage" api:"required"`
	// The current number of custom objects used.
	Usage int64 `json:"usage" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Limit       respjson.Field
		Percentage  respjson.Field
		Usage       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CustomObjectLimitResponse) RawJSON() string { return r.JSON.raw }
func (r *CustomObjectLimitResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CustomObjectRecordLimitResponse struct {
	ByObjectType []UsageForObjectType `json:"byObjectType" api:"required"`
	// The maximum number of custom object records allowed.
	OverallLimit int64 `json:"overallLimit" api:"required"`
	// The percentage of the overall custom object record limit that has been used.
	OverallPercentage float64 `json:"overallPercentage" api:"required"`
	// The total number of custom object records currently in use.
	OverallUsage int64 `json:"overallUsage" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ByObjectType      respjson.Field
		OverallLimit      respjson.Field
		OverallPercentage respjson.Field
		OverallUsage      respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CustomObjectRecordLimitResponse) RawJSON() string { return r.JSON.raw }
func (r *CustomObjectRecordLimitResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CustomPropertyLimitResponse struct {
	ByObjectType []LimitAndUsageForObjectType `json:"byObjectType" api:"required"`
	// The total limit for custom properties across all objects.
	OverallLimit int64 `json:"overallLimit" api:"required"`
	// The percentage of the overall custom property limit that has been used.
	OverallPercentage float64 `json:"overallPercentage" api:"required"`
	// The total number of custom properties currently in use across all objects.
	OverallUsage int64 `json:"overallUsage" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ByObjectType      respjson.Field
		OverallLimit      respjson.Field
		OverallPercentage respjson.Field
		OverallUsage      respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CustomPropertyLimitResponse) RawJSON() string { return r.JSON.raw }
func (r *CustomPropertyLimitResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LimitAndUsageForObjectType struct {
	// The maximum allowed count for the object type.
	Limit int64 `json:"limit" api:"required"`
	// The unique identifier for the object type.
	ObjectTypeID string `json:"objectTypeId" api:"required"`
	// The percentage of the limit that has been used.
	Percentage float64 `json:"percentage" api:"required"`
	// The plural label for the object type.
	PluralLabel string `json:"pluralLabel" api:"required"`
	// The singular label for the object type.
	SingularLabel string `json:"singularLabel" api:"required"`
	// The current usage count for the object type.
	Usage int64 `json:"usage" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Limit         respjson.Field
		ObjectTypeID  respjson.Field
		Percentage    respjson.Field
		PluralLabel   respjson.Field
		SingularLabel respjson.Field
		Usage         respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LimitAndUsageForObjectType) RawJSON() string { return r.JSON.raw }
func (r *LimitAndUsageForObjectType) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LimitsObjectTypeDefinition struct {
	// The unique identifier for the object type.
	ObjectTypeID string `json:"objectTypeId" api:"required"`
	// The plural form label for the object type.
	PluralLabel string `json:"pluralLabel" api:"required"`
	// The singular form label for the object type.
	SingularLabel string `json:"singularLabel" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ObjectTypeID  respjson.Field
		PluralLabel   respjson.Field
		SingularLabel respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LimitsObjectTypeDefinition) RawJSON() string { return r.JSON.raw }
func (r *LimitsObjectTypeDefinition) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NearLimitRecordSample struct {
	// The primary identifier of the record.
	Label string `json:"label" api:"required"`
	// The unique identifier for the object.
	ObjectID int64 `json:"objectId" api:"required"`
	// The percentage of the limit that has been used.
	Percentage float64 `json:"percentage" api:"required"`
	// The number of records currently in use.
	Usage int64 `json:"usage" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Label       respjson.Field
		ObjectID    respjson.Field
		Percentage  respjson.Field
		Usage       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NearLimitRecordSample) RawJSON() string { return r.JSON.raw }
func (r *NearLimitRecordSample) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ObjectTypeNearOrAtAssociationLimit struct {
	// Indicates whether there are records that have reached the association limit.
	HasRecordsAtLimit bool `json:"hasRecordsAtLimit" api:"required"`
	// Indicates whether there are records that are approaching the association limit.
	HasRecordsNearLimit bool `json:"hasRecordsNearLimit" api:"required"`
	// The unique identifier for the object type.
	ObjectTypeID string `json:"objectTypeId" api:"required"`
	// The plural form of the label for the object type.
	PluralLabel string `json:"pluralLabel" api:"required"`
	// The singular form of the label for the object type.
	SingularLabel string `json:"singularLabel" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		HasRecordsAtLimit   respjson.Field
		HasRecordsNearLimit respjson.Field
		ObjectTypeID        respjson.Field
		PluralLabel         respjson.Field
		SingularLabel       respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ObjectTypeNearOrAtAssociationLimit) RawJSON() string { return r.JSON.raw }
func (r *ObjectTypeNearOrAtAssociationLimit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PipelineLimitResponse struct {
	CustomObjectTypes         CustomObjectRecordLimitResponse `json:"customObjectTypes" api:"required"`
	HubSpotDefinedObjectTypes []LimitAndUsageForObjectType    `json:"hubspotDefinedObjectTypes" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CustomObjectTypes         respjson.Field
		HubSpotDefinedObjectTypes respjson.Field
		ExtraFields               map[string]respjson.Field
		raw                       string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PipelineLimitResponse) RawJSON() string { return r.JSON.raw }
func (r *PipelineLimitResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RecordLimitResponse struct {
	CustomObjectTypes         CustomObjectRecordLimitResponse `json:"customObjectTypes" api:"required"`
	HubSpotDefinedObjectTypes []LimitAndUsageForObjectType    `json:"hubspotDefinedObjectTypes" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CustomObjectTypes         respjson.Field
		HubSpotDefinedObjectTypes respjson.Field
		ExtraFields               map[string]respjson.Field
		raw                       string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RecordLimitResponse) RawJSON() string { return r.JSON.raw }
func (r *RecordLimitResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UsageForObjectType struct {
	// The unique identifier for the object type.
	ObjectTypeID string `json:"objectTypeId" api:"required"`
	// The plural form of the label for the object type.
	PluralLabel string `json:"pluralLabel" api:"required"`
	// The singular form of the label for the object type.
	SingularLabel string `json:"singularLabel" api:"required"`
	// The number of records used for the object type.
	Usage int64 `json:"usage" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ObjectTypeID  respjson.Field
		PluralLabel   respjson.Field
		SingularLabel respjson.Field
		Usage         respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UsageForObjectType) RawJSON() string { return r.JSON.raw }
func (r *UsageForObjectType) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LimitGetAssociationLabelLimitsParams struct {
	FromObjectTypeID param.Opt[string] `query:"fromObjectTypeId,omitzero" json:"-"`
	ToObjectTypeID   param.Opt[string] `query:"toObjectTypeId,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [LimitGetAssociationLabelLimitsParams]'s query parameters as
// `url.Values`.
func (r LimitGetAssociationLabelLimitsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type LimitGetAssociationRecordsLimitsByObjectTypeParams struct {
	FromObjectTypeID string `path:"fromObjectTypeId" api:"required" json:"-"`
	paramObj
}
