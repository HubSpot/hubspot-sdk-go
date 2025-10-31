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

func TestTransactionalSingleEmailSendWithOptionalParams(t *testing.T) {
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
		option.WithAccessToken("pat-na1-xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"),
	)
	_, err := client.Marketing.Transactional.SingleEmail.Send(context.TODO(), marketing.TransactionalSingleEmailSendParams{
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
				"0":  map[string]interface{}{},
				"1":  map[string]interface{}{},
				"2":  map[string]interface{}{},
				"3":  map[string]interface{}{},
				"4":  map[string]interface{}{},
				"5":  map[string]interface{}{},
				"6":  map[string]interface{}{},
				"7":  map[string]interface{}{},
				"8":  map[string]interface{}{},
				"9":  map[string]interface{}{},
				"10": map[string]interface{}{},
				"11": map[string]interface{}{},
				"12": map[string]interface{}{},
				"13": map[string]interface{}{},
				"14": map[string]interface{}{},
				"15": map[string]interface{}{},
				"16": map[string]interface{}{},
				"17": map[string]interface{}{},
				"18": map[string]interface{}{},
				"19": map[string]interface{}{},
				"20": map[string]interface{}{},
				"21": map[string]interface{}{},
				"22": map[string]interface{}{},
				"23": map[string]interface{}{},
				"24": map[string]interface{}{},
				"25": map[string]interface{}{},
				"26": map[string]interface{}{},
				"27": map[string]interface{}{},
				"28": map[string]interface{}{},
				"29": map[string]interface{}{},
				"30": map[string]interface{}{},
				"31": map[string]interface{}{},
				"32": map[string]interface{}{},
				"33": map[string]interface{}{},
				"34": map[string]interface{}{},
				"35": map[string]interface{}{},
				"36": map[string]interface{}{},
				"37": map[string]interface{}{},
				"38": map[string]interface{}{},
				"39": map[string]interface{}{},
				"40": map[string]interface{}{},
				"41": map[string]interface{}{},
				"42": map[string]interface{}{},
				"43": map[string]interface{}{},
				"44": map[string]interface{}{},
				"45": map[string]interface{}{},
				"46": map[string]interface{}{},
				"47": map[string]interface{}{},
				"48": map[string]interface{}{},
				"49": map[string]interface{}{},
				"50": map[string]interface{}{},
				"51": map[string]interface{}{},
				"52": map[string]interface{}{},
				"53": map[string]interface{}{},
				"54": map[string]interface{}{},
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
