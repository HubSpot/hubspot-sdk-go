// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cms

import (
	"encoding/json"
	"time"

	"github.com/stainless-sdks/hubspot-sdk-go/internal/apijson"
	"github.com/stainless-sdks/hubspot-sdk-go/option"
	"github.com/stainless-sdks/hubspot-sdk-go/packages/param"
	"github.com/stainless-sdks/hubspot-sdk-go/packages/respjson"
)

// CmService contains methods and other services that help with interacting with
// the hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCmService] method instead.
type CmService struct {
	Options      []option.RequestOption
	AuditLogs    AuditLogService
	Blogs        BlogService
	Domains      DomainService
	Hubdb        HubdbService
	MediaBridge  MediaBridgeService
	Pages        PageService
	SiteSearch   SiteSearchService
	SourceCode   SourceCodeService
	URLRedirects URLRedirectService
}

// NewCmService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewCmService(opts ...option.RequestOption) (r CmService) {
	r = CmService{}
	r.Options = opts
	r.AuditLogs = NewAuditLogService(opts...)
	r.Blogs = NewBlogService(opts...)
	r.Domains = NewDomainService(opts...)
	r.Hubdb = NewHubdbService(opts...)
	r.MediaBridge = NewMediaBridgeService(opts...)
	r.Pages = NewPageService(opts...)
	r.SiteSearch = NewSiteSearchService(opts...)
	r.SourceCode = NewSourceCodeService(opts...)
	r.URLRedirects = NewURLRedirectService(opts...)
	return
}

type Angle struct {
	Units string  `json:"units,required"`
	Value float64 `json:"value,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Units       respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Angle) RawJSON() string { return r.JSON.raw }
func (r *Angle) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this Angle to a AngleParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// AngleParam.Overrides()
func (r Angle) ToParam() AngleParam {
	return param.Override[AngleParam](json.RawMessage(r.RawJSON()))
}

// The properties Units, Value are required.
type AngleParam struct {
	Units string  `json:"units,required"`
	Value float64 `json:"value,required"`
	paramObj
}

func (r AngleParam) MarshalJSON() (data []byte, err error) {
	type shadow AngleParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AngleParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Request body object for attaching objects to multi-language groups.
//
// The properties ID, Language, PrimaryID are required.
type AttachToLangPrimaryRequestVNextParam struct {
	// ID of the object to add to a multi-language group.
	ID string `json:"id,required"`
	// Designated language of the object to add to a multi-language group.
	Language string `json:"language,required"`
	// ID of primary language object in multi-language group.
	PrimaryID string `json:"primaryId,required"`
	// Primary language of the multi-language group.
	PrimaryLanguage param.Opt[string] `json:"primaryLanguage,omitzero"`
	paramObj
}

func (r AttachToLangPrimaryRequestVNextParam) MarshalJSON() (data []byte, err error) {
	type shadow AttachToLangPrimaryRequestVNextParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AttachToLangPrimaryRequestVNextParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BackgroundImage struct {
	BackgroundPosition string `json:"backgroundPosition,required"`
	BackgroundSize     string `json:"backgroundSize,required"`
	ImageURL           string `json:"imageUrl,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BackgroundPosition respjson.Field
		BackgroundSize     respjson.Field
		ImageURL           respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BackgroundImage) RawJSON() string { return r.JSON.raw }
