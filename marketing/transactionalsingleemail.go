// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package marketing

import (
	"context"
	"net/http"
	"slices"

	"github.com/HubSpot/hubspot-sdk-go/internal/apijson"
	shimjson "github.com/HubSpot/hubspot-sdk-go/internal/encoding/json"
	"github.com/HubSpot/hubspot-sdk-go/internal/requestconfig"
	"github.com/HubSpot/hubspot-sdk-go/option"
)

// TransactionalSingleEmailService contains methods and other services that help
// with interacting with the hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTransactionalSingleEmailService] method instead.
type TransactionalSingleEmailService struct {
	options []option.RequestOption
}

// NewTransactionalSingleEmailService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewTransactionalSingleEmailService(opts ...option.RequestOption) (r TransactionalSingleEmailService) {
	r = TransactionalSingleEmailService{}
	r.options = opts
	return
}

// Asynchronously send a transactional email. Returns the status of the email send
// with a statusId that can be used to continuously query for the status using the
// Email Send Status API.
func (r *TransactionalSingleEmailService) Send(ctx context.Context, body TransactionalSingleEmailSendParams, opts ...option.RequestOption) (res *EmailSendStatusView, err error) {
	opts = slices.Concat(r.options, opts)
	path := "marketing/transactional/2026-03/single-email/send"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type TransactionalSingleEmailSendParams struct {
	PublicSingleSendRequestEgg PublicSingleSendRequestEggParam
	paramObj
}

func (r TransactionalSingleEmailSendParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.PublicSingleSendRequestEgg)
}
func (r *TransactionalSingleEmailSendParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
