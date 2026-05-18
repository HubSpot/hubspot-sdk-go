// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cms_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/HubSpot/hubspot-sdk-go"
	"github.com/HubSpot/hubspot-sdk-go/cms"
	"github.com/HubSpot/hubspot-sdk-go/internal/testutil"
	"github.com/HubSpot/hubspot-sdk-go/option"
	"github.com/HubSpot/hubspot-sdk-go/shared"
)

func TestPageSitePageBatchNewSitePages(t *testing.T) {
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
	_, err := client.Cms.Pages.SitePages.Batch.NewSitePages(context.TODO(), cms.PageSitePageBatchNewSitePagesParams{
		BatchInputPage: cms.BatchInputPageParam{
			Inputs: []cms.PagesPageParam{{
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
										Units: cms.AngleUnitsDegrees,
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
												Units: cms.SizeUnitsCh,
												Value: 0,
											},
											Top: cms.SizeParam{
												Units: cms.SizeUnitsCh,
												Value: 0,
											},
										},
										Padding: cms.PaddingParam{
											Bottom: cms.SizeParam{
												Units: cms.SizeUnitsCh,
												Value: 0,
											},
											Left: cms.SizeParam{
												Units: cms.SizeUnitsCh,
												Value: 0,
											},
											Right: cms.SizeParam{
												Units: cms.SizeUnitsCh,
												Value: 0,
											},
											Top: cms.SizeParam{
												Units: cms.SizeUnitsCh,
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
									Units: cms.AngleUnitsDegrees,
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
											Units: cms.SizeUnitsCh,
											Value: 0,
										},
										Top: cms.SizeParam{
											Units: cms.SizeUnitsCh,
											Value: 0,
										},
									},
									Padding: cms.PaddingParam{
										Bottom: cms.SizeParam{
											Units: cms.SizeUnitsCh,
											Value: 0,
										},
										Left: cms.SizeParam{
											Units: cms.SizeUnitsCh,
											Value: 0,
										},
										Right: cms.SizeParam{
											Units: cms.SizeUnitsCh,
											Value: 0,
										},
										Top: cms.SizeParam{
											Units: cms.SizeUnitsCh,
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

func TestPageSitePageBatchDeleteSitePages(t *testing.T) {
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
	err := client.Cms.Pages.SitePages.Batch.DeleteSitePages(context.TODO(), cms.PageSitePageBatchDeleteSitePagesParams{
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

func TestPageSitePageBatchGetSitePagesWithOptionalParams(t *testing.T) {
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
	_, err := client.Cms.Pages.SitePages.Batch.GetSitePages(context.TODO(), cms.PageSitePageBatchGetSitePagesParams{
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

func TestPageSitePageBatchUpdateSitePagesWithOptionalParams(t *testing.T) {
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
	_, err := client.Cms.Pages.SitePages.Batch.UpdateSitePages(context.TODO(), cms.PageSitePageBatchUpdateSitePagesParams{
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
