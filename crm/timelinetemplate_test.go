// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package crm_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/stainless-sdks/hubspot-sdk-go"
	"github.com/stainless-sdks/hubspot-sdk-go/crm"
	"github.com/stainless-sdks/hubspot-sdk-go/internal/testutil"
	"github.com/stainless-sdks/hubspot-sdk-go/option"
)

func TestTimelineTemplateNewWithOptionalParams(t *testing.T) {
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
	_, err := client.CRM.Timeline.Templates.New(
		context.TODO(),
		0,
		crm.TimelineTemplateNewParams{
			TimelineEventTemplateCreateRequest: crm.TimelineEventTemplateCreateRequestParam{
				Name:       "PetSpot Registration",
				ObjectType: "contacts",
				Tokens: []crm.TimelineEventTemplateTokenParam{{
					Label:              "Pet Name",
					Name:               "petName",
					Type:               crm.TimelineEventTemplateTokenTypeString,
					CreatedAt:          hubspotsdk.Time(time.Now()),
					ObjectPropertyName: hubspotsdk.String("customPropertyPetType"),
					Options: []crm.TimelineEventTemplateTokenOptionParam{{
						Label: "Dog",
						Value: "dog",
					}, {
						Label: "Cat",
						Value: "cat",
					}},
					UpdatedAt: hubspotsdk.Time(time.Now()),
				}, {
					Label:              "Pet Age",
					Name:               "petAge",
					Type:               crm.TimelineEventTemplateTokenTypeNumber,
					CreatedAt:          hubspotsdk.Time(time.Now()),
					ObjectPropertyName: hubspotsdk.String("customPropertyPetType"),
					Options: []crm.TimelineEventTemplateTokenOptionParam{{
						Label: "Dog",
						Value: "dog",
					}, {
						Label: "Cat",
						Value: "cat",
					}},
					UpdatedAt: hubspotsdk.Time(time.Now()),
				}, {
					Label:              "Pet Color",
					Name:               "petColor",
					Type:               crm.TimelineEventTemplateTokenTypeEnumeration,
					CreatedAt:          hubspotsdk.Time(time.Now()),
					ObjectPropertyName: hubspotsdk.String("customPropertyPetType"),
					Options: []crm.TimelineEventTemplateTokenOptionParam{{
						Label: "White",
						Value: "white",
					}, {
						Label: "Black",
						Value: "black",
					}, {
						Label: "Brown",
						Value: "brown",
					}, {
						Label: "Other",
						Value: "other",
					}},
					UpdatedAt: hubspotsdk.Time(time.Now()),
				}},
				DetailTemplate: hubspotsdk.String("Registration occurred at {{#formatDate timestamp}}{{/formatDate}}\n\n#### Questions\n{{#each extraData.questions}}\n  **{{question}}**: {{answer}}\n{{/each}}"),
				HeaderTemplate: hubspotsdk.String("Registered for [{{petName}}](https://my.petspot.com/pets/{{petName}})"),
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

func TestTimelineTemplateUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.CRM.Timeline.Templates.Update(
		context.TODO(),
		"eventTemplateId",
		crm.TimelineTemplateUpdateParams{
			AppID: 0,
			TimelineEventTemplateUpdateRequest: crm.TimelineEventTemplateUpdateRequestParam{
				ID:   "1001298",
				Name: "PetSpot Registration",
				Tokens: []crm.TimelineEventTemplateTokenParam{{
					Label:              "Pet Name",
					Name:               "petName",
					Type:               crm.TimelineEventTemplateTokenTypeString,
					CreatedAt:          hubspotsdk.Time(time.Now()),
					ObjectPropertyName: hubspotsdk.String("firstname"),
					Options: []crm.TimelineEventTemplateTokenOptionParam{{
						Label: "Dog",
						Value: "dog",
					}, {
						Label: "Cat",
						Value: "cat",
					}},
					UpdatedAt: hubspotsdk.Time(time.Now()),
				}, {
					Label:              "Pet Age",
					Name:               "petAge",
					Type:               crm.TimelineEventTemplateTokenTypeNumber,
					CreatedAt:          hubspotsdk.Time(time.Now()),
					ObjectPropertyName: hubspotsdk.String("customPropertyPetType"),
					Options: []crm.TimelineEventTemplateTokenOptionParam{{
						Label: "Dog",
						Value: "dog",
					}, {
						Label: "Cat",
						Value: "cat",
					}},
					UpdatedAt: hubspotsdk.Time(time.Now()),
				}, {
					Label:              "Pet Color",
					Name:               "petColor",
					Type:               crm.TimelineEventTemplateTokenTypeEnumeration,
					CreatedAt:          hubspotsdk.Time(time.Now()),
					ObjectPropertyName: hubspotsdk.String("customPropertyPetType"),
					Options: []crm.TimelineEventTemplateTokenOptionParam{{
						Label: "White",
						Value: "white",
					}, {
						Label: "Black",
						Value: "black",
					}, {
						Label: "Brown",
						Value: "brown",
					}, {
						Label: "Yellow",
						Value: "yellow",
					}, {
						Label: "Other",
						Value: "other",
					}},
					UpdatedAt: hubspotsdk.Time(time.Now()),
				}},
				DetailTemplate: hubspotsdk.String("Registration occurred at {{#formatDate timestamp}}{{/formatDate}}\n\n#### Questions\n{{#each extraData.questions}}\n  **{{question}}**: {{answer}}\n{{/each}}\n\nEDIT"),
				HeaderTemplate: hubspotsdk.String("Registered for [{{petName}}](https://my.petspot.com/pets/{{petName}})"),
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

func TestTimelineTemplateList(t *testing.T) {
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
	_, err := client.CRM.Timeline.Templates.List(context.TODO(), 0)
	if err != nil {
		var apierr *hubspotsdk.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestTimelineTemplateDelete(t *testing.T) {
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
	err := client.CRM.Timeline.Templates.Delete(
		context.TODO(),
		"eventTemplateId",
		crm.TimelineTemplateDeleteParams{
			AppID: 0,
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

func TestTimelineTemplateGet(t *testing.T) {
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
	_, err := client.CRM.Timeline.Templates.Get(
		context.TODO(),
		"eventTemplateId",
		crm.TimelineTemplateGetParams{
			AppID: 0,
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
