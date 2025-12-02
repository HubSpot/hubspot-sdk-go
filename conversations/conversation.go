// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package conversations

import (
	"encoding/json"
	"time"

	"github.com/stainless-sdks/hubspot-sdk-go/internal/apijson"
	"github.com/stainless-sdks/hubspot-sdk-go/option"
	"github.com/stainless-sdks/hubspot-sdk-go/packages/param"
	"github.com/stainless-sdks/hubspot-sdk-go/packages/respjson"
	"github.com/stainless-sdks/hubspot-sdk-go/shared"
)

// ConversationService contains methods and other services that help with
// interacting with the hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewConversationService] method instead.
type ConversationService struct {
	Options               []option.RequestOption
	Actors                ActorService
	ChannelAccounts       ChannelAccountService
	Channels              ChannelService
	CustomChannels        CustomChannelService
	Inboxes               InboxService
	Messages              MessageService
	Threads               ThreadService
	VisitorIdentification VisitorIdentificationService
}

// NewConversationService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewConversationService(opts ...option.RequestOption) (r ConversationService) {
	r = ConversationService{}
	r.Options = opts
	r.Actors = NewActorService(opts...)
	r.ChannelAccounts = NewChannelAccountService(opts...)
	r.Channels = NewChannelService(opts...)
	r.CustomChannels = NewCustomChannelService(opts...)
	r.Inboxes = NewInboxService(opts...)
	r.Messages = NewMessageService(opts...)
	r.Threads = NewThreadService(opts...)
	r.VisitorIdentification = NewVisitorIdentificationService(opts...)
	return
}

type AgentActor struct {
	ID string `json:"id,required"`
	// Any of "AGENT".
	Type   AgentActorType `json:"type,required"`
	Avatar string         `json:"avatar"`
	Email  string         `json:"email"`
	Name   string         `json:"name"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Type        respjson.Field
		Avatar      respjson.Field
		Email       respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AgentActor) RawJSON() string { return r.JSON.raw }
func (r *AgentActor) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AgentActorType string

const (
	AgentActorTypeAgent AgentActorType = "AGENT"
)

type BatchResponsePublicActor struct {
	CompletedAt time.Time          `json:"completedAt,required" format:"date-time"`
	Results     []PublicActorUnion `json:"results,required"`
	StartedAt   time.Time          `json:"startedAt,required" format:"date-time"`
	// Any of "CANCELED", "COMPLETE", "PENDING", "PROCESSING".
	Status      BatchResponsePublicActorStatus `json:"status,required"`
	Links       map[string]string              `json:"links"`
	RequestedAt time.Time                      `json:"requestedAt" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CompletedAt respjson.Field
		Results     respjson.Field
		StartedAt   respjson.Field
		Status      respjson.Field
		Links       respjson.Field
		RequestedAt respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BatchResponsePublicActor) RawJSON() string { return r.JSON.raw }
func (r *BatchResponsePublicActor) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BatchResponsePublicActorStatus string

const (
	BatchResponsePublicActorStatusCanceled   BatchResponsePublicActorStatus = "CANCELED"
	BatchResponsePublicActorStatusComplete   BatchResponsePublicActorStatus = "COMPLETE"
	BatchResponsePublicActorStatusPending    BatchResponsePublicActorStatus = "PENDING"
	BatchResponsePublicActorStatusProcessing BatchResponsePublicActorStatus = "PROCESSING"
)

