// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cms_test

import (
	"bytes"
	"context"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/stainless-sdks/hubspot-sdk-go"
	"github.com/stainless-sdks/hubspot-sdk-go/cms"
	"github.com/stainless-sdks/hubspot-sdk-go/internal/testutil"
	"github.com/stainless-sdks/hubspot-sdk-go/option"
)

func TestBlogPostBatchNew(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte("abc"))
	}))
	defer server.Close()
	baseURL := server.URL
	client := hubspotsdk.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAccessToken("pat-na1-xxxxxxxx-xxxx"),
	)
	resp, err := client.Cms.Blogs.Posts.Batch.New(context.TODO(), cms.BlogPostBatchNewParams{
		BatchInputBlogPost: cms.BatchInputBlogPostParam{
			Inputs: []cms.BlogPostParam{{
				ID:                  "id",
				AbStatus:            cms.BlogPostAbStatusAutomatedLoserVariant,
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
				CurrentState:                  cms.BlogPostCurrentStateAgentGenerated,
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
				Language:                      cms.BlogPostLanguageAa,
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
										Units: cms.AngleUnitsDeg,
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
										HorizontalSide: cms.SideOrCornerHorizontalSideCenter,
										VerticalSide:   cms.SideOrCornerVerticalSideBottom,
									},
								},
								BackgroundImage: cms.BackgroundImageParam{
									BackgroundPosition: "backgroundPosition",
									BackgroundSize:     "backgroundSize",
									ImageURL:           "imageUrl",
								},
								FlexboxPositioning:       cms.StylesFlexboxPositioningBottomCenter,
								ForceFullWidthSection:    true,
								MaxWidthSectionCentering: 0,
								VerticalAlignment:        cms.StylesVerticalAlignmentBottom,
								BreakpointStyles: map[string]cms.BreakpointStylesParam{
									"foo": {
										Hidden: true,
										Margin: cms.MarginParam{
											Bottom: cms.SizeParam{
												Units: cms.SizeUnitsUnknown0,
												Value: 0,
											},
											Top: cms.SizeParam{
												Units: cms.SizeUnitsUnknown0,
												Value: 0,
											},
										},
										Padding: cms.PaddingParam{
											Bottom: cms.SizeParam{
												Units: cms.SizeUnitsUnknown0,
												Value: 0,
											},
											Left: cms.SizeParam{
												Units: cms.SizeUnitsUnknown0,
												Value: 0,
											},
											Right: cms.SizeParam{
												Units: cms.SizeUnitsUnknown0,
												Value: 0,
											},
											Top: cms.SizeParam{
												Units: cms.SizeUnitsUnknown0,
												Value: 0,
											},
										},
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
									Units: cms.AngleUnitsDeg,
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
									HorizontalSide: cms.SideOrCornerHorizontalSideCenter,
									VerticalSide:   cms.SideOrCornerVerticalSideBottom,
								},
							},
							BackgroundImage: cms.BackgroundImageParam{
								BackgroundPosition: "backgroundPosition",
								BackgroundSize:     "backgroundSize",
								ImageURL:           "imageUrl",
							},
							FlexboxPositioning:       cms.StylesFlexboxPositioningBottomCenter,
							ForceFullWidthSection:    true,
							MaxWidthSectionCentering: 0,
							VerticalAlignment:        cms.StylesVerticalAlignmentBottom,
							BreakpointStyles: map[string]cms.BreakpointStylesParam{
								"foo": {
									Hidden: true,
									Margin: cms.MarginParam{
										Bottom: cms.SizeParam{
											Units: cms.SizeUnitsUnknown0,
											Value: 0,
										},
										Top: cms.SizeParam{
											Units: cms.SizeUnitsUnknown0,
											Value: 0,
										},
									},
									Padding: cms.PaddingParam{
										Bottom: cms.SizeParam{
											Units: cms.SizeUnitsUnknown0,
											Value: 0,
										},
										Left: cms.SizeParam{
											Units: cms.SizeUnitsUnknown0,
											Value: 0,
										},
										Right: cms.SizeParam{
											Units: cms.SizeUnitsUnknown0,
											Value: 0,
										},
										Top: cms.SizeParam{
											Units: cms.SizeUnitsUnknown0,
											Value: 0,
										},
									},
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
						CampaignName:             "campaignName",
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
	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		var apierr *hubspotsdk.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
	if !bytes.Equal(b, []byte("abc")) {
		t.Fatalf("return value not %s: %s", "abc", b)
	}
}

func TestBlogPostBatchUpdateWithOptionalParams(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte("abc"))
	}))
	defer server.Close()
	baseURL := server.URL
	client := hubspotsdk.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAccessToken("pat-na1-xxxxxxxx-xxxx"),
	)
	resp, err := client.Cms.Blogs.Posts.Batch.Update(context.TODO(), cms.BlogPostBatchUpdateParams{
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
	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		var apierr *hubspotsdk.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
	if !bytes.Equal(b, []byte("abc")) {
		t.Fatalf("return value not %s: %s", "abc", b)
	}
}

func TestBlogPostBatchDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
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
		BatchInputString: cms.BatchInputStringParam{
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
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte("abc"))
	}))
	defer server.Close()
	baseURL := server.URL
	client := hubspotsdk.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAccessToken("pat-na1-xxxxxxxx-xxxx"),
	)
	resp, err := client.Cms.Blogs.Posts.Batch.Get(context.TODO(), cms.BlogPostBatchGetParams{
		BatchInputString: cms.BatchInputStringParam{
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
	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		var apierr *hubspotsdk.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
	if !bytes.Equal(b, []byte("abc")) {
		t.Fatalf("return value not %s: %s", "abc", b)
	}
}
