// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package crm

import (
	"encoding/json"
	"time"

	"github.com/stainless-sdks/hubspot-sdk-go/internal/apijson"
	"github.com/stainless-sdks/hubspot-sdk-go/option"
	"github.com/stainless-sdks/hubspot-sdk-go/packages/param"
	"github.com/stainless-sdks/hubspot-sdk-go/packages/respjson"
)

// TimelineService contains methods and other services that help with interacting
// with the hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTimelineService] method instead.
type TimelineService struct {
	Options   []option.RequestOption
	Events    TimelineEventService
	Templates TimelineTemplateService
	Tokens    TimelineTokenService
}

// NewTimelineService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewTimelineService(opts ...option.RequestOption) (r TimelineService) {
	r = TimelineService{}
	r.Options = opts
	r.Events = NewTimelineEventService(opts...)
	r.Templates = NewTimelineTemplateService(opts...)
	r.Tokens = NewTimelineTokenService(opts...)
	return
}

// Used to create timeline events in batches.
//
// The property Inputs is required.
type BatchInputTimelineEventParam struct {
	// A collection of timeline events we want to create.
	Inputs []TimelineEventParam `json:"inputs,omitzero,required"`
	paramObj
}

func (r BatchInputTimelineEventParam) MarshalJSON() (data []byte, err error) {
	type shadow BatchInputTimelineEventParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BatchInputTimelineEventParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CollectionResponseTimelineEventTemplateNoPaging struct {
	Results []TimelineEventTemplate `json:"results,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Results     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CollectionResponseTimelineEventTemplateNoPaging) RawJSON() string { return r.JSON.raw }
func (r *CollectionResponseTimelineEventTemplateNoPaging) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The details Markdown rendered as HTML.
type EventDetail struct {
	// The details Markdown rendered as HTML.
	Details string `json:"details,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Details     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EventDetail) RawJSON() string { return r.JSON.raw }
func (r *EventDetail) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The state of the timeline event.
//
// The properties EventTemplateID, Tokens are required.
type TimelineEventParam struct {
	// The event template ID.
	EventTemplateID string `json:"eventTemplateId,required"`
	// A collection of token keys and values associated with the template tokens.
	Tokens map[string]string `json:"tokens,omitzero,required"`
	// Identifier for the event. This is optional, and we recommend you do not pass
	// this in. We will create one for you if you omit this. You can also use
	// `{{uuid}}` anywhere in the ID to generate a unique string, guaranteeing
	// uniqueness.
	ID param.Opt[string] `json:"id,omitzero"`
	// The event domain (often paired with utk).
	Domain param.Opt[string] `json:"domain,omitzero"`
	// The email address used for contact-specific events. This can be used to identify
	// existing contacts, create new ones, or change the email for an existing contact
	// (if paired with the `objectId`).
	Email param.Opt[string] `json:"email,omitzero"`
	// The CRM object identifier. This is required for every event other than contacts
	// (where utk or email can be used).
	ObjectID param.Opt[string] `json:"objectId,omitzero"`
	// The time the event occurred. If not passed in, the curren time will be assumed.
	// This is used to determine where an event is shown on a CRM object's timeline.
	Timestamp param.Opt[time.Time] `json:"timestamp,omitzero" format:"date-time"`
	// Use the `utk` parameter to associate an event with a contact by `usertoken`.
	// This is recommended if you don't know a user's email, but have an identifying
	// user token in your cookie.
	Utk param.Opt[string] `json:"utk,omitzero"`
	// Additional event-specific data that can be interpreted by the template's
	// markdown.
	ExtraData      any                      `json:"extraData,omitzero"`
	TimelineIFrame TimelineEventIFrameParam `json:"timelineIFrame,omitzero"`
	paramObj
}

func (r TimelineEventParam) MarshalJSON() (data []byte, err error) {
	type shadow TimelineEventParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TimelineEventParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TimelineEventIFrame struct {
	// The label of the modal window that displays the iframe contents.
	HeaderLabel string `json:"headerLabel,required"`
	// The height of the modal window in pixels.
	Height int64 `json:"height,required"`
	// The text displaying the link that will display the iframe.
	LinkLabel string `json:"linkLabel,required"`
	// The URI of the iframe contents.
	URL string `json:"url,required"`
	// The width of the modal window in pixels.
	Width int64 `json:"width,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		HeaderLabel respjson.Field
		Height      respjson.Field
		LinkLabel   respjson.Field
		URL         respjson.Field
		Width       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TimelineEventIFrame) RawJSON() string { return r.JSON.raw }
func (r *TimelineEventIFrame) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this TimelineEventIFrame to a TimelineEventIFrameParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// TimelineEventIFrameParam.Overrides()
func (r TimelineEventIFrame) ToParam() TimelineEventIFrameParam {
	return param.Override[TimelineEventIFrameParam](json.RawMessage(r.RawJSON()))
}

// The properties HeaderLabel, Height, LinkLabel, URL, Width are required.
type TimelineEventIFrameParam struct {
	// The label of the modal window that displays the iframe contents.
	HeaderLabel string `json:"headerLabel,required"`
	// The height of the modal window in pixels.
	Height int64 `json:"height,required"`
	// The text displaying the link that will display the iframe.
	LinkLabel string `json:"linkLabel,required"`
	// The URI of the iframe contents.
	URL string `json:"url,required"`
	// The width of the modal window in pixels.
	Width int64 `json:"width,required"`
	paramObj
}

func (r TimelineEventIFrameParam) MarshalJSON() (data []byte, err error) {
	type shadow TimelineEventIFrameParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TimelineEventIFrameParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The current state of the timeline event.
type TimelineEventResponse struct {
	// Identifier for the event. This should be unique to the app and event template.
	// If you use the same ID for different CRM objects, the last to be processed will
	// win and the first will not have a record. You can also use `{{uuid}}` anywhere
	// in the ID to generate a unique string, guaranteeing uniqueness.
	ID string `json:"id,required"`
	// The event template ID.
	EventTemplateID string `json:"eventTemplateId,required"`
	// The ObjectType associated with the EventTemplate.
	ObjectType string `json:"objectType,required"`
	// A collection of token keys and values associated with the template tokens.
	Tokens    map[string]string `json:"tokens,required"`
	CreatedAt time.Time         `json:"createdAt" format:"date-time"`
	// The event domain (often paired with utk).
	Domain string `json:"domain"`
	// The email address used for contact-specific events. This can be used to identify
	// existing contacts, create new ones, or change the email for an existing contact
	// (if paired with the `objectId`).
	Email string `json:"email"`
	// Additional event-specific data that can be interpreted by the template's
	// markdown.
	ExtraData any `json:"extraData"`
	// The CRM object identifier. This is required for every event other than contacts
	// (where utk or email can be used).
	ObjectID       string              `json:"objectId"`
	TimelineIFrame TimelineEventIFrame `json:"timelineIFrame"`
	// The time the event occurred. If not passed in, the curren time will be assumed.
	// This is used to determine where an event is shown on a CRM object's timeline.
	Timestamp time.Time `json:"timestamp" format:"date-time"`
	// Use the `utk` parameter to associate an event with a contact by `usertoken`.
	// This is recommended if you don't know a user's email, but have an identifying
	// user token in your cookie.
	Utk string `json:"utk"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		EventTemplateID respjson.Field
		ObjectType      respjson.Field
		Tokens          respjson.Field
		CreatedAt       respjson.Field
		Domain          respjson.Field
		Email           respjson.Field
		ExtraData       respjson.Field
		ObjectID        respjson.Field
		TimelineIFrame  respjson.Field
		Timestamp       respjson.Field
		Utk             respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TimelineEventResponse) RawJSON() string { return r.JSON.raw }
func (r *TimelineEventResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The current state of the template definition.
type TimelineEventTemplate struct {
	// The template ID.
	ID string `json:"id,required"`
	// The template name.
	Name string `json:"name,required"`
	// The type of CRM object this template is for. [Contacts, companies, tickets, and
	// deals] are supported.
	ObjectType string `json:"objectType,required"`
	// A collection of tokens that can be used as custom properties on the event and to
	// create fully fledged CRM objects.
	Tokens []TimelineEventTemplateToken `json:"tokens,required"`
	// The date and time that the Event Template was created, as an ISO 8601 timestamp.
	// Will be null if the template was created before Feb 18th, 2020.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// This uses Markdown syntax with Handlebars and event-specific data to render HTML
	// on a timeline when you expand the details.
	DetailTemplate string `json:"detailTemplate"`
	// This uses Markdown syntax with Handlebars and event-specific data to render HTML
	// on a timeline as a header.
	HeaderTemplate string `json:"headerTemplate"`
	// The date and time that the Event Template was last updated, as an ISO 8601
	// timestamp. Will be null if the template was created before Feb 18th, 2020.
	UpdatedAt time.Time `json:"updatedAt" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID             respjson.Field
		Name           respjson.Field
		ObjectType     respjson.Field
		Tokens         respjson.Field
		CreatedAt      respjson.Field
		DetailTemplate respjson.Field
		HeaderTemplate respjson.Field
		UpdatedAt      respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TimelineEventTemplate) RawJSON() string { return r.JSON.raw }
func (r *TimelineEventTemplate) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// State of the template definition being created.
//
// The properties Name, ObjectType, Tokens are required.
type TimelineEventTemplateCreateRequestParam struct {
	// The template name.
	Name string `json:"name,required"`
	// The type of CRM object this template is for. [Contacts, companies, tickets, and
	// deals] are supported.
	ObjectType string `json:"objectType,required"`
	// A collection of tokens that can be used as custom properties on the event and to
	// create fully fledged CRM objects.
	Tokens []TimelineEventTemplateTokenParam `json:"tokens,omitzero,required"`
	// This uses Markdown syntax with Handlebars and event-specific data to render HTML
	// on a timeline when you expand the details.
	DetailTemplate param.Opt[string] `json:"detailTemplate,omitzero"`
	// This uses Markdown syntax with Handlebars and event-specific data to render HTML
	// on a timeline as a header.
	HeaderTemplate param.Opt[string] `json:"headerTemplate,omitzero"`
	paramObj
}

func (r TimelineEventTemplateCreateRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow TimelineEventTemplateCreateRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TimelineEventTemplateCreateRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// State of the token definition.
type TimelineEventTemplateToken struct {
	// Used for list segmentation and reporting.
	Label string `json:"label,required"`
	// The name of the token referenced in the templates. This must be unique for the
	// specific template. It may only contain alphanumeric characters, periods, dashes,
	// or underscores (. - \_).
	Name string `json:"name,required"`
	// The data type of the token. You can currently choose from [string, number, date,
	// enumeration].
	//
	// Any of "date", "enumeration", "number", "string".
	Type TimelineEventTemplateTokenType `json:"type,required"`
	// The date and time that the Event Template Token was created, as an ISO 8601
	// timestamp. Will be null if the template was created before Feb 18th, 2020.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// The name of the CRM object property. This will populate the CRM object property
	// associated with the event. With enough of these, you can fully build CRM objects
	// via the Timeline API.
	ObjectPropertyName string `json:"objectPropertyName"`
	// If type is `enumeration`, we should have a list of options to choose from.
	Options []TimelineEventTemplateTokenOption `json:"options"`
	// The date and time that the Event Template Token was last updated, as an ISO 8601
	// timestamp. Will be null if the template was created before Feb 18th, 2020.
	UpdatedAt time.Time `json:"updatedAt" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Label              respjson.Field
		Name               respjson.Field
		Type               respjson.Field
		CreatedAt          respjson.Field
		ObjectPropertyName respjson.Field
		Options            respjson.Field
		UpdatedAt          respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TimelineEventTemplateToken) RawJSON() string { return r.JSON.raw }
func (r *TimelineEventTemplateToken) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this TimelineEventTemplateToken to a
// TimelineEventTemplateTokenParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// TimelineEventTemplateTokenParam.Overrides()
func (r TimelineEventTemplateToken) ToParam() TimelineEventTemplateTokenParam {
	return param.Override[TimelineEventTemplateTokenParam](json.RawMessage(r.RawJSON()))
}

// The data type of the token. You can currently choose from [string, number, date,
// enumeration].
type TimelineEventTemplateTokenType string

const (
	TimelineEventTemplateTokenTypeDate        TimelineEventTemplateTokenType = "date"
	TimelineEventTemplateTokenTypeEnumeration TimelineEventTemplateTokenType = "enumeration"
	TimelineEventTemplateTokenTypeNumber      TimelineEventTemplateTokenType = "number"
	TimelineEventTemplateTokenTypeString      TimelineEventTemplateTokenType = "string"
)

// State of the token definition.
//
// The properties Label, Name, Type are required.
type TimelineEventTemplateTokenParam struct {
	// Used for list segmentation and reporting.
	Label string `json:"label,required"`
	// The name of the token referenced in the templates. This must be unique for the
	// specific template. It may only contain alphanumeric characters, periods, dashes,
	// or underscores (. - \_).
	Name string `json:"name,required"`
	// The data type of the token. You can currently choose from [string, number, date,
	// enumeration].
	//
	// Any of "date", "enumeration", "number", "string".
	Type TimelineEventTemplateTokenType `json:"type,omitzero,required"`
	// The date and time that the Event Template Token was created, as an ISO 8601
	// timestamp. Will be null if the template was created before Feb 18th, 2020.
	CreatedAt param.Opt[time.Time] `json:"createdAt,omitzero" format:"date-time"`
	// The name of the CRM object property. This will populate the CRM object property
	// associated with the event. With enough of these, you can fully build CRM objects
	// via the Timeline API.
	ObjectPropertyName param.Opt[string] `json:"objectPropertyName,omitzero"`
	// The date and time that the Event Template Token was last updated, as an ISO 8601
	// timestamp. Will be null if the template was created before Feb 18th, 2020.
	UpdatedAt param.Opt[time.Time] `json:"updatedAt,omitzero" format:"date-time"`
	// If type is `enumeration`, we should have a list of options to choose from.
	Options []TimelineEventTemplateTokenOptionParam `json:"options,omitzero"`
	paramObj
}

func (r TimelineEventTemplateTokenParam) MarshalJSON() (data []byte, err error) {
	type shadow TimelineEventTemplateTokenParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TimelineEventTemplateTokenParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TimelineEventTemplateTokenOption struct {
	Label string `json:"label,required"`
	Value string `json:"value,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Label       respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TimelineEventTemplateTokenOption) RawJSON() string { return r.JSON.raw }
func (r *TimelineEventTemplateTokenOption) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this TimelineEventTemplateTokenOption to a
// TimelineEventTemplateTokenOptionParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// TimelineEventTemplateTokenOptionParam.Overrides()
func (r TimelineEventTemplateTokenOption) ToParam() TimelineEventTemplateTokenOptionParam {
	return param.Override[TimelineEventTemplateTokenOptionParam](json.RawMessage(r.RawJSON()))
}

// The properties Label, Value are required.
type TimelineEventTemplateTokenOptionParam struct {
	Label string `json:"label,required"`
	Value string `json:"value,required"`
	paramObj
}

func (r TimelineEventTemplateTokenOptionParam) MarshalJSON() (data []byte, err error) {
	type shadow TimelineEventTemplateTokenOptionParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TimelineEventTemplateTokenOptionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// State of the token definition for update requests.
//
// The property Label is required.
type TimelineEventTemplateTokenUpdateRequestParam struct {
	// Used for list segmentation and reporting.
	Label string `json:"label,required"`
	// The name of the CRM object property. This will populate the CRM object property
	// associated with the event. With enough of these, you can fully build CRM objects
	// via the Timeline API.
	ObjectPropertyName param.Opt[string] `json:"objectPropertyName,omitzero"`
	// If type is `enumeration`, we should have a list of options to choose from.
	Options []TimelineEventTemplateTokenOptionParam `json:"options,omitzero"`
	paramObj
}

func (r TimelineEventTemplateTokenUpdateRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow TimelineEventTemplateTokenUpdateRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TimelineEventTemplateTokenUpdateRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// State of the template definition being updated.
//
// The properties ID, Name, Tokens are required.
type TimelineEventTemplateUpdateRequestParam struct {
	// The template ID.
	ID string `json:"id,required"`
	// The template name.
	Name string `json:"name,required"`
	// A collection of tokens that can be used as custom properties on the event and to
	// create fully fledged CRM objects.
	Tokens []TimelineEventTemplateTokenParam `json:"tokens,omitzero,required"`
	// This uses Markdown syntax with Handlebars and event-specific data to render HTML
	// on a timeline when you expand the details.
	DetailTemplate param.Opt[string] `json:"detailTemplate,omitzero"`
	// This uses Markdown syntax with Handlebars and event-specific data to render HTML
	// on a timeline as a header.
	HeaderTemplate param.Opt[string] `json:"headerTemplate,omitzero"`
	paramObj
}

func (r TimelineEventTemplateUpdateRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow TimelineEventTemplateUpdateRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TimelineEventTemplateUpdateRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