type BotActor struct {
	ID string `json:"id,required"`
	// Any of "BOT".
	Type   BotActorType `json:"type,required"`
	Avatar string       `json:"avatar"`
	Name   string       `json:"name"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Type        respjson.Field
		Avatar      respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BotActor) RawJSON() string { return r.JSON.raw }
func (r *BotActor) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BotActorType string

const (
	BotActorTypeBot BotActorType = "BOT"
)

type CollectionResponsePublicMessageForwardPaging struct {
	Results []PublicMessageUnion `json:"results,required"`
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
func (r CollectionResponsePublicMessageForwardPaging) RawJSON() string { return r.JSON.raw }
func (r *CollectionResponsePublicMessageForwardPaging) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CollectionResponsePublicThreadForwardPaging struct {
	Results []PublicThread       `json:"results,required"`
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
func (r CollectionResponsePublicThreadForwardPaging) RawJSON() string { return r.JSON.raw }
func (r *CollectionResponsePublicThreadForwardPaging) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CollectionResponseWithTotalPublicChannelAccountForwardPaging struct {
	Results []PublicChannelAccount `json:"results,required"`
	Total   int64                  `json:"total,required"`
	Paging  shared.ForwardPaging   `json:"paging"`
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
func (r CollectionResponseWithTotalPublicChannelAccountForwardPaging) RawJSON() string {
	return r.JSON.raw
}
func (r *CollectionResponseWithTotalPublicChannelAccountForwardPaging) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CollectionResponseWithTotalPublicChannelForwardPaging struct {
	Results []PublicChannel      `json:"results,required"`
	Total   int64                `json:"total,required"`
	Paging  shared.ForwardPaging `json:"paging"`
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
func (r CollectionResponseWithTotalPublicChannelForwardPaging) RawJSON() string { return r.JSON.raw }
func (r *CollectionResponseWithTotalPublicChannelForwardPaging) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CollectionResponseWithTotalPublicInboxForwardPaging struct {
	Results []PublicInbox        `json:"results,required"`
	Total   int64                `json:"total,required"`
	Paging  shared.ForwardPaging `json:"paging"`
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
func (r CollectionResponseWithTotalPublicInboxForwardPaging) RawJSON() string { return r.JSON.raw }
func (r *CollectionResponseWithTotalPublicInboxForwardPaging) UnmarshalJSON(data []byte) error {
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

type ContactEmail struct {
	Email string `json:"email,required"`
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
	Email string `json:"email,required"`
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
	Phone string `json:"phone,required"`
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
	Phone string `json:"phone,required"`
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
	Addresses []ContactAddress `json:"addresses,required"`
	Emails    []ContactEmail   `json:"emails,required"`
	Phones    []ContactPhone   `json:"phones,required"`
	URLs      []ContactURL     `json:"urls,required"`
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
	Addresses []ContactAddressParam `json:"addresses,omitzero,required"`
	Emails    []ContactEmailParam   `json:"emails,omitzero,required"`
	Phones    []ContactPhoneParam   `json:"phones,omitzero,required"`
	URLs      []ContactURLParam     `json:"urls,omitzero,required"`
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
	URL string `json:"url,required"`
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
	URL string `json:"url,required"`
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

type ConversationsPublicConversationsMessage struct {
	ID                    string                                                   `json:"id,required"`
	Archived              bool                                                     `json:"archived,required"`
	Attachments           []ConversationsPublicConversationsMessageAttachmentUnion `json:"attachments,required"`
	ChannelAccountID      string                                                   `json:"channelAccountId,required"`
	ChannelID             string                                                   `json:"channelId,required"`
	Client                PublicClient                                             `json:"client,required"`
	ConversationsThreadID string                                                   `json:"conversationsThreadId,required"`
	CreatedAt             time.Time                                                `json:"createdAt,required" format:"date-time"`
	CreatedBy             string                                                   `json:"createdBy,required"`
	// Any of "INCOMING", "OUTGOING".
	Direction  ConversationsPublicConversationsMessageDirection `json:"direction,required"`
	Recipients []PublicRecipient                                `json:"recipients,required"`
	Senders    []PublicSender                                   `json:"senders,required"`
	Text       string                                           `json:"text,required"`
	// Any of "NOT_TRUNCATED", "TRUNCATED_TO_MOST_RECENT_REPLY", "TRUNCATED".
	TruncationStatus ConversationsPublicConversationsMessageTruncationStatus `json:"truncationStatus,required"`
	// Any of "MESSAGE".
	Type        ConversationsPublicConversationsMessageType `json:"type,required"`
	InReplyToID string                                      `json:"inReplyToId"`
	RichText    string                                      `json:"richText"`
	Status      PublicMessageStatus                         `json:"status"`
	Subject     string                                      `json:"subject"`
	UpdatedAt   time.Time                                   `json:"updatedAt" format:"date-time"`
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
func (r ConversationsPublicConversationsMessage) RawJSON() string { return r.JSON.raw }
func (r *ConversationsPublicConversationsMessage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ConversationsPublicConversationsMessageAttachmentUnion contains all possible
// properties and values from [PublicFile], [PublicLocation], [PublicContact],
// [PublicUnsupportedContent], [PublicMessageHeader], [PublicQuickReplies],
// [PublicWhatsAppTemplateMetadata], [PublicSocialMetadataAttachment].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ConversationsPublicConversationsMessageAttachmentUnion struct {
	// This field is a union of [string], [int64]
	FileID ConversationsPublicConversationsMessageAttachmentUnionFileID `json:"fileId"`
	// This field is from variant [PublicFile].
	FileUsageType string `json:"fileUsageType"`
	Type          string `json:"type"`
	Name          string `json:"name"`
	URL           string `json:"url"`
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
	MappedTemplateID string `json:"mappedTemplateId"`
	// This field is from variant [PublicWhatsAppTemplateMetadata].
	Parameters map[string]string `json:"parameters"`
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
		MappedTemplateID respjson.Field
		Parameters       respjson.Field
		SocialMetadata   respjson.Field
		raw              string
	} `json:"-"`
}

func (u ConversationsPublicConversationsMessageAttachmentUnion) AsFile() (v PublicFile) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConversationsPublicConversationsMessageAttachmentUnion) AsLocation() (v PublicLocation) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConversationsPublicConversationsMessageAttachmentUnion) AsContact() (v PublicContact) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConversationsPublicConversationsMessageAttachmentUnion) AsUnsupportedContent() (v PublicUnsupportedContent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConversationsPublicConversationsMessageAttachmentUnion) AsMessageHeader() (v PublicMessageHeader) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConversationsPublicConversationsMessageAttachmentUnion) AsQuickReplies() (v PublicQuickReplies) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConversationsPublicConversationsMessageAttachmentUnion) AsWhatsappTemplateMetadata() (v PublicWhatsAppTemplateMetadata) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConversationsPublicConversationsMessageAttachmentUnion) AsSocialMediaMetadata() (v PublicSocialMetadataAttachment) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ConversationsPublicConversationsMessageAttachmentUnion) RawJSON() string { return u.JSON.raw }

func (r *ConversationsPublicConversationsMessageAttachmentUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ConversationsPublicConversationsMessageAttachmentUnionFileID is an implicit
// subunion of [ConversationsPublicConversationsMessageAttachmentUnion].
// ConversationsPublicConversationsMessageAttachmentUnionFileID provides convenient
// access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [ConversationsPublicConversationsMessageAttachmentUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfInt]
type ConversationsPublicConversationsMessageAttachmentUnionFileID struct {
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

func (r *ConversationsPublicConversationsMessageAttachmentUnionFileID) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConversationsPublicConversationsMessageDirection string

const (
	ConversationsPublicConversationsMessageDirectionIncoming ConversationsPublicConversationsMessageDirection = "INCOMING"
	ConversationsPublicConversationsMessageDirectionOutgoing ConversationsPublicConversationsMessageDirection = "OUTGOING"
)

type ConversationsPublicConversationsMessageTruncationStatus string

const (
	ConversationsPublicConversationsMessageTruncationStatusNotTruncated               ConversationsPublicConversationsMessageTruncationStatus = "NOT_TRUNCATED"
	ConversationsPublicConversationsMessageTruncationStatusTruncatedToMostRecentReply ConversationsPublicConversationsMessageTruncationStatus = "TRUNCATED_TO_MOST_RECENT_REPLY"
	ConversationsPublicConversationsMessageTruncationStatusTruncated                  ConversationsPublicConversationsMessageTruncationStatus = "TRUNCATED"
)

type ConversationsPublicConversationsMessageType string

const (
	ConversationsPublicConversationsMessageTypeMessage ConversationsPublicConversationsMessageType = "MESSAGE"
)

type EmailActor struct {
	ID    string `json:"id,required"`
	Email string `json:"email,required"`
	// Any of "EMAIL".
	Type EmailActorType `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Email       respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmailActor) RawJSON() string { return r.JSON.raw }
func (r *EmailActor) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmailActorType string

const (
	EmailActorTypeEmail EmailActorType = "EMAIL"
)

