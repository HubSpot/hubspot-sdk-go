// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package conversations

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"
	"time"

	"github.com/stainless-sdks/hubspot-sdk-go/internal/apijson"
	"github.com/stainless-sdks/hubspot-sdk-go/internal/requestconfig"
	"github.com/stainless-sdks/hubspot-sdk-go/option"
	"github.com/stainless-sdks/hubspot-sdk-go/packages/param"
	"github.com/stainless-sdks/hubspot-sdk-go/packages/respjson"
	"github.com/stainless-sdks/hubspot-sdk-go/shared"
)

// CustomChannelService contains methods and other services that help with
// interacting with the Hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCustomChannelService] method instead.
type CustomChannelService struct {
	Options                     []option.RequestOption
	ChannelAccountStagingTokens CustomChannelChannelAccountStagingTokenService
	ChannelAccounts             CustomChannelChannelAccountService
	Messages                    CustomChannelMessageService
}

// NewCustomChannelService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCustomChannelService(opts ...option.RequestOption) (r CustomChannelService) {
	r = CustomChannelService{}
	r.Options = opts
	r.ChannelAccountStagingTokens = NewCustomChannelChannelAccountStagingTokenService(opts...)
	r.ChannelAccounts = NewCustomChannelChannelAccountService(opts...)
	r.Messages = NewCustomChannelMessageService(opts...)
	return
}

