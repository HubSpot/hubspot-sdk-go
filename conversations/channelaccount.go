// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package conversations

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/stainless-sdks/hubspot-sdk-go/internal/requestconfig"
	"github.com/stainless-sdks/hubspot-sdk-go/option"
)

// ChannelAccountService contains methods and other services that help with
// interacting with the Hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewChannelAccountService] method instead.
type ChannelAccountService struct {
	Options []option.RequestOption
}

// NewChannelAccountService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewChannelAccountService(opts ...option.RequestOption) (r ChannelAccountService) {
	r = ChannelAccountService{}
	r.Options = opts
	return
}

// Retrieve a list of channel accounts, with optional filters and sorting.
func (r *ChannelAccountService) List(ctx context.Context, opts ...option.RequestOption) (res *CollectionResponseWithTotalPublicChannelAccountForwardPaging, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "conversations/v3/conversations/channel-accounts"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Retrieve details of a single channel account using the channel account ID.
func (r *ChannelAccountService) Get(ctx context.Context, channelAccountID string, opts ...option.RequestOption) (res *PublicChannelAccount, err error) {
	opts = slices.Concat(r.Options, opts)
	if channelAccountID == "" {
		err = errors.New("missing required channelAccountId parameter")
		return
	}
	path := fmt.Sprintf("conversations/v3/conversations/channel-accounts/%s", channelAccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}
