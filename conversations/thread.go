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

// ThreadService contains methods and other services that help with interacting
// with the Hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewThreadService] method instead.
type ThreadService struct {
	Options []option.RequestOption
}

// NewThreadService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewThreadService(opts ...option.RequestOption) (r ThreadService) {
	r = ThreadService{}
	r.Options = opts
	return
}

// Updates a single thread. Either a thread's status can be updated, or the thread
// can be restored.
func (r *ThreadService) Update(ctx context.Context, threadID string, body ThreadUpdateParams, opts ...option.RequestOption) (res *PublicThread, err error) {
	opts = slices.Concat(r.Options, opts)
	if threadID == "" {
		err = errors.New("missing required threadId parameter")
		return
	}
	path := fmt.Sprintf("conversations/v3/conversations/threads/%s", threadID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Retrieve a list of threads, with optional filters and sorting.
func (r *ThreadService) List(ctx context.Context, opts ...option.RequestOption) (res *CollectionResponsePublicThreadForwardPaging, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "conversations/v3/conversations/threads"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Archives a single thread. The thread will be permanently deleted 30 days after
// placed in an archived state.
func (r *ThreadService) Delete(ctx context.Context, threadID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if threadID == "" {
		err = errors.New("missing required threadId parameter")
		return
	}
	path := fmt.Sprintf("conversations/v3/conversations/threads/%s", threadID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Retrieve a single thread by its ID
func (r *ThreadService) Get(ctx context.Context, threadID string, opts ...option.RequestOption) (res *PublicThread, err error) {
	opts = slices.Concat(r.Options, opts)
	if threadID == "" {
		err = errors.New("missing required threadId parameter")
		return
	}
	path := fmt.Sprintf("conversations/v3/conversations/threads/%s", threadID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type ThreadUpdateParams struct {
	PublicThreadUpdateRequest PublicThreadUpdateRequestParam
	paramObj
}

func (r ThreadUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.PublicThreadUpdateRequest)
}
func (r *ThreadUpdateParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.PublicThreadUpdateRequest)
}
