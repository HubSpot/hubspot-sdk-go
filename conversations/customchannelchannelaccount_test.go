// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package conversations_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/stainless-sdks/hubspot-sdk-go"
	"github.com/stainless-sdks/hubspot-sdk-go/conversations"
	"github.com/stainless-sdks/hubspot-sdk-go/internal/testutil"
	"github.com/stainless-sdks/hubspot-sdk-go/option"
)

func TestCustomChannelChannelAccountNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Conversations.CustomChannels.ChannelAccounts.New(
		context.TODO(),
		0,
		conversations.CustomChannelChannelAccountNewParams{
			PublicChannelAccountEgg: conversations.PublicChannelAccountEggParam{
				Authorized: true,
				InboxID:    "inboxId",
				Name:       "name",
				DeliveryIdentifier: conversations.PublicDeliveryIdentifierParam{
					Type:  conversations.PublicDeliveryIdentifierTypeChannelSpecificOpaqueID,
					Value: "value",
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

func TestCustomChannelChannelAccountUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Conversations.CustomChannels.ChannelAccounts.Update(
		context.TODO(),
		0,
		conversations.CustomChannelChannelAccountUpdateParams{
			ChannelID: 0,
			PublicChannelAccountUpdateRequest: conversations.PublicChannelAccountUpdateRequestParam{
				Authorized: hubspotsdk.Bool(true),
				Name:       hubspotsdk.String("name"),
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

func TestCustomChannelChannelAccountListWithOptionalParams(t *testing.T) {
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
	_, err := client.Conversations.CustomChannels.ChannelAccounts.List(
		context.TODO(),
		0,
		conversations.CustomChannelChannelAccountListParams{
			After:                   hubspotsdk.String("after"),
			Archived:                hubspotsdk.Bool(true),
			DefaultPageLength:       hubspotsdk.Int(0),
			DeliveryIdentifierType:  []string{"HS_EMAIL_ADDRESS"},
			DeliveryIdentifierValue: []string{"string"},
			Limit:                   hubspotsdk.Int(0),
			Sort:                    []string{"string"},
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

func TestCustomChannelChannelAccountUpdateStagingTokenWithOptionalParams(t *testing.T) {
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
	_, err := client.Conversations.CustomChannels.ChannelAccounts.UpdateStagingToken(
		context.TODO(),
		"accountToken",
		conversations.CustomChannelChannelAccountUpdateStagingTokenParams{
			ChannelID: 0,
			PublicChannelAccountStagingTokenUpdateRequest: conversations.PublicChannelAccountStagingTokenUpdateRequestParam{
				AccountName: hubspotsdk.String("accountName"),
				DeliveryIdentifier: conversations.PublicDeliveryIdentifierParam{
					Type:  conversations.PublicDeliveryIdentifierTypeChannelSpecificOpaqueID,
					Value: "value",
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
