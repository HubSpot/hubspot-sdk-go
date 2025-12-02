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
)

func TestFormNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Marketing.Forms.New(context.TODO(), marketing.FormNewParams{
		FormDefinitionCreateRequestBase: marketing.FormDefinitionCreateRequestBaseParam{
			Archived: true,
			Configuration: marketing.HubSpotFormConfigurationParam{
				AllowLinkToResetKnownValues: true,
				Archivable:                  true,
				Cloneable:                   true,
				CreateNewContactForNewEmail: true,
				Editable:                    true,
				Language:                    marketing.HubSpotFormConfigurationLanguageAf,
				NotifyContactOwner:          true,
				NotifyRecipients:            []string{"string"},
				PostSubmitAction: marketing.FormPostSubmitActionParam{
					Type:  marketing.FormPostSubmitActionTypeRedirectURL,
					Value: "value",
				},
				PrePopulateKnownValues: true,
				RecaptchaEnabled:       true,
				LifecycleStages: []marketing.LifecycleStageParam{{
					ObjectTypeID: "objectTypeId",
					Value:        "value",
				}},
			},
			CreatedAt: time.Now(),
			DisplayOptions: marketing.FormDisplayOptionsParam{
				RenderRawHTML: true,
				Style: marketing.FormStyleParam{
					BackgroundWidth:       "backgroundWidth",
					FontFamily:            "fontFamily",
					HelpTextColor:         "helpTextColor",
					HelpTextSize:          "helpTextSize",
					LabelTextColor:        "labelTextColor",
					LabelTextSize:         "labelTextSize",
					LegalConsentTextColor: "legalConsentTextColor",
					LegalConsentTextSize:  "legalConsentTextSize",
					SubmitAlignment:       marketing.FormStyleSubmitAlignmentCenter,
					SubmitColor:           "submitColor",
					SubmitFontColor:       "submitFontColor",
					SubmitSize:            "submitSize",
				},
				SubmitButtonText: "submitButtonText",
				Theme:            marketing.FormDisplayOptionsThemeCanvas,
				CssClass:         hubspotsdk.String("cssClass"),
			},
			FieldGroups: []marketing.FieldGroupParam{{
				Fields: []marketing.FieldGroupFieldUnionParam{{
					OfEmail: &marketing.EmailFieldParam{
						DependentFields: []marketing.DependentFieldParam{{
							DependentCondition: marketing.DependentFieldFilterParam{
								Operator:   marketing.DependentFieldFilterOperatorBetween,
								RangeEnd:   "rangeEnd",
								RangeStart: "rangeStart",
								Value:      "value",
								Values:     []string{"string"},
							},
							DependentField: marketing.DependentFieldDependentFieldUnionParam{
								OfPhone: &marketing.PhoneFieldParam{
									DependentFields:      []marketing.DependentFieldParam{},
									FieldType:            marketing.PhoneFieldFieldTypePhone,
									Hidden:               true,
									Label:                "label",
									Name:                 "name",
									ObjectTypeID:         "objectTypeId",
									Required:             true,
									UseCountryCodeSelect: true,
									Validation: marketing.PhoneFieldValidationParam{
										MaxAllowedDigits: 0,
										MinAllowedDigits: 0,
									},
									DefaultValue: hubspotsdk.String("defaultValue"),
									Description:  hubspotsdk.String("description"),
									Placeholder:  hubspotsdk.String("placeholder"),
								},
							},
						}},
						FieldType:    marketing.EmailFieldFieldTypeEmail,
						Hidden:       true,
						Label:        "label",
						Name:         "name",
						ObjectTypeID: "objectTypeId",
						Required:     true,
						Validation: marketing.EmailFieldValidationParam{
							BlockedEmailDomains: []string{"string"},
							UseDefaultBlockList: true,
						},
						DefaultValue: hubspotsdk.String("defaultValue"),
						Description:  hubspotsdk.String("description"),
						Placeholder:  hubspotsdk.String("placeholder"),
					},
				}},
				GroupType:    marketing.FieldGroupGroupTypeDefaultGroup,
				RichTextType: marketing.FieldGroupRichTextTypeImage,
				RichText:     hubspotsdk.String("richText"),
			}},
			FormType: marketing.FormDefinitionCreateRequestBaseFormTypeHubspot,
			LegalConsentOptions: marketing.FormDefinitionCreateRequestBaseLegalConsentOptionsUnionParam{
				OfNone: &marketing.LegalConsentOptionsNoneParam{
					Type: marketing.LegalConsentOptionsNoneTypeNone,
				},
			},
			Name:       "name",
			UpdatedAt:  time.Now(),
			ArchivedAt: hubspotsdk.Time(time.Now()),
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

func TestFormUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Marketing.Forms.Update(
		context.TODO(),
		"formId",
		marketing.FormUpdateParams{
			HubSpotFormDefinitionPatchRequest: marketing.HubSpotFormDefinitionPatchRequestParam{
				Archived: hubspotsdk.Bool(true),
				Configuration: marketing.HubSpotFormConfigurationParam{
					AllowLinkToResetKnownValues: true,
					Archivable:                  true,
					Cloneable:                   true,
					CreateNewContactForNewEmail: true,
					Editable:                    true,
					Language:                    marketing.HubSpotFormConfigurationLanguageAf,
					NotifyContactOwner:          true,
					NotifyRecipients:            []string{"string"},
					PostSubmitAction: marketing.FormPostSubmitActionParam{
						Type:  marketing.FormPostSubmitActionTypeRedirectURL,
						Value: "value",
					},
					PrePopulateKnownValues: true,
					RecaptchaEnabled:       true,
					LifecycleStages: []marketing.LifecycleStageParam{{
						ObjectTypeID: "objectTypeId",
						Value:        "value",
					}},
				},
				DisplayOptions: marketing.FormDisplayOptionsParam{
					RenderRawHTML: true,
					Style: marketing.FormStyleParam{
						BackgroundWidth:       "backgroundWidth",
						FontFamily:            "fontFamily",
						HelpTextColor:         "helpTextColor",
						HelpTextSize:          "helpTextSize",
						LabelTextColor:        "labelTextColor",
						LabelTextSize:         "labelTextSize",
						LegalConsentTextColor: "legalConsentTextColor",
						LegalConsentTextSize:  "legalConsentTextSize",
						SubmitAlignment:       marketing.FormStyleSubmitAlignmentCenter,
						SubmitColor:           "submitColor",
						SubmitFontColor:       "submitFontColor",
						SubmitSize:            "submitSize",
					},
					SubmitButtonText: "submitButtonText",
					Theme:            marketing.FormDisplayOptionsThemeCanvas,
					CssClass:         hubspotsdk.String("cssClass"),
				},
				FieldGroups: []marketing.FieldGroupParam{{
					Fields: []marketing.FieldGroupFieldUnionParam{{
						OfEmail: &marketing.EmailFieldParam{
							DependentFields: []marketing.DependentFieldParam{{
								DependentCondition: marketing.DependentFieldFilterParam{
									Operator:   marketing.DependentFieldFilterOperatorBetween,
									RangeEnd:   "rangeEnd",
									RangeStart: "rangeStart",
									Value:      "value",
									Values:     []string{"string"},
								},
								DependentField: marketing.DependentFieldDependentFieldUnionParam{
									OfPhone: &marketing.PhoneFieldParam{
										DependentFields:      []marketing.DependentFieldParam{},
										FieldType:            marketing.PhoneFieldFieldTypePhone,
										Hidden:               true,
										Label:                "label",
										Name:                 "name",
										ObjectTypeID:         "objectTypeId",
										Required:             true,
										UseCountryCodeSelect: true,
										Validation: marketing.PhoneFieldValidationParam{
											MaxAllowedDigits: 0,
											MinAllowedDigits: 0,
										},
										DefaultValue: hubspotsdk.String("defaultValue"),
										Description:  hubspotsdk.String("description"),
										Placeholder:  hubspotsdk.String("placeholder"),
									},
								},
							}},
							FieldType:    marketing.EmailFieldFieldTypeEmail,
							Hidden:       true,
							Label:        "label",
							Name:         "name",
							ObjectTypeID: "objectTypeId",
							Required:     true,
							Validation: marketing.EmailFieldValidationParam{
								BlockedEmailDomains: []string{"string"},
								UseDefaultBlockList: true,
							},
							DefaultValue: hubspotsdk.String("defaultValue"),
							Description:  hubspotsdk.String("description"),
							Placeholder:  hubspotsdk.String("placeholder"),
						},
					}},
					GroupType:    marketing.FieldGroupGroupTypeDefaultGroup,
					RichTextType: marketing.FieldGroupRichTextTypeImage,
					RichText:     hubspotsdk.String("richText"),
				}},
				LegalConsentOptions: marketing.HubSpotFormDefinitionPatchRequestLegalConsentOptionsUnionParam{
					OfNone: &marketing.LegalConsentOptionsNoneParam{
						Type: marketing.LegalConsentOptionsNoneTypeNone,
					},
				},
				Name: hubspotsdk.String("name"),
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

func TestFormListWithOptionalParams(t *testing.T) {
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
	_, err := client.Marketing.Forms.List(context.TODO(), marketing.FormListParams{
		After:     hubspotsdk.String("after"),
		Archived:  hubspotsdk.Bool(true),
		FormTypes: []string{"hubspot"},
		Limit:     hubspotsdk.Int(0),
	})
	if err != nil {
		var apierr *hubspotsdk.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestFormDelete(t *testing.T) {
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
	err := client.Marketing.Forms.Delete(context.TODO(), "formId")
	if err != nil {
		var apierr *hubspotsdk.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestFormGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Marketing.Forms.Get(
		context.TODO(),
		"formId",
		marketing.FormGetParams{
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

func TestFormReplaceWithOptionalParams(t *testing.T) {
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
	_, err := client.Marketing.Forms.Replace(
		context.TODO(),
		"formId",
		marketing.FormReplaceParams{
			HubSpotFormDefinition: marketing.HubSpotFormDefinitionParam{
				ID:       "id",
				Archived: true,
				Configuration: marketing.HubSpotFormConfigurationParam{
					AllowLinkToResetKnownValues: true,
					Archivable:                  true,
					Cloneable:                   true,
					CreateNewContactForNewEmail: true,
					Editable:                    true,
					Language:                    marketing.HubSpotFormConfigurationLanguageAf,
					NotifyContactOwner:          true,
					NotifyRecipients:            []string{"string"},
					PostSubmitAction: marketing.FormPostSubmitActionParam{
						Type:  marketing.FormPostSubmitActionTypeRedirectURL,
						Value: "value",
					},
					PrePopulateKnownValues: true,
					RecaptchaEnabled:       true,
					LifecycleStages: []marketing.LifecycleStageParam{{
						ObjectTypeID: "objectTypeId",
						Value:        "value",
					}},
				},
				CreatedAt: time.Now(),
				DisplayOptions: marketing.FormDisplayOptionsParam{
					RenderRawHTML: true,
					Style: marketing.FormStyleParam{
						BackgroundWidth:       "backgroundWidth",
						FontFamily:            "fontFamily",
						HelpTextColor:         "helpTextColor",
						HelpTextSize:          "helpTextSize",
						LabelTextColor:        "labelTextColor",
						LabelTextSize:         "labelTextSize",
						LegalConsentTextColor: "legalConsentTextColor",
						LegalConsentTextSize:  "legalConsentTextSize",
						SubmitAlignment:       marketing.FormStyleSubmitAlignmentCenter,
						SubmitColor:           "submitColor",
						SubmitFontColor:       "submitFontColor",
						SubmitSize:            "submitSize",
					},
					SubmitButtonText: "submitButtonText",
					Theme:            marketing.FormDisplayOptionsThemeCanvas,
					CssClass:         hubspotsdk.String("cssClass"),
				},
				FieldGroups: []marketing.FieldGroupParam{{
					Fields: []marketing.FieldGroupFieldUnionParam{{
						OfEmail: &marketing.EmailFieldParam{
							DependentFields: []marketing.DependentFieldParam{{
								DependentCondition: marketing.DependentFieldFilterParam{
									Operator:   marketing.DependentFieldFilterOperatorBetween,
									RangeEnd:   "rangeEnd",
									RangeStart: "rangeStart",
									Value:      "value",
									Values:     []string{"string"},
								},
								DependentField: marketing.DependentFieldDependentFieldUnionParam{
									OfPhone: &marketing.PhoneFieldParam{
										DependentFields:      []marketing.DependentFieldParam{},
										FieldType:            marketing.PhoneFieldFieldTypePhone,
										Hidden:               true,
										Label:                "label",
										Name:                 "name",
										ObjectTypeID:         "objectTypeId",
										Required:             true,
										UseCountryCodeSelect: true,
										Validation: marketing.PhoneFieldValidationParam{
											MaxAllowedDigits: 0,
											MinAllowedDigits: 0,
										},
										DefaultValue: hubspotsdk.String("defaultValue"),
										Description:  hubspotsdk.String("description"),
										Placeholder:  hubspotsdk.String("placeholder"),
									},
								},
							}},
							FieldType:    marketing.EmailFieldFieldTypeEmail,
							Hidden:       true,
							Label:        "label",
							Name:         "name",
							ObjectTypeID: "objectTypeId",
							Required:     true,
							Validation: marketing.EmailFieldValidationParam{
								BlockedEmailDomains: []string{"string"},
								UseDefaultBlockList: true,
							},
							DefaultValue: hubspotsdk.String("defaultValue"),
							Description:  hubspotsdk.String("description"),
							Placeholder:  hubspotsdk.String("placeholder"),
						},
					}},
					GroupType:    marketing.FieldGroupGroupTypeDefaultGroup,
					RichTextType: marketing.FieldGroupRichTextTypeImage,
					RichText:     hubspotsdk.String("richText"),
				}},
				FormType: marketing.HubSpotFormDefinitionFormTypeHubspot,
				LegalConsentOptions: marketing.HubSpotFormDefinitionLegalConsentOptionsUnionParam{
					OfNone: &marketing.LegalConsentOptionsNoneParam{
						Type: marketing.LegalConsentOptionsNoneTypeNone,
					},
				},
				Name:       "name",
				UpdatedAt:  time.Now(),
				ArchivedAt: hubspotsdk.Time(time.Now()),
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
