// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package automation_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/HubSpot/hubspot-sdk-go"
	"github.com/HubSpot/hubspot-sdk-go/automation"
	"github.com/HubSpot/hubspot-sdk-go/internal/testutil"
	"github.com/HubSpot/hubspot-sdk-go/option"
	"github.com/HubSpot/hubspot-sdk-go/shared"
)

func TestActionDefinitionNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Automation.Actions.Definitions.New(
		context.TODO(),
		0,
		automation.ActionDefinitionNewParams{
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
						Options: []shared.AutomationActionsOptionParam{{
							Description:  "description",
							DisplayOrder: 0,
							DoubleData:   0,
							Hidden:       true,
							Label:        "label",
							ReadOnly:     true,
							Value:        "value",
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

func TestActionDefinitionUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Automation.Actions.Definitions.Update(
		context.TODO(),
		"definitionId",
		automation.ActionDefinitionUpdateParams{
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
						Options: []shared.AutomationActionsOptionParam{{
							Description:  "description",
							DisplayOrder: 0,
							DoubleData:   0,
							Hidden:       true,
							Label:        "label",
							ReadOnly:     true,
							Value:        "value",
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

func TestActionDefinitionListWithOptionalParams(t *testing.T) {
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
	_, err := client.Automation.Actions.Definitions.List(
		context.TODO(),
		0,
		automation.ActionDefinitionListParams{
			After:    hubspotsdk.String("after"),
			Archived: hubspotsdk.Bool(true),
			Limit:    hubspotsdk.Int(0),
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

func TestActionDefinitionDelete(t *testing.T) {
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
	err := client.Automation.Actions.Definitions.Delete(
		context.TODO(),
		"definitionId",
		automation.ActionDefinitionDeleteParams{
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

func TestActionDefinitionNewRequiresObject(t *testing.T) {
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
	err := client.Automation.Actions.Definitions.NewRequiresObject(
		context.TODO(),
		"definitionId",
		automation.ActionDefinitionNewRequiresObjectParams{
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

func TestActionDefinitionGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Automation.Actions.Definitions.Get(
		context.TODO(),
		"definitionId",
		automation.ActionDefinitionGetParams{
			AppID:    0,
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

func TestActionDefinitionGetRequiresObject(t *testing.T) {
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
	_, err := client.Automation.Actions.Definitions.GetRequiresObject(
		context.TODO(),
		"definitionId",
		automation.ActionDefinitionGetRequiresObjectParams{
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
