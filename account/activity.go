// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package account

import (
	"context"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/HubSpot/hubspot-sdk-go/internal/apijson"
	"github.com/HubSpot/hubspot-sdk-go/internal/apiquery"
	"github.com/HubSpot/hubspot-sdk-go/internal/requestconfig"
	"github.com/HubSpot/hubspot-sdk-go/option"
	"github.com/HubSpot/hubspot-sdk-go/packages/pagination"
	"github.com/HubSpot/hubspot-sdk-go/packages/param"
	"github.com/HubSpot/hubspot-sdk-go/packages/respjson"
	"github.com/HubSpot/hubspot-sdk-go/shared"
)

// ActivityService contains methods and other services that help with interacting
// with the hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewActivityService] method instead.
type ActivityService struct {
	options []option.RequestOption
}

// NewActivityService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewActivityService(opts ...option.RequestOption) (r ActivityService) {
	r = ActivityService{}
	r.options = opts
	return
}

// Retrieve activity history for user actions related to approvals, content
// updates, CRM object updates, security activity, and more (Enterprise only).
// Learn more about
// [activities included in audit log exports](https://knowledge.hubspot.com/account-management/view-and-export-account-activity-history-in-a-centralized-audit-log?hubs_content=knowledge.hubspot.com/account-management/view-and-export-account-activity-history&hubs_content-cta=centralized%20audit%20log#data-included-in-the-centralized-audit-log).
func (r *ActivityService) ListAuditLogs(ctx context.Context, query ActivityListAuditLogsParams, opts ...option.RequestOption) (res *pagination.Page[PublicAPIUserActionEvent], err error) {
	var raw *http.Response
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "account-info/2026-03/activity/audit-logs"
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// Retrieve activity history for user actions related to approvals, content
// updates, CRM object updates, security activity, and more (Enterprise only).
// Learn more about
// [activities included in audit log exports](https://knowledge.hubspot.com/account-management/view-and-export-account-activity-history-in-a-centralized-audit-log?hubs_content=knowledge.hubspot.com/account-management/view-and-export-account-activity-history&hubs_content-cta=centralized%20audit%20log#data-included-in-the-centralized-audit-log).
func (r *ActivityService) ListAuditLogsAutoPaging(ctx context.Context, query ActivityListAuditLogsParams, opts ...option.RequestOption) *pagination.PageAutoPager[PublicAPIUserActionEvent] {
	return pagination.NewPageAutoPager(r.ListAuditLogs(ctx, query, opts...))
}

// Retrieve logs of user actions related to
// [login activity](https://knowledge.hubspot.com/account-management/view-and-export-account-activity-history#account-login-history).
func (r *ActivityService) ListLoginActivities(ctx context.Context, query ActivityListLoginActivitiesParams, opts ...option.RequestOption) (res *pagination.Page[PublicLoginAudit], err error) {
	var raw *http.Response
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "account-info/2026-03/activity/login"
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// Retrieve logs of user actions related to
// [login activity](https://knowledge.hubspot.com/account-management/view-and-export-account-activity-history#account-login-history).
func (r *ActivityService) ListLoginActivitiesAutoPaging(ctx context.Context, query ActivityListLoginActivitiesParams, opts ...option.RequestOption) *pagination.PageAutoPager[PublicLoginAudit] {
	return pagination.NewPageAutoPager(r.ListLoginActivities(ctx, query, opts...))
}

// Retrieve logs of user actions related to
// [security activity](https://knowledge.hubspot.com/account-management/view-and-export-account-activity-history#security-activity-history).
func (r *ActivityService) ListSecurityActivities(ctx context.Context, query ActivityListSecurityActivitiesParams, opts ...option.RequestOption) (res *pagination.Page[HydratedCriticalAction], err error) {
	var raw *http.Response
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "account-info/2026-03/activity/security"
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// Retrieve logs of user actions related to
// [security activity](https://knowledge.hubspot.com/account-management/view-and-export-account-activity-history#security-activity-history).
func (r *ActivityService) ListSecurityActivitiesAutoPaging(ctx context.Context, query ActivityListSecurityActivitiesParams, opts ...option.RequestOption) *pagination.PageAutoPager[HydratedCriticalAction] {
	return pagination.NewPageAutoPager(r.ListSecurityActivities(ctx, query, opts...))
}

type ActingUser struct {
	// The user's unique ID.
	UserID int64 `json:"userId" api:"required"`
	// The email address of the user who performed the action.
	UserEmail string `json:"userEmail"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		UserID      respjson.Field
		UserEmail   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ActingUser) RawJSON() string { return r.JSON.raw }
func (r *ActingUser) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CollectionResponseHydratedCriticalActionForwardPaging struct {
	Results []HydratedCriticalAction `json:"results" api:"required"`
	Paging  shared.ForwardPaging     `json:"paging"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Results     respjson.Field
		Paging      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CollectionResponseHydratedCriticalActionForwardPaging) RawJSON() string { return r.JSON.raw }
func (r *CollectionResponseHydratedCriticalActionForwardPaging) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CollectionResponsePublicAPIUserActionEventForwardPaging struct {
	Results []PublicAPIUserActionEvent `json:"results" api:"required"`
	Paging  shared.ForwardPaging       `json:"paging"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Results     respjson.Field
		Paging      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CollectionResponsePublicAPIUserActionEventForwardPaging) RawJSON() string { return r.JSON.raw }
func (r *CollectionResponsePublicAPIUserActionEventForwardPaging) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CollectionResponsePublicLoginAuditForwardPaging struct {
	Results []PublicLoginAudit   `json:"results" api:"required"`
	Paging  shared.ForwardPaging `json:"paging"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Results     respjson.Field
		Paging      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CollectionResponsePublicLoginAuditForwardPaging) RawJSON() string { return r.JSON.raw }
func (r *CollectionResponsePublicLoginAuditForwardPaging) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type HydratedCriticalAction struct {
	// The activity's unique ID.
	ID string `json:"id" api:"required"`
	// The time the activity took place.
	CreatedAt time.Time `json:"createdAt" api:"required" format:"date-time"`
	// The type of activity.
	//
	// Any of "ACCEPTANCE_TEST", "ACCOUNT_ADDED_TO_MULTI_ACCOUNT_ORGANIZATION",
	// "ACCOUNT_REMOVED_FROM_MULTI_ACCOUNT_ORGANIZATION", "ADD_ADMIN_PERMISSIONS",
	// "ADD_ADMIN_USER", "ADD_SINGLE_SIGN_ON", "ADD_TWO_FACTOR_AUTHENTICATION",
	// "ADD_USER", "ADD_WEBHOOK_IN_WORKFLOW", "ALLOWED_GEOLOCATIONS_DISABLED",
	// "ALLOWED_GEOLOCATIONS_ENABLED", "ALLOWED_LOGIN_METHODS_DELETE",
	// "ALLOWED_LOGIN_METHODS_UPDATE", "ATTACHMENT_LOGGING_DISABLED",
	// "ATTACHMENT_LOGGING_ENABLED", "AUTOMATED_INACTIVE_USER_DEACTIVATION_DISABLED",
	// "AUTOMATED_INACTIVE_USER_DEACTIVATION_ENABLED", "BOTS_WEBHOOK_POST",
	// "BOTS_WEBHOOK_UPDATE", "BOTS_WEBHOOK_VIEWED", "BULK_EMAIL_DOMAIN_CHANGE",
	// "CHANGE_AD_EVENT_CONSENT_SETTING", "CHANGE_AD_EVENT_DATA_SHARING_SETTING",
	// "CHANGE_PASSWORD", "CONTACT_DATA_EXPORT", "DATA_ACCESS_REQUEST_SUBMITTED",
	// "DATA_BACKUP_CREATED", "DATA_BACKUP_DOWNLOADED", "DATA_BACKUP_SCHEDULE_CREATED",
	// "DATA_BACKUP_SCHEDULE_DELETED", "DATA_BACKUP_SCHEDULE_UPDATED",
	// "DATA_RESTORE_COMPLETED", "DATA_SHARING_CONNECTION_ADDED",
	// "DATA_SHARING_CONNECTION_REMOVED", "DATASET_SYNC", "DEACTIVATE_USER",
	// "DOMAIN_BASED_INVITE_CREATED", "DOMAIN_BASED_INVITE_REMOVED",
	// "DOMAIN_BASED_INVITES_DISABLED", "DOMAIN_BASED_INVITES_ENABLED",
	// "EMAIL_TRACKING_DISABLED", "EMAIL_TRACKING_ENABLED", "EXPORT",
	// "EXPORT_APPROVAL", "EXPORT_DOWNLOAD", "EXPORT_USERS", "FORM_SUBMISSIONS_EXPORT",
	// "GDPR_DELETE", "GDPR_TOGGLE_DISABLED", "GDPR_TOGGLE_ENABLED", "HAPIKEY_CREATE",
	// "HAPIKEY_DEACTIVATE", "HAPIKEY_VIEW", "HUBSPOT_EMPLOYEE_ACCESS_DISABLED",
	// "HUBSPOT_EMPLOYEE_ACCESS_ENABLED", "IMPERSONATE_USER", "IMPORT",
	// "INSTALL_INTEGRATION", "IP_RESTRICTIONS_DISABLED", "IP_RESTRICTIONS_ENABLED",
	// "JOINED_PORTAL_VIA_DOMAIN_BASED_INVITE", "LEGAL_BASIS_REQUIREMENT_DISABLED",
	// "LEGAL_BASIS_REQUIREMENT_ENABLED", "MANUAL_PASSWORD_RESET_EMAIL_SEND",
	// "MANUAL_REGISTRATION_EMAIL_SEND", "MARKETING_CONTACTS_APP_SETTINGS_DISABLED",
	// "MARKETING_CONTACTS_APP_SETTINGS_ENABLED", "MERGE_REVERT",
	// "MODIFY_WEBHOOK_IN_WORKFLOW", "MULTI_ACCOUNT_REPORTING_CONNECTION_ADDED",
	// "MULTI_ACCOUNT_REPORTING_CONNECTION_REMOVED",
	// "MULTI_ACCOUNT_WORKFLOWS_CONNECTION_ADDED",
	// "MULTI_ACCOUNT_WORKFLOWS_CONNECTION_REMOVED", "NEVER_LOG_FOR_PORTAL_ADDITION",
	// "NEVER_LOG_FOR_PORTAL_DELETION", "NEVER_LOG_FOR_USER_ADDITION",
	// "NEVER_LOG_FOR_USER_DELETION", "PASSKEY_ADDED", "PASSKEY_DELETED",
	// "PAYMENT_ACCOUNT_CREATION", "PAYMENT_ACCOUNT_INFO_UPDATE",
	// "PAYMENT_BANK_ACCOUNT_CHANGE", "PAYMENT_ONBOARDING_LINK_SEND",
	// "PERSONAL_ACCESS_KEY_CREATE", "PERSONAL_ACCESS_KEY_DEACTIVATE",
	// "PERSONAL_ACCESS_KEY_ROTATE", "PERSONAL_ACCESS_KEY_VIEW",
	// "PRIVATE_APP_ACCESS_TOKEN_CREATE", "PRIVATE_APP_ACCESS_TOKEN_DEACTIVATE",
	// "PRIVATE_APP_ACCESS_TOKEN_ROTATE", "PRIVATE_APP_ACCESS_TOKEN_VIEW",
	// "PRIVATE_APP_CLIENT_SECRET_VIEW", "PRIVATE_APP_CLIENT_SECRET_WRITE",
	// "PRIVATE_APP_SCOPE_GROUPS_UPDATE", "PRODUCTION_DEPLOYMENT",
	// "PROPERTY_HISTORY_REVISION", "PUBLIC_APP_CLIENT_SECRET_VIEW",
	// "PUBLIC_APP_CLIENT_SECRET_WRITE", "REACTIVATE_USER", "REMOVE_ADMIN_PERMISSIONS",
	// "REMOVE_ADMIN_USER", "REMOVE_SINGLE_SIGN_ON",
	// "REMOVE_TWO_FACTOR_AUTHENTICATION", "REMOVE_USER", "REQUIRE_SINGLE_SIGN_ON",
	// "RESTRICTED_LIST_ADDED_TO_CONTENT", "SANDBOX_CREATION", "SANDBOX_DELETION",
	// "SANDBOX_SYNC", "SANDBOX_SYNC_TO_PRODUCTION",
	// "SECRET_ADDED_TO_SERVERLESS_FUNCTION", "SENSITIVE_DATA_DISABLED",
	// "SENSITIVE_DATA_ENABLED", "SEQUENCE_CLONED", "SEQUENCE_CREATED",
	// "SEQUENCE_ENROLLMENT_INITIATED", "SEQUENCE_ENROLLMENT_STATE_CHANGED",
	// "SEQUENCE_MODIFIED", "SERVICE_KEY_CREATE", "SERVICE_KEY_DEACTIVATE",
	// "SERVICE_KEY_PERMISSIONS_UPDATE", "SERVICE_KEY_REVEAL", "SERVICE_KEY_ROTATE",
	// "SMTP_TOKEN_CREATED", "SMTP_TOKEN_DELETED", "SMTP_TOKEN_PASSWORD_RESET",
	// "SMTP_TOKEN_RETRIEVED", "TEAM_ADDED", "TEAM_DELETED", "TEAM_USER_ADDED",
	// "TEAM_USER_DELETED", "TEMPLATE_DELETED", "TEMPLATE_MODIFIED",
	// "TOUCHLESS_PURCHASE", "UNIFIED_RESTORE_UNDO_EXECUTION", "UNINSTALL_INTEGRATION",
	// "UNREQUIRE_SINGLE_SIGN_ON", "WEBHOOK_SETTINGS_UPDATE",
	// "WEBHOOK_SUBSCRIPTION_CREATE", "WEBHOOK_SUBSCRIPTION_UPDATE".
	Type HydratedCriticalActionType `json:"type" api:"required"`
	// The user's unique ID.
	UserID int64 `json:"userId" api:"required"`
	// Email address of the user associated with the activity.
	ActingUser string `json:"actingUser"`
	// The approximate country code
	CountryCode string `json:"countryCode"`
	// A link to the URL where the action was taken in the account.
	InfoURL string `json:"infoUrl"`
	// IP address where the activity originated.
	IPAddress string `json:"ipAddress"`
	// The approximate location where the activity took place.
	Location string `json:"location"`
	// The ID of the affected object.
	ObjectID string `json:"objectId"`
	// The approximate region code
	RegionCode string `json:"regionCode"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		Type        respjson.Field
		UserID      respjson.Field
		ActingUser  respjson.Field
		CountryCode respjson.Field
		InfoURL     respjson.Field
		IPAddress   respjson.Field
		Location    respjson.Field
		ObjectID    respjson.Field
		RegionCode  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r HydratedCriticalAction) RawJSON() string { return r.JSON.raw }
func (r *HydratedCriticalAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The type of activity.
type HydratedCriticalActionType string

const (
	HydratedCriticalActionTypeAcceptanceTest                             HydratedCriticalActionType = "ACCEPTANCE_TEST"
	HydratedCriticalActionTypeAccountAddedToMultiAccountOrganization     HydratedCriticalActionType = "ACCOUNT_ADDED_TO_MULTI_ACCOUNT_ORGANIZATION"
	HydratedCriticalActionTypeAccountRemovedFromMultiAccountOrganization HydratedCriticalActionType = "ACCOUNT_REMOVED_FROM_MULTI_ACCOUNT_ORGANIZATION"
	HydratedCriticalActionTypeAddAdminPermissions                        HydratedCriticalActionType = "ADD_ADMIN_PERMISSIONS"
	HydratedCriticalActionTypeAddAdminUser                               HydratedCriticalActionType = "ADD_ADMIN_USER"
	HydratedCriticalActionTypeAddSingleSignOn                            HydratedCriticalActionType = "ADD_SINGLE_SIGN_ON"
	HydratedCriticalActionTypeAddTwoFactorAuthentication                 HydratedCriticalActionType = "ADD_TWO_FACTOR_AUTHENTICATION"
	HydratedCriticalActionTypeAddUser                                    HydratedCriticalActionType = "ADD_USER"
	HydratedCriticalActionTypeAddWebhookInWorkflow                       HydratedCriticalActionType = "ADD_WEBHOOK_IN_WORKFLOW"
	HydratedCriticalActionTypeAllowedGeolocationsDisabled                HydratedCriticalActionType = "ALLOWED_GEOLOCATIONS_DISABLED"
	HydratedCriticalActionTypeAllowedGeolocationsEnabled                 HydratedCriticalActionType = "ALLOWED_GEOLOCATIONS_ENABLED"
	HydratedCriticalActionTypeAllowedLoginMethodsDelete                  HydratedCriticalActionType = "ALLOWED_LOGIN_METHODS_DELETE"
	HydratedCriticalActionTypeAllowedLoginMethodsUpdate                  HydratedCriticalActionType = "ALLOWED_LOGIN_METHODS_UPDATE"
	HydratedCriticalActionTypeAttachmentLoggingDisabled                  HydratedCriticalActionType = "ATTACHMENT_LOGGING_DISABLED"
	HydratedCriticalActionTypeAttachmentLoggingEnabled                   HydratedCriticalActionType = "ATTACHMENT_LOGGING_ENABLED"
	HydratedCriticalActionTypeAutomatedInactiveUserDeactivationDisabled  HydratedCriticalActionType = "AUTOMATED_INACTIVE_USER_DEACTIVATION_DISABLED"
	HydratedCriticalActionTypeAutomatedInactiveUserDeactivationEnabled   HydratedCriticalActionType = "AUTOMATED_INACTIVE_USER_DEACTIVATION_ENABLED"
	HydratedCriticalActionTypeBotsWebhookPost                            HydratedCriticalActionType = "BOTS_WEBHOOK_POST"
	HydratedCriticalActionTypeBotsWebhookUpdate                          HydratedCriticalActionType = "BOTS_WEBHOOK_UPDATE"
	HydratedCriticalActionTypeBotsWebhookViewed                          HydratedCriticalActionType = "BOTS_WEBHOOK_VIEWED"
	HydratedCriticalActionTypeBulkEmailDomainChange                      HydratedCriticalActionType = "BULK_EMAIL_DOMAIN_CHANGE"
	HydratedCriticalActionTypeChangeAdEventConsentSetting                HydratedCriticalActionType = "CHANGE_AD_EVENT_CONSENT_SETTING"
	HydratedCriticalActionTypeChangeAdEventDataSharingSetting            HydratedCriticalActionType = "CHANGE_AD_EVENT_DATA_SHARING_SETTING"
	HydratedCriticalActionTypeChangePassword                             HydratedCriticalActionType = "CHANGE_PASSWORD"
	HydratedCriticalActionTypeContactDataExport                          HydratedCriticalActionType = "CONTACT_DATA_EXPORT"
	HydratedCriticalActionTypeDataAccessRequestSubmitted                 HydratedCriticalActionType = "DATA_ACCESS_REQUEST_SUBMITTED"
	HydratedCriticalActionTypeDataBackupCreated                          HydratedCriticalActionType = "DATA_BACKUP_CREATED"
	HydratedCriticalActionTypeDataBackupDownloaded                       HydratedCriticalActionType = "DATA_BACKUP_DOWNLOADED"
	HydratedCriticalActionTypeDataBackupScheduleCreated                  HydratedCriticalActionType = "DATA_BACKUP_SCHEDULE_CREATED"
	HydratedCriticalActionTypeDataBackupScheduleDeleted                  HydratedCriticalActionType = "DATA_BACKUP_SCHEDULE_DELETED"
	HydratedCriticalActionTypeDataBackupScheduleUpdated                  HydratedCriticalActionType = "DATA_BACKUP_SCHEDULE_UPDATED"
	HydratedCriticalActionTypeDataRestoreCompleted                       HydratedCriticalActionType = "DATA_RESTORE_COMPLETED"
	HydratedCriticalActionTypeDataSharingConnectionAdded                 HydratedCriticalActionType = "DATA_SHARING_CONNECTION_ADDED"
	HydratedCriticalActionTypeDataSharingConnectionRemoved               HydratedCriticalActionType = "DATA_SHARING_CONNECTION_REMOVED"
	HydratedCriticalActionTypeDatasetSync                                HydratedCriticalActionType = "DATASET_SYNC"
	HydratedCriticalActionTypeDeactivateUser                             HydratedCriticalActionType = "DEACTIVATE_USER"
	HydratedCriticalActionTypeDomainBasedInviteCreated                   HydratedCriticalActionType = "DOMAIN_BASED_INVITE_CREATED"
	HydratedCriticalActionTypeDomainBasedInviteRemoved                   HydratedCriticalActionType = "DOMAIN_BASED_INVITE_REMOVED"
	HydratedCriticalActionTypeDomainBasedInvitesDisabled                 HydratedCriticalActionType = "DOMAIN_BASED_INVITES_DISABLED"
	HydratedCriticalActionTypeDomainBasedInvitesEnabled                  HydratedCriticalActionType = "DOMAIN_BASED_INVITES_ENABLED"
	HydratedCriticalActionTypeEmailTrackingDisabled                      HydratedCriticalActionType = "EMAIL_TRACKING_DISABLED"
	HydratedCriticalActionTypeEmailTrackingEnabled                       HydratedCriticalActionType = "EMAIL_TRACKING_ENABLED"
	HydratedCriticalActionTypeExport                                     HydratedCriticalActionType = "EXPORT"
	HydratedCriticalActionTypeExportApproval                             HydratedCriticalActionType = "EXPORT_APPROVAL"
	HydratedCriticalActionTypeExportDownload                             HydratedCriticalActionType = "EXPORT_DOWNLOAD"
	HydratedCriticalActionTypeExportUsers                                HydratedCriticalActionType = "EXPORT_USERS"
	HydratedCriticalActionTypeFormSubmissionsExport                      HydratedCriticalActionType = "FORM_SUBMISSIONS_EXPORT"
	HydratedCriticalActionTypeGdprDelete                                 HydratedCriticalActionType = "GDPR_DELETE"
	HydratedCriticalActionTypeGdprToggleDisabled                         HydratedCriticalActionType = "GDPR_TOGGLE_DISABLED"
	HydratedCriticalActionTypeGdprToggleEnabled                          HydratedCriticalActionType = "GDPR_TOGGLE_ENABLED"
	HydratedCriticalActionTypeHapikeyCreate                              HydratedCriticalActionType = "HAPIKEY_CREATE"
	HydratedCriticalActionTypeHapikeyDeactivate                          HydratedCriticalActionType = "HAPIKEY_DEACTIVATE"
	HydratedCriticalActionTypeHapikeyView                                HydratedCriticalActionType = "HAPIKEY_VIEW"
	HydratedCriticalActionTypeHubSpotEmployeeAccessDisabled              HydratedCriticalActionType = "HUBSPOT_EMPLOYEE_ACCESS_DISABLED"
	HydratedCriticalActionTypeHubSpotEmployeeAccessEnabled               HydratedCriticalActionType = "HUBSPOT_EMPLOYEE_ACCESS_ENABLED"
	HydratedCriticalActionTypeImpersonateUser                            HydratedCriticalActionType = "IMPERSONATE_USER"
	HydratedCriticalActionTypeImport                                     HydratedCriticalActionType = "IMPORT"
	HydratedCriticalActionTypeInstallIntegration                         HydratedCriticalActionType = "INSTALL_INTEGRATION"
	HydratedCriticalActionTypeIPRestrictionsDisabled                     HydratedCriticalActionType = "IP_RESTRICTIONS_DISABLED"
	HydratedCriticalActionTypeIPRestrictionsEnabled                      HydratedCriticalActionType = "IP_RESTRICTIONS_ENABLED"
	HydratedCriticalActionTypeJoinedPortalViaDomainBasedInvite           HydratedCriticalActionType = "JOINED_PORTAL_VIA_DOMAIN_BASED_INVITE"
	HydratedCriticalActionTypeLegalBasisRequirementDisabled              HydratedCriticalActionType = "LEGAL_BASIS_REQUIREMENT_DISABLED"
	HydratedCriticalActionTypeLegalBasisRequirementEnabled               HydratedCriticalActionType = "LEGAL_BASIS_REQUIREMENT_ENABLED"
	HydratedCriticalActionTypeManualPasswordResetEmailSend               HydratedCriticalActionType = "MANUAL_PASSWORD_RESET_EMAIL_SEND"
	HydratedCriticalActionTypeManualRegistrationEmailSend                HydratedCriticalActionType = "MANUAL_REGISTRATION_EMAIL_SEND"
	HydratedCriticalActionTypeMarketingContactsAppSettingsDisabled       HydratedCriticalActionType = "MARKETING_CONTACTS_APP_SETTINGS_DISABLED"
	HydratedCriticalActionTypeMarketingContactsAppSettingsEnabled        HydratedCriticalActionType = "MARKETING_CONTACTS_APP_SETTINGS_ENABLED"
	HydratedCriticalActionTypeMergeRevert                                HydratedCriticalActionType = "MERGE_REVERT"
	HydratedCriticalActionTypeModifyWebhookInWorkflow                    HydratedCriticalActionType = "MODIFY_WEBHOOK_IN_WORKFLOW"
	HydratedCriticalActionTypeMultiAccountReportingConnectionAdded       HydratedCriticalActionType = "MULTI_ACCOUNT_REPORTING_CONNECTION_ADDED"
	HydratedCriticalActionTypeMultiAccountReportingConnectionRemoved     HydratedCriticalActionType = "MULTI_ACCOUNT_REPORTING_CONNECTION_REMOVED"
	HydratedCriticalActionTypeMultiAccountWorkflowsConnectionAdded       HydratedCriticalActionType = "MULTI_ACCOUNT_WORKFLOWS_CONNECTION_ADDED"
	HydratedCriticalActionTypeMultiAccountWorkflowsConnectionRemoved     HydratedCriticalActionType = "MULTI_ACCOUNT_WORKFLOWS_CONNECTION_REMOVED"
	HydratedCriticalActionTypeNeverLogForPortalAddition                  HydratedCriticalActionType = "NEVER_LOG_FOR_PORTAL_ADDITION"
	HydratedCriticalActionTypeNeverLogForPortalDeletion                  HydratedCriticalActionType = "NEVER_LOG_FOR_PORTAL_DELETION"
	HydratedCriticalActionTypeNeverLogForUserAddition                    HydratedCriticalActionType = "NEVER_LOG_FOR_USER_ADDITION"
	HydratedCriticalActionTypeNeverLogForUserDeletion                    HydratedCriticalActionType = "NEVER_LOG_FOR_USER_DELETION"
	HydratedCriticalActionTypePasskeyAdded                               HydratedCriticalActionType = "PASSKEY_ADDED"
	HydratedCriticalActionTypePasskeyDeleted                             HydratedCriticalActionType = "PASSKEY_DELETED"
	HydratedCriticalActionTypePaymentAccountCreation                     HydratedCriticalActionType = "PAYMENT_ACCOUNT_CREATION"
	HydratedCriticalActionTypePaymentAccountInfoUpdate                   HydratedCriticalActionType = "PAYMENT_ACCOUNT_INFO_UPDATE"
	HydratedCriticalActionTypePaymentBankAccountChange                   HydratedCriticalActionType = "PAYMENT_BANK_ACCOUNT_CHANGE"
	HydratedCriticalActionTypePaymentOnboardingLinkSend                  HydratedCriticalActionType = "PAYMENT_ONBOARDING_LINK_SEND"
	HydratedCriticalActionTypePersonalAccessKeyCreate                    HydratedCriticalActionType = "PERSONAL_ACCESS_KEY_CREATE"
	HydratedCriticalActionTypePersonalAccessKeyDeactivate                HydratedCriticalActionType = "PERSONAL_ACCESS_KEY_DEACTIVATE"
	HydratedCriticalActionTypePersonalAccessKeyRotate                    HydratedCriticalActionType = "PERSONAL_ACCESS_KEY_ROTATE"
	HydratedCriticalActionTypePersonalAccessKeyView                      HydratedCriticalActionType = "PERSONAL_ACCESS_KEY_VIEW"
	HydratedCriticalActionTypePrivateAppAccessTokenCreate                HydratedCriticalActionType = "PRIVATE_APP_ACCESS_TOKEN_CREATE"
	HydratedCriticalActionTypePrivateAppAccessTokenDeactivate            HydratedCriticalActionType = "PRIVATE_APP_ACCESS_TOKEN_DEACTIVATE"
	HydratedCriticalActionTypePrivateAppAccessTokenRotate                HydratedCriticalActionType = "PRIVATE_APP_ACCESS_TOKEN_ROTATE"
	HydratedCriticalActionTypePrivateAppAccessTokenView                  HydratedCriticalActionType = "PRIVATE_APP_ACCESS_TOKEN_VIEW"
	HydratedCriticalActionTypePrivateAppClientSecretView                 HydratedCriticalActionType = "PRIVATE_APP_CLIENT_SECRET_VIEW"
	HydratedCriticalActionTypePrivateAppClientSecretWrite                HydratedCriticalActionType = "PRIVATE_APP_CLIENT_SECRET_WRITE"
	HydratedCriticalActionTypePrivateAppScopeGroupsUpdate                HydratedCriticalActionType = "PRIVATE_APP_SCOPE_GROUPS_UPDATE"
	HydratedCriticalActionTypeProductionDeployment                       HydratedCriticalActionType = "PRODUCTION_DEPLOYMENT"
	HydratedCriticalActionTypePropertyHistoryRevision                    HydratedCriticalActionType = "PROPERTY_HISTORY_REVISION"
	HydratedCriticalActionTypePublicAppClientSecretView                  HydratedCriticalActionType = "PUBLIC_APP_CLIENT_SECRET_VIEW"
	HydratedCriticalActionTypePublicAppClientSecretWrite                 HydratedCriticalActionType = "PUBLIC_APP_CLIENT_SECRET_WRITE"
	HydratedCriticalActionTypeReactivateUser                             HydratedCriticalActionType = "REACTIVATE_USER"
	HydratedCriticalActionTypeRemoveAdminPermissions                     HydratedCriticalActionType = "REMOVE_ADMIN_PERMISSIONS"
	HydratedCriticalActionTypeRemoveAdminUser                            HydratedCriticalActionType = "REMOVE_ADMIN_USER"
	HydratedCriticalActionTypeRemoveSingleSignOn                         HydratedCriticalActionType = "REMOVE_SINGLE_SIGN_ON"
	HydratedCriticalActionTypeRemoveTwoFactorAuthentication              HydratedCriticalActionType = "REMOVE_TWO_FACTOR_AUTHENTICATION"
	HydratedCriticalActionTypeRemoveUser                                 HydratedCriticalActionType = "REMOVE_USER"
	HydratedCriticalActionTypeRequireSingleSignOn                        HydratedCriticalActionType = "REQUIRE_SINGLE_SIGN_ON"
	HydratedCriticalActionTypeRestrictedListAddedToContent               HydratedCriticalActionType = "RESTRICTED_LIST_ADDED_TO_CONTENT"
	HydratedCriticalActionTypeSandboxCreation                            HydratedCriticalActionType = "SANDBOX_CREATION"
	HydratedCriticalActionTypeSandboxDeletion                            HydratedCriticalActionType = "SANDBOX_DELETION"
	HydratedCriticalActionTypeSandboxSync                                HydratedCriticalActionType = "SANDBOX_SYNC"
	HydratedCriticalActionTypeSandboxSyncToProduction                    HydratedCriticalActionType = "SANDBOX_SYNC_TO_PRODUCTION"
	HydratedCriticalActionTypeSecretAddedToServerlessFunction            HydratedCriticalActionType = "SECRET_ADDED_TO_SERVERLESS_FUNCTION"
	HydratedCriticalActionTypeSensitiveDataDisabled                      HydratedCriticalActionType = "SENSITIVE_DATA_DISABLED"
	HydratedCriticalActionTypeSensitiveDataEnabled                       HydratedCriticalActionType = "SENSITIVE_DATA_ENABLED"
	HydratedCriticalActionTypeSequenceCloned                             HydratedCriticalActionType = "SEQUENCE_CLONED"
	HydratedCriticalActionTypeSequenceCreated                            HydratedCriticalActionType = "SEQUENCE_CREATED"
	HydratedCriticalActionTypeSequenceEnrollmentInitiated                HydratedCriticalActionType = "SEQUENCE_ENROLLMENT_INITIATED"
	HydratedCriticalActionTypeSequenceEnrollmentStateChanged             HydratedCriticalActionType = "SEQUENCE_ENROLLMENT_STATE_CHANGED"
	HydratedCriticalActionTypeSequenceModified                           HydratedCriticalActionType = "SEQUENCE_MODIFIED"
	HydratedCriticalActionTypeServiceKeyCreate                           HydratedCriticalActionType = "SERVICE_KEY_CREATE"
	HydratedCriticalActionTypeServiceKeyDeactivate                       HydratedCriticalActionType = "SERVICE_KEY_DEACTIVATE"
	HydratedCriticalActionTypeServiceKeyPermissionsUpdate                HydratedCriticalActionType = "SERVICE_KEY_PERMISSIONS_UPDATE"
	HydratedCriticalActionTypeServiceKeyReveal                           HydratedCriticalActionType = "SERVICE_KEY_REVEAL"
	HydratedCriticalActionTypeServiceKeyRotate                           HydratedCriticalActionType = "SERVICE_KEY_ROTATE"
	HydratedCriticalActionTypeSmtpTokenCreated                           HydratedCriticalActionType = "SMTP_TOKEN_CREATED"
	HydratedCriticalActionTypeSmtpTokenDeleted                           HydratedCriticalActionType = "SMTP_TOKEN_DELETED"
	HydratedCriticalActionTypeSmtpTokenPasswordReset                     HydratedCriticalActionType = "SMTP_TOKEN_PASSWORD_RESET"
	HydratedCriticalActionTypeSmtpTokenRetrieved                         HydratedCriticalActionType = "SMTP_TOKEN_RETRIEVED"
	HydratedCriticalActionTypeTeamAdded                                  HydratedCriticalActionType = "TEAM_ADDED"
	HydratedCriticalActionTypeTeamDeleted                                HydratedCriticalActionType = "TEAM_DELETED"
	HydratedCriticalActionTypeTeamUserAdded                              HydratedCriticalActionType = "TEAM_USER_ADDED"
	HydratedCriticalActionTypeTeamUserDeleted                            HydratedCriticalActionType = "TEAM_USER_DELETED"
	HydratedCriticalActionTypeTemplateDeleted                            HydratedCriticalActionType = "TEMPLATE_DELETED"
	HydratedCriticalActionTypeTemplateModified                           HydratedCriticalActionType = "TEMPLATE_MODIFIED"
	HydratedCriticalActionTypeTouchlessPurchase                          HydratedCriticalActionType = "TOUCHLESS_PURCHASE"
	HydratedCriticalActionTypeUnifiedRestoreUndoExecution                HydratedCriticalActionType = "UNIFIED_RESTORE_UNDO_EXECUTION"
	HydratedCriticalActionTypeUninstallIntegration                       HydratedCriticalActionType = "UNINSTALL_INTEGRATION"
	HydratedCriticalActionTypeUnrequireSingleSignOn                      HydratedCriticalActionType = "UNREQUIRE_SINGLE_SIGN_ON"
	HydratedCriticalActionTypeWebhookSettingsUpdate                      HydratedCriticalActionType = "WEBHOOK_SETTINGS_UPDATE"
	HydratedCriticalActionTypeWebhookSubscriptionCreate                  HydratedCriticalActionType = "WEBHOOK_SUBSCRIPTION_CREATE"
	HydratedCriticalActionTypeWebhookSubscriptionUpdate                  HydratedCriticalActionType = "WEBHOOK_SUBSCRIPTION_UPDATE"
)

type PublicAPIUserActionEvent struct {
	// The login activity's unique ID.
	ID         string     `json:"id" api:"required"`
	ActingUser ActingUser `json:"actingUser" api:"required"`
	// The type of action taken.
	Action string `json:"action" api:"required"`
	// The category of the activity.
	Category string `json:"category" api:"required"`
	// The time that the action occurred at.
	OccurredAt time.Time `json:"occurredAt" api:"required" format:"date-time"`
	// The subcategory of the activity.
	SubCategory string `json:"subCategory"`
	// The ID of the impacted object.
	TargetObjectID string `json:"targetObjectId"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID             respjson.Field
		ActingUser     respjson.Field
		Action         respjson.Field
		Category       respjson.Field
		OccurredAt     respjson.Field
		SubCategory    respjson.Field
		TargetObjectID respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PublicAPIUserActionEvent) RawJSON() string { return r.JSON.raw }
func (r *PublicAPIUserActionEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PublicLoginAudit struct {
	// The login activity's unique ID.
	ID string `json:"id" api:"required"`
	// The time the login took place.
	LoginAt time.Time `json:"loginAt" api:"required" format:"date-time"`
	// Whether the login was successful or not.
	LoginSucceeded bool `json:"loginSucceeded" api:"required"`
	// The approximate country code of the login
	CountryCode string `json:"countryCode"`
	// Email address of the user associated with the login.
	Email string `json:"email"`
	// IP address where the activity originated.
	IPAddress string `json:"ipAddress"`
	// The approximate location where the login activity originated.
	Location string `json:"location"`
	// The approximate region code of the login
	RegionCode string `json:"regionCode"`
	// Information about the device used for logging in.
	UserAgent string `json:"userAgent"`
	// The user's unique ID.
	UserID int64 `json:"userId"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID             respjson.Field
		LoginAt        respjson.Field
		LoginSucceeded respjson.Field
		CountryCode    respjson.Field
		Email          respjson.Field
		IPAddress      respjson.Field
		Location       respjson.Field
		RegionCode     respjson.Field
		UserAgent      respjson.Field
		UserID         respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PublicLoginAudit) RawJSON() string { return r.JSON.raw }
func (r *PublicLoginAudit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ActivityListAuditLogsParams struct {
	// The paging cursor token of the last successfully read resource will be returned
	// as the `paging.next.after` JSON property of a paged response containing more
	// results.
	After              param.Opt[string] `query:"after,omitzero" json:"-"`
	FillFinalTimestamp param.Opt[bool]   `query:"fillFinalTimestamp,omitzero" json:"-"`
	// The maximum number of results to display per page.
	Limit          param.Opt[int64]     `query:"limit,omitzero" json:"-"`
	OccurredAfter  param.Opt[time.Time] `query:"occurredAfter,omitzero" format:"date-time" json:"-"`
	OccurredBefore param.Opt[time.Time] `query:"occurredBefore,omitzero" format:"date-time" json:"-"`
	ActingUserID   []int64              `query:"actingUserId,omitzero" json:"-"`
	Sort           []string             `query:"sort,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ActivityListAuditLogsParams]'s query parameters as
// `url.Values`.
func (r ActivityListAuditLogsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ActivityListLoginActivitiesParams struct {
	// The paging cursor token of the last successfully read resource will be returned
	// as the `paging.next.after` JSON property of a paged response containing more
	// results.
	After param.Opt[string] `query:"after,omitzero" json:"-"`
	// The maximum number of results to display per page.
	Limit  param.Opt[int64] `query:"limit,omitzero" json:"-"`
	UserID param.Opt[int64] `query:"userId,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ActivityListLoginActivitiesParams]'s query parameters as
// `url.Values`.
func (r ActivityListLoginActivitiesParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ActivityListSecurityActivitiesParams struct {
	// The paging cursor token of the last successfully read resource will be returned
	// as the `paging.next.after` JSON property of a paged response containing more
	// results.
	After         param.Opt[string] `query:"after,omitzero" json:"-"`
	FromTimestamp param.Opt[int64]  `query:"fromTimestamp,omitzero" json:"-"`
	// The maximum number of results to display per page.
	Limit       param.Opt[int64] `query:"limit,omitzero" json:"-"`
	ToTimestamp param.Opt[int64] `query:"toTimestamp,omitzero" json:"-"`
	UserID      param.Opt[int64] `query:"userId,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ActivityListSecurityActivitiesParams]'s query parameters as
// `url.Values`.
func (r ActivityListSecurityActivitiesParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
