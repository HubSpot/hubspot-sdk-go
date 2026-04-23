// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package files_test

import (
	"bytes"
	"context"
	"errors"
	"io"
	"os"
	"testing"
	"time"

	"github.com/stainless-sdks/hubspot-sdk-go"
	"github.com/stainless-sdks/hubspot-sdk-go/files"
	"github.com/stainless-sdks/hubspot-sdk-go/internal/testutil"
	"github.com/stainless-sdks/hubspot-sdk-go/option"
)

func TestFileAssetNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Files.FileAssets.New(context.TODO(), files.FileAssetNewParams{
		FolderInput: files.FolderInputParam{
			Name:           "name",
			ParentFolderID: hubspotsdk.String("parentFolderId"),
			ParentPath:     hubspotsdk.String("parentPath"),
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

func TestFileAssetUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Files.FileAssets.Update(
		context.TODO(),
		"321669910225",
		files.FileAssetUpdateParams{
			FileUpdateInput: files.FileUpdateInputParam{
				ClearExpires:      true,
				Access:            files.FileUpdateInputAccessHiddenIndexable,
				ExpiresAt:         hubspotsdk.Time(time.Now()),
				IsUsableInContent: hubspotsdk.Bool(true),
				Name:              hubspotsdk.String("name"),
				ParentFolderID:    hubspotsdk.String("parentFolderId"),
				ParentFolderPath:  hubspotsdk.String("parentFolderPath"),
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

func TestFileAssetDelete(t *testing.T) {
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
	err := client.Files.FileAssets.Delete(context.TODO(), "321669910225")
	if err != nil {
		var apierr *hubspotsdk.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestFileAssetGdprDelete(t *testing.T) {
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
	err := client.Files.FileAssets.GdprDelete(context.TODO(), "321669910225")
	if err != nil {
		var apierr *hubspotsdk.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestFileAssetGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Files.FileAssets.Get(
		context.TODO(),
		"321669910225",
		files.FileAssetGetParams{
			Properties: []string{"string"},
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

func TestFileAssetGetImportTaskStatus(t *testing.T) {
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
	_, err := client.Files.FileAssets.GetImportTaskStatus(context.TODO(), "taskId")
	if err != nil {
		var apierr *hubspotsdk.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestFileAssetGetSignedURLWithOptionalParams(t *testing.T) {
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
	_, err := client.Files.FileAssets.GetSignedURL(
		context.TODO(),
		"321669910225",
		files.FileAssetGetSignedURLParams{
			ExpirationSeconds: hubspotsdk.Int(0),
			Size:              files.FileAssetGetSignedURLParamsSizeIcon,
			Upscale:           hubspotsdk.Bool(true),
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

func TestFileAssetImportFromURLAsyncWithOptionalParams(t *testing.T) {
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
	_, err := client.Files.FileAssets.ImportFromURLAsync(context.TODO(), files.FileAssetImportFromURLAsyncParams{
		ImportFromURLInput: files.ImportFromURLInputParam{
			Access:                      files.ImportFromURLInputAccessHiddenIndexable,
			DuplicateValidationScope:    files.ImportFromURLInputDuplicateValidationScopeEntirePortal,
			DuplicateValidationStrategy: files.ImportFromURLInputDuplicateValidationStrategyNone,
			Overwrite:                   true,
			ExpiresAt:                   hubspotsdk.Time(time.Now()),
			FolderID:                    hubspotsdk.String("folderId"),
			FolderPath:                  hubspotsdk.String("folderPath"),
			Name:                        hubspotsdk.String("name"),
			Ttl:                         hubspotsdk.String("ttl"),
			URL:                         hubspotsdk.String("url"),
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

func TestFileAssetReplaceWithOptionalParams(t *testing.T) {
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
	_, err := client.Files.FileAssets.Replace(
		context.TODO(),
		"321669910225",
		files.FileAssetReplaceParams{
			CharsetHunch: hubspotsdk.String("charsetHunch"),
			File:         io.Reader(bytes.NewBuffer([]byte("Example data"))),
			Options:      hubspotsdk.String("options"),
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

func TestFileAssetSearchWithOptionalParams(t *testing.T) {
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
	_, err := client.Files.FileAssets.Search(context.TODO(), files.FileAssetSearchParams{
		After:                 hubspotsdk.String("after"),
		AllowsAnonymousAccess: hubspotsdk.Bool(true),
		Before:                hubspotsdk.String("before"),
		CreatedAt:             hubspotsdk.Time(time.Now()),
		CreatedAtGte:          hubspotsdk.Time(time.Now()),
		CreatedAtLte:          hubspotsdk.Time(time.Now()),
		Encoding:              hubspotsdk.String("encoding"),
		ExpiresAt:             hubspotsdk.Time(time.Now()),
		ExpiresAtGte:          hubspotsdk.Time(time.Now()),
		ExpiresAtLte:          hubspotsdk.Time(time.Now()),
		Extension:             hubspotsdk.String("extension"),
		FileMd5:               hubspotsdk.String("fileMd5"),
		Height:                hubspotsdk.Int(0),
		HeightGte:             hubspotsdk.Int(0),
		HeightLte:             hubspotsdk.Int(0),
		IDGte:                 hubspotsdk.Int(0),
		IDLte:                 hubspotsdk.Int(0),
		IDs:                   []int64{0},
		IsUsableInContent:     hubspotsdk.Bool(true),
		Limit:                 hubspotsdk.Int(0),
		Name:                  hubspotsdk.String("name"),
		ParentFolderIDs:       []int64{0},
		Path:                  hubspotsdk.String("path"),
		Properties:            []string{"string"},
		Size:                  hubspotsdk.Int(0),
		SizeGte:               hubspotsdk.Int(0),
		SizeLte:               hubspotsdk.Int(0),
		Sort:                  []string{"string"},
		Type:                  hubspotsdk.String("type"),
		UpdatedAt:             hubspotsdk.Time(time.Now()),
		UpdatedAtGte:          hubspotsdk.Time(time.Now()),
		UpdatedAtLte:          hubspotsdk.Time(time.Now()),
		URL:                   hubspotsdk.String("url"),
		Width:                 hubspotsdk.Int(0),
		WidthGte:              hubspotsdk.Int(0),
		WidthLte:              hubspotsdk.Int(0),
	})
	if err != nil {
		var apierr *hubspotsdk.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestFileAssetUploadWithOptionalParams(t *testing.T) {
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
	_, err := client.Files.FileAssets.Upload(context.TODO(), files.FileAssetUploadParams{
		CharsetHunch: hubspotsdk.String("charsetHunch"),
		File:         io.Reader(bytes.NewBuffer([]byte("Example data"))),
		FileName:     hubspotsdk.String("fileName"),
		FolderID:     hubspotsdk.String("folderId"),
		FolderPath:   hubspotsdk.String("folderPath"),
		Options:      hubspotsdk.String("options"),
	})
	if err != nil {
		var apierr *hubspotsdk.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
