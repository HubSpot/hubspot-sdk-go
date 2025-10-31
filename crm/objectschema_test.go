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
	"github.com/stainless-sdks/hubspot-sdk-go/shared"
)

func TestObjectSchemaNewWithOptionalParams(t *testing.T) {
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
	_, err := client.CRM.Objects.Schemas.New(context.TODO(), crm.ObjectSchemaNewParams{
		ObjectSchemaEgg: crm.ObjectSchemaEggParam{
			AssociatedObjects: []string{"CONTACT"},
			Labels: shared.ObjectTypeDefinitionLabelsParam{
				Plural:   hubspotsdk.String("My objects"),
				Singular: hubspotsdk.String("My object"),
			},
			Name: "my_object",
			Properties: []crm.ObjectTypePropertyCreateParam{{
				FieldType:         "select",
				Label:             "My object property",
				Name:              "my_object_property",
				Type:              crm.ObjectTypePropertyCreateTypeEnumeration,
				Description:       hubspotsdk.String("description"),
				DisplayOrder:      hubspotsdk.Int(2),
				FormField:         hubspotsdk.Bool(true),
				GroupName:         hubspotsdk.String("my_object_information"),
				HasUniqueValue:    hubspotsdk.Bool(false),
				Hidden:            hubspotsdk.Bool(true),
				NumberDisplayHint: crm.ObjectTypePropertyCreateNumberDisplayHintUnformatted,
				Options: []shared.OptionInputParam{{
					DisplayOrder: 1,
					Hidden:       true,
					Label:        "Option A",
					Value:        "A",
					Description:  hubspotsdk.String("Choice number one"),
				}, {
					DisplayOrder: 2,
					Hidden:       true,
					Label:        "Option B",
					Value:        "B",
					Description:  hubspotsdk.String("Choice number two"),
				}},
				OptionSortStrategy:       crm.ObjectTypePropertyCreateOptionSortStrategyDisplayOrder,
				ReferencedObjectType:     hubspotsdk.String("referencedObjectType"),
				SearchableInGlobalSearch: hubspotsdk.Bool(true),
				ShowCurrencySymbol:       hubspotsdk.Bool(true),
				TextDisplayHint:          crm.ObjectTypePropertyCreateTextDisplayHintUnformattedSingleLine,
			}},
			RequiredProperties:         []string{"my_object_property"},
			Description:                hubspotsdk.String("description"),
			PrimaryDisplayProperty:     hubspotsdk.String("my_object_property"),
			SearchableProperties:       []string{"string"},
			SecondaryDisplayProperties: []string{"string"},
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

func TestObjectSchemaUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.CRM.Objects.Schemas.Update(
		context.TODO(),
		"objectType",
		crm.ObjectSchemaUpdateParams{
			ObjectTypeDefinitionPatch: crm.ObjectTypeDefinitionPatchParam{
				ClearDescription: hubspotsdk.Bool(true),
				Description:      hubspotsdk.String("description"),
				Labels: shared.ObjectTypeDefinitionLabelsParam{
					Plural:   hubspotsdk.String("plural"),
					Singular: hubspotsdk.String("singular"),
				},
				PrimaryDisplayProperty:     hubspotsdk.String("my_object_property"),
				RequiredProperties:         []string{"my_object_property"},
				Restorable:                 hubspotsdk.Bool(true),
				SearchableProperties:       []string{"my_object_property"},
				SecondaryDisplayProperties: []string{"string"},
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

func TestObjectSchemaListWithOptionalParams(t *testing.T) {
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
	_, err := client.CRM.Objects.Schemas.List(context.TODO(), crm.ObjectSchemaListParams{
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

func TestObjectSchemaDeleteWithOptionalParams(t *testing.T) {
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
	err := client.CRM.Objects.Schemas.Delete(
		context.TODO(),
		"objectType",
		crm.ObjectSchemaDeleteParams{
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

func TestObjectSchemaNewAssociationWithOptionalParams(t *testing.T) {
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
	_, err := client.CRM.Objects.Schemas.NewAssociation(
		context.TODO(),
		"objectType",
		crm.ObjectSchemaNewAssociationParams{
			AssociationDefinitionEgg: shared.AssociationDefinitionEggParam{
				FromObjectTypeID: "fromObjectTypeId",
				ToObjectTypeID:   "toObjectTypeId",
				Name:             hubspotsdk.String("name"),
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

func TestObjectSchemaDeleteAssociation(t *testing.T) {
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
	err := client.CRM.Objects.Schemas.DeleteAssociation(
		context.TODO(),
		"associationIdentifier",
		crm.ObjectSchemaDeleteAssociationParams{
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

func TestObjectSchemaGet(t *testing.T) {
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
	_, err := client.CRM.Objects.Schemas.Get(context.TODO(), "objectType")
	if err != nil {
		var apierr *hubspotsdk.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
