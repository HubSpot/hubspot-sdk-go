// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package crm

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/stainless-sdks/hubspot-sdk-go/internal/apijson"
	"github.com/stainless-sdks/hubspot-sdk-go/internal/requestconfig"
	"github.com/stainless-sdks/hubspot-sdk-go/option"
	"github.com/stainless-sdks/hubspot-sdk-go/packages/respjson"
)

// ObjectLibraryEnablementService contains methods and other services that help
// with interacting with the Hubspot API.
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

// Returns all objects in the object library and their enablement status
func (r *ObjectLibraryEnablementService) List(ctx context.Context, opts ...option.RequestOption) (res *ObjectLibraryEnablementListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "crm/v3/object-library/enablement"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Returns an object and its enablement status
func (r *ObjectLibraryEnablementService) Get(ctx context.Context, objectTypeID string, opts ...option.RequestOption) (res *ObjectLibraryEnablementGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if objectTypeID == "" {
		err = errors.New("missing required objectTypeId parameter")
		return
	}
	path := fmt.Sprintf("crm/v3/object-library/enablement/%s", objectTypeID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type ObjectLibraryEnablementListResponse struct {
	EnablementByObjectTypeID map[string]bool `json:"enablementByObjectTypeId,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		EnablementByObjectTypeID respjson.Field
		ExtraFields              map[string]respjson.Field
		raw                      string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ObjectLibraryEnablementListResponse) RawJSON() string { return r.JSON.raw }
func (r *ObjectLibraryEnablementListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ObjectLibraryEnablementGetResponse struct {
	Enablement bool `json:"enablement,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Enablement  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ObjectLibraryEnablementGetResponse) RawJSON() string { return r.JSON.raw }
func (r *ObjectLibraryEnablementGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
