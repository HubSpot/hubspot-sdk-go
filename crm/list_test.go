// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package crm_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/HubSpot/hubspot-sdk-go"
	"github.com/HubSpot/hubspot-sdk-go/crm"
	"github.com/HubSpot/hubspot-sdk-go/internal/testutil"
	"github.com/HubSpot/hubspot-sdk-go/option"
)

func TestListNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Crm.Lists.New(context.TODO(), crm.ListNewParams{
		ListCreateRequest: crm.ListCreateRequestParam{
			Name:           "name",
			ObjectTypeID:   "objectTypeId",
			ProcessingType: "processingType",
			CustomProperties: map[string]string{
				"foo": "string",
			},
			FilterBranch: crm.ListCreateRequestFilterBranchUnionParam{
				OfOr: &crm.PublicOrFilterBranchParam{
					FilterBranches: []crm.PublicOrFilterBranchFilterBranchUnionParam{{
						OfAnd: &crm.PublicAndFilterBranchParam{
							FilterBranches: []crm.PublicAndFilterBranchFilterBranchUnionParam{{
								OfNotAll: &crm.PublicNotAllFilterBranchParam{
									FilterBranches: []crm.PublicNotAllFilterBranchFilterBranchUnionParam{{
										OfNotAny: &crm.PublicNotAnyFilterBranchParam{
											FilterBranches: []crm.PublicNotAnyFilterBranchFilterBranchUnionParam{{
												OfRestricted: &crm.PublicRestrictedFilterBranchParam{
													FilterBranches: []crm.PublicRestrictedFilterBranchFilterBranchUnionParam{{
														OfUnifiedEvents: &crm.PublicUnifiedEventsFilterBranchParam{
															EventTypeID: "eventTypeId",
															FilterBranches: []crm.PublicUnifiedEventsFilterBranchFilterBranchUnionParam{{
																OfAssociation: &crm.PublicAssociationFilterBranchParam{
																	AssociationCategory: "associationCategory",
																	AssociationTypeID:   0,
																	FilterBranches: []crm.PublicAssociationFilterBranchFilterBranchUnionParam{{
																		OfOr: &crm.PublicOrFilterBranchParam{
																			FilterBranches:       []crm.PublicOrFilterBranchFilterBranchUnionParam{},
																			FilterBranchOperator: "filterBranchOperator",
																			FilterBranchType:     crm.PublicOrFilterBranchFilterBranchTypeOr,
																			Filters: []crm.PublicOrFilterBranchFilterUnionParam{{
																				OfProperty: &crm.PublicPropertyFilterParam{
																					FilterType: crm.PublicPropertyFilterFilterTypeProperty,
																					Operation: crm.PublicPropertyFilterOperationUnionParam{
																						OfBool: &crm.PublicBoolPropertyOperationParam{
																							IncludeObjectsWithNoValueSet: true,
																							OperationType:                crm.PublicBoolPropertyOperationOperationTypeBool,
																							Operator:                     "operator",
																							Value:                        true,
																						},
																					},
																					Property: "property",
																				},
																			}},
																		},
																	}},
																	FilterBranchOperator: "filterBranchOperator",
																	FilterBranchType:     crm.PublicAssociationFilterBranchFilterBranchTypeAssociation,
																	Filters: []crm.PublicAssociationFilterBranchFilterUnionParam{{
																		OfProperty: &crm.PublicPropertyFilterParam{
																			FilterType: crm.PublicPropertyFilterFilterTypeProperty,
																			Operation: crm.PublicPropertyFilterOperationUnionParam{
																				OfBool: &crm.PublicBoolPropertyOperationParam{
																					IncludeObjectsWithNoValueSet: true,
																					OperationType:                crm.PublicBoolPropertyOperationOperationTypeBool,
																					Operator:                     "operator",
																					Value:                        true,
																				},
																			},
																			Property: "property",
																		},
																	}},
																	ObjectTypeID: "objectTypeId",
																	Operator:     "operator",
																},
															}},
															FilterBranchOperator: "filterBranchOperator",
															FilterBranchType:     crm.PublicUnifiedEventsFilterBranchFilterBranchTypeUnifiedEvents,
															Filters: []crm.PublicUnifiedEventsFilterBranchFilterUnionParam{{
																OfProperty: &crm.PublicPropertyFilterParam{
																	FilterType: crm.PublicPropertyFilterFilterTypeProperty,
																	Operation: crm.PublicPropertyFilterOperationUnionParam{
																		OfBool: &crm.PublicBoolPropertyOperationParam{
																			IncludeObjectsWithNoValueSet: true,
																			OperationType:                crm.PublicBoolPropertyOperationOperationTypeBool,
																			Operator:                     "operator",
																			Value:                        true,
																		},
																	},
																	Property: "property",
																},
															}},
															Operator: crm.PublicUnifiedEventsFilterBranchOperatorHasCompleted,
															CoalescingRefineBy: crm.PublicUnifiedEventsFilterBranchCoalescingRefineByUnionParam{
																OfNumOccurrences: &crm.PublicNumOccurrencesRefineByParam{
																	Type:           crm.PublicNumOccurrencesRefineByTypeNumOccurrences,
																	MaxOccurrences: hubspotsdk.Int(0),
																	MinOccurrences: hubspotsdk.Int(0),
																},
															},
															PruningRefineBy: crm.PublicUnifiedEventsFilterBranchPruningRefineByUnionParam{
																OfNumOccurrences: &crm.PublicNumOccurrencesRefineByParam{
																	Type:           crm.PublicNumOccurrencesRefineByTypeNumOccurrences,
																	MaxOccurrences: hubspotsdk.Int(0),
																	MinOccurrences: hubspotsdk.Int(0),
																},
															},
														},
													}},
													FilterBranchOperator: "filterBranchOperator",
													FilterBranchType:     crm.PublicRestrictedFilterBranchFilterBranchTypeRestricted,
													Filters: []crm.PublicRestrictedFilterBranchFilterUnionParam{{
														OfProperty: &crm.PublicPropertyFilterParam{
															FilterType: crm.PublicPropertyFilterFilterTypeProperty,
															Operation: crm.PublicPropertyFilterOperationUnionParam{
																OfBool: &crm.PublicBoolPropertyOperationParam{
																	IncludeObjectsWithNoValueSet: true,
																	OperationType:                crm.PublicBoolPropertyOperationOperationTypeBool,
																	Operator:                     "operator",
																	Value:                        true,
																},
															},
															Property: "property",
														},
													}},
												},
											}},
											FilterBranchOperator: "filterBranchOperator",
											FilterBranchType:     crm.PublicNotAnyFilterBranchFilterBranchTypeNotAny,
											Filters: []crm.PublicNotAnyFilterBranchFilterUnionParam{{
												OfProperty: &crm.PublicPropertyFilterParam{
													FilterType: crm.PublicPropertyFilterFilterTypeProperty,
													Operation: crm.PublicPropertyFilterOperationUnionParam{
														OfBool: &crm.PublicBoolPropertyOperationParam{
															IncludeObjectsWithNoValueSet: true,
															OperationType:                crm.PublicBoolPropertyOperationOperationTypeBool,
															Operator:                     "operator",
															Value:                        true,
														},
													},
													Property: "property",
												},
											}},
										},
									}},
									FilterBranchOperator: "filterBranchOperator",
									FilterBranchType:     crm.PublicNotAllFilterBranchFilterBranchTypeNotAll,
									Filters: []crm.PublicNotAllFilterBranchFilterUnionParam{{
										OfProperty: &crm.PublicPropertyFilterParam{
											FilterType: crm.PublicPropertyFilterFilterTypeProperty,
											Operation: crm.PublicPropertyFilterOperationUnionParam{
												OfBool: &crm.PublicBoolPropertyOperationParam{
													IncludeObjectsWithNoValueSet: true,
													OperationType:                crm.PublicBoolPropertyOperationOperationTypeBool,
													Operator:                     "operator",
													Value:                        true,
												},
											},
											Property: "property",
										},
									}},
								},
							}},
							FilterBranchOperator: "filterBranchOperator",
							FilterBranchType:     crm.PublicAndFilterBranchFilterBranchTypeAnd,
							Filters: []crm.PublicAndFilterBranchFilterUnionParam{{
								OfProperty: &crm.PublicPropertyFilterParam{
									FilterType: crm.PublicPropertyFilterFilterTypeProperty,
									Operation: crm.PublicPropertyFilterOperationUnionParam{
										OfBool: &crm.PublicBoolPropertyOperationParam{
											IncludeObjectsWithNoValueSet: true,
											OperationType:                crm.PublicBoolPropertyOperationOperationTypeBool,
											Operator:                     "operator",
											Value:                        true,
										},
									},
									Property: "property",
								},
							}},
						},
					}},
					FilterBranchOperator: "filterBranchOperator",
					FilterBranchType:     crm.PublicOrFilterBranchFilterBranchTypeOr,
					Filters: []crm.PublicOrFilterBranchFilterUnionParam{{
						OfProperty: &crm.PublicPropertyFilterParam{
							FilterType: crm.PublicPropertyFilterFilterTypeProperty,
							Operation: crm.PublicPropertyFilterOperationUnionParam{
								OfBool: &crm.PublicBoolPropertyOperationParam{
									IncludeObjectsWithNoValueSet: true,
									OperationType:                crm.PublicBoolPropertyOperationOperationTypeBool,
									Operator:                     "operator",
									Value:                        true,
								},
							},
							Property: "property",
						},
					}},
				},
			},
			ListFolderID: hubspotsdk.Int(0),
			ListPermissions: crm.PublicListPermissionsParam{
				TeamsWithEditAccess: []int64{0},
				UsersWithEditAccess: []int64{0},
			},
			MembershipSettings: crm.PublicMembershipSettingsParam{
				IncludeUnassigned: hubspotsdk.Bool(true),
				MembershipTeamID:  hubspotsdk.Int(0),
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
}

func TestListListWithOptionalParams(t *testing.T) {
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
	_, err := client.Crm.Lists.List(context.TODO(), crm.ListListParams{
		IncludeFilters: hubspotsdk.Bool(true),
		ListIDs:        []string{"string"},
	})
	if err != nil {
		var apierr *hubspotsdk.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestListDelete(t *testing.T) {
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
	err := client.Crm.Lists.Delete(context.TODO(), "listId")
	if err != nil {
		var apierr *hubspotsdk.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestListAddAndRemoveMemberships(t *testing.T) {
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
	_, err := client.Crm.Lists.AddAndRemoveMemberships(
		context.TODO(),
		"listId",
		crm.ListAddAndRemoveMembershipsParams{
			MembershipChangeRequest: crm.MembershipChangeRequestParam{
				RecordIDsToAdd:    []string{"string"},
				RecordIDsToRemove: []string{"string"},
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

func TestListAddMemberships(t *testing.T) {
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
	_, err := client.Crm.Lists.AddMemberships(
		context.TODO(),
		"listId",
		crm.ListAddMembershipsParams{
			Body: []string{"string"},
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

func TestListAddMembershipsFrom(t *testing.T) {
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
	err := client.Crm.Lists.AddMembershipsFrom(
		context.TODO(),
		"sourceListId",
		crm.ListAddMembershipsFromParams{
			ListID: "listId",
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

func TestListBatchReadMemberships(t *testing.T) {
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
	_, err := client.Crm.Lists.BatchReadMemberships(context.TODO(), crm.ListBatchReadMembershipsParams{
		BatchInputRecordIDInput: crm.BatchInputRecordIDInputParam{
			Inputs: []crm.RecordIDInputParam{{
				ObjectTypeID: "objectTypeId",
				RecordID:     "recordId",
			}},
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

func TestListNewFolderWithOptionalParams(t *testing.T) {
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
	_, err := client.Crm.Lists.NewFolder(context.TODO(), crm.ListNewFolderParams{
		ListFolderCreateRequest: crm.ListFolderCreateRequestParam{
			Name:           "name",
			ParentFolderID: hubspotsdk.String("parentFolderId"),
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

func TestListNewIDMapping(t *testing.T) {
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
	_, err := client.Crm.Lists.NewIDMapping(context.TODO(), crm.ListNewIDMappingParams{
		Body: []string{"string"},
	})
	if err != nil {
		var apierr *hubspotsdk.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestListDeleteFolder(t *testing.T) {
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
	err := client.Crm.Lists.DeleteFolder(context.TODO(), "folderId")
	if err != nil {
		var apierr *hubspotsdk.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestListDeleteMemberships(t *testing.T) {
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
	err := client.Crm.Lists.DeleteMemberships(context.TODO(), "listId")
	if err != nil {
		var apierr *hubspotsdk.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestListGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Crm.Lists.Get(
		context.TODO(),
		"listId",
		crm.ListGetParams{
			IncludeFilters: hubspotsdk.Bool(true),
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

func TestListGetByObjectTypeAndNameWithOptionalParams(t *testing.T) {
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
	_, err := client.Crm.Lists.GetByObjectTypeAndName(
		context.TODO(),
		"listName",
		crm.ListGetByObjectTypeAndNameParams{
			ObjectTypeID:   "objectTypeId",
			IncludeFilters: hubspotsdk.Bool(true),
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

func TestListGetIDMappingWithOptionalParams(t *testing.T) {
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
	_, err := client.Crm.Lists.GetIDMapping(context.TODO(), crm.ListGetIDMappingParams{
		LegacyListID: hubspotsdk.String("legacyListId"),
	})
	if err != nil {
		var apierr *hubspotsdk.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestListGetMembershipsJoinOrderWithOptionalParams(t *testing.T) {
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
	_, err := client.Crm.Lists.GetMembershipsJoinOrder(
		context.TODO(),
		"listId",
		crm.ListGetMembershipsJoinOrderParams{
			After:  hubspotsdk.String("after"),
			Before: hubspotsdk.String("before"),
			Limit:  hubspotsdk.Int(0),
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

func TestListGetRecordMemberships(t *testing.T) {
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
	_, err := client.Crm.Lists.GetRecordMemberships(
		context.TODO(),
		"recordId",
		crm.ListGetRecordMembershipsParams{
			ObjectTypeID: "objectTypeId",
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

func TestListGetScheduleConversion(t *testing.T) {
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
	_, err := client.Crm.Lists.GetScheduleConversion(context.TODO(), "listId")
	if err != nil {
		var apierr *hubspotsdk.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestListGetSizeAndEditsHistoryBetweenWithOptionalParams(t *testing.T) {
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
	_, err := client.Crm.Lists.GetSizeAndEditsHistoryBetween(
		context.TODO(),
		"listId",
		crm.ListGetSizeAndEditsHistoryBetweenParams{
			EndDate:   hubspotsdk.Time(time.Now()),
			StartDate: hubspotsdk.Time(time.Now()),
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

func TestListListBySearchWithOptionalParams(t *testing.T) {
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
	_, err := client.Crm.Lists.ListBySearch(context.TODO(), crm.ListListBySearchParams{
		ListSearchRequest: crm.ListSearchRequestParam{
			ListIDs:                    []string{"string"},
			Offset:                     0,
			ProcessingTypes:            []string{"string"},
			AdditionalFilterProperties: []string{"string"},
			Count:                      hubspotsdk.Int(0),
			ObjectTypeID:               hubspotsdk.String("objectTypeId"),
			Query:                      hubspotsdk.String("query"),
			Sort:                       hubspotsdk.String("sort"),
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

func TestListListFoldersWithOptionalParams(t *testing.T) {
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
	_, err := client.Crm.Lists.ListFolders(context.TODO(), crm.ListListFoldersParams{
		FolderID: hubspotsdk.String("folderId"),
	})
	if err != nil {
		var apierr *hubspotsdk.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestListListMembershipsWithOptionalParams(t *testing.T) {
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
	_, err := client.Crm.Lists.ListMemberships(
		context.TODO(),
		"listId",
		crm.ListListMembershipsParams{
			After:  hubspotsdk.String("after"),
			Before: hubspotsdk.String("before"),
			Limit:  hubspotsdk.Int(0),
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

func TestListMoveFolder(t *testing.T) {
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
	_, err := client.Crm.Lists.MoveFolder(
		context.TODO(),
		"newParentFolderId",
		crm.ListMoveFolderParams{
			FolderID: "folderId",
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

func TestListMoveList(t *testing.T) {
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
	err := client.Crm.Lists.MoveList(context.TODO(), crm.ListMoveListParams{
		ListMoveRequest: crm.ListMoveRequestParam{
			ListID:      "listId",
			NewFolderID: "newFolderId",
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

func TestListRemoveMemberships(t *testing.T) {
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
	_, err := client.Crm.Lists.RemoveMemberships(
		context.TODO(),
		"listId",
		crm.ListRemoveMembershipsParams{
			Body: []string{"string"},
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

func TestListRenameFolderWithOptionalParams(t *testing.T) {
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
	_, err := client.Crm.Lists.RenameFolder(
		context.TODO(),
		"folderId",
		crm.ListRenameFolderParams{
			NewFolderName: hubspotsdk.String("newFolderName"),
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

func TestListRestore(t *testing.T) {
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
	err := client.Crm.Lists.Restore(context.TODO(), "listId")
	if err != nil {
		var apierr *hubspotsdk.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestListScheduleConversion(t *testing.T) {
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
	err := client.Crm.Lists.ScheduleConversion(context.TODO(), "listId")
	if err != nil {
		var apierr *hubspotsdk.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestListUpdateListFiltersWithOptionalParams(t *testing.T) {
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
	_, err := client.Crm.Lists.UpdateListFilters(
		context.TODO(),
		"listId",
		crm.ListUpdateListFiltersParams{
			ListFilterUpdateRequest: crm.ListFilterUpdateRequestParam{
				FilterBranch: crm.ListFilterUpdateRequestFilterBranchUnionParam{
					OfOr: &crm.PublicOrFilterBranchParam{
						FilterBranches: []crm.PublicOrFilterBranchFilterBranchUnionParam{{
							OfAnd: &crm.PublicAndFilterBranchParam{
								FilterBranches: []crm.PublicAndFilterBranchFilterBranchUnionParam{{
									OfNotAll: &crm.PublicNotAllFilterBranchParam{
										FilterBranches: []crm.PublicNotAllFilterBranchFilterBranchUnionParam{{
											OfNotAny: &crm.PublicNotAnyFilterBranchParam{
												FilterBranches: []crm.PublicNotAnyFilterBranchFilterBranchUnionParam{{
													OfRestricted: &crm.PublicRestrictedFilterBranchParam{
														FilterBranches: []crm.PublicRestrictedFilterBranchFilterBranchUnionParam{{
															OfUnifiedEvents: &crm.PublicUnifiedEventsFilterBranchParam{
																EventTypeID: "eventTypeId",
																FilterBranches: []crm.PublicUnifiedEventsFilterBranchFilterBranchUnionParam{{
																	OfAssociation: &crm.PublicAssociationFilterBranchParam{
																		AssociationCategory: "associationCategory",
																		AssociationTypeID:   0,
																		FilterBranches: []crm.PublicAssociationFilterBranchFilterBranchUnionParam{{
																			OfOr: &crm.PublicOrFilterBranchParam{
																				FilterBranches:       []crm.PublicOrFilterBranchFilterBranchUnionParam{},
																				FilterBranchOperator: "filterBranchOperator",
																				FilterBranchType:     crm.PublicOrFilterBranchFilterBranchTypeOr,
																				Filters: []crm.PublicOrFilterBranchFilterUnionParam{{
																					OfProperty: &crm.PublicPropertyFilterParam{
																						FilterType: crm.PublicPropertyFilterFilterTypeProperty,
																						Operation: crm.PublicPropertyFilterOperationUnionParam{
																							OfBool: &crm.PublicBoolPropertyOperationParam{
																								IncludeObjectsWithNoValueSet: true,
																								OperationType:                crm.PublicBoolPropertyOperationOperationTypeBool,
																								Operator:                     "operator",
																								Value:                        true,
																							},
																						},
																						Property: "property",
																					},
																				}},
																			},
																		}},
																		FilterBranchOperator: "filterBranchOperator",
																		FilterBranchType:     crm.PublicAssociationFilterBranchFilterBranchTypeAssociation,
																		Filters: []crm.PublicAssociationFilterBranchFilterUnionParam{{
																			OfProperty: &crm.PublicPropertyFilterParam{
																				FilterType: crm.PublicPropertyFilterFilterTypeProperty,
																				Operation: crm.PublicPropertyFilterOperationUnionParam{
																					OfBool: &crm.PublicBoolPropertyOperationParam{
																						IncludeObjectsWithNoValueSet: true,
																						OperationType:                crm.PublicBoolPropertyOperationOperationTypeBool,
																						Operator:                     "operator",
																						Value:                        true,
																					},
																				},
																				Property: "property",
																			},
																		}},
																		ObjectTypeID: "objectTypeId",
																		Operator:     "operator",
																	},
																}},
																FilterBranchOperator: "filterBranchOperator",
																FilterBranchType:     crm.PublicUnifiedEventsFilterBranchFilterBranchTypeUnifiedEvents,
																Filters: []crm.PublicUnifiedEventsFilterBranchFilterUnionParam{{
																	OfProperty: &crm.PublicPropertyFilterParam{
																		FilterType: crm.PublicPropertyFilterFilterTypeProperty,
																		Operation: crm.PublicPropertyFilterOperationUnionParam{
																			OfBool: &crm.PublicBoolPropertyOperationParam{
																				IncludeObjectsWithNoValueSet: true,
																				OperationType:                crm.PublicBoolPropertyOperationOperationTypeBool,
																				Operator:                     "operator",
																				Value:                        true,
																			},
																		},
																		Property: "property",
																	},
																}},
																Operator: crm.PublicUnifiedEventsFilterBranchOperatorHasCompleted,
																CoalescingRefineBy: crm.PublicUnifiedEventsFilterBranchCoalescingRefineByUnionParam{
																	OfNumOccurrences: &crm.PublicNumOccurrencesRefineByParam{
																		Type:           crm.PublicNumOccurrencesRefineByTypeNumOccurrences,
																		MaxOccurrences: hubspotsdk.Int(0),
																		MinOccurrences: hubspotsdk.Int(0),
																	},
																},
																PruningRefineBy: crm.PublicUnifiedEventsFilterBranchPruningRefineByUnionParam{
																	OfNumOccurrences: &crm.PublicNumOccurrencesRefineByParam{
																		Type:           crm.PublicNumOccurrencesRefineByTypeNumOccurrences,
																		MaxOccurrences: hubspotsdk.Int(0),
																		MinOccurrences: hubspotsdk.Int(0),
																	},
																},
															},
														}},
														FilterBranchOperator: "filterBranchOperator",
														FilterBranchType:     crm.PublicRestrictedFilterBranchFilterBranchTypeRestricted,
														Filters: []crm.PublicRestrictedFilterBranchFilterUnionParam{{
															OfProperty: &crm.PublicPropertyFilterParam{
																FilterType: crm.PublicPropertyFilterFilterTypeProperty,
																Operation: crm.PublicPropertyFilterOperationUnionParam{
																	OfBool: &crm.PublicBoolPropertyOperationParam{
																		IncludeObjectsWithNoValueSet: true,
																		OperationType:                crm.PublicBoolPropertyOperationOperationTypeBool,
																		Operator:                     "operator",
																		Value:                        true,
																	},
																},
																Property: "property",
															},
														}},
													},
												}},
												FilterBranchOperator: "filterBranchOperator",
												FilterBranchType:     crm.PublicNotAnyFilterBranchFilterBranchTypeNotAny,
												Filters: []crm.PublicNotAnyFilterBranchFilterUnionParam{{
													OfProperty: &crm.PublicPropertyFilterParam{
														FilterType: crm.PublicPropertyFilterFilterTypeProperty,
														Operation: crm.PublicPropertyFilterOperationUnionParam{
															OfBool: &crm.PublicBoolPropertyOperationParam{
																IncludeObjectsWithNoValueSet: true,
																OperationType:                crm.PublicBoolPropertyOperationOperationTypeBool,
																Operator:                     "operator",
																Value:                        true,
															},
														},
														Property: "property",
													},
												}},
											},
										}},
										FilterBranchOperator: "filterBranchOperator",
										FilterBranchType:     crm.PublicNotAllFilterBranchFilterBranchTypeNotAll,
										Filters: []crm.PublicNotAllFilterBranchFilterUnionParam{{
											OfProperty: &crm.PublicPropertyFilterParam{
												FilterType: crm.PublicPropertyFilterFilterTypeProperty,
												Operation: crm.PublicPropertyFilterOperationUnionParam{
													OfBool: &crm.PublicBoolPropertyOperationParam{
														IncludeObjectsWithNoValueSet: true,
														OperationType:                crm.PublicBoolPropertyOperationOperationTypeBool,
														Operator:                     "operator",
														Value:                        true,
													},
												},
												Property: "property",
											},
										}},
									},
								}},
								FilterBranchOperator: "filterBranchOperator",
								FilterBranchType:     crm.PublicAndFilterBranchFilterBranchTypeAnd,
								Filters: []crm.PublicAndFilterBranchFilterUnionParam{{
									OfProperty: &crm.PublicPropertyFilterParam{
										FilterType: crm.PublicPropertyFilterFilterTypeProperty,
										Operation: crm.PublicPropertyFilterOperationUnionParam{
											OfBool: &crm.PublicBoolPropertyOperationParam{
												IncludeObjectsWithNoValueSet: true,
												OperationType:                crm.PublicBoolPropertyOperationOperationTypeBool,
												Operator:                     "operator",
												Value:                        true,
											},
										},
										Property: "property",
									},
								}},
							},
						}},
						FilterBranchOperator: "filterBranchOperator",
						FilterBranchType:     crm.PublicOrFilterBranchFilterBranchTypeOr,
						Filters: []crm.PublicOrFilterBranchFilterUnionParam{{
							OfProperty: &crm.PublicPropertyFilterParam{
								FilterType: crm.PublicPropertyFilterFilterTypeProperty,
								Operation: crm.PublicPropertyFilterOperationUnionParam{
									OfBool: &crm.PublicBoolPropertyOperationParam{
										IncludeObjectsWithNoValueSet: true,
										OperationType:                crm.PublicBoolPropertyOperationOperationTypeBool,
										Operator:                     "operator",
										Value:                        true,
									},
								},
								Property: "property",
							},
						}},
					},
				},
			},
			EnrollObjectsInWorkflows: hubspotsdk.Bool(true),
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

func TestListUpdateListNameWithOptionalParams(t *testing.T) {
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
	_, err := client.Crm.Lists.UpdateListName(
		context.TODO(),
		"listId",
		crm.ListUpdateListNameParams{
			IncludeFilters: hubspotsdk.Bool(true),
			ListName:       hubspotsdk.String("listName"),
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

func TestListUpdateScheduleConversion(t *testing.T) {
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
	_, err := client.Crm.Lists.UpdateScheduleConversion(
		context.TODO(),
		"listId",
		crm.ListUpdateScheduleConversionParams{
			PublicListConversionTime: crm.PublicListConversionTimeUnionParam{
				OfConversionDate: &crm.PublicListConversionDateParam{
					ConversionType: crm.PublicListConversionDateConversionTypeConversionDate,
					Day:            0,
					Month:          0,
					Year:           0,
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