func (r *BackgroundImage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this BackgroundImage to a BackgroundImageParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// BackgroundImageParam.Overrides()
func (r BackgroundImage) ToParam() BackgroundImageParam {
	return param.Override[BackgroundImageParam](json.RawMessage(r.RawJSON()))
}

// The properties BackgroundPosition, BackgroundSize, ImageURL are required.
type BackgroundImageParam struct {
	BackgroundPosition string `json:"backgroundPosition,required"`
	BackgroundSize     string `json:"backgroundSize,required"`
	ImageURL           string `json:"imageUrl,required"`
	paramObj
}

func (r BackgroundImageParam) MarshalJSON() (data []byte, err error) {
	type shadow BackgroundImageParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BackgroundImageParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Wrapper for providing an array of JSON nodes as inputs.
//
// The property Inputs is required.
type BatchInputJsonNodeParam struct {
	// JSON nodes to input.
	Inputs []any `json:"inputs,omitzero,required"`
	paramObj
}

func (r BatchInputJsonNodeParam) MarshalJSON() (data []byte, err error) {
	type shadow BatchInputJsonNodeParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BatchInputJsonNodeParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ColorStop struct {
	// A color defined by RGB values.
	Color RgbaColor `json:"color,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Color       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ColorStop) RawJSON() string { return r.JSON.raw }
func (r *ColorStop) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this ColorStop to a ColorStopParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// ColorStopParam.Overrides()
func (r ColorStop) ToParam() ColorStopParam {
	return param.Override[ColorStopParam](json.RawMessage(r.RawJSON()))
}

// The property Color is required.
type ColorStopParam struct {
	// A color defined by RGB values.
	Color RgbaColorParam `json:"color,omitzero,required"`
	paramObj
}

func (r ColorStopParam) MarshalJSON() (data []byte, err error) {
	type shadow ColorStopParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ColorStopParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Request body object for cloning content.
//
// The property ID is required.
type ContentCloneRequestVNextParam struct {
	// ID of the object to be cloned.
	ID string `json:"id,required"`
	// Name of the cloned object.
	CloneName param.Opt[string] `json:"cloneName,omitzero"`
	paramObj
}

func (r ContentCloneRequestVNextParam) MarshalJSON() (data []byte, err error) {
	type shadow ContentCloneRequestVNextParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ContentCloneRequestVNextParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Request body object for scheduling the publish of content
//
// The properties ID, PublishDate are required.
type ContentScheduleRequestVNextParam struct {
	// The ID of the object to be scheduled.
	ID string `json:"id,required"`
	// The date the object should transition from scheduled to published.
	PublishDate time.Time `json:"publishDate,required" format:"date-time"`
	paramObj
}

func (r ContentScheduleRequestVNextParam) MarshalJSON() (data []byte, err error) {
	type shadow ContentScheduleRequestVNextParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ContentScheduleRequestVNextParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Request body object for detaching objects from multi-language groups.
//
// The property ID is required.
type DetachFromLangGroupRequestVNextParam struct {
	// ID of the object to remove from a multi-language group.
	ID string `json:"id,required"`
	paramObj
}

func (r DetachFromLangGroupRequestVNextParam) MarshalJSON() (data []byte, err error) {
	type shadow DetachFromLangGroupRequestVNextParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DetachFromLangGroupRequestVNextParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Gradient struct {
	Angle        Angle        `json:"angle,required"`
	Colors       []ColorStop  `json:"colors,required"`
	SideOrCorner SideOrCorner `json:"sideOrCorner,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Angle        respjson.Field
		Colors       respjson.Field
		SideOrCorner respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Gradient) RawJSON() string { return r.JSON.raw }
func (r *Gradient) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this Gradient to a GradientParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// GradientParam.Overrides()
func (r Gradient) ToParam() GradientParam {
	return param.Override[GradientParam](json.RawMessage(r.RawJSON()))
}

// The properties Angle, Colors, SideOrCorner are required.
type GradientParam struct {
	Angle        AngleParam        `json:"angle,omitzero,required"`
	Colors       []ColorStopParam  `json:"colors,omitzero,required"`
	SideOrCorner SideOrCornerParam `json:"sideOrCorner,omitzero,required"`
	paramObj
}

func (r GradientParam) MarshalJSON() (data []byte, err error) {
	type shadow GradientParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *GradientParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LayoutSection struct {
	Cells    []LayoutSection `json:"cells,required"`
	CssClass string          `json:"cssClass,required"`
	CssID    string          `json:"cssId,required"`
	CssStyle string          `json:"cssStyle,required"`
	Label    string          `json:"label,required"`
	Name     string          `json:"name,required"`
	// null
	Params      map[string]any             `json:"params,required"`
	RowMetaData []RowMetaData              `json:"rowMetaData,required"`
	Rows        []map[string]LayoutSection `json:"rows,required"`
	Styles      Styles                     `json:"styles,required"`
	Type        string                     `json:"type,required"`
	W           int64                      `json:"w,required"`
	X           int64                      `json:"x,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Cells       respjson.Field
		CssClass    respjson.Field
		CssID       respjson.Field
		CssStyle    respjson.Field
		Label       respjson.Field
		Name        respjson.Field
		Params      respjson.Field
		RowMetaData respjson.Field
		Rows        respjson.Field
		Styles      respjson.Field
		Type        respjson.Field
		W           respjson.Field
		X           respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LayoutSection) RawJSON() string { return r.JSON.raw }
func (r *LayoutSection) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this LayoutSection to a LayoutSectionParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// LayoutSectionParam.Overrides()
func (r LayoutSection) ToParam() LayoutSectionParam {
	return param.Override[LayoutSectionParam](json.RawMessage(r.RawJSON()))
}

// The properties Cells, CssClass, CssID, CssStyle, Label, Name, Params,
// RowMetaData, Rows, Styles, Type, W, X are required.
type LayoutSectionParam struct {
	Cells    []LayoutSectionParam `json:"cells,omitzero,required"`
	CssClass string               `json:"cssClass,required"`
	CssID    string               `json:"cssId,required"`
	CssStyle string               `json:"cssStyle,required"`
	Label    string               `json:"label,required"`
	Name     string               `json:"name,required"`
	// null
	Params      map[string]any                  `json:"params,omitzero,required"`
	RowMetaData []RowMetaDataParam              `json:"rowMetaData,omitzero,required"`
	Rows        []map[string]LayoutSectionParam `json:"rows,omitzero,required"`
	Styles      StylesParam                     `json:"styles,omitzero,required"`
	Type        string                          `json:"type,required"`
	W           int64                           `json:"w,required"`
	X           int64                           `json:"x,required"`
	paramObj
}

func (r LayoutSectionParam) MarshalJSON() (data []byte, err error) {
	type shadow LayoutSectionParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LayoutSectionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PublicAccessRule = any

// A color defined by RGB values.
type RgbaColor struct {
	// Alpha.
	A float64 `json:"a,required"`
	// Blue.
	B int64 `json:"b,required"`
	// Green.
	G int64 `json:"g,required"`
	// Red.
	R int64 `json:"r,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		A           respjson.Field
		B           respjson.Field
		G           respjson.Field
		R           respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RgbaColor) RawJSON() string { return r.JSON.raw }
func (r *RgbaColor) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this RgbaColor to a RgbaColorParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// RgbaColorParam.Overrides()
func (r RgbaColor) ToParam() RgbaColorParam {
	return param.Override[RgbaColorParam](json.RawMessage(r.RawJSON()))
}

// A color defined by RGB values.
//
// The properties A, B, G, R are required.
type RgbaColorParam struct {
	// Alpha.
	A float64 `json:"a,required"`
	// Blue.
	B int64 `json:"b,required"`
	// Green.
	G int64 `json:"g,required"`
	// Red.
	R int64 `json:"r,required"`
	paramObj
}

func (r RgbaColorParam) MarshalJSON() (data []byte, err error) {
	type shadow RgbaColorParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RgbaColorParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RowMetaData struct {
	CssClass string `json:"cssClass,required"`
	Styles   Styles `json:"styles,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CssClass    respjson.Field
		Styles      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RowMetaData) RawJSON() string { return r.JSON.raw }
func (r *RowMetaData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this RowMetaData to a RowMetaDataParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// RowMetaDataParam.Overrides()
func (r RowMetaData) ToParam() RowMetaDataParam {
	return param.Override[RowMetaDataParam](json.RawMessage(r.RawJSON()))
}

// The properties CssClass, Styles are required.
type RowMetaDataParam struct {
	CssClass string      `json:"cssClass,required"`
	Styles   StylesParam `json:"styles,omitzero,required"`
	paramObj
}

func (r RowMetaDataParam) MarshalJSON() (data []byte, err error) {
	type shadow RowMetaDataParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RowMetaDataParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Request body object for setting a new primary language.
//
// The property ID is required.
type SetNewLanguagePrimaryRequestVNextParam struct {
	// ID of object to set as primary in multi-language group.
	ID string `json:"id,required"`
	paramObj
}

func (r SetNewLanguagePrimaryRequestVNextParam) MarshalJSON() (data []byte, err error) {
	type shadow SetNewLanguagePrimaryRequestVNextParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SetNewLanguagePrimaryRequestVNextParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SideOrCorner struct {
	HorizontalSide string `json:"horizontalSide,required"`
	VerticalSide   string `json:"verticalSide,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		HorizontalSide respjson.Field
		VerticalSide   respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SideOrCorner) RawJSON() string { return r.JSON.raw }
func (r *SideOrCorner) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this SideOrCorner to a SideOrCornerParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// SideOrCornerParam.Overrides()
func (r SideOrCorner) ToParam() SideOrCornerParam {
	return param.Override[SideOrCornerParam](json.RawMessage(r.RawJSON()))
}

// The properties HorizontalSide, VerticalSide are required.
type SideOrCornerParam struct {
	HorizontalSide string `json:"horizontalSide,required"`
	VerticalSide   string `json:"verticalSide,required"`
	paramObj
}

func (r SideOrCornerParam) MarshalJSON() (data []byte, err error) {
	type shadow SideOrCornerParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SideOrCornerParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Styles struct {
	// A color defined by RGB values.
	BackgroundColor          RgbaColor                   `json:"backgroundColor,required"`
	BackgroundGradient       Gradient                    `json:"backgroundGradient,required"`
	BackgroundImage          BackgroundImage             `json:"backgroundImage,required"`
	FlexboxPositioning       string                      `json:"flexboxPositioning,required"`
	ForceFullWidthSection    bool                        `json:"forceFullWidthSection,required"`
	MaxWidthSectionCentering int64                       `json:"maxWidthSectionCentering,required"`
	VerticalAlignment        string                      `json:"verticalAlignment,required"`
	BreakpointStyles         map[string]BreakpointStyles `json:"breakpointStyles"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BackgroundColor          respjson.Field
		BackgroundGradient       respjson.Field
		BackgroundImage          respjson.Field
		FlexboxPositioning       respjson.Field
		ForceFullWidthSection    respjson.Field
		MaxWidthSectionCentering respjson.Field
		VerticalAlignment        respjson.Field
		BreakpointStyles         respjson.Field
		ExtraFields              map[string]respjson.Field
		raw                      string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Styles) RawJSON() string { return r.JSON.raw }
func (r *Styles) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this Styles to a StylesParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// StylesParam.Overrides()
func (r Styles) ToParam() StylesParam {
	return param.Override[StylesParam](json.RawMessage(r.RawJSON()))
}

// The properties BackgroundColor, BackgroundGradient, BackgroundImage,
// FlexboxPositioning, ForceFullWidthSection, MaxWidthSectionCentering,
// VerticalAlignment are required.
type StylesParam struct {
	// A color defined by RGB values.
	BackgroundColor          RgbaColorParam                   `json:"backgroundColor,omitzero,required"`
	BackgroundGradient       GradientParam                    `json:"backgroundGradient,omitzero,required"`
	BackgroundImage          BackgroundImageParam             `json:"backgroundImage,omitzero,required"`
	FlexboxPositioning       string                           `json:"flexboxPositioning,required"`
	ForceFullWidthSection    bool                             `json:"forceFullWidthSection,required"`
	MaxWidthSectionCentering int64                            `json:"maxWidthSectionCentering,required"`
	VerticalAlignment        string                           `json:"verticalAlignment,required"`
	BreakpointStyles         map[string]BreakpointStylesParam `json:"breakpointStyles,omitzero"`
	paramObj
}

func (r StylesParam) MarshalJSON() (data []byte, err error) {
	type shadow StylesParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *StylesParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Request object for updating languages within a multi-language group.
//
// The properties Languages, PrimaryID are required.
type UpdateLanguagesRequestVNextParam struct {
	// Map of object IDs to associated languages of object in the multi-language group.
	Languages map[string]string `json:"languages,omitzero,required"`
	// ID of the primary object in the multi-language group.
	PrimaryID string `json:"primaryId,required"`
	paramObj
}

func (r UpdateLanguagesRequestVNextParam) MarshalJSON() (data []byte, err error) {
	type shadow UpdateLanguagesRequestVNextParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UpdateLanguagesRequestVNextParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
