// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package crm

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"slices"
	"time"

	"github.com/stainless-sdks/hubspot-sdk-go/internal/apijson"
	shimjson "github.com/stainless-sdks/hubspot-sdk-go/internal/encoding/json"
	"github.com/stainless-sdks/hubspot-sdk-go/internal/requestconfig"
	"github.com/stainless-sdks/hubspot-sdk-go/option"
	"github.com/stainless-sdks/hubspot-sdk-go/packages/param"
	"github.com/stainless-sdks/hubspot-sdk-go/packages/respjson"
)

// ExtensionCardService contains methods and other services that help with
// interacting with the hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewExtensionCardService] method instead.
type ExtensionCardService struct {
	Options []option.RequestOption
}

// NewExtensionCardService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewExtensionCardService(opts ...option.RequestOption) (r ExtensionCardService) {
	r = ExtensionCardService{}
	r.Options = opts
	return
}

// Defines a new card that will become active on an account when this app is
// installed.
func (r *ExtensionCardService) New(ctx context.Context, appID int64, body ExtensionCardNewParams, opts ...option.RequestOption) (res *PublicCardResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("crm/v3/extensions/cards-dev/%v", appID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Update a card definition with new details.
func (r *ExtensionCardService) Update(ctx context.Context, cardID string, params ExtensionCardUpdateParams, opts ...option.RequestOption) (res *PublicCardResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if cardID == "" {
		err = errors.New("missing required cardId parameter")
		return
	}
	path := fmt.Sprintf("crm/v3/extensions/cards-dev/%v/%s", params.AppID, cardID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return
}

// Returns a list of cards for a given app.
func (r *ExtensionCardService) List(ctx context.Context, appID int64, opts ...option.RequestOption) (res *PublicCardListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("crm/v3/extensions/cards-dev/%v", appID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Permanently deletes a card definition with the given ID. Once deleted, data
// fetch requests for this card will no longer be sent to your service. This can't
// be undone.
func (r *ExtensionCardService) Delete(ctx context.Context, cardID string, body ExtensionCardDeleteParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if cardID == "" {
		err = errors.New("missing required cardId parameter")
		return
	}
	path := fmt.Sprintf("crm/v3/extensions/cards-dev/%v/%s", body.AppID, cardID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Returns the definition for a card with the given ID.
func (r *ExtensionCardService) Get(ctx context.Context, cardID string, query ExtensionCardGetParams, opts ...option.RequestOption) (res *PublicCardResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if cardID == "" {
		err = errors.New("missing required cardId parameter")
		return
	}
	path := fmt.Sprintf("crm/v3/extensions/cards-dev/%v/%s", query.AppID, cardID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Returns an example card detail response. This is the payload with displayed
// details for a card that will be shown to a user. An app should send this in
// response to the data fetch request.
func (r *ExtensionCardService) GetSampleResponse(ctx context.Context, opts ...option.RequestOption) (res *IntegratorCardPayloadResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "crm/v3/extensions/cards-dev/sample-response"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type ActionConfirmationBody struct {
	CancelButtonLabel  string `json:"cancelButtonLabel,required"`
	ConfirmButtonLabel string `json:"confirmButtonLabel,required"`
	Prompt             string `json:"prompt,required"`
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
	// Any of "CONNECT", "DELETE", "GET", "HEAD", "OPTIONS", "PATCH", "POST", "PUT",
	// "TRACE".
	HTTPMethod            ActionHookActionBodyHTTPMethod `json:"httpMethod,required"`
	PropertyNamesIncluded []string                       `json:"propertyNamesIncluded,required"`
	// Any of "ACTION_HOOK".
	Type         ActionHookActionBodyType `json:"type,required"`
	URL          string                   `json:"url,required"`
	Confirmation ActionConfirmationBody   `json:"confirmation"`
	Label        string                   `json:"label"`
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

type ActionHookActionBodyType string

const (
	ActionHookActionBodyTypeActionHook ActionHookActionBodyType = "ACTION_HOOK"
)

// Configuration for custom user actions on cards.
type CardActions struct {
	// A list of URL prefixes that will be accepted for card action URLs. If your data
	// fetch response includes an action URL that doesn't begin with one of these
	// values, it will result in an error and the card will not be displayed.
	BaseURLs []string `json:"baseUrls,required"`
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

// Configuration for custom user actions on cards.
//
// The property BaseURLs is required.
type CardActionsParam struct {
	// A list of URL prefixes that will be accepted for card action URLs. If your data
	// fetch response includes an action URL that doesn't begin with one of these
	// values, it will result in an error and the card will not be displayed.
	BaseURLs []string `json:"baseUrls,omitzero,required"`
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
	// Any of "CREATE", "DELETE", "UPDATE".
	ActionType    CardAuditResponseActionType `json:"actionType,required"`
	ApplicationID int64                       `json:"applicationId,required"`
	// Any of "APP", "EXTERNAL", "INTERNAL".
	AuthSource       CardAuditResponseAuthSource `json:"authSource,required"`
	ChangedAt        int64                       `json:"changedAt,required"`
	InitiatingUserID int64                       `json:"initiatingUserId,required"`
	ObjectTypeID     int64                       `json:"objectTypeId,required"`
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

type CardAuditResponseActionType string

const (
	CardAuditResponseActionTypeCreate CardAuditResponseActionType = "CREATE"
	CardAuditResponseActionTypeDelete CardAuditResponseActionType = "DELETE"
	CardAuditResponseActionTypeUpdate CardAuditResponseActionType = "UPDATE"
)

type CardAuditResponseAuthSource string

const (
	CardAuditResponseAuthSourceApp      CardAuditResponseAuthSource = "APP"
	CardAuditResponseAuthSourceExternal CardAuditResponseAuthSource = "EXTERNAL"
	CardAuditResponseAuthSourceInternal CardAuditResponseAuthSource = "INTERNAL"
)

// State of card definition to be created
//
// The properties Actions, Display, Fetch, Title are required.
type CardCreateRequestParam struct {
	// Configuration for custom user actions on cards.
	Actions CardActionsParam `json:"actions,omitzero,required"`
	// Configuration for displayed info on a card
	Display CardDisplayBodyParam `json:"display,omitzero,required"`
	// Configuration for this card's data fetch request.
	Fetch CardFetchBodyParam `json:"fetch,omitzero,required"`
	// The top-level title for this card. Displayed to users in the CRM UI.
	Title string `json:"title,required"`
	paramObj
}

func (r CardCreateRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow CardCreateRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CardCreateRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Configuration for displayed info on a card
type CardDisplayBody struct {
	// Card display properties. These will will be rendered as "label : value" pairs in
	// the card UI. See the [example card](#) in the overview docs for more details.
	Properties []CardDisplayProperty `json:"properties,required"`
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

// Configuration for displayed info on a card
//
// The property Properties is required.
type CardDisplayBodyParam struct {
	// Card display properties. These will will be rendered as "label : value" pairs in
	// the card UI. See the [example card](#) in the overview docs for more details.
	Properties []CardDisplayPropertyParam `json:"properties,omitzero,required"`
	paramObj
}

func (r CardDisplayBodyParam) MarshalJSON() (data []byte, err error) {
	type shadow CardDisplayBodyParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CardDisplayBodyParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Definition for a card display property.
type CardDisplayProperty struct {
	// Type of data represented by this property.
	//
	// Any of "BOOLEAN", "CURRENCY", "DATE", "DATETIME", "EMAIL", "LINK", "NUMERIC",
	// "STATUS", "STRING".
	DataType CardDisplayPropertyDataType `json:"dataType,required"`
	// The label for this property as you'd like it displayed to users.
	Label string `json:"label,required"`
	// An internal identifier for this property. This value must be unique TODO.
	Name string `json:"name,required"`
	// An array of available options that can be displayed. Only used in when
	// `dataType` is `STATUS`.
	Options []DisplayOption `json:"options,required"`
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

// Definition for a card display property.
//
// The properties DataType, Label, Name, Options are required.
type CardDisplayPropertyParam struct {
	// Type of data represented by this property.
	//
	// Any of "BOOLEAN", "CURRENCY", "DATE", "DATETIME", "EMAIL", "LINK", "NUMERIC",
	// "STATUS", "STRING".
	DataType CardDisplayPropertyDataType `json:"dataType,omitzero,required"`
	// The label for this property as you'd like it displayed to users.
	Label string `json:"label,required"`
	// An internal identifier for this property. This value must be unique TODO.
	Name string `json:"name,required"`
	// An array of available options that can be displayed. Only used in when
	// `dataType` is `STATUS`.
	Options []DisplayOptionParam `json:"options,omitzero,required"`
	paramObj
}

func (r CardDisplayPropertyParam) MarshalJSON() (data []byte, err error) {
	type shadow CardDisplayPropertyParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CardDisplayPropertyParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Configuration for this card's data fetch request.
//
// The properties ObjectTypes, TargetURL are required.
type CardFetchBodyParam struct {
	// An array of CRM object types where this card should be displayed. HubSpot will
	// call your data fetch URL whenever a user visits a record page of the types
	// defined here.
	ObjectTypes []CardObjectTypeBodyParam `json:"objectTypes,omitzero,required"`
	// URL to a service endpoints that will respond with card details. HubSpot will
	// call this endpoint each time a user visits a CRM record page where this card
	// should be displayed.
	TargetURL          string            `json:"targetUrl,required"`
	ServerlessFunction param.Opt[string] `json:"serverlessFunction,omitzero"`
	// Any of "EXTERNAL", "SERVERLESS".
	CardType CardFetchBodyCardType `json:"cardType,omitzero"`
	paramObj
}

func (r CardFetchBodyParam) MarshalJSON() (data []byte, err error) {
	type shadow CardFetchBodyParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CardFetchBodyParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CardFetchBodyCardType string

const (
	CardFetchBodyCardTypeExternal   CardFetchBodyCardType = "EXTERNAL"
	CardFetchBodyCardTypeServerless CardFetchBodyCardType = "SERVERLESS"
)

// Variant of CardFetchBody with fields as optional for patches
//
// The property ObjectTypes is required.
type CardFetchBodyPatchParam struct {
	// An array of CRM object types where this card should be displayed. HubSpot will
	// call your target URL whenever a user visits a record page of the types defined
	// here.
	ObjectTypes        []CardObjectTypeBodyParam `json:"objectTypes,omitzero,required"`
	ServerlessFunction param.Opt[string]         `json:"serverlessFunction,omitzero"`
	// URL to a service endpoint that will respond with details for this card. HubSpot
	// will call this endpoint each time a user visits a CRM record page where this
	// card should be displayed.
	TargetURL param.Opt[string] `json:"targetUrl,omitzero"`
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

type CardFetchBodyPatchCardType string

const (
	CardFetchBodyPatchCardTypeExternal   CardFetchBodyPatchCardType = "EXTERNAL"
	CardFetchBodyPatchCardTypeServerless CardFetchBodyPatchCardType = "SERVERLESS"
)

type CardObjectTypeBody struct {
	// A CRM object type where this card should be displayed.
	//
	// Any of "companies", "contacts", "deals", "marketing_events", "tickets".
	Name CardObjectTypeBodyName `json:"name,required"`
	// An array of properties that should be sent to this card's target URL when the
	// data fetch request is made. Must be valid properties for the corresponding CRM
	// object type.
	PropertiesToSend []string `json:"propertiesToSend,required"`
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
	Name CardObjectTypeBodyName `json:"name,omitzero,required"`
	// An array of properties that should be sent to this card's target URL when the
	// data fetch request is made. Must be valid properties for the corresponding CRM
	// object type.
	PropertiesToSend []string `json:"propertiesToSend,omitzero,required"`
	paramObj
}

func (r CardObjectTypeBodyParam) MarshalJSON() (data []byte, err error) {
	type shadow CardObjectTypeBodyParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CardObjectTypeBodyParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Body for a patch with optional fields
type CardPatchRequestParam struct {
	// The top-level title for this card. Displayed to users in the CRM UI.
	Title param.Opt[string] `json:"title,omitzero"`
	// Configuration for custom user actions on cards.
	Actions CardActionsParam `json:"actions,omitzero"`
	// Configuration for displayed info on a card
	Display CardDisplayBodyParam `json:"display,omitzero"`
	// Variant of CardFetchBody with fields as optional for patches
	Fetch CardFetchBodyPatchParam `json:"fetch,omitzero"`
	paramObj
}

func (r CardPatchRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow CardPatchRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CardPatchRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Option definition for STATUS dataTypes.
type DisplayOption struct {
	// The text that will be displayed to users for this option.
	Label string `json:"label,required"`
	// JSON-friendly unique name for option.
	Name string `json:"name,required"`
	// The type of status.
	//
	// Any of "DANGER", "DEFAULT", "INFO", "SUCCESS", "WARNING".
	Type DisplayOptionType `json:"type,required"`
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

// Option definition for STATUS dataTypes.
//
// The properties Label, Name, Type are required.
type DisplayOptionParam struct {
	// The text that will be displayed to users for this option.
	Label string `json:"label,required"`
	// JSON-friendly unique name for option.
	Name string `json:"name,required"`
	// The type of status.
	//
	// Any of "DANGER", "DEFAULT", "INFO", "SUCCESS", "WARNING".
	Type DisplayOptionType `json:"type,omitzero,required"`
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
	Height                int64    `json:"height,required"`
	PropertyNamesIncluded []string `json:"propertyNamesIncluded,required"`
	// Any of "IFRAME".
	Type  IFrameActionBodyType `json:"type,required"`
	URL   string               `json:"url,required"`
	Width int64                `json:"width,required"`
	Label string               `json:"label"`
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

type IFrameActionBodyType string

const (
	IFrameActionBodyTypeIframe IFrameActionBodyType = "IFRAME"
)

// The card details payload, sent to HubSpot by an app in response to a data fetch
// request when a user visits a CRM record page.
type IntegratorCardPayloadResponse struct {
	// The total number of card properties that will be sent in this response.
	TotalCount int64 `json:"totalCount,required"`
	// URL to a page the integrator has built that displays all details for this card.
	// This URL will be displayed to users under a `See more [x]` link if there are
	// more than five items in your response, where `[x]` is the value of `itemLabel`.
	AllItemsLinkURL string `json:"allItemsLinkUrl"`
	// The label to be used for the `allItemsLinkUrl` link (e.g. 'See more tickets').
	// If not provided, this falls back to the card's title.
	CardLabel string `json:"cardLabel"`
	// Any of "v1", "v3".
	ResponseVersion IntegratorCardPayloadResponseResponseVersion `json:"responseVersion"`
	// A list of up to five valid card sub categories.
	Sections        []IntegratorObjectResult `json:"sections"`
	TopLevelActions TopLevelActions          `json:"topLevelActions"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		TotalCount      respjson.Field
		AllItemsLinkURL respjson.Field
		CardLabel       respjson.Field
		ResponseVersion respjson.Field
		Sections        respjson.Field
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

type IntegratorCardPayloadResponseResponseVersion string

const (
	IntegratorCardPayloadResponseResponseVersionV1 IntegratorCardPayloadResponseResponseVersion = "v1"
	IntegratorCardPayloadResponseResponseVersionV3 IntegratorCardPayloadResponseResponseVersion = "v3"
)

type IntegratorObjectResult struct {
	ID      string                              `json:"id,required"`
	Actions []IntegratorObjectResultActionUnion `json:"actions,required"`
	Title   string                              `json:"title,required"`
	Tokens  []ObjectToken                       `json:"tokens,required"`
	LinkURL string                              `json:"linkUrl"`
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
	Value string `json:"value,required"`
	// Any of "BOOLEAN", "CURRENCY", "DATE", "DATETIME", "EMAIL", "LINK", "NUMERIC",
	// "STATUS", "STRING".
	DataType ObjectTokenDataType `json:"dataType"`
	Label    string              `json:"label"`
	Name     string              `json:"name"`
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
	ObjectTypes []CardObjectTypeBody `json:"objectTypes,required"`
	TargetURL   string               `json:"targetUrl,required"`
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
	Results []PublicCardResponse `json:"results,required"`
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
	ID string `json:"id,required"`
	// Configuration for custom user actions on cards.
	Actions      CardActions         `json:"actions,required"`
	AuditHistory []CardAuditResponse `json:"auditHistory,required"`
	// Configuration for displayed info on a card
	Display   CardDisplayBody     `json:"display,required"`
	Fetch     PublicCardFetchBody `json:"fetch,required"`
	Title     string              `json:"title,required"`
	CreatedAt time.Time           `json:"createdAt" format:"date-time"`
	UpdatedAt time.Time           `json:"updatedAt" format:"date-time"`
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
	Secondary []TopLevelActionsSecondaryUnion `json:"secondary,required"`
	Primary   TopLevelActionsPrimaryUnion     `json:"primary"`
	Settings  IFrameActionBody                `json:"settings"`
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

type ExtensionCardNewParams struct {
	// State of card definition to be created
	CardCreateRequest CardCreateRequestParam
	paramObj
}

func (r ExtensionCardNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.CardCreateRequest)
}
func (r *ExtensionCardNewParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.CardCreateRequest)
}

type ExtensionCardUpdateParams struct {
	AppID int64 `path:"appId,required" json:"-"`
	// Body for a patch with optional fields
	CardPatchRequest CardPatchRequestParam
	paramObj
}

func (r ExtensionCardUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.CardPatchRequest)
}
func (r *ExtensionCardUpdateParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.CardPatchRequest)
}

type ExtensionCardDeleteParams struct {
	AppID int64 `path:"appId,required" json:"-"`
	paramObj
}

type ExtensionCardGetParams struct {
	AppID int64 `path:"appId,required" json:"-"`
	paramObj
}
