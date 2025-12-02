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

func TestFileOperationUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Files.FileOperations.Update(
		context.TODO(),
		"321669910225",
		files.FileOperationUpdateParams{
			FileUpdateInput: files.FileUpdateInputParam{
				Access:            files.FileUpdateInputAccessHiddenIndexable,
				ClearExpires:      hubspotsdk.Bool(true),
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

func TestFileOperationDelete(t *testing.T) {
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
	err := client.Files.FileOperations.Delete(context.TODO(), "321669910225")
	if err != nil {
		var apierr *hubspotsdk.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestFileOperationGdprDelete(t *testing.T) {
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
	err := client.Files.FileOperations.GdprDelete(context.TODO(), "321669910225")
	if err != nil {
		var apierr *hubspotsdk.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestFileOperationGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Files.FileOperations.Get(
		context.TODO(),
		"321669910225",
		files.FileOperationGetParams{
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

func TestFileOperationGetByPathWithOptionalParams(t *testing.T) {
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
	_, err := client.Files.FileOperations.GetByPath(
		context.TODO(),
		"file_path",
		files.FileOperationGetByPathParams{
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

func TestFileOperationGetImportTaskStatus(t *testing.T) {
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
	_, err := client.Files.FileOperations.GetImportTaskStatus(context.TODO(), "taskId")
	if err != nil {
		var apierr *hubspotsdk.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestFileOperationGetSignedURLWithOptionalParams(t *testing.T) {
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
	_, err := client.Files.FileOperations.GetSignedURL(
		context.TODO(),
		"321669910225",
		files.FileOperationGetSignedURLParams{
			ExpirationSeconds: hubspotsdk.Int(0),
			Size:              files.FileOperationGetSignedURLParamsSizeIcon,
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

func TestFileOperationImportFromURLAsyncWithOptionalParams(t *testing.T) {
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
	_, err := client.Files.FileOperations.ImportFromURLAsync(context.TODO(), files.FileOperationImportFromURLAsyncParams{
		ImportFromURLInput: files.ImportFromURLInputParam{
			Access:                      files.ImportFromURLInputAccessHiddenIndexable,
			URL:                         "url",
			DuplicateValidationScope:    files.ImportFromURLInputDuplicateValidationScopeEntirePortal,
			DuplicateValidationStrategy: files.ImportFromURLInputDuplicateValidationStrategyNone,
			ExpiresAt:                   hubspotsdk.Time(time.Now()),
			FolderID:                    hubspotsdk.String("folderId"),
			FolderPath:                  hubspotsdk.String("folderPath"),
			Name:                        hubspotsdk.String("name"),
			Overwrite:                   hubspotsdk.Bool(true),
			Ttl:                         hubspotsdk.String("ttl"),
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

func TestFileOperationReplaceWithOptionalParams(t *testing.T) {
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
	_, err := client.Files.FileOperations.Replace(
		context.TODO(),
		"321669910225",
		files.FileOperationReplaceParams{
			CharsetHunch: hubspotsdk.String("charsetHunch"),
			File:         io.Reader(bytes.NewBuffer([]byte("some file contents"))),
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

func TestFileOperationSearchWithOptionalParams(t *testing.T) {
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
	_, err := client.Files.FileOperations.Search(context.TODO(), files.FileOperationSearchParams{
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

func TestFileOperationUploadWithOptionalParams(t *testing.T) {
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
	_, err := client.Files.FileOperations.Upload(context.TODO(), files.FileOperationUploadParams{
		CharsetHunch: hubspotsdk.String("charsetHunch"),
		File:         io.Reader(bytes.NewBuffer([]byte("some file contents"))),
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