type IntegratorActor struct {
	ID   string `json:"id,required"`
	Name string `json:"name,required"`
	// Any of "INTEGRATOR".
	Type   IntegratorActorType `json:"type,required"`
	Avatar string              `json:"avatar"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Name        respjson.Field
		Type        respjson.Field
		Avatar      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IntegratorActor) RawJSON() string { return r.JSON.raw }
func (r *IntegratorActor) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IntegratorActorType string

const (
	IntegratorActorTypeIntegrator IntegratorActorType = "INTEGRATOR"
)

type LlmActor struct {
	ID string `json:"id,required"`
	// Any of "LLM".
	Type   LlmActorType `json:"type,required"`
	Avatar string       `json:"avatar"`
	Name   string       `json:"name"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Type        respjson.Field
		Avatar      respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LlmActor) RawJSON() string { return r.JSON.raw }
func (r *LlmActor) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LlmActorType string

const (
	LlmActorTypeLlm LlmActorType = "LLM"
)

// PublicActorUnion contains all possible properties and values from [AgentActor],
// [BotActor], [IntegratorActor], [SystemActor], [VisitorActor], [EmailActor],
// [LlmActor].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type PublicActorUnion struct {
	ID     string `json:"id"`
	Type   string `json:"type"`
	Avatar string `json:"avatar"`
	Email  string `json:"email"`
	Name   string `json:"name"`
	JSON   struct {
		ID     respjson.Field
		Type   respjson.Field
		Avatar respjson.Field
		Email  respjson.Field
		Name   respjson.Field
		raw    string
	} `json:"-"`
}

func (u PublicActorUnion) AsAgentActor() (v AgentActor) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicActorUnion) AsBotActor() (v BotActor) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicActorUnion) AsIntegratorActor() (v IntegratorActor) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicActorUnion) AsSystemActor() (v SystemActor) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicActorUnion) AsVisitorActor() (v VisitorActor) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicActorUnion) AsEmailActor() (v EmailActor) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicActorUnion) AsLlmActor() (v LlmActor) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u PublicActorUnion) RawJSON() string { return u.JSON.raw }

