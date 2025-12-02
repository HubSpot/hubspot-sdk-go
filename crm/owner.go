// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package crm

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/stainless-sdks/hubspot-sdk-go/internal/apijson"
	"github.com/stainless-sdks/hubspot-sdk-go/internal/apiquery"
	"github.com/stainless-sdks/hubspot-sdk-go/internal/requestconfig"
	"github.com/stainless-sdks/hubspot-sdk-go/option"
	"github.com/stainless-sdks/hubspot-sdk-go/packages/pagination"
	"github.com/stainless-sdks/hubspot-sdk-go/packages/param"
	"github.com/stainless-sdks/hubspot-sdk-go/packages/respjson"
	"github.com/stainless-sdks/hubspot-sdk-go/settings"
	"github.com/stainless-sdks/hubspot-sdk-go/shared"
)

// OwnerService contains methods and other services that help with interacting with
// the hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewOwnerService] method instead.
type OwnerService struct {
	Options []option.RequestOption
}

// NewOwnerService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewOwnerService(opts ...option.RequestOption) (r OwnerService) {
	r = OwnerService{}
	r.Options = opts
	return
}

// Retrieve a paginated list of owners available in the account.
func (r *OwnerService) List(ctx context.Context, query OwnerListParams, opts ...option.RequestOption) (res *pagination.Page[PublicOwner], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "crm/v3/owners/"
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

// Retrieve a paginated list of owners available in the account.
func (r *OwnerService) ListAutoPaging(ctx context.Context, query OwnerListParams, opts ...option.RequestOption) *pagination.PageAutoPager[PublicOwner] {
	return pagination.NewPageAutoPager(r.List(ctx, query, opts...))
}

// Retrieve details of a specific owner using either their 'id' or 'userId'.
func (r *OwnerService) Get(ctx context.Context, ownerID int64, query OwnerGetParams, opts ...option.RequestOption) (res *PublicOwner, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("crm/v3/owners/%v", ownerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type CollectionResponsePublicOwnerForwardPaging struct {
	Results []PublicOwner        `json:"results,required"`
	Paging  shared.ForwardPaging `json:"paging"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Results     respjson.Field
		Paging      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CollectionResponsePublicOwnerForwardPaging) RawJSON() string { return r.JSON.raw }
func (r *CollectionResponsePublicOwnerForwardPaging) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PublicOwner struct {
	// The unique identifier of the owner.
	ID string `json:"id,required"`
	// Indicates whether the owner is archived.
	Archived bool `json:"archived,required"`
	// The date and time when the owner was created.
	CreatedAt time.Time `json:"createdAt,required" format:"date-time"`
	// The type of the owner, which can be either PERSON or QUEUE.
	//
	// Any of "PERSON", "QUEUE".
	Type PublicOwnerType `json:"type,required"`
	// The date and time when the owner was last updated.
	UpdatedAt time.Time `json:"updatedAt,required" format:"date-time"`
	// The email address of the owner.
	Email string `json:"email"`
	// The first name of the owner.
	FirstName string `json:"firstName"`
	// The last name of the owner.
	LastName string                `json:"lastName"`
	Teams    []settings.PublicTeam `json:"teams"`
	// The user ID of the owner.
	UserID int64 `json:"userId"`
	// The user ID of the owner, including inactive users.
	UserIDIncludingInactive int64 `json:"userIdIncludingInactive"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                      respjson.Field
		Archived                respjson.Field
		CreatedAt               respjson.Field
		Type                    respjson.Field
		UpdatedAt               respjson.Field
		Email                   respjson.Field
		FirstName               respjson.Field
		LastName                respjson.Field
		Teams                   respjson.Field
		UserID                  respjson.Field
		UserIDIncludingInactive respjson.Field
		ExtraFields             map[string]respjson.Field
		raw                     string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PublicOwner) RawJSON() string { return r.JSON.raw }
func (r *PublicOwner) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The type of the owner, which can be either PERSON or QUEUE.
type PublicOwnerType string

const (
	PublicOwnerTypePerson PublicOwnerType = "PERSON"
	PublicOwnerTypeQueue  PublicOwnerType = "QUEUE"
)

type OwnerListParams struct {
	After    param.Opt[string] `query:"after,omitzero" json:"-"`
	Archived param.Opt[bool]   `query:"archived,omitzero" json:"-"`
	Email    param.Opt[string] `query:"email,omitzero" json:"-"`
	Limit    param.Opt[int64]  `query:"limit,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [OwnerListParams]'s query parameters as `url.Values`.
func (r OwnerListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type OwnerGetParams struct {
	Archived param.Opt[bool] `query:"archived,omitzero" json:"-"`
	// Any of "id", "userId".
	IDProperty OwnerGetParamsIDProperty `query:"idProperty,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [OwnerGetParams]'s query parameters as `url.Values`.
func (r OwnerGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type OwnerGetParamsIDProperty string

const (
	OwnerGetParamsIDPropertyID     OwnerGetParamsIDProperty = "id"
	OwnerGetParamsIDPropertyUserID OwnerGetParamsIDProperty = "userId"
)
