// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package conversations

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/HubSpot/hubspot-sdk-go/internal/apijson"
	"github.com/HubSpot/hubspot-sdk-go/internal/apiquery"
	shimjson "github.com/HubSpot/hubspot-sdk-go/internal/encoding/json"
	"github.com/HubSpot/hubspot-sdk-go/internal/requestconfig"
	"github.com/HubSpot/hubspot-sdk-go/option"
	"github.com/HubSpot/hubspot-sdk-go/packages/pagination"
	"github.com/HubSpot/hubspot-sdk-go/packages/param"
	"github.com/HubSpot/hubspot-sdk-go/packages/respjson"
	"github.com/HubSpot/hubspot-sdk-go/shared"
)

// CustomChannelService contains methods and other services that help with
// interacting with the hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCustomChannelService] method instead.
type CustomChannelService struct {
	options         []option.RequestOption
	ChannelAccounts CustomChannelChannelAccountService
	Messages        CustomChannelMessageService
}

// NewCustomChannelService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCustomChannelService(opts ...option.RequestOption) (r CustomChannelService) {
	r = CustomChannelService{}
	r.options = opts
	r.ChannelAccounts = NewCustomChannelChannelAccountService(opts...)
	r.Messages = NewCustomChannelMessageService(opts...)
	return
}

