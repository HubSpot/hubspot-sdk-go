// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cms_test

import (
	"bytes"
	"context"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stainless-sdks/hubspot-sdk-go"
	"github.com/stainless-sdks/hubspot-sdk-go/cms"
	"github.com/stainless-sdks/hubspot-sdk-go/internal/testutil"
	"github.com/stainless-sdks/hubspot-sdk-go/option"
)

func TestBlogSettingMultiLanguageAttachToLangGroupWithOptionalParams(t *testing.T) {
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
	resp, err := client.Cms.Blogs.Settings.MultiLanguage.AttachToLangGroup(context.TODO(), cms.BlogSettingMultiLanguageAttachToLangGroupParams{
		AttachToLangPrimaryRequestVNext: cms.AttachToLangPrimaryRequestVNextParam{
			ID:              "id",
			Language:        cms.AttachToLangPrimaryRequestVNextLanguageAa,
			PrimaryID:       "primaryId",
			PrimaryLanguage: cms.AttachToLangPrimaryRequestVNextPrimaryLanguageAa,
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

func TestBlogSettingMultiLanguageNewLanguageVariationWithOptionalParams(t *testing.T) {
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
	_, err := client.Cms.Blogs.Settings.MultiLanguage.NewLanguageVariation(context.TODO(), cms.BlogSettingMultiLanguageNewLanguageVariationParams{
		BlogLanguageCloneRequestVNext: cms.BlogLanguageCloneRequestVNextParam{
			ID:              "id",
			Language:        hubspotsdk.String("language"),
			PrimaryLanguage: hubspotsdk.String("primaryLanguage"),
			Slug:            hubspotsdk.String("slug"),
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

func TestBlogSettingMultiLanguageDetachFromLangGroup(t *testing.T) {
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
	resp, err := client.Cms.Blogs.Settings.MultiLanguage.DetachFromLangGroup(context.TODO(), cms.BlogSettingMultiLanguageDetachFromLangGroupParams{
		DetachFromLangGroupRequestVNext: cms.DetachFromLangGroupRequestVNextParam{
			ID: "id",
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

func TestBlogSettingMultiLanguageSetNewLangPrimary(t *testing.T) {
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
	err := client.Cms.Blogs.Settings.MultiLanguage.SetNewLangPrimary(context.TODO(), cms.BlogSettingMultiLanguageSetNewLangPrimaryParams{
		SetNewLanguagePrimaryRequestVNext: cms.SetNewLanguagePrimaryRequestVNextParam{
			ID: "id",
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

func TestBlogSettingMultiLanguageUpdateLanguages(t *testing.T) {
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
	resp, err := client.Cms.Blogs.Settings.MultiLanguage.UpdateLanguages(context.TODO(), cms.BlogSettingMultiLanguageUpdateLanguagesParams{
		UpdateLanguagesRequestVNext: cms.UpdateLanguagesRequestVNextParam{
			Languages: map[string]string{
				"foo": "aa",
			},
			PrimaryID: "primaryId",
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
