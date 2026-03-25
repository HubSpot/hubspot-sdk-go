// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package data_studio

import (
	"github.com/stainless-sdks/hubspot-sdk-go/option"
)

// DataStudioService contains methods and other services that help with interacting
// with the hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDataStudioService] method instead.
type DataStudioService struct {
	Options    []option.RequestOption
	Datasource DatasourceService
}

// NewDataStudioService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewDataStudioService(opts ...option.RequestOption) (r DataStudioService) {
	r = DataStudioService{}
	r.Options = opts
	r.Datasource = NewDatasourceService(opts...)
	return
}
