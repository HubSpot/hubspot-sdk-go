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

func TestTimelineTokenNewWithOptionalParams(t *testing.T) {
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
	_, err := client.CRM.Timeline.Tokens.New(
		context.TODO(),
		"eventTemplateId",
		crm.TimelineTokenNewParams{
			AppID: 0,
			TimelineEventTemplateToken: crm.TimelineEventTemplateTokenParam{
				Label:              "Pet Type",
				Name:               "petType",
				Type:               crm.TimelineEventTemplateTokenTypeEnumeration,
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

func TestTimelineTokenUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.CRM.Timeline.Tokens.Update(
		context.TODO(),
		"tokenName",
		crm.TimelineTokenUpdateParams{
			AppID:           0,
			EventTemplateID: "eventTemplateId",
			TimelineEventTemplateTokenUpdateRequest: crm.TimelineEventTemplateTokenUpdateRequestParam{
				Label:              "petType edit",
				ObjectPropertyName: hubspotsdk.String("objectPropertyName"),
				Options: []crm.TimelineEventTemplateTokenOptionParam{{
					Label: "Dog",
					Value: "dog",
				}, {
					Label: "Cat",
					Value: "cat",
				}, {
					Label: "Bird",
					Value: "bird",
				}},
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

func TestTimelineTokenDelete(t *testing.T) {
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
	err := client.CRM.Timeline.Tokens.Delete(
		context.TODO(),
		"tokenName",
		crm.TimelineTokenDeleteParams{
			AppID:           0,
			EventTemplateID: "eventTemplateId",
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
