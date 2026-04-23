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

func TestBlogPostNew(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte("abc"))
	}))
	defer server.Close()
	baseURL := server.URL
	client := hubspotsdk.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAccessToken("My Access Token"),
	)
	resp, err := client.Cms.Blogs.Posts.New(context.TODO(), cms.BlogPostNewParams{
		BlogPost: cms.BlogPostParam{
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
											Units: cms.SizeUnitsPercent,
											Value: 0,
										},
										Top: cms.SizeParam{
											Units: cms.SizeUnitsPercent,
											Value: 0,
										},
									},
									Padding: cms.PaddingParam{
										Bottom: cms.SizeParam{
											Units: cms.SizeUnitsPercent,
											Value: 0,
										},
										Left: cms.SizeParam{
											Units: cms.SizeUnitsPercent,
											Value: 0,
										},
										Right: cms.SizeParam{
											Units: cms.SizeUnitsPercent,
											Value: 0,
										},
										Top: cms.SizeParam{
											Units: cms.SizeUnitsPercent,
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
										Units: cms.SizeUnitsPercent,
										Value: 0,
									},
									Top: cms.SizeParam{
										Units: cms.SizeUnitsPercent,
										Value: 0,
									},
								},
								Padding: cms.PaddingParam{
									Bottom: cms.SizeParam{
										Units: cms.SizeUnitsPercent,
										Value: 0,
									},
									Left: cms.SizeParam{
										Units: cms.SizeUnitsPercent,
										Value: 0,
									},
									Right: cms.SizeParam{
										Units: cms.SizeUnitsPercent,
										Value: 0,
									},
									Top: cms.SizeParam{
										Units: cms.SizeUnitsPercent,
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

func TestBlogPostUpdateWithOptionalParams(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte("abc"))
	}))
	defer server.Close()
	baseURL := server.URL
	client := hubspotsdk.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAccessToken("My Access Token"),
	)
	resp, err := client.Cms.Blogs.Posts.Update(
		context.TODO(),
		"objectId",
		cms.BlogPostUpdateParams{
			BlogPost: cms.BlogPostParam{
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
												Units: cms.SizeUnitsPercent,
												Value: 0,
											},
											Top: cms.SizeParam{
												Units: cms.SizeUnitsPercent,
												Value: 0,
											},
										},
										Padding: cms.PaddingParam{
											Bottom: cms.SizeParam{
												Units: cms.SizeUnitsPercent,
												Value: 0,
											},
											Left: cms.SizeParam{
												Units: cms.SizeUnitsPercent,
												Value: 0,
											},
											Right: cms.SizeParam{
												Units: cms.SizeUnitsPercent,
												Value: 0,
											},
											Top: cms.SizeParam{
												Units: cms.SizeUnitsPercent,
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
											Units: cms.SizeUnitsPercent,
											Value: 0,
										},
										Top: cms.SizeParam{
											Units: cms.SizeUnitsPercent,
											Value: 0,
										},
									},
									Padding: cms.PaddingParam{
										Bottom: cms.SizeParam{
											Units: cms.SizeUnitsPercent,
											Value: 0,
										},
										Left: cms.SizeParam{
											Units: cms.SizeUnitsPercent,
											Value: 0,
										},
										Right: cms.SizeParam{
											Units: cms.SizeUnitsPercent,
											Value: 0,
										},
										Top: cms.SizeParam{
											Units: cms.SizeUnitsPercent,
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
			},
			Archived: hubspotsdk.Bool(true),
		},
	)
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

