// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package marketing

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
	"github.com/stainless-sdks/hubspot-sdk-go/internal/apiquery"
	shimjson "github.com/stainless-sdks/hubspot-sdk-go/internal/encoding/json"
	"github.com/stainless-sdks/hubspot-sdk-go/internal/requestconfig"
	"github.com/stainless-sdks/hubspot-sdk-go/option"
	"github.com/stainless-sdks/hubspot-sdk-go/packages/pagination"
	"github.com/stainless-sdks/hubspot-sdk-go/packages/param"
	"github.com/stainless-sdks/hubspot-sdk-go/packages/respjson"
	"github.com/stainless-sdks/hubspot-sdk-go/shared"
)

// FormService contains methods and other services that help with interacting with
// the Hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewFormService] method instead.
type FormService struct {
	Options []option.RequestOption
}

// NewFormService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewFormService(opts ...option.RequestOption) (r FormService) {
	r = FormService{}
	r.Options = opts
	return
}

// Add a new `hubspot` form
func (r *FormService) New(ctx context.Context, body FormNewParams, opts ...option.RequestOption) (res *FormDefinitionBase, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "marketing/v3/forms/"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Update some of the form definition components
func (r *FormService) Update(ctx context.Context, formID string, body FormUpdateParams, opts ...option.RequestOption) (res *FormDefinitionBase, err error) {
	opts = slices.Concat(r.Options, opts)
	if formID == "" {
		err = errors.New("missing required formId parameter")
		return
	}
	path := fmt.Sprintf("marketing/v3/forms/%s", formID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Returns a list of forms based on the search filters. By default, it returns the
// first 20 `hubspot` forms
func (r *FormService) List(ctx context.Context, query FormListParams, opts ...option.RequestOption) (res *pagination.Page[HubSpotFormDefinition], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "marketing/v3/forms/"
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

// Returns a list of forms based on the search filters. By default, it returns the
// first 20 `hubspot` forms
func (r *FormService) ListAutoPaging(ctx context.Context, query FormListParams, opts ...option.RequestOption) *pagination.PageAutoPager[HubSpotFormDefinition] {
	return pagination.NewPageAutoPager(r.List(ctx, query, opts...))
}

// Archive a form definition. New submissions will not be accepted and the form
// definition will be permanently deleted after 3 months.
func (r *FormService) Delete(ctx context.Context, formID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if formID == "" {
		err = errors.New("missing required formId parameter")
		return
	}
	path := fmt.Sprintf("marketing/v3/forms/%s", formID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Returns a form based on the form ID provided.
func (r *FormService) Get(ctx context.Context, formID string, query FormGetParams, opts ...option.RequestOption) (res *FormDefinitionBase, err error) {
	opts = slices.Concat(r.Options, opts)
	if formID == "" {
		err = errors.New("missing required formId parameter")
		return
	}
	path := fmt.Sprintf("marketing/v3/forms/%s", formID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Update all fields of a hubspot form definition.
func (r *FormService) Replace(ctx context.Context, formID string, body FormReplaceParams, opts ...option.RequestOption) (res *FormDefinitionBase, err error) {
	opts = slices.Concat(r.Options, opts)
	if formID == "" {
		err = errors.New("missing required formId parameter")
		return
	}
	path := fmt.Sprintf("marketing/v3/forms/%s", formID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

type CollectionResponseFormDefinitionBaseForwardPaging struct {
	Results []HubSpotFormDefinition `json:"results,required"`
	Paging  shared.ForwardPaging    `json:"paging"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Results     respjson.Field
		Paging      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CollectionResponseFormDefinitionBaseForwardPaging) RawJSON() string { return r.JSON.raw }
func (r *CollectionResponseFormDefinitionBaseForwardPaging) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A form field used to select a date
type DatepickerField struct {
	// A list of other fields to make visible based on the value filled in for this
	// field.
	DependentFields []DependentField `json:"dependentFields,required"`
	// Determines how the field will be displayed and validated.
	//
	// Any of "datepicker".
	FieldType DatepickerFieldFieldType `json:"fieldType,required"`
	// Whether a field should be hidden or not. Hidden fields won't appear on the form,
	// but can be used to pass a value to a property without requiring the customer to
	// fill it in.
	Hidden bool `json:"hidden,required"`
	// The main label for the form field.
	Label string `json:"label,required"`
	// The identifier of the field. In combination with the object type ID, it must be
	// unique.
	Name string `json:"name,required"`
	// A unique ID for this field's CRM object type. For example a CONTACT field will
	// have the object type ID 0-1.
	ObjectTypeID string `json:"objectTypeId,required"`
	// Whether a value for this field is required when submitting the form.
	Required bool `json:"required,required"`
	// The value filled in by default. This value will be submitted unless the customer
	// modifies it.
	DefaultValue string `json:"defaultValue"`
	// Additional text helping the customer to complete the field.
	Description string `json:"description"`
	// The prompt text showing when the field isn't filled in.
	Placeholder string `json:"placeholder"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DependentFields respjson.Field
		FieldType       respjson.Field
		Hidden          respjson.Field
		Label           respjson.Field
		Name            respjson.Field
		ObjectTypeID    respjson.Field
		Required        respjson.Field
		DefaultValue    respjson.Field
		Description     respjson.Field
		Placeholder     respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DatepickerField) RawJSON() string { return r.JSON.raw }
func (r *DatepickerField) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this DatepickerField to a DatepickerFieldParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// DatepickerFieldParam.Overrides()
func (r DatepickerField) ToParam() DatepickerFieldParam {
	return param.Override[DatepickerFieldParam](json.RawMessage(r.RawJSON()))
}

// Determines how the field will be displayed and validated.
type DatepickerFieldFieldType string

const (
	DatepickerFieldFieldTypeDatepicker DatepickerFieldFieldType = "datepicker"
)

// A form field used to select a date
//
// The properties DependentFields, FieldType, Hidden, Label, Name, ObjectTypeID,
// Required are required.
type DatepickerFieldParam struct {
	// A list of other fields to make visible based on the value filled in for this
	// field.
	DependentFields []DependentFieldParam `json:"dependentFields,omitzero,required"`
	// Determines how the field will be displayed and validated.
	//
	// Any of "datepicker".
	FieldType DatepickerFieldFieldType `json:"fieldType,omitzero,required"`
	// Whether a field should be hidden or not. Hidden fields won't appear on the form,
	// but can be used to pass a value to a property without requiring the customer to
	// fill it in.
	Hidden bool `json:"hidden,required"`
	// The main label for the form field.
	Label string `json:"label,required"`
	// The identifier of the field. In combination with the object type ID, it must be
	// unique.
	Name string `json:"name,required"`
	// A unique ID for this field's CRM object type. For example a CONTACT field will
	// have the object type ID 0-1.
	ObjectTypeID string `json:"objectTypeId,required"`
	// Whether a value for this field is required when submitting the form.
	Required bool `json:"required,required"`
	// The value filled in by default. This value will be submitted unless the customer
	// modifies it.
	DefaultValue param.Opt[string] `json:"defaultValue,omitzero"`
	// Additional text helping the customer to complete the field.
	Description param.Opt[string] `json:"description,omitzero"`
	// The prompt text showing when the field isn't filled in.
	Placeholder param.Opt[string] `json:"placeholder,omitzero"`
	paramObj
}

func (r DatepickerFieldParam) MarshalJSON() (data []byte, err error) {
	type shadow DatepickerFieldParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DatepickerFieldParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A form field that will be displayed based on what the customer entered in
// another field.
type DependentField struct {
	// A condition based on customer input
	DependentCondition DependentFieldFilter `json:"dependentCondition,required"`
	// A form field used for collecting an email address.
	DependentField DependentFieldDependentFieldUnion `json:"dependentField,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DependentCondition respjson.Field
		DependentField     respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DependentField) RawJSON() string { return r.JSON.raw }
func (r *DependentField) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this DependentField to a DependentFieldParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// DependentFieldParam.Overrides()
func (r DependentField) ToParam() DependentFieldParam {
	return param.Override[DependentFieldParam](json.RawMessage(r.RawJSON()))
}

// DependentFieldDependentFieldUnion contains all possible properties and values
// from [EmailField], [PhoneField], [MobilePhoneField], [SingleLineTextField],
// [MultiLineTextField], [NumberField], [SingleCheckboxField],
// [MultipleCheckboxesField], [DropdownField], [RadioField], [DatepickerField],
// [FileField], [PaymentLinkRadioField].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type DependentFieldDependentFieldUnion struct {
	DependentFields []DependentField `json:"dependentFields"`
	FieldType       string           `json:"fieldType"`
	Hidden          bool             `json:"hidden"`
	Label           string           `json:"label"`
	Name            string           `json:"name"`
	ObjectTypeID    string           `json:"objectTypeId"`
	Required        bool             `json:"required"`
	// This field is a union of [EmailFieldValidation], [PhoneFieldValidation],
	// [NumberFieldValidation]
	Validation   DependentFieldDependentFieldUnionValidation `json:"validation"`
	DefaultValue string                                      `json:"defaultValue"`
	Description  string                                      `json:"description"`
	Placeholder  string                                      `json:"placeholder"`
	// This field is from variant [PhoneField].
	UseCountryCodeSelect bool                    `json:"useCountryCodeSelect"`
	DefaultValues        []string                `json:"defaultValues"`
	Options              []EnumeratedFieldOption `json:"options"`
	// This field is from variant [FileField].
	AllowMultipleFiles bool `json:"allowMultipleFiles"`
	JSON               struct {
		DependentFields      respjson.Field
		FieldType            respjson.Field
		Hidden               respjson.Field
		Label                respjson.Field
		Name                 respjson.Field
		ObjectTypeID         respjson.Field
		Required             respjson.Field
		Validation           respjson.Field
		DefaultValue         respjson.Field
		Description          respjson.Field
		Placeholder          respjson.Field
		UseCountryCodeSelect respjson.Field
		DefaultValues        respjson.Field
		Options              respjson.Field
		AllowMultipleFiles   respjson.Field
		raw                  string
	} `json:"-"`
}

func (u DependentFieldDependentFieldUnion) AsEmail() (v EmailField) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u DependentFieldDependentFieldUnion) AsPhone() (v PhoneField) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u DependentFieldDependentFieldUnion) AsMobilePhone() (v MobilePhoneField) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u DependentFieldDependentFieldUnion) AsSingleLineText() (v SingleLineTextField) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u DependentFieldDependentFieldUnion) AsMultiLineText() (v MultiLineTextField) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u DependentFieldDependentFieldUnion) AsNumber() (v NumberField) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u DependentFieldDependentFieldUnion) AsSingleCheckbox() (v SingleCheckboxField) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u DependentFieldDependentFieldUnion) AsMultipleCheckboxes() (v MultipleCheckboxesField) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u DependentFieldDependentFieldUnion) AsDropdown() (v DropdownField) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u DependentFieldDependentFieldUnion) AsRadio() (v RadioField) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u DependentFieldDependentFieldUnion) AsDatepicker() (v DatepickerField) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u DependentFieldDependentFieldUnion) AsFile() (v FileField) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u DependentFieldDependentFieldUnion) AsPaymentLinkRadio() (v PaymentLinkRadioField) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u DependentFieldDependentFieldUnion) RawJSON() string { return u.JSON.raw }

func (r *DependentFieldDependentFieldUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// DependentFieldDependentFieldUnionValidation is an implicit subunion of
// [DependentFieldDependentFieldUnion]. DependentFieldDependentFieldUnionValidation
// provides convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [DependentFieldDependentFieldUnion].
type DependentFieldDependentFieldUnionValidation struct {
	// This field is from variant [EmailFieldValidation].
	BlockedEmailDomains []string `json:"blockedEmailDomains"`
	// This field is from variant [EmailFieldValidation].
	UseDefaultBlockList bool  `json:"useDefaultBlockList"`
	MaxAllowedDigits    int64 `json:"maxAllowedDigits"`
	MinAllowedDigits    int64 `json:"minAllowedDigits"`
	JSON                struct {
		BlockedEmailDomains respjson.Field
		UseDefaultBlockList respjson.Field
		MaxAllowedDigits    respjson.Field
		MinAllowedDigits    respjson.Field
		raw                 string
	} `json:"-"`
}

func (r *DependentFieldDependentFieldUnionValidation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A form field that will be displayed based on what the customer entered in
// another field.
//
// The properties DependentCondition, DependentField are required.
type DependentFieldParam struct {
	// A condition based on customer input
	DependentCondition DependentFieldFilterParam `json:"dependentCondition,omitzero,required"`
	// A form field used for collecting an email address.
	DependentField DependentFieldDependentFieldUnionParam `json:"dependentField,omitzero,required"`
	paramObj
}

func (r DependentFieldParam) MarshalJSON() (data []byte, err error) {
	type shadow DependentFieldParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DependentFieldParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type DependentFieldDependentFieldUnionParam struct {
	OfEmail              *EmailFieldParam              `json:",omitzero,inline"`
	OfPhone              *PhoneFieldParam              `json:",omitzero,inline"`
	OfMobilePhone        *MobilePhoneFieldParam        `json:",omitzero,inline"`
	OfSingleLineText     *SingleLineTextFieldParam     `json:",omitzero,inline"`
	OfMultiLineText      *MultiLineTextFieldParam      `json:",omitzero,inline"`
	OfNumber             *NumberFieldParam             `json:",omitzero,inline"`
	OfSingleCheckbox     *SingleCheckboxFieldParam     `json:",omitzero,inline"`
	OfMultipleCheckboxes *MultipleCheckboxesFieldParam `json:",omitzero,inline"`
	OfDropdown           *DropdownFieldParam           `json:",omitzero,inline"`
	OfRadio              *RadioFieldParam              `json:",omitzero,inline"`
	OfDatepicker         *DatepickerFieldParam         `json:",omitzero,inline"`
	OfFile               *FileFieldParam               `json:",omitzero,inline"`
	OfPaymentLinkRadio   *PaymentLinkRadioFieldParam   `json:",omitzero,inline"`
	paramUnion
}

func (u DependentFieldDependentFieldUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfEmail,
		u.OfPhone,
		u.OfMobilePhone,
		u.OfSingleLineText,
		u.OfMultiLineText,
		u.OfNumber,
		u.OfSingleCheckbox,
		u.OfMultipleCheckboxes,
		u.OfDropdown,
		u.OfRadio,
		u.OfDatepicker,
		u.OfFile,
		u.OfPaymentLinkRadio)
}
func (u *DependentFieldDependentFieldUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *DependentFieldDependentFieldUnionParam) asAny() any {
	if !param.IsOmitted(u.OfEmail) {
		return u.OfEmail
	} else if !param.IsOmitted(u.OfPhone) {
		return u.OfPhone
	} else if !param.IsOmitted(u.OfMobilePhone) {
		return u.OfMobilePhone
	} else if !param.IsOmitted(u.OfSingleLineText) {
		return u.OfSingleLineText
	} else if !param.IsOmitted(u.OfMultiLineText) {
		return u.OfMultiLineText
	} else if !param.IsOmitted(u.OfNumber) {
		return u.OfNumber
	} else if !param.IsOmitted(u.OfSingleCheckbox) {
		return u.OfSingleCheckbox
	} else if !param.IsOmitted(u.OfMultipleCheckboxes) {
		return u.OfMultipleCheckboxes
	} else if !param.IsOmitted(u.OfDropdown) {
		return u.OfDropdown
	} else if !param.IsOmitted(u.OfRadio) {
		return u.OfRadio
	} else if !param.IsOmitted(u.OfDatepicker) {
		return u.OfDatepicker
	} else if !param.IsOmitted(u.OfFile) {
		return u.OfFile
	} else if !param.IsOmitted(u.OfPaymentLinkRadio) {
		return u.OfPaymentLinkRadio
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u DependentFieldDependentFieldUnionParam) GetUseCountryCodeSelect() *bool {
	if vt := u.OfPhone; vt != nil {
		return &vt.UseCountryCodeSelect
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u DependentFieldDependentFieldUnionParam) GetAllowMultipleFiles() *bool {
	if vt := u.OfFile; vt != nil {
		return &vt.AllowMultipleFiles
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u DependentFieldDependentFieldUnionParam) GetFieldType() *string {
	if vt := u.OfEmail; vt != nil {
		return (*string)(&vt.FieldType)
	} else if vt := u.OfPhone; vt != nil {
		return (*string)(&vt.FieldType)
	} else if vt := u.OfMobilePhone; vt != nil {
		return (*string)(&vt.FieldType)
	} else if vt := u.OfSingleLineText; vt != nil {
		return (*string)(&vt.FieldType)
	} else if vt := u.OfMultiLineText; vt != nil {
		return (*string)(&vt.FieldType)
	} else if vt := u.OfNumber; vt != nil {
		return (*string)(&vt.FieldType)
	} else if vt := u.OfSingleCheckbox; vt != nil {
		return (*string)(&vt.FieldType)
	} else if vt := u.OfMultipleCheckboxes; vt != nil {
		return (*string)(&vt.FieldType)
	} else if vt := u.OfDropdown; vt != nil {
		return (*string)(&vt.FieldType)
	} else if vt := u.OfRadio; vt != nil {
		return (*string)(&vt.FieldType)
	} else if vt := u.OfDatepicker; vt != nil {
		return (*string)(&vt.FieldType)
	} else if vt := u.OfFile; vt != nil {
		return (*string)(&vt.FieldType)
	} else if vt := u.OfPaymentLinkRadio; vt != nil {
		return (*string)(&vt.FieldType)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u DependentFieldDependentFieldUnionParam) GetHidden() *bool {
	if vt := u.OfEmail; vt != nil {
		return (*bool)(&vt.Hidden)
	} else if vt := u.OfPhone; vt != nil {
		return (*bool)(&vt.Hidden)
	} else if vt := u.OfMobilePhone; vt != nil {
		return (*bool)(&vt.Hidden)
	} else if vt := u.OfSingleLineText; vt != nil {
		return (*bool)(&vt.Hidden)
	} else if vt := u.OfMultiLineText; vt != nil {
		return (*bool)(&vt.Hidden)
	} else if vt := u.OfNumber; vt != nil {
		return (*bool)(&vt.Hidden)
	} else if vt := u.OfSingleCheckbox; vt != nil {
		return (*bool)(&vt.Hidden)
	} else if vt := u.OfMultipleCheckboxes; vt != nil {
		return (*bool)(&vt.Hidden)
	} else if vt := u.OfDropdown; vt != nil {
		return (*bool)(&vt.Hidden)
	} else if vt := u.OfRadio; vt != nil {
		return (*bool)(&vt.Hidden)
	} else if vt := u.OfDatepicker; vt != nil {
		return (*bool)(&vt.Hidden)
	} else if vt := u.OfFile; vt != nil {
		return (*bool)(&vt.Hidden)
	} else if vt := u.OfPaymentLinkRadio; vt != nil {
		return (*bool)(&vt.Hidden)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u DependentFieldDependentFieldUnionParam) GetLabel() *string {
	if vt := u.OfEmail; vt != nil {
		return (*string)(&vt.Label)
	} else if vt := u.OfPhone; vt != nil {
		return (*string)(&vt.Label)
	} else if vt := u.OfMobilePhone; vt != nil {
		return (*string)(&vt.Label)
	} else if vt := u.OfSingleLineText; vt != nil {
		return (*string)(&vt.Label)
	} else if vt := u.OfMultiLineText; vt != nil {
		return (*string)(&vt.Label)
	} else if vt := u.OfNumber; vt != nil {
		return (*string)(&vt.Label)
	} else if vt := u.OfSingleCheckbox; vt != nil {
		return (*string)(&vt.Label)
	} else if vt := u.OfMultipleCheckboxes; vt != nil {
		return (*string)(&vt.Label)
	} else if vt := u.OfDropdown; vt != nil {
		return (*string)(&vt.Label)
	} else if vt := u.OfRadio; vt != nil {
		return (*string)(&vt.Label)
	} else if vt := u.OfDatepicker; vt != nil {
		return (*string)(&vt.Label)
	} else if vt := u.OfFile; vt != nil {
		return (*string)(&vt.Label)
	} else if vt := u.OfPaymentLinkRadio; vt != nil {
		return (*string)(&vt.Label)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u DependentFieldDependentFieldUnionParam) GetName() *string {
	if vt := u.OfEmail; vt != nil {
		return (*string)(&vt.Name)
	} else if vt := u.OfPhone; vt != nil {
		return (*string)(&vt.Name)
	} else if vt := u.OfMobilePhone; vt != nil {
		return (*string)(&vt.Name)
	} else if vt := u.OfSingleLineText; vt != nil {
		return (*string)(&vt.Name)
	} else if vt := u.OfMultiLineText; vt != nil {
		return (*string)(&vt.Name)
	} else if vt := u.OfNumber; vt != nil {
		return (*string)(&vt.Name)
	} else if vt := u.OfSingleCheckbox; vt != nil {
		return (*string)(&vt.Name)
	} else if vt := u.OfMultipleCheckboxes; vt != nil {
		return (*string)(&vt.Name)
	} else if vt := u.OfDropdown; vt != nil {
		return (*string)(&vt.Name)
	} else if vt := u.OfRadio; vt != nil {
		return (*string)(&vt.Name)
	} else if vt := u.OfDatepicker; vt != nil {
		return (*string)(&vt.Name)
	} else if vt := u.OfFile; vt != nil {
		return (*string)(&vt.Name)
	} else if vt := u.OfPaymentLinkRadio; vt != nil {
		return (*string)(&vt.Name)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u DependentFieldDependentFieldUnionParam) GetObjectTypeID() *string {
	if vt := u.OfEmail; vt != nil {
		return (*string)(&vt.ObjectTypeID)
	} else if vt := u.OfPhone; vt != nil {
		return (*string)(&vt.ObjectTypeID)
	} else if vt := u.OfMobilePhone; vt != nil {
		return (*string)(&vt.ObjectTypeID)
	} else if vt := u.OfSingleLineText; vt != nil {
		return (*string)(&vt.ObjectTypeID)
	} else if vt := u.OfMultiLineText; vt != nil {
		return (*string)(&vt.ObjectTypeID)
	} else if vt := u.OfNumber; vt != nil {
		return (*string)(&vt.ObjectTypeID)
	} else if vt := u.OfSingleCheckbox; vt != nil {
		return (*string)(&vt.ObjectTypeID)
	} else if vt := u.OfMultipleCheckboxes; vt != nil {
		return (*string)(&vt.ObjectTypeID)
	} else if vt := u.OfDropdown; vt != nil {
		return (*string)(&vt.ObjectTypeID)
	} else if vt := u.OfRadio; vt != nil {
		return (*string)(&vt.ObjectTypeID)
	} else if vt := u.OfDatepicker; vt != nil {
		return (*string)(&vt.ObjectTypeID)
	} else if vt := u.OfFile; vt != nil {
		return (*string)(&vt.ObjectTypeID)
	} else if vt := u.OfPaymentLinkRadio; vt != nil {
		return (*string)(&vt.ObjectTypeID)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u DependentFieldDependentFieldUnionParam) GetRequired() *bool {
	if vt := u.OfEmail; vt != nil {
		return (*bool)(&vt.Required)
	} else if vt := u.OfPhone; vt != nil {
		return (*bool)(&vt.Required)
	} else if vt := u.OfMobilePhone; vt != nil {
		return (*bool)(&vt.Required)
	} else if vt := u.OfSingleLineText; vt != nil {
		return (*bool)(&vt.Required)
	} else if vt := u.OfMultiLineText; vt != nil {
		return (*bool)(&vt.Required)
	} else if vt := u.OfNumber; vt != nil {
		return (*bool)(&vt.Required)
	} else if vt := u.OfSingleCheckbox; vt != nil {
		return (*bool)(&vt.Required)
	} else if vt := u.OfMultipleCheckboxes; vt != nil {
		return (*bool)(&vt.Required)
	} else if vt := u.OfDropdown; vt != nil {
		return (*bool)(&vt.Required)
	} else if vt := u.OfRadio; vt != nil {
		return (*bool)(&vt.Required)
	} else if vt := u.OfDatepicker; vt != nil {
		return (*bool)(&vt.Required)
	} else if vt := u.OfFile; vt != nil {
		return (*bool)(&vt.Required)
	} else if vt := u.OfPaymentLinkRadio; vt != nil {
		return (*bool)(&vt.Required)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u DependentFieldDependentFieldUnionParam) GetDefaultValue() *string {
	if vt := u.OfEmail; vt != nil && vt.DefaultValue.Valid() {
		return &vt.DefaultValue.Value
	} else if vt := u.OfPhone; vt != nil && vt.DefaultValue.Valid() {
		return &vt.DefaultValue.Value
	} else if vt := u.OfMobilePhone; vt != nil && vt.DefaultValue.Valid() {
		return &vt.DefaultValue.Value
	} else if vt := u.OfSingleLineText; vt != nil && vt.DefaultValue.Valid() {
		return &vt.DefaultValue.Value
	} else if vt := u.OfMultiLineText; vt != nil && vt.DefaultValue.Valid() {
		return &vt.DefaultValue.Value
	} else if vt := u.OfNumber; vt != nil && vt.DefaultValue.Valid() {
		return &vt.DefaultValue.Value
	} else if vt := u.OfSingleCheckbox; vt != nil && vt.DefaultValue.Valid() {
		return &vt.DefaultValue.Value
	} else if vt := u.OfDatepicker; vt != nil && vt.DefaultValue.Valid() {
		return &vt.DefaultValue.Value
	} else if vt := u.OfFile; vt != nil && vt.DefaultValue.Valid() {
		return &vt.DefaultValue.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u DependentFieldDependentFieldUnionParam) GetDescription() *string {
	if vt := u.OfEmail; vt != nil && vt.Description.Valid() {
		return &vt.Description.Value
	} else if vt := u.OfPhone; vt != nil && vt.Description.Valid() {
		return &vt.Description.Value
	} else if vt := u.OfMobilePhone; vt != nil && vt.Description.Valid() {
		return &vt.Description.Value
	} else if vt := u.OfSingleLineText; vt != nil && vt.Description.Valid() {
		return &vt.Description.Value
	} else if vt := u.OfMultiLineText; vt != nil && vt.Description.Valid() {
		return &vt.Description.Value
	} else if vt := u.OfNumber; vt != nil && vt.Description.Valid() {
		return &vt.Description.Value
	} else if vt := u.OfSingleCheckbox; vt != nil && vt.Description.Valid() {
		return &vt.Description.Value
	} else if vt := u.OfMultipleCheckboxes; vt != nil && vt.Description.Valid() {
		return &vt.Description.Value
	} else if vt := u.OfDropdown; vt != nil && vt.Description.Valid() {
		return &vt.Description.Value
	} else if vt := u.OfRadio; vt != nil && vt.Description.Valid() {
		return &vt.Description.Value
	} else if vt := u.OfDatepicker; vt != nil && vt.Description.Valid() {
		return &vt.Description.Value
	} else if vt := u.OfFile; vt != nil && vt.Description.Valid() {
		return &vt.Description.Value
	} else if vt := u.OfPaymentLinkRadio; vt != nil && vt.Description.Valid() {
		return &vt.Description.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u DependentFieldDependentFieldUnionParam) GetPlaceholder() *string {
	if vt := u.OfEmail; vt != nil && vt.Placeholder.Valid() {
		return &vt.Placeholder.Value
	} else if vt := u.OfPhone; vt != nil && vt.Placeholder.Valid() {
		return &vt.Placeholder.Value
	} else if vt := u.OfMobilePhone; vt != nil && vt.Placeholder.Valid() {
		return &vt.Placeholder.Value
	} else if vt := u.OfSingleLineText; vt != nil && vt.Placeholder.Valid() {
		return &vt.Placeholder.Value
	} else if vt := u.OfMultiLineText; vt != nil && vt.Placeholder.Valid() {
		return &vt.Placeholder.Value
	} else if vt := u.OfNumber; vt != nil && vt.Placeholder.Valid() {
		return &vt.Placeholder.Value
	} else if vt := u.OfDropdown; vt != nil && vt.Placeholder.Valid() {
		return &vt.Placeholder.Value
	} else if vt := u.OfRadio; vt != nil && vt.Placeholder.Valid() {
		return &vt.Placeholder.Value
	} else if vt := u.OfDatepicker; vt != nil && vt.Placeholder.Valid() {
		return &vt.Placeholder.Value
	} else if vt := u.OfFile; vt != nil && vt.Placeholder.Valid() {
		return &vt.Placeholder.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's DependentFields property, if
// present.
func (u DependentFieldDependentFieldUnionParam) GetDependentFields() []DependentFieldParam {
	if vt := u.OfEmail; vt != nil {
		return vt.DependentFields
	} else if vt := u.OfPhone; vt != nil {
		return vt.DependentFields
	} else if vt := u.OfMobilePhone; vt != nil {
		return vt.DependentFields
	} else if vt := u.OfSingleLineText; vt != nil {
		return vt.DependentFields
	} else if vt := u.OfMultiLineText; vt != nil {
		return vt.DependentFields
	} else if vt := u.OfNumber; vt != nil {
		return vt.DependentFields
	} else if vt := u.OfSingleCheckbox; vt != nil {
		return vt.DependentFields
	} else if vt := u.OfMultipleCheckboxes; vt != nil {
		return vt.DependentFields
	} else if vt := u.OfDropdown; vt != nil {
		return vt.DependentFields
	} else if vt := u.OfRadio; vt != nil {
		return vt.DependentFields
	} else if vt := u.OfDatepicker; vt != nil {
		return vt.DependentFields
	} else if vt := u.OfFile; vt != nil {
		return vt.DependentFields
	} else if vt := u.OfPaymentLinkRadio; vt != nil {
		return vt.DependentFields
	}
	return nil
}

// Returns a subunion which exports methods to access subproperties
//
// Or use AsAny() to get the underlying value
func (u DependentFieldDependentFieldUnionParam) GetValidation() (res dependentFieldDependentFieldUnionParamValidation) {
	if vt := u.OfEmail; vt != nil {
		res.any = &vt.Validation
	} else if vt := u.OfPhone; vt != nil {
		res.any = &vt.Validation
	} else if vt := u.OfMobilePhone; vt != nil {
		res.any = &vt.Validation
	} else if vt := u.OfNumber; vt != nil {
		res.any = &vt.Validation
	}
	return
}

// Can have the runtime types [*EmailFieldValidationParam],
// [*PhoneFieldValidationParam], [*NumberFieldValidationParam]
type dependentFieldDependentFieldUnionParamValidation struct{ any }

// Use the following switch statement to get the type of the union:
//
//	switch u.AsAny().(type) {
//	case *marketing.EmailFieldValidationParam:
//	case *marketing.PhoneFieldValidationParam:
//	case *marketing.NumberFieldValidationParam:
//	default:
//	    fmt.Errorf("not present")
//	}
func (u dependentFieldDependentFieldUnionParamValidation) AsAny() any { return u.any }

// Returns a pointer to the underlying variant's property, if present.
func (u dependentFieldDependentFieldUnionParamValidation) GetBlockedEmailDomains() []string {
	switch vt := u.any.(type) {
	case *EmailFieldValidationParam:
		return vt.BlockedEmailDomains
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u dependentFieldDependentFieldUnionParamValidation) GetUseDefaultBlockList() *bool {
	switch vt := u.any.(type) {
	case *EmailFieldValidationParam:
		return &vt.UseDefaultBlockList
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u dependentFieldDependentFieldUnionParamValidation) GetMaxAllowedDigits() *int64 {
	switch vt := u.any.(type) {
	case *PhoneFieldValidationParam:
		return (*int64)(&vt.MaxAllowedDigits)
	case *NumberFieldValidationParam:
		return (*int64)(&vt.MaxAllowedDigits)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u dependentFieldDependentFieldUnionParamValidation) GetMinAllowedDigits() *int64 {
	switch vt := u.any.(type) {
	case *PhoneFieldValidationParam:
		return (*int64)(&vt.MinAllowedDigits)
	case *NumberFieldValidationParam:
		return (*int64)(&vt.MinAllowedDigits)
	}
	return nil
}

// Returns a pointer to the underlying variant's DefaultValues property, if
// present.
func (u DependentFieldDependentFieldUnionParam) GetDefaultValues() []string {
	if vt := u.OfMultipleCheckboxes; vt != nil {
		return vt.DefaultValues
	} else if vt := u.OfDropdown; vt != nil {
		return vt.DefaultValues
	} else if vt := u.OfRadio; vt != nil {
		return vt.DefaultValues
	} else if vt := u.OfPaymentLinkRadio; vt != nil {
		return vt.DefaultValues
	}
	return nil
}

// Returns a pointer to the underlying variant's Options property, if present.
func (u DependentFieldDependentFieldUnionParam) GetOptions() []EnumeratedFieldOptionParam {
	if vt := u.OfMultipleCheckboxes; vt != nil {
		return vt.Options
	} else if vt := u.OfDropdown; vt != nil {
		return vt.Options
	} else if vt := u.OfRadio; vt != nil {
		return vt.Options
	} else if vt := u.OfPaymentLinkRadio; vt != nil {
		return vt.Options
	}
	return nil
}

// A condition based on customer input
type DependentFieldFilter struct {
	// Any of "eq", "neq", "contains", "doesnt_contain", "str_starts_with",
	// "str_ends_with", "lt", "lte", "gt", "gte", "between", "not_between",
	// "within_time_reverse", "within_time", "set_any", "set_not_any", "set_all",
	// "set_not_all", "set_eq", "set_neq", "is_not_empty".
	Operator   DependentFieldFilterOperator `json:"operator,required"`
	RangeEnd   string                       `json:"rangeEnd,required"`
	RangeStart string                       `json:"rangeStart,required"`
	Value      string                       `json:"value,required"`
	Values     []string                     `json:"values,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Operator    respjson.Field
		RangeEnd    respjson.Field
		RangeStart  respjson.Field
		Value       respjson.Field
		Values      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DependentFieldFilter) RawJSON() string { return r.JSON.raw }
func (r *DependentFieldFilter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this DependentFieldFilter to a DependentFieldFilterParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// DependentFieldFilterParam.Overrides()
func (r DependentFieldFilter) ToParam() DependentFieldFilterParam {
	return param.Override[DependentFieldFilterParam](json.RawMessage(r.RawJSON()))
}

type DependentFieldFilterOperator string

const (
	DependentFieldFilterOperatorEq                DependentFieldFilterOperator = "eq"
	DependentFieldFilterOperatorNeq               DependentFieldFilterOperator = "neq"
	DependentFieldFilterOperatorContains          DependentFieldFilterOperator = "contains"
	DependentFieldFilterOperatorDoesntContain     DependentFieldFilterOperator = "doesnt_contain"
	DependentFieldFilterOperatorStrStartsWith     DependentFieldFilterOperator = "str_starts_with"
	DependentFieldFilterOperatorStrEndsWith       DependentFieldFilterOperator = "str_ends_with"
	DependentFieldFilterOperatorLt                DependentFieldFilterOperator = "lt"
	DependentFieldFilterOperatorLte               DependentFieldFilterOperator = "lte"
	DependentFieldFilterOperatorGt                DependentFieldFilterOperator = "gt"
	DependentFieldFilterOperatorGte               DependentFieldFilterOperator = "gte"
	DependentFieldFilterOperatorBetween           DependentFieldFilterOperator = "between"
	DependentFieldFilterOperatorNotBetween        DependentFieldFilterOperator = "not_between"
	DependentFieldFilterOperatorWithinTimeReverse DependentFieldFilterOperator = "within_time_reverse"
	DependentFieldFilterOperatorWithinTime        DependentFieldFilterOperator = "within_time"
	DependentFieldFilterOperatorSetAny            DependentFieldFilterOperator = "set_any"
	DependentFieldFilterOperatorSetNotAny         DependentFieldFilterOperator = "set_not_any"
	DependentFieldFilterOperatorSetAll            DependentFieldFilterOperator = "set_all"
	DependentFieldFilterOperatorSetNotAll         DependentFieldFilterOperator = "set_not_all"
	DependentFieldFilterOperatorSetEq             DependentFieldFilterOperator = "set_eq"
	DependentFieldFilterOperatorSetNeq            DependentFieldFilterOperator = "set_neq"
	DependentFieldFilterOperatorIsNotEmpty        DependentFieldFilterOperator = "is_not_empty"
)

// A condition based on customer input
//
// The properties Operator, RangeEnd, RangeStart, Value, Values are required.
type DependentFieldFilterParam struct {
	// Any of "eq", "neq", "contains", "doesnt_contain", "str_starts_with",
	// "str_ends_with", "lt", "lte", "gt", "gte", "between", "not_between",
	// "within_time_reverse", "within_time", "set_any", "set_not_any", "set_all",
	// "set_not_all", "set_eq", "set_neq", "is_not_empty".
	Operator   DependentFieldFilterOperator `json:"operator,omitzero,required"`
	RangeEnd   string                       `json:"rangeEnd,required"`
	RangeStart string                       `json:"rangeStart,required"`
	Value      string                       `json:"value,required"`
	Values     []string                     `json:"values,omitzero,required"`
	paramObj
}

func (r DependentFieldFilterParam) MarshalJSON() (data []byte, err error) {
	type shadow DependentFieldFilterParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DependentFieldFilterParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A field consisting of a drop down with multiple choices.
type DropdownField struct {
	// The values selected by default. Those values will be submitted unless the
	// customer modifies them.
	DefaultValues []string `json:"defaultValues,required"`
	// A list of other fields to make visible based on the value filled in for this
	// field.
	DependentFields []DependentField `json:"dependentFields,required"`
	// Determines how the field will be displayed and validated.
	//
	// Any of "dropdown".
	FieldType DropdownFieldFieldType `json:"fieldType,required"`
	// Whether a field should be hidden or not. Hidden fields won't appear on the form,
	// but can be used to pass a value to a property without requiring the customer to
	// fill it in.
	Hidden bool `json:"hidden,required"`
	// The main label for the form field.
	Label string `json:"label,required"`
	// The identifier of the field. In combination with the object type ID, it must be
	// unique.
	Name string `json:"name,required"`
	// A unique ID for this field's CRM object type. For example a CONTACT field will
	// have the object type ID 0-1.
	ObjectTypeID string `json:"objectTypeId,required"`
	// The list of available choices for this field.
	Options []EnumeratedFieldOption `json:"options,required"`
	// Whether a value for this field is required when submitting the form.
	Required bool `json:"required,required"`
	// Additional text helping the customer to complete the field.
	Description string `json:"description"`
	// The prompt text showing when the field isn't filled in.
	Placeholder string `json:"placeholder"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DefaultValues   respjson.Field
		DependentFields respjson.Field
		FieldType       respjson.Field
		Hidden          respjson.Field
		Label           respjson.Field
		Name            respjson.Field
		ObjectTypeID    respjson.Field
		Options         respjson.Field
		Required        respjson.Field
		Description     respjson.Field
		Placeholder     respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DropdownField) RawJSON() string { return r.JSON.raw }
func (r *DropdownField) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this DropdownField to a DropdownFieldParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// DropdownFieldParam.Overrides()
func (r DropdownField) ToParam() DropdownFieldParam {
	return param.Override[DropdownFieldParam](json.RawMessage(r.RawJSON()))
}

// Determines how the field will be displayed and validated.
type DropdownFieldFieldType string

const (
	DropdownFieldFieldTypeDropdown DropdownFieldFieldType = "dropdown"
)

// A field consisting of a drop down with multiple choices.
//
// The properties DefaultValues, DependentFields, FieldType, Hidden, Label, Name,
// ObjectTypeID, Options, Required are required.
type DropdownFieldParam struct {
	// The values selected by default. Those values will be submitted unless the
	// customer modifies them.
	DefaultValues []string `json:"defaultValues,omitzero,required"`
	// A list of other fields to make visible based on the value filled in for this
	// field.
	DependentFields []DependentFieldParam `json:"dependentFields,omitzero,required"`
	// Determines how the field will be displayed and validated.
	//
	// Any of "dropdown".
	FieldType DropdownFieldFieldType `json:"fieldType,omitzero,required"`
	// Whether a field should be hidden or not. Hidden fields won't appear on the form,
	// but can be used to pass a value to a property without requiring the customer to
	// fill it in.
	Hidden bool `json:"hidden,required"`
	// The main label for the form field.
	Label string `json:"label,required"`
	// The identifier of the field. In combination with the object type ID, it must be
	// unique.
	Name string `json:"name,required"`
	// A unique ID for this field's CRM object type. For example a CONTACT field will
	// have the object type ID 0-1.
	ObjectTypeID string `json:"objectTypeId,required"`
	// The list of available choices for this field.
	Options []EnumeratedFieldOptionParam `json:"options,omitzero,required"`
	// Whether a value for this field is required when submitting the form.
	Required bool `json:"required,required"`
	// Additional text helping the customer to complete the field.
	Description param.Opt[string] `json:"description,omitzero"`
	// The prompt text showing when the field isn't filled in.
	Placeholder param.Opt[string] `json:"placeholder,omitzero"`
	paramObj
}

func (r DropdownFieldParam) MarshalJSON() (data []byte, err error) {
	type shadow DropdownFieldParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DropdownFieldParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A form field used for collecting an email address.
type EmailField struct {
	// A list of other fields to make visible based on the value filled in for this
	// field.
	DependentFields []DependentField `json:"dependentFields,required"`
	// Determines how the field will be displayed and validated.
	//
	// Any of "email".
	FieldType EmailFieldFieldType `json:"fieldType,required"`
	// Whether a field should be hidden or not. Hidden fields won't appear on the form,
	// but can be used to pass a value to a property without requiring the customer to
	// fill it in.
	Hidden bool `json:"hidden,required"`
	// The main label for the form field.
	Label string `json:"label,required"`
	// The identifier of the field. In combination with the object type ID, it must be
	// unique.
	Name string `json:"name,required"`
	// A unique ID for this field's CRM object type. For example a CONTACT field will
	// have the object type ID 0-1.
	ObjectTypeID string `json:"objectTypeId,required"`
	// Whether a value for this field is required when submitting the form.
	Required bool `json:"required,required"`
	// Describes how an email address should be validated.
	Validation EmailFieldValidation `json:"validation,required"`
	// The value filled in by default. This value will be submitted unless the customer
	// modifies it.
	DefaultValue string `json:"defaultValue"`
	// Additional text helping the customer to complete the field.
	Description string `json:"description"`
	// The prompt text showing when the field isn't filled in.
	Placeholder string `json:"placeholder"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DependentFields respjson.Field
		FieldType       respjson.Field
		Hidden          respjson.Field
		Label           respjson.Field
		Name            respjson.Field
		ObjectTypeID    respjson.Field
		Required        respjson.Field
		Validation      respjson.Field
		DefaultValue    respjson.Field
		Description     respjson.Field
		Placeholder     respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmailField) RawJSON() string { return r.JSON.raw }
func (r *EmailField) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this EmailField to a EmailFieldParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// EmailFieldParam.Overrides()
func (r EmailField) ToParam() EmailFieldParam {
	return param.Override[EmailFieldParam](json.RawMessage(r.RawJSON()))
}

// Determines how the field will be displayed and validated.
type EmailFieldFieldType string

const (
	EmailFieldFieldTypeEmail EmailFieldFieldType = "email"
)

// A form field used for collecting an email address.
//
// The properties DependentFields, FieldType, Hidden, Label, Name, ObjectTypeID,
// Required, Validation are required.
type EmailFieldParam struct {
	// A list of other fields to make visible based on the value filled in for this
	// field.
	DependentFields []DependentFieldParam `json:"dependentFields,omitzero,required"`
	// Determines how the field will be displayed and validated.
	//
	// Any of "email".
	FieldType EmailFieldFieldType `json:"fieldType,omitzero,required"`
	// Whether a field should be hidden or not. Hidden fields won't appear on the form,
	// but can be used to pass a value to a property without requiring the customer to
	// fill it in.
	Hidden bool `json:"hidden,required"`
	// The main label for the form field.
	Label string `json:"label,required"`
	// The identifier of the field. In combination with the object type ID, it must be
	// unique.
	Name string `json:"name,required"`
	// A unique ID for this field's CRM object type. For example a CONTACT field will
	// have the object type ID 0-1.
	ObjectTypeID string `json:"objectTypeId,required"`
	// Whether a value for this field is required when submitting the form.
	Required bool `json:"required,required"`
	// Describes how an email address should be validated.
	Validation EmailFieldValidationParam `json:"validation,omitzero,required"`
	// The value filled in by default. This value will be submitted unless the customer
	// modifies it.
	DefaultValue param.Opt[string] `json:"defaultValue,omitzero"`
	// Additional text helping the customer to complete the field.
	Description param.Opt[string] `json:"description,omitzero"`
	// The prompt text showing when the field isn't filled in.
	Placeholder param.Opt[string] `json:"placeholder,omitzero"`
	paramObj
}

func (r EmailFieldParam) MarshalJSON() (data []byte, err error) {
	type shadow EmailFieldParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EmailFieldParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Describes how an email address should be validated.
type EmailFieldValidation struct {
	// A list of email domains to block.
	BlockedEmailDomains []string `json:"blockedEmailDomains,required"`
	// Whether to block the free email providers.
	UseDefaultBlockList bool `json:"useDefaultBlockList,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BlockedEmailDomains respjson.Field
		UseDefaultBlockList respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmailFieldValidation) RawJSON() string { return r.JSON.raw }
func (r *EmailFieldValidation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this EmailFieldValidation to a EmailFieldValidationParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// EmailFieldValidationParam.Overrides()
func (r EmailFieldValidation) ToParam() EmailFieldValidationParam {
	return param.Override[EmailFieldValidationParam](json.RawMessage(r.RawJSON()))
}

// Describes how an email address should be validated.
//
// The properties BlockedEmailDomains, UseDefaultBlockList are required.
type EmailFieldValidationParam struct {
	// A list of email domains to block.
	BlockedEmailDomains []string `json:"blockedEmailDomains,omitzero,required"`
	// Whether to block the free email providers.
	UseDefaultBlockList bool `json:"useDefaultBlockList,required"`
	paramObj
}

func (r EmailFieldValidationParam) MarshalJSON() (data []byte, err error) {
	type shadow EmailFieldValidationParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EmailFieldValidationParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EnumeratedFieldOption struct {
	// The order the choices will be displayed in.
	DisplayOrder int64 `json:"displayOrder,required"`
	// The visible label for this choice.
	Label string `json:"label,required"`
	// The value which will be submitted if this choice is selected.
	Value       string `json:"value,required"`
	Description string `json:"description"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DisplayOrder respjson.Field
		Label        respjson.Field
		Value        respjson.Field
		Description  respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EnumeratedFieldOption) RawJSON() string { return r.JSON.raw }
func (r *EnumeratedFieldOption) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this EnumeratedFieldOption to a EnumeratedFieldOptionParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// EnumeratedFieldOptionParam.Overrides()
func (r EnumeratedFieldOption) ToParam() EnumeratedFieldOptionParam {
	return param.Override[EnumeratedFieldOptionParam](json.RawMessage(r.RawJSON()))
}

// The properties DisplayOrder, Label, Value are required.
type EnumeratedFieldOptionParam struct {
	// The order the choices will be displayed in.
	DisplayOrder int64 `json:"displayOrder,required"`
	// The visible label for this choice.
	Label string `json:"label,required"`
	// The value which will be submitted if this choice is selected.
	Value       string            `json:"value,required"`
	Description param.Opt[string] `json:"description,omitzero"`
	paramObj
}

func (r EnumeratedFieldOptionParam) MarshalJSON() (data []byte, err error) {
	type shadow EnumeratedFieldOptionParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EnumeratedFieldOptionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A collection of up to three form fields usually displayed in a row.
type FieldGroup struct {
	// The form fields included in the group
	Fields []FieldGroupFieldUnion `json:"fields,required"`
	// Any of "default_group", "progressive", "queued".
	GroupType FieldGroupGroupType `json:"groupType,required"`
	// The type of rich text included. The default value is text.
	//
	// Any of "text", "image".
	RichTextType FieldGroupRichTextType `json:"richTextType,required"`
	// A block of rich text or an image. Those can be used to add extra information for
	// the customers filling in the form. If the field group includes fields, the rich
	// text will be displayed before the fields.
	RichText string `json:"richText"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Fields       respjson.Field
		GroupType    respjson.Field
		RichTextType respjson.Field
		RichText     respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FieldGroup) RawJSON() string { return r.JSON.raw }
func (r *FieldGroup) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this FieldGroup to a FieldGroupParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// FieldGroupParam.Overrides()
func (r FieldGroup) ToParam() FieldGroupParam {
	return param.Override[FieldGroupParam](json.RawMessage(r.RawJSON()))
}

// FieldGroupFieldUnion contains all possible properties and values from
// [EmailField], [PhoneField], [MobilePhoneField], [SingleLineTextField],
// [MultiLineTextField], [NumberField], [SingleCheckboxField],
// [MultipleCheckboxesField], [DropdownField], [RadioField], [DatepickerField],
// [FileField], [PaymentLinkRadioField].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type FieldGroupFieldUnion struct {
	DependentFields []DependentField `json:"dependentFields"`
	FieldType       string           `json:"fieldType"`
	Hidden          bool             `json:"hidden"`
	Label           string           `json:"label"`
	Name            string           `json:"name"`
	ObjectTypeID    string           `json:"objectTypeId"`
	Required        bool             `json:"required"`
	// This field is a union of [EmailFieldValidation], [PhoneFieldValidation],
	// [NumberFieldValidation]
	Validation   FieldGroupFieldUnionValidation `json:"validation"`
	DefaultValue string                         `json:"defaultValue"`
	Description  string                         `json:"description"`
	Placeholder  string                         `json:"placeholder"`
	// This field is from variant [PhoneField].
	UseCountryCodeSelect bool                    `json:"useCountryCodeSelect"`
	DefaultValues        []string                `json:"defaultValues"`
	Options              []EnumeratedFieldOption `json:"options"`
	// This field is from variant [FileField].
	AllowMultipleFiles bool `json:"allowMultipleFiles"`
	JSON               struct {
		DependentFields      respjson.Field
		FieldType            respjson.Field
		Hidden               respjson.Field
		Label                respjson.Field
		Name                 respjson.Field
		ObjectTypeID         respjson.Field
		Required             respjson.Field
		Validation           respjson.Field
		DefaultValue         respjson.Field
		Description          respjson.Field
		Placeholder          respjson.Field
		UseCountryCodeSelect respjson.Field
		DefaultValues        respjson.Field
		Options              respjson.Field
		AllowMultipleFiles   respjson.Field
		raw                  string
	} `json:"-"`
}

func (u FieldGroupFieldUnion) AsEmail() (v EmailField) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FieldGroupFieldUnion) AsPhone() (v PhoneField) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FieldGroupFieldUnion) AsMobilePhone() (v MobilePhoneField) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FieldGroupFieldUnion) AsSingleLineText() (v SingleLineTextField) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FieldGroupFieldUnion) AsMultiLineText() (v MultiLineTextField) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FieldGroupFieldUnion) AsNumber() (v NumberField) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FieldGroupFieldUnion) AsSingleCheckbox() (v SingleCheckboxField) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FieldGroupFieldUnion) AsMultipleCheckboxes() (v MultipleCheckboxesField) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FieldGroupFieldUnion) AsDropdown() (v DropdownField) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FieldGroupFieldUnion) AsRadio() (v RadioField) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FieldGroupFieldUnion) AsDatepicker() (v DatepickerField) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FieldGroupFieldUnion) AsFile() (v FileField) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FieldGroupFieldUnion) AsPaymentLinkRadio() (v PaymentLinkRadioField) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u FieldGroupFieldUnion) RawJSON() string { return u.JSON.raw }

func (r *FieldGroupFieldUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// FieldGroupFieldUnionValidation is an implicit subunion of
// [FieldGroupFieldUnion]. FieldGroupFieldUnionValidation provides convenient
// access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [FieldGroupFieldUnion].
type FieldGroupFieldUnionValidation struct {
	// This field is from variant [EmailFieldValidation].
	BlockedEmailDomains []string `json:"blockedEmailDomains"`
	// This field is from variant [EmailFieldValidation].
	UseDefaultBlockList bool  `json:"useDefaultBlockList"`
	MaxAllowedDigits    int64 `json:"maxAllowedDigits"`
	MinAllowedDigits    int64 `json:"minAllowedDigits"`
	JSON                struct {
		BlockedEmailDomains respjson.Field
		UseDefaultBlockList respjson.Field
		MaxAllowedDigits    respjson.Field
		MinAllowedDigits    respjson.Field
		raw                 string
	} `json:"-"`
}

func (r *FieldGroupFieldUnionValidation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FieldGroupGroupType string

const (
	FieldGroupGroupTypeDefaultGroup FieldGroupGroupType = "default_group"
	FieldGroupGroupTypeProgressive  FieldGroupGroupType = "progressive"
	FieldGroupGroupTypeQueued       FieldGroupGroupType = "queued"
)

// The type of rich text included. The default value is text.
type FieldGroupRichTextType string

const (
	FieldGroupRichTextTypeText  FieldGroupRichTextType = "text"
	FieldGroupRichTextTypeImage FieldGroupRichTextType = "image"
)

// A collection of up to three form fields usually displayed in a row.
//
// The properties Fields, GroupType, RichTextType are required.
type FieldGroupParam struct {
	// The form fields included in the group
	Fields []FieldGroupFieldUnionParam `json:"fields,omitzero,required"`
	// Any of "default_group", "progressive", "queued".
	GroupType FieldGroupGroupType `json:"groupType,omitzero,required"`
	// The type of rich text included. The default value is text.
	//
	// Any of "text", "image".
	RichTextType FieldGroupRichTextType `json:"richTextType,omitzero,required"`
	// A block of rich text or an image. Those can be used to add extra information for
	// the customers filling in the form. If the field group includes fields, the rich
	// text will be displayed before the fields.
	RichText param.Opt[string] `json:"richText,omitzero"`
	paramObj
}

func (r FieldGroupParam) MarshalJSON() (data []byte, err error) {
	type shadow FieldGroupParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FieldGroupParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type FieldGroupFieldUnionParam struct {
	OfEmail              *EmailFieldParam              `json:",omitzero,inline"`
	OfPhone              *PhoneFieldParam              `json:",omitzero,inline"`
	OfMobilePhone        *MobilePhoneFieldParam        `json:",omitzero,inline"`
	OfSingleLineText     *SingleLineTextFieldParam     `json:",omitzero,inline"`
	OfMultiLineText      *MultiLineTextFieldParam      `json:",omitzero,inline"`
	OfNumber             *NumberFieldParam             `json:",omitzero,inline"`
	OfSingleCheckbox     *SingleCheckboxFieldParam     `json:",omitzero,inline"`
	OfMultipleCheckboxes *MultipleCheckboxesFieldParam `json:",omitzero,inline"`
	OfDropdown           *DropdownFieldParam           `json:",omitzero,inline"`
	OfRadio              *RadioFieldParam              `json:",omitzero,inline"`
	OfDatepicker         *DatepickerFieldParam         `json:",omitzero,inline"`
	OfFile               *FileFieldParam               `json:",omitzero,inline"`
	OfPaymentLinkRadio   *PaymentLinkRadioFieldParam   `json:",omitzero,inline"`
	paramUnion
}

func (u FieldGroupFieldUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfEmail,
		u.OfPhone,
		u.OfMobilePhone,
		u.OfSingleLineText,
		u.OfMultiLineText,
		u.OfNumber,
		u.OfSingleCheckbox,
		u.OfMultipleCheckboxes,
		u.OfDropdown,
		u.OfRadio,
		u.OfDatepicker,
		u.OfFile,
		u.OfPaymentLinkRadio)
}
func (u *FieldGroupFieldUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *FieldGroupFieldUnionParam) asAny() any {
	if !param.IsOmitted(u.OfEmail) {
		return u.OfEmail
	} else if !param.IsOmitted(u.OfPhone) {
		return u.OfPhone
	} else if !param.IsOmitted(u.OfMobilePhone) {
		return u.OfMobilePhone
	} else if !param.IsOmitted(u.OfSingleLineText) {
		return u.OfSingleLineText
	} else if !param.IsOmitted(u.OfMultiLineText) {
		return u.OfMultiLineText
	} else if !param.IsOmitted(u.OfNumber) {
		return u.OfNumber
	} else if !param.IsOmitted(u.OfSingleCheckbox) {
		return u.OfSingleCheckbox
	} else if !param.IsOmitted(u.OfMultipleCheckboxes) {
		return u.OfMultipleCheckboxes
	} else if !param.IsOmitted(u.OfDropdown) {
		return u.OfDropdown
	} else if !param.IsOmitted(u.OfRadio) {
		return u.OfRadio
	} else if !param.IsOmitted(u.OfDatepicker) {
		return u.OfDatepicker
	} else if !param.IsOmitted(u.OfFile) {
		return u.OfFile
	} else if !param.IsOmitted(u.OfPaymentLinkRadio) {
		return u.OfPaymentLinkRadio
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u FieldGroupFieldUnionParam) GetUseCountryCodeSelect() *bool {
	if vt := u.OfPhone; vt != nil {
		return &vt.UseCountryCodeSelect
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u FieldGroupFieldUnionParam) GetAllowMultipleFiles() *bool {
	if vt := u.OfFile; vt != nil {
		return &vt.AllowMultipleFiles
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u FieldGroupFieldUnionParam) GetFieldType() *string {
	if vt := u.OfEmail; vt != nil {
		return (*string)(&vt.FieldType)
	} else if vt := u.OfPhone; vt != nil {
		return (*string)(&vt.FieldType)
	} else if vt := u.OfMobilePhone; vt != nil {
		return (*string)(&vt.FieldType)
	} else if vt := u.OfSingleLineText; vt != nil {
		return (*string)(&vt.FieldType)
	} else if vt := u.OfMultiLineText; vt != nil {
		return (*string)(&vt.FieldType)
	} else if vt := u.OfNumber; vt != nil {
		return (*string)(&vt.FieldType)
	} else if vt := u.OfSingleCheckbox; vt != nil {
		return (*string)(&vt.FieldType)
	} else if vt := u.OfMultipleCheckboxes; vt != nil {
		return (*string)(&vt.FieldType)
	} else if vt := u.OfDropdown; vt != nil {
		return (*string)(&vt.FieldType)
	} else if vt := u.OfRadio; vt != nil {
		return (*string)(&vt.FieldType)
	} else if vt := u.OfDatepicker; vt != nil {
		return (*string)(&vt.FieldType)
	} else if vt := u.OfFile; vt != nil {
		return (*string)(&vt.FieldType)
	} else if vt := u.OfPaymentLinkRadio; vt != nil {
		return (*string)(&vt.FieldType)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u FieldGroupFieldUnionParam) GetHidden() *bool {
	if vt := u.OfEmail; vt != nil {
		return (*bool)(&vt.Hidden)
	} else if vt := u.OfPhone; vt != nil {
		return (*bool)(&vt.Hidden)
	} else if vt := u.OfMobilePhone; vt != nil {
		return (*bool)(&vt.Hidden)
	} else if vt := u.OfSingleLineText; vt != nil {
		return (*bool)(&vt.Hidden)
	} else if vt := u.OfMultiLineText; vt != nil {
		return (*bool)(&vt.Hidden)
	} else if vt := u.OfNumber; vt != nil {
		return (*bool)(&vt.Hidden)
	} else if vt := u.OfSingleCheckbox; vt != nil {
		return (*bool)(&vt.Hidden)
	} else if vt := u.OfMultipleCheckboxes; vt != nil {
		return (*bool)(&vt.Hidden)
	} else if vt := u.OfDropdown; vt != nil {
		return (*bool)(&vt.Hidden)
	} else if vt := u.OfRadio; vt != nil {
		return (*bool)(&vt.Hidden)
	} else if vt := u.OfDatepicker; vt != nil {
		return (*bool)(&vt.Hidden)
	} else if vt := u.OfFile; vt != nil {
		return (*bool)(&vt.Hidden)
	} else if vt := u.OfPaymentLinkRadio; vt != nil {
		return (*bool)(&vt.Hidden)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u FieldGroupFieldUnionParam) GetLabel() *string {
	if vt := u.OfEmail; vt != nil {
		return (*string)(&vt.Label)
	} else if vt := u.OfPhone; vt != nil {
		return (*string)(&vt.Label)
	} else if vt := u.OfMobilePhone; vt != nil {
		return (*string)(&vt.Label)
	} else if vt := u.OfSingleLineText; vt != nil {
		return (*string)(&vt.Label)
	} else if vt := u.OfMultiLineText; vt != nil {
		return (*string)(&vt.Label)
	} else if vt := u.OfNumber; vt != nil {
		return (*string)(&vt.Label)
	} else if vt := u.OfSingleCheckbox; vt != nil {
		return (*string)(&vt.Label)
	} else if vt := u.OfMultipleCheckboxes; vt != nil {
		return (*string)(&vt.Label)
	} else if vt := u.OfDropdown; vt != nil {
		return (*string)(&vt.Label)
	} else if vt := u.OfRadio; vt != nil {
		return (*string)(&vt.Label)
	} else if vt := u.OfDatepicker; vt != nil {
		return (*string)(&vt.Label)
	} else if vt := u.OfFile; vt != nil {
		return (*string)(&vt.Label)
	} else if vt := u.OfPaymentLinkRadio; vt != nil {
		return (*string)(&vt.Label)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u FieldGroupFieldUnionParam) GetName() *string {
	if vt := u.OfEmail; vt != nil {
		return (*string)(&vt.Name)
	} else if vt := u.OfPhone; vt != nil {
		return (*string)(&vt.Name)
	} else if vt := u.OfMobilePhone; vt != nil {
		return (*string)(&vt.Name)
	} else if vt := u.OfSingleLineText; vt != nil {
		return (*string)(&vt.Name)
	} else if vt := u.OfMultiLineText; vt != nil {
		return (*string)(&vt.Name)
	} else if vt := u.OfNumber; vt != nil {
		return (*string)(&vt.Name)
	} else if vt := u.OfSingleCheckbox; vt != nil {
		return (*string)(&vt.Name)
	} else if vt := u.OfMultipleCheckboxes; vt != nil {
		return (*string)(&vt.Name)
	} else if vt := u.OfDropdown; vt != nil {
		return (*string)(&vt.Name)
	} else if vt := u.OfRadio; vt != nil {
		return (*string)(&vt.Name)
	} else if vt := u.OfDatepicker; vt != nil {
		return (*string)(&vt.Name)
	} else if vt := u.OfFile; vt != nil {
		return (*string)(&vt.Name)
	} else if vt := u.OfPaymentLinkRadio; vt != nil {
		return (*string)(&vt.Name)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u FieldGroupFieldUnionParam) GetObjectTypeID() *string {
	if vt := u.OfEmail; vt != nil {
		return (*string)(&vt.ObjectTypeID)
	} else if vt := u.OfPhone; vt != nil {
		return (*string)(&vt.ObjectTypeID)
	} else if vt := u.OfMobilePhone; vt != nil {
		return (*string)(&vt.ObjectTypeID)
	} else if vt := u.OfSingleLineText; vt != nil {
		return (*string)(&vt.ObjectTypeID)
	} else if vt := u.OfMultiLineText; vt != nil {
		return (*string)(&vt.ObjectTypeID)
	} else if vt := u.OfNumber; vt != nil {
		return (*string)(&vt.ObjectTypeID)
	} else if vt := u.OfSingleCheckbox; vt != nil {
		return (*string)(&vt.ObjectTypeID)
	} else if vt := u.OfMultipleCheckboxes; vt != nil {
		return (*string)(&vt.ObjectTypeID)
	} else if vt := u.OfDropdown; vt != nil {
		return (*string)(&vt.ObjectTypeID)
	} else if vt := u.OfRadio; vt != nil {
		return (*string)(&vt.ObjectTypeID)
	} else if vt := u.OfDatepicker; vt != nil {
		return (*string)(&vt.ObjectTypeID)
	} else if vt := u.OfFile; vt != nil {
		return (*string)(&vt.ObjectTypeID)
	} else if vt := u.OfPaymentLinkRadio; vt != nil {
		return (*string)(&vt.ObjectTypeID)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u FieldGroupFieldUnionParam) GetRequired() *bool {
	if vt := u.OfEmail; vt != nil {
		return (*bool)(&vt.Required)
	} else if vt := u.OfPhone; vt != nil {
		return (*bool)(&vt.Required)
	} else if vt := u.OfMobilePhone; vt != nil {
		return (*bool)(&vt.Required)
	} else if vt := u.OfSingleLineText; vt != nil {
		return (*bool)(&vt.Required)
	} else if vt := u.OfMultiLineText; vt != nil {
		return (*bool)(&vt.Required)
	} else if vt := u.OfNumber; vt != nil {
		return (*bool)(&vt.Required)
	} else if vt := u.OfSingleCheckbox; vt != nil {
		return (*bool)(&vt.Required)
	} else if vt := u.OfMultipleCheckboxes; vt != nil {
		return (*bool)(&vt.Required)
	} else if vt := u.OfDropdown; vt != nil {
		return (*bool)(&vt.Required)
	} else if vt := u.OfRadio; vt != nil {
		return (*bool)(&vt.Required)
	} else if vt := u.OfDatepicker; vt != nil {
		return (*bool)(&vt.Required)
	} else if vt := u.OfFile; vt != nil {
		return (*bool)(&vt.Required)
	} else if vt := u.OfPaymentLinkRadio; vt != nil {
		return (*bool)(&vt.Required)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u FieldGroupFieldUnionParam) GetDefaultValue() *string {
	if vt := u.OfEmail; vt != nil && vt.DefaultValue.Valid() {
		return &vt.DefaultValue.Value
	} else if vt := u.OfPhone; vt != nil && vt.DefaultValue.Valid() {
		return &vt.DefaultValue.Value
	} else if vt := u.OfMobilePhone; vt != nil && vt.DefaultValue.Valid() {
		return &vt.DefaultValue.Value
	} else if vt := u.OfSingleLineText; vt != nil && vt.DefaultValue.Valid() {
		return &vt.DefaultValue.Value
	} else if vt := u.OfMultiLineText; vt != nil && vt.DefaultValue.Valid() {
		return &vt.DefaultValue.Value
	} else if vt := u.OfNumber; vt != nil && vt.DefaultValue.Valid() {
		return &vt.DefaultValue.Value
	} else if vt := u.OfSingleCheckbox; vt != nil && vt.DefaultValue.Valid() {
		return &vt.DefaultValue.Value
	} else if vt := u.OfDatepicker; vt != nil && vt.DefaultValue.Valid() {
		return &vt.DefaultValue.Value
	} else if vt := u.OfFile; vt != nil && vt.DefaultValue.Valid() {
		return &vt.DefaultValue.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u FieldGroupFieldUnionParam) GetDescription() *string {
	if vt := u.OfEmail; vt != nil && vt.Description.Valid() {
		return &vt.Description.Value
	} else if vt := u.OfPhone; vt != nil && vt.Description.Valid() {
		return &vt.Description.Value
	} else if vt := u.OfMobilePhone; vt != nil && vt.Description.Valid() {
		return &vt.Description.Value
	} else if vt := u.OfSingleLineText; vt != nil && vt.Description.Valid() {
		return &vt.Description.Value
	} else if vt := u.OfMultiLineText; vt != nil && vt.Description.Valid() {
		return &vt.Description.Value
	} else if vt := u.OfNumber; vt != nil && vt.Description.Valid() {
		return &vt.Description.Value
	} else if vt := u.OfSingleCheckbox; vt != nil && vt.Description.Valid() {
		return &vt.Description.Value
	} else if vt := u.OfMultipleCheckboxes; vt != nil && vt.Description.Valid() {
		return &vt.Description.Value
	} else if vt := u.OfDropdown; vt != nil && vt.Description.Valid() {
		return &vt.Description.Value
	} else if vt := u.OfRadio; vt != nil && vt.Description.Valid() {
		return &vt.Description.Value
	} else if vt := u.OfDatepicker; vt != nil && vt.Description.Valid() {
		return &vt.Description.Value
	} else if vt := u.OfFile; vt != nil && vt.Description.Valid() {
		return &vt.Description.Value
	} else if vt := u.OfPaymentLinkRadio; vt != nil && vt.Description.Valid() {
		return &vt.Description.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u FieldGroupFieldUnionParam) GetPlaceholder() *string {
	if vt := u.OfEmail; vt != nil && vt.Placeholder.Valid() {
		return &vt.Placeholder.Value
	} else if vt := u.OfPhone; vt != nil && vt.Placeholder.Valid() {
		return &vt.Placeholder.Value
	} else if vt := u.OfMobilePhone; vt != nil && vt.Placeholder.Valid() {
		return &vt.Placeholder.Value
	} else if vt := u.OfSingleLineText; vt != nil && vt.Placeholder.Valid() {
		return &vt.Placeholder.Value
	} else if vt := u.OfMultiLineText; vt != nil && vt.Placeholder.Valid() {
		return &vt.Placeholder.Value
	} else if vt := u.OfNumber; vt != nil && vt.Placeholder.Valid() {
		return &vt.Placeholder.Value
	} else if vt := u.OfDropdown; vt != nil && vt.Placeholder.Valid() {
		return &vt.Placeholder.Value
	} else if vt := u.OfRadio; vt != nil && vt.Placeholder.Valid() {
		return &vt.Placeholder.Value
	} else if vt := u.OfDatepicker; vt != nil && vt.Placeholder.Valid() {
		return &vt.Placeholder.Value
	} else if vt := u.OfFile; vt != nil && vt.Placeholder.Valid() {
		return &vt.Placeholder.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's DependentFields property, if
// present.
func (u FieldGroupFieldUnionParam) GetDependentFields() []DependentFieldParam {
	if vt := u.OfEmail; vt != nil {
		return vt.DependentFields
	} else if vt := u.OfPhone; vt != nil {
		return vt.DependentFields
	} else if vt := u.OfMobilePhone; vt != nil {
		return vt.DependentFields
	} else if vt := u.OfSingleLineText; vt != nil {
		return vt.DependentFields
	} else if vt := u.OfMultiLineText; vt != nil {
		return vt.DependentFields
	} else if vt := u.OfNumber; vt != nil {
		return vt.DependentFields
	} else if vt := u.OfSingleCheckbox; vt != nil {
		return vt.DependentFields
	} else if vt := u.OfMultipleCheckboxes; vt != nil {
		return vt.DependentFields
	} else if vt := u.OfDropdown; vt != nil {
		return vt.DependentFields
	} else if vt := u.OfRadio; vt != nil {
		return vt.DependentFields
	} else if vt := u.OfDatepicker; vt != nil {
		return vt.DependentFields
	} else if vt := u.OfFile; vt != nil {
		return vt.DependentFields
	} else if vt := u.OfPaymentLinkRadio; vt != nil {
		return vt.DependentFields
	}
	return nil
}

// Returns a subunion which exports methods to access subproperties
//
// Or use AsAny() to get the underlying value
func (u FieldGroupFieldUnionParam) GetValidation() (res fieldGroupFieldUnionParamValidation) {
	if vt := u.OfEmail; vt != nil {
		res.any = &vt.Validation
	} else if vt := u.OfPhone; vt != nil {
		res.any = &vt.Validation
	} else if vt := u.OfMobilePhone; vt != nil {
		res.any = &vt.Validation
	} else if vt := u.OfNumber; vt != nil {
		res.any = &vt.Validation
	}
	return
}

// Can have the runtime types [*EmailFieldValidationParam],
// [*PhoneFieldValidationParam], [*NumberFieldValidationParam]
type fieldGroupFieldUnionParamValidation struct{ any }

// Use the following switch statement to get the type of the union:
//
//	switch u.AsAny().(type) {
//	case *marketing.EmailFieldValidationParam:
//	case *marketing.PhoneFieldValidationParam:
//	case *marketing.NumberFieldValidationParam:
//	default:
//	    fmt.Errorf("not present")
//	}
func (u fieldGroupFieldUnionParamValidation) AsAny() any { return u.any }

// Returns a pointer to the underlying variant's property, if present.
func (u fieldGroupFieldUnionParamValidation) GetBlockedEmailDomains() []string {
	switch vt := u.any.(type) {
	case *EmailFieldValidationParam:
		return vt.BlockedEmailDomains
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u fieldGroupFieldUnionParamValidation) GetUseDefaultBlockList() *bool {
	switch vt := u.any.(type) {
	case *EmailFieldValidationParam:
		return &vt.UseDefaultBlockList
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u fieldGroupFieldUnionParamValidation) GetMaxAllowedDigits() *int64 {
	switch vt := u.any.(type) {
	case *PhoneFieldValidationParam:
		return (*int64)(&vt.MaxAllowedDigits)
	case *NumberFieldValidationParam:
		return (*int64)(&vt.MaxAllowedDigits)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u fieldGroupFieldUnionParamValidation) GetMinAllowedDigits() *int64 {
	switch vt := u.any.(type) {
	case *PhoneFieldValidationParam:
		return (*int64)(&vt.MinAllowedDigits)
	case *NumberFieldValidationParam:
		return (*int64)(&vt.MinAllowedDigits)
	}
	return nil
}

// Returns a pointer to the underlying variant's DefaultValues property, if
// present.
func (u FieldGroupFieldUnionParam) GetDefaultValues() []string {
	if vt := u.OfMultipleCheckboxes; vt != nil {
		return vt.DefaultValues
	} else if vt := u.OfDropdown; vt != nil {
		return vt.DefaultValues
	} else if vt := u.OfRadio; vt != nil {
		return vt.DefaultValues
	} else if vt := u.OfPaymentLinkRadio; vt != nil {
		return vt.DefaultValues
	}
	return nil
}

// Returns a pointer to the underlying variant's Options property, if present.
func (u FieldGroupFieldUnionParam) GetOptions() []EnumeratedFieldOptionParam {
	if vt := u.OfMultipleCheckboxes; vt != nil {
		return vt.Options
	} else if vt := u.OfDropdown; vt != nil {
		return vt.Options
	} else if vt := u.OfRadio; vt != nil {
		return vt.Options
	} else if vt := u.OfPaymentLinkRadio; vt != nil {
		return vt.Options
	}
	return nil
}

// A form field used for uploading one or more files.
type FileField struct {
	// Whether to allow the upload of multiple files.
	AllowMultipleFiles bool `json:"allowMultipleFiles,required"`
	// A list of other fields to make visible based on the value filled in for this
	// field.
	DependentFields []DependentField `json:"dependentFields,required"`
	// Determines how the field will be displayed and validated.
	//
	// Any of "file".
	FieldType FileFieldFieldType `json:"fieldType,required"`
	// Whether a field should be hidden or not. Hidden fields won't appear on the form,
	// but can be used to pass a value to a property without requiring the customer to
	// fill it in.
	Hidden bool `json:"hidden,required"`
	// The main label for the form field.
	Label string `json:"label,required"`
	// The identifier of the field. In combination with the object type ID, it must be
	// unique.
	Name string `json:"name,required"`
	// A unique ID for this field's CRM object type. For example a CONTACT field will
	// have the object type ID 0-1.
	ObjectTypeID string `json:"objectTypeId,required"`
	// Whether a value for this field is required when submitting the form.
	Required bool `json:"required,required"`
	// The value filled in by default. This value will be submitted unless the customer
	// modifies it.
	DefaultValue string `json:"defaultValue"`
	// Additional text helping the customer to complete the field.
	Description string `json:"description"`
	// The prompt text showing when the field isn't filled in.
	Placeholder string `json:"placeholder"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AllowMultipleFiles respjson.Field
		DependentFields    respjson.Field
		FieldType          respjson.Field
		Hidden             respjson.Field
		Label              respjson.Field
		Name               respjson.Field
		ObjectTypeID       respjson.Field
		Required           respjson.Field
		DefaultValue       respjson.Field
		Description        respjson.Field
		Placeholder        respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FileField) RawJSON() string { return r.JSON.raw }
func (r *FileField) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this FileField to a FileFieldParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// FileFieldParam.Overrides()
func (r FileField) ToParam() FileFieldParam {
	return param.Override[FileFieldParam](json.RawMessage(r.RawJSON()))
}

// Determines how the field will be displayed and validated.
type FileFieldFieldType string

const (
	FileFieldFieldTypeFile FileFieldFieldType = "file"
)

// A form field used for uploading one or more files.
//
// The properties AllowMultipleFiles, DependentFields, FieldType, Hidden, Label,
// Name, ObjectTypeID, Required are required.
type FileFieldParam struct {
	// Whether to allow the upload of multiple files.
	AllowMultipleFiles bool `json:"allowMultipleFiles,required"`
	// A list of other fields to make visible based on the value filled in for this
	// field.
	DependentFields []DependentFieldParam `json:"dependentFields,omitzero,required"`
	// Determines how the field will be displayed and validated.
	//
	// Any of "file".
	FieldType FileFieldFieldType `json:"fieldType,omitzero,required"`
	// Whether a field should be hidden or not. Hidden fields won't appear on the form,
	// but can be used to pass a value to a property without requiring the customer to
	// fill it in.
	Hidden bool `json:"hidden,required"`
	// The main label for the form field.
	Label string `json:"label,required"`
	// The identifier of the field. In combination with the object type ID, it must be
	// unique.
	Name string `json:"name,required"`
	// A unique ID for this field's CRM object type. For example a CONTACT field will
	// have the object type ID 0-1.
	ObjectTypeID string `json:"objectTypeId,required"`
	// Whether a value for this field is required when submitting the form.
	Required bool `json:"required,required"`
	// The value filled in by default. This value will be submitted unless the customer
	// modifies it.
	DefaultValue param.Opt[string] `json:"defaultValue,omitzero"`
	// Additional text helping the customer to complete the field.
	Description param.Opt[string] `json:"description,omitzero"`
	// The prompt text showing when the field isn't filled in.
	Placeholder param.Opt[string] `json:"placeholder,omitzero"`
	paramObj
}

func (r FileFieldParam) MarshalJSON() (data []byte, err error) {
	type shadow FileFieldParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FileFieldParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FormDefinitionBase struct {
	ID            string                   `json:"id,required"`
	Archived      bool                     `json:"archived,required"`
	Configuration HubSpotFormConfiguration `json:"configuration,required"`
	CreatedAt     time.Time                `json:"createdAt,required" format:"date-time"`
	// Options for styling the form.
	DisplayOptions FormDisplayOptions `json:"displayOptions,required"`
	FieldGroups    []FieldGroup       `json:"fieldGroups,required"`
	// Any of "hubspot".
	FormType            FormDefinitionBaseFormType                 `json:"formType,required"`
	LegalConsentOptions FormDefinitionBaseLegalConsentOptionsUnion `json:"legalConsentOptions,required"`
	Name                string                                     `json:"name,required"`
	UpdatedAt           time.Time                                  `json:"updatedAt,required" format:"date-time"`
	ArchivedAt          time.Time                                  `json:"archivedAt" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                  respjson.Field
		Archived            respjson.Field
		Configuration       respjson.Field
		CreatedAt           respjson.Field
		DisplayOptions      respjson.Field
		FieldGroups         respjson.Field
		FormType            respjson.Field
		LegalConsentOptions respjson.Field
		Name                respjson.Field
		UpdatedAt           respjson.Field
		ArchivedAt          respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FormDefinitionBase) RawJSON() string { return r.JSON.raw }
func (r *FormDefinitionBase) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FormDefinitionBaseFormType string

const (
	FormDefinitionBaseFormTypeHubspot FormDefinitionBaseFormType = "hubspot"
)

// FormDefinitionBaseLegalConsentOptionsUnion contains all possible properties and
// values from [LegalConsentOptionsNone], [LegalConsentOptionsLegitimateInterest],
// [LegalConsentOptionsExplicitConsentToProcess],
// [LegalConsentOptionsImplicitConsentToProcess].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type FormDefinitionBaseLegalConsentOptionsUnion struct {
	Type string `json:"type"`
	// This field is from variant [LegalConsentOptionsLegitimateInterest].
	LawfulBasis LegalConsentOptionsLegitimateInterestLawfulBasis `json:"lawfulBasis"`
	PrivacyText string                                           `json:"privacyText"`
	// This field is from variant [LegalConsentOptionsLegitimateInterest].
	SubscriptionTypeIDs      []int64                `json:"subscriptionTypeIds"`
	CommunicationsCheckboxes []LegalConsentCheckbox `json:"communicationsCheckboxes"`
	CommunicationConsentText string                 `json:"communicationConsentText"`
	// This field is from variant [LegalConsentOptionsExplicitConsentToProcess].
	ConsentToProcessCheckboxLabel string `json:"consentToProcessCheckboxLabel"`
	// This field is from variant [LegalConsentOptionsExplicitConsentToProcess].
	ConsentToProcessFooterText string `json:"consentToProcessFooterText"`
	ConsentToProcessText       string `json:"consentToProcessText"`
	JSON                       struct {
		Type                          respjson.Field
		LawfulBasis                   respjson.Field
		PrivacyText                   respjson.Field
		SubscriptionTypeIDs           respjson.Field
		CommunicationsCheckboxes      respjson.Field
		CommunicationConsentText      respjson.Field
		ConsentToProcessCheckboxLabel respjson.Field
		ConsentToProcessFooterText    respjson.Field
		ConsentToProcessText          respjson.Field
		raw                           string
	} `json:"-"`
}

func (u FormDefinitionBaseLegalConsentOptionsUnion) AsNone() (v LegalConsentOptionsNone) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FormDefinitionBaseLegalConsentOptionsUnion) AsLegitimateInterest() (v LegalConsentOptionsLegitimateInterest) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FormDefinitionBaseLegalConsentOptionsUnion) AsExplicitConsentToProcess() (v LegalConsentOptionsExplicitConsentToProcess) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FormDefinitionBaseLegalConsentOptionsUnion) AsImplicitConsentToProcess() (v LegalConsentOptionsImplicitConsentToProcess) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u FormDefinitionBaseLegalConsentOptionsUnion) RawJSON() string { return u.JSON.raw }

func (r *FormDefinitionBaseLegalConsentOptionsUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Archived, Configuration, CreatedAt, DisplayOptions, FieldGroups,
// FormType, LegalConsentOptions, Name, UpdatedAt are required.
type FormDefinitionCreateRequestBaseParam struct {
	Archived      bool                          `json:"archived,required"`
	Configuration HubSpotFormConfigurationParam `json:"configuration,omitzero,required"`
	CreatedAt     time.Time                     `json:"createdAt,required" format:"date-time"`
	// Options for styling the form.
	DisplayOptions FormDisplayOptionsParam `json:"displayOptions,omitzero,required"`
	FieldGroups    []FieldGroupParam       `json:"fieldGroups,omitzero,required"`
	// Any of "hubspot".
	FormType            FormDefinitionCreateRequestBaseFormType                      `json:"formType,omitzero,required"`
	LegalConsentOptions FormDefinitionCreateRequestBaseLegalConsentOptionsUnionParam `json:"legalConsentOptions,omitzero,required"`
	Name                string                                                       `json:"name,required"`
	UpdatedAt           time.Time                                                    `json:"updatedAt,required" format:"date-time"`
	ArchivedAt          param.Opt[time.Time]                                         `json:"archivedAt,omitzero" format:"date-time"`
	paramObj
}

func (r FormDefinitionCreateRequestBaseParam) MarshalJSON() (data []byte, err error) {
	type shadow FormDefinitionCreateRequestBaseParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FormDefinitionCreateRequestBaseParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FormDefinitionCreateRequestBaseFormType string

const (
	FormDefinitionCreateRequestBaseFormTypeHubspot FormDefinitionCreateRequestBaseFormType = "hubspot"
)

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type FormDefinitionCreateRequestBaseLegalConsentOptionsUnionParam struct {
	OfNone                     *LegalConsentOptionsNoneParam                     `json:",omitzero,inline"`
	OfLegitimateInterest       *LegalConsentOptionsLegitimateInterestParam       `json:",omitzero,inline"`
	OfExplicitConsentToProcess *LegalConsentOptionsExplicitConsentToProcessParam `json:",omitzero,inline"`
	OfImplicitConsentToProcess *LegalConsentOptionsImplicitConsentToProcessParam `json:",omitzero,inline"`
	paramUnion
}

func (u FormDefinitionCreateRequestBaseLegalConsentOptionsUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfNone, u.OfLegitimateInterest, u.OfExplicitConsentToProcess, u.OfImplicitConsentToProcess)
}
func (u *FormDefinitionCreateRequestBaseLegalConsentOptionsUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *FormDefinitionCreateRequestBaseLegalConsentOptionsUnionParam) asAny() any {
	if !param.IsOmitted(u.OfNone) {
		return u.OfNone
	} else if !param.IsOmitted(u.OfLegitimateInterest) {
		return u.OfLegitimateInterest
	} else if !param.IsOmitted(u.OfExplicitConsentToProcess) {
		return u.OfExplicitConsentToProcess
	} else if !param.IsOmitted(u.OfImplicitConsentToProcess) {
		return u.OfImplicitConsentToProcess
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u FormDefinitionCreateRequestBaseLegalConsentOptionsUnionParam) GetLawfulBasis() *string {
	if vt := u.OfLegitimateInterest; vt != nil {
		return (*string)(&vt.LawfulBasis)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u FormDefinitionCreateRequestBaseLegalConsentOptionsUnionParam) GetSubscriptionTypeIDs() []int64 {
	if vt := u.OfLegitimateInterest; vt != nil {
		return vt.SubscriptionTypeIDs
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u FormDefinitionCreateRequestBaseLegalConsentOptionsUnionParam) GetConsentToProcessCheckboxLabel() *string {
	if vt := u.OfExplicitConsentToProcess; vt != nil && vt.ConsentToProcessCheckboxLabel.Valid() {
		return &vt.ConsentToProcessCheckboxLabel.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u FormDefinitionCreateRequestBaseLegalConsentOptionsUnionParam) GetConsentToProcessFooterText() *string {
	if vt := u.OfExplicitConsentToProcess; vt != nil && vt.ConsentToProcessFooterText.Valid() {
		return &vt.ConsentToProcessFooterText.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u FormDefinitionCreateRequestBaseLegalConsentOptionsUnionParam) GetType() *string {
	if vt := u.OfNone; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfLegitimateInterest; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfExplicitConsentToProcess; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfImplicitConsentToProcess; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u FormDefinitionCreateRequestBaseLegalConsentOptionsUnionParam) GetPrivacyText() *string {
	if vt := u.OfLegitimateInterest; vt != nil {
		return (*string)(&vt.PrivacyText)
	} else if vt := u.OfExplicitConsentToProcess; vt != nil {
		return (*string)(&vt.PrivacyText)
	} else if vt := u.OfImplicitConsentToProcess; vt != nil {
		return (*string)(&vt.PrivacyText)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u FormDefinitionCreateRequestBaseLegalConsentOptionsUnionParam) GetCommunicationConsentText() *string {
	if vt := u.OfExplicitConsentToProcess; vt != nil && vt.CommunicationConsentText.Valid() {
		return &vt.CommunicationConsentText.Value
	} else if vt := u.OfImplicitConsentToProcess; vt != nil && vt.CommunicationConsentText.Valid() {
		return &vt.CommunicationConsentText.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u FormDefinitionCreateRequestBaseLegalConsentOptionsUnionParam) GetConsentToProcessText() *string {
	if vt := u.OfExplicitConsentToProcess; vt != nil && vt.ConsentToProcessText.Valid() {
		return &vt.ConsentToProcessText.Value
	} else if vt := u.OfImplicitConsentToProcess; vt != nil && vt.ConsentToProcessText.Valid() {
		return &vt.ConsentToProcessText.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's CommunicationsCheckboxes property,
// if present.
func (u FormDefinitionCreateRequestBaseLegalConsentOptionsUnionParam) GetCommunicationsCheckboxes() []LegalConsentCheckboxParam {
	if vt := u.OfExplicitConsentToProcess; vt != nil {
		return vt.CommunicationsCheckboxes
	} else if vt := u.OfImplicitConsentToProcess; vt != nil {
		return vt.CommunicationsCheckboxes
	}
	return nil
}

// Options for styling the form.
type FormDisplayOptions struct {
	// Whether the form will render as raw HTML as opposed to inside an iFrame.
	RenderRawHTML bool `json:"renderRawHtml,required"`
	// Styling options for the form
	Style FormStyle `json:"style,required"`
	// The text displayed on the form submit button.
	SubmitButtonText string `json:"submitButtonText,required"`
	// The theme used for styling the input fields. This will not apply if the form is
	// added to a HubSpot CMS page.
	//
	// Any of "default_style", "canvas", "linear", "round", "sharp", "legacy".
	Theme    FormDisplayOptionsTheme `json:"theme,required"`
	CssClass string                  `json:"cssClass"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		RenderRawHTML    respjson.Field
		Style            respjson.Field
		SubmitButtonText respjson.Field
		Theme            respjson.Field
		CssClass         respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FormDisplayOptions) RawJSON() string { return r.JSON.raw }
func (r *FormDisplayOptions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this FormDisplayOptions to a FormDisplayOptionsParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// FormDisplayOptionsParam.Overrides()
func (r FormDisplayOptions) ToParam() FormDisplayOptionsParam {
	return param.Override[FormDisplayOptionsParam](json.RawMessage(r.RawJSON()))
}

// The theme used for styling the input fields. This will not apply if the form is
// added to a HubSpot CMS page.
type FormDisplayOptionsTheme string

const (
	FormDisplayOptionsThemeDefaultStyle FormDisplayOptionsTheme = "default_style"
	FormDisplayOptionsThemeCanvas       FormDisplayOptionsTheme = "canvas"
	FormDisplayOptionsThemeLinear       FormDisplayOptionsTheme = "linear"
	FormDisplayOptionsThemeRound        FormDisplayOptionsTheme = "round"
	FormDisplayOptionsThemeSharp        FormDisplayOptionsTheme = "sharp"
	FormDisplayOptionsThemeLegacy       FormDisplayOptionsTheme = "legacy"
)

// Options for styling the form.
//
// The properties RenderRawHTML, Style, SubmitButtonText, Theme are required.
type FormDisplayOptionsParam struct {
	// Whether the form will render as raw HTML as opposed to inside an iFrame.
	RenderRawHTML bool `json:"renderRawHtml,required"`
	// Styling options for the form
	Style FormStyleParam `json:"style,omitzero,required"`
	// The text displayed on the form submit button.
	SubmitButtonText string `json:"submitButtonText,required"`
	// The theme used for styling the input fields. This will not apply if the form is
	// added to a HubSpot CMS page.
	//
	// Any of "default_style", "canvas", "linear", "round", "sharp", "legacy".
	Theme    FormDisplayOptionsTheme `json:"theme,omitzero,required"`
	CssClass param.Opt[string]       `json:"cssClass,omitzero"`
	paramObj
}

func (r FormDisplayOptionsParam) MarshalJSON() (data []byte, err error) {
	type shadow FormDisplayOptionsParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FormDisplayOptionsParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// What should happen after the customer submits the form.
type FormPostSubmitAction struct {
	// The action to take after submit. The default action is displaying a thank you
	// message.
	//
	// Any of "thank_you", "redirect_url".
	Type FormPostSubmitActionType `json:"type,required"`
	// The thank you text or the page to redirect to.
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
func (r FormPostSubmitAction) RawJSON() string { return r.JSON.raw }
func (r *FormPostSubmitAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this FormPostSubmitAction to a FormPostSubmitActionParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// FormPostSubmitActionParam.Overrides()
func (r FormPostSubmitAction) ToParam() FormPostSubmitActionParam {
	return param.Override[FormPostSubmitActionParam](json.RawMessage(r.RawJSON()))
}

// The action to take after submit. The default action is displaying a thank you
// message.
type FormPostSubmitActionType string

const (
	FormPostSubmitActionTypeThankYou    FormPostSubmitActionType = "thank_you"
	FormPostSubmitActionTypeRedirectURL FormPostSubmitActionType = "redirect_url"
)

// What should happen after the customer submits the form.
//
// The properties Type, Value are required.
type FormPostSubmitActionParam struct {
	// The action to take after submit. The default action is displaying a thank you
	// message.
	//
	// Any of "thank_you", "redirect_url".
	Type FormPostSubmitActionType `json:"type,omitzero,required"`
	// The thank you text or the page to redirect to.
	Value string `json:"value,required"`
	paramObj
}

func (r FormPostSubmitActionParam) MarshalJSON() (data []byte, err error) {
	type shadow FormPostSubmitActionParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FormPostSubmitActionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Styling options for the form
type FormStyle struct {
	BackgroundWidth       string `json:"backgroundWidth,required"`
	FontFamily            string `json:"fontFamily,required"`
	HelpTextColor         string `json:"helpTextColor,required"`
	HelpTextSize          string `json:"helpTextSize,required"`
	LabelTextColor        string `json:"labelTextColor,required"`
	LabelTextSize         string `json:"labelTextSize,required"`
	LegalConsentTextColor string `json:"legalConsentTextColor,required"`
	LegalConsentTextSize  string `json:"legalConsentTextSize,required"`
	// Any of "left", "right", "center".
	SubmitAlignment FormStyleSubmitAlignment `json:"submitAlignment,required"`
	SubmitColor     string                   `json:"submitColor,required"`
	SubmitFontColor string                   `json:"submitFontColor,required"`
	SubmitSize      string                   `json:"submitSize,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BackgroundWidth       respjson.Field
		FontFamily            respjson.Field
		HelpTextColor         respjson.Field
		HelpTextSize          respjson.Field
		LabelTextColor        respjson.Field
		LabelTextSize         respjson.Field
		LegalConsentTextColor respjson.Field
		LegalConsentTextSize  respjson.Field
		SubmitAlignment       respjson.Field
		SubmitColor           respjson.Field
		SubmitFontColor       respjson.Field
		SubmitSize            respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FormStyle) RawJSON() string { return r.JSON.raw }
func (r *FormStyle) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this FormStyle to a FormStyleParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// FormStyleParam.Overrides()
func (r FormStyle) ToParam() FormStyleParam {
	return param.Override[FormStyleParam](json.RawMessage(r.RawJSON()))
}

type FormStyleSubmitAlignment string

const (
	FormStyleSubmitAlignmentLeft   FormStyleSubmitAlignment = "left"
	FormStyleSubmitAlignmentRight  FormStyleSubmitAlignment = "right"
	FormStyleSubmitAlignmentCenter FormStyleSubmitAlignment = "center"
)

// Styling options for the form
//
// The properties BackgroundWidth, FontFamily, HelpTextColor, HelpTextSize,
// LabelTextColor, LabelTextSize, LegalConsentTextColor, LegalConsentTextSize,
// SubmitAlignment, SubmitColor, SubmitFontColor, SubmitSize are required.
type FormStyleParam struct {
	BackgroundWidth       string `json:"backgroundWidth,required"`
	FontFamily            string `json:"fontFamily,required"`
	HelpTextColor         string `json:"helpTextColor,required"`
	HelpTextSize          string `json:"helpTextSize,required"`
	LabelTextColor        string `json:"labelTextColor,required"`
	LabelTextSize         string `json:"labelTextSize,required"`
	LegalConsentTextColor string `json:"legalConsentTextColor,required"`
	LegalConsentTextSize  string `json:"legalConsentTextSize,required"`
	// Any of "left", "right", "center".
	SubmitAlignment FormStyleSubmitAlignment `json:"submitAlignment,omitzero,required"`
	SubmitColor     string                   `json:"submitColor,required"`
	SubmitFontColor string                   `json:"submitFontColor,required"`
	SubmitSize      string                   `json:"submitSize,required"`
	paramObj
}

func (r FormStyleParam) MarshalJSON() (data []byte, err error) {
	type shadow FormStyleParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FormStyleParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type HubSpotFormConfiguration struct {
	// Whether to add a reset link to the form. This removes any pre-populated content
	// on the form and creates a new contact on submission.
	AllowLinkToResetKnownValues bool `json:"allowLinkToResetKnownValues,required"`
	// Whether the form can be archived.
	Archivable bool `json:"archivable,required"`
	// Whether the form can be cloned.
	Cloneable bool `json:"cloneable,required"`
	// Whether to create a new contact when a form is submitted with an email address
	// that doesn’t match any in your existing contacts records.
	CreateNewContactForNewEmail bool `json:"createNewContactForNewEmail,required"`
	// Whether the form can be edited.
	Editable bool `json:"editable,required"`
	// The language of the form.
	//
	// Any of "af", "ar-eg", "bg", "bn", "ca-es", "cs", "da", "de", "el", "en", "es",
	// "es-mx", "fi", "fr", "fr-ca", "he-il", "hr", "hu", "id", "it", "ja", "ko", "lt",
	// "ms", "nl", "no-no", "pl", "pt", "pt-br", "ro", "ru", "sk", "sl", "sv", "th",
	// "tl", "tr", "uk", "vi", "zh-cn", "zh-hk", "zh-tw".
	Language HubSpotFormConfigurationLanguage `json:"language,required"`
	// Whether to send a notification email to the contact owner when a submission is
	// received.
	NotifyContactOwner bool `json:"notifyContactOwner,required"`
	// The list of user IDs to receive a notification email when a submission is
	// received.
	NotifyRecipients []string `json:"notifyRecipients,required"`
	// What should happen after the customer submits the form.
	PostSubmitAction FormPostSubmitAction `json:"postSubmitAction,required"`
	// Whether contact fields should pre-populate with known information when a contact
	// returns to your site.
	PrePopulateKnownValues bool `json:"prePopulateKnownValues,required"`
	// Whether CAPTCHA (spam prevention) is enabled.
	RecaptchaEnabled bool             `json:"recaptchaEnabled,required"`
	LifecycleStages  []LifecycleStage `json:"lifecycleStages"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AllowLinkToResetKnownValues respjson.Field
		Archivable                  respjson.Field
		Cloneable                   respjson.Field
		CreateNewContactForNewEmail respjson.Field
		Editable                    respjson.Field
		Language                    respjson.Field
		NotifyContactOwner          respjson.Field
		NotifyRecipients            respjson.Field
		PostSubmitAction            respjson.Field
		PrePopulateKnownValues      respjson.Field
		RecaptchaEnabled            respjson.Field
		LifecycleStages             respjson.Field
		ExtraFields                 map[string]respjson.Field
		raw                         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r HubSpotFormConfiguration) RawJSON() string { return r.JSON.raw }
func (r *HubSpotFormConfiguration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this HubSpotFormConfiguration to a
// HubSpotFormConfigurationParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// HubSpotFormConfigurationParam.Overrides()
func (r HubSpotFormConfiguration) ToParam() HubSpotFormConfigurationParam {
	return param.Override[HubSpotFormConfigurationParam](json.RawMessage(r.RawJSON()))
}

// The language of the form.
type HubSpotFormConfigurationLanguage string

const (
	HubSpotFormConfigurationLanguageAf   HubSpotFormConfigurationLanguage = "af"
	HubSpotFormConfigurationLanguageArEg HubSpotFormConfigurationLanguage = "ar-eg"
	HubSpotFormConfigurationLanguageBg   HubSpotFormConfigurationLanguage = "bg"
	HubSpotFormConfigurationLanguageBn   HubSpotFormConfigurationLanguage = "bn"
	HubSpotFormConfigurationLanguageCaEs HubSpotFormConfigurationLanguage = "ca-es"
	HubSpotFormConfigurationLanguageCs   HubSpotFormConfigurationLanguage = "cs"
	HubSpotFormConfigurationLanguageDa   HubSpotFormConfigurationLanguage = "da"
	HubSpotFormConfigurationLanguageDe   HubSpotFormConfigurationLanguage = "de"
	HubSpotFormConfigurationLanguageEl   HubSpotFormConfigurationLanguage = "el"
	HubSpotFormConfigurationLanguageEn   HubSpotFormConfigurationLanguage = "en"
	HubSpotFormConfigurationLanguageEs   HubSpotFormConfigurationLanguage = "es"
	HubSpotFormConfigurationLanguageEsMx HubSpotFormConfigurationLanguage = "es-mx"
	HubSpotFormConfigurationLanguageFi   HubSpotFormConfigurationLanguage = "fi"
	HubSpotFormConfigurationLanguageFr   HubSpotFormConfigurationLanguage = "fr"
	HubSpotFormConfigurationLanguageFrCa HubSpotFormConfigurationLanguage = "fr-ca"
	HubSpotFormConfigurationLanguageHeIl HubSpotFormConfigurationLanguage = "he-il"
	HubSpotFormConfigurationLanguageHr   HubSpotFormConfigurationLanguage = "hr"
	HubSpotFormConfigurationLanguageHu   HubSpotFormConfigurationLanguage = "hu"
	HubSpotFormConfigurationLanguageID   HubSpotFormConfigurationLanguage = "id"
	HubSpotFormConfigurationLanguageIt   HubSpotFormConfigurationLanguage = "it"
	HubSpotFormConfigurationLanguageJa   HubSpotFormConfigurationLanguage = "ja"
	HubSpotFormConfigurationLanguageKo   HubSpotFormConfigurationLanguage = "ko"
	HubSpotFormConfigurationLanguageLt   HubSpotFormConfigurationLanguage = "lt"
	HubSpotFormConfigurationLanguageMs   HubSpotFormConfigurationLanguage = "ms"
	HubSpotFormConfigurationLanguageNl   HubSpotFormConfigurationLanguage = "nl"
	HubSpotFormConfigurationLanguageNoNo HubSpotFormConfigurationLanguage = "no-no"
	HubSpotFormConfigurationLanguagePl   HubSpotFormConfigurationLanguage = "pl"
	HubSpotFormConfigurationLanguagePt   HubSpotFormConfigurationLanguage = "pt"
	HubSpotFormConfigurationLanguagePtBr HubSpotFormConfigurationLanguage = "pt-br"
	HubSpotFormConfigurationLanguageRo   HubSpotFormConfigurationLanguage = "ro"
	HubSpotFormConfigurationLanguageRu   HubSpotFormConfigurationLanguage = "ru"
	HubSpotFormConfigurationLanguageSk   HubSpotFormConfigurationLanguage = "sk"
	HubSpotFormConfigurationLanguageSl   HubSpotFormConfigurationLanguage = "sl"
	HubSpotFormConfigurationLanguageSv   HubSpotFormConfigurationLanguage = "sv"
	HubSpotFormConfigurationLanguageTh   HubSpotFormConfigurationLanguage = "th"
	HubSpotFormConfigurationLanguageTl   HubSpotFormConfigurationLanguage = "tl"
	HubSpotFormConfigurationLanguageTr   HubSpotFormConfigurationLanguage = "tr"
	HubSpotFormConfigurationLanguageUk   HubSpotFormConfigurationLanguage = "uk"
	HubSpotFormConfigurationLanguageVi   HubSpotFormConfigurationLanguage = "vi"
	HubSpotFormConfigurationLanguageZhCn HubSpotFormConfigurationLanguage = "zh-cn"
	HubSpotFormConfigurationLanguageZhHk HubSpotFormConfigurationLanguage = "zh-hk"
	HubSpotFormConfigurationLanguageZhTw HubSpotFormConfigurationLanguage = "zh-tw"
)

// The properties AllowLinkToResetKnownValues, Archivable, Cloneable,
// CreateNewContactForNewEmail, Editable, Language, NotifyContactOwner,
// NotifyRecipients, PostSubmitAction, PrePopulateKnownValues, RecaptchaEnabled are
// required.
type HubSpotFormConfigurationParam struct {
	// Whether to add a reset link to the form. This removes any pre-populated content
	// on the form and creates a new contact on submission.
	AllowLinkToResetKnownValues bool `json:"allowLinkToResetKnownValues,required"`
	// Whether the form can be archived.
	Archivable bool `json:"archivable,required"`
	// Whether the form can be cloned.
	Cloneable bool `json:"cloneable,required"`
	// Whether to create a new contact when a form is submitted with an email address
	// that doesn’t match any in your existing contacts records.
	CreateNewContactForNewEmail bool `json:"createNewContactForNewEmail,required"`
	// Whether the form can be edited.
	Editable bool `json:"editable,required"`
	// The language of the form.
	//
	// Any of "af", "ar-eg", "bg", "bn", "ca-es", "cs", "da", "de", "el", "en", "es",
	// "es-mx", "fi", "fr", "fr-ca", "he-il", "hr", "hu", "id", "it", "ja", "ko", "lt",
	// "ms", "nl", "no-no", "pl", "pt", "pt-br", "ro", "ru", "sk", "sl", "sv", "th",
	// "tl", "tr", "uk", "vi", "zh-cn", "zh-hk", "zh-tw".
	Language HubSpotFormConfigurationLanguage `json:"language,omitzero,required"`
	// Whether to send a notification email to the contact owner when a submission is
	// received.
	NotifyContactOwner bool `json:"notifyContactOwner,required"`
	// The list of user IDs to receive a notification email when a submission is
	// received.
	NotifyRecipients []string `json:"notifyRecipients,omitzero,required"`
	// What should happen after the customer submits the form.
	PostSubmitAction FormPostSubmitActionParam `json:"postSubmitAction,omitzero,required"`
	// Whether contact fields should pre-populate with known information when a contact
	// returns to your site.
	PrePopulateKnownValues bool `json:"prePopulateKnownValues,required"`
	// Whether CAPTCHA (spam prevention) is enabled.
	RecaptchaEnabled bool                  `json:"recaptchaEnabled,required"`
	LifecycleStages  []LifecycleStageParam `json:"lifecycleStages,omitzero"`
	paramObj
}

func (r HubSpotFormConfigurationParam) MarshalJSON() (data []byte, err error) {
	type shadow HubSpotFormConfigurationParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *HubSpotFormConfigurationParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type HubSpotFormDefinition struct {
	ID            string                   `json:"id,required"`
	Archived      bool                     `json:"archived,required"`
	Configuration HubSpotFormConfiguration `json:"configuration,required"`
	CreatedAt     time.Time                `json:"createdAt,required" format:"date-time"`
	// Options for styling the form.
	DisplayOptions FormDisplayOptions `json:"displayOptions,required"`
	FieldGroups    []FieldGroup       `json:"fieldGroups,required"`
	// Any of "hubspot".
	FormType            HubSpotFormDefinitionFormType                 `json:"formType,required"`
	LegalConsentOptions HubSpotFormDefinitionLegalConsentOptionsUnion `json:"legalConsentOptions,required"`
	Name                string                                        `json:"name,required"`
	UpdatedAt           time.Time                                     `json:"updatedAt,required" format:"date-time"`
	ArchivedAt          time.Time                                     `json:"archivedAt" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                  respjson.Field
		Archived            respjson.Field
		Configuration       respjson.Field
		CreatedAt           respjson.Field
		DisplayOptions      respjson.Field
		FieldGroups         respjson.Field
		FormType            respjson.Field
		LegalConsentOptions respjson.Field
		Name                respjson.Field
		UpdatedAt           respjson.Field
		ArchivedAt          respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r HubSpotFormDefinition) RawJSON() string { return r.JSON.raw }
func (r *HubSpotFormDefinition) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this HubSpotFormDefinition to a HubSpotFormDefinitionParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// HubSpotFormDefinitionParam.Overrides()
func (r HubSpotFormDefinition) ToParam() HubSpotFormDefinitionParam {
	return param.Override[HubSpotFormDefinitionParam](json.RawMessage(r.RawJSON()))
}

type HubSpotFormDefinitionFormType string

const (
	HubSpotFormDefinitionFormTypeHubspot HubSpotFormDefinitionFormType = "hubspot"
)

// HubSpotFormDefinitionLegalConsentOptionsUnion contains all possible properties
// and values from [LegalConsentOptionsNone],
// [LegalConsentOptionsLegitimateInterest],
// [LegalConsentOptionsExplicitConsentToProcess],
// [LegalConsentOptionsImplicitConsentToProcess].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type HubSpotFormDefinitionLegalConsentOptionsUnion struct {
	Type string `json:"type"`
	// This field is from variant [LegalConsentOptionsLegitimateInterest].
	LawfulBasis LegalConsentOptionsLegitimateInterestLawfulBasis `json:"lawfulBasis"`
	PrivacyText string                                           `json:"privacyText"`
	// This field is from variant [LegalConsentOptionsLegitimateInterest].
	SubscriptionTypeIDs      []int64                `json:"subscriptionTypeIds"`
	CommunicationsCheckboxes []LegalConsentCheckbox `json:"communicationsCheckboxes"`
	CommunicationConsentText string                 `json:"communicationConsentText"`
	// This field is from variant [LegalConsentOptionsExplicitConsentToProcess].
	ConsentToProcessCheckboxLabel string `json:"consentToProcessCheckboxLabel"`
	// This field is from variant [LegalConsentOptionsExplicitConsentToProcess].
	ConsentToProcessFooterText string `json:"consentToProcessFooterText"`
	ConsentToProcessText       string `json:"consentToProcessText"`
	JSON                       struct {
		Type                          respjson.Field
		LawfulBasis                   respjson.Field
		PrivacyText                   respjson.Field
		SubscriptionTypeIDs           respjson.Field
		CommunicationsCheckboxes      respjson.Field
		CommunicationConsentText      respjson.Field
		ConsentToProcessCheckboxLabel respjson.Field
		ConsentToProcessFooterText    respjson.Field
		ConsentToProcessText          respjson.Field
		raw                           string
	} `json:"-"`
}

func (u HubSpotFormDefinitionLegalConsentOptionsUnion) AsNone() (v LegalConsentOptionsNone) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u HubSpotFormDefinitionLegalConsentOptionsUnion) AsLegitimateInterest() (v LegalConsentOptionsLegitimateInterest) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u HubSpotFormDefinitionLegalConsentOptionsUnion) AsExplicitConsentToProcess() (v LegalConsentOptionsExplicitConsentToProcess) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u HubSpotFormDefinitionLegalConsentOptionsUnion) AsImplicitConsentToProcess() (v LegalConsentOptionsImplicitConsentToProcess) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u HubSpotFormDefinitionLegalConsentOptionsUnion) RawJSON() string { return u.JSON.raw }

func (r *HubSpotFormDefinitionLegalConsentOptionsUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties ID, Archived, Configuration, CreatedAt, DisplayOptions,
// FieldGroups, FormType, LegalConsentOptions, Name, UpdatedAt are required.
type HubSpotFormDefinitionParam struct {
	ID            string                        `json:"id,required"`
	Archived      bool                          `json:"archived,required"`
	Configuration HubSpotFormConfigurationParam `json:"configuration,omitzero,required"`
	CreatedAt     time.Time                     `json:"createdAt,required" format:"date-time"`
	// Options for styling the form.
	DisplayOptions FormDisplayOptionsParam `json:"displayOptions,omitzero,required"`
	FieldGroups    []FieldGroupParam       `json:"fieldGroups,omitzero,required"`
	// Any of "hubspot".
	FormType            HubSpotFormDefinitionFormType                      `json:"formType,omitzero,required"`
	LegalConsentOptions HubSpotFormDefinitionLegalConsentOptionsUnionParam `json:"legalConsentOptions,omitzero,required"`
	Name                string                                             `json:"name,required"`
	UpdatedAt           time.Time                                          `json:"updatedAt,required" format:"date-time"`
	ArchivedAt          param.Opt[time.Time]                               `json:"archivedAt,omitzero" format:"date-time"`
	paramObj
}

func (r HubSpotFormDefinitionParam) MarshalJSON() (data []byte, err error) {
	type shadow HubSpotFormDefinitionParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *HubSpotFormDefinitionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type HubSpotFormDefinitionLegalConsentOptionsUnionParam struct {
	OfNone                     *LegalConsentOptionsNoneParam                     `json:",omitzero,inline"`
	OfLegitimateInterest       *LegalConsentOptionsLegitimateInterestParam       `json:",omitzero,inline"`
	OfExplicitConsentToProcess *LegalConsentOptionsExplicitConsentToProcessParam `json:",omitzero,inline"`
	OfImplicitConsentToProcess *LegalConsentOptionsImplicitConsentToProcessParam `json:",omitzero,inline"`
	paramUnion
}

func (u HubSpotFormDefinitionLegalConsentOptionsUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfNone, u.OfLegitimateInterest, u.OfExplicitConsentToProcess, u.OfImplicitConsentToProcess)
}
func (u *HubSpotFormDefinitionLegalConsentOptionsUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *HubSpotFormDefinitionLegalConsentOptionsUnionParam) asAny() any {
	if !param.IsOmitted(u.OfNone) {
		return u.OfNone
	} else if !param.IsOmitted(u.OfLegitimateInterest) {
		return u.OfLegitimateInterest
	} else if !param.IsOmitted(u.OfExplicitConsentToProcess) {
		return u.OfExplicitConsentToProcess
	} else if !param.IsOmitted(u.OfImplicitConsentToProcess) {
		return u.OfImplicitConsentToProcess
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u HubSpotFormDefinitionLegalConsentOptionsUnionParam) GetLawfulBasis() *string {
	if vt := u.OfLegitimateInterest; vt != nil {
		return (*string)(&vt.LawfulBasis)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u HubSpotFormDefinitionLegalConsentOptionsUnionParam) GetSubscriptionTypeIDs() []int64 {
	if vt := u.OfLegitimateInterest; vt != nil {
		return vt.SubscriptionTypeIDs
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u HubSpotFormDefinitionLegalConsentOptionsUnionParam) GetConsentToProcessCheckboxLabel() *string {
	if vt := u.OfExplicitConsentToProcess; vt != nil && vt.ConsentToProcessCheckboxLabel.Valid() {
		return &vt.ConsentToProcessCheckboxLabel.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u HubSpotFormDefinitionLegalConsentOptionsUnionParam) GetConsentToProcessFooterText() *string {
	if vt := u.OfExplicitConsentToProcess; vt != nil && vt.ConsentToProcessFooterText.Valid() {
		return &vt.ConsentToProcessFooterText.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u HubSpotFormDefinitionLegalConsentOptionsUnionParam) GetType() *string {
	if vt := u.OfNone; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfLegitimateInterest; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfExplicitConsentToProcess; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfImplicitConsentToProcess; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u HubSpotFormDefinitionLegalConsentOptionsUnionParam) GetPrivacyText() *string {
	if vt := u.OfLegitimateInterest; vt != nil {
		return (*string)(&vt.PrivacyText)
	} else if vt := u.OfExplicitConsentToProcess; vt != nil {
		return (*string)(&vt.PrivacyText)
	} else if vt := u.OfImplicitConsentToProcess; vt != nil {
		return (*string)(&vt.PrivacyText)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u HubSpotFormDefinitionLegalConsentOptionsUnionParam) GetCommunicationConsentText() *string {
	if vt := u.OfExplicitConsentToProcess; vt != nil && vt.CommunicationConsentText.Valid() {
		return &vt.CommunicationConsentText.Value
	} else if vt := u.OfImplicitConsentToProcess; vt != nil && vt.CommunicationConsentText.Valid() {
		return &vt.CommunicationConsentText.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u HubSpotFormDefinitionLegalConsentOptionsUnionParam) GetConsentToProcessText() *string {
	if vt := u.OfExplicitConsentToProcess; vt != nil && vt.ConsentToProcessText.Valid() {
		return &vt.ConsentToProcessText.Value
	} else if vt := u.OfImplicitConsentToProcess; vt != nil && vt.ConsentToProcessText.Valid() {
		return &vt.ConsentToProcessText.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's CommunicationsCheckboxes property,
// if present.
func (u HubSpotFormDefinitionLegalConsentOptionsUnionParam) GetCommunicationsCheckboxes() []LegalConsentCheckboxParam {
	if vt := u.OfExplicitConsentToProcess; vt != nil {
		return vt.CommunicationsCheckboxes
	} else if vt := u.OfImplicitConsentToProcess; vt != nil {
		return vt.CommunicationsCheckboxes
	}
	return nil
}

type HubSpotFormDefinitionPatchRequestParam struct {
	// Whether this form is archived.
	Archived param.Opt[bool] `json:"archived,omitzero"`
	// The name of the form. Expected to be unique for a hub.
	Name          param.Opt[string]             `json:"name,omitzero"`
	Configuration HubSpotFormConfigurationParam `json:"configuration,omitzero"`
	// Options for styling the form.
	DisplayOptions FormDisplayOptionsParam `json:"displayOptions,omitzero"`
	// The fields in the form, grouped in rows.
	FieldGroups         []FieldGroupParam                                              `json:"fieldGroups,omitzero"`
	LegalConsentOptions HubSpotFormDefinitionPatchRequestLegalConsentOptionsUnionParam `json:"legalConsentOptions,omitzero"`
	paramObj
}

func (r HubSpotFormDefinitionPatchRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow HubSpotFormDefinitionPatchRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *HubSpotFormDefinitionPatchRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type HubSpotFormDefinitionPatchRequestLegalConsentOptionsUnionParam struct {
	OfNone                     *LegalConsentOptionsNoneParam                     `json:",omitzero,inline"`
	OfLegitimateInterest       *LegalConsentOptionsLegitimateInterestParam       `json:",omitzero,inline"`
	OfExplicitConsentToProcess *LegalConsentOptionsExplicitConsentToProcessParam `json:",omitzero,inline"`
	OfImplicitConsentToProcess *LegalConsentOptionsImplicitConsentToProcessParam `json:",omitzero,inline"`
	paramUnion
}

func (u HubSpotFormDefinitionPatchRequestLegalConsentOptionsUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfNone, u.OfLegitimateInterest, u.OfExplicitConsentToProcess, u.OfImplicitConsentToProcess)
}
func (u *HubSpotFormDefinitionPatchRequestLegalConsentOptionsUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *HubSpotFormDefinitionPatchRequestLegalConsentOptionsUnionParam) asAny() any {
	if !param.IsOmitted(u.OfNone) {
		return u.OfNone
	} else if !param.IsOmitted(u.OfLegitimateInterest) {
		return u.OfLegitimateInterest
	} else if !param.IsOmitted(u.OfExplicitConsentToProcess) {
		return u.OfExplicitConsentToProcess
	} else if !param.IsOmitted(u.OfImplicitConsentToProcess) {
		return u.OfImplicitConsentToProcess
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u HubSpotFormDefinitionPatchRequestLegalConsentOptionsUnionParam) GetLawfulBasis() *string {
	if vt := u.OfLegitimateInterest; vt != nil {
		return (*string)(&vt.LawfulBasis)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u HubSpotFormDefinitionPatchRequestLegalConsentOptionsUnionParam) GetSubscriptionTypeIDs() []int64 {
	if vt := u.OfLegitimateInterest; vt != nil {
		return vt.SubscriptionTypeIDs
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u HubSpotFormDefinitionPatchRequestLegalConsentOptionsUnionParam) GetConsentToProcessCheckboxLabel() *string {
	if vt := u.OfExplicitConsentToProcess; vt != nil && vt.ConsentToProcessCheckboxLabel.Valid() {
		return &vt.ConsentToProcessCheckboxLabel.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u HubSpotFormDefinitionPatchRequestLegalConsentOptionsUnionParam) GetConsentToProcessFooterText() *string {
	if vt := u.OfExplicitConsentToProcess; vt != nil && vt.ConsentToProcessFooterText.Valid() {
		return &vt.ConsentToProcessFooterText.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u HubSpotFormDefinitionPatchRequestLegalConsentOptionsUnionParam) GetType() *string {
	if vt := u.OfNone; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfLegitimateInterest; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfExplicitConsentToProcess; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfImplicitConsentToProcess; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u HubSpotFormDefinitionPatchRequestLegalConsentOptionsUnionParam) GetPrivacyText() *string {
	if vt := u.OfLegitimateInterest; vt != nil {
		return (*string)(&vt.PrivacyText)
	} else if vt := u.OfExplicitConsentToProcess; vt != nil {
		return (*string)(&vt.PrivacyText)
	} else if vt := u.OfImplicitConsentToProcess; vt != nil {
		return (*string)(&vt.PrivacyText)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u HubSpotFormDefinitionPatchRequestLegalConsentOptionsUnionParam) GetCommunicationConsentText() *string {
	if vt := u.OfExplicitConsentToProcess; vt != nil && vt.CommunicationConsentText.Valid() {
		return &vt.CommunicationConsentText.Value
	} else if vt := u.OfImplicitConsentToProcess; vt != nil && vt.CommunicationConsentText.Valid() {
		return &vt.CommunicationConsentText.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u HubSpotFormDefinitionPatchRequestLegalConsentOptionsUnionParam) GetConsentToProcessText() *string {
	if vt := u.OfExplicitConsentToProcess; vt != nil && vt.ConsentToProcessText.Valid() {
		return &vt.ConsentToProcessText.Value
	} else if vt := u.OfImplicitConsentToProcess; vt != nil && vt.ConsentToProcessText.Valid() {
		return &vt.ConsentToProcessText.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's CommunicationsCheckboxes property,
// if present.
func (u HubSpotFormDefinitionPatchRequestLegalConsentOptionsUnionParam) GetCommunicationsCheckboxes() []LegalConsentCheckboxParam {
	if vt := u.OfExplicitConsentToProcess; vt != nil {
		return vt.CommunicationsCheckboxes
	} else if vt := u.OfImplicitConsentToProcess; vt != nil {
		return vt.CommunicationsCheckboxes
	}
	return nil
}

type LegalConsentCheckbox struct {
	// The main label for the form field.
	Label string `json:"label,required"`
	// Whether this checkbox is required when submitting the form.
	Required           bool  `json:"required,required"`
	SubscriptionTypeID int64 `json:"subscriptionTypeId,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Label              respjson.Field
		Required           respjson.Field
		SubscriptionTypeID respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LegalConsentCheckbox) RawJSON() string { return r.JSON.raw }
func (r *LegalConsentCheckbox) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this LegalConsentCheckbox to a LegalConsentCheckboxParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// LegalConsentCheckboxParam.Overrides()
func (r LegalConsentCheckbox) ToParam() LegalConsentCheckboxParam {
	return param.Override[LegalConsentCheckboxParam](json.RawMessage(r.RawJSON()))
}

// The properties Label, Required, SubscriptionTypeID are required.
type LegalConsentCheckboxParam struct {
	// The main label for the form field.
	Label string `json:"label,required"`
	// Whether this checkbox is required when submitting the form.
	Required           bool  `json:"required,required"`
	SubscriptionTypeID int64 `json:"subscriptionTypeId,required"`
	paramObj
}

func (r LegalConsentCheckboxParam) MarshalJSON() (data []byte, err error) {
	type shadow LegalConsentCheckboxParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LegalConsentCheckboxParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LegalConsentOptionsExplicitConsentToProcess struct {
	CommunicationsCheckboxes []LegalConsentCheckbox `json:"communicationsCheckboxes,required"`
	PrivacyText              string                 `json:"privacyText,required"`
	// Any of "explicit_consent_to_process".
	Type                          LegalConsentOptionsExplicitConsentToProcessType `json:"type,required"`
	CommunicationConsentText      string                                          `json:"communicationConsentText"`
	ConsentToProcessCheckboxLabel string                                          `json:"consentToProcessCheckboxLabel"`
	ConsentToProcessFooterText    string                                          `json:"consentToProcessFooterText"`
	ConsentToProcessText          string                                          `json:"consentToProcessText"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CommunicationsCheckboxes      respjson.Field
		PrivacyText                   respjson.Field
		Type                          respjson.Field
		CommunicationConsentText      respjson.Field
		ConsentToProcessCheckboxLabel respjson.Field
		ConsentToProcessFooterText    respjson.Field
		ConsentToProcessText          respjson.Field
		ExtraFields                   map[string]respjson.Field
		raw                           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LegalConsentOptionsExplicitConsentToProcess) RawJSON() string { return r.JSON.raw }
func (r *LegalConsentOptionsExplicitConsentToProcess) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this LegalConsentOptionsExplicitConsentToProcess to a
// LegalConsentOptionsExplicitConsentToProcessParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// LegalConsentOptionsExplicitConsentToProcessParam.Overrides()
func (r LegalConsentOptionsExplicitConsentToProcess) ToParam() LegalConsentOptionsExplicitConsentToProcessParam {
	return param.Override[LegalConsentOptionsExplicitConsentToProcessParam](json.RawMessage(r.RawJSON()))
}

type LegalConsentOptionsExplicitConsentToProcessType string

const (
	LegalConsentOptionsExplicitConsentToProcessTypeExplicitConsentToProcess LegalConsentOptionsExplicitConsentToProcessType = "explicit_consent_to_process"
)

// The properties CommunicationsCheckboxes, PrivacyText, Type are required.
type LegalConsentOptionsExplicitConsentToProcessParam struct {
	CommunicationsCheckboxes []LegalConsentCheckboxParam `json:"communicationsCheckboxes,omitzero,required"`
	PrivacyText              string                      `json:"privacyText,required"`
	// Any of "explicit_consent_to_process".
	Type                          LegalConsentOptionsExplicitConsentToProcessType `json:"type,omitzero,required"`
	CommunicationConsentText      param.Opt[string]                               `json:"communicationConsentText,omitzero"`
	ConsentToProcessCheckboxLabel param.Opt[string]                               `json:"consentToProcessCheckboxLabel,omitzero"`
	ConsentToProcessFooterText    param.Opt[string]                               `json:"consentToProcessFooterText,omitzero"`
	ConsentToProcessText          param.Opt[string]                               `json:"consentToProcessText,omitzero"`
	paramObj
}

func (r LegalConsentOptionsExplicitConsentToProcessParam) MarshalJSON() (data []byte, err error) {
	type shadow LegalConsentOptionsExplicitConsentToProcessParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LegalConsentOptionsExplicitConsentToProcessParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LegalConsentOptionsImplicitConsentToProcess struct {
	CommunicationsCheckboxes []LegalConsentCheckbox `json:"communicationsCheckboxes,required"`
	PrivacyText              string                 `json:"privacyText,required"`
	// Any of "implicit_consent_to_process".
	Type                     LegalConsentOptionsImplicitConsentToProcessType `json:"type,required"`
	CommunicationConsentText string                                          `json:"communicationConsentText"`
	ConsentToProcessText     string                                          `json:"consentToProcessText"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CommunicationsCheckboxes respjson.Field
		PrivacyText              respjson.Field
		Type                     respjson.Field
		CommunicationConsentText respjson.Field
		ConsentToProcessText     respjson.Field
		ExtraFields              map[string]respjson.Field
		raw                      string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LegalConsentOptionsImplicitConsentToProcess) RawJSON() string { return r.JSON.raw }
func (r *LegalConsentOptionsImplicitConsentToProcess) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this LegalConsentOptionsImplicitConsentToProcess to a
// LegalConsentOptionsImplicitConsentToProcessParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// LegalConsentOptionsImplicitConsentToProcessParam.Overrides()
func (r LegalConsentOptionsImplicitConsentToProcess) ToParam() LegalConsentOptionsImplicitConsentToProcessParam {
	return param.Override[LegalConsentOptionsImplicitConsentToProcessParam](json.RawMessage(r.RawJSON()))
}

type LegalConsentOptionsImplicitConsentToProcessType string

const (
	LegalConsentOptionsImplicitConsentToProcessTypeImplicitConsentToProcess LegalConsentOptionsImplicitConsentToProcessType = "implicit_consent_to_process"
)

// The properties CommunicationsCheckboxes, PrivacyText, Type are required.
type LegalConsentOptionsImplicitConsentToProcessParam struct {
	CommunicationsCheckboxes []LegalConsentCheckboxParam `json:"communicationsCheckboxes,omitzero,required"`
	PrivacyText              string                      `json:"privacyText,required"`
	// Any of "implicit_consent_to_process".
	Type                     LegalConsentOptionsImplicitConsentToProcessType `json:"type,omitzero,required"`
	CommunicationConsentText param.Opt[string]                               `json:"communicationConsentText,omitzero"`
	ConsentToProcessText     param.Opt[string]                               `json:"consentToProcessText,omitzero"`
	paramObj
}

func (r LegalConsentOptionsImplicitConsentToProcessParam) MarshalJSON() (data []byte, err error) {
	type shadow LegalConsentOptionsImplicitConsentToProcessParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LegalConsentOptionsImplicitConsentToProcessParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LegalConsentOptionsLegitimateInterest struct {
	// Any of "lead", "client", "other".
	LawfulBasis         LegalConsentOptionsLegitimateInterestLawfulBasis `json:"lawfulBasis,required"`
	PrivacyText         string                                           `json:"privacyText,required"`
	SubscriptionTypeIDs []int64                                          `json:"subscriptionTypeIds,required"`
	// Any of "legitimate_interest".
	Type LegalConsentOptionsLegitimateInterestType `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		LawfulBasis         respjson.Field
		PrivacyText         respjson.Field
		SubscriptionTypeIDs respjson.Field
		Type                respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LegalConsentOptionsLegitimateInterest) RawJSON() string { return r.JSON.raw }
func (r *LegalConsentOptionsLegitimateInterest) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this LegalConsentOptionsLegitimateInterest to a
// LegalConsentOptionsLegitimateInterestParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// LegalConsentOptionsLegitimateInterestParam.Overrides()
func (r LegalConsentOptionsLegitimateInterest) ToParam() LegalConsentOptionsLegitimateInterestParam {
	return param.Override[LegalConsentOptionsLegitimateInterestParam](json.RawMessage(r.RawJSON()))
}

type LegalConsentOptionsLegitimateInterestLawfulBasis string

const (
	LegalConsentOptionsLegitimateInterestLawfulBasisLead   LegalConsentOptionsLegitimateInterestLawfulBasis = "lead"
	LegalConsentOptionsLegitimateInterestLawfulBasisClient LegalConsentOptionsLegitimateInterestLawfulBasis = "client"
	LegalConsentOptionsLegitimateInterestLawfulBasisOther  LegalConsentOptionsLegitimateInterestLawfulBasis = "other"
)

type LegalConsentOptionsLegitimateInterestType string

const (
	LegalConsentOptionsLegitimateInterestTypeLegitimateInterest LegalConsentOptionsLegitimateInterestType = "legitimate_interest"
)

// The properties LawfulBasis, PrivacyText, SubscriptionTypeIDs, Type are required.
type LegalConsentOptionsLegitimateInterestParam struct {
	// Any of "lead", "client", "other".
	LawfulBasis         LegalConsentOptionsLegitimateInterestLawfulBasis `json:"lawfulBasis,omitzero,required"`
	PrivacyText         string                                           `json:"privacyText,required"`
	SubscriptionTypeIDs []int64                                          `json:"subscriptionTypeIds,omitzero,required"`
	// Any of "legitimate_interest".
	Type LegalConsentOptionsLegitimateInterestType `json:"type,omitzero,required"`
	paramObj
}

func (r LegalConsentOptionsLegitimateInterestParam) MarshalJSON() (data []byte, err error) {
	type shadow LegalConsentOptionsLegitimateInterestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LegalConsentOptionsLegitimateInterestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LegalConsentOptionsNone struct {
	// Any of "none".
	Type LegalConsentOptionsNoneType `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LegalConsentOptionsNone) RawJSON() string { return r.JSON.raw }
func (r *LegalConsentOptionsNone) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this LegalConsentOptionsNone to a LegalConsentOptionsNoneParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// LegalConsentOptionsNoneParam.Overrides()
func (r LegalConsentOptionsNone) ToParam() LegalConsentOptionsNoneParam {
	return param.Override[LegalConsentOptionsNoneParam](json.RawMessage(r.RawJSON()))
}

type LegalConsentOptionsNoneType string

const (
	LegalConsentOptionsNoneTypeNone LegalConsentOptionsNoneType = "none"
)

// The property Type is required.
type LegalConsentOptionsNoneParam struct {
	// Any of "none".
	Type LegalConsentOptionsNoneType `json:"type,omitzero,required"`
	paramObj
}

func (r LegalConsentOptionsNoneParam) MarshalJSON() (data []byte, err error) {
	type shadow LegalConsentOptionsNoneParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LegalConsentOptionsNoneParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LifecycleStage struct {
	// The objectTypeId for both contact and company
	ObjectTypeID string `json:"objectTypeId,required"`
	// The internal name of the contact's lifecycle stage set when submitting a form
	Value string `json:"value,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ObjectTypeID respjson.Field
		Value        respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LifecycleStage) RawJSON() string { return r.JSON.raw }
func (r *LifecycleStage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this LifecycleStage to a LifecycleStageParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// LifecycleStageParam.Overrides()
func (r LifecycleStage) ToParam() LifecycleStageParam {
	return param.Override[LifecycleStageParam](json.RawMessage(r.RawJSON()))
}

// The properties ObjectTypeID, Value are required.
type LifecycleStageParam struct {
	// The objectTypeId for both contact and company
	ObjectTypeID string `json:"objectTypeId,required"`
	// The internal name of the contact's lifecycle stage set when submitting a form
	Value string `json:"value,required"`
	paramObj
}

func (r LifecycleStageParam) MarshalJSON() (data []byte, err error) {
	type shadow LifecycleStageParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LifecycleStageParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A form field used for collecting a mobile phone number.
type MobilePhoneField struct {
	// A list of other fields to make visible based on the value filled in for this
	// field.
	DependentFields []DependentField `json:"dependentFields,required"`
	// Determines how the field will be displayed and validated.
	//
	// Any of "mobile_phone".
	FieldType MobilePhoneFieldFieldType `json:"fieldType,required"`
	// Whether a field should be hidden or not. Hidden fields won't appear on the form,
	// but can be used to pass a value to a property without requiring the customer to
	// fill it in.
	Hidden bool `json:"hidden,required"`
	// The main label for the form field.
	Label string `json:"label,required"`
	// The identifier of the field. In combination with the object type ID, it must be
	// unique.
	Name string `json:"name,required"`
	// A unique ID for this field's CRM object type. For example a CONTACT field will
	// have the object type ID 0-1.
	ObjectTypeID string `json:"objectTypeId,required"`
	// Whether a value for this field is required when submitting the form.
	Required bool `json:"required,required"`
	// Describes how a phone number should be validated.
	Validation PhoneFieldValidation `json:"validation,required"`
	// The value filled in by default. This value will be submitted unless the customer
	// modifies it.
	DefaultValue string `json:"defaultValue"`
	// Additional text helping the customer to complete the field.
	Description string `json:"description"`
	// The prompt text showing when the field isn't filled in.
	Placeholder string `json:"placeholder"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DependentFields respjson.Field
		FieldType       respjson.Field
		Hidden          respjson.Field
		Label           respjson.Field
		Name            respjson.Field
		ObjectTypeID    respjson.Field
		Required        respjson.Field
		Validation      respjson.Field
		DefaultValue    respjson.Field
		Description     respjson.Field
		Placeholder     respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MobilePhoneField) RawJSON() string { return r.JSON.raw }
func (r *MobilePhoneField) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this MobilePhoneField to a MobilePhoneFieldParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// MobilePhoneFieldParam.Overrides()
func (r MobilePhoneField) ToParam() MobilePhoneFieldParam {
	return param.Override[MobilePhoneFieldParam](json.RawMessage(r.RawJSON()))
}

// Determines how the field will be displayed and validated.
type MobilePhoneFieldFieldType string

const (
	MobilePhoneFieldFieldTypeMobilePhone MobilePhoneFieldFieldType = "mobile_phone"
)

// A form field used for collecting a mobile phone number.
//
// The properties DependentFields, FieldType, Hidden, Label, Name, ObjectTypeID,
// Required, Validation are required.
type MobilePhoneFieldParam struct {
	// A list of other fields to make visible based on the value filled in for this
	// field.
	DependentFields []DependentFieldParam `json:"dependentFields,omitzero,required"`
	// Determines how the field will be displayed and validated.
	//
	// Any of "mobile_phone".
	FieldType MobilePhoneFieldFieldType `json:"fieldType,omitzero,required"`
	// Whether a field should be hidden or not. Hidden fields won't appear on the form,
	// but can be used to pass a value to a property without requiring the customer to
	// fill it in.
	Hidden bool `json:"hidden,required"`
	// The main label for the form field.
	Label string `json:"label,required"`
	// The identifier of the field. In combination with the object type ID, it must be
	// unique.
	Name string `json:"name,required"`
	// A unique ID for this field's CRM object type. For example a CONTACT field will
	// have the object type ID 0-1.
	ObjectTypeID string `json:"objectTypeId,required"`
	// Whether a value for this field is required when submitting the form.
	Required bool `json:"required,required"`
	// Describes how a phone number should be validated.
	Validation PhoneFieldValidationParam `json:"validation,omitzero,required"`
	// The value filled in by default. This value will be submitted unless the customer
	// modifies it.
	DefaultValue param.Opt[string] `json:"defaultValue,omitzero"`
	// Additional text helping the customer to complete the field.
	Description param.Opt[string] `json:"description,omitzero"`
	// The prompt text showing when the field isn't filled in.
	Placeholder param.Opt[string] `json:"placeholder,omitzero"`
	paramObj
}

func (r MobilePhoneFieldParam) MarshalJSON() (data []byte, err error) {
	type shadow MobilePhoneFieldParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MobilePhoneFieldParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A form field consisting of a multiple-line text box.
type MultiLineTextField struct {
	// A list of other fields to make visible based on the value filled in for this
	// field.
	DependentFields []DependentField `json:"dependentFields,required"`
	// Determines how the field will be displayed and validated.
	//
	// Any of "multi_line_text".
	FieldType MultiLineTextFieldFieldType `json:"fieldType,required"`
	// Whether a field should be hidden or not. Hidden fields won't appear on the form,
	// but can be used to pass a value to a property without requiring the customer to
	// fill it in.
	Hidden bool `json:"hidden,required"`
	// The main label for the form field.
	Label string `json:"label,required"`
	// The identifier of the field. In combination with the object type ID, it must be
	// unique.
	Name string `json:"name,required"`
	// A unique ID for this field's CRM object type. For example a CONTACT field will
	// have the object type ID 0-1.
	ObjectTypeID string `json:"objectTypeId,required"`
	// Whether a value for this field is required when submitting the form.
	Required bool `json:"required,required"`
	// The value filled in by default. This value will be submitted unless the customer
	// modifies it.
	DefaultValue string `json:"defaultValue"`
	// Additional text helping the customer to complete the field.
	Description string `json:"description"`
	// The prompt text showing when the field isn't filled in.
	Placeholder string `json:"placeholder"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DependentFields respjson.Field
		FieldType       respjson.Field
		Hidden          respjson.Field
		Label           respjson.Field
		Name            respjson.Field
		ObjectTypeID    respjson.Field
		Required        respjson.Field
		DefaultValue    respjson.Field
		Description     respjson.Field
		Placeholder     respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MultiLineTextField) RawJSON() string { return r.JSON.raw }
func (r *MultiLineTextField) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this MultiLineTextField to a MultiLineTextFieldParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// MultiLineTextFieldParam.Overrides()
func (r MultiLineTextField) ToParam() MultiLineTextFieldParam {
	return param.Override[MultiLineTextFieldParam](json.RawMessage(r.RawJSON()))
}

// Determines how the field will be displayed and validated.
type MultiLineTextFieldFieldType string

const (
	MultiLineTextFieldFieldTypeMultiLineText MultiLineTextFieldFieldType = "multi_line_text"
)

// A form field consisting of a multiple-line text box.
//
// The properties DependentFields, FieldType, Hidden, Label, Name, ObjectTypeID,
// Required are required.
type MultiLineTextFieldParam struct {
	// A list of other fields to make visible based on the value filled in for this
	// field.
	DependentFields []DependentFieldParam `json:"dependentFields,omitzero,required"`
	// Determines how the field will be displayed and validated.
	//
	// Any of "multi_line_text".
	FieldType MultiLineTextFieldFieldType `json:"fieldType,omitzero,required"`
	// Whether a field should be hidden or not. Hidden fields won't appear on the form,
	// but can be used to pass a value to a property without requiring the customer to
	// fill it in.
	Hidden bool `json:"hidden,required"`
	// The main label for the form field.
	Label string `json:"label,required"`
	// The identifier of the field. In combination with the object type ID, it must be
	// unique.
	Name string `json:"name,required"`
	// A unique ID for this field's CRM object type. For example a CONTACT field will
	// have the object type ID 0-1.
	ObjectTypeID string `json:"objectTypeId,required"`
	// Whether a value for this field is required when submitting the form.
	Required bool `json:"required,required"`
	// The value filled in by default. This value will be submitted unless the customer
	// modifies it.
	DefaultValue param.Opt[string] `json:"defaultValue,omitzero"`
	// Additional text helping the customer to complete the field.
	Description param.Opt[string] `json:"description,omitzero"`
	// The prompt text showing when the field isn't filled in.
	Placeholder param.Opt[string] `json:"placeholder,omitzero"`
	paramObj
}

func (r MultiLineTextFieldParam) MarshalJSON() (data []byte, err error) {
	type shadow MultiLineTextFieldParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MultiLineTextFieldParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A form field consisting of a set of checkboxes allowing multiple choices to be
// selected at one time.
type MultipleCheckboxesField struct {
	// The values selected by default. Those values will be submitted unless the
	// customer modifies them.
	DefaultValues []string `json:"defaultValues,required"`
	// A list of other fields to make visible based on the value filled in for this
	// field.
	DependentFields []DependentField `json:"dependentFields,required"`
	// Determines how the field will be displayed and validated.
	//
	// Any of "multiple_checkboxes".
	FieldType MultipleCheckboxesFieldFieldType `json:"fieldType,required"`
	// Whether a field should be hidden or not. Hidden fields won't appear on the form,
	// but can be used to pass a value to a property without requiring the customer to
	// fill it in.
	Hidden bool `json:"hidden,required"`
	// The main label for the form field.
	Label string `json:"label,required"`
	// The identifier of the field. In combination with the object type ID, it must be
	// unique.
	Name string `json:"name,required"`
	// A unique ID for this field's CRM object type. For example a CONTACT field will
	// have the object type ID 0-1.
	ObjectTypeID string `json:"objectTypeId,required"`
	// The list of available choices for this field.
	Options []EnumeratedFieldOption `json:"options,required"`
	// Whether a value for this field is required when submitting the form.
	Required bool `json:"required,required"`
	// Additional text helping the customer to complete the field.
	Description string `json:"description"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DefaultValues   respjson.Field
		DependentFields respjson.Field
		FieldType       respjson.Field
		Hidden          respjson.Field
		Label           respjson.Field
		Name            respjson.Field
		ObjectTypeID    respjson.Field
		Options         respjson.Field
		Required        respjson.Field
		Description     respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MultipleCheckboxesField) RawJSON() string { return r.JSON.raw }
func (r *MultipleCheckboxesField) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this MultipleCheckboxesField to a MultipleCheckboxesFieldParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// MultipleCheckboxesFieldParam.Overrides()
func (r MultipleCheckboxesField) ToParam() MultipleCheckboxesFieldParam {
	return param.Override[MultipleCheckboxesFieldParam](json.RawMessage(r.RawJSON()))
}

// Determines how the field will be displayed and validated.
type MultipleCheckboxesFieldFieldType string

const (
	MultipleCheckboxesFieldFieldTypeMultipleCheckboxes MultipleCheckboxesFieldFieldType = "multiple_checkboxes"
)

// A form field consisting of a set of checkboxes allowing multiple choices to be
// selected at one time.
//
// The properties DefaultValues, DependentFields, FieldType, Hidden, Label, Name,
// ObjectTypeID, Options, Required are required.
type MultipleCheckboxesFieldParam struct {
	// The values selected by default. Those values will be submitted unless the
	// customer modifies them.
	DefaultValues []string `json:"defaultValues,omitzero,required"`
	// A list of other fields to make visible based on the value filled in for this
	// field.
	DependentFields []DependentFieldParam `json:"dependentFields,omitzero,required"`
	// Determines how the field will be displayed and validated.
	//
	// Any of "multiple_checkboxes".
	FieldType MultipleCheckboxesFieldFieldType `json:"fieldType,omitzero,required"`
	// Whether a field should be hidden or not. Hidden fields won't appear on the form,
	// but can be used to pass a value to a property without requiring the customer to
	// fill it in.
	Hidden bool `json:"hidden,required"`
	// The main label for the form field.
	Label string `json:"label,required"`
	// The identifier of the field. In combination with the object type ID, it must be
	// unique.
	Name string `json:"name,required"`
	// A unique ID for this field's CRM object type. For example a CONTACT field will
	// have the object type ID 0-1.
	ObjectTypeID string `json:"objectTypeId,required"`
	// The list of available choices for this field.
	Options []EnumeratedFieldOptionParam `json:"options,omitzero,required"`
	// Whether a value for this field is required when submitting the form.
	Required bool `json:"required,required"`
	// Additional text helping the customer to complete the field.
	Description param.Opt[string] `json:"description,omitzero"`
	paramObj
}

func (r MultipleCheckboxesFieldParam) MarshalJSON() (data []byte, err error) {
	type shadow MultipleCheckboxesFieldParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MultipleCheckboxesFieldParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A form field used for collecting a numeric value.
type NumberField struct {
	// A list of other fields to make visible based on the value filled in for this
	// field.
	DependentFields []DependentField `json:"dependentFields,required"`
	// Determines how the field will be displayed and validated.
	//
	// Any of "number".
	FieldType NumberFieldFieldType `json:"fieldType,required"`
	// Whether a field should be hidden or not. Hidden fields won't appear on the form,
	// but can be used to pass a value to a property without requiring the customer to
	// fill it in.
	Hidden bool `json:"hidden,required"`
	// The main label for the form field.
	Label string `json:"label,required"`
	// The identifier of the field. In combination with the object type ID, it must be
	// unique.
	Name string `json:"name,required"`
	// A unique ID for this field's CRM object type. For example a CONTACT field will
	// have the object type ID 0-1.
	ObjectTypeID string `json:"objectTypeId,required"`
	// Whether a value for this field is required when submitting the form.
	Required bool `json:"required,required"`
	// The value filled in by default. This value will be submitted unless the customer
	// modifies it.
	DefaultValue string `json:"defaultValue"`
	// Additional text helping the customer to complete the field.
	Description string `json:"description"`
	// The prompt text showing when the field isn't filled in.
	Placeholder string `json:"placeholder"`
	// Describes how a numeric value should be validated.
	Validation NumberFieldValidation `json:"validation"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DependentFields respjson.Field
		FieldType       respjson.Field
		Hidden          respjson.Field
		Label           respjson.Field
		Name            respjson.Field
		ObjectTypeID    respjson.Field
		Required        respjson.Field
		DefaultValue    respjson.Field
		Description     respjson.Field
		Placeholder     respjson.Field
		Validation      respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NumberField) RawJSON() string { return r.JSON.raw }
func (r *NumberField) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this NumberField to a NumberFieldParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// NumberFieldParam.Overrides()
func (r NumberField) ToParam() NumberFieldParam {
	return param.Override[NumberFieldParam](json.RawMessage(r.RawJSON()))
}

// Determines how the field will be displayed and validated.
type NumberFieldFieldType string

const (
	NumberFieldFieldTypeNumber NumberFieldFieldType = "number"
)

// A form field used for collecting a numeric value.
//
// The properties DependentFields, FieldType, Hidden, Label, Name, ObjectTypeID,
// Required are required.
type NumberFieldParam struct {
	// A list of other fields to make visible based on the value filled in for this
	// field.
	DependentFields []DependentFieldParam `json:"dependentFields,omitzero,required"`
	// Determines how the field will be displayed and validated.
	//
	// Any of "number".
	FieldType NumberFieldFieldType `json:"fieldType,omitzero,required"`
	// Whether a field should be hidden or not. Hidden fields won't appear on the form,
	// but can be used to pass a value to a property without requiring the customer to
	// fill it in.
	Hidden bool `json:"hidden,required"`
	// The main label for the form field.
	Label string `json:"label,required"`
	// The identifier of the field. In combination with the object type ID, it must be
	// unique.
	Name string `json:"name,required"`
	// A unique ID for this field's CRM object type. For example a CONTACT field will
	// have the object type ID 0-1.
	ObjectTypeID string `json:"objectTypeId,required"`
	// Whether a value for this field is required when submitting the form.
	Required bool `json:"required,required"`
	// The value filled in by default. This value will be submitted unless the customer
	// modifies it.
	DefaultValue param.Opt[string] `json:"defaultValue,omitzero"`
	// Additional text helping the customer to complete the field.
	Description param.Opt[string] `json:"description,omitzero"`
	// The prompt text showing when the field isn't filled in.
	Placeholder param.Opt[string] `json:"placeholder,omitzero"`
	// Describes how a numeric value should be validated.
	Validation NumberFieldValidationParam `json:"validation,omitzero"`
	paramObj
}

func (r NumberFieldParam) MarshalJSON() (data []byte, err error) {
	type shadow NumberFieldParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *NumberFieldParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Describes how a numeric value should be validated.
type NumberFieldValidation struct {
	MaxAllowedDigits int64 `json:"maxAllowedDigits,required"`
	MinAllowedDigits int64 `json:"minAllowedDigits,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		MaxAllowedDigits respjson.Field
		MinAllowedDigits respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NumberFieldValidation) RawJSON() string { return r.JSON.raw }
func (r *NumberFieldValidation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this NumberFieldValidation to a NumberFieldValidationParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// NumberFieldValidationParam.Overrides()
func (r NumberFieldValidation) ToParam() NumberFieldValidationParam {
	return param.Override[NumberFieldValidationParam](json.RawMessage(r.RawJSON()))
}

// Describes how a numeric value should be validated.
//
// The properties MaxAllowedDigits, MinAllowedDigits are required.
type NumberFieldValidationParam struct {
	MaxAllowedDigits int64 `json:"maxAllowedDigits,required"`
	MinAllowedDigits int64 `json:"minAllowedDigits,required"`
	paramObj
}

func (r NumberFieldValidationParam) MarshalJSON() (data []byte, err error) {
	type shadow NumberFieldValidationParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *NumberFieldValidationParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaymentLinkRadioField struct {
	DefaultValues   []string         `json:"defaultValues,required"`
	DependentFields []DependentField `json:"dependentFields,required"`
	// Any of "payment_link_radio".
	FieldType    PaymentLinkRadioFieldFieldType `json:"fieldType,required"`
	Hidden       bool                           `json:"hidden,required"`
	Label        string                         `json:"label,required"`
	Name         string                         `json:"name,required"`
	ObjectTypeID string                         `json:"objectTypeId,required"`
	Options      []EnumeratedFieldOption        `json:"options,required"`
	Required     bool                           `json:"required,required"`
	Description  string                         `json:"description"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DefaultValues   respjson.Field
		DependentFields respjson.Field
		FieldType       respjson.Field
		Hidden          respjson.Field
		Label           respjson.Field
		Name            respjson.Field
		ObjectTypeID    respjson.Field
		Options         respjson.Field
		Required        respjson.Field
		Description     respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaymentLinkRadioField) RawJSON() string { return r.JSON.raw }
func (r *PaymentLinkRadioField) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this PaymentLinkRadioField to a PaymentLinkRadioFieldParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// PaymentLinkRadioFieldParam.Overrides()
func (r PaymentLinkRadioField) ToParam() PaymentLinkRadioFieldParam {
	return param.Override[PaymentLinkRadioFieldParam](json.RawMessage(r.RawJSON()))
}

type PaymentLinkRadioFieldFieldType string

const (
	PaymentLinkRadioFieldFieldTypePaymentLinkRadio PaymentLinkRadioFieldFieldType = "payment_link_radio"
)

// The properties DefaultValues, DependentFields, FieldType, Hidden, Label, Name,
// ObjectTypeID, Options, Required are required.
type PaymentLinkRadioFieldParam struct {
	DefaultValues   []string              `json:"defaultValues,omitzero,required"`
	DependentFields []DependentFieldParam `json:"dependentFields,omitzero,required"`
	// Any of "payment_link_radio".
	FieldType    PaymentLinkRadioFieldFieldType `json:"fieldType,omitzero,required"`
	Hidden       bool                           `json:"hidden,required"`
	Label        string                         `json:"label,required"`
	Name         string                         `json:"name,required"`
	ObjectTypeID string                         `json:"objectTypeId,required"`
	Options      []EnumeratedFieldOptionParam   `json:"options,omitzero,required"`
	Required     bool                           `json:"required,required"`
	Description  param.Opt[string]              `json:"description,omitzero"`
	paramObj
}

func (r PaymentLinkRadioFieldParam) MarshalJSON() (data []byte, err error) {
	type shadow PaymentLinkRadioFieldParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PaymentLinkRadioFieldParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A form field used for collecting a phone number.
type PhoneField struct {
	// A list of other fields to make visible based on the value filled in for this
	// field.
	DependentFields []DependentField `json:"dependentFields,required"`
	// Determines how the field will be displayed and validated.
	//
	// Any of "phone".
	FieldType PhoneFieldFieldType `json:"fieldType,required"`
	// Whether a field should be hidden or not. Hidden fields won't appear on the form,
	// but can be used to pass a value to a property without requiring the customer to
	// fill it in.
	Hidden bool `json:"hidden,required"`
	// The main label for the form field.
	Label string `json:"label,required"`
	// The identifier of the field. In combination with the object type ID, it must be
	// unique.
	Name string `json:"name,required"`
	// A unique ID for this field's CRM object type. For example a CONTACT field will
	// have the object type ID 0-1.
	ObjectTypeID string `json:"objectTypeId,required"`
	// Whether a value for this field is required when submitting the form.
	Required bool `json:"required,required"`
	// Whether to display a country code drop down next to the phone field.
	UseCountryCodeSelect bool `json:"useCountryCodeSelect,required"`
	// Describes how a phone number should be validated.
	Validation PhoneFieldValidation `json:"validation,required"`
	// The value filled in by default. This value will be submitted unless the customer
	// modifies it.
	DefaultValue string `json:"defaultValue"`
	// Additional text helping the customer to complete the field.
	Description string `json:"description"`
	// The prompt text showing when the field isn't filled in.
	Placeholder string `json:"placeholder"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DependentFields      respjson.Field
		FieldType            respjson.Field
		Hidden               respjson.Field
		Label                respjson.Field
		Name                 respjson.Field
		ObjectTypeID         respjson.Field
		Required             respjson.Field
		UseCountryCodeSelect respjson.Field
		Validation           respjson.Field
		DefaultValue         respjson.Field
		Description          respjson.Field
		Placeholder          respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PhoneField) RawJSON() string { return r.JSON.raw }
func (r *PhoneField) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this PhoneField to a PhoneFieldParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// PhoneFieldParam.Overrides()
func (r PhoneField) ToParam() PhoneFieldParam {
	return param.Override[PhoneFieldParam](json.RawMessage(r.RawJSON()))
}

// Determines how the field will be displayed and validated.
type PhoneFieldFieldType string

const (
	PhoneFieldFieldTypePhone PhoneFieldFieldType = "phone"
)

// A form field used for collecting a phone number.
//
// The properties DependentFields, FieldType, Hidden, Label, Name, ObjectTypeID,
// Required, UseCountryCodeSelect, Validation are required.
type PhoneFieldParam struct {
	// A list of other fields to make visible based on the value filled in for this
	// field.
	DependentFields []DependentFieldParam `json:"dependentFields,omitzero,required"`
	// Determines how the field will be displayed and validated.
	//
	// Any of "phone".
	FieldType PhoneFieldFieldType `json:"fieldType,omitzero,required"`
	// Whether a field should be hidden or not. Hidden fields won't appear on the form,
	// but can be used to pass a value to a property without requiring the customer to
	// fill it in.
	Hidden bool `json:"hidden,required"`
	// The main label for the form field.
	Label string `json:"label,required"`
	// The identifier of the field. In combination with the object type ID, it must be
	// unique.
	Name string `json:"name,required"`
	// A unique ID for this field's CRM object type. For example a CONTACT field will
	// have the object type ID 0-1.
	ObjectTypeID string `json:"objectTypeId,required"`
	// Whether a value for this field is required when submitting the form.
	Required bool `json:"required,required"`
	// Whether to display a country code drop down next to the phone field.
	UseCountryCodeSelect bool `json:"useCountryCodeSelect,required"`
	// Describes how a phone number should be validated.
	Validation PhoneFieldValidationParam `json:"validation,omitzero,required"`
	// The value filled in by default. This value will be submitted unless the customer
	// modifies it.
	DefaultValue param.Opt[string] `json:"defaultValue,omitzero"`
	// Additional text helping the customer to complete the field.
	Description param.Opt[string] `json:"description,omitzero"`
	// The prompt text showing when the field isn't filled in.
	Placeholder param.Opt[string] `json:"placeholder,omitzero"`
	paramObj
}

func (r PhoneFieldParam) MarshalJSON() (data []byte, err error) {
	type shadow PhoneFieldParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PhoneFieldParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Describes how a phone number should be validated.
type PhoneFieldValidation struct {
	MaxAllowedDigits int64 `json:"maxAllowedDigits,required"`
	MinAllowedDigits int64 `json:"minAllowedDigits,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		MaxAllowedDigits respjson.Field
		MinAllowedDigits respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PhoneFieldValidation) RawJSON() string { return r.JSON.raw }
func (r *PhoneFieldValidation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this PhoneFieldValidation to a PhoneFieldValidationParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// PhoneFieldValidationParam.Overrides()
func (r PhoneFieldValidation) ToParam() PhoneFieldValidationParam {
	return param.Override[PhoneFieldValidationParam](json.RawMessage(r.RawJSON()))
}

// Describes how a phone number should be validated.
//
// The properties MaxAllowedDigits, MinAllowedDigits are required.
type PhoneFieldValidationParam struct {
	MaxAllowedDigits int64 `json:"maxAllowedDigits,required"`
	MinAllowedDigits int64 `json:"minAllowedDigits,required"`
	paramObj
}

func (r PhoneFieldValidationParam) MarshalJSON() (data []byte, err error) {
	type shadow PhoneFieldValidationParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PhoneFieldValidationParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A form field consisting of a set of radio options, out of which one can be
// selected at a time.
type RadioField struct {
	// The values selected by default. Those values will be submitted unless the
	// customer modifies them.
	DefaultValues []string `json:"defaultValues,required"`
	// A list of other fields to make visible based on the value filled in for this
	// field.
	DependentFields []DependentField `json:"dependentFields,required"`
	// Determines how the field will be displayed and validated.
	//
	// Any of "radio".
	FieldType RadioFieldFieldType `json:"fieldType,required"`
	// Whether a field should be hidden or not. Hidden fields won't appear on the form,
	// but can be used to pass a value to a property without requiring the customer to
	// fill it in.
	Hidden bool `json:"hidden,required"`
	// The main label for the form field.
	Label string `json:"label,required"`
	// The identifier of the field. In combination with the object type ID, it must be
	// unique.
	Name string `json:"name,required"`
	// A unique ID for this field's CRM object type. For example a CONTACT field will
	// have the object type ID 0-1.
	ObjectTypeID string `json:"objectTypeId,required"`
	// The list of available choices for this field.
	Options []EnumeratedFieldOption `json:"options,required"`
	// Whether a value for this field is required when submitting the form.
	Required bool `json:"required,required"`
	// Additional text helping the customer to complete the field.
	Description string `json:"description"`
	// The prompt text showing when the field isn't filled in.
	Placeholder string `json:"placeholder"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DefaultValues   respjson.Field
		DependentFields respjson.Field
		FieldType       respjson.Field
		Hidden          respjson.Field
		Label           respjson.Field
		Name            respjson.Field
		ObjectTypeID    respjson.Field
		Options         respjson.Field
		Required        respjson.Field
		Description     respjson.Field
		Placeholder     respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RadioField) RawJSON() string { return r.JSON.raw }
func (r *RadioField) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this RadioField to a RadioFieldParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// RadioFieldParam.Overrides()
func (r RadioField) ToParam() RadioFieldParam {
	return param.Override[RadioFieldParam](json.RawMessage(r.RawJSON()))
}

// Determines how the field will be displayed and validated.
type RadioFieldFieldType string

const (
	RadioFieldFieldTypeRadio RadioFieldFieldType = "radio"
)

// A form field consisting of a set of radio options, out of which one can be
// selected at a time.
//
// The properties DefaultValues, DependentFields, FieldType, Hidden, Label, Name,
// ObjectTypeID, Options, Required are required.
type RadioFieldParam struct {
	// The values selected by default. Those values will be submitted unless the
	// customer modifies them.
	DefaultValues []string `json:"defaultValues,omitzero,required"`
	// A list of other fields to make visible based on the value filled in for this
	// field.
	DependentFields []DependentFieldParam `json:"dependentFields,omitzero,required"`
	// Determines how the field will be displayed and validated.
	//
	// Any of "radio".
	FieldType RadioFieldFieldType `json:"fieldType,omitzero,required"`
	// Whether a field should be hidden or not. Hidden fields won't appear on the form,
	// but can be used to pass a value to a property without requiring the customer to
	// fill it in.
	Hidden bool `json:"hidden,required"`
	// The main label for the form field.
	Label string `json:"label,required"`
	// The identifier of the field. In combination with the object type ID, it must be
	// unique.
	Name string `json:"name,required"`
	// A unique ID for this field's CRM object type. For example a CONTACT field will
	// have the object type ID 0-1.
	ObjectTypeID string `json:"objectTypeId,required"`
	// The list of available choices for this field.
	Options []EnumeratedFieldOptionParam `json:"options,omitzero,required"`
	// Whether a value for this field is required when submitting the form.
	Required bool `json:"required,required"`
	// Additional text helping the customer to complete the field.
	Description param.Opt[string] `json:"description,omitzero"`
	// The prompt text showing when the field isn't filled in.
	Placeholder param.Opt[string] `json:"placeholder,omitzero"`
	paramObj
}

func (r RadioFieldParam) MarshalJSON() (data []byte, err error) {
	type shadow RadioFieldParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RadioFieldParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A form field consisting of a single checkbox.
type SingleCheckboxField struct {
	// A list of other fields to make visible based on the value filled in for this
	// field.
	DependentFields []DependentField `json:"dependentFields,required"`
	// Determines how the field will be displayed and validated.
	//
	// Any of "single_checkbox".
	FieldType SingleCheckboxFieldFieldType `json:"fieldType,required"`
	// Whether a field should be hidden or not. Hidden fields won't appear on the form,
	// but can be used to pass a value to a property without requiring the customer to
	// fill it in.
	Hidden bool `json:"hidden,required"`
	// The main label for the form field.
	Label string `json:"label,required"`
	// The identifier of the field. In combination with the object type ID, it must be
	// unique.
	Name string `json:"name,required"`
	// A unique ID for this field's CRM object type. For example a CONTACT field will
	// have the object type ID 0-1.
	ObjectTypeID string `json:"objectTypeId,required"`
	// Whether a value for this field is required when submitting the form.
	Required bool `json:"required,required"`
	// The value filled in by default. This value will be submitted unless the customer
	// modifies it.
	DefaultValue string `json:"defaultValue"`
	// Additional text helping the customer to complete the field.
	Description string `json:"description"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DependentFields respjson.Field
		FieldType       respjson.Field
		Hidden          respjson.Field
		Label           respjson.Field
		Name            respjson.Field
		ObjectTypeID    respjson.Field
		Required        respjson.Field
		DefaultValue    respjson.Field
		Description     respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SingleCheckboxField) RawJSON() string { return r.JSON.raw }
func (r *SingleCheckboxField) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this SingleCheckboxField to a SingleCheckboxFieldParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// SingleCheckboxFieldParam.Overrides()
func (r SingleCheckboxField) ToParam() SingleCheckboxFieldParam {
	return param.Override[SingleCheckboxFieldParam](json.RawMessage(r.RawJSON()))
}

// Determines how the field will be displayed and validated.
type SingleCheckboxFieldFieldType string

const (
	SingleCheckboxFieldFieldTypeSingleCheckbox SingleCheckboxFieldFieldType = "single_checkbox"
)

// A form field consisting of a single checkbox.
//
// The properties DependentFields, FieldType, Hidden, Label, Name, ObjectTypeID,
// Required are required.
type SingleCheckboxFieldParam struct {
	// A list of other fields to make visible based on the value filled in for this
	// field.
	DependentFields []DependentFieldParam `json:"dependentFields,omitzero,required"`
	// Determines how the field will be displayed and validated.
	//
	// Any of "single_checkbox".
	FieldType SingleCheckboxFieldFieldType `json:"fieldType,omitzero,required"`
	// Whether a field should be hidden or not. Hidden fields won't appear on the form,
	// but can be used to pass a value to a property without requiring the customer to
	// fill it in.
	Hidden bool `json:"hidden,required"`
	// The main label for the form field.
	Label string `json:"label,required"`
	// The identifier of the field. In combination with the object type ID, it must be
	// unique.
	Name string `json:"name,required"`
	// A unique ID for this field's CRM object type. For example a CONTACT field will
	// have the object type ID 0-1.
	ObjectTypeID string `json:"objectTypeId,required"`
	// Whether a value for this field is required when submitting the form.
	Required bool `json:"required,required"`
	// The value filled in by default. This value will be submitted unless the customer
	// modifies it.
	DefaultValue param.Opt[string] `json:"defaultValue,omitzero"`
	// Additional text helping the customer to complete the field.
	Description param.Opt[string] `json:"description,omitzero"`
	paramObj
}

func (r SingleCheckboxFieldParam) MarshalJSON() (data []byte, err error) {
	type shadow SingleCheckboxFieldParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SingleCheckboxFieldParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A form field consisting of a single-line text box.
type SingleLineTextField struct {
	// A list of other fields to make visible based on the value filled in for this
	// field.
	DependentFields []DependentField `json:"dependentFields,required"`
	// Determines how the field will be displayed and validated.
	//
	// Any of "single_line_text".
	FieldType SingleLineTextFieldFieldType `json:"fieldType,required"`
	// Whether a field should be hidden or not. Hidden fields won't appear on the form,
	// but can be used to pass a value to a property without requiring the customer to
	// fill it in.
	Hidden bool `json:"hidden,required"`
	// The main label for the form field.
	Label string `json:"label,required"`
	// The identifier of the field. In combination with the object type ID, it must be
	// unique.
	Name string `json:"name,required"`
	// A unique ID for this field's CRM object type. For example a CONTACT field will
	// have the object type ID 0-1.
	ObjectTypeID string `json:"objectTypeId,required"`
	// Whether a value for this field is required when submitting the form.
	Required bool `json:"required,required"`
	// The value filled in by default. This value will be submitted unless the customer
	// modifies it.
	DefaultValue string `json:"defaultValue"`
	// Additional text helping the customer to complete the field.
	Description string `json:"description"`
	// The prompt text showing when the field isn't filled in.
	Placeholder string `json:"placeholder"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DependentFields respjson.Field
		FieldType       respjson.Field
		Hidden          respjson.Field
		Label           respjson.Field
		Name            respjson.Field
		ObjectTypeID    respjson.Field
		Required        respjson.Field
		DefaultValue    respjson.Field
		Description     respjson.Field
		Placeholder     respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SingleLineTextField) RawJSON() string { return r.JSON.raw }
func (r *SingleLineTextField) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this SingleLineTextField to a SingleLineTextFieldParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// SingleLineTextFieldParam.Overrides()
func (r SingleLineTextField) ToParam() SingleLineTextFieldParam {
	return param.Override[SingleLineTextFieldParam](json.RawMessage(r.RawJSON()))
}

// Determines how the field will be displayed and validated.
type SingleLineTextFieldFieldType string

const (
	SingleLineTextFieldFieldTypeSingleLineText SingleLineTextFieldFieldType = "single_line_text"
)

// A form field consisting of a single-line text box.
//
// The properties DependentFields, FieldType, Hidden, Label, Name, ObjectTypeID,
// Required are required.
type SingleLineTextFieldParam struct {
	// A list of other fields to make visible based on the value filled in for this
	// field.
	DependentFields []DependentFieldParam `json:"dependentFields,omitzero,required"`
	// Determines how the field will be displayed and validated.
	//
	// Any of "single_line_text".
	FieldType SingleLineTextFieldFieldType `json:"fieldType,omitzero,required"`
	// Whether a field should be hidden or not. Hidden fields won't appear on the form,
	// but can be used to pass a value to a property without requiring the customer to
	// fill it in.
	Hidden bool `json:"hidden,required"`
	// The main label for the form field.
	Label string `json:"label,required"`
	// The identifier of the field. In combination with the object type ID, it must be
	// unique.
	Name string `json:"name,required"`
	// A unique ID for this field's CRM object type. For example a CONTACT field will
	// have the object type ID 0-1.
	ObjectTypeID string `json:"objectTypeId,required"`
	// Whether a value for this field is required when submitting the form.
	Required bool `json:"required,required"`
	// The value filled in by default. This value will be submitted unless the customer
	// modifies it.
	DefaultValue param.Opt[string] `json:"defaultValue,omitzero"`
	// Additional text helping the customer to complete the field.
	Description param.Opt[string] `json:"description,omitzero"`
	// The prompt text showing when the field isn't filled in.
	Placeholder param.Opt[string] `json:"placeholder,omitzero"`
	paramObj
}

func (r SingleLineTextFieldParam) MarshalJSON() (data []byte, err error) {
	type shadow SingleLineTextFieldParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SingleLineTextFieldParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FormNewParams struct {
	FormDefinitionCreateRequestBase FormDefinitionCreateRequestBaseParam
	paramObj
}

func (r FormNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.FormDefinitionCreateRequestBase)
}
func (r *FormNewParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.FormDefinitionCreateRequestBase)
}

type FormUpdateParams struct {
	HubSpotFormDefinitionPatchRequest HubSpotFormDefinitionPatchRequestParam
	paramObj
}

func (r FormUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.HubSpotFormDefinitionPatchRequest)
}
func (r *FormUpdateParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.HubSpotFormDefinitionPatchRequest)
}

type FormListParams struct {
	// The paging cursor token of the last successfully read resource will be returned
	// as the `paging.next.after` JSON property of a paged response containing more
	// results.
	After param.Opt[string] `query:"after,omitzero" json:"-"`
	// Whether to return only results that have been archived.
	Archived param.Opt[bool] `query:"archived,omitzero" json:"-"`
	// The maximum number of results to display per page.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// The form types to be included in the results.
	//
	// Any of "hubspot", "captured", "flow", "blog_comment", "all".
	FormTypes []string `query:"formTypes,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [FormListParams]'s query parameters as `url.Values`.
func (r FormListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type FormGetParams struct {
	// Whether to return only results that have been archived.
	Archived param.Opt[bool] `query:"archived,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [FormGetParams]'s query parameters as `url.Values`.
func (r FormGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type FormReplaceParams struct {
	HubSpotFormDefinition HubSpotFormDefinitionParam
	paramObj
}

func (r FormReplaceParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.HubSpotFormDefinition)
}
func (r *FormReplaceParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.HubSpotFormDefinition)
}