func (r *PublicActorUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PublicAssignmentMessage struct {
	ID                    string            `json:"id,required"`
	Archived              bool              `json:"archived,required"`
	Client                PublicClient      `json:"client,required"`
	ConversationsThreadID string            `json:"conversationsThreadId,required"`
	CreatedAt             time.Time         `json:"createdAt,required" format:"date-time"`
	CreatedBy             string            `json:"createdBy,required"`
	Recipients            []PublicRecipient `json:"recipients,required"`
	Senders               []PublicSender    `json:"senders,required"`
	// Any of "ASSIGNMENT".
	Type         PublicAssignmentMessageType `json:"type,required"`
	AssignedFrom string                      `json:"assignedFrom"`
	AssignedTo   string                      `json:"assignedTo"`
	UpdatedAt    time.Time                   `json:"updatedAt" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                    respjson.Field
		Archived              respjson.Field
		Client                respjson.Field
		ConversationsThreadID respjson.Field
		CreatedAt             respjson.Field
		CreatedBy             respjson.Field
		Recipients            respjson.Field
		Senders               respjson.Field
		Type                  respjson.Field
		AssignedFrom          respjson.Field
		AssignedTo            respjson.Field
		UpdatedAt             respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PublicAssignmentMessage) RawJSON() string { return r.JSON.raw }
func (r *PublicAssignmentMessage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PublicAssignmentMessageType string

const (
	PublicAssignmentMessageTypeAssignment PublicAssignmentMessageType = "ASSIGNMENT"
)

type PublicChannel struct {
	ID   string `json:"id,required"`
	Name string `json:"name,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PublicChannel) RawJSON() string { return r.JSON.raw }
func (r *PublicChannel) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PublicChannelAccount struct {
	ID                 string                   `json:"id,required"`
	Active             bool                     `json:"active,required"`
	Archived           bool                     `json:"archived,required"`
	Authorized         bool                     `json:"authorized,required"`
	ChannelID          string                   `json:"channelId,required"`
	CreatedAt          time.Time                `json:"createdAt,required" format:"date-time"`
	InboxID            string                   `json:"inboxId,required"`
	Name               string                   `json:"name,required"`
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

type PublicClient struct {
	// Any of "HUBSPOT", "INTEGRATION", "SYSTEM", "UNKNOWN".
	ClientType       PublicClientClientType `json:"clientType,required"`
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
	PublicClientClientTypeHubspot     PublicClientClientType = "HUBSPOT"
	PublicClientClientTypeIntegration PublicClientClientType = "INTEGRATION"
	PublicClientClientTypeSystem      PublicClientClientType = "SYSTEM"
	PublicClientClientTypeUnknown     PublicClientClientType = "UNKNOWN"
)

type PublicComment struct {
	ID                    string                         `json:"id,required"`
	Archived              bool                           `json:"archived,required"`
	Attachments           []PublicCommentAttachmentUnion `json:"attachments,required"`
	Client                PublicClient                   `json:"client,required"`
	ConversationsThreadID string                         `json:"conversationsThreadId,required"`
	CreatedAt             time.Time                      `json:"createdAt,required" format:"date-time"`
	CreatedBy             string                         `json:"createdBy,required"`
	Recipients            []PublicRecipient              `json:"recipients,required"`
	RichText              string                         `json:"richText,required"`
	Senders               []PublicSender                 `json:"senders,required"`
	Text                  string                         `json:"text,required"`
	// Any of "COMMENT".
	Type      PublicCommentType `json:"type,required"`
	UpdatedAt time.Time         `json:"updatedAt" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                    respjson.Field
		Archived              respjson.Field
		Attachments           respjson.Field
		Client                respjson.Field
		ConversationsThreadID respjson.Field
		CreatedAt             respjson.Field
		CreatedBy             respjson.Field
		Recipients            respjson.Field
		RichText              respjson.Field
		Senders               respjson.Field
		Text                  respjson.Field
		Type                  respjson.Field
		UpdatedAt             respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PublicComment) RawJSON() string { return r.JSON.raw }
func (r *PublicComment) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PublicCommentAttachmentUnion contains all possible properties and values from
// [PublicFile], [PublicLocation], [PublicContact], [PublicUnsupportedContent],
// [PublicMessageHeader], [PublicQuickReplies], [PublicWhatsAppTemplateMetadata],
// [PublicSocialMetadataAttachment].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type PublicCommentAttachmentUnion struct {
	// This field is a union of [string], [int64]
	FileID PublicCommentAttachmentUnionFileID `json:"fileId"`
	// This field is from variant [PublicFile].
	FileUsageType string `json:"fileUsageType"`
	Type          string `json:"type"`
	Name          string `json:"name"`
	URL           string `json:"url"`
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
	MappedTemplateID string `json:"mappedTemplateId"`
	// This field is from variant [PublicWhatsAppTemplateMetadata].
	Parameters map[string]string `json:"parameters"`
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
		MappedTemplateID respjson.Field
		Parameters       respjson.Field
		SocialMetadata   respjson.Field
		raw              string
	} `json:"-"`
}

func (u PublicCommentAttachmentUnion) AsFile() (v PublicFile) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicCommentAttachmentUnion) AsLocation() (v PublicLocation) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicCommentAttachmentUnion) AsContact() (v PublicContact) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicCommentAttachmentUnion) AsUnsupportedContent() (v PublicUnsupportedContent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicCommentAttachmentUnion) AsMessageHeader() (v PublicMessageHeader) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicCommentAttachmentUnion) AsQuickReplies() (v PublicQuickReplies) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicCommentAttachmentUnion) AsWhatsappTemplateMetadata() (v PublicWhatsAppTemplateMetadata) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicCommentAttachmentUnion) AsSocialMediaMetadata() (v PublicSocialMetadataAttachment) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u PublicCommentAttachmentUnion) RawJSON() string { return u.JSON.raw }

func (r *PublicCommentAttachmentUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PublicCommentAttachmentUnionFileID is an implicit subunion of
// [PublicCommentAttachmentUnion]. PublicCommentAttachmentUnionFileID provides
// convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [PublicCommentAttachmentUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfInt]
type PublicCommentAttachmentUnionFileID struct {
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

func (r *PublicCommentAttachmentUnionFileID) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PublicCommentType string

const (
	PublicCommentTypeComment PublicCommentType = "COMMENT"
)

// The properties Attachments, Text, Type are required.
type PublicCommentEggParam struct {
	Attachments []PublicCommentEggAttachmentUnionParam `json:"attachments,omitzero,required"`
	Text        string                                 `json:"text,required"`
	// Any of "COMMENT".
	Type     PublicCommentEggType `json:"type,omitzero,required"`
	RichText param.Opt[string]    `json:"richText,omitzero"`
	paramObj
}

func (r PublicCommentEggParam) MarshalJSON() (data []byte, err error) {
	type shadow PublicCommentEggParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PublicCommentEggParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type PublicCommentEggAttachmentUnionParam struct {
	OfFile                *PublicFileEggParam         `json:",omitzero,inline"`
	OfQuickReplies        *PublicQuickRepliesEggParam `json:",omitzero,inline"`
	OfSocialMediaMetadata *PublicSocialMediaEggParam  `json:",omitzero,inline"`
	paramUnion
}

func (u PublicCommentEggAttachmentUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfFile, u.OfQuickReplies, u.OfSocialMediaMetadata)
}
func (u *PublicCommentEggAttachmentUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *PublicCommentEggAttachmentUnionParam) asAny() any {
	if !param.IsOmitted(u.OfFile) {
		return u.OfFile
	} else if !param.IsOmitted(u.OfQuickReplies) {
		return u.OfQuickReplies
	} else if !param.IsOmitted(u.OfSocialMediaMetadata) {
		return u.OfSocialMediaMetadata
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u PublicCommentEggAttachmentUnionParam) GetFileID() *string {
	if vt := u.OfFile; vt != nil {
		return &vt.FileID
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u PublicCommentEggAttachmentUnionParam) GetQuickReplies() []QuickReplyParam {
	if vt := u.OfQuickReplies; vt != nil {
		return vt.QuickReplies
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u PublicCommentEggAttachmentUnionParam) GetSocialMetadata() *SocialMetadataParam {
	if vt := u.OfSocialMediaMetadata; vt != nil {
		return &vt.SocialMetadata
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u PublicCommentEggAttachmentUnionParam) GetType() *string {
	if vt := u.OfFile; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfQuickReplies; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfSocialMediaMetadata; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

type PublicCommentEggType string

const (
	PublicCommentEggTypeComment PublicCommentEggType = "COMMENT"
)

type PublicContact struct {
	ContactProfile ContactProfile `json:"contactProfile,required"`
	// Any of "CONTACT".
	Type PublicContactType `json:"type,required"`
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

// The properties Attachments, ChannelAccountID, ChannelID, Recipients,
// SenderActorID, Text, Type are required.
type PublicConversationsMessageEggParam struct {
	Attachments      []PublicConversationsMessageEggAttachmentUnionParam `json:"attachments,omitzero,required"`
	ChannelAccountID string                                              `json:"channelAccountId,required"`
	ChannelID        string                                              `json:"channelId,required"`
	Recipients       []PublicRecipientEggParam                           `json:"recipients,omitzero,required"`
	SenderActorID    string                                              `json:"senderActorId,required"`
	Text             string                                              `json:"text,required"`
	// Any of "MESSAGE".
	Type     PublicConversationsMessageEggType `json:"type,omitzero,required"`
	RichText param.Opt[string]                 `json:"richText,omitzero"`
	Subject  param.Opt[string]                 `json:"subject,omitzero"`
	paramObj
}

func (r PublicConversationsMessageEggParam) MarshalJSON() (data []byte, err error) {
	type shadow PublicConversationsMessageEggParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PublicConversationsMessageEggParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type PublicConversationsMessageEggAttachmentUnionParam struct {
	OfFile                *PublicFileEggParam         `json:",omitzero,inline"`
	OfQuickReplies        *PublicQuickRepliesEggParam `json:",omitzero,inline"`
	OfSocialMediaMetadata *PublicSocialMediaEggParam  `json:",omitzero,inline"`
	paramUnion
}

func (u PublicConversationsMessageEggAttachmentUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfFile, u.OfQuickReplies, u.OfSocialMediaMetadata)
}
func (u *PublicConversationsMessageEggAttachmentUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *PublicConversationsMessageEggAttachmentUnionParam) asAny() any {
	if !param.IsOmitted(u.OfFile) {
		return u.OfFile
	} else if !param.IsOmitted(u.OfQuickReplies) {
		return u.OfQuickReplies
	} else if !param.IsOmitted(u.OfSocialMediaMetadata) {
		return u.OfSocialMediaMetadata
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u PublicConversationsMessageEggAttachmentUnionParam) GetFileID() *string {
	if vt := u.OfFile; vt != nil {
		return &vt.FileID
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u PublicConversationsMessageEggAttachmentUnionParam) GetQuickReplies() []QuickReplyParam {
	if vt := u.OfQuickReplies; vt != nil {
		return vt.QuickReplies
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u PublicConversationsMessageEggAttachmentUnionParam) GetSocialMetadata() *SocialMetadataParam {
	if vt := u.OfSocialMediaMetadata; vt != nil {
		return &vt.SocialMetadata
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u PublicConversationsMessageEggAttachmentUnionParam) GetType() *string {
	if vt := u.OfFile; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfQuickReplies; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfSocialMediaMetadata; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

type PublicConversationsMessageEggType string

const (
	PublicConversationsMessageEggTypeMessage PublicConversationsMessageEggType = "MESSAGE"
)

type PublicDeliveryIdentifier struct {
	Type  string `json:"type,required"`
	Value string `json:"value,required"`
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

// The properties Type, Value are required.
type PublicDeliveryIdentifierParam struct {
	Type  string `json:"type,required"`
	Value string `json:"value,required"`
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
	FileID        string `json:"fileId,required"`
	FileUsageType string `json:"fileUsageType,required"`
	// Any of "FILE".
	Type PublicFileType `json:"type,required"`
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

type PublicFileType string

const (
	PublicFileTypeFile PublicFileType = "FILE"
)

// The properties FileID, Type are required.
type PublicFileEggParam struct {
	FileID string `json:"fileId,required"`
	// Any of "FILE".
	Type PublicFileEggType `json:"type,omitzero,required"`
	paramObj
}

func (r PublicFileEggParam) MarshalJSON() (data []byte, err error) {
	type shadow PublicFileEggParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PublicFileEggParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PublicFileEggType string

const (
	PublicFileEggTypeFile PublicFileEggType = "FILE"
)

type PublicInbox struct {
	ID         string    `json:"id,required"`
	Archived   bool      `json:"archived,required"`
	CreatedAt  time.Time `json:"createdAt,required" format:"date-time"`
	Name       string    `json:"name,required"`
	Type       string    `json:"type,required"`
	UpdatedAt  time.Time `json:"updatedAt,required" format:"date-time"`
	ArchivedAt time.Time `json:"archivedAt" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Archived    respjson.Field
		CreatedAt   respjson.Field
		Name        respjson.Field
		Type        respjson.Field
		UpdatedAt   respjson.Field
		ArchivedAt  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PublicInbox) RawJSON() string { return r.JSON.raw }
func (r *PublicInbox) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PublicLocation struct {
	Latitude  float64 `json:"latitude,required"`
	Longitude float64 `json:"longitude,required"`
	// Any of "LOCATION".
	Type    PublicLocationType `json:"type,required"`
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

// PublicMessageUnion contains all possible properties and values from
// [ConversationsPublicConversationsMessage], [PublicComment],
// [PublicWelcomeMessage], [PublicAssignmentMessage], [PublicThreadStatusChange],
// [PublicThreadInboxChange].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type PublicMessageUnion struct {
	ID       string `json:"id"`
	Archived bool   `json:"archived"`
	// This field is a union of
	// [[]ConversationsPublicConversationsMessageAttachmentUnion],
	// [[]PublicCommentAttachmentUnion]
	Attachments      PublicMessageUnionAttachments `json:"attachments"`
	ChannelAccountID string                        `json:"channelAccountId"`
	ChannelID        string                        `json:"channelId"`
	// This field is from variant [ConversationsPublicConversationsMessage].
	Client                PublicClient `json:"client"`
	ConversationsThreadID string       `json:"conversationsThreadId"`
	CreatedAt             time.Time    `json:"createdAt"`
	CreatedBy             string       `json:"createdBy"`
	// This field is from variant [ConversationsPublicConversationsMessage].
	Direction  ConversationsPublicConversationsMessageDirection `json:"direction"`
	Recipients []PublicRecipient                                `json:"recipients"`
	Senders    []PublicSender                                   `json:"senders"`
	Text       string                                           `json:"text"`
	// This field is from variant [ConversationsPublicConversationsMessage].
	TruncationStatus ConversationsPublicConversationsMessageTruncationStatus `json:"truncationStatus"`
	Type             string                                                  `json:"type"`
	// This field is from variant [ConversationsPublicConversationsMessage].
	InReplyToID string `json:"inReplyToId"`
	RichText    string `json:"richText"`
	// This field is from variant [ConversationsPublicConversationsMessage].
	Status PublicMessageStatus `json:"status"`
	// This field is from variant [ConversationsPublicConversationsMessage].
	Subject   string    `json:"subject"`
	UpdatedAt time.Time `json:"updatedAt"`
	// This field is from variant [PublicAssignmentMessage].
	AssignedFrom string `json:"assignedFrom"`
	// This field is from variant [PublicAssignmentMessage].
	AssignedTo string `json:"assignedTo"`
	// This field is from variant [PublicThreadStatusChange].
	NewStatus PublicThreadStatusChangeNewStatus `json:"newStatus"`
	// This field is from variant [PublicThreadInboxChange].
	FromInboxID string `json:"fromInboxId"`
	// This field is from variant [PublicThreadInboxChange].
	ToInboxID string `json:"toInboxId"`
	JSON      struct {
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
		AssignedFrom          respjson.Field
		AssignedTo            respjson.Field
		NewStatus             respjson.Field
		FromInboxID           respjson.Field
		ToInboxID             respjson.Field
		raw                   string
	} `json:"-"`
}

func (u PublicMessageUnion) AsConversationsPublicConversationsMessage() (v ConversationsPublicConversationsMessage) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicMessageUnion) AsPublicComment() (v PublicComment) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicMessageUnion) AsPublicWelcomeMessage() (v PublicWelcomeMessage) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicMessageUnion) AsPublicAssignmentMessage() (v PublicAssignmentMessage) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicMessageUnion) AsPublicThreadStatusChange() (v PublicThreadStatusChange) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PublicMessageUnion) AsPublicThreadInboxChange() (v PublicThreadInboxChange) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u PublicMessageUnion) RawJSON() string { return u.JSON.raw }

