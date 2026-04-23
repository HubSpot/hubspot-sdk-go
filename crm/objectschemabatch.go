// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package crm

import (
	"context"
	"net/http"
	"slices"

	"github.com/stainless-sdks/hubspot-sdk-go/internal/apijson"
	shimjson "github.com/stainless-sdks/hubspot-sdk-go/internal/encoding/json"
	"github.com/stainless-sdks/hubspot-sdk-go/internal/requestconfig"
	"github.com/stainless-sdks/hubspot-sdk-go/option"
)

// ObjectSchemaBatchService contains methods and other services that help with
// interacting with the hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewObjectSchemaBatchService] method instead.
type ObjectSchemaBatchService struct {
	options []option.RequestOption
}

// NewObjectSchemaBatchService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewObjectSchemaBatchService(opts ...option.RequestOption) (r ObjectSchemaBatchService) {
	r = ObjectSchemaBatchService{}
	r.options = opts
	return
}

// Retrieve details of multiple custom object schemas by providing a batch request
// with specified inputs. This operation allows you to fetch schema information,
// including properties and associations, for multiple custom objects in a single
// API call.
func (r *ObjectSchemaBatchService) Get(ctx context.Context, body ObjectSchemaBatchGetParams, opts ...option.RequestOption) (res *CollectionResponseObjectSchemaNoPaging, err error) {
	opts = slices.Concat(r.options, opts)
	path := "crm-object-schemas/2026-03/schemas/batch/read"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type ObjectSchemaBatchGetParams struct {
	ObjectSchemaBatchReadRequest ObjectSchemaBatchReadRequestParam
	paramObj
}

func (r ObjectSchemaBatchGetParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.ObjectSchemaBatchReadRequest)
}
func (r *ObjectSchemaBatchGetParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
