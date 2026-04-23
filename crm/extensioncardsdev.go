// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package crm

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/stainless-sdks/hubspot-sdk-go/internal/apijson"
	shimjson "github.com/stainless-sdks/hubspot-sdk-go/internal/encoding/json"
	"github.com/stainless-sdks/hubspot-sdk-go/internal/requestconfig"
	"github.com/stainless-sdks/hubspot-sdk-go/option"
	"github.com/stainless-sdks/hubspot-sdk-go/packages/param"
	"github.com/stainless-sdks/hubspot-sdk-go/packages/respjson"
)

// ExtensionCardsDevService contains methods and other services that help with
// interacting with the hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewExtensionCardsDevService] method instead.
type ExtensionCardsDevService struct {
	options []option.RequestOption
}

// NewExtensionCardsDevService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewExtensionCardsDevService(opts ...option.RequestOption) (r ExtensionCardsDevService) {
	r = ExtensionCardsDevService{}
	r.options = opts
	return
}

// Defines a new card that will become active on an account when this app is
// installed.
func (r *ExtensionCardsDevService) New(ctx context.Context, appID int64, body ExtensionCardsDevNewParams, opts ...option.RequestOption) (res *PublicCardResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := fmt.Sprintf("crm/extensions/cards-dev/2026-03/%v", appID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Update a card definition with new details.
func (r *ExtensionCardsDevService) Update(ctx context.Context, cardID string, params ExtensionCardsDevUpdateParams, opts ...option.RequestOption) (res *PublicCardResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if cardID == "" {
		err = errors.New("missing required cardId parameter")
		return nil, err
	}
	path := fmt.Sprintf("crm/extensions/cards-dev/2026-03/%v/%s", params.AppID, url.PathEscape(cardID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return res, err
}

// Permanently deletes a card definition with the given ID. Once deleted, data
// fetch requests for this card will no longer be sent to your service. This can't
// be undone.
func (r *ExtensionCardsDevService) Delete(ctx context.Context, cardID string, body ExtensionCardsDevDeleteParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if cardID == "" {
		err = errors.New("missing required cardId parameter")
		return err
	}
	path := fmt.Sprintf("crm/extensions/cards-dev/2026-03/%v/%s", body.AppID, url.PathEscape(cardID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Returns a list of cards for a given app.
func (r *ExtensionCardsDevService) Get(ctx context.Context, appID int64, opts ...option.RequestOption) (res *PublicCardListResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := fmt.Sprintf("crm/extensions/cards-dev/2026-03/%v", appID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Returns the definition for a card with the given ID.
func (r *ExtensionCardsDevService) GetByID(ctx context.Context, cardID string, query ExtensionCardsDevGetByIDParams, opts ...option.RequestOption) (res *PublicCardResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if cardID == "" {
		err = errors.New("missing required cardId parameter")
		return nil, err
	}
	path := fmt.Sprintf("crm/extensions/cards-dev/2026-03/%v/%s", query.AppID, url.PathEscape(cardID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Returns an example card detail response. This is the payload with displayed
// details for a card that will be shown to a user. An app should send this in
// response to the data fetch request.
func (r *ExtensionCardsDevService) GetSampleResponse(ctx context.Context, opts ...option.RequestOption) (res *IntegratorCardPayloadResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "crm/extensions/cards-dev/2026-03/sample-response"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

func (r *ExtensionCardsDevService) MigrateViews(ctx context.Context, appID int64, body ExtensionCardsDevMigrateViewsParams, opts ...option.RequestOption) (res *CardMigrateViewsResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := fmt.Sprintf("crm/extensions/cards-dev/2026-03/%v/views/migrate", appID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type ActionConfirmationBody struct {
	// The label for the button that cancels the action.
	CancelButtonLabel string `json:"cancelButtonLabel" api:"required"`
	// The label for the button that confirms the action.
	ConfirmButtonLabel string `json:"confirmButtonLabel" api:"required"`
	// The message displayed to the user to confirm the action.
	Prompt string `json:"prompt" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CancelButtonLabel  respjson.Field
		ConfirmButtonLabel respjson.Field
		Prompt             respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ActionConfirmationBody) RawJSON() string { return r.JSON.raw }
func (r *ActionConfirmationBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ActionHookActionBody struct {
	// The HTTP method to be used when making the call, which can be set to GET, POST,
	// PUT, DELETE, or PATCH. If using GET or DELETE
	//
	// Any of "CONNECT", "DELETE", "GET", "HEAD", "OPTIONS", "PATCH", "POST", "PUT",
	// "TRACE".
	HTTPMethod ActionHookActionBodyHTTPMethod `json:"httpMethod" api:"required"`
	// A list of property names that will be included on the action. See the
	// documentation for more information
	PropertyNamesIncluded []string `json:"propertyNamesIncluded" api:"required"`
	// The type of status.
	//
	// Any of "ACTION_HOOK".
	Type ActionHookActionBodyType `json:"type" api:"required"`
	// The URL endpoint that will be called when the action is triggered.
	URL          string                 `json:"url" api:"required"`
	Confirmation ActionConfirmationBody `json:"confirmation"`
	// The label for this property as you'd like it displayed to users.
	Label string `json:"label"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		HTTPMethod            respjson.Field
		PropertyNamesIncluded respjson.Field
		Type                  respjson.Field
		URL                   respjson.Field
		Confirmation          respjson.Field
		Label                 respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ActionHookActionBody) RawJSON() string { return r.JSON.raw }
func (r *ActionHookActionBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The HTTP method to be used when making the call, which can be set to GET, POST,
// PUT, DELETE, or PATCH. If using GET or DELETE
type ActionHookActionBodyHTTPMethod string

const (
	ActionHookActionBodyHTTPMethodConnect ActionHookActionBodyHTTPMethod = "CONNECT"
	ActionHookActionBodyHTTPMethodDelete  ActionHookActionBodyHTTPMethod = "DELETE"
	ActionHookActionBodyHTTPMethodGet     ActionHookActionBodyHTTPMethod = "GET"
	ActionHookActionBodyHTTPMethodHead    ActionHookActionBodyHTTPMethod = "HEAD"
	ActionHookActionBodyHTTPMethodOptions ActionHookActionBodyHTTPMethod = "OPTIONS"
	ActionHookActionBodyHTTPMethodPatch   ActionHookActionBodyHTTPMethod = "PATCH"
	ActionHookActionBodyHTTPMethodPost    ActionHookActionBodyHTTPMethod = "POST"
	ActionHookActionBodyHTTPMethodPut     ActionHookActionBodyHTTPMethod = "PUT"
	ActionHookActionBodyHTTPMethodTrace   ActionHookActionBodyHTTPMethod = "TRACE"
)

// The type of status.
type ActionHookActionBodyType string

const (
	ActionHookActionBodyTypeActionHook ActionHookActionBodyType = "ACTION_HOOK"
)

type CardActions struct {
	// A list of URL prefixes that will be accepted for card action URLs. If your data
	// fetch response includes an action URL that doesn't begin with one of these
	// values, it will result in an error and the card will not be displayed.
	BaseURLs []string `json:"baseUrls" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BaseURLs    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CardActions) RawJSON() string { return r.JSON.raw }
func (r *CardActions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this CardActions to a CardActionsParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// CardActionsParam.Overrides()
func (r CardActions) ToParam() CardActionsParam {
	return param.Override[CardActionsParam](json.RawMessage(r.RawJSON()))
}

// The property BaseURLs is required.
type CardActionsParam struct {
	// A list of URL prefixes that will be accepted for card action URLs. If your data
	// fetch response includes an action URL that doesn't begin with one of these
	// values, it will result in an error and the card will not be displayed.
	BaseURLs []string `json:"baseUrls,omitzero" api:"required"`
	paramObj
}

func (r CardActionsParam) MarshalJSON() (data []byte, err error) {
	type shadow CardActionsParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CardActionsParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CardAuditResponse struct {
	// The type of action performed, with possible values: CREATE, DELETE, UPDATE.
	//
	// Any of "CREATE", "DELETE", "UPDATE".
	ActionType CardAuditResponseActionType `json:"actionType" api:"required"`
	// The ID of the application associated with the card.
	ApplicationID int64 `json:"applicationId" api:"required"`
	// The source of authentication for the action, with possible values: APP,
	// EXTERNAL, INTERNAL.
	//
	// Any of "APP", "EXTERNAL", "INTERNAL".
	AuthSource CardAuditResponseAuthSource `json:"authSource" api:"required"`
	// The timestamp indicating when the change occurred.
	ChangedAt int64 `json:"changedAt" api:"required"`
	// The ID of the user who initiated the action.
	InitiatingUserID int64 `json:"initiatingUserId" api:"required"`
	// The ID of the card.
	ObjectTypeID int64 `json:"objectTypeId" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ActionType       respjson.Field
		ApplicationID    respjson.Field
		AuthSource       respjson.Field
		ChangedAt        respjson.Field
		InitiatingUserID respjson.Field
		ObjectTypeID     respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CardAuditResponse) RawJSON() string { return r.JSON.raw }
func (r *CardAuditResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The type of action performed, with possible values: CREATE, DELETE, UPDATE.
type CardAuditResponseActionType string

const (
	CardAuditResponseActionTypeCreate CardAuditResponseActionType = "CREATE"
	CardAuditResponseActionTypeDelete CardAuditResponseActionType = "DELETE"
	CardAuditResponseActionTypeUpdate CardAuditResponseActionType = "UPDATE"
)

// The source of authentication for the action, with possible values: APP,
// EXTERNAL, INTERNAL.
type CardAuditResponseAuthSource string

const (
	CardAuditResponseAuthSourceApp      CardAuditResponseAuthSource = "APP"
	CardAuditResponseAuthSourceExternal CardAuditResponseAuthSource = "EXTERNAL"
	CardAuditResponseAuthSourceInternal CardAuditResponseAuthSource = "INTERNAL"
)

// The properties Actions, Display, Fetch, Title are required.
type CardCreateRequestParam struct {
	Actions CardActionsParam     `json:"actions,omitzero" api:"required"`
	Display CardDisplayBodyParam `json:"display,omitzero" api:"required"`
	Fetch   CardFetchBodyParam   `json:"fetch,omitzero" api:"required"`
	// The top-level title for this card. Displayed to users in the CRM UI.
	Title string `json:"title" api:"required"`
	paramObj
}

func (r CardCreateRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow CardCreateRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CardCreateRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CardDisplayBody struct {
	// Card display properties. These will will be rendered as "label : value" pairs in
	// the card UI. See the [example card](#) in the overview docs for more details.
	Properties []CardDisplayProperty `json:"properties" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Properties  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CardDisplayBody) RawJSON() string { return r.JSON.raw }
func (r *CardDisplayBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this CardDisplayBody to a CardDisplayBodyParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// CardDisplayBodyParam.Overrides()
func (r CardDisplayBody) ToParam() CardDisplayBodyParam {
	return param.Override[CardDisplayBodyParam](json.RawMessage(r.RawJSON()))
}

// The property Properties is required.
type CardDisplayBodyParam struct {
	// Card display properties. These will will be rendered as "label : value" pairs in
	// the card UI. See the [example card](#) in the overview docs for more details.
	Properties []CardDisplayPropertyParam `json:"properties,omitzero" api:"required"`
	paramObj
}

func (r CardDisplayBodyParam) MarshalJSON() (data []byte, err error) {
	type shadow CardDisplayBodyParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CardDisplayBodyParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CardDisplayProperty struct {
	// Type of data represented by this property.
	//
	// Any of "BOOLEAN", "CURRENCY", "DATE", "DATETIME", "EMAIL", "LINK", "NUMERIC",
	// "STATUS", "STRING".
	DataType CardDisplayPropertyDataType `json:"dataType" api:"required"`
	// The label for this property as you'd like it displayed to users.
	Label string `json:"label" api:"required"`
	// An internal identifier for this property. This value must be unique TODO.
	Name string `json:"name" api:"required"`
	// An array of available options that can be displayed. Only used in when
	// `dataType` is `STATUS`.
	Options []DisplayOption `json:"options" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DataType    respjson.Field
		Label       respjson.Field
		Name        respjson.Field
		Options     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CardDisplayProperty) RawJSON() string { return r.JSON.raw }
func (r *CardDisplayProperty) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this CardDisplayProperty to a CardDisplayPropertyParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// CardDisplayPropertyParam.Overrides()
func (r CardDisplayProperty) ToParam() CardDisplayPropertyParam {
	return param.Override[CardDisplayPropertyParam](json.RawMessage(r.RawJSON()))
}

// Type of data represented by this property.
type CardDisplayPropertyDataType string

const (
	CardDisplayPropertyDataTypeBoolean  CardDisplayPropertyDataType = "BOOLEAN"
	CardDisplayPropertyDataTypeCurrency CardDisplayPropertyDataType = "CURRENCY"
	CardDisplayPropertyDataTypeDate     CardDisplayPropertyDataType = "DATE"
	CardDisplayPropertyDataTypeDatetime CardDisplayPropertyDataType = "DATETIME"
	CardDisplayPropertyDataTypeEmail    CardDisplayPropertyDataType = "EMAIL"
	CardDisplayPropertyDataTypeLink     CardDisplayPropertyDataType = "LINK"
	CardDisplayPropertyDataTypeNumeric  CardDisplayPropertyDataType = "NUMERIC"
	CardDisplayPropertyDataTypeStatus   CardDisplayPropertyDataType = "STATUS"
	CardDisplayPropertyDataTypeString   CardDisplayPropertyDataType = "STRING"
)

// The properties DataType, Label, Name, Options are required.
type CardDisplayPropertyParam struct {
	// Type of data represented by this property.
	//
	// Any of "BOOLEAN", "CURRENCY", "DATE", "DATETIME", "EMAIL", "LINK", "NUMERIC",
	// "STATUS", "STRING".
	DataType CardDisplayPropertyDataType `json:"dataType,omitzero" api:"required"`
	// The label for this property as you'd like it displayed to users.
	Label string `json:"label" api:"required"`
	// An internal identifier for this property. This value must be unique TODO.
	Name string `json:"name" api:"required"`
	// An array of available options that can be displayed. Only used in when
	// `dataType` is `STATUS`.
	Options []DisplayOptionParam `json:"options,omitzero" api:"required"`
	paramObj
}

func (r CardDisplayPropertyParam) MarshalJSON() (data []byte, err error) {
	type shadow CardDisplayPropertyParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CardDisplayPropertyParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties CardType, ObjectTypes, TargetURL are required.
type CardFetchBodyParam struct {
	// A deprecated field to determine the type of card returned.
	//
	// Any of "EXTERNAL", "SERVERLESS".
	CardType CardFetchBodyCardType `json:"cardType,omitzero" api:"required"`
	// An array of CRM object types where this card should be displayed. HubSpot will
	// call your data fetch URL whenever a user visits a record page of the types
	// defined here.
	ObjectTypes []CardObjectTypeBodyParam `json:"objectTypes,omitzero" api:"required"`
	// URL to a service endpoints that will respond with card details. HubSpot will
	// call this endpoint each time a user visits a CRM record page where this card
	// should be displayed.
	TargetURL string `json:"targetUrl" api:"required"`
	// A deprecated field to specify serverless functionality with the card
	ServerlessFunction param.Opt[string] `json:"serverlessFunction,omitzero"`
	paramObj
}

func (r CardFetchBodyParam) MarshalJSON() (data []byte, err error) {
	type shadow CardFetchBodyParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CardFetchBodyParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A deprecated field to determine the type of card returned.
type CardFetchBodyCardType string

const (
	CardFetchBodyCardTypeExternal   CardFetchBodyCardType = "EXTERNAL"
	CardFetchBodyCardTypeServerless CardFetchBodyCardType = "SERVERLESS"
)

// The property ObjectTypes is required.
type CardFetchBodyPatchParam struct {
	// An array of CRM object types where this card should be displayed. HubSpot will
	// call your target URL whenever a user visits a record page of the types defined
	// here.
	ObjectTypes []CardObjectTypeBodyParam `json:"objectTypes,omitzero" api:"required"`
	// A deprecated field to specify serverless functionality with the card
	ServerlessFunction param.Opt[string] `json:"serverlessFunction,omitzero"`
	// URL to a service endpoint that will respond with details for this card. HubSpot
	// will call this endpoint each time a user visits a CRM record page where this
	// card should be displayed.
	TargetURL param.Opt[string] `json:"targetUrl,omitzero"`
	// A deprecated field to determine the type of card returned.
	//
	// Any of "EXTERNAL", "SERVERLESS".
	CardType CardFetchBodyPatchCardType `json:"cardType,omitzero"`
	paramObj
}

func (r CardFetchBodyPatchParam) MarshalJSON() (data []byte, err error) {
	type shadow CardFetchBodyPatchParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CardFetchBodyPatchParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A deprecated field to determine the type of card returned.
type CardFetchBodyPatchCardType string

const (
	CardFetchBodyPatchCardTypeExternal   CardFetchBodyPatchCardType = "EXTERNAL"
	CardFetchBodyPatchCardTypeServerless CardFetchBodyPatchCardType = "SERVERLESS"
)

// The properties AllowDuplicateAppCardIDs, AppCardID, LegacyCrmCardID are
// required.
type CardMigrateViewsRequestParam struct {
	AllowDuplicateAppCardIDs bool             `json:"allowDuplicateAppCardIds" api:"required"`
	AppCardID                int64            `json:"appCardId" api:"required"`
	LegacyCrmCardID          int64            `json:"legacyCrmCardId" api:"required"`
	HelpdeskAppCardID        param.Opt[int64] `json:"helpdeskAppCardId,omitzero"`
	paramObj
}

func (r CardMigrateViewsRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow CardMigrateViewsRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CardMigrateViewsRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CardMigrateViewsResponse struct {
	// A human readable message describing the error along with remediation steps where
	// appropriate
	Message string `json:"message" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Message     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CardMigrateViewsResponse) RawJSON() string { return r.JSON.raw }
func (r *CardMigrateViewsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CardObjectTypeBody struct {
	// A CRM object type where this card should be displayed.
	//
	// Any of "companies", "contacts", "deals", "marketing_events", "tickets".
	Name CardObjectTypeBodyName `json:"name" api:"required"`
	// An array of properties that should be sent to this card's target URL when the
	// data fetch request is made. Must be valid properties for the corresponding CRM
	// object type.
	PropertiesToSend []string `json:"propertiesToSend" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name             respjson.Field
		PropertiesToSend respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CardObjectTypeBody) RawJSON() string { return r.JSON.raw }
func (r *CardObjectTypeBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this CardObjectTypeBody to a CardObjectTypeBodyParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// CardObjectTypeBodyParam.Overrides()
func (r CardObjectTypeBody) ToParam() CardObjectTypeBodyParam {
	return param.Override[CardObjectTypeBodyParam](json.RawMessage(r.RawJSON()))
}

// A CRM object type where this card should be displayed.
type CardObjectTypeBodyName string

const (
	CardObjectTypeBodyNameCompanies       CardObjectTypeBodyName = "companies"
	CardObjectTypeBodyNameContacts        CardObjectTypeBodyName = "contacts"
	CardObjectTypeBodyNameDeals           CardObjectTypeBodyName = "deals"
	CardObjectTypeBodyNameMarketingEvents CardObjectTypeBodyName = "marketing_events"
	CardObjectTypeBodyNameTickets         CardObjectTypeBodyName = "tickets"
)

// The properties Name, PropertiesToSend are required.
type CardObjectTypeBodyParam struct {
	// A CRM object type where this card should be displayed.
	//
	// Any of "companies", "contacts", "deals", "marketing_events", "tickets".
	Name CardObjectTypeBodyName `json:"name,omitzero" api:"required"`
	// An array of properties that should be sent to this card's target URL when the
	// data fetch request is made. Must be valid properties for the corresponding CRM
	// object type.
	PropertiesToSend []string `json:"propertiesToSend,omitzero" api:"required"`
	paramObj
}

func (r CardObjectTypeBodyParam) MarshalJSON() (data []byte, err error) {
	type shadow CardObjectTypeBodyParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CardObjectTypeBodyParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CardPatchRequestParam struct {
	// The top-level title for this card. Displayed to users in the CRM UI.
	Title   param.Opt[string]       `json:"title,omitzero"`
	Actions CardActionsParam        `json:"actions,omitzero"`
	Display CardDisplayBodyParam    `json:"display,omitzero"`
	Fetch   CardFetchBodyPatchParam `json:"fetch,omitzero"`
	paramObj
}

func (r CardPatchRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow CardPatchRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CardPatchRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DisplayOption struct {
	// The text that will be displayed to users for this option.
	Label string `json:"label" api:"required"`
	// JSON-friendly unique name for option.
	Name string `json:"name" api:"required"`
	// The type of status.
	//
	// Any of "DANGER", "DEFAULT", "INFO", "SUCCESS", "WARNING".
	Type DisplayOptionType `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Label       respjson.Field
		Name        respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DisplayOption) RawJSON() string { return r.JSON.raw }
func (r *DisplayOption) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this DisplayOption to a DisplayOptionParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// DisplayOptionParam.Overrides()
func (r DisplayOption) ToParam() DisplayOptionParam {
	return param.Override[DisplayOptionParam](json.RawMessage(r.RawJSON()))
}

// The type of status.
type DisplayOptionType string

const (
	DisplayOptionTypeDanger  DisplayOptionType = "DANGER"
	DisplayOptionTypeDefault DisplayOptionType = "DEFAULT"
	DisplayOptionTypeInfo    DisplayOptionType = "INFO"
	DisplayOptionTypeSuccess DisplayOptionType = "SUCCESS"
	DisplayOptionTypeWarning DisplayOptionType = "WARNING"
)

// The properties Label, Name, Type are required.
type DisplayOptionParam struct {
	// The text that will be displayed to users for this option.
	Label string `json:"label" api:"required"`
	// JSON-friendly unique name for option.
	Name string `json:"name" api:"required"`
	// The type of status.
	//
	// Any of "DANGER", "DEFAULT", "INFO", "SUCCESS", "WARNING".
	Type DisplayOptionType `json:"type,omitzero" api:"required"`
	paramObj
}

func (r DisplayOptionParam) MarshalJSON() (data []byte, err error) {
	type shadow DisplayOptionParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DisplayOptionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IFrameActionBody struct {
	// The height of the iframe in pixels.
	Height int64 `json:"height" api:"required"`
	// A list of property names that will be included on the url of the iframe.
	PropertyNamesIncluded []string `json:"propertyNamesIncluded" api:"required"`
	// The type of status.
	//
	// Any of "IFRAME".
	Type IFrameActionBodyType `json:"type" api:"required"`
	// The URL endpoint that will be loaded in the iframe when triggered.
	URL string `json:"url" api:"required"`
	// The width of the iframe in pixels.
	Width int64 `json:"width" api:"required"`
	// The label for this property as you'd like it displayed to users.
	Label string `json:"label"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Height                respjson.Field
		PropertyNamesIncluded respjson.Field
		Type                  respjson.Field
		URL                   respjson.Field
		Width                 respjson.Field
		Label                 respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IFrameActionBody) RawJSON() string { return r.JSON.raw }
func (r *IFrameActionBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The type of status.
type IFrameActionBodyType string

const (
	IFrameActionBodyTypeIframe IFrameActionBodyType = "IFRAME"
)

type IntegratorCardPayloadResponse struct {
	// The number version of the response.
	//
	// Any of "v1", "v3".
	ResponseVersion IntegratorCardPayloadResponseResponseVersion `json:"responseVersion" api:"required"`
	// A list of up to five valid card sub categories.
	Sections []IntegratorObjectResult `json:"sections" api:"required"`
	// The total number of card properties that will be sent in this response.
	TotalCount int64 `json:"totalCount" api:"required"`
	// URL to a page the integrator has built that displays all details for this card.
	// This URL will be displayed to users under a `See more [x]` link if there are
	// more than five items in your response, where `[x]` is the value of `itemLabel`.
	AllItemsLinkURL string `json:"allItemsLinkUrl"`
	// The label to be used for the `allItemsLinkUrl` link (e.g. 'See more tickets').
	// If not provided, this falls back to the card's title.
	CardLabel       string          `json:"cardLabel"`
	TopLevelActions TopLevelActions `json:"topLevelActions"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ResponseVersion respjson.Field
		Sections        respjson.Field
		TotalCount      respjson.Field
		AllItemsLinkURL respjson.Field
		CardLabel       respjson.Field
		TopLevelActions respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IntegratorCardPayloadResponse) RawJSON() string { return r.JSON.raw }
func (r *IntegratorCardPayloadResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The number version of the response.
type IntegratorCardPayloadResponseResponseVersion string

const (
	IntegratorCardPayloadResponseResponseVersionV1 IntegratorCardPayloadResponseResponseVersion = "v1"
	IntegratorCardPayloadResponseResponseVersionV3 IntegratorCardPayloadResponseResponseVersion = "v3"
)

type IntegratorObjectResult struct {
	// The unique identifier for the card.
	ID string `json:"id" api:"required"`
	// A list of actions associated with the card, which can include action hooks,
	// confirmation action hooks, or iframes.
	Actions []IntegratorObjectResultActionUnion `json:"actions" api:"required"`
	// The top-level title for this card. Displayed to users in the CRM UI.
	Title string `json:"title" api:"required"`
	// A collection of tokens representing specific properties related to the card.
	Tokens []ObjectToken `json:"tokens" api:"required"`
	// A URL used on the title of the card
	LinkURL string `json:"linkUrl"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Actions     respjson.Field
		Title       respjson.Field
		Tokens      respjson.Field
		LinkURL     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IntegratorObjectResult) RawJSON() string { return r.JSON.raw }
func (r *IntegratorObjectResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// IntegratorObjectResultActionUnion contains all possible properties and values
// from [ActionHookActionBody], [IFrameActionBody].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type IntegratorObjectResultActionUnion struct {
	// This field is from variant [ActionHookActionBody].
	HTTPMethod            ActionHookActionBodyHTTPMethod `json:"httpMethod"`
	PropertyNamesIncluded []string                       `json:"propertyNamesIncluded"`
	Type                  string                         `json:"type"`
	URL                   string                         `json:"url"`
	// This field is from variant [ActionHookActionBody].
	Confirmation ActionConfirmationBody `json:"confirmation"`
	Label        string                 `json:"label"`
	// This field is from variant [IFrameActionBody].
	Height int64 `json:"height"`
	// This field is from variant [IFrameActionBody].
	Width int64 `json:"width"`
	JSON  struct {
		HTTPMethod            respjson.Field
		PropertyNamesIncluded respjson.Field
		Type                  respjson.Field
		URL                   respjson.Field
		Confirmation          respjson.Field
		Label                 respjson.Field
		Height                respjson.Field
		Width                 respjson.Field
		raw                   string
	} `json:"-"`
}

func (u IntegratorObjectResultActionUnion) AsActionHook() (v ActionHookActionBody) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u IntegratorObjectResultActionUnion) AsIframe() (v IFrameActionBody) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u IntegratorObjectResultActionUnion) RawJSON() string { return u.JSON.raw }

func (r *IntegratorObjectResultActionUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ObjectToken struct {
	// The value of the property
	Value string `json:"value" api:"required"`
	// Type of data represented by this property.
	//
	// Any of "BOOLEAN", "CURRENCY", "DATE", "DATETIME", "EMAIL", "LINK", "NUMERIC",
	// "STATUS", "STRING".
	DataType ObjectTokenDataType `json:"dataType"`
	// The label for this property as you'd like it displayed to users.
	Label string `json:"label"`
	// An internal identifier for this property. This value must be unique TODO.
	Name string `json:"name"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Value       respjson.Field
		DataType    respjson.Field
		Label       respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ObjectToken) RawJSON() string { return r.JSON.raw }
func (r *ObjectToken) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Type of data represented by this property.
type ObjectTokenDataType string

const (
	ObjectTokenDataTypeBoolean  ObjectTokenDataType = "BOOLEAN"
	ObjectTokenDataTypeCurrency ObjectTokenDataType = "CURRENCY"
	ObjectTokenDataTypeDate     ObjectTokenDataType = "DATE"
	ObjectTokenDataTypeDatetime ObjectTokenDataType = "DATETIME"
	ObjectTokenDataTypeEmail    ObjectTokenDataType = "EMAIL"
	ObjectTokenDataTypeLink     ObjectTokenDataType = "LINK"
	ObjectTokenDataTypeNumeric  ObjectTokenDataType = "NUMERIC"
	ObjectTokenDataTypeStatus   ObjectTokenDataType = "STATUS"
	ObjectTokenDataTypeString   ObjectTokenDataType = "STRING"
)

type PublicCardFetchBody struct {
	// An array of CRM object types where this card should be displayed. HubSpot will
	// call your target URL whenever a user visits a record page of the types defined
	// here.
	ObjectTypes []CardObjectTypeBody `json:"objectTypes" api:"required"`
	// URL to a service endpoint that will respond with details for this card. HubSpot
	// will call this endpoint each time a user visits a CRM record page where this
	// card should be displayed.
	TargetURL string `json:"targetUrl" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ObjectTypes respjson.Field
		TargetURL   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PublicCardFetchBody) RawJSON() string { return r.JSON.raw }
func (r *PublicCardFetchBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PublicCardListResponse struct {
	// A list of card responses
	Results []PublicCardResponse `json:"results" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Results     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PublicCardListResponse) RawJSON() string { return r.JSON.raw }
func (r *PublicCardListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PublicCardResponse struct {
	// The unique id of the card.
	ID      string      `json:"id" api:"required"`
	Actions CardActions `json:"actions" api:"required"`
	// A list of actions performed on the card, including creation, deletion, and
	// updates.
	AuditHistory []CardAuditResponse `json:"auditHistory" api:"required"`
	Display      CardDisplayBody     `json:"display" api:"required"`
	Fetch        PublicCardFetchBody `json:"fetch" api:"required"`
	// The top-level title for this card. Displayed to users in the CRM UI.
	Title string `json:"title" api:"required"`
	// The date and time when the card was created.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// The date and time when the card was last updated.
	UpdatedAt time.Time `json:"updatedAt" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		Actions      respjson.Field
		AuditHistory respjson.Field
		Display      respjson.Field
		Fetch        respjson.Field
		Title        respjson.Field
		CreatedAt    respjson.Field
		UpdatedAt    respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PublicCardResponse) RawJSON() string { return r.JSON.raw }
func (r *PublicCardResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TopLevelActions struct {
	// Specifies a list of secondary actions for a card, each of which can be an action
	// hook or an iframe.
	Secondary []TopLevelActionsSecondaryUnion `json:"secondary" api:"required"`
	// Defines the primary action for a card, which can be either an action hook or an
	// iframe.
	Primary  TopLevelActionsPrimaryUnion `json:"primary"`
	Settings IFrameActionBody            `json:"settings"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Secondary   respjson.Field
		Primary     respjson.Field
		Settings    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TopLevelActions) RawJSON() string { return r.JSON.raw }
func (r *TopLevelActions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// TopLevelActionsSecondaryUnion contains all possible properties and values from
// [ActionHookActionBody], [IFrameActionBody].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type TopLevelActionsSecondaryUnion struct {
	// This field is from variant [ActionHookActionBody].
	HTTPMethod            ActionHookActionBodyHTTPMethod `json:"httpMethod"`
	PropertyNamesIncluded []string                       `json:"propertyNamesIncluded"`
	Type                  string                         `json:"type"`
	URL                   string                         `json:"url"`
	// This field is from variant [ActionHookActionBody].
	Confirmation ActionConfirmationBody `json:"confirmation"`
	Label        string                 `json:"label"`
	// This field is from variant [IFrameActionBody].
	Height int64 `json:"height"`
	// This field is from variant [IFrameActionBody].
	Width int64 `json:"width"`
	JSON  struct {
		HTTPMethod            respjson.Field
		PropertyNamesIncluded respjson.Field
		Type                  respjson.Field
		URL                   respjson.Field
		Confirmation          respjson.Field
		Label                 respjson.Field
		Height                respjson.Field
		Width                 respjson.Field
		raw                   string
	} `json:"-"`
}

func (u TopLevelActionsSecondaryUnion) AsActionHook() (v ActionHookActionBody) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TopLevelActionsSecondaryUnion) AsIframe() (v IFrameActionBody) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u TopLevelActionsSecondaryUnion) RawJSON() string { return u.JSON.raw }

func (r *TopLevelActionsSecondaryUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// TopLevelActionsPrimaryUnion contains all possible properties and values from
// [ActionHookActionBody], [IFrameActionBody].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type TopLevelActionsPrimaryUnion struct {
	// This field is from variant [ActionHookActionBody].
	HTTPMethod            ActionHookActionBodyHTTPMethod `json:"httpMethod"`
	PropertyNamesIncluded []string                       `json:"propertyNamesIncluded"`
	Type                  string                         `json:"type"`
	URL                   string                         `json:"url"`
	// This field is from variant [ActionHookActionBody].
	Confirmation ActionConfirmationBody `json:"confirmation"`
	Label        string                 `json:"label"`
	// This field is from variant [IFrameActionBody].
	Height int64 `json:"height"`
	// This field is from variant [IFrameActionBody].
	Width int64 `json:"width"`
	JSON  struct {
		HTTPMethod            respjson.Field
		PropertyNamesIncluded respjson.Field
		Type                  respjson.Field
		URL                   respjson.Field
		Confirmation          respjson.Field
		Label                 respjson.Field
		Height                respjson.Field
		Width                 respjson.Field
		raw                   string
	} `json:"-"`
}

func (u TopLevelActionsPrimaryUnion) AsActionHook() (v ActionHookActionBody) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TopLevelActionsPrimaryUnion) AsIframe() (v IFrameActionBody) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u TopLevelActionsPrimaryUnion) RawJSON() string { return u.JSON.raw }

func (r *TopLevelActionsPrimaryUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExtensionCardsDevNewParams struct {
	CardCreateRequest CardCreateRequestParam
	paramObj
}

func (r ExtensionCardsDevNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.CardCreateRequest)
}
func (r *ExtensionCardsDevNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExtensionCardsDevUpdateParams struct {
	AppID            int64 `path:"appId" api:"required" json:"-"`
	CardPatchRequest CardPatchRequestParam
	paramObj
}

func (r ExtensionCardsDevUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.CardPatchRequest)
}
func (r *ExtensionCardsDevUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExtensionCardsDevDeleteParams struct {
	AppID int64 `path:"appId" api:"required" json:"-"`
	paramObj
}

type ExtensionCardsDevGetByIDParams struct {
	AppID int64 `path:"appId" api:"required" json:"-"`
	paramObj
}

type ExtensionCardsDevMigrateViewsParams struct {
	CardMigrateViewsRequest CardMigrateViewsRequestParam
	paramObj
}

func (r ExtensionCardsDevMigrateViewsParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.CardMigrateViewsRequest)
}
func (r *ExtensionCardsDevMigrateViewsParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
