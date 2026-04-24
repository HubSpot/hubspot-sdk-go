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
)

func TestPageWebsitePageNew(t *testing.T) {
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
	_, err := client.Cms.Pages.WebsitePages.New(context.TODO(), cms.PageWebsitePageNewParams{
		PagesPage: cms.PagesPageParam{
			ID:                  "id",
			AbStatus:            cms.PagesPageAbStatusAutomatedLoserVariant,
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
			ContentTypeCategory:       cms.PagesPageContentTypeCategory0,
			Created:                   time.Now(),
			CreatedByID:               "createdById",
			CurrentlyPublished:        true,
			CurrentState:              cms.PagesPageCurrentStateAgentGenerated,
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
			Language:                  cms.PagesPageLanguageAa,
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
}

func TestPageWebsitePageUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Cms.Pages.WebsitePages.Update(
		context.TODO(),
		"objectId",
		cms.PageWebsitePageUpdateParams{
			PagesPage: cms.PagesPageParam{
				ID:                  "id",
				AbStatus:            cms.PagesPageAbStatusAutomatedLoserVariant,
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
				ContentTypeCategory:       cms.PagesPageContentTypeCategory0,
				Created:                   time.Now(),
				CreatedByID:               "createdById",
				CurrentlyPublished:        true,
				CurrentState:              cms.PagesPageCurrentStateAgentGenerated,
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
				Language:                  cms.PagesPageLanguageAa,
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
}

func TestPageWebsitePageListWithOptionalParams(t *testing.T) {
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
	_, err := client.Cms.Pages.WebsitePages.List(context.TODO(), cms.PageWebsitePageListParams{
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

func TestPageWebsitePageDeleteWithOptionalParams(t *testing.T) {
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
	err := client.Cms.Pages.WebsitePages.Delete(
		context.TODO(),
		"objectId",
		cms.PageWebsitePageDeleteParams{
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

func TestPageWebsitePageCloneWithOptionalParams(t *testing.T) {
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
	_, err := client.Cms.Pages.WebsitePages.Clone(context.TODO(), cms.PageWebsitePageCloneParams{
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

func TestPageWebsitePageGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Cms.Pages.WebsitePages.Get(
		context.TODO(),
		"objectId",
		cms.PageWebsitePageGetParams{
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

func TestPageWebsitePageGetDraft(t *testing.T) {
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
	_, err := client.Cms.Pages.WebsitePages.GetDraft(context.TODO(), "objectId")
	if err != nil {
		var apierr *hubspotsdk.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPageWebsitePagePublishDraft(t *testing.T) {
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
	err := client.Cms.Pages.WebsitePages.PublishDraft(context.TODO(), "objectId")
	if err != nil {
		var apierr *hubspotsdk.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPageWebsitePageSchedule(t *testing.T) {
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
	err := client.Cms.Pages.WebsitePages.Schedule(context.TODO(), cms.PageWebsitePageScheduleParams{
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

func TestPageWebsitePageSetNewLangPrimary(t *testing.T) {
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
	err := client.Cms.Pages.WebsitePages.SetNewLangPrimary(context.TODO(), cms.PageWebsitePageSetNewLangPrimaryParams{
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

func TestPageWebsitePageUpdateDraft(t *testing.T) {
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
	_, err := client.Cms.Pages.WebsitePages.UpdateDraft(
		context.TODO(),
		"objectId",
		cms.PageWebsitePageUpdateDraftParams{
			PagesPage: cms.PagesPageParam{
				ID:                  "id",
				AbStatus:            cms.PagesPageAbStatusAutomatedLoserVariant,
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
				ContentTypeCategory:       cms.PagesPageContentTypeCategory0,
				Created:                   time.Now(),
				CreatedByID:               "createdById",
				CurrentlyPublished:        true,
				CurrentState:              cms.PagesPageCurrentStateAgentGenerated,
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
				Language:                  cms.PagesPageLanguageAa,
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
}
