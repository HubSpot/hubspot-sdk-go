// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package automation_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/stainless-sdks/hubspot-sdk-go"
	"github.com/stainless-sdks/hubspot-sdk-go/automation"
	"github.com/stainless-sdks/hubspot-sdk-go/events"
	"github.com/stainless-sdks/hubspot-sdk-go/internal/testutil"
	"github.com/stainless-sdks/hubspot-sdk-go/option"
)

func TestActionNewWithOptionalParams(t *testing.T) {
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
		option.WithDeveloperAPIKey("My Developer API Key"),
	)
	_, err := client.Automation.Actions.New(
		context.TODO(),
		0,
		automation.ActionNewParams{
			PublicActionDefinitionEgg: automation.PublicActionDefinitionEggParam{
				ActionURL: "actionUrl",
				Functions: []automation.PublicActionFunctionParam{{
					FunctionSource: "functionSource",
					FunctionType:   automation.PublicActionFunctionFunctionTypePostActionExecution,
					ID:             hubspotsdk.String("id"),
				}},
				InputFields: []automation.PublicInputFieldDefinitionParam{{
					IsRequired: true,
					TypeDefinition: automation.PublicFieldTypeDefinitionParam{
						Name: "name",
						Options: []automation.PublicOptionParam{{
							Label:        "label",
							Value:        "value",
							Description:  hubspotsdk.String("description"),
							DisplayOrder: hubspotsdk.Int(0),
						}},
						Type:                 automation.PublicFieldTypeDefinitionTypeBool,
						Description:          hubspotsdk.String("description"),
						FieldType:            automation.PublicFieldTypeDefinitionFieldTypeBooleancheckbox,
						HelpText:             hubspotsdk.String("helpText"),
						Label:                hubspotsdk.String("label"),
						OptionsURL:           hubspotsdk.String("optionsUrl"),
						ReferencedObjectType: automation.PublicFieldTypeDefinitionReferencedObjectTypeOwner,
					},
					SupportedValueTypes: []string{"STATIC_VALUE"},
				}},
				Labels: map[string]automation.PublicActionLabelsParam{
					"foo": {
						ActionName:        "actionName",
						ActionCardContent: hubspotsdk.String("actionCardContent"),
						ActionDescription: hubspotsdk.String("actionDescription"),
						AppDisplayName:    hubspotsdk.String("appDisplayName"),
						ExecutionRules: map[string]string{
							"foo": "string",
						},
						InputFieldDescriptions: map[string]string{
							"foo": "string",
						},
						InputFieldLabels: map[string]string{
							"foo": "string",
						},
						InputFieldOptionLabels: map[string]map[string]string{
							"foo": {
								"foo": "string",
							},
						},
						OutputFieldLabels: map[string]string{
							"foo": "string",
						},
					},
				},
				ObjectTypes: []string{"string"},
				Published:   true,
				ArchivedAt:  hubspotsdk.Int(0),
				ExecutionRules: []automation.PublicExecutionTranslationRuleParam{{
					Conditions: map[string]any{
						"foo": map[string]any{},
					},
					LabelName: "labelName",
				}},
				InputFieldDependencies: []automation.PublicActionDefinitionEggInputFieldDependencyUnionParam{{
					OfSingleField: &automation.PublicSingleFieldDependencyParam{
						ControllingFieldName: "controllingFieldName",
						DependencyType:       automation.PublicSingleFieldDependencyDependencyTypeSingleField,
						DependentFieldNames:  []string{"string"},
					},
				}},
				ObjectRequestOptions: automation.PublicObjectRequestOptionsParam{
					Properties: []string{"string"},
				},
				OutputFields: []automation.OutputFieldDefinitionParam{{
					TypeDefinition: automation.FieldTypeDefinitionParam{
						ExternalOptions: true,
						Name:            "name",
						Options: []events.OptionParam{{
							Hidden:       true,
							Label:        "label",
							Value:        "value",
							Description:  hubspotsdk.String("description"),
							DisplayOrder: hubspotsdk.Int(0),
						}},
						Schema: automation.FieldTypeDefinitionSchemaUnionParam{
							OfInteger: &automation.IntegerFieldSchemaParam{
								Type:    automation.IntegerFieldSchemaTypeInteger,
								Maximum: hubspotsdk.Int(0),
								Minimum: hubspotsdk.Int(0),
							},
						},
						Type:                         automation.FieldTypeDefinitionTypeBool,
						UseChirp:                     true,
						Description:                  hubspotsdk.String("description"),
						ExternalOptionsReferenceType: hubspotsdk.String("externalOptionsReferenceType"),
						FieldType:                    automation.FieldTypeDefinitionFieldTypeBooleancheckbox,
						HelpText:                     hubspotsdk.String("helpText"),
						Label:                        hubspotsdk.String("label"),
						OptionsURL:                   hubspotsdk.String("optionsUrl"),
						ReferencedObjectType:         automation.FieldTypeDefinitionReferencedObjectTypeAbandonedCart,
					},
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

func TestActionUpdateWithOptionalParams(t *testing.T) {
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
		option.WithDeveloperAPIKey("My Developer API Key"),
	)
	_, err := client.Automation.Actions.Update(
		context.TODO(),
		"definitionId",
		automation.ActionUpdateParams{
			AppID: 0,
			PublicActionDefinitionPatch: automation.PublicActionDefinitionPatchParam{
				ActionURL: hubspotsdk.String("actionUrl"),
				ExecutionRules: []automation.PublicExecutionTranslationRuleParam{{
					Conditions: map[string]any{
						"foo": map[string]any{},
					},
					LabelName: "labelName",
				}},
				InputFieldDependencies: []automation.PublicActionDefinitionPatchInputFieldDependencyUnionParam{{
					OfSingleField: &automation.PublicSingleFieldDependencyParam{
						ControllingFieldName: "controllingFieldName",
						DependencyType:       automation.PublicSingleFieldDependencyDependencyTypeSingleField,
						DependentFieldNames:  []string{"string"},
					},
				}},
				InputFields: []automation.PublicInputFieldDefinitionParam{{
					IsRequired: true,
					TypeDefinition: automation.PublicFieldTypeDefinitionParam{
						Name: "name",
						Options: []automation.PublicOptionParam{{
							Label:        "label",
							Value:        "value",
							Description:  hubspotsdk.String("description"),
							DisplayOrder: hubspotsdk.Int(0),
						}},
						Type:                 automation.PublicFieldTypeDefinitionTypeBool,
						Description:          hubspotsdk.String("description"),
						FieldType:            automation.PublicFieldTypeDefinitionFieldTypeBooleancheckbox,
						HelpText:             hubspotsdk.String("helpText"),
						Label:                hubspotsdk.String("label"),
						OptionsURL:           hubspotsdk.String("optionsUrl"),
						ReferencedObjectType: automation.PublicFieldTypeDefinitionReferencedObjectTypeOwner,
					},
					SupportedValueTypes: []string{"STATIC_VALUE"},
				}},
				Labels: map[string]automation.PublicActionLabelsParam{
					"foo": {
						ActionName:        "actionName",
						ActionCardContent: hubspotsdk.String("actionCardContent"),
						ActionDescription: hubspotsdk.String("actionDescription"),
						AppDisplayName:    hubspotsdk.String("appDisplayName"),
						ExecutionRules: map[string]string{
							"foo": "string",
						},
						InputFieldDescriptions: map[string]string{
							"foo": "string",
						},
						InputFieldLabels: map[string]string{
							"foo": "string",
						},
						InputFieldOptionLabels: map[string]map[string]string{
							"foo": {
								"foo": "string",
							},
						},
						OutputFieldLabels: map[string]string{
							"foo": "string",
						},
					},
				},
				ObjectRequestOptions: automation.PublicObjectRequestOptionsParam{
					Properties: []string{"string"},
				},
				ObjectTypes: []string{"string"},
				OutputFields: []automation.OutputFieldDefinitionParam{{
					TypeDefinition: automation.FieldTypeDefinitionParam{
						ExternalOptions: true,
						Name:            "name",
						Options: []events.OptionParam{{
							Hidden:       true,
							Label:        "label",
							Value:        "value",
							Description:  hubspotsdk.String("description"),
							DisplayOrder: hubspotsdk.Int(0),
						}},
						Schema: automation.FieldTypeDefinitionSchemaUnionParam{
							OfInteger: &automation.IntegerFieldSchemaParam{
								Type:    automation.IntegerFieldSchemaTypeInteger,
								Maximum: hubspotsdk.Int(0),
								Minimum: hubspotsdk.Int(0),
							},
						},
						Type:                         automation.FieldTypeDefinitionTypeBool,
						UseChirp:                     true,
						Description:                  hubspotsdk.String("description"),
						ExternalOptionsReferenceType: hubspotsdk.String("externalOptionsReferenceType"),
						FieldType:                    automation.FieldTypeDefinitionFieldTypeBooleancheckbox,
						HelpText:                     hubspotsdk.String("helpText"),
						Label:                        hubspotsdk.String("label"),
						OptionsURL:                   hubspotsdk.String("optionsUrl"),
						ReferencedObjectType:         automation.FieldTypeDefinitionReferencedObjectTypeAbandonedCart,
					},
				}},
				Published: hubspotsdk.Bool(true),
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

func TestActionListWithOptionalParams(t *testing.T) {
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
		option.WithDeveloperAPIKey("My Developer API Key"),
	)
	_, err := client.Automation.Actions.List(
		context.TODO(),
		"definitionId",
		automation.ActionListParams{
			AppID: 0,
			After: hubspotsdk.String("after"),
			Limit: hubspotsdk.Int(0),
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

func TestActionDelete(t *testing.T) {
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
		option.WithDeveloperAPIKey("My Developer API Key"),
	)
	err := client.Automation.Actions.Delete(
		context.TODO(),
		"functionId",
		automation.ActionDeleteParams{
			AppID:        0,
			DefinitionID: "definitionId",
			FunctionType: automation.ActionDeleteParamsFunctionTypePostActionExecution,
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

func TestActionCompleteWithOptionalParams(t *testing.T) {
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
	err := client.Automation.Actions.Complete(
		context.TODO(),
		"callbackId",
		automation.ActionCompleteParams{
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

func TestActionCompleteBatch(t *testing.T) {
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
	err := client.Automation.Actions.CompleteBatch(context.TODO(), automation.ActionCompleteBatchParams{
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

func TestActionNewOrReplace(t *testing.T) {
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
		option.WithDeveloperAPIKey("My Developer API Key"),
	)
	_, err := client.Automation.Actions.NewOrReplace(
		context.TODO(),
		"functionId",
		automation.ActionNewOrReplaceParams{
			AppID:        0,
			DefinitionID: "definitionId",
			FunctionType: automation.ActionNewOrReplaceParamsFunctionTypePostActionExecution,
			Body:         "body",
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

func TestActionNewOrReplaceByFunctionType(t *testing.T) {
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
		option.WithDeveloperAPIKey("My Developer API Key"),
	)
	_, err := client.Automation.Actions.NewOrReplaceByFunctionType(
		context.TODO(),
		automation.ActionNewOrReplaceByFunctionTypeParamsFunctionTypePostActionExecution,
		automation.ActionNewOrReplaceByFunctionTypeParams{
			AppID:        0,
			DefinitionID: "definitionId",
			Body:         "body",
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

func TestActionNewRequiresObject(t *testing.T) {
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
		option.WithDeveloperAPIKey("My Developer API Key"),
	)
	err := client.Automation.Actions.NewRequiresObject(
		context.TODO(),
		"definitionId",
		automation.ActionNewRequiresObjectParams{
			AppID: 0,
			PublicActionDefinitionRequiresObjectRequest: automation.PublicActionDefinitionRequiresObjectRequestParam{
				RequiresObject: true,
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

func TestActionDeleteByFunctionType(t *testing.T) {
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
		option.WithDeveloperAPIKey("My Developer API Key"),
	)
	err := client.Automation.Actions.DeleteByFunctionType(
		context.TODO(),
		automation.ActionDeleteByFunctionTypeParamsFunctionTypePostActionExecution,
		automation.ActionDeleteByFunctionTypeParams{
			AppID:        0,
			DefinitionID: "definitionId",
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

func TestActionGet(t *testing.T) {
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
		option.WithDeveloperAPIKey("My Developer API Key"),
	)
	_, err := client.Automation.Actions.Get(
		context.TODO(),
		"revisionId",
		automation.ActionGetParams{
			AppID:        0,
			DefinitionID: "definitionId",
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

func TestActionGetByFunctionType(t *testing.T) {
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
		option.WithDeveloperAPIKey("My Developer API Key"),
	)
	_, err := client.Automation.Actions.GetByFunctionType(
		context.TODO(),
		automation.ActionGetByFunctionTypeParamsFunctionTypePostActionExecution,
		automation.ActionGetByFunctionTypeParams{
			AppID:        0,
			DefinitionID: "definitionId",
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

func TestActionGetRequiresObject(t *testing.T) {
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
		option.WithDeveloperAPIKey("My Developer API Key"),
	)
	_, err := client.Automation.Actions.GetRequiresObject(
		context.TODO(),
		"definitionId",
		automation.ActionGetRequiresObjectParams{
			AppID: 0,
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
