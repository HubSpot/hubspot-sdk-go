// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cms_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/stainless-sdks/hubspot-sdk-go"
	"github.com/stainless-sdks/hubspot-sdk-go/cms"
	"github.com/stainless-sdks/hubspot-sdk-go/internal/testutil"
	"github.com/stainless-sdks/hubspot-sdk-go/option"
	"github.com/stainless-sdks/hubspot-sdk-go/shared"
)

func TestHubdbRowNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Cms.Hubdb.Rows.New(
		context.TODO(),
		"tableIdOrName",
		cms.HubdbRowNewParams{
			HubDBTableRowV3Request: cms.HubDBTableRowV3RequestParam{
				ChildTableID: 0,
				DisplayIndex: 0,
				Values: map[string]cms.Variant{
					"foo": {
						"foo": "bar",
					},
				},
				Name: hubspotsdk.String("name"),
				Path: hubspotsdk.String("path"),
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

func TestHubdbRowListWithOptionalParams(t *testing.T) {
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
	_, err := client.Cms.Hubdb.Rows.List(
		context.TODO(),
		"tableIdOrName",
		cms.HubdbRowListParams{
			After:      hubspotsdk.String("after"),
			Archived:   hubspotsdk.Bool(true),
			Limit:      hubspotsdk.Int(0),
			Offset:     hubspotsdk.Int(0),
			Properties: []string{"string"},
			Sort:       []string{"string"},
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

func TestHubdbRowCloneBatch(t *testing.T) {
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
	_, err := client.Cms.Hubdb.Rows.CloneBatch(
		context.TODO(),
		"tableIdOrName",
		cms.HubdbRowCloneBatchParams{
			BatchInputHubDBTableRowBatchCloneRequest: cms.BatchInputHubDBTableRowBatchCloneRequestParam{
				Inputs: []cms.HubDBTableRowBatchCloneRequestParam{{
					ID:   "id",
					Name: hubspotsdk.String("name"),
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
}

func TestHubdbRowCloneDraftWithOptionalParams(t *testing.T) {
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
	_, err := client.Cms.Hubdb.Rows.CloneDraft(
		context.TODO(),
		"321669910225",
		cms.HubdbRowCloneDraftParams{
			TableIDOrName: "tableIdOrName",
			Name:          hubspotsdk.String("name"),
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

func TestHubdbRowNewBatch(t *testing.T) {
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
	_, err := client.Cms.Hubdb.Rows.NewBatch(
		context.TODO(),
		"tableIdOrName",
		cms.HubdbRowNewBatchParams{
			BatchInputHubDBTableRowV3Request: cms.BatchInputHubDBTableRowV3RequestParam{
				Inputs: []cms.HubDBTableRowV3RequestParam{{
					ChildTableID: 0,
					DisplayIndex: 0,
					Values: map[string]cms.Variant{
						"foo": {
							"foo": "bar",
						},
					},
					Name: hubspotsdk.String("name"),
					Path: hubspotsdk.String("path"),
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
}

func TestHubdbRowDeleteDraft(t *testing.T) {
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
	err := client.Cms.Hubdb.Rows.DeleteDraft(
		context.TODO(),
		"321669910225",
		cms.HubdbRowDeleteDraftParams{
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

func TestHubdbRowGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Cms.Hubdb.Rows.Get(
		context.TODO(),
		"321669910225",
		cms.HubdbRowGetParams{
			TableIDOrName: "tableIdOrName",
			Archived:      hubspotsdk.Bool(true),
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

func TestHubdbRowGetBatch(t *testing.T) {
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
	_, err := client.Cms.Hubdb.Rows.GetBatch(
		context.TODO(),
		"tableIdOrName",
		cms.HubdbRowGetBatchParams{
			BatchInputString: shared.BatchInputStringParam{
				Inputs: []string{"string"},
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

func TestHubdbRowGetDraftWithOptionalParams(t *testing.T) {
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
	_, err := client.Cms.Hubdb.Rows.GetDraft(
		context.TODO(),
		"321669910225",
		cms.HubdbRowGetDraftParams{
			TableIDOrName: "tableIdOrName",
			Archived:      hubspotsdk.Bool(true),
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

func TestHubdbRowGetDraftBatch(t *testing.T) {
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
	_, err := client.Cms.Hubdb.Rows.GetDraftBatch(
		context.TODO(),
		"tableIdOrName",
		cms.HubdbRowGetDraftBatchParams{
			BatchInputString: shared.BatchInputStringParam{
				Inputs: []string{"string"},
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

func TestHubdbRowPurgeBatch(t *testing.T) {
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
	err := client.Cms.Hubdb.Rows.PurgeBatch(
		context.TODO(),
		"tableIdOrName",
		cms.HubdbRowPurgeBatchParams{
			BatchInputString: shared.BatchInputStringParam{
				Inputs: []string{"string"},
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

func TestHubdbRowReplaceBatch(t *testing.T) {
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
	_, err := client.Cms.Hubdb.Rows.ReplaceBatch(
		context.TODO(),
		"tableIdOrName",
		cms.HubdbRowReplaceBatchParams{
			BatchInputHubDBTableRowV3BatchUpdateRequest: cms.BatchInputHubDBTableRowV3BatchUpdateRequestParam{
				Inputs: []cms.HubDBTableRowV3BatchUpdateRequestParam{{
					ChildTableID: 0,
					DisplayIndex: 0,
					Values: map[string]cms.Variant{
						"foo": {
							"foo": "bar",
						},
					},
					ID:   hubspotsdk.String("id"),
					Name: hubspotsdk.String("name"),
					Path: hubspotsdk.String("path"),
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
}

func TestHubdbRowReplaceDraftWithOptionalParams(t *testing.T) {
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
	_, err := client.Cms.Hubdb.Rows.ReplaceDraft(
		context.TODO(),
		"321669910225",
		cms.HubdbRowReplaceDraftParams{
			TableIDOrName: "tableIdOrName",
			HubDBTableRowV3Request: cms.HubDBTableRowV3RequestParam{
				ChildTableID: 0,
				DisplayIndex: 0,
				Values: map[string]cms.Variant{
					"foo": {
						"foo": "bar",
					},
				},
				Name: hubspotsdk.String("name"),
				Path: hubspotsdk.String("path"),
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

func TestHubdbRowUpdateBatch(t *testing.T) {
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
	_, err := client.Cms.Hubdb.Rows.UpdateBatch(
		context.TODO(),
		"tableIdOrName",
		cms.HubdbRowUpdateBatchParams{
			BatchInputHubDBTableRowV3BatchUpdateRequest: cms.BatchInputHubDBTableRowV3BatchUpdateRequestParam{
				Inputs: []cms.HubDBTableRowV3BatchUpdateRequestParam{{
					ChildTableID: 0,
					DisplayIndex: 0,
					Values: map[string]cms.Variant{
						"foo": {
							"foo": "bar",
						},
					},
					ID:   hubspotsdk.String("id"),
					Name: hubspotsdk.String("name"),
					Path: hubspotsdk.String("path"),
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
}

func TestHubdbRowUpdateDraftWithOptionalParams(t *testing.T) {
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
	_, err := client.Cms.Hubdb.Rows.UpdateDraft(
		context.TODO(),
		"321669910225",
		cms.HubdbRowUpdateDraftParams{
			TableIDOrName: "tableIdOrName",
			HubDBTableRowV3Request: cms.HubDBTableRowV3RequestParam{
				ChildTableID: 0,
				DisplayIndex: 0,
				Values: map[string]cms.Variant{
					"foo": {
						"foo": "bar",
					},
				},
				Name: hubspotsdk.String("name"),
				Path: hubspotsdk.String("path"),
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
