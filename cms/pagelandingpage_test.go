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

func TestPageLandingPageNew(t *testing.T) {
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
	err := client.Cms.Pages.LandingPages.New(context.TODO(), cms.PageLandingPageNewParams{
		Page: cms.PageParam{
			ID:                  "id",
			AbStatus:            cms.PageAbStatusAutomatedLoserVariant,
			AbTestID:            "abTestId",
			ArchivedAt:          time.Now(),
			ArchivedInDashboard: true,
			AttachedStylesheets: []map[string]any{{
				"foo": map[string]any{},
			}},
			AuthorName:                "authorName",
			Campaign:                  "campaign",
			CategoryID:                0,
			ContentGroupID:            "contentGroupId",
			ContentTypeCategory:       cms.PageContentTypeCategory0,
			Created:                   time.Now(),
			CreatedByID:               "createdById",
			CurrentlyPublished:        true,
			CurrentState:              cms.PageCurrentStateAutomated,
			Domain:                    "domain",
			DynamicPageDataSourceID:   "dynamicPageDataSourceId",
			DynamicPageDataSourceType: 0,
			DynamicPageHubDBTableID:   "dynamicPageHubDbTableId",
			EnableDomainStylesheets:   true,
			EnableLayoutStylesheets:   true,
			FeaturedImage:             "featuredImage",
			FeaturedImageAltText:      "featuredImageAltText",
			FolderID:                  "folderId",
			FooterHTML:                "footerHtml",
			HeadHTML:                  "headHtml",
			HTMLTitle:                 "htmlTitle",
			IncludeDefaultCustomCss:   true,
			Language:                  cms.PageLanguageAf,
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
			PageRedirected:           true,
			Password:                 "password",
			PublicAccessRules:        []cms.PublicAccessRule{map[string]any{}},
			PublicAccessRulesEnabled: true,
			PublishDate:              time.Now(),
			PublishImmediately:       true,
			Slug:                     "slug",
			State:                    "state",
			Subcategory:              "subcategory",
			TemplatePath:             "templatePath",
			ThemeSettingsValues: map[string]any{
				"foo": map[string]any{},
			},
			TranslatedFromID: "translatedFromId",
			Translations: map[string]cms.PagesContentLanguageVariationParam{
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

func TestPageLandingPageUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Cms.Pages.LandingPages.Update(
		context.TODO(),
		"objectId",
		cms.PageLandingPageUpdateParams{
			Page: cms.PageParam{
				ID:                  "id",
				AbStatus:            cms.PageAbStatusAutomatedLoserVariant,
				AbTestID:            "abTestId",
				ArchivedAt:          time.Now(),
				ArchivedInDashboard: true,
				AttachedStylesheets: []map[string]any{{
					"foo": map[string]any{},
				}},
				AuthorName:                "authorName",
				Campaign:                  "campaign",
				CategoryID:                0,
				ContentGroupID:            "contentGroupId",
				ContentTypeCategory:       cms.PageContentTypeCategory0,
				Created:                   time.Now(),
				CreatedByID:               "createdById",
				CurrentlyPublished:        true,
				CurrentState:              cms.PageCurrentStateAutomated,
				Domain:                    "domain",
				DynamicPageDataSourceID:   "dynamicPageDataSourceId",
				DynamicPageDataSourceType: 0,
				DynamicPageHubDBTableID:   "dynamicPageHubDbTableId",
				EnableDomainStylesheets:   true,
				EnableLayoutStylesheets:   true,
				FeaturedImage:             "featuredImage",
				FeaturedImageAltText:      "featuredImageAltText",
				FolderID:                  "folderId",
				FooterHTML:                "footerHtml",
				HeadHTML:                  "headHtml",
				HTMLTitle:                 "htmlTitle",
				IncludeDefaultCustomCss:   true,
				Language:                  cms.PageLanguageAf,
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
				PageRedirected:           true,
				Password:                 "password",
				PublicAccessRules:        []cms.PublicAccessRule{map[string]any{}},
				PublicAccessRulesEnabled: true,
				PublishDate:              time.Now(),
				PublishImmediately:       true,
				Slug:                     "slug",
				State:                    "state",
				Subcategory:              "subcategory",
				TemplatePath:             "templatePath",
				ThemeSettingsValues: map[string]any{
					"foo": map[string]any{},
				},
				TranslatedFromID: "translatedFromId",
				Translations: map[string]cms.PagesContentLanguageVariationParam{
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
}

func TestPageLandingPageListWithOptionalParams(t *testing.T) {
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
	_, err := client.Cms.Pages.LandingPages.List(context.TODO(), cms.PageLandingPageListParams{
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
}

func TestPageLandingPageDeleteWithOptionalParams(t *testing.T) {
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
	err := client.Cms.Pages.LandingPages.Delete(
		context.TODO(),
		"objectId",
		cms.PageLandingPageDeleteParams{
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

func TestPageLandingPageAttachToLangGroupWithOptionalParams(t *testing.T) {
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
	err := client.Cms.Pages.LandingPages.AttachToLangGroup(context.TODO(), cms.PageLandingPageAttachToLangGroupParams{
		AttachToLangPrimaryRequestVNext: cms.AttachToLangPrimaryRequestVNextParam{
			ID:              "id",
			Language:        "language",
			PrimaryID:       "primaryId",
			PrimaryLanguage: hubspotsdk.String("primaryLanguage"),
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

func TestPageLandingPageCloneWithOptionalParams(t *testing.T) {
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
	_, err := client.Cms.Pages.LandingPages.Clone(context.TODO(), cms.PageLandingPageCloneParams{
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
}

func TestPageLandingPageNewAbTestVariation(t *testing.T) {
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
	_, err := client.Cms.Pages.LandingPages.NewAbTestVariation(context.TODO(), cms.PageLandingPageNewAbTestVariationParams{
		AbTestCreateRequestVNext: shared.AbTestCreateRequestVNextParam{
			ContentID:     "contentId",
			VariationName: "variationName",
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

func TestPageLandingPageNewBatch(t *testing.T) {
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
	_, err := client.Cms.Pages.LandingPages.NewBatch(context.TODO(), cms.PageLandingPageNewBatchParams{
		BatchInputPage: cms.BatchInputPageParam{
			Inputs: []cms.PageParam{{
				ID:                  "id",
				AbStatus:            cms.PageAbStatusAutomatedLoserVariant,
				AbTestID:            "abTestId",
				ArchivedAt:          time.Now(),
				ArchivedInDashboard: true,
				AttachedStylesheets: []map[string]any{{
					"foo": map[string]any{},
				}},
				AuthorName:                "authorName",
				Campaign:                  "campaign",
				CategoryID:                0,
				ContentGroupID:            "contentGroupId",
				ContentTypeCategory:       cms.PageContentTypeCategory0,
				Created:                   time.Now(),
				CreatedByID:               "createdById",
				CurrentlyPublished:        true,
				CurrentState:              cms.PageCurrentStateAutomated,
				Domain:                    "domain",
				DynamicPageDataSourceID:   "dynamicPageDataSourceId",
				DynamicPageDataSourceType: 0,
				DynamicPageHubDBTableID:   "dynamicPageHubDbTableId",
				EnableDomainStylesheets:   true,
				EnableLayoutStylesheets:   true,
				FeaturedImage:             "featuredImage",
				FeaturedImageAltText:      "featuredImageAltText",
				FolderID:                  "folderId",
				FooterHTML:                "footerHtml",
				HeadHTML:                  "headHtml",
				HTMLTitle:                 "htmlTitle",
				IncludeDefaultCustomCss:   true,
				Language:                  cms.PageLanguageAf,
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
				PageRedirected:           true,
				Password:                 "password",
				PublicAccessRules:        []cms.PublicAccessRule{map[string]any{}},
				PublicAccessRulesEnabled: true,
				PublishDate:              time.Now(),
				PublishImmediately:       true,
				Slug:                     "slug",
				State:                    "state",
				Subcategory:              "subcategory",
				TemplatePath:             "templatePath",
				ThemeSettingsValues: map[string]any{
					"foo": map[string]any{},
				},
				TranslatedFromID: "translatedFromId",
				Translations: map[string]cms.PagesContentLanguageVariationParam{
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

func TestPageLandingPageNewFolder(t *testing.T) {
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
	_, err := client.Cms.Pages.LandingPages.NewFolder(context.TODO(), cms.PageLandingPageNewFolderParams{
		ContentFolder: cms.ContentFolderParam{
			ID:             "id",
			Category:       0,
			Created:        time.Now(),
			DeletedAt:      time.Now(),
			Name:           "name",
			ParentFolderID: 0,
			Updated:        time.Now(),
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

func TestPageLandingPageNewFoldersBatch(t *testing.T) {
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
	_, err := client.Cms.Pages.LandingPages.NewFoldersBatch(context.TODO(), cms.PageLandingPageNewFoldersBatchParams{
		BatchInputContentFolder: cms.BatchInputContentFolderParam{
			Inputs: []cms.ContentFolderParam{{
				ID:             "id",
				Category:       0,
				Created:        time.Now(),
				DeletedAt:      time.Now(),
				Name:           "name",
				ParentFolderID: 0,
				Updated:        time.Now(),
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

func TestPageLandingPageNewLanguageVariationWithOptionalParams(t *testing.T) {
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
	_, err := client.Cms.Pages.LandingPages.NewLanguageVariation(context.TODO(), cms.PageLandingPageNewLanguageVariationParams{
		ContentLanguageCloneRequestVNext: cms.ContentLanguageCloneRequestVNextParam{
			ID:              "id",
			Language:        hubspotsdk.String("language"),
			PrimaryLanguage: hubspotsdk.String("primaryLanguage"),
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

func TestPageLandingPageDeleteBatch(t *testing.T) {
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
	err := client.Cms.Pages.LandingPages.DeleteBatch(context.TODO(), cms.PageLandingPageDeleteBatchParams{
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

func TestPageLandingPageDeleteFolderWithOptionalParams(t *testing.T) {
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
	err := client.Cms.Pages.LandingPages.DeleteFolder(
		context.TODO(),
		"objectId",
		cms.PageLandingPageDeleteFolderParams{
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

func TestPageLandingPageDeleteFoldersBatch(t *testing.T) {
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
	err := client.Cms.Pages.LandingPages.DeleteFoldersBatch(context.TODO(), cms.PageLandingPageDeleteFoldersBatchParams{
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

func TestPageLandingPageDetachFromLangGroup(t *testing.T) {
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
	err := client.Cms.Pages.LandingPages.DetachFromLangGroup(context.TODO(), cms.PageLandingPageDetachFromLangGroupParams{
		DetachFromLangGroupRequestVNext: cms.DetachFromLangGroupRequestVNextParam{
			ID: "id",
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

func TestPageLandingPageEndAbTest(t *testing.T) {
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
	err := client.Cms.Pages.LandingPages.EndAbTest(context.TODO(), cms.PageLandingPageEndAbTestParams{
		AbTestEndRequestVNext: cms.AbTestEndRequestVNextParam{
			AbTestID: "abTestId",
			WinnerID: "winnerId",
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

func TestPageLandingPageGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Cms.Pages.LandingPages.Get(
		context.TODO(),
		"objectId",
		cms.PageLandingPageGetParams{
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
}

func TestPageLandingPageGetBatchWithOptionalParams(t *testing.T) {
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
	_, err := client.Cms.Pages.LandingPages.GetBatch(context.TODO(), cms.PageLandingPageGetBatchParams{
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

func TestPageLandingPageGetDraft(t *testing.T) {
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
	_, err := client.Cms.Pages.LandingPages.GetDraft(context.TODO(), "objectId")
	if err != nil {
		var apierr *hubspotsdk.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPageLandingPageGetFolderWithOptionalParams(t *testing.T) {
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
	_, err := client.Cms.Pages.LandingPages.GetFolder(
		context.TODO(),
		"objectId",
		cms.PageLandingPageGetFolderParams{
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
}

func TestPageLandingPageGetFolderRevision(t *testing.T) {
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
	_, err := client.Cms.Pages.LandingPages.GetFolderRevision(
		context.TODO(),
		"revisionId",
		cms.PageLandingPageGetFolderRevisionParams{
			ObjectID: "objectId",
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

func TestPageLandingPageGetFoldersBatchWithOptionalParams(t *testing.T) {
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
	_, err := client.Cms.Pages.LandingPages.GetFoldersBatch(context.TODO(), cms.PageLandingPageGetFoldersBatchParams{
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

func TestPageLandingPageGetRevision(t *testing.T) {
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
	_, err := client.Cms.Pages.LandingPages.GetRevision(
		context.TODO(),
		"revisionId",
		cms.PageLandingPageGetRevisionParams{
			ObjectID: "objectId",
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

func TestPageLandingPageListFolderRevisionsWithOptionalParams(t *testing.T) {
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
	_, err := client.Cms.Pages.LandingPages.ListFolderRevisions(
		context.TODO(),
		"objectId",
		cms.PageLandingPageListFolderRevisionsParams{
			After:  hubspotsdk.String("after"),
			Before: hubspotsdk.String("before"),
			Limit:  hubspotsdk.Int(0),
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

func TestPageLandingPageListFoldersWithOptionalParams(t *testing.T) {
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
	_, err := client.Cms.Pages.LandingPages.ListFolders(context.TODO(), cms.PageLandingPageListFoldersParams{
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
}

func TestPageLandingPageListRevisionsWithOptionalParams(t *testing.T) {
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
	_, err := client.Cms.Pages.LandingPages.ListRevisions(
		context.TODO(),
		"objectId",
		cms.PageLandingPageListRevisionsParams{
			After:  hubspotsdk.String("after"),
			Before: hubspotsdk.String("before"),
			Limit:  hubspotsdk.Int(0),
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

func TestPageLandingPagePublishDraft(t *testing.T) {
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
	err := client.Cms.Pages.LandingPages.PublishDraft(context.TODO(), "objectId")
	if err != nil {
		var apierr *hubspotsdk.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPageLandingPageRerunAbTest(t *testing.T) {
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
	err := client.Cms.Pages.LandingPages.RerunAbTest(context.TODO(), cms.PageLandingPageRerunAbTestParams{
		AbTestRerunRequestVNext: cms.AbTestRerunRequestVNextParam{
			AbTestID:    "abTestId",
			VariationID: "variationId",
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

func TestPageLandingPageResetDraft(t *testing.T) {
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
	err := client.Cms.Pages.LandingPages.ResetDraft(context.TODO(), "objectId")
	if err != nil {
		var apierr *hubspotsdk.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPageLandingPageRestoreFolderRevision(t *testing.T) {
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
	_, err := client.Cms.Pages.LandingPages.RestoreFolderRevision(
		context.TODO(),
		"revisionId",
		cms.PageLandingPageRestoreFolderRevisionParams{
			ObjectID: "objectId",
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

func TestPageLandingPageRestoreRevision(t *testing.T) {
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
	_, err := client.Cms.Pages.LandingPages.RestoreRevision(
		context.TODO(),
		"revisionId",
		cms.PageLandingPageRestoreRevisionParams{
			ObjectID: "objectId",
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

func TestPageLandingPageRestoreRevisionToDraft(t *testing.T) {
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
	_, err := client.Cms.Pages.LandingPages.RestoreRevisionToDraft(
		context.TODO(),
		0,
		cms.PageLandingPageRestoreRevisionToDraftParams{
			ObjectID: "objectId",
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

func TestPageLandingPageSchedule(t *testing.T) {
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
	err := client.Cms.Pages.LandingPages.Schedule(context.TODO(), cms.PageLandingPageScheduleParams{
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

func TestPageLandingPageSetNewLangPrimary(t *testing.T) {
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
	err := client.Cms.Pages.LandingPages.SetNewLangPrimary(context.TODO(), cms.PageLandingPageSetNewLangPrimaryParams{
		SetNewLanguagePrimaryRequestVNext: cms.SetNewLanguagePrimaryRequestVNextParam{
			ID: "id",
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

func TestPageLandingPageUpdateBatchWithOptionalParams(t *testing.T) {
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
	_, err := client.Cms.Pages.LandingPages.UpdateBatch(context.TODO(), cms.PageLandingPageUpdateBatchParams{
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

func TestPageLandingPageUpdateDraft(t *testing.T) {
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
	_, err := client.Cms.Pages.LandingPages.UpdateDraft(
		context.TODO(),
		"objectId",
		cms.PageLandingPageUpdateDraftParams{
			Page: cms.PageParam{
				ID:                  "id",
				AbStatus:            cms.PageAbStatusAutomatedLoserVariant,
				AbTestID:            "abTestId",
				ArchivedAt:          time.Now(),
				ArchivedInDashboard: true,
				AttachedStylesheets: []map[string]any{{
					"foo": map[string]any{},
				}},
				AuthorName:                "authorName",
				Campaign:                  "campaign",
				CategoryID:                0,
				ContentGroupID:            "contentGroupId",
				ContentTypeCategory:       cms.PageContentTypeCategory0,
				Created:                   time.Now(),
				CreatedByID:               "createdById",
				CurrentlyPublished:        true,
				CurrentState:              cms.PageCurrentStateAutomated,
				Domain:                    "domain",
				DynamicPageDataSourceID:   "dynamicPageDataSourceId",
				DynamicPageDataSourceType: 0,
				DynamicPageHubDBTableID:   "dynamicPageHubDbTableId",
				EnableDomainStylesheets:   true,
				EnableLayoutStylesheets:   true,
				FeaturedImage:             "featuredImage",
				FeaturedImageAltText:      "featuredImageAltText",
				FolderID:                  "folderId",
				FooterHTML:                "footerHtml",
				HeadHTML:                  "headHtml",
				HTMLTitle:                 "htmlTitle",
				IncludeDefaultCustomCss:   true,
				Language:                  cms.PageLanguageAf,
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
				PageRedirected:           true,
				Password:                 "password",
				PublicAccessRules:        []cms.PublicAccessRule{map[string]any{}},
				PublicAccessRulesEnabled: true,
				PublishDate:              time.Now(),
				PublishImmediately:       true,
				Slug:                     "slug",
				State:                    "state",
				Subcategory:              "subcategory",
				TemplatePath:             "templatePath",
				ThemeSettingsValues: map[string]any{
					"foo": map[string]any{},
				},
				TranslatedFromID: "translatedFromId",
				Translations: map[string]cms.PagesContentLanguageVariationParam{
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
}

func TestPageLandingPageUpdateFolderWithOptionalParams(t *testing.T) {
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
	_, err := client.Cms.Pages.LandingPages.UpdateFolder(
		context.TODO(),
		"objectId",
		cms.PageLandingPageUpdateFolderParams{
			ContentFolder: cms.ContentFolderParam{
				ID:             "id",
				Category:       0,
				Created:        time.Now(),
				DeletedAt:      time.Now(),
				Name:           "name",
				ParentFolderID: 0,
				Updated:        time.Now(),
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
}

func TestPageLandingPageUpdateFoldersBatchWithOptionalParams(t *testing.T) {
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
	_, err := client.Cms.Pages.LandingPages.UpdateFoldersBatch(context.TODO(), cms.PageLandingPageUpdateFoldersBatchParams{
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

func TestPageLandingPageUpdateLanguages(t *testing.T) {
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
	err := client.Cms.Pages.LandingPages.UpdateLanguages(context.TODO(), cms.PageLandingPageUpdateLanguagesParams{
		UpdateLanguagesRequestVNext: cms.UpdateLanguagesRequestVNextParam{
			Languages: map[string]string{
				"foo": "string",
			},
			PrimaryID: "primaryId",
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
