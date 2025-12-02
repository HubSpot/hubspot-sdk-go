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

func TestListNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Crm.Lists.New(context.TODO(), crm.ListNewParams{
		ListCreateRequest: crm.ListCreateRequestParam{
			Name:           "Dynamic Association List Example",
			ObjectTypeID:   "0-1",
			ProcessingType: "DYNAMIC",
			CustomProperties: map[string]string{
				"foo": "string",
			},
			FilterBranch: crm.ListCreateRequestFilterBranchUnionParam{
				OfOr: &shared.PublicOrFilterBranchParam{
					FilterBranches: []shared.PublicOrFilterBranchFilterBranchUnionParam{{
						OfAnd: &shared.PublicAndFilterBranchParam{
							FilterBranches: []shared.PublicAndFilterBranchFilterBranchUnionParam{{
								OfNotAll: &shared.PublicNotAllFilterBranchParam{
									FilterBranches: []shared.PublicNotAllFilterBranchFilterBranchUnionParam{{
										OfNotAny: &shared.PublicNotAnyFilterBranchParam{
											FilterBranches: []shared.PublicNotAnyFilterBranchFilterBranchUnionParam{{
												OfRestricted: &shared.PublicRestrictedFilterBranchParam{
													FilterBranches: []shared.PublicRestrictedFilterBranchFilterBranchUnionParam{{
														OfUnifiedEvents: &shared.PublicUnifiedEventsFilterBranchParam{
															EventTypeID: "eventTypeId",
															FilterBranches: []shared.PublicUnifiedEventsFilterBranchFilterBranchUnionParam{{
																OfPropertyAssociation: &shared.PublicPropertyAssociationFilterBranchParam{
																	FilterBranches: []shared.PublicPropertyAssociationFilterBranchFilterBranchUnionParam{{
																		OfAssociation: &shared.PublicAssociationFilterBranchParam{
																			AssociationCategory: "associationCategory",
																			AssociationTypeID:   0,
																			FilterBranches: []shared.PublicAssociationFilterBranchFilterBranchUnionParam{{
																				OfOr: &shared.PublicOrFilterBranchParam{
																					FilterBranches:       []shared.PublicOrFilterBranchFilterBranchUnionParam{},
																					FilterBranchOperator: "filterBranchOperator",
																					FilterBranchType:     shared.PublicOrFilterBranchFilterBranchTypeOr,
																					Filters: []shared.PublicOrFilterBranchFilterUnionParam{{
																						OfProperty: &shared.PublicPropertyFilterParam{
																							FilterType: shared.PublicPropertyFilterFilterTypeProperty,
																							Operation: shared.PublicPropertyFilterOperationUnionParam{
																								OfBool: &shared.PublicBoolPropertyOperationParam{
																									IncludeObjectsWithNoValueSet: true,
																									OperationType:                shared.PublicBoolPropertyOperationOperationTypeBool,
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
																			FilterBranchType:     shared.PublicAssociationFilterBranchFilterBranchTypeAssociation,
																			Filters: []shared.PublicAssociationFilterBranchFilterUnionParam{{
																				OfProperty: &shared.PublicPropertyFilterParam{
																					FilterType: shared.PublicPropertyFilterFilterTypeProperty,
																					Operation: shared.PublicPropertyFilterOperationUnionParam{
																						OfBool: &shared.PublicBoolPropertyOperationParam{
																							IncludeObjectsWithNoValueSet: true,
																							OperationType:                shared.PublicBoolPropertyOperationOperationTypeBool,
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
																	FilterBranchType:     shared.PublicPropertyAssociationFilterBranchFilterBranchTypePropertyAssociation,
																	Filters: []shared.PublicPropertyAssociationFilterBranchFilterUnionParam{{
																		OfProperty: &shared.PublicPropertyFilterParam{
																			FilterType: shared.PublicPropertyFilterFilterTypeProperty,
																			Operation: shared.PublicPropertyFilterOperationUnionParam{
																				OfBool: &shared.PublicBoolPropertyOperationParam{
																					IncludeObjectsWithNoValueSet: true,
																					OperationType:                shared.PublicBoolPropertyOperationOperationTypeBool,
																					Operator:                     "operator",
																					Value:                        true,
																				},
																			},
																			Property: "property",
																		},
																	}},
																	ObjectTypeID:         "objectTypeId",
																	Operator:             "operator",
																	PropertyWithObjectID: "propertyWithObjectId",
																},
															}},
															FilterBranchOperator: "filterBranchOperator",
															FilterBranchType:     shared.PublicUnifiedEventsFilterBranchFilterBranchTypeUnifiedEvents,
															Filters: []shared.PublicUnifiedEventsFilterBranchFilterUnionParam{{
																OfProperty: &shared.PublicPropertyFilterParam{
																	FilterType: shared.PublicPropertyFilterFilterTypeProperty,
																	Operation: shared.PublicPropertyFilterOperationUnionParam{
																		OfBool: &shared.PublicBoolPropertyOperationParam{
																			IncludeObjectsWithNoValueSet: true,
																			OperationType:                shared.PublicBoolPropertyOperationOperationTypeBool,
																			Operator:                     "operator",
																			Value:                        true,
																		},
																	},
																	Property: "property",
																},
															}},
															Operator: shared.PublicUnifiedEventsFilterBranchOperatorHasCompleted,
															CoalescingRefineBy: shared.PublicUnifiedEventsFilterBranchCoalescingRefineByUnionParam{
																OfNumOccurrences: &shared.PublicNumOccurrencesRefineByParam{
																	Type:           shared.PublicNumOccurrencesRefineByTypeNumOccurrences,
																	MaxOccurrences: hubspotsdk.Int(0),
																	MinOccurrences: hubspotsdk.Int(0),
																},
															},
														},
													}},
													FilterBranchOperator: "filterBranchOperator",
													FilterBranchType:     shared.PublicRestrictedFilterBranchFilterBranchTypeRestricted,
													Filters: []shared.PublicRestrictedFilterBranchFilterUnionParam{{
														OfProperty: &shared.PublicPropertyFilterParam{
															FilterType: shared.PublicPropertyFilterFilterTypeProperty,
															Operation: shared.PublicPropertyFilterOperationUnionParam{
																OfBool: &shared.PublicBoolPropertyOperationParam{
																	IncludeObjectsWithNoValueSet: true,
																	OperationType:                shared.PublicBoolPropertyOperationOperationTypeBool,
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
											FilterBranchType:     shared.PublicNotAnyFilterBranchFilterBranchTypeNotAny,
											Filters: []shared.PublicNotAnyFilterBranchFilterUnionParam{{
												OfProperty: &shared.PublicPropertyFilterParam{
													FilterType: shared.PublicPropertyFilterFilterTypeProperty,
													Operation: shared.PublicPropertyFilterOperationUnionParam{
														OfBool: &shared.PublicBoolPropertyOperationParam{
															IncludeObjectsWithNoValueSet: true,
															OperationType:                shared.PublicBoolPropertyOperationOperationTypeBool,
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
									FilterBranchType:     shared.PublicNotAllFilterBranchFilterBranchTypeNotAll,
									Filters: []shared.PublicNotAllFilterBranchFilterUnionParam{{
										OfProperty: &shared.PublicPropertyFilterParam{
											FilterType: shared.PublicPropertyFilterFilterTypeProperty,
											Operation: shared.PublicPropertyFilterOperationUnionParam{
												OfBool: &shared.PublicBoolPropertyOperationParam{
													IncludeObjectsWithNoValueSet: true,
													OperationType:                shared.PublicBoolPropertyOperationOperationTypeBool,
													Operator:                     "IS_EQUAL_TO",
													Value:                        true,
												},
											},
											Property: "hs_is_closed_won",
										},
									}},
								},
							}},
							FilterBranchOperator: "filterBranchOperator",
							FilterBranchType:     shared.PublicAndFilterBranchFilterBranchTypeAnd,
							Filters: []shared.PublicAndFilterBranchFilterUnionParam{{
								OfProperty: &shared.PublicPropertyFilterParam{
									FilterType: shared.PublicPropertyFilterFilterTypeProperty,
									Operation: shared.PublicPropertyFilterOperationUnionParam{
										OfBool: &shared.PublicBoolPropertyOperationParam{
											IncludeObjectsWithNoValueSet: true,
											OperationType:                shared.PublicBoolPropertyOperationOperationTypeBool,
											Operator:                     "IS_EQUAL_TO",
											Value:                        true,
										},
									},
									Property: "firstname",
								},
							}},
						},
					}},
					FilterBranchOperator: "filterBranchOperator",
					FilterBranchType:     shared.PublicOrFilterBranchFilterBranchTypeOr,
					Filters: []shared.PublicOrFilterBranchFilterUnionParam{{
						OfProperty: &shared.PublicPropertyFilterParam{
							FilterType: shared.PublicPropertyFilterFilterTypeProperty,
							Operation: shared.PublicPropertyFilterOperationUnionParam{
								OfBool: &shared.PublicBoolPropertyOperationParam{
									IncludeObjectsWithNoValueSet: true,
									OperationType:                shared.PublicBoolPropertyOperationOperationTypeBool,
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
	err := client.Crm.Lists.Delete(context.TODO(), "listId")
	if err != nil {
		var apierr *hubspotsdk.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestListDeleteScheduleConversion(t *testing.T) {
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
	err := client.Crm.Lists.DeleteScheduleConversion(context.TODO(), "listId")
	if err != nil {
		var apierr *hubspotsdk.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestListGetWithOptionalParams(t *testing.T) {
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

func TestListGetByObjectTypeIDAndNameWithOptionalParams(t *testing.T) {
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
	_, err := client.Crm.Lists.GetByObjectTypeIDAndName(
		context.TODO(),
		"listName",
		crm.ListGetByObjectTypeIDAndNameParams{
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

func TestListGetScheduleConversion(t *testing.T) {
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
	_, err := client.Crm.Lists.GetScheduleConversion(context.TODO(), "listId")
	if err != nil {
		var apierr *hubspotsdk.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestListRestore(t *testing.T) {
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
	_, err := client.Crm.Lists.ScheduleConversion(
		context.TODO(),
		"listId",
		crm.ListScheduleConversionParams{
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

func TestListSearchWithOptionalParams(t *testing.T) {
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
	_, err := client.Crm.Lists.Search(context.TODO(), crm.ListSearchParams{
		ListSearchRequest: crm.ListSearchRequestParam{
			AdditionalProperties: []string{"hs_list_size_week_delta"},
			Offset:               0,
			Count:                hubspotsdk.Int(100),
			ListIDs:              []string{"string"},
			ProcessingTypes:      []string{"string"},
			Query:                hubspotsdk.String("Test"),
			Sort:                 hubspotsdk.String("sort"),
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

func TestListUpdateFiltersWithOptionalParams(t *testing.T) {
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
	_, err := client.Crm.Lists.UpdateFilters(
		context.TODO(),
		"listId",
		crm.ListUpdateFiltersParams{
			ListFilterUpdateRequest: crm.ListFilterUpdateRequestParam{
				FilterBranch: crm.ListFilterUpdateRequestFilterBranchUnionParam{
					OfOr: &shared.PublicOrFilterBranchParam{
						FilterBranches: []shared.PublicOrFilterBranchFilterBranchUnionParam{{
							OfOr: &shared.PublicOrFilterBranchParam{
								FilterBranches: []shared.PublicOrFilterBranchFilterBranchUnionParam{{
									OfNotAll: &shared.PublicNotAllFilterBranchParam{
										FilterBranches: []shared.PublicNotAllFilterBranchFilterBranchUnionParam{{
											OfNotAny: &shared.PublicNotAnyFilterBranchParam{
												FilterBranches: []shared.PublicNotAnyFilterBranchFilterBranchUnionParam{{
													OfRestricted: &shared.PublicRestrictedFilterBranchParam{
														FilterBranches: []shared.PublicRestrictedFilterBranchFilterBranchUnionParam{{
															OfUnifiedEvents: &shared.PublicUnifiedEventsFilterBranchParam{
																EventTypeID: "eventTypeId",
																FilterBranches: []shared.PublicUnifiedEventsFilterBranchFilterBranchUnionParam{{
																	OfPropertyAssociation: &shared.PublicPropertyAssociationFilterBranchParam{
																		FilterBranches: []shared.PublicPropertyAssociationFilterBranchFilterBranchUnionParam{{
																			OfAssociation: &shared.PublicAssociationFilterBranchParam{
																				AssociationCategory: "associationCategory",
																				AssociationTypeID:   0,
																				FilterBranches: []shared.PublicAssociationFilterBranchFilterBranchUnionParam{{
																					OfOr: &shared.PublicOrFilterBranchParam{
																						FilterBranches:       []shared.PublicOrFilterBranchFilterBranchUnionParam{},
																						FilterBranchOperator: "filterBranchOperator",
																						FilterBranchType:     shared.PublicOrFilterBranchFilterBranchTypeOr,
																						Filters: []shared.PublicOrFilterBranchFilterUnionParam{{
																							OfProperty: &shared.PublicPropertyFilterParam{
																								FilterType: shared.PublicPropertyFilterFilterTypeProperty,
																								Operation: shared.PublicPropertyFilterOperationUnionParam{
																									OfBool: &shared.PublicBoolPropertyOperationParam{
																										IncludeObjectsWithNoValueSet: true,
																										OperationType:                shared.PublicBoolPropertyOperationOperationTypeBool,
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
																				FilterBranchType:     shared.PublicAssociationFilterBranchFilterBranchTypeAssociation,
																				Filters: []shared.PublicAssociationFilterBranchFilterUnionParam{{
																					OfProperty: &shared.PublicPropertyFilterParam{
																						FilterType: shared.PublicPropertyFilterFilterTypeProperty,
																						Operation: shared.PublicPropertyFilterOperationUnionParam{
																							OfBool: &shared.PublicBoolPropertyOperationParam{
																								IncludeObjectsWithNoValueSet: true,
																								OperationType:                shared.PublicBoolPropertyOperationOperationTypeBool,
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
																		FilterBranchType:     shared.PublicPropertyAssociationFilterBranchFilterBranchTypePropertyAssociation,
																		Filters: []shared.PublicPropertyAssociationFilterBranchFilterUnionParam{{
																			OfProperty: &shared.PublicPropertyFilterParam{
																				FilterType: shared.PublicPropertyFilterFilterTypeProperty,
																				Operation: shared.PublicPropertyFilterOperationUnionParam{
																					OfBool: &shared.PublicBoolPropertyOperationParam{
																						IncludeObjectsWithNoValueSet: true,
																						OperationType:                shared.PublicBoolPropertyOperationOperationTypeBool,
																						Operator:                     "operator",
																						Value:                        true,
																					},
																				},
																				Property: "property",
																			},
																		}},
																		ObjectTypeID:         "objectTypeId",
																		Operator:             "operator",
																		PropertyWithObjectID: "propertyWithObjectId",
																	},
																}},
																FilterBranchOperator: "filterBranchOperator",
																FilterBranchType:     shared.PublicUnifiedEventsFilterBranchFilterBranchTypeUnifiedEvents,
																Filters: []shared.PublicUnifiedEventsFilterBranchFilterUnionParam{{
																	OfProperty: &shared.PublicPropertyFilterParam{
																		FilterType: shared.PublicPropertyFilterFilterTypeProperty,
																		Operation: shared.PublicPropertyFilterOperationUnionParam{
																			OfBool: &shared.PublicBoolPropertyOperationParam{
																				IncludeObjectsWithNoValueSet: true,
																				OperationType:                shared.PublicBoolPropertyOperationOperationTypeBool,
																				Operator:                     "operator",
																				Value:                        true,
																			},
																		},
																		Property: "property",
																	},
																}},
																Operator: shared.PublicUnifiedEventsFilterBranchOperatorHasCompleted,
																CoalescingRefineBy: shared.PublicUnifiedEventsFilterBranchCoalescingRefineByUnionParam{
																	OfNumOccurrences: &shared.PublicNumOccurrencesRefineByParam{
																		Type:           shared.PublicNumOccurrencesRefineByTypeNumOccurrences,
																		MaxOccurrences: hubspotsdk.Int(0),
																		MinOccurrences: hubspotsdk.Int(0),
																	},
																},
															},
														}},
														FilterBranchOperator: "filterBranchOperator",
														FilterBranchType:     shared.PublicRestrictedFilterBranchFilterBranchTypeRestricted,
														Filters: []shared.PublicRestrictedFilterBranchFilterUnionParam{{
															OfProperty: &shared.PublicPropertyFilterParam{
																FilterType: shared.PublicPropertyFilterFilterTypeProperty,
																Operation: shared.PublicPropertyFilterOperationUnionParam{
																	OfBool: &shared.PublicBoolPropertyOperationParam{
																		IncludeObjectsWithNoValueSet: true,
																		OperationType:                shared.PublicBoolPropertyOperationOperationTypeBool,
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
												FilterBranchType:     shared.PublicNotAnyFilterBranchFilterBranchTypeNotAny,
												Filters: []shared.PublicNotAnyFilterBranchFilterUnionParam{{
													OfProperty: &shared.PublicPropertyFilterParam{
														FilterType: shared.PublicPropertyFilterFilterTypeProperty,
														Operation: shared.PublicPropertyFilterOperationUnionParam{
															OfBool: &shared.PublicBoolPropertyOperationParam{
																IncludeObjectsWithNoValueSet: true,
																OperationType:                shared.PublicBoolPropertyOperationOperationTypeBool,
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
										FilterBranchType:     shared.PublicNotAllFilterBranchFilterBranchTypeNotAll,
										Filters: []shared.PublicNotAllFilterBranchFilterUnionParam{{
											OfProperty: &shared.PublicPropertyFilterParam{
												FilterType: shared.PublicPropertyFilterFilterTypeProperty,
												Operation: shared.PublicPropertyFilterOperationUnionParam{
													OfBool: &shared.PublicBoolPropertyOperationParam{
														IncludeObjectsWithNoValueSet: true,
														OperationType:                shared.PublicBoolPropertyOperationOperationTypeBool,
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
								FilterBranchType:     shared.PublicOrFilterBranchFilterBranchType("AND"),
								Filters: []shared.PublicOrFilterBranchFilterUnionParam{{
									OfProperty: &shared.PublicPropertyFilterParam{
										FilterType: shared.PublicPropertyFilterFilterTypeProperty,
										Operation: shared.PublicPropertyFilterOperationUnionParam{
											OfBool: &shared.PublicBoolPropertyOperationParam{
												IncludeObjectsWithNoValueSet: true,
												OperationType:                shared.PublicBoolPropertyOperationOperationTypeBool,
												Operator:                     "IS_GREATER_THAN_OR_EQUAL_TO",
												Value:                        true,
											},
										},
										Property: "hs_predictivecontactscore_v2",
									},
								}, {
									OfProperty: &shared.PublicPropertyFilterParam{
										FilterType: shared.PublicPropertyFilterFilterTypeProperty,
										Operation: shared.PublicPropertyFilterOperationUnionParam{
											OfBool: &shared.PublicBoolPropertyOperationParam{
												IncludeObjectsWithNoValueSet: true,
												OperationType:                shared.PublicBoolPropertyOperationOperationTypeBool,
												Operator:                     "IS_UNKNOWN",
												Value:                        true,
											},
										},
										Property: "engagements_last_meeting_booked_source",
									},
								}, {
									OfEmailSubscription: &shared.PublicEmailSubscriptionFilterParam{
										AcceptedStatuses: []string{"OPT_IN"},
										FilterType:       shared.PublicEmailSubscriptionFilterFilterTypeEmailSubscription,
										SubscriptionIDs:  []string{"81537745", "321981152"},
										SubscriptionType: hubspotsdk.String("subscriptionType"),
									},
								}},
							},
						}},
						FilterBranchOperator: "filterBranchOperator",
						FilterBranchType:     shared.PublicOrFilterBranchFilterBranchTypeOr,
						Filters: []shared.PublicOrFilterBranchFilterUnionParam{{
							OfProperty: &shared.PublicPropertyFilterParam{
								FilterType: shared.PublicPropertyFilterFilterTypeProperty,
								Operation: shared.PublicPropertyFilterOperationUnionParam{
									OfBool: &shared.PublicBoolPropertyOperationParam{
										IncludeObjectsWithNoValueSet: true,
										OperationType:                shared.PublicBoolPropertyOperationOperationTypeBool,
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

func TestListUpdateNameWithOptionalParams(t *testing.T) {
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
	_, err := client.Crm.Lists.UpdateName(
		context.TODO(),
		"listId",
		crm.ListUpdateNameParams{
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
