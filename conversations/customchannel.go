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

	"github.com/stainless-sdks/hubspot-sdk-go/internal/apijson"
	"github.com/stainless-sdks/hubspot-sdk-go/internal/apiquery"
	shimjson "github.com/stainless-sdks/hubspot-sdk-go/internal/encoding/json"
	"github.com/stainless-sdks/hubspot-sdk-go/internal/requestconfig"
	"github.com/stainless-sdks/hubspot-sdk-go/option"
	"github.com/stainless-sdks/hubspot-sdk-go/packages/pagination"
	"github.com/stainless-sdks/hubspot-sdk-go/packages/param"
	"github.com/stainless-sdks/hubspot-sdk-go/packages/respjson"
	"github.com/stainless-sdks/hubspot-sdk-go/shared"
)

// CustomChannelService contains methods and other services that help with
// interacting with the hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCustomChannelService] method instead.
type CustomChannelService struct {
	Options                     []option.RequestOption
	ChannelAccountStagingTokens CustomChannelChannelAccountStagingTokenService
	ChannelAccounts             CustomChannelChannelAccountService
	Messages                    CustomChannelMessageService
}

// NewCustomChannelService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCustomChannelService(opts ...option.RequestOption) (r CustomChannelService) {
	r = CustomChannelService{}
	r.Options = opts
	r.ChannelAccountStagingTokens = NewCustomChannelChannelAccountStagingTokenService(opts...)
	r.ChannelAccounts = NewCustomChannelChannelAccountService(opts...)
	r.Messages = NewCustomChannelMessageService(opts...)
	return
}

