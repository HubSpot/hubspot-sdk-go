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

// MessageService contains methods and other services that help with interacting
// with the Hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMessageService] method instead.
type MessageService struct {
	Options []option.RequestOption
}

// NewMessageService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewMessageService(opts ...option.RequestOption) (r MessageService) {
	r = MessageService{}
	r.Options = opts
	return
}

// Send a new message on a thread at the current timestamp.
func (r *MessageService) New(ctx context.Context, threadID string, body MessageNewParams, opts ...option.RequestOption) (res *PublicMessageUnion, err error) {
	opts = slices.Concat(r.Options, opts)
	if threadID == "" {
		err = errors.New("missing required threadId parameter")
		return
	}
	path := fmt.Sprintf("conversations/v3/conversations/threads/%s/messages", threadID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieve the message history for a specific thread.
func (r *MessageService) List(ctx context.Context, threadID string, opts ...option.RequestOption) (res *CollectionResponsePublicMessageForwardPaging, err error) {
	opts = slices.Concat(r.Options, opts)
	if threadID == "" {
		err = errors.New("missing required threadId parameter")
		return
	}
	path := fmt.Sprintf("conversations/v3/conversations/threads/%s/messages", threadID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Retrieve a single message from a thread using the message ID.
func (r *MessageService) Get(ctx context.Context, messageID string, query MessageGetParams, opts ...option.RequestOption) (res *PublicMessageUnion, err error) {
	opts = slices.Concat(r.Options, opts)
	if query.ThreadID == "" {
		err = errors.New("missing required threadId parameter")
		return
	}
	if messageID == "" {
		err = errors.New("missing required messageId parameter")
		return
	}
	path := fmt.Sprintf("conversations/v3/conversations/threads/%s/messages/%s", query.ThreadID, messageID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Returns the complete original text and rich text bodies of a message. This will
// be different from the text and rich text in the message itself if the message's
// `truncationStatus` is anything other than `NOT_TRUNCATED`.
func (r *MessageService) GetOriginalContent(ctx context.Context, messageID string, query MessageGetOriginalContentParams, opts ...option.RequestOption) (res *PublicMessageContent, err error) {
	opts = slices.Concat(r.Options, opts)
	if query.ThreadID == "" {
		err = errors.New("missing required threadId parameter")
		return
	}
	if messageID == "" {
		err = errors.New("missing required messageId parameter")
		return
	}
	path := fmt.Sprintf("conversations/v3/conversations/threads/%s/messages/%s/original-content", query.ThreadID, messageID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type MessageNewParams struct {
	PublicMessageEgg PublicMessageEggUnionParam
	paramObj
}

func (r MessageNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.PublicMessageEgg)
}
func (r *MessageNewParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.PublicMessageEgg)
}

type MessageGetParams struct {
	ThreadID string `path:"threadId,required" json:"-"`
	paramObj
}

type MessageGetOriginalContentParams struct {
	ThreadID string `path:"threadId,required" json:"-"`
	paramObj
}
