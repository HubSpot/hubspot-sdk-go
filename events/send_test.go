// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package events_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/stainless-sdks/hubspot-sdk-go"
	"github.com/stainless-sdks/hubspot-sdk-go/events"
	"github.com/stainless-sdks/hubspot-sdk-go/internal/testutil"
	"github.com/stainless-sdks/hubspot-sdk-go/option"
)

func TestSendSendWithOptionalParams(t *testing.T) {
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
	err := client.Events.Send.Send(context.TODO(), events.SendSendParams{
		BehavioralEventHTTPCompletionRequest: events.BehavioralEventHTTPCompletionRequestParam{
			EventName:  "pe123456_account_login",
			Email:      hubspotsdk.String("mark.s@lumon.industries"),
			ObjectID:   hubspotsdk.String("089274502"),
			OccurredAt: hubspotsdk.Time(time.Now()),
			Properties: map[string]string{
				"0":  "{",
				"1":  "\"",
				"2":  "h",
				"3":  "s",
				"4":  "_",
				"5":  "p",
				"6":  "a",
				"7":  "g",
				"8":  "e",
				"9":  "_",
				"10": "i",
				"11": "d",
				"12": "\"",
				"13": ":",
				"14": "\"",
				"15": "1",
				"16": "2",
				"17": "3",
				"18": "4",
				"19": "5",
				"20": "6",
				"21": "7",
				"22": "8",
				"23": "9",
				"24": "0",
				"25": "\"",
				"26": ",",
				"27": "\"",
				"28": "h",
				"29": "s",
				"30": "_",
				"31": "e",
				"32": "l",
				"33": "e",
				"34": "m",
				"35": "e",
				"36": "n",
				"37": "t",
				"38": "_",
				"39": "i",
				"40": "d",
				"41": "\"",
				"42": ":",
				"43": "\"",
				"44": "l",
				"45": "o",
				"46": "g",
				"47": "i",
				"48": "n",
				"49": "-",
				"50": "b",
				"51": "u",
				"52": "t",
				"53": "t",
				"54": "o",
				"55": "n",
				"56": "\"",
				"57": ",",
				"58": "\"",
				"59": "h",
				"60": "s",
				"61": "_",
				"62": "p",
				"63": "a",
				"64": "g",
				"65": "e",
				"66": "_",
				"67": "t",
				"68": "i",
				"69": "t",
				"70": "l",
				"71": "e",
				"72": "\"",
				"73": ":",
				"74": "\"",
				"75": "h",
				"76": "o",
				"77": "m",
				"78": "e",
				"79": "p",
				"80": "a",
				"81": "g",
				"82": "e",
				"83": "\"",
				"84": "}",
			},
			Utk:  hubspotsdk.String("utk"),
			Uuid: hubspotsdk.String("uuid"),
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

func TestSendSendBatch(t *testing.T) {
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
	err := client.Events.Send.SendBatch(context.TODO(), events.SendSendBatchParams{
		BatchedBehavioralEventHTTPCompletionRequest: events.BatchedBehavioralEventHTTPCompletionRequestParam{
			Inputs: []events.BehavioralEventHTTPCompletionRequestParam{{
				EventName:  "pe123456_account_login",
				Email:      hubspotsdk.String("mark.s@lumon.industries"),
				ObjectID:   hubspotsdk.String("089274502"),
				OccurredAt: hubspotsdk.Time(time.Now()),
				Properties: map[string]string{
					"0":  "{",
					"1":  "\"",
					"2":  "h",
					"3":  "s",
					"4":  "_",
					"5":  "p",
					"6":  "a",
					"7":  "g",
					"8":  "e",
					"9":  "_",
					"10": "i",
					"11": "d",
					"12": "\"",
					"13": ":",
					"14": "\"",
					"15": "1",
					"16": "2",
					"17": "3",
					"18": "4",
					"19": "5",
					"20": "6",
					"21": "7",
					"22": "8",
					"23": "9",
					"24": "0",
					"25": "\"",
					"26": ",",
					"27": "\"",
					"28": "h",
					"29": "s",
					"30": "_",
					"31": "e",
					"32": "l",
					"33": "e",
					"34": "m",
					"35": "e",
					"36": "n",
					"37": "t",
					"38": "_",
					"39": "i",
					"40": "d",
					"41": "\"",
					"42": ":",
					"43": "\"",
					"44": "l",
					"45": "o",
					"46": "g",
					"47": "i",
					"48": "n",
					"49": "-",
					"50": "b",
					"51": "u",
					"52": "t",
					"53": "t",
					"54": "o",
					"55": "n",
					"56": "\"",
					"57": ",",
					"58": "\"",
					"59": "h",
					"60": "s",
					"61": "_",
					"62": "p",
					"63": "a",
					"64": "g",
					"65": "e",
					"66": "_",
					"67": "t",
					"68": "i",
					"69": "t",
					"70": "l",
					"71": "e",
					"72": "\"",
					"73": ":",
					"74": "\"",
					"75": "h",
					"76": "o",
					"77": "m",
					"78": "e",
					"79": "p",
					"80": "a",
					"81": "g",
					"82": "e",
					"83": "\"",
					"84": "}",
				},
				Utk:  hubspotsdk.String("utk"),
				Uuid: hubspotsdk.String("uuid"),
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