func (r *PublicMessageUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PublicMessageUnionAttachments is an implicit subunion of [PublicMessageUnion].
// PublicMessageUnionAttachments provides convenient access to the sub-properties
// of the union.
//
// For type safety it is recommended to directly use a variant of the
// [PublicMessageUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfConversationsPublicConversationsMessageAttachments
// OfPublicCommentAttachments]
type PublicMessageUnionAttachments struct {
	// This field will be present if the value is a
	// [[]ConversationsPublicConversationsMessageAttachmentUnion] instead of an object.
	OfConversationsPublicConversationsMessageAttachments []ConversationsPublicConversationsMessageAttachmentUnion `json:",inline"`
	// This field will be present if the value is a [[]PublicCommentAttachmentUnion]
	// instead of an object.
	OfPublicCommentAttachments []PublicCommentAttachmentUnion `json:",inline"`
	JSON                       struct {
		OfConversationsPublicConversationsMessageAttachments respjson.Field
		OfPublicCommentAttachments                           respjson.Field
		raw                                                  string
	} `json:"-"`
}

func (r *PublicMessageUnionAttachments) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PublicMessageContent struct {
	RichText string `json:"richText"`
	Text     string `json:"text"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		RichText    respjson.Field
		Text        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PublicMessageContent) RawJSON() string { return r.JSON.raw }
func (r *PublicMessageContent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func PublicMessageEggParamOfPublicCommentEgg(attachments []PublicCommentEggAttachmentUnionParam, text string, type_ PublicCommentEggType) PublicMessageEggUnionParam {
	var variant PublicCommentEggParam
	variant.Attachments = attachments
	variant.Text = text
	variant.Type = type_
	return PublicMessageEggUnionParam{OfPublicCommentEgg: &variant}
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type PublicMessageEggUnionParam struct {
	OfPublicConversationsMessageEgg *PublicConversationsMessageEggParam `json:",omitzero,inline"`
	OfPublicCommentEgg              *PublicCommentEggParam              `json:",omitzero,inline"`
	paramUnion
}

func (u PublicMessageEggUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfPublicConversationsMessageEgg, u.OfPublicCommentEgg)
}
func (u *PublicMessageEggUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *PublicMessageEggUnionParam) asAny() any {
	if !param.IsOmitted(u.OfPublicConversationsMessageEgg) {
		return u.OfPublicConversationsMessageEgg
	} else if !param.IsOmitted(u.OfPublicCommentEgg) {
		return u.OfPublicCommentEgg
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u PublicMessageEggUnionParam) GetChannelAccountID() *string {
	if vt := u.OfPublicConversationsMessageEgg; vt != nil {
		return &vt.ChannelAccountID
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u PublicMessageEggUnionParam) GetChannelID() *string {
	if vt := u.OfPublicConversationsMessageEgg; vt != nil {
		return &vt.ChannelID
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u PublicMessageEggUnionParam) GetRecipients() []PublicRecipientEggParam {
	if vt := u.OfPublicConversationsMessageEgg; vt != nil {
		return vt.Recipients
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u PublicMessageEggUnionParam) GetSenderActorID() *string {
	if vt := u.OfPublicConversationsMessageEgg; vt != nil {
		return &vt.SenderActorID
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u PublicMessageEggUnionParam) GetSubject() *string {
	if vt := u.OfPublicConversationsMessageEgg; vt != nil && vt.Subject.Valid() {
		return &vt.Subject.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u PublicMessageEggUnionParam) GetText() *string {
	if vt := u.OfPublicConversationsMessageEgg; vt != nil {
		return (*string)(&vt.Text)
	} else if vt := u.OfPublicCommentEgg; vt != nil {
		return (*string)(&vt.Text)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u PublicMessageEggUnionParam) GetType() *string {
	if vt := u.OfPublicConversationsMessageEgg; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfPublicCommentEgg; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u PublicMessageEggUnionParam) GetRichText() *string {
	if vt := u.OfPublicConversationsMessageEgg; vt != nil && vt.RichText.Valid() {
		return &vt.RichText.Value
	} else if vt := u.OfPublicCommentEgg; vt != nil && vt.RichText.Valid() {
		return &vt.RichText.Value
	}
	return nil
}

// Returns a subunion which exports methods to access subproperties
//
// Or use AsAny() to get the underlying value
func (u PublicMessageEggUnionParam) GetAttachments() (res publicMessageEggUnionParamAttachments) {
	if vt := u.OfPublicConversationsMessageEgg; vt != nil {
		res.any = &vt.Attachments
	} else if vt := u.OfPublicCommentEgg; vt != nil {
		res.any = &vt.Attachments
	}
	return
}

// Can have the runtime types
// [_[]PublicConversationsMessageEggAttachmentUnionParam],
// [_[]PublicCommentEggAttachmentUnionParam]
type publicMessageEggUnionParamAttachments struct{ any }

// Use the following switch statement to get the type of the union:
//
//	switch u.AsAny().(type) {
//	case *[]conversations.PublicConversationsMessageEggAttachmentUnionParam:
//	case *[]conversations.PublicCommentEggAttachmentUnionParam:
//	default:
//	    fmt.Errorf("not present")
//	}
func (u publicMessageEggUnionParamAttachments) AsAny() any { return u.any }

type PublicMessageFailureDetails struct {
	ErrorMessageTokens map[string]string `json:"errorMessageTokens,required"`
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
	Type   PublicMessageHeaderType `json:"type,required"`
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
	StatusType     PublicMessageStatusStatusType `json:"statusType,required"`
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
	AllowMultiSelect bool         `json:"allowMultiSelect,required"`
	AllowUserInput   bool         `json:"allowUserInput,required"`
	QuickReplies     []QuickReply `json:"quickReplies,required"`
	// Any of "QUICK_REPLIES".
	Type PublicQuickRepliesType `json:"type,required"`
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

// The properties QuickReplies, Type are required.
type PublicQuickRepliesEggParam struct {
	QuickReplies []QuickReplyParam `json:"quickReplies,omitzero,required"`
	// Any of "QUICK_REPLIES".
	Type PublicQuickRepliesEggType `json:"type,omitzero,required"`
	paramObj
}

func (r PublicQuickRepliesEggParam) MarshalJSON() (data []byte, err error) {
	type shadow PublicQuickRepliesEggParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PublicQuickRepliesEggParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PublicQuickRepliesEggType string

const (
	PublicQuickRepliesEggTypeQuickReplies PublicQuickRepliesEggType = "QUICK_REPLIES"
)

type PublicRecipient struct {
	DeliveryIdentifier PublicDeliveryIdentifier `json:"deliveryIdentifier,required"`
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

// The property DeliveryIdentifiers is required.
type PublicRecipientEggParam struct {
	DeliveryIdentifiers []PublicDeliveryIdentifierParam `json:"deliveryIdentifiers,omitzero,required"`
	ActorID             param.Opt[string]               `json:"actorId,omitzero"`
	Name                param.Opt[string]               `json:"name,omitzero"`
	RecipientField      param.Opt[string]               `json:"recipientField,omitzero"`
	DeliveryIdentifier  PublicDeliveryIdentifierParam   `json:"deliveryIdentifier,omitzero"`
	paramObj
}

func (r PublicRecipientEggParam) MarshalJSON() (data []byte, err error) {
	type shadow PublicRecipientEggParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PublicRecipientEggParam) UnmarshalJSON(data []byte) error {
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

// The properties SocialMetadata, Type are required.
type PublicSocialMediaEggParam struct {
	SocialMetadata SocialMetadataParam `json:"socialMetadata,omitzero,required"`
	// Any of "SOCIAL_MEDIA_METADATA".
	Type PublicSocialMediaEggType `json:"type,omitzero,required"`
	paramObj
}

func (r PublicSocialMediaEggParam) MarshalJSON() (data []byte, err error) {
	type shadow PublicSocialMediaEggParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PublicSocialMediaEggParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PublicSocialMediaEggType string

const (
	PublicSocialMediaEggTypeSocialMediaMetadata PublicSocialMediaEggType = "SOCIAL_MEDIA_METADATA"
)

type PublicSocialMetadataAttachment struct {
	SocialMetadata SocialMetadata `json:"socialMetadata,required"`
	// Any of "SOCIAL_MEDIA_METADATA".
	Type PublicSocialMetadataAttachmentType `json:"type,required"`
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

type PublicThread struct {
	ID                       string    `json:"id,required"`
	Archived                 bool      `json:"archived,required"`
	AssociatedContactID      string    `json:"associatedContactId,required"`
	CreatedAt                time.Time `json:"createdAt,required" format:"date-time"`
	InboxID                  string    `json:"inboxId,required"`
	OriginalChannelAccountID string    `json:"originalChannelAccountId,required"`
	OriginalChannelID        string    `json:"originalChannelId,required"`
	Spam                     bool      `json:"spam,required"`
	// Any of "CLOSED", "OPEN".
	Status                         PublicThreadStatus       `json:"status,required"`
	AssignedTo                     string                   `json:"assignedTo"`
	ClosedAt                       time.Time                `json:"closedAt" format:"date-time"`
	LatestMessageReceivedTimestamp time.Time                `json:"latestMessageReceivedTimestamp" format:"date-time"`
	LatestMessageSentTimestamp     time.Time                `json:"latestMessageSentTimestamp" format:"date-time"`
	LatestMessageTimestamp         time.Time                `json:"latestMessageTimestamp" format:"date-time"`
	ThreadAssociations             PublicThreadAssociations `json:"threadAssociations"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                             respjson.Field
		Archived                       respjson.Field
		AssociatedContactID            respjson.Field
		CreatedAt                      respjson.Field
		InboxID                        respjson.Field
		OriginalChannelAccountID       respjson.Field
		OriginalChannelID              respjson.Field
		Spam                           respjson.Field
		Status                         respjson.Field
		AssignedTo                     respjson.Field
		ClosedAt                       respjson.Field
		LatestMessageReceivedTimestamp respjson.Field
		LatestMessageSentTimestamp     respjson.Field
		LatestMessageTimestamp         respjson.Field
		ThreadAssociations             respjson.Field
		ExtraFields                    map[string]respjson.Field
		raw                            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PublicThread) RawJSON() string { return r.JSON.raw }
func (r *PublicThread) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PublicThreadStatus string

const (
	PublicThreadStatusClosed PublicThreadStatus = "CLOSED"
	PublicThreadStatusOpen   PublicThreadStatus = "OPEN"
)

type PublicThreadAssociations struct {
	AssociatedTicketID string `json:"associatedTicketId"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AssociatedTicketID respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PublicThreadAssociations) RawJSON() string { return r.JSON.raw }
func (r *PublicThreadAssociations) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PublicThreadInboxChange struct {
	ID                    string            `json:"id,required"`
	Archived              bool              `json:"archived,required"`
	Client                PublicClient      `json:"client,required"`
	ConversationsThreadID string            `json:"conversationsThreadId,required"`
	CreatedAt             time.Time         `json:"createdAt,required" format:"date-time"`
	CreatedBy             string            `json:"createdBy,required"`
	FromInboxID           string            `json:"fromInboxId,required"`
	Recipients            []PublicRecipient `json:"recipients,required"`
	Senders               []PublicSender    `json:"senders,required"`
	ToInboxID             string            `json:"toInboxId,required"`
	// Any of "THREAD_INBOX_CHANGE".
	Type      PublicThreadInboxChangeType `json:"type,required"`
	UpdatedAt time.Time                   `json:"updatedAt" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                    respjson.Field
		Archived              respjson.Field
		Client                respjson.Field
		ConversationsThreadID respjson.Field
		CreatedAt             respjson.Field
		CreatedBy             respjson.Field
		FromInboxID           respjson.Field
		Recipients            respjson.Field
		Senders               respjson.Field
		ToInboxID             respjson.Field
		Type                  respjson.Field
		UpdatedAt             respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PublicThreadInboxChange) RawJSON() string { return r.JSON.raw }
func (r *PublicThreadInboxChange) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PublicThreadInboxChangeType string

const (
	PublicThreadInboxChangeTypeThreadInboxChange PublicThreadInboxChangeType = "THREAD_INBOX_CHANGE"
)

type PublicThreadStatusChange struct {
	ID                    string       `json:"id,required"`
	Archived              bool         `json:"archived,required"`
	Client                PublicClient `json:"client,required"`
	ConversationsThreadID string       `json:"conversationsThreadId,required"`
	CreatedAt             time.Time    `json:"createdAt,required" format:"date-time"`
	CreatedBy             string       `json:"createdBy,required"`
	// Any of "OPEN", "CLOSED".
	NewStatus  PublicThreadStatusChangeNewStatus `json:"newStatus,required"`
	Recipients []PublicRecipient                 `json:"recipients,required"`
	Senders    []PublicSender                    `json:"senders,required"`
	// Any of "THREAD_STATUS_CHANGE".
	Type      PublicThreadStatusChangeType `json:"type,required"`
	UpdatedAt time.Time                    `json:"updatedAt" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                    respjson.Field
		Archived              respjson.Field
		Client                respjson.Field
		ConversationsThreadID respjson.Field
		CreatedAt             respjson.Field
		CreatedBy             respjson.Field
		NewStatus             respjson.Field
		Recipients            respjson.Field
		Senders               respjson.Field
		Type                  respjson.Field
		UpdatedAt             respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PublicThreadStatusChange) RawJSON() string { return r.JSON.raw }
func (r *PublicThreadStatusChange) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PublicThreadStatusChangeNewStatus string

const (
	PublicThreadStatusChangeNewStatusOpen   PublicThreadStatusChangeNewStatus = "OPEN"
	PublicThreadStatusChangeNewStatusClosed PublicThreadStatusChangeNewStatus = "CLOSED"
)

type PublicThreadStatusChangeType string

const (
	PublicThreadStatusChangeTypeThreadStatusChange PublicThreadStatusChangeType = "THREAD_STATUS_CHANGE"
)

type PublicThreadUpdateRequestParam struct {
	Archived param.Opt[bool] `json:"archived,omitzero"`
	// Any of "CLOSED", "OPEN".
	Status PublicThreadUpdateRequestStatus `json:"status,omitzero"`
	paramObj
}

func (r PublicThreadUpdateRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow PublicThreadUpdateRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PublicThreadUpdateRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PublicThreadUpdateRequestStatus string

const (
	PublicThreadUpdateRequestStatusClosed PublicThreadUpdateRequestStatus = "CLOSED"
	PublicThreadUpdateRequestStatusOpen   PublicThreadUpdateRequestStatus = "OPEN"
)

type PublicUnsupportedContent struct {
	// Any of "UNSUPPORTED_CONTENT".
	Type PublicUnsupportedContentType `json:"type,required"`
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

type PublicWelcomeMessage struct {
	ID                    string            `json:"id,required"`
	Archived              bool              `json:"archived,required"`
	ChannelAccountID      string            `json:"channelAccountId,required"`
	ChannelID             string            `json:"channelId,required"`
	Client                PublicClient      `json:"client,required"`
	ConversationsThreadID string            `json:"conversationsThreadId,required"`
	CreatedAt             time.Time         `json:"createdAt,required" format:"date-time"`
	CreatedBy             string            `json:"createdBy,required"`
	Recipients            []PublicRecipient `json:"recipients,required"`
	Senders               []PublicSender    `json:"senders,required"`
	Text                  string            `json:"text,required"`
	// Any of "WELCOME_MESSAGE".
	Type      PublicWelcomeMessageType `json:"type,required"`
	RichText  string                   `json:"richText"`
	UpdatedAt time.Time                `json:"updatedAt" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                    respjson.Field
		Archived              respjson.Field
		ChannelAccountID      respjson.Field
		ChannelID             respjson.Field
		Client                respjson.Field
		ConversationsThreadID respjson.Field
		CreatedAt             respjson.Field
		CreatedBy             respjson.Field
		Recipients            respjson.Field
		Senders               respjson.Field
		Text                  respjson.Field
		Type                  respjson.Field
		RichText              respjson.Field
		UpdatedAt             respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PublicWelcomeMessage) RawJSON() string { return r.JSON.raw }
func (r *PublicWelcomeMessage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PublicWelcomeMessageType string

const (
	PublicWelcomeMessageTypeWelcomeMessage PublicWelcomeMessageType = "WELCOME_MESSAGE"
)

type PublicWhatsAppTemplateMetadata struct {
	CrmObjectIDs     map[string]int64  `json:"crmObjectIds,required"`
	MappedTemplateID string            `json:"mappedTemplateId,required"`
	Parameters       map[string]string `json:"parameters,required"`
	// Any of "WHATSAPP_TEMPLATE_METADATA".
	Type PublicWhatsAppTemplateMetadataType `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CrmObjectIDs     respjson.Field
		MappedTemplateID respjson.Field
		Parameters       respjson.Field
		Type             respjson.Field
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

type QuickReply struct {
	Value     string `json:"value,required"`
	ValueType string `json:"valueType,required"`
	Label     string `json:"label"`
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

// The properties Value, ValueType are required.
type QuickReplyParam struct {
	Value     string            `json:"value,required"`
	ValueType string            `json:"valueType,required"`
	Label     param.Opt[string] `json:"label,omitzero"`
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
	MediaType      string `json:"mediaType,required"`
	ID             string `json:"id"`
	Description    string `json:"description"`
	MediaTitle     string `json:"mediaTitle"`
	MediaURL       string `json:"mediaUrl"`
	MediaURLString string `json:"mediaUrlString"`
	ThumbnailURL   string `json:"thumbnailUrl"`
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

// The property MediaType is required.
type SocialMetadataParam struct {
	MediaType      string            `json:"mediaType,required"`
	ID             param.Opt[string] `json:"id,omitzero"`
	Description    param.Opt[string] `json:"description,omitzero"`
	MediaTitle     param.Opt[string] `json:"mediaTitle,omitzero"`
	MediaURL       param.Opt[string] `json:"mediaUrl,omitzero"`
	MediaURLString param.Opt[string] `json:"mediaUrlString,omitzero"`
	ThumbnailURL   param.Opt[string] `json:"thumbnailUrl,omitzero"`
	paramObj
}

func (r SocialMetadataParam) MarshalJSON() (data []byte, err error) {
	type shadow SocialMetadataParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SocialMetadataParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SystemActor struct {
	ID string `json:"id,required"`
	// Any of "SYSTEM".
	Type SystemActorType `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SystemActor) RawJSON() string { return r.JSON.raw }
func (r *SystemActor) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SystemActorType string

const (
	SystemActorTypeSystem SystemActorType = "SYSTEM"
)

type VisitorActor struct {
	ID string `json:"id,required"`
	// Any of "VISITOR".
	Type   VisitorActorType `json:"type,required"`
	Avatar string           `json:"avatar"`
	Email  string           `json:"email"`
	Name   string           `json:"name"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Type        respjson.Field
		Avatar      respjson.Field
		Email       respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VisitorActor) RawJSON() string { return r.JSON.raw }
func (r *VisitorActor) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VisitorActorType string

const (
	VisitorActorTypeVisitor VisitorActorType = "VISITOR"
)
