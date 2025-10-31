// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package account

import (
	"context"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/stainless-sdks/hubspot-sdk-go/internal/apijson"
	"github.com/stainless-sdks/hubspot-sdk-go/internal/apiquery"
	"github.com/stainless-sdks/hubspot-sdk-go/internal/requestconfig"
	"github.com/stainless-sdks/hubspot-sdk-go/option"
	"github.com/stainless-sdks/hubspot-sdk-go/packages/param"
	"github.com/stainless-sdks/hubspot-sdk-go/packages/respjson"
	"github.com/stainless-sdks/hubspot-sdk-go/shared"
)

// ActivityService contains methods and other services that help with interacting
// with the Hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewActivityService] method instead.
type ActivityService struct {
	Options []option.RequestOption
}

// NewActivityService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewActivityService(opts ...option.RequestOption) (r ActivityService) {
	r = ActivityService{}
	r.Options = opts
	return
}

// Retrieve activity history for user actions related to approvals, content
// updates, CRM object updates, security activity, and more (Enterprise only).
// Learn more about
// [activities included in audit log exports](https://knowledge.hubspot.com/account-management/view-and-export-account-activity-history-in-a-centralized-audit-log?hubs_content=knowledge.hubspot.com/account-management/view-and-export-account-activity-history&hubs_content-cta=centralized%20audit%20log#data-included-in-the-centralized-audit-log).
func (r *ActivityService) ListAuditLogs(ctx context.Context, query ActivityListAuditLogsParams, opts ...option.RequestOption) (res *CollectionResponsePublicAPIUserActionEventForwardPaging, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "account-info/v3/activity/audit-logs"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Retrieve logs of user actions related to
// [login activity](https://knowledge.hubspot.com/account-management/view-and-export-account-activity-history#account-login-history).
func (r *ActivityService) ListLoginActivities(ctx context.Context, query ActivityListLoginActivitiesParams, opts ...option.RequestOption) (res *CollectionResponsePublicLoginAuditForwardPaging, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "account-info/v3/activity/login"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Retrieve logs of user actions related to
// [security activity](https://knowledge.hubspot.com/account-management/view-and-export-account-activity-history#security-activity-history).
func (r *ActivityService) ListSecurityActivities(ctx context.Context, query ActivityListSecurityActivitiesParams, opts ...option.RequestOption) (res *CollectionResponseHydratedCriticalActionForwardPaging, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "account-info/v3/activity/security"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type ActingUser struct {
	// The ID of the user who performed the action.
	UserID int64 `json:"userId,required"`
	// The email address of the user who performed the action.
	UserEmail string `json:"userEmail"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		UserID      respjson.Field
		UserEmail   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ActingUser) RawJSON() string { return r.JSON.raw }
func (r *ActingUser) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CollectionResponseHydratedCriticalActionForwardPaging struct {
	Results []HydratedCriticalAction `json:"results,required"`
	Paging  shared.ForwardPaging     `json:"paging"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Results     respjson.Field
		Paging      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CollectionResponseHydratedCriticalActionForwardPaging) RawJSON() string { return r.JSON.raw }
func (r *CollectionResponseHydratedCriticalActionForwardPaging) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CollectionResponsePublicAPIUserActionEventForwardPaging struct {
	Results []PublicAPIUserActionEvent `json:"results,required"`
	Paging  shared.ForwardPaging       `json:"paging"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Results     respjson.Field
		Paging      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CollectionResponsePublicAPIUserActionEventForwardPaging) RawJSON() string { return r.JSON.raw }
func (r *CollectionResponsePublicAPIUserActionEventForwardPaging) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CollectionResponsePublicLoginAuditForwardPaging struct {
	Results []PublicLoginAudit   `json:"results,required"`
	Paging  shared.ForwardPaging `json:"paging"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Results     respjson.Field
		Paging      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CollectionResponsePublicLoginAuditForwardPaging) RawJSON() string { return r.JSON.raw }
func (r *CollectionResponsePublicLoginAuditForwardPaging) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Details about the a particular security activity for a HubSpot account.
type HydratedCriticalAction struct {
	// The unique ID of the activity.
	ID string `json:"id,required"`
	// The time the activity took place.
	CreatedAt time.Time `json:"createdAt,required" format:"date-time"`
	// The type of activity.
	Type string `json:"type,required"`
	// The user's unique ID.
	UserID int64 `json:"userId,required"`
	// Email address of the user associated with the activity.
	ActingUser string `json:"actingUser"`
	// The approximate country code.
	CountryCode string `json:"countryCode"`
	// A link to the URL where the action was taken in the account.
	InfoURL string `json:"infoUrl"`
	// IP address where the activity originated.
	IPAddress string `json:"ipAddress"`
	Location  string `json:"location"`
	// The ID of the affected object.
	ObjectID string `json:"objectId"`
	// The approximate region code.
	RegionCode string `json:"regionCode"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		Type        respjson.Field
		UserID      respjson.Field
		ActingUser  respjson.Field
		CountryCode respjson.Field
		InfoURL     respjson.Field
		IPAddress   respjson.Field
		Location    respjson.Field
		ObjectID    respjson.Field
		RegionCode  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r HydratedCriticalAction) RawJSON() string { return r.JSON.raw }
func (r *HydratedCriticalAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PublicAPIUserActionEvent struct {
	// The unique ID of the activity.
	ID         string     `json:"id,required"`
	ActingUser ActingUser `json:"actingUser,required"`
	// The type of action taken.
	Action string `json:"action,required"`
	// The category of the activity.
	Category string `json:"category,required"`
	// The time that the action occurred at.
	OccurredAt time.Time `json:"occurredAt,required" format:"date-time"`
	// The subcategory of the activity.
	SubCategory string `json:"subCategory"`
	// The ID of the impacted object.
	TargetObjectID string `json:"targetObjectId"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID             respjson.Field
		ActingUser     respjson.Field
		Action         respjson.Field
		Category       respjson.Field
		OccurredAt     respjson.Field
		SubCategory    respjson.Field
		TargetObjectID respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PublicAPIUserActionEvent) RawJSON() string { return r.JSON.raw }
func (r *PublicAPIUserActionEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Details about the a particular login activity for a HubSpot account.
type PublicLoginAudit struct {
	// The login activity's unique ID.
	ID string `json:"id,required"`
	// The time the login took place.
	LoginAt time.Time `json:"loginAt,required" format:"date-time"`
	// Whether the login was successful or not.
	LoginSucceeded bool `json:"loginSucceeded,required"`
	// The approximate country code of the login.
	CountryCode string `json:"countryCode"`
	// Email address of the user associated with the login.
	Email string `json:"email"`
	// IP address where the activity originated.
	IPAddress string `json:"ipAddress"`
	Location  string `json:"location"`
	// The approximate region code of the login.
	RegionCode string `json:"regionCode"`
	// Information about the device used for logging in.
	UserAgent string `json:"userAgent"`
	// The user's unique ID.
	UserID int64 `json:"userId"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID             respjson.Field
		LoginAt        respjson.Field
		LoginSucceeded respjson.Field
		CountryCode    respjson.Field
		Email          respjson.Field
		IPAddress      respjson.Field
		Location       respjson.Field
		RegionCode     respjson.Field
		UserAgent      respjson.Field
		UserID         respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PublicLoginAudit) RawJSON() string { return r.JSON.raw }
func (r *PublicLoginAudit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ActivityListAuditLogsParams struct {
	// The paging cursor token of the last successfully read resource will be returned
	// as the `paging.next.after` JSON property of a paged response containing more
	// results.
	After param.Opt[string] `query:"after,omitzero" json:"-"`
	// The maximum number of results to display per page.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// A timestamp, as a starting point for retrieving activity logs.
	OccurredAfter param.Opt[time.Time] `query:"occurredAfter,omitzero" format:"date-time" json:"-"`
	// A timestamp, as an end point for retrieving activity logs.
	OccurredBefore param.Opt[time.Time] `query:"occurredBefore,omitzero" format:"date-time" json:"-"`
	// The ID of a user, for retrieving user-specific logs.
	ActingUserID []int64 `query:"actingUserId,omitzero" json:"-"`
	// Set to `occurredAt` to order results by the time of the event. By default,
	// events are ordered from oldest to newest.
	Sort []string `query:"sort,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ActivityListAuditLogsParams]'s query parameters as
// `url.Values`.
func (r ActivityListAuditLogsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ActivityListLoginActivitiesParams struct {
	// The cursor token value to get the next set of results. You can get this from the
	// `paging.next.after` JSON property of a paged response containing more results.
	After param.Opt[string] `query:"after,omitzero" json:"-"`
	// The maximum number of results to display per page. Max value of limit is 200.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// The ID of a user, for retrieving user-specific logs.
	UserID param.Opt[int64] `query:"userId,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ActivityListLoginActivitiesParams]'s query parameters as
// `url.Values`.
func (r ActivityListLoginActivitiesParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ActivityListSecurityActivitiesParams struct {
	// The cursor token value to get the next set of results. You can get this from the
	// `paging.next.after` JSON property of a paged response containing more results.
	After param.Opt[string] `query:"after,omitzero" json:"-"`
	// The start time, for retrieving logs within a specific timeframe.
	FromTimestamp param.Opt[int64] `query:"fromTimestamp,omitzero" json:"-"`
	// The maximum number of results to display per page. Max value of limit is 200.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// The end time, for retrieving logs within a specific timeframe.
	ToTimestamp param.Opt[int64] `query:"toTimestamp,omitzero" json:"-"`
	// The ID of a user, for retrieving user-specific logs.
	UserID param.Opt[int64] `query:"userId,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ActivityListSecurityActivitiesParams]'s query parameters as
// `url.Values`.
func (r ActivityListSecurityActivitiesParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
