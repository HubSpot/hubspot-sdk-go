// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package automation_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/stainless-sdks/hubspot-sdk-go"
	"github.com/stainless-sdks/hubspot-sdk-go/automation"
	"github.com/stainless-sdks/hubspot-sdk-go/internal/testutil"
	"github.com/stainless-sdks/hubspot-sdk-go/option"
)

func TestActionCallbackCompleteWithOptionalParams(t *testing.T) {
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
	err := client.Automation.Actions.Callbacks.Complete(
		context.TODO(),
		"callbackId",
		automation.ActionCallbackCompleteParams{
			CallbackCompletionRequest: automation.CallbackCompletionRequestParam{
				OutputFields: map[string]string{
					"foo": "string",
				},
				TypedOutputs:      map[string]any{},
				FailureReasonType: hubspotsdk.String("failureReasonType"),
				RequestContext: automation.CallbackCompletionRequestRequestContextUnionParam{
					OfWorkflows: &automation.WorkflowsRequestContextParam{
						Source:     automation.WorkflowsRequestContextSourceWorkflows,
						WorkflowID: 0,
						ActionExecutionIndexIdentifier: automation.ActionExecutionIndexIdentifierParam{
							ActionExecutionIndex: 0,
							EnrollmentID:         0,
						},
						ActionID: hubspotsdk.Int(0),
					},
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

func TestActionCallbackCompleteBatch(t *testing.T) {
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
	err := client.Automation.Actions.Callbacks.CompleteBatch(context.TODO(), automation.ActionCallbackCompleteBatchParams{
		BatchInputCallbackCompletionBatchRequest: automation.BatchInputCallbackCompletionBatchRequestParam{
			Inputs: []automation.CallbackCompletionBatchRequestParam{{
				CallbackID: "callbackId",
				OutputFields: map[string]string{
					"foo": "string",
				},
				TypedOutputs:      map[string]any{},
				FailureReasonType: hubspotsdk.String("failureReasonType"),
				RequestContext: automation.CallbackCompletionBatchRequestRequestContextUnionParam{
					OfWorkflows: &automation.WorkflowsRequestContextParam{
						Source:     automation.WorkflowsRequestContextSourceWorkflows,
						WorkflowID: 0,
						ActionExecutionIndexIdentifier: automation.ActionExecutionIndexIdentifierParam{
							ActionExecutionIndex: 0,
							EnrollmentID:         0,
						},
						ActionID: hubspotsdk.Int(0),
					},
				},
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
