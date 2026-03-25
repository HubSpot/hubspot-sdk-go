// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package data_studio_test

import (
	"bytes"
	"context"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/stainless-sdks/hubspot-sdk-go"
	"github.com/stainless-sdks/hubspot-sdk-go/data_studio"
	"github.com/stainless-sdks/hubspot-sdk-go/internal/testutil"
	"github.com/stainless-sdks/hubspot-sdk-go/option"
)

func TestDatasourceNewWithOptionalParams(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte("abc"))
	}))
	defer server.Close()
	baseURL := server.URL
	client := hubspotsdk.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAccessToken("pat-na1-xxxxxxxx-xxxx"),
	)
	resp, err := client.DataStudio.Datasource.New(context.TODO(), data_studio.DatasourceNewParams{
		FormDataMultiPart: data_studio.FormDataMultiPartParam{
			BodyParts: []data_studio.BodyPartParam{{
				ContentDisposition: data_studio.ContentDispositionParam{
					CreationDate:     time.Now(),
					FileName:         "fileName",
					ModificationDate: time.Now(),
					Parameters: map[string]string{
						"foo": "string",
					},
					ReadDate: time.Now(),
					Size:     0,
					Type:     "type",
				},
				Entity: map[string]any{},
				Headers: map[string][]string{
					"foo": {"string"},
				},
				MediaType: data_studio.MediaTypeParam{
					Parameters: map[string]string{
						"foo": "string",
					},
					Subtype:         "subtype",
					Type:            "type",
					WildcardSubtype: true,
					WildcardType:    true,
				},
				MessageBodyWorkers: map[string]any{},
				ParameterizedHeaders: map[string][]data_studio.ParameterizedHeaderParam{
					"foo": {{
						Parameters: map[string]string{
							"foo": "string",
						},
						Value: "value",
					}},
				},
				Providers: map[string]any{},
				Parent: data_studio.MultiPartParam{
					BodyParts: []data_studio.BodyPartParam{},
					ContentDisposition: data_studio.ContentDispositionParam{
						CreationDate:     time.Now(),
						FileName:         "fileName",
						ModificationDate: time.Now(),
						Parameters: map[string]string{
							"foo": "string",
						},
						ReadDate: time.Now(),
						Size:     0,
						Type:     "type",
					},
					Entity: map[string]any{},
					Headers: map[string][]string{
						"foo": {"string"},
					},
					MediaType: data_studio.MediaTypeParam{
						Parameters: map[string]string{
							"foo": "string",
						},
						Subtype:         "subtype",
						Type:            "type",
						WildcardSubtype: true,
						WildcardType:    true,
					},
					MessageBodyWorkers: map[string]any{},
					ParameterizedHeaders: map[string][]data_studio.ParameterizedHeaderParam{
						"foo": {{
							Parameters: map[string]string{
								"foo": "string",
							},
							Value: "value",
						}},
					},
					Providers: map[string]any{},
				},
			}},
			ContentDisposition: data_studio.ContentDispositionParam{
				CreationDate:     time.Now(),
				FileName:         "fileName",
				ModificationDate: time.Now(),
				Parameters: map[string]string{
					"foo": "string",
				},
				ReadDate: time.Now(),
				Size:     0,
				Type:     "type",
			},
			Entity: map[string]any{},
			Fields: map[string][]data_studio.FormDataBodyPartParam{
				"foo": {{
					ContentDisposition: data_studio.ContentDispositionParam{
						CreationDate:     time.Now(),
						FileName:         "fileName",
						ModificationDate: time.Now(),
						Parameters: map[string]string{
							"foo": "string",
						},
						ReadDate: time.Now(),
						Size:     0,
						Type:     "type",
					},
					Entity: map[string]any{},
					FormDataContentDisposition: data_studio.FormDataContentDispositionParam{
						CreationDate:     time.Now(),
						FileName:         "fileName",
						ModificationDate: time.Now(),
						Name:             "name",
						Parameters: map[string]string{
							"foo": "string",
						},
						ReadDate: time.Now(),
						Size:     0,
						Type:     "type",
					},
					Headers: map[string][]string{
						"foo": {"string"},
					},
					MediaType: data_studio.MediaTypeParam{
						Parameters: map[string]string{
							"foo": "string",
						},
						Subtype:         "subtype",
						Type:            "type",
						WildcardSubtype: true,
						WildcardType:    true,
					},
					MessageBodyWorkers: map[string]any{},
					Name:               "name",
					ParameterizedHeaders: map[string][]data_studio.ParameterizedHeaderParam{
						"foo": {{
							Parameters: map[string]string{
								"foo": "string",
							},
							Value: "value",
						}},
					},
					Providers: map[string]any{},
					Simple:    true,
					Value:     "value",
					Parent: data_studio.MultiPartParam{
						BodyParts: []data_studio.BodyPartParam{},
						ContentDisposition: data_studio.ContentDispositionParam{
							CreationDate:     time.Now(),
							FileName:         "fileName",
							ModificationDate: time.Now(),
							Parameters: map[string]string{
								"foo": "string",
							},
							ReadDate: time.Now(),
							Size:     0,
							Type:     "type",
						},
						Entity: map[string]any{},
						Headers: map[string][]string{
							"foo": {"string"},
						},
						MediaType: data_studio.MediaTypeParam{
							Parameters: map[string]string{
								"foo": "string",
							},
							Subtype:         "subtype",
							Type:            "type",
							WildcardSubtype: true,
							WildcardType:    true,
						},
						MessageBodyWorkers: map[string]any{},
						ParameterizedHeaders: map[string][]data_studio.ParameterizedHeaderParam{
							"foo": {{
								Parameters: map[string]string{
									"foo": "string",
								},
								Value: "value",
							}},
						},
						Providers: map[string]any{},
					},
				}},
			},
			Headers: map[string][]string{
				"foo": {"string"},
			},
			MediaType: data_studio.MediaTypeParam{
				Parameters: map[string]string{
					"foo": "string",
				},
				Subtype:         "subtype",
				Type:            "type",
				WildcardSubtype: true,
				WildcardType:    true,
			},
			MessageBodyWorkers: map[string]any{},
			ParameterizedHeaders: map[string][]data_studio.ParameterizedHeaderParam{
				"foo": {{
					Parameters: map[string]string{
						"foo": "string",
					},
					Value: "value",
				}},
			},
			Providers: map[string]any{},
			Parent: data_studio.MultiPartParam{
				BodyParts: []data_studio.BodyPartParam{},
				ContentDisposition: data_studio.ContentDispositionParam{
					CreationDate:     time.Now(),
					FileName:         "fileName",
					ModificationDate: time.Now(),
					Parameters: map[string]string{
						"foo": "string",
					},
					ReadDate: time.Now(),
					Size:     0,
					Type:     "type",
				},
				Entity: map[string]any{},
				Headers: map[string][]string{
					"foo": {"string"},
				},
				MediaType: data_studio.MediaTypeParam{
					Parameters: map[string]string{
						"foo": "string",
					},
					Subtype:         "subtype",
					Type:            "type",
					WildcardSubtype: true,
					WildcardType:    true,
				},
				MessageBodyWorkers: map[string]any{},
				ParameterizedHeaders: map[string][]data_studio.ParameterizedHeaderParam{
					"foo": {{
						Parameters: map[string]string{
							"foo": "string",
						},
						Value: "value",
					}},
				},
				Providers: map[string]any{},
			},
		},
	})
	if err != nil {
		var apierr *hubspotsdk.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		var apierr *hubspotsdk.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
	if !bytes.Equal(b, []byte("abc")) {
		t.Fatalf("return value not %s: %s", "abc", b)
	}
}

func TestDatasourceUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.DataStudio.Datasource.Update(
		context.TODO(),
		0,
		data_studio.DatasourceUpdateParams{
			FormDataMultiPart: data_studio.FormDataMultiPartParam{
				BodyParts: []data_studio.BodyPartParam{{
					ContentDisposition: data_studio.ContentDispositionParam{
						CreationDate:     time.Now(),
						FileName:         "fileName",
						ModificationDate: time.Now(),
						Parameters: map[string]string{
							"foo": "string",
						},
						ReadDate: time.Now(),
						Size:     0,
						Type:     "type",
					},
					Entity: map[string]any{},
					Headers: map[string][]string{
						"foo": {"string"},
					},
					MediaType: data_studio.MediaTypeParam{
						Parameters: map[string]string{
							"foo": "string",
						},
						Subtype:         "subtype",
						Type:            "type",
						WildcardSubtype: true,
						WildcardType:    true,
					},
					MessageBodyWorkers: map[string]any{},
					ParameterizedHeaders: map[string][]data_studio.ParameterizedHeaderParam{
						"foo": {{
							Parameters: map[string]string{
								"foo": "string",
							},
							Value: "value",
						}},
					},
					Providers: map[string]any{},
					Parent: data_studio.MultiPartParam{
						BodyParts: []data_studio.BodyPartParam{},
						ContentDisposition: data_studio.ContentDispositionParam{
							CreationDate:     time.Now(),
							FileName:         "fileName",
							ModificationDate: time.Now(),
							Parameters: map[string]string{
								"foo": "string",
							},
							ReadDate: time.Now(),
							Size:     0,
							Type:     "type",
						},
						Entity: map[string]any{},
						Headers: map[string][]string{
							"foo": {"string"},
						},
						MediaType: data_studio.MediaTypeParam{
							Parameters: map[string]string{
								"foo": "string",
							},
							Subtype:         "subtype",
							Type:            "type",
							WildcardSubtype: true,
							WildcardType:    true,
						},
						MessageBodyWorkers: map[string]any{},
						ParameterizedHeaders: map[string][]data_studio.ParameterizedHeaderParam{
							"foo": {{
								Parameters: map[string]string{
									"foo": "string",
								},
								Value: "value",
							}},
						},
						Providers: map[string]any{},
					},
				}},
				ContentDisposition: data_studio.ContentDispositionParam{
					CreationDate:     time.Now(),
					FileName:         "fileName",
					ModificationDate: time.Now(),
					Parameters: map[string]string{
						"foo": "string",
					},
					ReadDate: time.Now(),
					Size:     0,
					Type:     "type",
				},
				Entity: map[string]any{},
				Fields: map[string][]data_studio.FormDataBodyPartParam{
					"foo": {{
						ContentDisposition: data_studio.ContentDispositionParam{
							CreationDate:     time.Now(),
							FileName:         "fileName",
							ModificationDate: time.Now(),
							Parameters: map[string]string{
								"foo": "string",
							},
							ReadDate: time.Now(),
							Size:     0,
							Type:     "type",
						},
						Entity: map[string]any{},
						FormDataContentDisposition: data_studio.FormDataContentDispositionParam{
							CreationDate:     time.Now(),
							FileName:         "fileName",
							ModificationDate: time.Now(),
							Name:             "name",
							Parameters: map[string]string{
								"foo": "string",
							},
							ReadDate: time.Now(),
							Size:     0,
							Type:     "type",
						},
						Headers: map[string][]string{
							"foo": {"string"},
						},
						MediaType: data_studio.MediaTypeParam{
							Parameters: map[string]string{
								"foo": "string",
							},
							Subtype:         "subtype",
							Type:            "type",
							WildcardSubtype: true,
							WildcardType:    true,
						},
						MessageBodyWorkers: map[string]any{},
						Name:               "name",
						ParameterizedHeaders: map[string][]data_studio.ParameterizedHeaderParam{
							"foo": {{
								Parameters: map[string]string{
									"foo": "string",
								},
								Value: "value",
							}},
						},
						Providers: map[string]any{},
						Simple:    true,
						Value:     "value",
						Parent: data_studio.MultiPartParam{
							BodyParts: []data_studio.BodyPartParam{},
							ContentDisposition: data_studio.ContentDispositionParam{
								CreationDate:     time.Now(),
								FileName:         "fileName",
								ModificationDate: time.Now(),
								Parameters: map[string]string{
									"foo": "string",
								},
								ReadDate: time.Now(),
								Size:     0,
								Type:     "type",
							},
							Entity: map[string]any{},
							Headers: map[string][]string{
								"foo": {"string"},
							},
							MediaType: data_studio.MediaTypeParam{
								Parameters: map[string]string{
									"foo": "string",
								},
								Subtype:         "subtype",
								Type:            "type",
								WildcardSubtype: true,
								WildcardType:    true,
							},
							MessageBodyWorkers: map[string]any{},
							ParameterizedHeaders: map[string][]data_studio.ParameterizedHeaderParam{
								"foo": {{
									Parameters: map[string]string{
										"foo": "string",
									},
									Value: "value",
								}},
							},
							Providers: map[string]any{},
						},
					}},
				},
				Headers: map[string][]string{
					"foo": {"string"},
				},
				MediaType: data_studio.MediaTypeParam{
					Parameters: map[string]string{
						"foo": "string",
					},
					Subtype:         "subtype",
					Type:            "type",
					WildcardSubtype: true,
					WildcardType:    true,
				},
				MessageBodyWorkers: map[string]any{},
				ParameterizedHeaders: map[string][]data_studio.ParameterizedHeaderParam{
					"foo": {{
						Parameters: map[string]string{
							"foo": "string",
						},
						Value: "value",
					}},
				},
				Providers: map[string]any{},
				Parent: data_studio.MultiPartParam{
					BodyParts: []data_studio.BodyPartParam{},
					ContentDisposition: data_studio.ContentDispositionParam{
						CreationDate:     time.Now(),
						FileName:         "fileName",
						ModificationDate: time.Now(),
						Parameters: map[string]string{
							"foo": "string",
						},
						ReadDate: time.Now(),
						Size:     0,
						Type:     "type",
					},
					Entity: map[string]any{},
					Headers: map[string][]string{
						"foo": {"string"},
					},
					MediaType: data_studio.MediaTypeParam{
						Parameters: map[string]string{
							"foo": "string",
						},
						Subtype:         "subtype",
						Type:            "type",
						WildcardSubtype: true,
						WildcardType:    true,
					},
					MessageBodyWorkers: map[string]any{},
					ParameterizedHeaders: map[string][]data_studio.ParameterizedHeaderParam{
						"foo": {{
							Parameters: map[string]string{
								"foo": "string",
							},
							Value: "value",
						}},
					},
					Providers: map[string]any{},
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

func TestDatasourceDelete(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte("abc"))
	}))
	defer server.Close()
	baseURL := server.URL
	client := hubspotsdk.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAccessToken("pat-na1-xxxxxxxx-xxxx"),
	)
	resp, err := client.DataStudio.Datasource.Delete(context.TODO(), 0)
	if err != nil {
		var apierr *hubspotsdk.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		var apierr *hubspotsdk.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
	if !bytes.Equal(b, []byte("abc")) {
		t.Fatalf("return value not %s: %s", "abc", b)
	}
}

func TestDatasourceGet(t *testing.T) {
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
	_, err := client.DataStudio.Datasource.Get(context.TODO(), 0)
	if err != nil {
		var apierr *hubspotsdk.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
