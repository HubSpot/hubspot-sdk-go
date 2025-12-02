// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package conversations

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/stainless-sdks/hubspot-sdk-go/internal/apiquery"
	shimjson "github.com/stainless-sdks/hubspot-sdk-go/internal/encoding/json"
	"github.com/stainless-sdks/hubspot-sdk-go/internal/requestconfig"
	"github.com/stainless-sdks/hubspot-sdk-go/option"
	"github.com/stainless-sdks/hubspot-sdk-go/packages/param"
	"github.com/stainless-sdks/hubspot-sdk-go/shared"
)

// ActorService contains methods and other services that help with interacting with
// the hubspot API.
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

func (r *ActorService) BatchRead(ctx context.Context, params ActorBatchReadParams, opts ...option.RequestOption) (res *BatchResponsePublicActor, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "conversations/v3/conversations/actors/batch/read"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

func (r *ActorService) Get(ctx context.Context, actorID string, query ActorGetParams, opts ...option.RequestOption) (res *PublicActorUnion, err error) {
	opts = slices.Concat(r.Options, opts)
	if actorID == "" {
		err = errors.New("missing required actorId parameter")
		return
	}
	path := fmt.Sprintf("conversations/v3/conversations/actors/%s", actorID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type ActorBatchReadParams struct {
	// Wrapper for providing an array of strings as inputs.
	BatchInputString shared.BatchInputStringParam
	Property         param.Opt[string] `query:"property,omitzero" json:"-"`
	paramObj
}

func (r ActorBatchReadParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.BatchInputString)
}
func (r *ActorBatchReadParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.BatchInputString)
}

// URLQuery serializes [ActorBatchReadParams]'s query parameters as `url.Values`.
func (r ActorBatchReadParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ActorGetParams struct {
	Property param.Opt[string] `query:"property,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ActorGetParams]'s query parameters as `url.Values`.
func (r ActorGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
