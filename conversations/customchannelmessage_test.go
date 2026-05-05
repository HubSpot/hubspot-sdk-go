// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package conversations_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/HubSpot/hubspot-sdk-go"
	"github.com/HubSpot/hubspot-sdk-go/conversations"
	"github.com/HubSpot/hubspot-sdk-go/internal/testutil"
	"github.com/HubSpot/hubspot-sdk-go/option"
)

func TestCustomChannelMessageNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Conversations.CustomChannels.Messages.New(
		context.TODO(),
		0,
		conversations.CustomChannelMessageNewParams{
			ChannelIntegrationMessageEgg: conversations.ChannelIntegrationMessageEggParam{
				Attachments: []conversations.ChannelIntegrationMessageEggAttachmentUnionParam{{
					OfFile: &conversations.FileAttachmentParam{
						FileID:        "fileId",
						Type:          conversations.FileAttachmentTypeFile,
						FileUsageType: conversations.FileAttachmentFileUsageTypeAudio,
					},
				}},
				ChannelAccountID: "channelAccountId",
				MessageDirection: conversations.ChannelIntegrationMessageEggMessageDirectionIncoming,
				Recipients: []conversations.ChannelIntegrationParticipantParam{{
					DeliveryIdentifier: conversations.PublicDeliveryIdentifierParam{
						Type:  conversations.PublicDeliveryIdentifierTypeChannelSpecificOpaqueID,
						Value: "value",
					},
					Name:          hubspotsdk.String("name"),
					SenderActorID: hubspotsdk.String("senderActorId"),
				}},
				Senders: []conversations.ChannelIntegrationParticipantParam{{
					DeliveryIdentifier: conversations.PublicDeliveryIdentifierParam{
						Type:  conversations.PublicDeliveryIdentifierTypeChannelSpecificOpaqueID,
						Value: "value",
					},
					Name:          hubspotsdk.String("name"),
					SenderActorID: hubspotsdk.String("senderActorId"),
				}},
				Text:                     "text",
				Timestamp:                time.Now(),
				AssociateWithContactID:   hubspotsdk.Int(0),
				InReplyToID:              hubspotsdk.String("inReplyToId"),
				IntegrationIdempotencyID: hubspotsdk.String("integrationIdempotencyId"),
				IntegrationThreadID:      hubspotsdk.String("integrationThreadId"),
				PreResolvedContacts: conversations.PreResolvedContactsParam{
					Contacts: []conversations.PreResolvedContactParam{{
						ContactPropertiesLeadingToMatch: []string{"address"},
						ContactVid:                      0,
					}},
				},
				RichText: hubspotsdk.String("richText"),
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

func TestCustomChannelMessageUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Conversations.CustomChannels.Messages.Update(
		context.TODO(),
		"messageId",
		conversations.CustomChannelMessageUpdateParams{
			ChannelID: 0,
			PublicChannelIntegrationMessageUpdateRequest: conversations.PublicChannelIntegrationMessageUpdateRequestParam{
				StatusType:   conversations.PublicChannelIntegrationMessageUpdateRequestStatusTypeFailed,
				ErrorMessage: hubspotsdk.String("errorMessage"),
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

func TestCustomChannelMessageGet(t *testing.T) {
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
	_, err := client.Conversations.CustomChannels.Messages.Get(
		context.TODO(),
		"messageId",
		conversations.CustomChannelMessageGetParams{
			ChannelID: 0,
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
