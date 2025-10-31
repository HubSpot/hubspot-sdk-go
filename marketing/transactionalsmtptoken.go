// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package marketing

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/stainless-sdks/hubspot-sdk-go/internal/apiquery"
	shimjson "github.com/stainless-sdks/hubspot-sdk-go/internal/encoding/json"
	"github.com/stainless-sdks/hubspot-sdk-go/internal/requestconfig"
	"github.com/stainless-sdks/hubspot-sdk-go/option"
	"github.com/stainless-sdks/hubspot-sdk-go/packages/pagination"
	"github.com/stainless-sdks/hubspot-sdk-go/packages/param"
)

// TransactionalSmtpTokenService contains methods and other services that help with
// interacting with the Hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTransactionalSmtpTokenService] method instead.
type TransactionalSmtpTokenService struct {
	Options []option.RequestOption
}

// NewTransactionalSmtpTokenService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewTransactionalSmtpTokenService(opts ...option.RequestOption) (r TransactionalSmtpTokenService) {
	r = TransactionalSmtpTokenService{}
	r.Options = opts
	return
}

// Create a SMTP API token.
func (r *TransactionalSmtpTokenService) New(ctx context.Context, body TransactionalSmtpTokenNewParams, opts ...option.RequestOption) (res *SmtpAPITokenView, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "marketing/v3/transactional/smtp-tokens"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Query multiple SMTP API tokens by campaign name or a single token by
// emailCampaignId.
func (r *TransactionalSmtpTokenService) List(ctx context.Context, query TransactionalSmtpTokenListParams, opts ...option.RequestOption) (res *pagination.Page[SmtpAPITokenView], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "marketing/v3/transactional/smtp-tokens"
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
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if tokenID == "" {
		err = errors.New("missing required tokenId parameter")
		return
	}
	path := fmt.Sprintf("marketing/v3/transactional/smtp-tokens/%s", tokenID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Query a single token by ID.
func (r *TransactionalSmtpTokenService) Get(ctx context.Context, tokenID string, opts ...option.RequestOption) (res *SmtpAPITokenView, err error) {
	opts = slices.Concat(r.Options, opts)
	if tokenID == "" {
		err = errors.New("missing required tokenId parameter")
		return
	}
	path := fmt.Sprintf("marketing/v3/transactional/smtp-tokens/%s", tokenID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Allows the creation of a replacement password for a given token. Once the
// password is successfully reset, the old password for the token will be invalid.
func (r *TransactionalSmtpTokenService) ResetPassword(ctx context.Context, tokenID string, opts ...option.RequestOption) (res *SmtpAPITokenView, err error) {
	opts = slices.Concat(r.Options, opts)
	if tokenID == "" {
		err = errors.New("missing required tokenId parameter")
		return
	}
	path := fmt.Sprintf("marketing/v3/transactional/smtp-tokens/%s/password-reset", tokenID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

type TransactionalSmtpTokenNewParams struct {
	// A request object to create a SMTP API token
	SmtpAPITokenRequestEgg SmtpAPITokenRequestEggParam
	paramObj
}

func (r TransactionalSmtpTokenNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.SmtpAPITokenRequestEgg)
}
func (r *TransactionalSmtpTokenNewParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.SmtpAPITokenRequestEgg)
}

type TransactionalSmtpTokenListParams struct {
	// Starting point to get the next set of results.
	After param.Opt[string] `query:"after,omitzero" json:"-"`
	// A name for the campaign tied to the SMTP API token.
	CampaignName param.Opt[string] `query:"campaignName,omitzero" json:"-"`
	// Identifier assigned to the campaign provided during the token creation.
	EmailCampaignID param.Opt[string] `query:"emailCampaignId,omitzero" json:"-"`
	// Maximum number of tokens to return.
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
