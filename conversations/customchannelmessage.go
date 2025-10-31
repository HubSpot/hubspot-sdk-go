// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package conversations

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"
	"time"

	"github.com/stainless-sdks/hubspot-sdk-go/internal/apijson"
	"github.com/stainless-sdks/hubspot-sdk-go/internal/requestconfig"
	"github.com/stainless-sdks/hubspot-sdk-go/option"
	"github.com/stainless-sdks/hubspot-sdk-go/packages/param"
)

// CustomChannelMessageService contains methods and other services that help with
// interacting with the Hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCustomChannelMessageService] method instead.
type CustomChannelMessageService struct {
	Options []option.RequestOption
}

// NewCustomChannelMessageService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewCustomChannelMessageService(opts ...option.RequestOption) (r CustomChannelMessageService) {
	r = CustomChannelMessageService{}
	r.Options = opts
	return
}

// Publish a message over your custom channel
func (r *CustomChannelMessageService) New(ctx context.Context, channelID string, body CustomChannelMessageNewParams, opts ...option.RequestOption) (res *PublicConversationsMessage, err error) {
	opts = slices.Concat(r.Options, opts)
	if channelID == "" {
		err = errors.New("missing required channelId parameter")
		return
	}
	path := fmt.Sprintf("conversations/v3/custom-channels/%s/messages", channelID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Update a message's status to indicate if it was successfully sent, failed to
// send, or was read. For failed messages, this can also include the error message
// for the failure.
func (r *CustomChannelMessageService) Update(ctx context.Context, messageID string, params CustomChannelMessageUpdateParams, opts ...option.RequestOption) (res *PublicConversationsMessage, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.ChannelID == "" {
		err = errors.New("missing required channelId parameter")
		return
	}
	if messageID == "" {
		err = errors.New("missing required messageId parameter")
		return
	}
	path := fmt.Sprintf("conversations/v3/custom-channels/%s/messages/%s", params.ChannelID, messageID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return
}

// Get the details for a specific message sent over a custom channel
func (r *CustomChannelMessageService) Get(ctx context.Context, messageID string, query CustomChannelMessageGetParams, opts ...option.RequestOption) (res *PublicConversationsMessage, err error) {
	opts = slices.Concat(r.Options, opts)
	if query.ChannelID == "" {
		err = errors.New("missing required channelId parameter")
		return
	}
	if messageID == "" {
		err = errors.New("missing required messageId parameter")
		return
	}
	path := fmt.Sprintf("conversations/v3/custom-channels/%s/messages/%s", query.ChannelID, messageID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type CustomChannelMessageNewParams struct {
	Attachments         []CustomChannelMessageNewParamsAttachmentUnion `json:"attachments,omitzero,required"`
	ChannelAccountID    string                                         `json:"channelAccountId,required"`
	IntegrationThreadID string                                         `json:"integrationThreadId,required"`
	// Any of "INCOMING", "OUTGOING".
	MessageDirection         CustomChannelMessageNewParamsMessageDirection    `json:"messageDirection,omitzero,required"`
	Recipients               []CustomChannelMessageNewParamsRecipient         `json:"recipients,omitzero,required"`
	Senders                  []CustomChannelMessageNewParamsSender            `json:"senders,omitzero,required"`
	Text                     string                                           `json:"text,required"`
	Timestamp                time.Time                                        `json:"timestamp,required" format:"date-time"`
	InReplyToID              param.Opt[string]                                `json:"inReplyToId,omitzero"`
	IntegrationIdempotencyID param.Opt[string]                                `json:"integrationIdempotencyId,omitzero"`
	RichText                 param.Opt[string]                                `json:"richText,omitzero"`
	PreResolvedContacts      CustomChannelMessageNewParamsPreResolvedContacts `json:"preResolvedContacts,omitzero"`
	paramObj
}

func (r CustomChannelMessageNewParams) MarshalJSON() (data []byte, err error) {
	type shadow CustomChannelMessageNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CustomChannelMessageNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type CustomChannelMessageNewParamsAttachmentUnion struct {
	OfFile                *CustomChannelMessageNewParamsAttachmentFile                `json:",omitzero,inline"`
	OfLocation            *CustomChannelMessageNewParamsAttachmentLocation            `json:",omitzero,inline"`
	OfContact             *CustomChannelMessageNewParamsAttachmentContact             `json:",omitzero,inline"`
	OfUnsupportedContent  *CustomChannelMessageNewParamsAttachmentUnsupportedContent  `json:",omitzero,inline"`
	OfMessageHeader       *CustomChannelMessageNewParamsAttachmentMessageHeader       `json:",omitzero,inline"`
	OfQuickReplies        *CustomChannelMessageNewParamsAttachmentQuickReplies        `json:",omitzero,inline"`
	OfSocialMediaMetadata *CustomChannelMessageNewParamsAttachmentSocialMediaMetadata `json:",omitzero,inline"`
	paramUnion
}

func (u CustomChannelMessageNewParamsAttachmentUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfFile,
		u.OfLocation,
		u.OfContact,
		u.OfUnsupportedContent,
		u.OfMessageHeader,
		u.OfQuickReplies,
		u.OfSocialMediaMetadata)
}
func (u *CustomChannelMessageNewParamsAttachmentUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *CustomChannelMessageNewParamsAttachmentUnion) asAny() any {
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
func (u CustomChannelMessageNewParamsAttachmentUnion) GetFileUsageType() *string {
	if vt := u.OfFile; vt != nil && vt.FileUsageType.Valid() {
		return &vt.FileUsageType.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u CustomChannelMessageNewParamsAttachmentUnion) GetLatitude() *float64 {
	if vt := u.OfLocation; vt != nil {
		return &vt.Latitude
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u CustomChannelMessageNewParamsAttachmentUnion) GetLongitude() *float64 {
	if vt := u.OfLocation; vt != nil {
		return &vt.Longitude
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u CustomChannelMessageNewParamsAttachmentUnion) GetAddress() *string {
	if vt := u.OfLocation; vt != nil && vt.Address.Valid() {
		return &vt.Address.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u CustomChannelMessageNewParamsAttachmentUnion) GetName() *string {
	if vt := u.OfLocation; vt != nil && vt.Name.Valid() {
		return &vt.Name.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u CustomChannelMessageNewParamsAttachmentUnion) GetURL() *string {
	if vt := u.OfLocation; vt != nil && vt.URL.Valid() {
		return &vt.URL.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u CustomChannelMessageNewParamsAttachmentUnion) GetContactProfile() *ContactProfileParam {
	if vt := u.OfContact; vt != nil {
		return &vt.ContactProfile
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u CustomChannelMessageNewParamsAttachmentUnion) GetText() *string {
	if vt := u.OfMessageHeader; vt != nil && vt.Text.Valid() {
		return &vt.Text.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u CustomChannelMessageNewParamsAttachmentUnion) GetQuickReplies() []QuickReplyParam {
	if vt := u.OfQuickReplies; vt != nil {
		return vt.QuickReplies
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u CustomChannelMessageNewParamsAttachmentUnion) GetSocialMetadata() *SocialMetadataParam {
	if vt := u.OfSocialMediaMetadata; vt != nil {
		return &vt.SocialMetadata
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u CustomChannelMessageNewParamsAttachmentUnion) GetType() *string {
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
func (u CustomChannelMessageNewParamsAttachmentUnion) GetFileID() (res customChannelMessageNewParamsAttachmentUnionFileID) {
	if vt := u.OfFile; vt != nil {
		res.any = &vt.FileID
	} else if vt := u.OfMessageHeader; vt != nil && vt.FileID.Valid() {
		res.any = &vt.FileID.Value
	}
	return
}

// Can have the runtime types [*string], [*int64]
type customChannelMessageNewParamsAttachmentUnionFileID struct{ any }

// Use the following switch statement to get the type of the union:
//
//	switch u.AsAny().(type) {
//	case *string:
//	case *int64:
//	default:
//	    fmt.Errorf("not present")
//	}
func (u customChannelMessageNewParamsAttachmentUnionFileID) AsAny() any { return u.any }

// The properties FileID, Type are required.
type CustomChannelMessageNewParamsAttachmentFile struct {
	FileID string `json:"fileId,required"`
	// Any of "FILE".
	Type          string            `json:"type,omitzero,required"`
	FileUsageType param.Opt[string] `json:"fileUsageType,omitzero"`
	paramObj
}

func (r CustomChannelMessageNewParamsAttachmentFile) MarshalJSON() (data []byte, err error) {
	type shadow CustomChannelMessageNewParamsAttachmentFile
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CustomChannelMessageNewParamsAttachmentFile) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CustomChannelMessageNewParamsAttachmentFile](
		"type", "FILE",
	)
}

// The properties Latitude, Longitude, Type are required.
type CustomChannelMessageNewParamsAttachmentLocation struct {
	Latitude  float64 `json:"latitude,required"`
	Longitude float64 `json:"longitude,required"`
	// Any of "LOCATION".
	Type    string            `json:"type,omitzero,required"`
	Address param.Opt[string] `json:"address,omitzero"`
	Name    param.Opt[string] `json:"name,omitzero"`
	URL     param.Opt[string] `json:"url,omitzero"`
	paramObj
}

func (r CustomChannelMessageNewParamsAttachmentLocation) MarshalJSON() (data []byte, err error) {
	type shadow CustomChannelMessageNewParamsAttachmentLocation
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CustomChannelMessageNewParamsAttachmentLocation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CustomChannelMessageNewParamsAttachmentLocation](
		"type", "LOCATION",
	)
}

// The properties ContactProfile, Type are required.
type CustomChannelMessageNewParamsAttachmentContact struct {
	ContactProfile ContactProfileParam `json:"contactProfile,omitzero,required"`
	// Any of "CONTACT".
	Type string `json:"type,omitzero,required"`
	paramObj
}

func (r CustomChannelMessageNewParamsAttachmentContact) MarshalJSON() (data []byte, err error) {
	type shadow CustomChannelMessageNewParamsAttachmentContact
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CustomChannelMessageNewParamsAttachmentContact) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CustomChannelMessageNewParamsAttachmentContact](
		"type", "CONTACT",
	)
}

// The property Type is required.
type CustomChannelMessageNewParamsAttachmentUnsupportedContent struct {
	// Any of "UNSUPPORTED_CONTENT".
	Type string `json:"type,omitzero,required"`
	paramObj
}

func (r CustomChannelMessageNewParamsAttachmentUnsupportedContent) MarshalJSON() (data []byte, err error) {
	type shadow CustomChannelMessageNewParamsAttachmentUnsupportedContent
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CustomChannelMessageNewParamsAttachmentUnsupportedContent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CustomChannelMessageNewParamsAttachmentUnsupportedContent](
		"type", "UNSUPPORTED_CONTENT",
	)
}

// The property Type is required.
type CustomChannelMessageNewParamsAttachmentMessageHeader struct {
	// Any of "MESSAGE_HEADER".
	Type   string            `json:"type,omitzero,required"`
	FileID param.Opt[int64]  `json:"fileId,omitzero"`
	Text   param.Opt[string] `json:"text,omitzero"`
	paramObj
}

func (r CustomChannelMessageNewParamsAttachmentMessageHeader) MarshalJSON() (data []byte, err error) {
	type shadow CustomChannelMessageNewParamsAttachmentMessageHeader
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CustomChannelMessageNewParamsAttachmentMessageHeader) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CustomChannelMessageNewParamsAttachmentMessageHeader](
		"type", "MESSAGE_HEADER",
	)
}

// The properties QuickReplies, Type are required.
type CustomChannelMessageNewParamsAttachmentQuickReplies struct {
	QuickReplies []QuickReplyParam `json:"quickReplies,omitzero,required"`
	// Any of "QUICK_REPLIES".
	Type string `json:"type,omitzero,required"`
	paramObj
}

func (r CustomChannelMessageNewParamsAttachmentQuickReplies) MarshalJSON() (data []byte, err error) {
	type shadow CustomChannelMessageNewParamsAttachmentQuickReplies
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CustomChannelMessageNewParamsAttachmentQuickReplies) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CustomChannelMessageNewParamsAttachmentQuickReplies](
		"type", "QUICK_REPLIES",
	)
}

// The properties SocialMetadata, Type are required.
type CustomChannelMessageNewParamsAttachmentSocialMediaMetadata struct {
	SocialMetadata SocialMetadataParam `json:"socialMetadata,omitzero,required"`
	// Any of "SOCIAL_MEDIA_METADATA".
	Type string `json:"type,omitzero,required"`
	paramObj
}

func (r CustomChannelMessageNewParamsAttachmentSocialMediaMetadata) MarshalJSON() (data []byte, err error) {
	type shadow CustomChannelMessageNewParamsAttachmentSocialMediaMetadata
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CustomChannelMessageNewParamsAttachmentSocialMediaMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CustomChannelMessageNewParamsAttachmentSocialMediaMetadata](
		"type", "SOCIAL_MEDIA_METADATA",
	)
}

type CustomChannelMessageNewParamsMessageDirection string

const (
	CustomChannelMessageNewParamsMessageDirectionIncoming CustomChannelMessageNewParamsMessageDirection = "INCOMING"
	CustomChannelMessageNewParamsMessageDirectionOutgoing CustomChannelMessageNewParamsMessageDirection = "OUTGOING"
)

// The property DeliveryIdentifier is required.
type CustomChannelMessageNewParamsRecipient struct {
	DeliveryIdentifier PublicDeliveryIdentifierParam `json:"deliveryIdentifier,omitzero,required"`
	Name               param.Opt[string]             `json:"name,omitzero"`
	paramObj
}

func (r CustomChannelMessageNewParamsRecipient) MarshalJSON() (data []byte, err error) {
	type shadow CustomChannelMessageNewParamsRecipient
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CustomChannelMessageNewParamsRecipient) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property DeliveryIdentifier is required.
type CustomChannelMessageNewParamsSender struct {
	DeliveryIdentifier PublicDeliveryIdentifierParam `json:"deliveryIdentifier,omitzero,required"`
	Name               param.Opt[string]             `json:"name,omitzero"`
	paramObj
}

func (r CustomChannelMessageNewParamsSender) MarshalJSON() (data []byte, err error) {
	type shadow CustomChannelMessageNewParamsSender
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CustomChannelMessageNewParamsSender) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Contacts is required.
type CustomChannelMessageNewParamsPreResolvedContacts struct {
	Contacts []CustomChannelMessageNewParamsPreResolvedContactsContact `json:"contacts,omitzero,required"`
	paramObj
}

func (r CustomChannelMessageNewParamsPreResolvedContacts) MarshalJSON() (data []byte, err error) {
	type shadow CustomChannelMessageNewParamsPreResolvedContacts
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CustomChannelMessageNewParamsPreResolvedContacts) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties ContactPropertiesLeadingToMatch, ContactVid are required.
type CustomChannelMessageNewParamsPreResolvedContactsContact struct {
	ContactPropertiesLeadingToMatch []string `json:"contactPropertiesLeadingToMatch,omitzero,required"`
	ContactVid                      int64    `json:"contactVid,required"`
	paramObj
}

func (r CustomChannelMessageNewParamsPreResolvedContactsContact) MarshalJSON() (data []byte, err error) {
	type shadow CustomChannelMessageNewParamsPreResolvedContactsContact
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CustomChannelMessageNewParamsPreResolvedContactsContact) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CustomChannelMessageUpdateParams struct {
	ChannelID string `path:"channelId,required" json:"-"`
	// Valid status are SENT, FAILED, and READ
	//
	// Any of "SENT", "FAILED", "READ".
	StatusType   CustomChannelMessageUpdateParamsStatusType `json:"statusType,omitzero,required"`
	ErrorMessage param.Opt[string]                          `json:"errorMessage,omitzero"`
	paramObj
}

func (r CustomChannelMessageUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow CustomChannelMessageUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CustomChannelMessageUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Valid status are SENT, FAILED, and READ
type CustomChannelMessageUpdateParamsStatusType string

const (
	CustomChannelMessageUpdateParamsStatusTypeSent   CustomChannelMessageUpdateParamsStatusType = "SENT"
	CustomChannelMessageUpdateParamsStatusTypeFailed CustomChannelMessageUpdateParamsStatusType = "FAILED"
	CustomChannelMessageUpdateParamsStatusTypeRead   CustomChannelMessageUpdateParamsStatusType = "READ"
)

type CustomChannelMessageGetParams struct {
	ChannelID string `path:"channelId,required" json:"-"`
	paramObj
}
