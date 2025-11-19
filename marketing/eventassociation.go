// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package marketing

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/stainless-sdks/hubspot-sdk-go/internal/requestconfig"
	"github.com/stainless-sdks/hubspot-sdk-go/option"
)

// EventAssociationService contains methods and other services that help with
// interacting with the Hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEventAssociationService] method instead.
type EventAssociationService struct {
	Options []option.RequestOption
}

// NewEventAssociationService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewEventAssociationService(opts ...option.RequestOption) (r EventAssociationService) {
	r = EventAssociationService{}
	r.Options = opts
	return
}

// Gets lists associated with a marketing event by marketing event id
func (r *EventAssociationService) List(ctx context.Context, marketingEventID string, opts ...option.RequestOption) (res *CollectionResponseWithTotalPublicListNoPaging, err error) {
	opts = slices.Concat(r.Options, opts)
	if marketingEventID == "" {
		err = errors.New("missing required marketingEventId parameter")
		return
	}
	path := fmt.Sprintf("marketing/v3/marketing-events/associations/%s/lists", marketingEventID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Disassociates a list from a marketing event by marketing event id and ILS list
// id
func (r *EventAssociationService) Delete(ctx context.Context, listID string, body EventAssociationDeleteParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if body.MarketingEventID == "" {
		err = errors.New("missing required marketingEventId parameter")
		return
	}
	if listID == "" {
		err = errors.New("missing required listId parameter")
		return
	}
	path := fmt.Sprintf("marketing/v3/marketing-events/associations/%s/lists/%s", body.MarketingEventID, listID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Associates a list with a marketing event by marketing event id and ILS list id
func (r *EventAssociationService) Associate(ctx context.Context, listID string, body EventAssociationAssociateParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if body.MarketingEventID == "" {
		err = errors.New("missing required marketingEventId parameter")
		return
	}
	if listID == "" {
		err = errors.New("missing required listId parameter")
		return
	}
	path := fmt.Sprintf("marketing/v3/marketing-events/associations/%s/lists/%s", body.MarketingEventID, listID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, nil, nil, opts...)
	return
}

// Associates a list with a marketing event by external account id, external event
// id, and ILS list id
func (r *EventAssociationService) AssociateByExternalAccount(ctx context.Context, listID string, body EventAssociationAssociateByExternalAccountParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if body.ExternalAccountID == "" {
		err = errors.New("missing required externalAccountId parameter")
		return
	}
	if body.ExternalEventID == "" {
		err = errors.New("missing required externalEventId parameter")
		return
	}
	if listID == "" {
		err = errors.New("missing required listId parameter")
		return
	}
	path := fmt.Sprintf("marketing/v3/marketing-events/associations/%s/%s/lists/%s", body.ExternalAccountID, body.ExternalEventID, listID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, nil, nil, opts...)
	return
}

// Disassociates a list from a marketing event by external account id, external
// event id, and ILS list id
func (r *EventAssociationService) DeleteByExternalAccount(ctx context.Context, listID string, body EventAssociationDeleteByExternalAccountParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if body.ExternalAccountID == "" {
		err = errors.New("missing required externalAccountId parameter")
		return
	}
	if body.ExternalEventID == "" {
		err = errors.New("missing required externalEventId parameter")
		return
	}
	if listID == "" {
		err = errors.New("missing required listId parameter")
		return
	}
	path := fmt.Sprintf("marketing/v3/marketing-events/associations/%s/%s/lists/%s", body.ExternalAccountID, body.ExternalEventID, listID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Gets lists associated with a marketing event by external account id and external
// event id
func (r *EventAssociationService) ListByExternalAccount(ctx context.Context, externalEventID string, query EventAssociationListByExternalAccountParams, opts ...option.RequestOption) (res *CollectionResponseWithTotalPublicListNoPaging, err error) {
	opts = slices.Concat(r.Options, opts)
	if query.ExternalAccountID == "" {
		err = errors.New("missing required externalAccountId parameter")
		return
	}
	if externalEventID == "" {
		err = errors.New("missing required externalEventId parameter")
		return
	}
	path := fmt.Sprintf("marketing/v3/marketing-events/associations/%s/%s/lists", query.ExternalAccountID, externalEventID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type EventAssociationDeleteParams struct {
	MarketingEventID string `path:"marketingEventId,required" json:"-"`
	paramObj
}

type EventAssociationAssociateParams struct {
	MarketingEventID string `path:"marketingEventId,required" json:"-"`
	paramObj
}

type EventAssociationAssociateByExternalAccountParams struct {
	ExternalAccountID string `path:"externalAccountId,required" json:"-"`
	ExternalEventID   string `path:"externalEventId,required" json:"-"`
	paramObj
}

type EventAssociationDeleteByExternalAccountParams struct {
	ExternalAccountID string `path:"externalAccountId,required" json:"-"`
	ExternalEventID   string `path:"externalEventId,required" json:"-"`
	paramObj
}

type EventAssociationListByExternalAccountParams struct {
	ExternalAccountID string `path:"externalAccountId,required" json:"-"`
	paramObj
}
