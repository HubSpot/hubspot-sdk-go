// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package marketing_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/HubSpot/hubspot-sdk-go"
	"github.com/HubSpot/hubspot-sdk-go/internal/testutil"
	"github.com/HubSpot/hubspot-sdk-go/marketing"
	"github.com/HubSpot/hubspot-sdk-go/option"
	"github.com/HubSpot/hubspot-sdk-go/shared"
)

func TestEmailNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Marketing.Emails.New(context.TODO(), marketing.EmailNewParams{
		EmailCreateRequest: marketing.EmailCreateRequestParam{
			ActiveDomain:   hubspotsdk.String("activeDomain"),
			Archived:       hubspotsdk.Bool(true),
			BusinessUnitID: hubspotsdk.Int(0),
			Campaign:       hubspotsdk.String("campaign"),
			Content: marketing.PublicEmailContentParam{
				FlexAreas: map[string]any{
					"foo": map[string]any{},
				},
				PlainTextVersion: hubspotsdk.String("plainTextVersion"),
				SmartFields: map[string]marketing.SmartEmailField{
					"foo": map[string]any{},
				},
				StyleSettings: marketing.PublicEmailStyleSettingsParam{
					BackgroundColor:       hubspotsdk.String("backgroundColor"),
					BackgroundImage:       hubspotsdk.String("backgroundImage"),
					BackgroundImageType:   marketing.PublicEmailStyleSettingsBackgroundImageTypeRepeat,
					BodyBorderColor:       hubspotsdk.String("bodyBorderColor"),
					BodyBorderColorChoice: hubspotsdk.String("bodyBorderColorChoice"),
					BodyBorderWidth:       hubspotsdk.Float(0),
					BodyColor:             hubspotsdk.String("bodyColor"),
					ButtonStyleSettings: marketing.PublicButtonStyleSettingsParam{
						BackgroundColor: map[string]any{},
						CornerRadius:    hubspotsdk.Int(0),
						FontStyle: marketing.PublicFontStyleParam{
							Bold:      hubspotsdk.Bool(true),
							Color:     hubspotsdk.String("color"),
							Font:      hubspotsdk.String("font"),
							Italic:    hubspotsdk.Bool(true),
							Size:      hubspotsdk.Int(0),
							Underline: hubspotsdk.Bool(true),
						},
					},
					ColorPickerFavorite1: hubspotsdk.String("colorPickerFavorite1"),
					ColorPickerFavorite2: hubspotsdk.String("colorPickerFavorite2"),
					ColorPickerFavorite3: hubspotsdk.String("colorPickerFavorite3"),
					ColorPickerFavorite4: hubspotsdk.String("colorPickerFavorite4"),
					ColorPickerFavorite5: hubspotsdk.String("colorPickerFavorite5"),
					ColorPickerFavorite6: hubspotsdk.String("colorPickerFavorite6"),
					DividerStyleSettings: marketing.PublicDividerStyleSettingsParam{
						Color:    map[string]any{},
						Height:   hubspotsdk.Int(0),
						LineType: hubspotsdk.String("lineType"),
					},
					EmailBodyPadding: hubspotsdk.String("emailBodyPadding"),
					EmailBodyWidth:   hubspotsdk.String("emailBodyWidth"),
					HeadingOneFont: marketing.PublicFontStyleParam{
						Bold:      hubspotsdk.Bool(true),
						Color:     hubspotsdk.String("color"),
						Font:      hubspotsdk.String("font"),
						Italic:    hubspotsdk.Bool(true),
						Size:      hubspotsdk.Int(0),
						Underline: hubspotsdk.Bool(true),
					},
					HeadingTwoFont: marketing.PublicFontStyleParam{
						Bold:      hubspotsdk.Bool(true),
						Color:     hubspotsdk.String("color"),
						Font:      hubspotsdk.String("font"),
						Italic:    hubspotsdk.Bool(true),
						Size:      hubspotsdk.Int(0),
						Underline: hubspotsdk.Bool(true),
					},
					LinksFont: marketing.PublicFontStyleParam{
						Bold:      hubspotsdk.Bool(true),
						Color:     hubspotsdk.String("color"),
						Font:      hubspotsdk.String("font"),
						Italic:    hubspotsdk.Bool(true),
						Size:      hubspotsdk.Int(0),
						Underline: hubspotsdk.Bool(true),
					},
					PrimaryAccentColor:      hubspotsdk.String("primaryAccentColor"),
					PrimaryFont:             hubspotsdk.String("primaryFont"),
					PrimaryFontColor:        hubspotsdk.String("primaryFontColor"),
					PrimaryFontLineHeight:   hubspotsdk.String("primaryFontLineHeight"),
					PrimaryFontSize:         hubspotsdk.Float(0),
					SecondaryAccentColor:    hubspotsdk.String("secondaryAccentColor"),
					SecondaryFont:           hubspotsdk.String("secondaryFont"),
					SecondaryFontColor:      hubspotsdk.String("secondaryFontColor"),
					SecondaryFontLineHeight: hubspotsdk.String("secondaryFontLineHeight"),
					SecondaryFontSize:       hubspotsdk.Float(0),
				},
				TemplatePath: hubspotsdk.String("templatePath"),
				ThemeSettingsValues: map[string]any{
					"foo": map[string]any{},
				},
				WidgetContainers: map[string]any{
					"foo": map[string]any{},
				},
				Widgets: map[string]any{
					"foo": map[string]any{},
				},
			},
			FeedbackSurveyID: hubspotsdk.String("feedbackSurveyId"),
			FolderIDV2:       hubspotsdk.Int(0),
			From: marketing.PublicEmailFromDetailsParam{
				CustomReplyTo: hubspotsdk.String("customReplyTo"),
				FromName:      hubspotsdk.String("fromName"),
				ReplyTo:       hubspotsdk.String("replyTo"),
			},
			JitterSendTime: hubspotsdk.Bool(true),
			Language:       marketing.EmailCreateRequestLanguageAa,
			Name:           hubspotsdk.String("name"),
			PublishDate:    hubspotsdk.Time(time.Now()),
			RssData: marketing.PublicRssEmailDetailsParam{
				BlogEmailType:     hubspotsdk.String("blogEmailType"),
				BlogImageMaxWidth: hubspotsdk.Int(0),
				BlogLayout:        marketing.PublicRssEmailDetailsBlogLayoutFullPost,
				HubSpotBlogID:     hubspotsdk.String("hubspotBlogId"),
				MaxEntries:        hubspotsdk.Int(0),
				RssEntryTemplate:  hubspotsdk.String("rssEntryTemplate"),
				Timing: map[string]any{
					"foo": map[string]any{},
				},
				URL:                  hubspotsdk.String("url"),
				UseHeadlineAsSubject: hubspotsdk.Bool(true),
			},
			SendOnPublish: hubspotsdk.Bool(true),
			State:         marketing.EmailCreateRequestStateAgentGenerated,
			Subcategory:   marketing.EmailCreateRequestSubcategoryAbLoserVariant,
			Subject:       hubspotsdk.String("subject"),
			SubscriptionDetails: marketing.PublicEmailSubscriptionDetailsParam{
				OfficeLocationID:   hubspotsdk.String("officeLocationId"),
				PreferencesGroupID: hubspotsdk.String("preferencesGroupId"),
				SubscriptionID:     hubspotsdk.String("subscriptionId"),
				SubscriptionName:   hubspotsdk.String("subscriptionName"),
			},
			Testing: marketing.PublicEmailTestingDetailsParam{
				IsAbVariation:       true,
				AbSampleSizeDefault: marketing.PublicEmailTestingDetailsAbSampleSizeDefaultAutomatedLoserVariant,
				AbSamplingDefault:   marketing.PublicEmailTestingDetailsAbSamplingDefaultAutomatedLoserVariant,
				AbStatus:            marketing.PublicEmailTestingDetailsAbStatusAutomatedLoserVariant,
				AbSuccessMetric:     marketing.PublicEmailTestingDetailsAbSuccessMetricClicksByDelivered,
				AbTestPercentage:    hubspotsdk.Int(0),
				HoursToWait:         hubspotsdk.Int(0),
				TestID:              hubspotsdk.String("testId"),
			},
			To: marketing.PublicEmailToDetailsParam{
				ContactIDs: marketing.PublicEmailRecipientsParam{
					Exclude: []string{"string"},
					Include: []string{"string"},
				},
				ContactIlsLists: marketing.PublicEmailRecipientsParam{
					Exclude: []string{"string"},
					Include: []string{"string"},
				},
				ContactLists: marketing.PublicEmailRecipientsParam{
					Exclude: []string{"string"},
					Include: []string{"string"},
				},
				LimitSendFrequency: hubspotsdk.Bool(true),
				SuppressGraymail:   hubspotsdk.Bool(true),
			},
			Webversion: marketing.PublicWebversionDetailsParam{
				Domain:            hubspotsdk.String("domain"),
				Enabled:           hubspotsdk.Bool(true),
				ExpiresAt:         hubspotsdk.Time(time.Now()),
				IsPageRedirected:  hubspotsdk.Bool(true),
				MetaDescription:   hubspotsdk.String("metaDescription"),
				PageExpiryEnabled: hubspotsdk.Bool(true),
				RedirectToPageID:  hubspotsdk.String("redirectToPageId"),
				RedirectToURL:     hubspotsdk.String("redirectToUrl"),
				Slug:              hubspotsdk.String("slug"),
				Title:             hubspotsdk.String("title"),
				URL:               hubspotsdk.String("url"),
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

func TestEmailUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Marketing.Emails.Update(
		context.TODO(),
		"emailId",
		marketing.EmailUpdateParams{
			EmailUpdateRequest: marketing.EmailUpdateRequestParam{
				ActiveDomain:   hubspotsdk.String("activeDomain"),
				Archived:       hubspotsdk.Bool(true),
				BusinessUnitID: hubspotsdk.Int(0),
				Campaign:       hubspotsdk.String("campaign"),
				Content: marketing.PublicEmailContentParam{
					FlexAreas: map[string]any{
						"foo": map[string]any{},
					},
					PlainTextVersion: hubspotsdk.String("plainTextVersion"),
					SmartFields: map[string]marketing.SmartEmailField{
						"foo": map[string]any{},
					},
					StyleSettings: marketing.PublicEmailStyleSettingsParam{
						BackgroundColor:       hubspotsdk.String("backgroundColor"),
						BackgroundImage:       hubspotsdk.String("backgroundImage"),
						BackgroundImageType:   marketing.PublicEmailStyleSettingsBackgroundImageTypeRepeat,
						BodyBorderColor:       hubspotsdk.String("bodyBorderColor"),
						BodyBorderColorChoice: hubspotsdk.String("bodyBorderColorChoice"),
						BodyBorderWidth:       hubspotsdk.Float(0),
						BodyColor:             hubspotsdk.String("bodyColor"),
						ButtonStyleSettings: marketing.PublicButtonStyleSettingsParam{
							BackgroundColor: map[string]any{},
							CornerRadius:    hubspotsdk.Int(0),
							FontStyle: marketing.PublicFontStyleParam{
								Bold:      hubspotsdk.Bool(true),
								Color:     hubspotsdk.String("color"),
								Font:      hubspotsdk.String("font"),
								Italic:    hubspotsdk.Bool(true),
								Size:      hubspotsdk.Int(0),
								Underline: hubspotsdk.Bool(true),
							},
						},
						ColorPickerFavorite1: hubspotsdk.String("colorPickerFavorite1"),
						ColorPickerFavorite2: hubspotsdk.String("colorPickerFavorite2"),
						ColorPickerFavorite3: hubspotsdk.String("colorPickerFavorite3"),
						ColorPickerFavorite4: hubspotsdk.String("colorPickerFavorite4"),
						ColorPickerFavorite5: hubspotsdk.String("colorPickerFavorite5"),
						ColorPickerFavorite6: hubspotsdk.String("colorPickerFavorite6"),
						DividerStyleSettings: marketing.PublicDividerStyleSettingsParam{
							Color:    map[string]any{},
							Height:   hubspotsdk.Int(0),
							LineType: hubspotsdk.String("lineType"),
						},
						EmailBodyPadding: hubspotsdk.String("emailBodyPadding"),
						EmailBodyWidth:   hubspotsdk.String("emailBodyWidth"),
						HeadingOneFont: marketing.PublicFontStyleParam{
							Bold:      hubspotsdk.Bool(true),
							Color:     hubspotsdk.String("color"),
							Font:      hubspotsdk.String("font"),
							Italic:    hubspotsdk.Bool(true),
							Size:      hubspotsdk.Int(0),
							Underline: hubspotsdk.Bool(true),
						},
						HeadingTwoFont: marketing.PublicFontStyleParam{
							Bold:      hubspotsdk.Bool(true),
							Color:     hubspotsdk.String("color"),
							Font:      hubspotsdk.String("font"),
							Italic:    hubspotsdk.Bool(true),
							Size:      hubspotsdk.Int(0),
							Underline: hubspotsdk.Bool(true),
						},
						LinksFont: marketing.PublicFontStyleParam{
							Bold:      hubspotsdk.Bool(true),
							Color:     hubspotsdk.String("color"),
							Font:      hubspotsdk.String("font"),
							Italic:    hubspotsdk.Bool(true),
							Size:      hubspotsdk.Int(0),
							Underline: hubspotsdk.Bool(true),
						},
						PrimaryAccentColor:      hubspotsdk.String("primaryAccentColor"),
						PrimaryFont:             hubspotsdk.String("primaryFont"),
						PrimaryFontColor:        hubspotsdk.String("primaryFontColor"),
						PrimaryFontLineHeight:   hubspotsdk.String("primaryFontLineHeight"),
						PrimaryFontSize:         hubspotsdk.Float(0),
						SecondaryAccentColor:    hubspotsdk.String("secondaryAccentColor"),
						SecondaryFont:           hubspotsdk.String("secondaryFont"),
						SecondaryFontColor:      hubspotsdk.String("secondaryFontColor"),
						SecondaryFontLineHeight: hubspotsdk.String("secondaryFontLineHeight"),
						SecondaryFontSize:       hubspotsdk.Float(0),
					},
					TemplatePath: hubspotsdk.String("templatePath"),
					ThemeSettingsValues: map[string]any{
						"foo": map[string]any{},
					},
					WidgetContainers: map[string]any{
						"foo": map[string]any{},
					},
					Widgets: map[string]any{
						"foo": map[string]any{},
					},
				},
				FolderIDV2: hubspotsdk.Int(0),
				From: marketing.PublicEmailFromDetailsParam{
					CustomReplyTo: hubspotsdk.String("customReplyTo"),
					FromName:      hubspotsdk.String("fromName"),
					ReplyTo:       hubspotsdk.String("replyTo"),
				},
				JitterSendTime: hubspotsdk.Bool(true),
				Language:       marketing.EmailUpdateRequestLanguageAa,
				Name:           hubspotsdk.String("name"),
				PublishDate:    hubspotsdk.Time(time.Now()),
				RssData: marketing.PublicRssEmailDetailsParam{
					BlogEmailType:     hubspotsdk.String("blogEmailType"),
					BlogImageMaxWidth: hubspotsdk.Int(0),
					BlogLayout:        marketing.PublicRssEmailDetailsBlogLayoutFullPost,
					HubSpotBlogID:     hubspotsdk.String("hubspotBlogId"),
					MaxEntries:        hubspotsdk.Int(0),
					RssEntryTemplate:  hubspotsdk.String("rssEntryTemplate"),
					Timing: map[string]any{
						"foo": map[string]any{},
					},
					URL:                  hubspotsdk.String("url"),
					UseHeadlineAsSubject: hubspotsdk.Bool(true),
				},
				SendOnPublish: hubspotsdk.Bool(true),
				State:         marketing.EmailUpdateRequestStateAgentGenerated,
				Subcategory:   marketing.EmailUpdateRequestSubcategoryAbLoserVariant,
				Subject:       hubspotsdk.String("subject"),
				SubscriptionDetails: marketing.PublicEmailSubscriptionDetailsParam{
					OfficeLocationID:   hubspotsdk.String("officeLocationId"),
					PreferencesGroupID: hubspotsdk.String("preferencesGroupId"),
					SubscriptionID:     hubspotsdk.String("subscriptionId"),
					SubscriptionName:   hubspotsdk.String("subscriptionName"),
				},
				Testing: marketing.PublicEmailTestingDetailsParam{
					IsAbVariation:       true,
					AbSampleSizeDefault: marketing.PublicEmailTestingDetailsAbSampleSizeDefaultAutomatedLoserVariant,
					AbSamplingDefault:   marketing.PublicEmailTestingDetailsAbSamplingDefaultAutomatedLoserVariant,
					AbStatus:            marketing.PublicEmailTestingDetailsAbStatusAutomatedLoserVariant,
					AbSuccessMetric:     marketing.PublicEmailTestingDetailsAbSuccessMetricClicksByDelivered,
					AbTestPercentage:    hubspotsdk.Int(0),
					HoursToWait:         hubspotsdk.Int(0),
					TestID:              hubspotsdk.String("testId"),
				},
				To: marketing.PublicEmailToDetailsParam{
					ContactIDs: marketing.PublicEmailRecipientsParam{
						Exclude: []string{"string"},
						Include: []string{"string"},
					},
					ContactIlsLists: marketing.PublicEmailRecipientsParam{
						Exclude: []string{"string"},
						Include: []string{"string"},
					},
					ContactLists: marketing.PublicEmailRecipientsParam{
						Exclude: []string{"string"},
						Include: []string{"string"},
					},
					LimitSendFrequency: hubspotsdk.Bool(true),
					SuppressGraymail:   hubspotsdk.Bool(true),
				},
				Webversion: marketing.PublicWebversionDetailsParam{
					Domain:            hubspotsdk.String("domain"),
					Enabled:           hubspotsdk.Bool(true),
					ExpiresAt:         hubspotsdk.Time(time.Now()),
					IsPageRedirected:  hubspotsdk.Bool(true),
					MetaDescription:   hubspotsdk.String("metaDescription"),
					PageExpiryEnabled: hubspotsdk.Bool(true),
					RedirectToPageID:  hubspotsdk.String("redirectToPageId"),
					RedirectToURL:     hubspotsdk.String("redirectToUrl"),
					Slug:              hubspotsdk.String("slug"),
					Title:             hubspotsdk.String("title"),
					URL:               hubspotsdk.String("url"),
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

func TestEmailListWithOptionalParams(t *testing.T) {
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
	_, err := client.Marketing.Emails.List(context.TODO(), marketing.EmailListParams{
		After:                  hubspotsdk.String("after"),
		Archived:               hubspotsdk.Bool(true),
		Campaign:               hubspotsdk.String("campaign"),
		CreatedAfter:           hubspotsdk.Time(time.Now()),
		CreatedAt:              hubspotsdk.Time(time.Now()),
		CreatedBefore:          hubspotsdk.Time(time.Now()),
		IncludedProperties:     []string{"string"},
		IncludeStats:           hubspotsdk.Bool(true),
		IsPublished:            hubspotsdk.Bool(true),
		Limit:                  hubspotsdk.Int(0),
		MarketingCampaignNames: hubspotsdk.Bool(true),
		PublishedAfter:         hubspotsdk.Time(time.Now()),
		PublishedAt:            hubspotsdk.Time(time.Now()),
		PublishedBefore:        hubspotsdk.Time(time.Now()),
		Sort:                   []string{"string"},
		Type:                   marketing.EmailListParamsTypeAbEmail,
		UpdatedAfter:           hubspotsdk.Time(time.Now()),
		UpdatedAt:              hubspotsdk.Time(time.Now()),
		UpdatedBefore:          hubspotsdk.Time(time.Now()),
		VariantStats:           hubspotsdk.Bool(true),
		WorkflowNames:          hubspotsdk.Bool(true),
	})
	if err != nil {
		var apierr *hubspotsdk.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestEmailDeleteWithOptionalParams(t *testing.T) {
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
	err := client.Marketing.Emails.Delete(
		context.TODO(),
		"emailId",
		marketing.EmailDeleteParams{
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

func TestEmailCloneWithOptionalParams(t *testing.T) {
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
	_, err := client.Marketing.Emails.Clone(context.TODO(), marketing.EmailCloneParams{
		EmailCloneRequestVNext: marketing.EmailCloneRequestVNextParam{
			ID:        "id",
			CloneName: hubspotsdk.String("cloneName"),
			Language:  hubspotsdk.String("language"),
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

func TestEmailNewAbTestVariation(t *testing.T) {
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
	_, err := client.Marketing.Emails.NewAbTestVariation(context.TODO(), marketing.EmailNewAbTestVariationParams{
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

func TestEmailGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Marketing.Emails.Get(context.TODO(), marketing.EmailGetParams{
		EmailIDs:       []int64{0},
		EndTimestamp:   hubspotsdk.Time(time.Now()),
		Property:       hubspotsdk.String("property"),
		StartTimestamp: hubspotsdk.Time(time.Now()),
	})
	if err != nil {
		var apierr *hubspotsdk.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestEmailGetAbTestVariationWithOptionalParams(t *testing.T) {
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
	_, err := client.Marketing.Emails.GetAbTestVariation(
		context.TODO(),
		"emailId",
		marketing.EmailGetAbTestVariationParams{
			Archived:               hubspotsdk.Bool(true),
			IncludedProperties:     []string{"string"},
			IncludeStats:           hubspotsdk.Bool(true),
			MarketingCampaignNames: hubspotsdk.Bool(true),
			VariantStats:           hubspotsdk.Bool(true),
			WorkflowNames:          hubspotsdk.Bool(true),
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

func TestEmailGetDraft(t *testing.T) {
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
	_, err := client.Marketing.Emails.GetDraft(context.TODO(), "emailId")
	if err != nil {
		var apierr *hubspotsdk.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestEmailGetHistogramWithOptionalParams(t *testing.T) {
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
	_, err := client.Marketing.Emails.GetHistogram(context.TODO(), marketing.EmailGetHistogramParams{
		EmailIDs:       []int64{0},
		EndTimestamp:   hubspotsdk.Time(time.Now()),
		Interval:       marketing.EmailGetHistogramParamsIntervalDay,
		StartTimestamp: hubspotsdk.Time(time.Now()),
	})
	if err != nil {
		var apierr *hubspotsdk.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestEmailGetRevision(t *testing.T) {
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
	_, err := client.Marketing.Emails.GetRevision(
		context.TODO(),
		"revisionId",
		marketing.EmailGetRevisionParams{
			EmailID: "emailId",
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

func TestEmailListRevisionsWithOptionalParams(t *testing.T) {
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
	_, err := client.Marketing.Emails.ListRevisions(
		context.TODO(),
		"emailId",
		marketing.EmailListRevisionsParams{
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

func TestEmailPublish(t *testing.T) {
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
	err := client.Marketing.Emails.Publish(context.TODO(), "emailId")
	if err != nil {
		var apierr *hubspotsdk.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestEmailResetDraft(t *testing.T) {
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
	err := client.Marketing.Emails.ResetDraft(context.TODO(), "emailId")
	if err != nil {
		var apierr *hubspotsdk.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestEmailRestoreRevision(t *testing.T) {
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
	err := client.Marketing.Emails.RestoreRevision(
		context.TODO(),
		"revisionId",
		marketing.EmailRestoreRevisionParams{
			EmailID: "emailId",
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

func TestEmailRestoreRevisionToDraft(t *testing.T) {
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
	_, err := client.Marketing.Emails.RestoreRevisionToDraft(
		context.TODO(),
		0,
		marketing.EmailRestoreRevisionToDraftParams{
			EmailID: "emailId",
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

func TestEmailUnpublish(t *testing.T) {
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
	err := client.Marketing.Emails.Unpublish(context.TODO(), "emailId")
	if err != nil {
		var apierr *hubspotsdk.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestEmailUpdateDraftWithOptionalParams(t *testing.T) {
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
	_, err := client.Marketing.Emails.UpdateDraft(
		context.TODO(),
		"emailId",
		marketing.EmailUpdateDraftParams{
			EmailUpdateRequest: marketing.EmailUpdateRequestParam{
				ActiveDomain:   hubspotsdk.String("activeDomain"),
				Archived:       hubspotsdk.Bool(true),
				BusinessUnitID: hubspotsdk.Int(0),
				Campaign:       hubspotsdk.String("campaign"),
				Content: marketing.PublicEmailContentParam{
					FlexAreas: map[string]any{
						"foo": map[string]any{},
					},
					PlainTextVersion: hubspotsdk.String("plainTextVersion"),
					SmartFields: map[string]marketing.SmartEmailField{
						"foo": map[string]any{},
					},
					StyleSettings: marketing.PublicEmailStyleSettingsParam{
						BackgroundColor:       hubspotsdk.String("backgroundColor"),
						BackgroundImage:       hubspotsdk.String("backgroundImage"),
						BackgroundImageType:   marketing.PublicEmailStyleSettingsBackgroundImageTypeRepeat,
						BodyBorderColor:       hubspotsdk.String("bodyBorderColor"),
						BodyBorderColorChoice: hubspotsdk.String("bodyBorderColorChoice"),
						BodyBorderWidth:       hubspotsdk.Float(0),
						BodyColor:             hubspotsdk.String("bodyColor"),
						ButtonStyleSettings: marketing.PublicButtonStyleSettingsParam{
							BackgroundColor: map[string]any{},
							CornerRadius:    hubspotsdk.Int(0),
							FontStyle: marketing.PublicFontStyleParam{
								Bold:      hubspotsdk.Bool(true),
								Color:     hubspotsdk.String("color"),
								Font:      hubspotsdk.String("font"),
								Italic:    hubspotsdk.Bool(true),
								Size:      hubspotsdk.Int(0),
								Underline: hubspotsdk.Bool(true),
							},
						},
						ColorPickerFavorite1: hubspotsdk.String("colorPickerFavorite1"),
						ColorPickerFavorite2: hubspotsdk.String("colorPickerFavorite2"),
						ColorPickerFavorite3: hubspotsdk.String("colorPickerFavorite3"),
						ColorPickerFavorite4: hubspotsdk.String("colorPickerFavorite4"),
						ColorPickerFavorite5: hubspotsdk.String("colorPickerFavorite5"),
						ColorPickerFavorite6: hubspotsdk.String("colorPickerFavorite6"),
						DividerStyleSettings: marketing.PublicDividerStyleSettingsParam{
							Color:    map[string]any{},
							Height:   hubspotsdk.Int(0),
							LineType: hubspotsdk.String("lineType"),
						},
						EmailBodyPadding: hubspotsdk.String("emailBodyPadding"),
						EmailBodyWidth:   hubspotsdk.String("emailBodyWidth"),
						HeadingOneFont: marketing.PublicFontStyleParam{
							Bold:      hubspotsdk.Bool(true),
							Color:     hubspotsdk.String("color"),
							Font:      hubspotsdk.String("font"),
							Italic:    hubspotsdk.Bool(true),
							Size:      hubspotsdk.Int(0),
							Underline: hubspotsdk.Bool(true),
						},
						HeadingTwoFont: marketing.PublicFontStyleParam{
							Bold:      hubspotsdk.Bool(true),
							Color:     hubspotsdk.String("color"),
							Font:      hubspotsdk.String("font"),
							Italic:    hubspotsdk.Bool(true),
							Size:      hubspotsdk.Int(0),
							Underline: hubspotsdk.Bool(true),
						},
						LinksFont: marketing.PublicFontStyleParam{
							Bold:      hubspotsdk.Bool(true),
							Color:     hubspotsdk.String("color"),
							Font:      hubspotsdk.String("font"),
							Italic:    hubspotsdk.Bool(true),
							Size:      hubspotsdk.Int(0),
							Underline: hubspotsdk.Bool(true),
						},
						PrimaryAccentColor:      hubspotsdk.String("primaryAccentColor"),
						PrimaryFont:             hubspotsdk.String("primaryFont"),
						PrimaryFontColor:        hubspotsdk.String("primaryFontColor"),
						PrimaryFontLineHeight:   hubspotsdk.String("primaryFontLineHeight"),
						PrimaryFontSize:         hubspotsdk.Float(0),
						SecondaryAccentColor:    hubspotsdk.String("secondaryAccentColor"),
						SecondaryFont:           hubspotsdk.String("secondaryFont"),
						SecondaryFontColor:      hubspotsdk.String("secondaryFontColor"),
						SecondaryFontLineHeight: hubspotsdk.String("secondaryFontLineHeight"),
						SecondaryFontSize:       hubspotsdk.Float(0),
					},
					TemplatePath: hubspotsdk.String("templatePath"),
					ThemeSettingsValues: map[string]any{
						"foo": map[string]any{},
					},
					WidgetContainers: map[string]any{
						"foo": map[string]any{},
					},
					Widgets: map[string]any{
						"foo": map[string]any{},
					},
				},
				FolderIDV2: hubspotsdk.Int(0),
				From: marketing.PublicEmailFromDetailsParam{
					CustomReplyTo: hubspotsdk.String("customReplyTo"),
					FromName:      hubspotsdk.String("fromName"),
					ReplyTo:       hubspotsdk.String("replyTo"),
				},
				JitterSendTime: hubspotsdk.Bool(true),
				Language:       marketing.EmailUpdateRequestLanguageAa,
				Name:           hubspotsdk.String("name"),
				PublishDate:    hubspotsdk.Time(time.Now()),
				RssData: marketing.PublicRssEmailDetailsParam{
					BlogEmailType:     hubspotsdk.String("blogEmailType"),
					BlogImageMaxWidth: hubspotsdk.Int(0),
					BlogLayout:        marketing.PublicRssEmailDetailsBlogLayoutFullPost,
					HubSpotBlogID:     hubspotsdk.String("hubspotBlogId"),
					MaxEntries:        hubspotsdk.Int(0),
					RssEntryTemplate:  hubspotsdk.String("rssEntryTemplate"),
					Timing: map[string]any{
						"foo": map[string]any{},
					},
					URL:                  hubspotsdk.String("url"),
					UseHeadlineAsSubject: hubspotsdk.Bool(true),
				},
				SendOnPublish: hubspotsdk.Bool(true),
				State:         marketing.EmailUpdateRequestStateAgentGenerated,
				Subcategory:   marketing.EmailUpdateRequestSubcategoryAbLoserVariant,
				Subject:       hubspotsdk.String("subject"),
				SubscriptionDetails: marketing.PublicEmailSubscriptionDetailsParam{
					OfficeLocationID:   hubspotsdk.String("officeLocationId"),
					PreferencesGroupID: hubspotsdk.String("preferencesGroupId"),
					SubscriptionID:     hubspotsdk.String("subscriptionId"),
					SubscriptionName:   hubspotsdk.String("subscriptionName"),
				},
				Testing: marketing.PublicEmailTestingDetailsParam{
					IsAbVariation:       true,
					AbSampleSizeDefault: marketing.PublicEmailTestingDetailsAbSampleSizeDefaultAutomatedLoserVariant,
					AbSamplingDefault:   marketing.PublicEmailTestingDetailsAbSamplingDefaultAutomatedLoserVariant,
					AbStatus:            marketing.PublicEmailTestingDetailsAbStatusAutomatedLoserVariant,
					AbSuccessMetric:     marketing.PublicEmailTestingDetailsAbSuccessMetricClicksByDelivered,
					AbTestPercentage:    hubspotsdk.Int(0),
					HoursToWait:         hubspotsdk.Int(0),
					TestID:              hubspotsdk.String("testId"),
				},
				To: marketing.PublicEmailToDetailsParam{
					ContactIDs: marketing.PublicEmailRecipientsParam{
						Exclude: []string{"string"},
						Include: []string{"string"},
					},
					ContactIlsLists: marketing.PublicEmailRecipientsParam{
						Exclude: []string{"string"},
						Include: []string{"string"},
					},
					ContactLists: marketing.PublicEmailRecipientsParam{
						Exclude: []string{"string"},
						Include: []string{"string"},
					},
					LimitSendFrequency: hubspotsdk.Bool(true),
					SuppressGraymail:   hubspotsdk.Bool(true),
				},
				Webversion: marketing.PublicWebversionDetailsParam{
					Domain:            hubspotsdk.String("domain"),
					Enabled:           hubspotsdk.Bool(true),
					ExpiresAt:         hubspotsdk.Time(time.Now()),
					IsPageRedirected:  hubspotsdk.Bool(true),
					MetaDescription:   hubspotsdk.String("metaDescription"),
					PageExpiryEnabled: hubspotsdk.Bool(true),
					RedirectToPageID:  hubspotsdk.String("redirectToPageId"),
					RedirectToURL:     hubspotsdk.String("redirectToUrl"),
					Slug:              hubspotsdk.String("slug"),
					Title:             hubspotsdk.String("title"),
					URL:               hubspotsdk.String("url"),
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
