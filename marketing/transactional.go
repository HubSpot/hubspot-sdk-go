// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package marketing

import (
	"time"

	"github.com/stainless-sdks/hubspot-sdk-go/internal/apijson"
	"github.com/stainless-sdks/hubspot-sdk-go/option"
	"github.com/stainless-sdks/hubspot-sdk-go/packages/param"
	"github.com/stainless-sdks/hubspot-sdk-go/packages/respjson"
	"github.com/stainless-sdks/hubspot-sdk-go/shared"
)

// TransactionalService contains methods and other services that help with
// interacting with the hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTransactionalService] method instead.
type TransactionalService struct {
	options     []option.RequestOption
	SingleEmail TransactionalSingleEmailService
	SmtpTokens  TransactionalSmtpTokenService
}

// NewTransactionalService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewTransactionalService(opts ...option.RequestOption) (r TransactionalService) {
	r = TransactionalService{}
	r.options = opts
	r.SingleEmail = NewTransactionalSingleEmailService(opts...)
	r.SmtpTokens = NewTransactionalSmtpTokenService(opts...)
	return
}

type CollectionResponseSmtpAPITokenViewForwardPaging struct {
	Results []SmtpAPITokenView   `json:"results" api:"required"`
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
func (r CollectionResponseSmtpAPITokenViewForwardPaging) RawJSON() string { return r.JSON.raw }
func (r *CollectionResponseSmtpAPITokenViewForwardPaging) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties CampaignName, CreateContact are required.
type SmtpAPITokenRequestEggParam struct {
	// A name for the campaign tied to the SMTP API token.
	CampaignName string `json:"campaignName" api:"required"`
	// Indicates whether a contact should be created for email recipients.
	CreateContact bool `json:"createContact" api:"required"`
	paramObj
}

func (r SmtpAPITokenRequestEggParam) MarshalJSON() (data []byte, err error) {
	type shadow SmtpAPITokenRequestEggParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SmtpAPITokenRequestEggParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SmtpAPITokenView struct {
	// User name to log into the HubSpot SMTP server.
	ID string `json:"id" api:"required"`
	// A name for the campaign tied to the token.
	CampaignName string `json:"campaignName" api:"required"`
	// Indicates whether a contact should be created for email recipients.
	CreateContact bool `json:"createContact" api:"required"`
	// Timestamp generated when a token is created.
	CreatedAt time.Time `json:"createdAt" api:"required" format:"date-time"`
	// Email address of the user that sent the token creation request.
	CreatedBy string `json:"createdBy" api:"required"`
	// Identifier assigned to the campaign provided in the token creation request.
	EmailCampaignID string `json:"emailCampaignId" api:"required"`
	// Password used to log into the HubSpot SMTP server.
	Password string `json:"password"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		CampaignName    respjson.Field
		CreateContact   respjson.Field
		CreatedAt       respjson.Field
		CreatedBy       respjson.Field
		EmailCampaignID respjson.Field
		Password        respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SmtpAPITokenView) RawJSON() string { return r.JSON.raw }
func (r *SmtpAPITokenView) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
