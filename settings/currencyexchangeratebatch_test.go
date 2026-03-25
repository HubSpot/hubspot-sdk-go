// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package settings_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/stainless-sdks/hubspot-sdk-go"
	"github.com/stainless-sdks/hubspot-sdk-go/internal/testutil"
	"github.com/stainless-sdks/hubspot-sdk-go/option"
	"github.com/stainless-sdks/hubspot-sdk-go/settings"
	"github.com/stainless-sdks/hubspot-sdk-go/shared"
)

func TestCurrencyExchangeRateBatchNew(t *testing.T) {
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
		option.WithAccessToken("pat-na1-xxxxxxxx-xxxx"),
	)
	_, err := client.Settings.Currencies.ExchangeRates.Batch.New(context.TODO(), settings.CurrencyExchangeRateBatchNewParams{
		BatchInputExchangeRateCreateRequest: settings.BatchInputExchangeRateCreateRequestParam{
			Inputs: []settings.ExchangeRateCreateRequestParam{{
				ConversionRate:   0,
				FromCurrencyCode: settings.ExchangeRateCreateRequestFromCurrencyCodeAed,
				EffectiveAt:      hubspotsdk.Time(time.Now()),
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

func TestCurrencyExchangeRateBatchUpdate(t *testing.T) {
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
		option.WithAccessToken("pat-na1-xxxxxxxx-xxxx"),
	)
	_, err := client.Settings.Currencies.ExchangeRates.Batch.Update(context.TODO(), settings.CurrencyExchangeRateBatchUpdateParams{
		BatchInputExchangeRateUpdateRequest: settings.BatchInputExchangeRateUpdateRequestParam{
			Inputs: []settings.ExchangeRateUpdateRequestParam{{
				ID:             "id",
				ConversionRate: 0,
				EffectiveAt:    hubspotsdk.Time(time.Now()),
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

func TestCurrencyExchangeRateBatchGet(t *testing.T) {
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
		option.WithAccessToken("pat-na1-xxxxxxxx-xxxx"),
	)
	_, err := client.Settings.Currencies.ExchangeRates.Batch.Get(context.TODO(), settings.CurrencyExchangeRateBatchGetParams{
		BatchInputPublicObjectID: settings.BatchInputPublicObjectIDParam{
			Inputs: []shared.PublicObjectIDParam{{
				ID: "id",
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
