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

// CustomChannelMessageService contains methods and other services that help with
// interacting with the hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCustomChannelMessageService] method instead.
type CustomChannelMessageService struct {
	Options []option.RequestOption
}

// NewCustomChannelMessageService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewCustomChannelMessageService(opts ...option.RequestOption) (r CustomChannelMessageService) {
	r = CustomChannelMessageService{}
	r.Options = opts
	return
}

// Publish a message over your custom channel
func (r *CustomChannelMessageService) New(ctx context.Context, channelID int64, body CustomChannelMessageNewParams, opts ...option.RequestOption) (res *ConversationsPublicConversationsMessage, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("conversations/v3/custom-channels/%v/messages", channelID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Update a message's status to indicate if it was successfully sent, failed to
// send, or was read. For failed messages, this can also include the error message
// for the failure.
func (r *CustomChannelMessageService) Update(ctx context.Context, messageID string, params CustomChannelMessageUpdateParams, opts ...option.RequestOption) (res *ConversationsPublicConversationsMessage, err error) {
	opts = slices.Concat(r.Options, opts)
	if messageID == "" {
		err = errors.New("missing required messageId parameter")
		return
	}
	path := fmt.Sprintf("conversations/v3/custom-channels/%v/messages/%s", params.ChannelID, messageID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return
}

// Get the details for a specific message sent over a custom channel
func (r *CustomChannelMessageService) Get(ctx context.Context, messageID string, query CustomChannelMessageGetParams, opts ...option.RequestOption) (res *ConversationsPublicConversationsMessage, err error) {
	opts = slices.Concat(r.Options, opts)
	if messageID == "" {
		err = errors.New("missing required messageId parameter")
		return
	}
	path := fmt.Sprintf("conversations/v3/custom-channels/%v/messages/%s", query.ChannelID, messageID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type CustomChannelMessageNewParams struct {
	ChannelIntegrationMessageEgg ChannelIntegrationMessageEggParam
	paramObj
}

func (r CustomChannelMessageNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.ChannelIntegrationMessageEgg)
}
func (r *CustomChannelMessageNewParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.ChannelIntegrationMessageEgg)
}

type CustomChannelMessageUpdateParams struct {
	ChannelID                                    int64 `path:"channelId,required" json:"-"`
	PublicChannelIntegrationMessageUpdateRequest PublicChannelIntegrationMessageUpdateRequestParam
	paramObj
}

func (r CustomChannelMessageUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.PublicChannelIntegrationMessageUpdateRequest)
}
func (r *CustomChannelMessageUpdateParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.PublicChannelIntegrationMessageUpdateRequest)
}

type CustomChannelMessageGetParams struct {
	ChannelID int64 `path:"channelId,required" json:"-"`
	paramObj
}
