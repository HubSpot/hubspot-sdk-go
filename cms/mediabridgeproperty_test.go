// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cms_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/stainless-sdks/hubspot-sdk-go"
	"github.com/stainless-sdks/hubspot-sdk-go/cms"
	"github.com/stainless-sdks/hubspot-sdk-go/crm"
	"github.com/stainless-sdks/hubspot-sdk-go/internal/testutil"
	"github.com/stainless-sdks/hubspot-sdk-go/option"
	"github.com/stainless-sdks/hubspot-sdk-go/shared"
)

func TestMediaBridgePropertyNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Cms.MediaBridge.Properties.New(
		context.TODO(),
		"objectType",
		cms.MediaBridgePropertyNewParams{
			AppID: "appId",
			PropertyCreate: shared.PropertyCreateParam{
				FieldType:          shared.PropertyCreateFieldTypeBooleancheckbox,
				GroupName:          "groupName",
				Label:              "label",
				Name:               "name",
				Type:               shared.PropertyCreateTypeBool,
				CalculationFormula: hubspotsdk.String("calculationFormula"),
				DataSensitivity:    shared.PropertyCreateDataSensitivityNonSensitive,
				Description:        hubspotsdk.String("description"),
				DisplayOrder:       hubspotsdk.Int(0),
				ExternalOptions:    hubspotsdk.Bool(true),
				FormField:          hubspotsdk.Bool(true),
				HasUniqueValue:     hubspotsdk.Bool(true),
				Hidden:             hubspotsdk.Bool(true),
				Options: []shared.OptionInputParam{{
					DisplayOrder: 0,
					Hidden:       true,
					Label:        "label",
					Value:        "value",
					Description:  hubspotsdk.String("description"),
				}},
				ReferencedObjectType: hubspotsdk.String("referencedObjectType"),
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

func TestMediaBridgePropertyUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Cms.MediaBridge.Properties.Update(
		context.TODO(),
		"propertyName",
		cms.MediaBridgePropertyUpdateParams{
			AppID:              "appId",
			ObjectType:         "objectType",
			CalculationFormula: hubspotsdk.String("calculationFormula"),
			Description:        hubspotsdk.String("description"),
			DisplayOrder:       hubspotsdk.Int(0),
			FieldType:          cms.MediaBridgePropertyUpdateParamsFieldTypeBooleancheckbox,
			FormField:          hubspotsdk.Bool(true),
			GroupName:          hubspotsdk.String("groupName"),
			HasUniqueValue:     hubspotsdk.Bool(true),
			Hidden:             hubspotsdk.Bool(true),
			Label:              hubspotsdk.String("label"),
			Options: []shared.OptionInputParam{{
				DisplayOrder: 0,
				Hidden:       true,
				Label:        "label",
				Value:        "value",
				Description:  hubspotsdk.String("description"),
			}},
			Type: cms.MediaBridgePropertyUpdateParamsTypeBool,
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

func TestMediaBridgePropertyList(t *testing.T) {
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
	_, err := client.Cms.MediaBridge.Properties.List(
		context.TODO(),
		"objectType",
		cms.MediaBridgePropertyListParams{
			AppID: "appId",
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

func TestMediaBridgePropertyDelete(t *testing.T) {
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
	err := client.Cms.MediaBridge.Properties.Delete(
		context.TODO(),
		"propertyName",
		cms.MediaBridgePropertyDeleteParams{
			AppID:      "appId",
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

func TestMediaBridgePropertyArchiveBatch(t *testing.T) {
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
	err := client.Cms.MediaBridge.Properties.ArchiveBatch(
		context.TODO(),
		"objectType",
		cms.MediaBridgePropertyArchiveBatchParams{
			AppID: "appId",
			BatchInputPropertyName: shared.BatchInputPropertyNameParam{
				Inputs: []shared.PropertyNameParam{{
					Name: "name",
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

func TestMediaBridgePropertyNewBatch(t *testing.T) {
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
	_, err := client.Cms.MediaBridge.Properties.NewBatch(
		context.TODO(),
		"objectType",
		cms.MediaBridgePropertyNewBatchParams{
			AppID: "appId",
			BatchInputPropertyCreate: shared.BatchInputPropertyCreateParam{
				Inputs: []shared.PropertyCreateParam{{
					FieldType:          shared.PropertyCreateFieldTypeBooleancheckbox,
					GroupName:          "groupName",
					Label:              "label",
					Name:               "name",
					Type:               shared.PropertyCreateTypeBool,
					CalculationFormula: hubspotsdk.String("calculationFormula"),
					DataSensitivity:    shared.PropertyCreateDataSensitivityNonSensitive,
					Description:        hubspotsdk.String("description"),
					DisplayOrder:       hubspotsdk.Int(0),
					ExternalOptions:    hubspotsdk.Bool(true),
					FormField:          hubspotsdk.Bool(true),
					HasUniqueValue:     hubspotsdk.Bool(true),
					Hidden:             hubspotsdk.Bool(true),
					Options: []shared.OptionInputParam{{
						DisplayOrder: 0,
						Hidden:       true,
						Label:        "label",
						Value:        "value",
						Description:  hubspotsdk.String("description"),
					}},
					ReferencedObjectType: hubspotsdk.String("referencedObjectType"),
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

func TestMediaBridgePropertyGet(t *testing.T) {
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
	_, err := client.Cms.MediaBridge.Properties.Get(
		context.TODO(),
		"propertyName",
		cms.MediaBridgePropertyGetParams{
			AppID:      "appId",
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

func TestMediaBridgePropertyGetBatchWithOptionalParams(t *testing.T) {
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
	_, err := client.Cms.MediaBridge.Properties.GetBatch(
		context.TODO(),
		"objectType",
		cms.MediaBridgePropertyGetBatchParams{
			AppID: "appId",
			BatchReadInputPropertyName: crm.BatchReadInputPropertyNameParam{
				Archived: true,
				Inputs: []shared.PropertyNameParam{{
					Name: "name",
				}},
				DataSensitivity: crm.BatchReadInputPropertyNameDataSensitivityNonSensitive,
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