func (r *CustomChannelService) New(ctx context.Context, body CustomChannelNewParams, opts ...option.RequestOption) (res *PublicChannelIntegrationChannel, err error) {
	opts = slices.Concat(r.options, opts)
	path := "conversations/custom-channels/2026-03"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Update the capabilities for an existing. You can also use it to update the
// channel's webhookUri and its channelAccountConnectionRedirectUrl.
func (r *CustomChannelService) Update(ctx context.Context, channelID int64, body CustomChannelUpdateParams, opts ...option.RequestOption) (res *PublicChannelIntegrationChannel, err error) {
	opts = slices.Concat(r.options, opts)
	path := fmt.Sprintf("conversations/custom-channels/2026-03/%v", channelID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

func (r *CustomChannelService) List(ctx context.Context, query CustomChannelListParams, opts ...option.RequestOption) (res *pagination.Page[PublicChannelIntegrationChannel], err error) {
	var raw *http.Response
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "conversations/custom-channels/2026-03"
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

func (r *CustomChannelService) ListAutoPaging(ctx context.Context, query CustomChannelListParams, opts ...option.RequestOption) *pagination.PageAutoPager[PublicChannelIntegrationChannel] {
	return pagination.NewPageAutoPager(r.List(ctx, query, opts...))
}

// Archive an existing registered custom channel
func (r *CustomChannelService) Delete(ctx context.Context, channelID int64, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := fmt.Sprintf("conversations/custom-channels/2026-03/%v", channelID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Retrieve the details for a specific channel account. This contains all the
// metadata about your channel account, including its channel, associated inbox id,
// and delivery identifier information.
func (r *CustomChannelService) Get(ctx context.Context, channelAccountID int64, params CustomChannelGetParams, opts ...option.RequestOption) (res *PublicChannelAccount, err error) {
	opts = slices.Concat(r.options, opts)
	path := fmt.Sprintf("conversations/custom-channels/2026-03/%v/channel-accounts/%v", params.ChannelID, channelAccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// The properties Attachments, ChannelAccountID, MessageDirection, Recipients,
// Senders, Text, Timestamp are required.
type ChannelIntegrationMessageEggParam struct {
	Attachments      []ChannelIntegrationMessageEggAttachmentUnionParam `json:"attachments,omitzero" api:"required"`
	ChannelAccountID string                                             `json:"channelAccountId" api:"required"`
	// Any of "INCOMING", "OUTGOING".
	MessageDirection         ChannelIntegrationMessageEggMessageDirection `json:"messageDirection,omitzero" api:"required"`
	Recipients               []ChannelIntegrationParticipantParam         `json:"recipients,omitzero" api:"required"`
	Senders                  []ChannelIntegrationParticipantParam         `json:"senders,omitzero" api:"required"`
	Text                     string                                       `json:"text" api:"required"`
	Timestamp                time.Time                                    `json:"timestamp" api:"required" format:"date-time"`
	AssociateWithContactID   param.Opt[int64]                             `json:"associateWithContactId,omitzero"`
	InReplyToID              param.Opt[string]                            `json:"inReplyToId,omitzero"`
	IntegrationIdempotencyID param.Opt[string]                            `json:"integrationIdempotencyId,omitzero"`
	IntegrationThreadID      param.Opt[string]                            `json:"integrationThreadId,omitzero"`
	RichText                 param.Opt[string]                            `json:"richText,omitzero"`
	paramObj
}

func (r ChannelIntegrationMessageEggParam) MarshalJSON() (data []byte, err error) {
	type shadow ChannelIntegrationMessageEggParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ChannelIntegrationMessageEggParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ChannelIntegrationMessageEggAttachmentUnionParam struct {
	OfFile                *FileAttachmentParam                      `json:",omitzero,inline"`
	OfLocation            *LocationAttachmentParam                  `json:",omitzero,inline"`
	OfContact             *ContactAttachmentParam                   `json:",omitzero,inline"`
	OfUnsupportedContent  *UnsupportedContentAttachmentParam        `json:",omitzero,inline"`
	OfMessageHeader       *MessageHeaderAttachmentParam             `json:",omitzero,inline"`
	OfQuickReplies        *QuickRepliesAttachmentParam              `json:",omitzero,inline"`
	OfSocialMediaMetadata *SocialMetadataIntegrationAttachmentParam `json:",omitzero,inline"`
	paramUnion
}

func (u ChannelIntegrationMessageEggAttachmentUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfFile,
		u.OfLocation,
		u.OfContact,
		u.OfUnsupportedContent,
		u.OfMessageHeader,
		u.OfQuickReplies,
		u.OfSocialMediaMetadata)
}
func (u *ChannelIntegrationMessageEggAttachmentUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

type ChannelIntegrationMessageEggMessageDirection string

const (
	ChannelIntegrationMessageEggMessageDirectionIncoming ChannelIntegrationMessageEggMessageDirection = "INCOMING"
	ChannelIntegrationMessageEggMessageDirectionOutgoing ChannelIntegrationMessageEggMessageDirection = "OUTGOING"
)

// The property DeliveryIdentifier is required.
type ChannelIntegrationParticipantParam struct {
	DeliveryIdentifier PublicDeliveryIdentifierParam `json:"deliveryIdentifier,omitzero" api:"required"`
	Name               param.Opt[string]             `json:"name,omitzero"`
	SenderActorID      param.Opt[string]             `json:"senderActorId,omitzero"`
	paramObj
}

func (r ChannelIntegrationParticipantParam) MarshalJSON() (data []byte, err error) {
	type shadow ChannelIntegrationParticipantParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ChannelIntegrationParticipantParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CollectionResponseWithTotalPublicChannelAccount struct {
	Results []PublicChannelAccount `json:"results" api:"required"`
	Total   int64                  `json:"total" api:"required"`
	Paging  shared.Paging          `json:"paging"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Results     respjson.Field
		Total       respjson.Field
		Paging      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CollectionResponseWithTotalPublicChannelAccount) RawJSON() string { return r.JSON.raw }
func (r *CollectionResponseWithTotalPublicChannelAccount) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CollectionResponseWithTotalPublicChannelIntegrationChannel struct {
	Results []PublicChannelIntegrationChannel `json:"results" api:"required"`
	Total   int64                             `json:"total" api:"required"`
	Paging  shared.Paging                     `json:"paging"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Results     respjson.Field
		Total       respjson.Field
		Paging      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CollectionResponseWithTotalPublicChannelIntegrationChannel) RawJSON() string {
	return r.JSON.raw
}
func (r *CollectionResponseWithTotalPublicChannelIntegrationChannel) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ContactAddress struct {
	City        string `json:"city"`
	Country     string `json:"country"`
	CountryCode string `json:"countryCode"`
	State       string `json:"state"`
	Street      string `json:"street"`
	// Any of "HOME", "WORK".
	Type ContactAddressType `json:"type"`
	Zip  string             `json:"zip"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		City        respjson.Field
		Country     respjson.Field
		CountryCode respjson.Field
		State       respjson.Field
		Street      respjson.Field
		Type        respjson.Field
		Zip         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ContactAddress) RawJSON() string { return r.JSON.raw }
func (r *ContactAddress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this ContactAddress to a ContactAddressParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// ContactAddressParam.Overrides()
func (r ContactAddress) ToParam() ContactAddressParam {
	return param.Override[ContactAddressParam](json.RawMessage(r.RawJSON()))
}

type ContactAddressType string

const (
	ContactAddressTypeHome ContactAddressType = "HOME"
	ContactAddressTypeWork ContactAddressType = "WORK"
)

type ContactAddressParam struct {
	City        param.Opt[string] `json:"city,omitzero"`
	Country     param.Opt[string] `json:"country,omitzero"`
	CountryCode param.Opt[string] `json:"countryCode,omitzero"`
	State       param.Opt[string] `json:"state,omitzero"`
	Street      param.Opt[string] `json:"street,omitzero"`
	Zip         param.Opt[string] `json:"zip,omitzero"`
	// Any of "HOME", "WORK".
	Type ContactAddressType `json:"type,omitzero"`
	paramObj
}

func (r ContactAddressParam) MarshalJSON() (data []byte, err error) {
	type shadow ContactAddressParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ContactAddressParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties ContactProfile, Type are required.
type ContactAttachmentParam struct {
	ContactProfile ContactProfileParam `json:"contactProfile,omitzero" api:"required"`
	// Any of "CONTACT".
	Type ContactAttachmentType `json:"type,omitzero" api:"required"`
	paramObj
}

func (r ContactAttachmentParam) MarshalJSON() (data []byte, err error) {
	type shadow ContactAttachmentParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ContactAttachmentParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ContactAttachmentType string

const (
	ContactAttachmentTypeContact ContactAttachmentType = "CONTACT"
)

type ContactEmail struct {
	Email string `json:"email" api:"required"`
	// Any of "HOME", "WORK".
	Type ContactEmailType `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Email       respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ContactEmail) RawJSON() string { return r.JSON.raw }
func (r *ContactEmail) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this ContactEmail to a ContactEmailParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// ContactEmailParam.Overrides()
func (r ContactEmail) ToParam() ContactEmailParam {
	return param.Override[ContactEmailParam](json.RawMessage(r.RawJSON()))
}

type ContactEmailType string

const (
	ContactEmailTypeHome ContactEmailType = "HOME"
	ContactEmailTypeWork ContactEmailType = "WORK"
)

// The property Email is required.
type ContactEmailParam struct {
	Email string `json:"email" api:"required"`
	// Any of "HOME", "WORK".
	Type ContactEmailType `json:"type,omitzero"`
	paramObj
}

func (r ContactEmailParam) MarshalJSON() (data []byte, err error) {
	type shadow ContactEmailParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ContactEmailParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ContactName struct {
	FirstName  string `json:"firstName"`
	LastName   string `json:"lastName"`
	MiddleName string `json:"middleName"`
	Prefix     string `json:"prefix"`
	Suffix     string `json:"suffix"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FirstName   respjson.Field
		LastName    respjson.Field
		MiddleName  respjson.Field
		Prefix      respjson.Field
		Suffix      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ContactName) RawJSON() string { return r.JSON.raw }
func (r *ContactName) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this ContactName to a ContactNameParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// ContactNameParam.Overrides()
func (r ContactName) ToParam() ContactNameParam {
	return param.Override[ContactNameParam](json.RawMessage(r.RawJSON()))
}

type ContactNameParam struct {
	FirstName  param.Opt[string] `json:"firstName,omitzero"`
	LastName   param.Opt[string] `json:"lastName,omitzero"`
	MiddleName param.Opt[string] `json:"middleName,omitzero"`
	Prefix     param.Opt[string] `json:"prefix,omitzero"`
	Suffix     param.Opt[string] `json:"suffix,omitzero"`
	paramObj
}

func (r ContactNameParam) MarshalJSON() (data []byte, err error) {
	type shadow ContactNameParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ContactNameParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ContactOrg struct {
	Company    string `json:"company"`
	Department string `json:"department"`
	Title      string `json:"title"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Company     respjson.Field
		Department  respjson.Field
		Title       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ContactOrg) RawJSON() string { return r.JSON.raw }
func (r *ContactOrg) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this ContactOrg to a ContactOrgParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// ContactOrgParam.Overrides()
func (r ContactOrg) ToParam() ContactOrgParam {
	return param.Override[ContactOrgParam](json.RawMessage(r.RawJSON()))
}

type ContactOrgParam struct {
	Company    param.Opt[string] `json:"company,omitzero"`
	Department param.Opt[string] `json:"department,omitzero"`
	Title      param.Opt[string] `json:"title,omitzero"`
	paramObj
}

func (r ContactOrgParam) MarshalJSON() (data []byte, err error) {
	type shadow ContactOrgParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ContactOrgParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ContactPhone struct {
	Phone string `json:"phone" api:"required"`
	// Any of "CELL", "HOME", "MAIN", "WORK".
	Type ContactPhoneType `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Phone       respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ContactPhone) RawJSON() string { return r.JSON.raw }
func (r *ContactPhone) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this ContactPhone to a ContactPhoneParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// ContactPhoneParam.Overrides()
func (r ContactPhone) ToParam() ContactPhoneParam {
	return param.Override[ContactPhoneParam](json.RawMessage(r.RawJSON()))
}

type ContactPhoneType string

const (
	ContactPhoneTypeCell ContactPhoneType = "CELL"
	ContactPhoneTypeHome ContactPhoneType = "HOME"
	ContactPhoneTypeMain ContactPhoneType = "MAIN"
	ContactPhoneTypeWork ContactPhoneType = "WORK"
)

// The property Phone is required.
type ContactPhoneParam struct {
	Phone string `json:"phone" api:"required"`
	// Any of "CELL", "HOME", "MAIN", "WORK".
	Type ContactPhoneType `json:"type,omitzero"`
	paramObj
}

func (r ContactPhoneParam) MarshalJSON() (data []byte, err error) {
	type shadow ContactPhoneParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ContactPhoneParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ContactProfile struct {
	Addresses []ContactAddress `json:"addresses" api:"required"`
	Emails    []ContactEmail   `json:"emails" api:"required"`
	Phones    []ContactPhone   `json:"phones" api:"required"`
	URLs      []ContactURL     `json:"urls" api:"required"`
	Name      ContactName      `json:"name"`
	Org       ContactOrg       `json:"org"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Addresses   respjson.Field
		Emails      respjson.Field
		Phones      respjson.Field
		URLs        respjson.Field
		Name        respjson.Field
		Org         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ContactProfile) RawJSON() string { return r.JSON.raw }
func (r *ContactProfile) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this ContactProfile to a ContactProfileParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// ContactProfileParam.Overrides()
func (r ContactProfile) ToParam() ContactProfileParam {
	return param.Override[ContactProfileParam](json.RawMessage(r.RawJSON()))
}

// The properties Addresses, Emails, Phones, URLs are required.
type ContactProfileParam struct {
	Addresses []ContactAddressParam `json:"addresses,omitzero" api:"required"`
	Emails    []ContactEmailParam   `json:"emails,omitzero" api:"required"`
	Phones    []ContactPhoneParam   `json:"phones,omitzero" api:"required"`
	URLs      []ContactURLParam     `json:"urls,omitzero" api:"required"`
	Name      ContactNameParam      `json:"name,omitzero"`
	Org       ContactOrgParam       `json:"org,omitzero"`
	paramObj
}

func (r ContactProfileParam) MarshalJSON() (data []byte, err error) {
	type shadow ContactProfileParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ContactProfileParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ContactURL struct {
	URL string `json:"url" api:"required"`
	// Any of "HOME", "WORK".
	Type ContactURLType `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		URL         respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ContactURL) RawJSON() string { return r.JSON.raw }
func (r *ContactURL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this ContactURL to a ContactURLParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// ContactURLParam.Overrides()
func (r ContactURL) ToParam() ContactURLParam {
	return param.Override[ContactURLParam](json.RawMessage(r.RawJSON()))
}

type ContactURLType string

const (
	ContactURLTypeHome ContactURLType = "HOME"
	ContactURLTypeWork ContactURLType = "WORK"
)

// The property URL is required.
type ContactURLParam struct {
	URL string `json:"url" api:"required"`
	// Any of "HOME", "WORK".
	Type ContactURLType `json:"type,omitzero"`
	paramObj
}

func (r ContactURLParam) MarshalJSON() (data []byte, err error) {
	type shadow ContactURLParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ContactURLParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties FileID, Type are required.
type FileAttachmentParam struct {
	FileID string `json:"fileId" api:"required"`
	// Any of "FILE".
	Type FileAttachmentType `json:"type,omitzero" api:"required"`
	// Any of "AUDIO", "IMAGE", "OTHER", "STICKER", "VOICE_RECORDING".
	FileUsageType FileAttachmentFileUsageType `json:"fileUsageType,omitzero"`
	paramObj
}

func (r FileAttachmentParam) MarshalJSON() (data []byte, err error) {
	type shadow FileAttachmentParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FileAttachmentParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FileAttachmentType string

const (
	FileAttachmentTypeFile FileAttachmentType = "FILE"
)

type FileAttachmentFileUsageType string

const (
	FileAttachmentFileUsageTypeAudio          FileAttachmentFileUsageType = "AUDIO"
	FileAttachmentFileUsageTypeImage          FileAttachmentFileUsageType = "IMAGE"
	FileAttachmentFileUsageTypeOther          FileAttachmentFileUsageType = "OTHER"
	FileAttachmentFileUsageTypeSticker        FileAttachmentFileUsageType = "STICKER"
	FileAttachmentFileUsageTypeVoiceRecording FileAttachmentFileUsageType = "VOICE_RECORDING"
)

// The properties Latitude, Longitude, Type are required.
type LocationAttachmentParam struct {
	Latitude  float64 `json:"latitude" api:"required"`
	Longitude float64 `json:"longitude" api:"required"`
	// Any of "LOCATION".
	Type    LocationAttachmentType `json:"type,omitzero" api:"required"`
	Address param.Opt[string]      `json:"address,omitzero"`
	Name    param.Opt[string]      `json:"name,omitzero"`
	URL     param.Opt[string]      `json:"url,omitzero"`
	paramObj
}

func (r LocationAttachmentParam) MarshalJSON() (data []byte, err error) {
	type shadow LocationAttachmentParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LocationAttachmentParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LocationAttachmentType string

const (
	LocationAttachmentTypeLocation LocationAttachmentType = "LOCATION"
)

// The property Type is required.
type MessageHeaderAttachmentParam struct {
	// Any of "MESSAGE_HEADER".
	Type   MessageHeaderAttachmentType `json:"type,omitzero" api:"required"`
	FileID param.Opt[int64]            `json:"fileId,omitzero"`
	Text   param.Opt[string]           `json:"text,omitzero"`
	paramObj
}

func (r MessageHeaderAttachmentParam) MarshalJSON() (data []byte, err error) {
	type shadow MessageHeaderAttachmentParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MessageHeaderAttachmentParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MessageHeaderAttachmentType string

const (
	MessageHeaderAttachmentTypeMessageHeader MessageHeaderAttachmentType = "MESSAGE_HEADER"
)

type PublicChannelAccount struct {
	ID                 string                   `json:"id" api:"required"`
	Active             bool                     `json:"active" api:"required"`
	Archived           bool                     `json:"archived" api:"required"`
	Authorized         bool                     `json:"authorized" api:"required"`
	ChannelID          string                   `json:"channelId" api:"required"`
	CreatedAt          time.Time                `json:"createdAt" api:"required" format:"date-time"`
	InboxID            string                   `json:"inboxId" api:"required"`
	Name               string                   `json:"name" api:"required"`
	ArchivedAt         time.Time                `json:"archivedAt" format:"date-time"`
	DeliveryIdentifier PublicDeliveryIdentifier `json:"deliveryIdentifier"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                 respjson.Field
		Active             respjson.Field
		Archived           respjson.Field
		Authorized         respjson.Field
		ChannelID          respjson.Field
		CreatedAt          respjson.Field
		InboxID            respjson.Field
		Name               respjson.Field
		ArchivedAt         respjson.Field
		DeliveryIdentifier respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PublicChannelAccount) RawJSON() string { return r.JSON.raw }
func (r *PublicChannelAccount) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Authorized, InboxID, Name are required.
type PublicChannelAccountEggParam struct {
	Authorized         bool                          `json:"authorized" api:"required"`
	InboxID            string                        `json:"inboxId" api:"required"`
	Name               string                        `json:"name" api:"required"`
	DeliveryIdentifier PublicDeliveryIdentifierParam `json:"deliveryIdentifier,omitzero"`
	paramObj
}

func (r PublicChannelAccountEggParam) MarshalJSON() (data []byte, err error) {
	type shadow PublicChannelAccountEggParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PublicChannelAccountEggParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PublicChannelAccountStagingToken struct {
	AccountToken       string                   `json:"accountToken" api:"required"`
	CreatedAt          time.Time                `json:"createdAt" api:"required" format:"date-time"`
	GenericChannelID   int64                    `json:"genericChannelId" api:"required"`
	InboxID            int64                    `json:"inboxId" api:"required"`
	UserID             int64                    `json:"userId" api:"required"`
	AccountName        string                   `json:"accountName"`
	DeliveryIdentifier PublicDeliveryIdentifier `json:"deliveryIdentifier"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AccountToken       respjson.Field
		CreatedAt          respjson.Field
		GenericChannelID   respjson.Field
		InboxID            respjson.Field
		UserID             respjson.Field
		AccountName        respjson.Field
		DeliveryIdentifier respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PublicChannelAccountStagingToken) RawJSON() string { return r.JSON.raw }
func (r *PublicChannelAccountStagingToken) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PublicChannelAccountStagingTokenUpdateRequestParam struct {
	AccountName        param.Opt[string]             `json:"accountName,omitzero"`
	DeliveryIdentifier PublicDeliveryIdentifierParam `json:"deliveryIdentifier,omitzero"`
	paramObj
}

func (r PublicChannelAccountStagingTokenUpdateRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow PublicChannelAccountStagingTokenUpdateRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PublicChannelAccountStagingTokenUpdateRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PublicChannelAccountUpdateRequestParam struct {
	Authorized param.Opt[bool]   `json:"authorized,omitzero"`
	Name       param.Opt[string] `json:"name,omitzero"`
	paramObj
}

func (r PublicChannelAccountUpdateRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow PublicChannelAccountUpdateRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PublicChannelAccountUpdateRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PublicChannelIntegrationChannel struct {
	ID                                  string         `json:"id" api:"required"`
	Capabilities                        map[string]any `json:"capabilities" api:"required"`
	CreatedAt                           time.Time      `json:"createdAt" api:"required" format:"date-time"`
	Name                                string         `json:"name" api:"required"`
	ChannelAccountConnectionRedirectURL string         `json:"channelAccountConnectionRedirectUrl"`
	ChannelDescription                  string         `json:"channelDescription"`
	ChannelLogoURL                      string         `json:"channelLogoUrl"`
	WebhookURL                          string         `json:"webhookUrl"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                                  respjson.Field
		Capabilities                        respjson.Field
		CreatedAt                           respjson.Field
		Name                                respjson.Field
		ChannelAccountConnectionRedirectURL respjson.Field
		ChannelDescription                  respjson.Field
		ChannelLogoURL                      respjson.Field
		WebhookURL                          respjson.Field
		ExtraFields                         map[string]respjson.Field
		raw                                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PublicChannelIntegrationChannel) RawJSON() string { return r.JSON.raw }
func (r *PublicChannelIntegrationChannel) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Capabilities, Name are required.
type PublicChannelIntegrationChannelCreateParam struct {
	Capabilities                        map[string]any    `json:"capabilities,omitzero" api:"required"`
	Name                                string            `json:"name" api:"required"`
	ChannelAccountConnectionRedirectURL param.Opt[string] `json:"channelAccountConnectionRedirectUrl,omitzero"`
	ChannelDescription                  param.Opt[string] `json:"channelDescription,omitzero"`
	ChannelLogoURL                      param.Opt[string] `json:"channelLogoUrl,omitzero"`
	WebhookURL                          param.Opt[string] `json:"webhookUrl,omitzero"`
	paramObj
}

func (r PublicChannelIntegrationChannelCreateParam) MarshalJSON() (data []byte, err error) {
	type shadow PublicChannelIntegrationChannelCreateParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PublicChannelIntegrationChannelCreateParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Capabilities, ChannelAccountConnectionRedirectURL,
// ChannelDescription, ChannelLogoURL, Name, WebhookURL are required.
type PublicChannelIntegrationChannelPatchParam struct {
	Capabilities                        map[string]any `json:"capabilities,omitzero" api:"required"`
	ChannelAccountConnectionRedirectURL any            `json:"channelAccountConnectionRedirectUrl,omitzero" api:"required"`
	ChannelDescription                  any            `json:"channelDescription,omitzero" api:"required"`
	ChannelLogoURL                      any            `json:"channelLogoUrl,omitzero" api:"required"`
	Name                                any            `json:"name,omitzero" api:"required"`
	WebhookURL                          any            `json:"webhookUrl,omitzero" api:"required"`
	paramObj
}

func (r PublicChannelIntegrationChannelPatchParam) MarshalJSON() (data []byte, err error) {
	type shadow PublicChannelIntegrationChannelPatchParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PublicChannelIntegrationChannelPatchParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property StatusType is required.
type PublicChannelIntegrationMessageUpdateRequestParam struct {
	// Valid status are SENT, FAILED, and READ
	//
	// Any of "FAILED", "READ", "SENT".
	StatusType   PublicChannelIntegrationMessageUpdateRequestStatusType `json:"statusType,omitzero" api:"required"`
	ErrorMessage param.Opt[string]                                      `json:"errorMessage,omitzero"`
	paramObj
}

func (r PublicChannelIntegrationMessageUpdateRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow PublicChannelIntegrationMessageUpdateRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PublicChannelIntegrationMessageUpdateRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Valid status are SENT, FAILED, and READ
type PublicChannelIntegrationMessageUpdateRequestStatusType string

const (
	PublicChannelIntegrationMessageUpdateRequestStatusTypeFailed PublicChannelIntegrationMessageUpdateRequestStatusType = "FAILED"
	PublicChannelIntegrationMessageUpdateRequestStatusTypeRead   PublicChannelIntegrationMessageUpdateRequestStatusType = "READ"
	PublicChannelIntegrationMessageUpdateRequestStatusTypeSent   PublicChannelIntegrationMessageUpdateRequestStatusType = "SENT"
)

type PublicClient struct {
	// Any of "HUBSPOT", "INTEGRATION", "SYSTEM", "UNKNOWN".
	ClientType       PublicClientClientType `json:"clientType" api:"required"`
	IntegrationAppID int64                  `json:"integrationAppId"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClientType       respjson.Field
		IntegrationAppID respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PublicClient) RawJSON() string { return r.JSON.raw }
func (r *PublicClient) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PublicClientClientType string

const (
	PublicClientClientTypeHubSpot     PublicClientClientType = "HUBSPOT"
	PublicClientClientTypeIntegration PublicClientClientType = "INTEGRATION"
	PublicClientClientTypeSystem      PublicClientClientType = "SYSTEM"
	PublicClientClientTypeUnknown     PublicClientClientType = "UNKNOWN"
)

type PublicContact struct {
	ContactProfile ContactProfile `json:"contactProfile" api:"required"`
	// Any of "CONTACT".
	Type PublicContactType `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ContactProfile respjson.Field
		Type           respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PublicContact) RawJSON() string { return r.JSON.raw }
func (r *PublicContact) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PublicContactType string

const (
	PublicContactTypeContact PublicContactType = "CONTACT"
)

type PublicConversationsMessage struct {
	ID                    string                                      `json:"id" api:"required"`
	Archived              bool                                        `json:"archived" api:"required"`
	Attachments           []PublicConversationsMessageAttachmentUnion `json:"attachments" api:"required"`
	ChannelAccountID      string                                      `json:"channelAccountId" api:"required"`
	ChannelID             string                                      `json:"channelId" api:"required"`
	Client                PublicClient                                `json:"client" api:"required"`
	ConversationsThreadID string                                      `json:"conversationsThreadId" api:"required"`
	CreatedAt             time.Time                                   `json:"createdAt" api:"required" format:"date-time"`
	CreatedBy             string                                      `json:"createdBy" api:"required"`
	// Any of "INCOMING", "OUTGOING".
	Direction  PublicConversationsMessageDirection `json:"direction" api:"required"`
	Recipients []PublicRecipient                   `json:"recipients" api:"required"`
	Senders    []PublicSender                      `json:"senders" api:"required"`
	Text       string                              `json:"text" api:"required"`
	// Any of "NOT_TRUNCATED", "TRUNCATED", "TRUNCATED_TO_MOST_RECENT_REPLY".
	TruncationStatus PublicConversationsMessageTruncationStatus `json:"truncationStatus" api:"required"`
	// Any of "MESSAGE".
	Type        PublicConversationsMessageType `json:"type" api:"required"`
	InReplyToID string                         `json:"inReplyToId"`
	RichText    string                         `json:"richText"`
	Status      PublicMessageStatus            `json:"status"`
	Subject     string                         `json:"subject"`
	UpdatedAt   time.Time                      `json:"updatedAt" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                    respjson.Field
		Archived              respjson.Field
		Attachments           respjson.Field
		ChannelAccountID      respjson.Field
		ChannelID             respjson.Field
		Client                respjson.Field
		ConversationsThreadID respjson.Field
		CreatedAt             respjson.Field
		CreatedBy             respjson.Field
		Direction             respjson.Field
		Recipients            respjson.Field
		Senders               respjson.Field
		Text                  respjson.Field
		TruncationStatus      respjson.Field
		Type                  respjson.Field
		InReplyToID           respjson.Field
		RichText              respjson.Field
		Status                respjson.Field
		Subject               respjson.Field
		UpdatedAt             respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PublicConversationsMessage) RawJSON() string { return r.JSON.raw }
func (r *PublicConversationsMessage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PublicConversationsMessageAttachmentUnion contains all possible properties and
// values from [PublicFile], [PublicLocation], [PublicContact],
// [PublicUnsupportedContent], [PublicMessageHeader], [PublicQuickReplies],
// [PublicWhatsAppTemplateMetadata], [PublicSocialMetadataAttachment].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type PublicConversationsMessageAttachmentUnion struct {
	// This field is a union of [string], [int64]
	FileID PublicConversationsMessageAttachmentUnionFileID `json:"fileId"`
	// This field is from variant [PublicFile].
	FileUsageType PublicFileFileUsageType `json:"fileUsageType"`
	Type          string                  `json:"type"`
	Name          string                  `json:"name"`
	URL           string                  `json:"url"`
	// This field is from variant [PublicLocation].
	Latitude float64 `json:"latitude"`
	// This field is from variant [PublicLocation].
	Longitude float64 `json:"longitude"`
	// This field is from variant [PublicLocation].
	Address string `json:"address"`
	// This field is from variant [PublicContact].
	ContactProfile ContactProfile `json:"contactProfile"`
	// This field is from variant [PublicMessageHeader].
	Text string `json:"text"`
	// This field is from variant [PublicQuickReplies].
	AllowMultiSelect bool `json:"allowMultiSelect"`
	// This field is from variant [PublicQuickReplies].
	AllowUserInput bool `json:"allowUserInput"`
	// This field is from variant [PublicQuickReplies].
	QuickReplies []QuickReply `json:"quickReplies"`
	// This field is from variant [PublicWhatsAppTemplateMetadata].
	CrmObjectIDs map[string]int64 `json:"crmObjectIds"`
	// This field is from variant [PublicWhatsAppTemplateMetadata].
	Parameters map[string]string `json:"parameters"`
	// This field is from variant [PublicWhatsAppTemplateMetadata].
	ContentID int64 `json:"contentId"`
	// This field is from variant [PublicWhatsAppTemplateMetadata].
	MappedTemplateID int64 `json:"mappedTemplateId"`
	// This field is from variant [PublicWhatsAppTemplateMetadata].
	RootMicID int64 `json:"rootMicId"`
	// This field is from variant [PublicSocialMetadataAttachment].
	SocialMetadata SocialMetadata `json:"socialMetadata"`
	JSON           struct {
		FileID           respjson.Field
		FileUsageType    respjson.Field
		Type             respjson.Field
		Name             respjson.Field
		URL              respjson.Field
		Latitude         respjson.Field
		Longitude        respjson.Field
		Address          respjson.Field
		ContactProfile   respjson.Field
		Text             respjson.Field
		AllowMultiSelect respjson.Field
		AllowUserInput   respjson.Field
		QuickReplies     respjson.Field
		CrmObjectIDs     respjson.Field
		Parameters       respjson.Field
		ContentID        respjson.Field
		MappedTemplateID respjson.Field
		RootMicID        respjson.Field
		SocialMetadata   respjson.Field
		raw              string
	} `json:"-"`
}

func (u PublicConversationsMessageAttachmentUnion) AsFile() (v PublicFile) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicConversationsMessageAttachmentUnion) AsLocation() (v PublicLocation) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicConversationsMessageAttachmentUnion) AsContact() (v PublicContact) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicConversationsMessageAttachmentUnion) AsUnsupportedContent() (v PublicUnsupportedContent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicConversationsMessageAttachmentUnion) AsMessageHeader() (v PublicMessageHeader) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicConversationsMessageAttachmentUnion) AsQuickReplies() (v PublicQuickReplies) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicConversationsMessageAttachmentUnion) AsWhatsappTemplateMetadata() (v PublicWhatsAppTemplateMetadata) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicConversationsMessageAttachmentUnion) AsSocialMediaMetadata() (v PublicSocialMetadataAttachment) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u PublicConversationsMessageAttachmentUnion) RawJSON() string { return u.JSON.raw }

func (r *PublicConversationsMessageAttachmentUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PublicConversationsMessageAttachmentUnionFileID is an implicit subunion of
// [PublicConversationsMessageAttachmentUnion].
// PublicConversationsMessageAttachmentUnionFileID provides convenient access to
// the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [PublicConversationsMessageAttachmentUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfInt]
type PublicConversationsMessageAttachmentUnionFileID struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [int64] instead of an object.
	OfInt int64 `json:",inline"`
	JSON  struct {
		OfString respjson.Field
		OfInt    respjson.Field
		raw      string
	} `json:"-"`
}

func (r *PublicConversationsMessageAttachmentUnionFileID) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PublicConversationsMessageDirection string

const (
	PublicConversationsMessageDirectionIncoming PublicConversationsMessageDirection = "INCOMING"
	PublicConversationsMessageDirectionOutgoing PublicConversationsMessageDirection = "OUTGOING"
)

type PublicConversationsMessageTruncationStatus string

const (
	PublicConversationsMessageTruncationStatusNotTruncated               PublicConversationsMessageTruncationStatus = "NOT_TRUNCATED"
	PublicConversationsMessageTruncationStatusTruncated                  PublicConversationsMessageTruncationStatus = "TRUNCATED"
	PublicConversationsMessageTruncationStatusTruncatedToMostRecentReply PublicConversationsMessageTruncationStatus = "TRUNCATED_TO_MOST_RECENT_REPLY"
)

type PublicConversationsMessageType string

const (
	PublicConversationsMessageTypeMessage PublicConversationsMessageType = "MESSAGE"
)

type PublicDeliveryIdentifier struct {
	// Any of "CHANNEL_SPECIFIC_OPAQUE_ID", "HS_EMAIL_ADDRESS", "HS_PHONE_NUMBER",
	// "HS_SHORT_CODE".
	Type  PublicDeliveryIdentifierType `json:"type" api:"required"`
	Value string                       `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PublicDeliveryIdentifier) RawJSON() string { return r.JSON.raw }
func (r *PublicDeliveryIdentifier) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this PublicDeliveryIdentifier to a
// PublicDeliveryIdentifierParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// PublicDeliveryIdentifierParam.Overrides()
func (r PublicDeliveryIdentifier) ToParam() PublicDeliveryIdentifierParam {
	return param.Override[PublicDeliveryIdentifierParam](json.RawMessage(r.RawJSON()))
}

type PublicDeliveryIdentifierType string

const (
	PublicDeliveryIdentifierTypeChannelSpecificOpaqueID PublicDeliveryIdentifierType = "CHANNEL_SPECIFIC_OPAQUE_ID"
	PublicDeliveryIdentifierTypeHsEmailAddress          PublicDeliveryIdentifierType = "HS_EMAIL_ADDRESS"
	PublicDeliveryIdentifierTypeHsPhoneNumber           PublicDeliveryIdentifierType = "HS_PHONE_NUMBER"
	PublicDeliveryIdentifierTypeHsShortCode             PublicDeliveryIdentifierType = "HS_SHORT_CODE"
)

// The properties Type, Value are required.
type PublicDeliveryIdentifierParam struct {
	// Any of "CHANNEL_SPECIFIC_OPAQUE_ID", "HS_EMAIL_ADDRESS", "HS_PHONE_NUMBER",
	// "HS_SHORT_CODE".
	Type  PublicDeliveryIdentifierType `json:"type,omitzero" api:"required"`
	Value string                       `json:"value" api:"required"`
	paramObj
}

func (r PublicDeliveryIdentifierParam) MarshalJSON() (data []byte, err error) {
	type shadow PublicDeliveryIdentifierParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PublicDeliveryIdentifierParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PublicFile struct {
	FileID string `json:"fileId" api:"required"`
	// Any of "AUDIO", "IMAGE", "OTHER", "STICKER", "VOICE_RECORDING".
	FileUsageType PublicFileFileUsageType `json:"fileUsageType" api:"required"`
	// Any of "FILE".
	Type PublicFileType `json:"type" api:"required"`
	Name string         `json:"name"`
	URL  string         `json:"url"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FileID        respjson.Field
		FileUsageType respjson.Field
		Type          respjson.Field
		Name          respjson.Field
		URL           respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PublicFile) RawJSON() string { return r.JSON.raw }
func (r *PublicFile) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PublicFileFileUsageType string

const (
	PublicFileFileUsageTypeAudio          PublicFileFileUsageType = "AUDIO"
	PublicFileFileUsageTypeImage          PublicFileFileUsageType = "IMAGE"
	PublicFileFileUsageTypeOther          PublicFileFileUsageType = "OTHER"
	PublicFileFileUsageTypeSticker        PublicFileFileUsageType = "STICKER"
	PublicFileFileUsageTypeVoiceRecording PublicFileFileUsageType = "VOICE_RECORDING"
)

type PublicFileType string

const (
	PublicFileTypeFile PublicFileType = "FILE"
)

type PublicLocation struct {
	Latitude  float64 `json:"latitude" api:"required"`
	Longitude float64 `json:"longitude" api:"required"`
	// Any of "LOCATION".
	Type    PublicLocationType `json:"type" api:"required"`
	Address string             `json:"address"`
	Name    string             `json:"name"`
	URL     string             `json:"url"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Latitude    respjson.Field
		Longitude   respjson.Field
		Type        respjson.Field
		Address     respjson.Field
		Name        respjson.Field
		URL         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PublicLocation) RawJSON() string { return r.JSON.raw }
func (r *PublicLocation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PublicLocationType string

const (
	PublicLocationTypeLocation PublicLocationType = "LOCATION"
)

type PublicMessageFailureDetails struct {
	ErrorMessageTokens map[string]string `json:"errorMessageTokens" api:"required"`
	ErrorMessage       string            `json:"errorMessage"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ErrorMessageTokens respjson.Field
		ErrorMessage       respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PublicMessageFailureDetails) RawJSON() string { return r.JSON.raw }
func (r *PublicMessageFailureDetails) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PublicMessageHeader struct {
	// Any of "MESSAGE_HEADER".
	Type   PublicMessageHeaderType `json:"type" api:"required"`
	FileID int64                   `json:"fileId"`
	Text   string                  `json:"text"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		FileID      respjson.Field
		Text        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PublicMessageHeader) RawJSON() string { return r.JSON.raw }
func (r *PublicMessageHeader) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PublicMessageHeaderType string

const (
	PublicMessageHeaderTypeMessageHeader PublicMessageHeaderType = "MESSAGE_HEADER"
)

type PublicMessageStatus struct {
	// Any of "FAILED", "READ", "RECEIVED", "SENT".
	StatusType     PublicMessageStatusStatusType `json:"statusType" api:"required"`
	FailureDetails PublicMessageFailureDetails   `json:"failureDetails"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		StatusType     respjson.Field
		FailureDetails respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PublicMessageStatus) RawJSON() string { return r.JSON.raw }
func (r *PublicMessageStatus) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PublicMessageStatusStatusType string

const (
	PublicMessageStatusStatusTypeFailed   PublicMessageStatusStatusType = "FAILED"
	PublicMessageStatusStatusTypeRead     PublicMessageStatusStatusType = "READ"
	PublicMessageStatusStatusTypeReceived PublicMessageStatusStatusType = "RECEIVED"
	PublicMessageStatusStatusTypeSent     PublicMessageStatusStatusType = "SENT"
)

type PublicQuickReplies struct {
	AllowMultiSelect bool         `json:"allowMultiSelect" api:"required"`
	AllowUserInput   bool         `json:"allowUserInput" api:"required"`
	QuickReplies     []QuickReply `json:"quickReplies" api:"required"`
	// Any of "QUICK_REPLIES".
	Type PublicQuickRepliesType `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AllowMultiSelect respjson.Field
		AllowUserInput   respjson.Field
		QuickReplies     respjson.Field
		Type             respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PublicQuickReplies) RawJSON() string { return r.JSON.raw }
func (r *PublicQuickReplies) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PublicQuickRepliesType string

const (
	PublicQuickRepliesTypeQuickReplies PublicQuickRepliesType = "QUICK_REPLIES"
)

type PublicRecipient struct {
	DeliveryIdentifier PublicDeliveryIdentifier `json:"deliveryIdentifier" api:"required"`
	ActorID            string                   `json:"actorId"`
	Name               string                   `json:"name"`
	RecipientField     string                   `json:"recipientField"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DeliveryIdentifier respjson.Field
		ActorID            respjson.Field
		Name               respjson.Field
		RecipientField     respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PublicRecipient) RawJSON() string { return r.JSON.raw }
func (r *PublicRecipient) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PublicSender struct {
	ActorID            string                   `json:"actorId"`
	DeliveryIdentifier PublicDeliveryIdentifier `json:"deliveryIdentifier"`
	Name               string                   `json:"name"`
	SenderField        string                   `json:"senderField"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ActorID            respjson.Field
		DeliveryIdentifier respjson.Field
		Name               respjson.Field
		SenderField        respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PublicSender) RawJSON() string { return r.JSON.raw }
func (r *PublicSender) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PublicSocialMetadataAttachment struct {
	SocialMetadata SocialMetadata `json:"socialMetadata" api:"required"`
	// Any of "SOCIAL_MEDIA_METADATA".
	Type PublicSocialMetadataAttachmentType `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		SocialMetadata respjson.Field
		Type           respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PublicSocialMetadataAttachment) RawJSON() string { return r.JSON.raw }
func (r *PublicSocialMetadataAttachment) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PublicSocialMetadataAttachmentType string

const (
	PublicSocialMetadataAttachmentTypeSocialMediaMetadata PublicSocialMetadataAttachmentType = "SOCIAL_MEDIA_METADATA"
)

type PublicUnsupportedContent struct {
	// Any of "UNSUPPORTED_CONTENT".
	Type PublicUnsupportedContentType `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PublicUnsupportedContent) RawJSON() string { return r.JSON.raw }
func (r *PublicUnsupportedContent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PublicUnsupportedContentType string

const (
	PublicUnsupportedContentTypeUnsupportedContent PublicUnsupportedContentType = "UNSUPPORTED_CONTENT"
)

type PublicWhatsAppTemplateMetadata struct {
	CrmObjectIDs map[string]int64  `json:"crmObjectIds" api:"required"`
	Parameters   map[string]string `json:"parameters" api:"required"`
	// Any of "WHATSAPP_TEMPLATE_METADATA".
	Type             PublicWhatsAppTemplateMetadataType `json:"type" api:"required"`
	ContentID        int64                              `json:"contentId"`
	MappedTemplateID int64                              `json:"mappedTemplateId"`
	RootMicID        int64                              `json:"rootMicId"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CrmObjectIDs     respjson.Field
		Parameters       respjson.Field
		Type             respjson.Field
		ContentID        respjson.Field
		MappedTemplateID respjson.Field
		RootMicID        respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PublicWhatsAppTemplateMetadata) RawJSON() string { return r.JSON.raw }
func (r *PublicWhatsAppTemplateMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PublicWhatsAppTemplateMetadataType string

const (
	PublicWhatsAppTemplateMetadataTypeWhatsappTemplateMetadata PublicWhatsAppTemplateMetadataType = "WHATSAPP_TEMPLATE_METADATA"
)

// The properties QuickReplies, Type are required.
type QuickRepliesAttachmentParam struct {
	QuickReplies []QuickReplyParam `json:"quickReplies,omitzero" api:"required"`
	// Any of "QUICK_REPLIES".
	Type QuickRepliesAttachmentType `json:"type,omitzero" api:"required"`
	paramObj
}

func (r QuickRepliesAttachmentParam) MarshalJSON() (data []byte, err error) {
	type shadow QuickRepliesAttachmentParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *QuickRepliesAttachmentParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type QuickRepliesAttachmentType string

const (
	QuickRepliesAttachmentTypeQuickReplies QuickRepliesAttachmentType = "QUICK_REPLIES"
)

type QuickReply struct {
	Value string `json:"value" api:"required"`
	// Any of "TEXT", "URL".
	ValueType QuickReplyValueType `json:"valueType" api:"required"`
	Label     string              `json:"label"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Value       respjson.Field
		ValueType   respjson.Field
		Label       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r QuickReply) RawJSON() string { return r.JSON.raw }
func (r *QuickReply) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this QuickReply to a QuickReplyParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// QuickReplyParam.Overrides()
func (r QuickReply) ToParam() QuickReplyParam {
	return param.Override[QuickReplyParam](json.RawMessage(r.RawJSON()))
}

type QuickReplyValueType string

const (
	QuickReplyValueTypeText QuickReplyValueType = "TEXT"
	QuickReplyValueTypeURL  QuickReplyValueType = "URL"
)

// The properties Value, ValueType are required.
type QuickReplyParam struct {
	Value string `json:"value" api:"required"`
	// Any of "TEXT", "URL".
	ValueType QuickReplyValueType `json:"valueType,omitzero" api:"required"`
	Label     param.Opt[string]   `json:"label,omitzero"`
	paramObj
}

func (r QuickReplyParam) MarshalJSON() (data []byte, err error) {
	type shadow QuickReplyParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *QuickReplyParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SocialMetadata struct {
	// Any of "ARTICLE", "AUDIO", "CAROUSEL", "DOCUMENT", "GIF", "LINK", "NONE",
	// "PHOTO", "POLL", "STORY", "VIDEO".
	MediaType      SocialMetadataMediaType `json:"mediaType" api:"required"`
	ID             string                  `json:"id"`
	Description    string                  `json:"description"`
	MediaTitle     string                  `json:"mediaTitle"`
	MediaURL       string                  `json:"mediaUrl"`
	MediaURLString string                  `json:"mediaUrlString"`
	ThumbnailURL   string                  `json:"thumbnailUrl"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		MediaType      respjson.Field
		ID             respjson.Field
		Description    respjson.Field
		MediaTitle     respjson.Field
		MediaURL       respjson.Field
		MediaURLString respjson.Field
		ThumbnailURL   respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SocialMetadata) RawJSON() string { return r.JSON.raw }
func (r *SocialMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this SocialMetadata to a SocialMetadataParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// SocialMetadataParam.Overrides()
func (r SocialMetadata) ToParam() SocialMetadataParam {
	return param.Override[SocialMetadataParam](json.RawMessage(r.RawJSON()))
}

type SocialMetadataMediaType string

const (
	SocialMetadataMediaTypeArticle  SocialMetadataMediaType = "ARTICLE"
	SocialMetadataMediaTypeAudio    SocialMetadataMediaType = "AUDIO"
	SocialMetadataMediaTypeCarousel SocialMetadataMediaType = "CAROUSEL"
	SocialMetadataMediaTypeDocument SocialMetadataMediaType = "DOCUMENT"
	SocialMetadataMediaTypeGif      SocialMetadataMediaType = "GIF"
	SocialMetadataMediaTypeLink     SocialMetadataMediaType = "LINK"
	SocialMetadataMediaTypeNone     SocialMetadataMediaType = "NONE"
	SocialMetadataMediaTypePhoto    SocialMetadataMediaType = "PHOTO"
	SocialMetadataMediaTypePoll     SocialMetadataMediaType = "POLL"
	SocialMetadataMediaTypeStory    SocialMetadataMediaType = "STORY"
	SocialMetadataMediaTypeVideo    SocialMetadataMediaType = "VIDEO"
)

// The property MediaType is required.
type SocialMetadataParam struct {
	// Any of "ARTICLE", "AUDIO", "CAROUSEL", "DOCUMENT", "GIF", "LINK", "NONE",
	// "PHOTO", "POLL", "STORY", "VIDEO".
	MediaType      SocialMetadataMediaType `json:"mediaType,omitzero" api:"required"`
	ID             param.Opt[string]       `json:"id,omitzero"`
	Description    param.Opt[string]       `json:"description,omitzero"`
	MediaTitle     param.Opt[string]       `json:"mediaTitle,omitzero"`
	MediaURL       param.Opt[string]       `json:"mediaUrl,omitzero"`
	MediaURLString param.Opt[string]       `json:"mediaUrlString,omitzero"`
	ThumbnailURL   param.Opt[string]       `json:"thumbnailUrl,omitzero"`
	paramObj
}

func (r SocialMetadataParam) MarshalJSON() (data []byte, err error) {
	type shadow SocialMetadataParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SocialMetadataParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties SocialMetadata, Type are required.
type SocialMetadataIntegrationAttachmentParam struct {
	SocialMetadata SocialMetadataParam `json:"socialMetadata,omitzero" api:"required"`
	// Any of "SOCIAL_MEDIA_METADATA".
	Type SocialMetadataIntegrationAttachmentType `json:"type,omitzero" api:"required"`
	paramObj
}

func (r SocialMetadataIntegrationAttachmentParam) MarshalJSON() (data []byte, err error) {
	type shadow SocialMetadataIntegrationAttachmentParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SocialMetadataIntegrationAttachmentParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SocialMetadataIntegrationAttachmentType string

const (
	SocialMetadataIntegrationAttachmentTypeSocialMediaMetadata SocialMetadataIntegrationAttachmentType = "SOCIAL_MEDIA_METADATA"
)

// The property Type is required.
type UnsupportedContentAttachmentParam struct {
	// Any of "UNSUPPORTED_CONTENT".
	Type UnsupportedContentAttachmentType `json:"type,omitzero" api:"required"`
	paramObj
}

func (r UnsupportedContentAttachmentParam) MarshalJSON() (data []byte, err error) {
	type shadow UnsupportedContentAttachmentParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UnsupportedContentAttachmentParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UnsupportedContentAttachmentType string

const (
	UnsupportedContentAttachmentTypeUnsupportedContent UnsupportedContentAttachmentType = "UNSUPPORTED_CONTENT"
)

type CustomChannelNewParams struct {
	PublicChannelIntegrationChannelCreate PublicChannelIntegrationChannelCreateParam
	paramObj
}

func (r CustomChannelNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.PublicChannelIntegrationChannelCreate)
}
func (r *CustomChannelNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CustomChannelUpdateParams struct {
	PublicChannelIntegrationChannelPatch PublicChannelIntegrationChannelPatchParam
	paramObj
}

func (r CustomChannelUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.PublicChannelIntegrationChannelPatch)
}
func (r *CustomChannelUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CustomChannelListParams struct {
	// The paging cursor token of the last successfully read resource will be returned
	// as the `paging.next.after` JSON property of a paged response containing more
	// results.
	After             param.Opt[string] `query:"after,omitzero" json:"-"`
	DefaultPageLength param.Opt[int64]  `query:"defaultPageLength,omitzero" json:"-"`
	// The maximum number of results to display per page.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	Sort  []string         `query:"sort,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [CustomChannelListParams]'s query parameters as
// `url.Values`.
func (r CustomChannelListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type CustomChannelGetParams struct {
	ChannelID int64 `path:"channelId" api:"required" json:"-"`
	// Whether to return only results that have been archived.
	Archived param.Opt[bool] `query:"archived,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [CustomChannelGetParams]'s query parameters as `url.Values`.
func (r CustomChannelGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
