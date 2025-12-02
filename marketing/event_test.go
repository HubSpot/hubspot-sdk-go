// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package marketing_test

import (
	"bytes"
	"context"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/stainless-sdks/hubspot-sdk-go"
	"github.com/stainless-sdks/hubspot-sdk-go/internal/testutil"
	"github.com/stainless-sdks/hubspot-sdk-go/marketing"
	"github.com/stainless-sdks/hubspot-sdk-go/option"
)

func TestEventNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Marketing.Events.New(context.TODO(), marketing.EventNewParams{
		MarketingEventCreateRequestParams: marketing.MarketingEventCreateRequestParams{
			CustomProperties: []marketing.PropertyValueParam{{
				DataSensitivity:                    marketing.PropertyValueDataSensitivityHigh,
				IsEncrypted:                        true,
				IsLargeValue:                       true,
				Name:                               "name",
				PersistenceTimestamp:               0,
				RequestID:                          "requestId",
				SelectedByUser:                     true,
				SelectedByUserTimestamp:            0,
				Source:                             marketing.PropertyValueSourceAcademy,
				SourceID:                           "sourceId",
				SourceLabel:                        "sourceLabel",
				SourceMetadata:                     "sourceMetadata",
				SourceUpstreamDeployable:           "sourceUpstreamDeployable",
				SourceVid:                          []int64{0},
				Timestamp:                          0,
				Unit:                               "unit",
				UpdatedByUserID:                    0,
				UseTimestampAsPersistenceTimestamp: true,
				Value:                              "value",
			}},
			EventName:         "eventName",
			EventOrganizer:    "eventOrganizer",
			ExternalAccountID: "externalAccountId",
			ExternalEventID:   "externalEventId",
			EndDateTime:       hubspotsdk.Time(time.Now()),
			EventCancelled:    hubspotsdk.Bool(true),
			EventCompleted:    hubspotsdk.Bool(true),
			EventDescription:  hubspotsdk.String("eventDescription"),
			EventType:         hubspotsdk.String("eventType"),
			EventURL:          hubspotsdk.String("eventUrl"),
			StartDateTime:     hubspotsdk.Time(time.Now()),
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

func TestEventUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Marketing.Events.Update(
		context.TODO(),
		"objectId",
		marketing.EventUpdateParams{
			MarketingEventPublicUpdateRequestV2: marketing.MarketingEventPublicUpdateRequestV2Param{
				CustomProperties: []marketing.PropertyValueParam{{
					DataSensitivity:                    marketing.PropertyValueDataSensitivityHigh,
					IsEncrypted:                        true,
					IsLargeValue:                       true,
					Name:                               "name",
					PersistenceTimestamp:               0,
					RequestID:                          "requestId",
					SelectedByUser:                     true,
					SelectedByUserTimestamp:            0,
					Source:                             marketing.PropertyValueSourceAcademy,
					SourceID:                           "sourceId",
					SourceLabel:                        "sourceLabel",
					SourceMetadata:                     "sourceMetadata",
					SourceUpstreamDeployable:           "sourceUpstreamDeployable",
					SourceVid:                          []int64{0},
					Timestamp:                          0,
					Unit:                               "unit",
					UpdatedByUserID:                    0,
					UseTimestampAsPersistenceTimestamp: true,
					Value:                              "value",
				}},
				EndDateTime:      hubspotsdk.Time(time.Now()),
				EventCancelled:   hubspotsdk.Bool(true),
				EventDescription: hubspotsdk.String("eventDescription"),
				EventName:        hubspotsdk.String("eventName"),
				EventOrganizer:   hubspotsdk.String("eventOrganizer"),
				EventType:        hubspotsdk.String("eventType"),
				EventURL:         hubspotsdk.String("eventUrl"),
				StartDateTime:    hubspotsdk.Time(time.Now()),
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

func TestEventListWithOptionalParams(t *testing.T) {
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
	_, err := client.Marketing.Events.List(context.TODO(), marketing.EventListParams{
		After: hubspotsdk.String("after"),
		Limit: hubspotsdk.Int(0),
	})
	if err != nil {
		var apierr *hubspotsdk.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestEventDelete(t *testing.T) {
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
	err := client.Marketing.Events.Delete(context.TODO(), "objectId")
	if err != nil {
		var apierr *hubspotsdk.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestEventCancelByExternalEventID(t *testing.T) {
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
	_, err := client.Marketing.Events.CancelByExternalEventID(
		context.TODO(),
		"externalEventId",
		marketing.EventCancelByExternalEventIDParams{
			ExternalAccountID: "externalAccountId",
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

func TestEventCompleteByExternalEventID(t *testing.T) {
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
	_, err := client.Marketing.Events.CompleteByExternalEventID(
		context.TODO(),
		"externalEventId",
		marketing.EventCompleteByExternalEventIDParams{
			ExternalAccountID: "externalAccountId",
			MarketingEventCompleteRequestParams: marketing.MarketingEventCompleteRequestParams{
				EndDateTime:   time.Now(),
				StartDateTime: time.Now(),
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

func TestEventDeleteBatch(t *testing.T) {
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
	err := client.Marketing.Events.DeleteBatch(context.TODO(), marketing.EventDeleteBatchParams{
		BatchInputMarketingEventPublicObjectIDDeleteRequest: marketing.BatchInputMarketingEventPublicObjectIDDeleteRequestParam{
			Inputs: []marketing.MarketingEventPublicObjectIDDeleteRequestParam{{
				ObjectID: "objectId",
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

func TestEventDeleteBatchByExternalEventID(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte("abc"))
	}))
	defer server.Close()
	baseURL := server.URL
	client := hubspotsdk.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAccessToken("pat-na1-xxxxxxxx-xxxx"),
	)
	resp, err := client.Marketing.Events.DeleteBatchByExternalEventID(context.TODO(), marketing.EventDeleteBatchByExternalEventIDParams{
		BatchInputMarketingEventExternalUniqueIdentifier: marketing.BatchInputMarketingEventExternalUniqueIdentifierParam{
			Inputs: []marketing.MarketingEventExternalUniqueIdentifierParam{{
				AppID:             0,
				ExternalAccountID: "externalAccountId",
				ExternalEventID:   "externalEventId",
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
	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		var apierr *hubspotsdk.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
	if !bytes.Equal(b, []byte("abc")) {
		t.Fatalf("return value not %s: %s", "abc", b)
	}
}

func TestEventDeleteByExternalEventID(t *testing.T) {
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
	err := client.Marketing.Events.DeleteByExternalEventID(
		context.TODO(),
		"externalEventId",
		marketing.EventDeleteByExternalEventIDParams{
			ExternalAccountID: "externalAccountId",
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

func TestEventGet(t *testing.T) {
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
	_, err := client.Marketing.Events.Get(context.TODO(), "objectId")
	if err != nil {
		var apierr *hubspotsdk.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestEventGetByExternalEventID(t *testing.T) {
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
	_, err := client.Marketing.Events.GetByExternalEventID(
		context.TODO(),
		"externalEventId",
		marketing.EventGetByExternalEventIDParams{
			ExternalAccountID: "externalAccountId",
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

func TestEventSearchByExternalEventID(t *testing.T) {
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
	_, err := client.Marketing.Events.SearchByExternalEventID(context.TODO(), marketing.EventSearchByExternalEventIDParams{
		Q: "q",
	})
	if err != nil {
		var apierr *hubspotsdk.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestEventSearchIdentifiersByExternalEventID(t *testing.T) {
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
	_, err := client.Marketing.Events.SearchIdentifiersByExternalEventID(context.TODO(), "externalEventId")
	if err != nil {
		var apierr *hubspotsdk.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestEventUpdateBatch(t *testing.T) {
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
	_, err := client.Marketing.Events.UpdateBatch(context.TODO(), marketing.EventUpdateBatchParams{
		BatchInputMarketingEventPublicUpdateRequestFullV2: marketing.BatchInputMarketingEventPublicUpdateRequestFullV2Param{
			Inputs: []marketing.MarketingEventPublicUpdateRequestFullV2Param{{
				CustomProperties: []marketing.PropertyValueParam{{
					DataSensitivity:                    marketing.PropertyValueDataSensitivityHigh,
					IsEncrypted:                        true,
					IsLargeValue:                       true,
					Name:                               "name",
					PersistenceTimestamp:               0,
					RequestID:                          "requestId",
					SelectedByUser:                     true,
					SelectedByUserTimestamp:            0,
					Source:                             marketing.PropertyValueSourceAcademy,
					SourceID:                           "sourceId",
					SourceLabel:                        "sourceLabel",
					SourceMetadata:                     "sourceMetadata",
					SourceUpstreamDeployable:           "sourceUpstreamDeployable",
					SourceVid:                          []int64{0},
					Timestamp:                          0,
					Unit:                               "unit",
					UpdatedByUserID:                    0,
					UseTimestampAsPersistenceTimestamp: true,
					Value:                              "value",
				}},
				ObjectID:         "objectId",
				EndDateTime:      hubspotsdk.Time(time.Now()),
				EventCancelled:   hubspotsdk.Bool(true),
				EventDescription: hubspotsdk.String("eventDescription"),
				EventName:        hubspotsdk.String("eventName"),
				EventOrganizer:   hubspotsdk.String("eventOrganizer"),
				EventType:        hubspotsdk.String("eventType"),
				EventURL:         hubspotsdk.String("eventUrl"),
				StartDateTime:    hubspotsdk.Time(time.Now()),
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

func TestEventUpdateByExternalEventIDWithOptionalParams(t *testing.T) {
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
	_, err := client.Marketing.Events.UpdateByExternalEventID(
		context.TODO(),
		"externalEventId",
		marketing.EventUpdateByExternalEventIDParams{
			ExternalAccountID: "externalAccountId",
			MarketingEventUpdateRequestParams: marketing.MarketingEventUpdateRequestParams{
				CustomProperties: []marketing.PropertyValueParam{{
					DataSensitivity:                    marketing.PropertyValueDataSensitivityHigh,
					IsEncrypted:                        true,
					IsLargeValue:                       true,
					Name:                               "name",
					PersistenceTimestamp:               0,
					RequestID:                          "requestId",
					SelectedByUser:                     true,
					SelectedByUserTimestamp:            0,
					Source:                             marketing.PropertyValueSourceAcademy,
					SourceID:                           "sourceId",
					SourceLabel:                        "sourceLabel",
					SourceMetadata:                     "sourceMetadata",
					SourceUpstreamDeployable:           "sourceUpstreamDeployable",
					SourceVid:                          []int64{0},
					Timestamp:                          0,
					Unit:                               "unit",
					UpdatedByUserID:                    0,
					UseTimestampAsPersistenceTimestamp: true,
					Value:                              "value",
				}},
				EndDateTime:      hubspotsdk.Time(time.Now()),
				EventCancelled:   hubspotsdk.Bool(true),
				EventCompleted:   hubspotsdk.Bool(true),
				EventDescription: hubspotsdk.String("eventDescription"),
				EventName:        hubspotsdk.String("eventName"),
				EventOrganizer:   hubspotsdk.String("eventOrganizer"),
				EventType:        hubspotsdk.String("eventType"),
				EventURL:         hubspotsdk.String("eventUrl"),
				StartDateTime:    hubspotsdk.Time(time.Now()),
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

func TestEventUpsertBatch(t *testing.T) {
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
	_, err := client.Marketing.Events.UpsertBatch(context.TODO(), marketing.EventUpsertBatchParams{
		BatchInputMarketingEventCreateRequestParams: marketing.BatchInputMarketingEventCreateRequestParams{
			Inputs: []marketing.MarketingEventCreateRequestParams{{
				CustomProperties: []marketing.PropertyValueParam{{
					DataSensitivity:                    marketing.PropertyValueDataSensitivityHigh,
					IsEncrypted:                        true,
					IsLargeValue:                       true,
					Name:                               "name",
					PersistenceTimestamp:               0,
					RequestID:                          "requestId",
					SelectedByUser:                     true,
					SelectedByUserTimestamp:            0,
					Source:                             marketing.PropertyValueSourceAcademy,
					SourceID:                           "sourceId",
					SourceLabel:                        "sourceLabel",
					SourceMetadata:                     "sourceMetadata",
					SourceUpstreamDeployable:           "sourceUpstreamDeployable",
					SourceVid:                          []int64{0},
					Timestamp:                          0,
					Unit:                               "unit",
					UpdatedByUserID:                    0,
					UseTimestampAsPersistenceTimestamp: true,
					Value:                              "value",
				}},
				EventName:         "eventName",
				EventOrganizer:    "eventOrganizer",
				ExternalAccountID: "externalAccountId",
				ExternalEventID:   "externalEventId",
				EndDateTime:       hubspotsdk.Time(time.Now()),
				EventCancelled:    hubspotsdk.Bool(true),
				EventCompleted:    hubspotsdk.Bool(true),
				EventDescription:  hubspotsdk.String("eventDescription"),
				EventType:         hubspotsdk.String("eventType"),
				EventURL:          hubspotsdk.String("eventUrl"),
				StartDateTime:     hubspotsdk.Time(time.Now()),
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

func TestEventUpsertByExternalEventIDWithOptionalParams(t *testing.T) {
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
	_, err := client.Marketing.Events.UpsertByExternalEventID(
		context.TODO(),
		"externalEventId",
		marketing.EventUpsertByExternalEventIDParams{
			MarketingEventCreateRequestParams: marketing.MarketingEventCreateRequestParams{
				CustomProperties: []marketing.PropertyValueParam{{
					DataSensitivity:                    marketing.PropertyValueDataSensitivityHigh,
					IsEncrypted:                        true,
					IsLargeValue:                       true,
					Name:                               "name",
					PersistenceTimestamp:               0,
					RequestID:                          "requestId",
					SelectedByUser:                     true,
					SelectedByUserTimestamp:            0,
					Source:                             marketing.PropertyValueSourceAcademy,
					SourceID:                           "sourceId",
					SourceLabel:                        "sourceLabel",
					SourceMetadata:                     "sourceMetadata",
					SourceUpstreamDeployable:           "sourceUpstreamDeployable",
					SourceVid:                          []int64{0},
					Timestamp:                          0,
					Unit:                               "unit",
					UpdatedByUserID:                    0,
					UseTimestampAsPersistenceTimestamp: true,
					Value:                              "value",
				}},
				EventName:         "eventName",
				EventOrganizer:    "eventOrganizer",
				ExternalAccountID: "externalAccountId",
				ExternalEventID:   "externalEventId",
				EndDateTime:       hubspotsdk.Time(time.Now()),
				EventCancelled:    hubspotsdk.Bool(true),
				EventCompleted:    hubspotsdk.Bool(true),
				EventDescription:  hubspotsdk.String("eventDescription"),
				EventType:         hubspotsdk.String("eventType"),
				EventURL:          hubspotsdk.String("eventUrl"),
				StartDateTime:     hubspotsdk.Time(time.Now()),
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

func TestEventUpsertSubscriberStateByEmail(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte("abc"))
	}))
	defer server.Close()
	baseURL := server.URL
	client := hubspotsdk.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAccessToken("pat-na1-xxxxxxxx-xxxx"),
	)
	resp, err := client.Marketing.Events.UpsertSubscriberStateByEmail(
		context.TODO(),
		"subscriberState",
		marketing.EventUpsertSubscriberStateByEmailParams{
			ExternalEventID:   "externalEventId",
			ExternalAccountID: "externalAccountId",
			BatchInputMarketingEventEmailSubscriber: marketing.BatchInputMarketingEventEmailSubscriberParam{
				Inputs: []marketing.MarketingEventEmailSubscriberParam{{
					ContactProperties: map[string]string{
						"foo": "string",
					},
					Email:               "email",
					InteractionDateTime: 0,
					Properties: map[string]string{
						"foo": "string",
					},
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
	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		var apierr *hubspotsdk.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
	if !bytes.Equal(b, []byte("abc")) {
		t.Fatalf("return value not %s: %s", "abc", b)
	}
}

func TestEventUpsertSubscriberStateByID(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte("abc"))
	}))
	defer server.Close()
	baseURL := server.URL
	client := hubspotsdk.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAccessToken("pat-na1-xxxxxxxx-xxxx"),
	)
	resp, err := client.Marketing.Events.UpsertSubscriberStateByID(
		context.TODO(),
		"subscriberState",
		marketing.EventUpsertSubscriberStateByIDParams{
			ExternalEventID:   "externalEventId",
			ExternalAccountID: "externalAccountId",
			BatchInputMarketingEventSubscriber: marketing.BatchInputMarketingEventSubscriberParam{
				Inputs: []marketing.MarketingEventSubscriberParam{{
					InteractionDateTime: 0,
					Properties: map[string]string{
						"foo": "string",
					},
					Vid: 0,
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
	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		var apierr *hubspotsdk.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
	if !bytes.Equal(b, []byte("abc")) {
		t.Fatalf("return value not %s: %s", "abc", b)
	}
}
