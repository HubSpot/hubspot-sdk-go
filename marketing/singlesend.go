// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package marketing

import (
	"context"
	"net/http"
	"slices"

	"github.com/stainless-sdks/hubspot-sdk-go/internal/apijson"
	shimjson "github.com/stainless-sdks/hubspot-sdk-go/internal/encoding/json"
	"github.com/stainless-sdks/hubspot-sdk-go/internal/requestconfig"
	"github.com/stainless-sdks/hubspot-sdk-go/option"
)

// SingleSendService contains methods and other services that help with interacting
// with the hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSingleSendService] method instead.
type SingleSendService struct {
	options []option.RequestOption
}

// NewSingleSendService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewSingleSendService(opts ...option.RequestOption) (r SingleSendService) {
	r = SingleSendService{}
	r.options = opts
	return
}

// Send a template email to a specific recipient.
func (r *SingleSendService) New(ctx context.Context, body SingleSendNewParams, opts ...option.RequestOption) (res *EmailSendStatusView, err error) {
	opts = slices.Concat(r.options, opts)
	path := "marketing/email-campaigns/2026-03/single-send"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type SingleSendNewParams struct {
	PublicSingleSendRequestEgg PublicSingleSendRequestEggParam
	paramObj
}

func (r SingleSendNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.PublicSingleSendRequestEgg)
}
func (r *SingleSendNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
