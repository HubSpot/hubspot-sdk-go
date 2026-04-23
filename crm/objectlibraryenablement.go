// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package crm

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/stainless-sdks/hubspot-sdk-go/internal/apijson"
	"github.com/stainless-sdks/hubspot-sdk-go/internal/requestconfig"
	"github.com/stainless-sdks/hubspot-sdk-go/option"
	"github.com/stainless-sdks/hubspot-sdk-go/packages/respjson"
)

// ObjectLibraryEnablementService contains methods and other services that help
// with interacting with the hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewObjectLibraryEnablementService] method instead.
type ObjectLibraryEnablementService struct {
	options []option.RequestOption
}

// NewObjectLibraryEnablementService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewObjectLibraryEnablementService(opts ...option.RequestOption) (r ObjectLibraryEnablementService) {
	r = ObjectLibraryEnablementService{}
	r.options = opts
	return
}

func (r *ObjectLibraryEnablementService) GetAll(ctx context.Context, opts ...option.RequestOption) (res *PortalObjectTypeEnablementPublicResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "crm/object-library/2026-03/enablement"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

func (r *ObjectLibraryEnablementService) GetByObjectTypeID(ctx context.Context, objectTypeID string, opts ...option.RequestOption) (res *ObjectTypeEnablementPublicResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if objectTypeID == "" {
		err = errors.New("missing required objectTypeId parameter")
		return nil, err
	}
	path := fmt.Sprintf("crm/object-library/2026-03/enablement/%s", url.PathEscape(objectTypeID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type ObjectTypeEnablementPublicResponse struct {
	// Whether the object type is enabled or not
	Enablement bool `json:"enablement" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Enablement  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ObjectTypeEnablementPublicResponse) RawJSON() string { return r.JSON.raw }
func (r *ObjectTypeEnablementPublicResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PortalObjectTypeEnablementPublicResponse struct {
	// A map of objectTypeId to whether that object type is enabled or not
	EnablementByObjectTypeID map[string]bool `json:"enablementByObjectTypeId" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		EnablementByObjectTypeID respjson.Field
		ExtraFields              map[string]respjson.Field
		raw                      string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PortalObjectTypeEnablementPublicResponse) RawJSON() string { return r.JSON.raw }
func (r *PortalObjectTypeEnablementPublicResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
