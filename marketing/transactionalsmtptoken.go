// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package marketing

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/stainless-sdks/hubspot-sdk-go/internal/apijson"
	"github.com/stainless-sdks/hubspot-sdk-go/internal/apiquery"
	shimjson "github.com/stainless-sdks/hubspot-sdk-go/internal/encoding/json"
	"github.com/stainless-sdks/hubspot-sdk-go/internal/requestconfig"
	"github.com/stainless-sdks/hubspot-sdk-go/option"
	"github.com/stainless-sdks/hubspot-sdk-go/packages/pagination"
	"github.com/stainless-sdks/hubspot-sdk-go/packages/param"
)

// TransactionalSmtpTokenService contains methods and other services that help with
// interacting with the hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTransactionalSmtpTokenService] method instead.
type TransactionalSmtpTokenService struct {
	options []option.RequestOption
}

// NewTransactionalSmtpTokenService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewTransactionalSmtpTokenService(opts ...option.RequestOption) (r TransactionalSmtpTokenService) {
	r = TransactionalSmtpTokenService{}
	r.options = opts
	return
}

// Create a SMTP API token.
func (r *TransactionalSmtpTokenService) New(ctx context.Context, body TransactionalSmtpTokenNewParams, opts ...option.RequestOption) (res *SmtpAPITokenView, err error) {
	opts = slices.Concat(r.options, opts)
	path := "marketing/transactional/2026-03/smtp-tokens"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Query multiple SMTP API tokens by campaign name or a single token by
// emailCampaignId.
func (r *TransactionalSmtpTokenService) List(ctx context.Context, query TransactionalSmtpTokenListParams, opts ...option.RequestOption) (res *pagination.Page[SmtpAPITokenView], err error) {
	var raw *http.Response
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "marketing/transactional/2026-03/smtp-tokens"
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// Query multiple SMTP API tokens by campaign name or a single token by
// emailCampaignId.
func (r *TransactionalSmtpTokenService) ListAutoPaging(ctx context.Context, query TransactionalSmtpTokenListParams, opts ...option.RequestOption) *pagination.PageAutoPager[SmtpAPITokenView] {
	return pagination.NewPageAutoPager(r.List(ctx, query, opts...))
}

// Delete a single token by ID.
func (r *TransactionalSmtpTokenService) Delete(ctx context.Context, tokenID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if tokenID == "" {
		err = errors.New("missing required tokenId parameter")
		return err
	}
	path := fmt.Sprintf("marketing/transactional/2026-03/smtp-tokens/%s", url.PathEscape(tokenID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Query a single token by ID.
func (r *TransactionalSmtpTokenService) Get(ctx context.Context, tokenID string, opts ...option.RequestOption) (res *SmtpAPITokenView, err error) {
	opts = slices.Concat(r.options, opts)
	if tokenID == "" {
		err = errors.New("missing required tokenId parameter")
		return nil, err
	}
	path := fmt.Sprintf("marketing/transactional/2026-03/smtp-tokens/%s", url.PathEscape(tokenID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Allows the creation of a replacement password for a given token. Once the
// password is successfully reset, the old password for the token will be invalid.
func (r *TransactionalSmtpTokenService) ResetPassword(ctx context.Context, tokenID string, opts ...option.RequestOption) (res *SmtpAPITokenView, err error) {
	opts = slices.Concat(r.options, opts)
	if tokenID == "" {
		err = errors.New("missing required tokenId parameter")
		return nil, err
	}
	path := fmt.Sprintf("marketing/transactional/2026-03/smtp-tokens/%s/password-reset", url.PathEscape(tokenID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

type TransactionalSmtpTokenNewParams struct {
	SmtpAPITokenRequestEgg SmtpAPITokenRequestEggParam
	paramObj
}

func (r TransactionalSmtpTokenNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.SmtpAPITokenRequestEgg)
}
func (r *TransactionalSmtpTokenNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TransactionalSmtpTokenListParams struct {
	// The paging cursor token of the last successfully read resource will be returned
	// as the `paging.next.after` JSON property of a paged response containing more
	// results.
	After           param.Opt[string] `query:"after,omitzero" json:"-"`
	CampaignName    param.Opt[string] `query:"campaignName,omitzero" json:"-"`
	EmailCampaignID param.Opt[string] `query:"emailCampaignId,omitzero" json:"-"`
	// The maximum number of results to display per page.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [TransactionalSmtpTokenListParams]'s query parameters as
// `url.Values`.
func (r TransactionalSmtpTokenListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