func TestBlogPostListWithOptionalParams(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte("abc"))
	}))
	defer server.Close()
	baseURL := server.URL
	client := hubspotsdk.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAccessToken("My Access Token"),
	)
	resp, err := client.Cms.Blogs.Posts.List(context.TODO(), cms.BlogPostListParams{
		After:         hubspotsdk.String("after"),
		Archived:      hubspotsdk.Bool(true),
		CreatedAfter:  hubspotsdk.Time(time.Now()),
		CreatedAt:     hubspotsdk.Time(time.Now()),
		CreatedBefore: hubspotsdk.Time(time.Now()),
		Limit:         hubspotsdk.Int(0),
		Property:      hubspotsdk.String("property"),
		Sort:          []string{"string"},
		UpdatedAfter:  hubspotsdk.Time(time.Now()),
		UpdatedAt:     hubspotsdk.Time(time.Now()),
		UpdatedBefore: hubspotsdk.Time(time.Now()),
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

func TestBlogPostDeleteWithOptionalParams(t *testing.T) {
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
		option.WithAccessToken("My Access Token"),
	)
	err := client.Cms.Blogs.Posts.Delete(
		context.TODO(),
		"objectId",
		cms.BlogPostDeleteParams{
			Archived: hubspotsdk.Bool(true),
		},
	)
	if err != nil {
		var apierr *hubspotsdk.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestBlogPostCloneWithOptionalParams(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte("abc"))
	}))
	defer server.Close()
	baseURL := server.URL
	client := hubspotsdk.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAccessToken("My Access Token"),
	)
	resp, err := client.Cms.Blogs.Posts.Clone(context.TODO(), cms.BlogPostCloneParams{
		ContentCloneRequestVNext: cms.ContentCloneRequestVNextParam{
			ID:        "id",
			CloneName: hubspotsdk.String("cloneName"),
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

func TestBlogPostGetWithOptionalParams(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte("abc"))
	}))
	defer server.Close()
	baseURL := server.URL
	client := hubspotsdk.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAccessToken("My Access Token"),
	)
	resp, err := client.Cms.Blogs.Posts.Get(
		context.TODO(),
		"objectId",
		cms.BlogPostGetParams{
			Archived: hubspotsdk.Bool(true),
			Property: hubspotsdk.String("property"),
		},
	)
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

func TestBlogPostGetDraftByID(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte("abc"))
	}))
	defer server.Close()
	baseURL := server.URL
	client := hubspotsdk.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAccessToken("My Access Token"),
	)
	resp, err := client.Cms.Blogs.Posts.GetDraftByID(context.TODO(), "objectId")
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

