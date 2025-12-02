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
)

func TestMediaBridgeIntegratorSettingNewObjectDefinition(t *testing.T) {
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
	_, err := client.Cms.MediaBridge.IntegratorSettings.NewObjectDefinition(
		context.TODO(),
		0,
		cms.MediaBridgeIntegratorSettingNewObjectDefinitionParams{
			IntegratorObjectCreationRequest: cms.IntegratorObjectCreationRequestParam{
				MediaTypes: []string{"VIDEO"},
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

func TestMediaBridgeIntegratorSettingNewOembedDomainWithOptionalParams(t *testing.T) {
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
	_, err := client.Cms.MediaBridge.IntegratorSettings.NewOembedDomain(
		context.TODO(),
		0,
		cms.MediaBridgeIntegratorSettingNewOembedDomainParams{
			IntegratorOEmbedDomainRequest: cms.IntegratorOEmbedDomainRequestParam{
				Endpoints: cms.EndpointsParam{
					Discovery: true,
					Schemes:   []string{"string"},
					URL:       "url",
				},
				PortalID: hubspotsdk.Int(0),
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

func TestMediaBridgeIntegratorSettingDeleteOembedDomainWithOptionalParams(t *testing.T) {
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
	err := client.Cms.MediaBridge.IntegratorSettings.DeleteOembedDomain(
		context.TODO(),
		0,
		cms.MediaBridgeIntegratorSettingDeleteOembedDomainParams{
			ID:             hubspotsdk.Int(0),
			DomainPortalID: hubspotsdk.Int(0),
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

func TestMediaBridgeIntegratorSettingGetEventVisibilitySettings(t *testing.T) {
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
	_, err := client.Cms.MediaBridge.IntegratorSettings.GetEventVisibilitySettings(context.TODO(), 0)
	if err != nil {
		var apierr *hubspotsdk.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestMediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeWithOptionalParams(t *testing.T) {
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
	_, err := client.Cms.MediaBridge.IntegratorSettings.GetObjectDefinitionsByMediaType(
		context.TODO(),
		cms.MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeParamsMediaTypeAudio,
		cms.MediaBridgeIntegratorSettingGetObjectDefinitionsByMediaTypeParams{
			AppID:                 0,
			IncludeFullDefinition: hubspotsdk.Bool(true),
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

func TestMediaBridgeIntegratorSettingGetOembedDomain(t *testing.T) {
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
	_, err := client.Cms.MediaBridge.IntegratorSettings.GetOembedDomain(
		context.TODO(),
		"oEmbedDomainId",
		cms.MediaBridgeIntegratorSettingGetOembedDomainParams{
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

func TestMediaBridgeIntegratorSettingListOembedDomainsWithOptionalParams(t *testing.T) {
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
	_, err := client.Cms.MediaBridge.IntegratorSettings.ListOembedDomains(
		context.TODO(),
		0,
		cms.MediaBridgeIntegratorSettingListOembedDomainsParams{
			DomainPortalID: hubspotsdk.Int(0),
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

func TestMediaBridgeIntegratorSettingRegisterAppNameWithOptionalParams(t *testing.T) {
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
	_, err := client.Cms.MediaBridge.IntegratorSettings.RegisterAppName(
		context.TODO(),
		0,
		cms.MediaBridgeIntegratorSettingRegisterAppNameParams{
			MediaBridgeProviderPartial: cms.MediaBridgeProviderPartialParam{
				UpdatedAt: 0,
				Name:      hubspotsdk.String("name"),
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

func TestMediaBridgeIntegratorSettingUpdateAppNameWithOptionalParams(t *testing.T) {
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
	_, err := client.Cms.MediaBridge.IntegratorSettings.UpdateAppName(
		context.TODO(),
		0,
		cms.MediaBridgeIntegratorSettingUpdateAppNameParams{
			MediaBridgeProviderPartial: cms.MediaBridgeProviderPartialParam{
				UpdatedAt: 0,
				Name:      hubspotsdk.String("name"),
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

func TestMediaBridgeIntegratorSettingUpdateEventVisibilitySettingsWithOptionalParams(t *testing.T) {
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
	_, err := client.Cms.MediaBridge.IntegratorSettings.UpdateEventVisibilitySettings(
		context.TODO(),
		0,
		cms.MediaBridgeIntegratorSettingUpdateEventVisibilitySettingsParams{
			EventVisibilityChange: cms.EventVisibilityChangeParam{
				EventType:       cms.EventVisibilityChangeEventTypeAll,
				UpdatedAt:       0,
				ShowInReporting: hubspotsdk.Bool(true),
				ShowInTimeline:  hubspotsdk.Bool(true),
				ShowInWorkflows: hubspotsdk.Bool(true),
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

func TestMediaBridgeIntegratorSettingUpdateOembedDomainWithOptionalParams(t *testing.T) {
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
	_, err := client.Cms.MediaBridge.IntegratorSettings.UpdateOembedDomain(
		context.TODO(),
		"oEmbedDomainId",
		cms.MediaBridgeIntegratorSettingUpdateOembedDomainParams{
			AppID: 0,
			IntegratorOEmbedDomainRequest: cms.IntegratorOEmbedDomainRequestParam{
				Endpoints: cms.EndpointsParam{
					Discovery: true,
					Schemes:   []string{"string"},
					URL:       "url",
				},
				PortalID: hubspotsdk.Int(0),
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
