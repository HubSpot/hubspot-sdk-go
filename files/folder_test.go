// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package files_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/stainless-sdks/hubspot-sdk-go"
	"github.com/stainless-sdks/hubspot-sdk-go/files"
	"github.com/stainless-sdks/hubspot-sdk-go/internal/testutil"
	"github.com/stainless-sdks/hubspot-sdk-go/option"
)

func TestFolderGetUpdateAsyncStatus(t *testing.T) {
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
	_, err := client.Files.Folders.GetUpdateAsyncStatus(context.TODO(), "taskId")
	if err != nil {
		var apierr *hubspotsdk.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestFolderSearchWithOptionalParams(t *testing.T) {
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
	_, err := client.Files.Folders.Search(context.TODO(), files.FolderSearchParams{
		After:           hubspotsdk.String("after"),
		Before:          hubspotsdk.String("before"),
		CreatedAt:       hubspotsdk.Time(time.Now()),
		CreatedAtGte:    hubspotsdk.Time(time.Now()),
		CreatedAtLte:    hubspotsdk.Time(time.Now()),
		IDGte:           hubspotsdk.Int(0),
		IDLte:           hubspotsdk.Int(0),
		IDs:             []int64{0},
		Limit:           hubspotsdk.Int(0),
		Name:            hubspotsdk.String("name"),
		ParentFolderIDs: []int64{0},
		Path:            hubspotsdk.String("path"),
		Properties:      []string{"string"},
		Sort:            []string{"string"},
		UpdatedAt:       hubspotsdk.Time(time.Now()),
		UpdatedAtGte:    hubspotsdk.Time(time.Now()),
		UpdatedAtLte:    hubspotsdk.Time(time.Now()),
	})
	if err != nil {
		var apierr *hubspotsdk.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestFolderUpdateAsyncByIDWithOptionalParams(t *testing.T) {
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
	_, err := client.Files.Folders.UpdateAsyncByID(context.TODO(), files.FolderUpdateAsyncByIDParams{
		FolderUpdateInputWithID: files.FolderUpdateInputWithIDParam{
			ID:             "id",
			Name:           hubspotsdk.String("name"),
			ParentFolderID: hubspotsdk.Int(0),
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

func TestFolderUpdateByIDWithOptionalParams(t *testing.T) {
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
	_, err := client.Files.Folders.UpdateByID(
		context.TODO(),
		"321669910225",
		files.FolderUpdateByIDParams{
			FolderUpdateInput: files.FolderUpdateInputParam{
				Name:           hubspotsdk.String("name"),
				ParentFolderID: hubspotsdk.Int(0),
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
