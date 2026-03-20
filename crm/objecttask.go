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

// ObjectTaskService contains methods and other services that help with interacting
// with the hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewObjectTaskService] method instead.
type ObjectTaskService struct {
	Options []option.RequestOption
	Batch   ObjectTaskBatchService
}

// NewObjectTaskService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewObjectTaskService(opts ...option.RequestOption) (r ObjectTaskService) {
	r = ObjectTaskService{}
	r.Options = opts
	r.Batch = NewObjectTaskBatchService(opts...)
	return
}

// Move an Object identified by `{taskId}` to the recycling bin.
func (r *ObjectTaskService) Delete(ctx context.Context, objectID string, body ObjectTaskDeleteParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if body.ObjectType == "" {
		err = errors.New("missing required objectType parameter")
		return err
	}
	if objectID == "" {
		err = errors.New("missing required objectId parameter")
		return err
	}
	path := fmt.Sprintf("crm/objects/2026-03/%s/%s", body.ObjectType, objectID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

type ObjectTaskDeleteParams struct {
	ObjectType string `path:"objectType" api:"required" json:"-"`
	paramObj
}
