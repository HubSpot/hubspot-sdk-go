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

func TestPageSitePageNew(t *testing.T) {
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
	err := client.Cms.Pages.SitePages.New(context.TODO(), cms.PageSitePageNewParams{
		Page: cms.PageParam{
			ID:                  "id",
			AbStatus:            cms.PageAbStatusMaster,
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

func TestPageSitePageUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Cms.Pages.SitePages.Update(
		context.TODO(),
		"objectId",
		cms.PageSitePageUpdateParams{
			Page: cms.PageParam{
				ID:                  "id",
				AbStatus:            cms.PageAbStatusMaster,
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

func TestPageSitePageListWithOptionalParams(t *testing.T) {
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
	_, err := client.Cms.Pages.SitePages.List(context.TODO(), cms.PageSitePageListParams{
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

func TestPageSitePageDeleteWithOptionalParams(t *testing.T) {
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
	err := client.Cms.Pages.SitePages.Delete(
		context.TODO(),
		"objectId",
		cms.PageSitePageDeleteParams{
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

func TestPageSitePageAttachToLangGroupWithOptionalParams(t *testing.T) {
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
	err := client.Cms.Pages.SitePages.AttachToLangGroup(context.TODO(), cms.PageSitePageAttachToLangGroupParams{
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

func TestPageSitePageCloneWithOptionalParams(t *testing.T) {
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
	_, err := client.Cms.Pages.SitePages.Clone(context.TODO(), cms.PageSitePageCloneParams{
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

func TestPageSitePageNewAbTestVariation(t *testing.T) {
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
	_, err := client.Cms.Pages.SitePages.NewAbTestVariation(context.TODO(), cms.PageSitePageNewAbTestVariationParams{
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

func TestPageSitePageNewBatch(t *testing.T) {
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
	_, err := client.Cms.Pages.SitePages.NewBatch(context.TODO(), cms.PageSitePageNewBatchParams{
		BatchInputPage: cms.BatchInputPageParam{
			Inputs: []cms.PageParam{{
				ID:                  "id",
				AbStatus:            cms.PageAbStatusMaster,
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

func TestPageSitePageNewLanguageVariationWithOptionalParams(t *testing.T) {
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
	_, err := client.Cms.Pages.SitePages.NewLanguageVariation(context.TODO(), cms.PageSitePageNewLanguageVariationParams{
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

func TestPageSitePageDeleteBatch(t *testing.T) {
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
	err := client.Cms.Pages.SitePages.DeleteBatch(context.TODO(), cms.PageSitePageDeleteBatchParams{
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

func TestPageSitePageDetachFromLangGroup(t *testing.T) {
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
	err := client.Cms.Pages.SitePages.DetachFromLangGroup(context.TODO(), cms.PageSitePageDetachFromLangGroupParams{
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

func TestPageSitePageEndAbTest(t *testing.T) {
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
	err := client.Cms.Pages.SitePages.EndAbTest(context.TODO(), cms.PageSitePageEndAbTestParams{
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

func TestPageSitePageGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Cms.Pages.SitePages.Get(
		context.TODO(),
		"objectId",
		cms.PageSitePageGetParams{
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

func TestPageSitePageGetBatchWithOptionalParams(t *testing.T) {
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
	_, err := client.Cms.Pages.SitePages.GetBatch(context.TODO(), cms.PageSitePageGetBatchParams{
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

func TestPageSitePageGetDraft(t *testing.T) {
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
	_, err := client.Cms.Pages.SitePages.GetDraft(context.TODO(), "objectId")
	if err != nil {
		var apierr *hubspotsdk.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPageSitePageGetRevision(t *testing.T) {
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
	_, err := client.Cms.Pages.SitePages.GetRevision(
		context.TODO(),
		"revisionId",
		cms.PageSitePageGetRevisionParams{
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

func TestPageSitePageListRevisionsWithOptionalParams(t *testing.T) {
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
	_, err := client.Cms.Pages.SitePages.ListRevisions(
		context.TODO(),
		"objectId",
		cms.PageSitePageListRevisionsParams{
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

func TestPageSitePagePublishDraft(t *testing.T) {
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
	err := client.Cms.Pages.SitePages.PublishDraft(context.TODO(), "objectId")
	if err != nil {
		var apierr *hubspotsdk.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPageSitePageRerunAbTest(t *testing.T) {
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
	err := client.Cms.Pages.SitePages.RerunAbTest(context.TODO(), cms.PageSitePageRerunAbTestParams{
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

func TestPageSitePageResetDraft(t *testing.T) {
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
	err := client.Cms.Pages.SitePages.ResetDraft(context.TODO(), "objectId")
	if err != nil {
		var apierr *hubspotsdk.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPageSitePageRestoreRevision(t *testing.T) {
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
	_, err := client.Cms.Pages.SitePages.RestoreRevision(
		context.TODO(),
		"revisionId",
		cms.PageSitePageRestoreRevisionParams{
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

func TestPageSitePageRestoreRevisionToDraft(t *testing.T) {
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
	_, err := client.Cms.Pages.SitePages.RestoreRevisionToDraft(
		context.TODO(),
		0,
		cms.PageSitePageRestoreRevisionToDraftParams{
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

func TestPageSitePageSchedule(t *testing.T) {
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
	err := client.Cms.Pages.SitePages.Schedule(context.TODO(), cms.PageSitePageScheduleParams{
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

func TestPageSitePageSetNewLangPrimary(t *testing.T) {
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
	err := client.Cms.Pages.SitePages.SetNewLangPrimary(context.TODO(), cms.PageSitePageSetNewLangPrimaryParams{
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

func TestPageSitePageUpdateBatchWithOptionalParams(t *testing.T) {
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
	_, err := client.Cms.Pages.SitePages.UpdateBatch(context.TODO(), cms.PageSitePageUpdateBatchParams{
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

func TestPageSitePageUpdateDraft(t *testing.T) {
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
	_, err := client.Cms.Pages.SitePages.UpdateDraft(
		context.TODO(),
		"objectId",
		cms.PageSitePageUpdateDraftParams{
			Page: cms.PageParam{
				ID:                  "id",
				AbStatus:            cms.PageAbStatusMaster,
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

func TestPageSitePageUpdateLanguages(t *testing.T) {
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
	err := client.Cms.Pages.SitePages.UpdateLanguages(context.TODO(), cms.PageSitePageUpdateLanguagesParams{
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
