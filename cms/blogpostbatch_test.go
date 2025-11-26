// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cms_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/stainless-sdks/hubspot-sdk-go"
	"github.com/stainless-sdks/hubspot-sdk-go/cms"
	"github.com/stainless-sdks/hubspot-sdk-go/internal/testutil"
	"github.com/stainless-sdks/hubspot-sdk-go/option"
	"github.com/stainless-sdks/hubspot-sdk-go/shared"
)

func TestBlogPostBatchNew(t *testing.T) {
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := hubspotsdk.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAccessToken("pat-na1-xxxxxxxx-xxxx"),
	)
	_, err := client.Cms.Blogs.Posts.Batch.New(context.TODO(), cms.BlogPostBatchNewParams{
		BatchInputBlogPost: cms.BatchInputBlogPostParam{
			Inputs: []cms.BlogPostParam{{
				ID:                  "id",
				AbStatus:            cms.BlogPostAbStatusMaster,
				AbTestID:            "abTestId",
				ArchivedAt:          0,
				ArchivedInDashboard: true,
				AttachedStylesheets: []map[string]any{{
					"foo": map[string]any{},
				}},
				AuthorName:                    "authorName",
				BlogAuthorID:                  "blogAuthorId",
				Campaign:                      "campaign",
				CategoryID:                    0,
				ContentGroupID:                "contentGroupId",
				ContentTypeCategory:           cms.BlogPostContentTypeCategory0,
				Created:                       time.Now(),
				CreatedByID:                   "createdById",
				CurrentlyPublished:            true,
				CurrentState:                  cms.BlogPostCurrentStateAutomated,
				Domain:                        "domain",
				DynamicPageDataSourceID:       "dynamicPageDataSourceId",
				DynamicPageDataSourceType:     0,
				DynamicPageHubDBTableID:       "dynamicPageHubDbTableId",
				EnableDomainStylesheets:       true,
				EnableGoogleAmpOutputOverride: true,
				EnableLayoutStylesheets:       true,
				FeaturedImage:                 "featuredImage",
				FeaturedImageAltText:          "featuredImageAltText",
				FolderID:                      "folderId",
				FooterHTML:                    "footerHtml",
				HeadHTML:                      "headHtml",
				HTMLTitle:                     "htmlTitle",
				IncludeDefaultCustomCss:       true,
				Language:                      cms.BlogPostLanguageAf,
				LayoutSections: map[string]cms.LayoutSectionParam{
					"foo": {
						Cells:    []cms.LayoutSectionParam{},
						CssClass: "cssClass",
						CssID:    "cssId",
						CssStyle: "cssStyle",
						Label:    "label",
						Name:     "name",
						Params: map[string]any{
							"foo": map[string]any{},
						},
						RowMetaData: []cms.RowMetaDataParam{{
							CssClass: "cssClass",
							Styles: cms.StylesParam{
								BackgroundColor: cms.RgbaColorParam{
									A: 0,
									B: 0,
									G: 0,
									R: 0,
								},
								BackgroundGradient: cms.GradientParam{
									Angle: cms.AngleParam{
										Units: "units",
										Value: 0,
									},
									Colors: []cms.ColorStopParam{{
										Color: cms.RgbaColorParam{
											A: 0,
											B: 0,
											G: 0,
											R: 0,
										},
									}},
									SideOrCorner: cms.SideOrCornerParam{
										HorizontalSide: "horizontalSide",
										VerticalSide:   "verticalSide",
									},
								},
								BackgroundImage: cms.BackgroundImageParam{
									BackgroundPosition: "backgroundPosition",
									BackgroundSize:     "backgroundSize",
									ImageURL:           "imageUrl",
								},
								FlexboxPositioning:       "flexboxPositioning",
								ForceFullWidthSection:    true,
								MaxWidthSectionCentering: 0,
								VerticalAlignment:        "verticalAlignment",
								BreakpointStyles: map[string]cms.BreakpointStylesParam{
									"foo": {
										Hidden:  true,
										Margin:  map[string]any{},
										Padding: map[string]any{},
									},
								},
							},
						}},
						Rows: []map[string]cms.LayoutSectionParam{{}},
						Styles: cms.StylesParam{
							BackgroundColor: cms.RgbaColorParam{
								A: 0,
								B: 0,
								G: 0,
								R: 0,
							},
							BackgroundGradient: cms.GradientParam{
								Angle: cms.AngleParam{
									Units: "units",
									Value: 0,
								},
								Colors: []cms.ColorStopParam{{
									Color: cms.RgbaColorParam{
										A: 0,
										B: 0,
										G: 0,
										R: 0,
									},
								}},
								SideOrCorner: cms.SideOrCornerParam{
									HorizontalSide: "horizontalSide",
									VerticalSide:   "verticalSide",
								},
							},
							BackgroundImage: cms.BackgroundImageParam{
								BackgroundPosition: "backgroundPosition",
								BackgroundSize:     "backgroundSize",
								ImageURL:           "imageUrl",
							},
							FlexboxPositioning:       "flexboxPositioning",
							ForceFullWidthSection:    true,
							MaxWidthSectionCentering: 0,
							VerticalAlignment:        "verticalAlignment",
							BreakpointStyles: map[string]cms.BreakpointStylesParam{
								"foo": {
									Hidden:  true,
									Margin:  map[string]any{},
									Padding: map[string]any{},
								},
							},
						},
						Type: "type",
						W:    0,
						X:    0,
					},
				},
				LinkRelCanonicalURL:      "linkRelCanonicalUrl",
				MabExperimentID:          "mabExperimentId",
				MetaDescription:          "metaDescription",
				Name:                     "name",
				PageExpiryDate:           0,
				PageExpiryEnabled:        true,
				PageExpiryRedirectID:     0,
				PageExpiryRedirectURL:    "pageExpiryRedirectUrl",
				Password:                 "password",
				PostBody:                 "postBody",
				PostSummary:              "postSummary",
				PublicAccessRules:        []cms.PublicAccessRule{map[string]any{}},
				PublicAccessRulesEnabled: true,
				PublishDate:              time.Now(),
				PublishImmediately:       true,
				RssBody:                  "rssBody",
				RssSummary:               "rssSummary",
				Slug:                     "slug",
				State:                    "state",
				TagIDs:                   []int64{0},
				ThemeSettingsValues: map[string]any{
					"foo": map[string]any{},
				},
				TranslatedFromID: "translatedFromId",
				Translations: map[string]cms.ContentLanguageVariationParam{
					"foo": {
						ID:                       0,
						ArchivedInDashboard:      true,
						AuthorName:               "authorName",
						Campaign:                 "campaign",
						Created:                  time.Now(),
						Name:                     "name",
						Password:                 "password",
						PublicAccessRules:        []cms.PublicAccessRule{map[string]any{}},
						PublicAccessRulesEnabled: true,
						PublishDate:              time.Now(),
						Slug:                     "slug",
						State:                    "state",
						Updated:                  time.Now(),
						TagIDs:                   []int64{0},
					},
				},
				Updated:          time.Now(),
				UpdatedByID:      "updatedById",
				URL:              "url",
				UseFeaturedImage: true,
				WidgetContainers: map[string]any{
					"foo": map[string]any{},
				},
				Widgets: map[string]any{
					"foo": map[string]any{},
				},
			}},
		},
	})
	if err != nil {
		var apierr *hubspotsdk.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestBlogPostBatchUpdateWithOptionalParams(t *testing.T) {
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := hubspotsdk.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAccessToken("pat-na1-xxxxxxxx-xxxx"),
	)
	_, err := client.Cms.Blogs.Posts.Batch.Update(context.TODO(), cms.BlogPostBatchUpdateParams{
		BatchInputJsonNode: cms.BatchInputJsonNodeParam{
			Inputs: []any{map[string]any{}},
		},
		Archived: hubspotsdk.Bool(true),
	})
	if err != nil {
		var apierr *hubspotsdk.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestBlogPostBatchDelete(t *testing.T) {
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := hubspotsdk.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAccessToken("pat-na1-xxxxxxxx-xxxx"),
	)
	err := client.Cms.Blogs.Posts.Batch.Delete(context.TODO(), cms.BlogPostBatchDeleteParams{
		BatchInputString: shared.BatchInputStringParam{
			Inputs: []string{"string"},
		},
	})
	if err != nil {
		var apierr *hubspotsdk.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestBlogPostBatchGetWithOptionalParams(t *testing.T) {
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := hubspotsdk.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAccessToken("pat-na1-xxxxxxxx-xxxx"),
	)
	_, err := client.Cms.Blogs.Posts.Batch.Get(context.TODO(), cms.BlogPostBatchGetParams{
		BatchInputString: shared.BatchInputStringParam{
			Inputs: []string{"string"},
		},
		Archived: hubspotsdk.Bool(true),
	})
	if err != nil {
		var apierr *hubspotsdk.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
