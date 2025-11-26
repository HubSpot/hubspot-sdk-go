// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package marketing_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/stainless-sdks/hubspot-sdk-go"
	"github.com/stainless-sdks/hubspot-sdk-go/internal/testutil"
	"github.com/stainless-sdks/hubspot-sdk-go/marketing"
	"github.com/stainless-sdks/hubspot-sdk-go/option"
)

func TestSingleSendSendWithOptionalParams(t *testing.T) {
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
	_, err := client.Marketing.SingleSend.Send(context.TODO(), marketing.SingleSendSendParams{
		PublicSingleSendRequestEgg: marketing.PublicSingleSendRequestEggParam{
			EmailID: 0,
			Message: marketing.PublicSingleSendEmailParam{
				To:      "to",
				Bcc:     []string{"string"},
				Cc:      []string{"string"},
				From:    hubspotsdk.String("from"),
				ReplyTo: []string{"string"},
				SendID:  hubspotsdk.String("sendId"),
			},
			ContactProperties: map[string]string{
				"0":  "{",
				"1":  "\"",
				"2":  "l",
				"3":  "a",
				"4":  "s",
				"5":  "t",
				"6":  "n",
				"7":  "a",
				"8":  "m",
				"9":  "e",
				"10": "\"",
				"11": ":",
				"12": "\"",
				"13": "d",
				"14": "o",
				"15": "e",
				"16": "\"",
				"17": ",",
				"18": "\"",
				"19": "f",
				"20": "i",
				"21": "r",
				"22": "s",
				"23": "t",
				"24": "n",
				"25": "a",
				"26": "m",
				"27": "e",
				"28": "\"",
				"29": ":",
				"30": "\"",
				"31": "j",
				"32": "o",
				"33": "h",
				"34": "n",
				"35": "\"",
				"36": "}",
			},
			CustomProperties: map[string]any{
				"0":  map[string]any{},
				"1":  map[string]any{},
				"2":  map[string]any{},
				"3":  map[string]any{},
				"4":  map[string]any{},
				"5":  map[string]any{},
				"6":  map[string]any{},
				"7":  map[string]any{},
				"8":  map[string]any{},
				"9":  map[string]any{},
				"10": map[string]any{},
				"11": map[string]any{},
				"12": map[string]any{},
				"13": map[string]any{},
				"14": map[string]any{},
				"15": map[string]any{},
				"16": map[string]any{},
				"17": map[string]any{},
				"18": map[string]any{},
				"19": map[string]any{},
				"20": map[string]any{},
				"21": map[string]any{},
				"22": map[string]any{},
				"23": map[string]any{},
				"24": map[string]any{},
				"25": map[string]any{},
				"26": map[string]any{},
				"27": map[string]any{},
				"28": map[string]any{},
				"29": map[string]any{},
				"30": map[string]any{},
				"31": map[string]any{},
				"32": map[string]any{},
				"33": map[string]any{},
				"34": map[string]any{},
				"35": map[string]any{},
				"36": map[string]any{},
				"37": map[string]any{},
				"38": map[string]any{},
				"39": map[string]any{},
				"40": map[string]any{},
				"41": map[string]any{},
				"42": map[string]any{},
				"43": map[string]any{},
				"44": map[string]any{},
				"45": map[string]any{},
				"46": map[string]any{},
				"47": map[string]any{},
				"48": map[string]any{},
				"49": map[string]any{},
				"50": map[string]any{},
				"51": map[string]any{},
				"52": map[string]any{},
				"53": map[string]any{},
				"54": map[string]any{},
			},
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
