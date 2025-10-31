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
	"time"

	"github.com/stainless-sdks/hubspot-sdk-go"
	"github.com/stainless-sdks/hubspot-sdk-go/cms"
	"github.com/stainless-sdks/hubspot-sdk-go/internal/testutil"
	"github.com/stainless-sdks/hubspot-sdk-go/option"
	"github.com/stainless-sdks/hubspot-sdk-go/shared"
)

func TestHubdbTableNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Cms.Hubdb.Tables.New(context.TODO(), cms.HubdbTableNewParams{
		HubDBTableV3Request: cms.HubDBTableV3RequestParam{
			Label:                "label",
			Name:                 "name",
			AllowChildTables:     hubspotsdk.Bool(true),
			AllowPublicAPIAccess: hubspotsdk.Bool(true),
			Columns: []cms.ColumnRequestParam{{
				ID:    0,
				Label: "label",
				Name:  "name",
				Options: []shared.OptionParam{{
					Description:  "",
					DisplayOrder: 0,
					DoubleData:   0,
					Hidden:       false,
					Label:        "",
					ReadOnly:     false,
					Value:        "",
				}},
				Type:                  cms.ColumnRequestTypeNull,
				ForeignColumnID:       hubspotsdk.Int(0),
				ForeignTableID:        hubspotsdk.Int(0),
				MaxNumberOfCharacters: hubspotsdk.Int(0),
				MaxNumberOfOptions:    hubspotsdk.Int(0),
			}},
			DynamicMetaTags: map[string]int64{
				"foo": 0,
			},
			EnableChildTablePages: hubspotsdk.Bool(true),
			UseForPages:           hubspotsdk.Bool(true),
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

func TestHubdbTableListWithOptionalParams(t *testing.T) {
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
	_, err := client.Cms.Hubdb.Tables.List(context.TODO(), cms.HubdbTableListParams{
		After:                hubspotsdk.String("after"),
		Archived:             hubspotsdk.Bool(true),
		ContentType:          hubspotsdk.String("contentType"),
		CreatedAfter:         hubspotsdk.Time(time.Now()),
		CreatedAt:            hubspotsdk.Time(time.Now()),
		CreatedBefore:        hubspotsdk.Time(time.Now()),
		IsGetLocalizedSchema: hubspotsdk.Bool(true),
		Limit:                hubspotsdk.Int(0),
		Sort:                 []string{"string"},
		UpdatedAfter:         hubspotsdk.Time(time.Now()),
		UpdatedAt:            hubspotsdk.Time(time.Now()),
		UpdatedBefore:        hubspotsdk.Time(time.Now()),
	})
	if err != nil {
		var apierr *hubspotsdk.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestHubdbTableDelete(t *testing.T) {
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
	err := client.Cms.Hubdb.Tables.Delete(context.TODO(), "tableIdOrName")
	if err != nil {
		var apierr *hubspotsdk.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestHubdbTableCloneDraftWithOptionalParams(t *testing.T) {
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
	_, err := client.Cms.Hubdb.Tables.CloneDraft(
		context.TODO(),
		"tableIdOrName",
		cms.HubdbTableCloneDraftParams{
			HubDBTableCloneRequest: cms.HubDBTableCloneRequestParam{
				CopyRows:         true,
				IsHubspotDefined: true,
				NewLabel:         hubspotsdk.String("newLabel"),
				NewName:          hubspotsdk.String("newName"),
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

func TestHubdbTableDeleteVersion(t *testing.T) {
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
	err := client.Cms.Hubdb.Tables.DeleteVersion(
		context.TODO(),
		0,
		cms.HubdbTableDeleteVersionParams{
			TableIDOrName: "tableIdOrName",
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

func TestHubdbTableExportWithOptionalParams(t *testing.T) {
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
	resp, err := client.Cms.Hubdb.Tables.Export(
		context.TODO(),
		"tableIdOrName",
		cms.HubdbTableExportParams{
			Format: hubspotsdk.String("format"),
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

func TestHubdbTableExportDraftWithOptionalParams(t *testing.T) {
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
	resp, err := client.Cms.Hubdb.Tables.ExportDraft(
		context.TODO(),
		"tableIdOrName",
		cms.HubdbTableExportDraftParams{
			Format: hubspotsdk.String("format"),
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

func TestHubdbTableGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Cms.Hubdb.Tables.Get(
		context.TODO(),
		"tableIdOrName",
		cms.HubdbTableGetParams{
			Archived:             hubspotsdk.Bool(true),
			IncludeForeignIDs:    hubspotsdk.Bool(true),
			IsGetLocalizedSchema: hubspotsdk.Bool(true),
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

func TestHubdbTableGetDraftWithOptionalParams(t *testing.T) {
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
	_, err := client.Cms.Hubdb.Tables.GetDraft(
		context.TODO(),
		"tableIdOrName",
		cms.HubdbTableGetDraftParams{
			Archived:             hubspotsdk.Bool(true),
			IncludeForeignIDs:    hubspotsdk.Bool(true),
			IsGetLocalizedSchema: hubspotsdk.Bool(true),
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

func TestHubdbTableImportDraftWithOptionalParams(t *testing.T) {
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
	_, err := client.Cms.Hubdb.Tables.ImportDraft(
		context.TODO(),
		"tableIdOrName",
		cms.HubdbTableImportDraftParams{
			Config: hubspotsdk.String("config"),
			File:   io.Reader(bytes.NewBuffer([]byte("some file contents"))),
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

func TestHubdbTableListDraftWithOptionalParams(t *testing.T) {
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
	_, err := client.Cms.Hubdb.Tables.ListDraft(context.TODO(), cms.HubdbTableListDraftParams{
		After:                hubspotsdk.String("after"),
		Archived:             hubspotsdk.Bool(true),
		ContentType:          hubspotsdk.String("contentType"),
		CreatedAfter:         hubspotsdk.Time(time.Now()),
		CreatedAt:            hubspotsdk.Time(time.Now()),
		CreatedBefore:        hubspotsdk.Time(time.Now()),
		IsGetLocalizedSchema: hubspotsdk.Bool(true),
		Limit:                hubspotsdk.Int(0),
		Sort:                 []string{"string"},
		UpdatedAfter:         hubspotsdk.Time(time.Now()),
		UpdatedAt:            hubspotsdk.Time(time.Now()),
		UpdatedBefore:        hubspotsdk.Time(time.Now()),
	})
	if err != nil {
		var apierr *hubspotsdk.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestHubdbTablePublishDraftWithOptionalParams(t *testing.T) {
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
	_, err := client.Cms.Hubdb.Tables.PublishDraft(
		context.TODO(),
		"tableIdOrName",
		cms.HubdbTablePublishDraftParams{
			IncludeForeignIDs: hubspotsdk.Bool(true),
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

func TestHubdbTableResetDraftWithOptionalParams(t *testing.T) {
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
	_, err := client.Cms.Hubdb.Tables.ResetDraft(
		context.TODO(),
		"tableIdOrName",
		cms.HubdbTableResetDraftParams{
			IncludeForeignIDs: hubspotsdk.Bool(true),
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

func TestHubdbTableUnpublishWithOptionalParams(t *testing.T) {
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
	_, err := client.Cms.Hubdb.Tables.Unpublish(
		context.TODO(),
		"tableIdOrName",
		cms.HubdbTableUnpublishParams{
			IncludeForeignIDs: hubspotsdk.Bool(true),
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

func TestHubdbTableUpdateDraftWithOptionalParams(t *testing.T) {
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
	_, err := client.Cms.Hubdb.Tables.UpdateDraft(
		context.TODO(),
		"tableIdOrName",
		cms.HubdbTableUpdateDraftParams{
			HubDBTableV3Request: cms.HubDBTableV3RequestParam{
				Label:                "label",
				Name:                 "name",
				AllowChildTables:     hubspotsdk.Bool(true),
				AllowPublicAPIAccess: hubspotsdk.Bool(true),
				Columns: []cms.ColumnRequestParam{{
					ID:    0,
					Label: "label",
					Name:  "name",
					Options: []shared.OptionParam{{
						Description:  "",
						DisplayOrder: 0,
						DoubleData:   0,
						Hidden:       false,
						Label:        "",
						ReadOnly:     false,
						Value:        "",
					}},
					Type:                  cms.ColumnRequestTypeNull,
					ForeignColumnID:       hubspotsdk.Int(0),
					ForeignTableID:        hubspotsdk.Int(0),
					MaxNumberOfCharacters: hubspotsdk.Int(0),
					MaxNumberOfOptions:    hubspotsdk.Int(0),
				}},
				DynamicMetaTags: map[string]int64{
					"foo": 0,
				},
				EnableChildTablePages: hubspotsdk.Bool(true),
				UseForPages:           hubspotsdk.Bool(true),
			},
			Archived:             hubspotsdk.Bool(true),
			IncludeForeignIDs:    hubspotsdk.Bool(true),
			IsGetLocalizedSchema: hubspotsdk.Bool(true),
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
