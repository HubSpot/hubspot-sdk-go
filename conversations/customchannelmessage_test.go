// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package conversations_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/stainless-sdks/hubspot-sdk-go"
	"github.com/stainless-sdks/hubspot-sdk-go/conversations"
	"github.com/stainless-sdks/hubspot-sdk-go/internal/testutil"
	"github.com/stainless-sdks/hubspot-sdk-go/option"
)

func TestCustomChannelMessageNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Conversations.CustomChannels.Messages.New(
		context.TODO(),
		"channelId",
		conversations.CustomChannelMessageNewParams{
			Attachments: []conversations.CustomChannelMessageNewParamsAttachmentUnion{{
				OfFile: &conversations.CustomChannelMessageNewParamsAttachmentFile{
					FileID:        "fileId",
					Type:          "FILE",
					FileUsageType: hubspotsdk.String("fileUsageType"),
				},
			}},
			ChannelAccountID:    "channelAccountId",
			IntegrationThreadID: "integrationThreadId",
			MessageDirection:    conversations.CustomChannelMessageNewParamsMessageDirectionIncoming,
			Recipients: []conversations.CustomChannelMessageNewParamsRecipient{{
				DeliveryIdentifier: conversations.PublicDeliveryIdentifierParam{
					Type:  "type",
					Value: "value",
				},
				Name: hubspotsdk.String("name"),
			}},
			Senders: []conversations.CustomChannelMessageNewParamsSender{{
				DeliveryIdentifier: conversations.PublicDeliveryIdentifierParam{
					Type:  "type",
					Value: "value",
				},
				Name: hubspotsdk.String("name"),
			}},
			Text:                     "text",
			Timestamp:                time.Now(),
			InReplyToID:              hubspotsdk.String("inReplyToId"),
			IntegrationIdempotencyID: hubspotsdk.String("integrationIdempotencyId"),
			PreResolvedContacts: conversations.CustomChannelMessageNewParamsPreResolvedContacts{
				Contacts: []conversations.CustomChannelMessageNewParamsPreResolvedContactsContact{{
					ContactPropertiesLeadingToMatch: []string{"string"},
					ContactVid:                      0,
				}},
			},
			RichText: hubspotsdk.String("richText"),
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

func TestCustomChannelMessageUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Conversations.CustomChannels.Messages.Update(
		context.TODO(),
		"messageId",
		conversations.CustomChannelMessageUpdateParams{
			ChannelID:    "channelId",
			StatusType:   conversations.CustomChannelMessageUpdateParamsStatusTypeSent,
			ErrorMessage: hubspotsdk.String("errorMessage"),
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

func TestCustomChannelMessageGet(t *testing.T) {
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
	_, err := client.Conversations.CustomChannels.Messages.Get(
		context.TODO(),
		"messageId",
		conversations.CustomChannelMessageGetParams{
			ChannelID: "channelId",
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
