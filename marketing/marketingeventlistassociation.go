// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package marketing

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/HubSpot/hubspot-sdk-go/internal/requestconfig"
	"github.com/HubSpot/hubspot-sdk-go/option"
)

// MarketingEventListAssociationService contains methods and other services that
// help with interacting with the hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMarketingEventListAssociationService] method instead.
type MarketingEventListAssociationService struct {
	options []option.RequestOption
}

// NewMarketingEventListAssociationService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewMarketingEventListAssociationService(opts ...option.RequestOption) (r MarketingEventListAssociationService) {
	r = MarketingEventListAssociationService{}
	r.options = opts
	return
}

// Gets lists associated with a marketing event by marketing event id
func (r *MarketingEventListAssociationService) List(ctx context.Context, marketingEventID string, opts ...option.RequestOption) (res *CollectionResponseWithTotalPublicList, err error) {
	opts = slices.Concat(r.options, opts)
	if marketingEventID == "" {
		err = errors.New("missing required marketingEventId parameter")
		return nil, err
	}
	path := fmt.Sprintf("marketing/marketing-events/2026-03/associations/%s/lists", url.PathEscape(marketingEventID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Disassociates a list from a marketing event by marketing event id and ILS list
// id
func (r *MarketingEventListAssociationService) Delete(ctx context.Context, listID string, body MarketingEventListAssociationDeleteParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if body.MarketingEventID == "" {
		err = errors.New("missing required marketingEventId parameter")
		return err
	}
	if listID == "" {
		err = errors.New("missing required listId parameter")
		return err
	}
	path := fmt.Sprintf("marketing/marketing-events/2026-03/associations/%s/lists/%s", url.PathEscape(body.MarketingEventID), url.PathEscape(listID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Associates a list with a marketing event by marketing event id and ILS list id
func (r *MarketingEventListAssociationService) Associate(ctx context.Context, listID string, body MarketingEventListAssociationAssociateParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if body.MarketingEventID == "" {
		err = errors.New("missing required marketingEventId parameter")
		return err
	}
	if listID == "" {
		err = errors.New("missing required listId parameter")
		return err
	}
	path := fmt.Sprintf("marketing/marketing-events/2026-03/associations/%s/lists/%s", url.PathEscape(body.MarketingEventID), url.PathEscape(listID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, nil, nil, opts...)
	return err
}

// Associates a list with a marketing event by external account id, external event
// id, and ILS list id
func (r *MarketingEventListAssociationService) AssociateByExternalAccount(ctx context.Context, listID string, body MarketingEventListAssociationAssociateByExternalAccountParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if body.ExternalAccountID == "" {
		err = errors.New("missing required externalAccountId parameter")
		return err
	}
	if body.ExternalEventID == "" {
		err = errors.New("missing required externalEventId parameter")
		return err
	}
	if listID == "" {
		err = errors.New("missing required listId parameter")
		return err
	}
	path := fmt.Sprintf("marketing/marketing-events/2026-03/associations/%s/%s/lists/%s", url.PathEscape(body.ExternalAccountID), url.PathEscape(body.ExternalEventID), url.PathEscape(listID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, nil, nil, opts...)
	return err
}

// Disassociates a list from a marketing event by external account id, external
// event id, and ILS list id
func (r *MarketingEventListAssociationService) DeleteByExternalAccount(ctx context.Context, listID string, body MarketingEventListAssociationDeleteByExternalAccountParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if body.ExternalAccountID == "" {
		err = errors.New("missing required externalAccountId parameter")
		return err
	}
	if body.ExternalEventID == "" {
		err = errors.New("missing required externalEventId parameter")
		return err
	}
	if listID == "" {
		err = errors.New("missing required listId parameter")
		return err
	}
	path := fmt.Sprintf("marketing/marketing-events/2026-03/associations/%s/%s/lists/%s", url.PathEscape(body.ExternalAccountID), url.PathEscape(body.ExternalEventID), url.PathEscape(listID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Gets lists associated with a marketing event by external account id and external
// event id
func (r *MarketingEventListAssociationService) ListByExternalAccount(ctx context.Context, externalEventID string, query MarketingEventListAssociationListByExternalAccountParams, opts ...option.RequestOption) (res *CollectionResponseWithTotalPublicList, err error) {
	opts = slices.Concat(r.options, opts)
	if query.ExternalAccountID == "" {
		err = errors.New("missing required externalAccountId parameter")
		return nil, err
	}
	if externalEventID == "" {
		err = errors.New("missing required externalEventId parameter")
		return nil, err
	}
	path := fmt.Sprintf("marketing/marketing-events/2026-03/associations/%s/%s/lists", url.PathEscape(query.ExternalAccountID), url.PathEscape(externalEventID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type MarketingEventListAssociationDeleteParams struct {
	MarketingEventID string `path:"marketingEventId" api:"required" json:"-"`
	paramObj
}

type MarketingEventListAssociationAssociateParams struct {
	MarketingEventID string `path:"marketingEventId" api:"required" json:"-"`
	paramObj
}

type MarketingEventListAssociationAssociateByExternalAccountParams struct {
	ExternalAccountID string `path:"externalAccountId" api:"required" json:"-"`
	ExternalEventID   string `path:"externalEventId" api:"required" json:"-"`
	paramObj
}

type MarketingEventListAssociationDeleteByExternalAccountParams struct {
	ExternalAccountID string `path:"externalAccountId" api:"required" json:"-"`
	ExternalEventID   string `path:"externalEventId" api:"required" json:"-"`
	paramObj
}

type MarketingEventListAssociationListByExternalAccountParams struct {
	ExternalAccountID string `path:"externalAccountId" api:"required" json:"-"`
	paramObj
}
