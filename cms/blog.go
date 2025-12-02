// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cms

import (
	"github.com/stainless-sdks/hubspot-sdk-go/option"
)

// BlogService contains methods and other services that help with interacting with
// the hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBlogService] method instead.
type BlogService struct {
	Options  []option.RequestOption
	Authors  BlogAuthorService
	Posts    BlogPostService
	Settings BlogSettingService
	Tags     BlogTagService
}

// NewBlogService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewBlogService(opts ...option.RequestOption) (r BlogService) {
	r = BlogService{}
	r.Options = opts
	r.Authors = NewBlogAuthorService(opts...)
	r.Posts = NewBlogPostService(opts...)
	r.Settings = NewBlogSettingService(opts...)
	r.Tags = NewBlogTagService(opts...)
	return
}
