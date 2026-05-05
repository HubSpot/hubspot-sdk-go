// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cms

import (
	"encoding/json"
	"time"

	"github.com/HubSpot/hubspot-sdk-go/internal/apijson"
	"github.com/HubSpot/hubspot-sdk-go/option"
	"github.com/HubSpot/hubspot-sdk-go/packages/param"
	"github.com/HubSpot/hubspot-sdk-go/packages/respjson"
)

// CmService contains methods and other services that help with interacting with
// the hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCmService] method instead.
type CmService struct {
	options      []option.RequestOption
	AuditLogs    AuditLogService
	Blogs        BlogService
	Domains      DomainService
	Hubdb        HubdbService
	MediaBridge  MediaBridgeService
	Pages        PageService
	SiteSearch   SiteSearchService
	SourceCode   SourceCodeService
	URLMappings  URLMappingService
	URLRedirects URLRedirectService
}

// NewCmService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewCmService(opts ...option.RequestOption) (r CmService) {
	r = CmService{}
	r.options = opts
	r.AuditLogs = NewAuditLogService(opts...)
	r.Blogs = NewBlogService(opts...)
	r.Domains = NewDomainService(opts...)
	r.Hubdb = NewHubdbService(opts...)
	r.MediaBridge = NewMediaBridgeService(opts...)
	r.Pages = NewPageService(opts...)
	r.SiteSearch = NewSiteSearchService(opts...)
	r.SourceCode = NewSourceCodeService(opts...)
	r.URLMappings = NewURLMappingService(opts...)
	r.URLRedirects = NewURLRedirectService(opts...)
	return
}

