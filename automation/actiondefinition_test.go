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
	"github.com/stainless-sdks/hubspot-sdk-go/shared"
)

func TestActionDefinitionNewWithOptionalParams(t *testing.T) {
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
				InputFields: []automation.InputFieldDefinitionParam{{
					IsRequired: true,
					TypeDefinition: automation.FieldTypeDefinitionParam{
						ExternalOptions: true,
						Name:            "name",
						Options: []shared.OptionParam{{
							Hidden:       false,
							Label:        "Option A",
							Value:        "A",
							Description:  hubspotsdk.String("Choice number one"),
							DisplayOrder: hubspotsdk.Int(1),
						}},
						Type:                         automation.FieldTypeDefinitionTypeBool,
						Description:                  hubspotsdk.String("description"),
						ExternalOptionsReferenceType: hubspotsdk.String("externalOptionsReferenceType"),
						FieldType:                    automation.FieldTypeDefinitionFieldTypeBooleancheckbox,
						HelpText:                     hubspotsdk.String("helpText"),
						Label:                        hubspotsdk.String("label"),
						OptionsURL:                   hubspotsdk.String("optionsUrl"),
						ReferencedObjectType:         automation.FieldTypeDefinitionReferencedObjectTypeAbandonedCart,
					},
					AutomationFieldType: hubspotsdk.String("automationFieldType"),
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
						Options: []shared.OptionParam{{
							Hidden:       false,
							Label:        "Option A",
							Value:        "A",
							Description:  hubspotsdk.String("Choice number one"),
							DisplayOrder: hubspotsdk.Int(1),
						}},
						Type:                         automation.FieldTypeDefinitionTypeBool,
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
				InputFields: []automation.InputFieldDefinitionParam{{
					IsRequired: true,
					TypeDefinition: automation.FieldTypeDefinitionParam{
						ExternalOptions: true,
						Name:            "name",
						Options: []shared.OptionParam{{
							Hidden:       false,
							Label:        "Option A",
							Value:        "A",
							Description:  hubspotsdk.String("Choice number one"),
							DisplayOrder: hubspotsdk.Int(1),
						}},
						Type:                         automation.FieldTypeDefinitionTypeBool,
						Description:                  hubspotsdk.String("description"),
						ExternalOptionsReferenceType: hubspotsdk.String("externalOptionsReferenceType"),
						FieldType:                    automation.FieldTypeDefinitionFieldTypeBooleancheckbox,
						HelpText:                     hubspotsdk.String("helpText"),
						Label:                        hubspotsdk.String("label"),
						OptionsURL:                   hubspotsdk.String("optionsUrl"),
						ReferencedObjectType:         automation.FieldTypeDefinitionReferencedObjectTypeAbandonedCart,
					},
					AutomationFieldType: hubspotsdk.String("automationFieldType"),
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
						Options: []shared.OptionParam{{
							Hidden:       false,
							Label:        "Option A",
							Value:        "A",
							Description:  hubspotsdk.String("Choice number one"),
							DisplayOrder: hubspotsdk.Int(1),
						}},
						Type:                         automation.FieldTypeDefinitionTypeBool,
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

func TestActionDefinitionGetWithOptionalParams(t *testing.T) {
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
