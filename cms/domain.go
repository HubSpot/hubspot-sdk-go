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

	"github.com/HubSpot/hubspot-sdk-go/internal/apijson"
	"github.com/HubSpot/hubspot-sdk-go/internal/apiquery"
	"github.com/HubSpot/hubspot-sdk-go/internal/requestconfig"
	"github.com/HubSpot/hubspot-sdk-go/option"
	"github.com/HubSpot/hubspot-sdk-go/packages/pagination"
	"github.com/HubSpot/hubspot-sdk-go/packages/param"
	"github.com/HubSpot/hubspot-sdk-go/packages/respjson"
	"github.com/HubSpot/hubspot-sdk-go/shared"
)

// DomainService contains methods and other services that help with interacting
// with the hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDomainService] method instead.
type DomainService struct {
	options []option.RequestOption
}

// NewDomainService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewDomainService(opts ...option.RequestOption) (r DomainService) {
	r = DomainService{}
	r.options = opts
	return
}

func (r *DomainService) List(ctx context.Context, query DomainListParams, opts ...option.RequestOption) (res *pagination.Page[Domain], err error) {
	var raw *http.Response
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "cms/domains/2026-03"
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

func (r *DomainService) ListAutoPaging(ctx context.Context, query DomainListParams, opts ...option.RequestOption) *pagination.PageAutoPager[Domain] {
	return pagination.NewPageAutoPager(r.List(ctx, query, opts...))
}

// Returns a single domains with the id specified.
func (r *DomainService) Get(ctx context.Context, domainID string, opts ...option.RequestOption) (res *Domain, err error) {
	opts = slices.Concat(r.options, opts)
	if domainID == "" {
		err = errors.New("missing required domainId parameter")
		return nil, err
	}
	path := fmt.Sprintf("cms/domains/2026-03/%s", url.PathEscape(domainID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type CollectionResponseWithTotalDomain struct {
	// The results of the query.
	Results []Domain `json:"results" api:"required"`
	// The number of available results.
	Total  int64         `json:"total" api:"required"`
	Paging shared.Paging `json:"paging"`
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
func (r CollectionResponseWithTotalDomain) RawJSON() string { return r.JSON.raw }
func (r *CollectionResponseWithTotalDomain) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Domain struct {
	// The unique ID of this domain.
	ID string `json:"id" api:"required"`
	// The expected CNAME record for the domain.
	CorrectCname string `json:"correctCname" api:"required"`
	// The date and time when the domain was created.
	Created time.Time `json:"created" api:"required" format:"date-time"`
	// The actual domain or sub-domain. e.g. www.hubspot.com
	Domain string `json:"domain" api:"required"`
	// Whether the DNS for this domain is optimally configured for use with HubSpot.
	IsResolving bool `json:"isResolving" api:"required"`
	// Indicates whether SSL is enabled for the domain.
	IsSslEnabled bool `json:"isSslEnabled" api:"required"`
	// Indicates whether the domain is accessible only via SSL.
	IsSslOnly bool `json:"isSslOnly" api:"required"`
	// Whether the domain is used for CMS blog posts.
	IsUsedForBlogPost bool `json:"isUsedForBlogPost" api:"required"`
	// Whether the domain is used for CMS email web pages.
	IsUsedForEmail bool `json:"isUsedForEmail" api:"required"`
	// Whether the domain is used for CMS knowledge pages.
	IsUsedForKnowledge bool `json:"isUsedForKnowledge" api:"required"`
	// Whether the domain is used for CMS landing pages.
	IsUsedForLandingPage bool `json:"isUsedForLandingPage" api:"required"`
	// Whether the domain is used for CMS site pages.
	IsUsedForSitePage bool `json:"isUsedForSitePage" api:"required"`
	// Indicates whether the domain has been manually marked as resolving.
	ManuallyMarkedAsResolving bool `json:"manuallyMarkedAsResolving" api:"required"`
	// Indicates whether the domain is the primary domain for blog posts.
	PrimaryBlogPost bool `json:"primaryBlogPost" api:"required"`
	// Indicates whether the domain is the primary domain for email pages.
	PrimaryEmail bool `json:"primaryEmail" api:"required"`
	// Indicates whether the domain is the primary domain for knowledge pages.
	PrimaryKnowledge bool `json:"primaryKnowledge" api:"required"`
	// Indicates whether the domain is the primary domain for landing pages.
	PrimaryLandingPage bool `json:"primaryLandingPage" api:"required"`
	// Indicates whether the domain is the primary domain for site pages.
	PrimarySitePage bool `json:"primarySitePage" api:"required"`
	// Specifies the domain to which this domain is secondary.
	SecondaryToDomain string `json:"secondaryToDomain" api:"required"`
	// The date and time when the domain was last updated.
	Updated time.Time `json:"updated" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                        respjson.Field
		CorrectCname              respjson.Field
		Created                   respjson.Field
		Domain                    respjson.Field
		IsResolving               respjson.Field
		IsSslEnabled              respjson.Field
		IsSslOnly                 respjson.Field
		IsUsedForBlogPost         respjson.Field
		IsUsedForEmail            respjson.Field
		IsUsedForKnowledge        respjson.Field
		IsUsedForLandingPage      respjson.Field
		IsUsedForSitePage         respjson.Field
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
	Archived      param.Opt[bool]      `query:"archived,omitzero" json:"-"`
	CreatedAfter  param.Opt[time.Time] `query:"createdAfter,omitzero" format:"date-time" json:"-"`
	CreatedAt     param.Opt[time.Time] `query:"createdAt,omitzero" format:"date-time" json:"-"`
	CreatedBefore param.Opt[time.Time] `query:"createdBefore,omitzero" format:"date-time" json:"-"`
	// The maximum number of results to display per page.
	Limit         param.Opt[int64]     `query:"limit,omitzero" json:"-"`
	UpdatedAfter  param.Opt[time.Time] `query:"updatedAfter,omitzero" format:"date-time" json:"-"`
	UpdatedAt     param.Opt[time.Time] `query:"updatedAt,omitzero" format:"date-time" json:"-"`
	UpdatedBefore param.Opt[time.Time] `query:"updatedBefore,omitzero" format:"date-time" json:"-"`
	Sort          []string             `query:"sort,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [DomainListParams]'s query parameters as `url.Values`.
func (r DomainListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
