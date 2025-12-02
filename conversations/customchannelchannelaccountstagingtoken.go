// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package conversations

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"slices"

	shimjson "github.com/stainless-sdks/hubspot-sdk-go/internal/encoding/json"
	"github.com/stainless-sdks/hubspot-sdk-go/internal/requestconfig"
	"github.com/stainless-sdks/hubspot-sdk-go/option"
)

// CustomChannelChannelAccountStagingTokenService contains methods and other
// services that help with interacting with the hubspot API.
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
func (r *CustomChannelChannelAccountStagingTokenService) Update(ctx context.Context, accountToken string, params CustomChannelChannelAccountStagingTokenUpdateParams, opts ...option.RequestOption) (res *PublicChannelAccountStagingToken, err error) {
	opts = slices.Concat(r.Options, opts)
	if accountToken == "" {
		err = errors.New("missing required accountToken parameter")
		return
	}
	path := fmt.Sprintf("conversations/v3/custom-channels/%v/channel-account-staging-tokens/%s", params.ChannelID, accountToken)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return
}

type CustomChannelChannelAccountStagingTokenUpdateParams struct {
	ChannelID                                     int64 `path:"channelId,required" json:"-"`
	PublicChannelAccountStagingTokenUpdateRequest PublicChannelAccountStagingTokenUpdateRequestParam
	paramObj
}

func (r CustomChannelChannelAccountStagingTokenUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.PublicChannelAccountStagingTokenUpdateRequest)
}
func (r *CustomChannelChannelAccountStagingTokenUpdateParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.PublicChannelAccountStagingTokenUpdateRequest)
}
