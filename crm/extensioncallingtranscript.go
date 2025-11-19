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

// ExtensionCallingTranscriptService contains methods and other services that help
// with interacting with the Hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewExtensionCallingTranscriptService] method instead.
type ExtensionCallingTranscriptService struct {
	Options []option.RequestOption
}

// NewExtensionCallingTranscriptService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewExtensionCallingTranscriptService(opts ...option.RequestOption) (r ExtensionCallingTranscriptService) {
	r = ExtensionCallingTranscriptService{}
	r.Options = opts
	return
}

func (r *ExtensionCallingTranscriptService) New(ctx context.Context, body ExtensionCallingTranscriptNewParams, opts ...option.RequestOption) (res *TranscriptCreateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "crm/v3/extensions/calling/transcripts"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

func (r *ExtensionCallingTranscriptService) Delete(ctx context.Context, transcriptID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if transcriptID == "" {
		err = errors.New("missing required transcriptId parameter")
		return
	}
	path := fmt.Sprintf("crm/v3/extensions/calling/transcripts/%s", transcriptID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

func (r *ExtensionCallingTranscriptService) Get(ctx context.Context, transcriptID string, opts ...option.RequestOption) (res *TranscriptResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if transcriptID == "" {
		err = errors.New("missing required transcriptId parameter")
		return
	}
	path := fmt.Sprintf("crm/v3/extensions/calling/transcripts/%s", transcriptID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type Speaker struct {
	ID    string `json:"id,required"`
	Name  string `json:"name,required"`
	Email string `json:"email"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Name        respjson.Field
		Email       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Speaker) RawJSON() string { return r.JSON.raw }
func (r *Speaker) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this Speaker to a SpeakerParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// SpeakerParam.Overrides()
func (r Speaker) ToParam() SpeakerParam {
	return param.Override[SpeakerParam](json.RawMessage(r.RawJSON()))
}

// The properties ID, Name are required.
type SpeakerParam struct {
	ID    string            `json:"id,required"`
	Name  string            `json:"name,required"`
	Email param.Opt[string] `json:"email,omitzero"`
	paramObj
}

func (r SpeakerParam) MarshalJSON() (data []byte, err error) {
	type shadow SpeakerParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SpeakerParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties EngagementID, TranscriptCreateUtterances are required.
type TranscriptCreateRequestParam struct {
	EngagementID               int64                            `json:"engagementId,required"`
	TranscriptCreateUtterances []TranscriptCreateUtteranceParam `json:"transcriptCreateUtterances,omitzero,required"`
	paramObj
}

func (r TranscriptCreateRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow TranscriptCreateRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TranscriptCreateRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TranscriptCreateResponse struct {
	ID string `json:"id,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TranscriptCreateResponse) RawJSON() string { return r.JSON.raw }
func (r *TranscriptCreateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties EndTimeMillis, Speaker, StartTimeMillis, Text are required.
type TranscriptCreateUtteranceParam struct {
	EndTimeMillis   int64             `json:"endTimeMillis,required"`
	Speaker         SpeakerParam      `json:"speaker,omitzero,required"`
	StartTimeMillis int64             `json:"startTimeMillis,required"`
	Text            string            `json:"text,required"`
	LanguageCode    param.Opt[string] `json:"languageCode,omitzero"`
	paramObj
}

func (r TranscriptCreateUtteranceParam) MarshalJSON() (data []byte, err error) {
	type shadow TranscriptCreateUtteranceParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TranscriptCreateUtteranceParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TranscriptResponse struct {
	ID           string    `json:"id,required"`
	CreatedAt    time.Time `json:"createdAt,required" format:"date-time"`
	EngagementID int64     `json:"engagementId,required"`
	// Any of "HUBSPOT_GENERATED", "INTEGRATOR_GENERATED".
	TranscriptSource     TranscriptResponseTranscriptSource `json:"transcriptSource,required"`
	TranscriptUtterances []TranscriptUtterance              `json:"transcriptUtterances,required"`
	UpdatedAt            time.Time                          `json:"updatedAt,required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                   respjson.Field
		CreatedAt            respjson.Field
		EngagementID         respjson.Field
		TranscriptSource     respjson.Field
		TranscriptUtterances respjson.Field
		UpdatedAt            respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TranscriptResponse) RawJSON() string { return r.JSON.raw }
func (r *TranscriptResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TranscriptResponseTranscriptSource string

const (
	TranscriptResponseTranscriptSourceHubspotGenerated    TranscriptResponseTranscriptSource = "HUBSPOT_GENERATED"
	TranscriptResponseTranscriptSourceIntegratorGenerated TranscriptResponseTranscriptSource = "INTEGRATOR_GENERATED"
)

type TranscriptUtterance struct {
	ID              string  `json:"id,required"`
	EndTimeMillis   int64   `json:"endTimeMillis,required"`
	StartTimeMillis int64   `json:"startTimeMillis,required"`
	Text            string  `json:"text,required"`
	LanguageCode    string  `json:"languageCode"`
	Speaker         Speaker `json:"speaker"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		EndTimeMillis   respjson.Field
		StartTimeMillis respjson.Field
		Text            respjson.Field
		LanguageCode    respjson.Field
		Speaker         respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TranscriptUtterance) RawJSON() string { return r.JSON.raw }
func (r *TranscriptUtterance) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExtensionCallingTranscriptNewParams struct {
	TranscriptCreateRequest TranscriptCreateRequestParam
	paramObj
}

func (r ExtensionCallingTranscriptNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.TranscriptCreateRequest)
}
func (r *ExtensionCallingTranscriptNewParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.TranscriptCreateRequest)
}
