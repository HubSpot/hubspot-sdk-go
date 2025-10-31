// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package automation

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

// SequenceEnrollmentService contains methods and other services that help with
// interacting with the Hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSequenceEnrollmentService] method instead.
type SequenceEnrollmentService struct {
	Options []option.RequestOption
}

// NewSequenceEnrollmentService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewSequenceEnrollmentService(opts ...option.RequestOption) (r SequenceEnrollmentService) {
	r = SequenceEnrollmentService{}
	r.Options = opts
	return
}

// Enroll a contact into a sequence using the specified user ID and sequence
// details.
func (r *SequenceEnrollmentService) Enroll(ctx context.Context, body SequenceEnrollmentEnrollParams, opts ...option.RequestOption) (res *PublicSequenceEnrollmentLiteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "automation/v4/sequences/enrollments"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get the enrollment status of a contact in sequences by their contact ID.
func (r *SequenceEnrollmentService) GetByContactID(ctx context.Context, contactID string, opts ...option.RequestOption) (res *PublicSequenceEnrollmentResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if contactID == "" {
		err = errors.New("missing required contactId parameter")
		return
	}
	path := fmt.Sprintf("automation/v4/sequences/enrollments/contact/%s", contactID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type SequenceEnrollmentEnrollParams struct {
	PublicSequenceEnrollmentRequest PublicSequenceEnrollmentRequestParam
	paramObj
}

func (r SequenceEnrollmentEnrollParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.PublicSequenceEnrollmentRequest)
}
func (r *SequenceEnrollmentEnrollParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.PublicSequenceEnrollmentRequest)
}
