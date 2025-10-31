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

// InboxService contains methods and other services that help with interacting with
// the Hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewInboxService] method instead.
type InboxService struct {
	Options []option.RequestOption
}

// NewInboxService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewInboxService(opts ...option.RequestOption) (r InboxService) {
	r = InboxService{}
	r.Options = opts
	return
}

// Retrieve a list of conversations inboxes, with optional filters and sorting.
func (r *InboxService) List(ctx context.Context, opts ...option.RequestOption) (res *CollectionResponseWithTotalPublicInboxForwardPaging, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "conversations/v3/conversations/inboxes"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Retrieve details of a single conversations inbox using the inbox ID.
func (r *InboxService) Get(ctx context.Context, inboxID string, opts ...option.RequestOption) (res *PublicInbox, err error) {
	opts = slices.Concat(r.Options, opts)
	if inboxID == "" {
		err = errors.New("missing required inboxId parameter")
		return
	}
	path := fmt.Sprintf("conversations/v3/conversations/inboxes/%s", inboxID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}
