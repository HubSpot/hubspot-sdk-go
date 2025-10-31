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
	"github.com/stainless-sdks/hubspot-sdk-go/internal/apiquery"
	"github.com/stainless-sdks/hubspot-sdk-go/internal/requestconfig"
	"github.com/stainless-sdks/hubspot-sdk-go/option"
	"github.com/stainless-sdks/hubspot-sdk-go/packages/param"
	"github.com/stainless-sdks/hubspot-sdk-go/packages/respjson"
)

// LimitService contains methods and other services that help with interacting with
// the Hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewLimitService] method instead.
type LimitService struct {
	Options []option.RequestOption
}

// NewLimitService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewLimitService(opts ...option.RequestOption) (r LimitService) {
	r = LimitService{}
	r.Options = opts
	return
}

// Returns limits and usage for custom association labels
func (r *LimitService) GetAssociationLabelLimits(ctx context.Context, query LimitGetAssociationLabelLimitsParams, opts ...option.RequestOption) (res *CollectionResponseAssociationLabelLimitResponseNoPaging, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "crm/v3/limits/associations/labels"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Returns records approaching or at association limits between two objects
func (r *LimitService) GetAssociationRecordsLimitsByObjectType(ctx context.Context, toObjectTypeID string, query LimitGetAssociationRecordsLimitsByObjectTypeParams, opts ...option.RequestOption) (res *AssociationRecordLimitResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if query.FromObjectTypeID == "" {
		err = errors.New("missing required fromObjectTypeId parameter")
		return
	}
	if toObjectTypeID == "" {
		err = errors.New("missing required toObjectTypeId parameter")
		return
	}
	path := fmt.Sprintf("crm/v3/limits/associations/records/%s/%s", query.FromObjectTypeID, toObjectTypeID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Returns objects with records approaching or at association limits
func (r *LimitService) GetAssociationRecordsLimitsFromObjects(ctx context.Context, opts ...option.RequestOption) (res *CollectionResponseObjectTypeNearOrAtAssociationLimitNoPaging, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "crm/v3/limits/associations/records/from"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Returns objects for which the from object has records approaching or at
// association limits
func (r *LimitService) GetAssociationRecordsLimitsToObjects(ctx context.Context, fromObjectTypeID string, opts ...option.RequestOption) (res *CollectionResponseObjectTypeNearOrAtAssociationLimitNoPaging, err error) {
	opts = slices.Concat(r.Options, opts)
	if fromObjectTypeID == "" {
		err = errors.New("missing required fromObjectTypeId parameter")
		return
	}
	path := fmt.Sprintf("crm/v3/limits/associations/records/%s/to", fromObjectTypeID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Returns overall limit and per object usage for calculated properties
func (r *LimitService) GetCalculatedPropertyLimits(ctx context.Context, opts ...option.RequestOption) (res *CalculatedPropertyLimitResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "crm/v3/limits/calculated-properties"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Returns limits and usage for custom object schemas
func (r *LimitService) GetCustomObjectTypeLimits(ctx context.Context, opts ...option.RequestOption) (res *CustomObjectLimitResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "crm/v3/limits/custom-object-types"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Returns limits and usage per object for custom properties
func (r *LimitService) GetCustomPropertyLimits(ctx context.Context, opts ...option.RequestOption) (res *CustomPropertyLimitResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "crm/v3/limits/custom-properties"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Returns limits and usage per object for pipelines
func (r *LimitService) GetPipelineLimits(ctx context.Context, opts ...option.RequestOption) (res *PipelineLimitResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "crm/v3/limits/pipelines"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Returns limits and usage per object for records
func (r *LimitService) GetRecordLimits(ctx context.Context, opts ...option.RequestOption) (res *RecordLimitResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "crm/v3/limits/records"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type AssociationLabelLimitResponse struct {
	AllLabels []string `json:"allLabels,required"`
	// Defines an object type.
	FromObjectType ObjectTypeDefinition `json:"fromObjectType,required"`
	Limit          int64                `json:"limit,required"`
	Percentage     float64              `json:"percentage,required"`
	// Defines an object type.
	ToObjectType ObjectTypeDefinition `json:"toObjectType,required"`
	Usage        int64                `json:"usage,required"`
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
	AtLimitFromRecordSamples   []AtLimitRecordSample   `json:"atLimitFromRecordSamples,required"`
	Limit                      int64                   `json:"limit,required"`
	NearLimitFromRecordSamples []NearLimitRecordSample `json:"nearLimitFromRecordSamples,required"`
	TotalRecordsAtLimit        int64                   `json:"totalRecordsAtLimit,required"`
	TotalRecordsNearLimit      int64                   `json:"totalRecordsNearLimit,required"`
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
	Label    string `json:"label,required"`
	ObjectID int64  `json:"objectId,required"`
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
	ByObjectType      []UsageForObjectType `json:"byObjectType,required"`
	OverallLimit      int64                `json:"overallLimit,required"`
	OverallPercentage float64              `json:"overallPercentage,required"`
	OverallUsage      int64                `json:"overallUsage,required"`
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
	Results []AssociationLabelLimitResponse `json:"results,required"`
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
	Results []ObjectTypeNearOrAtAssociationLimit `json:"results,required"`
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
	Limit      int64   `json:"limit,required"`
	Percentage float64 `json:"percentage,required"`
	Usage      int64   `json:"usage,required"`
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
	ByObjectType      []UsageForObjectType `json:"byObjectType,required"`
	OverallLimit      int64                `json:"overallLimit,required"`
	OverallPercentage float64              `json:"overallPercentage,required"`
	OverallUsage      int64                `json:"overallUsage,required"`
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
	ByObjectType      []LimitAndUsageForObjectType `json:"byObjectType,required"`
	OverallLimit      int64                        `json:"overallLimit,required"`
	OverallPercentage float64                      `json:"overallPercentage,required"`
	OverallUsage      int64                        `json:"overallUsage,required"`
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
	Limit         int64   `json:"limit,required"`
	ObjectTypeID  string  `json:"objectTypeId,required"`
	Percentage    float64 `json:"percentage,required"`
	PluralLabel   string  `json:"pluralLabel,required"`
	SingularLabel string  `json:"singularLabel,required"`
	Usage         int64   `json:"usage,required"`
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

type NearLimitRecordSample struct {
	Label      string  `json:"label,required"`
	ObjectID   int64   `json:"objectId,required"`
	Percentage float64 `json:"percentage,required"`
	Usage      int64   `json:"usage,required"`
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
	HasRecordsAtLimit   bool   `json:"hasRecordsAtLimit,required"`
	HasRecordsNearLimit bool   `json:"hasRecordsNearLimit,required"`
	ObjectTypeID        string `json:"objectTypeId,required"`
	PluralLabel         string `json:"pluralLabel,required"`
	SingularLabel       string `json:"singularLabel,required"`
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
	CustomObjectTypes         CustomObjectRecordLimitResponse `json:"customObjectTypes,required"`
	HubspotDefinedObjectTypes []LimitAndUsageForObjectType    `json:"hubspotDefinedObjectTypes,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CustomObjectTypes         respjson.Field
		HubspotDefinedObjectTypes respjson.Field
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
	CustomObjectTypes         CustomObjectRecordLimitResponse `json:"customObjectTypes,required"`
	HubspotDefinedObjectTypes []LimitAndUsageForObjectType    `json:"hubspotDefinedObjectTypes,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CustomObjectTypes         respjson.Field
		HubspotDefinedObjectTypes respjson.Field
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
	ObjectTypeID  string `json:"objectTypeId,required"`
	PluralLabel   string `json:"pluralLabel,required"`
	SingularLabel string `json:"singularLabel,required"`
	Usage         int64  `json:"usage,required"`
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
	FromObjectTypeID string `path:"fromObjectTypeId,required" json:"-"`
	paramObj
}
