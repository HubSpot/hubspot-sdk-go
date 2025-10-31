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

// ChannelService contains methods and other services that help with interacting
// with the Hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewChannelService] method instead.
type ChannelService struct {
	Options []option.RequestOption
}

// NewChannelService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewChannelService(opts ...option.RequestOption) (r ChannelService) {
	r = ChannelService{}
	r.Options = opts
	return
}

// Retrieve a list of channels, with optional filters and sorting.
func (r *ChannelService) List(ctx context.Context, opts ...option.RequestOption) (res *CollectionResponseWithTotalPublicChannelForwardPaging, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "conversations/v3/conversations/channels"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Retrieve details of a single channel using the channel ID.
func (r *ChannelService) Get(ctx context.Context, channelID string, opts ...option.RequestOption) (res *PublicChannel, err error) {
	opts = slices.Concat(r.Options, opts)
	if channelID == "" {
		err = errors.New("missing required channelId parameter")
		return
	}
	path := fmt.Sprintf("conversations/v3/conversations/channels/%s", channelID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}
