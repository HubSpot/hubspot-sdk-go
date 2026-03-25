// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cms

import (
	"github.com/stainless-sdks/hubspot-sdk-go/option"
)

// CmService contains methods and other services that help with interacting with
// the hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCmService] method instead.
type CmService struct {
	Options []option.RequestOption
	Blogs   BlogService
}

// NewCmService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewCmService(opts ...option.RequestOption) (r CmService) {
	r = CmService{}
	r.Options = opts
	r.Blogs = NewBlogService(opts...)
	return
}
