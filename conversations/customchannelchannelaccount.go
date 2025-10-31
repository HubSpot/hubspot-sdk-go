// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package conversations

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/stainless-sdks/hubspot-sdk-go/internal/apijson"
	"github.com/stainless-sdks/hubspot-sdk-go/internal/requestconfig"
	"github.com/stainless-sdks/hubspot-sdk-go/option"
	"github.com/stainless-sdks/hubspot-sdk-go/packages/param"
)

// CustomChannelChannelAccountService contains methods and other services that help
// with interacting with the Hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCustomChannelChannelAccountService] method instead.
type CustomChannelChannelAccountService struct {
	Options []option.RequestOption
}

// NewCustomChannelChannelAccountService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewCustomChannelChannelAccountService(opts ...option.RequestOption) (r CustomChannelChannelAccountService) {
	r = CustomChannelChannelAccountService{}
	r.Options = opts
	return
}

// Create a new account for a channel. Multiple accounts can communicate over a
// single channel using different delivery identifiers.
func (r *CustomChannelChannelAccountService) New(ctx context.Context, channelID string, body CustomChannelChannelAccountNewParams, opts ...option.RequestOption) (res *PublicChannelAccount, err error) {
	opts = slices.Concat(r.Options, opts)
	if channelID == "" {
		err = errors.New("missing required channelId parameter")
		return
	}
	path := fmt.Sprintf("conversations/v3/custom-channels/%s/channel-accounts", channelID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// This API is used to update the name of the channel account and it's isAuthorized
// status. Setting to isAuthorized flag to False disables the channel account.
func (r *CustomChannelChannelAccountService) Update(ctx context.Context, channelAccountID string, params CustomChannelChannelAccountUpdateParams, opts ...option.RequestOption) (res *PublicChannelAccount, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.ChannelID == "" {
		err = errors.New("missing required channelId parameter")
		return
	}
	if channelAccountID == "" {
		err = errors.New("missing required channelAccountId parameter")
		return
	}
	path := fmt.Sprintf("conversations/v3/custom-channels/%s/channel-accounts/%s", params.ChannelID, channelAccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return
}

// Retrieve a list of accounts for a custom channel.
func (r *CustomChannelChannelAccountService) List(ctx context.Context, channelID string, opts ...option.RequestOption) (res *CollectionResponseWithTotalPublicChannelAccountForwardPaging, err error) {
	opts = slices.Concat(r.Options, opts)
	if channelID == "" {
		err = errors.New("missing required channelId parameter")
		return
	}
	path := fmt.Sprintf("conversations/v3/custom-channels/%s/channel-accounts", channelID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Retrieve the details for a specific channel account. This contains all the
// metadata about your channel account, including its channel, associated inbox id,
// and delivery identifier information.
func (r *CustomChannelChannelAccountService) Get(ctx context.Context, channelAccountID string, query CustomChannelChannelAccountGetParams, opts ...option.RequestOption) (res *PublicChannelAccount, err error) {
	opts = slices.Concat(r.Options, opts)
	if query.ChannelID == "" {
		err = errors.New("missing required channelId parameter")
		return
	}
	if channelAccountID == "" {
		err = errors.New("missing required channelAccountId parameter")
		return
	}
	path := fmt.Sprintf("conversations/v3/custom-channels/%s/channel-accounts/%s", query.ChannelID, channelAccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type CustomChannelChannelAccountNewParams struct {
	Authorized         bool                          `json:"authorized,required"`
	InboxID            string                        `json:"inboxId,required"`
	Name               string                        `json:"name,required"`
	DeliveryIdentifier PublicDeliveryIdentifierParam `json:"deliveryIdentifier,omitzero"`
	paramObj
}

func (r CustomChannelChannelAccountNewParams) MarshalJSON() (data []byte, err error) {
	type shadow CustomChannelChannelAccountNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CustomChannelChannelAccountNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CustomChannelChannelAccountUpdateParams struct {
	ChannelID  string            `path:"channelId,required" json:"-"`
	Authorized param.Opt[bool]   `json:"authorized,omitzero"`
	Name       param.Opt[string] `json:"name,omitzero"`
	paramObj
}

func (r CustomChannelChannelAccountUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow CustomChannelChannelAccountUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CustomChannelChannelAccountUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CustomChannelChannelAccountGetParams struct {
	ChannelID string `path:"channelId,required" json:"-"`
	paramObj
}