type Angle struct {
	// The unit of measurement for the angle.
	//
	// Any of "deg", "grad", "rad", "turn".
	Units AngleUnits `json:"units" api:"required"`
	// The numerical representation of the angle.
	Value float64 `json:"value" api:"required"`
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

// The unit of measurement for the angle.
type AngleUnits string

const (
	AngleUnitsDeg  AngleUnits = "deg"
	AngleUnitsGrad AngleUnits = "grad"
	AngleUnitsRad  AngleUnits = "rad"
	AngleUnitsTurn AngleUnits = "turn"
)

// The properties Units, Value are required.
type AngleParam struct {
	// The unit of measurement for the angle.
	//
	// Any of "deg", "grad", "rad", "turn".
	Units AngleUnits `json:"units,omitzero" api:"required"`
	// The numerical representation of the angle.
	Value float64 `json:"value" api:"required"`
	paramObj
}

func (r AngleParam) MarshalJSON() (data []byte, err error) {
	type shadow AngleParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AngleParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties ID, Language, PrimaryID are required.
type AttachToLangPrimaryRequestVNextParam struct {
	// ID of the object to add to a multi-language group.
	ID string `json:"id" api:"required"`
	// Designated language of the object to add to a multi-language group.
	//
	// Any of "aa", "ab", "ae", "af", "af-na", "af-za", "agq", "agq-cm", "ak", "ak-gh",
	// "am", "am-et", "an", "ann", "ann-ng", "ar", "ar-001", "ar-ae", "ar-bh", "ar-dj",
	// "ar-dz", "ar-eg", "ar-eh", "ar-er", "ar-il", "ar-iq", "ar-jo", "ar-km", "ar-kw",
	// "ar-lb", "ar-ly", "ar-ma", "ar-mr", "ar-om", "ar-ps", "ar-qa", "ar-sa", "ar-sd",
	// "ar-so", "ar-ss", "ar-sy", "ar-td", "ar-tn", "ar-ye", "as", "as-in", "asa",
	// "asa-tz", "ast", "ast-es", "av", "ay", "az", "az-az", "ba", "bas", "bas-cm",
	// "be", "be-by", "bem", "bem-zm", "bez", "bez-tz", "bg", "bg-bg", "bgc", "bgc-in",
	// "bho", "bho-in", "bi", "bm", "bm-ml", "bn", "bn-bd", "bn-in", "bo", "bo-cn",
	// "bo-in", "br", "br-fr", "brx", "brx-in", "bs", "bs-ba", "ca", "ca-ad", "ca-es",
	// "ca-fr", "ca-it", "ccp", "ccp-bd", "ccp-in", "ce", "ce-ru", "ceb", "ceb-ph",
	// "cgg", "cgg-ug", "ch", "chr", "chr-us", "ckb", "ckb-iq", "ckb-ir", "co", "cr",
	// "cs", "cs-cz", "cu", "cu-ru", "cv", "cv-ru", "cy", "cy-gb", "da", "da-dk",
	// "da-gl", "dav", "dav-ke", "de", "de-at", "de-be", "de-ch", "de-de", "de-gr",
	// "de-it", "de-li", "de-lu", "dje", "dje-ne", "doi", "doi-in", "dsb", "dsb-de",
	// "dua", "dua-cm", "dv", "dyo", "dyo-sn", "dz", "dz-bt", "ebu", "ebu-ke", "ee",
	// "ee-gh", "ee-tg", "el", "el-cy", "el-gr", "en", "en-001", "en-150", "en-ae",
	// "en-ag", "en-ai", "en-as", "en-at", "en-au", "en-bb", "en-be", "en-bi", "en-bm",
	// "en-bs", "en-bw", "en-bz", "en-ca", "en-cc", "en-ch", "en-ck", "en-cm", "en-cn",
	// "en-cx", "en-cy", "en-de", "en-dg", "en-dk", "en-dm", "en-ee", "en-eg", "en-er",
	// "en-es", "en-fi", "en-fj", "en-fk", "en-fm", "en-fr", "en-gb", "en-gd", "en-gg",
	// "en-gh", "en-gi", "en-gm", "en-gu", "en-gy", "en-hk", "en-id", "en-ie", "en-il",
	// "en-im", "en-in", "en-io", "en-je", "en-jm", "en-ke", "en-ki", "en-kn", "en-ky",
	// "en-lc", "en-lr", "en-ls", "en-lu", "en-mg", "en-mh", "en-mo", "en-mp", "en-ms",
	// "en-mt", "en-mu", "en-mv", "en-mw", "en-mx", "en-my", "en-na", "en-nf", "en-ng",
	// "en-nl", "en-nr", "en-nu", "en-nz", "en-pg", "en-ph", "en-pk", "en-pn", "en-pr",
	// "en-pt", "en-pw", "en-rw", "en-sb", "en-sc", "en-sd", "en-se", "en-sg", "en-sh",
	// "en-si", "en-sl", "en-ss", "en-sx", "en-sz", "en-tc", "en-th", "en-tk", "en-tn",
	// "en-to", "en-tt", "en-tv", "en-tz", "en-ug", "en-um", "en-us", "en-vc", "en-vg",
	// "en-vi", "en-vn", "en-vu", "en-ws", "en-za", "en-zm", "en-zw", "eo", "eo-001",
	// "es", "es-419", "es-ar", "es-bo", "es-br", "es-bz", "es-cl", "es-co", "es-cr",
	// "es-cu", "es-do", "es-ea", "es-ec", "es-es", "es-gq", "es-gt", "es-hn", "es-ic",
	// "es-mx", "es-ni", "es-pa", "es-pe", "es-ph", "es-pr", "es-py", "es-sv", "es-us",
	// "es-uy", "es-ve", "et", "et-ee", "eu", "eu-es", "ewo", "ewo-cm", "fa", "fa-af",
	// "fa-ir", "ff", "ff-bf", "ff-cm", "ff-gh", "ff-gm", "ff-gn", "ff-gw", "ff-lr",
	// "ff-mr", "ff-ne", "ff-ng", "ff-sl", "ff-sn", "fi", "fi-fi", "fil", "fil-ph",
	// "fj", "fo", "fo-dk", "fo-fo", "fr", "fr-be", "fr-bf", "fr-bi", "fr-bj", "fr-bl",
	// "fr-ca", "fr-cd", "fr-cf", "fr-cg", "fr-ch", "fr-ci", "fr-cm", "fr-dj", "fr-dz",
	// "fr-fr", "fr-ga", "fr-gf", "fr-gn", "fr-gp", "fr-gq", "fr-ht", "fr-km", "fr-lu",
	// "fr-ma", "fr-mc", "fr-mf", "fr-mg", "fr-ml", "fr-mq", "fr-mr", "fr-mu", "fr-nc",
	// "fr-ne", "fr-pf", "fr-pm", "fr-re", "fr-rw", "fr-sc", "fr-sn", "fr-sy", "fr-td",
	// "fr-tg", "fr-tn", "fr-vu", "fr-wf", "fr-yt", "frr", "frr-de", "fur", "fur-it",
	// "fy", "fy-nl", "ga", "ga-gb", "ga-ie", "gd", "gd-gb", "gl", "gl-es", "gn",
	// "gsw", "gsw-ch", "gsw-fr", "gsw-li", "gu", "gu-in", "guz", "guz-ke", "gv",
	// "gv-im", "ha", "ha-gh", "ha-ne", "ha-ng", "haw", "haw-us", "he", "he-il", "hi",
	// "hi-in", "hmn", "ho", "hr", "hr-ba", "hr-hr", "hsb", "hsb-de", "ht", "hu",
	// "hu-hu", "hy", "hy-am", "hz", "ia", "ia-001", "id", "id-id", "ie", "ig",
	// "ig-ng", "ii", "ii-cn", "ik", "io", "is", "is-is", "it", "it-ch", "it-it",
	// "it-sm", "it-va", "iu", "ja", "ja-jp", "jgo", "jgo-cm", "jmc", "jmc-tz", "jv",
	// "jv-id", "ka", "ka-ge", "kab", "kab-dz", "kam", "kam-ke", "kar", "kde",
	// "kde-tz", "kea", "kea-cv", "kg", "kgp", "kgp-br", "kh", "khq", "khq-ml", "ki",
	// "ki-ke", "kj", "kk", "kk-kz", "kkj", "kkj-cm", "kl", "kl-gl", "kln", "kln-ke",
	// "km", "km-kh", "kn", "kn-in", "ko", "ko-kp", "ko-kr", "kok", "kok-in", "kr",
	// "ks", "ks-in", "ksb", "ksb-tz", "ksf", "ksf-cm", "ksh", "ksh-de", "ku", "ku-tr",
	// "kv", "kw", "kw-gb", "ky", "ky-kg", "la", "lag", "lag-tz", "lb", "lb-lu", "lg",
	// "lg-ug", "li", "lkt", "lkt-us", "ln", "ln-ao", "ln-cd", "ln-cf", "ln-cg", "lo",
	// "lo-la", "lrc", "lrc-iq", "lrc-ir", "lt", "lt-lt", "lu", "lu-cd", "luo",
	// "luo-ke", "luy", "luy-ke", "lv", "lv-lv", "mai", "mai-in", "mas", "mas-ke",
	// "mas-tz", "mdf", "mdf-ru", "mer", "mer-ke", "mfe", "mfe-mu", "mg", "mg-mg",
	// "mgh", "mgh-mz", "mgo", "mgo-cm", "mh", "mi", "mi-nz", "mk", "mk-mk", "ml",
	// "ml-in", "mn", "mn-mn", "mni", "mni-in", "mr", "mr-in", "ms", "ms-bn", "ms-id",
	// "ms-my", "ms-sg", "mt", "mt-mt", "mua", "mua-cm", "my", "my-mm", "mzn",
	// "mzn-ir", "na", "naq", "naq-na", "nb", "nb-no", "nb-sj", "nd", "nd-zw", "nds",
	// "nds-de", "nds-nl", "ne", "ne-in", "ne-np", "ng", "nl", "nl-aw", "nl-be",
	// "nl-bq", "nl-ch", "nl-cw", "nl-lu", "nl-nl", "nl-sr", "nl-sx", "nmg", "nmg-cm",
	// "nn", "nn-no", "nnh", "nnh-cm", "no", "no-no", "nr", "nus", "nus-ss", "nv",
	// "ny", "nyn", "nyn-ug", "oc", "oc-es", "oc-fr", "oj", "om", "om-et", "om-ke",
	// "or", "or-in", "os", "os-ge", "os-ru", "pa", "pa-in", "pa-pk", "pcm", "pcm-ng",
	// "pi", "pis", "pis-sb", "pl", "pl-pl", "prg", "prg-001", "ps", "ps-af", "ps-pk",
	// "pt", "pt-ao", "pt-br", "pt-ch", "pt-cv", "pt-gq", "pt-gw", "pt-lu", "pt-mo",
	// "pt-mz", "pt-pt", "pt-st", "pt-tl", "qu", "qu-bo", "qu-ec", "qu-pe", "raj",
	// "raj-in", "rm", "rm-ch", "rn", "rn-bi", "ro", "ro-md", "ro-ro", "rof", "rof-tz",
	// "ru", "ru-by", "ru-kg", "ru-kz", "ru-md", "ru-ru", "ru-ua", "rw", "rw-rw",
	// "rwk", "rwk-tz", "sa", "sa-in", "sah", "sah-ru", "saq", "saq-ke", "sat",
	// "sat-in", "sbp", "sbp-tz", "sc", "sc-it", "sd", "sd-in", "sd-pk", "se", "se-fi",
	// "se-no", "se-se", "seh", "seh-mz", "ses", "ses-ml", "sg", "sg-cf", "shi",
	// "shi-ma", "si", "si-lk", "sk", "sk-sk", "sl", "sl-si", "sm", "smn", "smn-fi",
	// "sms", "sms-fi", "sn", "sn-zw", "so", "so-dj", "so-et", "so-ke", "so-so", "sq",
	// "sq-al", "sq-mk", "sq-xk", "sr", "sr-ba", "sr-cs", "sr-me", "sr-rs", "sr-xk",
	// "ss", "st", "su", "su-id", "sv", "sv-ax", "sv-fi", "sv-se", "sw", "sw-cd",
	// "sw-ke", "sw-tz", "sw-ug", "sy", "ta", "ta-in", "ta-lk", "ta-my", "ta-sg", "te",
	// "te-in", "teo", "teo-ke", "teo-ug", "tg", "tg-tj", "th", "th-th", "ti", "ti-er",
	// "ti-et", "tk", "tk-tm", "tl", "tn", "to", "to-to", "tok", "tok-001", "tr",
	// "tr-cy", "tr-tr", "ts", "tt", "tt-ru", "tw", "twq", "twq-ne", "ty", "tzm",
	// "tzm-ma", "ug", "ug-cn", "uk", "uk-ua", "ur", "ur-in", "ur-pk", "uz", "uz-af",
	// "uz-uz", "vai", "vai-lr", "ve", "vi", "vi-vn", "vo", "vo-001", "vun", "vun-tz",
	// "wa", "wae", "wae-ch", "wo", "wo-sn", "xh", "xh-za", "xog", "xog-ug", "yav",
	// "yav-cm", "yi", "yi-001", "yo", "yo-bj", "yo-ng", "yrl", "yrl-br", "yrl-co",
	// "yrl-ve", "yue", "yue-cn", "yue-hk", "za", "zgh", "zgh-ma", "zh", "zh-cn",
	// "zh-hans", "zh-hant", "zh-hk", "zh-mo", "zh-sg", "zh-tw", "zu", "zu-za".
	Language AttachToLangPrimaryRequestVNextLanguage `json:"language,omitzero" api:"required"`
	// ID of primary language object in multi-language group.
	PrimaryID string `json:"primaryId" api:"required"`
	// Primary language of the multi-language group.
	//
	// Any of "aa", "ab", "ae", "af", "af-na", "af-za", "agq", "agq-cm", "ak", "ak-gh",
	// "am", "am-et", "an", "ann", "ann-ng", "ar", "ar-001", "ar-ae", "ar-bh", "ar-dj",
	// "ar-dz", "ar-eg", "ar-eh", "ar-er", "ar-il", "ar-iq", "ar-jo", "ar-km", "ar-kw",
	// "ar-lb", "ar-ly", "ar-ma", "ar-mr", "ar-om", "ar-ps", "ar-qa", "ar-sa", "ar-sd",
	// "ar-so", "ar-ss", "ar-sy", "ar-td", "ar-tn", "ar-ye", "as", "as-in", "asa",
	// "asa-tz", "ast", "ast-es", "av", "ay", "az", "az-az", "ba", "bas", "bas-cm",
	// "be", "be-by", "bem", "bem-zm", "bez", "bez-tz", "bg", "bg-bg", "bgc", "bgc-in",
	// "bho", "bho-in", "bi", "bm", "bm-ml", "bn", "bn-bd", "bn-in", "bo", "bo-cn",
	// "bo-in", "br", "br-fr", "brx", "brx-in", "bs", "bs-ba", "ca", "ca-ad", "ca-es",
	// "ca-fr", "ca-it", "ccp", "ccp-bd", "ccp-in", "ce", "ce-ru", "ceb", "ceb-ph",
	// "cgg", "cgg-ug", "ch", "chr", "chr-us", "ckb", "ckb-iq", "ckb-ir", "co", "cr",
	// "cs", "cs-cz", "cu", "cu-ru", "cv", "cv-ru", "cy", "cy-gb", "da", "da-dk",
	// "da-gl", "dav", "dav-ke", "de", "de-at", "de-be", "de-ch", "de-de", "de-gr",
	// "de-it", "de-li", "de-lu", "dje", "dje-ne", "doi", "doi-in", "dsb", "dsb-de",
	// "dua", "dua-cm", "dv", "dyo", "dyo-sn", "dz", "dz-bt", "ebu", "ebu-ke", "ee",
	// "ee-gh", "ee-tg", "el", "el-cy", "el-gr", "en", "en-001", "en-150", "en-ae",
	// "en-ag", "en-ai", "en-as", "en-at", "en-au", "en-bb", "en-be", "en-bi", "en-bm",
	// "en-bs", "en-bw", "en-bz", "en-ca", "en-cc", "en-ch", "en-ck", "en-cm", "en-cn",
	// "en-cx", "en-cy", "en-de", "en-dg", "en-dk", "en-dm", "en-ee", "en-eg", "en-er",
	// "en-es", "en-fi", "en-fj", "en-fk", "en-fm", "en-fr", "en-gb", "en-gd", "en-gg",
	// "en-gh", "en-gi", "en-gm", "en-gu", "en-gy", "en-hk", "en-id", "en-ie", "en-il",
	// "en-im", "en-in", "en-io", "en-je", "en-jm", "en-ke", "en-ki", "en-kn", "en-ky",
	// "en-lc", "en-lr", "en-ls", "en-lu", "en-mg", "en-mh", "en-mo", "en-mp", "en-ms",
	// "en-mt", "en-mu", "en-mv", "en-mw", "en-mx", "en-my", "en-na", "en-nf", "en-ng",
	// "en-nl", "en-nr", "en-nu", "en-nz", "en-pg", "en-ph", "en-pk", "en-pn", "en-pr",
	// "en-pt", "en-pw", "en-rw", "en-sb", "en-sc", "en-sd", "en-se", "en-sg", "en-sh",
	// "en-si", "en-sl", "en-ss", "en-sx", "en-sz", "en-tc", "en-th", "en-tk", "en-tn",
	// "en-to", "en-tt", "en-tv", "en-tz", "en-ug", "en-um", "en-us", "en-vc", "en-vg",
	// "en-vi", "en-vn", "en-vu", "en-ws", "en-za", "en-zm", "en-zw", "eo", "eo-001",
	// "es", "es-419", "es-ar", "es-bo", "es-br", "es-bz", "es-cl", "es-co", "es-cr",
	// "es-cu", "es-do", "es-ea", "es-ec", "es-es", "es-gq", "es-gt", "es-hn", "es-ic",
	// "es-mx", "es-ni", "es-pa", "es-pe", "es-ph", "es-pr", "es-py", "es-sv", "es-us",
	// "es-uy", "es-ve", "et", "et-ee", "eu", "eu-es", "ewo", "ewo-cm", "fa", "fa-af",
	// "fa-ir", "ff", "ff-bf", "ff-cm", "ff-gh", "ff-gm", "ff-gn", "ff-gw", "ff-lr",
	// "ff-mr", "ff-ne", "ff-ng", "ff-sl", "ff-sn", "fi", "fi-fi", "fil", "fil-ph",
	// "fj", "fo", "fo-dk", "fo-fo", "fr", "fr-be", "fr-bf", "fr-bi", "fr-bj", "fr-bl",
	// "fr-ca", "fr-cd", "fr-cf", "fr-cg", "fr-ch", "fr-ci", "fr-cm", "fr-dj", "fr-dz",
	// "fr-fr", "fr-ga", "fr-gf", "fr-gn", "fr-gp", "fr-gq", "fr-ht", "fr-km", "fr-lu",
	// "fr-ma", "fr-mc", "fr-mf", "fr-mg", "fr-ml", "fr-mq", "fr-mr", "fr-mu", "fr-nc",
	// "fr-ne", "fr-pf", "fr-pm", "fr-re", "fr-rw", "fr-sc", "fr-sn", "fr-sy", "fr-td",
	// "fr-tg", "fr-tn", "fr-vu", "fr-wf", "fr-yt", "frr", "frr-de", "fur", "fur-it",
	// "fy", "fy-nl", "ga", "ga-gb", "ga-ie", "gd", "gd-gb", "gl", "gl-es", "gn",
	// "gsw", "gsw-ch", "gsw-fr", "gsw-li", "gu", "gu-in", "guz", "guz-ke", "gv",
	// "gv-im", "ha", "ha-gh", "ha-ne", "ha-ng", "haw", "haw-us", "he", "he-il", "hi",
	// "hi-in", "hmn", "ho", "hr", "hr-ba", "hr-hr", "hsb", "hsb-de", "ht", "hu",
	// "hu-hu", "hy", "hy-am", "hz", "ia", "ia-001", "id", "id-id", "ie", "ig",
	// "ig-ng", "ii", "ii-cn", "ik", "io", "is", "is-is", "it", "it-ch", "it-it",
	// "it-sm", "it-va", "iu", "ja", "ja-jp", "jgo", "jgo-cm", "jmc", "jmc-tz", "jv",
	// "jv-id", "ka", "ka-ge", "kab", "kab-dz", "kam", "kam-ke", "kar", "kde",
	// "kde-tz", "kea", "kea-cv", "kg", "kgp", "kgp-br", "kh", "khq", "khq-ml", "ki",
	// "ki-ke", "kj", "kk", "kk-kz", "kkj", "kkj-cm", "kl", "kl-gl", "kln", "kln-ke",
	// "km", "km-kh", "kn", "kn-in", "ko", "ko-kp", "ko-kr", "kok", "kok-in", "kr",
	// "ks", "ks-in", "ksb", "ksb-tz", "ksf", "ksf-cm", "ksh", "ksh-de", "ku", "ku-tr",
	// "kv", "kw", "kw-gb", "ky", "ky-kg", "la", "lag", "lag-tz", "lb", "lb-lu", "lg",
	// "lg-ug", "li", "lkt", "lkt-us", "ln", "ln-ao", "ln-cd", "ln-cf", "ln-cg", "lo",
	// "lo-la", "lrc", "lrc-iq", "lrc-ir", "lt", "lt-lt", "lu", "lu-cd", "luo",
	// "luo-ke", "luy", "luy-ke", "lv", "lv-lv", "mai", "mai-in", "mas", "mas-ke",
	// "mas-tz", "mdf", "mdf-ru", "mer", "mer-ke", "mfe", "mfe-mu", "mg", "mg-mg",
	// "mgh", "mgh-mz", "mgo", "mgo-cm", "mh", "mi", "mi-nz", "mk", "mk-mk", "ml",
	// "ml-in", "mn", "mn-mn", "mni", "mni-in", "mr", "mr-in", "ms", "ms-bn", "ms-id",
	// "ms-my", "ms-sg", "mt", "mt-mt", "mua", "mua-cm", "my", "my-mm", "mzn",
	// "mzn-ir", "na", "naq", "naq-na", "nb", "nb-no", "nb-sj", "nd", "nd-zw", "nds",
	// "nds-de", "nds-nl", "ne", "ne-in", "ne-np", "ng", "nl", "nl-aw", "nl-be",
	// "nl-bq", "nl-ch", "nl-cw", "nl-lu", "nl-nl", "nl-sr", "nl-sx", "nmg", "nmg-cm",
	// "nn", "nn-no", "nnh", "nnh-cm", "no", "no-no", "nr", "nus", "nus-ss", "nv",
	// "ny", "nyn", "nyn-ug", "oc", "oc-es", "oc-fr", "oj", "om", "om-et", "om-ke",
	// "or", "or-in", "os", "os-ge", "os-ru", "pa", "pa-in", "pa-pk", "pcm", "pcm-ng",
	// "pi", "pis", "pis-sb", "pl", "pl-pl", "prg", "prg-001", "ps", "ps-af", "ps-pk",
	// "pt", "pt-ao", "pt-br", "pt-ch", "pt-cv", "pt-gq", "pt-gw", "pt-lu", "pt-mo",
	// "pt-mz", "pt-pt", "pt-st", "pt-tl", "qu", "qu-bo", "qu-ec", "qu-pe", "raj",
	// "raj-in", "rm", "rm-ch", "rn", "rn-bi", "ro", "ro-md", "ro-ro", "rof", "rof-tz",
	// "ru", "ru-by", "ru-kg", "ru-kz", "ru-md", "ru-ru", "ru-ua", "rw", "rw-rw",
	// "rwk", "rwk-tz", "sa", "sa-in", "sah", "sah-ru", "saq", "saq-ke", "sat",
	// "sat-in", "sbp", "sbp-tz", "sc", "sc-it", "sd", "sd-in", "sd-pk", "se", "se-fi",
	// "se-no", "se-se", "seh", "seh-mz", "ses", "ses-ml", "sg", "sg-cf", "shi",
	// "shi-ma", "si", "si-lk", "sk", "sk-sk", "sl", "sl-si", "sm", "smn", "smn-fi",
	// "sms", "sms-fi", "sn", "sn-zw", "so", "so-dj", "so-et", "so-ke", "so-so", "sq",
	// "sq-al", "sq-mk", "sq-xk", "sr", "sr-ba", "sr-cs", "sr-me", "sr-rs", "sr-xk",
	// "ss", "st", "su", "su-id", "sv", "sv-ax", "sv-fi", "sv-se", "sw", "sw-cd",
	// "sw-ke", "sw-tz", "sw-ug", "sy", "ta", "ta-in", "ta-lk", "ta-my", "ta-sg", "te",
	// "te-in", "teo", "teo-ke", "teo-ug", "tg", "tg-tj", "th", "th-th", "ti", "ti-er",
	// "ti-et", "tk", "tk-tm", "tl", "tn", "to", "to-to", "tok", "tok-001", "tr",
	// "tr-cy", "tr-tr", "ts", "tt", "tt-ru", "tw", "twq", "twq-ne", "ty", "tzm",
	// "tzm-ma", "ug", "ug-cn", "uk", "uk-ua", "ur", "ur-in", "ur-pk", "uz", "uz-af",
	// "uz-uz", "vai", "vai-lr", "ve", "vi", "vi-vn", "vo", "vo-001", "vun", "vun-tz",
	// "wa", "wae", "wae-ch", "wo", "wo-sn", "xh", "xh-za", "xog", "xog-ug", "yav",
	// "yav-cm", "yi", "yi-001", "yo", "yo-bj", "yo-ng", "yrl", "yrl-br", "yrl-co",
	// "yrl-ve", "yue", "yue-cn", "yue-hk", "za", "zgh", "zgh-ma", "zh", "zh-cn",
	// "zh-hans", "zh-hant", "zh-hk", "zh-mo", "zh-sg", "zh-tw", "zu", "zu-za".
	PrimaryLanguage AttachToLangPrimaryRequestVNextPrimaryLanguage `json:"primaryLanguage,omitzero"`
	paramObj
}

func (r AttachToLangPrimaryRequestVNextParam) MarshalJSON() (data []byte, err error) {
	type shadow AttachToLangPrimaryRequestVNextParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AttachToLangPrimaryRequestVNextParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Designated language of the object to add to a multi-language group.
type AttachToLangPrimaryRequestVNextLanguage string

const (
	AttachToLangPrimaryRequestVNextLanguageAa     AttachToLangPrimaryRequestVNextLanguage = "aa"
	AttachToLangPrimaryRequestVNextLanguageAb     AttachToLangPrimaryRequestVNextLanguage = "ab"
	AttachToLangPrimaryRequestVNextLanguageAe     AttachToLangPrimaryRequestVNextLanguage = "ae"
	AttachToLangPrimaryRequestVNextLanguageAf     AttachToLangPrimaryRequestVNextLanguage = "af"
	AttachToLangPrimaryRequestVNextLanguageAfNa   AttachToLangPrimaryRequestVNextLanguage = "af-na"
	AttachToLangPrimaryRequestVNextLanguageAfZa   AttachToLangPrimaryRequestVNextLanguage = "af-za"
	AttachToLangPrimaryRequestVNextLanguageAgq    AttachToLangPrimaryRequestVNextLanguage = "agq"
	AttachToLangPrimaryRequestVNextLanguageAgqCm  AttachToLangPrimaryRequestVNextLanguage = "agq-cm"
	AttachToLangPrimaryRequestVNextLanguageAk     AttachToLangPrimaryRequestVNextLanguage = "ak"
	AttachToLangPrimaryRequestVNextLanguageAkGh   AttachToLangPrimaryRequestVNextLanguage = "ak-gh"
	AttachToLangPrimaryRequestVNextLanguageAm     AttachToLangPrimaryRequestVNextLanguage = "am"
	AttachToLangPrimaryRequestVNextLanguageAmEt   AttachToLangPrimaryRequestVNextLanguage = "am-et"
	AttachToLangPrimaryRequestVNextLanguageAn     AttachToLangPrimaryRequestVNextLanguage = "an"
	AttachToLangPrimaryRequestVNextLanguageAnn    AttachToLangPrimaryRequestVNextLanguage = "ann"
	AttachToLangPrimaryRequestVNextLanguageAnnNg  AttachToLangPrimaryRequestVNextLanguage = "ann-ng"
	AttachToLangPrimaryRequestVNextLanguageAr     AttachToLangPrimaryRequestVNextLanguage = "ar"
	AttachToLangPrimaryRequestVNextLanguageAr001  AttachToLangPrimaryRequestVNextLanguage = "ar-001"
	AttachToLangPrimaryRequestVNextLanguageArAe   AttachToLangPrimaryRequestVNextLanguage = "ar-ae"
	AttachToLangPrimaryRequestVNextLanguageArBh   AttachToLangPrimaryRequestVNextLanguage = "ar-bh"
	AttachToLangPrimaryRequestVNextLanguageArDj   AttachToLangPrimaryRequestVNextLanguage = "ar-dj"
	AttachToLangPrimaryRequestVNextLanguageArDz   AttachToLangPrimaryRequestVNextLanguage = "ar-dz"
	AttachToLangPrimaryRequestVNextLanguageArEg   AttachToLangPrimaryRequestVNextLanguage = "ar-eg"
	AttachToLangPrimaryRequestVNextLanguageArEh   AttachToLangPrimaryRequestVNextLanguage = "ar-eh"
	AttachToLangPrimaryRequestVNextLanguageArEr   AttachToLangPrimaryRequestVNextLanguage = "ar-er"
	AttachToLangPrimaryRequestVNextLanguageArIl   AttachToLangPrimaryRequestVNextLanguage = "ar-il"
	AttachToLangPrimaryRequestVNextLanguageArIq   AttachToLangPrimaryRequestVNextLanguage = "ar-iq"
	AttachToLangPrimaryRequestVNextLanguageArJo   AttachToLangPrimaryRequestVNextLanguage = "ar-jo"
	AttachToLangPrimaryRequestVNextLanguageArKm   AttachToLangPrimaryRequestVNextLanguage = "ar-km"
	AttachToLangPrimaryRequestVNextLanguageArKw   AttachToLangPrimaryRequestVNextLanguage = "ar-kw"
	AttachToLangPrimaryRequestVNextLanguageArLb   AttachToLangPrimaryRequestVNextLanguage = "ar-lb"
	AttachToLangPrimaryRequestVNextLanguageArLy   AttachToLangPrimaryRequestVNextLanguage = "ar-ly"
	AttachToLangPrimaryRequestVNextLanguageArMa   AttachToLangPrimaryRequestVNextLanguage = "ar-ma"
	AttachToLangPrimaryRequestVNextLanguageArMr   AttachToLangPrimaryRequestVNextLanguage = "ar-mr"
	AttachToLangPrimaryRequestVNextLanguageArOm   AttachToLangPrimaryRequestVNextLanguage = "ar-om"
	AttachToLangPrimaryRequestVNextLanguageArPs   AttachToLangPrimaryRequestVNextLanguage = "ar-ps"
	AttachToLangPrimaryRequestVNextLanguageArQa   AttachToLangPrimaryRequestVNextLanguage = "ar-qa"
	AttachToLangPrimaryRequestVNextLanguageArSa   AttachToLangPrimaryRequestVNextLanguage = "ar-sa"
	AttachToLangPrimaryRequestVNextLanguageArSd   AttachToLangPrimaryRequestVNextLanguage = "ar-sd"
	AttachToLangPrimaryRequestVNextLanguageArSo   AttachToLangPrimaryRequestVNextLanguage = "ar-so"
	AttachToLangPrimaryRequestVNextLanguageArSS   AttachToLangPrimaryRequestVNextLanguage = "ar-ss"
	AttachToLangPrimaryRequestVNextLanguageArSy   AttachToLangPrimaryRequestVNextLanguage = "ar-sy"
	AttachToLangPrimaryRequestVNextLanguageArTd   AttachToLangPrimaryRequestVNextLanguage = "ar-td"
	AttachToLangPrimaryRequestVNextLanguageArTn   AttachToLangPrimaryRequestVNextLanguage = "ar-tn"
	AttachToLangPrimaryRequestVNextLanguageArYe   AttachToLangPrimaryRequestVNextLanguage = "ar-ye"
	AttachToLangPrimaryRequestVNextLanguageAs     AttachToLangPrimaryRequestVNextLanguage = "as"
	AttachToLangPrimaryRequestVNextLanguageAsIn   AttachToLangPrimaryRequestVNextLanguage = "as-in"
	AttachToLangPrimaryRequestVNextLanguageAsa    AttachToLangPrimaryRequestVNextLanguage = "asa"
	AttachToLangPrimaryRequestVNextLanguageAsaTz  AttachToLangPrimaryRequestVNextLanguage = "asa-tz"
	AttachToLangPrimaryRequestVNextLanguageAst    AttachToLangPrimaryRequestVNextLanguage = "ast"
	AttachToLangPrimaryRequestVNextLanguageAstEs  AttachToLangPrimaryRequestVNextLanguage = "ast-es"
	AttachToLangPrimaryRequestVNextLanguageAv     AttachToLangPrimaryRequestVNextLanguage = "av"
	AttachToLangPrimaryRequestVNextLanguageAy     AttachToLangPrimaryRequestVNextLanguage = "ay"
	AttachToLangPrimaryRequestVNextLanguageAz     AttachToLangPrimaryRequestVNextLanguage = "az"
	AttachToLangPrimaryRequestVNextLanguageAzAz   AttachToLangPrimaryRequestVNextLanguage = "az-az"
	AttachToLangPrimaryRequestVNextLanguageBa     AttachToLangPrimaryRequestVNextLanguage = "ba"
	AttachToLangPrimaryRequestVNextLanguageBas    AttachToLangPrimaryRequestVNextLanguage = "bas"
	AttachToLangPrimaryRequestVNextLanguageBasCm  AttachToLangPrimaryRequestVNextLanguage = "bas-cm"
	AttachToLangPrimaryRequestVNextLanguageBe     AttachToLangPrimaryRequestVNextLanguage = "be"
	AttachToLangPrimaryRequestVNextLanguageBeBy   AttachToLangPrimaryRequestVNextLanguage = "be-by"
	AttachToLangPrimaryRequestVNextLanguageBem    AttachToLangPrimaryRequestVNextLanguage = "bem"
	AttachToLangPrimaryRequestVNextLanguageBemZm  AttachToLangPrimaryRequestVNextLanguage = "bem-zm"
	AttachToLangPrimaryRequestVNextLanguageBez    AttachToLangPrimaryRequestVNextLanguage = "bez"
	AttachToLangPrimaryRequestVNextLanguageBezTz  AttachToLangPrimaryRequestVNextLanguage = "bez-tz"
	AttachToLangPrimaryRequestVNextLanguageBg     AttachToLangPrimaryRequestVNextLanguage = "bg"
	AttachToLangPrimaryRequestVNextLanguageBgBg   AttachToLangPrimaryRequestVNextLanguage = "bg-bg"
	AttachToLangPrimaryRequestVNextLanguageBgc    AttachToLangPrimaryRequestVNextLanguage = "bgc"
	AttachToLangPrimaryRequestVNextLanguageBgcIn  AttachToLangPrimaryRequestVNextLanguage = "bgc-in"
	AttachToLangPrimaryRequestVNextLanguageBho    AttachToLangPrimaryRequestVNextLanguage = "bho"
	AttachToLangPrimaryRequestVNextLanguageBhoIn  AttachToLangPrimaryRequestVNextLanguage = "bho-in"
	AttachToLangPrimaryRequestVNextLanguageBi     AttachToLangPrimaryRequestVNextLanguage = "bi"
	AttachToLangPrimaryRequestVNextLanguageBm     AttachToLangPrimaryRequestVNextLanguage = "bm"
	AttachToLangPrimaryRequestVNextLanguageBmMl   AttachToLangPrimaryRequestVNextLanguage = "bm-ml"
	AttachToLangPrimaryRequestVNextLanguageBn     AttachToLangPrimaryRequestVNextLanguage = "bn"
	AttachToLangPrimaryRequestVNextLanguageBnBd   AttachToLangPrimaryRequestVNextLanguage = "bn-bd"
	AttachToLangPrimaryRequestVNextLanguageBnIn   AttachToLangPrimaryRequestVNextLanguage = "bn-in"
	AttachToLangPrimaryRequestVNextLanguageBo     AttachToLangPrimaryRequestVNextLanguage = "bo"
	AttachToLangPrimaryRequestVNextLanguageBoCn   AttachToLangPrimaryRequestVNextLanguage = "bo-cn"
	AttachToLangPrimaryRequestVNextLanguageBoIn   AttachToLangPrimaryRequestVNextLanguage = "bo-in"
	AttachToLangPrimaryRequestVNextLanguageBr     AttachToLangPrimaryRequestVNextLanguage = "br"
	AttachToLangPrimaryRequestVNextLanguageBrFr   AttachToLangPrimaryRequestVNextLanguage = "br-fr"
	AttachToLangPrimaryRequestVNextLanguageBrx    AttachToLangPrimaryRequestVNextLanguage = "brx"
	AttachToLangPrimaryRequestVNextLanguageBrxIn  AttachToLangPrimaryRequestVNextLanguage = "brx-in"
	AttachToLangPrimaryRequestVNextLanguageBs     AttachToLangPrimaryRequestVNextLanguage = "bs"
	AttachToLangPrimaryRequestVNextLanguageBsBa   AttachToLangPrimaryRequestVNextLanguage = "bs-ba"
	AttachToLangPrimaryRequestVNextLanguageCa     AttachToLangPrimaryRequestVNextLanguage = "ca"
	AttachToLangPrimaryRequestVNextLanguageCaAd   AttachToLangPrimaryRequestVNextLanguage = "ca-ad"
	AttachToLangPrimaryRequestVNextLanguageCaEs   AttachToLangPrimaryRequestVNextLanguage = "ca-es"
	AttachToLangPrimaryRequestVNextLanguageCaFr   AttachToLangPrimaryRequestVNextLanguage = "ca-fr"
	AttachToLangPrimaryRequestVNextLanguageCaIt   AttachToLangPrimaryRequestVNextLanguage = "ca-it"
	AttachToLangPrimaryRequestVNextLanguageCcp    AttachToLangPrimaryRequestVNextLanguage = "ccp"
	AttachToLangPrimaryRequestVNextLanguageCcpBd  AttachToLangPrimaryRequestVNextLanguage = "ccp-bd"
	AttachToLangPrimaryRequestVNextLanguageCcpIn  AttachToLangPrimaryRequestVNextLanguage = "ccp-in"
	AttachToLangPrimaryRequestVNextLanguageCe     AttachToLangPrimaryRequestVNextLanguage = "ce"
	AttachToLangPrimaryRequestVNextLanguageCeRu   AttachToLangPrimaryRequestVNextLanguage = "ce-ru"
	AttachToLangPrimaryRequestVNextLanguageCeb    AttachToLangPrimaryRequestVNextLanguage = "ceb"
	AttachToLangPrimaryRequestVNextLanguageCebPh  AttachToLangPrimaryRequestVNextLanguage = "ceb-ph"
	AttachToLangPrimaryRequestVNextLanguageCgg    AttachToLangPrimaryRequestVNextLanguage = "cgg"
	AttachToLangPrimaryRequestVNextLanguageCggUg  AttachToLangPrimaryRequestVNextLanguage = "cgg-ug"
	AttachToLangPrimaryRequestVNextLanguageCh     AttachToLangPrimaryRequestVNextLanguage = "ch"
	AttachToLangPrimaryRequestVNextLanguageChr    AttachToLangPrimaryRequestVNextLanguage = "chr"
	AttachToLangPrimaryRequestVNextLanguageChrUs  AttachToLangPrimaryRequestVNextLanguage = "chr-us"
	AttachToLangPrimaryRequestVNextLanguageCkb    AttachToLangPrimaryRequestVNextLanguage = "ckb"
	AttachToLangPrimaryRequestVNextLanguageCkbIq  AttachToLangPrimaryRequestVNextLanguage = "ckb-iq"
	AttachToLangPrimaryRequestVNextLanguageCkbIr  AttachToLangPrimaryRequestVNextLanguage = "ckb-ir"
	AttachToLangPrimaryRequestVNextLanguageCo     AttachToLangPrimaryRequestVNextLanguage = "co"
	AttachToLangPrimaryRequestVNextLanguageCr     AttachToLangPrimaryRequestVNextLanguage = "cr"
	AttachToLangPrimaryRequestVNextLanguageCs     AttachToLangPrimaryRequestVNextLanguage = "cs"
	AttachToLangPrimaryRequestVNextLanguageCsCz   AttachToLangPrimaryRequestVNextLanguage = "cs-cz"
	AttachToLangPrimaryRequestVNextLanguageCu     AttachToLangPrimaryRequestVNextLanguage = "cu"
	AttachToLangPrimaryRequestVNextLanguageCuRu   AttachToLangPrimaryRequestVNextLanguage = "cu-ru"
	AttachToLangPrimaryRequestVNextLanguageCv     AttachToLangPrimaryRequestVNextLanguage = "cv"
	AttachToLangPrimaryRequestVNextLanguageCvRu   AttachToLangPrimaryRequestVNextLanguage = "cv-ru"
	AttachToLangPrimaryRequestVNextLanguageCy     AttachToLangPrimaryRequestVNextLanguage = "cy"
	AttachToLangPrimaryRequestVNextLanguageCyGB   AttachToLangPrimaryRequestVNextLanguage = "cy-gb"
	AttachToLangPrimaryRequestVNextLanguageDa     AttachToLangPrimaryRequestVNextLanguage = "da"
	AttachToLangPrimaryRequestVNextLanguageDaDk   AttachToLangPrimaryRequestVNextLanguage = "da-dk"
	AttachToLangPrimaryRequestVNextLanguageDaGl   AttachToLangPrimaryRequestVNextLanguage = "da-gl"
	AttachToLangPrimaryRequestVNextLanguageDav    AttachToLangPrimaryRequestVNextLanguage = "dav"
	AttachToLangPrimaryRequestVNextLanguageDavKe  AttachToLangPrimaryRequestVNextLanguage = "dav-ke"
	AttachToLangPrimaryRequestVNextLanguageDe     AttachToLangPrimaryRequestVNextLanguage = "de"
	AttachToLangPrimaryRequestVNextLanguageDeAt   AttachToLangPrimaryRequestVNextLanguage = "de-at"
	AttachToLangPrimaryRequestVNextLanguageDeBe   AttachToLangPrimaryRequestVNextLanguage = "de-be"
	AttachToLangPrimaryRequestVNextLanguageDeCh   AttachToLangPrimaryRequestVNextLanguage = "de-ch"
	AttachToLangPrimaryRequestVNextLanguageDeDe   AttachToLangPrimaryRequestVNextLanguage = "de-de"
	AttachToLangPrimaryRequestVNextLanguageDeGr   AttachToLangPrimaryRequestVNextLanguage = "de-gr"
	AttachToLangPrimaryRequestVNextLanguageDeIt   AttachToLangPrimaryRequestVNextLanguage = "de-it"
	AttachToLangPrimaryRequestVNextLanguageDeLi   AttachToLangPrimaryRequestVNextLanguage = "de-li"
	AttachToLangPrimaryRequestVNextLanguageDeLu   AttachToLangPrimaryRequestVNextLanguage = "de-lu"
	AttachToLangPrimaryRequestVNextLanguageDje    AttachToLangPrimaryRequestVNextLanguage = "dje"
	AttachToLangPrimaryRequestVNextLanguageDjeNe  AttachToLangPrimaryRequestVNextLanguage = "dje-ne"
	AttachToLangPrimaryRequestVNextLanguageDoi    AttachToLangPrimaryRequestVNextLanguage = "doi"
	AttachToLangPrimaryRequestVNextLanguageDoiIn  AttachToLangPrimaryRequestVNextLanguage = "doi-in"
	AttachToLangPrimaryRequestVNextLanguageDsb    AttachToLangPrimaryRequestVNextLanguage = "dsb"
	AttachToLangPrimaryRequestVNextLanguageDsbDe  AttachToLangPrimaryRequestVNextLanguage = "dsb-de"
	AttachToLangPrimaryRequestVNextLanguageDua    AttachToLangPrimaryRequestVNextLanguage = "dua"
	AttachToLangPrimaryRequestVNextLanguageDuaCm  AttachToLangPrimaryRequestVNextLanguage = "dua-cm"
	AttachToLangPrimaryRequestVNextLanguageDv     AttachToLangPrimaryRequestVNextLanguage = "dv"
	AttachToLangPrimaryRequestVNextLanguageDyo    AttachToLangPrimaryRequestVNextLanguage = "dyo"
	AttachToLangPrimaryRequestVNextLanguageDyoSn  AttachToLangPrimaryRequestVNextLanguage = "dyo-sn"
	AttachToLangPrimaryRequestVNextLanguageDz     AttachToLangPrimaryRequestVNextLanguage = "dz"
	AttachToLangPrimaryRequestVNextLanguageDzBt   AttachToLangPrimaryRequestVNextLanguage = "dz-bt"
	AttachToLangPrimaryRequestVNextLanguageEbu    AttachToLangPrimaryRequestVNextLanguage = "ebu"
	AttachToLangPrimaryRequestVNextLanguageEbuKe  AttachToLangPrimaryRequestVNextLanguage = "ebu-ke"
	AttachToLangPrimaryRequestVNextLanguageEe     AttachToLangPrimaryRequestVNextLanguage = "ee"
	AttachToLangPrimaryRequestVNextLanguageEeGh   AttachToLangPrimaryRequestVNextLanguage = "ee-gh"
	AttachToLangPrimaryRequestVNextLanguageEeTg   AttachToLangPrimaryRequestVNextLanguage = "ee-tg"
	AttachToLangPrimaryRequestVNextLanguageEl     AttachToLangPrimaryRequestVNextLanguage = "el"
	AttachToLangPrimaryRequestVNextLanguageElCy   AttachToLangPrimaryRequestVNextLanguage = "el-cy"
	AttachToLangPrimaryRequestVNextLanguageElGr   AttachToLangPrimaryRequestVNextLanguage = "el-gr"
	AttachToLangPrimaryRequestVNextLanguageEn     AttachToLangPrimaryRequestVNextLanguage = "en"
	AttachToLangPrimaryRequestVNextLanguageEn001  AttachToLangPrimaryRequestVNextLanguage = "en-001"
	AttachToLangPrimaryRequestVNextLanguageEn150  AttachToLangPrimaryRequestVNextLanguage = "en-150"
	AttachToLangPrimaryRequestVNextLanguageEnAe   AttachToLangPrimaryRequestVNextLanguage = "en-ae"
	AttachToLangPrimaryRequestVNextLanguageEnAg   AttachToLangPrimaryRequestVNextLanguage = "en-ag"
	AttachToLangPrimaryRequestVNextLanguageEnAI   AttachToLangPrimaryRequestVNextLanguage = "en-ai"
	AttachToLangPrimaryRequestVNextLanguageEnAs   AttachToLangPrimaryRequestVNextLanguage = "en-as"
	AttachToLangPrimaryRequestVNextLanguageEnAt   AttachToLangPrimaryRequestVNextLanguage = "en-at"
	AttachToLangPrimaryRequestVNextLanguageEnAu   AttachToLangPrimaryRequestVNextLanguage = "en-au"
	AttachToLangPrimaryRequestVNextLanguageEnBb   AttachToLangPrimaryRequestVNextLanguage = "en-bb"
	AttachToLangPrimaryRequestVNextLanguageEnBe   AttachToLangPrimaryRequestVNextLanguage = "en-be"
	AttachToLangPrimaryRequestVNextLanguageEnBi   AttachToLangPrimaryRequestVNextLanguage = "en-bi"
	AttachToLangPrimaryRequestVNextLanguageEnBm   AttachToLangPrimaryRequestVNextLanguage = "en-bm"
	AttachToLangPrimaryRequestVNextLanguageEnBs   AttachToLangPrimaryRequestVNextLanguage = "en-bs"
	AttachToLangPrimaryRequestVNextLanguageEnBw   AttachToLangPrimaryRequestVNextLanguage = "en-bw"
	AttachToLangPrimaryRequestVNextLanguageEnBz   AttachToLangPrimaryRequestVNextLanguage = "en-bz"
	AttachToLangPrimaryRequestVNextLanguageEnCa   AttachToLangPrimaryRequestVNextLanguage = "en-ca"
	AttachToLangPrimaryRequestVNextLanguageEnCc   AttachToLangPrimaryRequestVNextLanguage = "en-cc"
	AttachToLangPrimaryRequestVNextLanguageEnCh   AttachToLangPrimaryRequestVNextLanguage = "en-ch"
	AttachToLangPrimaryRequestVNextLanguageEnCk   AttachToLangPrimaryRequestVNextLanguage = "en-ck"
	AttachToLangPrimaryRequestVNextLanguageEnCm   AttachToLangPrimaryRequestVNextLanguage = "en-cm"
	AttachToLangPrimaryRequestVNextLanguageEnCn   AttachToLangPrimaryRequestVNextLanguage = "en-cn"
	AttachToLangPrimaryRequestVNextLanguageEnCx   AttachToLangPrimaryRequestVNextLanguage = "en-cx"
	AttachToLangPrimaryRequestVNextLanguageEnCy   AttachToLangPrimaryRequestVNextLanguage = "en-cy"
	AttachToLangPrimaryRequestVNextLanguageEnDe   AttachToLangPrimaryRequestVNextLanguage = "en-de"
	AttachToLangPrimaryRequestVNextLanguageEnDg   AttachToLangPrimaryRequestVNextLanguage = "en-dg"
	AttachToLangPrimaryRequestVNextLanguageEnDk   AttachToLangPrimaryRequestVNextLanguage = "en-dk"
	AttachToLangPrimaryRequestVNextLanguageEnDm   AttachToLangPrimaryRequestVNextLanguage = "en-dm"
	AttachToLangPrimaryRequestVNextLanguageEnEe   AttachToLangPrimaryRequestVNextLanguage = "en-ee"
	AttachToLangPrimaryRequestVNextLanguageEnEg   AttachToLangPrimaryRequestVNextLanguage = "en-eg"
	AttachToLangPrimaryRequestVNextLanguageEnEr   AttachToLangPrimaryRequestVNextLanguage = "en-er"
	AttachToLangPrimaryRequestVNextLanguageEnEs   AttachToLangPrimaryRequestVNextLanguage = "en-es"
	AttachToLangPrimaryRequestVNextLanguageEnFi   AttachToLangPrimaryRequestVNextLanguage = "en-fi"
	AttachToLangPrimaryRequestVNextLanguageEnFj   AttachToLangPrimaryRequestVNextLanguage = "en-fj"
	AttachToLangPrimaryRequestVNextLanguageEnFk   AttachToLangPrimaryRequestVNextLanguage = "en-fk"
	AttachToLangPrimaryRequestVNextLanguageEnFm   AttachToLangPrimaryRequestVNextLanguage = "en-fm"
	AttachToLangPrimaryRequestVNextLanguageEnFr   AttachToLangPrimaryRequestVNextLanguage = "en-fr"
	AttachToLangPrimaryRequestVNextLanguageEnGB   AttachToLangPrimaryRequestVNextLanguage = "en-gb"
	AttachToLangPrimaryRequestVNextLanguageEnGd   AttachToLangPrimaryRequestVNextLanguage = "en-gd"
	AttachToLangPrimaryRequestVNextLanguageEnGg   AttachToLangPrimaryRequestVNextLanguage = "en-gg"
	AttachToLangPrimaryRequestVNextLanguageEnGh   AttachToLangPrimaryRequestVNextLanguage = "en-gh"
	AttachToLangPrimaryRequestVNextLanguageEnGi   AttachToLangPrimaryRequestVNextLanguage = "en-gi"
	AttachToLangPrimaryRequestVNextLanguageEnGm   AttachToLangPrimaryRequestVNextLanguage = "en-gm"
	AttachToLangPrimaryRequestVNextLanguageEnGu   AttachToLangPrimaryRequestVNextLanguage = "en-gu"
	AttachToLangPrimaryRequestVNextLanguageEnGy   AttachToLangPrimaryRequestVNextLanguage = "en-gy"
	AttachToLangPrimaryRequestVNextLanguageEnHk   AttachToLangPrimaryRequestVNextLanguage = "en-hk"
	AttachToLangPrimaryRequestVNextLanguageEnID   AttachToLangPrimaryRequestVNextLanguage = "en-id"
	AttachToLangPrimaryRequestVNextLanguageEnIe   AttachToLangPrimaryRequestVNextLanguage = "en-ie"
	AttachToLangPrimaryRequestVNextLanguageEnIl   AttachToLangPrimaryRequestVNextLanguage = "en-il"
	AttachToLangPrimaryRequestVNextLanguageEnIm   AttachToLangPrimaryRequestVNextLanguage = "en-im"
	AttachToLangPrimaryRequestVNextLanguageEnIn   AttachToLangPrimaryRequestVNextLanguage = "en-in"
	AttachToLangPrimaryRequestVNextLanguageEnIo   AttachToLangPrimaryRequestVNextLanguage = "en-io"
	AttachToLangPrimaryRequestVNextLanguageEnJe   AttachToLangPrimaryRequestVNextLanguage = "en-je"
	AttachToLangPrimaryRequestVNextLanguageEnJm   AttachToLangPrimaryRequestVNextLanguage = "en-jm"
	AttachToLangPrimaryRequestVNextLanguageEnKe   AttachToLangPrimaryRequestVNextLanguage = "en-ke"
	AttachToLangPrimaryRequestVNextLanguageEnKi   AttachToLangPrimaryRequestVNextLanguage = "en-ki"
	AttachToLangPrimaryRequestVNextLanguageEnKn   AttachToLangPrimaryRequestVNextLanguage = "en-kn"
	AttachToLangPrimaryRequestVNextLanguageEnKy   AttachToLangPrimaryRequestVNextLanguage = "en-ky"
	AttachToLangPrimaryRequestVNextLanguageEnLc   AttachToLangPrimaryRequestVNextLanguage = "en-lc"
	AttachToLangPrimaryRequestVNextLanguageEnLr   AttachToLangPrimaryRequestVNextLanguage = "en-lr"
	AttachToLangPrimaryRequestVNextLanguageEnLs   AttachToLangPrimaryRequestVNextLanguage = "en-ls"
	AttachToLangPrimaryRequestVNextLanguageEnLu   AttachToLangPrimaryRequestVNextLanguage = "en-lu"
	AttachToLangPrimaryRequestVNextLanguageEnMg   AttachToLangPrimaryRequestVNextLanguage = "en-mg"
	AttachToLangPrimaryRequestVNextLanguageEnMh   AttachToLangPrimaryRequestVNextLanguage = "en-mh"
	AttachToLangPrimaryRequestVNextLanguageEnMo   AttachToLangPrimaryRequestVNextLanguage = "en-mo"
	AttachToLangPrimaryRequestVNextLanguageEnMp   AttachToLangPrimaryRequestVNextLanguage = "en-mp"
	AttachToLangPrimaryRequestVNextLanguageEnMs   AttachToLangPrimaryRequestVNextLanguage = "en-ms"
	AttachToLangPrimaryRequestVNextLanguageEnMt   AttachToLangPrimaryRequestVNextLanguage = "en-mt"
	AttachToLangPrimaryRequestVNextLanguageEnMu   AttachToLangPrimaryRequestVNextLanguage = "en-mu"
	AttachToLangPrimaryRequestVNextLanguageEnMv   AttachToLangPrimaryRequestVNextLanguage = "en-mv"
	AttachToLangPrimaryRequestVNextLanguageEnMw   AttachToLangPrimaryRequestVNextLanguage = "en-mw"
	AttachToLangPrimaryRequestVNextLanguageEnMx   AttachToLangPrimaryRequestVNextLanguage = "en-mx"
	AttachToLangPrimaryRequestVNextLanguageEnMy   AttachToLangPrimaryRequestVNextLanguage = "en-my"
	AttachToLangPrimaryRequestVNextLanguageEnNa   AttachToLangPrimaryRequestVNextLanguage = "en-na"
	AttachToLangPrimaryRequestVNextLanguageEnNf   AttachToLangPrimaryRequestVNextLanguage = "en-nf"
	AttachToLangPrimaryRequestVNextLanguageEnNg   AttachToLangPrimaryRequestVNextLanguage = "en-ng"
	AttachToLangPrimaryRequestVNextLanguageEnNl   AttachToLangPrimaryRequestVNextLanguage = "en-nl"
	AttachToLangPrimaryRequestVNextLanguageEnNr   AttachToLangPrimaryRequestVNextLanguage = "en-nr"
	AttachToLangPrimaryRequestVNextLanguageEnNu   AttachToLangPrimaryRequestVNextLanguage = "en-nu"
	AttachToLangPrimaryRequestVNextLanguageEnNz   AttachToLangPrimaryRequestVNextLanguage = "en-nz"
	AttachToLangPrimaryRequestVNextLanguageEnPg   AttachToLangPrimaryRequestVNextLanguage = "en-pg"
	AttachToLangPrimaryRequestVNextLanguageEnPh   AttachToLangPrimaryRequestVNextLanguage = "en-ph"
	AttachToLangPrimaryRequestVNextLanguageEnPk   AttachToLangPrimaryRequestVNextLanguage = "en-pk"
	AttachToLangPrimaryRequestVNextLanguageEnPn   AttachToLangPrimaryRequestVNextLanguage = "en-pn"
	AttachToLangPrimaryRequestVNextLanguageEnPr   AttachToLangPrimaryRequestVNextLanguage = "en-pr"
	AttachToLangPrimaryRequestVNextLanguageEnPt   AttachToLangPrimaryRequestVNextLanguage = "en-pt"
	AttachToLangPrimaryRequestVNextLanguageEnPw   AttachToLangPrimaryRequestVNextLanguage = "en-pw"
	AttachToLangPrimaryRequestVNextLanguageEnRw   AttachToLangPrimaryRequestVNextLanguage = "en-rw"
	AttachToLangPrimaryRequestVNextLanguageEnSb   AttachToLangPrimaryRequestVNextLanguage = "en-sb"
	AttachToLangPrimaryRequestVNextLanguageEnSc   AttachToLangPrimaryRequestVNextLanguage = "en-sc"
	AttachToLangPrimaryRequestVNextLanguageEnSd   AttachToLangPrimaryRequestVNextLanguage = "en-sd"
	AttachToLangPrimaryRequestVNextLanguageEnSe   AttachToLangPrimaryRequestVNextLanguage = "en-se"
	AttachToLangPrimaryRequestVNextLanguageEnSg   AttachToLangPrimaryRequestVNextLanguage = "en-sg"
	AttachToLangPrimaryRequestVNextLanguageEnSh   AttachToLangPrimaryRequestVNextLanguage = "en-sh"
	AttachToLangPrimaryRequestVNextLanguageEnSi   AttachToLangPrimaryRequestVNextLanguage = "en-si"
	AttachToLangPrimaryRequestVNextLanguageEnSl   AttachToLangPrimaryRequestVNextLanguage = "en-sl"
	AttachToLangPrimaryRequestVNextLanguageEnSS   AttachToLangPrimaryRequestVNextLanguage = "en-ss"
	AttachToLangPrimaryRequestVNextLanguageEnSx   AttachToLangPrimaryRequestVNextLanguage = "en-sx"
	AttachToLangPrimaryRequestVNextLanguageEnSz   AttachToLangPrimaryRequestVNextLanguage = "en-sz"
	AttachToLangPrimaryRequestVNextLanguageEnTc   AttachToLangPrimaryRequestVNextLanguage = "en-tc"
	AttachToLangPrimaryRequestVNextLanguageEnTh   AttachToLangPrimaryRequestVNextLanguage = "en-th"
	AttachToLangPrimaryRequestVNextLanguageEnTk   AttachToLangPrimaryRequestVNextLanguage = "en-tk"
	AttachToLangPrimaryRequestVNextLanguageEnTn   AttachToLangPrimaryRequestVNextLanguage = "en-tn"
	AttachToLangPrimaryRequestVNextLanguageEnTo   AttachToLangPrimaryRequestVNextLanguage = "en-to"
	AttachToLangPrimaryRequestVNextLanguageEnTt   AttachToLangPrimaryRequestVNextLanguage = "en-tt"
	AttachToLangPrimaryRequestVNextLanguageEnTv   AttachToLangPrimaryRequestVNextLanguage = "en-tv"
	AttachToLangPrimaryRequestVNextLanguageEnTz   AttachToLangPrimaryRequestVNextLanguage = "en-tz"
	AttachToLangPrimaryRequestVNextLanguageEnUg   AttachToLangPrimaryRequestVNextLanguage = "en-ug"
	AttachToLangPrimaryRequestVNextLanguageEnUm   AttachToLangPrimaryRequestVNextLanguage = "en-um"
	AttachToLangPrimaryRequestVNextLanguageEnUs   AttachToLangPrimaryRequestVNextLanguage = "en-us"
	AttachToLangPrimaryRequestVNextLanguageEnVc   AttachToLangPrimaryRequestVNextLanguage = "en-vc"
	AttachToLangPrimaryRequestVNextLanguageEnVg   AttachToLangPrimaryRequestVNextLanguage = "en-vg"
	AttachToLangPrimaryRequestVNextLanguageEnVi   AttachToLangPrimaryRequestVNextLanguage = "en-vi"
	AttachToLangPrimaryRequestVNextLanguageEnVn   AttachToLangPrimaryRequestVNextLanguage = "en-vn"
	AttachToLangPrimaryRequestVNextLanguageEnVu   AttachToLangPrimaryRequestVNextLanguage = "en-vu"
	AttachToLangPrimaryRequestVNextLanguageEnWs   AttachToLangPrimaryRequestVNextLanguage = "en-ws"
	AttachToLangPrimaryRequestVNextLanguageEnZa   AttachToLangPrimaryRequestVNextLanguage = "en-za"
	AttachToLangPrimaryRequestVNextLanguageEnZm   AttachToLangPrimaryRequestVNextLanguage = "en-zm"
	AttachToLangPrimaryRequestVNextLanguageEnZw   AttachToLangPrimaryRequestVNextLanguage = "en-zw"
	AttachToLangPrimaryRequestVNextLanguageEo     AttachToLangPrimaryRequestVNextLanguage = "eo"
	AttachToLangPrimaryRequestVNextLanguageEo001  AttachToLangPrimaryRequestVNextLanguage = "eo-001"
	AttachToLangPrimaryRequestVNextLanguageEs     AttachToLangPrimaryRequestVNextLanguage = "es"
	AttachToLangPrimaryRequestVNextLanguageEs419  AttachToLangPrimaryRequestVNextLanguage = "es-419"
	AttachToLangPrimaryRequestVNextLanguageEsAr   AttachToLangPrimaryRequestVNextLanguage = "es-ar"
	AttachToLangPrimaryRequestVNextLanguageEsBo   AttachToLangPrimaryRequestVNextLanguage = "es-bo"
	AttachToLangPrimaryRequestVNextLanguageEsBr   AttachToLangPrimaryRequestVNextLanguage = "es-br"
	AttachToLangPrimaryRequestVNextLanguageEsBz   AttachToLangPrimaryRequestVNextLanguage = "es-bz"
	AttachToLangPrimaryRequestVNextLanguageEsCl   AttachToLangPrimaryRequestVNextLanguage = "es-cl"
	AttachToLangPrimaryRequestVNextLanguageEsCo   AttachToLangPrimaryRequestVNextLanguage = "es-co"
	AttachToLangPrimaryRequestVNextLanguageEsCr   AttachToLangPrimaryRequestVNextLanguage = "es-cr"
	AttachToLangPrimaryRequestVNextLanguageEsCu   AttachToLangPrimaryRequestVNextLanguage = "es-cu"
	AttachToLangPrimaryRequestVNextLanguageEsDo   AttachToLangPrimaryRequestVNextLanguage = "es-do"
	AttachToLangPrimaryRequestVNextLanguageEsEa   AttachToLangPrimaryRequestVNextLanguage = "es-ea"
	AttachToLangPrimaryRequestVNextLanguageEsEc   AttachToLangPrimaryRequestVNextLanguage = "es-ec"
	AttachToLangPrimaryRequestVNextLanguageEsEs   AttachToLangPrimaryRequestVNextLanguage = "es-es"
	AttachToLangPrimaryRequestVNextLanguageEsGq   AttachToLangPrimaryRequestVNextLanguage = "es-gq"
	AttachToLangPrimaryRequestVNextLanguageEsGt   AttachToLangPrimaryRequestVNextLanguage = "es-gt"
	AttachToLangPrimaryRequestVNextLanguageEsHn   AttachToLangPrimaryRequestVNextLanguage = "es-hn"
	AttachToLangPrimaryRequestVNextLanguageEsIc   AttachToLangPrimaryRequestVNextLanguage = "es-ic"
	AttachToLangPrimaryRequestVNextLanguageEsMx   AttachToLangPrimaryRequestVNextLanguage = "es-mx"
	AttachToLangPrimaryRequestVNextLanguageEsNi   AttachToLangPrimaryRequestVNextLanguage = "es-ni"
	AttachToLangPrimaryRequestVNextLanguageEsPa   AttachToLangPrimaryRequestVNextLanguage = "es-pa"
	AttachToLangPrimaryRequestVNextLanguageEsPe   AttachToLangPrimaryRequestVNextLanguage = "es-pe"
	AttachToLangPrimaryRequestVNextLanguageEsPh   AttachToLangPrimaryRequestVNextLanguage = "es-ph"
	AttachToLangPrimaryRequestVNextLanguageEsPr   AttachToLangPrimaryRequestVNextLanguage = "es-pr"
	AttachToLangPrimaryRequestVNextLanguageEsPy   AttachToLangPrimaryRequestVNextLanguage = "es-py"
	AttachToLangPrimaryRequestVNextLanguageEsSv   AttachToLangPrimaryRequestVNextLanguage = "es-sv"
	AttachToLangPrimaryRequestVNextLanguageEsUs   AttachToLangPrimaryRequestVNextLanguage = "es-us"
	AttachToLangPrimaryRequestVNextLanguageEsUy   AttachToLangPrimaryRequestVNextLanguage = "es-uy"
	AttachToLangPrimaryRequestVNextLanguageEsVe   AttachToLangPrimaryRequestVNextLanguage = "es-ve"
	AttachToLangPrimaryRequestVNextLanguageEt     AttachToLangPrimaryRequestVNextLanguage = "et"
	AttachToLangPrimaryRequestVNextLanguageEtEe   AttachToLangPrimaryRequestVNextLanguage = "et-ee"
	AttachToLangPrimaryRequestVNextLanguageEu     AttachToLangPrimaryRequestVNextLanguage = "eu"
	AttachToLangPrimaryRequestVNextLanguageEuEs   AttachToLangPrimaryRequestVNextLanguage = "eu-es"
	AttachToLangPrimaryRequestVNextLanguageEwo    AttachToLangPrimaryRequestVNextLanguage = "ewo"
	AttachToLangPrimaryRequestVNextLanguageEwoCm  AttachToLangPrimaryRequestVNextLanguage = "ewo-cm"
	AttachToLangPrimaryRequestVNextLanguageFa     AttachToLangPrimaryRequestVNextLanguage = "fa"
	AttachToLangPrimaryRequestVNextLanguageFaAf   AttachToLangPrimaryRequestVNextLanguage = "fa-af"
	AttachToLangPrimaryRequestVNextLanguageFaIr   AttachToLangPrimaryRequestVNextLanguage = "fa-ir"
	AttachToLangPrimaryRequestVNextLanguageFf     AttachToLangPrimaryRequestVNextLanguage = "ff"
	AttachToLangPrimaryRequestVNextLanguageFfBf   AttachToLangPrimaryRequestVNextLanguage = "ff-bf"
	AttachToLangPrimaryRequestVNextLanguageFfCm   AttachToLangPrimaryRequestVNextLanguage = "ff-cm"
	AttachToLangPrimaryRequestVNextLanguageFfGh   AttachToLangPrimaryRequestVNextLanguage = "ff-gh"
	AttachToLangPrimaryRequestVNextLanguageFfGm   AttachToLangPrimaryRequestVNextLanguage = "ff-gm"
	AttachToLangPrimaryRequestVNextLanguageFfGn   AttachToLangPrimaryRequestVNextLanguage = "ff-gn"
	AttachToLangPrimaryRequestVNextLanguageFfGw   AttachToLangPrimaryRequestVNextLanguage = "ff-gw"
	AttachToLangPrimaryRequestVNextLanguageFfLr   AttachToLangPrimaryRequestVNextLanguage = "ff-lr"
	AttachToLangPrimaryRequestVNextLanguageFfMr   AttachToLangPrimaryRequestVNextLanguage = "ff-mr"
	AttachToLangPrimaryRequestVNextLanguageFfNe   AttachToLangPrimaryRequestVNextLanguage = "ff-ne"
	AttachToLangPrimaryRequestVNextLanguageFfNg   AttachToLangPrimaryRequestVNextLanguage = "ff-ng"
	AttachToLangPrimaryRequestVNextLanguageFfSl   AttachToLangPrimaryRequestVNextLanguage = "ff-sl"
	AttachToLangPrimaryRequestVNextLanguageFfSn   AttachToLangPrimaryRequestVNextLanguage = "ff-sn"
	AttachToLangPrimaryRequestVNextLanguageFi     AttachToLangPrimaryRequestVNextLanguage = "fi"
	AttachToLangPrimaryRequestVNextLanguageFiFi   AttachToLangPrimaryRequestVNextLanguage = "fi-fi"
	AttachToLangPrimaryRequestVNextLanguageFil    AttachToLangPrimaryRequestVNextLanguage = "fil"
	AttachToLangPrimaryRequestVNextLanguageFilPh  AttachToLangPrimaryRequestVNextLanguage = "fil-ph"
	AttachToLangPrimaryRequestVNextLanguageFj     AttachToLangPrimaryRequestVNextLanguage = "fj"
	AttachToLangPrimaryRequestVNextLanguageFo     AttachToLangPrimaryRequestVNextLanguage = "fo"
	AttachToLangPrimaryRequestVNextLanguageFoDk   AttachToLangPrimaryRequestVNextLanguage = "fo-dk"
	AttachToLangPrimaryRequestVNextLanguageFoFo   AttachToLangPrimaryRequestVNextLanguage = "fo-fo"
	AttachToLangPrimaryRequestVNextLanguageFr     AttachToLangPrimaryRequestVNextLanguage = "fr"
	AttachToLangPrimaryRequestVNextLanguageFrBe   AttachToLangPrimaryRequestVNextLanguage = "fr-be"
	AttachToLangPrimaryRequestVNextLanguageFrBf   AttachToLangPrimaryRequestVNextLanguage = "fr-bf"
	AttachToLangPrimaryRequestVNextLanguageFrBi   AttachToLangPrimaryRequestVNextLanguage = "fr-bi"
	AttachToLangPrimaryRequestVNextLanguageFrBj   AttachToLangPrimaryRequestVNextLanguage = "fr-bj"
	AttachToLangPrimaryRequestVNextLanguageFrBl   AttachToLangPrimaryRequestVNextLanguage = "fr-bl"
	AttachToLangPrimaryRequestVNextLanguageFrCa   AttachToLangPrimaryRequestVNextLanguage = "fr-ca"
	AttachToLangPrimaryRequestVNextLanguageFrCd   AttachToLangPrimaryRequestVNextLanguage = "fr-cd"
	AttachToLangPrimaryRequestVNextLanguageFrCf   AttachToLangPrimaryRequestVNextLanguage = "fr-cf"
	AttachToLangPrimaryRequestVNextLanguageFrCg   AttachToLangPrimaryRequestVNextLanguage = "fr-cg"
	AttachToLangPrimaryRequestVNextLanguageFrCh   AttachToLangPrimaryRequestVNextLanguage = "fr-ch"
	AttachToLangPrimaryRequestVNextLanguageFrCi   AttachToLangPrimaryRequestVNextLanguage = "fr-ci"
	AttachToLangPrimaryRequestVNextLanguageFrCm   AttachToLangPrimaryRequestVNextLanguage = "fr-cm"
	AttachToLangPrimaryRequestVNextLanguageFrDj   AttachToLangPrimaryRequestVNextLanguage = "fr-dj"
	AttachToLangPrimaryRequestVNextLanguageFrDz   AttachToLangPrimaryRequestVNextLanguage = "fr-dz"
	AttachToLangPrimaryRequestVNextLanguageFrFr   AttachToLangPrimaryRequestVNextLanguage = "fr-fr"
	AttachToLangPrimaryRequestVNextLanguageFrGa   AttachToLangPrimaryRequestVNextLanguage = "fr-ga"
	AttachToLangPrimaryRequestVNextLanguageFrGf   AttachToLangPrimaryRequestVNextLanguage = "fr-gf"
	AttachToLangPrimaryRequestVNextLanguageFrGn   AttachToLangPrimaryRequestVNextLanguage = "fr-gn"
	AttachToLangPrimaryRequestVNextLanguageFrGp   AttachToLangPrimaryRequestVNextLanguage = "fr-gp"
	AttachToLangPrimaryRequestVNextLanguageFrGq   AttachToLangPrimaryRequestVNextLanguage = "fr-gq"
	AttachToLangPrimaryRequestVNextLanguageFrHt   AttachToLangPrimaryRequestVNextLanguage = "fr-ht"
	AttachToLangPrimaryRequestVNextLanguageFrKm   AttachToLangPrimaryRequestVNextLanguage = "fr-km"
	AttachToLangPrimaryRequestVNextLanguageFrLu   AttachToLangPrimaryRequestVNextLanguage = "fr-lu"
	AttachToLangPrimaryRequestVNextLanguageFrMa   AttachToLangPrimaryRequestVNextLanguage = "fr-ma"
	AttachToLangPrimaryRequestVNextLanguageFrMc   AttachToLangPrimaryRequestVNextLanguage = "fr-mc"
	AttachToLangPrimaryRequestVNextLanguageFrMf   AttachToLangPrimaryRequestVNextLanguage = "fr-mf"
	AttachToLangPrimaryRequestVNextLanguageFrMg   AttachToLangPrimaryRequestVNextLanguage = "fr-mg"
	AttachToLangPrimaryRequestVNextLanguageFrMl   AttachToLangPrimaryRequestVNextLanguage = "fr-ml"
	AttachToLangPrimaryRequestVNextLanguageFrMq   AttachToLangPrimaryRequestVNextLanguage = "fr-mq"
	AttachToLangPrimaryRequestVNextLanguageFrMr   AttachToLangPrimaryRequestVNextLanguage = "fr-mr"
	AttachToLangPrimaryRequestVNextLanguageFrMu   AttachToLangPrimaryRequestVNextLanguage = "fr-mu"
	AttachToLangPrimaryRequestVNextLanguageFrNc   AttachToLangPrimaryRequestVNextLanguage = "fr-nc"
	AttachToLangPrimaryRequestVNextLanguageFrNe   AttachToLangPrimaryRequestVNextLanguage = "fr-ne"
	AttachToLangPrimaryRequestVNextLanguageFrPf   AttachToLangPrimaryRequestVNextLanguage = "fr-pf"
	AttachToLangPrimaryRequestVNextLanguageFrPm   AttachToLangPrimaryRequestVNextLanguage = "fr-pm"
	AttachToLangPrimaryRequestVNextLanguageFrRe   AttachToLangPrimaryRequestVNextLanguage = "fr-re"
	AttachToLangPrimaryRequestVNextLanguageFrRw   AttachToLangPrimaryRequestVNextLanguage = "fr-rw"
	AttachToLangPrimaryRequestVNextLanguageFrSc   AttachToLangPrimaryRequestVNextLanguage = "fr-sc"
	AttachToLangPrimaryRequestVNextLanguageFrSn   AttachToLangPrimaryRequestVNextLanguage = "fr-sn"
	AttachToLangPrimaryRequestVNextLanguageFrSy   AttachToLangPrimaryRequestVNextLanguage = "fr-sy"
	AttachToLangPrimaryRequestVNextLanguageFrTd   AttachToLangPrimaryRequestVNextLanguage = "fr-td"
	AttachToLangPrimaryRequestVNextLanguageFrTg   AttachToLangPrimaryRequestVNextLanguage = "fr-tg"
	AttachToLangPrimaryRequestVNextLanguageFrTn   AttachToLangPrimaryRequestVNextLanguage = "fr-tn"
	AttachToLangPrimaryRequestVNextLanguageFrVu   AttachToLangPrimaryRequestVNextLanguage = "fr-vu"
	AttachToLangPrimaryRequestVNextLanguageFrWf   AttachToLangPrimaryRequestVNextLanguage = "fr-wf"
	AttachToLangPrimaryRequestVNextLanguageFrYt   AttachToLangPrimaryRequestVNextLanguage = "fr-yt"
	AttachToLangPrimaryRequestVNextLanguageFrr    AttachToLangPrimaryRequestVNextLanguage = "frr"
	AttachToLangPrimaryRequestVNextLanguageFrrDe  AttachToLangPrimaryRequestVNextLanguage = "frr-de"
	AttachToLangPrimaryRequestVNextLanguageFur    AttachToLangPrimaryRequestVNextLanguage = "fur"
	AttachToLangPrimaryRequestVNextLanguageFurIt  AttachToLangPrimaryRequestVNextLanguage = "fur-it"
	AttachToLangPrimaryRequestVNextLanguageFy     AttachToLangPrimaryRequestVNextLanguage = "fy"
	AttachToLangPrimaryRequestVNextLanguageFyNl   AttachToLangPrimaryRequestVNextLanguage = "fy-nl"
	AttachToLangPrimaryRequestVNextLanguageGa     AttachToLangPrimaryRequestVNextLanguage = "ga"
	AttachToLangPrimaryRequestVNextLanguageGaGB   AttachToLangPrimaryRequestVNextLanguage = "ga-gb"
	AttachToLangPrimaryRequestVNextLanguageGaIe   AttachToLangPrimaryRequestVNextLanguage = "ga-ie"
	AttachToLangPrimaryRequestVNextLanguageGd     AttachToLangPrimaryRequestVNextLanguage = "gd"
	AttachToLangPrimaryRequestVNextLanguageGdGB   AttachToLangPrimaryRequestVNextLanguage = "gd-gb"
	AttachToLangPrimaryRequestVNextLanguageGl     AttachToLangPrimaryRequestVNextLanguage = "gl"
	AttachToLangPrimaryRequestVNextLanguageGlEs   AttachToLangPrimaryRequestVNextLanguage = "gl-es"
	AttachToLangPrimaryRequestVNextLanguageGn     AttachToLangPrimaryRequestVNextLanguage = "gn"
	AttachToLangPrimaryRequestVNextLanguageGsw    AttachToLangPrimaryRequestVNextLanguage = "gsw"
	AttachToLangPrimaryRequestVNextLanguageGswCh  AttachToLangPrimaryRequestVNextLanguage = "gsw-ch"
	AttachToLangPrimaryRequestVNextLanguageGswFr  AttachToLangPrimaryRequestVNextLanguage = "gsw-fr"
	AttachToLangPrimaryRequestVNextLanguageGswLi  AttachToLangPrimaryRequestVNextLanguage = "gsw-li"
	AttachToLangPrimaryRequestVNextLanguageGu     AttachToLangPrimaryRequestVNextLanguage = "gu"
	AttachToLangPrimaryRequestVNextLanguageGuIn   AttachToLangPrimaryRequestVNextLanguage = "gu-in"
	AttachToLangPrimaryRequestVNextLanguageGuz    AttachToLangPrimaryRequestVNextLanguage = "guz"
	AttachToLangPrimaryRequestVNextLanguageGuzKe  AttachToLangPrimaryRequestVNextLanguage = "guz-ke"
	AttachToLangPrimaryRequestVNextLanguageGv     AttachToLangPrimaryRequestVNextLanguage = "gv"
	AttachToLangPrimaryRequestVNextLanguageGvIm   AttachToLangPrimaryRequestVNextLanguage = "gv-im"
	AttachToLangPrimaryRequestVNextLanguageHa     AttachToLangPrimaryRequestVNextLanguage = "ha"
	AttachToLangPrimaryRequestVNextLanguageHaGh   AttachToLangPrimaryRequestVNextLanguage = "ha-gh"
	AttachToLangPrimaryRequestVNextLanguageHaNe   AttachToLangPrimaryRequestVNextLanguage = "ha-ne"
	AttachToLangPrimaryRequestVNextLanguageHaNg   AttachToLangPrimaryRequestVNextLanguage = "ha-ng"
	AttachToLangPrimaryRequestVNextLanguageHaw    AttachToLangPrimaryRequestVNextLanguage = "haw"
	AttachToLangPrimaryRequestVNextLanguageHawUs  AttachToLangPrimaryRequestVNextLanguage = "haw-us"
	AttachToLangPrimaryRequestVNextLanguageHe     AttachToLangPrimaryRequestVNextLanguage = "he"
	AttachToLangPrimaryRequestVNextLanguageHeIl   AttachToLangPrimaryRequestVNextLanguage = "he-il"
	AttachToLangPrimaryRequestVNextLanguageHi     AttachToLangPrimaryRequestVNextLanguage = "hi"
	AttachToLangPrimaryRequestVNextLanguageHiIn   AttachToLangPrimaryRequestVNextLanguage = "hi-in"
	AttachToLangPrimaryRequestVNextLanguageHmn    AttachToLangPrimaryRequestVNextLanguage = "hmn"
	AttachToLangPrimaryRequestVNextLanguageHo     AttachToLangPrimaryRequestVNextLanguage = "ho"
	AttachToLangPrimaryRequestVNextLanguageHr     AttachToLangPrimaryRequestVNextLanguage = "hr"
	AttachToLangPrimaryRequestVNextLanguageHrBa   AttachToLangPrimaryRequestVNextLanguage = "hr-ba"
	AttachToLangPrimaryRequestVNextLanguageHrHr   AttachToLangPrimaryRequestVNextLanguage = "hr-hr"
	AttachToLangPrimaryRequestVNextLanguageHsb    AttachToLangPrimaryRequestVNextLanguage = "hsb"
	AttachToLangPrimaryRequestVNextLanguageHsbDe  AttachToLangPrimaryRequestVNextLanguage = "hsb-de"
	AttachToLangPrimaryRequestVNextLanguageHt     AttachToLangPrimaryRequestVNextLanguage = "ht"
	AttachToLangPrimaryRequestVNextLanguageHu     AttachToLangPrimaryRequestVNextLanguage = "hu"
	AttachToLangPrimaryRequestVNextLanguageHuHu   AttachToLangPrimaryRequestVNextLanguage = "hu-hu"
	AttachToLangPrimaryRequestVNextLanguageHy     AttachToLangPrimaryRequestVNextLanguage = "hy"
	AttachToLangPrimaryRequestVNextLanguageHyAm   AttachToLangPrimaryRequestVNextLanguage = "hy-am"
	AttachToLangPrimaryRequestVNextLanguageHz     AttachToLangPrimaryRequestVNextLanguage = "hz"
	AttachToLangPrimaryRequestVNextLanguageIa     AttachToLangPrimaryRequestVNextLanguage = "ia"
	AttachToLangPrimaryRequestVNextLanguageIa001  AttachToLangPrimaryRequestVNextLanguage = "ia-001"
	AttachToLangPrimaryRequestVNextLanguageID     AttachToLangPrimaryRequestVNextLanguage = "id"
	AttachToLangPrimaryRequestVNextLanguageIDID   AttachToLangPrimaryRequestVNextLanguage = "id-id"
	AttachToLangPrimaryRequestVNextLanguageIe     AttachToLangPrimaryRequestVNextLanguage = "ie"
	AttachToLangPrimaryRequestVNextLanguageIg     AttachToLangPrimaryRequestVNextLanguage = "ig"
	AttachToLangPrimaryRequestVNextLanguageIgNg   AttachToLangPrimaryRequestVNextLanguage = "ig-ng"
	AttachToLangPrimaryRequestVNextLanguageIi     AttachToLangPrimaryRequestVNextLanguage = "ii"
	AttachToLangPrimaryRequestVNextLanguageIiCn   AttachToLangPrimaryRequestVNextLanguage = "ii-cn"
	AttachToLangPrimaryRequestVNextLanguageIk     AttachToLangPrimaryRequestVNextLanguage = "ik"
	AttachToLangPrimaryRequestVNextLanguageIo     AttachToLangPrimaryRequestVNextLanguage = "io"
	AttachToLangPrimaryRequestVNextLanguageIs     AttachToLangPrimaryRequestVNextLanguage = "is"
	AttachToLangPrimaryRequestVNextLanguageIsIs   AttachToLangPrimaryRequestVNextLanguage = "is-is"
	AttachToLangPrimaryRequestVNextLanguageIt     AttachToLangPrimaryRequestVNextLanguage = "it"
	AttachToLangPrimaryRequestVNextLanguageItCh   AttachToLangPrimaryRequestVNextLanguage = "it-ch"
	AttachToLangPrimaryRequestVNextLanguageItIt   AttachToLangPrimaryRequestVNextLanguage = "it-it"
	AttachToLangPrimaryRequestVNextLanguageItSm   AttachToLangPrimaryRequestVNextLanguage = "it-sm"
	AttachToLangPrimaryRequestVNextLanguageItVa   AttachToLangPrimaryRequestVNextLanguage = "it-va"
	AttachToLangPrimaryRequestVNextLanguageIu     AttachToLangPrimaryRequestVNextLanguage = "iu"
	AttachToLangPrimaryRequestVNextLanguageJa     AttachToLangPrimaryRequestVNextLanguage = "ja"
	AttachToLangPrimaryRequestVNextLanguageJaJp   AttachToLangPrimaryRequestVNextLanguage = "ja-jp"
	AttachToLangPrimaryRequestVNextLanguageJgo    AttachToLangPrimaryRequestVNextLanguage = "jgo"
	AttachToLangPrimaryRequestVNextLanguageJgoCm  AttachToLangPrimaryRequestVNextLanguage = "jgo-cm"
	AttachToLangPrimaryRequestVNextLanguageJmc    AttachToLangPrimaryRequestVNextLanguage = "jmc"
	AttachToLangPrimaryRequestVNextLanguageJmcTz  AttachToLangPrimaryRequestVNextLanguage = "jmc-tz"
	AttachToLangPrimaryRequestVNextLanguageJv     AttachToLangPrimaryRequestVNextLanguage = "jv"
	AttachToLangPrimaryRequestVNextLanguageJvID   AttachToLangPrimaryRequestVNextLanguage = "jv-id"
	AttachToLangPrimaryRequestVNextLanguageKa     AttachToLangPrimaryRequestVNextLanguage = "ka"
	AttachToLangPrimaryRequestVNextLanguageKaGe   AttachToLangPrimaryRequestVNextLanguage = "ka-ge"
	AttachToLangPrimaryRequestVNextLanguageKab    AttachToLangPrimaryRequestVNextLanguage = "kab"
	AttachToLangPrimaryRequestVNextLanguageKabDz  AttachToLangPrimaryRequestVNextLanguage = "kab-dz"
	AttachToLangPrimaryRequestVNextLanguageKam    AttachToLangPrimaryRequestVNextLanguage = "kam"
	AttachToLangPrimaryRequestVNextLanguageKamKe  AttachToLangPrimaryRequestVNextLanguage = "kam-ke"
	AttachToLangPrimaryRequestVNextLanguageKar    AttachToLangPrimaryRequestVNextLanguage = "kar"
	AttachToLangPrimaryRequestVNextLanguageKde    AttachToLangPrimaryRequestVNextLanguage = "kde"
	AttachToLangPrimaryRequestVNextLanguageKdeTz  AttachToLangPrimaryRequestVNextLanguage = "kde-tz"
	AttachToLangPrimaryRequestVNextLanguageKea    AttachToLangPrimaryRequestVNextLanguage = "kea"
	AttachToLangPrimaryRequestVNextLanguageKeaCv  AttachToLangPrimaryRequestVNextLanguage = "kea-cv"
	AttachToLangPrimaryRequestVNextLanguageKg     AttachToLangPrimaryRequestVNextLanguage = "kg"
	AttachToLangPrimaryRequestVNextLanguageKgp    AttachToLangPrimaryRequestVNextLanguage = "kgp"
	AttachToLangPrimaryRequestVNextLanguageKgpBr  AttachToLangPrimaryRequestVNextLanguage = "kgp-br"
	AttachToLangPrimaryRequestVNextLanguageKh     AttachToLangPrimaryRequestVNextLanguage = "kh"
	AttachToLangPrimaryRequestVNextLanguageKhq    AttachToLangPrimaryRequestVNextLanguage = "khq"
	AttachToLangPrimaryRequestVNextLanguageKhqMl  AttachToLangPrimaryRequestVNextLanguage = "khq-ml"
	AttachToLangPrimaryRequestVNextLanguageKi     AttachToLangPrimaryRequestVNextLanguage = "ki"
	AttachToLangPrimaryRequestVNextLanguageKiKe   AttachToLangPrimaryRequestVNextLanguage = "ki-ke"
	AttachToLangPrimaryRequestVNextLanguageKj     AttachToLangPrimaryRequestVNextLanguage = "kj"
	AttachToLangPrimaryRequestVNextLanguageKk     AttachToLangPrimaryRequestVNextLanguage = "kk"
	AttachToLangPrimaryRequestVNextLanguageKkKz   AttachToLangPrimaryRequestVNextLanguage = "kk-kz"
	AttachToLangPrimaryRequestVNextLanguageKkj    AttachToLangPrimaryRequestVNextLanguage = "kkj"
	AttachToLangPrimaryRequestVNextLanguageKkjCm  AttachToLangPrimaryRequestVNextLanguage = "kkj-cm"
	AttachToLangPrimaryRequestVNextLanguageKl     AttachToLangPrimaryRequestVNextLanguage = "kl"
	AttachToLangPrimaryRequestVNextLanguageKlGl   AttachToLangPrimaryRequestVNextLanguage = "kl-gl"
	AttachToLangPrimaryRequestVNextLanguageKln    AttachToLangPrimaryRequestVNextLanguage = "kln"
	AttachToLangPrimaryRequestVNextLanguageKlnKe  AttachToLangPrimaryRequestVNextLanguage = "kln-ke"
	AttachToLangPrimaryRequestVNextLanguageKm     AttachToLangPrimaryRequestVNextLanguage = "km"
	AttachToLangPrimaryRequestVNextLanguageKmKh   AttachToLangPrimaryRequestVNextLanguage = "km-kh"
	AttachToLangPrimaryRequestVNextLanguageKn     AttachToLangPrimaryRequestVNextLanguage = "kn"
	AttachToLangPrimaryRequestVNextLanguageKnIn   AttachToLangPrimaryRequestVNextLanguage = "kn-in"
	AttachToLangPrimaryRequestVNextLanguageKo     AttachToLangPrimaryRequestVNextLanguage = "ko"
	AttachToLangPrimaryRequestVNextLanguageKoKp   AttachToLangPrimaryRequestVNextLanguage = "ko-kp"
	AttachToLangPrimaryRequestVNextLanguageKoKr   AttachToLangPrimaryRequestVNextLanguage = "ko-kr"
	AttachToLangPrimaryRequestVNextLanguageKok    AttachToLangPrimaryRequestVNextLanguage = "kok"
	AttachToLangPrimaryRequestVNextLanguageKokIn  AttachToLangPrimaryRequestVNextLanguage = "kok-in"
	AttachToLangPrimaryRequestVNextLanguageKr     AttachToLangPrimaryRequestVNextLanguage = "kr"
	AttachToLangPrimaryRequestVNextLanguageKs     AttachToLangPrimaryRequestVNextLanguage = "ks"
	AttachToLangPrimaryRequestVNextLanguageKsIn   AttachToLangPrimaryRequestVNextLanguage = "ks-in"
	AttachToLangPrimaryRequestVNextLanguageKsb    AttachToLangPrimaryRequestVNextLanguage = "ksb"
	AttachToLangPrimaryRequestVNextLanguageKsbTz  AttachToLangPrimaryRequestVNextLanguage = "ksb-tz"
	AttachToLangPrimaryRequestVNextLanguageKsf    AttachToLangPrimaryRequestVNextLanguage = "ksf"
	AttachToLangPrimaryRequestVNextLanguageKsfCm  AttachToLangPrimaryRequestVNextLanguage = "ksf-cm"
	AttachToLangPrimaryRequestVNextLanguageKsh    AttachToLangPrimaryRequestVNextLanguage = "ksh"
	AttachToLangPrimaryRequestVNextLanguageKshDe  AttachToLangPrimaryRequestVNextLanguage = "ksh-de"
	AttachToLangPrimaryRequestVNextLanguageKu     AttachToLangPrimaryRequestVNextLanguage = "ku"
	AttachToLangPrimaryRequestVNextLanguageKuTr   AttachToLangPrimaryRequestVNextLanguage = "ku-tr"
	AttachToLangPrimaryRequestVNextLanguageKv     AttachToLangPrimaryRequestVNextLanguage = "kv"
	AttachToLangPrimaryRequestVNextLanguageKw     AttachToLangPrimaryRequestVNextLanguage = "kw"
	AttachToLangPrimaryRequestVNextLanguageKwGB   AttachToLangPrimaryRequestVNextLanguage = "kw-gb"
	AttachToLangPrimaryRequestVNextLanguageKy     AttachToLangPrimaryRequestVNextLanguage = "ky"
	AttachToLangPrimaryRequestVNextLanguageKyKg   AttachToLangPrimaryRequestVNextLanguage = "ky-kg"
	AttachToLangPrimaryRequestVNextLanguageLa     AttachToLangPrimaryRequestVNextLanguage = "la"
	AttachToLangPrimaryRequestVNextLanguageLag    AttachToLangPrimaryRequestVNextLanguage = "lag"
	AttachToLangPrimaryRequestVNextLanguageLagTz  AttachToLangPrimaryRequestVNextLanguage = "lag-tz"
	AttachToLangPrimaryRequestVNextLanguageLb     AttachToLangPrimaryRequestVNextLanguage = "lb"
	AttachToLangPrimaryRequestVNextLanguageLbLu   AttachToLangPrimaryRequestVNextLanguage = "lb-lu"
	AttachToLangPrimaryRequestVNextLanguageLg     AttachToLangPrimaryRequestVNextLanguage = "lg"
	AttachToLangPrimaryRequestVNextLanguageLgUg   AttachToLangPrimaryRequestVNextLanguage = "lg-ug"
	AttachToLangPrimaryRequestVNextLanguageLi     AttachToLangPrimaryRequestVNextLanguage = "li"
	AttachToLangPrimaryRequestVNextLanguageLkt    AttachToLangPrimaryRequestVNextLanguage = "lkt"
	AttachToLangPrimaryRequestVNextLanguageLktUs  AttachToLangPrimaryRequestVNextLanguage = "lkt-us"
	AttachToLangPrimaryRequestVNextLanguageLn     AttachToLangPrimaryRequestVNextLanguage = "ln"
	AttachToLangPrimaryRequestVNextLanguageLnAo   AttachToLangPrimaryRequestVNextLanguage = "ln-ao"
	AttachToLangPrimaryRequestVNextLanguageLnCd   AttachToLangPrimaryRequestVNextLanguage = "ln-cd"
	AttachToLangPrimaryRequestVNextLanguageLnCf   AttachToLangPrimaryRequestVNextLanguage = "ln-cf"
	AttachToLangPrimaryRequestVNextLanguageLnCg   AttachToLangPrimaryRequestVNextLanguage = "ln-cg"
	AttachToLangPrimaryRequestVNextLanguageLo     AttachToLangPrimaryRequestVNextLanguage = "lo"
	AttachToLangPrimaryRequestVNextLanguageLoLa   AttachToLangPrimaryRequestVNextLanguage = "lo-la"
	AttachToLangPrimaryRequestVNextLanguageLrc    AttachToLangPrimaryRequestVNextLanguage = "lrc"
	AttachToLangPrimaryRequestVNextLanguageLrcIq  AttachToLangPrimaryRequestVNextLanguage = "lrc-iq"
	AttachToLangPrimaryRequestVNextLanguageLrcIr  AttachToLangPrimaryRequestVNextLanguage = "lrc-ir"
	AttachToLangPrimaryRequestVNextLanguageLt     AttachToLangPrimaryRequestVNextLanguage = "lt"
	AttachToLangPrimaryRequestVNextLanguageLtLt   AttachToLangPrimaryRequestVNextLanguage = "lt-lt"
	AttachToLangPrimaryRequestVNextLanguageLu     AttachToLangPrimaryRequestVNextLanguage = "lu"
	AttachToLangPrimaryRequestVNextLanguageLuCd   AttachToLangPrimaryRequestVNextLanguage = "lu-cd"
	AttachToLangPrimaryRequestVNextLanguageLuo    AttachToLangPrimaryRequestVNextLanguage = "luo"
	AttachToLangPrimaryRequestVNextLanguageLuoKe  AttachToLangPrimaryRequestVNextLanguage = "luo-ke"
	AttachToLangPrimaryRequestVNextLanguageLuy    AttachToLangPrimaryRequestVNextLanguage = "luy"
	AttachToLangPrimaryRequestVNextLanguageLuyKe  AttachToLangPrimaryRequestVNextLanguage = "luy-ke"
	AttachToLangPrimaryRequestVNextLanguageLv     AttachToLangPrimaryRequestVNextLanguage = "lv"
	AttachToLangPrimaryRequestVNextLanguageLvLv   AttachToLangPrimaryRequestVNextLanguage = "lv-lv"
	AttachToLangPrimaryRequestVNextLanguageMai    AttachToLangPrimaryRequestVNextLanguage = "mai"
	AttachToLangPrimaryRequestVNextLanguageMaiIn  AttachToLangPrimaryRequestVNextLanguage = "mai-in"
	AttachToLangPrimaryRequestVNextLanguageMas    AttachToLangPrimaryRequestVNextLanguage = "mas"
	AttachToLangPrimaryRequestVNextLanguageMasKe  AttachToLangPrimaryRequestVNextLanguage = "mas-ke"
	AttachToLangPrimaryRequestVNextLanguageMasTz  AttachToLangPrimaryRequestVNextLanguage = "mas-tz"
	AttachToLangPrimaryRequestVNextLanguageMdf    AttachToLangPrimaryRequestVNextLanguage = "mdf"
	AttachToLangPrimaryRequestVNextLanguageMdfRu  AttachToLangPrimaryRequestVNextLanguage = "mdf-ru"
	AttachToLangPrimaryRequestVNextLanguageMer    AttachToLangPrimaryRequestVNextLanguage = "mer"
	AttachToLangPrimaryRequestVNextLanguageMerKe  AttachToLangPrimaryRequestVNextLanguage = "mer-ke"
	AttachToLangPrimaryRequestVNextLanguageMfe    AttachToLangPrimaryRequestVNextLanguage = "mfe"
	AttachToLangPrimaryRequestVNextLanguageMfeMu  AttachToLangPrimaryRequestVNextLanguage = "mfe-mu"
	AttachToLangPrimaryRequestVNextLanguageMg     AttachToLangPrimaryRequestVNextLanguage = "mg"
	AttachToLangPrimaryRequestVNextLanguageMgMg   AttachToLangPrimaryRequestVNextLanguage = "mg-mg"
	AttachToLangPrimaryRequestVNextLanguageMgh    AttachToLangPrimaryRequestVNextLanguage = "mgh"
	AttachToLangPrimaryRequestVNextLanguageMghMz  AttachToLangPrimaryRequestVNextLanguage = "mgh-mz"
	AttachToLangPrimaryRequestVNextLanguageMgo    AttachToLangPrimaryRequestVNextLanguage = "mgo"
	AttachToLangPrimaryRequestVNextLanguageMgoCm  AttachToLangPrimaryRequestVNextLanguage = "mgo-cm"
	AttachToLangPrimaryRequestVNextLanguageMh     AttachToLangPrimaryRequestVNextLanguage = "mh"
	AttachToLangPrimaryRequestVNextLanguageMi     AttachToLangPrimaryRequestVNextLanguage = "mi"
	AttachToLangPrimaryRequestVNextLanguageMiNz   AttachToLangPrimaryRequestVNextLanguage = "mi-nz"
	AttachToLangPrimaryRequestVNextLanguageMk     AttachToLangPrimaryRequestVNextLanguage = "mk"
	AttachToLangPrimaryRequestVNextLanguageMkMk   AttachToLangPrimaryRequestVNextLanguage = "mk-mk"
	AttachToLangPrimaryRequestVNextLanguageMl     AttachToLangPrimaryRequestVNextLanguage = "ml"
	AttachToLangPrimaryRequestVNextLanguageMlIn   AttachToLangPrimaryRequestVNextLanguage = "ml-in"
	AttachToLangPrimaryRequestVNextLanguageMn     AttachToLangPrimaryRequestVNextLanguage = "mn"
	AttachToLangPrimaryRequestVNextLanguageMnMn   AttachToLangPrimaryRequestVNextLanguage = "mn-mn"
	AttachToLangPrimaryRequestVNextLanguageMni    AttachToLangPrimaryRequestVNextLanguage = "mni"
	AttachToLangPrimaryRequestVNextLanguageMniIn  AttachToLangPrimaryRequestVNextLanguage = "mni-in"
	AttachToLangPrimaryRequestVNextLanguageMr     AttachToLangPrimaryRequestVNextLanguage = "mr"
	AttachToLangPrimaryRequestVNextLanguageMrIn   AttachToLangPrimaryRequestVNextLanguage = "mr-in"
	AttachToLangPrimaryRequestVNextLanguageMs     AttachToLangPrimaryRequestVNextLanguage = "ms"
	AttachToLangPrimaryRequestVNextLanguageMsBn   AttachToLangPrimaryRequestVNextLanguage = "ms-bn"
	AttachToLangPrimaryRequestVNextLanguageMsID   AttachToLangPrimaryRequestVNextLanguage = "ms-id"
	AttachToLangPrimaryRequestVNextLanguageMsMy   AttachToLangPrimaryRequestVNextLanguage = "ms-my"
	AttachToLangPrimaryRequestVNextLanguageMsSg   AttachToLangPrimaryRequestVNextLanguage = "ms-sg"
	AttachToLangPrimaryRequestVNextLanguageMt     AttachToLangPrimaryRequestVNextLanguage = "mt"
	AttachToLangPrimaryRequestVNextLanguageMtMt   AttachToLangPrimaryRequestVNextLanguage = "mt-mt"
	AttachToLangPrimaryRequestVNextLanguageMua    AttachToLangPrimaryRequestVNextLanguage = "mua"
	AttachToLangPrimaryRequestVNextLanguageMuaCm  AttachToLangPrimaryRequestVNextLanguage = "mua-cm"
	AttachToLangPrimaryRequestVNextLanguageMy     AttachToLangPrimaryRequestVNextLanguage = "my"
	AttachToLangPrimaryRequestVNextLanguageMyMm   AttachToLangPrimaryRequestVNextLanguage = "my-mm"
	AttachToLangPrimaryRequestVNextLanguageMzn    AttachToLangPrimaryRequestVNextLanguage = "mzn"
	AttachToLangPrimaryRequestVNextLanguageMznIr  AttachToLangPrimaryRequestVNextLanguage = "mzn-ir"
	AttachToLangPrimaryRequestVNextLanguageNa     AttachToLangPrimaryRequestVNextLanguage = "na"
	AttachToLangPrimaryRequestVNextLanguageNaq    AttachToLangPrimaryRequestVNextLanguage = "naq"
	AttachToLangPrimaryRequestVNextLanguageNaqNa  AttachToLangPrimaryRequestVNextLanguage = "naq-na"
	AttachToLangPrimaryRequestVNextLanguageNb     AttachToLangPrimaryRequestVNextLanguage = "nb"
	AttachToLangPrimaryRequestVNextLanguageNbNo   AttachToLangPrimaryRequestVNextLanguage = "nb-no"
	AttachToLangPrimaryRequestVNextLanguageNbSj   AttachToLangPrimaryRequestVNextLanguage = "nb-sj"
	AttachToLangPrimaryRequestVNextLanguageNd     AttachToLangPrimaryRequestVNextLanguage = "nd"
	AttachToLangPrimaryRequestVNextLanguageNdZw   AttachToLangPrimaryRequestVNextLanguage = "nd-zw"
	AttachToLangPrimaryRequestVNextLanguageNds    AttachToLangPrimaryRequestVNextLanguage = "nds"
	AttachToLangPrimaryRequestVNextLanguageNdsDe  AttachToLangPrimaryRequestVNextLanguage = "nds-de"
	AttachToLangPrimaryRequestVNextLanguageNdsNl  AttachToLangPrimaryRequestVNextLanguage = "nds-nl"
	AttachToLangPrimaryRequestVNextLanguageNe     AttachToLangPrimaryRequestVNextLanguage = "ne"
	AttachToLangPrimaryRequestVNextLanguageNeIn   AttachToLangPrimaryRequestVNextLanguage = "ne-in"
	AttachToLangPrimaryRequestVNextLanguageNeNp   AttachToLangPrimaryRequestVNextLanguage = "ne-np"
	AttachToLangPrimaryRequestVNextLanguageNg     AttachToLangPrimaryRequestVNextLanguage = "ng"
	AttachToLangPrimaryRequestVNextLanguageNl     AttachToLangPrimaryRequestVNextLanguage = "nl"
	AttachToLangPrimaryRequestVNextLanguageNlAw   AttachToLangPrimaryRequestVNextLanguage = "nl-aw"
	AttachToLangPrimaryRequestVNextLanguageNlBe   AttachToLangPrimaryRequestVNextLanguage = "nl-be"
	AttachToLangPrimaryRequestVNextLanguageNlBq   AttachToLangPrimaryRequestVNextLanguage = "nl-bq"
	AttachToLangPrimaryRequestVNextLanguageNlCh   AttachToLangPrimaryRequestVNextLanguage = "nl-ch"
	AttachToLangPrimaryRequestVNextLanguageNlCw   AttachToLangPrimaryRequestVNextLanguage = "nl-cw"
	AttachToLangPrimaryRequestVNextLanguageNlLu   AttachToLangPrimaryRequestVNextLanguage = "nl-lu"
	AttachToLangPrimaryRequestVNextLanguageNlNl   AttachToLangPrimaryRequestVNextLanguage = "nl-nl"
	AttachToLangPrimaryRequestVNextLanguageNlSr   AttachToLangPrimaryRequestVNextLanguage = "nl-sr"
	AttachToLangPrimaryRequestVNextLanguageNlSx   AttachToLangPrimaryRequestVNextLanguage = "nl-sx"
	AttachToLangPrimaryRequestVNextLanguageNmg    AttachToLangPrimaryRequestVNextLanguage = "nmg"
	AttachToLangPrimaryRequestVNextLanguageNmgCm  AttachToLangPrimaryRequestVNextLanguage = "nmg-cm"
	AttachToLangPrimaryRequestVNextLanguageNn     AttachToLangPrimaryRequestVNextLanguage = "nn"
	AttachToLangPrimaryRequestVNextLanguageNnNo   AttachToLangPrimaryRequestVNextLanguage = "nn-no"
	AttachToLangPrimaryRequestVNextLanguageNnh    AttachToLangPrimaryRequestVNextLanguage = "nnh"
	AttachToLangPrimaryRequestVNextLanguageNnhCm  AttachToLangPrimaryRequestVNextLanguage = "nnh-cm"
	AttachToLangPrimaryRequestVNextLanguageNo     AttachToLangPrimaryRequestVNextLanguage = "no"
	AttachToLangPrimaryRequestVNextLanguageNoNo   AttachToLangPrimaryRequestVNextLanguage = "no-no"
	AttachToLangPrimaryRequestVNextLanguageNr     AttachToLangPrimaryRequestVNextLanguage = "nr"
	AttachToLangPrimaryRequestVNextLanguageNus    AttachToLangPrimaryRequestVNextLanguage = "nus"
	AttachToLangPrimaryRequestVNextLanguageNusSS  AttachToLangPrimaryRequestVNextLanguage = "nus-ss"
	AttachToLangPrimaryRequestVNextLanguageNv     AttachToLangPrimaryRequestVNextLanguage = "nv"
	AttachToLangPrimaryRequestVNextLanguageNy     AttachToLangPrimaryRequestVNextLanguage = "ny"
	AttachToLangPrimaryRequestVNextLanguageNyn    AttachToLangPrimaryRequestVNextLanguage = "nyn"
	AttachToLangPrimaryRequestVNextLanguageNynUg  AttachToLangPrimaryRequestVNextLanguage = "nyn-ug"
	AttachToLangPrimaryRequestVNextLanguageOc     AttachToLangPrimaryRequestVNextLanguage = "oc"
	AttachToLangPrimaryRequestVNextLanguageOcEs   AttachToLangPrimaryRequestVNextLanguage = "oc-es"
	AttachToLangPrimaryRequestVNextLanguageOcFr   AttachToLangPrimaryRequestVNextLanguage = "oc-fr"
	AttachToLangPrimaryRequestVNextLanguageOj     AttachToLangPrimaryRequestVNextLanguage = "oj"
	AttachToLangPrimaryRequestVNextLanguageOm     AttachToLangPrimaryRequestVNextLanguage = "om"
	AttachToLangPrimaryRequestVNextLanguageOmEt   AttachToLangPrimaryRequestVNextLanguage = "om-et"
	AttachToLangPrimaryRequestVNextLanguageOmKe   AttachToLangPrimaryRequestVNextLanguage = "om-ke"
	AttachToLangPrimaryRequestVNextLanguageOr     AttachToLangPrimaryRequestVNextLanguage = "or"
	AttachToLangPrimaryRequestVNextLanguageOrIn   AttachToLangPrimaryRequestVNextLanguage = "or-in"
	AttachToLangPrimaryRequestVNextLanguageOs     AttachToLangPrimaryRequestVNextLanguage = "os"
	AttachToLangPrimaryRequestVNextLanguageOsGe   AttachToLangPrimaryRequestVNextLanguage = "os-ge"
	AttachToLangPrimaryRequestVNextLanguageOsRu   AttachToLangPrimaryRequestVNextLanguage = "os-ru"
	AttachToLangPrimaryRequestVNextLanguagePa     AttachToLangPrimaryRequestVNextLanguage = "pa"
	AttachToLangPrimaryRequestVNextLanguagePaIn   AttachToLangPrimaryRequestVNextLanguage = "pa-in"
	AttachToLangPrimaryRequestVNextLanguagePaPk   AttachToLangPrimaryRequestVNextLanguage = "pa-pk"
	AttachToLangPrimaryRequestVNextLanguagePcm    AttachToLangPrimaryRequestVNextLanguage = "pcm"
	AttachToLangPrimaryRequestVNextLanguagePcmNg  AttachToLangPrimaryRequestVNextLanguage = "pcm-ng"
	AttachToLangPrimaryRequestVNextLanguagePi     AttachToLangPrimaryRequestVNextLanguage = "pi"
	AttachToLangPrimaryRequestVNextLanguagePis    AttachToLangPrimaryRequestVNextLanguage = "pis"
	AttachToLangPrimaryRequestVNextLanguagePisSb  AttachToLangPrimaryRequestVNextLanguage = "pis-sb"
	AttachToLangPrimaryRequestVNextLanguagePl     AttachToLangPrimaryRequestVNextLanguage = "pl"
	AttachToLangPrimaryRequestVNextLanguagePlPl   AttachToLangPrimaryRequestVNextLanguage = "pl-pl"
	AttachToLangPrimaryRequestVNextLanguagePrg    AttachToLangPrimaryRequestVNextLanguage = "prg"
	AttachToLangPrimaryRequestVNextLanguagePrg001 AttachToLangPrimaryRequestVNextLanguage = "prg-001"
	AttachToLangPrimaryRequestVNextLanguagePs     AttachToLangPrimaryRequestVNextLanguage = "ps"
	AttachToLangPrimaryRequestVNextLanguagePsAf   AttachToLangPrimaryRequestVNextLanguage = "ps-af"
	AttachToLangPrimaryRequestVNextLanguagePsPk   AttachToLangPrimaryRequestVNextLanguage = "ps-pk"
	AttachToLangPrimaryRequestVNextLanguagePt     AttachToLangPrimaryRequestVNextLanguage = "pt"
	AttachToLangPrimaryRequestVNextLanguagePtAo   AttachToLangPrimaryRequestVNextLanguage = "pt-ao"
	AttachToLangPrimaryRequestVNextLanguagePtBr   AttachToLangPrimaryRequestVNextLanguage = "pt-br"
	AttachToLangPrimaryRequestVNextLanguagePtCh   AttachToLangPrimaryRequestVNextLanguage = "pt-ch"
	AttachToLangPrimaryRequestVNextLanguagePtCv   AttachToLangPrimaryRequestVNextLanguage = "pt-cv"
	AttachToLangPrimaryRequestVNextLanguagePtGq   AttachToLangPrimaryRequestVNextLanguage = "pt-gq"
	AttachToLangPrimaryRequestVNextLanguagePtGw   AttachToLangPrimaryRequestVNextLanguage = "pt-gw"
	AttachToLangPrimaryRequestVNextLanguagePtLu   AttachToLangPrimaryRequestVNextLanguage = "pt-lu"
	AttachToLangPrimaryRequestVNextLanguagePtMo   AttachToLangPrimaryRequestVNextLanguage = "pt-mo"
	AttachToLangPrimaryRequestVNextLanguagePtMz   AttachToLangPrimaryRequestVNextLanguage = "pt-mz"
	AttachToLangPrimaryRequestVNextLanguagePtPt   AttachToLangPrimaryRequestVNextLanguage = "pt-pt"
	AttachToLangPrimaryRequestVNextLanguagePtSt   AttachToLangPrimaryRequestVNextLanguage = "pt-st"
	AttachToLangPrimaryRequestVNextLanguagePtTl   AttachToLangPrimaryRequestVNextLanguage = "pt-tl"
	AttachToLangPrimaryRequestVNextLanguageQu     AttachToLangPrimaryRequestVNextLanguage = "qu"
	AttachToLangPrimaryRequestVNextLanguageQuBo   AttachToLangPrimaryRequestVNextLanguage = "qu-bo"
	AttachToLangPrimaryRequestVNextLanguageQuEc   AttachToLangPrimaryRequestVNextLanguage = "qu-ec"
	AttachToLangPrimaryRequestVNextLanguageQuPe   AttachToLangPrimaryRequestVNextLanguage = "qu-pe"
	AttachToLangPrimaryRequestVNextLanguageRaj    AttachToLangPrimaryRequestVNextLanguage = "raj"
	AttachToLangPrimaryRequestVNextLanguageRajIn  AttachToLangPrimaryRequestVNextLanguage = "raj-in"
	AttachToLangPrimaryRequestVNextLanguageRm     AttachToLangPrimaryRequestVNextLanguage = "rm"
	AttachToLangPrimaryRequestVNextLanguageRmCh   AttachToLangPrimaryRequestVNextLanguage = "rm-ch"
	AttachToLangPrimaryRequestVNextLanguageRn     AttachToLangPrimaryRequestVNextLanguage = "rn"
	AttachToLangPrimaryRequestVNextLanguageRnBi   AttachToLangPrimaryRequestVNextLanguage = "rn-bi"
	AttachToLangPrimaryRequestVNextLanguageRo     AttachToLangPrimaryRequestVNextLanguage = "ro"
	AttachToLangPrimaryRequestVNextLanguageRoMd   AttachToLangPrimaryRequestVNextLanguage = "ro-md"
	AttachToLangPrimaryRequestVNextLanguageRoRo   AttachToLangPrimaryRequestVNextLanguage = "ro-ro"
	AttachToLangPrimaryRequestVNextLanguageRof    AttachToLangPrimaryRequestVNextLanguage = "rof"
	AttachToLangPrimaryRequestVNextLanguageRofTz  AttachToLangPrimaryRequestVNextLanguage = "rof-tz"
	AttachToLangPrimaryRequestVNextLanguageRu     AttachToLangPrimaryRequestVNextLanguage = "ru"
	AttachToLangPrimaryRequestVNextLanguageRuBy   AttachToLangPrimaryRequestVNextLanguage = "ru-by"
	AttachToLangPrimaryRequestVNextLanguageRuKg   AttachToLangPrimaryRequestVNextLanguage = "ru-kg"
	AttachToLangPrimaryRequestVNextLanguageRuKz   AttachToLangPrimaryRequestVNextLanguage = "ru-kz"
	AttachToLangPrimaryRequestVNextLanguageRuMd   AttachToLangPrimaryRequestVNextLanguage = "ru-md"
	AttachToLangPrimaryRequestVNextLanguageRuRu   AttachToLangPrimaryRequestVNextLanguage = "ru-ru"
	AttachToLangPrimaryRequestVNextLanguageRuUa   AttachToLangPrimaryRequestVNextLanguage = "ru-ua"
	AttachToLangPrimaryRequestVNextLanguageRw     AttachToLangPrimaryRequestVNextLanguage = "rw"
	AttachToLangPrimaryRequestVNextLanguageRwRw   AttachToLangPrimaryRequestVNextLanguage = "rw-rw"
	AttachToLangPrimaryRequestVNextLanguageRwk    AttachToLangPrimaryRequestVNextLanguage = "rwk"
	AttachToLangPrimaryRequestVNextLanguageRwkTz  AttachToLangPrimaryRequestVNextLanguage = "rwk-tz"
	AttachToLangPrimaryRequestVNextLanguageSa     AttachToLangPrimaryRequestVNextLanguage = "sa"
	AttachToLangPrimaryRequestVNextLanguageSaIn   AttachToLangPrimaryRequestVNextLanguage = "sa-in"
	AttachToLangPrimaryRequestVNextLanguageSah    AttachToLangPrimaryRequestVNextLanguage = "sah"
	AttachToLangPrimaryRequestVNextLanguageSahRu  AttachToLangPrimaryRequestVNextLanguage = "sah-ru"
	AttachToLangPrimaryRequestVNextLanguageSaq    AttachToLangPrimaryRequestVNextLanguage = "saq"
	AttachToLangPrimaryRequestVNextLanguageSaqKe  AttachToLangPrimaryRequestVNextLanguage = "saq-ke"
	AttachToLangPrimaryRequestVNextLanguageSat    AttachToLangPrimaryRequestVNextLanguage = "sat"
	AttachToLangPrimaryRequestVNextLanguageSatIn  AttachToLangPrimaryRequestVNextLanguage = "sat-in"
	AttachToLangPrimaryRequestVNextLanguageSbp    AttachToLangPrimaryRequestVNextLanguage = "sbp"
	AttachToLangPrimaryRequestVNextLanguageSbpTz  AttachToLangPrimaryRequestVNextLanguage = "sbp-tz"
	AttachToLangPrimaryRequestVNextLanguageSc     AttachToLangPrimaryRequestVNextLanguage = "sc"
	AttachToLangPrimaryRequestVNextLanguageScIt   AttachToLangPrimaryRequestVNextLanguage = "sc-it"
	AttachToLangPrimaryRequestVNextLanguageSd     AttachToLangPrimaryRequestVNextLanguage = "sd"
	AttachToLangPrimaryRequestVNextLanguageSdIn   AttachToLangPrimaryRequestVNextLanguage = "sd-in"
	AttachToLangPrimaryRequestVNextLanguageSdPk   AttachToLangPrimaryRequestVNextLanguage = "sd-pk"
	AttachToLangPrimaryRequestVNextLanguageSe     AttachToLangPrimaryRequestVNextLanguage = "se"
	AttachToLangPrimaryRequestVNextLanguageSeFi   AttachToLangPrimaryRequestVNextLanguage = "se-fi"
	AttachToLangPrimaryRequestVNextLanguageSeNo   AttachToLangPrimaryRequestVNextLanguage = "se-no"
	AttachToLangPrimaryRequestVNextLanguageSeSe   AttachToLangPrimaryRequestVNextLanguage = "se-se"
	AttachToLangPrimaryRequestVNextLanguageSeh    AttachToLangPrimaryRequestVNextLanguage = "seh"
	AttachToLangPrimaryRequestVNextLanguageSehMz  AttachToLangPrimaryRequestVNextLanguage = "seh-mz"
	AttachToLangPrimaryRequestVNextLanguageSes    AttachToLangPrimaryRequestVNextLanguage = "ses"
	AttachToLangPrimaryRequestVNextLanguageSesMl  AttachToLangPrimaryRequestVNextLanguage = "ses-ml"
	AttachToLangPrimaryRequestVNextLanguageSg     AttachToLangPrimaryRequestVNextLanguage = "sg"
	AttachToLangPrimaryRequestVNextLanguageSgCf   AttachToLangPrimaryRequestVNextLanguage = "sg-cf"
	AttachToLangPrimaryRequestVNextLanguageShi    AttachToLangPrimaryRequestVNextLanguage = "shi"
	AttachToLangPrimaryRequestVNextLanguageShiMa  AttachToLangPrimaryRequestVNextLanguage = "shi-ma"
	AttachToLangPrimaryRequestVNextLanguageSi     AttachToLangPrimaryRequestVNextLanguage = "si"
	AttachToLangPrimaryRequestVNextLanguageSiLk   AttachToLangPrimaryRequestVNextLanguage = "si-lk"
	AttachToLangPrimaryRequestVNextLanguageSk     AttachToLangPrimaryRequestVNextLanguage = "sk"
	AttachToLangPrimaryRequestVNextLanguageSkSk   AttachToLangPrimaryRequestVNextLanguage = "sk-sk"
	AttachToLangPrimaryRequestVNextLanguageSl     AttachToLangPrimaryRequestVNextLanguage = "sl"
	AttachToLangPrimaryRequestVNextLanguageSlSi   AttachToLangPrimaryRequestVNextLanguage = "sl-si"
	AttachToLangPrimaryRequestVNextLanguageSm     AttachToLangPrimaryRequestVNextLanguage = "sm"
	AttachToLangPrimaryRequestVNextLanguageSmn    AttachToLangPrimaryRequestVNextLanguage = "smn"
	AttachToLangPrimaryRequestVNextLanguageSmnFi  AttachToLangPrimaryRequestVNextLanguage = "smn-fi"
	AttachToLangPrimaryRequestVNextLanguageSMS    AttachToLangPrimaryRequestVNextLanguage = "sms"
	AttachToLangPrimaryRequestVNextLanguageSMSFi  AttachToLangPrimaryRequestVNextLanguage = "sms-fi"
	AttachToLangPrimaryRequestVNextLanguageSn     AttachToLangPrimaryRequestVNextLanguage = "sn"
	AttachToLangPrimaryRequestVNextLanguageSnZw   AttachToLangPrimaryRequestVNextLanguage = "sn-zw"
	AttachToLangPrimaryRequestVNextLanguageSo     AttachToLangPrimaryRequestVNextLanguage = "so"
	AttachToLangPrimaryRequestVNextLanguageSoDj   AttachToLangPrimaryRequestVNextLanguage = "so-dj"
	AttachToLangPrimaryRequestVNextLanguageSoEt   AttachToLangPrimaryRequestVNextLanguage = "so-et"
	AttachToLangPrimaryRequestVNextLanguageSoKe   AttachToLangPrimaryRequestVNextLanguage = "so-ke"
	AttachToLangPrimaryRequestVNextLanguageSoSo   AttachToLangPrimaryRequestVNextLanguage = "so-so"
	AttachToLangPrimaryRequestVNextLanguageSq     AttachToLangPrimaryRequestVNextLanguage = "sq"
	AttachToLangPrimaryRequestVNextLanguageSqAl   AttachToLangPrimaryRequestVNextLanguage = "sq-al"
	AttachToLangPrimaryRequestVNextLanguageSqMk   AttachToLangPrimaryRequestVNextLanguage = "sq-mk"
	AttachToLangPrimaryRequestVNextLanguageSqXk   AttachToLangPrimaryRequestVNextLanguage = "sq-xk"
	AttachToLangPrimaryRequestVNextLanguageSr     AttachToLangPrimaryRequestVNextLanguage = "sr"
	AttachToLangPrimaryRequestVNextLanguageSrBa   AttachToLangPrimaryRequestVNextLanguage = "sr-ba"
	AttachToLangPrimaryRequestVNextLanguageSrCs   AttachToLangPrimaryRequestVNextLanguage = "sr-cs"
	AttachToLangPrimaryRequestVNextLanguageSrMe   AttachToLangPrimaryRequestVNextLanguage = "sr-me"
	AttachToLangPrimaryRequestVNextLanguageSrRs   AttachToLangPrimaryRequestVNextLanguage = "sr-rs"
	AttachToLangPrimaryRequestVNextLanguageSrXk   AttachToLangPrimaryRequestVNextLanguage = "sr-xk"
	AttachToLangPrimaryRequestVNextLanguageSS     AttachToLangPrimaryRequestVNextLanguage = "ss"
	AttachToLangPrimaryRequestVNextLanguageSt     AttachToLangPrimaryRequestVNextLanguage = "st"
	AttachToLangPrimaryRequestVNextLanguageSu     AttachToLangPrimaryRequestVNextLanguage = "su"
	AttachToLangPrimaryRequestVNextLanguageSuID   AttachToLangPrimaryRequestVNextLanguage = "su-id"
	AttachToLangPrimaryRequestVNextLanguageSv     AttachToLangPrimaryRequestVNextLanguage = "sv"
	AttachToLangPrimaryRequestVNextLanguageSvAx   AttachToLangPrimaryRequestVNextLanguage = "sv-ax"
	AttachToLangPrimaryRequestVNextLanguageSvFi   AttachToLangPrimaryRequestVNextLanguage = "sv-fi"
	AttachToLangPrimaryRequestVNextLanguageSvSe   AttachToLangPrimaryRequestVNextLanguage = "sv-se"
	AttachToLangPrimaryRequestVNextLanguageSw     AttachToLangPrimaryRequestVNextLanguage = "sw"
	AttachToLangPrimaryRequestVNextLanguageSwCd   AttachToLangPrimaryRequestVNextLanguage = "sw-cd"
	AttachToLangPrimaryRequestVNextLanguageSwKe   AttachToLangPrimaryRequestVNextLanguage = "sw-ke"
	AttachToLangPrimaryRequestVNextLanguageSwTz   AttachToLangPrimaryRequestVNextLanguage = "sw-tz"
	AttachToLangPrimaryRequestVNextLanguageSwUg   AttachToLangPrimaryRequestVNextLanguage = "sw-ug"
	AttachToLangPrimaryRequestVNextLanguageSy     AttachToLangPrimaryRequestVNextLanguage = "sy"
	AttachToLangPrimaryRequestVNextLanguageTa     AttachToLangPrimaryRequestVNextLanguage = "ta"
	AttachToLangPrimaryRequestVNextLanguageTaIn   AttachToLangPrimaryRequestVNextLanguage = "ta-in"
	AttachToLangPrimaryRequestVNextLanguageTaLk   AttachToLangPrimaryRequestVNextLanguage = "ta-lk"
	AttachToLangPrimaryRequestVNextLanguageTaMy   AttachToLangPrimaryRequestVNextLanguage = "ta-my"
	AttachToLangPrimaryRequestVNextLanguageTaSg   AttachToLangPrimaryRequestVNextLanguage = "ta-sg"
	AttachToLangPrimaryRequestVNextLanguageTe     AttachToLangPrimaryRequestVNextLanguage = "te"
	AttachToLangPrimaryRequestVNextLanguageTeIn   AttachToLangPrimaryRequestVNextLanguage = "te-in"
	AttachToLangPrimaryRequestVNextLanguageTeo    AttachToLangPrimaryRequestVNextLanguage = "teo"
	AttachToLangPrimaryRequestVNextLanguageTeoKe  AttachToLangPrimaryRequestVNextLanguage = "teo-ke"
	AttachToLangPrimaryRequestVNextLanguageTeoUg  AttachToLangPrimaryRequestVNextLanguage = "teo-ug"
	AttachToLangPrimaryRequestVNextLanguageTg     AttachToLangPrimaryRequestVNextLanguage = "tg"
	AttachToLangPrimaryRequestVNextLanguageTgTj   AttachToLangPrimaryRequestVNextLanguage = "tg-tj"
	AttachToLangPrimaryRequestVNextLanguageTh     AttachToLangPrimaryRequestVNextLanguage = "th"
	AttachToLangPrimaryRequestVNextLanguageThTh   AttachToLangPrimaryRequestVNextLanguage = "th-th"
	AttachToLangPrimaryRequestVNextLanguageTi     AttachToLangPrimaryRequestVNextLanguage = "ti"
	AttachToLangPrimaryRequestVNextLanguageTiEr   AttachToLangPrimaryRequestVNextLanguage = "ti-er"
	AttachToLangPrimaryRequestVNextLanguageTiEt   AttachToLangPrimaryRequestVNextLanguage = "ti-et"
	AttachToLangPrimaryRequestVNextLanguageTk     AttachToLangPrimaryRequestVNextLanguage = "tk"
	AttachToLangPrimaryRequestVNextLanguageTkTm   AttachToLangPrimaryRequestVNextLanguage = "tk-tm"
	AttachToLangPrimaryRequestVNextLanguageTl     AttachToLangPrimaryRequestVNextLanguage = "tl"
	AttachToLangPrimaryRequestVNextLanguageTn     AttachToLangPrimaryRequestVNextLanguage = "tn"
	AttachToLangPrimaryRequestVNextLanguageTo     AttachToLangPrimaryRequestVNextLanguage = "to"
	AttachToLangPrimaryRequestVNextLanguageToTo   AttachToLangPrimaryRequestVNextLanguage = "to-to"
	AttachToLangPrimaryRequestVNextLanguageTok    AttachToLangPrimaryRequestVNextLanguage = "tok"
	AttachToLangPrimaryRequestVNextLanguageTok001 AttachToLangPrimaryRequestVNextLanguage = "tok-001"
	AttachToLangPrimaryRequestVNextLanguageTr     AttachToLangPrimaryRequestVNextLanguage = "tr"
	AttachToLangPrimaryRequestVNextLanguageTrCy   AttachToLangPrimaryRequestVNextLanguage = "tr-cy"
	AttachToLangPrimaryRequestVNextLanguageTrTr   AttachToLangPrimaryRequestVNextLanguage = "tr-tr"
	AttachToLangPrimaryRequestVNextLanguageTs     AttachToLangPrimaryRequestVNextLanguage = "ts"
	AttachToLangPrimaryRequestVNextLanguageTt     AttachToLangPrimaryRequestVNextLanguage = "tt"
	AttachToLangPrimaryRequestVNextLanguageTtRu   AttachToLangPrimaryRequestVNextLanguage = "tt-ru"
	AttachToLangPrimaryRequestVNextLanguageTw     AttachToLangPrimaryRequestVNextLanguage = "tw"
	AttachToLangPrimaryRequestVNextLanguageTwq    AttachToLangPrimaryRequestVNextLanguage = "twq"
	AttachToLangPrimaryRequestVNextLanguageTwqNe  AttachToLangPrimaryRequestVNextLanguage = "twq-ne"
	AttachToLangPrimaryRequestVNextLanguageTy     AttachToLangPrimaryRequestVNextLanguage = "ty"
	AttachToLangPrimaryRequestVNextLanguageTzm    AttachToLangPrimaryRequestVNextLanguage = "tzm"
	AttachToLangPrimaryRequestVNextLanguageTzmMa  AttachToLangPrimaryRequestVNextLanguage = "tzm-ma"
	AttachToLangPrimaryRequestVNextLanguageUg     AttachToLangPrimaryRequestVNextLanguage = "ug"
	AttachToLangPrimaryRequestVNextLanguageUgCn   AttachToLangPrimaryRequestVNextLanguage = "ug-cn"
	AttachToLangPrimaryRequestVNextLanguageUk     AttachToLangPrimaryRequestVNextLanguage = "uk"
	AttachToLangPrimaryRequestVNextLanguageUkUa   AttachToLangPrimaryRequestVNextLanguage = "uk-ua"
	AttachToLangPrimaryRequestVNextLanguageUr     AttachToLangPrimaryRequestVNextLanguage = "ur"
	AttachToLangPrimaryRequestVNextLanguageUrIn   AttachToLangPrimaryRequestVNextLanguage = "ur-in"
	AttachToLangPrimaryRequestVNextLanguageUrPk   AttachToLangPrimaryRequestVNextLanguage = "ur-pk"
	AttachToLangPrimaryRequestVNextLanguageUz     AttachToLangPrimaryRequestVNextLanguage = "uz"
	AttachToLangPrimaryRequestVNextLanguageUzAf   AttachToLangPrimaryRequestVNextLanguage = "uz-af"
	AttachToLangPrimaryRequestVNextLanguageUzUz   AttachToLangPrimaryRequestVNextLanguage = "uz-uz"
	AttachToLangPrimaryRequestVNextLanguageVai    AttachToLangPrimaryRequestVNextLanguage = "vai"
	AttachToLangPrimaryRequestVNextLanguageVaiLr  AttachToLangPrimaryRequestVNextLanguage = "vai-lr"
	AttachToLangPrimaryRequestVNextLanguageVe     AttachToLangPrimaryRequestVNextLanguage = "ve"
	AttachToLangPrimaryRequestVNextLanguageVi     AttachToLangPrimaryRequestVNextLanguage = "vi"
	AttachToLangPrimaryRequestVNextLanguageViVn   AttachToLangPrimaryRequestVNextLanguage = "vi-vn"
	AttachToLangPrimaryRequestVNextLanguageVo     AttachToLangPrimaryRequestVNextLanguage = "vo"
	AttachToLangPrimaryRequestVNextLanguageVo001  AttachToLangPrimaryRequestVNextLanguage = "vo-001"
	AttachToLangPrimaryRequestVNextLanguageVun    AttachToLangPrimaryRequestVNextLanguage = "vun"
	AttachToLangPrimaryRequestVNextLanguageVunTz  AttachToLangPrimaryRequestVNextLanguage = "vun-tz"
	AttachToLangPrimaryRequestVNextLanguageWa     AttachToLangPrimaryRequestVNextLanguage = "wa"
	AttachToLangPrimaryRequestVNextLanguageWae    AttachToLangPrimaryRequestVNextLanguage = "wae"
	AttachToLangPrimaryRequestVNextLanguageWaeCh  AttachToLangPrimaryRequestVNextLanguage = "wae-ch"
	AttachToLangPrimaryRequestVNextLanguageWo     AttachToLangPrimaryRequestVNextLanguage = "wo"
	AttachToLangPrimaryRequestVNextLanguageWoSn   AttachToLangPrimaryRequestVNextLanguage = "wo-sn"
	AttachToLangPrimaryRequestVNextLanguageXh     AttachToLangPrimaryRequestVNextLanguage = "xh"
	AttachToLangPrimaryRequestVNextLanguageXhZa   AttachToLangPrimaryRequestVNextLanguage = "xh-za"
	AttachToLangPrimaryRequestVNextLanguageXog    AttachToLangPrimaryRequestVNextLanguage = "xog"
	AttachToLangPrimaryRequestVNextLanguageXogUg  AttachToLangPrimaryRequestVNextLanguage = "xog-ug"
	AttachToLangPrimaryRequestVNextLanguageYav    AttachToLangPrimaryRequestVNextLanguage = "yav"
	AttachToLangPrimaryRequestVNextLanguageYavCm  AttachToLangPrimaryRequestVNextLanguage = "yav-cm"
	AttachToLangPrimaryRequestVNextLanguageYi     AttachToLangPrimaryRequestVNextLanguage = "yi"
	AttachToLangPrimaryRequestVNextLanguageYi001  AttachToLangPrimaryRequestVNextLanguage = "yi-001"
	AttachToLangPrimaryRequestVNextLanguageYo     AttachToLangPrimaryRequestVNextLanguage = "yo"
	AttachToLangPrimaryRequestVNextLanguageYoBj   AttachToLangPrimaryRequestVNextLanguage = "yo-bj"
	AttachToLangPrimaryRequestVNextLanguageYoNg   AttachToLangPrimaryRequestVNextLanguage = "yo-ng"
	AttachToLangPrimaryRequestVNextLanguageYrl    AttachToLangPrimaryRequestVNextLanguage = "yrl"
	AttachToLangPrimaryRequestVNextLanguageYrlBr  AttachToLangPrimaryRequestVNextLanguage = "yrl-br"
	AttachToLangPrimaryRequestVNextLanguageYrlCo  AttachToLangPrimaryRequestVNextLanguage = "yrl-co"
	AttachToLangPrimaryRequestVNextLanguageYrlVe  AttachToLangPrimaryRequestVNextLanguage = "yrl-ve"
	AttachToLangPrimaryRequestVNextLanguageYue    AttachToLangPrimaryRequestVNextLanguage = "yue"
	AttachToLangPrimaryRequestVNextLanguageYueCn  AttachToLangPrimaryRequestVNextLanguage = "yue-cn"
	AttachToLangPrimaryRequestVNextLanguageYueHk  AttachToLangPrimaryRequestVNextLanguage = "yue-hk"
	AttachToLangPrimaryRequestVNextLanguageZa     AttachToLangPrimaryRequestVNextLanguage = "za"
	AttachToLangPrimaryRequestVNextLanguageZgh    AttachToLangPrimaryRequestVNextLanguage = "zgh"
	AttachToLangPrimaryRequestVNextLanguageZghMa  AttachToLangPrimaryRequestVNextLanguage = "zgh-ma"
	AttachToLangPrimaryRequestVNextLanguageZh     AttachToLangPrimaryRequestVNextLanguage = "zh"
	AttachToLangPrimaryRequestVNextLanguageZhCn   AttachToLangPrimaryRequestVNextLanguage = "zh-cn"
	AttachToLangPrimaryRequestVNextLanguageZhHans AttachToLangPrimaryRequestVNextLanguage = "zh-hans"
	AttachToLangPrimaryRequestVNextLanguageZhHant AttachToLangPrimaryRequestVNextLanguage = "zh-hant"
	AttachToLangPrimaryRequestVNextLanguageZhHk   AttachToLangPrimaryRequestVNextLanguage = "zh-hk"
	AttachToLangPrimaryRequestVNextLanguageZhMo   AttachToLangPrimaryRequestVNextLanguage = "zh-mo"
	AttachToLangPrimaryRequestVNextLanguageZhSg   AttachToLangPrimaryRequestVNextLanguage = "zh-sg"
	AttachToLangPrimaryRequestVNextLanguageZhTw   AttachToLangPrimaryRequestVNextLanguage = "zh-tw"
	AttachToLangPrimaryRequestVNextLanguageZu     AttachToLangPrimaryRequestVNextLanguage = "zu"
	AttachToLangPrimaryRequestVNextLanguageZuZa   AttachToLangPrimaryRequestVNextLanguage = "zu-za"
)

// Primary language of the multi-language group.
type AttachToLangPrimaryRequestVNextPrimaryLanguage string

const (
	AttachToLangPrimaryRequestVNextPrimaryLanguageAa     AttachToLangPrimaryRequestVNextPrimaryLanguage = "aa"
	AttachToLangPrimaryRequestVNextPrimaryLanguageAb     AttachToLangPrimaryRequestVNextPrimaryLanguage = "ab"
	AttachToLangPrimaryRequestVNextPrimaryLanguageAe     AttachToLangPrimaryRequestVNextPrimaryLanguage = "ae"
	AttachToLangPrimaryRequestVNextPrimaryLanguageAf     AttachToLangPrimaryRequestVNextPrimaryLanguage = "af"
	AttachToLangPrimaryRequestVNextPrimaryLanguageAfNa   AttachToLangPrimaryRequestVNextPrimaryLanguage = "af-na"
	AttachToLangPrimaryRequestVNextPrimaryLanguageAfZa   AttachToLangPrimaryRequestVNextPrimaryLanguage = "af-za"
	AttachToLangPrimaryRequestVNextPrimaryLanguageAgq    AttachToLangPrimaryRequestVNextPrimaryLanguage = "agq"
	AttachToLangPrimaryRequestVNextPrimaryLanguageAgqCm  AttachToLangPrimaryRequestVNextPrimaryLanguage = "agq-cm"
	AttachToLangPrimaryRequestVNextPrimaryLanguageAk     AttachToLangPrimaryRequestVNextPrimaryLanguage = "ak"
	AttachToLangPrimaryRequestVNextPrimaryLanguageAkGh   AttachToLangPrimaryRequestVNextPrimaryLanguage = "ak-gh"
	AttachToLangPrimaryRequestVNextPrimaryLanguageAm     AttachToLangPrimaryRequestVNextPrimaryLanguage = "am"
	AttachToLangPrimaryRequestVNextPrimaryLanguageAmEt   AttachToLangPrimaryRequestVNextPrimaryLanguage = "am-et"
	AttachToLangPrimaryRequestVNextPrimaryLanguageAn     AttachToLangPrimaryRequestVNextPrimaryLanguage = "an"
	AttachToLangPrimaryRequestVNextPrimaryLanguageAnn    AttachToLangPrimaryRequestVNextPrimaryLanguage = "ann"
	AttachToLangPrimaryRequestVNextPrimaryLanguageAnnNg  AttachToLangPrimaryRequestVNextPrimaryLanguage = "ann-ng"
	AttachToLangPrimaryRequestVNextPrimaryLanguageAr     AttachToLangPrimaryRequestVNextPrimaryLanguage = "ar"
	AttachToLangPrimaryRequestVNextPrimaryLanguageAr001  AttachToLangPrimaryRequestVNextPrimaryLanguage = "ar-001"
	AttachToLangPrimaryRequestVNextPrimaryLanguageArAe   AttachToLangPrimaryRequestVNextPrimaryLanguage = "ar-ae"
	AttachToLangPrimaryRequestVNextPrimaryLanguageArBh   AttachToLangPrimaryRequestVNextPrimaryLanguage = "ar-bh"
	AttachToLangPrimaryRequestVNextPrimaryLanguageArDj   AttachToLangPrimaryRequestVNextPrimaryLanguage = "ar-dj"
	AttachToLangPrimaryRequestVNextPrimaryLanguageArDz   AttachToLangPrimaryRequestVNextPrimaryLanguage = "ar-dz"
	AttachToLangPrimaryRequestVNextPrimaryLanguageArEg   AttachToLangPrimaryRequestVNextPrimaryLanguage = "ar-eg"
	AttachToLangPrimaryRequestVNextPrimaryLanguageArEh   AttachToLangPrimaryRequestVNextPrimaryLanguage = "ar-eh"
	AttachToLangPrimaryRequestVNextPrimaryLanguageArEr   AttachToLangPrimaryRequestVNextPrimaryLanguage = "ar-er"
	AttachToLangPrimaryRequestVNextPrimaryLanguageArIl   AttachToLangPrimaryRequestVNextPrimaryLanguage = "ar-il"
	AttachToLangPrimaryRequestVNextPrimaryLanguageArIq   AttachToLangPrimaryRequestVNextPrimaryLanguage = "ar-iq"
	AttachToLangPrimaryRequestVNextPrimaryLanguageArJo   AttachToLangPrimaryRequestVNextPrimaryLanguage = "ar-jo"
	AttachToLangPrimaryRequestVNextPrimaryLanguageArKm   AttachToLangPrimaryRequestVNextPrimaryLanguage = "ar-km"
	AttachToLangPrimaryRequestVNextPrimaryLanguageArKw   AttachToLangPrimaryRequestVNextPrimaryLanguage = "ar-kw"
	AttachToLangPrimaryRequestVNextPrimaryLanguageArLb   AttachToLangPrimaryRequestVNextPrimaryLanguage = "ar-lb"
	AttachToLangPrimaryRequestVNextPrimaryLanguageArLy   AttachToLangPrimaryRequestVNextPrimaryLanguage = "ar-ly"
	AttachToLangPrimaryRequestVNextPrimaryLanguageArMa   AttachToLangPrimaryRequestVNextPrimaryLanguage = "ar-ma"
	AttachToLangPrimaryRequestVNextPrimaryLanguageArMr   AttachToLangPrimaryRequestVNextPrimaryLanguage = "ar-mr"
	AttachToLangPrimaryRequestVNextPrimaryLanguageArOm   AttachToLangPrimaryRequestVNextPrimaryLanguage = "ar-om"
	AttachToLangPrimaryRequestVNextPrimaryLanguageArPs   AttachToLangPrimaryRequestVNextPrimaryLanguage = "ar-ps"
	AttachToLangPrimaryRequestVNextPrimaryLanguageArQa   AttachToLangPrimaryRequestVNextPrimaryLanguage = "ar-qa"
	AttachToLangPrimaryRequestVNextPrimaryLanguageArSa   AttachToLangPrimaryRequestVNextPrimaryLanguage = "ar-sa"
	AttachToLangPrimaryRequestVNextPrimaryLanguageArSd   AttachToLangPrimaryRequestVNextPrimaryLanguage = "ar-sd"
	AttachToLangPrimaryRequestVNextPrimaryLanguageArSo   AttachToLangPrimaryRequestVNextPrimaryLanguage = "ar-so"
	AttachToLangPrimaryRequestVNextPrimaryLanguageArSS   AttachToLangPrimaryRequestVNextPrimaryLanguage = "ar-ss"
	AttachToLangPrimaryRequestVNextPrimaryLanguageArSy   AttachToLangPrimaryRequestVNextPrimaryLanguage = "ar-sy"
	AttachToLangPrimaryRequestVNextPrimaryLanguageArTd   AttachToLangPrimaryRequestVNextPrimaryLanguage = "ar-td"
	AttachToLangPrimaryRequestVNextPrimaryLanguageArTn   AttachToLangPrimaryRequestVNextPrimaryLanguage = "ar-tn"
	AttachToLangPrimaryRequestVNextPrimaryLanguageArYe   AttachToLangPrimaryRequestVNextPrimaryLanguage = "ar-ye"
	AttachToLangPrimaryRequestVNextPrimaryLanguageAs     AttachToLangPrimaryRequestVNextPrimaryLanguage = "as"
	AttachToLangPrimaryRequestVNextPrimaryLanguageAsIn   AttachToLangPrimaryRequestVNextPrimaryLanguage = "as-in"
	AttachToLangPrimaryRequestVNextPrimaryLanguageAsa    AttachToLangPrimaryRequestVNextPrimaryLanguage = "asa"
	AttachToLangPrimaryRequestVNextPrimaryLanguageAsaTz  AttachToLangPrimaryRequestVNextPrimaryLanguage = "asa-tz"
	AttachToLangPrimaryRequestVNextPrimaryLanguageAst    AttachToLangPrimaryRequestVNextPrimaryLanguage = "ast"
	AttachToLangPrimaryRequestVNextPrimaryLanguageAstEs  AttachToLangPrimaryRequestVNextPrimaryLanguage = "ast-es"
	AttachToLangPrimaryRequestVNextPrimaryLanguageAv     AttachToLangPrimaryRequestVNextPrimaryLanguage = "av"
	AttachToLangPrimaryRequestVNextPrimaryLanguageAy     AttachToLangPrimaryRequestVNextPrimaryLanguage = "ay"
	AttachToLangPrimaryRequestVNextPrimaryLanguageAz     AttachToLangPrimaryRequestVNextPrimaryLanguage = "az"
	AttachToLangPrimaryRequestVNextPrimaryLanguageAzAz   AttachToLangPrimaryRequestVNextPrimaryLanguage = "az-az"
	AttachToLangPrimaryRequestVNextPrimaryLanguageBa     AttachToLangPrimaryRequestVNextPrimaryLanguage = "ba"
	AttachToLangPrimaryRequestVNextPrimaryLanguageBas    AttachToLangPrimaryRequestVNextPrimaryLanguage = "bas"
	AttachToLangPrimaryRequestVNextPrimaryLanguageBasCm  AttachToLangPrimaryRequestVNextPrimaryLanguage = "bas-cm"
	AttachToLangPrimaryRequestVNextPrimaryLanguageBe     AttachToLangPrimaryRequestVNextPrimaryLanguage = "be"
	AttachToLangPrimaryRequestVNextPrimaryLanguageBeBy   AttachToLangPrimaryRequestVNextPrimaryLanguage = "be-by"
	AttachToLangPrimaryRequestVNextPrimaryLanguageBem    AttachToLangPrimaryRequestVNextPrimaryLanguage = "bem"
	AttachToLangPrimaryRequestVNextPrimaryLanguageBemZm  AttachToLangPrimaryRequestVNextPrimaryLanguage = "bem-zm"
	AttachToLangPrimaryRequestVNextPrimaryLanguageBez    AttachToLangPrimaryRequestVNextPrimaryLanguage = "bez"
	AttachToLangPrimaryRequestVNextPrimaryLanguageBezTz  AttachToLangPrimaryRequestVNextPrimaryLanguage = "bez-tz"
	AttachToLangPrimaryRequestVNextPrimaryLanguageBg     AttachToLangPrimaryRequestVNextPrimaryLanguage = "bg"
	AttachToLangPrimaryRequestVNextPrimaryLanguageBgBg   AttachToLangPrimaryRequestVNextPrimaryLanguage = "bg-bg"
	AttachToLangPrimaryRequestVNextPrimaryLanguageBgc    AttachToLangPrimaryRequestVNextPrimaryLanguage = "bgc"
	AttachToLangPrimaryRequestVNextPrimaryLanguageBgcIn  AttachToLangPrimaryRequestVNextPrimaryLanguage = "bgc-in"
	AttachToLangPrimaryRequestVNextPrimaryLanguageBho    AttachToLangPrimaryRequestVNextPrimaryLanguage = "bho"
	AttachToLangPrimaryRequestVNextPrimaryLanguageBhoIn  AttachToLangPrimaryRequestVNextPrimaryLanguage = "bho-in"
	AttachToLangPrimaryRequestVNextPrimaryLanguageBi     AttachToLangPrimaryRequestVNextPrimaryLanguage = "bi"
	AttachToLangPrimaryRequestVNextPrimaryLanguageBm     AttachToLangPrimaryRequestVNextPrimaryLanguage = "bm"
	AttachToLangPrimaryRequestVNextPrimaryLanguageBmMl   AttachToLangPrimaryRequestVNextPrimaryLanguage = "bm-ml"
	AttachToLangPrimaryRequestVNextPrimaryLanguageBn     AttachToLangPrimaryRequestVNextPrimaryLanguage = "bn"
	AttachToLangPrimaryRequestVNextPrimaryLanguageBnBd   AttachToLangPrimaryRequestVNextPrimaryLanguage = "bn-bd"
	AttachToLangPrimaryRequestVNextPrimaryLanguageBnIn   AttachToLangPrimaryRequestVNextPrimaryLanguage = "bn-in"
	AttachToLangPrimaryRequestVNextPrimaryLanguageBo     AttachToLangPrimaryRequestVNextPrimaryLanguage = "bo"
	AttachToLangPrimaryRequestVNextPrimaryLanguageBoCn   AttachToLangPrimaryRequestVNextPrimaryLanguage = "bo-cn"
	AttachToLangPrimaryRequestVNextPrimaryLanguageBoIn   AttachToLangPrimaryRequestVNextPrimaryLanguage = "bo-in"
	AttachToLangPrimaryRequestVNextPrimaryLanguageBr     AttachToLangPrimaryRequestVNextPrimaryLanguage = "br"
	AttachToLangPrimaryRequestVNextPrimaryLanguageBrFr   AttachToLangPrimaryRequestVNextPrimaryLanguage = "br-fr"
	AttachToLangPrimaryRequestVNextPrimaryLanguageBrx    AttachToLangPrimaryRequestVNextPrimaryLanguage = "brx"
	AttachToLangPrimaryRequestVNextPrimaryLanguageBrxIn  AttachToLangPrimaryRequestVNextPrimaryLanguage = "brx-in"
	AttachToLangPrimaryRequestVNextPrimaryLanguageBs     AttachToLangPrimaryRequestVNextPrimaryLanguage = "bs"
	AttachToLangPrimaryRequestVNextPrimaryLanguageBsBa   AttachToLangPrimaryRequestVNextPrimaryLanguage = "bs-ba"
	AttachToLangPrimaryRequestVNextPrimaryLanguageCa     AttachToLangPrimaryRequestVNextPrimaryLanguage = "ca"
	AttachToLangPrimaryRequestVNextPrimaryLanguageCaAd   AttachToLangPrimaryRequestVNextPrimaryLanguage = "ca-ad"
	AttachToLangPrimaryRequestVNextPrimaryLanguageCaEs   AttachToLangPrimaryRequestVNextPrimaryLanguage = "ca-es"
	AttachToLangPrimaryRequestVNextPrimaryLanguageCaFr   AttachToLangPrimaryRequestVNextPrimaryLanguage = "ca-fr"
	AttachToLangPrimaryRequestVNextPrimaryLanguageCaIt   AttachToLangPrimaryRequestVNextPrimaryLanguage = "ca-it"
	AttachToLangPrimaryRequestVNextPrimaryLanguageCcp    AttachToLangPrimaryRequestVNextPrimaryLanguage = "ccp"
	AttachToLangPrimaryRequestVNextPrimaryLanguageCcpBd  AttachToLangPrimaryRequestVNextPrimaryLanguage = "ccp-bd"
	AttachToLangPrimaryRequestVNextPrimaryLanguageCcpIn  AttachToLangPrimaryRequestVNextPrimaryLanguage = "ccp-in"
	AttachToLangPrimaryRequestVNextPrimaryLanguageCe     AttachToLangPrimaryRequestVNextPrimaryLanguage = "ce"
	AttachToLangPrimaryRequestVNextPrimaryLanguageCeRu   AttachToLangPrimaryRequestVNextPrimaryLanguage = "ce-ru"
	AttachToLangPrimaryRequestVNextPrimaryLanguageCeb    AttachToLangPrimaryRequestVNextPrimaryLanguage = "ceb"
	AttachToLangPrimaryRequestVNextPrimaryLanguageCebPh  AttachToLangPrimaryRequestVNextPrimaryLanguage = "ceb-ph"
	AttachToLangPrimaryRequestVNextPrimaryLanguageCgg    AttachToLangPrimaryRequestVNextPrimaryLanguage = "cgg"
	AttachToLangPrimaryRequestVNextPrimaryLanguageCggUg  AttachToLangPrimaryRequestVNextPrimaryLanguage = "cgg-ug"
	AttachToLangPrimaryRequestVNextPrimaryLanguageCh     AttachToLangPrimaryRequestVNextPrimaryLanguage = "ch"
	AttachToLangPrimaryRequestVNextPrimaryLanguageChr    AttachToLangPrimaryRequestVNextPrimaryLanguage = "chr"
	AttachToLangPrimaryRequestVNextPrimaryLanguageChrUs  AttachToLangPrimaryRequestVNextPrimaryLanguage = "chr-us"
	AttachToLangPrimaryRequestVNextPrimaryLanguageCkb    AttachToLangPrimaryRequestVNextPrimaryLanguage = "ckb"
	AttachToLangPrimaryRequestVNextPrimaryLanguageCkbIq  AttachToLangPrimaryRequestVNextPrimaryLanguage = "ckb-iq"
	AttachToLangPrimaryRequestVNextPrimaryLanguageCkbIr  AttachToLangPrimaryRequestVNextPrimaryLanguage = "ckb-ir"
	AttachToLangPrimaryRequestVNextPrimaryLanguageCo     AttachToLangPrimaryRequestVNextPrimaryLanguage = "co"
	AttachToLangPrimaryRequestVNextPrimaryLanguageCr     AttachToLangPrimaryRequestVNextPrimaryLanguage = "cr"
	AttachToLangPrimaryRequestVNextPrimaryLanguageCs     AttachToLangPrimaryRequestVNextPrimaryLanguage = "cs"
	AttachToLangPrimaryRequestVNextPrimaryLanguageCsCz   AttachToLangPrimaryRequestVNextPrimaryLanguage = "cs-cz"
	AttachToLangPrimaryRequestVNextPrimaryLanguageCu     AttachToLangPrimaryRequestVNextPrimaryLanguage = "cu"
	AttachToLangPrimaryRequestVNextPrimaryLanguageCuRu   AttachToLangPrimaryRequestVNextPrimaryLanguage = "cu-ru"
	AttachToLangPrimaryRequestVNextPrimaryLanguageCv     AttachToLangPrimaryRequestVNextPrimaryLanguage = "cv"
	AttachToLangPrimaryRequestVNextPrimaryLanguageCvRu   AttachToLangPrimaryRequestVNextPrimaryLanguage = "cv-ru"
	AttachToLangPrimaryRequestVNextPrimaryLanguageCy     AttachToLangPrimaryRequestVNextPrimaryLanguage = "cy"
	AttachToLangPrimaryRequestVNextPrimaryLanguageCyGB   AttachToLangPrimaryRequestVNextPrimaryLanguage = "cy-gb"
	AttachToLangPrimaryRequestVNextPrimaryLanguageDa     AttachToLangPrimaryRequestVNextPrimaryLanguage = "da"
	AttachToLangPrimaryRequestVNextPrimaryLanguageDaDk   AttachToLangPrimaryRequestVNextPrimaryLanguage = "da-dk"
	AttachToLangPrimaryRequestVNextPrimaryLanguageDaGl   AttachToLangPrimaryRequestVNextPrimaryLanguage = "da-gl"
	AttachToLangPrimaryRequestVNextPrimaryLanguageDav    AttachToLangPrimaryRequestVNextPrimaryLanguage = "dav"
	AttachToLangPrimaryRequestVNextPrimaryLanguageDavKe  AttachToLangPrimaryRequestVNextPrimaryLanguage = "dav-ke"
	AttachToLangPrimaryRequestVNextPrimaryLanguageDe     AttachToLangPrimaryRequestVNextPrimaryLanguage = "de"
	AttachToLangPrimaryRequestVNextPrimaryLanguageDeAt   AttachToLangPrimaryRequestVNextPrimaryLanguage = "de-at"
	AttachToLangPrimaryRequestVNextPrimaryLanguageDeBe   AttachToLangPrimaryRequestVNextPrimaryLanguage = "de-be"
	AttachToLangPrimaryRequestVNextPrimaryLanguageDeCh   AttachToLangPrimaryRequestVNextPrimaryLanguage = "de-ch"
	AttachToLangPrimaryRequestVNextPrimaryLanguageDeDe   AttachToLangPrimaryRequestVNextPrimaryLanguage = "de-de"
	AttachToLangPrimaryRequestVNextPrimaryLanguageDeGr   AttachToLangPrimaryRequestVNextPrimaryLanguage = "de-gr"
	AttachToLangPrimaryRequestVNextPrimaryLanguageDeIt   AttachToLangPrimaryRequestVNextPrimaryLanguage = "de-it"
	AttachToLangPrimaryRequestVNextPrimaryLanguageDeLi   AttachToLangPrimaryRequestVNextPrimaryLanguage = "de-li"
	AttachToLangPrimaryRequestVNextPrimaryLanguageDeLu   AttachToLangPrimaryRequestVNextPrimaryLanguage = "de-lu"
	AttachToLangPrimaryRequestVNextPrimaryLanguageDje    AttachToLangPrimaryRequestVNextPrimaryLanguage = "dje"
	AttachToLangPrimaryRequestVNextPrimaryLanguageDjeNe  AttachToLangPrimaryRequestVNextPrimaryLanguage = "dje-ne"
	AttachToLangPrimaryRequestVNextPrimaryLanguageDoi    AttachToLangPrimaryRequestVNextPrimaryLanguage = "doi"
	AttachToLangPrimaryRequestVNextPrimaryLanguageDoiIn  AttachToLangPrimaryRequestVNextPrimaryLanguage = "doi-in"
	AttachToLangPrimaryRequestVNextPrimaryLanguageDsb    AttachToLangPrimaryRequestVNextPrimaryLanguage = "dsb"
	AttachToLangPrimaryRequestVNextPrimaryLanguageDsbDe  AttachToLangPrimaryRequestVNextPrimaryLanguage = "dsb-de"
	AttachToLangPrimaryRequestVNextPrimaryLanguageDua    AttachToLangPrimaryRequestVNextPrimaryLanguage = "dua"
	AttachToLangPrimaryRequestVNextPrimaryLanguageDuaCm  AttachToLangPrimaryRequestVNextPrimaryLanguage = "dua-cm"
	AttachToLangPrimaryRequestVNextPrimaryLanguageDv     AttachToLangPrimaryRequestVNextPrimaryLanguage = "dv"
	AttachToLangPrimaryRequestVNextPrimaryLanguageDyo    AttachToLangPrimaryRequestVNextPrimaryLanguage = "dyo"
	AttachToLangPrimaryRequestVNextPrimaryLanguageDyoSn  AttachToLangPrimaryRequestVNextPrimaryLanguage = "dyo-sn"
	AttachToLangPrimaryRequestVNextPrimaryLanguageDz     AttachToLangPrimaryRequestVNextPrimaryLanguage = "dz"
	AttachToLangPrimaryRequestVNextPrimaryLanguageDzBt   AttachToLangPrimaryRequestVNextPrimaryLanguage = "dz-bt"
	AttachToLangPrimaryRequestVNextPrimaryLanguageEbu    AttachToLangPrimaryRequestVNextPrimaryLanguage = "ebu"
	AttachToLangPrimaryRequestVNextPrimaryLanguageEbuKe  AttachToLangPrimaryRequestVNextPrimaryLanguage = "ebu-ke"
	AttachToLangPrimaryRequestVNextPrimaryLanguageEe     AttachToLangPrimaryRequestVNextPrimaryLanguage = "ee"
	AttachToLangPrimaryRequestVNextPrimaryLanguageEeGh   AttachToLangPrimaryRequestVNextPrimaryLanguage = "ee-gh"
	AttachToLangPrimaryRequestVNextPrimaryLanguageEeTg   AttachToLangPrimaryRequestVNextPrimaryLanguage = "ee-tg"
	AttachToLangPrimaryRequestVNextPrimaryLanguageEl     AttachToLangPrimaryRequestVNextPrimaryLanguage = "el"
	AttachToLangPrimaryRequestVNextPrimaryLanguageElCy   AttachToLangPrimaryRequestVNextPrimaryLanguage = "el-cy"
	AttachToLangPrimaryRequestVNextPrimaryLanguageElGr   AttachToLangPrimaryRequestVNextPrimaryLanguage = "el-gr"
	AttachToLangPrimaryRequestVNextPrimaryLanguageEn     AttachToLangPrimaryRequestVNextPrimaryLanguage = "en"
	AttachToLangPrimaryRequestVNextPrimaryLanguageEn001  AttachToLangPrimaryRequestVNextPrimaryLanguage = "en-001"
	AttachToLangPrimaryRequestVNextPrimaryLanguageEn150  AttachToLangPrimaryRequestVNextPrimaryLanguage = "en-150"
	AttachToLangPrimaryRequestVNextPrimaryLanguageEnAe   AttachToLangPrimaryRequestVNextPrimaryLanguage = "en-ae"
	AttachToLangPrimaryRequestVNextPrimaryLanguageEnAg   AttachToLangPrimaryRequestVNextPrimaryLanguage = "en-ag"
	AttachToLangPrimaryRequestVNextPrimaryLanguageEnAI   AttachToLangPrimaryRequestVNextPrimaryLanguage = "en-ai"
	AttachToLangPrimaryRequestVNextPrimaryLanguageEnAs   AttachToLangPrimaryRequestVNextPrimaryLanguage = "en-as"
	AttachToLangPrimaryRequestVNextPrimaryLanguageEnAt   AttachToLangPrimaryRequestVNextPrimaryLanguage = "en-at"
	AttachToLangPrimaryRequestVNextPrimaryLanguageEnAu   AttachToLangPrimaryRequestVNextPrimaryLanguage = "en-au"
	AttachToLangPrimaryRequestVNextPrimaryLanguageEnBb   AttachToLangPrimaryRequestVNextPrimaryLanguage = "en-bb"
	AttachToLangPrimaryRequestVNextPrimaryLanguageEnBe   AttachToLangPrimaryRequestVNextPrimaryLanguage = "en-be"
	AttachToLangPrimaryRequestVNextPrimaryLanguageEnBi   AttachToLangPrimaryRequestVNextPrimaryLanguage = "en-bi"
	AttachToLangPrimaryRequestVNextPrimaryLanguageEnBm   AttachToLangPrimaryRequestVNextPrimaryLanguage = "en-bm"
	AttachToLangPrimaryRequestVNextPrimaryLanguageEnBs   AttachToLangPrimaryRequestVNextPrimaryLanguage = "en-bs"
	AttachToLangPrimaryRequestVNextPrimaryLanguageEnBw   AttachToLangPrimaryRequestVNextPrimaryLanguage = "en-bw"
	AttachToLangPrimaryRequestVNextPrimaryLanguageEnBz   AttachToLangPrimaryRequestVNextPrimaryLanguage = "en-bz"
	AttachToLangPrimaryRequestVNextPrimaryLanguageEnCa   AttachToLangPrimaryRequestVNextPrimaryLanguage = "en-ca"
	AttachToLangPrimaryRequestVNextPrimaryLanguageEnCc   AttachToLangPrimaryRequestVNextPrimaryLanguage = "en-cc"
	AttachToLangPrimaryRequestVNextPrimaryLanguageEnCh   AttachToLangPrimaryRequestVNextPrimaryLanguage = "en-ch"
	AttachToLangPrimaryRequestVNextPrimaryLanguageEnCk   AttachToLangPrimaryRequestVNextPrimaryLanguage = "en-ck"
	AttachToLangPrimaryRequestVNextPrimaryLanguageEnCm   AttachToLangPrimaryRequestVNextPrimaryLanguage = "en-cm"
	AttachToLangPrimaryRequestVNextPrimaryLanguageEnCn   AttachToLangPrimaryRequestVNextPrimaryLanguage = "en-cn"
	AttachToLangPrimaryRequestVNextPrimaryLanguageEnCx   AttachToLangPrimaryRequestVNextPrimaryLanguage = "en-cx"
	AttachToLangPrimaryRequestVNextPrimaryLanguageEnCy   AttachToLangPrimaryRequestVNextPrimaryLanguage = "en-cy"
	AttachToLangPrimaryRequestVNextPrimaryLanguageEnDe   AttachToLangPrimaryRequestVNextPrimaryLanguage = "en-de"
	AttachToLangPrimaryRequestVNextPrimaryLanguageEnDg   AttachToLangPrimaryRequestVNextPrimaryLanguage = "en-dg"
	AttachToLangPrimaryRequestVNextPrimaryLanguageEnDk   AttachToLangPrimaryRequestVNextPrimaryLanguage = "en-dk"
	AttachToLangPrimaryRequestVNextPrimaryLanguageEnDm   AttachToLangPrimaryRequestVNextPrimaryLanguage = "en-dm"
	AttachToLangPrimaryRequestVNextPrimaryLanguageEnEe   AttachToLangPrimaryRequestVNextPrimaryLanguage = "en-ee"
	AttachToLangPrimaryRequestVNextPrimaryLanguageEnEg   AttachToLangPrimaryRequestVNextPrimaryLanguage = "en-eg"
	AttachToLangPrimaryRequestVNextPrimaryLanguageEnEr   AttachToLangPrimaryRequestVNextPrimaryLanguage = "en-er"
	AttachToLangPrimaryRequestVNextPrimaryLanguageEnEs   AttachToLangPrimaryRequestVNextPrimaryLanguage = "en-es"
	AttachToLangPrimaryRequestVNextPrimaryLanguageEnFi   AttachToLangPrimaryRequestVNextPrimaryLanguage = "en-fi"
	AttachToLangPrimaryRequestVNextPrimaryLanguageEnFj   AttachToLangPrimaryRequestVNextPrimaryLanguage = "en-fj"
	AttachToLangPrimaryRequestVNextPrimaryLanguageEnFk   AttachToLangPrimaryRequestVNextPrimaryLanguage = "en-fk"
	AttachToLangPrimaryRequestVNextPrimaryLanguageEnFm   AttachToLangPrimaryRequestVNextPrimaryLanguage = "en-fm"
	AttachToLangPrimaryRequestVNextPrimaryLanguageEnFr   AttachToLangPrimaryRequestVNextPrimaryLanguage = "en-fr"
	AttachToLangPrimaryRequestVNextPrimaryLanguageEnGB   AttachToLangPrimaryRequestVNextPrimaryLanguage = "en-gb"
	AttachToLangPrimaryRequestVNextPrimaryLanguageEnGd   AttachToLangPrimaryRequestVNextPrimaryLanguage = "en-gd"
	AttachToLangPrimaryRequestVNextPrimaryLanguageEnGg   AttachToLangPrimaryRequestVNextPrimaryLanguage = "en-gg"
	AttachToLangPrimaryRequestVNextPrimaryLanguageEnGh   AttachToLangPrimaryRequestVNextPrimaryLanguage = "en-gh"
	AttachToLangPrimaryRequestVNextPrimaryLanguageEnGi   AttachToLangPrimaryRequestVNextPrimaryLanguage = "en-gi"
	AttachToLangPrimaryRequestVNextPrimaryLanguageEnGm   AttachToLangPrimaryRequestVNextPrimaryLanguage = "en-gm"
	AttachToLangPrimaryRequestVNextPrimaryLanguageEnGu   AttachToLangPrimaryRequestVNextPrimaryLanguage = "en-gu"
	AttachToLangPrimaryRequestVNextPrimaryLanguageEnGy   AttachToLangPrimaryRequestVNextPrimaryLanguage = "en-gy"
	AttachToLangPrimaryRequestVNextPrimaryLanguageEnHk   AttachToLangPrimaryRequestVNextPrimaryLanguage = "en-hk"
	AttachToLangPrimaryRequestVNextPrimaryLanguageEnID   AttachToLangPrimaryRequestVNextPrimaryLanguage = "en-id"
	AttachToLangPrimaryRequestVNextPrimaryLanguageEnIe   AttachToLangPrimaryRequestVNextPrimaryLanguage = "en-ie"
	AttachToLangPrimaryRequestVNextPrimaryLanguageEnIl   AttachToLangPrimaryRequestVNextPrimaryLanguage = "en-il"
	AttachToLangPrimaryRequestVNextPrimaryLanguageEnIm   AttachToLangPrimaryRequestVNextPrimaryLanguage = "en-im"
	AttachToLangPrimaryRequestVNextPrimaryLanguageEnIn   AttachToLangPrimaryRequestVNextPrimaryLanguage = "en-in"
	AttachToLangPrimaryRequestVNextPrimaryLanguageEnIo   AttachToLangPrimaryRequestVNextPrimaryLanguage = "en-io"
	AttachToLangPrimaryRequestVNextPrimaryLanguageEnJe   AttachToLangPrimaryRequestVNextPrimaryLanguage = "en-je"
	AttachToLangPrimaryRequestVNextPrimaryLanguageEnJm   AttachToLangPrimaryRequestVNextPrimaryLanguage = "en-jm"
	AttachToLangPrimaryRequestVNextPrimaryLanguageEnKe   AttachToLangPrimaryRequestVNextPrimaryLanguage = "en-ke"
	AttachToLangPrimaryRequestVNextPrimaryLanguageEnKi   AttachToLangPrimaryRequestVNextPrimaryLanguage = "en-ki"
	AttachToLangPrimaryRequestVNextPrimaryLanguageEnKn   AttachToLangPrimaryRequestVNextPrimaryLanguage = "en-kn"
	AttachToLangPrimaryRequestVNextPrimaryLanguageEnKy   AttachToLangPrimaryRequestVNextPrimaryLanguage = "en-ky"
	AttachToLangPrimaryRequestVNextPrimaryLanguageEnLc   AttachToLangPrimaryRequestVNextPrimaryLanguage = "en-lc"
	AttachToLangPrimaryRequestVNextPrimaryLanguageEnLr   AttachToLangPrimaryRequestVNextPrimaryLanguage = "en-lr"
	AttachToLangPrimaryRequestVNextPrimaryLanguageEnLs   AttachToLangPrimaryRequestVNextPrimaryLanguage = "en-ls"
	AttachToLangPrimaryRequestVNextPrimaryLanguageEnLu   AttachToLangPrimaryRequestVNextPrimaryLanguage = "en-lu"
	AttachToLangPrimaryRequestVNextPrimaryLanguageEnMg   AttachToLangPrimaryRequestVNextPrimaryLanguage = "en-mg"
	AttachToLangPrimaryRequestVNextPrimaryLanguageEnMh   AttachToLangPrimaryRequestVNextPrimaryLanguage = "en-mh"
	AttachToLangPrimaryRequestVNextPrimaryLanguageEnMo   AttachToLangPrimaryRequestVNextPrimaryLanguage = "en-mo"
	AttachToLangPrimaryRequestVNextPrimaryLanguageEnMp   AttachToLangPrimaryRequestVNextPrimaryLanguage = "en-mp"
	AttachToLangPrimaryRequestVNextPrimaryLanguageEnMs   AttachToLangPrimaryRequestVNextPrimaryLanguage = "en-ms"
	AttachToLangPrimaryRequestVNextPrimaryLanguageEnMt   AttachToLangPrimaryRequestVNextPrimaryLanguage = "en-mt"
	AttachToLangPrimaryRequestVNextPrimaryLanguageEnMu   AttachToLangPrimaryRequestVNextPrimaryLanguage = "en-mu"
	AttachToLangPrimaryRequestVNextPrimaryLanguageEnMv   AttachToLangPrimaryRequestVNextPrimaryLanguage = "en-mv"
	AttachToLangPrimaryRequestVNextPrimaryLanguageEnMw   AttachToLangPrimaryRequestVNextPrimaryLanguage = "en-mw"
	AttachToLangPrimaryRequestVNextPrimaryLanguageEnMx   AttachToLangPrimaryRequestVNextPrimaryLanguage = "en-mx"
	AttachToLangPrimaryRequestVNextPrimaryLanguageEnMy   AttachToLangPrimaryRequestVNextPrimaryLanguage = "en-my"
	AttachToLangPrimaryRequestVNextPrimaryLanguageEnNa   AttachToLangPrimaryRequestVNextPrimaryLanguage = "en-na"
	AttachToLangPrimaryRequestVNextPrimaryLanguageEnNf   AttachToLangPrimaryRequestVNextPrimaryLanguage = "en-nf"
	AttachToLangPrimaryRequestVNextPrimaryLanguageEnNg   AttachToLangPrimaryRequestVNextPrimaryLanguage = "en-ng"
	AttachToLangPrimaryRequestVNextPrimaryLanguageEnNl   AttachToLangPrimaryRequestVNextPrimaryLanguage = "en-nl"
	AttachToLangPrimaryRequestVNextPrimaryLanguageEnNr   AttachToLangPrimaryRequestVNextPrimaryLanguage = "en-nr"
	AttachToLangPrimaryRequestVNextPrimaryLanguageEnNu   AttachToLangPrimaryRequestVNextPrimaryLanguage = "en-nu"
	AttachToLangPrimaryRequestVNextPrimaryLanguageEnNz   AttachToLangPrimaryRequestVNextPrimaryLanguage = "en-nz"
	AttachToLangPrimaryRequestVNextPrimaryLanguageEnPg   AttachToLangPrimaryRequestVNextPrimaryLanguage = "en-pg"
	AttachToLangPrimaryRequestVNextPrimaryLanguageEnPh   AttachToLangPrimaryRequestVNextPrimaryLanguage = "en-ph"
	AttachToLangPrimaryRequestVNextPrimaryLanguageEnPk   AttachToLangPrimaryRequestVNextPrimaryLanguage = "en-pk"
	AttachToLangPrimaryRequestVNextPrimaryLanguageEnPn   AttachToLangPrimaryRequestVNextPrimaryLanguage = "en-pn"
	AttachToLangPrimaryRequestVNextPrimaryLanguageEnPr   AttachToLangPrimaryRequestVNextPrimaryLanguage = "en-pr"
	AttachToLangPrimaryRequestVNextPrimaryLanguageEnPt   AttachToLangPrimaryRequestVNextPrimaryLanguage = "en-pt"
	AttachToLangPrimaryRequestVNextPrimaryLanguageEnPw   AttachToLangPrimaryRequestVNextPrimaryLanguage = "en-pw"
	AttachToLangPrimaryRequestVNextPrimaryLanguageEnRw   AttachToLangPrimaryRequestVNextPrimaryLanguage = "en-rw"
	AttachToLangPrimaryRequestVNextPrimaryLanguageEnSb   AttachToLangPrimaryRequestVNextPrimaryLanguage = "en-sb"
	AttachToLangPrimaryRequestVNextPrimaryLanguageEnSc   AttachToLangPrimaryRequestVNextPrimaryLanguage = "en-sc"
	AttachToLangPrimaryRequestVNextPrimaryLanguageEnSd   AttachToLangPrimaryRequestVNextPrimaryLanguage = "en-sd"
	AttachToLangPrimaryRequestVNextPrimaryLanguageEnSe   AttachToLangPrimaryRequestVNextPrimaryLanguage = "en-se"
	AttachToLangPrimaryRequestVNextPrimaryLanguageEnSg   AttachToLangPrimaryRequestVNextPrimaryLanguage = "en-sg"
	AttachToLangPrimaryRequestVNextPrimaryLanguageEnSh   AttachToLangPrimaryRequestVNextPrimaryLanguage = "en-sh"
	AttachToLangPrimaryRequestVNextPrimaryLanguageEnSi   AttachToLangPrimaryRequestVNextPrimaryLanguage = "en-si"
	AttachToLangPrimaryRequestVNextPrimaryLanguageEnSl   AttachToLangPrimaryRequestVNextPrimaryLanguage = "en-sl"
	AttachToLangPrimaryRequestVNextPrimaryLanguageEnSS   AttachToLangPrimaryRequestVNextPrimaryLanguage = "en-ss"
	AttachToLangPrimaryRequestVNextPrimaryLanguageEnSx   AttachToLangPrimaryRequestVNextPrimaryLanguage = "en-sx"
	AttachToLangPrimaryRequestVNextPrimaryLanguageEnSz   AttachToLangPrimaryRequestVNextPrimaryLanguage = "en-sz"
	AttachToLangPrimaryRequestVNextPrimaryLanguageEnTc   AttachToLangPrimaryRequestVNextPrimaryLanguage = "en-tc"
	AttachToLangPrimaryRequestVNextPrimaryLanguageEnTh   AttachToLangPrimaryRequestVNextPrimaryLanguage = "en-th"
	AttachToLangPrimaryRequestVNextPrimaryLanguageEnTk   AttachToLangPrimaryRequestVNextPrimaryLanguage = "en-tk"
	AttachToLangPrimaryRequestVNextPrimaryLanguageEnTn   AttachToLangPrimaryRequestVNextPrimaryLanguage = "en-tn"
	AttachToLangPrimaryRequestVNextPrimaryLanguageEnTo   AttachToLangPrimaryRequestVNextPrimaryLanguage = "en-to"
	AttachToLangPrimaryRequestVNextPrimaryLanguageEnTt   AttachToLangPrimaryRequestVNextPrimaryLanguage = "en-tt"
	AttachToLangPrimaryRequestVNextPrimaryLanguageEnTv   AttachToLangPrimaryRequestVNextPrimaryLanguage = "en-tv"
	AttachToLangPrimaryRequestVNextPrimaryLanguageEnTz   AttachToLangPrimaryRequestVNextPrimaryLanguage = "en-tz"
	AttachToLangPrimaryRequestVNextPrimaryLanguageEnUg   AttachToLangPrimaryRequestVNextPrimaryLanguage = "en-ug"
	AttachToLangPrimaryRequestVNextPrimaryLanguageEnUm   AttachToLangPrimaryRequestVNextPrimaryLanguage = "en-um"
	AttachToLangPrimaryRequestVNextPrimaryLanguageEnUs   AttachToLangPrimaryRequestVNextPrimaryLanguage = "en-us"
	AttachToLangPrimaryRequestVNextPrimaryLanguageEnVc   AttachToLangPrimaryRequestVNextPrimaryLanguage = "en-vc"
	AttachToLangPrimaryRequestVNextPrimaryLanguageEnVg   AttachToLangPrimaryRequestVNextPrimaryLanguage = "en-vg"
	AttachToLangPrimaryRequestVNextPrimaryLanguageEnVi   AttachToLangPrimaryRequestVNextPrimaryLanguage = "en-vi"
	AttachToLangPrimaryRequestVNextPrimaryLanguageEnVn   AttachToLangPrimaryRequestVNextPrimaryLanguage = "en-vn"
	AttachToLangPrimaryRequestVNextPrimaryLanguageEnVu   AttachToLangPrimaryRequestVNextPrimaryLanguage = "en-vu"
	AttachToLangPrimaryRequestVNextPrimaryLanguageEnWs   AttachToLangPrimaryRequestVNextPrimaryLanguage = "en-ws"
	AttachToLangPrimaryRequestVNextPrimaryLanguageEnZa   AttachToLangPrimaryRequestVNextPrimaryLanguage = "en-za"
	AttachToLangPrimaryRequestVNextPrimaryLanguageEnZm   AttachToLangPrimaryRequestVNextPrimaryLanguage = "en-zm"
	AttachToLangPrimaryRequestVNextPrimaryLanguageEnZw   AttachToLangPrimaryRequestVNextPrimaryLanguage = "en-zw"
	AttachToLangPrimaryRequestVNextPrimaryLanguageEo     AttachToLangPrimaryRequestVNextPrimaryLanguage = "eo"
	AttachToLangPrimaryRequestVNextPrimaryLanguageEo001  AttachToLangPrimaryRequestVNextPrimaryLanguage = "eo-001"
	AttachToLangPrimaryRequestVNextPrimaryLanguageEs     AttachToLangPrimaryRequestVNextPrimaryLanguage = "es"
	AttachToLangPrimaryRequestVNextPrimaryLanguageEs419  AttachToLangPrimaryRequestVNextPrimaryLanguage = "es-419"
	AttachToLangPrimaryRequestVNextPrimaryLanguageEsAr   AttachToLangPrimaryRequestVNextPrimaryLanguage = "es-ar"
	AttachToLangPrimaryRequestVNextPrimaryLanguageEsBo   AttachToLangPrimaryRequestVNextPrimaryLanguage = "es-bo"
	AttachToLangPrimaryRequestVNextPrimaryLanguageEsBr   AttachToLangPrimaryRequestVNextPrimaryLanguage = "es-br"
	AttachToLangPrimaryRequestVNextPrimaryLanguageEsBz   AttachToLangPrimaryRequestVNextPrimaryLanguage = "es-bz"
	AttachToLangPrimaryRequestVNextPrimaryLanguageEsCl   AttachToLangPrimaryRequestVNextPrimaryLanguage = "es-cl"
	AttachToLangPrimaryRequestVNextPrimaryLanguageEsCo   AttachToLangPrimaryRequestVNextPrimaryLanguage = "es-co"
	AttachToLangPrimaryRequestVNextPrimaryLanguageEsCr   AttachToLangPrimaryRequestVNextPrimaryLanguage = "es-cr"
	AttachToLangPrimaryRequestVNextPrimaryLanguageEsCu   AttachToLangPrimaryRequestVNextPrimaryLanguage = "es-cu"
	AttachToLangPrimaryRequestVNextPrimaryLanguageEsDo   AttachToLangPrimaryRequestVNextPrimaryLanguage = "es-do"
	AttachToLangPrimaryRequestVNextPrimaryLanguageEsEa   AttachToLangPrimaryRequestVNextPrimaryLanguage = "es-ea"
	AttachToLangPrimaryRequestVNextPrimaryLanguageEsEc   AttachToLangPrimaryRequestVNextPrimaryLanguage = "es-ec"
	AttachToLangPrimaryRequestVNextPrimaryLanguageEsEs   AttachToLangPrimaryRequestVNextPrimaryLanguage = "es-es"
	AttachToLangPrimaryRequestVNextPrimaryLanguageEsGq   AttachToLangPrimaryRequestVNextPrimaryLanguage = "es-gq"
	AttachToLangPrimaryRequestVNextPrimaryLanguageEsGt   AttachToLangPrimaryRequestVNextPrimaryLanguage = "es-gt"
	AttachToLangPrimaryRequestVNextPrimaryLanguageEsHn   AttachToLangPrimaryRequestVNextPrimaryLanguage = "es-hn"
	AttachToLangPrimaryRequestVNextPrimaryLanguageEsIc   AttachToLangPrimaryRequestVNextPrimaryLanguage = "es-ic"
	AttachToLangPrimaryRequestVNextPrimaryLanguageEsMx   AttachToLangPrimaryRequestVNextPrimaryLanguage = "es-mx"
	AttachToLangPrimaryRequestVNextPrimaryLanguageEsNi   AttachToLangPrimaryRequestVNextPrimaryLanguage = "es-ni"
	AttachToLangPrimaryRequestVNextPrimaryLanguageEsPa   AttachToLangPrimaryRequestVNextPrimaryLanguage = "es-pa"
	AttachToLangPrimaryRequestVNextPrimaryLanguageEsPe   AttachToLangPrimaryRequestVNextPrimaryLanguage = "es-pe"
	AttachToLangPrimaryRequestVNextPrimaryLanguageEsPh   AttachToLangPrimaryRequestVNextPrimaryLanguage = "es-ph"
	AttachToLangPrimaryRequestVNextPrimaryLanguageEsPr   AttachToLangPrimaryRequestVNextPrimaryLanguage = "es-pr"
	AttachToLangPrimaryRequestVNextPrimaryLanguageEsPy   AttachToLangPrimaryRequestVNextPrimaryLanguage = "es-py"
	AttachToLangPrimaryRequestVNextPrimaryLanguageEsSv   AttachToLangPrimaryRequestVNextPrimaryLanguage = "es-sv"
	AttachToLangPrimaryRequestVNextPrimaryLanguageEsUs   AttachToLangPrimaryRequestVNextPrimaryLanguage = "es-us"
	AttachToLangPrimaryRequestVNextPrimaryLanguageEsUy   AttachToLangPrimaryRequestVNextPrimaryLanguage = "es-uy"
	AttachToLangPrimaryRequestVNextPrimaryLanguageEsVe   AttachToLangPrimaryRequestVNextPrimaryLanguage = "es-ve"
	AttachToLangPrimaryRequestVNextPrimaryLanguageEt     AttachToLangPrimaryRequestVNextPrimaryLanguage = "et"
	AttachToLangPrimaryRequestVNextPrimaryLanguageEtEe   AttachToLangPrimaryRequestVNextPrimaryLanguage = "et-ee"
	AttachToLangPrimaryRequestVNextPrimaryLanguageEu     AttachToLangPrimaryRequestVNextPrimaryLanguage = "eu"
	AttachToLangPrimaryRequestVNextPrimaryLanguageEuEs   AttachToLangPrimaryRequestVNextPrimaryLanguage = "eu-es"
	AttachToLangPrimaryRequestVNextPrimaryLanguageEwo    AttachToLangPrimaryRequestVNextPrimaryLanguage = "ewo"
	AttachToLangPrimaryRequestVNextPrimaryLanguageEwoCm  AttachToLangPrimaryRequestVNextPrimaryLanguage = "ewo-cm"
	AttachToLangPrimaryRequestVNextPrimaryLanguageFa     AttachToLangPrimaryRequestVNextPrimaryLanguage = "fa"
	AttachToLangPrimaryRequestVNextPrimaryLanguageFaAf   AttachToLangPrimaryRequestVNextPrimaryLanguage = "fa-af"
	AttachToLangPrimaryRequestVNextPrimaryLanguageFaIr   AttachToLangPrimaryRequestVNextPrimaryLanguage = "fa-ir"
	AttachToLangPrimaryRequestVNextPrimaryLanguageFf     AttachToLangPrimaryRequestVNextPrimaryLanguage = "ff"
	AttachToLangPrimaryRequestVNextPrimaryLanguageFfBf   AttachToLangPrimaryRequestVNextPrimaryLanguage = "ff-bf"
	AttachToLangPrimaryRequestVNextPrimaryLanguageFfCm   AttachToLangPrimaryRequestVNextPrimaryLanguage = "ff-cm"
	AttachToLangPrimaryRequestVNextPrimaryLanguageFfGh   AttachToLangPrimaryRequestVNextPrimaryLanguage = "ff-gh"
	AttachToLangPrimaryRequestVNextPrimaryLanguageFfGm   AttachToLangPrimaryRequestVNextPrimaryLanguage = "ff-gm"
	AttachToLangPrimaryRequestVNextPrimaryLanguageFfGn   AttachToLangPrimaryRequestVNextPrimaryLanguage = "ff-gn"
	AttachToLangPrimaryRequestVNextPrimaryLanguageFfGw   AttachToLangPrimaryRequestVNextPrimaryLanguage = "ff-gw"
	AttachToLangPrimaryRequestVNextPrimaryLanguageFfLr   AttachToLangPrimaryRequestVNextPrimaryLanguage = "ff-lr"
	AttachToLangPrimaryRequestVNextPrimaryLanguageFfMr   AttachToLangPrimaryRequestVNextPrimaryLanguage = "ff-mr"
	AttachToLangPrimaryRequestVNextPrimaryLanguageFfNe   AttachToLangPrimaryRequestVNextPrimaryLanguage = "ff-ne"
	AttachToLangPrimaryRequestVNextPrimaryLanguageFfNg   AttachToLangPrimaryRequestVNextPrimaryLanguage = "ff-ng"
	AttachToLangPrimaryRequestVNextPrimaryLanguageFfSl   AttachToLangPrimaryRequestVNextPrimaryLanguage = "ff-sl"
	AttachToLangPrimaryRequestVNextPrimaryLanguageFfSn   AttachToLangPrimaryRequestVNextPrimaryLanguage = "ff-sn"
	AttachToLangPrimaryRequestVNextPrimaryLanguageFi     AttachToLangPrimaryRequestVNextPrimaryLanguage = "fi"
	AttachToLangPrimaryRequestVNextPrimaryLanguageFiFi   AttachToLangPrimaryRequestVNextPrimaryLanguage = "fi-fi"
	AttachToLangPrimaryRequestVNextPrimaryLanguageFil    AttachToLangPrimaryRequestVNextPrimaryLanguage = "fil"
	AttachToLangPrimaryRequestVNextPrimaryLanguageFilPh  AttachToLangPrimaryRequestVNextPrimaryLanguage = "fil-ph"
	AttachToLangPrimaryRequestVNextPrimaryLanguageFj     AttachToLangPrimaryRequestVNextPrimaryLanguage = "fj"
	AttachToLangPrimaryRequestVNextPrimaryLanguageFo     AttachToLangPrimaryRequestVNextPrimaryLanguage = "fo"
	AttachToLangPrimaryRequestVNextPrimaryLanguageFoDk   AttachToLangPrimaryRequestVNextPrimaryLanguage = "fo-dk"
	AttachToLangPrimaryRequestVNextPrimaryLanguageFoFo   AttachToLangPrimaryRequestVNextPrimaryLanguage = "fo-fo"
	AttachToLangPrimaryRequestVNextPrimaryLanguageFr     AttachToLangPrimaryRequestVNextPrimaryLanguage = "fr"
	AttachToLangPrimaryRequestVNextPrimaryLanguageFrBe   AttachToLangPrimaryRequestVNextPrimaryLanguage = "fr-be"
	AttachToLangPrimaryRequestVNextPrimaryLanguageFrBf   AttachToLangPrimaryRequestVNextPrimaryLanguage = "fr-bf"
	AttachToLangPrimaryRequestVNextPrimaryLanguageFrBi   AttachToLangPrimaryRequestVNextPrimaryLanguage = "fr-bi"
	AttachToLangPrimaryRequestVNextPrimaryLanguageFrBj   AttachToLangPrimaryRequestVNextPrimaryLanguage = "fr-bj"
	AttachToLangPrimaryRequestVNextPrimaryLanguageFrBl   AttachToLangPrimaryRequestVNextPrimaryLanguage = "fr-bl"
	AttachToLangPrimaryRequestVNextPrimaryLanguageFrCa   AttachToLangPrimaryRequestVNextPrimaryLanguage = "fr-ca"
	AttachToLangPrimaryRequestVNextPrimaryLanguageFrCd   AttachToLangPrimaryRequestVNextPrimaryLanguage = "fr-cd"
	AttachToLangPrimaryRequestVNextPrimaryLanguageFrCf   AttachToLangPrimaryRequestVNextPrimaryLanguage = "fr-cf"
	AttachToLangPrimaryRequestVNextPrimaryLanguageFrCg   AttachToLangPrimaryRequestVNextPrimaryLanguage = "fr-cg"
	AttachToLangPrimaryRequestVNextPrimaryLanguageFrCh   AttachToLangPrimaryRequestVNextPrimaryLanguage = "fr-ch"
	AttachToLangPrimaryRequestVNextPrimaryLanguageFrCi   AttachToLangPrimaryRequestVNextPrimaryLanguage = "fr-ci"
	AttachToLangPrimaryRequestVNextPrimaryLanguageFrCm   AttachToLangPrimaryRequestVNextPrimaryLanguage = "fr-cm"
	AttachToLangPrimaryRequestVNextPrimaryLanguageFrDj   AttachToLangPrimaryRequestVNextPrimaryLanguage = "fr-dj"
	AttachToLangPrimaryRequestVNextPrimaryLanguageFrDz   AttachToLangPrimaryRequestVNextPrimaryLanguage = "fr-dz"
	AttachToLangPrimaryRequestVNextPrimaryLanguageFrFr   AttachToLangPrimaryRequestVNextPrimaryLanguage = "fr-fr"
	AttachToLangPrimaryRequestVNextPrimaryLanguageFrGa   AttachToLangPrimaryRequestVNextPrimaryLanguage = "fr-ga"
	AttachToLangPrimaryRequestVNextPrimaryLanguageFrGf   AttachToLangPrimaryRequestVNextPrimaryLanguage = "fr-gf"
	AttachToLangPrimaryRequestVNextPrimaryLanguageFrGn   AttachToLangPrimaryRequestVNextPrimaryLanguage = "fr-gn"
	AttachToLangPrimaryRequestVNextPrimaryLanguageFrGp   AttachToLangPrimaryRequestVNextPrimaryLanguage = "fr-gp"
	AttachToLangPrimaryRequestVNextPrimaryLanguageFrGq   AttachToLangPrimaryRequestVNextPrimaryLanguage = "fr-gq"
	AttachToLangPrimaryRequestVNextPrimaryLanguageFrHt   AttachToLangPrimaryRequestVNextPrimaryLanguage = "fr-ht"
	AttachToLangPrimaryRequestVNextPrimaryLanguageFrKm   AttachToLangPrimaryRequestVNextPrimaryLanguage = "fr-km"
	AttachToLangPrimaryRequestVNextPrimaryLanguageFrLu   AttachToLangPrimaryRequestVNextPrimaryLanguage = "fr-lu"
	AttachToLangPrimaryRequestVNextPrimaryLanguageFrMa   AttachToLangPrimaryRequestVNextPrimaryLanguage = "fr-ma"
	AttachToLangPrimaryRequestVNextPrimaryLanguageFrMc   AttachToLangPrimaryRequestVNextPrimaryLanguage = "fr-mc"
	AttachToLangPrimaryRequestVNextPrimaryLanguageFrMf   AttachToLangPrimaryRequestVNextPrimaryLanguage = "fr-mf"
	AttachToLangPrimaryRequestVNextPrimaryLanguageFrMg   AttachToLangPrimaryRequestVNextPrimaryLanguage = "fr-mg"
	AttachToLangPrimaryRequestVNextPrimaryLanguageFrMl   AttachToLangPrimaryRequestVNextPrimaryLanguage = "fr-ml"
	AttachToLangPrimaryRequestVNextPrimaryLanguageFrMq   AttachToLangPrimaryRequestVNextPrimaryLanguage = "fr-mq"
	AttachToLangPrimaryRequestVNextPrimaryLanguageFrMr   AttachToLangPrimaryRequestVNextPrimaryLanguage = "fr-mr"
	AttachToLangPrimaryRequestVNextPrimaryLanguageFrMu   AttachToLangPrimaryRequestVNextPrimaryLanguage = "fr-mu"
	AttachToLangPrimaryRequestVNextPrimaryLanguageFrNc   AttachToLangPrimaryRequestVNextPrimaryLanguage = "fr-nc"
	AttachToLangPrimaryRequestVNextPrimaryLanguageFrNe   AttachToLangPrimaryRequestVNextPrimaryLanguage = "fr-ne"
	AttachToLangPrimaryRequestVNextPrimaryLanguageFrPf   AttachToLangPrimaryRequestVNextPrimaryLanguage = "fr-pf"
	AttachToLangPrimaryRequestVNextPrimaryLanguageFrPm   AttachToLangPrimaryRequestVNextPrimaryLanguage = "fr-pm"
	AttachToLangPrimaryRequestVNextPrimaryLanguageFrRe   AttachToLangPrimaryRequestVNextPrimaryLanguage = "fr-re"
	AttachToLangPrimaryRequestVNextPrimaryLanguageFrRw   AttachToLangPrimaryRequestVNextPrimaryLanguage = "fr-rw"
	AttachToLangPrimaryRequestVNextPrimaryLanguageFrSc   AttachToLangPrimaryRequestVNextPrimaryLanguage = "fr-sc"
	AttachToLangPrimaryRequestVNextPrimaryLanguageFrSn   AttachToLangPrimaryRequestVNextPrimaryLanguage = "fr-sn"
	AttachToLangPrimaryRequestVNextPrimaryLanguageFrSy   AttachToLangPrimaryRequestVNextPrimaryLanguage = "fr-sy"
	AttachToLangPrimaryRequestVNextPrimaryLanguageFrTd   AttachToLangPrimaryRequestVNextPrimaryLanguage = "fr-td"
	AttachToLangPrimaryRequestVNextPrimaryLanguageFrTg   AttachToLangPrimaryRequestVNextPrimaryLanguage = "fr-tg"
	AttachToLangPrimaryRequestVNextPrimaryLanguageFrTn   AttachToLangPrimaryRequestVNextPrimaryLanguage = "fr-tn"
	AttachToLangPrimaryRequestVNextPrimaryLanguageFrVu   AttachToLangPrimaryRequestVNextPrimaryLanguage = "fr-vu"
	AttachToLangPrimaryRequestVNextPrimaryLanguageFrWf   AttachToLangPrimaryRequestVNextPrimaryLanguage = "fr-wf"
	AttachToLangPrimaryRequestVNextPrimaryLanguageFrYt   AttachToLangPrimaryRequestVNextPrimaryLanguage = "fr-yt"
	AttachToLangPrimaryRequestVNextPrimaryLanguageFrr    AttachToLangPrimaryRequestVNextPrimaryLanguage = "frr"
	AttachToLangPrimaryRequestVNextPrimaryLanguageFrrDe  AttachToLangPrimaryRequestVNextPrimaryLanguage = "frr-de"
	AttachToLangPrimaryRequestVNextPrimaryLanguageFur    AttachToLangPrimaryRequestVNextPrimaryLanguage = "fur"
	AttachToLangPrimaryRequestVNextPrimaryLanguageFurIt  AttachToLangPrimaryRequestVNextPrimaryLanguage = "fur-it"
	AttachToLangPrimaryRequestVNextPrimaryLanguageFy     AttachToLangPrimaryRequestVNextPrimaryLanguage = "fy"
	AttachToLangPrimaryRequestVNextPrimaryLanguageFyNl   AttachToLangPrimaryRequestVNextPrimaryLanguage = "fy-nl"
	AttachToLangPrimaryRequestVNextPrimaryLanguageGa     AttachToLangPrimaryRequestVNextPrimaryLanguage = "ga"
	AttachToLangPrimaryRequestVNextPrimaryLanguageGaGB   AttachToLangPrimaryRequestVNextPrimaryLanguage = "ga-gb"
	AttachToLangPrimaryRequestVNextPrimaryLanguageGaIe   AttachToLangPrimaryRequestVNextPrimaryLanguage = "ga-ie"
	AttachToLangPrimaryRequestVNextPrimaryLanguageGd     AttachToLangPrimaryRequestVNextPrimaryLanguage = "gd"
	AttachToLangPrimaryRequestVNextPrimaryLanguageGdGB   AttachToLangPrimaryRequestVNextPrimaryLanguage = "gd-gb"
	AttachToLangPrimaryRequestVNextPrimaryLanguageGl     AttachToLangPrimaryRequestVNextPrimaryLanguage = "gl"
	AttachToLangPrimaryRequestVNextPrimaryLanguageGlEs   AttachToLangPrimaryRequestVNextPrimaryLanguage = "gl-es"
	AttachToLangPrimaryRequestVNextPrimaryLanguageGn     AttachToLangPrimaryRequestVNextPrimaryLanguage = "gn"
	AttachToLangPrimaryRequestVNextPrimaryLanguageGsw    AttachToLangPrimaryRequestVNextPrimaryLanguage = "gsw"
	AttachToLangPrimaryRequestVNextPrimaryLanguageGswCh  AttachToLangPrimaryRequestVNextPrimaryLanguage = "gsw-ch"
	AttachToLangPrimaryRequestVNextPrimaryLanguageGswFr  AttachToLangPrimaryRequestVNextPrimaryLanguage = "gsw-fr"
	AttachToLangPrimaryRequestVNextPrimaryLanguageGswLi  AttachToLangPrimaryRequestVNextPrimaryLanguage = "gsw-li"
	AttachToLangPrimaryRequestVNextPrimaryLanguageGu     AttachToLangPrimaryRequestVNextPrimaryLanguage = "gu"
	AttachToLangPrimaryRequestVNextPrimaryLanguageGuIn   AttachToLangPrimaryRequestVNextPrimaryLanguage = "gu-in"
	AttachToLangPrimaryRequestVNextPrimaryLanguageGuz    AttachToLangPrimaryRequestVNextPrimaryLanguage = "guz"
	AttachToLangPrimaryRequestVNextPrimaryLanguageGuzKe  AttachToLangPrimaryRequestVNextPrimaryLanguage = "guz-ke"
	AttachToLangPrimaryRequestVNextPrimaryLanguageGv     AttachToLangPrimaryRequestVNextPrimaryLanguage = "gv"
	AttachToLangPrimaryRequestVNextPrimaryLanguageGvIm   AttachToLangPrimaryRequestVNextPrimaryLanguage = "gv-im"
	AttachToLangPrimaryRequestVNextPrimaryLanguageHa     AttachToLangPrimaryRequestVNextPrimaryLanguage = "ha"
	AttachToLangPrimaryRequestVNextPrimaryLanguageHaGh   AttachToLangPrimaryRequestVNextPrimaryLanguage = "ha-gh"
	AttachToLangPrimaryRequestVNextPrimaryLanguageHaNe   AttachToLangPrimaryRequestVNextPrimaryLanguage = "ha-ne"
	AttachToLangPrimaryRequestVNextPrimaryLanguageHaNg   AttachToLangPrimaryRequestVNextPrimaryLanguage = "ha-ng"
	AttachToLangPrimaryRequestVNextPrimaryLanguageHaw    AttachToLangPrimaryRequestVNextPrimaryLanguage = "haw"
	AttachToLangPrimaryRequestVNextPrimaryLanguageHawUs  AttachToLangPrimaryRequestVNextPrimaryLanguage = "haw-us"
	AttachToLangPrimaryRequestVNextPrimaryLanguageHe     AttachToLangPrimaryRequestVNextPrimaryLanguage = "he"
	AttachToLangPrimaryRequestVNextPrimaryLanguageHeIl   AttachToLangPrimaryRequestVNextPrimaryLanguage = "he-il"
	AttachToLangPrimaryRequestVNextPrimaryLanguageHi     AttachToLangPrimaryRequestVNextPrimaryLanguage = "hi"
	AttachToLangPrimaryRequestVNextPrimaryLanguageHiIn   AttachToLangPrimaryRequestVNextPrimaryLanguage = "hi-in"
	AttachToLangPrimaryRequestVNextPrimaryLanguageHmn    AttachToLangPrimaryRequestVNextPrimaryLanguage = "hmn"
	AttachToLangPrimaryRequestVNextPrimaryLanguageHo     AttachToLangPrimaryRequestVNextPrimaryLanguage = "ho"
	AttachToLangPrimaryRequestVNextPrimaryLanguageHr     AttachToLangPrimaryRequestVNextPrimaryLanguage = "hr"
	AttachToLangPrimaryRequestVNextPrimaryLanguageHrBa   AttachToLangPrimaryRequestVNextPrimaryLanguage = "hr-ba"
	AttachToLangPrimaryRequestVNextPrimaryLanguageHrHr   AttachToLangPrimaryRequestVNextPrimaryLanguage = "hr-hr"
	AttachToLangPrimaryRequestVNextPrimaryLanguageHsb    AttachToLangPrimaryRequestVNextPrimaryLanguage = "hsb"
	AttachToLangPrimaryRequestVNextPrimaryLanguageHsbDe  AttachToLangPrimaryRequestVNextPrimaryLanguage = "hsb-de"
	AttachToLangPrimaryRequestVNextPrimaryLanguageHt     AttachToLangPrimaryRequestVNextPrimaryLanguage = "ht"
	AttachToLangPrimaryRequestVNextPrimaryLanguageHu     AttachToLangPrimaryRequestVNextPrimaryLanguage = "hu"
	AttachToLangPrimaryRequestVNextPrimaryLanguageHuHu   AttachToLangPrimaryRequestVNextPrimaryLanguage = "hu-hu"
	AttachToLangPrimaryRequestVNextPrimaryLanguageHy     AttachToLangPrimaryRequestVNextPrimaryLanguage = "hy"
	AttachToLangPrimaryRequestVNextPrimaryLanguageHyAm   AttachToLangPrimaryRequestVNextPrimaryLanguage = "hy-am"
	AttachToLangPrimaryRequestVNextPrimaryLanguageHz     AttachToLangPrimaryRequestVNextPrimaryLanguage = "hz"
	AttachToLangPrimaryRequestVNextPrimaryLanguageIa     AttachToLangPrimaryRequestVNextPrimaryLanguage = "ia"
	AttachToLangPrimaryRequestVNextPrimaryLanguageIa001  AttachToLangPrimaryRequestVNextPrimaryLanguage = "ia-001"
	AttachToLangPrimaryRequestVNextPrimaryLanguageID     AttachToLangPrimaryRequestVNextPrimaryLanguage = "id"
	AttachToLangPrimaryRequestVNextPrimaryLanguageIDID   AttachToLangPrimaryRequestVNextPrimaryLanguage = "id-id"
	AttachToLangPrimaryRequestVNextPrimaryLanguageIe     AttachToLangPrimaryRequestVNextPrimaryLanguage = "ie"
	AttachToLangPrimaryRequestVNextPrimaryLanguageIg     AttachToLangPrimaryRequestVNextPrimaryLanguage = "ig"
	AttachToLangPrimaryRequestVNextPrimaryLanguageIgNg   AttachToLangPrimaryRequestVNextPrimaryLanguage = "ig-ng"
	AttachToLangPrimaryRequestVNextPrimaryLanguageIi     AttachToLangPrimaryRequestVNextPrimaryLanguage = "ii"
	AttachToLangPrimaryRequestVNextPrimaryLanguageIiCn   AttachToLangPrimaryRequestVNextPrimaryLanguage = "ii-cn"
	AttachToLangPrimaryRequestVNextPrimaryLanguageIk     AttachToLangPrimaryRequestVNextPrimaryLanguage = "ik"
	AttachToLangPrimaryRequestVNextPrimaryLanguageIo     AttachToLangPrimaryRequestVNextPrimaryLanguage = "io"
	AttachToLangPrimaryRequestVNextPrimaryLanguageIs     AttachToLangPrimaryRequestVNextPrimaryLanguage = "is"
	AttachToLangPrimaryRequestVNextPrimaryLanguageIsIs   AttachToLangPrimaryRequestVNextPrimaryLanguage = "is-is"
	AttachToLangPrimaryRequestVNextPrimaryLanguageIt     AttachToLangPrimaryRequestVNextPrimaryLanguage = "it"
	AttachToLangPrimaryRequestVNextPrimaryLanguageItCh   AttachToLangPrimaryRequestVNextPrimaryLanguage = "it-ch"
	AttachToLangPrimaryRequestVNextPrimaryLanguageItIt   AttachToLangPrimaryRequestVNextPrimaryLanguage = "it-it"
	AttachToLangPrimaryRequestVNextPrimaryLanguageItSm   AttachToLangPrimaryRequestVNextPrimaryLanguage = "it-sm"
	AttachToLangPrimaryRequestVNextPrimaryLanguageItVa   AttachToLangPrimaryRequestVNextPrimaryLanguage = "it-va"
	AttachToLangPrimaryRequestVNextPrimaryLanguageIu     AttachToLangPrimaryRequestVNextPrimaryLanguage = "iu"
	AttachToLangPrimaryRequestVNextPrimaryLanguageJa     AttachToLangPrimaryRequestVNextPrimaryLanguage = "ja"
	AttachToLangPrimaryRequestVNextPrimaryLanguageJaJp   AttachToLangPrimaryRequestVNextPrimaryLanguage = "ja-jp"
	AttachToLangPrimaryRequestVNextPrimaryLanguageJgo    AttachToLangPrimaryRequestVNextPrimaryLanguage = "jgo"
	AttachToLangPrimaryRequestVNextPrimaryLanguageJgoCm  AttachToLangPrimaryRequestVNextPrimaryLanguage = "jgo-cm"
	AttachToLangPrimaryRequestVNextPrimaryLanguageJmc    AttachToLangPrimaryRequestVNextPrimaryLanguage = "jmc"
	AttachToLangPrimaryRequestVNextPrimaryLanguageJmcTz  AttachToLangPrimaryRequestVNextPrimaryLanguage = "jmc-tz"
	AttachToLangPrimaryRequestVNextPrimaryLanguageJv     AttachToLangPrimaryRequestVNextPrimaryLanguage = "jv"
	AttachToLangPrimaryRequestVNextPrimaryLanguageJvID   AttachToLangPrimaryRequestVNextPrimaryLanguage = "jv-id"
	AttachToLangPrimaryRequestVNextPrimaryLanguageKa     AttachToLangPrimaryRequestVNextPrimaryLanguage = "ka"
	AttachToLangPrimaryRequestVNextPrimaryLanguageKaGe   AttachToLangPrimaryRequestVNextPrimaryLanguage = "ka-ge"
	AttachToLangPrimaryRequestVNextPrimaryLanguageKab    AttachToLangPrimaryRequestVNextPrimaryLanguage = "kab"
	AttachToLangPrimaryRequestVNextPrimaryLanguageKabDz  AttachToLangPrimaryRequestVNextPrimaryLanguage = "kab-dz"
	AttachToLangPrimaryRequestVNextPrimaryLanguageKam    AttachToLangPrimaryRequestVNextPrimaryLanguage = "kam"
	AttachToLangPrimaryRequestVNextPrimaryLanguageKamKe  AttachToLangPrimaryRequestVNextPrimaryLanguage = "kam-ke"
	AttachToLangPrimaryRequestVNextPrimaryLanguageKar    AttachToLangPrimaryRequestVNextPrimaryLanguage = "kar"
	AttachToLangPrimaryRequestVNextPrimaryLanguageKde    AttachToLangPrimaryRequestVNextPrimaryLanguage = "kde"
	AttachToLangPrimaryRequestVNextPrimaryLanguageKdeTz  AttachToLangPrimaryRequestVNextPrimaryLanguage = "kde-tz"
	AttachToLangPrimaryRequestVNextPrimaryLanguageKea    AttachToLangPrimaryRequestVNextPrimaryLanguage = "kea"
	AttachToLangPrimaryRequestVNextPrimaryLanguageKeaCv  AttachToLangPrimaryRequestVNextPrimaryLanguage = "kea-cv"
	AttachToLangPrimaryRequestVNextPrimaryLanguageKg     AttachToLangPrimaryRequestVNextPrimaryLanguage = "kg"
	AttachToLangPrimaryRequestVNextPrimaryLanguageKgp    AttachToLangPrimaryRequestVNextPrimaryLanguage = "kgp"
	AttachToLangPrimaryRequestVNextPrimaryLanguageKgpBr  AttachToLangPrimaryRequestVNextPrimaryLanguage = "kgp-br"
	AttachToLangPrimaryRequestVNextPrimaryLanguageKh     AttachToLangPrimaryRequestVNextPrimaryLanguage = "kh"
	AttachToLangPrimaryRequestVNextPrimaryLanguageKhq    AttachToLangPrimaryRequestVNextPrimaryLanguage = "khq"
	AttachToLangPrimaryRequestVNextPrimaryLanguageKhqMl  AttachToLangPrimaryRequestVNextPrimaryLanguage = "khq-ml"
	AttachToLangPrimaryRequestVNextPrimaryLanguageKi     AttachToLangPrimaryRequestVNextPrimaryLanguage = "ki"
	AttachToLangPrimaryRequestVNextPrimaryLanguageKiKe   AttachToLangPrimaryRequestVNextPrimaryLanguage = "ki-ke"
	AttachToLangPrimaryRequestVNextPrimaryLanguageKj     AttachToLangPrimaryRequestVNextPrimaryLanguage = "kj"
	AttachToLangPrimaryRequestVNextPrimaryLanguageKk     AttachToLangPrimaryRequestVNextPrimaryLanguage = "kk"
	AttachToLangPrimaryRequestVNextPrimaryLanguageKkKz   AttachToLangPrimaryRequestVNextPrimaryLanguage = "kk-kz"
	AttachToLangPrimaryRequestVNextPrimaryLanguageKkj    AttachToLangPrimaryRequestVNextPrimaryLanguage = "kkj"
	AttachToLangPrimaryRequestVNextPrimaryLanguageKkjCm  AttachToLangPrimaryRequestVNextPrimaryLanguage = "kkj-cm"
	AttachToLangPrimaryRequestVNextPrimaryLanguageKl     AttachToLangPrimaryRequestVNextPrimaryLanguage = "kl"
	AttachToLangPrimaryRequestVNextPrimaryLanguageKlGl   AttachToLangPrimaryRequestVNextPrimaryLanguage = "kl-gl"
	AttachToLangPrimaryRequestVNextPrimaryLanguageKln    AttachToLangPrimaryRequestVNextPrimaryLanguage = "kln"
	AttachToLangPrimaryRequestVNextPrimaryLanguageKlnKe  AttachToLangPrimaryRequestVNextPrimaryLanguage = "kln-ke"
	AttachToLangPrimaryRequestVNextPrimaryLanguageKm     AttachToLangPrimaryRequestVNextPrimaryLanguage = "km"
	AttachToLangPrimaryRequestVNextPrimaryLanguageKmKh   AttachToLangPrimaryRequestVNextPrimaryLanguage = "km-kh"
	AttachToLangPrimaryRequestVNextPrimaryLanguageKn     AttachToLangPrimaryRequestVNextPrimaryLanguage = "kn"
	AttachToLangPrimaryRequestVNextPrimaryLanguageKnIn   AttachToLangPrimaryRequestVNextPrimaryLanguage = "kn-in"
	AttachToLangPrimaryRequestVNextPrimaryLanguageKo     AttachToLangPrimaryRequestVNextPrimaryLanguage = "ko"
	AttachToLangPrimaryRequestVNextPrimaryLanguageKoKp   AttachToLangPrimaryRequestVNextPrimaryLanguage = "ko-kp"
	AttachToLangPrimaryRequestVNextPrimaryLanguageKoKr   AttachToLangPrimaryRequestVNextPrimaryLanguage = "ko-kr"
	AttachToLangPrimaryRequestVNextPrimaryLanguageKok    AttachToLangPrimaryRequestVNextPrimaryLanguage = "kok"
	AttachToLangPrimaryRequestVNextPrimaryLanguageKokIn  AttachToLangPrimaryRequestVNextPrimaryLanguage = "kok-in"
	AttachToLangPrimaryRequestVNextPrimaryLanguageKr     AttachToLangPrimaryRequestVNextPrimaryLanguage = "kr"
	AttachToLangPrimaryRequestVNextPrimaryLanguageKs     AttachToLangPrimaryRequestVNextPrimaryLanguage = "ks"
	AttachToLangPrimaryRequestVNextPrimaryLanguageKsIn   AttachToLangPrimaryRequestVNextPrimaryLanguage = "ks-in"
	AttachToLangPrimaryRequestVNextPrimaryLanguageKsb    AttachToLangPrimaryRequestVNextPrimaryLanguage = "ksb"
	AttachToLangPrimaryRequestVNextPrimaryLanguageKsbTz  AttachToLangPrimaryRequestVNextPrimaryLanguage = "ksb-tz"
	AttachToLangPrimaryRequestVNextPrimaryLanguageKsf    AttachToLangPrimaryRequestVNextPrimaryLanguage = "ksf"
	AttachToLangPrimaryRequestVNextPrimaryLanguageKsfCm  AttachToLangPrimaryRequestVNextPrimaryLanguage = "ksf-cm"
	AttachToLangPrimaryRequestVNextPrimaryLanguageKsh    AttachToLangPrimaryRequestVNextPrimaryLanguage = "ksh"
	AttachToLangPrimaryRequestVNextPrimaryLanguageKshDe  AttachToLangPrimaryRequestVNextPrimaryLanguage = "ksh-de"
	AttachToLangPrimaryRequestVNextPrimaryLanguageKu     AttachToLangPrimaryRequestVNextPrimaryLanguage = "ku"
	AttachToLangPrimaryRequestVNextPrimaryLanguageKuTr   AttachToLangPrimaryRequestVNextPrimaryLanguage = "ku-tr"
	AttachToLangPrimaryRequestVNextPrimaryLanguageKv     AttachToLangPrimaryRequestVNextPrimaryLanguage = "kv"
	AttachToLangPrimaryRequestVNextPrimaryLanguageKw     AttachToLangPrimaryRequestVNextPrimaryLanguage = "kw"
	AttachToLangPrimaryRequestVNextPrimaryLanguageKwGB   AttachToLangPrimaryRequestVNextPrimaryLanguage = "kw-gb"
	AttachToLangPrimaryRequestVNextPrimaryLanguageKy     AttachToLangPrimaryRequestVNextPrimaryLanguage = "ky"
	AttachToLangPrimaryRequestVNextPrimaryLanguageKyKg   AttachToLangPrimaryRequestVNextPrimaryLanguage = "ky-kg"
	AttachToLangPrimaryRequestVNextPrimaryLanguageLa     AttachToLangPrimaryRequestVNextPrimaryLanguage = "la"
	AttachToLangPrimaryRequestVNextPrimaryLanguageLag    AttachToLangPrimaryRequestVNextPrimaryLanguage = "lag"
	AttachToLangPrimaryRequestVNextPrimaryLanguageLagTz  AttachToLangPrimaryRequestVNextPrimaryLanguage = "lag-tz"
	AttachToLangPrimaryRequestVNextPrimaryLanguageLb     AttachToLangPrimaryRequestVNextPrimaryLanguage = "lb"
	AttachToLangPrimaryRequestVNextPrimaryLanguageLbLu   AttachToLangPrimaryRequestVNextPrimaryLanguage = "lb-lu"
	AttachToLangPrimaryRequestVNextPrimaryLanguageLg     AttachToLangPrimaryRequestVNextPrimaryLanguage = "lg"
	AttachToLangPrimaryRequestVNextPrimaryLanguageLgUg   AttachToLangPrimaryRequestVNextPrimaryLanguage = "lg-ug"
	AttachToLangPrimaryRequestVNextPrimaryLanguageLi     AttachToLangPrimaryRequestVNextPrimaryLanguage = "li"
	AttachToLangPrimaryRequestVNextPrimaryLanguageLkt    AttachToLangPrimaryRequestVNextPrimaryLanguage = "lkt"
	AttachToLangPrimaryRequestVNextPrimaryLanguageLktUs  AttachToLangPrimaryRequestVNextPrimaryLanguage = "lkt-us"
	AttachToLangPrimaryRequestVNextPrimaryLanguageLn     AttachToLangPrimaryRequestVNextPrimaryLanguage = "ln"
	AttachToLangPrimaryRequestVNextPrimaryLanguageLnAo   AttachToLangPrimaryRequestVNextPrimaryLanguage = "ln-ao"
	AttachToLangPrimaryRequestVNextPrimaryLanguageLnCd   AttachToLangPrimaryRequestVNextPrimaryLanguage = "ln-cd"
	AttachToLangPrimaryRequestVNextPrimaryLanguageLnCf   AttachToLangPrimaryRequestVNextPrimaryLanguage = "ln-cf"
	AttachToLangPrimaryRequestVNextPrimaryLanguageLnCg   AttachToLangPrimaryRequestVNextPrimaryLanguage = "ln-cg"
	AttachToLangPrimaryRequestVNextPrimaryLanguageLo     AttachToLangPrimaryRequestVNextPrimaryLanguage = "lo"
	AttachToLangPrimaryRequestVNextPrimaryLanguageLoLa   AttachToLangPrimaryRequestVNextPrimaryLanguage = "lo-la"
	AttachToLangPrimaryRequestVNextPrimaryLanguageLrc    AttachToLangPrimaryRequestVNextPrimaryLanguage = "lrc"
	AttachToLangPrimaryRequestVNextPrimaryLanguageLrcIq  AttachToLangPrimaryRequestVNextPrimaryLanguage = "lrc-iq"
	AttachToLangPrimaryRequestVNextPrimaryLanguageLrcIr  AttachToLangPrimaryRequestVNextPrimaryLanguage = "lrc-ir"
	AttachToLangPrimaryRequestVNextPrimaryLanguageLt     AttachToLangPrimaryRequestVNextPrimaryLanguage = "lt"
	AttachToLangPrimaryRequestVNextPrimaryLanguageLtLt   AttachToLangPrimaryRequestVNextPrimaryLanguage = "lt-lt"
	AttachToLangPrimaryRequestVNextPrimaryLanguageLu     AttachToLangPrimaryRequestVNextPrimaryLanguage = "lu"
	AttachToLangPrimaryRequestVNextPrimaryLanguageLuCd   AttachToLangPrimaryRequestVNextPrimaryLanguage = "lu-cd"
	AttachToLangPrimaryRequestVNextPrimaryLanguageLuo    AttachToLangPrimaryRequestVNextPrimaryLanguage = "luo"
	AttachToLangPrimaryRequestVNextPrimaryLanguageLuoKe  AttachToLangPrimaryRequestVNextPrimaryLanguage = "luo-ke"
	AttachToLangPrimaryRequestVNextPrimaryLanguageLuy    AttachToLangPrimaryRequestVNextPrimaryLanguage = "luy"
	AttachToLangPrimaryRequestVNextPrimaryLanguageLuyKe  AttachToLangPrimaryRequestVNextPrimaryLanguage = "luy-ke"
	AttachToLangPrimaryRequestVNextPrimaryLanguageLv     AttachToLangPrimaryRequestVNextPrimaryLanguage = "lv"
	AttachToLangPrimaryRequestVNextPrimaryLanguageLvLv   AttachToLangPrimaryRequestVNextPrimaryLanguage = "lv-lv"
	AttachToLangPrimaryRequestVNextPrimaryLanguageMai    AttachToLangPrimaryRequestVNextPrimaryLanguage = "mai"
	AttachToLangPrimaryRequestVNextPrimaryLanguageMaiIn  AttachToLangPrimaryRequestVNextPrimaryLanguage = "mai-in"
	AttachToLangPrimaryRequestVNextPrimaryLanguageMas    AttachToLangPrimaryRequestVNextPrimaryLanguage = "mas"
	AttachToLangPrimaryRequestVNextPrimaryLanguageMasKe  AttachToLangPrimaryRequestVNextPrimaryLanguage = "mas-ke"
	AttachToLangPrimaryRequestVNextPrimaryLanguageMasTz  AttachToLangPrimaryRequestVNextPrimaryLanguage = "mas-tz"
	AttachToLangPrimaryRequestVNextPrimaryLanguageMdf    AttachToLangPrimaryRequestVNextPrimaryLanguage = "mdf"
	AttachToLangPrimaryRequestVNextPrimaryLanguageMdfRu  AttachToLangPrimaryRequestVNextPrimaryLanguage = "mdf-ru"
	AttachToLangPrimaryRequestVNextPrimaryLanguageMer    AttachToLangPrimaryRequestVNextPrimaryLanguage = "mer"
	AttachToLangPrimaryRequestVNextPrimaryLanguageMerKe  AttachToLangPrimaryRequestVNextPrimaryLanguage = "mer-ke"
	AttachToLangPrimaryRequestVNextPrimaryLanguageMfe    AttachToLangPrimaryRequestVNextPrimaryLanguage = "mfe"
	AttachToLangPrimaryRequestVNextPrimaryLanguageMfeMu  AttachToLangPrimaryRequestVNextPrimaryLanguage = "mfe-mu"
	AttachToLangPrimaryRequestVNextPrimaryLanguageMg     AttachToLangPrimaryRequestVNextPrimaryLanguage = "mg"
	AttachToLangPrimaryRequestVNextPrimaryLanguageMgMg   AttachToLangPrimaryRequestVNextPrimaryLanguage = "mg-mg"
	AttachToLangPrimaryRequestVNextPrimaryLanguageMgh    AttachToLangPrimaryRequestVNextPrimaryLanguage = "mgh"
	AttachToLangPrimaryRequestVNextPrimaryLanguageMghMz  AttachToLangPrimaryRequestVNextPrimaryLanguage = "mgh-mz"
	AttachToLangPrimaryRequestVNextPrimaryLanguageMgo    AttachToLangPrimaryRequestVNextPrimaryLanguage = "mgo"
	AttachToLangPrimaryRequestVNextPrimaryLanguageMgoCm  AttachToLangPrimaryRequestVNextPrimaryLanguage = "mgo-cm"
	AttachToLangPrimaryRequestVNextPrimaryLanguageMh     AttachToLangPrimaryRequestVNextPrimaryLanguage = "mh"
	AttachToLangPrimaryRequestVNextPrimaryLanguageMi     AttachToLangPrimaryRequestVNextPrimaryLanguage = "mi"
	AttachToLangPrimaryRequestVNextPrimaryLanguageMiNz   AttachToLangPrimaryRequestVNextPrimaryLanguage = "mi-nz"
	AttachToLangPrimaryRequestVNextPrimaryLanguageMk     AttachToLangPrimaryRequestVNextPrimaryLanguage = "mk"
	AttachToLangPrimaryRequestVNextPrimaryLanguageMkMk   AttachToLangPrimaryRequestVNextPrimaryLanguage = "mk-mk"
	AttachToLangPrimaryRequestVNextPrimaryLanguageMl     AttachToLangPrimaryRequestVNextPrimaryLanguage = "ml"
	AttachToLangPrimaryRequestVNextPrimaryLanguageMlIn   AttachToLangPrimaryRequestVNextPrimaryLanguage = "ml-in"
	AttachToLangPrimaryRequestVNextPrimaryLanguageMn     AttachToLangPrimaryRequestVNextPrimaryLanguage = "mn"
	AttachToLangPrimaryRequestVNextPrimaryLanguageMnMn   AttachToLangPrimaryRequestVNextPrimaryLanguage = "mn-mn"
	AttachToLangPrimaryRequestVNextPrimaryLanguageMni    AttachToLangPrimaryRequestVNextPrimaryLanguage = "mni"
	AttachToLangPrimaryRequestVNextPrimaryLanguageMniIn  AttachToLangPrimaryRequestVNextPrimaryLanguage = "mni-in"
	AttachToLangPrimaryRequestVNextPrimaryLanguageMr     AttachToLangPrimaryRequestVNextPrimaryLanguage = "mr"
	AttachToLangPrimaryRequestVNextPrimaryLanguageMrIn   AttachToLangPrimaryRequestVNextPrimaryLanguage = "mr-in"
	AttachToLangPrimaryRequestVNextPrimaryLanguageMs     AttachToLangPrimaryRequestVNextPrimaryLanguage = "ms"
	AttachToLangPrimaryRequestVNextPrimaryLanguageMsBn   AttachToLangPrimaryRequestVNextPrimaryLanguage = "ms-bn"
	AttachToLangPrimaryRequestVNextPrimaryLanguageMsID   AttachToLangPrimaryRequestVNextPrimaryLanguage = "ms-id"
	AttachToLangPrimaryRequestVNextPrimaryLanguageMsMy   AttachToLangPrimaryRequestVNextPrimaryLanguage = "ms-my"
	AttachToLangPrimaryRequestVNextPrimaryLanguageMsSg   AttachToLangPrimaryRequestVNextPrimaryLanguage = "ms-sg"
	AttachToLangPrimaryRequestVNextPrimaryLanguageMt     AttachToLangPrimaryRequestVNextPrimaryLanguage = "mt"
	AttachToLangPrimaryRequestVNextPrimaryLanguageMtMt   AttachToLangPrimaryRequestVNextPrimaryLanguage = "mt-mt"
	AttachToLangPrimaryRequestVNextPrimaryLanguageMua    AttachToLangPrimaryRequestVNextPrimaryLanguage = "mua"
	AttachToLangPrimaryRequestVNextPrimaryLanguageMuaCm  AttachToLangPrimaryRequestVNextPrimaryLanguage = "mua-cm"
	AttachToLangPrimaryRequestVNextPrimaryLanguageMy     AttachToLangPrimaryRequestVNextPrimaryLanguage = "my"
	AttachToLangPrimaryRequestVNextPrimaryLanguageMyMm   AttachToLangPrimaryRequestVNextPrimaryLanguage = "my-mm"
	AttachToLangPrimaryRequestVNextPrimaryLanguageMzn    AttachToLangPrimaryRequestVNextPrimaryLanguage = "mzn"
	AttachToLangPrimaryRequestVNextPrimaryLanguageMznIr  AttachToLangPrimaryRequestVNextPrimaryLanguage = "mzn-ir"
	AttachToLangPrimaryRequestVNextPrimaryLanguageNa     AttachToLangPrimaryRequestVNextPrimaryLanguage = "na"
	AttachToLangPrimaryRequestVNextPrimaryLanguageNaq    AttachToLangPrimaryRequestVNextPrimaryLanguage = "naq"
	AttachToLangPrimaryRequestVNextPrimaryLanguageNaqNa  AttachToLangPrimaryRequestVNextPrimaryLanguage = "naq-na"
	AttachToLangPrimaryRequestVNextPrimaryLanguageNb     AttachToLangPrimaryRequestVNextPrimaryLanguage = "nb"
	AttachToLangPrimaryRequestVNextPrimaryLanguageNbNo   AttachToLangPrimaryRequestVNextPrimaryLanguage = "nb-no"
	AttachToLangPrimaryRequestVNextPrimaryLanguageNbSj   AttachToLangPrimaryRequestVNextPrimaryLanguage = "nb-sj"
	AttachToLangPrimaryRequestVNextPrimaryLanguageNd     AttachToLangPrimaryRequestVNextPrimaryLanguage = "nd"
	AttachToLangPrimaryRequestVNextPrimaryLanguageNdZw   AttachToLangPrimaryRequestVNextPrimaryLanguage = "nd-zw"
	AttachToLangPrimaryRequestVNextPrimaryLanguageNds    AttachToLangPrimaryRequestVNextPrimaryLanguage = "nds"
	AttachToLangPrimaryRequestVNextPrimaryLanguageNdsDe  AttachToLangPrimaryRequestVNextPrimaryLanguage = "nds-de"
	AttachToLangPrimaryRequestVNextPrimaryLanguageNdsNl  AttachToLangPrimaryRequestVNextPrimaryLanguage = "nds-nl"
	AttachToLangPrimaryRequestVNextPrimaryLanguageNe     AttachToLangPrimaryRequestVNextPrimaryLanguage = "ne"
	AttachToLangPrimaryRequestVNextPrimaryLanguageNeIn   AttachToLangPrimaryRequestVNextPrimaryLanguage = "ne-in"
	AttachToLangPrimaryRequestVNextPrimaryLanguageNeNp   AttachToLangPrimaryRequestVNextPrimaryLanguage = "ne-np"
	AttachToLangPrimaryRequestVNextPrimaryLanguageNg     AttachToLangPrimaryRequestVNextPrimaryLanguage = "ng"
	AttachToLangPrimaryRequestVNextPrimaryLanguageNl     AttachToLangPrimaryRequestVNextPrimaryLanguage = "nl"
	AttachToLangPrimaryRequestVNextPrimaryLanguageNlAw   AttachToLangPrimaryRequestVNextPrimaryLanguage = "nl-aw"
	AttachToLangPrimaryRequestVNextPrimaryLanguageNlBe   AttachToLangPrimaryRequestVNextPrimaryLanguage = "nl-be"
	AttachToLangPrimaryRequestVNextPrimaryLanguageNlBq   AttachToLangPrimaryRequestVNextPrimaryLanguage = "nl-bq"
	AttachToLangPrimaryRequestVNextPrimaryLanguageNlCh   AttachToLangPrimaryRequestVNextPrimaryLanguage = "nl-ch"
	AttachToLangPrimaryRequestVNextPrimaryLanguageNlCw   AttachToLangPrimaryRequestVNextPrimaryLanguage = "nl-cw"
	AttachToLangPrimaryRequestVNextPrimaryLanguageNlLu   AttachToLangPrimaryRequestVNextPrimaryLanguage = "nl-lu"
	AttachToLangPrimaryRequestVNextPrimaryLanguageNlNl   AttachToLangPrimaryRequestVNextPrimaryLanguage = "nl-nl"
	AttachToLangPrimaryRequestVNextPrimaryLanguageNlSr   AttachToLangPrimaryRequestVNextPrimaryLanguage = "nl-sr"
	AttachToLangPrimaryRequestVNextPrimaryLanguageNlSx   AttachToLangPrimaryRequestVNextPrimaryLanguage = "nl-sx"
	AttachToLangPrimaryRequestVNextPrimaryLanguageNmg    AttachToLangPrimaryRequestVNextPrimaryLanguage = "nmg"
	AttachToLangPrimaryRequestVNextPrimaryLanguageNmgCm  AttachToLangPrimaryRequestVNextPrimaryLanguage = "nmg-cm"
	AttachToLangPrimaryRequestVNextPrimaryLanguageNn     AttachToLangPrimaryRequestVNextPrimaryLanguage = "nn"
	AttachToLangPrimaryRequestVNextPrimaryLanguageNnNo   AttachToLangPrimaryRequestVNextPrimaryLanguage = "nn-no"
	AttachToLangPrimaryRequestVNextPrimaryLanguageNnh    AttachToLangPrimaryRequestVNextPrimaryLanguage = "nnh"
	AttachToLangPrimaryRequestVNextPrimaryLanguageNnhCm  AttachToLangPrimaryRequestVNextPrimaryLanguage = "nnh-cm"
	AttachToLangPrimaryRequestVNextPrimaryLanguageNo     AttachToLangPrimaryRequestVNextPrimaryLanguage = "no"
	AttachToLangPrimaryRequestVNextPrimaryLanguageNoNo   AttachToLangPrimaryRequestVNextPrimaryLanguage = "no-no"
	AttachToLangPrimaryRequestVNextPrimaryLanguageNr     AttachToLangPrimaryRequestVNextPrimaryLanguage = "nr"
	AttachToLangPrimaryRequestVNextPrimaryLanguageNus    AttachToLangPrimaryRequestVNextPrimaryLanguage = "nus"
	AttachToLangPrimaryRequestVNextPrimaryLanguageNusSS  AttachToLangPrimaryRequestVNextPrimaryLanguage = "nus-ss"
	AttachToLangPrimaryRequestVNextPrimaryLanguageNv     AttachToLangPrimaryRequestVNextPrimaryLanguage = "nv"
	AttachToLangPrimaryRequestVNextPrimaryLanguageNy     AttachToLangPrimaryRequestVNextPrimaryLanguage = "ny"
	AttachToLangPrimaryRequestVNextPrimaryLanguageNyn    AttachToLangPrimaryRequestVNextPrimaryLanguage = "nyn"
	AttachToLangPrimaryRequestVNextPrimaryLanguageNynUg  AttachToLangPrimaryRequestVNextPrimaryLanguage = "nyn-ug"
	AttachToLangPrimaryRequestVNextPrimaryLanguageOc     AttachToLangPrimaryRequestVNextPrimaryLanguage = "oc"
	AttachToLangPrimaryRequestVNextPrimaryLanguageOcEs   AttachToLangPrimaryRequestVNextPrimaryLanguage = "oc-es"
	AttachToLangPrimaryRequestVNextPrimaryLanguageOcFr   AttachToLangPrimaryRequestVNextPrimaryLanguage = "oc-fr"
	AttachToLangPrimaryRequestVNextPrimaryLanguageOj     AttachToLangPrimaryRequestVNextPrimaryLanguage = "oj"
	AttachToLangPrimaryRequestVNextPrimaryLanguageOm     AttachToLangPrimaryRequestVNextPrimaryLanguage = "om"
	AttachToLangPrimaryRequestVNextPrimaryLanguageOmEt   AttachToLangPrimaryRequestVNextPrimaryLanguage = "om-et"
	AttachToLangPrimaryRequestVNextPrimaryLanguageOmKe   AttachToLangPrimaryRequestVNextPrimaryLanguage = "om-ke"
	AttachToLangPrimaryRequestVNextPrimaryLanguageOr     AttachToLangPrimaryRequestVNextPrimaryLanguage = "or"
	AttachToLangPrimaryRequestVNextPrimaryLanguageOrIn   AttachToLangPrimaryRequestVNextPrimaryLanguage = "or-in"
	AttachToLangPrimaryRequestVNextPrimaryLanguageOs     AttachToLangPrimaryRequestVNextPrimaryLanguage = "os"
	AttachToLangPrimaryRequestVNextPrimaryLanguageOsGe   AttachToLangPrimaryRequestVNextPrimaryLanguage = "os-ge"
	AttachToLangPrimaryRequestVNextPrimaryLanguageOsRu   AttachToLangPrimaryRequestVNextPrimaryLanguage = "os-ru"
	AttachToLangPrimaryRequestVNextPrimaryLanguagePa     AttachToLangPrimaryRequestVNextPrimaryLanguage = "pa"
	AttachToLangPrimaryRequestVNextPrimaryLanguagePaIn   AttachToLangPrimaryRequestVNextPrimaryLanguage = "pa-in"
	AttachToLangPrimaryRequestVNextPrimaryLanguagePaPk   AttachToLangPrimaryRequestVNextPrimaryLanguage = "pa-pk"
	AttachToLangPrimaryRequestVNextPrimaryLanguagePcm    AttachToLangPrimaryRequestVNextPrimaryLanguage = "pcm"
	AttachToLangPrimaryRequestVNextPrimaryLanguagePcmNg  AttachToLangPrimaryRequestVNextPrimaryLanguage = "pcm-ng"
	AttachToLangPrimaryRequestVNextPrimaryLanguagePi     AttachToLangPrimaryRequestVNextPrimaryLanguage = "pi"
	AttachToLangPrimaryRequestVNextPrimaryLanguagePis    AttachToLangPrimaryRequestVNextPrimaryLanguage = "pis"
	AttachToLangPrimaryRequestVNextPrimaryLanguagePisSb  AttachToLangPrimaryRequestVNextPrimaryLanguage = "pis-sb"
	AttachToLangPrimaryRequestVNextPrimaryLanguagePl     AttachToLangPrimaryRequestVNextPrimaryLanguage = "pl"
	AttachToLangPrimaryRequestVNextPrimaryLanguagePlPl   AttachToLangPrimaryRequestVNextPrimaryLanguage = "pl-pl"
	AttachToLangPrimaryRequestVNextPrimaryLanguagePrg    AttachToLangPrimaryRequestVNextPrimaryLanguage = "prg"
	AttachToLangPrimaryRequestVNextPrimaryLanguagePrg001 AttachToLangPrimaryRequestVNextPrimaryLanguage = "prg-001"
	AttachToLangPrimaryRequestVNextPrimaryLanguagePs     AttachToLangPrimaryRequestVNextPrimaryLanguage = "ps"
	AttachToLangPrimaryRequestVNextPrimaryLanguagePsAf   AttachToLangPrimaryRequestVNextPrimaryLanguage = "ps-af"
	AttachToLangPrimaryRequestVNextPrimaryLanguagePsPk   AttachToLangPrimaryRequestVNextPrimaryLanguage = "ps-pk"
	AttachToLangPrimaryRequestVNextPrimaryLanguagePt     AttachToLangPrimaryRequestVNextPrimaryLanguage = "pt"
	AttachToLangPrimaryRequestVNextPrimaryLanguagePtAo   AttachToLangPrimaryRequestVNextPrimaryLanguage = "pt-ao"
	AttachToLangPrimaryRequestVNextPrimaryLanguagePtBr   AttachToLangPrimaryRequestVNextPrimaryLanguage = "pt-br"
	AttachToLangPrimaryRequestVNextPrimaryLanguagePtCh   AttachToLangPrimaryRequestVNextPrimaryLanguage = "pt-ch"
	AttachToLangPrimaryRequestVNextPrimaryLanguagePtCv   AttachToLangPrimaryRequestVNextPrimaryLanguage = "pt-cv"
	AttachToLangPrimaryRequestVNextPrimaryLanguagePtGq   AttachToLangPrimaryRequestVNextPrimaryLanguage = "pt-gq"
	AttachToLangPrimaryRequestVNextPrimaryLanguagePtGw   AttachToLangPrimaryRequestVNextPrimaryLanguage = "pt-gw"
	AttachToLangPrimaryRequestVNextPrimaryLanguagePtLu   AttachToLangPrimaryRequestVNextPrimaryLanguage = "pt-lu"
	AttachToLangPrimaryRequestVNextPrimaryLanguagePtMo   AttachToLangPrimaryRequestVNextPrimaryLanguage = "pt-mo"
	AttachToLangPrimaryRequestVNextPrimaryLanguagePtMz   AttachToLangPrimaryRequestVNextPrimaryLanguage = "pt-mz"
	AttachToLangPrimaryRequestVNextPrimaryLanguagePtPt   AttachToLangPrimaryRequestVNextPrimaryLanguage = "pt-pt"
	AttachToLangPrimaryRequestVNextPrimaryLanguagePtSt   AttachToLangPrimaryRequestVNextPrimaryLanguage = "pt-st"
	AttachToLangPrimaryRequestVNextPrimaryLanguagePtTl   AttachToLangPrimaryRequestVNextPrimaryLanguage = "pt-tl"
	AttachToLangPrimaryRequestVNextPrimaryLanguageQu     AttachToLangPrimaryRequestVNextPrimaryLanguage = "qu"
	AttachToLangPrimaryRequestVNextPrimaryLanguageQuBo   AttachToLangPrimaryRequestVNextPrimaryLanguage = "qu-bo"
	AttachToLangPrimaryRequestVNextPrimaryLanguageQuEc   AttachToLangPrimaryRequestVNextPrimaryLanguage = "qu-ec"
	AttachToLangPrimaryRequestVNextPrimaryLanguageQuPe   AttachToLangPrimaryRequestVNextPrimaryLanguage = "qu-pe"
	AttachToLangPrimaryRequestVNextPrimaryLanguageRaj    AttachToLangPrimaryRequestVNextPrimaryLanguage = "raj"
	AttachToLangPrimaryRequestVNextPrimaryLanguageRajIn  AttachToLangPrimaryRequestVNextPrimaryLanguage = "raj-in"
	AttachToLangPrimaryRequestVNextPrimaryLanguageRm     AttachToLangPrimaryRequestVNextPrimaryLanguage = "rm"
	AttachToLangPrimaryRequestVNextPrimaryLanguageRmCh   AttachToLangPrimaryRequestVNextPrimaryLanguage = "rm-ch"
	AttachToLangPrimaryRequestVNextPrimaryLanguageRn     AttachToLangPrimaryRequestVNextPrimaryLanguage = "rn"
	AttachToLangPrimaryRequestVNextPrimaryLanguageRnBi   AttachToLangPrimaryRequestVNextPrimaryLanguage = "rn-bi"
	AttachToLangPrimaryRequestVNextPrimaryLanguageRo     AttachToLangPrimaryRequestVNextPrimaryLanguage = "ro"
	AttachToLangPrimaryRequestVNextPrimaryLanguageRoMd   AttachToLangPrimaryRequestVNextPrimaryLanguage = "ro-md"
	AttachToLangPrimaryRequestVNextPrimaryLanguageRoRo   AttachToLangPrimaryRequestVNextPrimaryLanguage = "ro-ro"
	AttachToLangPrimaryRequestVNextPrimaryLanguageRof    AttachToLangPrimaryRequestVNextPrimaryLanguage = "rof"
	AttachToLangPrimaryRequestVNextPrimaryLanguageRofTz  AttachToLangPrimaryRequestVNextPrimaryLanguage = "rof-tz"
	AttachToLangPrimaryRequestVNextPrimaryLanguageRu     AttachToLangPrimaryRequestVNextPrimaryLanguage = "ru"
	AttachToLangPrimaryRequestVNextPrimaryLanguageRuBy   AttachToLangPrimaryRequestVNextPrimaryLanguage = "ru-by"
	AttachToLangPrimaryRequestVNextPrimaryLanguageRuKg   AttachToLangPrimaryRequestVNextPrimaryLanguage = "ru-kg"
	AttachToLangPrimaryRequestVNextPrimaryLanguageRuKz   AttachToLangPrimaryRequestVNextPrimaryLanguage = "ru-kz"
	AttachToLangPrimaryRequestVNextPrimaryLanguageRuMd   AttachToLangPrimaryRequestVNextPrimaryLanguage = "ru-md"
	AttachToLangPrimaryRequestVNextPrimaryLanguageRuRu   AttachToLangPrimaryRequestVNextPrimaryLanguage = "ru-ru"
	AttachToLangPrimaryRequestVNextPrimaryLanguageRuUa   AttachToLangPrimaryRequestVNextPrimaryLanguage = "ru-ua"
	AttachToLangPrimaryRequestVNextPrimaryLanguageRw     AttachToLangPrimaryRequestVNextPrimaryLanguage = "rw"
	AttachToLangPrimaryRequestVNextPrimaryLanguageRwRw   AttachToLangPrimaryRequestVNextPrimaryLanguage = "rw-rw"
	AttachToLangPrimaryRequestVNextPrimaryLanguageRwk    AttachToLangPrimaryRequestVNextPrimaryLanguage = "rwk"
	AttachToLangPrimaryRequestVNextPrimaryLanguageRwkTz  AttachToLangPrimaryRequestVNextPrimaryLanguage = "rwk-tz"
	AttachToLangPrimaryRequestVNextPrimaryLanguageSa     AttachToLangPrimaryRequestVNextPrimaryLanguage = "sa"
	AttachToLangPrimaryRequestVNextPrimaryLanguageSaIn   AttachToLangPrimaryRequestVNextPrimaryLanguage = "sa-in"
	AttachToLangPrimaryRequestVNextPrimaryLanguageSah    AttachToLangPrimaryRequestVNextPrimaryLanguage = "sah"
	AttachToLangPrimaryRequestVNextPrimaryLanguageSahRu  AttachToLangPrimaryRequestVNextPrimaryLanguage = "sah-ru"
	AttachToLangPrimaryRequestVNextPrimaryLanguageSaq    AttachToLangPrimaryRequestVNextPrimaryLanguage = "saq"
	AttachToLangPrimaryRequestVNextPrimaryLanguageSaqKe  AttachToLangPrimaryRequestVNextPrimaryLanguage = "saq-ke"
	AttachToLangPrimaryRequestVNextPrimaryLanguageSat    AttachToLangPrimaryRequestVNextPrimaryLanguage = "sat"
	AttachToLangPrimaryRequestVNextPrimaryLanguageSatIn  AttachToLangPrimaryRequestVNextPrimaryLanguage = "sat-in"
	AttachToLangPrimaryRequestVNextPrimaryLanguageSbp    AttachToLangPrimaryRequestVNextPrimaryLanguage = "sbp"
	AttachToLangPrimaryRequestVNextPrimaryLanguageSbpTz  AttachToLangPrimaryRequestVNextPrimaryLanguage = "sbp-tz"
	AttachToLangPrimaryRequestVNextPrimaryLanguageSc     AttachToLangPrimaryRequestVNextPrimaryLanguage = "sc"
	AttachToLangPrimaryRequestVNextPrimaryLanguageScIt   AttachToLangPrimaryRequestVNextPrimaryLanguage = "sc-it"
	AttachToLangPrimaryRequestVNextPrimaryLanguageSd     AttachToLangPrimaryRequestVNextPrimaryLanguage = "sd"
	AttachToLangPrimaryRequestVNextPrimaryLanguageSdIn   AttachToLangPrimaryRequestVNextPrimaryLanguage = "sd-in"
	AttachToLangPrimaryRequestVNextPrimaryLanguageSdPk   AttachToLangPrimaryRequestVNextPrimaryLanguage = "sd-pk"
	AttachToLangPrimaryRequestVNextPrimaryLanguageSe     AttachToLangPrimaryRequestVNextPrimaryLanguage = "se"
	AttachToLangPrimaryRequestVNextPrimaryLanguageSeFi   AttachToLangPrimaryRequestVNextPrimaryLanguage = "se-fi"
	AttachToLangPrimaryRequestVNextPrimaryLanguageSeNo   AttachToLangPrimaryRequestVNextPrimaryLanguage = "se-no"
	AttachToLangPrimaryRequestVNextPrimaryLanguageSeSe   AttachToLangPrimaryRequestVNextPrimaryLanguage = "se-se"
	AttachToLangPrimaryRequestVNextPrimaryLanguageSeh    AttachToLangPrimaryRequestVNextPrimaryLanguage = "seh"
	AttachToLangPrimaryRequestVNextPrimaryLanguageSehMz  AttachToLangPrimaryRequestVNextPrimaryLanguage = "seh-mz"
	AttachToLangPrimaryRequestVNextPrimaryLanguageSes    AttachToLangPrimaryRequestVNextPrimaryLanguage = "ses"
	AttachToLangPrimaryRequestVNextPrimaryLanguageSesMl  AttachToLangPrimaryRequestVNextPrimaryLanguage = "ses-ml"
	AttachToLangPrimaryRequestVNextPrimaryLanguageSg     AttachToLangPrimaryRequestVNextPrimaryLanguage = "sg"
	AttachToLangPrimaryRequestVNextPrimaryLanguageSgCf   AttachToLangPrimaryRequestVNextPrimaryLanguage = "sg-cf"
	AttachToLangPrimaryRequestVNextPrimaryLanguageShi    AttachToLangPrimaryRequestVNextPrimaryLanguage = "shi"
	AttachToLangPrimaryRequestVNextPrimaryLanguageShiMa  AttachToLangPrimaryRequestVNextPrimaryLanguage = "shi-ma"
	AttachToLangPrimaryRequestVNextPrimaryLanguageSi     AttachToLangPrimaryRequestVNextPrimaryLanguage = "si"
	AttachToLangPrimaryRequestVNextPrimaryLanguageSiLk   AttachToLangPrimaryRequestVNextPrimaryLanguage = "si-lk"
	AttachToLangPrimaryRequestVNextPrimaryLanguageSk     AttachToLangPrimaryRequestVNextPrimaryLanguage = "sk"
	AttachToLangPrimaryRequestVNextPrimaryLanguageSkSk   AttachToLangPrimaryRequestVNextPrimaryLanguage = "sk-sk"
	AttachToLangPrimaryRequestVNextPrimaryLanguageSl     AttachToLangPrimaryRequestVNextPrimaryLanguage = "sl"
	AttachToLangPrimaryRequestVNextPrimaryLanguageSlSi   AttachToLangPrimaryRequestVNextPrimaryLanguage = "sl-si"
	AttachToLangPrimaryRequestVNextPrimaryLanguageSm     AttachToLangPrimaryRequestVNextPrimaryLanguage = "sm"
	AttachToLangPrimaryRequestVNextPrimaryLanguageSmn    AttachToLangPrimaryRequestVNextPrimaryLanguage = "smn"
	AttachToLangPrimaryRequestVNextPrimaryLanguageSmnFi  AttachToLangPrimaryRequestVNextPrimaryLanguage = "smn-fi"
	AttachToLangPrimaryRequestVNextPrimaryLanguageSMS    AttachToLangPrimaryRequestVNextPrimaryLanguage = "sms"
	AttachToLangPrimaryRequestVNextPrimaryLanguageSMSFi  AttachToLangPrimaryRequestVNextPrimaryLanguage = "sms-fi"
	AttachToLangPrimaryRequestVNextPrimaryLanguageSn     AttachToLangPrimaryRequestVNextPrimaryLanguage = "sn"
	AttachToLangPrimaryRequestVNextPrimaryLanguageSnZw   AttachToLangPrimaryRequestVNextPrimaryLanguage = "sn-zw"
	AttachToLangPrimaryRequestVNextPrimaryLanguageSo     AttachToLangPrimaryRequestVNextPrimaryLanguage = "so"
	AttachToLangPrimaryRequestVNextPrimaryLanguageSoDj   AttachToLangPrimaryRequestVNextPrimaryLanguage = "so-dj"
	AttachToLangPrimaryRequestVNextPrimaryLanguageSoEt   AttachToLangPrimaryRequestVNextPrimaryLanguage = "so-et"
	AttachToLangPrimaryRequestVNextPrimaryLanguageSoKe   AttachToLangPrimaryRequestVNextPrimaryLanguage = "so-ke"
	AttachToLangPrimaryRequestVNextPrimaryLanguageSoSo   AttachToLangPrimaryRequestVNextPrimaryLanguage = "so-so"
	AttachToLangPrimaryRequestVNextPrimaryLanguageSq     AttachToLangPrimaryRequestVNextPrimaryLanguage = "sq"
	AttachToLangPrimaryRequestVNextPrimaryLanguageSqAl   AttachToLangPrimaryRequestVNextPrimaryLanguage = "sq-al"
	AttachToLangPrimaryRequestVNextPrimaryLanguageSqMk   AttachToLangPrimaryRequestVNextPrimaryLanguage = "sq-mk"
	AttachToLangPrimaryRequestVNextPrimaryLanguageSqXk   AttachToLangPrimaryRequestVNextPrimaryLanguage = "sq-xk"
	AttachToLangPrimaryRequestVNextPrimaryLanguageSr     AttachToLangPrimaryRequestVNextPrimaryLanguage = "sr"
	AttachToLangPrimaryRequestVNextPrimaryLanguageSrBa   AttachToLangPrimaryRequestVNextPrimaryLanguage = "sr-ba"
	AttachToLangPrimaryRequestVNextPrimaryLanguageSrCs   AttachToLangPrimaryRequestVNextPrimaryLanguage = "sr-cs"
	AttachToLangPrimaryRequestVNextPrimaryLanguageSrMe   AttachToLangPrimaryRequestVNextPrimaryLanguage = "sr-me"
	AttachToLangPrimaryRequestVNextPrimaryLanguageSrRs   AttachToLangPrimaryRequestVNextPrimaryLanguage = "sr-rs"
	AttachToLangPrimaryRequestVNextPrimaryLanguageSrXk   AttachToLangPrimaryRequestVNextPrimaryLanguage = "sr-xk"
	AttachToLangPrimaryRequestVNextPrimaryLanguageSS     AttachToLangPrimaryRequestVNextPrimaryLanguage = "ss"
	AttachToLangPrimaryRequestVNextPrimaryLanguageSt     AttachToLangPrimaryRequestVNextPrimaryLanguage = "st"
	AttachToLangPrimaryRequestVNextPrimaryLanguageSu     AttachToLangPrimaryRequestVNextPrimaryLanguage = "su"
	AttachToLangPrimaryRequestVNextPrimaryLanguageSuID   AttachToLangPrimaryRequestVNextPrimaryLanguage = "su-id"
	AttachToLangPrimaryRequestVNextPrimaryLanguageSv     AttachToLangPrimaryRequestVNextPrimaryLanguage = "sv"
	AttachToLangPrimaryRequestVNextPrimaryLanguageSvAx   AttachToLangPrimaryRequestVNextPrimaryLanguage = "sv-ax"
	AttachToLangPrimaryRequestVNextPrimaryLanguageSvFi   AttachToLangPrimaryRequestVNextPrimaryLanguage = "sv-fi"
	AttachToLangPrimaryRequestVNextPrimaryLanguageSvSe   AttachToLangPrimaryRequestVNextPrimaryLanguage = "sv-se"
	AttachToLangPrimaryRequestVNextPrimaryLanguageSw     AttachToLangPrimaryRequestVNextPrimaryLanguage = "sw"
	AttachToLangPrimaryRequestVNextPrimaryLanguageSwCd   AttachToLangPrimaryRequestVNextPrimaryLanguage = "sw-cd"
	AttachToLangPrimaryRequestVNextPrimaryLanguageSwKe   AttachToLangPrimaryRequestVNextPrimaryLanguage = "sw-ke"
	AttachToLangPrimaryRequestVNextPrimaryLanguageSwTz   AttachToLangPrimaryRequestVNextPrimaryLanguage = "sw-tz"
	AttachToLangPrimaryRequestVNextPrimaryLanguageSwUg   AttachToLangPrimaryRequestVNextPrimaryLanguage = "sw-ug"
	AttachToLangPrimaryRequestVNextPrimaryLanguageSy     AttachToLangPrimaryRequestVNextPrimaryLanguage = "sy"
	AttachToLangPrimaryRequestVNextPrimaryLanguageTa     AttachToLangPrimaryRequestVNextPrimaryLanguage = "ta"
	AttachToLangPrimaryRequestVNextPrimaryLanguageTaIn   AttachToLangPrimaryRequestVNextPrimaryLanguage = "ta-in"
	AttachToLangPrimaryRequestVNextPrimaryLanguageTaLk   AttachToLangPrimaryRequestVNextPrimaryLanguage = "ta-lk"
	AttachToLangPrimaryRequestVNextPrimaryLanguageTaMy   AttachToLangPrimaryRequestVNextPrimaryLanguage = "ta-my"
	AttachToLangPrimaryRequestVNextPrimaryLanguageTaSg   AttachToLangPrimaryRequestVNextPrimaryLanguage = "ta-sg"
	AttachToLangPrimaryRequestVNextPrimaryLanguageTe     AttachToLangPrimaryRequestVNextPrimaryLanguage = "te"
	AttachToLangPrimaryRequestVNextPrimaryLanguageTeIn   AttachToLangPrimaryRequestVNextPrimaryLanguage = "te-in"
	AttachToLangPrimaryRequestVNextPrimaryLanguageTeo    AttachToLangPrimaryRequestVNextPrimaryLanguage = "teo"
	AttachToLangPrimaryRequestVNextPrimaryLanguageTeoKe  AttachToLangPrimaryRequestVNextPrimaryLanguage = "teo-ke"
	AttachToLangPrimaryRequestVNextPrimaryLanguageTeoUg  AttachToLangPrimaryRequestVNextPrimaryLanguage = "teo-ug"
	AttachToLangPrimaryRequestVNextPrimaryLanguageTg     AttachToLangPrimaryRequestVNextPrimaryLanguage = "tg"
	AttachToLangPrimaryRequestVNextPrimaryLanguageTgTj   AttachToLangPrimaryRequestVNextPrimaryLanguage = "tg-tj"
	AttachToLangPrimaryRequestVNextPrimaryLanguageTh     AttachToLangPrimaryRequestVNextPrimaryLanguage = "th"
	AttachToLangPrimaryRequestVNextPrimaryLanguageThTh   AttachToLangPrimaryRequestVNextPrimaryLanguage = "th-th"
	AttachToLangPrimaryRequestVNextPrimaryLanguageTi     AttachToLangPrimaryRequestVNextPrimaryLanguage = "ti"
	AttachToLangPrimaryRequestVNextPrimaryLanguageTiEr   AttachToLangPrimaryRequestVNextPrimaryLanguage = "ti-er"
	AttachToLangPrimaryRequestVNextPrimaryLanguageTiEt   AttachToLangPrimaryRequestVNextPrimaryLanguage = "ti-et"
	AttachToLangPrimaryRequestVNextPrimaryLanguageTk     AttachToLangPrimaryRequestVNextPrimaryLanguage = "tk"
	AttachToLangPrimaryRequestVNextPrimaryLanguageTkTm   AttachToLangPrimaryRequestVNextPrimaryLanguage = "tk-tm"
	AttachToLangPrimaryRequestVNextPrimaryLanguageTl     AttachToLangPrimaryRequestVNextPrimaryLanguage = "tl"
	AttachToLangPrimaryRequestVNextPrimaryLanguageTn     AttachToLangPrimaryRequestVNextPrimaryLanguage = "tn"
	AttachToLangPrimaryRequestVNextPrimaryLanguageTo     AttachToLangPrimaryRequestVNextPrimaryLanguage = "to"
	AttachToLangPrimaryRequestVNextPrimaryLanguageToTo   AttachToLangPrimaryRequestVNextPrimaryLanguage = "to-to"
	AttachToLangPrimaryRequestVNextPrimaryLanguageTok    AttachToLangPrimaryRequestVNextPrimaryLanguage = "tok"
	AttachToLangPrimaryRequestVNextPrimaryLanguageTok001 AttachToLangPrimaryRequestVNextPrimaryLanguage = "tok-001"
	AttachToLangPrimaryRequestVNextPrimaryLanguageTr     AttachToLangPrimaryRequestVNextPrimaryLanguage = "tr"
	AttachToLangPrimaryRequestVNextPrimaryLanguageTrCy   AttachToLangPrimaryRequestVNextPrimaryLanguage = "tr-cy"
	AttachToLangPrimaryRequestVNextPrimaryLanguageTrTr   AttachToLangPrimaryRequestVNextPrimaryLanguage = "tr-tr"
	AttachToLangPrimaryRequestVNextPrimaryLanguageTs     AttachToLangPrimaryRequestVNextPrimaryLanguage = "ts"
	AttachToLangPrimaryRequestVNextPrimaryLanguageTt     AttachToLangPrimaryRequestVNextPrimaryLanguage = "tt"
	AttachToLangPrimaryRequestVNextPrimaryLanguageTtRu   AttachToLangPrimaryRequestVNextPrimaryLanguage = "tt-ru"
	AttachToLangPrimaryRequestVNextPrimaryLanguageTw     AttachToLangPrimaryRequestVNextPrimaryLanguage = "tw"
	AttachToLangPrimaryRequestVNextPrimaryLanguageTwq    AttachToLangPrimaryRequestVNextPrimaryLanguage = "twq"
	AttachToLangPrimaryRequestVNextPrimaryLanguageTwqNe  AttachToLangPrimaryRequestVNextPrimaryLanguage = "twq-ne"
	AttachToLangPrimaryRequestVNextPrimaryLanguageTy     AttachToLangPrimaryRequestVNextPrimaryLanguage = "ty"
	AttachToLangPrimaryRequestVNextPrimaryLanguageTzm    AttachToLangPrimaryRequestVNextPrimaryLanguage = "tzm"
	AttachToLangPrimaryRequestVNextPrimaryLanguageTzmMa  AttachToLangPrimaryRequestVNextPrimaryLanguage = "tzm-ma"
	AttachToLangPrimaryRequestVNextPrimaryLanguageUg     AttachToLangPrimaryRequestVNextPrimaryLanguage = "ug"
	AttachToLangPrimaryRequestVNextPrimaryLanguageUgCn   AttachToLangPrimaryRequestVNextPrimaryLanguage = "ug-cn"
	AttachToLangPrimaryRequestVNextPrimaryLanguageUk     AttachToLangPrimaryRequestVNextPrimaryLanguage = "uk"
	AttachToLangPrimaryRequestVNextPrimaryLanguageUkUa   AttachToLangPrimaryRequestVNextPrimaryLanguage = "uk-ua"
	AttachToLangPrimaryRequestVNextPrimaryLanguageUr     AttachToLangPrimaryRequestVNextPrimaryLanguage = "ur"
	AttachToLangPrimaryRequestVNextPrimaryLanguageUrIn   AttachToLangPrimaryRequestVNextPrimaryLanguage = "ur-in"
	AttachToLangPrimaryRequestVNextPrimaryLanguageUrPk   AttachToLangPrimaryRequestVNextPrimaryLanguage = "ur-pk"
	AttachToLangPrimaryRequestVNextPrimaryLanguageUz     AttachToLangPrimaryRequestVNextPrimaryLanguage = "uz"
	AttachToLangPrimaryRequestVNextPrimaryLanguageUzAf   AttachToLangPrimaryRequestVNextPrimaryLanguage = "uz-af"
	AttachToLangPrimaryRequestVNextPrimaryLanguageUzUz   AttachToLangPrimaryRequestVNextPrimaryLanguage = "uz-uz"
	AttachToLangPrimaryRequestVNextPrimaryLanguageVai    AttachToLangPrimaryRequestVNextPrimaryLanguage = "vai"
	AttachToLangPrimaryRequestVNextPrimaryLanguageVaiLr  AttachToLangPrimaryRequestVNextPrimaryLanguage = "vai-lr"
	AttachToLangPrimaryRequestVNextPrimaryLanguageVe     AttachToLangPrimaryRequestVNextPrimaryLanguage = "ve"
	AttachToLangPrimaryRequestVNextPrimaryLanguageVi     AttachToLangPrimaryRequestVNextPrimaryLanguage = "vi"
	AttachToLangPrimaryRequestVNextPrimaryLanguageViVn   AttachToLangPrimaryRequestVNextPrimaryLanguage = "vi-vn"
	AttachToLangPrimaryRequestVNextPrimaryLanguageVo     AttachToLangPrimaryRequestVNextPrimaryLanguage = "vo"
	AttachToLangPrimaryRequestVNextPrimaryLanguageVo001  AttachToLangPrimaryRequestVNextPrimaryLanguage = "vo-001"
	AttachToLangPrimaryRequestVNextPrimaryLanguageVun    AttachToLangPrimaryRequestVNextPrimaryLanguage = "vun"
	AttachToLangPrimaryRequestVNextPrimaryLanguageVunTz  AttachToLangPrimaryRequestVNextPrimaryLanguage = "vun-tz"
	AttachToLangPrimaryRequestVNextPrimaryLanguageWa     AttachToLangPrimaryRequestVNextPrimaryLanguage = "wa"
	AttachToLangPrimaryRequestVNextPrimaryLanguageWae    AttachToLangPrimaryRequestVNextPrimaryLanguage = "wae"
	AttachToLangPrimaryRequestVNextPrimaryLanguageWaeCh  AttachToLangPrimaryRequestVNextPrimaryLanguage = "wae-ch"
	AttachToLangPrimaryRequestVNextPrimaryLanguageWo     AttachToLangPrimaryRequestVNextPrimaryLanguage = "wo"
	AttachToLangPrimaryRequestVNextPrimaryLanguageWoSn   AttachToLangPrimaryRequestVNextPrimaryLanguage = "wo-sn"
	AttachToLangPrimaryRequestVNextPrimaryLanguageXh     AttachToLangPrimaryRequestVNextPrimaryLanguage = "xh"
	AttachToLangPrimaryRequestVNextPrimaryLanguageXhZa   AttachToLangPrimaryRequestVNextPrimaryLanguage = "xh-za"
	AttachToLangPrimaryRequestVNextPrimaryLanguageXog    AttachToLangPrimaryRequestVNextPrimaryLanguage = "xog"
	AttachToLangPrimaryRequestVNextPrimaryLanguageXogUg  AttachToLangPrimaryRequestVNextPrimaryLanguage = "xog-ug"
	AttachToLangPrimaryRequestVNextPrimaryLanguageYav    AttachToLangPrimaryRequestVNextPrimaryLanguage = "yav"
	AttachToLangPrimaryRequestVNextPrimaryLanguageYavCm  AttachToLangPrimaryRequestVNextPrimaryLanguage = "yav-cm"
	AttachToLangPrimaryRequestVNextPrimaryLanguageYi     AttachToLangPrimaryRequestVNextPrimaryLanguage = "yi"
	AttachToLangPrimaryRequestVNextPrimaryLanguageYi001  AttachToLangPrimaryRequestVNextPrimaryLanguage = "yi-001"
	AttachToLangPrimaryRequestVNextPrimaryLanguageYo     AttachToLangPrimaryRequestVNextPrimaryLanguage = "yo"
	AttachToLangPrimaryRequestVNextPrimaryLanguageYoBj   AttachToLangPrimaryRequestVNextPrimaryLanguage = "yo-bj"
	AttachToLangPrimaryRequestVNextPrimaryLanguageYoNg   AttachToLangPrimaryRequestVNextPrimaryLanguage = "yo-ng"
	AttachToLangPrimaryRequestVNextPrimaryLanguageYrl    AttachToLangPrimaryRequestVNextPrimaryLanguage = "yrl"
	AttachToLangPrimaryRequestVNextPrimaryLanguageYrlBr  AttachToLangPrimaryRequestVNextPrimaryLanguage = "yrl-br"
	AttachToLangPrimaryRequestVNextPrimaryLanguageYrlCo  AttachToLangPrimaryRequestVNextPrimaryLanguage = "yrl-co"
	AttachToLangPrimaryRequestVNextPrimaryLanguageYrlVe  AttachToLangPrimaryRequestVNextPrimaryLanguage = "yrl-ve"
	AttachToLangPrimaryRequestVNextPrimaryLanguageYue    AttachToLangPrimaryRequestVNextPrimaryLanguage = "yue"
	AttachToLangPrimaryRequestVNextPrimaryLanguageYueCn  AttachToLangPrimaryRequestVNextPrimaryLanguage = "yue-cn"
	AttachToLangPrimaryRequestVNextPrimaryLanguageYueHk  AttachToLangPrimaryRequestVNextPrimaryLanguage = "yue-hk"
	AttachToLangPrimaryRequestVNextPrimaryLanguageZa     AttachToLangPrimaryRequestVNextPrimaryLanguage = "za"
	AttachToLangPrimaryRequestVNextPrimaryLanguageZgh    AttachToLangPrimaryRequestVNextPrimaryLanguage = "zgh"
	AttachToLangPrimaryRequestVNextPrimaryLanguageZghMa  AttachToLangPrimaryRequestVNextPrimaryLanguage = "zgh-ma"
	AttachToLangPrimaryRequestVNextPrimaryLanguageZh     AttachToLangPrimaryRequestVNextPrimaryLanguage = "zh"
	AttachToLangPrimaryRequestVNextPrimaryLanguageZhCn   AttachToLangPrimaryRequestVNextPrimaryLanguage = "zh-cn"
	AttachToLangPrimaryRequestVNextPrimaryLanguageZhHans AttachToLangPrimaryRequestVNextPrimaryLanguage = "zh-hans"
	AttachToLangPrimaryRequestVNextPrimaryLanguageZhHant AttachToLangPrimaryRequestVNextPrimaryLanguage = "zh-hant"
	AttachToLangPrimaryRequestVNextPrimaryLanguageZhHk   AttachToLangPrimaryRequestVNextPrimaryLanguage = "zh-hk"
	AttachToLangPrimaryRequestVNextPrimaryLanguageZhMo   AttachToLangPrimaryRequestVNextPrimaryLanguage = "zh-mo"
	AttachToLangPrimaryRequestVNextPrimaryLanguageZhSg   AttachToLangPrimaryRequestVNextPrimaryLanguage = "zh-sg"
	AttachToLangPrimaryRequestVNextPrimaryLanguageZhTw   AttachToLangPrimaryRequestVNextPrimaryLanguage = "zh-tw"
	AttachToLangPrimaryRequestVNextPrimaryLanguageZu     AttachToLangPrimaryRequestVNextPrimaryLanguage = "zu"
	AttachToLangPrimaryRequestVNextPrimaryLanguageZuZa   AttachToLangPrimaryRequestVNextPrimaryLanguage = "zu-za"
)

type BackgroundImage struct {
	// Defines the position of the background image.
	BackgroundPosition string `json:"backgroundPosition" api:"required"`
	// Specifies the size of the background image.
	BackgroundSize string `json:"backgroundSize" api:"required"`
	// The URL of the background image.
	ImageURL string `json:"imageUrl" api:"required"`
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
	// Defines the position of the background image.
	BackgroundPosition string `json:"backgroundPosition" api:"required"`
	// Specifies the size of the background image.
	BackgroundSize string `json:"backgroundSize" api:"required"`
	// The URL of the background image.
	ImageURL string `json:"imageUrl" api:"required"`
	paramObj
}

func (r BackgroundImageParam) MarshalJSON() (data []byte, err error) {
	type shadow BackgroundImageParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BackgroundImageParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Inputs is required.
type BatchInputJsonNodeParam struct {
	// JSON nodes to input.
	Inputs []any `json:"inputs,omitzero" api:"required"`
	paramObj
}

func (r BatchInputJsonNodeParam) MarshalJSON() (data []byte, err error) {
	type shadow BatchInputJsonNodeParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BatchInputJsonNodeParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BreakpointStyles struct {
	// Boolean indicating if the breakpoint is visible.
	Hidden  bool    `json:"hidden" api:"required"`
	Margin  Margin  `json:"margin" api:"required"`
	Padding Padding `json:"padding" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Hidden      respjson.Field
		Margin      respjson.Field
		Padding     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BreakpointStyles) RawJSON() string { return r.JSON.raw }
func (r *BreakpointStyles) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this BreakpointStyles to a BreakpointStylesParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// BreakpointStylesParam.Overrides()
func (r BreakpointStyles) ToParam() BreakpointStylesParam {
	return param.Override[BreakpointStylesParam](json.RawMessage(r.RawJSON()))
}

// The properties Hidden, Margin, Padding are required.
type BreakpointStylesParam struct {
	// Boolean indicating if the breakpoint is visible.
	Hidden  bool         `json:"hidden" api:"required"`
	Margin  MarginParam  `json:"margin,omitzero" api:"required"`
	Padding PaddingParam `json:"padding,omitzero" api:"required"`
	paramObj
}

func (r BreakpointStylesParam) MarshalJSON() (data []byte, err error) {
	type shadow BreakpointStylesParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BreakpointStylesParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ColorStop struct {
	Color RgbaColor `json:"color" api:"required"`
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
	Color RgbaColorParam `json:"color,omitzero" api:"required"`
	paramObj
}

func (r ColorStopParam) MarshalJSON() (data []byte, err error) {
	type shadow ColorStopParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ColorStopParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property ID is required.
type ContentCloneRequestVNextParam struct {
	// ID of the object to be cloned.
	ID string `json:"id" api:"required"`
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

type ContentLanguageVariation struct {
	// ID of object to set as primary in multi-language group.
	ID int64 `json:"id" api:"required"`
	// If True, the variant will not show up in your dashboard, although the post could
	// still be live.
	ArchivedInDashboard bool `json:"archivedInDashboard" api:"required"`
	// The name of the user who last published the blog post. For posts that haven't
	// been published yet, this property will reflect the user who initially created
	// the draft.
	AuthorName string `json:"authorName" api:"required"`
	// The GUID of the marketing campaign this page is a part of.
	Campaign string `json:"campaign" api:"required"`
	// Name of the associated marketing campaign.
	CampaignName string `json:"campaignName" api:"required"`
	// The timestamp (ISO8601 format) when this Blog Post was created.
	Created time.Time `json:"created" api:"required" format:"date-time"`
	// The internal name of the content language variation.
	Name string `json:"name" api:"required"`
	// Set this to create a password protected page. Entering the password will be
	// required to view the page.
	Password          string             `json:"password" api:"required"`
	PublicAccessRules []PublicAccessRule `json:"publicAccessRules" api:"required"`
	// Boolean to determine whether or not to respect publicAccessRules.
	PublicAccessRulesEnabled bool `json:"publicAccessRulesEnabled" api:"required"`
	// The date (ISO8601 format) the page is to be published at.
	PublishDate time.Time `json:"publishDate" api:"required" format:"date-time"`
	// The path of the this page. This field is appended to the domain to construct the
	// url of this page.
	Slug string `json:"slug" api:"required"`
	// An ENUM describing the current state of this page.
	//
	// Maximum string length: 25
	State string `json:"state" api:"required"`
	// The timestamp (ISO8601 format) when this Blog Post was updated.
	Updated time.Time `json:"updated" api:"required" format:"date-time"`
	TagIDs  []int64   `json:"tagIds"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                       respjson.Field
		ArchivedInDashboard      respjson.Field
		AuthorName               respjson.Field
		Campaign                 respjson.Field
		CampaignName             respjson.Field
		Created                  respjson.Field
		Name                     respjson.Field
		Password                 respjson.Field
		PublicAccessRules        respjson.Field
		PublicAccessRulesEnabled respjson.Field
		PublishDate              respjson.Field
		Slug                     respjson.Field
		State                    respjson.Field
		Updated                  respjson.Field
		TagIDs                   respjson.Field
		ExtraFields              map[string]respjson.Field
		raw                      string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ContentLanguageVariation) RawJSON() string { return r.JSON.raw }
func (r *ContentLanguageVariation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this ContentLanguageVariation to a
// ContentLanguageVariationParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// ContentLanguageVariationParam.Overrides()
func (r ContentLanguageVariation) ToParam() ContentLanguageVariationParam {
	return param.Override[ContentLanguageVariationParam](json.RawMessage(r.RawJSON()))
}

// The properties ID, ArchivedInDashboard, AuthorName, Campaign, CampaignName,
// Created, Name, Password, PublicAccessRules, PublicAccessRulesEnabled,
// PublishDate, Slug, State, Updated are required.
type ContentLanguageVariationParam struct {
	// ID of object to set as primary in multi-language group.
	ID int64 `json:"id" api:"required"`
	// If True, the variant will not show up in your dashboard, although the post could
	// still be live.
	ArchivedInDashboard bool `json:"archivedInDashboard" api:"required"`
	// The name of the user who last published the blog post. For posts that haven't
	// been published yet, this property will reflect the user who initially created
	// the draft.
	AuthorName string `json:"authorName" api:"required"`
	// The GUID of the marketing campaign this page is a part of.
	Campaign string `json:"campaign" api:"required"`
	// Name of the associated marketing campaign.
	CampaignName string `json:"campaignName" api:"required"`
	// The timestamp (ISO8601 format) when this Blog Post was created.
	Created time.Time `json:"created" api:"required" format:"date-time"`
	// The internal name of the content language variation.
	Name string `json:"name" api:"required"`
	// Set this to create a password protected page. Entering the password will be
	// required to view the page.
	Password          string             `json:"password" api:"required"`
	PublicAccessRules []PublicAccessRule `json:"publicAccessRules,omitzero" api:"required"`
	// Boolean to determine whether or not to respect publicAccessRules.
	PublicAccessRulesEnabled bool `json:"publicAccessRulesEnabled" api:"required"`
	// The date (ISO8601 format) the page is to be published at.
	PublishDate time.Time `json:"publishDate" api:"required" format:"date-time"`
	// The path of the this page. This field is appended to the domain to construct the
	// url of this page.
	Slug string `json:"slug" api:"required"`
	// An ENUM describing the current state of this page.
	//
	// Maximum string length: 25
	State string `json:"state" api:"required"`
	// The timestamp (ISO8601 format) when this Blog Post was updated.
	Updated time.Time `json:"updated" api:"required" format:"date-time"`
	TagIDs  []int64   `json:"tagIds,omitzero"`
	paramObj
}

func (r ContentLanguageVariationParam) MarshalJSON() (data []byte, err error) {
	type shadow ContentLanguageVariationParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ContentLanguageVariationParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties ID, PublishDate are required.
type ContentScheduleRequestVNextParam struct {
	// The ID of the object to be scheduled.
	ID string `json:"id" api:"required"`
	// The date the object should transition from scheduled to published.
	PublishDate time.Time `json:"publishDate" api:"required" format:"date-time"`
	paramObj
}

func (r ContentScheduleRequestVNextParam) MarshalJSON() (data []byte, err error) {
	type shadow ContentScheduleRequestVNextParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ContentScheduleRequestVNextParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property ID is required.
type DetachFromLangGroupRequestVNextParam struct {
	// ID of the object to remove from a multi-language group.
	ID string `json:"id" api:"required"`
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
	Angle        Angle        `json:"angle" api:"required"`
	Colors       []ColorStop  `json:"colors" api:"required"`
	SideOrCorner SideOrCorner `json:"sideOrCorner" api:"required"`
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
	Angle        AngleParam        `json:"angle,omitzero" api:"required"`
	Colors       []ColorStopParam  `json:"colors,omitzero" api:"required"`
	SideOrCorner SideOrCornerParam `json:"sideOrCorner,omitzero" api:"required"`
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
	Cells []LayoutSection `json:"cells" api:"required"`
	// The CSS class applied to the layout section.
	CssClass string `json:"cssClass" api:"required"`
	// The CSS ID applied to the layout section.
	CssID string `json:"cssId" api:"required"`
	// Custom CSS styles applied to the layout section.
	CssStyle string `json:"cssStyle" api:"required"`
	// The label for the layout section.
	Label string `json:"label" api:"required"`
	// The name assigned to the layout section.
	Name string `json:"name" api:"required"`
	// null
	Params      map[string]any             `json:"params" api:"required"`
	RowMetaData []RowMetaData              `json:"rowMetaData" api:"required"`
	Rows        []map[string]LayoutSection `json:"rows" api:"required"`
	Styles      Styles                     `json:"styles" api:"required"`
	// The type of the layout section.
	Type string `json:"type" api:"required"`
	// The width of the layout section.
	W int64 `json:"w" api:"required"`
	// The x-coordinate position of the layout section.
	X int64 `json:"x" api:"required"`
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
	Cells []LayoutSectionParam `json:"cells,omitzero" api:"required"`
	// The CSS class applied to the layout section.
	CssClass string `json:"cssClass" api:"required"`
	// The CSS ID applied to the layout section.
	CssID string `json:"cssId" api:"required"`
	// Custom CSS styles applied to the layout section.
	CssStyle string `json:"cssStyle" api:"required"`
	// The label for the layout section.
	Label string `json:"label" api:"required"`
	// The name assigned to the layout section.
	Name string `json:"name" api:"required"`
	// null
	Params      map[string]any                  `json:"params,omitzero" api:"required"`
	RowMetaData []RowMetaDataParam              `json:"rowMetaData,omitzero" api:"required"`
	Rows        []map[string]LayoutSectionParam `json:"rows,omitzero" api:"required"`
	Styles      StylesParam                     `json:"styles,omitzero" api:"required"`
	// The type of the layout section.
	Type string `json:"type" api:"required"`
	// The width of the layout section.
	W int64 `json:"w" api:"required"`
	// The x-coordinate position of the layout section.
	X int64 `json:"x" api:"required"`
	paramObj
}

func (r LayoutSectionParam) MarshalJSON() (data []byte, err error) {
	type shadow LayoutSectionParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LayoutSectionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Margin struct {
	Bottom Size `json:"bottom" api:"required"`
	Top    Size `json:"top" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Bottom      respjson.Field
		Top         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Margin) RawJSON() string { return r.JSON.raw }
func (r *Margin) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this Margin to a MarginParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// MarginParam.Overrides()
func (r Margin) ToParam() MarginParam {
	return param.Override[MarginParam](json.RawMessage(r.RawJSON()))
}

// The properties Bottom, Top are required.
type MarginParam struct {
	Bottom SizeParam `json:"bottom,omitzero" api:"required"`
	Top    SizeParam `json:"top,omitzero" api:"required"`
	paramObj
}

func (r MarginParam) MarshalJSON() (data []byte, err error) {
	type shadow MarginParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MarginParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Padding struct {
	Bottom Size `json:"bottom" api:"required"`
	Left   Size `json:"left" api:"required"`
	Right  Size `json:"right" api:"required"`
	Top    Size `json:"top" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Bottom      respjson.Field
		Left        respjson.Field
		Right       respjson.Field
		Top         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Padding) RawJSON() string { return r.JSON.raw }
func (r *Padding) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this Padding to a PaddingParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// PaddingParam.Overrides()
func (r Padding) ToParam() PaddingParam {
	return param.Override[PaddingParam](json.RawMessage(r.RawJSON()))
}

// The properties Bottom, Left, Right, Top are required.
type PaddingParam struct {
	Bottom SizeParam `json:"bottom,omitzero" api:"required"`
	Left   SizeParam `json:"left,omitzero" api:"required"`
	Right  SizeParam `json:"right,omitzero" api:"required"`
	Top    SizeParam `json:"top,omitzero" api:"required"`
	paramObj
}

func (r PaddingParam) MarshalJSON() (data []byte, err error) {
	type shadow PaddingParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PaddingParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PublicAccessRule = any

type RgbaColor struct {
	// Alpha.
	A float64 `json:"a" api:"required"`
	// Blue.
	B int64 `json:"b" api:"required"`
	// Green.
	G int64 `json:"g" api:"required"`
	// Red.
	R int64 `json:"r" api:"required"`
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

// The properties A, B, G, R are required.
type RgbaColorParam struct {
	// Alpha.
	A float64 `json:"a" api:"required"`
	// Blue.
	B int64 `json:"b" api:"required"`
	// Green.
	G int64 `json:"g" api:"required"`
	// Red.
	R int64 `json:"r" api:"required"`
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
	// The CSS class applied to the row.
	CssClass string `json:"cssClass" api:"required"`
	Styles   Styles `json:"styles" api:"required"`
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
	// The CSS class applied to the row.
	CssClass string      `json:"cssClass" api:"required"`
	Styles   StylesParam `json:"styles,omitzero" api:"required"`
	paramObj
}

func (r RowMetaDataParam) MarshalJSON() (data []byte, err error) {
	type shadow RowMetaDataParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RowMetaDataParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property ID is required.
type SetNewLanguagePrimaryRequestVNextParam struct {
	// ID of object to set as primary in multi-language group.
	ID string `json:"id" api:"required"`
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
	// Specifies the horizontal side of an element.
	//
	// Any of "CENTER", "LEFT", "RIGHT".
	HorizontalSide SideOrCornerHorizontalSide `json:"horizontalSide" api:"required"`
	// Specifies the vertical side of an element.
	//
	// Any of "BOTTOM", "MIDDLE", "TOP".
	VerticalSide SideOrCornerVerticalSide `json:"verticalSide" api:"required"`
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

// Specifies the horizontal side of an element.
type SideOrCornerHorizontalSide string

const (
	SideOrCornerHorizontalSideCenter SideOrCornerHorizontalSide = "CENTER"
	SideOrCornerHorizontalSideLeft   SideOrCornerHorizontalSide = "LEFT"
	SideOrCornerHorizontalSideRight  SideOrCornerHorizontalSide = "RIGHT"
)

// Specifies the vertical side of an element.
type SideOrCornerVerticalSide string

const (
	SideOrCornerVerticalSideBottom SideOrCornerVerticalSide = "BOTTOM"
	SideOrCornerVerticalSideMiddle SideOrCornerVerticalSide = "MIDDLE"
	SideOrCornerVerticalSideTop    SideOrCornerVerticalSide = "TOP"
)

// The properties HorizontalSide, VerticalSide are required.
type SideOrCornerParam struct {
	// Specifies the horizontal side of an element.
	//
	// Any of "CENTER", "LEFT", "RIGHT".
	HorizontalSide SideOrCornerHorizontalSide `json:"horizontalSide,omitzero" api:"required"`
	// Specifies the vertical side of an element.
	//
	// Any of "BOTTOM", "MIDDLE", "TOP".
	VerticalSide SideOrCornerVerticalSide `json:"verticalSide,omitzero" api:"required"`
	paramObj
}

func (r SideOrCornerParam) MarshalJSON() (data []byte, err error) {
	type shadow SideOrCornerParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SideOrCornerParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Size struct {
	// Any of "%", "ch", "cm", "em", "ex", "in", "lh", "mm", "pc", "pt", "px", "Q",
	// "rem", "vh", "vmax", "vmin", "vw".
	Units SizeUnits `json:"units" api:"required"`
	Value float64   `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Units       respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Size) RawJSON() string { return r.JSON.raw }
func (r *Size) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this Size to a SizeParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// SizeParam.Overrides()
func (r Size) ToParam() SizeParam {
	return param.Override[SizeParam](json.RawMessage(r.RawJSON()))
}

type SizeUnits string

const (
	SizeUnitsPercent SizeUnits = "%"
	SizeUnitsCh      SizeUnits = "ch"
	SizeUnitsCm      SizeUnits = "cm"
	SizeUnitsEm      SizeUnits = "em"
	SizeUnitsEx      SizeUnits = "ex"
	SizeUnitsIn      SizeUnits = "in"
	SizeUnitsLh      SizeUnits = "lh"
	SizeUnitsMm      SizeUnits = "mm"
	SizeUnitsPc      SizeUnits = "pc"
	SizeUnitsPt      SizeUnits = "pt"
	SizeUnitsPx      SizeUnits = "px"
	SizeUnitsQ       SizeUnits = "Q"
	SizeUnitsRem     SizeUnits = "rem"
	SizeUnitsVh      SizeUnits = "vh"
	SizeUnitsVmax    SizeUnits = "vmax"
	SizeUnitsVmin    SizeUnits = "vmin"
	SizeUnitsVw      SizeUnits = "vw"
)

// The properties Units, Value are required.
type SizeParam struct {
	// Any of "%", "ch", "cm", "em", "ex", "in", "lh", "mm", "pc", "pt", "px", "Q",
	// "rem", "vh", "vmax", "vmin", "vw".
	Units SizeUnits `json:"units,omitzero" api:"required"`
	Value float64   `json:"value" api:"required"`
	paramObj
}

func (r SizeParam) MarshalJSON() (data []byte, err error) {
	type shadow SizeParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SizeParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Styles struct {
	BackgroundColor    RgbaColor       `json:"backgroundColor" api:"required"`
	BackgroundGradient Gradient        `json:"backgroundGradient" api:"required"`
	BackgroundImage    BackgroundImage `json:"backgroundImage" api:"required"`
	// Indicates whether flexbox positioning is enabled for the section.
	//
	// Any of "BOTTOM_CENTER", "BOTTOM_LEFT", "BOTTOM_RIGHT", "MIDDLE_CENTER",
	// "MIDDLE_LEFT", "MIDDLE_RIGHT", "TOP_CENTER", "TOP_LEFT", "TOP_RIGHT".
	FlexboxPositioning StylesFlexboxPositioning `json:"flexboxPositioning" api:"required"`
	// Determines if the section should be forced to full width.
	ForceFullWidthSection bool `json:"forceFullWidthSection" api:"required"`
	// Defines the maximum width for centering the section.
	MaxWidthSectionCentering int64 `json:"maxWidthSectionCentering" api:"required"`
	// Specifies the vertical alignment of elements within the section.
	//
	// Any of "BOTTOM", "MIDDLE", "TOP".
	VerticalAlignment StylesVerticalAlignment `json:"verticalAlignment" api:"required"`
	// Breakpoint CSS styles for margin, padding, etc...
	BreakpointStyles map[string]BreakpointStyles `json:"breakpointStyles"`
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

// Indicates whether flexbox positioning is enabled for the section.
type StylesFlexboxPositioning string

const (
	StylesFlexboxPositioningBottomCenter StylesFlexboxPositioning = "BOTTOM_CENTER"
	StylesFlexboxPositioningBottomLeft   StylesFlexboxPositioning = "BOTTOM_LEFT"
	StylesFlexboxPositioningBottomRight  StylesFlexboxPositioning = "BOTTOM_RIGHT"
	StylesFlexboxPositioningMiddleCenter StylesFlexboxPositioning = "MIDDLE_CENTER"
	StylesFlexboxPositioningMiddleLeft   StylesFlexboxPositioning = "MIDDLE_LEFT"
	StylesFlexboxPositioningMiddleRight  StylesFlexboxPositioning = "MIDDLE_RIGHT"
	StylesFlexboxPositioningTopCenter    StylesFlexboxPositioning = "TOP_CENTER"
	StylesFlexboxPositioningTopLeft      StylesFlexboxPositioning = "TOP_LEFT"
	StylesFlexboxPositioningTopRight     StylesFlexboxPositioning = "TOP_RIGHT"
)

// Specifies the vertical alignment of elements within the section.
type StylesVerticalAlignment string

const (
	StylesVerticalAlignmentBottom StylesVerticalAlignment = "BOTTOM"
	StylesVerticalAlignmentMiddle StylesVerticalAlignment = "MIDDLE"
	StylesVerticalAlignmentTop    StylesVerticalAlignment = "TOP"
)

// The properties BackgroundColor, BackgroundGradient, BackgroundImage,
// FlexboxPositioning, ForceFullWidthSection, MaxWidthSectionCentering,
// VerticalAlignment are required.
type StylesParam struct {
	BackgroundColor    RgbaColorParam       `json:"backgroundColor,omitzero" api:"required"`
	BackgroundGradient GradientParam        `json:"backgroundGradient,omitzero" api:"required"`
	BackgroundImage    BackgroundImageParam `json:"backgroundImage,omitzero" api:"required"`
	// Indicates whether flexbox positioning is enabled for the section.
	//
	// Any of "BOTTOM_CENTER", "BOTTOM_LEFT", "BOTTOM_RIGHT", "MIDDLE_CENTER",
	// "MIDDLE_LEFT", "MIDDLE_RIGHT", "TOP_CENTER", "TOP_LEFT", "TOP_RIGHT".
	FlexboxPositioning StylesFlexboxPositioning `json:"flexboxPositioning,omitzero" api:"required"`
	// Determines if the section should be forced to full width.
	ForceFullWidthSection bool `json:"forceFullWidthSection" api:"required"`
	// Defines the maximum width for centering the section.
	MaxWidthSectionCentering int64 `json:"maxWidthSectionCentering" api:"required"`
	// Specifies the vertical alignment of elements within the section.
	//
	// Any of "BOTTOM", "MIDDLE", "TOP".
	VerticalAlignment StylesVerticalAlignment `json:"verticalAlignment,omitzero" api:"required"`
	// Breakpoint CSS styles for margin, padding, etc...
	BreakpointStyles map[string]BreakpointStylesParam `json:"breakpointStyles,omitzero"`
	paramObj
}

func (r StylesParam) MarshalJSON() (data []byte, err error) {
	type shadow StylesParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *StylesParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Languages, PrimaryID are required.
type UpdateLanguagesRequestVNextParam struct {
	// Map of object IDs to associated languages of object in the multi-language group.
	//
	// Any of "aa", "ab", "ae", "af", "af-na", "af-za", "agq", "agq-cm", "ak", "ak-gh",
	// "am", "am-et", "an", "ann", "ann-ng", "ar", "ar-001", "ar-ae", "ar-bh", "ar-dj",
	// "ar-dz", "ar-eg", "ar-eh", "ar-er", "ar-il", "ar-iq", "ar-jo", "ar-km", "ar-kw",
	// "ar-lb", "ar-ly", "ar-ma", "ar-mr", "ar-om", "ar-ps", "ar-qa", "ar-sa", "ar-sd",
	// "ar-so", "ar-ss", "ar-sy", "ar-td", "ar-tn", "ar-ye", "as", "asa", "asa-tz",
	// "ast", "ast-es", "as-in", "av", "ay", "az", "az-az", "ba", "bas", "bas-cm",
	// "be", "bem", "bem-zm", "bez", "bez-tz", "be-by", "bg", "bgc", "bgc-in", "bg-bg",
	// "bi", "bho", "bho-in", "bm", "bm-ml", "bn", "bn-bd", "bn-in", "bo", "bo-cn",
	// "bo-in", "br", "brx", "brx-in", "br-fr", "bs", "bs-ba", "ca", "ca-ad", "ca-es",
	// "ca-fr", "ca-it", "ccp", "ccp-bd", "ccp-in", "ce", "ceb", "ceb-ph", "ce-ru",
	// "ch", "cgg", "cgg-ug", "chr", "chr-us", "ckb", "ckb-iq", "ckb-ir", "co", "cr",
	// "cs", "cs-cz", "cu", "cu-ru", "cv", "cv-ru", "cy", "cy-gb", "da", "dav",
	// "dav-ke", "da-dk", "da-gl", "de", "de-at", "de-be", "de-ch", "de-de", "de-gr",
	// "de-it", "de-li", "de-lu", "dje", "dje-ne", "doi", "doi-in", "dsb", "dsb-de",
	// "dua", "dua-cm", "dyo", "dyo-sn", "dv", "dz", "dz-bt", "ebu", "ebu-ke", "ee",
	// "ee-gh", "ee-tg", "el", "el-cy", "el-gr", "en", "en-001", "en-150", "en-ae",
	// "en-ag", "en-ai", "en-as", "en-at", "en-au", "en-bb", "en-be", "en-bi", "en-bm",
	// "en-bs", "en-bw", "en-bz", "en-ca", "en-cc", "en-ch", "en-ck", "en-cm", "en-cn",
	// "en-cx", "en-cy", "en-de", "en-dg", "en-dk", "en-dm", "en-ee", "en-eg", "en-er",
	// "en-es", "en-fi", "en-fj", "en-fk", "en-fm", "en-fr", "en-gb", "en-gd", "en-gg",
	// "en-gh", "en-gi", "en-gm", "en-gu", "en-gy", "en-hk", "en-id", "en-ie", "en-il",
	// "en-im", "en-in", "en-io", "en-je", "en-jm", "en-ke", "en-ki", "en-kn", "en-ky",
	// "en-lc", "en-lr", "en-ls", "en-lu", "en-mg", "en-mh", "en-mo", "en-mp", "en-ms",
	// "en-mt", "en-mu", "en-mv", "en-mw", "en-mx", "en-my", "en-na", "en-nf", "en-ng",
	// "en-nl", "en-nr", "en-nu", "en-nz", "en-pg", "en-ph", "en-pk", "en-pn", "en-pr",
	// "en-pt", "en-pw", "en-rw", "en-sb", "en-sc", "en-sd", "en-se", "en-sg", "en-sh",
	// "en-si", "en-sl", "en-ss", "en-sx", "en-sz", "en-tc", "en-th", "en-tk", "en-tn",
	// "en-to", "en-tt", "en-tv", "en-tz", "en-ug", "en-um", "en-us", "en-vc", "en-vg",
	// "en-vi", "en-vn", "en-vu", "en-ws", "en-za", "en-zm", "en-zw", "eo", "eo-001",
	// "es", "es-419", "es-ar", "es-bo", "es-br", "es-bz", "es-cl", "es-co", "es-cr",
	// "es-cu", "es-do", "es-ea", "es-ec", "es-es", "es-gq", "es-gt", "es-hn", "es-ic",
	// "es-mx", "es-ni", "es-pa", "es-pe", "es-ph", "es-pr", "es-py", "es-sv", "es-us",
	// "es-uy", "es-ve", "et", "et-ee", "eu", "eu-es", "ewo", "ewo-cm", "fa", "fa-af",
	// "fa-ir", "ff", "ff-bf", "ff-cm", "ff-gh", "ff-gm", "ff-gn", "ff-gw", "ff-lr",
	// "ff-mr", "ff-ne", "ff-ng", "ff-sl", "ff-sn", "fi", "fil", "fil-ph", "fi-fi",
	// "fj", "fo", "fo-dk", "fo-fo", "fr", "frr", "frr-de", "fr-be", "fr-bf", "fr-bi",
	// "fr-bj", "fr-bl", "fr-ca", "fr-cd", "fr-cf", "fr-cg", "fr-ch", "fr-ci", "fr-cm",
	// "fr-dj", "fr-dz", "fr-fr", "fr-ga", "fr-gf", "fr-gn", "fr-gp", "fr-gq", "fr-ht",
	// "fr-km", "fr-lu", "fr-ma", "fr-mc", "fr-mf", "fr-mg", "fr-ml", "fr-mq", "fr-mr",
	// "fr-mu", "fr-nc", "fr-ne", "fr-pf", "fr-pm", "fr-re", "fr-rw", "fr-sc", "fr-sn",
	// "fr-sy", "fr-td", "fr-tg", "fr-tn", "fr-vu", "fr-wf", "fr-yt", "fur", "fur-it",
	// "fy", "fy-nl", "ga", "ga-gb", "ga-ie", "gd", "gd-gb", "gl", "gl-es", "gn",
	// "gsw", "gsw-ch", "gsw-fr", "gsw-li", "gu", "guz", "guz-ke", "gu-in", "gv",
	// "gv-im", "ha", "haw", "haw-us", "ha-gh", "ha-ne", "ha-ng", "he", "he-il", "hi",
	// "hi-in", "hmn", "ho", "hr", "hr-ba", "hr-hr", "ht", "hsb", "hsb-de", "hu",
	// "hu-hu", "hy", "hy-am", "hz", "ia", "ia-001", "id", "ie", "ig", "ig-ng", "ii",
	// "ii-cn", "ik", "io", "id-id", "is", "is-is", "it", "it-ch", "it-it", "it-sm",
	// "it-va", "iu", "ja", "ja-jp", "jgo", "jgo-cm", "yi", "yi-001", "jmc", "jmc-tz",
	// "jv", "jv-id", "ka", "kab", "kab-dz", "kam", "kam-ke", "kar", "ka-ge", "kde",
	// "kde-tz", "kea", "kea-cv", "kgp", "kgp-br", "kg", "kh", "khq", "khq-ml", "ki",
	// "ki-ke", "kj", "kk", "kkj", "kkj-cm", "kk-kz", "kl", "kln", "kln-ke", "kl-gl",
	// "km", "km-kh", "kn", "kn-in", "ko", "kok", "kok-in", "ko-kp", "ko-kr", "kr",
	// "ks", "ksb", "ksb-tz", "ksf", "ksf-cm", "ksh", "ksh-de", "ks-in", "ku", "ku-tr",
	// "kv", "kw", "kw-gb", "ky", "ky-kg", "lag", "lag-tz", "la", "lb", "lb-lu", "lg",
	// "lg-ug", "lkt", "lkt-us", "li", "ln", "ln-ao", "ln-cd", "ln-cf", "ln-cg", "lo",
	// "lo-la", "lrc", "lrc-iq", "lrc-ir", "lt", "lt-lt", "lu", "luo", "luo-ke", "luy",
	// "luy-ke", "lu-cd", "lv", "lv-lv", "mai", "mai-in", "mas", "mas-ke", "mas-tz",
	// "mdf", "mdf-ru", "mer", "mer-ke", "mfe", "mfe-mu", "mg", "mgh", "mgh-mz", "mgo",
	// "mgo-cm", "mg-mg", "mh", "mi", "mi-nz", "mk", "mk-mk", "ml", "ml-in", "mn",
	// "mni", "mni-in", "mn-mn", "mr", "mr-in", "ms", "ms-bn", "ms-id", "ms-my",
	// "ms-sg", "mt", "mt-mt", "mua", "mua-cm", "my", "my-mm", "mzn", "mzn-ir", "naq",
	// "naq-na", "na", "nb", "nb-no", "nb-sj", "nd", "nds", "nds-de", "nds-nl",
	// "nd-zw", "ne", "ne-in", "ne-np", "ng", "nl", "nl-aw", "nl-be", "nl-bq", "nl-ch",
	// "nl-cw", "nl-lu", "nl-nl", "nl-sr", "nl-sx", "nmg", "nmg-cm", "nn", "nnh",
	// "nnh-cm", "nn-no", "nr", "nv", "ny", "no", "no-no", "nus", "nus-ss", "nyn",
	// "nyn-ug", "oc", "oc-es", "oc-fr", "oj", "om", "om-et", "om-ke", "or", "or-in",
	// "os", "os-ge", "os-ru", "pa", "pa-in", "pa-pk", "pcm", "pcm-ng", "pis",
	// "pis-sb", "pi", "pl", "pl-pl", "prg", "prg-001", "ps", "ps-af", "ps-pk", "pt",
	// "pt-ao", "pt-br", "pt-ch", "pt-cv", "pt-gq", "pt-gw", "pt-lu", "pt-mo", "pt-mz",
	// "pt-pt", "pt-st", "pt-tl", "qu", "qu-bo", "qu-ec", "qu-pe", "raj", "raj-in",
	// "rm", "rm-ch", "rn", "rn-bi", "ro", "rof", "rof-tz", "ro-md", "ro-ro", "ru",
	// "ru-by", "ru-kg", "ru-kz", "ru-md", "ru-ru", "ru-ua", "rw", "rwk", "rwk-tz",
	// "rw-rw", "sa", "sah", "sah-ru", "saq", "saq-ke", "sat", "sat-in", "sa-in",
	// "sbp", "sbp-tz", "sc", "sc-it", "sd", "sd-in", "sd-pk", "se", "seh", "seh-mz",
	// "ses", "ses-ml", "se-fi", "se-no", "se-se", "sg", "sg-cf", "shi", "shi-ma",
	// "si", "si-lk", "sk", "sk-sk", "sl", "sl-si", "sm", "smn", "smn-fi", "sms",
	// "sms-fi", "sn", "sn-zw", "so", "so-dj", "so-et", "so-ke", "so-so", "sq",
	// "sq-al", "sq-mk", "sq-xk", "sr", "sr-ba", "sr-cs", "sr-me", "sr-rs", "sr-xk",
	// "ss", "st", "su", "su-id", "sv", "sv-ax", "sv-fi", "sv-se", "sw", "sw-cd",
	// "sw-ke", "sw-tz", "sw-ug", "sy", "ta", "ta-in", "ta-lk", "ta-my", "ta-sg", "te",
	// "teo", "teo-ke", "teo-ug", "te-in", "tg", "tg-tj", "th", "th-th", "ti", "ti-er",
	// "ti-et", "tk", "tk-tm", "tl", "tn", "to", "tok", "tok-001", "to-to", "ts", "tr",
	// "tr-cy", "tr-tr", "tt", "tt-ru", "tw", "ty", "twq", "twq-ne", "tzm", "tzm-ma",
	// "ug", "ug-cn", "uk", "uk-ua", "ur", "ur-in", "ur-pk", "uz", "uz-af", "uz-uz",
	// "vai", "vai-lr", "ve", "vi", "vi-vn", "vo", "vo-001", "vun", "vun-tz", "wa",
	// "wae", "wae-ch", "wo", "wo-sn", "xh", "xh-za", "xog", "xog-ug", "yav", "yav-cm",
	// "yo", "yo-bj", "yo-ng", "yrl", "yrl-br", "yrl-co", "yrl-ve", "yue", "yue-cn",
	// "yue-hk", "zgh", "zgh-ma", "za", "zh", "zh-cn", "zh-hans", "zh-hant", "zh-hk",
	// "zh-mo", "zh-sg", "zh-tw", "zu", "zu-za".
	Languages map[string]string `json:"languages,omitzero" api:"required"`
	// ID of the primary object in the multi-language group.
	PrimaryID string `json:"primaryId" api:"required"`
	paramObj
}

func (r UpdateLanguagesRequestVNextParam) MarshalJSON() (data []byte, err error) {
	type shadow UpdateLanguagesRequestVNextParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UpdateLanguagesRequestVNextParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
