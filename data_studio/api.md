# DataStudio

## Datasource

Params Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/hubspot-sdk-go/data_studio">data_studio</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/hubspot-sdk-go/data_studio#BodyPartParam">BodyPartParam</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/hubspot-sdk-go/data_studio">data_studio</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/hubspot-sdk-go/data_studio#ContentDispositionParam">ContentDispositionParam</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/hubspot-sdk-go/data_studio">data_studio</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/hubspot-sdk-go/data_studio#FormDataBodyPartParam">FormDataBodyPartParam</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/hubspot-sdk-go/data_studio">data_studio</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/hubspot-sdk-go/data_studio#FormDataContentDispositionParam">FormDataContentDispositionParam</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/hubspot-sdk-go/data_studio">data_studio</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/hubspot-sdk-go/data_studio#FormDataMultiPartParam">FormDataMultiPartParam</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/hubspot-sdk-go/data_studio">data_studio</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/hubspot-sdk-go/data_studio#MediaTypeParam">MediaTypeParam</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/hubspot-sdk-go/data_studio">data_studio</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/hubspot-sdk-go/data_studio#MultiPartParam">MultiPartParam</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/hubspot-sdk-go/data_studio">data_studio</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/hubspot-sdk-go/data_studio#ParameterizedHeaderParam">ParameterizedHeaderParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/hubspot-sdk-go/data_studio">data_studio</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/hubspot-sdk-go/data_studio#DataSourceGetResponse">DataSourceGetResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/hubspot-sdk-go/data_studio">data_studio</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/hubspot-sdk-go/data_studio#DataSourceUpdateResponse">DataSourceUpdateResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/hubspot-sdk-go/data_studio">data_studio</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/hubspot-sdk-go/data_studio#FileColumn">FileColumn</a>

Methods:

- <code title="post /data-studio/2026-03/data-source">client.DataStudio.Datasource.<a href="https://pkg.go.dev/github.com/stainless-sdks/hubspot-sdk-go/data_studio#DatasourceService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/hubspot-sdk-go/data_studio">data_studio</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/hubspot-sdk-go/data_studio#DatasourceNewParams">DatasourceNewParams</a>) (\*http.Response, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /data-studio/2026-03/data-source/{datasourceId}">client.DataStudio.Datasource.<a href="https://pkg.go.dev/github.com/stainless-sdks/hubspot-sdk-go/data_studio#DatasourceService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, datasourceID <a href="https://pkg.go.dev/builtin#int64">int64</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/hubspot-sdk-go/data_studio">data_studio</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/hubspot-sdk-go/data_studio#DatasourceUpdateParams">DatasourceUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/hubspot-sdk-go/data_studio">data_studio</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/hubspot-sdk-go/data_studio#DataSourceUpdateResponse">DataSourceUpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /data-studio/2026-03/data-source/{datasourceId}">client.DataStudio.Datasource.<a href="https://pkg.go.dev/github.com/stainless-sdks/hubspot-sdk-go/data_studio#DatasourceService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, datasourceID <a href="https://pkg.go.dev/builtin#int64">int64</a>) (\*http.Response, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /data-studio/2026-03/data-source/{datasourceId}">client.DataStudio.Datasource.<a href="https://pkg.go.dev/github.com/stainless-sdks/hubspot-sdk-go/data_studio#DatasourceService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, datasourceID <a href="https://pkg.go.dev/builtin#int64">int64</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/hubspot-sdk-go/data_studio">data_studio</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/hubspot-sdk-go/data_studio#DataSourceGetResponse">DataSourceGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
