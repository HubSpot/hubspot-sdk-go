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
	"github.com/stainless-sdks/hubspot-sdk-go/shared"
)

func TestMarketingEventNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Marketing.MarketingEvents.New(context.TODO(), marketing.MarketingEventNewParams{
		MarketingEventCreateRequestParams: marketing.MarketingEventCreateRequestParams{
			CustomProperties: []shared.PropertyValueParam{{
				DataSensitivity:                    shared.PropertyValueDataSensitivityHigh,
				IsEncrypted:                        true,
				IsLargeValue:                       true,
				Name:                               "name",
				PersistenceTimestamp:               0,
				RequestID:                          "requestId",
				SelectedByUser:                     true,
				SelectedByUserTimestamp:            0,
				Source:                             shared.PropertyValueSourceAcademy,
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

func TestMarketingEventUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Marketing.MarketingEvents.Update(
		context.TODO(),
		"objectId",
		marketing.MarketingEventUpdateParams{
			MarketingEventPublicUpdateRequestV2: marketing.MarketingEventPublicUpdateRequestV2Param{
				CustomProperties: []shared.PropertyValueParam{{
					DataSensitivity:                    shared.PropertyValueDataSensitivityHigh,
					IsEncrypted:                        true,
					IsLargeValue:                       true,
					Name:                               "name",
					PersistenceTimestamp:               0,
					RequestID:                          "requestId",
					SelectedByUser:                     true,
					SelectedByUserTimestamp:            0,
					Source:                             shared.PropertyValueSourceAcademy,
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

func TestMarketingEventListWithOptionalParams(t *testing.T) {
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
	_, err := client.Marketing.MarketingEvents.List(context.TODO(), marketing.MarketingEventListParams{
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

func TestMarketingEventDelete(t *testing.T) {
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
	err := client.Marketing.MarketingEvents.Delete(context.TODO(), "objectId")
	if err != nil {
		var apierr *hubspotsdk.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestMarketingEventDeleteBatch(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte("abc"))
	}))
	defer server.Close()
	baseURL := server.URL
	client := hubspotsdk.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAccessToken("My Access Token"),
	)
	resp, err := client.Marketing.MarketingEvents.DeleteBatch(context.TODO(), marketing.MarketingEventDeleteBatchParams{
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

func TestMarketingEventDeleteBatchByExternalEventID(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte("abc"))
	}))
	defer server.Close()
	baseURL := server.URL
	client := hubspotsdk.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAccessToken("My Access Token"),
	)
	resp, err := client.Marketing.MarketingEvents.DeleteBatchByExternalEventID(context.TODO(), marketing.MarketingEventDeleteBatchByExternalEventIDParams{
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

func TestMarketingEventDeleteByExternalEventID(t *testing.T) {
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
	err := client.Marketing.MarketingEvents.DeleteByExternalEventID(
		context.TODO(),
		"externalEventId",
		marketing.MarketingEventDeleteByExternalEventIDParams{
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

func TestMarketingEventGet(t *testing.T) {
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
	_, err := client.Marketing.MarketingEvents.Get(context.TODO(), "objectId")
	if err != nil {
		var apierr *hubspotsdk.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestMarketingEventGetByExternalEventID(t *testing.T) {
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
	_, err := client.Marketing.MarketingEvents.GetByExternalEventID(
		context.TODO(),
		"externalEventId",
		marketing.MarketingEventGetByExternalEventIDParams{
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

func TestMarketingEventSearchByExternalEventID(t *testing.T) {
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
	_, err := client.Marketing.MarketingEvents.SearchByExternalEventID(context.TODO(), marketing.MarketingEventSearchByExternalEventIDParams{
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

func TestMarketingEventSearchIdentifiersByExternalEventID(t *testing.T) {
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
	_, err := client.Marketing.MarketingEvents.SearchIdentifiersByExternalEventID(context.TODO(), "externalEventId")
	if err != nil {
		var apierr *hubspotsdk.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestMarketingEventUpdateBatch(t *testing.T) {
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
	_, err := client.Marketing.MarketingEvents.UpdateBatch(context.TODO(), marketing.MarketingEventUpdateBatchParams{
		BatchInputMarketingEventPublicUpdateRequestFullV2: marketing.BatchInputMarketingEventPublicUpdateRequestFullV2Param{
			Inputs: []marketing.MarketingEventPublicUpdateRequestFullV2Param{{
				CustomProperties: []shared.PropertyValueParam{{
					DataSensitivity:                    shared.PropertyValueDataSensitivityHigh,
					IsEncrypted:                        true,
					IsLargeValue:                       true,
					Name:                               "name",
					PersistenceTimestamp:               0,
					RequestID:                          "requestId",
					SelectedByUser:                     true,
					SelectedByUserTimestamp:            0,
					Source:                             shared.PropertyValueSourceAcademy,
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

func TestMarketingEventUpdateByExternalEventIDWithOptionalParams(t *testing.T) {
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
	_, err := client.Marketing.MarketingEvents.UpdateByExternalEventID(
		context.TODO(),
		"externalEventId",
		marketing.MarketingEventUpdateByExternalEventIDParams{
			ExternalAccountID: "externalAccountId",
			MarketingEventUpdateRequestParams: marketing.MarketingEventUpdateRequestParams{
				CustomProperties: []shared.PropertyValueParam{{
					DataSensitivity:                    shared.PropertyValueDataSensitivityHigh,
					IsEncrypted:                        true,
					IsLargeValue:                       true,
					Name:                               "name",
					PersistenceTimestamp:               0,
					RequestID:                          "requestId",
					SelectedByUser:                     true,
					SelectedByUserTimestamp:            0,
					Source:                             shared.PropertyValueSourceAcademy,
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

func TestMarketingEventUpsertBatch(t *testing.T) {
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
	_, err := client.Marketing.MarketingEvents.UpsertBatch(context.TODO(), marketing.MarketingEventUpsertBatchParams{
		BatchInputMarketingEventCreateRequestParams: marketing.BatchInputMarketingEventCreateRequestParams{
			Inputs: []marketing.MarketingEventCreateRequestParams{{
				CustomProperties: []shared.PropertyValueParam{{
					DataSensitivity:                    shared.PropertyValueDataSensitivityHigh,
					IsEncrypted:                        true,
					IsLargeValue:                       true,
					Name:                               "name",
					PersistenceTimestamp:               0,
					RequestID:                          "requestId",
					SelectedByUser:                     true,
					SelectedByUserTimestamp:            0,
					Source:                             shared.PropertyValueSourceAcademy,
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

func TestMarketingEventUpsertByExternalEventIDWithOptionalParams(t *testing.T) {
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
	_, err := client.Marketing.MarketingEvents.UpsertByExternalEventID(
		context.TODO(),
		"externalEventId",
		marketing.MarketingEventUpsertByExternalEventIDParams{
			MarketingEventCreateRequestParams: marketing.MarketingEventCreateRequestParams{
				CustomProperties: []shared.PropertyValueParam{{
					DataSensitivity:                    shared.PropertyValueDataSensitivityHigh,
					IsEncrypted:                        true,
					IsLargeValue:                       true,
					Name:                               "name",
					PersistenceTimestamp:               0,
					RequestID:                          "requestId",
					SelectedByUser:                     true,
					SelectedByUserTimestamp:            0,
					Source:                             shared.PropertyValueSourceAcademy,
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
