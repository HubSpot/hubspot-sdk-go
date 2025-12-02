// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package crm

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/stainless-sdks/hubspot-sdk-go/internal/requestconfig"
	"github.com/stainless-sdks/hubspot-sdk-go/option"
)

// ObjectLibraryEnablementService contains methods and other services that help
// with interacting with the hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewObjectLibraryEnablementService] method instead.
type ObjectLibraryEnablementService struct {
	Options []option.RequestOption
}

// NewObjectLibraryEnablementService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewObjectLibraryEnablementService(opts ...option.RequestOption) (r ObjectLibraryEnablementService) {
	r = ObjectLibraryEnablementService{}
	r.Options = opts
	return
}

// For all object types supporting enablement, returns whether they're enabled or
// disabled
func (r *ObjectLibraryEnablementService) List(ctx context.Context, opts ...option.RequestOption) (res *PortalObjectTypeEnablementPublicResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "crm/v3/object-library/enablement"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Fetch whether object type is enabled
func (r *ObjectLibraryEnablementService) Get(ctx context.Context, objectTypeID string, opts ...option.RequestOption) (res *ObjectTypeEnablementPublicResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if objectTypeID == "" {
		err = errors.New("missing required objectTypeId parameter")
		return
	}
	path := fmt.Sprintf("crm/v3/object-library/enablement/%s", objectTypeID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}
