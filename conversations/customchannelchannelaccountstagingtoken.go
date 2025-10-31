// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package conversations

import (
	"context"
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
)

// CustomChannelChannelAccountStagingTokenService contains methods and other
// services that help with interacting with the Hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCustomChannelChannelAccountStagingTokenService] method instead.
type CustomChannelChannelAccountStagingTokenService struct {
	Options []option.RequestOption
}

// NewCustomChannelChannelAccountStagingTokenService generates a new service that
// applies the given options to each request. These options are applied after the
// parent client's options (if there is one), and before any request-specific
// options.
func NewCustomChannelChannelAccountStagingTokenService(opts ...option.RequestOption) (r CustomChannelChannelAccountStagingTokenService) {
	r = CustomChannelChannelAccountStagingTokenService{}
	r.Options = opts
	return
}

// Update a channel account staging token's account name and delivery identifier.
// This information will be applied to the channel account created from this
// staging token. This is used for public apps.
func (r *CustomChannelChannelAccountStagingTokenService) Update(ctx context.Context, accountToken string, params CustomChannelChannelAccountStagingTokenUpdateParams, opts ...option.RequestOption) (res *CustomChannelChannelAccountStagingTokenUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.ChannelID == "" {
		err = errors.New("missing required channelId parameter")
		return
	}
	if accountToken == "" {
		err = errors.New("missing required accountToken parameter")
		return
	}
	path := fmt.Sprintf("conversations/v3/custom-channels/%s/channel-account-staging-tokens/%s", params.ChannelID, accountToken)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return
}

type CustomChannelChannelAccountStagingTokenUpdateResponse struct {
	AccountToken       string                   `json:"accountToken,required"`
	CreatedAt          time.Time                `json:"createdAt,required" format:"date-time"`
	GenericChannelID   int64                    `json:"genericChannelId,required"`
	InboxID            int64                    `json:"inboxId,required"`
	UserID             int64                    `json:"userId,required"`
	AccountName        string                   `json:"accountName"`
	DeliveryIdentifier PublicDeliveryIdentifier `json:"deliveryIdentifier"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AccountToken       respjson.Field
		CreatedAt          respjson.Field
		GenericChannelID   respjson.Field
		InboxID            respjson.Field
		UserID             respjson.Field
		AccountName        respjson.Field
		DeliveryIdentifier respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CustomChannelChannelAccountStagingTokenUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *CustomChannelChannelAccountStagingTokenUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CustomChannelChannelAccountStagingTokenUpdateParams struct {
	ChannelID          string                        `path:"channelId,required" json:"-"`
	AccountName        string                        `json:"accountName,required"`
	DeliveryIdentifier PublicDeliveryIdentifierParam `json:"deliveryIdentifier,omitzero,required"`
	paramObj
}

func (r CustomChannelChannelAccountStagingTokenUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow CustomChannelChannelAccountStagingTokenUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CustomChannelChannelAccountStagingTokenUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
