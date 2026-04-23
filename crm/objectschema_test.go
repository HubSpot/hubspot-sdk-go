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
	_, err := client.Crm.ObjectSchemas.New(context.TODO(), crm.ObjectSchemaNewParams{
		ObjectSchemaEgg: crm.ObjectSchemaEggParam{
			AllowsSensitiveProperties: true,
			AssociatedObjects:         []string{"string"},
			Labels: shared.ObjectTypeDefinitionLabelsParam{
				Plural:   hubspotsdk.String("plural"),
				Singular: hubspotsdk.String("singular"),
			},
			Name: "name",
			Properties: []crm.ObjectTypePropertyCreateParam{{
				FieldType:                    "fieldType",
				Label:                        "label",
				Name:                         "name",
				Type:                         crm.ObjectTypePropertyCreateTypeBool,
				Description:                  hubspotsdk.String("description"),
				DisplayOrder:                 hubspotsdk.Int(0),
				ExternalOptionsReferenceType: hubspotsdk.String("externalOptionsReferenceType"),
				FormField:                    hubspotsdk.Bool(true),
				GroupName:                    hubspotsdk.String("groupName"),
				HasUniqueValue:               hubspotsdk.Bool(true),
				Hidden:                       hubspotsdk.Bool(true),
				NumberDisplayHint:            crm.ObjectTypePropertyCreateNumberDisplayHintCurrency,
				Options: []shared.OptionInputParam{{
					DisplayOrder: 0,
					Hidden:       true,
					Label:        "label",
					Value:        "value",
					Description:  hubspotsdk.String("description"),
				}},
				OptionSortStrategy:       crm.ObjectTypePropertyCreateOptionSortStrategyAlphabetical,
				ReferencedObjectType:     hubspotsdk.String("referencedObjectType"),
				SearchableInGlobalSearch: hubspotsdk.Bool(true),
				ShowCurrencySymbol:       hubspotsdk.Bool(true),
				TextDisplayHint:          crm.ObjectTypePropertyCreateTextDisplayHintDomainName,
			}},
			RequiredProperties:         []string{"string"},
			SearchableProperties:       []string{"string"},
			SecondaryDisplayProperties: []string{"string"},
			Description:                hubspotsdk.String("description"),
			PrimaryDisplayProperty:     hubspotsdk.String("primaryDisplayProperty"),
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
	_, err := client.Crm.ObjectSchemas.Update(
		context.TODO(),
		"objectType",
		crm.ObjectSchemaUpdateParams{
			ObjectTypeDefinitionPatch: shared.ObjectTypeDefinitionPatchParam{
				ClearDescription:          true,
				AllowsSensitiveProperties: hubspotsdk.Bool(true),
				Description:               hubspotsdk.String("description"),
				Labels: shared.ObjectTypeDefinitionLabelsParam{
					Plural:   hubspotsdk.String("plural"),
					Singular: hubspotsdk.String("singular"),
				},
				PrimaryDisplayProperty:     hubspotsdk.String("primaryDisplayProperty"),
				RequiredProperties:         []string{"string"},
				Restorable:                 hubspotsdk.Bool(true),
				SearchableProperties:       []string{"string"},
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
	_, err := client.Crm.ObjectSchemas.List(context.TODO(), crm.ObjectSchemaListParams{
		Archived:                      hubspotsdk.Bool(true),
		IncludeAssociationDefinitions: hubspotsdk.Bool(true),
		IncludeAuditMetadata:          hubspotsdk.Bool(true),
		IncludePropertyDefinitions:    hubspotsdk.Bool(true),
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
	err := client.Crm.ObjectSchemas.Delete(
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
	_, err := client.Crm.ObjectSchemas.NewAssociation(
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
	err := client.Crm.ObjectSchemas.DeleteAssociation(
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

func TestObjectSchemaGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Crm.ObjectSchemas.Get(
		context.TODO(),
		"objectType",
		crm.ObjectSchemaGetParams{
			IncludeAssociationDefinitions: hubspotsdk.Bool(true),
			IncludeAuditMetadata:          hubspotsdk.Bool(true),
			IncludePropertyDefinitions:    hubspotsdk.Bool(true),
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
