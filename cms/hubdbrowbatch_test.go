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

func TestHubdbRowBatchCloneBatch(t *testing.T) {
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
	_, err := client.Cms.Hubdb.Rows.Batch.CloneBatch(
		context.TODO(),
		"tableIdOrName",
		cms.HubdbRowBatchCloneBatchParams{
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

func TestHubdbRowBatchNewBatch(t *testing.T) {
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
	_, err := client.Cms.Hubdb.Rows.Batch.NewBatch(
		context.TODO(),
		"tableIdOrName",
		cms.HubdbRowBatchNewBatchParams{
			BatchInputHubDBTableRowV3Request: cms.BatchInputHubDBTableRowV3RequestParam{
				Inputs: []cms.HubDBTableRowV3RequestParam{{
					Values: map[string]cms.Variant{
						"foo": map[string]interface{}{},
					},
					ChildTableID: hubspotsdk.Int(0),
					DisplayIndex: hubspotsdk.Int(0),
					Name:         hubspotsdk.String("name"),
					Path:         hubspotsdk.String("path"),
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

func TestHubdbRowBatchGetBatch(t *testing.T) {
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
	_, err := client.Cms.Hubdb.Rows.Batch.GetBatch(
		context.TODO(),
		"tableIdOrName",
		cms.HubdbRowBatchGetBatchParams{
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

func TestHubdbRowBatchGetDraftBatch(t *testing.T) {
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
	_, err := client.Cms.Hubdb.Rows.Batch.GetDraftBatch(
		context.TODO(),
		"tableIdOrName",
		cms.HubdbRowBatchGetDraftBatchParams{
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

func TestHubdbRowBatchPurgeBatch(t *testing.T) {
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
	err := client.Cms.Hubdb.Rows.Batch.PurgeBatch(
		context.TODO(),
		"tableIdOrName",
		cms.HubdbRowBatchPurgeBatchParams{
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

func TestHubdbRowBatchReplaceBatch(t *testing.T) {
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
	_, err := client.Cms.Hubdb.Rows.Batch.ReplaceBatch(
		context.TODO(),
		"tableIdOrName",
		cms.HubdbRowBatchReplaceBatchParams{
			BatchInputHubDBTableRowV3BatchUpdateRequest: cms.BatchInputHubDBTableRowV3BatchUpdateRequestParam{
				Inputs: []cms.HubDBTableRowV3BatchUpdateRequestParam{{
					ID: "id",
					Values: map[string]cms.Variant{
						"foo": map[string]interface{}{},
					},
					ChildTableID: hubspotsdk.Int(0),
					DisplayIndex: hubspotsdk.Int(0),
					Name:         hubspotsdk.String("name"),
					Path:         hubspotsdk.String("path"),
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

func TestHubdbRowBatchUpdateBatch(t *testing.T) {
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
	_, err := client.Cms.Hubdb.Rows.Batch.UpdateBatch(
		context.TODO(),
		"tableIdOrName",
		cms.HubdbRowBatchUpdateBatchParams{
			BatchInputHubDBTableRowV3BatchUpdateRequest: cms.BatchInputHubDBTableRowV3BatchUpdateRequestParam{
				Inputs: []cms.HubDBTableRowV3BatchUpdateRequestParam{{
					ID: "id",
					Values: map[string]cms.Variant{
						"foo": map[string]interface{}{},
					},
					ChildTableID: hubspotsdk.Int(0),
					DisplayIndex: hubspotsdk.Int(0),
					Name:         hubspotsdk.String("name"),
					Path:         hubspotsdk.String("path"),
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