// Register a new channel along with its capabilities and the webhook url that will
// be used to receive messages published over the channel
func (r *CustomChannelService) New(ctx context.Context, body CustomChannelNewParams, opts ...option.RequestOption) (res *PublicChannelIntegrationChannel, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "conversations/v3/custom-channels/"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Update the capabilities for an existing. You can also use it to update the
// channel's webhookUri and its channelAccountConnectionRedirectUrl.
func (r *CustomChannelService) Update(ctx context.Context, channelID int64, body CustomChannelUpdateParams, opts ...option.RequestOption) (res *PublicChannelIntegrationChannel, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("conversations/v3/custom-channels/%v", channelID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Retrieve all custom channels associated with the app.
func (r *CustomChannelService) List(ctx context.Context, query CustomChannelListParams, opts ...option.RequestOption) (res *pagination.Page[PublicChannelIntegrationChannel], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "conversations/v3/custom-channels/"
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

// Retrieve all custom channels associated with the app.
func (r *CustomChannelService) ListAutoPaging(ctx context.Context, query CustomChannelListParams, opts ...option.RequestOption) *pagination.PageAutoPager[PublicChannelIntegrationChannel] {
	return pagination.NewPageAutoPager(r.List(ctx, query, opts...))
}

// Archive an existing registered custom channel
func (r *CustomChannelService) Delete(ctx context.Context, channelID int64, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := fmt.Sprintf("conversations/v3/custom-channels/%v", channelID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Retrieve the details about a custom channel. This API allows you to see a custom
// channel's current capabilties and other configuration metadata
func (r *CustomChannelService) Get(ctx context.Context, channelID int64, opts ...option.RequestOption) (res *PublicChannelIntegrationChannel, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("conversations/v3/custom-channels/%v", channelID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// The properties Attachments, ChannelAccountID, MessageDirection, Recipients,
// Senders, Text, Timestamp are required.
type ChannelIntegrationMessageEggParam struct {
	Attachments      []ChannelIntegrationMessageEggAttachmentUnionParam `json:"attachments,omitzero,required"`
	ChannelAccountID string                                             `json:"channelAccountId,required"`
	// Any of "INCOMING", "OUTGOING".
	MessageDirection         ChannelIntegrationMessageEggMessageDirection `json:"messageDirection,omitzero,required"`
	Recipients               []ChannelIntegrationParticipantParam         `json:"recipients,omitzero,required"`
	Senders                  []ChannelIntegrationParticipantParam         `json:"senders,omitzero,required"`
	Text                     string                                       `json:"text,required"`
	Timestamp                time.Time                                    `json:"timestamp,required" format:"date-time"`
	InReplyToID              param.Opt[string]                            `json:"inReplyToId,omitzero"`
	IntegrationIdempotencyID param.Opt[string]                            `json:"integrationIdempotencyId,omitzero"`
	IntegrationThreadID      param.Opt[string]                            `json:"integrationThreadId,omitzero"`
	RichText                 param.Opt[string]                            `json:"richText,omitzero"`
	PreResolvedContacts      PreResolvedContactsParam                     `json:"preResolvedContacts,omitzero"`
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

func (u *ChannelIntegrationMessageEggAttachmentUnionParam) asAny() any {
	if !param.IsOmitted(u.OfFile) {
		return u.OfFile
	} else if !param.IsOmitted(u.OfLocation) {
		return u.OfLocation
	} else if !param.IsOmitted(u.OfContact) {
		return u.OfContact
	} else if !param.IsOmitted(u.OfUnsupportedContent) {
		return u.OfUnsupportedContent
	} else if !param.IsOmitted(u.OfMessageHeader) {
		return u.OfMessageHeader
	} else if !param.IsOmitted(u.OfQuickReplies) {
		return u.OfQuickReplies
	} else if !param.IsOmitted(u.OfSocialMediaMetadata) {
		return u.OfSocialMediaMetadata
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ChannelIntegrationMessageEggAttachmentUnionParam) GetFileUsageType() *string {
	if vt := u.OfFile; vt != nil && vt.FileUsageType.Valid() {
		return &vt.FileUsageType.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ChannelIntegrationMessageEggAttachmentUnionParam) GetLatitude() *float64 {
	if vt := u.OfLocation; vt != nil {
		return &vt.Latitude
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ChannelIntegrationMessageEggAttachmentUnionParam) GetLongitude() *float64 {
	if vt := u.OfLocation; vt != nil {
		return &vt.Longitude
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ChannelIntegrationMessageEggAttachmentUnionParam) GetAddress() *string {
	if vt := u.OfLocation; vt != nil && vt.Address.Valid() {
		return &vt.Address.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ChannelIntegrationMessageEggAttachmentUnionParam) GetName() *string {
	if vt := u.OfLocation; vt != nil && vt.Name.Valid() {
		return &vt.Name.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ChannelIntegrationMessageEggAttachmentUnionParam) GetURL() *string {
	if vt := u.OfLocation; vt != nil && vt.URL.Valid() {
		return &vt.URL.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ChannelIntegrationMessageEggAttachmentUnionParam) GetContactProfile() *ContactProfileParam {
	if vt := u.OfContact; vt != nil {
		return &vt.ContactProfile
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ChannelIntegrationMessageEggAttachmentUnionParam) GetText() *string {
	if vt := u.OfMessageHeader; vt != nil && vt.Text.Valid() {
		return &vt.Text.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ChannelIntegrationMessageEggAttachmentUnionParam) GetQuickReplies() []QuickReplyParam {
	if vt := u.OfQuickReplies; vt != nil {
		return vt.QuickReplies
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ChannelIntegrationMessageEggAttachmentUnionParam) GetSocialMetadata() *SocialMetadataParam {
	if vt := u.OfSocialMediaMetadata; vt != nil {
		return &vt.SocialMetadata
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ChannelIntegrationMessageEggAttachmentUnionParam) GetType() *string {
	if vt := u.OfFile; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfLocation; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfContact; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfUnsupportedContent; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfMessageHeader; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfQuickReplies; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfSocialMediaMetadata; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

// Returns a subunion which exports methods to access subproperties
//
// Or use AsAny() to get the underlying value
func (u ChannelIntegrationMessageEggAttachmentUnionParam) GetFileID() (res channelIntegrationMessageEggAttachmentUnionParamFileID) {
	if vt := u.OfFile; vt != nil {
		res.any = &vt.FileID
	} else if vt := u.OfMessageHeader; vt != nil && vt.FileID.Valid() {
		res.any = &vt.FileID.Value
	}
	return
}

// Can have the runtime types [*string], [*int64]
type channelIntegrationMessageEggAttachmentUnionParamFileID struct{ any }

// Use the following switch statement to get the type of the union:
//
//	switch u.AsAny().(type) {
//	case *string:
//	case *int64:
//	default:
//	    fmt.Errorf("not present")
//	}
func (u channelIntegrationMessageEggAttachmentUnionParamFileID) AsAny() any { return u.any }

type ChannelIntegrationMessageEggMessageDirection string

const (
	ChannelIntegrationMessageEggMessageDirectionIncoming ChannelIntegrationMessageEggMessageDirection = "INCOMING"
	ChannelIntegrationMessageEggMessageDirectionOutgoing ChannelIntegrationMessageEggMessageDirection = "OUTGOING"
)

// The property DeliveryIdentifier is required.
type ChannelIntegrationParticipantParam struct {
	DeliveryIdentifier PublicDeliveryIdentifierParam `json:"deliveryIdentifier,omitzero,required"`
	Name               param.Opt[string]             `json:"name,omitzero"`
	paramObj
}

func (r ChannelIntegrationParticipantParam) MarshalJSON() (data []byte, err error) {
	type shadow ChannelIntegrationParticipantParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ChannelIntegrationParticipantParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CollectionResponseWithTotalPublicChannelIntegrationChannelForwardPaging struct {
	Results []PublicChannelIntegrationChannel `json:"results,required"`
	Total   int64                             `json:"total,required"`
	Paging  shared.ForwardPaging              `json:"paging"`
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
func (r CollectionResponseWithTotalPublicChannelIntegrationChannelForwardPaging) RawJSON() string {
	return r.JSON.raw
}
func (r *CollectionResponseWithTotalPublicChannelIntegrationChannelForwardPaging) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties ContactProfile, Type are required.
type ContactAttachmentParam struct {
	ContactProfile ContactProfileParam `json:"contactProfile,omitzero,required"`
	// Any of "CONTACT".
	Type ContactAttachmentType `json:"type,omitzero,required"`
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

// The properties FileID, Type are required.
type FileAttachmentParam struct {
	FileID string `json:"fileId,required"`
	// Any of "FILE".
	Type          FileAttachmentType `json:"type,omitzero,required"`
	FileUsageType param.Opt[string]  `json:"fileUsageType,omitzero"`
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

// The properties Latitude, Longitude, Type are required.
type LocationAttachmentParam struct {
	Latitude  float64 `json:"latitude,required"`
	Longitude float64 `json:"longitude,required"`
	// Any of "LOCATION".
	Type    LocationAttachmentType `json:"type,omitzero,required"`
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
	Type   MessageHeaderAttachmentType `json:"type,omitzero,required"`
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

// The properties ContactPropertiesLeadingToMatch, ContactVid are required.
type PreResolvedContactParam struct {
	ContactPropertiesLeadingToMatch []string `json:"contactPropertiesLeadingToMatch,omitzero,required"`
	ContactVid                      int64    `json:"contactVid,required"`
	paramObj
}

func (r PreResolvedContactParam) MarshalJSON() (data []byte, err error) {
	type shadow PreResolvedContactParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PreResolvedContactParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Contacts is required.
type PreResolvedContactsParam struct {
	Contacts []PreResolvedContactParam `json:"contacts,omitzero,required"`
	paramObj
}

func (r PreResolvedContactsParam) MarshalJSON() (data []byte, err error) {
	type shadow PreResolvedContactsParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PreResolvedContactsParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Authorized, InboxID, Name are required.
type PublicChannelAccountEggParam struct {
	Authorized         bool                          `json:"authorized,required"`
	InboxID            string                        `json:"inboxId,required"`
	Name               string                        `json:"name,required"`
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
	AccountToken       string                   `json:"accountToken,required"`
	CreatedAt          time.Time                `json:"createdAt,required" format:"date-time"`
	GenericChannelID   int64                    `json:"genericChannelId,required"`
	InboxID            int64                    `json:"inboxId,required"`
	UserID             int64                    `json:"userId,required"`
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

// The properties AccountName, DeliveryIdentifier are required.
type PublicChannelAccountStagingTokenUpdateRequestParam struct {
	AccountName        string                        `json:"accountName,required"`
	DeliveryIdentifier PublicDeliveryIdentifierParam `json:"deliveryIdentifier,omitzero,required"`
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
	ID                                  string         `json:"id,required"`
	Capabilities                        map[string]any `json:"capabilities,required"`
	CreatedAt                           time.Time      `json:"createdAt,required" format:"date-time"`
	Name                                string         `json:"name,required"`
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
	Capabilities                        map[string]any    `json:"capabilities,omitzero,required"`
	Name                                string            `json:"name,required"`
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
	Capabilities                        map[string]any `json:"capabilities,omitzero,required"`
	ChannelAccountConnectionRedirectURL any            `json:"channelAccountConnectionRedirectUrl,omitzero,required"`
	ChannelDescription                  any            `json:"channelDescription,omitzero,required"`
	ChannelLogoURL                      any            `json:"channelLogoUrl,omitzero,required"`
	Name                                any            `json:"name,omitzero,required"`
	WebhookURL                          any            `json:"webhookUrl,omitzero,required"`
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
	StatusType   PublicChannelIntegrationMessageUpdateRequestStatusType `json:"statusType,omitzero,required"`
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

// The properties QuickReplies, Type are required.
type QuickRepliesAttachmentParam struct {
	QuickReplies []QuickReplyParam `json:"quickReplies,omitzero,required"`
	// Any of "QUICK_REPLIES".
	Type QuickRepliesAttachmentType `json:"type,omitzero,required"`
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

// The properties SocialMetadata, Type are required.
type SocialMetadataIntegrationAttachmentParam struct {
	SocialMetadata SocialMetadataParam `json:"socialMetadata,omitzero,required"`
	// Any of "SOCIAL_MEDIA_METADATA".
	Type SocialMetadataIntegrationAttachmentType `json:"type,omitzero,required"`
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
	Type UnsupportedContentAttachmentType `json:"type,omitzero,required"`
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
	return json.Unmarshal(data, &r.PublicChannelIntegrationChannelCreate)
}

type CustomChannelUpdateParams struct {
	PublicChannelIntegrationChannelPatch PublicChannelIntegrationChannelPatchParam
	paramObj
}

func (r CustomChannelUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.PublicChannelIntegrationChannelPatch)
}
func (r *CustomChannelUpdateParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.PublicChannelIntegrationChannelPatch)
}

type CustomChannelListParams struct {
	// The paging cursor token of the last successfully read resource will be returned
	// as the `paging.next.after` JSON property of a paged response containing more
	// results.
	After param.Opt[string] `query:"after,omitzero" json:"-"`
	// Specify the default number of results to return per page.
	DefaultPageLength param.Opt[int64] `query:"defaultPageLength,omitzero" json:"-"`
	// The maximum number of results to display per page.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Specify the sorting order for the results.
	Sort []string `query:"sort,omitzero" json:"-"`
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
