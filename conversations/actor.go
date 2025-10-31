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
	"github.com/stainless-sdks/hubspot-sdk-go/shared"
)

// ActorService contains methods and other services that help with interacting with
// the Hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewActorService] method instead.
type ActorService struct {
	Options []option.RequestOption
}

// NewActorService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewActorService(opts ...option.RequestOption) (r ActorService) {
	r = ActorService{}
	r.Options = opts
	return
}

// Resolve a set of `ActorId`s to the underlying actors/participants.
func (r *ActorService) BatchRead(ctx context.Context, body ActorBatchReadParams, opts ...option.RequestOption) (res *BatchResponsePublicActor, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "conversations/v3/conversations/actors/batch/read"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieve details of a single actor using the actor ID.
func (r *ActorService) Get(ctx context.Context, actorID string, opts ...option.RequestOption) (res *PublicActorUnion, err error) {
	opts = slices.Concat(r.Options, opts)
	if actorID == "" {
		err = errors.New("missing required actorId parameter")
		return
	}
	path := fmt.Sprintf("conversations/v3/conversations/actors/%s", actorID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type ActorBatchReadParams struct {
	// Wrapper for providing an array of strings as inputs.
	BatchInputString shared.BatchInputStringParam
	paramObj
}

func (r ActorBatchReadParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.BatchInputString)
}
func (r *ActorBatchReadParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.BatchInputString)
}
