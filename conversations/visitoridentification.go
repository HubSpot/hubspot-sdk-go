// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package conversations

import (
	"context"
	"net/http"
	"slices"

	"github.com/stainless-sdks/hubspot-sdk-go/internal/apijson"
	shimjson "github.com/stainless-sdks/hubspot-sdk-go/internal/encoding/json"
	"github.com/stainless-sdks/hubspot-sdk-go/internal/requestconfig"
	"github.com/stainless-sdks/hubspot-sdk-go/option"
	"github.com/stainless-sdks/hubspot-sdk-go/packages/param"
	"github.com/stainless-sdks/hubspot-sdk-go/packages/respjson"
)

// VisitorIdentificationService contains methods and other services that help with
// interacting with the hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewVisitorIdentificationService] method instead.
type VisitorIdentificationService struct {
	options []option.RequestOption
}

// NewVisitorIdentificationService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewVisitorIdentificationService(opts ...option.RequestOption) (r VisitorIdentificationService) {
	r = VisitorIdentificationService{}
	r.options = opts
	return
}

// Generate an identification token for a website visitor who has been
// authenticated using your own system. An identification token returned from this
// API can be used to pass information about your already-authenticated visitor to
// the chat widget, so that it treats the visitor as a known contact. This allows
// support agents to recognize and assist the visitor more effectively.
func (r *VisitorIdentificationService) GenerateToken(ctx context.Context, body VisitorIdentificationGenerateTokenParams, opts ...option.RequestOption) (res *IdentificationTokenResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "visitor-identification/2026-03/tokens/create"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// The property Email is required.
type IdentificationTokenGenerationRequestParam struct {
	// The email of the visitor that you wish to identify
	Email string `json:"email" api:"required"`
	// The first name of the visitor that you wish to identify. This value will only be
	// set in HubSpot for new contacts and existing contacts where first name is
	// unknown. Optional.
	FirstName param.Opt[string] `json:"firstName,omitzero"`
	// The last name of the visitor that you wish to identify. This value will only be
	// set in HubSpot for new contacts and existing contacts where last name is
	// unknown. Optional.
	LastName param.Opt[string] `json:"lastName,omitzero"`
	paramObj
}

func (r IdentificationTokenGenerationRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow IdentificationTokenGenerationRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *IdentificationTokenGenerationRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IdentificationTokenResponse struct {
	// An identification token that allows the visitor to be treated as a known
	// contact.
	Token string `json:"token" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Token       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IdentificationTokenResponse) RawJSON() string { return r.JSON.raw }
func (r *IdentificationTokenResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VisitorIdentificationGenerateTokenParams struct {
	IdentificationTokenGenerationRequest IdentificationTokenGenerationRequestParam
	paramObj
}

func (r VisitorIdentificationGenerateTokenParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.IdentificationTokenGenerationRequest)
}
func (r *VisitorIdentificationGenerateTokenParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
