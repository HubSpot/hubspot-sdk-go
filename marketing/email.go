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

// EmailService contains methods and other services that help with interacting with
// the hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEmailService] method instead.
type EmailService struct {
	options []option.RequestOption
}

// NewEmailService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewEmailService(opts ...option.RequestOption) (r EmailService) {
	r = EmailService{}
	r.options = opts
	return
}

func (r *EmailService) New(ctx context.Context, body EmailNewParams, opts ...option.RequestOption) (res *PublicEmail, err error) {
	opts = slices.Concat(r.options, opts)
	path := "marketing/emails/2026-03"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Change properties of a marketing email.
func (r *EmailService) Update(ctx context.Context, emailID string, params EmailUpdateParams, opts ...option.RequestOption) (res *PublicEmail, err error) {
	opts = slices.Concat(r.options, opts)
	if emailID == "" {
		err = errors.New("missing required emailId parameter")
		return nil, err
	}
	path := fmt.Sprintf("marketing/emails/2026-03/%s", url.PathEscape(emailID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return res, err
}

func (r *EmailService) List(ctx context.Context, query EmailListParams, opts ...option.RequestOption) (res *pagination.Page[PublicEmail], err error) {
	var raw *http.Response
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "marketing/emails/2026-03"
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

func (r *EmailService) ListAutoPaging(ctx context.Context, query EmailListParams, opts ...option.RequestOption) *pagination.PageAutoPager[PublicEmail] {
	return pagination.NewPageAutoPager(r.List(ctx, query, opts...))
}

// Delete a marketing email by its ID
func (r *EmailService) Delete(ctx context.Context, emailID string, body EmailDeleteParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if emailID == "" {
		err = errors.New("missing required emailId parameter")
		return err
	}
	path := fmt.Sprintf("marketing/emails/2026-03/%s", url.PathEscape(emailID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, nil, opts...)
	return err
}

// This will create a duplicate email with the same properties as the original,
// with the exception of a unique ID.
func (r *EmailService) Clone(ctx context.Context, body EmailCloneParams, opts ...option.RequestOption) (res *PublicEmail, err error) {
	opts = slices.Concat(r.options, opts)
	path := "marketing/emails/2026-03/clone"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Create a variation of a marketing email for an A/B test. The new variation will
// be created as a draft. If an active variation already exists, a new one won't be
// created.
func (r *EmailService) NewAbTestVariation(ctx context.Context, body EmailNewAbTestVariationParams, opts ...option.RequestOption) (res *PublicEmail, err error) {
	opts = slices.Concat(r.options, opts)
	path := "marketing/emails/2026-03/ab-test/create-variation"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Use this endpoint to get aggregated statistics of emails sent in a specified
// time span. It also returns the list of emails that were sent during the time
// span.
func (r *EmailService) Get(ctx context.Context, query EmailGetParams, opts ...option.RequestOption) (res *AggregateEmailStatistics, err error) {
	opts = slices.Concat(r.options, opts)
	path := "marketing/emails/2026-03/statistics/list"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// This endpoint lets you obtain the variation of an A/B marketing email. If the
// email is variation A (master) it will return variation B (variant) and vice
// versa.
func (r *EmailService) GetAbTestVariation(ctx context.Context, emailID string, query EmailGetAbTestVariationParams, opts ...option.RequestOption) (res *PublicEmail, err error) {
	opts = slices.Concat(r.options, opts)
	if emailID == "" {
		err = errors.New("missing required emailId parameter")
		return nil, err
	}
	path := fmt.Sprintf("marketing/emails/2026-03/%s/ab-test/get-variation", url.PathEscape(emailID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Get the draft version of an email (if it exists). If no draft version exists,
// the published email is returned.
func (r *EmailService) GetDraft(ctx context.Context, emailID string, opts ...option.RequestOption) (res *PublicEmail, err error) {
	opts = slices.Concat(r.options, opts)
	if emailID == "" {
		err = errors.New("missing required emailId parameter")
		return nil, err
	}
	path := fmt.Sprintf("marketing/emails/2026-03/%s/draft", url.PathEscape(emailID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Get aggregated statistics in intervals for a specified time span. Each interval
// contains aggregated statistics of the emails that were sent in that time.
func (r *EmailService) GetHistogram(ctx context.Context, query EmailGetHistogramParams, opts ...option.RequestOption) (res *CollectionResponseWithTotalEmailStatisticInterval, err error) {
	opts = slices.Concat(r.options, opts)
	path := "marketing/emails/2026-03/statistics/histogram"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Get a specific revision of a marketing email.
func (r *EmailService) GetRevision(ctx context.Context, revisionID string, query EmailGetRevisionParams, opts ...option.RequestOption) (res *PublicEmailVersion, err error) {
	opts = slices.Concat(r.options, opts)
	if query.EmailID == "" {
		err = errors.New("missing required emailId parameter")
		return nil, err
	}
	if revisionID == "" {
		err = errors.New("missing required revisionId parameter")
		return nil, err
	}
	path := fmt.Sprintf("marketing/emails/2026-03/%s/revisions/%s", url.PathEscape(query.EmailID), url.PathEscape(revisionID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Get a list of all versions of a marketing email, with each entry including the
// full state of that particular version. To view the most recent version, sort by
// the updatedAt parameter.
func (r *EmailService) ListRevisions(ctx context.Context, emailID string, query EmailListRevisionsParams, opts ...option.RequestOption) (res *pagination.Page[VersionPublicEmail], err error) {
	var raw *http.Response
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if emailID == "" {
		err = errors.New("missing required emailId parameter")
		return nil, err
	}
	path := fmt.Sprintf("marketing/emails/2026-03/%s/revisions", url.PathEscape(emailID))
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

// Get a list of all versions of a marketing email, with each entry including the
// full state of that particular version. To view the most recent version, sort by
// the updatedAt parameter.
func (r *EmailService) ListRevisionsAutoPaging(ctx context.Context, emailID string, query EmailListRevisionsParams, opts ...option.RequestOption) *pagination.PageAutoPager[VersionPublicEmail] {
	return pagination.NewPageAutoPager(r.ListRevisions(ctx, emailID, query, opts...))
}

// If you have a Marketing Hub Enterprise account or the transactional email
// add-on, you can use this endpoint to publish an automated email or send/schedule
// a regular email.
func (r *EmailService) Publish(ctx context.Context, emailID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if emailID == "" {
		err = errors.New("missing required emailId parameter")
		return err
	}
	path := fmt.Sprintf("marketing/emails/2026-03/%s/publish", url.PathEscape(emailID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, nil, opts...)
	return err
}

// Resets the draft back to a copy of the live object.
func (r *EmailService) ResetDraft(ctx context.Context, emailID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if emailID == "" {
		err = errors.New("missing required emailId parameter")
		return err
	}
	path := fmt.Sprintf("marketing/emails/2026-03/%s/draft/reset", url.PathEscape(emailID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, nil, opts...)
	return err
}

// Restores a previous revision of a marketing email. The current revision becomes
// old, and the restored revision is given a new version number.
func (r *EmailService) RestoreRevision(ctx context.Context, revisionID string, body EmailRestoreRevisionParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if body.EmailID == "" {
		err = errors.New("missing required emailId parameter")
		return err
	}
	if revisionID == "" {
		err = errors.New("missing required revisionId parameter")
		return err
	}
	path := fmt.Sprintf("marketing/emails/2026-03/%s/revisions/%s/restore", url.PathEscape(body.EmailID), url.PathEscape(revisionID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, nil, opts...)
	return err
}

// Restores a previous revision of a marketing email to DRAFT state. If there is
// currently something in the draft for that object, it is overwritten.
func (r *EmailService) RestoreRevisionToDraft(ctx context.Context, revisionID int64, body EmailRestoreRevisionToDraftParams, opts ...option.RequestOption) (res *PublicEmail, err error) {
	opts = slices.Concat(r.options, opts)
	if body.EmailID == "" {
		err = errors.New("missing required emailId parameter")
		return nil, err
	}
	path := fmt.Sprintf("marketing/emails/2026-03/%s/revisions/%v/restore-to-draft", url.PathEscape(body.EmailID), revisionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

// If you have a Marketing Hub Enterprise account or the transactional email
// add-on, you can use this endpoint to unpublish an automated email or cancel a
// regular email. If the email is already in the process of being sent, canceling
// might not be possible.
func (r *EmailService) Unpublish(ctx context.Context, emailID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if emailID == "" {
		err = errors.New("missing required emailId parameter")
		return err
	}
	path := fmt.Sprintf("marketing/emails/2026-03/%s/unpublish", url.PathEscape(emailID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, nil, opts...)
	return err
}

// Create or update the draft version of a marketing email. If no draft exists, the
// system creates a draft from the current “live” email then applies the request
// body to that draft. The draft version only lives on the buffer—the email is not
// cloned.
func (r *EmailService) UpdateDraft(ctx context.Context, emailID string, body EmailUpdateDraftParams, opts ...option.RequestOption) (res *PublicEmail, err error) {
	opts = slices.Concat(r.options, opts)
	if emailID == "" {
		err = errors.New("missing required emailId parameter")
		return nil, err
	}
	path := fmt.Sprintf("marketing/emails/2026-03/%s/draft", url.PathEscape(emailID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

type AggregateEmailStatistics struct {
	Aggregate EmailStatisticsData `json:"aggregate" api:"required"`
	// The aggregated statistics per campaign.
	CampaignAggregations map[string]EmailStatisticsData `json:"campaignAggregations" api:"required"`
	// List of email IDs that were sent during the time span.
	Emails []int64 `json:"emails" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Aggregate            respjson.Field
		CampaignAggregations respjson.Field
		Emails               respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AggregateEmailStatistics) RawJSON() string { return r.JSON.raw }
func (r *AggregateEmailStatistics) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CollectionResponseWithTotalEmailStatisticInterval struct {
	Results []EmailStatisticInterval `json:"results" api:"required"`
	Total   int64                    `json:"total" api:"required"`
	Paging  shared.Paging            `json:"paging"`
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
func (r CollectionResponseWithTotalEmailStatisticInterval) RawJSON() string { return r.JSON.raw }
func (r *CollectionResponseWithTotalEmailStatisticInterval) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CollectionResponseWithTotalPublicEmail struct {
	Results []PublicEmail `json:"results" api:"required"`
	Total   int64         `json:"total" api:"required"`
	Paging  shared.Paging `json:"paging"`
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
func (r CollectionResponseWithTotalPublicEmail) RawJSON() string { return r.JSON.raw }
func (r *CollectionResponseWithTotalPublicEmail) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CollectionResponseWithTotalPublicEmailVersion struct {
	Results []VersionPublicEmail `json:"results" api:"required"`
	Total   int64                `json:"total" api:"required"`
	Paging  shared.Paging        `json:"paging"`
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
func (r CollectionResponseWithTotalPublicEmailVersion) RawJSON() string { return r.JSON.raw }
func (r *CollectionResponseWithTotalPublicEmailVersion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property ID is required.
type EmailCloneRequestVNextParam struct {
	// The email ID.
	ID string `json:"id" api:"required"`
	// The name to assign to the cloned email.
	CloneName param.Opt[string] `json:"cloneName,omitzero"`
	// The language code for the cloned email, such as 'en' for English.
	Language param.Opt[string] `json:"language,omitzero"`
	paramObj
}

func (r EmailCloneRequestVNextParam) MarshalJSON() (data []byte, err error) {
	type shadow EmailCloneRequestVNextParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EmailCloneRequestVNextParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmailCreateRequestParam struct {
	// The active domain of the email.
	ActiveDomain param.Opt[string] `json:"activeDomain,omitzero"`
	// Determines if the email is archived or not.
	Archived param.Opt[bool] `json:"archived,omitzero"`
	// The ID of the business unit associated with the email.
	BusinessUnitID param.Opt[int64] `json:"businessUnitId,omitzero"`
	// The ID of the campaign this email is associated to.
	Campaign param.Opt[string] `json:"campaign,omitzero"`
	// The ID of the feedback survey linked to the email.
	FeedbackSurveyID param.Opt[string] `json:"feedbackSurveyId,omitzero"`
	// The ID of the folder where the email will be stored.
	FolderIDV2 param.Opt[int64] `json:"folderIdV2,omitzero"`
	// Determines whether the email send time should be randomized to avoid sending all
	// emails at the exact same time.
	JitterSendTime param.Opt[bool] `json:"jitterSendTime,omitzero"`
	// The name of the email, as displayed on the email dashboard.
	Name param.Opt[string] `json:"name,omitzero"`
	// The date and time the email is scheduled for, in ISO8601 representation. This is
	// only used in local time or scheduled emails.
	PublishDate param.Opt[time.Time] `json:"publishDate,omitzero" format:"date-time"`
	// Determines whether the email will be sent immediately on publish.
	SendOnPublish param.Opt[bool] `json:"sendOnPublish,omitzero"`
	// The subject of the email.
	Subject param.Opt[string]           `json:"subject,omitzero"`
	Content PublicEmailContentParam     `json:"content,omitzero"`
	From    PublicEmailFromDetailsParam `json:"from,omitzero"`
	// The language code for the email, such as 'en' for English.
	//
	// Any of "aa", "ab", "ae", "af", "af-na", "af-za", "agq", "agq-cm", "ak", "ak-gh",
	// "am", "am-et", "an", "ann", "ann-ng", "ar", "ar-001", "ar-ae", "ar-bh", "ar-dj",
	// "ar-dz", "ar-eg", "ar-eh", "ar-er", "ar-il", "ar-iq", "ar-jo", "ar-km", "ar-kw",
	// "ar-lb", "ar-ly", "ar-ma", "ar-mr", "ar-om", "ar-ps", "ar-qa", "ar-sa", "ar-sd",
	// "ar-so", "ar-ss", "ar-sy", "ar-td", "ar-tn", "ar-ye", "as", "as-in", "asa",
	// "asa-tz", "ast", "ast-es", "av", "ay", "az", "az-az", "ba", "bal", "bal-pk",
	// "bas", "bas-cm", "be", "be-by", "bem", "bem-zm", "bez", "bez-tz", "bg", "bg-bg",
	// "bgc", "bgc-in", "bho", "bho-in", "bi", "blo", "blo-bj", "bm", "bm-ml", "bn",
	// "bn-bd", "bn-in", "bo", "bo-cn", "bo-in", "br", "br-fr", "brx", "brx-in", "bs",
	// "bs-ba", "ca", "ca-ad", "ca-es", "ca-fr", "ca-it", "ccp", "ccp-bd", "ccp-in",
	// "ce", "ce-ru", "ceb", "ceb-ph", "cgg", "cgg-ug", "ch", "chr", "chr-us", "ckb",
	// "ckb-iq", "ckb-ir", "co", "cr", "cs", "cs-cz", "csw", "csw-ca", "cu", "cu-ru",
	// "cv", "cv-ru", "cy", "cy-gb", "da", "da-dk", "da-gl", "dav", "dav-ke", "de",
	// "de-at", "de-be", "de-ch", "de-de", "de-gr", "de-it", "de-li", "de-lu", "dje",
	// "dje-ne", "doi", "doi-in", "dsb", "dsb-de", "dua", "dua-cm", "dv", "dyo",
	// "dyo-sn", "dz", "dz-bt", "ebu", "ebu-ke", "ee", "ee-gh", "ee-tg", "el", "el-cy",
	// "el-gr", "en", "en-001", "en-150", "en-ae", "en-ag", "en-ai", "en-as", "en-at",
	// "en-au", "en-bb", "en-be", "en-bi", "en-bm", "en-bs", "en-bw", "en-bz", "en-ca",
	// "en-cc", "en-ch", "en-ck", "en-cm", "en-cn", "en-cx", "en-cy", "en-cz", "en-de",
	// "en-dg", "en-dk", "en-dm", "en-ee", "en-eg", "en-er", "en-es", "en-fi", "en-fj",
	// "en-fk", "en-fm", "en-fr", "en-gb", "en-gd", "en-gg", "en-gh", "en-gi", "en-gm",
	// "en-gs", "en-gu", "en-gy", "en-hk", "en-hu", "en-id", "en-ie", "en-il", "en-im",
	// "en-in", "en-io", "en-it", "en-je", "en-jm", "en-ke", "en-ki", "en-kn", "en-ky",
	// "en-lc", "en-lr", "en-ls", "en-lu", "en-mg", "en-mh", "en-mo", "en-mp", "en-ms",
	// "en-mt", "en-mu", "en-mv", "en-mw", "en-mx", "en-my", "en-na", "en-nf", "en-ng",
	// "en-nl", "en-no", "en-nr", "en-nu", "en-nz", "en-pg", "en-ph", "en-pk", "en-pl",
	// "en-pn", "en-pr", "en-pt", "en-pw", "en-ro", "en-rw", "en-sb", "en-sc", "en-sd",
	// "en-se", "en-sg", "en-sh", "en-si", "en-sk", "en-sl", "en-ss", "en-sx", "en-sz",
	// "en-tc", "en-th", "en-tk", "en-tn", "en-to", "en-tt", "en-tv", "en-tz", "en-ug",
	// "en-um", "en-us", "en-vc", "en-vg", "en-vi", "en-vn", "en-vu", "en-ws", "en-za",
	// "en-zm", "en-zw", "eo", "eo-001", "es", "es-419", "es-ar", "es-bo", "es-br",
	// "es-bz", "es-cl", "es-co", "es-cr", "es-cu", "es-do", "es-ea", "es-ec", "es-es",
	// "es-gq", "es-gt", "es-hn", "es-ic", "es-mx", "es-ni", "es-pa", "es-pe", "es-ph",
	// "es-pr", "es-py", "es-sv", "es-us", "es-uy", "es-ve", "et", "et-ee", "eu",
	// "eu-es", "ewo", "ewo-cm", "fa", "fa-af", "fa-ir", "ff", "ff-bf", "ff-cm",
	// "ff-gh", "ff-gm", "ff-gn", "ff-gw", "ff-lr", "ff-mr", "ff-ne", "ff-ng", "ff-sl",
	// "ff-sn", "fi", "fi-fi", "fil", "fil-ph", "fj", "fo", "fo-dk", "fo-fo", "fr",
	// "fr-be", "fr-bf", "fr-bi", "fr-bj", "fr-bl", "fr-ca", "fr-cd", "fr-cf", "fr-cg",
	// "fr-ch", "fr-ci", "fr-cm", "fr-dj", "fr-dz", "fr-fr", "fr-ga", "fr-gf", "fr-gn",
	// "fr-gp", "fr-gq", "fr-ht", "fr-km", "fr-lu", "fr-ma", "fr-mc", "fr-mf", "fr-mg",
	// "fr-ml", "fr-mq", "fr-mr", "fr-mu", "fr-nc", "fr-ne", "fr-pf", "fr-pm", "fr-re",
	// "fr-rw", "fr-sc", "fr-sn", "fr-sy", "fr-td", "fr-tg", "fr-tn", "fr-vu", "fr-wf",
	// "fr-yt", "frr", "frr-de", "fur", "fur-it", "fy", "fy-nl", "ga", "ga-gb",
	// "ga-ie", "gaa", "gaa-gh", "gd", "gd-gb", "gl", "gl-es", "gn", "gsw", "gsw-ch",
	// "gsw-fr", "gsw-li", "gu", "gu-in", "guz", "guz-ke", "gv", "gv-im", "ha",
	// "ha-gh", "ha-ne", "ha-ng", "haw", "haw-us", "he", "he-il", "hi", "hi-in", "hmn",
	// "ho", "hr", "hr-ba", "hr-hr", "hsb", "hsb-de", "ht", "ht-ht", "hu", "hu-hu",
	// "hy", "hy-am", "hz", "ia", "ia-001", "id", "id-id", "ie", "ie-ee", "ig",
	// "ig-ng", "ii", "ii-cn", "ik", "io", "is", "is-is", "it", "it-ch", "it-it",
	// "it-sm", "it-va", "iu", "ja", "ja-jp", "jgo", "jgo-cm", "jmc", "jmc-tz", "jv",
	// "jv-id", "ka", "ka-ge", "kab", "kab-dz", "kam", "kam-ke", "kar", "kde",
	// "kde-tz", "kea", "kea-cv", "kg", "kgp", "kgp-br", "kh", "khq", "khq-ml", "ki",
	// "ki-ke", "kj", "kk", "kk-kz", "kkj", "kkj-cm", "kl", "kl-gl", "kln", "kln-ke",
	// "km", "km-kh", "kn", "kn-in", "ko", "ko-cn", "ko-kp", "ko-kr", "kok", "kok-in",
	// "kr", "ks", "ks-in", "ksb", "ksb-tz", "ksf", "ksf-cm", "ksh", "ksh-de", "ku",
	// "ku-tr", "kv", "kw", "kw-gb", "kxv", "kxv-in", "ky", "ky-kg", "la", "lag",
	// "lag-tz", "lb", "lb-lu", "lg", "lg-ug", "li", "lij", "lij-it", "lkt", "lkt-us",
	// "lmo", "lmo-it", "ln", "ln-ao", "ln-cd", "ln-cf", "ln-cg", "lo", "lo-la", "lrc",
	// "lrc-iq", "lrc-ir", "lt", "lt-lt", "lu", "lu-cd", "luo", "luo-ke", "luy",
	// "luy-ke", "lv", "lv-lv", "mai", "mai-in", "mas", "mas-ke", "mas-tz", "mdf",
	// "mdf-ru", "mer", "mer-ke", "mfe", "mfe-mu", "mg", "mg-mg", "mgh", "mgh-mz",
	// "mgo", "mgo-cm", "mh", "mi", "mi-nz", "mk", "mk-mk", "ml", "ml-in", "mn",
	// "mn-mn", "mni", "mni-in", "mr", "mr-in", "ms", "ms-bn", "ms-id", "ms-my",
	// "ms-sg", "mt", "mt-mt", "mua", "mua-cm", "my", "my-mm", "mzn", "mzn-ir", "na",
	// "naq", "naq-na", "nb", "nb-no", "nb-sj", "nd", "nd-zw", "nds", "nds-de",
	// "nds-nl", "ne", "ne-in", "ne-np", "ng", "nl", "nl-aw", "nl-be", "nl-bq",
	// "nl-ch", "nl-cw", "nl-lu", "nl-nl", "nl-sr", "nl-sx", "nmg", "nmg-cm", "nn",
	// "nn-no", "nnh", "nnh-cm", "no", "no-no", "nqo", "nqo-gn", "nr", "nso", "nso-za",
	// "nus", "nus-ss", "nv", "ny", "nyn", "nyn-ug", "oc", "oc-es", "oc-fr", "oj",
	// "om", "om-et", "om-ke", "or", "or-in", "os", "os-ge", "os-ru", "pa", "pa-in",
	// "pa-pk", "pcm", "pcm-ng", "pi", "pis", "pis-sb", "pl", "pl-pl", "prg",
	// "prg-001", "ps", "ps-af", "ps-pk", "pt", "pt-ao", "pt-br", "pt-ch", "pt-cv",
	// "pt-gq", "pt-gw", "pt-lu", "pt-mo", "pt-mz", "pt-pt", "pt-st", "pt-tl", "qu",
	// "qu-bo", "qu-ec", "qu-pe", "raj", "raj-in", "rm", "rm-ch", "rn", "rn-bi", "ro",
	// "ro-md", "ro-ro", "rof", "rof-tz", "ru", "ru-by", "ru-kg", "ru-kz", "ru-md",
	// "ru-ru", "ru-ua", "rw", "rw-rw", "rwk", "rwk-tz", "sa", "sa-in", "sah",
	// "sah-ru", "saq", "saq-ke", "sat", "sat-in", "sbp", "sbp-tz", "sc", "sc-it",
	// "sd", "sd-in", "sd-pk", "se", "se-fi", "se-no", "se-se", "seh", "seh-mz", "ses",
	// "ses-ml", "sg", "sg-cf", "shi", "shi-ma", "si", "si-lk", "sk", "sk-sk", "sl",
	// "sl-si", "sm", "smn", "smn-fi", "sms", "sms-fi", "sn", "sn-zw", "so", "so-dj",
	// "so-et", "so-ke", "so-so", "sq", "sq-al", "sq-mk", "sq-xk", "sr", "sr-ba",
	// "sr-cs", "sr-me", "sr-rs", "sr-xk", "ss", "st", "st-ls", "st-za", "su", "su-id",
	// "sv", "sv-ax", "sv-fi", "sv-se", "sw", "sw-cd", "sw-ke", "sw-tz", "sw-ug", "sy",
	// "syr", "syr-iq", "syr-sy", "szl", "szl-pl", "ta", "ta-in", "ta-lk", "ta-my",
	// "ta-sg", "te", "te-in", "teo", "teo-ke", "teo-ug", "tg", "tg-tj", "th", "th-th",
	// "ti", "ti-er", "ti-et", "tk", "tk-tm", "tl", "tn", "tn-bw", "tn-za", "to",
	// "to-to", "tok", "tok-001", "tr", "tr-cy", "tr-tr", "ts", "tt", "tt-ru", "tw",
	// "twq", "twq-ne", "ty", "tzm", "tzm-ma", "ug", "ug-cn", "uk", "uk-ua", "ur",
	// "ur-in", "ur-pk", "uz", "uz-af", "uz-uz", "vai", "vai-lr", "ve", "vec",
	// "vec-it", "vi", "vi-vn", "vmw", "vmw-mz", "vo", "vo-001", "vun", "vun-tz", "wa",
	// "wae", "wae-ch", "wo", "wo-sn", "xh", "xh-za", "xnr", "xnr-in", "xog", "xog-ug",
	// "yav", "yav-cm", "yi", "yi-001", "yi-ua", "yo", "yo-bj", "yo-ng", "yrl",
	// "yrl-br", "yrl-co", "yrl-ve", "yue", "yue-cn", "yue-hk", "yue-mo", "za",
	// "za-cn", "zgh", "zgh-ma", "zh", "zh-cn", "zh-hans", "zh-hant", "zh-hk", "zh-mo",
	// "zh-my", "zh-sg", "zh-tw", "zu", "zu-za".
	Language EmailCreateRequestLanguage `json:"language,omitzero"`
	RssData  PublicRssEmailDetailsParam `json:"rssData,omitzero"`
	// The email state.
	//
	// Any of "AGENT_GENERATED", "AUTOMATED", "AUTOMATED_AB", "AUTOMATED_AB_VARIANT",
	// "AUTOMATED_DRAFT", "AUTOMATED_DRAFT_AB", "AUTOMATED_DRAFT_ABVARIANT",
	// "AUTOMATED_FOR_FORM", "AUTOMATED_FOR_FORM_BUFFER", "AUTOMATED_FOR_FORM_DRAFT",
	// "AUTOMATED_FOR_FORM_LEGACY", "AUTOMATED_LOSER_ABVARIANT", "AUTOMATED_SENDING",
	// "BLOG_EMAIL_DRAFT", "BLOG_EMAIL_PUBLISHED", "DRAFT", "DRAFT_AB",
	// "DRAFT_AB_VARIANT", "ERROR", "LOSER_AB_VARIANT", "PAGE_STUB", "PRE_PROCESSING",
	// "PROCESSING", "PUBLISHED", "PUBLISHED_AB", "PUBLISHED_AB_VARIANT",
	// "PUBLISHED_OR_SCHEDULED", "RSS_TO_EMAIL_DRAFT", "RSS_TO_EMAIL_PUBLISHED",
	// "SCHEDULED", "SCHEDULED_AB", "SCHEDULED_OR_PUBLISHED".
	State EmailCreateRequestState `json:"state,omitzero"`
	// The email subcategory.
	//
	// Any of "ab_loser_variant", "ab_loser_variant_site_page", "ab_master",
	// "ab_master_site_page", "ab_variant", "ab_variant_site_page", "automated",
	// "automated_ab_master", "automated_ab_variant", "automated_for_crm",
	// "automated_for_custom_survey", "automated_for_deal",
	// "automated_for_feedback_ces", "automated_for_feedback_custom",
	// "automated_for_feedback_nps", "automated_for_form", "automated_for_form_buffer",
	// "automated_for_form_draft", "automated_for_form_legacy",
	// "automated_for_leadflow", "automated_for_ticket", "batch",
	// "blog_article_instance_layout", "blog_article_listing", "blog_author_detail",
	// "blog_email", "blog_email_child", "case_study", "case_study_instance_layout",
	// "case_study_listing", "discardable_stub", "imported_blog_post", "kb_404_page",
	// "kb_article_instance_layout", "kb_listing", "kb_search_results",
	// "kb_support_form", "landing_page", "legacy_blog_post", "legacy_page",
	// "localtime", "manage_preferences_email", "marketing_single_send_api",
	// "membership_email_verification", "membership_follow_up", "membership_otp_login",
	// "membership_password_reset", "membership_password_saved",
	// "membership_passwordless_auth", "membership_registration",
	// "membership_registration_follow_up", "membership_verification",
	// "normal_blog_post", "optin_email", "optin_followup_email",
	// "page_instance_layout", "page_stub", "performable_landing_page",
	// "performable_landing_page_cutover", "podcast_instance_layout",
	// "podcast_listing", "portal_content", "resubscribe_confirmation_email",
	// "resubscribe_email", "rss_to_email", "rss_to_email_child",
	// "scp_instance_layout_page", "scp_static_page", "single_send_api", "site_page",
	// "smtp_token", "staged_page", "ticket_closed_kickback_email",
	// "ticket_opened_kickback_email", "ticket_pipeline_automated", "UNKNOWN",
	// "unsubscribe_confirmation_email", "web_interactive".
	Subcategory         EmailCreateRequestSubcategory       `json:"subcategory,omitzero"`
	SubscriptionDetails PublicEmailSubscriptionDetailsParam `json:"subscriptionDetails,omitzero"`
	Testing             PublicEmailTestingDetailsParam      `json:"testing,omitzero"`
	To                  PublicEmailToDetailsParam           `json:"to,omitzero"`
	Webversion          PublicWebversionDetailsParam        `json:"webversion,omitzero"`
	paramObj
}

func (r EmailCreateRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow EmailCreateRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EmailCreateRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The language code for the email, such as 'en' for English.
type EmailCreateRequestLanguage string

const (
	EmailCreateRequestLanguageAa     EmailCreateRequestLanguage = "aa"
	EmailCreateRequestLanguageAb     EmailCreateRequestLanguage = "ab"
	EmailCreateRequestLanguageAe     EmailCreateRequestLanguage = "ae"
	EmailCreateRequestLanguageAf     EmailCreateRequestLanguage = "af"
	EmailCreateRequestLanguageAfNa   EmailCreateRequestLanguage = "af-na"
	EmailCreateRequestLanguageAfZa   EmailCreateRequestLanguage = "af-za"
	EmailCreateRequestLanguageAgq    EmailCreateRequestLanguage = "agq"
	EmailCreateRequestLanguageAgqCm  EmailCreateRequestLanguage = "agq-cm"
	EmailCreateRequestLanguageAk     EmailCreateRequestLanguage = "ak"
	EmailCreateRequestLanguageAkGh   EmailCreateRequestLanguage = "ak-gh"
	EmailCreateRequestLanguageAm     EmailCreateRequestLanguage = "am"
	EmailCreateRequestLanguageAmEt   EmailCreateRequestLanguage = "am-et"
	EmailCreateRequestLanguageAn     EmailCreateRequestLanguage = "an"
	EmailCreateRequestLanguageAnn    EmailCreateRequestLanguage = "ann"
	EmailCreateRequestLanguageAnnNg  EmailCreateRequestLanguage = "ann-ng"
	EmailCreateRequestLanguageAr     EmailCreateRequestLanguage = "ar"
	EmailCreateRequestLanguageAr001  EmailCreateRequestLanguage = "ar-001"
	EmailCreateRequestLanguageArAe   EmailCreateRequestLanguage = "ar-ae"
	EmailCreateRequestLanguageArBh   EmailCreateRequestLanguage = "ar-bh"
	EmailCreateRequestLanguageArDj   EmailCreateRequestLanguage = "ar-dj"
	EmailCreateRequestLanguageArDz   EmailCreateRequestLanguage = "ar-dz"
	EmailCreateRequestLanguageArEg   EmailCreateRequestLanguage = "ar-eg"
	EmailCreateRequestLanguageArEh   EmailCreateRequestLanguage = "ar-eh"
	EmailCreateRequestLanguageArEr   EmailCreateRequestLanguage = "ar-er"
	EmailCreateRequestLanguageArIl   EmailCreateRequestLanguage = "ar-il"
	EmailCreateRequestLanguageArIq   EmailCreateRequestLanguage = "ar-iq"
	EmailCreateRequestLanguageArJo   EmailCreateRequestLanguage = "ar-jo"
	EmailCreateRequestLanguageArKm   EmailCreateRequestLanguage = "ar-km"
	EmailCreateRequestLanguageArKw   EmailCreateRequestLanguage = "ar-kw"
	EmailCreateRequestLanguageArLb   EmailCreateRequestLanguage = "ar-lb"
	EmailCreateRequestLanguageArLy   EmailCreateRequestLanguage = "ar-ly"
	EmailCreateRequestLanguageArMa   EmailCreateRequestLanguage = "ar-ma"
	EmailCreateRequestLanguageArMr   EmailCreateRequestLanguage = "ar-mr"
	EmailCreateRequestLanguageArOm   EmailCreateRequestLanguage = "ar-om"
	EmailCreateRequestLanguageArPs   EmailCreateRequestLanguage = "ar-ps"
	EmailCreateRequestLanguageArQa   EmailCreateRequestLanguage = "ar-qa"
	EmailCreateRequestLanguageArSa   EmailCreateRequestLanguage = "ar-sa"
	EmailCreateRequestLanguageArSd   EmailCreateRequestLanguage = "ar-sd"
	EmailCreateRequestLanguageArSo   EmailCreateRequestLanguage = "ar-so"
	EmailCreateRequestLanguageArSS   EmailCreateRequestLanguage = "ar-ss"
	EmailCreateRequestLanguageArSy   EmailCreateRequestLanguage = "ar-sy"
	EmailCreateRequestLanguageArTd   EmailCreateRequestLanguage = "ar-td"
	EmailCreateRequestLanguageArTn   EmailCreateRequestLanguage = "ar-tn"
	EmailCreateRequestLanguageArYe   EmailCreateRequestLanguage = "ar-ye"
	EmailCreateRequestLanguageAs     EmailCreateRequestLanguage = "as"
	EmailCreateRequestLanguageAsIn   EmailCreateRequestLanguage = "as-in"
	EmailCreateRequestLanguageAsa    EmailCreateRequestLanguage = "asa"
	EmailCreateRequestLanguageAsaTz  EmailCreateRequestLanguage = "asa-tz"
	EmailCreateRequestLanguageAst    EmailCreateRequestLanguage = "ast"
	EmailCreateRequestLanguageAstEs  EmailCreateRequestLanguage = "ast-es"
	EmailCreateRequestLanguageAv     EmailCreateRequestLanguage = "av"
	EmailCreateRequestLanguageAy     EmailCreateRequestLanguage = "ay"
	EmailCreateRequestLanguageAz     EmailCreateRequestLanguage = "az"
	EmailCreateRequestLanguageAzAz   EmailCreateRequestLanguage = "az-az"
	EmailCreateRequestLanguageBa     EmailCreateRequestLanguage = "ba"
	EmailCreateRequestLanguageBal    EmailCreateRequestLanguage = "bal"
	EmailCreateRequestLanguageBalPk  EmailCreateRequestLanguage = "bal-pk"
	EmailCreateRequestLanguageBas    EmailCreateRequestLanguage = "bas"
	EmailCreateRequestLanguageBasCm  EmailCreateRequestLanguage = "bas-cm"
	EmailCreateRequestLanguageBe     EmailCreateRequestLanguage = "be"
	EmailCreateRequestLanguageBeBy   EmailCreateRequestLanguage = "be-by"
	EmailCreateRequestLanguageBem    EmailCreateRequestLanguage = "bem"
	EmailCreateRequestLanguageBemZm  EmailCreateRequestLanguage = "bem-zm"
	EmailCreateRequestLanguageBez    EmailCreateRequestLanguage = "bez"
	EmailCreateRequestLanguageBezTz  EmailCreateRequestLanguage = "bez-tz"
	EmailCreateRequestLanguageBg     EmailCreateRequestLanguage = "bg"
	EmailCreateRequestLanguageBgBg   EmailCreateRequestLanguage = "bg-bg"
	EmailCreateRequestLanguageBgc    EmailCreateRequestLanguage = "bgc"
	EmailCreateRequestLanguageBgcIn  EmailCreateRequestLanguage = "bgc-in"
	EmailCreateRequestLanguageBho    EmailCreateRequestLanguage = "bho"
	EmailCreateRequestLanguageBhoIn  EmailCreateRequestLanguage = "bho-in"
	EmailCreateRequestLanguageBi     EmailCreateRequestLanguage = "bi"
	EmailCreateRequestLanguageBlo    EmailCreateRequestLanguage = "blo"
	EmailCreateRequestLanguageBloBj  EmailCreateRequestLanguage = "blo-bj"
	EmailCreateRequestLanguageBm     EmailCreateRequestLanguage = "bm"
	EmailCreateRequestLanguageBmMl   EmailCreateRequestLanguage = "bm-ml"
	EmailCreateRequestLanguageBn     EmailCreateRequestLanguage = "bn"
	EmailCreateRequestLanguageBnBd   EmailCreateRequestLanguage = "bn-bd"
	EmailCreateRequestLanguageBnIn   EmailCreateRequestLanguage = "bn-in"
	EmailCreateRequestLanguageBo     EmailCreateRequestLanguage = "bo"
	EmailCreateRequestLanguageBoCn   EmailCreateRequestLanguage = "bo-cn"
	EmailCreateRequestLanguageBoIn   EmailCreateRequestLanguage = "bo-in"
	EmailCreateRequestLanguageBr     EmailCreateRequestLanguage = "br"
	EmailCreateRequestLanguageBrFr   EmailCreateRequestLanguage = "br-fr"
	EmailCreateRequestLanguageBrx    EmailCreateRequestLanguage = "brx"
	EmailCreateRequestLanguageBrxIn  EmailCreateRequestLanguage = "brx-in"
	EmailCreateRequestLanguageBs     EmailCreateRequestLanguage = "bs"
	EmailCreateRequestLanguageBsBa   EmailCreateRequestLanguage = "bs-ba"
	EmailCreateRequestLanguageCa     EmailCreateRequestLanguage = "ca"
	EmailCreateRequestLanguageCaAd   EmailCreateRequestLanguage = "ca-ad"
	EmailCreateRequestLanguageCaEs   EmailCreateRequestLanguage = "ca-es"
	EmailCreateRequestLanguageCaFr   EmailCreateRequestLanguage = "ca-fr"
	EmailCreateRequestLanguageCaIt   EmailCreateRequestLanguage = "ca-it"
	EmailCreateRequestLanguageCcp    EmailCreateRequestLanguage = "ccp"
	EmailCreateRequestLanguageCcpBd  EmailCreateRequestLanguage = "ccp-bd"
	EmailCreateRequestLanguageCcpIn  EmailCreateRequestLanguage = "ccp-in"
	EmailCreateRequestLanguageCe     EmailCreateRequestLanguage = "ce"
	EmailCreateRequestLanguageCeRu   EmailCreateRequestLanguage = "ce-ru"
	EmailCreateRequestLanguageCeb    EmailCreateRequestLanguage = "ceb"
	EmailCreateRequestLanguageCebPh  EmailCreateRequestLanguage = "ceb-ph"
	EmailCreateRequestLanguageCgg    EmailCreateRequestLanguage = "cgg"
	EmailCreateRequestLanguageCggUg  EmailCreateRequestLanguage = "cgg-ug"
	EmailCreateRequestLanguageCh     EmailCreateRequestLanguage = "ch"
	EmailCreateRequestLanguageChr    EmailCreateRequestLanguage = "chr"
	EmailCreateRequestLanguageChrUs  EmailCreateRequestLanguage = "chr-us"
	EmailCreateRequestLanguageCkb    EmailCreateRequestLanguage = "ckb"
	EmailCreateRequestLanguageCkbIq  EmailCreateRequestLanguage = "ckb-iq"
	EmailCreateRequestLanguageCkbIr  EmailCreateRequestLanguage = "ckb-ir"
	EmailCreateRequestLanguageCo     EmailCreateRequestLanguage = "co"
	EmailCreateRequestLanguageCr     EmailCreateRequestLanguage = "cr"
	EmailCreateRequestLanguageCs     EmailCreateRequestLanguage = "cs"
	EmailCreateRequestLanguageCsCz   EmailCreateRequestLanguage = "cs-cz"
	EmailCreateRequestLanguageCsw    EmailCreateRequestLanguage = "csw"
	EmailCreateRequestLanguageCswCa  EmailCreateRequestLanguage = "csw-ca"
	EmailCreateRequestLanguageCu     EmailCreateRequestLanguage = "cu"
	EmailCreateRequestLanguageCuRu   EmailCreateRequestLanguage = "cu-ru"
	EmailCreateRequestLanguageCv     EmailCreateRequestLanguage = "cv"
	EmailCreateRequestLanguageCvRu   EmailCreateRequestLanguage = "cv-ru"
	EmailCreateRequestLanguageCy     EmailCreateRequestLanguage = "cy"
	EmailCreateRequestLanguageCyGB   EmailCreateRequestLanguage = "cy-gb"
	EmailCreateRequestLanguageDa     EmailCreateRequestLanguage = "da"
	EmailCreateRequestLanguageDaDk   EmailCreateRequestLanguage = "da-dk"
	EmailCreateRequestLanguageDaGl   EmailCreateRequestLanguage = "da-gl"
	EmailCreateRequestLanguageDav    EmailCreateRequestLanguage = "dav"
	EmailCreateRequestLanguageDavKe  EmailCreateRequestLanguage = "dav-ke"
	EmailCreateRequestLanguageDe     EmailCreateRequestLanguage = "de"
	EmailCreateRequestLanguageDeAt   EmailCreateRequestLanguage = "de-at"
	EmailCreateRequestLanguageDeBe   EmailCreateRequestLanguage = "de-be"
	EmailCreateRequestLanguageDeCh   EmailCreateRequestLanguage = "de-ch"
	EmailCreateRequestLanguageDeDe   EmailCreateRequestLanguage = "de-de"
	EmailCreateRequestLanguageDeGr   EmailCreateRequestLanguage = "de-gr"
	EmailCreateRequestLanguageDeIt   EmailCreateRequestLanguage = "de-it"
	EmailCreateRequestLanguageDeLi   EmailCreateRequestLanguage = "de-li"
	EmailCreateRequestLanguageDeLu   EmailCreateRequestLanguage = "de-lu"
	EmailCreateRequestLanguageDje    EmailCreateRequestLanguage = "dje"
	EmailCreateRequestLanguageDjeNe  EmailCreateRequestLanguage = "dje-ne"
	EmailCreateRequestLanguageDoi    EmailCreateRequestLanguage = "doi"
	EmailCreateRequestLanguageDoiIn  EmailCreateRequestLanguage = "doi-in"
	EmailCreateRequestLanguageDsb    EmailCreateRequestLanguage = "dsb"
	EmailCreateRequestLanguageDsbDe  EmailCreateRequestLanguage = "dsb-de"
	EmailCreateRequestLanguageDua    EmailCreateRequestLanguage = "dua"
	EmailCreateRequestLanguageDuaCm  EmailCreateRequestLanguage = "dua-cm"
	EmailCreateRequestLanguageDv     EmailCreateRequestLanguage = "dv"
	EmailCreateRequestLanguageDyo    EmailCreateRequestLanguage = "dyo"
	EmailCreateRequestLanguageDyoSn  EmailCreateRequestLanguage = "dyo-sn"
	EmailCreateRequestLanguageDz     EmailCreateRequestLanguage = "dz"
	EmailCreateRequestLanguageDzBt   EmailCreateRequestLanguage = "dz-bt"
	EmailCreateRequestLanguageEbu    EmailCreateRequestLanguage = "ebu"
	EmailCreateRequestLanguageEbuKe  EmailCreateRequestLanguage = "ebu-ke"
	EmailCreateRequestLanguageEe     EmailCreateRequestLanguage = "ee"
	EmailCreateRequestLanguageEeGh   EmailCreateRequestLanguage = "ee-gh"
	EmailCreateRequestLanguageEeTg   EmailCreateRequestLanguage = "ee-tg"
	EmailCreateRequestLanguageEl     EmailCreateRequestLanguage = "el"
	EmailCreateRequestLanguageElCy   EmailCreateRequestLanguage = "el-cy"
	EmailCreateRequestLanguageElGr   EmailCreateRequestLanguage = "el-gr"
	EmailCreateRequestLanguageEn     EmailCreateRequestLanguage = "en"
	EmailCreateRequestLanguageEn001  EmailCreateRequestLanguage = "en-001"
	EmailCreateRequestLanguageEn150  EmailCreateRequestLanguage = "en-150"
	EmailCreateRequestLanguageEnAe   EmailCreateRequestLanguage = "en-ae"
	EmailCreateRequestLanguageEnAg   EmailCreateRequestLanguage = "en-ag"
	EmailCreateRequestLanguageEnAI   EmailCreateRequestLanguage = "en-ai"
	EmailCreateRequestLanguageEnAs   EmailCreateRequestLanguage = "en-as"
	EmailCreateRequestLanguageEnAt   EmailCreateRequestLanguage = "en-at"
	EmailCreateRequestLanguageEnAu   EmailCreateRequestLanguage = "en-au"
	EmailCreateRequestLanguageEnBb   EmailCreateRequestLanguage = "en-bb"
	EmailCreateRequestLanguageEnBe   EmailCreateRequestLanguage = "en-be"
	EmailCreateRequestLanguageEnBi   EmailCreateRequestLanguage = "en-bi"
	EmailCreateRequestLanguageEnBm   EmailCreateRequestLanguage = "en-bm"
	EmailCreateRequestLanguageEnBs   EmailCreateRequestLanguage = "en-bs"
	EmailCreateRequestLanguageEnBw   EmailCreateRequestLanguage = "en-bw"
	EmailCreateRequestLanguageEnBz   EmailCreateRequestLanguage = "en-bz"
	EmailCreateRequestLanguageEnCa   EmailCreateRequestLanguage = "en-ca"
	EmailCreateRequestLanguageEnCc   EmailCreateRequestLanguage = "en-cc"
	EmailCreateRequestLanguageEnCh   EmailCreateRequestLanguage = "en-ch"
	EmailCreateRequestLanguageEnCk   EmailCreateRequestLanguage = "en-ck"
	EmailCreateRequestLanguageEnCm   EmailCreateRequestLanguage = "en-cm"
	EmailCreateRequestLanguageEnCn   EmailCreateRequestLanguage = "en-cn"
	EmailCreateRequestLanguageEnCx   EmailCreateRequestLanguage = "en-cx"
	EmailCreateRequestLanguageEnCy   EmailCreateRequestLanguage = "en-cy"
	EmailCreateRequestLanguageEnCz   EmailCreateRequestLanguage = "en-cz"
	EmailCreateRequestLanguageEnDe   EmailCreateRequestLanguage = "en-de"
	EmailCreateRequestLanguageEnDg   EmailCreateRequestLanguage = "en-dg"
	EmailCreateRequestLanguageEnDk   EmailCreateRequestLanguage = "en-dk"
	EmailCreateRequestLanguageEnDm   EmailCreateRequestLanguage = "en-dm"
	EmailCreateRequestLanguageEnEe   EmailCreateRequestLanguage = "en-ee"
	EmailCreateRequestLanguageEnEg   EmailCreateRequestLanguage = "en-eg"
	EmailCreateRequestLanguageEnEr   EmailCreateRequestLanguage = "en-er"
	EmailCreateRequestLanguageEnEs   EmailCreateRequestLanguage = "en-es"
	EmailCreateRequestLanguageEnFi   EmailCreateRequestLanguage = "en-fi"
	EmailCreateRequestLanguageEnFj   EmailCreateRequestLanguage = "en-fj"
	EmailCreateRequestLanguageEnFk   EmailCreateRequestLanguage = "en-fk"
	EmailCreateRequestLanguageEnFm   EmailCreateRequestLanguage = "en-fm"
	EmailCreateRequestLanguageEnFr   EmailCreateRequestLanguage = "en-fr"
	EmailCreateRequestLanguageEnGB   EmailCreateRequestLanguage = "en-gb"
	EmailCreateRequestLanguageEnGd   EmailCreateRequestLanguage = "en-gd"
	EmailCreateRequestLanguageEnGg   EmailCreateRequestLanguage = "en-gg"
	EmailCreateRequestLanguageEnGh   EmailCreateRequestLanguage = "en-gh"
	EmailCreateRequestLanguageEnGi   EmailCreateRequestLanguage = "en-gi"
	EmailCreateRequestLanguageEnGm   EmailCreateRequestLanguage = "en-gm"
	EmailCreateRequestLanguageEnGs   EmailCreateRequestLanguage = "en-gs"
	EmailCreateRequestLanguageEnGu   EmailCreateRequestLanguage = "en-gu"
	EmailCreateRequestLanguageEnGy   EmailCreateRequestLanguage = "en-gy"
	EmailCreateRequestLanguageEnHk   EmailCreateRequestLanguage = "en-hk"
	EmailCreateRequestLanguageEnHu   EmailCreateRequestLanguage = "en-hu"
	EmailCreateRequestLanguageEnID   EmailCreateRequestLanguage = "en-id"
	EmailCreateRequestLanguageEnIe   EmailCreateRequestLanguage = "en-ie"
	EmailCreateRequestLanguageEnIl   EmailCreateRequestLanguage = "en-il"
	EmailCreateRequestLanguageEnIm   EmailCreateRequestLanguage = "en-im"
	EmailCreateRequestLanguageEnIn   EmailCreateRequestLanguage = "en-in"
	EmailCreateRequestLanguageEnIo   EmailCreateRequestLanguage = "en-io"
	EmailCreateRequestLanguageEnIt   EmailCreateRequestLanguage = "en-it"
	EmailCreateRequestLanguageEnJe   EmailCreateRequestLanguage = "en-je"
	EmailCreateRequestLanguageEnJm   EmailCreateRequestLanguage = "en-jm"
	EmailCreateRequestLanguageEnKe   EmailCreateRequestLanguage = "en-ke"
	EmailCreateRequestLanguageEnKi   EmailCreateRequestLanguage = "en-ki"
	EmailCreateRequestLanguageEnKn   EmailCreateRequestLanguage = "en-kn"
	EmailCreateRequestLanguageEnKy   EmailCreateRequestLanguage = "en-ky"
	EmailCreateRequestLanguageEnLc   EmailCreateRequestLanguage = "en-lc"
	EmailCreateRequestLanguageEnLr   EmailCreateRequestLanguage = "en-lr"
	EmailCreateRequestLanguageEnLs   EmailCreateRequestLanguage = "en-ls"
	EmailCreateRequestLanguageEnLu   EmailCreateRequestLanguage = "en-lu"
	EmailCreateRequestLanguageEnMg   EmailCreateRequestLanguage = "en-mg"
	EmailCreateRequestLanguageEnMh   EmailCreateRequestLanguage = "en-mh"
	EmailCreateRequestLanguageEnMo   EmailCreateRequestLanguage = "en-mo"
	EmailCreateRequestLanguageEnMp   EmailCreateRequestLanguage = "en-mp"
	EmailCreateRequestLanguageEnMs   EmailCreateRequestLanguage = "en-ms"
	EmailCreateRequestLanguageEnMt   EmailCreateRequestLanguage = "en-mt"
	EmailCreateRequestLanguageEnMu   EmailCreateRequestLanguage = "en-mu"
	EmailCreateRequestLanguageEnMv   EmailCreateRequestLanguage = "en-mv"
	EmailCreateRequestLanguageEnMw   EmailCreateRequestLanguage = "en-mw"
	EmailCreateRequestLanguageEnMx   EmailCreateRequestLanguage = "en-mx"
	EmailCreateRequestLanguageEnMy   EmailCreateRequestLanguage = "en-my"
	EmailCreateRequestLanguageEnNa   EmailCreateRequestLanguage = "en-na"
	EmailCreateRequestLanguageEnNf   EmailCreateRequestLanguage = "en-nf"
	EmailCreateRequestLanguageEnNg   EmailCreateRequestLanguage = "en-ng"
	EmailCreateRequestLanguageEnNl   EmailCreateRequestLanguage = "en-nl"
	EmailCreateRequestLanguageEnNo   EmailCreateRequestLanguage = "en-no"
	EmailCreateRequestLanguageEnNr   EmailCreateRequestLanguage = "en-nr"
	EmailCreateRequestLanguageEnNu   EmailCreateRequestLanguage = "en-nu"
	EmailCreateRequestLanguageEnNz   EmailCreateRequestLanguage = "en-nz"
	EmailCreateRequestLanguageEnPg   EmailCreateRequestLanguage = "en-pg"
	EmailCreateRequestLanguageEnPh   EmailCreateRequestLanguage = "en-ph"
	EmailCreateRequestLanguageEnPk   EmailCreateRequestLanguage = "en-pk"
	EmailCreateRequestLanguageEnPl   EmailCreateRequestLanguage = "en-pl"
	EmailCreateRequestLanguageEnPn   EmailCreateRequestLanguage = "en-pn"
	EmailCreateRequestLanguageEnPr   EmailCreateRequestLanguage = "en-pr"
	EmailCreateRequestLanguageEnPt   EmailCreateRequestLanguage = "en-pt"
	EmailCreateRequestLanguageEnPw   EmailCreateRequestLanguage = "en-pw"
	EmailCreateRequestLanguageEnRo   EmailCreateRequestLanguage = "en-ro"
	EmailCreateRequestLanguageEnRw   EmailCreateRequestLanguage = "en-rw"
	EmailCreateRequestLanguageEnSb   EmailCreateRequestLanguage = "en-sb"
	EmailCreateRequestLanguageEnSc   EmailCreateRequestLanguage = "en-sc"
	EmailCreateRequestLanguageEnSd   EmailCreateRequestLanguage = "en-sd"
	EmailCreateRequestLanguageEnSe   EmailCreateRequestLanguage = "en-se"
	EmailCreateRequestLanguageEnSg   EmailCreateRequestLanguage = "en-sg"
	EmailCreateRequestLanguageEnSh   EmailCreateRequestLanguage = "en-sh"
	EmailCreateRequestLanguageEnSi   EmailCreateRequestLanguage = "en-si"
	EmailCreateRequestLanguageEnSk   EmailCreateRequestLanguage = "en-sk"
	EmailCreateRequestLanguageEnSl   EmailCreateRequestLanguage = "en-sl"
	EmailCreateRequestLanguageEnSS   EmailCreateRequestLanguage = "en-ss"
	EmailCreateRequestLanguageEnSx   EmailCreateRequestLanguage = "en-sx"
	EmailCreateRequestLanguageEnSz   EmailCreateRequestLanguage = "en-sz"
	EmailCreateRequestLanguageEnTc   EmailCreateRequestLanguage = "en-tc"
	EmailCreateRequestLanguageEnTh   EmailCreateRequestLanguage = "en-th"
	EmailCreateRequestLanguageEnTk   EmailCreateRequestLanguage = "en-tk"
	EmailCreateRequestLanguageEnTn   EmailCreateRequestLanguage = "en-tn"
	EmailCreateRequestLanguageEnTo   EmailCreateRequestLanguage = "en-to"
	EmailCreateRequestLanguageEnTt   EmailCreateRequestLanguage = "en-tt"
	EmailCreateRequestLanguageEnTv   EmailCreateRequestLanguage = "en-tv"
	EmailCreateRequestLanguageEnTz   EmailCreateRequestLanguage = "en-tz"
	EmailCreateRequestLanguageEnUg   EmailCreateRequestLanguage = "en-ug"
	EmailCreateRequestLanguageEnUm   EmailCreateRequestLanguage = "en-um"
	EmailCreateRequestLanguageEnUs   EmailCreateRequestLanguage = "en-us"
	EmailCreateRequestLanguageEnVc   EmailCreateRequestLanguage = "en-vc"
	EmailCreateRequestLanguageEnVg   EmailCreateRequestLanguage = "en-vg"
	EmailCreateRequestLanguageEnVi   EmailCreateRequestLanguage = "en-vi"
	EmailCreateRequestLanguageEnVn   EmailCreateRequestLanguage = "en-vn"
	EmailCreateRequestLanguageEnVu   EmailCreateRequestLanguage = "en-vu"
	EmailCreateRequestLanguageEnWs   EmailCreateRequestLanguage = "en-ws"
	EmailCreateRequestLanguageEnZa   EmailCreateRequestLanguage = "en-za"
	EmailCreateRequestLanguageEnZm   EmailCreateRequestLanguage = "en-zm"
	EmailCreateRequestLanguageEnZw   EmailCreateRequestLanguage = "en-zw"
	EmailCreateRequestLanguageEo     EmailCreateRequestLanguage = "eo"
	EmailCreateRequestLanguageEo001  EmailCreateRequestLanguage = "eo-001"
	EmailCreateRequestLanguageEs     EmailCreateRequestLanguage = "es"
	EmailCreateRequestLanguageEs419  EmailCreateRequestLanguage = "es-419"
	EmailCreateRequestLanguageEsAr   EmailCreateRequestLanguage = "es-ar"
	EmailCreateRequestLanguageEsBo   EmailCreateRequestLanguage = "es-bo"
	EmailCreateRequestLanguageEsBr   EmailCreateRequestLanguage = "es-br"
	EmailCreateRequestLanguageEsBz   EmailCreateRequestLanguage = "es-bz"
	EmailCreateRequestLanguageEsCl   EmailCreateRequestLanguage = "es-cl"
	EmailCreateRequestLanguageEsCo   EmailCreateRequestLanguage = "es-co"
	EmailCreateRequestLanguageEsCr   EmailCreateRequestLanguage = "es-cr"
	EmailCreateRequestLanguageEsCu   EmailCreateRequestLanguage = "es-cu"
	EmailCreateRequestLanguageEsDo   EmailCreateRequestLanguage = "es-do"
	EmailCreateRequestLanguageEsEa   EmailCreateRequestLanguage = "es-ea"
	EmailCreateRequestLanguageEsEc   EmailCreateRequestLanguage = "es-ec"
	EmailCreateRequestLanguageEsEs   EmailCreateRequestLanguage = "es-es"
	EmailCreateRequestLanguageEsGq   EmailCreateRequestLanguage = "es-gq"
	EmailCreateRequestLanguageEsGt   EmailCreateRequestLanguage = "es-gt"
	EmailCreateRequestLanguageEsHn   EmailCreateRequestLanguage = "es-hn"
	EmailCreateRequestLanguageEsIc   EmailCreateRequestLanguage = "es-ic"
	EmailCreateRequestLanguageEsMx   EmailCreateRequestLanguage = "es-mx"
	EmailCreateRequestLanguageEsNi   EmailCreateRequestLanguage = "es-ni"
	EmailCreateRequestLanguageEsPa   EmailCreateRequestLanguage = "es-pa"
	EmailCreateRequestLanguageEsPe   EmailCreateRequestLanguage = "es-pe"
	EmailCreateRequestLanguageEsPh   EmailCreateRequestLanguage = "es-ph"
	EmailCreateRequestLanguageEsPr   EmailCreateRequestLanguage = "es-pr"
	EmailCreateRequestLanguageEsPy   EmailCreateRequestLanguage = "es-py"
	EmailCreateRequestLanguageEsSv   EmailCreateRequestLanguage = "es-sv"
	EmailCreateRequestLanguageEsUs   EmailCreateRequestLanguage = "es-us"
	EmailCreateRequestLanguageEsUy   EmailCreateRequestLanguage = "es-uy"
	EmailCreateRequestLanguageEsVe   EmailCreateRequestLanguage = "es-ve"
	EmailCreateRequestLanguageEt     EmailCreateRequestLanguage = "et"
	EmailCreateRequestLanguageEtEe   EmailCreateRequestLanguage = "et-ee"
	EmailCreateRequestLanguageEu     EmailCreateRequestLanguage = "eu"
	EmailCreateRequestLanguageEuEs   EmailCreateRequestLanguage = "eu-es"
	EmailCreateRequestLanguageEwo    EmailCreateRequestLanguage = "ewo"
	EmailCreateRequestLanguageEwoCm  EmailCreateRequestLanguage = "ewo-cm"
	EmailCreateRequestLanguageFa     EmailCreateRequestLanguage = "fa"
	EmailCreateRequestLanguageFaAf   EmailCreateRequestLanguage = "fa-af"
	EmailCreateRequestLanguageFaIr   EmailCreateRequestLanguage = "fa-ir"
	EmailCreateRequestLanguageFf     EmailCreateRequestLanguage = "ff"
	EmailCreateRequestLanguageFfBf   EmailCreateRequestLanguage = "ff-bf"
	EmailCreateRequestLanguageFfCm   EmailCreateRequestLanguage = "ff-cm"
	EmailCreateRequestLanguageFfGh   EmailCreateRequestLanguage = "ff-gh"
	EmailCreateRequestLanguageFfGm   EmailCreateRequestLanguage = "ff-gm"
	EmailCreateRequestLanguageFfGn   EmailCreateRequestLanguage = "ff-gn"
	EmailCreateRequestLanguageFfGw   EmailCreateRequestLanguage = "ff-gw"
	EmailCreateRequestLanguageFfLr   EmailCreateRequestLanguage = "ff-lr"
	EmailCreateRequestLanguageFfMr   EmailCreateRequestLanguage = "ff-mr"
	EmailCreateRequestLanguageFfNe   EmailCreateRequestLanguage = "ff-ne"
	EmailCreateRequestLanguageFfNg   EmailCreateRequestLanguage = "ff-ng"
	EmailCreateRequestLanguageFfSl   EmailCreateRequestLanguage = "ff-sl"
	EmailCreateRequestLanguageFfSn   EmailCreateRequestLanguage = "ff-sn"
	EmailCreateRequestLanguageFi     EmailCreateRequestLanguage = "fi"
	EmailCreateRequestLanguageFiFi   EmailCreateRequestLanguage = "fi-fi"
	EmailCreateRequestLanguageFil    EmailCreateRequestLanguage = "fil"
	EmailCreateRequestLanguageFilPh  EmailCreateRequestLanguage = "fil-ph"
	EmailCreateRequestLanguageFj     EmailCreateRequestLanguage = "fj"
	EmailCreateRequestLanguageFo     EmailCreateRequestLanguage = "fo"
	EmailCreateRequestLanguageFoDk   EmailCreateRequestLanguage = "fo-dk"
	EmailCreateRequestLanguageFoFo   EmailCreateRequestLanguage = "fo-fo"
	EmailCreateRequestLanguageFr     EmailCreateRequestLanguage = "fr"
	EmailCreateRequestLanguageFrBe   EmailCreateRequestLanguage = "fr-be"
	EmailCreateRequestLanguageFrBf   EmailCreateRequestLanguage = "fr-bf"
	EmailCreateRequestLanguageFrBi   EmailCreateRequestLanguage = "fr-bi"
	EmailCreateRequestLanguageFrBj   EmailCreateRequestLanguage = "fr-bj"
	EmailCreateRequestLanguageFrBl   EmailCreateRequestLanguage = "fr-bl"
	EmailCreateRequestLanguageFrCa   EmailCreateRequestLanguage = "fr-ca"
	EmailCreateRequestLanguageFrCd   EmailCreateRequestLanguage = "fr-cd"
	EmailCreateRequestLanguageFrCf   EmailCreateRequestLanguage = "fr-cf"
	EmailCreateRequestLanguageFrCg   EmailCreateRequestLanguage = "fr-cg"
	EmailCreateRequestLanguageFrCh   EmailCreateRequestLanguage = "fr-ch"
	EmailCreateRequestLanguageFrCi   EmailCreateRequestLanguage = "fr-ci"
	EmailCreateRequestLanguageFrCm   EmailCreateRequestLanguage = "fr-cm"
	EmailCreateRequestLanguageFrDj   EmailCreateRequestLanguage = "fr-dj"
	EmailCreateRequestLanguageFrDz   EmailCreateRequestLanguage = "fr-dz"
	EmailCreateRequestLanguageFrFr   EmailCreateRequestLanguage = "fr-fr"
	EmailCreateRequestLanguageFrGa   EmailCreateRequestLanguage = "fr-ga"
	EmailCreateRequestLanguageFrGf   EmailCreateRequestLanguage = "fr-gf"
	EmailCreateRequestLanguageFrGn   EmailCreateRequestLanguage = "fr-gn"
	EmailCreateRequestLanguageFrGp   EmailCreateRequestLanguage = "fr-gp"
	EmailCreateRequestLanguageFrGq   EmailCreateRequestLanguage = "fr-gq"
	EmailCreateRequestLanguageFrHt   EmailCreateRequestLanguage = "fr-ht"
	EmailCreateRequestLanguageFrKm   EmailCreateRequestLanguage = "fr-km"
	EmailCreateRequestLanguageFrLu   EmailCreateRequestLanguage = "fr-lu"
	EmailCreateRequestLanguageFrMa   EmailCreateRequestLanguage = "fr-ma"
	EmailCreateRequestLanguageFrMc   EmailCreateRequestLanguage = "fr-mc"
	EmailCreateRequestLanguageFrMf   EmailCreateRequestLanguage = "fr-mf"
	EmailCreateRequestLanguageFrMg   EmailCreateRequestLanguage = "fr-mg"
	EmailCreateRequestLanguageFrMl   EmailCreateRequestLanguage = "fr-ml"
	EmailCreateRequestLanguageFrMq   EmailCreateRequestLanguage = "fr-mq"
	EmailCreateRequestLanguageFrMr   EmailCreateRequestLanguage = "fr-mr"
	EmailCreateRequestLanguageFrMu   EmailCreateRequestLanguage = "fr-mu"
	EmailCreateRequestLanguageFrNc   EmailCreateRequestLanguage = "fr-nc"
	EmailCreateRequestLanguageFrNe   EmailCreateRequestLanguage = "fr-ne"
	EmailCreateRequestLanguageFrPf   EmailCreateRequestLanguage = "fr-pf"
	EmailCreateRequestLanguageFrPm   EmailCreateRequestLanguage = "fr-pm"
	EmailCreateRequestLanguageFrRe   EmailCreateRequestLanguage = "fr-re"
	EmailCreateRequestLanguageFrRw   EmailCreateRequestLanguage = "fr-rw"
	EmailCreateRequestLanguageFrSc   EmailCreateRequestLanguage = "fr-sc"
	EmailCreateRequestLanguageFrSn   EmailCreateRequestLanguage = "fr-sn"
	EmailCreateRequestLanguageFrSy   EmailCreateRequestLanguage = "fr-sy"
	EmailCreateRequestLanguageFrTd   EmailCreateRequestLanguage = "fr-td"
	EmailCreateRequestLanguageFrTg   EmailCreateRequestLanguage = "fr-tg"
	EmailCreateRequestLanguageFrTn   EmailCreateRequestLanguage = "fr-tn"
	EmailCreateRequestLanguageFrVu   EmailCreateRequestLanguage = "fr-vu"
	EmailCreateRequestLanguageFrWf   EmailCreateRequestLanguage = "fr-wf"
	EmailCreateRequestLanguageFrYt   EmailCreateRequestLanguage = "fr-yt"
	EmailCreateRequestLanguageFrr    EmailCreateRequestLanguage = "frr"
	EmailCreateRequestLanguageFrrDe  EmailCreateRequestLanguage = "frr-de"
	EmailCreateRequestLanguageFur    EmailCreateRequestLanguage = "fur"
	EmailCreateRequestLanguageFurIt  EmailCreateRequestLanguage = "fur-it"
	EmailCreateRequestLanguageFy     EmailCreateRequestLanguage = "fy"
	EmailCreateRequestLanguageFyNl   EmailCreateRequestLanguage = "fy-nl"
	EmailCreateRequestLanguageGa     EmailCreateRequestLanguage = "ga"
	EmailCreateRequestLanguageGaGB   EmailCreateRequestLanguage = "ga-gb"
	EmailCreateRequestLanguageGaIe   EmailCreateRequestLanguage = "ga-ie"
	EmailCreateRequestLanguageGaa    EmailCreateRequestLanguage = "gaa"
	EmailCreateRequestLanguageGaaGh  EmailCreateRequestLanguage = "gaa-gh"
	EmailCreateRequestLanguageGd     EmailCreateRequestLanguage = "gd"
	EmailCreateRequestLanguageGdGB   EmailCreateRequestLanguage = "gd-gb"
	EmailCreateRequestLanguageGl     EmailCreateRequestLanguage = "gl"
	EmailCreateRequestLanguageGlEs   EmailCreateRequestLanguage = "gl-es"
	EmailCreateRequestLanguageGn     EmailCreateRequestLanguage = "gn"
	EmailCreateRequestLanguageGsw    EmailCreateRequestLanguage = "gsw"
	EmailCreateRequestLanguageGswCh  EmailCreateRequestLanguage = "gsw-ch"
	EmailCreateRequestLanguageGswFr  EmailCreateRequestLanguage = "gsw-fr"
	EmailCreateRequestLanguageGswLi  EmailCreateRequestLanguage = "gsw-li"
	EmailCreateRequestLanguageGu     EmailCreateRequestLanguage = "gu"
	EmailCreateRequestLanguageGuIn   EmailCreateRequestLanguage = "gu-in"
	EmailCreateRequestLanguageGuz    EmailCreateRequestLanguage = "guz"
	EmailCreateRequestLanguageGuzKe  EmailCreateRequestLanguage = "guz-ke"
	EmailCreateRequestLanguageGv     EmailCreateRequestLanguage = "gv"
	EmailCreateRequestLanguageGvIm   EmailCreateRequestLanguage = "gv-im"
	EmailCreateRequestLanguageHa     EmailCreateRequestLanguage = "ha"
	EmailCreateRequestLanguageHaGh   EmailCreateRequestLanguage = "ha-gh"
	EmailCreateRequestLanguageHaNe   EmailCreateRequestLanguage = "ha-ne"
	EmailCreateRequestLanguageHaNg   EmailCreateRequestLanguage = "ha-ng"
	EmailCreateRequestLanguageHaw    EmailCreateRequestLanguage = "haw"
	EmailCreateRequestLanguageHawUs  EmailCreateRequestLanguage = "haw-us"
	EmailCreateRequestLanguageHe     EmailCreateRequestLanguage = "he"
	EmailCreateRequestLanguageHeIl   EmailCreateRequestLanguage = "he-il"
	EmailCreateRequestLanguageHi     EmailCreateRequestLanguage = "hi"
	EmailCreateRequestLanguageHiIn   EmailCreateRequestLanguage = "hi-in"
	EmailCreateRequestLanguageHmn    EmailCreateRequestLanguage = "hmn"
	EmailCreateRequestLanguageHo     EmailCreateRequestLanguage = "ho"
	EmailCreateRequestLanguageHr     EmailCreateRequestLanguage = "hr"
	EmailCreateRequestLanguageHrBa   EmailCreateRequestLanguage = "hr-ba"
	EmailCreateRequestLanguageHrHr   EmailCreateRequestLanguage = "hr-hr"
	EmailCreateRequestLanguageHsb    EmailCreateRequestLanguage = "hsb"
	EmailCreateRequestLanguageHsbDe  EmailCreateRequestLanguage = "hsb-de"
	EmailCreateRequestLanguageHt     EmailCreateRequestLanguage = "ht"
	EmailCreateRequestLanguageHtHt   EmailCreateRequestLanguage = "ht-ht"
	EmailCreateRequestLanguageHu     EmailCreateRequestLanguage = "hu"
	EmailCreateRequestLanguageHuHu   EmailCreateRequestLanguage = "hu-hu"
	EmailCreateRequestLanguageHy     EmailCreateRequestLanguage = "hy"
	EmailCreateRequestLanguageHyAm   EmailCreateRequestLanguage = "hy-am"
	EmailCreateRequestLanguageHz     EmailCreateRequestLanguage = "hz"
	EmailCreateRequestLanguageIa     EmailCreateRequestLanguage = "ia"
	EmailCreateRequestLanguageIa001  EmailCreateRequestLanguage = "ia-001"
	EmailCreateRequestLanguageID     EmailCreateRequestLanguage = "id"
	EmailCreateRequestLanguageIDID   EmailCreateRequestLanguage = "id-id"
	EmailCreateRequestLanguageIe     EmailCreateRequestLanguage = "ie"
	EmailCreateRequestLanguageIeEe   EmailCreateRequestLanguage = "ie-ee"
	EmailCreateRequestLanguageIg     EmailCreateRequestLanguage = "ig"
	EmailCreateRequestLanguageIgNg   EmailCreateRequestLanguage = "ig-ng"
	EmailCreateRequestLanguageIi     EmailCreateRequestLanguage = "ii"
	EmailCreateRequestLanguageIiCn   EmailCreateRequestLanguage = "ii-cn"
	EmailCreateRequestLanguageIk     EmailCreateRequestLanguage = "ik"
	EmailCreateRequestLanguageIo     EmailCreateRequestLanguage = "io"
	EmailCreateRequestLanguageIs     EmailCreateRequestLanguage = "is"
	EmailCreateRequestLanguageIsIs   EmailCreateRequestLanguage = "is-is"
	EmailCreateRequestLanguageIt     EmailCreateRequestLanguage = "it"
	EmailCreateRequestLanguageItCh   EmailCreateRequestLanguage = "it-ch"
	EmailCreateRequestLanguageItIt   EmailCreateRequestLanguage = "it-it"
	EmailCreateRequestLanguageItSm   EmailCreateRequestLanguage = "it-sm"
	EmailCreateRequestLanguageItVa   EmailCreateRequestLanguage = "it-va"
	EmailCreateRequestLanguageIu     EmailCreateRequestLanguage = "iu"
	EmailCreateRequestLanguageJa     EmailCreateRequestLanguage = "ja"
	EmailCreateRequestLanguageJaJp   EmailCreateRequestLanguage = "ja-jp"
	EmailCreateRequestLanguageJgo    EmailCreateRequestLanguage = "jgo"
	EmailCreateRequestLanguageJgoCm  EmailCreateRequestLanguage = "jgo-cm"
	EmailCreateRequestLanguageJmc    EmailCreateRequestLanguage = "jmc"
	EmailCreateRequestLanguageJmcTz  EmailCreateRequestLanguage = "jmc-tz"
	EmailCreateRequestLanguageJv     EmailCreateRequestLanguage = "jv"
	EmailCreateRequestLanguageJvID   EmailCreateRequestLanguage = "jv-id"
	EmailCreateRequestLanguageKa     EmailCreateRequestLanguage = "ka"
	EmailCreateRequestLanguageKaGe   EmailCreateRequestLanguage = "ka-ge"
	EmailCreateRequestLanguageKab    EmailCreateRequestLanguage = "kab"
	EmailCreateRequestLanguageKabDz  EmailCreateRequestLanguage = "kab-dz"
	EmailCreateRequestLanguageKam    EmailCreateRequestLanguage = "kam"
	EmailCreateRequestLanguageKamKe  EmailCreateRequestLanguage = "kam-ke"
	EmailCreateRequestLanguageKar    EmailCreateRequestLanguage = "kar"
	EmailCreateRequestLanguageKde    EmailCreateRequestLanguage = "kde"
	EmailCreateRequestLanguageKdeTz  EmailCreateRequestLanguage = "kde-tz"
	EmailCreateRequestLanguageKea    EmailCreateRequestLanguage = "kea"
	EmailCreateRequestLanguageKeaCv  EmailCreateRequestLanguage = "kea-cv"
	EmailCreateRequestLanguageKg     EmailCreateRequestLanguage = "kg"
	EmailCreateRequestLanguageKgp    EmailCreateRequestLanguage = "kgp"
	EmailCreateRequestLanguageKgpBr  EmailCreateRequestLanguage = "kgp-br"
	EmailCreateRequestLanguageKh     EmailCreateRequestLanguage = "kh"
	EmailCreateRequestLanguageKhq    EmailCreateRequestLanguage = "khq"
	EmailCreateRequestLanguageKhqMl  EmailCreateRequestLanguage = "khq-ml"
	EmailCreateRequestLanguageKi     EmailCreateRequestLanguage = "ki"
	EmailCreateRequestLanguageKiKe   EmailCreateRequestLanguage = "ki-ke"
	EmailCreateRequestLanguageKj     EmailCreateRequestLanguage = "kj"
	EmailCreateRequestLanguageKk     EmailCreateRequestLanguage = "kk"
	EmailCreateRequestLanguageKkKz   EmailCreateRequestLanguage = "kk-kz"
	EmailCreateRequestLanguageKkj    EmailCreateRequestLanguage = "kkj"
	EmailCreateRequestLanguageKkjCm  EmailCreateRequestLanguage = "kkj-cm"
	EmailCreateRequestLanguageKl     EmailCreateRequestLanguage = "kl"
	EmailCreateRequestLanguageKlGl   EmailCreateRequestLanguage = "kl-gl"
	EmailCreateRequestLanguageKln    EmailCreateRequestLanguage = "kln"
	EmailCreateRequestLanguageKlnKe  EmailCreateRequestLanguage = "kln-ke"
	EmailCreateRequestLanguageKm     EmailCreateRequestLanguage = "km"
	EmailCreateRequestLanguageKmKh   EmailCreateRequestLanguage = "km-kh"
	EmailCreateRequestLanguageKn     EmailCreateRequestLanguage = "kn"
	EmailCreateRequestLanguageKnIn   EmailCreateRequestLanguage = "kn-in"
	EmailCreateRequestLanguageKo     EmailCreateRequestLanguage = "ko"
	EmailCreateRequestLanguageKoCn   EmailCreateRequestLanguage = "ko-cn"
	EmailCreateRequestLanguageKoKp   EmailCreateRequestLanguage = "ko-kp"
	EmailCreateRequestLanguageKoKr   EmailCreateRequestLanguage = "ko-kr"
	EmailCreateRequestLanguageKok    EmailCreateRequestLanguage = "kok"
	EmailCreateRequestLanguageKokIn  EmailCreateRequestLanguage = "kok-in"
	EmailCreateRequestLanguageKr     EmailCreateRequestLanguage = "kr"
	EmailCreateRequestLanguageKs     EmailCreateRequestLanguage = "ks"
	EmailCreateRequestLanguageKsIn   EmailCreateRequestLanguage = "ks-in"
	EmailCreateRequestLanguageKsb    EmailCreateRequestLanguage = "ksb"
	EmailCreateRequestLanguageKsbTz  EmailCreateRequestLanguage = "ksb-tz"
	EmailCreateRequestLanguageKsf    EmailCreateRequestLanguage = "ksf"
	EmailCreateRequestLanguageKsfCm  EmailCreateRequestLanguage = "ksf-cm"
	EmailCreateRequestLanguageKsh    EmailCreateRequestLanguage = "ksh"
	EmailCreateRequestLanguageKshDe  EmailCreateRequestLanguage = "ksh-de"
	EmailCreateRequestLanguageKu     EmailCreateRequestLanguage = "ku"
	EmailCreateRequestLanguageKuTr   EmailCreateRequestLanguage = "ku-tr"
	EmailCreateRequestLanguageKv     EmailCreateRequestLanguage = "kv"
	EmailCreateRequestLanguageKw     EmailCreateRequestLanguage = "kw"
	EmailCreateRequestLanguageKwGB   EmailCreateRequestLanguage = "kw-gb"
	EmailCreateRequestLanguageKxv    EmailCreateRequestLanguage = "kxv"
	EmailCreateRequestLanguageKxvIn  EmailCreateRequestLanguage = "kxv-in"
	EmailCreateRequestLanguageKy     EmailCreateRequestLanguage = "ky"
	EmailCreateRequestLanguageKyKg   EmailCreateRequestLanguage = "ky-kg"
	EmailCreateRequestLanguageLa     EmailCreateRequestLanguage = "la"
	EmailCreateRequestLanguageLag    EmailCreateRequestLanguage = "lag"
	EmailCreateRequestLanguageLagTz  EmailCreateRequestLanguage = "lag-tz"
	EmailCreateRequestLanguageLb     EmailCreateRequestLanguage = "lb"
	EmailCreateRequestLanguageLbLu   EmailCreateRequestLanguage = "lb-lu"
	EmailCreateRequestLanguageLg     EmailCreateRequestLanguage = "lg"
	EmailCreateRequestLanguageLgUg   EmailCreateRequestLanguage = "lg-ug"
	EmailCreateRequestLanguageLi     EmailCreateRequestLanguage = "li"
	EmailCreateRequestLanguageLij    EmailCreateRequestLanguage = "lij"
	EmailCreateRequestLanguageLijIt  EmailCreateRequestLanguage = "lij-it"
	EmailCreateRequestLanguageLkt    EmailCreateRequestLanguage = "lkt"
	EmailCreateRequestLanguageLktUs  EmailCreateRequestLanguage = "lkt-us"
	EmailCreateRequestLanguageLmo    EmailCreateRequestLanguage = "lmo"
	EmailCreateRequestLanguageLmoIt  EmailCreateRequestLanguage = "lmo-it"
	EmailCreateRequestLanguageLn     EmailCreateRequestLanguage = "ln"
	EmailCreateRequestLanguageLnAo   EmailCreateRequestLanguage = "ln-ao"
	EmailCreateRequestLanguageLnCd   EmailCreateRequestLanguage = "ln-cd"
	EmailCreateRequestLanguageLnCf   EmailCreateRequestLanguage = "ln-cf"
	EmailCreateRequestLanguageLnCg   EmailCreateRequestLanguage = "ln-cg"
	EmailCreateRequestLanguageLo     EmailCreateRequestLanguage = "lo"
	EmailCreateRequestLanguageLoLa   EmailCreateRequestLanguage = "lo-la"
	EmailCreateRequestLanguageLrc    EmailCreateRequestLanguage = "lrc"
	EmailCreateRequestLanguageLrcIq  EmailCreateRequestLanguage = "lrc-iq"
	EmailCreateRequestLanguageLrcIr  EmailCreateRequestLanguage = "lrc-ir"
	EmailCreateRequestLanguageLt     EmailCreateRequestLanguage = "lt"
	EmailCreateRequestLanguageLtLt   EmailCreateRequestLanguage = "lt-lt"
	EmailCreateRequestLanguageLu     EmailCreateRequestLanguage = "lu"
	EmailCreateRequestLanguageLuCd   EmailCreateRequestLanguage = "lu-cd"
	EmailCreateRequestLanguageLuo    EmailCreateRequestLanguage = "luo"
	EmailCreateRequestLanguageLuoKe  EmailCreateRequestLanguage = "luo-ke"
	EmailCreateRequestLanguageLuy    EmailCreateRequestLanguage = "luy"
	EmailCreateRequestLanguageLuyKe  EmailCreateRequestLanguage = "luy-ke"
	EmailCreateRequestLanguageLv     EmailCreateRequestLanguage = "lv"
	EmailCreateRequestLanguageLvLv   EmailCreateRequestLanguage = "lv-lv"
	EmailCreateRequestLanguageMai    EmailCreateRequestLanguage = "mai"
	EmailCreateRequestLanguageMaiIn  EmailCreateRequestLanguage = "mai-in"
	EmailCreateRequestLanguageMas    EmailCreateRequestLanguage = "mas"
	EmailCreateRequestLanguageMasKe  EmailCreateRequestLanguage = "mas-ke"
	EmailCreateRequestLanguageMasTz  EmailCreateRequestLanguage = "mas-tz"
	EmailCreateRequestLanguageMdf    EmailCreateRequestLanguage = "mdf"
	EmailCreateRequestLanguageMdfRu  EmailCreateRequestLanguage = "mdf-ru"
	EmailCreateRequestLanguageMer    EmailCreateRequestLanguage = "mer"
	EmailCreateRequestLanguageMerKe  EmailCreateRequestLanguage = "mer-ke"
	EmailCreateRequestLanguageMfe    EmailCreateRequestLanguage = "mfe"
	EmailCreateRequestLanguageMfeMu  EmailCreateRequestLanguage = "mfe-mu"
	EmailCreateRequestLanguageMg     EmailCreateRequestLanguage = "mg"
	EmailCreateRequestLanguageMgMg   EmailCreateRequestLanguage = "mg-mg"
	EmailCreateRequestLanguageMgh    EmailCreateRequestLanguage = "mgh"
	EmailCreateRequestLanguageMghMz  EmailCreateRequestLanguage = "mgh-mz"
	EmailCreateRequestLanguageMgo    EmailCreateRequestLanguage = "mgo"
	EmailCreateRequestLanguageMgoCm  EmailCreateRequestLanguage = "mgo-cm"
	EmailCreateRequestLanguageMh     EmailCreateRequestLanguage = "mh"
	EmailCreateRequestLanguageMi     EmailCreateRequestLanguage = "mi"
	EmailCreateRequestLanguageMiNz   EmailCreateRequestLanguage = "mi-nz"
	EmailCreateRequestLanguageMk     EmailCreateRequestLanguage = "mk"
	EmailCreateRequestLanguageMkMk   EmailCreateRequestLanguage = "mk-mk"
	EmailCreateRequestLanguageMl     EmailCreateRequestLanguage = "ml"
	EmailCreateRequestLanguageMlIn   EmailCreateRequestLanguage = "ml-in"
	EmailCreateRequestLanguageMn     EmailCreateRequestLanguage = "mn"
	EmailCreateRequestLanguageMnMn   EmailCreateRequestLanguage = "mn-mn"
	EmailCreateRequestLanguageMni    EmailCreateRequestLanguage = "mni"
	EmailCreateRequestLanguageMniIn  EmailCreateRequestLanguage = "mni-in"
	EmailCreateRequestLanguageMr     EmailCreateRequestLanguage = "mr"
	EmailCreateRequestLanguageMrIn   EmailCreateRequestLanguage = "mr-in"
	EmailCreateRequestLanguageMs     EmailCreateRequestLanguage = "ms"
	EmailCreateRequestLanguageMsBn   EmailCreateRequestLanguage = "ms-bn"
	EmailCreateRequestLanguageMsID   EmailCreateRequestLanguage = "ms-id"
	EmailCreateRequestLanguageMsMy   EmailCreateRequestLanguage = "ms-my"
	EmailCreateRequestLanguageMsSg   EmailCreateRequestLanguage = "ms-sg"
	EmailCreateRequestLanguageMt     EmailCreateRequestLanguage = "mt"
	EmailCreateRequestLanguageMtMt   EmailCreateRequestLanguage = "mt-mt"
	EmailCreateRequestLanguageMua    EmailCreateRequestLanguage = "mua"
	EmailCreateRequestLanguageMuaCm  EmailCreateRequestLanguage = "mua-cm"
	EmailCreateRequestLanguageMy     EmailCreateRequestLanguage = "my"
	EmailCreateRequestLanguageMyMm   EmailCreateRequestLanguage = "my-mm"
	EmailCreateRequestLanguageMzn    EmailCreateRequestLanguage = "mzn"
	EmailCreateRequestLanguageMznIr  EmailCreateRequestLanguage = "mzn-ir"
	EmailCreateRequestLanguageNa     EmailCreateRequestLanguage = "na"
	EmailCreateRequestLanguageNaq    EmailCreateRequestLanguage = "naq"
	EmailCreateRequestLanguageNaqNa  EmailCreateRequestLanguage = "naq-na"
	EmailCreateRequestLanguageNb     EmailCreateRequestLanguage = "nb"
	EmailCreateRequestLanguageNbNo   EmailCreateRequestLanguage = "nb-no"
	EmailCreateRequestLanguageNbSj   EmailCreateRequestLanguage = "nb-sj"
	EmailCreateRequestLanguageNd     EmailCreateRequestLanguage = "nd"
	EmailCreateRequestLanguageNdZw   EmailCreateRequestLanguage = "nd-zw"
	EmailCreateRequestLanguageNds    EmailCreateRequestLanguage = "nds"
	EmailCreateRequestLanguageNdsDe  EmailCreateRequestLanguage = "nds-de"
	EmailCreateRequestLanguageNdsNl  EmailCreateRequestLanguage = "nds-nl"
	EmailCreateRequestLanguageNe     EmailCreateRequestLanguage = "ne"
	EmailCreateRequestLanguageNeIn   EmailCreateRequestLanguage = "ne-in"
	EmailCreateRequestLanguageNeNp   EmailCreateRequestLanguage = "ne-np"
	EmailCreateRequestLanguageNg     EmailCreateRequestLanguage = "ng"
	EmailCreateRequestLanguageNl     EmailCreateRequestLanguage = "nl"
	EmailCreateRequestLanguageNlAw   EmailCreateRequestLanguage = "nl-aw"
	EmailCreateRequestLanguageNlBe   EmailCreateRequestLanguage = "nl-be"
	EmailCreateRequestLanguageNlBq   EmailCreateRequestLanguage = "nl-bq"
	EmailCreateRequestLanguageNlCh   EmailCreateRequestLanguage = "nl-ch"
	EmailCreateRequestLanguageNlCw   EmailCreateRequestLanguage = "nl-cw"
	EmailCreateRequestLanguageNlLu   EmailCreateRequestLanguage = "nl-lu"
	EmailCreateRequestLanguageNlNl   EmailCreateRequestLanguage = "nl-nl"
	EmailCreateRequestLanguageNlSr   EmailCreateRequestLanguage = "nl-sr"
	EmailCreateRequestLanguageNlSx   EmailCreateRequestLanguage = "nl-sx"
	EmailCreateRequestLanguageNmg    EmailCreateRequestLanguage = "nmg"
	EmailCreateRequestLanguageNmgCm  EmailCreateRequestLanguage = "nmg-cm"
	EmailCreateRequestLanguageNn     EmailCreateRequestLanguage = "nn"
	EmailCreateRequestLanguageNnNo   EmailCreateRequestLanguage = "nn-no"
	EmailCreateRequestLanguageNnh    EmailCreateRequestLanguage = "nnh"
	EmailCreateRequestLanguageNnhCm  EmailCreateRequestLanguage = "nnh-cm"
	EmailCreateRequestLanguageNo     EmailCreateRequestLanguage = "no"
	EmailCreateRequestLanguageNoNo   EmailCreateRequestLanguage = "no-no"
	EmailCreateRequestLanguageNqo    EmailCreateRequestLanguage = "nqo"
	EmailCreateRequestLanguageNqoGn  EmailCreateRequestLanguage = "nqo-gn"
	EmailCreateRequestLanguageNr     EmailCreateRequestLanguage = "nr"
	EmailCreateRequestLanguageNso    EmailCreateRequestLanguage = "nso"
	EmailCreateRequestLanguageNsoZa  EmailCreateRequestLanguage = "nso-za"
	EmailCreateRequestLanguageNus    EmailCreateRequestLanguage = "nus"
	EmailCreateRequestLanguageNusSS  EmailCreateRequestLanguage = "nus-ss"
	EmailCreateRequestLanguageNv     EmailCreateRequestLanguage = "nv"
	EmailCreateRequestLanguageNy     EmailCreateRequestLanguage = "ny"
	EmailCreateRequestLanguageNyn    EmailCreateRequestLanguage = "nyn"
	EmailCreateRequestLanguageNynUg  EmailCreateRequestLanguage = "nyn-ug"
	EmailCreateRequestLanguageOc     EmailCreateRequestLanguage = "oc"
	EmailCreateRequestLanguageOcEs   EmailCreateRequestLanguage = "oc-es"
	EmailCreateRequestLanguageOcFr   EmailCreateRequestLanguage = "oc-fr"
	EmailCreateRequestLanguageOj     EmailCreateRequestLanguage = "oj"
	EmailCreateRequestLanguageOm     EmailCreateRequestLanguage = "om"
	EmailCreateRequestLanguageOmEt   EmailCreateRequestLanguage = "om-et"
	EmailCreateRequestLanguageOmKe   EmailCreateRequestLanguage = "om-ke"
	EmailCreateRequestLanguageOr     EmailCreateRequestLanguage = "or"
	EmailCreateRequestLanguageOrIn   EmailCreateRequestLanguage = "or-in"
	EmailCreateRequestLanguageOs     EmailCreateRequestLanguage = "os"
	EmailCreateRequestLanguageOsGe   EmailCreateRequestLanguage = "os-ge"
	EmailCreateRequestLanguageOsRu   EmailCreateRequestLanguage = "os-ru"
	EmailCreateRequestLanguagePa     EmailCreateRequestLanguage = "pa"
	EmailCreateRequestLanguagePaIn   EmailCreateRequestLanguage = "pa-in"
	EmailCreateRequestLanguagePaPk   EmailCreateRequestLanguage = "pa-pk"
	EmailCreateRequestLanguagePcm    EmailCreateRequestLanguage = "pcm"
	EmailCreateRequestLanguagePcmNg  EmailCreateRequestLanguage = "pcm-ng"
	EmailCreateRequestLanguagePi     EmailCreateRequestLanguage = "pi"
	EmailCreateRequestLanguagePis    EmailCreateRequestLanguage = "pis"
	EmailCreateRequestLanguagePisSb  EmailCreateRequestLanguage = "pis-sb"
	EmailCreateRequestLanguagePl     EmailCreateRequestLanguage = "pl"
	EmailCreateRequestLanguagePlPl   EmailCreateRequestLanguage = "pl-pl"
	EmailCreateRequestLanguagePrg    EmailCreateRequestLanguage = "prg"
	EmailCreateRequestLanguagePrg001 EmailCreateRequestLanguage = "prg-001"
	EmailCreateRequestLanguagePs     EmailCreateRequestLanguage = "ps"
	EmailCreateRequestLanguagePsAf   EmailCreateRequestLanguage = "ps-af"
	EmailCreateRequestLanguagePsPk   EmailCreateRequestLanguage = "ps-pk"
	EmailCreateRequestLanguagePt     EmailCreateRequestLanguage = "pt"
	EmailCreateRequestLanguagePtAo   EmailCreateRequestLanguage = "pt-ao"
	EmailCreateRequestLanguagePtBr   EmailCreateRequestLanguage = "pt-br"
	EmailCreateRequestLanguagePtCh   EmailCreateRequestLanguage = "pt-ch"
	EmailCreateRequestLanguagePtCv   EmailCreateRequestLanguage = "pt-cv"
	EmailCreateRequestLanguagePtGq   EmailCreateRequestLanguage = "pt-gq"
	EmailCreateRequestLanguagePtGw   EmailCreateRequestLanguage = "pt-gw"
	EmailCreateRequestLanguagePtLu   EmailCreateRequestLanguage = "pt-lu"
	EmailCreateRequestLanguagePtMo   EmailCreateRequestLanguage = "pt-mo"
	EmailCreateRequestLanguagePtMz   EmailCreateRequestLanguage = "pt-mz"
	EmailCreateRequestLanguagePtPt   EmailCreateRequestLanguage = "pt-pt"
	EmailCreateRequestLanguagePtSt   EmailCreateRequestLanguage = "pt-st"
	EmailCreateRequestLanguagePtTl   EmailCreateRequestLanguage = "pt-tl"
	EmailCreateRequestLanguageQu     EmailCreateRequestLanguage = "qu"
	EmailCreateRequestLanguageQuBo   EmailCreateRequestLanguage = "qu-bo"
	EmailCreateRequestLanguageQuEc   EmailCreateRequestLanguage = "qu-ec"
	EmailCreateRequestLanguageQuPe   EmailCreateRequestLanguage = "qu-pe"
	EmailCreateRequestLanguageRaj    EmailCreateRequestLanguage = "raj"
	EmailCreateRequestLanguageRajIn  EmailCreateRequestLanguage = "raj-in"
	EmailCreateRequestLanguageRm     EmailCreateRequestLanguage = "rm"
	EmailCreateRequestLanguageRmCh   EmailCreateRequestLanguage = "rm-ch"
	EmailCreateRequestLanguageRn     EmailCreateRequestLanguage = "rn"
	EmailCreateRequestLanguageRnBi   EmailCreateRequestLanguage = "rn-bi"
	EmailCreateRequestLanguageRo     EmailCreateRequestLanguage = "ro"
	EmailCreateRequestLanguageRoMd   EmailCreateRequestLanguage = "ro-md"
	EmailCreateRequestLanguageRoRo   EmailCreateRequestLanguage = "ro-ro"
	EmailCreateRequestLanguageRof    EmailCreateRequestLanguage = "rof"
	EmailCreateRequestLanguageRofTz  EmailCreateRequestLanguage = "rof-tz"
	EmailCreateRequestLanguageRu     EmailCreateRequestLanguage = "ru"
	EmailCreateRequestLanguageRuBy   EmailCreateRequestLanguage = "ru-by"
	EmailCreateRequestLanguageRuKg   EmailCreateRequestLanguage = "ru-kg"
	EmailCreateRequestLanguageRuKz   EmailCreateRequestLanguage = "ru-kz"
	EmailCreateRequestLanguageRuMd   EmailCreateRequestLanguage = "ru-md"
	EmailCreateRequestLanguageRuRu   EmailCreateRequestLanguage = "ru-ru"
	EmailCreateRequestLanguageRuUa   EmailCreateRequestLanguage = "ru-ua"
	EmailCreateRequestLanguageRw     EmailCreateRequestLanguage = "rw"
	EmailCreateRequestLanguageRwRw   EmailCreateRequestLanguage = "rw-rw"
	EmailCreateRequestLanguageRwk    EmailCreateRequestLanguage = "rwk"
	EmailCreateRequestLanguageRwkTz  EmailCreateRequestLanguage = "rwk-tz"
	EmailCreateRequestLanguageSa     EmailCreateRequestLanguage = "sa"
	EmailCreateRequestLanguageSaIn   EmailCreateRequestLanguage = "sa-in"
	EmailCreateRequestLanguageSah    EmailCreateRequestLanguage = "sah"
	EmailCreateRequestLanguageSahRu  EmailCreateRequestLanguage = "sah-ru"
	EmailCreateRequestLanguageSaq    EmailCreateRequestLanguage = "saq"
	EmailCreateRequestLanguageSaqKe  EmailCreateRequestLanguage = "saq-ke"
	EmailCreateRequestLanguageSat    EmailCreateRequestLanguage = "sat"
	EmailCreateRequestLanguageSatIn  EmailCreateRequestLanguage = "sat-in"
	EmailCreateRequestLanguageSbp    EmailCreateRequestLanguage = "sbp"
	EmailCreateRequestLanguageSbpTz  EmailCreateRequestLanguage = "sbp-tz"
	EmailCreateRequestLanguageSc     EmailCreateRequestLanguage = "sc"
	EmailCreateRequestLanguageScIt   EmailCreateRequestLanguage = "sc-it"
	EmailCreateRequestLanguageSd     EmailCreateRequestLanguage = "sd"
	EmailCreateRequestLanguageSdIn   EmailCreateRequestLanguage = "sd-in"
	EmailCreateRequestLanguageSdPk   EmailCreateRequestLanguage = "sd-pk"
	EmailCreateRequestLanguageSe     EmailCreateRequestLanguage = "se"
	EmailCreateRequestLanguageSeFi   EmailCreateRequestLanguage = "se-fi"
	EmailCreateRequestLanguageSeNo   EmailCreateRequestLanguage = "se-no"
	EmailCreateRequestLanguageSeSe   EmailCreateRequestLanguage = "se-se"
	EmailCreateRequestLanguageSeh    EmailCreateRequestLanguage = "seh"
	EmailCreateRequestLanguageSehMz  EmailCreateRequestLanguage = "seh-mz"
	EmailCreateRequestLanguageSes    EmailCreateRequestLanguage = "ses"
	EmailCreateRequestLanguageSesMl  EmailCreateRequestLanguage = "ses-ml"
	EmailCreateRequestLanguageSg     EmailCreateRequestLanguage = "sg"
	EmailCreateRequestLanguageSgCf   EmailCreateRequestLanguage = "sg-cf"
	EmailCreateRequestLanguageShi    EmailCreateRequestLanguage = "shi"
	EmailCreateRequestLanguageShiMa  EmailCreateRequestLanguage = "shi-ma"
	EmailCreateRequestLanguageSi     EmailCreateRequestLanguage = "si"
	EmailCreateRequestLanguageSiLk   EmailCreateRequestLanguage = "si-lk"
	EmailCreateRequestLanguageSk     EmailCreateRequestLanguage = "sk"
	EmailCreateRequestLanguageSkSk   EmailCreateRequestLanguage = "sk-sk"
	EmailCreateRequestLanguageSl     EmailCreateRequestLanguage = "sl"
	EmailCreateRequestLanguageSlSi   EmailCreateRequestLanguage = "sl-si"
	EmailCreateRequestLanguageSm     EmailCreateRequestLanguage = "sm"
	EmailCreateRequestLanguageSmn    EmailCreateRequestLanguage = "smn"
	EmailCreateRequestLanguageSmnFi  EmailCreateRequestLanguage = "smn-fi"
	EmailCreateRequestLanguageSMS    EmailCreateRequestLanguage = "sms"
	EmailCreateRequestLanguageSMSFi  EmailCreateRequestLanguage = "sms-fi"
	EmailCreateRequestLanguageSn     EmailCreateRequestLanguage = "sn"
	EmailCreateRequestLanguageSnZw   EmailCreateRequestLanguage = "sn-zw"
	EmailCreateRequestLanguageSo     EmailCreateRequestLanguage = "so"
	EmailCreateRequestLanguageSoDj   EmailCreateRequestLanguage = "so-dj"
	EmailCreateRequestLanguageSoEt   EmailCreateRequestLanguage = "so-et"
	EmailCreateRequestLanguageSoKe   EmailCreateRequestLanguage = "so-ke"
	EmailCreateRequestLanguageSoSo   EmailCreateRequestLanguage = "so-so"
	EmailCreateRequestLanguageSq     EmailCreateRequestLanguage = "sq"
	EmailCreateRequestLanguageSqAl   EmailCreateRequestLanguage = "sq-al"
	EmailCreateRequestLanguageSqMk   EmailCreateRequestLanguage = "sq-mk"
	EmailCreateRequestLanguageSqXk   EmailCreateRequestLanguage = "sq-xk"
	EmailCreateRequestLanguageSr     EmailCreateRequestLanguage = "sr"
	EmailCreateRequestLanguageSrBa   EmailCreateRequestLanguage = "sr-ba"
	EmailCreateRequestLanguageSrCs   EmailCreateRequestLanguage = "sr-cs"
	EmailCreateRequestLanguageSrMe   EmailCreateRequestLanguage = "sr-me"
	EmailCreateRequestLanguageSrRs   EmailCreateRequestLanguage = "sr-rs"
	EmailCreateRequestLanguageSrXk   EmailCreateRequestLanguage = "sr-xk"
	EmailCreateRequestLanguageSS     EmailCreateRequestLanguage = "ss"
	EmailCreateRequestLanguageSt     EmailCreateRequestLanguage = "st"
	EmailCreateRequestLanguageStLs   EmailCreateRequestLanguage = "st-ls"
	EmailCreateRequestLanguageStZa   EmailCreateRequestLanguage = "st-za"
	EmailCreateRequestLanguageSu     EmailCreateRequestLanguage = "su"
	EmailCreateRequestLanguageSuID   EmailCreateRequestLanguage = "su-id"
	EmailCreateRequestLanguageSv     EmailCreateRequestLanguage = "sv"
	EmailCreateRequestLanguageSvAx   EmailCreateRequestLanguage = "sv-ax"
	EmailCreateRequestLanguageSvFi   EmailCreateRequestLanguage = "sv-fi"
	EmailCreateRequestLanguageSvSe   EmailCreateRequestLanguage = "sv-se"
	EmailCreateRequestLanguageSw     EmailCreateRequestLanguage = "sw"
	EmailCreateRequestLanguageSwCd   EmailCreateRequestLanguage = "sw-cd"
	EmailCreateRequestLanguageSwKe   EmailCreateRequestLanguage = "sw-ke"
	EmailCreateRequestLanguageSwTz   EmailCreateRequestLanguage = "sw-tz"
	EmailCreateRequestLanguageSwUg   EmailCreateRequestLanguage = "sw-ug"
	EmailCreateRequestLanguageSy     EmailCreateRequestLanguage = "sy"
	EmailCreateRequestLanguageSyr    EmailCreateRequestLanguage = "syr"
	EmailCreateRequestLanguageSyrIq  EmailCreateRequestLanguage = "syr-iq"
	EmailCreateRequestLanguageSyrSy  EmailCreateRequestLanguage = "syr-sy"
	EmailCreateRequestLanguageSzl    EmailCreateRequestLanguage = "szl"
	EmailCreateRequestLanguageSzlPl  EmailCreateRequestLanguage = "szl-pl"
	EmailCreateRequestLanguageTa     EmailCreateRequestLanguage = "ta"
	EmailCreateRequestLanguageTaIn   EmailCreateRequestLanguage = "ta-in"
	EmailCreateRequestLanguageTaLk   EmailCreateRequestLanguage = "ta-lk"
	EmailCreateRequestLanguageTaMy   EmailCreateRequestLanguage = "ta-my"
	EmailCreateRequestLanguageTaSg   EmailCreateRequestLanguage = "ta-sg"
	EmailCreateRequestLanguageTe     EmailCreateRequestLanguage = "te"
	EmailCreateRequestLanguageTeIn   EmailCreateRequestLanguage = "te-in"
	EmailCreateRequestLanguageTeo    EmailCreateRequestLanguage = "teo"
	EmailCreateRequestLanguageTeoKe  EmailCreateRequestLanguage = "teo-ke"
	EmailCreateRequestLanguageTeoUg  EmailCreateRequestLanguage = "teo-ug"
	EmailCreateRequestLanguageTg     EmailCreateRequestLanguage = "tg"
	EmailCreateRequestLanguageTgTj   EmailCreateRequestLanguage = "tg-tj"
	EmailCreateRequestLanguageTh     EmailCreateRequestLanguage = "th"
	EmailCreateRequestLanguageThTh   EmailCreateRequestLanguage = "th-th"
	EmailCreateRequestLanguageTi     EmailCreateRequestLanguage = "ti"
	EmailCreateRequestLanguageTiEr   EmailCreateRequestLanguage = "ti-er"
	EmailCreateRequestLanguageTiEt   EmailCreateRequestLanguage = "ti-et"
	EmailCreateRequestLanguageTk     EmailCreateRequestLanguage = "tk"
	EmailCreateRequestLanguageTkTm   EmailCreateRequestLanguage = "tk-tm"
	EmailCreateRequestLanguageTl     EmailCreateRequestLanguage = "tl"
	EmailCreateRequestLanguageTn     EmailCreateRequestLanguage = "tn"
	EmailCreateRequestLanguageTnBw   EmailCreateRequestLanguage = "tn-bw"
	EmailCreateRequestLanguageTnZa   EmailCreateRequestLanguage = "tn-za"
	EmailCreateRequestLanguageTo     EmailCreateRequestLanguage = "to"
	EmailCreateRequestLanguageToTo   EmailCreateRequestLanguage = "to-to"
	EmailCreateRequestLanguageTok    EmailCreateRequestLanguage = "tok"
	EmailCreateRequestLanguageTok001 EmailCreateRequestLanguage = "tok-001"
	EmailCreateRequestLanguageTr     EmailCreateRequestLanguage = "tr"
	EmailCreateRequestLanguageTrCy   EmailCreateRequestLanguage = "tr-cy"
	EmailCreateRequestLanguageTrTr   EmailCreateRequestLanguage = "tr-tr"
	EmailCreateRequestLanguageTs     EmailCreateRequestLanguage = "ts"
	EmailCreateRequestLanguageTt     EmailCreateRequestLanguage = "tt"
	EmailCreateRequestLanguageTtRu   EmailCreateRequestLanguage = "tt-ru"
	EmailCreateRequestLanguageTw     EmailCreateRequestLanguage = "tw"
	EmailCreateRequestLanguageTwq    EmailCreateRequestLanguage = "twq"
	EmailCreateRequestLanguageTwqNe  EmailCreateRequestLanguage = "twq-ne"
	EmailCreateRequestLanguageTy     EmailCreateRequestLanguage = "ty"
	EmailCreateRequestLanguageTzm    EmailCreateRequestLanguage = "tzm"
	EmailCreateRequestLanguageTzmMa  EmailCreateRequestLanguage = "tzm-ma"
	EmailCreateRequestLanguageUg     EmailCreateRequestLanguage = "ug"
	EmailCreateRequestLanguageUgCn   EmailCreateRequestLanguage = "ug-cn"
	EmailCreateRequestLanguageUk     EmailCreateRequestLanguage = "uk"
	EmailCreateRequestLanguageUkUa   EmailCreateRequestLanguage = "uk-ua"
	EmailCreateRequestLanguageUr     EmailCreateRequestLanguage = "ur"
	EmailCreateRequestLanguageUrIn   EmailCreateRequestLanguage = "ur-in"
	EmailCreateRequestLanguageUrPk   EmailCreateRequestLanguage = "ur-pk"
	EmailCreateRequestLanguageUz     EmailCreateRequestLanguage = "uz"
	EmailCreateRequestLanguageUzAf   EmailCreateRequestLanguage = "uz-af"
	EmailCreateRequestLanguageUzUz   EmailCreateRequestLanguage = "uz-uz"
	EmailCreateRequestLanguageVai    EmailCreateRequestLanguage = "vai"
	EmailCreateRequestLanguageVaiLr  EmailCreateRequestLanguage = "vai-lr"
	EmailCreateRequestLanguageVe     EmailCreateRequestLanguage = "ve"
	EmailCreateRequestLanguageVec    EmailCreateRequestLanguage = "vec"
	EmailCreateRequestLanguageVecIt  EmailCreateRequestLanguage = "vec-it"
	EmailCreateRequestLanguageVi     EmailCreateRequestLanguage = "vi"
	EmailCreateRequestLanguageViVn   EmailCreateRequestLanguage = "vi-vn"
	EmailCreateRequestLanguageVmw    EmailCreateRequestLanguage = "vmw"
	EmailCreateRequestLanguageVmwMz  EmailCreateRequestLanguage = "vmw-mz"
	EmailCreateRequestLanguageVo     EmailCreateRequestLanguage = "vo"
	EmailCreateRequestLanguageVo001  EmailCreateRequestLanguage = "vo-001"
	EmailCreateRequestLanguageVun    EmailCreateRequestLanguage = "vun"
	EmailCreateRequestLanguageVunTz  EmailCreateRequestLanguage = "vun-tz"
	EmailCreateRequestLanguageWa     EmailCreateRequestLanguage = "wa"
	EmailCreateRequestLanguageWae    EmailCreateRequestLanguage = "wae"
	EmailCreateRequestLanguageWaeCh  EmailCreateRequestLanguage = "wae-ch"
	EmailCreateRequestLanguageWo     EmailCreateRequestLanguage = "wo"
	EmailCreateRequestLanguageWoSn   EmailCreateRequestLanguage = "wo-sn"
	EmailCreateRequestLanguageXh     EmailCreateRequestLanguage = "xh"
	EmailCreateRequestLanguageXhZa   EmailCreateRequestLanguage = "xh-za"
	EmailCreateRequestLanguageXnr    EmailCreateRequestLanguage = "xnr"
	EmailCreateRequestLanguageXnrIn  EmailCreateRequestLanguage = "xnr-in"
	EmailCreateRequestLanguageXog    EmailCreateRequestLanguage = "xog"
	EmailCreateRequestLanguageXogUg  EmailCreateRequestLanguage = "xog-ug"
	EmailCreateRequestLanguageYav    EmailCreateRequestLanguage = "yav"
	EmailCreateRequestLanguageYavCm  EmailCreateRequestLanguage = "yav-cm"
	EmailCreateRequestLanguageYi     EmailCreateRequestLanguage = "yi"
	EmailCreateRequestLanguageYi001  EmailCreateRequestLanguage = "yi-001"
	EmailCreateRequestLanguageYiUa   EmailCreateRequestLanguage = "yi-ua"
	EmailCreateRequestLanguageYo     EmailCreateRequestLanguage = "yo"
	EmailCreateRequestLanguageYoBj   EmailCreateRequestLanguage = "yo-bj"
	EmailCreateRequestLanguageYoNg   EmailCreateRequestLanguage = "yo-ng"
	EmailCreateRequestLanguageYrl    EmailCreateRequestLanguage = "yrl"
	EmailCreateRequestLanguageYrlBr  EmailCreateRequestLanguage = "yrl-br"
	EmailCreateRequestLanguageYrlCo  EmailCreateRequestLanguage = "yrl-co"
	EmailCreateRequestLanguageYrlVe  EmailCreateRequestLanguage = "yrl-ve"
	EmailCreateRequestLanguageYue    EmailCreateRequestLanguage = "yue"
	EmailCreateRequestLanguageYueCn  EmailCreateRequestLanguage = "yue-cn"
	EmailCreateRequestLanguageYueHk  EmailCreateRequestLanguage = "yue-hk"
	EmailCreateRequestLanguageYueMo  EmailCreateRequestLanguage = "yue-mo"
	EmailCreateRequestLanguageZa     EmailCreateRequestLanguage = "za"
	EmailCreateRequestLanguageZaCn   EmailCreateRequestLanguage = "za-cn"
	EmailCreateRequestLanguageZgh    EmailCreateRequestLanguage = "zgh"
	EmailCreateRequestLanguageZghMa  EmailCreateRequestLanguage = "zgh-ma"
	EmailCreateRequestLanguageZh     EmailCreateRequestLanguage = "zh"
	EmailCreateRequestLanguageZhCn   EmailCreateRequestLanguage = "zh-cn"
	EmailCreateRequestLanguageZhHans EmailCreateRequestLanguage = "zh-hans"
	EmailCreateRequestLanguageZhHant EmailCreateRequestLanguage = "zh-hant"
	EmailCreateRequestLanguageZhHk   EmailCreateRequestLanguage = "zh-hk"
	EmailCreateRequestLanguageZhMo   EmailCreateRequestLanguage = "zh-mo"
	EmailCreateRequestLanguageZhMy   EmailCreateRequestLanguage = "zh-my"
	EmailCreateRequestLanguageZhSg   EmailCreateRequestLanguage = "zh-sg"
	EmailCreateRequestLanguageZhTw   EmailCreateRequestLanguage = "zh-tw"
	EmailCreateRequestLanguageZu     EmailCreateRequestLanguage = "zu"
	EmailCreateRequestLanguageZuZa   EmailCreateRequestLanguage = "zu-za"
)

// The email state.
type EmailCreateRequestState string

const (
	EmailCreateRequestStateAgentGenerated          EmailCreateRequestState = "AGENT_GENERATED"
	EmailCreateRequestStateAutomated               EmailCreateRequestState = "AUTOMATED"
	EmailCreateRequestStateAutomatedAb             EmailCreateRequestState = "AUTOMATED_AB"
	EmailCreateRequestStateAutomatedAbVariant      EmailCreateRequestState = "AUTOMATED_AB_VARIANT"
	EmailCreateRequestStateAutomatedDraft          EmailCreateRequestState = "AUTOMATED_DRAFT"
	EmailCreateRequestStateAutomatedDraftAb        EmailCreateRequestState = "AUTOMATED_DRAFT_AB"
	EmailCreateRequestStateAutomatedDraftAbvariant EmailCreateRequestState = "AUTOMATED_DRAFT_ABVARIANT"
	EmailCreateRequestStateAutomatedForForm        EmailCreateRequestState = "AUTOMATED_FOR_FORM"
	EmailCreateRequestStateAutomatedForFormBuffer  EmailCreateRequestState = "AUTOMATED_FOR_FORM_BUFFER"
	EmailCreateRequestStateAutomatedForFormDraft   EmailCreateRequestState = "AUTOMATED_FOR_FORM_DRAFT"
	EmailCreateRequestStateAutomatedForFormLegacy  EmailCreateRequestState = "AUTOMATED_FOR_FORM_LEGACY"
	EmailCreateRequestStateAutomatedLoserAbvariant EmailCreateRequestState = "AUTOMATED_LOSER_ABVARIANT"
	EmailCreateRequestStateAutomatedSending        EmailCreateRequestState = "AUTOMATED_SENDING"
	EmailCreateRequestStateBlogEmailDraft          EmailCreateRequestState = "BLOG_EMAIL_DRAFT"
	EmailCreateRequestStateBlogEmailPublished      EmailCreateRequestState = "BLOG_EMAIL_PUBLISHED"
	EmailCreateRequestStateDraft                   EmailCreateRequestState = "DRAFT"
	EmailCreateRequestStateDraftAb                 EmailCreateRequestState = "DRAFT_AB"
	EmailCreateRequestStateDraftAbVariant          EmailCreateRequestState = "DRAFT_AB_VARIANT"
	EmailCreateRequestStateError                   EmailCreateRequestState = "ERROR"
	EmailCreateRequestStateLoserAbVariant          EmailCreateRequestState = "LOSER_AB_VARIANT"
	EmailCreateRequestStatePageStub                EmailCreateRequestState = "PAGE_STUB"
	EmailCreateRequestStatePreProcessing           EmailCreateRequestState = "PRE_PROCESSING"
	EmailCreateRequestStateProcessing              EmailCreateRequestState = "PROCESSING"
	EmailCreateRequestStatePublished               EmailCreateRequestState = "PUBLISHED"
	EmailCreateRequestStatePublishedAb             EmailCreateRequestState = "PUBLISHED_AB"
	EmailCreateRequestStatePublishedAbVariant      EmailCreateRequestState = "PUBLISHED_AB_VARIANT"
	EmailCreateRequestStatePublishedOrScheduled    EmailCreateRequestState = "PUBLISHED_OR_SCHEDULED"
	EmailCreateRequestStateRssToEmailDraft         EmailCreateRequestState = "RSS_TO_EMAIL_DRAFT"
	EmailCreateRequestStateRssToEmailPublished     EmailCreateRequestState = "RSS_TO_EMAIL_PUBLISHED"
	EmailCreateRequestStateScheduled               EmailCreateRequestState = "SCHEDULED"
	EmailCreateRequestStateScheduledAb             EmailCreateRequestState = "SCHEDULED_AB"
	EmailCreateRequestStateScheduledOrPublished    EmailCreateRequestState = "SCHEDULED_OR_PUBLISHED"
)

// The email subcategory.
type EmailCreateRequestSubcategory string

const (
	EmailCreateRequestSubcategoryAbLoserVariant                 EmailCreateRequestSubcategory = "ab_loser_variant"
	EmailCreateRequestSubcategoryAbLoserVariantSitePage         EmailCreateRequestSubcategory = "ab_loser_variant_site_page"
	EmailCreateRequestSubcategoryAbMaster                       EmailCreateRequestSubcategory = "ab_master"
	EmailCreateRequestSubcategoryAbMasterSitePage               EmailCreateRequestSubcategory = "ab_master_site_page"
	EmailCreateRequestSubcategoryAbVariant                      EmailCreateRequestSubcategory = "ab_variant"
	EmailCreateRequestSubcategoryAbVariantSitePage              EmailCreateRequestSubcategory = "ab_variant_site_page"
	EmailCreateRequestSubcategoryAutomated                      EmailCreateRequestSubcategory = "automated"
	EmailCreateRequestSubcategoryAutomatedAbMaster              EmailCreateRequestSubcategory = "automated_ab_master"
	EmailCreateRequestSubcategoryAutomatedAbVariant             EmailCreateRequestSubcategory = "automated_ab_variant"
	EmailCreateRequestSubcategoryAutomatedForCrm                EmailCreateRequestSubcategory = "automated_for_crm"
	EmailCreateRequestSubcategoryAutomatedForCustomSurvey       EmailCreateRequestSubcategory = "automated_for_custom_survey"
	EmailCreateRequestSubcategoryAutomatedForDeal               EmailCreateRequestSubcategory = "automated_for_deal"
	EmailCreateRequestSubcategoryAutomatedForFeedbackCes        EmailCreateRequestSubcategory = "automated_for_feedback_ces"
	EmailCreateRequestSubcategoryAutomatedForFeedbackCustom     EmailCreateRequestSubcategory = "automated_for_feedback_custom"
	EmailCreateRequestSubcategoryAutomatedForFeedbackNps        EmailCreateRequestSubcategory = "automated_for_feedback_nps"
	EmailCreateRequestSubcategoryAutomatedForForm               EmailCreateRequestSubcategory = "automated_for_form"
	EmailCreateRequestSubcategoryAutomatedForFormBuffer         EmailCreateRequestSubcategory = "automated_for_form_buffer"
	EmailCreateRequestSubcategoryAutomatedForFormDraft          EmailCreateRequestSubcategory = "automated_for_form_draft"
	EmailCreateRequestSubcategoryAutomatedForFormLegacy         EmailCreateRequestSubcategory = "automated_for_form_legacy"
	EmailCreateRequestSubcategoryAutomatedForLeadflow           EmailCreateRequestSubcategory = "automated_for_leadflow"
	EmailCreateRequestSubcategoryAutomatedForTicket             EmailCreateRequestSubcategory = "automated_for_ticket"
	EmailCreateRequestSubcategoryBatch                          EmailCreateRequestSubcategory = "batch"
	EmailCreateRequestSubcategoryBlogArticleInstanceLayout      EmailCreateRequestSubcategory = "blog_article_instance_layout"
	EmailCreateRequestSubcategoryBlogArticleListing             EmailCreateRequestSubcategory = "blog_article_listing"
	EmailCreateRequestSubcategoryBlogAuthorDetail               EmailCreateRequestSubcategory = "blog_author_detail"
	EmailCreateRequestSubcategoryBlogEmail                      EmailCreateRequestSubcategory = "blog_email"
	EmailCreateRequestSubcategoryBlogEmailChild                 EmailCreateRequestSubcategory = "blog_email_child"
	EmailCreateRequestSubcategoryCaseStudy                      EmailCreateRequestSubcategory = "case_study"
	EmailCreateRequestSubcategoryCaseStudyInstanceLayout        EmailCreateRequestSubcategory = "case_study_instance_layout"
	EmailCreateRequestSubcategoryCaseStudyListing               EmailCreateRequestSubcategory = "case_study_listing"
	EmailCreateRequestSubcategoryDiscardableStub                EmailCreateRequestSubcategory = "discardable_stub"
	EmailCreateRequestSubcategoryImportedBlogPost               EmailCreateRequestSubcategory = "imported_blog_post"
	EmailCreateRequestSubcategoryKB404Page                      EmailCreateRequestSubcategory = "kb_404_page"
	EmailCreateRequestSubcategoryKBArticleInstanceLayout        EmailCreateRequestSubcategory = "kb_article_instance_layout"
	EmailCreateRequestSubcategoryKBListing                      EmailCreateRequestSubcategory = "kb_listing"
	EmailCreateRequestSubcategoryKBSearchResults                EmailCreateRequestSubcategory = "kb_search_results"
	EmailCreateRequestSubcategoryKBSupportForm                  EmailCreateRequestSubcategory = "kb_support_form"
	EmailCreateRequestSubcategoryLandingPage                    EmailCreateRequestSubcategory = "landing_page"
	EmailCreateRequestSubcategoryLegacyBlogPost                 EmailCreateRequestSubcategory = "legacy_blog_post"
	EmailCreateRequestSubcategoryLegacyPage                     EmailCreateRequestSubcategory = "legacy_page"
	EmailCreateRequestSubcategoryLocaltime                      EmailCreateRequestSubcategory = "localtime"
	EmailCreateRequestSubcategoryManagePreferencesEmail         EmailCreateRequestSubcategory = "manage_preferences_email"
	EmailCreateRequestSubcategoryMarketingSingleSendAPI         EmailCreateRequestSubcategory = "marketing_single_send_api"
	EmailCreateRequestSubcategoryMembershipEmailVerification    EmailCreateRequestSubcategory = "membership_email_verification"
	EmailCreateRequestSubcategoryMembershipFollowUp             EmailCreateRequestSubcategory = "membership_follow_up"
	EmailCreateRequestSubcategoryMembershipOtpLogin             EmailCreateRequestSubcategory = "membership_otp_login"
	EmailCreateRequestSubcategoryMembershipPasswordReset        EmailCreateRequestSubcategory = "membership_password_reset"
	EmailCreateRequestSubcategoryMembershipPasswordSaved        EmailCreateRequestSubcategory = "membership_password_saved"
	EmailCreateRequestSubcategoryMembershipPasswordlessAuth     EmailCreateRequestSubcategory = "membership_passwordless_auth"
	EmailCreateRequestSubcategoryMembershipRegistration         EmailCreateRequestSubcategory = "membership_registration"
	EmailCreateRequestSubcategoryMembershipRegistrationFollowUp EmailCreateRequestSubcategory = "membership_registration_follow_up"
	EmailCreateRequestSubcategoryMembershipVerification         EmailCreateRequestSubcategory = "membership_verification"
	EmailCreateRequestSubcategoryNormalBlogPost                 EmailCreateRequestSubcategory = "normal_blog_post"
	EmailCreateRequestSubcategoryOptinEmail                     EmailCreateRequestSubcategory = "optin_email"
	EmailCreateRequestSubcategoryOptinFollowupEmail             EmailCreateRequestSubcategory = "optin_followup_email"
	EmailCreateRequestSubcategoryPageInstanceLayout             EmailCreateRequestSubcategory = "page_instance_layout"
	EmailCreateRequestSubcategoryPageStub                       EmailCreateRequestSubcategory = "page_stub"
	EmailCreateRequestSubcategoryPerformableLandingPage         EmailCreateRequestSubcategory = "performable_landing_page"
	EmailCreateRequestSubcategoryPerformableLandingPageCutover  EmailCreateRequestSubcategory = "performable_landing_page_cutover"
	EmailCreateRequestSubcategoryPodcastInstanceLayout          EmailCreateRequestSubcategory = "podcast_instance_layout"
	EmailCreateRequestSubcategoryPodcastListing                 EmailCreateRequestSubcategory = "podcast_listing"
	EmailCreateRequestSubcategoryPortalContent                  EmailCreateRequestSubcategory = "portal_content"
	EmailCreateRequestSubcategoryResubscribeConfirmationEmail   EmailCreateRequestSubcategory = "resubscribe_confirmation_email"
	EmailCreateRequestSubcategoryResubscribeEmail               EmailCreateRequestSubcategory = "resubscribe_email"
	EmailCreateRequestSubcategoryRssToEmail                     EmailCreateRequestSubcategory = "rss_to_email"
	EmailCreateRequestSubcategoryRssToEmailChild                EmailCreateRequestSubcategory = "rss_to_email_child"
	EmailCreateRequestSubcategoryScpInstanceLayoutPage          EmailCreateRequestSubcategory = "scp_instance_layout_page"
	EmailCreateRequestSubcategoryScpStaticPage                  EmailCreateRequestSubcategory = "scp_static_page"
	EmailCreateRequestSubcategorySingleSendAPI                  EmailCreateRequestSubcategory = "single_send_api"
	EmailCreateRequestSubcategorySitePage                       EmailCreateRequestSubcategory = "site_page"
	EmailCreateRequestSubcategorySmtpToken                      EmailCreateRequestSubcategory = "smtp_token"
	EmailCreateRequestSubcategoryStagedPage                     EmailCreateRequestSubcategory = "staged_page"
	EmailCreateRequestSubcategoryTicketClosedKickbackEmail      EmailCreateRequestSubcategory = "ticket_closed_kickback_email"
	EmailCreateRequestSubcategoryTicketOpenedKickbackEmail      EmailCreateRequestSubcategory = "ticket_opened_kickback_email"
	EmailCreateRequestSubcategoryTicketPipelineAutomated        EmailCreateRequestSubcategory = "ticket_pipeline_automated"
	EmailCreateRequestSubcategoryUnknown                        EmailCreateRequestSubcategory = "UNKNOWN"
	EmailCreateRequestSubcategoryUnsubscribeConfirmationEmail   EmailCreateRequestSubcategory = "unsubscribe_confirmation_email"
	EmailCreateRequestSubcategoryWebInteractive                 EmailCreateRequestSubcategory = "web_interactive"
)

type EmailStatisticInterval struct {
	Aggregations EmailStatisticsData `json:"aggregations" api:"required"`
	Interval     Interval            `json:"interval" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Aggregations respjson.Field
		Interval     respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmailStatisticInterval) RawJSON() string { return r.JSON.raw }
func (r *EmailStatisticInterval) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmailStatisticsData struct {
	// Counters like number of `sent`, `open` or `delivered`.
	Counters map[string]int64 `json:"counters" api:"required"`
	// Statistics by device.
	DeviceBreakdown map[string]map[string]int64 `json:"deviceBreakdown" api:"required"`
	// Number of emails that were dropped and bounced.
	QualifierStats map[string]map[string]int64 `json:"qualifierStats" api:"required"`
	// Ratios like `openratio` or `clickratio`
	Ratios map[string]float64 `json:"ratios" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Counters        respjson.Field
		DeviceBreakdown respjson.Field
		QualifierStats  respjson.Field
		Ratios          respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmailStatisticsData) RawJSON() string { return r.JSON.raw }
func (r *EmailStatisticsData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmailUpdateRequestParam struct {
	// The active domain of the email.
	ActiveDomain param.Opt[string] `json:"activeDomain,omitzero"`
	// Determines if the email is archived or not.
	Archived param.Opt[bool] `json:"archived,omitzero"`
	// The ID of the business unit associated with the email.
	BusinessUnitID param.Opt[int64] `json:"businessUnitId,omitzero"`
	// The ID of the campaign this email is associated to.
	Campaign param.Opt[string] `json:"campaign,omitzero"`
	// The ID of the folder where the email will be stored.
	FolderIDV2 param.Opt[int64] `json:"folderIdV2,omitzero"`
	// Determines whether the email send time should be randomized to avoid sending all
	// emails at the exact same time.
	JitterSendTime param.Opt[bool] `json:"jitterSendTime,omitzero"`
	// The name of the email, as displayed on the email dashboard.
	Name param.Opt[string] `json:"name,omitzero"`
	// The date and time the email is scheduled for, in ISO8601 representation. This is
	// only used in local time or scheduled emails.
	PublishDate param.Opt[time.Time] `json:"publishDate,omitzero" format:"date-time"`
	// Determines whether the email will be sent immediately on publish.
	SendOnPublish param.Opt[bool] `json:"sendOnPublish,omitzero"`
	// The subject of the email.
	Subject param.Opt[string]           `json:"subject,omitzero"`
	Content PublicEmailContentParam     `json:"content,omitzero"`
	From    PublicEmailFromDetailsParam `json:"from,omitzero"`
	// The language code for the email, such as 'en' for English.
	//
	// Any of "aa", "ab", "ae", "af", "af-na", "af-za", "agq", "agq-cm", "ak", "ak-gh",
	// "am", "am-et", "an", "ann", "ann-ng", "ar", "ar-001", "ar-ae", "ar-bh", "ar-dj",
	// "ar-dz", "ar-eg", "ar-eh", "ar-er", "ar-il", "ar-iq", "ar-jo", "ar-km", "ar-kw",
	// "ar-lb", "ar-ly", "ar-ma", "ar-mr", "ar-om", "ar-ps", "ar-qa", "ar-sa", "ar-sd",
	// "ar-so", "ar-ss", "ar-sy", "ar-td", "ar-tn", "ar-ye", "as", "as-in", "asa",
	// "asa-tz", "ast", "ast-es", "av", "ay", "az", "az-az", "ba", "bal", "bal-pk",
	// "bas", "bas-cm", "be", "be-by", "bem", "bem-zm", "bez", "bez-tz", "bg", "bg-bg",
	// "bgc", "bgc-in", "bho", "bho-in", "bi", "blo", "blo-bj", "bm", "bm-ml", "bn",
	// "bn-bd", "bn-in", "bo", "bo-cn", "bo-in", "br", "br-fr", "brx", "brx-in", "bs",
	// "bs-ba", "ca", "ca-ad", "ca-es", "ca-fr", "ca-it", "ccp", "ccp-bd", "ccp-in",
	// "ce", "ce-ru", "ceb", "ceb-ph", "cgg", "cgg-ug", "ch", "chr", "chr-us", "ckb",
	// "ckb-iq", "ckb-ir", "co", "cr", "cs", "cs-cz", "csw", "csw-ca", "cu", "cu-ru",
	// "cv", "cv-ru", "cy", "cy-gb", "da", "da-dk", "da-gl", "dav", "dav-ke", "de",
	// "de-at", "de-be", "de-ch", "de-de", "de-gr", "de-it", "de-li", "de-lu", "dje",
	// "dje-ne", "doi", "doi-in", "dsb", "dsb-de", "dua", "dua-cm", "dv", "dyo",
	// "dyo-sn", "dz", "dz-bt", "ebu", "ebu-ke", "ee", "ee-gh", "ee-tg", "el", "el-cy",
	// "el-gr", "en", "en-001", "en-150", "en-ae", "en-ag", "en-ai", "en-as", "en-at",
	// "en-au", "en-bb", "en-be", "en-bi", "en-bm", "en-bs", "en-bw", "en-bz", "en-ca",
	// "en-cc", "en-ch", "en-ck", "en-cm", "en-cn", "en-cx", "en-cy", "en-cz", "en-de",
	// "en-dg", "en-dk", "en-dm", "en-ee", "en-eg", "en-er", "en-es", "en-fi", "en-fj",
	// "en-fk", "en-fm", "en-fr", "en-gb", "en-gd", "en-gg", "en-gh", "en-gi", "en-gm",
	// "en-gs", "en-gu", "en-gy", "en-hk", "en-hu", "en-id", "en-ie", "en-il", "en-im",
	// "en-in", "en-io", "en-it", "en-je", "en-jm", "en-ke", "en-ki", "en-kn", "en-ky",
	// "en-lc", "en-lr", "en-ls", "en-lu", "en-mg", "en-mh", "en-mo", "en-mp", "en-ms",
	// "en-mt", "en-mu", "en-mv", "en-mw", "en-mx", "en-my", "en-na", "en-nf", "en-ng",
	// "en-nl", "en-no", "en-nr", "en-nu", "en-nz", "en-pg", "en-ph", "en-pk", "en-pl",
	// "en-pn", "en-pr", "en-pt", "en-pw", "en-ro", "en-rw", "en-sb", "en-sc", "en-sd",
	// "en-se", "en-sg", "en-sh", "en-si", "en-sk", "en-sl", "en-ss", "en-sx", "en-sz",
	// "en-tc", "en-th", "en-tk", "en-tn", "en-to", "en-tt", "en-tv", "en-tz", "en-ug",
	// "en-um", "en-us", "en-vc", "en-vg", "en-vi", "en-vn", "en-vu", "en-ws", "en-za",
	// "en-zm", "en-zw", "eo", "eo-001", "es", "es-419", "es-ar", "es-bo", "es-br",
	// "es-bz", "es-cl", "es-co", "es-cr", "es-cu", "es-do", "es-ea", "es-ec", "es-es",
	// "es-gq", "es-gt", "es-hn", "es-ic", "es-mx", "es-ni", "es-pa", "es-pe", "es-ph",
	// "es-pr", "es-py", "es-sv", "es-us", "es-uy", "es-ve", "et", "et-ee", "eu",
	// "eu-es", "ewo", "ewo-cm", "fa", "fa-af", "fa-ir", "ff", "ff-bf", "ff-cm",
	// "ff-gh", "ff-gm", "ff-gn", "ff-gw", "ff-lr", "ff-mr", "ff-ne", "ff-ng", "ff-sl",
	// "ff-sn", "fi", "fi-fi", "fil", "fil-ph", "fj", "fo", "fo-dk", "fo-fo", "fr",
	// "fr-be", "fr-bf", "fr-bi", "fr-bj", "fr-bl", "fr-ca", "fr-cd", "fr-cf", "fr-cg",
	// "fr-ch", "fr-ci", "fr-cm", "fr-dj", "fr-dz", "fr-fr", "fr-ga", "fr-gf", "fr-gn",
	// "fr-gp", "fr-gq", "fr-ht", "fr-km", "fr-lu", "fr-ma", "fr-mc", "fr-mf", "fr-mg",
	// "fr-ml", "fr-mq", "fr-mr", "fr-mu", "fr-nc", "fr-ne", "fr-pf", "fr-pm", "fr-re",
	// "fr-rw", "fr-sc", "fr-sn", "fr-sy", "fr-td", "fr-tg", "fr-tn", "fr-vu", "fr-wf",
	// "fr-yt", "frr", "frr-de", "fur", "fur-it", "fy", "fy-nl", "ga", "ga-gb",
	// "ga-ie", "gaa", "gaa-gh", "gd", "gd-gb", "gl", "gl-es", "gn", "gsw", "gsw-ch",
	// "gsw-fr", "gsw-li", "gu", "gu-in", "guz", "guz-ke", "gv", "gv-im", "ha",
	// "ha-gh", "ha-ne", "ha-ng", "haw", "haw-us", "he", "he-il", "hi", "hi-in", "hmn",
	// "ho", "hr", "hr-ba", "hr-hr", "hsb", "hsb-de", "ht", "ht-ht", "hu", "hu-hu",
	// "hy", "hy-am", "hz", "ia", "ia-001", "id", "id-id", "ie", "ie-ee", "ig",
	// "ig-ng", "ii", "ii-cn", "ik", "io", "is", "is-is", "it", "it-ch", "it-it",
	// "it-sm", "it-va", "iu", "ja", "ja-jp", "jgo", "jgo-cm", "jmc", "jmc-tz", "jv",
	// "jv-id", "ka", "ka-ge", "kab", "kab-dz", "kam", "kam-ke", "kar", "kde",
	// "kde-tz", "kea", "kea-cv", "kg", "kgp", "kgp-br", "kh", "khq", "khq-ml", "ki",
	// "ki-ke", "kj", "kk", "kk-kz", "kkj", "kkj-cm", "kl", "kl-gl", "kln", "kln-ke",
	// "km", "km-kh", "kn", "kn-in", "ko", "ko-cn", "ko-kp", "ko-kr", "kok", "kok-in",
	// "kr", "ks", "ks-in", "ksb", "ksb-tz", "ksf", "ksf-cm", "ksh", "ksh-de", "ku",
	// "ku-tr", "kv", "kw", "kw-gb", "kxv", "kxv-in", "ky", "ky-kg", "la", "lag",
	// "lag-tz", "lb", "lb-lu", "lg", "lg-ug", "li", "lij", "lij-it", "lkt", "lkt-us",
	// "lmo", "lmo-it", "ln", "ln-ao", "ln-cd", "ln-cf", "ln-cg", "lo", "lo-la", "lrc",
	// "lrc-iq", "lrc-ir", "lt", "lt-lt", "lu", "lu-cd", "luo", "luo-ke", "luy",
	// "luy-ke", "lv", "lv-lv", "mai", "mai-in", "mas", "mas-ke", "mas-tz", "mdf",
	// "mdf-ru", "mer", "mer-ke", "mfe", "mfe-mu", "mg", "mg-mg", "mgh", "mgh-mz",
	// "mgo", "mgo-cm", "mh", "mi", "mi-nz", "mk", "mk-mk", "ml", "ml-in", "mn",
	// "mn-mn", "mni", "mni-in", "mr", "mr-in", "ms", "ms-bn", "ms-id", "ms-my",
	// "ms-sg", "mt", "mt-mt", "mua", "mua-cm", "my", "my-mm", "mzn", "mzn-ir", "na",
	// "naq", "naq-na", "nb", "nb-no", "nb-sj", "nd", "nd-zw", "nds", "nds-de",
	// "nds-nl", "ne", "ne-in", "ne-np", "ng", "nl", "nl-aw", "nl-be", "nl-bq",
	// "nl-ch", "nl-cw", "nl-lu", "nl-nl", "nl-sr", "nl-sx", "nmg", "nmg-cm", "nn",
	// "nn-no", "nnh", "nnh-cm", "no", "no-no", "nqo", "nqo-gn", "nr", "nso", "nso-za",
	// "nus", "nus-ss", "nv", "ny", "nyn", "nyn-ug", "oc", "oc-es", "oc-fr", "oj",
	// "om", "om-et", "om-ke", "or", "or-in", "os", "os-ge", "os-ru", "pa", "pa-in",
	// "pa-pk", "pcm", "pcm-ng", "pi", "pis", "pis-sb", "pl", "pl-pl", "prg",
	// "prg-001", "ps", "ps-af", "ps-pk", "pt", "pt-ao", "pt-br", "pt-ch", "pt-cv",
	// "pt-gq", "pt-gw", "pt-lu", "pt-mo", "pt-mz", "pt-pt", "pt-st", "pt-tl", "qu",
	// "qu-bo", "qu-ec", "qu-pe", "raj", "raj-in", "rm", "rm-ch", "rn", "rn-bi", "ro",
	// "ro-md", "ro-ro", "rof", "rof-tz", "ru", "ru-by", "ru-kg", "ru-kz", "ru-md",
	// "ru-ru", "ru-ua", "rw", "rw-rw", "rwk", "rwk-tz", "sa", "sa-in", "sah",
	// "sah-ru", "saq", "saq-ke", "sat", "sat-in", "sbp", "sbp-tz", "sc", "sc-it",
	// "sd", "sd-in", "sd-pk", "se", "se-fi", "se-no", "se-se", "seh", "seh-mz", "ses",
	// "ses-ml", "sg", "sg-cf", "shi", "shi-ma", "si", "si-lk", "sk", "sk-sk", "sl",
	// "sl-si", "sm", "smn", "smn-fi", "sms", "sms-fi", "sn", "sn-zw", "so", "so-dj",
	// "so-et", "so-ke", "so-so", "sq", "sq-al", "sq-mk", "sq-xk", "sr", "sr-ba",
	// "sr-cs", "sr-me", "sr-rs", "sr-xk", "ss", "st", "st-ls", "st-za", "su", "su-id",
	// "sv", "sv-ax", "sv-fi", "sv-se", "sw", "sw-cd", "sw-ke", "sw-tz", "sw-ug", "sy",
	// "syr", "syr-iq", "syr-sy", "szl", "szl-pl", "ta", "ta-in", "ta-lk", "ta-my",
	// "ta-sg", "te", "te-in", "teo", "teo-ke", "teo-ug", "tg", "tg-tj", "th", "th-th",
	// "ti", "ti-er", "ti-et", "tk", "tk-tm", "tl", "tn", "tn-bw", "tn-za", "to",
	// "to-to", "tok", "tok-001", "tr", "tr-cy", "tr-tr", "ts", "tt", "tt-ru", "tw",
	// "twq", "twq-ne", "ty", "tzm", "tzm-ma", "ug", "ug-cn", "uk", "uk-ua", "ur",
	// "ur-in", "ur-pk", "uz", "uz-af", "uz-uz", "vai", "vai-lr", "ve", "vec",
	// "vec-it", "vi", "vi-vn", "vmw", "vmw-mz", "vo", "vo-001", "vun", "vun-tz", "wa",
	// "wae", "wae-ch", "wo", "wo-sn", "xh", "xh-za", "xnr", "xnr-in", "xog", "xog-ug",
	// "yav", "yav-cm", "yi", "yi-001", "yi-ua", "yo", "yo-bj", "yo-ng", "yrl",
	// "yrl-br", "yrl-co", "yrl-ve", "yue", "yue-cn", "yue-hk", "yue-mo", "za",
	// "za-cn", "zgh", "zgh-ma", "zh", "zh-cn", "zh-hans", "zh-hant", "zh-hk", "zh-mo",
	// "zh-my", "zh-sg", "zh-tw", "zu", "zu-za".
	Language EmailUpdateRequestLanguage `json:"language,omitzero"`
	RssData  PublicRssEmailDetailsParam `json:"rssData,omitzero"`
	// The email state.
	//
	// Any of "AGENT_GENERATED", "AUTOMATED", "AUTOMATED_AB", "AUTOMATED_AB_VARIANT",
	// "AUTOMATED_DRAFT", "AUTOMATED_DRAFT_AB", "AUTOMATED_DRAFT_ABVARIANT",
	// "AUTOMATED_FOR_FORM", "AUTOMATED_FOR_FORM_BUFFER", "AUTOMATED_FOR_FORM_DRAFT",
	// "AUTOMATED_FOR_FORM_LEGACY", "AUTOMATED_LOSER_ABVARIANT", "AUTOMATED_SENDING",
	// "BLOG_EMAIL_DRAFT", "BLOG_EMAIL_PUBLISHED", "DRAFT", "DRAFT_AB",
	// "DRAFT_AB_VARIANT", "ERROR", "LOSER_AB_VARIANT", "PAGE_STUB", "PRE_PROCESSING",
	// "PROCESSING", "PUBLISHED", "PUBLISHED_AB", "PUBLISHED_AB_VARIANT",
	// "PUBLISHED_OR_SCHEDULED", "RSS_TO_EMAIL_DRAFT", "RSS_TO_EMAIL_PUBLISHED",
	// "SCHEDULED", "SCHEDULED_AB", "SCHEDULED_OR_PUBLISHED".
	State EmailUpdateRequestState `json:"state,omitzero"`
	// The email subcategory.
	//
	// Any of "ab_loser_variant", "ab_loser_variant_site_page", "ab_master",
	// "ab_master_site_page", "ab_variant", "ab_variant_site_page", "automated",
	// "automated_ab_master", "automated_ab_variant", "automated_for_crm",
	// "automated_for_custom_survey", "automated_for_deal",
	// "automated_for_feedback_ces", "automated_for_feedback_custom",
	// "automated_for_feedback_nps", "automated_for_form", "automated_for_form_buffer",
	// "automated_for_form_draft", "automated_for_form_legacy",
	// "automated_for_leadflow", "automated_for_ticket", "batch",
	// "blog_article_instance_layout", "blog_article_listing", "blog_author_detail",
	// "blog_email", "blog_email_child", "case_study", "case_study_instance_layout",
	// "case_study_listing", "discardable_stub", "imported_blog_post", "kb_404_page",
	// "kb_article_instance_layout", "kb_listing", "kb_search_results",
	// "kb_support_form", "landing_page", "legacy_blog_post", "legacy_page",
	// "localtime", "manage_preferences_email", "marketing_single_send_api",
	// "membership_email_verification", "membership_follow_up", "membership_otp_login",
	// "membership_password_reset", "membership_password_saved",
	// "membership_passwordless_auth", "membership_registration",
	// "membership_registration_follow_up", "membership_verification",
	// "normal_blog_post", "optin_email", "optin_followup_email",
	// "page_instance_layout", "page_stub", "performable_landing_page",
	// "performable_landing_page_cutover", "podcast_instance_layout",
	// "podcast_listing", "portal_content", "resubscribe_confirmation_email",
	// "resubscribe_email", "rss_to_email", "rss_to_email_child",
	// "scp_instance_layout_page", "scp_static_page", "single_send_api", "site_page",
	// "smtp_token", "staged_page", "ticket_closed_kickback_email",
	// "ticket_opened_kickback_email", "ticket_pipeline_automated", "UNKNOWN",
	// "unsubscribe_confirmation_email", "web_interactive".
	Subcategory         EmailUpdateRequestSubcategory       `json:"subcategory,omitzero"`
	SubscriptionDetails PublicEmailSubscriptionDetailsParam `json:"subscriptionDetails,omitzero"`
	Testing             PublicEmailTestingDetailsParam      `json:"testing,omitzero"`
	To                  PublicEmailToDetailsParam           `json:"to,omitzero"`
	Webversion          PublicWebversionDetailsParam        `json:"webversion,omitzero"`
	paramObj
}

func (r EmailUpdateRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow EmailUpdateRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EmailUpdateRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The language code for the email, such as 'en' for English.
type EmailUpdateRequestLanguage string

const (
	EmailUpdateRequestLanguageAa     EmailUpdateRequestLanguage = "aa"
	EmailUpdateRequestLanguageAb     EmailUpdateRequestLanguage = "ab"
	EmailUpdateRequestLanguageAe     EmailUpdateRequestLanguage = "ae"
	EmailUpdateRequestLanguageAf     EmailUpdateRequestLanguage = "af"
	EmailUpdateRequestLanguageAfNa   EmailUpdateRequestLanguage = "af-na"
	EmailUpdateRequestLanguageAfZa   EmailUpdateRequestLanguage = "af-za"
	EmailUpdateRequestLanguageAgq    EmailUpdateRequestLanguage = "agq"
	EmailUpdateRequestLanguageAgqCm  EmailUpdateRequestLanguage = "agq-cm"
	EmailUpdateRequestLanguageAk     EmailUpdateRequestLanguage = "ak"
	EmailUpdateRequestLanguageAkGh   EmailUpdateRequestLanguage = "ak-gh"
	EmailUpdateRequestLanguageAm     EmailUpdateRequestLanguage = "am"
	EmailUpdateRequestLanguageAmEt   EmailUpdateRequestLanguage = "am-et"
	EmailUpdateRequestLanguageAn     EmailUpdateRequestLanguage = "an"
	EmailUpdateRequestLanguageAnn    EmailUpdateRequestLanguage = "ann"
	EmailUpdateRequestLanguageAnnNg  EmailUpdateRequestLanguage = "ann-ng"
	EmailUpdateRequestLanguageAr     EmailUpdateRequestLanguage = "ar"
	EmailUpdateRequestLanguageAr001  EmailUpdateRequestLanguage = "ar-001"
	EmailUpdateRequestLanguageArAe   EmailUpdateRequestLanguage = "ar-ae"
	EmailUpdateRequestLanguageArBh   EmailUpdateRequestLanguage = "ar-bh"
	EmailUpdateRequestLanguageArDj   EmailUpdateRequestLanguage = "ar-dj"
	EmailUpdateRequestLanguageArDz   EmailUpdateRequestLanguage = "ar-dz"
	EmailUpdateRequestLanguageArEg   EmailUpdateRequestLanguage = "ar-eg"
	EmailUpdateRequestLanguageArEh   EmailUpdateRequestLanguage = "ar-eh"
	EmailUpdateRequestLanguageArEr   EmailUpdateRequestLanguage = "ar-er"
	EmailUpdateRequestLanguageArIl   EmailUpdateRequestLanguage = "ar-il"
	EmailUpdateRequestLanguageArIq   EmailUpdateRequestLanguage = "ar-iq"
	EmailUpdateRequestLanguageArJo   EmailUpdateRequestLanguage = "ar-jo"
	EmailUpdateRequestLanguageArKm   EmailUpdateRequestLanguage = "ar-km"
	EmailUpdateRequestLanguageArKw   EmailUpdateRequestLanguage = "ar-kw"
	EmailUpdateRequestLanguageArLb   EmailUpdateRequestLanguage = "ar-lb"
	EmailUpdateRequestLanguageArLy   EmailUpdateRequestLanguage = "ar-ly"
	EmailUpdateRequestLanguageArMa   EmailUpdateRequestLanguage = "ar-ma"
	EmailUpdateRequestLanguageArMr   EmailUpdateRequestLanguage = "ar-mr"
	EmailUpdateRequestLanguageArOm   EmailUpdateRequestLanguage = "ar-om"
	EmailUpdateRequestLanguageArPs   EmailUpdateRequestLanguage = "ar-ps"
	EmailUpdateRequestLanguageArQa   EmailUpdateRequestLanguage = "ar-qa"
	EmailUpdateRequestLanguageArSa   EmailUpdateRequestLanguage = "ar-sa"
	EmailUpdateRequestLanguageArSd   EmailUpdateRequestLanguage = "ar-sd"
	EmailUpdateRequestLanguageArSo   EmailUpdateRequestLanguage = "ar-so"
	EmailUpdateRequestLanguageArSS   EmailUpdateRequestLanguage = "ar-ss"
	EmailUpdateRequestLanguageArSy   EmailUpdateRequestLanguage = "ar-sy"
	EmailUpdateRequestLanguageArTd   EmailUpdateRequestLanguage = "ar-td"
	EmailUpdateRequestLanguageArTn   EmailUpdateRequestLanguage = "ar-tn"
	EmailUpdateRequestLanguageArYe   EmailUpdateRequestLanguage = "ar-ye"
	EmailUpdateRequestLanguageAs     EmailUpdateRequestLanguage = "as"
	EmailUpdateRequestLanguageAsIn   EmailUpdateRequestLanguage = "as-in"
	EmailUpdateRequestLanguageAsa    EmailUpdateRequestLanguage = "asa"
	EmailUpdateRequestLanguageAsaTz  EmailUpdateRequestLanguage = "asa-tz"
	EmailUpdateRequestLanguageAst    EmailUpdateRequestLanguage = "ast"
	EmailUpdateRequestLanguageAstEs  EmailUpdateRequestLanguage = "ast-es"
	EmailUpdateRequestLanguageAv     EmailUpdateRequestLanguage = "av"
	EmailUpdateRequestLanguageAy     EmailUpdateRequestLanguage = "ay"
	EmailUpdateRequestLanguageAz     EmailUpdateRequestLanguage = "az"
	EmailUpdateRequestLanguageAzAz   EmailUpdateRequestLanguage = "az-az"
	EmailUpdateRequestLanguageBa     EmailUpdateRequestLanguage = "ba"
	EmailUpdateRequestLanguageBal    EmailUpdateRequestLanguage = "bal"
	EmailUpdateRequestLanguageBalPk  EmailUpdateRequestLanguage = "bal-pk"
	EmailUpdateRequestLanguageBas    EmailUpdateRequestLanguage = "bas"
	EmailUpdateRequestLanguageBasCm  EmailUpdateRequestLanguage = "bas-cm"
	EmailUpdateRequestLanguageBe     EmailUpdateRequestLanguage = "be"
	EmailUpdateRequestLanguageBeBy   EmailUpdateRequestLanguage = "be-by"
	EmailUpdateRequestLanguageBem    EmailUpdateRequestLanguage = "bem"
	EmailUpdateRequestLanguageBemZm  EmailUpdateRequestLanguage = "bem-zm"
	EmailUpdateRequestLanguageBez    EmailUpdateRequestLanguage = "bez"
	EmailUpdateRequestLanguageBezTz  EmailUpdateRequestLanguage = "bez-tz"
	EmailUpdateRequestLanguageBg     EmailUpdateRequestLanguage = "bg"
	EmailUpdateRequestLanguageBgBg   EmailUpdateRequestLanguage = "bg-bg"
	EmailUpdateRequestLanguageBgc    EmailUpdateRequestLanguage = "bgc"
	EmailUpdateRequestLanguageBgcIn  EmailUpdateRequestLanguage = "bgc-in"
	EmailUpdateRequestLanguageBho    EmailUpdateRequestLanguage = "bho"
	EmailUpdateRequestLanguageBhoIn  EmailUpdateRequestLanguage = "bho-in"
	EmailUpdateRequestLanguageBi     EmailUpdateRequestLanguage = "bi"
	EmailUpdateRequestLanguageBlo    EmailUpdateRequestLanguage = "blo"
	EmailUpdateRequestLanguageBloBj  EmailUpdateRequestLanguage = "blo-bj"
	EmailUpdateRequestLanguageBm     EmailUpdateRequestLanguage = "bm"
	EmailUpdateRequestLanguageBmMl   EmailUpdateRequestLanguage = "bm-ml"
	EmailUpdateRequestLanguageBn     EmailUpdateRequestLanguage = "bn"
	EmailUpdateRequestLanguageBnBd   EmailUpdateRequestLanguage = "bn-bd"
	EmailUpdateRequestLanguageBnIn   EmailUpdateRequestLanguage = "bn-in"
	EmailUpdateRequestLanguageBo     EmailUpdateRequestLanguage = "bo"
	EmailUpdateRequestLanguageBoCn   EmailUpdateRequestLanguage = "bo-cn"
	EmailUpdateRequestLanguageBoIn   EmailUpdateRequestLanguage = "bo-in"
	EmailUpdateRequestLanguageBr     EmailUpdateRequestLanguage = "br"
	EmailUpdateRequestLanguageBrFr   EmailUpdateRequestLanguage = "br-fr"
	EmailUpdateRequestLanguageBrx    EmailUpdateRequestLanguage = "brx"
	EmailUpdateRequestLanguageBrxIn  EmailUpdateRequestLanguage = "brx-in"
	EmailUpdateRequestLanguageBs     EmailUpdateRequestLanguage = "bs"
	EmailUpdateRequestLanguageBsBa   EmailUpdateRequestLanguage = "bs-ba"
	EmailUpdateRequestLanguageCa     EmailUpdateRequestLanguage = "ca"
	EmailUpdateRequestLanguageCaAd   EmailUpdateRequestLanguage = "ca-ad"
	EmailUpdateRequestLanguageCaEs   EmailUpdateRequestLanguage = "ca-es"
	EmailUpdateRequestLanguageCaFr   EmailUpdateRequestLanguage = "ca-fr"
	EmailUpdateRequestLanguageCaIt   EmailUpdateRequestLanguage = "ca-it"
	EmailUpdateRequestLanguageCcp    EmailUpdateRequestLanguage = "ccp"
	EmailUpdateRequestLanguageCcpBd  EmailUpdateRequestLanguage = "ccp-bd"
	EmailUpdateRequestLanguageCcpIn  EmailUpdateRequestLanguage = "ccp-in"
	EmailUpdateRequestLanguageCe     EmailUpdateRequestLanguage = "ce"
	EmailUpdateRequestLanguageCeRu   EmailUpdateRequestLanguage = "ce-ru"
	EmailUpdateRequestLanguageCeb    EmailUpdateRequestLanguage = "ceb"
	EmailUpdateRequestLanguageCebPh  EmailUpdateRequestLanguage = "ceb-ph"
	EmailUpdateRequestLanguageCgg    EmailUpdateRequestLanguage = "cgg"
	EmailUpdateRequestLanguageCggUg  EmailUpdateRequestLanguage = "cgg-ug"
	EmailUpdateRequestLanguageCh     EmailUpdateRequestLanguage = "ch"
	EmailUpdateRequestLanguageChr    EmailUpdateRequestLanguage = "chr"
	EmailUpdateRequestLanguageChrUs  EmailUpdateRequestLanguage = "chr-us"
	EmailUpdateRequestLanguageCkb    EmailUpdateRequestLanguage = "ckb"
	EmailUpdateRequestLanguageCkbIq  EmailUpdateRequestLanguage = "ckb-iq"
	EmailUpdateRequestLanguageCkbIr  EmailUpdateRequestLanguage = "ckb-ir"
	EmailUpdateRequestLanguageCo     EmailUpdateRequestLanguage = "co"
	EmailUpdateRequestLanguageCr     EmailUpdateRequestLanguage = "cr"
	EmailUpdateRequestLanguageCs     EmailUpdateRequestLanguage = "cs"
	EmailUpdateRequestLanguageCsCz   EmailUpdateRequestLanguage = "cs-cz"
	EmailUpdateRequestLanguageCsw    EmailUpdateRequestLanguage = "csw"
	EmailUpdateRequestLanguageCswCa  EmailUpdateRequestLanguage = "csw-ca"
	EmailUpdateRequestLanguageCu     EmailUpdateRequestLanguage = "cu"
	EmailUpdateRequestLanguageCuRu   EmailUpdateRequestLanguage = "cu-ru"
	EmailUpdateRequestLanguageCv     EmailUpdateRequestLanguage = "cv"
	EmailUpdateRequestLanguageCvRu   EmailUpdateRequestLanguage = "cv-ru"
	EmailUpdateRequestLanguageCy     EmailUpdateRequestLanguage = "cy"
	EmailUpdateRequestLanguageCyGB   EmailUpdateRequestLanguage = "cy-gb"
	EmailUpdateRequestLanguageDa     EmailUpdateRequestLanguage = "da"
	EmailUpdateRequestLanguageDaDk   EmailUpdateRequestLanguage = "da-dk"
	EmailUpdateRequestLanguageDaGl   EmailUpdateRequestLanguage = "da-gl"
	EmailUpdateRequestLanguageDav    EmailUpdateRequestLanguage = "dav"
	EmailUpdateRequestLanguageDavKe  EmailUpdateRequestLanguage = "dav-ke"
	EmailUpdateRequestLanguageDe     EmailUpdateRequestLanguage = "de"
	EmailUpdateRequestLanguageDeAt   EmailUpdateRequestLanguage = "de-at"
	EmailUpdateRequestLanguageDeBe   EmailUpdateRequestLanguage = "de-be"
	EmailUpdateRequestLanguageDeCh   EmailUpdateRequestLanguage = "de-ch"
	EmailUpdateRequestLanguageDeDe   EmailUpdateRequestLanguage = "de-de"
	EmailUpdateRequestLanguageDeGr   EmailUpdateRequestLanguage = "de-gr"
	EmailUpdateRequestLanguageDeIt   EmailUpdateRequestLanguage = "de-it"
	EmailUpdateRequestLanguageDeLi   EmailUpdateRequestLanguage = "de-li"
	EmailUpdateRequestLanguageDeLu   EmailUpdateRequestLanguage = "de-lu"
	EmailUpdateRequestLanguageDje    EmailUpdateRequestLanguage = "dje"
	EmailUpdateRequestLanguageDjeNe  EmailUpdateRequestLanguage = "dje-ne"
	EmailUpdateRequestLanguageDoi    EmailUpdateRequestLanguage = "doi"
	EmailUpdateRequestLanguageDoiIn  EmailUpdateRequestLanguage = "doi-in"
	EmailUpdateRequestLanguageDsb    EmailUpdateRequestLanguage = "dsb"
	EmailUpdateRequestLanguageDsbDe  EmailUpdateRequestLanguage = "dsb-de"
	EmailUpdateRequestLanguageDua    EmailUpdateRequestLanguage = "dua"
	EmailUpdateRequestLanguageDuaCm  EmailUpdateRequestLanguage = "dua-cm"
	EmailUpdateRequestLanguageDv     EmailUpdateRequestLanguage = "dv"
	EmailUpdateRequestLanguageDyo    EmailUpdateRequestLanguage = "dyo"
	EmailUpdateRequestLanguageDyoSn  EmailUpdateRequestLanguage = "dyo-sn"
	EmailUpdateRequestLanguageDz     EmailUpdateRequestLanguage = "dz"
	EmailUpdateRequestLanguageDzBt   EmailUpdateRequestLanguage = "dz-bt"
	EmailUpdateRequestLanguageEbu    EmailUpdateRequestLanguage = "ebu"
	EmailUpdateRequestLanguageEbuKe  EmailUpdateRequestLanguage = "ebu-ke"
	EmailUpdateRequestLanguageEe     EmailUpdateRequestLanguage = "ee"
	EmailUpdateRequestLanguageEeGh   EmailUpdateRequestLanguage = "ee-gh"
	EmailUpdateRequestLanguageEeTg   EmailUpdateRequestLanguage = "ee-tg"
	EmailUpdateRequestLanguageEl     EmailUpdateRequestLanguage = "el"
	EmailUpdateRequestLanguageElCy   EmailUpdateRequestLanguage = "el-cy"
	EmailUpdateRequestLanguageElGr   EmailUpdateRequestLanguage = "el-gr"
	EmailUpdateRequestLanguageEn     EmailUpdateRequestLanguage = "en"
	EmailUpdateRequestLanguageEn001  EmailUpdateRequestLanguage = "en-001"
	EmailUpdateRequestLanguageEn150  EmailUpdateRequestLanguage = "en-150"
	EmailUpdateRequestLanguageEnAe   EmailUpdateRequestLanguage = "en-ae"
	EmailUpdateRequestLanguageEnAg   EmailUpdateRequestLanguage = "en-ag"
	EmailUpdateRequestLanguageEnAI   EmailUpdateRequestLanguage = "en-ai"
	EmailUpdateRequestLanguageEnAs   EmailUpdateRequestLanguage = "en-as"
	EmailUpdateRequestLanguageEnAt   EmailUpdateRequestLanguage = "en-at"
	EmailUpdateRequestLanguageEnAu   EmailUpdateRequestLanguage = "en-au"
	EmailUpdateRequestLanguageEnBb   EmailUpdateRequestLanguage = "en-bb"
	EmailUpdateRequestLanguageEnBe   EmailUpdateRequestLanguage = "en-be"
	EmailUpdateRequestLanguageEnBi   EmailUpdateRequestLanguage = "en-bi"
	EmailUpdateRequestLanguageEnBm   EmailUpdateRequestLanguage = "en-bm"
	EmailUpdateRequestLanguageEnBs   EmailUpdateRequestLanguage = "en-bs"
	EmailUpdateRequestLanguageEnBw   EmailUpdateRequestLanguage = "en-bw"
	EmailUpdateRequestLanguageEnBz   EmailUpdateRequestLanguage = "en-bz"
	EmailUpdateRequestLanguageEnCa   EmailUpdateRequestLanguage = "en-ca"
	EmailUpdateRequestLanguageEnCc   EmailUpdateRequestLanguage = "en-cc"
	EmailUpdateRequestLanguageEnCh   EmailUpdateRequestLanguage = "en-ch"
	EmailUpdateRequestLanguageEnCk   EmailUpdateRequestLanguage = "en-ck"
	EmailUpdateRequestLanguageEnCm   EmailUpdateRequestLanguage = "en-cm"
	EmailUpdateRequestLanguageEnCn   EmailUpdateRequestLanguage = "en-cn"
	EmailUpdateRequestLanguageEnCx   EmailUpdateRequestLanguage = "en-cx"
	EmailUpdateRequestLanguageEnCy   EmailUpdateRequestLanguage = "en-cy"
	EmailUpdateRequestLanguageEnCz   EmailUpdateRequestLanguage = "en-cz"
	EmailUpdateRequestLanguageEnDe   EmailUpdateRequestLanguage = "en-de"
	EmailUpdateRequestLanguageEnDg   EmailUpdateRequestLanguage = "en-dg"
	EmailUpdateRequestLanguageEnDk   EmailUpdateRequestLanguage = "en-dk"
	EmailUpdateRequestLanguageEnDm   EmailUpdateRequestLanguage = "en-dm"
	EmailUpdateRequestLanguageEnEe   EmailUpdateRequestLanguage = "en-ee"
	EmailUpdateRequestLanguageEnEg   EmailUpdateRequestLanguage = "en-eg"
	EmailUpdateRequestLanguageEnEr   EmailUpdateRequestLanguage = "en-er"
	EmailUpdateRequestLanguageEnEs   EmailUpdateRequestLanguage = "en-es"
	EmailUpdateRequestLanguageEnFi   EmailUpdateRequestLanguage = "en-fi"
	EmailUpdateRequestLanguageEnFj   EmailUpdateRequestLanguage = "en-fj"
	EmailUpdateRequestLanguageEnFk   EmailUpdateRequestLanguage = "en-fk"
	EmailUpdateRequestLanguageEnFm   EmailUpdateRequestLanguage = "en-fm"
	EmailUpdateRequestLanguageEnFr   EmailUpdateRequestLanguage = "en-fr"
	EmailUpdateRequestLanguageEnGB   EmailUpdateRequestLanguage = "en-gb"
	EmailUpdateRequestLanguageEnGd   EmailUpdateRequestLanguage = "en-gd"
	EmailUpdateRequestLanguageEnGg   EmailUpdateRequestLanguage = "en-gg"
	EmailUpdateRequestLanguageEnGh   EmailUpdateRequestLanguage = "en-gh"
	EmailUpdateRequestLanguageEnGi   EmailUpdateRequestLanguage = "en-gi"
	EmailUpdateRequestLanguageEnGm   EmailUpdateRequestLanguage = "en-gm"
	EmailUpdateRequestLanguageEnGs   EmailUpdateRequestLanguage = "en-gs"
	EmailUpdateRequestLanguageEnGu   EmailUpdateRequestLanguage = "en-gu"
	EmailUpdateRequestLanguageEnGy   EmailUpdateRequestLanguage = "en-gy"
	EmailUpdateRequestLanguageEnHk   EmailUpdateRequestLanguage = "en-hk"
	EmailUpdateRequestLanguageEnHu   EmailUpdateRequestLanguage = "en-hu"
	EmailUpdateRequestLanguageEnID   EmailUpdateRequestLanguage = "en-id"
	EmailUpdateRequestLanguageEnIe   EmailUpdateRequestLanguage = "en-ie"
	EmailUpdateRequestLanguageEnIl   EmailUpdateRequestLanguage = "en-il"
	EmailUpdateRequestLanguageEnIm   EmailUpdateRequestLanguage = "en-im"
	EmailUpdateRequestLanguageEnIn   EmailUpdateRequestLanguage = "en-in"
	EmailUpdateRequestLanguageEnIo   EmailUpdateRequestLanguage = "en-io"
	EmailUpdateRequestLanguageEnIt   EmailUpdateRequestLanguage = "en-it"
	EmailUpdateRequestLanguageEnJe   EmailUpdateRequestLanguage = "en-je"
	EmailUpdateRequestLanguageEnJm   EmailUpdateRequestLanguage = "en-jm"
	EmailUpdateRequestLanguageEnKe   EmailUpdateRequestLanguage = "en-ke"
	EmailUpdateRequestLanguageEnKi   EmailUpdateRequestLanguage = "en-ki"
	EmailUpdateRequestLanguageEnKn   EmailUpdateRequestLanguage = "en-kn"
	EmailUpdateRequestLanguageEnKy   EmailUpdateRequestLanguage = "en-ky"
	EmailUpdateRequestLanguageEnLc   EmailUpdateRequestLanguage = "en-lc"
	EmailUpdateRequestLanguageEnLr   EmailUpdateRequestLanguage = "en-lr"
	EmailUpdateRequestLanguageEnLs   EmailUpdateRequestLanguage = "en-ls"
	EmailUpdateRequestLanguageEnLu   EmailUpdateRequestLanguage = "en-lu"
	EmailUpdateRequestLanguageEnMg   EmailUpdateRequestLanguage = "en-mg"
	EmailUpdateRequestLanguageEnMh   EmailUpdateRequestLanguage = "en-mh"
	EmailUpdateRequestLanguageEnMo   EmailUpdateRequestLanguage = "en-mo"
	EmailUpdateRequestLanguageEnMp   EmailUpdateRequestLanguage = "en-mp"
	EmailUpdateRequestLanguageEnMs   EmailUpdateRequestLanguage = "en-ms"
	EmailUpdateRequestLanguageEnMt   EmailUpdateRequestLanguage = "en-mt"
	EmailUpdateRequestLanguageEnMu   EmailUpdateRequestLanguage = "en-mu"
	EmailUpdateRequestLanguageEnMv   EmailUpdateRequestLanguage = "en-mv"
	EmailUpdateRequestLanguageEnMw   EmailUpdateRequestLanguage = "en-mw"
	EmailUpdateRequestLanguageEnMx   EmailUpdateRequestLanguage = "en-mx"
	EmailUpdateRequestLanguageEnMy   EmailUpdateRequestLanguage = "en-my"
	EmailUpdateRequestLanguageEnNa   EmailUpdateRequestLanguage = "en-na"
	EmailUpdateRequestLanguageEnNf   EmailUpdateRequestLanguage = "en-nf"
	EmailUpdateRequestLanguageEnNg   EmailUpdateRequestLanguage = "en-ng"
	EmailUpdateRequestLanguageEnNl   EmailUpdateRequestLanguage = "en-nl"
	EmailUpdateRequestLanguageEnNo   EmailUpdateRequestLanguage = "en-no"
	EmailUpdateRequestLanguageEnNr   EmailUpdateRequestLanguage = "en-nr"
	EmailUpdateRequestLanguageEnNu   EmailUpdateRequestLanguage = "en-nu"
	EmailUpdateRequestLanguageEnNz   EmailUpdateRequestLanguage = "en-nz"
	EmailUpdateRequestLanguageEnPg   EmailUpdateRequestLanguage = "en-pg"
	EmailUpdateRequestLanguageEnPh   EmailUpdateRequestLanguage = "en-ph"
	EmailUpdateRequestLanguageEnPk   EmailUpdateRequestLanguage = "en-pk"
	EmailUpdateRequestLanguageEnPl   EmailUpdateRequestLanguage = "en-pl"
	EmailUpdateRequestLanguageEnPn   EmailUpdateRequestLanguage = "en-pn"
	EmailUpdateRequestLanguageEnPr   EmailUpdateRequestLanguage = "en-pr"
	EmailUpdateRequestLanguageEnPt   EmailUpdateRequestLanguage = "en-pt"
	EmailUpdateRequestLanguageEnPw   EmailUpdateRequestLanguage = "en-pw"
	EmailUpdateRequestLanguageEnRo   EmailUpdateRequestLanguage = "en-ro"
	EmailUpdateRequestLanguageEnRw   EmailUpdateRequestLanguage = "en-rw"
	EmailUpdateRequestLanguageEnSb   EmailUpdateRequestLanguage = "en-sb"
	EmailUpdateRequestLanguageEnSc   EmailUpdateRequestLanguage = "en-sc"
	EmailUpdateRequestLanguageEnSd   EmailUpdateRequestLanguage = "en-sd"
	EmailUpdateRequestLanguageEnSe   EmailUpdateRequestLanguage = "en-se"
	EmailUpdateRequestLanguageEnSg   EmailUpdateRequestLanguage = "en-sg"
	EmailUpdateRequestLanguageEnSh   EmailUpdateRequestLanguage = "en-sh"
	EmailUpdateRequestLanguageEnSi   EmailUpdateRequestLanguage = "en-si"
	EmailUpdateRequestLanguageEnSk   EmailUpdateRequestLanguage = "en-sk"
	EmailUpdateRequestLanguageEnSl   EmailUpdateRequestLanguage = "en-sl"
	EmailUpdateRequestLanguageEnSS   EmailUpdateRequestLanguage = "en-ss"
	EmailUpdateRequestLanguageEnSx   EmailUpdateRequestLanguage = "en-sx"
	EmailUpdateRequestLanguageEnSz   EmailUpdateRequestLanguage = "en-sz"
	EmailUpdateRequestLanguageEnTc   EmailUpdateRequestLanguage = "en-tc"
	EmailUpdateRequestLanguageEnTh   EmailUpdateRequestLanguage = "en-th"
	EmailUpdateRequestLanguageEnTk   EmailUpdateRequestLanguage = "en-tk"
	EmailUpdateRequestLanguageEnTn   EmailUpdateRequestLanguage = "en-tn"
	EmailUpdateRequestLanguageEnTo   EmailUpdateRequestLanguage = "en-to"
	EmailUpdateRequestLanguageEnTt   EmailUpdateRequestLanguage = "en-tt"
	EmailUpdateRequestLanguageEnTv   EmailUpdateRequestLanguage = "en-tv"
	EmailUpdateRequestLanguageEnTz   EmailUpdateRequestLanguage = "en-tz"
	EmailUpdateRequestLanguageEnUg   EmailUpdateRequestLanguage = "en-ug"
	EmailUpdateRequestLanguageEnUm   EmailUpdateRequestLanguage = "en-um"
	EmailUpdateRequestLanguageEnUs   EmailUpdateRequestLanguage = "en-us"
	EmailUpdateRequestLanguageEnVc   EmailUpdateRequestLanguage = "en-vc"
	EmailUpdateRequestLanguageEnVg   EmailUpdateRequestLanguage = "en-vg"
	EmailUpdateRequestLanguageEnVi   EmailUpdateRequestLanguage = "en-vi"
	EmailUpdateRequestLanguageEnVn   EmailUpdateRequestLanguage = "en-vn"
	EmailUpdateRequestLanguageEnVu   EmailUpdateRequestLanguage = "en-vu"
	EmailUpdateRequestLanguageEnWs   EmailUpdateRequestLanguage = "en-ws"
	EmailUpdateRequestLanguageEnZa   EmailUpdateRequestLanguage = "en-za"
	EmailUpdateRequestLanguageEnZm   EmailUpdateRequestLanguage = "en-zm"
	EmailUpdateRequestLanguageEnZw   EmailUpdateRequestLanguage = "en-zw"
	EmailUpdateRequestLanguageEo     EmailUpdateRequestLanguage = "eo"
	EmailUpdateRequestLanguageEo001  EmailUpdateRequestLanguage = "eo-001"
	EmailUpdateRequestLanguageEs     EmailUpdateRequestLanguage = "es"
	EmailUpdateRequestLanguageEs419  EmailUpdateRequestLanguage = "es-419"
	EmailUpdateRequestLanguageEsAr   EmailUpdateRequestLanguage = "es-ar"
	EmailUpdateRequestLanguageEsBo   EmailUpdateRequestLanguage = "es-bo"
	EmailUpdateRequestLanguageEsBr   EmailUpdateRequestLanguage = "es-br"
	EmailUpdateRequestLanguageEsBz   EmailUpdateRequestLanguage = "es-bz"
	EmailUpdateRequestLanguageEsCl   EmailUpdateRequestLanguage = "es-cl"
	EmailUpdateRequestLanguageEsCo   EmailUpdateRequestLanguage = "es-co"
	EmailUpdateRequestLanguageEsCr   EmailUpdateRequestLanguage = "es-cr"
	EmailUpdateRequestLanguageEsCu   EmailUpdateRequestLanguage = "es-cu"
	EmailUpdateRequestLanguageEsDo   EmailUpdateRequestLanguage = "es-do"
	EmailUpdateRequestLanguageEsEa   EmailUpdateRequestLanguage = "es-ea"
	EmailUpdateRequestLanguageEsEc   EmailUpdateRequestLanguage = "es-ec"
	EmailUpdateRequestLanguageEsEs   EmailUpdateRequestLanguage = "es-es"
	EmailUpdateRequestLanguageEsGq   EmailUpdateRequestLanguage = "es-gq"
	EmailUpdateRequestLanguageEsGt   EmailUpdateRequestLanguage = "es-gt"
	EmailUpdateRequestLanguageEsHn   EmailUpdateRequestLanguage = "es-hn"
	EmailUpdateRequestLanguageEsIc   EmailUpdateRequestLanguage = "es-ic"
	EmailUpdateRequestLanguageEsMx   EmailUpdateRequestLanguage = "es-mx"
	EmailUpdateRequestLanguageEsNi   EmailUpdateRequestLanguage = "es-ni"
	EmailUpdateRequestLanguageEsPa   EmailUpdateRequestLanguage = "es-pa"
	EmailUpdateRequestLanguageEsPe   EmailUpdateRequestLanguage = "es-pe"
	EmailUpdateRequestLanguageEsPh   EmailUpdateRequestLanguage = "es-ph"
	EmailUpdateRequestLanguageEsPr   EmailUpdateRequestLanguage = "es-pr"
	EmailUpdateRequestLanguageEsPy   EmailUpdateRequestLanguage = "es-py"
	EmailUpdateRequestLanguageEsSv   EmailUpdateRequestLanguage = "es-sv"
	EmailUpdateRequestLanguageEsUs   EmailUpdateRequestLanguage = "es-us"
	EmailUpdateRequestLanguageEsUy   EmailUpdateRequestLanguage = "es-uy"
	EmailUpdateRequestLanguageEsVe   EmailUpdateRequestLanguage = "es-ve"
	EmailUpdateRequestLanguageEt     EmailUpdateRequestLanguage = "et"
	EmailUpdateRequestLanguageEtEe   EmailUpdateRequestLanguage = "et-ee"
	EmailUpdateRequestLanguageEu     EmailUpdateRequestLanguage = "eu"
	EmailUpdateRequestLanguageEuEs   EmailUpdateRequestLanguage = "eu-es"
	EmailUpdateRequestLanguageEwo    EmailUpdateRequestLanguage = "ewo"
	EmailUpdateRequestLanguageEwoCm  EmailUpdateRequestLanguage = "ewo-cm"
	EmailUpdateRequestLanguageFa     EmailUpdateRequestLanguage = "fa"
	EmailUpdateRequestLanguageFaAf   EmailUpdateRequestLanguage = "fa-af"
	EmailUpdateRequestLanguageFaIr   EmailUpdateRequestLanguage = "fa-ir"
	EmailUpdateRequestLanguageFf     EmailUpdateRequestLanguage = "ff"
	EmailUpdateRequestLanguageFfBf   EmailUpdateRequestLanguage = "ff-bf"
	EmailUpdateRequestLanguageFfCm   EmailUpdateRequestLanguage = "ff-cm"
	EmailUpdateRequestLanguageFfGh   EmailUpdateRequestLanguage = "ff-gh"
	EmailUpdateRequestLanguageFfGm   EmailUpdateRequestLanguage = "ff-gm"
	EmailUpdateRequestLanguageFfGn   EmailUpdateRequestLanguage = "ff-gn"
	EmailUpdateRequestLanguageFfGw   EmailUpdateRequestLanguage = "ff-gw"
	EmailUpdateRequestLanguageFfLr   EmailUpdateRequestLanguage = "ff-lr"
	EmailUpdateRequestLanguageFfMr   EmailUpdateRequestLanguage = "ff-mr"
	EmailUpdateRequestLanguageFfNe   EmailUpdateRequestLanguage = "ff-ne"
	EmailUpdateRequestLanguageFfNg   EmailUpdateRequestLanguage = "ff-ng"
	EmailUpdateRequestLanguageFfSl   EmailUpdateRequestLanguage = "ff-sl"
	EmailUpdateRequestLanguageFfSn   EmailUpdateRequestLanguage = "ff-sn"
	EmailUpdateRequestLanguageFi     EmailUpdateRequestLanguage = "fi"
	EmailUpdateRequestLanguageFiFi   EmailUpdateRequestLanguage = "fi-fi"
	EmailUpdateRequestLanguageFil    EmailUpdateRequestLanguage = "fil"
	EmailUpdateRequestLanguageFilPh  EmailUpdateRequestLanguage = "fil-ph"
	EmailUpdateRequestLanguageFj     EmailUpdateRequestLanguage = "fj"
	EmailUpdateRequestLanguageFo     EmailUpdateRequestLanguage = "fo"
	EmailUpdateRequestLanguageFoDk   EmailUpdateRequestLanguage = "fo-dk"
	EmailUpdateRequestLanguageFoFo   EmailUpdateRequestLanguage = "fo-fo"
	EmailUpdateRequestLanguageFr     EmailUpdateRequestLanguage = "fr"
	EmailUpdateRequestLanguageFrBe   EmailUpdateRequestLanguage = "fr-be"
	EmailUpdateRequestLanguageFrBf   EmailUpdateRequestLanguage = "fr-bf"
	EmailUpdateRequestLanguageFrBi   EmailUpdateRequestLanguage = "fr-bi"
	EmailUpdateRequestLanguageFrBj   EmailUpdateRequestLanguage = "fr-bj"
	EmailUpdateRequestLanguageFrBl   EmailUpdateRequestLanguage = "fr-bl"
	EmailUpdateRequestLanguageFrCa   EmailUpdateRequestLanguage = "fr-ca"
	EmailUpdateRequestLanguageFrCd   EmailUpdateRequestLanguage = "fr-cd"
	EmailUpdateRequestLanguageFrCf   EmailUpdateRequestLanguage = "fr-cf"
	EmailUpdateRequestLanguageFrCg   EmailUpdateRequestLanguage = "fr-cg"
	EmailUpdateRequestLanguageFrCh   EmailUpdateRequestLanguage = "fr-ch"
	EmailUpdateRequestLanguageFrCi   EmailUpdateRequestLanguage = "fr-ci"
	EmailUpdateRequestLanguageFrCm   EmailUpdateRequestLanguage = "fr-cm"
	EmailUpdateRequestLanguageFrDj   EmailUpdateRequestLanguage = "fr-dj"
	EmailUpdateRequestLanguageFrDz   EmailUpdateRequestLanguage = "fr-dz"
	EmailUpdateRequestLanguageFrFr   EmailUpdateRequestLanguage = "fr-fr"
	EmailUpdateRequestLanguageFrGa   EmailUpdateRequestLanguage = "fr-ga"
	EmailUpdateRequestLanguageFrGf   EmailUpdateRequestLanguage = "fr-gf"
	EmailUpdateRequestLanguageFrGn   EmailUpdateRequestLanguage = "fr-gn"
	EmailUpdateRequestLanguageFrGp   EmailUpdateRequestLanguage = "fr-gp"
	EmailUpdateRequestLanguageFrGq   EmailUpdateRequestLanguage = "fr-gq"
	EmailUpdateRequestLanguageFrHt   EmailUpdateRequestLanguage = "fr-ht"
	EmailUpdateRequestLanguageFrKm   EmailUpdateRequestLanguage = "fr-km"
	EmailUpdateRequestLanguageFrLu   EmailUpdateRequestLanguage = "fr-lu"
	EmailUpdateRequestLanguageFrMa   EmailUpdateRequestLanguage = "fr-ma"
	EmailUpdateRequestLanguageFrMc   EmailUpdateRequestLanguage = "fr-mc"
	EmailUpdateRequestLanguageFrMf   EmailUpdateRequestLanguage = "fr-mf"
	EmailUpdateRequestLanguageFrMg   EmailUpdateRequestLanguage = "fr-mg"
	EmailUpdateRequestLanguageFrMl   EmailUpdateRequestLanguage = "fr-ml"
	EmailUpdateRequestLanguageFrMq   EmailUpdateRequestLanguage = "fr-mq"
	EmailUpdateRequestLanguageFrMr   EmailUpdateRequestLanguage = "fr-mr"
	EmailUpdateRequestLanguageFrMu   EmailUpdateRequestLanguage = "fr-mu"
	EmailUpdateRequestLanguageFrNc   EmailUpdateRequestLanguage = "fr-nc"
	EmailUpdateRequestLanguageFrNe   EmailUpdateRequestLanguage = "fr-ne"
	EmailUpdateRequestLanguageFrPf   EmailUpdateRequestLanguage = "fr-pf"
	EmailUpdateRequestLanguageFrPm   EmailUpdateRequestLanguage = "fr-pm"
	EmailUpdateRequestLanguageFrRe   EmailUpdateRequestLanguage = "fr-re"
	EmailUpdateRequestLanguageFrRw   EmailUpdateRequestLanguage = "fr-rw"
	EmailUpdateRequestLanguageFrSc   EmailUpdateRequestLanguage = "fr-sc"
	EmailUpdateRequestLanguageFrSn   EmailUpdateRequestLanguage = "fr-sn"
	EmailUpdateRequestLanguageFrSy   EmailUpdateRequestLanguage = "fr-sy"
	EmailUpdateRequestLanguageFrTd   EmailUpdateRequestLanguage = "fr-td"
	EmailUpdateRequestLanguageFrTg   EmailUpdateRequestLanguage = "fr-tg"
	EmailUpdateRequestLanguageFrTn   EmailUpdateRequestLanguage = "fr-tn"
	EmailUpdateRequestLanguageFrVu   EmailUpdateRequestLanguage = "fr-vu"
	EmailUpdateRequestLanguageFrWf   EmailUpdateRequestLanguage = "fr-wf"
	EmailUpdateRequestLanguageFrYt   EmailUpdateRequestLanguage = "fr-yt"
	EmailUpdateRequestLanguageFrr    EmailUpdateRequestLanguage = "frr"
	EmailUpdateRequestLanguageFrrDe  EmailUpdateRequestLanguage = "frr-de"
	EmailUpdateRequestLanguageFur    EmailUpdateRequestLanguage = "fur"
	EmailUpdateRequestLanguageFurIt  EmailUpdateRequestLanguage = "fur-it"
	EmailUpdateRequestLanguageFy     EmailUpdateRequestLanguage = "fy"
	EmailUpdateRequestLanguageFyNl   EmailUpdateRequestLanguage = "fy-nl"
	EmailUpdateRequestLanguageGa     EmailUpdateRequestLanguage = "ga"
	EmailUpdateRequestLanguageGaGB   EmailUpdateRequestLanguage = "ga-gb"
	EmailUpdateRequestLanguageGaIe   EmailUpdateRequestLanguage = "ga-ie"
	EmailUpdateRequestLanguageGaa    EmailUpdateRequestLanguage = "gaa"
	EmailUpdateRequestLanguageGaaGh  EmailUpdateRequestLanguage = "gaa-gh"
	EmailUpdateRequestLanguageGd     EmailUpdateRequestLanguage = "gd"
	EmailUpdateRequestLanguageGdGB   EmailUpdateRequestLanguage = "gd-gb"
	EmailUpdateRequestLanguageGl     EmailUpdateRequestLanguage = "gl"
	EmailUpdateRequestLanguageGlEs   EmailUpdateRequestLanguage = "gl-es"
	EmailUpdateRequestLanguageGn     EmailUpdateRequestLanguage = "gn"
	EmailUpdateRequestLanguageGsw    EmailUpdateRequestLanguage = "gsw"
	EmailUpdateRequestLanguageGswCh  EmailUpdateRequestLanguage = "gsw-ch"
	EmailUpdateRequestLanguageGswFr  EmailUpdateRequestLanguage = "gsw-fr"
	EmailUpdateRequestLanguageGswLi  EmailUpdateRequestLanguage = "gsw-li"
	EmailUpdateRequestLanguageGu     EmailUpdateRequestLanguage = "gu"
	EmailUpdateRequestLanguageGuIn   EmailUpdateRequestLanguage = "gu-in"
	EmailUpdateRequestLanguageGuz    EmailUpdateRequestLanguage = "guz"
	EmailUpdateRequestLanguageGuzKe  EmailUpdateRequestLanguage = "guz-ke"
	EmailUpdateRequestLanguageGv     EmailUpdateRequestLanguage = "gv"
	EmailUpdateRequestLanguageGvIm   EmailUpdateRequestLanguage = "gv-im"
	EmailUpdateRequestLanguageHa     EmailUpdateRequestLanguage = "ha"
	EmailUpdateRequestLanguageHaGh   EmailUpdateRequestLanguage = "ha-gh"
	EmailUpdateRequestLanguageHaNe   EmailUpdateRequestLanguage = "ha-ne"
	EmailUpdateRequestLanguageHaNg   EmailUpdateRequestLanguage = "ha-ng"
	EmailUpdateRequestLanguageHaw    EmailUpdateRequestLanguage = "haw"
	EmailUpdateRequestLanguageHawUs  EmailUpdateRequestLanguage = "haw-us"
	EmailUpdateRequestLanguageHe     EmailUpdateRequestLanguage = "he"
	EmailUpdateRequestLanguageHeIl   EmailUpdateRequestLanguage = "he-il"
	EmailUpdateRequestLanguageHi     EmailUpdateRequestLanguage = "hi"
	EmailUpdateRequestLanguageHiIn   EmailUpdateRequestLanguage = "hi-in"
	EmailUpdateRequestLanguageHmn    EmailUpdateRequestLanguage = "hmn"
	EmailUpdateRequestLanguageHo     EmailUpdateRequestLanguage = "ho"
	EmailUpdateRequestLanguageHr     EmailUpdateRequestLanguage = "hr"
	EmailUpdateRequestLanguageHrBa   EmailUpdateRequestLanguage = "hr-ba"
	EmailUpdateRequestLanguageHrHr   EmailUpdateRequestLanguage = "hr-hr"
	EmailUpdateRequestLanguageHsb    EmailUpdateRequestLanguage = "hsb"
	EmailUpdateRequestLanguageHsbDe  EmailUpdateRequestLanguage = "hsb-de"
	EmailUpdateRequestLanguageHt     EmailUpdateRequestLanguage = "ht"
	EmailUpdateRequestLanguageHtHt   EmailUpdateRequestLanguage = "ht-ht"
	EmailUpdateRequestLanguageHu     EmailUpdateRequestLanguage = "hu"
	EmailUpdateRequestLanguageHuHu   EmailUpdateRequestLanguage = "hu-hu"
	EmailUpdateRequestLanguageHy     EmailUpdateRequestLanguage = "hy"
	EmailUpdateRequestLanguageHyAm   EmailUpdateRequestLanguage = "hy-am"
	EmailUpdateRequestLanguageHz     EmailUpdateRequestLanguage = "hz"
	EmailUpdateRequestLanguageIa     EmailUpdateRequestLanguage = "ia"
	EmailUpdateRequestLanguageIa001  EmailUpdateRequestLanguage = "ia-001"
	EmailUpdateRequestLanguageID     EmailUpdateRequestLanguage = "id"
	EmailUpdateRequestLanguageIDID   EmailUpdateRequestLanguage = "id-id"
	EmailUpdateRequestLanguageIe     EmailUpdateRequestLanguage = "ie"
	EmailUpdateRequestLanguageIeEe   EmailUpdateRequestLanguage = "ie-ee"
	EmailUpdateRequestLanguageIg     EmailUpdateRequestLanguage = "ig"
	EmailUpdateRequestLanguageIgNg   EmailUpdateRequestLanguage = "ig-ng"
	EmailUpdateRequestLanguageIi     EmailUpdateRequestLanguage = "ii"
	EmailUpdateRequestLanguageIiCn   EmailUpdateRequestLanguage = "ii-cn"
	EmailUpdateRequestLanguageIk     EmailUpdateRequestLanguage = "ik"
	EmailUpdateRequestLanguageIo     EmailUpdateRequestLanguage = "io"
	EmailUpdateRequestLanguageIs     EmailUpdateRequestLanguage = "is"
	EmailUpdateRequestLanguageIsIs   EmailUpdateRequestLanguage = "is-is"
	EmailUpdateRequestLanguageIt     EmailUpdateRequestLanguage = "it"
	EmailUpdateRequestLanguageItCh   EmailUpdateRequestLanguage = "it-ch"
	EmailUpdateRequestLanguageItIt   EmailUpdateRequestLanguage = "it-it"
	EmailUpdateRequestLanguageItSm   EmailUpdateRequestLanguage = "it-sm"
	EmailUpdateRequestLanguageItVa   EmailUpdateRequestLanguage = "it-va"
	EmailUpdateRequestLanguageIu     EmailUpdateRequestLanguage = "iu"
	EmailUpdateRequestLanguageJa     EmailUpdateRequestLanguage = "ja"
	EmailUpdateRequestLanguageJaJp   EmailUpdateRequestLanguage = "ja-jp"
	EmailUpdateRequestLanguageJgo    EmailUpdateRequestLanguage = "jgo"
	EmailUpdateRequestLanguageJgoCm  EmailUpdateRequestLanguage = "jgo-cm"
	EmailUpdateRequestLanguageJmc    EmailUpdateRequestLanguage = "jmc"
	EmailUpdateRequestLanguageJmcTz  EmailUpdateRequestLanguage = "jmc-tz"
	EmailUpdateRequestLanguageJv     EmailUpdateRequestLanguage = "jv"
	EmailUpdateRequestLanguageJvID   EmailUpdateRequestLanguage = "jv-id"
	EmailUpdateRequestLanguageKa     EmailUpdateRequestLanguage = "ka"
	EmailUpdateRequestLanguageKaGe   EmailUpdateRequestLanguage = "ka-ge"
	EmailUpdateRequestLanguageKab    EmailUpdateRequestLanguage = "kab"
	EmailUpdateRequestLanguageKabDz  EmailUpdateRequestLanguage = "kab-dz"
	EmailUpdateRequestLanguageKam    EmailUpdateRequestLanguage = "kam"
	EmailUpdateRequestLanguageKamKe  EmailUpdateRequestLanguage = "kam-ke"
	EmailUpdateRequestLanguageKar    EmailUpdateRequestLanguage = "kar"
	EmailUpdateRequestLanguageKde    EmailUpdateRequestLanguage = "kde"
	EmailUpdateRequestLanguageKdeTz  EmailUpdateRequestLanguage = "kde-tz"
	EmailUpdateRequestLanguageKea    EmailUpdateRequestLanguage = "kea"
	EmailUpdateRequestLanguageKeaCv  EmailUpdateRequestLanguage = "kea-cv"
	EmailUpdateRequestLanguageKg     EmailUpdateRequestLanguage = "kg"
	EmailUpdateRequestLanguageKgp    EmailUpdateRequestLanguage = "kgp"
	EmailUpdateRequestLanguageKgpBr  EmailUpdateRequestLanguage = "kgp-br"
	EmailUpdateRequestLanguageKh     EmailUpdateRequestLanguage = "kh"
	EmailUpdateRequestLanguageKhq    EmailUpdateRequestLanguage = "khq"
	EmailUpdateRequestLanguageKhqMl  EmailUpdateRequestLanguage = "khq-ml"
	EmailUpdateRequestLanguageKi     EmailUpdateRequestLanguage = "ki"
	EmailUpdateRequestLanguageKiKe   EmailUpdateRequestLanguage = "ki-ke"
	EmailUpdateRequestLanguageKj     EmailUpdateRequestLanguage = "kj"
	EmailUpdateRequestLanguageKk     EmailUpdateRequestLanguage = "kk"
	EmailUpdateRequestLanguageKkKz   EmailUpdateRequestLanguage = "kk-kz"
	EmailUpdateRequestLanguageKkj    EmailUpdateRequestLanguage = "kkj"
	EmailUpdateRequestLanguageKkjCm  EmailUpdateRequestLanguage = "kkj-cm"
	EmailUpdateRequestLanguageKl     EmailUpdateRequestLanguage = "kl"
	EmailUpdateRequestLanguageKlGl   EmailUpdateRequestLanguage = "kl-gl"
	EmailUpdateRequestLanguageKln    EmailUpdateRequestLanguage = "kln"
	EmailUpdateRequestLanguageKlnKe  EmailUpdateRequestLanguage = "kln-ke"
	EmailUpdateRequestLanguageKm     EmailUpdateRequestLanguage = "km"
	EmailUpdateRequestLanguageKmKh   EmailUpdateRequestLanguage = "km-kh"
	EmailUpdateRequestLanguageKn     EmailUpdateRequestLanguage = "kn"
	EmailUpdateRequestLanguageKnIn   EmailUpdateRequestLanguage = "kn-in"
	EmailUpdateRequestLanguageKo     EmailUpdateRequestLanguage = "ko"
	EmailUpdateRequestLanguageKoCn   EmailUpdateRequestLanguage = "ko-cn"
	EmailUpdateRequestLanguageKoKp   EmailUpdateRequestLanguage = "ko-kp"
	EmailUpdateRequestLanguageKoKr   EmailUpdateRequestLanguage = "ko-kr"
	EmailUpdateRequestLanguageKok    EmailUpdateRequestLanguage = "kok"
	EmailUpdateRequestLanguageKokIn  EmailUpdateRequestLanguage = "kok-in"
	EmailUpdateRequestLanguageKr     EmailUpdateRequestLanguage = "kr"
	EmailUpdateRequestLanguageKs     EmailUpdateRequestLanguage = "ks"
	EmailUpdateRequestLanguageKsIn   EmailUpdateRequestLanguage = "ks-in"
	EmailUpdateRequestLanguageKsb    EmailUpdateRequestLanguage = "ksb"
	EmailUpdateRequestLanguageKsbTz  EmailUpdateRequestLanguage = "ksb-tz"
	EmailUpdateRequestLanguageKsf    EmailUpdateRequestLanguage = "ksf"
	EmailUpdateRequestLanguageKsfCm  EmailUpdateRequestLanguage = "ksf-cm"
	EmailUpdateRequestLanguageKsh    EmailUpdateRequestLanguage = "ksh"
	EmailUpdateRequestLanguageKshDe  EmailUpdateRequestLanguage = "ksh-de"
	EmailUpdateRequestLanguageKu     EmailUpdateRequestLanguage = "ku"
	EmailUpdateRequestLanguageKuTr   EmailUpdateRequestLanguage = "ku-tr"
	EmailUpdateRequestLanguageKv     EmailUpdateRequestLanguage = "kv"
	EmailUpdateRequestLanguageKw     EmailUpdateRequestLanguage = "kw"
	EmailUpdateRequestLanguageKwGB   EmailUpdateRequestLanguage = "kw-gb"
	EmailUpdateRequestLanguageKxv    EmailUpdateRequestLanguage = "kxv"
	EmailUpdateRequestLanguageKxvIn  EmailUpdateRequestLanguage = "kxv-in"
	EmailUpdateRequestLanguageKy     EmailUpdateRequestLanguage = "ky"
	EmailUpdateRequestLanguageKyKg   EmailUpdateRequestLanguage = "ky-kg"
	EmailUpdateRequestLanguageLa     EmailUpdateRequestLanguage = "la"
	EmailUpdateRequestLanguageLag    EmailUpdateRequestLanguage = "lag"
	EmailUpdateRequestLanguageLagTz  EmailUpdateRequestLanguage = "lag-tz"
	EmailUpdateRequestLanguageLb     EmailUpdateRequestLanguage = "lb"
	EmailUpdateRequestLanguageLbLu   EmailUpdateRequestLanguage = "lb-lu"
	EmailUpdateRequestLanguageLg     EmailUpdateRequestLanguage = "lg"
	EmailUpdateRequestLanguageLgUg   EmailUpdateRequestLanguage = "lg-ug"
	EmailUpdateRequestLanguageLi     EmailUpdateRequestLanguage = "li"
	EmailUpdateRequestLanguageLij    EmailUpdateRequestLanguage = "lij"
	EmailUpdateRequestLanguageLijIt  EmailUpdateRequestLanguage = "lij-it"
	EmailUpdateRequestLanguageLkt    EmailUpdateRequestLanguage = "lkt"
	EmailUpdateRequestLanguageLktUs  EmailUpdateRequestLanguage = "lkt-us"
	EmailUpdateRequestLanguageLmo    EmailUpdateRequestLanguage = "lmo"
	EmailUpdateRequestLanguageLmoIt  EmailUpdateRequestLanguage = "lmo-it"
	EmailUpdateRequestLanguageLn     EmailUpdateRequestLanguage = "ln"
	EmailUpdateRequestLanguageLnAo   EmailUpdateRequestLanguage = "ln-ao"
	EmailUpdateRequestLanguageLnCd   EmailUpdateRequestLanguage = "ln-cd"
	EmailUpdateRequestLanguageLnCf   EmailUpdateRequestLanguage = "ln-cf"
	EmailUpdateRequestLanguageLnCg   EmailUpdateRequestLanguage = "ln-cg"
	EmailUpdateRequestLanguageLo     EmailUpdateRequestLanguage = "lo"
	EmailUpdateRequestLanguageLoLa   EmailUpdateRequestLanguage = "lo-la"
	EmailUpdateRequestLanguageLrc    EmailUpdateRequestLanguage = "lrc"
	EmailUpdateRequestLanguageLrcIq  EmailUpdateRequestLanguage = "lrc-iq"
	EmailUpdateRequestLanguageLrcIr  EmailUpdateRequestLanguage = "lrc-ir"
	EmailUpdateRequestLanguageLt     EmailUpdateRequestLanguage = "lt"
	EmailUpdateRequestLanguageLtLt   EmailUpdateRequestLanguage = "lt-lt"
	EmailUpdateRequestLanguageLu     EmailUpdateRequestLanguage = "lu"
	EmailUpdateRequestLanguageLuCd   EmailUpdateRequestLanguage = "lu-cd"
	EmailUpdateRequestLanguageLuo    EmailUpdateRequestLanguage = "luo"
	EmailUpdateRequestLanguageLuoKe  EmailUpdateRequestLanguage = "luo-ke"
	EmailUpdateRequestLanguageLuy    EmailUpdateRequestLanguage = "luy"
	EmailUpdateRequestLanguageLuyKe  EmailUpdateRequestLanguage = "luy-ke"
	EmailUpdateRequestLanguageLv     EmailUpdateRequestLanguage = "lv"
	EmailUpdateRequestLanguageLvLv   EmailUpdateRequestLanguage = "lv-lv"
	EmailUpdateRequestLanguageMai    EmailUpdateRequestLanguage = "mai"
	EmailUpdateRequestLanguageMaiIn  EmailUpdateRequestLanguage = "mai-in"
	EmailUpdateRequestLanguageMas    EmailUpdateRequestLanguage = "mas"
	EmailUpdateRequestLanguageMasKe  EmailUpdateRequestLanguage = "mas-ke"
	EmailUpdateRequestLanguageMasTz  EmailUpdateRequestLanguage = "mas-tz"
	EmailUpdateRequestLanguageMdf    EmailUpdateRequestLanguage = "mdf"
	EmailUpdateRequestLanguageMdfRu  EmailUpdateRequestLanguage = "mdf-ru"
	EmailUpdateRequestLanguageMer    EmailUpdateRequestLanguage = "mer"
	EmailUpdateRequestLanguageMerKe  EmailUpdateRequestLanguage = "mer-ke"
	EmailUpdateRequestLanguageMfe    EmailUpdateRequestLanguage = "mfe"
	EmailUpdateRequestLanguageMfeMu  EmailUpdateRequestLanguage = "mfe-mu"
	EmailUpdateRequestLanguageMg     EmailUpdateRequestLanguage = "mg"
	EmailUpdateRequestLanguageMgMg   EmailUpdateRequestLanguage = "mg-mg"
	EmailUpdateRequestLanguageMgh    EmailUpdateRequestLanguage = "mgh"
	EmailUpdateRequestLanguageMghMz  EmailUpdateRequestLanguage = "mgh-mz"
	EmailUpdateRequestLanguageMgo    EmailUpdateRequestLanguage = "mgo"
	EmailUpdateRequestLanguageMgoCm  EmailUpdateRequestLanguage = "mgo-cm"
	EmailUpdateRequestLanguageMh     EmailUpdateRequestLanguage = "mh"
	EmailUpdateRequestLanguageMi     EmailUpdateRequestLanguage = "mi"
	EmailUpdateRequestLanguageMiNz   EmailUpdateRequestLanguage = "mi-nz"
	EmailUpdateRequestLanguageMk     EmailUpdateRequestLanguage = "mk"
	EmailUpdateRequestLanguageMkMk   EmailUpdateRequestLanguage = "mk-mk"
	EmailUpdateRequestLanguageMl     EmailUpdateRequestLanguage = "ml"
	EmailUpdateRequestLanguageMlIn   EmailUpdateRequestLanguage = "ml-in"
	EmailUpdateRequestLanguageMn     EmailUpdateRequestLanguage = "mn"
	EmailUpdateRequestLanguageMnMn   EmailUpdateRequestLanguage = "mn-mn"
	EmailUpdateRequestLanguageMni    EmailUpdateRequestLanguage = "mni"
	EmailUpdateRequestLanguageMniIn  EmailUpdateRequestLanguage = "mni-in"
	EmailUpdateRequestLanguageMr     EmailUpdateRequestLanguage = "mr"
	EmailUpdateRequestLanguageMrIn   EmailUpdateRequestLanguage = "mr-in"
	EmailUpdateRequestLanguageMs     EmailUpdateRequestLanguage = "ms"
	EmailUpdateRequestLanguageMsBn   EmailUpdateRequestLanguage = "ms-bn"
	EmailUpdateRequestLanguageMsID   EmailUpdateRequestLanguage = "ms-id"
	EmailUpdateRequestLanguageMsMy   EmailUpdateRequestLanguage = "ms-my"
	EmailUpdateRequestLanguageMsSg   EmailUpdateRequestLanguage = "ms-sg"
	EmailUpdateRequestLanguageMt     EmailUpdateRequestLanguage = "mt"
	EmailUpdateRequestLanguageMtMt   EmailUpdateRequestLanguage = "mt-mt"
	EmailUpdateRequestLanguageMua    EmailUpdateRequestLanguage = "mua"
	EmailUpdateRequestLanguageMuaCm  EmailUpdateRequestLanguage = "mua-cm"
	EmailUpdateRequestLanguageMy     EmailUpdateRequestLanguage = "my"
	EmailUpdateRequestLanguageMyMm   EmailUpdateRequestLanguage = "my-mm"
	EmailUpdateRequestLanguageMzn    EmailUpdateRequestLanguage = "mzn"
	EmailUpdateRequestLanguageMznIr  EmailUpdateRequestLanguage = "mzn-ir"
	EmailUpdateRequestLanguageNa     EmailUpdateRequestLanguage = "na"
	EmailUpdateRequestLanguageNaq    EmailUpdateRequestLanguage = "naq"
	EmailUpdateRequestLanguageNaqNa  EmailUpdateRequestLanguage = "naq-na"
	EmailUpdateRequestLanguageNb     EmailUpdateRequestLanguage = "nb"
	EmailUpdateRequestLanguageNbNo   EmailUpdateRequestLanguage = "nb-no"
	EmailUpdateRequestLanguageNbSj   EmailUpdateRequestLanguage = "nb-sj"
	EmailUpdateRequestLanguageNd     EmailUpdateRequestLanguage = "nd"
	EmailUpdateRequestLanguageNdZw   EmailUpdateRequestLanguage = "nd-zw"
	EmailUpdateRequestLanguageNds    EmailUpdateRequestLanguage = "nds"
	EmailUpdateRequestLanguageNdsDe  EmailUpdateRequestLanguage = "nds-de"
	EmailUpdateRequestLanguageNdsNl  EmailUpdateRequestLanguage = "nds-nl"
	EmailUpdateRequestLanguageNe     EmailUpdateRequestLanguage = "ne"
	EmailUpdateRequestLanguageNeIn   EmailUpdateRequestLanguage = "ne-in"
	EmailUpdateRequestLanguageNeNp   EmailUpdateRequestLanguage = "ne-np"
	EmailUpdateRequestLanguageNg     EmailUpdateRequestLanguage = "ng"
	EmailUpdateRequestLanguageNl     EmailUpdateRequestLanguage = "nl"
	EmailUpdateRequestLanguageNlAw   EmailUpdateRequestLanguage = "nl-aw"
	EmailUpdateRequestLanguageNlBe   EmailUpdateRequestLanguage = "nl-be"
	EmailUpdateRequestLanguageNlBq   EmailUpdateRequestLanguage = "nl-bq"
	EmailUpdateRequestLanguageNlCh   EmailUpdateRequestLanguage = "nl-ch"
	EmailUpdateRequestLanguageNlCw   EmailUpdateRequestLanguage = "nl-cw"
	EmailUpdateRequestLanguageNlLu   EmailUpdateRequestLanguage = "nl-lu"
	EmailUpdateRequestLanguageNlNl   EmailUpdateRequestLanguage = "nl-nl"
	EmailUpdateRequestLanguageNlSr   EmailUpdateRequestLanguage = "nl-sr"
	EmailUpdateRequestLanguageNlSx   EmailUpdateRequestLanguage = "nl-sx"
	EmailUpdateRequestLanguageNmg    EmailUpdateRequestLanguage = "nmg"
	EmailUpdateRequestLanguageNmgCm  EmailUpdateRequestLanguage = "nmg-cm"
	EmailUpdateRequestLanguageNn     EmailUpdateRequestLanguage = "nn"
	EmailUpdateRequestLanguageNnNo   EmailUpdateRequestLanguage = "nn-no"
	EmailUpdateRequestLanguageNnh    EmailUpdateRequestLanguage = "nnh"
	EmailUpdateRequestLanguageNnhCm  EmailUpdateRequestLanguage = "nnh-cm"
	EmailUpdateRequestLanguageNo     EmailUpdateRequestLanguage = "no"
	EmailUpdateRequestLanguageNoNo   EmailUpdateRequestLanguage = "no-no"
	EmailUpdateRequestLanguageNqo    EmailUpdateRequestLanguage = "nqo"
	EmailUpdateRequestLanguageNqoGn  EmailUpdateRequestLanguage = "nqo-gn"
	EmailUpdateRequestLanguageNr     EmailUpdateRequestLanguage = "nr"
	EmailUpdateRequestLanguageNso    EmailUpdateRequestLanguage = "nso"
	EmailUpdateRequestLanguageNsoZa  EmailUpdateRequestLanguage = "nso-za"
	EmailUpdateRequestLanguageNus    EmailUpdateRequestLanguage = "nus"
	EmailUpdateRequestLanguageNusSS  EmailUpdateRequestLanguage = "nus-ss"
	EmailUpdateRequestLanguageNv     EmailUpdateRequestLanguage = "nv"
	EmailUpdateRequestLanguageNy     EmailUpdateRequestLanguage = "ny"
	EmailUpdateRequestLanguageNyn    EmailUpdateRequestLanguage = "nyn"
	EmailUpdateRequestLanguageNynUg  EmailUpdateRequestLanguage = "nyn-ug"
	EmailUpdateRequestLanguageOc     EmailUpdateRequestLanguage = "oc"
	EmailUpdateRequestLanguageOcEs   EmailUpdateRequestLanguage = "oc-es"
	EmailUpdateRequestLanguageOcFr   EmailUpdateRequestLanguage = "oc-fr"
	EmailUpdateRequestLanguageOj     EmailUpdateRequestLanguage = "oj"
	EmailUpdateRequestLanguageOm     EmailUpdateRequestLanguage = "om"
	EmailUpdateRequestLanguageOmEt   EmailUpdateRequestLanguage = "om-et"
	EmailUpdateRequestLanguageOmKe   EmailUpdateRequestLanguage = "om-ke"
	EmailUpdateRequestLanguageOr     EmailUpdateRequestLanguage = "or"
	EmailUpdateRequestLanguageOrIn   EmailUpdateRequestLanguage = "or-in"
	EmailUpdateRequestLanguageOs     EmailUpdateRequestLanguage = "os"
	EmailUpdateRequestLanguageOsGe   EmailUpdateRequestLanguage = "os-ge"
	EmailUpdateRequestLanguageOsRu   EmailUpdateRequestLanguage = "os-ru"
	EmailUpdateRequestLanguagePa     EmailUpdateRequestLanguage = "pa"
	EmailUpdateRequestLanguagePaIn   EmailUpdateRequestLanguage = "pa-in"
	EmailUpdateRequestLanguagePaPk   EmailUpdateRequestLanguage = "pa-pk"
	EmailUpdateRequestLanguagePcm    EmailUpdateRequestLanguage = "pcm"
	EmailUpdateRequestLanguagePcmNg  EmailUpdateRequestLanguage = "pcm-ng"
	EmailUpdateRequestLanguagePi     EmailUpdateRequestLanguage = "pi"
	EmailUpdateRequestLanguagePis    EmailUpdateRequestLanguage = "pis"
	EmailUpdateRequestLanguagePisSb  EmailUpdateRequestLanguage = "pis-sb"
	EmailUpdateRequestLanguagePl     EmailUpdateRequestLanguage = "pl"
	EmailUpdateRequestLanguagePlPl   EmailUpdateRequestLanguage = "pl-pl"
	EmailUpdateRequestLanguagePrg    EmailUpdateRequestLanguage = "prg"
	EmailUpdateRequestLanguagePrg001 EmailUpdateRequestLanguage = "prg-001"
	EmailUpdateRequestLanguagePs     EmailUpdateRequestLanguage = "ps"
	EmailUpdateRequestLanguagePsAf   EmailUpdateRequestLanguage = "ps-af"
	EmailUpdateRequestLanguagePsPk   EmailUpdateRequestLanguage = "ps-pk"
	EmailUpdateRequestLanguagePt     EmailUpdateRequestLanguage = "pt"
	EmailUpdateRequestLanguagePtAo   EmailUpdateRequestLanguage = "pt-ao"
	EmailUpdateRequestLanguagePtBr   EmailUpdateRequestLanguage = "pt-br"
	EmailUpdateRequestLanguagePtCh   EmailUpdateRequestLanguage = "pt-ch"
	EmailUpdateRequestLanguagePtCv   EmailUpdateRequestLanguage = "pt-cv"
	EmailUpdateRequestLanguagePtGq   EmailUpdateRequestLanguage = "pt-gq"
	EmailUpdateRequestLanguagePtGw   EmailUpdateRequestLanguage = "pt-gw"
	EmailUpdateRequestLanguagePtLu   EmailUpdateRequestLanguage = "pt-lu"
	EmailUpdateRequestLanguagePtMo   EmailUpdateRequestLanguage = "pt-mo"
	EmailUpdateRequestLanguagePtMz   EmailUpdateRequestLanguage = "pt-mz"
	EmailUpdateRequestLanguagePtPt   EmailUpdateRequestLanguage = "pt-pt"
	EmailUpdateRequestLanguagePtSt   EmailUpdateRequestLanguage = "pt-st"
	EmailUpdateRequestLanguagePtTl   EmailUpdateRequestLanguage = "pt-tl"
	EmailUpdateRequestLanguageQu     EmailUpdateRequestLanguage = "qu"
	EmailUpdateRequestLanguageQuBo   EmailUpdateRequestLanguage = "qu-bo"
	EmailUpdateRequestLanguageQuEc   EmailUpdateRequestLanguage = "qu-ec"
	EmailUpdateRequestLanguageQuPe   EmailUpdateRequestLanguage = "qu-pe"
	EmailUpdateRequestLanguageRaj    EmailUpdateRequestLanguage = "raj"
	EmailUpdateRequestLanguageRajIn  EmailUpdateRequestLanguage = "raj-in"
	EmailUpdateRequestLanguageRm     EmailUpdateRequestLanguage = "rm"
	EmailUpdateRequestLanguageRmCh   EmailUpdateRequestLanguage = "rm-ch"
	EmailUpdateRequestLanguageRn     EmailUpdateRequestLanguage = "rn"
	EmailUpdateRequestLanguageRnBi   EmailUpdateRequestLanguage = "rn-bi"
	EmailUpdateRequestLanguageRo     EmailUpdateRequestLanguage = "ro"
	EmailUpdateRequestLanguageRoMd   EmailUpdateRequestLanguage = "ro-md"
	EmailUpdateRequestLanguageRoRo   EmailUpdateRequestLanguage = "ro-ro"
	EmailUpdateRequestLanguageRof    EmailUpdateRequestLanguage = "rof"
	EmailUpdateRequestLanguageRofTz  EmailUpdateRequestLanguage = "rof-tz"
	EmailUpdateRequestLanguageRu     EmailUpdateRequestLanguage = "ru"
	EmailUpdateRequestLanguageRuBy   EmailUpdateRequestLanguage = "ru-by"
	EmailUpdateRequestLanguageRuKg   EmailUpdateRequestLanguage = "ru-kg"
	EmailUpdateRequestLanguageRuKz   EmailUpdateRequestLanguage = "ru-kz"
	EmailUpdateRequestLanguageRuMd   EmailUpdateRequestLanguage = "ru-md"
	EmailUpdateRequestLanguageRuRu   EmailUpdateRequestLanguage = "ru-ru"
	EmailUpdateRequestLanguageRuUa   EmailUpdateRequestLanguage = "ru-ua"
	EmailUpdateRequestLanguageRw     EmailUpdateRequestLanguage = "rw"
	EmailUpdateRequestLanguageRwRw   EmailUpdateRequestLanguage = "rw-rw"
	EmailUpdateRequestLanguageRwk    EmailUpdateRequestLanguage = "rwk"
	EmailUpdateRequestLanguageRwkTz  EmailUpdateRequestLanguage = "rwk-tz"
	EmailUpdateRequestLanguageSa     EmailUpdateRequestLanguage = "sa"
	EmailUpdateRequestLanguageSaIn   EmailUpdateRequestLanguage = "sa-in"
	EmailUpdateRequestLanguageSah    EmailUpdateRequestLanguage = "sah"
	EmailUpdateRequestLanguageSahRu  EmailUpdateRequestLanguage = "sah-ru"
	EmailUpdateRequestLanguageSaq    EmailUpdateRequestLanguage = "saq"
	EmailUpdateRequestLanguageSaqKe  EmailUpdateRequestLanguage = "saq-ke"
	EmailUpdateRequestLanguageSat    EmailUpdateRequestLanguage = "sat"
	EmailUpdateRequestLanguageSatIn  EmailUpdateRequestLanguage = "sat-in"
	EmailUpdateRequestLanguageSbp    EmailUpdateRequestLanguage = "sbp"
	EmailUpdateRequestLanguageSbpTz  EmailUpdateRequestLanguage = "sbp-tz"
	EmailUpdateRequestLanguageSc     EmailUpdateRequestLanguage = "sc"
	EmailUpdateRequestLanguageScIt   EmailUpdateRequestLanguage = "sc-it"
	EmailUpdateRequestLanguageSd     EmailUpdateRequestLanguage = "sd"
	EmailUpdateRequestLanguageSdIn   EmailUpdateRequestLanguage = "sd-in"
	EmailUpdateRequestLanguageSdPk   EmailUpdateRequestLanguage = "sd-pk"
	EmailUpdateRequestLanguageSe     EmailUpdateRequestLanguage = "se"
	EmailUpdateRequestLanguageSeFi   EmailUpdateRequestLanguage = "se-fi"
	EmailUpdateRequestLanguageSeNo   EmailUpdateRequestLanguage = "se-no"
	EmailUpdateRequestLanguageSeSe   EmailUpdateRequestLanguage = "se-se"
	EmailUpdateRequestLanguageSeh    EmailUpdateRequestLanguage = "seh"
	EmailUpdateRequestLanguageSehMz  EmailUpdateRequestLanguage = "seh-mz"
	EmailUpdateRequestLanguageSes    EmailUpdateRequestLanguage = "ses"
	EmailUpdateRequestLanguageSesMl  EmailUpdateRequestLanguage = "ses-ml"
	EmailUpdateRequestLanguageSg     EmailUpdateRequestLanguage = "sg"
	EmailUpdateRequestLanguageSgCf   EmailUpdateRequestLanguage = "sg-cf"
	EmailUpdateRequestLanguageShi    EmailUpdateRequestLanguage = "shi"
	EmailUpdateRequestLanguageShiMa  EmailUpdateRequestLanguage = "shi-ma"
	EmailUpdateRequestLanguageSi     EmailUpdateRequestLanguage = "si"
	EmailUpdateRequestLanguageSiLk   EmailUpdateRequestLanguage = "si-lk"
	EmailUpdateRequestLanguageSk     EmailUpdateRequestLanguage = "sk"
	EmailUpdateRequestLanguageSkSk   EmailUpdateRequestLanguage = "sk-sk"
	EmailUpdateRequestLanguageSl     EmailUpdateRequestLanguage = "sl"
	EmailUpdateRequestLanguageSlSi   EmailUpdateRequestLanguage = "sl-si"
	EmailUpdateRequestLanguageSm     EmailUpdateRequestLanguage = "sm"
	EmailUpdateRequestLanguageSmn    EmailUpdateRequestLanguage = "smn"
	EmailUpdateRequestLanguageSmnFi  EmailUpdateRequestLanguage = "smn-fi"
	EmailUpdateRequestLanguageSMS    EmailUpdateRequestLanguage = "sms"
	EmailUpdateRequestLanguageSMSFi  EmailUpdateRequestLanguage = "sms-fi"
	EmailUpdateRequestLanguageSn     EmailUpdateRequestLanguage = "sn"
	EmailUpdateRequestLanguageSnZw   EmailUpdateRequestLanguage = "sn-zw"
	EmailUpdateRequestLanguageSo     EmailUpdateRequestLanguage = "so"
	EmailUpdateRequestLanguageSoDj   EmailUpdateRequestLanguage = "so-dj"
	EmailUpdateRequestLanguageSoEt   EmailUpdateRequestLanguage = "so-et"
	EmailUpdateRequestLanguageSoKe   EmailUpdateRequestLanguage = "so-ke"
	EmailUpdateRequestLanguageSoSo   EmailUpdateRequestLanguage = "so-so"
	EmailUpdateRequestLanguageSq     EmailUpdateRequestLanguage = "sq"
	EmailUpdateRequestLanguageSqAl   EmailUpdateRequestLanguage = "sq-al"
	EmailUpdateRequestLanguageSqMk   EmailUpdateRequestLanguage = "sq-mk"
	EmailUpdateRequestLanguageSqXk   EmailUpdateRequestLanguage = "sq-xk"
	EmailUpdateRequestLanguageSr     EmailUpdateRequestLanguage = "sr"
	EmailUpdateRequestLanguageSrBa   EmailUpdateRequestLanguage = "sr-ba"
	EmailUpdateRequestLanguageSrCs   EmailUpdateRequestLanguage = "sr-cs"
	EmailUpdateRequestLanguageSrMe   EmailUpdateRequestLanguage = "sr-me"
	EmailUpdateRequestLanguageSrRs   EmailUpdateRequestLanguage = "sr-rs"
	EmailUpdateRequestLanguageSrXk   EmailUpdateRequestLanguage = "sr-xk"
	EmailUpdateRequestLanguageSS     EmailUpdateRequestLanguage = "ss"
	EmailUpdateRequestLanguageSt     EmailUpdateRequestLanguage = "st"
	EmailUpdateRequestLanguageStLs   EmailUpdateRequestLanguage = "st-ls"
	EmailUpdateRequestLanguageStZa   EmailUpdateRequestLanguage = "st-za"
	EmailUpdateRequestLanguageSu     EmailUpdateRequestLanguage = "su"
	EmailUpdateRequestLanguageSuID   EmailUpdateRequestLanguage = "su-id"
	EmailUpdateRequestLanguageSv     EmailUpdateRequestLanguage = "sv"
	EmailUpdateRequestLanguageSvAx   EmailUpdateRequestLanguage = "sv-ax"
	EmailUpdateRequestLanguageSvFi   EmailUpdateRequestLanguage = "sv-fi"
	EmailUpdateRequestLanguageSvSe   EmailUpdateRequestLanguage = "sv-se"
	EmailUpdateRequestLanguageSw     EmailUpdateRequestLanguage = "sw"
	EmailUpdateRequestLanguageSwCd   EmailUpdateRequestLanguage = "sw-cd"
	EmailUpdateRequestLanguageSwKe   EmailUpdateRequestLanguage = "sw-ke"
	EmailUpdateRequestLanguageSwTz   EmailUpdateRequestLanguage = "sw-tz"
	EmailUpdateRequestLanguageSwUg   EmailUpdateRequestLanguage = "sw-ug"
	EmailUpdateRequestLanguageSy     EmailUpdateRequestLanguage = "sy"
	EmailUpdateRequestLanguageSyr    EmailUpdateRequestLanguage = "syr"
	EmailUpdateRequestLanguageSyrIq  EmailUpdateRequestLanguage = "syr-iq"
	EmailUpdateRequestLanguageSyrSy  EmailUpdateRequestLanguage = "syr-sy"
	EmailUpdateRequestLanguageSzl    EmailUpdateRequestLanguage = "szl"
	EmailUpdateRequestLanguageSzlPl  EmailUpdateRequestLanguage = "szl-pl"
	EmailUpdateRequestLanguageTa     EmailUpdateRequestLanguage = "ta"
	EmailUpdateRequestLanguageTaIn   EmailUpdateRequestLanguage = "ta-in"
	EmailUpdateRequestLanguageTaLk   EmailUpdateRequestLanguage = "ta-lk"
	EmailUpdateRequestLanguageTaMy   EmailUpdateRequestLanguage = "ta-my"
	EmailUpdateRequestLanguageTaSg   EmailUpdateRequestLanguage = "ta-sg"
	EmailUpdateRequestLanguageTe     EmailUpdateRequestLanguage = "te"
	EmailUpdateRequestLanguageTeIn   EmailUpdateRequestLanguage = "te-in"
	EmailUpdateRequestLanguageTeo    EmailUpdateRequestLanguage = "teo"
	EmailUpdateRequestLanguageTeoKe  EmailUpdateRequestLanguage = "teo-ke"
	EmailUpdateRequestLanguageTeoUg  EmailUpdateRequestLanguage = "teo-ug"
	EmailUpdateRequestLanguageTg     EmailUpdateRequestLanguage = "tg"
	EmailUpdateRequestLanguageTgTj   EmailUpdateRequestLanguage = "tg-tj"
	EmailUpdateRequestLanguageTh     EmailUpdateRequestLanguage = "th"
	EmailUpdateRequestLanguageThTh   EmailUpdateRequestLanguage = "th-th"
	EmailUpdateRequestLanguageTi     EmailUpdateRequestLanguage = "ti"
	EmailUpdateRequestLanguageTiEr   EmailUpdateRequestLanguage = "ti-er"
	EmailUpdateRequestLanguageTiEt   EmailUpdateRequestLanguage = "ti-et"
	EmailUpdateRequestLanguageTk     EmailUpdateRequestLanguage = "tk"
	EmailUpdateRequestLanguageTkTm   EmailUpdateRequestLanguage = "tk-tm"
	EmailUpdateRequestLanguageTl     EmailUpdateRequestLanguage = "tl"
	EmailUpdateRequestLanguageTn     EmailUpdateRequestLanguage = "tn"
	EmailUpdateRequestLanguageTnBw   EmailUpdateRequestLanguage = "tn-bw"
	EmailUpdateRequestLanguageTnZa   EmailUpdateRequestLanguage = "tn-za"
	EmailUpdateRequestLanguageTo     EmailUpdateRequestLanguage = "to"
	EmailUpdateRequestLanguageToTo   EmailUpdateRequestLanguage = "to-to"
	EmailUpdateRequestLanguageTok    EmailUpdateRequestLanguage = "tok"
	EmailUpdateRequestLanguageTok001 EmailUpdateRequestLanguage = "tok-001"
	EmailUpdateRequestLanguageTr     EmailUpdateRequestLanguage = "tr"
	EmailUpdateRequestLanguageTrCy   EmailUpdateRequestLanguage = "tr-cy"
	EmailUpdateRequestLanguageTrTr   EmailUpdateRequestLanguage = "tr-tr"
	EmailUpdateRequestLanguageTs     EmailUpdateRequestLanguage = "ts"
	EmailUpdateRequestLanguageTt     EmailUpdateRequestLanguage = "tt"
	EmailUpdateRequestLanguageTtRu   EmailUpdateRequestLanguage = "tt-ru"
	EmailUpdateRequestLanguageTw     EmailUpdateRequestLanguage = "tw"
	EmailUpdateRequestLanguageTwq    EmailUpdateRequestLanguage = "twq"
	EmailUpdateRequestLanguageTwqNe  EmailUpdateRequestLanguage = "twq-ne"
	EmailUpdateRequestLanguageTy     EmailUpdateRequestLanguage = "ty"
	EmailUpdateRequestLanguageTzm    EmailUpdateRequestLanguage = "tzm"
	EmailUpdateRequestLanguageTzmMa  EmailUpdateRequestLanguage = "tzm-ma"
	EmailUpdateRequestLanguageUg     EmailUpdateRequestLanguage = "ug"
	EmailUpdateRequestLanguageUgCn   EmailUpdateRequestLanguage = "ug-cn"
	EmailUpdateRequestLanguageUk     EmailUpdateRequestLanguage = "uk"
	EmailUpdateRequestLanguageUkUa   EmailUpdateRequestLanguage = "uk-ua"
	EmailUpdateRequestLanguageUr     EmailUpdateRequestLanguage = "ur"
	EmailUpdateRequestLanguageUrIn   EmailUpdateRequestLanguage = "ur-in"
	EmailUpdateRequestLanguageUrPk   EmailUpdateRequestLanguage = "ur-pk"
	EmailUpdateRequestLanguageUz     EmailUpdateRequestLanguage = "uz"
	EmailUpdateRequestLanguageUzAf   EmailUpdateRequestLanguage = "uz-af"
	EmailUpdateRequestLanguageUzUz   EmailUpdateRequestLanguage = "uz-uz"
	EmailUpdateRequestLanguageVai    EmailUpdateRequestLanguage = "vai"
	EmailUpdateRequestLanguageVaiLr  EmailUpdateRequestLanguage = "vai-lr"
	EmailUpdateRequestLanguageVe     EmailUpdateRequestLanguage = "ve"
	EmailUpdateRequestLanguageVec    EmailUpdateRequestLanguage = "vec"
	EmailUpdateRequestLanguageVecIt  EmailUpdateRequestLanguage = "vec-it"
	EmailUpdateRequestLanguageVi     EmailUpdateRequestLanguage = "vi"
	EmailUpdateRequestLanguageViVn   EmailUpdateRequestLanguage = "vi-vn"
	EmailUpdateRequestLanguageVmw    EmailUpdateRequestLanguage = "vmw"
	EmailUpdateRequestLanguageVmwMz  EmailUpdateRequestLanguage = "vmw-mz"
	EmailUpdateRequestLanguageVo     EmailUpdateRequestLanguage = "vo"
	EmailUpdateRequestLanguageVo001  EmailUpdateRequestLanguage = "vo-001"
	EmailUpdateRequestLanguageVun    EmailUpdateRequestLanguage = "vun"
	EmailUpdateRequestLanguageVunTz  EmailUpdateRequestLanguage = "vun-tz"
	EmailUpdateRequestLanguageWa     EmailUpdateRequestLanguage = "wa"
	EmailUpdateRequestLanguageWae    EmailUpdateRequestLanguage = "wae"
	EmailUpdateRequestLanguageWaeCh  EmailUpdateRequestLanguage = "wae-ch"
	EmailUpdateRequestLanguageWo     EmailUpdateRequestLanguage = "wo"
	EmailUpdateRequestLanguageWoSn   EmailUpdateRequestLanguage = "wo-sn"
	EmailUpdateRequestLanguageXh     EmailUpdateRequestLanguage = "xh"
	EmailUpdateRequestLanguageXhZa   EmailUpdateRequestLanguage = "xh-za"
	EmailUpdateRequestLanguageXnr    EmailUpdateRequestLanguage = "xnr"
	EmailUpdateRequestLanguageXnrIn  EmailUpdateRequestLanguage = "xnr-in"
	EmailUpdateRequestLanguageXog    EmailUpdateRequestLanguage = "xog"
	EmailUpdateRequestLanguageXogUg  EmailUpdateRequestLanguage = "xog-ug"
	EmailUpdateRequestLanguageYav    EmailUpdateRequestLanguage = "yav"
	EmailUpdateRequestLanguageYavCm  EmailUpdateRequestLanguage = "yav-cm"
	EmailUpdateRequestLanguageYi     EmailUpdateRequestLanguage = "yi"
	EmailUpdateRequestLanguageYi001  EmailUpdateRequestLanguage = "yi-001"
	EmailUpdateRequestLanguageYiUa   EmailUpdateRequestLanguage = "yi-ua"
	EmailUpdateRequestLanguageYo     EmailUpdateRequestLanguage = "yo"
	EmailUpdateRequestLanguageYoBj   EmailUpdateRequestLanguage = "yo-bj"
	EmailUpdateRequestLanguageYoNg   EmailUpdateRequestLanguage = "yo-ng"
	EmailUpdateRequestLanguageYrl    EmailUpdateRequestLanguage = "yrl"
	EmailUpdateRequestLanguageYrlBr  EmailUpdateRequestLanguage = "yrl-br"
	EmailUpdateRequestLanguageYrlCo  EmailUpdateRequestLanguage = "yrl-co"
	EmailUpdateRequestLanguageYrlVe  EmailUpdateRequestLanguage = "yrl-ve"
	EmailUpdateRequestLanguageYue    EmailUpdateRequestLanguage = "yue"
	EmailUpdateRequestLanguageYueCn  EmailUpdateRequestLanguage = "yue-cn"
	EmailUpdateRequestLanguageYueHk  EmailUpdateRequestLanguage = "yue-hk"
	EmailUpdateRequestLanguageYueMo  EmailUpdateRequestLanguage = "yue-mo"
	EmailUpdateRequestLanguageZa     EmailUpdateRequestLanguage = "za"
	EmailUpdateRequestLanguageZaCn   EmailUpdateRequestLanguage = "za-cn"
	EmailUpdateRequestLanguageZgh    EmailUpdateRequestLanguage = "zgh"
	EmailUpdateRequestLanguageZghMa  EmailUpdateRequestLanguage = "zgh-ma"
	EmailUpdateRequestLanguageZh     EmailUpdateRequestLanguage = "zh"
	EmailUpdateRequestLanguageZhCn   EmailUpdateRequestLanguage = "zh-cn"
	EmailUpdateRequestLanguageZhHans EmailUpdateRequestLanguage = "zh-hans"
	EmailUpdateRequestLanguageZhHant EmailUpdateRequestLanguage = "zh-hant"
	EmailUpdateRequestLanguageZhHk   EmailUpdateRequestLanguage = "zh-hk"
	EmailUpdateRequestLanguageZhMo   EmailUpdateRequestLanguage = "zh-mo"
	EmailUpdateRequestLanguageZhMy   EmailUpdateRequestLanguage = "zh-my"
	EmailUpdateRequestLanguageZhSg   EmailUpdateRequestLanguage = "zh-sg"
	EmailUpdateRequestLanguageZhTw   EmailUpdateRequestLanguage = "zh-tw"
	EmailUpdateRequestLanguageZu     EmailUpdateRequestLanguage = "zu"
	EmailUpdateRequestLanguageZuZa   EmailUpdateRequestLanguage = "zu-za"
)

// The email state.
type EmailUpdateRequestState string

const (
	EmailUpdateRequestStateAgentGenerated          EmailUpdateRequestState = "AGENT_GENERATED"
	EmailUpdateRequestStateAutomated               EmailUpdateRequestState = "AUTOMATED"
	EmailUpdateRequestStateAutomatedAb             EmailUpdateRequestState = "AUTOMATED_AB"
	EmailUpdateRequestStateAutomatedAbVariant      EmailUpdateRequestState = "AUTOMATED_AB_VARIANT"
	EmailUpdateRequestStateAutomatedDraft          EmailUpdateRequestState = "AUTOMATED_DRAFT"
	EmailUpdateRequestStateAutomatedDraftAb        EmailUpdateRequestState = "AUTOMATED_DRAFT_AB"
	EmailUpdateRequestStateAutomatedDraftAbvariant EmailUpdateRequestState = "AUTOMATED_DRAFT_ABVARIANT"
	EmailUpdateRequestStateAutomatedForForm        EmailUpdateRequestState = "AUTOMATED_FOR_FORM"
	EmailUpdateRequestStateAutomatedForFormBuffer  EmailUpdateRequestState = "AUTOMATED_FOR_FORM_BUFFER"
	EmailUpdateRequestStateAutomatedForFormDraft   EmailUpdateRequestState = "AUTOMATED_FOR_FORM_DRAFT"
	EmailUpdateRequestStateAutomatedForFormLegacy  EmailUpdateRequestState = "AUTOMATED_FOR_FORM_LEGACY"
	EmailUpdateRequestStateAutomatedLoserAbvariant EmailUpdateRequestState = "AUTOMATED_LOSER_ABVARIANT"
	EmailUpdateRequestStateAutomatedSending        EmailUpdateRequestState = "AUTOMATED_SENDING"
	EmailUpdateRequestStateBlogEmailDraft          EmailUpdateRequestState = "BLOG_EMAIL_DRAFT"
	EmailUpdateRequestStateBlogEmailPublished      EmailUpdateRequestState = "BLOG_EMAIL_PUBLISHED"
	EmailUpdateRequestStateDraft                   EmailUpdateRequestState = "DRAFT"
	EmailUpdateRequestStateDraftAb                 EmailUpdateRequestState = "DRAFT_AB"
	EmailUpdateRequestStateDraftAbVariant          EmailUpdateRequestState = "DRAFT_AB_VARIANT"
	EmailUpdateRequestStateError                   EmailUpdateRequestState = "ERROR"
	EmailUpdateRequestStateLoserAbVariant          EmailUpdateRequestState = "LOSER_AB_VARIANT"
	EmailUpdateRequestStatePageStub                EmailUpdateRequestState = "PAGE_STUB"
	EmailUpdateRequestStatePreProcessing           EmailUpdateRequestState = "PRE_PROCESSING"
	EmailUpdateRequestStateProcessing              EmailUpdateRequestState = "PROCESSING"
	EmailUpdateRequestStatePublished               EmailUpdateRequestState = "PUBLISHED"
	EmailUpdateRequestStatePublishedAb             EmailUpdateRequestState = "PUBLISHED_AB"
	EmailUpdateRequestStatePublishedAbVariant      EmailUpdateRequestState = "PUBLISHED_AB_VARIANT"
	EmailUpdateRequestStatePublishedOrScheduled    EmailUpdateRequestState = "PUBLISHED_OR_SCHEDULED"
	EmailUpdateRequestStateRssToEmailDraft         EmailUpdateRequestState = "RSS_TO_EMAIL_DRAFT"
	EmailUpdateRequestStateRssToEmailPublished     EmailUpdateRequestState = "RSS_TO_EMAIL_PUBLISHED"
	EmailUpdateRequestStateScheduled               EmailUpdateRequestState = "SCHEDULED"
	EmailUpdateRequestStateScheduledAb             EmailUpdateRequestState = "SCHEDULED_AB"
	EmailUpdateRequestStateScheduledOrPublished    EmailUpdateRequestState = "SCHEDULED_OR_PUBLISHED"
)

// The email subcategory.
type EmailUpdateRequestSubcategory string

const (
	EmailUpdateRequestSubcategoryAbLoserVariant                 EmailUpdateRequestSubcategory = "ab_loser_variant"
	EmailUpdateRequestSubcategoryAbLoserVariantSitePage         EmailUpdateRequestSubcategory = "ab_loser_variant_site_page"
	EmailUpdateRequestSubcategoryAbMaster                       EmailUpdateRequestSubcategory = "ab_master"
	EmailUpdateRequestSubcategoryAbMasterSitePage               EmailUpdateRequestSubcategory = "ab_master_site_page"
	EmailUpdateRequestSubcategoryAbVariant                      EmailUpdateRequestSubcategory = "ab_variant"
	EmailUpdateRequestSubcategoryAbVariantSitePage              EmailUpdateRequestSubcategory = "ab_variant_site_page"
	EmailUpdateRequestSubcategoryAutomated                      EmailUpdateRequestSubcategory = "automated"
	EmailUpdateRequestSubcategoryAutomatedAbMaster              EmailUpdateRequestSubcategory = "automated_ab_master"
	EmailUpdateRequestSubcategoryAutomatedAbVariant             EmailUpdateRequestSubcategory = "automated_ab_variant"
	EmailUpdateRequestSubcategoryAutomatedForCrm                EmailUpdateRequestSubcategory = "automated_for_crm"
	EmailUpdateRequestSubcategoryAutomatedForCustomSurvey       EmailUpdateRequestSubcategory = "automated_for_custom_survey"
	EmailUpdateRequestSubcategoryAutomatedForDeal               EmailUpdateRequestSubcategory = "automated_for_deal"
	EmailUpdateRequestSubcategoryAutomatedForFeedbackCes        EmailUpdateRequestSubcategory = "automated_for_feedback_ces"
	EmailUpdateRequestSubcategoryAutomatedForFeedbackCustom     EmailUpdateRequestSubcategory = "automated_for_feedback_custom"
	EmailUpdateRequestSubcategoryAutomatedForFeedbackNps        EmailUpdateRequestSubcategory = "automated_for_feedback_nps"
	EmailUpdateRequestSubcategoryAutomatedForForm               EmailUpdateRequestSubcategory = "automated_for_form"
	EmailUpdateRequestSubcategoryAutomatedForFormBuffer         EmailUpdateRequestSubcategory = "automated_for_form_buffer"
	EmailUpdateRequestSubcategoryAutomatedForFormDraft          EmailUpdateRequestSubcategory = "automated_for_form_draft"
	EmailUpdateRequestSubcategoryAutomatedForFormLegacy         EmailUpdateRequestSubcategory = "automated_for_form_legacy"
	EmailUpdateRequestSubcategoryAutomatedForLeadflow           EmailUpdateRequestSubcategory = "automated_for_leadflow"
	EmailUpdateRequestSubcategoryAutomatedForTicket             EmailUpdateRequestSubcategory = "automated_for_ticket"
	EmailUpdateRequestSubcategoryBatch                          EmailUpdateRequestSubcategory = "batch"
	EmailUpdateRequestSubcategoryBlogArticleInstanceLayout      EmailUpdateRequestSubcategory = "blog_article_instance_layout"
	EmailUpdateRequestSubcategoryBlogArticleListing             EmailUpdateRequestSubcategory = "blog_article_listing"
	EmailUpdateRequestSubcategoryBlogAuthorDetail               EmailUpdateRequestSubcategory = "blog_author_detail"
	EmailUpdateRequestSubcategoryBlogEmail                      EmailUpdateRequestSubcategory = "blog_email"
	EmailUpdateRequestSubcategoryBlogEmailChild                 EmailUpdateRequestSubcategory = "blog_email_child"
	EmailUpdateRequestSubcategoryCaseStudy                      EmailUpdateRequestSubcategory = "case_study"
	EmailUpdateRequestSubcategoryCaseStudyInstanceLayout        EmailUpdateRequestSubcategory = "case_study_instance_layout"
	EmailUpdateRequestSubcategoryCaseStudyListing               EmailUpdateRequestSubcategory = "case_study_listing"
	EmailUpdateRequestSubcategoryDiscardableStub                EmailUpdateRequestSubcategory = "discardable_stub"
	EmailUpdateRequestSubcategoryImportedBlogPost               EmailUpdateRequestSubcategory = "imported_blog_post"
	EmailUpdateRequestSubcategoryKB404Page                      EmailUpdateRequestSubcategory = "kb_404_page"
	EmailUpdateRequestSubcategoryKBArticleInstanceLayout        EmailUpdateRequestSubcategory = "kb_article_instance_layout"
	EmailUpdateRequestSubcategoryKBListing                      EmailUpdateRequestSubcategory = "kb_listing"
	EmailUpdateRequestSubcategoryKBSearchResults                EmailUpdateRequestSubcategory = "kb_search_results"
	EmailUpdateRequestSubcategoryKBSupportForm                  EmailUpdateRequestSubcategory = "kb_support_form"
	EmailUpdateRequestSubcategoryLandingPage                    EmailUpdateRequestSubcategory = "landing_page"
	EmailUpdateRequestSubcategoryLegacyBlogPost                 EmailUpdateRequestSubcategory = "legacy_blog_post"
	EmailUpdateRequestSubcategoryLegacyPage                     EmailUpdateRequestSubcategory = "legacy_page"
	EmailUpdateRequestSubcategoryLocaltime                      EmailUpdateRequestSubcategory = "localtime"
	EmailUpdateRequestSubcategoryManagePreferencesEmail         EmailUpdateRequestSubcategory = "manage_preferences_email"
	EmailUpdateRequestSubcategoryMarketingSingleSendAPI         EmailUpdateRequestSubcategory = "marketing_single_send_api"
	EmailUpdateRequestSubcategoryMembershipEmailVerification    EmailUpdateRequestSubcategory = "membership_email_verification"
	EmailUpdateRequestSubcategoryMembershipFollowUp             EmailUpdateRequestSubcategory = "membership_follow_up"
	EmailUpdateRequestSubcategoryMembershipOtpLogin             EmailUpdateRequestSubcategory = "membership_otp_login"
	EmailUpdateRequestSubcategoryMembershipPasswordReset        EmailUpdateRequestSubcategory = "membership_password_reset"
	EmailUpdateRequestSubcategoryMembershipPasswordSaved        EmailUpdateRequestSubcategory = "membership_password_saved"
	EmailUpdateRequestSubcategoryMembershipPasswordlessAuth     EmailUpdateRequestSubcategory = "membership_passwordless_auth"
	EmailUpdateRequestSubcategoryMembershipRegistration         EmailUpdateRequestSubcategory = "membership_registration"
	EmailUpdateRequestSubcategoryMembershipRegistrationFollowUp EmailUpdateRequestSubcategory = "membership_registration_follow_up"
	EmailUpdateRequestSubcategoryMembershipVerification         EmailUpdateRequestSubcategory = "membership_verification"
	EmailUpdateRequestSubcategoryNormalBlogPost                 EmailUpdateRequestSubcategory = "normal_blog_post"
	EmailUpdateRequestSubcategoryOptinEmail                     EmailUpdateRequestSubcategory = "optin_email"
	EmailUpdateRequestSubcategoryOptinFollowupEmail             EmailUpdateRequestSubcategory = "optin_followup_email"
	EmailUpdateRequestSubcategoryPageInstanceLayout             EmailUpdateRequestSubcategory = "page_instance_layout"
	EmailUpdateRequestSubcategoryPageStub                       EmailUpdateRequestSubcategory = "page_stub"
	EmailUpdateRequestSubcategoryPerformableLandingPage         EmailUpdateRequestSubcategory = "performable_landing_page"
	EmailUpdateRequestSubcategoryPerformableLandingPageCutover  EmailUpdateRequestSubcategory = "performable_landing_page_cutover"
	EmailUpdateRequestSubcategoryPodcastInstanceLayout          EmailUpdateRequestSubcategory = "podcast_instance_layout"
	EmailUpdateRequestSubcategoryPodcastListing                 EmailUpdateRequestSubcategory = "podcast_listing"
	EmailUpdateRequestSubcategoryPortalContent                  EmailUpdateRequestSubcategory = "portal_content"
	EmailUpdateRequestSubcategoryResubscribeConfirmationEmail   EmailUpdateRequestSubcategory = "resubscribe_confirmation_email"
	EmailUpdateRequestSubcategoryResubscribeEmail               EmailUpdateRequestSubcategory = "resubscribe_email"
	EmailUpdateRequestSubcategoryRssToEmail                     EmailUpdateRequestSubcategory = "rss_to_email"
	EmailUpdateRequestSubcategoryRssToEmailChild                EmailUpdateRequestSubcategory = "rss_to_email_child"
	EmailUpdateRequestSubcategoryScpInstanceLayoutPage          EmailUpdateRequestSubcategory = "scp_instance_layout_page"
	EmailUpdateRequestSubcategoryScpStaticPage                  EmailUpdateRequestSubcategory = "scp_static_page"
	EmailUpdateRequestSubcategorySingleSendAPI                  EmailUpdateRequestSubcategory = "single_send_api"
	EmailUpdateRequestSubcategorySitePage                       EmailUpdateRequestSubcategory = "site_page"
	EmailUpdateRequestSubcategorySmtpToken                      EmailUpdateRequestSubcategory = "smtp_token"
	EmailUpdateRequestSubcategoryStagedPage                     EmailUpdateRequestSubcategory = "staged_page"
	EmailUpdateRequestSubcategoryTicketClosedKickbackEmail      EmailUpdateRequestSubcategory = "ticket_closed_kickback_email"
	EmailUpdateRequestSubcategoryTicketOpenedKickbackEmail      EmailUpdateRequestSubcategory = "ticket_opened_kickback_email"
	EmailUpdateRequestSubcategoryTicketPipelineAutomated        EmailUpdateRequestSubcategory = "ticket_pipeline_automated"
	EmailUpdateRequestSubcategoryUnknown                        EmailUpdateRequestSubcategory = "UNKNOWN"
	EmailUpdateRequestSubcategoryUnsubscribeConfirmationEmail   EmailUpdateRequestSubcategory = "unsubscribe_confirmation_email"
	EmailUpdateRequestSubcategoryWebInteractive                 EmailUpdateRequestSubcategory = "web_interactive"
)

type Interval struct {
	// The end timestamp of the interval, in ISO8601 format.
	End time.Time `json:"end" api:"required" format:"date-time"`
	// The start timestamp of the interval, in ISO8601 format.
	Start time.Time `json:"start" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		End         respjson.Field
		Start       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Interval) RawJSON() string { return r.JSON.raw }
func (r *Interval) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PublicButtonStyleSettings struct {
	BackgroundColor any             `json:"backgroundColor"`
	CornerRadius    int64           `json:"cornerRadius"`
	FontStyle       PublicFontStyle `json:"fontStyle"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BackgroundColor respjson.Field
		CornerRadius    respjson.Field
		FontStyle       respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PublicButtonStyleSettings) RawJSON() string { return r.JSON.raw }
func (r *PublicButtonStyleSettings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this PublicButtonStyleSettings to a
// PublicButtonStyleSettingsParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// PublicButtonStyleSettingsParam.Overrides()
func (r PublicButtonStyleSettings) ToParam() PublicButtonStyleSettingsParam {
	return param.Override[PublicButtonStyleSettingsParam](json.RawMessage(r.RawJSON()))
}

type PublicButtonStyleSettingsParam struct {
	CornerRadius    param.Opt[int64]     `json:"cornerRadius,omitzero"`
	BackgroundColor any                  `json:"backgroundColor,omitzero"`
	FontStyle       PublicFontStyleParam `json:"fontStyle,omitzero"`
	paramObj
}

func (r PublicButtonStyleSettingsParam) MarshalJSON() (data []byte, err error) {
	type shadow PublicButtonStyleSettingsParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PublicButtonStyleSettingsParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PublicDividerStyleSettings struct {
	Color    any    `json:"color"`
	Height   int64  `json:"height"`
	LineType string `json:"lineType"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Color       respjson.Field
		Height      respjson.Field
		LineType    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PublicDividerStyleSettings) RawJSON() string { return r.JSON.raw }
func (r *PublicDividerStyleSettings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this PublicDividerStyleSettings to a
// PublicDividerStyleSettingsParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// PublicDividerStyleSettingsParam.Overrides()
func (r PublicDividerStyleSettings) ToParam() PublicDividerStyleSettingsParam {
	return param.Override[PublicDividerStyleSettingsParam](json.RawMessage(r.RawJSON()))
}

type PublicDividerStyleSettingsParam struct {
	Height   param.Opt[int64]  `json:"height,omitzero"`
	LineType param.Opt[string] `json:"lineType,omitzero"`
	Color    any               `json:"color,omitzero"`
	paramObj
}

func (r PublicDividerStyleSettingsParam) MarshalJSON() (data []byte, err error) {
	type shadow PublicDividerStyleSettingsParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PublicDividerStyleSettingsParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PublicEmail struct {
	IsAb bool `json:"isAb" api:"required"`
	// The email ID.
	ID string `json:"id"`
	// The active domain of the email.
	ActiveDomain string `json:"activeDomain"`
	// List of emailCampaignIds.
	AllEmailCampaignIDs []string `json:"allEmailCampaignIds"`
	// Determines if the email is archived or not.
	Archived       bool   `json:"archived"`
	BusinessUnitID string `json:"businessUnitId"`
	// The ID of the campaign this email is associated to.
	Campaign string `json:"campaign"`
	// The name of the campaign.
	CampaignName string `json:"campaignName"`
	CampaignUtm  string `json:"campaignUtm"`
	// The ID of the email this email was cloned from.
	ClonedFrom string             `json:"clonedFrom"`
	Content    PublicEmailContent `json:"content"`
	// The date and time of the email's creation, in ISO8601 representation.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// The id of the user who created the email.
	CreatedByID string `json:"createdById"`
	// The date and time the email was deleted at, in ISO8601 representation.
	DeletedAt            time.Time `json:"deletedAt" format:"date-time"`
	EmailCampaignGroupID string    `json:"emailCampaignGroupId"`
	// Any of "DESIGN_MANAGER", "DRAG_AND_DROP".
	EmailTemplateMode PublicEmailEmailTemplateMode `json:"emailTemplateMode"`
	// The ID of the feedback survey linked to the email.
	FeedbackSurveyID string                 `json:"feedbackSurveyId"`
	FolderID         int64                  `json:"folderId"`
	FolderIDV2       int64                  `json:"folderIdV2"`
	From             PublicEmailFromDetails `json:"from"`
	// Returns the published status of the email. This is read only.
	IsPublished bool `json:"isPublished"`
	// Returns whether the email is a transactional email or not. This is read only.
	IsTransactional bool `json:"isTransactional"`
	JitterSendTime  bool `json:"jitterSendTime"`
	// Any of "aa", "ab", "ae", "af", "af-na", "af-za", "agq", "agq-cm", "ak", "ak-gh",
	// "am", "am-et", "an", "ann", "ann-ng", "ar", "ar-001", "ar-ae", "ar-bh", "ar-dj",
	// "ar-dz", "ar-eg", "ar-eh", "ar-er", "ar-il", "ar-iq", "ar-jo", "ar-km", "ar-kw",
	// "ar-lb", "ar-ly", "ar-ma", "ar-mr", "ar-om", "ar-ps", "ar-qa", "ar-sa", "ar-sd",
	// "ar-so", "ar-ss", "ar-sy", "ar-td", "ar-tn", "ar-ye", "as", "as-in", "asa",
	// "asa-tz", "ast", "ast-es", "av", "ay", "az", "az-az", "ba", "bal", "bal-pk",
	// "bas", "bas-cm", "be", "be-by", "bem", "bem-zm", "bez", "bez-tz", "bg", "bg-bg",
	// "bgc", "bgc-in", "bho", "bho-in", "bi", "blo", "blo-bj", "bm", "bm-ml", "bn",
	// "bn-bd", "bn-in", "bo", "bo-cn", "bo-in", "br", "br-fr", "brx", "brx-in", "bs",
	// "bs-ba", "ca", "ca-ad", "ca-es", "ca-fr", "ca-it", "ccp", "ccp-bd", "ccp-in",
	// "ce", "ce-ru", "ceb", "ceb-ph", "cgg", "cgg-ug", "ch", "chr", "chr-us", "ckb",
	// "ckb-iq", "ckb-ir", "co", "cr", "cs", "cs-cz", "csw", "csw-ca", "cu", "cu-ru",
	// "cv", "cv-ru", "cy", "cy-gb", "da", "da-dk", "da-gl", "dav", "dav-ke", "de",
	// "de-at", "de-be", "de-ch", "de-de", "de-gr", "de-it", "de-li", "de-lu", "dje",
	// "dje-ne", "doi", "doi-in", "dsb", "dsb-de", "dua", "dua-cm", "dv", "dyo",
	// "dyo-sn", "dz", "dz-bt", "ebu", "ebu-ke", "ee", "ee-gh", "ee-tg", "el", "el-cy",
	// "el-gr", "en", "en-001", "en-150", "en-ae", "en-ag", "en-ai", "en-as", "en-at",
	// "en-au", "en-bb", "en-be", "en-bi", "en-bm", "en-bs", "en-bw", "en-bz", "en-ca",
	// "en-cc", "en-ch", "en-ck", "en-cm", "en-cn", "en-cx", "en-cy", "en-cz", "en-de",
	// "en-dg", "en-dk", "en-dm", "en-ee", "en-eg", "en-er", "en-es", "en-fi", "en-fj",
	// "en-fk", "en-fm", "en-fr", "en-gb", "en-gd", "en-gg", "en-gh", "en-gi", "en-gm",
	// "en-gs", "en-gu", "en-gy", "en-hk", "en-hu", "en-id", "en-ie", "en-il", "en-im",
	// "en-in", "en-io", "en-it", "en-je", "en-jm", "en-ke", "en-ki", "en-kn", "en-ky",
	// "en-lc", "en-lr", "en-ls", "en-lu", "en-mg", "en-mh", "en-mo", "en-mp", "en-ms",
	// "en-mt", "en-mu", "en-mv", "en-mw", "en-mx", "en-my", "en-na", "en-nf", "en-ng",
	// "en-nl", "en-no", "en-nr", "en-nu", "en-nz", "en-pg", "en-ph", "en-pk", "en-pl",
	// "en-pn", "en-pr", "en-pt", "en-pw", "en-ro", "en-rw", "en-sb", "en-sc", "en-sd",
	// "en-se", "en-sg", "en-sh", "en-si", "en-sk", "en-sl", "en-ss", "en-sx", "en-sz",
	// "en-tc", "en-th", "en-tk", "en-tn", "en-to", "en-tt", "en-tv", "en-tz", "en-ug",
	// "en-um", "en-us", "en-vc", "en-vg", "en-vi", "en-vn", "en-vu", "en-ws", "en-za",
	// "en-zm", "en-zw", "eo", "eo-001", "es", "es-419", "es-ar", "es-bo", "es-br",
	// "es-bz", "es-cl", "es-co", "es-cr", "es-cu", "es-do", "es-ea", "es-ec", "es-es",
	// "es-gq", "es-gt", "es-hn", "es-ic", "es-mx", "es-ni", "es-pa", "es-pe", "es-ph",
	// "es-pr", "es-py", "es-sv", "es-us", "es-uy", "es-ve", "et", "et-ee", "eu",
	// "eu-es", "ewo", "ewo-cm", "fa", "fa-af", "fa-ir", "ff", "ff-bf", "ff-cm",
	// "ff-gh", "ff-gm", "ff-gn", "ff-gw", "ff-lr", "ff-mr", "ff-ne", "ff-ng", "ff-sl",
	// "ff-sn", "fi", "fi-fi", "fil", "fil-ph", "fj", "fo", "fo-dk", "fo-fo", "fr",
	// "fr-be", "fr-bf", "fr-bi", "fr-bj", "fr-bl", "fr-ca", "fr-cd", "fr-cf", "fr-cg",
	// "fr-ch", "fr-ci", "fr-cm", "fr-dj", "fr-dz", "fr-fr", "fr-ga", "fr-gf", "fr-gn",
	// "fr-gp", "fr-gq", "fr-ht", "fr-km", "fr-lu", "fr-ma", "fr-mc", "fr-mf", "fr-mg",
	// "fr-ml", "fr-mq", "fr-mr", "fr-mu", "fr-nc", "fr-ne", "fr-pf", "fr-pm", "fr-re",
	// "fr-rw", "fr-sc", "fr-sn", "fr-sy", "fr-td", "fr-tg", "fr-tn", "fr-vu", "fr-wf",
	// "fr-yt", "frr", "frr-de", "fur", "fur-it", "fy", "fy-nl", "ga", "ga-gb",
	// "ga-ie", "gaa", "gaa-gh", "gd", "gd-gb", "gl", "gl-es", "gn", "gsw", "gsw-ch",
	// "gsw-fr", "gsw-li", "gu", "gu-in", "guz", "guz-ke", "gv", "gv-im", "ha",
	// "ha-gh", "ha-ne", "ha-ng", "haw", "haw-us", "he", "he-il", "hi", "hi-in", "hmn",
	// "ho", "hr", "hr-ba", "hr-hr", "hsb", "hsb-de", "ht", "ht-ht", "hu", "hu-hu",
	// "hy", "hy-am", "hz", "ia", "ia-001", "id", "id-id", "ie", "ie-ee", "ig",
	// "ig-ng", "ii", "ii-cn", "ik", "io", "is", "is-is", "it", "it-ch", "it-it",
	// "it-sm", "it-va", "iu", "ja", "ja-jp", "jgo", "jgo-cm", "jmc", "jmc-tz", "jv",
	// "jv-id", "ka", "ka-ge", "kab", "kab-dz", "kam", "kam-ke", "kar", "kde",
	// "kde-tz", "kea", "kea-cv", "kg", "kgp", "kgp-br", "kh", "khq", "khq-ml", "ki",
	// "ki-ke", "kj", "kk", "kk-kz", "kkj", "kkj-cm", "kl", "kl-gl", "kln", "kln-ke",
	// "km", "km-kh", "kn", "kn-in", "ko", "ko-cn", "ko-kp", "ko-kr", "kok", "kok-in",
	// "kr", "ks", "ks-in", "ksb", "ksb-tz", "ksf", "ksf-cm", "ksh", "ksh-de", "ku",
	// "ku-tr", "kv", "kw", "kw-gb", "kxv", "kxv-in", "ky", "ky-kg", "la", "lag",
	// "lag-tz", "lb", "lb-lu", "lg", "lg-ug", "li", "lij", "lij-it", "lkt", "lkt-us",
	// "lmo", "lmo-it", "ln", "ln-ao", "ln-cd", "ln-cf", "ln-cg", "lo", "lo-la", "lrc",
	// "lrc-iq", "lrc-ir", "lt", "lt-lt", "lu", "lu-cd", "luo", "luo-ke", "luy",
	// "luy-ke", "lv", "lv-lv", "mai", "mai-in", "mas", "mas-ke", "mas-tz", "mdf",
	// "mdf-ru", "mer", "mer-ke", "mfe", "mfe-mu", "mg", "mg-mg", "mgh", "mgh-mz",
	// "mgo", "mgo-cm", "mh", "mi", "mi-nz", "mk", "mk-mk", "ml", "ml-in", "mn",
	// "mn-mn", "mni", "mni-in", "mr", "mr-in", "ms", "ms-bn", "ms-id", "ms-my",
	// "ms-sg", "mt", "mt-mt", "mua", "mua-cm", "my", "my-mm", "mzn", "mzn-ir", "na",
	// "naq", "naq-na", "nb", "nb-no", "nb-sj", "nd", "nd-zw", "nds", "nds-de",
	// "nds-nl", "ne", "ne-in", "ne-np", "ng", "nl", "nl-aw", "nl-be", "nl-bq",
	// "nl-ch", "nl-cw", "nl-lu", "nl-nl", "nl-sr", "nl-sx", "nmg", "nmg-cm", "nn",
	// "nn-no", "nnh", "nnh-cm", "no", "no-no", "nqo", "nqo-gn", "nr", "nso", "nso-za",
	// "nus", "nus-ss", "nv", "ny", "nyn", "nyn-ug", "oc", "oc-es", "oc-fr", "oj",
	// "om", "om-et", "om-ke", "or", "or-in", "os", "os-ge", "os-ru", "pa", "pa-in",
	// "pa-pk", "pcm", "pcm-ng", "pi", "pis", "pis-sb", "pl", "pl-pl", "prg",
	// "prg-001", "ps", "ps-af", "ps-pk", "pt", "pt-ao", "pt-br", "pt-ch", "pt-cv",
	// "pt-gq", "pt-gw", "pt-lu", "pt-mo", "pt-mz", "pt-pt", "pt-st", "pt-tl", "qu",
	// "qu-bo", "qu-ec", "qu-pe", "raj", "raj-in", "rm", "rm-ch", "rn", "rn-bi", "ro",
	// "ro-md", "ro-ro", "rof", "rof-tz", "ru", "ru-by", "ru-kg", "ru-kz", "ru-md",
	// "ru-ru", "ru-ua", "rw", "rw-rw", "rwk", "rwk-tz", "sa", "sa-in", "sah",
	// "sah-ru", "saq", "saq-ke", "sat", "sat-in", "sbp", "sbp-tz", "sc", "sc-it",
	// "sd", "sd-in", "sd-pk", "se", "se-fi", "se-no", "se-se", "seh", "seh-mz", "ses",
	// "ses-ml", "sg", "sg-cf", "shi", "shi-ma", "si", "si-lk", "sk", "sk-sk", "sl",
	// "sl-si", "sm", "smn", "smn-fi", "sms", "sms-fi", "sn", "sn-zw", "so", "so-dj",
	// "so-et", "so-ke", "so-so", "sq", "sq-al", "sq-mk", "sq-xk", "sr", "sr-ba",
	// "sr-cs", "sr-me", "sr-rs", "sr-xk", "ss", "st", "st-ls", "st-za", "su", "su-id",
	// "sv", "sv-ax", "sv-fi", "sv-se", "sw", "sw-cd", "sw-ke", "sw-tz", "sw-ug", "sy",
	// "syr", "syr-iq", "syr-sy", "szl", "szl-pl", "ta", "ta-in", "ta-lk", "ta-my",
	// "ta-sg", "te", "te-in", "teo", "teo-ke", "teo-ug", "tg", "tg-tj", "th", "th-th",
	// "ti", "ti-er", "ti-et", "tk", "tk-tm", "tl", "tn", "tn-bw", "tn-za", "to",
	// "to-to", "tok", "tok-001", "tr", "tr-cy", "tr-tr", "ts", "tt", "tt-ru", "tw",
	// "twq", "twq-ne", "ty", "tzm", "tzm-ma", "ug", "ug-cn", "uk", "uk-ua", "ur",
	// "ur-in", "ur-pk", "uz", "uz-af", "uz-uz", "vai", "vai-lr", "ve", "vec",
	// "vec-it", "vi", "vi-vn", "vmw", "vmw-mz", "vo", "vo-001", "vun", "vun-tz", "wa",
	// "wae", "wae-ch", "wo", "wo-sn", "xh", "xh-za", "xnr", "xnr-in", "xog", "xog-ug",
	// "yav", "yav-cm", "yi", "yi-001", "yi-ua", "yo", "yo-bj", "yo-ng", "yrl",
	// "yrl-br", "yrl-co", "yrl-ve", "yue", "yue-cn", "yue-hk", "yue-mo", "za",
	// "za-cn", "zgh", "zgh-ma", "zh", "zh-cn", "zh-hans", "zh-hant", "zh-hk", "zh-mo",
	// "zh-my", "zh-sg", "zh-tw", "zu", "zu-za".
	Language PublicEmailLanguage `json:"language"`
	// The name of the email, as displayed on the email dashboard.
	Name                   string `json:"name"`
	PreviewKey             string `json:"previewKey"`
	PrimaryEmailCampaignID string `json:"primaryEmailCampaignId"`
	// The date and time the email is scheduled for, in ISO8601 representation. This is
	// only used in local time or scheduled emails.
	PublishDate time.Time `json:"publishDate" format:"date-time"`
	// The date and time the email was published at, in ISO8601 representation.
	PublishedAt time.Time `json:"publishedAt" format:"date-time"`
	// Email of the user who published/sent the email.
	PublishedByEmail string `json:"publishedByEmail"`
	// The ID of the user who published the email.
	PublishedByID string `json:"publishedById"`
	// Name of the user who published the email.
	PublishedByName string                `json:"publishedByName"`
	RssData         PublicRssEmailDetails `json:"rssData"`
	// Determines whether the email will be sent immediately on publish.
	SendOnPublish bool `json:"sendOnPublish"`
	// The email state.
	//
	// Any of "AGENT_GENERATED", "AUTOMATED", "AUTOMATED_AB", "AUTOMATED_AB_VARIANT",
	// "AUTOMATED_DRAFT", "AUTOMATED_DRAFT_AB", "AUTOMATED_DRAFT_ABVARIANT",
	// "AUTOMATED_FOR_FORM", "AUTOMATED_FOR_FORM_BUFFER", "AUTOMATED_FOR_FORM_DRAFT",
	// "AUTOMATED_FOR_FORM_LEGACY", "AUTOMATED_LOSER_ABVARIANT", "AUTOMATED_SENDING",
	// "BLOG_EMAIL_DRAFT", "BLOG_EMAIL_PUBLISHED", "DRAFT", "DRAFT_AB",
	// "DRAFT_AB_VARIANT", "ERROR", "LOSER_AB_VARIANT", "PAGE_STUB", "PRE_PROCESSING",
	// "PROCESSING", "PUBLISHED", "PUBLISHED_AB", "PUBLISHED_AB_VARIANT",
	// "PUBLISHED_OR_SCHEDULED", "RSS_TO_EMAIL_DRAFT", "RSS_TO_EMAIL_PUBLISHED",
	// "SCHEDULED", "SCHEDULED_AB", "SCHEDULED_OR_PUBLISHED".
	State PublicEmailState    `json:"state"`
	Stats EmailStatisticsData `json:"stats"`
	// The email subcategory.
	Subcategory string `json:"subcategory"`
	// The subject of the email.
	Subject             string                         `json:"subject"`
	SubscriptionDetails PublicEmailSubscriptionDetails `json:"subscriptionDetails"`
	TeamsWithAccess     []string                       `json:"teamsWithAccess"`
	Testing             PublicEmailTestingDetails      `json:"testing"`
	To                  PublicEmailToDetails           `json:"to"`
	// The email type, this is derived from other properties on the email such as
	// subcategory.
	//
	// Any of "AB_EMAIL", "AUTOMATED_AB_EMAIL", "AUTOMATED_EMAIL", "BATCH_EMAIL",
	// "BLOG_EMAIL", "BLOG_EMAIL_CHILD", "FEEDBACK_CES_EMAIL", "FEEDBACK_CUSTOM_EMAIL",
	// "FEEDBACK_CUSTOM_SURVEY_EMAIL", "FEEDBACK_NPS_EMAIL", "FOLLOWUP_EMAIL",
	// "LEADFLOW_EMAIL", "LOCALTIME_EMAIL", "MANAGE_PREFERENCES_EMAIL",
	// "MARKETING_SINGLE_SEND_API", "MEMBERSHIP_EMAIL_VERIFICATION_EMAIL",
	// "MEMBERSHIP_FOLLOW_UP_EMAIL", "MEMBERSHIP_OTP_LOGIN_EMAIL",
	// "MEMBERSHIP_PASSWORD_RESET_EMAIL", "MEMBERSHIP_PASSWORD_SAVED_EMAIL",
	// "MEMBERSHIP_PASSWORDLESS_AUTH_EMAIL", "MEMBERSHIP_REGISTRATION_EMAIL",
	// "MEMBERSHIP_REGISTRATION_FOLLOW_UP_EMAIL", "MEMBERSHIP_VERIFICATION_EMAIL",
	// "OPTIN_EMAIL", "OPTIN_FOLLOWUP_EMAIL", "RESUBSCRIBE_EMAIL", "RSS_EMAIL",
	// "RSS_EMAIL_CHILD", "SINGLE_SEND_API", "SMTP_TOKEN", "TICKET_EMAIL".
	Type          PublicEmailType `json:"type"`
	UnpublishedAt time.Time       `json:"unpublishedAt" format:"date-time"`
	// The date and time of the last update to the email, in ISO8601 representation.
	UpdatedAt time.Time `json:"updatedAt" format:"date-time"`
	// The id of the user who last updated the email.
	UpdatedByID     string                  `json:"updatedById"`
	UsersWithAccess []string                `json:"usersWithAccess"`
	Webversion      PublicWebversionDetails `json:"webversion"`
	// Names of workflows in which the email is used within a "send email" action.
	WorkflowNames []string `json:"workflowNames"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		IsAb                   respjson.Field
		ID                     respjson.Field
		ActiveDomain           respjson.Field
		AllEmailCampaignIDs    respjson.Field
		Archived               respjson.Field
		BusinessUnitID         respjson.Field
		Campaign               respjson.Field
		CampaignName           respjson.Field
		CampaignUtm            respjson.Field
		ClonedFrom             respjson.Field
		Content                respjson.Field
		CreatedAt              respjson.Field
		CreatedByID            respjson.Field
		DeletedAt              respjson.Field
		EmailCampaignGroupID   respjson.Field
		EmailTemplateMode      respjson.Field
		FeedbackSurveyID       respjson.Field
		FolderID               respjson.Field
		FolderIDV2             respjson.Field
		From                   respjson.Field
		IsPublished            respjson.Field
		IsTransactional        respjson.Field
		JitterSendTime         respjson.Field
		Language               respjson.Field
		Name                   respjson.Field
		PreviewKey             respjson.Field
		PrimaryEmailCampaignID respjson.Field
		PublishDate            respjson.Field
		PublishedAt            respjson.Field
		PublishedByEmail       respjson.Field
		PublishedByID          respjson.Field
		PublishedByName        respjson.Field
		RssData                respjson.Field
		SendOnPublish          respjson.Field
		State                  respjson.Field
		Stats                  respjson.Field
		Subcategory            respjson.Field
		Subject                respjson.Field
		SubscriptionDetails    respjson.Field
		TeamsWithAccess        respjson.Field
		Testing                respjson.Field
		To                     respjson.Field
		Type                   respjson.Field
		UnpublishedAt          respjson.Field
		UpdatedAt              respjson.Field
		UpdatedByID            respjson.Field
		UsersWithAccess        respjson.Field
		Webversion             respjson.Field
		WorkflowNames          respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PublicEmail) RawJSON() string { return r.JSON.raw }
func (r *PublicEmail) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PublicEmailEmailTemplateMode string

const (
	PublicEmailEmailTemplateModeDesignManager PublicEmailEmailTemplateMode = "DESIGN_MANAGER"
	PublicEmailEmailTemplateModeDragAndDrop   PublicEmailEmailTemplateMode = "DRAG_AND_DROP"
)

type PublicEmailLanguage string

const (
	PublicEmailLanguageAa     PublicEmailLanguage = "aa"
	PublicEmailLanguageAb     PublicEmailLanguage = "ab"
	PublicEmailLanguageAe     PublicEmailLanguage = "ae"
	PublicEmailLanguageAf     PublicEmailLanguage = "af"
	PublicEmailLanguageAfNa   PublicEmailLanguage = "af-na"
	PublicEmailLanguageAfZa   PublicEmailLanguage = "af-za"
	PublicEmailLanguageAgq    PublicEmailLanguage = "agq"
	PublicEmailLanguageAgqCm  PublicEmailLanguage = "agq-cm"
	PublicEmailLanguageAk     PublicEmailLanguage = "ak"
	PublicEmailLanguageAkGh   PublicEmailLanguage = "ak-gh"
	PublicEmailLanguageAm     PublicEmailLanguage = "am"
	PublicEmailLanguageAmEt   PublicEmailLanguage = "am-et"
	PublicEmailLanguageAn     PublicEmailLanguage = "an"
	PublicEmailLanguageAnn    PublicEmailLanguage = "ann"
	PublicEmailLanguageAnnNg  PublicEmailLanguage = "ann-ng"
	PublicEmailLanguageAr     PublicEmailLanguage = "ar"
	PublicEmailLanguageAr001  PublicEmailLanguage = "ar-001"
	PublicEmailLanguageArAe   PublicEmailLanguage = "ar-ae"
	PublicEmailLanguageArBh   PublicEmailLanguage = "ar-bh"
	PublicEmailLanguageArDj   PublicEmailLanguage = "ar-dj"
	PublicEmailLanguageArDz   PublicEmailLanguage = "ar-dz"
	PublicEmailLanguageArEg   PublicEmailLanguage = "ar-eg"
	PublicEmailLanguageArEh   PublicEmailLanguage = "ar-eh"
	PublicEmailLanguageArEr   PublicEmailLanguage = "ar-er"
	PublicEmailLanguageArIl   PublicEmailLanguage = "ar-il"
	PublicEmailLanguageArIq   PublicEmailLanguage = "ar-iq"
	PublicEmailLanguageArJo   PublicEmailLanguage = "ar-jo"
	PublicEmailLanguageArKm   PublicEmailLanguage = "ar-km"
	PublicEmailLanguageArKw   PublicEmailLanguage = "ar-kw"
	PublicEmailLanguageArLb   PublicEmailLanguage = "ar-lb"
	PublicEmailLanguageArLy   PublicEmailLanguage = "ar-ly"
	PublicEmailLanguageArMa   PublicEmailLanguage = "ar-ma"
	PublicEmailLanguageArMr   PublicEmailLanguage = "ar-mr"
	PublicEmailLanguageArOm   PublicEmailLanguage = "ar-om"
	PublicEmailLanguageArPs   PublicEmailLanguage = "ar-ps"
	PublicEmailLanguageArQa   PublicEmailLanguage = "ar-qa"
	PublicEmailLanguageArSa   PublicEmailLanguage = "ar-sa"
	PublicEmailLanguageArSd   PublicEmailLanguage = "ar-sd"
	PublicEmailLanguageArSo   PublicEmailLanguage = "ar-so"
	PublicEmailLanguageArSS   PublicEmailLanguage = "ar-ss"
	PublicEmailLanguageArSy   PublicEmailLanguage = "ar-sy"
	PublicEmailLanguageArTd   PublicEmailLanguage = "ar-td"
	PublicEmailLanguageArTn   PublicEmailLanguage = "ar-tn"
	PublicEmailLanguageArYe   PublicEmailLanguage = "ar-ye"
	PublicEmailLanguageAs     PublicEmailLanguage = "as"
	PublicEmailLanguageAsIn   PublicEmailLanguage = "as-in"
	PublicEmailLanguageAsa    PublicEmailLanguage = "asa"
	PublicEmailLanguageAsaTz  PublicEmailLanguage = "asa-tz"
	PublicEmailLanguageAst    PublicEmailLanguage = "ast"
	PublicEmailLanguageAstEs  PublicEmailLanguage = "ast-es"
	PublicEmailLanguageAv     PublicEmailLanguage = "av"
	PublicEmailLanguageAy     PublicEmailLanguage = "ay"
	PublicEmailLanguageAz     PublicEmailLanguage = "az"
	PublicEmailLanguageAzAz   PublicEmailLanguage = "az-az"
	PublicEmailLanguageBa     PublicEmailLanguage = "ba"
	PublicEmailLanguageBal    PublicEmailLanguage = "bal"
	PublicEmailLanguageBalPk  PublicEmailLanguage = "bal-pk"
	PublicEmailLanguageBas    PublicEmailLanguage = "bas"
	PublicEmailLanguageBasCm  PublicEmailLanguage = "bas-cm"
	PublicEmailLanguageBe     PublicEmailLanguage = "be"
	PublicEmailLanguageBeBy   PublicEmailLanguage = "be-by"
	PublicEmailLanguageBem    PublicEmailLanguage = "bem"
	PublicEmailLanguageBemZm  PublicEmailLanguage = "bem-zm"
	PublicEmailLanguageBez    PublicEmailLanguage = "bez"
	PublicEmailLanguageBezTz  PublicEmailLanguage = "bez-tz"
	PublicEmailLanguageBg     PublicEmailLanguage = "bg"
	PublicEmailLanguageBgBg   PublicEmailLanguage = "bg-bg"
	PublicEmailLanguageBgc    PublicEmailLanguage = "bgc"
	PublicEmailLanguageBgcIn  PublicEmailLanguage = "bgc-in"
	PublicEmailLanguageBho    PublicEmailLanguage = "bho"
	PublicEmailLanguageBhoIn  PublicEmailLanguage = "bho-in"
	PublicEmailLanguageBi     PublicEmailLanguage = "bi"
	PublicEmailLanguageBlo    PublicEmailLanguage = "blo"
	PublicEmailLanguageBloBj  PublicEmailLanguage = "blo-bj"
	PublicEmailLanguageBm     PublicEmailLanguage = "bm"
	PublicEmailLanguageBmMl   PublicEmailLanguage = "bm-ml"
	PublicEmailLanguageBn     PublicEmailLanguage = "bn"
	PublicEmailLanguageBnBd   PublicEmailLanguage = "bn-bd"
	PublicEmailLanguageBnIn   PublicEmailLanguage = "bn-in"
	PublicEmailLanguageBo     PublicEmailLanguage = "bo"
	PublicEmailLanguageBoCn   PublicEmailLanguage = "bo-cn"
	PublicEmailLanguageBoIn   PublicEmailLanguage = "bo-in"
	PublicEmailLanguageBr     PublicEmailLanguage = "br"
	PublicEmailLanguageBrFr   PublicEmailLanguage = "br-fr"
	PublicEmailLanguageBrx    PublicEmailLanguage = "brx"
	PublicEmailLanguageBrxIn  PublicEmailLanguage = "brx-in"
	PublicEmailLanguageBs     PublicEmailLanguage = "bs"
	PublicEmailLanguageBsBa   PublicEmailLanguage = "bs-ba"
	PublicEmailLanguageCa     PublicEmailLanguage = "ca"
	PublicEmailLanguageCaAd   PublicEmailLanguage = "ca-ad"
	PublicEmailLanguageCaEs   PublicEmailLanguage = "ca-es"
	PublicEmailLanguageCaFr   PublicEmailLanguage = "ca-fr"
	PublicEmailLanguageCaIt   PublicEmailLanguage = "ca-it"
	PublicEmailLanguageCcp    PublicEmailLanguage = "ccp"
	PublicEmailLanguageCcpBd  PublicEmailLanguage = "ccp-bd"
	PublicEmailLanguageCcpIn  PublicEmailLanguage = "ccp-in"
	PublicEmailLanguageCe     PublicEmailLanguage = "ce"
	PublicEmailLanguageCeRu   PublicEmailLanguage = "ce-ru"
	PublicEmailLanguageCeb    PublicEmailLanguage = "ceb"
	PublicEmailLanguageCebPh  PublicEmailLanguage = "ceb-ph"
	PublicEmailLanguageCgg    PublicEmailLanguage = "cgg"
	PublicEmailLanguageCggUg  PublicEmailLanguage = "cgg-ug"
	PublicEmailLanguageCh     PublicEmailLanguage = "ch"
	PublicEmailLanguageChr    PublicEmailLanguage = "chr"
	PublicEmailLanguageChrUs  PublicEmailLanguage = "chr-us"
	PublicEmailLanguageCkb    PublicEmailLanguage = "ckb"
	PublicEmailLanguageCkbIq  PublicEmailLanguage = "ckb-iq"
	PublicEmailLanguageCkbIr  PublicEmailLanguage = "ckb-ir"
	PublicEmailLanguageCo     PublicEmailLanguage = "co"
	PublicEmailLanguageCr     PublicEmailLanguage = "cr"
	PublicEmailLanguageCs     PublicEmailLanguage = "cs"
	PublicEmailLanguageCsCz   PublicEmailLanguage = "cs-cz"
	PublicEmailLanguageCsw    PublicEmailLanguage = "csw"
	PublicEmailLanguageCswCa  PublicEmailLanguage = "csw-ca"
	PublicEmailLanguageCu     PublicEmailLanguage = "cu"
	PublicEmailLanguageCuRu   PublicEmailLanguage = "cu-ru"
	PublicEmailLanguageCv     PublicEmailLanguage = "cv"
	PublicEmailLanguageCvRu   PublicEmailLanguage = "cv-ru"
	PublicEmailLanguageCy     PublicEmailLanguage = "cy"
	PublicEmailLanguageCyGB   PublicEmailLanguage = "cy-gb"
	PublicEmailLanguageDa     PublicEmailLanguage = "da"
	PublicEmailLanguageDaDk   PublicEmailLanguage = "da-dk"
	PublicEmailLanguageDaGl   PublicEmailLanguage = "da-gl"
	PublicEmailLanguageDav    PublicEmailLanguage = "dav"
	PublicEmailLanguageDavKe  PublicEmailLanguage = "dav-ke"
	PublicEmailLanguageDe     PublicEmailLanguage = "de"
	PublicEmailLanguageDeAt   PublicEmailLanguage = "de-at"
	PublicEmailLanguageDeBe   PublicEmailLanguage = "de-be"
	PublicEmailLanguageDeCh   PublicEmailLanguage = "de-ch"
	PublicEmailLanguageDeDe   PublicEmailLanguage = "de-de"
	PublicEmailLanguageDeGr   PublicEmailLanguage = "de-gr"
	PublicEmailLanguageDeIt   PublicEmailLanguage = "de-it"
	PublicEmailLanguageDeLi   PublicEmailLanguage = "de-li"
	PublicEmailLanguageDeLu   PublicEmailLanguage = "de-lu"
	PublicEmailLanguageDje    PublicEmailLanguage = "dje"
	PublicEmailLanguageDjeNe  PublicEmailLanguage = "dje-ne"
	PublicEmailLanguageDoi    PublicEmailLanguage = "doi"
	PublicEmailLanguageDoiIn  PublicEmailLanguage = "doi-in"
	PublicEmailLanguageDsb    PublicEmailLanguage = "dsb"
	PublicEmailLanguageDsbDe  PublicEmailLanguage = "dsb-de"
	PublicEmailLanguageDua    PublicEmailLanguage = "dua"
	PublicEmailLanguageDuaCm  PublicEmailLanguage = "dua-cm"
	PublicEmailLanguageDv     PublicEmailLanguage = "dv"
	PublicEmailLanguageDyo    PublicEmailLanguage = "dyo"
	PublicEmailLanguageDyoSn  PublicEmailLanguage = "dyo-sn"
	PublicEmailLanguageDz     PublicEmailLanguage = "dz"
	PublicEmailLanguageDzBt   PublicEmailLanguage = "dz-bt"
	PublicEmailLanguageEbu    PublicEmailLanguage = "ebu"
	PublicEmailLanguageEbuKe  PublicEmailLanguage = "ebu-ke"
	PublicEmailLanguageEe     PublicEmailLanguage = "ee"
	PublicEmailLanguageEeGh   PublicEmailLanguage = "ee-gh"
	PublicEmailLanguageEeTg   PublicEmailLanguage = "ee-tg"
	PublicEmailLanguageEl     PublicEmailLanguage = "el"
	PublicEmailLanguageElCy   PublicEmailLanguage = "el-cy"
	PublicEmailLanguageElGr   PublicEmailLanguage = "el-gr"
	PublicEmailLanguageEn     PublicEmailLanguage = "en"
	PublicEmailLanguageEn001  PublicEmailLanguage = "en-001"
	PublicEmailLanguageEn150  PublicEmailLanguage = "en-150"
	PublicEmailLanguageEnAe   PublicEmailLanguage = "en-ae"
	PublicEmailLanguageEnAg   PublicEmailLanguage = "en-ag"
	PublicEmailLanguageEnAI   PublicEmailLanguage = "en-ai"
	PublicEmailLanguageEnAs   PublicEmailLanguage = "en-as"
	PublicEmailLanguageEnAt   PublicEmailLanguage = "en-at"
	PublicEmailLanguageEnAu   PublicEmailLanguage = "en-au"
	PublicEmailLanguageEnBb   PublicEmailLanguage = "en-bb"
	PublicEmailLanguageEnBe   PublicEmailLanguage = "en-be"
	PublicEmailLanguageEnBi   PublicEmailLanguage = "en-bi"
	PublicEmailLanguageEnBm   PublicEmailLanguage = "en-bm"
	PublicEmailLanguageEnBs   PublicEmailLanguage = "en-bs"
	PublicEmailLanguageEnBw   PublicEmailLanguage = "en-bw"
	PublicEmailLanguageEnBz   PublicEmailLanguage = "en-bz"
	PublicEmailLanguageEnCa   PublicEmailLanguage = "en-ca"
	PublicEmailLanguageEnCc   PublicEmailLanguage = "en-cc"
	PublicEmailLanguageEnCh   PublicEmailLanguage = "en-ch"
	PublicEmailLanguageEnCk   PublicEmailLanguage = "en-ck"
	PublicEmailLanguageEnCm   PublicEmailLanguage = "en-cm"
	PublicEmailLanguageEnCn   PublicEmailLanguage = "en-cn"
	PublicEmailLanguageEnCx   PublicEmailLanguage = "en-cx"
	PublicEmailLanguageEnCy   PublicEmailLanguage = "en-cy"
	PublicEmailLanguageEnCz   PublicEmailLanguage = "en-cz"
	PublicEmailLanguageEnDe   PublicEmailLanguage = "en-de"
	PublicEmailLanguageEnDg   PublicEmailLanguage = "en-dg"
	PublicEmailLanguageEnDk   PublicEmailLanguage = "en-dk"
	PublicEmailLanguageEnDm   PublicEmailLanguage = "en-dm"
	PublicEmailLanguageEnEe   PublicEmailLanguage = "en-ee"
	PublicEmailLanguageEnEg   PublicEmailLanguage = "en-eg"
	PublicEmailLanguageEnEr   PublicEmailLanguage = "en-er"
	PublicEmailLanguageEnEs   PublicEmailLanguage = "en-es"
	PublicEmailLanguageEnFi   PublicEmailLanguage = "en-fi"
	PublicEmailLanguageEnFj   PublicEmailLanguage = "en-fj"
	PublicEmailLanguageEnFk   PublicEmailLanguage = "en-fk"
	PublicEmailLanguageEnFm   PublicEmailLanguage = "en-fm"
	PublicEmailLanguageEnFr   PublicEmailLanguage = "en-fr"
	PublicEmailLanguageEnGB   PublicEmailLanguage = "en-gb"
	PublicEmailLanguageEnGd   PublicEmailLanguage = "en-gd"
	PublicEmailLanguageEnGg   PublicEmailLanguage = "en-gg"
	PublicEmailLanguageEnGh   PublicEmailLanguage = "en-gh"
	PublicEmailLanguageEnGi   PublicEmailLanguage = "en-gi"
	PublicEmailLanguageEnGm   PublicEmailLanguage = "en-gm"
	PublicEmailLanguageEnGs   PublicEmailLanguage = "en-gs"
	PublicEmailLanguageEnGu   PublicEmailLanguage = "en-gu"
	PublicEmailLanguageEnGy   PublicEmailLanguage = "en-gy"
	PublicEmailLanguageEnHk   PublicEmailLanguage = "en-hk"
	PublicEmailLanguageEnHu   PublicEmailLanguage = "en-hu"
	PublicEmailLanguageEnID   PublicEmailLanguage = "en-id"
	PublicEmailLanguageEnIe   PublicEmailLanguage = "en-ie"
	PublicEmailLanguageEnIl   PublicEmailLanguage = "en-il"
	PublicEmailLanguageEnIm   PublicEmailLanguage = "en-im"
	PublicEmailLanguageEnIn   PublicEmailLanguage = "en-in"
	PublicEmailLanguageEnIo   PublicEmailLanguage = "en-io"
	PublicEmailLanguageEnIt   PublicEmailLanguage = "en-it"
	PublicEmailLanguageEnJe   PublicEmailLanguage = "en-je"
	PublicEmailLanguageEnJm   PublicEmailLanguage = "en-jm"
	PublicEmailLanguageEnKe   PublicEmailLanguage = "en-ke"
	PublicEmailLanguageEnKi   PublicEmailLanguage = "en-ki"
	PublicEmailLanguageEnKn   PublicEmailLanguage = "en-kn"
	PublicEmailLanguageEnKy   PublicEmailLanguage = "en-ky"
	PublicEmailLanguageEnLc   PublicEmailLanguage = "en-lc"
	PublicEmailLanguageEnLr   PublicEmailLanguage = "en-lr"
	PublicEmailLanguageEnLs   PublicEmailLanguage = "en-ls"
	PublicEmailLanguageEnLu   PublicEmailLanguage = "en-lu"
	PublicEmailLanguageEnMg   PublicEmailLanguage = "en-mg"
	PublicEmailLanguageEnMh   PublicEmailLanguage = "en-mh"
	PublicEmailLanguageEnMo   PublicEmailLanguage = "en-mo"
	PublicEmailLanguageEnMp   PublicEmailLanguage = "en-mp"
	PublicEmailLanguageEnMs   PublicEmailLanguage = "en-ms"
	PublicEmailLanguageEnMt   PublicEmailLanguage = "en-mt"
	PublicEmailLanguageEnMu   PublicEmailLanguage = "en-mu"
	PublicEmailLanguageEnMv   PublicEmailLanguage = "en-mv"
	PublicEmailLanguageEnMw   PublicEmailLanguage = "en-mw"
	PublicEmailLanguageEnMx   PublicEmailLanguage = "en-mx"
	PublicEmailLanguageEnMy   PublicEmailLanguage = "en-my"
	PublicEmailLanguageEnNa   PublicEmailLanguage = "en-na"
	PublicEmailLanguageEnNf   PublicEmailLanguage = "en-nf"
	PublicEmailLanguageEnNg   PublicEmailLanguage = "en-ng"
	PublicEmailLanguageEnNl   PublicEmailLanguage = "en-nl"
	PublicEmailLanguageEnNo   PublicEmailLanguage = "en-no"
	PublicEmailLanguageEnNr   PublicEmailLanguage = "en-nr"
	PublicEmailLanguageEnNu   PublicEmailLanguage = "en-nu"
	PublicEmailLanguageEnNz   PublicEmailLanguage = "en-nz"
	PublicEmailLanguageEnPg   PublicEmailLanguage = "en-pg"
	PublicEmailLanguageEnPh   PublicEmailLanguage = "en-ph"
	PublicEmailLanguageEnPk   PublicEmailLanguage = "en-pk"
	PublicEmailLanguageEnPl   PublicEmailLanguage = "en-pl"
	PublicEmailLanguageEnPn   PublicEmailLanguage = "en-pn"
	PublicEmailLanguageEnPr   PublicEmailLanguage = "en-pr"
	PublicEmailLanguageEnPt   PublicEmailLanguage = "en-pt"
	PublicEmailLanguageEnPw   PublicEmailLanguage = "en-pw"
	PublicEmailLanguageEnRo   PublicEmailLanguage = "en-ro"
	PublicEmailLanguageEnRw   PublicEmailLanguage = "en-rw"
	PublicEmailLanguageEnSb   PublicEmailLanguage = "en-sb"
	PublicEmailLanguageEnSc   PublicEmailLanguage = "en-sc"
	PublicEmailLanguageEnSd   PublicEmailLanguage = "en-sd"
	PublicEmailLanguageEnSe   PublicEmailLanguage = "en-se"
	PublicEmailLanguageEnSg   PublicEmailLanguage = "en-sg"
	PublicEmailLanguageEnSh   PublicEmailLanguage = "en-sh"
	PublicEmailLanguageEnSi   PublicEmailLanguage = "en-si"
	PublicEmailLanguageEnSk   PublicEmailLanguage = "en-sk"
	PublicEmailLanguageEnSl   PublicEmailLanguage = "en-sl"
	PublicEmailLanguageEnSS   PublicEmailLanguage = "en-ss"
	PublicEmailLanguageEnSx   PublicEmailLanguage = "en-sx"
	PublicEmailLanguageEnSz   PublicEmailLanguage = "en-sz"
	PublicEmailLanguageEnTc   PublicEmailLanguage = "en-tc"
	PublicEmailLanguageEnTh   PublicEmailLanguage = "en-th"
	PublicEmailLanguageEnTk   PublicEmailLanguage = "en-tk"
	PublicEmailLanguageEnTn   PublicEmailLanguage = "en-tn"
	PublicEmailLanguageEnTo   PublicEmailLanguage = "en-to"
	PublicEmailLanguageEnTt   PublicEmailLanguage = "en-tt"
	PublicEmailLanguageEnTv   PublicEmailLanguage = "en-tv"
	PublicEmailLanguageEnTz   PublicEmailLanguage = "en-tz"
	PublicEmailLanguageEnUg   PublicEmailLanguage = "en-ug"
	PublicEmailLanguageEnUm   PublicEmailLanguage = "en-um"
	PublicEmailLanguageEnUs   PublicEmailLanguage = "en-us"
	PublicEmailLanguageEnVc   PublicEmailLanguage = "en-vc"
	PublicEmailLanguageEnVg   PublicEmailLanguage = "en-vg"
	PublicEmailLanguageEnVi   PublicEmailLanguage = "en-vi"
	PublicEmailLanguageEnVn   PublicEmailLanguage = "en-vn"
	PublicEmailLanguageEnVu   PublicEmailLanguage = "en-vu"
	PublicEmailLanguageEnWs   PublicEmailLanguage = "en-ws"
	PublicEmailLanguageEnZa   PublicEmailLanguage = "en-za"
	PublicEmailLanguageEnZm   PublicEmailLanguage = "en-zm"
	PublicEmailLanguageEnZw   PublicEmailLanguage = "en-zw"
	PublicEmailLanguageEo     PublicEmailLanguage = "eo"
	PublicEmailLanguageEo001  PublicEmailLanguage = "eo-001"
	PublicEmailLanguageEs     PublicEmailLanguage = "es"
	PublicEmailLanguageEs419  PublicEmailLanguage = "es-419"
	PublicEmailLanguageEsAr   PublicEmailLanguage = "es-ar"
	PublicEmailLanguageEsBo   PublicEmailLanguage = "es-bo"
	PublicEmailLanguageEsBr   PublicEmailLanguage = "es-br"
	PublicEmailLanguageEsBz   PublicEmailLanguage = "es-bz"
	PublicEmailLanguageEsCl   PublicEmailLanguage = "es-cl"
	PublicEmailLanguageEsCo   PublicEmailLanguage = "es-co"
	PublicEmailLanguageEsCr   PublicEmailLanguage = "es-cr"
	PublicEmailLanguageEsCu   PublicEmailLanguage = "es-cu"
	PublicEmailLanguageEsDo   PublicEmailLanguage = "es-do"
	PublicEmailLanguageEsEa   PublicEmailLanguage = "es-ea"
	PublicEmailLanguageEsEc   PublicEmailLanguage = "es-ec"
	PublicEmailLanguageEsEs   PublicEmailLanguage = "es-es"
	PublicEmailLanguageEsGq   PublicEmailLanguage = "es-gq"
	PublicEmailLanguageEsGt   PublicEmailLanguage = "es-gt"
	PublicEmailLanguageEsHn   PublicEmailLanguage = "es-hn"
	PublicEmailLanguageEsIc   PublicEmailLanguage = "es-ic"
	PublicEmailLanguageEsMx   PublicEmailLanguage = "es-mx"
	PublicEmailLanguageEsNi   PublicEmailLanguage = "es-ni"
	PublicEmailLanguageEsPa   PublicEmailLanguage = "es-pa"
	PublicEmailLanguageEsPe   PublicEmailLanguage = "es-pe"
	PublicEmailLanguageEsPh   PublicEmailLanguage = "es-ph"
	PublicEmailLanguageEsPr   PublicEmailLanguage = "es-pr"
	PublicEmailLanguageEsPy   PublicEmailLanguage = "es-py"
	PublicEmailLanguageEsSv   PublicEmailLanguage = "es-sv"
	PublicEmailLanguageEsUs   PublicEmailLanguage = "es-us"
	PublicEmailLanguageEsUy   PublicEmailLanguage = "es-uy"
	PublicEmailLanguageEsVe   PublicEmailLanguage = "es-ve"
	PublicEmailLanguageEt     PublicEmailLanguage = "et"
	PublicEmailLanguageEtEe   PublicEmailLanguage = "et-ee"
	PublicEmailLanguageEu     PublicEmailLanguage = "eu"
	PublicEmailLanguageEuEs   PublicEmailLanguage = "eu-es"
	PublicEmailLanguageEwo    PublicEmailLanguage = "ewo"
	PublicEmailLanguageEwoCm  PublicEmailLanguage = "ewo-cm"
	PublicEmailLanguageFa     PublicEmailLanguage = "fa"
	PublicEmailLanguageFaAf   PublicEmailLanguage = "fa-af"
	PublicEmailLanguageFaIr   PublicEmailLanguage = "fa-ir"
	PublicEmailLanguageFf     PublicEmailLanguage = "ff"
	PublicEmailLanguageFfBf   PublicEmailLanguage = "ff-bf"
	PublicEmailLanguageFfCm   PublicEmailLanguage = "ff-cm"
	PublicEmailLanguageFfGh   PublicEmailLanguage = "ff-gh"
	PublicEmailLanguageFfGm   PublicEmailLanguage = "ff-gm"
	PublicEmailLanguageFfGn   PublicEmailLanguage = "ff-gn"
	PublicEmailLanguageFfGw   PublicEmailLanguage = "ff-gw"
	PublicEmailLanguageFfLr   PublicEmailLanguage = "ff-lr"
	PublicEmailLanguageFfMr   PublicEmailLanguage = "ff-mr"
	PublicEmailLanguageFfNe   PublicEmailLanguage = "ff-ne"
	PublicEmailLanguageFfNg   PublicEmailLanguage = "ff-ng"
	PublicEmailLanguageFfSl   PublicEmailLanguage = "ff-sl"
	PublicEmailLanguageFfSn   PublicEmailLanguage = "ff-sn"
	PublicEmailLanguageFi     PublicEmailLanguage = "fi"
	PublicEmailLanguageFiFi   PublicEmailLanguage = "fi-fi"
	PublicEmailLanguageFil    PublicEmailLanguage = "fil"
	PublicEmailLanguageFilPh  PublicEmailLanguage = "fil-ph"
	PublicEmailLanguageFj     PublicEmailLanguage = "fj"
	PublicEmailLanguageFo     PublicEmailLanguage = "fo"
	PublicEmailLanguageFoDk   PublicEmailLanguage = "fo-dk"
	PublicEmailLanguageFoFo   PublicEmailLanguage = "fo-fo"
	PublicEmailLanguageFr     PublicEmailLanguage = "fr"
	PublicEmailLanguageFrBe   PublicEmailLanguage = "fr-be"
	PublicEmailLanguageFrBf   PublicEmailLanguage = "fr-bf"
	PublicEmailLanguageFrBi   PublicEmailLanguage = "fr-bi"
	PublicEmailLanguageFrBj   PublicEmailLanguage = "fr-bj"
	PublicEmailLanguageFrBl   PublicEmailLanguage = "fr-bl"
	PublicEmailLanguageFrCa   PublicEmailLanguage = "fr-ca"
	PublicEmailLanguageFrCd   PublicEmailLanguage = "fr-cd"
	PublicEmailLanguageFrCf   PublicEmailLanguage = "fr-cf"
	PublicEmailLanguageFrCg   PublicEmailLanguage = "fr-cg"
	PublicEmailLanguageFrCh   PublicEmailLanguage = "fr-ch"
	PublicEmailLanguageFrCi   PublicEmailLanguage = "fr-ci"
	PublicEmailLanguageFrCm   PublicEmailLanguage = "fr-cm"
	PublicEmailLanguageFrDj   PublicEmailLanguage = "fr-dj"
	PublicEmailLanguageFrDz   PublicEmailLanguage = "fr-dz"
	PublicEmailLanguageFrFr   PublicEmailLanguage = "fr-fr"
	PublicEmailLanguageFrGa   PublicEmailLanguage = "fr-ga"
	PublicEmailLanguageFrGf   PublicEmailLanguage = "fr-gf"
	PublicEmailLanguageFrGn   PublicEmailLanguage = "fr-gn"
	PublicEmailLanguageFrGp   PublicEmailLanguage = "fr-gp"
	PublicEmailLanguageFrGq   PublicEmailLanguage = "fr-gq"
	PublicEmailLanguageFrHt   PublicEmailLanguage = "fr-ht"
	PublicEmailLanguageFrKm   PublicEmailLanguage = "fr-km"
	PublicEmailLanguageFrLu   PublicEmailLanguage = "fr-lu"
	PublicEmailLanguageFrMa   PublicEmailLanguage = "fr-ma"
	PublicEmailLanguageFrMc   PublicEmailLanguage = "fr-mc"
	PublicEmailLanguageFrMf   PublicEmailLanguage = "fr-mf"
	PublicEmailLanguageFrMg   PublicEmailLanguage = "fr-mg"
	PublicEmailLanguageFrMl   PublicEmailLanguage = "fr-ml"
	PublicEmailLanguageFrMq   PublicEmailLanguage = "fr-mq"
	PublicEmailLanguageFrMr   PublicEmailLanguage = "fr-mr"
	PublicEmailLanguageFrMu   PublicEmailLanguage = "fr-mu"
	PublicEmailLanguageFrNc   PublicEmailLanguage = "fr-nc"
	PublicEmailLanguageFrNe   PublicEmailLanguage = "fr-ne"
	PublicEmailLanguageFrPf   PublicEmailLanguage = "fr-pf"
	PublicEmailLanguageFrPm   PublicEmailLanguage = "fr-pm"
	PublicEmailLanguageFrRe   PublicEmailLanguage = "fr-re"
	PublicEmailLanguageFrRw   PublicEmailLanguage = "fr-rw"
	PublicEmailLanguageFrSc   PublicEmailLanguage = "fr-sc"
	PublicEmailLanguageFrSn   PublicEmailLanguage = "fr-sn"
	PublicEmailLanguageFrSy   PublicEmailLanguage = "fr-sy"
	PublicEmailLanguageFrTd   PublicEmailLanguage = "fr-td"
	PublicEmailLanguageFrTg   PublicEmailLanguage = "fr-tg"
	PublicEmailLanguageFrTn   PublicEmailLanguage = "fr-tn"
	PublicEmailLanguageFrVu   PublicEmailLanguage = "fr-vu"
	PublicEmailLanguageFrWf   PublicEmailLanguage = "fr-wf"
	PublicEmailLanguageFrYt   PublicEmailLanguage = "fr-yt"
	PublicEmailLanguageFrr    PublicEmailLanguage = "frr"
	PublicEmailLanguageFrrDe  PublicEmailLanguage = "frr-de"
	PublicEmailLanguageFur    PublicEmailLanguage = "fur"
	PublicEmailLanguageFurIt  PublicEmailLanguage = "fur-it"
	PublicEmailLanguageFy     PublicEmailLanguage = "fy"
	PublicEmailLanguageFyNl   PublicEmailLanguage = "fy-nl"
	PublicEmailLanguageGa     PublicEmailLanguage = "ga"
	PublicEmailLanguageGaGB   PublicEmailLanguage = "ga-gb"
	PublicEmailLanguageGaIe   PublicEmailLanguage = "ga-ie"
	PublicEmailLanguageGaa    PublicEmailLanguage = "gaa"
	PublicEmailLanguageGaaGh  PublicEmailLanguage = "gaa-gh"
	PublicEmailLanguageGd     PublicEmailLanguage = "gd"
	PublicEmailLanguageGdGB   PublicEmailLanguage = "gd-gb"
	PublicEmailLanguageGl     PublicEmailLanguage = "gl"
	PublicEmailLanguageGlEs   PublicEmailLanguage = "gl-es"
	PublicEmailLanguageGn     PublicEmailLanguage = "gn"
	PublicEmailLanguageGsw    PublicEmailLanguage = "gsw"
	PublicEmailLanguageGswCh  PublicEmailLanguage = "gsw-ch"
	PublicEmailLanguageGswFr  PublicEmailLanguage = "gsw-fr"
	PublicEmailLanguageGswLi  PublicEmailLanguage = "gsw-li"
	PublicEmailLanguageGu     PublicEmailLanguage = "gu"
	PublicEmailLanguageGuIn   PublicEmailLanguage = "gu-in"
	PublicEmailLanguageGuz    PublicEmailLanguage = "guz"
	PublicEmailLanguageGuzKe  PublicEmailLanguage = "guz-ke"
	PublicEmailLanguageGv     PublicEmailLanguage = "gv"
	PublicEmailLanguageGvIm   PublicEmailLanguage = "gv-im"
	PublicEmailLanguageHa     PublicEmailLanguage = "ha"
	PublicEmailLanguageHaGh   PublicEmailLanguage = "ha-gh"
	PublicEmailLanguageHaNe   PublicEmailLanguage = "ha-ne"
	PublicEmailLanguageHaNg   PublicEmailLanguage = "ha-ng"
	PublicEmailLanguageHaw    PublicEmailLanguage = "haw"
	PublicEmailLanguageHawUs  PublicEmailLanguage = "haw-us"
	PublicEmailLanguageHe     PublicEmailLanguage = "he"
	PublicEmailLanguageHeIl   PublicEmailLanguage = "he-il"
	PublicEmailLanguageHi     PublicEmailLanguage = "hi"
	PublicEmailLanguageHiIn   PublicEmailLanguage = "hi-in"
	PublicEmailLanguageHmn    PublicEmailLanguage = "hmn"
	PublicEmailLanguageHo     PublicEmailLanguage = "ho"
	PublicEmailLanguageHr     PublicEmailLanguage = "hr"
	PublicEmailLanguageHrBa   PublicEmailLanguage = "hr-ba"
	PublicEmailLanguageHrHr   PublicEmailLanguage = "hr-hr"
	PublicEmailLanguageHsb    PublicEmailLanguage = "hsb"
	PublicEmailLanguageHsbDe  PublicEmailLanguage = "hsb-de"
	PublicEmailLanguageHt     PublicEmailLanguage = "ht"
	PublicEmailLanguageHtHt   PublicEmailLanguage = "ht-ht"
	PublicEmailLanguageHu     PublicEmailLanguage = "hu"
	PublicEmailLanguageHuHu   PublicEmailLanguage = "hu-hu"
	PublicEmailLanguageHy     PublicEmailLanguage = "hy"
	PublicEmailLanguageHyAm   PublicEmailLanguage = "hy-am"
	PublicEmailLanguageHz     PublicEmailLanguage = "hz"
	PublicEmailLanguageIa     PublicEmailLanguage = "ia"
	PublicEmailLanguageIa001  PublicEmailLanguage = "ia-001"
	PublicEmailLanguageID     PublicEmailLanguage = "id"
	PublicEmailLanguageIDID   PublicEmailLanguage = "id-id"
	PublicEmailLanguageIe     PublicEmailLanguage = "ie"
	PublicEmailLanguageIeEe   PublicEmailLanguage = "ie-ee"
	PublicEmailLanguageIg     PublicEmailLanguage = "ig"
	PublicEmailLanguageIgNg   PublicEmailLanguage = "ig-ng"
	PublicEmailLanguageIi     PublicEmailLanguage = "ii"
	PublicEmailLanguageIiCn   PublicEmailLanguage = "ii-cn"
	PublicEmailLanguageIk     PublicEmailLanguage = "ik"
	PublicEmailLanguageIo     PublicEmailLanguage = "io"
	PublicEmailLanguageIs     PublicEmailLanguage = "is"
	PublicEmailLanguageIsIs   PublicEmailLanguage = "is-is"
	PublicEmailLanguageIt     PublicEmailLanguage = "it"
	PublicEmailLanguageItCh   PublicEmailLanguage = "it-ch"
	PublicEmailLanguageItIt   PublicEmailLanguage = "it-it"
	PublicEmailLanguageItSm   PublicEmailLanguage = "it-sm"
	PublicEmailLanguageItVa   PublicEmailLanguage = "it-va"
	PublicEmailLanguageIu     PublicEmailLanguage = "iu"
	PublicEmailLanguageJa     PublicEmailLanguage = "ja"
	PublicEmailLanguageJaJp   PublicEmailLanguage = "ja-jp"
	PublicEmailLanguageJgo    PublicEmailLanguage = "jgo"
	PublicEmailLanguageJgoCm  PublicEmailLanguage = "jgo-cm"
	PublicEmailLanguageJmc    PublicEmailLanguage = "jmc"
	PublicEmailLanguageJmcTz  PublicEmailLanguage = "jmc-tz"
	PublicEmailLanguageJv     PublicEmailLanguage = "jv"
	PublicEmailLanguageJvID   PublicEmailLanguage = "jv-id"
	PublicEmailLanguageKa     PublicEmailLanguage = "ka"
	PublicEmailLanguageKaGe   PublicEmailLanguage = "ka-ge"
	PublicEmailLanguageKab    PublicEmailLanguage = "kab"
	PublicEmailLanguageKabDz  PublicEmailLanguage = "kab-dz"
	PublicEmailLanguageKam    PublicEmailLanguage = "kam"
	PublicEmailLanguageKamKe  PublicEmailLanguage = "kam-ke"
	PublicEmailLanguageKar    PublicEmailLanguage = "kar"
	PublicEmailLanguageKde    PublicEmailLanguage = "kde"
	PublicEmailLanguageKdeTz  PublicEmailLanguage = "kde-tz"
	PublicEmailLanguageKea    PublicEmailLanguage = "kea"
	PublicEmailLanguageKeaCv  PublicEmailLanguage = "kea-cv"
	PublicEmailLanguageKg     PublicEmailLanguage = "kg"
	PublicEmailLanguageKgp    PublicEmailLanguage = "kgp"
	PublicEmailLanguageKgpBr  PublicEmailLanguage = "kgp-br"
	PublicEmailLanguageKh     PublicEmailLanguage = "kh"
	PublicEmailLanguageKhq    PublicEmailLanguage = "khq"
	PublicEmailLanguageKhqMl  PublicEmailLanguage = "khq-ml"
	PublicEmailLanguageKi     PublicEmailLanguage = "ki"
	PublicEmailLanguageKiKe   PublicEmailLanguage = "ki-ke"
	PublicEmailLanguageKj     PublicEmailLanguage = "kj"
	PublicEmailLanguageKk     PublicEmailLanguage = "kk"
	PublicEmailLanguageKkKz   PublicEmailLanguage = "kk-kz"
	PublicEmailLanguageKkj    PublicEmailLanguage = "kkj"
	PublicEmailLanguageKkjCm  PublicEmailLanguage = "kkj-cm"
	PublicEmailLanguageKl     PublicEmailLanguage = "kl"
	PublicEmailLanguageKlGl   PublicEmailLanguage = "kl-gl"
	PublicEmailLanguageKln    PublicEmailLanguage = "kln"
	PublicEmailLanguageKlnKe  PublicEmailLanguage = "kln-ke"
	PublicEmailLanguageKm     PublicEmailLanguage = "km"
	PublicEmailLanguageKmKh   PublicEmailLanguage = "km-kh"
	PublicEmailLanguageKn     PublicEmailLanguage = "kn"
	PublicEmailLanguageKnIn   PublicEmailLanguage = "kn-in"
	PublicEmailLanguageKo     PublicEmailLanguage = "ko"
	PublicEmailLanguageKoCn   PublicEmailLanguage = "ko-cn"
	PublicEmailLanguageKoKp   PublicEmailLanguage = "ko-kp"
	PublicEmailLanguageKoKr   PublicEmailLanguage = "ko-kr"
	PublicEmailLanguageKok    PublicEmailLanguage = "kok"
	PublicEmailLanguageKokIn  PublicEmailLanguage = "kok-in"
	PublicEmailLanguageKr     PublicEmailLanguage = "kr"
	PublicEmailLanguageKs     PublicEmailLanguage = "ks"
	PublicEmailLanguageKsIn   PublicEmailLanguage = "ks-in"
	PublicEmailLanguageKsb    PublicEmailLanguage = "ksb"
	PublicEmailLanguageKsbTz  PublicEmailLanguage = "ksb-tz"
	PublicEmailLanguageKsf    PublicEmailLanguage = "ksf"
	PublicEmailLanguageKsfCm  PublicEmailLanguage = "ksf-cm"
	PublicEmailLanguageKsh    PublicEmailLanguage = "ksh"
	PublicEmailLanguageKshDe  PublicEmailLanguage = "ksh-de"
	PublicEmailLanguageKu     PublicEmailLanguage = "ku"
	PublicEmailLanguageKuTr   PublicEmailLanguage = "ku-tr"
	PublicEmailLanguageKv     PublicEmailLanguage = "kv"
	PublicEmailLanguageKw     PublicEmailLanguage = "kw"
	PublicEmailLanguageKwGB   PublicEmailLanguage = "kw-gb"
	PublicEmailLanguageKxv    PublicEmailLanguage = "kxv"
	PublicEmailLanguageKxvIn  PublicEmailLanguage = "kxv-in"
	PublicEmailLanguageKy     PublicEmailLanguage = "ky"
	PublicEmailLanguageKyKg   PublicEmailLanguage = "ky-kg"
	PublicEmailLanguageLa     PublicEmailLanguage = "la"
	PublicEmailLanguageLag    PublicEmailLanguage = "lag"
	PublicEmailLanguageLagTz  PublicEmailLanguage = "lag-tz"
	PublicEmailLanguageLb     PublicEmailLanguage = "lb"
	PublicEmailLanguageLbLu   PublicEmailLanguage = "lb-lu"
	PublicEmailLanguageLg     PublicEmailLanguage = "lg"
	PublicEmailLanguageLgUg   PublicEmailLanguage = "lg-ug"
	PublicEmailLanguageLi     PublicEmailLanguage = "li"
	PublicEmailLanguageLij    PublicEmailLanguage = "lij"
	PublicEmailLanguageLijIt  PublicEmailLanguage = "lij-it"
	PublicEmailLanguageLkt    PublicEmailLanguage = "lkt"
	PublicEmailLanguageLktUs  PublicEmailLanguage = "lkt-us"
	PublicEmailLanguageLmo    PublicEmailLanguage = "lmo"
	PublicEmailLanguageLmoIt  PublicEmailLanguage = "lmo-it"
	PublicEmailLanguageLn     PublicEmailLanguage = "ln"
	PublicEmailLanguageLnAo   PublicEmailLanguage = "ln-ao"
	PublicEmailLanguageLnCd   PublicEmailLanguage = "ln-cd"
	PublicEmailLanguageLnCf   PublicEmailLanguage = "ln-cf"
	PublicEmailLanguageLnCg   PublicEmailLanguage = "ln-cg"
	PublicEmailLanguageLo     PublicEmailLanguage = "lo"
	PublicEmailLanguageLoLa   PublicEmailLanguage = "lo-la"
	PublicEmailLanguageLrc    PublicEmailLanguage = "lrc"
	PublicEmailLanguageLrcIq  PublicEmailLanguage = "lrc-iq"
	PublicEmailLanguageLrcIr  PublicEmailLanguage = "lrc-ir"
	PublicEmailLanguageLt     PublicEmailLanguage = "lt"
	PublicEmailLanguageLtLt   PublicEmailLanguage = "lt-lt"
	PublicEmailLanguageLu     PublicEmailLanguage = "lu"
	PublicEmailLanguageLuCd   PublicEmailLanguage = "lu-cd"
	PublicEmailLanguageLuo    PublicEmailLanguage = "luo"
	PublicEmailLanguageLuoKe  PublicEmailLanguage = "luo-ke"
	PublicEmailLanguageLuy    PublicEmailLanguage = "luy"
	PublicEmailLanguageLuyKe  PublicEmailLanguage = "luy-ke"
	PublicEmailLanguageLv     PublicEmailLanguage = "lv"
	PublicEmailLanguageLvLv   PublicEmailLanguage = "lv-lv"
	PublicEmailLanguageMai    PublicEmailLanguage = "mai"
	PublicEmailLanguageMaiIn  PublicEmailLanguage = "mai-in"
	PublicEmailLanguageMas    PublicEmailLanguage = "mas"
	PublicEmailLanguageMasKe  PublicEmailLanguage = "mas-ke"
	PublicEmailLanguageMasTz  PublicEmailLanguage = "mas-tz"
	PublicEmailLanguageMdf    PublicEmailLanguage = "mdf"
	PublicEmailLanguageMdfRu  PublicEmailLanguage = "mdf-ru"
	PublicEmailLanguageMer    PublicEmailLanguage = "mer"
	PublicEmailLanguageMerKe  PublicEmailLanguage = "mer-ke"
	PublicEmailLanguageMfe    PublicEmailLanguage = "mfe"
	PublicEmailLanguageMfeMu  PublicEmailLanguage = "mfe-mu"
	PublicEmailLanguageMg     PublicEmailLanguage = "mg"
	PublicEmailLanguageMgMg   PublicEmailLanguage = "mg-mg"
	PublicEmailLanguageMgh    PublicEmailLanguage = "mgh"
	PublicEmailLanguageMghMz  PublicEmailLanguage = "mgh-mz"
	PublicEmailLanguageMgo    PublicEmailLanguage = "mgo"
	PublicEmailLanguageMgoCm  PublicEmailLanguage = "mgo-cm"
	PublicEmailLanguageMh     PublicEmailLanguage = "mh"
	PublicEmailLanguageMi     PublicEmailLanguage = "mi"
	PublicEmailLanguageMiNz   PublicEmailLanguage = "mi-nz"
	PublicEmailLanguageMk     PublicEmailLanguage = "mk"
	PublicEmailLanguageMkMk   PublicEmailLanguage = "mk-mk"
	PublicEmailLanguageMl     PublicEmailLanguage = "ml"
	PublicEmailLanguageMlIn   PublicEmailLanguage = "ml-in"
	PublicEmailLanguageMn     PublicEmailLanguage = "mn"
	PublicEmailLanguageMnMn   PublicEmailLanguage = "mn-mn"
	PublicEmailLanguageMni    PublicEmailLanguage = "mni"
	PublicEmailLanguageMniIn  PublicEmailLanguage = "mni-in"
	PublicEmailLanguageMr     PublicEmailLanguage = "mr"
	PublicEmailLanguageMrIn   PublicEmailLanguage = "mr-in"
	PublicEmailLanguageMs     PublicEmailLanguage = "ms"
	PublicEmailLanguageMsBn   PublicEmailLanguage = "ms-bn"
	PublicEmailLanguageMsID   PublicEmailLanguage = "ms-id"
	PublicEmailLanguageMsMy   PublicEmailLanguage = "ms-my"
	PublicEmailLanguageMsSg   PublicEmailLanguage = "ms-sg"
	PublicEmailLanguageMt     PublicEmailLanguage = "mt"
	PublicEmailLanguageMtMt   PublicEmailLanguage = "mt-mt"
	PublicEmailLanguageMua    PublicEmailLanguage = "mua"
	PublicEmailLanguageMuaCm  PublicEmailLanguage = "mua-cm"
	PublicEmailLanguageMy     PublicEmailLanguage = "my"
	PublicEmailLanguageMyMm   PublicEmailLanguage = "my-mm"
	PublicEmailLanguageMzn    PublicEmailLanguage = "mzn"
	PublicEmailLanguageMznIr  PublicEmailLanguage = "mzn-ir"
	PublicEmailLanguageNa     PublicEmailLanguage = "na"
	PublicEmailLanguageNaq    PublicEmailLanguage = "naq"
	PublicEmailLanguageNaqNa  PublicEmailLanguage = "naq-na"
	PublicEmailLanguageNb     PublicEmailLanguage = "nb"
	PublicEmailLanguageNbNo   PublicEmailLanguage = "nb-no"
	PublicEmailLanguageNbSj   PublicEmailLanguage = "nb-sj"
	PublicEmailLanguageNd     PublicEmailLanguage = "nd"
	PublicEmailLanguageNdZw   PublicEmailLanguage = "nd-zw"
	PublicEmailLanguageNds    PublicEmailLanguage = "nds"
	PublicEmailLanguageNdsDe  PublicEmailLanguage = "nds-de"
	PublicEmailLanguageNdsNl  PublicEmailLanguage = "nds-nl"
	PublicEmailLanguageNe     PublicEmailLanguage = "ne"
	PublicEmailLanguageNeIn   PublicEmailLanguage = "ne-in"
	PublicEmailLanguageNeNp   PublicEmailLanguage = "ne-np"
	PublicEmailLanguageNg     PublicEmailLanguage = "ng"
	PublicEmailLanguageNl     PublicEmailLanguage = "nl"
	PublicEmailLanguageNlAw   PublicEmailLanguage = "nl-aw"
	PublicEmailLanguageNlBe   PublicEmailLanguage = "nl-be"
	PublicEmailLanguageNlBq   PublicEmailLanguage = "nl-bq"
	PublicEmailLanguageNlCh   PublicEmailLanguage = "nl-ch"
	PublicEmailLanguageNlCw   PublicEmailLanguage = "nl-cw"
	PublicEmailLanguageNlLu   PublicEmailLanguage = "nl-lu"
	PublicEmailLanguageNlNl   PublicEmailLanguage = "nl-nl"
	PublicEmailLanguageNlSr   PublicEmailLanguage = "nl-sr"
	PublicEmailLanguageNlSx   PublicEmailLanguage = "nl-sx"
	PublicEmailLanguageNmg    PublicEmailLanguage = "nmg"
	PublicEmailLanguageNmgCm  PublicEmailLanguage = "nmg-cm"
	PublicEmailLanguageNn     PublicEmailLanguage = "nn"
	PublicEmailLanguageNnNo   PublicEmailLanguage = "nn-no"
	PublicEmailLanguageNnh    PublicEmailLanguage = "nnh"
	PublicEmailLanguageNnhCm  PublicEmailLanguage = "nnh-cm"
	PublicEmailLanguageNo     PublicEmailLanguage = "no"
	PublicEmailLanguageNoNo   PublicEmailLanguage = "no-no"
	PublicEmailLanguageNqo    PublicEmailLanguage = "nqo"
	PublicEmailLanguageNqoGn  PublicEmailLanguage = "nqo-gn"
	PublicEmailLanguageNr     PublicEmailLanguage = "nr"
	PublicEmailLanguageNso    PublicEmailLanguage = "nso"
	PublicEmailLanguageNsoZa  PublicEmailLanguage = "nso-za"
	PublicEmailLanguageNus    PublicEmailLanguage = "nus"
	PublicEmailLanguageNusSS  PublicEmailLanguage = "nus-ss"
	PublicEmailLanguageNv     PublicEmailLanguage = "nv"
	PublicEmailLanguageNy     PublicEmailLanguage = "ny"
	PublicEmailLanguageNyn    PublicEmailLanguage = "nyn"
	PublicEmailLanguageNynUg  PublicEmailLanguage = "nyn-ug"
	PublicEmailLanguageOc     PublicEmailLanguage = "oc"
	PublicEmailLanguageOcEs   PublicEmailLanguage = "oc-es"
	PublicEmailLanguageOcFr   PublicEmailLanguage = "oc-fr"
	PublicEmailLanguageOj     PublicEmailLanguage = "oj"
	PublicEmailLanguageOm     PublicEmailLanguage = "om"
	PublicEmailLanguageOmEt   PublicEmailLanguage = "om-et"
	PublicEmailLanguageOmKe   PublicEmailLanguage = "om-ke"
	PublicEmailLanguageOr     PublicEmailLanguage = "or"
	PublicEmailLanguageOrIn   PublicEmailLanguage = "or-in"
	PublicEmailLanguageOs     PublicEmailLanguage = "os"
	PublicEmailLanguageOsGe   PublicEmailLanguage = "os-ge"
	PublicEmailLanguageOsRu   PublicEmailLanguage = "os-ru"
	PublicEmailLanguagePa     PublicEmailLanguage = "pa"
	PublicEmailLanguagePaIn   PublicEmailLanguage = "pa-in"
	PublicEmailLanguagePaPk   PublicEmailLanguage = "pa-pk"
	PublicEmailLanguagePcm    PublicEmailLanguage = "pcm"
	PublicEmailLanguagePcmNg  PublicEmailLanguage = "pcm-ng"
	PublicEmailLanguagePi     PublicEmailLanguage = "pi"
	PublicEmailLanguagePis    PublicEmailLanguage = "pis"
	PublicEmailLanguagePisSb  PublicEmailLanguage = "pis-sb"
	PublicEmailLanguagePl     PublicEmailLanguage = "pl"
	PublicEmailLanguagePlPl   PublicEmailLanguage = "pl-pl"
	PublicEmailLanguagePrg    PublicEmailLanguage = "prg"
	PublicEmailLanguagePrg001 PublicEmailLanguage = "prg-001"
	PublicEmailLanguagePs     PublicEmailLanguage = "ps"
	PublicEmailLanguagePsAf   PublicEmailLanguage = "ps-af"
	PublicEmailLanguagePsPk   PublicEmailLanguage = "ps-pk"
	PublicEmailLanguagePt     PublicEmailLanguage = "pt"
	PublicEmailLanguagePtAo   PublicEmailLanguage = "pt-ao"
	PublicEmailLanguagePtBr   PublicEmailLanguage = "pt-br"
	PublicEmailLanguagePtCh   PublicEmailLanguage = "pt-ch"
	PublicEmailLanguagePtCv   PublicEmailLanguage = "pt-cv"
	PublicEmailLanguagePtGq   PublicEmailLanguage = "pt-gq"
	PublicEmailLanguagePtGw   PublicEmailLanguage = "pt-gw"
	PublicEmailLanguagePtLu   PublicEmailLanguage = "pt-lu"
	PublicEmailLanguagePtMo   PublicEmailLanguage = "pt-mo"
	PublicEmailLanguagePtMz   PublicEmailLanguage = "pt-mz"
	PublicEmailLanguagePtPt   PublicEmailLanguage = "pt-pt"
	PublicEmailLanguagePtSt   PublicEmailLanguage = "pt-st"
	PublicEmailLanguagePtTl   PublicEmailLanguage = "pt-tl"
	PublicEmailLanguageQu     PublicEmailLanguage = "qu"
	PublicEmailLanguageQuBo   PublicEmailLanguage = "qu-bo"
	PublicEmailLanguageQuEc   PublicEmailLanguage = "qu-ec"
	PublicEmailLanguageQuPe   PublicEmailLanguage = "qu-pe"
	PublicEmailLanguageRaj    PublicEmailLanguage = "raj"
	PublicEmailLanguageRajIn  PublicEmailLanguage = "raj-in"
	PublicEmailLanguageRm     PublicEmailLanguage = "rm"
	PublicEmailLanguageRmCh   PublicEmailLanguage = "rm-ch"
	PublicEmailLanguageRn     PublicEmailLanguage = "rn"
	PublicEmailLanguageRnBi   PublicEmailLanguage = "rn-bi"
	PublicEmailLanguageRo     PublicEmailLanguage = "ro"
	PublicEmailLanguageRoMd   PublicEmailLanguage = "ro-md"
	PublicEmailLanguageRoRo   PublicEmailLanguage = "ro-ro"
	PublicEmailLanguageRof    PublicEmailLanguage = "rof"
	PublicEmailLanguageRofTz  PublicEmailLanguage = "rof-tz"
	PublicEmailLanguageRu     PublicEmailLanguage = "ru"
	PublicEmailLanguageRuBy   PublicEmailLanguage = "ru-by"
	PublicEmailLanguageRuKg   PublicEmailLanguage = "ru-kg"
	PublicEmailLanguageRuKz   PublicEmailLanguage = "ru-kz"
	PublicEmailLanguageRuMd   PublicEmailLanguage = "ru-md"
	PublicEmailLanguageRuRu   PublicEmailLanguage = "ru-ru"
	PublicEmailLanguageRuUa   PublicEmailLanguage = "ru-ua"
	PublicEmailLanguageRw     PublicEmailLanguage = "rw"
	PublicEmailLanguageRwRw   PublicEmailLanguage = "rw-rw"
	PublicEmailLanguageRwk    PublicEmailLanguage = "rwk"
	PublicEmailLanguageRwkTz  PublicEmailLanguage = "rwk-tz"
	PublicEmailLanguageSa     PublicEmailLanguage = "sa"
	PublicEmailLanguageSaIn   PublicEmailLanguage = "sa-in"
	PublicEmailLanguageSah    PublicEmailLanguage = "sah"
	PublicEmailLanguageSahRu  PublicEmailLanguage = "sah-ru"
	PublicEmailLanguageSaq    PublicEmailLanguage = "saq"
	PublicEmailLanguageSaqKe  PublicEmailLanguage = "saq-ke"
	PublicEmailLanguageSat    PublicEmailLanguage = "sat"
	PublicEmailLanguageSatIn  PublicEmailLanguage = "sat-in"
	PublicEmailLanguageSbp    PublicEmailLanguage = "sbp"
	PublicEmailLanguageSbpTz  PublicEmailLanguage = "sbp-tz"
	PublicEmailLanguageSc     PublicEmailLanguage = "sc"
	PublicEmailLanguageScIt   PublicEmailLanguage = "sc-it"
	PublicEmailLanguageSd     PublicEmailLanguage = "sd"
	PublicEmailLanguageSdIn   PublicEmailLanguage = "sd-in"
	PublicEmailLanguageSdPk   PublicEmailLanguage = "sd-pk"
	PublicEmailLanguageSe     PublicEmailLanguage = "se"
	PublicEmailLanguageSeFi   PublicEmailLanguage = "se-fi"
	PublicEmailLanguageSeNo   PublicEmailLanguage = "se-no"
	PublicEmailLanguageSeSe   PublicEmailLanguage = "se-se"
	PublicEmailLanguageSeh    PublicEmailLanguage = "seh"
	PublicEmailLanguageSehMz  PublicEmailLanguage = "seh-mz"
	PublicEmailLanguageSes    PublicEmailLanguage = "ses"
	PublicEmailLanguageSesMl  PublicEmailLanguage = "ses-ml"
	PublicEmailLanguageSg     PublicEmailLanguage = "sg"
	PublicEmailLanguageSgCf   PublicEmailLanguage = "sg-cf"
	PublicEmailLanguageShi    PublicEmailLanguage = "shi"
	PublicEmailLanguageShiMa  PublicEmailLanguage = "shi-ma"
	PublicEmailLanguageSi     PublicEmailLanguage = "si"
	PublicEmailLanguageSiLk   PublicEmailLanguage = "si-lk"
	PublicEmailLanguageSk     PublicEmailLanguage = "sk"
	PublicEmailLanguageSkSk   PublicEmailLanguage = "sk-sk"
	PublicEmailLanguageSl     PublicEmailLanguage = "sl"
	PublicEmailLanguageSlSi   PublicEmailLanguage = "sl-si"
	PublicEmailLanguageSm     PublicEmailLanguage = "sm"
	PublicEmailLanguageSmn    PublicEmailLanguage = "smn"
	PublicEmailLanguageSmnFi  PublicEmailLanguage = "smn-fi"
	PublicEmailLanguageSMS    PublicEmailLanguage = "sms"
	PublicEmailLanguageSMSFi  PublicEmailLanguage = "sms-fi"
	PublicEmailLanguageSn     PublicEmailLanguage = "sn"
	PublicEmailLanguageSnZw   PublicEmailLanguage = "sn-zw"
	PublicEmailLanguageSo     PublicEmailLanguage = "so"
	PublicEmailLanguageSoDj   PublicEmailLanguage = "so-dj"
	PublicEmailLanguageSoEt   PublicEmailLanguage = "so-et"
	PublicEmailLanguageSoKe   PublicEmailLanguage = "so-ke"
	PublicEmailLanguageSoSo   PublicEmailLanguage = "so-so"
	PublicEmailLanguageSq     PublicEmailLanguage = "sq"
	PublicEmailLanguageSqAl   PublicEmailLanguage = "sq-al"
	PublicEmailLanguageSqMk   PublicEmailLanguage = "sq-mk"
	PublicEmailLanguageSqXk   PublicEmailLanguage = "sq-xk"
	PublicEmailLanguageSr     PublicEmailLanguage = "sr"
	PublicEmailLanguageSrBa   PublicEmailLanguage = "sr-ba"
	PublicEmailLanguageSrCs   PublicEmailLanguage = "sr-cs"
	PublicEmailLanguageSrMe   PublicEmailLanguage = "sr-me"
	PublicEmailLanguageSrRs   PublicEmailLanguage = "sr-rs"
	PublicEmailLanguageSrXk   PublicEmailLanguage = "sr-xk"
	PublicEmailLanguageSS     PublicEmailLanguage = "ss"
	PublicEmailLanguageSt     PublicEmailLanguage = "st"
	PublicEmailLanguageStLs   PublicEmailLanguage = "st-ls"
	PublicEmailLanguageStZa   PublicEmailLanguage = "st-za"
	PublicEmailLanguageSu     PublicEmailLanguage = "su"
	PublicEmailLanguageSuID   PublicEmailLanguage = "su-id"
	PublicEmailLanguageSv     PublicEmailLanguage = "sv"
	PublicEmailLanguageSvAx   PublicEmailLanguage = "sv-ax"
	PublicEmailLanguageSvFi   PublicEmailLanguage = "sv-fi"
	PublicEmailLanguageSvSe   PublicEmailLanguage = "sv-se"
	PublicEmailLanguageSw     PublicEmailLanguage = "sw"
	PublicEmailLanguageSwCd   PublicEmailLanguage = "sw-cd"
	PublicEmailLanguageSwKe   PublicEmailLanguage = "sw-ke"
	PublicEmailLanguageSwTz   PublicEmailLanguage = "sw-tz"
	PublicEmailLanguageSwUg   PublicEmailLanguage = "sw-ug"
	PublicEmailLanguageSy     PublicEmailLanguage = "sy"
	PublicEmailLanguageSyr    PublicEmailLanguage = "syr"
	PublicEmailLanguageSyrIq  PublicEmailLanguage = "syr-iq"
	PublicEmailLanguageSyrSy  PublicEmailLanguage = "syr-sy"
	PublicEmailLanguageSzl    PublicEmailLanguage = "szl"
	PublicEmailLanguageSzlPl  PublicEmailLanguage = "szl-pl"
	PublicEmailLanguageTa     PublicEmailLanguage = "ta"
	PublicEmailLanguageTaIn   PublicEmailLanguage = "ta-in"
	PublicEmailLanguageTaLk   PublicEmailLanguage = "ta-lk"
	PublicEmailLanguageTaMy   PublicEmailLanguage = "ta-my"
	PublicEmailLanguageTaSg   PublicEmailLanguage = "ta-sg"
	PublicEmailLanguageTe     PublicEmailLanguage = "te"
	PublicEmailLanguageTeIn   PublicEmailLanguage = "te-in"
	PublicEmailLanguageTeo    PublicEmailLanguage = "teo"
	PublicEmailLanguageTeoKe  PublicEmailLanguage = "teo-ke"
	PublicEmailLanguageTeoUg  PublicEmailLanguage = "teo-ug"
	PublicEmailLanguageTg     PublicEmailLanguage = "tg"
	PublicEmailLanguageTgTj   PublicEmailLanguage = "tg-tj"
	PublicEmailLanguageTh     PublicEmailLanguage = "th"
	PublicEmailLanguageThTh   PublicEmailLanguage = "th-th"
	PublicEmailLanguageTi     PublicEmailLanguage = "ti"
	PublicEmailLanguageTiEr   PublicEmailLanguage = "ti-er"
	PublicEmailLanguageTiEt   PublicEmailLanguage = "ti-et"
	PublicEmailLanguageTk     PublicEmailLanguage = "tk"
	PublicEmailLanguageTkTm   PublicEmailLanguage = "tk-tm"
	PublicEmailLanguageTl     PublicEmailLanguage = "tl"
	PublicEmailLanguageTn     PublicEmailLanguage = "tn"
	PublicEmailLanguageTnBw   PublicEmailLanguage = "tn-bw"
	PublicEmailLanguageTnZa   PublicEmailLanguage = "tn-za"
	PublicEmailLanguageTo     PublicEmailLanguage = "to"
	PublicEmailLanguageToTo   PublicEmailLanguage = "to-to"
	PublicEmailLanguageTok    PublicEmailLanguage = "tok"
	PublicEmailLanguageTok001 PublicEmailLanguage = "tok-001"
	PublicEmailLanguageTr     PublicEmailLanguage = "tr"
	PublicEmailLanguageTrCy   PublicEmailLanguage = "tr-cy"
	PublicEmailLanguageTrTr   PublicEmailLanguage = "tr-tr"
	PublicEmailLanguageTs     PublicEmailLanguage = "ts"
	PublicEmailLanguageTt     PublicEmailLanguage = "tt"
	PublicEmailLanguageTtRu   PublicEmailLanguage = "tt-ru"
	PublicEmailLanguageTw     PublicEmailLanguage = "tw"
	PublicEmailLanguageTwq    PublicEmailLanguage = "twq"
	PublicEmailLanguageTwqNe  PublicEmailLanguage = "twq-ne"
	PublicEmailLanguageTy     PublicEmailLanguage = "ty"
	PublicEmailLanguageTzm    PublicEmailLanguage = "tzm"
	PublicEmailLanguageTzmMa  PublicEmailLanguage = "tzm-ma"
	PublicEmailLanguageUg     PublicEmailLanguage = "ug"
	PublicEmailLanguageUgCn   PublicEmailLanguage = "ug-cn"
	PublicEmailLanguageUk     PublicEmailLanguage = "uk"
	PublicEmailLanguageUkUa   PublicEmailLanguage = "uk-ua"
	PublicEmailLanguageUr     PublicEmailLanguage = "ur"
	PublicEmailLanguageUrIn   PublicEmailLanguage = "ur-in"
	PublicEmailLanguageUrPk   PublicEmailLanguage = "ur-pk"
	PublicEmailLanguageUz     PublicEmailLanguage = "uz"
	PublicEmailLanguageUzAf   PublicEmailLanguage = "uz-af"
	PublicEmailLanguageUzUz   PublicEmailLanguage = "uz-uz"
	PublicEmailLanguageVai    PublicEmailLanguage = "vai"
	PublicEmailLanguageVaiLr  PublicEmailLanguage = "vai-lr"
	PublicEmailLanguageVe     PublicEmailLanguage = "ve"
	PublicEmailLanguageVec    PublicEmailLanguage = "vec"
	PublicEmailLanguageVecIt  PublicEmailLanguage = "vec-it"
	PublicEmailLanguageVi     PublicEmailLanguage = "vi"
	PublicEmailLanguageViVn   PublicEmailLanguage = "vi-vn"
	PublicEmailLanguageVmw    PublicEmailLanguage = "vmw"
	PublicEmailLanguageVmwMz  PublicEmailLanguage = "vmw-mz"
	PublicEmailLanguageVo     PublicEmailLanguage = "vo"
	PublicEmailLanguageVo001  PublicEmailLanguage = "vo-001"
	PublicEmailLanguageVun    PublicEmailLanguage = "vun"
	PublicEmailLanguageVunTz  PublicEmailLanguage = "vun-tz"
	PublicEmailLanguageWa     PublicEmailLanguage = "wa"
	PublicEmailLanguageWae    PublicEmailLanguage = "wae"
	PublicEmailLanguageWaeCh  PublicEmailLanguage = "wae-ch"
	PublicEmailLanguageWo     PublicEmailLanguage = "wo"
	PublicEmailLanguageWoSn   PublicEmailLanguage = "wo-sn"
	PublicEmailLanguageXh     PublicEmailLanguage = "xh"
	PublicEmailLanguageXhZa   PublicEmailLanguage = "xh-za"
	PublicEmailLanguageXnr    PublicEmailLanguage = "xnr"
	PublicEmailLanguageXnrIn  PublicEmailLanguage = "xnr-in"
	PublicEmailLanguageXog    PublicEmailLanguage = "xog"
	PublicEmailLanguageXogUg  PublicEmailLanguage = "xog-ug"
	PublicEmailLanguageYav    PublicEmailLanguage = "yav"
	PublicEmailLanguageYavCm  PublicEmailLanguage = "yav-cm"
	PublicEmailLanguageYi     PublicEmailLanguage = "yi"
	PublicEmailLanguageYi001  PublicEmailLanguage = "yi-001"
	PublicEmailLanguageYiUa   PublicEmailLanguage = "yi-ua"
	PublicEmailLanguageYo     PublicEmailLanguage = "yo"
	PublicEmailLanguageYoBj   PublicEmailLanguage = "yo-bj"
	PublicEmailLanguageYoNg   PublicEmailLanguage = "yo-ng"
	PublicEmailLanguageYrl    PublicEmailLanguage = "yrl"
	PublicEmailLanguageYrlBr  PublicEmailLanguage = "yrl-br"
	PublicEmailLanguageYrlCo  PublicEmailLanguage = "yrl-co"
	PublicEmailLanguageYrlVe  PublicEmailLanguage = "yrl-ve"
	PublicEmailLanguageYue    PublicEmailLanguage = "yue"
	PublicEmailLanguageYueCn  PublicEmailLanguage = "yue-cn"
	PublicEmailLanguageYueHk  PublicEmailLanguage = "yue-hk"
	PublicEmailLanguageYueMo  PublicEmailLanguage = "yue-mo"
	PublicEmailLanguageZa     PublicEmailLanguage = "za"
	PublicEmailLanguageZaCn   PublicEmailLanguage = "za-cn"
	PublicEmailLanguageZgh    PublicEmailLanguage = "zgh"
	PublicEmailLanguageZghMa  PublicEmailLanguage = "zgh-ma"
	PublicEmailLanguageZh     PublicEmailLanguage = "zh"
	PublicEmailLanguageZhCn   PublicEmailLanguage = "zh-cn"
	PublicEmailLanguageZhHans PublicEmailLanguage = "zh-hans"
	PublicEmailLanguageZhHant PublicEmailLanguage = "zh-hant"
	PublicEmailLanguageZhHk   PublicEmailLanguage = "zh-hk"
	PublicEmailLanguageZhMo   PublicEmailLanguage = "zh-mo"
	PublicEmailLanguageZhMy   PublicEmailLanguage = "zh-my"
	PublicEmailLanguageZhSg   PublicEmailLanguage = "zh-sg"
	PublicEmailLanguageZhTw   PublicEmailLanguage = "zh-tw"
	PublicEmailLanguageZu     PublicEmailLanguage = "zu"
	PublicEmailLanguageZuZa   PublicEmailLanguage = "zu-za"
)

// The email state.
type PublicEmailState string

const (
	PublicEmailStateAgentGenerated          PublicEmailState = "AGENT_GENERATED"
	PublicEmailStateAutomated               PublicEmailState = "AUTOMATED"
	PublicEmailStateAutomatedAb             PublicEmailState = "AUTOMATED_AB"
	PublicEmailStateAutomatedAbVariant      PublicEmailState = "AUTOMATED_AB_VARIANT"
	PublicEmailStateAutomatedDraft          PublicEmailState = "AUTOMATED_DRAFT"
	PublicEmailStateAutomatedDraftAb        PublicEmailState = "AUTOMATED_DRAFT_AB"
	PublicEmailStateAutomatedDraftAbvariant PublicEmailState = "AUTOMATED_DRAFT_ABVARIANT"
	PublicEmailStateAutomatedForForm        PublicEmailState = "AUTOMATED_FOR_FORM"
	PublicEmailStateAutomatedForFormBuffer  PublicEmailState = "AUTOMATED_FOR_FORM_BUFFER"
	PublicEmailStateAutomatedForFormDraft   PublicEmailState = "AUTOMATED_FOR_FORM_DRAFT"
	PublicEmailStateAutomatedForFormLegacy  PublicEmailState = "AUTOMATED_FOR_FORM_LEGACY"
	PublicEmailStateAutomatedLoserAbvariant PublicEmailState = "AUTOMATED_LOSER_ABVARIANT"
	PublicEmailStateAutomatedSending        PublicEmailState = "AUTOMATED_SENDING"
	PublicEmailStateBlogEmailDraft          PublicEmailState = "BLOG_EMAIL_DRAFT"
	PublicEmailStateBlogEmailPublished      PublicEmailState = "BLOG_EMAIL_PUBLISHED"
	PublicEmailStateDraft                   PublicEmailState = "DRAFT"
	PublicEmailStateDraftAb                 PublicEmailState = "DRAFT_AB"
	PublicEmailStateDraftAbVariant          PublicEmailState = "DRAFT_AB_VARIANT"
	PublicEmailStateError                   PublicEmailState = "ERROR"
	PublicEmailStateLoserAbVariant          PublicEmailState = "LOSER_AB_VARIANT"
	PublicEmailStatePageStub                PublicEmailState = "PAGE_STUB"
	PublicEmailStatePreProcessing           PublicEmailState = "PRE_PROCESSING"
	PublicEmailStateProcessing              PublicEmailState = "PROCESSING"
	PublicEmailStatePublished               PublicEmailState = "PUBLISHED"
	PublicEmailStatePublishedAb             PublicEmailState = "PUBLISHED_AB"
	PublicEmailStatePublishedAbVariant      PublicEmailState = "PUBLISHED_AB_VARIANT"
	PublicEmailStatePublishedOrScheduled    PublicEmailState = "PUBLISHED_OR_SCHEDULED"
	PublicEmailStateRssToEmailDraft         PublicEmailState = "RSS_TO_EMAIL_DRAFT"
	PublicEmailStateRssToEmailPublished     PublicEmailState = "RSS_TO_EMAIL_PUBLISHED"
	PublicEmailStateScheduled               PublicEmailState = "SCHEDULED"
	PublicEmailStateScheduledAb             PublicEmailState = "SCHEDULED_AB"
	PublicEmailStateScheduledOrPublished    PublicEmailState = "SCHEDULED_OR_PUBLISHED"
)

// The email type, this is derived from other properties on the email such as
// subcategory.
type PublicEmailType string

const (
	PublicEmailTypeAbEmail                             PublicEmailType = "AB_EMAIL"
	PublicEmailTypeAutomatedAbEmail                    PublicEmailType = "AUTOMATED_AB_EMAIL"
	PublicEmailTypeAutomatedEmail                      PublicEmailType = "AUTOMATED_EMAIL"
	PublicEmailTypeBatchEmail                          PublicEmailType = "BATCH_EMAIL"
	PublicEmailTypeBlogEmail                           PublicEmailType = "BLOG_EMAIL"
	PublicEmailTypeBlogEmailChild                      PublicEmailType = "BLOG_EMAIL_CHILD"
	PublicEmailTypeFeedbackCesEmail                    PublicEmailType = "FEEDBACK_CES_EMAIL"
	PublicEmailTypeFeedbackCustomEmail                 PublicEmailType = "FEEDBACK_CUSTOM_EMAIL"
	PublicEmailTypeFeedbackCustomSurveyEmail           PublicEmailType = "FEEDBACK_CUSTOM_SURVEY_EMAIL"
	PublicEmailTypeFeedbackNpsEmail                    PublicEmailType = "FEEDBACK_NPS_EMAIL"
	PublicEmailTypeFollowupEmail                       PublicEmailType = "FOLLOWUP_EMAIL"
	PublicEmailTypeLeadflowEmail                       PublicEmailType = "LEADFLOW_EMAIL"
	PublicEmailTypeLocaltimeEmail                      PublicEmailType = "LOCALTIME_EMAIL"
	PublicEmailTypeManagePreferencesEmail              PublicEmailType = "MANAGE_PREFERENCES_EMAIL"
	PublicEmailTypeMarketingSingleSendAPI              PublicEmailType = "MARKETING_SINGLE_SEND_API"
	PublicEmailTypeMembershipEmailVerificationEmail    PublicEmailType = "MEMBERSHIP_EMAIL_VERIFICATION_EMAIL"
	PublicEmailTypeMembershipFollowUpEmail             PublicEmailType = "MEMBERSHIP_FOLLOW_UP_EMAIL"
	PublicEmailTypeMembershipOtpLoginEmail             PublicEmailType = "MEMBERSHIP_OTP_LOGIN_EMAIL"
	PublicEmailTypeMembershipPasswordResetEmail        PublicEmailType = "MEMBERSHIP_PASSWORD_RESET_EMAIL"
	PublicEmailTypeMembershipPasswordSavedEmail        PublicEmailType = "MEMBERSHIP_PASSWORD_SAVED_EMAIL"
	PublicEmailTypeMembershipPasswordlessAuthEmail     PublicEmailType = "MEMBERSHIP_PASSWORDLESS_AUTH_EMAIL"
	PublicEmailTypeMembershipRegistrationEmail         PublicEmailType = "MEMBERSHIP_REGISTRATION_EMAIL"
	PublicEmailTypeMembershipRegistrationFollowUpEmail PublicEmailType = "MEMBERSHIP_REGISTRATION_FOLLOW_UP_EMAIL"
	PublicEmailTypeMembershipVerificationEmail         PublicEmailType = "MEMBERSHIP_VERIFICATION_EMAIL"
	PublicEmailTypeOptinEmail                          PublicEmailType = "OPTIN_EMAIL"
	PublicEmailTypeOptinFollowupEmail                  PublicEmailType = "OPTIN_FOLLOWUP_EMAIL"
	PublicEmailTypeResubscribeEmail                    PublicEmailType = "RESUBSCRIBE_EMAIL"
	PublicEmailTypeRssEmail                            PublicEmailType = "RSS_EMAIL"
	PublicEmailTypeRssEmailChild                       PublicEmailType = "RSS_EMAIL_CHILD"
	PublicEmailTypeSingleSendAPI                       PublicEmailType = "SINGLE_SEND_API"
	PublicEmailTypeSmtpToken                           PublicEmailType = "SMTP_TOKEN"
	PublicEmailTypeTicketEmail                         PublicEmailType = "TICKET_EMAIL"
)

type PublicEmailContent struct {
	FlexAreas           map[string]any             `json:"flexAreas"`
	PlainTextVersion    string                     `json:"plainTextVersion"`
	SmartFields         map[string]SmartEmailField `json:"smartFields"`
	StyleSettings       PublicEmailStyleSettings   `json:"styleSettings"`
	TemplatePath        string                     `json:"templatePath"`
	ThemeSettingsValues map[string]any             `json:"themeSettingsValues"`
	WidgetContainers    map[string]any             `json:"widgetContainers"`
	Widgets             map[string]any             `json:"widgets"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FlexAreas           respjson.Field
		PlainTextVersion    respjson.Field
		SmartFields         respjson.Field
		StyleSettings       respjson.Field
		TemplatePath        respjson.Field
		ThemeSettingsValues respjson.Field
		WidgetContainers    respjson.Field
		Widgets             respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PublicEmailContent) RawJSON() string { return r.JSON.raw }
func (r *PublicEmailContent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this PublicEmailContent to a PublicEmailContentParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// PublicEmailContentParam.Overrides()
func (r PublicEmailContent) ToParam() PublicEmailContentParam {
	return param.Override[PublicEmailContentParam](json.RawMessage(r.RawJSON()))
}

type PublicEmailContentParam struct {
	PlainTextVersion    param.Opt[string]             `json:"plainTextVersion,omitzero"`
	TemplatePath        param.Opt[string]             `json:"templatePath,omitzero"`
	FlexAreas           map[string]any                `json:"flexAreas,omitzero"`
	SmartFields         map[string]SmartEmailField    `json:"smartFields,omitzero"`
	StyleSettings       PublicEmailStyleSettingsParam `json:"styleSettings,omitzero"`
	ThemeSettingsValues map[string]any                `json:"themeSettingsValues,omitzero"`
	WidgetContainers    map[string]any                `json:"widgetContainers,omitzero"`
	Widgets             map[string]any                `json:"widgets,omitzero"`
	paramObj
}

func (r PublicEmailContentParam) MarshalJSON() (data []byte, err error) {
	type shadow PublicEmailContentParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PublicEmailContentParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PublicEmailFromDetails struct {
	// The reply to recipients will see.
	CustomReplyTo string `json:"customReplyTo"`
	// The name recipients will see.
	FromName string `json:"fromName"`
	// The from address and reply to email address (if no customReplyTo defined)
	// recipients will see.
	ReplyTo string `json:"replyTo"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CustomReplyTo respjson.Field
		FromName      respjson.Field
		ReplyTo       respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PublicEmailFromDetails) RawJSON() string { return r.JSON.raw }
func (r *PublicEmailFromDetails) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this PublicEmailFromDetails to a PublicEmailFromDetailsParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// PublicEmailFromDetailsParam.Overrides()
func (r PublicEmailFromDetails) ToParam() PublicEmailFromDetailsParam {
	return param.Override[PublicEmailFromDetailsParam](json.RawMessage(r.RawJSON()))
}

type PublicEmailFromDetailsParam struct {
	// The reply to recipients will see.
	CustomReplyTo param.Opt[string] `json:"customReplyTo,omitzero"`
	// The name recipients will see.
	FromName param.Opt[string] `json:"fromName,omitzero"`
	// The from address and reply to email address (if no customReplyTo defined)
	// recipients will see.
	ReplyTo param.Opt[string] `json:"replyTo,omitzero"`
	paramObj
}

func (r PublicEmailFromDetailsParam) MarshalJSON() (data []byte, err error) {
	type shadow PublicEmailFromDetailsParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PublicEmailFromDetailsParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PublicEmailRecipients struct {
	// Excluded IDs.
	Exclude []string `json:"exclude"`
	// Included IDs.
	Include []string `json:"include"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Exclude     respjson.Field
		Include     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PublicEmailRecipients) RawJSON() string { return r.JSON.raw }
func (r *PublicEmailRecipients) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this PublicEmailRecipients to a PublicEmailRecipientsParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// PublicEmailRecipientsParam.Overrides()
func (r PublicEmailRecipients) ToParam() PublicEmailRecipientsParam {
	return param.Override[PublicEmailRecipientsParam](json.RawMessage(r.RawJSON()))
}

type PublicEmailRecipientsParam struct {
	// Excluded IDs.
	Exclude []string `json:"exclude,omitzero"`
	// Included IDs.
	Include []string `json:"include,omitzero"`
	paramObj
}

func (r PublicEmailRecipientsParam) MarshalJSON() (data []byte, err error) {
	type shadow PublicEmailRecipientsParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PublicEmailRecipientsParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PublicEmailStyleSettings struct {
	BackgroundColor string `json:"backgroundColor"`
	BackgroundImage string `json:"backgroundImage"`
	// Any of "REPEAT", "SINGLE", "STRETCH".
	BackgroundImageType     PublicEmailStyleSettingsBackgroundImageType `json:"backgroundImageType"`
	BodyBorderColor         string                                      `json:"bodyBorderColor"`
	BodyBorderColorChoice   string                                      `json:"bodyBorderColorChoice"`
	BodyBorderWidth         float64                                     `json:"bodyBorderWidth"`
	BodyColor               string                                      `json:"bodyColor"`
	ButtonStyleSettings     PublicButtonStyleSettings                   `json:"buttonStyleSettings"`
	ColorPickerFavorite1    string                                      `json:"colorPickerFavorite1"`
	ColorPickerFavorite2    string                                      `json:"colorPickerFavorite2"`
	ColorPickerFavorite3    string                                      `json:"colorPickerFavorite3"`
	ColorPickerFavorite4    string                                      `json:"colorPickerFavorite4"`
	ColorPickerFavorite5    string                                      `json:"colorPickerFavorite5"`
	ColorPickerFavorite6    string                                      `json:"colorPickerFavorite6"`
	DividerStyleSettings    PublicDividerStyleSettings                  `json:"dividerStyleSettings"`
	EmailBodyPadding        string                                      `json:"emailBodyPadding"`
	EmailBodyWidth          string                                      `json:"emailBodyWidth"`
	HeadingOneFont          PublicFontStyle                             `json:"headingOneFont"`
	HeadingTwoFont          PublicFontStyle                             `json:"headingTwoFont"`
	LinksFont               PublicFontStyle                             `json:"linksFont"`
	PrimaryAccentColor      string                                      `json:"primaryAccentColor"`
	PrimaryFont             string                                      `json:"primaryFont"`
	PrimaryFontColor        string                                      `json:"primaryFontColor"`
	PrimaryFontLineHeight   string                                      `json:"primaryFontLineHeight"`
	PrimaryFontSize         float64                                     `json:"primaryFontSize"`
	SecondaryAccentColor    string                                      `json:"secondaryAccentColor"`
	SecondaryFont           string                                      `json:"secondaryFont"`
	SecondaryFontColor      string                                      `json:"secondaryFontColor"`
	SecondaryFontLineHeight string                                      `json:"secondaryFontLineHeight"`
	SecondaryFontSize       float64                                     `json:"secondaryFontSize"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BackgroundColor         respjson.Field
		BackgroundImage         respjson.Field
		BackgroundImageType     respjson.Field
		BodyBorderColor         respjson.Field
		BodyBorderColorChoice   respjson.Field
		BodyBorderWidth         respjson.Field
		BodyColor               respjson.Field
		ButtonStyleSettings     respjson.Field
		ColorPickerFavorite1    respjson.Field
		ColorPickerFavorite2    respjson.Field
		ColorPickerFavorite3    respjson.Field
		ColorPickerFavorite4    respjson.Field
		ColorPickerFavorite5    respjson.Field
		ColorPickerFavorite6    respjson.Field
		DividerStyleSettings    respjson.Field
		EmailBodyPadding        respjson.Field
		EmailBodyWidth          respjson.Field
		HeadingOneFont          respjson.Field
		HeadingTwoFont          respjson.Field
		LinksFont               respjson.Field
		PrimaryAccentColor      respjson.Field
		PrimaryFont             respjson.Field
		PrimaryFontColor        respjson.Field
		PrimaryFontLineHeight   respjson.Field
		PrimaryFontSize         respjson.Field
		SecondaryAccentColor    respjson.Field
		SecondaryFont           respjson.Field
		SecondaryFontColor      respjson.Field
		SecondaryFontLineHeight respjson.Field
		SecondaryFontSize       respjson.Field
		ExtraFields             map[string]respjson.Field
		raw                     string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PublicEmailStyleSettings) RawJSON() string { return r.JSON.raw }
func (r *PublicEmailStyleSettings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this PublicEmailStyleSettings to a
// PublicEmailStyleSettingsParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// PublicEmailStyleSettingsParam.Overrides()
func (r PublicEmailStyleSettings) ToParam() PublicEmailStyleSettingsParam {
	return param.Override[PublicEmailStyleSettingsParam](json.RawMessage(r.RawJSON()))
}

type PublicEmailStyleSettingsBackgroundImageType string

const (
	PublicEmailStyleSettingsBackgroundImageTypeRepeat  PublicEmailStyleSettingsBackgroundImageType = "REPEAT"
	PublicEmailStyleSettingsBackgroundImageTypeSingle  PublicEmailStyleSettingsBackgroundImageType = "SINGLE"
	PublicEmailStyleSettingsBackgroundImageTypeStretch PublicEmailStyleSettingsBackgroundImageType = "STRETCH"
)

type PublicEmailStyleSettingsParam struct {
	BackgroundColor         param.Opt[string]  `json:"backgroundColor,omitzero"`
	BackgroundImage         param.Opt[string]  `json:"backgroundImage,omitzero"`
	BodyBorderColor         param.Opt[string]  `json:"bodyBorderColor,omitzero"`
	BodyBorderColorChoice   param.Opt[string]  `json:"bodyBorderColorChoice,omitzero"`
	BodyBorderWidth         param.Opt[float64] `json:"bodyBorderWidth,omitzero"`
	BodyColor               param.Opt[string]  `json:"bodyColor,omitzero"`
	ColorPickerFavorite1    param.Opt[string]  `json:"colorPickerFavorite1,omitzero"`
	ColorPickerFavorite2    param.Opt[string]  `json:"colorPickerFavorite2,omitzero"`
	ColorPickerFavorite3    param.Opt[string]  `json:"colorPickerFavorite3,omitzero"`
	ColorPickerFavorite4    param.Opt[string]  `json:"colorPickerFavorite4,omitzero"`
	ColorPickerFavorite5    param.Opt[string]  `json:"colorPickerFavorite5,omitzero"`
	ColorPickerFavorite6    param.Opt[string]  `json:"colorPickerFavorite6,omitzero"`
	EmailBodyPadding        param.Opt[string]  `json:"emailBodyPadding,omitzero"`
	EmailBodyWidth          param.Opt[string]  `json:"emailBodyWidth,omitzero"`
	PrimaryAccentColor      param.Opt[string]  `json:"primaryAccentColor,omitzero"`
	PrimaryFont             param.Opt[string]  `json:"primaryFont,omitzero"`
	PrimaryFontColor        param.Opt[string]  `json:"primaryFontColor,omitzero"`
	PrimaryFontLineHeight   param.Opt[string]  `json:"primaryFontLineHeight,omitzero"`
	PrimaryFontSize         param.Opt[float64] `json:"primaryFontSize,omitzero"`
	SecondaryAccentColor    param.Opt[string]  `json:"secondaryAccentColor,omitzero"`
	SecondaryFont           param.Opt[string]  `json:"secondaryFont,omitzero"`
	SecondaryFontColor      param.Opt[string]  `json:"secondaryFontColor,omitzero"`
	SecondaryFontLineHeight param.Opt[string]  `json:"secondaryFontLineHeight,omitzero"`
	SecondaryFontSize       param.Opt[float64] `json:"secondaryFontSize,omitzero"`
	// Any of "REPEAT", "SINGLE", "STRETCH".
	BackgroundImageType  PublicEmailStyleSettingsBackgroundImageType `json:"backgroundImageType,omitzero"`
	ButtonStyleSettings  PublicButtonStyleSettingsParam              `json:"buttonStyleSettings,omitzero"`
	DividerStyleSettings PublicDividerStyleSettingsParam             `json:"dividerStyleSettings,omitzero"`
	HeadingOneFont       PublicFontStyleParam                        `json:"headingOneFont,omitzero"`
	HeadingTwoFont       PublicFontStyleParam                        `json:"headingTwoFont,omitzero"`
	LinksFont            PublicFontStyleParam                        `json:"linksFont,omitzero"`
	paramObj
}

func (r PublicEmailStyleSettingsParam) MarshalJSON() (data []byte, err error) {
	type shadow PublicEmailStyleSettingsParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PublicEmailStyleSettingsParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PublicEmailSubscriptionDetails struct {
	// ID of the selected office location.
	OfficeLocationID   string `json:"officeLocationId"`
	PreferencesGroupID string `json:"preferencesGroupId"`
	// ID of the subscription.
	SubscriptionID   string `json:"subscriptionId"`
	SubscriptionName string `json:"subscriptionName"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		OfficeLocationID   respjson.Field
		PreferencesGroupID respjson.Field
		SubscriptionID     respjson.Field
		SubscriptionName   respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PublicEmailSubscriptionDetails) RawJSON() string { return r.JSON.raw }
func (r *PublicEmailSubscriptionDetails) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this PublicEmailSubscriptionDetails to a
// PublicEmailSubscriptionDetailsParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// PublicEmailSubscriptionDetailsParam.Overrides()
func (r PublicEmailSubscriptionDetails) ToParam() PublicEmailSubscriptionDetailsParam {
	return param.Override[PublicEmailSubscriptionDetailsParam](json.RawMessage(r.RawJSON()))
}

type PublicEmailSubscriptionDetailsParam struct {
	// ID of the selected office location.
	OfficeLocationID   param.Opt[string] `json:"officeLocationId,omitzero"`
	PreferencesGroupID param.Opt[string] `json:"preferencesGroupId,omitzero"`
	// ID of the subscription.
	SubscriptionID   param.Opt[string] `json:"subscriptionId,omitzero"`
	SubscriptionName param.Opt[string] `json:"subscriptionName,omitzero"`
	paramObj
}

func (r PublicEmailSubscriptionDetailsParam) MarshalJSON() (data []byte, err error) {
	type shadow PublicEmailSubscriptionDetailsParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PublicEmailSubscriptionDetailsParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PublicEmailTestingDetails struct {
	IsAbVariation bool `json:"isAbVariation" api:"required"`
	// Version of the email that should be sent if there are too few recipients to
	// conduct an AB test.
	//
	// Any of "automated_loser_variant", "automated_master", "automated_variant",
	// "loser_variant", "mab_master", "mab_variant", "master", "variant".
	AbSampleSizeDefault PublicEmailTestingDetailsAbSampleSizeDefault `json:"abSampleSizeDefault"`
	// Version of the email that should be sent if the results are inconclusive after
	// the test period, master or variant.
	//
	// Any of "automated_loser_variant", "automated_master", "automated_variant",
	// "loser_variant", "mab_master", "mab_variant", "master", "variant".
	AbSamplingDefault PublicEmailTestingDetailsAbSamplingDefault `json:"abSamplingDefault"`
	// Status of the AB test.
	//
	// Any of "automated_loser_variant", "automated_master", "automated_variant",
	// "loser_variant", "mab_master", "mab_variant", "master", "variant".
	AbStatus PublicEmailTestingDetailsAbStatus `json:"abStatus"`
	// Metric to determine the version that will be sent to the remaining contacts.
	//
	// Any of "CLICKS_BY_DELIVERED", "CLICKS_BY_OPENS", "OPENS_BY_DELIVERED".
	AbSuccessMetric PublicEmailTestingDetailsAbSuccessMetric `json:"abSuccessMetric"`
	// The size of your test group.
	AbTestPercentage int64 `json:"abTestPercentage"`
	// Time limit on gathering test results. After this time is up, the winning version
	// will be sent to the remaining contacts.
	HoursToWait int64 `json:"hoursToWait"`
	// The ID of the AB test.
	TestID string `json:"testId"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		IsAbVariation       respjson.Field
		AbSampleSizeDefault respjson.Field
		AbSamplingDefault   respjson.Field
		AbStatus            respjson.Field
		AbSuccessMetric     respjson.Field
		AbTestPercentage    respjson.Field
		HoursToWait         respjson.Field
		TestID              respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PublicEmailTestingDetails) RawJSON() string { return r.JSON.raw }
func (r *PublicEmailTestingDetails) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this PublicEmailTestingDetails to a
// PublicEmailTestingDetailsParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// PublicEmailTestingDetailsParam.Overrides()
func (r PublicEmailTestingDetails) ToParam() PublicEmailTestingDetailsParam {
	return param.Override[PublicEmailTestingDetailsParam](json.RawMessage(r.RawJSON()))
}

// Version of the email that should be sent if there are too few recipients to
// conduct an AB test.
type PublicEmailTestingDetailsAbSampleSizeDefault string

const (
	PublicEmailTestingDetailsAbSampleSizeDefaultAutomatedLoserVariant PublicEmailTestingDetailsAbSampleSizeDefault = "automated_loser_variant"
	PublicEmailTestingDetailsAbSampleSizeDefaultAutomatedMaster       PublicEmailTestingDetailsAbSampleSizeDefault = "automated_master"
	PublicEmailTestingDetailsAbSampleSizeDefaultAutomatedVariant      PublicEmailTestingDetailsAbSampleSizeDefault = "automated_variant"
	PublicEmailTestingDetailsAbSampleSizeDefaultLoserVariant          PublicEmailTestingDetailsAbSampleSizeDefault = "loser_variant"
	PublicEmailTestingDetailsAbSampleSizeDefaultMabMaster             PublicEmailTestingDetailsAbSampleSizeDefault = "mab_master"
	PublicEmailTestingDetailsAbSampleSizeDefaultMabVariant            PublicEmailTestingDetailsAbSampleSizeDefault = "mab_variant"
	PublicEmailTestingDetailsAbSampleSizeDefaultMaster                PublicEmailTestingDetailsAbSampleSizeDefault = "master"
	PublicEmailTestingDetailsAbSampleSizeDefaultVariant               PublicEmailTestingDetailsAbSampleSizeDefault = "variant"
)

// Version of the email that should be sent if the results are inconclusive after
// the test period, master or variant.
type PublicEmailTestingDetailsAbSamplingDefault string

const (
	PublicEmailTestingDetailsAbSamplingDefaultAutomatedLoserVariant PublicEmailTestingDetailsAbSamplingDefault = "automated_loser_variant"
	PublicEmailTestingDetailsAbSamplingDefaultAutomatedMaster       PublicEmailTestingDetailsAbSamplingDefault = "automated_master"
	PublicEmailTestingDetailsAbSamplingDefaultAutomatedVariant      PublicEmailTestingDetailsAbSamplingDefault = "automated_variant"
	PublicEmailTestingDetailsAbSamplingDefaultLoserVariant          PublicEmailTestingDetailsAbSamplingDefault = "loser_variant"
	PublicEmailTestingDetailsAbSamplingDefaultMabMaster             PublicEmailTestingDetailsAbSamplingDefault = "mab_master"
	PublicEmailTestingDetailsAbSamplingDefaultMabVariant            PublicEmailTestingDetailsAbSamplingDefault = "mab_variant"
	PublicEmailTestingDetailsAbSamplingDefaultMaster                PublicEmailTestingDetailsAbSamplingDefault = "master"
	PublicEmailTestingDetailsAbSamplingDefaultVariant               PublicEmailTestingDetailsAbSamplingDefault = "variant"
)

// Status of the AB test.
type PublicEmailTestingDetailsAbStatus string

const (
	PublicEmailTestingDetailsAbStatusAutomatedLoserVariant PublicEmailTestingDetailsAbStatus = "automated_loser_variant"
	PublicEmailTestingDetailsAbStatusAutomatedMaster       PublicEmailTestingDetailsAbStatus = "automated_master"
	PublicEmailTestingDetailsAbStatusAutomatedVariant      PublicEmailTestingDetailsAbStatus = "automated_variant"
	PublicEmailTestingDetailsAbStatusLoserVariant          PublicEmailTestingDetailsAbStatus = "loser_variant"
	PublicEmailTestingDetailsAbStatusMabMaster             PublicEmailTestingDetailsAbStatus = "mab_master"
	PublicEmailTestingDetailsAbStatusMabVariant            PublicEmailTestingDetailsAbStatus = "mab_variant"
	PublicEmailTestingDetailsAbStatusMaster                PublicEmailTestingDetailsAbStatus = "master"
	PublicEmailTestingDetailsAbStatusVariant               PublicEmailTestingDetailsAbStatus = "variant"
)

// Metric to determine the version that will be sent to the remaining contacts.
type PublicEmailTestingDetailsAbSuccessMetric string

const (
	PublicEmailTestingDetailsAbSuccessMetricClicksByDelivered PublicEmailTestingDetailsAbSuccessMetric = "CLICKS_BY_DELIVERED"
	PublicEmailTestingDetailsAbSuccessMetricClicksByOpens     PublicEmailTestingDetailsAbSuccessMetric = "CLICKS_BY_OPENS"
	PublicEmailTestingDetailsAbSuccessMetricOpensByDelivered  PublicEmailTestingDetailsAbSuccessMetric = "OPENS_BY_DELIVERED"
)

// The property IsAbVariation is required.
type PublicEmailTestingDetailsParam struct {
	IsAbVariation bool `json:"isAbVariation" api:"required"`
	// The size of your test group.
	AbTestPercentage param.Opt[int64] `json:"abTestPercentage,omitzero"`
	// Time limit on gathering test results. After this time is up, the winning version
	// will be sent to the remaining contacts.
	HoursToWait param.Opt[int64] `json:"hoursToWait,omitzero"`
	// The ID of the AB test.
	TestID param.Opt[string] `json:"testId,omitzero"`
	// Version of the email that should be sent if there are too few recipients to
	// conduct an AB test.
	//
	// Any of "automated_loser_variant", "automated_master", "automated_variant",
	// "loser_variant", "mab_master", "mab_variant", "master", "variant".
	AbSampleSizeDefault PublicEmailTestingDetailsAbSampleSizeDefault `json:"abSampleSizeDefault,omitzero"`
	// Version of the email that should be sent if the results are inconclusive after
	// the test period, master or variant.
	//
	// Any of "automated_loser_variant", "automated_master", "automated_variant",
	// "loser_variant", "mab_master", "mab_variant", "master", "variant".
	AbSamplingDefault PublicEmailTestingDetailsAbSamplingDefault `json:"abSamplingDefault,omitzero"`
	// Status of the AB test.
	//
	// Any of "automated_loser_variant", "automated_master", "automated_variant",
	// "loser_variant", "mab_master", "mab_variant", "master", "variant".
	AbStatus PublicEmailTestingDetailsAbStatus `json:"abStatus,omitzero"`
	// Metric to determine the version that will be sent to the remaining contacts.
	//
	// Any of "CLICKS_BY_DELIVERED", "CLICKS_BY_OPENS", "OPENS_BY_DELIVERED".
	AbSuccessMetric PublicEmailTestingDetailsAbSuccessMetric `json:"abSuccessMetric,omitzero"`
	paramObj
}

func (r PublicEmailTestingDetailsParam) MarshalJSON() (data []byte, err error) {
	type shadow PublicEmailTestingDetailsParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PublicEmailTestingDetailsParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PublicEmailToDetails struct {
	ContactIDs         PublicEmailRecipients `json:"contactIds"`
	ContactIlsLists    PublicEmailRecipients `json:"contactIlsLists"`
	ContactLists       PublicEmailRecipients `json:"contactLists"`
	LimitSendFrequency bool                  `json:"limitSendFrequency"`
	// Whether to send to unengaged contacts (false) or not (true).
	SuppressGraymail bool `json:"suppressGraymail"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ContactIDs         respjson.Field
		ContactIlsLists    respjson.Field
		ContactLists       respjson.Field
		LimitSendFrequency respjson.Field
		SuppressGraymail   respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PublicEmailToDetails) RawJSON() string { return r.JSON.raw }
func (r *PublicEmailToDetails) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this PublicEmailToDetails to a PublicEmailToDetailsParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// PublicEmailToDetailsParam.Overrides()
func (r PublicEmailToDetails) ToParam() PublicEmailToDetailsParam {
	return param.Override[PublicEmailToDetailsParam](json.RawMessage(r.RawJSON()))
}

type PublicEmailToDetailsParam struct {
	LimitSendFrequency param.Opt[bool] `json:"limitSendFrequency,omitzero"`
	// Whether to send to unengaged contacts (false) or not (true).
	SuppressGraymail param.Opt[bool]            `json:"suppressGraymail,omitzero"`
	ContactIDs       PublicEmailRecipientsParam `json:"contactIds,omitzero"`
	ContactIlsLists  PublicEmailRecipientsParam `json:"contactIlsLists,omitzero"`
	ContactLists     PublicEmailRecipientsParam `json:"contactLists,omitzero"`
	paramObj
}

func (r PublicEmailToDetailsParam) MarshalJSON() (data []byte, err error) {
	type shadow PublicEmailToDetailsParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PublicEmailToDetailsParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PublicEmailVersion struct {
	ID        string             `json:"id" api:"required"`
	Object    PublicEmail        `json:"object" api:"required"`
	UpdatedAt time.Time          `json:"updatedAt" api:"required" format:"date-time"`
	User      shared.VersionUser `json:"user" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Object      respjson.Field
		UpdatedAt   respjson.Field
		User        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PublicEmailVersion) RawJSON() string { return r.JSON.raw }
func (r *PublicEmailVersion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PublicFontStyle struct {
	Bold      bool   `json:"bold"`
	Color     string `json:"color"`
	Font      string `json:"font"`
	Italic    bool   `json:"italic"`
	Size      int64  `json:"size"`
	Underline bool   `json:"underline"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Bold        respjson.Field
		Color       respjson.Field
		Font        respjson.Field
		Italic      respjson.Field
		Size        respjson.Field
		Underline   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PublicFontStyle) RawJSON() string { return r.JSON.raw }
func (r *PublicFontStyle) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this PublicFontStyle to a PublicFontStyleParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// PublicFontStyleParam.Overrides()
func (r PublicFontStyle) ToParam() PublicFontStyleParam {
	return param.Override[PublicFontStyleParam](json.RawMessage(r.RawJSON()))
}

type PublicFontStyleParam struct {
	Bold      param.Opt[bool]   `json:"bold,omitzero"`
	Color     param.Opt[string] `json:"color,omitzero"`
	Font      param.Opt[string] `json:"font,omitzero"`
	Italic    param.Opt[bool]   `json:"italic,omitzero"`
	Size      param.Opt[int64]  `json:"size,omitzero"`
	Underline param.Opt[bool]   `json:"underline,omitzero"`
	paramObj
}

func (r PublicFontStyleParam) MarshalJSON() (data []byte, err error) {
	type shadow PublicFontStyleParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PublicFontStyleParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PublicRssEmailDetails struct {
	BlogEmailType     string `json:"blogEmailType"`
	BlogImageMaxWidth int64  `json:"blogImageMaxWidth"`
	// Any of "FULL_POST", "SUMMARY_NO_FEATURED_IMAGE", "SUMMARY_WITH_FEATURED_IMAGE".
	BlogLayout           PublicRssEmailDetailsBlogLayout `json:"blogLayout"`
	HubSpotBlogID        string                          `json:"hubspotBlogId"`
	MaxEntries           int64                           `json:"maxEntries"`
	RssEntryTemplate     string                          `json:"rssEntryTemplate"`
	Timing               map[string]any                  `json:"timing"`
	URL                  string                          `json:"url"`
	UseHeadlineAsSubject bool                            `json:"useHeadlineAsSubject"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BlogEmailType        respjson.Field
		BlogImageMaxWidth    respjson.Field
		BlogLayout           respjson.Field
		HubSpotBlogID        respjson.Field
		MaxEntries           respjson.Field
		RssEntryTemplate     respjson.Field
		Timing               respjson.Field
		URL                  respjson.Field
		UseHeadlineAsSubject respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PublicRssEmailDetails) RawJSON() string { return r.JSON.raw }
func (r *PublicRssEmailDetails) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this PublicRssEmailDetails to a PublicRssEmailDetailsParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// PublicRssEmailDetailsParam.Overrides()
func (r PublicRssEmailDetails) ToParam() PublicRssEmailDetailsParam {
	return param.Override[PublicRssEmailDetailsParam](json.RawMessage(r.RawJSON()))
}

type PublicRssEmailDetailsBlogLayout string

const (
	PublicRssEmailDetailsBlogLayoutFullPost                 PublicRssEmailDetailsBlogLayout = "FULL_POST"
	PublicRssEmailDetailsBlogLayoutSummaryNoFeaturedImage   PublicRssEmailDetailsBlogLayout = "SUMMARY_NO_FEATURED_IMAGE"
	PublicRssEmailDetailsBlogLayoutSummaryWithFeaturedImage PublicRssEmailDetailsBlogLayout = "SUMMARY_WITH_FEATURED_IMAGE"
)

type PublicRssEmailDetailsParam struct {
	BlogEmailType        param.Opt[string] `json:"blogEmailType,omitzero"`
	BlogImageMaxWidth    param.Opt[int64]  `json:"blogImageMaxWidth,omitzero"`
	HubSpotBlogID        param.Opt[string] `json:"hubspotBlogId,omitzero"`
	MaxEntries           param.Opt[int64]  `json:"maxEntries,omitzero"`
	RssEntryTemplate     param.Opt[string] `json:"rssEntryTemplate,omitzero"`
	URL                  param.Opt[string] `json:"url,omitzero"`
	UseHeadlineAsSubject param.Opt[bool]   `json:"useHeadlineAsSubject,omitzero"`
	// Any of "FULL_POST", "SUMMARY_NO_FEATURED_IMAGE", "SUMMARY_WITH_FEATURED_IMAGE".
	BlogLayout PublicRssEmailDetailsBlogLayout `json:"blogLayout,omitzero"`
	Timing     map[string]any                  `json:"timing,omitzero"`
	paramObj
}

func (r PublicRssEmailDetailsParam) MarshalJSON() (data []byte, err error) {
	type shadow PublicRssEmailDetailsParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PublicRssEmailDetailsParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PublicWebversionDetails struct {
	Domain            string    `json:"domain"`
	Enabled           bool      `json:"enabled"`
	ExpiresAt         time.Time `json:"expiresAt" format:"date-time"`
	IsPageRedirected  bool      `json:"isPageRedirected"`
	MetaDescription   string    `json:"metaDescription"`
	PageExpiryEnabled bool      `json:"pageExpiryEnabled"`
	RedirectToPageID  string    `json:"redirectToPageId"`
	RedirectToURL     string    `json:"redirectToUrl"`
	Slug              string    `json:"slug"`
	Title             string    `json:"title"`
	URL               string    `json:"url"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Domain            respjson.Field
		Enabled           respjson.Field
		ExpiresAt         respjson.Field
		IsPageRedirected  respjson.Field
		MetaDescription   respjson.Field
		PageExpiryEnabled respjson.Field
		RedirectToPageID  respjson.Field
		RedirectToURL     respjson.Field
		Slug              respjson.Field
		Title             respjson.Field
		URL               respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PublicWebversionDetails) RawJSON() string { return r.JSON.raw }
func (r *PublicWebversionDetails) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this PublicWebversionDetails to a PublicWebversionDetailsParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// PublicWebversionDetailsParam.Overrides()
func (r PublicWebversionDetails) ToParam() PublicWebversionDetailsParam {
	return param.Override[PublicWebversionDetailsParam](json.RawMessage(r.RawJSON()))
}

type PublicWebversionDetailsParam struct {
	Domain            param.Opt[string]    `json:"domain,omitzero"`
	Enabled           param.Opt[bool]      `json:"enabled,omitzero"`
	ExpiresAt         param.Opt[time.Time] `json:"expiresAt,omitzero" format:"date-time"`
	IsPageRedirected  param.Opt[bool]      `json:"isPageRedirected,omitzero"`
	MetaDescription   param.Opt[string]    `json:"metaDescription,omitzero"`
	PageExpiryEnabled param.Opt[bool]      `json:"pageExpiryEnabled,omitzero"`
	RedirectToPageID  param.Opt[string]    `json:"redirectToPageId,omitzero"`
	RedirectToURL     param.Opt[string]    `json:"redirectToUrl,omitzero"`
	Slug              param.Opt[string]    `json:"slug,omitzero"`
	Title             param.Opt[string]    `json:"title,omitzero"`
	URL               param.Opt[string]    `json:"url,omitzero"`
	paramObj
}

func (r PublicWebversionDetailsParam) MarshalJSON() (data []byte, err error) {
	type shadow PublicWebversionDetailsParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PublicWebversionDetailsParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SmartEmailField = any

type VersionPublicEmail struct {
	// ID of this marketing email version.
	ID     string      `json:"id" api:"required"`
	Object PublicEmail `json:"object" api:"required"`
	// The date and time of the last update to the email, in ISO8601 representation.
	UpdatedAt time.Time          `json:"updatedAt" api:"required" format:"date-time"`
	User      shared.VersionUser `json:"user" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Object      respjson.Field
		UpdatedAt   respjson.Field
		User        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VersionPublicEmail) RawJSON() string { return r.JSON.raw }
func (r *VersionPublicEmail) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmailNewParams struct {
	EmailCreateRequest EmailCreateRequestParam
	paramObj
}

func (r EmailNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.EmailCreateRequest)
}
func (r *EmailNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmailUpdateParams struct {
	EmailUpdateRequest EmailUpdateRequestParam
	// Whether to return only results that have been archived.
	Archived param.Opt[bool] `query:"archived,omitzero" json:"-"`
	paramObj
}

func (r EmailUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.EmailUpdateRequest)
}
func (r *EmailUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [EmailUpdateParams]'s query parameters as `url.Values`.
func (r EmailUpdateParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type EmailListParams struct {
	// The paging cursor token of the last successfully read resource will be returned
	// as the `paging.next.after` JSON property of a paged response containing more
	// results.
	After param.Opt[string] `query:"after,omitzero" json:"-"`
	// Whether to return only results that have been archived.
	Archived      param.Opt[bool]      `query:"archived,omitzero" json:"-"`
	Campaign      param.Opt[string]    `query:"campaign,omitzero" json:"-"`
	CreatedAfter  param.Opt[time.Time] `query:"createdAfter,omitzero" format:"date-time" json:"-"`
	CreatedAt     param.Opt[time.Time] `query:"createdAt,omitzero" format:"date-time" json:"-"`
	CreatedBefore param.Opt[time.Time] `query:"createdBefore,omitzero" format:"date-time" json:"-"`
	IncludeStats  param.Opt[bool]      `query:"includeStats,omitzero" json:"-"`
	IsPublished   param.Opt[bool]      `query:"isPublished,omitzero" json:"-"`
	// The maximum number of results to display per page.
	Limit                  param.Opt[int64]     `query:"limit,omitzero" json:"-"`
	MarketingCampaignNames param.Opt[bool]      `query:"marketingCampaignNames,omitzero" json:"-"`
	PublishedAfter         param.Opt[time.Time] `query:"publishedAfter,omitzero" format:"date-time" json:"-"`
	PublishedAt            param.Opt[time.Time] `query:"publishedAt,omitzero" format:"date-time" json:"-"`
	PublishedBefore        param.Opt[time.Time] `query:"publishedBefore,omitzero" format:"date-time" json:"-"`
	UpdatedAfter           param.Opt[time.Time] `query:"updatedAfter,omitzero" format:"date-time" json:"-"`
	UpdatedAt              param.Opt[time.Time] `query:"updatedAt,omitzero" format:"date-time" json:"-"`
	UpdatedBefore          param.Opt[time.Time] `query:"updatedBefore,omitzero" format:"date-time" json:"-"`
	VariantStats           param.Opt[bool]      `query:"variantStats,omitzero" json:"-"`
	WorkflowNames          param.Opt[bool]      `query:"workflowNames,omitzero" json:"-"`
	IncludedProperties     []string             `query:"includedProperties,omitzero" json:"-"`
	Sort                   []string             `query:"sort,omitzero" json:"-"`
	// Any of "AB_EMAIL", "AUTOMATED_AB_EMAIL", "AUTOMATED_EMAIL", "BATCH_EMAIL",
	// "BLOG_EMAIL", "BLOG_EMAIL_CHILD", "FEEDBACK_CES_EMAIL", "FEEDBACK_CUSTOM_EMAIL",
	// "FEEDBACK_CUSTOM_SURVEY_EMAIL", "FEEDBACK_NPS_EMAIL", "FOLLOWUP_EMAIL",
	// "LEADFLOW_EMAIL", "LOCALTIME_EMAIL", "MANAGE_PREFERENCES_EMAIL",
	// "MARKETING_SINGLE_SEND_API", "MEMBERSHIP_EMAIL_VERIFICATION_EMAIL",
	// "MEMBERSHIP_FOLLOW_UP_EMAIL", "MEMBERSHIP_OTP_LOGIN_EMAIL",
	// "MEMBERSHIP_PASSWORD_RESET_EMAIL", "MEMBERSHIP_PASSWORD_SAVED_EMAIL",
	// "MEMBERSHIP_PASSWORDLESS_AUTH_EMAIL", "MEMBERSHIP_REGISTRATION_EMAIL",
	// "MEMBERSHIP_REGISTRATION_FOLLOW_UP_EMAIL", "MEMBERSHIP_VERIFICATION_EMAIL",
	// "OPTIN_EMAIL", "OPTIN_FOLLOWUP_EMAIL", "RESUBSCRIBE_EMAIL", "RSS_EMAIL",
	// "RSS_EMAIL_CHILD", "SINGLE_SEND_API", "SMTP_TOKEN", "TICKET_EMAIL".
	Type EmailListParamsType `query:"type,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [EmailListParams]'s query parameters as `url.Values`.
func (r EmailListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type EmailListParamsType string

const (
	EmailListParamsTypeAbEmail                             EmailListParamsType = "AB_EMAIL"
	EmailListParamsTypeAutomatedAbEmail                    EmailListParamsType = "AUTOMATED_AB_EMAIL"
	EmailListParamsTypeAutomatedEmail                      EmailListParamsType = "AUTOMATED_EMAIL"
	EmailListParamsTypeBatchEmail                          EmailListParamsType = "BATCH_EMAIL"
	EmailListParamsTypeBlogEmail                           EmailListParamsType = "BLOG_EMAIL"
	EmailListParamsTypeBlogEmailChild                      EmailListParamsType = "BLOG_EMAIL_CHILD"
	EmailListParamsTypeFeedbackCesEmail                    EmailListParamsType = "FEEDBACK_CES_EMAIL"
	EmailListParamsTypeFeedbackCustomEmail                 EmailListParamsType = "FEEDBACK_CUSTOM_EMAIL"
	EmailListParamsTypeFeedbackCustomSurveyEmail           EmailListParamsType = "FEEDBACK_CUSTOM_SURVEY_EMAIL"
	EmailListParamsTypeFeedbackNpsEmail                    EmailListParamsType = "FEEDBACK_NPS_EMAIL"
	EmailListParamsTypeFollowupEmail                       EmailListParamsType = "FOLLOWUP_EMAIL"
	EmailListParamsTypeLeadflowEmail                       EmailListParamsType = "LEADFLOW_EMAIL"
	EmailListParamsTypeLocaltimeEmail                      EmailListParamsType = "LOCALTIME_EMAIL"
	EmailListParamsTypeManagePreferencesEmail              EmailListParamsType = "MANAGE_PREFERENCES_EMAIL"
	EmailListParamsTypeMarketingSingleSendAPI              EmailListParamsType = "MARKETING_SINGLE_SEND_API"
	EmailListParamsTypeMembershipEmailVerificationEmail    EmailListParamsType = "MEMBERSHIP_EMAIL_VERIFICATION_EMAIL"
	EmailListParamsTypeMembershipFollowUpEmail             EmailListParamsType = "MEMBERSHIP_FOLLOW_UP_EMAIL"
	EmailListParamsTypeMembershipOtpLoginEmail             EmailListParamsType = "MEMBERSHIP_OTP_LOGIN_EMAIL"
	EmailListParamsTypeMembershipPasswordResetEmail        EmailListParamsType = "MEMBERSHIP_PASSWORD_RESET_EMAIL"
	EmailListParamsTypeMembershipPasswordSavedEmail        EmailListParamsType = "MEMBERSHIP_PASSWORD_SAVED_EMAIL"
	EmailListParamsTypeMembershipPasswordlessAuthEmail     EmailListParamsType = "MEMBERSHIP_PASSWORDLESS_AUTH_EMAIL"
	EmailListParamsTypeMembershipRegistrationEmail         EmailListParamsType = "MEMBERSHIP_REGISTRATION_EMAIL"
	EmailListParamsTypeMembershipRegistrationFollowUpEmail EmailListParamsType = "MEMBERSHIP_REGISTRATION_FOLLOW_UP_EMAIL"
	EmailListParamsTypeMembershipVerificationEmail         EmailListParamsType = "MEMBERSHIP_VERIFICATION_EMAIL"
	EmailListParamsTypeOptinEmail                          EmailListParamsType = "OPTIN_EMAIL"
	EmailListParamsTypeOptinFollowupEmail                  EmailListParamsType = "OPTIN_FOLLOWUP_EMAIL"
	EmailListParamsTypeResubscribeEmail                    EmailListParamsType = "RESUBSCRIBE_EMAIL"
	EmailListParamsTypeRssEmail                            EmailListParamsType = "RSS_EMAIL"
	EmailListParamsTypeRssEmailChild                       EmailListParamsType = "RSS_EMAIL_CHILD"
	EmailListParamsTypeSingleSendAPI                       EmailListParamsType = "SINGLE_SEND_API"
	EmailListParamsTypeSmtpToken                           EmailListParamsType = "SMTP_TOKEN"
	EmailListParamsTypeTicketEmail                         EmailListParamsType = "TICKET_EMAIL"
)

type EmailDeleteParams struct {
	// Whether to return only results that have been archived.
	Archived param.Opt[bool] `query:"archived,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [EmailDeleteParams]'s query parameters as `url.Values`.
func (r EmailDeleteParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type EmailCloneParams struct {
	EmailCloneRequestVNext EmailCloneRequestVNextParam
	paramObj
}

func (r EmailCloneParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.EmailCloneRequestVNext)
}
func (r *EmailCloneParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmailNewAbTestVariationParams struct {
	AbTestCreateRequestVNext shared.AbTestCreateRequestVNextParam
	paramObj
}

func (r EmailNewAbTestVariationParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.AbTestCreateRequestVNext)
}
func (r *EmailNewAbTestVariationParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmailGetParams struct {
	EndTimestamp   param.Opt[time.Time] `query:"endTimestamp,omitzero" format:"date-time" json:"-"`
	Property       param.Opt[string]    `query:"property,omitzero" json:"-"`
	StartTimestamp param.Opt[time.Time] `query:"startTimestamp,omitzero" format:"date-time" json:"-"`
	EmailIDs       []int64              `query:"emailIds,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [EmailGetParams]'s query parameters as `url.Values`.
func (r EmailGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type EmailGetAbTestVariationParams struct {
	// Whether to return only results that have been archived.
	Archived               param.Opt[bool] `query:"archived,omitzero" json:"-"`
	IncludeStats           param.Opt[bool] `query:"includeStats,omitzero" json:"-"`
	MarketingCampaignNames param.Opt[bool] `query:"marketingCampaignNames,omitzero" json:"-"`
	VariantStats           param.Opt[bool] `query:"variantStats,omitzero" json:"-"`
	WorkflowNames          param.Opt[bool] `query:"workflowNames,omitzero" json:"-"`
	IncludedProperties     []string        `query:"includedProperties,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [EmailGetAbTestVariationParams]'s query parameters as
// `url.Values`.
func (r EmailGetAbTestVariationParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type EmailGetHistogramParams struct {
	EndTimestamp   param.Opt[time.Time] `query:"endTimestamp,omitzero" format:"date-time" json:"-"`
	StartTimestamp param.Opt[time.Time] `query:"startTimestamp,omitzero" format:"date-time" json:"-"`
	EmailIDs       []int64              `query:"emailIds,omitzero" json:"-"`
	// Any of "DAY", "HOUR", "MINUTE", "MONTH", "QUARTER", "QUARTER_HOUR", "SECOND",
	// "WEEK", "YEAR".
	Interval EmailGetHistogramParamsInterval `query:"interval,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [EmailGetHistogramParams]'s query parameters as
// `url.Values`.
func (r EmailGetHistogramParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type EmailGetHistogramParamsInterval string

const (
	EmailGetHistogramParamsIntervalDay         EmailGetHistogramParamsInterval = "DAY"
	EmailGetHistogramParamsIntervalHour        EmailGetHistogramParamsInterval = "HOUR"
	EmailGetHistogramParamsIntervalMinute      EmailGetHistogramParamsInterval = "MINUTE"
	EmailGetHistogramParamsIntervalMonth       EmailGetHistogramParamsInterval = "MONTH"
	EmailGetHistogramParamsIntervalQuarter     EmailGetHistogramParamsInterval = "QUARTER"
	EmailGetHistogramParamsIntervalQuarterHour EmailGetHistogramParamsInterval = "QUARTER_HOUR"
	EmailGetHistogramParamsIntervalSecond      EmailGetHistogramParamsInterval = "SECOND"
	EmailGetHistogramParamsIntervalWeek        EmailGetHistogramParamsInterval = "WEEK"
	EmailGetHistogramParamsIntervalYear        EmailGetHistogramParamsInterval = "YEAR"
)

type EmailGetRevisionParams struct {
	EmailID string `path:"emailId" api:"required" json:"-"`
	paramObj
}

type EmailListRevisionsParams struct {
	// The paging cursor token of the last successfully read resource will be returned
	// as the `paging.next.after` JSON property of a paged response containing more
	// results.
	After  param.Opt[string] `query:"after,omitzero" json:"-"`
	Before param.Opt[string] `query:"before,omitzero" json:"-"`
	// The maximum number of results to display per page.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [EmailListRevisionsParams]'s query parameters as
// `url.Values`.
func (r EmailListRevisionsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type EmailRestoreRevisionParams struct {
	EmailID string `path:"emailId" api:"required" json:"-"`
	paramObj
}

type EmailRestoreRevisionToDraftParams struct {
	EmailID string `path:"emailId" api:"required" json:"-"`
	paramObj
}

type EmailUpdateDraftParams struct {
	EmailUpdateRequest EmailUpdateRequestParam
	paramObj
}

func (r EmailUpdateDraftParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.EmailUpdateRequest)
}
func (r *EmailUpdateDraftParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
