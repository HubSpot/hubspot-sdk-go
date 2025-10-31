// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package marketing

import (
	"context"
	"encoding/json"
	"net/http"
	"slices"

	shimjson "github.com/stainless-sdks/hubspot-sdk-go/internal/encoding/json"
	"github.com/stainless-sdks/hubspot-sdk-go/internal/requestconfig"
	"github.com/stainless-sdks/hubspot-sdk-go/option"
)

// SingleSendService contains methods and other services that help with interacting
// with the Hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSingleSendService] method instead.
type SingleSendService struct {
	Options []option.RequestOption
}

// NewSingleSendService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewSingleSendService(opts ...option.RequestOption) (r SingleSendService) {
	r = SingleSendService{}
	r.Options = opts
	return
}

// Send a template email to a specific recipient.
func (r *SingleSendService) Send(ctx context.Context, body SingleSendSendParams, opts ...option.RequestOption) (res *EmailSendStatusView, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "marketing/v4/email/single-send"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type SingleSendSendParams struct {
	// A request to send a single email asynchronously.
	PublicSingleSendRequestEgg PublicSingleSendRequestEggParam
	paramObj
}

func (r SingleSendSendParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.PublicSingleSendRequestEgg)
}
func (r *SingleSendSendParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.PublicSingleSendRequestEgg)
}
