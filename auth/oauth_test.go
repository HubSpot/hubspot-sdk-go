// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package auth_test

import (
	"bytes"
	"context"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/HubSpot/hubspot-sdk-go"
	"github.com/HubSpot/hubspot-sdk-go/auth"
	"github.com/HubSpot/hubspot-sdk-go/internal/testutil"
	"github.com/HubSpot/hubspot-sdk-go/option"
)

func TestOAuthNewTokenWithOptionalParams(t *testing.T) {
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
	_, err := client.Auth.OAuth.NewToken(context.TODO(), auth.OAuthNewTokenParams{
		ClientID:     hubspotsdk.String("client_id"),
		ClientSecret: hubspotsdk.String("client_secret"),
		Code:         hubspotsdk.String("code"),
		CodeVerifier: hubspotsdk.String("code_verifier"),
		GrantType:    auth.OAuthNewTokenParamsGrantTypeAuthorizationCode,
		RedirectUri:  hubspotsdk.String("redirect_uri"),
		RefreshToken: hubspotsdk.String("refresh_token"),
		Scope:        hubspotsdk.String("scope"),
	})
	if err != nil {
		var apierr *hubspotsdk.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestOAuthIntrospectTokenWithOptionalParams(t *testing.T) {
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
	_, err := client.Auth.OAuth.IntrospectToken(context.TODO(), auth.OAuthIntrospectTokenParams{
		Token:         hubspotsdk.String("token"),
		ClientID:      hubspotsdk.String("client_id"),
		ClientSecret:  hubspotsdk.String("client_secret"),
		TokenTypeHint: hubspotsdk.String("token_type_hint"),
	})
	if err != nil {
		var apierr *hubspotsdk.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestOAuthRevokeTokenWithOptionalParams(t *testing.T) {
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
	resp, err := client.Auth.OAuth.RevokeToken(context.TODO(), auth.OAuthRevokeTokenParams{
		Token:         hubspotsdk.String("token"),
		ClientID:      hubspotsdk.String("client_id"),
		ClientSecret:  hubspotsdk.String("client_secret"),
		TokenTypeHint: hubspotsdk.String("token_type_hint"),
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
