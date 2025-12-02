// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cms

import (
	"context"
	"errors"
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
	"github.com/stainless-sdks/hubspot-sdk-go/shared"
)

// DomainService contains methods and other services that help with interacting
// with the hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDomainService] method instead.
type DomainService struct {
	Options []option.RequestOption
}

// NewDomainService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewDomainService(opts ...option.RequestOption) (r DomainService) {
	r = DomainService{}
	r.Options = opts
	return
}

// Returns all existing domains that have been created. Results can be limited and
// filtered by creation or updated date.
func (r *DomainService) List(ctx context.Context, query DomainListParams, opts ...option.RequestOption) (res *pagination.Page[Domain], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "cms/v3/domains/"
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

// Returns all existing domains that have been created. Results can be limited and
// filtered by creation or updated date.
func (r *DomainService) ListAutoPaging(ctx context.Context, query DomainListParams, opts ...option.RequestOption) *pagination.PageAutoPager[Domain] {
	return pagination.NewPageAutoPager(r.List(ctx, query, opts...))
}

// Returns a single domains with the id specified.
func (r *DomainService) Get(ctx context.Context, domainID string, opts ...option.RequestOption) (res *Domain, err error) {
	opts = slices.Concat(r.Options, opts)
	if domainID == "" {
		err = errors.New("missing required domainId parameter")
		return
	}
	path := fmt.Sprintf("cms/v3/domains/%s", domainID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type CollectionResponseWithTotalDomainForwardPaging struct {
	Results []Domain             `json:"results,required"`
	Total   int64                `json:"total,required"`
	Paging  shared.ForwardPaging `json:"paging"`
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
func (r CollectionResponseWithTotalDomainForwardPaging) RawJSON() string { return r.JSON.raw }
func (r *CollectionResponseWithTotalDomainForwardPaging) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Domain struct {
	// The unique ID of this domain.
	ID string `json:"id,required"`
	// The actual domain or sub-domain. e.g. www.hubspot.com
	Domain string `json:"domain,required"`
	// Whether the DNS for this domain is optimally configured for use with HubSpot.
	IsResolving bool `json:"isResolving,required"`
	// Whether the domain is used for CMS blog posts.
	IsUsedForBlogPost bool `json:"isUsedForBlogPost,required"`
	// Whether the domain is used for CMS email web pages.
	IsUsedForEmail bool `json:"isUsedForEmail,required"`
	// Whether the domain is used for CMS knowledge pages.
	IsUsedForKnowledge bool `json:"isUsedForKnowledge,required"`
	// Whether the domain is used for CMS landing pages.
	IsUsedForLandingPage bool `json:"isUsedForLandingPage,required"`
	// Whether the domain is used for CMS site pages.
	IsUsedForSitePage         bool      `json:"isUsedForSitePage,required"`
	CorrectCname              string    `json:"correctCname"`
	Created                   time.Time `json:"created" format:"date-time"`
	IsSslEnabled              bool      `json:"isSslEnabled"`
	IsSslOnly                 bool      `json:"isSslOnly"`
	ManuallyMarkedAsResolving bool      `json:"manuallyMarkedAsResolving"`
	PrimaryBlogPost           bool      `json:"primaryBlogPost"`
	PrimaryEmail              bool      `json:"primaryEmail"`
	PrimaryKnowledge          bool      `json:"primaryKnowledge"`
	PrimaryLandingPage        bool      `json:"primaryLandingPage"`
	PrimarySitePage           bool      `json:"primarySitePage"`
	SecondaryToDomain         string    `json:"secondaryToDomain"`
	Updated                   time.Time `json:"updated" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                        respjson.Field
		Domain                    respjson.Field
		IsResolving               respjson.Field
		IsUsedForBlogPost         respjson.Field
		IsUsedForEmail            respjson.Field
		IsUsedForKnowledge        respjson.Field
		IsUsedForLandingPage      respjson.Field
		IsUsedForSitePage         respjson.Field
		CorrectCname              respjson.Field
		Created                   respjson.Field
		IsSslEnabled              respjson.Field
		IsSslOnly                 respjson.Field
		ManuallyMarkedAsResolving respjson.Field
		PrimaryBlogPost           respjson.Field
		PrimaryEmail              respjson.Field
		PrimaryKnowledge          respjson.Field
		PrimaryLandingPage        respjson.Field
		PrimarySitePage           respjson.Field
		SecondaryToDomain         respjson.Field
		Updated                   respjson.Field
		ExtraFields               map[string]respjson.Field
		raw                       string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Domain) RawJSON() string { return r.JSON.raw }
func (r *Domain) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DomainListParams struct {
	// The paging cursor token of the last successfully read resource will be returned
	// as the `paging.next.after` JSON property of a paged response containing more
	// results.
	After param.Opt[string] `query:"after,omitzero" json:"-"`
	// Whether to return only results that have been archived.
	Archived param.Opt[bool] `query:"archived,omitzero" json:"-"`
	// Only return domains created after this date.
	CreatedAfter param.Opt[time.Time] `query:"createdAfter,omitzero" format:"date-time" json:"-"`
	// Only return domains created at this date.
	CreatedAt param.Opt[time.Time] `query:"createdAt,omitzero" format:"date-time" json:"-"`
	// Only return domains created before this date.
	CreatedBefore param.Opt[time.Time] `query:"createdBefore,omitzero" format:"date-time" json:"-"`
	// Maximum number of results per page.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Only return domains updated after this date.
	UpdatedAfter param.Opt[time.Time] `query:"updatedAfter,omitzero" format:"date-time" json:"-"`
	// Only return domains updated at this date.
	UpdatedAt param.Opt[time.Time] `query:"updatedAt,omitzero" format:"date-time" json:"-"`
	// Only return domains updated before this date.
	UpdatedBefore param.Opt[time.Time] `query:"updatedBefore,omitzero" format:"date-time" json:"-"`
	// Specifies the order in which the domains are returned.
	Sort []string `query:"sort,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [DomainListParams]'s query parameters as `url.Values`.
func (r DomainListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