func TestBlogPostListAuthorsWithOptionalParams(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte("abc"))
	}))
	defer server.Close()
	baseURL := server.URL
	client := hubspotsdk.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAccessToken("My Access Token"),
	)
	resp, err := client.Cms.Blogs.Posts.ListAuthors(context.TODO(), cms.BlogPostListAuthorsParams{
		After:         hubspotsdk.String("after"),
		Archived:      hubspotsdk.Bool(true),
		CreatedAfter:  hubspotsdk.Time(time.Now()),
		CreatedAt:     hubspotsdk.Time(time.Now()),
		CreatedBefore: hubspotsdk.Time(time.Now()),
		Limit:         hubspotsdk.Int(0),
		Property:      hubspotsdk.String("property"),
		Sort:          []string{"string"},
		UpdatedAfter:  hubspotsdk.Time(time.Now()),
		UpdatedAt:     hubspotsdk.Time(time.Now()),
		UpdatedBefore: hubspotsdk.Time(time.Now()),
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

func TestBlogPostListTagsWithOptionalParams(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte("abc"))
	}))
	defer server.Close()
	baseURL := server.URL
	client := hubspotsdk.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAccessToken("My Access Token"),
	)
	resp, err := client.Cms.Blogs.Posts.ListTags(context.TODO(), cms.BlogPostListTagsParams{
		After:         hubspotsdk.String("after"),
		Archived:      hubspotsdk.Bool(true),
		CreatedAfter:  hubspotsdk.Time(time.Now()),
		CreatedAt:     hubspotsdk.Time(time.Now()),
		CreatedBefore: hubspotsdk.Time(time.Now()),
		Limit:         hubspotsdk.Int(0),
		Property:      hubspotsdk.String("property"),
		Sort:          []string{"string"},
		UpdatedAfter:  hubspotsdk.Time(time.Now()),
		UpdatedAt:     hubspotsdk.Time(time.Now()),
		UpdatedBefore: hubspotsdk.Time(time.Now()),
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

func TestBlogPostPushLive(t *testing.T) {
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
		option.WithAccessToken("My Access Token"),
	)
	err := client.Cms.Blogs.Posts.PushLive(context.TODO(), "objectId")
	if err != nil {
		var apierr *hubspotsdk.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestBlogPostQueryWithOptionalParams(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte("abc"))
	}))
	defer server.Close()
	baseURL := server.URL
	client := hubspotsdk.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAccessToken("My Access Token"),
	)
	resp, err := client.Cms.Blogs.Posts.Query(context.TODO(), cms.BlogPostQueryParams{
		After:         hubspotsdk.String("after"),
		Archived:      hubspotsdk.Bool(true),
		CreatedAfter:  hubspotsdk.Time(time.Now()),
		CreatedAt:     hubspotsdk.Time(time.Now()),
		CreatedBefore: hubspotsdk.Time(time.Now()),
		Limit:         hubspotsdk.Int(0),
		Property:      hubspotsdk.String("property"),
		Sort:          []string{"string"},
		UpdatedAfter:  hubspotsdk.Time(time.Now()),
		UpdatedAt:     hubspotsdk.Time(time.Now()),
		UpdatedBefore: hubspotsdk.Time(time.Now()),
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

func TestBlogPostQueryAuthorsWithOptionalParams(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte("abc"))
	}))
	defer server.Close()
	baseURL := server.URL
	client := hubspotsdk.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAccessToken("My Access Token"),
	)
	resp, err := client.Cms.Blogs.Posts.QueryAuthors(context.TODO(), cms.BlogPostQueryAuthorsParams{
		After:         hubspotsdk.String("after"),
		Archived:      hubspotsdk.Bool(true),
		CreatedAfter:  hubspotsdk.Time(time.Now()),
		CreatedAt:     hubspotsdk.Time(time.Now()),
		CreatedBefore: hubspotsdk.Time(time.Now()),
		Limit:         hubspotsdk.Int(0),
		Property:      hubspotsdk.String("property"),
		Sort:          []string{"string"},
		UpdatedAfter:  hubspotsdk.Time(time.Now()),
		UpdatedAt:     hubspotsdk.Time(time.Now()),
		UpdatedBefore: hubspotsdk.Time(time.Now()),
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

func TestBlogPostQueryTagsWithOptionalParams(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte("abc"))
	}))
	defer server.Close()
	baseURL := server.URL
	client := hubspotsdk.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAccessToken("My Access Token"),
	)
	resp, err := client.Cms.Blogs.Posts.QueryTags(context.TODO(), cms.BlogPostQueryTagsParams{
		After:         hubspotsdk.String("after"),
		Archived:      hubspotsdk.Bool(true),
		CreatedAfter:  hubspotsdk.Time(time.Now()),
		CreatedAt:     hubspotsdk.Time(time.Now()),
		CreatedBefore: hubspotsdk.Time(time.Now()),
		Limit:         hubspotsdk.Int(0),
		Property:      hubspotsdk.String("property"),
		Sort:          []string{"string"},
		UpdatedAfter:  hubspotsdk.Time(time.Now()),
		UpdatedAt:     hubspotsdk.Time(time.Now()),
		UpdatedBefore: hubspotsdk.Time(time.Now()),
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

func TestBlogPostResetDraft(t *testing.T) {
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
		option.WithAccessToken("My Access Token"),
	)
	err := client.Cms.Blogs.Posts.ResetDraft(context.TODO(), "objectId")
	if err != nil {
		var apierr *hubspotsdk.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestBlogPostSchedule(t *testing.T) {
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
		option.WithAccessToken("My Access Token"),
	)
	err := client.Cms.Blogs.Posts.Schedule(context.TODO(), cms.BlogPostScheduleParams{
		ContentScheduleRequestVNext: cms.ContentScheduleRequestVNextParam{
			ID:          "id",
			PublishDate: time.Now(),
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

func TestBlogPostUpdateDraft(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte("abc"))
	}))
	defer server.Close()
	baseURL := server.URL
	client := hubspotsdk.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAccessToken("My Access Token"),
	)
	resp, err := client.Cms.Blogs.Posts.UpdateDraft(
		context.TODO(),
		"objectId",
		cms.BlogPostUpdateDraftParams{
			BlogPost: cms.BlogPostParam{
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
												Units: cms.SizeUnitsPercent,
												Value: 0,
											},
											Top: cms.SizeParam{
												Units: cms.SizeUnitsPercent,
												Value: 0,
											},
										},
										Padding: cms.PaddingParam{
											Bottom: cms.SizeParam{
												Units: cms.SizeUnitsPercent,
												Value: 0,
											},
											Left: cms.SizeParam{
												Units: cms.SizeUnitsPercent,
												Value: 0,
											},
											Right: cms.SizeParam{
												Units: cms.SizeUnitsPercent,
												Value: 0,
											},
											Top: cms.SizeParam{
												Units: cms.SizeUnitsPercent,
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
											Units: cms.SizeUnitsPercent,
											Value: 0,
										},
										Top: cms.SizeParam{
											Units: cms.SizeUnitsPercent,
											Value: 0,
										},
									},
									Padding: cms.PaddingParam{
										Bottom: cms.SizeParam{
											Units: cms.SizeUnitsPercent,
											Value: 0,
										},
										Left: cms.SizeParam{
											Units: cms.SizeUnitsPercent,
											Value: 0,
										},
										Right: cms.SizeParam{
											Units: cms.SizeUnitsPercent,
											Value: 0,
										},
										Top: cms.SizeParam{
											Units: cms.SizeUnitsPercent,
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
			},
		},
	)
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
