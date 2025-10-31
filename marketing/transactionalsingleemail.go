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

// TransactionalSingleEmailService contains methods and other services that help
// with interacting with the Hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTransactionalSingleEmailService] method instead.
type TransactionalSingleEmailService struct {
	Options []option.RequestOption
}

// NewTransactionalSingleEmailService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewTransactionalSingleEmailService(opts ...option.RequestOption) (r TransactionalSingleEmailService) {
	r = TransactionalSingleEmailService{}
	r.Options = opts
	return
}

// Asynchronously send a transactional email. Returns the status of the email send
// with a statusId that can be used to continuously query for the status using the
// Email Send Status API.
func (r *TransactionalSingleEmailService) Send(ctx context.Context, body TransactionalSingleEmailSendParams, opts ...option.RequestOption) (res *EmailSendStatusView, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "marketing/v3/transactional/single-email/send"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type TransactionalSingleEmailSendParams struct {
	// A request to send a single email asynchronously.
	PublicSingleSendRequestEgg PublicSingleSendRequestEggParam
	paramObj
}

func (r TransactionalSingleEmailSendParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.PublicSingleSendRequestEgg)
}
func (r *TransactionalSingleEmailSendParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.PublicSingleSendRequestEgg)
}
