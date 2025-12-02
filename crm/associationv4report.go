// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package crm

import (
	"context"
	"fmt"
	"net/http"
	"slices"

	"github.com/stainless-sdks/hubspot-sdk-go/internal/requestconfig"
	"github.com/stainless-sdks/hubspot-sdk-go/option"
)

// AssociationV4ReportService contains methods and other services that help with
// interacting with the hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAssociationV4ReportService] method instead.
type AssociationV4ReportService struct {
	Options []option.RequestOption
}

// NewAssociationV4ReportService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAssociationV4ReportService(opts ...option.RequestOption) (r AssociationV4ReportService) {
	r = AssociationV4ReportService{}
	r.Options = opts
	return
}

// Requests a report of all objects in the portal which have a high usage of
// associations
func (r *AssociationV4ReportService) RequestHighUsageReport(ctx context.Context, userID int64, opts ...option.RequestOption) (res *ReportCreationResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("crm/v4/associations/usage/high-usage-report/%v", userID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}
