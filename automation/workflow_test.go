// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package automation_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/stainless-sdks/hubspot-sdk-go"
	"github.com/stainless-sdks/hubspot-sdk-go/automation"
	"github.com/stainless-sdks/hubspot-sdk-go/internal/testutil"
	"github.com/stainless-sdks/hubspot-sdk-go/option"
	"github.com/stainless-sdks/hubspot-sdk-go/shared"
)

func TestWorkflowNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Automation.Workflows.New(context.TODO(), automation.WorkflowNewParams{
		APIFlowCreateRequest: automation.APIFlowCreateRequestUnionParam{
			OfAPIContactFlowCreateRequest: &automation.APIContactFlowCreateRequestParam{
				Actions: []automation.APIContactFlowCreateRequestActionUnionParam{{
					OfStaticBranch: &automation.APIStaticBranchActionParam{
						ActionID: "actionId",
						InputValue: automation.APIStaticBranchActionInputValueUnionParam{
							OfFieldData: &automation.APIActionDataValueParam{
								ActionID: "actionId",
								DataKey:  "dataKey",
								Type:     automation.APIActionDataValueTypeFieldData,
							},
						},
						StaticBranches: []automation.APIStaticBranchParam{{
							BranchValue: "branchValue",
							Connection: automation.APIConnectionParam{
								EdgeType:     "edgeType",
								NextActionID: "nextActionId",
							},
						}},
						Type: automation.APIStaticBranchActionTypeStaticBranch,
						DefaultBranch: automation.APIConnectionParam{
							EdgeType:     "edgeType",
							NextActionID: "nextActionId",
						},
						DefaultBranchName: hubspotsdk.String("defaultBranchName"),
					},
				}},
				BlockedDates: []automation.APIBlockedDateParam{{
					DayOfMonth: 0,
					Month:      automation.APIBlockedDateMonthApril,
					Year:       hubspotsdk.Int(0),
				}},
				CanEnrollFromSalesforce: true,
				CustomProperties: map[string]string{
					"foo": "string",
				},
				DataSources: []automation.APIContactFlowCreateRequestDataSourceUnionParam{{
					OfAssociation: &automation.APIAssociationDataSourceParam{
						AssociationCategory: automation.APIAssociationDataSourceAssociationCategoryHubspotDefined,
						AssociationTypeID:   0,
						Name:                "name",
						ObjectTypeID:        "objectTypeId",
						Type:                automation.APIAssociationDataSourceTypeAssociation,
						SortBy: automation.APISortParam{
							Order:    automation.APISortOrderAsc,
							Property: "property",
							Missing:  hubspotsdk.String("missing"),
						},
					},
				}},
				FlowType:           automation.APIContactFlowCreateRequestFlowTypeWorkflow,
				IsEnabled:          true,
				ObjectTypeID:       "objectTypeId",
				SuppressionListIDs: []int64{0},
				TimeWindows: []automation.APITimeWindowParam{{
					Day: automation.APITimeWindowDayFriday,
					EndTime: automation.APITimeOfDayParam{
						Hour:   0,
						Minute: 0,
					},
					StartTime: automation.APITimeOfDayParam{
						Hour:   0,
						Minute: 0,
					},
				}},
				Type:        automation.APIContactFlowCreateRequestTypeContactFlow,
				Description: hubspotsdk.String("description"),
				EnrollmentCriteria: automation.APIContactFlowCreateRequestEnrollmentCriteriaUnionParam{
					OfListBased: &automation.APIListBasedEnrollmentCriteriaParam{
						ListFilterBranch: automation.APIListBasedEnrollmentCriteriaListFilterBranchUnionParam{
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
										FilterBranchType:     shared.PublicAndFilterBranchFilterBranchTypeAnd,
										Filters: []shared.PublicAndFilterBranchFilterUnionParam{{
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
						ReEnrollmentTriggersFilterBranches: []automation.APIListBasedEnrollmentCriteriaReEnrollmentTriggersFilterBranchUnionParam{{
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
						ShouldReEnroll:                    true,
						Type:                              automation.APIListBasedEnrollmentCriteriaTypeListBased,
						UnEnrollObjectsNotMeetingCriteria: true,
					},
				},
				EnrollmentSchedule: automation.APIContactFlowCreateRequestEnrollmentScheduleUnionParam{
					OfDaily: &automation.APIDailyEnrollmentScheduleParam{
						TimeOfDay: automation.APITimeOfDayParam{
							Hour:   0,
							Minute: 0,
						},
						Type: automation.APIDailyEnrollmentScheduleTypeDaily,
					},
				},
				EventAnchor: automation.APIContactFlowCreateRequestEventAnchorUnionParam{
					OfContactPropertyAnchor: &automation.APIContactPropertyAnchorParam{
						ContactProperty: "contactProperty",
						Type:            automation.APIContactPropertyAnchorTypeContactPropertyAnchor,
					},
				},
				GoalFilterBranch: automation.APIContactFlowCreateRequestGoalFilterBranchUnionParam{
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
				},
				Name:          hubspotsdk.String("name"),
				StartActionID: hubspotsdk.String("startActionId"),
				UnEnrollmentSetting: automation.APIUnEnrollmentSettingParam{
					FlowIDs: []string{"string"},
					Type:    automation.APIUnEnrollmentSettingTypeAll,
				},
				Uuid: hubspotsdk.String("uuid"),
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

func TestWorkflowUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Automation.Workflows.Update(
		context.TODO(),
		"flowId",
		automation.WorkflowUpdateParams{
			APIFlowPutRequest: automation.APIFlowPutRequestUnionParam{
				OfAPIContactFlowPutRequest: &automation.APIContactFlowPutRequestParam{
					Actions: []automation.APIContactFlowPutRequestActionUnionParam{{
						OfStaticBranch: &automation.APIStaticBranchActionParam{
							ActionID: "actionId",
							InputValue: automation.APIStaticBranchActionInputValueUnionParam{
								OfFieldData: &automation.APIActionDataValueParam{
									ActionID: "actionId",
									DataKey:  "dataKey",
									Type:     automation.APIActionDataValueTypeFieldData,
								},
							},
							StaticBranches: []automation.APIStaticBranchParam{{
								BranchValue: "branchValue",
								Connection: automation.APIConnectionParam{
									EdgeType:     "edgeType",
									NextActionID: "nextActionId",
								},
							}},
							Type: automation.APIStaticBranchActionTypeStaticBranch,
							DefaultBranch: automation.APIConnectionParam{
								EdgeType:     "edgeType",
								NextActionID: "nextActionId",
							},
							DefaultBranchName: hubspotsdk.String("defaultBranchName"),
						},
					}},
					BlockedDates: []automation.APIBlockedDateParam{{
						DayOfMonth: 0,
						Month:      automation.APIBlockedDateMonthApril,
						Year:       hubspotsdk.Int(0),
					}},
					CanEnrollFromSalesforce: true,
					CustomProperties: map[string]string{
						"foo": "string",
					},
					IsEnabled:          true,
					RevisionID:         "revisionId",
					SuppressionListIDs: []int64{0},
					TimeWindows: []automation.APITimeWindowParam{{
						Day: automation.APITimeWindowDayFriday,
						EndTime: automation.APITimeOfDayParam{
							Hour:   0,
							Minute: 0,
						},
						StartTime: automation.APITimeOfDayParam{
							Hour:   0,
							Minute: 0,
						},
					}},
					Type:        automation.APIContactFlowPutRequestTypeContactFlow,
					Description: hubspotsdk.String("description"),
					EnrollmentCriteria: automation.APIContactFlowPutRequestEnrollmentCriteriaUnionParam{
						OfListBased: &automation.APIListBasedEnrollmentCriteriaParam{
							ListFilterBranch: automation.APIListBasedEnrollmentCriteriaListFilterBranchUnionParam{
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
											FilterBranchType:     shared.PublicAndFilterBranchFilterBranchTypeAnd,
											Filters: []shared.PublicAndFilterBranchFilterUnionParam{{
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
							ReEnrollmentTriggersFilterBranches: []automation.APIListBasedEnrollmentCriteriaReEnrollmentTriggersFilterBranchUnionParam{{
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
							ShouldReEnroll:                    true,
							Type:                              automation.APIListBasedEnrollmentCriteriaTypeListBased,
							UnEnrollObjectsNotMeetingCriteria: true,
						},
					},
					EnrollmentSchedule: automation.APIContactFlowPutRequestEnrollmentScheduleUnionParam{
						OfDaily: &automation.APIDailyEnrollmentScheduleParam{
							TimeOfDay: automation.APITimeOfDayParam{
								Hour:   0,
								Minute: 0,
							},
							Type: automation.APIDailyEnrollmentScheduleTypeDaily,
						},
					},
					EventAnchor: automation.APIContactFlowPutRequestEventAnchorUnionParam{
						OfContactPropertyAnchor: &automation.APIContactPropertyAnchorParam{
							ContactProperty: "contactProperty",
							Type:            automation.APIContactPropertyAnchorTypeContactPropertyAnchor,
						},
					},
					GoalFilterBranch: automation.APIContactFlowPutRequestGoalFilterBranchUnionParam{
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
					},
					Name:          hubspotsdk.String("name"),
					StartActionID: hubspotsdk.String("startActionId"),
					UnEnrollmentSetting: automation.APIUnEnrollmentSettingParam{
						FlowIDs: []string{"string"},
						Type:    automation.APIUnEnrollmentSettingTypeAll,
					},
					Uuid: hubspotsdk.String("uuid"),
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

func TestWorkflowListWithOptionalParams(t *testing.T) {
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
	_, err := client.Automation.Workflows.List(context.TODO(), automation.WorkflowListParams{
		After: hubspotsdk.String("after"),
		Limit: hubspotsdk.Int(0),
	})
	if err != nil {
		var apierr *hubspotsdk.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestWorkflowDelete(t *testing.T) {
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
	err := client.Automation.Workflows.Delete(context.TODO(), 0)
	if err != nil {
		var apierr *hubspotsdk.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestWorkflowBatchGet(t *testing.T) {
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
	_, err := client.Automation.Workflows.BatchGet(context.TODO(), automation.WorkflowBatchGetParams{
		APIFlowBatchInput: automation.APIFlowBatchInputParam{
			Inputs: []automation.APIFlowBatchFetchFlowIDCoordinateParam{{
				FlowID: "flowId",
				Type:   automation.APIFlowBatchFetchFlowIDCoordinateTypeFlowID,
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

func TestWorkflowBatchGetIDMappings(t *testing.T) {
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
	_, err := client.Automation.Workflows.BatchGetIDMappings(context.TODO(), automation.WorkflowBatchGetIDMappingsParams{
		APIFlowBatchMigrationInput: automation.APIFlowBatchMigrationInputParam{
			Inputs: []automation.APIFlowBatchMigrationInputInputUnionParam{{
				OfFlowID: &automation.APIFlowBatchFetchMigrationFlowIDCoordinateParam{
					FlowMigrationStatuses: "flowMigrationStatuses",
					Type:                  automation.APIFlowBatchFetchMigrationFlowIDCoordinateTypeFlowID,
				},
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

func TestWorkflowGet(t *testing.T) {
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
	_, err := client.Automation.Workflows.Get(context.TODO(), "flowId")
	if err != nil {
		var apierr *hubspotsdk.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestWorkflowListEmailCampaignsWithOptionalParams(t *testing.T) {
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
	_, err := client.Automation.Workflows.ListEmailCampaigns(context.TODO(), automation.WorkflowListEmailCampaignsParams{
		After:  hubspotsdk.String("after"),
		Before: hubspotsdk.String("before"),
		FlowID: []string{"string"},
		Limit:  hubspotsdk.Int(0),
	})
	if err != nil {
		var apierr *hubspotsdk.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
