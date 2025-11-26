// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package marketing_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/stainless-sdks/hubspot-sdk-go"
	"github.com/stainless-sdks/hubspot-sdk-go/internal/testutil"
	"github.com/stainless-sdks/hubspot-sdk-go/marketing"
	"github.com/stainless-sdks/hubspot-sdk-go/option"
	"github.com/stainless-sdks/hubspot-sdk-go/shared"
)

func TestEmailNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Marketing.Emails.New(context.TODO(), marketing.EmailNewParams{
		EmailCreateRequest: marketing.EmailCreateRequestParam{
			Name:           "My subject",
			ActiveDomain:   hubspotsdk.String("test.hs-sites.com"),
			Archived:       hubspotsdk.Bool(false),
			BusinessUnitID: hubspotsdk.Int(0),
			Campaign:       hubspotsdk.String("1b7f51a6-33c1-44d6-ba28-fe81f655dced"),
			Content: marketing.PublicEmailContentParam{
				FlexAreas: map[string]any{
					"main": map[string]any{},
				},
				PlainTextVersion: hubspotsdk.String("This is custom! View in browser ({{view_as_page_url}})\n\nHello {{ contact.firstname }},\n\nPlain text emails have minimal formatting so your reader can really focus on what you have to say. Introduce yourself and explain why you’re reaching out.\n\nEvery email should try to lead the reader to some kind of action. Use this space to describe why the reader should want to click on the link below. Put the link on its own line to really draw their eye to it.\n\nLink text\n\nNow it’s time to wrap up your email. Before your signature, thank the recipient for reading. You can also invite them to send this email to any of their colleagues who might be interested.\n\nAll the best,\n\nYour full name\n\nYour job title\n\nOther contact information\n\n{{site_settings.company_name}}, {{site_settings.company_street_address_1}}, {{site_settings.company_street_address_2}}, {{site_settings.company_city}}, {{site_settings.company_state}} {{site_settings.company_zip}}, {{site_settings.company_country}}, {{site_settings.company_phone}}\n\nUnsubscribe ({{unsubscribe_link_all}})\n\nManage preferences ({{unsubscribe_link}})"),
				SmartFields: map[string]marketing.SmartEmailField{
					"foo": map[string]any{},
				},
				StyleSettings: marketing.PublicEmailStyleSettingsParam{
					BackgroundColor:       hubspotsdk.String("backgroundColor"),
					BackgroundImage:       hubspotsdk.String("backgroundImage"),
					BackgroundImageType:   hubspotsdk.String("backgroundImageType"),
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
					"module-0-1-1":           map[string]any{},
					"module-1-1-1":           map[string]any{},
					"module_160676180617911": map[string]any{},
					"preview_text":           map[string]any{},
				},
			},
			FeedbackSurveyID: hubspotsdk.String("feedbackSurveyId"),
			FolderIDV2:       hubspotsdk.Int(0),
			From: marketing.PublicEmailFromDetailsParam{
				CustomReplyTo: hubspotsdk.String("customReplyTo"),
				FromName:      hubspotsdk.String("Bruce Wayne"),
				ReplyTo:       hubspotsdk.String("test@hubspot.com"),
			},
			JitterSendTime: hubspotsdk.Bool(true),
			Language:       marketing.EmailCreateRequestLanguageAf,
			PublishDate:    hubspotsdk.Time(time.Now()),
			RssData: marketing.PublicRssEmailDetailsParam{
				BlogEmailType:     hubspotsdk.String("blogEmailType"),
				BlogImageMaxWidth: hubspotsdk.Int(0),
				BlogLayout:        hubspotsdk.String("blogLayout"),
				HubspotBlogID:     hubspotsdk.String("hubspotBlogId"),
				MaxEntries:        hubspotsdk.Int(0),
				RssEntryTemplate:  hubspotsdk.String("rssEntryTemplate"),
				Timing: map[string]any{
					"foo": map[string]any{},
				},
				URL:                  hubspotsdk.String("url"),
				UseHeadlineAsSubject: hubspotsdk.Bool(true),
			},
			SendOnPublish: hubspotsdk.Bool(true),
			State:         marketing.EmailCreateRequestStateDraft,
			Subcategory:   marketing.EmailCreateRequestSubcategoryBatch,
			Subject:       hubspotsdk.String("My subject"),
			SubscriptionDetails: marketing.PublicEmailSubscriptionDetailsParam{
				OfficeLocationID:   hubspotsdk.String("5449392956"),
				PreferencesGroupID: hubspotsdk.String("preferencesGroupId"),
				SubscriptionID:     hubspotsdk.String("subscriptionId"),
				SubscriptionName:   hubspotsdk.String("subscriptionName"),
			},
			Testing: marketing.PublicEmailTestingDetailsParam{
				AbSampleSizeDefault: marketing.PublicEmailTestingDetailsAbSampleSizeDefaultMaster,
				AbSamplingDefault:   marketing.PublicEmailTestingDetailsAbSamplingDefaultMaster,
				AbStatus:            marketing.PublicEmailTestingDetailsAbStatusMaster,
				AbSuccessMetric:     marketing.PublicEmailTestingDetailsAbSuccessMetricClicksByOpens,
				AbTestPercentage:    hubspotsdk.Int(0),
				HoursToWait:         hubspotsdk.Int(0),
				IsAbVariation:       hubspotsdk.Bool(true),
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
				MetaDescription:   hubspotsdk.String(""),
				PageExpiryEnabled: hubspotsdk.Bool(true),
				RedirectToPageID:  hubspotsdk.String("redirectToPageId"),
				RedirectToURL:     hubspotsdk.String("http://www.example.org"),
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
	_, err := client.Marketing.Emails.Update(
		context.TODO(),
		"emailId",
		marketing.EmailUpdateParams{
			EmailUpdateRequest: marketing.EmailUpdateRequestParam{
				ActiveDomain:   hubspotsdk.String("test.hs-sites.com"),
				Archived:       hubspotsdk.Bool(false),
				BusinessUnitID: hubspotsdk.Int(0),
				Campaign:       hubspotsdk.String("1b7f51a6-33c1-44d6-ba28-fe81f655dced"),
				Content: marketing.PublicEmailContentParam{
					FlexAreas: map[string]any{
						"main": map[string]any{},
					},
					PlainTextVersion: hubspotsdk.String("This is custom! View in browser ({{view_as_page_url}})\n\nHello {{ contact.firstname }},\n\nPlain text emails have minimal formatting so your reader can really focus on what you have to say. Introduce yourself and explain why you’re reaching out.\n\nEvery email should try to lead the reader to some kind of action. Use this space to describe why the reader should want to click on the link below. Put the link on its own line to really draw their eye to it.\n\nLink text\n\nNow it’s time to wrap up your email. Before your signature, thank the recipient for reading. You can also invite them to send this email to any of their colleagues who might be interested.\n\nAll the best,\n\nYour full name\n\nYour job title\n\nOther contact information\n\n{{site_settings.company_name}}, {{site_settings.company_street_address_1}}, {{site_settings.company_street_address_2}}, {{site_settings.company_city}}, {{site_settings.company_state}} {{site_settings.company_zip}}, {{site_settings.company_country}}, {{site_settings.company_phone}}\n\nUnsubscribe ({{unsubscribe_link_all}})\n\nManage preferences ({{unsubscribe_link}})"),
					SmartFields: map[string]marketing.SmartEmailField{
						"foo": map[string]any{},
					},
					StyleSettings: marketing.PublicEmailStyleSettingsParam{
						BackgroundColor:       hubspotsdk.String("backgroundColor"),
						BackgroundImage:       hubspotsdk.String("backgroundImage"),
						BackgroundImageType:   hubspotsdk.String("backgroundImageType"),
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
						"module-0-1-1":           map[string]any{},
						"module-1-1-1":           map[string]any{},
						"module_160676180617911": map[string]any{},
						"preview_text":           map[string]any{},
					},
				},
				FolderIDV2: hubspotsdk.Int(0),
				From: marketing.PublicEmailFromDetailsParam{
					CustomReplyTo: hubspotsdk.String("customReplyTo"),
					FromName:      hubspotsdk.String("Bruce Wayne"),
					ReplyTo:       hubspotsdk.String("test@hubspot.com"),
				},
				JitterSendTime: hubspotsdk.Bool(true),
				Language:       marketing.EmailUpdateRequestLanguageAf,
				Name:           hubspotsdk.String("My subject"),
				PublishDate:    hubspotsdk.Time(time.Now()),
				RssData: marketing.PublicRssEmailDetailsParam{
					BlogEmailType:     hubspotsdk.String("blogEmailType"),
					BlogImageMaxWidth: hubspotsdk.Int(0),
					BlogLayout:        hubspotsdk.String("blogLayout"),
					HubspotBlogID:     hubspotsdk.String("hubspotBlogId"),
					MaxEntries:        hubspotsdk.Int(0),
					RssEntryTemplate:  hubspotsdk.String("rssEntryTemplate"),
					Timing: map[string]any{
						"foo": map[string]any{},
					},
					URL:                  hubspotsdk.String("url"),
					UseHeadlineAsSubject: hubspotsdk.Bool(true),
				},
				SendOnPublish: hubspotsdk.Bool(true),
				State:         marketing.EmailUpdateRequestStateDraft,
				Subcategory:   marketing.EmailUpdateRequestSubcategoryBatch,
				Subject:       hubspotsdk.String("My subject"),
				SubscriptionDetails: marketing.PublicEmailSubscriptionDetailsParam{
					OfficeLocationID:   hubspotsdk.String("5449392956"),
					PreferencesGroupID: hubspotsdk.String("preferencesGroupId"),
					SubscriptionID:     hubspotsdk.String("subscriptionId"),
					SubscriptionName:   hubspotsdk.String("subscriptionName"),
				},
				Testing: marketing.PublicEmailTestingDetailsParam{
					AbSampleSizeDefault: marketing.PublicEmailTestingDetailsAbSampleSizeDefaultMaster,
					AbSamplingDefault:   marketing.PublicEmailTestingDetailsAbSamplingDefaultMaster,
					AbStatus:            marketing.PublicEmailTestingDetailsAbStatusMaster,
					AbSuccessMetric:     marketing.PublicEmailTestingDetailsAbSuccessMetricClicksByOpens,
					AbTestPercentage:    hubspotsdk.Int(0),
					HoursToWait:         hubspotsdk.Int(0),
					IsAbVariation:       hubspotsdk.Bool(true),
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
					MetaDescription:   hubspotsdk.String(""),
					PageExpiryEnabled: hubspotsdk.Bool(true),
					RedirectToPageID:  hubspotsdk.String("redirectToPageId"),
					RedirectToURL:     hubspotsdk.String("http://www.example.org"),
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
		Sort:                   []string{"string"},
		Type:                   marketing.EmailListParamsTypeAbEmail,
		UpdatedAfter:           hubspotsdk.Time(time.Now()),
		UpdatedAt:              hubspotsdk.Time(time.Now()),
		UpdatedBefore:          hubspotsdk.Time(time.Now()),
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
	_, err := client.Marketing.Emails.Get(
		context.TODO(),
		"emailId",
		marketing.EmailGetParams{
			Archived:               hubspotsdk.Bool(true),
			IncludedProperties:     []string{"string"},
			IncludeStats:           hubspotsdk.Bool(true),
			MarketingCampaignNames: hubspotsdk.Bool(true),
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

func TestEmailGetAbTestVariation(t *testing.T) {
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
	_, err := client.Marketing.Emails.GetAbTestVariation(context.TODO(), "emailId")
	if err != nil {
		var apierr *hubspotsdk.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestEmailGetDraft(t *testing.T) {
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
	_, err := client.Marketing.Emails.GetDraft(context.TODO(), "emailId")
	if err != nil {
		var apierr *hubspotsdk.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestEmailGetRevision(t *testing.T) {
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
	_, err := client.Marketing.Emails.UpdateDraft(
		context.TODO(),
		"emailId",
		marketing.EmailUpdateDraftParams{
			EmailUpdateRequest: marketing.EmailUpdateRequestParam{
				ActiveDomain:   hubspotsdk.String("test.hs-sites.com"),
				Archived:       hubspotsdk.Bool(false),
				BusinessUnitID: hubspotsdk.Int(0),
				Campaign:       hubspotsdk.String("1b7f51a6-33c1-44d6-ba28-fe81f655dced"),
				Content: marketing.PublicEmailContentParam{
					FlexAreas: map[string]any{
						"main": map[string]any{},
					},
					PlainTextVersion: hubspotsdk.String("This is custom! View in browser ({{view_as_page_url}})\n\nHello {{ contact.firstname }},\n\nPlain text emails have minimal formatting so your reader can really focus on what you have to say. Introduce yourself and explain why you’re reaching out.\n\nEvery email should try to lead the reader to some kind of action. Use this space to describe why the reader should want to click on the link below. Put the link on its own line to really draw their eye to it.\n\nLink text\n\nNow it’s time to wrap up your email. Before your signature, thank the recipient for reading. You can also invite them to send this email to any of their colleagues who might be interested.\n\nAll the best,\n\nYour full name\n\nYour job title\n\nOther contact information\n\n{{site_settings.company_name}}, {{site_settings.company_street_address_1}}, {{site_settings.company_street_address_2}}, {{site_settings.company_city}}, {{site_settings.company_state}} {{site_settings.company_zip}}, {{site_settings.company_country}}, {{site_settings.company_phone}}\n\nUnsubscribe ({{unsubscribe_link_all}})\n\nManage preferences ({{unsubscribe_link}})"),
					SmartFields: map[string]marketing.SmartEmailField{
						"foo": map[string]any{},
					},
					StyleSettings: marketing.PublicEmailStyleSettingsParam{
						BackgroundColor:       hubspotsdk.String("backgroundColor"),
						BackgroundImage:       hubspotsdk.String("backgroundImage"),
						BackgroundImageType:   hubspotsdk.String("backgroundImageType"),
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
						"module-0-1-1":           map[string]any{},
						"module-1-1-1":           map[string]any{},
						"module_160676180617911": map[string]any{},
						"preview_text":           map[string]any{},
					},
				},
				FolderIDV2: hubspotsdk.Int(0),
				From: marketing.PublicEmailFromDetailsParam{
					CustomReplyTo: hubspotsdk.String("customReplyTo"),
					FromName:      hubspotsdk.String("Bruce Wayne"),
					ReplyTo:       hubspotsdk.String("test@hubspot.com"),
				},
				JitterSendTime: hubspotsdk.Bool(true),
				Language:       marketing.EmailUpdateRequestLanguageAf,
				Name:           hubspotsdk.String("My subject"),
				PublishDate:    hubspotsdk.Time(time.Now()),
				RssData: marketing.PublicRssEmailDetailsParam{
					BlogEmailType:     hubspotsdk.String("blogEmailType"),
					BlogImageMaxWidth: hubspotsdk.Int(0),
					BlogLayout:        hubspotsdk.String("blogLayout"),
					HubspotBlogID:     hubspotsdk.String("hubspotBlogId"),
					MaxEntries:        hubspotsdk.Int(0),
					RssEntryTemplate:  hubspotsdk.String("rssEntryTemplate"),
					Timing: map[string]any{
						"foo": map[string]any{},
					},
					URL:                  hubspotsdk.String("url"),
					UseHeadlineAsSubject: hubspotsdk.Bool(true),
				},
				SendOnPublish: hubspotsdk.Bool(true),
				State:         marketing.EmailUpdateRequestStateDraft,
				Subcategory:   marketing.EmailUpdateRequestSubcategoryBatch,
				Subject:       hubspotsdk.String("My subject"),
				SubscriptionDetails: marketing.PublicEmailSubscriptionDetailsParam{
					OfficeLocationID:   hubspotsdk.String("5449392956"),
					PreferencesGroupID: hubspotsdk.String("preferencesGroupId"),
					SubscriptionID:     hubspotsdk.String("subscriptionId"),
					SubscriptionName:   hubspotsdk.String("subscriptionName"),
				},
				Testing: marketing.PublicEmailTestingDetailsParam{
					AbSampleSizeDefault: marketing.PublicEmailTestingDetailsAbSampleSizeDefaultMaster,
					AbSamplingDefault:   marketing.PublicEmailTestingDetailsAbSamplingDefaultMaster,
					AbStatus:            marketing.PublicEmailTestingDetailsAbStatusMaster,
					AbSuccessMetric:     marketing.PublicEmailTestingDetailsAbSuccessMetricClicksByOpens,
					AbTestPercentage:    hubspotsdk.Int(0),
					HoursToWait:         hubspotsdk.Int(0),
					IsAbVariation:       hubspotsdk.Bool(true),
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
					MetaDescription:   hubspotsdk.String(""),
					PageExpiryEnabled: hubspotsdk.Bool(true),
					RedirectToPageID:  hubspotsdk.String("redirectToPageId"),
					RedirectToURL:     hubspotsdk.String("http://www.example.org"),
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