// Register a new channel along with its capabilities and the webhook url that will
// be used to receive messages published over the channel
func (r *CustomChannelService) New(ctx context.Context, body CustomChannelNewParams, opts ...option.RequestOption) (res *CustomChannelNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "conversations/v3/custom-channels/"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Update the capabilities for an existing. You can also use it to update the
// channel's webhookUri and its channelAccountConnectionRedirectUrl.
func (r *CustomChannelService) Update(ctx context.Context, channelID string, body CustomChannelUpdateParams, opts ...option.RequestOption) (res *CustomChannelUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if channelID == "" {
		err = errors.New("missing required channelId parameter")
		return
	}
	path := fmt.Sprintf("conversations/v3/custom-channels/%s", channelID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Retrieve all custom channels associated with the app.
func (r *CustomChannelService) List(ctx context.Context, opts ...option.RequestOption) (res *CustomChannelListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "conversations/v3/custom-channels/"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Archive an existing registered custom channel
func (r *CustomChannelService) Delete(ctx context.Context, channelID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if channelID == "" {
		err = errors.New("missing required channelId parameter")
		return
	}
	path := fmt.Sprintf("conversations/v3/custom-channels/%s", channelID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Retrieve the details about a custom channel. This API allows you to see a custom
// channel's current capabilties and other configuration metadata
func (r *CustomChannelService) Get(ctx context.Context, channelID string, opts ...option.RequestOption) (res *CustomChannelGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if channelID == "" {
		err = errors.New("missing required channelId parameter")
		return
	}
	path := fmt.Sprintf("conversations/v3/custom-channels/%s", channelID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type CustomChannelNewResponse struct {
	ID                                  string         `json:"id,required"`
	Capabilities                        map[string]any `json:"capabilities,required"`
	CreatedAt                           time.Time      `json:"createdAt,required" format:"date-time"`
	Name                                string         `json:"name,required"`
	ChannelAccountConnectionRedirectURL string         `json:"channelAccountConnectionRedirectUrl"`
	ChannelDescription                  string         `json:"channelDescription"`
	ChannelLogoURL                      string         `json:"channelLogoUrl"`
	WebhookURL                          string         `json:"webhookUrl"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                                  respjson.Field
		Capabilities                        respjson.Field
		CreatedAt                           respjson.Field
		Name                                respjson.Field
		ChannelAccountConnectionRedirectURL respjson.Field
		ChannelDescription                  respjson.Field
		ChannelLogoURL                      respjson.Field
		WebhookURL                          respjson.Field
		ExtraFields                         map[string]respjson.Field
		raw                                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CustomChannelNewResponse) RawJSON() string { return r.JSON.raw }
func (r *CustomChannelNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CustomChannelUpdateResponse struct {
	ID                                  string         `json:"id,required"`
	Capabilities                        map[string]any `json:"capabilities,required"`
	CreatedAt                           time.Time      `json:"createdAt,required" format:"date-time"`
	Name                                string         `json:"name,required"`
	ChannelAccountConnectionRedirectURL string         `json:"channelAccountConnectionRedirectUrl"`
	ChannelDescription                  string         `json:"channelDescription"`
	ChannelLogoURL                      string         `json:"channelLogoUrl"`
	WebhookURL                          string         `json:"webhookUrl"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                                  respjson.Field
		Capabilities                        respjson.Field
		CreatedAt                           respjson.Field
		Name                                respjson.Field
		ChannelAccountConnectionRedirectURL respjson.Field
		ChannelDescription                  respjson.Field
		ChannelLogoURL                      respjson.Field
		WebhookURL                          respjson.Field
		ExtraFields                         map[string]respjson.Field
		raw                                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CustomChannelUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *CustomChannelUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CustomChannelListResponse struct {
	Results []CustomChannelListResponseResult `json:"results,required"`
	Total   int64                             `json:"total,required"`
	Paging  shared.ForwardPaging              `json:"paging"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Results     respjson.Field
		Total       respjson.Field
		Paging      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CustomChannelListResponse) RawJSON() string { return r.JSON.raw }
func (r *CustomChannelListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CustomChannelListResponseResult struct {
	ID                                  string         `json:"id,required"`
	Capabilities                        map[string]any `json:"capabilities,required"`
	CreatedAt                           time.Time      `json:"createdAt,required" format:"date-time"`
	Name                                string         `json:"name,required"`
	ChannelAccountConnectionRedirectURL string         `json:"channelAccountConnectionRedirectUrl"`
	ChannelDescription                  string         `json:"channelDescription"`
	ChannelLogoURL                      string         `json:"channelLogoUrl"`
	WebhookURL                          string         `json:"webhookUrl"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                                  respjson.Field
		Capabilities                        respjson.Field
		CreatedAt                           respjson.Field
		Name                                respjson.Field
		ChannelAccountConnectionRedirectURL respjson.Field
		ChannelDescription                  respjson.Field
		ChannelLogoURL                      respjson.Field
		WebhookURL                          respjson.Field
		ExtraFields                         map[string]respjson.Field
		raw                                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CustomChannelListResponseResult) RawJSON() string { return r.JSON.raw }
func (r *CustomChannelListResponseResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CustomChannelGetResponse struct {
	ID                                  string         `json:"id,required"`
	Capabilities                        map[string]any `json:"capabilities,required"`
	CreatedAt                           time.Time      `json:"createdAt,required" format:"date-time"`
	Name                                string         `json:"name,required"`
	ChannelAccountConnectionRedirectURL string         `json:"channelAccountConnectionRedirectUrl"`
	ChannelDescription                  string         `json:"channelDescription"`
	ChannelLogoURL                      string         `json:"channelLogoUrl"`
	WebhookURL                          string         `json:"webhookUrl"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                                  respjson.Field
		Capabilities                        respjson.Field
		CreatedAt                           respjson.Field
		Name                                respjson.Field
		ChannelAccountConnectionRedirectURL respjson.Field
		ChannelDescription                  respjson.Field
		ChannelLogoURL                      respjson.Field
		WebhookURL                          respjson.Field
		ExtraFields                         map[string]respjson.Field
		raw                                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CustomChannelGetResponse) RawJSON() string { return r.JSON.raw }
func (r *CustomChannelGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CustomChannelNewParams struct {
	Capabilities                        map[string]any    `json:"capabilities,omitzero,required"`
	Name                                string            `json:"name,required"`
	ChannelAccountConnectionRedirectURL param.Opt[string] `json:"channelAccountConnectionRedirectUrl,omitzero"`
	ChannelDescription                  param.Opt[string] `json:"channelDescription,omitzero"`
	ChannelLogoURL                      param.Opt[string] `json:"channelLogoUrl,omitzero"`
	WebhookURL                          param.Opt[string] `json:"webhookUrl,omitzero"`
	paramObj
}

func (r CustomChannelNewParams) MarshalJSON() (data []byte, err error) {
	type shadow CustomChannelNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CustomChannelNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CustomChannelUpdateParams struct {
	Capabilities                        map[string]any `json:"capabilities,omitzero,required"`
	ChannelDescription                  any            `json:"channelDescription,omitzero,required"`
	ChannelLogoURL                      any            `json:"channelLogoUrl,omitzero,required"`
	ChannelAccountConnectionRedirectURL any            `json:"channelAccountConnectionRedirectUrl,omitzero"`
	Name                                any            `json:"name,omitzero"`
	WebhookURL                          any            `json:"webhookUrl,omitzero"`
	paramObj
}

func (r CustomChannelUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow CustomChannelUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CustomChannelUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
