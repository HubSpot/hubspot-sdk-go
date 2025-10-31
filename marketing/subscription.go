// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package marketing

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"slices"
	"time"

	"github.com/stainless-sdks/hubspot-sdk-go/internal/apijson"
	shimjson "github.com/stainless-sdks/hubspot-sdk-go/internal/encoding/json"
	"github.com/stainless-sdks/hubspot-sdk-go/internal/requestconfig"
	"github.com/stainless-sdks/hubspot-sdk-go/option"
	"github.com/stainless-sdks/hubspot-sdk-go/packages/param"
	"github.com/stainless-sdks/hubspot-sdk-go/packages/respjson"
)

// SubscriptionService contains methods and other services that help with
// interacting with the Hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSubscriptionService] method instead.
type SubscriptionService struct {
	Options []option.RequestOption
	V4      SubscriptionV4Service
}

// NewSubscriptionService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewSubscriptionService(opts ...option.RequestOption) (r SubscriptionService) {
	r = SubscriptionService{}
	r.Options = opts
	r.V4 = NewSubscriptionV4Service(opts...)
	return
}

// Get a list of all subscription definitions for the portal
func (r *SubscriptionService) List(ctx context.Context, opts ...option.RequestOption) (res *SubscriptionDefinitionsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "communication-preferences/v3/definitions"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Returns a list of subscriptions and their status for a given contact.
func (r *SubscriptionService) GetEmailStatus(ctx context.Context, emailAddress string, opts ...option.RequestOption) (res *PublicSubscriptionStatusesResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if emailAddress == "" {
		err = errors.New("missing required emailAddress parameter")
		return
	}
	path := fmt.Sprintf("communication-preferences/v3/status/email/%s", emailAddress)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Subscribes a contact to the given subscription type. This API is not valid to
// use for subscribing a contact at a brand or portal level and will return an
// error.
func (r *SubscriptionService) Subscribe(ctx context.Context, body SubscriptionSubscribeParams, opts ...option.RequestOption) (res *PublicSubscriptionStatus, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "communication-preferences/v3/subscribe"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Unsubscribes a contact from the given subscription type. This API is not valid
// to use for unsubscribing a contact at a brand or portal level and will return an
// error.
func (r *SubscriptionService) Unsubscribe(ctx context.Context, body SubscriptionUnsubscribeParams, opts ...option.RequestOption) (res *PublicSubscriptionStatus, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "communication-preferences/v3/unsubscribe"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type PublicSubscriptionStatus struct {
	// The ID for the subscription.
	ID string `json:"id,required"`
	// A description of the subscription.
	Description string `json:"description,required"`
	// The name of the subscription.
	Name string `json:"name,required"`
	// Where the status is determined from e.g. PORTAL_WIDE_STATUS if the contact opted
	// out from the portal.
	//
	// Any of "PORTAL_WIDE_STATUS", "BRAND_WIDE_STATUS", "SUBSCRIPTION_STATUS".
	SourceOfStatus PublicSubscriptionStatusSourceOfStatus `json:"sourceOfStatus,required"`
	// Whether the contact is subscribed.
	//
	// Any of "SUBSCRIBED", "NOT_SUBSCRIBED".
	Status PublicSubscriptionStatusStatus `json:"status,required"`
	// The ID of the brand that the subscription is associated with, if there is one.
	BrandID int64 `json:"brandId"`
	// The legal reason for the current status of the subscription.
	//
	// Any of "LEGITIMATE_INTEREST_PQL", "LEGITIMATE_INTEREST_CLIENT",
	// "PERFORMANCE_OF_CONTRACT", "CONSENT_WITH_NOTICE", "NON_GDPR",
	// "PROCESS_AND_STORE", "LEGITIMATE_INTEREST_OTHER".
	LegalBasis PublicSubscriptionStatusLegalBasis `json:"legalBasis"`
	// A more detailed explanation to go with the legal basis.
	LegalBasisExplanation string `json:"legalBasisExplanation"`
	// The name of the preferences group that the subscription is associated with.
	PreferenceGroupName string `json:"preferenceGroupName"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                    respjson.Field
		Description           respjson.Field
		Name                  respjson.Field
		SourceOfStatus        respjson.Field
		Status                respjson.Field
		BrandID               respjson.Field
		LegalBasis            respjson.Field
		LegalBasisExplanation respjson.Field
		PreferenceGroupName   respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PublicSubscriptionStatus) RawJSON() string { return r.JSON.raw }
func (r *PublicSubscriptionStatus) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Where the status is determined from e.g. PORTAL_WIDE_STATUS if the contact opted
// out from the portal.
type PublicSubscriptionStatusSourceOfStatus string

const (
	PublicSubscriptionStatusSourceOfStatusPortalWideStatus   PublicSubscriptionStatusSourceOfStatus = "PORTAL_WIDE_STATUS"
	PublicSubscriptionStatusSourceOfStatusBrandWideStatus    PublicSubscriptionStatusSourceOfStatus = "BRAND_WIDE_STATUS"
	PublicSubscriptionStatusSourceOfStatusSubscriptionStatus PublicSubscriptionStatusSourceOfStatus = "SUBSCRIPTION_STATUS"
)

// Whether the contact is subscribed.
type PublicSubscriptionStatusStatus string

const (
	PublicSubscriptionStatusStatusSubscribed    PublicSubscriptionStatusStatus = "SUBSCRIBED"
	PublicSubscriptionStatusStatusNotSubscribed PublicSubscriptionStatusStatus = "NOT_SUBSCRIBED"
)

// The legal reason for the current status of the subscription.
type PublicSubscriptionStatusLegalBasis string

const (
	PublicSubscriptionStatusLegalBasisLegitimateInterestPql    PublicSubscriptionStatusLegalBasis = "LEGITIMATE_INTEREST_PQL"
	PublicSubscriptionStatusLegalBasisLegitimateInterestClient PublicSubscriptionStatusLegalBasis = "LEGITIMATE_INTEREST_CLIENT"
	PublicSubscriptionStatusLegalBasisPerformanceOfContract    PublicSubscriptionStatusLegalBasis = "PERFORMANCE_OF_CONTRACT"
	PublicSubscriptionStatusLegalBasisConsentWithNotice        PublicSubscriptionStatusLegalBasis = "CONSENT_WITH_NOTICE"
	PublicSubscriptionStatusLegalBasisNonGdpr                  PublicSubscriptionStatusLegalBasis = "NON_GDPR"
	PublicSubscriptionStatusLegalBasisProcessAndStore          PublicSubscriptionStatusLegalBasis = "PROCESS_AND_STORE"
	PublicSubscriptionStatusLegalBasisLegitimateInterestOther  PublicSubscriptionStatusLegalBasis = "LEGITIMATE_INTEREST_OTHER"
)

type PublicSubscriptionStatusesResponse struct {
	// Email address of the contact.
	Recipient string `json:"recipient,required"`
	// A list of all of the contact's subscriptions statuses.
	SubscriptionStatuses []PublicSubscriptionStatus `json:"subscriptionStatuses,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Recipient            respjson.Field
		SubscriptionStatuses respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PublicSubscriptionStatusesResponse) RawJSON() string { return r.JSON.raw }
func (r *PublicSubscriptionStatusesResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties EmailAddress, SubscriptionID are required.
type PublicUpdateSubscriptionStatusRequestParam struct {
	// Contact's email address.
	EmailAddress string `json:"emailAddress,required"`
	// ID of the subscription being updated for the contact.
	SubscriptionID string `json:"subscriptionId,required"`
	// A more detailed explanation to go with the legal basis (required for GDPR
	// enabled portals).
	LegalBasisExplanation param.Opt[string] `json:"legalBasisExplanation,omitzero"`
	// Legal basis for updating the contact's status (required for GDPR enabled
	// portals).
	//
	// Any of "LEGITIMATE_INTEREST_PQL", "LEGITIMATE_INTEREST_CLIENT",
	// "PERFORMANCE_OF_CONTRACT", "CONSENT_WITH_NOTICE", "NON_GDPR",
	// "PROCESS_AND_STORE", "LEGITIMATE_INTEREST_OTHER".
	LegalBasis PublicUpdateSubscriptionStatusRequestLegalBasis `json:"legalBasis,omitzero"`
	paramObj
}

func (r PublicUpdateSubscriptionStatusRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow PublicUpdateSubscriptionStatusRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PublicUpdateSubscriptionStatusRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Legal basis for updating the contact's status (required for GDPR enabled
// portals).
type PublicUpdateSubscriptionStatusRequestLegalBasis string

const (
	PublicUpdateSubscriptionStatusRequestLegalBasisLegitimateInterestPql    PublicUpdateSubscriptionStatusRequestLegalBasis = "LEGITIMATE_INTEREST_PQL"
	PublicUpdateSubscriptionStatusRequestLegalBasisLegitimateInterestClient PublicUpdateSubscriptionStatusRequestLegalBasis = "LEGITIMATE_INTEREST_CLIENT"
	PublicUpdateSubscriptionStatusRequestLegalBasisPerformanceOfContract    PublicUpdateSubscriptionStatusRequestLegalBasis = "PERFORMANCE_OF_CONTRACT"
	PublicUpdateSubscriptionStatusRequestLegalBasisConsentWithNotice        PublicUpdateSubscriptionStatusRequestLegalBasis = "CONSENT_WITH_NOTICE"
	PublicUpdateSubscriptionStatusRequestLegalBasisNonGdpr                  PublicUpdateSubscriptionStatusRequestLegalBasis = "NON_GDPR"
	PublicUpdateSubscriptionStatusRequestLegalBasisProcessAndStore          PublicUpdateSubscriptionStatusRequestLegalBasis = "PROCESS_AND_STORE"
	PublicUpdateSubscriptionStatusRequestLegalBasisLegitimateInterestOther  PublicUpdateSubscriptionStatusRequestLegalBasis = "LEGITIMATE_INTEREST_OTHER"
)

type SubscriptionDefinition struct {
	// The ID of the definition.
	ID string `json:"id,required"`
	// Time at which the definition was created.
	CreatedAt time.Time `json:"createdAt,required" format:"date-time"`
	// A description of the subscription.
	Description string `json:"description,required"`
	// Whether the definition is active or archived.
	IsActive bool `json:"isActive,required"`
	// A subscription definition created by HubSpot.
	IsDefault bool `json:"isDefault,required"`
	// A default description that is used by some HubSpot tools and cannot be edited.
	IsInternal bool `json:"isInternal,required"`
	// The name of the subscription.
	Name string `json:"name,required"`
	// Time at which the definition was last updated.
	UpdatedAt time.Time `json:"updatedAt,required" format:"date-time"`
	// The ID of the business unit associated with the subscription definition.
	BusinessUnitID int64 `json:"businessUnitId"`
	// The method or technology used to contact.
	CommunicationMethod string `json:"communicationMethod"`
	// The purpose of this subscription or the department in your organization that
	// uses it.
	Purpose string `json:"purpose"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                  respjson.Field
		CreatedAt           respjson.Field
		Description         respjson.Field
		IsActive            respjson.Field
		IsDefault           respjson.Field
		IsInternal          respjson.Field
		Name                respjson.Field
		UpdatedAt           respjson.Field
		BusinessUnitID      respjson.Field
		CommunicationMethod respjson.Field
		Purpose             respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SubscriptionDefinition) RawJSON() string { return r.JSON.raw }
func (r *SubscriptionDefinition) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SubscriptionDefinitionsResponse struct {
	// A list of all subscription definitions.
	SubscriptionDefinitions []SubscriptionDefinition `json:"subscriptionDefinitions,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		SubscriptionDefinitions respjson.Field
		ExtraFields             map[string]respjson.Field
		raw                     string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SubscriptionDefinitionsResponse) RawJSON() string { return r.JSON.raw }
func (r *SubscriptionDefinitionsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SubscriptionSubscribeParams struct {
	PublicUpdateSubscriptionStatusRequest PublicUpdateSubscriptionStatusRequestParam
	paramObj
}

func (r SubscriptionSubscribeParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.PublicUpdateSubscriptionStatusRequest)
}
func (r *SubscriptionSubscribeParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.PublicUpdateSubscriptionStatusRequest)
}

type SubscriptionUnsubscribeParams struct {
	PublicUpdateSubscriptionStatusRequest PublicUpdateSubscriptionStatusRequestParam
	paramObj
}

func (r SubscriptionUnsubscribeParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.PublicUpdateSubscriptionStatusRequest)
}
func (r *SubscriptionUnsubscribeParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.PublicUpdateSubscriptionStatusRequest)
}
