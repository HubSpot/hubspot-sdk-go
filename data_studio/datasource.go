// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package data_studio

import (
	"bytes"
	"context"
	"fmt"
	"mime/multipart"
	"net/http"
	"slices"
	"time"

	"github.com/stainless-sdks/hubspot-sdk-go/internal/apiform"
	"github.com/stainless-sdks/hubspot-sdk-go/internal/apijson"
	"github.com/stainless-sdks/hubspot-sdk-go/internal/requestconfig"
	"github.com/stainless-sdks/hubspot-sdk-go/option"
	"github.com/stainless-sdks/hubspot-sdk-go/packages/param"
	"github.com/stainless-sdks/hubspot-sdk-go/packages/respjson"
)

// DatasourceService contains methods and other services that help with interacting
// with the hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDatasourceService] method instead.
type DatasourceService struct {
	Options []option.RequestOption
}

// NewDatasourceService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewDatasourceService(opts ...option.RequestOption) (r DatasourceService) {
	r = DatasourceService{}
	r.Options = opts
	return
}

func (r *DatasourceService) New(ctx context.Context, body DatasourceNewParams, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "data-studio/2026-03/data-source"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

func (r *DatasourceService) Update(ctx context.Context, datasourceID int64, body DatasourceUpdateParams, opts ...option.RequestOption) (res *DataSourceUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("data-studio/2026-03/data-source/%v", datasourceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return res, err
}

func (r *DatasourceService) Delete(ctx context.Context, datasourceID int64, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := fmt.Sprintf("data-studio/2026-03/data-source/%v", datasourceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

func (r *DatasourceService) Get(ctx context.Context, datasourceID int64, opts ...option.RequestOption) (res *DataSourceGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("data-studio/2026-03/data-source/%v", datasourceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// The properties ContentDisposition, Entity, Headers, MediaType,
// MessageBodyWorkers, ParameterizedHeaders, Providers are required.
type BodyPartParam struct {
	ContentDisposition ContentDispositionParam `json:"contentDisposition,omitzero" api:"required"`
	// An object representing the actual content or payload of the body part.
	Entity any `json:"entity,omitzero" api:"required"`
	// An object containing the headers associated with this body part, where each
	// header can have multiple string values.
	Headers   map[string][]string `json:"headers,omitzero" api:"required"`
	MediaType MediaTypeParam      `json:"mediaType,omitzero" api:"required"`
	// An object representing workers that handle the processing of the message body.
	MessageBodyWorkers any `json:"messageBodyWorkers,omitzero" api:"required"`
	// An object containing headers with parameters, where each header can have
	// multiple ParameterizedHeader objects.
	ParameterizedHeaders map[string][]ParameterizedHeaderParam `json:"parameterizedHeaders,omitzero" api:"required"`
	// An object representing providers that supply additional handling or processing
	// for the body part.
	Providers any            `json:"providers,omitzero" api:"required"`
	Parent    MultiPartParam `json:"parent,omitzero"`
	paramObj
}

func (r BodyPartParam) MarshalJSON() (data []byte, err error) {
	type shadow BodyPartParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BodyPartParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties CreationDate, FileName, ModificationDate, Parameters, ReadDate,
// Size, Type are required.
type ContentDispositionParam struct {
	// The date and time when the file was created, formatted as a date-time string.
	CreationDate time.Time `json:"creationDate" api:"required" format:"date-time"`
	// The name of the file as a string, indicating the file's name in the content
	// disposition.
	FileName string `json:"fileName" api:"required"`
	// The date and time when the file was last modified, formatted as a date-time
	// string.
	ModificationDate time.Time `json:"modificationDate" api:"required" format:"date-time"`
	// An object containing additional parameters for the content disposition, with
	// each parameter represented as a key-value pair of strings.
	Parameters map[string]string `json:"parameters,omitzero" api:"required"`
	// The date and time when the file was last read, formatted as a date-time string.
	ReadDate time.Time `json:"readDate" api:"required" format:"date-time"`
	// The size of the file as an integer, representing the file's size in bytes.
	Size int64 `json:"size" api:"required"`
	// The type of content disposition, typically a string indicating how the content
	// should be handled.
	Type string `json:"type" api:"required"`
	paramObj
}

func (r ContentDispositionParam) MarshalJSON() (data []byte, err error) {
	type shadow ContentDispositionParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ContentDispositionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DataSourceGetResponse struct {
	// An array of FileColumn objects representing the columns in the data source.
	Columns []FileColumn `json:"columns" api:"required"`
	// The creation date and time of the data source, represented as a string.
	CreatedAt string `json:"createdAt" api:"required"`
	// The unique identifier for the data source, represented as a 64-bit integer.
	DatasourceID int64 `json:"datasourceId" api:"required"`
	// The name of the data source, represented as a string.
	DatasourceName string `json:"datasourceName" api:"required"`
	// The type of the data source, which is a string with a valid value of 'FILE'.
	//
	// Any of "FILE".
	DatasourceType DataSourceGetResponseDatasourceType `json:"datasourceType" api:"required"`
	// The status of the last data ingestion process, represented as a string. Valid
	// values include 'SUCCESSFUL', 'IN_PROGRESS', and 'FAILED'.
	//
	// Any of "FAILED", "IN_PROGRESS", "SUCCESSFUL".
	LastIngestionStatus DataSourceGetResponseLastIngestionStatus `json:"lastIngestionStatus" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Columns             respjson.Field
		CreatedAt           respjson.Field
		DatasourceID        respjson.Field
		DatasourceName      respjson.Field
		DatasourceType      respjson.Field
		LastIngestionStatus respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DataSourceGetResponse) RawJSON() string { return r.JSON.raw }
func (r *DataSourceGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The type of the data source, which is a string with a valid value of 'FILE'.
type DataSourceGetResponseDatasourceType string

const (
	DataSourceGetResponseDatasourceTypeFile DataSourceGetResponseDatasourceType = "FILE"
)

// The status of the last data ingestion process, represented as a string. Valid
// values include 'SUCCESSFUL', 'IN_PROGRESS', and 'FAILED'.
type DataSourceGetResponseLastIngestionStatus string

const (
	DataSourceGetResponseLastIngestionStatusFailed     DataSourceGetResponseLastIngestionStatus = "FAILED"
	DataSourceGetResponseLastIngestionStatusInProgress DataSourceGetResponseLastIngestionStatus = "IN_PROGRESS"
	DataSourceGetResponseLastIngestionStatusSuccessful DataSourceGetResponseLastIngestionStatus = "SUCCESSFUL"
)

type DataSourceUpdateResponse struct {
	// The unique identifier for the data source. It is an integer formatted as int64.
	DatasourceID int64 `json:"datasourceId" api:"required"`
	// The name of the data source. It is a string.
	DatasourceName string `json:"datasourceName" api:"required"`
	// A URL string that provides a preview link for the data source.
	PreviewLink string `json:"previewLink" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DatasourceID   respjson.Field
		DatasourceName respjson.Field
		PreviewLink    respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DataSourceUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *DataSourceUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FileColumn struct {
	// The name of the column, represented as a string.
	Name string `json:"name" api:"required"`
	// The data type of the column, represented as a string.
	Type string `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FileColumn) RawJSON() string { return r.JSON.raw }
func (r *FileColumn) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties ContentDisposition, Entity, FormDataContentDisposition, Headers,
// MediaType, MessageBodyWorkers, Name, ParameterizedHeaders, Providers, Simple,
// Value are required.
type FormDataBodyPartParam struct {
	ContentDisposition ContentDispositionParam `json:"contentDisposition,omitzero" api:"required"`
	// An object representing the entity of the form data part, which contains the
	// actual data being submitted.
	Entity                     any                             `json:"entity,omitzero" api:"required"`
	FormDataContentDisposition FormDataContentDispositionParam `json:"formDataContentDisposition,omitzero" api:"required"`
	// An object containing the headers associated with this form data part, where each
	// header can have multiple string values.
	Headers   map[string][]string `json:"headers,omitzero" api:"required"`
	MediaType MediaTypeParam      `json:"mediaType,omitzero" api:"required"`
	// An object representing the message body workers, which are responsible for
	// processing the body of the message.
	MessageBodyWorkers any `json:"messageBodyWorkers,omitzero" api:"required"`
	// The name of the form data part, typically used to identify the part within the
	// multipart request.
	Name string `json:"name" api:"required"`
	// An object containing parameterized headers, where each header can have multiple
	// values represented as ParameterizedHeader objects.
	ParameterizedHeaders map[string][]ParameterizedHeaderParam `json:"parameterizedHeaders,omitzero" api:"required"`
	// An object representing the providers associated with this form data part.
	Providers any `json:"providers,omitzero" api:"required"`
	// A boolean indicating whether the form data part is simple, typically meaning it
	// does not contain complex nested structures.
	Simple bool `json:"simple" api:"required"`
	// The string value of the form data part, representing the actual data being
	// submitted as a string.
	Value  string         `json:"value" api:"required"`
	Parent MultiPartParam `json:"parent,omitzero"`
	paramObj
}

func (r FormDataBodyPartParam) MarshalJSON() (data []byte, err error) {
	type shadow FormDataBodyPartParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FormDataBodyPartParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties CreationDate, FileName, ModificationDate, Name, Parameters,
// ReadDate, Size, Type are required.
type FormDataContentDispositionParam struct {
	// The date and time when the file was created, in ISO 8601 format.
	CreationDate time.Time `json:"creationDate" api:"required" format:"date-time"`
	// A string indicating the name of the file associated with this content
	// disposition.
	FileName string `json:"fileName" api:"required"`
	// The date and time when the file was last modified, in ISO 8601 format.
	ModificationDate time.Time `json:"modificationDate" api:"required" format:"date-time"`
	// A string representing the name associated with this content disposition.
	Name string `json:"name" api:"required"`
	// An object containing additional parameters for the content disposition, with
	// each parameter represented as a string.
	Parameters map[string]string `json:"parameters,omitzero" api:"required"`
	// The date and time when the file was last read, in ISO 8601 format.
	ReadDate time.Time `json:"readDate" api:"required" format:"date-time"`
	// An integer representing the size of the file in bytes.
	Size int64 `json:"size" api:"required"`
	// A string representing the type of content disposition.
	Type string `json:"type" api:"required"`
	paramObj
}

func (r FormDataContentDispositionParam) MarshalJSON() (data []byte, err error) {
	type shadow FormDataContentDispositionParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FormDataContentDispositionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties BodyParts, ContentDisposition, Entity, Fields, Headers,
// MediaType, MessageBodyWorkers, ParameterizedHeaders, Providers are required.
type FormDataMultiPartParam struct {
	// An array of BodyPart objects, each representing a part of the multipart form
	// data.
	BodyParts          []BodyPartParam         `json:"bodyParts,omitzero" api:"required"`
	ContentDisposition ContentDispositionParam `json:"contentDisposition,omitzero" api:"required"`
	// An object representing the entity of the multipart form data, containing the
	// actual data to be processed.
	Entity any `json:"entity,omitzero" api:"required"`
	// An object containing fields of the multipart form data, where each field can
	// have multiple FormDataBodyPart items.
	Fields map[string][]FormDataBodyPartParam `json:"fields,omitzero" api:"required"`
	// An object containing headers associated with the multipart form data, where each
	// header can have multiple string values.
	Headers   map[string][]string `json:"headers,omitzero" api:"required"`
	MediaType MediaTypeParam      `json:"mediaType,omitzero" api:"required"`
	// An object representing workers that process the message body of the multipart
	// form data.
	MessageBodyWorkers any `json:"messageBodyWorkers,omitzero" api:"required"`
	// An object containing parameterized headers, where each header can have multiple
	// ParameterizedHeader items.
	ParameterizedHeaders map[string][]ParameterizedHeaderParam `json:"parameterizedHeaders,omitzero" api:"required"`
	// An object representing providers associated with the multipart form data.
	Providers any            `json:"providers,omitzero" api:"required"`
	Parent    MultiPartParam `json:"parent,omitzero"`
	paramObj
}

func (r FormDataMultiPartParam) MarshalJSON() (data []byte, err error) {
	type shadow FormDataMultiPartParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FormDataMultiPartParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Parameters, Subtype, Type, WildcardSubtype, WildcardType are
// required.
type MediaTypeParam struct {
	// An object containing additional parameters for the media type, where each
	// key-value pair is a string.
	Parameters map[string]string `json:"parameters,omitzero" api:"required"`
	// The specific subtype of the media, represented as a string.
	Subtype string `json:"subtype" api:"required"`
	// The primary type of the media, represented as a string.
	Type string `json:"type" api:"required"`
	// A boolean indicating whether the media subtype is a wildcard.
	WildcardSubtype bool `json:"wildcardSubtype" api:"required"`
	// A boolean indicating whether the media type is a wildcard.
	WildcardType bool `json:"wildcardType" api:"required"`
	paramObj
}

func (r MediaTypeParam) MarshalJSON() (data []byte, err error) {
	type shadow MediaTypeParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MediaTypeParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties BodyParts, ContentDisposition, Entity, Headers, MediaType,
// MessageBodyWorkers, ParameterizedHeaders, Providers are required.
type MultiPartParam struct {
	// An array of BodyPart objects, each representing a distinct part of the multipart
	// entity.
	BodyParts          []BodyPartParam         `json:"bodyParts,omitzero" api:"required"`
	ContentDisposition ContentDispositionParam `json:"contentDisposition,omitzero" api:"required"`
	// An object that holds the main content or payload of the multipart entity.
	Entity any `json:"entity,omitzero" api:"required"`
	// An object containing a map of header names to their respective values, where
	// each value is an array of strings.
	Headers   map[string][]string `json:"headers,omitzero" api:"required"`
	MediaType MediaTypeParam      `json:"mediaType,omitzero" api:"required"`
	// An object that may contain workers for processing the message body, though its
	// specific properties are not detailed.
	MessageBodyWorkers any `json:"messageBodyWorkers,omitzero" api:"required"`
	// An object containing a map of header names to arrays of ParameterizedHeader
	// objects, which include additional parameters for each header.
	ParameterizedHeaders map[string][]ParameterizedHeaderParam `json:"parameterizedHeaders,omitzero" api:"required"`
	// An object that may contain providers related to the multipart entity, though its
	// specific properties are not detailed.
	Providers any             `json:"providers,omitzero" api:"required"`
	Parent    *MultiPartParam `json:"parent,omitzero"`
	paramObj
}

func (r MultiPartParam) MarshalJSON() (data []byte, err error) {
	type shadow MultiPartParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MultiPartParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Parameters, Value are required.
type ParameterizedHeaderParam struct {
	// An object containing additional parameters for the header, where each key is a
	// parameter name and each value is a string representing the parameter's value.
	Parameters map[string]string `json:"parameters,omitzero" api:"required"`
	// A string representing the main value of the header.
	Value string `json:"value" api:"required"`
	paramObj
}

func (r ParameterizedHeaderParam) MarshalJSON() (data []byte, err error) {
	type shadow ParameterizedHeaderParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ParameterizedHeaderParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DatasourceNewParams struct {
	FormDataMultiPart FormDataMultiPartParam
	paramObj
}

func (r DatasourceNewParams) MarshalMultipart() (data []byte, contentType string, err error) {
	buf := bytes.NewBuffer(nil)
	writer := multipart.NewWriter(buf)
	err = apiform.MarshalRoot(r.FormDataMultiPart, writer)
	if err == nil {
		err = apiform.WriteExtras(writer, r.ExtraFields())
	}
	if err != nil {
		writer.Close()
		return nil, "", err
	}
	err = writer.Close()
	if err != nil {
		return nil, "", err
	}
	return buf.Bytes(), writer.FormDataContentType(), nil
}

type DatasourceUpdateParams struct {
	FormDataMultiPart FormDataMultiPartParam
	paramObj
}

func (r DatasourceUpdateParams) MarshalMultipart() (data []byte, contentType string, err error) {
	buf := bytes.NewBuffer(nil)
	writer := multipart.NewWriter(buf)
	err = apiform.MarshalRoot(r.FormDataMultiPart, writer)
	if err == nil {
		err = apiform.WriteExtras(writer, r.ExtraFields())
	}
	if err != nil {
		writer.Close()
		return nil, "", err
	}
	err = writer.Close()
	if err != nil {
		return nil, "", err
	}
	return buf.Bytes(), writer.FormDataContentType(), nil
}
