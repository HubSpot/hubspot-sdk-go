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
	"github.com/HubSpot/hubspot-sdk-go/internal/apiquery"
	"github.com/HubSpot/hubspot-sdk-go/internal/requestconfig"
	"github.com/HubSpot/hubspot-sdk-go/option"
	"github.com/HubSpot/hubspot-sdk-go/packages/param"
	"github.com/HubSpot/hubspot-sdk-go/packages/respjson"
)

// SiteSearchService contains methods and other services that help with interacting
// with the hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSiteSearchService] method instead.
type SiteSearchService struct {
	options []option.RequestOption
}

// NewSiteSearchService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewSiteSearchService(opts ...option.RequestOption) (r SiteSearchService) {
	r = SiteSearchService{}
	r.options = opts
	return
}

// Return all indexed data for an asset (e.g., page, blog post, HubDB table),
// specified by ID. This is useful when debugging why a particular asset is not
// returned from a custom search.
func (r *SiteSearchService) GetIndexedData(ctx context.Context, contentID string, query SiteSearchGetIndexedDataParams, opts ...option.RequestOption) (res *IndexedData, err error) {
	opts = slices.Concat(r.options, opts)
	if contentID == "" {
		err = errors.New("missing required contentId parameter")
		return nil, err
	}
	path := fmt.Sprintf("cms/site-search/2026-03/indexed-data/%s", url.PathEscape(contentID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

type IndexedData struct {
	// The ID of the document in HubSpot.
	ID string `json:"id" api:"required"`
	// The indexed fields in HubSpot.
	Fields map[string]IndexedField `json:"fields" api:"required"`
	// The type of document. Can be `SITE_PAGE`, `LANDING_PAGE`, `BLOG_POST`,
	// `LISTING_PAGE`, or `KNOWLEDGE_ARTICLE`.
	//
	// Any of "BLOG_POST", "KNOWLEDGE_ARTICLE", "LANDING_PAGE", "LISTING_PAGE",
	// "SITE_PAGE", "STRUCTURED_CONTENT".
	Type IndexedDataType `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Fields      respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IndexedData) RawJSON() string { return r.JSON.raw }
func (r *IndexedData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The type of document. Can be `SITE_PAGE`, `LANDING_PAGE`, `BLOG_POST`,
// `LISTING_PAGE`, or `KNOWLEDGE_ARTICLE`.
type IndexedDataType string

const (
	IndexedDataTypeBlogPost          IndexedDataType = "BLOG_POST"
	IndexedDataTypeKnowledgeArticle  IndexedDataType = "KNOWLEDGE_ARTICLE"
	IndexedDataTypeLandingPage       IndexedDataType = "LANDING_PAGE"
	IndexedDataTypeListingPage       IndexedDataType = "LISTING_PAGE"
	IndexedDataTypeSitePage          IndexedDataType = "SITE_PAGE"
	IndexedDataTypeStructuredContent IndexedDataType = "STRUCTURED_CONTENT"
)

type IndexedField struct {
	// Indicates whether the field is a metadata field.
	MetadataField bool `json:"metadataField" api:"required"`
	// The name of the indexed field.
	Name string `json:"name" api:"required"`
	// The primary value of the indexed field.
	Value  any   `json:"value" api:"required"`
	Values []any `json:"values" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		MetadataField respjson.Field
		Name          respjson.Field
		Value         respjson.Field
		Values        respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IndexedField) RawJSON() string { return r.JSON.raw }
func (r *IndexedField) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SiteSearchGetIndexedDataParams struct {
	Type param.Opt[string] `query:"type,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SiteSearchGetIndexedDataParams]'s query parameters as
// `url.Values`.
func (r SiteSearchGetIndexedDataParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
