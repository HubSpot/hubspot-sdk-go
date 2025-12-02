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

func TestMessageNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Conversations.Messages.New(
		context.TODO(),
		0,
		conversations.MessageNewParams{
			PublicMessageEgg: conversations.PublicMessageEggUnionParam{
				OfPublicConversationsMessageEgg: &conversations.PublicConversationsMessageEggParam{
					Attachments: []conversations.PublicConversationsMessageEggAttachmentUnionParam{{
						OfFile: &conversations.PublicFileEggParam{
							FileID: "fileId",
							Type:   conversations.PublicFileEggTypeFile,
						},
					}},
					ChannelAccountID: "channelAccountId",
					ChannelID:        "channelId",
					Recipients: []conversations.PublicRecipientEggParam{{
						DeliveryIdentifiers: []conversations.PublicDeliveryIdentifierParam{{
							Type:  "type",
							Value: "value",
						}},
						ActorID: hubspotsdk.String("actorId"),
						DeliveryIdentifier: conversations.PublicDeliveryIdentifierParam{
							Type:  "type",
							Value: "value",
						},
						Name:           hubspotsdk.String("name"),
						RecipientField: hubspotsdk.String("recipientField"),
					}},
					SenderActorID: "senderActorId",
					Text:          "text",
					Type:          conversations.PublicConversationsMessageEggTypeMessage,
					RichText:      hubspotsdk.String("richText"),
					Subject:       hubspotsdk.String("subject"),
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

func TestMessageListWithOptionalParams(t *testing.T) {
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
	_, err := client.Conversations.Messages.List(
		context.TODO(),
		0,
		conversations.MessageListParams{
			After:    hubspotsdk.String("after"),
			Archived: hubspotsdk.Bool(true),
			Limit:    hubspotsdk.Int(0),
			Property: hubspotsdk.String("property"),
			Sort:     []string{"string"},
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

func TestMessageGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Conversations.Messages.Get(
		context.TODO(),
		"messageId",
		conversations.MessageGetParams{
			ThreadID: 0,
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

func TestMessageGetOriginalContentWithOptionalParams(t *testing.T) {
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
	_, err := client.Conversations.Messages.GetOriginalContent(
		context.TODO(),
		"messageId",
		conversations.MessageGetOriginalContentParams{
			ThreadID: 0,
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
