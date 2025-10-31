// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cms_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/stainless-sdks/hubspot-sdk-go"
	"github.com/stainless-sdks/hubspot-sdk-go/cms"
	"github.com/stainless-sdks/hubspot-sdk-go/internal/testutil"
	"github.com/stainless-sdks/hubspot-sdk-go/option"
	"github.com/stainless-sdks/hubspot-sdk-go/shared"
)

func TestBlogAuthorNew(t *testing.T) {
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
	_, err := client.Cms.Blogs.Authors.New(context.TODO(), cms.BlogAuthorNewParams{
		BlogAuthor: cms.BlogAuthorParam{
			ID:               "id",
			Avatar:           "avatar",
			Bio:              "bio",
			Created:          time.Now(),
			DeletedAt:        time.Now(),
			DisplayName:      "displayName",
			Email:            "email",
			Facebook:         "facebook",
			FullName:         "fullName",
			Language:         cms.BlogAuthorLanguageAf,
			Linkedin:         "linkedin",
			Name:             "name",
			Slug:             "slug",
			TranslatedFromID: 0,
			Twitter:          "twitter",
			Updated:          time.Now(),
			Website:          "website",
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

func TestBlogAuthorUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Cms.Blogs.Authors.Update(
		context.TODO(),
		"objectId",
		cms.BlogAuthorUpdateParams{
			BlogAuthor: cms.BlogAuthorParam{
				ID:               "id",
				Avatar:           "avatar",
				Bio:              "bio",
				Created:          time.Now(),
				DeletedAt:        time.Now(),
				DisplayName:      "displayName",
				Email:            "email",
				Facebook:         "facebook",
				FullName:         "fullName",
				Language:         cms.BlogAuthorLanguageAf,
				Linkedin:         "linkedin",
				Name:             "name",
				Slug:             "slug",
				TranslatedFromID: 0,
				Twitter:          "twitter",
				Updated:          time.Now(),
				Website:          "website",
			},
			Archived: hubspotsdk.Bool(true),
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

func TestBlogAuthorListWithOptionalParams(t *testing.T) {
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
	_, err := client.Cms.Blogs.Authors.List(context.TODO(), cms.BlogAuthorListParams{
		After:         hubspotsdk.String("after"),
		Archived:      hubspotsdk.Bool(true),
		CreatedAfter:  hubspotsdk.Time(time.Now()),
		CreatedAt:     hubspotsdk.Time(time.Now()),
		CreatedBefore: hubspotsdk.Time(time.Now()),
		Limit:         hubspotsdk.Int(0),
		Property:      hubspotsdk.String("property"),
		Sort:          []string{"string"},
		UpdatedAfter:  hubspotsdk.Time(time.Now()),
		UpdatedAt:     hubspotsdk.Time(time.Now()),
		UpdatedBefore: hubspotsdk.Time(time.Now()),
	})
	if err != nil {
		var apierr *hubspotsdk.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestBlogAuthorDeleteWithOptionalParams(t *testing.T) {
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
	err := client.Cms.Blogs.Authors.Delete(
		context.TODO(),
		"objectId",
		cms.BlogAuthorDeleteParams{
			Archived: hubspotsdk.Bool(true),
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

func TestBlogAuthorAttachToLangGroupWithOptionalParams(t *testing.T) {
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
	err := client.Cms.Blogs.Authors.AttachToLangGroup(context.TODO(), cms.BlogAuthorAttachToLangGroupParams{
		AttachToLangPrimaryRequestVNext: cms.AttachToLangPrimaryRequestVNextParam{
			ID:              "id",
			Language:        "language",
			PrimaryID:       "primaryId",
			PrimaryLanguage: hubspotsdk.String("primaryLanguage"),
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

func TestBlogAuthorNewBatch(t *testing.T) {
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
	_, err := client.Cms.Blogs.Authors.NewBatch(context.TODO(), cms.BlogAuthorNewBatchParams{
		BatchInputBlogAuthor: cms.BatchInputBlogAuthorParam{
			Inputs: []cms.BlogAuthorParam{{
				ID:               "id",
				Avatar:           "avatar",
				Bio:              "bio",
				Created:          time.Now(),
				DeletedAt:        time.Now(),
				DisplayName:      "displayName",
				Email:            "email",
				Facebook:         "facebook",
				FullName:         "fullName",
				Language:         cms.BlogAuthorLanguageAf,
				Linkedin:         "linkedin",
				Name:             "name",
				Slug:             "slug",
				TranslatedFromID: 0,
				Twitter:          "twitter",
				Updated:          time.Now(),
				Website:          "website",
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

func TestBlogAuthorNewLanguageVariationWithOptionalParams(t *testing.T) {
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
	_, err := client.Cms.Blogs.Authors.NewLanguageVariation(context.TODO(), cms.BlogAuthorNewLanguageVariationParams{
		BlogAuthorCloneRequestVNext: cms.BlogAuthorCloneRequestVNextParam{
			ID: "id",
			BlogAuthor: cms.BlogAuthorParam{
				ID:               "id",
				Avatar:           "avatar",
				Bio:              "bio",
				Created:          time.Now(),
				DeletedAt:        time.Now(),
				DisplayName:      "displayName",
				Email:            "email",
				Facebook:         "facebook",
				FullName:         "fullName",
				Language:         cms.BlogAuthorLanguageAf,
				Linkedin:         "linkedin",
				Name:             "name",
				Slug:             "slug",
				TranslatedFromID: 0,
				Twitter:          "twitter",
				Updated:          time.Now(),
				Website:          "website",
			},
			Language:        hubspotsdk.String("language"),
			PrimaryLanguage: hubspotsdk.String("primaryLanguage"),
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

func TestBlogAuthorDeleteBatch(t *testing.T) {
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
	err := client.Cms.Blogs.Authors.DeleteBatch(context.TODO(), cms.BlogAuthorDeleteBatchParams{
		BatchInputString: shared.BatchInputStringParam{
			Inputs: []string{"string"},
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

func TestBlogAuthorDetachFromLangGroup(t *testing.T) {
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
	err := client.Cms.Blogs.Authors.DetachFromLangGroup(context.TODO(), cms.BlogAuthorDetachFromLangGroupParams{
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
}

func TestBlogAuthorGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Cms.Blogs.Authors.Get(
		context.TODO(),
		"objectId",
		cms.BlogAuthorGetParams{
			Archived: hubspotsdk.Bool(true),
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

func TestBlogAuthorGetBatchWithOptionalParams(t *testing.T) {
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
	_, err := client.Cms.Blogs.Authors.GetBatch(context.TODO(), cms.BlogAuthorGetBatchParams{
		BatchInputString: shared.BatchInputStringParam{
			Inputs: []string{"string"},
		},
		Archived: hubspotsdk.Bool(true),
	})
	if err != nil {
		var apierr *hubspotsdk.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestBlogAuthorSetNewLangPrimary(t *testing.T) {
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
	err := client.Cms.Blogs.Authors.SetNewLangPrimary(context.TODO(), cms.BlogAuthorSetNewLangPrimaryParams{
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

func TestBlogAuthorUpdateBatchWithOptionalParams(t *testing.T) {
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
	_, err := client.Cms.Blogs.Authors.UpdateBatch(context.TODO(), cms.BlogAuthorUpdateBatchParams{
		BatchInputJsonNode: cms.BatchInputJsonNodeParam{
			Inputs: []any{map[string]interface{}{}},
		},
		Archived: hubspotsdk.Bool(true),
	})
	if err != nil {
		var apierr *hubspotsdk.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestBlogAuthorUpdateLanguages(t *testing.T) {
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
	err := client.Cms.Blogs.Authors.UpdateLanguages(context.TODO(), cms.BlogAuthorUpdateLanguagesParams{
		UpdateLanguagesRequestVNext: cms.UpdateLanguagesRequestVNextParam{
			Languages: map[string]string{
				"foo": "string",
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
}
