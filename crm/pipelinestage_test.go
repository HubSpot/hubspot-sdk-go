// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package crm_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/stainless-sdks/hubspot-sdk-go"
	"github.com/stainless-sdks/hubspot-sdk-go/crm"
	"github.com/stainless-sdks/hubspot-sdk-go/internal/testutil"
	"github.com/stainless-sdks/hubspot-sdk-go/option"
)

func TestPipelineStageNewWithOptionalParams(t *testing.T) {
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
	_, err := client.CRM.Pipelines.Stages.New(
		context.TODO(),
		"pipelineId",
		crm.PipelineStageNewParams{
			ObjectType: "objectType",
			PipelineStageInput: crm.PipelineStageInputParam{
				DisplayOrder: 1,
				Label:        "Done",
				Metadata: map[string]string{
					"ticketState": "CLOSED",
				},
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

func TestPipelineStageUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.CRM.Pipelines.Stages.Update(
		context.TODO(),
		"stageId",
		crm.PipelineStageUpdateParams{
			ObjectType: "objectType",
			PipelineID: "pipelineId",
			PipelineStagePatchInput: crm.PipelineStagePatchInputParam{
				Archived:     hubspotsdk.Bool(true),
				DisplayOrder: hubspotsdk.Int(1),
				Label:        hubspotsdk.String("Done"),
				Metadata: map[string]string{
					"ticketState": "CLOSED",
				},
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

func TestPipelineStageList(t *testing.T) {
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
	_, err := client.CRM.Pipelines.Stages.List(
		context.TODO(),
		"pipelineId",
		crm.PipelineStageListParams{
			ObjectType: "objectType",
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

func TestPipelineStageDelete(t *testing.T) {
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
	err := client.CRM.Pipelines.Stages.Delete(
		context.TODO(),
		"stageId",
		crm.PipelineStageDeleteParams{
			ObjectType: "objectType",
			PipelineID: "pipelineId",
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

func TestPipelineStageGet(t *testing.T) {
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
	_, err := client.CRM.Pipelines.Stages.Get(
		context.TODO(),
		"stageId",
		crm.PipelineStageGetParams{
			ObjectType: "objectType",
			PipelineID: "pipelineId",
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

func TestPipelineStageGetAudit(t *testing.T) {
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
	_, err := client.CRM.Pipelines.Stages.GetAudit(
		context.TODO(),
		"stageId",
		crm.PipelineStageGetAuditParams{
			ObjectType: "objectType",
			PipelineID: "pipelineId",
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

func TestPipelineStageReplaceWithOptionalParams(t *testing.T) {
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
	_, err := client.CRM.Pipelines.Stages.Replace(
		context.TODO(),
		"stageId",
		crm.PipelineStageReplaceParams{
			ObjectType: "objectType",
			PipelineID: "pipelineId",
			PipelineStageInput: crm.PipelineStageInputParam{
				DisplayOrder: 1,
				Label:        "Done",
				Metadata: map[string]string{
					"ticketState": "CLOSED",
				},
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
