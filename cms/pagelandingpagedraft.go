// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cms

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/HubSpot/hubspot-sdk-go/internal/apijson"
	shimjson "github.com/HubSpot/hubspot-sdk-go/internal/encoding/json"
	"github.com/HubSpot/hubspot-sdk-go/internal/requestconfig"
	"github.com/HubSpot/hubspot-sdk-go/option"
)

// PageLandingPageDraftService contains methods and other services that help with
// interacting with the hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPageLandingPageDraftService] method instead.
type PageLandingPageDraftService struct {
	options []option.RequestOption
}

// NewPageLandingPageDraftService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewPageLandingPageDraftService(opts ...option.RequestOption) (r PageLandingPageDraftService) {
	r = PageLandingPageDraftService{}
	r.options = opts
	return
}

// Partially updates the draft version of a single landing page, specified by its
// ID. You only need to specify the column values that you are modifying.
func (r *PageLandingPageDraftService) Update(ctx context.Context, objectID string, body PageLandingPageDraftUpdateParams, opts ...option.RequestOption) (res *PagesPage, err error) {
	opts = slices.Concat(r.options, opts)
	if objectID == "" {
		err = errors.New("missing required objectId parameter")
		return nil, err
	}
	path := fmt.Sprintf("cms/pages/2026-03/landing-pages/%s/draft", url.PathEscape(objectID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

// Retrieve the full draft version of a landing page, specified by page ID.
func (r *PageLandingPageDraftService) Get(ctx context.Context, objectID string, opts ...option.RequestOption) (res *PagesPage, err error) {
	opts = slices.Concat(r.options, opts)
	if objectID == "" {
		err = errors.New("missing required objectId parameter")
		return nil, err
	}
	path := fmt.Sprintf("cms/pages/2026-03/landing-pages/%s/draft", url.PathEscape(objectID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Take any changes from the draft version of the Landing Page and apply them to
// the live version.
func (r *PageLandingPageDraftService) PushLive(ctx context.Context, objectID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if objectID == "" {
		err = errors.New("missing required objectId parameter")
		return err
	}
	path := fmt.Sprintf("cms/pages/2026-03/landing-pages/%s/draft/push-live", url.PathEscape(objectID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, nil, opts...)
	return err
}

// Discards any edits and resets the draft to match the live version.
func (r *PageLandingPageDraftService) Reset(ctx context.Context, objectID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if objectID == "" {
		err = errors.New("missing required objectId parameter")
		return err
	}
	path := fmt.Sprintf("cms/pages/2026-03/landing-pages/%s/draft/reset", url.PathEscape(objectID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, nil, opts...)
	return err
}

type PageLandingPageDraftUpdateParams struct {
	PagesPage PagesPageParam
	paramObj
}

func (r PageLandingPageDraftUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.PagesPage)
}
func (r *PageLandingPageDraftUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
